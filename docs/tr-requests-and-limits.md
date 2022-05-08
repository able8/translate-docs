# [A Practical Guide to Setting Kubernetes Requests and Limits](https://blog.kubecost.com/blog/requests-and-limits/)

# [设置Kubernetes请求和限制的实用指南](https://blog.kubecost.com/blog/requests-and-limits/)

7 minute read

7分钟阅读

Setting Kubernetes requests and limits effectively has a major impact on application performance, stability, and cost. And yet working with many teams over the past year has shown us that determining the right values for these parameters is hard. For this reason, we have created this short guide and are launching a new product to help teams more accurately set Kubernetes requests and limits for their applications.

有效地设置 Kubernetes 请求和限制对应用程序的性能、稳定性和成本有重大影响。
然而，在过去一年中与许多团队合作向我们表明，确定这些参数的正确值是困难的。
出于这个原因，我们创建了这个简短的指南，并正在推出一款新产品，以帮助团队更准确地为其应用程序设置 Kubernetes 请求和限制。

## The Basics

##  基础

Resource requests and limits are optional parameters specified at the container level. Kubernetes computes a Pod’s request and limit as the sum of requests and limits across all of its containers. Kubernetes then uses these parameters for scheduling and resource allocation decisions.

资源请求和限制是在容器级别指定的可选参数。 Kubernetes 将 Pod 的请求和限制计算为其所有容器的请求和限制的总和。然后 Kubernetes 使用这些参数进行调度和资源分配决策。

