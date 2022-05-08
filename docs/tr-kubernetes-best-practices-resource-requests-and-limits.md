# Kubernetes best practices: Resource requests and limits

# Kubernetes 最佳实践：资源请求和限制

May 11, 2018

2018 年 5 月 11 日

When Kubernetes schedules a Pod, it’s important that the containers have enough resources to actually run. If you schedule a large application on a node with limited resources, it is possible for the node to run out of memory or CPU resources and for things to stop working!

当 Kubernetes 调度 Pod 时，重要的是容器有足够的资源来实际运行。如果您在资源有限的节点上调度大型应用程序，则该节点可能会耗尽内存或 CPU 资源并停止工作！

It’s also possible for applications to take up more resources than they should. This could be caused by a team spinning up more replicas than they need to artificially decrease latency (hey, it's easier to spin up more copies than make your code more efficient!), to a bad configuration change that causes a program to go out of control and use 100% of the available CPU. Regardless of whether the issue is caused by a bad developer, bad code, or bad luck, what’s important is that you be in control.

应用程序也可能占用比应有的更多资源。这可能是由于团队创建的副本数量超过了人为减少延迟所需的数量（嘿，创建更多副本比让您的代码更高效更容易！），以及导致程序退出的错误配置更改控制和使用 100% 的可用 CPU。不管问题是由糟糕的开发人员、糟糕的代码还是运气不好引起的，重要的是你能控制住。

In this episode of Kubernetes best practices, let’s take a look at how you can solve these problems using resource requests and limits.

在本期 Kubernetes 最佳实践中，我们来看看如何使用资源请求和限制来解决这些问题。

### Requests and Limits

### 请求和限制

Requests and limits are the mechanisms Kubernetes uses to control resources such as CPU and memory. Requests are what the container is guaranteed to get. If a container requests a resource, Kubernetes will only schedule it on a node that can give it that resource. Limits, on the other hand, make sure a container never goes above a certain value. The container is only allowed to go up to the limit, and then it is restricted.

请求和限制是 Kubernetes 用来控制 CPU 和内存等资源的机制。请求是容器保证得到的。如果容器请求资源，Kubernetes 只会将其调度到可以为其提供该资源的节点上。另一方面，限制确保容器永远不会超过某个值。容器只允许上升到极限，然后被限制。

It is important to remember that the limit can never be lower than the request. If you try this, Kubernetes will throw an error and won’t let you run the container.

重要的是要记住限制永远不能低于请求。如果你尝试这个，Kubernetes 会抛出一个错误并且不会让你运行容器。

Requests and limits are on a per-container basis. While Pods usually contain a single container, it’s common to see Pods with multiple containers as well. Each container in the Pod gets its own individual limit and request, but because Pods are always scheduled as a group, you need to add the limits and requests for each container together to get an aggregate value for the Pod.

请求和限制基于每个容器。虽然 Pod 通常包含一个容器，但通常也会看到具有多个容器的 Pod。 Pod 中的每个容器都有自己的单独限制和请求，但由于 Pod 总是作为一个组进行调度，因此您需要将每个容器的限制和请求加在一起以获得 Pod 的聚合值。

To control what requests and limits a container can have, you can set quotas at the Container level and at the Namespace level. If you want to learn more about Namespaces, see [this previous installment](https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-organizing-with-namespaces) from our blog series !

要控制容器可以拥有的请求和限制，您可以在容器级别和命名空间级别设置配额。如果您想了解有关命名空间的更多信息，请参阅我们的博客系列中的 [上一期](https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-organizing-with-namespaces) ！

Let’s see how these work.

让我们看看这些是如何工作的。



### Container settings

### 容器设置

There are two types of resources: CPU and Memory. The Kubernetes scheduler uses these to figure out where to run your pods.

有两种类型的资源：CPU 和内存。 Kubernetes 调度程序使用这些来确定在哪里运行您的 pod。

[Here are the docs for these resources.](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/)

[这里是这些资源的文档。](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/)

If you are running in [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine/) (GKE), the default Namespace already has some requests and limits set up for you.

如果您在 [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine/)(GKE) 中运行，则默认命名空间已经为您设置了一些请求和限制。

![google-namespace-requests-limits93uk.PNG](https://storage.googleapis.com/gweb-cloudblog-publish/images/google-namespace-requests-limits93uk.max-500x500.PNG)

These default settings are okay for “Hello World”, but it is important to change them to fit your app.

这些默认设置适用于“Hello World”，但更改它们以适合您的应用程序很重要。

A typical Pod spec for resources might look something like this. This pod has two containers:

典型的 Pod 资源规范可能看起来像这样。这个 pod 有两个容器：

![gcp-container-pod-specil8w.PNG](https://storage.googleapis.com/gweb-cloudblog-publish/images/gcp-container-pod-specil8w.max-400x400.PNG)

Each container in the Pod can set its own requests and limits, and these are all additive. So in the above example, the Pod has a total request of 500 mCPU and 128 MiB of memory, and a total limit of 1 CPU and 256MiB of memory.

Pod 中的每个容器都可以设置自己的请求和限制，这些都是附加的。所以在上面的例子中，Pod 的总请求量为 500 mCPU 和 128 MiB 内存，总限制为 1 个 CPU 和 256MiB 内存。

**CPU**

**中央处理器**

CPU resources are defined in millicores. If your container needs two full cores to run, you would put the value “2000m”. If your container only needs ¼ of a core, you would put a value of “250m”. 

CPU 资源以毫秒为单位定义。如果你的容器需要两个完整的核心来运行，你可以输入“2000m”。如果您的容器只需要 ¼ 个核心，您将输入“250m”的值。

One thing to keep in mind about CPU requests is that if you put in a value larger than the core count of your biggest node, your pod will never be scheduled. Let’s say you have a pod that needs four cores, but your Kubernetes cluster is comprised of dual core VMs—your pod will never be scheduled!

关于 CPU 请求要记住的一件事是，如果您输入的值大于最大节点的核心数，您的 pod 将永远不会被调度。假设您有一个需要四个核心的 pod，但您的 Kubernetes 集群由双核 VM 组成——您的 pod 永远不会被调度！

Unless your app is specifically designed to take advantage of multiple cores (scientific computing and some databases come to mind), it is usually a best practice to keep the CPU request at ‘1’ or below, and run more replicas to scale it out. This gives the system more flexibility and reliability.

除非您的应用程序专门设计为利用多核（想到科学计算和一些数据库），否则将 CPU 请求保持在“1”或以下通常是最佳实践，并运行更多副本以扩展它。这使系统具有更大的灵活性和可靠性。

It’s when it comes to CPU limits that things get interesting. CPU is considered a “compressible” resource. If your app starts hitting your CPU limits, Kubernetes starts throttling your container. This means the CPU will be artificially restricted, giving your app potentially worse performance! However, it won’t be terminated or evicted. You can use a [liveness health check](https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-setting-up-health-checks-with-readiness-and-liveness-probes) to make sure performance has not been impacted.

当涉及到 CPU 限制时，事情就变得有趣了。 CPU 被视为“可压缩”资源。如果您的应用开始达到 CPU 限制，Kubernetes 就会开始限制您的容器。这意味着 CPU 将受到人为限制，使您的应用程序性能可能更差！但是，它不会被终止或驱逐。您可以使用 [liveness 健康检查](https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-setting-up-health-checks-with-readiness-and-liveness-探针)以确保性能没有受到影响。

**Memory**

Memory resources are defined in bytes. Normally, you give a [mebibyte](https://en.wikipedia.org/wiki/Byte#Multiple-byte_units) value for memory (this is basically the same thing as a megabyte), but you can give anything from [bytes to petabytes](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#meaning-of-memory).

内存资源以字节为单位。通常，您为内存提供 [mebibyte](https://en.wikipedia.org/wiki/Byte#Multiple-byte_units)值（这与兆字节基本相同)，但您可以提供 [bytes到 PB](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#meaning-of-memory)。

Just like CPU, if you put in a memory request that is larger than the amount of memory on your nodes, the pod will never be scheduled.

就像 CPU 一样，如果您提出的内存请求大于节点上的内存量，则该 pod 将永远不会被调度。

Unlike CPU resources, memory cannot be compressed. Because there is no way to throttle memory usage, if a container goes past its memory limit it will be terminated. If your pod is managed by a [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/), [StatefulSet](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/), [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/), or another type of controller, then the controller spins up a replacement.

与 CPU 资源不同，内存无法压缩。因为没有办法限制内存使用，如果容器超过其内存限制，它将被终止。如果您的 pod 由 [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)、[StatefulSet](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)、[DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) 或其他类型的控制器，然后控制器启动替换。

**Nodes**

**节点**

It is important to remember that you cannot set requests that are larger than resources provided by your nodes. For example, if you have a cluster of dual-core machines, a Pod with a request of 2.5 cores will never be scheduled! You can find the total resources for Kubernetes Engine VMs [here](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture#node_allocatable).

请务必记住，您不能设置大于节点提供的资源的请求。例如，如果您有一个双核机器集群，那么请求 2.5 核的 Pod 将永远不会被调度！您可以在 [此处](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture#node_allocatable) 找到 Kubernetes Engine 虚拟机的总资源。

### Namespace settings

### 命名空间设置

In an ideal world, Kubernetes’ Container settings would be good enough to take care of everything, but the world is a dark and terrible place. People can easily forget to set the resources, or a rogue team can set the requests and limits very high and take up more than their fair share of the cluster.

在一个理想的世界里，Kubernetes 的 Container 设置足以应付一切，但这个世界是一个黑暗而可怕的地方。人们很容易忘记设置资源，或者流氓团队可以将请求和限制设置得非常高，并占用超过他们在集群中的公平份额。

To prevent these scenarios, you can set up ResourceQuotas and LimitRanges at the Namespace level.

为防止出现这些情况，您可以在 Namespace 级别设置 ResourceQuotas 和 LimitRanges。

**ResourceQuotas**

**资源配额**

After creating Namespaces, you can lock them down using ResourceQuotas. ResourceQuotas are very powerful, but let’s just look at how you can use them to restrict CPU and Memory resource usage.

创建命名空间后，您可以使用 ResourceQuotas 锁定它们。 ResourceQuotas 非常强大，但让我们看看如何使用它们来限制 CPU 和 Memory 资源的使用。

A Quota for resources might look something like this:

资源配额可能如下所示：

![gcp-resourcequota3qo9.PNG](https://storage.googleapis.com/gweb-cloudblog-publish/images/gcp-resourcequota3qo9.max-300x300.PNG)

Looking at this example, you can see there are four sections. Configuring each of these sections is optional.

查看此示例，您可以看到有四个部分。配置这些部分中的每一个都是可选的。

**requests.cpu** is the maximum combined CPU requests in millicores for all the containers in the Namespace. In the above example, you can have 50 containers with 10m requests, five containers with 100m requests, or even one container with a 500m request. As long as the total requested CPU in the Namespace is less than 500m! 

**requests.cpu** 是命名空间中所有容器的最大组合 CPU 请求（以毫秒为单位）。在上面的例子中，你可以有 50 个 10m 请求的容器，5 个 100m 请求的容器，甚至一个 500m 请求的容器。只要在 Namespace 中请求的 CPU 总量小于 500m！

**requests.memory** is the maximum combined Memory requests for all the containers in the Namespace. In the above example, you can have 50 containers with 2MiB requests, five containers with 20MiB CPU requests, or even a single container with a 100MiB request. As long as the total requested Memory in the Namespace is less than 100MiB!

**requests.memory** 是命名空间中所有容器的最大组合内存请求。在上面的示例中，您可以拥有 50 个具有 2MiB 请求的容器，5 个具有 20MiB CPU 请求的容器，甚至是一个具有 100MiB 请求的容器。只要命名空间中请求的总内存小于 100MiB！

**limits.cpu** is the maximum combined CPU limits for all the containers in the Namespace. It’s just like requests.cpu but for the limit.

**limits.cpu** 是命名空间中所有容器的最大组合 CPU 限制。它就像 requests.cpu 一样，但有限制。

**limits.memory** is the maximum combined Memory limits for all containers in the Namespace. It’s just like requests.memory but for the limit.

**limits.memory** 是命名空间中所有容器的最大组合内存限制。它就像 requests.memory 但有限制。

If you are using a production and development Namespace (in contrast to a Namespace per team or service), a common pattern is to put no quota on the production Namespace and strict quotas on the development Namespace. This allows production to take all the resources it needs in case of a spike in traffic.

如果您使用的是生产和开发命名空间（与每个团队或服务的命名空间相反），一种常见的模式是对生产命名空间不设置配额，而对开发命名空间设置严格的配额。这允许生产在流量激增的情况下获取所需的所有资源。

**LimitRange**

**限制范围**

You can also create a LimitRange in your Namespace. Unlike a Quota, which looks at the Namespace as a whole, a LimitRange applies to an individual container. This can help prevent people from creating super tiny or super large containers inside the Namespace.

您还可以在命名空间中创建一个 LimitRange。与将命名空间视为一个整体的配额不同，LimitRange 适用于单个容器。这有助于防止人们在命名空间内创建超小型或超大型容器。

A LimitRange might look something like this:

LimitRange 可能看起来像这样：

![gcp-limit-range228w.PNG](https://storage.googleapis.com/gweb-cloudblog-publish/images/gcp-limit-range228w.max-400x400.PNG)

Looking at this example, you can see there are four sections. Again, setting each of these sections is optional.

查看此示例，您可以看到有四个部分。同样，设置这些部分中的每一个都是可选的。

The **default section** sets up the default **limits** for a container in a pod. If you set these values in the limitRange, any containers that don’t explicitly set these themselves will get assigned the default values.

**default 部分** 为 pod 中的容器设置默认的 **limits**。如果您在 limitRange 中设置这些值，则任何未明确设置这些值的容器都将被分配默认值。

The **defaultRequest section** sets up the default **requests** for a container in a pod. If you set these values in the limitRange, any containers that don’t explicitly set these themselves will get assigned the default values.

**defaultRequest 部分** 为 pod 中的容器设置默认的 **requests**。如果您在 limitRange 中设置这些值，则任何未明确设置这些值的容器都将被分配默认值。

The **max section** will set up the **maximum limits** that a container in a Pod can set. The default section cannot be higher than this value. Likewise, limits set on a container cannot be higher than this value. It is important to note that if this value is set and the default section is not, any containers that don’t explicitly set these values themselves will get assigned the max values as the limit.

**max 部分** 将设置 Pod 中的容器可以设置的**最大限制**。默认部分不能高于此值。同样，对容器设置的限制不能高于此值。需要注意的是，如果设置了这个值并且没有设置默认部分，那么任何没有明确设置这些值的容器都会被分配最大值作为限制。

The **min section** sets up the **minimum Requests** that a container in a Pod can set. The defaultRequest section cannot be lower than this value. Likewise, requests set on a container cannot be lower than this value either. It is important to note that if this value is set and the defaultRequest section is not, the min value becomes the defaultRequest value too.

**min 部分** 设置 Pod 中的容器可以设置的 **minimum Requests**。 defaultRequest 部分不能低于此值。同样，在容器上设置的请求也不能低于此值。需要注意的是，如果设置了这个值并且没有设置 defaultRequest 部分，那么最小值也会变成 defaultRequest 值。

### The lifecycle of a Kubernetes Pod

### Kubernetes Pod 的生命周期

At the end of the day, these resources requests are used by the Kubernetes scheduler to run your workloads. It is important to understand how this works so you can tune your containers correctly.

最终，Kubernetes 调度程序使用这些资源请求来运行您的工作负载。了解其工作原理非常重要，这样您才能正确调整容器。

Let’s say you want to run a Pod on your Cluster. Assuming the Pod specifications are valid, the Kubernetes scheduler will use round-robin load balancing to pick a Node to run your workload.

假设您想在集群上运行一个 Pod。假设 Pod 规范有效，Kubernetes 调度程序将使用循环负载平衡来选择一个节点来运行您的工作负载。

_Note: The exception to this is if you use a nodeSelector or similar mechanism to force Kubernetes to schedule your Pod in a specific place. The resource checks still occur when you use a nodeSelector, but Kubernetes will only check nodes that have the required label._

_注意：如果您使用 nodeSelector 或类似机制来强制 Kubernetes 将您的 Pod 调度到特定位置，则例外情况。使用 nodeSelector 时仍会进行资源检查，但 Kubernetes 只会检查具有所需标签的节点。_

Kubernetes then checks to see if the Node has enough resources to fulfill the resources requests on the Pod’s containers. If it doesn’t, it moves on to the next node. 

Kubernetes 然后检查节点是否有足够的资源来满足 Pod 容器上的资源请求。如果没有，它会移动到下一个节点。

If none of the Nodes in the system have resources left to fill the requests, then Pods go into a “pending” state. By using [GKE](https://cloud.google.com/kubernetes-engine/) features such as the [Node Autoscaler](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-autoscaler), Kubernetes Engine can automatically detect this state and create more Nodes automatically. If there is excess capacity, the autoscaler can also scale down and remove Nodes to save you money!

如果系统中的所有节点都没有剩余资源来填充请求，则 Pod 进入“待处理”状态。通过使用 [GKE](https://cloud.google.com/kubernetes-engine/) 功能，例如 [Node Autoscaler](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-autoscaler)，Kubernetes Engine 可以自动检测这种状态并自动创建更多节点。如果容量过剩，自动扩缩器还可以缩减并移除节点以节省您的资金！

But what about limits? As you know, limits can be higher than the requests. What if you have a Node where the sum of all the container Limits is actually higher than the resources available on the machine?

但是限制呢？如您所知，限制可能高于请求。如果您有一个节点，其中所有容器限制的总和实际上高于机器上可用的资源怎么办？

At this point, Kubernetes goes into something called an “[overcommitted state](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/resource-qos.md#qos-classes) .” Here is where things get interesting. Because CPU can be compressed, Kubernetes will make sure your containers get the CPU they requested and will throttle the rest. Memory cannot be compressed, so Kubernetes needs to start making decisions on what containers to terminate if the Node runs out of memory.

此时，Kubernetes 进入了一种称为“[过度使用状态](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/resource-qos.md#qos-classes) 。”这就是事情变得有趣的地方。因为 CPU 可以被压缩，Kubernetes 将确保您的容器获得它们请求的 CPU，并限制其余的。内存无法压缩，因此 Kubernetes 需要开始决定在节点内存不足时终止哪些容器。

Let’s imagine a scenario where we have a machine that is running out of memory. What will Kubernetes do?

让我们想象一个场景，我们有一台内存不足的机器。 Kubernetes 会做什么？

_Note: The following is true for Kubernetes 1.9 and above. In previous versions, it uses a slightly different process. See this doc for an in-depth explanation._

_注意：以下内容适用于 Kubernetes 1.9 及更高版本。在以前的版本中，它使用稍微不同的过程。请参阅此文档以获得深入的解释。_

Kubernetes looks for Pods that are using more resources than they requested. If your Pod’s containers have no requests, then by default they are using more than they requested, so these are prime candidates for termination. Other prime candidates are containers that have gone over their request but are still under their limit.

Kubernetes 会查找使用的资源超过其请求的 Pod。如果您的 Pod 的容器没有请求，那么默认情况下它们使用的数量超过了请求的数量，因此这些是终止的主要候选者。其他主要候选者是已经超过他们的请求但仍低于其限制的容器。

If Kubernetes finds multiple pods that have gone over their requests, it will then rank these by the Pod's [priority](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/scheduling/pod-priority-api.md), and terminate the lowest priority pods first. If all the Pods have the same priority, Kubernetes terminates the Pod that’s the most over its request.

如果 Kubernetes 发现多个 Pod 已经处理了他们的请求，它将按照 Pod 的 [priority](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/scheduling/pod-priority-api.md)，并首先终止最低优先级的 pod。如果所有 Pod 具有相同的优先级，Kubernetes 将终止其请求最多的 Pod。

In very rare scenarios, Kubernetes might be forced to terminate Pods that are still within their requests. This can happen when critical system components, like the kubelet or docker, start taking more resources than were reserved for them.

在极少数情况下，Kubernetes 可能会被迫终止仍在其请求范围内的 Pod。当关键系统组件（如 kubelet 或 docker）开始占用比为它们保留的资源更多的资源时，就会发生这种情况。

### Conclusion

###  结论

While your Kubernetes cluster might work fine without setting resource requests and limits, you will start running into stability issues as your teams and projects grow. Adding requests and limits to your Pods and Namespaces only takes a little extra effort, and can save you from running into many headaches down the line!

虽然您的 Kubernetes 集群可能在不设置资源请求和限制的情况下正常工作，但随着团队和项目的增长，您将开始遇到稳定性问题。向您的 Pod 和命名空间添加请求和限制只需要一点额外的努力，并且可以让您免于遇到许多令人头疼的问题！

https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-resource-requests-and-limits 

https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-resource-requests-and-limits

