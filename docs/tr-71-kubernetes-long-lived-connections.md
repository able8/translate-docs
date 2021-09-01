# Load balancing and scaling long-lived connections in Kubernetes

# Kubernetes 中的负载平衡和扩展长连接

Published in February 2020

> **TL;DR:** Kubernetes doesn't load balance long-lived connections, and some Pods might receive more requests than others. If you're using HTTP/2, gRPC,  RSockets, AMQP or any other long-lived connection such as a database  connection, you might want to consider client-side load balancing.

> **TL;DR:** Kubernetes 不会对长期连接进行负载均衡，并且某些 Pod 可能会收到比其他 Pod 更多的请求。如果您使用 HTTP/2、gRPC、RSockets、AMQP 或任何其他长期连接（例如数据库连接），您可能需要考虑客户端负载平衡。

Kubernetes offers two convenient abstractions to deploy apps: Services and Deployments. Deployments describe a recipe for what kind and how many copies of your app should run at any given time. Each app is deployed as a Pod, and an IP address is assigned to it.  Services, on the other hand, are similar to load balancers. They are designed to distribute the traffic to a set of Pods.

Kubernetes 提供了两个方便的抽象来部署应用程序：服务和部署。部署描述了在任何给定时间应该运行的应用程序的种类和数量的方法。每个应用程序都部署为一个 Pod，并为其分配一个 IP 地址。另一方面，服务类似于负载均衡器。它们旨在将流量分配到一组 Pod。