![resource recs](http://blog.kubecost.com/assets/images/k8s-recs-ands-limits.png)

**Requests**

**要求**

Pods will get the amount of _memory_ they request. If they exceed their memory request, they could be killed if another pod happens to need this memory. Pods are only ever killed when using less memory than requested if critical system or high priority workloads need the memory.

Pod 将获得它们请求的 _memory_ 数量。如果他们超出了他们的内存请求，如果另一个 pod 碰巧需要这个内存，他们可能会被杀死。只有当关键系统或高优先级工作负载需要内存时，Pod 才会在使用的内存少于请求的内存时被杀死。

Similarly, each container in a Pod is allocated the amount of _CPU_ it requests, if available. It may be allocated additional CPU cycles if available resources are not needed by other running Pods/Jobs.

类似地，Pod 中的每个容器都分配了它请求的 _CPU_ 数量（如果可用）。如果其他正在运行的 Pod/Jobs 不需要可用资源，它可能会被分配额外的 CPU 周期。

Note: if a Pod’s total requests are not available on a single node, then the Pod will remain in a Pending state (i.e. not running) until these resources become available.

注意：如果 Pod 的总请求在单个节点上不可用，则 Pod 将保持在 Pending 状态（即未运行），直到这些资源可用。

**Limits**

**限制**

Resource limits help the Kubernetes scheduler better handle resource contention. When a Pod uses more memory than its limit, its processes will be killed by the kernel to protect other applications in the cluster. Pods will be CPU throttled when they exceed their CPU limit.

资源限制有助于 Kubernetes 调度程序更好地处理资源争用。当 Pod 使用的内存超过其限制时，其进程将被内核杀死以保护集群中的其他应用程序。当 Pod 超过其 CPU 限制时，它们将受到 CPU 限制。

If no limit is set, then the pods can use excess memory and CPU when available.

如果未设置限制，则 Pod 可以在可用时使用多余的内存和 CPU。

Here’s a quick example of how you set requests and limits on a container spec:

下面是一个快速示例，说明如何在容器规范上设置请求和限制：

```
apiVersion: v1
kind: Pod
metadata:
name: hello-app
spec:
containers:
  - name: wp
    image: wordpress
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"

```

CPU requests are set in cpu units where 1000 millicpu (“m”) equals 1 vCPU or 1 Core. So 250m CPU equals ¼ of a CPU. Memory can be set with Ti, Gi, Mi, or Ki units. For more advanced technical info on the mechanics of these parameters, we suggest these articles:

CPU 请求以 cpu 单位设置，其中 1000 millicpu (“m”) 等于 1 个 vCPU 或 1 个核心。所以 250m CPU 等于 CPU 的 ¼。内存可以设置为 Ti、Gi、Mi 或 Ki 单位。有关这些参数机制的更高级技术信息，我们建议阅读以下文章：

- [https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/resource-qos.md](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/resource-qos.md)
- [https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/)

## The Tradeoffs

## 权衡

Determining the right level for requests and limits is about managing trade-offs, as shown in the following tables.
When setting requests, there is inherently a tradeoff between the cost of running an application and the performance/outage risk for this application.
Balancing these risks depends on the relative cost of extra CPU/RAM compared to the expected cost of an application throttle or outage event.
For example, if allocating another 1 Gb of RAM ($5 cost) reduces the risk of an application outage event ($10,000 cost) by 1% then it would be worth the additional cost of these compute resources.

确定请求和限制的正确级别是关于管理权衡，如下表所示。
在设置请求时，在运行应用程序的成本和该应用程序的性能/中断风险之间存在固有的权衡。
平衡这些风险取决于额外 CPU/RAM 的相对成本与应用程序限制或中断事件的预期成本相比。
例如，如果再分配 1 Gb RAM（成本 5 美元）将应用程序中断事件的风险（成本 10,000 美元）降低 1%，那么这些计算资源的额外成本是值得的。

RequestToo lowToo highCPUStarvation – may not get CPU cycles neededInefficiency – requires extra CPUs to schedule other PodsMemoryKill risk – may be terminated if other pods need memoryInefficiency – requires extra RAM to schedule other Pods 

RequestToo lowToo highCPUStarvation - 可能无法获得所需的 CPU 周期效率低下 - 需要额外的 CPU 来调度其他 PodsMemoryKill 风险 - 如果其他 Pod 需要内存，可能会被终止效率低下 - 需要额外的 RAM 来调度其他 Pod

When setting limits, the tradeoffs are similar but not quite the same. The tradeoff here is the relative performance of individual applications on your shared infrastructure vs the total cost of running these applications. For example, setting the aggregated amount of CPU limits higher than the allocated number of CPUs exposes applications to potential throttling risk. Provisioning additional CPUs (i.e. increase spend) is one potential answer while reducing CPU limits for certain applications (i.e. increase throttling risk) is another.

在设置限制时，权衡是相似的，但并不完全相同。这里的权衡是共享基础架构上单个应用程序的相对性能与运行这些应用程序的总成本。例如，将 CPU 限制的总量设置为高于分配的 CPU 数量会使应用程序面临潜在的限制风险。提供额外的 CPU（即增加支出）是一个潜在的答案，而降低某些应用程序的 CPU 限制（即增加节流风险）是另一个答案。



In the following section, we present a framework for managing these tradeoffs more effectively.

在下一节中，我们将提出一个更有效地管理这些权衡的框架。

## Determining the right values

## 确定正确的值

When setting requests, start by determining the acceptable probability of a container’s usage exceeding its request in a specific time window, e.g. 24 hours. To predict this in future periods, we can analyze historical resource usage. As an example, allowing usage to be above a request threshold with 0.01 probability (i.e. three nines) means that it will, on average, face increased risk of throttling or being killed 1.44 minutes a day.

设置请求时，首先确定容器在特定时间窗口内的使用超过其请求的可接受概率，例如24小时。为了预测未来时期的情况，我们可以分析历史资源使用情况。例如，允许使用率以 0.01 的概率（即三个 9）高于请求阈值意味着它平均每天面临 1.44 分钟的节流或被杀死的风险。

You can classify applications into different availability tiers and apply these rules of thumb for targeting the appropriate level of availability:

您可以将应用程序分类为不同的可用性层，并应用这些经验法则来定位适当的可用性级别：

TierRequestLimitCritical / Highly Available99.99th percentile + 100% headroom2x request or as higher if resources availableProduction / Non-critical99th + 50% headroom2x requestDev / Experimental95th or consider namespace quotas\*1.5x request or consider namespace quotas\*

TierRequestLimitCritical /高度可用99.99% + 100% headroom2x 请求或更高，如果资源可用生产/非关键99th + 50% headroom2x requestDev / Experimental95th 或考虑命名空间配额\*1.5x 请求或考虑命名空间配额\*

This approach of analyzing historical usage patterns typically provides both a good representation of the future and is easy to understand/introspect. Applying extra headroom allows for fluctuations that may have been missed by your historical sampling. We recommend measuring usage over 1 week at a minimum and setting thresholds based on the specific availability requirements of your pod.

这种分析历史使用模式的方法通常既能很好地代表未来，又易于理解/反思。应用额外的净空允许您的历史采样可能遗漏的波动。我们建议至少测量 1 周以上的使用情况，并根据您的 pod 的特定可用性要求设置阈值。

_\*Developers sharing experimental clusters may rely on broader protection from [resource quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/).
Quotas set aggregate caps at the namespace level and can help protect tasks like long-running batch or ML jobs from getting killed because someone improperly specified resources in another namespace._

_\*共享实验集群的开发人员可能依赖更广泛的保护以免受 [资源配额](https://kubernetes.io/docs/concepts/policy/resource-quotas/)。
配额在命名空间级别设置聚合上限，有助于保护长期运行的批处理或 ML 作业等任务不会因为有人在另一个命名空间中不正确地指定资源而被杀死。_

## Our Solution

## 我们的解决方案

Seeing the difficulty of setting these parameters correctly and managing them over time motivated us to create a solution in the Kubecost product to directly generate recommendations for your applications. Our recommendations are based on a configurable Availability Tiers (e.g. Production or Dev), which is easily tracked by namespace or other concepts directly in the Kubecost product.

看到正确设置这些参数并随着时间的推移对其进行管理的难度促使我们在 Kubecost 产品中创建一个解决方案，以直接为您的应用程序生成建议。我们的建议基于可配置的可用性层（例如生产或开发），可以通过命名空间或其他概念直接在 Kubecost 产品中轻松跟踪。

![resource recs](http://blog.kubecost.com/assets/images/kubecost-request-recs.png)

In addition to providing request recommendations, this solution also proactively detects out of memory and CPU throttle risks. The full Kubecost product is available via a single Helm command ( [install options](http://docs.kubecost.com/install)) and these recommendations can easily be viewed for each container in the Namespace view. Our commercial product is free for small clusters, comes with a free trial for larger clusters, and is based on the [Kubecost open source project](https://github.com/kubecost/cost-model).

除了提供请求建议外，该解决方案还主动检测内存不足和 CPU 节流风险。完整的 Kubecost 产品可通过单个 Helm 命令（[安装选项](http://docs.kubecost.com/install))获得，并且可以在命名空间视图中轻松查看每个容器的这些建议。我们的商业产品对小型集群是免费的，为大型集群提供免费试用，并且基于 [Kubecost 开源项目](https://github.com/kubecost/cost-model)。

Also, here are sample Prometheus queries if you want to calculate these metrics yourself!

此外，如果您想自己计算这些指标，这里有一些 Prometheus 查询示例！

**Memory Request (Production Tier)**

**内存请求（生产层）**

We recommend `container_memory_working_set_bytes` because this metric excludes cached data and is what Kubernetes uses for OOM/scheduling decisions. More info in this [article](https://medium.com/faun/how-much-is-too-much-the-linux-oomkiller-and-used-memory-d32186f29c9d).

我们推荐使用 `container_memory_working_set_bytes`，因为该指标不包括缓存数据，并且是 Kubernetes 用于 OOM/调度决策的指标。这篇[文章](https://medium.com/faun/how-much-is-too-much-the-linux-oomkiller-and-used-memory-d32186f29c9d)中的更多信息。

```
1.5 * avg(quantile_over_time(.99,container_memory_working_set_bytes{container_name!="POD",container_name!=""}[7d])) by (container_name,pod_name,namespace)

```

**CPU Request (Production Tier)**

**CPU 请求（生产层）**

First, create a recording rule with this expression. Note we recommend using `irate` to capture short-term spike is resource needs.

首先，用这个表达式创建一个记录规则。请注意，我们建议使用 `irate` 来捕获短期峰值是资源需求。

```
avg( irate(container_cpu_usage_seconds_total{container_name!="POD", container_name!=""}[5m]) ) by (container_name,pod_name,namespace)

```

Then run this query:

然后运行这个查询：

```
1.5 * quantile_over_time(.99,container_cpu_usage_irate[7d])
```

## Vertical pod autoscaling 

## 垂直 Pod 自动缩放

The goal of vertical pod autoscaling (VPA) is to remove the need to worry about specifying values for a container’s CPU and memory requests. It can be a great solution in certain situations, often with stateless workloads, but you should note that this tool is still in beta as of September 2019. Here are some of the practical limitations to be aware of:

垂直 Pod 自动缩放 (VPA) 的目标是无需担心为容器的 CPU 和内存请求指定值。在某些情况下（通常是无状态工作负载），它可能是一个很好的解决方案，但您应该注意，截至 2019 年 9 月，此工具仍处于测试阶段。以下是一些需要注意的实际限制：

- Pods are evicted and need to restart when the VerticalPodAutoscaler needs to change the Pod’s resource requests.
- VPA can cause performance risk and outages if not configured correctly and adds observability complexity.
- To appropriately handle scale up events, it’s recommended that Cluster Autoscaler also be enabled to handle the increased resource requirements sizes of your workloads.
- The VPA requires careful tuning to implement a tier-based solution with different parameters for highly available apps, dev, prod, staging etc.

- 当 VerticalPodAutoscaler 需要更改 Pod 的资源请求时，Pod 被驱逐并需要重新启动。
- 如果配置不正确，VPA 可能会导致性能风险和中断，并增加可观察性的复杂性。
- 为了适当地处理扩展事件，建议还启用 Cluster Autoscaler 以处理工作负载增加的资源需求大小。
- VPA 需要仔细调整，以实现具有不同参数的基于层的解决方案，用于高可用性应用程序、开发、生产、登台等。

We advise teams to be cautious when using VPA for critical production workloads. It introduces complexity to your infrastructure and you should adequately ensure that your deployments and VPA itself are configured correctly. Risks aside, it can be a great solution when applied correctly.

我们建议团队在将 VPA 用于关键生产工作负载时要谨慎。它给您的基础架构带来了复杂性，您应该充分确保您的部署和 VPA 本身配置正确。撇开风险不谈，如果应用得当，它可能是一个很好的解决方案。

More info on VPA is available here:

有关 VPA 的更多信息可在此处获得：

- [https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)
- [https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/vertical-pod-autoscaler.md](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/vertical-pod-autoscaler.md)

## Conclusion

##  结论

Setting requests and limits effectively can make or break application performance and reliability in Kubernetes. This set of guidelines and this new Kubecost tool can help you manage the inherent risks and tradeoffs when solution like vertical pod autoscaling are not the right fit. Our recommendations combined cost data and health insights are available in Kubecost today to help you make informed decisions.

有效地设置请求和限制可以决定或破坏 Kubernetes 中的应用程序性能和可靠性。当垂直 Pod 自动缩放等解决方案不适合时，这套指南和这个新的 Kubecost 工具可以帮助您管理固有风险和权衡。我们的建议结合了成本数据和健康见解，现在可以在 Kubecost 中获得，以帮助您做出明智的决定。

## About Kubecost

## 关于 Kubecost

[Kubecost](http://kubecost.com) provides cost and capacity management tools purpose built for Kubernetes. We help teams do Kubernetes chargeback and more. Reach out via email (team@kubecost), [Slack](https://join.slack.com/t/kubecost/shared_invite/enQtNTA2MjQ1NDUyODE5LWFjYzIzNWE4MDkzMmUyZGU4NjkwMzMyMjIyM2E0NGNmYjExZjBiNjk1YzY5ZDI0ZTNhZDg4NjlkMGRkYzFlZTU) or visit [our website](http://kubecost.com) if you want to talk shop or learn more!

[Kubecost](http://kubecost.com) 提供专为 Kubernetes 构建的成本和容量管理工具。我们帮助团队进行 Kubernetes 退款等。 Reach out via email (team@kubecost), [Slack](https://join.slack.com/t/kubecost/shared_invite/enQtNTA2MjQ1NDUyODE5LWFjYzIzNWE4MDkzMmUyZGU4NjkwMzMyMjIyM2E0NGNmYjExZjBiNjk1YzY5ZDI0ZTNhZDg4NjlkMGRkYzFlZTU) or visit [our website](http://kubecost.com) if you want谈论商店或了解更多信息！

A big thanks to all that have given us feedback and made contributions!

非常感谢所有给我们反馈和贡献的人！

**Tags:** [cost monitoring](http://blog.kubecost.com/tags/#cost-monitoring),[Kubecost](http://blog.kubecost.com/tags/#kubecost)

**标签：**[成本监控](http://blog.kubecost.com/tags/#cost-monitoring),[Kubecost](http://blog.kubecost.com/tags/#kubecost)

**Categories:**[blog](http://blog.kubecost.com/categories/#blog)

**类别：**[博客](http://blog.kubecost.com/categories/#blog)

**Updated:** September 25, 2019

**更新时间：** 2019 年 9 月 25 日

https://blog.kubecost.com/blog/requests-and-limits 

