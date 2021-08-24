## Everything you Need to Know about Kubernetes Quality of Service (QoS) Classes

## 您需要了解的有关 Kubernetes 服务质量 (QoS) 类的所有信息

In this blog post we will take a deep dive into Kubernetes QOS classes. We will start by looking at the factors that determine whether a pod is assigned a Guaranteed, Burstable or BestEffort QOS class. We will then look at how QOS class impacts the way pods are scheduled by the Kubernetes Scheduler, how it impacts the eviction order of pods by the Kubelet as well as what happens to them during node OOM events.

在这篇博文中，我们将深入探讨 Kubernetes QOS 类。我们将首先查看决定 pod 是否被分配有保证、突发或 BestEffort QOS 类的因素。然后，我们将看看 QOS 类如何影响 Kubernetes 调度程序调度 Pod 的方式，它如何影响 Kubelet 驱逐 Pod 的顺序，以及在节点 OOM 事件期间它们会发生什么。

March 22, 2019 From: https://www.replex.io/blog/everything-you-need-to-know-about-kubernetes-quality-of-service-qos-classes

2019 年 3 月 22 日来自：https://www.replex.io/blog/everything-you-need-to-know-about-kubernetes-quality-of-service-qos-classes

Quality of Service (QoS) class is a Kubernetes concept which determines the scheduling and eviction priority of pods. [QoS class](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#qos-classes) is used by the Kubernetes scheduler to make decisions about scheduling pods onto nodes.

服务质量 (QoS) 类是 Kubernetes 的一个概念，它决定了 pod 的调度和驱逐优先级。 [QoS 类](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#qos-classes) 被 Kubernetes 调度器用来决定将 pod 调度到节点上。

Kubelet uses it to govern both the order in which pods are evicted as well as to allow more complex pod placement decisions using advanced [CPU management policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management -policies/).

Kubelet 使用它来管理 pod 被驱逐的顺序，以及使用高级 [CPU 管理策略](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management) 允许更复杂的 pod 放置决策-政策/）。

QoS class is assigned to pods by Kubernetes itself. DevOps can, however, control the QoS class assigned to a pod by playing around with [resource requests and limits](/blog/5-ways-to-manage-your-kubernetes-resource-usage) for individual containers inside the pod.

QoS 类由 Kubernetes 本身分配给 Pod。然而，DevOps 可以通过为 pod 内的各个容器处理 [资源请求和限制](/blog/5-ways-to-manage-your-kubernetes-resource-usage) 来控制分配给 pod 的 QoS 类。

There are three QoS classes in Kubernetes:

Kubernetes 中有 3 个 QoS 类：

*   Guaranteed
*   Burstable
*   BestEffort

* 保证
* 可爆
* 尽力而为

Let’s go through the different QoS classes and see how they work together with the Kubernetes Scheduler and Kubelet.

让我们来看看不同的 QoS 类，看看它们如何与 Kubernetes Scheduler 和 Kubelet 一起工作。

Guaranteed
----------

### How Can I Assign a QoS class of Guaranteed to a Pod?

保证
### 如何为 Pod 分配有保证的 QoS 等级？

For a pod to be placed in the Guaranteed QoS class, every container in the pod must have a CPU and memory limit. Kubernetes will automatically assign CPU and memory request values (that are equal to the CPU and memory limit values) to the containers inside this pod and will assign it the Guaranteed QoS class.

要将 Pod 置于保证 QoS 类中，Pod 中的每个容器都必须具有 CPU 和内存限制。 Kubernetes 将自动为该 Pod 内的容器分配 CPU 和内存请求值（等于 CPU 和内存限制值），并将为其分配保证 QoS 类。

Pods with explicit and equal values for both CPU requests and limits and memory requests and limits are also placed in the Guaranteed QoS class.

具有 CPU 请求和限制以及内存请求和限制的显式和相等值的 Pod 也放置在有保证的 QoS 类中。

### How does the Kubernetes Scheduler Handle Guaranteed Pods?

### Kubernetes 调度程序如何处理有保证的 Pod？

The [Kubernetes scheduler](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/) assigns Guaranteed pods only to nodes which have enough resources to fulfil their CPU and memory requests. The Scheduler does this by ensuring that the sum of both memory and CPU requests for all containers (running and newly scheduled) is lower than the total capacity of the node.

[Kubernetes 调度程序](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/) 仅将有保证的 pod 分配给具有足够资源来满足其 CPU 和内存请求的节点。调度器通过确保所有容器（正在运行的和新调度的）的内存和 CPU 请求的总和低于节点的总容量来做到这一点。

### Can Guaranteed Pods be Allocated Exclusive CPU Cores?

### 可以为有保证的 Pod 分配独占 CPU 内核吗？

The default CPU management policy of Kubernetes is “None”. Under this policy Guaranteed pods run in the shared CPU pool on a node. The shared CPU pool contains all the CPU resources on the node minus the ones reserved by the Kubelet using **\--kube-reserved** or **\--system-reserved**.

Kubernetes 的默认 CPU 管理策略是“无”。在此策略下，有保证的 Pod 在节点上的共享 CPU 池中运行。共享 CPU 池包含节点上的所有 CPU 资源，减去 Kubelet 使用 **\--kube-reserved** 或​​ **\--system-reserved** 保留的资源。

Guaranteed pods can, however, be allocated exclusive use of CPU cores with a static CPU management policy. To be granted exclusive use of CPU cores under this policy, Guaranteed pods also need to have CPU request values in integers. Guaranteed pods with fractional CPU request values will still run in the shared CPU pool under the static CPU management policy.

但是，可以通过静态 CPU 管理策略为有保证的 Pod 分配 CPU 内核的独占使用权。要根据此策略获得 CPU 内核的独占使用权，有保证的 Pod 还需要具有整数形式的 CPU 请求值。在静态 CPU 管理策略下，具有部分 CPU 请求值的有保证的 Pod 仍将在共享 CPU 池中运行。

### What Stops a Guaranteed Pod from being Scheduled onto a Node?

### 是什么阻止了有保证的 Pod 被调度到节点上？

Guaranteed pods cannot be scheduled onto nodes for which the Kubelet reports a [DiskPressure node condition](https://kubernetes.io/docs/concepts/architecture/nodes/#condition). DiskPressure is a node condition which is triggered when the available disk space and inodes on either the node’s root filesystem or image filesystem hit an eviction threshold. When the node reports a DiskPressure condition, the Scheduler stops scheduling any new Guaranteed pods onto the node.

无法将有保证的 Pod 调度到 Kubelet 报告 [DiskPressure 节点条件](https://kubernetes.io/docs/concepts/architecture/nodes/#condition) 的节点上。 DiskPressure 是一种节点条件，当节点的根文件系统或映像文件系统上的可用磁盘空间和 inode 达到驱逐阈值时会触发。当节点报告 DiskPressure 条件时，调度程序停止将任何新的有保证的 pod 调度到节点上。

Burstable
---------

### How Can I assign a QoS class of Burstable to a Pod?

可爆
### 如何为 Pod 分配 Burstable 的 QoS 类？

A pod is assigned a [Burstable QoS](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#create-a-pod-that-gets-assigned-a-qos -class-of-burstable) class if at least one container in that pod has a memory or CPU request.

一个 Pod 被分配了一个 [Burstable QoS](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#create-a-pod-that-gets-assigned-a-qos -class-of-burstable) 类，如果该 Pod 中至少有一个容器有内存或 CPU 请求。

### How does the Kubernetes Scheduler Handle Burstable Pods?

### Kubernetes 调度程序如何处理 Burstable Pod？

The Kubernetes scheduler will not be able to ensure that Burstable pods are placed onto nodes that have enough resources for them.

Kubernetes 调度程序将无法确保 Burstable pod 被放置在具有足够资源的节点上。

### Can Burstable Pods be Allocated Exclusive CPU cores?

### Burstable Pod 可以分配独占 CPU 内核吗？

Burstable pods run in the shared resources pool of nodes along with BestEffort and Guaranteed pods under the default “None” CPU management policy. It is not possible to allocate exclusive CPU cores to Burstable pods.

Burstable Pod 与 BestEffort 和 Guaranteed Pod 一起在默认的“None”CPU 管理策略下运行在节点的共享资源池中。无法将独占 CPU 内核分配给 Burstable pod。

### What Stops a Burstable Pod from Being Scheduled onto a Node? 
### 是什么阻止了 Burstable Pod 被调度到节点上？
As with Guaranteed pods, BestEffort pods also cannot be scheduled onto nodes under DiskPressure. The Kubernetes scheduler will not schedule any new Burstable pods onto a node with the condition DiskPressure.

与有保证的 pod 一样，BestEffort pod 也不能调度到 DiskPressure 下的节点上。 Kubernetes 调度程序不会将任何新的 Burstable pod 调度到具有 DiskPressure 条件的节点上。

BestEffort
----------

### How can I Assign a QoS Class of BestEffort to a Pod?

尽力而为
### 如何为 Pod 分配 BestEffort 的 QoS 等级？

A pod is assigned a [BestEffort QoS class](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#create-a-pod-that-gets-assigned-a- qos-class-of-besteffort) if none of it's containers have CPU or memory requests and limits.

一个 Pod 被分配了一个 [BestEffort QoS 类](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#create-a-pod-that-gets-assigned-a- qos-class-of-besteffort），如果它的容器都没有 CPU 或内存请求和限制。

### How does the Kubernetes Scheduler handle BestEffort pods?

### Kubernetes 调度程序如何处理 BestEffort pod？

BestEffort pods are not guaranteed to be placed on to pods that have enough resources for them. They are, however, able to use any amount of free CPU and memory resources on the node. This can at times lead to resource contention with other pods, where BestEffort pods hog resources and do not leave enough resource headroom for other pods to consume resources within resource limits.

BestEffort pod 不能保证被放置在有足够资源的 pod 上。但是，它们能够使用节点上任意数量的空闲 CPU 和内存资源。这有时会导致与其他 pod 的资源争用，其中 BestEffort pod 会占用资源，并且没有为其他 pod 留出足够的资源空间来消耗资源限制内的资源。

### Can BestEffort Pods be Allocated Exclusive CPU cores?

### BestEffort Pod 可以分配独占 CPU 内核吗？

As with pods which have a Burstable QoS class, BestEffort pods also run in the shared resources pool on a node and cannot be granted exclusive CPU resource usage.

与具有 Burstable QoS 类的 Pod 一样，BestEffort Pod 也在节点上的共享资源池中运行，并且不能被授予独占 CPU 资源使用权。

### What Stops a BestEffort Pod from Being Scheduled onto a Node?

### 是什么阻止了 BestEffort Pod 被调度到节点上？

BestEffort pods cannot be scheduled onto nodes under both DiskPressure and [MemoryPressure](https://kubernetes.io/docs/concepts/architecture/nodes/#condition). A node reports MemoryPressure condition if it has lower levels of memory available then a predefined threshold. The Kubernetes Scheduler will, in turn, stop scheduling any new BestEffort pods onto the node.

BestEffort pod 不能同时调度到 DiskPressure 和 [MemoryPressure](https://kubernetes.io/docs/concepts/architecture/nodes/#condition) 下的节点上。如果节点的可用内存级别低于预定义的阈值，则它会报告 MemoryPressure 条件。反过来，Kubernetes 调度程序将停止将任何新的 BestEffort pod 调度到节点上。

Next we will look at how Kubelet handles evictions for pods of all three QoS classes. We will also see how a pod's QOS class impacts what happens to it when the node runs out of memory.

接下来我们将看看 Kubelet 如何处理所有三个 QoS 类的 pod 的驱逐。我们还将看到 pod 的 QOS 类如何影响节点内存不足时发生的情况。

How are Guaranteed, Burstable and BestEffort Pods Evicted by the Kubelet?
-------------------------------------------------------------------------

[Pod evictions](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/#eviction-policy) are initiated by the Kubelet when the node starts running low on compute resources. These evictions are meant to reclaim resources to avoid a system out of memory (OOM) event. DevOps can specify thresholds for resources which when breached trigger pod evictions by the Kubelet.

Kubelet 是如何驱逐有保证的、突发的和 BestEffort Pod 的？
[Pod evictions](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/#eviction-policy) 当节点开始计算资源不足时由 Kubelet 发起。这些驱逐旨在回收资源以避免系统内存不足 (OOM) 事件。 DevOps 可以为资源指定阈值，当违反这些阈值时，Kubelet 会触发 Pod 驱逐。

The QoS class of a pod does affect the order in which it is chosen for eviction by the Kubelet. Kubelet first evicts BestEffort and Burstable pods using resources above requests. The order of eviction depends on the priority assigned to each pod and the amount of resources being consumed above request.

Pod 的 QoS 等级确实会影响 Kubelet 选择它进行驱逐的顺序。 Kubelet 首先使用请求之上的资源驱逐 BestEffort 和 Burstable pod。驱逐的顺序取决于分配给每个 pod 的优先级以及在请求之上消耗的资源量。

Guaranteed and Burstable pods not exceeding resource requests are evicted next based on which ones have the lowest priority.

不超过资源请求的有保证的和突发的 pod 将根据优先级最低的 pod 被逐出。

Both Guaranteed and Burstable pods whose resource usage is lower than the requested amount are never evicted because of the resource usage of another pod. They might, however, be evicted if system daemons start using more resources than reserved. In this case, Guaranteed and Burstable pods with the lowest priority are evicted first.

资源使用量低于请求数量的保证和突发 Pod 都不会因为另一个 Pod 的资源使用而被驱逐。然而，如果系统守护进程开始使用比保留更多的资源，它们可能会被驱逐。在这种情况下，优先级最低的保证和突发 pod 将首先被驱逐。

When responding to DiskPressure node condition, the Kubelet [first evicts BestEffort pods followed by Burstable pods](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/#evicting-end-user-pods ). Only when there are no BestEffort or Burstable pods left are Guaranteed pods evicted.

当响应 DiskPressure 节点条件时，Kubelet [首先驱逐 BestEffort pod，然后是 Burstable pod](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/#evicting-end-user-pods ）。只有当没有剩余的 BestEffort 或 Burstable pod 时，保证 pod 才会被驱逐。

What is the Node Out of Memory (OOM) Behaviour for Guaranteed, Burstable and BestEffort pods?
---------------------------------------------------------------------------------------------

If the node runs out of memory before the Kubelet can reclaim it, the [oom\_killer kicks in](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/#node-oom- behavior) to kill containers based on their oom\_score. The oom\_score is calculated by the oom\_killer for each container and is based on the percentage of memory the container uses on the node as compared to what it requested plus the oom\_score\_adj score.

保证、突发和 BestEffort pod 的节点内存不足 (OOM) 行为是什么？
如果节点在 Kubelet 可以回收它之前耗尽内存，[oom\_killer 开始](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/#node-oom-行为）根据它们的 oom\_score 杀死容器。 oom\_score 由每个容器的 oom\_killer 计算，并基于容器在节点上使用的内存与其请求的内存百分比加上 oom\_score\_adj 分数。

The [oom\_score\_adj](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/#node-oom-behavior) for each container is governed by the QoS class of the pod it belongs to. For a container inside a Guaranteed pod the oom\_score\_adj is “-998”, for a Burstable pod container it is “1000” and for a BestEffort pod container “min(max(2, 1000 - (1000 \* memoryRequestBytes) / machineMemoryCapacityBytes), 999)”. 
每个容器的 [oom\_score\_adj](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/#node-oom-behavior) 由 pod 的 QoS 类管理它属于。对于有保证的 pod 内的容器，oom\_score\_adj 是“-998”，对于 Burstable pod 容器是“1000”，对于 BestEffort pod 容器是“min(max(2, 1000 - (1000 \* memoryRequestBytes)) /machineMemoryCapacityBytes), 999)”。
The oom\_killer first terminates containers that belong to pods with the lowest QoS class and which most exceed the requested resources. This means that containers belonging to a pod with a better QoS class (like Guaranteed) have a lower probability of being killed than one’s with Burstable or BestEffort QoS class.

oom\_killer 首先终止属于具有最低 QoS 等级的 pod 并且最超出所请求资源的容器。这意味着属于具有更好 QoS 等级（如保证）的 pod 的容器被杀死的可能性低于具有 Burstable 或 BestEffort QoS 等级的容器。

This, however, is not true of all cases. Since the oom\_killer also considers memory usage vs request, a container with a better QoS class might have a higher oom\_score because of excessive memory usage and thus might be killed first.

然而，这并非适用于所有情况。由于 oom\_killer 还考虑内存使用与请求，具有更好 QoS 类的容器可能会因为内存使用过多而具有更高的 oom\_score，因此可能会首先被杀死。

Conclusion
----------

QOS class determines the order in which the [Kubernetes scheduler](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/) schedules pods as well as the order in which they are evicted by the Kubelet. DevOps can influence the QOS class of a pod by assigning resource limits and/or requests to individual containers belonging to the pod.

结论
QOS 类决定了 [Kubernetes scheduler](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/) 调度 pod 的顺序以及它们被驱逐的顺序通过 Kubelet。 DevOps 可以通过为属于 Pod 的各个容器分配资源限制和/或请求来影响 Pod 的 QOS 类别。

The QOS class of a pod can in some cases impact the resource utilization of individual nodes. Since resource requests and limits are mostly set based on guesstimates there are cases where Guaranteed and Burstable QOS pods have a resource footprint that is much higher than required. This can lead to sitautions where pods do not utilize the requested resources efficiently. 
在某些情况下，Pod 的 QOS 等级会影响单个节点的资源利用率。由于资源请求和限制主要是根据猜测设置的，因此在某些情况下，保证和突发 QOS pod 的资源占用远高于所需。这可能导致 pod 无法有效利用请求的资源的情况。
