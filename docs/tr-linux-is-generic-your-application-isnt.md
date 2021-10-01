# Linux Is Generic; Your Application Isn’t

# Linux 是通用的；您的应用程序不是

August 3, 2020

Imagine you’re a software reliability engineer with your service-level objectives clearly laid out on your Grafana dashboard. Suddenly, an alert appears: Your 99th percentile latency is going through the roof!

想象一下，您是一名软件可靠性工程师，您的服务级别目标清楚地列在您的 Grafana 仪表板上。突然，一条警报出现：您的第 99 个百分位数的延迟正在飙升！

A quick glance at the relevant metrics reveals the culprit: A spike in traffic has exceeded the current system throughput. It’s clearly an issue caused by lack of capacity, that adding a few extra instances to the cluster should solve, right? After all, more machines means more capacity. You remain calm as the charts stabilize and the alert fades away. There’s no cause for concern; you automated the scaling process a long time ago.

快速浏览一下相关指标就会发现罪魁祸首：流量高峰超过了当前的系统吞吐量。显然是容量不足导致的问题，向集群添加一些额外的实例应该可以解决，对吗？毕竟，更多的机器意味着更多的容量。当图表趋于稳定且警报逐渐消失时，您会保持冷静。无需担心；您很久以前就自动化了扩展过程。

But this story could have resulted in a very different, far worse scenario. Capacity is easily accessible these days, especially in cloud environments that offer simplified scaling processes. But the ease of provisioning and scaling also incurs rising infrastructure costs.

但这个故事可能会导致一个非常不同的、更糟糕的情况。如今，容量可以轻松访问，尤其是在提供简化扩展过程的云环境中。但易于配置和扩展也会导致基础设施成本上升。

In order to truly overcome the tradeoff between capacity and costs, organizations must maximize application performance. Improving application performance would not only result in significant cost reduction but also a higher quality of service – serving customers at speed to improve customer experience and increase revenues.

为了真正克服容量和成本之间的权衡，组织必须最大限度地提高应用程序性能。提高应用程序性能不仅会显着降低成本，还会带来更高的服务质量——快速服务客户以改善客户体验并增加收入。

Moreover, in many cases handling performance issues or improving performance can’t necessarily be achieved simply by adding more machines due to application bottlenecks or resource management inefficiencies. To paraphrase a wise uncle: with greater compute power (adding more machines) doesn’t come greater performance. Unfortunately, in most cases today, improving performance will require architecture changes or code refactoring.

此外，在许多情况下，由于应用程序瓶颈或资源管理效率低下，仅通过添加更多机器不一定能解决性能问题或提高性能。套用一位睿智的叔叔的话：更高的计算能力（添加更多机器）并不会带来更高的性能。不幸的是，在今天的大多数情况下，提高性能将需要架构更改或代码重构。

# Optimizing the Software

# 优化软件

Since increasing the node count rarely helps with improving the performance, let’s explore what can be accelerated. The stack is composed of hardware, an operating system (OS), libraries, and the application, among other components. Making improvements at the hardware level is not always feasible, especially when running in the cloud. Vertical scaling is usually limited to the available predefined instance types, which can inflate your cloud bill in the blink of an eye.

由于增加节点数量很少有助于提高性能，让我们探索可以加速的内容。该堆栈由硬件、操作系统 (OS)、库和应用程序以及其他组件组成。在硬件级别进行改进并不总是可行的，尤其是在云中运行时。垂直扩展通常仅限于可用的预定义实例类型，这可能会在眨眼间增加您的云账单。

Rather than spending more on your cloud bill, consider checking how fast the userspace code is. Application developers have a lot on their plates already. They must ensure a proper domain model, maintainable architecture, and timely feature delivery, and so there is only so much they can optimize.

与其在云账单上花费更多，不如考虑检查用户空间代码的速度。应用程序开发人员已经有很多事情要做。他们必须确保正确的领域模型、可维护的架构和及时的功能交付，因此他们只能优化这么多。

While it is possible to invest R&D efforts and time in replacing a poorly performing library with a more performant one or doing the occasional performance-focused rewrite in the hopes of resolving the issue, it’s possible neither will work.

虽然有可能投入研发工作和时间来用性能更好的库替换性能不佳的库，或者偶尔进行以性能为中心的重写以希望解决问题，但这两种方法都行不通。

When fidgeting with the hardware isn’t an option and the developers are unavailable to help, another option is to attempt to address the problem at the OS level.

当无法解决硬件问题并且开发人员无法提供帮助时，另一种选择是尝试在操作系统级别解决问题。

