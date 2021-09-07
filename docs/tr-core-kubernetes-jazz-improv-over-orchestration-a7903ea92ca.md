# Core Kubernetes: Jazz Improv over Orchestration

[May 30, 2017](http://blog.heptio.com/core-kubernetes-jazz-improv-over-orchestration-a7903ea92ca?source=post_page-----a7903ea92ca--------------------------------)·7 min read

This is the first in a series of blog posts that details some of the inner workings of Kubernetes. If you are simply an operator or user of Kubernetes you don’t necessarily need to understand these details. But if you prefer depth-first learning and really want to understand the details of how things work, this is for you.

这是一系列博客文章中的第一篇，详细介绍了 Kubernetes 的一些内部工作原理。如果您只是 Kubernetes 的操作员或用户，则不一定需要了解这些细节。但是，如果您更喜欢深度优先学习并且真的想了解事物工作原理的细节，那么这适合您。

This article assumes a working knowledge of Kubernetes. I’m not going to define what Kubernetes is or the core components (e.g. Pod, Node, Kubelet).

本文假设您具备 Kubernetes 的应用知识。我不打算定义 Kubernetes 是什么或核心组件（例如 Pod、Node、Kubelet）。

In this article we talk about the core moving parts and how they work with each other to make things happen. The general class of systems like Kubernetes is commonly called _container orchestration_. But orchestration implies there is a central conductor with an up front plan. However, this isn’t really a great description of Kubernetes. Instead, Kubernetes is more like jazz improv. There is a set of actors that are playing off of each other to coordinate and react.

在本文中，我们将讨论核心运动部件以及它们如何相互配合以实现目标。像 Kubernetes 这样的一般系统类别通常称为 _容器编排_。但编排意味着有一个有预先计划的中央指挥。然而，这并不是对 Kubernetes 的一个很好的描述。相反，Kubernetes 更像是爵士即兴表演。有一组演员相互配合以进行协调和反应。

We’ll start by going over the core components and what they do. Then we’ll look at a typical flow that schedules and runs a Pod.

我们将首先介绍核心组件及其功能。然后我们将查看一个典型的调度和运行 Pod 的流程。

![](https://miro.medium.com/max/1400/1*S005AgJ66RndcSACBW1NeQ.png)

# **Datastore: etcd**

# **数据存储：etcd**

etcd is the core state store for Kubernetes. While there are important in-memory caches throughout the system, etcd is considered the system of record.

etcd 是 Kubernetes 的核心状态存储。虽然整个系统都有重要的内存缓存，但 etcd 被认为是记录系统。

Quick summary of etcd: etcd is a clustered database that prizes consistency above partition tolerance. Systems of this class (ZooKeeper, parts of Consul) are patterned after a system developed at Google called [chubby](https://research.google.com/archive/chubby.html). These systems are often called “lock servers” as they can be used to coordinate locking in a distributed systems. Personally, I find that name a bit confusing. The data model for etcd (and chubby) is a simple hierarchy of keys that store simple unstructured values. It actually looks a lot like a file system. Interestingly, at Google, chubby is most frequently accessed using an abstracted `File` interface that works across local files, object stores, etc. The highly consistent nature, however, provides for strict ordering of writes and allows clients to do atomic updates of a set of values.

etcd 的快速总结：etcd 是一个集群数据库，它重视一致性高于分区容忍度。此类系统（ZooKeeper，Consul 的一部分）模仿了 Google 开发的一个名为 [chubby](https://research.google.com/archive/chubby.html) 的系统。这些系统通常被称为“锁服务器”，因为它们可用于协调分布式系统中的锁。就个人而言，我觉得这个名字有点令人困惑。 etcd（和 chubby)的数据模型是一个简单的键层次结构，用于存储简单的非结构化值。它实际上看起来很像一个文件系统。有趣的是，在谷歌，最常使用抽象的“文件”接口访问 chubby，该接口跨本地文件、对象存储等工作。然而，高度一致的性质提供了严格的写入顺序，并允许客户端执行原子更新值集。

Managing state reliably is one of the more difficult things to do in any system. In a distributed system it is even more difficult as it brings in many subtle algorithms like [raft](https://en.wikipedia.org/wiki/Raft_(computer_science)) or [paxos](https://en.wikipedia.org/wiki/Paxos_(computer_science)). By using etcd, Kubernetes itself can concentrate on other parts of the system.

可靠地管理状态是在任何系统中都比较困难的事情之一。在分布式系统中，这更加困难，因为它引入了许多微妙的算法，如 [raft](https://en.wikipedia.org/wiki/Raft_(computer_science)) 或 [paxos](https://en.wikipedia).org/wiki/Paxos_(computer_science))。通过使用 etcd，Kubernetes 本身可以专注于系统的其他部分。

The idea of **watch** in etcd (and similar systems) is critical for how Kubernetes works. These systems allow clients to perform a lightweight subscription for changes to parts of the key namespace. Clients get notified immediately when something they are watching changes. This can be used as a coordination mechanism between components of the distributed system. One component can write to etcd and other componenents can immediately react to that change.

etcd（和类似系统）中 **watch** 的想法对于 Kubernetes 的工作方式至关重要。这些系统允许客户端对密钥命名空间的部分更改执行轻量级订阅。当客户正在观看的内容发生变化时，他们会立即收到通知。这可以用作分布式系统组件之间的协调机制。一个组件可以写入 etcd，其他组件可以立即对该更改做出反应。

One way to think of this is as an inversion of the common pubsub mechanisms. In many queue systems, the topics store no real user data but the messages that are published to those topics contain rich data. For systems like etcd the keys (analogous to topics) store the real data while the messages (notifications of changes) contain no unique rich information. In other words, for queues the topics are simple and the messages rich while systems like etcd are the opposite.

一种思考方式是将常见的 pubsub 机制倒置。在许多队列系统中，主题不存储真正的用户数据，但发布到这些主题的消息包含丰富的数据。对于像 etcd 这样的系统，键（类似于主题）存储真实数据，而消息（更改通知）不包含唯一的丰富信息。换句话说，对于队列来说，主题很简单，消息很丰富，而像 etcd 这样的系统则相反。

The common pattern is for clients to mirror a subset of the database in memory and then react to changes of that database. Watches are used as an efficient mechanism to keep that cache up to date. If the watch fails for some reason, the client can fall back to polling at the cost of increased load, network traffic and latency.

常见的模式是客户端在内存中镜像数据库的一个子集，然后对该数据库的更改做出反应。手表被用作保持缓存最新的有效机制。如果监视由于某种原因失败，客户端可以以增加负载、网络流量和延迟为代价回退到轮询。

# **Policy Layer: API Server** 

# **策略层：API 服务器**

The heart of Kubernetes is a component that is, creatively, called the API Server. **This is the only component in the system that talks to etcd.** In fact, etcd is really an implementation detail of the API Server and it is theoretically possible to back Kubernetes with some other storage system.

Kubernetes 的核心是一个组件，创造性地称为 API 服务器。 **这是系统中唯一与 etcd 对话的组件。** 事实上，etcd 确实是 API Server 的一个实现细节，理论上可以用其他一些存储系统支持 Kubernetes。

The API Server is a policy component that provides filtered access to etcd. Its responsibilities are relatively generic in nature and it is currently being broken out so that it can be used as a control plane nexus for other types of systems.

API Server 是一个策略组件，提供对 etcd 的过滤访问。它的职责本质上是相对通用的，目前正在分解，以便它可以用作其他类型系统的控制平面连接。

The main currency of the API Server is a resource. These are exposed via a simple REST API. There is a standard structure to most of these resources that enables some expanded features. The nature and reasoning for that API structure is left as a topic for a future post. Regardless, the API Server allows various components to create, read, write, update and watch for changes of resources.

API Server 的主要货币是一种资源。这些是通过一个简单的 REST API 公开的。大多数这些资源都有一个标准结构，可以启用一些扩展功能。该 API 结构的性质和推理将作为未来帖子的主题。无论如何，API Server 允许各种组件创建、读取、写入、更新和监视资源的变化。

Let’s detail the responsibilities of the API Server:

让我们详细说明 API Server 的职责：

1. **Authentication and authorization.** Kubernetes has a pluggable auth system. There are some built in mechanisms for both authentication users and authorizing those users to access resources. In addition there are methods to call out to external services (potentially self-hosted on Kubernetes) to provide these services. This type of extensiblity is core to how Kubernetes is built.
2. Next, the API Server runs a set of**admission controllers** that can reject or modify requests. These allow policy to be applied and default values to be set. This is a critical place for making sure that the data entering the system is valid while the API Server client is still waiting for request confirmation. While these admission controllers are currently compiled in to the API Server, there is ongoing work to make this be another extensibility mechanism.
3. The API server helps with**API versioning**. A critical problem when versioning APIs is to allow for the representation of the resources to evolve. Fields will be added, deprecated, re-organized and in other ways transformed. The API Server stores a “true” representation of a resource in etcd and converts/renders that resource depending on the version of the API being satisfied. Planning for versioning and the evolution of APIs has been a key effort for Kubernetes since early in the project. This is part of what allows Kubernetes to offer a decent [deprecation policy](https://kubernetes.io/docs/reference/deprecation-policy/) relatively early in its lifecycle.

1. **身份验证和授权。** Kubernetes 有一个可插拔的身份验证系统。有一些内置机制用于身份验证用户和授权这些用户访问资源。此外，还有一些方法可以调用外部服务（可能自托管在 Kubernetes 上）来提供这些服务。这种类型的可扩展性是 Kubernetes 构建方式的核心。
2. 接下来，API Server 运行一组**准入控制器**，可以拒绝或修改请求。这些允许应用策略并设置默认值。这是确保在 API Server 客户端仍在等待请求确认时进入系统的数据有效的关键位置。虽然这些准入控制器目前已编译到 API 服务器中，但正在进行工作以使其成为另一种可扩展性机制。
3. API 服务器有助于**API 版本控制**。版本化 API 时的一个关键问题是允许资源的表示不断发展。字段将被添加、弃用、重新组织和以其他方式转换。 API 服务器在 etcd 中存储资源的“真实”表示，并根据满足的 API 版本转换/呈现该资源。自项目早期以来，版本控制和 API 演变的规划一直是 Kubernetes 的一项关键工作。这是允许 Kubernetes 在其生命周期的早期提供体面的 [弃用策略](https://kubernetes.io/docs/reference/deprecation-policy/) 的一部分。

A critical feature of the API Server is that it _also_ supports the idea of **watch**. This means that clients of the API Server can employ the same coordination patterns as with etcd. Most coordination in Kubernetes consists of a component writing to an API Server resource that another component is watching. The second component will then react to changes almost immediately.

API Server 的一个关键特性是它_也_支持 **watch** 的想法。这意味着 API 服务器的客户端可以使用与 etcd 相同的协调模式。 Kubernetes 中的大多数协调包括一个组件写入另一个组件正在监视的 API 服务器资源。然后，第二个组件将几乎立即对变化做出反应。

# **Business Logic: Controller Manager and Scheduler**

# **业务逻辑：控制器管理器和调度器**

The last piece of the puzzle is the code that actually makes the thing work! These are the components that coordinate through the API Server. These are bundled into separate servers called the **Controller Manager** and the **Scheduler**. The choice to break these out was so they couldn’t “cheat”. If the core parts of the system had to talk to the API Server like every other component it would help ensure that we were building an extensible system from the start. The fact that there are just two of these is an accident of history. They could conceivably be combined into one big binary or broken out into a dozen+ separate servers.

最后一块拼图是使这件事真正起作用的代码！这些是通过 API Server 协调的组件。这些被捆绑到称为 **Controller Manager** 和 **Scheduler** 的单独服务器中。打破这些的选择是为了让他们不能“作弊”。如果系统的核心部分必须像其他组件一样与 API 服务器通信，这将有助于确保我们从一开始就构建了一个可扩展的系统。事实上，只有其中两个是历史的偶然。可以想象，它们可以组合成一个大的二进制文件，或者分解成十几个单独的服务器。

The components here do all sorts of things to make the system work. The scheduler, specifically, (a) looks for Pods that aren't assigned to a node (unbound Pods), (b) examines the state of the cluster (cached in memory), (c) picks a node that has free space and meets other constraints, and (d) binds that Pod to a node. 

这里的组件做各种各样的事情来使系统工作。具体而言，调度程序 (a) 查找未分配给节点的 Pod（未绑定的 Pod），(b) 检查集群的状态（缓存在内存中），(c) 选择具有可用空间的节点，并且满足其他约束，并且 (d) 将该 Pod 绑定到一个节点。

Similarly, there is code (“controller”) in the Controller Manager to implement the behavior of a ReplicaSet. (As a reminder, the ReplicaSet ensures that there are a set number of replicas of a Pod Template running at any one time) This controller will watch both the ReplicaSet resource and a set of Pods based on the selector in that resource. It then takes action to create/destroy Pods in order to maintain a stable set of Pods as described in the ReplicaSet. Most controllers follow this type of pattern.

类似地，控制器管理器中有代码（“控制器”）来实现 ReplicaSet 的行为。 （提醒一下，ReplicaSet 确保在任何时候运行的 Pod 模板都有一定数量的副本）该控制器将监视 ReplicaSet 资源和基于该资源中的选择器的一组 Pod。然后它采取行动来创建/销毁 Pod，以维护一组稳定的 Pod，如 ReplicaSet 中所述。大多数控制器遵循这种类型的模式。

# **Node Agent: Kubelet**

# **节点代理：Kubelet**

Finally, there is the agent that sits on the node. This also authenticates to the API Server like any other component. It is responsible for watching the set of Pods that are bound to its node and making sure those Pods are running. It then reports back status as things change with respect to those Pods.

最后，还有位于节点上的代理。这也像任何其他组件一样向 API 服务器进行身份验证。它负责监视绑定到其节点的一组 Pod 并确保这些 Pod 正在运行。然后，它会在与这些 Pod 相关的情况发生变化时报告状态。

# **A Typical Flow**

# **典型流程**

To help understand how this works, let’s work through an example of how things get done in Kubernetes.

为了帮助理解这是如何工作的，让我们通过一个例子来说明事情是如何在 Kubernetes 中完成的。

![](https://miro.medium.com/max/60/1*WDJmiyarVfcsDp6X1-lLFQ.png?q=500)

This sequence diagram shows how a typical flow works for scheduling a Pod. This shows the (somewhat rare) case where a user is creating a Pod directly. More typically, the user will create something like a ReplicaSet and it will be the ReplicaSet that creates the Pod.

此序列图显示了典型流程如何用于调度 Pod。这显示了用户直接创建 Pod 的（有点罕见的）情况。更典型的是，用户将创建类似 ReplicaSet 的东西，而创建 Pod 的将是 ReplicaSet。

The basic flow:

基本流程：

1. The user creates a Pod via the API Server and the API server writes it to etcd.
2. The scheduler notices an “unbound” Pod and decides which node to run that Pod on. It writes that binding back to the API Server.
3. The Kubelet notices a change in the set of Pods that are bound to its node. It, in turn, runs the container via the container runtime (i.e. Docker).
4. The Kubelet monitors the status of the Pod via the container runtime. As things change, the Kubelet will reflect the current status back to the API Server.

1. 用户通过API Server创建Pod，API Server写入etcd。

2. 调度器注意到一个“未绑定”的 Pod 并决定在哪个节点上运行该 Pod。它将该绑定写回 API 服务器。
3. Kubelet 注意到绑定到其节点的一组 Pod 发生了变化。反过来，它通过容器运行时（即 Docker）运行容器。
4. Kubelet 通过容器运行时监控 Pod 的状态。随着情况发生变化，Kubelet 会将当前状态反映回 API Server。

# **Summing Up**

#  **加起来**

By using the API Server as a central coordination point, Kubernetes is able to have a set of components interact with each other in a loosely coupled manner. Hopefully this gives you an idea of how Kubernetes is more jazz improv than orchestration.

通过使用 API Server 作为中央协调点，Kubernetes 能够让一组组件以松散耦合的方式相互交互。希望这能让您了解 Kubernetes 如何比编排更爵士即兴。

Please give us feedback on this article and suggestions for future “under the covers” type pieces. Hit me up on twitter at [@jbeda](https://twitter.com/jbeda) or [@heptio](https://twitter.com/heptio).

请向我们提供有关本文的反馈以及对未来“幕后”类型作品的建议。在 twitter 上联系我 [@jbeda](https://twitter.com/jbeda) 或 [@heptio](https://twitter.com/heptio)。

[**Heptio**](https://blog.heptio.com/?source=post_sidebar--------------------------post_sidebar-----------)



