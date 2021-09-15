# Making sense out of cloud-native buzz

# 理解云原生的嗡嗡声

December 10, 2020 (Updated: August 14, 2021)

[Ranting](http://iximiuz.com/en/categories/?category=Ranting)

I've been trying to wrap my head around the tremendous growth of the [cloud-native zoo](https://landscape.cncf.io/) for quite some time. But recently I was listening to [a wonderful podcast episode](https://kubernetespodcast.com/episode/129-linkerd/) with the Linkerd creator [Thomas Rampelberg](https://twitter.com/grampelberg) and he kindly reminded me one thing about... **microservices**. Long story short, despite the common belief that microservices solve technical problems, the most appealing part of the microservice architecture apparently has something to do with **solving organisational problems** such as allocating teams to development areas or tackling software modernisation campaigns. And on the contrary, while helping with the org problems, **microservices rather create new technical challenges**!

一段时间以来，我一直试图围绕 [云原生动物园](https://landscape.cncf.io/) 的巨大增长进行思考。但最近我和 Linkerd 的创造者 [Thomas Rampelberg](https://twitter.com/grampelberg) 一起听了 [一个精彩的播客集](https://kubernetespodcast.com/episode/129-linkerd/)，他很友好让我想起一件事……**微服务**。长话短说，尽管普遍认为微服务可以解决技术问题，但微服务架构中最吸引人的部分显然与**解决组织问题**有关，例如将团队分配到开发领域或解决软件现代化活动。相反，在帮助解决组织问题的同时，**微服务反而创造了新的技术挑战**！

_Disclaimer: This article is not about Microservices vs Monolith._

_免责声明：本文不是关于微服务与 Monolith。_

That made me rethink the need for all those projects constituting the cloud-native landscape. From now on I can't help but see an awful load of projects solving all kinds of technical problems originated by the transition to the microservice paradigm.

这让我重新思考构成云原生景观的所有项目的必要性。从现在开始，我不禁看到大量项目解决了由过渡到微服务范式而产生的各种技术问题。

![](http://iximiuz.com/making-sense-out-of-cloud-native-buzz/kdpv-2000-opt.png)

Starting from the foundational project, "The Production-Grade Container Orchestration" system we all know and love - [Kubernetes](https://landscape.cncf.io/?selected=kubernetes). While it's not limited by that, the idea of [grouping containers Pods into Services](https://kubernetes.io/docs/concepts/services-networking/service/) is foundational for Kubernetes. But there is even more, Kubernetes provides other higher-level capabilities vital for microservice architectures such as [service discovery](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/#entering-kubernetes-realm) or autoscaling out of the box. But if you cannot benefit from Kubernetes goodies because of... reasons, fear not! CNCF has a decent list of [Coordination & Service Discovery](https://landscape.cncf.io/card-mode?category=coordination-service-discovery&grouping=category) offerings (ZooKeeper, etcd, Netflix Eureka, etc) that can teach an old-fashioned infrastructure some new tricks.

从基础项目开始，我们都知道和喜爱的“生产级容器编排”系统——[Kubernetes](https://landscape.cncf.io/?selected=kubernetes)。虽然不受此限制，但 [将容器 Pod 分组为服务](https://kubernetes.io/docs/concepts/services-networking/service/) 的想法是 Kubernetes 的基础。但更重要的是，Kubernetes 提供了其他对微服务架构至关重要的更高级别的功能，例如 [服务发现](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/#entering-kubernetes-realm) 或开箱即用的自动缩放。但是，如果您因为……原因而无法从 Kubernetes 的好处中受益，请不要害怕！ CNCF 有一份不错的 [Coordination & Service Discovery](https://landscape.cncf.io/card-mode?category=coordination-service-discovery&grouping=category) 产品列表（ZooKeeper、etcd、Netflix Eureka 等)，可以教一个老式的基础设施一些新的技巧。

Going further, the increase in the number of services naturally leads to the surge of network communications. Hence, we need to have a reliable remote procedure call layer. Luckily, the CNCF landscape has an [RPC section](https://landscape.cncf.io/card-mode?category=remote-procedure-call&grouping=category) for you with such prominent projects as gRPC or Apache Thrift.

更进一步，服务数量的增加自然会导致网络通信的激增。因此，我们需要有一个可靠的远程过程调用层。幸运的是，CNCF 领域有一个 [RPC 部分](https://landscape.cncf.io/card-mode?category=remote-procedure-call&grouping=category) 为您提供 gRPC 或 Apache Thrift 等突出项目。

In microservice-verse, you learn pretty quickly that an uncontrolled service-to-service communication leads to a mess... that can be eliminated with a [mesh](https://servicemesh.io/). Projects like Linkerd, Maesh, or Istio are great representatives of the [Service Mesh category](https://landscape.cncf.io/card-mode?category=service-mesh&grouping=category). And of course, under the hood of almost every service mesh, there is a piece of black networking magic proxy like Envoy, Traefik, or linkerd2-proxy. Well, apparently there is also a [Service Proxy category](https://landscape.cncf.io/card-mode?category=service-proxy&grouping=category) gathering such proxies.

在 microservice-verse 中，您很快就会了解到不受控制的服务到服务通信会导致混乱……这可以通过 [mesh](https://servicemesh.io/) 消除。 Linkerd、Maesh 或 Istio 等项目是 [Service Mesh 类别](https://landscape.cncf.io/card-mode?category=service-mesh&grouping=category) 的杰出代表。当然，在几乎每个服务网格的幕后，都有一个黑色网络代理，比如 Envoy、Traefik 或 linkerd2-proxy。嗯，显然还有一个[服务代理类别](https://landscape.cncf.io/card-mode?category=service-proxy&grouping=category)收集这样的代理。

Having transparent service-to-service communication doesn't make every network request legitimate. In the microservice architecture, you cannot trust every service. Only requests from well-identified clients should be allowed. And CNCF landscape has something for you again - [spiffe](https://landscape.cncf.io/card-mode?category=key-management&grouping=category&selected=spiffe) and [spire](https://landscape.cncf.io/card-mode?category=key-management&grouping=category&selected=spire) projects to provide the distributed identity.

具有透明的服务到服务通信并不能使每个网络请求都合法。在微服务架构中，你不能信任每一个服务。只应允许来自身份明确的客户的请求。 CNCF 景观再次为您准备了一些东西 - [spiffe](https://landscape.cncf.io/card-mode?category=key-management&grouping=category&selected=spiffe) 和 [spire](https://landscape.cncf.io/card-mode?category=key-management&grouping=category&selected=spire) 项目来提供分布式身份。

Lots of services mean lots of APIs. These APIs need to be consolidated before being exposed to the outside world. Guess what? There is a whole [API Gateway category](https://landscape.cncf.io/card-mode?category=api-gateway&grouping=category) at your disposal. 

大量的服务意味着大量的 API。这些 API 在暴露给外界之前需要进行整合。你猜怎么着？有一个完整的 [API 网关类别](https://landscape.cncf.io/card-mode?category=api-gateway&grouping=category) 供您使用。

Chasing bugs in microservice architectures is much trickier than in monoliths. A request may hit tens of services on its way to be served. Tracing request paths, collecting and aggregating logs from all the services are great engineering problems deserving their own categories - [Tracing](https://landscape.cncf.io/card-mode?category=tracing&grouping=category) and [Logging]( https://landscape.cncf.io/card-mode?category=logging&grouping=category). And of course, there is the third pillar of observability - [Metrics and Monitoring](https://landscape.cncf.io/card-mode?category=monitoring&grouping=category).

在微服务架构中追查 bug 比在单体应用中要棘手得多。一个请求在被服务的过程中可能会遇到数十个服务。跟踪请求路径、收集和聚合来自所有服务的日志是伟大的工程问题，值得拥有自己的类别 - [Tracing](https://landscape.cncf.io/card-mode?category=tracing&grouping=category) 和 [Logging]( https://landscape.cncf.io/card-mode?category=logging&grouping=category)。当然，还有第三个可观察性支柱——[指标和监控](https://landscape.cncf.io/card-mode?category=monitoring&grouping=category)。

_Containers vs microservices_ always sounded to me like a chicken and egg problem. Anyway, it's clear now that microservices and containers work well together. Thus, we can consider the sheer amount of container-related projects like [runtimes](https://landscape.cncf.io/card-mode?category=container-runtime&grouping=category) or [container registries](https://landscape.cncf.io/card-mode?category=container-registry&grouping=category) as microservice-supporting as well.

_容器 vs 微服务_在我看来总是像鸡和蛋的问题。无论如何，现在很明显微服务和容器可以很好地协同工作。因此，我们可以考虑大量与容器相关的项目，例如 [runtimes](https://landscape.cncf.io/card-mode?category=container-runtime&grouping=category) 或 [container registries](https://Landscape.cncf.io/card-mode?category=container-registry&grouping=category)也支持微服务。

I could continue giving more and more examples here, but I hope that by this time I already managed to make my point. Sometimes, the correlation is more obvious, and sometimes there is some level of indirection but to me, it looks like the vast majority of the projects under the CNCF umbrella exist to solve the technical problems that arose only because we started using lots of smaller services in our web applications.

我可以继续在这里举出越来越多的例子，但我希望此时我已经设法表达了我的观点。有时，相关性更明显，有时存在某种程度的间接性，但在我看来，CNCF 保护伞下的绝大多数项目都是为了解决技术问题而存在的，这些问题只是因为我们开始使用许多较小的服务在我们的网络应用程序中。

For curious, I recommend the following series of articles diving into the CNCF landscape, this time without any poor speculations 🙈:

出于好奇，我推荐以下系列文章深入了解 CNCF 领域，这一次没有任何糟糕的猜测🙈：

- [An Introduction to the Cloud Native Landscape](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/)
- [The Cloud Native Landscape: The Provisioning Layer Explained](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)
- [The Cloud Native Landscape: The Runtime Layer Explained](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)
- [The Cloud Native Landscape: The Orchestration and Management Layer](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/)

- [云原生景观简介](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/)
- [云原生景观：供应层解释](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)
- [云原生景观：运行时层解释](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)
- [云原生景观：编排和管理层](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/)

[microservices,](javascript: void 0) [cloud-native,](javascript: void 0) [cncf](javascript: void 0)

[微服务,](javascript: void 0) [云原生,](javascript: void 0) [cncf](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

