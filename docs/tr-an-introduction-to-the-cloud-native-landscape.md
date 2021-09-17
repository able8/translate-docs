# An Introduction to the Cloud Native Landscape

# 云原生景观简介

#### 21 Jul 2020 9:52am,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini")

If you’ve researched cloud native applications and technologies, you’ve probably come across the [Cloud Native Computing Foundation (CNCF)](https://landscape.cncf.io/) cloud native landscape map. Unsurprisingly, the sheer scale of it can be overwhelming. So many categories and so many technologies. How do you make sense of it?

如果您研究过云原生应用程序和技术，您可能会遇到过[云原生计算基金会 (CNCF)](https://landscape.cncf.io/) 云原生景观图。不出所料，它的庞大规模可能是压倒性的。这么多类别和这么多技术。你怎么理解它？

As with anything else, if you break it down and analyze it one piece at the time, you’ll find it’s not that complex and makes a lot of sense. In fact, the map is neatly organized by functionality, and once you understand what each category represents, navigating it becomes a lot easier.

与其他任何事情一样，如果您将其分解并一次分析，您会发现它并不那么复杂并且很有意义。事实上，地图按功能整齐地组织起来，一旦您了解每个类别代表什么，导航就会变得容易得多。

In this article, the first in a series, we’ll break this mammoth landscape down and provide a high-level overview of the entire landscape, its layers, columns and categories. In follow up articles, we’ll zoom into each layer and column and provide more details on what each category is, what problem it solves, and how.

在本系列文章的第一篇文章中，我们将分解这个庞大的景观，并提供整个景观、其层、列和类别的高级概述。在后续文章中，我们将放大每一层和每一列，并提供有关每个类别是什么、它解决什么问题以及如何解决的更多详细信息。

![CNCF landscape](https://cdn.thenewstack.io/media/2020/07/6b2f1a36-screen-shot-2020-07-17-at-8.17.23-am.png)

## The Four Layers of the Cloud Native Landscape

## 云原生景观的四层

First, let’s strip all individual technologies from the landscape and look at the categories. There are different “rows” reflecting architectural layers each with its own set of subcategories. In the first layer, you have tools to provision infrastructure, that’s your foundation. Then you start adding tooling needed to run and manage apps such as the runtime and orchestration layer. At the very top you have tools to define and develop your application, such as databases, image building, and CI/CD tools (we’ll discuss each of these below).

首先，让我们从景观中剥离所有单独的技术并查看类别。有不同的“行”反映了建筑层，每一个都有自己的子类别集。在第一层，您拥有用于配置基础设施的工具，这是您的基础。然后开始添加运行和管理应用程序所需的工具，例如运行时和编排层。在最顶端，您拥有定义和开发应用程序的工具，例如数据库、图像构建和 CI/CD 工具（我们将在下面讨论这些工具中的每一个）。

![CNCF landscape categories](https://cdn.thenewstack.io/media/2020/07/794834eb-screen-shot-2020-07-17-at-2.39.02-pm.png)

For now, what you should remember is that the landscape starts with the infrastructure and, with each layer, moves closer to the actual app. That’s what these layers represent (we’ll address the two “columns” running across those layers later). Let’s explore each layer at a time, starting with the bottom.

现在，你应该记住的是，景观从基础设施开始，随着每一层，都更接近实际的应用程序。这就是这些层所代表的内容（我们稍后将讨论跨越这些层的两个“列”）。让我们一次探索每一层，从底部开始。

### 1\. The Provisioning Layer

### 1\.供应层

Provisioning refers to the tools involved in creating and hardening the foundation on which cloud native applications are built. It covers everything from automating the creation, management, and configuration of infrastructure to scanning, signing and storing container images. Provisioning even extends into the security space by providing tools that allow you to set and enforce policies, build authentication and authorization into your apps and platforms, and handle secrets distribution.

配置是指创建和强化构建云原生应用程序的基础所涉及的工具。它涵盖了从自动化基础设施的创建、管理和配置到扫描、签名和存储容器映像的所有内容。通过提供允许您设置和实施策略、在应用程序和平台中构建身份验证和授权以及处理机密分发的工具，配置甚至扩展到安全领域。

In the provisioning layer, you’ll find:

在配置层中，您会发现：

- **Automation and configuration tooling** to help engineers build computing environments without human intervention.
- **Container registries** store executable files of the apps.
- **Security and** **compliance** frameworks address different security areas.
- **Key management** solutions help with encryption to ensure only authorized users have access to the application.

- **自动化和配置工具**，帮助工程师构建无需人工干预的计算环境。
- **容器注册表** 存储应用程序的可执行文件。
- **安全和** **合规性** 框架针对不同的安全领域。
- **密钥管理**解决方案有助于加密，以确保只有授权用户才能访问应用程序。

These tools allow engineers to codify all infrastructure specifics, so that the system can spin new environments up and down as needed, ensuring they are consistent and secure.

这些工具允许工程师对所有基础设施细节进行编码，以便系统可以根据需要启动和关闭新环境，确保它们的一致性和安全性。

### 2\. The Runtime Layer

### 2. 运行时层

Next, is the runtime layer. Runtime is one of those terms that can be confusing. Like many terms in IT, there is no strict definition and it can be used differently, depending on the context. In a narrow sense, runtime is a sandbox on a specific machine prepared to run an app — the bare minimum an app needs. In the widest of senses, runtime is any tool the app needs to run.

接下来是运行时层。运行时是可能令人困惑的术语之一。与 IT 中的许多术语一样，没有严格的定义，可以根据上下文使用不同的术语。从狭义上讲，运行时是准备运行应用程序的特定机器上的沙箱——应用程序所需的最低限度。在最广泛的意义上，运行时是应用程序需要运行的任何工具。

In the CNCF cloud native landscape, runtime is defined somewhere in between focusing on the components that matter for the containerized apps in particular: what they need to run, remember, and communicate. They include:

在 CNCF 云原生环境中，运行时的定义介于两者之间，重点关注对容器化应用程序特别重要的组件：它们需要运行、记住和通信的内容。它们包括：

- **Cloud native storage** provides virtualized disks or persistence for containerized apps.
- **Container runtime** delivers the constraint, resource, and security considerations for containers and executes the files with the codified app. 

- **云原生存储**为容器化应用程序提供虚拟化磁盘或持久性。
- **容器运行时** 为容器提供约束、资源和安全注意事项，并使用编码的应用程序执行文件。

- **Cloud native networking**, the network over which nodes (machines or processes) of a [distributed system](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/) are connected and communicate.

- **云原生网络**，[分布式系统](https://thenewstack.io/primer-distributed-systems-and-cloud-native-computing/)的节点（机器或进程)所在的网络连接和通信。

### 3\. The Orchestration and Management Layer

### 3. 编排和管理层

Once you automate infrastructure provisioning following security and compliance standards (provisioning layer) and set up the tools the app needs to run (runtime layer), engineers must figure out how to orchestrate and manage their apps. The orchestration and management layer deals with how all containerized services (app components) are managed as a group. They need to identify other services, communicate with one another, and coordinate. Inherently scalable, cloud native apps rely on automation and resilience, enabled by this layer.

一旦您按照安全性和合规性标准（配置层）自动化基础设施配置并设置应用程序需要运行的工具（运行时层），工程师必须弄清楚如何编排和管理他们的应用程序。编排和管理层处理如何将所有容器化服务（应用程序组件）作为一个组进行管理。他们需要识别其他服务、相互通信和协调。本质上可扩展的云原生应用程序依赖于由这一层实现的自动化和弹性。

In this layer you’ll find:

在这一层你会发现：

- **Orchestration and scheduling** to deploy and manage container clusters ensuring they are resilient, loosely-coupled, and scalable. In fact, the orchestration tool, in most cases [Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-care/), is what makes a cluster by managing containers and the operating environment
- **Coordination and service discovery** so services (app components) can locate and communicate with one another.
- **Remote procedure call (RPC)**, a technique enabling a service on one node to communicate with a service on a different node connected through a network.
- **Service proxy** is an intermediary placed between services through which they communicate. The sole purpose of the proxy is to exert more control over service communication, it doesn’t add anything to the communication itself. These proxies are crucial to service meshes mentioned below.
- **API gateway**, an abstraction layer through which external applications can communicate.
- **Service mesh** is similar to the API gateway in the sense that it’s a dedicated infrastructure layer through which apps communicate, but it provides policy-driven _internal_ service-to-service communication. Additionally, it may include everything from traffic encryption to service discovery, to application observability.

- **编排和调度**以部署和管理容器集群，确保它们具有弹性、松散耦合和可扩展性。其实编排工具，在大多数情况下[Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-care/)，是通过管理容器和运行环境来构建集群的
- **协调和服务发现**，因此服务（应用程序组件）可以相互定位和通信。
- **远程过程调用 (RPC)**，一种使一个节点上的服务能够与通过网络连接的不同节点上的服务进行通信的技术。
- **服务代理**是放置在服务之间的中介，通过它们进行通信。代理的唯一目的是对服务通信施加更多控制，它不会向通信本身添加任何内容。这些代理对于下面提到的服务网格至关重要。
- **API 网关**，一个抽象层，外部应用程序可以通过它进行通信。
- **服务网格**类似于 API 网关，因为它是应用程序通信的专用基础设施层，但它提供策略驱动的 _内部_ 服务到服务通信。此外，它可能包括从流量加密到服务发现，再到应用程序可观察性的所有内容。

### 4\. The Application Definition and Development Layer

### 4. 应用定义和开发层

Now let’s move to the top layer. As the name suggests, the application definition and development layer focus on the tools that enable engineers to build apps and allow them to function. Everything discussed above was related to building a reliable, secure environment and providing all needed app dependencies.

现在让我们移动到顶层。顾名思义，应用程序定义和开发层专注于使工程师能够构建应用程序并使其运行的工具。上面讨论的所有内容都与构建可靠、安全的环境和提供所有需要的应用程序依赖项有关。

Under this category you’ll see:

在此类别下，您将看到：

- **Databases** enabling apps to collect data in an organized manner.
- **Streaming and messaging** enable apps to send and receive messages (events and streams). It’s not a networking layer, but rather a tool to queue and process messages.
- **Application definition and image build** are services that help configure, maintain, and run container images (the executable files of an app).
- **Continuous integration and delivery (CI/CD)** allow developers to automatically test that their code works with the codebase (the rest of the app) and, if their team is mature enough, even automate deployment into production.

- **数据库** 使应用程序能够以有组织的方式收集数据。
- **流和消息传递**使应用程序能够发送和接收消息（事件和流）。它不是网络层，而是用于排队和处理消息的工具。
- **应用程序定义和映像构建** 是帮助配置、维护和运行容器映像（应用程序的可执行文件）的服务。
- **持续集成和交付 (CI/CD)** 允许开发人员自动测试他们的代码是否适用于代码库（应用程序的其余部分），如果他们的团队足够成熟，甚至可以自动部署到生产中。

## Tools Running Across All Layers

## 跨所有层运行的工具

Going back to the category overview, we’ll explore the two columns running across all layers. Observability and analysis are tools that monitor all layers. Platforms, on the other hand, bundle multiple technologies within these layers into one solution, including observability and analysis.

回到类别概览，我们将探索跨越所有层的两列。可观察性和分析是监控所有层的工具。另一方面，平台将这些层中的多种技术捆绑到一个解决方案中，包括可观察性和分析。

### ![cncf landscape rows](https://cdn.thenewstack.io/media/2020/07/50c10db4-screen-shot-2020-07-17-at-2.39.23-pm.png)

### Observability and Analysis 

### 可观察性和分析

To limit service disruption and help drive down MRRT (meantime to resolution), you’ll need to monitor and analyze every aspect of your application so any anomaly gets detected and rectified right away. Failures _will_ occur in complex environments and these tools help make them less impactful by helping identify and resolve failures as quickly as possible. Since this category runs across and monitors all layers, it’s on the side and not embedded in a specific layer.

为了限制服务中断并帮助降低 MRRT（平均解决时间），您需要监控和分析应用程序的各个方面，以便立即检测和纠正任何异常。故障_将_发生在复杂的环境中，这些工具通过帮助尽快识别和解决故障来帮助降低它们的影响。由于此类别贯穿并监控所有层，因此它位于侧面，而不是嵌入到特定层中。

Here you’ll find:

在这里你会发现：

- **Logging** tools to collect event logs (info about processes).
- **Monitoring** solutions to collect metrics (numerical system parameters, such as RAM availability).
- **Tracing** goes one step further than monitoring and monitors the propagation of user requests. This is relevant in the context of service meshes.
- **Chaos engineering** are tools to test software in production to identify weaknesses and fix them before they impact service delivery.

- **Logging** 工具来收集事件日志（有关进程的信息）。
- **监控**解决方案以收集指标（数字系统参数，例如 RAM 可用性）。
- **跟踪**比监视和监视用户请求的传播更进一步。这在服务网格的上下文中是相关的。
- **混沌工程**是在生产中测试软件以识别弱点并在它们影响服务交付之前修复它们的工具。

### Platforms

### 平台

As we’ve seen, each of these modules solves a particular problem. Storage alone does not provide all you need to manage your app. You’ll need an orchestration tool, container runtime, service discovery, networking, API gateway, etc. Covering multiple layers, platforms bundle different tools together solving a larger problem.

正如我们所见，这些模块中的每一个都解决了一个特定的问题。仅存储并不能提供管理应用程序所需的一切。您需要一个编排工具、容器运行时、服务发现、网络、API 网关等。平台覆盖多个层，将不同的工具捆绑在一起解决一个更大的问题。

Configuring and fine-tuning different modules so they are reliable and secure and ensuring all the technologies it leverages are updated and vulnerabilities patched is no easy task. With platforms, users don’t have to worry about these details — a real value add.

配置和微调不同的模块，使其可靠和安全，并确保它所利用的所有技术都得到更新并修补漏洞并非易事。有了平台，用户就不必担心这些细节——真正的增值。

You’ll probably notice, the categories all revolve around Kubernetes. That’s because Kubernetes, while only one piece of the puzzle, is at the core of the cloud native stack. The CNCF, by the way, was created with Kubernetes as its first seeding project; all other projects followed later.

您可能会注意到，所有类别都围绕 Kubernetes。这是因为 Kubernetes 虽然只是拼图的一部分，但却是云原生堆栈的核心。顺便说一下，CNCF 是用 Kubernetes 作为第一个种子项目创建的；所有其他项目随后跟进。

Platforms can be categorized in four groups:

平台可以分为四组：

- **Kubernetes distributions** take the unmodified, open source code (although some modify it) and add additional features their market needs around it.
- **Hosted Kubernetes** (aka managed Kubernetes) is similar to a distribution but it’s managed by your provider on their or on your own infrastructure.
- **Kubernetes installers** are exactly that, they automate the installation and configuration process of Kubernetes.
- **PaaS / container services** are similar to hosted Kubernetes, but include a broad set of application deployment tools (generally a subset from the cloud native landscape).

- **Kubernetes 发行版**采用未经修改的开源代码（尽管有些对其进行了修改）并围绕其添加了市场需要的附加功能。
- **托管 Kubernetes**（又名托管 Kubernetes）类似于发行版，但由您的提供商在他们或您自己的基础架构上管理。
- **Kubernetes 安装程序** 就是这样，它们自动执行 Kubernetes 的安装和配置过程。
- **PaaS / 容器服务**类似于托管 Kubernetes，但包括一组广泛的应用程序部署工具（通常是云原生环境的一个子集）。

## Conclusion

##  结论

In each category, there are different tools aimed at solving the same or similar problems. Some are pre-cloud native technologies adapted to the new reality, while others are completely new. Differences lie in their implementation and design approaches. There is no perfect technology that checks all the boxes. In most cases technology is limited by design and architectural choices — there is always a tradeoff.

在每个类别中，都有不同的工具旨在解决相同或相似的问题。有些是适应新现实的云原生技术，而有些则是全新的。不同之处在于它们的实现和设计方法。没有完美的技术可以检查所有的框。在大多数情况下，技术受到设计和架构选择的限制——总有一个权衡。

When selecting the stack, engineers must carefully consider each capability and tradeoff to identify the best option for their use case. While this brings additional complexity, it’s never been more feasible to choose a data storage, infrastructure management, messaging system, etc. that best fits the application’s needs. Architecting systems today is a lot easier than in a pre-cloud native world. And, if architected appropriately, cloud native technologies offer powerful and much-needed flexibility. In today’s fast-changing technology ecosystem, that is likely one of the most important capabilities.

在选择堆栈时，工程师必须仔细考虑每个功能和权衡，以确定适合其用例的最佳选项。虽然这会带来额外的复杂性，但选择最适合应用程序需求的数据存储、基础设施管理、消息系统等从未如此可行。今天的系统架构比在云原生世界中容易得多。而且，如果架构适当，云原生技术可提供强大且急需的灵活性。在当今瞬息万变的技术生态系统中，这可能是最重要的能力之一。

We hope this quick overview was helpful. Stay tuned for our follow up articles to learn more about each layer and column! 

我们希望这个快速概述对您有所帮助。请继续关注我们的后续文章，以了解有关每一层和每一列的更多信息！

**_As always, thanks to [Oleg Chunihkin](https://www.linkedin.com/in/olegch/) for all his input and also to [Jason Morgan](https://www.linkedin.com/in/jasonmorgan2/), my co-author for the upcoming more detailed landscape articles (very excited about that!). And a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._**

**_一如既往，感谢 [Oleg Chunihkin](https://www.linkedin.com/in/olegch/) 的所有投入以及 [Jason Morgan](https://www.linkedin.com/in)，我即将发表的更详细的景观文章的合著者（对此感到非常兴奋！)。非常特别感谢来自 CNCF 的 [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/)，他非常友好地审阅了文章以确保一切准确。_**



At this time, The New Stack does not allow comments directly on this website. We invite all readers who wish to discuss a story to visit us on [Twitter](https://twitter.com/thenewstack) or [Facebook](https://www.facebook.com/thenewstack/). We also welcome your news tips and feedback via email: [feedback@thenewstack.io](mailto:feedback@thenewstack.io).

目前，The New Stack 不允许直接在本网站上发表评论。我们邀请所有希望讨论故事的读者在 [Twitter](https://twitter.com/thenewstack) 或 [Facebook](https://www.facebook.com/thenewstack/)上访问我们。我们也欢迎您通过电子邮件提供新闻提示和反馈：[feedback@thenewstack.io](mailto:feedback@thenewstack.io)。

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Real. 

The New Stack 是 Insight Partners 的全资子公司。 TNS 所有者 Insight Partners 是以下公司的投资者：Real。

