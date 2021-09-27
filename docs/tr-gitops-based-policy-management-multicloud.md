# GitOps-based Policy Management: How to Scale in a Multi-Node, Multicloud World

# 基于 GitOps 的策略管理：如何在多节点、多云世界中扩展

January 19, 2021

2021 年 1 月 19 日

Companies want to benefit from an increase in engineering velocity and innovation that comes with adopting Kubernetes and cloud native. But getting there requires a solid strategy for fleets that can scale on a variety of environments on multiple clusters that span on-premise, across multiple clouds and maybe even the edge. Organizations need an automated and consistent approach for managing clusters – one that they can use efficiently no matter how many clusters or clouds or other configurations are included in the mix.

公司希望从采用 Kubernetes 和云原生所带来的工程速度和创新的提高中受益。但是，要实现这一目标，需要为车队制定可靠的战略，以便可以在跨本地、跨多个云甚至边缘的多个集群上的各种环境中进行扩展。组织需要一种自动化且一致的方法来管理集群——无论混合中包含多少集群、云或其他配置，他们都可以有效地使用这种方法。

In this post, we'll walk you through the common challenges faced in multi-cluster environments, and discuss how [GitOps and effective policy management](https://go.weave.works/GitOps_on_AWS_Managing_GRC_for_k8s.html) simplifies large-scale Kubernetes deployments anywhere.

