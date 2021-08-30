# Architecting Kubernetes clusters — choosing the best autoscaling strategy

# 架构 Kubernetes 集群——选择最佳的自动扩展策略

Published in June 2021

2021 年 6 月出版

![Architecting Kubernetes clusters — choosing the best autoscaling strategy](https://learnk8s.io/a/0498c77a1b646af0f3fc42a03f0171fb.svg)

*TL;DR: Scaling pods and nodes in a Kubernetes cluster could take several  minutes with the default settings. Learn how to size your cluster nodes, configure the Horizontal and Cluster Autoscaler, and overprovision your cluster for faster scaling.*

*TL;DR：在默认设置下，扩展 Kubernetes 集群中的 pod 和节点可能需要几分钟时间。了解如何调整集群节点的大小、配置水平和集群自动缩放器以及过度配置集群以加快扩展速度。*

**Table of content:**

**表中的内容：**

- [When autoscaling pods goes wrong](https://learnk8s.io/kubernetes-autoscaling-strategies#when-autoscaling-pods-goes-wrong)
- [How the Cluster Autoscaler works in Kubernetes](https://learnk8s.io/kubernetes-autoscaling-strategies#how-the-cluster-autoscaler-works-in-kubernetes)
- [Exploring pod autoscaling lead time](https://learnk8s.io/kubernetes-autoscaling-strategies#exploring-pod-autoscaling-lead-time)
- [Choosing the optimal instance size for a Kubernetes node](https://learnk8s.io/kubernetes-autoscaling-strategies#choosing-the-optimal-instance-size-for-a-kubernetes-node)
- [Overprovisioning nodes in your Kubernetes cluster](https://learnk8s.io/kubernetes-autoscaling-strategies#overprovisioning-nodes-in-your-kubernetes-cluster)
- [Selecting the correct memory and CPU requests for your Pods](https://learnk8s.io/kubernetes-autoscaling-strategies#selecting-the-correct-memory-and-cpu-requests-for-your-pods)
- [What about downscaling a cluster?](https://learnk8s.io/kubernetes-autoscaling-strategies#what-about-downscaling-a-cluster-)
- [Why not autoscaling based on memory or CPU?](https://learnk8s.io/kubernetes-autoscaling-strategies#why-not-autoscaling-based-on-memory-or-cpu-)

- [当自动缩放 pod 出错时](https://learnk8s.io/kubernetes-autoscaling-strategies#when-autoscaling-pods-goes-wrong)
- [集群自动缩放器在 Kubernetes 中的工作原理](https://learnk8s.io/kubernetes-autoscaling-strategies#how-the-cluster-autoscaler-works-in-kubernetes)
- [探索 pod 自动缩放前置时间](https://learnk8s.io/kubernetes-autoscaling-strategies#exploring-pod-autoscaling-lead-time)
- [选择 Kubernetes 节点的最佳实例大小](https://learnk8s.io/kubernetes-autoscaling-strategies#choosing-the-optimal-instance-size-for-a-kubernetes-node)
- [Kubernetes 集群中的过度配置节点](https://learnk8s.io/kubernetes-autoscaling-strategies#overprovisioning-nodes-in-your-kubernetes-cluster)
- [为您的 Pod 选择正确的内存和 CPU 请求](https://learnk8s.io/kubernetes-autoscaling-strategies#selecting-the-correct-memory-and-cpu-requests-for-your-pods)
- [缩小集群怎么样？](https://learnk8s.io/kubernetes-autoscaling-strategies#what-about-downscaling-a-cluster-)
- [为什么不基于内存或 CPU 进行自动缩放？](https://learnk8s.io/kubernetes-autoscaling-strategies#why-not-autoscaling-based-on-memory-or-cpu-)

In Kubernetes, several things are referred to as "autoscaling", including:

在 Kubernetes 中，有几件事被称为“自动缩放”，包括：

- [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/).
- [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler).
- [Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).

- [Horizo​​ntal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizo​​ntal-pod-autoscale/)。
- [垂直 Pod 自动缩放器](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)。
- [集群自动缩放器](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler)。

Those autoscalers belong to different categories because they address other concerns.

这些自动缩放器属于不同的类别，因为它们解决了其他问题。

The **Horizontal Pod Autoscaler (HPA)** is designed to increase the replicas in your deployments.

**Horizo​​ntal Pod Autoscaler (HPA)** 旨在增加部署中的副本。

As your application receives more traffic, you could have the autoscaler  adjusting the number of replicas to handle more requests.

随着您的应用程序收到更多流量，您可以让自动缩放器调整副本数量以处理更多请求。

- ![The Horizontal Pod Autoscaler (HPA) inspects metrics such as memory and CPU at a regular interval.](https://learnk8s.io/a/b39ab8cff1b9a9f443aa9ef798c9dffb.svg)

  1/2

1/2

  The Horizontal Pod Autoscaler (HPA) inspects metrics such as memory and CPU at a regular interval.

Horizo​​ntal Pod Autoscaler (HPA) 会定期检查内存和 CPU 等指标。

  Next

下一个

The **Vertical Pod Autoscaler (VPA)** is useful when you can't create more copies of your Pods, but you still need to handle more traffic.

当您无法创建更多 Pod 副本但仍需要处理更多流量时，**垂直 Pod 自动缩放器 (VPA)** 非常有用。

As an example, you can't scale a database (easily) only by adding more Pods.

例如，您不能仅通过添加更多 Pod 来（轻松）扩展数据库。

A database might require sharding or configuring read-only replicas.

数据库可能需要分片或配置只读副本。

But you can make a database handle more connections by increasing the memory and CPU available to it.

但是您可以通过增加可用的内存和 CPU 来使数据库处理更多连接。

That's precisely the purpose of the vertical autoscaler — increasing the size of the Pod.

这正是垂直自动缩放器的目的——增加 Pod 的大小。

- ![You can't only increase the number of replicas to scale a database in Kubernetes.](https://learnk8s.io/a/41e2f5317f5dc17ddd0fcb9846ceeae1.svg)

  1/2

1/2

  You can't only increase the number of replicas to scale a database in Kubernetes.

您不能仅通过增加副本数量来扩展 Kubernetes 中的数据库。

  Next

下一个

Lastly, the **Cluster Autoscaler (CA)**.

最后，**Cluster Autoscaler (CA)**。

When your cluster runs low on resources, the Cluster Autoscaler provision a new compute unit and adds it to the cluster.

当您的集群资源不足时，Cluster Autoscaler 会配置一个新的计算单元并将其添加到集群中。

If there are too many empty nodes, the cluster autoscaler will remove them to reduce costs.

如果空节点过多，集群自动扩缩器会移除它们以降低成本。

- ![When you scale your pods in Kubernetes, you might run out of space.](https://learnk8s.io/a/7389f3827b5ba289475c2f03893b0d9e.svg)

  1/3

1/3

  When you scale your pods in Kubernetes, you might run out of space.

当您在 Kubernetes 中扩展 Pod 时，您可能会耗尽空间。

  Next

下一个

While these components all "autoscale" something, they are entirely unrelated to each other.

虽然这些组件都“自动缩放”了一些东西，但它们彼此完全无关。

They all address very different use cases and use other concepts and mechanisms.

它们都针对非常不同的用例并使用其他概念和机制。

And they are developed in separate projects and can be used independently from each other.

它们是在单独的项目中开发的，可以相互独立使用。

**However, scaling your cluster requires fine-tuning the setting of the autoscalers so that they work in concert.**

**但是，扩展集群需要微调自动缩放器的设置，以便它们协同工作。**

Let's have a look at an example.

让我们看一个例子。

## When autoscaling pods goes wrong

## 当自动缩放 Pod 出错时

Imagine having an application that requires and uses 1.5GB of memory and 0.25 vCPU at all times.

想象一下，有一个应用程序始终需要并使用 1.5GB 内存和 0.25 个 vCPU。

You provisioned a cluster with a single node of 8GB and 2 vCPU — it should  be able to fit four pods perfectly (and have a little bit of extra space left).

您配置了一个具有 8GB 和 2 个 vCPU 的单个节点的集群——它应该能够完美地容纳四个 pod（并且还有一点额外的空间）。

![A single node cluster with 8GB of memory and 2 vCPU](https://learnk8s.io/a/458f663479dc10d60ba31648fbb225fb.svg)

You deploy a single Pod and set up: 
您部署单个 Pod 并设置：
1. An **Horizontal Pod Autoscaler** adds a replica every 10 incoming requests (i.e. if you have 40 concurrent requests, it should scale to 4 replicas).
2. A **Cluster Autoscaler** to create more nodes when resources are low.

1. **Horizo​​ntal Pod Autoscaler** 每 10 个传入请求添加一个副本（即如果您有 40 个并发请求，它应该扩展到 4 个副本）。
2. **Cluster Autoscaler** 在资源不足时创建更多节点。

> The Horizontal Pod Autoscaler can scale the replicas in your deployment  using Custom Metrics such as the queries per second (QPS) from an  Ingress controller.

> Horizo​​ntal Pod Autoscaler 可以使用自定义指标来扩展部署中的副本，例如来自 Ingress 控制器的每秒查询数 (QPS)。

You start driving traffic 30 concurrent requests to your cluster and observe the following:

您开始为集群增加 30 个并发请求并观察以下情况：

1. The **Horizontal Pod Autoscaler** starts scaling the Pods.
2. Two more Pods are created.
3. The **Cluster Autoscaler** doesn't trigger — no new node is created in the cluster.

1. **Horizo​​ntal Pod Autoscaler** 开始扩展 Pod。
2. 又创建了两个 Pod。
3. **Cluster Autoscaler** 不会触发——没有在集群中创建新节点。

It makes sense since there's enough space for one more Pod in that node.

这是有道理的，因为该节点中有足够的空间容纳另一个 Pod。

![Scaling three replicas in a single node](https://learnk8s.io/a/2af2246616bdd048e4ef698bb4c9648c.svg)

You further increase the traffic to 40 concurrent requests and observe again:

您进一步将流量增加到 40 个并发请求并再次观察：

1. The **Horizontal Pod Autoscaler** creates one more Pod.
2. The Pod is pending and cannot be deployed.
3. The **Cluster Autoscaler** triggers creating a new node.
4. The node is provisioned in 4 minutes. After that, the pending Pod is deployed.

1. **Horizo​​ntal Pod Autoscaler** 再创建一个 Pod。
2. Pod 处于待定状态，无法部署。
3. **Cluster Autoscaler** 触发创建新节点。
4. 节点在 4 分钟内开通。之后，挂起的 Pod 被部署。

- ![When you scale to four replicas, the fourth replicas isn't deployed in the first node.Instead, it stays "Pending".](https://learnk8s.io/a/1ffab2618c899dd9652e55ea8c973404.svg)

相反，它保持“待定”。](https://learnk8s.io/a/1ffab2618c899dd9652e55ea8c973404.svg)

  1/2 When you scale to four replicas, the fourth replicas isn't deployed in the first node. Instead, it stays *"Pending"*. Next

1/2 当您扩展到四个副本时，第四个副本不会部署在第一个节点中。相反，它保持*“待定”*。下一个

*Why is the fourth Pod not deployed in the first node?*

*为什么第一个节点没有部署第四个Pod？*

Pods deployed in your Kubernetes cluster consume resources such as memory, CPU and storage.

部署在 Kubernetes 集群中的 Pod 会消耗内存、CPU 和存储等资源。

However, on the same node, **the operating system and the kubelet require memory and CPU too.**

但是，在同一个节点上，**操作系统和 kubelet 也需要内存和 CPU。**

In a Kubernetes worker node's memory and CPU are divided into:

在 Kubernetes 工作节点中，内存和 CPU 分为：

1. **Resources needed to run the operating system and system daemons** such as SSH, systemd, etc.
2. **Resources necessary to run Kubernetes agents** such as the Kubelet, the container runtime, [node problem detector](https://github.com/kubernetes/node-problem-detector), etc.
3. **Resources available to Pods.**
4. **Resources reserved to the [eviction threshold](https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/#eviction-thresholds)**.

1. **运行操作系统和系统守护进程所需的资源**如SSH、systemd等。
2.**运行Kubernetes代理所需的资源**，如Kubelet、容器运行时、【节点问题检测器】(https://github.com/kubernetes/node-problem-detector)等。
3. **Pod 可用的资源。**
4. **资源保留给[驱逐阈值](https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/#eviction-thresholds)**。

![Resources in a Kubernetes cluster are consumed by Pods, the operating system, kubelet and eviction threshold](https://learnk8s.io/a/bbaadb833f5689978cfdd0d2fcd9b7ac.svg)

As you can guess, [all of those quotas are customisable](https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/#eviction-thresholds), but you need to account for them.

正如您所猜测的，[所有这些配额都是可定制的](https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/#eviction-thresholds)，但您需要考虑到它们。

In an 8GB and 2 vCPU virtual machine, you can expect:

在 8GB 和 2 个 vCPU 虚拟机中，您可以预期：

- 100MB of memory and 0.1 vCPU to be reserved for the operating system.
- 1.8GB of memory and 0.07 vCPU to be reserved for the Kubelet.
- 100MB of memory for the eviction threshold.

- 为操作系统保留 100MB 内存和 0.1 个 vCPU。
- 为 Kubelet 保留 1.8GB 内存和 0.07 个 vCPU。
- 100MB 内存用于驱逐阈值。

**The remaining ~6GB of memory and 1.83 vCPU are usable by the Pods.**

**剩余的 ~6GB 内存和 1.83 vCPU 可供 Pod 使用。**

If your cluster runs a DeamonSet such as kube-proxy, you should further reduce the available memory and CPU.

如果您的集群运行 kube-proxy 等 DeamonSet，则应进一步减少可用内存和 CPU。

Considering kube-proxy has requests of 128MB and 0.1 vCPU, only ~5.9GB of memory and 1.73 vCPU are available to run Pods.

考虑到 kube-proxy 有 128MB 的请求和 0.1 个 vCPU，只有约 5.9GB 的内存和 1.73 个 vCPU 可用于运行 Pod。

Running a CNI like Flannel and a log collector such as Fluentd will further reduce your resource footprint.

运行像 Flannel 这样的 CNI 和像 Fluentd 这样的日志收集器将进一步减少你的资源占用。

After accounting for all the extra resources, you have space left for only three pods.

考虑到所有额外资源后，您只剩下三个 Pod 的空间。

```
OS                  100MB, 0.1 vCPU   +
 Kubelet             1.8GB, 0.07 vCPU  +
 Eviction threshold  100MB, 0 vCPU     +
 Daemonsets          128MB, 0.1 vCPU   +
 ======================================
 Used                2.1GB, 0.27 vCPU

 ======================================
 Available to Pods   5.9GB, 1.73 vCPU

 Pod requests        1.5GB, 0.25 vCPU
 ======================================
 Total (4 Pods)        6GB, 1vCPU
 ```

 
The fourth stays "Pending" unless it can be deployed on another node.

第四个保持“待定”，除非它可以部署在另一个节点上。

*Since the Cluster Autoscaler knows that there's no space for a fourth Pod, why doesn't it provision a new node?*

*既然 Cluster Autoscaler 知道没有空间容纳第四个 Pod，为什么不配置一个新节点？*

*Why does it wait for the Pod to be pending before it triggers creating a node?*

*为什么它会在触发创建节点之前等待 Pod 挂起？*

## How the Cluster Autoscaler works in Kubernetes

## Cluster Autoscaler 如何在 Kubernetes 中工作

**The Cluster Autoscaler doesn't look at memory or CPU available when it triggers the autoscaling.**

**集群自动缩放器在触发自动缩放时不会查看可用的内存或 CPU。**

Instead, the Cluster Autoscaler reacts to events and checks for any unschedulable Pods every 10 seconds.

相反，Cluster Autoscaler 会对事件做出反应并每 10 秒检查一次任何不可调度的 Pod。

A pod is unschedulable when the scheduler is unable to find a node that can accommodate it.

当调度程序无法找到可以容纳它的节点时，Pod 是不可调度的。

For example, when a Pod requests 1 vCPU but the cluster has only 0.5 vCPU  available, the scheduler marks the Pod as unschedulable.

例如，当 Pod 请求 1 个 vCPU 但集群只有 0.5 个 vCPU 可用时，调度程序会将 Pod 标记为不可调度。

**That's when the Cluster Autoscaler initiates creating a new node.** 
**那时 Cluster Autoscaler 开始创建新节点。**
The Cluster Autoscaler scans the current cluster and [checks if any of the unschedulable pods would fit on in a new node.](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md# what-are-expanders)

Cluster Autoscaler 扫描当前集群并[检查任何不可调度的 pod 是否适合新节点。](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#什么是扩展器）

If you have a cluster with several node types (often also referred to as  node groups or node pools), the Cluster Autoscaler will pick one of them using the following strategies:

如果您的集群具有多种节点类型（通常也称为节点组或节点池），则集群自动缩放器将使用以下策略选择其中一种：

- **Random** — picks a node type at random. This is the default strategy.
- **Most pods** — selects the node group that would schedule the most pods.
- **Least waste** — selects the node group with the least idle CPU after scale-up.
- **Price** — select the node group that will cost the least (only works for GCP at the moment).
- **Priority** — selects the node group with the highest priority (and you manually assign priorities).

- **Random** — 随机选择一个节点类型。这是默认策略。
- **大多数 pods** — 选择将调度最多 pod 的节点组。
- **Least waste** — 选择扩展后空闲 CPU 最少的节点组。
- **价格** — 选择成本最低的节点组（目前仅适用于 GCP）。
- **Priority** — 选择具有最高优先级的节点组（并且您手动分配优先级）。

Once the node type is identified, the Cluster Autoscaler will call the relevant API to provision a new compute resource.

一旦确定了节点类型，Cluster Autoscaler 将调用相关 API 来提供新的计算资源。

If you're using AWS, the Cluster Autoscaler will provision a new EC2 instance.

如果您使用的是 AWS，集群自动缩放器将预置一个新的 EC2 实例。

On Azure, it will create a new Virtual Machine and on GCP, a new Compute Engine.

在 Azure 上，它将创建一个新的虚拟机，并在 GCP 上创建一个新的计算引擎。

It may take some time before the created nodes appear in Kubernetes.

创建的节点可能需要一些时间才能出现在 Kubernetes 中。

Once the compute resource is ready, the [node is initialised](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet-tls-bootstrapping/) and added to the cluster where unscheduled Pods can be deployed.

计算资源准备就绪后，[节点已初始化](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet-tls-bootstrapping/) 并添加到集群中，其中未调度的 Pod 可以被部署。

**Unfortunately, provisioning new nodes is usually slow.**

**不幸的是，配置新节点通常很慢。**

It might take several minutes to provision a new compute unit.

供应新的计算单元可能需要几分钟时间。

*But let's dive into the numbers.*

*但让我们深入研究数字。*

## Exploring pod autoscaling lead time

## 探索 Pod 自动缩放提前期

The time it takes to create a new Pod on a new Node is determined by four major factors:

在新节点上创建新 Pod 所需的时间由四个主要因素决定：

1. **Horizontal Pod Autoscaler reaction time.**
2. **Cluster Autoscaler reaction time.**
3. **Node provisioning time.**
4. **Pod creation time.**

1. **Horizo​​ntal Pod Autoscaler 反应时间。**
2. **Cluster Autoscaler 反应时间。**
3. **节点配置时间。**
4. **Pod 创建时间。**

By default, [pods' CPU and memory usage is scraped by kubelet every 10 seconds.](https://github.com/kubernetes/kubernetes/blob/2da8d1c18fb9406bd8bb9a51da58d5f8108cb8f7/pkg/kubelet/kubelet.go#L1855)

默认情况下，[pods 的 CPU 和内存使用情况每 10 秒被 kubelet 抓取一次。](https://github.com/kubernetes/kubernetes/blob/2da8d1c18fb9406bd8bb9a51da58d5f8108cb8f7/pkg/kubelet/kubelet.go)#L

[Every minute, the Metrics Server will aggregate those metrics](https://github.com/kubernetes-sigs/metrics-server/blob/master/FAQ.md#how-often-metrics-are-scraped) and expose them to the rest of the Kubernetes API.

[每分钟，Metrics Server 都会聚合这些指标](https://github.com/kubernetes-sigs/metrics-server/blob/master/FAQ.md#how-often-metrics-are-scraped) 并公开它们Kubernetes API 的其余部分。

The Horizontal Pod Autoscaler controller is in charge of checking the metrics and deciding to scale up or down your replicas.

Horizo​​ntal Pod Autoscaler 控制器负责检查指标并决定扩大或缩小您的副本。

By default, the [Horizontal Pod Autoscaler checks Pods metrics every 15 seconds.](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#how-does-the-horizontal-pod- autoscaler-work)

默认情况下，[Horizo​​ntal Pod Autoscaler 每 15 秒检查一次 Pods 指标。](https://kubernetes.io/docs/tasks/run-application/horizo​​ntal-pod-autoscale/#how-does-the-horizo​​ntal-pod-自动缩放工作）

![The Horizontal Pod Autoscaler can take up to 1 minute and a half to trigger the autoscaling](https://learnk8s.io/a/f147f5a02119f3320d03c689397f3b48.svg)

[The Cluster Autoscaler checks for unschedulable Pods in the cluster every 10 seconds.](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-does-scale-up-work )

[Cluster Autoscaler 每 10 秒检查一次集群中不可调度的 Pod。](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-does-scale-up-work )

Once one or more Pods are detected, it will run an algorithm to decide:

一旦检测到一个或多个 Pod，它将运行一种算法来决定：

1. **How many nodes** are necessary to deploy all pending Pods.
2. **What type of node group** should be created.

1. **需要多少个节点**来部署所有待处理的 Pod。
2. **应该创建什么类型的节点组**。

The entire process should take:

整个过程应该：

- **No more than 30 seconds** on clusters with **less than 100 nodes** with up to 30 pods each. The average latency should be about 5 seconds.
- **No more than 60 seconds** on cluster with **100 to 1000 nodes**. The average latency should be about 15 seconds.

- **不超过 30 秒**，在 **少于 100 个节点**且每个节点最多 30 个 Pod 的集群上。平均延迟应该是大约 5 秒。
- **在具有 **100 到 1000 个节点**的集群上不超过 60 秒**。平均延迟应约为 15 秒。

![The Cluster Autoscaler takes 30 seconds to decide to create a small node](https://learnk8s.io/a/54ed073d5b8e6c852ba0a7cb19dedec0.svg)

Then, there's the Node provisioning time, which depends mainly on the cloud provider.

然后是节点配置时间，这主要取决于云提供商。

**It's pretty standard for a new compute resource to be provisioned in 3 to 5 minutes.**

**在 3 到 5 分钟内供应新计算资源是非常标准的。**

![Creating a virtual machine on a cloud provider could take several minutes](https://learnk8s.io/a/fdaa1ad973d840c63de2430ae6efdab6.svg)

Lastly, the Pod has to be created by the container runtime.

最后，Pod 必须由容器运行时创建。

Launching a container shouldn't take more than few milliseconds, but **downloading the container image could take several seconds.**

启动容器不会超过几毫秒，但**下载容器映像可能需要几秒钟。**

If you're not caching your container images, downloading an image from the container registry could take from a couple of seconds up to a minute,  depending on the size and number of layers.

如果您不缓存容器映像，则从容器注册表下载映像可能需要几秒钟到一分钟的时间，具体取决于层的大小和数量。

![Downloading a container image could take time and affect scaling](https://learnk8s.io/a/8c0b05d9016ee14a1a6f9952c9d3aa8a.svg)

So the total timing for trigger the autoscaling when there is no space in the current cluster is:

因此，当前集群中没有空间时触发自动缩放的总时间为：

1. The Horizontal Pod Autoscaler might take up to 1m30s to increase the number of replicas.
2. The Cluster Autoscaler should take less than 30 seconds for a cluster with  less than 100 nodes and less than a minute for a cluster with more than  100 nodes.
3. The cloud provider might take 3 to 5 minutes to create the computer resource. 
1. Horizo​​ntal Pod Autoscaler 可能需要长达 1 分 30 秒来增加副本数量。
2. 对于少于 100 个节点的集群，Cluster Autoscaler 应该花费不到 30 秒的时间，对于超过 100 个节点的集群，集群 Autoscaler 应该花费不到 1 分钟。
3. 云提供商可能需要 3 到 5 分钟来创建计算机资源。
4. The container runtime could take up to 30 seconds to download the container image.

4. 容器运行时可能需要长达 30 秒才能下载容器映像。

In the worse case, with a small cluster, you have:

在最坏的情况下，对于一个小集群，你有：

```
HPA delay:          1m30s +
 CA delay:           0m30s +
 Cloud provider:     4m    +
 Container runtime:  0m30s +
 =========================
 Total               6m30s
 ```

 
With a cluster with more than 100 nodes, the total delay could be up to 7 minutes.

对于超过 100 个节点的集群，总延迟可能高达 7 分钟。

*Are you happy to wait for 7 minutes before you have more Pods to handle a sudden surge in traffic?*

*在您有更多 Pod 来处理突然激增的流量之前，您是否愿意等待 7 分钟？*

*How can you tune the autoscaling to reduce the 7 minutes scaling time if you need a new node?*

*如果需要新节点，如何调整自动缩放以减少 7 分钟的缩放时间？*

You could change:

你可以改变：

- The refresh time for the Horizontal Pod Autoscaler (controlled by the `--horizontal-pod-autoscaler-sync-period` flag, default is 15 seconds).
- The interval for metrics scraping in the Metrics Server (controlled by the `metric-resolution` flag, default 60 seconds).
- How frequently the cluster autoscaler scans for unscheduled Pods (controlled by the `scan-interval` flag, default 10 seconds).
- How you cache the image on the local node ([with a tool such as kube-fledged](https://github.com/senthilrch/kube-fledged)).

- Horizo​​ntal Pod Autoscaler 的刷新时间（由 `--horizo​​ntal-pod-autoscaler-sync-period` 标志控制，默认为 15 秒）。
- 指标服务器中指标抓取的间隔（由`metric-resolution` 标志控制，默认为 60 秒）。
- 集群自动缩放器扫描未调度 Pod 的频率（由 `scan-interval` 标志控制，默认 10 秒）。
- 如何在本地节点上缓存图像（[使用 kube-fledged 等工具](https://github.com/senthilrch/kube-fledged)）。

But even if you were to tune those settings to a tiny number, you will still be limited by the cloud provider provisioning time.

但即使您将这些设置调整为很小的数字，您仍然会受到云提供商供应时间的限制。

*So, how could you fix that?*

*所以，你怎么能解决这个问题？*

Since you can't change the provisioning time, you will need a workaround this time.

由于您无法更改配置时间，因此这次您需要一种解决方法。

You could try two things:

你可以尝试两件事：

1. **Avoid creating new nodes,** if possible.
2. **Creating nodes proactively** so that they are already provisioned when you need them.

1. **尽可能避免创建新节点**。
2. **主动创建节点**，以便在您需要时已配置它们。

*Let's have a look at the options one at a time.*

*让我们一次看一个选项。*

## Choosing the optimal instance size for a Kubernetes node

## 为 Kubernetes 节点选择最佳实例大小

**Choosing the right instance type for your cluster has dramatic consequences on your scaling strategy.**

**为您的集群选择正确的实例类型会对您的扩展策略产生巨大影响。**

*Consider the following scenario.*

*考虑以下场景。*

You have an application that requests 1GB of memory and 0.1 vCPU.

您有一个请求 1GB 内存和 0.1 个 vCPU 的应用程序。

You provision a node that has 4GB of memory and 1 vCPU.

您供应一个具有 4GB 内存和 1 个 vCPU 的节点。

After reserving memory and CPU for the kubelet, operating system and eviction threshold, you are left with ~2.5GB of memory and 0.7 vCPU that can be  used for running Pods.

在为 kubelet、操作系统和驱逐阈值保留内存和 CPU 后，您将拥有约 2.5GB 的内存和 0.7 个可用于运行 Pod 的 vCPU。

![Choosing a smaller instance can affect scaling](https://learnk8s.io/a/22fe8bec231d15671512cf797ff3d1cc.svg)

Your node has space for only two Pods.

您的节点只有两个 Pod 的空间。

Every time you scale your replicas, you are likely to incur in up to 7  minutes delay (the lead time to trigger the Horizontal Pod Autoscaler,  Cluster Autoscaler and provisioning the compute resource on the cloud  provider).

每次扩展副本时，您可能会产生最多 7 分钟的延迟（触发 Horizo​​ntal Pod Autoscaler、Cluster Autoscaler 和在云提供商上配置计算资源的前置时间）。

*Let's have a look at what happens if you decide to use a 64GB memory and 16 vCPU node instead.*

*让我们看看如果您决定使用 64GB 内存和 16 个 vCPU 节点会发生什么。*

After reserving memory and CPU for the kubelet, operating system and eviction threshold, you are left with ~58.32GB of memory and 15.8 vCPU that can  be used for running Pods.

为 kubelet、操作系统和驱逐阈值保留内存和 CPU 后，您将拥有约 58.32GB 的内存和 15.8 个 vCPU，可用于运行 Pod。

**The available space can host 58 Pods, and you are likely to need a new node only when you have more than 58 replicas.**

**可用空间可以托管 58 个 Pod，只有当你有超过 58 个副本时，你才有可能需要一个新节点。**

![Choosing a larger instance can affect scaling](https://learnk8s.io/a/691e6c33c3d1db3bb5906512552c5bac.svg)

Also, every time a node is added to the cluster, several pods can be deployed.

此外，每次向集群中添加一个节点时，都可以部署多个 Pod。

There is less chance to trigger *again* the Cluster Autoscaler (and provisioning new compute units on the cloud provider).

很少有机会*再次*触发集群自动缩放器（并在云提供商上配置新的计算单元）。

Choosing large instance types also has another benefit.

选择大型实例类型还有另一个好处。

**The ratio between resource reserved for Kubelet, operating system and  eviction threshold and available resources to run Pods is greater.**

**为 Kubelet 预留的资源、操作系统和驱逐阈值与运行 Pod 的可用资源之间的比率更大。**

Have a look at this graph that pictures the memory available to pods.

看看这张图，它描绘了 Pod 可用的内存。

![Memory available to pods based on different instance types](https://learnk8s.io/a/6829f5509ca3044e1d3166a1f6218eac.svg)

As the instance size increase, you can notice that (in proportion) the resources available to pods increase.

随着实例大小的增加，您可以注意到（按比例）可用于 Pod 的资源增加。

In other words, you are utilising your resources more efficiently than having two instances of half of the size.

换句话说，与拥有两个大小一半的实例相比，您可以更有效地利用资源。

*Should you select the biggest instance all the time?*

*你应该一直选择最大的实例吗？*

**There's a peak in efficiency dictated by how many Pods you can have on the node.**

**效率的峰值取决于节点上可以拥有的 Pod 数量。**

Some cloud providers limit the number of Pods to 110 (i.e. GKE). Others have limits dictated by the underlying network on a per-instance basis (i.e. AWS).

一些云提供商将 Pod 的数量限制为 110（即 GKE）。其他一些限制是由底层网络基于每个实例（即 AWS）决定的。

> [You can inspect the limits from most cloud providers here.](https://docs.google.com/spreadsheets/d/1RPpyDOLFmcgxMCpABDzrsBYWpPYCIBuvAoUQLwOGoQw/edit#gid=907731238)

> [您可以在此处检查大多数云提供商的限制。](https://docs.google.com/spreadsheets/d/1RPpyDOLFmcgxMCpABDzrsBYWpPYCIBuvAoUQLwOGoQw/edit#gid=907731238)

**And choosing a larger instance type is not always a good option.**

**选择更大的实例类型并不总是一个好的选择。**

You should also consider:

您还应该考虑：

1. **Blast radius** — if you have only a few nodes, then the impact of a failing node is bigger than if you have many nodes.
2. **Autoscaling is less cost-effective** as the next increment is a (very) large node.

1. **爆炸半径**——如果你只有几个节点，那么一个失败节点的影响比你有很多节点的影响更大。
2. **自动缩放的成本效益较低**，因为下一个增量是一个（非常）大的节点。

Assuming you have selected the right instance type for your cluster, you might  still face a delay in provisioning the new compute unit.

假设您为集群选择了正确的实例类型，您在配置新计算单元时可能仍会遇到延迟。

*How can you work around that?*

*你如何解决这个问题？*

*What if instead of creating a new node when it's time to scale, you create the same node ahead of time?*

*如果不是在需要扩展时创建新节点，而是提前创建相同的节点怎么办？*

## Overprovisioning nodes in your Kubernetes cluster 
## 在 Kubernetes 集群中过度配置节点
If you can afford to have a spare node available at all times, you could:

如果您可以负担得起随时可用的备用节点，您可以：

1. Create a node and leave it empty.
2. As soon as there's a Pod in the empty node, you create another empty node.

1. 创建一个节点并将其留空。
2. 一旦空节点中有一个 Pod，你就创建另一个空节点。

**In other words, you teach the autoscaler always to have a spare empty node if you need to scale.**

**换句话说，如果你需要扩展，你教自动扩展器总是有一个备用的空节点。**

- ![When you decide to overprovision a cluster, a node is always empty and ready to deploy Pods.](https://learnk8s.io/a/dfe3abd74845075ae2f9fc9b4cf896c3.svg)

  1/3 When you decide to overprovision a cluster, a node is always empty and ready to deploy Pods. Next

1/3 当您决定过度配置集群时，节点始终为空并准备好部署 Pod。下一个

**It's a trade-off: you incur an extra cost (one empty compute unit available at all times), but you gain in speed.**

**这是一种权衡：你会产生额外的成本（一个空的计算单元始终可用），但速度会提高。**

With this strategy, you can scale your fleet much quicker.

使用此策略，您可以更快地扩展您的车队。

*But there's bad and good news.*

*但也有好消息和坏消息。*

The bad news is that the Cluster Autoscaler doesn't have this functionality built-in.

坏消息是 Cluster Autoscaler 没有内置此功能。

**It cannot be configured to be proactive, and there is no flag to "always provision an empty node".**

**它不能被配置为主动，并且没有“总是提供一个空节点”的标志。**

The good news is that you can still fake it.

好消息是你仍然可以伪造它。

*Let me explain.*

*让我解释。*

**You could run a Deployment with enough requests to reserve an entire node.**

**您可以运行具有足够请求的部署来保留整个节点。**

You could think about this pod as a placeholder — it is meant to reserve space, not use any resource.

您可以将此 pod 视为占位符——它旨在保留空间，而不是使用任何资源。

As soon as a real Pod is created, you could evict the placeholder and deploy the Pod.

一旦创建了真正的 Pod，您就可以驱逐占位符并部署 Pod。

- ![In an overprovisioned cluster you have a Pod as a placeholder with low priority.](https://learnk8s.io/a/5297ee233777a99ef552f288299ffb80.svg)

  1/3

1/3

  In an overprovisioned cluster you have a Pod as a placeholder with low priority.

在过度配置的集群中，您有一个 Pod 作为低优先级的占位符。

  Next

下一个

Notice how this time, you still have to wait 5 minutes for the node to be  added to the cluster, but you can keep using the current node.

请注意，这一次，您仍然需要等待 5 分钟才能将节点添加到集群中，但您可以继续使用当前节点。

In the meantime, a new node is provisioned in the background.

同时，在后台供应了一个新节点。

*How can you achieve that?*

*你怎么能做到这一点？*

**Overprovisioning can be configured using deployment running a pod that sleeps forever.**

**可以使用运行永久休眠的 pod 的部署来配置过度配置。**

overprovision.yaml

过度配置.yaml

```
apiVersion: apps/v1
 kind: Deployment
 metadata:
   name: overprovisioning
 spec:
   replicas: 1
   selector:
     matchLabels:
       run: overprovisioning
   template:
     metadata:
       labels:
         run: overprovisioning
     spec:
       containers:
         - name: pause
           image: k8s.gcr.io/pause
           resources:
             requests:
               cpu: '1739m'
               memory: '5.9G'
 ```

 
**You should pay extra attention to the memory and CPU requests.**

**您应该特别注意内存和 CPU 请求。**

The scheduler uses those values to decide where to deploy a Pod.

调度程序使用这些值来决定部署 Pod 的位置。

In this particular case, they are used to reserve the space.

在这种特殊情况下，它们用于保留空间。

You could provision a single large pod that has roughly the requests matching the available node resources.

您可以配置一个大型 Pod，该 Pod 的请求大致与可用节点资源相匹配。

**Please make sure that you account for resources consumed by the kubelet, operating system, kube-proxy, etc.**

**请确保您考虑了 kubelet、操作系统、kube-proxy 等消耗的资源**

If your node instance is 2 vCPU and 8GB of memory and the available space  for pods is 1.73 vCPU and ~5.9GB of memory, your pause pod should match  the latter.

如果您的节点实例是 2 个 vCPU 和 8GB 内存，并且 Pod 的可用空间是 1.73 个 vCPU 和 ~5.9GB 内存，则您的暂停 Pod 应该与后者匹配。

![Sizing sleep pod in overprovisioned clusters](https://learnk8s.io/a/2866bf31d1616ce105778a5f4a5820a9.svg)

To make sure that the Pod is evicted as soon as a real Pod is created, you can use [Priorities and Preemptions.](https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/)

为了确保在创建真正的 Pod 后立即驱逐 Pod，您可以使用 [Priorities and Preemptions.](https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/)

**Pod Priority indicates the importance of a Pod relative to other Pods.**

**Pod Priority 表示一个 Pod 相对于其他 Pod 的重要性。**

When a Pod cannot be scheduled, the scheduler tries to preempt (evict) lower priority Pods to schedule the Pending pod.

当无法调度 Pod 时，调度程序会尝试抢占（驱逐）较低优先级的 Pod 以调度 Pending Pod。

You can configure Pod Priorities in your cluster with a PodPriorityClass:

您可以使用 PodPriorityClass 在集群中配置 Pod 优先级：

priority.yaml

优先级.yaml

```
apiVersion: scheduling.k8s.io/v1beta1
 kind: PriorityClass
 metadata:
   name: overprovisioning
 value: -1
 globalDefault: false
 description: 'Priority class used by overprovisioning.'
 ```

 
Since the default priority for a Pod is `0` and the `overprovisioning` PriorityClass has a value of `-1`, those Pods are the first to be evicted when the cluster runs out of space.

由于 Pod 的默认优先级为 `0` 并且 `overprovisioning` PriorityClass 的值为 `-1`，当集群空间不足时，这些 Pod 将首先被驱逐。

PriorityClass also has two optional fields: `globalDefault` and `description`.

PriorityClass 还有两个可选字段：`globalDefault` 和 `description`。

- The `description` is a human-readable memo of what the PriorityClass is about.
- The `globalDefault` field indicates that the value of this PriorityClass should be used for Pods without a `priorityClassName`. Only one PriorityClass with `globalDefault` set to `true` can exist in the system.

- `description` 是关于 PriorityClass 内容的可读备忘录。
- `globalDefault` 字段表示这个 PriorityClass 的值应该用于没有 `priorityClassName` 的 Pod。系统中只能存在一个“globalDefault”设置为“true”的PriorityClass。

You can assign the priority to your sleep Pod with:

您可以使用以下命令为 sleep Pod 分配优先级：

overprovision.yaml

过度配置.yaml

```
apiVersion: apps/v1
 kind: Deployment
 metadata:
   name: overprovisioning
 spec:
   replicas: 1
   selector:
     matchLabels:
       run: overprovisioning
   template:
     metadata:
       labels:
         run: overprovisioning
     spec:
       priorityClassName: overprovisioning
       containers:
         - name: reserve-resources
           image: k8s.gcr.io/pause
           resources:
             requests:
               cpu: '1739m'
               memory: '5.9G'
 ```

 
*The setup is complete!*

*设置完成！*

When there are not enough resources in the cluster, the pause pod is preempted, and new pods take their place. 
当集群中没有足够的资源时，暂停 Pod 会被抢占，并由新的 Pod 取而代之。
Since the pause pod become unschedulable, it forces the Cluster Autoscaler to add more nodes to the cluster.

由于暂停 pod 变得不可调度，它会强制集群自动缩放器向集群添加更多节点。

*Now that you're ready to overprovision your cluster, it's worth having a look at optimising your applications for scaling.*

*既然您已准备好过度配置集群，那么值得考虑优化您的应用程序以进行扩展。*

## Selecting the correct memory and CPU requests for your Pods

## 为 Pod 选择正确的内存和 CPU 请求

**The cluster autoscaler makes scaling decisions based on the presence of pending pods.**

**集群自动扩缩器根据挂起的 pod 的存在做出扩缩决策。**

The Kubernetes scheduler assigns (or not) a Pod to a Node based on its memory and CPU requests.

Kubernetes 调度程序根据节点的内存和 CPU 请求将 Pod 分配（或不分配）给节点。

Hence, it's essential to set the correct requests on your workloads, or you  might be triggering your autoscaler too late (or too early).

因此，必须为您的工作负载设置正确的请求，否则您可能会过晚（或过早）触发自动缩放器。

*Let's have a look at an example.*

*让我们看一个例子。*

You decide to profile an application, and you found out that:

您决定分析一个应用程序，并发现：

- Under average load, the application consumes 512MB of memory and 0.25 vCPU.
- At peak, the application should consume up to 4GB of memory and 1 vCPU.

- 在平均负载下，应用程序消耗 512MB 内存和 0.25 个 vCPU。
- 在高峰期，应用程序最多应消耗 4GB 内存和 1 个 vCPU。

![Setting the right memory and CPU requests](https://learnk8s.io/a/d44981930b1e0d843e696dd67f3b3b43.svg)

The limit for your container is 4GB of memory and 1 vCPU.

您的容器的限制是 4GB 内存和 1 个 vCPU。

*However, what about the requests?*

*但是，请求呢？*

The scheduler uses the Pod's memory and CPU requests to select the best node before creating the Pod.

调度器在创建 Pod 之前使用 Pod 的内存和 CPU 请求来选择最佳节点。

So you could:

所以你可以：

1. **Set requests lower than the actual average usage.**
2. Be conservative and **assign requests closer to the limit.**
3. **Set requests to match the actual limits.**

1. **设置低于实际平均使用量的请求。**
2. 保守一点，**分配更接近限制的请求。**
3. **设置请求以匹配实际限制。**

- ![You could assign requests that are lower than the average app consumption.](https://learnk8s.io/a/27a58b9c37d2e1ba3386f02c85c110ab.svg)

  1/3

1/3

  You could assign requests that are **lower** than the average app consumption.

您可以分配**低于**平均应用消耗的请求。

  Next

下一个

**Defining requests lower than the actual usage is problematic since your nodes will be often overcommitted.**

**定义低于实际使用的请求是有问题的，因为您的节点经常会被过度使用。**

As an example, you can assign 256MB of memory as a memory request.

例如，您可以分配 256MB 的内存作为内存请求。

The Kubernetes scheduler can fit twice as many Pods for each node.

Kubernetes 调度程序可以为每个节点安装两倍的 Pod。

However, Pods use twice as much memory in practice and start competing for  resources (CPU) and being evicted (not enough memory on the Node).

然而，Pod 在实践中使用两倍的内存并开始竞争资源 (CPU) 并被驱逐（节点上没有足够的内存）。

![Overcommitting nodes](https://learnk8s.io/a/22fdf4559d43e563b3b9b0472ea68969.svg)

**Overcommitting nodes can lead to excessive evictions, more work for the kubelet and a lot of rescheduling.**

**过度使用节点会导致过多的驱逐、更多的 kubelet 工作和大量的重新调度。**

*What happens if you set the request to the same value of the limit?*

*如果您将请求设置为与限制相同的值会怎样？*

You can set request and limits to the same values.

您可以将请求和限制设置为相同的值。

In Kubernetes, this is often referred to as [Guaranteed Quality of Service class](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#qos-classes) and refers to the fact that it's improbable that the pod will be terminated and evicted.

在 Kubernetes 中，这通常称为 [Guaranted Quality of Service class](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#qos-classes)，指的是事实上，pod 不太可能被终止和驱逐。

The Kubernetes scheduler will reserve the entire CPU and memory for the Pod on the assigned node.

Kubernetes 调度程序将为分配的节点上的 Pod 保留整个 CPU 和内存。

**Pods with Guaranteed Quality of Service are stable but also inefficient.**

**具有服务质量保证的 Pod 稳定但效率低下。**

If your app uses 512MB of memory on average, but you reserve 4GB for it, you have 3.5GB unused most of the time.

如果您的应用平均使用 512MB 的内存，但您为它预留了 4GB，那么大部分时间您有 3.5GB 未使用。

![Overcommitting nodes](https://learnk8s.io/a/3661626fe6a72a79770b9f8e2139e015.svg)

*Is it worth it?*

*这值得么？*

If you want extra stability, yes.

如果你想要额外的稳定性，是的。

If you want efficiency, you might want to lower the requests and leave a gap between those and the limit.

如果您想要效率，您可能希望降低请求并在这些请求与限制之间留出差距。

This is often referred to as **Burstable Quality of Service class** and refers to the fact that the Pod baseline consumption can occasionally burst into using more memory and CPU.

这通常被称为 **Burstable Quality of Service 类**，指的是 Pod 基线消耗偶尔会突然使用更多内存和 CPU。

When your requests match the app's actual usage, the scheduler will pack your pods in your nodes efficiently.

当您的请求与应用程序的实际使用相匹配时，调度程序将有效地将您的 pod 打包到您的节点中。

**Occasionally, the app might require more memory or CPU.**

**有时，应用程序可能需要更多内存或 CPU。**

1. If there are resources in the Node, the app will use them before returning to the baseline consumption.
2. If the node is low on resources, the pod will compete for resources (CPU), and the kubelet might try to evict the Pod (memory).

1.如果Node中有资源，app会在回到基线消耗之前使用它们。
2. 如果节点资源不足，Pod 会争抢资源（CPU），kubelet 可能会尝试驱逐 Pod（内存）。

*Should you use Guaranteed or Burstable quality of Service?*

*您应该使用保证服务质量还是突发服务质量？*

*It depends.*

*这取决于。*

1. **Use Guaranteed Quality of Service (requests equal to limits) when you want to minimise rescheduling and evictions for the Pod.** An excellent example is a Pod for a database.
2. **Use Burstable Quality of Service (requests to match actual average usage)  when you want to optimise your cluster and use the resources wisely.** If you have a web application or a REST API, you might want to use a Burstable Quality of Service.

1. **当您想最小化 Pod 的重新调度和驱逐时，请使用有保证的服务质量（请求等于限制）。** 一个很好的例子是数据库的 Pod。
2. **当您想要优化集群并明智地使用资源时，请使用 Burstable Quality of Service（请求匹配实际平均使用情况）。**如果您有 Web 应用程序或 REST API，您可能想要使用 Burstable服务质量。

*How do you select the correct requests and limits values?*

*您如何选择正确的请求和限制值？*

**You should profile the application and measure memory and CPU consumption when idle, under load and at peak.**

**您应该分析应用程序并测量空闲、负载和高峰时的内存和 CPU 消耗。**

A more straightforward strategy consists of deploying the Vertical Pod Autoscaler and wait for it to suggest the correct values.

更直接的策略包括部署 Vertical Pod Autoscaler 并等待它建议正确的值。

The Vertical Pod Autoscaler collects the data from the Pod and applies a regression model to extrapolate requests and limits.

Vertical Pod Autoscaler 从 Pod 收集数据并应用回归模型来推断请求和限制。

[You can learn more about how to do that in this article.](https://learnk8s.io/setting-cpu-memory-limits-requests)

[您可以在本文中了解有关如何执行此操作的更多信息。](https://learnk8s.io/setting-cpu-memory-limits-requests)

## What about downscaling a cluster? 
## 缩小集群怎么样？
**Every 10 seconds, the Cluster Autoscaler decides to remove a node only when the request utilization falls below 50%.**

**每 10 秒，只有当请求利用率低于 50% 时，Cluster Autoscaler 才会决定删除节点。**

In other words, for all the pods on the same node, it sums the CPU and memory requests.

换句话说，对于同一节点上的所有 Pod，它会汇总 CPU 和内存请求。

If they are lower than half of the node's capacity, the Cluster Autoscaler will consider the current node for downscaling.

如果它们低于节点容量的一半，Cluster Autoscaler 将考虑当前节点进行缩减。

> It's worth noting that the Cluster Autoscaler does not consider actual CPU  and memory usage or limits and instead only looks at resource requests.

> 值得注意的是，Cluster Autoscaler 不考虑实际的 CPU 和内存使用或限制，而只考虑资源请求。

Before the node is removed, the Cluster Autoscaler executes:

在移除节点之前，Cluster Autoscaler 执行：

- [Pods checks](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-types-of-pods-can-prevent-ca-from-removing-a- node) to make sure that the Pods can be moved to other nodes.
- [Nodes checks](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#i-have-a-couple-of-nodes-with-low-utilization-but- they-are-not-scaled-down-why) to prevent nodes from being destroyed prematurely.

- [Pods 检查](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-types-of-pods-can-prevent-ca-from-removing-a-节点）以确保 Pod 可以移动到其他节点。
- [节点检查](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#i-have-a-couple-of-nodes-with-low-utilization-but-它们不是按比例缩小的原因）以防止节点过早被破坏。

If the checks pass, the Cluster Autoscaler will remove the node from the cluster.

如果检查通过，Cluster Autoscaler 将从集群中删除节点。

## Why not autoscaling based on memory or CPU?

## 为什么不基于内存或 CPU 自动缩放？

**CPU or memory-based cluster autoscalers don't care about pods when scaling up and down.**

**CPU 或基于内存的集群自动缩放器在扩展和缩减时不关心 Pod。**

Imagine having a cluster with a single node and setting up the autoscaler to  add a new node with the CPU reaches 80% of the total capacity.

想象一下，有一个只有一个节点的集群，并设置自动缩放器来添加一个新节点，CPU 达到总容量的 80%。

You decide to create a Deployment with 3 replicas.

您决定创建一个具有 3 个副本的部署。

The combined resource usage for the three pods reaches 85% of the CPU.

三个 Pod 的综合资源使用率达到了 CPU 的 85%。

A new node is provisioned.

提供了一个新节点。

*What if you don't need any more pods?*

*如果您不需要更多豆荚怎么办？*

You have a full node idling — not great.

你有一个完整的节点空闲——不是很好。

**Usage of these type of autoscalers with Kubernetes is discouraged.**

**不鼓励在 Kubernetes 中使用这些类型的自动缩放器。**

## Summary

＃＃ 概括

Defining and implementing a successful scaling strategy in Kubernetes requires you to master several subjects:

在 Kubernetes 中定义和实施成功的扩展策略需要您掌握几个主题：

- Allocatable resources in Kubernetes nodes.
- Fine-tuning refresh intervals for Metrics Server, Horizontal Pod Autoscaler and Cluster Autoscalers.
- Architecting cluster and node instance sizes.
- Container image caching.
- Application benchmarking and profiling.

- Kubernetes 节点中的可分配资源。
- 微调 Metrics Server、Horizo​​ntal Pod Autoscaler 和 Cluster Autoscalers 的刷新间隔。
- 构建集群和节点实例大小。
- 容器图像缓存。
- 应用程序基准测试和分析。

But with the proper monitoring tool, you can iteratively test your scaling  strategy and tune the speed and costs of your cluster. 
但是使用适当的监控工具，您可以反复测试您的扩展策略并调整集群的速度和成本。
