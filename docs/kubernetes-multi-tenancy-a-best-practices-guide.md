# Kubernetes Multi-Tenancy – A Best Practices Guide

Sep 21, 2020
https://loft.sh/blog/kubernetes-multi-tenancy-a-best-practices-guide

Table of Contents

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

Kubernetes multi-tenancy is a topic that more and more organizations are interested in as their Kubernetes usage spreads out. However, since Kubernetes is not a multi-tenant system per se, getting multi-tenancy right comes with some challenges.

In this article, I will describe these challenges and how to overcome them as well as some useful tools for Kubernetes multi-tenancy. Before, I will explain what Kubernetes multi-tenancy actually means, what the differences between soft and hard multi-tenancy are, and why this is such a relevant topic right now.

## [\#](http://loft.sh\#what-is-kubernetes-multi-tenancy) What Is Kubernetes Multi-Tenancy

Multi-tenancy in Kubernetes means that a cluster and its control plane are shared by multiple users, workloads, or applications. It is the opposite of single-tenancy, where only one user uses a whole Kubernetes cluster.

There are different types of multi-tenancy, ranging from soft multi-tenancy to hard multi-tenancy.

### [\#](http://loft.sh\#soft-multi-tenancy) Soft Multi-Tenancy

Soft multi-tenancy is a form of multi-tenancy that does not have a strict isolation of the different users, workloads, or applications. It is thus an appropriate solution for trusted and known tenants, i.e. tenants that will not voluntarily abuse each other such as engineers within the same organization. The isolation between users is rather focused on preventing accidents and cannot prevent attacks on other tenants.

In terms of Kubernetes, soft multi-tenancy is typically associated with simple Kubernetes namespaces that the individual tenants are working in.

### [\#](http://loft.sh\#hard-multi-tenancy) Hard Multi-Tenancy

Hard multi-tenancy enforces stricter isolation of tenants and so also prevents negative consequences of malicious behavior from other tenants. In addition to trusted tenants, it can thus also be used for tenants you do not trust such as many unconnected users and people from different organizations.

To implement hard multi-tenancy in Kubernetes, you need a more advanced configuration for namespaces or [virtual Clusters (vClusters)](http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide).

### [\#](http://loft.sh\#soft-vs-hard-multi-tenancy) Soft vs. Hard Multi-Tenancy

While it sounds like soft and hard multi-tenancy are distinct opposites, they are rather endpoints of a range of different implementations. Some even argue that perfect hard multi-tenancy in Kubernetes cannot exist because Kubernetes is designed in a way that there is only one central control plane that always has to be shared by tenants introducing a potential vulnerability.

In general, one can still say that hard multi-tenancy is generally “better” (as it can be used for trusted and non-trusted tenants), while soft multi-tenancy is easier to implement.

## [\#](http://loft.sh\#limitations-of-multi-tenant-kubernetes) Limitations of Multi-Tenant Kubernetes

By implementing multi-tenancy, you automatically introduce “limitations” to your Kubernetes cluster because the tenants will be technically restricted compared to users of a single-tenant cluster and/or the tenants must consider the other tenants. An example of the limitations of a namespace-based multi-tenancy is that the tenants are not able to use CRDs, install Helm charts that use RBAC, or change cluster-wide settings such as the Kubernetes version.

Using virtual clusters mitigates these limitations as tenants have access to cluster-wide settings and resources in their vClusters. For this, virtual clusters feel much more like “real” clusters compared to namespaces, which is why vCluster-based multi-tenancy is a form of multi-tenancy that is relatively close to single-tenancy.

Nevertheless, introducing any form of multi-tenancy adds a layer of complexity to your system and comes with some restrictions for the tenants.

## [\#](http://loft.sh\#reasons-for-kubernetes-multi-tenancy) Reasons For Kubernetes Multi-Tenancy

You now might ask why Kubernetes multi-tenancy is relevant and why you could not just use many single-tenant clusters instead.

