# The Cloud Native Landscape: Platforms Explained

# 云原生景观：平台解释

#### 17 Mar 2021 12:08pm,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack. io/author/jason-morgan/ "Posts by Jason Morgan")

#### 2021 年 3 月 17 日下午 12:08，作者：[Catherine Paganini](https://thenewstack.io/author/catherine-paganini/“Catherine Paganini 的帖子”) 和 [Jason Morgan](https://thenewstack.io/author/jason-morgan/“杰森摩根的帖子”)

![](https://cdn.thenewstack.io/media/2021/03/2244406d-background-1126047_640.jpg)

_This post is part of an ongoing series from the Cloud Native Computing Foundation's_ [_Business Value Subcommittee_](https://lists.cncf.io/g/cncf-business-value) _co-chairs_ [_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) _and_ [_Jason Morgan_](https://thenewstack.io/author/jason-morgan/) _that focuses on explaining each category of the [cloud native landscape](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/) to a non-technical audience as well as engineers just getting started with cloud native. See also installments on the layers for [application definition development](https://thenewstack.io/the-cloud-native-landscape-the-application-definition-and-development-layer/), the [runtime](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/), the [orchestration and management](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/), and the [provisioning.](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)_

_这篇文章是云原生计算基金会正在进行的系列文章的一部分_[_商业价值小组委员会_](https://lists.cncf.io/g/cncf-business-value)_co-chairs_[_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) _and_ [_Jason Morgan_](https://thenewstack.io/author/jason-morgan/)_重点讲解[云原生Landscape](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/) 面向非技术受众以及刚开始使用云原生的工程师。另请参阅有关[应用程序定义开发](https://thenewstack.io/the-cloud-native-landscape-the-application-definition-and-development-layer/)、[运行时](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)，[编排和管理](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/) 和 [provisioning.](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)_

There isn’t anything inherently new in these platforms. Everything they do can be done by one of the tools in these layers or the observability and analysis column. You could certainly build your own platform, and in fact, many organizations do. However, configuring and fine-tuning the different modules reliably and securely while ensuring that all technologies are always updated and vulnerabilities patched is no easy task—you’ll need a dedicated team to build and maintain it. If you don’t have the required bandwidth and/or know-how, your team is likely better off with a platform. For some organizations, especially those with small engineering teams, platforms are the only way to adopt a cloud native approach.

在这些平台中没有任何本质上的新东西。他们所做的一切都可以通过这些层中的工具之一或可观察性和分析列来完成。您当然可以构建自己的平台，事实上，许多组织都这样做。但是，在确保所有技术始终更新并修补漏洞的同时可靠、安全地配置和微调不同的模块并非易事——您需要一个专门的团队来构建和维护它。如果您没有所需的带宽和/或专业知识，您的团队可能会更好地使用平台。对于某些组织，尤其是那些拥有小型工程团队的组织，平台是采用云原生方法的唯一途径。

You'll probably notice, all platforms [revolve around Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-care/). That’s because Kubernetes, is at the core of the cloud native stack.

您可能会注意到，所有平台 [围绕 Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-care/)。这是因为 Kubernetes 是云原生堆栈的核心。

### Sidenote

###  边注

When looking at the [Cloud Native Landscape](https://landscape.cncf.io/), you’ll note a few distinctions:

在查看 [云原生景观](https://landscape.cncf.io/) 时，您会注意到一些区别：

- Projects in large boxes are Cloud Native Computing Foundation-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- Projects in small white boxes are open source projects.
- Products in gray boxes are proprietary.

- 大盒子中的项目是云原生计算基金会托管的开源项目。有的还在孵化阶段（浅蓝色/紫色框），有的则是毕业项目（深蓝色框）。
- 小白盒中的项目是开源项目。
- 灰色框中的产品是专有的。

Please note that even during the time of this writing, we saw new projects becoming part of the Cloud Native Computing Foundation (CNCF) so always refer to the actual landscape — things are moving fast!

请注意，即使在撰写本文时，我们也看到新项目成为云原生计算基金会 (CNCF) 的一部分，因此请始终参考实际情况——事情进展很快！

## Kubernetes Distributions

## Kubernetes 发行版

### What It Is

###  这是什么

A distribution, or distro, is when a vendor takes core Kubernetes — that’s the unmodified, open source code (although some modify it) — and packages it for redistribution. Usually, that entails finding and validating the Kubernetes software and providing a mechanism handling cluster installation and upgrades. Many Kubernetes distributions include other proprietary or open source applications.

发行版或发行版是指供应商采用核心 Kubernetes——即未经修改的开源代码（尽管有些人对其进行了修改）——并将其打包以进行重新分发。通常，这需要查找和验证 Kubernetes 软件并提供处理集群安装和升级的机制。许多 Kubernetes 发行版包括其他专有或开源应用程序。

### What’s the Problem They Solve

### 他们解决的问题是什么

[Open source Kubernetes](https://github.com/kubernetes/kubernetes) doesn’t specify a particular installation tool and leaves many setup configuration choices to the user. Additionally, there is limited support for issues as they arise beyond community resources like [Community Forums](https://discuss.kubernetes.io/),[StackOverflow](http://stackoverflow.com/questions/tagged/kubernetes) , [Slack](https://slack.k8s.io/) or Discord. 

[开源 Kubernetes](https://github.com/kubernetes/kubernetes) 没有指定特定的安装工具，并为用户留下了许多设置配置选择。此外，对于超出社区资源（例如 [社区论坛](https://discuss.kubernetes.io/)、[StackOverflow](http://stackoverflow.com/questions/tagged/kubernetes)的问题的支持有限) , [Slack](https://slack.k8s.io/) 或 Discord。

While using Kubernetes has gotten easier over time, it can be challenging to find and use the open source installers. Users need to understand what versions to use, where to get them, and if a particular component is compatible with another. They also need to decide what software will be deployed to their clusters and what settings to use to ensure their platforms are secure, stable, and performant. All this requires deep Kubernetes expertise that may not be readily available in-house.

随着时间的推移，使用 Kubernetes 变得越来越容易，但找到和使用开源安装程序可能具有挑战性。用户需要了解要使用的版本、从何处获取以及特定组件是否与另一个组件兼容。他们还需要决定将哪些软件部署到他们的集群以及使用哪些设置来确保他们的平台安全、稳定和高性能。所有这一切都需要深厚的 Kubernetes 专业知识，而这些专业知识在内部可能并不容易获得。

### How It Helps

### 它如何帮助

Kubernetes distributions provide a trusted and reliable way to install Kubernetes and provide opinionated defaults that create a better and more secure operating environment. A Kubernetes distribution gives vendors and projects the control and predictability they need to provide support for a customer as they go through the lifecycle of deploying, maintaining, and upgrading their Kubernetes clusters.

Kubernetes 发行版提供了一种值得信赖且可靠的 Kubernetes 安装方式，并提供了自以为是的默认设置，以创建更好、更安全的操作环境。 Kubernetes 发行版为供应商和项目提供了他们在部署、维护和升级其 Kubernetes 集群的生命周期中为客户提供支持所需的控制和可预测性。

That predictability enables distribution providers to support users when they have production issues. Distributions also often provide a tested and supported upgrade path that allows users to keep their Kubernetes clusters up to date. Additionally, distributions often provide software to deploy on top of Kubernetes that makes it easier to use.

这种可预测性使分销商能够在用户遇到生产问题时为他们提供支持。发行版通常还提供经过测试和支持的升级路径，允许用户保持他们的 Kubernetes 集群是最新的。此外，发行版通常会提供部署在 Kubernetes 之上的软件，使其更易于使用。

Distributions significantly ease and speed up Kubernetes adoption. Since the expertise needed to configure and fine-tune the clusters is coded into the platform, organizations can get up and running with cloud native tools without having to hire additional engineers with specialized expertise.

发行版显着简化并加快了 Kubernetes 的采用。由于配置和微调集群所需的专业知识已编码到平台中，因此组织可以使用云原生工具启动和运行，而无需聘请具有专业知识的额外工程师。

### Technical 101

### 技术 101

If you’ve installed Kubernetes, you’ve likely used something like kubeadm to get your cluster up and running. Even then, you probably had to decide on a CNI (Container Network Interface), install, and configure it. Then, you might have added some storage classes, a tool to handle log messages, maybe an ingress controller, and the list goes on. A Kubernetes distribution will automate some, or all, of that setup. It will also ship with configuration settings based on its own interpretation of best practice or an intelligent default. Additionally, most distributions will come with some extensions or add-ons bundled and tested to ensure you can get going with your new cluster as quickly as possible.

如果您已经安装了 Kubernetes，那么您可能已经使用了 kubeadm 之类的东西来启动和运行集群。即便如此，您可能也必须决定 CNI（容器网络接口）、安装和配置它。然后，您可能添加了一些存储类、一个处理日志消息的工具、一个入口控制器，等等。 Kubernetes 发行版将自动化部分或全部设置。它还将根据自己对最佳实践的解释或智能默认值提供配置设置。此外，大多数发行版都会附带一些扩展或附加组件，并进行了捆绑和测试，以确保您可以尽快开始使用新集群。

Let’s take [Kublr](https://kublr.com/) as an example. With Kubernetes at its core, this platform bundles technologies from mainly three layers: provisioning, runtime, orchestration and management, and observability and analysis. All modules are preconfigured with a few options to choose from and ready to go.Different platforms have different focal points. In the case of Kublr, the focus is more on the operations side, while other platforms may focus more on developer tooling.

我们以 [Kublr](https://kublr.com/) 为例。该平台以 Kubernetes 为核心，主要从三个层面捆绑技术：供应、运行时、编排和管理以及可观察性和分析。所有模块都预先配置了一些可供选择的选项，随时可用。不同的平台有不同的重点。就 Kublr 而言，重点更多地放在运营方面，而其他平台可能更侧重于开发人员工具。

There are a lot of options in this category. As of this writing, [k3s](https://k3s.io) is the only CNCF project. There are a lot of great open source and commercial options available, including, Microk8s from Canonical, k3s, Tanzu Kubernetes Grid from [VMware](https://tanzu.vmware.com?utm_content=inline-mention), Docker Enterprise from [ Mirantis](https://www.mirantis.com/?utm_content=inline-mention), Rancher from Suse, and of course [Red Hat](https://www.openshift.com/try?utm_content=inline-mention)'s Openshift. We didn’t have time to mention even close to half the Kubernetes distributions, and we encourage you to think carefully about your needs when you begin evaluating distributions.

这个类别有很多选择。在撰写本文时，[k3s](https://k3s.io) 是唯一的 CNCF 项目。有很多很棒的开源和商业选项可用，包括 Canonical 的 Microk8s、k3s、[VMware] 的 Tanzu Kubernetes Grid(https://tanzu.vmware.com?utm_content=inline-mention)、[VMware] 的 Docker Enterprise。 Mirantis](https://www.mirantis.com/?utm_content=inline-mention)，来自 Suse 的 Rancher，当然还有 [Red Hat](https://www.openshift.com/try?utm_content=inline-mention) 的 Openshift。我们甚至没有时间提及将近一半的 Kubernetes 发行版，我们鼓励您在开始评估发行版时仔细考虑您的需求。

![Kubernetes distributions](https://cdn.thenewstack.io/media/2021/03/342aedc7-screen-shot-2021-03-16-at-9.18.35-am.png)

## Hosted Kubernetes

## 托管 Kubernetes

### What It Is 

###  这是什么

Hosted Kubernetes is a service offered by infrastructure providers like [Amazon Web Services](https://aws.amazon.com/?utm_content=inline-mention)(AWS), DigitalOcean, Azure, or Google, allowing customers to spin up a Kubernetes cluster on-demand. The cloud provider takes responsibility for managing part of the Kubernetes cluster, usually called the control plane. They are similar to distributions but managed by the cloud provider on _their_ infrastructure.

托管 Kubernetes 是由 [Amazon Web Services](https://aws.amazon.com/?utm_content=inline-mention)(AWS)、DigitalOcean、Azure 或 Google 等基础设施提供商提供的服务，允许客户启动Kubernetes 集群按需。云提供商负责管理 Kubernetes 集群的一部分，通常称为控制平面。它们类似于发行版，但由云提供商在他们的基础设施上管理。

### What’s the Problem They Solve

### 他们解决的问题是什么

Hosted Kubernetes allows teams to get started with Kubernetes without knowing or doing anything beyond setting up an account with a cloud vendor. It solves four of the five Ws of getting started with Kubernetes. Who (manages it): your cloud provider; what: their hosted Kubernetes offering; when: now; and where: on the cloud providers infrastructure. The why is up to you.

托管 Kubernetes 允许团队开始使用 Kubernetes，除了在云供应商处设置帐户之外，无需知道或做任何事情。它解决了 Kubernetes 入门的五个 W 中的四个。谁（管理它）：您的云提供商；什么：他们托管的 Kubernetes 产品；什么时候：现在；在哪里：在云提供商基础设施上。原因取决于你。

### How It Helps

### 它如何帮助

Since the provider takes care of all management details, hosted Kubernetes is the easiest way to get started with cloud native. All users have to do, is develop their apps and deploy them on the hosted Kubernetes services — it’s incredibly convenient. The hosted offering allows users to spin up a Kubernetes cluster and get started right away,\* while taking some responsibility for the cluster availability. It’s worth noting that with the extra convenience of these services comes some reduced flexibility. The offering is bound to the cloud provider, and Kubernetes users don’t have access to the Kubernetes control plane, so some configuration options are limited.

由于提供商负责所有管理细节，托管 Kubernetes 是开始使用云原生的最简单方法。所有用户要做的就是开发他们的应用程序并将它们部署在托管的 Kubernetes 服务上——这非常方便。托管产品允许用户启动 Kubernetes 集群并立即开始，\* 同时对集群可用性承担一些责任。值得注意的是，随着这些服务的额外便利，灵活性有所降低。该产品与云提供商绑定，Kubernetes 用户无权访问 Kubernetes 控制平面，因此某些配置选项受到限制。

_\\* Slight exception for EKS from AWS as it also requires users to take some additional steps to prepare their clusters._

_\\* AWS EKS 的轻微例外，因为它还需要用户采取一些额外的步骤来准备他们的集群。_

### Technical 101

### 技术 101

Hosted Kubernetes are on-demand Kubernetes clusters provided by a vendor, usually an infrastructure hosting provider. The vendor takes responsibility for provisioning the cluster and managing the Kubernetes control plane. Again, the notable exception is EKS, where individual node provisioning is left up to the client.

托管 Kubernetes 是由供应商（通常是基础架构托管提供商）提供的按需 Kubernetes 集群。供应商负责配置集群和管理 Kubernetes 控制平面。同样，值得注意的例外是 EKS，其中单独的节点供应由客户端决定。

Hosted Kubernetes allows an organization to quickly provision new clusters and reduce their operational risk by outsourcing infrastructure component management to another organization. The main trade-offs are that you'll likely be charged for the control plane management (GKE ran into a [bit of controversy](https://www.theregister.com/2020/03/05/google_reintroduces_management_fee_for_kubernetes_clusters/) around price changes last year) and that you'll be limited in what you can do. Managed clusters provide stricter limits on configuring your Kubernetes cluster than DIY Kubernetes clusters.

托管 Kubernetes 允许组织通过将基础架构组件管理外包给另一个组织来快速配置新集群并降低其运营风险。主要的权衡是您可能需要为控制平面管理付费（GKE 遇到了 [有点争议](https://www.theregister.com/2020/03/05/google_reintroduces_management_fee_for_kubernetes_clusters/)围绕价格去年发生了变化)，并且您可以做的事情会受到限制。与 DIY Kubernetes 集群相比，托管集群对 Kubernetes 集群的配置提供了更严格的限制。

There are numerous vendors and projects in this space and, at the time of this writing, no CNCF projects.

这个领域有许多供应商和项目，在撰写本文时，还没有 CNCF 项目。

![Hosted Kubernetes](https://cdn.thenewstack.io/media/2021/03/a061e08f-screen-shot-2021-03-16-at-9.21.08-am.png)

## Kubernetes Installer

## Kubernetes 安装程序

### What It Is

###  这是什么

Kubernetes installers help install Kubernetes on a machine. They automate the Kubernetes installation and configuration process and may even help with upgrades. Kubernetes installers are often coupled with or used by Kubernetes distributions or hosted Kubernetes offerings.

Kubernetes 安装程序帮助在机器上安装 Kubernetes。他们自动执行 Kubernetes 安装和配置过程，甚至可以帮助升级。 Kubernetes 安装程序通常与 Kubernetes 发行版或托管 Kubernetes 产品耦合或使用。

### What’s the Problem They Solve

### 他们解决的问题是什么

Similar to Kubernetes distributions, Kubernetes installers simplify getting started with Kubernetes. [Open source Kubernetes](https://github.com/kubernetes/kubernetes) relies on installers like kubeadm, which, as of this writing, is part of the Certified Kubernetes Administrator certification test to get Kubernetes clusters up and running.

与 Kubernetes 发行版类似，Kubernetes 安装程序简化了 Kubernetes 的入门。 [开源 Kubernetes](https://github.com/kubernetes/kubernetes) 依赖于像 kubeadm 这样的安装程序，在撰写本文时，它是 Kubernetes 管理员认证测试的一部分，用于启动和运行 Kubernetes 集群。

### How It Helps

### 它如何帮助

Kubernetes installers ease the Kubernetes installation process. Like distributions, they provide a vetted source for the source code and version. They also often ship with opinionated Kubernetes environment configurations. Kubernetes installers like-kind (Kubernetes in Docker) allow you to get a Kubernetes cluster with a single command.

Kubernetes 安装程序简化了 Kubernetes 安装过程。与发行版一样，它们为源代码和版本提供经过审查的来源。它们还经常附带固执的 Kubernetes 环境配置。 Kubernetes 安装程序（如 Docker 中的 Kubernetes）允许您使用单个命令获取 Kubernetes 集群。

### Technical 101 

### 技术 101

Whether you're installing Kubernetes locally on Docker, spinning up and provisioning new virtual machines, or preparing new physical servers, you're going to need a tool to handle all the preparation of various Kubernetes components (unless you're looking to do it [the hard way](https://github.com/kelseyhightower/kubernetes-the-hard-way)).

无论您是在 Docker 上本地安装 Kubernetes、启动和配置新的虚拟机，还是准备新的物理服务器，您都需要一个工具来处理各种 Kubernetes 组件的所有准备工作（除非您打算这样做） [艰难的方式](https://github.com/kelseyhightower/kubernetes-the-hard-way))。

Kubernetes installers simplify that process. Some handle spinning up nodes and others merely configure nodes you’ve already provisioned. They all offer various levels of automation and are each suited for different use cases. When getting started with an installer, start by understanding your needs, then pick an installer that addresses them. At the time of this writing, kubeadm is considered so fundamental to the Kubernetes ecosystem that it’s included as part of the CKA, certified Kubernetes administrator exam. Minikube, kind, kops, and kubespray are all CNCF-owned Kubernetes installer projects.

Kubernetes 安装程序简化了这个过程。一些处理旋转节点，而另一些仅配置您已经配置的节点。它们都提供不同级别的自动化，并且每个都适用于不同的用例。开始使用安装程序时，首先要了解您的需求，然后选择一个可以解决这些需求的安装程序。在撰写本文时，kubeadm 被认为是 Kubernetes 生态系统的基础，因此它被包含在 CKA 认证的 Kubernetes 管理员考试中。 Minikube、kind、kops 和 kubespray 都是 CNCF 拥有的 Kubernetes 安装程序项目。

![Kubernetes installer](https://cdn.thenewstack.io/media/2021/03/2891980c-screen-shot-2021-03-16-at-9.24.16-am.png)

## PaaS / Container Service

## PaaS / 容器服务

### What It Is

###  这是什么

A [platform as a service](https://en.wikipedia.org/wiki/Platform_as_a_service), or PaaS, is an environment that allows users to run applications without necessarily understanding or knowing about the underlying compute resources. PaaS and container services in this category are mechanisms to either host a PaaS for developers or host services they can use.

[平台即服务](https://en.wikipedia.org/wiki/Platform_as_a_service) 或 PaaS 是一种环境，它允许用户运行应用程序而不必了解或了解底层计算资源。此类别中的 PaaS 和容器服务是为开发人员托管 PaaS 或托管他们可以使用的服务的机制。

### What’s the Problem They Solve

### 他们解决的问题是什么

In this series, we’ve talked a lot about the tools and technologies around “cloud native.” A PaaS attempts to connect many of the technologies found in this landscape in a way that provides direct value to developers. It answers the following questions: how will I run applications in various environments and, once running, how will my team and users interact with them?

在本系列中，我们围绕“云原生”讨论了很多工具和技术。 PaaS 试图以一种为开发人员提供直接价值的方式连接在这个领域中发现的许多技术。它回答了以下问题：我将如何在各种环境中运行应用程序，一旦运行，我的团队和用户将如何与它们交互？

### How It Helps

### 它如何帮助

PaaS provides opinions and choices around how to piece together the various open and closed source tools needed to run applications. Many offerings include tools that handle PaaS installation and upgrades and the mechanisms to convert application code into a running application. Additionally, PaaS handle the runtime needs of application instances, including on-demand scaling of individual components and visibility into performance and log messages of individual apps.

PaaS 提供了关于如何将运行应用程序所需的各种开源和闭源工具组合在一起的意见和选择。许多产品包括处理 PaaS 安装和升级的工具以及将应用程序代码转换为正在运行的应用程序的机制。此外，PaaS 处理应用程序实例的运行时需求，包括按需扩展单个组件以及查看单个应用程序的性能和日志消息。

### Technical 101

### 技术 101

Organizations are adopting cloud native technologies to achieve specific business or organizational objectives. A PaaS provides a quicker path to value than building a custom application platform. Tools like Heroku or Cloud Foundry Application Runtime help organizations get up and running with new applications quickly. They excel at providing the tools needed to run [12 factor](https://12factor.net/) or cloud native applications.

组织正在采用云原生技术来实现特定的业务或组织目标。 PaaS 提供了比构建自定义应用程序平台更快的价值途径。 Heroku 或 Cloud Foundry Application Runtime 等工具可帮助组织快速启动并运行新应用程序。他们擅长提供运行 [12 factor](https://12factor.net/) 或云原生应用程序所需的工具。

Any PaaS comes with its own set of trade-offs and restrictions. Most only work with a subset of languages or application types and the opinions and decisions baked into them may or may not be a good fit for your needs. Stateless applications tend to do very well in a PaaS but stateful applications like databases usually don’t. There are currently no CNCF projects in this space but most of the offerings are open source and Cloud Foundry is managed by the Cloud Foundry Foundation.

任何 PaaS 都有自己的一套权衡和限制。大多数只适用于语言或应用程序类型的子集，并且融入其中的意见和决定可能适合也可能不适合您的需求。无状态应用程序在 PaaS 中往往做得很好，但像数据库这样的有状态应用程序通常不会。目前在这个领域没有 CNCF 项目，但大多数产品都是开源的，Cloud Foundry 由 Cloud Foundry 基金会管理。

![Kubernetes PaaS / Container as a Service](https://cdn.thenewstack.io/media/2021/03/a12a4c10-screen-shot-2021-03-16-at-9.26.50-am.png)

## Conclusion

##  结论

As we’ve seen there are multiple tools that help ease Kubernetes adoption. From Kubernetes distributions and hosted Kubernetes to more barebones installers or PaaS, they all take some of the installation and configuration burden and pre-package it for you. Each solution comes with its own “flavor.” Vendor opinions about what’s important and appropriate are built into the solution. 

正如我们所见，有多种工具可以帮助简化 Kubernetes 的采用。从 Kubernetes 发行版和托管 Kubernetes 到更多准系统安装程序或 PaaS，它们都承担了一些安装和配置负担，并为您预先打包。每个解决方案都有自己的“风味”。供应商关于什么是重要的和合适的意见被纳入解决方案。

Before adopting any of these, you’ll need to do some research to identify the best solution for your particular use case. Will you likely encounter advanced Kubernetes scenarios where you’ll need control over the control plane? Then, hosted solutions are likely not a good fit. Do you have a small team that manages “standard” workloads and need to offload as many operational tasks as possible? Then, hosted solutions may be a great fit. Is portability important? What about production-readiness? There are multiple aspects to consider. There is no “one best tool,” but there certainly is an optimal tool for your use case. Hopefully, this article will help you narrow your search down to the right “bucket.”

在采用其中任何一个之前，您需要进行一些研究以确定适合您特定用例的最佳解决方案。您可能会遇到需要控制控制平面的高级 Kubernetes 场景吗？那么，托管解决方案可能不太适合。您是否有一个管理“标准”工作负载并需要卸载尽可能多的操作任务的小团队？然后，托管解决方案可能非常适合。便携性重要吗？生产准备情况如何？有多个方面需要考虑。没有“最好的工具”，但肯定有适合您的用例的最佳工具。希望本文能帮助您将搜索范围缩小到正确的“桶”。

This concludes the platform “column” of the CNCF landscape. Next, we’ll tackle the last article of this series, the observability and analysis “column.”

CNCF 景观的平台“专栏”到此结束。接下来，我们将讨论本系列的最后一篇文章，即可观察性和分析“专栏”。

_As always, a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._

_一如既往，非常感谢来自 CNCF 的 [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/)，他非常友好地审阅了这篇文章，以确保其准确无误。_

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Mirantis, Docker, Bit.

The New Stack 是 Insight Partners 的全资子公司。 TNS 所有者 Insight Partners 是以下公司的投资者：Mirantis、Docker、Bit。

Amazon Web Services, CNCF, Mirantis, Red Hat and VMware are sponsors of The New Stack.

Amazon Web Services、CNCF、Mirantis、Red Hat 和 VMware 是 The New Stack 的赞助商。

Image par [kalhh](https://pixabay.com/fr/users/kalhh-86169/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1126047) de [Pixabay](https://pixabay.com/fr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1126047). 

图片标准 [kalhh](https://pixabay.com/fr/users/kalhh-86169/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1126047) de [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1126047)。

