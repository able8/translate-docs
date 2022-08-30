# Monitoring and Observability With USE and RED

# 使用 USE 和 RED 进行监控和可观察性

https://orangematter.solarwinds.com/2017/10/05/monitoring-and-observability-with-use-and-red/

October 5, 2017

Modern systems can emit thousands or millions of metrics, and modern monitoring tools can collect them all. Faced with such an abundance of data, it can be difficult to know where to start looking when trying to diagnose a problem. And when you’re not in diagnosis mode, but you just want to know whether there’s a problem at all, you might have the same difficulty. What are the most important KPIs coming from your systems? I've written extensively about this before, but this time I want to refer to acronyms—USE and RED respectively—that are easy to remember and provide good high-level guidance for system observability .

现代系统可以发出数千或数百万个指标，现代监控工具可以收集所有指标。面对如此丰富的数据，在尝试诊断问题时可能很难知道从哪里开始寻找。而当你不处于诊断模式，但你只是想知道是否有问题时，你可能会遇到同样的困难。来自您的系统的最重要的 KPI 是什么？我之前已经写过很多关于这方面的文章，但这次我想分别引用首字母缩略词——USE 和 RED——它们很容易记住并为系统可观察性提供良好的高级指导.

## The USE Method

## 使用方法

USE is an acronym for [Utilization, Saturation, and Errors](http://www.brendangregg.com/usemethod.html). Brendan Gregg suggests using it to get started quickly when you’re diving into an unknown system: “I developed the USE method to teach others how to solve common performance issues quickly, without overlooking important areas. Like an emergency checklist in a flight manual, it is intended to be simple, straightforward, complete, and fast. ”A summary of USE is “For every resource, check utilization, saturation, and errors.” What do those things mean? Brendan defines the terminology:

USE 是 [Utilization, Saturation, and Errors](http://www.brendangregg.com/usemethod.html) 的首字母缩写词。 Brendan Gregg 建议在您潜入未知系统时使用它来快速入门：“我开发了 USE 方法来教其他人如何快速解决常见的性能问题，而不会忽略重要的领域。就像飞行手册中的紧急检查表一样，它旨在简单、直接、完整和快速。” USE 的总结是“对于每个资源，检查利用率、饱和度和错误”。那些东西是什么意思？ Brendan 定义了术语：

- Utilization: the average time the resource was busy servicing work
- Saturation: the degree to which the resource has extra work which it can't service, often queued
- Errors: the count of error events

- 利用率：资源忙于服务工作的平均时间
- 饱和度：资源有无法服务的额外工作的程度，通常是排队的
- 错误：错误事件的计数

This disambiguates utilization and saturation, making it clear utilization is “busy time %” and saturation is “backlog.” These terms are very different from things a person might confuse with them, such as “disk utilization” as an expression of how much disk space is left.

这消除了利用率和饱和度的歧义，明确了利用率是“忙碌时间百分比”，而饱和度是“积压”。这些术语与人们可能会混淆的东西非常不同，例如“磁盘利用率”作为剩余磁盘空间的表达。

## The RED Method

## RED 方法

I first saw this acronym in a [talk on monitoring microservices](https://www.slideshare.net/weaveworks/monitoring-microservices) in 2015. The acronym stands for Rate, Errors, and Duration. These are request-scoped, not resource-scoped as the USE method is. Duration is explicitly taken to mean _distributions_, not averages.

我在 2015 年的 [关于监控微服务的谈话](https://www.slideshare.net/weaveworks/monitoring-microservices) 中第一次看到这个首字母缩写词。首字母缩写词代表速率、错误和持续时间。这些是请求范围的，而不是资源范围的 USE 方法。持续时间明确表示_distributions_，而不是平均值。

## USE and RED: Two Sides of the Same Coin 

## USE 和 RED：同一枚硬币的两个面

What may not be obvious is USE and RED complement one another. The USE method is an internal, service-centric view. The system or service’s workload is assumed, and USE directs attention to the _resources_ handling the workload. The goal is to understand how these resources behave in the presence of the load. The RED method, on the other hand, is about the _workload_ itself, and treats the service as a black box. It’s an externally-visible view of the behavior of the workload as serviced by the resources. I define _workload_ as a population of requests over a period of time. I've spoken and written extensively before about the importance of measuring the workload, since the system's _raison d'etre_ is to do useful work. Taken together, RED and USE comprise minimally complete, maximally useful observability—a way to understand _both_aspects of a system:  its users/customers and the work they request, and its resources/components and how they react to the workload. (I include users in the system. Users aren't separate from the system; they're an inextricable part of it.) I often refer to this duality as the "Zen of Performance," a holistic, unified system performance worldview I' m developing. It's work in progress!

可能不明显的是 USE 和 RED 相辅相成。 USE 方法是一种内部的、以服务为中心的视图。假设系统或服务的工作负载，USE 将注意力集中在处理工作负载的资源上。目标是了解这些资源在负载存在时的行为方式。另一方面，RED 方法是关于_workload_ 本身的，并将服务视为一个黑盒子。这是资源服务的工作负载行为的外部可见视图。我将_workload_定义为一段时间内的一组请求。我之前已经就测量工作负载的重要性发表过大量的演讲和文章，因为系统的_存在理由_是做有用的工作。总之，RED 和 USE 构成了最小完整、最大有用的可观察性——一种理解系统：它的用户/客户和他们请求的工作，它的资源/组件以及他们对工作负载的反应。 （我将用户包括在系统中。用户不是与系统分离的；他们是系统不可分割的一部分。）我经常将这种二元性称为“性能之禅”，这是一种整体的、统一的系统性能世界观。发展中。正在进行中！

## Mapping USE and RED to Standard Terminology

## 将 USE 和 RED 映射到标准术语

USE and RED are convenient, and part of the reason they’re so valuable is their atoms map directly to standard concepts that are core performance metrics:

USE 和 RED 很方便，它们如此有价值的部分原因是它们的原子直接映射到作为核心性能指标的标准概念：

- U = Utilization, as canonically defined
- S = Concurrency
- E = Error Rate, as a throughput metric
- R = Request Throughput, in requests per second
- E = Request Error Rate, as either a throughput metric or a fraction of overall throughput
- D = Latency, Residence Time, or Response Time; all three are widely used



- U = 利用率，如规范定义的那样
- S = 并发
- E = 错误率，作为吞吐量指标
- R = 请求吞吐量，以每秒请求数为单位
- E = 请求错误率，作为吞吐量指标或总吞吐量的一部分
- D = 延迟、停留时间或响应时间；这三个都被广泛使用

To learn more about why these metrics are so fundamental to performance and observability, listen to Jon Moore's talk on why [API admission control](https://www.youtube.com/watch?v=m64SWl9bfvk) should use concurrency instead of throughput . And, for further reading, consider my eBooks on [queuing theory](https://www.vividcortex.com/resources/queueing-theory) and the [Universal Scalability Law](https://www.vividcortex.com/resources/universal-scalability-law/). In conclusion, if you’re unsure which metrics are most useful for both monitoring and diagnosis, USE and RED are great places to start. 

要详细了解为什么这些指标对性能和可观察性如此重要，请听 Jon Moore 关于为什么 [API 准入控制](https://www.youtube.com/watch?v=m64SWl9bfvk) 应该使用并发而不是吞吐量的演讲.并且，为了进一步阅读，请参考我关于 [排队理论](https://www.vividcortex.com/resources/queueing-theory) 和 [通用可扩展性定律](https://www.vividcortex.com/resources)的电子书/通用可扩展性法律。总之，如果您不确定哪些指标对监控和诊断最有用，USE 和 RED 是很好的起点。