# A Trip Down Memory Lane

# 记忆之路之旅

The history of computing has not only been about smaller transistors and faster clocks. Back in the mainframe days, machines ran a single program at a time that was encoded on a punch card and inserted by a computer operator. 

计算的历史不仅仅是关于更小的晶体管和更快的时钟。回到大型机时代，机器一次运行一个程序，该程序在穿孔卡上编码并由计算机操作员插入。

Then business people came along looking for a way to make computers more efficient since, during a program switch, the machines remained idle. This (and a fair amount of ingenuity on the part of early computer scientists) led to the creation of the operating system, a program that executed other programs and managed resource allocation between them. Operating systems, and Linux as one, were designed for users behind the keyboard running simultaneous tasks, therefore the operating system resource management was designed to provide this illusion of parallelism to users behind the keyboard by optimizing internal resource management to achieve high interactivity and fairness.

然后商界人士开始寻找提高计算机效率的方法，因为在程序切换期间，计算机一直处于闲置状态。这（以及早期计算机科学家的相当多的独创性）导致了操作系统的创建，这是一个执行其他程序并管理它们之间资源分配的程序。操作系统和 Linux 是为同时运行任务的键盘后面的用户设计的，因此操作系统资源管理旨在通过优化内部资源管理来实现高交互性和公平性，为键盘后面的用户提供这种并行的错觉。

The steadily decreasing price of servers enabled a vast range of potential allocation. Servers were still scarce and ran heterogeneous workloads ranging from web servers to long-running batch computations. Yet they all were running the same operating system more or less: Linux. The OS had to have sensible defaults so that it could perform well in diverse conditions and fit with different kinds of hardware from many different vendors.

服务器价格的稳步下降实现了广泛的潜在分配。服务器仍然稀缺，运行从 Web 服务器到长时间运行的批量计算的异构工作负载。然而，他们或多或少都在运行相同的操作系统：Linux。操作系统必须有合理的默认设置，以便它可以在不同的条件下良好运行，并适合来自许多不同供应商的不同类型的硬件。

Today, it’s not uncommon to command a fleet of generic virtual Linux boxes that are mainly focused on running a specific application, a microservice. Due to the inherent modular approach, microservices have known consistent resource usage characteristics and patterns. But the OS underneath hasn’t changed much. It still behaves as if it were supposed to execute multiple programs and share resources between them, which isn’t necessarily the most efficient in such a case and doesn’t provide optimal performance for the application.

今天，指挥一组主要专注于运行特定应用程序（微服务）的通用虚拟 Linux 机器的情况并不少见。由于固有的模块化方法，微服务具有一致的资源使用特征和模式。但是下面的操作系统并没有太大变化。它仍然表现得好像它应该执行多个程序并在它们之间共享资源，这在这种情况下不一定是最有效的，并且不会为应用程序提供最佳性能。

# Optimizing the OS

# 优化操作系统

There are a number of potential performance improvements that can be tested and applied directly at the OS level. Tuning sys controls can have a significant impact on the performance of many subcomponents, such as networking.

有许多潜在的性能改进可以直接在操作系统级别进行测试和应用。调整系统控制会对许多子组件的性能产生重大影响，例如网络。

High speed NICs may require setting net.core.netdev\_max\_backlog much higher than the default to prevent filling up the card’s ring buffer, which can lead to packet loss. In addition, the initial value for net.core.somaxconn may prove far too low for proper machine saturation. And those are just two examples.

高速 NIC 可能需要将 net.core.netdev\_max\_backlog 设置得比默认值高得多，以防止填满卡的环形缓冲区，这会导致数据包丢失。此外，net.core.somaxconn 的初始值可能证明对于适当的机器饱和来说太低了。这些只是两个例子。

The I/O scheduler may be worth looking into as well. For example, with databases, the default CFQ (Completely Fair Queuing) can yield results that are inferior to those of the deadline scheduler. On the other hand, a noop scheduler allows you to avoid having to schedule I/O operations twice in cloud environments. After all, the VM’s hypervisor often manages the hardware already.

I/O 调度器也可能值得研究。例如，对于数据库，默认的 CFQ（完全公平队列）可能会产生比截止时间调度程序差的结果。另一方面，noop 调度程序允许您避免在云环境中两次调度 I/O 操作。毕竟，VM 的管理程序通常已经在管理硬件。

No matter what your approach, though, constant meticulous measurement in a production-like environment (or even the production itself, if you’re into chaos engineering) is recommended. Performance tuning is a highly advanced field that requires specialistic system knowledge. It’s also easy to make silly mistakes, such as forgetting to reload kernel parameter preferences after tweaking them.

