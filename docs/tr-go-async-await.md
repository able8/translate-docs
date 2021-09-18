# If at first you don’t succeed  call it version 1.0

# 如果一开始你没有成功，就叫它 1.0 版

The await/async concurrency pattern in Golang Subscribe to [my newsletter](https://tinyletter.com/made2591) to be informed about my new blog posts, talks and activities.

Golang 中的 await/async 并发模式订阅 [我的时事通讯](https://tinyletter.com/made2591) 以了解我的新博文、演讲和活动。

2020-01-02

2020-01-02

### Introduction

###  介绍

First of all...happy new year! I decided after a while to come back online speaking about Golang. In this post, I will focus on **parallelism** and **concurrency** and how you can achieve the same behavioral pattern you can achieve with Node.js using `await/async` statements, without the difficulties (hopefully) of dealing with [Single Threaded Event Loop](https://en.wikipedia.org/wiki/Node.js) and these primitives (that, btw, keep things really simple). Let's start!

首先……新年快乐！过了一会儿，我决定回到网上谈论 Golang。在这篇文章中，我将重点讨论 **并行** 和 **并发** 以及如何使用 `await/async` 语句实现与 Node.js 相同的行为模式，而不会有（希望如此）处理 [单线程事件循环](https://en.wikipedia.org/wiki/Node.js)和这些原语（顺便说一句，让事情变得非常简单)。开始吧！

#### A bit of confusion

#### 有点混乱

Concurrency and parallelism are two terms that are bound to come  across often when looking into multitasking and are often used  interchangeably. However, they mean two distinctly different things. Concurrency is all about the following:

并发性和并行性是在研究多任务时经常会遇到的两个术语，并且经常互换使用。然而，它们意味着两个截然不同的事物。并发是关于以下所有内容：

> Dealing with many things at once.

> 同时处理很多事情。

This means that we manage to get multiple things done at once in a  given period of time. However, we will only be doing a single thing at a time. This tends to happen in programs where one task is waiting and  the program decides to run another task in the idle time. Pretty simple.

这意味着我们设法在给定的时间内同时完成多项工作。但是，我们一次只会做一件事。这往往发生在一个任务正在等待并且程序决定在空闲时间运行另一个任务的程序中。很简单。

On the other hand, parallelism means

另一方面，并行性意味着

> Doing many things at once.

> 同时做很多事情。

This means that even if we have two tasks, they are *continuously working without any breaks in between them*. The two tasks run independently and are not influenced by each other in any manner.

这意味着即使我们有两个任务，它们也在*连续工作，它们之间没有任何中断*。这两个任务独立运行，互不影响。

We can also say that concurrency is the composition of independently  executing processes, while parallelism is the simultaneous execution of  (possibly related) computations[1](https://madeddu.xyz/posts/go-async-await/#fn:golang-blog). However, the components running in parallel, even inside a single  application, have might have to communicate with each other. These  communications happen between the components of even the simplest  applications, and the overhead is generally low in concurrent systems. In the case when components run in parallel in multiple cores, this  communication overhead could be (and generally is) higher. Hence  parallel programs do not always result in faster execution times.

我们也可以说并发是独立执行进程的组合，而并行是同时执行（可能相关的）计算[1](https://madeddu.xyz/posts/go-async-await/#fn:golang-博客）。然而，并行运行的组件，即使在单个应用程序中，也可能必须相互通信。这些通信发生在即使是最简单的应用程序的组件之间，并且在并发系统中开销通常很低。在组件在多个内核中并行运行的情况下，这种通信开销可能（并且通常)更高。因此，并行程序并不总是导致更快的执行时间。

Concurrency is an inherent part of the Go programming language, and it's handled using `goroutines` and `channels`.

并发是 Go 编程语言的一个固有部分，它使用 `goroutines` 和 `channels` 来处理。

#### goroutines

#### 协程

If you look at the golang-by-example tour[2](https://madeddu.xyz/posts/go-async-await/#fn:golang-by-example-goroutines), the definition of a goroutine is as simple as the following:

如果你查看 golang-by-example tour[2](https://madeddu.xyz/posts/go-async-await/#fn:golang-by-example-goroutines)，goroutine 的定义是简单如下：

> A goroutine is a lightweight thread managed by the Go runtime.

> goroutine 是由 Go 运行时管理的轻量级线程。

The Golang official site states[3](https://madeddu.xyz/posts/go-async-await/#fn:golang-goroutines) that they're called *goroutines* because the existing terms—threads, coroutines, processes , and so  on—convey inaccurate connotations. A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same  address space. It is lightweight, costing little more than the  allocation of stack space. And the stacks start small, so they are  cheap, and grow by allocating (and freeing) heap storage as required.

Golang 官方网站指出 [3](https://madeddu.xyz/posts/go-async-await/#fn:golang-goroutines) 它们被称为 *goroutines* 因为现有的术语——线程、协程、进程，等等——传达不准确的内涵。 goroutine 有一个简单的模型：它是一个与同一地址空间中的其他 goroutine 并发执行的函数。它是轻量级的，成本比分配堆栈空间多一点。并且堆栈开始时很小，因此它们很便宜，并且通过根据需要分配（和释放)堆存储来增长。

Moreover, the goroutines are multiplexed to a fewer number of OS  threads, thus there might be only one thread in a program with thousands of goroutines. If one goroutine should block, such as while waiting for I/O, then another OS thread is created and the remaining goroutines are moved to the new OS thread and continue to run. Their design hides many of the complexities of thread creation and management.

此外，goroutine 被多路复用到较少数量的 OS 线程，因此在具有数千个 goroutine 的程序中可能只有一个线程。如果一个 goroutine 应该阻塞，例如在等待 I/O 时，则会创建另一个 OS 线程，并将剩余的 goroutine 移动到新的 OS 线程并继续运行。它们的设计隐藏了线程创建和管理的许多复杂性。

The cost of creating a Goroutine is tiny when compared to a thread. Hence it's common for Go applications to have thousands of goroutines  running concurrently.

与线程相比，创建 Goroutine 的成本很小。因此，Go 应用程序通常会同时运行数千个 goroutine。

Prefix a function or method call with the go keyword to run the call  in a new goroutine. When the call completes, the goroutine exits,  silently. (The effect is similar to the Unix shell's & notation for  running a command in the background.). Let's make an example.

使用 go 关键字为函数或方法调用添加前缀，以在新的 goroutine 中运行调用。当调用完成时，goroutine 静默退出。 （效果类似于在后台运行命令的 Unix shell 的 & 符号。）。让我们举个例子。

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0;i < 5;i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

You can easily test the pseudo-random behavior of printed output: the first time say is called doesn't block the execution of the main  function, thus the *hello* string appears interleaved by the world string (5 times each). Let's go ahead introducing the concept of channels.

您可以轻松测试打印输出的伪随机行为：第一次调用 say 不会阻止主函数的执行，因此 *hello* 字符串与世界字符串交错出现（每个 5 次）。让我们继续介绍渠道的概念。

#### channels

#### 频道

Channels are the pipes that connect concurrent goroutines. You can  send values into channels from one goroutine and receive those values  into another goroutine, thus make many goroutines communicate between  each other - actually, orchestrate them - using channels (aka...  memory). In fact, it's pretty known that Go's approach to concurrency  differs from the traditional use of (not only) threads, but shared  memory as well. Philosophically, the idea behind Go can be summarized by the following sentence:

通道是连接并发 goroutine 的管道。您可以将值从一个 goroutine 发送到通道中，然后将这些值接收到另一个 goroutine 中，从而使许多 goroutine 相互通信 - 实际上，编排它们 - 使用通道（又名......内存）。事实上，众所周知，Go 的并发方法不同于（不仅）线程的传统使用，还不同于共享内存。从哲学上讲，Go 背后的思想可以用下面这句话来概括：

> Don't communicate by sharing memory; share memory by communicating.

> 不要通过共享内存来交流；通过通信共享内存。

Channels can be thought of as a pipe using which goroutines  communicate. They allow you to pass references to data structures  between goroutines, so if you consider this as passing around ownership  of the data (the ability to read and write it), they become a powerful  and expressive synchronization mechanism. Moreover, channels by design  prevent race conditions from happening when accessing shared memory  using goroutines.

通道可以被认为是一个管道，使用 goroutine 进行通信。它们允许您在 goroutine 之间传递对数据结构的引用，因此如果您认为这是传递数据的所有权（读取和写入数据的能力），它们将成为一种强大且富有表现力的同步机制。此外，通道的设计可以防止在使用 goroutine 访问共享内存时发生竞争条件。

![img](https://i.imgur.com/KFAKaeX.jpg)

Let's make an example.

让我们举个例子。

```go
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
}
```

Another way to think about the channels is as *typed conduits through which you can send and receive values* with the channel operator, `<-`. In the code above, `c <- sum` send `v` to the channel `sum` and `x, y := <-c, <-c` receive two times `v` from the channel `c` and assign the respective values to `x` and `y`. You can play more with channels in the Golang-tour[4](https://madeddu.xyz/posts/go-async-await/#fn:golang-by-example-channels) and follow the codewalk[5]( https://madeddu.xyz/posts/go-async-await/#fn:golang-channels-codewalk) after that to get confidence with the use of them.

考虑通道的另一种方式是*类型化的管道，您可以通过它使用通道运算符“<-”发送和接收值*。在上面的代码中，`c <- sum` 将`v` 发送到通道`sum` 和`x, y := <-c, <-c` 从通道`c` 接收两次`v` 并赋值`x` 和 `y` 的相应值。您可以在 Golang-tour[4](https://madeddu.xyz/posts/go-async-await/#fn:golang-by-example-channels) 中更多地使用频道并按照代码步 [5]( https://madeddu.xyz/posts/go-async-await/#fn:golang-channels-codewalk) 之后，才能对它们的使用充满信心。

Done? Then we should be ready and pretty confident to map this concept to the well known `async/await` pattern. Let's go ahead!

完毕？然后我们应该准备好并且非常有信心将这个概念映射到众所周知的 `async/await` 模式。让我们继续！

### async/await

### 异步/等待

Since I'm not a Javascript expert (neither a Golang one, I'm sorry about that), you definitely know more than me about `Promise` and `async/await`. The simplest use case you can think about is the following:

由于我不是 Javascript 专家（也不是 Golang 专家，对此我很抱歉），你肯定比我更了解 `Promise` 和 `async/await`。您可以想到的最简单的用例如下：

```js
const sleep = require('util').promisify(setTimeout)
async function myAsyncFunction() {
    await sleep(2000)
    return 2
};

(async function() {
    const result = await myAsyncFunction();
    // outputs `2` after two seconds
    console.log(result);
})();
```

What this code does should be simple to understand: it simply  simulates a workload of 2 seconds and asynchronously waits for it to be  completed. Also, since the run of a script from the shell is  synchronous, you have to *await* for the execution of `myAsyncFunction` from inside an async context, otherwise the `Node.js` runtime will complaint. You should be able to copy and paste the code inside a `test.js` file and run it from the console with `node test.js`.

这段代码所做的应该很容易理解：它只是模拟 2 秒的工作负载并异步等待它完成。此外，由于 shell 中的脚本运行是同步的，你必须 *await* 在异步上下文中执行 `myAsyncFunction`，否则 `Node.js` 运行时会抱怨。您应该能够将代码复制并粘贴到 `test.js` 文件中，并使用 `node test.js` 从控制台运行它。

How can we achieve the same behavior with a Golang script?

我们如何使用 Golang 脚本实现相同的行为？

```go
package main

import (
    "fmt"
    "time"
)

func myAsyncFunction() <-chan int32 {
    r := make(chan int32)
    go func() {
        defer close(r)
        // func() core (meaning, the operation to be completed)
        time.Sleep(time.Second * 2)
        r <- 2
    }()
    return r
}

func main() {
    r := <-myAsyncFunction()
    // outputs `2` after two seconds
    fmt.Println(r)
}
```

As you can see, we used both a `goroutine` and a `channel`, introduced in the beginning. Let's see in detail the pattern used to  implement the async mechanism. First of all, the async function  explicitly returns a `<-chan [your_type]` where `your_type` could be whatever you want. In this case, it's a simple `int32` number. Within the function you want to run asynchronously, create a channel by using the `make(chan [your_type])` and return the created channel at the end of the function. Finally, start an anonymous goroutine by the `go myAsyncFunction() {...}` and implement the function's logic inside that anonymous function. Return the result by sending the value to the channel. At the beginning  of the anonymous function, add defer close(r) to close the channel once  done.

如您所见，我们使用了开头介绍的 `goroutine` 和 `channel`。让我们详细看看用于实现异步机制的模式。首先，异步函数显式返回一个 `<-chan [your_type]`，其中 `your_type` 可以是任何你想要的。在这种情况下，它是一个简单的“int32”数字。在要异步运行的函数中，使用`make(chan [your_type])`创建一个通道，并在函数末尾返回创建的通道。最后，通过 `go myAsyncFunction() {...}` 启动一个匿名 goroutine，并在该匿名函数中实现函数的逻辑。通过将值发送到通道来返回结果。在匿名函数的开头，添加 defer close(r) 以在完成后关闭通道。

To "await" behavior is implemented by simply read the value from channel, with `r := <-myAsyncFunction()`. And This Is It.

“等待”行为是通过简单地从通道读取值来实现的，使用 `r := <-myAsyncFunction()`。这就是它。

### Promise.all()

### Promise.all()

Unfortunately, things get more complicated as soon as you realized what you can do with `async/await`: another common scenario is when you start multiple async tasks then  wait for all of them to finish and gather their results. Doing that is  quite simple in Javascript (it is? it depends I guess). A pretty-simple  to describe a way to achieve it is by using the `Promise.all()` primitive:

不幸的是，一旦你意识到你可以用 `async/await` 做什么，事情就会变得更加复杂：另一种常见的情况是当你启动多个异步任务然后等待所有任务完成并收集它们的结果时。在 Javascript 中这样做非常简单（是吗？这取决于我猜）。一个非常简单的描述实现它的方法是使用 `Promise.all()` 原语：

```js
const myAsyncFunction = (s) => {
    return new Promise((resolve) => {
        setTimeout(() => resolve(s), 2000);
    })
};

(async function() {
    const result = await Promise.all([
        myAsyncFunction(2),
        myAsyncFunction(3)
    ]);
    // outputs `[2, 3]` after three seconds
    console.log(result);
})();
```

The await this time is done across a list of `Promises`: pay attention, because of the `.all()` signature takes an array as input. The `.all()` resolve all promises passed as an iterable object, short-circuits when  an input value is rejected, is resolved successfully when all the  promises in the array are resolved and rejected at first rejected of  them.

这次 await 是在一个 `Promises` 列表中完成的：注意，因为 `.all()` 签名将一个数组作为输入。 `.all()` 解析作为可迭代对象传递的所有承诺，当输入值被拒绝时短路，当数组中的所有承诺被解析并在第一次被拒绝时被拒绝时成功解析。

We achieve the same behavior with a Golang script:

我们使用 Golang 脚本实现了相同的行为：

```go
package main

import (
    "fmt"
    "time"
)

func myAsyncFunction(s int32) <-chan int32 {
    r := make(chan int32)
    go func() {
        defer close(r)
        // func() core (meaning, the operation to be completed)
        time.Sleep(time.Second * 2)
        r <- s
    }()
    return r
}

func main() {
    firstChannel, secondChannel := myAsyncFunction(2), myAsyncFunction(3)
    first, second := <-firstChannel, <-secondChannel
    // outputs `2, 3` after three seconds
    fmt.Println(first, second)
}
```

In both snippets of code we just packaged a function taking as  parameter the number of seconds to simulate a workload. The await is  implemented using the channels receive operation, nothing more than the `<-` operator.

在这两个代码片段中，我们只是打包了一个函数，将秒数作为参数来模拟工作负载。 await 是使用通道接收操作实现的，无非是 `<-` 操作符。

### Promise.race()

### Promise.race()

Sometimes, a piece of data can be received from several sources to  avoid high latencies, or there're cases that multiple results are  generated but they're equivalent and the only first response is  consumed. This first-response-win pattern is quite popular.

有时，可以从多个来源接收一条数据以避免高延迟，或者在某些情况下会生成多个结果但它们是等效的并且只使用第一个响应。这种第一反应赢的模式很受欢迎。

```js
const myAsyncFunction = (s) => {
    return new Promise((resolve) => {
        setTimeout(() => resolve(s), 2000);
    })
};

(async function() {
    const result = await Promise.race([
        myAsyncFunction(2),
        myAsyncFunction(3)
    ]);
    // outputs `2` after three seconds
    console.log(result);
})();
```

The expected behavior is that `2` is always returned before the second Promise returned by `myAsyncFunction(3)` got resolved. This is natural due to the nature of `.race()` that implements the first-win pattern mentioned above. In Golang, this  can be obtained similarly by using the select statement: let's make an  example.

预期的行为是在由 `myAsyncFunction(3)` 返回的第二个 Promise 得到解决之前总是返回 `2`。这是很自然的，因为实现了上面提到的先赢模式的 `.race()` 的性质。在 Golang 中，使用 select 语句也可以类似的获得：我们举个例子。

```go
package main

import (
    "fmt"
    "time"
)

func myAsyncFunction(s int32) <-chan int32 {
    r := make(chan int32)
    go func() {
        defer close(r)
        // func() core (meaning, the operation to be completed)
        time.Sleep(time.Second * 2)
        r <- s
    }()
    return r
}

func main() {
    var r int32
    select {
        case r = <-myAsyncFunction(2):
        case r = <-myAsyncFunction(3):
    }
    // outputs `2` after three seconds
    fmt.Println(r)
}
```

The cool thing about channels is that you can use Go's `select` statement to implement concurrency patterns and **wait on multiple channel operations**. In the snippet above, we use select to await both of the values  simultaneously, choosing, in this case, the first one that arrives: once again, `2` is always returned before a value appear is retrieved from the channel populated by the ` myAsyncFunction(3)`.

通道很酷的一点是你可以使用 Go 的 `select` 语句来实现并发模式并**等待多个通道操作**。在上面的代码片段中，我们使用 select 同时等待两个值，在这种情况下，选择第一个到达的值：再一次，在从由 ` myAsyncFunction(3)`。

However, we've seen that basic sends and receives on channels are blocking. We can use select with a `default` clause to implement non-blocking sends, receives, and even non-blocking multi-way selects. Let's take the example exposed by the gobyexample[6](https://madeddu.xyz/posts/go-async-await/#fn:golang-not-blocking-select) site.

但是，我们已经看到通道上的基本发送和接收是阻塞的。我们可以使用带有 `default` 子句的 select 来实现非阻塞发送、接收，甚至非阻塞多路选择。让我们以 gobyexample[6](https://madeddu.xyz/posts/go-async-await/#fn:golang-not-blocking-select) 站点公开的示例为例。

```go
package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
        case msg := <-messages:
            fmt.Println("received message", msg)
        default:
            fmt.Println("no message received")
    }

    msg := "hi"
    select {
        case messages <- msg:
            fmt.Println("sent message", msg)
        default:
            fmt.Println("no message sent")
    }

    select {
        case msg := <-messages:
            fmt.Println("received message", msg)
        case sig := <-signals:
            fmt.Println("received signal", sig)
        default:
            fmt.Println("no activity")
    }
}
```

The code above implements a non-blocking receive. If a value is available on `messages` then select will take the `<-messages` case with that value. If not it will immediately take the default case. A non-blocking send works similarly. Here `msg` cannot be sent to the messages channel, because the channel has no  buffer and there is no receiver. Therefore the default case is selected. We can use multiple cases above the default clause to implement a  multi-way non-blocking select. Here we attempt non-blocking receives on  both messages and signals.

上面的代码实现了一个非阻塞接收。如果在 `messages` 上有可用的值，则 select 将采用带有该值的 `<-messages` 大小写。如果不是，它将立即采用默认情况。非阻塞发送的工作原理类似。这里的 `msg` 无法发送到消息通道，因为通道没有缓冲区，也没有接收器。因此选择了默认情况。我们可以在 default 子句之上使用多个 case 来实现多路非阻塞选择。在这里，我们尝试对消息和信号进行非阻塞接收。

### Conclusion

###  结论

As you can see, the `await/async` basic patterns are easily portable to a Golang code. But... this was just a tasting: you can get so much more using `buffered channels`, `signals` and `context`. I will talk about all of this next time! Stay tuned and thank you for reading.

如您所见，`await/async` 基本模式很容易移植到 Golang 代码中。但是……这只是一次品尝：您可以使用“缓冲通道”、“信号”和“上下文”获得更多信息。下次我会讲这一切！请继续关注并感谢您的阅读。

If you like this post, please upvote it on HackerNews [here](https://news.ycombinator.com/submitted?id=made2591).

如果你喜欢这篇文章，请在 HackerNews [这里](https://news.ycombinator.com/submitted?id=made2591) 上点赞。

------

1. https://blog.golang.org/concurrency-is-not-parallelism [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-blog)
2. https://tour.golang.org/concurrency/1 [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-by-example-goroutines)
3. https://golang.org/doc/effective_go.html#goroutines [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-goroutines)
4. https://tour.golang.org/concurrency/2 [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-by-example-channels)
5. https://golang.org/doc/codewalk/sharemem/ [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-channels-codewalk)
6. https://gobyexample.com/non-blocking-channel-operations [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-not-blocking-select)

1. https://blog.golang.org/concurrency-is-not-parallelism [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-blog)
2. https://tour.golang.org/concurrency/1 [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-by-example-goroutines)
3. https://golang.org/doc/effective_go.html#goroutines [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-goroutines)
4. https://tour.golang.org/concurrency/2 [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-by-example-channels)
5. https://golang.org/doc/codewalk/sharemem/ [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-channels-codewalk)
6. https://gobyexample.com/non-blocking-channel-operations [[return\]](https://madeddu.xyz/posts/go-async-await/#fnref:golang-not-blocking-select)

Subscribe to [my newsletter](https://tinyletter.com/made2591) to be informed about my new blog posts, talks and activities. 

订阅 [我的时事通讯](https://tinyletter.com/made2591) 以了解我的新博文、演讲和活动。

