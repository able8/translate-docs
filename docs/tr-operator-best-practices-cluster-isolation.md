# Best practices for cluster isolation in Azure Kubernetes Service (AKS)

# Azure Kubernetes 服务 (AKS) 中集群隔离的最佳实践

- 08/22/2022  https://docs.microsoft.com/en-us/azure/aks/operator-best-practices-cluster-isolation

### In this article

###  在本文中

As you manage clusters in Azure Kubernetes Service (AKS), you often need to isolate teams and workloads. AKS provides flexibility in how you can run multi-tenant clusters and isolate resources. To maximize your investment in Kubernetes, first understand and implement AKS multi-tenancy and isolation features.

在 Azure Kubernetes 服务 (AKS) 中管理群集时，通常需要隔离团队和工作负载。 AKS 在如何运行多租户群集和隔离资源方面提供了灵活性。为了最大限度地提高您对 Kubernetes 的投资，首先要了解并实施 AKS 多租户和隔离功能。

This best practices article focuses on isolation for cluster operators. In this article, you learn how to:

这篇最佳实践文章侧重于集群操作员的隔离。在本文中，您将学习如何：

- Plan for multi-tenant clusters and separation of resources
- Use logical or physical isolation in your AKS clusters

- 规划多租户集群和资源分离
- 在 AKS 群集中使用逻辑或物理隔离

## Design clusters for multi-tenancy

## 为多租户设计集群

