# Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)

# Prometheus Cheat Sheet - 如何加入多个指标（向量匹配）

June 13, 2021 (Updated: July 1, 2021)

2021 年 6 月 13 日（更新：2021 年 7 月 1 日）

[Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)

[普罗米修斯](http://iximiuz.com/en/categories/?category=Prometheus)

This article is part of the **[Learning Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)** series:

本文是**[学习普罗米修斯](http://iximiuz.com/en/categories/?category=Prometheus)**系列的一部分：

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

- [Prometheus 不是 TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [如何通过 Prometheus Playground 学习 PromQL](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus 备忘单 - 基础知识（指标、标签、时间序列、抓取）](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- Prometheus 备忘单 - 如何加入多个指标（向量匹配）
- [Prometheus 备忘单 - 移动平均、最大值、最小值等（随时间的聚合）](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

PromQL looks neat and powerful. And at first sight, simple. But when you start using it for real, you'll quickly notice that it's far from being trivial. Searching the Internet for query explanation rarely helps - most articles focus on pretty high-level overviews of the language's most basic capabilities. For example, when I needed to match multiple metrics using the common labels, I quickly found myself reading the [code implementing binary operations on vectors](https://github.com/prometheus/prometheus/blob/c0c22ed04200a8d24d1d5719f605c85710f0d008/promql/engine.go#L1833-L1942). Without a solid understanding of the matching rules, I constantly stumbled upon various query execution errors, such as complaints about missing `group_left` or `group_right` modifier. Reading the code, feeding my local Prometheus playground with artificial metrics, running test queries, and validating assumptions, finally helped me understand how multiple metrics can be joined together. Below are my findings.

PromQL 看起来简洁而强大。乍一看，很简单。但是当您开始真正使用它时，您会很快注意到它远非微不足道。在 Internet 上搜索查询解释很少有帮助 - 大多数文章都集中在语言最基本功能的相当高级的概述上。例如，当我需要使用公共标签匹配多个指标时，我很快发现自己阅读了[在向量上实现二进制操作的代码](https://github.com/prometheus/prometheus/blob/c0c22ed04200a8d24d1d5719f605c85710f0d008/promql/engine.去#L1833-L1942)。在对匹配规则没有深入了解的情况下，我经常偶然发现各种查询执行错误，例如抱怨缺少“group_left”或“group_right”修饰符。阅读代码、为我的本地 Prometheus 游乐场提供人工指标、运行测试查询和验证假设，最终帮助我理解了如何将多个指标连接在一起。以下是我的发现。

## PromQL binary operators

## PromQL 二元运算符

PromQL comes with 15 binary operators that can be divided into three groups by operation type:

PromQL 带有 15 个二元运算符，可以按操作类型分为三组：

- arithmetic`+ - / * ^ %`
- comparison`< > <= >= == !=`
- logical/set`and, unless, or`

- 算术`+ - /* ^ %`
- 比较`< > <= >= == !=`
- 逻辑/设置`和，除非，或`

Binary operations are defined for different types of operands - scalar/scalar, scalar/vector, and vector/vector. And the last pair of operands, vector/vector, is the most puzzling one because the subtle vector matching rules may differ depending on the cardinality of the sides or operation type.

二元运算是为不同类型的操作数定义的——标量/标量、标量/向量和向量/向量。最后一对操作数，向量/向量，是最令人费解的，因为微妙的向量匹配规则可能会根据边的基数或操作类型而有所不同。

## One-to-one vector matching

## 一对一向量匹配

The following diagram tries to shed some light on:

下图试图阐明：

- how_one-to-one_ vector matching works
- how`on` and `ignoring` modifiers reduce the set of labels to be used for matching
- what problem this label set reduction can cause

- how_one-to-one_矢量匹配工作
- how`on` 和 `ignoring` 修饰符减少了用于匹配的标签集
- 这个标签集减少会导致什么问题

[![PromQL one-to-one vector matching - arithmetic and comparison operations](http://iximiuz.com/prometheus-vector-matching/vector-matching-1-2000-opt.png)](http://iximiuz.com/prometheus-vector-matching/vector-matching-1.png)

iximiuz.com/prometheus-vector-matching/vector-matching-1.png)

_PromQL one-to-one vector matching - arithmetic and comparison operations (clickable, 1.1 MB)._

_PromQL 一对一向量匹配 - 算术和比较运算（可点击，1.1 MB）。_

_**NOTE:** check out [this article](https://www.robustperception.io/whats-in-a-__name__) to learn why some operations preserve the metric name in the result and some produce nameless series. _

_**注意：** 查看 [这篇文章](https://www.robustperception.io/whats-in-a-__name__) 了解为什么有些操作会在结果中保留度量名称，而有些操作会产生无名系列。 _

## One-to-many and many-to-one vector matching

## 一对多和多对一向量匹配

One-to-one matching is the most straightforward one. Most of the time, `on` or `ignoring` modifiers help to make a query return something reasonable. But the pitfall here is that some query results may show a one-to-one cardinality only by coincidence. For instance, when our assumptions of the potential set of values for a given label are erroneous. So, it's just so happened that the query showed a one-to-one data relationship on a selected time range, but it can be one-to-many in general. Or many-to-one.

一对一匹配是最直接的一种。大多数情况下，`on` 或 `ignoring` 修饰符有助于使查询返回合理的内容。但这里的陷阱是，某些查询结果可能只是巧合地显示一对一的基数。例如，当我们对给定标签的潜在值集的假设是错误的时。因此，查询在选定的时间范围内显示一对一的数据关系是碰巧的，但通常可以是一对多的。或者多对一。

Luckily, Prometheus does support _many-to-one_ and _one-to-many_ vector matching. But it has to be specified explicitly by adding either `group_left` or `group_right` modifier to a query. Otherwise, the following error might be returned during the query execution:

幸运的是，Prometheus 确实支持 _many-to-one_ 和 _one-to-many_ 向量匹配。但必须通过向查询添加 `group_left` 或 `group_right` 修饰符来明确指定。否则，在查询执行过程中可能会返回以下错误：

```
multiple matches for labels: many-to-one matching must be explicit (group_left/group_right)

```

Unless logical binary operator `and|unless|or` is used, Prometheus always considers at least one side of the binary operation as having the cardinality of _"one"_. If during a query execution Prometheus finds a collision (label-wise) on the _"one"_ side, the query will fail with the following error:

除非使用逻辑二元运算符 `and|unless|or`，Prometheus 总是认为二元运算的至少一侧具有 _"one"_ 的基数。如果在查询执行期间 Prometheus 在 _"one"_ 端发现冲突（标签方式），查询将失败并显示以下错误：

```
found duplicate series for the match group <keys/values> on the <left|right> hand-side of the operation: <op>;
many-to-many matching not allowed: matching labels must be unique on one side

```

Interesting, that even if the _"one"_ side doesn't have collisions and `group_left` or `group_right` is specified, a query can still fail with:

有趣的是，即使 _"one"_ 侧没有冲突并且指定了 `group_left` 或 `group_right`，查询仍然可能失败：

```
multiple matches for labels: grouping labels must ensure unique matches

```

It can happen because, for every element on the _"many"_ side, Prometheus should find no more than one element from the _"one"_ side. Otherwise, the query result would become ambiguous. If the requested label matching doesn't allow to build an unambiguous result, Prometheus just fails the query.

之所以会发生这种情况，是因为对于 _"many"_ 一侧的每个元素，Prometheus 应该从 _"one"_ 一侧找到不超过一个元素。否则，查询结果将变得不明确。如果请求的标签匹配不允许构建明确的结果，Prometheus 只会使查询失败。

[![PromQL many-to-one and one-to-many vector matching - arithmetic and comparison operations](http://iximiuz.com/prometheus-vector-matching/vector-matching-2-2000-opt.png)](http://iximiuz.com/prometheus-vector-matching/vector-matching-2.png)

](http://iximiuz.com/prometheus-vector-matching/vector-matching-2.png)

_PromQL many-to-one and one-to-many vector matching - arithmetic and comparison operations (clickable, 1.2 MB)._

_PromQL 多对一和一对多向量匹配 - 算术和比较操作（可点击，1.2 MB）。_

## Many-to-many vector matching (logical/set operations)

## 多对多向量匹配（逻辑/集合操作）

Logical ( _aka_ set) binary operators `and`, `unless`, and `or` surprisingly adhere to a simpler vector matching logic. These operations are always _many-to-many_. Hence no `group_left` or `group_right` may be needed. The following diagram focuses on how logical/set operations behave in Prometheus:

逻辑（_aka_ set）二元运算符 `and`、`unless` 和 `or` 令人惊讶地遵循更简单的向量匹配逻辑。这些操作总是_多对多_。因此，可能不需要“group_left”或“group_right”。下图重点介绍了 Prometheus 中的逻辑/集合操作的行为方式：

[![PromQL many-to-many vector matching - logical/set operations](http://iximiuz.com/prometheus-vector-matching/vector-matching-3-2000-opt.png)](http://iximiuz.com/prometheus-vector-matching/vector-matching-3.png)

iximiuz.com/prometheus-vector-matching/vector-matching-3.png)

_PromQL many-to-many vector matching - logical/set operations (clickable, 1.4 MB)._

_PromQL 多对多向量匹配 - 逻辑/集合操作（可点击，1.4 MB）。_

## Instead of conclusion

## 而不是结论

If you find this useful, check out [other my Prometheus drawing](https://twitter.com/iximiuz/status/1402315573766328322). And if the idea of querying files such as Nginx or Envoy access logs with PromQL-like syntax sounds interesting to you, please give the `pq` project a star on GitHub. It really fuels me up:

如果您觉得这有用，请查看 [其他我的普罗米修斯绘图](https://twitter.com/iximiuz/status/1402315573766328322)。如果您觉得使用类似 PromQL 的语法查询 Nginx 或 Envoy 访问日志等文件的想法很有趣，请在 GitHub 上给 `pq` 项目一个星星。这真的让我兴奋不已：

[pq - parse and query files with PromQL-like syntax](https://github.com/iximiuz/pq)

[pq - 使用类似 PromQL 的语法解析和查询文件](https://github.com/iximiuz/pq)

## Resources

##  资源

- [Prometheus repo on GitHub](https://github.com/prometheus/prometheus/)
- [Prometheus docs on operators](https://prometheus.io/docs/prometheus/latest/querying/operators/) \- turns out it's a very good resource when you already understand the internals. Can't really call it beginner-friendly, though.

- [GitHub 上的 Prometheus 存储库](https://github.com/prometheus/prometheus/)
- [Prometheus 操作员文档](https://prometheus.io/docs/prometheus/latest/querying/operators/) \- 当您已经了解内部结构时，它是一个非常好的资源。不过，不能真正称其为初学者友好的。

#### See other posts in the series

#### 查看该系列的其他帖子

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

- [Prometheus 不是 TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [如何通过 Prometheus Playground 学习 PromQL](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus 备忘单 - 基础知识（指标、标签、时间序列、抓取）](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus 备忘单 - 移动平均、最大值、最小值等（随时间的聚合）](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

[prometheus,](javascript: void 0) [promql,](javascript: void 0) [labels,](javascript: void 0) [metrics,](javascript: void 0) [vector,](javascript: void 0) [matching,](javascript: void 0) [join](javascript: void 0)

[prometheus,](javascript: void 0) [promql,](javascript: void 0) [labels,](javascript: void 0) [metrics,](javascript: void 0) [vector,](javascript: void 0) [匹配,](javascript: void 0) [join](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