在这篇文章中，我们将带您了解多集群环境中面临的常见挑战，并讨论 [GitOps 和有效的策略管理](https://go.weave.works/GitOps_on_AWS_Managing_GRC_for_k8s.html) 如何简化大规模 Kubernetes随处部署。

## Diverse cluster stacks add complexity

## 不同的集群堆栈增加了复杂性

Kubernetes helps you innovate more quickly provided that you can manage the complexity of configuring platforms for it. One of the biggest challenges is maintaining and deploying consistent Kubernetes developer platforms used by multiple teams who need to operate across multiple environments and clouds.

Kubernetes 可以帮助您更快地进行创新，前提是您可以管理为其配置平台的复杂性。最大的挑战之一是维护和部署由需要跨多个环境和云运行的多个团队使用的一致 Kubernetes 开发人员平台。

Implementing Kubernetes involves much more than simply spinning up a cluster. On top of the base Kubernetes install, there are the core add-ons you need to run it, both within its infrastructure, including a way to monitor its health as well as the tools and applications your development teams require for CD pipelines, code tracing and logging. In addition to this, you may also need to consider any tools for specialized business requirements like machine learning or edge computing whose applications also need to be configured to work with Kubernetes.

实施 Kubernetes 不仅仅是简单地启动一个集群。在基本 Kubernetes 安装之上，还有运行它所需的核心附加组件，包括在其基础架构中，包括监控其运行状况的方法以及开发团队用于 CD 管道、代码跟踪所需的工具和应用程序和日志。除此之外，您可能还需要考虑满足特殊业务需求的任何工具，例如机器学习或边缘计算，其应用程序也需要配置为与 Kubernetes 一起使用。

Wherever your clusters run, fleets require an identical configuration that works. In addition to this, the configuration must also be quickly and easily upgraded with security patches and other required maintenance with minimal to no down time.

无论您的集群在哪里运行，车队都需要一个相同的配置才能工作。除此之外，还必须使用安全补丁和其他必需的维护快速轻松地升级配置，并且停机时间最少甚至没有。

## Manage cluster configuration definitions with GitOps

## 使用 GitOps 管理集群配置定义

Almost every element of Kubernetes uses declarative configuration: the cluster, and each component within it as well as the applications running on it. This allows for the platform configuration to be stored together in Git. With GitOps, reconcilers and software agents running inside the cluster ensures that the cluster and the apps that run on it are always up to date. In the case of a drift in configuration, an alert is triggered and the cluster can be automatically reconciled with a known state stored in Git.

Kubernetes 的几乎每个元素都使用声明式配置：集群、其中的每个组件以及在其上运行的应用程序。这允许将平台配置一起存储在 Git 中。借助 GitOps，在集群内运行的协调器和软件代理可确保集群和在其上运行的应用程序始终处于最新状态。在配置发生偏移的情况下，会触发警报，并且集群可以自动与存储在 Git 中的已知状态进行协调。

![Principles_of_GitOps_2021.png](https://images.contentstack.io/v3/assets/blt300387d93dabf50e/blt2be043a8abafc0e7/6005f1ed3e567f1011da0a13/Principles_of_GitOps_2021.png)

This approach comes with the benefits of:

这种方法具有以下优点：

- Git-backed security guarantees provenance that includes a built-in audit trail of who did what.
- Increased reliability with secure built-in GitOps to automate the upgrade process across fleets.
- Scalable cluster management across fleets of clusters and applications.
- Save time and costs with complete platform configurations pushed to Git.
- Proactive monitoring with alerts on cluster configuration drift.

- Git 支持的安全保证出处，包括内置的审计跟踪谁做了什么。
- 通过安全的内置 GitOps 提高可靠性，以自动化跨车队的升级过程。
- 跨集群和应用程序的可扩展集群管理。
- 将完整的平台配置推送到 Git，从而节省时间和成本。
- 主动监控集群配置漂移警报。

See, [Guide to GitOps](https://www.weave.works/technologies/gitops/) for more information on its benefits.

请参阅 [GitOps 指南](https://www.weave.works/technologies/gitops/)，了解有关其优势的更多信息。

## What does Git-based policy look like?

## 基于 Git 的策略是什么样的？

At the core of reproducible, correct cluster configuration is the ability to manage policy with GitOps. This is a standard component in the [Weave Kubernetes Platform (WKP)](https://www.weave.works/product/enterprise-kubernetes-platform/) that can help meet business and regulatory compliance requirements more efficiently. Policies and rules can be set up by Platform teams to determine roles and permissions on who can commit changes to the base Kubernetes configuration.

可重现、正确的集群配置的核心是能够使用 GitOps 管理策略。这是 [Weave Kubernetes Platform (WKP)](https://www.weave.works/product/enterprise-kubernetes-platform/) 中的一个标准组件，可以帮助更有效地满足业务和法规遵从性要求。平台团队可以设置策略和规则，以确定谁可以提交对基本 Kubernetes 配置的更改的角色和权限。

![GitOps_Policy_Manager.png](https://images.contentstack.io/v3/assets/blt300387d93dabf50e/bltde36e4e0e2e513dd/6005f2116215cf0f9a18aed5/GitOps_Policy_Manager.png)

Role-based access control (RBAC) permissions can be checked in and confirmed in Git at commit time. Before any changes are applied feedback is provided. 

可以在提交时在 Git 中签入和确认基于角色的访问控制 (RBAC) 权限。在应用任何更改之前提供反馈。

The GitOps Policy Manager implements a set of Git-based rules built on top of the [Open Policy Agent (OPA)](https://www.openpolicyagent.org/) standard that is managed by pull request. This ensures that cluster changes are only initiated by the roles that are permitted to do so. Additional security including Role-based Access Control (RBAC) can also be applied for both teams and namespaces through the [WKP workspaces feature](https://www.weave.works/blog/wkp-team-workspaces-rbac).

GitOps 策略管理器实现了一组基于 Git 的规则，这些规则建立在由拉取请求管理的 [开放策略代理 (OPA)](https://www.openpolicyagent.org/) 标准之上。这确保集群更改仅由允许这样做的角色启动。还可以通过 [WKP 工作区功能](https://www.weave.works/blog/wkp-team-workspaces-rbac) 为团队和命名空间应用额外的安全性，包括基于角色的访问控制 (RBAC)。

See [Security with GitOps](https://www.weave.works/use-cases/security-with-gitops/) for more information.

有关更多信息，请参阅 [GitOps 的安全性](https://www.weave.works/use-cases/security-with-gitops/)。

## Self-service Kubernetes with guardrails

## 带有护栏的自助 Kubernetes

With policies and GitOps, you don’t have to choose between configuration consistency and tool choice. Instead, you can let your developers use tools they need to build their delivery pipelines. When all of your configuration information is stored in Git, entire cluster stacks can be pushed out in a consistent and scalable way, no matter how many clusters you are maintaining and deploying.

使用策略和 GitOps，您不必在配置一致性和工具选择之间做出选择。相反，您可以让您的开发人员使用他们构建交付管道所需的工具。当您的所有配置信息都存储在 Git 中时，无论您维护和部署多少个集群，都可以以一致且可扩展的方式推出整个集群堆栈。

[Read how Mettle, a financial services company](https://www.weave.works/blog/case-study-mettle-leverages-gitops-for-self-service-developer-platform) used GitOps and Weaveworks solutions to implement a self-service Kubernetes.

[阅读金融服务公司 Mettle 如何使用 GitOps 和 Weaveworks 解决方案实施自助服务 Kubernetes。

## Managing fleets of clusters

## 管理集群队列

Larger organizations need to manage fleets or multiple clusters on premise, on clouds and at the edge. Another obvious benefit of GitOps policy management is that it allows you to apply the same configurations to as many clusters as you want in an automated fashion.

较大的组织需要在本地、云端和边缘管理队列或多个集群。 GitOps 策略管理的另一个明显好处是，它允许您以自动化方式将相同的配置应用于任意数量的集群。

With GitOps, there is no need to generate or apply custom YAML files for each cluster in your infrastructure. You can let tools do that hard and tedious work for you so that your team can focus on crafting the best configurations for your workloads, instead of applying them.

使用 GitOps，无需为基础架构中的每个集群生成或应用自定义 YAML 文件。您可以让工具为您完成那些艰巨而乏味的工作，这样您的团队就可以专注于为您的工作负载制定最佳配置，而不是应用它们。

This approach not only makes managing dozens of clusters much more efficient, but it also increases security by reducing the risk of manual configuration mistakes or oversights.

这种方法不仅使管理数十个集群的效率大大提高，而且还通过降低手动配置错误或疏忽的风险来提高安全性。

### Achieving consistency between local environments and the cloud

### 实现本地环境和云之间的一致性

Infrastructure in most organizations today includes a mix of on-premise resources and multiple public clouds. Managing clusters in a consistent way on heterogeneous infrastructure can be difficult and error prone without the simplicity and automation that GitOps and git-backed policy management provides.

当今大多数组织中的基础架构包括内部部署资源和多个公共云的组合。如果没有 GitOps 和 git 支持的策略管理提供的简单性和自动化，在异构基础架构上以一致的方式管理集群可能会很困难且容易出错。

Managing configuration with GitOps avoids the problem of having to maintain or secure bespoke and snowflake clusters. You can write a single set of configuration files, tailoring them where necessary to address the nuances of different clouds or on-premise clusters or for specific tools, and then deploy them to all parts of your infrastructure using a common workflow and a secure process.

使用 GitOps 管理配置避免了必须维护或保护定制和雪花集群的问题。您可以编写一组配置文件，在必要时对其进行定制，以解决不同云或本地集群或特定工具的细微差别，然后使用通用工作流和安全流程将它们部署到基础架构的所有部分。

## Streamlining access control across the organization

## 简化整个组织的访问控制

When manually managing fleets of clusters spread over multiple clouds, access control and user roles can quickly get messy. You would typically end up with a different set of roles and access-control policies for each cluster, which is a recipe for oversights and security holes.

当手动管理分布在多个云中的集群队列时，访问控制和用户角色很快就会变得混乱。您通常会为每个集群获得一组不同的角色和访问控制策略，这是造成疏忽和安全漏洞的秘诀。

With [GitOps-based management and team workspaces](https://www.weave.works/blog/wkp-team-workspaces-rbac) that are an integral part of the Weave Kubernetes Platform, however, it becomes feasible to define a single set of access-control policies, keep them in Git and apply them to all clusters or to specific engineering teams. This is true not only because you can apply the configurations in an automated way, but also because having a centralized, automated approach to policy management allows you to control the number of people who have access to your clusters in the first place.

然而，使用 [基于 GitOps 的管理和团队工作区](https://www.weave.works/blog/wkp-team-workspaces-rbac) 作为 Weave Kubernetes 平台的一个组成部分，定义一个一组访问控制策略，将它们保存在 Git 中并将它们应用到所有集群或特定的工程团队。之所以如此，不仅是因为您可以以自动化方式应用配置，还因为拥有集中、自动化的策略管理方法可以让您首先控制有权访问集群的人数。

And that, of course, is a huge security benefit as well as a manageability benefit.

当然，这是一个巨大的安全优势和可管理性优势。

## Conclusion 

##  结论

GitOps is key to rolling out highly scalable Kubernetes strategies that are reliable, efficient and secure. Manual approaches to configuration may work when you have only one or two clusters to manage; but there is no way to work with fleets of clusters without the help of GitOps policy management. You’ll waste far too much time, your configurations will be inconsistent and you will probably make mistakes that lead to security issues. GitOps addresses all of these challenges to enable a Kubernetes management strategy that is truly consistent and scalable.

GitOps 是推出可靠、高效和安全的高度可扩展的 Kubernetes 策略的关键。当您只有一两个集群需要管理时，手动配置方法可能会奏效；但是如果没有 GitOps 策略管理的帮助，就无法使用集群队列。你会浪费太多时间，你的配置会不一致，你可能会犯导致安全问题的错误。 GitOps 解决了所有这些挑战，以实现真正一致且可扩展的 Kubernetes 管理策略。

## Learn more about the GitOps policy manager in Weave Kubernetes Platform

## 了解有关 Weave Kubernetes 平台中 GitOps 策略管理器的更多信息

The [Weave Kubernetes Platform (WKP)](https://www.weave.works/product/enterprise-kubernetes-platform/) is a production ready platform with GitOps as the underlying architecture and developer experience. It simplifies cluster configuration and management across your organization by bringing together all the tools, services, and components that your team needs to run into a single platform. WKP also provides policy and Git-based rules to specify, audit, and control who can change what in the cluster configuration.

[Weave Kubernetes Platform (WKP)](https://www.weave.works/product/enterprise-kubernetes-platform/) 是一个以 GitOps 作为底层架构和开发者体验的生产就绪平台。它通过将团队运行所需的所有工具、服务和组件整合到一个平台中，简化了整个组织的集群配置和管理。 WKP 还提供策略和基于 Git 的规则来指定、审计和控制谁可以更改集群配置中的内容。

[Learn more about the features of Weave Kubernetes Platform in this on-demand webinar](https://vimeo.com/489580000)

[在此点播网络研讨会中了解有关 Weave Kubernetes 平台功能的更多信息](https://vimeo.com/489580000)

* * *

* * *

## About Anita Buehrle

## 关于 Anita Buehrle

Anita has over 20 years experience in software development. She’s written technical guides for the X Windows server company, Hummingbird (now OpenText) and also at Algorithmics, Inc. She’s managed product delivery teams, and developed and marketed her own mobile apps. Currently, Anita leads content and other market-driven initiatives at Weaveworks. 

Anita 拥有超过 20 年的软件开发经验。她为 X Windows 服务器公司 Hummingbird（现为 OpenText）以及 Algorithmics, Inc. 编写技术指南。她管理产品交付团队，并开发和营销自己的移动应用程序。目前，Anita 在 Weaveworks 领导内容和其他市场驱动的计划。

