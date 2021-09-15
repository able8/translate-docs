# The Cloud Native Landscape: The Orchestration and Management Layer

# 云原生景观：编排和管理层

#### 15 Dec 2020 11:54am,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack. io/author/jason-morgan/ "Posts by Jason Morgan")

#### 2020 年 12 月 15 日上午 11:54，作者：[Catherine Paganini](https://thenewstack.io/author/catherine-paganini/“Catherine Paganini 的帖子”) 和 [Jason Morgan](https://thenewstack.io/author/jason-morgan/“杰森摩根的帖子”)

![](https://cdn.thenewstack.io/media/2020/12/57e7bc17-meteora-3717220_640.jpg)

_This post is part of an ongoing series from_ [_CNCF Business Value Subcommittee_](https://lists.cncf.io/g/cncf-business-value) _co-chairs_ [_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) _and_ [_Jason Morgan_](https://thenewstack.io/author/jason-morgan/) _that focuses on explaining each category of the cloud native landscape to a non-technical audience as well as engineers just getting started with cloud native._

_这篇文章是来自_[_CNCF商业价值小组委员会_](https://lists.cncf.io/g/cncf-business-value)_co-chairs_[_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) _and_ [_Jason Morgan_](https://thenewstack.io/author/jason-morgan/) _专注于向非技术受众解释云原生景观的每个类别以及刚刚开始使用云原生的工程师。_

[![](https://cdn.thenewstack.io/media/2020/12/5984c027-screen-shot-2020-12-08-at-8.51.01-am.png)\
\
Catherine Paganini\
\
Catherine is Head of Marketing at Buoyant, the creator of Linkerd. A marketing leader turned cloud native evangelist, Catherine is passionate about educating business leaders on the new stack and the critical flexibility it provides.](https://www.linkedin.com/in/catherinepaganini/en/)

\
凯瑟琳·帕格尼尼\
\
Catherine 是 Linkerd 的创建者 Buoyant 的营销主管。 Catherine 是一名营销领导者，后来成为云原生布道者，她热衷于就新堆栈及其提供的关键灵活性对业务领导者进行教育。](https://www.linkedin.com/in/catherinepaganini/en/)

The orchestration and management layer is the third layer in the [Cloud Native Computing Foundation’s cloud native landscape](https://landscape.cncf.io). Before tackling tools in this category, engineers have presumably already automated infrastructure provisioning following security and compliance standards ( [provisioning layer](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)) and set up the runtime for the application ( [runtime layer](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)). Now they must figure out how to orchestrate and manage all app components as a group. Components must identify one another to communicate and coordinate to accomplish a common goal. Inherently scalable, cloud native apps rely on automation and resilience, enabled by these tools.

编排和管理层是 [云原生计算基金会云原生景观](https://landscape.cncf.io)中的第三层。在处理此类工具之前，工程师可能已经按照安全性和合规性标准自动进行基础设施配置（[配置层](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)) 并为应用程序设置运行时 ([runtime layer](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/))。现在，他们必须弄清楚如何将所有应用程序组件作为一个组进行编排和管理。组件必须相互识别以进行通信和协调以实现共同目标。本质上可扩展的云原生应用依赖于这些工具所支持的自动化和弹性。

### **_Sidenote:_**

###  **_边注：_**

[![](https://cdn.thenewstack.io/media/2020/08/572c2815-jason.png)\
\
Jason Morgan\
\
Jason Morgan, a Solutions Engineer with VMware, focuses on helping customers build and mature microservices platforms. Passionate about helping others on their cloud native journey, Jason enjoys sharing lessons learned with the broader developer community.](https://blog.59s.io/)

\
杰森摩根\
\
Jason Morgan 是 VMware 的解决方案工程师，专注于帮助客户构建和成熟的微服务平台。 Jason 热衷于在云原生之旅中帮助他人，他喜欢与更广泛的开发人员社区分享经验教训。](https://blog.59s.io/)

When looking at the [Cloud Native Landscape](https://landscape.cncf.io), you’ll note a few distinctions:

在查看 [云原生景观](https://landscape.cncf.io) 时，您会注意到一些区别：

- **Projects in large boxes** are CNCF-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- **Projects in small white boxes** are open source projects.
- **Products/projects in gray boxes** are proprietary.

- **大盒子中的项目**是 CNCF 托管的开源项目。有的还在孵化阶段（浅蓝色/紫色框），有的则是毕业项目（深蓝色框）。
- **小白盒中的项目**是开源项目。
- **灰色框中的产品/项目** 是专有的。

Please note that even during the time of this writing, we saw new projects becoming part of the CNCF so always refer to the actual landscape — things are moving fast!

请注意，即使在撰写本文时，我们也看到新项目成为 CNCF 的一部分，因此请始终参考实际情况——事情进展很快！

## Orchestration and Scheduling

## 编排和调度

### What It Is

###  这是什么

Orchestration and scheduling refer to running and managing containers, a new-ish way to package and ship applications, across a cluster. A cluster is a group of machines, physical or virtual, connected over a network.

编排和调度是指运行和管理容器，这是一种跨集群打包和传送应用程序的新方式。集群是一组通过网络连接的物理或虚拟机器。

Container orchestrators (and schedulers) are somewhat similar to the operating system (OS) on your laptop which manages all your apps (e.g. Microsoft 360, Slack, Zoom, etc.). The OS executes the apps you want to use and schedules when which app gets to use your laptop’s CPU and other hardware resources. 

容器编排器（和调度器）有点类似于笔记本电脑上的操作系统 (OS)，它管理着你的所有应用程序（例如 Microsoft 360、Slack、Zoom 等）。操作系统执行您想要使用的应用程序并安排哪个应用程序使用您的笔记本电脑的 CPU 和其他硬件资源。

While running everything on a single machine is great, most applications today are a lot bigger than one computer can possibly handle. Think Gmail or Netflix. These massive apps are distributed across multiple machines forming a [distributed application](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/). Most modern-day applications are distributed, and that requires software that is able to manage all components running across these different machines. In short, you need a “cluster OS.” That’s where orchestration tools come in.

虽然在一台机器上运行所有东西很棒，但今天的大多数应用程序都比一台计算机可能处理的要大得多。想想 Gmail 或 Netflix。这些海量应用分布在多台机器上，形成一个[分布式应用](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/)。大多数现代应用程序都是分布式的，这需要能够管理在这些不同机器上运行的所有组件的软件。简而言之，您需要一个“集群操作系统”。这就是编排工具的用武之地。

If you’ve read our previous articles, you probably noticed that containers come up time and again. Their ability to run apps in many different environments is key. And so are container orchestrators which, in most cases, is [Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-care/). Containers and Kubernetes are both central to cloud native architectures, which is why we hear so much about them.

如果您阅读过我们之前的文章，您可能会注意到容器一次又一次地出现。他们在许多不同环境中运行应用程序的能力是关键。容器编排器也是如此，在大多数情况下，它是 [Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-关心/)。容器和 Kubernetes 都是云原生架构的核心，这就是我们经常听到它们的原因。

### Problem It Addresses

### 它解决的问题

In cloud native architectures, applications are broken down into small components, or services, each placed in a container. You may have heard of them referred to as microservices. Instead of having one large application, you now have multiple small services each in need of resources, monitoring, and fixing if a problem occurs. While it is feasible to do those things manually for a single service, you’ll need automated processes when you have hundreds of containers.

在云原生架构中，应用程序被分解成小的组件或服务，每个组件或服务都放置在一个容器中。您可能听说过它们被称为微服务。您现在拥有多个小型服务，而不是拥有一个大型应用程序，每个服务都需要资源、监控和在出现问题时进行修复。虽然为单个服务手动执行这些操作是可行的，但当您拥有数百个容器时，您将需要自动化流程。

### How It Helps

### 它如何帮助

Container orchestrators automate container management. But what does that mean in practice? Let’s answer that for Kubernetes specifically since it’s the de facto container orchestrator.

容器编排器使容器管理自动化。但这在实践中意味着什么？让我们专门针对 Kubernetes 回答这个问题，因为它是事实上的容器编排器。

Kubernetes does something called desired state reconciliation: it matches the _current_ state of containers within a cluster to the _desired_ state. The desired state is specified by the engineer in a file (e.g. ten instances of service A running on three nodes, i.e. machines, with access to database B, etc.) and continuously compares against the actual state. If the desired and actual state don’t match, Kubernetes reconciles them by creating or destroying objects (e.g. if a container crashes, it will spin a new one up).

Kubernetes 做了一些称为期望状态协调的事情：它将集群内容器的 _current_ 状态与 _desired_ 状态相匹配。所需状态由工程师在文件中指定（例如，在三个节点上运行的服务 A 的十个实例，即可以访问数据库 B 的机器等），并不断与实际状态进行比较。如果期望状态和实际状态不匹配，Kubernetes 会通过创建或销毁对象来协调它们（例如，如果容器崩溃，它将旋转一个新的）。

In short, Kubernetes allows you to treat a cluster as one computer. It focuses only on what that environment should look like and handles the implementation details for you.

简而言之，Kubernetes 允许您将集群视为一台计算机。它只关注环境应该是什么样子，并为您处理实现细节。

### Technical 101

### 技术 101

Kubernetes lives in the orchestration and scheduling section along with other container orchestrators like Docker Swarm and Mesos. Its basic purpose is to allow you to manage a number of disparate computers as a single pool of resources. On top of that, it allows you to manage them in a declarative way, i.e. instead of telling Kubernetes how to do something you provide a definition of what you want to be done. This allows you to maintain the desired state in one or more YAML files and apply it more or less unchanged to any Kubernetes cluster. The orchestrator itself then creates anything that’s missing or deletes anything that should no longer exist. While Kubernetes isn’t the only orchestrator that the CNCF hosts as a project (both Crossplane and Volcano are incubating projects) it is the most commonly used and actively maintained.

Kubernetes 与其他容器编排器（如 Docker Swarm 和 Mesos）一起存在于编排和调度部分。它的基本目的是允许您将许多不同的计算机作为单个资源池进行管理。最重要的是，它允许您以声明性方式管理它们，即您无需告诉 Kubernetes 如何做某事，而是提供您想要完成的工作的定义。这允许您在一个或多个 YAML 文件中维护所需的状态，并将其应用到任何 Kubernetes 集群或多或少不变。然后，协调器本身会创建丢失的任何内容或删除不应再存在的任何内容。虽然 Kubernetes 不是 CNCF 作为项目托管的唯一编排器（Crossplane 和 Volcano 都是孵化项目），但它是最常用和积极维护的。

**Buzzwords**

**流行语**

**Popular Projects**

**热门项目**

- Cluster
- Scheduler
- Orchestration

- 簇
- 调度程序
- 编排

- Kubernetes
- Docker Swarm
- Mesos

- Kubernetes
- Docker Swarm
- 金币

![Scheduling and orchestration](https://cdn.thenewstack.io/media/2020/12/cf616dae-screen-shot-2020-12-08-at-8.21.04-am.png)

## Coordination and Service Discovery

## 协调和服务发现

### What It Is

###  这是什么

As we’ve seen, modern applications are composed of multiple individual services that need to collaborate to provide value to the end user. To collaborate, they communicate over a network (discussed in our [runtime layer article](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)). And to communicate, they must first locate one another. Service discovery is the process of figuring out how to do that.

正如我们所见，现代应用程序由多个单独的服务组成，这些服务需要协作才能为最终用户提供价值。为了协作，他们通过网络进行通信（在我们的[运行时层文章](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)中讨论)。要进行交流，他们必须首先找到彼此。服务发现是弄清楚如何做到这一点的过程。

### Problem It Addresses 

### 它解决的问题

Cloud native architectures are dynamic and fluid, meaning they are constantly changing. When a container crashes on one node, a new container is spun up on a different node to replace it. Or, when an app scales, replicas are spread out throughout the network. There is no one place where a particular service is; the location of everything is constantly changing. Tools in this category keep track of services within the network so services can find one another when needed.

云原生架构是动态和流动的，这意味着它们在不断变化。当一个容器在一个节点上崩溃时，一个新的容器会在另一个节点上启动以替换它。或者，当应用扩展时，副本会分布在整个网络中。没有特定服务所在的地方；一切事物的位置都在不断变化。此类别中的工具会跟踪网络内的服务，以便服务可以在需要时找到彼此。

### How It Helps

### 它如何帮助

Service discovery tools address this problem by providing a common place to find and potentially identify individual services. There are basically two types of tools in this category: (1) Service discovery engines are database-like tools that store info about what services exist and how to locate them. And (2) name resolution tools (e.g. Core DNS) receive service location requests and return network address information.

服务发现工具通过提供一个公共位置来查找和潜在地识别单个服务来解决这个问题。该类别中基本上有两种类型的工具： (1) 服务发现引擎是类似数据库的工具，用于存储有关存在哪些服务以及如何定位它们的信息。 (2) 名称解析工具（例如 Core DNS）接收服务位置请求并返回网络地址信息。

### **_Sidenote:_**

###  **_边注：_**

In Kubernetes, to make a pod reachable a new abstraction layer called “service”is introduced. Services provide a single stable address for a dynamically changing group of pods.

在 Kubernetes 中，为了使 pod 可访问，引入了一个称为“服务”的新抽象层。服务为动态变化的 pod 组提供一个稳定的地址。

Please note that “service” may have different meanings in different contexts, which can be quite confusing. The term “services” generally refers to the service placed inside a container/pod. It’s the app component or microservice with a specific function within the actual app (e.g. face recognition algorithm of your iPhone).

请注意，“服务”在不同的上下文中可能有不同的含义，这可能会令人困惑。术语“服务”通常是指放置在容器/pod 内的服务。它是实际应用程序中具有特定功能的应用程序组件或微服务（例如 iPhone 的人脸识别算法）。

A Kubernetes service, on the other hand, is the abstraction that helps pods find and connect to each other. It is an entry point for a service (functionality) as a collection of processes or pods. In Kubernetes, when you create a service (abstraction), you create a group of pods that together provide a service (functionality) with a single endpoint (entry) which is the Kubernetes service.

另一方面，Kubernetes 服务是帮助 Pod 找到并相互连接的抽象。它是作为进程或 Pod 集合的服务（功能）的入口点。在 Kubernetes 中，当您创建服务（抽象）时，您会创建一组 pod，这些 pod 一起提供具有单个端点（条目）的服务（功能），即 Kubernetes 服务。

### Technical 101

### 技术 101

As distributed systems became more and more prevalent, traditional DNS processes and traditional load balancers were often unable to keep up with changing endpoint information. To make up for these shortcomings, service discovery tools were created to handle individual application instances rapidly registering and deregistering themselves. Some options such as etcd and CoreDNS are Kubernetes native, others have custom libraries or tools to allow services to operate effectively. CoreDNS and etcd are CNCF projects and are built into Kubernetes.

随着分布式系统变得越来越普遍，传统的 DNS 流程和传统的负载均衡器往往无法跟上不断变化的端点信息。为了弥补这些缺点，创建了服务发现工具来处理各个应用程序实例的快速注册和注销。某些选项（例如 etcd 和 CoreDNS）是 Kubernetes 原生的，其他选项则具有自定义库或工具以允许服务有效运行。 CoreDNS 和 etcd 是 CNCF 项目，内置于 Kubernetes 中。

**Buzzwords**

**流行语**

**Popular Projects**

**热门项目**

- DNS
- Service Discovery

- DNS
- 服务发现

- CoreDNS
- etcd
- Zookeeper
- Eureka

- 核心DNS
- etcd
- 动物园管理员
- 尤里卡

![Coordination and service discovery](https://cdn.thenewstack.io/media/2020/12/673ee53a-screen-shot-2020-12-08-at-8.23.02-am.png)

## Remote Procedure Call

## 远程过程调用

### What It Is

###  这是什么

Remote Procedure Call (RPC) is a particular technique enabling applications to talk to each other. It represents one of a few ways for applications to structure their communications with one another. This section is not particularly relevant to non-developers.

远程过程调用 (RPC) 是一种使应用程序能够相互通信的特殊技术。它代表了应用程序构建彼此通信的几种方式之一。本节与非开发人员不是特别相关。

### Problem It Addresses

### 它解决的问题

Modern apps are composed of numerous individual services that must communicate in order to collaborate. RPC is one option for handling the communication between applications.

现代应用程序由众多单独的服务组成，这些服务必须进行通信才能进行协作。 RPC 是处理应用程序之间通信的一种选择。

### How It Helps

### 它如何帮助

RPC provides a tightly coupled and highly opinionated way of handling communication between services. It allows for bandwidth-efficient communications and many languages enable RPC interface implementations. RPC is not the only nor the most common way to address this problem.

RPC 提供了一种紧密耦合且高度自以为是的方式来处理服务之间的通信。它允许带宽高效的通信，并且许多语言支持 RPC 接口实现。 RPC 不是解决此问题的唯一方法，也不是最常用的方法。

### Technical 101

### 技术 101

RPC provides a highly structured and tightly coupled interface for communications between services. There are a lot of potential benefits with RPC: It makes coding the connection easier, it allows for extremely efficient use of the network layer and well-structured communications between services. RPC has also been criticized for creating brittle connection points and forcing users to do coordinated upgrades for multiple services. gRPC is a particularly popular RPC implementation and has been adopted by the CNCF.

RPC 为服务之间的通信提供了高度结构化和紧密耦合的接口。 RPC 有很多潜在的好处：它使连接编码变得更容易，它允许极其有效地使用网络层和服务之间结构良好的通信。 RPC 也因创建脆弱的连接点并迫使用户对多个服务进行协调升级而受到批评。 gRPC 是一种特别流行的 RPC 实现，并已被 CNCF 采用。

**Buzzwords**

**流行语**

**Popular Projects**

**热门项目**

- gRPC

- gRPC

- gRPC

- gRPC

![Remote procedure call](https://cdn.thenewstack.io/media/2020/12/1e4e9de8-screen-shot-2020-12-08-at-8.26.32-am.png)

## Service Proxy

## 服务代理

### What It Is 

###  这是什么

A service proxy is a tool that intercepts traffic to or from a given service, applies some logic to it, then forwards that traffic to another service. It essentially acts as a “go-between” that collects information about network traffic and/or applies rules to it. This can be as simple as serving as a load balancer that forwards traffic to individual applications or as complex as an interconnected mesh of proxies running side by side with individual containerized applications handling all network connections.

服务代理是一种工具，它拦截进出给定服务的流量，对其应用一些逻辑，然后将该流量转发到另一个服务。它本质上充当“中间人”，收集有关网络流量的信息和/或对其应用规则。这可以像负载均衡器一样简单，将流量转发到各个应用程序，也可以像代理的互连网格一样复杂，与处理所有网络连接的各个容器化应用程序并排运行。

While a service proxy is useful in and of itself, especially when driving traffic from the broader network into a Kubernetes cluster, service proxies are also building blocks for other systems, such as API gateways or service meshes, which we’ll discuss below.

虽然服务代理本身很有用，尤其是在将流量从更广泛的网络驱动到 Kubernetes 集群时，服务代理也是其他系统的构建块，例如 API 网关或服务网格，我们将在下面讨论。

### Problem It Addresses

### 它解决的问题

Applications should send and receive network traffic in a controlled manner. To keep track of the traffic and potentially transform or redirect it, we need to collect data. Traditionally, the code enabling data collection and network traffic management was embedded within each application.

应用程序应以受控方式发送和接收网络流量。为了跟踪流量并可能对其进行转换或重定向，我们需要收集数据。传统上，支持数据收集和网络流量管理的代码嵌入在每个应用程序中。

A service proxy allows us to “externalize” this functionality. No longer does it need to live within the apps. Instead, it’s now embedded into the platform layer (where your apps run). This is incredibly powerful because it allows developers to fully focus on writing application logic, your value generating code, while the universal task of handling traffic is managed by the platform team (whose responsibility it is in the first place). By centralizing the distribution and management of globally needed service functionality (e.g. routing or TLS termination) from a single, common location, communication between services is more reliable, secure, and performant.

服务代理允许我们“外部化”这个功能。它不再需要存在于应用程序中。相反，它现在嵌入到平台层（您的应用程序运行的地方）。这是非常强大的，因为它允许开发人员完全专注于编写应用程序逻辑，您的价值生成代码，而处理流量的通用任务由平台团队管理（首先是其职责）。通过从单个公共位置集中分发和管理全球所需的服务功能（例如路由或 TLS 终止），服务之间的通信更加可靠、安全和高效。

### How It Helps

### 它如何帮助

Proxies act as gatekeepers between the user and services or between different services. With this unique positioning, they provide insight into what type of communication is happening. Based on their insight, they determine where to send a particular request or even deny it entirely.

代理充当用户和服务之间或不同服务之间的看门人。凭借这种独特的定位，他们可以深入了解正在发生的通信类型。根据他们的洞察力，他们决定将特定请求发送到哪里，甚至完全拒绝它。

Proxies gather critical data, manage routing (spreading traffic evenly among services or rerouting if some services break down), encrypt connections, and cache content (reducing resource consumption).

代理收集关键数据、管理路由（在服务之间均匀分布流量或在某些服务出现故障时重新路由）、加密连接和缓存内容（减少资源消耗）。

### Technical 101

### 技术 101

Service proxies work by intercepting traffic between services, performing some logic on them, then potentially allowing traffic to move on. By putting a centrally controlled set of capabilities into this proxy, administrators are able to accomplish several things. They can gather detailed metrics about inter-service communication, protect services from being overloaded, and apply other common standards to services, like mutual TLS. Service proxies are fundamental to other tools like service meshes as they provide a way to enforce higher-level policies to all network traffic.

服务代理的工作原理是拦截服务之间的流量，对它们执行一些逻辑，然后可能允许流量继续前进。通过将集中控制的一组功能放入此代理，管理员可以完成几件事。他们可以收集有关服务间通信的详细指标，保护服务免于过载，并将其他通用标准应用于服务，例如双向 TLS。服务代理是服务网格等其他工具的基础，因为它们提供了一种对所有网络流量实施更高级别策略的方法。

Please note, the CNCF includes load balancers and ingress providers into this category. Envoy, Contour, and BFE are all CNCF projects.

请注意，CNCF 将负载平衡器和入口提供程序包含在此类别中。 Envoy、Contour、BFE 都是 CNCF 项目。

**Buzzwords**

**流行语**

**Popular Projects**

**热门项目**

- Service proxy
- Ingress

- 服务代理
- 入口

- Envoy
- Contour
- NGINX

- 特使
- 轮廓
- NGINX

![Service proxy](https://cdn.thenewstack.io/media/2020/12/ed807753-screen-shot-2020-12-08-at-8.28.02-am.png)

## API Gateway

## API 网关

### What It Is

###  这是什么

While humans generally interact with computer programs via a GUI (graphical user interface) such as a webpage or a (desktop) application, computers interact with each other through APIs (application programming interfaces). But an API shouldn’t be confused with an API gateway.

虽然人类通常通过网页或（桌面）应用程序等 GUI（图形用户界面）与计算机程序交互，但计算机通过 API（应用程序编程接口）相互交互。但是不应将 API 与 API 网关混淆。

An API gateway allows organizations to move key functions, such as authorizing or limiting the number of requests between applications, to a centrally managed location. It also functions as a common interface to (often external) API consumers.

API 网关允许组织将关键功能（例如授权或限制应用程序之间的请求数量）移动到集中管理的位置。它还充当（通常是外部）API 使用者的通用接口。

Through API gateways, organizations can centrally control (limit or enable) interactions between apps and keep track of them, enabling things like chargeback, authentication, and protecting services from being overused (aka rate-limiting).

通过 API 网关，组织可以集中控制（限制或启用）应用程序之间的交互并跟踪它们，从而实现诸如退款、身份验证和保护服务不被过度使用（又名速率限制）之类的功能。

### Example 

###  例子

Take Amazon store cards. To offer them, Amazon partners with a bank that will issue and manage all Amazon store cards. In return, the bank will keep, let’s say, $1 per transaction. To authorize the retailer to request new cards, keep track of transactions, and maybe even restrict the number of requested cards per minute, the bank will use an API gateway. All that functionality is encoded into the gateway, not the services using it. Services just worry about issuing cards.

以亚马逊商店卡为例。为了提供这些服务，亚马逊与一家银行合作，该银行将发行和管理所有亚马逊商店卡。作为回报，银行将保留每笔交易 1 美元。为了授权零售商申请新卡、跟踪交易，甚至可能限制每分钟申请的卡数量，银行将使用 API 网关。所有这些功能都被编码到网关中，而不是使用它的服务。服务只是担心发卡。

### Problem It Addresses

### 它解决的问题

While most containers and core applications have an API, an API gateway is more than just an API. An API gateway simplifies how organizations manage and apply rules to all interactions.

虽然大多数容器和核心应用程序都有一个 API，但 API 网关不仅仅是一个 API。 API 网关简化了组织管理和应用规则到所有交互的方式。

API gateways allow developers to write and maintain less custom code (the system functionality is encoded into the API gateway, remember?). They also enable teams to see and control the interactions between application users and the applications themselves.

API 网关允许开发人员编写和维护较少的自定义代码（系统功能被编码到 API 网关中，还记得吗？）。它们还使团队能够查看和控制应用程序用户与应用程序本身之间的交互。

### How It Helps

### 它如何帮助

An API gateway sits between the users and the application. It acts as a go-between that takes the messages (requests) from the users and forwards them to the appropriate service. But before handing the request off, it evaluates whether the user is allowed to do what they’re trying to do and records details about who made the request and how many requests they’ve made.

API 网关位于用户和应用程序之间。它充当中间人，从用户那里获取消息（请求）并将它们转发到适当的服务。但在提交请求之前，它会评估是否允许用户做他们想做的事情，并记录有关谁发出请求以及他们发出了多少请求的详细信息。

Put simply, an API gateway provides a single point of entry with a common user interface for app users. It also enables you to handoff tasks otherwise implemented within the app to the gateway, saving developer time and money.

简而言之，API 网关为应用程序用户提供了具有通用用户界面的单一入口点。它还使您能够将应用程序中以其他方式实现的任务移交给网关，从而节省开发人员的时间和金钱。

### Technical 101

### 技术 101

Like many categories in this layer, an API gateway takes custom code out of our apps and brings it into a central system. The API gateway works by intercepting calls to backend services, performing some kind of value add activity like validating authorization, collecting metrics, or transforming requests, then performing whatever action it deems appropriate. API gateways serve as a common entry point for a set of downstream applications while at the same time providing a place where teams can inject business logic to handle authorization, rate limiting, and chargeback. They allow application developers to abstract away changes to their downstream APIs from their customers and offload tasks like onboarding new customers to the gateway.

与该层中的许多类别一样，API 网关从我们的应用程序中提取自定义代码并将其带入中央系统。 API 网关的工作原理是拦截对后端服务的调用，执行某种增值活动，例如验证授权、收集指标或转换请求，然后执行它认为合适的任何操作。 API 网关充当一组下游应用程序的公共入口点，同时提供了一个地方，团队可以在其中注入业务逻辑来处理授权、速率限制和退款。它们允许应用程序开发人员从客户那里抽象出对其下游 API 的更改，并将新客户入职等任务卸载到网关。

**Buzzwords**

**流行语**

**Popular Projects/Products**

**热门项目/产品**

- API gateway

- API 网关

- Kong
- Mulesoft
- Ambassador

- 孔
- Mulesoft
- 大使

![API gateway](https://cdn.thenewstack.io/media/2020/12/0550d18a-screen-shot-2020-12-08-at-8.29.46-am.png)

## Service Mesh

## 服务网格

### What It Is

###  这是什么

If you’ve been reading a little bit about cloud native, the term service mesh probably rings a bell. It’s been getting quite a bit of attention lately. According to long-time TNS contributor Janakiram MSV, “after Kubernetes, the service mesh technology has become the most critical component of the cloud native stack.” So pay attention to this one — you’ll likely hear a lot more about it.

如果您一直在阅读有关云原生的一些内容，那么术语服务网格可能会敲响警钟。最近受到了相当多的关注。根据长期 TNS 贡献者 Janakiram MSV 的说法，“在 Kubernetes 之后，服务网格技术已成为云原生堆栈中最关键的组件。”所以请注意这个——你可能会听到更多关于它的信息。

Service meshes manage traffic (i.e. communication) between services. They enable platform teams to add reliability, observability, and security features uniformly across all services running within a cluster without requiring any code changes.

服务网格管理服务之间的流量（即通信）。它们使平台团队能够在集群内运行的所有服务中统一添加可靠性、可观察性和安全性功能，而无需更改任何代码。

### Problem It Addresses

### 它解决的问题

In a cloud native world, we are dealing with multiple services all in need to communicate. This means a lot more traffic is going back and forth on an inherently unreliable and often slow network. To address this new set of challenges, engineers must implement additional functionality. Prior to the service mesh, that functionality had to be encoded into every single application. This custom code often became a source of technical debt and provided new avenues for failures or vulnerabilities.

在云原生世界中，我们正在处理需要通信的多种服务。这意味着更多的流量在本质上不可靠且通常很慢的网络上来回传输。为了应对这一系列新挑战，工程师必须实现额外的功能。在服务网格出现之前，必须将该功能编码到每个应用程序中。这种自定义代码经常成为技术债务的来源，并为失败或漏洞提供了新的途径。

### How It Helps

### 它如何帮助

Service meshes add reliability, observability, and security features uniformly across all services on a platform layer without touching the app code. They are compatible with any programming language, allowing development teams to focus on writing business logic.

服务网格在平台层上的所有服务中统一添加可靠性、可观察性和安全性功能，而无需触及应用程序代码。它们与任何编程语言兼容，允许开发团队专注于编写业务逻辑。

### **_Sidenote:_** 

###  **_边注：_**

Since traditionally, these service mesh features had to be coded into each service, each time a new service was released or updated, the developer had to ensure these features were functional, too, providing a lot of room for human error. And here’s a dirty little secret, developers prefer focusing on business logic (value-generating functionalities) rather than building reliability, observability, and security features. For the platform owners, on the other hand, these are core capabilities and central to everything they do. Making developers responsible for adding features that platform owners need is inherently problematic. This, by the way, also applies to general-purpose proxies and API gateways mentioned above. Service meshes and API gateways solve that very issue as they are implemented by the platform owners and applied universally across all services.

由于传统上，这些服务网格功能必须编码到每个服务中，每次发布或更新新服务时，开发人员都必须确保这些功能也能正常运行，从而为人为错误提供了很大的空间。这里有一个肮脏的小秘密，开发人员更喜欢专注于业务逻辑（价值生成功能），而不是构建可靠性、可观察性和安全功能。另一方面，对于平台所有者来说，这些是核心能力，也是他们所做一切的核心。让开发人员负责添加平台所有者需要的功能本质上是有问题的。顺便说一下，这也适用于上面提到的通用代理和 API 网关。服务网格和 API 网关解决了这个问题，因为它们由平台所有者实施并普遍应用于所有服务。

### Technical 101

### 技术 101

Service meshes bind all services running on a cluster together via service proxies creating a mesh of services, hence service mesh. These are managed and controlled through the service mesh control plane. Service meshes allow platform owners to perform common actions or collect data on applications without having developers write custom logic.

服务网格通过创建服务网格的服务代理将集群上运行的所有服务绑定在一起，从而创建服务网格。这些通过服务网格控制平面进行管理和控制。服务网格允许平台所有者在应用程序上执行常见操作或收集数据，而无需开发人员编写自定义逻辑。

In essence, a service mesh is an infrastructure layer that manages inter-service communications by providing command and control signals to a network, or mesh, of service proxies. Its power lies in its ability to provide key system functionality without having to modify the applications.

本质上，服务网格是一个基础设施层，它通过向服务代理的网络或网格提供命令和控制信号来管理服务间通信。它的强大之处在于它无需修改应用程序即可提供关键系统功能的能力。

Some service meshes use a general-purpose service proxy (see above) for their data plane. Others use a dedicated proxy; Linkerd, for example, uses the [Linkerd2-proxy “micro proxy”](https://linkerd.io/) to gain an advantage in performance and resource consumption. These proxies are uniformly attached to each service through so-called sidecars. Sidecar refers to the fact that the proxy runs in its own container but lives in the same pod. Just like a motorcycle sidecar, it’s a separate module attached to the motorcycle, following it wherever it goes.

一些服务网格为其数据平面使用通用服务代理（见上文）。其他人使用专用代理；例如，Linkerd 使用 [Linkerd2-proxy “微代理”](https://linkerd.io/) 在性能和资源消耗方面获得优势。这些代理通过所谓的边车统一附加到每个服务。 Sidecar 是指代理运行在自己的容器中，但位于同一个 Pod 中。就像摩托车边车一样，它是一个独立的模块，附在摩托车上，无论它走到哪里都跟在它后面。

### Example:

###  例子：

Take circuit breaking. In microservice environments, individual components often fail or begin running slowly. Without a service mesh, developers would have to write custom logic to handle downstream failures gracefully and potentially set cooldown timers to avoid that upstream services continually request responses from degraded or failed downstream services. With a service mesh, that logic is handled at a platform level.

采取断路措施。在微服务环境中，单个组件经常出现故障或开始运行缓慢。如果没有服务网格，开发人员将不得不编写自定义逻辑来优雅地处理下游故障，并可能设置冷却计时器以避免上游服务不断请求来自降级或失败的下游服务的响应。使用服务网格，该逻辑在平台级别进行处理。

Service meshes provide many useful features, including the ability to surface detailed metrics, encrypt all traffic, limit what operations are authorized by what service, provide additional plugins for other tools, and much more. For more detailed information, check out the [service mesh interface](https://smi-spec.io/) specification.

服务网格提供了许多有用的功能，包括显示详细指标的能力、加密所有流量、限制哪些操作由哪些服务授权、为其他工具提供额外插件等等。有关更多详细信息，请查看 [服务网格接口](https://smi-spec.io/) 规范。

**Buzzwords**

**流行语**

**Popular Projects**

**热门项目**

- Service mesh
- Sidecar
- Data plane
- Control plane

- 服务网格
- 边车
- 数据平面
- 控制平面

- Linkerd
- Consul
- Istio

- 链接器
- 领事
- Istio

![Service mesh](https://cdn.thenewstack.io/media/2020/12/458f49aa-screen-shot-2020-12-08-at-8.30.08-am.png)

## Conclusion 

##  结论

As we’ve seen, tools in this layer deal with how all these independent containerized services are managed as a group. Orchestration and scheduling tools are some sort of cluster OS managing containerized applications across your cluster. Coordination and service discovery, service proxies, and service meshes ensure services can find each other and communicate effectively in order to collaborate as one cohesive app. API gateways are an additional layer providing even more control over service communication, in particular between external applications.In our next article, we’ll discuss the application definition and development layer — the last layer of the CNCF landscape. It covers databases, streaming and messaging, application definition, and image build, as well as continuous integration and delivery.

正如我们所见，这一层中的工具处理如何将所有这些独立的容器化服务作为一个组进行管理。编排和调度工具是某种集群操作系统，用于管理集群中的容器化应用程序。协调和服务发现、服务代理和服务网格确保服务可以找到彼此并有效通信，以便作为一个有凝聚力的应用程序进行协作。 API 网关是一个附加层，它提供对服务通信的更多控制，尤其是外部应用程序之间的通信。在下一篇文章中，我们将讨论应用程序定义和开发层——CNCF 格局的最后一层。它涵盖了数据库、流媒体和消息传递、应用程序定义和映像构建，以及持续集成和交付。

_As always, a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._

_一如既往，非常感谢来自 CNCF 的 [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/)，他非常友好地审阅了这篇文章，以确保其准确无误。_

The Cloud Native Computing Foundation and VMware are sponsors of The New Stack.

云原生计算基金会和 VMware 是 The New Stack 的赞助商。

Feature image [Antonios Ntoumas](https://pixabay.com/fr/users/atlantios-4957810/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3717220) de [Pixabay](https://pixabay.com/fr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3717220).

特色图片 [Antonios Ntoumas](https://pixabay.com/fr/users/atlantios-4957810/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3717220) de [Pixabay](https://pixabay)(https://pixabay)/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3717220)。

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: MADE, Docker, Ambassador, Prevalent, Bit. 

The New Stack 是 Insight Partners 的全资子公司。 TNS 所有者 Insight Partners 是以下公司的投资者：MADE、Docker、Ambassador、Prevalent、Bit。

