# How we use Kubernetes at Asana

# 我们如何在 Asana 使用 Kubernetes

By
[Tony Liang](https://blog.asana.com/author/tonyliangasana-com/ "Posts by Tony Liang")
Feb 11, 2021

经过
[Tony Liang](https://blog.asana.com/author/tonyliangasana-com/ “Tony Liang 的帖子”)
2021 年 2 月 11 日

At Asana, we use Kubernetes to deploy and manage services independently from our monolith infrastructure. We encountered a few pain points when initially using Kubernetes, and built a framework to standardize the creation and maintenance of Kubernetes applications, aptly named KubeApps. Over the last two years, the Infrastructure Platform team has been making improvements to the KubeApps framework to make deploying services easier at Asana. In this post, we’ll explore the problems we aimed to solve in Kubernetes and the ways we did that with the KubeApp framework.

在 Asana，我们使用 Kubernetes 独立于我们的整体基础架构部署和管理服务。我们在最初使用 Kubernetes 时遇到了一些痛点，并构建了一个框架来规范 Kubernetes 应用程序的创建和维护，恰如其分地命名为 KubeApps。在过去的两年中，基础设施平台团队一直在改进 KubeApps 框架，以便在 Asana 更轻松地部署服务。在这篇文章中，我们将探讨我们旨在在 Kubernetes 中解决的问题以及我们使用 KubeApp 框架解决问题的方法。

### Background

###  背景

In an early blog post about Asana's [legacy deployment system](https://blog.asana.com/2017/06/asana-server-deployment-pragmatic-approach-maintaining-legacy-deployment-system/), we discussed how our core infrastructure runs as a monolith of AWS EC2 instances with custom configuration scripts for each service. This system was not scalable: Pushing new code to the monolith required all hosts to be reconfigured; adding new services required a new set of custom configuration scripts that were difficult to maintain and carried a risk of adding instability to deployments to the monolith.

在一篇关于 Asana 的 [遗留部署系统] (https://blog.asana.com/2017/06/asana-server-deployment-pragmatic-approach-maintaining-legacy-deployment-system/) 的早期博客文章中，我们讨论了如何我们的核心基础设施作为 AWS EC2 实例的整体运行，每个服务都有自定义配置脚本。该系统不可扩展：将新代码推送到单体应用需要重新配置所有主机；添加新服务需要一组新的自定义配置脚本，这些脚本难以维护，并且存在增加整体部署不稳定的风险。

When we considered how we wanted to build out the infrastructure behind [Luna2](https://blog.asana.com/2017/08/performance-asana-app-rewrite/), we decided to use Kubernetes to enable us to deploy and manage services independently from the monolith. The container orchestration solution worked well for us, and our broader Engineering team expressed enough interest in using Kubernetes for other systems at Asana that we generalized these tools to allow other teams to build and deploy containerized services.

当我们考虑如何构建 [Luna2](https://blog.asana.com/2017/08/performance-asana-app-rewrite/) 背后的基础设施时，我们决定使用 Kubernetes 来使我们能够部署并独立于单体管理服务。容器编排解决方案对我们来说效果很好，我们更广泛的工程团队对将 Kubernetes 用于 Asana 的其他系统表示了足够的兴趣，因此我们将这些工具推广到允许其他团队构建和部署容器化服务。

### Kubernetes at Asana

### Kubernetes 在 Asana

Using Kubernetes has made managing services much easier because we are able to delineate clear boundaries between processes and the environments where they run. As a result, application developers are able to build and update services without worrying about negatively impacting other existing services.

使用 Kubernetes 使管理服务变得更加容易，因为我们能够在进程和它们运行的环境之间划定清晰的界限。因此，应用程序开发人员能够构建和更新服务，而不必担心会对其他现有服务产生负面影响。

While Kubernetes helped us solve challenges with scaling out our infrastructure beyond the EC2 monolith, we still needed to handle some problems that Kubernetes did not solve for us out of the box at the time.[1]

虽然 Kubernetes 帮助我们解决了将基础设施扩展到 EC2 单体之外的挑战，但我们仍然需要处理一些当时 Kubernetes 没有立即为我们解决的问题。[1]

- AWS resource management is beyond the scope of Kubernetes. Developers would need to provide pre-created ASGs and provision them to run Docker containers as well as any other resources that may be required for the application to run.
- Instrumenting common tooling like metrics collection and secrets permissions across all services was tedious and we would find ourselves repeating work like setting up ELBs and managing configuration groups for each application.
- Continuous delivery is not built into Kubernetes. At the Kubernetes abstraction layer, the application logic is expected to be packaged into images readily available in a container registry.

- AWS 资源管理超出了 Kubernetes 的范围。开发人员需要提供预先创建的 ASG 并配置它们以运行 Docker 容器以及应用程序运行可能需要的任何其他资源。
- 在所有服务中检测指标收集和机密权限等通用工具是乏味的，我们会发现自己重复工作，例如为每个应用程序设置 ELB 和管理配置组。
- Kubernetes 未内置持续交付。在 Kubernetes 抽象层，应用程序逻辑预计会被打包到容器注册表中随时可用的图像中。

In order to solve these pain points, we built a framework to standardize the creation and maintenance of Kubernetes applications, aptly named KubeApps. The KubeApp framework was initially designed to handle all aspects of deployment from gathering the necessary resources to updating DNS during deployments and providing smooth transitions between code versions. Each KubeApp is defined by a set of hardware requirements, pod definitions and service specifications.

为了解决这些痛点，我们构建了一个框架来规范Kubernetes应用的创建和维护，恰如其分地命名为KubeApps。 KubeApp 框架最初旨在处理部署的所有方面，从收集必要的资源到在部署期间更新 DNS 以及提供代码版本之间的平滑转换。每个 KubeApp 都由一组硬件要求、pod 定义和服务规范定义。

### Configuration as code

### 配置即代码

All of our KubeApp configurations exist as Python code. This allows for programmatically generated configurations and ensures consistency via shared libraries. These configurations allow us to declare and configure external resources needed for the KubeApp.

我们所有的 KubeApp 配置都以 Python 代码的形式存在。这允许以编程方式生成配置并通过共享库确保一致性。这些配置允许我们声明和配置 KubeApp 所需的外部资源。

In the image below, we have the specifications for a sample webapp and how it can be represented in our KubeApp configuration system. The deployment will make requests for hardware, ELB and image builds. We provide default specifications to handle DNS configurations and include pod definitions to handle metrics and logging.

在下图中，我们有示例 Web 应用程序的规范以及它如何在我们的 KubeApp 配置系统中表示。部署将请求硬件、ELB 和镜像构建。我们提供默认规范来处理 DNS 配置，并包括 pod 定义来处理指标和日志记录。

![](https://blog.asana.com/wp-content/post-images/sample-kubeapp-fixed-1024x492.png)

### Docker images with continuous deployment 

### 持续部署的 Docker 镜像

Application code is packaged into Docker images using either [Bazel Docker images](https://github.com/bazelbuild/rules_docker) or via traditional Dockerfiles. Our KubeApps framework will build and push these images as part of our CD pipeline, so images for any service will be readily available on our container registry.

应用程序代码使用 [Bazel Docker 镜像](https://github.com/bazelbuild/rules_docker) 或通过传统的 Dockerfile 打包到 Docker 镜像中。我们的 KubeApps 框架将构建和推送这些镜像作为我们 CD 管道的一部分，因此任何服务的镜像都可以在我们的容器注册表中随时可用。

We use [Bazel](https://bazel.build/) to build and manage dependencies for most of our code at Asana. It’s particularly useful with containerized applications because it enforces explicit dependency declarations and packages third party libraries into built executables. This provides a deterministic output when setting up container environments for each application.

我们使用 [Bazel](https://bazel.build/) 为我们在 Asana 的大部分代码构建和管理依赖项。它对容器化应用程序特别有用，因为它强制执行显式依赖声明并将第三方库打包到构建的可执行文件中。这在为每个应用程序设置容器环境时提供了确定性的输出。

### One Cluster per KubeApp

### 每个 KubeApp 一个集群

A decision that many teams come across when building on Kubernetes is how to divide services up between clusters[2]. Many services deployed to a single Kubernetes cluster is cost effective and allows for efficient administration because the services share master node resources and operational work only needs to happen once (i.e. Kubernetes version upgrades, deployments). However, the single cluster model does not provide hard multitenancy and lends itself to being a single point of failure. There are also scaling limitations of a single Kubernetes cluster, so a split will need to be considered once a deployment grows to a certain size.

许多团队在构建 Kubernetes 时遇到的一个决定是如何在集群之间划分服务[2]。部署到单个 Kubernetes 集群的许多服务具有成本效益并允许高效管理，因为这些服务共享主节点资源并且操作工作只需要发生一次（即 Kubernetes 版本升级、部署）。但是，单集群模型不提供硬多租户，并且容易成为单点故障。单个 Kubernetes 集群也存在扩展限制，因此一旦部署增长到一定规模，就需要考虑拆分。

At Asana, each KubeApp is deployed in its own Kubernetes cluster via AWS EKS. This approach gives us strong guarantees on application security and resilience. Since each cluster is responsible for one service, we don’t need to worry about resource contention between services and the impact of a cluster failure is limited to a single service.

在 Asana，每个 KubeApp 都通过 AWS EKS 部署在自己的 Kubernetes 集群中。这种方法为我们提供了对应用程序安全性和弹性的有力保证。由于每个集群负责一个服务，我们不需要担心服务之间的资源争用，集群故障的影响仅限于单个服务。

Managing multiple clusters can be tricky because the default tooling for Kubernetes is only able to interface with one cluster at a time. However, our team has built tooling in our KubeApps framework to manage multiple clusters at once. We’ve also found that this model empowers individual KubeApp owners to take on cluster management work independently (i.e. upgrading nodes, scaling deployment sizes, etc).

管理多个集群可能很棘手，因为 Kubernetes 的默认工具一次只能与一个集群交互。但是，我们的团队在 KubeApps 框架中构建了工具来同时管理多个集群。我们还发现，该模型使各个 KubeApp 所有者能够独立承担集群管理工作（即升级节点、扩展部署规模等）。

### KubeApp Deployment Workflow

### KubeApp 部署工作流程

![](https://lh5.googleusercontent.com/leyZ7uA7vja_BRj_XrcY5Xr65H9grKvfekSSQwe_wc2fuy4AL-0Loepb_33zUq7TVBW39H7zPKI_rg4KPmromyyHlX4YTrYzrPpbbP55AY-4tpqQJYVNaPav_OSRmrmo3rRqR7s)

KubeApp deployments are driven through a central management hub that we named “kubecontrol” which can be configured to run updates automatically via crons or manually by developers. The steps that happen during a KubeApp deployment are as follows:

KubeApp 部署是通过我们命名为“kubecontrol”的中央管理中心驱动的，该中心可以配置为通过 cron 自动运行更新或由开发人员手动运行。 KubeApp 部署过程中发生的步骤如下：

1. A KubeApp update or create is triggered via the command line on kubecontrol
2. From the application specs, we request the set of resources required for the KubeApp (ASGs, spot instances, etc) and a new EKS cluster is created.
3. We make a request to our image builder service to compile the docker image at the given code version. The image builder will compile the code for the KubeApp and commit the image to ECR (Elastic Container Registry) if it does not already exist there.
4. Once all the required resources are built, we hand off the component specifications to the Kubernetes cluster to pull required docker containers from ECR and deploy them onto the nodes.

1. 通过 kubecontrol 上的命令行触发 KubeApp 更新或创建
2. 根据应用规范，我们请求 KubeApp 所需的资源集（ASG、spot 实例等）并创建一个新的 EKS 集群。
3. 我们向镜像构建器服务发出请求，以编译给定代码版本的 docker 镜像。镜像构建器将为 KubeApp 编译代码并将镜像提交到 ECR（弹性容器注册表）（如果那里尚不存在）。
4. 构建好所有所需资源后，我们将组件规范交给 Kubernetes 集群，以从 ECR 中拉取所需的 docker 容器并将它们部署到节点上。

Full updates of KubeApps are blue/green deployments and require a new EKS cluster AWS resources to be launched and configured. Once the newly launched KubeApp is verified to be in a working state, we switch load to the new cluster and tear down the old one. KubeApps also have the option for a rolling-update which will just update images on a running cluster. This allows for quick seamless transitions between code versions without needing to spin up an entirely new cluster.

KubeApps 的完整更新是蓝/绿部署，需要启动和配置新的 EKS 集群 AWS 资源。一旦验证新启动的 KubeApp 处于工作状态，我们将负载切换到新集群并拆除旧集群。 KubeApps 还提供滚动更新选项，它只会更新正在运行的集群上的图像。这允许在代码版本之间快速无缝转换，而无需启动全新的集群。

### KubeApp Management Console 

### KubeApp 管理控制台

Until recently, the only way for a developer to directly monitor or manage a KubeApp was to ssh into kubecontrol and interface with their app via the CLI. Information about deployments was not easily searchable so users would need to search through logs to figure out when a specific version of code was deployed to a KubeApp. In order to provide more clarity and observability to KubeApp users, we’ve built out a KubeApp Management Console (KMC) that would be responsible for recording historical information about past deployments. Eventually, we would like to use this interface to provide a centralized web-based interface for users to interact with KubeApps.

直到最近，开发人员直接监控或管理 KubeApp 的唯一方法是通过 ssh 进入 kubecontrol 并通过 CLI 与他们的应用程序交互。有关部署的信息不容易搜索，因此用户需要搜索日志来确定特定版本的代码何时部署到 KubeApp。为了向 KubeApp 用户提供更多清晰度和可观察性，我们构建了一个 KubeApp 管理控制台 (KMC)，负责记录有关过去部署的历史信息。最终，我们希望使用这个接口为用户提供一个集中的基于 Web 的接口，以便用户与 KubeApps 进行交互。

### State of KubeApps today and what’s next

### KubeApps 今天的状态以及接下来会发生什么

We currently have over 60 KubeApps running at Asana that support a wide variety of workloads ranging from [release management](https://blog.asana.com/2021/01/asana-engineering-ships-web-application-releases/) to [distributed caching](https://blog.asana.com/2020/09/worldstore-distributed-caching-reactivity-part-1/#close). We still maintain the monolith of EC2 instances, but we are in the process of reimagining these services as containerized processes running in KubeApps.

我们目前有 60 多个 KubeApps 在 Asana 上运行，支持各种工作负载，包括 [发布管理](https://blog.asana.com/2021/01/asana-engineering-ships-web-application-releases/)到[分布式缓存](https://blog.asana.com/2020/09/worldstore-distributed-caching-reactivity-part-1/#close)。我们仍然维护 EC2 实例的整体，但我们正在将这些服务重新设想为在 KubeApps 中运行的容器化进程。

The Infrastructure Platform team will continue to iterate and build new functionalities on the KubeApps framework. In the future, we plan to extend support for more types of architectures (ARM64) and infrastructure providers (AWS Fargate). We also plan to build tools to support a better development experience on KubeApps by making them launchable in sandboxed or local environments. These changes, among others, will enable the KubeApps framework to be extensible for any workload our engineers may need.

基础设施平台团队将继续在 KubeApps 框架上迭代和构建新功能。未来，我们计划扩展对更多类型架构 (ARM64) 和基础设施提供商 (AWS Fargate) 的支持。我们还计划构建工具来支持更好的 KubeApps 开发体验，使它们可以在沙盒或本地环境中启动。除其他外，这些更改将使 KubeApps 框架能够针对我们的工程师可能需要的任何工作负载进行扩展。

[1] Kubernetes does have a [Cloud Controller Manager](https://kubernetes.io/docs/concepts/architecture/cloud-controller/) that can manage node objects, configure routes between containers, and integrate with cloud infrastructure components.

[1] Kubernetes 确实有一个[云控制器管理器](https://kubernetes.io/docs/concepts/architecture/cloud-controller/)，可以管理节点对象，配置容器之间的路由，并与云基础设施组件集成。

[2] Some discussion about the tradeoffs between many clusters vs few clusters are available in this article from learnk8s.io: [Architecting Kubernetes clusters — how many should you have?](https://learnk8s.io/how-many-clusters)

[2] 这篇来自 learnk8s.io 的文章提供了一些关于多集群与少数集群之间权衡的讨论：[架构 Kubernetes 集群——你应该拥有多少？](https://learnk8s.io/how-many-clusters)

Special thanks to

特别感谢

Eldar Bogdanov, Tony Liang, Kriti Singh, Natan Dubitski, Ashley Waxman, Steve Landey, Misha Kestler, Susan Thomakos

埃尔达·博格丹诺夫、托尼·梁、克里蒂·辛格、纳坦·杜比茨基、阿什利·瓦克斯曼、史蒂夫·兰迪、米莎·凯斯特勒、苏珊·托马斯

Would you recommend this article?
Yes /
No 

你会推荐这篇文章吗？
是的 /
不

