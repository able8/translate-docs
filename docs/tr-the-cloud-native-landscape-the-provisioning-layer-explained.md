# The Cloud Native Landscape: The Provisioning Layer Explained

# 云原生景观：供应层解释

#### 20 Aug 2020 12:13pm, by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack. io/author/jason-morgan/ "Posts by Jason Morgan")

_This post is part of an ongoing series from [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/) and Jason Morgan that focuses on explaining each category of the cloud native landscape. Both are co-organizers of the [Kubernetes Community Days DC](https://kubernetescommunitydays.org/events/2021-washington-dc/) and the [DC Kubernetes meetup group](https://www.meetup.com/All-Things-Kubernetes-k8s-DC/)._

_这篇文章是 [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/) 和 Jason Morgan 正在进行的系列文章的一部分，该系列专注于解释云原生景观的每个类别。两者都是 [Kubernetes Community Days DC](https://kubernetescommunitydays.org/events/2021-washington-dc/) 和 [DC Kubernetes meetup group](https://www.meetup.com/万事万物-Kubernetes-k8s-DC/)._

Catherine is Head of Marketing at Buoyant, the creator of Linkerd. A marketing leader turned cloud native evangelist, Catherine is passionate about educating business leaders on the new stack and the critical flexibility it provides.

Catherine 是 Linkerd 的创建者 Buoyant 的营销主管。 Catherine 是一名营销领导者，后来成为云原生布道者，她热衷于就新堆栈及其提供的关键灵活性对业务领导者进行教育。

In our [introduction to the cloud native landscape](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/), we provided a high-level overview of the [Cloud Native Computing Foundation](https://www.cncf.io/)'s cloud native ecosystem. This article is the first in a series that examines each layer at the time. Non-technical readers will learn what the tools in each category are, what problem they solve, and how they address it. We also added a short technical 101 section for those engineers who are just getting started with cloud native.

在我们的[云原生景观简介](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/)中，我们提供了云原生计算基金会的高级概述的云原生生态系统。本文是当时检查每一层的系列文章中的第一篇。非技术读者将了解每个类别中的工具是什么、它们解决什么问题以及如何解决这些问题。我们还为刚开始使用云原生的工程师添加了一个简短的技术 101 部分。

The first layer in the cloud native landscape is provisioning. This encompasses the tools that are used to create and harden the foundation on which cloud native apps are built, including how the infrastructure is created, managed, and configured — automatically — as well as scanning, signing, and storing container images. The layer also extends to security with tools that enable policies to be set and enforced, authentication and authorization to be built into apps and platforms, and the handling of secrets distribution.

云原生环境中的第一层是配置。这包括用于创建和强化构建云原生应用程序的基础的工具，包括如何自动创建、管理和配置基础设施，以及扫描、签名和存储容器映像。该层还通过工具扩展到安全性，这些工具支持设置和执行策略、将身份验证和授权内置到应用程序和平台中，以及处理机密分发。

![provisioning layer](https://cdn.thenewstack.io/media/2020/08/5b39e1f1-screen-shot-2020-08-06-at-9.42.29-am.png)

When looking at the [cloud native landscape](https://landscape.cncf.io), you’ll note a few distinctions:

在查看 [云原生景观](https://landscape.cncf.io) 时，您会注意到一些区别：

- Projects in large boxes are CNCF-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- Projects in small white boxes are open source projects.
- Products in gray boxes are proprietary products.

- 大盒子中的项目是 CNCF 托管的开源项目。有的还在孵化阶段（浅蓝色/紫色框），有的则是毕业项目（深蓝色框）。
- 小白盒中的项目是开源项目。
- 灰色框内的产品是专有产品。

Please note that even during the time of this writing, we saw new projects becoming part of the CNCF so always refer to the actual landscape — things are moving fast!

请注意，即使在撰写本文时，我们也看到新项目成为 CNCF 的一部分，因此请始终参考实际情况——事情进展很快！

Ok, let’s have a look at each category of the provisioning layer, the role it plays, and how these technologies help adapt applications to a new cloud native reality.

好的，让我们来看看配置层的每个类别、它所扮演的角色，以及这些技术如何帮助应用程序适应新的云原生现实。

## Automation and Configuration

## 自动化和配置

### What It Is

###  这是什么

Automation and configuration tools speed up the creation and configuration of compute resources (virtual machines, networks, firewall rules, load balancers, etc.). These tools may handle different parts of the provisioning process or try to control everything end to end. Most provide the ability to integrate with other projects and products in the space.

自动化和配置工具可加快计算资源（虚拟机、网络、防火墙规则、负载平衡器等）的创建和配置。这些工具可能会处理供应过程的不同部分，或者尝试端到端地控制一切。大多数提供与空间中的其他项目和产品集成的能力。

### Problem it addresses

### 它解决的问题

Traditionally, IT processes relied on lengthy and labor-intensive manual release cycles, typically between three to six months. Those cycles came with lots of human processes and controls that slowed down changes to production environments. These slow-release cycles and static environments aren’t compatible with cloud native development. To deliver on rapid development cycles, infrastructure must be provisioned dynamically and without human intervention.

传统上，IT 流程依赖于冗长且劳动密集型的手动发布周期，通常在三到六个月之间。这些周期伴随着大量人工流程和控制，减缓了生产环境的变化。这些缓慢的发布周期和静态环境与云原生开发不兼容。为了实现快速的开发周期，必须动态配置基础设施，无需人工干预。

### How it helps

### 它如何帮助

Tools of this category allow engineers to build computing environments without human intervention. By codifying the environment setup it becomes reproducible with the click of a button. While manual setup is error-prone, once codified, environment creation matches the exact desired state — a huge advantage.

此类工具允许工程师在没有人工干预的情况下构建计算环境。通过对环境设置进行编码，单击按钮即可重现。虽然手动设置容易出错，但一旦编码，环境创建与确切的期望状态相匹配——这是一个巨大的优势。

While tools may take different approaches, they all aim at reducing the required work to provision resources through automation.

虽然工具可能采用不同的方法，但它们都旨在减少通过自动化配置资源所需的工作。

### Technical 101 

### 技术 101

As we move from old-style human-driven provisioning to the new on-demand scaling model demanded by the cloud we find that the patterns and tools we used before no longer meet our needs. Your organization can’t afford to maintain a large 24×7 staff whose job it is to create, configure, and manage servers. Automated tools like Terraform reduce the level of effort required to scale tens of servers and associated network and up to hundreds of firewall rules. Tools like Puppet, Chef, and Ansible provision and/or configure these new servers and applications programmatically as they are spun up and allow them to be consumed by developers.

当我们从旧式的人工驱动配置转向云所要求的新的按需扩展模型时，我们发现我们之前使用的模式和工具不再满足我们的需求。您的组织无法维持庞大的 24×7 员工，其工作是创建、配置和管理服务器。 Terraform 等自动化工具减少了扩展数十台服务器和相关网络以及多达数百个防火墙规则所需的工作量。 Puppet、Chef 和 Ansible 等工具在这些新服务器和应用程序启动时以编程方式提供和/或配置它们，并允许开发人员使用它们。

Some tools interact directly with the infrastructure APIs provided by platforms like AWS or vSphere, while others focus on configuring the individual machines to make them part of a Kubernetes cluster. Many, like Chef and Terraform, can interoperate to provision and configure the environment. Others, like OpenStack, exist to provide an Infrastructure-as-a-Service (IaaS) environment that other tools could consume. Fundamentally, you’ll need one or more tools in this space as part of laying down the computing environment, CPU, memory, storage, and networking, for your Kubernetes clusters. You’ll also need a subset of these to create and manage the Kubernetes clusters themselves.

一些工具直接与 AWS 或 vSphere 等平台提供的基础设施 API 交互，而其他工具则专注于配置单个机器以使其成为 Kubernetes 集群的一部分。许多，如 Chef 和 Terraform，可以互操作以供应和配置环境。其他工具（如 OpenStack）的存在是为了提供其他工具可以使用的基础设施即服务 (IaaS) 环境。从根本上说，在为 Kubernetes 集群布置计算环境、CPU、内存、存储和网络时，您需要在这个领域使用一个或多个工具。您还需要其中的一个子集来自己创建和管理 Kubernetes 集群。

At the time of this writing, there are three CNCF projects in this space: KubeEdge, a Sandbox CNCF project, as well as Kubespray and Kops (the latter two are Kubernetes subprojects and belong thus to the CNCF although they aren't yet listed on the landscape). Most of the tools in this category offer an open source as well as a paid version.

在撰写本文时，该领域有三个 CNCF 项目：KubeEdge，一个 Sandbox CNCF 项目，以及 Kubespray 和 Kops（后两个是 Kubernetes 子项目，因此属于 CNCF，尽管它们尚未在景观）。此类别中的大多数工具都提供开源和付费版本。

Buzz words Popular Projects/Products

流行语热门项目/产品

- Infrastructure-as-Code (IaC)
- Automation
- Declarative Configuration

- 基础设施即代码 (IaC)
- 自动化
- 声明式配置

- Chef
- Puppet
- Ansible
- Terraform

- 厨师
- 傀儡
- Ansible
- 地形

![automation and config](https://cdn.thenewstack.io/media/2020/08/df17f0e2-screen-shot-2020-08-06-at-9.52.34-am.png)

## Container Registry

## 容器注册表

### What It Is

###  这是什么

Before defining container registries, let’s first discuss three tightly related concepts:

在定义容器注册表之前，让我们首先讨论三个紧密相关的概念：

1. A container is a set of compute constraints used to execute a process. Processes launched within containers are tricked to believe they are running on their own dedicated computer vs. a machine shared with other processes (similar to virtual machines). In short, containers allow you to run your code in a controlled fashion no matter where it is.
2. An image is the set of archive files needed to run a container and its process. You could see it as a form of template on which you can create an unlimited number of containers.
3. A repository, or just repo, is a space to store images.

1. 容器是一组用于执行进程的计算约束。在容器中启动的进程被欺骗相信它们是在自己的专用计算机上运行，而不是与其他进程共享的计算机（类似于虚拟机）。简而言之，容器允许您以受控方式运行代码，无论它在哪里。
2. 镜像是运行容器及其进程所需的一组归档文件。您可以将其视为一种模板形式，您可以在其上创建无限数量的容器。
3. 存储库，或简称为 repo，是存储图像的空间。

Back to container registries. Container registries are specialized web applications to categorize and store repositories.

回到容器注册。容器注册表是专门用于对存储库进行分类和存储的 Web 应用程序。

In summary, images contain the information needed to execute a program (within a container) and are stored in repositories which in turn are categorized and grouped in registries. Tools that build, run, and manage containers need access to those images. Access is provided by referencing to the registry (the path to access the image).

总之，图像包含执行程序（在容器内）所需的信息，并存储在存储库中，而存储库又在注册表中进行分类和分组。构建、运行和管理容器的工具需要访问这些镜像。访问是通过引用注册表（访问图像的路径）来提供的。

![Registry, repo, containers](https://cdn.thenewstack.io/media/2020/08/b348c280-screen-shot-2020-08-06-at-9.59.11-am.png)

### Problem It Addresses

### 它解决的问题

Cloud native applications are packaged and run as containers. Container registries store and provide these container images.

云原生应用程序被打包并作为容器运行。容器注册表存储并提供这些容器镜像。

### How It Helps

### 它如何帮助

By centrally storing all container images in one place, they are easily accessible for any developer working on that app.

通过将所有容器映像集中存储在一个地方，任何开发该应用程序的开发人员都可以轻松访问它们。

### Technical 101 

### 技术 101

Container registry tools exist to either store and distribute images or to enhance an existing registry in some way. Fundamentally, a registry is a kind of web API that allows container engines to store and retrieve images. Many provide interfaces to allow container scanning or signing tools to enhance the security of the images they store. Some specialize in distributing or duplicating images in a particularly efficient manner. Any environment using containers will need to use one or more registries.

存在容器注册表工具来存储和分发映像或以某种方式增强现有注册表。从根本上说，注册中心是一种允许容器引擎存储和检索图像的 Web API。许多提供接口以允许容器扫描或签名工具来增强它们存储的图像的安全性。有些专门以特别有效的方式分发或复制图像。任何使用容器的环境都需要使用一个或多个注册表。

Tools in this space can provide integrations to scan, sign, and inspect the images they store. At the time of this writing Dragonfly and Harbor are CNCF projects and Harbor recently gained the distinction of [being the first](https://goharbor.io/blog/harbor-2.0/) OCI compliant registry. Each major cloud provider provides its own hosted registry and many other registries can be deployed standalone or directly into your Kubernetes cluster via tools like Helm.

此空间中的工具可以提供集成以扫描、签名和检查它们存储的图像。在撰写本文时，Dragonfly 和 Harbor 是 CNCF 项目，Harbor 最近获得了 [成为第一个](https://goharbor.io/blog/harbor-2.0/) OCI 兼容注册表的区别。每个主要的云提供商都提供自己的托管注册表，许多其他注册表可以独立部署或通过 Helm 等工具直接部署到您的 Kubernetes 集群中。

BuzzwordsPopular Projects/Products

流行语热门项目/产品

- Container
- OCI Image
- Registry

- 容器
- OCI 图像
- 注册表

- Docker Hub
- Harbor
- Hosted registries from AWS, Azure, and GCP
- Artifactory

- Docker 中心
- 港口
- 来自 AWS、Azure 和 GCP 的托管注册表
- 工艺品

![Container registries](https://cdn.thenewstack.io/media/2020/08/ee64a02c-screen-shot-2020-08-06-at-10.02.34-am.png)

## Security and Compliance

## 安全性和合规性

### What It Is

###  这是什么

Cloud native applications are designed to be rapidly iterated on. Think of the continuous flow of updates your iPhone apps get — everyday, they are evolving, presumably getting better. In order to release code on a regular cadence, we must ensure that our code and our operating environment are secure and only accessed by authorized engineers. Tools and projects in this section represent some of the things needed to create and run modern applications in a secure fashion.

云原生应用程序旨在快速迭代。想想你的 iPhone 应用程序获得的持续更新流——每天，它们都在发展，可能会变得更好。为了定期发布代码，我们必须确保我们的代码和我们的操作环境是安全的，并且只有经过授权的工程师才能访问。本节中的工具和项目代表了以安全方式创建和运行现代应用程序所需的一些东西。

### Problem It Addresses

### 它解决的问题

These tools and projects help you harden, monitor, and enforce security for your platforms and applications. From the container to your Kubernetes environment, they enable you to set policies (for compliance), get insights into existing vulnerabilities, catch misconfigurations, and harden the containers and clusters.

这些工具和项目可帮助您加强、监控和实施平台和应用程序的安全性。从容器到您的 Kubernetes 环境，它们使您能够设置策略（用于合规性）、深入了解现有漏洞、捕获错误配置以及强化容器和集群。

### How It Helps

### 它如何帮助

In order to securely run containers they must be scanned for known vulnerabilities and signed to ensure they haven’t been tampered with. Kubernetes itself defaults to extremely permissive access control settings that are unsuitable for production use. Furthermore, Kubernetes clusters are an attractive target to anyone looking to attack your systems. The tools and projects in this space help harden the cluster and provide tooling to detect when the system is behaving abnormally.

为了安全地运行容器，必须扫描它们是否存在已知漏洞并进行签名以确保它们未被篡改。 Kubernetes 本身默认为极其宽松的访问控制设置，不适合生产使用。此外，Kubernetes 集群对于任何想要攻击您的系统的人来说都是一个有吸引力的目标。该领域中的工具和项目有助于强化集群并提供工具来检测系统何时出现异常行为。

### Technical 101

### 技术 101

In order to operate securely in a dynamic and rapidly evolving environment we must treat security as part of the platform and application development lifecycle. The tools in this space are an extremely varied group and seek to solve different portions of the problem. Most of the tooling falls into one of the following categories:

为了在动态且快速发展的环境中安全运行，我们必须将安全视为平台和应用程序开发生命周期的一部分。这个领域的工具是一个极其多样化的群体，试图解决问题的不同部分。大多数工具属于以下类别之一：

- Audit and compliance
- Path to production hardening tools
   - Code scanning
   - Vulnerability scanning
   - Image signing
- Policy creation and enforcement
- Network layer security

- 审计与合规
- 生产硬化工具的途径
  - 扫码
  - 漏洞扫描
  - 图像签名
- 政策创建和执行
- 网络层安全

Some of these tools and projects will rarely be used directly, like Trivy, Claire, and Notary which are leveraged by registries or other scanning tools. Others are key hardening components of a modern application platform, like Falco or Open Policy Agent (OPA).

其中一些工具和项目很少被直接使用，例如注册机构或其他扫描工具利用的 Trivy、Claire 和 Notary。其他是现代应用程序平台的关键强化组件，如 Falco 或开放策略代理 (OPA)。

There are a number of mature vendors providing solutions in this space, as well as startups founded explicitly on bringing Kubernetes native frameworks to market. At the time of this writing Falco, Notary/TUF, and OPA are the only CNCF projects in this space.

有许多成熟的供应商在这个领域提供解决方案，以及明确将 Kubernetes 原生框架推向市场的初创公司。在撰写本文时，Falco、Notary/TUF 和 OPA 是该领域仅有的 CNCF 项目。

BuzzwordsPopular Projects/Products

流行语热门项目/产品

- Image scanning
- Image signing
- Policy enforcement
- Audit

- 图像扫描
- 图像签名
- 政策执行
- 审计

- OPA
- Falco
- Sonobuoy

- OPA
- 法尔科
- 声纳浮标

![security and compliance](https://cdn.thenewstack.io/media/2020/08/d60f1b0b-screen-shot-2020-08-06-at-10.06.50-am.png)

## Key (and Identity) Management

## 密钥（和身份）管理

### What It Is

###  这是什么

Before we go into key management, let’s first define cryptographic keys. A key is a string of characters used to encrypt or sign data. Like a physical key, it locks (encrypts) data so that only someone with the right key can unlock (decrypt) it. 

在我们进入密钥管理之前，让我们首先定义加密密钥。密钥是用于加密或签名数据的字符串。就像物理钥匙一样，它会锁定（加密）数据，以便只有拥有正确密钥的人才能解锁（解密）它。

As applications and operations adapt to a new cloud native world, security tools are evolving to meet new security needs. The tools and projects in this category cover everything from how to securely store passwords and other secrets (sensitive data such as API keys, encryption keys, etc.) to how to safely eliminate passwords and secrets from your microservices environment.

随着应用程序和操作适应新的云原生世界，安全工具也在不断发展以满足新的安全需求。此类别中的工具和项目涵盖了从如何安全地存储密码和其他机密（API 密钥、加密密钥等敏感数据）到如何从微服务环境中安全地消除密码和机密的所有内容。

### Problem It Addresses

### 它解决的问题

Cloud native environments are highly dynamic calling for secret distribution that is on-demand, entirely programmatic (no humans in the loop), and automated. Applications must also know if a given request comes from a valid source (authentication) and if that request has the right to do whatever it’s trying to do (authorization). This is commonly referred to as AuthN and AuthZ.

云原生环境是高度动态的，需要按需、完全程序化（没有人参与）和自动化的秘密分发。应用程序还必须知道给定的请求是否来自有效的来源（身份验证），以及该请求是否有权做它想做的任何事情（授权）。这通常称为 AuthN 和 AuthZ。

### How It Helps

### 它如何帮助

Each tool or project takes a different approach but they all provide a way to either securely distribute secrets and keys, or they provide a service or specification related to authentication, authorization, or both.

每个工具或项目都采用不同的方法，但它们都提供了一种安全分发机密和密钥的方法，或者它们提供与身份验证、授权或两者相关的服务或规范。

### Technical 101

### 技术 101

Tools in this category can be grouped into two sets: While some tools focus on key generation, storage, management and rotation, the other group focuses on single sign-on and identity management. Vault, for instance, is a rather generic key management tool allowing you to manage different types of keys. Keycloak, on the other hand, is an identity broker which can be used to manage access keys for different services.

此类别中的工具可分为两组：有些工具侧重于密钥生成、存储、管理和轮换，而另一组侧重于单点登录和身份管理。例如，Vault 是一个相当通用的密钥管理工具，允许您管理不同类型的密钥。另一方面，Keycloak 是一个身份代理，可用于管理不同服务的访问密钥。

At the time of this writing SPIFFE/SPIRE are the only CNCF projects in this space, and most tools offer an open source as well as paid version.

在撰写本文时，SPIFFE/SPIRE 是该领域仅有的 CNCF 项目，大多数工具都提供开源和付费版本。

BuzzwordsPopular Projects

流行语热门项目

- AuthN and AuthZ
- Identity
- Access
- Secrets

- AuthN 和 AuthZ
- 身份
- 使用权
- 秘密

- Vault
- Spiffe
- OAuth2

- 保险库
- 斯皮夫
- OAuth2

![Key management](https://cdn.thenewstack.io/media/2020/08/62e5a458-screen-shot-2020-08-06-at-10.13.56-am.png)

As we’ve seen the provisioning layer focuses on building the foundation of your cloud native platforms and applications, with tools handling everything from infrastructure provisioning to container registries to security. This piece is intended to be the first in a series of articles detailing the cloud native landscape. In the next article we’ll focus on the runtime layer and explore cloud native storage, container runtime, and networking.

正如我们所见，供应层专注于构建您的云原生平台和应用程序的基础，使用工具处理从基础设施供应到容器注册再到安全性的所有内容。本文旨在成为详细介绍云原生景观的系列文章中的第一篇。在下一篇文章中，我们将重点关注运行时层并探索云原生存储、容器运行时和网络。

_A very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate. Also, a big thanks to [Oleg Chunihkin](https://www.linkedin.com/in/olegch/) for all his input early on in this project._

_非常感谢来自 CNCF 的 [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/)，他非常友好地审阅了文章以确保其准确无误。另外，非常感谢 [Oleg Chunihkin](https://www.linkedin.com/in/olegch/) 在这个项目早期的所有投入。_

The Cloud Native Computing Foundation is a sponsor of The New Stack.

云原生计算基金会是 The New Stack 的赞助商。

Feature image by [torstensimon](https://pixabay.com/users/torstensimon-5039407/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=2211588) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=2211588).

[torstensimon](https://pixabay.com/users/torstensimon-5039407/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=2211588)的特色图片来自[Pixabay](https://pixabay.com/链接归因&utm_medium=referral&utm_campaign=image&utm_content=2211588)。

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Docker. 

The New Stack 是 Insight Partners 的全资子公司。 TNS 所有者 Insight Partners 是以下公司的投资者： Docker。

