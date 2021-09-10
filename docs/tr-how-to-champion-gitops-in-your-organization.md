# How to Champion GitOps in Your Organization

# 如何在您的组织中支持 GitOps

August 04, 2020

By now you’ve heard all about GitOps and are convinced that GitOps is the most efficient way for development teams to go faster without them having to become Kubernetes gurus. However, making the switch to a cloud native technical solution may be the simplest part in your Kubernetes journey. Getting buy-in from your peers and championing GitOps throughout your organization could very well be the most challenging aspect of the cloud native transition.

到目前为止，您已经听说了关于 GitOps 的所有信息，并且确信 GitOps 是开发团队无需成为 Kubernetes 专家即可加快速度的最有效方式。但是，切换到云原生技术解决方案可能是您 Kubernetes 旅程中最简单的部分。获得同行的认可并在整个组织中倡导 GitOps 很可能是云原生过渡中最具挑战性的方面。

At a recent GitOps Days virtual event held earlier this spring, we hosted a roundtable discussion with four Kubernetes and GitOps hands-on practitioners. Most of our panelists had recently implemented GitOps and self-service development platforms for Kubernetes in their organizations. In this discussion, the panelists offer a lot of common sense advice on the pitfalls to avoid when engineering a platform, and more importantly, they dig down on some practical strategies to use when you're on-boarding and educating development teams who are adopting GitOps for self-service developer platforms in your own organization.

在今年春季早些时候举行的最近 GitOps Days 虚拟活动中，我们与四位 Kubernetes 和 GitOps 实践从业者举行了圆桌讨论。我们的大多数小组成员最近都在他们的组织中为 Kubernetes 实施了 GitOps 和自助服务开发平台。在本次讨论中，小组成员提供了许多关于在设计平台时要避免的陷阱的常识性建议，更重要的是，他们深入研究了一些实用策略，以便在您入职和教育正在采用的开发团队时使用用于您自己组织中的自助开发人员平台的 GitOps。

## Our panel of GitOps and Kubernetes practitioners

## 我们的 GitOps 和 Kubernetes 从业者小组

