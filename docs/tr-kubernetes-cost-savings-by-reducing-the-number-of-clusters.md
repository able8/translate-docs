# Kubernetes Cost Savings By Reducing The Number Of Clusters

# Kubernetes 通过减少集群数量来节省成本

Oct 5, 2020

2020 年 10 月 5 日

Table of Contents

目录

- [Advantages of fewer clusters](http://loft.sh#advantages-of-fewer-clusters)
- [Disadvantages of fewer clusters](http://loft.sh#disadvantages-of-fewer-clusters)
- [Trade-Off: Efficiency vs. Stability](http://loft.sh#trade-off-efficiency-vs-stability)
- [How to Safely Reduce The Number of Required Clusters](http://loft.sh#how-to-safely-reduce-the-number-of-required-clusters)
   - [Make Differentiated Decisions](http://loft.sh#make-differentiated-decisions)
   - [Use Virtual Clusters](http://loft.sh#use-virtual-clusters)
   - [Implement Effective Multi-Tenancy](http://loft.sh#implement-effective-multi-tenancy)
- [Conclusion](http://loft.sh#conclusion)

- [更少集群的优势](http://loft.sh#advantages-of-fewer-clusters)
- [集群少的缺点](http://loft.sh#disadvantages-of-fewer-clusters)
- [权衡：效率与稳定性](http://loft.sh#trade-off-efficiency-vs-stability)
- [如何安全地减少所需集群的数量](http://loft.sh#how-to-safely-reduce-the-number-of-required-clusters)
  - [做出差异化决策](http://loft.sh#make-differentiated-decisions)
  - [使用虚拟集群](http://loft.sh#use-virtual-clusters)
  - [实施有效的多租户](http://loft.sh#implement-effective-multi-tenancy)
- [结论](http://loft.sh#conclusion)

When you are using Kubernetes at a larger scale and at different stages (development, testing, production), you will sooner or later face the question of how many clusters you should run. To find the right answer to this question is not easy as having many clusters has advantages and disadvantages compared to running only one or a few clusters, as discussed [in this article](https://learnk8s.io/how-many-clusters).

当你在更大规模和不同阶段（开发、测试、生产）使用 Kubernetes 时，你迟早会面临应该运行多少个集群的问题。要找到这个问题的正确答案并不容易，因为与仅运行一个或几个集群相比，拥有多个集群具有优势和劣势，正如[在本文中讨论的](https://learnk8s.io/how-many-clusters)。

However, your answer to this question also determines how much you will have to pay for your Kubernetes system. In general, it is cheaper to run only a few clusters, which is why I will explain in this article how you can reduce the number of clusters and thus save Kubernetes cost without negatively impacting your system.

但是，您对这个问题的回答也决定了您需要为 Kubernetes 系统支付多少费用。一般来说，只运行几个集群会更便宜，这就是为什么我将在本文中解释如何减少集群数量，从而在不影响系统的情况下节省 Kubernetes 成本。

## [\#](http://loft.sh\#advantages-of-fewer-clusters) Advantages of fewer clusters

## [\#](http://loft.sh\#advantages-of-fewer-clusters) 更少集群的优势

**No Redundancies:** Running many clusters means that you also have many API servers, etcds and other parts of the control plane. This leads to a lot of redundancy and inefficiency as most of these components will be underutilized. However, if you only run a few clusters, all applications and users of the same cluster can share a control plane, which will drive utilization up and cost down significantly.

**无冗余：** 运行许多集群意味着您还有许多 API 服务器、etcd 和控制平面的其他部分。这会导致大量冗余和低效，因为这些组件中的大多数都没有得到充分利用。但是，如果您只运行几个集群，则同一集群的所有应用程序和用户可以共享一个控制平面，这将大大提高利用率并显着降低成本。

**No Cluster Management Fees:** Some public cloud providers, including AWS and Google Cloud, charge their users a flat cluster management fee for every cluster. Naturally, if you have fewer clusters, you will pay less for the cluster management. The cluster management fee is particularly important for situations with many small clusters as the relative cost (about $70 per month per cluster) can represent more than 50% of the total cost here.

**无集群管理费：** 包括 AWS 和 Google Cloud 在内的一些公共云提供商会针对每个集群向其用户收取固定的集群管理费。自然，如果您拥有较少的集群，您将为集群管理支付更少的费用。集群管理费对于有许多小型集群的情况尤为重要，因为相对成本（每个集群每月约 70 美元）可能占总成本的 50% 以上。

**Efficient Administration:** In general, it is easier to manage and supervise a system with a limited number of clusters because you can get a much better overview of the system and do not have to repeat the same processes many times (e.g. updating every cluster). For this, reducing the number of clusters also reduces the admin effort for the clusters and will so ultimately lead to additional cost savings. 

**高效管理：** 一般来说，管理和监督集群数量有限的系统更容易，因为您可以更好地了解系统，并且不必多次重复相同的过程（例如更新每个集群）。为此，减少集群的数量也减少了集群的管理工作，最终将导致额外的成本节约。

> Cluster sharing, which leads to fewer clusters, is also the basis for an [internal Kubernetes platform](http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) that provides the engineers with [self-service namespaces](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) improving the [Kubernetes development workflow](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters). This illustrates that a reduction of the number of clusters can also be a positive side-effect of improvements on other parts of your Kubernetes system.

> 集群共享，导致集群更少，也是[内部Kubernetes平台]的基础(http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes -cost-savings-by-reducing-the-number-of-clusters) 为工程师提供[自助服务命名空间](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) 改进 [Kubernetes 开发工作流程](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters)。这说明集群数量的减少也可能是改进 Kubernetes 系统其他部分的积极副作用。

## [\#](http://loft.sh\#disadvantages-of-fewer-clusters) Disadvantages of fewer clusters

## [\#](http://loft.sh\#disadvantages-of-fewer-clusters) 集群少的缺点

**Single Point of Failure:** If more applications are running or more users are working on the same cluster, all of them will be impacted at the same time when the cluster crashes. This means that the stability of the cluster will be more important as the cluster becomes a “single” point of failure. Of course, “single” in single point of failure does not necessarily mean that there is only one cluster but the same goes analogously for a low number of clusters.

**单点故障：**如果有更多应用程序正在运行或更多用户在同一个集群上工作，那么当集群崩溃时，所有这些都会同时受到影响。这意味着随着集群成为“单”故障点，集群的稳定性将变得更加重要。当然，单点故障中的“单一”并不一定意味着只有一个集群，但对于少量集群也是如此。

**Multi-Tenancy:** If a cluster is shared by several users and applications, you need to implement an efficient [Kubernetes multi-tenancy](http://loft.sh/blog/kubernetes-multi-tenancy-best-practices-guide/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters). This can be quite challenging because it does not only comprise the isolation of tenants but also user management and enforcement of limits. Especially hard multi-tenancy, i.e. a system that securely isolates tenants that are not trusted, is hard to establish with Kubernetes as some components are always shared within one cluster.

**多租户：**如果一个集群被多个用户和应用共享，需要实现一个高效的 [Kubernetes多租户](http://loft.sh/blog/kubernetes-multi-tenancy-best-实践指南/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters)。这可能非常具有挑战性，因为它不仅包括租户的隔离，还包括用户管理和限制的执行。特别是硬多租户，即安全隔离不受信任的租户的系统，很难用 Kubernetes 建立，因为某些组件总是在一个集群中共享。

**Technical Limitations:** While a single Kubernetes cluster is quite scalable, its scalability will be limited at some point. For example, [Kubernetes v1.19 supports clusters with up to 5,000 nodes, 150,000 total pods, and 300,000 total containers](https://kubernetes.io/docs/setup/best-practices/cluster-large/). These limits are of course pretty high, but you might already encounter scalability problems before reaching the limits, e.g. due to networking capacity. Therefore, it is also technically infeasible to run everything on just a single cluster if you reach a very large scale.

**技术限制：** 虽然单个 Kubernetes 集群具有很强的可扩展性，但它的可扩展性在某些时候会受到限制。例如，[Kubernetes v1.19 支持最多 5000 个节点、150000 个 pod 和 300000 个容器的集群](https://kubernetes.io/docs/setup/best-practices/cluster-large/)。这些限制当然非常高，但您可能在达到限制之前就已经遇到了可扩展性问题，例如由于网络容量。因此，如果您达到非常大的规模，在单个集群上运行所有内容在技术上也是不可行的。

## [\#](http://loft.sh\#trade-off-efficiency-vs-stability) Trade-Off: Efficiency vs. Stability

## [\#](http://loft.sh\#trade-off-efficiency-vs-stability) 权衡：效率与稳定性

Simply put, in determining the right number of clusters for your use case, you face a trade-off between (cost-)efficiency and stability of your system. Therefore, it is clear that your goal should not be to end up with just a single cluster and try to run everything there. To efficiently save Kubernetes cost, you should rather use the lowest reasonably possible number of clusters that still ensures the stability of your system. And this “right” number heavily depends on the size of your system, the number of users, and the nature of your software.

简而言之，在为您的用例确定正确的集群数量时，您需要在（成本）效率和系统稳定性之间进行权衡。因此，很明显，您的目标不应该是最终只有一个集群并尝试在那里运行所有内容。为了有效地节省 Kubernetes 成本，您应该使用尽可能低的合理数量的集群，但仍能确保系统的稳定性。而这个“正确”的数字在很大程度上取决于您的系统大小、用户数量和软件的性质。

However, there are some approaches that can help you to drive down the number of required clusters and so will save you cost:

但是，有一些方法可以帮助您减少所需集群的数量，从而节省成本：

> Reducing the number of clusters is just one of several methods to save costs. In my [guide to reducing Kubernetes cost](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters), I also explain some other approaches that make your Kubernetes system more cost-efficient.

> 减少集群数量只是节省成本的几种方法之一。在我的 [降低 Kubernetes 成本指南](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters)，我还解释了一些使您的 Kubernetes 系统更具成本效益的其他方法。

## [\#](http://loft.sh\#how-to-safely-reduce-the-number-of-required-clusters) How to Safely Reduce The Number of Required Clusters

## [\#](http://loft.sh\#how-to-safely-reduce-the-number-of-required-clusters) 如何安全地减少所需集群的数量

### [\#](http://loft.sh\#make-differentiated-decisions) Make Differentiated Decisions 

### [\#](http://loft.sh\#make-differentiated-decisions) 做出差异化决策

Since there is no one-size-fits-all answer to the questions of how many clusters you should run, you need to make an individual decision for your use case. However, it is not enough to make this decision once, but you should rather repeat it for every part of your application and for every group of users.

由于对于应该运行多少个集群的问题没有一刀切的答案，因此您需要为您的用例做出单独的决定。但是，仅做出一次决定是不够的，您应该对应用程序的每个部分和每个用户组重复它。

This means that you should not just adopt a dogmatic policy of “every engineer and every application get an own cluster” or “we are doing everything in one cluster”. You really need to evaluate if it makes sense in the situation at hand.

这意味着您不应该仅仅采用“每个工程师和每个应用程序都有一个自己的集群”或“我们在一个集群中做所有事情”的教条政策。你真的需要评估它在手头的情况下是否有意义。

For example, there may be some applications, especially very critical applications running in production, that should get their own cluster so that these applications cannot be impacted by others. Other applications may share a cluster either because they are less important or are only used for developing and testing, which makes them less critical. However, also critical applications may run on a shared cluster, e.g. if they heavily depend on each other and would not work anyway if one of them failed.

例如，可能有一些应用程序，尤其是在生产中运行的非常关键的应用程序，应该拥有自己的集群，这样这些应用程序就不会受到其他应用程序的影响。其他应用程序可能会共享一个集群，因为它们不太重要，或者仅用于开发和测试，这使得它们不太重要。但是，关键应用程序也可以在共享集群上运行，例如如果它们严重依赖彼此，并且如果其中一个失败，无论如何都不会工作。

You also need to make the cluster decision for your engineers: Usually, not every engineer needs an own cluster but namespaces are often enough, so you might provide them with [self-service namespaces](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) that run on just one cluster. Still, some engineers may have special requirements and need to work directly on the Kubernetes configuration. Since this is not possible with simple namespaces, these engineers should get individual clusters or alternatively virtual clusters, which I will describe next.

您还需要为您的工程师做出集群决策：通常，并非每个工程师都需要自己的集群，但命名空间通常就足够了，因此您可以为他们提供 [自助服务命名空间](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters)只在一个集群上运行。不过，有些工程师可能有特殊要求，需要直接在 Kubernetes 配置上工作。由于这对于简单的命名空间是不可能的，因此这些工程师应该获得单独的集群或虚拟集群，我将在下面进行描述。

> If you want to learn more about the decision about clusters for development purposes, take a look at this [comparison of individual clusters and shared clusters for development](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters).

> 如果您想了解更多关于集群用于开发目的的决策，请查看此[用于开发的单个集群和共享集群的比较](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters)。

### [\#](http://loft.sh\#use-virtual-clusters) Use Virtual Clusters

### [\#](http://loft.sh\#use-virtual-clusters) 使用虚拟集群

Another way to decrease the number of clusters is to use [virtual clusters (vClusters)](http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters). Virtual clusters are very similar to “real” clusters but are virtualized and several virtual clusters can thus run on one physical cluster. Virtual clusters have some advantages that allow you to replace physical clusters with vClusters: 

另一种减少集群数量的方法是使用[虚拟集群（vClusters）](http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-通过减少集群数量来节省成本)。虚拟集群与“真实”集群非常相似，但它们是虚拟化的，因此可以在一个物理集群上运行多个虚拟集群。虚拟集群有一些优势，可以让您用 vCluster 替换物理集群：

Virtual clusters solve the multi-tenancy issue as they provide better isolation of tenants than namespaces, which is why it is easier to implement hard multi-tenancy with vClusters. Additionally, the tenants can configure their vCluster freely and independently from others, so that [vClusters are very good development environments](http://loft.sh/blog/kubernetes-virtual-clusters-as-development-environments/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) even for engineers that need to configure Kubernetes or use a specific version of it. Finally, you can use [vClusters for cluster sharding](https://loft.sh/use-cases/kubernetes-cluster-sharding?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters), so it is, for example, possible to handle more requests and thus push the feasible technical limits of Kubernetes.

虚拟集群解决了多租户问题，因为它们提供了比命名空间更好的租户隔离，这就是为什么使用 vCluster 更容易实现硬多租户的原因。此外，租户可以独立于其他人自由地配置他们的 vCluster，因此 [vCluster 是非常好的开发环境](http://loft.sh/blog/kubernetes-virtual-clusters-as-development-environments/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) 即使对于需要配置 Kubernetes 或使用特定版本的工程师也是如此。最后，您可以使用 [vClusters 进行集群分片](https://loft.sh/use-cases/kubernetes-cluster-sharding?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters)，因此可以处理更多请求，从而推动 Kubernetes 的可行技术极限。

Since virtual clusters are still running on a single physical cluster, this physical cluster remains a single point of failure. However, as the configurations and additional installations can be done on the vCluster level, the underlying cluster can be rather simple with only basic components, which makes this cluster less error-prone and thus more stable.

由于虚拟集群仍在单个物理集群上运行，因此该物理集群仍然是单点故障。但是，由于可以在 vCluster 级别上进行配置和附加安装，因此底层集群可以非常简单，只有基本组件，这使得该集群不易出错，因此更稳定。

### [\#](http://loft.sh\#implement-effective-multi-tenancy) Implement Effective Multi-Tenancy

### [\#](http://loft.sh\#implement-effective-multi-tenancy) 实施有效的多租户

One of the reasons why some companies prefer to use many clusters is that they are not sure how to implement multi-tenancy and how much effort this is. If more organizations get [best practices for Kubernetes multi-tenancy](http://loft.sh/blog/kubernetes-multi-tenancy-best-practices-guide/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) right, they could reduce the number of clusters, especially for non-production use cases.

一些公司喜欢使用多个集群的原因之一是他们不确定如何实现多租户以及这需要付出多少努力。如果更多组织获得 [Kubernetes 多租户的最佳实践](http://loft.sh/blog/kubernetes-multi-tenancy-best-practices-guide/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) 对，它们可以减少集群的数量，尤其是对于非生产用例。

Fortunately, Kubernetes has some helpful features for multi-tenancy, such as Role-based access control (RBAC) and Resource Quotas. In general, setting up a good multi-tenant system is mostly a one-time effort and requires good user management and strictly enforced user limits. To get this, some companies such as [Spotify](https://www.youtube.com/watch?v=vLrxOhZ6Wrg) started to build their own [internal Kubernetes platforms](http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters), but it is also possible to buy out-of-the-box solutions for this, such as [Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters).

幸运的是，Kubernetes 为多租户提供了一些有用的功能，例如基于角色的访问控制 (RBAC) 和资源配额。一般来说，建立一个好的多租户系统大多是一次性的，需要良好的用户管理和严格的用户限制。为此，一些公司如 [Spotify](https://www.youtube.com/watch?v=vLrxOhZ6Wrg) 开始构建自己的[内部 Kubernetes 平台](http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters)，但也可以购买开箱即用的解决方案为此，例如 [Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters)。

## [\#](http://loft.sh\#conclusion) Conclusion

## [\#](http://loft.sh\#conclusion) 结论

Reducing the number of clusters can lead to significant cost savings. However, just using a single cluster is often not the right solution as this could negatively impact the stability of your system. Finding the right number of clusters for your use case is therefore not an easy task because there is no general rule of how many clusters are optimal.

减少集群的数量可以显着节省成本。但是，仅使用单个集群通常不是正确的解决方案，因为这可能会对系统的稳定性产生负面影响。因此，为您的用例找到正确数量的集群并不是一件容易的事，因为没有关于多少集群是最优的一般规则。

It is rather a question you have to answer for every application and every group of engineers individually by assessing their situation and needs. However, it is possible to reduce the minimum number of necessary clusters by replacing some clusters with virtual clusters and by implementing efficient multi-tenancy.

这是一个您必须通过评估他们的情况和需求来为每个应用程序和每个工程师组单独回答的问题。但是，可以通过用虚拟集群替换一些集群和实施高效的多租户来减少必要集群的最小数量。

This will allow you to improve the cost-efficiency of your Kubernetes system without a negative impact on its stability.

这将使您能够提高 Kubernetes 系统的成本效益，而不会对其稳定性产生负面影响。

https://loft.sh/blog/kubernetes-cost-savings-by-reducing-the-number-of-clusters 

https://loft.sh/blog/kubernetes-cost-savings-by-reducing-the-number-of-clusters

