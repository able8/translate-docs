# Delaying Shutdown to Wait for Pod Deletion Propagation

# 延迟关闭以等待 Pod 删除传播

Delaying Shutdown of Pods In Kubernetes

延迟关闭 Kubernetes 中的 Pod

This is part 3 of [our journey](https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33) to implementing a zero downtime update of our Kubernetes cluster. In [part 2 of the series](https://blog.gruntwork.io/gracefully-shutting-down-pods-in-a-kubernetes-cluster-328aecec90d), we mitigated downtime from an unclean shutdown by leveraging lifecycle  hooks to implement a graceful termination of our application pods. However, we also learned that the Pod potentially continues to receive  traffic after the shutdown sequence is initiated. This means that the  end clients may get error messages because they are routed to a pod that is no longer able to service traffic. Ideally, we would like the pods  to stop receiving traffic immediately after they have been evicted. To  mitigate this, we must first understand why this is happening.

这是 [我们的旅程](https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33) 的第 3 部分，以实现我们的 Kubernetes 集群的零停机更新。在 [系列的第 2 部分](https://blog.gruntwork.io/gracefully-shutting-down-pods-in-a-kubernetes-cluster-328aecec90d) 中，我们通过利用生命周期钩子来减少不正常关机造成的停机时间实现我们的应用程序 pod 的优雅终止。但是，我们还了解到，在启动关闭序列后，Pod 可能会继续接收流量。这意味着最终客户端可能会收到错误消息，因为它们被路由到不再能够为流量提供服务的 pod。理想情况下，我们希望 Pod 在被驱逐后立即停止接收流量。为了缓解这种情况，我们必须首先了解为什么会发生这种情况。

*A lot of the information in this post was learned from the book “*[*Kubernetes in Action*](https://www.manning.com/books/kubernetes-in-action)*” by Marko Lukša. You can find an excerpt of the relevant section* [*here*](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/)*. In addition to the material covered here, the book provides an  excellent overview of best practices for running applications on  Kubernetes and we highly recommend reading it.*

*这篇文章中的很多信息都是从 Marko Lukša 的“*[*Kubernetes in Action*](https://www.manning.com/books/kubernetes-in-action)*”一书中学到的。您可以找到相关部分的摘录* [*此处*](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/)*。除了此处涵盖的材料外，本书还对在 Kubernetes 上运行应用程序的最佳实践进行了出色的概述，我们强烈建议您阅读。*

## Pod Shutdown Sequence

## Pod 关闭序列

In [the previous post](https://blog.gruntwork.io/gracefully-shutting-down-pods-in-a-kubernetes-cluster-328aecec90d#d7e5), we covered the Pod eviction lifecycle. If you recall from the post, the first step in the eviction sequence is for the pod to be deleted, which starts a chain of events that ultimately results in the pod being  removed from the system. However, what we didn’t talk about is how the  pod gets deregistered from the `Service` so that it stops receiving traffic.

在 [上一篇文章](https://blog.gruntwork.io/gracefully-shutting-down-pods-in-a-kubernetes-cluster-328aecec90d#d7e5) 中，我们介绍了 Pod 驱逐生命周期。如果你还记得这篇文章，驱逐序列的第一步是删除 pod，这会启动一系列事件，最终导致 pod 从系统中删除。然而，我们没有讨论的是 pod 如何从“Service”中注销，从而停止接收流量。

So what causes the pod to be removed from the `Service` ? To understand this, we need to drop one layer deeper into understanding what happens when a pod is removed from the cluster.

那么是什么导致 pod 从 `Service` 中移除？要理解这一点，我们需要更深入地了解从集群中删除 pod 时会发生什么。

When a pod is removed from the cluster via the API, all that is happening is that the pod is marked for deletion in the metadata server. This sends a pod deletion notification to all relevant subsystems that then handle  it:

当通过 API 从集群中删除一个 Pod 时，所发生的一切就是该 Pod 在元数据服务器中被标记为删除。这会向所有相关子系统发送一个 pod 删除通知，然后处理它：

- The `kubelet` running the pod starts the shutdown sequence described in the previous post.
 - The `kube-proxy` daemon running on all the nodes will remove the pod’s ip address from `iptables`.
 - The endpoints controller will remove the pod from the list of valid endponts, in turn removing the pod from the `Service` .

- 运行 pod 的 `kubelet` 启动上一篇文章中描述的关闭序列。
- 在所有节点上运行的 `kube-proxy` 守护进程将从 `iptables` 中删除 pod 的 IP 地址。
- 端点控制器将从有效端点列表中删除 pod，进而从 `Service` 中删除该 pod。

You don’t need to know the details of each system. The important point here is that there are multiple systems involved, potentially running on  different nodes, with the sequences happening in parallel. Because of  this, it is very likely that the pod runs the `preStop` hook and receives the `TERM` signal well before the pod is removed from all active lists. This is  why the pod continues to receive traffic even after the shutdown  sequence is initiated.

您无需了解每个系统的详细信息。这里的重点是涉及多个系统，可能运行在不同的节点上，序列并行发生。因此，在从所有活动列表中删除 pod 之前，pod 很可能运行 `preStop` 钩子并接收到 `TERM` 信号。这就是为什么即使在启动关闭序列后 pod 仍继续接收流量的原因。

## Mitigating the Problem

## 缓解问题

On the surface it may seem that what we want to do here is chain the  sequence of events so that the pod isn’t shutdown until after it has  been deregistered from all the relevant subsystems. However, this is  hard to do in practice due to the distributed nature of the Kuberenetes  system. What happens if one of the nodes experiences a network  partition? Do you indefinitely wait for the propagation? What if that  node comes back online? What if you have 1000s of nodes you have to wait for? 10s of thousands?

从表面上看，我们在这里想要做的似乎是将事件序列链接起来，以便 pod 在从所有相关子系统注销之前不会关闭。然而，由于 Kuberenetes 系统的分布式特性，这在实践中很难做到。如果其中一个节点遇到网络分区会怎样？您是否无限期地等待传播？如果该节点重新上线怎么办？如果您有 1000 个节点需要等待怎么办？几十万？

Unfortunately there isn’t a perfect solution here to prevent all outages. What we can do however, is to introduce a sufficient delay in the shutdown sequence to capture 99% of the cases. To do this, we introduce a `sleep` in the `preStop` hook that delays the shutdown sequence. Let’s look at how this works in our example.

不幸的是，这里没有完美的解决方案来防止所有中断。然而，我们可以做的是在关闭序列中引入足够的延迟以捕获 99% 的情况。为此，我们在 `preStop` 钩子中引入了一个 `sleep` 来延迟关闭序列。让我们看看它在我们的例子中是如何工作的。

We will need to update our config to introduce the delay as part of the `preStop` hook. In “Kubernetes in Action”, Lukša recommends 5–10 seconds, so here we will use 5 seconds:

我们需要更新我们的配置以引入延迟作为 `preStop` 钩子的一部分。在“Kubernetes in Action”中，Lukša 推荐 5-10 秒，所以这里我们将使用 5 秒：

```
 lifecycle:
   preStop:
     exec:
       command: [
         "sh", "-c",
         # Introduce a delay to the shutdown sequence to wait for the
         # pod eviction event to propagate. Then, gracefully shutdown
         # nginx.
         "sleep 5 && /usr/sbin/nginx -s quit",
       ]
 ``` 


Now let’s walk through what happens during the shutdown sequence in our example. Like in the previous post, we will start with `kubectl drain` , which will evict the pods on the nodes. This will send a delete pod event that notifies the `kubelet` and the Endpoint Controller (which manages the `Service` endpoints) simultaneously. Here, we assume the `preStop` hook starts before the controller removes the pod.


现在让我们看看在我们的示例中关闭序列期间会发生什么。和上一篇文章一样，我们将从 `kubectl drain` 开始，它将驱逐节点上的 pod。这将发送一个删除 pod 事件，同时通知 `kubelet` 和 Endpoint Controller（管理 `Service` 端点）。在这里，我们假设 `preStop` 钩子在控制器移除 pod 之前启动。


![img](https://miro.medium.com/max/1400/1*ClRbAVug7UJz4lts4lIwqg.png)

Drain node will remove the pod, which in turn sends a deletion event.

Drain 节点将移除 pod，而 pod 又会发送一个删除事件。

At this point the `preStop` hook starts, which will delay the shutdown sequence by 5 seconds. During this time, the Endpoint Controller will remove the pod:


此时`preStop`钩子开始，这将延迟5秒的关闭序列。在此期间，端点控制器将移除 pod：


![img](https://miro.medium.com/max/1400/1*b8JZKTQdIIz-fMr7DKlSkQ.png)

The pod is removed from the controller while the shutdown sequence is delayed.

当关闭序列延迟时，吊舱从控制器中移除。

Note that during this delay, the pod is still up so even if it receives  connections, the pod is still able to handle the connections. Additionally, if any clients try to connect after the pod is removed  from the controller, they will not be routed to the pod being shutdown. So in this scenario, assuming the Controller handles the event during  the delay period, there will be no downtime.

请注意，在此延迟期间，Pod 仍在运行，因此即使它接收到连接，Pod 仍然能够处理连接。此外，如果任何客户端在 pod 从控制器中移除后尝试连接，它们将不会被路由到正在关闭的 pod。所以在这个场景中，假设Controller在延迟期间处理事件，就不会出现宕机。

Finally, to complete the picture, the `preStop` hook comes out of the sleep and shuts down the Nginx pod, removing the pod from the node:


最后，为了完成图片，`preStop` 钩子从睡眠中出来并关闭 Nginx pod，从节点中移除 pod：


![img](https://miro.medium.com/max/1400/1*iKrPivAfmA-Ooxzi6dBASg.png)


![img](https://miro.medium.com/max/1400/1*jYpTNfVS5M1tkiQtAomDWQ.png)

At this point, it is safe to do any upgrades on Node 1, including  rebooting the node to load a new kernel version. We can also shutdown  the node if we had launched a new node to house the workload that was  already running on it (see [the next section on PodDisruptionBudgets](https://blog.gruntwork.io/delaying-shutdown-to-wait -for-pod-deletion-propagation-445f779a8304#89eb)).

此时，可以安全地在节点 1 上进行任何升级，包括重新启动节点以加载新的内核版本。如果我们启动了一个新节点来容纳已经在其上运行的工作负载，我们也可以关闭该节点（请参阅[关于 PodDisruptionBudgets 的下一节](https://blog.gruntwork.io/delaying-shutdown-to-wait -for-pod-deletion-propagation-445f779a8304#89eb))。

## Recreating Pods

## 重新创建 Pod

If you made it this far, you might be wondering how we recreate the Pods  that were originally scheduled on the node. Now we know how to  gracefully shutdown the Pods, but what if it is critical to get back to  the original number of Pods running? This is where the `Deployment` resource comes into play.

如果您做到了这一点，您可能想知道我们如何重新创建最初在节点上调度的 Pod。现在我们知道如何优雅地关闭 Pod，但是如果恢复到原来运行的 Pod 数量很重要呢？这就是“部署”资源发挥作用​​的地方。

The `Deployment` resource is called [a controller,](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) and it does the work of maintaining a specified desired state on the  cluster. If you recall our resource config, we do not directly create  Pods in the config. Instead, we use `Deployment` resources to automatically manage our Pods for us by providing it a template for how to create the Pods. This is what the `template` section is in our config:

`Deployment` 资源称为 [a controller,](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)，它负责在集群上维护指定的所需状态。如果您还记得我们的资源配置，我们不会直接在配置中创建 Pod。相反，我们使用“部署”资源为我们自动管理我们的 Pod，为我们提供如何创建 Pod 的模板。这是我们配置中的`template`部分：

```
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
 ```

 
This specifies that Pods in our `Deployment` should be created with the label `app: nginx` and one container running the `nginx:1.15` image, exposing port `80` .

这指定了我们的 `Deployment` 中的 Pod 应该使用标签 `app: nginx` 和一个运行 `nginx:1.15` 镜像的容器创建，暴露端口 `80`。

In addition to the Pod template, we also provide a spec to the `Deployment` resource specifying the number of replicas it should maintain:

除了 Pod 模板，我们还为 `Deployment` 资源提供了一个规范，指定它应该维护的副本数量：

```
 spec:
   replicas: 2
 ```

 
This notifies the controller that it should try to maintain 2 Pods running  on the cluster. Anytime the number of running Pods drops, the controller will automatically create a new one to replace it. So in our case, the  when we evict the Pod from the node during a drain operation, the `Deployment` controller automatically recreates it on one of the other available nodes.

这会通知控制器它应该尝试维护在集群上运行的 2 个 Pod。每当正在运行的 Pod 数量下降时，控制器会自动创建一个新的来替换它。因此，在我们的例子中，当我们在 Drain 操作期间从节点中驱逐 Pod 时，“Deployment”控制器会自动在其他可用节点之一上重新创建它。

## Summary

＃＃ 概括

In sum, with adequate delays in the `preStop` hooks and graceful termination, we can now shutdown our pods gracefully on a single node. And with `Deployment` resources, we can automatically recreate the shut down Pods. But what  if we want to replace all the nodes in the cluster at once?

总之，通过`preStop` 钩子的足够延迟和优雅终止，我们现在可以在单个节点上优雅地关闭我们的 pod。使用“部署”资源，我们可以自动重新创建关闭的 Pod。但是如果我们想一次性替换集群中的所有节点呢？

If we naively drain all the nodes, we could end up with downtime because  the service load balancer may end up with no pods available. Worse, for a stateful system, we may wipe out our quorum, causing extended downtime  as the new pods come up and have to go through leader election waiting  for a quorum of nodes to come back up.

如果我们天真地耗尽所有节点，我们最终可能会停机，因为服务负载均衡器可能最终没有可用的 pod。更糟糕的是，对于有状态的系统，我们可能会消灭我们的法定人数，随着新 Pod 的出现而导致停机时间延长，并且必须通过领导者选举等待法定人数的节点恢复。

If we instead try to drain nodes one at a time, we could end up with new  Pods being launched on the remaining old nodes. This risks a situation  where we might end up with all our Pod replicas running on one of the  old nodes, so that when we get to drain that one, we still lose all our  Pod replicas. 
如果我们尝试一次排空一个节点，我们最终可能会在剩余的旧节点上启动新的 Pod。这存在一种风险，即我们可能最终将所有 Pod 副本都运行在一个旧节点上，因此当我们耗尽该节点时，我们仍然会丢失所有 Pod 副本。
To handle this situation, Kubernetes offers a feature called `PodDisruptionBudgets`, which indicate tolerance for the number of pods that can be shutdown at any given point in time. In [the next and final part of our series](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085), we will cover how we can use this to control the number of drain events that can happen concurrently despite our naive approach to issue a  drain call for all nodes in parallel.

为了处理这种情况，Kubernetes 提供了一个名为“PodDisruptionBudgets”的功能，它表示可以在任何给定时间点关闭的 pod 数量的容差。在 [我们系列的下一部分也是最后一部分](https://blog.gruntwork.io/avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets-ef6a4baa5085) 中，我们将介绍如何使用它来尽管我们为所有节点并行发出排放调用的幼稚方法，但控制可以同时发生的排放事件的数量。

*To get a fully implemented version of zero downtime Kubernetes cluster updates on AWS and more, check out* [*Gruntwork.io*](http://gruntwork.io)*.*

*要在 AWS 等上获得完全实施的零停机 Kubernetes 集群更新版本，请查看* [*Gruntwork.io*](http://gruntwork.io)*.*

The Gruntwork Blog

Gruntwork 博客

- [Kubernetes](https://blog.gruntwork.io/tagged/kubernetes) 
- [Kubernetes](https://blog.gruntwork.io/tagged/kubernetes)