Theoretically, it is possible to use many single-tenant clusters instead of a shared multi-tenant cluster. Some companies actually do this at the moment, e.g. to provide developers access to Kubernetes. However, such a solution is very inefficient and thus expensive, especially at a larger scale.

This is something that many organizations realize now when their [adoption of Kubernetes spreads within their organization](http://loft.sh/blog/why-adopting-kubernetes-is-not-the-solution/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) and more engineers get Kubernetes access: While it is very simple and not so costly to create one cluster per tenant/user during the initial experimentation phase with Kubernetes, it becomes a huge problem if (almost) every developer in an organization gets an own cluster during the later stages of [the cloud-native adoption journey](http://loft.sh/blog/the-journey-of-adopting-cloud-native-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide). Suddenly, you end up with dozens, hundreds, or even thousands of clusters that all cost money (the cluster management fees of public cloud providers also start to matter a lot now) and need to be efficiently managed, which is far from trivial.

At this stage, the benefits of having a multi-tenant system outweigh the additional complications of it. It is simply much easier to manage a Kubernetes system with one or just a few clusters and sharing a cluster is also more efficient in terms of resource utilization as redundancies can be reduced.

> For more information about this topic, also take a look at my article about [a comparison of individual clusters and shared clusters](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide).

## [\#](http://loft.sh\#kubernetes-multi-tenancy-implementation-challenges) Kubernetes Multi-Tenancy Implementation Challenges

There are three major challenges that you will face when implementing Kubernetes multi-tenancy:

### [\#](http://loft.sh\#user-management) User Management

**Challenge:** Most companies already have a user management system for their engineers in place somewhere. This could be in GitHub, Microsoft, Google, or any other service. Since you do not want to manage all users twice or even more often, you need to enable Single-Sign-On (SSO) for your Kubernetes system.

**Solution:** With the CNCF sandbox project [dex](https://github.com/dexidp/dex), you can provide the tenants with an SSO-option. Dex is an OpenID Connect and OAuth2 provider that supports various identity providers including LDAP and SAML and thus can be used with [many user management systems](https://github.com/dexidp/dex#connectors). Therefore, dex is a very good option to solve the user management challenge for your multi-tenancy system with Kubernetes.

### [\#](http://loft.sh\#fair-resource-sharing) Fair Resource Sharing

**Challenge:** Since all tenants share the same underlying resources, i.e. network and computing resources, the second challenge is to ensure that the available resources are shared fairly between the tenants. This is important because you do not want to have one tenant consume all or an excessive amount of resources (accidentally or voluntarily) leaving the others unable to work. For this, you need to make sure that every tenant has appropriate usage limits.

**Solution:** The resource consumption of tenants can be limited Kubernetes-natively with Resource Quotas. Since users must specify CPU and memory limits if quotas are enabled for these resources, it makes sense to also set smart defaults via LimitRanges.

### [\#](http://loft.sh\#isolation) Isolation

**Challenge:** The third challenge is to isolate the different tenants from each other. This prevents tenants from interfering with each other. As described above, the degree of isolation is determining if you have [soft or hard multi-tenancy](http://loft.sh#soft-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) in place and if the system should only be used by trusted tenants or if it is also secured against voluntary attacks.

**Solution:** Kubernetes namespaces serve as basic isolation for tenants. Alternatively, [vClusters provide even more isolation than namespaces](https://loft.sh/features/virtual-kubernetes-clusters?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide), while they also give tenants more flexibility. Another factor that should be considered for isolating tenants is network traffic: Network policies should be configured in a way that by default all traffic is denied and only traffic within the same namespace, internet traffic for containers, and requests to the DNS are allowed. This makes it much harder for tenants to attack or interfere with each other.

## [\#](http://loft.sh\#available-multi-tenancy-solutions) Available Multi-Tenancy Solutions

Some useful solutions exist already that help you to implement multi-tenancy with Kubernetes. Besides the previously mentioned [dex](https://github.com/dexidp/dex), [kiosk](https://github.com/kiosk-sh/kiosk) and [loft](https://loft.sh/features/kubernetes-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) are worth mentioning here.

### [\#](http://loft.sh\#kiosk) kiosk

[kiosk](https://github.com/kiosk-sh/kiosk) is an open-source multi-tenancy extension for Kubernetes. It is designed as a lightweight, pluggable, and customizable solution for any Kubernetes cluster and solves some of the multi-tenancy challenges in a simple way. This includes account separation, resource consumption limitation on a user level, and namespace templates for secure tenant isolation and self-service namespace initialization. While it does not provide an in-built user management, kiosk is a very good building block to develop an own multi-tenancy system.

Even though kiosk has been released only a few months ago, it is already included in the [EKS Best Practices Guide for multi-tenancy by AWS](https://aws.github.io/aws-eks-best-practices/security/docs/multitenancy/). You can find a [detailed guide to set up kiosk with EKS here](https://aws.amazon.com/de/blogs/containers/set-up-soft-multi-tenancy-with-kiosk-on-amazon-elastic-kubernetes-service/).

### [\#](http://loft.sh\#loft) Loft

[Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) is based on kiosk and provides a comprehensive solution for a multi-tenancy platform. Loft can be installed into any Kubernetes cluster and then lets tenants create namespaces and virtual Clusters on-demand. It cares for the user management (including SSO) and user isolation and lets the cluster admins define usage limits, so all previously mentioned multi-tenancy problems are resolved. Loft provides some additional features such as a sleep mode that leads to reduced cloud computing costs by shutting down unused namespaces and vClusters.

Loft is a commercial offering with a free tier and is also included in the [EKS Best Practices Guide for multi-tenancy](https://aws.github.io/aws-eks-best-practices/multitenancy). While it was originally focused on multi-tenant development use cases, it can also be used for production use cases such as [cluster sharding](https://loft.sh/use-cases/kubernetes-cluster-sharding?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide) or to [run multiple instances of a product in a shared cluster](https://loft.sh/use-cases/cloud-native-managed-products?utm_medium=reader&utm_source=other&utm_campaign=blog_kubernetes-multi-tenancy-a-best-practices-guide)

## [\#](http://loft.sh\#conclusion) Conclusion

Kubernetes multi-tenancy is one of the challenges that many organizations face at the moment as their Kubernetes adoption progresses and single-tenant solutions become increasingly infeasible.

When implementing multi-tenancy with Kubernetes, you need to decide if you need hard multi-tenancy or if soft multi-tenancy is enough. In any case, you need to solve three major problems: How to manage the users/tenants, how to limit their resource usage, and how to isolate them from each other. Here, several tools such as dex, kiosk, and loft can help you so you get multi-tenancy with Kubernetes right more easily.

[Multi-Tenancy](http://loft.sh/blog/tags/multi-tenancy/) [Guides](http://loft.sh/blog/tags/guides/) [vcluster](http://loft.sh/blog/tags/vcluster/)

[Share on Twitter](https://twitter.com/intent/tweet?text=Kubernetes%20Multi-Tenancy%20%e2%80%93%20A%20Best%20Practices%20Guide&hashtags=kubernetes%2Cloft%2Ccontainers%2Cdevops&url=https%3a%2f%2floft.sh%2fblog%2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=twitter%26utm_content=blog-share) [Share on Linkedin](https://www.linkedin.com/sharing/share-offsite/?title=Kubernetes%20Multi-Tenancy%20%e2%80%93%20A%20Best%20Practices%20Guide&url=https%3a%2f%2floft.sh%2fblog%2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=linkedin%26utm_content=blog-share&via=loft_sh) [Share on Facebook](https://facebook.com/sharer.php?u=https%3a%2f%2floft.sh%2fblog%2fkubernetes-multi-tenancy-a-best-practices-guide%2f%3Futm_source=social%26utm_medium=fb%26utm_content=blog-share)
