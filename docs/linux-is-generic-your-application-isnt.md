# Linux Is Generic; Your Application Isn’t

August 3, 2020

Imagine you’re a software reliability engineer with your service-level objectives clearly laid out on your Grafana dashboard. Suddenly, an alert appears: Your 99th percentile latency is going through the roof!

A quick glance at the relevant metrics reveals the culprit: A spike in traffic has exceeded the current system throughput. It’s clearly an issue caused by lack of capacity, that adding a few extra instances to the cluster should solve, right? After all, more machines means more capacity. You remain calm as the charts stabilize and the alert fades away. There’s no cause for concern; you automated the scaling process a long time ago.

But this story could have resulted in a very different, far worse scenario. Capacity is easily accessible these days, especially in cloud environments that offer simplified scaling processes. But the ease of provisioning and scaling also incurs rising infrastructure costs.

In order to truly overcome the tradeoff between capacity and costs, organizations must maximize application performance. Improving application performance would not only result in significant cost reduction but also a higher quality of service – serving customers at speed to improve customer experience and increase revenues.

Moreover, in many cases handling performance issues or improving performance can’t necessarily be achieved simply by adding more machines due to application bottlenecks or resource management inefficiencies. To paraphrase a wise uncle: with greater compute power (adding more machines) doesn’t come greater performance. Unfortunately, in most cases today, improving performance will require architecture changes or code refactoring.

# Optimizing the Software

Since increasing the node count rarely helps with improving the performance, let’s explore what can be accelerated. The stack is composed of hardware, an operating system (OS), libraries, and the application, among other components. Making improvements at the hardware level is not always feasible, especially when running in the cloud. Vertical scaling is usually limited to the available predefined instance types, which can inflate your cloud bill in the blink of an eye.

Rather than spending more on your cloud bill, consider checking how fast the userspace code is. Application developers have a lot on their plates already. They must ensure a proper domain model, maintainable architecture, and timely feature delivery, and so there is only so much they can optimize.

While it is possible to invest R&D efforts and time in replacing a poorly performing library with a more performant one or doing the occasional performance-focused rewrite in the hopes of resolving the issue, it’s possible neither will work.

When fidgeting with the hardware isn’t an option and the developers are unavailable to help, another option is to attempt to address the problem at the OS level.

# A Trip Down Memory Lane

The history of computing has not only been about smaller transistors and faster clocks. Back in the mainframe days, machines ran a single program at a time that was encoded on a punch card and inserted by a computer operator.

Then business people came along looking for a way to make computers more efficient since, during a program switch, the machines remained idle. This (and a fair amount of ingenuity on the part of early computer scientists) led to the creation of the operating system, a program that executed other programs and managed resource allocation between them.  Operating systems, and Linux as one, were designed for users behind the keyboard running simultaneous tasks, therefore the operating system resource management was designed to provide this illusion of parallelism to users behind the keyboard by optimizing internal resource management to achieve high interactivity and fairness.

The steadily decreasing price of servers enabled a vast range of potential allocation. Servers were still scarce and ran heterogeneous workloads ranging from web servers to long-running batch computations. Yet they all were running the same operating system more or less: Linux. The OS had to have sensible defaults so that it could perform well in diverse conditions and fit with different kinds of hardware from many different vendors.

Today, it’s not uncommon to command a fleet of generic virtual Linux boxes that are mainly focused on running a specific application, a microservice. Due to the inherent modular approach, microservices have known consistent resource usage characteristics and patterns. But the OS underneath hasn’t changed much. It still behaves as if it were supposed to execute multiple programs and share resources between them, which isn’t necessarily the most efficient in such a case and doesn’t provide optimal performance for the application.

# Optimizing the OS

There are a number of potential performance improvements that can be tested and applied directly at the OS level. Tuning sys controls can have a significant impact on the performance of many subcomponents, such as networking.

High speed NICs may require setting net.core.netdev\_max\_backlog much higher than the default to prevent filling up the card’s ring buffer, which can lead to packet loss. In addition, the initial value for net.core.somaxconn may prove far too low for proper machine saturation. And those are just two examples.

The I/O scheduler may be worth looking into as well. For example, with databases, the default CFQ (Completely Fair Queuing) can yield results that are inferior to those of the deadline scheduler. On the other hand, a noop scheduler allows you to avoid having to schedule I/O operations twice in cloud environments. After all, the VM’s hypervisor often manages the hardware already.

No matter what your approach, though, constant meticulous measurement in a production-like environment (or even the production itself, if you’re into chaos engineering) is recommended. Performance tuning is a highly advanced field that requires specialistic system knowledge. It’s also easy to make silly mistakes, such as forgetting to reload kernel parameter preferences after tweaking them.

# Advanced OS Tweaks

Some OS features are not so easily accessible for performance tuning. For example, the Linux process scheduler utilizes the Completely Fair Scheduler (CFS), which is perfectly sensible in _most_ cases. However, it can sometimes create a significant [performance gap that cannot be easily found](https://blog.acolyer.org/2016/04/26/the-linux-scheduler-a-decade-of-wasted-cores/) with standard profiling tools. And even if you do discover it, you can’t simply change a parameter in one of the configuration files; rather a kernel patch and a rebuild is required.

Let’s say you’re perfectly fine with the algorithm, but you’d simply like to state that some threads are more important than others. By default you can’t do this, as the niceness setting only works at the process level.

I/O-bound applications are also complicated. Even when using raw sockets and epoll, there is no runtime mechanism to provide selection logic or even a priority to the sockets in the queue. And there is no way for a kernel to know the performance budget for a request.

# Additional Means for Performance Optimization

In an ideal world, your application will consist of purpose-built operating systems tailored for each microservice and exploiting every opportunity to boost performance and synergize with the application. In such an ideal world, the internal resource management mechanisms within the operating system will be tailored to the application-specific utility function to drive optimized performance and in turn, also deliver reduced infrastructure costs.

Unfortunately, we’re not there yet, and such solutions are currently only available to corporate giants who can afford to hire a few dozen people to do just that full time. In the words of William Gibson, “The future is already here – it’s just not evenly distributed.”

So what is left to those with finite budgets? A new approach for real-time continuous optimization that enables organizations to leverage AI-driven infrastructure optimizations that are suited specifically to the running workload.

Using application-driven scheduling and prioritization algorithms, it is possible to identify contended resources, bottlenecks, and prioritization opportunities and solve them in real-time.

These innovative solutions leverage application’s specific resource usage patterns, the data flow, analyzing CPU scheduling order, oversubscribed locks, memory, network, and disk access patterns, and more.

This approach ensures the most efficient use of compute resources, resulting in the need for fewer VMs, less compute resources, reducing costs significantly while delivering better performance.
