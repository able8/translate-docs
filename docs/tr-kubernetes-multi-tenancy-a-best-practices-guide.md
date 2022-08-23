# Kubernetes Multi-Tenancy – A Best Practices Guide

# Kubernetes 多租户——最佳实践指南

Sep 21, 2020
https://loft.sh/blog/kubernetes-multi-tenancy-a-best-practices-guide

2020 年 9 月 21 日
https://loft.sh/blog/kubernetes-multi-tenancy-a-best-practices-guide

Table of Contents

目录

- [What Is Kubernetes Multi-Tenancy](http://loft.sh#what-is-kubernetes-multi-tenancy)
   - [Soft Multi-Tenancy](http://loft.sh#soft-multi-tenancy)
   - [Hard Multi-Tenancy](http://loft.sh#hard-multi-tenancy)
   - [Soft vs. Hard Multi-Tenancy](http://loft.sh#soft-vs-hard-multi-tenancy)
- [Limitations of Multi-Tenant Kubernetes](http://loft.sh#limitations-of-multi-tenant-kubernetes)
- [Reasons For Kubernetes Multi-Tenancy](http://loft.sh#reasons-for-kubernetes-multi-tenancy)
- [Kubernetes Multi-Tenancy Implementation Challenges](http://loft.sh#kubernetes-multi-tenancy-implementation-challenges)
   - [User Management](http://loft.sh#user-management)
   - [Fair Resource Sharing](http://loft.sh#fair-resource-sharing)
   - [Isolation](http://loft.sh#isolation)
- [Available Multi-Tenancy Solutions](http://loft.sh#available-multi-tenancy-solutions)
   - [kiosk](http://loft.sh#kiosk)
   - [Loft](http://loft.sh#loft)
- [Conclusion](http://loft.sh#conclusion)

- [什么是 Kubernetes 多租户](http://loft.sh#what-is-kubernetes-multi-tenancy)
  - [软多租户](http://loft.sh#soft-multi-tenancy)
  - [硬多租户](http://loft.sh#hard-multi-tenancy)
  - [软与硬多租户](http://loft.sh#soft-vs-hard-multi-tenancy)
- [多租户Kubernetes的局限性](http://loft.sh#limitations-of-multi-tenant-kubernetes)
- [Kubernetes 多租户的原因](http://loft.sh#reasons-for-kubernetes-multi-tenancy)
- [Kubernetes 多租户实施挑战](http://loft.sh#kubernetes-multi-tenancy-implementation-challenges)
  - [用户管理](http://loft.sh#user-management)
  - [公平资源共享](http://loft.sh#fair-resource-sharing)
  - [隔离](http://loft.sh#isolation)
- [可用的多租户解决方案](http://loft.sh#available-multi-tenancy-solutions)
  - [信息亭](http://loft.sh#kiosk)
  - [阁楼](http://loft.sh#loft)
- [结论](http://loft.sh#conclusion)

Kubernetes multi-tenancy is a topic that more and more organizations are interested in as their Kubernetes usage spreads out. However, since Kubernetes is not a multi-tenant system per se, getting multi-tenancy right comes with some challenges.

随着 Kubernetes 使用范围的扩大，越来越多的组织对 Kubernetes 多租户感兴趣。然而，由于 Kubernetes 本身并不是一个多租户系统，因此要正确实现多租户会带来一些挑战。

In this article, I will describe these challenges and how to overcome them as well as some useful tools for Kubernetes multi-tenancy. Before, I will explain what Kubernetes multi-tenancy actually means, what the differences between soft and hard multi-tenancy are, and why this is such a relevant topic right now.

在本文中，我将描述这些挑战以及如何克服它们，以及一些用于 Kubernetes 多租户的有用工具。之前，我将解释 Kubernetes 多租户的真正含义，软多租户和硬多租户之间的区别，以及为什么这是一个如此相关的话题。

## [\#](http://loft.sh\#what-is-kubernetes-multi-tenancy) What Is Kubernetes Multi-Tenancy

## [\#](http://loft.sh\#what-is-kubernetes-multi-tenancy) 什么是 Kubernetes 多租户

Multi-tenancy in Kubernetes means that a cluster and its control plane are shared by multiple users, workloads, or applications. It is the opposite of single-tenancy, where only one user uses a whole Kubernetes cluster.

Kubernetes 中的多租户意味着集群及其控制平面由多个用户、工作负载或应用程序共享。它与单租户相反，只有一个用户使用整个 Kubernetes 集群。

There are different types of multi-tenancy, ranging from soft multi-tenancy to hard multi-tenancy.

有不同类型的多租户，从软多租户到硬多租户。

### [\#](http://loft.sh\#soft-multi-tenancy) Soft Multi-Tenancy

### [\#](http://loft.sh\#soft-multi-tenancy) 软多租户

Soft multi-tenancy is a form of multi-tenancy that does not have a strict isolation of the different users, workloads, or applications. It is thus an appropriate solution for trusted and known tenants, i.e. tenants that will not voluntarily abuse each other such as engineers within the same organization. The isolation between users is rather focused on preventing accidents and cannot prevent attacks on other tenants.

软多租户是一种多租户形式，它没有严格隔离不同的用户、工作负载或应用程序。因此，对于受信任和已知的租户来说，这是一个合适的解决方案，即不会自愿互相虐待的租户，例如同一组织内的工程师。用户之间的隔离更侧重于防止事故发生，不能防止对其他租户的攻击。

In terms of Kubernetes, soft multi-tenancy is typically associated with simple Kubernetes namespaces that the individual tenants are working in.

就 Kubernetes 而言，软多租户通常与单个租户正在使用的简单 Kubernetes 命名空间相关联。

### [\#](http://loft.sh\#hard-multi-tenancy) Hard Multi-Tenancy

### [\#](http://loft.sh\#hard-multi-tenancy) 硬多租户

Hard multi-tenancy enforces stricter isolation of tenants and so also prevents negative consequences of malicious behavior from other tenants. In addition to trusted tenants, it can thus also be used for tenants you do not trust such as many unconnected users and people from different organizations.

硬多租户强制执行更严格的租户隔离，因此还可以防止来自其他租户的恶意行为的负面后果。除了受信任的租户之外，它还可以用于您不信任的租户，例如许多未连接的用户和来自不同组织的人员。

To implement hard multi-tenancy in Kubernetes, you need a more advanced configuration for namespaces or [virtual Clusters (vClusters)](http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide).

要在 Kubernetes 中实现硬多租户，您需要对命名空间或 [虚拟集群 (vClusters)] 进行更高级的配置 (http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium =reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)。

### [\#](http://loft.sh\#soft-vs-hard-multi-tenancy) Soft vs. Hard Multi-Tenancy

### [\#](http://loft.sh\#soft-vs-hard-multi-tenancy) 软与硬多租户

While it sounds like soft and hard multi-tenancy are distinct opposites, they are rather endpoints of a range of different implementations. Some even argue that perfect hard multi-tenancy in Kubernetes cannot exist because Kubernetes is designed in a way that there is only one central control plane that always has to be shared by tenants introducing a potential vulnerability.

虽然听起来软和硬多租户是截然不同的对立面，但它们是一系列不同实现的端点。一些人甚至认为 Kubernetes 中不存在完美的硬多租户，因为 Kubernetes 的设计方式是只有一个中央控制平面必须始终由租户共享，这会引入潜在的漏洞。

In general, one can still say that hard multi-tenancy is generally “better” (as it can be used for trusted and non-trusted tenants), while soft multi-tenancy is easier to implement.

总的来说，人们仍然可以说硬多租户通常“更好”（因为它可以用于受信任和不受信任的租户），而软多租户更容易实现。

## [\#](http://loft.sh\#limitations-of-multi-tenant-kubernetes) Limitations of Multi-Tenant Kubernetes 

## [\#](http://loft.sh\#limitations-of-multi-tenant-kubernetes) Kubernetes 多租户的局限性

By implementing multi-tenancy, you automatically introduce “limitations” to your Kubernetes cluster because the tenants will be technically restricted compared to users of a single-tenant cluster and/or the tenants must consider the other tenants. An example of the limitations of a namespace-based multi-tenancy is that the tenants are not able to use CRDs, install Helm charts that use RBAC, or change cluster-wide settings such as the Kubernetes version.

通过实施多租户，您会自动为 Kubernetes 集群引入“限制”，因为与单租户集群的用户相比，租户在技术上会受到限制和/或租户必须考虑其他租户。基于命名空间的多租户限制的一个示例是租户无法使用 CRD、安装使用 RBAC 的 Helm 图表或更改集群范围的设置，例如 Kubernetes 版本。

Using virtual clusters mitigates these limitations as tenants have access to cluster-wide settings and resources in their vClusters. For this, virtual clusters feel much more like “real” clusters compared to namespaces, which is why vCluster-based multi-tenancy is a form of multi-tenancy that is relatively close to single-tenancy.

使用虚拟集群可以缓解这些限制，因为租户可以访问其 vCluster 中的集群范围设置和资源。为此，与命名空间相比，虚拟集群感觉更像“真实”集群，这就是为什么基于 vCluster 的多租户是一种相对接近单租户的多租户形式。

Nevertheless, introducing any form of multi-tenancy adds a layer of complexity to your system and comes with some restrictions for the tenants.

尽管如此，引入任何形式的多租户都会给您的系统增加一层复杂性，并对租户带来一些限制。

## [\#](http://loft.sh\#reasons-for-kubernetes-multi-tenancy) Reasons For Kubernetes Multi-Tenancy

## [\#](http://loft.sh\#reasons-for-kubernetes-multi-tenancy) Kubernetes 多租户的原因

You now might ask why Kubernetes multi-tenancy is relevant and why you could not just use many single-tenant clusters instead.

您现在可能会问为什么 Kubernetes 多租户是相关的，以及为什么不能只使用许多单租户集群。

Theoretically, it is possible to use many single-tenant clusters instead of a shared multi-tenant cluster. Some companies actually do this at the moment, e.g. to provide developers access to Kubernetes. However, such a solution is very inefficient and thus expensive, especially at a larger scale.

理论上，可以使用多个单租户集群来代替共享的多租户集群。一些公司目前实际上是这样做的，例如。为开发人员提供对 Kubernetes 的访问权限。然而，这种解决方案效率很低，因此很昂贵，尤其是在更大的规模上。

This is something that many organizations realize now when their [adoption of Kubernetes spreads within their organization](http://loft.sh/blog/why-adopting-kubernetes-is-not-the-solution/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) and more engineers get Kubernetes access: While it is very simple and not so costly to create one cluster per tenant/user during the initial experimentation phase with Kubernetes, it becomes a huge problem if (almost) every developer in an organization gets an own cluster during the later stages of [the cloud-native adoption journey](http://loft.sh/blog/the-journey-of-adopting-cloud-native-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide). Suddenly, you end up with dozens, hundreds, or even thousands of clusters that all cost money (the cluster management fees of public cloud providers also start to matter a lot now) and need to be efficiently managed, which is far from trivial.

当 [Kubernetes 的采用在他们的组织内传播] 时，许多组织现在都意识到了这一点（http://loft.sh/blog/why-adopting-kubernetes-is-not-the-solution/?utm_medium=reader&utm_source=other&utm_campaign =blog_kubernetes-multi-tenancy-a-best-practices-guide) 和更多工程师获得 Kubernetes 访问权限：虽然在 Kubernetes 的初始试验阶段为每个租户/用户创建一个集群非常简单且成本不高，但它变成了一个如果（几乎）组织中的每个开发人员在 [云原生采用之旅] 的后期阶段都拥有自己的集群，这将是一个巨大的问题（http://loft.sh/blog/the-journey-of-adopting-cloud-native -development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)。突然之间，你最终会拥有数十个、数百个甚至数千个集群，这些集群都需要花钱（公共云提供商的集群管理费用现在也开始变得很重要）并且需要进行有效管理，这绝非易事。

At this stage, the benefits of having a multi-tenant system outweigh the additional complications of it. It is simply much easier to manage a Kubernetes system with one or just a few clusters and sharing a cluster is also more efficient in terms of resource utilization as redundancies can be reduced.

在这个阶段，拥有多租户系统的好处超过了它带来的额外复杂性。管理具有一个或几个集群的 Kubernetes 系统要容易得多，并且共享一个集群在资源利用方面也更有效，因为可以减少冗余。

> For more information about this topic, also take a look at my article about [a comparison of individual clusters and shared clusters](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide).

> 关于这个话题的更多信息，也可以看看我关于 [单个集群和共享集群的比较]的文章（http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-多租户最佳实践指南）。

## [\#](http://loft.sh\#kubernetes-multi-tenancy-implementation-challenges) Kubernetes Multi-Tenancy Implementation Challenges

## [\#](http://loft.sh\#kubernetes-multi-tenancy-implementation-challenges) Kubernetes 多租户实施挑战

There are three major challenges that you will face when implementing Kubernetes multi-tenancy:

在实施 Kubernetes 多租户时，您将面临三个主要挑战：

### [\#](http://loft.sh\#user-management) User Management

### [\#](http://loft.sh\#user-management) 用户管理

**Challenge:** Most companies already have a user management system for their engineers in place somewhere. This could be in GitHub, Microsoft, Google, or any other service. Since you do not want to manage all users twice or even more often, you need to enable Single-Sign-On (SSO) for your Kubernetes system. 

**挑战：** 大多数公司已经在某个地方为其工程师配备了用户管理系统。这可能在 GitHub、Microsoft、Google 或任何其他服务中。由于您不想两次甚至更频繁地管理所有用户，因此您需要为 Kubernetes 系统启用单点登录 (SSO)。

**Solution:** With the CNCF sandbox project [dex](https://github.com/dexidp/dex), you can provide the tenants with an SSO-option. Dex is an OpenID Connect and OAuth2 provider that supports various identity providers including LDAP and SAML and thus can be used with [many user management systems](https://github.com/dexidp/dex#connectors). Therefore, dex is a very good option to solve the user management challenge for your multi-tenancy system with Kubernetes.

**解决方案：** 通过 CNCF 沙箱项目 [dex](https://github.com/dexidp/dex)，您可以为租户提供 SSO 选项。 Dex 是一个 OpenID Connect 和 OAuth2 提供程序，支持包括 LDAP 和 SAML 在内的各种身份提供程序，因此可以与 [许多用户管理系统](https://github.com/dexidp/dex#connectors) 一起使用。因此，使用 Kubernetes 解决多租户系统的用户管理挑战，dex 是一个非常好的选择。

### [\#](http://loft.sh\#fair-resource-sharing) Fair Resource Sharing

### [\#](http://loft.sh\#fair-resource-sharing) 公平资源共享

**Challenge:** Since all tenants share the same underlying resources, i.e. network and computing resources, the second challenge is to ensure that the available resources are shared fairly between the tenants. This is important because you do not want to have one tenant consume all or an excessive amount of resources (accidentally or voluntarily) leaving the others unable to work. For this, you need to make sure that every tenant has appropriate usage limits.

**挑战：**由于所有租户共享相同的底层资源，即网络和计算资源，第二个挑战是确保租户之间公平共享可用资源。这很重要，因为您不希望一个租户消耗全部或过多的资源（意外或自愿）而导致其他租户无法工作。为此，您需要确保每个租户都有适当的使用限制。

**Solution:** The resource consumption of tenants can be limited Kubernetes-natively with Resource Quotas. Since users must specify CPU and memory limits if quotas are enabled for these resources, it makes sense to also set smart defaults via LimitRanges.

**解决方案：** Kubernetes-native 可以通过 Resource Quotas 限制租户的资源消耗。如果为这些资源启用配额，用户必须指定 CPU 和内存限制，因此还可以通过 LimitRanges 设置智能默认值。

### [\#](http://loft.sh\#isolation) Isolation

### [\#](http://loft.sh\#isolation) 隔离

**Challenge:** The third challenge is to isolate the different tenants from each other. This prevents tenants from interfering with each other. As described above, the degree of isolation is determining if you have [soft or hard multi-tenancy](http://loft.sh#soft-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) in place and if the system should only be used by trusted tenants or if it is also secured against voluntary attacks.

**挑战：** 第三个挑战是将不同的租户相互隔离。这可以防止租户相互干扰。如上所述，隔离程度决定了您是否拥有 [软或硬多租户](http://loft.sh#soft-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-最佳实践指南)以及该系统是否应仅由受信任的租户使用，或者是否还可以防止自愿攻击。

**Solution:** Kubernetes namespaces serve as basic isolation for tenants. Alternatively, [vClusters provide even more isolation than namespaces](https://loft.sh/features/virtual-kubernetes-clusters?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide), while they also give tenants more flexibility. Another factor that should be considered for isolating tenants is network traffic: Network policies should be configured in a way that by default all traffic is denied and only traffic within the same namespace, internet traffic for containers, and requests to the DNS are allowed. This makes it much harder for tenants to attack or interfere with each other.

**解决方案：** Kubernetes 命名空间用作租户的基本隔离。或者，[vCluster 提供比命名空间更多的隔离](https://loft.sh/features/virtual-kubernetes-clusters?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)，而它们还为租户提供了更大的灵活性。隔离租户应考虑的另一个因素是网络流量：网络策略的配置方式应默认拒绝所有流量，仅允许同一命名空间内的流量、容器的互联网流量和对 DNS 的请求。这使得租户更难相互攻击或干扰。

## [\#](http://loft.sh\#available-multi-tenancy-solutions) Available Multi-Tenancy Solutions

## [\#](http://loft.sh\#available-multi-tenancy-solutions) 可用的多租户解决方案

Some useful solutions exist already that help you to implement multi-tenancy with Kubernetes. Besides the previously mentioned [dex](https://github.com/dexidp/dex),[kiosk](https://github.com/kiosk-sh/kiosk) and [loft](https://loft.sh/features/kubernetes-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) are worth mentioning here.

一些有用的解决方案已经存在，可以帮助您使用 Kubernetes 实现多租户。除了前面提到的[dex](https://github.com/dexidp/dex)、[kiosk](https://github.com/kiosk-sh/kiosk)和[loft](https://loft.sh/features/kubernetes-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)在这里值得一提。

### [\#](http://loft.sh\#kiosk) kiosk

### [\#](http://loft.sh\#kiosk) 信息亭

[kiosk](https://github.com/kiosk-sh/kiosk) is an open-source multi-tenancy extension for Kubernetes. It is designed as a lightweight, pluggable, and customizable solution for any Kubernetes cluster and solves some of the multi-tenancy challenges in a simple way. This includes account separation, resource consumption limitation on a user level, and namespace templates for secure tenant isolation and self-service namespace initialization. While it does not provide an in-built user management, kiosk is a very good building block to develop an own multi-tenancy system.

[kiosk](https://github.com/kiosk-sh/kiosk) 是 Kubernetes 的开源多租户扩展。它被设计为适用于任何 Kubernetes 集群的轻量级、可插拔和可定制的解决方案，并以简单的方式解决了一些多租户挑战。这包括帐户分离、用户级别的资源消耗限制以及用于安全租户隔离和自助命名空间初始化的命名空间模板。虽然它不提供内置的用户管理，但信息亭是开发自己的多租户系统的一个非常好的构建块。

Even though kiosk has been released only a few months ago, it is already included in the [EKS Best Practices Guide for multi-tenancy by AWS](https://aws.github.io/aws-eks-best-practices/security/docs/multitenancy/). You can find a [detailed guide to set up kiosk with EKS here](https://aws.amazon.com/de/blogs/containers/set-up-soft-multi-tenancy-with-kiosk-on-amazon-elastic-kubernetes-service/).

尽管 kiosk 仅在几个月前发布，但它已经包含在 [AWS 的多租户 EKS 最佳实践指南](https://aws.github.io/aws-eks-best-practices/security/docs/multitenancy/)。您可以在此处找到 [使用 EKS 设置信息亭的详细指南](https://aws.amazon.com/de/blogs/containers/set-up-soft-multi-tenancy-with-kiosk-on-amazon-弹性 kubernetes 服务/)。

### [\#](http://loft.sh\#loft) Loft 

### [\#](http://loft.sh\#loft) 阁楼

[Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) is based on kiosk and provides a comprehensive solution for a multi-tenancy platform. Loft can be installed into any Kubernetes cluster and then lets tenants create namespaces and virtual Clusters on-demand. It cares for the user management (including SSO) and user isolation and lets the cluster admins define usage limits, so all previously mentioned multi-tenancy problems are resolved. Loft provides some additional features such as a sleep mode that leads to reduced cloud computing costs by shutting down unused namespaces and vClusters.

[Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)基于kiosk，为多租户平台提供全面的解决方案。 Loft 可以安装到任何 Kubernetes 集群中，然后让租户按需创建命名空间和虚拟集群。它关心用户管理（包括 SSO)和用户隔离，并让集群管理员定义使用限制，因此前面提到的所有多租户问题都得到了解决。 Loft 提供了一些附加功能，例如睡眠模式，通过关闭未使用的命名空间和 vCluster 来降低云计算成本。

Loft is a commercial offering with a free tier and is also included in the [EKS Best Practices Guide for multi-tenancy](https://aws.github.io/aws-eks-best-practices/multitenancy). While it was originally focused on multi-tenant development use cases, it can also be used for production use cases such as [cluster sharding](https://loft.sh/use-cases/kubernetes-cluster-sharding?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) or to [run multiple instances of a product in a shared cluster](https://loft.sh/use-cases/cloud-native-managed-products?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)

Loft 是具有免费层级的商业产品，也包含在 [EKS 多租户最佳实践指南](https://aws.github.io/aws-eks-best-practices/multitenancy) 中。虽然它最初专注于多租户开发用例，但它也可以用于生产用例，例如 [集群分片](https://loft.sh/use-cases/kubernetes-cluster-sharding?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)或[在共享集群中运行产品的多个实例](https://loft.sh/use-cases/cloud-native-managed-products?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)

## [\#](http://loft.sh\#conclusion) Conclusion

## [\#](http://loft.sh\#conclusion) 结论

Kubernetes multi-tenancy is one of the challenges that many organizations face at the moment as their Kubernetes adoption progresses and single-tenant solutions become increasingly infeasible.

Kubernetes 多租户是许多组织目前面临的挑战之一，因为他们的 Kubernetes 采用不断发展，单租户解决方案变得越来越不可行。

When implementing multi-tenancy with Kubernetes, you need to decide if you need hard multi-tenancy or if soft multi-tenancy is enough. In any case, you need to solve three major problems: How to manage the users/tenants, how to limit their resource usage, and how to isolate them from each other. Here, several tools such as dex, kiosk, and loft can help you so you get multi-tenancy with Kubernetes right more easily.

在使用 Kubernetes 实现多租户时，您需要决定是需要硬多租户还是软多租户就足够了。无论如何，您需要解决三个主要问题：如何管理用户/租户，如何限制他们的资源使用，以及如何相互隔离。在这里，dex、kiosk 和 loft 等多种工具可以帮助您更轻松地使用 Kubernetes 进行多租户。

[Multi-Tenancy](http://loft.sh/blog/tags/multi-tenancy/)[Guides](http://loft.sh/blog/tags/guides/) [vcluster](http://loft.sh/blog/tags/vcluster/)

[多租户](http://loft.sh/blog/tags/multi-tenancy/)[指南](http://loft.sh/blog/tags/guides/) [vcluster](http://loft.sh/blog/tags/vcluster/)

[Share on Twitter](https://twitter.com/intent/tweet?text=Kubernetes%20Multi-Tenancy%20%e2%80%93%20A%20Best%20Practices%20Guide&hashtags=kubernetes%2Cloft%2Ccontainers%2Cdevops&url=https%3a%2f%2floft.sh%2fblog%2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=twitter%26utm_content=blog-share) [Share on Linkedin](https: //www.linkedin.com/sharing/share-offsite/?title=Kubernetes%20Multi-Tenancy%20%e2%80%93%20A%20Best%20Practices%20Guide&url=https%3a%2f%2floft.sh%2fblog %2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=linkedin%26utm_content=blog-share&via=loft_sh) [Share on Facebook](https://facebook.com/sharer.php?u=https%3a%2f%2floft.sh%2fblog%2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=fb%26utm_content=blog-share) 

[分享到 Twitter](https://twitter.com/intent/tweet?text=Kubernetes%20Multi-Tenancy%20%e2%80%93%20A%20Best%20Practices%20Guide&hashtags=kubernetes%2Cloft%2Ccontainers%2Cdevops&url=https%3a%2f%2floft.sh%2fblog%2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=twitter%26utm_content=blog-share) [在Linkedin上分享](https: //www.linkedin.com/sharing/share-offsite/?title=Kubernetes%20Multi-Tenancy%20%e2%80%93%20A%20Best%20Practices%20Guide&url=https%3a%2f%2floft.sh%2fblog %2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=linkedin%26utm_content=blog-share&via=loft_sh) [在 Facebook 上分享](https://facebook.com/sharer.php?u=https%3a%2f%2floft.sh%2fblog%2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=fb%26utm_content=blog-share)

