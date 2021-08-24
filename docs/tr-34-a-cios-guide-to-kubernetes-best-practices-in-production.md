# A CIOs Guide to Kubernetes Best Practices in Production

# CIO 生产 Kubernetes 最佳实践指南

In this article, we will dive into Kubernetes best practices for CIOs and CTOs. It is based on our blog series outlining best practices for DevOps and Kubernetes admins and provides a broader more zoomed-out view of best practices in production.

在本文中，我们将深入探讨为首席信息官和首席技术官的 Kubernetes 最佳实践。它基于我们的博客系列，概述了 DevOps 和 Kubernetes 管理员的最佳实践，并提供了生产中最佳实践的更广泛、更缩小的视图。

August 13, 2019 From: https://www.replex.io/blog/a-cios-guide-to-kubernetes-best-practices-in-production

Chief Information Officer (CIO)

We recently wrote a series of articles about [Kubernetes best practices](http://www.replex.io/blog/kubernetes-in-production-readiness-checklist-and-best-practices-for-resource-management) in production . The series outlines Kubernetes best practices from a resource management, disaster recovery, availability, security, scalability, monitoring and governance perspective. It digs into the internals of Kubernetes and is aimed towards DevOps teams in the trenches, getting their hands dirty with Kubernetes on a daily basis.

我们最近写了一系列关于[Kubernetes最佳实践](http://www.replex.io/blog/kubernetes-in-production-readiness-checklist-and-best-practices-for-resource-management)在生产中的文章。该系列从资源管理、灾难恢复、可用性、安全性、可扩展性、监控和治理的角度概述了 Kubernetes 最佳实践。它深入研究了 Kubernetes 的内部结构，并针对处于战壕中的 DevOps 团队，让他们每天接触 Kubernetes。

With this article, we intend to provide a more zoomed out view of Kubernetes best practices in production. The article distils out the main learnings from the earlier series and is targeted towards CIOs and CTOs. It will dig into best practices using some of the same attributes of production workloads from the previous series, including monitoring, availability, security and governance.

在本文中，我们打算提供 Kubernetes 生产中最佳实践的缩小视图。这篇文章从早期系列中提炼出主要知识，并针对 CIO 和 CTO。它将使用上一系列中生产工作负载的一些相同属性深入研究最佳实践，包括监控、可用性、安全性和治理。

We will also outline best practices for CIOs and CTOs to help align organisational culture with the new realities of distributed DevOps and SRE teams, the increasing skill overlap between traditional devs and ops and new CI/CD paradigms for software release cycles. So let’s jump right into it.

我们还将概述 CIO 和 CTO 的最佳实践，以帮助组织文化与分布式 DevOps 和 SRE 团队的新现实、传统开发人员和运维人员之间日益增加的技能重叠以及软件发布周期的新 CI/CD 范式保持一致。所以让我们直接进入它。

## Kubernetes Best Practices in Production: Monitoring

## Kubernetes 生产最佳实践：监控

### Monitoring for Cloud-Native Applications

### 云原生应用程序监控

The [cloud-native](https://github.com/cncf/toc/blob/master/DEFINITION.md) set of tools have changed the way software is developed, deployed and managed. This new toolset has necessitated a shift in the way both the tools themselves as well as the applications propped up by them are monitored.

[云原生](https://github.com/cncf/toc/blob/master/DEFINITION.md) 工具集改变了软件的开发、部署和管理方式。这个新的工具集需要改变工具本身以及由它们支持的应用程序的监控方式。

The same is true of Kubernetes, which introduces a number of new abstractions on both the hardware as well as the application layer. Any [monitoring pipeline](http://www.replex.io/blog/kubernetes-in-production-the-ultimate-guide-to-monitoring-resource-metrics) for Kubernetes needs to take both these new abstractions as well as its resource management model into account.

Kubernetes 也是如此，它在硬件和应用层都引入了许多新的抽象。 Kubernetes 的任何 [监控管道](http://www.replex.io/blog/kubernetes-in-production-the-ultimate-guide-to-monitoring-resource-metrics) 都需要采用这些新的抽象以及其资源管理模型考虑在内。

This means that in addition to monitoring historically relevant infrastructure metrics like CPU and RAM utilisation for cloud VMs and physical machines, logical abstractions like pods, services and replica sets also need to be considered.

这意味着除了监控历史相关的基础设施指标（如云虚拟机和物理机的 CPU 和 RAM 利用率）外，还需要考虑逻辑抽象，如 pod、服务和副本集。

### Observability Paradigm for Kubernetes

### Kubernetes 的可观察性范式

More importantly, however, Kubernetes monitoring needs to pivot to a new observability paradigm. Traditionally organisations have relied on black box monitoring methods to monitor infrastructure and applications. Black box monitoring observes only the external behaviour of a system.

然而，更重要的是，Kubernetes 监控需要转向新的可观察性范式。传统上，组织依靠黑盒监控方法来监控基础设施和应用程序。黑盒监控仅观察系统的外部行为。

In the cloud-native age of containers, orchestration, and microservices, monitoring needs to move beyond black box monitoring. Black box monitoring can still serve as the baseline for a monitoring strategy but it needs to be complemented by newer white box monitoring methods more suited to the distributed, ephemeral nature of containers and Kubernetes.

在容器、编排和微服务的云原生时代，监控需要超越黑盒监控。黑盒监控仍然可以作为监控策略的基准，但它需要由更适合容器和 Kubernetes 的分布式、短暂性质的较新的白盒监控方法来补充。

[Observability](https://thenewstack.io/monitoring-and-observability-whats-the-difference-and-why-does-it-matter/) encompasses both traditional black box monitoring methods in addition to newer monitoring paradigms like logging , tracing and metrics (together known as white box monitoring). Observability pipelines decouple data collection from data ingestion by introducing a buffer.

[可观察性](https://thenewstack.io/monitoring-and-observability-whats-the-difference-and-why-does-it-matter/) 除了更新的监控范式（如日志记录）之外，还包括传统的黑盒监控方法、跟踪和指标（统称为白盒监控）。可观察性管道通过引入缓冲区将数据收集与数据摄取分离。

The pipeline serves as the central repository of traces, metrics, logs and events which are then forwarded to the appropriate service using a data router. This mitigates the need to have agents for each destination running on each host and reduces the number of integrations that need to be maintained. It also allows enterprises to avoid vendor lock-in and quickly test new SaaS-based monitoring services.

管道作为跟踪、指标、日志和事件的中央存储库，然后使用数据路由器转发到适当的服务。这减少了在每个主机上运行每个目的地的代理的需要，并减少了需要维护的集成数量。它还允许企业避免供应商锁定并快速测试新的基于 SaaS 的监控服务。

### Observability Best Practices 
### 可观察性最佳实践
Observability aims to understand the internals of a system and how it works to quickly debug and resolve issues in production. Since it integrates logs, traces and metrics into traditional monitoring pipelines it covers much more ground and requires a lot more effort to deploy.

可观察性旨在了解系统的内部结构以及它如何在生产中快速调试和解决问题。由于它将日志、跟踪和指标集成到传统的监控管道中，因此它涵盖了更多的领域并且需要更多的努力来部署。

A best practice, therefore, is for CIOs and CTOs to gradually build towards a full observability pipeline for their cloud-native environments by integrating elements of [white box monitoring](https://landing.google.com/sre/sre-book /chapters/monitoring-distributed-systems/) over time.

因此，最佳实践是让 CIO 和 CTO 通过集成 [白盒监控](https://landing.google.com/sre/sre-book /chapters/monitoring-distributed-systems/) 随着时间的推移。

The adoption of cloud-native technologies has also resulted in much more overlap between traditional dev and ops teams. Observability pipelines allow organisations to better integrate these teams by helping build a culture based on facts and feedback.

云原生技术的采用也导致传统开发和运营团队之间出现更多重叠。可观察性管道通过帮助建立基于事实和反馈的文化，使组织能够更好地整合这些团队。

## Kubernetes Best Practices in Production:

## Kubernetes 生产最佳实践：

## High availability, Backup and Disaster Recovery

## 高可用性、备份和灾难恢复

High availability and disaster recovery are crucial elements of any enterprise application. Orchestration engines like Kubernetes introduce additional layers which have to be considered when designing highly available architectures.

高可用性和灾难恢复是任何企业应用程序的关键要素。像 Kubernetes 这样的编排引擎引入了额外的层，在设计高可用性架构时必须考虑这些层。

### Multi-Layered High Availability

### 多层高可用性

[Highly available Kubernetes](http://www.replex.io/kubernetes-production-readiness-and-best-practices-checklist) environments can be seen in terms of two distinct layers or levels. The bottom-most layer is the infrastructure layer, which can refer to any number of public cloud providers or physical infrastructure in a data centre. Next is the orchestration layer which includes both hardware and software abstractions like nodes, clusters, containers and pods as well as other application components.

[高可用 Kubernetes](http://www.replex.io/kubernetes-production-readiness-and-best-practices-checklist) 环境可以从两个不同的层或级别来看。最底层是基础设施层，可以指数据中心内任意数量的公共云提供商或物理基础设施。接下来是编排层，它包括硬件和软件抽象，如节点、集群、容器和 Pod 以及其他应用程序组件。

### High Availability on the IAAS and On-premises Layer

### IAAS 和本地层的高可用性

Public cloud providers provide a number of high availability mechanisms for compute, storage and networking that should serve as a baseline for any Kubernetes environment. CIOs and CTOs also need to bake in redundancy into compute, storage and networking equipment supporting Kubernetes environments in on-premise data centres.

公共云提供商为计算、存储和网络提供了许多高可用性机制，这些机制应该作为任何 Kubernetes 环境的基准。 CIO 和 CTO 还需要在计算、存储和网络设备中加入冗余，以支持本地数据中心的 Kubernetes 环境。

### High Availability on the Orchestration Layer

### 编排层的高可用性

On the orchestration layer, a [multi-master Kubernetes](http://www.replex.io/blog/kubernetes-in-production-readiness-checklist-and-best-practices) cluster is a good starting point. Master nodes should also be distributed across cloud provider zones to ensure they are not affected by outages in any one zone.

在编排层，[multi-master Kubernetes](http://www.replex.io/blog/kubernetes-in-production-readiness-checklist-and-best-practices)集群是一个很好的起点。主节点还应该分布在云提供商区域之间，以确保它们不受任何一个区域中断的影响。

Availability on the orchestration layer, however, needs to move beyond simple multi-master clusters. A best practice is to provision a minimum of 3 master nodes distributed across multiple zones. Similarly, worker nodes should also be distributed across zones for high availability.

然而，编排层的可用性需要超越简单的多主集群。最佳实践是提供至少 3 个分布在多个区域中的主节点。同样，工作节点也应该跨区域分布以获得高可用性。

In addition to having at least 3 master nodes, a best practice is to replicate the etcd master component and place it on dedicated nodes. It is recommended to have at least [5 etcd members](https://etcd.io/docs/v3.3.12/faq/) for production clusters.

除了至少有 3 个主节点之外，最佳实践是复制 etcd 主组件并将其放置在专用节点上。对于生产集群，建议至少有 [5 个 etcd 成员](https://etcd.io/docs/v3.3.12/faq/)。

On the application layer, CIOs and CTOs need to ensure the use of native Kubernetes controllers like statefulsets or deployments. These will ensure that the desired number of pod replicas are always up and running.

在应用层，CIO 和 CTO 需要确保使用原生 Kubernetes 控制器，如 statefulsets 或部署。这些将确保所需数量的 pod 副本始终启动并运行。

### Backup and Disaster Recovery

### 备份和灾难恢复

Backup and disaster recovery should also figure at the top of every CIOs to-do list for Kubernetes clusters in production. The etcd master component is responsible for storing the cluster state and configuration. Having a plan for regular etcd backups is, therefore, a best practice. Stateful workloads on Kubernetes leverage [persistent volumes](http://www.replex.io/blog/the-ultimate-kubernetes-cost-guide-adding-persistent-storage-to-the-mix) which also need to be backed up.

对于生产中的 Kubernetes 集群，备份和灾难恢复也应该排在每个 CIO 的待办事项列表的首位。 etcd master 组件负责存储集群状态和配置。因此，制定定期 etcd 备份计划是最佳实践。 Kubernetes 上的有状态工作负载利用 [持久卷](http://www.replex.io/blog/the-ultimate-kubernetes-cost-guide-adding-persistent-storage-to-the-mix) 也需要支持起来。

Backup and disaster recovery are important elements of mission-critical enterprise applications. CTOs and CIOs need to have a well thought out and comprehensive high availability, backup and disaster recovery mechanism for Kubernetes, that encompasses all layers.

备份和灾难恢复是关键任务企业应用程序的重要元素。 CTO 和 CIO 需要为 Kubernetes 制定一个深思熟虑且全面的高可用性、备份和灾难恢复机制，包括所有层。

## Kubernetes Best Practices in Production:

## Kubernetes 生产最佳实践：

## Distributed Devops and SRE

## 分布式 DevOps 和 SRE

The future of enterprise software is moving towards containerised [microservices](https://martinfowler.com/articles/microservices.html) based distributed applications deployed on Kubernetes with the cloud as an underlying layer. This new cloud-native landscape needs to be reflected in the way dev and ops teams are organised internally as well as in the software release cycle.

企业软件的未来正朝着基于容器化 [微服务](https://martinfowler.com/articles/microservices.html) 的分布式应用程序发展，这些应用程序部署在 Kubernetes 上，以云为底层。这种新的云原生环境需要反映在开发和运营团队的内部组织方式以及软件发布周期中。

### Cloud-Native Roles and Teams 
### 云原生角色和团队
Kubernetes and cloud-native technologies have changed traditional dev and ops roles, broken down siloed dev and ops teams as well as changing the entire software release lifecycle. Given these paradigm changes, we will outline best practices for CIOs and CTOs in terms of role definitions, team composition and new paradigms for developing and deploying software.

Kubernetes 和云原生技术改变了传统的开发和运营角色，打破了孤立的开发和运营团队，并改变了整个软件发布生命周期。鉴于这些范式变化，我们将在角色定义、团队组成以及开发和部署软件的新范式方面概述 CIO 和 CTO 的最佳实践。

[DevOps](https://landing.google.com/sre/sre-book/chapters/introduction/) has already broken up the siloed development, testing and operations teams that serviced traditional monolithic applications. More and more developer teams are internalising ops skills.

[DevOps](https://landing.google.com/sre/sre-book/chapters/introduction/) 已经打破了为传统单体应用服务的孤立的开发、测试和运营团队。越来越多的开发团队正在内化运维技能。

In the new cloud-native world, however, the boundary between dev and ops has blurred even more. CIOs and CTOs need to ensure that every DevOps team has the required skills and knowledge to automate, monitor and optimize the distributed, cloud-native applications being developed. They should also have the required skills to ensure highly available and scalable applications, implement networking as well as onboard the tools required throughout the application lifecycle.

然而，在新的云原生世界中，开发和运营之间的界限更加模糊。 CIO 和 CTO 需要确保每个 DevOps 团队都具备自动化、监控和优化正在开发的分布式云原生应用程序所需的技能和知识。他们还应该具备确保高度可用和可扩展的应用程序、实施网络以及加载整个应用程序生命周期所需的工具所需的技能。

### DevOps and SRE

### DevOps 和 SRE

One way to inject these skills into already existing DevOps teams is to move towards [SRE](https://landing.google.com/sre/sre-book/toc/index.html). SRE is an implementation of DevOps, developed internally by Google that pushes for an even more overlapping skill set for individual developers. SREs typically divide their time equally between development and ops responsibilities.

将这些技能注入现有 DevOps 团队的一种方法是转向 [SRE](https://landing.google.com/sre/sre-book/toc/index.html)。 SRE 是 DevOps 的一种实现，由 Google 内部开发，旨在为个人开发人员提供更多重叠的技能集。 SRE 通常在开发和运维职责之间平均分配他们的时间。

In the context of Kubernetes, a best practice for CIOs and CTOs is to sprinkle SREs among DevOps teams. These SREs would, in turn, be responsible for both development as well as managing performance, on-boarding tools, building in automation and monitoring.

在 Kubernetes 的背景下，CIO 和 CTO 的最佳实践是在 DevOps 团队中散布 SRE。反过来，这些 SRE 将负责开发和管理性能、入门工具、构建自动化和监控。

### The Role of Central IT

### 中央 IT 的作用

The increasingly distributed nature of enterprise applications translating into distributed DevOps teams, however, does not mean that central IT loses its significance. There does need to be some degree of [control and oversight](http://www.replex.io/cost-controlling) over these teams.

然而，企业应用程序日益分散的性质转化为分布式 DevOps 团队，但这并不意味着中央 IT 失去其重要性。确实需要对这些团队进行某种程度的[控制和监督](http://www.replex.io/cost-controlling)。

Even though organisations increasingly prefer developers with cross-domain knowledge of ops, overlapping skills do tend to dilute both development and ops.

尽管组织越来越喜欢具有跨领域运维知识的开发人员，但重叠的技能确实会稀释开发和运维。

A best practice, therefore, is to have a central IT team that includes personnel with ops and infrastructure skill sets. This skill set will enable central IT to provide DevOps teams with critical services that are shared by those teams. It will also ensure that organisations avoid wasted effort due to distributed teams figuring out solutions to shared problems.

因此，最佳实践是拥有一个中央 IT 团队，其中包括具有操作和基础设施技能集的人员。该技能集将使中央 IT 能够为 DevOps 团队提供这些团队共享的关键服务。它还将确保组织避免因分布式团队找出共享问题的解决方案而浪费精力。

Both the cloud and Kubernetes itself have made it increasingly easier for teams to provision and consume resources. The cloud-native movement and DevOps also emphasize on agility and the ability to self-service resources. This can at times lead to an explosion in the number of compute resources provisioned and can potentially lead to wastage and inefficient resource usage. A strong central IT team will be able to [govern](http://www.replex.io/blog/kubernetes-in-production-best-practices-for-governance-cost-management-and-security-and-access -control) these distributed teams and avoid the fallouts from self-service and ballooning resources. They will also be able to hold teams accountable.

云和 Kubernetes 本身都让团队越来越容易地配置和使用资源。云原生运动和 DevOps 还强调敏捷性和自助服务资源的能力。这有时会导致配置的计算资源数量激增，并可能导致资源浪费和低效使用。强大的中央 IT 团队将能够治理这些分布式团队，并避免自助服务和不断膨胀的资源带来的后果。他们还将能够让团队承担责任。

## Kubernetes Best Practices in Production:

## Kubernetes 生产最佳实践：

## CI/CD:

In the same way that Kubernetes and the wider cloud-native technology toolset made CIOs rethink traditional dev and ops roles, it has also required a new way of thinking about build and release cycles. Containerised, microservices based applications, developed, deployed and managed by distributed teams, are not very suited to traditional one-dimensional build and release pipelines.

就像 Kubernetes 和更广泛的云原生技术工具集让 CIO 重新思考传统的开发和运营角色一样，它还需要一种新的思考构建和发布周期的方式。由分布式团队开发、部署和管理的容器化、基于微服务的应用程序不太适合传统的一维构建和发布管道。

### CI/CD for Distributed Teams

### 分布式团队的 CI/CD

A best practice for CIOs, therefore, is to support distributed teams with a well-tooled and thought-out [CI/CD](https://medium.com/@nirespire/what-is-cicd-concepts-in-continuous -integration-and-deployment-4fe3f6625007) pipeline. A robust CI/CD pipeline is essential to fully realising the benefits of faster release cycles and agility promised by Kubernetes and cloud-native technologies. There are a number of tools that CIOs and CTOs can use to deploy CI/CD pipelines. These include Jenkins, TravisCI, GitLab CI and Spinnaker.

因此，对于 CIO 来说，最佳实践是通过精心设计和深思熟虑的 [CI/CD](https://medium.com/@nirespire/what-is-cicd-concepts-in-continuous -integration-and-deployment-4fe3f6625007) 管道。强大的 CI/CD 管道对于完全实现 Kubernetes 和云原生技术承诺的更快发布周期和敏捷性的好处至关重要。 CIO 和 CTO 可以使用许多工具来部署 CI/CD 管道。其中包括 Jenkins、TravisCI、GitLab CI 和 Spinnaker。

### Continuous Integration (CI) 
### 持续集成 (CI)
CI/CD is a broad concept and touches on aspects of development, testing and operations. When deploying a CI/CD pipeline from scratch a good place to start is with the developer team. [Continuous integration](https://en.wikipedia.org/wiki/Continuous_integration) is a subset of CI/CD that aims to increase the frequency of code merges and automate build and test processes.

CI/CD 是一个广泛的概念，涉及开发、测试和运营的各个方面。从头开始部署 CI/CD 管道时，最好的起点是开发团队。 [持续集成](https://en.wikipedia.org/wiki/Continuous_integration) 是 CI/CD 的一个子集，旨在增加代码合并的频率以及自动化构建和测试过程。

Instead of developing new features in isolation, developers are encouraged to merge code into the main pipeline as frequently as possible. An automated build is created from these code changes which is then run through a suite of automated tests. Getting developer teams to adopt CI best practices will ensure that code changes and new features are always ready to be pushed out to production.

鼓励开发人员尽可能频繁地将代码合并到主管道中，而不是孤立地开发新功能。自动构建是从这些代码更改中创建的，然后通过一套自动化测试运行。让开发人员团队采用 CI 最佳实践将确保代码更改和新功能始终准备好推向生产。

### Continuous Delivery and Deployment (CD)

### 持续交付和部署 (CD)

Once CI practices are firmly in place, CIOs and CTOs can then move on to [continuous delivery](https://en.wikipedia.org/wiki/Continuous_delivery) and deployment. Continuous delivery is an extension of continuous integration where code changes are run through more rigorous tests and ultimately deployed to an environment that closely mirrors the production environment.

一旦 CI 实践牢固到位，CIO 和 CTO 就可以继续进行[持续交付](https://en.wikipedia.org/wiki/Continuous_delivery) 和部署。持续交付是持续集成的扩展，其中代码更改通过更严格的测试运行，并最终部署到与生产环境密切反映的环境中。

With continuous delivery there is often a human element involved making decisions about when and how frequently to push code into production. [Continuous deployment](https://medium.com/@nirespire/what-is-cicd-concepts-in-continuous-integration-and-deployment-4fe3f6625007) automates the entire pipeline by automatically pushing code into production once it passes the automated builds and tests defined in both the integration and delivery phases.

对于持续交付，通常涉及人为因素决定何时以及以何种频率将代码投入生产。 [持续部署](https://medium.com/@nirespire/what-is-cicd-concepts-in-continuous-integration-and-deployment-4fe3f6625007) 通过在代码通过后自动将代码推送到生产中来自动化整个管道在集成和交付阶段定义的自动化构建和测试。

### CI/CD Best Practices

### CI/CD 最佳实践

Agile distributed teams working in isolation can at times lead to an explosion in the number of isolated build pipelines. To avoid this, a best practice for CIOs is to make the CI/CD pipeline the [only way to push code](https://docs.microsoft.com/en-us/azure/architecture/microservices/ci-cd- kubernetes) into production. This will ensure that all code changes are pushed into a unified build pipeline and are subjected to the a consistent set of integration and test suites.

孤立工作的敏捷分布式团队有时会导致孤立的构建管道数量激增。为避免这种情况，CIO 的最佳实践是使 CI/CD 管道成为[推送代码的唯一方法](https://docs.microsoft.com/en-us/azure/architecture/microservices/ci-cd-kubernetes)投入生产。这将确保所有代码更改都被推送到统一的构建管道中，并受到一组一致的集成和测试套件的约束。

Distributed teams also tend to use a number of different tools and frameworks. CIOs need to ensure that the CICD pipeline is flexible enough to accommodate this usage.

分布式团队也倾向于使用许多不同的工具和框架。 CIO 需要确保 CICD 管道足够灵活以适应这种使用。

Another best practice is to encourage a culture of small incremental code changes and frequent merges among developer teams. Smaller changes are easier to integrate and roll back and minimise the fallout if something goes wrong.

另一个最佳实践是鼓励开发团队之间的小增量代码更改和频繁合并的文化。较小的更改更容易集成和回滚，并在出现问题时最大限度地减少影响。

CIOs also need to institute a [build once policy](https://docs.microsoft.com/en-us/azure/architecture/microservices/ci-cd-kubernetes) at the start of the pipeline. This ensures that later phases of the CI/CD pipeline have a consistent build to work with. It also avoids any inconsistencies that can creep in when using multiple build tools.

CIO 还需要在管道开始时制定 [一次构建策略](https://docs.microsoft.com/en-us/azure/architecture/microservices/ci-cd-kubernetes)。这可确保 CI/CD 管道的后期阶段具有一致的构建可供使用。它还避免了在使用多个构建工具时可能出现的任何不一致。

Additionally, CIOs need to strike a balance between the extent of the testing regime they push code changes through and the speed of the pipeline itself. More rigorous testing regimes while minimising the chances of bad code being pushed to production also have a time overhead.

此外，CIO 需要在他们推动代码更改的测试机制的范围和管道本身的速度之间取得平衡。更严格的测试制度同时最大限度地减少将不良代码推向生产的机会也有时间开销。

CI/CD pipelines even though championing decentralisation and agility do still need to be governed by central IT for major feature releases. CIOs and CTOs need to ensure they strike a balance between governance and oversight from central IT and the agility and flexibility of distributed teams. The need to ensure a degree of oversight that while allowing them control does not impact the release velocity of software and teams.

尽管支持去中心化和敏捷性，CI/CD 管道仍然需要由中央 IT 管理以发布主要功能。 CIO 和 CTO 需要确保他们在中央 IT 的治理和监督与分布式团队的敏捷性和灵活性之间取得平衡。需要确保一定程度的监督，在允许他们控制的同时不会影响软件和团队的发布速度。

## Kubernetes Best Practices in Production:

## Kubernetes 生产最佳实践：

## Choosing the Right Kubernetes Distribution:

## 选择正确的 Kubernetes 发行版：

Even though Kubernetes on its own is vastly feature rich, mission-critical enterprise workloads need to be supported by more feature rich variants to provide required service levels.

尽管 Kubernetes 本身的功能非常丰富，但任务关键型企业工作负载需要更多功能丰富的变体来支持，以提供所需的服务级别。

### Managed Kubernetes

### 托管 Kubernetes

There are a number of [managed Kubernetes offerings](http://www.replex.io/blog/the-ultimate-kubernetes-cost-guide-aws-vs-gce-vs-azure-vs-digital-ocean) from public cloud providers that CIOs and CTOs can evaluate. These managed offerings take over some of the heavy lifting involved in managing upgrades, patches and HA. Public cloud provider offerings do, however, restrict Kubernetes environments to a specific vendor and might not fit well with a future hybrid or multi-cloud strategy.

有许多 [托管 Kubernetes 产品](http://www.replex.io/blog/the-ultimate-kubernetes-cost-guide-aws-vs-gce-vs-azure-vs-digital-ocean) CIO 和 CTO 可以评估的公共云提供商。这些托管产品接管了涉及管理升级、补丁和 HA 的一些繁重工作。然而，公共云提供商的产品确实将 Kubernetes 环境限制为特定供应商，并且可能不适合未来的混合或多云战略。

Commercial value added Kubernetes distributions are also available from vendors like Red Hat, Docker, Heptia, Pivotal and Rancher. Below we will outline some of the features CIOs and CTOs need to look for when choosing one.

Red Hat、Docker、Heptia、Pivo​​tal 和 Rancher 等供应商也提供商业增值 Kubernetes 发行版。下面我们将概述 CIO 和 CTO 在选择时需要寻找的一些功能。

### Feature Set for Kubernetes Distributions 
### Kubernetes 发行版的功能集
**High availability and disaster recovery:** CIOs and CTOs need to look for distributions that support high availability out of the box. This would include support for multi-master architectures, highly available etcd components as well as backup and recovery.

**高可用性和灾难恢复：** CIO 和 CTO 需要寻找支持开箱即用的高可用性的发行版。这将包括对多主架构、高度可用的 etcd 组件以及备份和恢复的支持。

**Hybrid and multi-cloud support:** Vendor lock-in is a very real concern for the modern enterprise. To ensure Kubernetes environments are portable, CIOs need to choose distributions that support a wide range of deployment models, from on-premise to hybrid and multi-cloud. Support for creating and managing multiple clusters is another feature that should be evaluated.

**混合和多云支持** 供应商锁定是现代企业非常关心的问题。为确保 Kubernetes 环境具有可移植性，CIO 需要选择支持各种部署模型的发行版，从内部部署到混合云和多云。对创建和管理多个集群的支持是另一个需要评估的特性。

**Management, upgrades and Operational support:** Managed Kubernetes offerings also need to be evaluated based on ease of setup, installation, and cluster creation as well as day 2 operations including upgrades, monitoring and troubleshooting. A baseline requirement should be support for fully automated cluster upgrades with zero downtime. The solution chosen should also allow upgrades to be triggered manually. Monitoring, health checks, cluster and node metrics and alerts and notifications should also be a standard part.

**管理、升级和运营支持：**托管 Kubernetes 产品还需要根据设置、安装和集群创建的难易程度以及第 2 天的操作（包括升级、监控和故障排除）进行评估。基线要求应该是支持零停机时间的全自动集群升级。选择的解决方案还应允许手动触发升级。监控、健康检查、集群和节点指标以及警报和通知也应该是标准部分。

**Identity and access management:** Identity and access management are important both in terms of security as well as governance. CIOs need to ensure that the Kuberntes distribution they choose supports integration with already existing authentication and authorization tools being used internally. RBAC and granular access control are also important feature sets that should be supported.

**身份和访问管理：**身份和访问管理在安全和治理方面都很重要。 CIO 需要确保他们选择的 Kuberntes 发行版支持与内部使用的现有身份验证和授权工具的集成。 RBAC 和细粒度访问控制也是应该支持的重要功能集。

**Networking and Storage:** The Kubernetes networking model is highly configurable and can be implemented using a number of options. The distribution chosen should either have a native software-defined networking solution that covers the wide range of requirements imposed by different applications or infrastructure or support one of the more popular CNI based networking implementations including Flannel, Calico, kube-router or OVN etc. CIOs also need to ensure that the Kubernetes distribution they choose supports at a minimum, either flexvolume or CSI integration with storage providers as well as deployment on multiple cloud providers and on-premise.

**网络和存储：** Kubernetes 网络模型是高度可配置的，可以使用多种选项来实现。选择的发行版应该具有原生软件定义的网络解决方案，涵盖不同应用程序或基础设施强加的广泛要求，或者支持一种更流行的基于 CNI 的网络实现，包括 Flannel、Calico、kube-router 或 OVN 等。还需要确保他们选择的 Kubernetes 发行版至少支持 flexvolume 或 CSI 与存储提供商的集成，以及在多个云提供商和内部部署上的部署。

**Deploy, manage and upgrade applications:** Kubernetes distributions being considered by CIOs also need to support a comprehensive solution for deploying, managing, and upgrading applications. A helm based, application catalog that aggregates both private and public chart repositories should be a minimal requirement. 

**部署、管理和升级应用程序：** CIO 考虑的 Kubernetes 发行版还需要支持用于部署、管理和升级应用程序的综合解决方案。一个基于 helm 的、聚合私有和公共图表存储库的应用程序目录应该是最低要求。

