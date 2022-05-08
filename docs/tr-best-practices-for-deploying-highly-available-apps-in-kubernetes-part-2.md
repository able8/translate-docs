# Best practices for deploying highly available apps in Kubernetes. Part 2

# 在 Kubernetes 中部署高可用应用的最佳实践。第2部分

16 August 2021

In [Part One](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/), we shared some recommendations for a number of Kubernetes mechanisms that facilitate the deployment of highly available applications. We discussed aspects of the scheduler operation, update strategies, priorities, probes, etc. This second part discusses the remaining three crucial topics: PodDisruptionBudget, HorizontalPodAutoscaler, and VerticalPodAutoscaler (the numbering continues from [Part One](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/)).

在 [Part One](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/) 中，我们分享了一些 Kubernetes 机制的建议这有助于部署高可用性应用程序。我们讨论了调度器操作、更新策略、优先级、探测等方面。第二部分讨论了剩下的三个关键主题：PodDisruptionBudget、HorizontalPodAutoscaler 和 VerticalPodAutoscaler。

## 9\. PodDisruptionBudget

The [PodDisruptionBudget](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#pod-disruption-budgets)(PDB) mechanism is a must-have for applications running in production. It provides you the means to specify a maximum limit to the number of application Pods that can be unavailable simultaneously. In [Part One](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/), we discussed some methods that are instrumental in avoiding potentially risky situations: running several application replicas, specifying `podAntiAffinity` (to prevent several Pods from being assigned to the same node), etc.

[PodDisruptionBudget](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#pod-disruption-budgets)(PDB) 机制是在生产中运行的应用程序的必备条件。它为您提供了指定同时不可用的应用程序 Pod 数量的最大限制的方法。在 [Part One](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/) 中，我们讨论了一些有助于避免潜在风险情况：运行多个应用程序副本，指定 `podAntiAffinity`（以防止将多个 Pod 分配到同一个节点)等。

However, you may encounter a situation in which more than one K8s node becomes unavailable at the same time. For example, suppose you decide to switch instances to more powerful ones. There may be other reasons besides that, but that’s beyond the scope of this article. The thing is that several nodes get removed at the same time. “But that’s Kubernetes!” you might say. ”Everything is ephemeral here! The Pods will get moved to other nodes, so what’s the deal?” Well, let’s have a look.

但是，您可能会遇到多个 K8s 节点同时不可用的情况。例如，假设您决定将实例切换到更强大的实例。除此之外可能还有其他原因，但这超出了本文的范围。问题是几个节点同时被删除。 “但那是 Kubernetes！”你可能会说。 “这里的一切都是短暂的！ Pod 将被移动到其他节点，所以有什么关系？”好吧，让我们看看。

Suppose the application has three replicas. The load is evenly distributed between them, while the Pods are distributed across the nodes. In this case, the application will continue to run even if one of the replicas fails. However, the failure of two replicas will result in a service degradation: one single Pod simply cannot handle the entire load on its own. The clients will start getting 5XX errors. (Of course, you can set a rate limit in the nginx container; in that case, the error will be _429 Too Many Requests_. Still, the service will degrade nevertheless).

假设应用程序具有三个副本。负载均匀分布在它们之间，而 Pod 分布在节点之间。在这种情况下，即使其中一个副本发生故障，应用程序也会继续运行。但是，两个副本的故障将导致服务降级：一个 Pod 根本无法独自处理整个负载。客户端将开始收到 5XX 错误。 （当然，您可以在 nginx 容器中设置速率限制；在这种情况下，错误将是_429 Too Many Requests_。不过，服务仍然会降级）。

And that’s where PodDisruptionBudget comes to help. Let’s take a look at its manifest:

这就是 PodDisruptionBudget 提供帮助的地方。让我们看一下它的清单：

```
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
name: app-pdb
spec:
maxUnavailable: 1
selector:
    matchLabels:
      app: app
```

The manifest is pretty straightforward; you are probably familiar with most of its fields, `maxUnavailable` being the most interesting among them. This field sets the maximum number of Pods that can be simultaneously unavailable. This can be either an absolute number or a percentage.

清单非常简单；您可能熟悉它的大部分字段，其中“maxUnavailable”是最有趣的。该字段设置可以同时不可用的 Pod 的最大数量。这可以是绝对数字或百分比。

Suppose the PDB is configured for the application. What will happen in the event that, for some reason, two or more nodes start evicting application Pods? The above PDB only allows for one Pod to be evicted at a time. Therefore, the second node will wait until the number of replicas reverts to the pre-eviction level, and only then will the second replica be evicted.

假设为应用程序配置了 PDB。如果出于某种原因，两个或更多节点开始驱逐应用程序 Pod，会发生什么情况？上面的 PDB 只允许一次驱逐一个 Pod。因此，第二个节点将等到副本数恢复到驱逐前的水平，然后才会驱逐第二个副本。

As an alternative, you can also set a `minAvailable` parameter. For example:

作为替代方案，您还可以设置一个 `minAvailable` 参数。例如：

```
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
name: app-pdb
spec:
minAvailable: 80%
selector:
    matchLabels:
      app: app
```

This parameter ensures that at least 80% of replicas are available in the cluster at all times. Thus, only 20% of replicas can be evicted if necessary. `minAvailable` can be either an absolute number or a percentage.

此参数确保集群中至少 80% 的副本始终可用。因此，如有必要，只能驱逐 20% 的副本。 `minAvailable` 可以是绝对数或百分比。

But there is a catch: there have to be enough nodes in the cluster that satisfy the `podAntiAffinity` criteria. Otherwise, you may encounter a situation in which a replica gets evicted, but the scheduler cannot re-deploy it due to a lack of suitable nodes. As a result, draining a node will take forever to complete, and will get you two application replicas instead of three. Granted, you can invoke `kubectl describe` for a _Pending_ Pod to see what’s going on and eliminate the problem. But still, it is better to prevent this kind of situation from happening.

但是有一个问题：集群中必须有足够的节点满足“podAntiAffinity”标准。否则，您可能会遇到副本被驱逐的情况，但由于缺少合适的节点，调度程序无法重新部署它。因此，排空一个节点需要很长时间才能完成，并且会得到两个应用程序副本，而不是三个。当然，您可以为 _Pending_ Pod 调用 `kubectl describe` 来查看发生了什么并消除问题。但是，最好还是防止这种情况发生。

To summarize, **always configure the PDB for critical components of your system**.

总而言之，**始终为系统的关键组件配置 PDB**。

## 10\. HorizontalPodAutoscaler



Let’s consider another situation: what happens if an application has an unexpected load that is significantly higher than usual? Yes, you can scale up the cluster manually, but that is not the method we use.

让我们考虑另一种情况：如果应用程序的意外负载明显高于平时，会发生什么情况？是的，您可以手动扩展集群，但这不是我们使用的方法。

That is where [HorizontalPodAutoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)(HPA) comes in. With HPA, you can choose a metric and use it as a trigger for scaling the cluster up/down automatically, depending on the metric's value. Imagine that on one quiet night your cluster suddenly gets blasted with a massive uptick in traffic, say, Reddit users have found out about your service. The CPU load (or some other Pod metric) increases, hits the threshold, and then HPA comes into play. It scales up the cluster, thus distributing the load between a larger number of Pods.

这就是 [HorizontalPodAutoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)(HPA) 的用武之地。使用 HPA，您可以选择一个指标并将其用作触发器用于根据指标的值自动向上/向下扩展集群。想象一下，在一个安静的夜晚，您的集群突然因流量大幅增加而爆炸，例如，Reddit 用户发现了您的服务。 CPU 负载（或其他一些 Pod 指标)增加，达到阈值，然后 HPA 开始发挥作用。它扩展了集群，从而在大量 Pod 之间分配负载。

Thanks to that, all the incoming requests are processed successfully. Just as important, after the load returns to the average level, **HPA scales the cluster down to reduce infrastructure costs and save money**. Sounds great, doesn’t it?

多亏了这一点，所有传入的请求都被成功处理。同样重要的是，在负载恢复到平均水平后，**HPA 会缩小集群以降低基础架构成本并节省资金**。听起来很棒，不是吗？

Let’s see how exactly HPA calculates the number of replicas to be added. Here is the formula from the documentation:

让我们看看 HPA 如何准确计算要添加的副本数。这是文档中的公式：

`desiredReplicas = ceil[currentReplicas * ( currentMetricValue / desiredMetricValue )]` 

Now suppose that:

现在假设：

- the current number of replicas is 3;
- the current metric value is 100;
- the metric threshold value is 60;

- 当前副本数为 3；
- 当前指标值为 100；
- 度量阈值为 60；

In this case, the resulting number is `3 * ( 100 / 60 )`, i.e. “about” 5 replicas (HPA will round the result up). Thus, the application will gain two more replicas. But that is not the end of the story: HPA will continue to calculate the number of replicas required (using the formula above) to scale down the cluster if the load decreases.

在这种情况下，结果数字是 `3 * ( 100 / 60 )`，即“大约” 5 个副本（HPA 会将结果向上取整）。因此，应用程序将获得另外两个副本。但这并不是故事的结局：如果负载减少，HPA 将继续计算所需的副本数量（使用上面的公式）以缩减集群。

And that brings us to the most exciting part. What metric should you use? The first thing that comes to mind is one of the primary metrics, such as CPU or Memory utilization. And that will work if your CPU and Memory consumption is directly proportional to the incoming load. But what if the Pods are handling different requests? Some requests require many CPU cycles, others may consume a lot of memory, and still others only demand minimum resources.

这将我们带到了最激动人心的部分。您应该使用什么指标？首先想到的是主要指标之一，例如 CPU 或内存利用率。如果您的 CPU 和内存消耗与传入负载成正比，这将起作用。但是如果 Pod 处理不同的请求呢？有些请求需要很多 CPU 周期，有些可能会消耗大量内存，还有一些只需要最少的资源。

Let’s take a look, for example, at the RabbitMQ queue and the instances processing it. Suppose there are ten messages in the queue. Monitoring shows that messages are being dequeued (as per RabbitMQ’s terminology) steadily and regularly. That is, we feel that ten messages in the queue on average is okay. But then the load suddenly increases, and the queue grows to 100 messages. However, the workers’ CPU and Memory consumption stays the same: they are steadily processing the queue, leaving about 80-90 messages in it.

例如，让我们看一下 RabbitMQ 队列和处理它的实例。假设队列中有十条消息。监控显示消息正在稳定且定期地出队（根据 RabbitMQ 的术语）。也就是说，我们觉得队列中平均有 10 条消息是可以的。但随后负载突然增加，队列增长到 100 条消息。但是，worker 的 CPU 和内存消耗保持不变：他们正在稳定地处理队列，在其中留下大约 80-90 条消息。

But **what if we use a custom metric** that describes the number of messages in the queue? Let’s configure our custom metric as follows:

但是**如果我们使用描述队列中消息数量的自定义指标**会怎样？让我们按如下方式配置我们的自定义指标：

- the current number of replicas is 3;
- the current metric value is 80;
- the metric threshold value is 15. 

- 当前副本数为 3；
- 当前指标值为 80；
- 度量阈值为 15。

Thus, `3 * ( 80 / 15 ) = 16`. In this case, HPA can increase the number of workers to 16, and they quickly process all the messages in the queue (at which point HPA will decrease their number again). However, all the required infrastructure must be ready to accommodate this number of Pods. That is, they must fit on the existing nodes, or new nodes must be provisioned by the infrastructure provider (cloud provider) in the case that [Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) is used. In other words, we are back to planning cluster resources.

因此，`3 * ( 80 / 15 ) = 16`。在这种情况下，HPA 可以将工作人员的数量增加到 16 个，并且他们可以快速处理队列中的所有消息（此时 HPA 将再次减少他们的数量）。但是，所有必需的基础设施都必须准备好容纳这个数量的 Pod。也就是说，它们必须适合现有节点，或者在 [Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/)的情况下必须由基础设施提供商（云提供商）提供新节点使用集群自动缩放器)。换句话说，我们又回到了规划集群资源上。

