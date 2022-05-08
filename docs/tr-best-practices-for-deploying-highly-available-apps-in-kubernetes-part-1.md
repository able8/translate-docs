# Best practices for deploying highly available apps in Kubernetes. Part 1

# 在 Kubernetes 中部署高可用应用的最佳实践。第1部分

As you know, deploying a basic viable app configuration in Kubernetes is a breeze. On the other hand, trying to make your application as available and fault-tolerant as possible inevitably entails a great number of hurdles and pitfalls. In this article, we break down what we believe to be the most important rules when it comes to deploying high-availability applications in Kubernetes and sharing them in a concise way.

如您所知，在 Kubernetes 中部署一个基本可行的应用程序配置是轻而易举的事。另一方面，试图使您的应用程序尽可能地可用和容错不可避免地会带来大量的障碍和陷阱。在本文中，我们分解了我们认为在 Kubernetes 中部署高可用性应用程序并以简洁的方式共享它们时最重要的规则。

Note that we will not be using any features that aren’t available right out-of-the-box. What we will also not do is lock into specific CD solutions, and we will omit the issues of templating/generating Kubernetes manifests. In this article, we only discuss the final structure of Kubernetes manifests when deploying to the cluster.

请注意，我们不会使用任何开箱即用的功能。我们也不会做的是锁定特定的 CD 解决方案，我们将省略模板/生成 Kubernetes 清单的问题。在本文中，我们只讨论 Kubernetes manifests 在部署到集群时的最终结构。

## 1\. Number of replicas

## 1. 副本数

You need at least two replicas for the application to be considered minimally available. But why, you may ask, is a single replica not enough? The problem is that many entities in Kubernetes (Node, Pod, ReplicaSet, etc.) are ephemeral, i.e. under certain conditions, they may be automatically deleted/recreated. Obviously, the Kubernetes cluster and applications running in it must account for that.

您至少需要两个副本才能将应用程序视为最低可用。但是，您可能会问，为什么单个副本还不够？问题是 Kubernetes 中的许多实体（Node、Pod、ReplicaSet 等）都是短暂的，即在某些条件下，它们可能会被自动删除/重新创建。显然，Kubernetes 集群和在其中运行的应用程序必须考虑到这一点。

For example, when the autoscaler scales down your number of nodes, some of those nodes will be deleted, including the Pods running on them. If the sole instance of your application is running on one of the nodes to be deleted, you may find your application completely unavailable, though this is usually short-lived. In general, if you only have one replica of the application, any abnormal termination of it will result in downtime. In other words, you must have **at least two running replicas of the application.**

例如，当自动缩放器缩减您的节点数量时，其中一些节点将被删除，包括在它们上运行的 Pod。如果您的应用程序的唯一实例正在要删除的节点之一上运行，您可能会发现您的应用程序完全不可用，尽管这通常是短暂的。一般来说，如果你只有一个应用程序的副本，任何异常终止都会导致停机。换句话说，您必须拥有**至少两个正在运行的应用程序副本。**

The more replicas there are, the milder of a decline there will be in your application’s computing capacity in the event that some replica fails. For example, suppose you have two replicas and one fails due to network issues on a node. The load that the application can handle will be cut in half (with only one of the two replicas available). Of course, the new replica will be scheduled on a new node, and the load capacity of the application will be fully restored. But until then, increasing the load can lead to service disruptions, which is why you **must have some replicas in reserve.**

副本越多，在某些副本失败的情况下，应用程序的计算能力下降的幅度就越小。例如，假设您有两个副本，其中一个由于节点上的网络问题而失败。应用程序可以处理的负载将减半（只有两个副本中的一个可用）。当然，新的副本会被调度到一个新的节点上，完全恢复应用的负载能力。但在那之前，增加负载会导致服务中断，这就是为什么您**必须保留一些副本。**

_The above recommendations are relevant to cases in which there is no HorizontalPodAutoscaler used. The best alternative for applications that have more than a few replicas is to configure HorizontalPodAutoscaler and let it manage the number of replicas. We will focus on HorizontalPodAutoscaler in the next article._

_以上建议适用于没有使用 HorizontalPodAutoscaler 的情况。对于具有多个副本的应用程序，最好的替代方案是配置 HorizontalPodAutoscaler 并让它管理副本的数量。我们将在下一篇文章中重点介绍HorizontalPodAutoscaler。_

## 2\. The update strategy 

## 2. 更新策略

