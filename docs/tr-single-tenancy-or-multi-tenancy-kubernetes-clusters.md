# Choosing Single-Tenancy or Multi-Tenancy Kubernetes Clusters

# 选择单租户或多租户 Kubernetes 集群

29 July 2020 by [Yassine Jaffoo](http://www.appvia.io#author)

2020 年 7 月 29 日，作者 [Yassine Jaffo](http://www.appvia.io#author)

https://www.appvia.io/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters

https://www.appvia.io/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters

You reach a fork in the road when setting up the architecture for your Kubernetes workloads: Do you want to operate with **many** individualized clusters or **one** large one? It’s an important decision because the direction that you go will define how your teams work, how you allocate budget and how customer data is managed.

在为 Kubernetes 工作负载设置架构时，您遇到了一个岔路口：您想使用**许多**个性化集群还是**一个**大型集群？这是一个重要的决定，因为您前进的方向将决定您的团队如何工作、如何分配预算以及如何管理客户数据。

[Kubernetes gives you](https://www.appvia.io/blog/intro-guide-to-kubernetes) the flexibility when making these choices, so it really is entirely up to the structure of your team(s), and the goals and abilities of the organization. Instead of a solely technical decision, it needs to be one that considers the impact on the overall business. Below we’ll look at the different models that you can use to structure your Kubernetes clusters and the pros and cons of each of them.

[Kubernetes 为您提供](https://www.appvia.io/blog/intro-guide-to-kubernetes) 做出这些选择时的灵活性，因此这完全取决于您团队的结构，并且组织的目标和能力。它需要考虑到对整体业务的影响，而不是单纯的技术决策。下面我们将介绍可用于构建 Kubernetes 集群的不同模型以及它们各自的优缺点。

### What is Single-Tenancy?

### 什么是单租户？

![](https://www.appvia.io/media/pages/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters/4e8a267748-1629216749/single-tenancy-v-multi-tenancy-04-640x.jpg)

640x.jpg)

Single-tenancy is when a single application or workload, dedicated to a single tenant, lives on a single cluster. Emphasis on singularity. In this model, every cluster is purpose-built, and there is a high-degree of separation: workloads, data and teams are all separated. Within the single-tenancy structure, there are tweaks and variations that you can make to your clusters to best suit your customers and your team. More on those later.

单租户是指专用于单个租户的单个应用程序或工作负载位于单个集群上。强调单一性。在这个模型中，每个集群都是专门构建的，并且存在高度的分离：工作负载、数据和团队都是分离的。在单租户结构中，您可以对集群进行调整和变化，以最适合您的客户和团队。稍后再谈。

### What is Multi-Tenancy?

### 什么是多租户？

![](https://www.appvia.io/media/pages/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters/8fd9d8faaa-1629216810/single-tenancy-v-multi-tenancy-05-640x.jpg)

640x.jpg)

Multi-tenancy, on the other hand, is when _multiple_ tenants operate on a single cluster. A tenant could be defined as users, clients and/or workloads, but the important distinction of this model is that all of the tenants share the cluster’s resources. It’s best-practice to isolate tenants as much as possible within a multi-tenant cluster so that one tenant can’t attack others or monopolise the shared resources.

另一方面，多租户是指_多个_租户在单个集群上运行。租户可以定义为用户、客户端和/或工作负载，但此模型的重要区别在于所有租户共享集群的资源。最好的做法是在多租户集群中尽可能地隔离租户，这样一个租户就不能攻击其他租户或垄断共享资源。

### Comparing the two

### 比较两者

When looking at the different architectures, consider these five core components: **Cost**, **Security**, **Reliability** and **Operational Overhead.** It's likely that you'll see a clear front-runner by matching up the strengths of one model with the priorities and capabilities of your business, so let's scope single-tenancy and multi-tenancy out side by side against each of these considerations.

在查看不同的架构时，请考虑以下五个核心组件：**成本**、**安全性**、**可靠性**和**运营开销**。您可能会看到明显的领先者通过将一个模型的优势与您的业务的优先级和能力相匹配，让我们根据这些考虑中的每一个并排确定单租户和多租户的范围。

![](https://www.appvia.io/media/pages/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters/7421f8f7ca-1629216832/single-tenancy-v-multi-tenancy-06-640x.jpg)

640x.jpg)

### 1\. How much is it going to cost?

### 1. 要花多少钱？

#### SINGLE-TENANCY

#### 单租

It’s naturally more expensive to operate many clusters, because you’ll require a lot of resources. Each and every Kubernetes cluster has a set of management resource requirements: master nodes, control plane components, monitoring solutions etc.

运行多个集群自然会更昂贵，因为您需要大量资源。每个 Kubernetes 集群都有一组管理资源需求：主节点、控制平面组件、监控解决方案等。

There’s no getting around it - if you have a bunch of smaller clusters, you’re going to dedicate a good portion of your overall resources to support the bare-bones management functions mentioned above for each one.

没有办法绕过它 - 如果您有一堆较小的集群，您将投入整体资源的很大一部分来支持上面提到的每个集群的基本管理功能。

The extra resource, of course, comes at a cost.

当然，额外的资源是有代价的。

#### MULTI-TENANCY

####  多租户

Because the aforementioned resources are shared between workloads in a multi-tenancy model, you need less of them. So, it costs less.

由于上述资源在多租户模型中的工作负载之间共享，因此您需要的资源更少。所以，它的成本更低。

You’ll be able to reuse all cluster-wide services like Ingress controllers, monitoring, logging and load balancers. A good example of resource-saving within the multi-tenancy model is the amount of master nodes you’ll need. Each cluster requires three master nodes, so if you have 10+ single-tenant clusters you’re looking at _at least_ 30 master nodes in comparison to the humble three nodes on your large, shared cluster.

您将能够重用所有集群范围的服务，例如 Ingress 控制器、监控、日志记录和负载均衡器。多租户模型中节省资源的一个很好的例子是您需要的主节点数量。每个集群需要三个主节点，因此如果您有 10 个以上的单租户集群，那么与大型共享集群上的三个节点相比，您正在查看_至少_30 个主节点。

While you don’t have to purchase additional resources, consider the somewhat abstract cost of potential security breaches and administrative nuances, which could soar much higher than the initial cost for an extra thirty master nodes.

虽然您不必购买额外的资源，但请考虑潜在安全漏洞和管理细微差别的抽象成本，这可能比额外增加 30 个主节点的初始成本高得多。

### 2\. How secure is your customer data?

### 2. 您的客户数据有多安全？

#### SINGLE-TENANCY 

#### 单租

Single-tenant clusters ‘take home the gold’ when it comes to security. It’s a huge positive in this respect that the individual clusters don’t share resources, because it creates strong isolation between customers and applications.

在安全性方面，单租户集群“带回家”。在这方面，单个集群不共享资源是一个巨大的积极因素，因为它在客户和应用程序之间建立了强大的隔离。

#### MULTI-TENANCY

####  多租户

The level of isolation you see with a single-tenant cluster of course isn’t the same with a multi-tenant cluster. If several applications are running within the same cluster, and sharing resources as we mentioned above, they’re going to be interacting with each-other.

您在单租户集群中看到的隔离级别当然与多租户集群不同。如果多个应用程序在同一个集群中运行，并且如上所述共享资源，它们将相互交互。

This could put an application, and your customers’ data, at risk when applications interact in an undesirable way. Kubernetes has safe-guards in place to prevent security breaches as much as possible, like Role Based Access Control (RBAC), PodSecurityPolicies and NetworkPolicies. But you need experience on your team to be able to manage these security capabilities in the right way.

当应用程序以不良方式交互时，这可能会使应用程序和您的客户数据面临风险。 Kubernetes 有安全措施来尽可能防止安全漏洞，例如基于角色的访问控制 (RBAC)、PodSecurityPolicies 和 NetworkPolicies。但是您需要团队经验才能以正确的方式管理这些安全功能。

You can’t prevent every security breach from happening, so you should be prepared for the added risk.

您无法阻止所有安全漏洞的发生，因此您应该为增加的风险做好准备。

### 3\. How reliable is it?

### 3. 它有多可靠？

#### SINGLE-TENANCY

#### 单租

From a reliability standpoint, there are a few things to consider with a single-tenant approach:

从可靠性的角度来看，单租户方法需要考虑以下几点：

1. If a cluster breaks, it only affects that workload. You can mitigate the risk of multiple things failing at the same time.
2. You’ll likely have fewer users tinkering with things and making changes, and that controlled access naturally lowers the risk of something breaking in the first place.

1. 如果集群中断，它只会影响该工作负载。您可以降低多件事情同时失败的风险。
2. 修补和更改的用户可能会减少，而受控访问自然会降低首先出现问题的风险。

#### MULTI-TENANCY

####  多租户

The exact opposite is true for the multi-tenant cluster. There’s a single point of failure, so if one thing breaks the whole cluster and all of the workloads will be down. Also, with more users working on the cluster, there’s a greater opportunity for things to go awry.

多租户集群的情况正好相反。有一个单点故障，所以如果一件事情破坏了整个集群，所有的工作负载都会停止。此外，随着更多用户在集群上工作，出现问题的机会也更大。

It’s also important for you to realize before committing to a multi-tenant cluster is that it can only get _so_ big before it breaks. Kubernetes has [defined upper limits](https://kubernetes.io/docs/setup/best-practices/cluster-large/) of a cluster at around 5000 nodes, 150,000 Pods and 300,000 containers. But that doesn't mean you can even get to that size without issues - you could start seeing [issues with clusters](https://events19.lfasiallc.com/wp-content/uploads/2017/11/BoF_-Not-One-Size-Fits-All-How-to-Size-Kubernetes-Clusters_Guang-Ya-Liu-_-Sahdev-Zala.pdf) that have upwards of just 500 nodes. You’ll want to keep a close eye on the growth of your cluster, in particular the strain on the Kubernetes control plane, to keep it running efficiently.

在提交到多租户集群之前，您还需要意识到，在它崩溃之前它只能变大。 Kubernetes 有 [已定义的上限](https://kubernetes.io/docs/setup/best-practices/cluster-large/) 在大约 5000 个节点、150,000 个 Pod 和 300,000 个容器上的集群。但这并不意味着您甚至可以毫无问题地达到该规模 - 您可能会开始看到 [集群问题](https://events19.lfasiallc.com/wp-content/uploads/2017/11/BoF_-Not-One-Size-Fits-All-How-to-Size-Kubernetes-Clusters_Guang-Ya-Liu-_-Sahdev-Zala.pdf)只有 500 个以上的节点。您需要密切关注集群的增长，尤其是 Kubernetes 控制平面上的压力，以保持其高效运行。

### 4. What’s the operational overhead?

### 4. 运营开销是多少？

#### SINGLE-TENANCY

#### 单租

There are certain complexities to managing many clusters, involving the set-up and maintenance of authentication, authorization and other frameworks. If you go this route, it's in your best interest to adopt automation that makes these processes faster and less prone to error, whether the automation is set up by your team or [t](https://www.appvia.io/solutions/kore) [hrough a product/service](https://www.appvia.io/solutions/kore).

管理许多集群存在一定的复杂性，涉及身份验证、授权和其他框架的设置和维护。如果您走这条路，那么采用自动化使这些流程更快、更不容易出错符合您的最大利益，无论自动化是由您的团队设置还是 [t](https://www.appvia.io/solutions/kore) [通过产品/服务](https://www.appvia.io/solutions/kore)。

You can manage the operational overhead by configuring your single-tenancy clusters in a more efficient way. A single tenant cluster can either be based on team structure, application stack, or deployment environment.

您可以通过以更有效的方式配置单租户集群来管理运营开销。单个租户集群可以基于团队结构、应用程序堆栈或部署环境。

#### MULTI-TENANCY

####  多租户

Overall, managing a single cluster (or a few clusters) is easier than administrating many clusters. There are less administrative tasks and no need to worry about having to create new environments for new customers etc. You do most of the set-up a single time, and don't have to think about it again versus having to repeat the process constantly in a single-tenancy model.

总体而言，管理单个集群（或几个集群）比管理多个集群更容易。管理任务更少，无需担心必须为新客户创建新环境等。您只需一次完成大部分设置，无需再次考虑，而不必不断重复该过程在单租户模型中。

### Making the best all-around choice

### 做出最好的全能选择

There are strong considerations of each tenancy model, and the nuanced iterations that can be created within each, but you should have an overarching idea of the pros and cons of each one for your organization. What you choose for your workloads depends on what you want to achieve and the capabilities of your team.

每种租赁模型都需要考虑很多因素，并且可以在每个租赁模型中创建细微的迭代，但您应该对每种租赁模型对您的组织的优缺点有一个总体了解。您为工作负载选择的内容取决于您想要实现的目标以及团队的能力。

If a single-tenancy model is the best option, but you don’t have the specialized team to manage the running of all of your individualized clusters, consider a service or product to create that for your business. 

如果单租户模型是最佳选择，但您没有专门的团队来管理所有个性化集群的运行，请考虑为您的业务创建服务或产品。

You don’t want to increase the cognitive load of your development team if they don’t have the infrastructure background or knowledge. [Wayfinder](http://www.appvia.io/products/wayfinder) enables the automation of deployment environments and provides security guardrails for teams using Kubernetes. 

如果您的开发团队没有基础架构背景或知识，您不想增加他们的认知负担。 [Wayfinder](http://www.appvia.io/products/wayfinder) 实现部署环境的自动化，并为使用 Kubernetes 的团队提供安全护栏。