Now let’s take a look at some manifests:

现在让我们看一些清单：

```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
name: php-apache
spec:
scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: php-apache
minReplicas: 1
maxReplicas: 10
targetCPUUtilizationPercentage: 50
```

This one is simple. As soon as the CPU load reaches 50%, HPA starts scaling the number of replicas to a maximum of 10.

这个很简单。一旦 CPU 负载达到 50%，HPA 就会开始将副本数量扩展到最多 10 个。

Here is a more interesting one:

这里有一个更有趣的：

```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
name: worker
spec:
scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: worker
minReplicas: 1
maxReplicas: 10
metrics:
  - type: External
    external:
      metric:
        name: queue_messages
      target:
        type: AverageValue
        averageValue: 15
```

Note that in this example, HPA uses the [custom metric](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-custom-metrics). It will base its scaling decisions on the size of the queue ( `queue_messages` metric). Given that the average number of messages in the queue is 10, we set the threshold to 15. This way, you can manage the number of replicas more accurately. As you can see, the custom metric enables more accurate cluster autoscaling than, say, a CPU-based metric.

请注意，在此示例中，HPA 使用 [自定义指标](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-custom-metrics)。它将根据队列的大小（`queue_messages`指标)做出扩展决策。鉴于队列中的平均消息数为 10，我们将阈值设置为 15。这样，您可以更准确地管理副本数。如您所见，与基于 CPU 的指标相比，自定义指标支持更准确的集群自动缩放。

