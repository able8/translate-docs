# How to Build a Custom Kubernetes Ingress Controller in Go

# 如何在 Go 中构建自定义 Kubernetes Ingress Controller

Oct 10, 2019 at 11:52AM

2019 年 10 月 10 日上午 11:52

    Caleb Doxsey

迦勒·多克西

Recently I switched from GKE to Digital Ocean as my Managed  Kubernetes provider. As part of the switch I wanted to also start using a Kubernetes Ingress Controller to map incoming HTTP requests to specific services. After some frustration attempting to get the NGINX Ingress  Controller working, I ended up rolling my own in an afternoon. This blog post will explain how I did that.

最近，我从 GKE 切换到 Digital Ocean 作为我的托管 Kubernetes 提供商。作为切换的一部分，我还想开始使用 Kubernetes 入口控制器将传入的 HTTP 请求映射到特定服务。在尝试让 NGINX Ingress Controller 工作时遇到了一些挫折之后，我在一个下午结束了自己的工作。这篇博文将解释我是如何做到的。

Even if you have no need to make your own ingress controller, the  steps described below are generally useful both for Kubernetes  development as well as building HTTP proxies in Go.

即使您不需要制作自己的入口控制器，下面描述的步骤通常对 Kubernetes 开发以及在 Go 中构建 HTTP 代理都很有用。

### Overview

###  概述

