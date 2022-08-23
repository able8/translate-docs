# Three Tenancy Models For Kubernetes

# Kubernetes 的三种租户模型

Thursday, April 15, 2021
https://kubernetes.io/blog/2021/04/15/three-tenancy-models-for-kubernetes

2021 年 4 月 15 日星期四
https://kubernetes.io/blog/2021/04/15/three-tenancy-models-for-kubernetes

**Authors:** Ryan Bezdicek (Medtronic), Jim Bugwadia (Nirmata), Tasha Drew (VMware), Fei Guo (Alibaba), Adrian Ludwin (Google)

**作者：** Ryan Bezdicek (Medtronic)、Jim Bugwadia (Nirmata)、Tasha Drew (VMware)、Fei Guo (阿里巴巴)、Adrian Ludwin (Google)

Kubernetes clusters are typically used by several teams in an organization. In other cases, Kubernetes may be used to deliver applications to end users requiring segmentation and isolation of resources across users from different organizations. Secure sharing of Kubernetes control plane and worker node resources allows maximizing productivity and saving costs in both cases.

Kubernetes 集群通常由组织中的多个团队使用。在其他情况下，Kubernetes 可用于将应用程序交付给需要在来自不同组织的用户之间分割和隔离资源的最终用户。 Kubernetes 控制平面和工作节点资源的安全共享允许在这两种情况下最大限度地提高生产力并节省成本。

The Kubernetes Multi-Tenancy Working Group is chartered with defining tenancy models for Kubernetes and making it easier to operationalize tenancy related use cases. This blog post, from the working group members, describes three common tenancy models and introduces related working group projects.

Kubernetes 多租户工作组的职责是为 Kubernetes 定义租户模型，并使其更容易操作与租户相关的用例。这篇来自工作组成员的博文描述了三种常见的租户模型，并介绍了相关的工作组项目。

