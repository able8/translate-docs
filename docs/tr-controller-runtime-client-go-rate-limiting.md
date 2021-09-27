# Rate Limiting in controller-runtime and client-go

# 控制器运行时和客户端运行中的速率限制

## February 14, 2021

## 2021 年 2 月 14 日

If you have written a [Kubernetes controller](https://kubernetes.io/docs/concepts/architecture/controller/), you are likely [familiar with `controller-runtime`](https://github.com/kubernetes-sigs/controller-runtime), or [at least `client-go`](https://github.com/kubernetes/client-go). `controller-runtime` is a framework for building controllers that allows consumers to setup multiple controllers that are all handled under a controller manager. Behind the the scenes, `controller-runtime` is using `client-go` to communicate with the Kubernetes [API server](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/) to watch for changes on resources and pass them along to the relevant controllers. It handles a number of aspects of plumbing up these controllers, including caching, queuing, and more. One of these components is *rate limiting*.

如果你写过一个 [Kubernetes 控制器](https://kubernetes.io/docs/concepts/architecture/controller/)，你很可能 [熟悉`controller-runtime`](https://github.com/kubernetes-sigs/controller-runtime) 或 [至少`client-go`](https://github.com/kubernetes/client-go)。 `controller-runtime` 是一个构建控制器的框架，它允许消费者设置多个控制器，这些控制器都在控制器管理器下处理。在幕后，`controller-runtime` 使用 `client-go` 与 Kubernetes [API 服务器](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/) 观察资源的变化并将它们传递给相关的控制器。它处理连接这些控制器的许多方面，包括缓存、排队等。这些组件之一是*速率限制*。

## What is rate limiting?

## 什么是速率限制？

> This section is a very basic overview of rate limiting. You may skip it if you are already well versed in the concepts, but it may be useful for framing later sections.

> 本节是对速率限制的一个非常基本的概述。如果您已经精通概念，则可以跳过它，但它可能对构建后面的部分很有用。

Rate limiting has existed in software since computer networking was invented, and before that in many other human processes. In fact, as we discuss rate limiting, you’ll likely see many parallels to daily tasks you perform and organizational patterns within companies and communities.

自计算机网络发明以来，软件中就存在速率限制，而在此之前，许多其他人类过程中就存在速率限制。事实上，在我们讨论速率限制时，您可能会看到与您执行的日常任务以及公司和社区内的组织模式有许多相似之处。

Rate limiting is necessary for effective communication to occur between any two parties. Software communicates by passing messages between different processes of execution, whether it be via the operating system, dedicated hardware facilities, across a network, or some combination of the three. In a client-server model, the client is typically requesting that the server perform some work on its behalf. It takes time for the server to perform that work, meaning that if many clients ask it to do work at the same time and the server does not have the capacity to handle them all, it has a choice to make. It can:

速率限制对于任何两方之间的有效通信都是必要的。软件通过在不同执行进程之间传递消息来进行通信，无论是通过操作系统、专用硬件设施、网络还是三者的某种组合。在客户端-服务器模型中，客户端通常请求服务器代表它执行某些工作。服务器执行该工作需要时间，这意味着如果许多客户端要求它同时完成工作而服务器没有能力处理所有这些工作，则它可以做出选择。它可以：

1. Drop requests with no response.
2. Wait to respond to requests until the work can be performed to completion.
3. Respond to requests indicating that the work cannot be performed at this time, but that the client should ask again at a future time.
4. Add the work to a queue and respond to the request to with a message saying that it will let the client know when it completes the work.

1. 删除请求而没有响应。
2. 等待响应请求，直到工作可以执行完成。
3. 回应表示此时无法执行工作，但客户应在未来再次询问的请求。
4. 将工作添加到队列中，并用一条消息响应请求，说当它完成工作时会通知客户端。

If the client and server know each other well (i.e. their communication patterns are well-known to each other), any of these methods could be valid communication models. Think about your relationships with other people in your life. You may know people who communicate in a wide variety of manners, but you likely can work effectively with all of them if the communication style is well-known.

如果客户端和服务器彼此非常了解（即他们的通信模式彼此熟知），那么这些方法中的任何一种都可以是有效的通信模型。想想你在生活中与其他人的关系。您可能认识以各种方式进行交流的人，但如果交流方式众所周知，您可能可以与所有人有效地合作。

> As an example, my partner plans things ahead and doesn’t like unexpected changes. My college roommate, on the other hand, does not like planning and prefers to make decisions at the last moment. I may prefer one communcation style to the other, but I know each of them quite well and can live effectively with either because we can each adjust our communication patterns appropriately.

> 例如，我的搭档提前计划，不喜欢意外的变化。另一方面，我的大学室友不喜欢计划，更喜欢在最后一刻做出决定。我可能更喜欢一种交流方式而不是另一种，但我对它们中的每一种都非常了解，并且可以有效地适应其中的任何一种，因为我们每个人都可以适当地调整我们的交流方式。

Unfortunately, like humans, software can be unreliable. For example, a server may say that it will respond to requests with a future time at which the client should ask again for work to be performed, but the connection between the client and the server could be blocked, causing requests to be dropped. Likewise, a client could be receiving responses saying that work cannot be performed until a future time, but it may continue requesting for the work to be done immediately. For these reasons (and many more that we will not explore today), both server-side and client-side rate limiting are necessary for scalable, resilient systems.

不幸的是，与人类一样，软件也可能不可靠。例如，服务器可能会说它将在客户端应再次请求执行工作的未来时间响应请求，但是客户端和服务器之间的连接可能会被阻止，从而导致请求被丢弃。同样，客户端可能会收到响应，说在未来某个时间之前无法执行工作，但它可能会继续请求立即完成工作。由于这些原因（以及更多我们今天不会探讨的原因），服务器端和客户端的速率限制对于可扩展的、有弹性的系统都是必要的。

Because `controller-runtime` and `client-go` are frameworks to build Kubernetes controllers, which are *clients* of the Kubernetes API server, we will mostly be focusing on client-side rate limiting today.

因为`controller-runtime` 和`client-go` 是构建Kubernetes 控制器的框架，它们是Kubernetes API 服务器的*客户端*，我们今天将主要关注客户端速率限制。

## What’s in a controller?

## 控制器中有什么？

> Skip this section if you are already familiar with the basics of how `controller-runtime` works. 

> 如果您已经熟悉“控制器运行时”工作原理的基础知识，请跳过本节。

`controller-runtime` exposes the [controller abstraction](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/controller/controller.go#L61) to consumers primarily through the execution of a [reconciliation loop](https://book.kubebuilder.io/cronjob-tutorial/controller-implementation.html#implementing-a-controller) that is implemented by the consumer and passed to the framework. Here is an example of a simple [`Reconciler`](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/reconcile/reconcile.go#L89) that could be passed to a ` controller-runtime` controller:

`controller-runtime` 主要通过执行 a 向消费者公开 [控制器抽象](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/controller/controller.go#L61)[协调循环](https://book.kubebuilder.io/cronjob-tutorial/controller-implementation.html#implementing-a-controller) 由消费者实现并传递给框架。这是一个简单的 [`Reconciler`](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/reconcile/reconcile.go#L89) 的例子控制器运行时`控制器：

```go
type Reconciler struct {}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
    fmt.Println("Reconciling something!")
    return reconcile.Result{}, nil
}
```

> You can see a more sophisticated reconciliation loop [here](https://github.com/crossplane/crossplane/blob/5f239dbb7c399a8d544518be2be23ad16f98a71d/internal/controller/pkg/manager/reconciler.go#L216).

> 你可以看到一个更复杂的调节循环 [这里](https://github.com/crossplane/crossplane/blob/5f239dbb7c399a8d544518be2be23ad16f98a71d/internal/controller/pkg/manager/reconciler.go#L216)。

`controller-runtime` accepts this reconciliation loop via a controller [builder pattern](https://en.wikipedia.org/wiki/Builder_pattern) implementation that allows a consumer to specify resources that should trigger this reconciliation to run. Here is an example of a controller that would be triggered by any CRUD operations on a `Pod`:

`controller-runtime` 通过控制器 [builder pattern](https://en.wikipedia.org/wiki/Builder_pattern) 实现接受此协调循环，该实现允许使用者指定应触发此协调运行的资源。下面是一个控制器的例子，它会被 Pod 上的任何 CRUD 操作触发：

```go
ctrl.NewControllerManagedBy(mgr).
    Named("my-pod-controller").
    For(&v1.Pod{}).
    Complete(&Reconciler{})
```

Ignoring the controller manager (`mgr`) for a moment, you can see that we are passing a name for the controller (`my-pod-controller`), the type we want it to reconcile (`v1.Pod`), as well the `&Reconciler{}` that will actually perform, well, the reconciliation. There are other options that we will explore later (and some that we will not) that can be passed into this controller builder to further customize its behavior.

暂时忽略控制器管理器（`mgr`），您可以看到我们正在传递控制器的名称（`my-pod-controller`），我们希望它协调的类型（`v1.Pod`），以及实际执行协调的 `&Reconciler{}`。我们稍后将探索其他选项（有些我们不会），可以将其传递到此控制器构建器以进一步自定义其行为。

Every `Reconciler` is required to implement the `Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error)` method. `controller-runtime` will invoke this method when changes to “watched” objects occur. Information about the object that triggered the reconciliation (in this case a `Pod`) is passed via the [`reconcile.Request`](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/reconcile/reconcile.go#L47). Within the reconciliation loop, the consumer may choose to use that information to get the object from the API server using `client-go` or the [`Client`](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/client/interfaces.go#L100) abstraction exposed by `controller-runtime`. Let’s expand our `Reconciler` a bit more:

每个 `Reconciler` 都需要实现 `Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error)` 方法。当“监视”对象发生更改时，`controller-runtime` 将调用此方法。有关触发协调的对象（在本例中为 `Pod`）的信息通过 [`reconcile.Request`](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7adakg24805/p7adakg4805/p /reconcile/reconcile.go#L47）。在协调循环中，消费者可以选择使用该信息从 API 服务器使用 `client-go` 或 [`Client`](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/client/interfaces.go#L100)由`controller-runtime`公开的抽象。让我们进一步扩展我们的 `Reconciler`：

```go
type Reconciler struct {
    client client.Client
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
    p := &v1.Pod{}
    if err := r.client.Get(ctx, req.NamespacedName, p);err != nil {
        reconcile.Result{}, err
    }
    fmt.Println(p.Name)
    return reconcile.Result{}, nil
}
```

Now we use the [`NamespacedName`](https://github.com/kubernetes/apimachinery/blob/c93b0f84892eb6bcbc80b312ae70729d8168bc7e/pkg/types/namespacedname.go#L27) to get the object that triggered the reconciliation. We may fail to get the object from the API server, in which case we return `err`. Otherwise, we will return `nil`. 

现在我们使用 [`NamespacedName`](https://github.com/kubernetes/apimachinery/blob/c93b0f84892eb6bcbc80b312ae70729d8168bc7e/pkg/types/namespacedname.go#L27) 来获取触发协调的对象。我们可能无法从 API 服务器获取对象，在这种情况下我们会返回 `err`。否则，我们将返回 `nil`。

> Note: when saying that we are making a request to the API server, it doesn’t necessarily mean we are literally making a network call. As mentioned earlier, `controller-runtime` employs [caching](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/cache/cache.go#L41) to make the operations of controllers managed by a controller manager more efficient. We will not go into depth on how caching is performed, but if enough people find this post useful I’ll write one up on that as well.

> 注意：当说我们正在向 API 服务器发出请求时，并不一定意味着我们实际上是在进行网络调用。正如前面提到的，`controller-runtime` 使用[缓存](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/cache/cache.go#L41)来进行控制器的操作由控制器管理器管理更有效率。我们不会深入讨论缓存是如何执行的，但如果有足够多的人认为这篇文章有用，我也会写一篇关于它的文章。

This is as far as we will go with implementing our tiny controller because we have reached a place where we might encounter rate limiting, and we have probably already lost half the readers who already know how to implement a Kubernetes controller and think this is just another tutorial on doing so. If you want to learn more about controller design, take a look at one of the resources below or search on YouTube for one of the countless talks that have been given on the topic.

这就是我们将要实现的微型控制器，因为我们已经到了可能会遇到速率限制的地方，而且我们可能已经失去了一半已经知道如何实现 Kubernetes 控制器并认为这只是另一个问题的读者关于这样做的教程。如果您想了解有关控制器设计的更多信息，请查看以下资源之一或在 YouTube 上搜索有关该主题的无数演讲之一。

- Kubebuilder CronJob Tutorial
   - The canonical “build your first controller” tutorial.
- “A deep dive into Kubernetes controllers” - Bitnami Blog
   - An old but good post that goes a bit beyond “pass this thing to `controller-runtime` and let the magic ensue!".
- Building an Enterprise Infrastructure Control Plane on Kubernetes
   - A tutorial from KubeCon NA 2020 I gave with [Steven Borrelli](https://twitter.com/stevendborrelli)(Mastercard) on building [Crossplane](https://crossplane.io/) providers. We go through building controllers for a GitHub operator and talk about the various things happening behind the scenes. We also [describe `crossplane-runtime`](https://github.com/crossplane/crossplane-runtime), a framework built on top of `controller-runtime` that caters to writing controllers for external infrastructure providers.
- TGI Kubernetes 145: Duck Typing in Kubernetes
   - This is a recent livestream by [Scott Nichols](https://twitter.com/n3wscott) where he talks about some advanced controller patterns used in projects like [Knative](https://knative.dev/), Crossplane, and more. Scott is also leading some [upstream standardization](https://docs.google.com/document/d/1Bud636dMcAQjXe6xfOMBzT0YYqOj1rx3EELxrq2YQv8/edit?usp=sharing) in this area.

- Kubebuilder CronJob 教程
  - 规范的“构建您的第一个控制器”教程。
- “深入了解 Kubernetes 控制器” - Bitnami 博客
  - 一个古老但很好的帖子，它超越了“将这个东西传递给‘控制器运行时’并让魔法随之而来！”。
- 在 Kubernetes 上构建企业基础设施控制平面
  - 来自 KubeCon NA 2020 的教程，我与 [Steven Borrelli](https://twitter.com/stevendborrelli)(Mastercard) 一起提供了关于构建 [Crossplane](https://crossplane.io/) 提供商的教程。我们为 GitHub 操作员构建控制器，并讨论幕后发生的各种事情。我们还 [描述 `crossplane-runtime`](https://github.com/crossplane/crossplane-runtime)，这是一个构建在 `controller-runtime` 之上的框架，用于为外部基础设施提供商编写控制器。
- TGI Kubernetes 145：Kubernetes 中的鸭子打字
  - 这是 [Scott Nichols](https://twitter.com/n3wscott) 最近的直播，在那里他谈到了在 [Knative](https://knative.dev/)、Crossplane、和更多。 Scott 还在该领域领导了一些 [上游标准化](https://docs.google.com/document/d/1Bud636dMcAQjXe6xfOMBzT0YYqOj1rx3EELxrq2YQv8/edit?usp=sharing)。

## What happens when we fail?

## 当我们失败时会发生什么？

This is not an existential question, we are just talking about when we return an error from our reconciliation loop. Like most things in software, and in life, the answer is: “it depends”. Let's take a look at the [`reconcile.Result` struct](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/reconcile/reconcile.go#L27) and see what options we have for telling `controller-runtime` what to do next.

这不是一个存在的问题，我们只是在讨论何时从我们的协调循环中返回错误。就像软件和生活中的大多数事情一样，答案是：“视情况而定”。我们来看看 [`reconcile.Result` 结构体](https://github.com/kubernetes-sigs/controller-runtime/blob/e1a725df2743147795e5dfc8275365f7ada24805/pkg/reconcile/reconcile.go) 看看有什么#L 选项。必须告诉`controller-runtime`接下来要做什么。

```go
// Result contains the result of a Reconciler invocation.
type Result struct {
  // Requeue tells the Controller to requeue the reconcile key.Defaults to false.
  Requeue bool

  // RequeueAfter if greater than 0, tells the Controller to requeue the reconcile key after the Duration.
  // Implies that Requeue is true, there is no need to set Requeue to true at the same time as RequeueAfter.
  RequeueAfter time.Duration
}
```

So we can either say requeue after some period of time, requeue immediately, or don’t requeue at all. Then there is also the `error` we return. The permutations of these values, and how `controller-runtime` will respond to them are outlined below (you can also look directly at the [source](https://github.com/kubernetes-sigs/controller-runtime/blob/16bf3ad036b908d897543c415fcc0bafc5cec711/pkg/internal/controller/controller.go#L297)).

所以我们可以说一段时间后重新排队，立即重新排队，或者根本不重新排队。然后还有我们返回的 `error`。下面概述了这些值的排列以及`controller-runtime` 将如何响应它们（您也可以直接查看 [source](https://github.com/kubernetes-sigs/controller-runtime/blob/16bf3ad036b908d897543c415fcc0bafc5cec711/pkg/internal/controller/controller.go#L297))。

| Requeue | RequeueAfter | Error | Result                                  |
| ------- |------------ |----- |--------------------------------------- |
| any     | any          | !nil  | Requeue with rate limiting. |
| true    | 0            | nil   | Requeue with rate limiting. |
| any     | >0           | nil   | Requeue after specified `RequeueAfter`. |
| false   | 0            | nil   | Do not requeue. | 

|重新排队 | RequeueAfter |错误 |结果 |
| |任何 |任何 | !nil |使用速率限制重新排队。 |
|真实| 0 |零|使用速率限制重新排队。 |
|任何 | >0 |零|在指定的“RequeueAfter”之后重新排队。 |
|假| 0 |零 |不要重新排队。 |

We are primarily interested in the first two cases as the latter two are fairly self explanatory. Both of the first two essentially result in the same outcome (with some difference in what logs and metrics are written). The specific call being made is `c.Queue.AddRateLimited(req)`. This is similar to the `RequeueAfter` result except in that case the call is `c.Queue.AddAfter(req, result.RequeueAfter)` and we are passing the exact time to wait until the request is added back to queue. So how long do we wait in the rate limited case?

我们主要对前两种情况感兴趣，因为后两种情况是不言自明的。前两个本质上都会产生相同的结果（在写入的日志和指标方面存在一些差异）。正在进行的特定调用是`c.Queue.AddRateLimited(req)`。这类似于 `RequeueAfter` 结果，除了在这种情况下调用是 `c.Queue.AddAfter(req, result.RequeueAfter)` 并且我们正在传递等待请求被添加回队列的确切时间。那么在限速的情况下我们要等多久呢？

## The Default Controller Rate Limiter

## 默认控制器速率限制器

Earlier when we talked about rate limiting we honed in on the case where we wanted to be resilient to failed communication between a specific client and a specific server. You’ll notice that `controller-runtime` does not have the luxury of doing the same. In our tiny toy `Reconciler`, we returned an error when we failed to get an object from the API server, but that didn’t have to be the case. It could be true that we instead were returning an error on a condition that did not involve the API server, such as the status of the `Pod` not having a certain condition. `controller-runtime` has to accommodate for general error cases, and because of that, uses a fairly generic rate limiting strategy by default.

早些时候，当我们讨论速率限制时，我们仔细研究了我们希望对特定客户端和特定服务器之间的失败通信具有弹性的情况。你会注意到`controller-runtime` 没有这样做的奢侈。在我们的小玩具“Reconciler”中，当我们无法从 API 服务器获取对象时，我们返回了一个错误，但事实并非如此。确实，我们在不涉及 API 服务器的条件下返回错误，例如“Pod”的状态没有特定条件。 `controller-runtime` 必须适应一般的错误情况，因此，默认情况下使用相当通用的速率限制策略。

The default implementation is borrowed from `client-go` and gets set [during controller construction](https://github.com/kubernetes-sigs/controller-runtime/blob/16bf3ad036b908d897543c415fcc0bafc5cec711/pkg/controller/controller.go#L117) . Let’s hop over to the `client-go` codebase and see what this `workqueue.DefaultControllerRateLimiter()` looks like:

默认实现是从`client-go` 借来的，并在[控制器构建期间]设置(https://github.com/kubernetes-sigs/controller-runtime/blob/16bf3ad036b908d897543c415fcc0bafc5cec711/pkg/controller/controller.go#L117) .让我们跳到 `client-go` 代码库，看看这个 `workqueue.DefaultControllerRateLimiter()` 是什么样的：

```go
// DefaultControllerRateLimiter is a no-arg constructor for a default rate limiter for a workqueue.It has
// both overall and per-item rate limiting.The overall is a token bucket and the per-item is exponential
func DefaultControllerRateLimiter() RateLimiter {
  return NewMaxOfRateLimiter(
    NewItemExponentialFailureRateLimiter(5*time.Millisecond, 1000*time.Second),
    // 10 qps, 100 bucket size.This is only for retry speed and its only the overall factor (not per item)
    &BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
  )
}
```

We can get a pretty good idea of what is going on by just looking at the comments. It [returns a `MaxOfRateLimiter`](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/default_rate_limiters.go#L191), which, as you would guess, takes the value from each limiter passed to it and returns the maximum when the `RateLimiterInterface` [method calls](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/rate_limiting_queue.go#L20) it implements are invoked .

只需查看评论，我们就可以很好地了解正在发生的事情。它[返回一个`MaxOfRateLimiter`](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/default_rate_limiters.go#L191)，它从你猜测的每个值中获取传递给它的限制器并在调用 `RateLimiterInterface` [方法调用](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/rate_limiting_queue.goLs)时返回最大值。 .

The two rate limiters that are passed to it are [an `ItemExponentialFailureRateLimiter`](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/default_rate_limiters.go#L67) and [ a`BucketRateLimiter `](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/default_rate_limiters.go#L48). The comments give us a helpful hint that the former is “per-item” while the latter is only the “overall factor”.

传递给它的两个速率限制器是[一个`ItemExponentialFailureRateLimiter`](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/default_rate_7Bucker)和[a`LimitergoLimiter]。`](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/default_rate_limiters.go#L48)。评论给了我们一个有用的提示，前者是“每个项目”，而后者只是“整体因素”。

Looking at the `ItemExponentialFailureRateLimiter` first, we can see that it internally has a `map` of failures, a base delay, and a max delay (also a `failuresLock` mutex, which is necessary due to the fact that `controller-runtime ` allows for concurrent reconciler invocations for a single controller).

首先查看`ItemExponentialFailureRateLimiter`，我们可以看到它内部有一个失败的`map`、一个基本延迟和一个最大延迟（也是一个`failuresLock`互斥锁，由于`controller-runtime ` 允许对单个控制器进行并发协调器调用）。

```go
// ItemExponentialFailureRateLimiter does a simple baseDelay*2^<num-failures> limit
// dealing with max failures and expiration are up to the caller
type ItemExponentialFailureRateLimiter struct {
  failuresLock sync.Mutex
  failures     map[interface{}]int

  baseDelay time.Duration
  maxDelay  time.Duration
}
```

The struct description goes so far to provide the actual formula that is used for rate limiting: `baseDelay*2^<num-failures>`. This is an extremely common algorithm known as **exponential backoff**. It can be boiled down to “for every consecutive failure we will double the amount of time we wait to try again”. As expected, the amount of time we wait will continue exponentially as long we continue to fail, meaning the wait time could grow *extremely* large. To guard against this, we pass a `maxDelay`, indicating that no matter how bad it gets, we don’t want to wait longer than that value.

到目前为止，结构描述提供了用于速率限制的实际公式：`baseDelay*2^<num-failures>`。这是一种非常常见的算法，称为**指数退避**。可以归结为“对于每一次连续失败，我们将等待再次尝试的时间增加一倍”。正如预期的那样，只要我们继续失败，我们等待的时间就会呈指数增长，这意味着等待时间可能会增长*非常*大。为了防止这种情况，我们传递了一个 `maxDelay`，表明无论它变得多么糟糕，我们都不想等待超过该值。

The other important concept here is that this rate limiter is “per-item”. That means that in our toy controller, if we continuously fail reconciling a `Pod` named `one` in the `default` namespace, and our requeue delay grows exponentially, we will not start off with a huge delay the first time we fail to reconcile the `Pod` named `two` in the `default` namespace.

这里的另一个重要概念是这个速率限制器是“每项”。这意味着在我们的玩具控制器中，如果我们在“default”命名空间中协调名为“one”的“Pod”不断失败，并且我们的重新排队延迟呈指数增长，我们将不会在第一次失败时以巨大的延迟开始协调 `default` 命名空间中名为 `two` 的 `Pod`。

Using the default `baseDelay` of `.005s` and `maxDelay` of `1000s`, we end up with a requeue backoff that looks like this:

使用默认的 `baseDelay` 为 `.005s` 和 `maxDelay` 为 `1000s`，我们最终得到一个如下所示的重新排队退避：

![single-item](https://i.imgur.com/z4eBQTu.png)

Great! If a particular object that we are reconciling constantly causes errors, we will backoff to only trying every `16.67m` (`1000s`), which is quite infrequent. Within the first second, where the delays are the shortest, we are only going to requeue ~7 times.

伟大的！如果我们正在协调的特定对象不断导致错误，我们将退避到仅尝试每`16.67m`（`1000s`），这种情况很少见。在延迟最短的第一秒内，我们只会重新排队约 7 次。

![single-item-req](https://danielmangum.com/static/rate_limit_item_requeue.png)

However, in a large cluster, such as one that is responsible for managing all of your organization's infrastructure as Kubernetes [CustomResourceDefinitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) (as we see quite frequently with Crossplane users), a single controller could be reconciling many objects. When considering rate limiting, we frequently quantify how aggressive a strategy is by the *maximum requests per second* (or in this case maximum requeues). If we only used the per-item rate limiter, that number is unbounded. For example, in the case that 10,000 objects were being watched by a single controller and they all started failing continuously at the same moment, there would be somewhere between 70,000 and 80,000 requeues *within the first second*.

但是，在大型集群中，例如负责将组织的所有基础架构管理为 Kubernetes [CustomResourceDefinitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)（正如我们经常在 Crossplane 用户中看到的那样），单个控制器可以协调许多对象。在考虑速率限制时，我们经常通过*每秒最大请求数*（或在这种情况下最大重新排队)来量化策略的激进程度。如果我们只使用每个项目的速率限制器，那么这个数字是无限的。例如，如果单个控制器正在监视 10,000 个对象，并且它们都在同一时刻开始连续失败，那么*在第一秒内*将有 70,000 到 80,000 个重新排队。

> Note: all of our measurements are approximate as we are not considering execution time of the actual reconciliation and related operations. Additionally, we are not considering enqueues that result from actual changes to the object as reported by informers. However, the concepts still hold, most notably that the upper limit is unbounded.

> 注意：我们所有的测量都是近似的，因为我们没有考虑实际对账和相关操作的执行时间。此外，我们不考虑由告密者报告的对象的实际更改导致的排队。然而，这些概念仍然成立，最值得注意的是上限是无界的。

![req-per-sec](https://danielmangum.com/static/rate_limit_many_requeue.png)

This is where the `BucketRateLimiter` comes into play. It is a wrapper around [`Limiter`](https://github.com/golang/time/blob/7e3f01d253248a0a5694eb5b7376dfea18b6397e/rate/rate.go#L55) from the golang [`golang.org/x/time/rate` ](https://github.com/golang/time/tree/master/rate) package. The helpful documentation on `Limiter` tells us that *“Informally, in any large enough time interval, the Limiter limits the rate to r tokens per second, with a maximum burst size of b events”*. It is an implementation of a [token bucket algorithm](https://en.wikipedia.org/wiki/Token_bucket) and guards against situations such as the worst case failure described above. In `client-go`, the defaults passed to the `BucketRateLimiter` tell us that the bucket starts out with 100 tokens, which is also the maximum it can hold, and will be refilled at a rate of 10 tokens per second. The 100 token limit is known as the *maximum burst* as it the ceiling of requeues that can happen at the same instant. Looking back at our 10,000 objects failing reconciliation at the same time example, using the `BucketRateLimiter` *alone* makes our graph look a lot more friendly.

这就是“BucketRateLimiter”发挥作用的地方。它是来自 golang [`golang.org/x/time/rate` 的 [`Limiter`](https://github.com/golang/time/blob/7e3f01d253248a0a5694eb5b7376dfea18b6397e/rate/rate.go#L55)的包装器](https://github.com/golang/time/tree/master/rate) 包。关于 `Limiter` 的有用文档告诉我们*“非正式地，在任何足够大的时间间隔内，Limiter 将速率限制为每秒 r 个令牌，最大突发大小为 b 个事件”*。它是 [令牌桶算法](https://en.wikipedia.org/wiki/Token_bucket) 的实现，可以防止上述最坏情况失败等情况。在 `client-go` 中，传递给 `BucketRateLimiter` 的默认值告诉我们该存储桶以 100 个令牌开始，这也是它可以容纳的最大值，并且将以每秒 10 个令牌的速度重新填充。 100 个代币的限制被称为*最大突发*，因为它是同一时刻可能发生的重新排队的上限。回顾我们的 10,000 个对象同时协调失败的例子，使用 `BucketRateLimiter` *单独* 使我们的图表看起来更加友好。

![token-bucket](https://danielmangum.com/static/rate_limit_bucket.png)

As with most examples in this post, this is a contrived scenario, but it serves to represent the concept. At `t=0` the first 100 failed reconciliations requeue immediately because the token bucket has a token available, meaning they can be added back to the queue immediately (ie the call to the `BucketRateLimiter`'s `When()` method returns a delay of `0s`). The 101st failure goes to the token bucket only to find it empty, and is told wait for ~`.1s` for the next available token (because our rate is 10 tokens per second). The 102nd failure comes to the bucket also seeing it empty and is told to wait ~`.2s` for the next available token. This delay will increase for every subsequent failure before a new token is added to the bucket. Eventually, the failures will hopefully occur less frequently, meaning that the bucket can refill to its maximum 100 tokens and accommodate future large bursts.

与本文中的大多数示例一样，这是一个人为的场景，但它可以代表这个概念。在 `t=0` 时，前 100 个失败的协调会立即重新排队，因为令牌桶有一个可用的令牌，这意味着它们可以立即添加回队列（即调用 `BucketRateLimiter` 的 `When()` 方法返回'0s' 的延迟）。第 101 次失败进入令牌桶后发现它是空的，并被告知等待 ~`.1s` 以获取下一个可用令牌（因为我们的速率是每秒 10 个令牌）。第 102 次失败出现在存储桶中，也看到它是空的，并被告知等待 ~`.2s` 以获取下一个可用令牌。在将新令牌添加到存储桶之前，每次后续失败都会增加此延迟。最终，失败的发生频率有望降低，这意味着存储桶可以重新填充到其最大 100 个令牌并适应未来的大爆发。

The token bucket rate limiting strategy also guards against a [thundering herd](https://en.wikipedia.org/wiki/Thundering_herd_problem) problem, where all items are requeued on the same schedule, meaning that massive spikes will occur at backoff intervals . Simply increasing the delay on the per-item rate limiter will not provide the same protection because, while there may be longer delays between requeues, all items may requeue at the same future time, causing massive pressure on the client and server.

令牌桶速率限制策略还可以防止出现 [thundering herd](https://en.wikipedia.org/wiki/Thundering_herd_problem) 问题，即所有项目都按照相同的时间表重新排队，这意味着会以退避间隔出现大规模峰值.简单地增加 per-item 速率限制器的延迟不会提供相同的保护，因为虽然重新排队之间可能会有更长的延迟，但所有 item 可能会在同一时间重新排队，从而对客户端和服务器造成巨大压力。

So now that we have seen `ItemExponentialFailureRateLimiter` and the `BucketRateLimiter` work in isolation, you are likely wondering what it looks like when they are used in tandem with the `MaxOfRateLimiter`. In the worst case scenario we have been examining, it looks a lot like a slightly smoother version of the `BucketRateLimiter` graph. Why is that the case? Well, for the first 100 failures that come to the bucket, they are going to be told that they can requeue immediately by the `BucketRateLimiter`, but the `ItemExponentialFailureRateLimiter` is going to say “not so fast” because `0.005 * 2^ 0 = 0.005`, is greater than `0`, so it will be selected by the `MaxOfRateLimiter` as the delay value to return.

因此，既然我们已经看到了 `ItemExponentialFailureRateLimiter` 和 `BucketRateLimiter` 独立工作，您可能想知道当它们与 `MaxOfRateLimiter` 一起使用时会是什么样子。在我们一直在研究的最坏情况下，它看起来很像“BucketRateLimiter”图的稍微平滑的版本。为什么会这样？好吧，对于到达存储桶的前 100 个失败，他们将被告知他们可以立即通过 `BucketRateLimiter` 重新排队，但是 `ItemExponentialFailureRateLimiter` 会说“不那么快”，因为 `0.005 * 2^ 0 = 0.005`，大于`0`，所以会被`MaxOfRateLimiter`选择作为延迟值返回。

Similarly, in a different scenario where the 101st failure occurs on an item that has already failed 5 or more times, the `BucketRateLimiter` will say that it needs to wait `.1s` for a token to be available, but the `ItemExponentialFailureRateLimiter` will say it needs to wait longer than that (`0.005 * 2^5 = 0.16`). The `MaxOfRateLimiter` will select the latter.

类似地，在第 101 次失败发生在已经失败 5 次或更多次的项目上的不同场景中，`BucketRateLimiter` 会说它需要等待 `.1s` 以获得可用的令牌，但是 `ItemExponentialFailureRateLimiter`会说它需要等待更长的时间（`0.005 * 2^5 = 0.16`）。 `MaxOfRateLimiter` 将选择后者。

However, in most cases, there will not be 10,000 unique objects all returning failures at the same instant, and the requeues per second will look like a curve falling somewhere beneath the maximum controller requeue limit established by the `BucketRateLimiter`.

但是，在大多数情况下，不会有 10,000 个唯一对象在同一时刻都返回故障，并且每秒重新排队看起来像一条曲线，低于由“BucketRateLimiter”建立的最大控制器重新排队限制。

## Using Your Own Rate Limiter 

## 使用您自己的速率限制器

Now that we have explored the default implementation, it is time to consider whether it is acceptable for your use case. If not, you can [supply your own rate limiter](https://github.com/kubernetes-sigs/controller-runtime/issues/631). The controller builder we used when constructing our toy controller allows us to override the `RateLimiter` controller option.

现在我们已经探索了默认实现，是时候考虑它是否适合您的用例。如果没有，您可以[提供您自己的速率限制器](https://github.com/kubernetes-sigs/controller-runtime/issues/631)。我们在构建玩具控制器时使用的控制器构建器允许我们覆盖 `RateLimiter` 控制器选项。

```go
ctrl.NewControllerManagedBy(mgr).
    Named("my-pod-controller").
    WithOptions(controller.Options{
        RateLimiter: coolpackage.NewAwesomeRateLimiter(),
    }).
    For(&v1.Pod{}).
    Complete(&Reconciler{})
```

A few examples you of why you might want to pass your own rate limiter include:

您可能希望通过自己的速率限制器的一些示例包括：

- Overriding the default values passed to the rate limiters that `controller-runtime` uses by default. For instance, you may want a more or less aggressive per-item strategy.
- Using a different rate limiting strategy entirely. `client-go` has [an `ItemFastSlowRateLimiter`](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/default_rate_limiters.go#L125) that could make more sense in certain scenarios.
- Rate limiting at a different level. Maybe you have many controllers all hitting the same external API and you want to limit the amount of times per second that can happen at the controller manager level. Passing the same `BucketRateLimiter` to all controllers would provide that effect.

- 覆盖传递给“controller-runtime”默认使用的速率限制器的默认值。例如，您可能想要一个或多或少激进的每项策略。
- 完全使用不同的速率限制策略。 `client-go` 有 [一个 `ItemFastSlowRateLimiter`](https://github.com/kubernetes/client-go/blob/20732a1bc198ab57de644af498fa75e73fa44c08/util/workqueue/default_rate_limiters.go#L125)在某些情况下可能更有意义)
- 不同级别的速率限制。也许您有许多控制器都访问相同的外部 API，并且您想限制控制器管理器级别每秒可能发生的次数。将相同的 `BucketRateLimiter` 传递给所有控制器将提供这种效果。

## Wrapping Up

##  包起来

As Kubernetes controllers expand to cover more and more use cases, the default happy path is not necessarily the right answer in every scenario. Using different patterns, or even exposing the knobs to end users of the controllers so they can tune them to their environments, might make sense.

随着 Kubernetes 控制器扩展以涵盖越来越多的用例，默认的快乐路径不一定在每个场景中都是正确的答案。使用不同的模式，甚至将旋钮暴露给控制器的最终用户，以便他们可以根据自己的环境调整它们，可能是有意义的。

Send me a message [@hasheddan](https://twitter.com/hasheddan) on Twitter for any questions or comments! 

如有任何问题或意见，请在 Twitter 上向我发送消息 [@hasheddan](https://twitter.com/hasheddan)！

