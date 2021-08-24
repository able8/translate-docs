# Kubernetes Autoscaling in Production: Best Practices for Cluster Autoscaler, HPA and VPA

# Kubernetes 生产中的自动缩放：集群自动缩放器、HPA 和 VPA 的最佳实践

In this article we will take a deep dive into Kubernetes autoscaling tools including the cluster autoscaler, the horizontal pod autoscaler and the vertical pod autoscaler. We will also identify best practices that developers, DevOps and Kubernetes administrators should follow when configuring these tools.

在本文中，我们将深入探讨 Kubernetes 自动缩放工具，包括集群自动缩放器、水平 pod 自动缩放器和垂直 pod 自动缩放器。我们还将确定开发人员、DevOps 和 Kubernetes 管理员在配置这些工具时应遵循的最佳实践。

December 5, 2019 From: https://www.replex.io/blog/kubernetes-in-production-best-practices-for-cluster-autoscaler-hpa-and-vpa

Kubernetes is inherently scalable. It has a number of tools that allow both applications as well as the infrastructure they are hosted on to scale in and out based on demand, efficiency and a number of other metrics.

Kubernetes 本质上是可扩展的。它有许多工具，允许应用程序及其托管的基础架构根据需求、效率和许多其他指标进行扩展和扩展。

In this article we will take a deep dive into these autoscaling tools and identify best practices for working with them. This list of best practices is targeted towards developers, DevOps and Kubernetes administrators tasked with application development and delivery on Kubernetes as well as managing and operating these applications once they are in production.

在本文中，我们将深入探讨这些自动缩放工具，并确定使用它们的最佳实践。此最佳实践列表面向开发人员、DevOps 和 Kubernetes 管理员，他们的任务是在 Kubernetes 上开发和交付应用程序，以及在这些应用程序投入生产后对其进行管理和操作。


## Kubernetes Autoscaling

## Kubernetes 自动缩放

Kubernetes has three scalability tools.Two of these, the [Horizontal pod autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) (HPA) and the [Vertical pod autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) (VPA), function on the application abstraction layer. The cluster autoscaler works on the infrastructure layer.

