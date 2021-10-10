# Kubernetes Backup vs Disaster Recovery

# Kubernetes 备份与灾难恢复

###### April 17, 2020

###### 2020 年 4 月 17 日

When it comes to choosing your Kubernetes backup and restore strategy or designing a disaster recovery solution, there are varying levels of business and technical requirements to consider. To help you better understand these differences, we’ll walk you through what we think Kubernetes backup and restore should contain and how it compares to what Kubernetes disaster recovery should provide. Ultimately, we’ll explain why it is often necessary to have both for proper data protection.

在选择 Kubernetes 备份和恢复策略或设计灾难恢复解决方案时，需要考虑不同级别的业务和技术要求。为了帮助您更好地理解这些差异，我们将引导您了解我们认为 Kubernetes 备份和恢复应包含的内容，以及它与 Kubernetes 灾难恢复应提供的内容的比较。最后，我们将解释为什么通常需要同时拥有两者才能进行适当的数据保护。

Let’s begin with some really useful definitions provided in [this article](https://gomindsight.com/insights/blog/difference-between-disaster-recovery-and-backups/) written by writer Siobhan Climer of Mindsight.

让我们从 [本文](https://gomindsight.com/insights/blog/difference-between-disaster-recovery-and-backups/) 中提供的一些非常有用的定义开始，该定义由 Mindsight 的作家 Siobhan Climer 撰写。

> “ **_Disaster Recovery (DR)_**:  _A strategic security planning model that seeks to protect an enterprise from the effects of natural or human-induced disaster, such as a tornado or cyber attack. A DR plan aims to maintain critical functions before, during, and after a disaster event, thereby causing minimal disruption to business continuity._”
>
> “ **_Backup_**: _The copying of data into a secondary form (i.e. archive file), which can be used to restore the original file in the event of a disaster event._”

> “ **_灾难恢复 (DR)_**：_一种战略安全规划模型，旨在保护企业免受自然或人为灾难（例如龙卷风或网络攻击）的影响。 DR 计划旨在在灾难事件发生之前、期间和之后维护关键功能，从而最大限度地减少对业务连续性的干扰。_”
>
> “ **_Backup_**：_将数据复制到二级形式（即存档文件）中，可用于在发生灾难事件时恢复原始文件。_”

As you can see, these two strategies  are not the same, even though they both provide vital solutions for protecting applications and data. Let’s dive into how these apply to Kuberenetes.

如您所见，这两种策略并不相同，尽管它们都提供了保护应用程序和数据的重要解决方案。让我们深入探讨这些如何应用于 Kuberenetes。

### Backup and Restore

###  备份还原

![](https://www.portworx.com/wp-content/uploads/2020/04/2-1024x232.png)

Backup and recovery for Kubernetes should be focused on the backup of the entire application from the local Kubernetes cluster to a secondary location, which is often offsite. A secondary location could be object storage in public or private clouds or storage available on-prem in different regions or failure domains. Backup solutions can also have multiple backup targets and are often configured to protect from system failure or to complete a compliance or regulatory checklist for an application. When it comes to understanding what a “backup” is for Kubernetes, it means that the backup software is able to understand the semantics of what makes up the application, such as Kubernetes YAML resources—eg, secrets, service accounts, stateful sets, CRDS , app config—and persistent data as one. Only backing up data is not sufficient for a Kubernetes-level backup.

Kubernetes 的备份和恢复应侧重于将整个应用程序从本地 Kubernetes 集群备份到辅助位置，这通常是异地的。次要位置可以是公共或私有云中的对象存储，也可以是不同区域或故障域中的本地可用存储。备份解决方案还可以有多个备份目标，并且通常配置为防止系统故障或完成应用程序的合规性或监管检查表。当谈到理解 Kubernetes 的“备份”是什么时，这意味着备份软件能够理解构成应用程序的语义，例如 Kubernetes YAML 资源——例如，机密、服务帐户、有状态集、CRDS 、应用程序配置和持久数据合二为一。对于 Kubernetes 级别的备份，仅备份数据是不够的。

Backup and restore for Kubernetes should contain the following:

Kubernetes 的备份和恢复应包含以下内容：

- Backup Targets (where to store backed up data)
- Credentials and security (to back data up securely)
- Application and data container-level granularity (cannot just backup vm/server in k8s)
- Ability to have selective backups (backup app only, app + data, entire namespace, single pod, etc.)
- Tooling, such as schedules, jobs, retention, audit and search
- Restore capabilities to and from clusters/namespaces (restore to different namespace or Kubernetes cluster)

- 备份目标（存储备份数据的位置）
- 凭据和安全性（安全备份数据）
- 应用和数据容器级粒度（不能只备份k8s中的vm/server）
- 能够进行选择性备份（仅备份应用程序、应用程序 + 数据、整个命名空间、单个 Pod 等）
- 工具，例如日程安排、工作、保留、审计和搜索
- 从集群/命名空间恢复功能（恢复到不同的命名空间或 Kubernetes 集群）

See [Kubernetes Backup Tools: Comparing Cohesity, Kasten, OpenEBS, Portworx, Rancher Longhorn, and Velero](https://www.portworx.com/kubernetes-backup-tools/) for an overview of Kubernetes backup options.

有关 Kubernetes 备份选项的概述，请参阅 [Kubernetes 备份工具：比较 Cohesity、Kasten、OpenEBS、Portworx、Rancher Longhorn 和 Velero](https://www.portworx.com/kubernetes-backup-tools/)。

## How Kubernetes Backup & Restore differs from Traditional Backup & Restore

## Kubernetes 备份和恢复与传统备份和恢复有何不同

![](https://www.portworx.com/wp-content/uploads/2020/04/3.png)

First, let’s define “traditional.” What we mean by this is a system that is focused on the hypervisor and server layers, not the microservice layer. These systems are in **no way** useless; however, they cannot be dropped into a Kubernetes environment and expected to work seamlessly. These systems typically operate at a higher-level object, such as a virtual machine (VM) or server. While systems that target VM and server-based workloads are valuable, they do not translate well to Kubernetes. Targeting the VM, server, or disk the application is using is not enough anymore—the backup system must understand these differences. Kubernetes-focused solutions still need certain aspects of traditional backup systems, though, such as schedules, jobs, retention, encryption, and tiering.

首先，让我们定义“传统”。我们的意思是一个专注于管理程序和服务器层的系统，而不是微服务层。这些系统**绝不**没用；然而，它们不能被放入 Kubernetes 环境并期望无缝工作。这些系统通常在更高级别的对象上运行，例如虚拟机 (VM) 或服务器。虽然针对 VM 和基于服务器的工作负载的系统很有价值，但它们并不能很好地转换为 Kubernetes。仅针对应用程序正在使用的 VM、服务器或磁盘已经不够了——备份系统必须了解这些差异。不过，以 Kubernetes 为中心的解决方案仍然需要传统备份系统的某些方面，例如计划、作业、保留、加密和分层。

### Disaster Recovery

###  灾难恢复

![](https://www.portworx.com/wp-content/uploads/2020/04/4-1024x250.png)

Disaster recovery (DR) for Kubernetes should focus on the nuances of recovery time objectives and recovery point objectives that translate to how much data loss and downtime a business critical application can endure while keeping critical functions available. DR solutions must also treat the Kubernetes YAML resources and persistent state as a single entity to protect from applications available from a primary site to a DR site. The DR system must also handle the varying levels of RPO and RTO and, depending on the cost and business requirements, protect at these different levels, such as including mechanisms for custom schedules and data replication configurations to meet the varying requirements.

Kubernetes 的灾难恢复 (DR) 应关注恢复时间目标和恢复点目标之间的细微差别，这将转化为业务关键应用程序在保持关键功能可用的同时可以承受多少数据丢失和停机时间。 DR 解决方案还必须将 Kubernetes YAML 资源和持久状态视为单个实体，以保护从主站点到 DR 站点可用的应用程序。 DR 系统还必须处理不同级别的 RPO 和 RTO，并根据成本和业务要求，在这些不同级别进行保护，例如包括自定义计划和数据复制配置的机制以满足不同的要求。

One example of this would be Zero RPO. These are cases where applications cannot incur data loss or downtime. Another example would be low RPO, where requirements are a bit more relaxed and RTOs of hours may be suitable.

一个例子是零 RPO。在这些情况下，应用程序不会导致数据丢失或停机。另一个例子是低 RPO，在这种情况下，要求稍微宽松一些，并且几个小时的 RTO 可能是合适的。

The DR system must also (to some degree) understand how the application is configured to be run in the recovery site. This means it should be able to understand metadata—such as labels or replica numbers—so the recovered application will start properly in the DR site without human intervention. If these types of Kubernetes APIs are not understood, this could lead to disjointed recovery and functionality and ultimately downtime or data loss.

DR 系统还必须（在某种程度上）了解应用程序是如何配置为在恢复站点中运行的。这意味着它应该能够理解元数据（例如标签或副本编号），因此恢复的应用程序将在 DR 站点中正确启动，而无需人工干预。如果不了解这些类型的 Kubernetes API，可能会导致恢复和功能脱节，并最终导致停机或数据丢失。

DR for Kubernetes should contain the following:

Kubernetes 的 DR 应包含以下内容：

- Application and data container-level granularity (understand “what is an application” in k8s)
- Ability to instantiate applications for recovery (understand k8s APIs)
- Ability to cover low RTO and RPO use cases (can be asynchronous)
- Ability to achieve Zero RPO for mission critical apps (synchronous replication)
- Support Active/Active or Active/Standby (be configurable)
- Offer tooling, such as schedules, suspensions, and mutations
- Ability to be selective (app only, app + data, entire ns, etc)

- 应用和数据容器级粒度（理解k8s中的“什么是应用”）
- 能够实例化应用程序以进行恢复（了解 k8s API）
- 能够覆盖低 RTO 和 RPO 用例（可以是异步的）
- 能够实现关键任务应用程序的零 RPO（同步复制）
- 支持 Active/Active 或 Active/Standby（可配置）
- 提供工具，例如时间表、暂停和突变
- 具有选择性（仅限应用程序、应用程序 + 数据、整个 ns 等）

## How Kubernetes Disaster Recovery Differs from Traditional DR

## Kubernetes 灾难恢复与传统灾难恢复有何不同

![](https://www.portworx.com/wp-content/uploads/2020/04/5-1024x495.png)

Disaster recovery in the era of Kubernetes and containerization means that traditional systems cannot target “an application” or “a namespace” in Kubernetes because moving a VM or disk from one site to another means there is no insight into whether that VM or disk has what an application needs. Abstractions in orchestrators today mean that DR solutions must be able to capture and move application metadata that the orchestration system can understand as well as link it to the persistent state the application may be using. Lastly, the DR systems must be able to bring those pieces up in the DR site as a cohesive unit. Systems that protect at the VM and Disk level miss the mark and need the finer details of what makes up an application on Kubernetes. DR for Kubernetes still needs many aspects of traditional solutions, though, such as schedules, witness procedures, and replication policies to attack RPO and RTO requirements from the business perspective.

Kubernetes 和容器化时代的灾难恢复意味着传统系统无法针对 Kubernetes 中的“应用程序”或“命名空间”，因为将 VM 或磁盘从一个站点移动到另一个站点意味着无法洞察该 VM 或磁盘是否具有一个应用程序需要。当今编排器中的抽象意味着 DR 解决方案必须能够捕获和移动编排系统可以理解的应用程序元数据，并将其链接到应用程序可能正在使用的持久状态。最后，DR 系统必须能够将这些部分作为一个有凝聚力的单元整合到 DR 站点中。在 VM 和磁盘级别进行保护的系统没有达到目标，需要了解构成 Kubernetes 应用程序的更详细信息。但是，Kubernetes 的 DR 仍然需要传统解决方案的许多方面，例如计划、见证程序和复制策略，以从业务角度解决 RPO 和 RTO 需求。

### Disaster Recovery vs Backup and Restore

### 灾难恢复与备份和恢复

Kubernetes is in production and is here to stay. Given this, applications on Kubernetes need proper backup and restore as well as disaster recovery when the business requirements call for it. Not all applications are the same, and not all budgets are the same, so it’s imperative that we understand the differences between what can be achieved with backup and DR techniques with Kubernetes. Kubernetes demands a new approach to application and data protection, and there are emerging technologies that can do this, including PX-DR and PX-Backup, but it is important to understand the application and business requirements first and apply the solutions you need.

Kubernetes 已投入生产，并将继续存在。鉴于此，Kubernetes 上的应用程序需要适当的备份和恢复以及在业务需求需要时进行灾难恢复。并非所有应用程序都相同，也并非所有预算都相同，因此我们必须了解使用 Kubernetes 备份和灾难恢复技术可以实现的目标之间的差异。 Kubernetes 需要一种新的应用程序和数据保护方法，并且有新兴技术可以做到这一点，包括 PX-DR 和 PX-Backup，但首先了解应用程序和业务需求并应用您需要的解决方案很重要。

### Related Posts

###  相关文章

- [Zero RPO Disaster Recovery For Kubernetes](https://portworx.com/blog/synchronizing-kubernetes-clusters-disaster-recovery/)

- [Kubernetes 零 RPO 灾难恢复](https://portworx.com/blog/synchronizing-kubernetes-clusters-disaster-recovery/)

We recently discussed different disaster recovery strategies for Kubernetes and covered the different types of failure… 

我们最近讨论了 Kubernetes 的不同灾难恢复策略，并涵盖了不同类型的故障……

