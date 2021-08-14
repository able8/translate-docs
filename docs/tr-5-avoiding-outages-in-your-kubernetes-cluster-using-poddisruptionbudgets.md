# Avoiding Outages in your Kubernetes Cluster using PodDisruptionBudgets

# 使用 PodDisruptionBudgets 避免 Kubernetes 集群中断

Preventing Pods from Eviction with Pod Disruption Budgets in Kubernetes

使用 Kubernetes 中的 Pod 中断预算防止 Pod 被驱逐

This is part 4 of [our journey](https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33) to implementing a zero downtime update of our Kubernetes cluster. In the previous two posts ([part 2](https://blog.gruntwork.io/gracefully-shutting-down-pods-in-a-kubernetes-cluster-328aecec90d) and [part 3](https://blog .gruntwork.io/delaying-shutdown-to-wait-for-pod-deletion-propagation-445f779a8304)), we focused on how to gracefully shutdown the existing Pods in our cluster. We covered how to use `preStop` hooks to gracefully shutdown pods and why it is important to add delays in the sequence to wait for the deletion event to propagate through the cluster. This can handle terminating one pod, but does not prevent us  from shutting down too many pods such that our service can’t function. In this post, we will use [PodDisruptionBudgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work) or PDB for short to mitigate this risk.

这是 [我们的旅程](https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33) 的第 4 部分，以实现我们的 Kubernetes 集群的零停机更新。在前两篇文章（[第 2 部分](https://blog.gruntwork.io/gracefully-shutting-down-pods-in-a-kubernetes-cluster-328aecec90d) 和 [第 3 部分](https://blog .gruntwork.io/delaying-shutdown-to-wait-for-pod-deletion-propagation-445f779a8304))，我们专注于如何优雅地关闭集群中现有的 Pod。我们介绍了如何使用 `preStop` 钩子优雅地关闭 pod，以及为什么在序列中添加延迟以等待删除事件在集群中传播很重要。这可以处理终止一个 pod，但并不能阻止我们关闭太多 pod，从而导致我们的服务无法运行。在这篇文章中，我们将使用 [PodDisruptionBudgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work) 或简称 PDB 来降低这种风险。

## PodDisruptionBudgets: Budgeting the Number of Faults to Tolerate

## PodDisruptionBudgets：预算可容忍的故障数量

A pod disruption budget is an indicator of the number of disruptions that can be tolerated at a given time for a class of pods (a budget of  faults). Whenever a disruption to the pods in a service is calculated to cause the service to drop below the budget, the operation is paused  until it can maintain the budget. This means that the drain event could  be temporarily halted while it waits for more pods to become available  such that the budget isn’t crossed by evicting the pods.

Pod 中断预算是在给定时间对于一类 Pod（故障预算）可以容忍的中断数量的指标。每当计算出服务中 Pod 的中断导致服务低于预算时，操作就会暂停，直到它可以维持预算。这意味着可以在等待更多 Pod 可用时暂时停止耗尽事件，这样驱逐 Pod 就不会超出预算。

To configure a pod disruption budget, we will create a `PodDisruptionBudget` resource that matches the pods in the service. For example, if we  wanted to create a pod disruption budget where we always want at least 1 Nginx pod to be available for our example deployment, we will apply the following config:

为了配置 Pod 中断预算，我们将创建一个与服务中的 Pod 匹配的 `PodDisruptionBudget` 资源。例如，如果我们想创建一个 pod 中断预算，我们总是希望至少有 1 个 Nginx pod 可用于我们的示例部署，我们将应用以下配置：

```
 apiVersion: policy/v1beta1
 kind: PodDisruptionBudget
 metadata:
   name: nginx-pdb
 spec:
   minAvailable: 1
   selector:
     matchLabels:
       app: nginx
 ```

 
This indicates to Kubernetes that we want at least 1 pod that matches the label `app: nginx` to be available at any given time. Using this, we can induce Kubernetes to wait for the pod in one drain request to be replaced before evicting the pods in a second drain request.

这向 Kubernetes 表明我们希望在任何给定时间至少有 1 个与标签“app: nginx”匹配的 pod 可用。使用这个，我们可以诱导 Kubernetes 在第二个排水请求中驱逐 Pod 之前等待一个排水请求中的 pod 被替换。

## Example

＃＃ 例子

To illustrate how this works, let’s go back to our example. For the sake  of simplicity, we will ignore any prestop hooks, readiness probes, and  service requests in this example. We will also assume that we want to do a one to one replacement of the cluster nodes. This means that we will  expand our cluster by doubling the number of nodes, with the new nodes  running the new image.

为了说明这是如何工作的，让我们回到我们的例子。为简单起见，我们将在此示例中忽略任何 prestop 挂钩、就绪探测和服务请求。我们还将假设我们想要对集群节点进行一对一的替换。这意味着我们将通过将节点数量加倍来扩展集群，新节点运行新映像。

So starting with our original cluster of two nodes:

所以从我们最初的两个节点集群开始：

![img](https://miro.medium.com/max/1400/1*txdjdbrj195vSXazKRhCUw.png)

We provision two additional nodes here running the new VM images. We will  eventually replace all the Pods on the old nodes on to the new ones:

我们在这里预置了两个额外的节点来运行新的 VM 映像。我们最终会将旧节点上的所有 Pod 替换为新节点：

![img](https://miro.medium.com/max/1400/1*OPsctr1Z8XhJ4gof8GfUbw.png)

To replace the Pods, we will first need to drain the old nodes. In this  example, let’s see what happens when we concurrently issue the drain  command to both nodes that were running our Nginx pods. The drain  request will be issued in two threads (in practice, this is just two  terminal tabs), each managing the drain sequence for one of the nodes.

要更换 Pod，我们首先需要排空旧节点。在这个例子中，让我们看看当我们同时向运行 Nginx pod 的两个节点发出排放命令时会发生什么。排放请求将在两个线程中发出（实际上，这只是两个终端选项卡），每个线程管理节点之一的排放序列。

Note that up to this point we were simplifying the examples by assuming that the drain command immediately issues an eviction request. In reality,  the drain operation involves tainting nodes (with the `NoSchedule` taint) first so that new pods won’t be scheduled on the nodes. For this example, we will look at the two phases individually.

请注意，到目前为止，我们通过假设排放命令立即发出驱逐请求来简化示例。实际上，drain 操作首先涉及污染节点（带有 `NoSchedule` 污染），这样新的 pod 就不会被调度到节点上。对于此示例，我们将分别查看两个阶段。

So to start, the two threads managing the drain sequence will taint the nodes so that new pods won’t be scheduled:

因此，首先，管理排放序列的两个线程将污染节点，从而不会调度新的 pod：

![img](https://miro.medium.com/max/1400/1*l8XTAFW-ia17FZbWs-omgw.png)

After the tainting completes, the drain threads will start evicting the pods  on the nodes. As part of this, the drain thread will query the control  plane to see if the eviction will cause the service to drop below the  configured Pod Disruption Budget (PDB). 
污染完成后，drain 线程将开始驱逐节点上的 pod。作为其中的一部分，drain 线程将查询控制平面以查看驱逐是否会导致服务低于配置的 Pod 中断预算 (PDB)。
Note that the control plane will serialize the requests, processing one PDB  inquiry at a time. As such, in this case, the control plane will respond to one of the requests with a success, while failing the other. This is because the first request is based on 2 pods being available. Allowing  this request would drop the number of pods available to 1, which means  the budget is maintained. When it allows the request to proceed, one of  the pods is then evicted, thereby becoming unavailable. At that point,  when the second request is processed, the control plane will reject it  because allowing that request would drop the number of available pods  down to 0, dropping below our configured budget.

请注意，控制平面将序列化请求，一次处理一个 PDB 查询。因此，在这种情况下，控制平面将成功响应其中一个请求，而另一个失败。这是因为第一个请求基于 2 个可用的 pod。允许此请求会将可用的 pod 数量降至 1，这意味着预算得以维持。当它允许请求继续时，其中一个 pod 将被驱逐，从而变得不可用。此时，当第二个请求被处理时，控制平面将拒绝它，因为允许该请求会将可用 Pod 的数量降至 0，低于我们配置的预算。

Given that, in this example we will assume that node 1 was the one that got  the successful response. In this case, the drain thread for node 1 will  proceed to evict the pods, while the drain thread for node 2 will wait  and try again later:

鉴于此，在此示例中，我们将假设节点 1 是获得成功响应的节点。在这种情况下，节点 1 的排放线程将继续驱逐 pod，而节点 2 的排放线程将等待并稍后重试：

![img](https://miro.medium.com/max/1400/1*F2PF4lNrG2f9kdUpoy9vKg.png)

![img](https://miro.medium.com/max/1400/1*r4I3akpcARD5K2KmeZ2qPg.png)

When the pod in node 1 is evicted, it is immediately recreated in one of the available nodes by the `Deployment` controller. In this case, since our old nodes are tainted with the `NoSchedule` taint, the scheduler will choose one of the new nodes:

当节点 1 中的 pod 被驱逐时，“Deployment”控制器会立即在可用节点之一中重新创建它。在这种情况下，由于我们的旧节点被“NoSchedule”污染，调度程序将选择新节点之一：

![img](https://miro.medium.com/max/1400/1*Ydl0N4Jn1DEmchCrgJbUiA.png)

At this point, now that the pod has been replaced successfully on the new  node and the original node is drained, the thread for draining node 1  completes.

至此，新节点上的pod已经替换成功，原节点排空，排空节点1的线程完成。

From this point on, when the drain thread for node 2 tries again to query  the control plane about the PDB again, it will succeed. This is because  there is a pod running that is not in consideration for eviction, so  allowing the drain thread for node 2 to progress further won’t drop the  number of available pods down below the budget. So the thread progresses to evict the pods and eventually completes the eviction process:

从这点开始，当节点2的drain线程再次尝试再次查询有关PDB的控制平面时，它将成功。这是因为有一个正在运行的 Pod 不考虑驱逐，因此允许节点 2 的排放线程进一步推进不会将可用 Pod 的数量降低到预算以下。因此线程继续驱逐 Pod 并最终完成驱逐过程：

![img](https://miro.medium.com/max/1400/1*5xNcTfZ1N7N7oyKmjXuIjA.png)

![img](https://miro.medium.com/max/1400/1*eoZ48zIbCeqfk6Gbr6aGoQ.png)

![img](https://miro.medium.com/max/1400/1*riu9fA1JjOynt7S5uBPj6w.png)

With that, we have successfully migrated both pods to the new nodes, without ever having a situation where we have no pods to service the  application. Moreover, we did not need to have any coordination logic  between the two threads, as Kubernetes handled all that for us based on  the config we provided!

有了这个，我们已经成功地将两个 pod 迁移到了新节点，从来没有出现过没有 pod 来为应用程序提供服务的情况。此外，我们不需要在两个线程之间有任何协调逻辑，因为 Kubernetes 根据我们提供的配置为我们处理了所有这些！

# Summary

＃ 概括

So to tie all that together, in this blog post series we covered:

因此，为了将所有这些联系在一起，在这个博客文章系列中，我们介绍了：

- How to use lifecycle hooks to implement the ability to gracefully shutdown  our applications so that they are not abruptly terminated.
 - How pods are removed from the system and why it is necessary to introduce delays in the shutdown sequence.
 - How to specify pod disruption budgets to ensure that we always have a  certain number of pods available to continuously service a functioning  application in the face of disruption.

- 如何使用生命周期钩子实现优雅关闭我们的应用程序的能力，以便它们不会突然终止。
- 如何从系统中移除 Pod 以及为什么需要在关闭序列中引入延迟。
- 如何指定 pod 中断预算以确保我们始终有一定数量的 pod 可用，以便在遇到中断时持续为正常运行的应用程序提供服务。

When these features are all used together, we are able to achieve our goal of a zero downtime rollout of instance updates!

当这些功能全部一起使用时，我们就能够实现实例更新的零停机时间推出的目标！

But don’t just take my word for it! Go ahead and take this configuration  out for a spin. You can even write automated tests using terratest, by  leveraging the functions in [the k8s module](https://godoc.org/github.com/gruntwork-io/terratest/modules/k8s), and [the ability to continuously check an endpoint](https://godoc.org/github.com/gruntwork-io/terratest/modules/http-helper#ContinuouslyCheckUrl). After all, one of the important [lessons we learned from writing 300k lines of infrastructure code](https://blog.gruntwork.io/5-lessons-learned-from-writing-over-300-000-lines-of- infrastructure-code-36ba7fadeac1) is that infrastructure code without automated tests is broken.

但不要只相信我的话！继续并尝试使用此配置。您甚至可以使用 terratest 编写自动化测试，利用 [k8s 模块](https://godoc.org/github.com/gruntwork-io/terratest/modules/k8s) 中的功能和 [持续检查的能力一个端点]（https://godoc.org/github.com/gruntwork-io/terratest/modules/http-helper#ContinuouslyCheckUrl）。毕竟，[我们从编写 30 万行基础设施代码中学到的经验教训]（https://blog.gruntwork.io/5-lessons-learned-from-writing-over-300-000-lines-of- Infrastructure-code-36ba7fadeac1) 是没有自动化测试的基础架构代码被破坏。

*To get a fully implemented version of zero downtime Kubernetes cluster updates on AWS and more, check out* [*Gruntwork.io*](http://gruntwork.io)*.*

*要在 AWS 等上获得完全实施的零停机 Kubernetes 集群更新版本，请查看* [*Gruntwork.io*](http://gruntwork.io)*.*

The Gruntwork Blog 
Gruntwork 博客
