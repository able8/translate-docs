# Zero Downtime Server Updates For Your Kubernetes Cluster

# Kubernetes 集群的零停机服务器更新

> From: https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33

Rolling Update for your Kubernetes Cluster

Kubernetes 集群的滚动更新

At some point during the lifetime of your Kubernetes cluster, you will  need to perform maintenance on the underlying nodes. This may include  package updates, kernel upgrades, or deploying new VM images. This is  considered a ["Voluntary Disruption"](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#voluntary-and-involuntary-disruptions) in Kubernetes.

在 Kubernetes 集群的生命周期中的某个时刻，您需要对底层节点执行维护。这可能包括软件包更新、内核升级或部署新的 VM 镜像。这在 Kubernetes 中被视为 [“自愿中断”](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#voluntary-and-involved-disruptions)。

This is part of a 4 part blog series:

1. This post
2. [Gracefully shutting down Pods](https://medium.com/p/328aecec90d)
3. [Delaying Shutdown to Wait for Pod Deletion Propagation](https://medium.com/p/445f779a8304)
4. [Avoiding Outages with PodDisruptionBudgets](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085)

这是 4 部分博客系列的一部分：

1. 这个帖子
2. [优雅地关闭 Pods](https://medium.com/p/328aecec90d)
3. [延迟关机等待Pod删除传播](https://medium.com/p/445f779a8304)
4. [使用 PodDisruptionBudgets 避免中断](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085)

In this series, we will be covering all the tools that Kubernetes provides to achieve a zero downtime update for the underlying worker nodes in  your cluster.

在本系列中，我们将介绍 Kubernetes 提供的所有工具，以实现集群中底层工作节点的零停机更新。

## Stating the problem

## 说明问题

We will start with a naive approach, identify challenges and potential  risks of the approach, and incrementally build up to solve each problem  that we identify throughout the series. We will finish with a config  that leverages lifecycle hooks, readiness probes, and Pod disruption  budgets to achieve our zero downtime rollout.

我们将从一个简单的方法开始，确定该方法的挑战和潜在风险，并逐步改进我们在整个系列中确定的每个问题。我们将完成一个配置，该配置利用生命周期钩子、准备就绪探测和 Pod 中断预算来实现我们的零停机部署。

To start our journey, let’s look at a concrete example. Suppose we have a  two node Kubernetes cluster running an application with two Pods backing a `Service` resource:

为了开始我们的旅程，让我们看一个具体的例子。假设我们有一个两节点的 Kubernetes 集群运行一个应用程序，其中两个 Pod 支持一个“服务”资源：

![img](https://miro.medium.com/max/1400/1*WExl8VYi2FpkhzOTBcy16Q.png)

Our starting point with 2 Nginx pods and a Service running on our 2 node Kubernetes cluster.

我们的起点是 2 个 Nginx pod 和一个在我们的 2 节点 Kubernetes 集群上运行的服务。

We want to upgrade the kernel version of the two underlying worker nodes  in our cluster. How will we do this? A naive approach would be to launch new nodes with the updated config and then shutdown the old nodes once  the new nodes were launched. While this works, there are a few issues  with this approach:

我们要升级集群中两个底层工作节点的内核版本。我们将如何做到这一点？一种天真的方法是使用更新的配置启动新节点，然后在新节点启动后关闭旧节点。虽然这是有效的，但这种方法存在一些问题：

- When you shutdown the old nodes, you would be taking down the running pods  with it. What if the pods need to clean up for a graceful shutdown? The  underlying VM technology might not wait for the clean up process.
 - What if you shutdown all the nodes at the same time? You could have a brief  outage while the pods are relaunched into the new nodes.

- 当您关闭旧节点时，您将同时关闭正在运行的 Pod。如果 Pod 需要清理以正常关闭怎么办？底层 VM 技术可能不会等待清理过程。
- 如果您同时关闭所有节点会怎样？当 Pod 重新启动到新节点时，您可能会出现短暂中断。

What we want is a way to gracefully migrate the pods off of the old nodes to ensure that none of our workloads are running while we make changes to  the node. Or if we are doing a full replacement of the cluster as in the example (e.g replacing VM images), we want to move the workloads off of the old nodes to the new nodes. In both cases, we want to prevent new  pods from being scheduled on the old node and then evict all the running pods off of it. We can use the `kubectl drain` command to achieve this.

我们想要的是一种从旧节点优雅地迁移 pod 的方法，以确保在我们对节点进行更改时没有任何工作负载在运行。或者，如果我们像示例中那样完全替换集群（例如替换 VM 映像），我们希望将工作负载从旧节点转移到新节点。在这两种情况下，我们都希望防止在旧节点上调度新 pod，然后将所有正在运行的 pod 驱逐出它。我们可以使用 `kubectl drain` 命令来实现这一点。

## Rescheduling Pods off of a node

## 从节点重新调度 Pod

The drain operation achieves the goal of rescheduling all the Pods off of  the node. During the drain operation, the node is marked as  unschedulable (the `NoSchedule`  taint). This prevents new pods from being scheduled on the node. Afterwards, the drain operation starts evicting the pods from the node,  shutting down the containers that are currently running on the node by  sending the `TERM` signal to the underlying containers of the pods.

Drain 操作实现了将所有 Pod 重新调度离开节点的目标。在 Drain 操作期间，节点被标记为不可调度（“NoSchedule”污点）。这可以防止在节点上调度新的 pod。之后，drain 操作开始从节点驱逐 Pod，通过向 Pod 的底层容器发送“TERM”信号来关闭当前在节点上运行的容器。

Although `kubectl drain` will gracefully handle pod eviction, there are still two factors that  could cause service disruption during the drain operation:

尽管 `kubectl drain` 会优雅地处理 pod eviction，但仍有两个因素可能在 Drain 操作期间导致服务中断：

- Your service application needs to be able to gracefully handle the `TERM` signal. When a Pod is evicted, Kubernetes will send the `TERM` signal to the container, and then will wait for a configurable amount  of time for the container to shutdown after giving the signal before  forcefully terminating it. However, if your containers do not handle the signal gracefully, you could still shutdown the pods uncleanly if it is in the middle of doing work (e.g committing a database transaction).
 - You lose all the pods servicing your application. Your service could  experience downtime while the new containers are being started on the  new nodes, or, if you did not deploy your pods with controllers, they  could end up never restarting.

- 您的服务应用程序需要能够优雅地处理“TERM”信号。当 Pod 被驱逐时，Kubernetes 会向容器发送“TERM”信号，然后在发出信号后等待可配置的时间让容器关闭，然后再强行终止它。但是，如果您的容器不能正常处理信号，如果 pod 正在执行工作（例如提交数据库事务），您仍然可以不干净地关闭 pod。
- 您丢失了为您的应用程序提供服务的所有 pod。在新节点上启动新容器时，您的服务可能会停机，或者，如果您没有使用控制器部署 pod，它们最终可能永远不会重新启动。

## Avoiding outages

## 避免中断

To minimize downtime from a voluntary disruption like draining a node,  Kubernetes provides the following disruption handling features:

- [Graceful termination](https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods)
 - [Lifecycle hooks](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/)
 - [PodDisruptionBudgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work) 

为了最大限度地减少因节点耗尽等自愿中断造成的停机时间，Kubernetes 提供了以下中断处理功能：

- [优雅终止](https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods)

- [生命周期钩子](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/)

- [PodDisruptionBudgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work)

In the rest of the series, we will use these features of Kubernetes to  mitigate service disruption from an eviction event. To make it easier to follow along, we will use our example above with the following resource config:

在本系列的其余部分，我们将使用 Kubernetes 的这些功能来减轻驱逐事件造成的服务中断。为了更容易理解，我们将使用上面的示例和以下资源配置：

```
 ---
 apiVersion: apps/v1
 kind: Deployment
 metadata:
   name: nginx-deployment
   labels:
     app: nginx
 spec:
   replicas: 2
   selector:
     matchLabels:
       app: nginx
   template:
     metadata:
       labels:
         app: nginx
     spec:
       containers:
       - name: nginx
         image: nginx:1.15
         ports:
         - containerPort: 80
 ---
 kind: Service
 apiVersion: v1
 metadata:
   name: nginx-service
 spec:
   selector:
     app: nginx
   ports:
   - protocol: TCP
     targetPort: 80
     port: 80
```


This config is a minimal example of a `Deployment` resource that manages multiple nginx pods (in our case two). This  resource will work towards maintaining two nginx pods in the cluster. Additionally, the config will provision a `Service` resource that can be used to access the nginx pods within the cluster.

此配置是管理多个 nginx pod（在我们的示例中为两个）的“部署”资源的最小示例。该资源将致力于在集群中维护两个 nginx pod。此外，配置将提供一个“服务”资源，可用于访问集群内的 nginx pod。

We will incrementally add to this throughout this series to build up to a  final config that implements all of the features Kubernetes provides to  minimize downtime during a maintenance operation. Here is our roadmap:

- [Gracefully shutting down Pods](https://medium.com/p/328aecec90d)

 - [Delaying Shutdown to Wait for Pod Deletion Propagation](https://medium.com/p/445f779a8304)

 - [Avoiding outages with PodDisruptionBudgets](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085

我们将在本系列中逐步添加，以构建最终配置，该配置实现了 Kubernetes 提供的所有功能，以最大限度地减少维护操作期间的停机时间。这是我们的路线图：

- [优雅地关闭 Pod](https://medium.com/p/328aecec90d)

- [延迟关闭以等待 Pod 删除传播](https://medium.com/p/445f779a8304)

- [使用 PodDisruptionBudgets 避免中断](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085)

Head on over to [the next post](https://medium.com/p/328aecec90d) to learn how you can leverage lifecycle hooks to gracefully shutdown your Pods.

前往 [下一篇文章](https://medium.com/p/328aecec90d) 了解如何利用生命周期钩子优雅地关闭 Pod。

*To get a fully implemented and tested version of zero downtime Kubernetes cluster updates on AWS and more, check out* [*Gruntwork.io*](http://gruntwork.io)*.*

*要在 AWS 等上获得完全实施和测试的零停机 Kubernetes 集群更新版本，请查看* [*Gruntwork.io*](http://gruntwork.io)*.*

[Gruntwork](https://blog.gruntwork.io/?source=post_sidebar--------------------------post_sidebar----- ------)


Thanks to Yevgeniy Brikman.

感谢 Yevgeniy Brikman。

- [Kubernetes](https://blog.gruntwork.io/tagged/kubernetes)

The Gruntwork Blog 
Gruntwork 博客# Zero Downtime Server Updates For Your Kubernetes Cluster

# Kubernetes 集群的零停机服务器更新

Rolling Update for your Kubernetes Cluster

Kubernetes 集群的滚动更新

At some point during the lifetime of your Kubernetes cluster, you will  need to perform maintenance on the underlying nodes. This may include  package updates, kernel upgrades, or deploying new VM images. This is  considered a ["Voluntary Disruption"](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#voluntary-and-involuntary-disruptions) in Kubernetes.

在 Kubernetes 集群的生命周期中的某个时刻，您需要对底层节点执行维护。这可能包括软件包更新、内核升级或部署新的 VM 镜像。这在 Kubernetes 中被视为 [“自愿中断”](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#voluntary-and-involved-disruptions)。

This is part of a 4 part blog series:

1. This post
2. [Gracefully shutting down Pods](https://medium.com/p/328aecec90d)
3. [Delaying Shutdown to Wait for Pod Deletion Propagation](https://medium.com/p/445f779a8304)
4. [Avoiding Outages with PodDisruptionBudgets](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085)

这是 4 部分博客系列的一部分：
1. 这个帖子
2. [优雅地关闭 Pods](https://medium.com/p/328aecec90d)
3. [延迟关机等待Pod删除传播](https://medium.com/p/445f779a8304)
4. [使用 PodDisruptionBudgets 避免中断](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085)

In this series, we will be covering all the tools that Kubernetes provides to achieve a zero downtime update for the underlying worker nodes in  your cluster.

在本系列中，我们将介绍 Kubernetes 提供的所有工具，以实现集群中底层工作节点的零停机更新。

## Stating the problem

## 说明问题

We will start with a naive approach, identify challenges and potential  risks of the approach, and incrementally build up to solve each problem  that we identify throughout the series. We will finish with a config  that leverages lifecycle hooks, readiness probes, and Pod disruption  budgets to achieve our zero downtime rollout.

我们将从一个简单的方法开始，确定该方法的挑战和潜在风险，并逐步改进我们在整个系列中确定的每个问题。我们将完成一个配置，该配置利用生命周期钩子、准备就绪探测和 Pod 中断预算来实现我们的零停机部署。

To start our journey, let’s look at a concrete example. Suppose we have a  two node Kubernetes cluster running an application with two Pods backing a `Service` resource:

为了开始我们的旅程，让我们看一个具体的例子。假设我们有一个两节点的 Kubernetes 集群运行一个应用程序，其中两个 Pod 支持一个“服务”资源：

![img](https://miro.medium.com/max/1400/1*WExl8VYi2FpkhzOTBcy16Q.png)

Our starting point with 2 Nginx pods and a Service running on our 2 node Kubernetes cluster.

我们的起点是 2 个 Nginx pod 和一个在我们的 2 节点 Kubernetes 集群上运行的服务。

We want to upgrade the kernel version of the two underlying worker nodes  in our cluster. How will we do this? A naive approach would be to launch new nodes with the updated config and then shutdown the old nodes once  the new nodes were launched. While this works, there are a few issues  with this approach:

我们要升级集群中两个底层工作节点的内核版本。我们将如何做到这一点？一种天真的方法是使用更新的配置启动新节点，然后在新节点启动后关闭旧节点。虽然这是有效的，但这种方法存在一些问题：

- When you shutdown the old nodes, you would be taking down the running pods  with it. What if the pods need to clean up for a graceful shutdown? The  underlying VM technology might not wait for the clean up process.
 - What if you shutdown all the nodes at the same time? You could have a brief  outage while the pods are relaunched into the new nodes.

- 当您关闭旧节点时，您将同时关闭正在运行的 Pod。如果 Pod 需要清理以正常关闭怎么办？底层 VM 技术可能不会等待清理过程。
- 如果您同时关闭所有节点会怎样？当 Pod 重新启动到新节点时，您可能会出现短暂中断。

What we want is a way to gracefully migrate the pods off of the old nodes to ensure that none of our workloads are running while we make changes to  the node. Or if we are doing a full replacement of the cluster as in the example (e.g replacing VM images), we want to move the workloads off of the old nodes to the new nodes. In both cases, we want to prevent new  pods from being scheduled on the old node and then evict all the running pods off of it. We can use the `kubectl drain` command to achieve this.

我们想要的是一种从旧节点优雅地迁移 pod 的方法，以确保在我们对节点进行更改时没有任何工作负载在运行。或者，如果我们像示例中那样完全替换集群（例如替换 VM 映像），我们希望将工作负载从旧节点转移到新节点。在这两种情况下，我们都希望防止在旧节点上调度新 pod，然后将所有正在运行的 pod 驱逐出它。我们可以使用 `kubectl drain` 命令来实现这一点。

## Rescheduling Pods off of a node

## 从节点重新调度 Pod

The drain operation achieves the goal of rescheduling all the Pods off of  the node. During the drain operation, the node is marked as  unschedulable (the `NoSchedule`  taint). This prevents new pods from being scheduled on the node. Afterwards, the drain operation starts evicting the pods from the node,  shutting down the containers that are currently running on the node by  sending the `TERM` signal to the underlying containers of the pods.

Drain 操作实现了将所有 Pod 重新调度离开节点的目标。在 Drain 操作期间，节点被标记为不可调度（“NoSchedule”污点）。这可以防止在节点上调度新的 pod。之后，drain 操作开始从节点驱逐 Pod，通过向 Pod 的底层容器发送“TERM”信号来关闭当前在节点上运行的容器。

Although `kubectl drain` will gracefully handle pod eviction, there are still two factors that  could cause service disruption during the drain operation:

尽管 `kubectl drain` 会优雅地处理 pod eviction，但仍有两个因素可能在 Drain 操作期间导致服务中断：

- Your service application needs to be able to gracefully handle the `TERM` signal. When a Pod is evicted, Kubernetes will send the `TERM` signal to the container, and then will wait for a configurable amount  of time for the container to shutdown after giving the signal before  forcefully terminating it. However, if your containers do not handle the signal gracefully, you could still shutdown the pods uncleanly if it is in the middle of doing work (e.g committing a database transaction).
 - You lose all the pods servicing your application. Your service could  experience downtime while the new containers are being started on the  new nodes, or, if you did not deploy your pods with controllers, they  could end up never restarting.

- 您的服务应用程序需要能够优雅地处理“TERM”信号。当 Pod 被驱逐时，Kubernetes 会向容器发送“TERM”信号，然后在发出信号后等待可配置的时间让容器关闭，然后再强行终止它。但是，如果您的容器不能正常处理信号，如果 pod 正在执行工作（例如提交数据库事务），您仍然可以不干净地关闭 pod。
- 您丢失了为您的应用程序提供服务的所有 pod。在新节点上启动新容器时，您的服务可能会停机，或者，如果您没有使用控制器部署 pod，它们最终可能永远不会重新启动。

## Avoiding outages

## 避免中断

To minimize downtime from a voluntary disruption like draining a node,  Kubernetes provides the following disruption handling features:

- [Graceful termination](https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods)
 - [Lifecycle hooks](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/)
 - [PodDisruptionBudgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work) 

为了最大限度地减少因节点耗尽等自愿中断造成的停机时间，Kubernetes 提供了以下中断处理功能：

- [优雅终止](https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods)

- [生命周期钩子](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/)

- [PodDisruptionBudgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work)

In the rest of the series, we will use these features of Kubernetes to  mitigate service disruption from an eviction event. To make it easier to follow along, we will use our example above with the following resource config:

在本系列的其余部分，我们将使用 Kubernetes 的这些功能来减轻驱逐事件造成的服务中断。为了更容易理解，我们将使用上面的示例和以下资源配置：

```
 ---
 apiVersion: apps/v1
 kind: Deployment
 metadata:
   name: nginx-deployment
   labels:
     app: nginx
 spec:
   replicas: 2
   selector:
     matchLabels:
       app: nginx
   template:
     metadata:
       labels:
         app: nginx
     spec:
       containers:
       - name: nginx
         image: nginx:1.15
         ports:
         - containerPort: 80
 ---
 kind: Service
 apiVersion: v1
 metadata:
   name: nginx-service
 spec:
   selector:
     app: nginx
   ports:
   - protocol: TCP
     targetPort: 80
     port: 80
```


This config is a minimal example of a `Deployment` resource that manages multiple nginx pods (in our case two). This  resource will work towards maintaining two nginx pods in the cluster. Additionally, the config will provision a `Service` resource that can be used to access the nginx pods within the cluster.

此配置是管理多个 nginx pod（在我们的示例中为两个）的“部署”资源的最小示例。该资源将致力于在集群中维护两个 nginx pod。此外，配置将提供一个“服务”资源，可用于访问集群内的 nginx pod。

We will incrementally add to this throughout this series to build up to a  final config that implements all of the features Kubernetes provides to  minimize downtime during a maintenance operation. Here is our roadmap:

- [Gracefully shutting down Pods](https://medium.com/p/328aecec90d)

 - [Delaying Shutdown to Wait for Pod Deletion Propagation](https://medium.com/p/445f779a8304)

 - [Avoiding outages with PodDisruptionBudgets](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085

我们将在本系列中逐步添加，以构建最终配置，该配置实现了 Kubernetes 提供的所有功能，以最大限度地减少维护操作期间的停机时间。这是我们的路线图：

- [优雅地关闭 Pod](https://medium.com/p/328aecec90d)

- [延迟关闭以等待 Pod 删除传播](https://medium.com/p/445f779a8304)

- [使用 PodDisruptionBudgets 避免中断](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085)

Head on over to [the next post](https://medium.com/p/328aecec90d) to learn how you can leverage lifecycle hooks to gracefully shutdown your Pods.

前往 [下一篇文章](https://medium.com/p/328aecec90d) 了解如何利用生命周期钩子优雅地关闭 Pod。

*To get a fully implemented and tested version of zero downtime Kubernetes cluster updates on AWS and more, check out* [*Gruntwork.io*](http://gruntwork.io)*.*

*要在 AWS 等上获得完全实施和测试的零停机 Kubernetes 集群更新版本，请查看* [*Gruntwork.io*](http://gruntwork.io)*.*

Thanks to Yevgeniy Brikman.

The Gruntwork Blog 
Gruntwork 博客
- [Kubernetes](https://blog.gruntwork.io/tagged/kubernetes)