# Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)

June 13, 2021 (Updated: July 1, 2021)

[Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)

This article is part of the **[Learning Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)** series:

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)


PromQL looks neat and powerful. And at first sight, simple. But when you start using it for real, you'll quickly notice that it's far from being trivial. Searching the Internet for query explanation rarely helps - most articles focus on pretty high-level overviews of the language's most basic capabilities. For example, when I needed to match multiple metrics using the common labels, I quickly found myself reading the [code implementing binary operations on vectors](https://github.com/prometheus/prometheus/blob/c0c22ed04200a8d24d1d5719f605c85710f0d008/promql/engine.go#L1833-L1942). Without a solid understanding of the matching rules, I constantly stumbled upon various query execution errors, such as complaints about missing `group_left` or `group_right` modifier. Reading the code, feeding my local Prometheus playground with artificial metrics, running test queries, and validating assumptions, finally helped me understand how multiple metrics can be joined together. Below are my findings.

## PromQL binary operators

PromQL comes with 15 binary operators that can be divided into three groups by operation type:

- arithmetic`+ - / * ^ %`
- comparison`< > <= >= == !=`
- logical/set`and, unless, or`

Binary operations are defined for different types of operands - scalar/scalar, scalar/vector, and vector/vector. And the last pair of operands, vector/vector, is the most puzzling one because the subtle vector matching rules may differ depending on the cardinality of the sides or operation type.

## One-to-one vector matching

The following diagram tries to shed some light on:

- how_one-to-one_ vector matching works
- how`on` and `ignoring` modifiers reduce the set of labels to be used for matching
- what problem this label set reduction can cause

[![PromQL one-to-one vector matching - arithmetic and comparison operations](http://iximiuz.com/prometheus-vector-matching/vector-matching-1-2000-opt.png)](http://iximiuz.com/prometheus-vector-matching/vector-matching-1.png)

_PromQL one-to-one vector matching - arithmetic and comparison operations (clickable, 1.1 MB)._

_**NOTE:** check out [this article](https://www.robustperception.io/whats-in-a-__name__) to learn why some operations preserve the metric name in the result and some produce nameless series._

## One-to-many and many-to-one vector matching

One-to-one matching is the most straightforward one. Most of the time, `on` or `ignoring` modifiers help to make a query return something reasonable. But the pitfall here is that some query results may show a one-to-one cardinality only by coincidence. For instance, when our assumptions of the potential set of values for a given label are erroneous. So, it's just so happened that the query showed a one-to-one data relationship on a selected time range, but it can be one-to-many in general. Or many-to-one.

Luckily, Prometheus does support _many-to-one_ and _one-to-many_ vector matching. But it has to be specified explicitly by adding either `group_left` or `group_right` modifier to a query. Otherwise, the following error might be returned during the query execution:

```
multiple matches for labels: many-to-one matching must be explicit (group_left/group_right)

```

Unless logical binary operator `and|unless|or` is used, Prometheus always considers at least one side of the binary operation as having the cardinality of _"one"_. If during a query execution Prometheus finds a collision (label-wise) on the _"one"_ side, the query will fail with the following error:

```
found duplicate series for the match group <keys/values> on the <left|right> hand-side of the operation: <op>;
many-to-many matching not allowed: matching labels must be unique on one side

```

Interesting, that even if the _"one"_ side doesn't have collisions and `group_left` or `group_right` is specified, a query can still fail with:

```
multiple matches for labels: grouping labels must ensure unique matches

```

It can happen because, for every element on the _"many"_ side, Prometheus should find no more than one element from the _"one"_ side. Otherwise, the query result would become ambiguous. If the requested label matching doesn't allow to build an unambiguous result, Prometheus just fails the query.

[![PromQL many-to-one and one-to-many vector matching - arithmetic and comparison operations](http://iximiuz.com/prometheus-vector-matching/vector-matching-2-2000-opt.png)](http://iximiuz.com/prometheus-vector-matching/vector-matching-2.png)

_PromQL many-to-one and one-to-many vector matching - arithmetic and comparison operations (clickable, 1.2 MB)._

## Many-to-many vector matching (logical/set operations)

Logical ( _aka_ set) binary operators `and`, `unless`, and `or` surprisingly adhere to a simpler vector matching logic. These operations are always _many-to-many_. Hence no `group_left` or `group_right` may be needed. The following diagram focuses on how logical/set operations behave in Prometheus:

[![PromQL many-to-many vector matching - logical/set operations](http://iximiuz.com/prometheus-vector-matching/vector-matching-3-2000-opt.png)](http://iximiuz.com/prometheus-vector-matching/vector-matching-3.png)

_PromQL many-to-many vector matching - logical/set operations (clickable, 1.4 MB)._

## Instead of conclusion

If you find this useful, check out [other my Prometheus drawing](https://twitter.com/iximiuz/status/1402315573766328322). And if the idea of querying files such as Nginx or Envoy access logs with PromQL-like syntax sounds interesting to you, please give the `pq` project a star on GitHub. It really fuels me up:

[pq - parse and query files with PromQL-like syntax](https://github.com/iximiuz/pq)

## Resources

- [Prometheus repo on GitHub](https://github.com/prometheus/prometheus/)
- [Prometheus docs on operators](https://prometheus.io/docs/prometheus/latest/querying/operators/) \- turns out it's a very good resource when you already understand the internals. Can't really call it beginner-friendly, though.

#### See other posts in the series

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)](http://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

[prometheus,](javascript: void 0) [promql,](javascript: void 0) [labels,](javascript: void 0) [metrics,](javascript: void 0) [vector,](javascript: void 0) [matching,](javascript: void 0) [join](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

