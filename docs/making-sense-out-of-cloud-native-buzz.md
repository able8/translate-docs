# Making sense out of cloud-native buzz

December 10, 2020 (Updated: August 14, 2021)

[Ranting](http://iximiuz.com/en/categories/?category=Ranting)

I've been trying to wrap my head around the tremendous growth of the [cloud-native zoo](https://landscape.cncf.io/) for quite some time. But recently I was listening to [a wonderful podcast episode](https://kubernetespodcast.com/episode/129-linkerd/) with the Linkerd creator [Thomas Rampelberg](https://twitter.com/grampelberg) and he kindly reminded me one thing about... **microservices**. Long story short, despite the common belief that microservices solve technical problems, the most appealing part of the microservice architecture apparently has something to do with **solving organisational problems** such as allocating teams to development areas or tackling software modernisation campaigns. And on the contrary, while helping with the org problems, **microservices rather create new technical challenges**!

_Disclaimer: This article is not about Microservices vs Monolith._

That made me rethink the need for all those projects constituting the cloud-native landscape. From now on I can't help but see an awful load of projects solving all kinds of technical problems originated by the transition to the microservice paradigm.

![](http://iximiuz.com/making-sense-out-of-cloud-native-buzz/kdpv-2000-opt.png)

Starting from the foundational project, "The Production-Grade Container Orchestration" system we all know and love - [Kubernetes](https://landscape.cncf.io/?selected=kubernetes). While it's not limited by that, the idea of [grouping containers Pods into Services](https://kubernetes.io/docs/concepts/services-networking/service/) is foundational for Kubernetes. But there is even more, Kubernetes provides other higher-level capabilities vital for microservice architectures such as [service discovery](http://iximiuz.com/en/posts/service-discovery-in-kubernetes/#entering-kubernetes-realm) or autoscaling out of the box. But if you cannot benefit from Kubernetes goodies because of... reasons, fear not! CNCF has a decent list of [Coordination & Service Discovery](https://landscape.cncf.io/card-mode?category=coordination-service-discovery&grouping=category) offerings (ZooKeeper, etcd, Netflix Eureka, etc) that can teach an old-fashioned infrastructure some new tricks.

Going further, the increase in the number of services naturally leads to the surge of network communications. Hence, we need to have a reliable remote procedure call layer. Luckily, the CNCF landscape has an [RPC section](https://landscape.cncf.io/card-mode?category=remote-procedure-call&grouping=category) for you with such prominent projects as gRPC or Apache Thrift.

In microservice-verse, you learn pretty quickly that an uncontrolled service-to-service communication leads to a mess... that can be eliminated with a [mesh](https://servicemesh.io/). Projects like Linkerd, Maesh, or Istio are great representatives of the [Service Mesh category](https://landscape.cncf.io/card-mode?category=service-mesh&grouping=category). And of course, under the hood of almost every service mesh, there is a piece of black networking magic proxy like Envoy, Traefik, or linkerd2-proxy. Well, apparently there is also a [Service Proxy category](https://landscape.cncf.io/card-mode?category=service-proxy&grouping=category) gathering such proxies.

Having transparent service-to-service communication doesn't make every network request legitimate. In the microservice architecture, you cannot trust every service. Only requests from well-identified clients should be allowed. And CNCF landscape has something for you again - [spiffe](https://landscape.cncf.io/card-mode?category=key-management&grouping=category&selected=spiffe) and [spire](https://landscape.cncf.io/card-mode?category=key-management&grouping=category&selected=spire) projects to provide the distributed identity.

Lots of services mean lots of APIs. These APIs need to be consolidated before being exposed to the outside world. Guess what? There is a whole [API Gateway category](https://landscape.cncf.io/card-mode?category=api-gateway&grouping=category) at your disposal.

Chasing bugs in microservice architectures is much trickier than in monoliths. A request may hit tens of services on its way to be served. Tracing request paths, collecting and aggregating logs from all the services are great engineering problems deserving their own categories - [Tracing](https://landscape.cncf.io/card-mode?category=tracing&grouping=category) and [Logging](https://landscape.cncf.io/card-mode?category=logging&grouping=category). And of course, there is the third pillar of observability - [Metrics and Monitoring](https://landscape.cncf.io/card-mode?category=monitoring&grouping=category).

_Containers vs microservices_ always sounded to me like a chicken and egg problem. Anyway, it's clear now that microservices and containers work well together. Thus, we can consider the sheer amount of container-related projects like [runtimes](https://landscape.cncf.io/card-mode?category=container-runtime&grouping=category) or [container registries](https://landscape.cncf.io/card-mode?category=container-registry&grouping=category) as microservice-supporting as well.

I could continue giving more and more examples here, but I hope that by this time I already managed to make my point. Sometimes, the correlation is more obvious, and sometimes there is some level of indirection but to me, it looks like the vast majority of the projects under the CNCF umbrella exist to solve the technical problems that arose only because we started using lots of smaller services in our web applications.

For curious, I recommend the following series of articles diving into the CNCF landscape, this time without any poor speculations 🙈:

- [An Introduction to the Cloud Native Landscape](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/)
- [The Cloud Native Landscape: The Provisioning Layer Explained](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)
- [The Cloud Native Landscape: The Runtime Layer Explained](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)
- [The Cloud Native Landscape: The Orchestration and Management Layer](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/)

[microservices,](javascript: void 0) [cloud-native,](javascript: void 0) [cncf](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