The default update strategy for Deployment entails a reduction of the number of old+new ReplicaSet Pods with a `Ready` status of 75% of their pre-update amount. Thus, during the update, the computing capacity of an application may drop to 75% of its regular level, and that may lead to a partial failure (degradation of the application’s performance). The `strategy.RollingUpdate.maxUnavailable` parameter allows you to configure the maximum percentage of Pods that can become unavailable during an update. Therefore, either make sure that your application runs smoothly even in the event that 25% of your Pods are unavailable or lower the `maxUnavailable` parameter. Note that the `maxUnavailable` parameter is rounded down.

Deployment 的默认更新策略需要减少旧 + 新 ReplicaSet Pod 的数量，其 `Ready` 状态为更新前数量的 75%。因此，在更新过程中，应用程序的计算能力可能会下降到其正常水平的 75%，这可能会导致部分故障（应用程序性能下降）。 `strategy.RollingUpdate.maxUnavailable` 参数允许您配置在更新期间可能变为不可用的 Pod 的最大百分比。因此，要么确保你的应用程序在 25% 的 Pod 不可用的情况下也能顺利运行，要么降低 `maxUnavailable` 参数。请注意，“maxUnavailable”参数是向下舍入的。

There’s a little trick to the default update strategy ( `RollingUpdate`): the application will temporarily have not only a few replicas, but two different versions (the old one and the new one) running concurrently as well. Therefore, if running different replicas and different versions of the application side by side is unfeasible for some reason, then you can use `strategy.type: Recreate`. Under the `Recreate` strategy, all the existing Pods are killed before the new Pods are created. This results in a short-lived downtime.

默认更新策略（`RollingUpdate`）有一个小技巧：应用程序将暂时不仅有几个副本，而且还会同时运行两个不同的版本（旧版本和新版本）。因此，如果由于某种原因并行运行不同副本和不同版本的应用程序是不可行的，那么您可以使用`strategy.type: Recreate`。在 `Recreate` 策略下，所有现有的 Pod 都会在新的 Pod 被创建之前被杀死。这会导致短暂的停机时间。

