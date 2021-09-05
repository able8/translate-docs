# Kubernetes production best practices

# Kubernetes 生产最佳实践

From: https://learnk8s.io/production-best-practices/

A curated checklist of best practices designed to help you release to production.

精心策划的最佳实践清单，旨在帮助您发布到生产环境

This checklist provides actionable best practices for deploying secure, scalable, and resilient services on Kubernetes.

此清单提供了可操作的最佳实践，用于在 Kubernetes 上部署安全、可扩展和有弹性的服务。

The content is open source and available [in this repository](https://github.com/learnk8s/kubernetes-production-best-practices). If you think there are missing best practices or they are not right, consider submitting an issue.

内容是开源的，可 [在此存储库中](https://github.com/learnk8s/kubernetes-production-best-practices)。如果您认为缺少最佳实践或它们不正确，请考虑提交问题。

Categories 类别

- [1. Application development](https://learnk8s.io/production-best-practices/#application-development)
- [2. Governance](https://learnk8s.io/production-best-practices/#governance)
- [3. Cluster configuration](https://learnk8s.io/production-best-practices/#cluster-configuration)

- [1.应用开发](https://learnk8s.io/production-best-practices/#application-development)
- [2.治理](https://learnk8s.io/production-best-practices/#governance)
- [3.集群配置](https://learnk8s.io/production-best-practices/#cluster-configuration)

# 1. Application development

# 1. 应用程序开发

Best practices for application development on Kubernetes.

在 Kubernetes 上开发应用程序的最佳实践。

## Health checks

## 健康检查

Kubernetes offers two mechanisms to track the lifecycle of your containers and Pods: liveness and readiness probes.

Kubernetes 提供了两种机制来跟踪容器和 Pod 的生命周期：活性探测和就绪探测。

**The readiness probe determines when a container can receive traffic.**

**就绪探针确定容器何时可以接收流量。*

The kubelet executes the checks and decides if the app can receive traffic or not.

kubelet 执行检查并决定应用程序是否可以接收流量。

**The liveness probe determines when a container should be restarted.**

**活性探针确定何时应重新启动容器。**

The kubelet executes the check and decides if the container should be restarted.

kubelet 执行检查并决定是否应该重新启动容器。

**Resources:**

**资源：**

- The official Kubernetes documentation offers some practical advice on how to [configure Liveness, Readiness and Startup Probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).
- [Liveness probes are dangerous](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html) has some information on how to set (or not) dependencies in your readiness probes.

- Kubernetes 官方文档提供了一些关于如何[配置 Liveness、Readiness 和 Startup Probes] 的实用建议
- [Liveness 探测器很危险](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html)有一些关于如何在你的就绪探测器中设置（或不设置)依赖项的信息。

### Containers have Readiness probes

### 容器有就绪探测

> Please note that there's no default value for readiness and liveness.

> 请注意，readiness 和 liveness 没有默认值。

If you don't set the readiness probe, the kubelet assumes that the app is ready to receive traffic as soon as the container starts.

如果您没有设置就绪探测，kubelet 会假设应用程序在容器启动后就准备好接收流量。

If the container takes 2 minutes to start, all the requests to it will fail for those 2 minutes.

如果容器需要 2 分钟才能启动，那么在这 2 分钟内对它的所有请求都将失败。

### Containers crash when there's a fatal error

### 容器在出现致命错误时崩溃

If the application reaches an unrecoverable error, [you should let it crash](https://blog.colinbreck.com/kubernetes-liveness-and-readiness-probes-revisited-how-to-avoid-shooting-yourself-in-the-other-foot/#letitcrash).

如果应用程序遇到不可恢复的错误，[你应该让它崩溃](https://blog.colinbreck.com/kubernetes-liveness-and-readiness-probes-revisited-how-to-avoid-shooting-yourself-in-另一只脚/#letitcrash)。

Examples of such unrecoverable errors are:

- an uncaught exception
- a typo in the code (for dynamic languages)
- unable to load a header or dependency

此类不可恢复错误的示例包括：

- 一个未捕获的异常
- 代码中的错字（对于动态语言）
- 无法加载标题或依赖项

Please note that you should not signal a failing Liveness probe.

请注意，您不应发出失败的 Liveness 探测信号。

Instead, you should immediately exit the process and let the kubelet restart the container.

相反，您应该立即退出进程并让 kubelet 重新启动容器。

### Configure a passive Liveness probe

### 配置一个被动的 Liveness 探针

The Liveness probe is designed to restart your container when it's stuck.

Liveness 探针旨在在您的容器卡住时重新启动它。

Consider the following scenario: if your application is processing an infinite loop, there's no way to exit or ask for help.

考虑以下场景：如果您的应用程序正在处理无限循环，则无法退出或寻求帮助。

When the process is consuming 100% CPU, it won't have time to reply to the (other) Readiness probe checks, and it will be eventually removed from the Service.

当进程消耗 100% CPU 时，它将没有时间回复（其他）Readiness 探针检查，它最终将从服务中删除。

However, the Pod is still registered as an active replica for the current Deployment.

但是，Pod 仍然注册为当前部署的活动副本。

If you don't have a Liveness probe, it stays _Running_ but detached from the Service.

如果您没有 Liveness 探测器，它会保持 _Running_ 状态，但会与服务分离。

In other words, not only is the process not serving any requests, but it is also consuming resources.

换句话说，该进程不仅不处理任何请求，而且还在消耗资源。

_What should you do?_

_你该怎么办？_

1. Expose an endpoint from your app
1. The endpoint always replies with a success response
1. Consume the endpoint from the Liveness probe



1. 从你的应用公开一个端点
1. 端点总是回复成功
1. 从Liveness探针消费端点

Please note that you should not use the Liveness probe to handle fatal errors in your app and request Kubernetes to restart the app.

请注意，您不应使用 Liveness 探针来处理应用程序中的致命错误并请求 Kubernetes 重新启动应用程序。

Instead, you should let the app crash.

相反，您应该让应用程序崩溃。

The Liveness probe should be used as a recovery mechanism only in case the process is not responsive.

只有在进程没有响应的情况下，才应将 Liveness 探针用作恢复机制。

### Liveness probes values aren't the same as the Readiness

### Liveness 探针值与 Readiness 不同

When Liveness and Readiness probes are pointing to the same endpoint, the effects of the probes are combined.

当 Liveness 和 Readiness 探针指向同一端点时，探针的效果会结合在一起。

When the app signals that it's not ready or live, the kubelet detaches the container from the Service and delete it **at the same time**.

当应用程序发出未准备好或未运行的信号时，kubelet 将容器与服务分离并**同时**删除它。

You might notice dropping connections because the container does not have enough time to drain the current connections or process the incoming ones.

您可能会注意到断开连接，因为容器没有足够的时间来耗尽当前连接或处理传入的连接。

You can dig deeper in the following [article that discussed graceful shutdown](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/).

您可以在以下 [讨论正常关机的文章](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/) 中深入挖掘。

## Apps are independent 

## 应用程序是独立的

You might be tempted to signal the readiness of your app only if all of the dependencies such as databases or backend API are also ready.

仅当所有依赖项（例如数据库或后端 API）也准备就绪时，您才可能想要发出应用程序准备就绪的信号。

If the app connects to a database, you might think that returning a failing readiness probe until the database is _ready_ is a good idea — it is not.

如果应用程序连接到数据库，您可能会认为在数据库准备就绪之前返回失败的就绪探测器是一个好主意 - 事实并非如此。

Consider the following scenario: you have one front-end app that depends on a backend API.

考虑以下场景：您有一个依赖后端 API 的前端应用程序。

If the API is flaky (e.g. it's unavailable from time to time due to a bug), the readiness probe fails, and the dependent readiness in the front-end app fail as well.

如果 API 不稳定（例如，由于 bug 有时不可用），则就绪探测失败，并且前端应用程序中的依赖就绪也失败。

And you have downtime.

你有停机时间。

More in general, **a failure in a dependency downstream could propagate to all apps upstream** and eventually, bring down your front-end facing layer as well.

更一般地说，**下游依赖项中的故障可能会传播到上游的所有应用程序**，最终也会导致您的前端层失效。

### The Readiness probes are independent

### 就绪探针是独立的

The readiness probe doesn't include dependencies to services such as:

- databases
- database migrations
- APIs
- third party services

就绪探针不包括对服务的依赖，例如：

- 数据库
- 数据库迁移
- API
- 第三方服务

You can [explore what happens when there're dependencies in the readiness probes in this essay](https://blog.colinbreck.com/kubernetes-liveness-and-readiness-probes-how-to-avoid-shooting-yourself-in-the-foot/#shootingyourselfinthefootwithreadinessprobes).

您可以[探索本文中的就绪探针存在依赖关系时会发生什么](https://blog.colinbreck.com/kubernetes-liveness-and-readiness-probes-how-to-avoid-shooting-yourself-在脚/#shootingyourselfinthefootwithreadinessprobes)。

### The app retries connecting to dependent services

### 应用程序重试连接到依赖服务

When the app starts, it shouldn't crash because a dependency such as a database isn't ready.

当应用程序启动时，它不应崩溃，因为诸如数据库之类的依赖项尚未准备好。

Instead, the app should keep retrying to connect to the database until it succeeds.

相反，应用程序应该不断重试连接到数据库，直到成功。

Kubernetes expects that application components can be started in any order.

Kubernetes 期望应用程序组件可以以任何顺序启动。

When you make sure that your app can reconnect to a dependency such as a database you know you can deliver a more robust and resilient service.

当您确保您的应用程序可以重新连接到依赖项（例如数据库）时，您就知道您可以提供更强大和更有弹性的服务。

## Graceful shutdown

## 优雅关机

When a Pod is deleted, you don't want to terminate all connections abruptly.

当 Pod 被删除时，您不希望突然终止所有连接。

Instead, you should wait for the existing connection to drain and stop processing new ones.

相反，您应该等待现有连接耗尽并停止处理新连接。

Please notice that, when a Pod is terminated, the endpoints for that Pod are removed from the Service.

请注意，当 Pod 终止时，该 Pod 的端点将从服务中删除。

However, it might take some time before component such as kube-proxy or the Ingress controller is notified of the change.

但是，可能需要一些时间才能将更改通知给 kube-proxy 或 Ingress 控制器等组件。

You can find a detail explanation on how graceful shutdown works in [handling client requests correctly with Kubernetes](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/).

您可以在 [使用 Kubernetes 正确处理客户端请求](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/) 中找到有关优雅关闭如何工作的详细说明。

The correct graceful shutdown sequence is:

1. upon receiving SIGTERM
1. the server stops accepting new connections
1. completes all active requests
1. then immediately kills all keepalive connections and
1. the process exits

正确的正常关机顺序是：

1. 收到SIGTERM后
1. 服务器停止接受新连接
1. 完成所有活动请求
1. 然后立即杀死所有保持活动的连接并
1. 进程退出

You can [test that your app gracefully shuts down with this tool: kube-sigterm-test](https://github.com/mikkeloscar/kube-sigterm-test).

您可以[使用此工具测试您的应用是否正常关闭：kube-sigterm-test](https://github.com/mikkeloscar/kube-sigterm-test)。

### The app doesn't shut down on SIGTERM, but it gracefully terminates connections

### 应用程序不会在 SIGTERM 上关闭，但它会优雅地终止连接

It might take some time before a component such as kube-proxy or the Ingress controller is notified of the endpoint changes.

可能需要一些时间才能将端点更改通知给 kube-proxy 或 Ingress 控制器等组件。

Hence, traffic might still flow to the Pod despite it being marked as terminated.

因此，尽管 Pod 被标记为已终止，但流量可能仍会流向该 Pod。

The app should stop accepting new requests on all remaining connections, and close these once the outgoing queue is drained.

应用程序应停止接受所有剩余连接上的新请求，并在传出队列耗尽后关闭这些请求。

If you need a refresher on how endpoints are propagated in your cluster, [read this article on how to handle client requests properly](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/) .

如果您需要复习如何在集群中传播端点，[阅读有关如何正确处理客户端请求的文章](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/) .

### The app still processes incoming requests in the grace period

### 应用程序仍然在宽限期内处理传入的请求

You might want to consider using the container lifecycle events such as [the preStop handler](https://kubernetes.io/docs/tasks/configure-pod-container/attach-handler-lifecycle-event/#define-poststart-and-prestop-handlers) to customize what happened before a Pod is deleted.

您可能需要考虑使用容器生命周期事件，例如 [preStop 处理程序](https://kubernetes.io/docs/tasks/configure-pod-container/attach-handler-lifecycle-event/#define-poststart-and-prestop-handlers) 来自定义删除 Pod 之前发生的事情。

### The CMD in the `Dockerfile` forwards the SIGTERM to the process

### `Dockerfile` 中的 CMD 将 SIGTERM 转发给进程

You can be notified when the Pod is about to be terminated by capturing the SIGTERM signal in your app.

通过在您的应用程序中捕获 SIGTERM 信号，您可以在 Pod 即将终止时收到通知。

You should also pay attention to [forwarding the signal to the right process in your container](https://pracucci.com/graceful-shutdown-of-kubernetes-pods.html).

您还应该注意[将信号转发到容器中的正确进程](https://pracucci.com/graceful-shutdown-of-kubernetes-pods.html)。

### Close all idle keep-alive sockets

### 关闭所有空闲的保持活动的套接字

If the calling app is not closing the TCP connection (e.g. using TCP keep-alive or a connection pool) it will connect to one Pod and not use the other Pods in that Service.

如果调用应用程序没有关闭 TCP 连接（例如使用 TCP 保持连接或连接池），它将连接到一个 Pod，而不使用该服务中的其他 Pod。

_But what happens when a Pod is deleted?_

_但是当一个 Pod 被删除时会发生什么？_

Ideally, the request should go to another Pod.

理想情况下，请求应该转到另一个 Pod。

However, the calling app has a long-lived connection open with the Pod that is about to be terminated, and it will keep using it.

但是，调用应用程序与即将终止的 Pod 建立了一个长期连接，它将继续使用它。

On the other hand, you shouldn't abruptly terminate long-lived connections.

另一方面，您不应该突然终止长期连接。

Instead, you should terminate them before shutting down the app. 

相反，您应该在关闭应用程序之前终止它们。

You can read about keep-alive connections on this article about [gracefully shutting down a Nodejs HTTP server](http://dillonbuchanan.com/programming/gracefully-shutting-down-a-nodejs-http-server/).

您可以在这篇关于[优雅地关闭 Nodejs HTTP 服务器](http://dillonbuchanan.com/programming/gracefully-shutting-down-a-nodejs-http-server/) 的文章中阅读有关保持活动连接的信息。

## Fault tolerance

## 容错

Your cluster nodes could disappear at any time for several reasons:

- a hardware failure of the physical machine
- cloud provider or hypervisor failure
- a kernel panic

由于以下几个原因，您的集群节点可能随时消失：

- 物理机的硬件故障
- 云提供商或管理程序故障
- 内核恐慌

Pods deployed in those nodes are lost too.

部署在这些节点中的 Pod 也会丢失。

Also, there are other scenarios where Pods could be deleted:

- directly deleting a pod (accident)
- draining a node
- removing a pod from a node to permit another Pod to fit on that node

此外，还有其他场景可以删除 Pod：

- 直接删除一个pod（意外）
- 排空一个节点
- 从节点中删除一个 Pod 以允许另一个 Pod 适合该节点

Any of the above scenarios could affect the availability of your app and potentially cause downtime.

上述任何一种情况都可能影响您的应用程序的可用性并可能导致停机。

You should protect from a scenario where all of your Pods are made unavailable, and you aren't able to serve live traffic.

您应该防止出现所有 Pod 都不可用且无法提供实时流量的情况。

### Run more than one replica for your Deployment

### 为您的部署运行多个副本

Never run a single Pod individually.

永远不要单独运行单个 Pod。

Instead consider deploying your Pod as part of a Deployment, DaemonSet, ReplicaSet or StatefulSet.

而是考虑将您的 Pod 部署为 Deployment、DaemonSet、ReplicaSet 或 StatefulSet 的一部分。

[Running more than one instance your of your Pods guarantees that deleting a single Pod won't cause downtime](https://cloudmark.github.io/Node-Management-In-GKE/#replicas).

[运行多个 Pod 实例可确保删除单个 Pod 不会导致停机](https://cloudmark.github.io/Node-Management-In-GKE/#replicas)。

### Avoid Pods being placed into a single node

### 避免将 Pod 放置到单个节点中

**Even if you run several copies of your Pods, there are no guarantees that losing a node won't take down your service.**

**即使您运行多个 Pod 副本，也不能保证丢失节点不会中断您的服务。**

Consider the following scenario: you have 11 replicas on a single cluster node.

考虑以下场景：您在单个集群节点上有 11 个副本。

If the node is made unavailable, the 11 replicas are lost, and you have downtime.

如果该节点不可用，则 11 个副本将丢失，并且您将停机。

[You should apply anti-affinity rules to your Deployments so that Pods are spread in all the nodes of your cluster](https://cloudmark.github.io/Node-Management-In-GKE/#pod-anti-affinity-rules).

[您应该将反关联规则应用于您的部署，以便 Pod 分布在集群的所有节点中](https://cloudmark.github.io/Node-Management-In-GKE/#pod-anti-affinity-规则)。

The [inter-pod affinity and anti-affinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity) documentation describe how you can you could change your Pod to be located (or not) in the same node.

[pod 间亲和性和反亲和性](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity) 文档描述了如何您可以将 Pod 更改为位于（或不位于)同一节点中。

### Set Pod disruption budgets

### 设置 Pod 中断预算

When a node is drained, all the Pods on that node are deleted and rescheduled.

当一个节点耗尽时，该节点上的所有 Pod 都会被删除并重新调度。

_But what if you are under heavy load and you can't lose more than 50% of your Pods?_

_但是，如果您的负载很重并且丢失的 Pod 不能超过 50% 怎么办？_

The drain event could affect your availability.

排空事件可能会影响您的可用性。

To protect the Deployments from unexpected events that could take down several Pods at the same time, you can define Pod Disruption Budget.

为了保护部署免受可能同时关闭多个 Pod 的意外事件的影响，您可以定义 Pod 中断预算。

Imagine saying: _"Kubernetes, please make sure that there are always at least 5 Pods running for my app"._

想象一下：_“Kubernetes，请确保始终至少有 5 个 Pod 在为我的应用程序运行”。_

Kubernetes will prevent the drain event if the final state results in less than 5 Pods for that Deployment.

如果最终状态导致该部署的 Pod 少于 5 个，Kubernetes 将阻止耗尽事件。

The official documentation is an excellent place to start to understand [Pod Disruption Budgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/).

官方文档是开始了解 [Pod Disruption Budgets](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/) 的绝佳起点。

## Resources utilisation

## 资源利用率

You can think about the Kubernetes as a skilled Tetris player.

您可以将 Kubernetes 视为熟练的俄罗斯方块玩家。

Docker containers are the blocks; servers are the boards, and the scheduler is the player.

Docker 容器是块；服务器是板，调度器是播放器。

![Kubernetes is the best Tetris player](tetris.svg)

To maximise the efficiency of the scheduler, you should share with Kubernetes details such as resource utilisation, workload priorities and overheads.

为了最大限度地提高调度程序的效率，您应该与 Kubernetes 共享详细信息，例如资源利用率、工作负载优先级和开销。

### Set memory limits and requests for all containers

### 为所有容器设置内存限制和请求

Resource limits are used to constrain how much CPU and memory your containers can utilise and are set using the resources property of a `containerSpec`.

资源限制用于限制您的容器可以使用多少 CPU 和内存，并使用 `containerSpec` 的资源属性进行设置。

The scheduler uses those as one of metrics to decide which node is best suited for the current Pod.

调度程序使用这些作为指标之一来决定哪个节点最适合当前 Pod。

A container without a memory limit has memory utilisation of zero — according to the scheduler.

根据调度程序，没有内存限制的容器的内存利用率为零。

An unlimited number of Pods if schedulable on any nodes leading to resource overcommitment and potential node (and kubelet) crashes.

如果可在任何节点上调度会导致资源过度使用和潜在节点（和 kubelet）崩溃，则 Pod 数量不受限制。

The same applies to CPU limits.

这同样适用于 CPU 限制。

_But should you always set limits and requests for memory and CPU?_

_但是你应该总是为内存和 CPU 设置限制和请求吗？_

Yes and no.

是和否。

If your process goes over the memory limit, the process is terminated.

如果您的进程超过内存限制，则该进程将终止。

Since CPU is a compressible resource, if your container goes over the limit, the process is throttled.

由于 CPU 是一种可压缩资源，如果您的容器超过限制，则进程会受到限制。

Even if it could have used some of the CPU that was available at that moment.

即使它可以使用当时可用的一些 CPU。

**[CPU limits are hard.](https://www.reddit.com/r/kubernetes/comments/cmp7jj/multithreading_in_a_container_with_limited/ew52fcj/)**

**[CPU 限制很难。](https://www.reddit.com/r/kubernetes/comments/cmp7jj/multithreading_in_a_container_with_limited/ew52fcj/)**

If you wish to dig deeper into CPU and memory limits you should check out the following articles: 

如果您想更深入地了解 CPU 和内存限制，您应该查看以下文章：

- [Understanding resource limits in kubernetes: memory](https://medium.com/@betz.mark/understanding-resource-limits-in-kubernetes-memory-6b41e9a955f9)
- [Understanding resource limits in kubernetes: cpu time](https://medium.com/@betz.mark/understanding-resource-limits-in-kubernetes-cpu-time-9eff74d3161b)

- [了解 kubernetes 中的资源限制：内存](https://medium.com/@betz.mark/understanding-resource-limits-in-kubernetes-memory-6b41e9a955f9)
- [了解kubernetes中的资源限制：cpu时间](https://medium.com/@betz.mark/understanding-resource-limits-in-kubernetes-cpu-time-9eff74d3161b)

> Please note that if you are not sure what should be the _right_ CPU or memory limit, you can use the [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) in Kubernetes with the recommendation mode turned on. The autoscaler profiles your app and recommends limits for it.

> 请注意，如果您不确定应该是什么_正确的 CPU 或内存限制，您可以使用 [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) 在 Kubernetes 中打开推荐模式。自动调节程序分析您的应用并为其推荐限制。

### Set CPU request to 1 CPU or below

### 将 CPU 请求设置为 1 个 CPU 或以下

Unless you have computational intensive jobs, [it is recommended to set the request to 1 CPU or below](https://www.youtube.com/watch?v=xjpHggHKm78).

除非您有计算密集型作业，[建议将请求设置为 1 个 CPU 或以下](https://www.youtube.com/watch?v=xjpHggHKm78)。

### Disable CPU limits — unless you have a good use case

### 禁用 CPU 限制——除非你有一个很好的用例

CPU is measured as CPU timeunits per timeunit.

CPU 以每个时间单位的 CPU 时间单位来衡量。

`cpu: 1` means 1 CPU second per second.

`cpu: 1` 表示每秒 1 个 CPU。

If you have 1 thread, you can't consume more than 1 CPU second per second.

如果您有 1 个线程，则每秒不能消耗超过 1 个 CPU。

If you have 2 threads, you can consume 1 CPU second in 0.5 seconds.

如果你有 2 个线程，你可以在 0.5 秒内消耗 1 个 CPU 秒。

8 threads can consume 1 CPU second in 0.125 seconds.

8 个线程可以在 0.125 秒内消耗 1 个 CPU 秒。

After that, your process is throttled.

之后，您的进程将受到限制。

If you're not sure about what's the best settings for your app, it's better not to set the CPU limits.

如果您不确定应用程序的最佳设置是什么，最好不要设置 CPU 限制。

If you wish to learn more, [this article digs deeper in CPU requests and limits](https://medium.com/@betz.mark/understanding-resource-limits-in-kubernetes-cpu-time-9eff74d3161b).

如果您想了解更多信息，[本文深入挖掘 CPU 请求和限制](https://medium.com/@betz.mark/understanding-resource-limits-in-kubernetes-cpu-time-9eff74d3161b)。

### The namespace has a LimitRange

### 命名空间有一个 LimitRange

If you think you might forget to set memory and CPU limits, you should consider using a LimitRange object to define the standard size for a container deployed in the current namespace.

如果您认为您可能忘记设置内存和 CPU 限制，您应该考虑使用 LimitRange 对象来定义部署在当前命名空间中的容器的标准大小。

[The official documentation about LimitRange](https://kubernetes.io/docs/concepts/policy/limit-range/) is an excellent place to start.

[关于 LimitRange 的官方文档](https://kubernetes.io/docs/concepts/policy/limit-range/) 是一个很好的起点。

### Set an appropriate Quality of Service (QoS) for Pods

### 为 Pod 设置适当的服务质量 (QoS)

When a node goes into an overcommitted state (i.e. using too many resources) Kubernetes tries to evict some of the Pod in that Node.

当节点进入过度使用状态（即使用过多资源）时，Kubernetes 会尝试驱逐该节点中的某些 Pod。

Kubernetes ranks and evicts the Pods according to a well-defined logic.

Kubernetes 根据明确定义的逻辑对 Pod 进行排名和驱逐。

You can find more about [configuring the quality of service for your Pods](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/) on the official documentation.

您可以在官方文档中找到更多关于[为您的 Pod 配置服务质量](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/) 的信息。

## Tagging resources

## 标记资源

Labels are the mechanism you use to organize Kubernetes objects.

标签是用于组织 Kubernetes 对象的机制。

A label is a key-value pair without any pre-defined meaning.

标签是没有任何预定义含义的键值对。

They can be applied to all resources in your cluster from Pods to Service, Ingress manifests, Endpoints, etc.

它们可以应用于集群中的所有资源，从 Pod 到服务、入口清单、端点等。

You can use labels to categorize resources by purpose, owner, environment, or other criteria.

您可以使用标签按用途、所有者、环境或其他条件对资源进行分类。

So you could choose a label to tag a Pod in an environment such as "this pod is running in production" or "the payment team owns that Deployment".

因此，您可以选择一个标签来标记环境中的 Pod，例如“此 Pod 正在生产中运行”或“支付团队拥有该部署”。

You can also omit labels altogether.

您也可以完全省略标签。

However, you might want to consider using labels to cover the following categories:

- technical labels such as the environment
- labels for automation
- label related to your business such as cost-centre allocation
- label related to security such as compliance requirements

但是，您可能需要考虑使用标签来涵盖以下类别：

- 技术标签，如环境
- 自动化标签
- 与您的业务相关的标签，例如成本中心分配
- 与安全相关的标签，例如合规性要求

### Resources have technical labels defined

### 资源定义了技术标签

You could tag your Pods with:

- `name`, the name of the application such "User API"
- `instance`, a unique name identifying the instance of an application (you could use the container image tag)
- `version`, the current version of the appl (an incremental counter)
- `component`, the component within the architecture such as "API" or "database"
- `part-of`, the name of a higher-level application this one is part of such as "payment gateway"
- `managed-by`, the tool being used to manage the operation of an application such as "kubectl" or "Helm"

您可以使用以下标签标记您的 Pod：

- `name`，应用程序的名称，例如“用户 API”
- `instance`，一个标识应用程序实例的唯一名称（您可以使用容器图像标签）
- `version`，应用程序的当前版本（增量计数器）
- `component`，架构内的组件，例如“API”或“数据库”
- `part-of`，这是一个更高级别的应用程序的名称，例如“支付网关”
- `managed-by`，用于管理应用程序操作的工具，例如“kubectl”或“Helm”

Here's an example on how you could use such labels in a Deployment:

```yaml|highlight=6-11,20-24|title=deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deployment
  labels:
    app.kubernetes.io/name: user-api
    app.kubernetes.io/instance: user-api-5fa65d2
    app.kubernetes.io/version: "42"
    app.kubernetes.io/component: api
    app.kubernetes.io/part-of: payment-gateway
    app.kubernetes.io/managed-by: kubectl
spec:
  replicas: 3
  selector:
    matchLabels:
      application: my-app
  template:
    metadata:
      labels:
        app.kubernetes.io/name: user-api
        app.kubernetes.io/instance: user-api-5fa65d2
        app.kubernetes.io/version: "42"
        app.kubernetes.io/component: api
        app.kubernetes.io/part-of: payment-gateway
    spec:
      containers:
      - name: app
        image: myapp
```




Those labels are [recommended by the official documentation](https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/).

这些标签是[官方文档推荐的](https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/)。

> Please not that you're recommended to tag **all resources**.

> 请不要建议您标记**所有资源**。

### Resources have business labels defined

### 资源定义了业务标签

You could tag your Pods with:

- `owner`, used to identify who is responsible for the resource
- `project`, used to determine the project that the resource belongs to
- `business-unit`, used to identify the cost centre or business unit associated with a resource; typically for cost allocation and tracking

您可以使用以下标签标记您的 Pod：

- `owner`，用于标识谁对资源负责
- `project`，用于确定资源所属的项目
- `business-unit`，用于标识与资源关联的成本中心或业务单位；通常用于成本分配和跟踪

Here's an example on how you could use such labels in a Deployment:

```yaml|highlight=6-8,17-19|title=deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deployment
  labels:
    owner: payment-team
    project: fraud-detection
    business-unit: "80432"
spec:
  replicas: 3
  selector:
    matchLabels:
      application: my-app
  template:
    metadata:
      labels:
        owner: payment-team
        project: fraud-detection
        business-unit: "80432"
    spec:
      containers:
      - name: app
        image: myapp
```


You can explore labels and [tagging for resources on the AWS tagging strategy page](https://aws.amazon.com/answers/account-management/aws-tagging-strategies/).

您可以探索标签和 [AWS 标记策略页面上的资源标记](https://aws.amazon.com/answers/account-management/aws-tagging-strategies/)。

The article isn't specific to Kubernetes but explores some of the most common strategies for tagging resources.

本文并非专门针对 Kubernetes，而是探讨了一些最常见的资源标记策略。

> Please not that you're recommended to tag **all resources**.

> 请不要建议您标记**所有资源**。

### Resources have security labels defined

### 资源定义了安全标签

You could tag your Pods with:

- `confidentiality`, an identifier for the specific data-confidentiality level a resource supports
- `compliance`, an identifier for workloads designed to adhere to specific compliance requirements

您可以使用以下标签标记您的 Pod：

- `confidentiality`，资源支持的特定数据机密级别的标识符
- `compliance`，工作负载的标识符，旨在遵守特定的合规性要求

Here's an example on how you could use such labels in a Deployment:

```yaml|highlight=6-11,20-24|title=deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deployment
  labels:
    confidentiality: official
    compliance: pci
spec:
  replicas: 3
  selector:
    matchLabels:
      application: my-app
  template:
    metadata:
      labels:
        confidentiality: official
        compliance: pci
    spec:
      containers:
      - name: app
        image: myapp
```


You can explore label and [tagging for resources on the AWS tagging strategy page](https://aws.amazon.com/answers/account-management/aws-tagging-strategies/).

您可以探索标签和 [AWS 标记策略页面上的资源标记](https://aws.amazon.com/answers/account-management/aws-tagging-strategies/)。

The article isn't specific to Kubernetes but explores some of the most common strategies for tagging resources.

本文并非专门针对 Kubernetes，而是探讨了一些最常见的资源标记策略。

> Please not that you're recommended to tag **all resources**.

> 请不要建议您标记**所有资源**。

## Logging

## 记录

Application logs can help you understand what is happening inside your app.

应用程序日志可以帮助您了解应用程序内部发生的情况。

The logs are particularly useful for debugging problems and monitoring app activity.

日志对于调试问题和监控应用程序活动特别有用。

### The application logs to `stdout` and `stderr`

### 应用程序记录到 `stdout` 和 `stderr`

There are two logging strategies: _passive_ and _active_.

有两种日志记录策略：_passive_ 和 _active_。

Apps that use passive logging are unaware of the logging infrastructure and log messages to standard outputs.

使用被动日志记录的应用程序不知道日志记录基础结构并将消息记录到标准输出。

This best practice is part of [the twelve-factor app](https://12factor.net/logs).

此最佳实践是 [十二要素应用程序](https://12factor.net/logs) 的一部分。

In active logging, the app makes network connections to intermediate aggregators, sends data to third-party logging services, or writes directly to a database or index.

在主动日志记录中，应用程序与中间聚合器建立网络连接，将数据发送到第三方日志记录服务，或直接写入数据库或索引。

Active logging is considered an antipattern, and it should be avoided.

主动日志记录被认为是一种反模式，应该避免使用。

### Avoid sidecars for logging (if you can)

### 避免使用 sidecar 进行日志记录（如果可以的话）

If you wish to [apply log transformations to an application with a non-standard log event model](https://rclayton.silvrback.com/container-services-logging-with-docker#effective-logging-infrastructure), you may want to use a sidecar container.

如果您希望 [将日志转换应用于具有非标准日志事件模型的应用程序](https://rclayton.silvrback.com/container-services-logging-with-docker#effective-logging-infrastructure)，您可以想使用边车容器。

With a sidecar container, you can normalise the log entries before they are shipped elsewhere.

使用 sidecar 容器，您可以在将日志条目发送到其他地方之前对其进行规范化。

For example, you may want to transform Apache logs into Logstash JSON format before shipping it to the logging infrastructure.

例如，您可能希望在将 Apache 日志传送到日志记录基础架构之前将其转换为 Logstash JSON 格式。

However, if you have control over the application, you could output the right format, to begin with.

但是，如果您可以控制应用程序，则可以从一开始就输出正确的格式。

You could save on running an extra container for each Pod in your cluster.

您可以节省为集群中的每个 Pod 运行一个额外的容器。

## Scaling

## 缩放

### Containers do not store any state in their local filesystem

### 容器不在其本地文件系统中存储任何状态

Containers have a local filesystem and you might be tempted to use it for persisting data.

容器有一个本地文件系统，你可能想用它来持久化数据。

However, storing persistent data in a container's local filesystem prevents the encompassing Pod from being scaled horizontally (that is, by adding or removing replicas of the Pod). 

但是，将持久数据存储在容器的本地文件系统中会阻止包含的 Pod 被水平扩展（即，通过添加或删除 Pod 的副本）。

This is because, by using the local filesystem, each container maintains its own "state", which means that the states of Pod replicas may diverge over time. This results in inconsistent behaviour from the user's point of view (for example, a specific piece of user information is available when the request hits one Pod, but not when the request hits another Pod).

这是因为，通过使用本地文件系统，每个容器维护自己的“状态”，这意味着 Pod 副本的状态可能会随着时间的推移而发散。这会导致从用户的角度来看不一致的行为（例如，当请求命中一个 Pod 时，特定的用户信息可用，但当请求命中另一个 Pod 时不可用）。

Instead, any persistent information should be saved at a central place outside the Pods. For example, in a PersistentVolume in the cluster, or even better in some storage service outside the cluster.

相反，任何持久信息都应该保存在 Pod 之外的中心位置。例如在集群中的一个 PersistentVolume 中，或者在集群外的一些存储服务中甚至更好。

### Use the Horizontal Pod Autoscaler for apps with variable usage patterns

### 将 Horizontal Pod Autoscaler 用于具有可变使用模式的应用程序

The [Horizontal Pod Autoscaler (HPA)](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) is a built-in Kubernetes feature that monitors your application and automatically adds or removes Pod replicas based on the current usage.

[Horizontal Pod Autoscaler (HPA)](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) 是 Kubernetes 的一项内置功能，可监控您的应用程序并自动添加或删除 Pod基于当前使用情况的副本。

Configuring the HPA allows your app to stay available and responsive under any traffic conditions, including unexpected spikes.

配置 HPA 可让您的应用在任何流量条件下（包括意外峰值）保持可用和响应。

To configure the HPA to autoscale your app, you have to create a [HorizontalPodAutoscaler](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#horizontalpodautoscaler-v1-autoscaling) resource, which defines what metric to monitor for your app.

要将 HPA 配置为自动缩放您的应用程序，您必须创建一个 [HorizontalPodAutoscaler](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#horizontalpodautoscaler-v1-autoscaling) 资源，其中定义要为您的应用监控的指标。

The HPA can monitor either built-in resource metric (CPU and memory usage of your Pods) or custom metrics. In the case of custom metrics, you are also responsible for collecting and exposing these metrics, which you can do, for example, with [Prometheus](https://prometheus.io/) and the [Prometheus Adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter).

HPA 可以监控内置资源指标（Pod 的 CPU 和内存使用情况）或自定义指标。对于自定义指标，您还负责收集和公开这些指标，例如，您可以使用 [Prometheus](https://prometheus.io/) 和 [Prometheus Adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter)。

### Don't use the Vertical Pod Autoscaler while it's still in beta

### 不要在仍处于测试阶段的 Vertical Pod Autoscaler 使用

Analogous to the [Horizontal Pod Autoscaler (HPA)](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/), there exists the [Vertical Pod Autoscaler (VPA)](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler).

类似于[Horizontal Pod Autoscaler (HPA)](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)，存在[Vertical Pod Autoscaler (VPA)](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)。

The VPA can automatically adapt the resource requests and limits of your Pods so that when a Pod needs more resources, it can get them (increasing/decreasing the resources of a single Pod is called _vertical scaling_, as opposed to _horizontal scaling_, which means increasing /decreasing the number of replicas of a Pod).

VPA 可以自动调整你的 Pod 的资源请求和限制，以便当一个 Pod 需要更多资源时，它可以得到它们（增加/减少单个 Pod 的资源称为_垂直缩放_，而不是_水平缩放_，这意味着增加/减少单个 Pod 的资源） /减少 Pod 的副本数）。

This can be useful for scaling applications that can't be scaled horizontally.

这对于扩展无法水平扩展的应用程序非常有用。

However, the VPA is curently in beta and it has [some known limitations](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#limitations-of-beta-version) (for example , scaling a Pod by changing its resource requirements, requires the Pod to be killed and restarted).

然而，VPA 目前处于测试阶段，它有 [一些已知的限制](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#limitations-of-beta-version)（例如，通过更改 Pod 的资源需求来扩展 Pod，需要杀死并重新启动 Pod)。

Given these limitations, and the fact that most applications on Kubernetes can be scaled horizontally anyway, it is recommended to not use the VPA in production (at least until there is a stable version).

鉴于这些限制，以及 Kubernetes 上的大多数应用程序无论如何都可以水平扩展的事实，建议不要在生产中使用 VPA（至少在有稳定版本之前）。

### Use the Cluster Autoscaler if you have highly varying workloads

### 如果您的工作负载变化很大，请使用 Cluster Autoscaler

The [Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) is another type of "autoscaler" (besides the [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) and [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)).

[Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) 是另一种类型的“autoscaler”（除了 [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) 和 [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler))。

The Cluster Autoscaler can automatically scale the size of your cluster by adding or removing worker nodes.

Cluster Autoscaler 可以通过添加或删除工作节点来自动扩展集群的大小。

A scale-up operation happens when a Pod fails to be scheduled because of insufficient resources on the existing worker nodes. In this case, the Cluster Autoscaler creates a new worker node, so that the Pod can be scheduled. Similarly, when the utilisation of the existing worker nodes is low, the Cluster Autoscaler can scale down by evicting all the workloads from one of the worker nodes and removing it.

当 Pod 由于现有工作节点上的资源不足而无法调度时，会发生纵向扩展操作。在这种情况下，Cluster Autoscaler 会创建一个新的工作节点，以便可以调度 Pod。类似地，当现有工作节点的利用率较低时，Cluster Autoscaler 可以通过从其中一个工作节点驱逐所有工作负载并将其删除来缩小规模。

Using the Cluster Autoscaler makes sense for highly variable workloads, for example, when the number of Pods may multiply in a short time, and then go back to the previous value. In such scenarios, the Cluster Autoscaler allows you to meet the demand spikes without wasting resources by overprovisioning worker nodes. 

使用 Cluster Autoscaler 对高度可变的工作负载很有意义，例如，当 Pod 的数量可能在短时间内成倍增加，然后又回到之前的值时。在这种情况下，Cluster Autoscaler 允许您通过过度配置工作节点来满足需求高峰，而不会浪费资源。

However, if your workloads do not vary so much, it may not be worth to set up the Cluster Autoscaler, as it may never be triggered. If your workloads grow slowly and monotonically, it may be enough to monitor the utilisations of your existing worker nodes and add an additional worker node manually when they reach a critical value.

但是，如果您的工作负载变化不大，则可能不值得设置 Cluster Autoscaler，因为它可能永远不会被触发。如果您的工作负载缓慢且单调地增长，监控现有工作节点的利用率并在达到临界值时手动添加额外的工作节点可能就足够了。

## 2. Configuration and secrets

## 配置和秘密

### Externalise all configuration

### 外部化所有配置

Configuration should be maintained outside the application code.

配置应该在应用程序代码之外维护。

This has several benefits. First, changing the configuration does not require recompiling the application. Second, the configuration can be updated when the application is running. Third, the same code can be used in different environments.

这有几个好处。首先，更改配置不需要重新编译应用程序。其次，可以在应用程序运行时更新配置。第三，相同的代码可以在不同的环境中使用。

In Kubernetes, the configuration can be saved in ConfigMaps, which can then be mounted into containers as volumes are passed in as environment variables.

在 Kubernetes 中，配置可以保存在 ConfigMaps 中，然后可以将其挂载到容器中，因为卷作为环境变量传入。

Save only non-sensitive configuration in ConfigMaps. For sensitive information (such as credentials), use the Secret resource.

在 ConfigMaps 中仅保存非敏感配置。对于敏感信息（例如凭据），请使用 Secret 资源。

### Mount Secrets as volumes, not enviroment variables

### 将 Secrets 挂载为卷，而不是环境变量

The content of Secret resources should be mounted into containers as volumes rather than passed in as environment variables.

Secret 资源的内容应该作为卷挂载到容器中，而不是作为环境变量传入。

This is to prevent that the secret values appear in the command that was used to start the container, which may be inspected by individuals that shouldn't have access to the secret values.

这是为了防止秘密值出现在用于启动容器的命令中，这可能会被不应访问秘密值的个人检查。



# Cluster configuration

# 集群配置

Cluster configuration best practices.

集群配置最佳实践。

## Approved Kubernetes configuration

## 批准的 Kubernetes 配置

Kubernetes is flexible and can be configured in several different ways.

Kubernetes 很灵活，可以通过多种不同的方式进行配置。

But how do you know what's the recommended configuration for your cluster?

但是您如何知道集群的推荐配置是什么？

The best option is to compare your cluster with a standard reference.

最好的选择是将您的集群与标准参考进行比较。

In the case of Kubernetes, the reference is the Centre for Internet Security (CIS) benchmark.

对于 Kubernetes，参考是互联网安全中心 (CIS) 基准。

### The cluster passes the CIS benchmark

### 集群通过CIS基准测试

The Center for Internet Security provides several guidelines and benchmark tests for best practices in securing your code.

互联网安全中心为保护代码的最佳实践提供了一些指南和基准测试。

They also maintain a benchmark for Kubernetes which you can [download from the official website](https://www.cisecurity.org/benchmark/kubernetes/).

他们还为 Kubernetes 维护了一个基准，您可以[从官方网站下载](https://www.cisecurity.org/benchmark/kubernetes/)。

While you can read the lengthy guide and manually check if your cluster is compliant, an easier way is to download and execute [`kube-bench`](https://github.com/aquasecurity/kube-bench).

虽然您可以阅读冗长的指南并手动检查您的集群是否合规，但更简单的方法是下载并执行 [`kube-bench`](https://github.com/aquasecurity/kube-bench)。

[`kube-bench`](https://github.com/aquasecurity/kube-bench) is a tool designed to automate the CIS Kubernetes benchmark and report on misconfigurations in your cluster.

[`kube-bench`](https://github.com/aquasecurity/kube-bench) 是一种工具，旨在自动化 CIS Kubernetes 基准测试并报告集群中的错误配置。

Example output:

```terminal|title=bash
[INFO] 1 Master Node Security Configuration
[INFO] 1.1 API Server
[WARN] 1.1.1 Ensure that the --anonymous-auth argument is set to false (Not Scored)
[PASS] 1.1.2 Ensure that the --basic-auth-file argument is not set (Scored)
[PASS] 1.1.3 Ensure that the --insecure-allow-any-token argument is not set (Not Scored)
[PASS] 1.1.4 Ensure that the --kubelet-https argument is set to true (Scored)
[PASS] 1.1.5 Ensure that the --insecure-bind-address argument is not set (Scored)
[PASS] 1.1.6 Ensure that the --insecure-port argument is set to 0 (Scored)
[PASS] 1.1.7 Ensure that the --secure-port argument is not set to 0 (Scored)
[FAIL] 1.1.8 Ensure that the --profiling argument is set to false (Scored)
```


> Please note that it is not possible to inspect the master nodes of managed clusters such as GKE, EKS and AKS, using `kube-bench`. The master nodes are controlled and managed by the cloud provider.

> 请注意，无法使用 `kube-bench` 检查托管集群（例如 GKE、EKS 和 AKS）的主节点。主节点由云提供商控制和管理。

### Disable metadata cloud providers metada API

### 禁用元数据云提供程序元数据 API

Cloud platforms (AWS, Azure, GCE, etc.) often expose metadata services locally to instances.

云平台（AWS、Azure、GCE 等）通常在本地向实例公开元数据服务。

By default, these APIs are accessible by pods running on an instance and can contain cloud credentials for that node, or provisioning data such as kubelet credentials.

默认情况下，这些 API 可由在实例上运行的 pod 访问，并且可以包含该节点的云凭据，或供应数据，例如 kubelet 凭据。

These credentials can be used to escalate within the cluster or to other cloud services under the same account.

这些凭据可用于在集群内升级或升级到同一帐户下的其他云服务。

### Restrict access to alpha or beta features

### 限制访问 alpha 或 beta 功能

Alpha and beta Kubernetes features are in active development and may have limitations or bugs that result in security vulnerabilities.

Alpha 和 Beta Kubernetes 功能正在积极开发中，可能存在导致安全漏洞的限制或错误。

Always assess the value an alpha or beta feature may provide against the possible risk to your security posture.

始终评估 alpha 或 beta 功能可能为您的安全状况可能带来的风险提供的价值。

When in doubt, disable features you do not use.

如有疑问，请禁用您不使用的功能。

## Authentication

##  验证

When you use `kubectl`, you authenticate yourself against the kube-api server component.

当您使用 `kubectl` 时，您可以针对 kube-api 服务器组件对自己进行身份验证。

Kubernetes supports different authentication strategies:

- **Static Tokens**: are difficult to invalidate and should be avoided
- **Bootstrap Tokens**: same as static tokens above 

Kubernetes 支持不同的认证策略：

- **静态令牌**：难以失效，应避免
- **Bootstrap Tokens**：与上面的静态令牌相同

- **Basic Authentication** transmits credentials over the network in cleartext
- **X509 client certs** requires renewing and redistributing client certs regularly
- **Service Account Tokens** are the preferred authentication strategy for applications and workloads running in the cluster
- **OpenID Connect (OIDC) Tokens**: best authentication strategy for end-users as OIDC integrates with your identity provider such as AD, AWS IAM, GCP IAM, etc.

- **基本身份验证**以明文形式通过网络传输凭据
- **X509 客户端证书** 需要定期更新和重新分发客户端证书
- **服务帐户令牌**是集群中运行的应用程序和工作负载的首选身份验证策略
- **OpenID Connect (OIDC) 令牌**：针对最终用户的最佳身份验证策略，因为 OIDC 与您的身份提供商（例如 AD、AWS IAM、GCP IAM 等）集成。

You can learn about the strategies in more detail [in the official documentation](https://kubernetes.io/docs/reference/access-authn-authz/authentication/).

您可以[在官方文档中](https://kubernetes.io/docs/reference/access-authn-authz/authentication/) 更详细地了解这些策略。

### Use OpenID (OIDC) tokens as a user authentication strategy

### 使用 OpenID (OIDC) 令牌作为用户身份验证策略

Kubernetes supports various authentication methods, including OpenID Connect (OIDC).

Kubernetes 支持各种身份验证方法，包括 OpenID Connect (OIDC)。

OpenID Connect allows single sign-on (SSO) such as your Google Identity to connect to a Kubernetes cluster and other development tools.

OpenID Connect 允许单点登录 (SSO)（例如您的 Google 身份）连接到 Kubernetes 集群和其他开发工具。

You don't need to remember or manage credentials separately.

您无需单独记住或管理凭据。

You could have several clusters connect to the same OpenID provider.

您可以将多个集群连接到同一个 OpenID 提供程序。

You can [learn more about the OpenID connect in Kubernetes](https://thenewstack.io/kubernetes-single-sign-one-less-identity/) in this article.

您可以在本文中[了解有关 Kubernetes 中 OpenID 连接的更多信息](https://thenewstack.io/kubernetes-single-sign-one-less-identity/)。

## Role-Based Access Control (RBAC)

## 基于角色的访问控制 (RBAC)

Role-Based Access Control (RBAC) allows you to define policies on how to access resources in your cluster.

基于角色的访问控制 (RBAC) 允许您定义有关如何访问集群中资源的策略。

### ServiceAccount tokens are for applications and controllers **only**

### ServiceAccount 令牌用于应用程序和控制器**仅**

Service Account Tokens should not be used for end-users trying to interact with Kubernetes clusters, but they are the preferred authentication strategy for applications and workloads running on Kubernetes.

服务帐户令牌不应用于尝试与 Kubernetes 集群交互的最终用户，但它们是在 Kubernetes 上运行的应用程序和工作负载的首选身份验证策略。

## Logging setup

## 日志设置

You should collect and centrally store logs from all the workloads running in the cluster and from the cluster components themselves.

您应该从集群中运行的所有工作负载和集群组件本身收集并集中存储日志。

### There's a retention and archival strategy for logs

### 有日志的保留和归档策略

You should retain 30-45 days of historical logs.

您应该保留 30-45 天的历史日志。

### Logs are collected from Nodes, Control Plane, Auditing

### 日志从节点、控制平面、审计收集

What to collect logs from:

- Nodes (kubelet, container runtime)
- Control plane (API server, scheduler, controller manager)
- Kubernetes auditing (all requests to the API server)

从什么收集日志：

- 节点（kubelet、容器运行时）
- 控制平面（API 服务器、调度程序、控制器管理器）
- Kubernetes 审计（对 API 服务器的所有请求）

What you should collect:

- Application name. Retrieved from metadata labels.
- Application instance. Retrieved from metadata labels.
- Application version. Retrieved from metadata labels.
- Cluster ID. Retrieved from Kubernetes cluster.
- Container name. Retrieved from Kubernetes API.
- Cluster node running this container. Retrieved from Kubernetes cluster.
- Pod name running the container. Retrieved from Kubernetes cluster.
- The namespace. Retrieved from Kubernetes cluster.

你应该收集什么：

- 应用名称。从元数据标签中检索。
- 应用实例。从元数据标签中检索。
- 应用程序版本。从元数据标签中检索。
- 集群 ID。从 Kubernetes 集群中检索。
- 容器名称。从 Kubernetes API 检索。
- 运行此容器的集群节点。从 Kubernetes 集群中检索。
- 运行容器的 Pod 名称。从 Kubernetes 集群中检索。
- 命名空间。从 Kubernetes 集群中检索。

### Prefer a daemon on each node to collect the logs instead of sidecars

### 更喜欢每个节点上的守护进程来收集日志而不是边车

Applications should log to stdout rather than to files.

应用程序应该登录到标准输出而不是文件。

[A daemon on each node can collect the logs from the container runtime](https://rclayton.silvrback.com/container-services-logging-with-docker#effective-logging-infrastructure) (if logging to files, a sidecar container for each pod might be necessary).

[每个节点上的守护进程可以从容器运行时收集日志](https://rclayton.silvrback.com/container-services-logging-with-docker#effective-logging-infrastructure)（如果记录到文件，一个sidecar可能需要每个 pod 的容器)。

### Provision a log aggregation tool

### 提供日志聚合工具

Use a log aggregation tool such as EFK stack (Elasticsearch, Fluentd, Kibana), DataDog, Sumo Logic, Sysdig, GCP Stackdriver, Azure Monitor, AWS CloudWatch.

使用日志聚合工具，例如 EFK 堆栈（Elasticsearch、Fluentd、Kibana）、DataDog、Sumo Logic、Sysdig、GCP Stackdriver、Azure Monitor、AWS CloudWatch。

# 3. Governance

# 3. 治理

Best practices for creating, managing and administering namespaces.

创建、管理和管理命名空间的最佳实践。

## Namespace limits

## 命名空间限制

When you decide to segregate your cluster in namespaces, you should protect against misuses in resources.

当您决定在命名空间中隔离集群时，您应该防止资源滥用。

You shouldn't allow your user to use more resources than what you agreed in advance.

您不应该允许您的用户使用比您事先同意的更多的资源。

Cluster administrators can set constraints to limit the number of objects or amount of computing resources that are used in your project with quotas and limit ranges.

集群管理员可以通过配额和限制范围设置约束来限制项目中使用的对象数量或计算资源数量。

You should check out the official documentation if you need a refresher on [limit ranges](https://kubernetes.io/docs/concepts/policy/limit-range/)

如果你需要复习[限制范围](https://kubernetes.io/docs/concepts/policy/limit-range/)，你应该查看官方文档

### Namespaces have LimitRange

### 命名空间有 LimitRange

Containers without limits can lead to resource contention with other containers and unoptimized consumption of computing resources.

没有限制的容器会导致与其他容器的资源争用和计算资源的未优化消耗。

Kubernetes has two features for constraining resource utilisation: ResourceQuota and LimitRange.

Kubernetes 有两个限制资源使用的特性：ResourceQuota 和 LimitRange。

With the LimitRange object, you can define default values for resource requests and limits for individual containers inside namespaces.

使用 LimitRange 对象，您可以为命名空间内的各个容器定义资源请求和限制的默认值。

Any container created inside that namespace, without request and limit values explicitly specified, is assigned the default values. 

在该命名空间内创建的任何容器，没有明确指定请求和限制值，都会被分配默认值。

You should check out the official documentation if you need a refresher on [resource quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/).

如果您需要复习[资源配额](https://kubernetes.io/docs/concepts/policy/resource-quotas/)，请查看官方文档。

### Namespaces have ResourceQuotas

### 命名空间具有 ResourceQuotas

With ResourceQuotas, you can limit the total resource consumption of all containers inside a Namespace.

使用 ResourceQuotas，您可以限制 Namespace 内所有容器的总资源消耗。

Defining a resource quota for a namespace limits the total amount of CPU, memory or storage resources that can be consumed by all containers belonging to that namespace.

为命名空间定义资源配额会限制属于该命名空间的所有容器可以消耗的 CPU、内存或存储资源的总量。

You can also set quotas for other Kubernetes objects such as the number of Pods in the current namespace.

您还可以为其他 Kubernetes 对象设置配额，例如当前命名空间中的 Pod 数量。

If you're thinking that someone could exploit your cluster and create 20000 ConfigMaps, using the LimitRange is how you can prevent that.

如果您认为有人可以利用您的集群并创建 20000 个 ConfigMap，那么使用 LimitRange 可以防止这种情况发生。

## Pod security policies

## Pod 安全策略

When a Pod is deployed into the cluster, you should guard against:

- the container being compromised
- the container using resources on the node that are not allowed such as process, network or file system

当一个 Pod 部署到集群中时，你应该防范：

- 容器受到威胁
- 容器使用节点上不允许的资源，例如进程、网络或文件系统

More in general, you should restrict what the Pod can do to the bare minimum.

更一般地说，您应该将 Pod 可以做的事情限制在最低限度。

### Enable Pod Security Policies

### 启用 Pod 安全策略

For example, you can use Kubernetes Pod security policies for restricting:

- Access the host process or network namespace
- Running privileged containers
- The user that the container is running as
- Access the host filesystem
- Linux capabilities, Seccomp or SELinux profiles

例如，您可以使用 Kubernetes Pod 安全策略来限制：

- 访问主机进程或网络命名空间
- 运行特权容器
- 容器运行的用户
- 访问主机文件系统
- Linux 功能、Seccomp 或 SELinux 配置文件

Choosing the right policy depends on the nature of your cluster.

选择正确的策略取决于集群的性质。

The following article explains some of the [Kubernetes Pod Security Policy best practices](https://resources.whitesourcesoftware.com/blog-whitesource/kubernetes-pod-security-policy)

以下文章解释了一些 [Kubernetes Pod 安全策略最佳实践](https://resources.whitesourcesoftware.com/blog-whitesource/kubernetes-pod-security-policy)

### Disable privileged containers

### 禁用特权容器

In a Pod, containers can run in "privileged" mode and have almost unrestricted access to resources on the host system.

在 Pod 中，容器可以在“特权”模式下运行，并且几乎可以不受限制地访问主机系统上的资源。

While there are specific use cases where this level of access is necessary, in general, it's a security risk to let your containers do this.

虽然在某些特定用例中需要这种级别的访问，但一般而言，让您的容器执行此操作存在安全风险。

Valid uses cases for privileged Pods include using hardware on the node such as GPUs.

特权 Pod 的有效用例包括在节点上使用硬件，例如 GPU。

You can [learn more about security contexts and privileges containers from this article](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/).

您可以[从本文中了解有关安全上下文和权限容器的更多信息](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/)。

### Use a read-only filesystem in containers

### 在容器中使用只读文件系统

Running a read-only file system in your containers forces your containers to be immutable.

在您的容器中运行只读文件系统会强制您的容器不可变。

Not only does this mitigate some old (and risky) practices such as hot patching, but also helps you prevent the risks of malicious processes storing or manipulating data inside a container.

这不仅可以缓解一些旧的（且有风险的）做法，例如热补丁，还可以帮助您防止恶意进程在容器内存储或操纵数据的风险。

Running containers with a read-only file system might sound straightforward, but it might come with some complexity.

使用只读文件系统运行容器听起来很简单，但可能会带来一些复杂性。

_What if you need to write logs or store files in a temporary folder?_

_如果您需要在临时文件夹中写入日志或存储文件怎么办？_

You can learn about the trade-offs in this article on [running containers securely in production](https://medium.com/@axbaretto/running-docker-containers-securely-in-production-98b8104ef68).

您可以在有关 [在生产中安全运行容器](https://medium.com/@axbaretto/running-docker-containers-securely-in-production-98b8104ef68) 的这篇文章中了解权衡。

### Prevent containers from running as root

### 防止容器以 root 身份运行

A process running in a container is no different from any other process on the host, except it has a small piece of metadata that declares that it's in a container.

在容器中运行的进程与主机上的任何其他进程没有区别，只是它有一小段元数据声明它在容器中。

Hence, root in a container is the same root (uid 0) as on the host machine.

因此，容器中的 root 与主机上的 root (uid 0) 相同。

If a user manages to break out of an application running as root in a container, they may be able to gain access to the host with the same root user.

如果用户设法突破在容器中以 root 身份运行的应用程序，他们可能能够以相同的 root 用户访问主机。

Configuring containers to use unprivileged users, is the best way to prevent privilege escalation attacks.

将容器配置为使用非特权用户，是防止提权攻击的最佳方法。

If you wish to learn more, the follow [article offers some detailed explanation examples of what happens when you run your containers as root](https://medium.com/@mccode/processes-in-containers-should-not-run-as-root-2feae3f0df3b).

如果您想了解更多信息，请参阅[文章提供了一些详细的解释示例，说明当您以 root 身份运行容器时会发生什么](https://medium.com/@mccode/processes-in-containers-should-not-run-as-root-2feae3f0df3b)。

### Limit capabilities

### 限制能力

Linux capabilities give processes the ability to do some of the many privileged operations only the root user can do by default.

Linux 功能使进程能够执行许多特权操作中的一些，默认情况下只有 root 用户才能执行。

For example, `CAP_CHOWN` allows a process to "make arbitrary changes to file UIDs and GIDs".

例如，`CAP_CHOWN` 允许进程“对文件 UID 和 GID 进行任意更改”。

Even if your process doesn't run as `root`, there's a chance that a process could use those root-like features by escalating privileges.

即使您的进程不是以“root”身份运行，进程也有可能通过提升权限来使用这些类似 root 的功能。

In other words, you should enable only the capabilities that you need if you don't want to be compromised.

换句话说，如果您不想受到损害，您应该只启用您需要的功能。

_But what capabilities should be enabled and why?_

_但是应该启用哪些功能以及为什么？_

The following two articles dive into the theory and practical best-practices about capabilities in the Linux Kernel: 

以下两篇文章深入探讨了有关 Linux 内核功能的理论和实践最佳实践：

- [Linux Capabilities: Why They Exist and How They Work](https://blog.container-solutions.com/linux-capabilities-why-they-exist-and-how-they-work)
- [Linux Capabilities In Practice](https://blog.container-solutions.com/linux-capabilities-in-practice)

- [Linux 功能：它们为何存在以及它们如何工作](https://blog.container-solutions.com/linux-capabilities-why-they-exist-and-how-they-work)
- [实践中的 Linux 功能](https://blog.container-solutions.com/linux-capabilities-in-practice)

### Prevent privilege escalation

### 防止权限提升

You should run your container with privilege escalation turned off to prevent escalating privileges using `setuid` or `setgid` binaries.

您应该在关闭权限提升的情况下运行容器，以防止使用 `setuid` 或 `setgid` 二进制文件提升权限。

## Network policies

## 网络政策

A Kubernetes network must adhere to three basic rules:

1. **containers can talk to any other container in the network**, and there's no translation of addresses in the process — i.e. no NAT is involved
1. **nodes in the cluster can talk to any other container in the network and vice-versa**. Even in this case, there's no translation of addresses — i.e. no NAT
1. **a container's IP address is always the same**, independently if seen from another container or itself.

Kubernetes 网络必须遵守三个基本规则：

1. **容器可以与网络中的任何其他容器通信**，并且在此过程中没有地址转换——即不涉及 NAT
1. **集群中的节点可以与网络中的任何其他容器通信，反之亦然**。即使在这种情况下，也没有地址转换——即没有 NAT
1. **一个容器的IP地址总是相同的**，无论是从另一个容器还是它本身来看都是独立的。

The first rule isn't helping if you plan to segregate your cluster in smaller chunks and have isolation between namespaces.

如果您计划将集群隔离为较小的块并在命名空间之间进行隔离，则第一条规则无济于事。

_Imagine if a user in your cluster were able to use any other service in the cluster._

_想象一下您集群中的用户是否能够使用集群中的任何其他服务。_

Now, _imagine if a malicious user in the cluster were to obtain access to the cluster_ — they could make requests to the whole cluster.

现在，_想象一下，如果集群中的恶意用户要获得对集群的访问权限_——他们可以向整个集群发出请求。

To fix that, you can define how Pods should be allowed to communicate in the current namespace and cross-namespace using Network Policies.

为了解决这个问题，您可以使用网络策略定义如何允许 Pod 在当前命名空间和跨命名空间中进行通信。

### Enable network policies

### 启用网络策略

Kubernetes network policies specify the access permissions for groups of pods, much like security groups in the cloud are used to control access to VM instances.

Kubernetes 网络策略指定了 Pod 组的访问权限，就像云中的安全组用于控制对 VM 实例的访问一样。

In other words, it creates firewalls between pods running on a Kubernetes cluster.

换句话说，它在 Kubernetes 集群上运行的 Pod 之间创建防火墙。

If you are not familiar with Network Policies, you can read [Securing Kubernetes Cluster Networking](https://ahmet.im/blog/kubernetes-network-policy/).

如果您不熟悉网络策略，可以阅读 [Securing Kubernetes Cluster Networking](https://ahmet.im/blog/kubernetes-network-policy/)。

### There's a conservative NetworkPolicy in every namespace

### 每个命名空间中都有一个保守的 NetworkPolicy

This repository contains various use cases of Kubernetes Network Policies and samples YAML files to leverage in your setup. If you ever wondered [how to drop/restrict traffic to applications running on Kubernetes](https://github.com/ahmetb/kubernetes-network-policy-recipes), read on.

此存储库包含 Kubernetes 网络策略的各种用例和示例 YAML 文件以在您的设置中利用。如果您想知道 [如何删除/限制在 Kubernetes 上运行的应用程序的流量](https://github.com/ahmetb/kubernetes-network-policy-recipes)，请继续阅读。

## Role-Based Access Control (RBAC) policies

## 基于角色的访问控制 (RBAC) 策略

Role-Based Access Control (RBAC) allows you to define policies on how to access resources in your cluster.

基于角色的访问控制 (RBAC) 允许您定义有关如何访问集群中资源的策略。

It's common practice to give away the least permission needed, _but what is practical and how do you quantify the least privilege?_

通常的做法是放弃所需的最少权限，_但是什么是实用的，您如何量化最少权限？_

Fine-grained policies provide greater security but require more effort to administrate.

细粒度策略提供更高的安全性，但需要更多的管理工作。

Broader grants can give unnecessary API access to service accounts but are easier to controls.

更广泛的授权可以为服务帐户提供不必要的 API 访问权限，但更易于控制。

_Should you create a single policy per namespace and share it?_

_是否应该为每个命名空间创建一个策略并共享它？_

_Or perhaps it's better to have them on a more granular basis?_

_或者最好将它们放在更细粒度的基础上？_

There's no one-size-fits-all approach, and you should judge your requirements case by case.

没有一刀切的方法，您应该逐案判断您的需求。

_But where do you start?_

_但你从哪儿开始呢？_

If you start with a Role with empty rules, you can add all the resources that you need one by one and still be sure that you're not giving away too much.

如果您从一个规则为空的角色开始，您可以一一添加您需要的所有资源，并且仍然确保您不会放弃太多。

### Disable auto-mounting of the default ServiceAccount

### 禁用默认 ServiceAccount 的自动挂载

Please note that [the default ServiceAccount is automatically mounted into the file system of all Pods](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server).

请注意 [默认ServiceAccount会自动挂载到所有Pod的文件系统](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server)。

You might want to disable that and provide more granular policies.

您可能希望禁用它并提供更精细的策略。

### RBAC policies are set to the least amount of privileges necessary

### RBAC 策略设置为所需的最少权限

It's challenging to find good advice on how to set up your RBAC rules. In [3 realistic approaches to Kubernetes RBAC](https://thenewstack.io/three-realistic-approaches-to-kubernetes-rbac/), you can find three practical scenarios and practical advice on how to get started.

很难找到关于如何设置 RBAC 规则的好建议。在 [Kubernetes RBAC 的 3 种现实方法](https://thenewstack.io/three-realistic-approaches-to-kubernetes-rbac/) 中，您可以找到三个实用场景和有关如何入门的实用建议。

### RBAC policies are granular and not shared

### RBAC 策略是细粒度的，不共享

Zalando has a concise policy to define roles and ServiceAccounts.

Zalando 有一个简洁的策略来定义角色和 ServiceAccounts。

First, they describe their requirements:

- Users should be able to deploy, but they shouldn't be allowed to read Secrets for example
- Admins should get full access to all resources
- Applications should not gain write access to the Kubernetes API by default
- It should be possible to write to the Kubernetes API for some uses.

首先，他们描述了他们的要求：

- 用户应该能够部署，但他们不应该被允许读取例如 Secrets
- 管理员应该获得对所有资源的完全访问权限
- 默认情况下，应用程序不应获得对 Kubernetes API 的写访问权限
- 应该可以为某些用途写入 Kubernetes API。

The four requirements translate into five separate Roles:

- ReadOnly
- PowerUser
- Operator
- Controller
- Admin 

这四个要求转化为五个独立的角色：

- 只读
- 超级用户
- 操作员
- 控制器
- 行政

You can read about [their decision in this link](https://kubernetes-on-aws.readthedocs.io/en/latest/dev-guide/arch/access-control/adr-004-roles-and-service-accounts.html).

您可以阅读 [他们在此链接中的决定](https://kubernetes-on-aws.readthedocs.io/en/latest/dev-guide/arch/access-control/adr-004-roles-and-service-帐户.html)。

## Custom policies

## 自定义策略

Even if you're able to assign policies in your cluster to resources such as Secrets and Pods, there are some cases where Pod Security Policies (PSPs), Role-Based Access Control (RBAC), and Network Policies fall short.

即使您能够将集群中的策略分配给 Secrets 和 Pod 等资源，在某些情况下，Pod 安全策略 (PSP)、基于角色的访问控制 (RBAC) 和网络策略仍会失效。

As an example, you might want to avoid downloading containers from the public internet and prefer to approve those containers first.

例如，您可能希望避免从公共 Internet 下载容器，而是希望先批准这些容器。

Perhaps you have an internal registry, and only the images in this registry can be deployed in your cluster.

也许您有一个内部注册中心，并且只有这个注册中心中的镜像才能部署到您的集群中。

_How do you enforce that only **trusted containers** can be deployed in the cluster?_

_你如何强制在集群中只能部署**受信任的容器**？_

There's no RBAC policy for that.

没有 RBAC 政策。

Network policies won't work.

网络策略不起作用。

_What should you do?_

_你该怎么办？_

You could use the [Admission controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to vet resources that are submitted to the cluster.

您可以使用 [Admission controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) 来审查提交给集群的资源。

### Allow deploying containers only from known registries

### 仅允许从已知注册中心部署容器

One of the most common custom policies that you might want to consider is to restrict the images that can be deployed in your cluster.

您可能要考虑的最常见的自定义策略之一是限制可以部署在集群中的映像。

[The following tutorial explains how you can use the Open Policy Agent to restrict not approved images](https://blog.openpolicyagent.org/securing-the-kubernetes-api-with-open-policy-agent-ce93af0552c3#3c6e) .

[以下教程解释了如何使用 Open Policy Agent 来限制未批准的图像](https://blog.openpolicyagent.org/securing-the-kubernetes-api-with-open-policy-agent-ce93af0552c3#3c6e) .

### Enforce uniqueness in Ingress hostnames

### 强制 Ingress 主机名的唯一性

When a user creates an Ingress manifest, they can use any hostname in it.

当用户创建 Ingress 清单时，他们可以使用其中的任何主机名。

```yaml|highlight=7|title=ingress.yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: example-ingress
spec:
  rules:
    - host: first.example.com
      http:
        paths:
          - backend:
              serviceName: service
              servicePort: 80
```


However, you might want to prevent users using **the same hostname multiple times** and overriding each other.

但是，您可能希望防止用户**多次使用相同的主机名** 并相互覆盖。

The official documentation for the Open Policy Agent has [a tutorial on how to check Ingress resources as part of the validation webhook](https://www.openpolicyagent.org/docs/latest/kubernetes-tutorial/#4-define-a-policy-and-load-it-into-opa-via-kubernetes).

Open Policy Agent 的官方文档有[关于如何检查 Ingress 资源作为验证 webhook 的一部分的教程](https://www.openpolicyagent.org/docs/latest/kubernetes-tutorial/#4-define-a-policy-and-load-it-into-opa-via-kubernetes)。

### Only use approved domain names in the Ingress hostnames

### 仅在 Ingress 主机名中使用批准的域名

When a user creates an Ingress manifest, they can use any hostname in it.
a

当用户创建 Ingress 清单时，他们可以使用其中的任何主机名。
一种

```yaml|highlight=7|title=ingress.yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: example-ingress
spec:
  rules:
    - host: first.example.com
      http:
        paths:
          - backend:
              serviceName: service
              servicePort: 80
```


However, you might want to prevent users using **invalid hostnames**.

但是，您可能希望阻止用户使用**无效主机名**。

The official documentation for the Open Policy Agent has [a tutorial on how to check Ingress resources as part of the validation webhook](https://www.openpolicyagent.org/docs/latest/kubernetes-tutorial/#4-define-a-policy-and-load-it-into-opa-via-kubernetes). 

Open Policy Agent 的官方文档有[关于如何检查 Ingress 资源作为验证 webhook 的一部分的教程](https://www.openpolicyagent.org/docs/latest/kubernetes-tutorial/#4-define-a-policy-and-load-it-into-opa-via-kubernetes)。