### Additional features

###  附加的功能

The HPA configuration options are pretty diverse. For example, you can combine different metrics. In the manifest below, CPU utilization and queue size are used to trigger scaling decisions.

HPA 配置选项非常多样化。例如，您可以组合不同的指标。在下面的清单中，CPU 利用率和队列大小用于触发扩展决策。

```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
name: worker
spec:
scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: worker
minReplicas: 1
maxReplicas: 10
metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 50
  - type: External
    external:
      metric:
        name: queue_messages
      target:
        type: AverageValue
        averageValue: 15
```

What calculation algorithm does HPA apply? Well, it uses the highest calculated number of replicas regardless of the metric exploited. For example, if the calculation based on the CPU metric shows that 5 replicas need to be added while the queue size-based metric gives only 3 Pods, HPA will use the larger value and add 5 Pods.

HPA 应用什么计算算法？好吧，它使用计算出的最高副本数，而不管利用的指标如何。例如，如果基于 CPU 指标的计算显示需要添加 5 个副本，而基于队列大小的指标仅给出 3 个 Pod，则 HPA 将使用较大的值并添加 5 个 Pod。

With the release of Kubernetes 1.18, you now [have the ability](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-configurable-scaling-behavior) to define `scaleUp` and `scaleDown` policies. For example:

