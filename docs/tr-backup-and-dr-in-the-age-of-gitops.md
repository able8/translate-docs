# Backup and DR in the Age of GitOps

# GitOps 时代的备份和灾难恢复

August 10, 2020

2020 年 8 月 10 日

Whether you call it GitOps, infrastructure as code (IaC) or simply using a CI/CD pipeline to automatically deploy changes to an application, we can all agree that moving to a model where your application and resource configuration is defined as code is beneficial to everyone involved—particularly the business it supports. This code is usually saved in git, the most popular source version control system for cloud-native projects; provides an automatic and verifiable change capture process; and simplifies application rollout. More importantly, it also prevents “snowflakes,” wherein the configuration of a deployed application differs from the definition because of manual edits.

无论您将其称为 GitOps、基础设施即代码 (IaC) 还是简单地使用 CI/CD 管道自动将更改部署到应用程序，我们都同意转移到将您的应用程序和资源配置定义为代码的模型有利于每个人都参与其中——尤其是它支持的业务。这段代码通常保存在 git 中，这是云原生项目最流行的源代码版本控制系统；提供自动且可验证的变更捕获过程；并简化应用程序部署。更重要的是，它还可以防止“雪花”，即由于手动编辑，已部署应用程序的配置与定义不同。

But even with these capabilities available to developers and operators, the chances of catastrophic business disruption remain. System safety and scalability should still be an area of concern for developers and operators even if GitOps is on point. As such, the role of backup and disaster recovery capabilities, once the domain of legacy data center operators, needs to be elevated in GitOps’ operating models.

但即使开发人员和运营商可以使用这些功能，灾难性业务中断的可能性仍然存在。即使 GitOps 是正确的，系统安全性和可扩展性仍然应该是开发人员和运营商关注的领域。因此，备份和灾难恢复功能的作用，曾经是传统数据中心运营商的领域，需要在 GitOps 的运营模式中提升。

## It’s All About the Data

## 一切都与数据有关

Any automated way of deploying an application will only bring back Kubernetes objects and configuration. Hence, any persistent data or volumes used by applications are not captured in version control; therefore, bringing back any stateful service such as a relational database or NoSQL system requires that the entire application stack, its data and the dependencies of the stack on the data be discovered, tracked and captured.

任何部署应用程序的自动化方式都只会带回 Kubernetes 对象和配置。因此，应用程序使用的任何持久数据或卷都不会在版本控制中捕获；因此，恢复关系数据库或 NoSQL 系统等任何有状态服务都需要发现、跟踪和捕获整个应用程序堆栈、其数据以及堆栈对数据的依赖关系。

With any application redeployment coming in from an automated system, the backup platform must support intelligent “data-only restores.” This is one of the areas where Kubernetes-native data capabilities are critical; the ability to track the relationship between stateful services (including replication) and then bring back only the volumes or data associated with the application after the application has been redeployed via a CI/CD or GitOps pipeline can be massively helpful for developers.

对于来自自动化系统的任何应用程序重新部署，备份平台必须支持智能的“仅数据恢复”。这是 Kubernetes 原生数据能力至关重要的领域之一；跟踪有状态服务（包括复制）之间的关系，然后在通过 CI/CD 或 GitOps 管道重新部署应用程序后仅带回与应用程序关联的卷或数据的能力对开发人员非常有帮助。

## CI/CD at Scale

## 大规模 CI/CD

Enterprises that run Kubernetes at scale and as shared infrastructure have clusters with hundreds of applications from multiple teams and business units. As different groups make independent technology choices, different CI/CD systems and deployment mechanisms in large companies range from GitOps and CI/CD-based Helm installs to manual deployments. This is the reality in complex environments.

大规模运行 Kubernetes 并作为共享基础架构运行的企业拥有包含来自多个团队和业务部门的数百个应用程序的集群。由于不同的群体做出独立的技术选择，大公司中不同的 CI/CD 系统和部署机制范围从 GitOps 和基于 CI/CD 的 Helm 安装到手动部署。这是复杂环境中的现实。