We were very fortunate to host Niraj Amin, Director, Cloud Platform Architect at Fidelity Investments, Steve Wade ( [@swade1987](https://twitter.com/swade1987)), the Platform Lead at Mettle, Javeria Khan ( [@javeriak \_](https://twitter.com/javeriak_)), Senior SRE at Palo Alto Networks, and Kingdon Barrett ( [@yebyen](https://twitter.com/yebyen)), Application Developer at University of Notre Dame OIT. Cornelia Davis, Weaveworks, CTO moderated the panel.

我们非常幸运地接待了 Niraj Amin，Fidelity Investments 的云平台架构师总监，Steve Wade ([@swade1987](https://twitter.com/swade1987))，Mettle 的平台负责人，Javeria Khan ([@javeriak \_](https://twitter.com/javeriak_))，Palo Alto Networks 的高级 SRE，以及 Kingdon Barrett ([@yebyen](https://twitter.com/yebyen))，圣母大学的应用程序开发人员OIT 夫人。 Weaveworks 的首席技术官 Cornelia Davis 主持了小组讨论。

## What do we mean by a self-service developer platform?

## 自助开发者平台是什么意思？

Before we discuss how you go about championing GitOps in your own organization, let’s step back and talk about what we even mean when we say “self-service Kubernetes platform”.

在我们讨论如何在自己的组织中倡导 GitOps 之前，让我们退后一步，谈谈我们所说的“自助服务 Kubernetes 平台”的含义。

Platform teams in an organization are increasingly responsible for providing a set of developer services to developers. Developers on the other hand are responsible for delivering applications to the company’s end-consumers. And so, to a large extent what we’re exploring is the relationship between those teams within an organization.

组织中的平台团队越来越负责向开发人员提供一组开发人员服务。另一方面，开发人员负责向公司的最终消费者交付应用程序。因此，在很大程度上，我们正在探索的是组织内这些团队之间的关系。

### Guardrails and security in place

### 护栏和安全到位

At Fidelity, a dedicated platform team manages their Kubernetes implementation. As a platform team they serve the needs of developers. In particular their job is to get out of the way of the developer so they can do their job as efficiently as possible.

在 Fidelity，一个专门的平台团队负责管理他们的 Kubernetes 实施。作为一个平台团队，他们满足开发人员的需求。特别是他们的工作是避开开发人员，以便他们尽可能高效地完成工作。

Niraj clarified, “...when I talk about platforms, what we’re really talking about is the infrastructure component of things. Obviously both EKS and Kubernetes play a big role. We also have fifteen or sixteen different components like the ELB ingress controller, external DNS and other components that are open source, plus we provide the autoscalers to developers as well.”

Niraj 澄清说：“……当我谈论平台时，我们真正谈论的是事物的基础设施组件。显然 EKS 和 Kubernetes 都发挥着重要作用。我们还有 15 到 16 个不同的组件，例如 ELB 入口控制器、外部 DNS 和其他开源组件，此外我们还向开发人员提供自动缩放器。”

At Fidelity, GitOps allows for all platform configurations to be bundled and versioned in Git. Developers never have to worry about operational tasks, as upgrades are now automated. The result is a standard platform with guardrails and security in place that any development team can deploy to and from.

在 Fidelity，GitOps 允许在 Git 中捆绑所有平台配置并对其进行版本控制。开发人员永远不必担心操作任务，因为升级现在是自动化的。结果是一个带有护栏和安全性的标准平台，任何开发团队都可以部署到位。

Similarly at Palo Alto, Javeria Khan, Senior SRE says that “as platform builders we are trying to solve infrastructure issues and remove that burden from the developer. We especially want to ensure that they don’t inadvertently make a change that can cause harm to the entire system.“

同样在帕洛阿尔托，高级 SRE 的 Javeria Khan 说：“作为平台构建者，我们正在努力解决基础设施问题并减轻开发人员的负担。我们特别希望确保他们不会无意中做出可能对整个系统造成损害的更改。”

Steve Wade from Mettle says that implementing GitOps is a way of providing an abstraction layer on top of Kubernetes, “For us the platform needs to enable a self-service mechanism for developers. Essentially the platform is there for them to be able to bring business value to Mettle and our customers. We use GitOps as an abstraction layer for developers to onboard new microservices into the platform.”

来自 Mettle 的 Steve Wade 表示，实施 GitOps 是在 Kubernetes 之上提供抽象层的一种方式，“对我们来说，平台需要为开发人员启用自助服务机制。从本质上讲，该平台可以让他们为 Mettle 和我们的客户带来商业价值。我们使用 GitOps 作为抽象层，让开发人员将新的微服务载入平台。”

## Balance between control and flexibility 

## 控制和灵活性之间的平衡

When implementing a developer platform, there is tension between the need for control while at the same time providing some flexibility in tool choice and in certain cloud native patterns. Maintaining the balance between those two elements can be tricky but all participants agreed on the need for implementing constraints and guardrails.

在实现开发人员平台时，控制需求与同时在工具选择和某些云原生模式中提供一些灵活性之间存在紧张关系。保持这两个要素之间的平衡可能很棘手，但所有参与者都同意实施约束和护栏的必要性。

Throughout the discussion, the panelists boiled down their advice into these five practical tips for championing GitOps in your own organization:

在整个讨论过程中，小组成员将他们的建议归结为以下五个实用技巧，以在您自己的组织中支持 GitOps：

### \#1 Define, collaborate and document common cloud native patterns

### \#1 定义、协作和记录常见的云原生模式

At Fidelity the platform team is very transparent in terms of what security they have in place and why it needs to be there. To enforce those security requirements and to be more flexible, they need to support multiple cloud native patterns. For example, they document and provide support for numerous strategies on how to manage secrets and how you might go about using persistent data, among others.

在 Fidelity，平台团队在他们拥有什么安全性以及为什么需要安全性方面非常透明。为了强制执行这些安全要求并更加灵活，他们需要支持多种云原生模式。例如，它们记录并支持有关如何管理机密以及如何使用持久数据等的众多策略。

The Mettle platform team developed and documented a number of microservices patterns together with the development team as a way to illustrate the problems with doing things the old way vs the new way. To keep the information flowing, the platform team at Mettle created a wiki to document all of these patterns.

Mettle 平台团队与开发团队一起开发并记录了许多微服务模式，以此来说明以旧方式与新方式做事的问题。为了保持信息畅通，Mettle 的平台团队创建了一个 wiki 来记录所有这些模式。

### \#2 Take small steps & iterate the process

### \#2 采取小步骤并迭代过程

One of the best ways to start according to Steve is with a component that you're familiar with and that has a small blast radius. In this way, if that one thing doesn't work out, then it's not going to be too difficult to regroup.

根据 Steve 的说法，最好的开始方式之一是使用您熟悉且爆炸半径较小的组件。这样一来，如果那一件事没有解决，那么重组也不会太难。

For [Mettle's GitOps journey](https://www.weave.works/blog/case-study-mettle-leverages-gitops-for-self-service-developer-platform), they began with the platform workload, leaving the developer workloads, and focused on how to update the platform itself. They iterated on a number of different approaches for different aspects: how to deploy the ingress controller; or how to deploy the Prometheus monitoring stack.

对于 [Mettle 的 GitOps 之旅](https://www.weave.works/blog/case-study-mettle-leverages-gitops-for-self-service-developer-platform)，他们从平台工作量开始，离开了开发者工作负载，并专注于如何更新平台本身。他们针对不同方面迭代了许多不同的方法：如何部署入口控制器；或如何部署 Prometheus 监控堆栈。

> “As an operator and as a platform architect, I started small. That means rolling out to staging and Dev environments first, for example, when you're adding GitOps tools like Flux or Flagger to your Kubernetes environments, enable them on staging clusters first to get a feel of what you like and what you don't like about it. This allows you to decide which features tie in better with your environment and how you’re going to integrate them with your workflows.” -- **Javeria Khan, Senior SRE, Palo Alto**

> “作为运营商和平台架构师，我从小处开始。这意味着首先部署到暂存和开发环境，例如，当您将 Flux 或 Flagger 等 GitOps 工具添加到 Kubernetes 环境时，首先在暂存集群上启用它们以了解您喜欢什么和不喜欢什么喜欢它。这使您可以决定哪些功能更适合您的环境，以及如何将它们与您的工作流程集成。” -- **Javeria Khan，高级 SRE，帕洛阿尔托**

### \#3 Develop a good UX from your local machine on to staging

### \#3 从本地机器开发一个好的用户体验到登台

The other important part is to develop some sort of sandbox environment on developer machines so they can experiment with processes on their own. There are many different types of tools they can use such as [Kind](https://kind.sigs.k8s.io/), [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) and maybe even [Ignite](https://github.com/weaveworks/ignite). In addition to those tools, you can also take advantage of public Helm charts and public images for experimentation. Steve suggests to deploy an NGINX ingress controller in Minikube using [Flux](https://www.weave.works/oss/flux/). And once you’ve built the path, developers will understand the processes much better.

另一个重要部分是在开发人员机器上开发某种沙箱环境，以便他们可以自己试验流程。他们可以使用许多不同类型的工具，例如 [Kind](https://kind.sigs.k8s.io/)、[Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) 甚至 [Ignite](https://github.com/weaveworks/ignite)。除了这些工具之外，您还可以利用公共 Helm 图表和公共图像进行实验。 Steve 建议使用 [Flux](https://www.weave.works/oss/flux/) 在 Minikube 中部署 NGINX 入口控制器。一旦你建立了路径，开发人员就会更好地理解流程。

### \#4 Host brown bag information sessions

### \#4 主持棕色包信息会议

Javeria Khan, in addition to documenting cloud native patterns, took it one step further and suggests that people often learn in different ways. Javeria suggests organizing and recording brown bag sessions for any patterns and strategies you collaborate on and then saving those recordings so that people can view them in their own time.

除了记录云原生模式之外，Javeria Khan 更进一步，并建议人们经常以不同的方式学习。 Javeria 建议为您合作的任何模式和策略组织和录制棕色包会话，然后保存这些录音，以便人们可以在自己的时间查看它们。

### \#5 Over communicate changes

### \#5 过度沟通变化

All panelists agreed that you can never over-communicate your changes. Experiment with different ways of communicating changes, both written and verbal, and both as formal and informal presentations.

所有小组成员都同意，您永远不能过度传达您的更改。尝试不同的交流方式，包括书面和口头，以及正式和非正式的演示。

View the entire panel for more questions and answers on the length of our panelists journey and other great questions asked from the viewing audience: 

查看整个小组，了解有关小组成员旅程长度的更多问题和答案以及观众提出的其他重要问题：

For more tips on teaching your team GitOps you can refer to the [GitOps Conversation Kit](https://gitops-community.github.io/kit/), join one of our [upcoming webinars](https://www.weave.works/press/events/) or invite the [Weaveworks team for personalized training](https://www.weave.works/services/kubernetes-training/).

有关教授团队 GitOps 的更多技巧，您可以参考 [GitOps 对话工具包](https://gitops-community.github.io/kit/)，加入我们的[即将举行的网络研讨会](https://www.gitops-community.github.io/kit/)。 weave.works/press/events/) 或邀请 [Weaveworks 团队进行个性化培训](https://www.weave.works/services/kubernetes-training/)。

[Contact us for information on how to champion GitOps in your organization.](https://www.weave.works/services/kubernetes-training/)

[有关如何在您的组织中支持 GitOps 的信息，请联系我们。](https://www.weave.works/services/kubernetes-training/)

* * *

* * *

## About Anita Buehrle

## 关于 Anita Buehrle

Anita has over 20 years experience in software development. She’s written technical guides for the X Windows server company, Hummingbird (now OpenText) and also at Algorithmics, Inc. She’s managed product delivery teams, and developed and marketed her own mobile apps. Currently, Anita leads content and other market-driven initiatives at Weaveworks. 

Anita 拥有超过 20 年的软件开发经验。她为 X Windows 服务器公司 Hummingbird（现为 OpenText）和 Algorithmics, Inc. 编写技术指南。她管理产品交付团队，并开发和营销她自己的移动应用程序。目前，Anita 在 Weaveworks 领导内容和其他市场驱动的计划。

