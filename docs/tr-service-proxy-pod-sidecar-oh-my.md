# Sidecar Proxy Pattern - The Basis Of Service Mesh

# Sidecar 代理模式——服务网格的基础

## Service proxy, pod, sidecar, oh my!

## 服务代理、pod、sidecar，哦，天哪！

September 6, 2020 (Updated: August 7, 2021)

## How services talk to each other?

## 服务如何相互通信？

Imagine you're developing a service... For certainty, let's call it _A_. It's going to provide some public HTTP API to its clients. However, to serve requests it needs to call another service. Let's call this upstream service - _B_.

想象一下，您正在开发一项服务……可以肯定的是，我们将其称为 _A_。它将为其客户端提供一些公共 HTTP API。然而，为了服务请求，它需要调用另一个服务。我们称这个上游服务为 - _B_。

![Service A talks to Service B directly.](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/10-service-a-service-b.png)

Obviously, neither network nor service _B_ is ideal. If service _A_ wants to decrease the impact of the failing upstream requests on its public API success rate, it has to do something about errors. For instance, it could start retrying failed requests.

显然，网络和服务 _B_ 都不是理想的。如果服务 _A_ 想要减少失败的上游请求对其公共 API 成功率的影响，它必须对错误做一些事情。例如，它可以开始重试失败的请求。

![Service A retries failed requests Service B.](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/20-service-a-service-b-with-retries.png)

Implementation of the retry mechanism requires some code changes in the service _A_, but the codebase is fresh, there are tons of advanced HTTP libraries, so you just need to grab one... Easy-peasy, right?

重试机制的实现需要在服务 _A_ 中进行一些代码更改，但是代码库是新鲜的，有大量高级 HTTP 库，所以你只需要抓住一个......轻松，对吧？

Unfortunately, this simplicity is not always the case. Replace service _A_ with service _Z_ that was written 10 years ago in some esoteric language by a developer that already retired. Or add to the equitation services _Q_, _U_, and _X_ written by different teams in three different languages. As a result, the cumulative cost of the company-wide retry mechanism implementation in the code gets really high...

不幸的是，这种简单性并非总是如此。将服务 _A_ 替换为服务 _Z_，该服务是 10 年前由已退休的开发人员用某种深奥的语言编写的。或者添加由不同团队以三种不同语言编写的马术服务 _Q_、_U_ 和 _X_。结果，代码中全公司重试机制实现的累积成本变得非常高......

![Service Mesh example](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/30-service-qux-service-b.png)

But what if retries are not the only thing you need? Proper _request timeouts_ have to be ensured as well. And how about _distributed tracing_? It'd be nice to correlate the whole request tree with the original customer transaction by propagating some additional HTTP headers. However, every such capability would make the HTTP libraries even more bloated...

但是，如果重试不是您唯一需要的呢？还必须确保适当的_请求超时_。那么_分布式跟踪_呢？通过传播一些额外的 HTTP 标头，将整个请求树与原始客户事务相关联会很好。然而，每一个这样的功能都会使 HTTP 库更加臃肿......

## What is a sidecar proxy?

## 什么是边车代理？

Let's try to go one level higher... or lower? 🤔

让我们尝试提高一级……或更低？ 🤔

In our original setup, service _A_ has been communicating with service _B_ directly. But what if we put an intermediary infrastructure component in between those services? Thanks to containerization, orchestration, devops, add a buzz word of your choice here, nowadays, it became so simple to configure infrastructure, that the cost of adding another infra component is often lower than the cost of writing application code...

在我们最初的设置中，服务 _A_ 一直与服务 _B_ 直接通信。但是如果我们在这些服务之间放置一个中间基础设施组件呢？感谢容器化、编排、devops，在这里添加您选择的流行语，如今，配置基础设施变得如此简单，添加另一个基础设施组件的成本通常低于编写应用程序代码的成本……

![Sidecar Proxy Pattern visualized](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/40-service-a-sidecar-service-b.png)

For the sake of simplicity, let's call the box enclosing the service _A_ and the secret intermediary component _a server_ (bare metal or virtual, doesn't really matter). And now it's about time to introduce one of the fancy words from the article's title. Any piece of software running on the server _alongside_ the primary service and helping it do its job is called _a sidecar_. I hope, the idea behind the name is more or less straightforward here.

为简单起见，我们将包含服务 _A_ 和秘密中介组件 _a server_ 的盒子称为（裸机或虚拟，并不重要）。现在是时候介绍文章标题中的一个花哨词了。任何在服务器上_与主要服务一起运行并帮助其完成工作的软件都称为_a sidecar_。我希望，这个名字背后的想法在这里或多或少是直截了当的。

