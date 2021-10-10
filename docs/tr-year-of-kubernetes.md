# What we learned after a year of GitLab.com on Kubernetes

# 在 GitLab.com 上使用 Kubernetes 一年后我们学到了什么

John Jarvis

约翰·贾维斯

Sep 16, 2020·10 min read· [Leave a comment](http://about.gitlab.com#disqus_thread)

2020 年 9 月 16 日·10 分钟阅读· [发表评论](http://about.gitlab.com#disqus_thread)

* * *

* * *

![](http://about.gitlab.com/images/blogimages/a_year_of_k8s/nico-e-AAbjUJsgjvE-unsplash.jpg)

For about a year now, the infrastructure department has been working on migrating all services that run on GitLab.com to Kubernetes. The effort has not been without challenges, not only with moving services to Kubernetes but also managing a hybrid deployment during the transition. We have learned a number of lessons along the way that we will explore in this post.

大约一年来，基础设施部门一直致力于将 GitLab.com 上运行的所有服务迁移到 Kubernetes。这项工作并非没有挑战，不仅是将服务迁移到 Kubernetes，还包括在过渡期间管理混合部署。在此过程中，我们学到了许多经验教训，我们将在本文中探讨这些经验教训。

Since the very beginning of GitLab.com, servers for the website have run in the cloud on virtual machines. These VMs are managed by Chef and installed using our [official Linux package](http://about.gitlab.com/install/#ubuntu).
When an application update is required, [our deployment strategy](https://gitlab.com/gitlab-org/release/docs/-/blob/master/general/deploy/gitlab-com-deployer.md) is to simply upgrade fleets of servers in a coordinated rolling fashion using a CI pipeline.
This method, while slow and a bit [boring](http://about.gitlab.com/handbook/values/#boring-solutions), ensures that GitLab.com is using the same installation methods and configuration as our self-managed customers who use Linux packages.
We use this method because it is especially important that any pain or joy felt by the community when installing or configuring self-managed GitLab is also felt by GitLab.com.
This approach worked well for us for a time but as GitLab.com has grown to hosting over 10 million projects we realized it would no longer serve our needs for scaling and deployments.

从 GitLab.com 一开始，网站的服务器就在虚拟机上运行在云中。这些 VM 由 Chef 管理并使用我们的 [官方 Linux 包](http://about.gitlab.com/install/#ubuntu) 进行安装。
当需要更新应用程序时，[我们的部署策略](https://gitlab.com/gitlab-org/release/docs/-/blob/master/general/deploy/gitlab-com-deployer.md) 就是简单的使用 CI 管道以协调的滚动方式升级服务器队列。
这种方法虽然缓慢而且有点[无聊](http://about.gitlab.com/handbook/values/#boring-solutions)，但可以确保 GitLab.com 使用与我们自行管理的相同的安装方法和配置使用 Linux 软件包的客户。
我们使用这种方法是因为社区在安装或配置自我管理的 GitLab 时感受到的任何痛苦或快乐也让 GitLab.com 感受到尤为重要。
这种方法在一段时间内对我们很有效，但随着 GitLab.com 已经发展到托管超过 1000 万个项目，我们意识到它不再满足我们对扩展和部署的需求。

## Enter Kubernetes and cloud native GitLab

## 进入Kubernetes和云原生GitLab

We created the [GitLab Charts](https://gitlab.com/gitlab-org/charts) project in 2017 to prepare GitLab for deployments in the cloud and enable self-managed users to install GitLab into a Kubernetes cluster. We knew then that running GitLab.com on Kubernetes would benefit the SaaS platform for scaling, deployments, and efficient use of compute resources. At the time though there were still many application features that depended on NFS mounts that delayed our migration off of VMs.

我们在 2017 年创建了 [GitLab Charts](https://gitlab.com/gitlab-org/charts) 项目，以准备 GitLab 在云中的部署，并使自我管理的用户能够将 GitLab 安装到 Kubernetes 集群中。我们当时就知道，在 Kubernetes 上运行 GitLab.com 将有利于 SaaS 平台进行扩展、部署和有效利用计算资源。虽然当时仍有许多应用程序功能依赖于 NFS 挂载，这延迟了我们从 VM 的迁移。

The push for cloud native and Kubernetes gave engineering an opportunity to plan a gradual transition that removes some of the network storage dependencies on the application while continuing to develop new features. Since we started planning the migration in the summer of 2019, most of these limitations have been resolved and the journey to running all of GitLab.com on Kubernetes is now well underway!

对云原生和 Kubernetes 的推动为工程提供了一个计划逐步过渡的机会，该过渡消除了对应用程序的一些网络存储依赖，同时继续开发新功能。自从我们在 2019 年夏天开始计划迁移以来，这些限制中的大部分都已得到解决，在 Kubernetes 上运行所有 GitLab.com 的旅程现在正在顺利进行！

## Running GitLab.com on Kubernetes

## 在 Kubernetes 上运行 GitLab.com

For GitLab.com we use a single regional GKE cluster that services all application traffic. To minimize the complexity of the (already complex) migration we focus on services that don't depend on local storage or NFS. While GitLab.com is running from mostly monolithic Rails codebase, we route traffic depending on workload characteristics to different endpoints which are isolated into their own node pools.

对于 GitLab.com，我们使用单个区域 GKE 集群来为所有应用程序流量提供服务。为了最大限度地降低（已经很复杂）迁移的复杂性，我们专注于不依赖于本地存储或 NFS 的服务。虽然 GitLab.com 主要从单一的 Rails 代码库运行，但我们根据工作负载特性将流量路由到不同的端点，这些端点被隔离到自己的节点池中。

On the frontend these types are divided into web, API, git SSH/HTTPs requests, and Registry.
On the backend we divide our queued jobs into different characteristics depending on [predefined resource boundaries](http://about.gitlab.com/blog/2020/06/24/scaling-our-use-of-sidekiq/) that allow us to set Service-level Objective (SLO) targets for a range of different workloads.

在前端，这些类型分为 Web、API、git SSH/HTTPs 请求和注册表。
在后端，我们根据 [预定义的资源边界](http://about.gitlab.com/blog/2020/06/24/scaling-our-use-of-sidekiq/) 将排队的作业划分为不同的特征，允许我们为一系列不同的工作负载设置服务级别目标 (SLO) 目标。

All of these GitLab.com services are configured with the unmodified GitLab Helm chart, which configures them in sub-charts that can be selectively enabled as we gradually migrate services to the cluster.
While we opted to not include some of our stateful services such as Redis, Postgres, GitLab Pages, and Gitaly, when the migration to Kubernetes is finished it will drastically reduce the number of VMs that we currently manage with Chef.

所有这些 GitLab.com 服务都使用未修改的 GitLab Helm 图表进行配置，该图表将它们配置在子图表中，随着我们逐渐将服务迁移到集群，可以有选择地启用这些子图表。
虽然我们选择不包括我们的一些有状态服务，如 Redis、Postgres、GitLab Pages 和 Gitaly，但当迁移到 Kubernetes 完成后，它将大大减少我们目前使用 Chef 管理的虚拟机数量。

## Transparency and managing the Kubernetes configuration

## 透明度和管理 Kubernetes 配置

All configuration is managed in GitLab itself in three configuration projects using Terraform and Helm.
While we use GitLab to run GitLab wherever possible, we maintain a separate GitLab installation for operations.
This is done to ensure we do not depend on the availability of GitLab.com for deployments and upgrades of GitLab.com. 

所有配置都在 GitLab 中使用 Terraform 和 Helm 在三个配置项目中进行管理。
虽然我们尽可能使用 GitLab 来运行 GitLab，但我们为操作维护了一个单独的 GitLab 安装。
这样做是为了确保我们不依赖 GitLab.com 的可用性来部署和升级 GitLab.com。

Even though our pipelines that execute against the Kubernetes cluster run on this separate GitLab deployment, the code repositories are mirrored and publicly viewable at the following locations:

即使我们针对 Kubernetes 集群执行的管道在这个单独的 GitLab 部署上运行，代码存储库也会在以下位置进行镜像和公开查看：

- [k8s-workloads/gitlab-com](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-com): GitLab.com configuration wrapper for the GitLab Helm chart.
- [k8s-workloads/gitlab-helmfiles](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-helmfiles/): Contains the configuration for services that are not directly related to the GitLab application. This includes configurations for cluster logging and monitoring and integrations like PlantUML.
- [gitlab-com-infrastructure](https://gitlab.com/gitlab-com/gitlab-com-infrastructure): Terraform configuration for the Kubernetes and legacy VM infrastructure. All the resources necessary to run the cluster are configured here, including the cluster, node pools, service accounts, and IP address reservations.

- [k8s-workloads/gitlab-com](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-com)：GitLab.com 的 GitLab Helm 图表配置包装器。
- [k8s-workloads/gitlab-helmfiles](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-helmfiles/)：包含与 GitLab 没有直接关系的服务的配置应用。这包括集群日志记录和监控以及 PlantUML 等集成的配置。
- [gitlab-com-infrastructure](https://gitlab.com/gitlab-com/gitlab-com-infrastructure)：Kubernetes 和传统 VM 基础架构的 Terraform 配置。运行集群所需的所有资源都在此处配置，包括集群、节点池、服务帐户和 IP 地址预留。

[![hpa](http://about.gitlab.com/images/blogimages/a_year_of_k8s/hpa.png)](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-com/-/merge_requests/315#note_390180361)
Whenever a change is proposed, a public [short summary](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-com/-/merge_requests/315#note_390180361) is displayed, with a link to detailed diff that an SRE reviews before applying changes to the cluster.

com/-/merge_requests/315#note_390180361)
每当提出更改时，都会显示公开的[简短摘要](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-com/-/merge_requests/315#note_390180361)，并带有链接到 SRE 在对集群应用更改之前审查的详细差异。

For SREs, we link to a detailed diff on our operations GitLab instance that has limited access.
This allows employees and the community, who do not have access to the operational project which is limited to SREs, to have visibility into proposed config changes.
By having a public GitLab instance for code, and a private instance for [CI pipelines](http://about.gitlab.com/stages-devops-lifecycle/continuous-integration/), we are able to keep a single workflow while at the same time ensuring we don't have a dependency on GitLab.com for configuration updates.

对于 SRE，我们链接到我们的操作 GitLab 实例的详细差异，该实例具有有限的访问权限。
这使无法访问仅限于 SRE 的运营项目的员工和社区能够了解提议的配置更改。
通过拥有用于代码的公共 GitLab 实例和用于 [CI 管道](http://about.gitlab.com/stages-devops-lifecycle/continuous-integration/) 的私有实例，我们能够在保持单一工作流的同时同时确保我们不依赖 GitLab.com 进行配置更新。

## The lessons we learned along the way

## 我们一路上学到的教训

We have learned a few things along the way, lessons that we are applying to future migrations and new deployments into Kubernetes.

在此过程中，我们学到了一些东西，我们将这些经验教训应用到未来迁移和新部署到 Kubernetes 中。

### Increased billing from cross-AZ traffic

### 增加跨可用区流量的计费

![git egress](http://about.gitlab.com/images/blogimages/a_year_of_k8s/git_egress.png)
Daily egress bytes/day from the Git storage fleet on GitLab.com.

GitLab.com 上 Git 存储队列的每日出口字节数/天。

Google divides its network into regions and regions are divided into availability zones (AZs).
Because of the large amount of bandwidth required for Git hosting, it is important we are cognizant of network egress. For internal network traffic, egress is only free-of-charge if it remains in a single AZ.
At the time of writing this blog post, we deliver approximately 100TB on a typical work day for just Git repositories.
On legacy VM topology, services that were previously colocated on the same VMs are now running in Kubernetes pods.
This mean some network traffic that was previously local to a VM can now potentially traverse availability zones.

Google 将其网络划分为区域，而区域又划分为可用区 (AZ)。
由于 Git 托管需要大量带宽，因此了解网络出口很重要。对于内部网络流量，只有保留在单个可用区中的出口才免费。
在撰写本博文时，我们在一个典型工作日内仅为 Git 存储库交付大约 100TB。
在旧的 VM 拓扑中，以前在同一 VM 上共存的服务现在在 Kubernetes pod 中运行。
这意味着以前在 VM 本地的一些网络流量现在可能会穿越可用性区域。

Regional GKE clusters provide the convenience of spanning multiple availability zones for redundancy.
We are considering [splitting the regional GKE cluster into single zonal clusters](https://gitlab.com/gitlab-com/gl-infra/delivery/-/issues/1175) for services that use a lot of bandwidth to avoid network egress charges while maintaining redundancy at the cluster level.

区域 GKE 集群提供了跨越多个可用区以实现冗余的便利。
我们正在考虑[将区域 GKE 集群拆分为单个区域集群](https://gitlab.com/gitlab-com/gl-infra/delivery/-/issues/1175) 用于使用大量带宽以避免网络出口收费，同时保持集群级别的冗余。

### Resource limits, requests, and scaling

### 资源限制、请求和扩展

![replicas](http://about.gitlab.com/images/blogimages/a_year_of_k8s/replicas.png)
Number of replicas servicing production traffic on registry.gitlab.com, Registry traffic reaches it peak at ~15:00UTC.

为 registry.gitlab.com 上的生产流量提供服务的副本数量，注册表流量在大约 15:00UTC 达到峰值。

Our migration story began in August 2019 when we migrated the GitLab Container Registry to Kubernetes, the first service to move.
Though this was a critical and high traffic service, it was a good choice for the first migration because it is a stateless application with only a few external dependencies.
The first challenge we experienced was the large number of evicted pods, due to memory constraints on our nodes. 

我们的迁移故事始于 2019 年 8 月，当时我们将 GitLab Container Registry 迁移到了 Kubernetes，这是第一个迁移的服务。
尽管这是一项关键且高流量的服务，但它是第一次迁移的不错选择，因为它是一个只有少数外部依赖项的无状态应用程序。
由于我们节点的内存限制，我们遇到的第一个挑战是大量被驱逐的 pod。

This required multiple changes to requests and limits. We found that with an application that increases its memory utilization over time, low requests (which reserves memory for each pod) and a generous hard limit on utilization was a recipe for node saturation and a high rate of evictions.
To adjust for this [we eventually decided to use higher requests and lower limit](https://gitlab.com/gitlab-com/gl-infra/delivery/-/issues/998#note_388983696) which took pressure off of the nodes and allowed pods to be recycled without putting too much pressure on the node.
After experiencing this once, we start our migrations with generous requests and limits that are close to the same value, and adjust down as needed.

这需要对请求和限制进行多次更改。我们发现，随着时间的推移，应用程序会增加其内存利用率，低请求（为每个 pod 保留内存）和对利用率的严格限制是节点饱和和高驱逐率的一个原因。
为了对此进行调整 [我们最终决定使用更高的请求和更低的限制](https://gitlab.com/gitlab-com/gl-infra/delivery/-/issues/998#note_388983696) 这减轻了节点的压力并允许 pod 被回收而不会对节点施加太大压力。
在经历过一次之后，我们以接近相同值的慷慨请求和限制开始我们的迁移，并根据需要进行调整。

### Metrics and logging

### 指标和日志记录

![registry-general](http://about.gitlab.com/images/blogimages/a_year_of_k8s/registry-general.png)
The Infrastructure department focuses on latency, error rates and saturation that have [Service-level objectives (SLOs)](https://en.wikipedia.org/wiki/Service-level_objective) that tie into our [overall system availability](https://gitlab.com/gitlab-com/dashboards-gitlab-com/-/metrics/sla-dashboard.yml?environment=1790496&duration_seconds=86400).

基础设施部门专注于延迟、错误率和饱和度，这些具有 [服务级别目标 (SLO)](https://en.wikipedia.org/wiki/Service-level_objective) 与我们的 [整体系统可用性](https://gitlab.com/gitlab-com/dashboards-gitlab-com/-/metrics/sla-dashboard.yml?environment=1790496&duration_seconds=86400)。

Over the past year, one of the major changes in the infrastructure department was improvements to how we monitor and manage SLOs.
SLOs allowed us to set targets on individual services which were monitored closely during the migration.
Yet even with this improved observability, we can't always see problems right away with our metric reporting and alerting.
For example, focusing on latency and error rates may not adequately cover all uses of the service that is being migrated.
We discovered this problem very early with some of the workloads that were moved into the cluster. This challenge was particularly acute when we had to validate features that do not receive many requests but have very specific configuration dependencies.
One of the key migration lessons was to also evaluate more than just monitoring metrics, but also logs, and the long-tail of errors in our monitoring.
Now for every migration we include a detailed list of log queries and plan a clear rollback procedures that can be handed off from one shift to the next in case of issues.

在过去的一年中，基础设施部门的主要变化之一是改进了我们监控和管理 SLO 的方式。
SLO 允许我们为在迁移过程中受到密切监控的单个服务设定目标。
然而，即使有了这种改进的可观察性，我们也不能总是立即发现我们的指标报告和警报问题。
例如，关注延迟和错误率可能无法充分涵盖正在迁移的服务的所有用途。
我们很早就在将一些工作负载移入集群时发现了这个问题。当我们必须验证未收到很多请求但具有非常具体的配置依赖项的功能时，这一挑战尤为严峻。
重要的迁移课程之一是不仅要评估监控指标，还要评估日志以及监控中的长尾错误。
现在，对于每次迁移，我们都包含一个详细的日志查询列表，并计划一个明确的回滚程序，以便在出现问题时可以从一个班次切换到下一个班次。

Serving the same requests on legacy VM infrastructure and Kubernetes simultaneously presented a unique challenge.
Unlike a lift-and-shift migration, running on legacy VMs and Kubernetes at the same time requires that our observability is compatible with both and combines metrics into one view.
Most importantly, we are using the same dashboards and log queries to ensure the observability is consistent during the transition period.

同时在传统 VM 基础架构和 Kubernetes 上处理相同的请求是一项独特的挑战。
与直接迁移不同，同时在传统 VM 和 Kubernetes 上运行要求我们的可观察性与两者兼容，并将指标组合到一个视图中。
最重要的是，我们使用相同的仪表板和日志查询来确保过渡期间的可观察性一致。

### Shifting traffic to the new cluster

### 将流量转移到新集群

For GitLab.com we maintain a segmentation of our fleet named the [canary stage](http://about.gitlab.com/handbook/engineering/#canary-testing).
This canary fleet services our internal projects, [or can be enabled by users](https://next.gitlab.com), and is deployed to first for infrastructure and application changes.
The first service we migrated started with taking limited traffic internally and we are continuing to use this method to ensure we are meeting our SLOs before committing all traffic to the cluster.
What this means for the migration is requests to internal projects are first routed to Kubernetes and then we slowly move other traffic to the cluster using HAProxy backend weighting.
We learned in the process of moving from VMs to Kubernetes that it was extremely beneficial for us to have an easy way to move traffic between the old and new infrastructure, and to keep legacy infrastructure available for rollback in the first few days after the migration.

对于 GitLab.com，我们维护了一个名为 [canary stage](http://about.gitlab.com/handbook/engineering/#canary-testing) 的车队分段。
这个金丝雀车队为我们的内部项目提供服务，[或可以由用户启用](https://next.gitlab.com)，并首先部署以进行基础设施和应用程序更改。
我们迁移的第一个服务从内部获取有限流量开始，我们将继续使用这种方法来确保在将所有流量提交到集群之前满足我们的 SLO。
这对于迁移意味着对内部项目的请求首先路由到 Kubernetes，然后我们使用 HAProxy 后端加权将其他流量慢慢移动到集群。
我们在从 VM 迁移到 Kubernetes 的过程中了解到，在新旧基础架构之间移动流量的简单方法，以及在迁移后的前几天保持旧基础架构可用于回滚对我们来说非常有益。

### Reserved pod capacity and utilization

### 预留 pod 容量和利用率

One problem we identified early was, while our pod start times for the Registry service were very short, our start times for Sidekiq took as long as [two minutes](https://gitlab.com/gitlab-org/charts/gitlab/-/issues/1775).
The long Sidekiq start times posed a challenge when we started moving workloads to Kubernetes for workers that need to process jobs quickly and scale fast. 

我们早期发现的一个问题是，虽然我们的 Registry 服务的 pod 启动时间很短，但我们的 Sidekiq 启动时间却长达 [两分钟](https://gitlab.com/gitlab-org/charts/gitlab/-/issues/1775)。
当我们开始将工作负载转移到 Kubernetes 以供需要快速处理作业和快速扩展的工作人员时，Sidekiq 的较长启动时间构成了挑战。

The lesson here was while the Horizontal Pod Autoscaler (HPA) works well in Kubernetes for adapting to increased traffic, it is also important to evaluate workload characteristics and set reserved pod capacity, especially for uneven demand.
In our case, we saw a sudden spike in jobs which caused a large scaling event which saturated CPU before we could scale the node pool.
While it is tempting to squeeze as much as possible out of the cluster, after experiencing some initial performance problems we now start with a generous pod budget and scale down later, while keeping a close eye on SLOs.
The pod start times for Sidekiq service have improved significantly and now average about 40 seconds. [Improving the pod start times](https://gitlab.com/gitlab-org/charts/gitlab/-/issues/1775) benefited GitLab.com as well as all the self-managed customers using the official GitLab Helm chart.

这里的教训是，虽然 Kubernetes 中的 Horizontal Pod Autoscaler (HPA) 在适应增加的流量方面运行良好，但评估工作负载特征和设置预留 Pod 容量也很重要，尤其是对于不均衡的需求。
在我们的例子中，我们看到作业突然激增，导致大规模扩展事件，在我们可以扩展节点池之前饱和 CPU。
虽然从集群中挤出尽可能多的东西很诱人，但在遇到一些初始性能问题后，我们现在从慷慨的 pod 预算开始，然后缩减规模，同时密切关注 SLO。
Sidekiq 服务的 pod 启动时间显着改善，现在平均约为 40 秒。 [改进 pod 启动时间](https://gitlab.com/gitlab-org/charts/gitlab/-/issues/1775) 使 GitLab.com 以及所有使用官方 GitLab Helm 图表的自我管理客户受益。

After transitioning each service, we enjoyed many benefits of using Kubernetes in production, including much faster and safer deploys of the application, scaling, and more efficient resource allocation.
The migration benefits extend beyond GitLab.com. With each improvement of the official Helm chart, we provide additional benefits to our self-managed customers.

转换每项服务后，我们享受了在生产中使用 Kubernetes 的许多好处，包括更快、更安全地部署应用程序、扩展和更有效的资源分配。
迁移优势超出了 GitLab.com。随着官方 Helm 图表的每一次改进，我们都会为自我管理的客户提供额外的好处。

We hope you enjoyed reading about our Kubernetes migration journey. As we continue to migrate more services to the cluster you can read more at following links:

我们希望您喜欢阅读我们的 Kubernetes 迁移之旅。随着我们继续将更多服务迁移到集群，您可以在以下链接中阅读更多内容：

- [Why are we migrating to Kubernetes?](http://about.gitlab.com/handbook/engineering/infrastructure/production/kubernetes/gitlab-com/)
- [GitLab.com on Kubernetes](http://about.gitlab.com/handbook/engineering/infrastructure/production/architecture/#gitlab-com-on-kubernetes)
- [Tracking epic for the GitLab.com Kubernetes Migration](https://gitlab.com/groups/gitlab-com/gl-infra/-/epics/112) 

- [我们为什么要迁移到 Kubernetes？](http://about.gitlab.com/handbook/engineering/infrastructure/production/kubernetes/gitlab-com/)
- [Kubernetes 上的 GitLab.com](http://about.gitlab.com/handbook/engineering/infrastructure/production/architecture/#gitlab-com-on-kubernetes)
- [GitLab.com Kubernetes 迁移的跟踪史诗](https://gitlab.com/groups/gitlab-com/gl-infra/-/epics/112)

