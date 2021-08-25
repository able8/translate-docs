# Concurrency Patterns in Go: sync.WaitGroup

# Go 中的并发模式：sync.WaitGroup

From: https://www.calhoun.io/concurrency-patterns-in-go-sync-waitgroup

来自：https://www.calhoun.io/concurrency-patterns-in-go-sync-waitgroup

When working with concurrent code, the first thing most people  notice is that the rest of their code doesn’t wait for the concurrent  code to finish before moving on. For instance, imagine that we wanted to send a message to a few services before shutting down, and we started  with the following code:

在处理并发代码时，大多数人注意到的第一件事是他们的其余代码不会等待并发代码完成后再继续。例如，假设我们想在关闭之前向一些服务发送消息，我们从以下代码开始：

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func notify(services ...string) {
	for _, service := range services {
		go func(s string) {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Finished notifying %s...\n", s)
		}(service)
	}
	fmt.Println("All services notified!")
}

func main() {
	notify("Service-1", "Service-2", "Service-3")
	// Running this outputs "All services notified!" but we
	// won't see any of the services outputting their finished messages!
}
 ```

 
If we were to run this code with some sleeps in place to similuate  latency, we would see an “All services notified!” message output, but  none of the “Finished notifying …” messages would ever print out,  suggesting that our applicaiton doesn’t wait for these messages to send  before shutting down. That is going to be an issue!

如果我们在适当的睡眠状态下运行此代码以模拟延迟，我们将看到“所有服务已通知！”消息输出，但“已完成通知……”消息都不会打印出来，这表明我们的应用程序在关闭之前不会等待这些消息发送。这将是一个问题！

One way to solve this problem is to use a [sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup). This is a type provided by the standard library that makes it easy to  say, “I have N tasks that I need to run concurrently, wait for them to  complete, and then resume my code.”

解决此问题的一种方法是使用 [sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup)。这是标准库提供的一种类型，可以很容易地说，“我有 N 个任务需要并发运行，等待它们完成，然后恢复我的代码。”

To use a `sync.WaitGroup` we do roughly four things:

要使用`sync.WaitGroup`，我们大致做四件事：

1. Declare the `sync.WaitGroup`
2. Add to the WaitGroup queue
3. Tell our code to wait on the WaitGroup queue to reach zero before proceeding
4. Inside each goroutine, mark items in the queue as done

1. 声明`sync.WaitGroup`
2.添加到WaitGroup队列
3. 告诉我们的代码在继续之前等待 WaitGroup 队列达到零
4. 在每个 goroutine 中，将队列中的项目标记为完成

The code below shows this, and we will discuss the code after you give it a read.

下面的代码显示了这一点，我们将在您阅读后讨论该代码。

```go
func notify(services ...string) {
	var wg sync.WaitGroup

	for _, service := range services {
		wg.Add(1)
		go func(s string) {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Finished notifying %s...\n", s)
			wg.Done()
		}(service)
	}

	wg.Wait()
	fmt.Println("All services notified!")
}
 ```

 
At the very start of our code we achieve (1) by declaring the  sync.WaitGroup. We do this before calling any goroutines so it is  available for each goroutine.

在我们代码的最开始，我们通过声明 sync.WaitGroup 来实现 (1)。我们在调用任何 goroutine 之前执行此操作，以便每个 goroutine 都可以使用它。

Next we need to add items to the WaitGroup queue. We do this by calling [Add(n)](https://golang.org/pkg/sync/#WaitGroup.Add), where `n` is the number of items we want to add to the queue. This means we can call `Add(5)` once if we know we want to wait on five tasks, or in this case we opt to call `Add(1)` for each iteration of the loop. Both approaches work fine, and the code above could easily be replaced with something like:

接下来我们需要将项目添加到 WaitGroup 队列。我们通过调用 [Add(n)](https://golang.org/pkg/sync/#WaitGroup.Add) 来做到这一点，其中 `n` 是我们想要添加到队列中的项目数。这意味着如果我们知道要等待五个任务，我们可以调用一次 Add(5) ，或者在这种情况下我们选择为循环的每次迭代调用 Add(1) 。两种方法都可以正常工作，上面的代码可以很容易地替换为：

```go
wg.Add(len(services))
for _, service := range services {
  go func(s string) {
    fmt.Printf("Starting to notifing %s...\n", s)
    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
    fmt.Printf("Finished notifying %s...\n", s)
    wg.Done()
  }(service)
}
```

 
In either case, **I recommend calling Add() outside of the concurrent code to ensure it runs immediately**. If we were to instead place this inside of the goroutine it is possible that the program will get to the wg.Wait() line before the goroutine  has run, in which case wg.Wait() won't have anything to wait on and we  will be in the same position we were in before. This is shown on the Go  Playground here: https://play.golang.org/p/Yl4f_5We6s7

无论哪种情况，**我建议在并发代码之外调用 Add() 以确保它立即运行**。如果我们将它放在 goroutine 中，那么程序可能会在 goroutine 运行之前到达 wg.Wait() 行，在这种情况下 wg.Wait() 将没有任何东西可以等待并且我们将处于与以前相同的位置。这显示在 Go Playground 上：https://play.golang.org/p/Yl4f_5We6s7

We need to also mark items in our WaitGroup queue as complete. To do this we call [Done()](https://golang.org/pkg/sync/#WaitGroup.Done), which unlike `Add()`, does not take an argument and needs to be called for as many items are  in the WaitGroup queue. Because this is dependent on the code running in a goroutine, the call to `Done()` should be run inside of the goroutine we want to wait on. If we were to instead call `Done()` inside the for loop but NOT inside the goroutine it would mark every  task as complete before the goroutine actually runs. This is shown on  the Go Playground here: https://play.golang.org/p/i2vu2vGjYgB

我们还需要将 WaitGroup 队列中的项目标记为完成。为此，我们调用 [Done()](https://golang.org/pkg/sync/#WaitGroup.Done)，它与 ​​`Add()` 不同，它不接受参数并且需要调用尽可能多的参数项目在 WaitGroup 队列中。因为这取决于在 goroutine 中运行的代码，所以对 `Done()` 的调用应该在我们想要等待的 goroutine 内运行。如果我们在 for 循环中而不是在 goroutine 中调用 `Done()`，它会在 goroutine 实际运行之前将每个任务标记为完成。这显示在 Go Playground 上：https://play.golang.org/p/i2vu2vGjYgB

Finally, we need to wait on all the items queued up in the WaitGroup to finish. We do this by calling [Wait()](https://golang.org/pkg/sync/#WaitGroup.Wait), which causes our program to wait there until the WaitGroup’s queue is cleared. 
最后，我们需要等待所有在 WaitGroup 中排队的项目完成。我们通过调用 [Wait()](https://golang.org/pkg/sync/#WaitGroup.Wait) 来做到这一点，这会导致我们的程序在那里等待，直到 WaitGroup 的队列被清除。
It is worth noting that this pattern works best when we don’t care  about gathering any results from the goroutines. If we find ourselves in a situation where we need data returned from each goroutine, it will  likely be easier to use channels to communicate that information. For  instance, the following code is very similar to the wait group example,  but it uses a channel to receive a message from each goroutine after it  completes.

值得注意的是，当我们不关心从 goroutines 收集任何结果时，这种模式最有效。如果我们发现自己需要从每个 goroutine 返回数据，那么使用通道来传达这些信息可能会更容易。例如，以下代码与等待组示例非常相似，但它使用一个通道在每个 goroutine 完成后接收来自每个 goroutine 的消息。

```go
func notify(services ...string) {
	res := make(chan string)
	count := 0

	for _, service := range services {
		count++
		go func(s string) {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			res <- fmt.Sprintf("Finished %s", s)
		}(service)
	}

	for i := 0; i < count; i++ {
		fmt.Println(<-res)
	}

	fmt.Println("All services notified!")
}
 ```

 
*Why don’t we just use channels all the time?* Based on this last example, we can do everything a sync.WaitGroup does with a channel, so why the new type?

*为什么我们不一直使用通道？* 基于最后一个例子，我们可以做一个 sync.WaitGroup 对一个通道所做的一切，那么为什么要使用新类型呢？

The short answer is that `sync.WaitGroup` is a bit clearer when we don’t care about the data being returned from the goroutine. It signifies to other developers that we just want to wait for a group of  goroutines to complete, whereas a channel communicates that we are  interested in results from those goroutines.

简短的回答是，当我们不关心从 goroutine 返回的数据时，`sync.WaitGroup` 会更清晰一些。它向其他开发人员表示我们只想等待一组 goroutines 完成，而 channel 则表示我们对这些 goroutines 的结果感兴趣。

That’s it for this one. In future articles I’ll present a few more  concurrency patterns, and we can continue to expand upon the subject. If you happen to have a specific use case and want to share it (or request I write about it), just [send me an email](mailto:jon@calhoun.io).

就是这个。在以后的文章中，我将介绍更多并发模式，我们可以继续扩展这个主题。如果你碰巧有一个特定的用例并想分享它（或要求我写它），只需[给我发一封电子邮件](mailto:jon@calhoun.io)。

This article is part of the series, [Concurrency Patterns in Go](https://www.calhoun.io/concurrency-patterns-in-go/). 
本文是 [Go 中的并发模式](https://www.calhoun.io/concurrency-patterns-in-go/) 系列的一部分。