随着 Kubernetes 1.18 的发布，您现在 [有能力](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-configurable-scaling-behavior) 定义`scaleUp` 和 `scaleDown` 政策。例如：

```
behavior:
scaleDown:
    stabilizationWindowSeconds: 60
    policies:
    - type: Percent
      value: 5
      periodSeconds: 20
   - type: Pods
      value: 5
      periodSeconds: 60
    selectPolicy: Min
scaleUp:
    stabilizationWindowSeconds: 0
    policies:
    - type: Percent
      value: 100
      periodSeconds: 10
```

As you can see in the manifest above, it features two sections. The first one ( `scaleDown`) defines the scaling down parameters while the second ( `scaleUp`) is used for scaling up. Each section features the `stabilizationWindowSeconds`. This helps prevent what is referred to as “flapping” (or unnecessary scaling) as the number of replicas continues to oscillate. This parameter essentially serves as a timeout after the number of replicas is changed. 

正如您在上面的清单中所见，它具有两个部分。第一个（`scaleDown`）定义缩小参数，而第二个（`scaleUp`）用于放大。每个部分都有“stabilizationWindowSeconds”。这有助于防止所谓的“抖动”（或不必要的缩放），因为副本的数量继续波动。这个参数本质上是在副本数改变后的超时。

Now let’s talk about the policies. The `scaleDown` policy allows you to specify the percent of Pods ( `type: Percent`) to scale down over a specific period of time. If the load features a cyclical pattern, what you have to do is decrease the percentage and increase the duration period. In that case, as the load decreases, HPA will not kill a large number of Pods at once (according to its formula) but will do so gradually instead. Furthermore, you can set the maximum number of Pods ( `type: Pods`) that HPA is allowed to kill over the specified time period.

现在让我们谈谈政策。 `scaleDown` 策略允许您指定要在特定时间段内缩减的 Pod 百分比（`type: Percent`）。如果负载具有周期性模式，您需要做的是降低百分比并增加持续时间。在这种情况下，随着负载的减少，HPA 不会一次杀死大量的 Pod（根据它的公式），而是逐渐地这样做。此外，您可以设置 HPA 在指定时间段内允许杀死的最大 Pod 数（`type: Pods`）。

Note the `selectPolicy: Min` parameter. What that means is HPA uses the policy that affects the minimum number of Pods. Thus, HPA will choose a percent value if it (5% in the example above) is less than the numeric alternative (5 Pods in the example above). Conversely, the `selectPolicy: Max` policy will have the opposite effect.

注意 `selectPolicy: Min` 参数。这意味着 HPA 使用影响最小 Pod 数量的策略。因此，如果它（上例中为 5%）小于数字替代值（上例中为 5 个 Pod），HPA 将选择一个百分比值。相反，`selectPolicy: Max` 策略会产生相反的效果。

Similar parameters are used in the `scaleUp` section. Note that in most situations, the cluster must scale up (almost) instantly since even a slight delay can affect users and their experience. For that reason, `stabilizationWindowSeconds` is set to 0 in this section. If the load has a cyclical pattern, HPA can increase the replica count to `maxReplicas` (as defined in the HPA manifest) if necessary. Our policy allows HPA to add up to 100% to the currently running replicas every 10 seconds ( `periodSeconds: 10`).

