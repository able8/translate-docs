# The Cloud Native Landscape: The Runtime Layer Explained

# 云原生景观：解释运行时层

#### 29 Sep 2020 1:24pm,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack. io/author/jason-morgan/ "Posts by Jason Morgan")

#### 2020 年 9 月 29 日下午 1:24，作者：[Catherine Paganini](https://thenewstack.io/author/catherine-paganini/“Catherine Paganini 的帖子”) 和 [Jason Morgan](https://thenewstack.io/author/jason-morgan/“杰森摩根的帖子”)

_This post is part of an ongoing series from_ [_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) _and_ [_Jason Morgan_](https://thenewstack.io/author/jason-morgan/) _that focuses on explaining each category of the cloud native landscape to a non-technical audience as well as engineers just getting started with cloud native._

_这篇文章是来自_[_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category)_和_[_Jason Morgan_](https://thenewstack.io/author/jason-morgan/) _专注于向非技术受众以及刚刚开始使用云原生的工程师解释云原生环境的每个类别。_

Jason Morgan, a Solutions Engineer with VMware, focuses on helping customers build and mature microservices platforms. Passionate about helping others on their cloud native journey, Jason enjoys sharing lessons learned with the broader developer community.](https://blog.59s.io/)

Jason Morgan 是 VMware 的解决方案工程师，专注于帮助客户构建和成熟的微服务平台。 Jason 热衷于在云原生之旅中帮助他人，他喜欢与更广泛的开发人员社区分享经验教训。](https://blog.59s.io/)

In our previous article, we explored the [provisioning layer of the Cloud Native Computing Foundation's cloud native landscape](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/) which focuses on building the foundation of your cloud native platforms and applications. This article zooms into the runtime layer encompassing everything a container needs in order to run in a cloud native environment. That means the code used to start a container, referred to as the runtime engine; the tools to make persistent storage available to containers; and those that manage the container environment networks.

在我们之前的文章中，我们探讨了 [云原生计算基金会云原生景观的供应层](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)构建您的云原生平台和应用程序的基础。本文放大了运行时层，其中包含容器在云原生环境中运行所需的一切。这意味着用于启动容器的代码，称为运行时引擎；为容器提供持久存储的工具；以及那些管理容器环境网络的人。

But note, these resources shouldn’t be confused with the networking and storage work handled by the infrastructure and provisioning layer concerned with getting the container platform running. Tools in this category are used by the containers directly to start/stop, store data, and talk to each other.

但请注意，不应将这些资源与与让容器平台运行相关的基础设施和配置层处理的网络和存储工作相混淆。容器直接使用此类别中的工具来启动/停止、存储数据和相互通信。

![runtime layer ](https://cdn.thenewstack.io/media/2020/09/756d2023-screen-shot-2020-09-23-at-6.34.09-pm.png)

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

## Cloud Native Storage

## 云原生存储

### What It Is

###  这是什么

Storage is where the persistent data of an app is stored, often referred to as persistent volume. Easy access to it is critical for the app to function reliably. Generally, when we say persistent data we mean storing things like databases, messages, or any other information we want to ensure doesn’t disappear when an app gets restarted.

存储是存储应用程序持久数据的地方，通常称为持久卷。轻松访问它对于应用程序可靠运行至关重要。通常，当我们说持久数据时，我们指的是存储诸如数据库、消息或任何其他我们希望确保在应用程序重新启动时不会消失的信息。

### Problem It Addresses

### 它解决的问题

Cloud native architectures are fluid, flexible, and elastic, making persisting data between restarts challenging. To scale up and down or self-heal, containerized apps are continuously created and deleted, changing physical location over time. Therefore, cloud native storage must be provided in a node-independent fashion. To store data, however, you’ll need hardware — a disk to be specific — and disks, just like any other hardware, are infrastructure-bound. That’s the first big challenge.

云原生架构是流动的、灵活的和有弹性的，这使得在重启之间持久化数据具有挑战性。为了扩大和缩小规模或自我修复，容器化应用不断被创建和删除，随着时间的推移改变物理位置。因此，必须以独立于节点的方式提供云原生存储。然而，要存储数据，您需要硬件——具体来说是磁盘——而磁盘，就像任何其他硬件一样，受基础设施的限制。这是第一个大挑战。

Then there is the actual storage interface which can change significantly between datacenters (in the old world, each infrastructure had their own storage solution with its own interface), making portability really tough. And lastly, to benefit from the elasticity of the cloud, storage must be provisioned in an automated fashion as manual provisioning and autoscaling aren’t compatible.

然后是实际的存储接口，它可以在数据中心之间发生显着变化（在旧世界中，每个基础设施都有自己的存储解决方案和自己的接口），这使得可移植性变得非常困难。最后，为了从云的弹性中受益，必须以自动方式配置存储，因为手动配置和自动扩展不兼容。

Cloud native storage is tailored to this new cloud native reality.

云原生存储是为这种新的云原生现实量身定制的。

### How It Helps 

### 它如何帮助

The tools in this category help either a) provide cloud native storage options for containers, b) standardize the interfaces between containers and storage providers or c) provide data protection through backup and restore operations. The former means storage that uses a cloud native compatible container storage interface (aka tools in the second category) and which can be provisioned automatically, enabling autoscaling and self-healing by eliminating the human bottleneck.

此类别中的工具有助于 a) 为容器提供云原生存储选项，b) 标准化容器和存储提供商之间的接口，或 c) 通过备份和恢复操作提供数据保护。前者是指使用云原生兼容容器存储接口（即第二类工具）并且可以自动配置的存储，通过消除人为瓶颈来实现自动扩展和自我修复。

