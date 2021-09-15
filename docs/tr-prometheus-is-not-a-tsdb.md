# Prometheus Is Not a TSDB

# Prometheus 不是 TSDB

July 24, 2021

2021 年 7 月 24 日

[Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)

[普罗米修斯](http://iximiuz.com/en/categories/?category=Prometheus)

This article is part of the **[Learning Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)** series:

本文是**[学习普罗米修斯](http://iximiuz.com/en/categories/?category=Prometheus)**系列的一部分：

- Prometheus Is Not a TSDB
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

- Prometheus 不是 TSDB
- [如何通过 Prometheus Playground 学习 PromQL](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus 备忘单 - 基础知识（指标、标签、时间序列、抓取）](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - 如何加入多个指标（向量匹配）](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus 备忘单 - 移动平均、最大值、最小值等（随时间的聚合）](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

_Misconception_ \- the right word to explain my early Prometheus experience. I came to Prometheus with vast Graphite and moderate InfluxDB experience. In my eyes, Graphite was a highly performant but fairly limited system. Metrics in Graphite are just strings (well, dotted), and the values are always stored aggregated with the lowest possible resolution of 1 second. But due to these limitations, Graphite is fast. In contrast, InfluxDB adopts [Metrics 2.0](http://metrics20.org/) format with multiple tags and fields per metric. It also allows the storage of non-aggregated data points with impressive nanosecond precision. But this power needs to be used carefully. Otherwise, you'll get all sorts of performance issues.

_误解_ \- 解释我早期普罗米修斯经历的正确词。我带着大量的 Graphite 和适度的 InfluxDB 经验来到 Prometheus。在我看来，Graphite 是一个高性能但相当有限的系统。 Graphite 中的指标只是字符串（好吧，虚线），并且这些值总是以尽可能低的 1 秒分辨率聚合存储。但由于这些限制，Graphite 速度很快。相比之下，InfluxDB 采用 [Metrics 2.0](http://metrics20.org/) 格式，每个指标具有多个标签和字段。它还允许以令人印象深刻的纳秒精度存储非聚合数据点。但这种力量需要谨慎使用。否则，您将遇到各种性能问题。

For some reason, I expected Prometheus to reside somewhere in between these two systems. A kinda-sorta system that takes the best of both worlds: rich labeled metrics, non-aggregated values, and high query performance.

出于某种原因，我希望 Prometheus 位于这两个系统之间的某个位置。一种兼具两全其美的系统：丰富的标记指标、非聚合值和高查询性能。

And at first, it indeed felt as such! But then I started noticing that I cannot really explain some of the query results. Like at all. Or sometimes, I couldn't find evidence in metrics that just had to be there. Like the metrics were showing me a different picture than I was observing with my eyes while analyzing raw data such as web server access logs.

而一开始，确实是这样的感觉！但后来我开始注意到我无法真正解释一些查询结果。就像。或者有时，我无法在必须存在的指标中找到证据。就像在分析原始数据（例如 Web 服务器访问日志）时，指标向我展示的图片与我用眼睛观察的图片不同。

So, I started looking for more details. I wanted to understand precisely how metrics are collected, how they are stored, what a query execution model is, et cetera, et cetera. And at first, I was shocked by my findings! Oftentimes, the Prometheus behavior didn't make any sense, especially comparing to Graphite or InfluxDB! But then it occurred to me that I was missing one important detail...

所以，我开始寻找更多的细节。我想准确地了解指标是如何收集的，它们是如何存储的，什么是查询执行模型，等等等等。起初，我对我的发现感到震惊！通常，Prometheus 的行为没有任何意义，尤其是与 Graphite 或 InfluxDB 相比！但后来我突然想到我遗漏了一个重要的细节......

Both Graphite and InfluxDB are pure time-series databases (TSDB). Yes, they are often used as metric storage for monitoring purposes. But every particular setup of these systems comes with certain trade-offs and bolt-on additions addressing performance or reliability concerns. For instance, there is often a statsd-like daemon in front of your Graphite doing preaggregation; you use different rollup strategies for older data points, etc. But normally, you are aware of that. So when you query the last couple of days of metrics, you expect them to have a secondly precision. But when you query something a week or a month old, you already know that each data point represents a minute of aggregated data, not a second.

Graphite 和 InfluxDB 都是纯时间序列数据库 (TSDB)。是的，它们通常用作用于监控目的的指标存储。但是这些系统的每个特定设置都伴随着某些权衡和附加功能，以解决性能或可靠性问题。例如，在您的 Graphite 前面通常有一个类似 statsd 的守护进程进行预聚合；您对较旧的数据点等使用不同的汇总策略。但通常情况下，您知道这一点。因此，当您查询最近几天的指标时，您希望它们具有第二个精度。但是当您查询一周或一个月前的内容时，您已经知道每个数据点代表一分钟的聚合数据，而不是一秒。

However, **Prometheus is not a TSDB.**

但是，**Prometheus 不是 TSDB。**

![Prometheus is not a TSDB](http://iximiuz.com/prometheus-is-not-a-tsdb/kdpv-2000-opt.png)

Prometheus is a monitoring system that just [happens to use a TSDB under the hood](https://prometheus.io/docs/introduction/comparison/). So, some of these trade-offs one would usually find in a typical Graphite installation were already taken by Prometheus authors. But maybe not just all of them were made clear in the documentation 🙈 

Prometheus 是一个监控系统，它只是 [碰巧在幕后使用 TSDB](https://prometheus.io/docs/introduction/comparison/)。因此，在典型的 Graphite 安装中通常会发现的一些权衡已经被 Prometheus 的作者采用了。但也许不是所有的都在文档中说清楚了🙈

So, when I would find a certain query producing quirky results, it would be due to my misunderstanding of Prometheus as a whole, not just its query execution model. Aiming for results that would be reasonable for a pure TSDB apparently may be way too expensive in a monitoring context. Prometheus, as a metric collection system, is tailored for monitoring purposes from day one. It does provide good collection, storage, and query performance. But it may sacrifice the data precision (scraping metrics every 10-30 seconds), or completeness (tolerating missing scrapes with a 5m long lookback delta), or extrapolate your latency distribution instead of keeping the actual measurements (that's how [histogram\_quantile( )](https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile) actually works).

所以，当我发现某个查询产生了古怪的结果时，那是由于我对 Prometheus 整体的误解，而不仅仅是它的查询执行模型。在监控环境中，针对纯 TSDB 合理的结果显然可能过于昂贵。 Prometheus 作为一个指标收集系统，从一开始就为监控目的而量身定制。它确实提供了良好的收集、存储和查询性能。但它可能会牺牲数据精度（每 10-30 秒抓取一次指标）或完整性（以 5m 长的回溯增量容忍缺失的抓取），或者推断您的延迟分布而不是保留实际测量值（这就是 [histogram\_quantile( )](https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile)实际上有效)。

Once I realized this difference, I reconsidered my expectations from the system. However, I still needed a better way to understand the PromQL's query execution model than the official documentation provides. There is also plenty of Prometheus and PromQL articles out there, but they give you just a shallow overview of the query language. Only a sweeping search finally revealed these two much more profound resources - [PromLabs' blog](https://promlabs.com/blog/) and [Robust Perception's blog](https://www.robustperception.io/blog). Both are apparently written by Prometheus creators.

一旦我意识到这种差异，我就会重新考虑我对系统的期望。但是，我仍然需要一种比官方文档提供的更好的方式来理解 PromQL 的查询执行模型。还有很多 Prometheus 和 PromQL 文章，但它们只是对查询语言进行了粗略的概述。经过一番搜索，终于发现了这两个更深刻的资源——[PromLabs的博客](https://promlabs.com/blog/)和[RobustPerception的博客](https://www.robustperception.io/blog)。两者显然都是由普罗米修斯的创作者编写的。

But practice is even better than reading 😉 So, I went ahead and hacked a local Prometheus playground to solidify my learnings. And of course, I'm always striving to write down stuff on the way, mostly for future me, but feel free to check out my other Prometheus posts:

但是练习比阅读更好 😉 所以，我继续黑进了当地的 Prometheus 游乐场来巩固我的学习。当然，我一直在努力写下东西，主要是为了未来的我，但请随时查看我的其他普罗米修斯帖子：

![PromQL diagrams](http://iximiuz.com/prometheus-is-not-a-tsdb/diagrams-2000-opt.png)

- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

- [如何通过 Prometheus Playground 学习 PromQL](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus 备忘单 - 基础知识（指标、标签、时间序列、抓取）](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - 如何加入多个指标（向量匹配）](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus 备忘单 - 移动平均、最大值、最小值等（随时间的聚合）](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

[prometheus,](javascript: void 0) [promql,](javascript: void 0) [tsdb,](javascript: void 0) [monitoring](javascript: void 0)

[prometheus,](javascript: void 0) [promql,](javascript: void 0) [tsdb,](javascript: void 0) [监控](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

