# How to Build a Custom Kubernetes Ingress Controller in Go  

Oct 10, 2019 at 11:52AM

​    Caleb Doxsey  

Recently I switched from GKE to Digital Ocean as my Managed  Kubernetes provider. As part of the switch I wanted to also start using a Kubernetes Ingress Controller to map incoming HTTP requests to specific services. After some frustration attempting to get the NGINX Ingress  Controller working, I ended up rolling my own in an afternoon. This blog post will explain how I did that.

Even if you have no need to make your own ingress controller, the  steps described below are generally useful both for Kubernetes  development as well as building HTTP proxies in Go.

### Overview

A description of my basic Kubernetes setup can be found in a prior blog post: [Kubernetes: The Surprisingly Affordable Platform for Personal Projects](https://www.doxsey.net/blog/kubernetes--the-surprisingly-affordable-platform-for-personal-projects). I maintain several web applications in a Kubernetes cluster. Each of  those web applications is made up of a Deployment and a corresponding  Service:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: doxsey-www
  labels:
    app: doxsey-www
spec:
  replicas: 1
  selector:
    matchLabels:
      app: doxsey-www
  template:
    metadata:
      labels:
        app: doxsey-www
    spec:
      containers:
        - name: doxsey-www
          image: quay.io/calebdoxsey/www:v1.4.0
          ports:
            - containerPort: 9002
---
kind: Service
apiVersion: v1
metadata:
  namespace: default
  name: doxsey-www
spec:
  selector:
    app: doxsey-www
  ports:
    - protocol: TCP
      port: 9002
      targetPort: 9002
```

With this service + deployment in place the web application can be reached by going to `doxsey-www.default.svc.cluster` from within the Kubernetes cluster. But how do clients on the internet access the service?

Previously the way I solved this was by running NGINX as a DaemonSet on each node attached to the host network:

```
spec:
  hostNetwork: true
  containers:
    - image: nginx:1.15.3-alpine
      name: nginx
      ports:
        - name: http
          containerPort: 80
          hostPort: 80
```

By setting `hostNetwork: true` the container will bind  port 80 of the node itself, not just the container, and by so doing we  can reach NGINX by hitting port 80 of our public IP address. I then open ports 80 and 443 in the firewall and use a [custom application](https://github.com/calebdoxsey/kubernetes-cloudflare-sync) to synchronize the node public IP addresses as A records in my DNS provider.

Although this approach works I ran into a few problems:

1. It requires creating an NGINX configuration file, which can be a bit arcane and surprisingly difficult to get right. If I ever add new  deployments or domain names, that configuration file has to manually  modified to account for this.
2. TLS certs, though dynamically created with letsencrypt, are also manually configured in the NGINX configuration.
3. NGINX's default behavior is to crash if a DNS lookup on a backend  fails when first starting — which can easily happen if all the replicas  in a deployment are unavailable. Although this bizarre behavior can be  fixed in the enterprise version of NGINX, I never did find a good  solution for how to fix it in the open source version.
4. Because I relied on ephemeral nodes which disappeared every 24  hours, the DNS lookup problem sometimes resulted in no NGINX daemons  being available.

So this time around I decided to go a different route: using Ingress objects.

### Ingress Objects

Kubernetes has support for mapping external domains to internal  services via Ingress objects. The Ingress object I'm using looks like  this:

```
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: doxsey-www
spec:
  tls:
    - hosts:
        - "*.doxsey.net"
      secretName: doxsey-net-tls
  rules:
    - host: www.doxsey.net
      http:
        paths:
          - path: /
            backend:
              serviceName: doxsey-www
              servicePort: 9002
```

This manifest declares how incoming HTTP traffic should be routed to backend services:

- Any HTTP request to the `www.doxsey.net` domain, for any path will be routed to the `doxsey-www` service.
- If this is an HTTPS request, and the domain matches `*.doxsey.net`, the `doxsey-net-tls` certificate will be used for the request.

This configuration is significantly simpler than the manual  configuration I had to do with NGINX and it also has native support for  TLS certificates as they are stored in Kubernetes (as Kubernetes  Secrets). Furthermore adjusting the manifest for a Service or Ingress  will automatically adjust the routing rules accordingly.

### Ingress Controllers

In addition to the Ingress object, we have to install an Ingress  Controller. This Controller is responsible for reading the ingress  manifest rules and actually implementing the desired proxy behavior. In  other words, an Ingress manifest is merely a declarative intent.

The default ingress controller is the [NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx), though there are [many others](https://caylent.com/kubernetes-top-ingress-controllers). Although it was a fun exercise to implement my own, **you should probably use one of these**. There can be a lot of subtlety in handling all the edge cases,  particularly when dealing with lots of services or rules, and those edge cases and bugs have probably been fixed in more widely used  controllers.

Nevertheless building a custom ingress controller is surprisingly straightforward.

### A Custom Ingress Controller

*The code for this application is available [here](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller).* 

To build a custom ingress controller in Go we need to create an application which will do the following:

1. Query the Kubernetes API for Services, Ingresses and Secrets and listen for changes.
2. Load the TLS certificates so that they can be used to terminate HTTP requests.
3. Construct a routing table to be used by the HTTP server based on the loaded Kubernetes data. This routing table should be efficient since  all incoming HTTP traffic will go through it.
4. Listen on `:80` and `:443` for incoming HTTP requests. A backend will be looked up according to the routing table and an `httputil.ReverseProxy` will be used to proxy the request and response. For `:443` the appropriate TLS cert will be used for the secure connection.

Let's take each of these in turn.

#### Querying Kubernetes

A Kubernetes client can be created by getting a rest config and calling `NewForConfig`:

```
config, err := rest.InClusterConfig()
if err != nil {
    log.Fatal().Err(err).Msg("failed to get kubernetes configuration")
}
client, err := kubernetes.NewForConfig(config)
if err != nil {
    log.Fatal().Err(err).Msg("failed to create kubernetes client")
}
```

From there we create a `Watcher` and `Payload`. The `Watcher` is responsible for querying Kubernetes and creating `Payload`s. The `Payload` consists of all the Kubernetes data needed to fulfill HTTP requests:

```
// A Payload is a collection of Kubernetes data loaded by the watcher.
type Payload struct {
    Ingresses       []IngressPayload
    TLSCertificates map[string]*tls.Certificate
}

// An IngressPayload is an ingress + its service ports.
type IngressPayload struct {
    Ingress      *extensionsv1beta1.Ingress
    ServicePorts map[string]map[string]int
}
```

Ingresses can reference backend service ports by name in addition to  port, so we populate that data by retrieving the corresponding service  definition.

The `Watcher` has a single `Run(ctx context.Context) error` method and contains two fields:

```
// A Watcher watches for ingresses in the kubernetes cluster
type Watcher struct {
    client   kubernetes.Interface
    onChange func(*Payload)
}
```

With this approach the `onChange` function will be called anytime we detect that something has changed. There are a couple other ways we could build this API:

1. We could use a Channel for updates:

   ```
   type Watcher struct {
       Updates chan *Payload
   }
   ```

2. We could use an iterator like `bufio.Scanner`

   ```
   func (w *Watcher) OnUpdate() bool
   func (w *Watcher) Payload() *Payload
   func (w *Watcher) Err() error
   ```

   Used like:

   ```
   for watcher.OnUpdate() {
       // ...
   }
   ```

Because of how we use the `Watcher` it doesn't make a big difference which approach is taken. Notice that the `Watcher` doesn't know anything about HTTP routing. Using the [separation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns) design principle helps to make the code easier to understand.

To implement `Run` we use the `k8s.io/client-go/informers`package. This package provides a type-safe, efficient mechanism for retrieving,  listing and watching Kubernetes objects. We create a `SharedInformerFactory` along with `Lister`s for each object we're interested in:

```
func (w *Watcher) Run(ctx context.Context) error {
    factory := informers.NewSharedInformerFactory(w.client, time.Minute)
    secretLister := factory.Core().V1().Secrets().Lister()
    serviceLister := factory.Core().V1().Services().Lister()
    ingressLister := factory.Extensions().V1beta1().Ingresses().Lister()
```

We then define a local `onChange` function which will be  called for anytime a change is detected. Rather than utilize special  rules for each type of change, it's easier to just rebuild everything  from scratch every time. We should only look into more specialized logic if we discover a performance bottleneck. This is especially true  because our `Watcher` runs in a different goroutine from our  HTTP handler. We can essentially build the payload in the background  without affecting any ongoing requests, so if it takes a few seconds,  it's not a big deal.

We start by listing the ingresses:

```
ingresses, err := ingressLister.List(labels.Everything())
if err != nil {
    log.Error().Err(err).Msg("failed to list ingresses")
    return
}
```

Then for each ingress, if there's one or more TLS rules, we load those from the secrets:

```
for _, rec := range ingress.Spec.TLS {
    if rec.SecretName != "" {
        secret, err := secretLister.Secrets(ingress.Namespace).Get(rec.SecretName)
        if err != nil {
            log.Error().Err(err).Str("namespace", ingress.Namespace).Str("name", rec.SecretName).Msg("unknown secret")
            continue
        }
        cert, err := tls.X509KeyPair(secret.Data["tls.crt"], secret.Data["tls.key"])
        if err != nil {
            log.Error().Err(err).Str("namespace", ingress.Namespace).Str("name", rec.SecretName).Msg("invalid tls certificate")
            continue
        }
        payload.TLSCertificates[rec.SecretName] = &cert
    }
}
```

Go has excellent support for cryptography built-in, which makes this code very simple. For the actual HTTP rules I made an `addBackend` helper function:

```
addBackend := func(ingressPayload *IngressPayload, backend extensionsv1beta1.IngressBackend) {
    svc, err := serviceLister.Services(ingressPayload.Ingress.Namespace).Get(backend.ServiceName)
    if err != nil {
        log.Error().Err(err).Str("namespace", ingressPayload.Ingress.Namespace).Str("name", backend.ServiceName).Msg("unknown service")
    } else {
        m := make(map[string]int)
        for _, port := range svc.Spec.Ports {
            m[port.Name] = int(port.Port)
        }
        ingressPayload.ServicePorts[svc.Name] = m
    }
}
```

This gets called for each HTTP rule, as well as the optional default rule:

```
if ingress.Spec.Backend != nil {
    addBackend(&ingressPayload, *ingress.Spec.Backend)
}
for _, rule := range ingress.Spec.Rules {
    if rule.HTTP != nil {
        continue
    }
    for _, path := range rule.HTTP.Paths {
        addBackend(&ingressPayload, path.Backend)
    }
}
```

And then we call the `onChange` callback:

```
w.onChange(payload)
```

The local `onChange` function is invoked any time something changes, so the final step is to start our informers:

```
var wg sync.WaitGroup
wg.Add(1)
go func() {
    informer := factory.Core().V1().Secrets().Informer()
    informer.AddEventHandler(handler)
    informer.Run(ctx.Done())
    wg.Done()
}()

wg.Add(1)
go func() {
    informer := factory.Extensions().V1beta1().Ingresses().Informer()
    informer.AddEventHandler(handler)
    informer.Run(ctx.Done())
    wg.Done()
}()

wg.Add(1)
go func() {
    informer := factory.Core().V1().Services().Informer()
    informer.AddEventHandler(handler)
    informer.Run(ctx.Done())
    wg.Done()
}()

wg.Wait()
```

The same handler is used for each informer:

```
debounced := debounce.New(time.Second)
handler := cache.ResourceEventHandlerFuncs{
    AddFunc: func(obj interface{}) {
        debounced(onChange)
    },
    UpdateFunc: func(oldObj, newObj interface{}) {
        debounced(onChange)
    },
    DeleteFunc: func(obj interface{}) {
        debounced(onChange)
    },
}
```

[Debouncing](https://godoc.org/github.com/bep/debounce) is a way of avoiding duplicate events. We set a small delay, and if an  additional event occurs before we hit the delay, we restart the timer.  Using a debouncer makes it likely that the first `onChange` event will include all of the ingresses, services and secrets, rather than receiving a partial view of the current state.

And that's basically it for the watcher. You can see the source code [here](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller/blob/master/watcher/watcher.go).

#### Routing Table

Our goal with the routing table is to make it efficient to query by  pre-computing most of the lookup information. There's usually a  trade-off here between simple solutions which don't have the best  performance characterics and hyper-specialized data structures, which  often have subtle bugs and are difficult to understand and maintain.

One of the easiest ways of implementing a query interface which is  both efficient and easy to understand is to use maps. Maps give us `O(1)` lookup and its quite difficult to do much better (though some sort of  trie might be worthwhile for path prefix / regexp matching). I used a  hybrid approach where an initial lookup is done with a map and if  multiple entries are found after than, a slice is used (which is `O(n)`, but `n` is typically 1)

A routing table consists of two maps:

```
type RoutingTable struct {
    certificatesByHost map[string]map[string]*tls.Certificate
    backendsByHost     map[string][]routingTableBackend
}

// NewRoutingTable creates a new RoutingTable.
func NewRoutingTable(payload *watcher.Payload) *RoutingTable {
    rt := &RoutingTable{
        certificatesByHost: make(map[string]map[string]*tls.Certificate),
        backendsByHost:     make(map[string][]routingTableBackend),
    }
    rt.init(payload)
    return rt
}
```

These correspond to two methods:

```
// GetCertificate gets a certificate.
func (rt *RoutingTable) GetCertificate(sni string) (*tls.Certificate, error) {
    hostCerts, ok := rt.certificatesByHost[sni]
    if ok {
        for h, cert := range hostCerts {
            if rt.matches(sni, h) {
                return cert, nil
            }
        }
    }
    return nil, errors.New("certificate not found")
}

// GetBackend gets the backend for the given host and path.
func (rt *RoutingTable) GetBackend(host, path string) (*url.URL, error) {
    // strip the port
    if idx := strings.IndexByte(host, ':'); idx > 0 {
        host = host[:idx]
    }
    backends := rt.backendsByHost[host]
    for _, backend := range backends {
        if backend.matches(path) {
            return backend.url, nil
        }
    }
    return nil, errors.New("backend not found")
}
```

`GetCertificate` is used to get the TLS certificate used for secure connections. `GetBackend` is used by the HTTP handler to proxy the request to the backend. For the TLS certificate we have a `matches` method to handle wildcard certs:

```
func (rt *RoutingTable) matches(sni string, certHost string) bool {
    for strings.HasPrefix(certHost, "*.") {
        if idx := strings.IndexByte(sni, '.'); idx >= 0 {
            sni = sni[idx+1:]
        } else {
            return false
        }
        certHost = certHost[2:]
    }
    return sni == certHost
}
```

For the backend the `matches` method is actually a regular expression (because the definition of an Ingress path is a regular expression):

```
type routingTableBackend struct {
    pathRE *regexp.Regexp
    url    *url.URL
}

func (rtb routingTableBackend) matches(path string) bool {
    if rtb.pathRE == nil {
        return true
    }
    return rtb.pathRE.MatchString(path)
}
```

You can see how these maps are constructed [here](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller/blob/master/server/route.go).

#### HTTP Server

For the HTTP Server I decided to use an API configured via [functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis). There's a private `config` struct:

```
type config struct {
    host    string
    port    int
    tlsPort int
}
```

An `Option` type:

```
// An Option modifies the config.
type Option func(*config)
```

And functions to set the options. Like `WithHost`:

```
// WithHost sets the host to bind in the config.
func WithHost(host string) Option {
    return func(cfg *config) {
        cfg.host = host
    }
}
```

Our server struct and constructor look like this:

```
// A Server serves HTTP pages.
type Server struct {
    cfg          *config
    routingTable atomic.Value
    ready        *Event
}

// New creates a new Server.
func New(options ...Option) *Server {
    cfg := defaultConfig()
    for _, o := range options {
        o(cfg)
    }
    s := &Server{
        cfg:   cfg,
        ready: NewEvent(),
    }
    s.routingTable.Store(NewRoutingTable(nil))
    return s
}
```

By using *sane defaults* this type of API makes most client usage very easy (most of the time you just say `New()`) while still providing flexibility to change options when needed. This  approach to APIs has become quite common in Go. For a widely used  example check out [gRPC](https://godoc.org/google.golang.org/grpc#Dial).

In addition to the config, our server also has a pointer to the routing table and a ready `Event` we use to signal the first time the payload is set. More on that in a minute, but first notice we use an `atomic.Value` to store the routing table. Why do that?

Go programs are not thread-safe. If our routing table is modified  while the HTTP handler is attempting to read it, we may end up with  corrupt state or our program crashing. Because of this we need to  prevent the simultaneous reading and writing of a shared data structure. There are different ways to achieve this:

1. The way I opted to go was to use an `atomic.Value`. This type provides a `Load` and `Store` method which allows you to atomically read/write the value. Since we  rebuild the routing table on every change we can safely swap the old and new routing tables in a single operation. This is quite similar to the `ReadMostly` example from the [documentation](https://godoc.org/sync/atomic#Value):

   > The following example shows how to maintain a scalable frequently  read, but infrequently updated data structure using copy-on-write idiom.

   One downside to this approach is that the type of the value stored has to be asserted at runtime:

   ```
   s.routingTable.Load().(*RoutingTable).GetBackend(r.Host, r.URL.Path)
   ```

2. We could use a `Mutex` or `RWMutex` instead to control access to the critical region:

   ```
   // read
   s.mu.RLock()
   backendURL, err := s.routingTable.GetBackend(r.Host, r.URL.Path)
   s.mu.RUnlock()
   
   // write
   rt := NewRoutingTable(payload)
   s.mu.Lock()
   s.routingTable = rt
   s.mu.Unlock()
   ```

   This approach is very similar to the `atomic.Value`, but `RWMutex`s don't scale as well as the `atomic.Value`. With a large number of goroutines / CPU cores you may have issues with [thread contention](https://github.com/golang/go/issues/17973).

3. We could make the routing table itself thread safe. Instead of a `map` we could use `sync.Map` and add methods to dynamically update the routing table instead of rebuilding it each time.

   In general I would avoid this approach. It makes the code harder to  understand and maintain, and often adds unnecessary overhead if you  don't actually end up having multiple goroutines accessing the data  structure. Instead do your synchronization at the next level up  (basically wherever you end up starting goroutines).

   Global, shared maps are generally a code-smell in Go programs and,  for that matter, in any programming where you want to utilize a large  number of CPU cores.

The actual Server `ServeHTTP` method looks like this:

```
// ServeHTTP serves an HTTP request.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    backendURL, err := s.routingTable.Load().(*RoutingTable).GetBackend(r.Host, r.URL.Path)
    if err != nil {
        http.Error(w, "upstream server not found", http.StatusNotFound)
        return
    }
    log.Info().Str("host", r.Host).Str("path", r.URL.Path).Str("backend", backendURL.String()).Msg("proxying request")
    p := httputil.NewSingleHostReverseProxy(backendURL)
    p.ErrorLog = stdlog.New(log.Logger, "", 0)
    p.ServeHTTP(w, r)
}
```

The [`httputil`](https://godoc.org/net/http/httputil) package has a Reverse Proxy implementation that we can leverage for  HTTP server. It takes a URL, forwards the request to that URL, and sends the response back to the client.

The server code can be found [here](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller/blob/master/server/server.go)

#### Main

Stitching all the components together, our `main` function looks like this:

```
func main() {
    flag.StringVar(&host, "host", "0.0.0.0", "the host to bind")
    flag.IntVar(&port, "port", 80, "the insecure http port")
    flag.IntVar(&tlsPort, "tls-port", 443, "the secure https port")
    flag.Parse()

    client, err := kubernetes.NewForConfig(getKubernetesConfig())
    if err != nil {
        log.Fatal().Err(err).Msg("failed to create kubernetes client")
    }

    s := server.New(server.WithHost(host), server.WithPort(port), server.WithTLSPort(tlsPort))
    w := watcher.New(client, func(payload *watcher.Payload) {
        s.Update(payload)
    })

    var eg errgroup.Group
    eg.Go(func() error {
        return s.Run(context.TODO())
    })
    eg.Go(func() error {
        return w.Run(context.TODO())
    })
    if err := eg.Wait(); err != nil {
        log.Fatal().Err(err).Send()
    }
}
```

#### Kubernetes Configuration

With the server code in place we can set it up in Kubernetes as a DaemonController running on each node:

```
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: kubernetes-simple-ingress-controller
  namespace: default
  labels:
    app: ingress-controller
spec:
  selector:
    matchLabels:
      app: ingress-controller
  template:
    metadata:
      labels:
        app: ingress-controller
    spec:
      hostNetwork: true
      dnsPolicy: ClusterFirstWithHostNet
      serviceAccountName: kubernetes-simple-ingress-controller
      containers:
        - name: kubernetes-simple-ingress-controller
          image: quay.io/calebdoxsey/kubernetes-simple-ingress-controller:v0.1.0
          ports:
            - name: http
              containerPort: 80
            - name: https
              containerPort: 443
```

And that's it. The aforementioned [`kubernetes-cloudflare-sync`](https://github.com/calebdoxsey/kubernetes-cloudflare-sync) syncs the node public IPs with cloudflare, so this custom HTTP server  will receive any requests that end up on port 80 or 443 of the node  itself and they will be proxied to the backend services running in  Kubernetes.

### Conclusion

So that's how I built a Kubernetes Ingress Controller from scratch.  I'm sure I missed something along the way... but I'm actually using this application to serve this blog, so it at least kind of works.