### Technical 101

### 技术 101

Cloud native storage is largely made possible by the Container Storage Interface (CSI) which allows a standard API for providing file and block storage to containers. There are a number of tools in this space, both open source and vendor-provided that leverage the CSI to provide on-demand storage to containers. In addition to that extremely important functionality, we have a number of other tools and technologies which aim to solve storage problems in the cloud native space. Minio is a popular project that, among other things, provides an S3-compatible API for object storage. Tools like Velero help simplify the process of backing up and restoring both the Kubernetes clusters themselves as well as persistent data used by the applications.

容器存储接口 (CSI) 在很大程度上使云原生存储成为可能，该接口允许使用标准 API 为容器提供文件和块存储。在这个领域有许多工具，包括开源的和供应商提供的，它们利用 CSI 为容器提供按需存储。除了这个极其重要的功能之外，我们还有许多其他工具和技术，旨在解决云原生空间中的存储问题。 Minio 是一个流行的项目，其中包括为对象存储提供与 S3 兼容的 API。 Velero 等工具有助于简化备份和恢复 Kubernetes 集群本身以及应用程序使用的持久数据的过程。

**Buzzwords**

**流行语**

**Popular Projects/Products**

**热门项目/产品**

- CSI
- Storage API
- Backup and Restore

- CSI
- 存储 API
- 备份还原

- **Minio**
- **CSI**
- **Ceph + Rook**
- **Velero**

- **迷你小**
- **CSI**
- **Ceph + Rook**
- **Velero**