![In this diagram you have three instances of a single app and a load balancer.](https://learnk8s.io/a/c62b844725054a789045193c5af5d245.svg)

1/4 In this diagram you have three instances of a single app and a load balancer. Next

1/4 在此图中，您有单个应用程序和负载均衡器的三个实例。下一个

It is often useful to think about Services as a collection of IP address.

将服务视为 IP 地址的集合通常很有用。

Every time you make a request to a Service, one of the IP addresses from that list is selected and used as the destination.

每次向服务发出请求时，都会选择该列表中的一个 IP 地址并将其用作目标。

![Imagine issuing a request such as curl 10.96.45.152 to the Service.](https://learnk8s.io/a/bfad9f818be47768e09e7e87aa9eafbe.svg)

1/3 Imagine issuing a request such as `curl 10.96.45.152` to the Service. Next

1/3 想象一下向 Service 发出一个诸如“curl 10.96.45.152”之类的请求。下一个

If you have two apps such as a front-end and a backend, you can use a  Deployment and a Service for each and deploy them in the cluster.

如果您有两个应用程序，例如前端和后端，您可以为每个应用程序使用一个部署和一个服务，并将它们部署在集群中。

When the front-end app makes a request, it doesn't need to know how many Pods are connected to the backend Service.

当前端应用发出请求时，它不需要知道有多少 Pod 连接到后端 Service。

It could be one Pod, tens or hundreds.

它可以是一个 Pod、数十个或数百个。

The front-end app isn't aware of the individual IP addresses of the backend app either.

前端应用程序也不知道后端应用程序的各个 IP 地址。

When it wants to make a request, that request is sent to the backend Service which has an IP address that doesn't change.

当它想要发出请求时，该请求被发送到具有不变 IP 地址的后端服务。

![The red Pod issues a request to an internal (beige) component.Instead of choosing one of the Pod as the destination, the red Pod issues the request to the Service.](https://learnk8s.io/a/5a7c9a225135ac88cfa88fbb24b8745b.svg)

红色 Pod 没有选择其中一个 Pod 作为目的地，而是向 Service 发出请求。](https://learnk8s.io/a/5a7c9a225135ac88cfa88fbb24b8745b.svg)

1/4 The red Pod issues a request to an internal (beige) component. Instead of  choosing one of the Pod as the destination, the red Pod issues the  request to the Service. Next

1/4 红色 Pod 向内部（米色）组件发出请求。红色 Pod 不是选择其中一个 Pod 作为目的地，而是向 Service 发出请求。下一个

*But what's the load balancing strategy for the Service?*

*但是服务的负载平衡策略是什么？*

*It is round-robin, right?*

*这是循环，对吧？*

Sort of.

有点。

## Load balancing in Kubernetes Services

## Kubernetes 服务中的负载平衡

Kubernetes Services don't exist.

Kubernetes 服务不存在。

There's no process listening on the IP address and port of the Service.

没有进程侦听服务的 IP 地址和端口。

> You can check that this is the case by accessing any node in your Kubernetes cluster and executing `netstat -ntlp`.

> 您可以通过访问 Kubernetes 集群中的任何节点并执行 `netstat -ntlp` 来检查是否是这种情况。

Even the IP address can't be found anywhere.

甚至在任何地方都找不到IP地址。

The IP address for a Service is allocated by the control plane in the controller manager and stored in the database — etcd.

Service 的 IP 地址由控制器管理器中的控制平面分配并存储在数据库 - etcd 中。

That same IP address is then used by another component: kube-proxy.

然后另一个组件使用相同的 IP 地址：kube-proxy。

Kube-proxy reads the list of IP addresses for all Services and writes a collection of iptables rules in every node.

Kube-proxy 读取所有服务的 IP 地址列表，并在每个节点中写入一组 iptables 规则。

The rules are meant to say: "if you see this Service IP address, instead  rewrite the request and pick one of the Pod as the destination".

规则旨在说明：“如果您看到此服务 IP 地址，请改写请求并选择其中一个 Pod 作为目的地”。

The Service IP address is used only as a placeholder — that's why there is no process listening on the IP address or port.

服务 IP 地址仅用作占位符 — 这就是为什么没有进程侦听 IP 地址或端口的原因。

![Consider a cluster with three Nodes.Each Node has a Pod deployed.](https://learnk8s.io/a/a3e195d7ee05e50797e7444aacabc31d.svg)

每个节点都部署了一个 Pod。](https://learnk8s.io/a/a3e195d7ee05e50797e7444aacabc31d.svg)

1/8 Consider a cluster with three Nodes. Each Node has a Pod deployed. Next

1/8 考虑一个具有三个节点的集群。每个节点都部署了一个 Pod。下一个

*Does iptables use round-robin?*

*iptables 是否使用循环？*

No, iptables is primarily used for firewalls, and it is not designed to do load balancing.

不，iptables 主要用于防火墙，它不是为了做负载平衡而设计的。

However, you could [craft a smart set of rules that could make iptables behave like a load balancer](https://scalingo.com/blog/iptables#load-balancing).

但是，您可以[制定一组智能规则，使 iptables 表现得像负载均衡器](https://scalingo.com/blog/iptables#load-balancing)。

And this is precisely what happens in Kubernetes.

而这正是 Kubernetes 中发生的事情。

If you have three Pods, kube-proxy writes the following rules:

如果你有三个 Pod，kube-proxy 会写如下规则：

1. select Pod 1 as the destination with a likelihood of 33%. Otherwise, move to the next rule
2. choose Pod 2 as the destination with a probability of 50%. Otherwise, move to the following rule
3. select Pod 3 as the destination (no probability)


1. 选择 Pod 1 作为目的地，可能性为 33%。否则，移至下一条规则
2. 以50%的概率选择Pod 2作为目的地。否则，转到以下规则
3. 选择Pod 3作为目的地（无概率）

The compound probability is that Pod 1, Pod 2 and Pod 3 have all have a one-third chance (33%) to be selected.

复合概率是 Pod 1、Pod 2 和 Pod 3 都有三分之一的机会 (33%) 被选中。

![iptables rules for three Pods](https://learnk8s.io/a/fbbcbf56da96099c15314b9af4a0dd77.svg)

Also, there's no guarantee that Pod 2 is selected after Pod 1 as the destination.

此外，不能保证在 Pod 1 之后选择 Pod 2 作为目的地。

> Iptables use the [statistic module](http://ipset.netfilter.org/iptables-extensions.man.html#lbCD) with `random` mode. So the load balancing algorithm is random.

> iptables 使用 [统计模块](http://ipset.netfilter.org/iptables-extensions.man.html#lbCD) 和 `random` 模式。所以负载均衡算法是随机的。

Now that you're familiar with how Services work let's have a look at more exciting scenarios.

现在您已经熟悉了服务的工作方式，让我们来看看更令人兴奋的场景。

## Long-lived connections don't scale out of the box in Kubernetes 
## 长期连接不会在 Kubernetes 中开箱即用

With every HTTP request started from the front-end to the backend, a new TCP connection is opened and closed.

每个 HTTP 请求从前端开始到后端，都会打开和关闭一个新的 TCP 连接。

If the front-end makes 100 HTTP requests per second to the backend, 100  different TCP connections are opened and closed in that second.

如果前端每秒向后端发出 100 个 HTTP 请求，那一秒内将打开和关闭 100 个不同的 TCP 连接。

You can improve the latency and save resources if you open a TCP connection and reuse it for any subsequent HTTP requests.

如果您打开 TCP 连接并将其重用于任何后续 HTTP 请求，则可以改善延迟并节省资源。

The HTTP protocol has a feature called HTTP keep-alive, or HTTP connection  reuse that uses a single TCP connection to send and receive multiple  HTTP requests and responses.

HTTP 协议有一个称为 HTTP keep-alive 或 HTTP 连接重用的功能，它使用单个 TCP 连接来发送和接收多个 HTTP 请求和响应。

![Opening and closing connections VS HTTP connection reuse](https://learnk8s.io/a/83f9bebffbbbee8d6a25315b8679f9df.svg)

It doesn't work out of the box; your server and client should be configured to use it.

它不能开箱即用；您的服务器和客户端应该配置为使用它。

The change itself is straightforward, and it's available in most languages and frameworks.

更改本身很简单，并且在大多数语言和框架中都可用。

Here a few examples on how to implement keep-alive in different languages:

这里有一些关于如何在不同语言中实现 keep-alive 的示例：

- [Keep-alive in Node.js](https://medium.com/@onufrienkos/keep-alive-connection-on-inter-service-http-requests-3f2de73ffa1)
- [Keep-alive in Spring boot](https://www.baeldung.com/httpclient-connection-management)
- [Keep-alive in Python](https://blog.insightdatascience.com/learning-about-the-http-connection-keep-alive-header-7ebe0efa209d)
- [Keep-alive in .NET](https://docs.microsoft.com/en-us/dotnet/api/system.net.httpwebrequest.keepalive?view=netframework-4.8)

- [在 Node.js 中保持活动状态](https://medium.com/@onufrienkos/keep-alive-connection-on-inter-service-http-requests-3f2de73ffa1)
- [在 Spring 启动时保持活动状态](https://www.baeldung.com/httpclient-connection-management)
- [在 Python 中保持活动](https://blog.insightdatascience.com/learning-about-the-http-connection-keep-alive-header-7ebe0efa209d)
- [在 .NET 中保持活动状态](https://docs.microsoft.com/en-us/dotnet/api/system.net.httpwebrequest.keepalive?view=netframework-4.8)

*What happens when you use keep-alive with a Kubernetes Service?*

*当你对 Kubernetes 服务使用 keep-alive 时会发生什么？*

Let's imagine that front-end and backend support keep-alive.

让我们想象一下前端和后端都支持 keep-alive。

You have a single instance of the front-end and three replicas for the backend.

您有一个前端实例和三个后端副本。

The front-end makes the first request to the backend and opens the TCP connection.

前端向后端发出第一个请求并打开 TCP 连接。

The request reaches the Service, and one of the Pod is selected as the destination.

请求到达 Service，选择其中一个 Pod 作为目的地。

The backend Pod replies and the front-end receives the response.

后端 Pod 回复，前端接收响应。

But instead of closing the TCP connection, it is kept open for subsequent HTTP requests.

但是它并没有关闭 TCP 连接，而是为后续的 HTTP 请求保持打开状态。

*What happens when the front-end issues more requests?*

*当前端发出更多请求时会发生什么？*

They are sent to the same Pod.

它们被发送到同一个 Pod。

*Isn't iptables supposed to distribute the traffic?*

*iptables 不应该分配流量吗？*

It is.

这是。

There is a single TCP connection open, and iptables rule were invocated the first time.

有一个 TCP 连接打开，并且第一次调用了 iptables 规则。

One of the three Pods was selected as the destination.

三个 Pod 之一被选为目的地。

Since all subsequent requests are channelled through the same TCP connection, [iptables isn't invoked anymore.](https://scalingo.com/blog/iptables#load-balancing)

由于所有后续请求都通过相同的 TCP 连接传送，因此 [iptables 不再被调用。](https://scalingo.com/blog/iptables#load-balancing)

- ![The red Pod issues a request to the Service.](https://learnk8s.io/a/3d9434491906a7f16962b3ca12cf896e.svg)

1/5 The red Pod issues a request to the Service. Next

1/5 红色 Pod 向 Service 发出请求。下一个

So you have now achieved better latency and throughput, but you lost the ability to scale your backend.

因此，您现在已经实现了更好的延迟和吞吐量，但您失去了扩展后端的能力。

Even if you have two backend Pods that can receive requests from the frontend Pod, only one is actively used.

即使你有两个后端 Pod 可以接收来自前端 Pod 的请求，也只有一个是主动使用的。

*Is it fixable?*

*它可以修复吗？*

Since Kubernetes doesn't know how to load balance persistent connections, you could step in and fix it yourself.

由于 Kubernetes 不知道如何对持久连接进行负载平衡，因此您可以自行介入并修复它。

Services are a collection of IP addresses and ports — usually called endpoints.

服务是 IP 地址和端口的集合——通常称为端点。

Your app could retrieve the list of endpoints from the Service and decide how to distribute the requests.

您的应用程序可以从服务中检索端点列表并决定如何分发请求。

As a first try, you could open a persistent connection to every Pod and round-robin requests to them.

作为第一次尝试，您可以打开与每个 Pod 的持久连接并向它们发送循环请求。

Or you could [implement more sophisticated load balancing algorithms](https://blog.twitter.com/engineering/en_us/topics/infrastructure/2019/daperture-load-balancer.html).

或者你可以[实现更复杂的负载平衡算法](https://blog.twitter.com/engineering/en_us/topics/infrastructure/2019/daperture-load-balancer.html)。

The client-side code that executes the load balancing should follow the logic below:

执行负载均衡的客户端代码应遵循以下逻辑：

1. retrieve a list of endpoints from the Service
2. for each of them, open a connection and keep it open
3. when you need to make a request, pick one of the open connections
4. on a regular interval refresh the list of endpoints and remove or add new connections


1. 从服务中检索端点列表
2. 为每个人打开一个连接并保持打开状态
3. 当您需要提出请求时，选择其中一个打开的连接
4. 定期刷新端点列表并删除或添加新连接

![Instead of having the red Pod issuing a request to your Service, you could load balance the request client-side.](https://learnk8s.io/a/bd7b0d3f4b3d509a113719b1f5d12e48.svg)

1/4 Instead of having the red Pod issuing a request to your Service, you could load balance the request client-side. Next

1/4 您可以在客户端对请求进行负载平衡，而不是让红色 Pod 向您的服务发出请求。下一个

*Does this problem apply only to HTTP keep-alive?*

*此问题是否仅适用于 HTTP keep-alive？*

## Client-side load balancing

## 客户端负载均衡

HTTP isn't the only protocol that can benefit from long-lived TCP connections.

HTTP 不是唯一可以从长期 TCP 连接中受益的协议。

If your app uses a database, the connection isn't opened and closed every time you wish to retrieve a record or a document.

如果您的应用程序使用数据库，则每次您希望检索记录或文档时都不会打开和关闭连接。

Instead, the TCP connection is established once and kept open.

相反，TCP 连接建立一次并保持打开状态。

If your database is deployed in Kubernetes using a Service, you might experience the same issues as the previous example.

如果您的数据库使用 Service 部署在 Kubernetes 中，您可能会遇到与前一个示例相同的问题。

There's one replica in your database that is utilised more than the others.

您的数据库中有一个副本比其他副本使用得更多。

Kube-proxy and Kubernetes don't help to balance persistent connections.

Kube-proxy 和 Kubernetes 无助于平衡持久连接。

Instead, you should take care of load balancing the requests to your database.

相反，您应该负责对数据库请求进行负载平衡。

Depending on the library that you use to connect to the database, you might have different options.

根据用于连接到数据库的库，您可能有不同的选择。

The following example is from a clustered MySQL database called from Node.js:

以下示例来自从 Node.js 调用的集群 MySQL 数据库：

index.js

```
var mysql = require('mysql'); 

var poolCluster = mysql.createPoolCluster();
var endpoints = /* retrieve endpoints from the Service */
for (var [index, endpoint] of endpoints) {
   poolCluster.add(`mysql-replica-${index}`, endpoint);
}
// Make queries to the clustered MySQL database
```

As you can imagine, several other protocols work over long-lived TCP connections.

可以想象，其他几种协议在长期存在的 TCP 连接上工作。

Here you can read a few examples:

在这里，您可以阅读一些示例：

- Websockets and secured WebSockets
- HTTP/2
- gRPC
- RSockets
- AMQP

- Websockets 和安全的 WebSockets
- HTTP/2
- gRPC
- RSockets
- AMQP

You might recognise most of the protocols above.

您可能会认识上面的大多数协议。

*So if these protocols are so popular, why isn't there a standard answer to load balancing?*

*因此，如果这些协议如此流行，为什么没有负载平衡的标准答案？*

*Why does the logic have to be moved into the client?*

*为什么必须将逻辑移到客户端？*

*Is there a native solution in Kubernetes?*

*Kubernetes 有原生解决方案吗？*

Kube-proxy and iptables are designed to cover the most popular use cases of deployments in a Kubernetes cluster.

Kube-proxy 和 iptables 旨在涵盖 Kubernetes 集群中最流行的部署用例。

But they are mostly there for convenience.

但它们主要是为了方便。

If you're using a web service that exposes a REST API, then you're in luck — this use case usually doesn't reuse TCP connections, and you can use  any Kubernetes Service.

如果您使用的是公开 REST API 的 Web 服务，那么您很幸运——这个用例通常不会重用 TCP 连接，您可以使用任何 Kubernetes 服务。

But as  soon as you start using persistent TCP connections, you should look into how you can evenly distribute the load to your backends.

但是一旦您开始使用持久 TCP 连接，您就应该研究如何将负载平均分配到后端。

Kubernetes doesn't cover that specific use case out of the box.

Kubernetes 没有涵盖开箱即用的特定用例。

However, there's something that could help.

但是，有一些东西可以提供帮助。

## Load balancing long-lived connections in Kubernetes

## Kubernetes 中的负载均衡长连接

Kubernetes has four different kinds of Services:

Kubernetes 有四种不同类型的服务：

1. ClusterIP
2. NodePort
3. LoadBalancer
4. Headless

1. 集群IP
2. 节点端口
3. 负载均衡器
4. 无头

The first three Services have a virtual IP address that is used by kube-proxy to create iptables rules.

前三个服务有一个虚拟 IP 地址，kube-proxy 使用它来创建 iptables 规则。

But the fundamental building block of all kinds of the Services is the Headless Service.

但是各种服务的基本构建块是无头服务。

The headless Service doesn't have an assigned IP address and is only a  mechanism to collect a list of Pod IP addresses and ports (also called  endpoints).

Headless Service 没有分配的 IP 地址，它只是一种收集 Pod IP 地址和端口（也称为端点）列表的机制。

Every other Service is built on top of the Headless Service.

所有其他服务都建立在无头服务之上。

The ClusterIP Service is a Headless Service with some extra features:

ClusterIP 服务是具有一些额外功能的 Headless 服务：

- the control plane assigns it an IP address
- kube-proxy iterates through all the IP addresses and creates iptables rules

- 控制平面为其分配一个 IP 地址
- kube-proxy 遍历所有 IP 地址并创建 iptables 规则

So you could ignore kube-proxy all together and always use the list of  endpoints collected by the Headless Service to load balance requests client-side.

因此，您可以完全忽略 kube-proxy，并始终使用 Headless Service 收集的端点列表来负载平衡客户端请求。

*But can you imagine adding that logic to all apps deployed in the cluster?* If you have an existing fleet of applications, this might sound like an impossible task. But there's an alternative.

*但是您能想象将这种逻辑添加到集群中部署的所有应用程序吗？* 如果您有一组现有的应用程序，这听起来可能是一项不可能完成的任务。但还有一个选择。

## Service meshes to the rescue

## 服务网格来救援

You probably already noticed that the client-side load balancing strategy is quite standard.

您可能已经注意到客户端负载平衡策略是非常标准的。

When the app starts, it should
- retrieve a list of IP addresses from the Service
- open and maintain a pool of connections
- periodically refresh the pool by adding and removing endpoints

当应用程序启动时，它应该
- 从服务中检索 IP 地址列表
- 打开并维护一个连接池
- 通过添加和删除端点定期刷新池

As soon as it wishes to make a request, it should:

一旦它希望提出请求，它应该：

- pick one of the available connections using a predefined logic such as round-robin
- issue the request

- 使用预定义的逻辑（例如循环）选择可用连接之一
- 发出请求

The steps above are valid for WebSockets connections as well as gRPC and AMQP. You could extract that logic in a separate library and share it with all apps. Instead of writing a library from scratch, you could use a Service mesh such as Istio or Linkerd.

上述步骤适用于 WebSockets 连接以及 gRPC 和 AMQP。 您可以在单独的库中提取该逻辑并与所有应用程序共享。您可以使用服务网格，例如 Istio 或 Linkerd，而不是从头开始编写库。

Service meshes augment your app with a new process that:

服务网格通过一个新流程来增强您的应用：

- automatically discovers IP addresses Services
- inspects connections such as WebSockets and gRPC
- load-balances requests using the right protocol

- 自动发现 IP 地址服务
- 检查 WebSockets 和 gRPC 等连接
- 使用正确的协议负载平衡请求

Service meshes can help you to manage the traffic inside your cluster, but they aren't exactly lightweight.

服务网格可以帮助您管理集群内的流量，但它们并不完全是轻量级的。

Other options include using a library such as Netflix Ribbon, a programmable proxy such as Envoy or just ignore it.

其他选项包括使用诸如 Netflix Ribbon 之类的库、诸如 Envoy 之类的可编程代理或直接忽略它。

*What happens if you ignore it?*

*如果你忽略它会怎样？*

You can ignore the load balancing and still don't notice any change.

您可以忽略负载平衡，但仍然不会注意到任何变化。

There are a couple of scenarios that you should consider.

您应该考虑几种情况。

If you have more clients than servers, there should be limited issues.

如果您的客户端多于服务器，则问题应该有限。

Imagine you have five clients opening persistent connections to two servers.

假设您有五个客户端打开与两台服务器的持久连接。

Even if there's no load balancing, both servers likely utilised.

即使没有负载平衡，两台服务器也可能被利用。

![More clients than servers](https://learnk8s.io/a/317efc5e7a260b1f477a95e3b1f101f9.svg)

The connections might not be distributed evenly (perhaps four ended up connecting to the same server), but overall there's a good chance that  both servers are utilised.

连接可能不会均匀分布（可能有四个最终连接到同一台服务器），但总体而言很有可能同时使用两台服务器。

What's more problematic is the opposite scenario.

更成问题的是相反的情况。

If you have fewer clients and more servers, you might have some underutilised resources and a potential bottleneck.

如果您有更少的客户端和更多的服务器，您可能会有一些未充分利用的资源和潜在的瓶颈。

Imagine having two clients and five servers.

想象一下有两个客户端和五个服务器。

At best, two persistent connections to two servers are opened.

充其量，打开了到两个服务器的两个持久连接。

The remaining servers are not used at all.

其余的服务器根本没有使用。

![More servers than clients](https://learnk8s.io/a/cada9a6e57b5440f779aef1f6c758494.svg)

If the two servers can't handle the traffic generated by the clients, horizontal scaling won't help.

如果两台服务器无法处理客户端产生的流量，水平扩展将无济于事。

## Summary

## 概括

Kubernetes Services are designed to cover most common uses for web applications.  However, as soon as you start working with application protocols that use  persistent TCP connections, such as databases, gRPC, or WebSockets, they fall apart. Kubernetes doesn't offer any built-in mechanism to load balance long-lived TCP connections.

Kubernetes 服务旨在涵盖 Web 应用程序的最常见用途。 但是，一旦您开始使用使用持久 TCP 连接的应用程序协议（例如数据库、gRPC 或 WebSockets），它们就会崩溃。Kubernetes 不提供任何内置机制来负载平衡长期 TCP 连接。

Instead, you should code your application so that it can retrieve and load balance upstreams client-side.

相反，您应该编码您的应用程序，以便它可以检索和负载平衡上游客户端。

Many thanks to [Daniel Weibel](https://medium.com/@weibeld), [Gergely Risko](https://github.com/errge) and [Salman Iqbal](https://twitter.com/soulmaniqbal ) for offering some invaluable suggestions.

非常感谢 [Daniel Weibel](https://medium.com/@weibeld)、[Gergely Risko](https://github.com/errge) 和 [Salman Iqbal](https://twitter.com/soulmaniqbal ) 提供了一些宝贵的建议。

And to [Chris Hanson](https://twitter.com/CloudNativChris) who suggested to include the detailed explanation (and flow chart) on how iptables rules work in practice. 
还有 [Chris Hanson](https://twitter.com/CloudNativChris)，他建议包含关于 iptables 规则在实践中如何工作的详细解释（和流程图）。