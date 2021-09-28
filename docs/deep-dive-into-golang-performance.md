# Deep Dive into Golang Performance

February 22, 2021

According to the [2020 StackOverflow Developer Survey](https://insights.stackoverflow.com/survey/2020#most-loved-dreaded-and-wanted) and the [TIOBE index](https://www.tiobe.com/tiobe-index/go/), Go (or Golang) has gained more traction in recent years, especially for backend developers and DevOps teams working on infrastructure automation. This article discusses what makes Go an attractive programming language for developers when it comes to performance.

## Introduction to Go

First, let’s cover some of Go’s high-level properties.

### Go Facts

For those new to this programming language, some of the key facts you need to know about Golangs include:

- _**Open-source:** Go’s entire [implementation](https://github.com/golang/go) and [specification](https://golang.org/ref/spec) was published under licenses guaranteeing open access, which means that any user can observe its evolution. Go is also guided by the open-source community. Even Google imports Go code from the public repository._

- **Backed by Google:** Created at Google, Go is currently maintained by Google developers, along with multiple contributors from the open-source community.

- **Balanced:** The fact that Go is compiled, garbage-collected, statically typed, and natively concurrent makes it remarkable for its compilation and running time. But Go also has a beautifully clean syntax, which makes it an expressive language as well. It feels like writing code in an interpreted, [syntax-friendly](https://talks.golang.org/2015/simplicity-is-complicated.slide#11) language like Python, but with C performance application.

- **Supported:** Go runs on Linux, Windows, and Mac OS. It is also supported by cloud providers like Google Cloud and AWS. Go is a [main citizen within the GCP ecosystem](https://cloud.google.com/go/home).

### Go in Production

What products have been built using Go? Go is used for products that demand global scale, like the ubiquitous container tools [Docker](https://github.com/docker/cli) and [Kubernetes](https://github.com/kubernetes/kubernetes). The number of platforms currently running on top of Kubernetes says a lot about Go’s capabilities.

In the webspace, [Hugo](https://gohugo.io/) is billed as “the world’s fastest framework for building websites.” (Hugo is a static website generator that can serve pages in under 1 ms.)

Go is also used in [CoreOS](https://github.com/coreos?language=go) (as of this writing, acquired by RedHat). And [Dropbox](https://github.com/dropbox?language=go) has a decent set of networking and infrastructure utilities written in Go.

And it should be no surprise that [Google relies heavily on Go](https://talks.golang.org/2013/go-sreops.slide#1). In 2012, in fact, Rob Pike noted, “Go is a programming language designed by Google to help solve Google’s problems, and Google has big problems.”

### When to Use Go

As you may have surmised by now, when you need a backend component or service that supports global scaling, Go might be an option to consider. Go’s capabilities match with projects that:

- Involve networking or distributed processing.
- Involve REST and GRPC APIs running on cloud providers.
- Are related to infrastructure automation.
- Cover general tooling for OS or networking management.

## Go’s Performance Capabilities

Go has implemented different strategies on verticals like concurrency, system calls, task scheduling, and memory modeling, among others. All these strategies add up to a great balance between speed and robustness. But what makes Go one of the best programming languages when it comes to performance?

### Concurrency

Go implements a variant of the [CSP model](https://medium.com/@niteshagarwal_/communicating-sequential-processes-golang-a3d6d5d4b25e), in which [channels](https://tour.golang.org/concurrency/2) are the preferred method for two Goroutines (a user space thread-like, with a few kilobytes in its stack) to share data. This approach is actually the opposite of that frequently used with other languages like Ruby or Python—a global shared data structure, with synchronization primitives for exclusive access (semaphores, locks, queues, etc.). Keeping these global data structures consistent across all units involves a lot of overhead.

By following the CSP model, Go makes it possible to have concurrent constructions as primitives of the language. By default, Go knows how to deal with multiple tasks at once, and knows how to pass data between them. This, of course, translates to low latency with intercommunicating Goroutines. In Go, in the context of multithreading, you don’t write data to common storage. You create Goroutines to share data via channels. And because there is no need for exclusive access to global data structures, you gain speed.

It is important to note that you can also use mutex (or lock) mechanisms in Go, but that isn’t the [default approach](https://golang.org/ref/mem) for a concurrent program.

### Threading Model

Go operates under an [M:N threading model](https://flylib.com/books/en/3.19.1.51/1/). In an M:N model, there are units of work under the user space (the Goroutines or G in the scheduler lexicon) which are scheduled to be run by the language runtime on OS threads (or M in the scheduler lexicon) on machine processors (or P in the scheduler lexicon). A Goroutine is defined as a lightweight thread managed by the Go runtime. Different Goroutines (G) can be executed on different OS threads (M), but at any given time, only one OS thread can be run on a CPU (P). In the user space, you achieve concurrency as the Goroutines work cooperatively. In the presence of a blocking operation (network, I/O or system call), another Goroutine can be assigned to the OS thread.

Once the blocking call ends, the runtime will try to reassign the previous Goroutine to an available OS thread. It’s possible to achieve parallelism here, because once the Goroutines are assigned to an OS thread, the OS can decide to distribute its threads’ execution through its multiple cores.

By having multiple Goroutines assigned to OS threads—thus being run cooperatively (or in parallel if two OS threads are run simultaneously on different cores)—you get an efficient use of your machine’s CPUs, because all cores will be available for running your program’s functions.

### Goroutines

Goroutines live within the user thread space. In comparison to OS threads, their operations cost less: The overhead for assigning them, suspending them, and resuming them is lower than the overhead required by OS threads. Goroutines and channels are two of the most important primitives Go offers for concurrency. One important aspect of Goroutines is that expressing them in terms of code is fairly easy. You simply put the keyword go before the function you want to schedule to be run outside of the main thread.

But how do Goroutines help make Go more performant? The [minimal stack required for a Goroutine](https://github.com/golang/go/blob/8f2db14cd35bbd674cb2988a508306de6655e425/src/runtime/stack.go#L72) to exist is 2 KB. Goroutines can increase their stack on runtime if they see the need for more space, but overall, they are memory-friendly. This means their management overhead is minimal. In other words, you can have more working units being processed with a decent quantity of memory, and that translates into efficiency and speed.

### Task Scheduling

Go comes with its own [runtime scheduler.](https://morsmachine.dk/go-scheduler) The language does not rely on the native OS thread/process scheduler, but it cooperates with it. Because the scheduler is an independent component, it has the flexibility for implementing optimizations. All these optimizations aim for one thing: to avoid too much [preemption](https://medium.com/a-journey-with-go/go-goroutine-and-preemption-d6bc2aa2f4b7) of the OS Goroutines, which would result in suspending and resuming the functions’ execution, an expensive operation.

Next, we are going to highlight some specific optimizations done by the scheduler in order to avoid preemption.

### Work Stealing

Generally, there are two ways to distribute workloads across CPUs. The first one is _work sharing_, in which busy processors send threads to other, less busy processors with the hope they will be taken and executed. The second method is _work stealing_, in which an idle processor is constantly looking to steal other processor threads. [Go uses work stealing](https://github.com/golang/go/blob/f2eea4c1dc37886939c010daff89c03d5a3825be/src/runtime/proc.go#L2481).

How does the work stealing approach help make Go faster? The migration of threads between processors is expensive, as it involves context switch operations. Under the stealing paradigm, this phenomenon occurs less frequently, resulting in less overhead.

### Spinning Threads

The scheduler also implements a particular strategy called [spinning threads](https://github.com/golang/go/blob/f2eea4c1dc37886939c010daff89c03d5a3825be/src/runtime/proc.go#L54), which tries to fairly distribute as many OS threads across processors as possible. Go runtime not only reduces the frequency of thread migrations between processors, it is also capable of moving an OS thread with no work assigned to another processor. This can balance CPU usage and power.

When you have all CPUs working with fairly distributed workloads, you are avoiding resource underutilization, which, again, translates to resource efficiency and speed.

### System Calls

What strategy does the Go scheduler follow for handling [system calls](https://about.sourcegraph.com/go/a-go-guide-to-syscalls/)? It turns out that it also helps reduce overhead overall. Let’s see how.

For system calls expected to be slow, the scheduler applies a pessimistic approach. It makes the OS thread release the processor in which it’s been running, just before the system call. Then, after the system call ends, the scheduler tries to reacquire the processor if it’s available. Otherwise, it’s enqueued by the scheduler until it finds a new available processor. The inconvenience of this approach is the overhead required for dropping and reacquiring a processor.

However, the scheduler uses a second approach for system calls that are known to be fast—an optimistic approach. With this approach, the OS thread running the Goroutine with the system call does not release the processor, but it flags it. Then, after a few microseconds (20 to be precise), another independent special Goroutine (the sysmon Goroutine) checks for all flagged processors. If they are still running the heavy Goroutine that involves the system call, the scheduler takes their processors away, so they’re suspended. If the stolen processor is still available once the system call ends, the Goroutine can continue executing. Otherwise, it will need to be scheduled for execution again (until a processor becomes available).

### Conclusion

In this article, we have covered the different strategies the Go language takes with concurrency and task scheduling. Go’s scheduler strategies and its [compiler optimizations](https://medium.com/a-journey-with-go/go-overview-of-the-compiler-4e5a153ca889) are what make Go so performant. Go’s balance between speed, robustness, and friendly syntax makes it a great option for specialized networking and web applications.