We will also be presenting on this content and discussing different use cases at our Kubecon EU 2021 panel session, [Multi-tenancy vs. Multi-cluster: When Should you Use What?](https://sched.co/iE66).

我们还将在我们的 Kubecon EU 2021 小组会议上介绍此内容并讨论不同的用例，[多租户与多集群：您何时应该使用什么？](https://sched.co/iE66)。

## Namespaces as a Service

## 命名空间即服务

With the _namespaces-as-a-service_ model, tenants share a cluster and tenant workloads are restricted to a set of Namespaces assigned to the tenant. The cluster control plane resources like the API server and scheduler, and worker node resources like CPU, memory, etc. are available for use across all tenants.

使用 _namespaces-as-a-service_ 模型，租户共享一个集群，租户工作负载被限制在一组分配给租户的命名空间中。 API 服务器和调度程序等集群控制平面资源以及 CPU、内存等工作节点资源可供所有租户使用。

To isolate tenant workloads, each namespace must also contain:

为了隔离租户工作负载，每个命名空间还必须包含：

- **[role bindings](http://kubernetes.io/docs/reference/access-authn-authz/rbac/#rolebinding-and-clusterrolebinding):** for controlling access to the namespace
- **[network policies](http://kubernetes.io/docs/concepts/services-networking/network-policies/):** to prevent network traffic across tenants
- **[resource quotas](http://kubernetes.io/docs/concepts/policy/resource-quotas/):** to limit usage and ensure fairness across tenants

- **[角色绑定](http://kubernetes.io/docs/reference/access-authn-authz/rbac/#rolebinding-and-clusterrolebinding):** 用于控制对命名空间的访问
- **[网络策略](http://kubernetes.io/docs/concepts/services-networking/network-policies/):** 防止跨租户的网络流量
- **[资源配额](http://kubernetes.io/docs/concepts/policy/resource-quotas/):** 限制使用并确保跨租户的公平性

With this model, tenants share cluster-wide resources like ClusterRoles and CustomResourceDefinitions (CRDs) and hence cannot create or update these cluster-wide resources.

使用此模型，租户共享集群范围的资源，如 ClusterRoles 和 CustomResourceDefinitions (CRD)，因此无法创建或更新这些集群范围的资源。

The [Hierarchical Namespace Controller (HNC)](http://kubernetes.io/blog/2020/08/14/introducing-hierarchical-namespaces/) project makes it easier to manage namespace based tenancy by allowing users to create additional namespaces under a namespace, and propagating resources within the namespace hierarchy. This allows self-service namespaces for tenants, without requiring cluster-wide permissions.

[Hierarchical Namespace Controller (HNC)](http://kubernetes.io/blog/2020/08/14/introducing-hierarchical-namespaces/) 项目通过允许用户在命名空间，并在命名空间层次结构中传播资源。这允许租户使用自助服务命名空间，而无需集群范围的权限。

The [Multi-Tenancy Benchmarks (MTB)](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks) project provides benchmarks and a command-line tool that performs several configuration and runtime checks to report if tenant namespaces are properly isolated and the necessary security controls are implemented.

[Multi-Tenancy Benchmarks (MTB)](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks) 项目提供了基准测试和一个命令行工具，用于执行多个配置和运行时检查以报告租户命名空间是否正确隔离并实施了必要的安全控制。

## Clusters as a Service

## 集群即服务

With the _clusters-as-a-service_ usage model, each tenant gets their own cluster. This model allows tenants to have different versions of cluster-wide resources such as CRDs, and provides full isolation of the Kubernetes control plane.

使用 _clusters-as-a-service_ 使用模型，每个租户都有自己的集群。该模型允许租户拥有 CRD 等不同版本的集群范围资源，并提供 Kubernetes 控制平面的完全隔离。

The tenant clusters may be provisioned using projects like [Cluster API (CAPI)](https://cluster-api.sigs.k8s.io/) where a management cluster is used to provision multiple workload clusters. A workload cluster is assigned to a tenant and tenants have full control over cluster resources. Note that in most enterprises a central platform team may be responsible for managing required add-on services such as security and monitoring services, and for providing cluster lifecycle management services such as patching and upgrades. A tenant administrator may be restricted from modifying the centrally managed services and other critical cluster information.

租户集群可以使用 [Cluster API (CAPI)](https://cluster-api.sigs.k8s.io/) 等项目来配置，其中管理集群用于配置多个工作负载集群。工作负载集群分配给租户，租户可以完全控制集群资源。请注意，在大多数企业中，中央平台团队可能负责管理所需的附加服务，例如安全和监控服务，并提供集群生命周期管理服务，例如补丁和升级。租户管理员可能被限制修改集中管理的服务和其他关键集群信息。

## Control planes as a Service 

## 控制平面即服务

In a variation of the _clusters-as-a-service_ model, the tenant cluster may be a **virtual cluster** where each tenant gets their own dedicated Kubernetes control plane but share worker node resources. As with other forms of virtualization, users of a virtual cluster see no significant differences between a virtual cluster and other Kubernetes clusters. This is sometimes referred to as `Control Planes as a Service` (CPaaS).

在 _clusters-as-a-service_ 模型的变体中，租户集群可能是一个**虚拟集群**，其中每个租户都有自己的专用 Kubernetes 控制平面，但共享工作节点资源。与其他形式的虚拟化一样，虚拟集群的用户认为虚拟集群与其他 Kubernetes 集群之间没有显着差异。这有时被称为“控制平面即服务”（CPaaS）。

A virtual cluster of this type shares worker node resources and workload state independent control plane components, like the scheduler. Other workload aware control-plane components, like the API server, are created on a per-tenant basis to allow overlaps, and additional components are used to synchronize and manage state across the per-tenant control plane and the underlying shared cluster resources. With this model users can manage their own cluster-wide resources.

这种类型的虚拟集群共享工作节点资源和工作负载状态独立的控制平面组件，如调度程序。其他工作负载感知控制平面组件（如 API 服务器）是在每个租户的基础上创建的以允许重叠，并且其他组件用于跨每个租户控制平面和底层共享集群资源同步和管理状态。使用此模型，用户可以管理他们自己的集群范围内的资源。

The [Virtual Cluster](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/virtualcluster) project implements this model, where a `supercluster` is shared by multiple `virtual clusters`. The [Cluster API Nested](https://github.com/kubernetes-sigs/cluster-api-provider-nested) project is extending this work to conform to the CAPI model, allowing use of familiar API resources to provision and manage virtual clusters.

[虚拟集群](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/virtualcluster) 项目实现了这个模型，其中一个“超级集群”由多个“虚拟集群”共享。 [集群 API 嵌套](https://github.com/kubernetes-sigs/cluster-api-provider-nested) 项目正在扩展这项工作以符合 CAPI 模型，允许使用熟悉的 API 资源来配置和管理虚拟集群。

## Security considerations

## 安全注意事项

Cloud native security involves different system layers and lifecycle phases as described in the [Cloud Native Security Whitepaper](http://kubernetes.io/blog/2020/11/18/cloud-native-security-for-your-clusters) from CNCF SIG Security. Without proper security measures implemented across all layers and phases, Kubernetes tenant isolation can be compromised and a security breach with one tenant can threaten other tenants.

云原生安全涉及不同的系统层和生命周期阶段，如 [云原生安全白皮书](http://kubernetes.io/blog/2020/11/18/cloud-native-security-for-your-clusters) 中所述CNCF SIG 安全。如果没有在所有层和阶段实施适当的安全措施，Kubernetes 租户隔离可能会受到影响，一个租户的安全漏洞可能会威胁到其他租户。

It is important for any new user to Kubernetes to realize that the default installation of a new upstream Kubernetes cluster is not secure, and you are going to need to invest in hardening it in order to avoid security issues.

对于 Kubernetes 的任何新用户来说，重要的是要意识到新的上游 Kubernetes 集群的默认安装是不安全的，您将需要投资加固它以避免安全问题。

At a minimum, the following security measures are required:

至少需要以下安全措施：

- image scanning: container image vulnerabilities can be exploited to execute commands and access additional resources.
- [RBAC](http://kubernetes.io/docs/reference/access-authn-authz/rbac/): for _namespaces-as-a-service_ user roles and permissions must be properly configured at a per-namespace level; for other models tenants may need to be restricted from accessing centrally managed add-on services and other cluster-wide resources.
- [network policies](http://kubernetes.io/docs/concepts/services-networking/network-policies/): for _namespaces-as-a-service_ default network policies that deny all ingress and egress traffic are recommended to prevent cross-tenant network traffic and may also be used as a best practice for other tenancy models.
- [Kubernetes Pod Security Standards](http://kubernetes.io/docs/concepts/security/pod-security-standards/): to enforce Pod hardening best practices the `Restricted` policy is recommended as the default for tenant workloads with exclusions configured only as needed.
- [CIS Benchmarks for Kubernetes](https://www.cisecurity.org/benchmark/kubernetes/): the CIS Benchmarks for Kubernetes guidelines should be used to properly configure Kubernetes control-plane and worker node components.

- 镜像扫描：容器镜像漏洞可以被利用来执行命令和访问额外的资源。
- [RBAC](http://kubernetes.io/docs/reference/access-authn-authz/rbac/)：对于_namespaces-as-a-service_，必须在每个命名空间级别正确配置用户角色和权限；对于其他模型，可能需要限制租户访问集中管理的附加服务和其他集群范围的资源。
- [网络策略](http://kubernetes.io/docs/concepts/services-networking/network-policies/)：对于_namespaces-as-a-service_默认网络策略，建议拒绝所有入口和出口流量，以防止跨租户网络流量，也可用作其他租户模型的最佳实践。
- [Kubernetes Pod 安全标准](http://kubernetes.io/docs/concepts/security/pod-security-standards/)：为了执行 Pod 强化最佳实践，建议将“受限”策略作为租户工作负载的默认值仅根据需要配置排除项。
- [Kubernetes 的 CIS 基准](https://www.cisecurity.org/benchmark/kubernetes/)：应使用 Kubernetes 的 CIS 基准指南正确配置 Kubernetes 控制平面和工作节点组件。

Additional recommendations include using:

其他建议包括使用：

- policy engines: for configuration security best practices, such as only allowing trusted registries.
- runtime scanners: to detect and report runtime security events.
- VM-based container sandboxing: for stronger data plane isolation.

- 策略引擎：用于配置安全最佳实践，例如仅允许受信任的注册表。
- 运行时扫描器：检测和报告运行时安全事件。
- 基于虚拟机的容器沙盒：用于更强的数据平面隔离。

While proper security is required independently of tenancy models, not having essential security controls like [pod security](http://kubernetes.io/docs/concepts/security/pod-security-standards/) in a shared cluster provides attackers with means to compromise tenancy models and possibly access sensitive information across tenants increasing the overall risk profile.

虽然独立于租赁模型需要适当的安全性，但在共享集群中没有像 [pod security](http://kubernetes.io/docs/concepts/security/pod-security-standards/) 这样的基本安全控制为攻击者提供了手段破坏租赁模式并可能访问跨租户的敏感信息，从而增加整体风险状况。

## Summary

##  概括

A 2020 CNCF survey showed that production Kubernetes usage has increased by over 300% since 2016. As an increasing number of Kubernetes workloads move to production, organizations are looking for ways to share Kubernetes resources across teams for agility and cost savings. 

2020 年 CNCF 调查显示，自 2016 年以来，生产 Kubernetes 的使用量增加了 300% 以上。随着越来越多的 Kubernetes 工作负载转向生产，组织正在寻找跨团队共享 Kubernetes 资源的方法，以实现敏捷性和成本节约。

The **namespaces as a service** tenancy model allows sharing clusters and hence enables resource efficiencies. However, it requires proper security configurations and has limitations as all tenants share the same cluster-wide resources.

**命名空间即服务**租户模型允许共享集群，从而提高资源效率。但是，它需要适当的安全配置并且有局限性，因为所有租户共享相同的集群范围资源。

The **clusters as a service** tenancy model addresses these limitations, but with higher management and resource overhead.

**集群即服务**租赁模型解决了这些限制，但管理和资源开销更高。

The **control planes as a service** model provides a way to share resources of a single Kubernetes cluster and also let tenants manage their own cluster-wide resources. Sharing worker node resources increases resource effeciencies, but also exposes cross tenant security and isolation concerns that exist for shared clusters.

**控制平面即服务**模型提供了一种共享单个 Kubernetes 集群资源的方法，还允许租户管理他们自己的集群范围内的资源。共享工作节点资源提高了资源效率，但也暴露了共享集群存在的跨租户安全和隔离问题。

In many cases, organizations will use multiple tenancy models to address different use cases and as different product and development teams will have varying needs. Following security and management best practices, such as applying [Pod Security Standards](http://kubernetes.io/docs/concepts/security/pod-security-standards/) and not using the `default` namespace, makes it easer to switch from one model to another.

在许多情况下，组织将使用多种租赁模型来解决不同的用例，因为不同的产品和开发团队会有不同的需求。遵循安全和管理最佳实践，例如应用 [Pod 安全标准](http://kubernetes.io/docs/concepts/security/pod-security-standards/) 并且不使用 `default` 命名空间，可以更轻松地从一种模型切换到另一种模型。

The [Kubernetes Multi-Tenancy Working Group](https://github.com/kubernetes-sigs/multi-tenancy) has created several projects like [Hierarchical Namespaces Controller](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/hnc), [Virtual Cluster](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/virtualcluster) / [CAPI Nested](https://github.com/kubernetes-sigs/cluster-api-provider-nested), and [Multi-Tenancy Benchmarks](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks) to make it easier to provision and manage multi-tenancy models.

[Kubernetes 多租户工作组](https://github.com/kubernetes-sigs/multi-tenancy) 已经创建了多个项目，例如 [Hierarchical Namespaces Controller](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/hnc), [虚拟集群](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/virtualcluster) / [CAPI嵌套](https:///github.com/kubernetes-sigs/cluster-api-provider-nested) 和 [Multi-Tenancy Benchmarks](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks) 来制作更容易配置和管理多租户模型。

If you are interested in multi-tenancy topics, or would like to share your use cases, please join us in an upcoming [community meeting](https://github.com/kubernetes/community/blob/master/wg-multitenancy/README.md) or reach out on the _wg-multitenancy channel_ on the [Kubernetes slack](https://slack.k8s.io/). 

如果您对多租户主题感兴趣，或者想分享您的用例，请加入我们即将举行的 [社区会议](https://github.com/kubernetes/community/blob/master/wg-multitenancy/README.md) 或联系 [Kubernetes slack] (https://slack.k8s.io/) 上的 _wg-multitenancy 频道_。