_Other deployment strategies (blue-green, canary, etc.) can often provide a much better alternative to the RollingUpdate strategy. However, we are not taking them into account in this article since their implementation depends on the software used to deploy the application. That goes beyond the scope of this article (here is a_ [_great article_](https://www.weave.works/blog/kubernetes-deployment-strategies) _on the topic that we recommend and is well worth the read)._

_其他部署策略（蓝绿、金丝雀等）通常可以提供比 RollingUpdate 策略更好的替代方案。但是，我们没有在本文中考虑它们，因为它们的实现取决于用于部署应用程序的软件。这超出了本文的范围（这里有一篇_[_great article_](https://www.weave.works/blog/kubernetes-deployment-strategies)_关于我们推荐并且非常值得一读的主题)。_

## 3\. Uniform replicas distribution across nodes

## 3. 跨节点的均匀副本分布

It is very important that you distribute Pods of the application across different nodes if you have multiple replicas of the application. To do so, **you can instruct your scheduler to avoid starting multiple Pods of the same Deployment on the same node:**

如果您有应用程序的多个副本，那么跨不同节点分发应用程序的 Pod 非常重要。为此，**您可以指示调度程序避免在同一节点上启动同一 Deployment 的多个 Pod：**

```
       affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchLabels:
                  app: testapp
              topologyKey: kubernetes.io/hostname
```

It is better to use `preferredDuringSchedulingaffinity` instead of `requiredDuringScheduling`. The latter may render it impossible to start new Pods if the number of nodes required for the new Pods is larger than the number of nodes available. Still, the `requiredDuringScheduling` affinity might come in handy when the number of nodes and application replicas is known in advance and you need to be sure that two Pods will not end up on the same node.

最好使用 `preferredDuringSchedulingaffinity` 而不是 `requiredDuringScheduling`。如果新 Pod 所需的节点数大于可用节点数，后者可能会导致无法启动新 Pod。尽管如此，当预先知道节点和应用程序副本的数量并且您需要确保两个 Pod 不会最终在同一个节点上时，“requiredDuringScheduling”亲和性可能会派上用场。

## 4\. Priority

## 4. 优先事项

`priorityClassName` represents your Pod priority. The scheduler uses it to decide which Pods are to be scheduled first and which Pods should be evicted first if there is no space for Pods left on the nodes.

`priorityClassName` 代表你的 Pod 优先级。调度程序使用它来决定首先调度哪些 Pod，如果节点上没有剩余 Pod 空间，则应该首先驱逐哪些 Pod。

You will need to add several [PriorityClass](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass) type resources and map Pods to them using `priorityClassName`. Here is an example of how `PriorityClasses` may vary:

您将需要添加几个 [PriorityClass](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass) 类型的资源并使用 `priorityClassName` 将 Pod 映射到它们。以下是“PriorityClasses”如何变化的示例：

- _Cluster_. Priority > 10000. Cluster-critical components, such as kube-apiserver.
- _Daemonsets_. Priority: 10000. Usually, it is not advised for DaemonSet Pods to be evicted from cluster nodes and replaced by ordinary applications.
- _Production-high_. Priority: 9000. Stateful applications.
- _Production-medium._ Priority: 8000. Stateless applications.
- _Production-low_. Priority: 7000. Less critical applications.
- _Default_. Priority: 0. Non-production applications.
- _Cluster_。优先级 > 10000。集群关键组件，例如 kube-apiserver。
- _守护进程_。优先级：10000。通常不建议将 DaemonSet Pod 从集群节点中逐出并替换为普通应用程序。

Setting priorities will help you to avoid sudden evictions of critical components. Also, critical applications will evict less important applications if there is a lack of node resources.

设置优先级将帮助您避免突然驱逐关键组件。此外，如果缺乏节点资源，关键应用程序将驱逐不太重要的应用程序。

## 5\. Stopping processes in containers

## 5. 停止容器中的进程

The signal specified in `STOPSIGNAL` (usually, the `TERM` signal) is sent to the process to stop it. However, some applications cannot handle it properly and cannot manage to shut down gracefully. The same is true for applications running in Kubernetes.

`STOPSIGNAL` 中指定的信号（通常是 `TERM` 信号）被发送到进程以停止它。但是，某些应用程序无法正确处理它并且无法正常关闭。在 Kubernetes 中运行的应用程序也是如此。

For example, in order to shut down nginx properly, you will need a `preStop` hook like this:

例如，为了正确关闭 nginx，您将需要一个 `preStop` 钩子，如下所示：

```
lifecycle:
preStop:
    exec:
      command:
      - /bin/sh
      - -ec
      - |
        sleep 3
        nginx -s quit
```

A brief explanation for this listing: 

此清单的简要说明：

1. `sleep 3` prevents race conditions that may be caused by an endpoint being deleted.
2. `nginx -s quit` shuts down nginx properly. This line isn’t required for more up-to-date images since the `STOPSIGNAL: SIGQUIT` parameter is set there by default.



1. `sleep 3` 防止可能由删除端点引起的竞争条件。
2. `nginx -s quit` 正常关闭nginx。更新的图像不需要此行，因为此处默认设置了 `STOPSIGNAL: SIGQUIT` 参数。

_(You can learn more about graceful shutdowns for nginx bundled with PHP-FPM in_ [_our other article_](https://blog.flant.com/graceful-shutdown-in-kubernetes-is-not-always-trivial/)_.)_

_（您可以在_ [_我们的另一篇文章_](https://blog.flant.com/graceful-shutdown-in-kubernetes-is-not-always-trivial/)中了解更多关于与 PHP-FPM 捆绑的 nginx 的优雅关闭_.)_

The way `STOPSIGNAL` is handled depends on the application itself. In practice, for most applications, you have to Google the way `STOPSIGNAL` is handled. If the signal is not handled appropriately the `preStop` hook can help you solve the problem. Another option is to replace `STOPSIGNAL` with a signal that the application can handle properly (and permit it to shut down gracefully).

处理“STOPSIGNAL”的方式取决于应用程序本身。实际上，对于大多数应用程序，您必须使用 Google 处理“STOPSIGNAL”的方式。如果信号处理不当，`preStop` 挂钩可以帮助您解决问题。另一种选择是将“STOPSIGNAL”替换为应用程序可以正确处理的信号（并允许它正常关闭）。

`terminationGracePeriodSeconds` is another crucial parameter important in shutting down the application. It specifies the time period for which the application is to shut down gracefully. If the application does not terminate within this time frame (30 seconds by default), it will receive a `KILL` signal. Thus, you will need to increase the terminationGracePeriodSeconds parameter if you think that running the `preStop` hook and/or shutting down the application at the `STOPSIGNAL` may take more than 30 seconds. For example, you may need to increase it if some requests from your web service clients take a long time to complete (e.g. requests that involve downloading large files).

`terminationGracePeriodSeconds` 是关闭应用程序的另一个重要参数。它指定应用程序正常关闭的时间段。如果应用程序没有在这个时间范围内（默认为 30 秒）终止，它将收到一个 `KILL` 信号。因此，如果您认为运行 `preStop` 挂钩和/或在 `STOPSIGNAL` 处关闭应用程序可能需要超过 30 秒，则需要增加terminationGracePeriodSeconds 参数。例如，如果来自 Web 服务客户端的某些请求需要很长时间才能完成（例如涉及下载大文件的请求），您可能需要增加它。

It is worth noting that the `preStop` hook has a locking mechanism, i.e. `STOPSIGNAL` may be sent only after the `preStop` hook has finished running. At the same time, the `terminationGracePeriodSeconds` countdown _continues during the_ _preStop_ _hook execution_. All the hook-induced processes, as well as the processes running in the container, will be `KILL` ed after `terminationGracePeriodSeconds` is over.

值得注意的是，`preStop` 钩子有一个锁定机制，即只有在 `preStop` 钩子完成运行后才能发送 `STOPSIGNAL`。同时，`terminationGracePeriodSeconds` 倒计时_在_preStop__hook 执行_期间_继续。在 `terminationGracePeriodSeconds` 结束后，所有钩子引发的进程以及在容器中运行的进程都将被 `KILL` 删除。

Also, some applications have specific settings that set the deadline at which point the application must terminate (for example, the `--timeout` option in Sidekiq). Therefore, in each case, you have to make sure that if the application has this setting, it has a slightly lower value than that of `terminationGracePeriodSeconds`.

此外，某些应用程序具有特定设置，用于设置应用程序必须终止的截止日期（例如，Sidekiq 中的 `--timeout` 选项）。因此，在每种情况下，您都必须确保如果应用程序具有此设置，则它的值略低于 `terminationGracePeriodSeconds` 的值。

## 6\. Reserving resources

## 6. 预留资源

The scheduler uses a Pod’s `resources.requests` to decide which node to place the Pod on. For instance, a Pod cannot be scheduled on a Node that does not have enough free (i.e., _non-requested_) resources to cover that Pod’s resource requests. On the other hand, `resources.limits` allow you to limit Pods’ resource consumption that heavily exceeds their respective requests. A good tip is to **set limits equal to requests**. Setting limits at much higher than requests may lead to a situation when some of a node’s Pods not getting the requested resources. This may lead to the failure of other applications on the node (or even the node itself). Kubernetes assigns a [QoS class](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod) to each Pod based on its resource scheme. K8s then uses QoS classes to make decisions about which Pods should be evicted from the nodes.

调度程序使用 Pod 的“resources.requests”来决定将 Pod 放置在哪个节点上。例如，无法在没有足够空闲（即_non-requested_）资源来覆盖该 Pod 资源请求的节点上调度 Pod。另一方面，`resources.limits` 允许您限制 Pod 的资源消耗，这些资源消耗大大超过了它们各自的请求。一个好的提示是**设置限制等于请求**。将限制设置为远高于请求可能会导致某些节点的 Pod 无法获得请求的资源的情况。这可能会导致节点上的其他应用程序（甚至节点本身）出现故障。 Kubernetes 根据其资源方案为每个 Pod 分配一个 [QoS 类](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod)。然后，K8s 使用 QoS 类来决定应该从节点中驱逐哪些 Pod。

Therefore, you **have to set both requests and limits for both the CPU and memory.** The only thing you can/should omit is the CPU limit if the [Linux kernel version is older than 5.4](https://engineering.indeedblog.com/blog/2019/12/cpu-throttling-regression-fix/) (in the case of EL7/CentOS7, the kernel version must be older than 3.10.0-1062.8.1.el7). 

因此，您**必须同时设置 CPU 和内存的请求和限制。** 如果[Linux 内核版本早于 5.4](https://engineering.indeedblog.com/blog/2019/12/cpu-throttling-regression-fix/)，您唯一可以/应该省略的是 CPU 限制（对于EL7/CentOS7，内核版本必须早于3.10.0-1062.8.1.el7）。

Furthermore, the memory consumption of some applications tends to grow in an unlimited fashion. A good example of that is Redis used for caching or an application that basically runs “on its own”. To limit their impact on other applications on the node, you can (and should) set limits for the amount of memory to be consumed. The only problem with that is the application will be `KILL` ed when this limit is reached. Applications cannot predict/handle this signal, and this will probably prevent them from shutting down correctly. That is why, in addition to Kubernetes limits, we **highly recommend using application-specific mechanisms for limiting memory consumption** so that it does not exceed (or come close to) the amount set in a Pod's `limits.memory` parameter .

此外，某些应用程序的内存消耗往往以无限的方式增长。一个很好的例子是用于缓存的 Redis 或基本上“独立”运行的应用程序。为了限制它们对节点上其他应用程序的影响，您可以（并且应该）设置要消耗的内存量的限制。唯一的问题是当达到这个限制时，应用程序将被 `KILL` 编辑。应用程序无法预测/处理此信号，这可能会阻止它们正确关闭。这就是为什么，除了 Kubernetes 限制之外，我们**强烈建议使用特定于应用程序的机制来限制内存消耗**，这样它就不会超过（或接近）Pod 的 `limits.memory` 参数中设置的数量.

Here is a Redis configuration that can help you with this:

这是一个 Redis 配置，可以帮助您解决这个问题：

```
maxmemory 500mb   # if the amount of data exceeds 500 MB...
maxmemory-policy allkeys-lru   # ...Redis would delete rarely used keys
```

As for Sidekiq, you can use the [Sidekiq worker killer](https://github.com/klaxit/sidekiq-worker-killer):

至于 Sidekiq，你可以使用 [Sidekiq 工人杀手](https://github.com/klaxit/sidekiq-worker-killer)：

```
require 'sidekiq/worker_killer'
Sidekiq.configure_server do |config|
config.server_middleware do |chain|
    # Terminate Sidekiq correctly when it consumes 500 MB
    chain.add Sidekiq::WorkerKiller, max_rss: 500
end
end
```

It is clear that in all these cases that `limits.memory` needs to be higher than the thresholds for triggering the above mechanisms.

很明显，在所有这些情况下，`limits.memory` 需要高于触发上述机制的阈值。

_In the next article, we’ll discuss using VerticalPodAutoscaler to allocate resources automatically._

_在下一篇文章中，我们将讨论使用 VerticalPodAutoscaler 自动分配资源。_

## 7\. Probes

## 7. 探头

In Kubernetes, probes (health checks) are used to determine whether it is possible to switch traffic to the application ( _readiness_ probes) and whether the application needs to be restarted ( _liveness_ probes). They play an important role in updating Deployments and starting new Pods in general.

在 Kubernetes 中，探针（健康检查）用于确定是否可以将流量切换到应用程序（_readiness_probes）以及是否需要重新启动应用程序（_liveness_probes）。它们在更新 Deployment 和启动新 Pod 方面发挥着重要作用。

First of all, we would like to provide a general recommendation for all probe types: **set a high value for the** **`timeoutSeconds` parameter**. A default value of one second is way too low, and it will have a critical impact on readinessProbe & livenessProbe. If `timeoutSeconds` is too low, an increase in the application response time (which often takes place simultaneously for all Pods due to Service load balancing) may either result in these Pods being removed from load balancing (in the case of a failed readiness probe ) or, what's worse, in cascading container restarts (in the case of a failed liveness probe).

首先，我们想为所有探针类型提供一般建议：**为** **`timeoutSeconds` 参数**设置一个高值。一秒的默认值太低了，它会对 readinessProbe 和 livenessProbe 产生严重影响。如果 `timeoutSeconds` 太低，应用程序响应时间的增加（由于服务负载平衡，这通常会同时发生在所有 Pod 上）可能会导致这些 Pod 被从负载平衡中移除（在就绪探测失败的情况下） )，或者更糟糕的是，在级联容器重新启动时（在活性探测失败的情况下）。

### 7.1. Liveness probe

### 7.1 活体探针

In practice, the liveness probe is not as widely used as you may have thought. Its purpose is to restart a container if, for example, the application is frozen. However, in real life, such app deadlocks are an exception rather than the rule. If the application demonstrates partial functionality for some reason (e.g., it cannot restore connection to a database after it has been broken), you have to fix that in the application, rather than “inventing” livenessProbe-based workarounds.

在实践中，活性探针的使用并不像您想象的那么广泛。其目的是在例如应用程序被冻结时重新启动容器。然而，在现实生活中，这样的应用程序死锁是一个例外而不是规则。如果应用程序出于某种原因展示了部分功能（例如，它无法在数据库中断后恢复与数据库的连接），您必须在应用程序中修复它，而不是“发明”基于 livenessProbe 的解决方法。

While you can use livenessProbe to check for these kinds of states, we recommend either **not using livenessProbe by default** or only performing some **basic liveness-testing, such as testing for the TCP connection (remember to set a high timeout value)**. This way, the application will be restarted in response to an apparent deadlock without risking falling into the trap of a loop of restarts (i.e. restarting it won’t help). 

虽然您可以使用 livenessProbe 检查这些类型的状态，但我们建议**默认不使用 livenessProbe** 或仅执行一些**基本的 liveness-testing，例如测试 TCP 连接（请记住设置高超时价值）**。这样，应用程序将重新启动以响应明显的死锁，而不会陷入重新启动循环的陷阱（即重新启动它无济于事）。

Risks related to a poorly configured livenessProbe are serious. In the most common cases, livenessProbe fails due to increased load on the application (it simply cannot make it within the time specified by the timeout parameter) or due to the state of external dependencies that are currently down being checked (directly or indirectly). In the latter case, all the containers will be restarted. In the best case scenario, this would result in nothing, but in the worst case, this would render the application  completely unavailable, probably long-term. Long-term total unavailability of an application (if it has a large number of replicas) may result if most Pods’ containers are restarted within a short time period. Some containers are likely to become READY faster than others, and the entire load will be distributed over this limited number of running containers. That will end up causing livenessProbe timeouts, which will trigger even more restarts.

与配置不当的 livenessProbe 相关的风险非常严重。在最常见的情况下，livenessProbe 失败的原因是应用程序的负载增加（它根本无法在 timeout 参数指定的时间内完成）或由于当前正在检查的外部依赖项的状态（直接或间接）。在后一种情况下，所有容器都将重新启动。在最好的情况下，这不会产生任何结果，但在最坏的情况下，这会使应用程序完全不可用，可能是长期的。如果大多数 Pod 的容器在短时间内重新启动，可能会导致应用程序（如果它有大量副本）长期完全不可用。一些容器可能比其他容器更快地准备好，并且整个负载将分布在这个有限数量的运行容器上。这最终会导致 livenessProbe 超时，这将触发更多的重启。

Also, ensure that livenessProbe does not stop responding if your application has a limit on the number of established connections and that limit has been reached. Usually, you have to dedicate a separate application thread/process to livenessProbe to avoid such problems. For example, if your application has 11 threads (one thread per client), you can limit the number of clients to 10, ensuring that there is an idle thread available for livenessProbe.

此外，如果您的应用程序对已建立的连接数有限制并且已达到该限制，请确保 livenessProbe 不会停止响应。通常，您必须将单独的应用程序线程/进程专用于 livenessProbe 以避免此类问题。例如，如果您的应用程序有 11 个线程（每个客户端一个线程），您可以将客户端数量限制为 10 个，确保有空闲线程可用于 livenessProbe。

And, of course, do not add any external dependency checks to your livenessProbe.

当然，不要向 livenessProbe 添加任何外部依赖项检查。

See [this article](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html) for more information on liveness probe issues and how to prevent them.

请参阅 [本文](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html) 了解有关活动探测问题以及如何预防这些问题的更多信息。

### 7.2. Readiness probe

### 7.2 准备探测

The design of readinessProbe has turned out not to be very successful. readinessProbe combines two different functions:

事实证明，readinessProbe 的设计并不是很成功。 readinessProbe 结合了两个不同的功能：

- it finds out if an application is available during the container start;
- it checks if an application remains available after the container has been successfully started.

- 在容器启动期间确定应用程序是否可用；
- 它检查容器成功启动后应用程序是否仍然可用。

In practice, the first function is required in the vast majority of cases, while the second is only needed as often as the livenessProbe. The poorly configured readinessProbe can cause issues similar to those of livenessProbe. In the worst case, they can also end up causing long-term unavailability for the application.

在实践中，绝大多数情况下都需要第一个函数，而第二个函数只需要和 livenessProbe 一样频繁。配置不当的 readinessProbe 可能会导致类似于 livenessProbe 的问题。在最坏的情况下，它们还可能最终导致应用程序长期不可用。

When readinessProbe fails, the Pod ceases to receive traffic. In most cases, such behavior is of little use, since the traffic is usually distributed more or less evenly between the Pods. Thus, generally, readinessProbe either works everywhere or does not work on a large number of Pods at once. There are situations when such behavior can be useful. However, in my experience, that is for the most part under exceptional cases.

当 readinessProbe 失败时，Pod 停止接收流量。在大多数情况下，这种行为几乎没有用，因为流量通常在 Pod 之间或多或少地均匀分布。因此，通常情况下，readinessProbe 要么在任何地方工作，要么不能同时在大量 Pod 上工作。在某些情况下，这种行为可能有用。但是，根据我的经验，这在很大程度上是在特殊情况下。

Still, readinessProbe comes with another crucial feature: it helps determine when a newly created container can handle the traffic so as not to forward load to an application that isn’t ready yet. This readinessProbe feature, au contraire, is necessary at all times.

尽管如此，readinessProbe 还带有另一个关键特性：它有助于确定新创建的容器何时可以处理流量，以免将负载转发到尚未准备好的应用程序。这个 readinessProbe 功能，相反，在任何时候都是必要的。

In other words, one feature of readinessProbe is in high demand, while the other is not necessary at all. This dilemma was solved with the introduction of startupProbe. It first appeared in Kubernetes 1.16 becoming beta in v1.18 and stable in v1.20. Thus, you're best off **using readinessProbe to check if an application is ready in Kubernetes versions below 1.18, but startupProbe – in Kubernetes versions 1.18 and up.** Then again, you can use readinessProbe in Kubernetes 1.18+ if you have any need to stop traffic to individual Pods after the application has been started.

换句话说，readinessProbe 的一个特性需求量很大，而另一个根本不需要。随着 startupProbe 的引入，这个困境得到了解决。它首次出现在 Kubernetes 1.16 中，在 v1.18 中成为 beta 版，在 v1.20 中稳定。因此，您最好 **使用 readinessProbe 来检查应用程序是否在 Kubernetes 1.18 以下版本中准备就绪，但 startupProbe - 在 Kubernetes 版本 1.18 及更高版本中。**再说一次，如果您有，您可以在 Kubernetes 1.18+ 中使用 readinessProbe在应用程序启动后需要停止到各个 Pod 的流量。

### 7.3. Startup probe 

### 7.3 启动探针

startupProbe checks if an application in the container is ready. Then it marks the current Pod as ready to receive traffic or goes on updating/restarting the Deployment. Unlike readinessProbe, startupProbe stops working after the container has been started. We do not advise using startupProbe for checking external dependencies: its failure would trigger a container restart, which may eventually cause the Pod to go `CrashLoopBackOff`. In this state, the delay between attempts to restart a failed container can be as high as five minutes. It may lead to unnecessary downtime since, despite the application being _ready to be restarted_, the container continues to wait until the end of the `CrashLoopBackOff` period before trying to restart.

startupProbe 检查容器中的应用程序是否准备就绪。然后它将当前 Pod 标记为准备好接收流量或继续更新/重新启动部署。与 readinessProbe 不同，startupProbe 在容器启动后停止工作。我们不建议使用 startupProbe 来检查外部依赖关系：它的失败会触发容器重启，最终可能导致 Pod 进入 `CrashLoopBackOff`。在这种状态下，尝试重新启动失败容器之间的延迟可能高达五分钟。这可能会导致不必要的停机，因为尽管应用程序已_准备好重新启动_，但容器仍会继续等待直到 `CrashLoopBackOff` 周期结束，然后再尝试重新启动。

You should use startupProbe if your application receives traffic and your Kubernetes version is 1.18 or higher.

如果您的应用程序接收到流量并且您的 Kubernetes 版本是 1.18 或更高版本，则应该使用 startupProbe。

Also, note that increasing `failureThreshold` instead of setting `initialDelaySeconds` is the preferred method for configuring the probe. This will allow the container to become available as quickly as possible.

另外，请注意增加 `failureThreshold` 而不是设置 `initialDelaySeconds` 是配置探针的首选方法。这将允许容器尽快变得可用。

## 8\. Checking external dependencies

## 8. 检查外部依赖项

As you know, readinessProbe is often used for checking external dependencies (e.g. databases). While this approach has the right to exist, you'd be well advised to separate your means of checking for external dependencies and your means of checking whether the application in the Pod is running at full capacity (and cutting off the sending of traffic to it is a good idea as well).

如您所知，readinessProbe 通常用于检查外部依赖项（例如数据库）。虽然这种方法有权存在，但建议您将检查外部依赖项的方法与检查 Pod 中的应用程序是否满负荷运行（并切断向其发送流量）的方法分开也是个好主意）。

You can use `initContainers` to check external dependencies before running the main containers’ startupProbe/readinessProbe. It’s pretty clear that in that case, you will no longer need to check external dependencies using readinessProbe. `initContainers` do not require changes to the application code. You do not need to embed additional tools to use them for checking external dependencies in the application containers. Usually, they are reasonably easy to implement:

在运行主容器的 startupProbe/readinessProbe 之前，您可以使用 `initContainers` 检查外部依赖项。很明显，在这种情况下，您将不再需要使用 readinessProbe 检查外部依赖项。 `initContainers` 不需要更改应用程序代码。您不需要嵌入额外的工具来使用它们来检查应用程序容器中的外部依赖关系。通常，它们相当容易实现：

```
       initContainers:
      - name: wait-postgres
        image: postgres:12.1-alpine
        command:
        - sh
        - -ec
        - |
          until (pg_isready -h example.org -p 5432 -U postgres);do
            sleep 1
          done
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 50m
            memory: 50Mi
      - name: wait-redis
        image: redis:6.0.10-alpine3.13
        command:
        - sh
        - -ec
        - |
          until (redis-cli -u redis://redis:6379/0 ping);do
            sleep 1
          done
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 50m
            memory: 50Mi
```

## Complete example

## 完整示例

To sum it up, here is a complete example of the production-grade Deployment of a stateless application that comprises all the recommendations provided above.

总而言之，这是一个无状态应用程序的生产级部署的完整示例，其中包含上面提供的所有建议。

_You will need Kubernetes 1.18 or higher and Ubuntu/Debian-based nodes with kernel version 5.4 or higher._

_您将需要 Kubernetes 1.18 或更高版本以及基于 Ubuntu/Debian 且内核版本为 5.4 或更高版本的节点。_

```
apiVersion: apps/v1
kind: Deployment
metadata:
name: testapp
spec:
replicas: 10
selector:
    matchLabels:
      app: testapp
template:
    metadata:
      labels:
        app: testapp
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchLabels:
                  app: testapp
              topologyKey: kubernetes.io/hostname
      priorityClassName: production-medium
      terminationGracePeriodSeconds: 40
      initContainers:
      - name: wait-postgres
        image: postgres:12.1-alpine
        command:
        - sh
        - -ec
        - |
          until (pg_isready -h example.org -p 5432 -U postgres);do
            sleep 1
          done
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 50m
            memory: 50Mi
      containers:
      - name: backend
        image: my-app-image:1.11.1
        command:
        - run
        - app
        - --trigger-graceful-shutdown-if-memory-usage-is-higher-than
        - 450Mi
        - --timeout-seconds-for-graceful-shutdown
        - 35s
        startupProbe:
          httpGet:
            path: /simple-startup-check-no-external-dependencies
            port: 80
          timeoutSeconds: 7
          failureThreshold: 12
        lifecycle:
          preStop:
            exec:
              ["sh", "-ec", "#command to shutdown gracefully if needed"]
        resources:
          requests:
            cpu: 200m
            memory: 500Mi
          limits:
            cpu: 200m
            memory: 500Mi
```

## In the next article

## 在下一篇

There are several other important topics that need to be addressed, such as `PodDisruptionBudget`, `HorizontalPodAutoscaler`, and `VerticalPodAutoscaler`. We will discuss them in Part 2 of this article **(UPDATE: [Part 2 is out!](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-2/))** — subscribe to our blog below and/or follow [our Twitter](https://twitter.com/flant_com) not to miss it! Please share your best practices for deploying applications (or, if need be, you can correct/supplement the ones discussed above).

还有其他几个重要的主题需要解决，例如 `PodDisruptionBudget`、`HorizontalPodAutoscaler` 和 `VerticalPodAutoscaler`。我们将在本文的第 2 部分中讨论它们**（更新：[第 2 部分已发布！](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-2/))** — 订阅我们下面的博客和/或关注 [我们的 Twitter](https://twitter.com/flant_com)不要错过它！请分享您部署应用程序的最佳实践（或者，如果需要，您可以更正/补充上面讨论的那些)。

## Related posts:

##  相关文章：

- [Comparing Ingress controllers for Kubernetes](https://blog.flant.com/comparing-ingress-controllers-for-kubernetes/ "Comparing Ingress controllers for Kubernetes")
- [Migrating your app to Kubernetes: what to do with files?](https://blog.flant.com/migrating-your-app-to-kubernetes-what-to-do-with-files/ "Migrating your app to Kubernetes: what to do with files?")
- [How we enjoyed upgrading a bunch of Kubernetes clusters from v1.16 to v1.19](https://blog.flant.com/how-we-enjoyed-upgrading-kubernetes-clusters-from-v1-16-to -v1-19/ "How we enjoyed upgrading a bunch of Kubernetes clusters from v1.16 to v1.19")

- [比较 Kubernetes 的 Ingress 控制器](https://blog.flant.com/comparing-ingress-controllers-for-kubernetes/ "比较 Kubernetes 的 Ingress 控制器")
- [将您的应用迁移到 Kubernetes：如何处理文件？](https://blog.flant.com/migrating-your-app-to-kubernetes-what-to-do-with-files/ "迁移您的应用到 Kubernetes：如何处理文件？”)
- [我们如何享受将一堆 Kubernetes 集群从 v1.16 升级到 v1.19](https://blog.flant.com/how-we-enjoyed-upgrading-kubernetes-clusters-from-v1-16-to -v1-19/“我们多么享受将一堆 Kubernetes 集群从 v1.16 升级到 v1.19”)



https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1