Kubernetes lets you logically isolate teams and workloads in the same cluster. The goal is to provide the least number of privileges, scoped to the resources each team needs. A Kubernetes [Namespace](http://docs.microsoft.com/concepts-clusters-workloads#namespaces) creates a logical isolation boundary. Additional Kubernetes features and considerations for isolation and multi-tenancy include the following areas:

Kubernetes 允许您在逻辑上隔离同一集群中的团队和工作负载。目标是提供最少数量的特权，范围仅限于每个团队需要的资源。 Kubernetes [命名空间](http://docs.microsoft.com/concepts-clusters-workloads#namespaces) 创建了一个逻辑隔离边界。隔离和多租户的其他 Kubernetes 功能和注意事项包括以下方面：

### Scheduling

### 调度

_Scheduling_ uses basic features such as resource quotas and pod disruption budgets. For more information about these features, see [Best practices for basic scheduler features in AKS](http://docs.microsoft.com/operator-best-practices-scheduler).

_Scheduling_ 使用资源配额和 Pod 中断预算等基本功能。有关这些功能的详细信息，请参阅 [AKS 中基本调度程序功能的最佳做法](http://docs.microsoft.com/operator-best-practices-scheduler)。

More advanced scheduler features include:

更高级的调度程序功能包括：

- Taints and tolerations
- Node selectors
- Node and pod affinity or anti-affinity.

- 污点和容忍度
- 节点选择器
- 节点和 pod 亲和性或反亲和性。

For more information about these features, see [Best practices for advanced scheduler features in AKS](http://docs.microsoft.com/operator-best-practices-advanced-scheduler).

有关这些功能的详细信息，请参阅 [AKS 中高级调度程序功能的最佳做法](http://docs.microsoft.com/operator-best-practices-advanced-scheduler)。

### Networking

###  联网

_Networking_ uses network policies to control the flow of traffic in and out of pods.

_Networking_ 使用网络策略来控制进出 Pod 的流量。

### Authentication and authorization

### 身份验证和授权

_Authentication and authorization_ uses:

_身份验证和授权_使用：

- Role-based access control (RBAC)
- Azure Active Directory (AD) integration
- Pod identities
- Secrets in Azure Key Vault

- 基于角色的访问控制 (RBAC)
- Azure Active Directory (AD) 集成
- 豆荚身份
- Azure Key Vault 中的秘密

For more information about these features, see [Best practices for authentication and authorization in AKS](http://docs.microsoft.com/operator-best-practices-identity).

有关这些功能的详细信息，请参阅 [AKS 中身份验证和授权的最佳做法](http://docs.microsoft.com/operator-best-practices-identity)。

### Containers

### 容器

_Containers_ include:

_容器_包括：

- The Azure Policy Add-on for AKS to enforce pod security.
- The use of pod security admission.
- Scanning both images and the runtime for vulnerabilities.
- Using App Armor or Seccomp (Secure Computing) to restrict container access to the underlying node.

- 用于 AKS 的 Azure Policy 加载项以强制执行 pod 安全性。
- 使用吊舱安全入场。
- 扫描图像和运行时的漏洞。
- 使用 App Armor 或 Seccomp（安全计算）来限制容器对底层节点的访问。

## Logically isolate clusters

## 逻辑隔离集群

> **Best practice guidance**
>
> Separate teams and projects using _logical isolation_. Minimize the number of physical AKS clusters you deploy to isolate teams or applications.

> **最佳实践指南**
>
> 使用_逻辑隔离_分隔团队和项目。尽量减少为隔离团队或应用程序而部署的物理 AKS 群集的数量。

With logical isolation, a single AKS cluster can be used for multiple workloads, teams, or environments. Kubernetes [Namespaces](http://docs.microsoft.com/concepts-clusters-workloads#namespaces) form the logical isolation boundary for workloads and resources.

通过逻辑隔离，单个 AKS 群集可用于多个工作负载、团队或环境。 Kubernetes [命名空间](http://docs.microsoft.com/concepts-clusters-workloads#namespaces) 形成了工作负载和资源的逻辑隔离边界。

![Logical isolation of a Kubernetes cluster in AKS](https://docs.microsoft.com/en-us/azure/aks/media/operator-best-practices-cluster-isolation/logical-isolation.png)

Logical separation of clusters usually provides a higher pod density than physically isolated clusters, with less excess compute capacity sitting idle in the cluster. When combined with the Kubernetes cluster autoscaler, you can scale the number of nodes up or down to meet demands. This best practice approach to autoscaling minimizes costs by running only the number of nodes required.

集群的逻辑分离通常提供比物理隔离集群更高的 pod 密度，并且集群中闲置的多余计算容量更少。与 Kubernetes 集群自动扩展程序结合使用时，您可以扩展或缩减节点数量以满足需求。这种自动缩放的最佳实践方法通过仅运行所需数量的节点来最大程度地降低成本。

Currently, Kubernetes environments aren't completely safe for hostile multi-tenant usage. In a multi-tenant environment, multiple tenants are working on a common, shared infrastructure. If all tenants cannot be trusted, you will need extra planning to prevent tenants from impacting the security and service of others.

目前，Kubernetes 环境对于恶意的多租户使用并不完全安全。在多租户环境中，多个租户在一个通用的共享基础架构上工作。如果不能信任所有租户，您将需要进行额外规划，以防止租户影响他人的安全和服务。

Additional security features, like Kubernetes RBAC for nodes, efficiently block exploits. For true security when running hostile multi-tenant workloads, you should only trust a hypervisor. The security domain for Kubernetes becomes the entire cluster, not an individual node.

其他安全功能，例如用于节点的 Kubernetes RBAC，可以有效地阻止漏洞利用。为了在运行恶意多租户工作负载时实现真正的安全性，您应该只信任虚拟机管理程序。 Kubernetes 的安全域成为整个集群，而不是单个节点。

For these types of hostile multi-tenant workloads, you should use physically isolated clusters.

对于这些类型的恶意多租户工作负载，您应该使用物理隔离的集群。

## Physically isolate clusters

## 物理隔离集群

> **Best practice guidance**
> 

> **最佳实践指南**
>

> Minimize the use of physical isolation for each separate team or application deployment. Instead, use _logical_ isolation, as discussed in the previous section.

> 尽量减少对每个单独的团队或应用程序部署使用物理隔离。相反，请使用 _logical_ 隔离，如上一节所述。

Physically separating AKS clusters is a common approach to cluster isolation. In this isolation model, teams or workloads are assigned their own AKS cluster. While physical isolation might look like the easiest way to isolate workloads or teams, it adds management and financial overhead. Now, you must maintain these multiple clusters and individually provide access and assign permissions. You'll also be billed for each the individual node.

物理分离 AKS 群集是群集隔离的常用方法。在此隔离模型中，为团队或工作负载分配了自己的 AKS 群集。虽然物理隔离可能看起来是隔离工作负载或团队的最简单方法，但它会增加管理和财务开销。现在，您必须维护这些多个集群并单独提供访问和分配权限。您还将为每个单独的节点付费。

![Physical isolation of individual Kubernetes clusters in AKS](https://docs.microsoft.com/en-us/azure/aks/media/operator-best-practices-cluster-isolation/physical-isolation.png)

Physically separate clusters usually have a low pod density. Since each team or workload has their own AKS cluster, the cluster is often over-provisioned with compute resources. Often, a small number of pods are scheduled on those nodes. Unclaimed node capacity can't be used for applications or services in development by other teams. These excess resources contribute to the additional costs in physically separate clusters. 

物理上独立的集群通常具有较低的 pod 密度。由于每个团队或工作负载都有自己的 AKS 群集，因此群集通常会过度配置计算资源。通常，在这些节点上会安排少量的 pod。无人认领的节点容量不能用于其他团队正在开发的应用程序或服务。这些多余的资源会导致物理上分离的集群的额外成本。

