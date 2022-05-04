# Implementing Graceful Shutdown in Go

# 在 Go 中实现优雅关闭

From: https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/

Tech Lead at RudderStack

RudderStack 技术主管

Shutting down gracefully is important for any long lasting process, especially  for one that handles some kind of state. For example, what if you wanted to shutdown the database that supports your application and the db  process didn't flush the current state to the disk, or what if you  wanted to shut down a web server with thousands of connections but  didn't wait for the requests to finish Not only does shutting down  gracefully positively affect the user experience, it also eases internal operations, leading to happier engineers and less stressed SREs.

优雅地关闭对于任何持久的过程都很重要，尤其是对于处理某种状态的过程。例如，如果您想关闭支持您的应用程序的数据库并且 db 进程没有将当前状态刷新到磁盘，或者如果您想关闭具有数千个连接但没有等待的 Web 服务器怎么办让请求完成 优雅地关闭不仅对用户体验产生积极影响，还简化了内部操作，从而使工程师更快乐，SRE 压力更小。

To shutdown gracefully is for the program to terminate after:

正常关闭是程序在以下情况下终止：

- All pending processes (web request, loops) are completed - no new processes should start and no new web requests should be accepted.
- Closing all open connections to external services and databases.

- 所有待处理的进程（Web 请求、循环）都已完成 - 不应启动新进程，也不应接受新的 Web 请求。
- 关闭与外部服务和数据库的所有打开连接。

There are a couple of things we must figure out in order to shutdown gracefully:

为了优雅地关闭，我们必须弄清楚几件事：

- **When should we shutdown** *-* Are all pending processes completed, and how we can know this? What if a processes is stuck?
- **How we communicate with processes** - The previous task requires some kind of communication. This is  especially true if we are building a modern, asynchronous, and highly  concurrent application. So, how can we tell them to shutdown and also  know when they've done that?

- **我们什么时候应该关闭** *-* 所有待处理的进程是否都已完成，我们如何知道这一点？如果一个进程卡住了怎么办？
- **我们如何与进程通信** - 上一个任务需要某种通信。如果我们正在构建一个现代、异步和高度并发的应用程序，则尤其如此。那么，我们如何告诉他们关闭并知道他们何时关闭？

