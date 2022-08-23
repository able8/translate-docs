# Approaches to implementing multi-tenancy in SaaS applications

# 在 SaaS 应用程序中实现多租户的方法

May 19, 2022 From: https://developers.redhat.com/articles/2022/05/09/approaches-implementing-multi-tenancy-saas-applications#

This article discusses architectural approaches for separating and isolating SaaS tenants to provide [_multi-tenancy_](https://www.redhat.com/en/topics/cloud-computing/what-is-multitenancy), the provisioning of services to multiple clients in different organizations. For the approaches, the type and level of isolation provided are compared, along with their tradeoffs.

本文讨论了分离和隔离 SaaS 租户以提供 [_multi-tenancy_](https://www.redhat.com/en/topics/cloud-computing/what-is-multitenancy)、向多个不同组织的客户。对于这些方法，比较了所提供的隔离类型和级别，以及它们的权衡。

The approaches laid out in different sections of the article are not mutually exclusive and can be combined to provide the levels of separation and isolation necessary to satisfy the requirements of your SaaS customers and markets. We'll also discuss how to incorporate existing single-tenant applications into a SaaS environment.

本文不同部分中列出的方法并不相互排斥，可以组合起来提供满足 SaaS 客户和市场需求所需的分离和隔离级别。我们还将讨论如何将现有的单租户应用程序整合到 SaaS 环境中。

## Multi-tenancy considerations

## 多租户注意事项

The resources that are shared across multiple tenants vary, based on the architecture of the SaaS application. The level of sharing can be very high in large business-to-consumer (B2C) SaaS applications. In these cases, a given application instance potentially handles requests from thousands of unrelated tenants. For a more sensitive business-to-business (B2B) application, each tenant might get a dedicated application instance, though it could still be running on shared infrastructure.

跨多个租户共享的资源因 SaaS 应用程序的体系结构而异。在大型企业对消费者 (B2C) SaaS 应用程序中，共享级别可能非常高。在这些情况下，给定的应用程序实例可能会处理来自数千个不相关租户的请求。对于更敏感的企业对企业 (B2B) 应用程序，每个租户可能会获得一个专用的应用程序实例，尽管它仍然可以在共享基础架构上运行。

Perhaps the most important aspect of a SaaS service is making sure tenants feel they have privacy and that no other tenant can see their data or activities. Additionally, tenant data must be stored securely and reliably. The integrity of each tenant's data must be protected from accidental loss, modification, or tampering.

也许 SaaS 服务最重要的方面是确保租户觉得他们有隐私，并且没有其他租户可以看到他们的数据或活动。此外，租户数据必须安全可靠地存储。必须保护每个租户数据的完整性，以免意外丢失、修改或篡改。

When architecting a SaaS application, consider the level of isolation and controls required for the service you plan to offer. The requirements vary depending on the nature of the application, the sensitivity of the data, and the type of market. The contract or terms of service for the application usually detail what the provider is promising. However, regulatory or government agencies might issue compliance requirements for specific industries, such as financial services or health care, that you must adhere to when providing a SaaS application to customers under their jurisdiction.

在构建 SaaS 应用程序时，请考虑您计划提供的服务所需的隔离和控制级别。要求因应用程序的性质、数据的敏感性和市场类型而异。应用程序的合同或服务条款通常会详细说明提供商的承诺。但是，监管机构或政府机构可能会针对特定行业（例如金融服务或医疗保健）发布合规性要求，您在向其管辖范围内的客户提供 SaaS 应用程序时必须遵守这些要求。

It is worth considering that a one-size-fits-all approach might not be enough to balance the security requirements and costs for all prospective customers. Customers from highly regulated industries or groups might not be the target for your SaaS application. However, if you do attract customers from financial services, health care, or government organizations, you might have to adhere to requirements you didn't originally anticipate.

值得考虑的是，一刀切的方法可能不足以平衡所有潜在客户的安全要求和成本。来自高度监管的行业或群体的客户可能不是您的 SaaS 应用程序的目标。但是，如果您确实吸引了金融服务、医疗保健或政府组织的客户，您可能必须遵守您最初没有预料到的要求。

Also, some customers are under data sovereignty regulations that specify where the physical machines storing data must reside. These requirements potentially dictate data center locations and even which cloud providers can be used to host the SaaS services the customers consume.

此外，一些客户受数据主权法规的约束，这些法规规定了存储数据的物理机器必须驻留在何处。这些要求可能决定数据中心的位置，甚至可以使用哪些云提供商来托管客户使用的 SaaS 服务。

## Using application logic to provide multi-tenancy

## 使用应用程序逻辑提供多租户

When you think of SaaS applications, the ones that come to mind are probably those from large technology companies that are offered to millions of consumers. These applications are built from the ground up for multi-tenancy. They are designed for high-density deployments where large numbers of users can be served using minimal per-user resources. To achieve the necessary efficiency, a very high degree of sharing is usually necessary.

当您想到 SaaS 应用程序时，脑海中浮现的可能是大型科技公司提供给数百万消费者的应用程序。这些应用程序是为多租户而构建的。它们专为高密度部署而设计，其中可以使用最少的每用户资源为大量用户提供服务。为了达到必要的效率，通常需要非常高度的共享。

In an application design where each software instance handles requests from multiple tenants, the infrastructure generally can't provide the needed isolation. Instead, application logic is typically used to implement the controls that isolate tenants in application-level tenancy.

在每个软件实例处理来自多个租户的请求的应用程序设计中，基础设施通常无法提供所需的隔离。相反，应用程序逻辑通常用于实现在应用程序级租户中隔离租户的控件。

In this model, a single instance of the application serves requests from multiple tenants. The code is responsible for keeping each tenant's data separate and making sure other tenants can't see any data or activity that does not belong to them. 

在此模型中，应用程序的单个实例为来自多个租户的请求提供服务。该代码负责将每个租户的数据分开，并确保其他租户看不到任何不属于他们的数据或活动。

The controls in application-level tenancy must be implemented by software developers. Therefore, extensive [automated](http://developers.redhat.com/topics/automation) testing and strict adherence to secure coding best practices are critical for application tenancy, as any bugs in the application could accidentally disclose tenant data or compromise data integrity. If a vulnerability in the application is found and exploited, relying solely on application logic to implement isolation could make it difficult to contain and mitigate the potential impact to tenants that share the same instances.

应用程序级租赁中的控制必须由软件开发人员实施。因此，广泛的 [自动化](http://developers.redhat.com/topics/automation) 测试和严格遵守安全编码最佳实践对于应用程序租赁至关重要，因为应用程序中的任何错误都可能意外泄露租户数据或损害数据正直。如果发现并利用应用程序中的漏洞，仅依靠应用程序逻辑来实现隔离可能会难以控制和减轻对共享相同实例的租户的潜在影响。

Application tenancy is best for applications that are designed for multi-tenancy as they are developed. Retrofitting an existing application to support multiple tenants can require extensive software development work as well as extensive testing. The costs and time necessary to modify an existing application could be high. For those applications, it might be better to consider the next two approaches we'll discuss in this article, where infrastructure-level controls provide the necessary separation and isolation.

应用程序租赁最适合在开发时为多租户设计的应用程序。改造现有应用程序以支持多个租户可能需要大量的软件开发工作以及大量测试。修改现有应用程序所需的成本和时间可能很高。对于这些应用程序，最好考虑我们将在本文中讨论的接下来的两种方法，其中基础设施级别的控制提供了必要的分离和隔离。

### Advantages of application-level tenancy

### 应用程序级租赁的优势

Implementing multi-tenancy in the code of your application itself is appealing for a number of reasons:

在应用程序本身的代码中实现多租户很有吸引力，原因有很多：

- No infrastructure support is necessary to implement application-level tenancy.
- Higher tenant density is possible, resulting in lower infrastructure costs because application instances can be shared by multiple tenants.
- In many cases, no additional processes or infrastructure need to be provisioned when adding tenants.

- 无需基础设施支持即可实施应用程序级租赁。
- 更高的租户密度是可能的，从而降低基础设施成本，因为应用程序实例可以由多个租户共享。
- 在许多情况下，添加租户时不需要配置额外的流程或基础设施。

### Disadvantages of application-level tenancy

### 应用程序级租户的缺点

However, there are also reasons you might want to avoid this architecture:

但是，您可能还希望避免这种架构的原因：

- This approach entails high development costs, because isolation must be implemented by software developers.
- Extensive security testing is needed whenever changes are made to validate isolation is correctly enforced.
- This approach provides a single layer of defense, so bugs or vulnerabilities in the application could accidentally disclose or corrupt tenant data.
- Security reviews by outside auditors or regulators can be difficult and labor intensive, because they need to understand and validate the controls implemented by the application.
- Change management occurs at a very coarse granularity. Upgrades, downgrades, or patches affect the entire tenant population or large subsets of it, because application instances span multiple tenants.
- It is difficult and costly to retrofit existing applications that were not designed for multi-tenant SaaS deployments.

- 这种方法需要高昂的开发成本，因为隔离必须由软件开发人员实施。
- 每当进行更改以验证隔离是否正确实施时，都需要进行广泛的安全测试。
- 此方法提供单层防御，因此应用程序中的错误或漏洞可能会意外泄露或损坏租户数据。
- 外部审计员或监管机构的安全审查可能很困难且需要大量人力，因为他们需要了解和验证应用程序实施的控制措施。
- 变更管理以非常粗略的粒度进行。升级、降级或补丁会影响整个租户群体或其大部分子集，因为应用程序实例跨越多个租户。
- 改造不是为多租户 SaaS 部署设计的现有应用程序既困难又昂贵。

## Isolating tenants on the same cluster using namespaces

## 使用命名空间隔离同一集群上的租户

An alternative to the application-level tenancy described in the previous section is to use the capabilities available in [Kubernetes](http://developers.redhat.com/topics/kubernetes) to isolate applications that share a cluster. Using these controls, it is possible to run multiple instances of the same unmodified single-tenant application on the same cluster while keeping them logically separated. _Namespace-level tenancy_ can allow existing applications that were not written for SaaS to be deployed as a SaaS solution.

上一节中描述的应用程序级租户的替代方法是使用 [Kubernetes](http://developers.redhat.com/topics/kubernetes) 中可用的功能来隔离共享集群的应用程序。使用这些控件，可以在同一集群上运行相同未修改的单租户应用程序的多个实例，同时保持它们在逻辑上分离。 _命名空间级租户_可以允许将不是为 SaaS 编写的现有应用程序部署为 SaaS 解决方案。

Namespaces within Kubernetes provide a mechanism for isolating resources into different groups that keep user communities or application groups separate even while they share a cluster. For multi-tenancy, namespaces provide isolation between tenant workloads. Because Kubernetes enforces separation of namespaces, an application running inside of a namespace can't accidentally or intentionally access data or processes in a different namespace.

Kubernetes 中的命名空间提供了一种将资源隔离到不同组中的机制，即使用户社区或应用程序组共享一个集群，它们也可以保持分离。对于多租户，命名空间提供租户工作负载之间的隔离。因为 Kubernetes 强制执行命名空间的分离，所以在命名空间内运行的应用程序不能意外或有意访问不同命名空间中的数据或进程。

Kubernetes namespaces provide ways to use a single-tenant application in a multi-tenant SaaS deployment. The most common and perhaps most straightforward approach is to create a namespace for each tenant, and deploy the complete application stack and data store into each tenant's namespace. The namespace provides isolation between each application instance. This approach is referred to as _namespace-level tenancy_, and is illustrated in Figure 1. 

Kubernetes 命名空间提供了在多租户 SaaS 部署中使用单租户应用程序的方法。最常见也是最直接的方法是为每个租户创建一个命名空间，并将完整的应用程序堆栈和数据存储部署到每个租户的命名空间中。命名空间提供每个应用程序实例之间的隔离。这种方法被称为 _namespace-leveltenancy_，如图 1 所示。

[![Two tenants are shown with namespace-level tenancy, storing each application's resources including data in its own namespace.](http://developers.redhat.com/sites/default/files/styles/article_full_width_1440px_w/public/image1_7.png?itok=FdzkVvnl)](http://developers.redhat.com/sites/default/files/image1_7.png)

png?itok=FdzkVvnl)](http://developers.redhat.com/sites/default/files/image1_7.png)

Figure 1. Two tenants are shown with namespace-level tenancy, storing each application's resources including data in its own namespace.

图 1. 显示了两个具有命名空间级别的租户，将每个应用程序的资源（包括数据）存储在其自己的命名空间中。

Figure 1: Two tenants are shown with namespace-level tenancy, storing each application's resources including data in its own namespace.

图 1：显示了两个具有命名空间级别的租户的租户，将每个应用程序的资源（包括数据）存储在其自己的命名空间中。

To provide higher degrees of isolation and better control over resources, cluster administrators can control which worker nodes are available for each tenant's pods. Specific tenants can be configured to use only dedicated worker nodes, which could run on dedicated hardware if desired.

为了提供更高程度的隔离和更好的资源控制，集群管理员可以控制每个租户的 pod 可以使用哪些工作节点。可以将特定租户配置为仅使用专用工作节点，如果需要，这些节点可以在专用硬件上运行。

With namespace-level tenancy, a single control plane manages all the application instances and worker nodes used by all tenants of the cluster. The cluster's control plane contains configuration information for all tenant application instances running on that cluster. Anyone with cluster administration access can view and change configuration information for all tenants of the cluster.

使用命名空间级别的租赁，单个控制平面管理集群的所有租户使用的所有应用程序实例和工作节点。集群的控制平面包含在该集群上运行的所有租户应用程序实例的配置信息。任何拥有集群管理访问权限的人都可以查看和更改集群所有租户的配置信息。

### Advantages of namespace-level tenancy

### 命名空间级租赁的优点

This architecture is appealing for a number of reasons:

这种架构吸引人的原因有很多：

- Single-tenant applications can be used in SaaS deployments by using namespace controls to run isolated application instances for each tenant. The application does not need to be modified for SaaS use, because the cluster provides multi-tenancy control.
- The impacts of bugs and vulnerabilities within the application or application components are contained to a single tenant running in that namespace.
- Individual tenants can have customized deployments, making it easier to add features to a select set or add third-party integrations.
- Resource consumption is easy to view at the level of an individual tenant.
- The resources available for scaling the application up or down can be optimized at the level of a single tenant.
- Application changes can be managed flexibly at a granular level. The cadence of new releases and patches can be optimized on a per-tenant basis.

- 通过使用命名空间控制为每个租户运行隔离的应用程序实例，可以在 SaaS 部署中使用单租户应用程序。应用程序不需要为 SaaS 使用而修改，因为集群提供了多租户控制。
- 应用程序或应用程序组件中的错误和漏洞的影响包含在该命名空间中运行的单个租户中。
- 单个租户可以进行自定义部署，从而更轻松地向选择集添加功能或添加第三方集成。
- 资源消耗很容易在单个租户级别查看。
- 可用于向上或向下扩展应用程序的资源可以在单个租户级别进行优化。
- 可以在粒度级别灵活管理应用程序更改。可以在每个租户的基础上优化新版本和补丁的节奏。

### Disadvantages of namespace-level tenancy

### 命名空间级租赁的缺点

However, there are reasons why you might decide not to use this architecture:

但是，您可能决定不使用此架构是有原因的：

- Deployment density is lower and infrastructure costs are higher than in application-level tenancy, because the cluster has to run multiple application instances that are not shared. In the simplest case, a complete application stack is run for each tenant.
- Tenant onboarding requires the cluster to create a new project or namespace along with the necessary services deployed with it. However, this process can be automated.
- Cluster-wide scaling and resource management can be more challenging, because changes might be necessary for each namespace or project.
- Change management at the cluster level still has potential impacts on all tenants.

- 部署密度较低，基础设施成本高于应用程序级租户，因为集群必须运行多个不共享的应用程序实例。在最简单的情况下，为每个租户运行一个完整的应用程序堆栈。
- 租户入职要求集群创建一个新项目或命名空间以及随其部署的必要服务。然而，这个过程可以自动化。
- 集群范围的扩展和资源管理可能更具挑战性，因为每个命名空间或项目都可能需要进行更改。
- 集群级别的变更管理仍然对所有租户产生潜在影响。

### Utilizing namespaces with Red Hat

### 在 Red Hat 中使用命名空间

[Red Hat OpenShift](http://developers.redhat.com/openshift) has a number of features that improve security and isolation beyond what is available in upstream Kubernetes. These controls provide additional protection for namespace-level tenancy:

[Red Hat OpenShift](http://developers.redhat.com/openshift) 有许多功能可以提高安全性和隔离性，超出了上游 Kubernetes 中可用的功能。这些控件为命名空间级别的租用提供了额外的保护：

- [Security-enhanced Linux (SELinux)](https://www.redhat.com/en/topics/linux/what-is-selinux) provides mandatory and discretionary access controls for files and processes at the operating system level, going beyond the default mechanisms in Linux.
- [Security context constraints](https://docs.openshift.com/container-platform/4.10/authentication/managing-security-context-constraints.html) limit the capabilities of individual containers.
- [Network-level isolation](https://docs.openshift.com/container-platform/4.10/networking/network_policy/multitenant-network-policy.html) prevents containers in one namespace from connecting to containers in other namespaces.

- [Security-enhanced Linux (SELinux)](https://www.redhat.com/en/topics/linux/what-is-selinux) 在操作系统级别为文件和进程提供强制和自由访问控制，去超出了 Linux 中的默认机制。
- [安全上下文约束](https://docs.openshift.com/container-platform/4.10/authentication/managing-security-context-constraints.html) 限制单个容器的能力。
- [网络级隔离](https://docs.openshift.com/container-platform/4.10/networking/network_policy/multitenant-network-policy.html) 防止一个命名空间中的容器连接到其他命名空间中的容器。

A future article will go into more detail on the security considerations for SaaS applications. For more information, please see:

未来的文章将更详细地介绍 SaaS 应用程序的安全注意事项。欲了解更多信息，请参阅：

- [Red Hat OpenShift documentation: Projects and users](https://docs.openshift.com/online/pro/architecture/core_concepts/projects_and_users.html)
- [Red Hat OpenShift security guide](https://www.redhat.com/en/resources/openshift-security-guide-ebook) 

- [红帽 OpenShift 文档：项目和用户](https://docs.openshift.com/online/pro/architecture/core_concepts/projects_and_users.html)
- [红帽 OpenShift 安全指南](https://www.redhat.com/en/resources/openshift-security-guide-ebook)

- [Kubernetes documentation: Namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)

- [Kubernetes 文档：命名空间](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)

## Using single-tenant clusters for strong isolation

## 使用单租户集群进行强隔离

To provide a very high level of isolation, a tenant can be assigned to a dedicated Kubernetes or Red Hat OpenShift cluster. In this case, no resources are shared between tenants—not even the cluster's control plane.

为了提供非常高的隔离级别，可以将租户分配到专用的 Kubernetes 或 Red Hat OpenShift 集群。在这种情况下，租户之间不共享任何资源——甚至集群的控制平面也不共享。

There are a number of cases where single-tenant clusters can be a good fit:

在许多情况下，单租户集群非常适合：

- A dedicated cluster can be a good option for a tenant that is very risk-averse and is willing to pay for a higher level of service.
- Single-tenant clusters can be used to address compliance requirements or legal regulations, particularly for highly regulated industries.
- The hardware the cluster runs on might need to be physically located within a specific national entity to comply with data sovereignty laws.
- Using single-tenant clusters allows cluster administrative access to be partitioned so only specific administrators have control over a given tenant's cluster.

- 对于非常规避风险并愿意为更高水平的服务付费的租户来说，专用集群可能是一个不错的选择。
- 单租户集群可用于解决合规要求或法律法规，特别是对于高度监管的行业。
- 运行集群的硬件可能需要物理上位于特定的国家实体内，以遵守数据主权法律。
- 使用单租户集群允许对集群管理访问进行分区，因此只有特定管理员才能控制给定租户的集群。

Figure 2 shows an overview of cluster-level tenancy.

图 2 显示了集群级租户的概览。

[![Two separate clusters, each containing not only its own application resources but infrastructure such as control planes.](http://developers.redhat.com/sites/default/files/styles/article_full_width_1440px_w/public/image2_3.png?itok=aoV2MSTx)](http://developers.redhat.com/sites/default/files/image2_3.png)

itok=aoV2MSTx)](http://developers.redhat.com/sites/default/files/image2_3.png)

Figure 2. Two separate clusters, each containing not only its own application resources but infrastructure such as control planes.

图 2. 两个独立的集群，每个集群不仅包含自己的应用程序资源，还包含控制平面等基础设施。

Figure 2: Two separate clusters, each containing not only its own application resources but infrastructure elements such as control planes.

图 2：两个独立的集群，每个集群不仅包含自己的应用程序资源，还包含控制平面等基础设施元素。

Single-tenant clusters provide a very high level of protection to mitigate the impact of security breaches. As discussed earlier, namespace-level tenancy limits any vulnerability within the application or application's components to the tenant using that namespace. However, parts of the technology stack are still shared when multiple tenants use the same worker node and control plane. A vulnerability in these shared parts of the stack could disclose or compromise the integrity of other tenant data. The shared components include:

单租户集群提供了非常高水平的保护，以减轻安全漏洞的影响。如前所述，命名空间级别的租赁将应用程序或应用程序组件中的任何漏洞限制为使用该命名空间的租户。但是，当多个租户使用相同的工作节点和控制平面时，部分技术堆栈仍然是共享的。堆栈的这些共享部分中的漏洞可能会泄露或损害其他租户数据的完整性。共享组件包括：

- **Container engine:** The container engine is responsible for configuring the isolation between containers and the host system. A bug in the container engine could allow a process running inside a container to break out.
- **Linux kernel and host operating system:** A vulnerability in the Linux kernel or operating system components could allow an intruder to bypass the system-level controls that isolate users, processes, or containers.
- **Hardware such as CPUs and memory:** A hardware defect could be exploited to access the memory of other processes running on the system. The [Meltdown and Spectre attacks](https://meltdownattack.com) against weaknesses in CPU design have shown that this risk is not theoretical and could be exploited to gain access to sensitive information such as encryption keys.
- **Control plane:** A bug in the implementation of the control plane could allow an intruder to bypass cluster isolation controls.

- **容器引擎：**容器引擎负责配置容器与宿主系统之间的隔离。容器引擎中的错误可能允许在容器内运行的进程中断。
- **Linux 内核和主机操作系统：** Linux 内核或操作系统组件中的漏洞可能允许入侵者绕过隔离用户、进程或容器的系统级控制。
- **硬件，例如 CPU 和内存：** 可以利用硬件缺陷来访问系统上运行的其他进程的内存。针对 CPU 设计弱点的 [Meltdown 和 Spectre 攻击](https://meltdownattack.com) 表明，这种风险不是理论上的，可以被利用来获取敏感信息，例如加密密钥。
- **控制平面：** 控制平面实现中的错误可能允许入侵者绕过集群隔离控制。

The good news that is vulnerabilities in these lower levels of the technology stack occur far less frequently than application-level vulnerabilities. The best practice to guard against security vulnerabilities is to use a defense-in-depth strategy, where an attack would have to breach multiple layers of protection.

好消息是，这些较低级别的技术堆栈中的漏洞发生的频率远低于应用程序级别的漏洞。防范安全漏洞的最佳实践是使用纵深防御策略，其中攻击必须破坏多层保护。

Single-tenant clusters are a very strong line of defense for situations that require this level of guaranteed security. A future article in this SaaS architecture checklist series will cover a number of the tools that can be used as part of a defense-in-depth security framework.

对于需要这种级别的安全保障的情况，单租户集群是一道非常坚固的防线。此 SaaS 架构清单系列中的后续文章将介绍一些可用作纵深防御安全框架一部分的工具。

### Advantages of single-tenant clusters

### 单租户集群的优势

This architecture is appealing for the following reasons:

这种架构之所以吸引人，原因如下：

- This design offers the highest level of isolation and security among the three approaches outlined in this article. Deployments do not share components of the technology stack, keeping even the container engine, Linux kernel, and control plane separate.
- This design also offers the highest degree of per-tenant customization. An application can run on completely separate hardware, in different clouds or data centers, and even in different countries. 

- 此设计在本文概述的三种方法中提供了最高级别的隔离和安全性。部署不共享技术堆栈的组件，甚至将容器引擎、Linux 内核和控制平面分开。
- 这种设计还提供了最高程度的每租户定制。一个应用程序可以在完全独立的硬件上运行，在不同的云或数据中心，甚至在不同的国家。

- Change management cadence and scheduling for cluster maintenance and upgrades can be tailored to each tenant's needs.

- 集群维护和升级的变更管理节奏和调度可以根据每个租户的需求进行定制。

### Disadvantages of single-tenant clusters

### 单租户集群的缺点

You might decide not to use this architecture for the following reasons:

您可能出于以下原因决定不使用此架构：

- Because there is no sharing between tenants, this design has the lowest tenant density and therefore the highest infrastructure costs among the three approaches.
- This design also involves the most cluster administration, because a cluster needs to be deployed and maintained for each tenant. However, automation and tools can significantly reduce the effort required.

- 由于租户之间没有共享，因此此设计的租户密度最低，因此在三种方法中基础设施成本最高。
- 这种设计也涉及最多的集群管理，因为需要为每个租户部署和维护一个集群。但是，自动化和工具可以显着减少所需的工作量。

### Utilizing Kubernetes clusters with Red Hat

### 在 Red Hat 中使用 Kubernetes 集群

There are a number of options to make the deployment and management of multiple Kubernetes or Red Hat OpenShift clusters faster and easier.

有许多选项可以更快、更轻松地部署和管理多个 Kubernetes 或 Red Hat OpenShift 集群。

[Red Hat Advanced Cluster Management for Kubernetes](https://www.redhat.com/en/technologies/management/advanced-cluster-management) allows you to manage multiple Kubernetes or Red Hat OpenShift clusters from a single management interface. You can use Red Hat Advanced Cluster Management for Kubernetes's unified multi-cluster lifecycle management to create, update, upgrade, and decommission clusters as needed for your SaaS deployments. Some of the managed services for Red Hat OpenShift on public clouds can manage multiple clusters.

[Red Hat Advanced Cluster Management for Kubernetes](https://www.redhat.com/en/technologies/management/advanced-cluster-management) 允许您从单个管理界面管理多个 Kubernetes 或 Red Hat OpenShift 集群。您可以使用 Red Hat Advanced Cluster Management for Kubernetes 的统一多集群生命周期管理来根据 SaaS 部署的需要创建、更新、升级和停用集群。公共云上的一些 Red Hat OpenShift 托管服务可以管理多个集群。

A feature planned for a future release of Red Hat Advanced Cluster Management for Kubernetes is _hosted control planes_. This feature allows the control plane for new clusters to be run on the worker nodes of an existing management cluster, which reduces the infrastructure and deployment time required. The upstream project for hosting OpenShift clusters at scale is called [HyperShift](https://cloud.redhat.com/blog/a-guide-to-red-hat-hypershift-on-bare-metal). A future article will discuss multi-cluster management.

Red Hat Advanced Cluster Management for Kubernetes 未来版本中计划的一项功能是_托管控制平面_。此功能允许新集群的控制平面在现有管理集群的工作节点上运行，从而减少所需的基础架构和部署时间。用于大规模托管 OpenShift 集群的上游项目称为 [HyperShift](https://cloud.redhat.com/blog/a-guide-to-red-hat-hypershift-on-bare-metal)。未来的文章将讨论多集群管理。

## Comparison of multi-tenancy approaches

## 多租户方法的比较

This article discussed three approaches to providing multi-tenancy for SaaS applications:

本文讨论了为 SaaS 应用程序提供多租户的三种方法：

- **Application-level tenancy:** Using application logic to isolate tenants.
- **Namespace-level tenancy:** Providing isolation between tenant application instances through Kubernetes namespaces or OpenShift projects.
- **Cluster-level tenancy:** Using a dedicated Kubernetes cluster for each tenant to provide a very high level of isolation and security.

- **应用程序级租户：**使用应用程序逻辑隔离租户。
- **命名空间级租户：** 通过 Kubernetes 命名空间或 OpenShift 项目在租户应用程序实例之间提供隔离。
- **集群级租户：**为每个租户使用专用的 Kubernetes 集群，以提供非常高级别的隔离和安全性。

The first approach requires the application to be written for multi-tenancy. The latter two approaches can be used to deploy existing single-tenant applications as SaaS applications. Table 1 compares the three approaches.

第一种方法需要为多租户编写应用程序。后两种方法可用于将现有的单租户应用程序部署为 SaaS 应用程序。表 1 比较了这三种方法。

Table 1: Comparison of multi-tenancy approachesApplication-level

表 1：多租户方法的比较应用程序级

tenancyNamespace-level tenancyCluster-level tenancyTenant isolation enforced byApplication logicSeparate namespacesSeparate clustersLevel of infrastructure shared across tenantsHigh-allMedium-mixedLow-noneTenant density possibleHighMediumLowInfrastructure costs per tenantLowMediumHighSecurity protectionsLowMediumHighAbility to monitor and scale per tenantLowMediumHighAbility to customize deployment for each tenantLowMediumHigh

租户命名空间级租户集群级租户由应用程序逻辑强制执行的租户隔离单独的命名空间单独的集群跨租户共享的基础设施级别高-全部中混合低-无租户密度可能高中低每个租户的基础设施成本低中高安全保护低中高能够监控和扩展每个租户低中高能够为每个租户自定义部署低中高

These three approaches are not mutually exclusive and can be combined to provide the best fit for your application, or to address the requirements of specific tenants. For example, an application that implements multi-tenancy could be deployed as multiple instances with either namespace-level or cluster-level tenancy to provide enhanced security protections and to contain the potential impact of vulnerabilities to a limited number of tenants.

这三种方法不是相互排斥的，可以组合使用以提供最适合您的应用程序的方法，或满足特定租户的要求。例如，实现多租户的应用程序可以部署为具有命名空间级别或集群级别租户的多个实例，以提供增强的安全保护并遏制漏洞对有限数量租户的潜在影响。

A SaaS provider could offer multiple levels of service to address the requirements of tenants that are more risk-averse and are willing to pay for higher levels of service, versus those that are more cost-conscious and less sensitive to risk. Cost-conscious tenants can be served by environments with high degrees of sharing. The most risk-sensitive tenants could be offered the option of dedicated hardware to meet their requirements. 

SaaS 提供商可以提供多层次的服务，以满足那些更厌恶风险并愿意为更高水平的服务付费的租户的需求，而不是那些更注重成本且对风险不太敏感的租户的需求。高度共享的环境可以为注重成本的租户提供服务。可以为对风险最敏感的租户提供专用硬件选项来满足他们的要求。

Future articles in this series will cover topics such as security and the options available for a defense-in-depth strategy. Another article will address single-tenant workloads that are difficult to containerize and will explain how these can be deployed for SaaS use as virtual machines using Red Hat OpenShift Virtualization with namespace-level isolation.

本系列的后续文章将涵盖安全性和可用于纵深防御策略的选项等主题。另一篇文章将讨论难以容器化的单租户工作负载，并将解释如何使用具有命名空间级别隔离的 Red Hat OpenShift Virtualization 将这些工作负载部署为 SaaS 用作虚拟机。

## Partner with Red Hat to build your SaaS

## 与红帽合作构建您的 SaaS

[Red Hat SaaS Foundations](https://connect.redhat.com/en/partner-with-us/red-hat-saas-foundations) is a partner program designed for building enterprise-grade SaaS solutions on Red Hat OpenShift or Red Hat Enterprise Linux platforms, and deploying them across multiple cloud and non-cloud footprints. [Email](http://mailto:saas@redhat.com) us to learn more.

[Red Hat SaaS Foundations](https://connect.redhat.com/en/partner-with-us/red-hat-saas-foundations) 是一个合作伙伴计划，旨在在 Red Hat OpenShift 或红帽企业 Linux 平台，并将它们部署到多个云和非云足迹。 [电子邮件](http://mailto:saas@redhat.com) 我们以了解更多信息。

_Last updated:
June 2, 2022_ 

_最近更新时间：
2022年6月2日_

