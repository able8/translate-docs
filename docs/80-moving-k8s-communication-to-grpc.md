# Moving k8s communication to gRPC

2021/03/20 From: https://blog.cloudflare.com/moving-k8s-communication-to-grpc/

Over the past year and a half, Cloudflare has been hard at work moving our  back-end services running in our non-edge locations from bare metal  solutions and Mesos Marathon to a more unified approach using [Kubernetes(K8s)](https://kubernetes.io/). We chose Kubernetes because it allowed us to split up our monolithic  application into many different microservices with granular control of  communication.

For example, a [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) in Kubernetes can provide high availability by ensuring that the correct number of pods are always available. A [Pod](https://kubernetes.io/docs/concepts/workloads/pods/) in Kubernetes is similar to a container in [Docker](https://www.docker.com/). Both are responsible for running the actual application. These pods can then be exposed through a Kubernetes [Service](https://kubernetes.io/docs/concepts/services-networking/service/) to abstract away the number of replicas by providing a single endpoint  that load balances to the pods behind it. The services can then be  exposed to the Internet via an [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/). Lastly, a network policy can protect against unwanted communication by  ensuring the correct policies are applied to the application. These  policies can include L3 or L4 rules.

The diagram below shows a simple example of this setup.

![The diagram below shows a simple example of this setup.](https://blog.cloudflare.com/content/images/2021/03/2-3.png)

Though Kubernetes does an excellent job at providing the tools for  communication and traffic management, it does not help the developer  decide the best way to communicate between the applications running on  the pods. Throughout this blog we will look at some of the decisions we  made and why we made them to discuss the pros and cons of two commonly  used API architectures, REST and gRPC.

### Out with the old, in with the new

When the DNS team first moved to Kubernetes, all of our pod-to-pod  communication was done through REST APIs and in many cases also included Kafka. The general communication flow was as follows:

![When the DNS team first moved to Kubernetes, all of our pod-to-pod communication was done through REST APIs and in many cases also included Kafka. ](https://blog.cloudflare.com/content/images/2021/03/1-5.png)

We use Kafka because it allows us to handle large spikes in volume without losing information. For example, during a Secondary DNS Zone zone  transfer, Service A tells Service B that the zone is ready to be  published to the edge. Service B then calls Service A’s REST API,  generates the zone, and pushes it to the edge. If you want more  information about how this works, I wrote an entire blog post about the [Secondary DNS pipeline](https://blog.cloudflare.com/secondary-dns-deep-dive/) at Cloudflare.

HTTP worked well for most communication between these two services. However, as we scaled up and added new endpoints, we realized that as long as we control both ends of the communication, we could improve the usability  and performance of our communication. In addition, sending large DNS  zones over the network using HTTP often caused issues with sizing  constraints and compression.

In contrast, gRPC can easily stream  data between client and server and is commonly used in microservice  architecture. These qualities made gRPC the obvious replacement for our  REST APIs.

### gRPC Usability

Often  overlooked from a developer’s perspective, HTTP client libraries are  clunky and require code that defines paths, handles parameters, and  deals with responses in bytes. gRPC abstracts all of this away and makes network calls feel like any other function calls defined for a struct.

The example below shows a very basic schema to set up a GRPC client/server system. As a result of gRPC using [protobuf](https://developers.google.com/protocol-buffers) for serialization, it is largely language agnostic. Once a schema is defined, the *protoc* command can be used to generate code for [many languages](https://grpc.io/docs/languages/).

Protocol Buffer data is structured as *messages,* with each *message* containing information stored in the form of fields. The fields are  strongly typed, providing type safety unlike JSON or XML. Two messages  have been defined, *Hello* and *HelloResponse*. Next we define a service called *HelloWorldHandler* which contains one RPC function called *SayHello* that must be implemented if any object wants to call themselves a *HelloWorldHandler*.

Simple Proto:

```js
message Hello{
   string Name = 1;
}

message HelloResponse{}

service HelloWorldHandler {
   rpc SayHello(Hello) returns (HelloResponse){}
}
```

Once we run our *protoc* command, we are ready to write the server-side code. In order to implement the *HelloWorldHandler*, we must define a struct that implements all of the RPC functions specified in the protobuf schema above*.* In this case, the struct *Server* defines a function *SayHello* that takes in two parameters, context and **pb.Hello*. **pb.Hello* was previously specified in the schema and contains one field, *Name. SayHello* must also return the **pbHelloResponse* which has been defined without fields for simplicity.

Inside the main function, we create a TCP listener, create a new gRPC server, and then register our handler as a *HelloWorldHandlerServer*. After calling *Serve* on our gRPC server, clients will be able to communicate with the server through the function *SayHello*.

Simple Server:

```js
type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *pb.Hello) (*pb.HelloResponse, error) {
    fmt.Println("%s says hello\n", in.Name)
    return &pb.HelloResponse{}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    gRPCServer := gRPC.NewServer()
    handler := Server{}
    pb.RegisterHelloWorldHandlerServer(gRPCServer, &handler)
    if err := gRPCServer.Serve(lis); err != nil {
        panic(err)
    }
}
```

Finally, we need to implement the gRPC Client. First, we establish a TCP connection with the server. Then, we create a new *pb.HandlerClient*. The client is able to call the server's *SayHello* function by passing in a **pb.Hello* object.

Simple Client:

```js
conn, err := gRPC.Dial("127.0.0.1:8080", gRPC.WithInsecure())
if err != nil {
    panic(err)
}
client := pb.NewHelloWorldHandlerClient(conn)
client.SayHello(context.Background(), &pb.Hello{Name: "alex"})
```

Though I have removed some code for simplicity, these *services* and *messages* can become quite complex if needed. The most important thing to  understand is that when a server attempts to announce itself as a *HelloWorldHandlerServer*, it is required to implement the RPC functions as specified within the  protobuf schema. This agreement between the client and server makes  cross-language network calls feel like regular function calls.

In addition to the basic Unary server described above, gRPC lets you decide between four types of service methods:

- **Unary** (example above): client sends a single request to the server and gets a single response back, just like a normal function call.
- **Server Streaming:** server returns a stream of messages in response to a client's request.
- **Client Streaming:** client sends a stream of messages to the server and the server replies  in a single message, usually once the client has finished streaming.
- **Bi-directional Streaming:** the client and server can both send streams of messages to each other asynchronously.

### gRPC Performance

Not all HTTP connections are created equal. Though Golang natively supports HTTP/2, the HTTP/2 transport must be set by the client and the server  must also support HTTP/2. Before moving to gRPC, we were still using  HTTP/1.1 for client connections. We could have switched to HTTP/2 for  performance gains, but we would have lost some of the benefits of native protobuf compression and usability changes.

The best option  available in HTTP/1.1 is pipelining. Pipelining means that although  requests can share a connection, they must queue up one after the other  until the request in front completes. HTTP/2 improved pipelining by  using connection multiplexing. Multiplexing allows for multiple requests to be sent on the same connection and at the same time.

HTTP REST APIs generally use JSON for their request and response format. Protobuf is the native request/response format of gRPC because it has a standard schema agreed upon by the client and server during registration. In  addition, protobuf is known to be significantly faster than JSON due to  its serialization speeds. I’ve run some benchmarks on my laptop, source  code can be found [here](https://github.com/Fattouche/protobuf-benchmark).

![Protobuf performs better in small, medium, and large data sizes.](https://blog.cloudflare.com/content/images/2021/03/image1-26.png)

As you can see, protobuf performs better in small, medium, and large data  sizes. It is faster per operation, smaller after marshalling, and scales well with input size. This becomes even more noticeable when  unmarshaling very large data sets. Protobuf takes 96.4ns/op but JSON  takes 22647ns/op, a 235X reduction in time! For large DNS zones, this  efficiency makes a massive difference in the time it takes us to go from record change in our API to serving it at the edge.

Combining the benefits of HTTP/2 and protobuf showed almost no performance change  from our application’s point of view. This is likely due to the fact  that our pods were already so close together that our connection times  were already very low. In addition, most of our gRPC calls are done with small amounts of data where the difference is negligible. One thing  that we did notice **—** likely related to the multiplexing of HTTP/2 **—** was greater efficiency when writing newly created/edited/deleted records to the edge. Our latency spikes dropped in both amplitude and frequency.

![Our latency spikes dropped in both amplitude and frequency.](https://blog.cloudflare.com/content/images/2021/03/image2-19.png)

### gRPC Security

One of the best features in Kubernetes is the NetworkPolicy. This allows developers to control what goes in and what goes out.

```js
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: test-network-policy
  namespace: default
spec:
  podSelector:
    matchLabels:
      role: db
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - ipBlock:
        cidr: 172.17.0.0/16
        except:
        - 172.17.1.0/24
    - namespaceSelector:
        matchLabels:
          project: myproject
    - podSelector:
        matchLabels:
          role: frontend
    ports:
    - protocol: TCP
      port: 6379
  egress:
  - to:
    - ipBlock:
        cidr: 10.0.0.0/24
    ports:
    - protocol: TCP
      port: 5978
```

In this example, taken from the [Kubernetes docs](https://kubernetes.io/docs/concepts/services-networking/network-policies/), we can see that this will create a network policy called  test-network-policy. This policy controls both ingress and egress  communication to or from any pod that matches the role *db* and enforces the following rules:

Ingress connections allowed:

- Any pod in default namespace with label “role=frontend”
- Any pod in any namespace that has a label “project=myproject”
- Any source IP address in 172.17.0.0/16 except for 172.17.1.0/24

Egress connections allowed:

- Any dest IP address in 10.0.0.0/24

NetworkPolicies do a fantastic job of protecting APIs at the network level, however,  they do nothing to protect APIs at the application level. If you wanted  to control which endpoints can be accessed within the API, you would  need k8s to be able to not only distinguish between pods, but also  endpoints within those pods. These concerns led us to [per RPC credentials](https://grpc.io/docs/guides/auth/). Per RPC credentials are easy to set up on top of the pre-existing gRPC  code. All you need to do is add interceptors to both your stream and  unary handlers.

```js
func (s *Server) UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    // Get the targeted function
    functionInfo := strings.Split(info.FullMethod, "/")
    function := functionInfo[len(functionInfo)-1]
    md, _ := metadata.FromIncomingContext(ctx)

    // Authenticate
    err := authenticateClient(md.Get("username")[0], md.Get("password")[0], function)
    // Blocked
    if err != nil {
        return nil, err
    }
    // Verified
    return handler(ctx, req)
}
```

In this example code snippet, we are grabbing the username, password, and requested function from the info object. We then authenticate  against the client to make sure that it has correct rights to call that  function. This interceptor will run before any of the other functions  get called, which means one implementation protects all functions. The  client would initialize its secure connection and send credentials like  so:

```js
transportCreds, err := credentials.NewClientTLSFromFile(certFile, "")
if err != nil {
    return nil, err
}
perRPCCreds := Creds{Password: grpcPassword, User: user}
conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(transportCreds), grpc.WithPerRPCCredentials(perRPCCreds))
if err != nil {
    return nil, err
}
client:= pb.NewRecordHandlerClient(conn)
// Can now start using the client
```

Here the client first verifies that the server matches with the  certFile. This step ensures that the client does not accidentally send  its password to a bad actor. Next, the client initializes the *perRPCCreds* struct with its username and password and dials the server with that  information. Any time the client makes a call to an rpc defined  function, its credentials will be verified by the server.

### Next Steps

Our next step is to remove the need for many applications to access the  database and ultimately DRY up our codebase by pulling all DNS-related  code into a single API, accessed from one gRPC interface. This removes  the potential for mistakes in individual applications and makes updating our database schema easier. It also gives us more granular control over which functions can be accessed rather than which tables can be  accessed.

So far, the DNS team is very happy with the results of  our gRPC migration. However, we still have a long way to go before we  can move entirely away from REST. We are also patiently waiting for [HTTP/3 support](https://github.com/grpc/grpc/issues/19126) for gRPC so that we can take advantage of those super [quic](https://en.wikipedia.org/wiki/QUIC) speeds!
