# **The USA TODAY NETWORK’s SRE team’s journey into Docker and Kubernetes. Was it worth it?**

# **今日美国网络的 SRE 团队进入 Docker 和 Kubernetes 的旅程。它值得吗？**

Jul 17, 2018 From: https://medium.com/usa-today-network/devops-you-build-it-you-run-it-8f972343eb8e

Docker and Kubernetes are some of the top trending technologies in our field with many companies working to move their applications to a container-based environment. The Site Reliability Team at USA TODAY NETWORK jumped at the opportunity to move our desktop web platform from Chef-based to container-based infrastructure.

Docker 和 Kubernetes 是我们领域中一些最热门的技术，许多公司都在努力将他们的应用程序迁移到基于容器的环境中。 USA TODAY NETWORK 的站点可靠性团队抓住了将我们的桌面 Web 平台从基于 Chef 的基础架构迁移到基于容器的基础架构的机会。

A little over a year ago, we were running more than 1,000 virtual machines configured with Chef to support 125+ desktop web applications. After managing several of our own Chef-based Kubernetes clusters, and finally landing on GKE-based clusters, we now run just over 2,000 containers and 700 pods to support the same applications.

一年多前，我们运行了 1000 多台配置了 Chef 的虚拟机，以支持 125 多个桌面 Web 应用程序。在管理了我们自己的几个基于 Chef 的 Kubernetes 集群并最终登陆基于 GKE 的集群之后，我们现在运行了超过 2,000 个容器和 700 个 Pod 来支持相同的应用程序。

We didn’t just migrate to run on the latest cool technology. Our ultimate goal for this project — and as site reliability engineers — was to make quantifiable improvements to our desktop web platform. Here are some questions we asked ourselves to confirm whether we met that objective.

我们不只是为了使用最新的酷技术而迁移。作为站点可靠性工程师，我们对这个项目的最终目标是对我们的桌面网络平台进行可量化的改进。以下是我们问自己的一些问题，以确认我们是否达到了该目标。

**Are we at least breaking even?**

**我们至少收支平衡了吗？**

First and foremost, are we saving the company money or at least cost-neutral? In our Chef-based architecture we ran two instances of each of our applications out of the east region and one out of the west. For larger sites we had to run larger servers — and more of them — to handle smaller increases in traffic so we didn’t have to wait for the VMs to scale. For some of our smaller sites, we would never scale. Essentially, we were over-provisioned until and unless there was significant traffic that required us to scale.

首先，我们是在为公司节省资金还是至少是成本中性？在我们基于 Chef 的架构中，我们在东部地区和西部地区运行每个应用程序的两个实例。对于更大的站点，我们必须运行更大的服务器——而且数量更多——以处理较小的流量增长，因此我们不必等待虚拟机扩展。对于我们的一些较小的站点，我们永远不会扩展。从本质上讲，我们被过度配置，除非有大量流量需要我们进行扩展。

