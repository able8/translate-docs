# Dive Deep Into Resource Requests and Limits in Kubernetes

# 深入了解 Kubernetes 中的资源请求和限制

### This article will be helpful for you to understand how Kubernetes requests and limits work, and why they can  work in an expected way.

### 本文将有助于您了解 Kubernetes 请求和限制是如何工作的，以及为什么它们可以以预期的方式工作。

As you create resources in a Kubernetes cluster, you may have encountered the following scenarios:

在 Kubernetes 集群中创建资源时，您可能会遇到以下场景：

1. No CPU requests or low CPU requests specified for workloads, which  means more Pods “seem” to be able to work on the same node. During  traffic bursts, your CPU is maxed out with a longer delay while some of  your machines may have a CPU soft lockup.
2. Likewise, no memory requests or low memory requests specified for  workloads. Some Pods, especially those running Java business apps, will  keep restarting while they can actually run normally in local tests.
3. In a Kubernetes cluster, workloads are usually not scheduled evenly across nodes. In most cases, in particular, memory resources are  unevenly distributed, which means some nodes can see much higher memory  utilization than other nodes. As the de facto standard in container  orchestration, Kubernetes should have an effective scheduler that  ensures the even distribution of resources. But, is it really the case?



1. 没有为工作负载指定 CPU 请求或低 CPU 请求，这意味着更多 Pod “似乎”能够在同一个节点上工作。在流量爆发期间，您的 CPU 会以较长的延迟被最大化，而您的某些机器可能会出现 CPU 软锁定。
2. 同样，没有为工作负载指定内存请求或低内存请求。一些 Pod，尤其是那些运行 Java 业务应用程序的 Pod，会在本地测试中实际上可以正常运行的同时不断重启。
3. 在 Kubernetes 集群中，工作负载通常不会跨节点均匀调度。特别是在大多数情况下，内存资源分布不均，这意味着某些节点的内存利用率比其他节点高得多。作为容器编排的事实标准，Kubernetes 应该有一个有效的调度器来确保资源的均匀分布。但是，真的是这样吗？

Generally, cluster administrators can do nothing but restart the  cluster if the above issues happen amid traffic bursts when all of your  machines hang and SSH login fails. In this article, we will dive deep  into Kubernetes requests and limits by analyzing possible issues and  discussing the best practices for them. If you are also interested in  the underlying mechanism, you can also find the analysis from the  perspective of source code. Hopefully, this article will be helpful for  you to understand how Kubernetes requests and limits work, and why they  can work in the expected way.

通常，如果在所有机器都挂起并且 SSH 登录失败的情况下，如果在流量突发的情况下发生上述问题，则集群管理员只能重启集群。在本文中，我们将通过分析可能的问题并讨论它们的最佳实践，深入探讨 Kubernetes 的请求和限制。如果你也对底层机制感兴趣，也可以从源码的角度找分析。希望本文能帮助您了解 Kubernetes 请求和限制是如何工作的，以及为什么它们可以按预期方式工作。

## Concepts

## 概念

To make full use of resources in a Kubernetes cluster and improve  scheduling efficiency, Kubernetes uses requests and limits to control  resource allocation for containers. Each container can have its own  requests and limits. These two parameters are specified by `resources.requests` and `resources.limits`. Generally speaking, requests are more important in scheduling while limits are more important in running.

为了充分利用 Kubernetes 集群中的资源，提高调度效率，Kubernetes 使用请求和限制来控制容器的资源分配。每个容器都可以有自己的请求和限制。这两个参数由`resources.requests`和`resources.limits`指定。一般来说，请求在调度中更重要，而限制在运行中更重要。

```
resources:
    requests:
        cpu: 50m
        memory: 50Mi
    limits:
        cpu: 100m
        memory: 100Mi
```

Requests define the smallest amount of resources a container needs. For example, for a container running Spring Boot business, the requests  specified must be the minimum amount of resources that a Java Virtual  Machine (JVM) needs to consume in the container image. If you only  specify a low memory request, the odds are that the Kubernetes scheduler tends to schedule the Pod to the node which doesn’t have sufficient  resources to run the JVM. Namely, the Pod cannot use more memory which  the JVM bootup process needs. As a result, the Pod keeps restarting.

请求定义了容器需要的最小资源量。例如，对于运行 Spring Boot 业务的容器，指定的请求必须是 Java 虚拟机（JVM）在容器镜像中需要消耗的最小资源量。如果你只指定一个低内存请求，很可能 Kubernetes 调度器倾向于将 Pod 调度到没有足够资源来运行 JVM 的节点上。也就是说，Pod 无法使用 JVM 启动过程所需的更多内存。结果，Pod 不断重启。

