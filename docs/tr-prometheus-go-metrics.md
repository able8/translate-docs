# Exploring Prometheus Go client metrics - Povilas Versockas

# 探索 Prometheus Go 客户端指标 - Povilas Versockas

Published April 12, 2018

In this post I want to explore go metrics, which are exported by `client_golang` via `promhttp.Handler()` call. These metrics help you better understand how Go works & gives you better operational sanity, when your on call.

在这篇文章中，我想探索 go 指标，这些指标是通过 `promhttp.Handler()` 调用由 `client_golang` 导出的。这些指标可帮助您更好地了解 Go 的工作原理，并在您随叫随到时为您提供更好的运营意识。

Interested in learning more Prometheus? Check out [Monitoring Systems and Services with Prometheus](https://povilasv.me/out/prometheus), it’s awesome course that will get you up to speed.

有兴趣了解更多普罗米修斯吗？查看 [使用 Prometheus 监控系统和服务](https://povilasv.me/out/prometheus)，这是一门很棒的课程，可以让您快速上手。

Let’s start with a simple program, registering prom handler and listening on `8080` port:

让我们从一个简单的程序开始，注册 prom 处理程序并监听 8080 端口：

```
package main

import (
    "log"
    "net/http"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```


When you hit your metrics endpoint, you will get something like:

当你点击你的指标端点时，你会得到类似的东西：

```
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 3.5101e-05
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 6
...
process_open_fds 12
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.1272192e+07
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 4.74484736e+08
```


On initialisation `client_golang` registers 2 Prometheus collectors:

在初始化时，`client_golang` 注册了 2 个 Prometheus 收集器：

- [Process Collector](https://godoc.org/github.com/prometheus/client_golang/prometheus#NewProcessCollector) – which collects basic Linux process information like CPU, memory, file descriptor usage and start time.
- [Go Collector](https://godoc.org/github.com/prometheus/client_golang/prometheus#NewGoCollector) – which collects information about Go’s runtime like details about GC, number of gouroutines and OS threads.

- [Process Collector](https://godoc.org/github.com/prometheus/client_golang/prometheus#NewProcessCollector) – 收集基本的 Linux 进程信息，如 CPU、内存、文件描述符使用情况和启动时间。
- [Go Collector](https://godoc.org/github.com/prometheus/client_golang/prometheus#NewGoCollector) – 收集有关 Go 运行时的信息，例如有关 GC、gouroutine 和操作系统线程的详细信息。

## Process Collector

## 进程收集器

What this collector does is reads *proc* file system. *proc* file system exposes internal kernel data structures, which is used to obtain information about the system.[1](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23 #fn-1692-1)

这个收集器的作用是读取 *proc* 文件系统。 *proc* 文件系统公开内核内部数据结构，用于获取有关系统的信息。[1](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23 #fn-1692-1)

So Prometheus client reads `/proc/PID/stat` file, which looks like this:

所以 Prometheus 客户端读取 `/proc/PID/stat` 文件，它看起来像这样：

```
1 (sh) S 0 1 1 34816 8 4194560 674 43 9 1 5 0 0 0 20 0 1 0 89724 1581056 209 18446744073709551615 94672542621696 94672543427732 140730737801568 0 0 0 0 2637828 65538 1 0 0 17 3 0 0 0 0 0 94672545527192 94672545542787 94672557428736 140730737807231140730737807234 140730737807234 140730737807344 0
```


You can get human readable variant of this information using `cat /proc/PID/status`.

您可以使用 `cat /proc/PID/status` 获取此信息的人类可读变体。

**process_cpu_seconds_total** – it uses **utime** – number of ticks executing code in user mode, measured in jiffies, with **stime** – jiffies spent in the system mode, executing code on behalf of the process (like doing system calls). A *[jiffy](https://elinux.org/Kernel_Timer_Systems)* is the time between two ticks of the system timer interrupt. [2](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-2)

**process_cpu_seconds_total** – 它使用 **utime** – 在用户模式下执行代码的滴答数，以 jiffies 为单位，**stime** – 在系统模式下花费的 jiffies，代表进程执行代码（例如进行系统调用）。 A *[jiffy](https://elinux.org/Kernel_Timer_Systems)* 是系统定时器中断两次滴答之间的时间。 [2](关于:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-2)

**process_cpu_seconds_total** equals to sum of **utime** and **stime** and divide by USER_HZ. This makes sense, as dividing **number of scheduler ticks** by **Hz(ticks per second)** produces total time in seconds operating system has been running the process. [3](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-3)

**process_cpu_seconds_total** 等于 **utime** 和 **stime** 之和除以 USER_HZ。这是有道理的，因为将 ** 调度程序滴答数** 除以 **Hz（每秒滴答数）** 产生操作系统运行该进程的总时间（以秒为单位）。 [3](关于:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-3)

**process_virtual_memory_bytes** – uses **vsize** – virtual memory size is the amount of address space that a process is  managing. This includes all types of memory, both in RAM and swapped  out.

**process_virtual_memory_bytes** – 使用 **vsize** – 虚拟内存大小是进程管理的地址空间量。这包括所有类型的内存，包括 RAM 中的和换出的。

**process_resident_memory_bytes** – multiplies **rss** – resident set memory size is number of memory pages the process has in real memory, with pagesize [4](about:reader?url=https%3A%2F%2Fpovilasv.me %2Fprometheus-go-metrics%2F%23#fn-1692-4). This results in the amount of memory that belongs specifically to that  process in bytes. This excludes swapped out memory pages. 

**process_resident_memory_bytes** – 乘以 **rss** – 常驻集内存大小是进程在实际内存中的内存页数，pagesize [4](about:reader?url=https%3A%2F%2Fpovilasv.me %2Fprometheus-go-metrics%2F%23#fn-1692-4)。这会导致专门属于该进程的内存量（以字节为单位)。这不包括换出的内存页面。

**process_start_time_seconds** – uses **start_time** – time the process started after system boot, which is expressed in jiffies and **btime** from `/proc/stat` which shows the time at which the system booted in seconds since the Unix epoch. **start_time** is divided by *USER_HZ* in order to get the value in seconds.

**process_start_time_seconds** – 使用 **start_time** – 系统启动后进程启动的时间，以 jiffies 和 **btime** 表示，来自 `/proc/stat`，以秒为单位显示系统启动的时间自 Unix 时代以来。 **start_time** 除以 *USER_HZ* 以获得以秒为单位的值。

**process_open_fds** – counts the number of files in `/proc/PID/fd` directory. This shows currently open regular files, sockets, pseudo terminals, etc.

**process_open_fds** – 计算 `/proc/PID/fd` 目录中的文件数。这显示当前打开的常规文件、套接字、伪终端等。

**process_max_fds** – reads /proc/{PID}/limits and uses Soft Limit from “Max Open Files” row. The interesting bit here is that `/limits` lists Soft and Hard limits.
As it turns out, the soft limit is the value that the kernel enforces for  the corresponding resource and the Hard limit acts as a ceiling for the  soft limit.
An unprivileged process may only set its soft limit to a value up to the hard limit and (irreversibly) lower its Hard limit. [5](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-5)

**process_max_fds** – 读取 /proc/{PID}/limits 并使用“Max Open Files”行中的软限制。这里有趣的一点是`/limits` 列出了软限制和硬限制。
事实证明，软限制是内核为相应资源强制执行的值，而硬限制充当软限制的上限。
非特权进程只能将其软限制设置为硬限制并（不可逆转地）降低其硬限制。 [5](关于:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-5)

In Go you can use `err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{Cur: 9, Max: 10})` to set your limit.

在 Go 中，你可以使用 `err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{Cur: 9, Max: 10})` 来设置你的限制。

## Go Collector



Go Collector’s most of the metrics are taken from `runtime`, `runtime/debug` packages.

Go Collector 的大部分指标取自 `runtime`、`runtime/debug` 包。

**go_goroutines** – calls out to `runtime.NumGoroutine()`, which computes the value based off the scheduler struct and global `allglen` variable. As all the values in sched struct can be changed concurently  there is this funny check where if computed value is less than 1 it  becomes 1.

**go_goroutines** – 调用 `runtime.NumGoroutine()`，它根据调度程序结构和全局 `allglen` 变量计算值。由于 sched 结构中的所有值都可以同时更改，因此有一个有趣的检查，如果计算值小于 1，则它变为 1。

**go_threads** – calls out to `runtime.CreateThreadProfile()`, which reads off global `allm` variable. If you don’t know what M’s or G’s you can read my [blogpost about it](https://povilasv.me/go-scheduler/).

**go_threads** – 调用 `runtime.CreateThreadProfile()`，它读取全局 `allm` 变量。如果您不知道 M 或 G 是什么，您可以阅读我的 [关于它的博文](https://povilasv.me/go-scheduler/)。

**go_gc_duration_seconds** – calls out to `debug.ReadGCStats()` with `PauseQuantile` set to 5, which returns us the minimum, 25%, 50%, 75%, and maximum  pause times. Then it manualy creates a Summary type from pause  quantiles, `NumGC` var and `PauseTotal` seconds. It’s cool how well `GCStats` struct fits the prom’s Summary type. [6](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-6)

**go_gc_duration_seconds** – 调用 `debug.ReadGCStats()`，并将 `PauseQuantile` 设置为 5，这会返回最小、25%、50%、75% 和最大暂停时间。然后它根据暂停分位数、`NumGC` var 和 `PauseTotal` 秒手动创建一个摘要类型。 `GCStats` 结构体非常适合舞会的摘要类型，这很酷。 [6](关于:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-6)

**go_info** – this provides us with Go version. It’s pretty clever, it calls out to `runtime.Version()` and set’s that as a `version` label and then always returns value of `1` for this gauge metric.

**go_info** – 这为我们提供了 Go 版本。它非常聪明，它调用 `runtime.Version()` 并将其设置为 `version` 标签，然后始终为该仪表指标返回值 `1`。

### Memory



Go Collector provides us with a lot of metrics about memory and GC.
All those metrics are from `runtime.ReadMemStats()`, which gives us metrics from [MemStats](https://golang.org/pkg/runtime/#MemStats) struct.
One thing that worries me, is that `runtime.ReadMemStats()` has a explicit call to make a stop-the-world pause [7](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus- go-metrics%2F%23#fn-1692-7).
So I wonder how much actual cost this pause introduces?
As during stop-the-world pause, all goroutines are paused, so that GC can  run. I’ll probably do a comparison of an app with and without  instrumentation in a later post.

Go Collector 为我们提供了很多关于内存和 GC 的指标。
所有这些指标都来自 `runtime.ReadMemStats()`，它为我们提供了来自 [MemStats](https://golang.org/pkg/runtime/#MemStats) 结构的指标。
令我担心的一件事是，`runtime.ReadMemStats()` 有一个显式调用来停止世界暂停 [7](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus- go-metrics%2F%23#fn-1692-7)。
所以我想知道这次暂停会带来多少实际成本？
在 stop-the-world 暂停期间，所有 goroutine 都会暂停，以便 GC 可以运行。我可能会在以后的文章中对有和没有仪器的应用程序进行比较。

We already seen that Linux provides us with rss/vsize metrics for  memory stats, so naturally the question arises which metrics to use, the ones provided in MemStats or rss/vsize?

我们已经看到 Linux 为我们提供了内存统计数据的 rss/vsize 指标，所以很自然地出现了使用哪些指标的问题，在 MemStats 或 rss/vsize 中提供的指标？

The good part about resident set size and virtual memory size is that it’s based off Linux primitives and is programming language agnostic.
So in theory you could instrument any program and you would know how much  memory it consumes (as long as you name your metrics consistently, ie *process_virtual_memory_bytes* and *process_resident_memory_bytes*.).
In practice, however when Go process starts up it takes a lot of virtual  memory beforehand, such a simple program like the one above takes up to  544MiB of vsize on my machine (x86_64 Ubuntu), which is a bit confusing. RSS shows around 7mib, which is closer to the actual usage.

常驻集大小和虚拟内存大小的好处在于它基于 Linux 原语并且与编程语言无关。
所以理论上你可以检测任何程序并且你会知道它消耗了多少内存（只要你一致地命名你的指标，即 *process_virtual_memory_bytes* 和 *process_resident_memory_bytes*。）。
但是在实际中，Go 进程启动时会预先占用大量虚拟内存，像上面这样的简单程序在我的机器（x86_64 Ubuntu）上占用了 544MiB 的 vsize，这有点令人困惑。 RSS显示在7mib左右，更接近实际使用情况。

On the other hand using Go runtime based metrics gives more fined  grained information on what is happening in your running application.
You should be able to find out more easily whether your program has a memory leak, how long GC took, how much it reclaimed.
Also, it should point you into right direction when you are optimizing program’s memory allocations. 

另一方面，使用基于 Go 运行时的指标可以提供有关正在运行的应用程序中发生的事情的更细粒度的信息。
您应该能够更轻松地找出您的程序是否存在内存泄漏、GC 花费了多长时间、回收了多少。
此外，当您优化程序的内存分配时，它应该为您指明正确的方向。

I haven't looked in detail how Go GC and memory model works, a part of it's concurrency model [8](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23 #fn-1692-8), so this bit is still new to me.

我没有详细研究 Go GC 和内存模型是如何工作的，它是并发模型的一部分 [8](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23 #fn-1692-8)，所以这一点对我来说仍然是新的。

So let’s take a look at those metrics:

那么让我们来看看这些指标：

**go_memstats_alloc_bytes** – a metric which shows how much bytes of memory is allocated on the [Heap](https://en.wikipedia.org/wiki/Memory_management#HEAP) for the Objects. The value is same as **go_memstats_heap_alloc_bytes**. This metric counts all reachable heap objects plus unreachable objects, GC has not yet freed.

**go_memstats_alloc_bytes** – 一个指标，显示在 [堆](https://en.wikipedia.org/wiki/Memory_management#HEAP) 上为对象分配了多少字节的内存。该值与 **go_memstats_heap_alloc_bytes** 相同。该指标计算所有可到达的堆对象加上不可到达的对象，GC 尚未释放。

**go_memstats_alloc_bytes_total** – this metric  increases as objects are allocated in the Heap, but doesn’t decrease  when they are freed. I think it is immensly useful, as it is only  increasing number and has same nice properties that [Prometheus Counter](https://povilasv.me/prometheus-tracking-request-duration/) has. Doing `rate()` on it should show us how many bytes/s of memory app consumes and is “durable” across restarts and scrape misses.

**go_memstats_alloc_bytes_total** – 该指标随着对象在堆中分配而增加，但在释放对象时不会减少。我认为它非常有用，因为它的数量只会增加，并且具有与 [Prometheus Counter](https://povilasv.me/prometheus-tracking-request-duration/) 相同的良好属性。对它执行 `rate()` 应该向我们展示应用程序消耗多少字节/秒的内存，并且在重新启动和抓取未命中时是“持久的”。

**go_memstats_sys_bytes** – it’s a metric, which measures how many bytes of memory in total is taken from system by Go. It sums all the **sys** metrics described below.

**go_memstats_sys_bytes** – 它是一个度量标准，用于衡量 Go 从系统中获取的内存字节总数。它总结了下面描述的所有 **sys** 指标。

**go_memstats_lookups_total** – counts how many pointer dereferences happened. This is a counter value so you can use `rate()` to lookups/s.

**go_memstats_lookups_total** – 计算指针取消引用发生的次数。这是一个计数器值，因此您可以使用 `rate()` 来查找/秒。

**go_memstats_mallocs_total** – shows how many heap objects are allocated. This is a counter value so you can use `rate()` to objects allocated/s.

**go_memstats_mallocs_total** – 显示分配了多少堆对象。这是一个计数器值，因此您可以对分配的对象使用 `rate()`。

**go_memstats_frees_total** – shows how many heap objects are freed. This is a counter value so you can use `rate()` to objects allocated/s. Note you can get number of live objects with `go_memstats_mallocs_total` – `go_memstats_frees_total`.

**go_memstats_frees_total** – 显示释放了多少堆对象。这是一个计数器值，因此您可以对分配的对象使用 `rate()`。请注意，您可以使用 `go_memstats_mallocs_total` – `go_memstats_frees_total` 获取活动对象的数量。

Turns out that Go organizes memory in spans, which are contiguous regions of memory of 8K or larger. There are 3 types of Spans:
\1) **idle** – span, that has no objects and can be released back to the OS, or reused for heap allocation, or reused for stack memory.
\2) **in use** – span, that has atleast one heap object and may have space for more.
\3) **stack** – span, which is used for goroutine stack. This span can live either in stack or in heap, but not in both.

事实证明，Go 以跨度组织内存，跨度是 8K 或更大的连续内存区域。有 3 种类型的 Span：
\1) **idle** – span，没有对象，可以释放回操作系统，或重用于堆分配，或重用于堆栈内存。
\2) **in use** – 跨度，至少有一个堆对象并且可能有更多空间。
\3) **stack** – span，用于 goroutine 栈。此跨度可以存在于堆栈中或堆中，但不能同时存在于两者中。

### Heap memory metrics

### 堆内存指标

**go_memstats_heap_alloc_bytes** – same as **go_memstats_alloc_bytes**.

**go_memstats_heap_alloc_bytes** – 与 **go_memstats_alloc_bytes** 相同。

**go_memstats_heap_sys_bytes** – bytes of memory obtained for the heap from OS. This includes [virtual address space](https://en.wikipedia.org/wiki/Virtual_address_space) that has been resevered, but not yet used and virtual address space  which was returned to OS after it became unused. This metric estimates  the largest size the heap.

**go_memstats_heap_sys_bytes** – 从操作系统获得的堆内存字节数。这包括已保留但尚未使用的 [虚拟地址空间](https://en.wikipedia.org/wiki/Virtual_address_space) 以及在未使用后返回给操作系统的虚拟地址空间。该指标估计堆的最大大小。

**go_memstats_heap_idle_bytes** – shows how many bytes are in **idle** spans.

**go_memstats_heap_idle_bytes** – 显示 **idle** 跨度中有多少字节。

**go_memstats_heap_idle_bytes** minus **go_memstats_heap_released_bytes** estimates how many bytes of memory could be released, but is kept by  runtime, so that runtime can allocate objects on the heap without asking OS for more memory.

**go_memstats_heap_idle_bytes** 减去 **go_memstats_heap_released_bytes** 估计可以释放多少字节的内存，但由运行时保留，以便运行时可以在堆上分配对象，而无需向操作系统请求更多内存。

**go_memstats_heap_inuse_bytes** – shows how many bytes in **in-use** spans.

**go_memstats_heap_inuse_bytes** – 显示 **in-use** span 中的字节数。

**go_memstats_heap_inuse_bytes** minus **go_memstats_heap_alloc_bytes** shows how many bytes of memory has been allocated for the heap, but not currently used.

**go_memstats_heap_inuse_bytes** 减去 **go_memstats_heap_alloc_bytes** 显示已为堆分配但当前未使用的内存字节数。

**go_memstats_heap_released_bytes** – shows how many bytes of idle spans were returned to the OS.

**go_memstats_heap_released_bytes** – 显示返回给操作系统的空闲跨度字节数。

**go_memstats_heap_objects** – shows how many objects are allocated on the heap. This changes as GC is performed and new objects are allocated.

**go_memstats_heap_objects** – 显示在堆上分配了多少对象。这会随着 GC 的执行和新对象的分配而改变。

### Stack memory metrics

### 堆栈内存指标

**go_memstats_stack_inuse_bytes** – shows how many bytes of memory is used by stack memory spans, which have atleast one object  in them. Go doc says, that stack memory spans can only be used for other stack spans, i.e. there is no mixing of heap objects and stack objects  in one memory span.

**go_memstats_stack_inuse_bytes** – 显示堆栈内存跨度使用了多少字节的内存，其中至少有一个对象。 Go doc 说，堆栈内存跨度只能用于其他堆栈跨度，即在一个内存跨度中不能混合堆对象和堆栈对象。

**go_memstats_stack_sys_bytes** – shows how many bytes of stack memory is obtained from OS. It’s **go_memstats_stack_inuse_bytes** plus any memory obtained for OS thread stack.

**go_memstats_stack_sys_bytes** – 显示从操作系统获得的堆栈内存字节数。它是 **go_memstats_stack_inuse_bytes** 加上为操作系统线程堆栈获得的任何内存。

There is no **go_memstats_stack_idle_bytes**, as unused stack spans are counted towards **go_memstats_heap_idle_bytes**.

没有 **go_memstats_stack_idle_bytes**，因为未使用的堆栈跨度计入 **go_memstats_heap_idle_bytes**。

### Off-heap memory metrics

### 堆外内存指标

These metrics are bytes allocated for runtime internal structures,  that are not allocated on the heap, because they implement the heap.

这些指标是为运行时内部结构分配的字节，它们没有在堆上分配，因为它们实现了堆。

**go_memstats_mspan_inuse_bytes** – shows how many bytes are in use by `mspan` structures. 

**go_memstats_mspan_inuse_bytes** – 显示 `mspan` 结构正在使用多少字节。

**go_memstats_mspan_sys_bytes** – shows how many bytes are obtained from OS for `mspan` structures.

**go_memstats_mspan_sys_bytes** – 显示从操作系统获取的 `mspan` 结构的字节数。

**go_memstats_mcache_inuse_bytes** – shows how many bytes are in use by `mcache` structures.

**go_memstats_mcache_inuse_bytes** – 显示 `mcache` 结构正在使用多少字节。

**go_memstats_mcache_sys_bytes** – shows how many bytes are obtained from OS for `mcache` structures.

**go_memstats_mcache_sys_bytes** – 显示从操作系统获取的 `mcache` 结构的字节数。

**go_memstats_buck_hash_sys_bytes** – shows how many bytes of memory are in bucket hash tables, which are used for profiling.

**go_memstats_buck_hash_sys_bytes** – 显示用于分析的存储桶哈希表中的内存字节数。

**go_memstats_gc_sys_bytes** – shows how many in garbage collection metadata.

**go_memstats_gc_sys_bytes** – 显示垃圾收集元数据中有多少。

**go_memstats_other_sys_bytes** – **go_memstats_other_sys_bytes** shows how many bytes of memory are used for other runtime allocations.

**go_memstats_other_sys_bytes** – **go_memstats_other_sys_bytes** 显示用于其他运行时分配的内存字节数。

**go_memstats_next_gc_bytes** – shows the target heap size of the next GC cycle. GC’s goal is to keep **go_memstats_heap_alloc_bytes** less than this value.

**go_memstats_next_gc_bytes** – 显示下一个 GC 周期的目标堆大小。 GC 的目标是让 **go_memstats_heap_alloc_bytes** 小于这个值。

**go_memstats_last_gc_time_seconds** – contains unix timestamp when last GC finished.

**go_memstats_last_gc_time_seconds** – 包含上次 GC 完成时的 unix 时间戳。

**go_memstats_last_gc_cpu_fraction** – shows the fraction of this program’s available CPU time used by GC since the program started.
This metric is also provided in GODEBUG=gctrace=1.

**go_memstats_last_gc_cpu_fraction** – 显示自程序启动以来 GC 使用的该程序可用 CPU 时间的比例。
该指标也在 GODEBUG=gctrace=1 中提供。

## Playing around with numbers

## 玩数字

So it’s a lot of metrics and a lot of information.
I think the best way to learn is to just play around with it, so in this part I’ll do just that.
So I’ll be using the same program that is above.
Here are the dump from `/metrics` (edited for space), which I’m going to use:

所以它有很多指标和很多信息。
我认为最好的学习方法就是玩弄它，所以在这部分我会这样做。
所以我将使用与上面相同的程序。
以下是来自`/metrics`（为空间编辑）的转储，我将使用它：

```
process_resident_memory_bytes 1.09568e+07

process_virtual_memory_bytes 6.46668288e+08

go_memstats_heap_alloc_bytes 2.24344e+06

go_memstats_heap_idle_bytes 6.3643648e+07

go_memstats_heap_inuse_bytes 3.039232e+06

go_memstats_heap_objects 6498

go_memstats_heap_released_bytes 0

go_memstats_heap_sys_bytes 6.668288e+07

go_memstats_lookups_total 0

go_memstats_frees_total 12209

go_memstats_mallocs_total 18707

go_memstats_buck_hash_sys_bytes 1.443899e+06

go_memstats_mcache_inuse_bytes 6912

go_memstats_mcache_sys_bytes 16384

go_memstats_mspan_inuse_bytes 25840

go_memstats_mspan_sys_bytes 32768

go_memstats_other_sys_bytes 1.310909e+06

go_memstats_stack_inuse_bytes 425984

go_memstats_stack_sys_bytes 425984

go_memstats_sys_bytes 7.2284408e+07

go_memstats_next_gc_bytes 4.194304e+06

go_memstats_gc_cpu_fraction 1.421928536233557e-06

go_memstats_gc_sys_bytes 2.371584e+06

go_memstats_last_gc_time_seconds 1.5235057190167596e+09
```


rss = 1.09568e+07 = 10956800 bytes = 10700 KiB = 10.4 MiB

vsize = 6.46668288e+08 = 646668288 bytes = 631512 KiB = 616.7 MiB

heap_alloc_bytes = 2.24344e+06 = 2243440 = 2190 KiB = 2.1 MiB

heap_inuse_bytes = 3.039232e+06 = 3039232 = 2968 KiB = 2,9 MiB

heap_idle_bytes = 6.3643648e+07 = 63643648 = 62152 KiB = 60.6 MiB

heap_released_bytes = 0

heap_sys_bytes = 6.668288e+07 = 66682880 = 65120 KiB = 63.6 MiB

frees_total = 12209

mallocs_total = 18707

mallocs_total = 18707

mspan_inuse_bytes = 25840 = 25.2 KiB

mspan_inuse_bytes = 25840 = 25.2 KiB

mspan_sys_bytes = 32768 = 32 KiB

mspan_sys_bytes = 32768 = 32 KiB

mcache_inuse_bytes = 6912 = 6.8 KiB

mcache_inuse_bytes = 6912 = 6.8 KiB

mcache_sys_bytes = 16384 = 12 KiB

mcache_sys_bytes = 16384 = 12 KiB

buck_hash_sys_bytes = 1.443899e+06 = 1443899 = 1410 KiB = 1.4 MiB

buck_hash_sys_bytes = 1.443899e+06 = 1443899 = 1410 KiB = 1.4 MiB

gc_sys_bytes = 2.371584e+06 = 2371584 = 2316 KiB = 2.3 MiB

gc_sys_bytes = 2.371584e+06 = 2371584 = 2316 KiB = 2.3 MiB

other_sys_bytes = 1.310909e+06 = 1310909 = 1280,2 KiB = 1.3MiB

other_sys_bytes = 1.310909e+06 = 1310909 = 1280,2 KiB = 1.3MiB

stack_inuse_bytes = 425984 = 416 KiB

stack_inuse_bytes = 425984 = 416 KiB

stack_sys_bytes = 425984 = 416 KiB

stack_sys_bytes = 425984 = 416 KiB

sys_bytes = 7.2284408e+07 = 72284408 = 70590.2 KiB = 68.9 MiB

sys_bytes = 7.2284408e+07 = 72284408 = 70590.2 KiB = 68.9 MiB

next_gc_bytes = 4.194304e+06 = 4194304 = 4096 KiB = 4 MiB

next_gc_bytes = 4.194304e+06 = 4194304 = 4096 KiB = 4 MiB

gc_cpu_fraction = 1.421928536233557e-06 = 0.000001

gc_cpu_fraction = 1.421928536233557e-06 = 0.000001

last_gc_time_seconds = 1.5235057190167596e+09 = Thu, 12 Apr 2018 05:47:59 GMT

last_gc_time_seconds = 1.5235057190167596e+09 = 2018 年 4 月 12 日星期四 05:47:59 GMT

Interesting bit is that **heap_inuse_bytes** is more than **heap_alloc_bytes**.
I think **heap_alloc_bytes** shows how many bytes are in terms of **objects** and **heap_inuse_bytes** shows bytes of memory in terms of **spans**.
Dividing **heap_inuse_bytes** by size of the span gives: 3039232 / 8192 = 371 span.

有趣的是 **heap_inuse_bytes** 不仅仅是 **heap_alloc_bytes**。
我认为 **heap_alloc_bytes** 显示了 **objects** 的字节数，而 **heap_inuse_bytes** 显示了 **spans** 的内存字节数。
**heap_inuse_bytes** 除以跨度的大小得出：3039232 / 8192 = 371 跨度。

**heap_inuse_bytes** minus **heap_alloc_bytes**, should show the amount of free space that we have in in-use spans, which is 2,9 MiB – 2.1 MiB = 0.8 MiB.
This roughly means that we can allocate 0.8 MiB of objects on the heap without using new memory spans.
But we should keep in mind memory fragmentation.
Imagine if you have a new bytes slice of 10K bytes, the memory could be in the  position, where it doesn’t have a contiguous block of 10K bytes + slice  header, so it would need use a new span, instead of reusing 

**heap_inuse_bytes** 减去 **heap_alloc_bytes**，应显示我们在使用中的跨度中的可用空间量，即 2.9 MiB – 2.1 MiB = 0.8 MiB。
这大致意味着我们可以在不使用新内存跨度的情况下在堆上分配 0.8 MiB 的对象。
但是我们应该记住内存碎片。
想象一下，如果你有一个 10K 字节的新字节切片，内存可能位于没有 10K 字节的连续块 + 切片头的位置，因此它需要使用新的跨度，而不是重用

**heap_idle_bytes** minus **heap_released_byte** shows that we have around 60.6 MiB of unused spans, which are reserved  from OS and could be returned to OS. It’s 63643648/8192 = 7769 spans.

**heap_idle_bytes** 减去 **heap_released_byte** 表明我们有大约 60.6 MiB 的未使用跨度，它们是从操作系统保留的，可以返回给操作系统。它是 63643648/8192 = 7769 个跨度。

**heap_sys_bytes**, which is 63.6MiB estimates the largest size the heap has had. It’s 66682880/8192 = 8140 spans.

**heap_sys_bytes**，这是 63.6MiB 估计堆的最大大小。它是 66682880/8192 = 8140 个跨度。

**mallocs_total** shows that we allocated 18707 objects  and freed 12209. So currently, we have 18707-12209 = 6498 objects. We  can find the average size of the object dividing *heap_alloc_bytes* over live objects, which is 6498. The result is 2243440 / 6498 = 345.3 bytes.
(This is probably a stupid metric, as object vary a lot in size and we should do histograms instead)

**mallocs_total** 显示我们分配了 18707 个对象并释放了 12209 个。所以目前，我们有 18707-12209 = 6498 个对象。我们可以找到在活动对象上划分 *heap_alloc_bytes* 的对象的平均大小，即 6498。结果是 2243440 / 6498 = 345.3 字节。
（这可能是一个愚蠢的指标，因为对象的大小变化很大，我们应该做直方图）

So **sys_bytes** should be a sum of all ***sys** metrics. So let’s check that:
sys_bytes == mspan_sys_bytes + mcache_sys_bytes + buck_hash_sys_bytes +  gc_sys_bytes + other_sys_bytes + stack_sys_bytes + heap_sys_bytes.
So, we have 72284408 == 32768 + 16384 + 1443899 + 2371584 + 1310909 +  425984 + 66682880, which is 72284408 == 72284408, which is correct.

所以 **sys_bytes** 应该是所有 ***sys** 指标的总和。所以让我们检查一下：
sys_bytes == mspan_sys_bytes + mcache_sys_bytes + buck_hash_sys_bytes + gc_sys_bytes + other_sys_bytes + stack_sys_bytes + heap_sys_bytes。
所以，我们有 72284408 == 32768 + 16384 + 1443899 + 2371584 + 1310909 + 425984 + 66682880，即 72284408 == 72284408，这是正确的。

The interesting detail about **sys_bytes**, is that it’s 68,9 MiB it’s how many bytes of memory in total taken from OS. Meanwhile, OS’es **vsize** gives you 616,7MiB and rss 10.4 MiB. So all these numbers don’t really match up.

关于 **sys_bytes** 的一个有趣的细节是，它是 68,9 MiB，它是从操作系统中总共占用了多少字节的内存。同时，操作系统的 **vsize** 为您提供 616,7MiB 和 rss 10.4 MiB。所以所有这些数字并不真正匹配。

As I understand it so part of our memory could be in OS’s memory  pages which are in swap or in filesystem (not in RAM), so this would  explain why rss is smaller than **sys_bytes**.

据我了解，我们的内存的一部分可能位于 OS 的内存页面中，这些页面位于交换或文件系统中（不在 RAM 中），所以这可以解释为什么 rss 小于 **sys_bytes**。

And **vsize** contains a lot of things, like mapped libc, pthreads libs, etc. You can explore `/proc/PID/maps` and `/proc/PID/smaps` file, to see what is being currently mapped .

**vsize** 包含很多东西，比如映射的 libc、pthreads 库等。你可以浏览 `/proc/PID/maps` 和 `/proc/PID/smaps` 文件，看看当前正在映射什么.

**gc_cpu_fraction** is running crazy low, `0.000001` of CPU time is used for GC. That’s really really cool. (Although this program doesn’t produce much garbage)

**gc_cpu_fraction** 运行得非常低，“0.000001”的 CPU 时间用于 GC。这真的很酷。 （虽然这个程序不会产生太多垃圾）

**next_gc_bytes** shows that the target for GC is to keep **heap_alloc_bytes** under 4 MiB, as **heap_alloc_bytes** is currently at 2.1 MiB the target is achieved.

**next_gc_bytes** 表明 GC 的目标是将 **heap_alloc_bytes** 保持在 4 MiB 以下，因为 **heap_alloc_bytes** 目前为 2.1 MiB，目标已实现。

## Conclusion

##  结论

I love Go and the fact that it exposes so much useful information in  it’s packages and users like you and me can just call a function and get the information. Alos Prometheus is really a great tool to monitor  applications.

我喜欢 Go 并且事实上它在它的包中公开了这么多有用的信息，像你我这样的用户只需调用一个函数就可以获取信息。 Alos Prometheus 确实是监控应用程序的绝佳工具。

Do you want get better at Prometheus? Check out [Monitoring Systems and Services with Prometheus](https://povilasv.me/out/prometheus). I definitely recommend the module.

你想在普罗米修斯变得更好吗？查看 [使用 Prometheus 监控系统和服务](https://povilasv.me/out/prometheus)。我绝对推荐该模块。

It was really cool playing around and reading about Linux & Go,  so I’m thinking of doing part 2 of this post. Maybe go look into metrics provided by [cAdvisor](https://github.com/google/cadvisor) or show how to use some of the metrics described here in dashboards/alerts with Prometheus.

玩和阅读 Linux & Go 真的很酷，所以我正在考虑做这篇文章的第 2 部分。也许可以查看 [cAdvisor](https://github.com/google/cadvisor) 提供的指标，或者展示如何使用 Prometheus 的仪表板/警报中描述的一些指标。

Also, once [vgo](https://research.swtch.com/vgo) get’s  integrated (and I really really hope it does, cause it’s like the best  package manager I ever used). Then we should be able to inspect  dependencies from some go runtime package, which would be really cool! Imagine writing a custom prom collector, which would go through all your dependencies, check for new versions and if found wouldgive you back a  number of outdated pkgs, something like `go_num_outdated_pkgs` metric.
This way you could write an alert if your service get’s terribly outdated. Or check that your live dependency hashes don’t match current hashes?

此外，一旦 [vgo](https://research.swtch.com/vgo)被集成（我真的很希望它能做到，因为它就像我用过的最好的包管理器)。然后我们应该能够检查一些 go 运行时包的依赖关系，这真的很酷！想象一下编写一个自定义的舞会收集器，它会检查你所有的依赖项，检查新版本，如果找到，会给你一些过时的 pkgs，比如 `go_num_outdated_pkgs` 指标。
通过这种方式，如果您的服务非常过时，您可以编写警报。或者检查您的实时依赖哈希是否与当前哈希不匹配？

If you like the post, hit the up arrow button on the reddit and see you soon.

如果你喜欢这篇文章，请点击 reddit 上的向上箭头按钮，很快就会见到你。

 