Faced with this reality, a backup platform is necessary for teams that are still [ramping up their cloud-native development](https://containerjournal.com/topics/container-ecosystems/the-keys-to-successful-infrastructure-configuration/). However, in the fast-moving Kubernetes ecosystem and the 50-plus ways to deploy a containerized application, some applications tend to be long-lived. The systems they were deployed with aren’t being maintained anymore and might not be compatible with the installed Kubernetes version (e.g., applications installed with Helm v2, that will be EOL soon). Being able to restore these important applications (they are long-lived for a reason) quickly in case of accidental or malicious failure will be critical.

面对这一现实，对于仍在[加速他们的云原生开发](https://containerjournal.com/topics/container-ecosystems/the-keys-to-successful-infrastructure-configuration)的团队来说，一个备份平台是必要的/)。然而，在快速发展的 Kubernetes 生态系统和 50 多种部署容器化应用程序的方式中，有些应用程序往往是长期存在的。他们部署的系统不再维护，并且可能与安装的 Kubernetes 版本不兼容（例如，安装了 Helm v2 的应用程序，很快就会停产）。能够在意外或恶意故障的情况下快速恢复这些重要的应用程序（它们因某种原因长期存在)将是至关重要的。

Additionally, when multiple CI/CD systems are in use, it is difficult for an operations group to reach out to hundreds of developer teams to retrigger deployment in case of failure or when applications need to be moved across clusters to perform a Kubernetes version upgrade. 

此外，当使用多个 CI/CD 系统时，如果出现故障或需要跨集群移动应用程序以执行 Kubernetes 版本升级，运维团队很难与数百个开发团队联系以重新触发部署。

Sopra Steria, a European information technology consultancy, recently had this problem when the company had to move 170+ applications to [upgrade OpenShift versions](https://blog.kasten.io/kasten-and-red-hat-migration-and-backup-for-openshift) and couldn't get responses from some teams that were running applications on their cluster. Integrating its backup platform into DevOps workflows enabled Sopra Steria to capture both stateful and stateless applications and move them over without downtime and allowed its developer teams to resync their pipelines to the new clusters when they found free time.

欧洲信息技术咨询公司 Sopra Steria 最近遇到了这个问题，该公司不得不将 170 多个应用程序迁移到 [升级 OpenShift 版本](https://blog.kasten.io/kasten-and-red-hat-migration-and-backup-for-openshift)并且无法从一些在其集群上运行应用程序的团队获得响应。将其备份平台集成到 DevOps 工作流程中，Sopra Steria 能够捕获有状态和无状态的应用程序并在不停机的情况下移动它们，并允许其开发团队在空闲时间将他们的管道重新同步到新集群。

## Disaster Recovery

##  灾难恢复

As teams go through Kubernetes business continuity planning, they need to ensure that Kubernetes and all its applications can be restored quickly in case of disaster. Disaster recovery (DR) can take many forms including cross-cluster, cross-region/data center and cross-cloud.

在团队进行 Kubernetes 业务连续性规划时，他们需要确保 Kubernetes 及其所有应用程序能够在发生灾难时快速恢复。灾难恢复 (DR) 可以采用多种形式，包括跨集群、跨区域/数据中心和跨云。

While pushing configuration to a version control system and using CI/CD will be beneficial here too, the DR problem doesn’t go away. The version control system must be high availability across all fault domains to be able to recover quickly. DR capabilities should be able to support both full backups and granular restores. So, if the version control system (eg git) has been recovered in a separate fault domain, the CI/CD system should be able to deploy the application configuration from the checked-in state while selectively bringing only the data back from its DR copies . If the version control system is not available, the application configuration and all related Kubernetes should be restored from the backup along with the data. However, because of its application-centricity and Kubernetes-native knowledge, this state will be quickly reconciled with your GitOps or CI/CD system when the source is available again.

虽然将配置推送到版本控制系统并使用 CI/CD 在这里也有好处，但 DR 问题不会消失。版本控制系统必须在所有故障域中都具有高可用性，以便能够快速恢复。 DR 功能应该能够支持完整备份和粒度还原。因此，如果版本控制系统（例如 git）已在单独的故障域中恢复，则 CI/CD 系统应该能够从签入状态部署应用程序配置，同时有选择地仅从其 DR 副本中带回数据.如果版本控制系统不可用，则应用程序配置和所有相关 Kubernetes 应与数据一起从备份中恢复。但是，由于其以应用程序为中心和 Kubernetes 原生知识，当源再次可用时，这种状态将很快与您的 GitOps 或 CI/CD 系统协调。

There are several other reasons why a Kubernetes-native backup platform that understands runtime application state is most appropriate for integration with a deployment pipeline. Some application deployments have external side effects (e.g., a DNS update or external load balancer creation) that are dynamically named and are non-deterministic. These changes should be captured by a system that understands runtime state, and this should be brought back in on restore, along with the static resources from the automated deployment system.

了解运行时应用程序状态的 Kubernetes 原生备份平台最适合与部署管道集成，还有其他几个原因。某些应用程序部署具有动态命名且不确定的外部副作用（例如，DNS 更新或外部负载均衡器创建）。这些更改应该由了解运行时状态的系统捕获，并且应该在恢复时与来自自动部署系统的静态资源一起带回。

Finally, if clusters are in a regulated environment (eg, US financial services regulated by the SEC) and your backups need to prove that they captured what was really running versus what the desired state of the world should have been, a platform that can capture the runtime state will be critical. Even if the sync period between configuration changes and deployment can be small on average, the window can drift in case of controller or CI/CD failures, manual additions of untracked resources or malicious changes.

最后，如果集群处于受监管的环境中（例如，受美国证券交易委员会监管的美国金融服务），并且您的备份需要证明它们捕获了真正运行的内容与期望的世界状态相比，一个可以捕获的平台运行时状态将是关键的。即使配置更改和部署之间的平均同步周期可能很小，但在控制器或 CI/CD 故障、手动添加未跟踪资源或恶意更改的情况下，窗口可能会漂移。

Interoperability with GitOps, IaC, and CI/CD systems is a must as organizations increasingly are deploying these systems to improve IT operations and enhance business success. Even with these powerful primitives available, the need for backup and disaster recovery is as important as ever, and deploying a backup solution will be critical for safety and scale. However, it is imperative that such a backup system be truly cloud-native that can integrate into GitOps and CI/CD workflows. Legacy VM-based systems will simply not work in the new cloud-native world we live in today!

与 GitOps、IaC 和 CI/CD 系统的互操作性是必须的，因为越来越多的组织正在部署这些系统以改进 IT 运营并提高业务成功率。即使有这些强大的原语可用，备份和灾难恢复的需求也一如既往地重要，部署备份解决方案对于安全性和规模至关重要。然而，这样的备份系统必须是真正的云原生，可以集成到 GitOps 和 CI/CD 工作流中。传统的基于 VM 的系统在我们今天生活的新的云原生世界中根本无法工作！

- [Click to share on Twitter (Opens in new window)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=twitter " Click to share on Twitter")
- [Click to share on Facebook (Opens in new window)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=facebook " Click to share on Facebook")
- [Click to share on LinkedIn (Opens in new window)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=linkedin " Click to share on LinkedIn")
- [Click to share on Reddit (Opens in new window)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=reddit " Click to share on Reddit")

- [点击在 Twitter 上分享（在新窗口中打开）](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=twitter "单击以在 Twitter 上分享")
- [点击分享到 Facebook (在新窗口中打开)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=facebook "单击以在 Facebook 上分享”)
- [点击在 LinkedIn 上分享 (在新窗口中打开)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=linkedin "单击以在 LinkedIn 上分享")
- [点击在 Reddit 上分享（在新窗口中打开）](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=reddit "单击以在 Reddit 上分享")

### _Related_ 

###  _有关的_

- [← New Training Course Prepares Developers to Create Enterprise Blockchain Applications](https://containerjournal.com/news/news-releases/new-training-course-prepares-developers-to-create-enterprise-blockchain-applications/)
- [Survey: Containers Driving Microservices Transition →](https://containerjournal.com/topics/container-ecosystems/survey-containers-driving-microservices-transition/)

- [← 新培训课程让开发人员准备创建企业区块链应用程序](https://containerjournal.com/news/news-releases/new-training-course-prepares-developers-to-create-enterprise-blockchain-applications/)
- [调查：容器推动微服务转型→](https://containerjournal.com/topics/container-ecosystems/survey-containers-driving-microservices-transition/)

![](https://containerjournal.com/wp-content/uploads/2020/08/kasten_niraj_HR-150x150.jpg)

#### Niraj Tolia

#### 尼拉吉托利亚

Niraj Tolia is the CEO and Co-Founder at Kasten and is interested in all things Kubernetes. He has played multiple roles in the past, including the Senior Director of Engineering for Dell EMC's CloudBoost family of products and the VP of Engineering and Chief Architect at Maginatics (acquired by EMC). Niraj received his PhD, MS, and BS in Computer Engineering from Carnegie Mellon University.

Niraj Tolia 是 Kasten 的首席执行官兼联合创始人，对 Kubernetes 的所有事物都感兴趣。他过去曾担任过多个角色，包括 Dell EMC CloudBoost 系列产品的高级工程总监和 Maginatics（被 EMC 收购）的工程副总裁兼首席架构师。 Niraj 在卡内基梅隆大学获得计算机工程博士、硕士和学士学位。

Niraj Tolia has 1 posts and counting. [See all posts by Niraj Tolia](https://containerjournal.com/author/niraj-tolia/) 

Niraj Tolia 有 1 个帖子，并且还在增加中。 [查看 Niraj Tolia 的所有帖子](https://containerjournal.com/author/niraj-tolia/)