On the other hand, limits determine the maximum amount of resources  that a container can use, preventing resource shortage or machine  crashes due to excessive resource consumption. If it is set to `0`, it means no resource limit for the container. In particular, if you set `limits` without specifying `requests`, Kubernetes considers the value of `requests` is the same as that of `limits` by default.

另一方面，限制决定了容器可以使用的最大资源量，防止资源短缺或由于资源消耗过多而导致机器崩溃。如果设置为“0”，则表示容器没有资源限制。特别是，如果你设置了 `limits` 而没有指定 `requests`，Kubernetes 会默认认为 `requests` 的值与 `limits` 的值相同。

Requests and limits apply to two types of resources - compressible  (for example, CPU) and incompressible (for example, memory). For  incompressible resources, appropriate limits are extremely important.

请求和限制适用于两种类型的资源 - 可压缩（例如 CPU）和不可压缩（例如内存）。对于不可压缩资源，适当的限制非常重要。

Here is a brief summary of requests and limits:

以下是请求和限制的简要摘要：

- If services in a Pod are using more CPU resources than the limits  specified, the Pod will be restricted but will not be killed. If no  limit is set, a Pod can use all idle CPU resources.
- If a Pod is using more memory resources than the limits specified,  container processes in the Pod will be killed due to OOM. In this case,  Kubernetes tends to restart the container on the original node or simply create another Pod.
- 0 <= requests <=Node Allocatable; requests <= limits <= Infinity.

- 如果 Pod 中的服务使用的 CPU 资源超过指定的限制，则 Pod 将被限制但不会被杀死。如果没有设置限制，Pod 可以使用所有空闲的 CPU 资源。
- 如果 Pod 使用的内存资源超过指定的限制，则 Pod 中的容器进程将因 OOM 而被杀死。在这种情况下，Kubernetes 倾向于在原始节点上重新启动容器或简单地创建另一个 Pod。
- 0 <= 请求 <= 节点可分配；请求 <= 限制 <= 无限。

## Scenario Analysis

## 场景分析

After we look at the concepts of requests and limits, let’s go back to the three scenarios mentioned at the beginning.

看完requests和limits的概念，我们再回到开头提到的三个场景。

### Scenario 1 

### 场景 1

First and foremost, you need to know that CPU resources and memory  resources are completely different. CPU resources are compressible. The  distribution and management of CPU are based on Completely Fair  Scheduler (CFS) and Cgroups. To put it in a simple way, if the Service  in a Pod is using more CPU resources than the CPU limits specified, it  will be throttled by Kubernetes. For Pods without CPU limits, once idle  CPU resources are running out, the amount of CPU resources allocated  before will gradually decrease. In both situations, ultimately, Pods  will be unable to handle external requests, resulting in a longer delay  and response time.

首先，你需要知道 CPU 资源和内存资源是完全不同的。 CPU 资源是可压缩的。 CPU的分配和管理基于完全公平调度器（CFS）和Cgroups。简单来说，如果 Pod 中的 Service 使用的 CPU 资源超过了指定的 CPU 限制，它将被 Kubernetes 限制。对于没有 CPU 限制的 Pod，一旦空闲的 CPU 资源耗尽，之前分配的 CPU 资源量会逐渐减少。在这两种情况下，最终 Pod 都将无法处理外部请求，从而导致更长的延迟和响应时间。

### Scenario 2

### 场景 2

On the contrary, memory cannot be compressed and Pods cannot share  memory resources. This means the allocation of new memory resources will definitely fail if memory is running out.

相反，内存不能被压缩，Pods 不能共享内存资源。这意味着如果内存耗尽，新内存资源的分配肯定会失败。

Some processes in a Pod need a certain amount of memory exclusively  in initialization. For example, a JVM applies for a certain amount of  memory upon start-up. If the specified memory request is less than the  memory applied by the JVM, the memory application will fail (OOM-kill). As a result, the Pod will keep restarting and failing.

Pod 中的某些进程在初始化时需要一定数量的内存。例如，JVM 在启动时会申请一定数量的内存。如果指定的内存请求小于JVM申请的内存，内存申请会失败（OOM-kill）。因此，Pod 将不断重新启动和失败。

### Scenario 3

### 场景 3

