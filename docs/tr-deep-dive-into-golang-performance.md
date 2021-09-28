# Deep Dive into Golang Performance

# 深入研究 Golang 性能

February 22, 2021

2021 年 2 月 22 日

According to the [2020 StackOverflow Developer Survey](https://insights.stackoverflow.com/survey/2020#most-loved-dreaded-and-wanted) and the [TIOBE index](https://www.tiobe.com/tiobe-index/go/), Go (or Golang) has gained more traction in recent years, especially for backend developers and DevOps teams working on infrastructure automation. This article discusses what makes Go an attractive programming language for developers when it comes to performance.

根据 [2020 StackOverflow 开发者调查](https://insights.stackoverflow.com/survey/2020#most-loved-dreaded-and-wanted) 和 [TIOBE 指数](https://www.tiobe.com/tiobe-index/go/)，Go（或 Golang)近年来获得了更多关注，尤其是对于从事基础设施自动化的后端开发人员和 DevOps 团队而言。本文讨论了在性能方面，是什么让 Go 成为对开发人员有吸引力的编程语言。

## Introduction to Go

## Go 简介

First, let’s cover some of Go’s high-level properties.

首先，让我们介绍一些 Go 的高级属性。

### Go Facts

### 去事实

For those new to this programming language, some of the key facts you need to know about Golangs include:

对于这种编程语言的新手，您需要了解的有关 Golangs 的一些关键事实包括：

- _**Open-source:** Go's entire [implementation](https://github.com/golang/go) and [specification](https://golang.org/ref/spec) was published under licenses guaranteeing open access, which means that any user can observe its evolution. Go is also guided by the open-source community. Even Google imports Go code from the public repository._

- _**开源：** Go 的整个 [实现](https://github.com/golang/go) 和 [规范](https://golang.org/ref/spec) 是在保证许可的情况下发布的开放访问，这意味着任何用户都可以观察其演变。 Go 也受到开源社区的指导。甚至 Google 从公共存储库中导入 Go 代码。_

- **Backed by Google:** Created at Google, Go is currently maintained by Google developers, along with multiple contributors from the open-source community.

- **由 Google 支持：** 在 Google 创建，Go 目前由 Google 开发人员以及来自开源社区的多个贡献者维护。

- **Balanced:** The fact that Go is compiled, garbage-collected, statically typed, and natively concurrent makes it remarkable for its compilation and running time. But Go also has a beautifully clean syntax, which makes it an expressive language as well. It feels like writing code in an interpreted, [syntax-friendly](https://talks.golang.org/2015/simplicity-is-complicated.slide#11) language like Python, but with C performance application.

- **平衡：** Go 被编译、垃圾收集、静态类型和本地并发的事实使其在编译和运行时间方面非常出色。但是 Go 也有漂亮干净的语法，这也使它成为一种富有表现力的语言。感觉就像用解释性的、[语法友好](https://talks.golang.org/2015/simplicity-is-complicated.slide#11) 语言（如 Python)编写代码，但使用 C 性能应用程序。

- **Supported:** Go runs on Linux, Windows, and Mac OS. It is also supported by cloud providers like Google Cloud and AWS. Go is a [main citizen within the GCP ecosystem](https://cloud.google.com/go/home).

- **支持：** Go 在 Linux、Windows 和 Mac OS 上运行。谷歌云和 AWS 等云提供商也支持它。 Go 是 [GCP 生态系统中的主要公民](https://cloud.google.com/go/home)。

### Go in Production

### 投入生产

What products have been built using Go? Go is used for products that demand global scale, like the ubiquitous container tools [Docker](https://github.com/docker/cli) and [Kubernetes](https://github.com/kubernetes/kubernetes). The number of platforms currently running on top of Kubernetes says a lot about Go’s capabilities.

使用 Go 构建了哪些产品？ Go 用于需要全球规模的产品，如无处不在的容器工具 [Docker](https://github.com/docker/cli) 和 [Kubernetes](https://github.com/kubernetes/kubernetes)。当前在 Kubernetes 之上运行的平台数量很大程度上说明了 Go 的功能。

In the webspace, [Hugo](https://gohugo.io/) is billed as “the world’s fastest framework for building websites.” (Hugo is a static website generator that can serve pages in under 1 ms.)

在网络空间中，[Hugo](https://gohugo.io/) 被誉为“世界上最快的网站构建框架”。 （Hugo 是一个静态网站生成器，可以在 1 毫秒内提供页面。)

Go is also used in [CoreOS](https://github.com/coreos?language=go) (as of this writing, acquired by RedHat). And [Dropbox](https://github.com/dropbox?language=go) has a decent set of networking and infrastructure utilities written in Go.

[CoreOS](https://github.com/coreos?language=go) 中也使用了 Go（在撰写本文时，已被 RedHat 收购)。 [Dropbox](https://github.com/dropbox?language=go) 有一套用 Go 编写的不错的网络和基础设施实用程序。

And it should be no surprise that [Google relies heavily on Go](https://talks.golang.org/2013/go-sreops.slide#1). In 2012, in fact, Rob Pike noted, “Go is a programming language designed by Google to help solve Google’s problems, and Google has big problems.”

[Google 严重依赖 Go](https://talks.golang.org/2013/go-sreops.slide#1) 也就不足为奇了。事实上，Rob Pike 在 2012 年就指出，“Go 是谷歌设计的一种编程语言，旨在帮助解决谷歌的问题，而谷歌有大问题。”

### When to Use Go

### 何时使用 Go

As you may have surmised by now, when you need a backend component or service that supports global scaling, Go might be an option to consider. Go’s capabilities match with projects that:

正如您现在可能已经猜到的那样，当您需要一个支持全局扩展的后端组件或服务时，Go 可能是一个可以考虑的选择。 Go 的功能与以下项目相匹配：

- Involve networking or distributed processing.
- Involve REST and GRPC APIs running on cloud providers.
- Are related to infrastructure automation.
- Cover general tooling for OS or networking management.

- 涉及网络或分布式处理。
- 涉及在云提供商上运行的 REST 和 GRPC API。
- 与基础设施自动化有关。
- 涵盖操作系统或网络管理的通用工具。

## Go’s Performance Capabilities

## Go 的性能能力

Go has implemented different strategies on verticals like concurrency, system calls, task scheduling, and memory modeling, among others. All these strategies add up to a great balance between speed and robustness. But what makes Go one of the best programming languages when it comes to performance?

Go 在并发、系统调用、任务调度和内存建模等垂直领域实施了不同的策略。所有这些策略加起来在速度和鲁棒性之间取得了很好的平衡。但是，在性能方面，是什么让 Go 成为最好的编程语言之一？

### Concurrency 

### 并发

Go implements a variant of the [CSP model](https://medium.com/@niteshagarwal_/communicating-sequential-processes-golang-a3d6d5d4b25e), in which [channels](https://tour.golang.org/concurrency/2) are the preferred method for two Goroutines (a user space thread-like, with a few kilobytes in its stack) to share data. This approach is actually the opposite of that frequently used with other languages like Ruby or Python—a global shared data structure, with synchronization primitives for exclusive access (semaphores, locks, queues, etc.). Keeping these global data structures consistent across all units involves a lot of overhead.

Go 实现了 [CSP 模型](https://medium.com/@niteshagarwal_/communicating-sequential-processes-golang-a3d6d5d4b25e) 的变体，其中 [channels](https://tour.golang.org/concurrency/2) 是两个 Goroutines（类似用户空间的线程，堆栈中有几千字节）共享数据的首选方法。这种方法实际上与 Ruby 或 Python 等其他语言中经常使用的方法相反——一种全局共享数据结构，具有用于独占访问的同步原语（信号量、锁、队列等)。使这些全局数据结构在所有单元中保持一致涉及大量开销。

By following the CSP model, Go makes it possible to have concurrent constructions as primitives of the language. By default, Go knows how to deal with multiple tasks at once, and knows how to pass data between them. This, of course, translates to low latency with intercommunicating Goroutines. In Go, in the context of multithreading, you don’t write data to common storage. You create Goroutines to share data via channels. And because there is no need for exclusive access to global data structures, you gain speed.

通过遵循 CSP 模型，Go 可以将并发构造作为语言的原语。默认情况下，Go 知道如何同时处理多个任务，并且知道如何在它们之间传递数据。当然，这转化为互通 Goroutines 的低延迟。在 Go 中，在多线程上下文中，您不会将数据写入公共存储。您创建 Goroutines 以通过通道共享数据。并且因为不需要对全局数据结构进行独占访问，所以您可以获得速度。

It is important to note that you can also use mutex (or lock) mechanisms in Go, but that isn’t the [default approach](https://golang.org/ref/mem) for a concurrent program.

需要注意的是，您也可以在 Go 中使用互斥（或锁定）机制，但这不是并发程序的 [默认方法](https://golang.org/ref/mem)。

### Threading Model

### 线程模型

Go operates under an [M:N threading model](https://flylib.com/books/en/3.19.1.51/1/). In an M:N model, there are units of work under the user space (the Goroutines or G in the scheduler lexicon) which are scheduled to be run by the language runtime on OS threads (or M in the scheduler lexicon) on machine processors (or P in the scheduler lexicon). A Goroutine is defined as a lightweight thread managed by the Go runtime. Different Goroutines (G) can be executed on different OS threads (M), but at any given time, only one OS thread can be run on a CPU (P). In the user space, you achieve concurrency as the Goroutines work cooperatively. In the presence of a blocking operation (network, I/O or system call), another Goroutine can be assigned to the OS thread.

Go 在 [M:N 线程模型](https://flylib.com/books/en/3.19.1.51/1/) 下运行。在 M:N 模型中，用户空间下的工作单元（调度程序词典中的 Goroutines 或 G）被安排由机器处理器上的操作系统线程（或调度程序词典中的 M）上的语言运行时运行（或调度程序词典中的 P）。 Goroutine 被定义为由 Go 运行时管理的轻量级线程。不同的 Goroutines (G) 可以在不同的 OS 线程 (M) 上执行，但在任何给定时间，一个 CPU (P) 上只能运行一个 OS 线程。在用户空间中，您可以通过 Goroutine 协同工作来实现并发。在存在阻塞操作（网络、I/O 或系统调用)时，可以将另一个 Goroutine 分配给 OS 线程。

Once the blocking call ends, the runtime will try to reassign the previous Goroutine to an available OS thread. It’s possible to achieve parallelism here, because once the Goroutines are assigned to an OS thread, the OS can decide to distribute its threads’ execution through its multiple cores.

一旦阻塞调用结束，运行时将尝试将之前的 Goroutine 重新分配给可用的 OS 线程。在这里实现并行是可能的，因为一旦 Goroutine 被分配给一个 OS 线程，OS 就可以决定通过它的多个内核分配其线程的执行。

By having multiple Goroutines assigned to OS threads—thus being run cooperatively (or in parallel if two OS threads are run simultaneously on different cores)—you get an efficient use of your machine's CPUs, because all cores will be available for running your program's functions .

通过将多个 Goroutine 分配给 OS 线程——从而协作运行（如果两个 OS 线程在不同的内核上同时运行，则并行运行）——您可以有效地利用机器的 CPU，因为所有内核都可用于运行您的程序功能.

### Goroutines

### 协程

Goroutines live within the user thread space. In comparison to OS threads, their operations cost less: The overhead for assigning them, suspending them, and resuming them is lower than the overhead required by OS threads. Goroutines and channels are two of the most important primitives Go offers for concurrency. One important aspect of Goroutines is that expressing them in terms of code is fairly easy. You simply put the keyword go before the function you want to schedule to be run outside of the main thread.

Goroutines 存在于用户线程空间中。与 OS 线程相比，它们的操作成本更低：分配、挂起和恢复它们的开销低于 OS 线程所需的开销。 Goroutines 和通道是 Go 为并发提供的两个最重要的原语。 Goroutines 的一个重要方面是用代码表达它们相当容易。您只需将关键字 go 放在要安排在主线程之外运行的函数之前。

But how do Goroutines help make Go more performant? The [minimal stack required for a Goroutine](https://github.com/golang/go/blob/8f2db14cd35bbd674cb2988a508306de6655e425/src/runtime/stack.go#L72) to exist is 2 KB. Goroutines can increase their stack on runtime if they see the need for more space, but overall, they are memory-friendly. This means their management overhead is minimal. In other words, you can have more working units being processed with a decent quantity of memory, and that translates into efficiency and speed.

但是 Goroutines 如何帮助提高 Go 的性能？ [Goroutine 所需的最小堆栈](https://github.com/golang/go/blob/8f2db14cd35bbd674cb2988a508306de6655e425/src/runtime/stack.go#L72) 是 2 KB。如果 Goroutines 看到需要更多空间，它们可以在运行时增加它们的堆栈，但总的来说，它们是内存友好的。这意味着他们的管理开销很小。换句话说，您可以使用相当数量的内存处理更多工作单元，这转化为效率和速度。

### Task Scheduling 

### 任务调度

Go comes with its own [runtime scheduler.](https://morsmachine.dk/go-scheduler) The language does not rely on the native OS thread/process scheduler, but it cooperates with it. Because the scheduler is an independent component, it has the flexibility for implementing optimizations. All these optimizations aim for one thing: to avoid too much [preemption](https://medium.com/a-journey-with-go/go-goroutine-and-preemption-d6bc2aa2f4b7) of the OS Goroutines, which would result in suspending and resuming the functions' execution, an expensive operation.

Go 自带 [运行时调度器。](https://morsmachine.dk/go-scheduler) 该语言不依赖于原生 OS 线程/进程调度器，而是与之配合。因为调度器是一个独立的组件，所以它具有实现优化的灵活性。所有这些优化的目的是为了一件事：避免过多的 [抢占](https://medium.com/a-journey-with-go/go-goroutine-and-preemption-d6bc2aa2f4b7)，这会导致在暂停和恢复函数的执行时，这是一项昂贵的操作。

Next, we are going to highlight some specific optimizations done by the scheduler in order to avoid preemption.

接下来，我们将重点介绍调度程序为避免抢占而进行的一些特定优化。

### Work Stealing

### 工作窃取

Generally, there are two ways to distribute workloads across CPUs. The first one is _work sharing_, in which busy processors send threads to other, less busy processors with the hope they will be taken and executed. The second method is _work stealing_, in which an idle processor is constantly looking to steal other processor threads. [Go uses work stealing](https://github.com/golang/go/blob/f2eea4c1dc37886939c010daff89c03d5a3825be/src/runtime/proc.go#L2481).

通常，有两种方法可以跨 CPU 分配工作负载。第一个是_工作共享_，在这种情况下，繁忙的处理器将线程发送到其他不那么繁忙的处理器，希望它们能被占用并执行。第二种方法是_工作窃取_，其中空闲处理器不断地寻找窃取其他处理器线程。 [Go 使用工作窃取](https://github.com/golang/go/blob/f2eea4c1dc37886939c010daff89c03d5a3825be/src/runtime/proc.go#L2481)。

How does the work stealing approach help make Go faster? The migration of threads between processors is expensive, as it involves context switch operations. Under the stealing paradigm, this phenomenon occurs less frequently, resulting in less overhead.

工作窃取方法如何帮助使 Go 更快？处理器之间的线程迁移很昂贵，因为它涉及上下文切换操作。在窃取范式下，这种现象发生的频率较低，从而导致开销较少。

### Spinning Threads

### 旋转线程

The scheduler also implements a particular strategy called [spinning threads](https://github.com/golang/go/blob/f2eea4c1dc37886939c010daff89c03d5a3825be/src/runtime/proc.go#L54), which tries to fairly distribute as many OS threads across processors as possible. Go runtime not only reduces the frequency of thread migrations between processors, it is also capable of moving an OS thread with no work assigned to another processor. This can balance CPU usage and power.

调度程序还实现了一个称为 [旋转线程](https://github.com/golang/go/blob/f2eea4c1dc37886939c010daff89c03d5a3825be/src/runtime/proc.go#L54) 的特定策略，它试图公平地分配尽可能多的操作系统线程处理器尽可能。 Go 运行时不仅降低了处理器之间线程迁移的频率，还能够在没有工作分配给另一个处理器的情况下移动 OS 线程。这可以平衡 CPU 使用率和功率。

When you have all CPUs working with fairly distributed workloads, you are avoiding resource underutilization, which, again, translates to resource efficiency and speed.

当您让所有 CPU 处理相当分布的工作负载时，您就避免了资源未充分利用，这又转化为资源效率和速度。

### System Calls

### 系统调用

What strategy does the Go scheduler follow for handling [system calls](https://about.sourcegraph.com/go/a-go-guide-to-syscalls/)? It turns out that it also helps reduce overhead overall. Let’s see how.

Go 调度程序遵循什么策略来处理 [系统调用](https://about.sourcegraph.com/go/a-go-guide-to-syscalls/)？事实证明，它还有助于减少总体开销。让我们看看如何。

For system calls expected to be slow, the scheduler applies a pessimistic approach. It makes the OS thread release the processor in which it’s been running, just before the system call. Then, after the system call ends, the scheduler tries to reacquire the processor if it’s available. Otherwise, it’s enqueued by the scheduler until it finds a new available processor. The inconvenience of this approach is the overhead required for dropping and reacquiring a processor.

对于预期很慢的系统调用，调度程序采用悲观方法。它使操作系统线程在系统调用之前释放它正在运行的处理器。然后，在系统调用结束后，调度程序会尝试重新获取可用的处理器。否则，它会被调度器排入队列，直到找到新的可用处理器。这种方法的不便之处在于丢弃和重新获取处理器所需的开销。

However, the scheduler uses a second approach for system calls that are known to be fast—an optimistic approach. With this approach, the OS thread running the Goroutine with the system call does not release the processor, but it flags it. Then, after a few microseconds (20 to be precise), another independent special Goroutine (the sysmon Goroutine) checks for all flagged processors. If they are still running the heavy Goroutine that involves the system call, the scheduler takes their processors away, so they’re suspended. If the stolen processor is still available once the system call ends, the Goroutine can continue executing. Otherwise, it will need to be scheduled for execution again (until a processor becomes available).

然而，调度程序对已知速度快的系统调用使用第二种方法——乐观方法。使用这种方法，运行带有系统调用的 Goroutine 的 OS 线程不会释放处理器，而是标记它。然后，在几微秒（准确地说是 20 微秒）之后，另一个独立的特殊 Goroutine（sysmon Goroutine）会检查所有标记的处理器。如果它们仍在运行涉及系统调用的繁重 Goroutine，则调度程序会将它们的处理器拿走，因此它们会被挂起。如果系统调用结束后被盗处理器仍然可用，则 Goroutine 可以继续执行。否则，它将需要再次安排执行（直到处理器可用）。

### Conclusion

###  结论

In this article, we have covered the different strategies the Go language takes with concurrency and task scheduling. Go’s scheduler strategies and its [compiler optimizations](https://medium.com/a-journey-with-go/go-overview-of-the-compiler-4e5a153ca889) are what make Go so performant. Go’s balance between speed, robustness, and friendly syntax makes it a great option for specialized networking and web applications. 

在本文中，我们介绍了 Go 语言在并发和任务调度方面采用的不同策略。 Go 的调度器策略及其 [编译器优化](https://medium.com/a-journey-with-go/go-overview-of-the-compiler-4e5a153ca889) 使 Go 如此高效。 Go 在速度、健壮性和友好语法之间的平衡使其成为专业网络和 Web 应用程序的绝佳选择。

