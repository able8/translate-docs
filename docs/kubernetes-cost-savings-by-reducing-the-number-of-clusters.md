# Kubernetes Cost Savings By Reducing The Number Of Clusters

Oct 5, 2020

Table of Contents

- [Advantages of fewer clusters](http://loft.sh#advantages-of-fewer-clusters)
- [Disadvantages of fewer clusters](http://loft.sh#disadvantages-of-fewer-clusters)
- [Trade-Off: Efficiency vs. Stability](http://loft.sh#trade-off-efficiency-vs-stability)
- [How to Safely Reduce The Number of Required Clusters](http://loft.sh#how-to-safely-reduce-the-number-of-required-clusters)
  - [Make Differentiated Decisions](http://loft.sh#make-differentiated-decisions)
  - [Use Virtual Clusters](http://loft.sh#use-virtual-clusters)
  - [Implement Effective Multi-Tenancy](http://loft.sh#implement-effective-multi-tenancy)
- [Conclusion](http://loft.sh#conclusion)

When you are using Kubernetes at a larger scale and at different stages (development, testing, production), you will sooner or later face the question of how many clusters you should run. To find the right answer to this question is not easy as having many clusters has advantages and disadvantages compared to running only one or a few clusters, as discussed [in this article](https://learnk8s.io/how-many-clusters).

However, your answer to this question also determines how much you will have to pay for your Kubernetes system. In general, it is cheaper to run only a few clusters, which is why I will explain in this article how you can reduce the number of clusters and thus save Kubernetes cost without negatively impacting your system.

## [\#](http://loft.sh\#advantages-of-fewer-clusters) Advantages of fewer clusters

**No Redundancies:** Running many clusters means that you also have many API servers, etcds and other parts of the control plane. This leads to a lot of redundancy and inefficiency as most of these components will be underutilized. However, if you only run a few clusters, all applications and users of the same cluster can share a control plane, which will drive utilization up and cost down significantly.

**No Cluster Management Fees:** Some public cloud providers, including AWS and Google Cloud, charge their users a flat cluster management fee for every cluster. Naturally, if you have fewer clusters, you will pay less for the cluster management. The cluster management fee is particularly important for situations with many small clusters as the relative cost (about $70 per month per cluster) can represent more than 50% of the total cost here.

**Efficient Administration:** In general, it is easier to manage and supervise a system with a limited number of clusters because you can get a much better overview of the system and do not have to repeat the same processes many times (e.g. updating every cluster). For this, reducing the number of clusters also reduces the admin effort for the clusters and will so ultimately lead to additional cost savings.

> Cluster sharing, which leads to fewer clusters, is also the basis for an [internal Kubernetes platform](http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) that provides the engineers with [self-service namespaces](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) improving the [Kubernetes development workflow](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters). This illustrates that a reduction of the number of clusters can also be a positive side-effect of improvements on other parts of your Kubernetes system.

## [\#](http://loft.sh\#disadvantages-of-fewer-clusters) Disadvantages of fewer clusters

**Single Point of Failure:** If more applications are running or more users are working on the same cluster, all of them will be impacted at the same time when the cluster crashes. This means that the stability of the cluster will be more important as the cluster becomes a “single” point of failure. Of course, “single” in single point of failure does not necessarily mean that there is only one cluster but the same goes analogously for a low number of clusters.

**Multi-Tenancy:** If a cluster is shared by several users and applications, you need to implement an efficient [Kubernetes multi-tenancy](http://loft.sh/blog/kubernetes-multi-tenancy-best-practices-guide/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters). This can be quite challenging because it does not only comprise the isolation of tenants but also user management and enforcement of limits. Especially hard multi-tenancy, i.e. a system that securely isolates tenants that are not trusted, is hard to establish with Kubernetes as some components are always shared within one cluster.

**Technical Limitations:** While a single Kubernetes cluster is quite scalable, its scalability will be limited at some point. For example, [Kubernetes v1.19 supports clusters with up to 5,000 nodes, 150,000 total pods, and 300,000 total containers](https://kubernetes.io/docs/setup/best-practices/cluster-large/). These limits are of course pretty high, but you might already encounter scalability problems before reaching the limits, e.g. due to networking capacity. Therefore, it is also technically infeasible to run everything on just a single cluster if you reach a very large scale.

## [\#](http://loft.sh\#trade-off-efficiency-vs-stability) Trade-Off: Efficiency vs. Stability

Simply put, in determining the right number of clusters for your use case, you face a trade-off between (cost-)efficiency and stability of your system. Therefore, it is clear that your goal should not be to end up with just a single cluster and try to run everything there. To efficiently save Kubernetes cost, you should rather use the lowest reasonably possible number of clusters that still ensures the stability of your system. And this “right” number heavily depends on the size of your system, the number of users, and the nature of your software.

However, there are some approaches that can help you to drive down the number of required clusters and so will save you cost:

> Reducing the number of clusters is just one of several methods to save costs. In my [guide to reducing Kubernetes cost](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters), I also explain some other approaches that make your Kubernetes system more cost-efficient.

## [\#](http://loft.sh\#how-to-safely-reduce-the-number-of-required-clusters) How to Safely Reduce The Number of Required Clusters

### [\#](http://loft.sh\#make-differentiated-decisions) Make Differentiated Decisions

Since there is no one-size-fits-all answer to the questions of how many clusters you should run, you need to make an individual decision for your use case. However, it is not enough to make this decision once, but you should rather repeat it for every part of your application and for every group of users.

This means that you should not just adopt a dogmatic policy of “every engineer and every application get an own cluster” or “we are doing everything in one cluster”. You really need to evaluate if it makes sense in the situation at hand.

For example, there may be some applications, especially very critical applications running in production, that should get their own cluster so that these applications cannot be impacted by others. Other applications may share a cluster either because they are less important or are only used for developing and testing, which makes them less critical. However, also critical applications may run on a shared cluster, e.g. if they heavily depend on each other and would not work anyway if one of them failed.

You also need to make the cluster decision for your engineers: Usually, not every engineer needs an own cluster but namespaces are often enough, so you might provide them with [self-service namespaces](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) that run on just one cluster. Still, some engineers may have special requirements and need to work directly on the Kubernetes configuration. Since this is not possible with simple namespaces, these engineers should get individual clusters or alternatively virtual clusters, which I will describe next.

> If you want to learn more about the decision about clusters for development purposes, take a look at this [comparison of individual clusters and shared clusters for development](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters).

### [\#](http://loft.sh\#use-virtual-clusters) Use Virtual Clusters

Another way to decrease the number of clusters is to use [virtual clusters (vClusters)](http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters). Virtual clusters are very similar to “real” clusters but are virtualized and several virtual clusters can thus run on one physical cluster. Virtual clusters have some advantages that allow you to replace physical clusters with vClusters:

Virtual clusters solve the multi-tenancy issue as they provide better isolation of tenants than namespaces, which is why it is easier to implement hard multi-tenancy with vClusters. Additionally, the tenants can configure their vCluster freely and independently from others, so that [vClusters are very good development environments](http://loft.sh/blog/kubernetes-virtual-clusters-as-development-environments/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) even for engineers that need to configure Kubernetes or use a specific version of it. Finally, you can use [vClusters for cluster sharding](https://loft.sh/use-cases/kubernetes-cluster-sharding?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters), so it is, for example, possible to handle more requests and thus push the feasible technical limits of Kubernetes.

Since virtual clusters are still running on a single physical cluster, this physical cluster remains a single point of failure. However, as the configurations and additional installations can be done on the vCluster level, the underlying cluster can be rather simple with only basic components, which makes this cluster less error-prone and thus more stable.

### [\#](http://loft.sh\#implement-effective-multi-tenancy) Implement Effective Multi-Tenancy

One of the reasons why some companies prefer to use many clusters is that they are not sure how to implement multi-tenancy and how much effort this is. If more organizations get [best practices for Kubernetes multi-tenancy](http://loft.sh/blog/kubernetes-multi-tenancy-best-practices-guide/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters) right, they could reduce the number of clusters, especially for non-production use cases.

Fortunately, Kubernetes has some helpful features for multi-tenancy, such as Role-based access control (RBAC) and Resource Quotas. In general, setting up a good multi-tenant system is mostly a one-time effort and requires good user management and strictly enforced user limits. To get this, some companies such as [Spotify](https://www.youtube.com/watch?v=vLrxOhZ6Wrg) started to build their own [internal Kubernetes platforms](http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters), but it is also possible to buy out-of-the-box solutions for this, such as [Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-cost-savings-by-reducing-the-number-of-clusters).

## [\#](http://loft.sh\#conclusion) Conclusion

Reducing the number of clusters can lead to significant cost savings. However, just using a single cluster is often not the right solution as this could negatively impact the stability of your system. Finding the right number of clusters for your use case is therefore not an easy task because there is no general rule of how many clusters are optimal.

It is rather a question you have to answer for every application and every group of engineers individually by assessing their situation and needs. However, it is possible to reduce the minimum number of necessary clusters by replacing some clusters with virtual clusters and by implementing efficient multi-tenancy.

This will allow you to improve the cost-efficiency of your Kubernetes system without a negative impact on its stability.

https://loft.sh/blog/kubernetes-cost-savings-by-reducing-the-number-of-clusters