When a Pod is being created, Kubernetes needs to allocate or  provision different resources including CPU and memory in a balanced and comprehensive way. Meanwhile, the Kubernetes scheduling algorithm  entails a variety of factors, such as `NodeResourcesLeastAllocated` and Pod affinity. The reason why memory resources are often unevenly  distributed is that for apps, memory is considered scarcer than other  resources.

在创建 Pod 时，Kubernetes 需要以均衡、全面的方式分配或配置不同的资源，包括 CPU 和内存。同时，Kubernetes 调度算法需要考虑多种因素，比如 NodeResourcesLeastAllocated 和 Pod 亲和性。内存资源经常分布不均的原因是，对于应用程序来说，内存被认为比其他资源更稀缺。

Besides, a Kubernetes scheduler works based on the current status of a cluster. In other words, when new Pods are created, the scheduler  selects an optimal node for Pods to run on according to the resource  specification of the cluster at that moment. This is where potential  issues may happen as Kubernetes clusters are highly dynamic. For  example, to maintain a node, you may need to cordon it and all the Pods  running on it will be scheduled to other nodes. The problem is, after  the maintenance, these Pods will not be automatically scheduled back to  the original node. This is because a running Pod cannot be rescheduled  by Kubernetes itself to another node once it is bound to a node at the  beginning.

此外，Kubernetes 调度程序基于集群的当前状态工作。换句话说，当新的 Pod 被创建时，调度器会根据当时集群的资源规格为 Pod 选择一个最优的节点来运行。由于 Kubernetes 集群是高度动态的，因此可能会出现潜在问题。例如，要维护一个节点，您可能需要将其封锁，并且在其上运行的所有 Pod 都将被调度到其他节点。问题是，维护后这些 Pod 不会自动调度回原来的节点。这是因为一个运行中的 Pod 不能被 Kubernetes 自己重新调度到另一个节点上，一旦它一开始就绑定到一个节点上。

## Best Practices

## 最佳实践

We can know from the above analysis that cluster stability has a  direct bearing on your app’s performance. Temporary resource shortage is often the major cause of cluster instability, which cloud mean app  malfunctions or even node failures. Here, we would like to introduce two ways to improve cluster stability.

从上面的分析我们可以知道，集群的稳定性直接关系到你的应用程序的性能。临时资源短缺往往是集群不稳定的主要原因，云意味着应用程序故障甚至节点故障。在这里，我们想介绍两种提高集群稳定性的方法。

First, reserve a certain amount of system resources by [editing the kubelet configuration file](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/). This is especially important when you are dealing with incompressible compute resources, such as memory or disk space.

首先，通过 [编辑kubelet配置文件](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/)预留一定的系统资源。当您处理不可压缩的计算资源（例如内存或磁盘空间)时，这一点尤其重要。

Second, configure appropriate [Quality of Service (QoS) classes](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/) for Pods. Kubernetes uses QoS classes to determine the scheduling and  eviction priority of Pods. Different Pods can be assigned different QoS  classes, including `Guaranteed` (highest priority), `Burstable` and `BestEffort` (lowest priority).

其次，为 Pod 配置适当的 [服务质量 (QoS) 类](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/)。 Kubernetes 使用 QoS 类来确定 Pod 的调度和驱逐优先级。可以为不同的 Pod 分配不同的 QoS 等级，包括“Guaranteed”（最高优先级）、“Burstable”和“BestEffort”（最低优先级)。

- `Guaranteed`. Every container in the Pod, including init containers, must have requests and limits specified for CPU and memory, and they must be equal.
- `Burstable`. At least one container in the Pod has requests specified for CPU or memory.
- `BestEffort`. No container in the Pod has requests and limits specified for CPU and memory.

- `保证`。 Pod 中的每个容器，包括 init 容器，都必须为 CPU 和内存指定请求和限制，并且它们必须相等。
- `突发性`。 Pod 中至少有一个容器具有指定的 CPU 或内存请求。
- `尽力而为`。 Pod 中的任何容器都没有为 CPU 和内存指定的请求和限制。

***Note:\***

*With CPU management policies of Kubelet, you can set CPU affinity for a specific Pod. For more information, see [the Kubernetes documentation](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/).* 

*通过 Kubelet 的 CPU 管理策略，您可以为特定 Pod 设置 CPU 亲和性。有关更多信息，请参阅 [Kubernetes 文档](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/)。*

