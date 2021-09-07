# Introducing istiod: simplifying the control plane

# 介绍 istiod：简化控制平面

Istiod consolidates the Istio control plane components into a single binary.

Istiod 将 Istio 控制平面组件整合到一个二进制文件中。

Mar 19, 2020 \| By Craig Box - Google

Microservices are a great pattern when they map services to disparate teams that deliver them, or when the value of independent rollout and the value of independent scale are greater than the cost of orchestration. We regularly talk to customers and teams running Istio in the real world, and they told us that none of these were the case for the Istio control plane. So, in Istio 1.5, we’ve changed how Istio is packaged, consolidating the control plane functionality into a single binary called **istiod**.

当微服务将服务映射到交付它们的不同团队时，或者当独立推出的价值和独立规模的价值大于编排成本时，微服务是一种很好的模式。我们定期与在现实世界中运行 Istio 的客户和团队交谈，他们告诉我们，这些都不是 Istio 控制平面的情况。因此，在 Istio 1.5 中，我们改变了 Istio 的打包方式，将控制平面功能整合到一个名为 **istiod** 的二进制文件中。

## History of the Istio control plane

## Istio 控制平面的历史

Istio implements a pattern that has been in use at both Google and IBM for many years, which later became known as “service mesh”. By pairing client and server processes with proxy servers, they act as an application-aware _data plane_ that’s not simply moving packets around hosts, or pulses over wires.

Istio 实现了一种已在 Google 和 IBM 使用多年的模式，后来被称为“服务网格”。通过将客户端和服务器进程与代理服务器配对，它们充当应用程序感知的_数据平面_，而不是简单地在主机周围移动数据包或通过电线传输脉冲。

This pattern helps the world come to terms with _microservices_: fine-grained, loosely-coupled services connected via lightweight protocols. The common cross-platform and cross-language standards like HTTP and gRPC that replace proprietary transports, and the widespread presence of the needed libraries, empower different teams to write different parts of an overall architecture in whatever language makes the most sense. Furthermore, each service can scale independently as needed. A desire to implement security, observability and traffic control for such a network powers Istio’s popularity.

这种模式帮助世界接受_微服务_：通过轻量级协议连接的细粒度、松散耦合的服务。通用的跨平台和跨语言标准（如 HTTP 和 gRPC）取代了专有传输，以及所需库的广泛存在，使不同的团队能够以最有意义的任何语言编写整体架构的不同部分。此外，每个服务都可以根据需要独立扩展。为这样的网络实现安全性、可观察性和流量控制的愿望推动了 Istio 的流行。

Istio’s _control plane_ is, itself, a modern, cloud-native application. Thus, it was built from the start as a set of microservices. Individual Istio components like service discovery (Pilot), configuration (Galley), certificate generation (Citadel) and extensibility (Mixer) were all written and deployed as separate microservices. The need for these components to communicate securely and be observable, provided opportunities for Istio to eat its own dogfood (or “drink its own champagne”, to use a more French version of the metaphor!).

Istio 的 _control plane_ 本身就是一个现代的云原生应用程序。因此，它从一开始就是作为一组微服务构建的。服务发现 (Pilot)、配置 (Galley)、证书生成 (Citadel) 和可扩展性 (Mixer) 等各个 Istio 组件都是作为单独的微服务编写和部署的。这些组件需要安全地通信和可观察，为 Istio 提供了吃自己的狗粮（或“喝自己的香槟”，用更法语版的比喻！）的机会。

## The cost of complexity

## 复杂度的代价

Good teams look back upon their choices and, with the benefit of hindsight, revisit them. Generally, when a team adopts microservices and their inherent complexity, they look for improvements in other areas to justify the tradeoffs. Let’s look at the Istio control plane through that lens.

优秀的团队会回顾他们的选择，并在事后重新审视他们。通常，当一个团队采用微服务及其固有的复杂性时，他们会寻找其他领域的改进来证明权衡的合理性。让我们通过这个镜头来看看 Istio 控制平面。

- **Microservices empower you to write in different languages.** The data plane (the Envoy proxy) is written in C++, and this boundary benefits from a clean separation in terms of the xDS APIs. However, all of the Istio control plane components are written in Go. We were able to choose the appropriate language for the appropriate job: highly performant C++ for the proxy, but accessible and speedy-development for everything else.

- **微服务使您能够使用不同的语言进行编写。** 数据平面（Envoy 代理）是用 C++ 编写的，这一边界得益于 xDS API 的清晰分离。但是，所有 Istio 控制平面组件都是用 Go 编写的。我们能够为适当的工作选择适当的语言：代理的高性能 C++，但其他一切都可以访问和快速开发。

