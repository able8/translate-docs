# Cloud Native Backups, Disaster Recovery and Migrations on Kubernetes

# Kubernetes 上的云原生备份、灾难恢复和迁移

28 Jul 2020 Murat Karslioglu



Murat is an infrastructure architect with experience in storage, distributed systems and enterprise infrastructure development. He is VP of Product at MayaData, as well as one of the maintainers of the CNCF project _OpenEBS_ and author of [Kubernetes - A Complete DevOps Cookbook](https://www.linkedin.com/in/muratkarslioglu/)

Murat 是一位基础架构架构师，在存储、分布式系统和企业基础架构开发方面拥有丰富的经验。他是 MayaData 的产品副总裁，也是 CNCF 项目 _OpenEBS_ 的维护者之一和 _Kubernetes - A Complete DevOps Cookbook_ 的作者。

Data is increasingly being run on Kubernetes, and increasingly in a way that supports cloud native architectures. In [a recent Cloud Native Computing Foundation webinar](https://www.cncf.io/webinars/kubernetes-for-storage-an-overview/), [Kiran Mova](https://www.linkedin.com/in/kiranmova/), the leader of the CNCF OpenEBS project and co-founder of MayaData, explained that loosely coupled architectures enable loosely coupled teams. The benefit being that loosely-coupled teams are unblocked; they’re able to iterate as and when they can, free from the scourge of meetings, long release cycles, change control boards and more.

数据越来越多地在 Kubernetes 上运行，并且越来越多地以支持云原生架构的方式运行。在[最近的云原生计算基金会网络研讨会](https://www.cncf.io/webinars/kubernetes-for-storage-an-overview/)中，[Kiran Mova](https://www.linkedin.com/in/kiranmova/)，CNCF OpenEBS 项目的领导者和 MayaData 的联合创始人，他解释说松散耦合的架构支持松散耦合的团队。好处是松散耦合的团队不受阻碍；他们能够尽可能地进行迭代，而不受会议、漫长的发布周期、变更控制板等的困扰。

One implication of this loosely-coupled, per workload approach — called **Container Attached Storage** — is that backups have to change as well. They need to change because the governance model is changing (small teams are in charge now) and because the fundamental architecture of loosely coupled workloads is different.

这种松散耦合、每个工作负载的方法（称为**容器附加存储**）的一个含义是备份也必须更改。他们需要改变，因为治理模型正在发生变化（现在由小团队负责）并且因为松散耦合工作负载的基本架构不同。

So what changes when data is deployed as an enormous number of loosely coupled workloads, as opposed to as a few centrally managed databases?

那么，当数据部署为大量松散耦合的工作负载而不是几个集中管理的数据库时，会发生什么变化？

## **Challenges to Traditional Backups**

## **传统备份的挑战**

Traditionally backups could be performed bottom-up, for example backing up all the data on a particular volume, Kubernetes node, or virtual machine.

传统上可以自下而上执行备份，例如备份特定卷、Kubernetes 节点或虚拟机上的所有数据。

But what happens when instead of one or two workloads on a server or an ESX host, you have 110 or more pods; and each of these pods runs a workload, which is being moved by Kubernetes to other nodes as needed?

但是，当您拥有 110 个或更多 Pod 而不是服务器或 ESX 主机上的一两个工作负载时会发生什么？每个 pod 都运行一个工作负载，Kubernetes 会根据需要将其移动到其他节点吗？

The answer is that if you can only backup at the cluster level, then your granularity is only that of a cluster. However, the small teams running workloads in your environment don’t manage clusters; they manage workloads. So in the event that they need their workload restored, they have to ask someone else to recover for them an entire Kubernetes cluster. And then they need to sort through this cluster for the particular data and application that they need. This breaks the control and autonomy of these small teams, which slows everyone down. It is a shared dependency and it puts your data at risk.

答案是，如果您只能在集群级别进行备份，那么您的粒度只是集群的粒度。但是，在您的环境中运行工作负载的小团队不管理集群；他们管理工作负载。因此，如果他们需要恢复工作负载，他们必须请其他人为他们恢复整个 Kubernetes 集群。然后他们需要对这个集群进行分类，以获得他们需要的特定数据和应用程序。这打破了这些小团队的控制和自主权，让每个人都慢下来。它是一个共享依赖项，它会使您的数据处于危险之中。

**Sponsor Note**

**赞助商备注**

![sponsor logo](https://cdn.thenewstack.io/media/2020/07/2769a255-mayadata@2x.png)

MayaData delivers data agility. MayaData sponsors two Cloud Native Computing Foundation (CNCF) projects, OpenEBS – the leading open-source container attached storage solution – and Litmus – the leading Kubernetes native chaos engineering project. Well-known users of MayaData products include the CNCF itself, Bloomberg, Comcast, Arista, Orange, Intuit, and others.

MayaData 提供数据敏捷性。 MayaData 赞助了两个云原生计算基金会 (CNCF) 项目，OpenEBS——领先的开源容器附加存储解决方案——和 Litmus——领先的 Kubernetes 原生混沌工程项目。 MayaData 产品的知名用户包括 CNCF 本身、Bloomberg、Comcast、Arista、Orange、Intuit 等。

## **Cloud Native Per Workload Backups, Disaster Recovery and Migration** 

## **每个工作负载的云原生备份、灾难恢复和迁移**

Contrast this to the experience of small teams running their workloads on a Container Attached Storage enabled Kubernetes environment, with data protection included via Kubera from MayaData. In this case, each small team remains in control of their workloads. If they feel that their workload needs to be cloned because they want to develop against a read-only copy of their production data, they can do that (assuming Kubernetes RBAC controls allows it). And if this small team feels that their particular use of PostgreSQL suggests they should have a cross AZ deployment for the primary DB cluster and also streaming or near real-time backups to another cloud, they can do that as well. They can also author a common pattern for production PostgreSQL for their environment and have that storage class become the default for all their deploys; and even share that pattern with others in their organization.

将此与小型团队在支持容器附加存储的 Kubernetes 环境中运行工作负载的体验形成对比，其中数据保护通过来自 MayaData 的 Kubera 提供。在这种情况下，每个小团队都可以控制他们的工作量。如果他们觉得他们的工作负载需要克隆，因为他们想针对生产数据的只读副本进行开发，他们可以这样做（假设 Kubernetes RBAC 控件允许）。如果这个小团队认为他们对 PostgreSQL 的特殊使用表明他们应该为主数据库集群进行跨可用区部署，并且还可以流式传输或近实时备份到另一个云，他们也可以这样做。他们还可以为他们的环境编写生产 PostgreSQL 的通用模式，并使该存储类成为他们所有部署的默认值；甚至与组织中的其他人分享这种模式。

The autonomy of small teams to make decisions and iterate quickly is fundamental to agility. Loosely coupled architectures also require the separation of the state of workloads from infrastructure. Your “per workload backup” should be able to restore or migrate workloads to a different cloud platform, even though that other cloud platform has a different storage platform – otherwise, you are tightly coupled to the underlying platform.

小团队做出决策和快速迭代的自主权是敏捷性的基础。松散耦合的架构还需要将工作负载的状态与基础设施分离。您的“每个工作负载备份”应该能够将工作负载恢复或迁移到不同的云平台，即使其他云平台具有不同的存储平台——否则，您将与底层平台紧密耦合。

Additionally, in many organizations, the platform teams need to be able to see the overall behavior of the platform. In our experience, some of the most important responsibilities of the platform teams include the protection of important data, capacity planning as data usage grows, and the management of one or more underlying cloud or hardware environments. Kubera gives these platform teams the tools they need to manage the resilience, capacity, and even the cost and performance of the overall environment.

此外，在许多组织中，平台团队需要能够看到平台的整体行为。根据我们的经验，平台团队的一些最重要的职责包括保护重要数据、随着数据使用量的增长进行容量规划，以及管理一个或多个底层云或硬件环境。 Kubera 为这些平台团队提供了管理整体环境的弹性、容量甚至成本和性能所需的工具。

## **Learn More**

##  **了解更多**

Kubera from MayaData is available with a free forever tier. All you need to have running is a Kubernetes cluster — Kubera will connect to it and give you preconfigured off cluster logging and alerting for your stateful workloads, reporting and visualization, storage provisioning and more — in addition to the per workload backups, disaster recovery and migration mentioned in this article.

MayaData 的 Kubera 提供永久免费层。您只需要运行一个 Kubernetes 集群——Kubera 将连接到它，并为您的有状态工作负载、报告和可视化、存储配置等提供预配置的集群日志和警报——除了每个工作负载的备份、灾难恢复和本文中提到的迁移。

Kubera now includes integration with Cloudian Hyperstore at no additional costs — and Kubera starts at $49 per user per month. [Try it now — we are looking forward to your feedback](https://account.mayadata.io/signup).

Kubera 现在包括与 Cloudian Hyperstore 的集成，无需额外费用 — Kubera 起价为每位用户每月 49 美元。 [立即尝试 - 我们期待您的反馈](https://account.mayadata.io/signup)。

_We will be discussing and demonstrating the new requirements of per workload Kubernetes backups, as well as data management and protection, backup, recovery and cloud migration, on August 6 with our partner, the leader in enterprise-class object storage, Cloudian._ _ [Register now to learn more](https://go.mayadata.io/data-protection-for-kubernetes-webinar)._

_我们将在 8 月 6 日与我们的合作伙伴、企业级对象存储领域的领导者 Cloudian 讨论和展示每个工作负载 Kubernetes 备份以及数据管理和保护、备份、恢复和云迁移的新要求。_ _ [立即注册以了解更多](https://go.mayadata.io/data-protection-for-kubernetes-webinar)._

Feature image via [Pixabay](https://pixabay.com/photos/squirrel-feeding-nuts-nager-garden-4382005/).

特征图片来自 [Pixabay](https://pixabay.com/photos/squirrel-feeding-nuts-nager-garden-4382005/)。

_At this time, The New Stack does not allow comments directly on this website. We invite all readers who wish to discuss a story to visit us on [Twitter](https://twitter.com/thenewstack) or [Facebook](https://www.facebook.com/thenewstack/). We also welcome your news tips and feedback via email: [feedback@thenewstack.io](mailto:feedback@thenewstack.io)._

_此时，The New Stack 不允许直接在本网站上发表评论。我们邀请所有希望讨论故事的读者在 [Twitter](https://twitter.com/thenewstack) 或 [Facebook](https://www.facebook.com/thenewstack/)上访问我们。我们也欢迎您通过电子邮件提供新闻提示和反馈：[feedback@thenewstack.io](mailto:feedback@thenewstack.io)._

[Contributed](https://thenewstack.io/tag/contributed/)[Sponsored](https://thenewstack.io/tag/sponsored/)

[贡献](https://thenewstack.io/tag/contributed/)[赞助](https://thenewstack.io/tag/sponted/)

![](https://cdn.thenewstack.io/static/img/The-New-Stack-Updates-Logo.svg)



