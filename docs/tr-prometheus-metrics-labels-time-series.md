# Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)

# Prometheus 备忘单 - 基础知识（指标、标签、时间序列、抓取）

July 24, 2021

2021 年 7 月 24 日

[Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)

[普罗米修斯](http://iximiuz.com/en/categories/?category=Prometheus)

This article is part of the **[Learning Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)** series:

本文是**[学习普罗米修斯](http://iximiuz.com/en/categories/?category=Prometheus)**系列的一部分：

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)
- [Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

- [Prometheus 不是 TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [如何通过 Prometheus Playground 学习 PromQL](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- Prometheus 备忘单 - 基础知识（指标、标签、时间序列、抓取）
- [Prometheus Cheat Sheet - 如何加入多个指标（向量匹配）](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus 备忘单 - 移动平均、最大值、最小值等（随时间的聚合）](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

Here we focus on the most basic Prometheus concepts - metrics, labels, scrapes, and time series.

在这里，我们专注于最基本的 Prometheus 概念 - 指标、标签、刮擦和时间序列。

## What is a metric?

## 什么是指标？

In Prometheus, everything revolves around [_metrics_](https://prometheus.io/docs/concepts/data_model/). A _metric_ is a feature (i.e., a characteristic) of a system that is being measured. Typical examples of metrics are:

在 Prometheus 中，一切都围绕着 [_metrics_](https://prometheus.io/docs/concepts/data_model/)。 _metric_ 是被测量系统的特征（即特性)。指标的典型示例是：

- http\_requests\_total
- http\_request\_size\_bytes
- system\_memory\_used\_bytes
- node\_network\_receive\_bytes\_total

- http\_requests\_total
- http\_request\_size\_bytes
- system\_memory\_used\_bytes
- 节点\_network\_receive\_bytes\_total

![Prometheus metrics](http://iximiuz.com/prometheus-metrics-labels-time-series/metrics-2000-opt.png)

## What is a label?

## 什么是标签？

The idea of a _metric_ seems fairly simple. However, there is a problem with such simplicity. On the diagram above, Prometheus monitors several application servers simultaneously. Each of these servers reports `mem_used_bytes` metric. At any given time, how should Prometheus store multiple samples behind a single metric name?

_metric_ 的想法似乎相当简单。但是，这种简单性存在问题。在上图中，Prometheus 同时监控多个应用服务器。这些服务器中的每一个都报告“mem_used_bytes”指标。在任何给定时间，Prometheus 应该如何在单个度量名称后面存储多个样本？

The first option is _aggregation_. Prometheus could sum up all the bytes and store the total memory usage of the whole fleet. Or compute an average memory usage and store it. Or min/max memory usage. Or compute and store all of those together. However, there is always a problem with storing only aggregated metrics - we wouldn't be able to pin down a particular server with a bizarre memory usage pattern based on such data.

第一个选项是_aggregation_。 Prometheus 可以汇总所有字节并存储整个队列的总内存使用量。或者计算平均内存使用量并存储它。或最小/最大内存使用量。或者将所有这些计算并存储在一起。但是，仅存储聚合指标总是存在问题 - 我们无法根据此类数据确定具有奇怪内存使用模式的特定服务器。

Luckily, Prometheus uses another approach - it can differentiate samples with the same metric name by labeling them. A _label_ is a certain attribute of a metric. Generally, labels are populated by metric producers (servers in the example above). However, in Prometheus, it's possible to enrich a metric with some static labels based on the producer's identity while recording it on the Prometheus node's side. In the wild, it's common for a Prometheus metric to carry multiple labels.

幸运的是，Prometheus 使用了另一种方法——它可以通过标记来区分具有相同度量名称的样本。 _label_ 是指标的某个属性。通常，标签由指标生成器（上例中的服务器）填充。但是，在 Prometheus 中，可以根据生产者的身份使用一些静态标签来丰富指标，同时将其记录在 Prometheus 节点一侧。在野外，Prometheus 指标带有多个标签是很常见的。

Typical examples of labels are:

标签的典型例子是：

- instance - an_instance_ (a server or cronjob process) of a _job_ being monitored in the `<host>:<port>` form
- job - a name of a logical group of_instances_ sharing the same purpose
- endpoint - name of an HTTP API endpoint
- method - HTTP method
- status\_code - HTTP status code

- 实例 - 以 `<host>:<port>` 形式监控的 _job_ 的 an_instance_（服务器或 cronjob 进程）
- 作业 - 具有相同目的的逻辑组的名称_实例_
- 端点 - HTTP API 端点的名称
- 方法 - HTTP 方法
- status\_code - HTTP 状态码

![Prometheus labels](http://iximiuz.com/prometheus-metrics-labels-time-series/labels-2000-opt.png)

## What is scraping?

## 什么是刮擦？

There are two principally different approaches to collect metrics. A monitoring system can have a _passive_ or _active_ collector component. In the case of a passive collector, samples are constantly _pushed_ by active instances to the collector. In contrast, an active collector periodically _pulls_ samples from instances that passively expose them.

有两种主要不同的方法来收集指标。监控系统可以有一个 _passive_ 或 _active_ 收集器组件。在被动收集器的情况下，样本不断被主动实例_推送_到收集器。相比之下，主动收集器会定期从被动公开它们的实例中抽取样本。

Prometheus uses a _pull model_, and the metric collection process is called _scraping_.

Prometheus 使用_pull 模型_，度量收集过程称为_scraping_。

In a system with a _passive collector_, there is no need to register monitored instances upfront. Instead, you need to communicate the address of the collector endpoint to the instances, so they could start pushing data. However, in the case of an _active collector_, one should supply the list of instances to be scraped beforehand. Or teach the monitoring system how to build such a list dynamically using one of the supported service discovery mechanisms. 

在带有_被动收集器_的系统中，无需预先注册受监控的实例。相反，您需要将收集器端点的地址传达给实例，以便它们可以开始推送数据。但是，对于 _active 收集器_，应该事先提供要抓取的实例列表。或者教监控系统如何使用支持的服务发现机制之一动态构建这样的列表。

In Prometheus, scraping is configured via providing a static list of `<host>:<port>` **_scraping targets_**. It's also possible to configure a service-discovery-specific (consul, docker, kubernetes, ec2, etc.) endpoint to fetch such a list at runtime. You also need to specify a **_scrape interval_** \- a delay between any two consecutive scrapes. Surprisingly or not, it's common for such an interval to be several seconds or even tens of seconds long.

在 Prometheus 中，抓取是通过提供`<host>:<port>` **_scraping targets_** 的静态列表来配置的。还可以配置特定于服务发现的（consul、docker、kubernetes、ec2 等）端点以在运行时获取此类列表。您还需要指定 **_scrape interval_** \- 任何两个连续刮擦之间的延迟。不管是否令人惊讶，这样的间隔通常有几秒甚至几十秒长。

For a monitoring system, the design choice to use a pull model and relatively long scrape intervals have some interesting repercussions...

对于监控系统，使用拉模型和相对较长的抓取间隔的设计选择会产生一些有趣的影响......

## What is a time series in Prometheus?

## Prometheus 中的时间序列是什么？

_**Side note 1:** Despite being born in the age of distributed systems, every Prometheus server node is autonomous. I.e., there is no distributed metric storage in the default Prometheus setup, and every node acts as a self-sufficient monitoring server with local metric storage. It simplifies a lot of things, including the following explanation, because we don't need to think of how to merge overlapping series from different Prometheus nodes_ 😉

_**旁注 1：** 尽管诞生于分布式系统时代，但每个 Prometheus 服务器节点都是自治的。即，在默认的 Prometheus 设置中没有分布式度量存储，每个节点都充当具有本地度量存储的自给自足的监控服务器。它简化了很多事情，包括下面的解释，因为我们不需要考虑如何合并来自不同 Prometheus 节点的重叠系列_😉

![Prometheus series](http://iximiuz.com/prometheus-metrics-labels-time-series/time-series-2000-opt.png)

In general, a stream of timestamped values is called a **_time series_**. In the above example, there are four different time series. But only two metric names. I.e., a time series in Prometheus is defined by a combination of a metric name and a particular set of key-value labels.

通常，带时间戳的值流称为**_时间序列_**。在上面的例子中，有四个不同的时间序列。但只有两个指标名称。即，Prometheus 中的时间序列由度量名称和一组特定的键值标签的组合定义。

_**Side note 2:** Values are always floating-point numbers; timestamps are integers storing the number of milliseconds since the Unix epoch._

_**旁注 2：** 值总是浮点数；时间戳是存储自 Unix 纪元以来的毫秒数的整数。_

Every such time series is stored separately on the Prometheus node in the form of an append-only file. Since a series is defined by the label value(s), one needs to be careful with labels that might have high cardinality.

每个这样的时间序列都以仅附加文件的形式单独存储在 Prometheus 节点上。由于一系列是由标签值定义的，因此需要小心可能具有高基数的标签。

The terms _time series_, _series_, and _metric_ are often used interchangeably. However, in Prometheus, a metric technically means a group of [time] series.

术语_时间序列_、_系列_和_度量_经常互换使用。但是，在 Prometheus 中，度量在技术上意味着一组 [时间] 系列。

## Downsides of active scraping

## 主动抓取的缺点

Since it's a single node scraping multiple distributed endpoints with potentially different performance and network conditions, the exact sample timestamps will (although most of the time just slightly) vary for every scrape. Because of that and of the potential loss of some scrapes, the interval between two samples in a given time series is neither constant nor multiplication of the scrape interval. Remember the repercussions I mentioned above?

由于它是单个节点抓取具有潜在不同性能和网络条件的多个分布式端点，因此每次抓取的确切样本时间戳（尽管大多数时间只是略有不同）会有所不同。由于这一点以及一些刮擦的潜在损失，给定时间序列中两个样本之间的间隔既不是常数也不是刮擦间隔的乘积。还记得我上面提到的影响吗？

![Missing scrapes illustration](http://iximiuz.com/prometheus-metrics-labels-time-series/scrape-interval-drift-2000-opt.png)

_Prometheus node scraping two services every 10 seconds - actual samples aren't ideally aligned in time._

_Prometheus 节点每 10 秒抓取两个服务 - 实际样本在时间上并不理想。_

There is another interesting, more important pitfall to be aware of. If a target reports a _gauge_ (i.e., _instant measurement_) metric that changes more frequently than it's scraped, the intermediate values will never be seen by the Prometheus node. Thus, it may cause _blindness_ of the monitoring system to some bizarre patterns:

还有另一个有趣的、更重要的陷阱需要注意。如果目标报告的 _gauge_（即 _instant measure_）指标的变化频率高于其抓取频率，则 Prometheus 节点将永远不会看到中间值。因此，它可能会导致监控系统_盲目_出现一些奇怪的模式：

Obviously, _counter_ (i.e., monotonically incrementing measurement) metrics don't have such a problem.

显然，_counter_（即单调递增的度量）指标没有这样的问题。

#### See other posts in the series

#### 查看该系列的其他帖子

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

- [Prometheus 不是 TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [如何通过 Prometheus Playground 学习 PromQL](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - 如何加入多个指标（向量匹配）](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus 备忘单 - 移动平均、最大值、最小值等（随时间的聚合）](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

[prometheus,](javascript: void 0) [metric,](javascript: void 0) [label,](javascript: void 0) [scraping](javascript: void 0)

[prometheus,](javascript: void 0) [metric,](javascript: void 0) [label,](javascript: void 0) [scraping](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

