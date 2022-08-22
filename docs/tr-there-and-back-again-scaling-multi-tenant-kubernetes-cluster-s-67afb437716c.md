# There and Back Again — Scaling Multi-Tenant Kubernetes Cluster(s)
May 12, 2020

# 来来回回——扩展多租户 Kubernetes 集群
2020 年 5 月 12 日

From: https://medium.com/usa-today-network/there-and-back-again-scaling-multi-tenant-kubernetes-cluster-s-67afb437716c

_Everyone loves a good war story_.

_每个人都喜欢精彩的战争故事_。

They say there are lessons to be learned in IT “war stories”. But maybe the real lessons are what happened afterwards, not the event or even what led up to the event? Just as the event is not the whole story, nor is any particular tool the whole story.

他们说，在 IT“战争故事”中可以吸取教训。但也许真正的教训是之后发生的事情，而不是事件，甚至是导致事件的原因？正如事件不是故事的全部，任何特定工具也不是故事的全部。

At Gannett, we’ve got some war stories. Plenty, in fact. When you are working for the largest local news organization, you have to rapidly adapt to changing landscape or be left behind.

在甘尼特，我们有一些战争故事。很多，事实上。当您为最大的本地新闻机构工作时，您必须迅速适应不断变化的环境，否则就会被甩在后面。

