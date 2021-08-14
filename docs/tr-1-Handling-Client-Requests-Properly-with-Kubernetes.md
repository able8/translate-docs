# [Handling Client Requests Properly with Kubernetes](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/)

# [使用 Kubernetes 正确处理客户端请求](https://freecontent.manning.com/handling-client-requests-properly-with-kubernetes/)

From [Kubernetes in Action](https://www.manning.com/books/kubernetes-in-action)

摘自 [Kubernetes 实战](https://www.manning.com/books/kubernetes-in-action)

It goes without saying that we want client requests to be handled properly. We obviously don’t want to see broken connections when pods are starting up or shutting down. By itself, Kubernetes doesn’t ensure this never happens. Your app needs to follow a few rules to prevent broken connections. This article discusses those rules.

不用说，我们希望正确处理客户端请求。我们显然不希望在 pod 启动或关闭时看到断开的连接。 Kubernetes 本身并不能确保这种情况永远不会发生。您的应用需要遵循一些规则来防止连接断开。本文讨论了这些规则。

**Making sure all client requests are handled properly**

**确保所有客户端请求都得到正确处理**

Let’s look at the pod’s lifecycle from the perspective of the pod’s clients (clients consuming the service the pod is providing). We want to make sure that our clients’ requests are properly handled, because if connections start breaking when pods start up or shut down, we’re in trouble. Kubernetes by itself doesn’t guarantee that this won’t happen, so let’s see what we need to do to prevent it from happening.

让我们从 pod 的客户端（使用 pod 提供的服务的客户端）的角度来看一下 pod 的生命周期。我们希望确保我们客户的请求得到正确处理，因为如果在 pod 启动或关闭时连接开始中断，我们就有麻烦了。 Kubernetes 本身并不能保证这不会发生，所以让我们看看我们需要做些什么来防止它发生。

**Preventing broken client connections when a pod is starting up**

**在 Pod 启动时防止客户端连接断开**

Ensuring each connection is handled properly at pod startup is fairly simple, if you understand how services and service endpoints work. When a pod is started, it’s added as an endpoint to all the services, whose label selector matches the pod’s labels. The pod also needs to signal to Kubernetes that it’s ready. Until it is, it won’t become a service endpoint and therefore won’t receive any requests from clients.

如果您了解服务和服务端点的工作原理，确保在 pod 启动时正确处理每个连接是相当简单的。当 Pod 启动时，它会作为端点添加到所有标签选择器与 Pod 的标签匹配的服务。Pod 还需要向 Kubernetes 发出信号，表明它已准备就绪。在此之前，它不会成为服务端点，因此不会收到来自客户端的任何请求。

If you don’t specify a readiness probe in your pod spec, the pod is considered ready all the time. This means it’ll start receiving requests almost immediately – as soon as the first Kube-Proxy updates the iptables rules on its node and the first client pod tries to connect to the service. If your app isn’t ready to accept connections by then, clients will see “connection refused” types of errors.

如果您没有在 pod 规范中指定就绪探针，则 pod 一直被认为是就绪的。这意味着它几乎会立即开始接收请求——只要第一个 Kube-Proxy 更新其节点上的 iptables 规则并且第一个客户端 pod 尝试连接到服务。如果您的应用到那时还没有准备好接受连接，客户端将看到“连接被拒绝”类型的错误。

All you need to do is make sure that your readiness probe returns success only when your app is ready to properly handle incoming requests. A good first step is to add an HTTP GET readiness probe and point it to the base URL of your app. In a lot of cases, that gets you far enough and saves you from having to implement a special readiness endpoint in your app.

您需要做的就是确保只有当您的应用程序准备好正确处理传入请求时，您的就绪探针才会返回成功。一个好的第一步是添加一个 HTTP GET 就绪探针并将其指向您的应用程序的基本 URL。在很多情况下，这足以让您走得更远，并使您不必在您的应用程序中实现特殊的就绪端点。

**Preventing broken connections during pod shutdown**

**防止在 pod 关闭期间断开连接**

Now let’s see what happens at the other end of a pod’s life – when the pod is deleted and its containers are terminated. The pod’s containers should start shutting down cleanly as soon as they receive a SIGTERM signal (or even before that – when its pre-stop hook is executed), but does this ensure all client requests will be handled properly?

现在让我们看看在 pod 生命周期的另一端会发生什么——当 pod 被删除并且它的容器被终止时。 pod 的容器应该在收到 SIGTERM 信号后立即开始关闭（或者甚至在此之前——当它的 pre-stop 钩子被执行时），但这是否确保所有客户端请求都能得到正确处理？

How should the app behave when it receives a termination signal? Should it continue accepting requests? What about requests that have already been received, but haven’t completed yet? What about persistent HTTP connections, which may be in between requests, but are open (there’s no active request on the connection)? Before we can answer those questions, we need to take a detailed look at the chain of events that unfolds across the cluster when a pod is deleted.

应用程序在收到终止信号时应该如何表现？它应该继续接受请求吗？已经收到但尚未完成的请求呢？持久性 HTTP 连接怎么样，它可能在请求之间，但是是打开的（连接上没有活动请求）？在回答这些问题之前，我们需要详细查看删除 pod 时在集群中呈现的事件链。

**Understanding the sequence of events occurring at pod deletion**

**了解 pod 删除时发生的事件顺序**

You need to always keep in mind that the components in a Kubernetes cluster run as separate processes on multiple machines. They aren’t part of a single big monolithic process. It takes time for all the components to be on the same page regarding the state of the cluster. Let’s explore this fact by looking at what happens across the cluster when a pod is deleted.

您需要始终牢记 Kubernetes 集群中的组件在多台机器上作为单独的进程运行。它们不是一个大的单体程序的一部分。所有组件都需要时间在集群状态上达成共识。让我们通过查看删除 Pod 时集群中发生的情况来探索这一事实。

When a request for a pod deletion is received by the API server, it first modifies the state in *etcd* and then notifies its watchers of the deletion. Among those watchers are the Kubelet and the Endpoints controller. The two sequences of events, which happen in parallel (marked with either A or B), are shown in figure 1.

当 API 服务器收到 pod 删除请求时，它首先修改 *etcd* 中的状态，然后通知其删除事件的观察者。这些观察者包括 Kubelet 和 Endpoints 控制器。并行发生的两个事件序列（用 A 或 B 标记）如图 1 所示。

------

![img](https://freecontent.manning.com/wp-content/uploads/Luksa_HCRPwK_01.png)

Figure 1 Sequence of events that occurs when a pod is deleted

图 1 删除 Pod 时发生的事件序列

------

In the A sequence of events, you'll see that as soon as the Kubelet receives the notification that the pod should be terminated, it initiates the shutdown sequence (run the pre-stop hook, send SIGTERM, wait for some time and then forcibly kill the container if it hasn't yet terminated on its own). If the app responds to the SIGTERM by immediately ceasing to receive client requests, any client trying to connect to it receives a Connection refused error. The time it takes for this to happen from the time the pod is deleted is relatively short, because of the direct path from the API server to the Kubelet.

在A序列事件中，你会看到Kubelet一收到Pod应该被终止的通知，它就会启动关闭序列（运行pre-stop钩子，发送SIGTERM，等待一段时间然后强行如果容器尚未自行终止，则终止容器）。如果应用程序通过立即停止接收客户端请求来响应 SIGTERM，则任何尝试连接到它的客户端都会收到连接拒绝错误。由于从 API 服务器到 Kubelet 的直接路径，因此从删除 Pod 到发生这种情况所需的时间相对较短。

Now, let’s look at what happens in the other sequence of events – the one leading up to the pod being removed from the iptables rules (sequence B in the figure). When the Endpoints controller (which runs in the Controller Manager in the Kubernetes Control Plane) receives the notification of the pod being deleted, it removes the pod as an endpoint in all services that the pod is a part of. It does this by modifying the Endpoints API object by sending a REST request to the API server. The API server then notifies everyone watching the Endpoints object. Among those watchers are the Kube-Proxies running on the worker nodes. Each of these proxies updates the *iptables* rules on its node, which is what prevents new connections from being forwarded to the terminating pod. An important detail here is that removing the iptables rules has no effect on existing connections – clients who’re already connected to the pod can still send additional requests to the pod through those existing connections.

现在，让我们看看在其他事件序列中发生了什么——导致 pod 从 iptables 规则中删除的事件（图中的序列 B）。当 Endpoints 控制器（在 Kubernetes 控制平面的控制器管理器中运行）收到 pod 被删除的通知时，它会删除 pod 作为 pod 所属的所有服务中的端点。它通过向 API 服务器发送 REST 请求来修改 Endpoints API 对象来实现此目的。 API 服务器然后通知每个观察 Endpoints 的对象。这些观察者包括在工作节点上运行的 Kube-Proxies。这些代理中的每个代理都会更新其节点上的 *iptables* 规则，这是防止新连接被转发到终止的 pod。这里的一个重要细节是删除 iptables 规则对现有连接没有影响——已经连接到 pod 的客户端仍然可以通过这些现有连接向 pod 发送额外的请求。

Both of these sequences of events happen in parallel. Most likely, the time it takes to shut down the app process in the pod is slightly shorter than the time required for the *iptables* rules to be updated. This is because the chain of events that leads to iptables rules being updated is considerably longer (see figure 2), because the event must first reach the Endpoints controller, which sends a new request to the API server, and then the API server must notify the Kube-Proxy, before the proxy finally modifies the *iptables* rules. This means there’s a high probability of the SIGTERM signal being sent well before the iptables rules are updated on all nodes.

这两个事件序列并行发生。最有可能的是，在 pod 中关闭应用进程所需的时间比更新 *iptables* 规则所需的时间略短。这是因为导致 iptables 规则更新的事件链相当长（见图 2），因为事件必须首先到达 Endpoints 控制器，后者向 API 服务器发送新请求，然后 API 服务器必须通知Kube-Proxy，在代理最终修改 *iptables* 规则之前。这意味着在所有节点上更新 iptables 规则之前，SIGTERM 信号很有可能被发送。

------

![img](https://freecontent.manning.com/wp-content/uploads/Luksa_HCRPwK_02.png)

**Figure 2** Timeline of events when pod is deleted

**图 2** 删除 pod 时的事件时间线

------

The result is that the pod may still receive client requests after it has received the termination signal. If the app stops accepting connections immediately, it causes clients to receive “connection refused” types of errors (like what happens at pod startup if your app isn't capable of accepting connections immediately and you don't define a readiness probe for it) .

结果是 pod 在收到终止信号后仍然可能会收到客户端请求。如果应用程序立即停止接受连接，它会导致客户端收到“连接被拒绝”类型的错误（就像在您的应用程序无法立即接受连接并且您没有为它定义就绪探测，pod 启动时会发生的那样）。

**Solving the problem**

**解决问题**

Googling for solutions to this problem makes it seem that adding a readiness probe to your pod solves this problem. Supposedly, all you need to do is make the readiness probe start failing as soon as the pod receives the SIGTERM. This is supposed to cause the pod to be removed as the endpoint of the service. The removal would happen only after the readiness probe fails for a few consecutive times (this is configurable in the readiness probe spec). And, obviously, the removal then still needs to reach the Kube-Proxy before the pod is removed from *iptables* rules.

在谷歌上搜索这个问题的解决方案，似乎在你的 pod 中添加一个就绪探针可以解决这个问题。据说，您需要做的就是在 pod 收到 SIGTERM 后立即使就绪探针开始失败。这应该会导致 pod 作为服务的端点被删除。只有在就绪探测器连续几次失败后才会删除（这可以在就绪探测器规范中配置）。而且，显然，在从 *iptables* 规则中删除 pod 之前，删除仍然需要到达 Kube-Proxy。

In reality, the readiness probe has absolutely no bearing on the whole process at all. The Endpoints controller removes the pod from the service endpoints as soon as it receives notice of the pod being deleted (when the *deletionTimestamp* field in the pod’s spec is no longer null). From that point on, the result of the readiness probe is irrelevant.

实际上，就绪探针与整个过程完全没有关系。 Endpoints 控制器一旦收到 pod 被删除的通知（当 pod 规范中的 *deletionTimestamp* 字段不再为空时）就会从服务端点中删除 pod。从那时起，就绪探测的结果就无关紧要了。

**What’s the proper solution to the problem? How can we make sure all requests are handled fully? **
**解决问题的正确方法是什么？我们如何确保所有请求都得到充分处理？**

Well, it’s clear that the pod needs to keep accepting connections even after it receives the termination signal, up until all the Kube-proxies have finished updating the iptables rules. Well, it’s not only the Kube-Proxies. There may also be Ingress controllers or load balancers forwarding connections to the pod directly, without going through the service (iptables). This also includes clients using client-side load-balancing. To ensure none of the clients experience broken connections, you’d have to wait until all of them somehow notify you they’ll no longer forward connections to the pod.

嗯，很明显，即使在收到终止信号后，pod 也需要继续接受连接，直到所有 Kube 代理都完成了 iptables 规则的更新。嗯，不仅仅是 Kube-Proxies。也可能有 Ingress 控制器或负载均衡器直接将连接转发到 pod，而无需通过服务（iptables）。这还包括使用客户端负载平衡的客户端。为了确保所有客户端都不会遇到连接断开的情况，您必须等到所有客户端都以某种方式通知您他们不再将连接转发到 pod。

This is impossible, because all those components are distributed across many different computers. Even if you knew the location of each and every one of them and could wait until all of them say it’s ok to shut down the pod, what would you do if one of them doesn’t respond? How long do you wait for the response? Remember, during that time, you’re holding up the shutdown process.

这是不可能的，因为所有这些组件都分布在许多不同的计算机上。即使你知道每个组件的位置并且可以等到他们都说可以关闭 pod，如果其中一个没有响应，你会怎么做？你要等回复多久？请记住，在此期间，您正在阻止关机过程。

The only reasonable thing you can do is wait for a long enough time to ensure all of the proxies have done their job. But how long is long enough? A few seconds should be enough in most situations, but obviously, there’s no guarantee it’ll suffice every time. When the API server or the Endpoints controller are overloaded, it may take longer for the notification to reach the Kube-Proxy. It’s important to understand that you can’t solve the problem perfectly, but even a five or ten second delay should improve the user experience considerably. You can use a longer delay, but don't go overboard, because the delay prevents the container from shutting down promptly and causes the pod to still be shown in lists long after it has been deleted, which is always frustrating to the user deleting the pod.

您可以做的唯一合理的事情是等待足够长的时间以确保所有代理都完成了他们的工作。但多久才够长？在大多数情况下，几秒钟就足够了，但显然，不能保证每次都足够。当 API 服务器或 Endpoints 控制器过载时，通知到达 Kube-Proxy 可能需要更长的时间。重要的是要了解您无法完美解决问题，但即使是五到十秒的延迟也应该可以显着改善用户体验。您可以使用更长的延迟，但不要过火，因为延迟会阻止容器立即关闭，并导致 pod 在删除后很长时间内仍会显示在列表中，这总是让删除该容器的用户感到沮丧荚。

 

Properly shutting down an application includes these steps:

- Wait for a few seconds, then stop accepting new connections,
 - Close all keep-alive connections that aren’t in the middle of a request,
 - Wait for all active requests to finish, and then
 - Shut down completely.

 正确关闭应用程序包括以下步骤：

- 等待几秒钟，然后停止接受新连接，
- 关闭所有不在请求中间的保持活动连接，
- 等待所有活动请求完成，然后
- 完全关闭。

 

To understand what’s happening with the connections and requests during this process, examine figure 3 carefully.

要了解在此过程中连接和请求发生了什么，请仔细检查图 3。

------

![img](https://freecontent.manning.com/wp-content/uploads/Luksa_HCRPwK_03.png)

**Figure 3** Properly handling existing and new connections after receiving a termination signal

**图 3** 收到终止信号后正确处理现有连接和新连接

------

Not as simple as exiting the process immediately upon receiving the termination signal, right? Is it worth going through all this? That’s for you to decide. But the least you can do is add a pre-stop hook that waits a few seconds. Something like this, perhaps:

不像收到终止信号就立即退出进程那么简单，对吧？值得经历这一切吗？那由你来决定。但是您至少可以添加一个等待几秒钟的预停止钩子。像这样的事情，也许：

```
   lifecycle:
    preStop:
     exec:
      command:
      - sh
      - -c
      - "sleep 5"
```


This way, you don’t need to modify the code of your app at all. If your app already ensures all in-flight requests are processed completely, this pre-stop delay may be all you need.

这样，您根本不需要修改应用程序的代码。如果您的应用程序已经确保完全处理所有进行中的请求，则此停止前延迟可能就是您所需要的。

That’s all for this article.

这就是本文的全部内容。

------

For more, check out the entire book on liveBook [here](https://livebook.manning.com/#!/book/kubernetes-in-action/).

有关更多信息，请查看 liveBook [此处](https://livebook.manning.com/#!/book/kubernetes-in-action/) 上的整本书。

------

