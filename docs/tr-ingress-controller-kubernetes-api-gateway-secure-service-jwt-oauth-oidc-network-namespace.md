## Why and How of Kubernetes Ingress (and Networking)

## Kubernetes Ingress（和网络）的原因和方式

[Saaras Inc.](http://getenroute.io/author/saaras-inc.)August 07, 2021

[Saaras Inc.](http://getenroute.io/author/saaras-inc.)2021 年 8 月 7 日

![Why and How of Kubernetes Ingress (and Networking)](https://getenroute.io/img/EnrouteIngressDetail.jpeg)

Services running in Kubernetes are not accessible on public or private cloud. This is how Kubernetes is designed considering service security in mind.

无法在公共或私有云上访问在 Kubernetes 中运行的服务。这就是考虑到服务安全性的 Kubernetes 设计方式。

Securely allowing access to a service outside the cluster requires some understanding of how networking is setup and the different requirements driving the networking choices.

安全地允许访问集群外的服务需要对网络的设置方式以及推动网络选择的不同要求有一定的了解。

We briefly start by exploring what is expected from a kubernetes cluster when it comes to service isolation, service scaling and service delivery. Once the high level requirements are laid out, it is easier to understand the significance of different constructs and abstractions.

我们首先简要地探讨在服务隔离、服务扩展和服务交付方面对 Kubernetes 集群的期望。一旦列出了高级需求，就更容易理解不同构造和抽象的重要性。

We conclude by contrasting the advantages of using an Ingress to run a layer of L7 policy (or proxies) in front of a service running inside Kubernetes.

我们通过对比使用 Ingress 在 Kubernetes 内部运行的服务前运行 L7 策略（或代理）层的优势得出结论。

### Understanding the scheme of Kubernetes Networking

### 理解Kubernetes Networking的方案

Understanding Kubernetes Ingress is key to running microservices and securely accessing those services. This article is an attempt to demystify how Kubernetes networking is setup.

了解 Kubernetes Ingress 是运行微服务和安全访问这些服务的关键。本文试图揭开 Kubernetes 网络是如何设置的神秘面纱。

We look at networking when a service is created, the different Kubernetes artifacts created, the networking machinery required to meet different requirements.

我们在创建服务时查看网络、创建的不同 Kubernetes 工件、满足不同需求所需的网络机制。

We also describe the significance of different types of IPs like External-IP, Node-IP, Cluster-IP, Pod-IP and describe how traffic passes through each one of them.

我们还描述了不同类型 IP（如外部 IP、节点 IP、集群 IP、Pod-IP）的重要性，并描述了流量如何通过它们中的每一个。

Starting with cluster networking requirements provides us an opportunity of why networking is setup the way it is.

从集群网络要求开始，我们有机会了解为什么网络是这样设置的。

#### Cluster Networking Requirements

#### 集群网络要求

**Cluster networking** in Kubernetes has several **requirements**

**Kubernetes 中的集群网络**有几个**要求**

- **security and isolation of service(s)**
- **connection, networking and IP allocation for pods**
- **setup networking to build a cluster abstraction out of several physical nodes**
- **load balancing of traffic across multiple instances of a service**
- **controlling external access to a service**
- **working with Kubernetes networking in public and private cloud environments.**

- **服务的安全和隔离**
- **Pod 的连接、网络和 IP 分配**
- **设置网络以从多个物理节点构建集群抽象**
- **跨服务的多个实例的流量负载平衡**
- **控制对服务的外部访问**
- **在公共和私有云环境中使用 Kubernetes 网络。**

To understand these different aspects of Kubernetes networking, we start by describing what happens when a service is created in a pod all the way to accessing that service in public and private cloud.

为了理解 Kubernetes 网络的这些不同方面，我们首先描述从在 pod 中创建服务一直到在公共和私有云中访问该服务时会发生什么。

We highlight the need for Ingress and how it fits into the overall Kubernetes networking model.

我们强调了对 Ingress 的需求以及它如何适应整个 Kubernetes 网络模型。

#### Network isolation of Service running in Kubernetes Pods

#### Kubernetes Pod 中运行的 Service 的网络隔离

Let us consider a simple Kubernetes cluster with two nodes

让我们考虑一个具有两个节点的简单 Kubernetes 集群

![](https://d33wubrfki0l68.cloudfront.net/9ba692de6b1b8123ba6657013707b7c33ad44229/35bb5/img/k8s-cluster-no-svc.png)

Kubernetes orchestrates containers or pods (which is a group of containers). **When Kubernetes creates a pod, it is run in its own isolated network (using network namespace).**

Kubernetes 编排容器或 pod（这是一组容器）。 **当 Kubernetes 创建一个 Pod 时，它运行在自己的隔离网络中（使用网络命名空间）。**

The diagram below shows two pods created on each node.

下图显示了在每个节点上创建的两个 Pod。

![](https://d33wubrfki0l68.cloudfront.net/9db66d843ff1651a60729bbb73357e8e8dc66837/66bf6/img/k8s-pod-ip.png)

What does this mean for a service? A service runs inside a pod in the pod’s network. **An IP address allocated on this pod network (for the service) isn’t accessible outside the pod.**

这对服务意味着什么？服务在 pod 网络中的 pod 内运行。 **在此 Pod 网络上分配的 IP 地址（用于服务）在 Pod 外部无法访问。**

So how do you access this service?

那么如何访问这项服务呢？

#### Making Service in Pod accessible to host network stack

#### 使主机网络堆栈可以访问 Pod 中的服务

Kubernetes builds an abstraction of a cluster on top of multiple physical nodes or machines. The physical nodes have their own network stack. A pod created by Kubernetes creates an isolated network stack for the services that run inside the pod.

Kubernetes 在多个物理节点或机器之上构建集群的抽象。物理节点有自己的网络堆栈。 Kubernetes 创建的 pod 为在 pod 内运行的服务创建了一个隔离的网络堆栈。

To reach this service (or IP address inside the pod), there needs to be routing/bridging that creates a path between the pod network and the host network. **Container Networking Interface or CNI sets up the networking associated with creating a traffic path between the node and pod.** Popular examples of CNI are calico, cilium, flannel, etc.

要访问此服务（或 Pod 内的 IP 地址），需要路由/桥接在 Pod 网络和主机网络之间创建路径。 **容器网络接口或 CNI 设置与在节点和 Pod 之间创建流量路径相关的网络。** CNI 的流行示例是 calico、cilium、flannel 等。

![](https://d33wubrfki0l68.cloudfront.net/c24a73d5ffe4bae6a90b11575322c7cffe5a1b75/6e56a/img/k8s-pod-ip-with-cni.png)

When Kubernetes creates a pod, it invokes the CNI callbacks. These callbacks result in the invocation of CNI provider services to setup IP addresses for the pod and connecting the pod network with the host network.

当 Kubernetes 创建一个 pod 时，它会调用 CNI 回调。这些回调导致调用 CNI 提供程序服务来为 pod 设置 IP 地址并将 pod 网络与主机网络连接。

#### Making a Service accessible across Node boundaries 

#### 跨节点访问服务

A service resides in a pod or several pods. Each of these pods can reside on one physical node or multiple physical nodes. As an example, say a service is spread across two pods that reside on two physical nodes.

一个服务驻留在一个或多个 Pod 中。这些 Pod 中的每一个都可以驻留在一个物理节点或多个物理节点上。例如，假设一项服务分布在驻留在两个物理节点上的两个 Pod 中。

When traffic is destined to this service (spread on two pods across two nodes), how does Kubernetes load balance traffic across them?

当流量以该服务为目的地（分布在两个节点的两个 Pod 上）时，Kubernetes 如何负载均衡它们之间的流量？

Kubernetes uses an abstraction of a Cluster IP. **Any traffic destined to Cluster IP is load-balanced across pods (in which the service runs)**.

Kubernetes 使用集群 IP 的抽象。 **任何发往集群 IP 的流量都在 pod（服务在其中运行）之间进行负载平衡**。

To load balance to service instances in pods, networking is setup to reach the service in these pods. These pods may be running on different physical Nodes of the cluster. Wiring up a Cluster IP for a service ensures, that traffic sent to Cluster IP can be sent to all pods that run the service; regardless of the physical Node on which pods runs.

为了对 Pod 中的服务实例进行负载平衡，需要设置网络以访问这些 Pod 中的服务。这些 Pod 可能运行在集群的不同物理节点上。为服务连接集群 IP 可确保发送到集群 IP 的流量可以发送到运行该服务的所有 pod；无论 Pod 运行在哪个物理节点上。

![](https://d33wubrfki0l68.cloudfront.net/e51cdd0b4cbafc371435491ad8bc092e147e2772/13c47/img/k8s-cluster-ip-service.png)

Implementation and realization of Cluster IP is achieved by kube-proxy component and mechanisms like iptables, ipvs or user space traffic steering.

Cluster IP 的实现和实现是通过 kube-proxy 组件和 iptables、ipvs 或用户空间流量导向等机制实现的。

#### Accessing a service from outside the cluster

#### 从集群外部访问服务

Traffic destined to a ClusterIP is load-balanced across pods that may span multiple physical nodes. But the ClusterIP is only accessible from the nodes in the cluster. Or, put another way, **networking in Kubernetes ensures that external access to ClusterIP is restricted.**

发往 ClusterIP 的流量在可能跨越多个物理节点的 pod 之间进行负载平衡。但是 ClusterIP 只能从集群中的节点访问。或者，换句话说，**Kubernetes 中的网络确保限制对 ClusterIP 的外部访问。**

Accessing the ClusterIP outside the cluster needs an explicit declaration to make it accessible outside the Nodes in a Kubernetes Cluster. This is `NodePort`

在集群外访问 ClusterIP 需要一个显式声明，以使其可在 Kubernetes 集群中的节点外访问。这是`NodePort`

**A `NodePort` in Kubernetes wires up an Node IP (and port) with ClusterIP.**

**Kubernetes 中的 `NodePort` 将节点 IP（和端口）与 ClusterIP 连接起来。**

Defining a `NodePort` provides an IP address on the local network. **Traffic sent to this NodePort IP (and port) is then routed to ClusterIP and eventually load balanced to the pods (and services).**

定义一个 `NodePort` 会在本地网络上提供一个 IP 地址。 **发送到这个 NodePort IP（和端口）的流量然后被路由到 ClusterIP 并最终负载平衡到 pods（和服务）。**

![](https://d33wubrfki0l68.cloudfront.net/b75cd368ad5928594d04f675c403d94958d4536a/c6dba/img/k8s-node-port.png)

#### Accessing service in Kubernetes on a public cloud

#### 在公共云上访问 Kubernetes 中的服务

A `NodePort` makes a service accessible outside the cluster, but the IP address is only available locally. A `LoadBalancer` service is a way to associate a public IP (or DNS) with the `NodePort` service.

`NodePort` 使服务可在集群外访问，但 IP 地址仅在本地可用。 `LoadBalancer` 服务是一种将公共 IP（或 DNS）与 `NodePort` 服务相关联的方法。

When a `LoadBalancer` type of service is created in a Kubernetes cluster, it allocates a public IP and sets up the load balancer on the cloud provider (like AWS, GCP, OCI, Azure etc.). **The cloud load balancer is configured to pipe traffic sent to the external IP to the `NodePort` service.**

当在 Kubernetes 集群中创建“LoadBalancer”类型的服务时，它会分配一个公共 IP 并在云提供商（如 AWS、GCP、OCI、Azure 等）上设置负载均衡器。 **云负载均衡器配置为将发送到外部 IP 的流量通过管道传输到 `NodePort` 服务。**

![](https://d33wubrfki0l68.cloudfront.net/665c1caa1fb4dad7707ea9ab1ad0debe7b33e755/8b20c/img/k8s-external-ip-service.png)

#### Accessing service in Kubernetes on a private cloud

#### 在私有云上访问 Kubernetes 中的服务

When running in a private cloud, creating a `LoadBalancer` type of service needs a Kubernetes controller that can provision a load balancer. One such implementation is [MetalLb](https://metallb.universe.tf/), which allocates an IP to route external traffic inside a cluster.

在私有云中运行时，创建“LoadBalancer”类型的服务需要一个可以提供负载均衡器的 Kubernetes 控制器。其中一种实现是 [MetalLb](https://metallb.universe.tf/)，它分配一个 IP 来路由集群内的外部流量。

### Accessing Service on a Public Cloud with or without Ingress

### 使用或不使用 Ingress 访问公共云上的服务

There are a couple of ways to access a service running inside a Kubernetes Cluster on a public cloud. On a public cloud, when a service is of type `LoadBalancer` an External IP is allocated for external access.

有几种方法可以访问在公共云上的 Kubernetes 集群中运行的服务。在公共云上，当服务类型为“LoadBalancer”时，会为外部访问分配一个外部 IP。

- **A service can be directly declared as a `LoadBalancer` type.**

- **服务可以直接声明为`LoadBalancer`类型。**

- **Alternatively, the Ingress service that controls and configures a proxy can be declared of type `LoadBalancer`. Routes and policies can then be created on this Ingress service, to route external traffic to the destination service.**

- **或者，控制和配置代理的 Ingress 服务可以声明为 `LoadBalancer` 类型。然后可以在此 Ingress 服务上创建路由和策略，以将外部流量路由到目标服务。**

A proxy like Envoy/Nginx/HAProxy can receive all external traffic entering a cluster by running it as a service and defining this service of type `LoadBalancer`. These proxies can be configured using L7 routing and security rules. A collection of such rules forms the Ingress rules.

像 Envoy/Nginx/HAProxy 这样的代理可以通过将其作为服务运行并定义该服务类型为“LoadBalancer”来接收进入集群的所有外部流量。这些代理可以使用 L7 路由和安全规则进行配置。这些规则的集合构成了入口规则。

#### Without Ingress - Directly accessing a service by making it a service of type `LoadBalancer`

#### 没有 Ingress - 通过将其设为 `LoadBalancer` 类型的服务来直接访问服务

**When a service is declared as a type `LoadBalancer`, it directly receives traffic from the external load balancer.** In the diagram below, the service `helloenroute` service is declared of type `LoadBalancer`. It directly receives traffic from the external load balancer.

**当一个服务被声明为`LoadBalancer`类型时，它直接从外部负载均衡器接收流量。*在下图中，服务`helloenroute`服务被声明为`LoadBalancer`类型。它直接从外部负载均衡器接收流量。

![](https://d33wubrfki0l68.cloudfront.net/81664a4db818bf4814a6e1e0fd6916a10632cd64/22e67/img/helloenroutedirectaccess.png)

#### With Ingress - Putting the Service behind a Proxy that is externally accessible through a `LoadBalancer`

#### 使用 Ingress - 将服务置于可通过 `LoadBalancer` 外部访问的代理之后

A layer of L7 proxies can be placed before the service to apply L7 routing and policies. To achieve this, an Ingress Controller is needed.

可以在服务之前放置一层 L7 代理以应用 L7 路由和策略。为此，需要一个入口控制器。

**An Ingress Controller is a service inside a Kubernetes cluster that is configured as type `LoadBalancer` to receive external traffic. The Ingress Controller uses the defined L7 routing rules and L7 policy to route traffic to the service.**

**Ingress Controller 是 Kubernetes 集群内的一项服务，配置为“LoadBalancer”类型以接收外部流量。入口控制器使用定义的 L7 路由规则和 L7 策略将流量路由到服务。**

In the example below, `helloenroute` service receives traffic from the EnRoute Ingress Controller which receives traffic from the external load balancer.

在下面的示例中，`helloenroute` 服务接收来自 EnRoute Ingress Controller 的流量，EnRoute Ingress Controller 接收来自外部负载均衡器的流量。

![](https://d33wubrfki0l68.cloudfront.net/72423798806754ee35c5981f16cd2e08e1a838b4/a21b1/img/enrouteingresslayer.png)

### Advantages of Using EnRoute Ingress Controller Proxy

### 使用 EnRoute 入口控制器代理的优势

There are several distinct advantages of running an Ingress Controller and enforcing policies at Ingress.

运行 Ingress Controller 并在 Ingress 执行策略有几个明显的优势。

- Ingress provides a portable mechanism to enforce policy inside the Kubernetes Cluster. Policies enforced inside a cluster are easier to port across clouds.
- Multiple proxies can be horizontally scaled using Kubernetes service scaling. Elasticity of L7 fabric makes it easier to operate and scale it
- L7 policies can be hosted along with services inside the cluster with cluster-native state storage
- Keeping L7 policy closer to services simplifies policy enforcement and troubleshooting of services and APIs.

- Ingress 提供了一种可移植的机制来在 Kubernetes 集群内实施策略。在集群内实施的策略更容易跨云移植。
- 可以使用 Kubernetes 服务扩展来水平扩展多个代理。 L7 织物的弹性使其更易于操作和扩展
- L7 策略可以与集群内的服务一起托管在集群本地状态存储中
- 使 L7 策略更接近服务可简化服务和 API 的策略实施和故障排除。

#### Plugins for fine-grained traffic control

#### 用于细粒度流量控制的插件

EnRoute uses Envoy as the underlying proxy to provide the L7 ingress function. EnRoute has a modular architecture that closely reflects Envoy’s extensible model.

EnRoute 使用 Envoy 作为底层代理来提供 L7 入口功能。 EnRoute 具有模块化架构，紧密反映了 Envoy 的可扩展模型。

[Plugins/Filters](http://getenroute.io/features/) can be defined at the route level or service level to enforce L7 policies at Ingress. EnRoute provides advanced rate-limiting plugin in the community edition completely free without any limits. [Clocking your APIs and micro-services](http://getenroute.io/blog/why-every-api-needs-a-clock/) using deep L7 state is a critical need and EnRoutes flexible rate-limiting funtion provides immense flexibility to match a variety of rate-limiting use-cases.

[插件/过滤器](http://getenroute.io/features/) 可以在路由级别或服务级别定义，以在 Ingress 执行 L7 策略。 EnRoute在社区版中提供了完全免费的高级限速插件，没有任何限制。 [为您的 API 和微服务计时](http://getenroute.io/blog/why-every-api-needs-a-clock/) 使用深度 L7 状态是一项关键需求，而 EnRoutes 灵活的速率限制功能提供了巨大的灵活匹配各种限速用例。

[EnRoute Enterprise](http://getenroute.io/features/) includes support and enterprise plugins that help secure traffic at Kubernetes Ingress.

[EnRoute Enterprise](http://getenroute.io/features/) 包括帮助保护 Kubernetes Ingress 流量的支持和企业插件。

Tip: To quickly skim through the article, read through the bold text to get an overall understanding. 

提示：要快速浏览文章，请通读粗体文本以获得整体理解。