When I started looking into shutdown at RudderStack, I saw a number of anti patterns that we were following—for example using *os.Exit(1)* (more on this later)—and decided it was time to implement a graceful shutdown mechanism for [Rudder Server](https://github.com/rudderlabs/rudder-server/). At RudderStack we are building an important part of the modern data  stack. RudderStack is responsible for capturing, processing, and  delivering data to important parts of a company's infrastructure. So,  making sure everything is predictable and ensuring there is no chance  for data loss whenever we have to interact with a service is incredibly  important. This gave me two main goals with graceful shutdown:

当我开始研究 RudderStack 的关闭时，我看到了一些我们正在遵循的反模式——例如使用 *os.Exit(1)*（稍后会详细介绍）——并决定是时候实现一个优雅的关闭机制了用于 [Rudder Server](https://github.com/rudderlabs/rudder-server/)。在 RudderStack，我们正在构建现代数据堆栈的重要组成部分。 RudderStack 负责捕获、处理和交付数据到公司基础设施的重要部分。因此，确保一切都是可预测的，并确保在我们必须与服务交互时不会丢失数据是非常重要的。这给了我两个优雅关闭的主要目标：

1. Ensure that no data loss can happen during a shutdown.
2. Introduce better service control to enable integration testing.

1. 确保关机期间不会发生数据丢失。
2. 引入更好的服务控制来实现集成测试。

Rudder Server is written in Go and my initial research on how to properly  implement graceful shutdown didn't return much information. So, I  decided to publish my experience in implementing this pattern on Rudder  Server.

Rudder Server 是用 Go 编写的，我对如何正确实现优雅关闭的初步研究并没有返回太多信息。因此，我决定发布我在 Rudder Server 上实现此模式的经验。

In this post you'll find a number of anti patterns and  learn how to make exiting a graceful process with a couple of different  approaches. I'll also include a number of examples for common libraries  and some advanced patterns. Let's dive in.

在这篇文章中，您将找到许多反模式，并学习如何使用几种不同的方法来使退出过程变得优雅。我还将包含一些常见库和一些高级模式的示例。让我们潜入水中。

## Anti-patterns

## 反模式

### Block artificially

### 人为拦截

The first anti-pattern is the idea of blocking the main go routine without  actually waiting on anything. Here's an example toy implementation:

第一个反模式是阻塞主 go 例程而不实际等待任何东西的想法。这是一个示例玩具实现：

```text
func KeepProcessAlive() {
    var ch chan int
    <-ch
}

func main() {
    ...
    KeepProcessAlive()
}
```

### os.Exit()

### os.Exit()

Calling `os.Exit(1)` while other go routines are still running is essentially equal to  SIGKILL, no chance for closing open connections and finishing inflight  requests and processing.

在其他 go 例程仍在运行时调用 `os.Exit(1)` 本质上等于 SIGKILL，没有机会关闭打开的连接并完成飞行中的请求和处理。

```text
go func() {
        <-ch
        os.Exit(1)
}()

go func () {

    for ... {

    }
}()
```

## How to make it graceful in Go

## 如何在 Go 中使其优雅

In order to gracefully shutdown a service there are two things you need to understand:

为了优雅地关闭服务，您需要了解两件事：

1. How to wait for all the running go routines to exit
2. How to propagate the termination signal to multiple go routines

1.如何等待所有正在运行的goroutines退出
2.如何将终止信号传播到多个goroutine

Go provides all the tools we need to properly implement (1) and (2). Let's take a look at these in more detail.

Go 提供了我们正确实现 (1) 和 (2) 所需的所有工具。让我们更详细地看一下这些。

### Wait for go-routines to finish

### 等待 go-routines 完成

Go provides sufficient ways for controlling concurrency. Let's see what options are available on waiting go routines.

Go 提供了足够的方法来控制并发性。让我们看看等待执行例程中有哪些可用选项。

#### Using channel

#### 使用频道

Simplest solution, using channel primitive.

最简单的解决方案，使用通道原语。

1. We create an empty struct channel `make(chan struct{}, 1)` (empty struct requires no memory).
2. Every child go routine should **publish to the channel when it is done** (defer can be useful here).
3. The parent go routine should **consume from the channel as many times as the expected go routines**.

1.我们创建一个空的struct channel `make(chan struct{}, 1)`（空的struct不需要内存）。
2. 每个孩子的 goroutine 都应该**完成后发布到频道**（延迟在这里很有用）。
3. 父 goroutine 应该**从频道消费的次数与预期的 goroutines 一样多**。

The example can clear things up:

这个例子可以澄清事情：

```text
func run(ctx) {
  wait := make(chan struct{}, 1)

    go func() {
        defer func() {
        wait <- struct{}{}
        }()
        for {
            select {
        case <-ctx.Done():
                fmt.Println("Break the loop")
                break;
            case <-time.After(1 * time.Second):
                fmt.Println("Hello in a loop")
            }
        }
    }()

    go func() {
        defer func() {
        wait <- struct{}{}
        }()
        for {
            select {
        case <-ctx.Done():
                fmt.Println("Break the loop")
                break;
            case <-time.After(1 * time.Second):
                fmt.Println("Ciao in a loop")
            }
        }
    }()

    // wait for two goroutines to finish
    <-wait
    <-wait

    fmt.Println("Main done")
}
```

*Note: This is mostly useful when waiting on a single go routine.*

*注意：这在等待单次执行例程时最有用。*

#### With WaitGroup

#### 使用 WaitGroup

The channel solution can be a bit ugly, especially with multiple go routines.

通道解决方案可能有点难看，尤其是在有多个 go 例程的情况下。

[sync.WaitGroup](https://pkg.go.dev/sync#WaitGroup/) is a standard library package, that can be used as a more idiomatic way to achieve the above.

[sync.WaitGroup](https://pkg.go.dev/sync#WaitGroup/) 是一个标准库包，可以用作实现上述目标的更惯用的方式。

You can also see another [example of waitgroups](https://gobyexample.com/waitgroups/) in use.

您还可以查看正在使用的另一个 [waitgroups 示例](https://gobyexample.com/waitgroups/)。

```text
func run(ctx) {
    var wg sync.WaitGroup

    wg.Add(1)
  go func() {
        defer wg.Done()
        for {
            select {
        case <-ctx.Done():
                fmt.Println("Break the loop")
                return;
            case <-time.After(1 * time.Second):
                fmt.Println("Hello in a loop")
            }
        }
    }()

  wg.Add(1)
    go func() {
        defer wg.Done()
        for {
            select {
        case <-ctx.Done():
                fmt.Println("Break the loop")
                return;
            case <-time.After(1 * time.Second):
                fmt.Println("Ciao in a loop")
            }
        }
    }()

  wg.Wait()
    fmt.Println("Main done")
}
```

#### With errgroup

#### 带错误组

The [`sync/errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup/) package exposes a better way to do this.

[`sync/errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup/) 包提供了一种更好的方法来执行此操作。

- The two `errgroup`'s methods `.Go` and `.Wait` are more readable and easier to maintain in comparison to `WaitGroup`.
- In addition, as its name suggests it does error propagation and cancels  the context in order to terminate the other go-routines in case of an  error.

- 与 `WaitGroup` 相比，`errgroup` 的两个方法`.Go` 和 `.Wait` 更具可读性和更易于维护。
- 此外，顾名思义，它会进行错误传播并取消上下文，以便在发生错误时终止其他 go-routines。

```text
func run(ctx) {
    g, gCtx := errgroup.WithContext(ctx)
    g.Go(func() error {
        for {
            select {
        case <-gCtx.Done():
                fmt.Println("Break the loop")
                return nil;
            case <-time.After(1 * time.Second):
                fmt.Println("Hello in a loop")
            }
        }
    })

    g.Go(func() error {
        for {
            select {
        case <-gCtx.Done():
                fmt.Println("Break the loop")
                return nil;
            case <-time.After(1 * time.Second):
                fmt.Println("Ciao in a loop")
            }
        }
    }()

  err := g.Wait()
    if err != nil {
        fmt.Println("Error group: ", err)
    }
    fmt.Println("Main done")
}
```

## Terminating a process

## 终止进程

Even if we have figured out how to properly communicate the state of  processes and wait for them, we still have to implement termination. Let's see how this can be done with a simple example, introducing all  the necessary Go primitives.

即使我们已经弄清楚如何正确地传达进程的状态并等待它们，我们仍然必须实现终止。让我们通过一个简单的例子来看看如何做到这一点，介绍所有必要的 Go 原语。

Let's start with a very simple "Hello in a loop" example:

让我们从一个非常简单的“循环中的Hello”示例开始：

```text
func main() {
    for {
        time.Sleep(1 * time.Second)
        fmt.Println("Hello in a loop")
    }
}
```

### Introducing signal handling

### 介绍信号处理

Listen for an OS signal to stop the progress:

监听 OS 信号以停止进度：

```text
exit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
```

- We need to use `os.Interrupt` to gracefully shutdown on `Ctrl+C` which is **SIGINT**
- `syscall.**SIGTERM**` is the usual signal for termination and the default one (it can be [modified](https://docs.docker.com/engine/reference/builder/#stopsignal/)) for [docker](https://docs.docker.com/engine/reference/commandline/stop/) containers, which is also used by [kubernetes](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination/).
- Read more about `signal` in the [package documentation](https://pkg.go.dev/os/signal/) and [go by example](https://gobyexample.com/signals/).

- 我们需要使用 `os.Interrupt` 在 `Ctrl+C` 上正常关闭，即 **SIGINT**
- `syscall.**SIGTERM**` 是终止的常用信号和默认信号（可以[修改](https://docs.docker.com/engine/reference/builder/#stopsignal/）)[docker](https://docs.docker.com/engine/reference/commandline/stop/) 容器，[kubernetes](https://kubernetes.io/docs/concepts/workloads/pods/吊舱生命周期/#吊舱终止/)。
- 在 [package documentation](https://pkg.go.dev/os/signal/) 和 [go by example](https://gobyexample.com/signals/) 中阅读有关 `signal` 的更多信息。

### Breaking the loop

### 打破循环

Now that we have a way to capture signals, we need to find a way to interrupt the loop.

现在我们有了捕获信号的方法，我们需要找到一种方法来中断循环。

#### Non-Blocking Channel Select

#### 非阻塞频道选择

`select` gives you the ability to consume from multiple channels in each `case`.

`select` 使您能够在每个 `case` 中从多个渠道进行消费。

You can review the following resources to get a better understanding:

您可以查看以下资源以获得更好的理解：

- [https://gobyexample.com/non-blocking-channel-operations](https://gobyexample.com/non-blocking-channel-operations/)
- [https://tour.golang.org/concurrency/5](https://tour.golang.org/concurrency/5/)
- [https://gobyexample.com/timeouts](https://gobyexample.com/timeouts/)

- [https://gobyexample.com/non-blocking-channel-operations](https://gobyexample.com/non-blocking-channel-operations/)
- [https://tour.golang.org/concurrency/5](https://tour.golang.org/concurrency/5/)
- [https://gobyexample.com/timeouts](https://gobyexample.com/timeouts/)

Our simple hello for loop, now stops on termination signal:

我们简单的 hello for 循环，现在在终止信号处停止：

```text
func main() {
    c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    
    for {
        select {
    case <-c:
            fmt.Println("Break the loop")
            return;
        case <-time.After(1 * time.Second):
            fmt.Println("Hello in a loop")
        }
    }
}
```

***Note:** We had to change the* `time.Sleep(1 * time.Second)` *to* `time.After(1 * time.Second)`

***注意：** 我们必须将 * `time.Sleep(1 * time.Second)` *to* `time.After(1 * time.Second)`

### How to do it using Context

### 如何使用上下文来做到这一点

Context is a very useful interface in go, that should be used and propagated in all blocking functions. It enables the propagation of cancelation  throughout the program. 

Context 是 go 中一个非常有用的接口，应该在所有阻塞函数中使用和传播。它可以在整个程序中传播取消。

It is considered good practice for `ctx context.Context` to be the first argument in every method or function that is used directly or indirectly for external dependencies.

`ctx context.Context` 作为直接或间接用于外部依赖项的每个方法或函数中的第一个参数被认为是一种很好的做法。

A very detailed article about context: [https://go.dev/blog/context](https://go.dev/blog/context/)

关于上下文的一篇非常详细的文章：[https://go.dev/blog/context](https://go.dev/blog/context/)

### Channel sharing issue

### 频道分享问题

Let's examine how context properties could help in a more complex situation.

让我们研究一下上下文属性如何在更复杂的情况下提供帮助。

*Having multiple loops running in parallel, using channels (counter-example):*

*使用通道并行运行多个循环（反例）：*

```text
// COUNTER EXAMPLE, DO NOT USE THIS CODE
func main() {
    exit := make(chan os.Signal, 1)
    signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
    
  // This will not work as expected!!
    var wg sync.WaitGroup

    wg.Add(1)
  go func() {
        defer wg.Done()
        for {
            select {
        case <-exit: // Only one go routine will get the termination signal
                fmt.Println("Break the loop: hello")
                break;
            case <-time.After(1 * time.Second):
                fmt.Println("Hello in a loop")
            }
        }
    }()

    wg.Add(1)
  go func() {
        defer wg.Done()
        for {
            select {
        case <-exit: // Only one go routine will get the termination signal
                fmt.Println("Break the loop: ciao")
                break;
            case <-time.After(1 * time.Second):
                fmt.Println("Ciao in a loop")
            }
        }
    }()

    wg.Wait()
    fmt.Println("Main done")
}
```

*Why is this not going to work?*

*为什么这不起作用？*

Go channels do not work in a **broadcast** way, only one go routine will receive a single `os.Signal`. Also, there is no guarantee which go routine will receive it.

Go 频道不能以 **broadcast** 方式工作，只有一个 go 例程会收到一个 `os.Signal`。此外，不能保证哪个 goroutine 会收到它。

```
wait := make(chan struct{}{}, 2)
```

Context can help us make the above work, let's see how.

上下文可以帮助我们完成上述工作，让我们看看如何。

#### Using Context for termination

#### 使用上下文终止

Let's try to fix this problem by introducing [`context.WithCancel`](https://pkg.go.dev/context#WithCancel/)

让我们尝试通过引入 [`context.WithCancel`](https://pkg.go.dev/context#WithCancel/) 来解决这个问题

```text
func main() {
  ctx, cancel := context.WithCancel(context.Background())

    go func() {
        exit := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt, syscall.SIGTERM)
        cancel()
    }()

    var wg sync.WaitGroup

    wg.Add(1)
  go func() {
        defer wg.Done()
        for {
            select {
        case <-ctx.Done():
                fmt.Println("Break the loop")
                break;
            case <-time.After(1 * time.Second):
                fmt.Println("Hello in a loop")
            }
        }
    }()

    wg.Add(1)
  go func() {
        defer wg.Done()
        for {
            select {
        case <-ctx.Done():
                fmt.Println("Break the loop")
                break;
            case <-time.After(1 * time.Second):
                fmt.Println("Ciao in a loop")
            }
        }
    }()

    wg.Wait()
    fmt.Println("Main done")
}
```

Essentially the `cancel()` is broadcasted to all the go-routines that call `.Done()`.

本质上，`cancel()` 被广播到所有调用 `.Done()` 的 go-routines。

*The returned context's Done channel is closed when the returned cancel  function is called or when the parent context's Done channel is closed,  whichever happens first.*

*返回的上下文的完成通道在调用返回的取消函数或父上下文的完成通道关闭时关闭，以先发生者为准。*

### NotifyContext

### 通知上下文

In go 1.16 a new helpful method was introduced in signal package, [singal.NotifyContext](https://pkg.go.dev/os/signal#NotifyContext/):

在 go 1.16 中，在信号包中引入了一种新的有用方法，[singal.NotifyContext](https://pkg.go.dev/os/signal#NotifyContext/)：

```text
func NotifyContext(parent context.Context, signals ...os.Signal) (ctx context.Context, stop context.CancelFunc)
```

*NotifyContext returns a copy of the parent context that is marked done (its Done  channel is closed) when one of the listed signals arrives, when the  returned stop function is called, or when the parent context's Done  channel is closed, whichever happens first. *

*NotifyContext 在列出的信号之一到达、调用返回的停止函数或父上下文的 Done 通道关闭时（以先发生者为准）返回标记为完成的父上下文的副本（其完成通道已关闭）。 *

Using NotifyContext can simplify the example above to:

使用 NotifyContext 可以将上面的示例简化为：

```undefined
func main() {
  ctx, stop := context.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
  defer stop()

    var wg sync.WaitGroup

    wg.Add(1)
  go func() {
        defer wg.Done()
        for {
            select {
        case <-ctx.Done():
                fmt.Println("Break the loop")
                break;
            case <-time.After(1 * time.Second):
                fmt.Println("Hello in a loop")
            }
        }
    }()

    wg.Add(1)
  go func() {
        defer wg.Done()
        for {
            select {
        case <-ctx.Done():
                fmt.Println("Break the loop")
                break;
            case <-time.After(1 * time.Second):
                fmt.Println("Ciao in a loop")
            }
        }
    }()

    wg.Wait()
    fmt.Println("Main done")
}
```

*A full working example can be found under our [example repo](https://github.com/rudderlabs/graceful-shutdown-examples/tree/main/signal/)*

*可以在我们的 [example repo](https://github.com/rudderlabs/graceful-shutdown-examples/tree/main/signal/) 下找到完整的工作示例*

## Common libraries

## 常用库

### HTTP server

### HTTP 服务器

The examples above included a `for` loop for simplification, but let's examine something more practical.

上面的例子包括一个用于简化的`for`循环，但让我们来看看更实用的东西。

During a non-graceful shutdown, inflight HTTP requests could face the following issues:

在非正常关闭期间，正在进行的 HTTP 请求可能会遇到以下问题：

- They never get a response back, so they timeout.
- Some progress has been made, but it is interrupted halfway, causing a waste  of resources or data inconsistencies if transactions are not used  properly. 

- 他们永远不会得到回应，所以他们超时。
- 取得了一些进展，但中途中断，如果交易使用不当，造成资源浪费或数据不一致。

- A connection to an external dependency is closed by another go routine, so the request can not progress further.

- 与外部依赖项的连接被另一个 go 例程关闭，因此请求无法进一步进行。

*⚠️ **Having your HTTP server shutting down gracefully is really important.** In a cloud-native environment services/pods shutdown multiple times  within a day either for autoscaling, applying a configuration, or  deploying a new version of a service. Thus, the impact of interrupted or timeout requests can be significant in the service's SLAs.*

*⚠️ **让您的 HTTP 服务器正常关闭非常重要。** 在云原生环境中，服务/pod 在一天内多次关闭，以自动扩展、应用配置或部署新版本的服务。因此，中断或超时请求的影响可能会对服务的 SLA 产生重大影响。*

Fortunately, go provides a way to gracefully shutdown an HTTP server.

幸运的是，go 提供了一种优雅地关闭 HTTP 服务器的方法。

Let us see how it's done:

让我们看看它是如何完成的：

```text
func main() {

    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
        signal.Notify(c, os.Interrupt, syscall.SIGTERM)

        <-c
        cancel()
    }()

    db, err := repo.SetupPostgresDB(ctx, getConfig("DB_DSN", "root@tcp(127.0.0.1:3306)/service"))
    if err != nil {
        panic(err)
    }

    httpServer := &http.Server{
        Addr:    ":8000",
    }

    g, gCtx := errgroup.WithContext(ctx)
    g.Go(func() error {
        return httpServer.ListenAndServe()
    })
    g.Go(func() error {
        <-gCtx.Done()
        return httpServer.Shutdown(context.Background())
    })

    if err := g.Wait();err != nil {
        fmt.Printf("exit reason: %s \n", err)
    }
}
```

We are using two go routines:

我们正在使用两个 go 例程：

1. run **`httpServer.ListenAndServe()`** as usual
2. wait for `<-gCtx.Done()` and then call **`httpServer.Shutdown(context.Background())`**

1. 像往常一样运行**`httpServer.ListenAndServe()`**
2.等待`<-gCtx.Done()`然后调用**`httpServer.Shutdown(context.Background())`**

It is important to read the package documentation in order to understand how this works:

阅读软件包文档以了解其工作原理非常重要：

> Shutdown gracefully shuts down the server **without interrupting any active connections**.

> Shutdown 优雅地关闭服务器**而不会中断任何活动连接**。

Nice, but how?

很好，但是怎么做？

> Shutdown works by first closing all open listeners, then closing all idle  connections, and then waiting indefinitely for connections to return to  idle and then shut down.

> 关闭首先关闭所有打开的侦听器，然后关闭所有空闲连接，然后无限期地等待连接返回空闲，然后关闭。

Why do I have to provide a context?

为什么我必须提供上下文？

> If the provided context expires before the shutdown is complete, Shutdown  returns the context's error, otherwise it returns any error returned  from closing the Server's underlying Listener(s).

> 如果提供的上下文在关闭完成之前过期，则关闭返回上下文的错误，否则返回关闭服务器的底层侦听器返回的任何错误。

In the example, we chose to provide **`context.Background()`** which has no expiration.

在示例中，我们选择提供没有过期时间的**`context.Background()`**。

#### Canceling long running requests

#### 取消长时间运行的请求

When `.Shutdown` is method is called the serve stop accepting new connections and it waits for existing once to finish before `.ListenAndServe()` may return.

当调用 .Shutdown 方法时，服务停止接受新连接，并等待现有的一次完成，然后 .ListenAndServe() 可能会返回。

There are cases where http requests require quite a long time to be  terminated. That could be a for instance a long running job or a  websocket connection.

在某些情况下，http 请求需要很长时间才能终止。例如，这可能是一个长时间运行的作业或一个 websocket 连接。

So, what is the best way to terminate those gracefully and not hang waiting for them to finish?

那么，优雅地终止这些而不是挂起等待它们完成的最佳方法是什么？

The answer comes into two parts:

答案分为两部分：

1. First of all you need to extract the context from http.Request `ctx := req.Context()` and use this context to terminate your long running process.
2. Use [BaseContext](https://pkg.go.dev/net/http#Server/) (introduced in go1.13), to pass your main ctx as the context in every request

1. 首先，您需要从 http.Request `ctx := req.Context()` 中提取上下文，并使用此上下文来终止您长时间运行的进程。
2. 使用 [BaseContext](https://pkg.go.dev/net/http#Server/)（在 go1.13 中引入)，在每个请求中传递你的主 ctx 作为上下文

> BaseContext optionally specifies a function that returns the base context for incoming requests on this server.
> The provided Listener is the specific Listener that's about to start accepting requests.
> If BaseContext is nil, the default is context.Background().
> If non-nil, it must return a non-nil context.

> BaseContext 可选地指定一个函数，该函数返回此服务器上传入请求的基本上下文。
> 提供的侦听器是即将开始接受请求的特定侦听器。
> 如果 BaseContext 为 nil，则默认为 context.Background()。
> 如果非零，它必须返回一个非零上下文。

In the example bellow, a dummy http handler keeps printing in stdout `Hello in a loop`, it will stop either when the request is canceled or the instance receives a termination signal.

在下面的示例中，一个虚拟的 http 处理程序在 stdout `Hello in a loop` 中保持打印，当请求被取消或实例接收到终止信号时，它将停止。

```text
func main() {
    mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    httpServer := &http.Server{
        Addr: ":8000",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()

            for {
                select {
                case <-ctx.Done():
                    fmt.Println("Graceful handler exit")
                    w.WriteHeader(http.StatusOK)
                    return
                case <-time.After(1 * time.Second):
                    fmt.Println("Hello in a loop")
                }
            }
        }),
        BaseContext: func(_ net.Listener) context.Context {
            return mainCtx
        },
    }
    g, gCtx := errgroup.WithContext(mainCtx)
    g.Go(func() error {
        return httpServer.ListenAndServe()
    })
    g.Go(func() error {
        <-gCtx.Done()
        return httpServer.Shutdown(context.Background())
    })

    if err := g.Wait();err != nil {
        fmt.Printf("exit reason: %s \n", err)
    }
}
```

A full working example can be found under our [example repo](https://github.com/rudderlabs/graceful-shutdown-examples/tree/main/httpserver/), feel free to experiment by commenting out `BaseContextor` or ` httpServer.Shutdown.`

一个完整的工作示例可以在我们的 [example repo](https://github.com/rudderlabs/graceful-shutdown-examples/tree/main/httpserver/) 下找到，随时通过注释掉 `BaseContextor` 或 ` httpServer.Shutdown.`

### HTTP Client

### HTTP 客户端

Go standard libraries provides a way to pass a context when making an HTTP request: [NewRequestWithContext](https://pkg.go.dev/net/http#NewRequestWithContext/)

Go 标准库提供了一种在发出 HTTP 请求时传递上下文的方法：[NewRequestWithContext](https://pkg.go.dev/net/http#NewRequestWithContext/)

Let's see how the following code can be refactored to use it:

让我们看看如何重构以下代码以使用它：

```text
resp, err := netClient.Post(uri, "application/json; charset=utf-8",
                    bytes.NewBuffer(payload))
...
```

The equivalent with passing ctx:

通过 ctx 的等价物：

```text
req, err := http.NewRequestWithContext(ctx, "POST", uri, bytes.NewBuffer(payload))
if err != nil {
            return err
}
req.Header.Set("Content-Type", "application/json; charset=utf-8")

resp, err := netClient.Do(req)
...
```

The following techniques are necessary for more advanced use cases. For  instance, if you are using a pool of workers or you have a chain of  component dependencies that need to shutdown in order.

对于更高级的用例，以下技术是必需的。例如，如果您正在使用一个工作人员池，或者您有一个需要按顺序关闭的组件依赖链。

### Draining Worker Channels

### 排水工人渠道

When you have worker go routines that are consuming/producing from/to a  channel, special care must be taken to make sure no items are left in  the channels when the process shuts down. To do this we need to utilize  go `close` method on the channel. Here's a great overview on [closing channels](https://gobyexample.com/closing-channels/), and a more [advanced article](https://go101.org/article/channel-closing.html/) on the topic.

当您有工作程序从/向通道消费/生产时，必须特别注意确保在进程关闭时没有项目留在通道中。为此，我们需要在通道上使用 go `close` 方法。以下是关于 [关闭频道](https://gobyexample.com/closing-channels/) 的精彩概述，以及更多 [高级文章](https://go101.org/article/channel-closing.html/)话题。

Two things to remember about closing a channel:

关于关闭频道要记住两件事：

- Writing to a close channel will result in a panic
- When reading for a channel, you can use `value, ok <- ch` . Reading from a close channel will return all the buffered items. Once the buffer items are "drained", the channel will return zero `value` and `ok` will be false. *Note: While the channel still has items `ok` will be true.*
- Alternative you can do a `range` on the channel `for value := range ch {` . In this case the for loop will stop when no more items are left on  the channel and the channel is closed. This is much prettier than the  approach above, but not always possible.

- 写入关闭的频道会导致恐慌
- 读取频道时，您可以使用 `value, ok <- ch` 。从关闭通道读取将返回所有缓冲项。一旦缓冲区项被“耗尽”，通道将返回零 `value` 并且 `ok` 将为 false。 *注意：当频道仍然有项目时，`ok` 将是真的。*
- 或者，您可以在频道 `for value := range ch {` 上做一个 `range`。在这种情况下，当通道上没有剩余项目并且通道关闭时，for 循环将停止。这比上面的方法漂亮得多，但并不总是可能的。

The points above conclude to the following:

以上几点总结如下：

- If you have a **single worker writing to the channel**, close the channel once you are done:

- 如果您有一个 **单个工作人员向频道写入数据**，请在完成后关闭频道：

```text
go func() {
    defer close(ch) // close after write is no longer possible
    for {
            select {
                case <-ctx.Done():
                    return
                ...
        ch <- value // write to the channel only happens inside the loop
    }
}()
```

- If you have **multiple workers writing to the same channel**, close the channel after waiting for all workers to finish:

- 如果您有**多个工人写入同一个频道**，请在等待所有工人完成后关闭频道：

```text
g, gCtx := errgroup.WithContext(ctx)
ch = make(...) // channel will be written from multiple workers
for w := range workers { // create n number of workers
  g.Go(func() error {
        return w.Run(ctx, ch) // workers will publish
    })
}
g.Wait() // we need to wait for all workers to stop
close(ch) // and then close the channel
```

- If you're reading from a channel, exit only when the channel has no more  data. Essentially it's the responsibility of the writer to stop the  readers, by closing the channel:

- 如果您正在从频道读取数据，请仅在频道没有更多数据时退出。本质上，作者有责任通过关闭频道来阻止读者：

```text
for v := range ch {

}

// or
for {
   select {
      case v, ok <- ch:
         if !ok { // nothing left to read
           return;
         }
                 foo(v) // process `v` normally
      case ...:
                    ...
   }
}
```

- If a worker is both reading and writing, the worker should stop when the  channel that it is reading from has no more data, and then close the  writer.

- 如果一个worker同时在读和写，worker应该在它正在读取的channel没有更多数据时停止，然后关闭writer。

## Graceful methods

## 优雅的方法

We have seen several  techniques so far for gracefully terminating a piece of long running  code. It is also useful to examine how components can expose exported  methods that can be called and then facilitate gracefully shutdown.

到目前为止，我们已经看到了几种优雅地终止一段长时间运行的代码的技术。检查组件如何公开可以调用的导出方法然后促进正常关闭也很有用。

### Blocking with ctx

### 使用 ctx 阻塞

This is the most common approach and the easier to understand and implement.

这是最常见的方法，也更容易理解和实施。

- You call a method
- You pass it a context
- The method blocks
- It returns in case of an error or when context is cancelled / timeout.

- 你调用一个方法
- 你传递给它一个上下文
- 方法块
- 它在发生错误或上下文被取消/超时时返回。

```text
// calling:
err := srv.Run(ctx, ...)

// implementation

func (srv *Service) Run(ctx context.Context, ...) {
    ...
    ...

    for {
        ...
    select {
            case <- ctx.Done()
                return ctx.Err() // Depending on our business logic,
                                                 //   we may or may not want to return a ctx error:
                                                 //   https://pkg.go.dev/context#pkg-variables
         }
    }
```

### Setup/Shutdown 

### 设置/关机

There are cases when blocking with ctx code is not the best approach. This is the case when we want greater control over when `.Shutdown` happens. This approach is a bit more complex and there is also the danger of people forgetting to call `.Shutdown`.

在某些情况下，使用 ctx 代码进行阻塞并不是最好的方法。当我们想要更好地控制 .Shutdown 何时发生时，就是这种情况。这种方法有点复杂，而且还有人们忘记调用`.Shutdown`的危险。

#### Use case

#### 用例

The code bellow demonstrates why this pattern might be useful. We want to  make sure that db Shutdown happens only after the Service is no longer  running, because the Service is depending on the database to run for it  to work.

下面的代码演示了为什么这种模式可能有用。我们要确保 db Shutdown 仅在服务不再运行后发生，因为服务依赖于运行的数据库才能工作。

By calling `db.Shutdown()` on defer, we ensure it runs after `g.Wait` returns:

通过在 defer 上调用 `db.Shutdown()`，我们确保它在 `g.Wait` 返回后运行：

```text
// calling:
func () {
   err := db.Setup() // will not block
   defer db.Shutdown()

   svc := Service{
     DB: db
   }

   g.Run(...
     svc.Run(ctx, ...)
   )
   g.Wait()
}
```

#### Implementation example

#### 实现示例

####

####

```text
type Database struct {
  ...
    cancel func()
  wait func() err
}

func (db *Database) Setup() {
    // ...
    // ...

  ctx, cancel := context.WithCancel(context.Background())
    g, gCtx := errgroup.WithContext(ctx)

  db.cancel = cancel
  db.wait = g.Wait

    for {
        ...
    select {
            case <- ctx.Done()
                return ctx.Err() // Depending on our business logic,
                                                 //   we may or may not want to return a ctx error:
                                                 //   https://pkg.go.dev/context#pkg-variables
         }
    }
}

func (db *Database) Shutdown() error {
    db.cancel()
    return db.wait()
}
```

## Final Thoughts

##  最后的想法

Terminating your long-running services gracefully is an important pattern that you  will have to implement sooner or later. This is especially true for  systems like RudderStack that act as middlewares where many connections  to external services exist and high volumes of data are handled  concurrently.

优雅地终止长期运行的服务是您迟早必须实施的重要模式。对于像 RudderStack 这样充当中间件的系统尤其如此，其中存在许多与外部服务的连接并且同时处理大量数据。

Go offers all the tools we need to implement this  pattern, and selecting the right ones depends a lot on your use case. My intention for this post was to act as a guide to help choose the right  tools for your case. If you have any questions, please reach out, and if you like solving problems like this check our [Careers page] 

Go 提供了我们实现此模式所需的所有工具，选择正确的工具在很大程度上取决于您的用例。我写这篇文章的目的是作为指导，帮助您为您的案例选择正确的工具。如果您有任何问题，请联系我们，如果您喜欢解决此类问题，请查看我们的 [Careers page]

