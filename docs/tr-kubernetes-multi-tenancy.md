# Understanding Multi-Tenancy in Kubernetes

# 了解 Kubernetes 中的多租户

https://www.thinksys.com/devops/kubernetes-multi-tenancy/

A [**container orchestration**](http://www.thinksys.com/devops/all-about-containers/) system is one of the most widely used automating software scaling, deployment, and management. Created by Google and managed by **Cloud Native Computing** Foundation, **[Kubernetes](http://www.thinksys.com/devops/understanding-kubernetes-architecture/)** is an **open-source container orchestration system** used by legions of organizations. Container orchestration helps in operational tasks like networking, provisioning, and deploying in containerized workloads. Several organizations can use multiple workloads in a **single Kubernetes cluster** that shares the same infrastructure. This strategy used by enterprises is called multi-tenancy. This article is all about **understanding multi-tenancy in Kubernetes**, including its use cases, best practices, and how it helps organizations in the cloud-native space.

[**容器编排**](http://www.thinksys.com/devops/all-about-containers/) 系统是使用最广泛的自动化软件扩展、部署和管理系统之一。由 Google 创建并由 **Cloud Native Computing** Foundation 管理，**[Kubernetes](http://www.thinksys.com/devops/understanding-kubernetes-architecture/)** 是一个**开源容器大量组织使用的编排系统**。容器编排有助于在容器化工作负载中完成网络、配置和部署等操作任务。多个组织可以在共享相同基础架构的**单个 Kubernetes 集群**中使用多个工作负载。企业使用的这种策略称为多租户。本文全部是关于**了解 Kubernetes 中的多租户**，包括它的用例、最佳实践以及它如何帮助云原生空间中的组织。

## **What is a Multi-Tenancy in Kubernetes?**

## **什么是 Kubernetes 中的多租户？**

Tenants are separate entities in an organization but share several shared components like infrastructure. Every tenant in an organization is given some shared components along with some isolation. Achieving this isolation can be done in multiple ways, like giving every tenant a unique server or virtual machine. Even though it is effective, the efficiency is compromised here as it will be costly and will not use the resources the right way. This is the part where **multi-tenancy in Kubernetes** comes into play.

租户是组织中的独立实体，但共享多个共享组件，例如基础架构。组织中的每个租户都有一些共享组件以及一些隔离。可以通过多种方式实现这种隔离，例如为每个租户提供唯一的服务器或虚拟机。即使它是有效的，这里的效率也会受到影响，因为它成本高昂并且不会以正确的方式使用资源。这是 **Kubernetes 中的多租户** 发挥作用的部分。

![Kubernetes Multi Tenancy ](https://www.thinksys.com/wp-content/uploads/2022/06/MultiTanancy-Kubernates.jpg)

**[Multi-tenancy](http://www.thinksys.com/cloud/multi-tenant-architecture-cloud-apps/)** is the capability to run different entities' workloads in a single cluster shared by different tenants . The organizations running multi-tenant clusters have to isolate each tenant to avoid any potential damage caused by a malicious tenant. In simpler terms, multi-tenancy is allocating an isolated space to an entity in a cluster while also giving some shared cluster components. This model is used in organizations running several applications in a single environment or where different teams exist but share a single **Kubernetes environment**.

**[多租户](http://www.thinksys.com/cloud/multi-tenant-architecture-cloud-apps/)** 是在由不同租户共享的单个集群中运行不同实体的工作负载的能力.运行多租户集群的组织必须隔离每个租户，以避免恶意租户造成的任何潜在损害。简单来说，多租户就是为集群中的一个实体分配一个孤立的空间，同时也提供一些共享的集群组件。此模型用于在单个环境中运行多个应用程序的组织，或者存在不同团队但共享单个 **Kubernetes 环境**的组织。

## **Types of Multi-Tenancy in Kubernetes**

## **Kubernetes 中的多租户类型**

- **Soft Multi-Tenancy :** Soft multi-tenancy is mainly done when numerous projects or departments are running in an organization or when there are trusted tenants. It can be implemented through the Kubernetes namespace multi-tenancy. In this type of multi-tenancy, the Kubernetes cluster is separated among different users but without extreme or strict isolation. The primary reason why this type of multi-tenancy is implemented in the Kubernetes cluster is to assist with resource separation and avoid any accidental access to resources. As isolation is not strict in this type, deliberate attacks by one tenant on another cannot be prevented or minimized. This type is only preferred for trusted tenants in a Kubernetes environment with that thing in mind.
- **Hard Multi-Tenancy :** Organizations having legions of tenants in a single Kubernetes cluster may have both trusted and untrusted tenants. Implementing hard multi-tenancy in Kubernetes applies much stricter isolation than soft mule-tenancy, hindering tenants from influencing each other. Even the malicious tenants cannot affect any other tenants in the cluster. As this multi-tenant type comes with stricter isolation, enforcing it is also more complicated. The virtual cluster and namespace configuration is the tricky part of its implementation.

- **软多租户：** 软多租户主要在组织中运行多个项目或部门或有受信任的租户时完成。它可以通过 Kubernetes 命名空间多租户来实现。在这种类型的多租户中，Kubernetes 集群在不同的用户之间是分开的，但没有极端或严格的隔离。在 Kubernetes 集群中实现这种类型的多租户的主要原因是为了协助资源分离并避免任何意外访问资源。由于这种类型的隔离并不严格，因此无法防止或最小化一个租户对另一个租户的蓄意攻击。考虑到这一点，这种类型仅适用于 Kubernetes 环境中受信任的租户。
- **硬多租户：** 在单个 Kubernetes 集群中拥有大量租户的组织可能同时拥有受信任和不受信任的租户。在 Kubernetes 中实现硬多租户比软 mule 租户更严格的隔离，阻碍租户相互影响。即使是恶意租户也无法影响集群中的任何其他租户。由于这种多租户类型具有更严格的隔离，因此执行起来也更加复杂。虚拟集群和命名空间配置是其实现的棘手部分。

## **Single Tenancy VS Multi-Tenancy(How Multi Tenancy is different from Single tenancy?)**

## **单租户 VS 多租户（多租户与单租户有何不同？）**

Both **single and multi-tenancy** are different from each other in many ways. Before you make your next move, it is crucial to understand the key differences each tenancy type will bring. 

**单租户和多租户**在很多方面都彼此不同。在您采取下一步行动之前，了解每种租赁类型将带来的主要差异至关重要。

- **Cost :** Every tenant will have a unique cluster with a control plane, master nodes, monitoring, and other management components in a single tenancy. Implementing and managing a separate cluster with management components for each tenant is highly expensive, especially in a large organization with tons of tenants. On the other hand, multi-tenancy gives an isolated space to each tenant but on a single cluster where the organization can reuse resources, making it a cost-effective option.

- **成本：** 每个租户都将拥有一个独特的集群，在单个租户中包含一个控制平面、主节点、监控和其他管理组件。为每个租户实施和管理带有管理组件的单独集群非常昂贵，尤其是在拥有大量租户的大型组织中。另一方面，多租户为每个租户提供了一个独立的空间，但在组织可以重用资源的单个集群上，使其成为一种具有成本效益的选择。

- **Complexity:** Having a single cluster for each tenant is not just time-consuming but also complex. Even though the process can be automated through managed services, still it should be done for each cluster. A multi-tenancy method is less complicated as there is a need for setting up only one cluster for multiple tenants.

- **复杂性：** 为每个租户拥有一个集群不仅耗时而且复杂。尽管该过程可以通过托管服务自动化，但仍应为每个集群完成。多租户方法不太复杂，因为只需为多个租户设置一个集群。

- **Security:** As there is a single cluster for every tenant, their isolation is natural; hence, none of the malicious tenants can affect any other. Due to this reason, a single tenancy is preferred for untrusted tenants. However, Kubernetes multi-tenancy security management can be a tussle as tenants share a single cluster. Though the security can be enhanced with Role-Based Access Control and hard multi-tenancy, still it is extra work that requires additional time and effort.

- **安全性：**由于每个租户都有一个集群，因此它们的隔离是自然的；因此，任何恶意租户都不会影响其他任何租户。由于这个原因，不受信任的租户首选单一租赁。然而，Kubernetes 多租户安全管理可能是一场争斗，因为租户共享一个集群。尽管可以通过基于角色的访问控制和硬多租户来增强安全性，但它仍然是需要额外时间和精力的额外工作。

## **Multi-Tenancy Models in Kubernetes:**

## **Kubernetes 中的多租户模型：**

Kubernetes multi-tenancy models help in making its use cases easier and uncomplicated. Depending on the organizations and teams, these models can be implemented for the best outcome. Three most commonly implemented multi-tenancy models can be applied to projects.

Kubernetes 多租户模型有助于使其用例更容易和简单。根据组织和团队的不同，可以实施这些模型以获得最佳结果。三种最常用的多租户模型可以应用于项目。

1. **Namespaces-as-a-Service:** In this model, every tenant shares a cluster where their workloads are restricted only to a set of namespaces allocated to the particular tenant. However, all the control plane resources, including the scheduler and AP server, CPU, and memory, are accessible by all the tenants across the cluster. The namespaces-as-a-service model allows tenants to share all the cluster resources, hindering clusters from updating or creating any of such resources. When isolating tenant workloads, their namespace should contain role bindings, resource quota, and network policies. Adding these to the namespaces is necessary as they help control access to the namespace, limit usage in the tenants, and prevent network traffic in all the tenants.
2. **Clusters as a Service:** In clusters as a service multi-tenancy model, every tenant is given a cluster where they can use cluster-wide resources. Moreover, they can have a Kubernetes control plane with complete isolation where management cluster projects are used to provision multiple workload clusters. Furthermore, the tenants are provided with a workload cluster that provides complete control of the cluster resources. The central platform teams manage add-on services like security, monitoring, upgrading, patching, and cluster lifestyle management services. However, certain limitations exist for the tenant admin to modify the services above.
3. **Control Planes as a Service:** The control planes as a service variant of the earlier mentioned CaaS model. However, a tenant may be assigned a virtual cluster where they will be given an exclusive control plane. This model is applicable when the virtual cluster’s users cannot determine the differences between a Kubernetes cluster and a virtual cluster. Even though a virtual cluster is allocated to tenants, they still have to share worker node resources and certain control plane components. This model is implemented by a virtual cluster project where several VCs share a super-cluster.



1. **命名空间即服务：** 在此模型中，每个租户共享一个集群，其中他们的工作负载仅限于分配给特定租户的一组命名空间。但是，集群中的所有租户都可以访问所有控制平面资源，包括调度程序和 AP 服务器、CPU 和内存。命名空间即服务模型允许租户共享所有集群资源，阻止集群更新或创建任何此类资源。隔离租户工作负载时，其命名空间应包含角色绑定、资源配额和网络策略。将这些添加到命名空间是必要的，因为它们有助于控制对命名空间的访问，限制租户的使用，并防止所有租户的网络流量。
2. **集群即服务：** 在集群即服务多租户模型中，每个租户都被分配了一个集群，他们可以在其中使用集群范围的资源。此外，他们可以拥有一个完全隔离的 Kubernetes 控制平面，其中管理集群项目用于配置多个工作负载集群。此外，为租户提供了一个工作负载集群，该集群提供对集群资源的完全控制。中央平台团队管理附加服务，如安全、监控、升级、修补和集群生活方式管理服务。但是，租户管理员修改上述服务存在某些限制。
3. **控制平面即服务：** 控制平面即前面提到的 CaaS 模型的服务变体。但是，可以为租户分配一个虚拟集群，在那里他们将获得一个独占的控制平面。该模型适用于虚拟集群的用户无法确定 Kubernetes 集群和虚拟集群之间的差异的情况。即使将虚拟集群分配给租户，他们仍然必须共享工作节点资源和某些控制平面组件。该模型由多个 VC 共享一个超级集群的虚拟集群项目实现。

## **Kubernetes Multi-Tenancy Use Cases**

## **Kubernetes 多租户用例**

Here are the most frequently use cases of the **Kubernetes multi-tenancy models**: 

以下是 **Kubernetes 多租户模型**最常见的用例：

1. **SaaS Provider Multi-Tenancy:** The Software-as-a-Service control plane and the customer instance are the tenants of a SaaS provider’s cluster. Every application instance is organized with its namespaces and the SaaS control plane components to take full leverage of the namespace policies. Every end-user has to use the interface provided by SaaS, which communicates with the Kubernetes control plane. This process has to be followed by every user as they cannot communicate with the control plane directly. The biggest example of SaaS provider multi-tenancy is a blogging platform running on a multi-tenant cluster. Here, the platform gives them a control plane, and their user’s blog will have a separate namespace. The users can use all the services through its interface without viewing the operation of the cluster.
2. **Enterprise Multi-Tenancy:** When it comes to tenants in an enterprise, they are mainly different teams of the same organization that comes with a namespace. However, managing these tenants per cluster in an alternative multi-tenancy model is complicated. Furthermore, the network traffic between tenants should be defined correctly. This task in multi-tenancy can be accomplished through Kubernetes network policy where the cluster users can be categorized into cluster-admin, namespace admin, and developer.



1. **SaaS 提供商多租户：** 软件即服务控制平面和客户实例是 SaaS 提供商集群的租户。每个应用程序实例都使用其命名空间和 SaaS 控制平面组件进行组织，以充分利用命名空间策略。每个最终用户都必须使用 SaaS 提供的接口，该接口与 Kubernetes 控制平面进行通信。每个用户都必须遵循此过程，因为他们无法直接与控制平面通信。 SaaS 提供商多租户的最大例子是在多租户集群上运行的博客平台。在这里，平台为他们提供了一个控制平面，他们的用户的博客将有一个单独的命名空间。用户可以通过其界面使用所有服务，而无需查看集群的运行情况。
2. **Enterprise Multi-Tenancy：**企业中的租户，主要是同一个组织的不同团队，带有一个命名空间。但是，在另一种多租户模型中管理每个集群的这些租户很复杂。此外，应正确定义租户之间的网络流量。多租户中的这项任务可以通过 Kubernetes 网络策略来完成，其中集群用户可以分为集群管理员、命名空间管理员和开发人员。

A cluster admin will handle the cluster and its tenants with authority to create, update, delete, and read any policy object. Moreover, they can also create and assign namespaces. The next is the namespace administrator, who will handle single tenants in the namespace. The last role is of the developer who can read, delete, create, and update namespace non-policy objects. However, this role is limited because their authority lies within their accessible namespaces.
3. **Multiple Applications on a Single Cluster:** There are stances when organizations want to host multiple applications on a single cluster. This need can be fulfilled through multi-tenancy, where you can host several related and unrelated applications that require a scalable platform but in a single cluster.
4. **Hosting Trusted and Untrusted Tenants:** An organization has to work with both trusted and untrusted tenants that may be malicious. Without a doubt, the organization never wants to compromise tenants’ security. This is the part where the multi-tenancy cluster comes into play. Here, the organization can share the infrastructure with both types of tenants without worrying about security. They can host apps needed by the internal teams along with the external entities that may require access to your cluster for workloads.

集群管理员将处理集群及其租户，并有权创建、更新、删除和读取任何策略对象。此外，他们还可以创建和分配命名空间。接下来是命名空间管理员，他将处理命名空间中的单个租户。最后一个角色是可以读取、删除、创建和更新命名空间非策略对象的开发人员。但是，此角色是有限的，因为他们的权限位于其可访问的名称空间内。
3. **单个集群上的多个应用程序：** 当组织希望在单个集群上托管多个应用程序时，会有一些立场。这种需求可以通过多租户来满足，您可以在其中托管多个需要可扩展平台但在单个集群中的相关和不相关的应用程序。
4. **托管受信任和不受信任的租户：** 组织必须与可能存在恶意的受信任和不受信任的租户合作。毫无疑问，该组织从不想损害租户的安全。这是多租户集群发挥作用的部分。在这里，组织可以与两种类型的租户共享基础架构，而无需担心安全性。他们可以托管内部团队所需的应用程序以及可能需要访问集群以进行工作负载的外部实体。

## **Best Practices for Kubernetes Multi-Tenancy**

## **Kubernetes 多租户的最佳实践**

Kubernetes multi-tenancy can be used for many different use cases. However, the right practices must be followed to get the most out of it. With that in mind, here are some of the best practices for Kubernetes multi-tenancy.

Kubernetes 多租户可用于许多不同的用例。但是，必须遵循正确的做法才能充分利用它。考虑到这一点，这里有一些 Kubernetes 多租户的最佳实践。

- **Cluster Personas:** Creating a hierarchy of personas will maintain transparency within the process and avoid clashes in the team. Considering that fact, it is best to create a hierarchy of cluster personas based on the actions they can perform and the permissions they require to accomplish their tasks. There are four different personas in a multi-tenancy environment; cluster view, cluster-admin, tenant admin, and tenant user. 

- **集群角色：**创建角色层次结构将保持流程的透明度并避免团队中的冲突。考虑到这一事实，最好根据他们可以执行的操作和完成任务所需的权限来创建集群角色的层次结构。多租户环境中有四种不同的角色；集群视图、集群管理员、租户管理员和租户用户。

- **Role-Based Access Control (RBAC):** No matter from where the request originates, every create, read, delete, and update operation is done through the Kubernetes API server in multi-tenancy. When there are multiple tenants in a cluster, it is essential to keep it as secure as possible. Enabling Kubernetes multi-tenancy RBAC in an API server will help get better control over the applications and users in the cluster. Every RBAC has four API objects: Role, RoleBinding, ClusterRole, and ClusterRoleBinding. Furthermore, disabling attribute-based access control is also recommended.
- **Namespace Categorization:** In a multi-tenant Kubernetes environment, the namespace is among the most crucial aspects. One of the best practices regarding namespace is categorizing it into different groups. The most commonly used groups are system, service, and tenant namespace. The system namespace is exclusive for system pods, whereas the service namespace should run apps whose access is required by other namespaces in the cluster. However, a tenant namespace is a group used to run services and applications which do not need access from any other namespace.
- **Label Namespaces:** Another excellent yet underrated practice in multi-tenancy is labeling namespaces. This practice helps metadata applications to understand the reason for using resources. Labeling the namespaces will help understand the metrics whenever necessary or filter the application’s data easily.
- **Use Network Policy:** In a multi-tenant environment, it is essential to isolate tenant namespaces. This can be done by using network policies that let the cluster admins control the communication of group pods. Admins should use network policy resources to isolate tenant namespaces.
- **Limit Shared Resource Usage:** When multiple tenants exist in a cluster, they are bound to use shared resources. Sometimes, these resources can be wasted by a tenant, reducing the outcome for other tenants. A great way to eradicate this issue is by limiting the shared resource usage by implementing Kubernetes namespace resource quotas. Through this quota, you can control the total resource usage by a single tenant.



- **基于角色的访问控制 (RBAC)：** 无论请求来自何处，每次创建、读取、删除和更新操作都是通过多租户的 Kubernetes API 服务器完成的。当集群中有多个租户时，必须尽可能保证其安全。在 API 服务器中启用 Kubernetes 多租户 RBAC 将有助于更好地控制集群中的应用程序和用户。每个 RBAC 都有四个 API 对象：Role、RoleBinding、ClusterRole 和 ClusterRoleBinding。此外，还建议禁用基于属性的访问控制。
- **命名空间分类：** 在多租户 Kubernetes 环境中，命名空间是最关键的方面之一。关于命名空间的最佳实践之一是将其分类为不同的组。最常用的组是系统、服务和租户命名空间。系统命名空间是系统 pod 专有的，而服务命名空间应该运行集群中其他命名空间需要访问的应用程序。但是，租户命名空间是用于运行不需要从任何其他命名空间访问的服务和应用程序的组。
- **标签命名空间：** 多租户中另一个优秀但被低估的做法是标签命名空间。这种做法有助于元数据应用程序了解使用资源的原因。标记命名空间将有助于在必要时了解指标或轻松过滤应用程序的数据。
- **使用网络策略：** 在多租户环境中，隔离租户命名空间是必不可少的。这可以通过使用允许集群管理员控制组 pod 通信的网络策略来完成。管理员应该使用网络策略资源来隔离租户命名空间。
- **限制共享资源使用：** 当集群中存在多个租户时，它们必然会使用共享资源。有时，这些资源可能会被租户浪费，从而减少其他租户的结果。消除此问题的一个好方法是通过实施 Kubernetes 命名空间资源配额来限制共享资源的使用。通过此配额，您可以控制单个租户的总资源使用量。

## **Kubernetes Multi-Tenancy Cluster Setup**

## **Kubernetes 多租户集群设置**

Experts recommend having a single large cluster in a multi-tenancy environment rather than having multiple small clusters for different tenants. Here is a quick guide on setting up clusters in a multi-tenancy environment.

专家建议在多租户环境中使用单个大型集群，而不是为不同租户使用多个小型集群。这是在多租户环境中设置集群的快速指南。

### **Step 1: Partition cluster depending on the workload using namespaces:**

### **步骤 1：使用命名空间根据工作负载对集群进行分区：**

The first step is to set up a cluster based on the development workload requirement. The basic cluster configurations come with four nodes with a single CPU and four gigabytes of RAM. This process includes setting up two teams in a cluster with a separate namespace for each team. Though the Kubernetes cluster comes with a default namespace, new namespaces can be created for the teams. For instance, the following namespace can be created for the leadership and virtual teams.

第一步是根据开发工作负载需求设置集群。基本集群配置带有四个节点、一个 CPU 和 4 GB 的 RAM。此过程包括在集群中设置两个团队，每个团队都有一个单独的命名空间。尽管 Kubernetes 集群带有默认命名空间，但可以为团队创建新的命名空间。例如，可以为领导团队和虚拟团队创建以下命名空间。

> _kubectl create namespace team-leadership_
>
> _kubectl create namespace team-virtual_



Furthermore, creating a sample application in these namespaces is also part of this step. This process will guide you in deploying an Nginx pod in the created namespaces using the below-mentioned command.

### **Step 2: Grant access control to the teams:**

### **第 2 步：授予团队访问控制权：**

Once the namespaces and the applications are ready, it is time to provide them access control. To accomplish this task, the first thing to do is to create a service account for the team and assign the IAM role. Once that is done, you need to create a Kubernetes role with basic CRUD permissions. Moreover, you need to assign this role to the IAM service account created earlier in the process.

一旦命名空间和应用程序准备就绪，就该为它们提供访问控制了。要完成此任务，首先要做的是为团队创建一个服务帐户并分配 IAM 角色。完成后，您需要创建一个具有基本 CRUD 权限的 Kubernetes 角色。此外，您需要将此角色分配给在该流程之前创建的 IAM 服务账户。

### **Step 3: Test the Access:** 

### **第 3 步：测试访问权限：**

Now that the roles are assigned, the right practice is to test the access. To do that, you need to download the JSON key of the service account and try to log in. After you have successfully logged in to the service account, make sure to access the app in the namespace. Follow the same process for every namespace that you have created.

现在已经分配了角色，正确的做法是测试访问权限。为此，您需要下载服务帐户的 JSON 密钥并尝试登录。成功登录服务帐户后，请确保访问命名空间中的应用程序。对您创建的每个命名空间都遵循相同的过程。

### **Step 4: Assign Resources to the namespace:**

### **第 4 步：将资源分配给命名空间：**

When there are multiple tenants or namespaces in a single cluster, they share some resources, so resource allocation is necessary. This allocation can be achieved by using ResourceQuota Kubernetes, which you will configure for namespaces regarding the resources they can utilize like total storage space, CPU, memory, pods, and services. You can restrict the resource utilization in each namespace so that none of the tenants overutilize resources.

当单个集群中有多个租户或命名空间时，它们共享一些资源，因此需要进行资源分配。这种分配可以通过使用 ResourceQuota Kubernetes 来实现，您将为命名空间配置它们可以利用的资源，例如总存储空间、CPU、内存、pod 和服务。您可以限制每个命名空间中的资源利用率，以免租户过度使用资源。

### **Step 5: Resource Utilization Monitoring:**

### **第 5 步：资源利用率监控：**

As you have already allocated resources to each namespace in the previous step, it is time to watch resource utilization. When new use cases are added to the namespace, the resource utilization may change. With that in mind, it is always advised to understand the resource usage pattern to make sure that every namespace gets the right amount of resources as per their usage.

由于您已经在上一步中为每个命名空间分配了资源，是时候观察资源利用率了。当新的用例添加到命名空间时，资源利用率可能会发生变化。考虑到这一点，始终建议您了解资源使用模式，以确保每个命名空间都能根据其使用情况获得正确数量的资源。

## What is Hypernetes in Kubernetes Multi-Tenancy?

## Kubernetes 多租户中的 Hypernetes 是什么？

Often, organizations run containers inside a VM to enhance its security. Undeniably, it is an effective method, but it comes with certain issues like the inability to manage container networks uniformly through the IaaS layer or the lack of centrally scheduled resources in containers. In that case, the alternative to this method is Hypernetes. A Kubernetes multi-tenant Distro adds a Hyper-based container execution engine, a container SDN network, cinder-based persistent storage, authentication, and authorization to the Kubernetes.

通常，组织在 VM 内运行容器以增强其安全性。不可否认，这是一种有效的方法，但它也存在一些问题，例如无法通过 IaaS 层统一管理容器网络，或者容器中缺乏集中调度的资源。在这种情况下，此方法的替代方法是 Hypernetes。 Kubernetes 多租户 Distro 为 Kubernetes 添加了基于 Hyper 的容器执行引擎、容器 SDN 网络、基于 cinder 的持久存储、身份验证和授权。

Furthermore, Hypernetes adds certain components to Kubernetes like isolated tenants managed by keystone, Layer 2 isolation network for tenants, isolation of containers through virtualization-based container execution engine, and persistent storage. Apart from that, Hypernetes provides numerous components based on Kubernetes through different plugins.

此外，Hypernetes 还为 Kubernetes 添加了某些组件，例如由 keystone 管理的隔离租户、租户的第 2 层隔离网络、通过基于虚拟化的容器执行引擎隔离容器以及持久存储。除此之外，Hypernetes 通过不同的插件提供了许多基于 Kubernetes 的组件。

### **Conclusion:**

###  **结论：**

Undoubtedly, using multiple clusters for each tenant is not a practical way of containerizing applications. The **Kubernetes multi-tenancy** has been proven to be an efficient way of storing applications. It is cost-effective but saves container setup time and a lot of resources as well. As multi-tenancy does not come out of the box in Kubernetes, organizations may require professional assistance. [ThinkSys Inc](http://www.thinksys.com/get-in-touch/) can provide you with the unique strategies to implement multi-tenancy in Kubernetes that will expand its overall usability and attain efficient resource utilization. Furthermore, ThinkSys’ dedicated Kubernetes toolset ensures effective multi-tenancy implementation in a cluster.

毫无疑问，为每个租户使用多个集群并不是容器化应用程序的实用方法。 **Kubernetes 多租户**已被证明是一种存储应用程序的有效方式。它具有成本效益，但也节省了容器设置时间和大量资源。由于多租户在 Kubernetes 中不是开箱即用的，因此组织可能需要专业帮助。 [ThinkSys Inc](http://www.thinksys.com/get-in-touch/) 可以为您提供在 Kubernetes 中实现多租户的独特策略，这将扩大其整体可用性并实现高效的资源利用。此外，ThinkSys 的专用 Kubernetes 工具集可确保在集群中有效地实施多租户。

[Chat with ThinkSys Kubernetes Experts to Implement multi-tenancy in Kubernetes](http://www.thinksys.com/get-in-touch/)

 [与ThinkSys Kubernetes专家畅谈在Kubernetes中实现多租户](http://www.thinksys.com/get-in-touch/)

## **Frequently Asked Questions**(Kubernetes)

## **常见问题**（Kubernetes）

### What is a Multi-Tenant SaaS platform?

### 什么是多租户 SaaS 平台？

A multi-tenant SaaS platform is software that serves several customers on a single infrastructure and database. However, every customer’s data is always isolated, which other tenants cannot access. They share all the shared resources but not their data.

多租户 SaaS 平台是在单个基础架构和数据库上为多个客户提供服务的软件。但是，每个客户的数据始终是孤立的，其他租户无法访问。他们共享所有共享资源，但不共享数据。

### What is a Multi-Tenant Schema?

### 什么是多租户架构？

A multi-tenant schema is when the application determines which schema to connect to for a tenant after connecting to a database.

多租户架构是指应用程序在连接到数据库后确定租户连接到哪个架构。

### What is Kubernetes Multi-tenancy?

### 什么是 Kubernetes 多租户？

A Kubernetes multi-tenancy is an architecture that helps run workloads of different entities in a single cluster but with isolation. Here, the workloads are also called tenants, which share the same cluster and its resources but are kept separate.

Kubernetes 多租户是一种有助于在单个集群中运行不同实体的工作负载但具有隔离性的架构。在这里，工作负载也称为租户，它们共享相同的集群及其资源，但保持独立。

### What do you mean by Multi-Tenancy?

### 多租户是什么意思？

Multi-tenancy is when a single cluster serves multiple tenants rather than creating a separate cluster for each tenant. Every tenant shares the cluster along with the database. However, their data is always isolated.

多租户是指单个集群为多个租户提供服务，而不是为每个租户创建单独的集群。每个租户与数据库共享集群。但是，他们的数据始终是孤立的。

### Is AWS Multi-Tenant? 

### 是 AWS 多租户吗？

The AWS supports multi-tenancy where SaaS applications can have multiple tenants with isolation. The level of isolation in Kubernetes multi-tenancy AWS and the shared resources is influenced by factors like domain nature, AWS services, and the multi-architecture model.

AWS 支持多租户，其中 SaaS 应用程序可以有多个隔离的租户。 Kubernetes 多租户 AWS 和共享资源的隔离级别受域性质、AWS 服务和多架构模型等因素的影响。

### Can Kubernetes run on multiple machines?

### Kubernetes 可以在多台机器上运行吗？

Kubernetes has the ability that allows containers to run on several machines, be it physical, on-premises, virtual, or cloud. Moreover, these containers support all the major operating systems and share the same to run on multiple machines.

Kubernetes 具有允许容器在多台机器上运行的能力，无论是物理的、本地的、虚拟的还是云的。此外，这些容器支持所有主要操作系统并共享相同的操作系统以在多台机器上运行。

### What is the difference between Single tenants and Multi-tenant?

### 单租户和多租户有什么区别？

The most significant difference between a single tenant and a multi-tenant is that the former will provide a separate database to a customer. In contrast, a multi-tenant can serve multiple customers with a single database. Multi-tenancy is proven to be cost-effective and resource-efficient for large organizations.

单租户和多租户之间最显着的区别是前者将为客户提供单独的数据库。相比之下，多租户可以使用单个数据库为多个客户提供服务。事实证明，对于大型组织而言，多租户具有成本效益和资源效率。

### What is Multi-tenant deployment?

### 什么是多租户部署？

A Kubernetes multi-tenant deployment is when multiple software instances run on a single cluster. This cluster will have multiple tenants who will share the resources while being in isolation simultaneously.

Kubernetes 多租户部署是指多个软件实例在单个集群上运行。该集群将有多个租户，他们将在隔离的同时共享资源。

### What is a Multi-Tenant Cluster?

### 什么是多租户集群？

Several customers share a multi-tenant cluster called tenants. The cluster operators will isolate the tenants where they will allocate resources for each tenant depending on the requirements. 

多个客户共享一个称为租户的多租户集群。集群操作员将隔离租户，他们将根据需求为每个租户分配资源。

