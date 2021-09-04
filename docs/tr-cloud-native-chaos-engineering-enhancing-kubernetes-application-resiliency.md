## Extending cloud native principles to chaos engineering

## 将云原生原理扩展到混沌工程

Faults are bound to happen no matter how hard you test to find them before putting your software into production – clouds and availability zones will have issues, networks will drop, and, yes, bugs will make their presence felt. Resilience is how well a system withstands such faults – a highly resilient system, for example, one built with loosely coupled microservices that can themselves be restarted and scaled easily, overcomes such faults without impacting users. Chaos Engineering is the practice of injecting faults into a system before they naturally occur. Chaos Engineering is now accepted as an essential approach for ensuring that today’s frequently changing and highly complex systems are achieving the resilience required. Through chaos engineering, unanticipated failure scenarios can be discovered and corrected before causing user issues.

无论您在将软件投入生产之前如何努力测试以找到它们，故障都一定会发生——云和可用区会出现问题，网络会掉线，而且，是的，错误会让人感觉到它们的存在。弹性是指系统承受此类故障的能力——例如，一个高度弹性的系统，由松散耦合的微服务构建而成，这些微服务本身可以轻松重启和扩展，可以在不影响用户的情况下克服此类故障。混沌工程是在故障自然发生之前将故障注入系统的实践。混沌工程现在被认为是确保当今频繁变化和高度复杂的系统实现所需弹性的基本方法。通过混沌工程，可以在导致用户问题之前发现和纠正未预料到的故障场景。

Broad adoption has made Kubernetes one of the most important platforms for software development and operations. The word “Cloud Native” is an overloaded term that has been co-opted by many traditional vendors to mean almost anything; even CNCF has allowed the use of the term cloud native to describe technologies that predate the cloud native pattern by, in some cases, decades. For the purposes of this blog, I’d like to use a more technical definition of cloud native; cloud native is here defined as an architecture where the components are microservices that are loosely coupled and, more specifically, are deployed in containers that are orchestrated by Kubernetes and related projects.

Kubernetes 的广泛采用使 Kubernetes 成为最重要的软件开发和运营平台之一。 “云原生”这个词是一个过度使用的术语，许多传统供应商都将其用于表示几乎任何东西；甚至 CNCF 也允许使用术语云原生来描述在某些情况下比云原生模式早几十年的技术。出于本博客的目的，我想使用更技术性的云原生定义；云原生在这里被定义为一种架构，其中组件是松散耦合的微服务，更具体地说，部署在由 Kubernetes 和相关项目编排的容器中。

In this blog, I would like to introduce a relatively new or less frequently used term called “Cloud Native Chaos Engineering”, defined as engineering practices focused on (and built on) Kubernetes environments, applications, microservices, and infrastructure.

在这篇博客中，我想介绍一个相对较新或不太常用的术语，称为“云原生混沌工程”，定义为专注于（并构建于）Kubernetes 环境、应用程序、微服务和基础设施的工程实践。

CNCF is, first and foremost, an open-source community (while some projects may not be strictly cloud native, they are all open-source). If Kubernetes had not been open-source, it would not have become the defacto platform for software development and operations. With that in mind, I’d like to stake the claim that Cloud Native Chaos Engineering is necessarily based on open source technologies.

CNCF 首先是一个开源社区（虽然有些项目可能不是严格的云原生项目，但它们都是开源的）。如果 Kubernetes 没有开源，它就不会成为软件开发和运营的实际平台。考虑到这一点，我想声明云原生混沌工程必然基于开源技术。

## 4 principles of a cloud native chaos engineering framework

## 云原生混沌工程框架的4个原则

1. **Open source** – The framework has to be completely open-source under the Apache2 License to encourage broader community participation and inspection. The number of applications moving to the Kubernetes platform is limitless. At such a large scale, only the Open Chaos model will thrive and get the required adoption.

   **开源** – 该框架必须在 Apache2 许可下完全开源，以鼓励更广泛的社区参与和检查。迁移到 Kubernetes 平台的应用程序数量是无限的。在如此大规模的情况下，只有开放混沌模型才能蓬勃发展并获得所需的采用。

2. **CRDs for Chaos Management** – Kubernetes native – defined here as using Kubernetes CRDs as APIs for both Developers and SREs to build and orchestrate chaos testing. The CRDs act as standard APIs to provision and manage the chaos.

   **用于混沌管理的 CRD** – Kubernetes 原生 – 此处定义为使用 Kubernetes CRD 作为 API，供开发人员和 SRE 构建和编排混沌测试。 CRD 充当标准 API 来供应和管理混乱。

3. **Extensible and pluggable** – One lesson learned why cloud native approaches are winning is that their components can be relatively easily swapped out and new ones introduced as needed. Any standard chaos library or functionality developed by other open-source developers should be able to be integrated into and orchestrated for testing via this pluggable framework. 

   **可扩展和可插拔** - 云原生方法之所以能够获胜的一个教训是，它们的组件可以相对轻松地更换，并根据需要引入新的组件。由其他开源开发人员开发的任何标准混沌库或功能都应该能够通过这个可插拔框架集成和编排进行测试。