- **Microservices empower you to allow different teams to manage services individually.**. In the vast majority of Istio installations, all the components are installed and operated by a single team or individual. The componentization done within Istio is aligned along the boundaries of the development teams who build it. This would make sense if the Istio components were delivered as a managed service by the people who wrote them, but this is not the case! Making life simpler for the development teams had an outsized impact of the usability for the orders-of-magnitude more users.

- **微服务使您能够允许不同的团队单独管理服务。**。在绝大多数 Istio 安装中，所有组件都由单个团队或个人安装和操作。在 Istio 中完成的组件化与构建它的开发团队的边界保持一致。如果 Istio 组件由编写它们的人作为托管服务交付，这将是有意义的，但事实并非如此！让开发团队的生活更简单对数量级更多用户的可用性产生了巨大的影响。

- **Microservices empower you to decouple versions, and release different components at different times.** All the components of the control plane have always been released at the same version, at the same time. We have never tested or supported running different versions of (for example) Citadel and Pilot.

- **微服务使您能够解耦版本，并在不同时间发布不同的组件。** 控制平面的所有组件始终在同一版本、同一时间发布。我们从未测试或支持运行不同版本的（例如）Citadel 和 Pilot。

- **Microservices empower you to scale components independently.** In Istio 1.5, control plane costs are dominated by a single feature: serving the Envoy xDS APIs that program the data plane. Every other feature has a marginal cost, which means there is very little value to having those features in separately-scalable microservices. 

- **微服务使您能够独立扩展组件。** 在 Istio 1.5 中，控制平面成本由单一功能主导：为对数据平面进行编程的 Envoy xDS API 提供服务。每个其他功能都有边际成本，这意味着在可单独扩展的微服务中拥有这些功能几乎没有价值。

- **Microservices empower you to maintain security boundaries.** Another good reason to separate an application into different microservices is if they have different security roles. Multiple Istio microservices like the sidecar injector, the Envoy bootstrap, Citadel, and Pilot hold nearly equivalent permissions to change the proxy configuration. Therefore, exploiting any of these services would cause near equivalent damage. When you deploy Istio, all the components are installed by default into the same Kubernetes namespace, offering limited security isolation.


- **微服务使您能够维护安全边界。** 将应用程序分成不同的微服务的另一个很好的理由是它们是否具有不同的安全角色。多个 Istio 微服务，如 sidecar 注入器、Envoy 引导程序、Citadel 和 Pilot，拥有几乎相同的权限来更改代理配置。因此，利用这些服务中的任何一个都会造成几乎同等的损害。部署 Istio 时，所有组件默认安装在同一个 Kubernetes 命名空间中，提供有限的安全隔离。


## The benefit of consolidation: introducing istiod

## 整合的好处：引入 istiod

Having established that many of the common benefits of microservices didn't apply to the Istio control plane, we decided to unify them into a single binary: **istiod** (the 'd' is for [daemon](https://en.wikipedia.org/wiki/Daemon_%28computing%29)).

