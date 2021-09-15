# Service Discovery in Kubernetes - Combining the Best of Two Worlds

# Kubernetes 中的服务发现 - 结合两个世界的优点

December 6, 2020 (Updated: July 31, 2021)

Before jumping to any Kubernetes specifics, let's talk about the service discovery problem in general.

在跳到任何 Kubernetes 细节之前，让我们先谈谈服务发现问题。

## What is Service Discovery

## 什么是服务发现

In the world of web service development, it's a common practice to run multiple copies of a service at the same time. Every such copy is a separate instance of the service represented by a network endpoint (i.e. some _IP_ and _port_) exposing the service API. Traditionally, virtual or physical machines have been used to host such endpoints, with the shift towards containers in more recent times. Having multiple instances of the service running simultaneously increases its availability and helps to adjust the service capacity to meet the traffic demand. On the other hand, it also complicates the overall setup - before accessing the service, a client (the term _client_ is intentionally used loosely here; oftentimes a client of some service is another service) needs to figure out the actual IP address and the port it should use. The situation becomes even more tricky if we add the ephemeral nature of instances to the equation. New instances come and existing instances go because of the non-zero failure rate, up- and downscaling, or maintenance. That's how a so-called **_service discovery_** problem arises.

在 Web 服务开发领域，同时运行一个服务的多个副本是一种常见的做法。每个这样的副本都是由暴露服务 API 的网络端点（即某些 _IP_ 和 _port_）表示的服务的单独实例。传统上，虚拟机或物理机已被用于托管此类端点，最近逐渐转向容器。同时运行多个服务实例可提高其可用性，并有助于调整服务容量以满足流量需求。另一方面，它也使整体设置复杂化——在访问服务之前，客户端（术语 _client_ 在这里故意松散地使用；通常某个服务的客户端是另一个服务）需要弄清楚实际的 IP 地址和端口它应该使用。如果我们将实例的短暂性质添加到等式中，情况会变得更加棘手。由于非零故障率、扩展和缩减或维护，新实例出现和现有实例消失。这就是所谓的**_服务发现_**问题的产生方式。

