# Moving k8s communication to gRPC

# 将 k8s 通信转移到 gRPC

2021/03/20 From: https://blog.cloudflare.com/moving-k8s-communication-to-grpc/

Over the past year and a half, Cloudflare has been hard at work moving our  back-end services running in our non-edge locations from bare metal  solutions and Mesos Marathon to a more unified approach using [Kubernetes(K8s)](https:/ /kubernetes.io/). We chose Kubernetes because it allowed us to split up our monolithic application into many different microservices with granular control of communication.

在过去的一年半里，Cloudflare 一直在努力将我们在非边缘位置运行的后端服务从裸机解决方案和 Mesos Marathon 迁移到使用 [Kubernetes(K8s)](https:/ /kubernetes.io/)。我们选择 Kubernetes 是因为它允许我们将单体应用程序拆分为许多不同的微服务，并对通信进行精细控制。

For example, a [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) in Kubernetes can provide high availability by ensuring that the correct number of pods are always available. A [Pod](https://kubernetes.io/docs/concepts/workloads/pods/) in Kubernetes is similar to a container in [Docker](https://www.docker.com/). Both are responsible for running the actual application. These pods can then be exposed through a Kubernetes [Service](https://kubernetes.io/docs/concepts/services-networking/service/) to abstract away the number of replicas by providing a single endpoint  that load balances to the pods behind it. The services can then be  exposed to the Internet via an [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/). Lastly, a network policy can protect against unwanted communication by  ensuring the correct policies are applied to the application. These  policies can include L3 or L4 rules.

例如，Kubernetes 中的 [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) 可以通过确保正确数量的 pod 始终可用来提供高可用性。 Kubernetes 中的 [Pod](https://kubernetes.io/docs/concepts/workloads/pods/) 类似于 [Docker](https://www.docker.com/) 中的容器。两者都负责运行实际的应用程序。然后可以通过 Kubernetes [服务](https://kubernetes.io/docs/concepts/services-networking/service/) 公开这些 pod，通过提供一个负载均衡到 pod 的端点来抽象出副本的数量在它后面。然后可以通过 [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) 将服务暴露给 Internet。最后，网络策略可以通过确保将正确的策略应用于应用程序来防止不需要的通信。这些策略可以包括 L3 或 L4 规则。

The diagram below shows a simple example of this setup.

下图显示了此设置的一个简单示例。

![The diagram below shows a simple example of this setup.](https://blog.cloudflare.com/content/images/2021/03/2-3.png)

Though Kubernetes does an excellent job at providing the tools for  communication and traffic management, it does not help the developer  decide the best way to communicate between the applications running on  the pods. Throughout this blog we will look at some of the decisions we  made and why we made them to discuss the pros and cons of two commonly  used API architectures, REST and gRPC.

尽管 Kubernetes 在提供通信和流量管理工具方面做得非常出色，但它并不能帮助开发人员决定在 pod 上运行的应用程序之间进行通信的最佳方式。在本博客中，我们将介绍我们做出的一些决定以及我们做出这些决定的原因，以讨论两种常用 API 架构 REST 和 gRPC 的优缺点。

### Out with the old, in with the new

### 旧的，新的

When the DNS team first moved to Kubernetes, all of our pod-to-pod  communication was done through REST APIs and in many cases also included Kafka. The general communication flow was as follows:

当 DNS 团队第一次迁移到 Kubernetes 时，我们所有的 pod 到 pod 通信都是通过 REST API 完成的，在许多情况下还包括 Kafka。一般通信流程如下：

![When the DNS team first moved to Kubernetes, all of our pod-to-pod communication was done through REST APIs and in many cases also included Kafka.](https://blog.cloudflare.com/content/images/2021/03/1-5.png)



We use Kafka because it allows us to handle large spikes in volume without losing information. For example, during a Secondary DNS Zone zone  transfer, Service A tells Service B that the zone is ready to be  published to the edge. Service B then calls Service A’s REST API,  generates the zone, and pushes it to the edge. If you want more  information about how this works, I wrote an entire blog post about the [Secondary DNS pipeline](https://blog.cloudflare.com/secondary-dns-deep-dive/) at Cloudflare.

我们使用 Kafka 是因为它允许我们在不丢失信息的情况下处理大量的峰值。例如，在次要 DNS 区域传输期间，服务 A 告诉服务 B 该区域已准备好发布到边缘。然后服务 B 调用服务 A 的 REST API，生成区域，并将其推送到边缘。如果您想了解有关其工作原理的更多信息，我在 Cloudflare 上写了一篇关于 [辅助 DNS 管道](https://blog.cloudflare.com/secondary-dns-deep-dive/) 的完整博客文章。

HTTP worked well for most communication between these two services. However, as we scaled up and added new endpoints, we realized that as long as we control both ends of the communication, we could improve the usability  and performance of our communication. In addition, sending large DNS  zones over the network using HTTP often caused issues with sizing  constraints and compression.

对于这两个服务之间的大多数通信，HTTP 运行良好。然而，随着我们扩大规模和添加新端点，我们意识到只要我们控制通信的两端，我们就可以提高通信的可用性和性能。此外，使用 HTTP 通过网络发送大型 DNS 区域通常会导致大小限制和压缩问题。

In contrast, gRPC can easily stream  data between client and server and is commonly used in microservice  architecture. These qualities made gRPC the obvious replacement for our  REST APIs.

相比之下，gRPC 可以轻松地在客户端和服务器之间传输数据，并且常用于微服务架构中。这些品质使 gRPC 成为我们 REST API 的明显替代品。

### gRPC Usability

### gRPC 可用性

Often  overlooked from a developer’s perspective, HTTP client libraries are  clunky and require code that defines paths, handles parameters, and  deals with responses in bytes. gRPC abstracts all of this away and makes network calls feel like any other function calls defined for a struct.

从开发人员的角度来看，HTTP 客户端库经常被忽视，并且需要以字节为单位定义路径、处理参数和处理响应的代码。 gRPC 将所有这些抽象出来，并使网络调用感觉就像为结构定义的任何其他函数调用一样。

The example below shows a very basic schema to set up a gRPC client/server system. As a result of gRPC using [protobuf](https://developers.google.com/protocol-buffers) for serialization, it is largely language agnostic. Once a schema is defined, the *protoc* command can be used to generate code for [many languages](https://grpc.io/docs/languages/). 

下面的示例显示了一个非常基本的模式来设置 gRPC 客户端/服务器系统。由于 gRPC 使用 [protobuf](https://developers.google.com/protocol-buffers) 进行序列化，它在很大程度上与语言无关。定义模式后，*protoc* 命令可用于为[多种语言](https://grpc.io/docs/languages/) 生成代码。

Protocol Buffer data is structured as *messages,* with each *message* containing information stored in the form of fields. The fields are  strongly typed, providing type safety unlike JSON or XML. Two messages  have been defined, *Hello* and *HelloResponse*. Next we define a service called *HelloWorldHandler* which contains one RPC function called *SayHello* that must be implemented if any object wants to call themselves a *HelloWorldHandler*.

协议缓冲区数据的结构为 *messages*，每个 *message* 包含以字段形式存储的信息。这些字段是强类型的，提供了与 JSON 或 XML 不同的类型安全性。已经定义了两个消息，*Hello* 和 *HelloResponse*。接下来，我们定义一个名为 *HelloWorldHandler* 的服务，其中包含一个名为 *SayHello* 的 RPC 函数，如果任何对象想要将自己称为 *HelloWorldHandler*，则必须实现该函数。

Simple Proto:

简单的原型：

```js
 message Hello{
    string Name = 1;
 }

 message HelloResponse{}

 service HelloWorldHandler {
    rpc SayHello(Hello) returns (HelloResponse){}
 }
```


Once we run our *protoc* command, we are ready to write the server-side code. In order to implement the *HelloWorldHandler*, we must define a struct that implements all of the RPC functions specified in the protobuf schema above*.* In this case, the struct *Server* defines a function *SayHello* that takes in two parameters , context and **pb.Hello*. **pb.Hello* was previously specified in the schema and contains one field, *Name. SayHello* must also return the **pbHelloResponse* which has been defined without fields for simplicity.

一旦我们运行了 *protoc* 命令，我们就可以编写服务器端代码了。为了实现 *HelloWorldHandler*，我们必须定义一个结构体来实现上面 protobuf 模式中指定的所有 RPC 函数*。* 在这种情况下，结构体 *Server* 定义了一个函数 *SayHello*，它接受两个参数，上下文和**pb.Hello*。 **pb.Hello* 之前已在架构中指定并包含一个字段 *Name。 SayHello* 还必须返回 **pbHelloResponse*，为了简单起见，该响应已定义为不带字段。

Inside the main function, we create a TCP listener, create a new gRPC server, and then register our handler as a *HelloWorldHandlerServer*. After calling *Serve* on our gRPC server, clients will be able to communicate with the server through the function *SayHello*.

在 main 函数中，我们创建了一个 TCP 侦听器，创建了一个新的 gRPC 服务器，然后将我们的处理程序注册为 *HelloWorldHandlerServer*。在我们的 gRPC 服务器上调用 *Serve* 后，客户端将能够通过函数 *SayHello* 与服务器通信。

Simple Server:

简单服务器：

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

最后，我们需要实现 gRPC Client。首先，我们与服务器建立 TCP 连接。然后，我们创建一个新的 *pb.HandlerClient*。客户端可以通过传入一个 **pb.Hello* 对象来调用服务器的 *SayHello* 函数。

Simple Client:

简单客户端：

```js
conn, err := gRPC.Dial("127.0.0.1:8080", gRPC.WithInsecure())
 if err != nil {
     panic(err)
 }
 client := pb.NewHelloWorldHandlerClient(conn)
 client.SayHello(context.Background(), &pb.Hello{Name: "alex"})
```


Though I have removed some code for simplicity, these *services* and *messages* can become quite complex if needed. The most important thing to  understand is that when a server attempts to announce itself as a *HelloWorldHandlerServer*, it is required to implement the RPC functions as specified within the  protobuf schema. This agreement between the client and server makes  cross-language network calls feel like regular function calls.

尽管为了简单起见我删除了一些代码，但如果需要，这些 *services* 和 *messages* 可能会变得非常复杂。要理解的最重要的事情是，当服务器尝试将自己声明为 *HelloWorldHandlerServer* 时，需要实现 protobuf 模式中指定的 RPC 功能。客户端和服务器之间的这种协议使跨语言网络调用感觉就像常规函数调用一样。

In addition to the basic Unary server described above, gRPC lets you decide between four types of service methods:

除了上述基本的一元服务器之外，gRPC 还允许您在四种类型的服务方法之间进行选择：

- **Unary** (example above): client sends a single request to the server and gets a single response back, just like a normal function call.
- **Server Streaming:** server returns a stream of messages in response to a client's request.
- **Client Streaming:** client sends a stream of messages to the server and the server replies  in a single message, usually once the client has finished streaming.
- **Bi-directional Streaming:** the client and server can both send streams of messages to each other asynchronously.

- **一元**（上面的例子）：客户端向服务器发送一个请求并得到一个响应，就像一个普通的函数调用。
- **Server Streaming：** 服务器返回消息流以响应客户端的请求。
- **客户端流式传输：** 客户端向服务器发送消息流，服务器在单个消息中回复，通常是在客户端完成流式传输后。
- **双向流：** 客户端和服务器都可以异步地相互发送消息流。

### gRPC Performance

### gRPC 性能

Not all HTTP connections are created equal. Though Golang natively supports HTTP/2, the HTTP/2 transport must be set by the client and the server  must also support HTTP/2. Before moving to gRPC, we were still using  HTTP/1.1 for client connections. We could have switched to HTTP/2 for  performance gains, but we would have lost some of the benefits of native protobuf compression and usability changes.

并非所有的 HTTP 连接都是平等的。虽然 Golang 本身支持 HTTP/2，但 HTTP/2 传输必须由客户端设置，服务器也必须支持 HTTP/2。在转向 gRPC 之前，我们仍然使用 HTTP/1.1 进行客户端连接。我们本可以切换到 HTTP/2 以获得性能提升，但我们会失去原生 protobuf 压缩和可用性更改的一些好处。

The best option  available in HTTP/1.1 is pipelining. Pipelining means that although  requests can share a connection, they must queue up one after the other  until the request in front completes. HTTP/2 improved pipelining by  using connection multiplexing. Multiplexing allows for multiple requests to be sent on the same connection and at the same time. 
HTTP/1.1 中可用的最佳选项是流水线。流水线意味着虽然请求可以共享一个连接，但它们必须一个接一个地排队，直到前面的请求完成。 HTTP/2 通过使用连接多路复用改进了流水线。多路复用允许在同一连接上同时发送多个请求。

HTTP REST APIs generally use JSON for their request and response format. Protobuf is the native request/response format of gRPC because it has a standard schema agreed upon by the client and server during registration. In  addition, protobuf is known to be significantly faster than JSON due to  its serialization speeds. I’ve run some benchmarks on my laptop, source  code can be found [here](https://github.com/Fattouche/protobuf-benchmark).

HTTP REST API 通常使用 JSON 作为其请求和响应格式。 Protobuf 是 gRPC 的原生请求/响应格式，因为它具有客户端和服务器在注册期间商定的标准架构。此外，众所周知，protobuf 的序列化速度明显快于 JSON。我已经在我的笔记本电脑上运行了一些基准测试，源代码可以在这里找到（https://github.com/Fattouche/protobuf-benchmark）。

![Protobuf performs better in small, medium, and large data sizes.](https://blog.cloudflare.com/content/images/2021/03/image1-26.png)

As you can see, protobuf performs better in small, medium, and large data  sizes. It is faster per operation, smaller after marshalling, and scales well with input size. This becomes even more noticeable when  unmarshaling very large data sets. Protobuf takes 96.4ns/op but JSON  takes 22647ns/op, a 235X reduction in time! For large DNS zones, this  efficiency makes a massive difference in the time it takes us to go from record change in our API to serving it at the edge.

如您所见，protobuf 在小、中、大数据规模上表现更好。每个操作速度更快，编组后更小，并且可以很好地扩展输入大小。当解组非常大的数据集时，这变得更加明显。 Protobuf 需要 96.4ns/op，但 JSON 需要 22647ns/op，时间减少了 235 倍！对于大型 DNS 区域，这种效率对我们从 API 中的记录更改到在边缘提供服务所需的时间产生巨大影响。

Combining the benefits of HTTP/2 and protobuf showed almost no performance change  from our application’s point of view. This is likely due to the fact  that our pods were already so close together that our connection times  were already very low. In addition, most of our gRPC calls are done with small amounts of data where the difference is negligible. One thing  that we did notice **—** likely related to the multiplexing of HTTP/2 **—** was greater efficiency when writing newly created/edited/deleted records to the edge. Our latency spikes dropped in both amplitude and frequency.

从我们的应用程序的角度来看，结合 HTTP/2 和 protobuf 的优点几乎没有性能变化。这可能是因为我们的 pod 已经非常靠近，以至于我们的连接时间已经非常短。此外，我们的大部分 gRPC 调用都是使用少量数据完成的，其中差异可以忽略不计。我们确实注意到 **—** 可能与 HTTP/2 的多路复用有关的一件事 **—** 在将新创建/编辑/删除的记录写入边缘时效率更高。我们的延迟峰值在幅度和频率上都有所下降。

![Our latency spikes dropped in both amplitude and frequency.](https://blog.cloudflare.com/content/images/2021/03/image2-19.png)

### gRPC Security

### gRPC 安全性

One of the best features in Kubernetes is the NetworkPolicy. This allows developers to control what goes in and what goes out.

Kubernetes 中最好的特性之一是 NetworkPolicy。这允许开发人员控制进入和退出的内容。

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


In this example, taken from the [Kubernetes docs](https://kubernetes.io/docs/concepts/services-networking/network-policies/), we can see that this will create a network policy called  test-network-policy . This policy controls both ingress and egress  communication to or from any pod that matches the role *db* and enforces the following rules:

在这个例子中，取自 [Kubernetes docs](https://kubernetes.io/docs/concepts/services-networking/network-policies/)，我们可以看到这将创建一个名为 test-network-policy 的网络策略.此策略控制与角色 *db* 匹配的任何 pod 的进出通信，并强制执行以下规则：

Ingress connections allowed:

允许入口连接：

- Any pod in default namespace with label “role=frontend”
- Any pod in any namespace that has a label “project=myproject”
- Any source IP address in 172.17.0.0/16 except for 172.17.1.0/24

- 默认命名空间中带有“role=frontend”标签的任何 pod
- 任何命名空间中带有“project=myproject”标签的任何 pod
- 除 172.17.1.0/24 外的 172.17.0.0/16 中的任何源 IP 地址

Egress connections allowed:

允许出口连接：

- Any dest IP address in 10.0.0.0/24

- 10.0.0.0/24 中的任何目标 IP 地址

NetworkPolicies do a fantastic job of protecting APIs at the network level, however,  they do nothing to protect APIs at the application level. If you wanted  to control which endpoints can be accessed within the API, you would  need k8s to be able to not only distinguish between pods, but also  endpoints within those pods. These concerns led us to [per RPC credentials](https://grpc.io/docs/guides/auth/). Per RPC credentials are easy to set up on top of the pre-existing gRPC  code. All you need to do is add interceptors to both your stream and  unary handlers.

NetworkPolicies 在网络级别保护 API 方面做得非常出色，但是，它们在应用程序级别保护 API 没有任何作用。如果您想控制可以在 API 中访问哪些端点，则需要 k8s 不仅能够区分 Pod，还能够区分这些 Pod 中的端点。这些担忧使我们[每个 RPC 凭据](https://grpc.io/docs/guides/auth/)。每个 RPC 凭据很容易在预先存在的 gRPC 代码之上设置。您需要做的就是将拦截器添加到流和一元处理程序中。

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

在这个示例代码片段中，我们从 info 对象中获取用户名、密码和请求的函数。然后我们对客户端进行身份验证以确保它具有调用该函数的正确权限。该拦截器将在调用任何其他函数之前运行，这意味着一个实现可以保护所有函数。客户端将初始化其安全连接并发送凭据，如下所示：

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

这里客户端首先验证服务器是否与 certFile 匹配。此步骤可确保客户端不会意外地将其密码发送给不良行为者。接下来，客户端使用其用户名和密码初始化 *perRPCCreds* 结构，并使用该信息拨打服务器。每当客户端调用 rpc 定义的函数时，服务器都会验证其凭据。

### Next Steps

### 下一步

Our next step is to remove the need for many applications to access the  database and ultimately DRY up our codebase by pulling all DNS-related  code into a single API, accessed from one gRPC interface. This removes  the potential for mistakes in individual applications and makes updating our database schema easier. It also gives us more granular control over which functions can be accessed rather than which tables can be  accessed.

我们的下一步是消除许多应用程序访问数据库的需要，并最终通过将所有与 DNS 相关的代码提取到单个 API 中来干掉我们的代码库，从一个 gRPC 接口访问。这消除了单个应用程序中出错的可能性，并使更新我们的数据库架构更容易。它还使我们能够更精细地控制可以访问哪些函数而不是可以访问哪些表。

So far, the DNS team is very happy with the results of  our gRPC migration. However, we still have a long way to go before we  can move entirely away from REST. We are also patiently waiting for [HTTP/3 support](https://github.com/grpc/grpc/issues/19126) for gRPC so that we can take advantage of those super [quic](https://en. wikipedia.org/wiki/QUIC) speeds! 
到目前为止，DNS 团队对我们的 gRPC 迁移结果非常满意。然而，在我们完全摆脱 REST 之前，我们还有很长的路要走。我们也在耐心等待 gRPC 的 [HTTP/3 支持](https://github.com/grpc/grpc/issues/19126)，以便我们可以利用那些超级 [quic](https://en.wikipedia.org/wiki/QUIC)速度！