![](https://miro.medium.com/max/1400/1*IlAKo4iJnvxqzotDZZGGOw.png)

## From Here — A Kubernetes Cluster for Everyone

## From Here — 适合所有人的 Kubernetes 集群

Kubernetes is one such tool that we use to bridge that gap between where we’ve come from and where we need to go. Our first success with Kubernetes was in November 2016, when our home-grown Kubernetes clusters carried USA Today’s election coverage. It was such a success that we quickly started building out as many Kubernetes clusters as development teams wanted and were willing to manage using [Chef](https://www.chef.io/) and [Scalr](https://www.scalr.com/).

Kubernetes 就是这样一种工具，我们用它来弥合我们来自哪里和我们需要去哪里之间的鸿沟。我们在 Kubernetes 上的第一次成功是在 2016 年 11 月，当时我们的本土 Kubernetes 集群进行了《今日美国》的选举报道。非常成功，我们迅速开始构建尽可能多的 Kubernetes 集群，开发团队想要并愿意使用 [Chef](https://www.chef.io/) 和 [Scalr](https://www.scalr.com/)。

Listening to this talk, you’ll quickly realize how complicated and difficult it is to run Kubernetes the “hard way”. We built up an amazing infrastructure to automate the deployment and management of the 20+ clusters using Chef and Scalr. However, it was still hard, especially on the development teams who wanted to deploy their applications quickly without a lot of hassle. It was still a big step forward, but not far enough.

听完这个演讲，你会很快意识到以“硬方式”运行 Kubernetes 是多么的复杂和困难。我们建立了一个惊人的基础设施，使用 Chef 和 Scalar 自动部署和管理 20 多个集群。然而，这仍然很困难，尤其是对于希望快速部署应用程序而没有太多麻烦的开发团队。这仍然是向前迈出的一大步，但还远远不够。

[**The USA TODAY NETWORK’s SRE team’s journey into Docker and Kubernetes. Was it worth it?** ](http://medium.com/usa-today-network/the-usa-today-networks-sre-team-s-journey-into-docker-and-kubernetes-was-it-worth-it-d20840757f05)

[**今日美国网络的 SRE 团队进入 Docker 和 Kubernetes 的旅程。值得吗？** ](http://medium.com/usa-today-network/the-usa-today-networks-sre-team-s-journey-into-docker-and-kubernetes-was-it-值得-d20840757f05)

[**DevOps — You Build It, You Run It**](http://medium.com/usa-today-network/devops-you-build-it-you-run-it-8f972343eb8e)

[**DevOps — 你构建它，你运行它**](http://medium.com/usa-today-network/devops-you-build-it-you-run-it-8f972343eb8e)

## To There — A Shared, Managed Kubernetes Cluster for Everyone

## 到那里——一个共享的、托管的 Kubernetes 集群供所有人使用

_“Provide a resilient, optimized, feature rich, and easy to use platform that increases speed of innovation while reducing developer toil.”_

_“提供弹性、优化、功能丰富且易于使用的平台，可提高创新速度，同时减少开发人员的工作量。”_

I pick up the story again in early 2018. It’s becoming clear that asking development teams to run their own Kubernetes clusters does not, in fact, “ _reduce developer toil_”. A new approach is needed. [Google’s Kubernetes Engine](https://cloud.google.com/kubernetes-engine) offering is gathering speed and mindshare. At the same time, Gannett is migrating a large portion of our cloud infrastructure from AWS to GCP.

我在 2018 年初再次提起这个故事。很明显，要求开发团队运行自己的 Kubernetes 集群实际上并不能“减少开发人员的工作量”。需要一种新的方法。 [Google 的 Kubernetes 引擎](https://cloud.google.com/kubernetes-engine) 产品正在加速速度和注意力。与此同时，Gannett 正在将我们的大部分云基础设施从 AWS 迁移到 GCP。

A simple, elegant solution appears before us — form a new team to manage shared GKE clusters for everyone! A hybrid model where an operations team, a managed, secure service and RBAC come together to provide Developers with all the access they need to run the applications how they want.

一个简单、优雅的解决方案出现在我们面前——组建一个新团队来为每个人管理共享的 GKE 集群！一种混合模式，运营团队、托管的安全服务和 RBAC 结合在一起，为开发人员提供所需的所有访问权限，以按照他们的意愿运行应用程序。

![Division of responsibilities between teams.](https://miro.medium.com/max/1400/0*F4RTWt6_G9TaN5Cv)

The solution seems straightforward enough. We can follow the best practice documents around multi-tenancy. We create a namespace per team. We create an admin service account in each namespace and give those credentials to the various development teams. We implement Kubernetes integration with Hashicorp Vault. We reduce developer toil by taking back most of the work around maintaining a Kubernetes cluster, including collecting logs and basic metrics in a centralized location. We run two production clusters in one GCP project and the pre-production another project.

解决方案似乎很简单。我们可以遵循关于多租户的最佳实践文档。我们为每个团队创建一个命名空间。我们在每个命名空间中创建一个管理员服务帐户，并将这些凭据提供给各个开发团队。我们实现 Kubernetes 与 Hashicorp Vault 的集成。我们通过收回维护 Kubernetes 集群的大部分工作来减少开发人员的工作量，包括在一个集中位置收集日志和基本指标。我们在一个 GCP 项目中运行两个生产集群，在另一个项目中运行预生产。

![](https://miro.medium.com/max/960/1*5_J9HtUvj2I7lC0rKxbNgw.png)

We start to hit a few issues here and there. We purposely did not apply limit on namespaces. We can’t predict what teams will be needing to run their production applications and why throttle access to readily available resources? A deployment which goes bad and steals resources wasn’t entirely unexpected and could easily be dealt with. We naively thought that by splitting teams across workloads and nodepools would be adequate.

我们开始在这里和那里遇到一些问题。我们故意不对命名空间施加限制。我们无法预测需要哪些团队来运行他们的生产应用程序，以及为什么要限制对现成资源的访问？出现故障并窃取资源的部署并非完全出乎意料，并且可以轻松处理。我们天真地认为，将团队分散到工作负载和节点池中就足够了。

We didn’t anticipate what the combined and highly diverse load would do to GKE. We begin to hit rare bugs in the OSS Kubernetes kernel. The Jenkins K8S plugin we use triggers goroutines to leak and destabilize the entire cluster. We request Google support restart our Master API server every few days to prevent the cluster from crashing over the course of several weeks. A bug in the OSS Linux kernel gets repeatedly triggered by all of the containers starting and stopping on a single node. We start proactively monitoring nodes and rebooting them every few hours until the bug is remedied.

我们没有预料到组合和高度多样化的负载会对 GKE 造成什么影响。我们开始发现 OSS Kubernetes 内核中的罕见错误。我们使用的 Jenkins K8S 插件触发 goroutine 来泄漏和破坏整个集群。我们要求 Google 支持每隔几天重新启动我们的 Master API 服务器，以防止集群在几周内崩溃。 OSS Linux 内核中的错误会被单个节点上的所有容器启动和停止反复触发。我们开始主动监控节点并每隔几个小时重新启动一次，直到错误得到修复。

Our methodology for maintaining and updating clusters stops scaling well. We originally started with a dedicated helm chart & [Concourse](https://concourse-ci.org/) job per team. With over 40 teams on three clusters, the helm charts and jobs are becoming harder to maintain and prone to simple errors that delete teams’ namespaces and deployments. The clusters were created with [Terraform](https://www.terraform.io/) files. Those templates are now dangerous to re-apply, as the state drift will cause entire clusters to be deleted and recreated. Ad-hoc documentation starts to spring up attempting to document all the special knobs that were turned on which cluster to fix which issue. Building a new cluster is easily a week long process and not likely to replicate everything. We don’t mention the words “disaster recovery” anymore.

我们维护和更新集群的方法停止了很好的扩展。我们最初从每个团队的专用舵图和 [Concourse](https://concourse-ci.org/) 工作开始。由于三个集群中有 40 多个团队，掌舵图和作业变得越来越难以维护，并且容易出现删除团队命名空间和部署的简单错误。集群是使用 [Terraform](https://www.terraform.io/) 文件创建的。现在重新应用这些模板很危险，因为状态漂移会导致整个集群被删除和重新创建。临时文档开始涌现，试图记录打开哪个集群以解决哪个问题的所有特殊旋钮。构建一个新的集群很容易是一个为期一周的过程，而且不可能复制所有内容。我们不再提及“灾难恢复”这两个词。

Cluster upgrades take days and can cause multiple teams’ applications to break. Keeping the clusters up to date is now a break/fix situation only. Migrating to new node pools is near impossible. Each new node requires the Kubernetes service controller to update all of the backends to add the new nodes to every GCP Forwarding Rule in the whole cluster. We have thousands of services resulting in about 30 minutes of delay migrating to a new node and all of the other nodes being updated.

集群升级需要几天时间，并且可能导致多个团队的应用程序中断。使集群保持最新状态现在只是一种中断/修复情况。迁移到新的节点池几乎是不可能的。每个新节点都需要 Kubernetes 服务控制器更新所有后端，以将新节点添加到整个集群中的每个 GCP 转发规则。我们有数以千计的服务，导致迁移到新节点大约需要 30 分钟的延迟，并且所有其他节点都在更新。

Less easy to identify and rectify are the “hard quotas”, which are documented but not visible in the GCP quotas page. We found them while troubleshooting problems through cryptic log messages or side comments from support. Some limits are hard set and in most cases cannot be adjusted. Sometimes we were able to request Google engineering increase these limits, a bit.

不太容易识别和纠正的是“硬配额”，它们记录在案但在 GCP 配额页面中不可见。我们在通过神秘的日志消息或支持人员的旁注解决问题时发现了它们。有些限制是硬设置的，在大多数情况下无法调整。有时我们可以要求 Google 工程部门稍微提高这些限制。

- Internal Forwarding Rules has a default maximum of 50
- Maximum services/node limits

- 内部转发规则的默认最大值为 50
- 最大服务/节点限制

As the months wore on and the outages stacked up, it became all too clear that we had become too multi-tenant for our own good. Kubernetes is built around the idea of scaling pods and nodes, not more and more services with only a handful of pods. It expects workloads to be similar-ish and not having a large degree of churn. We were mixing too many workloads and deployments in one place.

随着几个月的过去和停电的叠加，很明显，我们已经变得过于多租户了，这对我们自己来说是不利的。 Kubernetes 是围绕扩展 pod 和节点的想法构建的，而不是越来越多的服务，只有少数几个 pod。它期望工作负载是相似的，并且不会有很大程度的流失。我们在一个地方混合了太多的工作负载和部署。

## And Back Again — Managed Kubernetes Clusters for Everyone 

## 又回来了——为所有人管理 Kubernetes 集群

In 2019, it was time to go back to where we began. All 40 development teams working on the same clusters wasn’t working super well anymore. The question became how we take the best of our previous two iterations of Kubernetes and bring that forward? How can we create shared clusters that share costs by allowing teams to use the same nodes? How can we isolate some workloads, yet be multi-tenant? How can we increase our resiliency by implementing updated GKE features like regional-master and VPN hub & spoke networking? How can we manage the managed service back like we did in the Scalr days, but with less effort. Mergers are looming, budget cuts and shrinking staff are likely to be in our near term future.

2019年，是时候回到我们开始的地方了。在同一个集群上工作的所有 40 个开发团队不再工作得非常好。问题变成了我们如何充分利用前两次 Kubernetes 迭代并将其向前推进？我们如何通过允许团队使用相同的节点来创建共享成本的共享集群？我们如何隔离一些工作负载，但又是多租户的？我们如何通过实施更新的 GKE 功能（例如区域主机和 VPN 中心辐射网络）来提高弹性？我们如何才能像在 Scalr 时代那样管理托管服务，但工作量更少。合并迫在眉睫，预算削减和裁员可能会出现在我们近期的未来。

We began by looking at alternative methods for managing clusters other than our home grown Concourse jobs. These jobs did what they needed to do but were going to be painful to extend to a dynamically group set of new clusters. The process to build a new cluster had drifted organically over the year and was no longer reproducible. Teams were hard coded with IPs, service accounts credentials and on-the-fly deployment customizations.

我们首先寻找管理集群的替代方法，而不是我们自己种植的 Concourse 工作。这些工作完成了他们需要做的事情，但扩展到一组动态的新集群会很痛苦。一年来，构建新集群的过程有机地发生了变化，不再可重复。使用 IP、服务帐户凭据和动态部署自定义对团队进行硬编码。

_Goal 1) Divide and Conquer_

_目标 1) 分而治之_

We needed to find a middle ground between every team having a dedicated GKE cluster and all teams being on the same cluster. The obvious solutions of splitting by “mission critical” applications on some clusters, but not others wouldn’t work. Every team believes their applications are mission critical and did we really want all of those on a single cluster anyways? We went with the broader categorization of Production, Pre-Production and “Tools”.

我们需要在拥有专用 GKE 集群的每个团队和位于同一集群上的所有团队之间找到一个中间地带。在某些集群上通过“关键任务”应用程序进行拆分的明显解决方案是行不通的。每个团队都认为他们的应用程序是关键任务，我们真的希望所有这些应用程序都在一个集群上吗？我们采用了更广泛的生产、预生产和“工具”分类。

![](https://miro.medium.com/max/1240/1*0DZh2T-mlnjedYkUvMLWlA.png)

The second categorization is more of a psychological one. What are the individual development teams’ tolerance and desire for new features? What level of risks are they willing to accept to always have access to the latest versions and newest features? Some teams are actively waiting for Alpha features to become Beta features. Other teams would much prefer to never hear the word “alpha” or “beta” in regards to their production environments.

第二种分类更多是一种心理分类。各个开发团队对新功能的容忍度和渴望如何？他们愿意接受什么样的风险才能始终访问最新版本和最新功能？一些团队正在积极等待 Alpha 功能成为 Beta 功能。其他团队更愿意在他们的生产环境中永远不要听到“alpha”或“beta”这个词。

![](https://miro.medium.com/max/1280/1*gA9Vf6eqSNQpszBHblpndA.png)

_Goal 2) Automate All The Things_

_目标 2) 自动化所有事情_

We could all agree that more GKE clusters were needed. But how can we do that simply without adding more complexity to our already overscheduled workloads? We did look at other solutions for managing clusters. None of them did everything we needed to do. Some were too expensive. Some trivialized actions to the point where they were no longer reproducible. Other tools would take a huge knowledge lift for us to implement, on top of an already painfully diverse workload.

我们都同意需要更多的 GKE 集群。但是，我们如何才能做到这一点，而又不给我们已经超额安排的工作负载增加更多复杂性呢？我们确实查看了其他管理集群的解决方案。他们都没有做我们需要做的一切。有些太贵了。一些琐碎的动作到了无法重现的地步。在已经令人痛苦的多样化工作负载之上，其他工具将需要我们大量的知识才能实施。

Our ideal solution would cover all of these things:

我们理想的解决方案将涵盖所有这些内容：

- Cheap and allows us to re-use existing tools like Hashicorp Terraform, Concourse and [Hashicorp Vault](https://www.vaultproject.io/)
- Custom clusters, yet all clusters share same base configurations
- Dynamically generate cluster list and credentials
- Modify and update many clusters simultaneously, yet track each cluster’s unique applied yaml via git (GitOps Rules!)

- 便宜且允许我们重复使用现有工具，如 Hashicorp Terraform、Concourse 和 [Hashicorp Vault](https://www.vaultproject.io/)
- 自定义集群，但所有集群共享相同的基本配置
- 动态生成集群列表和凭证
- 同时修改和更新多个集群，但通过 git 跟踪每个集群的唯一应用 yaml（GitOps 规则！）

To briefly summarize, [GitOps upholds the principle that Git is the one and only source of truth. GitOps requires the desired state of the system to be stored in version control such that anyone can view the entire audit trail of changes. All changes to the desired state are fully traceable commits associated with committer information, commit IDs and time stamps. This means that both the application and the infrastructure are now versioned artifacts and can be audited using the gold standards of software development and delivery.](https://www.cloudbees.com/gitops/what-is-gitops) Our new solution must meet these clearly stated goals.

简而言之，[GitOps 坚持 Git 是唯一且唯一的事实来源的原则。 GitOps 要求将所需的系统状态存储在版本控制中，以便任何人都可以查看更改的整个审计跟踪。对所需状态的所有更改都是与提交者信息、提交 ID 和时间戳相关联的完全可跟踪的提交。这意味着应用程序和基础架构现在都是版本化的工件，并且可以使用软件开发和交付的黄金标准进行审计。](https://www.cloudbees.com/gitops/what-is-gitops) 我们的新解决方案必须达到这些明确规定的目标。

We were able to leverage our existing Terraform pipelines to automate building of the GKE clusters and all of the supporting GCP resources.

我们能够利用现有的 Terraform 管道来自动构建 GKE 集群和所有支持的 GCP 资源。

![](https://miro.medium.com/max/1400/1*dUMZyPUYQXy0LAJxJN4JRQ.png)

Next came the deployment yaml’s. Our clusters still need monitoring, cost reporting and logging applications along with Hashicorp Vault integration and several other internal tools. Some clusters will need special configurations or an entirely different set of applications. We can re-use our cluster designations to apply these things.

接下来是部署 yaml 的。我们的集群仍然需要监控、成本报告和日志记录应用程序以及 Hashicorp Vault 集成和其他几个内部工具。一些集群将需要特殊配置或一组完全不同的应用程序。我们可以重新使用我们的集群名称来应用这些东西。

We realized we didn’t need a complicated new product to add to our already very diverse set of toolings. A bash script with some simple looping would solve this problem for us. We could add it to Concourse as another job and no longer worry about which cluster was which type and who was using it when.

我们意识到我们不需要复杂的新产品来添加到我们已经非常多样化的工具集中。带有一些简单循环的 bash 脚本将为我们解决这个问题。我们可以将它作为另一个工作添加到 Concourse 中，而不再担心哪个集群是哪种类型以及谁在何时使用它。

We generate a finalized version of every custer’s unique Kubernetes manifest and save it to git. We like helm charts, but it can be tricky to see the final applied configuration without connecting to each cluster and Tiller. We also wanted a way to document changes naturally. And remember all of those teams who needed custom namespaces and RBAC rules? We built a `setup-namespace` helm chart to create the namespace, RBAC roles and other requirements for every valid combination of clusters, namespaces and gsuite groups.

我们生成每个客户独特的 Kubernetes 清单的最终版本，并将其保存到 git。我们喜欢 helm 图表，但如果不连接到每个集群和 Tiller，就很难看到最终应用的配置。我们还想要一种自然地记录更改的方法。还记得那些需要自定义命名空间和 RBAC 规则的团队吗？我们构建了一个 `setup-namespace` helm 图表来为集群、命名空间和 gsuite 组的每个有效组合创建命名空间、RBAC 角色和其他要求。

![](https://miro.medium.com/max/1400/1*7FyLXqW0nM501eXIQzXPKQ.png)

Function to generate chart manifests

生成图表清单的功能

Applying the configurations was not much more complicated. Combining a gcloud service account and GCP labels on the projects & clusters, we could dynamically identify all of the clusters and how to connect to them. The next step was to simply apply the relevant generated manifests to the matching clusters.

应用配置并不复杂。结合项目和集群上的 gcloud 服务帐户和 GCP 标签，我们可以动态识别所有集群以及如何连接到它们。下一步是简单地将相关生成的清单应用于匹配的集群。

![](https://miro.medium.com/max/1400/1*y3gfa9Pa8yGC_JE4pPZ1tQ.png)

logic for planning, applying changes to clusters

规划逻辑，将更改应用于集群

All of the layers of complexity are simplified to a several, repeated Concourse jobs. An update to the cluster configurations is tested with a \`pr-test\` and must pass a `apply-approved-prs` job, which consists of applying and verifying the new manifest on the Sandbox clusters. If a PR can’t be applied without failing, we want to know about it before it is merged into master. After changes are merged in, a `create-release-tag` job is run and creates a git release tag. Anytime the `release-plan` or `release-apply` jobs are run, they will look for changes between this release tag and what is on the cluster already. Only changed manifests are applied to keep the changes to a minimum.

所有的复杂层都被简化为几个重复的 Concourse 作业。集群配置的更新使用 \`pr-test\` 进行测试，并且必须通过 `apply-approved-prs` 作业，该作业包括在沙盒集群上应用和验证新清单。如果一个 PR 不能在没有失败的情况下应用，我们想在它被合并到 master 之前知道它。合并更改后，将运行“create-release-tag”作业并创建一个 git 发布标签。每当运行 `release-plan` 或 `release-apply` 作业时，它们都会查找此发布标记与集群上已有内容之间的更改。仅应用更改的清单以将更改保持在最低限度。

![View of minimal Concourse jobs required to manage all the clusters.](https://miro.medium.com/max/1400/1*OXxOGcg784krrfP2qgVDgg.png)

_Goal 3) Rebuild With Newest High Availability Features_

_目标 3) 使用最新的高可用性功能重建_

Since early 2018, multiple new GKE and GCP features were released that could dramatically decrease our outages by implementing highly resilient networking and GKE masters. We redesigned our networking between GCP projects to be a hub and spoke model taking advantage of new products like HA VPN Gateways.

自 2018 年初以来，我们发布了多个新的 GKE 和 GCP 功能，这些功能可以通过实施高弹性网络和 GKE 主服务器来显着减少我们的中断。我们重新设计了 GCP 项目之间的网络，使其成为利用 HA VPN 网关等新产品的中心辐射模型。

- Highly resilient external and internal networking
- Private, not public, GKE clusters
- Istio enabled
- `gke-security-group` RBAC and IAM gsuite group enabled
- Regional multiple masters, instead of a single zonal master

- 高弹性的外部和内部网络
- 私有而非公共的 GKE 集群
- 启用 Istio
- `gke-security-group` RBAC 和 IAM gsuite 组已启用
- 区域多个主人，而不是单个区域主人

# The Final Story

# 最后的故事

What has this journey taught us? Running on the latest technology absolutely allows us to adapt quickly and continue to actively reduce our cloud spending. No debate there. However, it comes with a hidden cost of any solution never solving the problem for long. In our four year’s experience with Kubernetes, we’ve never regretted the choice. Our team’s unofficial motto “ _become comfortable with change_” has never been more true.

这段旅程教会了我们什么？在最新技术上运行绝对可以让我们快速适应并继续积极减少我们的云支出。没有辩论。然而，它伴随着任何解决方案的隐藏成本，永远无法解决问题。在我们使用 Kubernetes 的四年经验中，我们从未后悔过这个选择。我们团队的非官方座右铭“_适应变化_”从未如此真实。

[Some rights reserved](http://creativecommons.org/licenses/by/4.0/)

[保留部分权利](http://creativecommons.org/licenses/by/4.0/)

## [More from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----67afb437716c--------------------------------)

Gannett is the largest local news organization in America. Our brands include USA TODAY and 250+ newsrooms spanning 46 states. Gannett’s vastly expanded local-to-national platform reaches over 50% of the U.S. digital audience, including more Millennials than Buzzfeed.

Gannett 是美国最大的地方新闻机构。我们的品牌包括《今日美国》和遍布 46 个州的 250 多家新闻编辑室。 Gannett 大大扩展的本地到国家平台覆盖了超过 50% 的美国数字受众，其中包括比 Buzzfeed 更多的千禧一代。

[Read more from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----67afb437716c--------------------------------)