![Service discovery problem](http://iximiuz.com/service-discovery-in-kubernetes/service-discovery-problem-4000-opt.png)

_Service discovery problem._

_服务发现问题。_

## Server-Side Service Discovery

## 服务器端服务发现

A pretty common way of solving the service discovery problem is putting a load balancer _aka_ reverse proxy (e.g. Nginx or HAProxy) in front of the group of instances constituting a single service. An address (i.e. DNS name or less frequently IP) of such a load balancer is a much more stable piece of information. It can be communicated to the clients on development or configuration stages and assumed invariable during a single client lifespan. Then, from the client standpoint, accessing the multi-instance service is no different from accessing a single network endpoint. In other words, [the service discovery happens completely on the server-side](https://microservices.io/patterns/server-side-discovery.html).

解决服务发现问题的一种非常常见的方法是在构成单个服务的实例组之前放置一个负载均衡器_aka_反向代理（例如 Nginx 或 HAProxy）。这种负载均衡器的地址（即 DNS 名称或频率较低的 IP）是一条更稳定的信息。它可以在开发或配置阶段传达给客户端，并假设在单个客户端生命周期内保持不变。然后，从客户端的角度来看，访问多实例服务与访问单个网络端点没有什么不同。换句话说，[服务发现完全发生在服务器端](https://microservices.io/patterns/server-side-discovery.html)。

![Server-side service discovery example](http://iximiuz.com/service-discovery-in-kubernetes/server-side-service-discovery-4000-opt.png)

_Server-side service discovery._

_服务器端服务发现。_

The load balancer abstracts the volatile set of service instances away from the clients. However, the load balancer itself needs to be aware of the up to date state of the service fleet. This can be achieved by adding a [_service registry_](https://microservices.io/patterns/service-registry.html) component. On the startup, an instance needs to be added to the registry database. Upon the termination, the instance needs to be removed from it. One of the main tasks of the load balancer is to dynamically update its routing rules based on the service registry information.

负载平衡器从客户端抽象出一组易失的服务实例。但是，负载均衡器本身需要了解服务队列的最新状态。这可以通过添加 [_service registry_](https://microservices.io/patterns/service-registry.html) 组件来实现。在启动时，需要将一个实例添加到注册表数据库中。终止后，需要将实例从中删除。负载均衡器的主要任务之一是根据服务注册信息动态更新其路由规则。

While it looks very appealing from the client standpoint, the server-side service discovery may quickly reveal its downsides, especially in highly-loaded environments. The load balancer component is a **single point of failure and a potential throughput bottleneck**. To overcome this, the load balancing layer needs to be designed with a reasonable level of redundancy. Additionally, the load balancer may **need to be aware of all the communication protocols** used between services and clients and there will be always an **extra network hop** on the request path.

虽然从客户端的角度来看它看起来很有吸引力，但服务器端服务发现可能很快就会暴露出它的缺点，尤其是在高负载环境中。负载均衡器组件是**单点故障和潜在的吞吐量瓶颈**。为了克服这个问题，负载均衡层需要设计有合理的冗余度。此外，负载均衡器可能**需要了解服务和客户端之间使用的所有通信协议**，并且请求路径上总会有**额外的网络跃点**。

## Client-Side Service Discovery 

## 客户端服务发现

Can we solve the service discovery problem without introducing the centralized load balancing component? Sure! If we keep the _service registry_ component around, we can teach the clients to look up the service instance addresses in the _service registry_ directly. After fetching the full list of IP addresses constituting the service fleet, a client could pick up an instance based on the load balancing strategy at its disposal. In such a case, [the service discovery would be happening solely on the client-side](https://microservices.io/patterns/client-side-discovery.html). Probably the most prominent real-world implementation of the client-side service discovery is Netflix [Eureka](https://github.com/Netflix/eureka) and [Ribbon](https://github.com/Netflix/ribbon) projects.

在不引入集中式负载均衡组件的情况下，能否解决服务发现问题？当然！如果我们保留 _service registry_ 组件，我们可以教客户端直接在 _service registry_ 中查找服务实例地址。在获取构成服务队列的完整 IP 地址列表后，客户端可以根据其处理的负载平衡策略选择一个实例。在这种情况下，[服务发现将仅发生在客户端](https://microservices.io/patterns/client-side-discovery.html)。客户端服务发现的最突出的现实世界实现可能是 Netflix [Eureka](https://github.com/Netflix/eureka) 和 [Ribbon](https://github.com/Netflix/ribbon)项目。

![Client-side service discovery example](http://iximiuz.com/service-discovery-in-kubernetes/client-side-service-discovery-4000-opt.png)

_Client-side service discovery._

_客户端服务发现。_

The benefits of the client-side approach mostly come from the absence of the load balancer. There is neither a **single point of failure** nor a potential **throughput bottleneck** in the system design. There is also one less moving part which is usually a good thing and **no extra network hops on the packet path**.

客户端方法的好处主要来自没有负载平衡器。系统设计中既不存在**单点故障**，也不存在潜在的**吞吐量瓶颈**。还有一个较少移动的部分，这通常是一件好事，**数据包路径上没有额外的网络跃点**。

However, as with the server-side service discovery, there are some significant drawbacks as well. Client-side service discovery **couples clients with the _service registry_**. It **requires some integration code** to be written for every programming language or framework in your ecosystem. And obviously, this **extra logic complicates the clients**. There seem to be an effort to [offload the client-side service discovery logic to the service proxy sidecars](http://iximiuz.com/en/posts/service-proxy-pod-sidecar-oh-my/) but that's already a different story...

但是，与服务器端服务发现一样，也存在一些明显的缺点。客户端服务发现**将客户端与 _service registry_** 结合起来。它**需要为生态系统中的每种编程语言或框架编写一些集成代码**。显然，这个额外的逻辑使客户端复杂化**。似乎有努力[将客户端服务发现逻辑卸载到服务代理边车](http://iximiuz.com/en/posts/service-proxy-pod-sidecar-oh-my/) 但就是这样已经是另一个故事了...

## DNS and Service Discovery

## DNS 和服务发现

Service discovery can be implemented in more than just two ways. There is a well-known [DNS-based service discovery (DNS-SD)](https://en.wikipedia.org/wiki/Zero-configuration_networking#DNS-based_service_discovery) approach that actually even predates the massive spread of microservices. Quoting the Wikipedia, _"a client discovers the list of available instances for a given service type by querying the DNS PTR record of that service type's name; the server returns zero or more names of the form `<Service>.<Domain>` , each corresponding to a SRV/TXT record pair. The SRV record resolves to the domain name providing the instance..."_ and then another DNS query can be used to resolve a chosen instance's domain name to the actual IP address. Wow, so many layers of indirection, sounds fun 🙈

服务发现可以通过不止两种方式实现。有一种众所周知的 [基于 DNS 的服务发现 (DNS-SD)](https://en.wikipedia.org/wiki/Zero-configuration_networking#DNS-based_service_discovery) 方法实际上甚至早于微服务的大规模传播。引用维基百科，_“客户端通过查询该服务类型名称的 DNS PTR 记录来发现给定服务类型的可用实例列表；服务器返回零个或多个形式为 `<Service>.<Domain>` 的名称，每个对应一个 SRV/TXT 记录对。SRV 记录解析为提供实例的域名..."_ 然后可以使用另一个 DNS 查询将所选实例的域名解析为实际 IP 地址。哇，这么多层的间接，听起来很有趣🙈

I never worked with DNS-SD, but to me, it doesn't sound like full-fledged service discovery. Rather, DNS is used as a _service registry_, and then depending on the dislocation of the code that knows how to query and interpret DNS-SD records, we can get either a [canonical client-side](https://github.com/benschw/srv-lb) or a [canonical server-side](https://www.haproxy.com/blog/dns-service-discovery-haproxy/) implementation.

我从未使用过 DNS-SD，但对我来说，这听起来不像是成熟的服务发现。相反，DNS 被用作_服务注册中心_，然后根据知道如何查询和解释 DNS-SD 记录的代码的错位，我们可以得到一个[规范客户端](https://github.com/benschw/srv-lb) 或 [规范服务器端](https://www.haproxy.com/blog/dns-service-discovery-haproxy/) 实现。

Alternatively, [Round-robin DNS](https://en.wikipedia.org/wiki/Round-robin_DNS) can be (ab)used for service discovery. Even though it was originally designed for load balancing, or rather load distribution, having multiple A records for the same hostname (i.e. service name) returned in a rotating manner implicitly abstracts multiple replicas behind a single service name.

或者，[Round-robin DNS](https://en.wikipedia.org/wiki/Round-robin_DNS)可以（ab）用于服务发现。尽管它最初是为负载平衡或负载分配而设计的，但以轮换方式返回的相同主机名（即服务名称)的多个 A 记录隐式地抽象了单个服务名称后面的多个副本。

In any case, **DNS has a significant drawback when used for service discovery**. Updating DNS records is a slow procedure. There are **multiple layers of caches**, including the client-side libraries, and historically record's TTL isn't strictly respected. As a result, propagating the change in the set of service instances to all the clients can take a while.

在任何情况下，**DNS 在用于服务发现时都有一个明显的缺点**。更新 DNS 记录是一个缓慢的过程。有**多层缓存**，包括客户端库，并且历史记录的 TTL 没有严格遵守。因此，将一组服务实例中的更改传播到所有客户端可能需要一段时间。

## Service Discovery in Kubernetes

## Kubernetes 中的服务发现

First, let's play an analogy game! 

首先，让我们玩一个类比游戏！

If I were to draw some analogies between Kubernetes and more traditional architectures, I'd compare Kubernetes [_Pods_](https://kubernetes.io/docs/concepts/workloads/pods/) with service instances. _Pods_ are [many things](https://www.reddit.com/r/kubernetes/comments/k0zc0m/kubernetes_pods_are_logical_hosts_or_simply/?utm_source=share&utm_medium=web2x&context=3) to [many people](https://twitter.com/iximiuz/status/1331630336707596288?s=20), however, when it comes to networking [the documentation clearly states that](https://kubernetes.io/docs/concepts/cluster-administration/networking/#the-kubernetes-network-model) _"...Pods can be treated much like VMs or physical hosts from the perspectives of port allocation, naming, service discovery, load balancing, application configuration, and migration."_

如果我要在 Kubernetes 和更传统的架构之间进行一些类比，我会将 Kubernetes [_Pods_](https://kubernetes.io/docs/concepts/workloads/pods/) 与服务实例进行比较。 _Pods_ 是 [许多东西](https://www.reddit.com/r/kubernetes/comments/k0zc0m/kubernetes_pods_are_logical_hosts_or_simply/?utm_source=share&utm_medium=web2x&context=3) 对 [许多人](https://twitter.com/iximiuz/status/1331630336707596288?s=20)，然而，当涉及到网络时[文档明确指出](https://kubernetes.io/docs/concepts/cluster-administration/networking/#the-kubernetes-network-model) _"...从端口分配、命名、服务发现、负载均衡、应用程序配置和迁移的角度来看，Pods 可以像虚拟机或物理主机一样对待。"_

If _Pods_ correspond to individual instances of a service, I'd expect a similar analogy for the _service_, as a logical grouping of instances, itself. And indeed there is a suitable concept in Kubernetes called... surprise, surprise, a [_Service_](https://kubernetes.io/docs/concepts/services-networking/service/). _" In Kubernetes, a Service is an abstraction which defines a logical set of Pods and a policy by which to access them (sometimes this pattern is called a micro-service)."_

如果 _Pods_ 对应于服务的单个实例，我希望 _service_ 有类似的类比，作为实例的逻辑分组本身。事实上，Kubernetes 中有一个合适的概念叫做...惊喜，惊喜，一个 [_Service_](https://kubernetes.io/docs/concepts/services-networking/service/)。 _“在 Kubernetes 中，服务是一种抽象，它定义了一组逻辑 Pod 和访问它们的策略（有时这种模式称为微服务)。”_

Strengthening the analogy, the set of Pods making up a Service should be also considered ephemeral because neither the Pods' headcount nor the final set of IP addresses is stable. Thus, in Kubernetes, the problem of providing a reliable **service discovery** remains actual.

加强类比，组成 Service 的一组 Pod 也应该被视为短暂的，因为 Pod 的人数和最终的 IP 地址集都不是稳定的。因此，在 Kubernetes 中，提供可靠的**服务发现**的问题仍然存在。

While creating a new Service, one should choose a name that will be used to refer to the set of Pods constituting the service. Among other things, the Service maintains an up to date list of IP addresses of its Pods organized as an [_Endpoints_](https://kubernetes.io/docs/reference/glossary/?all=true#term-endpoints) (or [_EndpontSlice_](https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/) since Kubernetes 1.17) object. [Citing the documentation](https://kubernetes.io/docs/concepts/services-networking/service/#cloud-native-service-discovery) one more time, _"...if you're able to use Kubernetes APIs for service discovery in your application, you can query the API server for Endpoints,_ [I guess using the Service name] _that get updated whenever the set of Pods in a Service changes."_ Well, sounds like an invitation to implement a cloud-native Kubernetes-native client-side service discovery with the Kubernetes control plane playing (in particular) the role of the _service registry_.

在创建新服务时，应该选择一个名称，用于指代构成服务的一组 Pod。除其他外，该服务维护其 Pod 的最新 IP 地址列表，该列表组织为 [_Endpoints_](https://kubernetes.io/docs/reference/glossary/?all=true#term-endpoints)（或[_EndpontSlice_](https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/) 自 Kubernetes 1.17) 对象。 [引用文档](https://kubernetes.io/docs/concepts/services-networking/service/#cloud-native-service-discovery) 再来一次，_"...如果你能够使用 Kubernetes在您的应用程序中用于服务发现的 API，您可以查询 API 服务器以获取端点，_ [我猜使用服务名称] _每当服务中的一组 Pod 发生变化时都会更新。”_ 好吧，听起来像是实现一个云原生 Kubernetes 原生客户端服务发现，Kubernetes 控制平面扮演（特别是)_service registry_ 的角色。

![Kubernetes-native client-side service discovery](http://iximiuz.com/service-discovery-in-kubernetes/kube-native-service-discovery-4000-opt.png)

_Kubernetes-native client-side service discovery._

_Kubernetes 原生客户端服务发现。_

However, the only real-world usage of this mechanism I've stumbled upon so far was in the [_service mesh_](https://linkerd.io/2020/07/23/under-the-hood-of-linkerds-state-of-the-art-rust-proxy-linkerd2-proxy/#the-life-of-a-request) kind of software. It's a bit unfair to mention it here though because _service mesh_ itself is needed to provide, in particular, the service discovery mechanism for its users. So, if you're aware of the client-side service discovery implementations leveraging Kubernetes Endpoints API please drop a comment below.

但是，到目前为止，我偶然发现的这种机制的唯一实际用法是在 [_service mesh_](https://linkerd.io/2020/07/23/under-the-hood-of-linkerds-state-of-the-art-rust-proxy-linkerd2-proxy/#the-life-of-a-request) 类型的软件。在这里提及它有点不公平，因为需要 _service mesh_ 本身来为其用户提供服务发现机制。因此，如果您了解利用 Kubernetes Endpoints API 的客户端服务发现实现，请在下面发表评论。

Luckily, as with many other things in Kubernetes, there's more than one way to skin a cat to get the service discovery done. And the applications that weren't born cloud-native (i.e. 99% of them) will likely find the next service-discovery mechanism much more appealing.

幸运的是，与 Kubernetes 中的许多其他东西一样，有不止一种方法可以给猫剥皮来完成服务发现。而那些不是原生云的应用程序（即 99%）可能会发现下一个服务发现机制更具吸引力。

## Network-Side Service Discovery

## 网络端服务发现

_Disclaimer: I've no idea if there is such thing as network-side service discovery in other domains and I've never seen the usage of this term in the microservices world. But I find it funny and suitable for the purpose of this paragraph._ 

_免责声明：我不知道其他领域是否有网络端服务发现这样的东西，而且我从未在微服务世界中看到过这个术语的用法。但我觉得这很有趣，适合本段的目的。_

In Kubernetes, the name of a Service object must be a valid [DNS label name](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names). It's not a coincidence. When the DNS add-on is enabled ( [and I guess it's almost always the case](https://kubernetes.io/docs/concepts/services-networking/service/#dns)), every Service gets a DNS record like `<service-name>.<namespace-name>`. Obviously, this name can be used by applications to access the service and it simplifies the life of the clients to the highest possible extent. A single well-known address behind every Service **eradicates the need for any service discovery logic on the client-side**.

在 Kubernetes 中，Service 对象的名称必须是有效的 [DNS 标签名称](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names)。这不是巧合。当启用 DNS 附加组件时（[我猜它几乎总是这样](https://kubernetes.io/docs/concepts/services-networking/service/#dns))，每个服务都会获得一个 DNS 记录，例如`<服务名称>.<命名空间名称>`。显然，应用程序可以使用该名称来访问服务，并且最大限度地简化了客户端的生命周期。每个服务背后的一个众所周知的地址**消除了客户端对任何服务发现逻辑的需求**。

However, as we already know, DNS is often ill-suited for service discovery and [the Kubernetes ecosystem is not an exception](https://kubernetes.io/docs/concepts/services-networking/service/#why-not-use-round-robin-dns). Therefore, instead of using round-robin DNS to list Pods' IP addresses, Kubernetes introduces one more IP address for every service. This IP address is called `clusterIP` (not to be confused with the [`ClusterIP` service type](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types)). Similar to the DNS name, this address can be used to transparently access Pods constituting the service.

然而，正如我们已经知道的，DNS 通常不适合服务发现，[Kubernetes 生态系统也不例外](https://kubernetes.io/docs/concepts/services-networking/service/#why-not-使用round-robin-dns）。因此，Kubernetes 不是使用循环 DNS 来列出 Pod 的 IP 地址，而是为每个服务引入了一个 IP 地址。这个IP地址称为`clusterIP`（不要与[`ClusterIP`服务类型](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types）混淆))。与 DNS 名称类似，该地址可用于透明访问构成服务的 Pod。

_NB: there is actually no hard dependency on DNS for Kubernetes applications. Clients can always learn a `clusterIP` of a service by inspecting their [environment variables](https://kubernetes.io/docs/concepts/services-networking/service/#environment-variables). Upon a pod startup **for every running service** Kubernetes injects a couple of env variables looking like `<service-name>_SERVICE_HOST` and `<service-name>_SERVICE_PORT`._

_NB：Kubernetes 应用程序实际上没有对 DNS 的硬依赖。客户端总是可以通过检查他们的 [环境变量]来了解服务的 `clusterIP`。在 pod 启动时**对于每个正在运行的服务** Kubernetes 注入几个环境变量，看起来像 `<service-name>_SERVICE_HOST` 和 `<service-name>_SERVICE_PORT`。_

Ok, here is one more analogy. The resulting (logical) setup looks much like a load balancer or reverse proxy sitting in front of the set of virtual machines.

好的，这里还有一个类比。结果（逻辑）设置看起来很像位于一组虚拟机前面的负载平衡器或反向代理。

![Server-side (logically) service discovery in Kubernetes](http://iximiuz.com/service-discovery-in-kubernetes/kube-logical-service-discovery-4000-opt.png)

_Server-side (logically) service discovery in Kubernetes._

_Kubernetes 中的服务器端（逻辑上）服务发现。_

But there is more to it than just that. This `clusterIP` is a so-called virtual address. When I stumbled upon the concept of the virtual IP for the first time it was a real mind-bender.

但还有更多的东西。这个`clusterIP`就是所谓的虚拟地址。当我第一次偶然发现虚拟 IP 的概念时，我真的很困惑。

A _virtual IP_ basically means that there is no single network interface in the whole system carrying it around! Instead, there is a super-powerful and likely underestimated background component called [_kube-proxy_](https://kubernetes.io/docs/concepts/overview/components/#kube-proxy) that magically makes all the Pods (and even Nodes) [thinking the Service IPs do exist](https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies): _"Every node in a Kubernetes cluster runs a `kube-proxy`. `kube-proxy` is responsible for implementing a form of virtual IP for Services of type other than ExternalName."_ 

_虚拟 IP_ 基本上意味着整个系统中没有单个网络接口承载它！相反，有一个名为 [_kube-proxy_](https://kubernetes.io/docs/concepts/overview/components/#kube-proxy) 的超级强大且可能被低估的后台组件，它神奇地使所有 Pod（甚至Nodes) [认为 Service IP 确实存在](https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies)：_"Kubernetes 集群中的每个节点都运行一个 `kube-proxy`。`kube-proxy` 负责为非 ExternalName 类型的服务实现一种形式的虚拟 IP。”_

Funnily enough, the _kube-proxy_ component is actually a [misnomer](https://en.wikipedia.org/wiki/Misnomer). I.e. it's not really a proxy anymore, although [it was born as a true user space proxy](https://github.com/kubernetes/kubernetes/issues/1107). I'm not going to dive into implementation details here, there is plenty of information on the Internet including the official Kubernetes documentation and [this great article](https://arthurchiao.art/blog/cracking-k8s-node-proxy/) of Arthur Chiao. Long story short - _kube-proxy_ **operates on the network layer** using such Linux capabilities as iptables or IPVS and transparently substitutes the destination `clusterIP` with an IP address of some service's Pod. Thus, **_kube-proxy_ is one of the main implementers of the service discovery and load balancing in the cluster**.

有趣的是，_kube-proxy_ 组件实际上是一个 [misnomer](https://en.wikipedia.org/wiki/Misnomer)。 IE。它不再是真正的代理，尽管 [它是作为真正的用户空间代理诞生的](https://github.com/kubernetes/kubernetes/issues/1107)。我不打算在这里深入研究实现细节，互联网上有很多信息，包括官方 Kubernetes 文档和 [这篇很棒的文章](https://arthurchio.art/blog/cracking-k8s-node-proxy/) 亚瑟·乔。长话短说 - _kube-proxy_**在网络层**运行，使用诸如 iptables 或 IPVS 之类的 Linux 功能，并透明地将目标 `clusterIP` 替换为某些服务的 Pod 的 IP 地址。因此，**_kube-proxy_ 是集群中服务发现和负载均衡的主要实现者之一**。

![Kube-proxy implements service discovery in Kubernetes](http://iximiuz.com/service-discovery-in-kubernetes/kube-proxy-service-discovery-4000-opt.png)

_Kube-proxy implements service discovery in Kubernetes._

_Kube-proxy 在 Kubernetes 中实现服务发现。_

The _kube-proxy_ component turns every Kubernetes node into a service proxy (just another fancy name for a client-side proxy) and all pod-to-pod traffic always goes through its local service proxy.

_kube-proxy_ 组件将每个 Kubernetes 节点变成一个服务代理（客户端代理的另一个花哨名称），并且所有 pod 到 pod 的流量总是通过其本地服务代理。

Now, let's see what does it mean from the service discovery standpoint. Since there are as many self-sufficient proxies as the number of nodes in the cluster, there is **no single point of failure** when it comes to load balancing. Unlike the canonical server-side service discovery technique with the centralized load balancer component, _kube-proxy_-based service discovery follows the decentralized approach with all the nodes sharing a comparable amount of traffic. Hence, **the probability of getting a throughput bottleneck is also much lower.** On top of that, there is **no extra network hop** on the packet's path because every Pod contacts its node-local copy of proxy.

现在，让我们看看从服务发现的角度来看这意味着什么。由于自给自足的代理数量与集群中的节点数量一样多，因此在负载均衡方面**没有单点故障**。与具有集中式负载均衡器组件的规范服务器端服务发现技术不同，基于 _kube-proxy_ 的服务发现遵循分散式方法，所有节点共享相当数量的流量。因此，**出现吞吐量瓶颈的可能性也低得多。** 最重要的是，在数据包的路径上**没有额外的网络跃点**，因为每个 Pod 都联系其节点本地的代理副本。

**Thereby, Kubernetes takes the best of both worlds. As with the server-side service discovery, clients can simply access a single endpoint, a stable Service IP address, i.e. there is no need for advanced logic on the application side. At the same time, physically the service discovery and load balancing happen on every cluster node, i.e. close to the client-side. Thus, there are no traditional downsides of the server-side service discovery.**

**因此，Kubernetes 两全其美。与服务器端服务发现一样，客户端可以简单地访问单个端点、一个稳定的服务 IP 地址，即不需要应用程序端的高级逻辑。同时，服务发现和负载均衡在物理上发生在每个集群节点上，即靠近客户端。因此，服务器端服务发现没有传统的缺点。**

Since the implementation of the service discovery in Kubernetes heavily relies on the Linux network stack, I'm inclined to call it a _network-side_ service discovery. Although, the term _service-side_ service discovery might work as well.

由于 Kubernetes 中服务发现的实现严重依赖 Linux 网络堆栈，因此我倾向于将其称为 _network-side_ 服务发现。虽然，术语 _service-side_ 服务发现也可能适用。

## Conclusion

##  结论

Kubernetes tries hard to make the transition from more traditional virtual or bare-metal ecosystems to containers simple. Kubernetes NAT-less networking model, Pods, and Services allow familiar designs to be reapplied without significant adjustments. But Kubernetes goes even further and provides [a very reliable and elegant solution for the in-cluster service discovery and load balancing problems](https://kubernetespodcast.com/episode/129-linkerd/) out of the box. On top of that, the provided solution turned out to be easy to extend and that gave birth to such an amazing piece of software as a [Kubernetes-native service mesh](https://github.com/linkerd/linkerd2).

Kubernetes 努力使从更传统的虚拟或裸机生态系统向容器的过渡变得简单。 Kubernetes 无 NAT 网络模型、Pod 和服务允许重新应用熟悉的设计，而无需进行重大调整。但是 Kubernetes 走得更远，提供了 [一个非常可靠和优雅的集群内服务发现和负载平衡问题的解决方案](https://kubernetespodcast.com/episode/129-linkerd/) 开箱即用。最重要的是，所提供的解决方案被证明很容易扩展，并催生了这样一个惊人的软件，如 [Kubernetes 原生服务网格](https://github.com/linkerd/linkerd2)。

_Disclaimer: This article intentionally omits the questions of external service (Service type `ExternalName`) discovering and discovering of the Kubernetes services from the outside world (Ingress Controller). These two deserve a dedicated article each._

_免责声明：本文有意省略了外部服务（Service type `ExternalName`）发现和从外界（Ingress Controller）发现Kubernetes服务的问题。这两个人都应该有一篇专门的文章。_

### Further reading

### 进一步阅读

- [Pattern: Server-side service discovery](https://microservices.io/patterns/server-side-discovery.html)
- [Pattern: Client-side service discovery](https://microservices.io/patterns/client-side-discovery.html)
- [Pattern: Service registry](https://microservices.io/patterns/service-registry.html)
- [Service Discovery in a Microservices Architecture](https://www.nginx.com/blog/service-discovery-in-a-microservices-architecture/)
- [Microservices: Client Side Load Balancing](https://www.linkedin.com/pulse/microservices-client-side-load-balancing-amit-kumar-sharma/)
- [Kubernetes Podcast from Google: Linkerd, with Thomas Rampelberg](https://kubernetespodcast.com/episode/129-linkerd/) 

- [模式：服务端服务发现](https://microservices.io/patterns/server-side-discovery.html)
- [模式：客户端服务发现](https://microservices.io/patterns/client-side-discovery.html)
- [模式：服务注册](https://microservices.io/patterns/service-registry.html)
- [微服务架构中的服务发现](https://www.nginx.com/blog/service-discovery-in-a-microservices-architecture/)
- [微服务：客户端负载均衡](https://www.linkedin.com/pulse/microservices-client-side-load-balancing-amit-kumar-sharma/)
- [来自 Google 的 Kubernetes 播客：Linkerd，与 Thomas Rampelberg](https://kubernetespodcast.com/episode/129-linkerd/)

- [Baker Street: Avoiding Bottlenecks with a Client-Side Load Balancer for Microservices](https://thenewstack.io/baker-street-avoiding-bottlenecks-with-a-client-side-load-balancer-for-microservices/)

- [贝克街：使用微服务的客户端负载均衡器避免瓶颈](https://thenewstack.io/baker-street-avoiding-bottlenecks-with-a-client-side-load-balancer-for-microservices/)

### Other posts you may like

### 您可能喜欢的其他帖子

- [Service proxy, pod, sidecar, oh my!](http://iximiuz.com/en/posts/service-proxy-pod-sidecar-oh-my/)
- [Exploring Kubernetes Operator Pattern](http://iximiuz.com/en/posts/kubernetes-operator-pattern/)

- [服务代理, pod, sidecar, oh my!](http://iximiuz.com/en/posts/service-proxy-pod-sidecar-oh-my/)
- [探索Kubernetes Operator模式](http://iximiuz.com/en/posts/kubernetes-operator-pattern/)

[kubernetes,](javascript: void 0) [service-discovery](javascript: void 0)

[kubernetes,](javascript: void 0) [service-discovery](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

