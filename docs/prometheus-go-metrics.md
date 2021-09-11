# Exploring Prometheus Go client metrics - Povilas Versockas

Published April 12, 2018

In this post I want to explore go metrics, which are exported by `client_golang` via `promhttp.Handler()` call. These metrics help you better understand how Go works & gives you better operational sanity, when your on call. 

Interested in learning more Prometheus? Check out [Monitoring Systems and Services with Prometheus](https://povilasv.me/out/prometheus), it’s awesome course that will get you up to speed.

Let’s start with a simple program, registering prom handler and listening on `8080` port:

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

- [Process Collector](https://godoc.org/github.com/prometheus/client_golang/prometheus#NewProcessCollector) – which collects basic Linux process information like CPU, memory, file descriptor usage and start time.
- [Go Collector](https://godoc.org/github.com/prometheus/client_golang/prometheus#NewGoCollector) – which collects information about Go’s runtime like details about GC, number of gouroutines and OS threads.

## Process Collector

What this collector does is reads *proc* file system. *proc* file system exposes internal kernel data structures, which is used to obtain information about the system.[1](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-1)

So Prometheus client reads `/proc/PID/stat` file, which looks like this:

```
1 (sh) S 0 1 1 34816 8 4194560 674 43 9 1 5 0 0 0 20 0 1 0 89724 1581056 209 18446744073709551615 94672542621696 94672543427732 140730737801568 0 0 0 0 2637828 65538 1 0 0 17 3 0 0 0 0 0 94672545527192 94672545542787 94672557428736 140730737807231 140730737807234 140730737807234 140730737807344 0
```

You can get human readable variant of this information using `cat /proc/PID/status`.

**process_cpu_seconds_total** – it uses **utime** – number of ticks executing code in user mode, measured in jiffies, with **stime** – jiffies spent in the system mode, executing code on behalf of the process (like doing system calls). A *[jiffy](https://elinux.org/Kernel_Timer_Systems)* is the time between two ticks of the system timer interrupt. [2](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-2)

**process_cpu_seconds_total** equals to sum of **utime** and **stime** and divide by USER_HZ. This makes sense, as dividing **number of scheduler ticks** by **Hz(ticks per second)** produces total time in seconds operating system has been running the process. [3](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-3)

**process_virtual_memory_bytes** – uses **vsize** – virtual memory size is the amount of address space that a process is  managing. This includes all types of memory, both in RAM and swapped  out.

**process_resident_memory_bytes** – multiplies **rss** – resident set memory size is number of memory pages the process has in real memory, with pagesize [4](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-4). This results in the amount of memory that belongs specifically to that  process in bytes. This excludes swapped out memory pages.

**process_start_time_seconds** – uses **start_time** – time the process started after system boot, which is expressed in jiffies and **btime** from `/proc/stat` which shows the time at which the system booted in seconds since the Unix epoch. **start_time** is divided by *USER_HZ* in order to get the value in seconds.

**process_open_fds** – counts the number of files in `/proc/PID/fd` directory. This shows currently open regular files, sockets, pseudo terminals, etc.

**process_max_fds** – reads /proc/{PID}/limits and uses Soft Limit from “Max Open Files” row. The interesting bit here is that `/limits` lists Soft and Hard limits.
As it turns out, the soft limit is the value that the kernel enforces for  the corresponding resource and the Hard limit acts as a ceiling for the  soft limit.
An unprivileged process may only set its soft limit to a value up to the hard limit and (irreversibly) lower its Hard limit. [5](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-5)

In Go you can use `err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{Cur: 9, Max: 10})` to set your limit.

## Go Collector

Go Collector’s most of the metrics are taken from `runtime`, `runtime/debug` packages.

**go_goroutines** – calls out to `runtime.NumGoroutine()`, which computes the value based off the scheduler struct and global `allglen` variable. As all the values in sched struct can be changed concurently  there is this funny check where if computed value is less than 1 it  becomes 1.

**go_threads** – calls out to `runtime.CreateThreadProfile()`, which reads off global `allm` variable. If you don’t know what M’s or G’s you can read my [blogpost about it](https://povilasv.me/go-scheduler/).

**go_gc_duration_seconds** – calls out to `debug.ReadGCStats()` with `PauseQuantile` set to 5, which returns us the minimum, 25%, 50%, 75%, and maximum  pause times. Then it manualy creates a Summary type from pause  quantiles, `NumGC` var and `PauseTotal` seconds. It’s cool how well `GCStats` struct fits the prom’s Summary type. [6](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-6)

**go_info** – this provides us with Go version. It’s pretty clever, it calls out to `runtime.Version()` and set’s that as a `version` label and then always returns value of `1` for this gauge metric.

### Memory

Go Collector provides us with a lot of metrics about memory and GC.
All those metrics are from `runtime.ReadMemStats()`, which gives us metrics from [MemStats](https://golang.org/pkg/runtime/#MemStats) struct.
One thing that worries me, is that `runtime.ReadMemStats()` has a explicit call to make a stop-the-world pause [7](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-7).
So I wonder how much actual cost this pause introduces?
As during stop-the-world pause, all goroutines are paused, so that GC can  run. I’ll probably do a comparison of an app with and without  instrumentation in a later post.

We already seen that Linux provides us with rss/vsize metrics for  memory stats, so naturally the question arises which metrics to use, the ones provided in MemStats or rss/vsize?

The good part about resident set size and virtual memory size is that it’s based off Linux primitives and is programming language agnostic.
So in theory you could instrument any program and you would know how much  memory it consumes (as long as you name your metrics consistently, ie *process_virtual_memory_bytes* and *process_resident_memory_bytes*.).
In practice, however when Go process starts up it takes a lot of virtual  memory beforehand, such a simple program like the one above takes up to  544MiB of vsize on my machine (x86_64 Ubuntu), which is a bit confusing. RSS shows around 7mib, which is closer to the actual usage.

On the other hand using Go runtime based metrics gives more fined  grained information on what is happening in your running application.
You should be able to find out more easily whether your program has a memory leak, how long GC took, how much it reclaimed.
Also, it should point you into right direction when you are optimizing program’s memory allocations.

I haven’t looked in detail how Go GC and memory model works, a part of it’s concurrency model [8](about:reader?url=https%3A%2F%2Fpovilasv.me%2Fprometheus-go-metrics%2F%23#fn-1692-8), so this bit is still new to me.

So let’s take a look at those metrics:

**go_memstats_alloc_bytes** – a metric which shows how much bytes of memory is allocated on the [Heap](https://en.wikipedia.org/wiki/Memory_management#HEAP) for the Objects. The value is same as **go_memstats_heap_alloc_bytes**. This metric counts all reachable heap objects plus unreachable objects, GC has not yet freed.

**go_memstats_alloc_bytes_total** – this metric  increases as objects are allocated in the Heap, but doesn’t decrease  when they are freed. I think it is immensly useful, as it is only  increasing number and has same nice properties that [Prometheus Counter](https://povilasv.me/prometheus-tracking-request-duration/) has. Doing `rate()` on it should show us how many bytes/s of memory app consumes and is “durable” across restarts and scrape misses.

**go_memstats_sys_bytes** – it’s a metric, which measures how many bytes of memory in total is taken from system by Go. It sums all the **sys** metrics described below.

**go_memstats_lookups_total** – counts how many pointer dereferences happened. This is a counter value so you can use `rate()` to lookups/s.

**go_memstats_mallocs_total** – shows how many heap objects are allocated. This is a counter value so you can use `rate()` to objects allocated/s.

**go_memstats_frees_total** – shows how many heap objects are freed. This is a counter value so you can use `rate()` to objects allocated/s. Note you can get number of live objects with `go_memstats_mallocs_total` – `go_memstats_frees_total`.

Turns out that Go organizes memory in spans, which are contiguous regions of memory of 8K or larger. There are 3 types of Spans:
\1) **idle** – span, that has no objects and can be released back to the OS, or reused for heap allocation, or reused for stack memory.
\2) **in use** – span, that has atleast one heap object and may have space for more.
\3) **stack** – span, which is used for goroutine stack. This span can live either in stack or in heap, but not in both.

### Heap memory metrics

**go_memstats_heap_alloc_bytes** – same as **go_memstats_alloc_bytes**.

**go_memstats_heap_sys_bytes** – bytes of memory obtained for the heap from OS. This includes [virtual address space](https://en.wikipedia.org/wiki/Virtual_address_space) that has been resevered, but not yet used and virtual address space  which was returned to OS after it became unused. This metric estimates  the largest size the heap.

**go_memstats_heap_idle_bytes** – shows how many bytes are in **idle** spans.

**go_memstats_heap_idle_bytes** minus **go_memstats_heap_released_bytes** estimates how many bytes of memory could be released, but is kept by  runtime, so that runtime can allocate objects on the heap without asking OS for more memory.

**go_memstats_heap_inuse_bytes** – shows how many bytes in **in-use** spans.

**go_memstats_heap_inuse_bytes** minus **go_memstats_heap_alloc_bytes** shows how many bytes of memory has been allocated for the heap, but not currently used.

**go_memstats_heap_released_bytes** – shows how many bytes of idle spans were returned to the OS.

**go_memstats_heap_objects** – shows how many objects are allocated on the heap. This changes as GC is performed and new objects are allocated.

### Stack memory metrics

**go_memstats_stack_inuse_bytes** – shows how many bytes of memory is used by stack memory spans, which have atleast one object  in them. Go doc says, that stack memory spans can only be used for other stack spans, i.e. there is no mixing of heap objects and stack objects  in one memory span.

**go_memstats_stack_sys_bytes** – shows how many bytes of stack memory is obtained from OS. It’s **go_memstats_stack_inuse_bytes** plus any memory obtained for OS thread stack.

There is no **go_memstats_stack_idle_bytes**, as unused stack spans are counted towards **go_memstats_heap_idle_bytes**.

### Off-heap memory metrics

These metrics are bytes allocated for runtime internal structures,  that are not allocated on the heap, because they implement the heap.

**go_memstats_mspan_inuse_bytes** – shows how many bytes are in use by `mspan` structures.

**go_memstats_mspan_sys_bytes** – shows how many bytes are obtained from OS for `mspan` structures.

**go_memstats_mcache_inuse_bytes** – shows how many bytes are in use by `mcache` structures.

**go_memstats_mcache_sys_bytes** – shows how many bytes are obtained from OS for `mcache` structures.

**go_memstats_buck_hash_sys_bytes** – shows how many bytes of memory are in bucket hash tables, which are used for profiling.

**go_memstats_gc_sys_bytes** – shows how many in garbage collection metadata.

**go_memstats_other_sys_bytes** – **go_memstats_other_sys_bytes** shows how many bytes of memory are used for other runtime allocations.

**go_memstats_next_gc_bytes** – shows the target heap size of the next GC cycle. GC’s goal is to keep **go_memstats_heap_alloc_bytes** less than this value.

**go_memstats_last_gc_time_seconds** – contains unix timestamp when last GC finished.

**go_memstats_last_gc_cpu_fraction** – shows the fraction of this program’s available CPU time used by GC since the program started.
This metric is also provided in GODEBUG=gctrace=1.

## Playing around with numbers

So it’s a lot of metrics and a lot of information.
I think the best way to learn is to just play around with it, so in this part I’ll do just that.
So I’ll be using the same program that is above.
Here are the dump from `/metrics` (edited for space), which I’m going to use:

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

mspan_inuse_bytes = 25840 = 25.2 KiB

mspan_sys_bytes = 32768 = 32 KiB

mcache_inuse_bytes = 6912 = 6.8 KiB

mcache_sys_bytes = 16384 = 12 KiB

buck_hash_sys_bytes = 1.443899e+06 = 1443899 = 1410 KiB = 1.4 MiB

gc_sys_bytes = 2.371584e+06 = 2371584 = 2316 KiB = 2.3 MiB

other_sys_bytes = 1.310909e+06 = 1310909 = 1280,2 KiB = 1.3MiB

stack_inuse_bytes = 425984 = 416 KiB

stack_sys_bytes = 425984 = 416 KiB

sys_bytes = 7.2284408e+07 = 72284408 = 70590.2 KiB = 68.9 MiB

next_gc_bytes = 4.194304e+06 = 4194304 = 4096 KiB = 4 MiB

gc_cpu_fraction = 1.421928536233557e-06 = 0.000001

last_gc_time_seconds = 1.5235057190167596e+09 = Thu, 12 Apr 2018 05:47:59 GMT

Interesting bit is that **heap_inuse_bytes** is more than **heap_alloc_bytes**.
I think **heap_alloc_bytes** shows how many bytes are in terms of **objects** and **heap_inuse_bytes** shows bytes of memory in terms of **spans**.
Dividing **heap_inuse_bytes** by size of the span gives: 3039232 / 8192 = 371 span.

**heap_inuse_bytes** minus **heap_alloc_bytes**, should show the amount of free space that we have in in-use spans, which is 2,9 MiB – 2.1 MiB = 0.8 MiB.
This roughly means that we can allocate 0.8 MiB of objects on the heap without using new memory spans.
But we should keep in mind memory fragmentation.
Imagine if you have a new bytes slice of 10K bytes, the memory could be in the  position, where it doesn’t have a contiguous block of 10K bytes + slice  header, so it would need use a new span, instead of reusing

**heap_idle_bytes** minus **heap_released_byte** shows that we have around 60.6 MiB of unused spans, which are reserved  from OS and could be returned to OS. It’s 63643648/8192 = 7769 spans.

**heap_sys_bytes**, which is 63.6MiB estimates the largest size the heap has had. It’s 66682880/8192 = 8140 spans.

**mallocs_total** shows that we allocated 18707 objects  and freed 12209. So currently, we have 18707-12209 = 6498 objects. We  can find the average size of the object dividing *heap_alloc_bytes* over live objects, which is 6498. The result is 2243440 / 6498 = 345.3 bytes.
(This is probably a stupid metric, as object vary a lot in size and we should do histograms instead)

So **sys_bytes** should be a sum of all ***sys** metrics. So let’s check that:
sys_bytes == mspan_sys_bytes + mcache_sys_bytes + buck_hash_sys_bytes +  gc_sys_bytes + other_sys_bytes + stack_sys_bytes + heap_sys_bytes.
So, we have 72284408 == 32768 + 16384 + 1443899 + 2371584 + 1310909 +  425984 + 66682880, which is 72284408 == 72284408, which is correct.

The interesting detail about **sys_bytes**, is that it’s 68,9 MiB it’s how many bytes of memory in total taken from OS. Meanwhile, OS’es **vsize** gives you 616,7MiB and rss 10.4 MiB. So all these numbers don’t really match up.

As I understand it so part of our memory could be in OS’s memory  pages which are in swap or in filesystem (not in RAM), so this would  explain why rss is smaller than **sys_bytes**.

And **vsize** contains a lot of things, like mapped libc, pthreads libs, etc. You can explore `/proc/PID/maps` and `/proc/PID/smaps` file, to see what is being currently mapped.

**gc_cpu_fraction** is running crazy low, `0.000001` of CPU time is used for GC. That’s really really cool. (Although this program doesn’t produce much garbage)

**next_gc_bytes** shows that the target for GC is to keep **heap_alloc_bytes** under 4 MiB, as **heap_alloc_bytes** is currently at 2.1 MiB the target is achieved.

## Conclusion

I love Go and the fact that it exposes so much useful information in  it’s packages and users like you and me can just call a function and get the information. Alos Prometheus is really a great tool to monitor  applications.

Do you want get better at Prometheus? Check out [Monitoring Systems and Services with Prometheus](https://povilasv.me/out/prometheus). I definitely recommend the module.

It was really cool playing around and reading about Linux & Go,  so I’m thinking of doing part 2 of this post. Maybe go look into metrics provided by [cAdvisor](https://github.com/google/cadvisor) or show how to use some of the metrics described here in dashboards/alerts with Prometheus.

Also, once [vgo](https://research.swtch.com/vgo) get’s  integrated (and I really really hope it does, cause it’s like the best  package manager I ever used). Then we should be able to inspect  dependencies from some go runtime package, which would be really cool!  Imagine writing a custom prom collector, which would go through all your dependencies, check for new versions and if found wouldgive you back a  number of outdated pkgs, something like `go_num_outdated_pkgs` metric.
This way you could write an alert if your service get’s terribly outdated.  Or check that your live dependency hashes don’t match current hashes?

If you like the post, hit the up arrow button on the reddit and see you soon.

​          