4. **Broad Community adoption**– Once we have the APIs, Operator, and plugin framework, we have all the ingredients needed for a common way of injecting chaos. The chaos will be run against a well-known infrastructure like Kubernetes or applications like databases or other infrastructure components like storage or networking. These chaos experiments can be reused, and a broad-based community is useful for identifying and contributing to other high-value scenarios. Hence a Chaos Engineering framework should provide a central hub or forge where open-source chaos experiments are shared, and collaboration via code is enabled.

   **广泛的社区采用**——一旦我们有了 API、Operator 和插件框架，我们就拥有了注入混乱的通用方式所需的所有成分。混乱将针对众所周知的基础设施（如 Kubernetes）或应用程序（如数据库）或其他基础设施组件（如存储或网络）运行。这些混沌实验可以重复使用，一个基础广泛的社区对于识别和贡献其他高价值场景很有用。因此，混沌工程框架应该提供一个中央枢纽或锻造厂，在那里共享开源混沌实验，并通过代码实现协作。

![](https://www.cncf.io/wp-content/uploads/2020/08/image2-1-768x730-1.png)

## Introduction to Litmus

## 石蕊简介

Litmus is a cloud native chaos Engineering framework for Kubernetes. It is unique in fulfilling all 4 of the above parameters. Litmus originally started as a chaos toolset to run E2E pipelines for the CNCF SandBox project OpenEBS – powering, for example, [OpenEBS.ci](https://openebs.ci/) – and has evolved into a completely open-source framework for building and operating chaos tests on Kubernetes based systems. It consists of four main components-

Litmus 是 Kubernetes 的云原生混沌工程框架。它在满足上述所有 4 个参数方面是独一无二的。 Litmus 最初是作为一个混沌工具集来运行 CNCF SandBox 项目 OpenEBS 的 E2E 管道——例如，为 [OpenEBS.ci](https://openebs.ci/) 提供支持——并且已经发展成为一个完全开源的框架在基于 Kubernetes 的系统上构建和运行混沌测试。它由四个主要部分组成——

- Chaos CRDs or API
- Chaos Operator
- Chaos libraries and plugin framework
- Chaos Hub

- 混沌 CRD 或 API
- 混沌操作符
- 混沌库和插件框架
- 混沌中心

![](https://www.cncf.io/wp-content/uploads/2020/08/image4-1-768x320-1.png)

## Chaos API

## 混沌API

Currently, Litmus provides three APIs:

- ChaosEngine
- ChaosExperiment
- ChaosResult

目前，Litmus 提供了三个 API：

- 混沌引擎
- 混沌实验
- 混沌结果

**ChaosEngine:**

**混沌引擎：**

ChaosEngine CR is created for a given application and is tagged with appLabel. This CR ties one or more ChaosExperiments to an application.

ChaosEngine CR 是为给定的应用程序创建的，并用 appLabel 标记。此 CR 将一个或多个 ChaosExperiments 与应用程序联系起来。

**ChaosExperiment:** ChaosExperiment CR is created to hold and operate the details of actual chaos on an application. It defines the type of experiment and key parameters of the experiment.

**ChaosExperiment：**创建 ChaosExperiment CR 是为了保存和操作应用程序上实际混乱的详细信息。它定义了实验的类型和实验的关键参数。

**ChaosResult:** ChaosResult CR is created by the operator after an experiment is run. One ChaosResult CR is maintained per ChaosEngine. The ChaosResult CR is useful in making sense of a given ChaosExperiment. This CR is used for generating chaos analytics which can be extremely useful – for example when certain components are upgraded between the chaos experiments, and the results need to be easily compared

**ChaosResult：** ChaosResult CR 由操作员在实验运行后创建。每个 ChaosEngine 维护一个 ChaosResult CR。 ChaosResult CR 有助于理解给定的 ChaosExperiment。此 CR 用于生成非常有用的混沌分析 - 例如，当某些组件在混沌实验之间升级时，需要轻松比较结果

## Chaos Operator

## 混沌Operator

The Litmus Operator is implemented using the Operator-SDK. This operator manages the lifecycle of the chaos CRs. The lifecycle of Litmus itself can be managed using this operator as it follows the lifecycle management API requirements. The chaos operator is also available at [operatorhub.io](https://operatorhub.io/operator/litmuschaos)

Litmus Operator 是使用 Operator-SDK 实现的。该操作员管理混乱 CR 的生命周期。 Litmus 本身的生命周期可以使用此运算符进行管理，因为它遵循生命周期管理 API 要求。 [operatorhub.io](https://operatorhub.io/operator/litmuschaos) 也提供了混沌运算符

## Chaos libraries and external plugins

## 混沌库和外部插件

The actual injection of chaos is done by chaos libraries or chaos executors. For example, the Litmus project has already built a chaos library called “LitmusLib”. LitmusLib is aware of how to kill a pod, how to introduce a CPU hog, how to hog memory or how to kill a node, and several other faults and degradations. Like LitmusLib, there are other open-source chaos projects like Pumba or PowerfulSeal. The CNCF landscape has more details of various chaos engineering projects. As shown below, the Litmus plugin framework allows other chaos projects to make use of Litmus for chaos orchestration. For example, one can create a chaos chart for the pod-kill experiment using Pumba or PowerfulSeal and execute it via the Litmus framework.

混沌的实际注入是由混沌库或混沌执行器完成的。例如，Litmus 项目已经构建了一个名为“LitmusLib”的混沌库。 LitmusLib 知道如何杀死 pod、如何引入 CPU hog、如何占用内存或如何杀死节点，以及其他一些故障和降级。与 LitmusLib 一样，还有其他开源混沌项目，如 Pumba 或 PowerfulSeal。 CNCF 版图有更多关于各种混沌工程项目的细节。如下图所示，Litmus 插件框架允许其他混沌项目利用 Litmus 进行混沌编排。例如，您可以使用 Pumba 或 PowerfulSeal 为 pod-kill 实验创建一个混沌图，并通过 Litmus 框架执行它。

![](https://www.cncf.io/wp-content/uploads/2020/08/image1-1-768x718-1.png)

_\\* PowerfulSeal and Pumba are shown as examples._

_\\* 以 PowerSeal 和 Pumba 为例。_

## Chaos Hub

## 混沌中心

Chaos charts are located at [hub.litmuschaos.io](https://hub.litmuschaos.io/). ChaosHub brings all the reusable chaos experiments together. Application developers and SRE share their chaos experiences for others to reuse. The goal of the hub is to have the developers share the failure tests that they are using to validate their applications in CI pipelines to their users, who are typically SREs.

混沌图表位于 [hub.litmuschaos.io](https://hub.litmuschaos.io/)。 ChaosHub 将所有可重用的混沌实验结合在一起。应用程序开发人员和 SRE 分享他们的混沌经验以供其他人重用。该中心的目标是让开发人员将他们用来验证 CI 管道中的应用程序的失败测试分享给他们的用户，这些用户通常是 SRE。

![](https://www.cncf.io/wp-content/uploads/2020/08/image5-1-768x384-1.png)

Currently, the chaos hub contains charts for Kubernetes chaos and OpenEBS chaos. We expect to receive more contributions from the community going forward.

目前，混乱中心包含 Kubernetes 混乱和 OpenEBS 混乱的图表。我们希望在未来收到更多来自社区的贡献。

## Example use cases of Litmus: 

## Litmus 的示例用例：

The most simple use case of Litmus is application developers using Litmus in the development phase itself. Chaos Engineering has been limited to the Production environment, and lately, we are seeing this practice being adopted in CI pipelines. But with Litmus, chaos testing is possible during development as well. Like Unit Testing, Integration Testing, and Behavior-Driven Testing, Chaos Testing is a test philosophy for developers to carry out the negative test scenarios to test the resiliency of the code before the code is merged to the repository. Chaos testing can be appended very easily to the application, as shown below:

![](https://www.cncf.io/wp-content/uploads/2020/08/dev-litmus-1-1.gif)

Other use cases of Litmus are for inducing chaos in CI pipelines and production environments.

Litmus 最简单的用例是应用程序开发人员在开发阶段本身使用 Litmus。混沌工程仅限于生产环境，最近，我们看到在 CI 管道中采用了这种做法。但是使用 Litmus，在开发过程中也可以进行混沌测试。与单元测试、集成测试和行为驱动测试一样，混沌测试是一种测试哲学，供开发人员在代码合并到存储库之前执行负面测试场景以测试代码的弹性。混沌测试可以很容易地附加到应用程序中，如下所示：

Litmus 的其他用例是在 CI 管道和生产环境中引起混乱。

## Summary

##  概括

With the introduction of chaos operator, chaos CRDs, and chaos hub, Litmus has all the key ingredients of cloud native Chaos Engineering.

随着混沌算子、混沌 CRD 和混沌中心的引入，Litmus 拥有云原生混沌工程的所有关键要素。

### Important links:

GitHub: [github.com/litmuschaos](https://github.com/litmuschaos/litmus)

### 重要链接：

GitHub：[github.com/litmuschaos](https://github.com/litmuschaos/litmus)

Twitter: [@litmuschaos](https://twitter.com/litmuschaos)

推特：[@litmuschaos](https://twitter.com/litmuschaos)

Chaos Charts: [hub.litmuschaos.io](https://hub.litmuschaos.io/)

混沌图表：[hub.litmuschaos.io](https://hub.litmuschaos.io/)

Community Slack: [#litmus channel on K8S Slack](https://kubernetes.slack.com/messages/CNXNB0ZTN) 

社区 Slack：[K8S Slack 上的#litmus 频道](https://kubernetes.slack.com/messages/CNXNB0ZTN)

