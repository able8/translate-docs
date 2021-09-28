# Horizontal Pod Autoscaler

# 水平 Pod 自动缩放器

The Horizontal Pod Autoscaler automatically scales the number of Pods in a replication controller, deployment, replica set or stateful set based on observed CPU utilization (or, with [custom metrics](https://git.k8s.io/community/contributors/design-proposals/instrumentation/custom-metrics-api.md) support, on some other application-provided metrics). Note that Horizontal Pod Autoscaling does not apply to objects that can't be scaled, for example, DaemonSets.

Horizontal Pod Autoscaler 根据观察到的 CPU 利用率（或使用 [自定义指标](https://git.k8s.io/community/contributors/design-proposals/instrumentation/custom-metrics-api.md）支持，其他一些应用程序提供的指标)。请注意，Horizontal Pod Autoscaling 不适用于无法缩放的对象，例如 DaemonSet。

The Horizontal Pod Autoscaler is implemented as a Kubernetes API  resource and a controller. The resource determines the behavior of the controller. The controller periodically adjusts the number of replicas in a  replication controller or deployment to match the observed metrics such  as average CPU utilisation, average memory utilisation or any other  custom metric to the target specified by the user.

Horizontal Pod Autoscaler 实现为 Kubernetes API 资源和控制器。资源决定了控制器的行为。控制器会定期调整复制控制器或部署中的副本数量，以将观察到的指标（例如平均 CPU 利用率、平均内存利用率或任何其他自定义指标）与用户指定的目标相匹配。

## How does the Horizontal Pod Autoscaler work?

## Horizontal Pod Autoscaler 如何工作？

![Horizontal Pod Autoscaler diagram](https://d33wubrfki0l68.cloudfront.net/4fe1ef7265a93f5f564bd3fbb0269ebd10b73b4e/1775d/images/docs/horizontal-pod-autoscaler.svg)

The Horizontal Pod Autoscaler is implemented as a control loop, with a period controlled by the controller manager's `--horizontal-pod-autoscaler-sync-period` flag (with a default value of 15 seconds).

Horizontal Pod Autoscaler 被实现为一个控制循环，其周期由控制器管理器的 `--horizontal-pod-autoscaler-sync-period` 标志控制（默认值为 15 秒）。

During each period, the controller manager queries the resource utilization against the metrics specified in each HorizontalPodAutoscaler definition. The controller manager obtains the metrics from either the resource metrics API (for per-pod resource metrics), or the custom metrics API (for all other metrics).

在每个时间段内，控制器管理器根据每个 HorizontalPodAutoscaler 定义中指定的指标查询资源利用率。控制器管理器从资源指标 API（对于每个 Pod 资源指标）或自定义指标 API（对于所有其他指标）获取指标。

- For per-pod resource metrics (like CPU), the controller fetches the metrics from the resource metrics API for each Pod targeted by the HorizontalPodAutoscaler. Then, if a target utilization value is set, the controller calculates the utilization value as a percentage of the equivalent resource request on the containers in each Pod. If a target raw value is set, the raw metric values are used directly. The controller then takes the mean of the utilization or the raw value (depending on the type of target specified) across all targeted Pods, and produces a ratio used to scale the number of desired replicas.

   - 对于每个 Pod 的资源指标（如 CPU），控制器从 HorizontalPodAutoscaler 所针对的每个 Pod 的资源指标 API 中获取指标。然后，如果设置了目标利用率值，则控制器将利用率值计算为每个 Pod 中容器上的等效资源请求的百分比。如果设置了目标原始值，则直接使用原始度量值。然后，控制器取所有目标 Pod 的利用率或原始值（取决于指定的目标类型）的平均值，并生成用于扩展所需副本数量的比率。

  Please note that if some of the Pod's containers do not have the relevant resource request set, CPU utilization for the Pod will not be defined and the autoscaler will not take any action for that metric. See the [algorithm details](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details) section below for more information about how the autoscaling algorithm works.

请注意，如果 Pod 的某些容器没有设置相关的资源请求，则不会定义 Pod 的 CPU 利用率，并且自动缩放器不会针对该指标采取任何行动。有关自动缩放算法如何工作的更多信息，请参阅下面的 [算法详细信息](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details) 部分。

- For per-pod custom metrics, the controller functions similarly to per-pod resource metrics, except that it works with raw values, not utilization values.

- 对于每个 Pod 的自定义指标，控制器的功能类似于每个 Pod 的资源指标，不同之处在于它使用原始值，而不是利用率值。

- For object metrics and external metrics, a single metric is fetched, which describes the object in question. This metric is compared to the target value, to produce a ratio as above. In the `autoscaling/v2beta2` API version, this value can optionally be divided by the number of Pods before the comparison is made.

- 对于对象指标和外部指标，获取单个指标，该指标描述相关对象。将该指标与目标值进行比较，以产生上述比率。在`autoscaling/v2beta2` API 版本中，在进行比较之前，可以选择将此值除以 Pod 的数量。

The HorizontalPodAutoscaler normally fetches metrics from a series of aggregated APIs (`metrics.k8s.io`, `custom.metrics.k8s.io`, and `external.metrics.k8s.io`). The `metrics.k8s.io` API is usually provided by metrics-server, which needs to be launched separately. For more information about resource metrics, see [Metrics Server](https://kubernetes.io/docs/tasks/debug-application-cluster/resource-metrics-pipeline/#metrics-server).

HorizontalPodAutoscaler 通常从一系列聚合的 API（`metrics.k8s.io`、`custom.metrics.k8s.io` 和 `external.metrics.k8s.io`）中获取指标。 `metrics.k8s.io` API 通常由 metrics-server 提供，需要单独启动。有关资源指标的更多信息，请参阅 [指标服务器](https://kubernetes.io/docs/tasks/debug-application-cluster/resource-metrics-pipeline/#metrics-server)。

See [Support for metrics APIs](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-metrics-apis) for more details. 

有关更多详细信息，请参阅 [支持指标 API](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-metrics-apis)。

The autoscaler accesses corresponding scalable controllers (such as replication controllers, deployments, and replica sets) by using the scale sub-resource. Scale is an interface that allows you to dynamically set the number of replicas and examine each of their current states. More details on scale sub-resource can be found [here](https://git.k8s.io/community/contributors/design-proposals/autoscaling/horizontal-pod-autoscaler.md#scale-subresource).

autoscaler 使用 scale 子资源访问对应的可伸缩控制器（如复制控制器、部署和副本集）。 Scale 是一个界面，允许您动态设置副本的数量并检查它们的每个当前状态。有关 scale 子资源的更多详细信息，请参见 [此处](https://git.k8s.io/community/contributors/design-proposals/autoscaling/horizontal-pod-autoscaler.md#scale-subresource)。

### Algorithm Details

### 算法细节

From the most basic perspective, the Horizontal Pod Autoscaler controller operates on the ratio between desired metric value and current metric value:

从最基本的角度来看，Horizontal Pod Autoscaler 控制器根据所需指标值和当前指标值之间的比率进行操作：

```
desiredReplicas = ceil[currentReplicas * ( currentMetricValue / desiredMetricValue )]
```

For example, if the current metric value is `200m`, and the desired value is `100m`, the number of replicas will be doubled, since `200.0 / 100.0 == 2.0` If the current value is instead `50m`, we 'll halve the number of replicas, since `50.0 / 100.0 == 0.5`. We'll skip scaling if the ratio is sufficiently close to 1.0 (within a globally-configurable tolerance, from the `--horizontal-pod-autoscaler-tolerance` flag, which defaults to 0.1).

例如，如果当前度量值为‘200m’，而期望值为‘100m’，则副本数将翻倍，因为‘200.0 / 100.0 == 2.0’如果当前值为‘50m’，我们将副本数量减半，因为`50.0 / 100.0 == 0.5`。如果比率足够接近 1.0（在全局可配置容差范围内，来自 `--horizontal-pod-autoscaler-tolerance` 标志，默认为 0.1），我们将跳过缩放。

When a `targetAverageValue` or `targetAverageUtilization` is specified, the `currentMetricValue` is computed by taking the average of the given metric across all Pods in the HorizontalPodAutoscaler's scale target. Before checking the tolerance and deciding on the final values, we take pod readiness and missing metrics into consideration, however.

当指定了 `targetAverageValue` 或 `targetAverageUtilization` 时，通过取 HorizontalPodAutoscaler 的缩放目标中所有 Pod 的给定指标的平均值来计算 `currentMetricValue`。然而，在检查容差和决定最终值之前，我们会考虑 pod 准备情况和缺失的指标。

All Pods with a deletion timestamp set (i.e. Pods in the process of being shut down) and all failed Pods are discarded.

设置了删除时间戳的所有 Pod（即处于关闭过程中的 Pod）和所有失败的 Pod 都将被丢弃。

If a particular Pod is missing metrics, it is set aside for later; Pods with missing metrics will be used to adjust the final scaling amount.

如果特定 Pod 缺少指标，则将其留待以后使用；缺少指标的 Pod 将用于调整最终的扩展量。

When scaling on CPU, if any pod has yet to become ready (i.e. it's still initializing) *or* the most recent metric point for the pod was before it became ready, that pod is set aside as well.

在 CPU 上扩展时，如果任何 pod 尚未准备好（即它仍在初始化）*或* pod 的最新指标点是在它准备好之前，则该 pod 也会被搁置一旁。

Due to technical constraints, the HorizontalPodAutoscaler controller cannot exactly determine the first time a pod becomes ready when determining whether to set aside certain CPU metrics. Instead, it considers a Pod "not yet ready" if it's unready and transitioned to unready within a short, configurable window of time since it started. This value is configured with the `--horizontal-pod-autoscaler-initial-readiness-delay` flag, and its default is 30 seconds. Once a pod has become ready, it considers any transition to ready to be the first if it occurred within a longer, configurable time since it started. This value is configured with the `--horizontal-pod-autoscaler-cpu-initialization-period` flag, and its default is 5 minutes.

由于技术限制，HorizontalPodAutoscaler 控制器在确定是否留出某些 CPU 指标时无法准确确定 Pod 的第一次准备就绪。相反，如果 Pod 未准备好并在它启动后的一个可配置的短时间窗口内转换为未准备好，它会认为 Pod“尚未准备好”。这个值是用 `--horizontal-pod-autoscaler-initial-readiness-delay` 标志配置的，它的默认值为 30 秒。一旦 Pod 准备就绪，如果它发生在自启动以来更长的、可配置的时间内，它就会将任何过渡到就绪状态视为第一个。这个值是用`--horizontal-pod-autoscaler-cpu-initialization-period`标志配置的，它的默认值为5分钟。

The `currentMetricValue / desiredMetricValue` base scale ratio is then calculated using the remaining pods not set aside or discarded from above.

然后使用剩余的 pod 计算“currentMetricValue/desiredMetricValue”基本比例比率，而不是从上面留出或丢弃。

If there were any missing metrics, we recompute the average more conservatively, assuming those pods were consuming 100% of the desired value in case of a scale down, and 0% in case of a scale up. This dampens the magnitude of any potential scale.

如果有任何缺失的指标，我们会更保守地重新计算平均值，假设这些 Pod 在缩减时消耗了 100% 的期望值，在扩展时消耗了 0%。这会抑制任何潜在规模的大小。

Furthermore, if any not-yet-ready pods were present, and we would have scaled up without factoring in missing metrics or not-yet-ready pods, we conservatively assume the not-yet-ready pods are consuming 0% of the desired metric , further dampening the magnitude of a scale up.

此外，如果存在任何尚未准备好的 pod，并且我们会在不考虑缺失指标或尚未准备好的 pod 的情况下进行扩展，我们保守地假设尚未准备好的 pod 消耗了所需指标的 0% ，进一步抑制放大的幅度。

After factoring in the not-yet-ready pods and missing metrics, we recalculate the usage ratio. If the new ratio reverses the scale direction, or is within the tolerance, we skip scaling. Otherwise, we use the new ratio to scale.

在考虑尚未准备好的 pod 和缺失的指标后，我们重新计算使用率。如果新比率与缩放方向相反，或者在容差范围内，我们将跳过缩放。否则，我们使用新的比率进行缩放。

Note that the *original* value for the average utilization is reported back via the HorizontalPodAutoscaler status, without factoring in the not-yet-ready pods or missing metrics, even when the new usage ratio is used. 

请注意，平均利用率的*原始*值是通过 HorizontalPodAutoscaler 状态报告回来的，即使使用了新的使用率，也不会考虑尚未准备好的 pod 或缺少的指标。

If multiple metrics are specified in a HorizontalPodAutoscaler, this calculation is done for each metric, and then the largest of the desired replica counts is chosen. If any of these metrics cannot be converted into a desired replica count (e.g. due to an error fetching the metrics from the metrics APIs) and a scale down is suggested by the metrics which can be fetched, scaling is skipped. This means that the HPA is still capable of scaling up if one or more metrics give a `desiredReplicas` greater than the current value.

如果在 HorizontalPodAutoscaler 中指定了多个指标，则对每个指标进行此计算，然后选择所需副本计数中最大的一个。如果这些指标中的任何一个无法转换为所需的副本计数（例如，由于从指标 API 获取指标的错误）并且可以获取的指标建议缩小规模，则跳过缩放。这意味着如果一个或多个指标给出的“desiredReplicas”大于当前值，HPA 仍然能够扩展。

Finally, right before HPA scales the target, the scale recommendation is recorded. The controller considers all recommendations within a configurable window choosing the highest recommendation from within that window. This value can be configured using the `--horizontal-pod-autoscaler-downscale-stabilization` flag, which defaults to 5 minutes. This means that scaledowns will occur gradually, smoothing out the impact of rapidly fluctuating metric values.

最后，就在 HPA 缩放目标之前，记录缩放建议。控制器考虑可配置窗口内的所有推荐，从该窗口中选择最高推荐。这个值可以使用 `--horizontal-pod-autoscaler-downscale-stabilization` 标志配置，默认为 5 分钟。这意味着缩减将逐渐发生，从而消除快速波动的指标值的影响。

## API Object

## API 对象

The Horizontal Pod Autoscaler is an API resource in the Kubernetes `autoscaling` API group. The current stable version, which only includes support for CPU autoscaling, can be found in the `autoscaling/v1` API version.

Horizontal Pod Autoscaler 是 Kubernetes `autoscaling` API 组中的一个 API 资源。当前的稳定版本仅支持 CPU 自动缩放，可以在 `autoscaling/v1` API 版本中找到。

The beta version, which includes support for scaling on memory and custom metrics, can be found in `autoscaling/v2beta2`. The new fields introduced in `autoscaling/v2beta2` are preserved as annotations when working with `autoscaling/v1`.

测试版包括对内存扩展和自定义指标的支持，可以在 `autoscaling/v2beta2` 中找到。在使用“autoscaling/v1”时，“autoscaling/v2beta2”中引入的新字段作为注释保留。

When you create a HorizontalPodAutoscaler API object, make sure the name specified is a valid [DNS subdomain name](https://kubernetes.io/docs/concepts/overview/working-with-objects/names#dns-subdomain-names) . More details about the API object can be found at [HorizontalPodAutoscaler Object](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#horizontalpodautoscaler-v1-autoscaling).

创建 HorizontalPodAutoscaler API 对象时，请确保指定的名称是有效的 [DNS 子域名](https://kubernetes.io/docs/concepts/overview/working-with-objects/names#dns-subdomain-names) .有关 API 对象的更多详细信息，请参见 [HorizontalPodAutoscaler 对象](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#horizontalpodautoscaler-v1-autoscaling)。

## Support for Horizontal Pod Autoscaler in kubectl

## 支持 kubectl 中的 Horizontal Pod Autoscaler

Horizontal Pod Autoscaler, like every API resource, is supported in a standard way by `kubectl`. We can create a new autoscaler using `kubectl create` command. We can list autoscalers by `kubectl get hpa` and get detailed description by `kubectl describe hpa`. Finally, we can delete an autoscaler using `kubectl delete hpa`.

Horizontal Pod Autoscaler 与每个 API 资源一样，由 `kubectl` 以标准方式支持。我们可以使用 `kubectl create` 命令创建一个新的自动缩放器。我们可以通过 `kubectl get hpa` 列出 autoscalers，通过 `kubectl describe hpa` 获取详细描述。最后，我们可以使用 `kubectl delete hpa` 删除自动缩放器。

In addition, there is a special `kubectl autoscale` command for creating a HorizontalPodAutoscaler object. For instance, executing `kubectl autoscale rs foo --min=2 --max=5 --cpu-percent=80` will create an autoscaler for replication set *foo*, with target CPU utilization set to `80%` and the number of replicas between 2 and 5. The detailed documentation of `kubectl autoscale` can be found [here](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands/#autoscale).

此外，还有一个特殊的 `kubectl autoscale` 命令用于创建 HorizontalPodAutoscaler 对象。例如，执行 `kubectl autoscale rs foo --min=2 --max=5 --cpu-percent=80` 将为复制集 *foo* 创建一个自动缩放器，目标 CPU 利用率设置为 `80%`，并且2 到 5 之间的副本数。可以在 [此处](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands/#autoscale) 找到“kubectl autoscale”的详细文档。

## Autoscaling during rolling update

## 在滚动更新期间自动缩放

Kubernetes lets you perform a rolling update on a Deployment. In that case, the Deployment manages the underlying ReplicaSets for you. When you configure autoscaling for a Deployment, you bind a HorizontalPodAutoscaler to a single Deployment. The HorizontalPodAutoscaler manages the `replicas` field of the Deployment. The deployment controller is responsible for setting the `replicas` of the underlying ReplicaSets so that they add up to a suitable number during the rollout and also afterwards.

Kubernetes 允许您对部署执行滚动更新。在这种情况下，Deployment 会为您管理底层的 ReplicaSet。为 Deployment 配置自动缩放时，您将 HorizontalPodAutoscaler 绑定到单个 Deployment。 HorizontalPodAutoscaler 管理 Deployment 的 `replicas` 字段。部署控制器负责设置底层 ReplicaSet 的“副本”，以便它们在部署期间和之后加起来为合适的数量。

If you perform a rolling update of a StatefulSet that has an autoscaled number of replicas, the StatefulSet directly manages its set of Pods (there is no intermediate resource similar to ReplicaSet).

如果您对具有自动缩放数量的副本的 StatefulSet 执行滚动更新，则 StatefulSet 直接管理其 Pod 集（没有类似于 ReplicaSet 的中间资源）。

## Support for cooldown/delay

## 支持冷却/延迟

When managing the scale of a group of replicas using the Horizontal Pod Autoscaler, it is possible that the number of replicas keeps fluctuating frequently due to the dynamic nature of the metrics evaluated. This is sometimes referred to as *thrashing*.

在使用 Horizontal Pod Autoscaler 管理一组副本的规模时，由于所评估指标的动态特性，副本数量可能会经常波动。这有时被称为*颠簸*。

Starting from v1.6, a cluster operator can mitigate this problem by tuning the global HPA settings exposed as flags for the `kube-controller-manager` component:

从 v1.6 开始，集群操作员可以通过调整作为 `kube-controller-manager` 组件标志公开的全局 HPA 设置来缓解这个问题：

Starting from v1.12, a new algorithmic update removes the need for the upscale delay. 

从 v1.12 开始，新的算法更新消除了对高档延迟的需要。

- `--horizontal-pod-autoscaler-downscale-stabilization`: Specifies the duration of the downscale stabilization time window. Horizontal Pod Autoscaler remembers the historical recommended sizes and only acts on the largest size within this time window. The default value is 5 minutes (`5m0s`).

- `--horizontal-pod-autoscaler-downscale-stabilization`：指定缩减稳定时间窗口的持续时间。 Horizontal Pod Autoscaler 会记住历史推荐的大小，并且只对这个时间窗口内的最大大小起作用。默认值为 5 分钟（`5m0s`）。

> **Note:** When tuning these parameter values, a cluster operator should be aware of the possible consequences. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.

> **注意：** 在调整这些参数值时，集群操作员应该意识到可能的后果。如果延迟（冷却）值设置得太长，可能会有人抱怨 Horizontal Pod Autoscaler 对工作负载变化没有响应。但是，如果延迟值设置得太短，副本集的规模可能会像往常一样不断抖动。

## Support for resource metrics

## 支持资源指标

Any HPA target can be scaled based on the resource usage of the pods in the scaling target. When defining the pod specification the resource requests like `cpu` and `memory` should be specified. This is used to determine the resource utilization and used by the HPA controller to scale the target up or down. To use resource utilization based scaling specify a metric source like this:

任何 HPA 目标都可以根据扩展目标中 Pod 的资源使用情况进行扩展。在定义 pod 规范时，应该指定像 `cpu` 和 `memory` 这样的资源请求。这用于确定资源利用率，并由 HPA 控制器用于扩大或缩小目标。要使用基于资源利用率的扩展，请指定如下指标源：

```yaml
type: Resource
resource:
  name: cpu
  target:
    type: Utilization
    averageUtilization: 60
```

With this metric the HPA controller will keep the average utilization of the pods in the scaling target at 60%. Utilization is the ratio between the current usage of resource to the requested resources of the pod. See [Algorithm](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details) for more details about how the utilization is calculated and averaged.

使用此指标，HPA 控制器会将扩展目标中 Pod 的平均利用率保持在 60%。利用率是当前资源使用量与 Pod 请求资源之间的比率。有关如何计算和平均利用率的更多详细信息，请参阅 [算法](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details)。

> **Note:** Since the resource usages of all the containers are summed up the total pod utilization may not accurately represent the individual container resource usage. This could lead to situations where a single container might be running with high usage and the HPA will not scale out because the overall pod usage is still within acceptable limits.

> **注意：** 由于所有容器的资源使用情况是相加的，因此总的 pod 使用情况可能无法准确代表单个容器的资源使用情况。这可能导致单个容器可能以高使用率运行并且 HPA 不会横向扩展的情况，因为整体 pod 使用率仍在可接受的范围内。

### Container Resource Metrics

### 容器资源指标

**FEATURE STATE:** `Kubernetes v1.20 [alpha]`

**功能状态：** `Kubernetes v1.20 [alpha]`

`HorizontalPodAutoscaler` also supports a container metric source where the HPA can track the resource usage of individual containers across a set of Pods, in order to scale the target resource. This lets you configure scaling thresholds for the containers that matter most in a particular Pod. For example, if you have a web application and a logging sidecar, you can scale based on the resource use of the web application, ignoring the sidecar container and its resource use.

`HorizontalPodAutoscaler` 还支持容器指标源，HPA 可以在其中跟踪一组 Pod 中各个容器的资源使用情况，以扩展目标资源。这使您可以为特定 Pod 中最重要的容器配置扩展阈值。例如，如果您有一个 Web 应用程序和一个日志 sidecar，您可以根据 Web 应用程序的资源使用情况进行扩展，而忽略 sidecar 容器及其资源使用情况。

If you revise the target resource to have a new Pod specification with a different set of containers, you should revise the HPA spec if that newly added container should also be used for scaling. If the specified container in the metric source is not present or only present in a subset of the pods then those pods are ignored and the recommendation is recalculated. See [Algorithm](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details) for more details about the calculation. To use container resources for autoscaling define a metric source as follows:

如果您修改目标资源以获得具有不同容器集的新 Pod 规范，并且如果新添加的容器也应用于扩展，则您应该修改 HPA 规范。如果指标源中的指定容器不存在或仅存在于 pod 的子集中，则这些 pod 将被忽略并重新计算建议。有关计算的更多详细信息，请参阅[算法](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details)。要使用容器资源进行自动缩放，请按如下方式定义指标源：

```yaml
type: ContainerResource
containerResource:
  name: cpu
  container: application
  target:
    type: Utilization
    averageUtilization: 60
```

In the above example the HPA controller scales the target such that the average utilization of the cpu in the `application` container of all the pods is 60%.

在上面的例子中，HPA 控制器缩放目标，使得所有 pod 的“应用程序”容器中 CPU 的平均利用率为 60%。

> **Note:**
>
> If you change the name of a container that a HorizontalPodAutoscaler is tracking, you can make that change in a specific order to ensure scaling remains available and effective whilst the change is being applied. Before you update the resource that defines the container (such as a Deployment), you should update the associated HPA to track both the new and old container names. This way, the HPA is able to calculate a scaling recommendation throughout the update process.
>
> Once you have rolled out the container name change to the workload resource, tidy up by removing the old container name from the HPA specification.

> **注意：**
>
> 如果您更改 HorizontalPodAutoscaler 正在跟踪的容器的名称，您可以按特定顺序进行更改，以确保在应用更改时缩放仍然可用和有效。在更新定义容器的资源（例如部署）之前，您应该更新关联的 HPA 以跟踪新旧容器名称。这样，HPA 就能够在整个更新过程中计算扩展建议。
>
> 将容器名称更改推出到工作负载资源后，通过从 HPA 规范中删除旧容器名称来进行整理。

## Support for multiple metrics 

## 支持多个指标

Kubernetes 1.6 adds support for scaling based on multiple metrics. You can use the `autoscaling/v2beta2` API version to specify multiple metrics for the Horizontal Pod Autoscaler to scale on. Then, the Horizontal Pod Autoscaler controller will evaluate each metric, and propose a new scale based on that metric. The largest of the proposed scales will be used as the new scale.

Kubernetes 1.6 添加了对基于多个指标的扩展支持。您可以使用 `autoscaling/v2beta2` API 版本为 Horizontal Pod Autoscaler 指定多个指标以进行扩展。然后，Horizontal Pod Autoscaler 控制器将评估每个指标，并基于该指标提出一个新的规模。建议的比例尺中最大的将用作新比例尺。

## Support for custom metrics

## 支持自定义指标

> **Note:** Kubernetes 1.2 added alpha support for  scaling based on application-specific metrics using special annotations. Support for these annotations was removed in Kubernetes 1.6 in favor of  the new autoscaling API. While the old method for collecting custom metrics is still available, these metrics will not be available  for use by the Horizontal Pod Autoscaler, and the former annotations for specifying which custom metrics to scale on are no  longer honored by the Horizontal Pod Autoscaler controller.

> **注意：** Kubernetes 1.2 添加了 alpha 支持，以使用特殊注释基于特定于应用程序的指标进行缩放。 Kubernetes 1.6 中删除了对这些注释的支持，以支持新的自动缩放 API。虽然收集自定义指标的旧方法仍然可用，但 Horizontal Pod Autoscaler 将无法使用这些指标，并且 Horizontal Pod Autoscaler 控制器不再支持以前用于指定要扩展的自定义指标的注释。

Kubernetes 1.6 adds support for making use of custom metrics in the Horizontal Pod Autoscaler. You can add custom metrics for the Horizontal Pod Autoscaler to use in the `autoscaling/v2beta2` API. Kubernetes then queries the new custom metrics API to fetch the values of the appropriate custom metrics.

Kubernetes 1.6 添加了对在 Horizontal Pod Autoscaler 中使用自定义指标的支持。您可以为 Horizontal Pod Autoscaler 添加自定义指标以在 `autoscaling/v2beta2` API 中使用。 Kubernetes 然后查询新的自定义指标 API 以获取适当的自定义指标的值。

See [Support for metrics APIs](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-metrics-apis) for the requirements.

有关要求，请参阅 [支持指标 API](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-metrics-apis)。

## Support for metrics APIs

## 支持指标 API

By default, the HorizontalPodAutoscaler controller retrieves metrics from a series of APIs. In order for it to access these APIs, cluster administrators must ensure that:

默认情况下，HorizontalPodAutoscaler 控制器从一系列 API 中检索指标。为了访问这些 API，集群管理员必须确保：

- The [API aggregation layer](https://kubernetes.io/docs/tasks/extend-kubernetes/configure-aggregation-layer/) is enabled.
- The corresponding APIs are registered:
   - For resource metrics, this is the `metrics.k8s.io` API, generally provided by [metrics-server](https://github.com/kubernetes-sigs/metrics-server). It can be launched as a cluster addon.
   - For custom metrics, this is the `custom.metrics.k8s.io` API. It's provided by "adapter" API servers provided by metrics solution vendors. Check with your metrics pipeline, or the [list of known solutions](https://github.com/kubernetes/metrics/blob/master/IMPLEMENTATIONS.md#custom-metrics-api). If you would like to write your own, check out the [boilerplate](https://github.com/kubernetes-sigs/custom-metrics-apiserver) to get started.
   - For external metrics, this is the `external.metrics.k8s.io` API. It may be provided by the custom metrics adapters provided above.

- [API聚合层](https://kubernetes.io/docs/tasks/extend-kubernetes/configure-aggregation-layer/)已启用。
- 注册了相应的 API：
  - 对于资源指标，这是`metrics.k8s.io` API，一般由[metrics-server](https://github.com/kubernetes-sigs/metrics-server)提供。它可以作为集群插件启动。
  - 对于自定义指标，这是`custom.metrics.k8s.io` API。它由指标解决方案供应商提供的“适配器”API 服务器提供。检查您的指标管道，或 [已知解决方案列表](https://github.com/kubernetes/metrics/blob/master/IMPLEMENTATIONS.md#custom-metrics-api)。如果您想自己编写，请查看[boilerplate](https://github.com/kubernetes-sigs/custom-metrics-apiserver) 以开始使用。
  - 对于外部指标，这是`external.metrics.k8s.io` API。它可以由上面提供的自定义指标适配器提供。

For more information on these different metrics paths and how they differ please see the relevant design proposals for [the HPA V2](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/hpa-v2.md), [custom.metrics.k8s.io](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/custom-metrics-api.md) and [ external.metrics.k8s.io](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/external-metrics-api.md).

有关这些不同指标路径及其差异的更多信息，请参阅 [HPA V2] 的相关设计提案（https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/hpa -v2.md)、[custom.metrics.k8s.io](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/custom-metrics-api.md) 和 [ external.metrics.k8s.io](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/external-metrics-api.md)。

For examples of how to use them see [the walkthrough for using custom metrics](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#autoscaling-on-multiple-metrics-and-custom-metrics) and [the walkthrough for using external metrics](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#autoscaling-on-metrics-not-related-to-kubernetes-objects).

有关如何使用它们的示例，请参阅 [使用自定义指标的演练](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#autoscaling-on-multiple-metrics-and-custom-metrics) 和 [使用外部指标的演练](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#autoscaling-on-metrics-not-related-to-kubernetes-objects)。

## Support for configurable scaling behavior 

## 支持可配置的缩放行为

Starting from [v1.18](https://github.com/kubernetes/enhancements/blob/master/keps/sig-autoscaling/853-configurable-hpa-scale-velocity/README.md) the `v2beta2` API allows scaling behavior to be configured through the HPA `behavior` field. Behaviors are specified separately for scaling up and down in `scaleUp` or `scaleDown` section under the `behavior` field. A stabilization window can be specified for both directions which prevents the flapping of the number of the replicas in the scaling target. Similarly specifying scaling policies controls the rate of change of replicas while scaling.

从 [v1.18](https://github.com/kubernetes/enhancements/blob/master/keps/sig-autoscaling/853-configurable-hpa-scale-velocity/README.md) 开始，`v2beta2` API 允许要通过 HPA `behavior` 字段配置的缩放行为。在 `behavior` 字段下的 `scaleUp` 或 `scaleDown` 部分中分别指定用于放大和缩小的行为。可以为两个方向指定一个稳定窗口，以防止缩放目标中的副本数量波动。类似地，指定扩展策略可以控制扩展时副本的变化率。

### Scaling Policies

### 扩展策略

One or more scaling policies can be specified in the `behavior` section of the spec. When multiple policies are specified the policy which allows the highest amount of change is the policy which is selected by default. The following example shows this behavior while scaling down:

可以在规范的“行为”部分指定一个或多个扩展策略。当指定多个策略时，允许最大更改量的策略是默认选择的策略。以下示例显示了缩小时的这种行为：

```yaml
behavior:
  scaleDown:
    policies:
    - type: Pods
      value: 4
      periodSeconds: 60
    - type: Percent
      value: 10
      periodSeconds: 60
```

`periodSeconds` indicates the length of time in the past for which the policy must hold true. The first policy *(Pods)* allows at most 4 replicas to be scaled down in one minute. The second policy *(Percent)* allows at most 10% of the current replicas to be scaled down in one minute.

`periodSeconds` 表示该策略必须在过去的时间长度。第一个策略 *(Pods)* 允许在一分钟内最多缩减 4 个副本。第二个策略 *(Percent)* 允许在一分钟内缩减最多 10% 的当前副本。

Since by default the policy which allows the highest amount of change is selected, the second policy will only be used when the number of pod replicas is more than 40. With 40 or less replicas, the first policy will be applied. For instance if there are 80 replicas and the target has to be scaled down to 10 replicas then during the first step 8 replicas will be reduced. In the next iteration when the number of replicas is 72, 10% of the pods is 7.2 but the number is rounded up to 8. On each loop of the autoscaler controller the number of pods to be change is re-calculated based on the number of current replicas. When the number of replicas falls below 40 the first policy *(Pods)* is applied and 4 replicas will be reduced at a time.

由于默认选择了允许最大更改量的策略，因此仅当 pod 副本数大于 40 时才会使用第二个策略。如果副本数为 40 或更少，则将应用第一个策略。例如，如果有 80 个副本并且目标必须缩小到 10 个副本，那么在第一步期间将减少 8 个副本。在下一次迭代中，当副本数为 72 时，10% 的 Pod 为 7.2，但数字向上取整为 8。在 autoscaler 控制器的每个循环中，根据数量重新计算要更改的 Pod 数量当前副本的数量。当副本数低于 40 时，将应用第一个策略 *(Pods)*，一次减少 4 个副本。

The policy selection can be changed by specifying the `selectPolicy` field for a scaling direction. By setting the value to `Min` which would select the policy which allows the smallest change in the replica count. Setting the value to `Disabled` completely disables scaling in that direction.

可以通过为缩放方向指定 `selectPolicy` 字段来更改策略选择。通过将值设置为“Min”，这将选择允许副本计数变化最小的策略。将值设置为“禁用”将完全禁用该方向的缩放。

### Stabilization Window

### 稳定窗口

The stabilization window is used to restrict the flapping of replicas when the metrics used for scaling keep fluctuating. The stabilization window is used by the autoscaling algorithm to consider the computed desired state from the past to prevent scaling. In the following example the stabilization window is specified for `scaleDown`.

当用于扩展的指标不断波动时，稳定窗口用于限制副本的摆动。自动缩放算法使用稳定窗口来考虑从过去计算出的所需状态以防止缩放。在下面的示例中，为“scaleDown”指定了稳定窗口。

```yaml
scaleDown:
  stabilizationWindowSeconds: 300
```

When the metrics indicate that the target should be scaled down the algorithm looks into previously computed desired states and uses the highest value from the specified interval. In above example all desired states from the past 5 minutes will be considered.

当指标表明目标应按比例缩小时，算法会查看先前计算的所需状态并使用指定间隔中的最高值。在上面的示例中，将考虑过去 5 分钟的所有期望状态。

### Default Behavior

### 默认行为

To use the custom scaling not all fields have to be specified. Only values which need to be customized can be specified. These custom values are merged with default values. The default values match the existing behavior in the HPA algorithm.

要使用自定义缩放，不必指定所有字段。只能指定需要自定义的值。这些自定义值与默认值合并。默认值与 HPA 算法中的现有行为相匹配。

```yaml
behavior:
  scaleDown:
    stabilizationWindowSeconds: 300
    policies:
    - type: Percent
      value: 100
      periodSeconds: 15
  scaleUp:
    stabilizationWindowSeconds: 0
    policies:
    - type: Percent
      value: 100
      periodSeconds: 15
    - type: Pods
      value: 4
      periodSeconds: 15
    selectPolicy: Max
```

For scaling down the stabilization window is *300* seconds (or the value of the `--horizontal-pod-autoscaler-downscale-stabilization` flag if provided). There is only a single policy for scaling down which allows a 100% of the currently running replicas to be removed which means the scaling target can be scaled down to the minimum allowed replicas. For scaling up there is no stabilization window. When the metrics indicate that the target should be scaled up the target is scaled up immediately. There are 2 policies where 4 pods or a 100% of the currently running replicas will be added every 15 seconds till the HPA reaches its steady state.

对于缩小稳定窗口是 *300* 秒（或 `--horizontal-pod-autoscaler-downscale-stabilization` 标志的值，如果提供的话）。只有一个用于缩减的策略允许 100% 的当前正在运行的副本被删除，这意味着可以将扩展目标缩减到允许的最小副本。对于放大，没有稳定窗口。当指标表明应该扩大目标时，目标会立即扩大。有 2 个策略，每 15 秒添加 4 个 pod 或 100% 的当前运行的副本，直到 HPA 达到稳定状态。

### Example: change downscale stabilization window

### 示例：更改缩小稳定窗口

To provide a custom downscale stabilization window of 1 minute, the following behavior would be added to the HPA:

为了提供 1 分钟的自定义缩减稳定窗口，将向 HPA 添加以下行为：

```yaml
behavior:
  scaleDown:
    stabilizationWindowSeconds: 60
```

### Example: limit scale down rate

### 示例：限制缩小率

To limit the rate at which pods are removed by the HPA to 10% per minute, the following behavior would be added to the HPA:

要将 HPA 删除 Pod 的速率限制为每分钟 10%，将向 HPA 添加以下行为：

```yaml
behavior:
  scaleDown:
    policies:
    - type: Percent
      value: 10
      periodSeconds: 60
```

To ensure that no more than 5 Pods are removed per minute, you can add a second scale-down policy with a fixed size of 5, and set `selectPolicy` to minimum. Setting `selectPolicy` to `Min` means that the autoscaler chooses the policy that affects the smallest number of Pods:

为确保每分钟删除的 Pod 不超过 5 个，您可以添加固定大小为 5 的第二个缩减策略，并将 `selectPolicy` 设置为最小值。将 `selectPolicy` 设置为 `Min` 意味着自动缩放器选择影响最少 Pod 数量的策略：

```yaml
behavior:
  scaleDown:
    policies:
    - type: Percent
      value: 10
      periodSeconds: 60
    - type: Pods
      value: 5
      periodSeconds: 60
    selectPolicy: Min
```

### Example: disable scale down

### 示例：禁用缩小

The `selectPolicy` value of `Disabled` turns off scaling the given direction. So to prevent downscaling the following policy would be used:

`Disabled` 的 `selectPolicy` 值关闭给定方向的缩放。因此，为了防止缩小规模，将使用以下策略：

```yaml
behavior:
  scaleDown:
    selectPolicy: Disabled
```

## Implicit maintenance-mode deactivation

## 隐式维护模式停用

You can implicitly deactivate the HPA for a target without the need to change the HPA configuration itself. If the target's desired replica count is set to 0, and the HPA's minimum replica count is greater than 0, the HPA stops adjusting the target (and sets the `ScalingActive` Condition on itself to `false`) until you reactivate it by manually adjusting the target's desired replica count or HPA's minimum replica count.

您可以隐式停用目标的 HPA，而无需更改 HPA 配置本身。如果目标的所需副本数设置为 0，并且 HPA 的最小副本数大于 0，则 HPA 将停止调整目标（并将自身的“ScalingActive”条件设置为“false”），直到您通过手动调整重新激活它目标所需的副本数或 HPA 的最小副本数。

## What's next

##  下一步是什么

- Design documentation: [Horizontal Pod Autoscaling](https://git.k8s.io/community/contributors/design-proposals/autoscaling/horizontal-pod-autoscaler.md).
- kubectl autoscale command: [kubectl autoscale](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands/#autoscale).
- Usage example of [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/).

- 设计文档：[Horizontal Pod Autoscaling](https://git.k8s.io/community/contributors/design-proposals/autoscaling/horizontal-pod-autoscaler.md)。
- kubectl autoscale 命令：[kubectl autoscale](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands/#autoscale)。
- [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/)的使用示例。

## Feedback

##  回馈

Was this page helpful?

此页面有用吗？

Last modified    September 09, 2021 at 6:50 AM PST    : [Improvement: Remove Heapster content from HPA. (#29547) (975bd9e9b)](https://github.com/kubernetes/website/commit/975bd9e9b758153b9f41f3197f4a0b27b4fdf12a) 

上次修改时间为 2021 年 9 月 9 日太平洋标准时间上午 6:50：[改进：从 HPA 中删除 Heapster 内容。 (#29547) (975bd9e9b)](https://github.com/kubernetes/website/commit/975bd9e9b758153b9f41f3197f4a0b27b4fdf12a)