A description of my basic Kubernetes setup can be found in a prior blog post: [Kubernetes: The Surprisingly Affordable Platform for Personal Projects](https://www.doxsey.net/blog/kubernetes--the-surprisingly-affordable-platform-for-personal-projects). I maintain several web applications in a Kubernetes cluster. Each of  those web applications is made up of a Deployment and a corresponding  Service:

可以在之前的博客文章中找到对我的基本 Kubernetes 设置的描述：[Kubernetes：用于个人项目的令人惊讶的负担得起的平台](https://www.doxsey.net/blog/kubernetes--the-surprisingly-affordable-platform-对于个人项目)。我在 Kubernetes 集群中维护了几个 Web 应用程序。这些 Web 应用程序中的每一个都由一个部署和一个相应的服务组成：

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

有了这个服务 + 部署，就可以通过从 Kubernetes 集群中转到 `doxsey-www.default.svc.cluster` 来访问 Web 应用程序。但是 Internet 上的客户端如何访问该服务呢？

Previously the way I solved this was by running NGINX as a DaemonSet on each node attached to the host network:

以前我解决这个问题的方法是在连接到主机网络的每个节点上运行 NGINX 作为 DaemonSet：

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

通过设置 `hostNetwork: true`，容器将绑定节点本身的端口 80，而不仅仅是容器，这样我们就可以通过访问公共 IP 地址的端口 80 来访问 NGINX。然后我在防火墙中打开端口 80 和 443，并使用[自定义应用程序](https://github.com/calebdoxsey/kubernetes-cloudflare-sync) 将节点公共 IP 地址同步为我的 DNS 提供商中的 A 记录。

Although this approach works I ran into a few problems:

虽然这种方法有效，但我遇到了一些问题：

1. It requires creating an NGINX configuration file, which can be a bit arcane and surprisingly difficult to get right. If I ever add new  deployments or domain names, that configuration file has to manually  modified to account for this.
2. TLS certs, though dynamically created with letsencrypt, are also manually configured in the NGINX configuration.
3. NGINX's default behavior is to crash if a DNS lookup on a backend  fails when first starting — which can easily happen if all the replicas  in a deployment are unavailable. Although this bizarre behavior can be  fixed in the enterprise version of NGINX, I never did find a good  solution for how to fix it in the open source version.
4. Because I relied on ephemeral nodes which disappeared every 24  hours, the DNS lookup problem sometimes resulted in no NGINX daemons  being available.

1. 它需要创建一个 NGINX 配置文件，这可能有点神秘，而且很难做到正确。如果我添加了新的部署或域名，则必须手动修改该配置文件以解决此问题。
2. TLS 证书虽然是用letsencrypt 动态创建的，但也是在NGINX 配置中手动配置的。
3. 如果首次启动时后端的 DNS 查找失败，NGINX 的默认行为是崩溃——如果部署中的所有副本都不可用，这很容易发生。虽然这种奇怪的行为可以在 NGINX 的企业版中修复，但我一直没有找到如何在开源版本中修复它的好的解决方案。
4. 因为我依赖于每 24 小时消失一次的临时节点，DNS 查找问题有时会导致没有可用的 NGINX 守护进程。

So this time around I decided to go a different route: using Ingress objects.

所以这一次我决定走不同的路线：使用 Ingress 对象。

### Ingress Objects

### 入口对象

Kubernetes has support for mapping external domains to internal  services via Ingress objects. The Ingress object I'm using looks like  this:

Kubernetes 支持通过 Ingress 对象将外部域映射到内部服务。我正在使用的 Ingress 对象如下所示：

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

此清单声明应如何将传入的 HTTP 流量路由到后端服务：

- Any HTTP request to the `www.doxsey.net` domain, for any path will be routed to the `doxsey-www` service.
- If this is an HTTPS request, and the domain matches `*.doxsey.net`, the `doxsey-net-tls` certificate will be used for the request. 

- 对 `www.doxsey.net` 域的任何 HTTP 请求，对于任何路径，都将路由到 `doxsey-www` 服务。
- 如果这是一个 HTTPS 请求，并且域与 `*.doxsey.net` 匹配，`doxsey-net-tls` 证书将用于请求。

This configuration is significantly simpler than the manual  configuration I had to do with NGINX and it also has native support for  TLS certificates as they are stored in Kubernetes (as Kubernetes  Secrets). Furthermore adjusting the manifest for a Service or Ingress  will automatically adjust the routing rules accordingly.

此配置比我必须使用 NGINX 进行的手动配置简单得多，并且它还对 TLS 证书提供本机支持，因为它们存储在 Kubernetes 中（作为 Kubernetes Secrets）。此外，调整服务或入口的清单将相应地自动调整路由规则。

### Ingress Controllers

### 入口控制器

In addition to the Ingress object, we have to install an Ingress  Controller. This Controller is responsible for reading the ingress  manifest rules and actually implementing the desired proxy behavior. In  other words, an Ingress manifest is merely a declarative intent.

除了 Ingress 对象，我们还必须安装一个 Ingress Controller。该控制器负责读取入口清单规则并实际实现所需的代理行为。换句话说，Ingress manifest 只是一个声明性的意图。

The default ingress controller is the [NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx), though there are [many others](https://caylent.com/kubernetes-top-ingress-controllers). Although it was a fun exercise to implement my own, **you should probably use one of these**. There can be a lot of subtlety in handling all the edge cases,  particularly when dealing with lots of services or rules, and those edge cases and bugs have probably been fixed in more widely used  controllers.

默认的入口控制器是 [NGINX 入口控制器](https://github.com/kubernetes/ingress-nginx)，尽管还有 [许多其他](https://caylent.com/kubernetes-top-ingress-controllers)。虽然实现我自己的练习很有趣，但**您可能应该使用其中一个**。处理所有边缘情况可能有很多微妙之处，尤其是在处理大量服务或规则时，这些边缘情况和错误可能已在更广泛使用的控制器中得到修复。

Nevertheless building a custom ingress controller is surprisingly straightforward.

然而，构建自定义入口控制器非常简单。

### A Custom Ingress Controller

### 自定义入口控制器

*The code for this application is available [here](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller).*

*此应用程序的代码可在 [此处](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller) 获得。*

To build a custom ingress controller in Go we need to create an application which will do the following:

要在 Go 中构建自定义入口控制器，我们需要创建一个应用程序，它将执行以下操作：

1. Query the Kubernetes API for Services, Ingresses and Secrets and listen for changes.
2. Load the TLS certificates so that they can be used to terminate HTTP requests.
3. Construct a routing table to be used by the HTTP server based on the loaded Kubernetes data. This routing table should be efficient since  all incoming HTTP traffic will go through it.
4. Listen on `:80` and `:443` for incoming HTTP requests. A backend will be looked up according to the routing table and an `httputil.ReverseProxy` will be used to proxy the request and response. For `:443` the appropriate TLS cert will be used for the secure connection.

1. 查询 Kubernetes API 的服务、入口和秘密并监听变化。
2. 加载 TLS 证书，以便它们可用于终止 HTTP 请求。
3、根据加载的Kubernetes数据构建HTTP服务器使用的路由表。这个路由表应该是有效的，因为所有传入的 HTTP 流量都将通过它。
4. 监听传入 HTTP 请求的 `:80` 和 `:443`。将根据路由表查找后端，并使用 httputil.ReverseProxy 代理请求和响应。对于 `:443`，适当的 TLS 证书将用于安全连接。

Let's take each of these in turn.

让我们依次考虑这些。

#### Querying Kubernetes

#### 查询 Kubernetes

A Kubernetes client can be created by getting a rest config and calling `NewForConfig`:

Kubernetes 客户端可以通过获取休息配置并调用 `NewForConfig` 来创建：

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

从那里我们创建了一个 `Watcher` 和 `Payload`。 `Watcher` 负责查询 Kubernetes 并创建 `Payload`。 `Payload` 包含完成 HTTP 请求所需的所有 Kubernetes 数据：

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

除了端口之外，入口还可以通过名称引用后端服务端口，因此我们通过检索相应的服务定义来填充该数据。

The `Watcher` has a single `Run(ctx context.Context) error` method and contains two fields:

`Watcher` 有一个 `Run(ctx context.Context) error` 方法并包含两个字段：

```
// A Watcher watches for ingresses in the kubernetes cluster
type Watcher struct {
    client   kubernetes.Interface
    onChange func(*Payload)
}
```

With this approach the `onChange` function will be called anytime we detect that something has changed. There are a couple other ways we could build this API:

使用这种方法，只要我们检测到某些内容发生了变化，就会调用 `onChange` 函数。我们可以通过其他几种方式构建此 API：

1. We could use a Channel for updates:

    1. 我们可以使用 Channel 进行更新：

   ```
    type Watcher struct {
       Updates chan *Payload
   }
   ```

2. We could use an iterator like `bufio.Scanner`

    2. 我们可以使用像 `bufio.Scanner` 这样的迭代器

   ```
    func (w *Watcher) OnUpdate() bool
   func (w *Watcher) Payload() *Payload
   func (w *Watcher) Err() error
   ```

   
Used like:

    像这样使用：

   ```
    for watcher.OnUpdate() {
       // ...
   }
   ```

Because of how we use the `Watcher` it doesn't make a big difference which approach is taken. Notice that the `Watcher` doesn't know anything about HTTP routing. Using the [separation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns) design principle helps to make the code easier to understand.

由于我们如何使用`Watcher`，采用哪种方法并没有太大区别。请注意，`Watcher` 对 HTTP 路由一无所知。使用[关注点分离](https://en.wikipedia.org/wiki/Separation_of_concerns) 设计原则有助于使代码更易于理解。

To implement `Run` we use the `k8s.io/client-go/informers`package. This package provides a type-safe, efficient mechanism for retrieving,  listing and watching Kubernetes objects. We create a `SharedInformerFactory` along with `Lister`s for each object we're interested in:

为了实现 `Run`，我们使用 `k8s.io/client-go/informers` 包。该包提供了一种类型安全、高效的检索、列出和监视 Kubernetes 对象的机制。我们为我们感兴趣的每个对象创建一个 `SharedInformerFactory` 和 `Lister`s：

```
func (w *Watcher) Run(ctx context.Context) error {
    factory := informers.NewSharedInformerFactory(w.client, time.Minute)
    secretLister := factory.Core().V1().Secrets().Lister()
    serviceLister := factory.Core().V1().Services().Lister()
    ingressLister := factory.Extensions().V1beta1().Ingresses().Lister()
```

We then define a local `onChange` function which will be  called for anytime a change is detected. Rather than utilize special  rules for each type of change, it's easier to just rebuild everything  from scratch every time. We should only look into more specialized logic if we discover a performance bottleneck. This is especially true  because our `Watcher` runs in a different goroutine from our  HTTP handler. We can essentially build the payload in the background  without affecting any ongoing requests, so if it takes a few seconds,  it's not a big deal.

然后我们定义了一个本地 `onChange` 函数，该函数将在检测到更改时调用。与其对每种类型的更改都使用特殊规则，不如每次从头开始重新构建所有内容更容易。如果我们发现性能瓶颈，我们应该只研究更专业的逻辑。尤其如此，因为我们的 `Watcher` 运行在与 HTTP 处理程序不同的 goroutine 中。我们基本上可以在后台构建有效负载，而不会影响任何正在进行的请求，所以如果需要几秒钟，这没什么大不了的。

We start by listing the ingresses:

我们首先列出入口：

```
ingresses, err := ingressLister.List(labels.Everything())
if err != nil {
    log.Error().Err(err).Msg("failed to list ingresses")
    return
}
```

Then for each ingress, if there's one or more TLS rules, we load those from the secrets:

然后对于每个入口，如果有一个或多个 TLS 规则，我们从秘密加载这些规则：

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

Go 对内置的密码学有很好的支持，这使得这段代码非常简单。对于实际的 HTTP 规则，我创建了一个 `addBackend` 辅助函数：

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

这会为每个 HTTP 规则以及可选的默认规则调用：

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

然后我们调用 `onChange` 回调：

```
w.onChange(payload)
```

The local `onChange` function is invoked any time something changes, so the final step is to start our informers:

本地 `onChange` 函数会在任何事情发生变化时被调用，所以最后一步是启动我们的 Informers：

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

每个通知者使用相同的处理程序：

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

[Debouncing](https://godoc.org/github.com/bep/debounce) is a way of avoiding duplicate events. We set a small delay, and if an  additional event occurs before we hit the delay, we restart the timer. Using a debouncer makes it likely that the first `onChange` event will include all of the ingresses, services and secrets, rather than receiving a partial view of the current state.

[Debounce](https://godoc.org/github.com/bep/debounce) 是一种避免重复事件的方法。我们设置了一个小的延迟，如果在我们达到延迟之前发生了额外的事件，我们会重新启动计时器。使用 debouncer 使得第一个 `onChange` 事件可能包括所有的入口、服务和秘密，而不是接收当前状态的部分视图。

And that's basically it for the watcher. You can see the source code [here](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller/blob/master/watcher/watcher.go).

对于观察者来说，基本上就是这样。你可以在[这里](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller/blob/master/watcher/watcher.go)查看源代码。

#### Routing Table

#### 路由表

Our goal with the routing table is to make it efficient to query by  pre-computing most of the lookup information. There's usually a  trade-off here between simple solutions which don't have the best  performance characterics and hyper-specialized data structures, which  often have subtle bugs and are difficult to understand and maintain.

我们对路由表的目标是通过预先计算大部分查找信息来提高查询效率。通常在没有最佳性能特征的简单解决方案和通常具有微妙错误且难以理解和维护的超专业数据结构之间进行权衡。

One of the easiest ways of implementing a query interface which is  both efficient and easy to understand is to use maps. Maps give us `O(1)` lookup and its quite difficult to do much better (though some sort of  trie might be worthwhile for path prefix / regexp matching). I used a  hybrid approach where an initial lookup is done with a map and if  multiple entries are found after than, a slice is used (which is `O(n)`, but `n` is typically 1)

实现既高效又易于理解的查询接口的最简单方法之一是使用地图。 Maps 为我们提供了`O(1)` 查找，并且很难做得更好（尽管对于路径前缀/正则表达式匹配，某种trie 可能是值得的）。我使用了一种混合方法，其中使用地图完成初始查找，如果在之后找到多个条目，则使用切片（这是“O(n)”，但“n”通常为 1）

A routing table consists of two maps:

路由表由两个映射组成：

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

这些对应于两种方法：

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
    if idx := strings.IndexByte(host, ':');idx > 0 {
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

`GetCertificate` 用于获取用于安全连接的 TLS 证书。 HTTP 处理程序使用`GetBackend` 将请求代理到后端。对于 TLS 证书，我们有一个 `matches` 方法来处理通配符证书：

```
func (rt *RoutingTable) matches(sni string, certHost string) bool {
    for strings.HasPrefix(certHost, "*.") {
        if idx := strings.IndexByte(sni, '.');idx >= 0 {
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

对于后端，matches 方法实际上是一个正则表达式（因为 Ingress 路径的定义是一个正则表达式）：

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

您可以在 [此处](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller/blob/master/server/route.go)查看这些地图的构建方式。

#### HTTP Server

#### HTTP 服务器

For the HTTP Server I decided to use an API configured via [functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis). There's a private `config` struct:

对于 HTTP 服务器，我决定使用通过 [功能选项](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) 配置的 API。有一个私有的 `config` 结构：

```
type config struct {
    host    string
    port    int
    tlsPort int
}
```

An `Option` type:

`Option` 类型：

```
// An Option modifies the config.
type Option func(*config)
```

And functions to set the options. Like `WithHost`:

以及设置选项的功能。像`WithHost`：

```
// WithHost sets the host to bind in the config.
func WithHost(host string) Option {
    return func(cfg *config) {
        cfg.host = host
    }
}
```

Our server struct and constructor look like this:

我们的服务器结构和构造函数如下所示：

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

通过使用 *sane defaults* 这种类型的 API 使大多数客户端的使用变得非常简单（大多数时候你只说“New()”），同时仍然提供在需要时更改选项的灵活性。这种 API 方法在 Go 中变得非常普遍。有关广泛使用的示例，请查看 [gRPC](https://godoc.org/google.golang.org/grpc#Dial)。

In addition to the config, our server also has a pointer to the routing table and a ready `Event` we use to signal the first time the payload is set. More on that in a minute, but first notice we use an `atomic.Value` to store the routing table. Why do that?

除了配置之外，我们的服务器还有一个指向路由表的指针和一个准备好的“事件”，我们用来在第一次设置有效负载时发出信号。稍后会详细介绍，但首先请注意我们使用 `atomic.Value` 来存储路由表。为什么要这样做？

Go programs are not thread-safe. If our routing table is modified  while the HTTP handler is attempting to read it, we may end up with  corrupt state or our program crashing. Because of this we need to  prevent the simultaneous reading and writing of a shared data structure. There are different ways to achieve this:

Go 程序不是线程安全的。如果我们的路由表在 HTTP 处理程序试图读取它时被修改，我们最终可能会出现损坏状态或我们的程序崩溃。因此，我们需要防止同时读取和写入共享数据结构。有不同的方法可以实现这一点：

1. The way I opted to go was to use an `atomic.Value`. This type provides a `Load` and `Store` method which allows you to atomically read/write the value. Since we  rebuild the routing table on every change we can safely swap the old and new routing tables in a single operation. This is quite similar to the `ReadMostly` example from the [documentation](https://godoc.org/sync/atomic#Value):

    1. 我选择的方式是使用 `atomic.Value`。这种类型提供了一个 `Load` 和 `Store` 方法，允许你原子地读/写值。由于我们在每次更改时都重建路由表，因此我们可以在一次操作中安全地交换旧路由表和新路由表。这与 [文档](https://godoc.org/sync/atomic#Value) 中的“ReadMostly”示例非常相似：

   > The following example shows how to maintain a scalable frequently  read, but infrequently updated data structure using copy-on-write idiom.

    > 以下示例展示了如何使用写时复制习语来维护可扩展的频繁读取但不频繁更新的数据结构。

   One downside to this approach is that the type of the value stored has to be asserted at runtime:

    这种方法的一个缺点是必须在运行时断言存储的值的类型：

   ```
    s.routingTable.Load().(*RoutingTable).GetBackend(r.Host, r.URL.Path)
   ```

2. We could use a `Mutex` or `RWMutex` instead to control access to the critical region:

    2. 我们可以使用 `Mutex` 或 `RWMutex` 代替来控制对关键区域的访问：

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

这种方法与 `atomic.Value` 非常相似，但 `RWMutex` 的伸缩性不如 `atomic.Value`。使用大量 goroutine/CPU 内核时，您可能会遇到 [线程争用](https://github.com/golang/go/issues/17973) 的问题。

3. We could make the routing table itself thread safe. Instead of a `map` we could use `sync.Map` and add methods to dynamically update the routing table instead of rebuilding it each time.

    3.我们可以使路由表本身线程安全。我们可以使用 `sync.Map` 并添加方法来动态更新路由表，而不是每次都重新构建它，而不是使用 `map`。

   In general I would avoid this approach. It makes the code harder to  understand and maintain, and often adds unnecessary overhead if you  don't actually end up having multiple goroutines accessing the data  structure. Instead do your synchronization at the next level up  (basically wherever you end up starting goroutines).

    一般来说，我会避免这种方法。它使代码更难理解和维护，并且如果您实际上最终没有让多个 goroutine 访问数据结构，通常会增加不必要的开销。而是在下一个级别进行同步（基本上在您最终启动 goroutines 的任何地方）。

   Global, shared maps are generally a code-smell in Go programs and,  for that matter, in any programming where you want to utilize a large  number of CPU cores.

全局共享映射通常是 Go 程序中的一种代码味道，就此而言，在您想要使用大量 CPU 内核的任何编程中。

The actual Server `ServeHTTP` method looks like this:

实际的服务器 `ServeHTTP` 方法如下所示：

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

[`httputil`](https://godoc.org/net/http/httputil) 包有一个反向代理实现，我们可以将其用于 HTTP 服务器。它接受一个 URL，将请求转发到该 URL，并将响应发送回客户端。

The server code can be found [here](https://github.com/calebdoxsey/kubernetes-simple-ingress-controller/blob/master/server/server.go)

服务器代码可以在这里找到（https://github.com/calebdoxsey/kubernetes-simple-ingress-controller/blob/master/server/server.go）

#### Main

####  主要的

Stitching all the components together, our `main` function looks like this:

将所有组件拼接在一起，我们的 main 函数如下所示：

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
    if err := eg.Wait();err != nil {
        log.Fatal().Err(err).Send()
    }
}
```

#### Kubernetes Configuration

#### Kubernetes 配置

With the server code in place we can set it up in Kubernetes as a DaemonController running on each node:

有了服务器代码，我们可以在 Kubernetes 中将其设置为在每个节点上运行的 DaemonController：

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

就是这样。前面提到的 [`kubernetes-cloudflare-sync`](https://github.com/calebdoxsey/kubernetes-cloudflare-sync) 将节点公共 IP 与 cloudflare 同步，因此此自定义 HTTP 服务器将接收最终到达端口的任何请求80 或 443 个节点本身，它们将被代理到在 Kubernetes 中运行的后端服务。

### Conclusion

###  结论

So that's how I built a Kubernetes Ingress Controller from scratch. I'm sure I missed something along the way... but I'm actually using this application to serve this blog, so it at least kind of works. 

这就是我从头开始构建 Kubernetes Ingress Controller 的方式。我确定我在此过程中遗漏了一些东西……但我实际上是在使用这个应用程序来为这个博客提供服务，所以它至少是有效的。