Kubernetes 有三个可扩展性工具。 其中两个，[Horizo​​ntal pod autoscaler](https://kubernetes.io/docs/tasks/run-application/horizo​​ntal-pod-autoscale/) (HPA) 和 [Vertical pod autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) (VPA)，应用抽象层上的功能。集群自动缩放器在基础设施层工作。

In this article we will outline best practices for all three auto scaling tools. Let’s start with the cluster autoscaler.

在本文中，我们将概述所有三种自动缩放工具的最佳实践。让我们从集群自动缩放器开始。

### What is the Cluster Autoscaler?

### 什么是集群自动缩放器？

The [cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) is a Kubernetes tool that increases or decreases the size of a Kubernetes cluster (by adding or removing nodes), based on the presence of pending pods and node utilization metrics.

[cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) 是一个 Kubernetes 工具，用于增加或减少 Kubernetes 集群的大小（通过添加或删除节点），基于挂起的 pod 和节点利用率指标的存在。

The cluster autocaler
- Adds nodes to a cluster whenever it detects pending pods that could not be scheduled due to resource shortages.
- Removes nodes from a cluster, whenever the utilization of a node falls below a certain threshold defined by the cluster administrator.

集群自动调节器
- 每当检测到由于资源短缺而无法调度的挂起 Pod 时，将节点添加到集群中。
- 从集群中删除节点，只要节点的利用率低于集群管理员定义的某个阈值。

The cluster autoscaler is a great tool to ensure that the underlying cluster infrastructure is elastic and scalable and can meet the changing demands of the workloads on top.

集群自动缩放器是一个很好的工具，可确保底层集群基础架构具有弹性和可扩展性，并且可以满足顶层工作负载不断变化的需求。

Let's now move on the cluster autoscaler best practices.

现在让我们继续讨论集群自动缩放器最佳实践。

## **Cluster Autoscaler Best Practices**

## **集群自动扩缩器最佳实践**

### **Use the Correct Version of Cluster Autoscaler**

### **使用正确版本的集群自动缩放器**

Kubernetes is a fast moving platform with new releases and features being released periodically. A best practice when deploying the cluster autoscaler is to ensure that you use it with the recommended Kubernetes version. [Here](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler#releases) is a complete list of the compatibility of different cluster autoscaler versions with Kubernetes versions.

Kubernetes 是一个快速移动的平台，定期发布新版本和功能。部署集群自动缩放器时的最佳实践是确保将其与推荐的 Kubernetes 版本一起使用。 [这里](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler#releases) 是不同集群 autoscaler 版本与 Kubernetes 版本兼容性的完整列表。

**Ensure Cluster Nodes have the Same Capacity**
The cluster autoscaler only functions correctly with kubernetes node groups/instance groups that have nodes with the same capacity. One reason for this is the underlying cluster autoscaler [assumption](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-does-scale-up-work) that each individual node in the node group has the same CPU and memory capacity. Based on these assumptions, it creates template nodes for each node group and makes autoscaling decisions based on that template node. 

**确保集群节点具有相同的容量**
集群自动缩放器仅适用于具有相同容量节点的 kubernetes 节点组/实例组。一个原因是底层集群自动缩放器 [假设](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-does-scale-up-work)，每个人节点组中的节点具有相同的 CPU 和内存容量。基于这些假设，它为每个节点组创建模板节点，并根据该模板节点做出自动缩放决策。

A best practice, therefore is to ensure that the instance group being autoscaled via the cluster autoscaler has instances/nodes of the same type. For public cloud providers like AWS, this might not be optimal, since diversification and availability considerations dictate the use of multiple instance types. The cluster autoscaler does [support](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-does-scale-up-work) node groups with mixed instance types. A best practice however, is to ensure that these instance types have the same resource footprint i.e. they have the same amount of CPU and memory resources.

因此，最佳实践是确保通过集群自动缩放器自动缩放的实例组具有相同类型的实例/节点。对于像 AWS 这样的公共云提供商，这可能不是最佳选择，因为多样化和可用性考虑决定了使用多种实例类型。集群自动缩放器 [support](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-does-scale-up-work) 具有混合实例类型的节点组。然而，最佳实践是确保这些实例类型具有相同的资源占用，即它们具有相同数量的 CPU 和内存资源。

**Ensure Every Pod has Resource Requests Defined**
Since the cluster autoscaler makes scaling decisions based on the scheduling status of pods and the utilization of individual nodes, specifying resource requests is essential for it to function correctly.

**确保每个 Pod 都定义了资源请求**
由于集群自动缩放器根据 pod 的调度状态和单个节点的利用率做出缩放决策，因此指定资源请求对其正常运行至关重要。

Take cluster scale-down. The cluster autoscaler will scale down any nodes that have a utilization less than a specified threshold. [Utilization](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-the-parameters-to-ca) is calculated as the sum of requested resources divided by the capacity. Utilization calculations could be thrown off by the presence of any pods or containers without resource requests and could lead to suboptimal functioning.

采取集群缩减。集群自动缩放器将缩减利用率低于指定阈值的任何节点。 [利用率](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-the-parameters-to-ca) 计算为请求资源的总和除以容量。没有资源请求的任何 Pod 或容器的存在可能会导致利用率计算中断，并可能导致功能欠佳。

A best practice therefore is to ensure that all pods, scheduled to run in an autoscaled node group/instance group, have resource requests specified.

因此，最佳实践是确保所有计划在自动缩放的节点组/实例组中运行的 pod 都指定了资源请求。

**Specify PodDisruptionBudget for kube-system Pods**
Kube-system pods by default prevent the cluster autoscaler from [scaling down](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-to-set-pdbs-to- enable-ca-to-move-kube-system-pods) the nodes they are running on. In situations where these pods end up on different nodes, they can also prevent a cluster from scaling down.

**为 kube-system Pod 指定 PodDisruptionBudget**
Kube-system pods 默认阻止集群自动缩放器 [scaling down](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-to-set-pdbs-to-enable-ca-to-move-kube-system-pods) 他们正在运行的节点。在这些 pod 最终位于不同节点的情况下，它们还可以防止集群缩小。

To avoid situations where nodes cannot be scaled down due to the presence of system pods, a best practice is to specify a pod disruption budget for these pods. [Pod disruption budgets](https://kubernetes.io/docs/tasks/run-application/configure-pdb/) allows kubernetes administrators to avoid disruptions to critical pods and ensure that a desired number of these pods is always running.

为避免由于系统 Pod 的存在而导致节点无法缩减的情况，最佳实践是为这些 Pod 指定 Pod 中断预算。 [Pod 中断预算](https://kubernetes.io/docs/tasks/run-application/configure-pdb/) 允许 kubernetes 管理员避免对关键 Pod 造成中断，并确保所需数量的这些 Pod 始终运行。

While specifying a disruption budget for system pods it is important to consider the number of replicas of these pods that are provisioned by default.

在为系统 Pod 指定中断预算时，重要的是要考虑默认配置的这些 Pod 的副本数量。

Kube-dns is the only system pod that has multiple running replicas by default. Most other system pods run as single instance pods and restarting them could result in disruptions to the cluster.

Kube-dns 是唯一一个默认有多个运行副本的系统 pod。大多数其他系统 Pod 作为单实例 Pod 运行，重新启动它们可能会导致集群中断。

A best practice in this context is to avoid building in a disruption budget for single instance pods like the metrics-server.

这种情况下的最佳实践是避免为指标服务器等单实例 Pod 建立中断预算。

**Specify PodDisruptionBudget for Application Pods**
In addition to specifying a pod disruption budget for system pods, another best practice is to also specify a pod disruption budget for [application pods](https://kubernetes.io/docs/tasks/run-application/configure-pdb/# protecting-an-application-with-a-poddisruptionbudget). This will ensure that the cluster autoscaler does not scale down pod replicas beyond a certain minimum number and will protect critical applications from disruptions and ensure high availability.

**为应用程序 Pod 指定 PodDisruptionBudget**
除了为系统 pod 指定 pod 中断预算之外，另一个最佳实践是还为 application pods 指定 pod 中断预算。这将确保集群自动缩放器不会将 pod 副本缩减超过某个最小数量，并将保护关键应用程序免受中断并确保高可用性。

Pod disruption budgets can be specified using the `.spec.minAvailable` and `.spec.maxUnavailable` fields. `.spec.minAvailable` specifies the number of pods that must be available after the eviction, as an absolute number or a percentage. Similarly `.spec.maxUnavailable` sets out the maximum number of pods that can be unavailable after the eviction expressed either as an absolute number or a percentage.

Pod 中断预算可以使用 `.spec.minAvailable` 和 `.spec.maxUnavailable` 字段指定。 `.spec.minAvailable` 指定驱逐后必须可用的 pod 数量，以绝对数量或百分比表示。类似地，`.spec.maxUnavailable` 列出了驱逐后可以不可用的最大 pod 数量，以绝对数量或百分比表示。

**Avoid using the Cluster Autoscaler with more than 1000 Node Clusters**
For the cluster autoscaler to remain responsive it is important to ensure that the cluster does not exceed a certain size. The [official](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/proposals/scalability_tests.md#ca-scales-to-1000-nodes) scalability and responsiveness service level for the cluster autoscaler is set at 1000 nodes with each node running 30 pods. Here is a complete writeup of the scale up and scale down results using a test setup with a 1000 node cluster.

**避免对超过 1000 个节点集群使用集群自动缩放器**
为了让集群自动缩放器保持响应，确保集群不超过特定大小非常重要。 [官方](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/proposals/scalability_tests.md#ca-scales-to-1000-nodes) 集群自动缩放器的可扩展性和响应性服务级别设置为 1000 个节点，每个节点运行 30 个 Pod。这是使用具有 1000 个节点集群的测试设置放大和缩小结果的完整记录。

A best practice therefore is to avoid cluster sprawl and ensure that the cluster footprint does not exceed the specified scalability limit.

因此，最佳实践是避免集群蔓延并确保集群足迹不超过指定的可扩展性限制。

**Ensure Resource Availability for the Cluster Autoscaler Pod** 
**确保集群 Autoscaler Pod 的资源可用性**
For larger clusters it is important to ensure resource availability for the cluster autoscaler. A best practice in this context is to set resource requests of the cluster autoscaler pod to a minimum of [1 CPU](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/proposals/scalability_tests.md #ca-scales-to-1000-nodes).

对于较大的集群，确保集群自动缩放器的资源可用性非常重要。这种情况下的最佳实践是将集群自动缩放器 pod 的资源请求设置为至少 1 CPU。

It is also important to ensure that the node the cluster autoscaler pod is running on has enough resources available to support it. Running the cluster autoscaler pod on a node with resource pressure, could lead to degraded performance or the cluster autoscaler becoming non responsive.

确保运行集群自动缩放器 pod 的节点有足够的可用资源来支持它也很重要。在具有资源压力的节点上运行集群自动缩放器 pod，可能会导致性能下降或集群自动缩放器变得无响应。

**Ensure Resource Requests are Close to Actual Usage**
As mentioned before the cluster autoscaler makes scaling decisions based on the presence of pending pods and the utilization of individual nodes. Node utilization is calculated as the sum of requested resources of all pods divided by the capacity.

**确保资源请求接近实际使用**
如前所述，集群自动缩放器根据待处理 pod 的存在和各个节点的利用率做出扩展决策。节点利用率计算为所有 Pod 的请求资源总和除以容量。

However, most developers tend to over provision [resource requests](http://www.replex.io/blog/5-ways-to-manage-your-kubernetes-resource-usage). This can at times lead to situations where pods are not utilizing requested resources efficiently, leading to a lower overall node utilization. However since the total resource requests are high the cluster autoscaler calculates a higher utilization level for the node and might not scale it down.

但是，大多数开发人员倾向于过度提供 [资源请求](http://www.replex.io/blog/5-ways-to-manage-your-kubernetes-resource-usage)。这有时会导致 Pod 没有有效利用请求的资源的情况，从而导致整体节点利用率较低。但是，由于总资源请求很高，集群自动缩放程序会为节点计算更高的利用率级别，并且可能不会将其缩小。

A best practice therefore is to ensure that pod's requested resources are comparable to the actual [resource usage](http://www.replex.io/blog/kubernetes-in-production-questions-you-need-to-be-asking )/consumption. Using the virtual pod autoscaler (VPA), is a good starting point. These decisions can also be based on historical resource usage and consumption of pods.

因此，最佳实践是确保 pod 请求的资源与实际的[资源使用情况](http://www.replex.io/blog/kubernetes-in-production-questions-you-need-to-be-asking)/消费。使用虚拟 Pod 自动缩放器 (VPA) 是一个很好的起点。这些决策也可以基于 Pod 的历史资源使用和消耗。

**Over-provision Cluster to Ensure head room for Critical pods**
The cluster autoscaler has a [service level objective (SLO)](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-the-service-level-objectives- for-cluster-autoscaler) of 30 seconds latency between the time a pod is marked as unschedulable to the time that it requests a scale-up to the cloud provider. This latency benchmark is for smaller clusters of less than 100 nodes. For larger clusters of up to a 1000 nodes this latency is expected to be around the 60 second mark.

**过度配置集群以确保关键 Pod 的空间**
集群自动缩放器有一个服务级别目标 (SLO)从 pod 被标记为不可调度的时间到它请求扩展到云提供商的时间之间有 30 秒的延迟。此延迟基准适用于少于 100 个节点的较小集群。对于多达 1000 个节点的较大集群，此延迟预计将在 60 秒左右。

The actual time that it takes for the pod to be scheduled as a result of the scale up request and a new node being provisioned, depends on the cloud provider. This could very well mean a delay of several minutes.

由于扩展请求和供应新节点而安排 pod 所需的实际时间取决于云提供商。这很可能意味着几分钟的延迟。

To avoid this delay and ensure that pods spend as little time as possible in unschedulable state, a best practice is to over provision the cluster. This can be accomplished using a deployment running pause pods.

为了避免这种延迟并确保 pod 尽可能少地处于不可调度状态，最佳实践是过度配置集群。这可以使用运行暂停 Pod 的部署来完成。

Pause pods are dummy pods that are spun up exclusively to reserve space for other higher priority pods. Since pause pods are assigned a very low [priority](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/), the kubernetes scheduler will remove them to make space for unscheduled pods with a higher priority . This essentially means that critical pods do not have to wait for a new node to be provisioned by the cloud provider and can be quickly scheduled on the already existing nodes, replacing the pause pods.

暂停 pod 是虚拟 pod，专门旋转起来为其他更高优先级的 pod 保留空间。由于暂停 pod 被分配了非常低的 [优先级](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/)，kubernetes 调度程序将删除它们，为具有更高优先级的未调度 pod 腾出空间.这实质上意味着关键 Pod 不必等待云提供商提供新节点，并且可以在现有节点上快速调度，取代暂停 Pod。

Once the pause pods re-spawn they become unschedulable resulting in the cluster scaling up. Cluster over provisioned head room can be controlled by specifying the size of the pause pods.

一旦暂停 Pod 重新生成，它们就会变得不可调度，从而导致集群扩展。可以通过指定暂停 Pod 的大小来控制集群过度配置的净空。

### To recap here are the recommended best practices for the cluster autoscaler on Kubernetes:

### 在这里回顾一下 Kubernetes 上集群自动缩放器的推荐最佳实践：

- Use the correct version of Cluster autoscaler
- Ensure cluster nodes have the same capacity:
- Ensure every pod has resource requests defined
- Specify PodDisruptionBudget for kube-system pods
- Specify PodDisruptionBudget for application pods
- Avoid using the Cluster autoscaler with more than 1000 node clusters
- Ensure resource availability for the cluster autoscaler pod
- Ensure resource requests are close to actual usage
- Over-provision cluster to ensure head room for critical pods

- 使用正确版本的 Cluster autoscaler
- 确保集群节点具有相同的容量：
- 确保每个 Pod 都定义了资源请求
- 为 kube-system pod 指定 PodDisruptionBudget
- 为应用程序 Pod 指定 PodDisruptionBudget
- 避免对超过 1000 个节点集群使用集群自动缩放器
- 确保集群自动缩放器 pod 的资源可用性
- 确保资源请求接近实际使用情况
- 过度配置集群以确保关键 Pod 的空间

Let us now move on to the horizontal pod autoscaler (HPA).

现在让我们继续讨论水平 Pod 自动缩放器 (HPA)。

## **HorizontalPodAutoscaler (HPA) Best Practices** 
## **Horizo​​ntalPodAutoscaler (HPA) 最佳实践**
HPA scales the number of pods in a replication controller, deployment, replica set or stateful set based on CPU utilization. HPA can also be configured to make scaling decisions based on [custom](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-custom-metrics) or [external](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#autoscaling-on-metrics-not-related-to-kubernetes-objects) metrics.

HPA 根据 CPU 利用率扩展复制控制器、部署、副本集或有状态集中的 pod 数量。 HPA 还可以配置为基于自定义或外部指标做出扩展决策。

HPA is a great tool to ensure that critical applications are elastic and can scale out to meet increasing demand as well scale down to ensure optimal [resource usage](http://www.replex.io/blog/7-things-you- can-do-today-to-reduce-aws-kubernetes-costs).

HPA 是一个很好的工具，可以确保关键应用程序具有弹性，可以横向扩展以满足不断增长的需求，也可以缩减规模以确保最佳资源使用。

**Ensure all Pods have Resource Requests Configured**
HPA makes scaling decisions based on the observed CPU utilisation values of pods that are part of a Kubernetes controller. [Utilisation](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/horizontal-pod-autoscaler.md#horizontalpodautoscaler-object) values are calculated as a percentage of the resource requests of individual pods. Missing resource request values for some containers might throw off the utilisation calculations of the HPA controller leading to suboptimal operation and scaling decisions.

**确保所有 Pod 都配置了资源请求**
HPA 根据观察到的作为 Kubernetes 控制器一部分的 Pod 的 CPU 利用率值做出扩展决策。 [利用率](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/horizo​​ntal-pod-autoscaler.md#horizo​​ntalpodautoscaler-object) 值计算为资源请求的百分比个别豆荚。某些容器缺少资源请求值可能会影响 HPA 控制器的利用率计算，从而导致次优操作和扩展决策。

A best practice therefore is to ensure that resource request values are configured for all containers of each individual pod, that is a part of the Kubernetes controller being scaled using HPA.

因此，最佳实践是确保为每个单独 pod 的所有容器配置资源请求值，这是使用 HPA 扩展的 Kubernetes 控制器的一部分。

**Install metrics-server**
HPA makes scaling decisions based on per-pod resource metrics retrieved from the resource metrics API (metrics.k8s.io). The metrics.k8s.io API is provided by the metrics-server. A best practice therefore is to [launch](https://kubernetes.io/docs/tasks/debug-application-cluster/resource-metrics-pipeline/#metrics-server) metrics-server in your Kubernetes cluster as a cluster add -on.

**安装指标服务器**
HPA 根据从资源指标 API (metrics.k8s.io) 检索到的每个 Pod 资源指标做出扩展决策。 metrics.k8s.io API 由 metrics-server 提供。因此，最佳实践是在您的 Kubernetes 集群中 [启动](https://kubernetes.io/docs/tasks/debug-application-cluster/resource-metrics-pipeline/#metrics-server) metrics-server 作为集群添加-在。

In addition to this, another best practice is to set `--horizontal-pod-autoscaler-use-rest-clients` to `true ` or unset. This is important since setting this flag to `false` will revert to Heapster which is deprecated as of Kubernetes 1.11.

除此之外，另一个最佳实践是将 `--horizo​​ntal-pod-autoscaler-use-rest-clients` 设置为 `true ` 或未设置。这很重要，因为将此标志设置为 `false` 将恢复为 Heapster，自 Kubernetes 1.11 起已弃用。

### Configure Custom or External Metrics

### 配置自定义或外部指标

The HPA can also make scaling decisions based on custom or external metrics. There are two types of custom metrics supported: pod and object metrics. [Pod metrics](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#autoscaling-on-multiple-metrics-and-custom-metricshttps://kubernetes.io/ docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/%23autoscaling-on-multiple-metrics-and-custom-metrics) are averaged across all pods and as such only support `target ` type of `AverageValue`. Object metrics can describe any other object in the same namespace and support `target` types of both `Value` and `AverageValue`.

HPA 还可以根据自定义或外部指标做出扩展决策。支持两种类型的自定义指标：pod 和对象指标。 [Pod 指标](https://kubernetes.io/docs/tasks/run-application/horizo​​ntal-pod-autoscale-walkthrough/#autoscaling-on-multiple-metrics-and-custom-metricshttps://kubernetes.io/ docs/tasks/run-application/horizo​​ntal-pod-autoscale-walkthrough/%23autoscaling-on-multiple-metrics-and-custom-metrics) 在所有 pod 中取平均值，因此仅支持 `target` 类型的 `AverageValue`。对象度量可以描述同一命名空间中的任何其他对象，并支持 `Value` 和 `AverageValue` 的 `target` 类型。

A best practice when configuring custom metrics is to ensure that the correct `target` type is used for pod and object metrics.

配置自定义指标时的最佳实践是确保为 pod 和对象指标使用正确的“目标”类型。

External metrics allow HPA to autoscale applications based on metrics provided by third party monitoring systems. External metrics support `target` types of both `Value` and `AverageValue`.

外部指标允许 HPA 根据第三方监控系统提供的指标自动扩展应用程序。外部指标支持 `Value` 和 `AverageValue` 的 `target` 类型。

### Prefer Custom Metrics over External Metrics whenever Possible

### 尽可能使用自定义指标而不是外部指标

A best practice when deciding between custom and external metrics (when such a choice is possible) is to [prefer](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#autoscaling -on-metrics-not-related-to-kubernetes-objects) custom metrics. One reason for this is the fact that the external metrics API takes a lot more effort to secure as compared to custom metrics API and could potentially allow access to all metrics.

在自定义指标和外部指标之间做出决定（当可以进行此类选择时）的最佳实践是 [首选](https://kubernetes.io/docs/tasks/run-application/horizo​​ntal-pod-autoscale-walkthrough/#autoscaling -on-metrics-not-related-to-kubernetes-objects) 自定义指标。其原因之一是，与自定义指标 API 相比，外部指标 API 需要付出更多努力来确保安全，并且可能允许访问所有指标。

### Configure Cooldown Period

### 配置冷却时间

The dynamic nature of the metrics being evaluated by the HPA may at times lead to scaling events in quick succession without a period between those scaling events. This leads to [thrashing](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-cooldown-delay) where the number of replicas fluctuates frequently and is not desirable.

HPA 评估的指标的动态特性有时可能会导致快速连续的扩展事件，而这些扩展事件之间没有时间间隔。这会导致 [thrashing](https://kubernetes.io/docs/tasks/run-application/horizo​​ntal-pod-autoscale/#support-for-cooldown-delay)，其中副本数量频繁波动且不理想。

To get around this and specify a cool down period a best practice is to configure the `--horizontal-pod-autoscaler-downscale-stabilization ` flag passed to the kube-controller-manager. This flag has a default value of 5 minutes and specifies the duration HPA waits after a downscale event before initiating another downscale operation.

要解决此问题并指定冷却时间，最佳做法是配置传递给 kube-controller-manager 的 `--horizo​​ntal-pod-autoscaler-downscale-stabilization ` 标志。此标志的默认值为 5 分钟，并指定 HPA 在缩减事件之后启动另一个缩减操作之前等待的持续时间。

Kubernetes admins should also take into account the unique requirements of their applications when deciding on an optimal value for this duration. 
Kubernetes 管理员在决定此持续时间的最佳值时还应考虑其应用程序的独特要求。

By default the HPA tolerates a 10% change in the desired to actual metrics ratio before scaling. Depending on application requirements, this value can be changed by configuring the `horizontal-pod-autoscaler-tolerance` flag. Other [configurable flags](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#algorithm-details) include -- `horizontal-pod-autoscaler-cpu-initialization-period` duration ,  `horizontal-pod-autoscaler-initial-readiness-delay` duration and `horizontal-pod-autoscaler-sync-period` duration. All of these can be configured based on unique cluster or application requirements.

默认情况下，HPA 允许在缩放前所需的实际指标比率发生 10% 的变化。根据应用程序要求，可以通过配置 `horizo​​ntal-pod-autoscaler-tolerance` 标志来更改此值。其他 [可配置标志](https://kubernetes.io/docs/tasks/run-application/horizo​​ntal-pod-autoscale/#algorithm-details) 包括 -- `horizo​​ntal-pod-autoscaler-cpu-initialization-period` 持续时间, `horizo​​ntal-pod-autoscaler-initial-readiness-delay` 持续时间和 `horizo​​ntal-pod-autoscaler-sync-period` 持续时间。所有这些都可以根据独特的集群或应用程序要求进行配置。

### To recap here are the Horizontal Pod Autoscaler (HPA) best practices.

### 在这里回顾一下 Horizo​​ntal Pod Autoscaler (HPA) 最佳实践。

- Ensure all pods have resource requests specified
- Install metrics-server
- Configure custom or external metrics
- Prefer Custom metrics over external metrics
- Configure cool-down period

- 确保所有 Pod 都指定了资源请求
- 安装指标服务器
- 配置自定义或外部指标
- 比外部指标更喜欢自定义指标
- 配置冷却时间

## Vertical Pod Autoscaler (VPA) Best Practices

## 垂直 Pod Autoscaler (VPA) 最佳实践

Next we will review best practices for the Vertical Pod Autoscaler (VPA). [VPA](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) automatically sets the resource request and limit values of containers based on usage. VPA aims to reduce the maintenance overhead of configuring resource requests and limits for containers and improve the utilization of cluster resources.

接下来，我们将回顾 Vertical Pod Autoscaler (VPA) 的最佳实践。 [VPA](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)根据使用情况自动设置容器的资源请求和限制值。 VPA旨在减少为容器配置资源请求和限制的维护开销，提高集群资源的利用率。

The VerticalPodAutoscaler can:

VerticalPodAutoscaler 可以：

- Reduce the request value for containers whose resource usage is consistently lower than the requested amount.
- Increase request values for containers that consistently use a high percentage of resources requested.
- Automatically set resource limit values based on limit to request ratios specified as part of the container template.

- 降低资源使用量始终低于请求量的容器的请求值。
- 为始终使用大量请求资源的容器增加请求值。
- 根据作为容器模板一部分指定的请求比率限制自动设置资源限制值。

### Use the Correct Kubernetes Version

### 使用正确的 Kubernetes 版本

Version 0.4 and later of the VerticalPodAutoscaler requires custom resource definition capabilities and can therefore [NOT](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#installation) be used with Kubernetes versions older than 1.11. For earlier Kubernetes versions it is recommended to use version 0.3 of the VerticalPodAutoscaler.

VerticalPodAutoscaler 0.4 及更高版本需要自定义资源定义功能，因此可以 [NOT](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#installation) 与 Kubernetes 版本早于1.11.对于较早的 Kubernetes 版本，建议使用 VerticalPodAutoscaler 的 0.3 版。

## Install metrics-server and Prometheus

##安装metrics-server和Prometheus

VPA makes scaling decisions based on usage and utilization metrics from both Prometheus and metrics-server.

VPA 根据来自 Prometheus 和 metrics-server 的使用和利用率指标做出扩展决策。

[Recommender](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/pkg/recommender/README.md) is the main component of the VerticalPodAutoscaler and is responsible for computing recommended resources and generating a recommendation model. For [running pods](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/vertical-pod-autoscaler.md#recommender), the recommender component receives real-time usage and utilization metrics from the metrics-server via the metrics API and makes scaling decisions based on them. A best practice therefore is to ensure that the metrics-server is running in your Kubernetes cluster.

[Recommender](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/pkg/recommender/README.md)是VerticalPodAutoscaler的主要组件，负责计算推荐资源并生成推荐模型。对于 [running pods](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/vertical-pod-autoscaler.md#recommender)，推荐组件接收实时使用和通过指标 API 来自指标服务器的利用率指标，并根据它们做出扩展决策。因此，最佳实践是确保指标服务器在您的 Kubernetes 集群中运行。

Unlike HPA however, VPA also requires prometheus. The [history storage](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/vertical-pod-autoscaler.md#history-storage) component of VPA, consumes utilization signals and OOM events and stores them persistently and is backed up by Prometheus. On startup the recommender fetches this data from the history storage and keeps it in memory.

然而，与 HPA 不同的是，VPA 还需要 prometheus。 VPA 的 [history storage](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/vertical-pod-autoscaler.md#history-storage) 组件，消耗利用率信号和OOM 事件并持久存储它们，并由 Prometheus 备份。在启动时，推荐器从历史存储中获取这些数据并将其保存在内存中。

For the recommender to pull in this [historical](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/pkg/recommender/README.md#running) data, a best practice is to install Prometheus in your cluster and configure it to receive metrics from cadvisor. Also ensure that metrics from cAdvisor have the label `job=kubernetes-cadvisor`.

对于推荐者拉入这个[历史](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/pkg/recommender/README.md#running)数据，最佳实践是在集群中安装 Prometheus 并将其配置为从 cadvisor 接收指标。还要确保来自 cAdvisor 的指标具有标签“job=kubernetes-cadvisor”。

[Another best practice](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/FAQ.md#how-can-i-use-prometheus-as-a-history-provider- for-the-vpa-recommender) is to set the `--storage=prometheus` and the `--prometheus-address=<your-prometheus-address>` flags in the VerticalPodAutoscaler deployment:

[另一个最佳实践](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/FAQ.md#how-can-i-use-prometheus-as-a-history-provider- for-the-vpa-recommender) 是在 VerticalPodAutoscaler 部署中设置 `--storage=prometheus` 和 `--prometheus-address=<your-prometheus-address>` 标志：

This is what the spec looks like:

这是规范的样子：

```
spec:
 containers:
 - args:
 - --v=4
 - --storage=prometheus
 - --prometheus-address=http://prometheus.default.svc.cluster.local:9090

```


Also make sure you update the `--prometheus-address` flag with the name of the actual namespace that Prometheus is running in.

还要确保使用 Prometheus 运行的实际命名空间的名称更新 `--prometheus-address` 标志。

### Avoid using HPA and VPA in tandem

### 避免同时使用 HPA 和 VPA

HPA and VPA are currently incompatible and a best practice is to avoid using both together for the same set of pods. VPA can however be used with HPA that is configured to use either external or custom metrics.

HPA 和 VPA 目前不兼容，最佳实践是避免将两者一起用于同一组 pod。但是，VPA 可以与配置为使用外部或自定义指标的 HPA 一起使用。

### Use VPA together with Cluster autoscaler 
### 将 VPA 与集群自动缩放器一起使用
A best practice when configuring VPA is to use it in combination with the cluster autoscaler. The recommender component of VPA might at times recommend resource request values that exceed the available resources. This leads to resource pressure and might result in some pods going into pending state. Having the cluster autoscaler running mitigates this behaviour since it spins up new nodes in response to pending pods.

配置 VPA 时的最佳实践是将其与集群自动缩放器结合使用。 VPA 的推荐组件有时可能会推荐超过可用资源的资源请求值。这会导致资源压力，并可能导致某些 Pod 进入挂起状态。运行集群自动缩放器可以缓解这种行为，因为它会启动新节点以响应挂起的 pod。

### To recap here are the recommended best practices for VPA:

### 在这里回顾一下推荐的 VPA 最佳实践：

- Use the Correct Kubernetes Version
- Install metrics-server and Prometheus
- Avoid using HPA and VPA in tandem
- Use VPA together with Cluster autoscaler 
- 使用正确的 Kubernetes 版本
- 安装 metrics-server 和 Prometheus
- 避免同时使用 HPA 和 VPA
- 将 VPA 与 Cluster autoscaler 一起使用