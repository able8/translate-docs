# Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)

# Prometheus 备忘单 - 移动平均、最大值、最小值等（随时间的聚合）

June 28, 2021 (Updated: July 2, 2021)

2021 年 6 月 28 日（更新：2021 年 7 月 2 日）

[Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)

[普罗米修斯](http://iximiuz.com/en/categories/?category=Prometheus)

This article is part of the **[Learning Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)** series:

本文是**[学习普罗米修斯](http://iximiuz.com/en/categories/?category=Prometheus)**系列的一部分：

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)

- [Prometheus 不是 TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [如何通过 Prometheus Playground 学习 PromQL](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus 备忘单 - 基础知识（指标、标签、时间序列、抓取）](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - 如何加入多个指标（向量匹配）](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- Prometheus 备忘单 - 移动平均、最大值、最小值等（随时间的聚合）

When you have a long series of numbers, such as server memory consumption scraped 10-secondly, it's a natural desire to derive another, probably more meaningful series from it, by applying a _moving window function_. For instance, [_moving average_](https://en.wikipedia.org/wiki/Moving_average) or _moving quantile_ can give you much more readable results by smoothing some spikes.

当您有一个很长的数字系列时，例如服务器内存消耗每 10 秒被刮掉一次，很自然地希望通过应用_移动窗口函数_从中推导出另一个可能更有意义的系列。例如，[_moving average_](https://en.wikipedia.org/wiki/Moving_average) 或 _moving quantile_ 可以通过平滑一些尖峰为您提供更具可读性的结果。

Prometheus has a bunch of functions called `<smth>_over_time()`. They can be applied only to range vectors. It essentially makes them window aggregation functions. Every such function takes in a range vector and produces an instant vector with elements being per-series aggregations.

Prometheus 有一堆函数叫做 `<smth>_over_time()`。它们只能应用于范围向量。它本质上使它们成为窗口聚合函数。每个这样的函数都接受一个范围向量并产生一个即时向量，其中元素是每个系列的聚合。

For people like me who normally grasp code faster than text, here is some pseudocode of the aggregation logic:

对于像我这样通常比文本更快地掌握代码的人，这里是一些聚合逻辑的伪代码：

```python
# Input vector example.
range_vector = [
    ({"lab1": "val1", "lab2": "val2"}, [(12, 1624722138), (11, 1624722148), (17, 1624722158)]),
    ({"lab1": "val1", "lab2": "val2"}, [(14, 1624722138), (10, 1624722148), (13, 1624722158)]),
    ({"lab1": "val1", "lab2": "val2"}, [(16, 1624722138), (12, 1624722148), (15, 1624722158)]),
    ({"lab1": "val1", "lab2": "val2"}, [(12, 1624722138), (17, 1624722148), (18, 1624722158)]),
]

# agg_func examples: `sum`, `min`, `max`, `avg`, `last`, etc.

def agg_over_time(range_vector, agg_func, timestamp):
    # The future instant vector.
    instant_vector = {"timestamp": timestamp, "elements": []}

    for (labels, samples) in range_vector:
        # Every instant vector element is
        # an aggregation of multiple samples.
        sample = agg_func(samples)
        instant_vector["elements"].append((labels, sample))

    # Notice, that the timestamp of the resulting instant vector
    # is the timestamp of the query execution.I.e., it may not
    # match any of the timestamps in the input range vector.
    return instant_vector

```

Almost all the functions in the aggregation family accept just a single parameter - a _range vector_. It means that the _over time_ part, i.e., the duration of the aggregation period, comes from the range vector definition itself.

聚合系列中的几乎所有函数都只接受一个参数 - _range 向量_。这意味着_over time_部分，即聚合期的持续时间，来自范围向量定义本身。

The only way to construct a range vector in PromQL is by appending a bracketed duration to a vector selector. E.g. `http_requests_total[5m]`. Therefore, an `<agg>_over_time()` function can be applied only to a vector selector, meaning the aggregation will always be done using raw scrapes.

在 PromQL 中构造范围向量的唯一方法是将括号内的持续时间附加到向量选择器。例如。 `http_requests_total[5m]`。因此，`<agg>_over_time()` 函数只能应用于向量选择器，这意味着聚合将始终使用原始数据完成。

Another motivation for aggregating metrics over time comes from the fact that instant vectors of raw scrapes often are just glimpses of the actual state of your system. To facilitate graph plotting, Prometheus offers [_range API queries_](https://twitter.com/iximiuz/status/1402315573766328322). Logically, a response of a range API query consists of a series of instant vectors, where each vector is one _resolution step_ away from another. Depending on how lucky you are, you can miss important insights hidden in between any two consecutive vectors. 

随着时间的推移聚合指标的另一个动机来自这样一个事实，即原始数据的即时向量通常只是系统实际状态的一瞥。为了方便绘图，Prometheus 提供了 [_range API 查询_](https://twitter.com/iximiuz/status/1402315573766328322)。从逻辑上讲，范围 API 查询的响应由一系列即时向量组成，其中每个向量与另一个向量相距 _resolution step_。根据您的幸运程度，您可能会错过隐藏在任何两个连续向量之间的重要见解。

Unlike scalar elements of an instant vector, each element of the range vector is actually an array. Elements of such an array are scrape values falling into the corresponding time bucket. By applying an `<agg>_over_time` function to these buckets, we achieve a _moving window_ effect that potentially can incorporate all data points without leaving any gaps. If the window duration is at least one resolution step long, none of the available data points will be discarded.

与即时向量的标量元素不同，范围向量的每个元素实际上都是一个数组。这种数组的元素是落入相应时间段的刮取值。通过将 `<agg>_over_time` 函数应用于这些存储桶，我们实现了 _moving window_ 效果，可能会合并所有数据点而不会留下任何间隙。如果窗口持续时间至少是一个分辨率步长，则不会丢弃任何可用数据点。

![PromQL Functions - aggregation over time](http://iximiuz.com/prometheus-functions-agg-over-time/agg_over_time-2000-opt.png)

_Grafana tip: use `$__interval` built-in variable for your range vector durations._

_Grafana 提示：使用 `$__interval` 内置变量作为范围向量持续时间。_

### min, max, avg, sum, stddev, stdvar over time

### min、max、avg、sum、stddev、stdvar随着时间的推移

The naming makes the purpose of these functions quite obvious. For instance, `avg_over_time()` is what you may use to compute a **moving average** of some metric. Similarly, `stddev_over_time()` can be used to produce a **moving standard deviation**.

命名使这些函数的目的非常明显。例如，您可以使用 `avg_over_time()` 来计算某个指标的 **移动平均值**。类似地，`stddev_over_time()` 可用于产生**移动标准偏差**。

However, there is always a caveat. Prometheus supports different _logical types_ of metrics - _gauges_, _counters_, _histograms_, and _summaries_. Surprisingly or not, physically, there is no difference between these types. All of them are just stored as series of _float numbers_. PromQL doesn't really distinguish between the logical types of metrics either. What only matters for PromQL is an expression type. I.e., it wouldn't allow you to call a function that expects an instant vector with a range vector argument. However, a range vector of _gauges_ is physically indistinguishable from a range vector of _counters_. And here we go...

然而，总是有一个警告。 Prometheus 支持不同的 _logical types_ 指标 - _gauges_、_counters_、_histograms_ 和 _summaries_。无论是否令人惊讶，在物理上，这些类型之间没有区别。所有这些都只是存储为一系列_浮点数_。 PromQL 也没有真正区分指标的逻辑类型。 PromQL 唯一重要的是表达式类型。即，它不允许您调用需要具有范围向量参数的即时向量的函数。然而，_gauges_ 的范围向量与_counters_ 的范围向量在物理上是无法区分的。现在我们开始...

**functions `min_over_time()`, `max_over_time()`, `avg_over_time()`, `sum_over_time()`, `stddev_over_time()`, and `stdvar_over_time()` makes sense to use only with gauge metrics.**

**函数`min_over_time()`、`max_over_time()`、`avg_over_time()`、`sum_over_time()`、`stddev_over_time()` 和`stdvar_over_time()` 仅适用于仪表指标。**

Just think of it for a second. If you call, say, `min_over_time()` on a range vector of counters, it'd always return the left-most sample.

想一想。如果你在计数器的范围向量上调用`min_over_time()`，它总是返回最左边的样本。

### quantile\_over\_time()

### 分位数\_over\_time()

`quantile_over_time(scalar, range-vector)` \- this is the only one that takes two parameters: a scalar defining the quantile to compute and a range vector.

`quantile_over_time(scalar, range-vector)` \- 这是唯一需要两个参数的参数：定义要计算的分位数的标量和范围向量。

### last\_over\_time()

### last\_over\_time()

`last_over_time(http_requests_total[5m])` is pretty much the same as just `http_requests_total`. If we take the very last sample from every element of a range vector, the resulting vector will be identical to a regular instant vector query. The only usage that comes to mind is to be able to manipulate the [_lookback window_](https://twitter.com/iximiuz/status/1402315573766328322) on the fly.

`last_over_time(http_requests_total[5m])` 与 `http_requests_total` 几乎相同。如果我们从范围向量的每个元素中获取最后一个样本，则结果向量将与常规即时向量查询相同。唯一想到的用法是能够即时操纵 [_lookback window_](https://twitter.com/iximiuz/status/1402315573766328322)。

While constructing an instant vector, Prometheus looks up to 5 minutes back from the queried timestamp. For every metric, the very last sample falling into the `[t - 5m, t]` time range is then added to the resulting vector. The `5m` delta is a heuristical value that can be changed by providing `--query.lookback-delta` command-line parameter. However, changing it would require a server restart. Via using `last_over_time(<vector_selector>[<duration>])` we can control the lookback delta duration at runtime.

在构建即时向量时，Prometheus 从查询的时间戳开始最多查找 5 分钟。对于每个度量，落入`[t - 5m, t]` 时间范围内的最后一个样本然后被添加到结果向量中。 `5m` delta 是一个启发式值，可以通过提供 `--query.lookback-delta` 命令行参数来更改。但是，更改它需要重新启动服务器。通过使用`last_over_time(<vector_selector>[<duration>])`，我们可以在运行时控制回溯增量持续时间。

### count\_over\_time()

### 计数\_over\_time()

`count_over_time()` counts the number of samples for each element of the range vector. The actual values of the samples are disregarded. Thus, the resulting instant vector will just represent the number of scrapes per series falling into the requested time frame.

`count_over_time()` 计算范围向量的每个元素的样本数。样品的实际值被忽略。因此，生成的即时向量将仅表示落入请求的时间范围内的每个系列的刮擦次数。

#### See other posts in the series

#### 查看该系列的其他帖子

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)](http://iximiuz.com/en/posts/prometheus-vector-matching/)

- [Prometheus 不是 TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [如何通过 Prometheus Playground 学习 PromQL](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus 备忘单 - 基础知识（指标、标签、时间序列、抓取）](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - 如何加入多个指标（向量匹配）](http://iximiuz.com/en/posts/prometheus-vector-matching/)

[prometheus,](javascript: void 0) [promql,](javascript: void 0) [functions,](javascript: void 0) [aggregation,](javascript: void 0) [range-vector,](javascript: void 0) [moving-window](javascript: void 0)

[prometheus,](javascript: void 0) [promql,](javascript: void 0) [functions,](javascript: void 0) [aggregation,](javascript: void 0) [range-vector,](javascript: void 0) [移动窗口](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_ 

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