When resources are running out, your cluster will first kill Pods with a QoS class of `BestEffort`, followed by `Burstable`. In other words, Pods that have the lowest priority are terminated  first. If you have enough resources, you can assign all Pods the class  of `Guaranteed`. This can be considered as a trade-off  between compute resources and performance and stability. You may expect  higher overheads but your cluster can work more efficiently. At the same time, to improve resource utilization, you can assign Pods running  business services the class of `Guaranteed`. For other services, assign them the class of `Burstable` or `BestEffort` according to their priority.

当资源耗尽时，您的集群将首先杀死 QoS 类为“BestEffort”的 Pod，然后是“Burstable”。换句话说，优先级最低的 Pod 首先被终止。如果您有足够的资源，您可以为所有 Pod 分配 `Guaranteed` 类。这可以被认为是计算资源与性能和稳定性之间的权衡。您可能会期望更高的开销，但您的集群可以更有效地工作。同时，为了提高资源利用率，可以为运行业务服务的 Pod 分配 `Guaranteed` 类。对于其他服务，根据它们的优先级为它们分配“Burstable”或“BestEffort”类。

Next, we will use the [KubeSphere container platform](https://kubesphere.io/) as an example to see how to gracefully configure resources for Pods.

接下来，我们将以 [KubeSphere 容器平台](https://kubesphere.io/) 为例，看看如何优雅地为 Pod 配置资源。

...

For containers running key business processes, they need to handle more  traffic than other containers.In reality, there is no panacea and you  need to make careful and comprehensive decisions on requests and limits  of these containers. Think about the following questions:

实际上，没有灵丹妙药，您需要对这些容器的请求和限制做出谨慎而全面的决定。思考以下问题：

1. Are your containers CPU-intensive or IO-intensive?
2. Are they highly available?
3. What are the upstream and downstream objects of your service?



1. 你的容器是 CPU 密集型还是 IO 密集型？
2. 它们是否高度可用？
3. 你们服务的上下游对象是什么？

If you look at a load of containers over a long period of time, you  may find it periodic. In this connection, the historical monitoring data can serve as an important reference as you configure requests and  limits. On the back of Prometheus, which is integrated into the  platform, KubeSphere features a powerful and holistic observability  system that monitors resources at a granular level. Vertically, it  covers data from clusters to Pods. Horizontally, it tracks information  about CPU, memory, network, and storage. Generally, you can specify  requests based on the average value of historical data while limits need to be higher than the average. That said, you may need to make some  adjustments to your final decision as needed.

如果您长时间查看大量容器，您可能会发现它是周期性的。在这方面，历史监控数据可以作为您配置请求和限制时的重要参考。在集成到平台中的 Prometheus 的背面，KubeSphere 具有强大的整体可观察性系统，可以在粒度级别监控资源。纵向上，它涵盖了从集群到 Pod 的数据。它横向跟踪有关 CPU、内存、网络和存储的信息。一般可以根据历史数据的平均值来指定请求，而limit需要高于平均值。也就是说，您可能需要根据需要对最终决定进行一些调整。

## Source Code Analysis

## 源码分析

Now that you know some best practices for configuring requests and limits, let’s dive deeper into the source code.

现在您已经了解了一些配置请求和限制的最佳实践，让我们更深入地研究源代码。

### Requests and Scheduling

### 请求和调度

The following code shows the relation between the requests of a Pod and the requests of containers in the Pod.

下面的代码展示了 Pod 的请求和 Pod 中容器的请求之间的关系。

```go
func computePodResourceRequest(pod *v1.Pod) *preFilterState {
    result := &preFilterState{}
    for _, container := range pod.Spec.Containers {
        result.Add(container.Resources.Requests)
    }

    // take max_resource(sum_pod, any_init_container)
    for _, container := range pod.Spec.InitContainers {
        result.SetMaxResource(container.Resources.Requests)
    }

    // If Overhead is being utilized, add to the total requests for the pod
    if pod.Spec.Overhead != nil && utilfeature.DefaultFeatureGate.Enabled(features.PodOverhead) {
        result.Add(pod.Spec.Overhead)
    }

    return result
}
...
func (f *Fit) PreFilter(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod) *framework.Status {
    cycleState.Write(preFilterStateKey, computePodResourceRequest(pod))
    return nil
}
...
func getPreFilterState(cycleState *framework.CycleState) (*preFilterState, error) {
    c, err := cycleState.Read(preFilterStateKey)
    if err != nil {
        // preFilterState doesn't exist, likely PreFilter wasn't invoked.
        return nil, fmt.Errorf("error reading %q from cycleState: %v", preFilterStateKey, err)
    }

    s, ok := c.(*preFilterState)
    if !ok {
        return nil, fmt.Errorf("%+v  convert to NodeResourcesFit.preFilterState error", c)
    }
    return s, nil
}
...
func (f *Fit) Filter(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
    s, err := getPreFilterState(cycleState)
    if err != nil {
        return framework.NewStatus(framework.Error, err.Error())
    }

    insufficientResources := fitsRequest(s, nodeInfo, f.ignoredResources, f.ignoredResourceGroups)

    if len(insufficientResources) != 0 {
        // We will keep all failure reasons.
        failureReasons := make([]string, 0, len(insufficientResources))
        for _, r := range insufficientResources {
            failureReasons = append(failureReasons, r.Reason)
        }
        return framework.NewStatus(framework.Unschedulable, failureReasons...)
    }
    return nil
}
```

It can be seen from the code above that the scheduler (schedule  thread) calculates the resources required by the Pod to be scheduled. Specifically, it calculates the total requests of init containers and  the total requests of working containers respectively according to Pod  specifications. The greater one will be used. Note that for lightweight  virtual machines (e.g., kata-container), their own resource consumption  of virtualization needs to be counted in caches. In the following `Filter` stage, all nodes will be checked to see if they meet the conditions.

从上面的代码可以看出，调度器（调度线程）计算了待调度的Pod所需的资源。具体来说，它根据 Pod 规范分别计算 init 容器的总请求数和工作容器的总请求数。较大的将被使用。请注意，对于轻量级虚拟机（例如 kata-container），它们自己的虚拟化资源消耗需要计入缓存中。在接下来的 `Filter` 阶段，将检查所有节点是否满足条件。

***Note\***

***笔记\***

*The scheduling process entails different stages, including `Pre filter`, `Filter`, `Post filter` and `Score`. For more information, see [filter and score nodes](https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/#kube-scheduler-implementation).*

*调度过程需要不同的阶段，包括`Pre filter`，`Filter`，`Post filter`和`Score`。有关详细信息，请参阅 [过滤器和评分节点](https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/#kube-scheduler-implementation)。*

After the filtering, if there is only one applicable node, the Pod  will be scheduled to it. If there are multiple applicable Pods, the  scheduler will select the node with the highest weighted scores sum. Scoring is based on a variety of factors as [scheduling plugins](https://kubernetes.io/docs/reference/scheduling/config/#scheduling-plugins) implement one or more extension points. Note that the value of `requests` and the value of `limits` impact directly on the ultimate result of the plugin `NodeResourcesLeastAllocated`. Here is the source code:

过滤后，如果只有一个适用的节点，则将 Pod 调度到该节点上。如果有多个适用的 Pod，调度器会选择权重总和最高的节点。评分基于多种因素，因为 [调度插件](https://kubernetes.io/docs/reference/scheduling/config/#scheduling-plugins) 实现了一个或多个扩展点。请注意，`requests` 的值和 `limits` 的值直接影响插件 `NodeResourcesLeastAllocated` 的最终结果。这是源代码：

```go
func leastResourceScorer(resToWeightMap resourceToWeightMap) func(resourceToValueMap, resourceToValueMap, bool, int, int) int64 {
    return func(requested, allocable resourceToValueMap, includeVolumes bool, requestedVolumes int, allocatableVolumes int) int64 {
        var nodeScore, weightSum int64
        for resource, weight := range resToWeightMap {
            resourceScore := leastRequestedScore(requested[resource], allocable[resource])
            nodeScore += resourceScore * weight
            weightSum += weight
        }
        return nodeScore / weightSum
    }
}
...
func leastRequestedScore(requested, capacity int64) int64 {
    if capacity == 0 {
        return 0
    }
    if requested > capacity {
        return 0
    }

    return ((capacity - requested) * int64(framework.MaxNodeScore)) / capacity
}
```

For `NodeResourcesLeastAllocated`, a node will get higher  scores if it has more resources for the same Pod. In other words, a Pod  will be more likely to be scheduled to the node with sufficient  resources.

对于 `NodeResourcesLeastAllocated`，如果一个节点为同一个 Pod 拥有更多资源，它将获得更高的分数。换句话说，一个 Pod 将更有可能被调度到资源充足的节点上。

When a Pod is being created, Kubernetes needs to allocate different  resources including CPU and memory. Each kind of resource has a weight  (the `resToWeightMap` structure in the source code). As a  whole, they tell the Kubernetes scheduler what the best decision may be  to achieve resource balance. In the `Score` stage, the scheduler also uses other plugins for scoring in addition to `NodeResourcesLeastAllocated`, such as `InterPodAffinity`.

在创建 Pod 时，Kubernetes 需要分配不同的资源，包括 CPU 和内存。每种资源都有一个权重（源代码中的`resToWeightMap`结构）。作为一个整体，它们告诉 Kubernetes 调度程序什么是实现资源平衡的最佳决策。在 `Score` 阶段，调度器除了 `NodeResourcesLeastAllocated` 外，还使用其他插件进行评分，例如 `InterPodAffinity`。

### QoS and Scheduling

### QoS 和调度

As a resource protection mechanism in Kubernetes, QoS is mainly used  to control incompressible resources such as memory. It also impacts the  OOM score of different Pods and containers. When a node is running out  of memory, the kernel (OOM Killer) kills Pods of low priority (higher  scores mean lower priority). Here is the source code:

QoS作为Kubernetes中的一种资源保护机制，主要用于控制内存等不可压缩资源。它还会影响不同 Pod 和容器的 OOM 分数。当一个节点内存不足时，内核（OOM Killer）会杀死低优先级的 Pod（分数越高，优先级越低）。这是源代码：

```go
func GetContainerOOMScoreAdjust(pod *v1.Pod, container *v1.Container, memoryCapacity int64) int {
    if types.IsCriticalPod(pod) {
        // Critical pods should be the last to get killed.
        return guaranteedOOMScoreAdj
    }

    switch v1qos.GetPodQOS(pod) {
    case v1.PodQOSGuaranteed:
        // Guaranteed containers should be the last to get killed.
        return guaranteedOOMScoreAdj
    case v1.PodQOSBestEffort:
        return besteffortOOMScoreAdj
    }

    // Burstable containers are a middle tier, between Guaranteed and Best-Effort.Ideally,
    // we want to protect Burstable containers that consume less memory than requested.
    // The formula below is a heuristic.A container requesting for 10% of a system's
    // memory will have an OOM score adjust of 900. If a process in container Y
    // uses over 10% of memory, its OOM score will be 1000. The idea is that containers
    // which use more than their request will have an OOM score of 1000 and will be prime
    // targets for OOM kills.
    // Note that this is a heuristic, it won't work if a container has many small processes.
    memoryRequest := container.Resources.Requests.Memory().Value()
    oomScoreAdjust := 1000 - (1000*memoryRequest)/memoryCapacity
    // A guaranteed pod using 100% of memory can have an OOM score of 10. Ensure
    // that burstable pods have a higher OOM score adjustment.
    if int(oomScoreAdjust) < (1000 + guaranteedOOMScoreAdj) {
        return (1000 + guaranteedOOMScoreAdj)
    }
    // Give burstable pods a higher chance of survival over besteffort pods.
    if int(oomScoreAdjust) == besteffortOOMScoreAdj {
        return int(oomScoreAdjust - 1)
    }
    return int(oomScoreAdjust)
}
```

## Summary

##  概括

As a portable and extensible open-source platform, Kubernetes is born for managing containerized workloads and Services. It boasts a  comprehensive, fast-growing ecosystem that has helped it secure its  position as the de facto standard in container orchestration. That being said, it is not always easy for users to learn Kubernetes and this is  where KubeSphere comes to play its part. KubeSphere empowers users to  perform virtually all the operations on its dashboard while they also  have the option to use the built-in web kubectl tool to run commands. This article focuses on requests and limits, their underlying logic in  Kubernetes, as well as how to use KubeSphere to configure them for easy  operation and maintenance of your cluster.

作为一个可移植和可扩展的开源平台，Kubernetes 为管理容器化工作负载和服务而生。它拥有一个全面的、快速发展的生态系统，帮助它巩固了其作为容器编排事实上的标准的地位。话虽如此，用户学习 Kubernetes 并不总是那么容易，而这正是 KubeSphere 发挥作用的地方。 KubeSphere 使用户能够在其仪表板上执行几乎所有操作，同时他们还可以选择使用内置的 Web kubectl 工具来运行命令。本文重点介绍请求和限制，它们在 Kubernetes 中的底层逻辑，以及如何使用 KubeSphere 配置它们，以便于您的集群的运维。

https://dzone.com/articles/dive-deep-into-resource-requests-and-limits-in-kub

