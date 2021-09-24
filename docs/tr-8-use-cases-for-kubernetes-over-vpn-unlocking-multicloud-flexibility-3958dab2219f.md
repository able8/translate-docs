# 8 Use Cases for Kubernetes over VPN: Unlocking Multicloud Flexibility

# 8 Kubernetes over VPN 用例：解锁多云灵活性

[Jun 29 2021](http://itnext.io/8-use-cases-for-kubernetes-over-vpn-unlocking-multicloud-flexibility-3958dab2219f?source=post_page-----3958dab2219f--------------------------------)·8 min read

This is a higher-level take on another story I wrote about [how to run Kubernetes across multiple clouds](http://itnext.io/how-to-deploy-a-single-kubernetes-cluster-across-multiple-clouds-using-k3s-and-wireguard-a5ae176a6e81). Before explaining _how_, I should have answered _why?_ So, let’s start with the problem.

这是对我写的另一个故事的更高层次的看法 [如何跨多个云运行 Kubernetes](http://itnext.io/how-to-deploy-a-single-kubernetes-cluster-across-multiple-clouds-using-k3s-and-wireguard-a5ae176a6e81)。在解释_how_之前，我应该回答_why？_所以，让我们从问题开始。

# **What’s the Problem with Multicloud?**

# **多云有什么问题？**

Kubernetes has become the de facto cloud native computing standard in a short period of time. Released in 2014, there is a Kubernetes distribution for every platform, and most enterprises have a Kubernetes strategy.

Kubernetes 在短时间内成为事实上的云原生计算标准。 Kubernetes 发布于 2014 年，每个平台都有一个 Kubernetes 发行版，大多数企业都有 Kubernetes 策略。

This has introduced a set of problems that enterprises are still learning to solve. Chief among these problems is **inter-cloud operability**, or, the ability to manage clusters and cluster-based applications across a variety of environments.

这引入了一系列企业仍在学习解决的问题。其中最主要的问题是**云间可操作性**，即跨各种环境管理集群和基于集群的应用程序的能力。

Typical solution architectures focus on deploying multiple clusters into various environments, and then coordinating the management of both infrastructure and applications across these environments.

典型的解决方案架构侧重于将多个集群部署到各种环境中，然后协调跨这些环境的基础设施和应用程序的管理。

In many scenarios where environments are strictly segmented, such solutions are necessary. However, these architectures introduce problems of their own. They require additional complex software which must be learned and operated. They require an operations team to manage many clusters simultaneously. They also require substantial resource overhead, because at least one full cluster is required per environment.

在很多环境严格分割的场景下，这样的解决方案是很有必要的。然而，这些体系结构引入了它们自己的问题。它们需要额外的复杂软件，必须学习和操作。他们需要一个运营团队同时管理多个集群。它们还需要大量资源开销，因为每个环境至少需要一个完整的集群。

# How can we solve this?

# 我们如何解决这个问题？

Enterprises typically overlook an alternative approach to these problems: the mesh VPN. This is not to be confused with a “service mesh,” an unrelated Kubernetes concept. The mesh VPN creates a “virtual” cloud environment for a cluster, which may be composed of many physical environments.

企业通常会忽略解决这些问题的替代方法：网状 VPN。这不要与“服务网格”混淆，这是一个不相关的 Kubernetes 概念。网状 VPN 为集群创建了一个“虚拟”云环境，该集群可能由许多物理环境组成。

With a mesh VPN, enterprises can manage a **single cluster** across **multiple environments**. This method is relatively simple and has several key benefits:

借助网状 VPN，企业可以跨**多个环境**管理**单个集群**。这种方法相对简单，有几个关键的好处：

- Cloud burst into new environments
- Save on resource overhead
- No new skills or tools required

- 云突进新环境
- 节省资源开销
- 无需新技能或工具

A mesh VPN also provides an added benefit to Kubernetes clusters outside of time and savings: **security**.

除了时间和节省之外，网状 VPN 还为 Kubernetes 集群提供了额外的好处：**安全**。

A Kubernetes cluster enabled with a mesh VPN has encrypted traffic between all nodes, and enables new patterns for secure access from and to the cluster.

启用了网状 VPN 的 Kubernetes 集群对所有节点之间的流量进行了加密，并启用了进出集群的安全访问的新模式。

# What is a Mesh VPN?

# 什么是网状 VPN？

![](https://miro.medium.com/max/900/0*URDLYJlaaebN6Mlu?q=20)

A Mesh VPN is a virtual private network where every computer has a direct connection to every other computer over a private IP address.

Mesh VPN 是一个虚拟专用网络，其中每台计算机都通过专用 IP 地址直接连接到其他每台计算机。

In the context of Kubernetes, it is a virtual subnet where you deploy your worker nodes, which does not require those nodes to be deployed in the same location.

在 Kubernetes 的上下文中，它是一个虚拟子网，您可以在其中部署工作程序节点，不需要将这些节点部署在同一位置。

This enables platform teams to deploy clusters with nodes placed in arbitrary environments. For instance, you can have an on-prem data center which can “cloud burst” into a public cloud environment if more resources are needed. From the cluster’s perspective, it is just one normal Kubernetes cluster. It has no idea its nodes are placed in different locations. Simple management is enabled via node selectors.

这使平台团队能够部署具有位于任意环境中的节点的集群。例如，您可以拥有一个本地数据中心，如果需要更多资源，它可以“云爆发”到公共云环境中。从集群的角度来看，它只是一个普通的 Kubernetes 集群。它不知道它的节点被放置在不同的位置。通过节点选择器启用简单管理。

An enterprise could use any one of the available mesh VPNs, including Nebula, Tailscale, Twingate, Netmaker, and others. However, it is critically important to use a mesh VPN based on **kernel WireGuard**.

企业可以使用任何一种可用的网状 VPN，包括 Nebula、Tailscale、Twingate、Netmaker 等。但是，使用基于 **kernel WireGuard** 的网状 VPN 至关重要。

![](https://miro.medium.com/max/900/0*fQukEPouNwGQxYtI?q=20)

VPN’s have historically caused significant latency, and can lead to a 30% reduction in bandwidth or more. Additionally, they are usually complex and heavy. WireGuard, a breakthrough VPN technology, eliminates these problems. It achieves near-parity in speed to the same network without WireGuard, and has a very simple implementation relative to older technologies like OpenVPN and IPSec.

VPN 在历史上造成了显着的延迟，并可能导致带宽减少 30% 或更多。此外，它们通常既复杂又笨重。 WireGuard 是一项突破性的 VPN 技术，可消除这些问题。它在没有 WireGuard 的情况下实现了与同一网络几乎相同的速度，并且相对于 OpenVPN 和 IPSec 等旧技术具有非常简单的实现。

Only a few of the latest mesh VPN’s (such as Netmaker) utilize kernel WireGuard in order to maximize speed.

只有少数最新的网状 VPN（例如 Netmaker）使用内核 WireGuard 以最大限度地提高速度。

# What are the limitations?

# 有什么限制？

Before going into the use cases, let’s cover a few limitations. There are reasons this is not a widely used pattern today, though they are due to some misconceptions which can be solved. 

在进入用例之前，让我们先介绍一些限制。这不是今天广泛使用的模式是有原因的，尽管它们是由于一些可以解决的误解。

The largest misconception is that Kubernetes cannot run with high latency: it _can,_ but _etcd_ cannot. This means that your master nodes either must be co-located, or use a database other than etcd. For instance, SQL with K3S has no such problem. MicroK8s also runs with Dqlite, a distributed version of SQL which is tolerant to latency.

最大的误解是 Kubernetes 不能以高延迟运行：它 _可以，_ 但 _etcd_ 不能。这意味着您的主节点必须位于同一位置，或者使用 etcd 以外的数据库。比如SQL with K3S就没有这个问题。 MicroK8s 还与 Dqlite 一起运行，这是一个分布式 SQL 版本，可以容忍延迟。

In addition, operators are used to the latency cost of running with traditional VPN’s like IPSec and OpenVPN. They may not know how fast WireGuard is.

此外，运营商习惯于使用 IPSec 和 OpenVPN 等传统 VPN 运行的延迟成本。他们可能不知道 WireGuard 有多快。

There are three reasonable factors to consider that should dissuade some companies from running with this strategy: bandwidth pricing, corporate firewalls, and application-level latency.

有三个合理的因素可以阻止一些公司采用这种策略：带宽定价、企业防火墙和应用程序级延迟。

If you are using a cloud provider with very high egress data charges, you may want to run an analysis of what those charges will look like, since your nodes will be sending data back and forth between clouds. Still, this price may be lower than the price of running duplicated infra in each environment. DigitalOcean’s bandwidth pricing is likely much lower than running extra infrastructure.

如果您使用的是出口数据费用非常高的云提供商，您可能需要分析这些费用的情况，因为您的节点将在云之间来回发送数据。尽管如此，这个价格可能低于在每个环境中运行重复基础设施的价格。 DigitalOcean 的带宽定价可能远低于运行额外的基础设施。

If you are in a corporate environment with heavy firewall restrictions between environments, requiring layers of permissions to run data between them, this may also create a challenge in setting up this topology. Still, it may be worth considering if you already need to run such integrations.

如果您在企业环境中，环境之间有严格的防火墙限制，需要多层权限才能在它们之间运行数据，这也可能会给设置此拓扑带来挑战。不过，如果您已经需要运行此类集成，则可能值得考虑。

Finally, there are many cases where your applications cannot tolerate high latency. In these cases, you can solve the problem with simple affinity rules and node selectors: Per location, give nodes a label, and then set up affinity rules on your cluster so that pods in applications will schedule to one group or the other by default.

最后，在很多情况下，您的应用程序不能容忍高延迟。在这些情况下，您可以使用简单的关联规则和节点选择器来解决问题：每个位置，给节点一个标签，然后在您的集群上设置关联规则，以便应用程序中的 pod 在默认情况下调度到一个组或另一个组。

There are some cases where it makes more sense to run duplicated infrastructure to avoid dealing with potential inter-cloud issues, but there are less of these cases than you might think. Still, it is always important to consider the potential costs and hazards involved.

在某些情况下，运行重复的基础架构以避免处理潜在的云间问题更有意义，但这些情况比您想象的要少。尽管如此，考虑所涉及的潜在成本和危害总是很重要的。

# What are the use cases?

# 用例是什么？

Assuming a company is not heavily impacted by the above limitations, a mesh VPN enables many valuable use cases, which both decrease cost and complexity while enabling powerful deployment patterns. We discuss these different use cases below.

假设一家公司不受上述限制的严重影响，网状 VPN 可以实现许多有价值的用例，既降低成本和复杂性，同时启用强大的部署模式。我们将在下面讨论这些不同的用例。

![](https://miro.medium.com/max/600/0*83CJvB0Hk0c0R_K2?q=20)

**Case 0: Regular Cluster**

**案例 0：常规集群**

If you are planning to deploy a cluster, it is always worth considering a WireGuard kernel-based mesh VPN underneath. You will have a normal functioning cluster with negligible performance differences. However, you enable your cluster to run with enhanced topologies in the future, with the added benefit of encrypted inter-node traffic.

如果您计划部署集群，那么始终值得考虑在其下使用基于 WireGuard 内核的网状 VPN。您将拥有一个正常运行的集群，其性能差异可以忽略不计。但是，您可以让您的集群在未来以增强的拓扑运行，并获得加密节点间流量的额外好处。

**Case 1: Wide Nodes**

**案例 1：宽节点**

In this scenario, imagine a cluster from Case 0, but now you want to deploy an application that is integrated with an Azure service. You simply deploy a VM in Azure, add the VM to your mesh, and install the node. You now have the below topology. Migrating an application between the two clouds is as simple as changing the node selector.

在此方案中，假设案例 0 中的群集，但现在您想要部署与 Azure 服务集成的应用程序。您只需在 Azure 中部署一个 VM，将该 VM 添加到您的网格中，然后安装该节点。您现在拥有以下拓扑。在两个云之间迁移应用程序就像更改节点选择器一样简单。

![](https://miro.medium.com/max/900/0*gE3w6w2gvxGteEGV?q=20)

![](https://miro.medium.com/max/600/0*re8_QyXLE_YRNEzi?q=20)

**Case 2: Cloud Bursting**

**案例 2：云爆发**

In this scenario, an enterprise has an on-prem data center they are using for Kubernetes. They have deployed the mesh VPN underneath their cluster, which enables them to add nodes arbitrarily from a cloud environment (DigitalOcean, AWS, etc.). This can be very useful for cases where you need to scale your application quickly but have limited resources locally.

在这种情况下，企业拥有一个用于 Kubernetes 的本地数据中心。他们在集群下部署了网状 VPN，这使他们能够从云环境（DigitalOcean、AWS 等）中任意添加节点。这对于需要快速扩展应用程序但本地资源有限的情况非常有用。

![](https://miro.medium.com/max/600/0*8Bp-71SJQDTCidHx?q=20)

**Case 3: Cloud Control**

**案例 3：云控制**

In this scenario, an enterprise has substantial on-prem resources, but does not want to manage the control plane locally. Master node failure can be catastrophic and they’d rather use cloud resources. Because of this, the enterprise utilizes a cloud-based control plane but uses on-prem workers. This can bring down cloud costs substantially.

在这种情况下，企业拥有大量本地资源，但不想在本地管理控制平面。主节点故障可能是灾难性的，他们宁愿使用云资源。因此，企业使用基于云的控制平面，但使用本地员工。这可以大大降低云成本。

This scenario can also be applied to using **cloud-based storage nodes**, another critical component where failure is not an option.

此场景也可应用于使用**基于云的存储节点**，这是另一个无法选择故障的关键组件。

**Case 4: Distributed/Edge Nodes** 

**案例 4：分布式/边缘节点**

In this scenario, worker nodes are all placed in different environments. This use case can also be applied to edge environments, where you have one central control plane pushing out applications (perhaps via daemonset) to all of the edge nodes.

在这种情况下，工作节点都放置在不同的环境中。此用例也可应用于边缘环境，在这种环境中，您有一个中央控制平面将应用程序（可能通过 daemonset）推送到所有边缘节点。

![](https://miro.medium.com/max/900/0*UTVZCDwpWwaYP4c7?q=20)

**Case 5: Distributed Clusters**

**案例 5：分布式集群**

This scenario is more complex and requires a deeper dive, but with the right architecture, the entire cluster may be run across arbitrary clouds, including the master nodes. The main limitation here is etcd, which does not tolerate high latency. Such a pattern may require a non-etcd cluster database.

这个场景更复杂，需要更深入的研究，但有了正确的架构，整个集群可以跨任意云运行，包括主节点。这里的主要限制是 etcd，它不能容忍高延迟。这种模式可能需要一个非 etcd 集群数据库。

**Case 6: Connected Clusters**

**案例 6：连接的集群**

In this scenario, clusters are connected over the mesh VPN and can communicate with each other directly, meaning a microservice-based application can be split between the two clusters. This may be used in conjunction with multi-cluster management tools.

在这种情况下，集群通过网状 VPN 连接并且可以直接相互通信，这意味着基于微服务的应用程序可以在两个集群之间拆分。这可以与多集群管理工具结合使用。

![](https://miro.medium.com/max/900/0*krXPHPaHhzMFTWyh?q=20)

**Case 7: Secure, private access (inbound)**

**案例 7：安全、私密的访问（入站）**

In this scenario, an enterprise only wants trusted users or applications to access cluster resources. Here, the mesh VPN acts similarly to a corporate VPN, where one must attach to the VPN in order to gain access. This can also be performed at a lower level to allow outside access to the service and pod networks within the Kubernetes cluster, a useful feature for development and operations teams.

在这种情况下，企业只希望受信任的用户或应用程序访问集群资源。在这里，网状 VPN 的作用类似于企业 VPN，必须连接到 VPN 才能获得访问权限。这也可以在较低级别执行，以允许外部访问 Kubernetes 集群内的服务和 pod 网络，这是一个对开发和运营团队有用的功能。

![](https://miro.medium.com/max/900/0*ptp8f6Gi_3e_8rBF?q=20)

**Case 8: Secure, private access (outbound)**

**案例 8：安全、私密的访问（出站）**

In this scenario, a Kubernetes-based application needs secure access to a non-Kubernetes application such as a SQL database. The database VM is added to the VPN and the application now has secure, instant access (see above). This method can be used to add any arbitrary non-Kubernetes application to the Kubernetes network and allow pods to use the service.

在这种情况下，基于 Kubernetes 的应用程序需要安全访问非 Kubernetes 应用程序，例如 SQL 数据库。数据库虚拟机被添加到 VPN 中，应用程序现在可以安全、即时访问（见上文）。此方法可用于将任意非 Kubernetes 应用程序添加到 Kubernetes 网络并允许 pod 使用该服务。

# **Options for a Mesh VPN**

# **网状 VPN 的选项**

Several options exist for a mesh VPN today, including Tailscale, Netmaker, Kilo, Nebula, and others. Some of the above use cases become less feasible based on the choice of VPN. In addition, only Netmaker and Kilo offer kernel WireGuard as an option, a key consideration to minimize latency.

当今的网状 VPN 存在多种选择，包括 Tailscale、Netmaker、Kilo、Nebula 等。根据 VPN 的选择，上述一些用例变得不太可行。此外，只有 Netmaker 和 Kilo 提供内核 WireGuard 作为选项，这是最小化延迟的关键考虑因素。

A company can also “roll their own” mesh using WireGuard directly, but this becomes very difficult to manage at scale and requires significant manual intervention.

公司也可以直接使用 WireGuard 来“滚动他们自己的”网格，但这变得非常难以大规模管理，并且需要大量的人工干预。

GRAVITL has designed [Netmaker](https://gravitl.com/netmaker) to handle all of the above use cases while being largely automated and based on kernel WireGuard. In addition, since it is a generalized mesh VPN, it can be used for intergrating external services. That said, every company should make their own determination on which network virtualization tool is most suited to their environments.

GRAVITL 设计了 [Netmaker](https://gravitl.com/netmaker) 来处理上述所有用例，同时在很大程度上实现自动化并基于内核 WireGuard。此外，由于它是一个广义的网状 VPN，它可以用于集成外部服务。也就是说，每家公司都应该自行决定哪种网络虚拟化工具最适合其环境。

# **Conclusion**

#  **结论**

We have discussed the current state of multi-cloud Kubernetes clusters, the need for a mesh VPN within Kubernetes, and the different use cases it enables.

我们已经讨论了多云 Kubernetes 集群的当前状态、Kubernetes 内对网状 VPN 的需求以及它支持的不同用例。

It is recommended that platform owners consider deploying a mesh VPN underneath their clusters, even if they do not have a current-state use case, as it does not significantly impact cluster performance, and allows for quick enablement of the above topologies when needed.

建议平台所有者考虑在其集群下部署网状 VPN，即使他们没有当前状态的用例，因为它不会显着影响集群性能，并且允许在需要时快速启用上述拓扑。

For more information about mesh VPN's or [Netmaker](https://gravitl.com/netmaker), contact GRAVITL at [info@gravitl.com](mailto:info@gravitl.com), or visit our website at [https: //gravitl.com/book](https://gravitl.com/book) .

有关网状 VPN 或 [Netmaker](https://gravitl.com/netmaker) 的更多信息，请通过 [info@gravitl.com](mailto:info@gravitl.com) 联系 GRAVITL，或访问我们的网站 [https: //gravitl.com/book](https://gravitl.com/book) .

[**ITNEXT**](https://itnext.io/?source=post_sidebar--------------------------post_sidebar-----------)

ITNEXT is a platform for IT developers & software engineers… 

ITNEXT 是一个面向 IT 开发人员和软件工程师的平台……

