# Best practices for enterprise multi-tenancy

# 企业多租户的最佳实践

From: https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy

Multi-tenancy in Google Kubernetes Engine (GKE) refers to one or more clusters that are shared between tenants. In Kubernetes, a *tenant* can be defined as any of the following:

Google Kubernetes Engine (GKE) 中的多租户是指租户之间共享的一个或多个集群。在 Kubernetes 中，*tenant* 可以定义为以下任何一种：

- A team responsible for developing and operating one or more workloads.
- A set of related workloads, whether operated by one or more teams.
- A single workload, such as a Deployment.

- 负责开发和操作一个或多个工作负载的团队。
- 一组相关的工作负载，无论是由一个或多个团队操作。
- 单个工作负载，例如部署。

[Cluster multi-tenancy](https://cloud.google.com/kubernetes-engine/docs/concepts/multitenancy-overview#what_is_multi-tenancy) is often implemented to reduce costs or to consistently apply administration policies across tenants. However, incorrectly configuring a GKE cluster or its associated GKE resources can result in unachieved cost savings, incorrect policy application, or destructive interactions between different tenants' workloads.

[集群多租户](https://cloud.google.com/kubernetes-engine/docs/concepts/multitenancy-overview#what_is_multi-tenancy) 通常用于降低成本或跨租户一致地应用管理策略。但是，错误地配置 GKE 集群或其关联的 GKE 资源可能会导致无法节省成本、错误的策略应用或不同租户工作负载之间的破坏性交互。

This guide provides best practices to safely and efficiently set up multiple multi-tenant clusters for an enterprise organization.

本指南提供了为企业组织安全有效地设置多个多租户集群的最佳实践。

**Note:** For a summarized checklist of all the best practices, see the  [Checklist summary](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#checklist) at the bottom of this guide.

**注意：** 有关所有最佳实践的汇总清单，请参阅 [清单摘要](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#checklist)，网址为本指南的底部。

## Assumptions and requirements

## 假设和要求

The best practices in this guide are based on a multi-tenant use case for an enterprise environment, which has the following assumptions and requirements:

本指南中的最佳实践基于企业环境的多租户用例，具有以下假设和要求：

- The organization is a single company that has many tenants (two or more application/service teams) that use Kubernetes and would like to share computing and administrative resources.
- Each tenant is a single team developing a single workload.
- Other than the application/service teams, there are other teams that also utilize and manage clusters, including platform team members, cluster administrators, auditors, etc.
- The platform team owns the clusters and defines the amount of resources each tenant team can use; each tenant can request more.
- Each tenant team should be able to deploy their application through the Kubernetes API without having to communicate with the platform team.
- Each tenant should not be able to affect other tenants in the shared cluster, except via explicit design decisions like API calls, shared data sources, etc.

- 该组织是一家拥有许多使用 Kubernetes 并希望共享计算和管理资源的租户（两个或多个应用程序/服务团队）的单一公司。
- 每个租户都是一个开发单一工作负载的团队。
- 除了应用/服务团队之外，还有其他团队也在使用和管理集群，包括平台团队成员、集群管理员、审计员等。
- 平台团队拥有集群并定义每个租户团队可以使用的资源量；每个租户都可以要求更多。
- 每个租户团队都应该能够通过 Kubernetes API 部署他们的应用程序，而无需与平台团队沟通。
- 每个租户不应影响共享集群中的其他租户，除非通过 API 调用、共享数据源等明确的设计决策。

This setup will serve as a model from which we can demonstrate multi-tenant best practices. While this setup might not perfectly describe all enterprise organizations, it can be easily extended to cover similar scenarios.

这个设置将作为一个模型，我们可以从中展示多租户的最佳实践。虽然此设置可能无法完美地描述所有企业组织，但它可以轻松扩展以涵盖类似的场景。

**Note:** For Terraform modules and sample deployments, see the  [GoogleCloudPlatform/gke-enterprise-mt](https://github.com/GoogleCloudPlatform/gke-enterprise-mt)  GitHub repository.

**注意：** 对于 Terraform 模块和示例部署，请参阅 [GoogleCloudPlatform/gke-enterprise-mt](https://github.com/GoogleCloudPlatform/gke-enterprise-mt) GitHub 存储库。

## Setting up folders, projects and clusters

## 设置文件夹、项目和集群

**Best practices**:

**最佳实践**：

- [Establish a folder and project hierarchy](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#folder-hierarchy)
- [Assign roles using IAM](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#assign-iam-roles).
- [Centralize network control with Shared VPCs](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#network-control).
- [Create one cluster admin project per cluster](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-cluster).
- [Make clusters private](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-cluster).
- [Ensure the control plane for the cluster is regional](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-cluster).
- [Ensure nodes in your cluster span at least three zones](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-cluster).
- [Autoscale cluster nodes and resources](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#autoscale-cluster).
- [Schedule maintenance windows for off-peak hours](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#maintenance-window).
- [Set up HTTP(S) Load Balancing with Ingress](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#load-balancing).

- [建立文件夹和项目层次结构](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#folder-hierarchy)
- [使用 IAM 分配角色](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#assign-iam-roles)。
- [使用共享 VPC 集中网络控制](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#network-control)。
- [为每个集群创建一个集群管理项目](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-cluster)。
- [将集群设为私有](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-cluster)。
- [确保集群的控制平面是区域性的](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-cluster)。
- [确保集群中的节点至少跨越三个区域](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-cluster)。
- [自动缩放集群节点和资源](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#autoscale-cluster)。
- [安排非高峰时段的维护时段](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#maintenance-window)。
- [使用 Ingress 设置 HTTP(S) 负载平衡](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#load-balancing)。

### Establish a folder and project hierarchy 

### 建立文件夹和项目层次结构

To capture how your organization manages Google Cloud resources and to enforce a separation of concerns, use [folders](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy#folders) and [ projects](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy#projects). Folders allow different teams to set policies that cascade across multiple projects, while projects can be used to segregate environments (for example, production vs. staging) and teams from each other. For example, most organizations have a team to manage network infrastructure and a different team to manage clusters. Each technology is considered a separate piece of the stack requiring its own level of expertise, troubleshooting and access.

要了解您的组织如何管理 Google Cloud 资源并强制执行关注点分离，请使用 [folders](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy#folders) 和 [项目](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy#projects)。文件夹允许不同的团队设置跨多个项目级联的策略，而项目可用于将环境（例如，生产与暂存）和团队相互隔离。例如，大多数组织都有一个团队来管理网络基础设施，而另一个团队来管理集群。每种技术都被视为堆栈的一个单独部分，需要其自身的专业知识水平、故障排除和访问权限。

A parent folder can contain up to 300 folders, and you can nest folders up to 10 levels deep. If you have over 300 tenants, you can arrange the tenants into nested hierarchies to stay within the limit. For more information about folders, see [Creating and Managing Folders](https://cloud.google.com/resource-manager/docs/creating-managing-folders).

一个父文件夹最多可以包含 300 个文件夹，并且您可以嵌套最多 10 层的文件夹。如果您有超过 300 个租户，您可以将租户安排到嵌套层次结构中以保持在限制范围内。有关文件夹的详细信息，请参阅[创建和管理文件夹](https://cloud.google.com/resource-manager/docs/creating-managing-folders)。

Demonstrating this practice

展示这种做法

For our enterprise environment, we created three top-level folders dedicated to resources for each of the following teams:

对于我们的企业环境，我们为以下每个团队创建了三个专用于资源的顶级文件夹：

- **Network Team**: A folder dedicated for the network team to  manage network resources. This folder contains subfolders for the tenant  network and the cluster network(s), which we discuss further in the    [Centralize network control](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#network-control) section. Each subfolder contains one project per environment (development, staging, and    production) to host the virtual private clouds (VPCs) that    provide all network connectivity in the organization.
- **Cluster Team**: A folder dedicated for the platform team to    manage clusters per environment. This folder contains a subfolder for each    environment (development, staging, and production), each of which contains    one or more projects to accommodate the clusters.
- **Tenants**: A folder dedicated for managing tenants. This    folder contains a subfolder for each tenant to host their non-cluster    resources, each of which may contain one or more projects (or even    subfolders) as required by the individual tenant.

- **Network Team**：专用于网络团队管理网络资源的文件夹。此文件夹包含租户网络和集群网络的子文件夹，我们将在 [集中网络控制](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#network-control) 中进一步讨论多租户部分。每个子文件夹包含每个环境（开发、暂存和生产）的一个项目，用于托管提供组织中所有网络连接的虚拟私有云 (VPC)。

- **Cluster Team**：专门供平台团队管理每个环境的集群的文件夹。此文件夹包含每个环境（开发、登台和生产）的子文件夹，每个环境都包含一个或多个项目以容纳集群。
- **Tenants**：专门用于管理租户的文件夹。此文件夹包含一个子文件夹，供每个租户托管他们的非集群资源，每个租户可能包含一个或多个项目（甚至是子文件夹），以满足单个租户的要求。

  [     ![Folder hierarchy](https://cloud.google.com/static/kubernetes-engine/images/enterprise-folder-hierarchy.svg)   ](https://cloud.google.com/static/kubernetes-engine/images/enterprise-folder-hierarchy.svg)  **Figure 1:** Folder hierarchy

​	 **图 1：** 文件夹层次结构

Note that we recommend per-environment projects for the network and tenant teams, but per-environment folders for the cluster team, where each folder groups projects for each environment (for example, the production folder contains production projects). The reason for this configuration is that the cluster team has specialized segregation needs, and projects are the primary method for segregating resources in Google Cloud. For example, the cluster team might choose to host only one cluster in each project for the following reasons:

请注意，我们建议为网络和租户团队使用每个环境的项目，但为集群团队推荐每个环境的文件夹，其中每个文件夹对每个环境的项目进行分组（例如，生产文件夹包含生产项目）。这种配置的原因是集群团队有专门的隔离需求，而项目是 GCP 中隔离资源的主要方法。例如，集群团队可能出于以下原因选择在每个项目中仅托管一个集群：

- *Cluster configuration*: Some configurations, such as Identity and Access Management (IAM),    are per-project. Placing different clusters in different projects ensures    that a misconfiguration in one project will not affect all clusters in an    environment simultaneously, and allows you to progressively roll out and    validate changes to your configuration.
- *Workload security*: By default, workloads running in different projects are far more segregated from one another than workloads in the  same project. Hosting clusters in dedicated projects ensures that a    compromised, misbehaving or malicious workload in one cluster has limited impact.
- *Resource quota*: Quotas are established and enforced per-project. Spreading clusters across projects limits the impact of a single    workload (for example, in an autoscaled cluster) from exhausting the entire    environment's limits. 

- *集群配置*：某些配置，例如身份和访问管理 (IAM)，是每个项目的。将不同的集群放置在不同的项目中可确保一个项目中的错误配置不会同时影响环境中的所有集群，并允许您逐步推出和验证对配置的更改。
- *工作负载安全*：默认情况下，在不同项目中运行的工作负载比在同一个项目中运行的工作负载相互隔离得更多。在专用项目中托管集群可确保在一个集群中受到损害、行为不端或恶意工作负载的影响有限。
- *资源配额*：配额是按项目建立和执行的。跨项目分布集群可以限制单个工作负载（例如，在自动扩展的集群中）的影响，以免耗尽整个环境的限制。

It may still be useful to apply certain low-risk policies to "all production clusters", regardless of the projects in which they are segregated. The cluster team's per-environment folders allows these kinds of policies to be easily applied. These folders can also be used with aggregated log sinks, allowing for easy per-environment log exporting.

将某些低风险策略应用于“所有生产集群”可能仍然有用，无论它们被隔离在哪个项目中。集群团队的每个环境文件夹允许轻松应用这些类型的策略。这些文件夹也可以与聚合日志接收器一起使用，从而可以轻松导出每个环境的日志。

This recommended topology can easily be extended or simplified depending on your organization's needs. For example, smaller organizations with looser service level objectives (SLOs) may choose to keep all their per-environment clusters in a single project, in which case the per-environment folders are unnecessary. It is also valid to reduce the number of clusters to fit your needs.

这个推荐的拓扑结构可以根据您组织的需要轻松扩展或简化。例如，具有较宽松的服务级别目标 (SLO) 的小型组织可能会选择将其所有每个环境的集群保存在一个项目中，在这种情况下，每个环境的文件夹是不必要的。减少集群数量以满足您的需求也是有效的。

### Assign roles using IAM

### 使用 IAM 分配角色

You can control access to Google Cloud resources through [IAM](https://cloud.google.com/iam/docs/overview) policies. Start by identifying the groups needed for your organization and their scope of operations, then assign the appropriate [IAM role](https://cloud.google.com/iam/docs/understanding-roles) to the group. Use Google Groups to efficiently assign and manage IAM for users.

您可以通过 [IAM](https://cloud.google.com/iam/docs/overview) 政策控制对 Google Cloud 资源的访问。首先确定您的组织所需的组及其操作范围，然后为该组分配适当的 [IAM 角色](https://cloud.google.com/iam/docs/understanding-roles)。使用 Google 网上论坛为用户有效分配和管理 IAM。

Demonstrating this practice

展示这种做法

For our enterprise environment, we defined the following groups and role assignments:

对于我们的企业环境，我们定义了以下组和角色分配：

| Group                                                        | Function                                                     | IAM roles                                                    |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Org Admin                                                    | Organizes the structure of the resources used by the organization. | Organization Administrator, Billing Account Creator, Billing Account  User, Shared VPC Admin, Project Creator |
| Folder Admin                                                 | Creates and manages folders and projects in the organization's        folders. | Folder Admin, Project Creator, Billing Account User          |
| Network Admin                                                | Creates networks, VPCs, subnets, firewall rules, and IP Address Management (IPAM). | Compute Network Admin                                        |
| Security Admin                                               | Manages all logs (and audit logs), secret management, isolation and        incident response. | Compute Security Admin                                       |
| Auditor                                                      | Reviews security events logs and system configurations.      | Private Logs Viewer                                          |
| Cluster Admin                                                | Manages all clusters, including node pools, instances and system        workloads. | Kubernetes Engine Admin                                      |
| Tenant Admin[1](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#fn1) | Manages all tenant namespaces and tenant users.              | Kubernetes Engine Viewer                                     |
| Tenant Developer[1](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#fn1) | Manages and troubleshoots workloads in the tenant namespaces. | Kubernetes Engine Viewer                                     |

1Tenant groups require additional access control in  [Kubernetes RBAC](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-rbac).[↩](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#fnref1)

1租户组在 [Kubernetes RBAC](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-rbac) 中需要额外的访问控制。[↩](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#fnref1)

### Centralize network control 

### 集中网络控制

To maintain centralized control over network resources, such as subnets, routes, and firewalls, use [Shared VPC networks](https://cloud.google.com/vpc/docs/shared-vpc). Resources in a Shared VPC can communicate with each other securely and efficiently across project boundaries using internal IPs. Each Shared VPC network is defined and owned by a centralized *host project*, and can be used by one or more *service projects*.

要保持对网络资源（例如子网、路由和防火墙）的集中控制，请使用 [共享 VPC 网络](https://cloud.google.com/vpc/docs/shared-vpc)。共享 VPC 中的资源可以使用内部 IP 跨项目边界安全高效地相互通信。每个共享 VPC 网络都由一个集中的*宿主项目*定义和拥有，并且可以由一个或多个*服务项目*使用。

Using Shared VPC and IAM, you can separate network administration from project administration. This separation helps you implement the principle of least privilege. For example, a centralized network team can administer the network without having any permissions into the participating projects. Similarly, the project admins can manage their project resources without any permissions to manipulate the shared network.

使用共享 VPC 和 IAM，您可以将网络管理与项目管理分开。这种分离有助于您实施最小特权原则。例如，一个集中的网络团队可以在没有任何参与项目权限的情况下管理网络。同样，项目管理员可以在没有任何权限操作共享网络的情况下管理他们的项目资源。

When you set up a Shared VPC, you must configure the subnets and their secondary IP ranges in the VPC. To determine the subnet size, you need to know the expected number of tenants, the number of Pods and Services they are expected to run, and the maximum and average Pod size. Calculating the total cluster capacity needed will allow for an understanding of the desired instance size, and this provides the total node count. With the total number of nodes, the total IP space consumed can be calculated, and this can provide the desired subnet size.

设置共享 VPC 时，您必须在 VPC 中配置子网及其辅助 IP 范围。要确定子网大小，您需要知道预期的租户数量、预期运行的 Pod 和服务数量以及最大和平均 Pod 大小。计算所需的总集群容量将有助于了解所需的实例大小，这提供了总节点数。使用节点总数，可以计算消耗的总 IP 空间，这可以提供所需的子网大小。

Here are some factors that you should also consider when setting up your network:

以下是设置网络时还应考虑的一些因素：

- The maximum number of service projects that can be attached to a host project is [1,000](https://cloud.google.com/vpc/docs/quota#shared-vpc), and the maximum number of Shared VPC host projects in a single organization is [100](https://cloud.google.com/vpc/docs/quota#shared-vpc).
- The Node, Pod, and Services [IP ranges](https://cloud.google.com/kubernetes-engine/docs/concepts/alias-ips#cluster_sizing) must all be unique. You cannot create a subnet whose primary and secondary IP address ranges overlap.
- The maximum number of Pods and Services for a given GKE cluster is limited by the size of the cluster's secondary ranges.
- The [maximum number of nodes](https://cloud.google.com/kubernetes-engine/docs/concepts/alias-ips#node_limiters) in the cluster is limited by the size of the cluster's subnet's primary IP address range and the cluster's Pod address range.
- For flexibility and more control over IP address management, you can [configure the maximum number of Pods](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr) that can run on a node. By reducing the number of Pods per node, you also reduce the CIDR range allocated per node, requiring fewer IP addresses.

- 一个宿主项目可以附加的服务项目的最大数量为[1,000](https://cloud.google.com/vpc/docs/quota#shared-vpc)，共享VPC宿主项目的最大数量在单个组织中是[100](https://cloud.google.com/vpc/docs/quota#shared-vpc)。
- 节点、Pod 和服务 [IP 范围](https://cloud.google.com/kubernetes-engine/docs/concepts/alias-ips#cluster_sizing) 都必须是唯一的。您不能创建主要和次要 IP 地址范围重叠的子网。
- 给定 GKE 集群的 Pod 和服务的最大数量受集群次要范围大小的限制。
- 集群中的 [最大节点数](https://cloud.google.com/kubernetes-engine/docs/concepts/alias-ips#node_limiters)受集群子网主IP地址范围大小的限制，集群的 Pod 地址范围。
- 为了更灵活地控制 IP 地址管理，您可以 [配置 Pod 的最大数量](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)在一个节点上运行。通过减少每个节点的 Pod 数量，您还可以减少为每个节点分配的 CIDR 范围，从而需要更少的 IP 地址。

To help calculate subnets for your clusters, you can use the [GKE IPAM calculator](https://github.com/GoogleCloudPlatform/gke-ip-address-management) open source tool. IP Address Management (IPAM) enables efficient use of IP space/subnets and avoids having overlaps in ranges, which prevents connectivity options down the road. For information on network ranges in a VPC cluster, see [Creating a VPC-native cluster](https://cloud.google.com/kubernetes-engine/docs/concepts/alias-ips#cluster_sizing).

为了帮助计算集群的子网，您可以使用 [GKE IPAM 计算器](https://github.com/GoogleCloudPlatform/gke-ip-address-management) 开源工具。 IP 地址管理 (IPAM) 可有效利用 IP 空间/子网，并避免范围重叠，从而阻止未来的连接选项。有关 VPC 集群中网络范围的信息，请参阅[创建 VPC 原生集群](https://cloud.google.com/kubernetes-engine/docs/concepts/alias-ips#cluster_sizing)。

Tenants that require further isolation for resources that run outside the shared clusters (such as dedicated Compute Engine VMs) may use their own VPC, which is peered to the Shared VPC run by the networking team. This provides additional security at the cost of increased complexity and numerous other limitations. For more information on peering, see [Using VPC Network Peering](https://cloud.google.com/vpc/docs/using-vpc-peering). In the example below, all tenants have chosen to share a single (per-environment) tenant VPC.

需要进一步隔离在共享集群之外运行的资源（例如专用 Compute Engine 虚拟机）的租户可以使用他们自己的 VPC，该 VPC 与网络团队运行的共享 VPC 对等。这以增加复杂性和许多其他限制为代价提供了额外的安全性。有关对等互连的更多信息，请参阅[使用 VPC 网络对等互连](https://cloud.google.com/vpc/docs/using-vpc-peering)。在下面的示例中，所有租户都选择共享一个（每个环境)租户 VPC。

Demonstrating this practice 

展示这种做法

Our organization has a dedicated network team to manage both the tenant networks and the cluster networks. The Cluster Network folder contains a host project for each environment to host a Shared VPC. This means that the development, staging, and production environments each have their own Shared VPC networks for their service projects to connect to. Each service project contains a cluster that is connected to the associated subnet for each environment.

我们的组织有一个专门的网络团队来管理租户网络和集群网络。 Cluster Network 文件夹包含一个宿主项目，供每个环境托管一个共享 VPC。这意味着开发、暂存和生产环境都有自己的共享 VPC 网络供其服务项目连接。每个服务项目都包含一个集群，该集群连接到每个环境的关联子网。

The Tenant Network folder also contains a host project per environment, and each project hosts a Shared VPC. Tenants A and B are service projects of the tenant network host project and share the same subnet for their non-cluster resources, to reduce networking overhead/IP space and allow the network team to easily control the network and related resources. Each tenant network is peered to the corresponding cluster network in the same environment.

   租户网络文件夹还包含每个环境的宿主项目，并且每个项目都托管一个共享 VPC。租户 A 和 B 是租户网络宿主项目的服务项目，它们的非集群资源共享同一个子网，以减少网络开销/IP 空间，让网络团队轻松控制网络和相关资源。每个租户网络都与同一环境中的相应集群网络对等。

  [     ![Folder hierarchy](https://cloud.google.com/static/kubernetes-engine/images/enterprise-project-architecture.svg)   ](https://cloud.google.com/static/kubernetes-engine/images/enterprise-project-architecture.svg)  **Figure 2:** Project architecture for    Shared VPC networks

**图 2：** 共享 VPC 网络的项目架构

To accommodate each cluster's potential future growth, we created the following CIDR ranges for our networks:

为了适应每个集群的潜在未来增长，我们为我们的网络创建了以下 CIDR 范围：

| Network                     | Subnet             | CIDR Range     | No. of addresses |
| --------------------------- |------------------ |-------------- |---------------- |
| Tenant Network              | Tenant subnet      | `10.0.0.0/16`  | 65,536           |
| Each tenant per environment | `/22-/25`          | 1024 - 128     | |
| Development Network         | Development subnet | `10.17.0.0/16` | 65,536           |
| Pod secondary IP range      | `10.16.0.0/16`     | 65,536         | |
| Service secondary IP range  | `10.18.0.0/16`     | 65,536         | |
| Control plane IP range      | `10.19.0.0/28`     | 16             | |
| Staging Network             | Staging subnet     | `10.33.0.0/16` | 65,536           |
| Pod secondary IP range      | `10.32.0.0/16`     | 65,536         | |
| Service secondary IP range  | `10.34.0.0/16`     | 65,536         | |
| Control plane IP range      | `10.35.0.0/28`     | 16             | |
| Production Network          | Production subnet  | `10.49.0.0/16` | 65,536           |
| Pod secondary IP range      | `10.48.0.0/16`     | 65,536         | |
| Service secondary IP range  | `10.50.0.0/16`     | 65,536         | |
| Control plane IP range      | `10.51.0.0/28`     | 16             | |



### Creating reliable and highly available clusters

### 创建可靠且高可用的集群

Design your cluster architecture for high availability and reliability by implementing the following recommendations:

通过实施以下建议来设计集群架构以实现高可用性和可靠性：

- Create one cluster admin project per cluster to reduce the risk of project-level configurations (for example, IAM bindings) adversely affecting many clusters, and to help provide separation for quota and billing. Cluster admin projects are separate from *tenant* projects, which individual tenants use to manage, for example, their Google Cloud resources.
- Make the production cluster [private](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters) to disable access to the nodes and manage access to the control plane. We also recommend using private clusters for development and staging environments. 

- 每个集群创建一个集群管理项目，以降低项目级配置（例如，IAM 绑定）对许多集群产生不利影响的风险，并帮助提供配额和计费分离。集群管理项目与 *tenant* 项目是分开的，各个租户使用这些项目来管理他们的 GCP 资源等。
- 将生产集群设为 [私有](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters) 以禁用对节点的访问并管理对控制平面的访问。我们还建议将私有集群用于开发和暂存环境。

- Ensure the control plane for the cluster is [regional](https://cloud.google.com/kubernetes-engine/docs/concepts/regional-clusters) to provide high availability for multi-tenancy; any disruptions to the control plane will impact tenants. Please note, there are [cost implications](https://cloud.google.com/kubernetes-engine/docs/concepts/regional-clusters#pricing) with running regional clusters. [Autopilot clusters](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#comparison) are pre-configured as regional clusters.
- Ensure the nodes in your cluster span at least three zones to achieve zonal reliability. For information about the cost of egress between zones in the same region, see the [network pricing](https://cloud.google.com/vpc/network-pricing#general) documentation.

- 确保集群的控制平面是 [regional](https://cloud.google.com/kubernetes-engine/docs/concepts/regional-clusters) 以提供多租户的高可用性；控制平面的任何中断都会影响租户。请注意，运行区域集群会产生 [成本影响](https://cloud.google.com/kubernetes-engine/docs/concepts/regional-clusters#pricing)。 [Autopilot 集群](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#comparison) 预配置为区域集群。

- 确保集群中的节点至少跨越三个区域以实现区域可靠性。有关同一区域中可用区之间的出口成本的信息，请参阅 [网络定价](https://cloud.google.com/vpc/network-pricing#general) 文档。

  [     ![A private regional cluster with a regional control plane running in three zones](https://cloud.google.com/static/kubernetes-engine/images/enterprise-regional-cluster-and-planes.svg)   ](https://cloud.google.com/static/kubernetes-engine/images/enterprise-regional-cluster-and-planes.svg)  **Figure 3:** A private regional cluster with a regional control plane running in three zones

​     **图 3：** 具有在三个区域中运行的区域控制平面的私有区域集群.

#### Autoscale cluster nodes and resources

#### 自动缩放集群节点和资源

To accommodate the demands of your tenants, automatically scale nodes in your cluster by enabling [autoscaling](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-autoscaler). Autoscaling helps systems appear responsive and healthy when heavy workloads are deployed by various tenants in their namespaces, or to respond to zonal outages.

为了满足租户的需求，请通过启用 [autoscaling](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-autoscaler) 自动扩展集群中的节点。当各种租户在其命名空间中部署繁重的工作负载或响应区域性中断时，自动缩放可帮助系统显示出响应性和健康性。

When you enable autoscaling, you specify the minimum and maximum number of nodes in a cluster based on the expected workload sizes. By specifying the maximum number of nodes, you can ensure there is enough space for all Pods in the cluster, regardless of the namespace they run in. Cluster autoscaling rescales node pools based on the min/max boundary, helping to reduce operational costs when the system load decreases, and avoid Pods going into a pending state when there aren't enough available cluster resources. To determine the maximum number of nodes, identify the maximum amount of CPU and memory that each tenant requires, and add those amounts together to get the total capacity that the cluster should be able to handle if all tenants were at the limit. Using the maximum number of nodes, you can then choose instance sizes and counts, taking into consideration the IP subnet space made available to the cluster.

启用自动缩放时，您可以根据预期的工作负载大小指定集群中的最小和最大节点数。通过指定最大节点数，您可以确保集群中的所有 Pod 都有足够的空间，而不管它们在哪个命名空间中运行。集群自动缩放根据最小/最大边界重新缩放节点池，有助于降低运行成本。系统负载降低，避免 Pod 在没有足够可用集群资源时进入挂起状态。要确定最大节点数，请确定每个租户需要的最大 CPU 和内存量，然后将这些数量相加以获得集群在所有租户都处于限制时应该能够处理的总容量。使用最大节点数，您可以选择实例大小和计数，同时考虑到集群可用的 IP 子网空间。

Use Pod autoscaling to automatically scale Pods based on resource demands. [Horizontal Pod Autoscaler (HPA)](https://cloud.google.com/kubernetes-engine/docs/concepts/horizontalpodautoscaler) scales the number of Pod replicas based on CPU/memory utilization or custom metrics. [Vertical Pod Autoscaling (VPA)](https://cloud.google.com/kubernetes-engine/docs/concepts/verticalpodautoscaler) can be used to automatically scale Pods resource demands. It should not be used with HPA unless custom metrics are available as the two autoscalers can compete with each other. For this reason, start with HPA and only later VPA when needed.

使用 Pod 自动缩放功能可根据资源需求自动缩放 Pod。 [Horizontal Pod Autoscaler (HPA)](https://cloud.google.com/kubernetes-engine/docs/concepts/horizontalpodautoscaler) 根据 CPU/内存利用率或自定义指标扩展 Pod 副本的数量。 [垂直 Pod 自动缩放 (VPA)](https://cloud.google.com/kubernetes-engine/docs/concepts/verticalpodautoscaler) 可用于自动缩放 Pod 资源需求。它不应与 HPA 一起使用，除非自定义指标可用，因为两个自动缩放器可以相互竞争。出于这个原因，从 HPA 开始，然后在需要时才使用 VPA。

#### Determine the size of your cluster

#### 确定集群的大小

When determining the size of your cluster, here are some important factors to consider:

在确定集群的大小时，需要考虑以下一些重要因素：

- The sizing of your cluster is dependent on the type of workloads you plan to run. If your workloads have greater density, the cost efficiency is higher but there is also a greater chance for resource contention.
- The minimum size of a cluster is defined by the number of zones it spans: one node for a zonal cluster and three nodes for a regional cluster.
- Per project, there is a maximum of 50 clusters per zone, plus 50 regional clusters per region. 

- 集群的大小取决于您计划运行的工作负载类型。如果您的工作负载密度更高，则成本效率更高，但资源争用的可能性也更大。
- 集群的最小大小由其跨越的区域数量定义：一个节点用于区域集群，三个节点用于区域集群。
- 每个项目，每个区域最多有 50 个集群，每个区域加上 50 个区域集群。

- Per cluster, there is a maximum of 15,000 nodes per cluster (5,000 for GKE versions up to 1.17), 1,000 nodes per node pool, 1,000 nodes per cluster (if you use the GKE Ingress controller), 256 Pods per node (110 for GKE versions older than 1.23.5-gke.1300), 150,000 Pods per cluster, and 300,000 containers per cluster. Refer to the [Quotas and limits page](https://cloud.google.com/kubernetes-engine/quotas) for additional information.

- 每个集群，每个集群最多有 15,000 个节点（GKE 版本最高 1.17 为 5,000），每个节点池 1,000 个节点，每个集群 1,000 个节点（如果您使用 GKE Ingress 控制器），每个节点 256 个 Pod（110 个GKE 版本早于 1.23.5-gke.1300)，每个集群 150,000 个 Pod，每个集群 300,000 个容器。有关更多信息，请参阅[配额和限制页面](https://cloud.google.com/kubernetes-engine/quotas)。

#### Schedule maintenance windows

#### 安排维护时段

To reduce downtimes during cluster/node upgrades and maintenance, schedule [maintenance windows](https://cloud.google.com/kubernetes-engine/docs/concepts/maintenance-windows-and-exclusions) to occur during off-peak hours . During upgrades, there can be temporary disruptions when workloads are moved to recreate nodes. To ensure minimal impact of such disruptions, schedule upgrades for off-peak hours and design your application deployments to handle partial disruptions seamlessly, if possible.

为了减少集群/节点升级和维护期间的停机时间，请将 [维护时段](https://cloud.google.com/kubernetes-engine/docs/concepts/maintenance-windows-and-exclusions)安排在非高峰时段发生.在升级期间，当工作负载被移动以重新创建节点时，可能会出现临时中断。为确保将此类中断的影响降至最低，请在非高峰时间安排升级，并设计您的应用程序部署以无缝处理部分中断（如果可能)。

#### Set up HTTP(S) Load Balancing with Ingress

#### 使用 Ingress 设置 HTTP(S) 负载平衡

To help with the management of your tenants' published [Services](https://cloud.google.com/kubernetes-engine/docs/concepts/service) and the management of incoming traffic to those Services, create an [HTTP(s ) load balancer](https://cloud.google.com/load-balancing/docs/load-balancing-overview) to allow a single ingress per cluster, where each tenant's Services are registered with the cluster's [Ingress](https://cloud.google.com/kubernetes-engine/docs/concepts/ingress) resource. You can create and configure an HTTP(S) load balancer by creating a Kubernetes Ingress resource, which defines how traffic reaches your Services and how the traffic is routed to your tenant's application. By registering Services with the Ingress resource, the Services' naming convention becomes consistent, showing a single ingress, such as `tenanta.example.com` and `tenantb.example.com`.

为了帮助管理您的租户已发布的 [服务](https://cloud.google.com/kubernetes-engine/docs/concepts/service) 以及管理这些服务的传入流量，请创建一个 [HTTP(s ) 负载均衡器](https://cloud.google.com/load-balancing/docs/load-balancing-overview) 允许每个集群有一个入口，其中每个租户的服务都注册到集群的 [入口](https://cloud.google.com/kubernetes-engine/docs/concepts/ingress) 资源。您可以通过创建 Kubernetes Ingress 资源来创建和配置 HTTP(S) 负载均衡器，该资源定义流量如何到达您的服务以及流量如何路由到租户的应用程序。通过使用 Ingress 资源注册服务，服务的命名约定变得一致，显示单个入口，例如 `tenanta.example.com` 和 `tenantb.example.com`。

## Securing the cluster for multi-tenancy

## 保护集群以实现多租户

**Best practices**:

**最佳实践**：

[Control Pod communication with network policies](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#network-policies).
[Run workloads with GKE Sandbox](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#gke-sandbox).
[Create Pod security policies](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#psps).
[Use Workload Identity to grant access to Google Cloud services](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#workload-identity).
[Restrict network access to the control plane](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#control-plane).

[使用网络策略控制 Pod 通信](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#network-policies)。
[使用 GKE 沙盒运行工作负载](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#gke-sandbox)。
[创建 Pod 安全策略](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#psps)。
[使用 Workload Identity 授予对 Google Cloud 服务的访问权限](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#workload-identity)。
[限制对控制平面的网络访问](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#control-plane)。

### Control Pod communication with network policies

### 使用网络策略控制 Pod 通信

To control network communication between Pods in each of your cluster's namespaces, create [network policies](https://cloud.google.com/kubernetes-engine/docs/how-to/network-policy) based on your tenants' requirements. As an initial recommendation, you should block traffic between namespaces that host different tenants' applications. Your cluster administrator can apply a `deny-all` network policy to deny all ingress traffic to avoid Pods from one namespace accidentally sending traffic to Services or databases in other namespaces.

要控制每个集群命名空间中 Pod 之间的网络通信，请根据租户的要求创建 [网络策略](https://cloud.google.com/kubernetes-engine/docs/how-to/network-policy)。作为最初的建议，您应该阻止托管不同租户应用程序的命名空间之间的流量。您的集群管理员可以应用 `deny-all` 网络策略来拒绝所有入口流量，以避免来自一个命名空间的 Pod 意外地将流量发送到其他命名空间中的服务或数据库。

As an example, here's a network policy that restricts ingress from all other namespaces to the `tenant-a` namespace:

例如，这是一个限制从所有其他命名空间进入 `tenant-a` 命名空间的网络策略：

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-all
  namespace: tenant-a
spec:
  podSelector:
    matchLabels:
  ingress:
  - from:
    - podSelector: {}
```

### Run workloads with GKE Sandbox

### 使用 GKE 沙盒运行工作负载

Clusters that run untrusted workloads are more exposed to security vulnerabilities than other clusters. Using [GKE Sandbox](https://cloud.google.com/kubernetes-engine/docs/concepts/sandbox-pods), you can harden the isolation boundaries between workloads for your multi-tenant environment. For security management, we recommend starting with GKE Sandbox and then using Pod security policies to fill in any gaps. 

运行不受信任的工作负载的集群比其他集群更容易受到安全漏洞的影响。使用 [GKE Sandbox](https://cloud.google.com/kubernetes-engine/docs/concepts/sandbox-pods)，您可以强化多租户环境的工作负载之间的隔离边界。对于安全管理，我们建议从 GKE Sandbox 开始，然后使用 Pod 安全策略来填补任何空白。

GKE Sandbox is based on [gVisor](https://gvisor.dev/), an open source container sandboxing project, and provides additional isolation for multi-tenant workloads by adding an extra layer between your containers and host OS. Container runtimes often run as a privileged user on the node and have access to most system calls into the host kernel. In a multi-tenant cluster, one malicious tenant can gain access to the host kernel and to other tenant's data. GKE Sandbox mitigates these threats by reducing the need for containers to interact with the host by shrinking the attack surface of the host and restricting the movement of malicious actors.

GKE Sandbox 基于开源容器沙盒项目 [gVisor](https://gvisor.dev/)，并通过在容器和主机操作系统之间添加一个额外的层来为多租户工作负载提供额外的隔离。容器运行时通常在节点上以特权用户身份运行，并且可以访问对主机内核的大多数系统调用。在多租户集群中，一个恶意租户可以访问主机内核和其他租户的数据。 GKE Sandbox 通过缩小主机的攻击面并限制恶意行为者的移动来减少容器与主机交互的需求，从而缓解这些威胁。

GKE Sandbox provides two isolation boundaries between the container and the host OS:

GKE Sandbox 在容器和主机操作系统之间提供了两个隔离边界：

- A user-space kernel, written in Go, that handles system calls and limits interaction with the host kernel. Each Pod has its own isolated user-space kernel.
- The user-space kernel also runs inside namespaces and seccomp filtering system calls.

- 用 Go 编写的用户空间内核，用于处理系统调用并限制与主机内核的交互。每个 Pod 都有自己独立的用户空间内核。
- 用户空间内核也在命名空间和 seccomp 过滤系统调用中运行。

**Note:** GKE Sandbox is not supported in [Autopilot clusters](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#security_limitations).

**注意：** [Autopilot 集群](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#security_limitations) 不支持 GKE Sandbox。

### Create Pod security policies

### 创建 Pod 安全策略

**Warning:** Kubernetes has officially [deprecated PodSecurityPolicy](https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/) in version 1.21. PodSecurityPolicy will be shut down in version 1.25. For information about alternatives, refer to [PodSecurityPolicy deprecation](https://cloud.google.com/kubernetes-engine/docs/deprecations/podsecuritypolicy).

**警告：** Kubernetes 在 1.21 版本中已正式 [deprecated PodSecurityPolicy](https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/)。 PodSecurityPolicy 将在 1.25 版本中关闭。有关替代方案的信息，请参阅 [PodSecurityPolicy deprecation](https://cloud.google.com/kubernetes-engine/docs/deprecations/podsecuritypolicy)。

To prevent Pods from running in a cluster, create a [Pod Security Policy (PSP)](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies), which specifies conditions that Pods must meet in a cluster. You implement Pod Security Policy control by enabling the admission controller and by authorizing the target Pod's service account to use the policy. You can authorize the use of policies for a Pod in [Kubernetes Role-Based Access Control (RBAC)](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)  by binding the Pod's `serviceAccount ` to a role that has access to use the policies.

要防止 Pod 在集群中运行，请创建 [Pod 安全策略 (PSP)](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies)，其中指定条件Pod 必须在集群中相遇。您可以通过启用准入控制器并授权目标 Pod 的服务帐户使用该策略来实现 Pod 安全策略控制。您可以在 [Kubernetes Role-Based Access Control (RBAC)](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) 中通过绑定 Pod 的 `serviceAccount ` 到有权使用策略的角色。

When defining a PSP, we recommend defining the most restrictive policy bound to `system:authenticated` and more permissive policies bound as needed for exceptions.

在定义 PSP 时，我们建议定义绑定到 `system:authenticated` 的最严格的策略，并根据需要为异常定义更宽松的策略。

As an example, here's a restrictive PSP that requires users to run as unprivileged users, blocks possible escalations to root, and requires the use of several security mechanisms:

例如，这是一个限制性 PSP，它要求用户以非特权用户身份运行，阻止可能升级到 root，并需要使用多种安全机制：

```yaml
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: restricted
spec:
  privileged: false
  # Required to prevent escalations to root.
  allowPrivilegeEscalation: false
  # The following is redundant with non-root + disallow privilege
  # escalation, but we can provide it for defense in depth.
  requiredDropCapabilities:
    - ALL
  # Allow core volume types.
  volumes:
    - 'configMap'
    - 'emptyDir'
    - 'projected'
    - 'secret'
    - 'downwardAPI'
    # Assume that persistentVolumes set up by the cluster admin
    # are safe to use.
    - 'persistentVolumeClaim'
  hostNetwork: false
  hostIPC: false
  hostPID: false
  runAsUser:
    # Require the container to run without root privileges.
    rule: 'MustRunAsNonRoot'
  seLinux:
    # Assumes the nodes are using AppArmor rather than SELinux.
    rule: 'RunAsAny'
  supplementalGroups:
    rule: 'MustRunAs'
    ranges:
      # Forbid adding the root group.
      - min: 1
        max: 65535
  fsGroup:
    rule: 'MustRunAs'
    ranges:
      # Forbid adding the root group.
      - min: 1
        max: 65535
```

Set the following parameters to avoid privilege escalations on the containers:

设置以下参数以避免容器上的特权升级：

- To ensure that no child process of a container can gain more privileges than its parent, set the `allowPrivilegeEscalation` parameter to `false`.
- To disallow escalation privileges outside of the container, disable access to the components of the Host namespaces (`hostNetwork`, `hostIPC`, and `hostPID`). This also blocks snooping on network activity of other Pods on the same node.

- 为确保容器的子进程不能获得比其父进程更多的权限，请将 `allowPrivilegeEscalation` 参数设置为 `false`。
- 要禁止容器外的升级权限，请禁用对主机命名空间组件（`hostNetwork`、`hostIPC` 和 `hostPID`）的访问。这也阻止了对同一节点上其他 Pod 的网络活动的窥探。

**Note:** Pod security policies are not supported in [Autopilot clusters](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#pod_security_policies).

**注意：** [Autopilot 集群](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#pod_security_policies) 不支持 Pod 安全策略。

### Use Workload Identity to grant access to Google Cloud services 

### 使用 Workload Identity 授予对 Google Cloud 服务的访问权限

To securely grant workloads access to Google Cloud services, enable [Workload Identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity) in the cluster. Workload Identity helps administrators manage Kubernetes service accounts that Kubernetes workloads use to access Google Cloud services. When you create a cluster with Workload Identity enabled, an identity namespace is established for the project that the cluster is housed in. The identity namespace allows the cluster to automatically authenticate service accounts for GKE applications by mapping the Kubernetes service account name to a virtual Google service account handle, which is used for IAM binding of tenant Kubernetes service accounts.

要安全地授予工作负载访问 Google Cloud 服务的权限，请在集群中启用 [Workload Identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity)。 Workload Identity 可帮助管理员管理 Kubernetes 工作负载用来访问 Google Cloud 服务的 Kubernetes 服务帐号。当您创建启用了 Workload Identity 的集群时，会为该集群所在的项目建立一个身份命名空间。身份命名空间允许集群通过将 Kubernetes 服务帐户名称映射到虚拟 Google 来自动验证 GKE 应用程序的服务帐户服务账号句柄，用于租户Kubernetes服务账号的IAM绑定。

### Restrict network access to the control plane

### 限制对控制平面的网络访问

To protect your control plane, restrict access to authorized networks. In GKE, when you enable [authorized networks](https://cloud.google.com/kubernetes-engine/docs/how-to/authorized-networks), you can authorize up to 50 CIDR ranges and allow IP addresses only in those ranges to access your control plane. GKE already uses Transport Layer Security (TLS) and authentication to provide secure access to your control plane endpoint from the public internet. By using authorized networks, you can further restrict access to specified sets of IP addresses.

为了保护您的控制平面，请限制对授权网络的访问。在 GKE 中，当您启用 [授权网络](https://cloud.google.com/kubernetes-engine/docs/how-to/authorized-networks) 时，您最多可以授权 50 个 CIDR 范围，并且仅允许 IP 地址在这些范围来访问您的控制平面。 GKE 已经使用传输层安全 (TLS) 和身份验证来提供从公共互联网对控制平面端点的安全访问。通过使用授权网络，您可以进一步限制对指定 IP 地址集的访问。

## Tenant provisioning

## 租户配置

**Best practices**:

**最佳实践**：

[Create tenant projects](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-projects).
[Use RBAC to refine tenant access](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-rbac).
[Create namespaces for isolation between tenants](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-namespaces).

[创建租户项目](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-projects)。
[使用 RBAC 优化租户访问](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-rbac)。
[为租户之间的隔离创建命名空间](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#create-namespaces)。

### Create tenant projects

### 创建租户项目

To host a tenant's non-cluster resources, create a service project for each tenant. These service projects contain logical resources specific to the tenant applications (for example, logs, monitoring, storage buckets, service accounts, etc.). All tenant service projects are connected to the Shared VPC in the tenant host project.

要托管租户的非集群资源，请为每个租户创建一个服务项目。这些服务项目包含特定于租户应用程序的逻辑资源（例如，日志、监控、存储桶、服务帐户等）。所有租户服务项目都连接到租户宿主项目中的共享 VPC。

### Use RBAC to refine tenant access

### 使用 RBAC 优化租户访问

Define finer-grained access to cluster resources for your tenants by using [Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/). On top of the read-only access initially granted with IAM to tenant groups, define namespace-wide Kubernetes RBAC roles and bindings for each tenant group.

使用 [Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) 为您的租户定义对集群资源的更细粒度的访问。除了最初通过 IAM 授予租户组的只读访问权限之外，还可以为每个租户组定义命名空间范围的 Kubernetes RBAC 角色和绑定。

Earlier we identified two tenant groups: tenant admins and tenant developers. For those groups, we define the following RBAC roles and access:

之前我们确定了两个租户组：租户管理员和租户开发人员。对于这些组，我们定义了以下 RBAC 角色和访问权限：

| Group            | Kubernetes RBAC role              | Description                                                  |
| ---------------- |--------------------------------- |------------------------------------------------------------ |
| Tenant Admin     | namespace admin                   | Grants access to list and watch deployments in their namespace. Grants access to add and remove users in the tenant group. |
| Tenant Developer | namespace admin, namespace viewer | Grants access to create/edit/delete Pods, deployments, Services,        configmaps in their namespace. |



In addition to creating RBAC roles and bindings that assign Google Workspace or Cloud Identity groups various permissions inside their namespace, Tenant admins often require the ability to manage users in each of those groups. Based on your organization's requirements, this can be handled by either delegating Google Workspace or Cloud Identity permissions to the Tenant admin to manage their own group membership or by the Tenant admin engaging with a team in your organization that has Google Workspace or Cloud Identity permissions to handle those changes.

除了创建 RBAC 角色和绑定以在其命名空间内为 G Suite 或 Cloud Identity 组分配各种权限外，租户管理员通常需要能够管理每个组中的用户。根据您组织的要求，可以通过将 Google Workspace 或 Cloud Identity 权限委派给租户管理员来管理他们自己的群组成员资格，或者由租户管理员与您组织中拥有 Google Workspace 或 Cloud Identity 权限的团队合作来处理处理这些变化。

Demonstrating this practice

展示这种做法

For our enterprise model, we created a manifest with the following Kubernetes RBAC roles, binded to the tenant groups mentioned above:

对于我们的企业模型，我们创建了一个清单，其中包含以下 Kubernetes RBAC 角色，绑定到上述租户组：

- **namespace admin**: Defined with the `admin`    `ClusterRole` in a `RoleBinding` to allow read and    write access for resources in its namespace, including the ability to create    roles and role bindings in the namespace. 

- **命名空间管理员**：使用 `RoleBinding` 中的 `admin` `ClusterRole` 定义，以允许对其命名空间中的资源进行读写访问，包括在命名空间中创建角色和角色绑定的能力。

- **namespace editor**: Defined with the `edit`    `ClusterRole` in a `RoleBinding` to allow read/write    access to Pods, deployments, Services, configmaps in the tenant namespace.
- **namespace viewer**: Defined with the `view`    `ClusterRole` in a `RoleBinding` to allow read-only    access to Pods, deployments, Services, configmaps in the tenant namespace.

- **命名空间编辑器**：使用 `RoleBinding` 中的 `edit` `ClusterRole` 定义，以允许对租户命名空间中的 Pod、部署、服务、配置映射进行读/写访问。
- **命名空间查看器**：使用 `RoleBinding` 中的 `view` `ClusterRole` 定义，以允许对租户命名空间中的 Pod、部署、服务、配置映射进行只读访问。

You can use IAM and RBAC permissions together with namespaces to restrict user interactions with cluster resources on console. For more information, see [   Enable access and view cluster resources by namespace](https://cloud.google.com/kubernetes-engine/docs/how-to/restrict-resources-access-by-namespace).

您可以将 IAM 和 RBAC 权限与命名空间一起使用，以限制用户在控制台上与集群资源的交互。有关详细信息，请参阅 [按命名空间启用访问和查看集群资源](https://cloud.google.com/kubernetes-engine/docs/how-to/restrict-resources-access-by-namespace)。

#### Use Google Groups to bind permissions

#### 使用谷歌群组绑定权限

To efficiently manage tenant permissions in a cluster, you can bind RBAC permissions to your [Google Groups](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control#google-groups-for-gke). The membership of those groups are maintained by your Google Workspace administrators, so your cluster administrators do not need detailed information about your users.

为了有效地管理集群中的租户权限，您可以将 RBAC 权限绑定到您的 [Google 群组](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control#google-groups-for-gke)。这些群组的成员资格由您的 G Suite 管理员维护，因此您的集群管理员不需要您的用户的详细信息。

As an example, we have a Google Group named `tenant-admins@mydomain.com` and a user named `admin1@mydomain.com` is a member of that group, the following binding provides the user with admin access to the `tenant -a` namespace:

例如，我们有一个名为 `tenant-admins@mydomain.com` 的 Google 组，并且名为 `admin1@mydomain.com` 的用户是该组的成员，以下绑定为用户提供了对 `tenant 的管理员访问权限-a` 命名空间：

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: tenant-a
  name: tenant-admin-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: tenant-admin
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: "tenant-admins@mydomain.com"
```

### Create namespaces

### 创建命名空间

To provide a logical isolation between tenants that are on the same cluster, implement [namespaces](https://kubernetes.io/docs/tasks/administer-cluster/namespaces/). As part of the Kubernetes RBAC process, the cluster admin creates namespaces for each tenant group. The Tenant admin manages users (tenant developers) within their respective tenant namespace. Tenant developers are then able to use cluster and tenant specific resources to deploy their applications.

要在同一集群上的租户之间提供逻辑隔离，请实现 [namespaces](https://kubernetes.io/docs/tasks/administer-cluster/namespaces/)。作为 Kubernetes RBAC 流程的一部分，集群管理员为每个租户组创建命名空间。租户管理员在其各自的租户命名空间内管理用户（租户开发人员)。然后，租户开发人员能够使用集群和租户特定资源来部署他们的应用程序。

#### Avoid reaching namespace limits

#### 避免达到命名空间限制

The theoretical maximum number of namespaces in a cluster is 10,000, though in practice there are many factors that could prevent you from reaching this limit. For example, you might reach the cluster-wide maximum number of Pods (150,000) and nodes (5,000) before you reach the maximum number of namespaces; other factors (such as the number of Secrets) can further reduce the effective limits. As a result, a good initial rule of thumb is to only attempt to approach the theoretical limit of one constraint at a time, and stay approximately one order of magnitude away from the other limits, unless experimentation shows that your use cases work well. If you need more resources than can be supported by a single cluster, you should create more clusters. For information about Kubernetes scalability, see the [Kubernetes Scalability thresholds](https://github.com/kubernetes/community/blob/master/sig-scalability/configs-and-limits/thresholds.md)  article.

理论上，集群中命名空间的最大数量为 10,000，但实际上有许多因素可能会阻止您达到此限制。例如，在达到最大命名空间数量之前，您可能会达到集群范围内的最大 Pod 数量（150,000）和节点（5,000）；其他因素（例如 Secret 的数量）可以进一步降低有效限制。因此，一个好的初始经验法则是一次只尝试接近一个约束的理论限制，并与其他限制保持大约一个数量级，除非实验表明您的用例运行良好。如果您需要的资源超出单个集群的支持范围，则应创建更多集群。有关 Kubernetes 可扩展性的信息，请参阅 [Kubernetes 可扩展性阈值](https://github.com/kubernetes/community/blob/master/sig-scalability/configs-and-limits/thresholds.md) 文章。

#### Standardize namespace naming

#### 标准化命名空间命名

To ease deployments across multiple environments that are hosted in different clusters, standardize the namespace naming convention you use. For example, avoid tying the environment name (development, staging, and production) to the namespace name and instead use the same name across environments. By using the same name, you avoid having to change the config files across environments.

要简化跨托管在不同集群中的多个环境的部署，请标准化您使用的命名空间命名约定。例如，避免将环境名称（开发、登台和生产）与命名空间名称绑定，而是在不同环境中使用相同的名称。通过使用相同的名称，您可以避免跨环境更改配置文件。

#### Create service accounts for tenant workloads 

#### 为租户工作负载创建服务帐户

Create a tenant-specific Google service account for each distinct workload in a tenant namespace. This provides a form of security, ensuring that tenants can manage service accounts for the workloads that they own/deploy in their respective namespaces. The Kubernetes service account for each namespace is mapped to one Google service account by using [Workload Identity](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#workload-identity).

为租户命名空间中的每个不同工作负载创建特定于租户的 Google 服务帐户。这提供了一种安全形式，确保租户可以管理他们在各自命名空间中拥有/部署的工作负载的服务帐户。使用 [Workload Identity](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#workload-identity) 将每个命名空间的 Kubernetes 服务帐户映射到一个 Google 服务帐户。

### Enforce resource quotas

### 强制资源配额

To ensure all tenants that share a cluster have fair access to the cluster resources, enforce [resources quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/). Create a resource quota for each namespace based on the number of Pods deployed by each tenant, and the amount of memory and CPU required by each Pod.

为确保共享集群的所有租户都能公平访问集群资源，请强制执行 [resources quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/)。根据每个租户部署的 Pod 数量，以及每个 Pod 所需的内存和 CPU 量，为每个命名空间创建资源配额。

The following example defines a resource quota where Pods in the `tenant-a` namespace can request up to 16 CPU and 64 GB of memory, and the maximum CPU is 32 and the maximum memory is 72 GB.

以下示例定义了一个资源配额，其中 `tenant-a` 命名空间中的 Pod 最多可以请求 16 个 CPU 和 64 GB 内存，最大 CPU 为 32，最大内存为 72 GB。

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: tenant-a
spec:
  hard: "1"
    requests.cpu: "16"
    requests.memory: 64Gi
    limits.cpu: "32"
    limits.memory: 72Gi
```

## Monitoring, logging and usage

## 监控、记录和使用

**Best practices**:

**最佳实践**：

[Track usage metrics](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#usage-metrics).
[Provide tenant-specific logs](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-logs).
[Provide tenant-specific monitoring](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-monitoring).

[跟踪使用指标](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#usage-metrics)。
[提供特定于租户的日志](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-logs)。
[提供特定于租户的监控](https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#tenant-monitoring)。

### Track usage metrics

### 跟踪使用指标

To obtain cost breakdowns on individual namespaces and labels in a cluster, you can enable [GKE usage metering](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-usage-metering). GKE usage metering tracks information about resource requests and resource usage of a cluster's workloads, which you can further break down by namespaces and labels. With GKE usage metering, you can approximate the cost breakdown for departments/teams that are sharing a cluster, understand the usage patterns of individual applications (or even components of a single application), help cluster admins triage spikes in usage, and provide better capacity planning and budgeting.

要获取集群中各个命名空间和标签的成本明细，您可以启用 [GKE 使用计量](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-usage-metering)。 GKE 使用计量跟踪有关集群工作负载的资源请求和资源使用情况的信息，您可以按命名空间和标签进一步细分这些信息。借助 GKE 使用计量，您可以估算共享集群的部门/团队的成本细分，了解单个应用程序（甚至单个应用程序的组件)的使用模式，帮助集群管理员对使用高峰进行分类，并提供更好的容量计划和预算。

When you enable GKE usage metering on the multi-tenant cluster, resource usage records are written to a BigQuery table. You can export tenant-specific metrics to BigQuery datasets in the corresponding tenant project, which auditors can then analyze to determine cost breakdowns. Auditors can visualize GKE usage metering data by creating dashboards with plug-and-play Google Data Studio templates.

当您在多租户集群上启用 GKE 使用计量时，资源使用记录将写入 BigQuery 表。您可以将特定于租户的指标导出到相应租户项目中的 BigQuery 数据集，然后审计人员可以对其进行分析以确定成本明细。审核员可以通过使用即插即用的 Google Data Studio 模板创建仪表板来可视化 GKE 使用计量数据。

**Note:** GKE usage metering is not supported in [Autopilot clusters](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#add-ons).

**注意：** [Autopilot 集群](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#add-ons) 不支持 GKE 使用计量。

### Provide tenant-specific logs

### 提供特定于租户的日志

To provide tenants with log data specific to their project workloads, use Cloud Monitoring's [Log Router](https://cloud.google.com/logging/docs/routing/overview). To create tenant-specific logs, the cluster admin creates a [sink](https://cloud.google.com/logging/docs/routing/overview#sinks) to export log entries to a [log bucket](https://cloud.google.com/logging/docs/buckets) created in the tenant's Google Cloud project. For details on how to configure these types of logs, see [Multi-tenant logging on GKE](https://cloud.google.com/stackdriver/docs/solutions/kubernetes-engine/multi-tenant-logging).

要为租户提供特定于其项目工作负载的日志数据，请使用 Cloud Monitoring 的 [日志路由器](https://cloud.google.com/logging/docs/routing/overview)。要创建特定于租户的日志，集群管理员会创建一个[sink](https://cloud.google.com/logging/docs/routing/overview#sinks) 以将日志条目导出到 [log bucket](https:///cloud.google.com/logging/docs/buckets)在租户的 Google Cloud 项目中创建。有关如何配置这些类型的日志的详细信息，请参阅[GKE 上的多租户日志记录](https://cloud.google.com/stackdriver/docs/solutions/kubernetes-engine/multi-tenant-logging)。

### Provide tenant-specific monitoring 

### 提供租户特定的监控

To provide tenant-specific monitoring, the cluster admin can use a dedicated namespace that contains a [Prometheus](https://prometheus.io/)  to Stackdriver adapter ([prometheus-to-sd](https://github.com/GoogleCloudPlatform/k8s-stackdriver/tree/master/prometheus-to-sd)) with a per namespace config. This configuration ensures tenants can only monitor their own metrics in their projects. However, the downside to this design is the extra cost of managing your own Prometheus deployment(s).

为了提供特定于租户的监控，集群管理员可以使用包含 [Prometheus](https://prometheus.io/) 到 Stackdriver 适配器 ([prometheus-to-sd](https://github.com)的专用命名空间/GoogleCloudPlatform/k8s-stackdriver/tree/master/prometheus-to-sd)) 每个命名空间配置。这种配置确保租户只能在他们的项目中监控他们自己的指标。但是，这种设计的缺点是管理您自己的 Prometheus 部署的额外成本。

Here are other options you could consider for providing tenant-specific monitoring:

以下是您可以考虑提供特定租户监控的其他选项：

- Teams accept shared tenancy within the Cloud Monitoring environment and allow tenants to have visibility into all metrics in the project.
- Deploy a single [Grafana](https://grafana.com/) instance per tenant, which communicates with the shared Cloud Monitoring environment. Configure the Grafana instance to only view the metrics from a particular namespace. The downside to this option is the cost and overhead of managing these additional deployments of Grafana. For more information, see [Using Cloud Monitoring in Grafana](https://grafana.com/docs/grafana/latest/datasources/google-cloud-monitoring/).

- 团队接受 Cloud Monitoring 环境中的共享租赁，并允许租户了解项目中的所有指标。
- 为每个租户部署一个 [Grafana](https://grafana.com/) 实例，该实例与共享的 Cloud Monitoring 环境进行通信。将 Grafana 实例配置为仅查看来自特定命名空间的指标。此选项的缺点是管理这些额外的 Grafana 部署的成本和开销。如需了解详情，请参阅 [在 Grafana 中使用 Cloud Monitoring](https://grafana.com/docs/grafana/latest/datasources/google-cloud-monitoring/)。

## Checklist summary

## 清单摘要

The following table summarizes the tasks that are recommended for creating multi-tenant clusters in an enterprise organization:

下表总结了建议在企业组织中创建多租户集群的任务：

https://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy#checklist

## What's next

##  下一步是什么

- For more information on security, see [Hardening your cluster's security](https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster).
- For more information on VPC networks, see [Best practices and reference architectures for VPC design](https://cloud.google.com/solutions/best-practices-vpc-design).
- For more enterprise best practices, see [Google Cloud Architecture Framework](https://cloud.google.com/architecture/framework). 

- 有关安全性的更多信息，请参阅[强化集群的安全性](https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster)。
- 有关 VPC 网络的更多信息，请参阅[VPC 设计的最佳实践和参考架构](https://cloud.google.com/solutions/best-practices-vpc-design)。
- 有关更多企业最佳实践，请参阅 [Google Cloud 架构框架](https://cloud.google.com/architecture/framework)。

