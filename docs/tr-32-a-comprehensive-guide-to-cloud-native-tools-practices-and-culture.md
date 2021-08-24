# A Comprehensive Guide to Cloud Native: Tools, Practices and Culture

# 云原生综合指南：工具、实践和文化

In this guide to Cloud Native we outline the tools, cultural practices and concepts that come together to create a cloud native environment. We also outline the features of cloud native applications and review the tools that are the essential building blocks of cloud native environments.

在本云原生指南中，我们概述了共同创建云原生环境的工具、文化实践和概念。我们还概述了云原生应用程序的功能，并回顾了作为云原生环境基本构建块的工具。

February 3, 2020 From: https://www.replex.io/blog/a-comprehensive-guide-to-cloud-native-tools-practices-and-culture


Let's start by stating the obvious: Cloud native is "not" about the cloud. Counter intuitive as that may sound cloud native practices are not limited by the underlying infrastructure and can be adopted across any number of public, private, hybrid or on-premises infrastructures.

让我们从显而易见的事情开始：云原生与云“无关”。可能听起来违反直觉的云原生实践不受底层基础设施的限制，可以在任何数量的公共、私有、混合或本地基础设施中采用。

The same is true of cloud native tools, which can be deployed across "traditional" infrastructure technologies. These technologies however do need to emulate the cloud delivery model, which we will get into later. This essentially means that instead of being a cornerstone of cloud native, the cloud is only one of a number of tools that comprise the [cloud native landscape](https://landscape.cncf.io/).

云原生工具也是如此，它们可以跨“传统”基础设施技术进行部署。然而，这些技术确实需要模拟云交付模型，我们将在稍后介绍。这实质上意味着，云不是云原生的基石，而是构成 [云原生景观](https://landscape.cncf.io/) 的众多工具之一。

Let’s also answer the why of cloud native. Cloud native, as most other technologies in the IT landscape, is driven by the demands for speed and agility placed on the modern software delivery life cycle. Modern applications need to evolve quickly to meet changing customer demands or gain competitive advantages with innovative new features. Cloud native practices and tools allow organisations to do just that.

我们也来回答一下为什么是云原生。与 IT 领域中的大多数其他技术一样，云原生是由现代软件交付生命周期对速度和敏捷性的需求驱动的。现代应用程序需要快速发展以满足不断变化的客户需求或通过创新的新功能获得竞争优势。云原生实践和工具允许组织做到这一点。

## What is Cloud Native?

## 什么是云原生？

[Cloud native](https://github.com/cncf/toc/blob/master/DEFINITION.md) is set of tools and practices that allows organisations to build, deploy and operate software applications more frequently, predictably and reliably. Any organisation that builds and operates software applications using cloud native tools and adopts cloud native practices to do so can be said to be cloud native irrespective of the underlying infrastructure.

[云原生](https://github.com/cncf/toc/blob/master/DEFINITION.md) 是一组工具和实践，允许组织更频繁、可预测和可靠地构建、部署和运行软件应用程序。任何使用云原生工具构建和运行软件应用程序并采用云原生实践来这样做的组织都可以说是云原生的，而不管底层基础设施如何。

The distinction between tools and practices is an important one which is mostly ignored when describing a cloud native architecture. Cloud native tools are the specific technology pieces that go into the cloud native puzzle, practices refer to the underlying cultural dynamics of working with those tools. We will cover cloud native cultural practices in more detail later on in the article.

工具和实践之间的区别是一个重要的区别，在描述云原生架构时通常会忽略这一区别。云原生工具是进入云原生难题的特定技术部分，实践是指使用这些工具的潜在文化动态。我们将在本文后面更详细地介绍云原生文化实践。

Cloud native is a continuously evolving concept as new tools are developed and new practices take root in today's fast moving IT landscape. However we can identify a baseline of tools and practices that are common to most cloud native architectures. In the next few sections we will identify and describe these tools and practices.

随着新工具的开发和新实践在当今快速发展的 IT 环境中扎根，云原生是一个不断发展的概念。但是，我们可以确定大多数云原生架构通用的工具和实践的基线。在接下来的几节中，我们将确定并描述这些工具和实践。

The order in which these concepts are presented does not reflect their relative importance or the order in which they should be implemented. Cloud native, like consciousness in the human brain, is an [emergent quality](https://sciencing.com/emergent-properties-8232868.html) of organisations that implement a specific set of tools and adopt the cultural practices that go along with them.

这些概念的呈现顺序并不反映它们的相对重要性或它们应该实施的顺序。云原生，就像人脑中的意识一样，是组织的一种[紧急质量](https://sciencing.com/emergent-properties-8232868.html)，这些组织实施了一组特定的工具并采用了随之而来的文化实践与他们。

## Cloud Native Tools

## 云原生工具

### Cloud Delivery Model

### 云交付模型

The cloud or more precisely the [cloud delivery model](https://www.sciencedirect.com/topics/computer-science/cloud-delivery-model) of provisioning, consuming and managing resources is an essential part of cloud native architectures. Cloud native tools and cultural practices can only be adopted in an architecture that supports the on-demand dynamic provisioning of storage, compute and networking resources. This can be either a public cloud provider or an in-house private cloud solution that provides a cloud-like delivery model to IT teams.

云，或者更准确地说，提供、消费和管理资源的 [云交付模型](https://www.sciencedirect.com/topics/computer-science/cloud-delivery-model) 是云原生架构的重要组成部分。云原生工具和文化实践只能在支持按需动态配置存储、计算和网络资源的架构中采用。这可以是公共云提供商，也可以是为 IT 团队提供类似云的交付模型的内部私有云解决方案。

### Containers

### 容器

Containers are the lifeblood of cloud native applications. Individual containers package application code and all the resources required to run that code in a discrete unit of software. Containerisation makes applications more manageable allowing other technology pieces of the cloud native landscape to chip in and provide new and innovative solutions for application design, scalability, security and reliability etc.

容器是云原生应用程序的命脉。单个容器将应用程序代码和运行该代码所需的所有资源打包在一个独立的软件单元中。容器化使应用程序更易于管理，允许云原生环境的其他技术部分参与进来，并为应用程序设计、可扩展性、安全性和可靠性等提供新的创新解决方案。

Containerised applications are much more [portable](https://www.replex.io/blog/virtualization-vs-containers-scalability-portability-and-resource-utilization) as compared to their VM based counterparts and use the underlying resources more efficiently. They also have a much lower management and operational overhead. 
与基于 VM 的应用程序相比，容器化应用程序更[可移植](https://www.replex.io/blog/virtualization-vs-containers-scalability-portability-and-resource-utilization)，并且更多地使用底层资源有效地。它们的管理和运营开销也低得多。
Since [containers](http://www.replex.io/blog/virtualization-vs-containers-scalability-portability-and-resource-utilization) are platform agnostic, they result in lower integration issues and reduced testing and debugging as part of the developer workflow. The ease with which containers can be created, destroyed and updated leads to an overall acceleration in speed to market for new application features, allowing organisations to keep up with changing customer demands.

由于 [containers](http://www.replex.io/blog/virtualization-vs-containers-scalability-portability-and-resource-utilization) 与平台无关，它们导致较低的集成问题并减少测试和调试作为一部分开发人员工作流程。容器可以轻松创建、销毁和更新，从而整体加快了新应用程序功能的上市速度，使组织能够跟上不断变化的客户需求。

### Microservices

### 微服务

Microservices architecture is a way of architecting applications as a collection of services. It breaks down applications into easy manageable 'microservices', each performing a specific business function, owned by an individual team and communicating with other related [microservices](https://en.wikipedia.org/wiki/Microservices).

微服务架构是一种将应用程序架构为服务集合的方式。它将应用程序分解为易于管理的“微服务”，每个都执行特定的业务功能，由单个团队拥有并与其他相关 [微服务](https://en.wikipedia.org/wiki/Microservices) 进行通信。

Microservices architecture fits in nicely with the cloud native cultural practice of self contained, agile, autonomous teams with native knowledge and skills to develop, test, deploy and operate each microservice.

微服务架构非常适合具有本地知识和技能的自包含、敏捷、自治团队的云原生文化实践，以开发、测试、部署和操作每个微服务。

Breaking up applications in this way allows application components to be developed, deployed, managed and operated independently. This has major implications for developer productivity and deployment speed of applications. The loosely coupled nature of microservices applications also means that production issues in any one microservice do not lead to application wide outages. This makes it easier to contain production issues as well as respond and recover quickly.

以这种方式拆分应用程序允许独立开发、部署、管理和操作应用程序组件。这对开发人员的生产力和应用程序的部署速度有重大影响。微服务应用程序的松散耦合特性也意味着任何一个微服务中的生产问题都不会导致应用程序范围内的中断。这使得控制生产问题以及快速响应和恢复变得更加容易。

### Service Mesh

### 服务网格

Breaking applications down into a collection of loosely coupled microservices leads to an increase in the volume of service to service communication. Most cloud native applications comprise of hundreds of microservices, communicating in complex webs.

将应用程序分解为一组松散耦合的微服务会导致服务间通信量的增加。大多数云原生应用程序由数百个微服务组成，在复杂的网络中进行通信。

Service meshes manage this complex web of service to service communication at scale and make it secure, fast and reliable. [Istio](https://github.com/istio/istio), Lankerd and Consul are prominent examples from the cloud native landscape. Service meshes work by decoupling communication protocols from application code and abstracting it to an infrastructure layer atop TCP/IP. This reduces the overhead for developers who can concentrate on building new features rather then managing managing communcations.

服务网格管理这个复杂的服务网络，大规模地进行服务通信，并使其安全、快速和可靠。 [Istio](https://github.com/istio/istio)、Lankerd 和 Consul 是云原生领域的突出例子。服务网格通过将通信协议与应用程序代码解耦并将其抽象到 TCP/IP 之上的基础设施层来工作。这减少了开发人员的开销，他们可以专注于构建新功能而不是管理通信。

### Continuous Integration and Delivery

### 持续集成和交付

Continuous integration and delivery can refer to both a set of practices as well as the tools that support those practices, aimed at accelerating software development cycles and making them more robust and reliable. [CICD tools](https://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-continuous-integration-and-delivery-cicd) automate crucial stages of the software release cycle. They also encourage a culture of shared responsibility, small incremental changes to application code, continuously integrating and testing those changes and using version control. CICD practices also extend to the delivery and deployment stages by ensuring new features are [production ready](https://www.replex.io/kubernetes-production-readiness-and-best-practices-checklist) once they go through the automated integration and testing phase.

持续集成和交付既可以指一组实践，也可以指支持这些实践的工具，旨在加快软件开发周期并使它们更加健壮和可靠。 [CICD 工具](https://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-continuous-integration-and-delivery-cicd) 自动化软件发布周期的关键阶段。他们还鼓励分担责任的文化，对应用程序代码进行小的增量更改，不断集成和测试这些更改并使用版本控制。 CICD 实践还扩展到交付和部署阶段，确保新功能在通过自动化后 [生产就绪](https://www.replex.io/kubernetes-production-readiness-and-best-practices-checklist)集成和测试阶段。

### Orchestration

### 编排

Modern enterprise applications span multiple containerised microservices deployed across a number of public and private cloud hosts. Deploying, operating and scaling the fleets of containerised microservices making up these applications, while ensuring high availability is no easy task. This is where container orchestrators like Kubernetes shine.

现代企业应用程序跨越多个容器化微服务，部署在许多公共和私有云主机上。部署、操作和扩展构成这些应用程序的容器化微服务队列，同时确保高可用性并非易事。这就是像 Kubernetes 这样的容器编排器大放异彩的地方。

[Kubernetes](http://kubernetes.io/) makes it easier to provision, deploy and scale fleets of containerised microservices applications. It handles most of the mechanics of placing containers on hosts, load balancing across hosts as well as removing and re-spawning containers under the hood. Network and storage abstractions coupled with standardised resource and configuration definitions add an additional layer of portability on top of containerisation.

[Kubernetes](http://kubernetes.io/) 可以更轻松地提供、部署和扩展容器化微服务应用程序队列。它处理在主机上放置容器、跨主机负载平衡以及在后台删除和重新生成容器的大部分机制。网络和存储抽象加上标准化的资源和配置定义，在容器化之上增加了一层额外的可移植性。

All of this makes Kubernetes an indispensable cog in cloud native environments. Kubernetes by itself does not qualify as cloud native, however no environment can truly be cloud native without some sort of orchestration engine at its heart.

所有这些都使 Kubernetes 成为云原生环境中不可或缺的齿轮。 Kubernetes 本身并不具备云原生的资格，但是如果没有某种编排引擎的核心，任何环境都不能真正成为云原生。

## Cloud Native Culture

## 云原生文化

Culture is a nebulous concept. While most agree that it [exists](https://hbr.org/2013/05/what-is-organizational-culture) and influences behaviour and practices within organisations it is not easy to pin down.

文化是一个模糊的概念。虽然大多数人都同意它 [存在](https://hbr.org/2013/05/what-is-organizational-culture) 并影响组织内的行为和实践，但并不容易确定。

### Cloud native Culture vs Cloud native tools 
### 云原生文化 vs 云原生工具
The same is true of the cultural component of cloud native systems and architectures. Cloud native culture is mostly confused with cloud native tools. While cloud native tools are relatively easy to integrate into pre existing workflows for building and releasing software products, culture is hard to adopt or even define.

云原生系统和架构的文化组件也是如此。云原生文化大多与云原生工具混淆。虽然云原生工具相对容易集成到用于构建和发布软件产品的现有工作流程中，但文化很难采用甚至定义。

For the purposes of this article, we define culture, as a collection of practices centered around the way organisations build, release and operate software products. Culture is the way these organisations build products, not the tools they use to build them.

出于本文的目的，我们将文化定义为围绕组织构建、发布和运行软件产品的方式的一系列实践。文化是这些组织构建产品的方式，而不是他们用来构建产品的工具。

### Cloud native Practices

### 云原生实践

A [recent survey](https://www.replex.io/replex-state-of-kubernetes-report-2019-download) by the Replex team of IT practitioners at KubeCon Barcelona identifies cultural change as the biggest obstacle to cloud native adoption. It comes out ahead of complexity, planning and deployment and the lack of internal interest in terms of relative difficulty. In this section we will identify some of these practices that epitomise cloud native culture. Let’s start with the most obvious candidate, DevOps.

[最近的调查](https://www.replex.io/replex-state-of-kubernetes-report-2019-download) 在 KubeCon Barcelona 上由 Replex 的 IT 从业人员团队进行的一项调查表明，文化变革是云原生的最大障碍通过。它出现在复杂性、规划和部署以及相对困难方面缺乏内部兴趣之前。在本节中，我们将确定一些体现云原生文化的实践。让我们从最明显的候选者 DevOps 开始。

#### DevOps and SRE

#### DevOps 和 SRE

Even though DevOps predates cloud native it is widely considered an essential component of cloud native systems or at the very least an essential on ramp to cloud native culture. DevOps breaks down the silos that traditional dev and ops teams operated in by encouraging and facilitating open two way communication and collaboration. This leads to better team integration and promises accelerated software delivery and innovation.

尽管 DevOps 早于云原生，但它被广泛认为是云原生系统的重要组成部分，或者至少是向云原生文化过渡的重要组成部分。 DevOps 通过鼓励和促进开放的双向沟通​​和协作，打破了传统开发和运营团队运作的孤岛。这会导致更好的团队整合，并承诺加速软件交付和创新。

[SRE](https://landing.google.com/sre/sre-book/toc/index.html), an iteration of DevOps developed internally by google, takes this one step further by taking a software centric view of operations. It encourages traditional developers to internalize operations skills including networking, system administration and automation.

[SRE](https://landing.google.com/sre/sre-book/toc/index.html) 是 google 内部开发的 DevOps 迭代，通过采用以软件为中心的操作视图，更进一步。它鼓励传统开发人员将操作技能内化，包括网络、系统管理和自动化。

#### Microservices and CICD

#### 微服务和 CICD

The way in which cloud native applications are architected requires the formation of close knit cross-functional teams responsible for individual components (microservices) of applications. These teams have end to end responsibility for developing, testing, deploying and operating these components and therefore need to internalize a broad set of skills.

云原生应用程序的架构方式需要组成紧密的跨职能团队，负责应用程序的各个组件（微服务）。这些团队对开发、测试、部署和操作这些组件负有端到端的责任，因此需要内化广泛的技能。

Better alignment with rapidly changing customer demands and the drive to gain competitive advantages with new features requires organisations to adopt a culture of frequent, small, rapid releases. CICD practices such as building in automation, shared responsibility and being [production ready](https://www.replex.io/kubernetes-production-readiness-and-best-practices-checklist) at all times are also crucial components of cloud native culture.

更好地与快速变化的客户需求保持一致以及通过新功能获得竞争优势的动力要求组织采用频繁、小规模、快速发布的文化。 CICD 实践，例如始终构建自动化、分担责任和 [生产就绪](https://www.replex.io/kubernetes-production-readiness-and-best-practices-checklist) 也是云的重要组成部分本土文化。

## Cloud Native Applications

## 云原生应用

Now that we have wrapped our heads around the concept of a cloud native architecture let’s take another stab at defining a cloud native application.

现在我们已经了解了云原生架构的概念，让我们再次尝试定义云原生应用程序。

A cloud native application is a set of multiple loosely coupled containerised microservices, deployed using an orchestration engine like Kubernetes with the cloud or a cloud like delivery model as an underlying layer.

云原生应用程序是一组多个松散耦合的容器化微服务，使用编排引擎（如 Kubernetes）以云或类似云的交付模型作为底层进行部署。

Cloud native applications are not static entities however and evolve continuously in response to the external environment. This is where the wider constellation of supporting tools comes into play, which need to be adopted to enable teams to develop, test and deploy code more frequently and reliably.

然而，云原生应用程序不是静态实体，而是根据外部环境不断发展。这就是更广泛的支持工具群发挥作用的地方，需要采用这些工具使团队能够更频繁、更可靠地开发、测试和部署代码。

All of this is supported by an underlying cultural layer which advocates the removal of silos in the software life cycle. One way to accomplish this is to create small, agile independent teams with end to end responsibility for developing, testing, deploying and operating applications. These cross functional teams comprising developers, DevOps and SREs would also be responsible for integrating tools, building in automation, monitoring, self-healing, managing performance and ensuring high availability.

所有这一切都得到了一个底层文化层的支持，它提倡在软件生命周期中消除孤岛。实现这一目标的一种方法是创建小型、敏捷的独立团队，负责开发、测试、部署和操作应用程序的端到端责任。这些由开发人员、DevOps 和 SRE 组成的跨职能团队还将负责集成工具、构建自动化、监控、自我修复、管理性能和确保高可用性。

### Cloud Native Application Features

### 云原生应用功能

So what distinguishes cloud native applications from traditional ones? In this section we will briefly review some features of cloud native applications that give them an edge over traditional applications.

那么云原生应用与传统应用的区别是什么？在本节中，我们将简要回顾云原生应用程序的一些特性，这些特性使它们比传统应用程序更具优势。

#### Scalable 
#### 可扩展
Cloudnative applications are inherently scalable. Some of this can be attributed to the in-built scalability of the underlying technologies they are supported by. Take Kubernetes for example, it can [scale](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) both applications and the underlying infrastructure based on a number of business, application and server metrics. The same holds true for the cloud, which has built-in scalability mechanisms for most services.

云原生应用程序具有固有的可扩展性。其中一些可归因于它们所支持的底层技术的内置可扩展性。以Kubernetes为例，它可以[扩展](https://kubernetes.io/docs/tasks/run-application/horizo​​ntal-pod-autoscale/)基于多个业务、应用和服务器的应用和底层基础设施指标。云也是如此，它具有适用于大多数服务的内置可扩展性机制。

#### Observable

#### 可观察的

As opposed to monitoring, [observability](https://www.replex.io/blog/a-cios-guide-to-kubernetes-best-practices-in-production) is a property of a system. Systems are observable if their current state and in turn performance can be inferred based only on outputs. The support for traces, metrics, events and logs, in most of the underlying technology means that cloud native applications are highly observable.

与监控相反，[可观察性](https://www.replex.io/blog/a-cios-guide-to-kubernetes-best-practices-in-production) 是系统的属性。如果系统的当前状态和性能只能根据输出推断出来，则系统是可观察的。大多数底层技术中对跟踪、指标、事件和日志的支持意味着云原生应用程序是高度可观察的。

#### Resilient and Manageable

#### 弹性和可管理性

Reliability and resilience is another feature of cloud native applications. The fact that they are composed of multiple loosely coupled services means that application wide shutdowns are rare. Since problems are contained, disaster recovery is also relatively easier. The disaster recovery mechanisms of the underlying technology and the cloud native cultural practices of version control and using git as a single source of truth add to the ease of disaster recovery.

可靠性和弹性是云原生应用程序的另一个特性。它们由多个松散耦合的服务组成的事实意味着应用程序范围内的关闭很少见。由于问题被遏制了，容灾也相对容易一些。底层技术的容灾机制以及版本控制和使用git作为单一事实来源的云原生文化实践增加了容灾的便利性。

Manageability is another aspect which is easier because of the loosely coupled nature of cloud native applications. New features can be easily integrated and are immediately deployable using a number of deployment techniques. Individual application components can be updated without turning the lights off all over.

由于云原生应用程序的松散耦合特性，可管理性是另一个更容易的方面。新功能可以轻松集成，并且可以使用多种部署技术立即部署。可以在不完全关闭灯的情况下更新单个应用程序组件。

#### Immutability and API driven

#### 不变性和 API 驱动

[Immutability](https://github.com/cncf/toc/blob/master/DEFINITION.md) is a cultural aspect of cloud native architectures and plays an important role in how these applications are managed. Immutability is the practice of replacing containers instead of updating them in place. This practice is increasingly taking hold across cloud native environments and is cited by the CNCF as a defining feature of cloud native architectures and applications.

[不变性](https://github.com/cncf/toc/blob/master/DEFINITION.md) 是云原生架构的一个文化方面，在如何管理这些应用程序方面发挥着重要作用。不变性是替换容器而不是更新它们的做法。这种做法越来越多地在云原生环境中占据一席之地，并被 CNCF 引用为云原生架构和应用程序的定义特征。

Declarative APIs is another feature that the CNCF cites. [Declarative APIs](https://github.com/cncf/toc/blob/master/DEFINITION.md) concentrate on outcomes rather than explicitly mapping out a set of actions. Cloud native applications are usually integrated with lightweight APIs such as representational state transfer (REST), Google’s open source remote procedure call (gRPC) or NATS.

声明式 API 是 CNCF 引用的另一个特性。 [声明式 API](https://github.com/cncf/toc/blob/master/DEFINITION.md) 专注于结果，而不是明确映射出一组操作。云原生应用程序通常与轻量级 API 集成，例如表征状态传输 (REST)、谷歌的开源远程过程调用 (gRPC) 或 NATS。

## CNCF Cloud native Landscape

## CNCF 云原生景观

The [CNCF cloud native landscape](https://landscape.cncf.io/) is a collection of open source and third party tools under the umbrella of the cloud native computing foundation. These tools cover most aspects of cloud native environments and are aimed at helping companies build end-to-end technology stacks for developing, deploying, operating and monitoring cloud native applications.

[CNCF 云原生景观](https://landscape.cncf.io/) 是云原生计算基金会旗下的开源和第三方工具集合。这些工具涵盖了云原生环境的大部分方面，旨在帮助公司构建端到端的技术堆栈，用于开发、部署、操作和监控云原生应用程序。

In the next section we will review some of these cloud native tools targeted towards specific aspects of cloud native environments. Let’s start with developer tools.

在下一节中，我们将回顾其中一些针对云原生环境特定方面的云原生工具。让我们从开发者工具开始。

### Cloud native Developer tools

### 云原生开发者工具

Kubernetes, while making it easier to deploy and operate containerized applications, has also introduced a number of abstractions into the application development workflow. These new abstractions include but are not limited to pods, nodes, namespaces and deployments. Developers need to familiarize themselves with these new abstractions and incorporate them into already existing development workflows.

Kubernetes 在使部署和操作容器化应用程序变得更容易的同时，还在应用程序开发工作流中引入了许多抽象。这些新的抽象包括但不限于 pod、节点、命名空间和部署。开发人员需要熟悉这些新的抽象并将它们合并到现有的开发工作流程中。

The following tools allow developers to do just that by decluttering the development pipeline for Kubernetes based cloud native applications and reducing management overhead for developers.

以下工具允许开发人员通过整理基于 Kubernetes 的云原生应用程序的开发管道并减少开发人员的管理开销来做到这一点。

#### Draft

#### 草稿

[Draft](https://draft.sh/) aids developers by providing two main features: Draft create which automatically spins up the artifacts needed to run Kubernetes applications and Draft up which builds container images from code and deploys it to a Kubernetes cluster

[Draft](https://draft.sh/) 通过提供两个主要功能来帮助开发人员：Draft create 自动启动运行 Kubernetes 应用程序所需的工件和 Draft up，它从代码构建容器映像并将其部署到 Kubernetes 集群

#### Skaffold

####脚手架

[Skaffold](https://github.com/GoogleContainerTools/skaffold) allows developers to iterate on application code locally, build container images and deploy it to local or remote clusters as well as providing a minimal CICD pipeline.

[Skaffold](https://github.com/GoogleContainerTools/skaffold) 允许开发人员在本地迭代应用程序代码，构建容器镜像并将其部署到本地或远程集群，并提供最小的 CICD 管道。

#### Telepresence

####远程呈现

[Telepresence](http://telepresence.io/) accelerates application development by allowing developers to develop services locally, connect those services to remote clusters and automatically triggers updates whenever changes occur locally.

[Telepresence](http://telepresence.io/) 通过允许开发人员在本地开发服务、将这些服务连接到远程集群并在本地发生更改时自动触发更新来加速应用程序开发。

#### Okteto 
#### 奥克特托
[Okteto](https://okteto.com/) allows developers to spin up development environments in remote Kubernetes clusters, detects local code changes and synchronizes them to remote dev environments. By doing this it enables developers to work with their favorite tools locally, accelerates application development and reduces integration issues.

[Okteto](https://okteto.com/) 允许开发人员在远程 Kubernetes 集群中启动开发环境，检测本地代码更改并将其同步到远程开发环境。通过这样做，它使开发人员能够在本地使用他们喜欢的工具，加速应用程序开发并减少集成问题。

### Cloud native CICD tools

### 云原生 CICD 工具

CICD tools aim to accelerate application development and delivery as well as reduce integration and production issues. Most CICD tools predate cloud native architectures, and as such are not aligned with the specific requirements of these architectures. Some CICD providers have however developed variants targeted towards such architectures. Besides these completely new CICD tools built from the ground up for cloud native architectures have also cropped up. In the next section we will review some of these tools.

CICD 工具旨在加速应用程序开发和交付，并减少集成和生产问题。大多数 CICD 工具早于云原生架构，因此不符合这些架构的特定要求。然而，一些 CICD 提供商开发了针对此类架构的变体。除了这些从头开始为云原生架构构建的全新 CICD 工具之外，也出现了。在下一节中，我们将回顾其中一些工具。

#### Jenkins X

####詹金斯X

[Jenkins X](https://jenkins-x.io/) allows developers to build CICD pipelines without having to know the internals of Kubernetes or keeping up with the ever-growing list of its functionalities. It has a number of nifty build-in features including baseline CICD pipelines that incorporate DevOps and GitOps best practices, team and preview environments and feedback on issues and pull requests.

[Jenkins X](https://jenkins-x.io/) 允许开发人员构建 CICD 管道，而无需了解 Kubernetes 的内部结构或跟上其不断增长的功能列表。它具有许多漂亮的内置功能，包括包含 DevOps 和 GitOps 最佳实践的基线 CICD 管道、团队和预览环境以及对问题和拉取请求的反馈。

#### Gitlab

### GitLab

Gitlab is another feature rich CICD platform which integrates into the larger [Gitlab](https://about.gitlab.com/) suite of tools. One interesting feature is Auto DevOps which spins up predefined CICD pipelines whenever new projects are created. Deploy boards are another and help DevOps monitor the health and status of CI environments running on Kubernetes.

Gitlab 是另一个功能丰富的 CICD 平台，它集成到更大的 [Gitlab](https://about.gitlab.com/) 工具套件中。一个有趣的功能是 Auto DevOps，它在创建新项目时启动预定义的 CICD 管道。部署板是另一个，帮助 DevOps 监控在 Kubernetes 上运行的 CI 环境的健康和状态。

#### Argo

Argo is built from the ground up for Kubernetes based cloud native applications and leverages Kubernetes CRDs to implement CICD pipelines. This allows pipelines to be managed using native Kubernetes tools like kubectl and also means that they have a much broader integration with other Kubernetes services. [Argo](https://argoproj.github.io/argo-cd/) monitors live applications and compares it to the one kept under version control. In the event of any diversion it automatically triggers a synching mechanism.

Argo 是为基于 Kubernetes 的云原生应用程序从头开始构建的，并利用 Kubernetes CRD 来实现 CICD 管道。这允许使用 kubectl 等原生 Kubernetes 工具管理管道，也意味着它们与其他 Kubernetes 服务有更广泛的集成。 [Argo](https://argoproj.github.io/argo-cd/) 监控实时应用程序并将其与版本控制下的应用程序进行比较。如果发生任何转移，它会自动触发同步机制。

#### GoCD

#### GoCD

[GoCD](https://www.gocd.org/) can be easily installed in Kubernetes clusters using a Helm chart which spins up a GoCD server and elastic agents as pods. CD pipelines can then be defined as either json or yaml files. GoCD also allows DevOps to import sample pipelines to get up and running quickly and configure them with native Kubernetes artefacts like secrets, service accounts and API tokens.

[GoCD](https://www.gocd.org/) 可以使用 Helm 图表轻松安装在 Kubernetes 集群中，该图表将 GoCD 服务器和弹性代理作为 pod 启动。然后可以将 CD 管道定义为 json 或 yaml 文件。 GoCD 还允许 DevOps 导入示例管道以快速启动和运行，并使用本地 Kubernetes 人工制品（如机密、服务帐户和 API 令牌）对其进行配置。

### Cloud Native Network Tools

### 云原生网络工具

Cloud native applications are usually split into functional pieces called micro services that communicate with each other to perform higher order application functions. This leads to an increase in the volume of communications between application components and a correspondingly larger networking footprint. Cloud VMs, containers and Kubernetes, which are essential underlying technologies supporting cloud native applications, have their own networking requirements.

云原生应用程序通常被拆分为称为微服务的功能块，它们相互通信以执行更高阶的应用程序功能。这导致应用程序组件之间的通信量增加，并且相应地增加了网络占用空间。云虚拟机、容器和 Kubernetes 是支持云原生应用必不可少的底层技术，它们都有自己的网络需求。

Cloud native networks tools make it easier to manage networking for cloud native applications. In the next section we will review some of these tools featured in the CNCF cloud native landscape.

云原生网络工具可以更轻松地管理云原生应用程序的网络。在下一节中，我们将回顾 CNCF 云原生景观中的一些工具。

#### Weavenet

####织网

[Weavenet](https://www.weave.works/docs/net/latest/overview/) is a virtual layer 2 network that connects containers on the same host or across multiple hosts. As opposed to most other cloud native network tools, it does not require an external data store, significantly reducing the overhead for developers and DevOps managing these networks. Weavenet uses VXLan encapsulation, encrypts traffic using either ESP of IPsec or NACL, supports partially connected networks and automatically forwards traffic via the fastest path between two hosts. Besides this it also supports Kubernetes network policies, service discovery and load balancing.

[Weavenet](https://www.weave.works/docs/net/latest/overview/) 是一个虚拟的第 2 层网络，用于连接同一主机或跨多个主机上的容器。与大多数其他云原生网络工具不同，它不需要外部数据存储，从而显着降低了开发人员和 DevOps 管理这些网络的开销。 Weavenet 使用 VXLan 封装，使用 IPsec 的 ESP 或 NACL 对流量进行加密，支持部分连接的网络并通过两个主机之间的最快路径自动转发流量。除此之外，它还支持 Kubernetes 网络策略、服务发现和负载均衡。

#### Calico 
#### 印花布
[Calico](https://www.projectcalico.org/) functions without encapsulation as a layer 3 network as well as with encapsulation using either VXlan or IP-In-IP. It can also dynamically switch between the two based on whether traffic traverses a subnet boundary or stays within it. Calico can use both etcd as well as the Kubernetes API datastore and supports TLS encryption for communication between etcd and Calico components. It natively integrates into managed Kubernetes services from most public cloud providers, and supports both unicast and anycast IP. It also provides a native network policy resource that expands on the Kubernetes network policy feature set.

[Calico](https://www.projectcalico.org/) 无需封装即可作为第 3 层网络运行，也可以使用 VXlan 或 IP-In-IP 进行封装。它还可以根据流量是穿越子网边界还是停留在子网边界内，在两者之间动态切换。 Calico 可以同时使用 etcd 和 Kubernetes API 数据存储，并支持 TLS 加密用于 etcd 和 Calico 组件之间的通信。它本机集成到来自大多数公共云提供商的托管 Kubernetes 服务中，并支持单播和任播 IP。它还提供了扩展 Kubernetes 网络策略功能集的本机网络策略资源。

#### Flannel

####法兰绒

[Flannel](https://coreos.com/flannel/docs/latest/) deploys a virtual layer 2 overlay network that spans across the entire cluster and can use both VXlan and host-gw as a backend. UDP can also be used, however it is only recommended for debugging. Similar to most other cloud native networks it uses etcd as a datastore, and can be run using IPsec or Wireguard backends to encrypt traffic. Unlike most other cloud native networks however, Flannel does not support the Kubetnretes network policy resource.

[Flannel](https://coreos.com/flannel/docs/latest/) 部署了一个跨越整个集群的虚拟二层覆盖网络，可以同时使用 VXlan 和 host-gw 作为后端。也可以使用 UDP，但仅推荐用于调试。与大多数其他云原生网络类似，它使用 etcd 作为数据存储，并且可以使用 IPsec 或 Wireguard 后端运行以加密流量。然而，与大多数其他云原生网络不同，Flannel 不支持 Kubetnretes 网络策略资源。

#### Cilium

#### 纤毛

[Cilium](https://cilium.io/) is an open source cloud native network built on top of BPF. BPF enables it to perform filtering at the kernel level as well as support highly scalable [load balancing](https://docs.cilium.io/en/stable/intro/#load-balancing) for traffic between containers and to external services . Cilium supports both VXlan and Geneve encapsulation, can be configured to use either etcd or Consul as data stores and also supports IPsec encryption. It extends the Kubetnetes network policy resource to add support for layer 7 policy enforcement on ingress and egress for Http and kafka as well as egress support for CIDRs.

[Cilium](https://cilium.io/) 是一个建立在 BPF 之上的开源云原生网络。 BPF 使其能够在内核级别执行过滤，并支持高度可扩展的[负载平衡]（https://docs.cilium.io/en/stable/intro/#load-balancing）用于容器之间和外部服务的流量. Cilium 支持 VXlan 和 Geneve 两种封装，可以配置为使用 etcd 或 Consul 作为数据存储，还支持 IPsec 加密。它扩展了 Kubetnetes 网络策略资源，以添加对 Http 和 kafka 的入口和出口的第 7 层策略实施的支持以及对 CIDR 的出口支持。

#### Contiv

####康提

[Contiv](https://contiv.io/) is an open source cloud native network from Cisco that supports multiple operational modes including L2, L3, overlay and ACI. It uses etcd as a datastore and has its own built-in network policy resource that replaces the vanilla Kubernetes network policy resource. The built-in policy resource supports both bandwidth policies, that allow users to control the overall resource use of a group of containers and isolation that allows them to control the access of a group of containers. Contiv supports overlapping IPs across hosts and enables multi-tenant support. It uses the DNS protocol for service discovery, does not require queries to external data stores for IP or port information, has built-in support for service load balancing and allows admins to manage users, authorization and LDAP.

[Contiv](https://contiv.io/) 是 Cisco 的开源云原生网络，支持 L2、L3、overlay 和 ACI 等多种操作模式。它使用 etcd 作为数据存储，并拥有自己的内置网络策略资源来替代 vanilla Kubernetes 网络策略资源。内置的策略资源支持带宽策略，允许用户控制一组容器的整体资源使用和隔离，允许他们控制一组容器的访问。 Contiv 支持跨主机的重叠 IP 并启用多租户支持。它使用 DNS 协议进行服务发现，不需要向外部数据存储查询 IP 或端口信息，内置支持服务负载平衡，并允许管理员管理用户、授权和 LDAP。

### Cloud Native Service Mesh tools

### 云原生服务网格工具

In the previous section we outlined multiple [cloud native network tools](http://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-cloud-native-network-tools). At first sight these tools seem to have a lot in common with service meshes. There are some important differences however. Managing inter-service communications at the scale required by today’s microservices based enterprise applications quickly becomes infeasible with existing networking tools. Securing, monitoring and orchestrating these communications as well as implementing observability paradigms like tracing and logging add additional complexity.

在上一节中，我们概述了多个云原生网络工具。乍一看，这些工具似乎与服务网格有很多共同之处。然而，有一些重要的区别。以当今基于微服务的企业应用程序所需的规模管理服务间通信很快变得无法使用现有的网络工具。保护、监控和协调这些通信以及实施跟踪和日志记录等可观察性范式增加了额外的复杂性。

Service meshes operate alongside cloud native network tools and extend their feature-set by adding security, orchestration, tracing and logging for inter-service communication. Service mesh tools create an abstraction layer on top of microservices allowing DevOps to manage, orchestrate, monitor, observe and secure the communications between those services.

服务网格与云原生网络工具一起运行，并通过为服务间通信添加安全性、编排、跟踪和日志记录来扩展其功能集。服务网格工具在微服务之上创建了一个抽象层，允许 DevOps 管理、编排、监控、观察和保护这些服务之间的通信。

Service meshes are composed of two main components: the control plane and the data plane. In Kubernetes environments the data plane is usually a proxy like envoy deployed alongside a microservice as a side-car container. Proxies handle all traffic to and from the microservice based on policies configured in the control plane.

服务网格由两个主要组件组成：控制平面和数据平面。在 Kubernetes 环境中，数据平面通常是一个类似于 Envoy 的代理，作为 side-car 容器与微服务一起部署。代理根据控制平面中配置的策略处理进出微服务的所有流量。

#### **Istio** 
[Istio](http://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-cloud-native-network-tools) is one of the most popular service mesh tools in the CNCF cloud native landscape. It is very well integrated into Kubernetes, both in standalone Kubernetes environments as well as managed Kubernetes offerings from major cloud providers. Istio uses an extended version of the envoy proxy and deploys it alongside each microservice pod in Kubernetes environments. It has a broad feature set allowing DevOps to configure and create policies for circuit breakers, timeouts, retries, AB/testing, canary rollouts, and staged rollouts. Security features include support for mTLS encryption, authentication and authorization as well as certificate management. DevOps can also monitor service metrics including ones for latency and traffic and gain access to distributed traces and logs.

[Istio](http://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-cloud-native-network-tools) 是 CNCF 云原生中最流行的服务网格工具之一景观。它很好地集成到 Kubernetes 中，无论是在独立的 Kubernetes 环境中还是在主要云提供商的托管 Kubernetes 产品中。 Istio 使用特使代理的扩展版本，并将其与 Kubernetes 环境中的每个微服务 pod 一起部署。它具有广泛的功能集，允许 DevOps 为断路器、超时、重试、AB/测试、金丝雀推出和分阶段推出配置和创建策略。安全功能包括支持 mTLS 加密、身份验证和授权以及证书管理。 DevOps 还可以监控服务指标，包括延迟和流量指标，并获得对分布式跟踪和日志的访问权限。

#### **Consul**

Similar to Istio, [Consul](http://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-cloud-native-network-tools) from Hashicorp uses envoy as a proxy and can be easily installed in cloud native Kubernetes environments as well as managed Kubernetes offerings. Consul works by injecting a Connect sidecar (running the envoy proxy) alongside each pod in the cluster. Consul’s L7 traffic management feature supports A/B testing, Blue/Green deployments, circuit breaking, fault injections and bespoke policies to manage ingress and egress traffic. Services can be registered manually or automatically (using container orchestrators) with a dedicated registry that keeps track of all running services and their health. ACLs allow DevOps to manage authentication and authorization and secure inter-service communication. It also supports mTLS encryption and provides multiple certificate management tools including a built-in CA system. Metrics are captured for all envoy proxies in a prometheus time series which can then be graphed in Grafana. Distributed traces and logs are also supported as part of the observability feature-set.

与 Istio 类似，来自 Hashicorp 的 [Consul](http://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-cloud-native-network-tools) 使用 envoy 作为代理，可以可轻松安装在云原生 Kubernetes 环境以及托管 Kubernetes 产品中。 Consul 通过在集群中的每个 pod 旁边注入一个 Connect sidecar（运行特使代理）来工作。 Consul 的 L7 流量管理功能支持 A/B 测试、蓝/绿部署、断路器、故障注入和定制策略来管理入口和出口流量。服务可以手动或自动（使用容器编排器）注册到专用注册表，该注册表跟踪所有正在运行的服务及其健康状况。 ACL 允许 DevOps 管理身份验证和授权以及安全的服务间通信。它还支持 mTLS 加密并提供包括内置 CA 系统在内的多种证书管理工具。在普罗米修斯时间序列中捕获所有特使代理的指标，然后可以在 Grafana 中绘制图表。作为可观察性功能集的一部分，还支持分布式跟踪和日志。

#### **Kuma**

[Kuma](http://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-cloud-native-network-tools) is an open source platform agnostic service mesh from Kong, that operates equally well across multiple platforms including Kuberntes, VMs and bare metal. It uses the envoy proxy and stores all of its state and configuration in the Kubernetes API server. Kuma injects an instance of the kuma-dp sidecar container alongside each service pod. Kuma-dp in turn invokes the envoy proxy and connects to the Kuma control plane, which can be used to create and configure the service mesh. Once installed DevOps can configure routing rules for Blue/Green deployments and canary releases as well as manage communication dependencies between services. Inter-service traffic is encrypted using mTLS, which can also be used for AuthN/Z. It also provides both a native certificate authority as a well as support for multiple third party ones. DevOps can collect metrics across all data planes using Prometheus and graph it using pre-built Grafana dashboards. They can also configure policies for health checks, distributed tracing and logs. 

[Kuma](http://www.replex.io/blog/kubernetes-and-cloud-native-application-checklist-cloud-native-network-tools) 是一个来自 Kong 的开源平台不可知的服务网格，它同样运行跨多个平台，包括 Kuberntes、VM 和裸机。它使用特使代理并将其所有状态和配置存储在 Kubernetes API 服务器中。 Kuma 在每个服务 pod 旁边注入 kuma-dp sidecar 容器的一个实例。 Kuma-dp 依次调用 envoy 代理并连接到 Kuma 控制平面，可用于创建和配置服务网格。安装 DevOps 后，可以为蓝/绿部署和金丝雀版本配置路由规则，并管理服务之间的通信依赖关系。服务间流量使用 mTLS 加密，也可用于 AuthN/Z。它还提供本地证书颁发机构以及对多个第三方证书颁发机构的支持。 DevOps 可以使用 Prometheus 收集所有数据平面的指标，并使用预构建的 Grafana 仪表板绘制图表。他们还可以为健康检查、分布式跟踪和日志配置策略。