`scaleUp` 部分使用了类似的参数。请注意，在大多数情况下，集群必须（几乎）立即扩展，因为即使是轻微的延迟也会影响用户及其体验。因此，在本节中，`stabilizationWindowSeconds` 设置为 0。如果负载具有循环模式，HPA 可以在必要时将副本计数增加到“maxReplicas”（如 HPA 清单中所定义）。我们的策略允许 HPA 每 10 秒（`periodSeconds: 10`）向当前运行的副本添加最多 100% 的数据。

Finally, you can set the `selectPolicy` parameter to `Disabled` to turn off scaling in the given direction:

最后，您可以将 `selectPolicy` 参数设置为 `Disabled` 以关闭给定方向的缩放：

```
behavior:
scaleDown:
    selectPolicy: Disabled
```

Most of the time that policies are used is when HPA does not work as expected. **Policies provide flexibility but render the manifest harder to grasp.**

大多数情况下，使用策略是在 HPA 无法按预期工作时。 **政策提供了灵活性，但使清单更难掌握。**

Recently, HPA [became capable](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#container-resource-metrics) to track the resource usage of individual containers across a set of Pods (introduced as an alpha feature in Kubernetes 1.20).

最近，HPA [变得有能力](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#container-resource-metrics) 可以跨一组 Pod 跟踪单个容器的资源使用情况（作为 Kubernetes 1.20 中的 alpha 功能引入)。

### HPA: summary

### HPA：总结

Let us conclude this section with an example of the complete HPA manifest:

让我们以一个完整的 HPA 清单示例结束本节：

```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
name: worker
spec:
scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: worker
minReplicas: 1
maxReplicas: 10
metrics:
  - type: External
    external:
      metric:
        name: queue_messages
      target:
        type: AverageValue
        averageValue: 15
behavior:
    scaleDown:
      stabilizationWindowSeconds: 60
      policies:
      - type: Percent
        value: 5
        periodSeconds: 20
     - type: Pods
        value: 5
        periodSeconds: 60
      selectPolicy: Min
    scaleUp:
      stabilizationWindowSeconds: 0
      policies:
      - type: Percent
        value: 100
        periodSeconds: 10
```

Please note this example is provided for informational purposes only. You will need to adapt it to suit the specifics of your own operation.

请注意，此示例仅供参考。您需要对其进行调整以适应您自己操作的具体情况。

Horizontal Pod Autoscaler resume: **HPA is perfect for production environments. But you have to be careful and forward-thinking when choosing metrics for HPA.** A mistaken metric or an incorrect threshold will result in either a waste of resources (from unnecessary replicas) or service degradation (if the number of replicas is not enough ). Closely monitor the behavior of the application and test it until you’ve achieved the right balance.

Horizontal Pod Autoscaler resume：**HPA 非常适合生产环境。但是在为 HPA 选择指标时，您必须谨慎并具有前瞻性。** 错误的指标或不正确的阈值将导致资源浪费（来自不必要的副本）或服务降级（如果副本数量不足） ）。密切监视应用程序的行为并对其进行测试，直到达到适当的平衡。

## 11\. VerticalPodAutoscaler

## 11.  VerticalPodAutoscaler

[VPA](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) analyzes the resource requirements of the containers and sets (if the corresponding mode is enabled) their limits and requests.

[VPA](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)分析容器的资源需求并设置（如果启用了相应的模式)它们的限制和请求。

Suppose you have deployed a new app version with some new functions and it turns out that, say, the imported library is a huge resource eater, or the code isn’t very well optimized. In other words, the application resource requirements have increased. You failed to notice this during testing (since it is hard to load the application in the same way as in production). 

假设您已经部署了一个具有一些新功能的新应用程序版本，结果发现，例如，导入的库是一个巨大的资源消耗者，或者代码没有得到很好的优化。换句话说，应用程序资源需求增加了。您在测试期间没有注意到这一点（因为很难以与生产中相同的方式加载应用程序）。

And, of course, the relevant requests and limits had been set for the app before an update begins. And now the application reaches the memory limit, and its Pod gets killed due to OOM. VPA can prevent this! At first glance, **VPA looks like a great tool that should be used whenever and wherever possible. But in real life that isn’t always necessarily the case**, and you have to bear in mind the finer details involved.

当然，在更新开始之前，已经为应用程序设置了相关请求和限制。现在应用程序达到了内存限制，并且它的 Pod 因 OOM 而被杀死。 VPA可以防止这种情况！乍一看，**VPA 看起来是一个很棒的工具，应该随时随地使用。但在现实生活中，情况并非总是如此**，您必须牢记所涉及的更精细的细节。

The main problem (it isn’t solved yet) is that the Pod needs to be restarted for resource changes to take effect. In the future, VPA will modify them without restarting the Pod, but for now, it simply isn’t capable of doing that. But no need to worry. That isn’t a big deal if you have a “well-written” application that is always ready for redeployment (say, it has a large number of replicas; its PodAntiAffinity, PodDistruptionBudget, HorizontalPodAutoscaler are carefully configured; etc.). In that case, you (probably) won’t even notice the VPA activity.

主要问题（尚未解决）是需要重新启动 Pod 才能使资源更改生效。将来，VPA 将在不重新启动 Pod 的情况下修改它们，但目前，它根本无法做到这一点。但无需担心。如果您有一个“编写良好”的应用程序，并且始终准备好重新部署（例如，它有大量副本；它的 PodAntiAffinity、PodDistruptionBudget、HorizontalPodAutoscaler 都经过精心配置；等等），这没什么大不了的。在这种情况下，您（可能）甚至不会注意到 VPA 活动。

Sadly, there are other less pleasant scenarios that may occur like: the application not taking redeployment very well, the number of replicas being limited due to a lack of nodes, our application running as a StatefulSet, etc. In the worst-case scenario, the Pods' resource consumption grows due to an increased load, HPA starts to scale up the cluster, and then, suddenly, VPA proceeds to modify the resource parameters and restarts the Pods. As a result, this high load gets distributed across the rest of the Pods. Some of them may crash, rendering things even worse and resulting in a chain reaction of failure.

可悲的是，可能会出现其他不太令人愉快的情况，例如：应用程序没有很好地重新部署，由于缺少节点而导致副本数量受到限制，我们的应用程序作为 StatefulSet 运行等。在最坏的情况下， Pod 的资源消耗由于负载增加而增加，HPA 开始扩展集群，然后，突然，VPA 继续修改资源参数并重新启动 Pod。结果，这个高负载分布在其余的 Pod 上。其中一些可能会崩溃，使事情变得更糟，并导致失败的连锁反应。

That is why having a profound understanding of various VPA operating modes is important. Let’s start with the simplest one — “Off”.

这就是为什么深入了解各种 VPA 操作模式很重要的原因。让我们从最简单的开始——“关闭”。

**Off mode**

**关闭模式**

All this mode does is calculate the resource consumption of Pods and make recommendations. Looking ahead, I would like to note that at Flant **we use this mode in the majority of cases** (and we recommend it). But first, let’s look at a few examples.

该模式所做的只是计算 Pod 的资源消耗并提出建议。展望未来，我想指出，在 Flant **我们在大多数情况下都使用这种模式**（我们推荐它）。但首先，让我们看几个例子。

Some basic manifests follow below:

一些基本清单如下：

```
apiVersion: autoscaling.k8s.io/v1
kind: VerticalPodAutoscaler
metadata:
name: my-app-vpa
spec:
targetRef:
    apiVersion: "apps/v1"
    kind: Deployment
    name: my-app
updatePolicy:
    updateMode: "Recreate"
    containerPolicies:
      - containerName: "*"
        minAllowed:
          cpu: 100m
          memory: 250Mi
        maxAllowed:
          cpu: 1
          memory: 500Mi
        controlledResources: ["cpu", "memory"]
        controlledValues: RequestsAndLimits
```

We will not go into detail about this manifest’s parameters: [this article](https://povilasv.me/vertical-pod-autoscaling-the-definitive-guide/) provides a detailed description of the features and aspects of VPA. In short, we specify the VPA target ( `targetRef`) and select the update policy. Additionally, we specify the upper and lower limits for the resources VPA can use. The primary focus is on the `updateMode` field. In “Recreate” or “Auto” mode, VPA will recreate Pods with all consequences (until an above-mentioned patch for in-place Pod resource parameters update becomes available). Since we don’t want it, we use the “Off” mode:

我们不会详细介绍此清单的参数：[本文](https://povilasv.me/vertical-pod-autoscaling-the-definitive-guide/) 提供了 VPA 的功能和方面的详细描述。简而言之，我们指定 VPA 目标（`targetRef`）并选择更新策略。此外，我们指定了 VPA 可以使用的资源的上限和下限。主要关注的是“updateMode”字段。在“重新创建”或“自动”模式下，VPA 将重新创建 Pod 并承担所有后果（直到上述用于就地 Pod 资源参数更新的补丁可用)。由于我们不想要它，我们使用“关闭”模式：

```
apiVersion: autoscaling.k8s.io/v1
kind: VerticalPodAutoscaler
metadata:
name: my-app-vpa
spec:
targetRef:
    apiVersion: "apps/v1"
    kind: Deployment
    name: my-app
updatePolicy:
    updateMode: "Off"   # !!!
resourcePolicy:
    containerPolicies:
      - containerName: "*"
        controlledResources: ["cpu", "memory"]
```

VPA starts collecting metrics. You can use the `kubectl describe vpa` command to see the recommendations (just let VPA run for a few minutes):

VPA 开始收集指标。您可以使用 `kubectl describe vpa` 命令查看建议（只需让 VPA 运行几分钟）：

```
Recommendation:
    Container Recommendations:
      Container Name:  nginx
      Lower Bound:
        Cpu:     25m
        Memory:  52428800
      Target:
        Cpu:     25m
        Memory:  52428800
      Uncapped Target:
        Cpu:     25m
        Memory:  52428800
      Upper Bound:
        Cpu:     25m
        Memory:  52428800
```

The VPA recommendations will be more accurate after a couple of days (a week, a month, etc.) of running. And then is the perfect time to adjust limits in the application manifest. That way, you can avoid OOM kills due to a lack of resources and save on infrastructure (if initial requests/limits are too high).

运行几天（一周、一个月等）后，VPA 建议将更加准确。然后是在应用程序清单中调整限制的最佳时机。这样，您可以避免因缺乏资源而导致 OOM 终止并节省基础架构（如果初始请求/限制太高）。

Now, let’s talk about some of the details of using VPA.

现在，让我们谈谈使用 VPA 的一些细节。

### Other VPA modes

### 其他 VPA 模式

Note that in “Initial” Mode, VPA assigns resources when Pods are started and never changes them later. Thus, VPA will set low requests/limits for newly created Pods if the load was relatively low over the past week. It may lead to problems if the load suddenly increases because the requests/limits will be much lower than what is required for such a load. This mode may come in handy if your load is uniformly distributed and grows in a linear fashion.

请注意，在“初始”模式下，VPA 在 Pod 启动时分配资源，以后不再更改它们。因此，如果过去一周的负载相对较低，VPA 将为新创建的 Pod 设置较低的请求/限制。如果负载突然增加，可能会导致问题，因为请求/限制将远低于此类负载所需的要求。如果您的负载均匀分布并以线性方式增长，则此模式可能会派上用场。

In “Auto” mode, VPA recreates the Pods. Thus, the application must handle the restart properly. If it cannot shutdown gracefully (i.e. by closing the existing connections correctly and so on), you will most likely catch some avoidable 5XX errors. Using Auto mode with a StatefulSet is rarely advisable: imagine VPA attempting to add PostgreSQL resources to production…

在“自动”模式下，VPA 会重新创建 Pod。因此，应用程序必须正确处理重启。如果它不能正常关闭（即通过正确关闭现有连接等），您很可能会遇到一些可避免的 5XX 错误。很少建议使用带有 StatefulSet 的 Auto 模式：想象一下 VPA 试图将 PostgreSQL 资源添加到生产环境......

As for the dev environment, you can freely experiment to find the level of resources to use (later) in production that is acceptable to you. Suppose you want to use VPA in the “Initial” mode and we have Redis in the cluster using the `maxmemory` parameter. You will most likely need to change it to adjust it to your needs. The problem is Redis doesn’t care about the limits at the cgroups level. In other words, you are risking a lot if `maxmemory` is, say, 2GB while your Pod’s memory is capped at 1GB. But how can you set `maxmemory` to be the same as the limit? Well, there is a way! You can use the VPA-recommended values:

至于开发环境，您可以自由试验以找到您可以接受的（稍后）生产中使用的资源级别。假设您想在“初始”模式下使用 VPA，并且我们在集群中使用“maxmemory”参数有 Redis。您很可能需要对其进行更改以适应您的需要。问题是 Redis 不关心 cgroups 级别的限制。换句话说，如果 `maxmemory` 是 2GB 而你的 Pod 的内存上限为 1GB，你会冒很大的风险。但是如何将 `maxmemory` 设置为与限制相同？嗯，有办法！您可以使用 VPA 推荐的值：

```
apiVersion: apps/v1
kind: Deployment
metadata:
name: redis
labels:
    app: redis
spec:
replicas: 1
selector:
    matchLabels:
      app: redis
template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
      - name: redis
        image: redis:6.2.1
        ports:
        - containerPort: 6379
        resources:
           requests:
             memory: "100Mi"
             cpu: "256m"
           limits:
             memory: "100Mi"
             cpu: "256m"
        env:
          - name: MY_MEM_REQUEST
            valueFrom:
              resourceFieldRef:
                containerName: app
                resource: requests.memory
          - name: MY_MEM_LIMIT
            valueFrom:
              resourceFieldRef:
                containerName: app
                resource: limits.memory
```

You can use environment variables to obtain the memory limit (and subtract, say, 10% from that for application needs) and set the resulting value as `maxmemory`. You will probably have to do something about the init container that uses `sed` to process the Redis config since the default Redis container image does not support passing `maxmemory` using an environment variable. Nevertheless, this solution is quite functional.

您可以使用环境变量来获取内存限制（并从应用程序需求中减去 10%）并将结果值设置为“maxmemory”。您可能需要对使用 `sed` 处理 Redis 配置的 init 容器做一些事情，因为默认的 Redis 容器映像不支持使用环境变量传递 `maxmemory`。然而，这个解决方案非常实用。

Finally, I would like to turn your attention to the fact that VPA evicts the DaemonSet Pods all at once, en masse. We are currently working on a [patch](https://github.com/kubernetes/kubernetes/pull/98307) that fixes this.

最后，我想将您的注意力转移到 VPA 一次性全部驱逐 DaemonSet Pod 的事实。我们目前正在研究解决此问题的 [patch](https://github.com/kubernetes/kubernetes/pull/98307)。

### Final VPA recommendations

### 最终 VPA 建议

“Off” mode is suitable for the majority of cases.

“关闭”模式适用于大多数情况。

You can experiment with “Auto” and “Initial” modes in the dev environment.

您可以在开发环境中试验“自动”和“初始”模式。

**Only use VPA in production if you have already accumulated recommendations and tested them thoroughly**. In addition, you have to clearly understand what you are doing and why you are doing it.

**如果您已经积累了建议并对其进行了彻底的测试，请仅在生产中使用 VPA **。此外，您必须清楚地了解自己在做什么以及为什么要这样做。

In the meantime, we are eagerly anticipating in-place (restart-free) updates for Pod resources. 

与此同时，我们热切期待 Pod 资源的就地（免重启）更新。

Note that there are some limitations associated with joint use of HPA and VPA. For instance, VPA should not be used together with HPA if the CPU- or Memory-based metric is used as a trigger. The reason is that when the threshold is reached, VPA increases resource requests/limits while HPA adds new replicas. Consequently, the load will drop off sharply, and the process will go in reverse, resulting in “flapping”. The [official documentation](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#known-limitations) sheds more light on the existing limitations.

请注意，联合使用 HPA 和 VPA 存在一些限制。例如，如果使用基于 CPU 或内存的指标作为触发器，则 VPA 不应与 HPA 一起使用。原因是当达到阈值时，VPA 会增加资源请求/限制，而 HPA 会添加新副本。因此，负载将急剧下降，过程将反向进行，导致“颤动”。 [官方文档](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#known-limitations) 更清楚地说明了现有的限制。

## Conclusion

##  结论

This concludes our review of best practices for deploying HA applications in Kubernetes. Please share your thoughts and suggestions!

我们对在 Kubernetes 中部署 HA 应用程序的最佳实践的回顾到此结束。请分享您的想法和建议！

## Related posts:

##  相关文章：

- [Migrating your app to Kubernetes: what to do with files?](https://blog.flant.com/migrating-your-app-to-kubernetes-what-to-do-with-files/ "Migrating your app to Kubernetes: what to do with files?")
- [ConfigMaps in Kubernetes: how they work and what you should remember](https://blog.flant.com/configmaps-in-kubernetes-how-they-work-and-what-you-should-remember/ "ConfigMaps in Kubernetes: how they work and what you should remember")
- [Best practices for deploying highly available apps in Kubernetes. Part 1](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/ "Best practices for deploying highly available apps in Kubernetes. Part 1" )

- [将您的应用迁移到 Kubernetes：如何处理文件？](https://blog.flant.com/migrating-your-app-to-kubernetes-what-to-do-with-files/ "迁移您的应用到 Kubernetes：如何处理文件？”)
- [Kubernetes 中的 ConfigMap：它们是如何工作的以及你应该记住什么](https://blog.flant.com/configmaps-in-kubernetes-how-they-work-and-what-you-should-remember/ "ConfigMaps在 Kubernetes 中：它们是如何工作的以及你应该记住什么”)
- [在 Kubernetes 中部署高可用性应用程序的最佳实践。第 1 部分](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/ “在 Kubernetes 中部署高可用性应用程序的最佳实践。第 1 部分” )

https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-2

