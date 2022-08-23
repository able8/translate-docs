# Difference Between multi-cluster, multi-master, multi-tenant & federated Kubernetes

# 多集群、多主机、多租户和联合 Kubernetes 之间的区别

Last updated June 9, 2021

https://platform9.com/blog/difference-between-multi-cluster-multi-master-multi-tenant-federated-kubernetes

Just as there are many parts that comprise Kubernetes, there are multiple ways to go about [Kubernetes deployment](https://platform9.com/docs/deploy-kubernetes-the-ultimate-guide/). The best approach for your needs depends on your team’s technical expertise, your infrastructure availability (or lack thereof), your capital expenditure and ROI goals, and more.

正如 Kubernetes 的组成部分很多一样，Kubernetes 部署 也有多种方法。满足您需求的最佳方法取决于您团队的技术专长、您的基础设施可用性（或缺乏）、您的资本支出和 ROI 目标等等。

There are multiple ways to scale a Kubernetes environment to meet the needs of your organization’s requirements and manage the infrastructure as it grows. While they have similar-sounding names, they have very different applications in the real world. Read on to see how multi-tenant, multi-cluster, multi-master, and Federation fit into the mix.

有多种方法可以扩展 Kubernetes 环境，以满足您组织的需求并随着基础设施的发展对其进行管理。尽管它们的名称听起来相似，但它们在现实世界中的应用却截然不同。继续阅读以了解多租户、多集群、多主机和联合如何融入其中。

## Summary of Deployment Models

## 部署模型总结

The deployment models for Kubernetes range from single-server acting as master and worker with a single tenant, all the way to multiple multi-node clusters across multiple data centers, with federation enabled for some applications.

Kubernetes 的部署模型范围从作为主服务器的单服务器和具有单个租户的工作人员，一直到跨多个数据中心的多个多节点集群，并为某些应用程序启用了联合。

In this article we will focus on the more common deployment models and how they are different, as opposed to how they are similar. Because they all use the same base product, they have most functionality in common.

在本文中，我们将重点介绍更常见的部署模型以及它们的不同之处，而不是它们的相似之处。因为它们都使用相同的基础产品，所以它们具有大多数共同点。

Some key differences include single-server models which are typically optimized for the use of developers, and not meant for production. Whereas production systems recommend a minimum of multi-master (with an odd number of nodes) and often end up with multi-cluster, and Federation becomes a very real consideration when the deployments need to scale to thousands of nodes and across regions.

一些关键差异包括通常为开发人员使用而优化的单服务器模型，而不是用于生产。而生产系统推荐最少的多主服务器（具有奇数个节点）并且通常以多集群告终，当部署需要扩展到数千个节点和跨区域时，联合成为一个非常现实的考虑因素。

## Single-tenant

## 单租户

A simple deployment, often using a single server, is perfect for a developer or QA team member who is trying to ensure containers are being built properly, and all scripts are Kubernetes compatible. These single-server deployments use a mini distribution such as minikube, or a limited deployment of a full suite like Platform9 Managed Kubernetes Free Tier (PMKFT).

一个简单的部署，通常使用单个服务器，非常适合试图确保正确构建容器并且所有脚本都与 Kubernetes 兼容的开发人员或 QA 团队成员。这些单服务器部署使用 minikube 等迷你发行版，或 Platform9 托管 Kubernetes 免费层 (PMKFT) 等完整套件的有限部署。

These single-server single-tenant deployments run everything, not only on the same host but in the same name space. This is NOT a good practice for production, as it exposes application, deployment, and pod-specific configuration items – from secrets, to resource items like storage volumes, to all containers on the node. In addition, the network is wide open, so there are no restrictions on container interactions, which is against both network and security best practices.

这些单服务器单租户部署运行一切，不仅在同一主机上，而且在同一名称空间中。这对于生产来说不是一个好的实践，因为它暴露了应用程序、部署和特定于 pod 的配置项——从机密到存储卷等资源项，再到节点上的所有容器。此外，网络是完全开放的，因此对容器交互没有任何限制，这既违反了网络最佳实践，也违反了安全最佳实践。

[![](https://platform9.com/wp-content/uploads/2020/01/Single-Server-Single-Tenant-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Single-Server-Single-Tenant.jpeg)

## Multi-tenant

##  多租户

Once you passed the basic stage of just having Kubernetes running – whether it was a mini distribution or a full Kubernetes install with only the basic default networking enabled – you have, no doubt, run into the issue where services can access whatever they want, with no control. This makes enforcing any kind of security model or application flow nearly impossible, as developers can simply elect to ignore it and call anything they want.

一旦你通过了让 Kubernetes 运行的基本阶段——无论是迷你发行版还是只启用基本默认网络的完整 Kubernetes 安装——毫无疑问，你会遇到服务可以访问任何他们想要的东西的问题，使用无控制。这使得执行任何类型的安全模型或应用程序流程几乎是不可能的，因为开发人员可以简单地选择忽略它并调用他们想要的任何东西。

Multi-tenant is the term for when a program, like a Kubernetes cluster, is configured to allow multiple areas within its purview to operate in isolation. This is critical when any kind of data is used, especially if it falls under financial or privacy regulations. The additional benefits are that multiple development teams, QA staff, or other entities can work in parallel within the same cluster with their own quotas and security profiles, without having any negative impact on the other teams also working in the cluster.

多租户是指将程序（如 Kubernetes 集群）配置为允许其权限内的多个区域独立运行的术语。当使用任何类型的数据时，这一点至关重要，尤其是在它属于财务或隐私法规的情况下。额外的好处是，多个开发团队、QA 人员或其他实体可以在同一个集群中并行工作，具有自己的配额和安全配置文件，而不会对也在集群中工作的其他团队产生任何负面影响。

Within Kubernetes, this is primarily accomplished at the network layer for deployed pods that can only see pods within their name space by default, and can be granted access via network policies to get access to other namespaces and the applications they are running. 

在 Kubernetes 中，这主要是在网络层为已部署的 Pod 完成的，默认情况下只能在其名称空间内看到 Pod，并且可以通过网络策略授予访问权限以访问其他命名空间和它们正在运行的应用程序。

Additional technology, called Service Meshes, can be added to environments that are multi-tenant. The purpose is to increase network resilience between pods, enforce authentication and authorization policies, and to increase compliance to defined network policies. Service Meshes can also be leveraged for inter-cluster communications, but that is beyond what most deployments currently use them for.

可以将称为服务网格的附加技术添加到多租户环境中。目的是提高 pod 之间的网络弹性，强制执行身份验证和授权策略，并提高对已定义网络策略的遵从性。服务网格也可用于集群间通信，但这超出了大多数部署当前使用它们的范围。

[![](https://platform9.com/wp-content/uploads/2020/01/Single-Server-Multi-tenant-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Single-Server-Multi-tenant.jpeg)

## Multi-master

## 多主

When Kubernetes is going to be used to support any kind of real workload, it is best to start down the path of having clusters be multi-master. By having more than one master server, the Kubernetes cluster can continue to be administered even in the event that not only a container like the api server fails, but if an entire server crashes taking out instances of all the operators and other controllers that are deployed and required to maintain a steady state across all the worker nodes.

当要使用 Kubernetes 来支持任何类型的实际工作负载时，最好从让集群成为多主节点的路径开始。通过拥有多个主服务器，即使不仅像 api 服务器这样的容器发生故障，而且如果整个服务器崩溃并取出所有已部署的操作员和其他控制器的实例，Kubernetes 集群也可以继续进行管理并且需要在所有工作节点上保持稳定状态。

Special consideration needs to be paid to etcd, which is the most commonly-used datastore for Kubernetes clusters. As etcd requires a quorum to support read-write operations, it needs to have a minimum of three nodes to be highly available and recommends having an odd number that will increase based on the number of worker nodes in the cluster. By the time you get to about 100 worker nodes, you’ll be having five to seven etcd servers to handle the load. Some Kubernetes distributions have etcd share the same servers as the master control plane containers, which works great as long as the individual nodes are sized appropriately.

需要特别考虑 etcd，它是 Kubernetes 集群最常用的数据存储。由于 etcd 需要一个仲裁来支持读写操作，因此它需要至少三个节点才能具有高可用性，并且建议使用一个奇数，该奇数会根据集群中工作节点的数量而增加。当您达到大约 100 个工作节点时，您将拥有 5 到 7 个 etcd 服务器来处理负载。一些 Kubernetes 发行版的 etcd 与主控制平面容器共享相同的服务器，只要各个节点的大小合适，它就可以很好地工作。

[![](https://platform9.com/wp-content/uploads/2020/01/Multi-Master-Multi-Tenant-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Multi-Master-Multi-Tenant.jpeg)

## Multi-cluster

## 多集群

Multi-cluster environments are typically configured as multi-master and multi-tenant, as well. If you are big enough to need multiple clusters, there are probably high-availability and isolation requirements driven by at least a few of the groups that are actively building and deploying containers within your hybrid cloud infrastructure.

多集群环境通常也配置为多主机和多租户。如果您的规模大到需要多个集群，那么至少有一些团队可能会推动高可用性和隔离要求，这些团队正在您的混合云基础架构中积极构建和部署容器。

A multi-cluster deployment is simply deploying multiple clusters across one or more data centers, ideally with a single management interface to simplify life; although, that is not always the case. Companies like Platform9 offer products with multi-cluster management capabilities that can work with infrastructure provided by all the biggest cloud providers, in addition to on-premises infrastructure like OpenStack and VMware.

多集群部署只是跨一个或多个数据中心部署多个集群，理想情况下使用单一管理界面来简化生活；但是，情况并非总是如此。 Platform9 等公司提供的产品具有多集群管理功能，除了 OpenStack 和 VMware 等本地基础架构外，这些产品还可以与所有最大的云提供商提供的基础架构一起使用。

Reasons you would get into a multi-cluster scenario is separation of development and production; or, since it is not recommended to have a Kubernetes cluster that spans more than a single datacenter, having one or more regions in use across single- or multiple-cloud providers.

进入多集群场景的原因是开发和生产分离；或者，由于不建议拥有跨越多个数据中心的 Kubernetes 集群，因此在单个或多个云提供商之间使用一个或多个区域。

[![](https://platform9.com/wp-content/uploads/2020/01/Multi-Cluster-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Multi-Cluster.jpeg)

## Federation

## 联邦

Kubernetes Federation is the cherry on top of a multi-cluster deployment as it allows parts of the configuration to be synchronized across specific namespaces in the clusters that are federated. Where this becomes invaluable is when you have the same application and services spread across multiple datacenters or cloud regions, for reasons like capacity and avoiding a single point of failure. 

Kubernetes Federation 是多集群部署之上的樱桃，因为它允许部分配置在联合集群中的特定命名空间之间同步。当您出于容量和避免单点故障等原因将相同的应用程序和服务分布在多个数据中心或云区域中时，这变得非常宝贵。

By having these applications managed through Federation, the CI/CD or another deployment process can issue a single Kubernetes Deployment. It can roll out an application update across a truly global infrastructure while coordinating between the member clusters to ensure the chosen deployment model is being adhered to regardless of what cloud provider or region it is located in. This could be a hard cutover to the new application , canary deployments, or even just a gradual cutover leveraging a service mesh so there are no interrupted user sessions.

通过通过联合管理这些应用程序，CI/CD 或其他部署过程可以发布单个 Kubernetes 部署。它可以在真正的全球基础架构中推出应用程序更新，同时在成员集群之间进行协调，以确保无论其位于哪个云提供商或区域，都遵循所选的部署模型。这可能是对新应用程序的硬切换，金丝雀部署，甚至只是利用服务网格的逐步切换，因此没有中断的用户会话。

[![](https://platform9.com/wp-content/uploads/2020/01/Federated-Clusters-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Federated-Clusters.jpeg)

## Conclusion

##  结论

Kubernetes has a deployment model to meet the needs of any organization regardless of the scale they operate at, or their individual requirements. Within many organizations, multiple different deployment models are often used to support different phases of an application’s lifecycle and the needs of different business units. From the CTO office just needing a wide open playground to try things, to the shopping cart requiring truly isolated services to be PCI compliant.

Kubernetes 有一个部署模型，可以满足任何组织的需求，无论他们的运营规模如何，或者他们的个人需求如何。在许多组织中，经常使用多种不同的部署模型来支持应用程序生命周期的不同阶段以及不同业务部门的需求。从 CTO 办公室只需要一个开阔的游乐场来尝试事物，到购物车需要真正隔离的服务才能符合 PCI 标准。

Platform9 is the world’s #1 open distributed cloud service, offering the power of the public cloud on infrastructure of customers’ choice — powered by Kubernetes and cloud-native technologies. Public clouds are walled gardens, and DIY is difficult and time consuming. Platform9 offers a third option — an open and faster option — enabling a better way to go cloud-native. Platform9's service powers 40K+ nodes across private, public, and edge clouds. Innovative enterprises like Juniper, Kingfisher Plc, Mavenir, Redfin, and Cloudera achieve 4x faster time-to-market, up to 90% reduction in operational costs, and 99.9% uptime . Platform9 is an inclusive, globally distributed company backed by leading investors.

Platform9 是世界排名第一的开放分布式云服务，在客户选择的基础设施上提供公共云的强大功能——由 Kubernetes 和云原生技术提供支持。公共云是有围墙的花园，DIY 既困难又耗时。 Platform9 提供了第三种选择——一种开放且更快的选择——为实现云原生提供了一种更好的方式。 Platform9 的服务为私有云、公共云和边缘云中的 40K+ 节点提供支持。Juniper、Kingfisher Plc、Mavenir、Redfin 和 Cloudera 等创新企业实现了 4 倍的上市时间、高达 90% 的运营成本降低和 99.9% 的正常运行时间。Platform9 是一家由领先投资者支持的包容性、全球分布的公司。

Latest posts by Platform9 ( [see all](https://platform9.com/blog/author/platform9/)) 

