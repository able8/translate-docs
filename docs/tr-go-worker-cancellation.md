# Go - graceful shutdown of worker goroutines

# Go - 优雅地关闭 worker goroutines

### 05 October 2019

In this blog post we’ll take a look at gracefully shutting down a Go program having worker goroutines performing tasks that must be completed before allowing the program to shut down.

在这篇博文中，我们将看看如何优雅地关闭 Go 程序，让工作程序 goroutine 执行必须在允许程序关闭之前完成的任务。

# Introduction

#  介绍

In a recent project we had a use-case where a Go-based microservice was consuming events emitted from a 3rd party library. These events would undergo a bit of processing before resulting in a call to an external service. That external service handles each request quite slowly, but can on the other hand handle many concurrent requests. Therefore, we implemented a simple internal worker-pool to fanout incoming events onto several concurrently executing goroutines.

在最近的一个项目中，我们有一个用例，其中基于 Go 的微服务正在使用从 3rd 方库发出的事件。这些事件在调用外部服务之前会进行一些处理。该外部服务处理每个请求的速度非常慢，但另一方面可以处理许多并发请求。因此，我们实现了一个简单的内部工作池来将传入事件扇出到多个并发执行的 goroutine 上。

On a general level, it looks like this:

在一般层面上，它看起来像这样：

