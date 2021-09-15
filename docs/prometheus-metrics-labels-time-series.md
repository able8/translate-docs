# Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)

July 24, 2021

[Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)

This article is part of the **[Learning Prometheus](http://iximiuz.com/en/categories/?category=Prometheus)** series:

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- Prometheus Cheat Sheet - Basics (Metrics, Labels, Time Series, Scraping)
- [Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)


Here we focus on the most basic Prometheus concepts - metrics, labels, scrapes, and time series.

## What is a metric?

In Prometheus, everything revolves around [_metrics_](https://prometheus.io/docs/concepts/data_model/). A _metric_ is a feature (i.e., a characteristic) of a system that is being measured. Typical examples of metrics are:

- http\_requests\_total
- http\_request\_size\_bytes
- system\_memory\_used\_bytes
- node\_network\_receive\_bytes\_total

![Prometheus metrics](http://iximiuz.com/prometheus-metrics-labels-time-series/metrics-2000-opt.png)

## What is a label?

The idea of a _metric_ seems fairly simple. However, there is a problem with such simplicity. On the diagram above, Prometheus monitors several application servers simultaneously. Each of these servers reports `mem_used_bytes` metric. At any given time, how should Prometheus store multiple samples behind a single metric name?

The first option is _aggregation_. Prometheus could sum up all the bytes and store the total memory usage of the whole fleet. Or compute an average memory usage and store it. Or min/max memory usage. Or compute and store all of those together. However, there is always a problem with storing only aggregated metrics - we wouldn't be able to pin down a particular server with a bizarre memory usage pattern based on such data.

Luckily, Prometheus uses another approach - it can differentiate samples with the same metric name by labeling them. A _label_ is a certain attribute of a metric. Generally, labels are populated by metric producers (servers in the example above). However, in Prometheus, it's possible to enrich a metric with some static labels based on the producer's identity while recording it on the Prometheus node's side. In the wild, it's common for a Prometheus metric to carry multiple labels.

Typical examples of labels are:

- instance - an_instance_ (a server or cronjob process) of a _job_ being monitored in the `<host>:<port>` form
- job - a name of a logical group of_instances_ sharing the same purpose
- endpoint - name of an HTTP API endpoint
- method - HTTP method
- status\_code - HTTP status code

![Prometheus labels](http://iximiuz.com/prometheus-metrics-labels-time-series/labels-2000-opt.png)

## What is scraping?

There are two principally different approaches to collect metrics. A monitoring system can have a _passive_ or _active_ collector component. In the case of a passive collector, samples are constantly _pushed_ by active instances to the collector. In contrast, an active collector periodically _pulls_ samples from instances that passively expose them.

Prometheus uses a _pull model_, and the metric collection process is called _scraping_.

In a system with a _passive collector_, there is no need to register monitored instances upfront. Instead, you need to communicate the address of the collector endpoint to the instances, so they could start pushing data. However, in the case of an _active collector_, one should supply the list of instances to be scraped beforehand. Or teach the monitoring system how to build such a list dynamically using one of the supported service discovery mechanisms.

In Prometheus, scraping is configured via providing a static list of `<host>:<port>` **_scraping targets_**. It's also possible to configure a service-discovery-specific (consul, docker, kubernetes, ec2, etc.) endpoint to fetch such a list at runtime. You also need to specify a **_scrape interval_** \- a delay between any two consecutive scrapes. Surprisingly or not, it's common for such an interval to be several seconds or even tens of seconds long.

For a monitoring system, the design choice to use a pull model and relatively long scrape intervals have some interesting repercussions...

## What is a time series in Prometheus?

_**Side note 1:** Despite being born in the age of distributed systems, every Prometheus server node is autonomous. I.e., there is no distributed metric storage in the default Prometheus setup, and every node acts as a self-sufficient monitoring server with local metric storage. It simplifies a lot of things, including the following explanation, because we don't need to think of how to merge overlapping series from different Prometheus nodes_ 😉

![Prometheus series](http://iximiuz.com/prometheus-metrics-labels-time-series/time-series-2000-opt.png)

In general, a stream of timestamped values is called a **_time series_**. In the above example, there are four different time series. But only two metric names. I.e., a time series in Prometheus is defined by a combination of a metric name and a particular set of key-value labels.

_**Side note 2:** Values are always floating-point numbers; timestamps are integers storing the number of milliseconds since the Unix epoch._

Every such time series is stored separately on the Prometheus node in the form of an append-only file. Since a series is defined by the label value(s), one needs to be careful with labels that might have high cardinality.

The terms _time series_, _series_, and _metric_ are often used interchangeably. However, in Prometheus, a metric technically means a group of [time] series.

## Downsides of active scraping

Since it's a single node scraping multiple distributed endpoints with potentially different performance and network conditions, the exact sample timestamps will (although most of the time just slightly) vary for every scrape. Because of that and of the potential loss of some scrapes, the interval between two samples in a given time series is neither constant nor multiplication of the scrape interval. Remember the repercussions I mentioned above?

![Missing scrapes illustration](http://iximiuz.com/prometheus-metrics-labels-time-series/scrape-interval-drift-2000-opt.png)

_Prometheus node scraping two services every 10 seconds - actual samples aren't ideally aligned in time._

There is another interesting, more important pitfall to be aware of. If a target reports a _gauge_ (i.e., _instant measurement_) metric that changes more frequently than it's scraped, the intermediate values will never be seen by the Prometheus node. Thus, it may cause _blindness_ of the monitoring system to some bizarre patterns:

Obviously, _counter_ (i.e., monotonically incrementing measurement) metrics don't have such a problem.

#### See other posts in the series

- [Prometheus Is Not a TSDB](http://iximiuz.com/en/posts/prometheus-is-not-a-tsdb/)
- [How to learn PromQL with Prometheus Playground](http://iximiuz.com/en/posts/prometheus-learning-promql/)
- [Prometheus Cheat Sheet - How to Join Multiple Metrics (Vector Matching)](http://iximiuz.com/en/posts/prometheus-vector-matching/)
- [Prometheus Cheat Sheet - Moving Average, Max, Min, etc (Aggregation Over Time)](http://iximiuz.com/en/posts/prometheus-functions-agg-over-time/)

[prometheus,](javascript: void 0) [metric,](javascript: void 0) [label,](javascript: void 0) [scraping](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