But getting back to the service-to-service communication problem, what sidecar should we use to keep the service code free of the low-level details such as retries or request tracing? Well, the needed piece of software is called a _service proxy_. Probably, the most widely used implementation of the service proxy in the real world is [envoy](https://www.envoyproxy.io/). 

但是回到服务到服务的通信问题，我们应该使用什么 sidecar 来保持服务代码不受重试或请求跟踪等低级细节的影响？好吧，所需的软件称为_服务代理_。可能，现实世界中使用最广泛的服务代理实现是[envoy](https://www.envoyproxy.io/)。

The idea of the service proxy is the following: instead of accessing the service _B_ directly, code in the service _A_ now will be sending requests to the service proxy sidecar. Since both of the processes run on the same server, the loopback network interface (i.e. `127.0.0.1` _aka_ `localhost`) is perfectly suitable for this part of the communication. On every received HTTP request, the service proxy sidecar will make a request to the upstream service using the external network interface of the server. The response from the upstream will be eventually forwarded back by the sidecar to the service _A_.

服务代理的想法如下：服务_A_中的代码现在将向服务代理sidecar发送请求，而不是直接访问服务_B_。由于两个进程运行在同一台服务器上，环回网络接口（即`127.0.0.1`_aka_`localhost`）非常适合这部分通信。在每个收到的 HTTP 请求上，服务代理 sidecar 将使用服务器的外部网络接口向上游服务发出请求。来自上游的响应最终会被 sidecar 转发回服务 _A_。

I think, at this time, it's already obvious where the retry, timeouts, tracing, etc. logic should reside. Having this kind of functionality provided by a separate sidecar process makes enhancing any service written in any language with such capabilities rather trivial.

我认为，此时，重试、超时、跟踪等逻辑应该驻留在何处已经很明显了。拥有由单独的 sidecar 进程提供的这种功能使得增强以任何语言编写的具有此类功能的任何服务变得相当简单。

Interestingly enough, that service proxy could be used not only for outgoing traffic (egress) but also for the incoming traffic (ingress) of the service _A_. Usually, there is plenty of cross-cutting things that can be tackled on the ingress stage. For instance, proxy sidecars can do _SSL_ termination, request authentication, and more. A detailed diagram of a single server setup could look something like that:

有趣的是，该服务代理不仅可以用于传出流量（出口），还可以用于服务 _A_ 的传入流量（入口）。通常，有很多跨领域的事情可以在入口阶段解决。例如，代理 sidecar 可以执行 _SSL_ 终止、请求身份验证等。单个服务器设置的详细图表可能如下所示：

![Local service proxy intercepting ingress and egress traffic](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/50-single-host-sidecar.png)

Probably, the last fancy term we are going to cover here is _a pod_. People have been deploying code using virtual machines or bare metal servers for a long time... A server itself is already a good abstraction and a unit of encapsulation. For instance, every server has at least one external network interface, a network loopback interface for the internal [IPC](https://en.wikipedia.org/wiki/Inter-process_communication) needs, and it can run a bunch of processes sharing access to these communication means. Servers are usually addressable within the private network of the company by their IPs. Last but not least, it's pretty common to use a whole server for a single purpose (otherwise, maintenance quickly becomes a nightmare). I.e. you may have a group of identical servers running instances of service _A_, another group of servers each running an instance of service _B_, etc. So, why on earth would anybody want something better than a server? 

可能，我们将在这里介绍的最后一个花哨术语是 _a pod_。长期以来，人们一直在使用虚拟机或裸机服务器来部署代码……服务器本身已经是一个很好的抽象和封装单元。例如，每台服务器至少有一个外部网络接口，一个用于内部 [IPC](https://en.wikipedia.org/wiki/Inter-process_communication) 需求的网络环回接口，它可以运行一堆进程共享对这些通信手段的访问。服务器通常可通过其 IP 在公司的专用网络内进行寻址。最后但并非最不重要的是，将整个服务器用于单一目的是很常见的（否则，维护很快就会变成一场噩梦)。 IE。您可能有一组相同的服务器运行服务 _A_ 的实例，另一组服务器每个运行一个服务 _B_ 的实例，等等。那么，为什么有人想要比服务器更好的东西呢？

Despite being a good abstraction, the orchestration overhead servers introduce is often too high. So people started thinking about how to package applications more efficiently and that's how we got containers. Well, probably you know that _Docker_ and _container_ had been kind of a synonym for a long time and folks from Docker have been actively advocating for _"a one process per container"_ model. Obviously, this model is pretty different from the widely used _server_ abstraction where multiple processes are allowed to work side by side. And that's how we got the concept of _pods_. A pod is just a group of containers sharing a bunch of namespaces. If we now run a single process per container all of the processes in the pod will still share the common execution environment. In particular, the network namespace. Thus, all the containers in the pod will have a shared loopback interface and a shared external interface with an IP address assigned to it. Then it's up to the orchestration layer (say hi to Kubernetes) how to make all the pods reachable within the network by their IPs. And that's how people reinvented servers...

尽管是一个很好的抽象，但服务器引入的编排开销通常太高。所以人们开始考虑如何更有效地打包应用程序，这就是我们获得容器的方式。好吧，您可能知道 _Docker_ 和 _container_ 长期以来一直是同义词，Docker 的人们一直在积极倡导_“每个容器一个进程”_ 模型。显然，这个模型与广泛使用的 _server_ 抽象非常不同，后者允许多个进程并排工作。这就是我们如何得到 _pods_ 的概念。 Pod 只是一组共享一堆命名空间的容器。如果我们现在为每个容器运行一个进程，那么 pod 中的所有进程仍将共享公共执行环境。特别是网络命名空间。因此，Pod 中的所有容器都将拥有一个共享的环回接口和一个共享的外部接口，并为其分配了 IP 地址。然后由编排层（跟 Kubernetes 打个招呼）如何让所有 pod 在网络内通过它们的 IP 访问。这就是人们如何改造服务器......

So, getting back to all those blue boxes enclosing the service process and the sidecar on the diagrams above - we can think of them as being either a virtual machine, a bare metal server, or a pod. All three of them are more or less interchangeable abstractions.

因此，回到上图中包含服务流程和边车的所有蓝色框 - 我们可以将它们视为虚拟机、裸机服务器或 Pod。这三个或多或少都是可互换的抽象。

To summarize, let's try to visualize how the service to service communication could look like with the proxy sidecars:

总而言之，让我们尝试可视化使用代理边车的服务到服务通信的样子：

![Mesh of services talking to each other through sidecar proxies](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/60-service-to-service-topology.png)

_Example of service to service communication topology, a.k.a. service mesh._

_服务到服务通信拓扑的示例，也就是服务网格。_

## Sidecar proxy example (practical part)

## Sidecar 代理示例（实战部分）

Since the only way to really understand something is to write a blog post about it implement it yourself, let's quickly hack a [demo environment](https://github.com/iximiuz/envoy-playground).

由于真正理解某事的唯一方法是写一篇关于它自己实现它的博客文章，让我们快速破解一个[演示环境](https://github.com/iximiuz/envoy-playground)。

#### Service A talks to service B directly

#### 服务 A 直接与服务 B 对话

We will start from the simple setup where service _A_ will be accessing service _B_ directly:

我们将从简单的设置开始，其中服务 _A_ 将直接访问服务 _B_：

![Multi-service demo setup](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/70-demo-direct.png)

The code of the [service _A_](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-a) is relatively straightforward. It's just a simple HTTP server that makes a call to its upstream service _B_ on every client request. Depending on the response from the upstream, _A_ returns either an HTTP 200 or HTTP 500 to the client.

[service_A_](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-a)的代码比较简单。它只是一个简单的 HTTP 服务器，它在每个客户端请求上调用其上游服务 _B_。根据上游的响应，_A_ 向客户端返回 HTTP 200 或 HTTP 500。

```go
package main

// ...

var requestCounter = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "service_a_requests_total",
        Help: "The total number of requests received by Service A.",
    },
    []string{"status"},
)

func handler(w http.ResponseWriter, r *http.Request) {
    resp, err := httpGet(os.Getenv("UPSTREAM_SERVICE"))
    if err == nil {
        fmt.Fprintln(w, "Service A: upstream responded with:", resp)
        requestCounter.WithLabelValues("2xx").Inc()
    } else {
        http.Error(w, fmt.Sprintf("Service A: upstream failed with: %v", err.Error()),
            http.StatusInternalServerError)
        requestCounter.WithLabelValues("5xx").Inc()
    }
}

func main() {
    // init prometheus /metrics endpoint

    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(
        os.Getenv("SERVICE_HOST")+":"+os.Getenv("SERVICE_PORT"), nil))
}

```


_(see full version on [GitHub](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-a/main.go))_

_（请参阅 [GitHub] 上的完整版本（https://github.com/iximiuz/envoy-playground/tree/master/basics/service-a/main.go））_

Notice that instead of hard-coding, we use `SERVICE_HOST` and `SERVICE_PORT` env variables to specify the host and port of the HTTP API endpoint. It'll come in handy soon. Additionally, the code of the service relies on the `UPSTREAM_SERVICE` env variable when accessing the upstream service _B_.

请注意，我们使用 `SERVICE_HOST` 和 `SERVICE_PORT` 环境变量来指定 HTTP API 端点的主机和端口，而不是硬编码。很快就会派上用场。此外，服务的代码在访问上游服务 _B_ 时依赖于 `UPSTREAM_SERVICE` 环境变量。

To get some visibility, the code is instrumented with the primitive counter metric `service_a_requests_total` that gets incremented on every incoming request. We will use an instance of [prometheus](https://github.com/iximiuz/envoy-playground/tree/master/basics/prometheus) service to scrape the metrics exposed by the service _A_. 

为了获得一些可见性，代码使用原始计数器指标“service_a_requests_total”进行检测，该指标在每个传入请求时递增。我们将使用 [prometheus](https://github.com/iximiuz/envoy-playground/tree/master/basics/prometheus) 服务的实例来抓取服务 _A_ 公开的指标。

The implementation of the upstream [service _B_](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-b) is trivial as well. It's yet another HTTP server. However its behavior is rather close to a static endpoint.

上游 [service _B_](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-b) 的实现也很简单。它是另一个 HTTP 服务器。然而，它的行为与静态端点相当接近。

```go
package main

// ...

var ERROR_RATE int

var (
    requestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "service_b_requests_total",
            Help: "The total number of requests received by Service B.",
        },
        []string{"status"},
    )
)

func handler(w http.ResponseWriter, r *http.Request) {
    if rand.Intn(100) >= ERROR_RATE {
        fmt.Fprintln(w, "Service B: Yay! nounce", rand.Uint32())
        requestCounter.WithLabelValues("2xx").Inc()
    } else {
        http.Error(w, fmt.Sprintf("Service B: Ooops... nounce %v", rand.Uint32()),
            http.StatusInternalServerError)
        requestCounter.WithLabelValues("5xx").Inc()
    }
}

func main() {
    // set ERROR_RATE
    // init prometheus /metrics endpoint

    http.HandleFunc("/", handler)

    // Listen on all interfaces
    log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVICE_PORT"), nil))
}

```


_(see full version on [GitHub](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-b/main.go))_

_（请参阅 [GitHub] 上的完整版本（https://github.com/iximiuz/envoy-playground/tree/master/basics/service-b/main.go））_

Probably the only interesting part here is `ERROR_RATE`. The service is designed to fail requests with some constant rate, i.e. if `ERROR_RATE` is _20_, approximately 20% of requests will fail with HTTP 500 status code. As with the service _A_, we will use prometheus to scrape basic usage statistics, see the counter `service_b_requests_total`.

可能这里唯一有趣的部分是“ERROR_RATE”。该服务旨在以某种恒定速率使请求失败，即如果“ERROR_RATE”为 _20_，则大约 20% 的请求将失败并显示 HTTP 500 状态代码。与服务 _A_ 一样，我们将使用 prometheus 来抓取基本使用统计信息，请参阅计数器“service_b_requests_total”。

Now it's time to launch the services and wire them up together. We are going to use [podman](https://github.com/containers/podman) to build and run services. Mostly because unlike Docker, podman [supports the concept of pods out of the box](https://developers.redhat.com/blog/2019/01/15/podman-managing-containers-pods/). Heck, look at its name, it's **POD** man 🐵

现在是启动服务并将它们连接在一起的时候了。我们将使用 [podman](https://github.com/containers/podman) 来构建和运行服务。主要是因为与 Docker 不同，podman [支持开箱即用的 pod 概念](https://developers.redhat.com/blog/2019/01/15/podman-managing-containers-pods/)。哎呀，看它的名字，它是**POD**人🐵

We will start from creating the service _B_ since it's a dependency of the service _A_. Clone the [demo repository](https://github.com/iximiuz/envoy-playground) and run the following commands from its root (a Linux host with installed podman is assumed):

我们将从创建服务 _B_ 开始，因为它是服务 _A_ 的依赖项。克隆 [demo 存储库](https://github.com/iximiuz/envoy-playground) 并从其根目录运行以下命令（假设是安装了 podman 的 Linux 主机)：

_Click here to see service B Dockerfile._

_点击这里查看服务B Dockerfile。_

```dockerfile
FROM golang:1.15

# Build
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go mod download
RUN go build -o service-b

# Run
ENV ERROR_RATE=30
ENV SERVICE_PORT=80
ENV METRICS_PORT=8081
CMD ["/app/service-b"]
```


[source on GitHub](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-b/Dockerfile)

[GitHub 上的来源](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-b/Dockerfile)

```bash
# Build service B image
$ sudo podman build -t service-b -f service-b/Dockerfile

# Create a pod (read "server") for service B
$ sudo podman pod create --name service-b-pod

# Start service B container in the pod
$ sudo podman run --name service-b -d --rm --pod service-b-pod \
    -e ERROR_RATE=20 service-b

# Keep pod's IP address for future use
$ POD_B_IP=$(sudo podman inspect -f "{{.NetworkSettings.IPAddress}}" \
    $(sudo podman pod inspect -f "{{.InfraContainerID}}" service-b-pod))

$ echo $POD_B_IP
> 10.88.0.164  # output on my machine

```


Notice that the server is listening on pod's external network interface, port _80_:

请注意，服务器正在侦听 pod 的外部网络接口端口 _80_：

```bash
$ curl $POD_B_IP
> Service B: Yay!nounce 3494557023
$ curl $POD_B_IP
> Service B: Yay!nounce 1634910179
$ curl $POD_B_IP
> Service B: Yay!nounce 2013866549
$ curl $POD_B_IP
> Service B: Ooops... nounce 1258862891

```


Now we are ready to proceed with the service _A_. First, let's create a pod:

现在我们准备好继续服务_A_。首先，让我们创建一个 pod：

```bash
# Create a pod (read "server") for service A
$ sudo podman pod create --name service-a-pod \
    --add-host b.service:$POD_B_IP --publish 8080:80

```


Notice how we injected a DNS record like `b.service 10.88.0.164`. Since both pods reside in the same podman network, they can reach each other using assigned IP addresses. However, as of the time of writing this, podman doesn't provide DNS support for pods (yet). So, we have to maintain the mappings manually. Of course, we could use the plain IP address of the _B_'s pod while accessing the upstream from the service _A_ code. However, it's always nice to have human-readable hostnames instead of raw IP addresses. We will also see how this technique comes in handy with the envoy proxy sidecar below.

注意我们是如何注入一个像“b.service 10.88.0.164”这样的 DNS 记录的。由于两个 Pod 位于同一个 Podman 网络中，因此它们可以使用分配的 IP 地址相互访问。但是，截至撰写本文时，podman 尚不为 pod 提供 DNS 支持。因此，我们必须手动维护映射。当然，我们可以在从服务 _A_ 代码访问上游时使用 _B_ 的 pod 的普通 IP 地址。然而，拥有人类可读的主机名而不是原始 IP 地址总是好的。我们还将看到这种技术如何在下面的特使代理边车中派上用场。

Let's continue with the service itself. We need to build it and run inside the pod we've just created.

让我们继续服务本身。我们需要构建它并在我们刚刚创建的 pod 中运行。

_Click here to see service A Dockerfile._

_点击这里查看服务A Dockerfile。_

```dockerfile
FROM golang:1.15

# Build
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go mod download
RUN go build -o service-a

# Run
ENV SERVICE_HOST="0.0.0.0"
ENV SERVICE_PORT=80

ENV METRICS_HOST="0.0.0.0"
ENV METRICS_PORT=8081

ENV UPSTREAM_SERVICE="http://b.service/"

CMD ["/app/service-a"]

```


[source on GitHub](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-a/Dockerfile-service)

[GitHub 上的来源](https://github.com/iximiuz/envoy-playground/tree/master/basics/service-a/Dockerfile-service)

```bash
# Build service A image
$ sudo podman build -t service-a -f service-a/Dockerfile-service

# Start service A container in the pod
$ sudo podman run --name service-a -d --rm --pod service-a-pod \
   -e SERVICE_HOST=0.0.0.0 -e SERVICE_PORT=80 \
   -e UPSTREAM_SERVICE=http://b.service:80 \
service-a

# Keep pod's IP address for future use
$ POD_A_IP=$(sudo podman inspect -f "{{.NetworkSettings.IPAddress}}" \
    $(sudo podman pod inspect -f "{{.InfraContainerID}}" service-a-pod))

$ echo $POD_A_IP
> 10.88.0.165  # output on my machine

```


Remember the diagram from the beginning of this section. At this part of the exercise service _A_ has to be directly exposed to the outside world (i.e. the host machine) and it has to communicate with the service _B_ directly as well. That's why we made service _A_ listening on the pod's external network interface using `-e SERVICE_HOST=0.0.0.0 -e SERVICE_PORT=80` and provided it with the knowledge how to reach the service _B_ `-e UPSTREAM_SERVICE=http://b.service:80`.

记住本节开头的图表。在练习的这一部分，服务 _A_ 必须直接暴露给外部世界（即主机），并且它也必须直接与服务 _B_ 通信。这就是为什么我们使用 `-e SERVICE_HOST=0.0.0.0 -e SERVICE_PORT=80` 让服务 _A_ 监听 pod 的外部网络接口，并为它提供如何访问服务 _B_ `-e UPSTREAM_SERVICE=http://b 的知识.服务：80`。

The last preparation before pouring some traffic - starting a [prometheus](https://github.com/iximiuz/envoy-playground/tree/master/basics/prometheus) node:

注入流量前的最后准备——启动一个 [prometheus](https://github.com/iximiuz/envoy-playground/tree/master/basics/prometheus) 节点：

```bash
# Scrape configs
$ cat prometheus/prometheus.yml
> scrape_configs:
>   - job_name: service-a
>     scrape_interval: 1s
>     static_configs:
>       - targets: ['a.service:8081']
>
>   - job_name: service-b
>     scrape_interval: 1s
>     static_configs:
>       - targets: ['b.service:8081']

# Dockerfile
$ cat prometheus/Dockerfile
> FROM prom/prometheus:v2.20.1
> COPY prometheus.yml /etc/prometheus/prometheus.yml

# Build & run
$ sudo podman build -t envoy-prom -f prometheus/Dockerfile

$ sudo podman run --name envoy-prom -d --rm \
   --publish 9090:9090 \
   --add-host a.service:$POD_A_IP \
   --add-host b.service:$POD_B_IP \
envoy-prom

```


At this time, we have two services within a shared network. They can talk to each other using their IP addresses. Additionally, port _80_ of the service _A_ is mapped to the port _8080_ of the host machine and prometheus is exposed on the port _9090_. I intentionally made these two ports mapped on `0.0.0.0` 'cause I run the demo inside a VirtualBox machine. This way, I can access prometheus graphical interface from the laptop's host operating system via using `<vm_ip_address>:9090/graph`.

此时，我们在共享网络中有两个服务。他们可以使用他们的 IP 地址相互交谈。另外，服务_A_的_80_端口映射到主机的_8080_端口，prometheus暴露在_9090_端口。我故意将这两个端口映射到“0.0.0.0”，因为我在 VirtualBox 机器内运行演示。这样，我就可以通过使用 `<vm_ip_address>:9090/graph` 从笔记本电脑的主机操作系统访问 prometheus 图形界面。

Finally, we can send some traffic to the service _A_ and see what happens:

最后，我们可以向服务 _A_ 发送一些流量，看看会发生什么：

```bash
$ for _ in {1..1000};do curl --silent localhost:8080;done |sort |uniq -w 24 -c
>    1000
>     208 Service A: upstream failed with: HTTP 500 - Service B: Ooops... nounce 1007409508
>     792 Service A: upstream responded with: Service B: Yay!nounce 1008262846

```


Yay! 🎉 As expected, ca. 20% of the upstream requests failed with the HTTP 500 status code. Let's take a look at the prometheus metrics to see the per-service statistics:

好极了！ 🎉正如预期的那样，大约。 20% 的上游请求失败并显示 HTTP 500 状态代码。让我们看一下 prometheus 指标以查看每个服务的统计信息：

![Service A - 20% of outgoing requests failed](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/prom-service-a-direct.png)

_`service_a_requests_total`_

![Service B - 20% of incoming requests failed](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/prom-service-b-direct.png)

_`service_b_requests_total`_ 

Well, I believe it's not a surprise that both services handled each 1000 of requests and the service _A_ failed as many requests as the service _B_.

好吧，我相信这两个服务每处理 1000 个请求并且服务 _A_ 失败的请求与服务 _B_ 一样多，这并不奇怪。

#### Service A talks to service B through envoy proxy sidecar

#### 服务 A 通过特使代理 sidecar 与服务 B 对话

Let's enhance our setup by adding a service proxy sidecar to the service _A_. For the sake of simplicity of this demo, the only thing the sidecar will be doing is making up to 2 retries of the failed HTTP requests. Hopefully, it'll improve the overall service _A_ success rate. The desired setup will look as follows:

让我们通过向服务 _A_ 添加服务代理边车来增强我们的设置。为了这个演示的简单起见，sidecar 唯一要做的就是对失败的 HTTP 请求进行最多 2 次重试。希望它会提高整体服务_A_成功率。所需的设置如下所示：

![Multi-service demo setup with sidecar proxy intercepting traffic](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/80-demo-sidecar.png)

The main difference is that with the sidecar all incoming and outgoing requests will be passing through envoy. In contrast with the previous section setup, service _A_ will neither be exposed publicly nor will be allowed to contact service _B_ directly.

主要区别在于，使用 sidecar，所有传入和传出的请求都将通过 Envoy。与上一节设置相比，服务 _A_ 既不会公开，也不允许直接联系服务 _B_。

Let's review two scenarios:

让我们回顾两个场景：

- **Ingress**: a client sends an HTTP request to `$POD_A_IP`. It hits the envoy sidecar listening on `$POD_A_IP:80`. Envoy, in turn, makes a request to the service _A_ container listening on the pod's `localhost:8000`. Once the envoy process gets the response from the service _A_ container, it forwards it back to the client.

- **Ingress**：客户端向 `$POD_A_IP` 发送 HTTP 请求。它击中了监听 `$POD_A_IP:80` 的特使边车。反过来，Envoy 向服务 _A_ 容器发出请求，该容器侦听 pod 的“localhost:8000”。一旦特使进程从服务 _A_ 容器获得响应，它将其转发回客户端。

- **Egress**: service _A_ gets a request from envoy. To handle it, the service process needs to access the upstream service _B_. The service _A_ container sends a request to another envoy listener sitting on pod's `localhost:9001`. Additionally, service _A_ specifies the HTTP Host header `b.local.service` allowing envoy to route this request appropriately. When envoy receives a request on `localhost:9001` it knows that it's egress traffic. It checks the Host header and if it looks like the service _B_, it makes a request to `$POD_B_IP`.


- **Egress**：服务 _A_ 从 Envoy 获取请求。为了处理它，服务进程需要访问上游服务_B_。服务 _A_ 容器向位于 pod 的“localhost:9001”上的另一个特使侦听器发送请求。此外，服务 _A_ 指定了 HTTP 主机标头 `b.local.service`，允许 Envoy 适当地路由此请求。当 Envoy 在 localhost:9001 上收到请求时，它知道这是出口流量。它检查主机头，如果它看起来像服务 _B_，它会向`$POD_B_IP` 发出请求。


Configuring envoy could quickly become tricky due to its huge set of capabilities. However, the [official documentation](https://www.envoyproxy.io/docs/envoy/latest/) is a great place to start. It not only describes the configuration format but also highlights some best practices and explains some concepts. In particular, I suggest these two articles ["Life of a Request"](https://www.envoyproxy.io/docs/envoy/v1.15.0/intro/life_of_A_request) and ["Service to service only"](https://www.envoyproxy.io/docs/envoy/v1.15.0/intro/deployment_types/service_to_service) for a better understanding of the material.

由于其庞大的功能集，配置 envoy 很快就会变得棘手。然而，[官方文档](https://www.envoyproxy.io/docs/envoy/latest/) 是一个很好的起点。它不仅描述了配置格式，还突出了一些最佳实践并解释了一些概念。特别推荐这两篇文章 ["Life of a Request"](https://www.envoyproxy.io/docs/envoy/v1.15.0/intro/life_of_A_request)和["Service to service only"](https://www.envoyproxy.io/docs/envoy/v1.15.0/intro/life_of_A_request)以更好地理解材料。

From a very high-level overview, Envoy could be seen as a bunch of pipelines. A pipeline starts from the listener and then connected through a set of filters to some number of clusters, where a cluster is just a logical group of network endpoints. Trying to be less abstract:

从一个非常高级的概述来看，Envoy 可以被视为一堆管道。管道从侦听器开始，然后通过一组过滤器连接到一定数量的集群，其中集群只是网络端点的逻辑组。尽量不那么抽象：

```
# Ingress
listener 0.0.0.0:80
       |
http_connection_manager (filter)
       |
http_router (filter)
       |
local_service (cluster) [127.0.0.1:8000]

# Egress
listener 127.0.0.1:9001
       |
http_connection_manager (filter)
       |
http_router (filter)
       |
remote_service_b (cluster) [b.service:80]

```


_Click here to see envoy.yaml file._

_点击这里查看 envoy.yaml 文件。_

```yaml
static_resources:
listeners:
# Ingress
  - address:
      socket_address:
        address: 0.0.0.0
        port_value: 80
    filter_chains:
    - filters:
      - name: envoy.filters.network.http_connection_manager
        typed_config:
          "@type": type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager
          codec_type: auto
          stat_prefix: ingress_http
          access_log:
          - name: envoy.access_loggers.file
            typed_config:
              "@type": type.googleapis.com/envoy.config.accesslog.v2.FileAccessLog
              path: /dev/stdout
          route_config:
            name: ingress_route
            virtual_hosts:
            - name: local_service
              domains:
              - "*"
              routes:
              - match:
                  prefix: "/"
                route:
                  cluster: local_service
          http_filters:
          - name: envoy.filters.http.router
            typed_config: {}
# Egress
  - address:
      socket_address:
        address: 127.0.0.1
        port_value: 9001
    filter_chains:
    - filters:
      - name: envoy.filters.network.http_connection_manager
        typed_config:
          "@type": type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager
          codec_type: auto
          stat_prefix: egress_http
          access_log:
          - name: envoy.access_loggers.file
            typed_config:
              "@type": type.googleapis.com/envoy.config.accesslog.v2.FileAccessLog
              path: /dev/stdout
          route_config:
            name: egress_route
            virtual_hosts:
            - name: remote_service_b
              domains:
              - "b.local.service"
              - "b.local.service:*"
              retry_policy:
                retry_on: 5xx
                num_retries: 2
              routes:
              - match:
                  prefix: "/"
                route:
                  cluster: remote_service_b
          http_filters:
          - name: envoy.filters.http.router
            typed_config: {}

clusters:
  - name: local_service
    connect_timeout: 0.25s
    type: strict_dns
    lb_policy: round_robin
    load_assignment:
      cluster_name: local_service
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 8000
  - name: remote_service_b
    connect_timeout: 0.25s
    type: strict_dns
    lb_policy: round_robin
    load_assignment:
      cluster_name: remote_service_b
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: b.service
                port_value: 80

admin:
access_log_path: "/dev/stdout"
address:
    socket_address:
      # Beware: it's insecure to expose admin interface publicly
      address: 0.0.0.0
      port_value: 9901

```

[source on GitHub](https://github.com/iximiuz/envoy-playground/blob/master/basics/service-a/envoy.yaml)

Envoy is famous for its observability capabilities. It exposes various statistic information and luckily for us, it supports the prometheus metrics format out of the box. We can extend the prometheus scrape configs adding the following section:

Envoy 以其可观察性能力而闻名。它公开了各种统计信息，幸运的是，它支持开箱即用的普罗米修斯指标格式。我们可以扩展 prometheus 抓取配置，添加以下部分：

```bash
# prometheus/prometheus.yml

  - job_name: service-a-envoy
    scrape_interval: 1s
    metrics_path: /stats/prometheus
    static_configs:
      - targets: ['a.service:9901']

```


To build the envoy sidecar image we can run:

要构建 envoy sidecar 镜像，我们可以运行：

```bash
$ sudo podman build -t service-a-envoy -f service-a/Dockerfile-envoy

```


We don't need to rebuild service images since they've been made configurable via environment variables. However, we need to recreate the service _A_ to make it listening on the pod's `localhost:8000`.

我们不需要重建服务映像，因为它们已经通过环境变量进行了配置。但是，我们需要重新创建服务 _A_ 以使其侦听 pod 的“localhost:8000”。

```bash
# Stop existing container and remove pod
$ sudo podman kill service-a
$ sudo podman pod rm service-a-pod

$ sudo podman pod create --name service-a-pod \
   --add-host b.local.service:127.0.0.1 \
   --add-host b.service:$POD_B_IP \
   --publish 8080:80

$ sudo podman run --name service-a -d --rm --pod service-a-pod \
   -e SERVICE_HOST=127.0.0.1 \
   -e SERVICE_PORT=8000 \
   -e UPSTREAM_SERVICE=http://b.local.service:9001 \
service-a

$ sudo podman run --name service-a-envoy -d --rm --pod service-a-pod \
   -e ENVOY_UID=0 -e ENVOY_GID=0 service-a-envoy

$ POD_A_IP=$(sudo podman inspect -f "{{.NetworkSettings.IPAddress}}" \
    $(sudo podman pod inspect -f "{{.InfraContainerID}}" service-a-pod))

```


Let's see what happens if we pour some more traffic:

让我们看看如果我们注入更多流量会发生什么：

```bash
$ for _ in {1..1000};do curl --silent localhost:8080;done |sort |uniq -w 24 -c
>    1000
>       9 Service A: upstream failed with: HTTP 500 - Service B: Ooops... nounce 1263663296
>     991 Service A: upstream responded with: Service B: Yay!nounce 1003014939

```


Hooray! 🎉 Seems like the success rate of the service _A_ jumped from 80% to 99%! Well, that's great, but also as expected. The original probability to get HTTP 500 from the service _A_ was equal to the probability of the service _B_ to fail a request since the network conditions are kind of ideal here. But since the introduction of the envoy sidecar, service _A_ got a superpower of retries. The probability to fail 3 consequent requests with a 20% chance of a single attempt failure is `0.2 * 0.2 * 0.2 = 0.008`, i.e very close to 1%. Thus, we theoretically confirmed the observed 99% success rate.

万岁！ 🎉 好像服务_A_的成功率从80%跳到了99%！嗯，这很棒，但也符合预期。从服务 _A_ 获得 HTTP 500 的原始概率等于服务 _B_ 请求失败的概率，因为这里的网络条件有点理想。但是自从引入了 envoy sidecar，服务 _A_ 获得了重试的超级能力。失败 3 个后续请求的概率为 20% 的单次尝试失败概率为“0.2 * 0.2 * 0.2 = 0.008”，即非常接近 1%。因此，我们从理论上证实了观察到的 99% 成功率。

Last but not least, let's check out the metrics. We will start from the familiar `service_a_requests_total` counter:

最后但并非最不重要的一点，让我们来看看指标。我们将从熟悉的“service_a_requests_total”计数器开始：

![Service A - only 1% of outgoing requests failed](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/prom-service-a-envoy.png)

_`service_a_requests_total`_

Well, seems like service _A_ again got 1000 requests, but this time it failed only a tiny fraction of it. What's up with service _B_?

好吧，似乎 service _A_ 再次收到了 1000 个请求，但这次它只失败了其中的一小部分。服务_B_怎么了？

![Service B - still 20% of incoming requests failed](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/prom-service-b-envoy.png)

_`service_b_requests_total`_



Here we definitely can see the change. Instead of the original 1000, this time service _B_ got about 1250 requests in total. However, only about 1000 have been served successfully.

在这里我们绝对可以看到变化。这次服务 _B_ 总共收到了大约 1250 个请求，而不是原来的 1000 个。但是，只有大约 1000 个成功送达。

What can the envoy sidecar tell us?

特使边车可以告诉我们什么？

![Envoy local cluster stats](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/prom-envoy-local-service.png)

_`envoy_cluster_upstream_rq{envoy_cluster_name="local_service"}`_



![Envoy remote cluster stats](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/prom-envoy-remote-service.png)

_`envoy_cluster_upstream_rq{envoy_cluster_name="remote_service_b"}`_



While both `local_service` and `remote_service_b` clusters don't shed much light on the actual number of retries that were made, there is another metric we can check:

虽然“local_service”和“remote_service_b”集群都没有说明实际重试次数，但我们可以检查另一个指标：

![Envoy retries stats](http://iximiuz.com/service-proxy-pod-sidecar-oh-my/prom-envoy-retries.png)

_`envoy_cluster_retry_upstream_rq{envoy_cluster_name="remote_service_b"}`_



Perfect, we managed to confirm that all those ~250 extra requests the service _B_ received are actually retries originated by the envoy sidecar!

完美，我们设法确认服务 _B_ 收到的所有大约 250 个额外请求实际上是由特使 sidecar 发起的重试！

## Instead of conclusion 

## 而不是结论

I hope you enjoyed playing with all these pods and sidecars as much as I did. It's always beneficial to build such demos from time to time because often the amount of insights you're going to get while working on it is underestimated. So, I encourage everyone to get the hands dirty and share your findings! See you next time!

我希望你和我一样喜欢玩所有这些吊舱和边车。时不时地构建这样的演示总是有益的，因为通常会低估您在工作时获得的洞察力。因此，我鼓励每个人都动手并分享您的发现！下次见！

Make code, not war!

编写代码，而不是战争！

### Other posts you may like

### 您可能喜欢的其他帖子

- [Service Discovery in Kubernetes - Combining the Best of Two Worlds](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/)
- [Traefik: canary deployments with weighted load balancing](http://iximiuz.com/en/posts/traefik-canary-deployments-with-weighted-load-balancing/)

- [Kubernetes 中的服务发现 - 结合两全其美](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/)
- [Traefik：具有加权负载平衡的金丝雀部署](http://iximiuz.com/en/posts/traefik-canary-deployments-with-weighted-load-balancing/)

[envoy,](javascript: void 0) [microservices,](javascript: void 0) [architecture](javascript: void 0)

[envoy,](javascript: void 0) [微服务,](javascript: void 0) [架构](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