![figure 1](http://callistaenterprise.se/assets/blogg/goblog/other/blog-oct2019-1.png)

However, we needed to make sure that if the microservice is shutting down, any currently running requests to the external service must be allowed to finish with the outcome persisted to our internal backend.

但是，我们需要确保如果微服务关闭，则必须允许任何当前正在运行的外部服务请求完成，并将结果持久化到我们的内部后端。

## Worker pools and Sigterm handling

## 工作池和 Sigterm 处理

The [worker pool pattern](https://gobyexample.com/worker-pools) is a well known Go pattern for worker-pools and there are numerous examples on how to do [SIGTERM notification based graceful shutdown](https://gobyexample.com/signals), but we realized that a number of our requirements made our use-case somewhat more complicated.

[worker pool pattern](https://gobyexample.com/worker-pools) 是一个众所周知的工作池 Go 模式，并且有很多关于如何执行 [基于 SIGTERM 通知的优雅关闭](https://gobyexample.com/worker-pools)的示例gobyexample.com/signals)，但我们意识到我们的一些要求使我们的用例变得更加复杂。

When the program receives a SIGTERM / SIGINT signal - for example, from our container orchestrator scaling down the number of replicas - any currently executing worker goroutines must be allowed to finish their long-running work before terminating the program.

当程序收到 SIGTERM / SIGINT 信号时——例如，来自我们的容器编排器缩小副本数量——必须允许任何当前正在执行的工作协程在终止程序之前完成其长时间运行的工作。

To make things slightly more complicated, we had no control over the producer-side library. We get to register a callback function which is invoked each time the library has a new event for us. The library blocks until the callback function has finished executing and then invokes it again if there’s more events.

为了让事情稍微复杂一点，我们无法控制生产者端库。我们可以注册一个回调函数，每次库有新事件时都会调用该函数。库会阻塞直到回调函数执行完毕，然后如果有更多事件再次调用它。

The worker-pool goroutines are fed with events to process using the standard “range over channel” construct, e.g:

工作池 goroutine 被提供事件以使用标准的“通道范围”构造进行处理，例如：

```go
func workerFunc() {
    for event := range jobsChan { // blocks until an event is received or channel is closed.
        // handle the event...
    }
}
```

Which means that the cleanest way to let a worker “finish” is to close that “jobsChan” channel.

这意味着让工人“完成”的最干净的方法是关闭“jobsChan”频道。

## Closing on the producer side

## 在生产者端关闭

One of the first things you first learn about closing channels in Go is that the program will panic if a send occurs on a closed channel. This boils down to a very simple rule:

关于在 Go 中关闭通道，您首先了解的第一件事是，如果在关闭的通道上发生发送，程序将发生恐慌。这归结为一个非常简单的规则：

```
"Always close a channel on the producer side"
```

What’s the producer side anyway? Well, typically the _goroutine_ that is putting events on that channel:

制片方到底是什么？好吧，通常是在该通道上放置事件的 _goroutine_：

```go
func callbackFunc(event int) {
    jobsChan<-event
}
```

This is our callbackFunc that we’ve registered with the external library that passes events to us. _(To keep these examples simple, I’ve replaced the real events with a simple int as payload.)_

这是我们的 callbackFunc，我们已经向外部库注册了它，该库将事件传递给我们。 _（为了让这些例子保持简单，我用一个简单的 int 作为有效载荷替换了真实事件。）_

How do you _safely_ protect the piece of code above from sending on a closed channel? It’s not trivial to go down the route of Mutexes and boolean flags and if-statements to determine if some _other_ goroutine has closed the channel and to control whether a send should be allowed or not. Very open to potential race conditions and non-deterministic behaviour.

你如何_安全地_保护上面的一段代码不被发送到一个封闭的频道？沿着 Mutexes 和布尔标志以及 if 语句的路线确定某个 _other_ goroutine 是否关闭了通道并控制是否应该允许发送并非易事。对潜在的竞争条件和非确定性行为非常开放。

Our solution was to introduce an intermediate channel and an internal “consumer” that acts as a proxy between the callback and the jobs channel:

我们的解决方案是引入一个中间通道和一个内部“消费者”，作为回调和作业通道之间的代理：

![figure 2](http://callistaenterprise.se/assets/blogg/goblog/other/blog-oct2019-2.png)

The consumer function looks like this:

消费者函数如下所示：

```go
func startConsumer(ctx context.Context) {
    // Loop until a ctx.Done() is received.Note that select{} blocks until either case happens
    for {
            select {
            case event := <-intermediateChan:
                jobsChan <- event
            case _ <- ctx.Done():
                close(jobsChan)
                return             // exit this function so we don't consume anything more from the intermediate chan
            }
    }
}

```

Ok, wait a minute. What’s this “select” and “ctx.Done()”? 

好的，等一下。这个“选择”和“ctx.Done()”是什么？

The [select](https://gobyexample.com/select) statement is IMHO one of the absolutely most awesome things about Go. It allows waiting and coordinating on multiple channels. In this particular case, we will either receive an event on the intermediate channel to pass to the jobsChan or a cancellation signal from the [context.Context](https://golang.org/pkg/context/#WithCancel).

[select](https://gobyexample.com/select) 语句是恕我直言，Go 中绝对最棒的事情之一。它允许在多个通道上等待和协调。在这种特殊情况下，我们将在中间通道上接收一个事件以传递给 jobsChan，或者接收来自 [context.Context](https://golang.org/pkg/context/#WithCancel) 的取消信号。

The _return_ statement after closing the jobsChan will step us out from the for-loop and function which makes sure _no new events can be passed_ to the jobsChan and _no events_ will be consumed from the intermediateChan.

关闭jobsChan 后的_return_ 语句将使我们退出for 循环和函数，这确保_没有新事件可以传递_ 到jobsChan，并且_no events_ 将从intermediateChan 消耗。

So either an event is passed to the jobs (which the workers consume from) or the jobsChan is closed _in the same goroutine_ as the producer.

因此，要么将事件传递给作业（工作人员从中消费），要么将 jobsChan _在与生产者相同的 goroutine_ 中关闭。

Closing the jobsChan means all workers will stop ranging on the jobsChan on the consumer side:

关闭jobsChan 意味着所有worker 将停止在consumer 端的jobsChan 上运行：

```go
for event := range jobsChan { // <- on the close(jobsChan), all goroutines waiting for jobs here will exit the for-loop
    // handle the event...
}

```

## Issuing the cancellation signal

## 发出取消信号

Waiting for a Go program to terminate is a well-known pattern:

等待 Go 程序终止是一个众所周知的模式：

```go
func main() {
    ... rest of program ...

    termChan := make(chan os.Signal)
    signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
    <-termChan // Blocks here until either SIGINT or SIGTERM is received.
    // what now?
}

```

At “what now?” the main goroutine resumes execution after capturing SIGINT or SIGTERM. We need to tell the consumer goroutine that passes events from the intermediateChan to the jobsChan to close the jobsChan across goroutine boundaries.

在“现在呢？”主协程在捕获 SIGINT 或 SIGTERM 后恢复执行。我们需要告诉消费者 goroutine 将事件从 middleChan 传递到jobsChan 以关闭跨goroutine 边界的jobsChan。

Again, it is technically possible - but rather awkward and error prone - to solve this using Mutexes and conditional statements. Instead, we’ll utilize the cancellation support of context.Context we touched on earlier.

同样，使用互斥锁和条件语句解决这个问题在技术上是可行的——但相当笨拙且容易出错。相反，我们将利用我们之前提到的 context.Context 的取消支持。

Somewhere in our _func main()_ we set up a root background context with cancellation support:

在我们 _func main()_ 的某个地方，我们设置了一个支持取消的根背景上下文：

```go
func main() {
    ctx, cancelFunc := context.WithCancel(ctx.Background())
    // ... some omitted code ...

    go startConsumer(ctx) // pass the cancellable context to the consumer function

    // ... some more omitted code ...
    <-termChan

    cancelFunc() // call the cancelfunc to notify the consumer it's time to shut stuff down.
}

```

This is how the _<-ctx.Done()_ select case becomes invoked, starting the graceful teardown of channels and workers.

这就是 _<-ctx.Done()_ select case 被调用的方式，开始优雅地拆除通道和工作线程。

## Using WaitGroups

## 使用等待组

There's only one problem with the solution above - the program will exit immediately after the _cancelFunc()_ invocation, which means our worker goroutines having in-flight invocations won't have time to finish, potentially leaving transactions in our system in an indeterminate state .

上面的解决方案只有一个问题 - 程序将在 _cancelFunc()_ 调用后立即退出，这意味着我们正在进行调用的工作程序 goroutine 将没有时间完成，这可能会使我们系统中的事务处于不确定状态.

We need to halt shutdown until all workers report that they’re done with whatever stuff they were doing. Enter [sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup) that lets us wait for an arbitrary number of goroutines to finish!

我们需要停止关闭，直到所有工人报告他们已经完成了他们正在做的任何事情。输入 [sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup) 让我们等待任意数量的 goroutine 完成！

When starting our workers, we pass along a pointer to a WaitGroup created in _func main()_:

在启动我们的工作进程时，我们传递一个指向在 _func main()_ 中创建的 WaitGroup 的指针：

```go
const numberOfWorkers = 4

func main() {
    // ... omitted ...
    wg := &sync.WaitGroup{}
    wg.Add(numberOfWorkers)

    // Start [workerPoolSize] workers
    for i := 0;i < workerPoolSize;i++ {
        go workerFunc(wg)
    }

    // ... more omitted stuff ...

    <-termChan    // wait for SIGINT / SIGTERM
    cancelFunc()  // send the shutdown signal through the context.Context
    wg.Wait()     // program will wait here until all worker goroutines have reported that they're done
    fmt.Println("Workers done, shutting down!")
}

```

This changes our worker startup function a little:

这稍微改变了我们的工作启动函数：

```go
func workerFunc(wg *sync.WaitGroup) {
    defer wg.Done() // Mark this goroutine as done!once the function exits
    for event := range jobsChan {
        // handle the event...
    }
}

```

The _wg.Done()_ decrements the waitgroup by one and once the internal counter reaches 0, the main goroutine will continue past the _wg.Wait()_. Graceful shutdown complete!

_wg.Done()_ 将等待组减 1，一旦内部计数器达到 0，主 goroutine 将继续通过 _wg.Wait()_。优雅关机完成！

### Running

###  跑步

The source for the final program comes in the next section. In it, I’ve added some logging statements so we can follow what happens.

最终程序的源代码在下一节。在其中，我添加了一些日志记录语句，以便我们可以跟踪发生的情况。

Here’s the output for an execution of the program with 4 worker goroutines, where I use Ctrl+C to stop the program:

这是使用 4 个工作 goroutine 执行程序的输出，我使用 Ctrl+C 停止程序：

```
$ go run main.go
Worker 3 starting
Worker 2 starting
Worker 1 starting
Worker 0 starting
Worker 3 finished processing job 0
Worker 0 finished processing job 3
^C*********************************     <-- HERE I PRESS CTRL+C
Shutdown signal received
*********************************
Worker 3 finished processing job 4
Worker 2 finished processing job 1
Worker 1 finished processing job 2
Consumer received cancellation signal, closing jobsChan!<-- Here, the consumer receives the <-ctx.Done()
Worker 3 finished processing job 6
Worker 0 finished processing job 5
Worker 1 finished processing job 8
Worker 2 finished processing job 7
Worker 0 finished processing job 10
Worker 0 interrupted                    <-- Worker 0 has finished job #10, 3 left
Worker 2 finished processing job 12
Worker 2 interrupted                    <-- Worker 2 has finished job #12, 2 left
Worker 3 finished processing job 9
Worker 3 interrupted                    <-- Worker 3 has finished job #9, 1 left
Worker 1 finished processing job 11
Worker 1 interrupted                    <-- Worker 1 has finished job #11, all done
All workers done, shutting down!

```

As one may observe that the point in time where the consumer receives _<-ctx.Done()_ is actually non-deterministic due to how the Go runtime schedules communications on channels into the select statement. The Go specification says:

正如人们可能观察到的那样，消费者接收 _<-ctx.Done()_ 的时间点实际上是不确定的，因为 Go 运行时如何将通道上的通信调度到 select 语句中。 Go 规范说：

```
"If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection."
```

This is why jobs may be passed to the workers even after CTRL+C is pressed.

这就是为什么即使在按下 CTRL+C 后作业也可能传递给工作人员的原因。

Another peculiar thing is that it seems as jobs (jobs 9-12) are passed to the workers even _after_ the jobsChan has been closed. Well - they were actually passed _before_ the channel was closed. The reason for this is our use of a buffered channel with 4 “slots”. This means that if all four workers have consumed a job from the channel and are processing them, there are potentially four new jobs waiting to be consumed by workers on the channel given that our 3rd party producer is constantly passing new events to us at a higher rate than our workers can manage. Closing a channel doesn’t affect data already buffered into the channel - Go allows consumers to consume those.

另一个奇怪的事情是，即使在jobsChan 关闭之后，工作（工作9-12）似乎也会传递给工作人员。好吧 - 他们实际上是在 _before_ 通道关闭之前通过的。这样做的原因是我们使用了一个带有 4 个“插槽”的缓冲通道。这意味着，如果所有四个工作人员都从通道中消费了一个工作并正在处理它们，那么鉴于我们的 3rd 方生产者不断地以更高的速度将新事件传递给我们，可能有四个新工作等待通道上的工作人员消费。速度超出了我们的工人可以管理的范围。关闭通道不会影响已经缓冲到通道中的数据 - Go 允许消费者使用这些数据。

If we change the jobsChan to be unbuffered:

如果我们将 jobsChan 更改为无缓冲：

```
jobsChan := make(chan int)

```

And run again:

并再次运行：

```
$ go run main.go
.... omitted for brevity ....
^C*********************************
Shutdown signal received
*********************************
Worker 3 finished processing job 3
Worker 3 started job 5
Worker 0 finished processing job 4
Worker 0 started job 6
Consumer received cancellation signal, closing jobsChan!<-- again, it may take some time until the consumer is handed <-ctx.Done()
Consumer closed jobsChan
Worker 1 finished processing job 1     <-- From here on, we see that each worker finishes exactly one job before being interrupted.
Worker 1 interrupted
Worker 2 finished processing job 2
Worker 2 interrupted
Worker 0 finished processing job 6
Worker 0 interrupted
Worker 3 finished processing job 5
Worker 3 interrupted
All workers done, shutting down!

```

This time we don’t see any “unexpected” jobs being consumed by the workers after the channel was closed. Having a channel buffer the same size as the number of workers is a however a common optimization to keep workers fed with data without unnecessary stalling on the producer side.

这次我们没有看到在通道关闭后工作人员消耗了任何“意外”的工作。然而，拥有与工作线程数量相同大小的通道缓冲区是一种常见的优化，可以让工作线程接收数据，而不会在生产者端造成不必要的停顿。

# The full program

# 完整程序

The snippets above are somewhat simplified to keep them as concise as possible. The full program with some structs for encapsulation and simulation of the 3rd-party producer follows here:

上面的代码片段经过了一定程度的简化，以使其尽可能简洁。带有一些用于封装和模拟第 3 方生产者的结构的完整程序如下：

```go
package main

import (
    "context"
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

const workerPoolSize = 4

func main() {
    // create the consumer
    consumer := Consumer{
        ingestChan: make(chan int, 1),
        jobsChan:   make(chan int, workerPoolSize),
    }

    // Simulate external lib sending us 10 events per second
    producer := Producer{callbackFunc: consumer.callbackFunc}
    go producer.start()

    // Set up cancellation context and waitgroup
    ctx, cancelFunc := context.WithCancel(context.Background())
    wg := &sync.WaitGroup{}

    // Start consumer with cancellation context passed
    go consumer.startConsumer(ctx)

    // Start workers and Add [workerPoolSize] to WaitGroup
    wg.Add(workerPoolSize)
    for i := 0;i < workerPoolSize;i++ {
        go consumer.workerFunc(wg, i)
    }

    // Handle sigterm and await termChan signal
    termChan := make(chan os.Signal)
    signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

    <-termChan         // Blocks here until interrupted

    // Handle shutdown
    fmt.Println("*********************************\nShutdown signal received\n*********************************")
    cancelFunc()       // Signal cancellation to context.Context
    wg.Wait()          // Block here until are workers are done

    fmt.Println("All workers done, shutting down!")
}

```

The consumer struct:

消费者结构：

```go
// -- Consumer below here!
type Consumer struct {
    ingestChan chan int
    jobsChan   chan int
}

// callbackFunc is invoked each time the external lib passes an event to us.
func (c Consumer) callbackFunc(event int) {
    c.ingestChan <- event
}

// workerFunc starts a single worker function that will range on the jobsChan until that channel closes.
func (c Consumer) workerFunc(wg *sync.WaitGroup, index int) {
    defer wg.Done()

    fmt.Printf("Worker %d starting\n", index)
    for eventIndex := range c.jobsChan {
        // simulate work  taking between 1-3 seconds
        fmt.Printf("Worker %d started job %d\n", index, eventIndex)
        time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
        fmt.Printf("Worker %d finished processing job %d\n", index, eventIndex)
    }
    fmt.Printf("Worker %d interrupted\n", index)
}

// startConsumer acts as the proxy between the ingestChan and jobsChan, with a select to support graceful shutdown.
func (c Consumer) startConsumer(ctx context.Context) {
    for {
        select {
        case job := <-c.ingestChan:
            c.jobsChan <- job
        case <-ctx.Done():
            fmt.Println("Consumer received cancellation signal, closing jobsChan!")
            close(c.jobsChan)
            fmt.Println("Consumer closed jobsChan")
            return
        }
    }
}

```

Finally, the Producer struct that simulates our external library:

最后，模拟我们外部库的 Producer 结构：

```go
// -- Producer simulates an external library that invokes the
// registered callback when it has new data for us once per 100ms.
type Producer struct {
    callbackFunc func(event int)
}
func (p Producer) start() {
    eventIndex := 0
    for {
        p.callbackFunc(eventIndex)
        eventIndex++
        time.Sleep(time.Millisecond * 100)
    }
}

```

# Summary

#  概括

I hope this little blog post have given a simple example of goroutine-based worker pools and how to gracefully shut them down using context-based cancellation, WaitGroups and producer-side closing of channels.

我希望这篇小博文给出了一个简单的基于 goroutine 的工作池示例，以及如何使用基于上下文的取消、WaitGroups 和生产者端关闭通道来优雅地关闭它们。