![cloud native storage](https://cdn.thenewstack.io/media/2020/09/3160a300-screen-shot-2020-09-23-at-6.39.03-pm.png)

## Container Runtime

## 容器运行时

### What It Is

###  这是什么

As discussed in the [provisioning layer article](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/), a container is a set of **compute constraints** used to execute (that's tech-speak for launch) an application. Containerized apps believe they are running on their own dedicated computer and are oblivious that they are sharing resources with other processes (similar to virtual machines).

如 [供应层文章](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)中所述，容器是一组**计算约束**，用于执行（这是启动的技术术语）应用程序。容器化应用程序相信它们在自己的专用计算机上运行，并且没有意识到它们正在与其他进程（类似于虚拟机)共享资源。

The container _runtime_ is the software that executes containerized (or “constrained”) applications. Without the runtime, you only have the container image, the file specifying how the containerized app should look like. The runtime will start an app within a container and provide it with the needed resources.

容器 _runtime_ 是执行容器化（或“受限”）应用程序的软件。没有运行时，您只有容器映像，该文件指定容器化应用程序的外观。运行时将在容器内启动一个应用程序并为其提供所需的资源。

### Problem It Addresses

### 它解决的问题

Container images (the files with the application specs) must be launched in a standardized, secure, and isolated way. **Standardized** because you need standard operating rules no matter where they are running. **Secure**, well, because you don’t want anyone who shouldn’t access it to do so. And **isolated** because you don’t want the app to affect or be affected by other apps (for instance, if a co-located application crashes). Isolation basically functions as protection. Additionally, the application must be provided resources, from CPU to storage to memory.

容器镜像（包含应用程序规范的文件）必须以标准化、安全和隔离的方式启动。 **标准化**，因为无论在何处运行，您都需要标准的操作规则。 **安全**，好吧，因为您不希望任何不应该访问它的人这样做。并且**隔离**，因为您不希望该应用程序影响其他应用程序或受到其他应用程序的影响（例如，如果位于同一位置的应用程序崩溃）。隔离基本上起到保护的作用。此外，必须为应用程序提供资源，从 CPU 到存储再到内存。

### How It Helps

### 它如何帮助

The container runtime does all that. It launches apps in a standardized fashion across all environments and sets security boundaries. The latter is where some of these tools differ. Runtimes like CRI-O or gVisor have hardened their security boundaries. The runtime also sets resource limits for the container. Without it, the app could consume resources as needed, potentially taking resources away from other apps, so you always need to set limits.

容器运行时完成所有这些。它以标准化的方式在所有环境中启动应用程序并设置安全边界。后者是其中一些工具的不同之处。像 CRI-O 或 gVisor 这样的运行时已经强化了它们的安全边界。运行时还为容器设置资源限制。没有它，应用程序可能会根据需要消耗资源，可能会从其他应用程序中夺走资源，因此您始终需要设置限制。

### Technical 101

### 技术 101

Not all tools in this category are created equal. Containerd (part of the famous Docker product) and CRI-O are standard container runtime implementations. Then there are tools that expand the use of containers to other technologies, such as Kata which allows you to run containers as VMs. Others aim at solving a specific container-related problem such as gVisor which provides an additional security layer between containers and the OS.

并非此类别中的所有工具都是平等的。 Containerd（著名的 Docker 产品的一部分）和 CRI-O 是标准的容器运行时实现。还有一些工具可以将容器的使用扩展到其他技术，例如 Kata，它允许您将容器作为 VM 运行。其他人旨在解决特定的容器相关问题，例如 gVisor，它在容器和操作系统之间提供额外的安全层。

**Buzzwords**

**流行语**

**Popular Projects/Products**

**热门项目/产品**

- Container
- MicroVM

- 容器
- 微型虚拟机

- **Containerd**
- **CRI-O**
- **Kata**
- **gVisor**
- **Firecracker**

- **集装箱**
- **CRI-O**
- **卡塔**
- **gVisor**
- **鞭炮**

![Container runtime](https://cdn.thenewstack.io/media/2020/09/6d332514-screen-shot-2020-09-23-at-6.48.22-pm.png)

## Cloud Native Networking

## 云原生网络

### What It Is 

###  这是什么

Containers talk to each other and to the infrastructure layer through a cloud native network. [Distributed applications](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/) have multiple components that use the network for different purposes.Tools in this category overlay a virtual network on top of existing networks specifically for apps to communicate, referred to as an _overlay_ network.

容器通过云原生网络相互通信并与基础设施层通信。 [分布式应用程序](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/) 有多个组件，将网络用于不同目的。该类别中的工具覆盖了一个虚拟网络专门用于应用程序通信的现有网络，称为 _overlay_ 网络。

### Problem It Addresses

### 它解决的问题

While it’s common to refer to the code running in a container as an app, the reality is that most containers hold only a small specific set of functionalities of a larger application. Modern applications such as Netflix or Gmail are actually composed of a number of these smaller components each running in its own container. For all these independent pieces to function as a cohesive application, containers need to communicate with each other privately. Tools in this category provide that private communication network.

虽然通常将容器中运行的代码称为应用程序，但现实情况是，大多数容器仅包含大型应用程序的一小部分特定功能。 Netflix 或 Gmail 等现代应用程序实际上由许多这些较小的组件组成，每个组件都在自己的容器中运行。为了让所有这些独立的部分作为一个有凝聚力的应用程序运行，容器需要彼此私下通信。此类别中的工具提供该专用通信网络。

Additionally, messages exchanged between these containers may be private, sensitive, or extremely important. This leads to additional requirements such as providing isolation for the various components and the ability to inspect traffic to identify network issues. In some use cases, you may want to extend these networks and network policies (e.g. firewall and access rules) so your app can connect to virtual machines or services running externally to our container network.

此外，这些容器之间交换的消息可能是私密的、敏感的或极其重要的。这导致了额外的要求，例如为各种组件提供隔离以及检查流量以识别网络问题的能力。在某些用例中，您可能希望扩展这些网络和网络策略（例如防火墙和访问规则），以便您的应用程序可以连接到在我们的容器网络外部运行的虚拟机或服务。

### How It Helps

### 它如何帮助

Projects and products in this category use the CNCF project – Container Network Interface (CNI) to provide networking functionalities to containerized applications. Some tools, like Flannel, are rather minimalistic providing bare-bones connectivity to containers. Others, such as NSX-T provide a full software-defined networking layer creating an isolated virtual network for every Kubernetes namespace.

此类别中的项目和产品使用 CNCF 项目 – 容器网络接口 (CNI) 为容器化应用程序提供网络功能。一些工具，如 Flannel，相当简约，提供与容器的基本连接。其他的，如 NSX-T 提供了一个完整的软件定义网络层，为每个 Kubernetes 命名空间创建一个隔离的虚拟网络。

At a minimum, a container network needs to assign IP addresses to pods (that’s where containerized apps run in Kubernetes), that allows other processes to access it.

至少，容器网络需要为 Pod 分配 IP 地址（这是容器化应用程序在 Kubernetes 中运行的地方），以允许其他进程访问它。

### Technical 101

### 技术 101

Similar to storage, the variety and innovation in this space is largely made possible by the CNCF project CNI (Container Networking Interface) which standardizes how network layers provide functionalities to pods.Selecting the right container network for your Kubernetes environment is critical and you've got a number of tools to choose from. Weave Net, Antrea, Calico, and Flannel all provide effective open source networking layers. Their functionalities vary widely and your choice should be ultimately driven by your specific needs.

与存储类似，该领域的多样性和创新在很大程度上是由 CNCF 项目 CNI（容器网络接口）实现的，该项目标准化了网络层如何为 pod 提供功能。为您的 Kubernetes 环境选择正确的容器网络至关重要，您已经有许多工具可供选择。 Weave Net、Antrea、Calico 和 Flannel 都提供了有效的开源网络层。它们的功能差异很大，您的选择最终应由您的特定需求决定。

Additionally,there are many vendors ready to support and extend your Kubernetes networks with Software Defined Networking (SDN) tools that allow you to gain additional insights into network traffic, enforce network policies, and even extend your container networks and policies to your broader datacenter.

此外，有许多供应商准备使用软件定义网络 (SDN) 工具支持和扩展您的 Kubernetes 网络，这些工具使您能够深入了解网络流量、实施网络策略，甚至将容器网络和策略扩展到更广泛的数据中心。

**Buzzwords**

**流行语**

**Popular Projects/Products**

**热门项目/产品**

- SDN
- Network Overlay
- CNI

- SDN
- 网络覆盖
- CNI

- **Calico**
- **Weave Net**
- **Flannel**
- **Antrea**
- **NSX-T**

- **印花布**
- **织网**
- **法兰绒**
- **安特里亚**
- **NSX-T**

![cloud native network](https://cdn.thenewstack.io/media/2020/09/01e7d965-screen-shot-2020-09-23-at-6.51.02-pm.png)

This concludes our overview of the runtime layer which provides all the tools containers need to run in a cloud native environment. From storage that gives apps easy and fast access to data needed to run reliably, to the container runtime which executes the application code, to the network over which containerized apps communicate. In our next article, we’ll focus on the orchestration and management layer which deals with how all these containerized apps are managed as a group.

我们对运行时层的概述到此结束，它提供了容器在云原生环境中运行所需的所有工具。从使应用程序能够轻松快速地访问可靠运行所需的数据的存储，到执行应用程序代码的容器运行时，再到容器化应用程序通过其进行通信的网络。在我们的下一篇文章中，我们将关注编排和管理层，它处理如何将所有这些容器化应用程序作为一个组进行管理。

_As always, a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._

_一如既往，非常感谢来自 CNCF 的 [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/)，他非常友好地审阅了这篇文章，以确保其准确无误。_

The Cloud Native Computing Foundation and VMware are sponsors of The New Stack.

云原生计算基金会和 VMware 是 The New Stack 的赞助商。

Feature Image by [Candid\_Shots](https://pixabay.com/users/Candid_Shots-11873433/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=5582775) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=5582775).

[Candid\_Shots](https://pixabay.com/users/Candid_Shots-11873433/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=5582775) 来自 https://pixabay. utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=5582775)。

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: MADE, Docker, Famous. 

The New Stack 是 Insight Partners 的全资子公司。 TNS 所有者 Insight Partners 是以下公司的投资者：MADE、Docker、Famous。

