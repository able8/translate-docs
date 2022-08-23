# How to Save More Than 2/3 of Engineers’ Kubernetes Cost

# 如何为工程师节省 2/3 以上的 Kubernetes 成本

Jul 7, 2020

2020 年 7 月 7 日

Table of Contents

目录

- [The Cost of the Cloud](http://loft.sh#the-cost-of-the-cloud)
- [Cloud-Native Workflows Cause Additional Cloud Cost](http://loft.sh#cloud-native-workflows-cause-additional-cloud-cost)
- [Inefficiency 1: Too Many Clusters](http://loft.sh#inefficiency-1-too-many-clusters)
   - [Problem: Many clusters are expensive](http://loft.sh#problem-many-clusters-are-expensive)
   - [Solution: Share a cluster](http://loft.sh#solution-share-a-cluster)
   - [Maximum Savings](http://loft.sh#maximum-savings)
   - [Example Calculation](http://loft.sh#example-calculation)
- [Inefficiency 2: Unused Resources](http://loft.sh#inefficiency-2-unused-resources)
   - [Problem: Computing Resources are often unused](http://loft.sh#problem-computing-resources-are-often-unused)
   - [Solution: Let spaces “sleep” when unused](http://loft.sh#solution-let-spaces-sleep-when-unused)
   - [Maximum Savings](http://loft.sh#maximum-savings-1)
   - [Example Calculation](http://loft.sh#example-calculation-1)
- [How to get these savings](http://loft.sh#how-to-get-these-savings)
   - [Example Calculation](http://loft.sh#example-calculation-2)
- [Private Clouds](http://loft.sh#private-clouds)
- [Conclusion](http://loft.sh#conclusion)

- [云的成本](http://loft.sh#the-cost-of-the-cloud)
- [云原生工作流导致额外的云成本](http://loft.sh#cloud-native-workflows-cause-additional-cloud-cost)
- [低效率1：集群过多](http://loft.sh#inefficiency-1-too-many-clusters)
  - [问题：许多集群都很昂贵](http://loft.sh#problem-many-clusters-are-expensive)
  - [解决方案：共享集群](http://loft.sh#solution-share-a-cluster)
  - [最大节省](http://loft.sh#maximum-savings)
  - [示例计算](http://loft.sh#example-calculation)
- [低效率2：未使用的资源](http://loft.sh#inefficiency-2-unused-resources)
  - [问题：计算资源经常被使用](http://loft.sh#problem-computing-resources-are-often-unused)
  - [解决方案：让空间在未使用时“休眠”](http://loft.sh#solution-let-spaces-sleep-when-unused)
  - [最大节省](http://loft.sh#maximum-savings-1)
  - [示例计算](http://loft.sh#example-calculation-1)
- [如何获得这些储蓄](http://loft.sh#how-to-get-these-savings)
  - [示例计算](http://loft.sh#example-calculation-2)
- [私云](http://loft.sh#private-clouds)
- [结论](http://loft.sh#conclusion)

## [\#](http://loft.sh\#the-cost-of-the-cloud) The Cost of the Cloud

## [\#](http://loft.sh\#the-cost-of-the-cloud) 云的成本

Using the cloud is not cheap. If you are using a public cloud, you will get a monthly bill for computing resources, traffic, and for any additional services. But also for a private cloud, the cost is substantial, taking into account the hardware, datacenter, and maintenance cost. (For simplicity, I will assume the public cloud case in the following, but I added a paragraph about private clouds at the end.)

使用云并不便宜。如果您使用的是公共云，您将获得计算资源、流量和任何附加服务的月度账单。但对于私有云来说，考虑到硬件、数据中心和维护成本，成本也很高。 （为简单起见，我将在下面假设公共云案例，但我在最后添加了一段关于私有云的内容。）

The widespread adoption of Kubernetes has not changed this situation fundamentally. [Some say it has only increased cost](https://www.reddit.com/r/kubernetes/comments/a91c5n/is_running_kubernetes_always_expensive/) because the Kubernetes clusters themselves consume resources (e.g. Master node, LoadBalancer, etcd). [Others are convinced that Kubernetes could at least reduce the cost](https://techbeacon.com/devops/one-year-using-kubernetes-production-lessons-learned) because it makes the use of the cloud resources more efficient. Overall, it probably depends on your setting and perspective, but the cost of the cloud remains an issue either way.

Kubernetes 的广泛采用并没有从根本上改变这种情况。 [有人说它只会增加成本](https://www.reddit.com/r/kubernetes/comments/a91c5n/is_running_kubernetes_always_expensive/) 因为 Kubernetes 集群本身会消耗资源（例如 Master 节点、LoadBalancer 等)。 [其他人相信 Kubernetes 至少可以降低成本](https://techbeacon.com/devops/one-year-using-kubernetes-production-lessons-learned) 因为它使云资源的使用更加高效。总体而言，这可能取决于您的设置和观点，但无论哪种方式，云的成本仍然是一个问题。

## [\#](http://loft.sh\#cloud-native-workflows-cause-additional-cloud-cost) Cloud-Native Workflows Cause Additional Cloud Cost

## [\#](http://loft.sh\#cloud-native-workflows-cause-additional-cloud-cost) 云原生工作流导致额外的云成本

For this, [reducing Kubernetes cloud computing cost](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost) is always attractive, especially with more and more companies adopting Kubernetes throughout the organization and introducing cloud-native development, which includes giving engineers access to [Kubernetes for development](http://loft.sh/blog/kubernetes-development-environments-comparison/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost), testing, experimentation, or CI/CD (level 2 or level 3 on the cloud-native journey [according to this article](http://loft.sh/blog/the-journey-of-adopting-cloud-native-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost)).

为此， [降低 Kubernetes 云计算成本](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-Engineers-kubernetes-cost) 总是很有吸引力，尤其是随着越来越多的公司在整个组织中采用 Kubernetes 并引入云原生开发，其中包括让工程师访问 [Kubernetes 进行开发](http://loft.sh/blog/kubernetes-development-environments-comparison/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost)、测试、实验或 CI/CD（2 级或云原生之旅的第 3 级 [根据本文](http://loft.sh/blog/the-journey-of-adopting-cloud-native-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-节省超过 2-3 个工程师的 kubernetes 成本）)。

In these cases, companies are sometimes hesitant to spend a lot of money on the cloud as this, contrary to spending money on production workloads, is not directly connected to revenue and customer benefit. Still, [giving developers direct Kubernetes access can make a lot of sense and can pay off in terms of higher stability and more efficient workflows, too](http://loft.sh/blog/is-kubernetes-still-just-an-ops-topic/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost). 

在这些情况下，公司有时不愿在云上花很多钱，因为这与在生产工作负载上花钱相反，与收入和客户利益没有直接关系。尽管如此，[让开发人员直接访问 Kubernetes 很有意义，并且还可以在更高的稳定性和更高效的工作流程方面获得回报](http://loft.sh/blog/is-kubernetes-still-just-an-ops-topic/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost)。

Making cloud-native development cheaper and easier can thus be an important driver for Kubernetes adoption, especially in small and medium-sized organizations. I will therefore describe 2 “obvious” inefficiencies of the engineers’ use of the cloud and easy ways of eliminating them to significantly reduce the cost of using Kubernetes. Since developers use the cloud differently than production systems, these cost-saving opportunities are most relevant for non-production environments.

因此，让云原生开发更便宜、更容易成为 Kubernetes 采用的重要驱动力，尤其是在中小型组织中。因此，我将描述工程师使用云的两个“明显”低效率以及消除它们以显着降低使用 Kubernetes 成本的简单方法。由于开发人员使用云的方式不同于生产系统，因此这些节省成本的机会与非生产环境最为相关。

## [\#](http://loft.sh\#inefficiency-1-too-many-clusters) Inefficiency 1: Too Many Clusters

## [\#](http://loft.sh\#inefficiency-1-too-many-clusters) 效率低下1：集群太多

### [\#](http://loft.sh\#problem-many-clusters-are-expensive) Problem: Many clusters are expensive

### [\#](http://loft.sh\#problem-many-clusters-are-expensive) 问题：很多集群都很昂贵

An easy option to give every engineer access to a cloud environment is to give everyone an own cluster. However, this is pretty inefficient for several reasons.

让每个工程师都可以访问云环境的一个简单选择是给每个人一个自己的集群。但是，由于几个原因，这是非常低效的。

- **Cluster Management Fee.** Big public cloud providers such as [AWS](https://aws.amazon.com/eks/pricing/?nc1=h_ls) and [Google Cloud](https://cloud.google.com/kubernetes-engine/pricing) charge a cluster management fee for their Kubernetes offerings. If the cluster is running all the time, only the cluster fee will add more than $70 to your bill for every cluster. And this is not including any computing resources that will come on top of that.
- **Redundant Kubernetes Functionality.** Since every cluster is running independently of the others, they all need to provide the whole Kubernetes functionality themselves. That means that, for example, every cluster needs its own load balancer that could actually be shared by many developers. This makes many clusters rather inefficient from a computing resource perspective.
- **No Central Oversight.** Many clusters make it hard to oversee the whole system. Even with cluster management tools, it is not always clear what each cluster is used for and if it is still used at all. In many cases, clusters will be created for a specific purpose such as a test, but often, they will not be deleted afterward so they continue to run and cost you money without any benefit.

- **集群管理费。** 大型公共云提供商，例如 [AWS](https://aws.amazon.com/eks/pricing/?nc1=h_ls) 和 [Google Cloud](https://cloud.google.com/kubernetes-engine/pricing)对其 Kubernetes 产品收取集群管理费。如果集群一直在运行，只有集群费用会为每个集群增加 70 多美元的账单。这还不包括任何计算资源。
- **冗余 Kubernetes 功能。** 由于每个集群都独立于其他集群运行，因此它们都需要自己提供整个 Kubernetes 功能。这意味着，例如，每个集群都需要自己的负载均衡器，实际上可以由许多开发人员共享。从计算资源的角度来看，这使得许多集群相当低效。
- **没有中央监督。** 许多集群使得监督整个系统变得困难。即使使用集群管理工具，也并不总是清楚每个集群的用途以及是否仍在使用。在许多情况下，会为特定目的（例如测试）创建集群，但通常不会在之后删除它们，因此它们会继续运行并花费您的资金而没有任何好处。

### [\#](http://loft.sh\#solution-share-a-cluster) Solution: Share a cluster

### [\#](http://loft.sh\#solution-share-a-cluster) 解决方法：共享一个集群

The solution to the first inefficiency of many non-production clusters is to share clusters among engineers. That means that only one cluster is used and developers are working with their own namespaces that limit them and isolate them from each other. [I wrote a separate article about a comparison of individual clusters and shared clusters](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost), but regarding the described cost efficiencies, sharing a cluster has the following advantages:

许多非生产集群的第一个低效率的解决方案是在工程师之间共享集群。这意味着只使用一个集群，并且开发人员正在使用他们自己的命名空间来限制它们并将它们彼此隔离。  [我单独写了一篇单独的集群和共享集群比较的文章](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost)，但就所描述的成本效率而言，共享集群具有以下优势：

- **Shared Cluster Management Fee.** Instead of paying the management fee for every developer as there are many clusters, you only need to pay it once. So, the cost per developer is reduced from about $70 per month to $70 divided by the number of developers and is thus decreasing with the number of developers.
- **Share Kubernetes Functionality.** Similar to the management fee, also the Kubernetes functionality can be shared and thus does not need to run redundantly. With more developers sharing it, the load will of course also be higher, but there will definitely still be cost savings from sharing a cluster.
- **Easier oversight.** In a shared cluster, it is much easier to control the users, especially if they are created with specialized solutions for this use case, such as [Loft](https://loft.sh). The cluster-admin has thus less effort to control the system and can shut down unused spaces easily. This shutdown can even be triggered automatically with a [“sleep mode”](http://loft.sh/docs/sleep-mode/basics?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost) (for more, see below).

- **共享集群管理费。**无需为每个开发者支付管理费，因为有很多集群，您只需支付一次。因此，每位开发人员的成本从每月约 70 美元减少到 70 美元除以开发人员的数量，因此随着开发人员的数量而减少。
- **共享 Kubernetes 功能。** 与管理费类似，Kubernetes 功能也可以共享，因此不需要冗余运行。随着更多的开发者共享它，负载当然也会更高，但是共享一个集群肯定会节省成本。
- **更容易监督。**在共享集群中，控制用户要容易得多，特别是如果他们是使用针对此用例的专门解决方案创建的，例如 [Loft](https://loft.sh)。因此，集群管理员控制系统的工作量更少，并且可以轻松关闭未使用的空间。这种关闭甚至可以通过 [“睡眠模式”](http://loft.sh/docs/sleep-mode/basics?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost）（更多信息，见下文)。

### [\#](http://loft.sh\#maximum-savings) Maximum Savings 

### [\#](http://loft.sh\#maximum-savings) 最大节省

The [Kubernetes cost savings from reducing the number of clusters](https://medium.com/faun/kubernetes-cost-savings-by-reducing-the-number-of-clusters-336378680995) by sharing a cluster are highest if there are **many clusters** used before, which is typical for **larger teams** or if clusters are used for **many purposes** such as in-cluster development, testing, and staging. The saving effect relative to the total cost is also particularly high if the **clusters are rather small and thus cheap**. In these cases, a cluster management fee and the redundancy is a comparably expensive cost driver.

[Kubernetes 减少集群数量所节省的成本](https://medium.com/faun/kubernetes-cost-savings-by-reducing-the-number-of-clusters-336378680995) 是最高的，如果之前使用过**许多集群**，这对于**较大的团队**是典型的，或者如果集群用于**多种目的**，例如集群内开发、测试和登台。如果**集群相当小并且因此便宜**，则相对于总成本的节省效果也特别高。在这些情况下，集群管理费和冗余是一个相对昂贵的成本驱动因素。

### [\#](http://loft.sh\#example-calculation) Example Calculation

### [\#](http://loft.sh\#example-calculation) 示例计算

Let’s take a look at a simple example to demonstrate the cost-saving opportunities of reducing the number of clusters by Kubernetes cluster sharing:

让我们看一个简单的例子来展示通过 Kubernetes 集群共享减少集群数量的成本节约机会：

If in a small team of 5 developers, everyone gets an own Kubernetes cluster and the clusters are always running, you have to pay $360 only for the cluster management fee. If they share one cluster, you only have to pay the cluster management fee for one cluster, which is $72 per month. In such a small team, you would thus already save $288 just in cluster management fees. (For GKE, one “zonal cluster” is free, so you would pay $288 in cluster management fees for 5 clusters and $0 for one cluster. Overall, the savings remain unchanged.)

如果在一个由 5 名开发人员组成的小团队中，每个人都有一个自己的 Kubernetes 集群并且集群一直在运行，那么您只需支付 360 美元的集群管理费。如果他们共享一个集群，您只需支付一个集群的集群管理费，即每月 72 美元。因此，在这样一个小团队中，仅集群管理费用就可以节省 288 美元。 （对于 GKE，一个“区域集群”是免费的，因此您需要为 5 个集群支付 288 美元的集群管理费，为一个集群支付 0 美元。总体而言，节省的费用保持不变。）

On top of that, you could save additional money by sharing the basic Kubernetes features and by a better control which prevents unused clusters. However, since these factors are hard to estimate, I do not add them to the calculation.

最重要的是，您可以通过共享基本的 Kubernetes 功能和更好的控制来防止未使用的集群来节省额外的资金。但是，由于这些因素很难估计，所以我没有将它们添加到计算中。

Of course, if you use commercial tools such as Loft, you need to pay for these tools, too. In the example of 5 developers using Loft, the cost would be $75.

当然，如果你使用 Loft 等商业工具，也需要为这些工具付费。在 5 个开发人员使用 Loft 的示例中，成本为 75 美元。

**Your total savings from sharing a cluster would so be at least $213 (60%) for only a small team of 5.**

**对于一个只有 5 人的小团队，您从共享集群中节省的总成本至少为 213 美元 (60%)。**

![](http://loft.sh/blog/images/content/cost-savings-clustersharing.jpg?nf_resize=fit&w=1040)

## [\#](http://loft.sh\#inefficiency-2-unused-resources) Inefficiency 2: Unused Resources

## [\#](http://loft.sh\#inefficiency-2-unused-resources) 效率低下2：未使用的资源

### [\#](http://loft.sh\#problem-computing-resources-are-often-unused) Problem: Computing Resources are often unused

### [\#](http://loft.sh\#problem-computing-resources-are-often-unused) 问题：计算资源经常未被使用

In non-production settings, the Kubernetes clusters are actually unused most of the time, while they cost you money all the time. One option to prevent this is to tell your engineers to reliably delete their stuff or shut down their spaces when they do not use them anymore. This is, however, not so easy to implement in reality. Sometimes, engineers may just forget about it as distractions come in or they are not at work when they should shut down their spaces. A few examples, when computing resources are typically unused:

在非生产环境中，Kubernetes 集群实际上大部分时间都没有使用，而它们一直在花钱。防止这种情况的一种选择是告诉您的工程师在他们不再使用它们时可靠地删除他们的东西或关闭他们的空间。然而，这在现实中实现起来并不容易。有时，工程师可能会因为分心而忘记它，或者当他们应该关闭他们的空间时他们不在工作。一些示例，当计算资源通常未被使用时：

- **In-cluster development.** Developers use Kubernetes in the cloud already during development, e.g. with special tools such as [DevSpace](https://github.com/devspace-cloud/devspace) or [Skaffold](https://github.com/GoogleContainerTools/skaffold). The used spaces continue to run even after work, at night, during weekends, holidays, or sick leave of the engineer.
- **Experiments.** Sometimes, engineers need to use the power of the cloud to test their applications. This is often the case for ML/AI applications that require a lot of computing resources or specific hardware. While these tests run for a while, the underlying hardware is sometimes even running when there are actually no experiments executed.
- **Staging/Testing.** Developers may push their finished code to a staging or testing environment where it is inspected before it can run in production. Similar to the in-cluster development situation, there are usually long periods of time when no tests are executed, no developer is pushing their code and no one is looking at the application in staging, e.g. at night. Still, the cloud environment continues to run and cost you money.

- **集群内开发。** 开发人员在开发过程中已经在云中使用 Kubernetes，例如使用特殊工具，例如 [DevSpace](https://github.com/devspace-cloud/devspace) 或 [Skaffold](https://github.com/GoogleContainerTools/skaffold)。即使在下班后、晚上、周末、节假日或工程师病假期间，使用的空间也会继续运行。
- **实验。**有时，工程师需要使用云的力量来测试他们的应用程序。对于需要大量计算资源或特定硬件的 ML/AI 应用程序来说，情况往往如此。虽然这些测试运行了一段时间，但底层硬件有时甚至在实际上没有执行任何实验时也在运行。
- **暂存/测试。** 开发人员可以将完成的代码推送到暂存或测试环境，在该环境中对其进行检查，然后才能在生产中运行。与集群内开发情况类似，通常有很长一段时间没有执行测试，没有开发人员推送他们的代码，也没有人在 staging 中查看应用程序，例如晚上。尽管如此，云环境仍在继续运行并花费您金钱。

### [\#](http://loft.sh\#solution-let-spaces-sleep-when-unused) Solution: Let spaces “sleep” when unused 

### [\#](http://loft.sh\#solution-let-spaces-sleep-when-unused) 解决方案：让空间在未使用时“休眠”

The solution to this problem is to activate a so-called sleep mode of unused spaces in the cloud. This means that the containers will be terminated while all persistent data will be preserved. The computing resources can thus be scaled down and you do not have to pay for them anymore. When they are needed again, they can be scaled up and the containers will be restarted automatically again with all the persistent data immediately available. In tools such as [Loft](https://loft.sh), a sleep mode can either be started manually or automatically after a pre-defined period of time.

这个问题的解决方案是激活云中未使用空间的所谓睡眠模式。这意味着容器将被终止，而所有持久数据都将被保留。计算资源因此可以按比例缩小，您不必再为它们付费。当再次需要它们时，它们可以按比例放大，容器将再次自动重新启动，所有持久数据都立即可用。在 [Loft](https://loft.sh) 等工具中，可以手动启动睡眠模式，也可以在预定义的时间段后自动启动睡眠模式。

Overall, this means that the cloud resources in Kubernetes that the engineers use are flexibly available, and you only pay during the times they are really needed. The advantage of an automatic sleep mode is that it does not affect the engineers at all, they do neither have to start nor to end it, so their [Kubernetes workflow](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost) remains as efficient as possible.

总体而言，这意味着工程师使用的 Kubernetes 中的云资源是灵活可用的，您只需在真正需要的时候付费。自动休眠模式的优点是它完全不影响工程师，他们不必开始也不必结束它，所以他们的 [Kubernetes 工作流程](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost) 保持尽可能高效。

### [\#](http://loft.sh\#maximum-savings-1) Maximum Savings

### [\#](http://loft.sh\#maximum-savings-1) 最大节省

The cost savings of a sleep mode are of course highest when there are a lot of computing resources generally used, i.e. **in larger teams or with complex applications and experiments with high resource requirements**. Another factor for the potential of the sleep mode is the **length of the idle times**.

当通常使用大量计算资源时，睡眠模式的成本节省当然最高，即**在较大的团队中或具有高资源要求的复杂应用和实验**。休眠模式的潜力的另一个因素是**空闲时间的长度**。

For a typical engineer with a 40 hour work week, the savings can be about 75% (128 of 168 hours per week are non-working hours, not accounting for holidays, sick leave, or other tasks of the engineer that do not require computing resources such as meetings). If the Kubernetes environment is used for some other tasks such as testing or experimentation, the savings can be even higher (potentially, there are days during which 100% could be saved when no experiments are run during that day).

对于一个每周工作 40 小时的典型工程师来说，节省大约 75%（每周 168 小时中有 128 小时是非工作时间，不考虑假期、病假或工程师的其他不需要计算的任务）资源，例如会议）。如果将 Kubernetes 环境用于其他一些任务，例如测试或实验，则节省的费用可能会更高（如果当天没有运行任何实验，可能会节省 100%）。

### [\#](http://loft.sh\#example-calculation-1) Example Calculation

### [\#](http://loft.sh\#example-calculation-1) 示例计算

Now, we can look at a similar example as before with a team of 5 engineers. Since we are now looking at computing costs only, we will ignore cluster management fees. We also assume that the average engineer works 8 hours a day on 5 days per week. For simplicity, we further ignore additional factors such as holidays, sick leave, and other tasks that do not require a Kubernetes access, e.g. meetings. In sum, [these factors can of course be substantial](https://www.quora.com/How-many-hours-in-average-do-software-engineers-Developers-or-coders-code-in-a-day-Do-developers-code-8-hours-every-day) but would only support the result of the example in terms of further saving opportunities.

现在，我们可以看一个与以前类似的例子，团队由 5 名工程师组成。由于我们现在只关注计算成本，我们将忽略集群管理费用。我们还假设普通工程师每周工作 5 天，每天工作 8 小时。为简单起见，我们进一步忽略了其他因素，例如假期、病假和其他不需要 Kubernetes 访问权限的任务，例如会议。总之，[这些因素当然可以很大](https://www.quora.com/How-many-hours-in-average-do-software-engineers-Developers-or-coders-code-in-a-day-Do-developers-code-8-hours-every-day)，但仅支持示例的结果以进一步节省机会。

Now, let’s say that the engineers need 4vCPUs and 15GB RAM each to run the application they are working on. Depending on the software, this could, of course, be quite a lot (e.g. for a simple website) or not nearly enough (e.g. for complex ML applications), so there is some variance here. In GKE, this is equal to an “n1-standard-4” machine that costs about $97 per month.

现在，假设工程师每个需要 4 个 vCPU 和 15GB RAM 来运行他们正在处理的应用程序。根据软件的不同，这当然可能很多（例如对于一个简单的网站）或远远不够（例如对于复杂的 ML 应用程序），因此这里存在一些差异。在 GKE 中，这相当于一台每月花费约 97 美元的“n1-standard-4”机器。

For the team of 5 engineers, you would so face computing resource cost of $485 per month. Since the engineers are only working 40 hours per week, the computing resources are not used 128 hours per week, which is more than 75% of the time.

对于 5 名工程师的团队，您将面临每月 485 美元的计算资源成本。由于工程师每周只工作 40 小时，计算资源每周没有使用 128 小时，超过 75% 的时间。

With a sleep mode configured in a way that sends unused spaces to sleep after 1 hour of inactivity, the computing resources are only active for 45 hours per week and the unused time is reduced to 5 hours (1 hour per day). For this, you only need to pay about 27% of the computing resource cost and you can reduce the time during which the resources are unused from 75% to 12.5%.

睡眠模式配置为在 1 小时不活动后将未使用空间发送到睡眠状态，计算资源每周仅活动 45 小时，未使用时间减少到 5 小时（每天 1 小时）。为此，您只需要支付大约 27% 的计算资源成本，就可以将资源未使用的时间从 75% 减少到 12.5%。

**In absolute terms, this means that you only have to pay about $130 for all 5 engineers per month saving about $355. Again, even after deducting $75 for tools to enable a sleep mode, such as Loft, you can get net savings of $280 every month.**

**绝对值，这意味着您每月只需为所有 5 名工程师支付约 130 美元，节省约 355 美元。同样，即使扣除了 75 美元用于启用睡眠模式的工具（如阁楼），您每月也可以获得 280 美元的净节省。**

![](http://loft.sh/blog/images/content/cost-savings-sleepmode.jpg?nf_resize=fit&w=1040)

## [\#](http://loft.sh\#how-to-get-these-savings) How to get these savings 

## [\#](http://loft.sh\#how-to-get-these-savings) 如何获得这些储蓄

With solutions such as [Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost), that is based on [kiosk](https://github.com/kiosk-sh/kiosk), it is possible to eliminate both of the described inefficiencies. Loft is a multi-tenancy manager for Kubernetes that allows engineers to create namespaces in a shared cluster on-demand. It also comes with an in-built sleep mode feature to save costs for unused computing resources.

使用 [Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost)等解决方案，基于[kiosk](https://github.com/kiosk-sh/kiosk)，可以消除所描述的两种低效率。 Loft 是 Kubernetes 的多租户管理器，允许工程师按需在共享集群中创建命名空间。它还具有内置的睡眠模式功能，可节省未使用的计算资源的成本。

To use Loft, you install it in a Kubernetes cluster that then becomes your management cluster for Loft. Afterward, you can connect other clusters (or the management cluster itself) to Loft and in these connected clusters, the engineers can start their isolated namespaces. Overall, Loft achieves the solutions I described before out-of-the-box:

要使用 Loft，请将其安装在 Kubernetes 集群中，然后该集群将成为 Loft 的管理集群。之后，您可以将其他集群（或管理集群本身）连接到 Loft，并且在这些连接的集群中，工程师可以启动他们隔离的命名空间。总的来说，Loft 实现了我之前描述的开箱即用的解决方案：

- **Shared Clusters.** Loft transforms your cluster into a multi-tenancy platform on which engineers can create their own namespaces whenever they need them. The user isolation and management are cared for by Loft and it also lets you limit the resource consumption on different levels (per user, per space). This also allows the cluster-admin to centrally control and manage the users and they can even take a direct look at the engineers’ spaces with space sharing to see if they are still in use.
- **Sleep Mode.** Another central feature of Loft is its sleep mode that automatically scales down unused and idle spaces after an individually determined period of time. When the space is needed again, it will be woken up automatically, too, and the engineers can proceed from where they ended. For this, engineers do not have to think about the cloud infrastructure anymore and can work freely and uninterruptedly.

- **共享集群。** Loft 将您的集群转变为一个多租户平台，工程师可以在需要时创建自己的命名空间。用户隔离和管理由 Loft 负责，它还允许您限制不同级别（每个用户、每个空间）的资源消耗。这也允许集群管理员集中控制和管理用户，他们甚至可以通过空间共享直接查看工程师的空间，看看他们是否仍在使用。
- **睡眠模式。** Loft 的另一个核心功能是其睡眠模式，可在单独确定的时间段后自动缩小未使用和空闲空间。当再次需要该空间时，它也会自动唤醒，工程师可以从他们结束的地方继续。为此，工程师不必再考虑云基础设施，可以自由不间断地工作。

### [\#](http://loft.sh\#example-calculation-2) Example Calculation

### [\#](http://loft.sh\#example-calculation-2) 示例计算

Combining the two examples from above, one can calculate the total cost savings of introducing shared clusters and a sleep mode for the sample engineering team of 5. For this, one only has to sum up the cost savings from both examples. The only thing that needs to be adjusted is the cost to get cluster-sharing and the sleep mode. Since both features are included in Loft, the usage fee for it only has to be accounted for once.

结合上面的两个示例，可以计算出引入共享集群和睡眠模式为 5 人的示例工程团队所节省的总成本。为此，只需将两个示例的成本节省相加即可。唯一需要调整的是集群共享和睡眠模式的成本。由于这两个功能都包含在阁楼中，因此只需计算一次使用费。

**Overall, it is so possible to save more than two-thirds of the cost, which is $568 for the sample engineering team. Reducing the cost per engineer from $169 to $55.4 makes providing engineers with direct access to Kubernetes in the cloud also feasible for smaller organizations.**

**总体而言，可以节省三分之二以上的成本，即样品工程团队的 568 美元。将每位工程师的成本从 169 美元降低到 55.4 美元，使工程师能够直接访问云中的 Kubernetes，这对于小型组织来说也是可行的。**

![](http://loft.sh/blog/images/content/cost-savings-clustersharing-sleepmode.jpg?nf_resize=fit&w=1040)

## [\#](http://loft.sh\#private-clouds) Private Clouds

## [\#](http://loft.sh\#private-clouds) 私有云

The described problems and solutions are mostly applicable to private clouds, too. Here, sometimes the usage-cost relationship is not so clear, especially if you own the hardware yourself. In these cases, the computing cost for the cloud is rather a fixed cost than a variable cost, except for power and network traffic cost. Some other additional factors such as cluster management fees are not an issue in private clouds, but the maintenance effort is usually a bigger challenge.

所描述的问题和解决方案也大多适用于私有云。在这里，有时使用成本关系并不那么清楚，尤其是如果您自己拥有硬件。在这些情况下，除了电力和网络流量成本之外，云计算的成本是固定成本而不是可变成本。其他一些额外的因素，例如集群管理费用，在私有云中不是问题，但维护工作通常是一个更大的挑战。

If you apply the described solutions of cluster sharing and a sleep mode in a private cloud, it can still save you cost, e.g. because central Kubernetes features are not running redundantly. Overall, the direct cost effect is potentially smaller, but there is an additional indirect effect: If you are using sleep mode and can configure user limits centrally, it is easier to “oversubscribe” the available computing resources as you already know that not all of them will be used at the same time.

如果您在私有云中应用所描述的集群共享和睡眠模式的解决方案，它仍然可以为您节省成本，例如因为中央 Kubernetes 功能没有冗余运行。总体而言，直接成本影响可能较小，但还有一个额外的间接影响：如果您使用睡眠模式并且可以集中配置用户限制，则更容易“超额订阅”可用计算资源，因为您已经知道并非所有它们将同时使用。

For example, you could activate the sleep mode so that all spaces for in-cluster development for developers are sleeping after work and then schedule complex processes such as ML experiments to run at night in spaces that will sleep again in the morning when the engineers' workday starts. Another example would be remote teams in different time zones that can alternately use the same computing resources. 

例如，您可以激活睡眠模式，以便开发人员在下班后进行集群内开发的所有空间都处于睡眠状态，然后安排复杂的流程（例如 ML 实验）在晚上运行，这些空间将在早上工程师工作时再次睡眠。工作日开始。另一个例子是不同时区的远程团队可以交替使用相同的计算资源。

Therefore, there is also a savings impact in private clouds that might be less immediate but if you are able to use the available computing resources more efficiently, this is generally beneficial. This can eventually lead to fewer hardware needs, which will also save you a lot of cost.

因此，私有云也有可能不太直接的节省影响，但如果您能够更有效地使用可用的计算资源，这通常是有益的。这最终会导致更少的硬件需求，这也将为您节省大量成本。

## [\#](http://loft.sh\#conclusion) Conclusion

## [\#](http://loft.sh\#conclusion) 结论

With engineers’ access to cloud environments, the cost of the cloud becomes an even more important issue. While simple solutions with many clusters and no central control might be a first start, they are definitely not efficient in terms of cost. Over time, more companies should thus move to more advanced solutions to limit the number of clusters and to reduce the time computing resources cost money while they are not used.

随着工程师对云环境的访问，云的成本成为一个更加重要的问题。虽然具有许多集群且没有中央控制的简单解决方案可能是第一次开始，但它们在成本方面绝对不是有效的。随着时间的推移，更多的公司应该转向更先进的解决方案，以限制集群的数量，并减少计算资源在不使用时花费金钱的时间。

With cluster sharing among engineers and an individually configured sleep mode, a lot of computing resource costs can be saved. Fortunately, the implementation of these solutions with off-the-shelf solutions such as [Loft](https://loft.sh) is now easier than ever and available even to small and medium-sized companies and not only to big players that might build such solutions themselves.

通过工程师之间的集群共享和单独配置的休眠模式，可以节省大量的计算资源成本。幸运的是，使用 [Loft](https://loft.sh) 等现成解决方案实施这些解决方案现在比以往任何时候都更容易，甚至适用于中小型公司，而不仅仅是大型企业可能会自己构建这样的解决方案。

In the provided example, the cost can so be reduced by more than two-thirds leading to savings of about $570 per month in a 5-person team. Maybe even more importantly, the cost per engineer is reduced from deterrent $169 to moderate $55.4 per month. This makes the adoption of Kubernetes throughout whole organizations not only cheaper and more attractive but also enables even smaller companies to start their cloud-native transformation.

在所提供的示例中，成本可以减少三分之二以上，从而在 5 人团队中每月节省约 570 美元。也许更重要的是，每位工程师的成本从每月 169 美元降低到 55.4 美元。这使得在整个组织中采用 Kubernetes 不仅成本更低、更具吸引力，而且还使更小的公司能够开始其云原生转型。

https://loft.sh/blog/how-to-save-more-than-2-3-of-engineers-kubernetes-cost 

https://loft.sh/blog/how-to-save-more-than-2-3-of-engineers-kubernetes-cost