With Kubernetes, we were able to do a little [math](http://medium.com/usa-today-network/scaling-at-scale-i-was-told-there-would-be-no-math-6e7bdbf74a3b) around the sizing of our pods and get closer to only allocating what our sites need for typical traffic — leveraging horizontal pod autoscaling when traffic increases. We were also able to go from three regions down to two, since we could scale pods faster.

使用 Kubernetes，我们能够做一些 [math](http://medium.com/usa-today-network/scaling-at-scale-i-was-told-there-would-be-no-math-6e7bdbf74a3b) 围绕我们的 pod 的大小，并更接近于只分配我们的站点需要的典型流量——当流量增加时利用水平 pod 自动缩放。我们还能够从三个区域减少到两个区域，因为我们可以更快地扩展 pod。

After the dust settled, moving from VMs to containers reduced daily instance costs by 40%+.

尘埃落定后，从虚拟机迁移到容器后，每日实例成本降低了 40% 以上。

**Are we meeting our Key Performance Indicators?**

**我们是否达到我们的关键绩效指标？**

Our goal wasn’t necessarily to improve application performance, but to deliver the same performance (response times) in our new environment. After all, the applications didn’t change and, as described above, we saved money and reduced overall resources.

我们的目标不一定是提高应用程序性能，而是在我们的新环境中提供相同的性能（响应时间）。毕竟，应用程序没有改变，如上所述，我们节省了资金并减少了整体资源。

Where we ended up seeing performance gains was with scaling. With our Chef-based environment, we would have to wait for our cloud provider to provision a new virtual machine and then wait for the Chef run to complete. During periods of increased traffic, we could see degraded origin performance as we were waiting for the new virtual machines to spin up. This could take up to ten minutes in some instances.

我们最终看到性能提升的地方是扩展。在我们基于 Chef 的环境中，我们将不得不等待我们的云提供商提供一个新的虚拟机，然后等待 Chef 运行完成。在流量增加期间，我们可能会看到原始性能下降，因为我们正在等待新的虚拟机启动。在某些情况下，这可能需要十分钟。

On Kubernetes, we leveraged horizontal pod autoscaling to scale our pods based off of CPU utilization. Instead of waiting ten minutes for the VM and Chef run, we got the new pods up in less than 30 seconds. This allows us to handle traffic peaks better and, in turn, see better performance on our origin.

在 Kubernetes 上，我们利用水平 pod 自动缩放来根据 CPU 利用率来扩展我们的 pod。我们无需等待 10 分钟等待 VM 和 Chef 运行，而是在 30 秒内启动了新的 Pod。这使我们能够更好地处理流量峰值，进而在我们的源站上看到更好的性能。

The metrics indicate an improvement in availability: In addition to APM stats holding steady over time, we’ve seen a dramatic (3X) decrease in monthly alerts.

这些指标表明可用性有所改善：除了 APM 统计数据随着时间的推移保持稳定外，我们还发现每月警报数量急剧减少 (3 倍)。

**Are we doing less manual work?**

**我们做的手工工作减少了吗？**

As SREs, we don’t like manual tasks. We especially don’t like having to wake up in the middle of the night to manually resolve issues that we can handle with automation. If you ask me and my teammates, this may be the most important item on our list. By simply leveraging Kubernetes’ liveness and readiness probes, our sites essentially fix themselves. 

作为 SRE，我们不喜欢手动任务。我们特别不喜欢在半夜醒来手动解决我们可以通过自动化处理的问题。如果你问我和我的队友，这可能是我们名单上最重要的项目。通过简单地利用 Kubernetes 的活跃性和就绪性探测，我们的站点基本上可以自我修复。

In the Chef-based environment, our servers lost connectivity or crashed for many possible reasons which all required manual intervention from our team. We were running more than 1,000 virtual machines for our desktop platform, so issues were bound to arise. Our Help Desk would often page our on-call person, requiring them to logon and manually intervene to resolve issues.

在基于 Chef 的环境中，我们的服务器由于许多可能的原因而失去连接或崩溃，所有这些都需要我们团队的手动干预。我们为我们的桌面平台运行了 1,000 多台虚拟机，因此必然会出现问题。我们的帮助台通常会呼叫我们的待命人员，要求他们登录并手动干预以解决问题。

Now that we have implemented liveness and readiness probes, if something causes issues on the pod, it’s terminated automatically and a new pod is created. This may not cover every case, like when there is a dependent micro-service having issues, but it covers most manual interventions that our team would need to take.

现在我们已经实现了 liveness 和 readiness 探针，如果有什么东西在 pod 上引起问题，它会自动终止并创建一个新的 pod。这可能无法涵盖所有情况，例如当依赖微服务出现问题时，但它涵盖了我们团队需要采取的大多数手动干预。

Even when we need to intervene manually, we’ve seen drastic improvements. For example, we might need to redeploy sites to pick up new backend configurations or deploy the site with a new version. Previously, a single site could take 15–20 minutes to deploy. And to deploy all 125+ of our sites, it could take two hours or longer.

即使我们需要手动干预，我们也看到了巨大的改进。例如，我们可能需要重新部署站点以获取新的后端配置或使用新版本部署站点。以前，部署单个站点可能需要 15-20 分钟。要部署我们所有 125 多个站点，可能需要两个小时或更长时间。

Now leveraging helm-charts we can deploy a single site in under a minute to both the east and west regions. We can deploy all 125+ sites to both regions in under 20 minutes. We are also more confident in our deployment jobs on Kubernetes and our team members can focus on other sprint work instead of having to monitor our deployment jobs or wait for a site to redeploy.

现在利用 helm-charts，我们可以在一分钟内将一个站点部署到东部和西部地区。我们可以在 20 分钟内将所有 125 多个站点部署到这两个区域。我们对 Kubernetes 上的部署工作也更有信心，我们的团队成员可以专注于其他 sprint 工作，而不必监控我们的部署工作或等待站点重新部署。

You may be wondering: “What about the Kubernetes management side of things? Doesn’t that add more support and maintenance than the Chef environment?” When we started down the Kubernetes path, we rolled our own Kubernetes solution, based on Chef. At that time, the answer probably would have been yes. Just a few months ago, we still managed our own etcd clusters and masters. However, recently, we have migrated to GKE where Google manages the etcd cluster and masters for us. Our cluster operations team now handles RBAC, cluster upgrades, cluster scaling and implementing new functionality within our Kubernetes clusters.

您可能想知道：“Kubernetes 管理方面呢？这不是比 Chef 环境增加了更多的支持和维护吗？”当我们开始走 Kubernetes 之路时，我们推出了基于 Chef 的自己的 Kubernetes 解决方案。那时，答案可能是肯定的。就在几个月前，我们还在管理自己的 etcd 集群和 master。然而，最近，我们已经迁移到 GKE，Google 为我们管理 etcd 集群和 master。我们的集群运营团队现在处理 RBAC、集群升级、集群扩展和在我们的 Kubernetes 集群中实施新功能。

**Was it worth it?**

**它值得吗？**

In a word, yes! We have checked every box off our list when migrating our app to docker and Kubernetes. We have had several learning moments along the way, but those only improved our team and our entire organization’s move into container-based applications. We have seen a reduction in cost, improved our performance through faster scaling, and we are far less likely to be woken up in the middle of the night since migrating our desktop web platform to GKE. We are now a Docker- and Kubernetes-first shop and push all of our teams towards our shared platform.

一句话，对！将我们的应用程序迁移到 docker 和 Kubernetes 时，我们已经检查了列表中的每个框。在此过程中，我们经历了几次学习时刻，但这些时刻只改善了我们的团队和整个组织转向基于容器的应用程序。我们已经看到成本降低，通过更快的扩展提高了性能，并且自从将我们的桌面 Web 平台迁移到 GKE 后，我们在半夜被吵醒的可能性要小得多。我们现在是 Docker 和 Kubernetes 优先的商店，并将我们所有的团队推向我们的共享平台。

## [More from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----d20840757f05--------------------------------)

Gannett is the largest local news organization in America. Our brands include USA TODAY and 250+ newsrooms spanning 46 states. Gannett’s vastly expanded local-to-national platform reaches over 50% of the U.S. digital audience, including more Millennials than Buzzfeed.

Gannett 是美国最大的地方新闻机构。我们的品牌包括《今日美国》和遍布 46 个州的 250 多家新闻编辑室。 Gannett 大大扩展的本地到国家平台覆盖了超过 50% 的美国数字受众，其中包括比 Buzzfeed 更多的千禧一代。

[Read more from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----d20840757f05--------------------------------)