但是，无论您采用何种方法，都建议在类似生产的环境（甚至是生产本身，如果您从事混沌工程）中进行持续细致的测量。性能调优是一个非常先进的领域，需要专业的系统知识。也很容易犯一些愚蠢的错误，例如在调整后忘记重新加载内核参数首选项。

# Advanced OS Tweaks

# 高级操作系统调整

Some OS features are not so easily accessible for performance tuning. For example, the Linux process scheduler utilizes the Completely Fair Scheduler (CFS), which is perfectly sensible in _most_ cases. However, it can sometimes create a significant [performance gap that cannot be easily found](https://blog.acolyer.org/2016/04/26/the-linux-scheduler-a-decade-of-wasted-cores/) with standard profiling tools. And even if you do discover it, you can’t simply change a parameter in one of the configuration files; rather a kernel patch and a rebuild is required.

某些操作系统功能不太容易用于性能调优。例如，Linux 进程调度程序使用完全公平调度程序 (CFS)，这在_大多数_情况下是非常明智的。然而，它有时会造成显着的[不易发现的性能差距](https://blog.acolyer.org/2016/04/26/the-linux-scheduler-a-decade-of-wasted-cores/) 使用标准分析工具。即使您确实发现了它，您也不能简单地更改其中一个配置文件中的参数；而是需要内核补丁和重建。

Let’s say you’re perfectly fine with the algorithm, but you’d simply like to state that some threads are more important than others. By default you can’t do this, as the niceness setting only works at the process level.

假设您对算法完全满意，但您只想说明某些线程比其他线程更重要。默认情况下您不能这样做，因为 niceness 设置仅适用于流程级别。

I/O-bound applications are also complicated. Even when using raw sockets and epoll, there is no runtime mechanism to provide selection logic or even a priority to the sockets in the queue. And there is no way for a kernel to know the performance budget for a request.

I/O 绑定应用程序也很复杂。即使使用原始套接字和 epoll，也没有运行时机制来为队列中的套接字提供选择逻辑甚至优先级。内核无法知道请求的性能预算。

# Additional Means for Performance Optimization 

# 额外的性能优化手段

In an ideal world, your application will consist of purpose-built operating systems tailored for each microservice and exploiting every opportunity to boost performance and synergize with the application. In such an ideal world, the internal resource management mechanisms within the operating system will be tailored to the application-specific utility function to drive optimized performance and in turn, also deliver reduced infrastructure costs.

在理想的世界中，您的应用程序将由为每个微服务量身定制的专用操作系统组成，并利用每一个机会来提高性能并与应用程序协同工作。在这样的理想世界中，操作系统内的内部资源管理机制将针对特定于应用程序的实用程序功能进行定制，以推动优化性能，进而降低基础设施成本。

Unfortunately, we’re not there yet, and such solutions are currently only available to corporate giants who can afford to hire a few dozen people to do just that full time. In the words of William Gibson, “The future is already here – it’s just not evenly distributed.”

不幸的是，我们还没有做到这一点，此类解决方案目前仅适用于有能力雇佣几十人来做全职工作的企业巨头。用威廉吉布森的话来说，“未来已经到来——只是分布不均。”

So what is left to those with finite budgets? A new approach for real-time continuous optimization that enables organizations to leverage AI-driven infrastructure optimizations that are suited specifically to the running workload.

那么，那些预算有限的人还剩下什么呢？一种实时持续优化的新方法，使组织能够利用专门适用于正在运行的工作负载的 AI 驱动的基础架构优化。

Using application-driven scheduling and prioritization algorithms, it is possible to identify contended resources, bottlenecks, and prioritization opportunities and solve them in real-time.

使用应用程序驱动的调度和优先级算法，可以识别竞争资源、瓶颈和优先级机会并实时解决它们。

These innovative solutions leverage application’s specific resource usage patterns, the data flow, analyzing CPU scheduling order, oversubscribed locks, memory, network, and disk access patterns, and more.

这些创新解决方案利用应用程序的特定资源使用模式、数据流、分析 CPU 调度顺序、超额订阅锁、内存、网络和磁盘访问模式等。

This approach ensures the most efficient use of compute resources, resulting in the need for fewer VMs, less compute resources, reducing costs significantly while delivering better performance. 

这种方法可确保最有效地使用计算资源，从而减少对虚拟机和计算资源的需求，显着降低成本，同时提供更好的性能。