确定微服务的许多共同优势不适用于 Istio 控制平面后，我们决定将它们统一为一个二进制文件：**istiod**（“d”代表 [daemon](https://en.wikipedia.org/wiki/Daemon_%28computing%29))。

Let’s look at the benefits of the new packaging:

让我们看看新包装的好处：

- **Installation becomes easier.** Fewer Kubernetes deployments and associated configurations are required, so the set of configuration options and flags for Istio is reduced significantly. In the simplest case, **_you can start the Istio control plane, with all features enabled, by starting a single Pod._**

- **安装变得更容易。** 需要更少的 Kubernetes 部署和相关配置，因此 Istio 的配置选项和标志集显着减少。在最简单的情况下，**_您可以通过启动单个 Pod 来启动 Istio 控制平面，并启用所有功能。_**

- **Configuration becomes easier.** Many of the configuration options that Istio has today are ways to orchestrate the control plane components, and so are no longer needed. You also no longer need to change cluster-wide `PodSecurityPolicy` to deploy Istio.

- **配置变得更容易。** Istio 今天拥有的许多配置选项都是编排控制平面组件的方法，因此不再需要。您也不再需要更改集群范围的 `PodSecurityPolicy` 来部署 Istio。

- **Using VMs becomes easier.** To add a workload to a mesh, you now just need to install one agent and the generated certificates. That agent connects back to only a single service.

- **使用虚拟机变得更容易。**要将工作负载添加到网格，您现在只需要安装一个代理和生成的证书。该代理仅连接回单个服务。

- **Maintenance becomes easier.** Installing, upgrading, and removing Istio no longer require a complicated dance of version dependencies and startup orders. For example: To upgrade, you only need to start a new istiod version alongside your existing control plane, canary it, and then move all traffic over to it.

- **维护变得更容易。** 安装、升级和删除 Istio 不再需要复杂的版本依赖关系和启动顺序。例如：要升级，您只需要在现有控制平面旁边启动一个新的 istiod 版本，对其进行 Canary，然后将所有流量转移到它。

- **Scalability becomes easier.** There is now only one component to scale.

- **可扩展性变得更容易。** 现在只有一个组件可以扩展。

- **Debugging becomes easier.** Fewer components means less cross-component environmental debugging.

- **调试变得更容易。** 更少的组件意味着更少的跨组件环境调试。

- **Startup time goes down.** Components no longer need to wait for each other to start in a defined order.

- **启动时间减少。** 组件不再需要按照定义的顺序等待彼此启动。

- **Resource usage goes down and responsiveness goes up.** Communication between components becomes guaranteed, and not subject to gRPC size limits. Caches can be shared safely, which decreases the resource footprint as a result.


- **资源使用率下降，响应性上升。** 组件之间的通信得到保证，不受 gRPC 大小限制。缓存可以安全共享，从而减少资源占用。


istiod unifies functionality that Pilot, Galley, Citadel and the sidecar injector previously performed, into a single binary.

istiod 将 Pilot、Galley、Citadel 和 sidecar 注入器之前执行的功能统一到一个二进制文件中。

A separate component, the istio-agent, helps each sidecar connect to the mesh by securely passing configuration and secrets to the Envoy proxies. While the agent, strictly speaking, is still part of the control plane, it runs on a per-pod basis. We’ve further simplified by rolling per-node functionality that used to run as a DaemonSet, into that per-pod agent.

一个单独的组件 istio-agent 通过将配置和机密安全地传递给 Envoy 代理，帮助每个边车连接到网格。严格来说，虽然代理仍然是控制平面的一部分，但它在每个 Pod 的基础上运行。我们通过将过去作为 DaemonSet 运行的每个节点功能滚动到每个 Pod 代理中来进一步简化。

## Extra for experts

## 额外的专家

There will still be some cases where you might want to run Istio components independently, or replace certain components.

在某些情况下，您可能希望独立运行 Istio 组件，或者替换某些组件。

Some users might want to use a Certificate Authority (CA) outside the mesh, and we have [documentation on how to do that](http://istio.io/latest/docs/tasks/security/cert-management/plugin-ca-cert/). If you do your certificate provisioning using a different tool, we can use that instead of the built-in CA.

一些用户可能想要在网格之外使用证书颁发机构 (CA)，我们有 [关于如何做到这一点的文档](http://istio.io/latest/docs/tasks/security/cert-management/plugin-ca-证书/)。如果您使用不同的工具进行证书配置，我们可以使用它而不是内置 CA。

## Moving forward

##  向前进

At its heart, istiod is just a packaging and optimization change. It’s built on the same code and API contracts as the separate components, and remains covered by our comprehensive test suite. This gives us confidence in making it the default in Istio 1.5. The service is now called `istiod` \- you’ll see an `istio-pilot` for existing proxies as the upgrade process completes.

从本质上讲，istiod 只是一个打包和优化更改。它构建在与单独组件相同的代码和 API 契约之上，并且仍然由我们的综合测试套件覆盖。这让我们有信心将其设为 Istio 1.5 的默认设置。该服务现在称为 `istiod` \- 当升级过程完成时，您将看到现有代理的 `istio-pilot`。

While the move to istiod may seem like a big change, and is a huge improvement for the people who _administer_ and _maintain_ the mesh, it won’t make the day-to-day life of _using_ Istio any different. istiod is not changing any of the APIs used to configure your mesh, so your existing processes will all stay the same. 

虽然迁移到 istiod 似乎是一个很大的变化，并且对于_管理_和_维护_网格的人来说是一个巨大的进步，但它不会使_使用_Istio的日常生活有任何不同。 istiod 不会更改用于配置网格的任何 API，因此您现有的流程将保持不变。

Does this change imply that microservice are a mistake for _all_ workloads and architectures? Of course not. They are a tool in a toolbelt, and they work best when they are reflected in your organizational reality. Instead, this change shows a willingness in the project to change based on user feedback, and a continued focus on simplification for all users. Microservices have to be right sized, and we believe we have found the right size for Istio. 

这种变化是否意味着微服务对于_所有_工作负载和架构来说都是错误的？当然不是。它们是工具带中的工具，当它们反映在您的组织现实中时，它们会发挥最佳作用。相反，此更改表明项目愿意根据用户反馈进行更改，并继续关注所有用户的简化。微服务必须有合适的规模，我们相信我们已经为 Istio 找到了合适的规模。

