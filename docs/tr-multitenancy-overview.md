# Cluster multi-tenancy

# 集群多租户

https://cloud.google.com/kubernetes-engine/docs/concepts/multitenancy-overview 

This page explains [cluster multi-tenancy](http://cloud.google.com#what_is_multi-tenancy) on Google Kubernetes Engine (GKE). This includes clusters shared by different users at a single organization, and clusters that are shared by per-customer instances of a software as a service (SaaS) application. Cluster multi-tenancy is an alternative to managing many single-tenant clusters.

本页面介绍了 Google Kubernetes Engine (GKE) 上的 [集群多租户](http://cloud.google.com#what_is_multi-tenancy)。这包括由单个组织中的不同用户共享的集群，以及由软件即服务(SaaS) 应用程序的每个客户实例共享的集群。集群多租户是管理许多单租户集群的替代方法。

This page also summarizes the Kubernetes and GKE features that can be used to manage multi-tenant clusters.

本页还总结了可用于管理多租户集群的 Kubernetes 和 GKE 功能。

## What is multi-tenancy?

## 什么是多租户？

A multi-tenant cluster is shared by multiple users and/or workloads which are referred to as "tenants". The operators of multi-tenant clusters must isolate tenants from each other to minimize the damage that a compromised or malicious tenant can do to the cluster and other tenants. Also, cluster resources must be fairly allocated among tenants.

多租户集群由称为“租户”的多个用户和/或工作负载共享。多租户集群的运营商必须将租户彼此隔离，以尽量减少受感染或恶意租户对集群和其他租户造成的损害。此外，集群资源必须在租户之间公平分配。

When you plan a multi-tenant architecture you should consider the layers of resource isolation in Kubernetes: cluster, namespace, node, Pod, and container. You should also consider the security implications of sharing different types of resources among tenants. For example, scheduling Pods from different tenants on the same node could reduce the number of machines needed in the cluster. On the other hand, you might need to prevent certain workloads from being colocated. For example, you might not allow untrusted code from outside of your organization to run on the same node as containers that process sensitive information.

当您规划多租户架构时，您应该考虑 Kubernetes 中的资源隔离层：集群、命名空间、节点、Pod 和容器。您还应该考虑在租户之间共享不同类型资源的安全隐患。例如，在同一节点上调度来自不同租户的 Pod 可以减少集群中所需的机器数量。另一方面，您可能需要防止某些工作负载被托管。例如，您可能不允许来自组织外部的不受信任的代码与处理敏感信息的容器在同一节点上运行。

Although Kubernetes cannot guarantee perfectly secure isolation between tenants, it does offer features that may be sufficient for specific use cases. You can separate each tenant and their Kubernetes resources into their own [namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/). You can then use [policies](https://kubernetes.io/docs/concepts/policy/) to enforce tenant isolation. Policies are usually scoped by namespace and can be used to restrict API access, to constrain resource usage, and to restrict what containers are allowed to do.

尽管 Kubernetes 不能保证租户之间的完全安全隔离，但它确实提供了可能足以满足特定用例的功能。您可以将每个租户及其 Kubernetes 资源分离到自己的 [命名空间](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)。然后，您可以使用 [policies](https://kubernetes.io/docs/concepts/policy/) 来强制执行租户隔离。策略通常由命名空间限定，可用于限制 API 访问、限制资源使用以及限制允许容器执行的操作。

The tenants of a multi-tenant cluster share:

多租户集群的租户共享：

- [Extensions](https://kubernetes.io/docs/concepts/extend-kubernetes/extend-cluster/), [controllers](https://kubernetes.io/docs/reference/glossary/?fundamental=true#term-controller), add-ons, and [custom resource definitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)(CRDs).
- The cluster [control plane](https://kubernetes.io/docs/concepts/architecture/control-plane-node-communication/). This implies that the cluster operations, security, and auditing are centralized.

- [扩展](https://kubernetes.io/docs/concepts/extend-kubernetes/extend-cluster/)、[控制器](https://kubernetes.io/docs/reference/glossary/?fundamental=true#term-controller)、附加组件和 [自定义资源定义](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)(CRD)。
- 集群[控制平面](https://kubernetes.io/docs/concepts/architecture/control-plane-node-communication/)。这意味着集群操作、安全和审计是集中的。

Operating a multi-tenant cluster has several advantages over operating multiple, single-tenant clusters:

与操作多个单租户集群相比，操作多租户集群有几个优势：

- Reduced management overhead
- Reduced resource fragmentation
- No need to wait for cluster creation for new tenants

- 减少管理开销
- 减少资源碎片
- 新租户无需等待创建集群

## Multi-tenancy use cases

## 多租户用例

This section describes how you could configure a cluster for various multi-tenancy use cases.

本节介绍如何为各种配置集群多租户用例。

### Enterprise multi-tenancy

### 企业多租户

In an enterprise environment, the tenants of a cluster are distinct teams within the organization. Typically, each tenant has a corresponding namespace. Alternative models of multi-tenancy with a tenant per cluster, or a tenant per Google Cloud project, are harder to manage. Network traffic within a namespace is unrestricted. Network traffic between namespaces must be explicitly allowed. These policies can be enforced using Kubernetes [network policy](http://cloud.google.com#network_policies).

在企业环境中，集群的租户是组织内不同的团队。通常，每个租户都有一个对应的命名空间。每个集群一个租户或每个 Google Cloud 项目一个租户的多租户替代模型更难管理。命名空间内的网络流量不受限制。必须明确允许命名空间之间的网络流量。可以使用 Kubernetes [网络策略](http://cloud.google.com#network_policies) 强制执行这些策略。

The users of the cluster are divided into three different roles, depending on their privilege:

集群的用户根据他们的权限分为三个不同的角色：

Cluster administrator 集群管理员

This role is for administrators of the entire cluster, who manage all tenants. Cluster administrators can create, read, update, and delete any policy object. They can create namespaces and assign them to namespace administrators.

此角色适用于管理所有租户的整个集群的管理员。集群管理员可以创建、读取、更新和删除任何策略对象。他们可以创建命名空间并将它们分配给命名空间管理员。

Namespace administrator 命名空间管理员

This role is for administrators of specific, single tenants. A namespace administrator can manage the users in their namespace.

此角色适用于特定的单一租户的管理员。命名空间管理员可以管理其命名空间中的用户。

Developer  开发者

Members of this role can create, read, update, and delete namespaced non-policy objects like [Pods](http://cloud.google.com/kubernetes-engine/docs/concepts/pod), [Jobs](http://cloud.google.com/kubernetes-engine/docs/how-to/jobs), and [Ingresses](http://cloud.google.com/kubernetes-engine/docs/concepts/ingress). Developers only have these privileges in the namespaces they have access to.

此角色的成员可以创建、读取、更新和删除命名空间的非策略对象，例如 [Pods](http://cloud.google.com/kubernetes-engine/docs/concepts/pod)、[Jobs](http://cloud.google.com/kubernetes-engine/docs/how-to/jobs) 和 [Ingresses](http://cloud.google.com/kubernetes-engine/docs/concepts/ingress)。开发商只有他们有权访问的命名空间中的这些权限。

![](http://cloud.google.com/static/kubernetes-engine/images/enterprise-multitenancy.svg)

For information on setting up multiple multi-tenant clusters for an enterprise organization, see [Best practices for enterprise multi-tenancy](http://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy) .

有关为企业组织设置多个多租户集群的信息，请参阅[企业多租户的最佳实践](http://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy) .

### SaaS provider multi-tenancy

### SaaS 提供商多租户

The tenants of a SaaS provider's cluster are the per-customer instances of the application, and the SaaS's control plane. To take advantage of namespace-scoped policies, the application instances should be organized into their own namespaces, as should components of the SaaS's control plane. End users can't interact with the Kubernetes control plane directly, they use the SaaS's interface instead, which in turn interacts with the Kubernetes control plane.

SaaS 提供商集群的租户是应用程序的每个客户实例，以及 SaaS 的控制平面。为了利用命名空间范围的策略，应用程序实例应该组织到它们自己的命名空间中，SaaS 控制平面的组件也应该如此。最终用户不能直接与 Kubernetes 控制平面交互，他们使用 SaaS 的接口，而后者又与 Kubernetes 控制平面交互。

For example, a blogging platform could run on a multi-tenant cluster. In this case, the tenants are each customer's blog instance and the platform's own control plane. The platform's control plane and each hosted blog would all run in separate namespaces. Customers would create and delete blogs, update the blogging software versions through the platform's interface with no visibility into how the cluster operates.

例如，博客平台可以在多租户集群上运行。在这种情况下，租户是每个客户的博客实例和平台自己的控制平面。该平台的控制平面和每个托管博客都将在单独的命名空间中运行。客户将创建和删除博客，通过平台界面更新博客软件版本，而无法了解集群的运行方式。

![](http://cloud.google.com/static/kubernetes-engine/images/saas-multitenancy.svg)

## Multi-tenancy policy enforcement

## 多租户策略执行

GKE and Kubernetes provide several features that can be used to manage multi-tenant clusters. The following sections give an overview of these features.

GKE 和 Kubernetes 提供了一些可用于管理多租户集群的功能。以下部分概述了这些功能。

### Access control

###  访问控制

GKE has two access control systems: Identity and Access Management (IAM) and role-based access control (RBAC). 

IAM is Google Cloud's access control system for managing authentication and authorization for GCP resources. You use IAM to grant users access to GKE and Kubernetes resources. RBAC is built into Kubernetes and grants granular permissions for specific resources and operations within your clusters.

GKE 有两个访问控制系统：身份和访问管理 (IAM) 和基于角色的访问控制 (RBAC)。
IAM 是 Google Cloud 的访问控制系统，用于管理 GCP 资源的身份验证和授权。您使用 IAM 授予用户对 GKE 和 Kubernetes 资源的访问权限。 RBAC 内置在 Kubernetes 中，并为集群中的特定资源和操作授予精细权限。

Refer to the [Access control overview](http://cloud.google.com/kubernetes-engine/docs/concepts/access-control) for more information about these options and when to use each.

请参阅[访问控制概述](http://cloud.google.com/kubernetes-engine/docs/concepts/access-control) 了解有关这些选项以及何时使用每个选项的更多信息。

Refer to the [RBAC how-to guide](http://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control) and the [IAM how-to guide](http://cloud.google.com/kubernetes-engine/docs/how-to/iam) to learn how to use these access control systems.

请参阅 [RBAC 操作指南](http://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control) 和 [IAM 操作指南](http://cloud.google.com/kubernetes-engine/docs/how-to/iam)来学习如何使用这些访问控制系统。

You can use IAM and RBAC permissions together with namespaces to restrict user interactions with cluster resources on console. For more information, see [Enable access and view cluster resources by namespace](http://cloud.google.com/kubernetes-engine/docs/how-to/restrict-resources-access-by-namespace).

您可以将 IAM 和 RBAC 权限与命名空间一起使用来限制用户在控制台上与集群资源交互。有关详细信息，请参阅 [启用按命名空间访问和查看集群资源](http://cloud.google.com/kubernetes-engine/docs/how-to/restrict-resources-access-by-namespace)。

### Network policies

### 网络策略

Cluster [network policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/) give you control over the communication between your cluster's Pods. Policies specify which namespaces, labels, and IP address ranges a Pod can communicate with.

集群 [网络策略](https://kubernetes.io/docs/concepts/services-networking/network-policies/) 让您可以控制集群 Pod 之间的通信。策略指定 Pod 可以与之通信的命名空间、标签和 IP 地址范围。

See the [network policy how-to](http://cloud.google.com/kubernetes-engine/docs/how-to/network-policy) for instructions on enabling network policy enforcement on GKE.

请参阅[网络政策操作方法 ](http://cloud.google.com/kubernetes-engine/docs/how-to/network-policy) 有关在 GKE 上启用网络策略实施的说明。

Follow the [network policy tutorial](http://cloud.google.com/kubernetes-engine/docs/tutorials/network-policy) to learn how to write network policies.

按照 [网络政策教程](http://cloud.google.com/kubernetes-engine/docs/tutorials/network-policy) 了解如何编写网络策略。

### Resource quotas

### 资源配额

Resource quotas manage the amount of resources used by the objects in a namespace. You can set quotas in terms of CPU and memory usage, or in terms of object counts. Resource quotas let you ensure that no tenant uses more than its assigned share of cluster resources.

资源配额管理命名空间中的对象使用的资源量。您可以根据 CPU 和内存使用情况或对象计数来设置配额。资源配额可让您确保没有租户使用超过其分配的集群资源份额。

Refer to the [resource quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas) documentation for more information.

有关详细信息，请参阅 [资源配额](https://kubernetes.io/docs/concepts/policy/resource-quotas) 文档。

### Pod security policies

### Pod 安全策略

**Warning:** Kubernetes has officially [deprecated PodSecurityPolicy](https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/) in version 1.21. PodSecurityPolicy will be shut down in version 1.25. For information about alternatives, refer to [PodSecurityPolicy deprecation](http://cloud.google.com/kubernetes-engine/docs/deprecations/podsecuritypolicy). 

**Note:** For Autopilot clusters, [Pod security policies](http://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#pod_security_policies) are not supported.

**警告：** Kubernetes 已正式 [deprecated PodSecurityPolicy](https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/) 在 1.21 版中。 PodSecurityPolicy 将在 1.25 版本中关闭。有关替代方案的信息，请参阅 [PodSecurityPolicy deprecation](http://cloud.google.com/kubernetes-engine/docs/deprecations/podsecuritypolicy)。**注意：** 对于 Autopilot 集群，[Pod 安全策略](http://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#pod_security_policies) 不支持。

[PodSecurityPolicies ](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) are a Kubernetes API type that validate requests to create and update Pods.
PodSecurityPolicies define default values and requirements for security-sensitive fields of Pod specification. You can create policies that restrict the deployment of Pods that access the host filesystem, networks, PID
namespaces, [volumes](http://cloud.google.com/kubernetes-engine/docs/concepts/volumes), and more.

[PodSecurityPolicies](https://kubernetes.io/docs/concepts/policy/pod-security-policy/)
是一种 Kubernetes API 类型，用于验证创建和更新 Pod 的请求。PodSecurityPolicies 定义默认值和要求
Pod 规范的安全敏感字段。您可以创建策略限制访问主机文件系统、网络、PID 的 Pod 的部署命名空间、[volumes](http://cloud.google.com/kubernetes-engine/docs/concepts/volumes) 等等。

Refer to the [PodSecurityPolicies how-to](http://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) for more.

有关更多信息，请参阅 [PodSecurityPolicies how-to](http://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies)。

### Pod anti-affinity

### Pod 反亲和性

**Warning:** Pod anti-affinity rules can be circumvented by malicious tenants. The example below should only be used with clusters with trusted tenants, or with tenants who don't have direct access to the Kubernetes control plane.

**警告：** Pod 反关联规则可能会被恶意租户规避。这下面的示例应仅用于具有受信任租户的集群，或无法直接访问 Kubernetes 控制平面的租户。

You can use [Pod anti-affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity) to prevent Pods from different tenants from being scheduled on the same node. Anti-affinity constraints are based on Pod [labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/). For example, the Pod specification below describes a Pod with the label `"team":"billing"`, and an anti-affinity rule that prevents the Pod from being scheduled alongside Pods without the label.

您可以使用 [Pod 反亲和性](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity)防止来自不同租户的 Pod 被调度在同一个节点上。 反亲和约束基于 Pod [标签](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)。 
例如，下面的 Pod 规范描述了一个带有 `"team":"billing"` 标签的 Pod，以及防止 Pod 被调度在没有标签的 Pod 旁边。

```
apiVersion: v1
kind: Pod
metadata:
name: bar
labels:
    team: "billing"
spec:
affinity:
    podAntiAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
      - topologyKey: "kubernetes.io/hostname"
        labelSelector:
          matchExpressions:
          - key: "team"
            operator: NotIn
            values: ["billing"]

```

The drawback to this technique is that malicious users could circumvent the rule by applying the `team: billing` label to an arbitrary Pod. Pod anti-affinity alone cannot securely enforce policy on clusters with untrusted tenants. Refer to the [Pod anti-affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity) documentation for more information.

这种技术的缺点是恶意用户可以通过将“团队：计费”标签应用于任意 Pod 来规避规则。仅 Pod 反亲和性无法安全地对具有不受信任的租户的集群执行策略。更多信息请参考 [Pod 反亲和性](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity) 文档。

### Dedicated nodes with taints and tolerations

### 带有污点和容忍度的专用节点

**Warning:** Policies enforced by node taints and tolerations can be circumvented by malicious tenants. The example below should only be used with clusters with trusted tenants, or with tenants who don't have direct access to the Kubernetes control plane.

**警告：** 由节点污染和容忍强制执行的策略可能会被恶意租户规避。下面的示例只能用于具有受信任租户的集群，或者不能直接访问 Kubernetes 控制平面的租户。

Node taints are another way to control workload scheduling. You can use node taints to reserve specialized nodes for use by certain tenants. For example, you can dedicate [GPU equipped nodes](http://cloud.google.com/kubernetes-engine/docs/how-to/gpus) to the specific tenants whose workloads require GPUs. For Autopilot clusters, node tolerations are supported only for [workload separation](http://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning#workload_separation).

节点污点是控制工作负载调度的另一种方式。您可以使用节点污点来保留专用节点以供某些租户使用。例如，您可以将 [配备 GPU 的节点](http://cloud.google.com/kubernetes-engine/docs/how-to/gpus) 专用于工作负载需要 GPU 的特定租户。对于 Autopilot 集群，仅 [工作负载分离](http://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning#workload_separation) 支持节点容差。

Node taints are automatically added by node auto-provisioning as needed.

节点污点由节点自动配置根据需要自动添加。

To dedicate a [node pool](http://cloud.google.com/kubernetes-engine/docs/concepts/node-pools) to a certain tenant, apply a taint with `effect: "NoSchedule"` to the node pool . Then only Pods with a corresponding toleration can be scheduled to nodes in the node pool.

要将 [节点池](http://cloud.google.com/kubernetes-engine/docs/concepts/node-pools) 专用于某个租户，请对节点池应用具有 `effect: "NoSchedule"` 的污点。那么只有具有相应容忍度的 Pod 才能被调度到节点池中的节点上。

The drawback to this technique is that malicious users could add a toleration to their Pods to get access to the dedicated node pool. Node taints and tolerations alone cannot securely enforce policy on clusters with untrusted tenants.

这种技术的缺点是恶意用户可以向他们的 Pod 添加一个容忍度来访问专用节点池。仅靠节点污染和容忍度无法安全地对具有不受信任租户的集群执行策略。

See the [node taints how-to page](http://cloud.google.com/kubernetes-engine/docs/how-to/node-taints) to learn how to control scheduling with node taints.

请参阅 [node taints how-to page ](http://cloud.google.com/kubernetes-engine/docs/how-to/node-taints) 了解如何使用节点污点控制调度。

## What's next

##  下一步是什么

- Watch the[Kubernetes Multi-tenancy talk](https://www.youtube.com/watch?v=RkY8u1_f5yY) from Google Cloud Next '18.
- Read the[Best practices for enterprise multi-tenancy](http://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy). 

- 观看来自 Google Cloud Next '18 的 [Kubernetes 多租户演讲](https://www.youtube.com/watch?v=RkY8u1_f5yY)。
- 阅读[企业多租户的最佳实践](http://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy)。

- Learn how to [Optimize resource usage in a multi-tenant GKE cluster using node auto-provisioning](http://cloud.google.com/solutions/optimizing-resources-in-multi-tenant-gke-clusters-with-auto-provisioning).

- 了解如何[使用节点自动配置优化多租户 GKE 集群中的资源使用](http://cloud.google.com/solutions/optimizing-resources-in-multi-tenant-gke-clusters-with-自动配置)。
