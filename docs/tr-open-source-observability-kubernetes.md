# Introduction to open source observability on Kubernetes

# Kubernetes 开源可观察性介绍

## In the first article in this series, learn the signals, mechanisms, tools, and platforms you can use to observe services running on Kubernetes.

## 在本系列的第一篇文章中，学习可用于观察在 Kubernetes 上运行的服务的信号、机制、工具和平台。

07 Oct 2019

With the advent of DevOps, engineering teams are taking on more and more ownership of the reliability of their services. While some chafe at the increased operational burden, others welcome the opportunity to treat service reliability as a key feature, invest in the necessary capabilities to measure and improve reliability, and deliver the best possible customer experiences.

随着 DevOps 的出现，工程团队对其服务的可靠性承担了越来越多的责任。虽然有些人对增加的运营负担感到恼火，但其他人欢迎有机会将服务可靠性视为关键特征，投资必要的能力来衡量和提高可靠性，并提供尽可能最好的客户体验。

This change is measured explicitly in the [2019 Accelerate State of DevOps Report](https://cloud.google.com/blog/products/devops-sre/the-2019-accelerate-state-of-devops-elite-performance-productivity-and-scaling). One of its most interesting conclusions (as written in the summary) is:

> "Delivering software quickly, **reliably**, and safely is at the heart of technology transformation and organizational performance. We see continued evidence that software speed, stability, and **availability** __ contribute to organizational performance (including profitability, productivity, and customer satisfaction). Our highest performers are twice as likely to meet or exceed their organizational performance goals."

[2019 年 DevOps 加速状态报告](https://cloud.google.com/blog/products/devops-sre/the-2019-accelerate-state-of-devops-elite-performance-生产力和规模）。其最有趣的结论之一（如摘要中所写)是：

> “快速、**可靠地**、安全地交付软件是技术转型和组织绩效的核心。我们看到持续的证据表明软件速度、稳定性和**可用性** 有助于组织绩效（包括盈利能力、生产力和客户满意度）。我们表现最好的人达到或超过其组织绩效目标的可能性是其两倍。”

The full [report](https://services.google.com/fh/files/misc/state-of-devops-2019.pdf) says:

> " **Low performers use more proprietary software than high and elite performers**: The cost to maintain and support proprietary software can be prohibitive, prompting high and elite performers to use open source solutions. This is in line with results from previous reports . In fact, the 2018 Accelerate State of DevOps Report found that elite performers were 1.75 times more likely to make extensive use of open source components, libraries, and platforms."

完整的 [报告](https://services.google.com/fh/files/misc/state-of-devops-2019.pdf) 说：

> " **低绩效者比优秀者和精英者使用更多的专有软件**：维护和支持专有软件的成本高得令人望而却步，促使高绩效者和精英者使用开源解决方案。这与之前报告的结果一致. 事实上，2018 年 DevOps 加速状态报告发现，精英执行者广泛使用开源组件、库和平台的可能性是其他人的 1.75 倍。”

This is a strong testament to the value of open source as a general accelerator of performance. Combining these two conclusions leads to the rather obvious thesis for this series:

> Reliability is a critical feature, observability is a necessary component of reliability, and open source tooling is at least _A_ right approach, if not _THE_ right approach.

这有力地证明了开源作为通用性能加速器的价值。结合这两个结论，可以得出本系列相当明显的论点：

> 可靠性是一个关键特性，可观察性是可靠性的必要组成部分，开源工具至少是 _A_ 正确的方法，如果不是 _THE_ 正确的方法。

This article, the first in a series, will introduce the types of signals engineers typically rely on and the mechanisms, tools, and platforms that you can use to instrument services running on Kubernetes to emit these signals, ingest and store them, and use and interpret them.

本文是该系列的第一篇，将介绍工程师通常依赖的信号类型，以及可用于检测在 Kubernetes 上运行的服务以发出这些信号、摄取和存储它们以及使用和使用的机制、工具和平台。解释它们。

From there, the series will continue with hands-on tutorials, where I will walk through getting started with each of the tools and technologies. By the end, you should be well-equipped to start improving the observability of your own systems!

从那里开始，该系列将继续提供动手教程，在那里我将逐步介绍每种工具和技术的入门。最后，您应该准备好开始提高自己系统的可观察性！

## What is observability?

## 什么是可观察性？

While observability as a general [concept in control theory](https://en.wikipedia.org/wiki/Observability) has been around since at least 1960, its applicability to digital systems and services is rather new and in some ways an evolution of how these systems have been monitored for the last two decades. You are likely familiar with the necessity of monitoring services to ensure you know about issues before your users are impacted. You are also likely familiar with the idea of using metric data to better understand the health and state of a system, especially in the context of troubleshooting during an incident or debugging.

虽然可观察性作为一般[控制理论中的概念](https://en.wikipedia.org/wiki/Observability) 至少从 1960 年就已经存在，但它对数字系统和服务的适用性是相当新的，并且在某些方面是一种演变过去二十年如何监控这些系统。您可能熟悉监控服务的必要性，以确保您在用户受到影响之前了解问题。您可能还熟悉使用指标数据更好地了解系统的运行状况和状态的想法，尤其是在事件或调试期间进行故障排除的上下文中。

The key differentiation between monitoring and observability is that observability is an inherent property of a system or service, rather than something someone does to the system, which is what monitoring fundamentally is. [Cindy Sridharan](https://twitter.com/copyconstruct), author of a free [e-book](https://t.co/0gOgZp88Jn?amp=1) on observability in distributed systems, does a great job of explaining the difference in an excellent [Medium article](https://medium.com/@copyconstruct/monitoring-and-observability-8417d1952e1c). 

监控和可观察性之间的关键区别在于，可观察性是系统或服务的固有属性，而不是某人对系统所做的事情，而这正是监控的根本所在。 [Cindy Sridharan](https://twitter.com/copyconstruct)，一本关于分布式系统可观察性的免费[电子书](https://t.co/0gOgZp88Jn?amp=1) 的作者，做得很好在一篇优秀的 [Medium 文章](https://medium.com/@copyconstruct/monitoring-and-observability-8417d1952e1c)中解释差异。

It is important to distinguish between these two terms because observability, as a property of the service you build, is your responsibility. As a service developer and owner, you have full control over the signals your system emits, how and where those signals are ingested and stored, and how they're utilized. This is in contrast to "monitoring," which may be done by others (and by you) to measure the availability and performance of your service and generate alerts to let you know that service reliability has degraded.

区分这两个术语很重要，因为可观察性作为您构建的服务的属性，是您的责任。作为服务开发人员和所有者，您可以完全控制系统发出的信号、这些信号的摄取和存储方式和位置，以及它们的使用方式。这与“监控”形成对比，“监控”可能由其他人（和您）完成，以衡量您的服务的可用性和性能，并生成警报以让您知道服务可靠性已降低。

## Signals

## 信号

Now that you understand the idea of observability as a property of a system that you control and that is explicitly manifested as the signals you instruct your system to emit, it's important to understand and describe the kinds of signals generally considered in this context.

既然您已将可观察性的概念理解为您控制的系统的一种属性，并且明确表现为您指示系统发出的信号，那么理解和描述在此上下文中通常考虑的信号类型就很重要了。

### What are metrics?

### 什么是指标？

A metric is a fundamental type of signal that can be emitted by a service or the infrastructure it's running on. At its most basic, it is the combination of:

 指标是一种基本类型的信号，可以由服务或其运行的基础设施发出。在最基本的情况下，它是以下各项的组合：

1. Some identifier, hopefully descriptive, that indicates what the metric represents

   一些标识符，希望是描述性的，表明指标代表什么

2. A series of data points, each of which contains two elements:

   一系列数据点，每个数据点包含两个元素：


     a. The timestamp at which the data point was generated (or ingested)
     生成（或摄取）数据点的时间戳


    b. A numeric value representing the state of the thing you're measuring at that time
    	一个数值，表示您当时正在测量的事物的状态

Time-series metrics have been and remain the key data structure used in monitoring and observability practice and are the primary way that the state and health of a system are represented over time. They are also the primary mechanism for alerting, but that practice and others (like incident management, on-call, and postmortems) are outside the scope here. For now, the focus is on how to instrument systems to emit metrics, how to store them, and how to use them for charts and dashboards to help you visualize the current and historical state of your system.

时间序列指标一直是并且仍然是监控和可观察性实践中使用的关键数据结构，并且是随着时间的推移表示系统状态和健康状况的主要方式。它们也是警报的主要机制，但这种做法和其他做法（如事件管理、随叫随到和事后分析）超出了此处的范围。目前，重点是如何检测系统以发出指标、如何存储它们，以及如何将它们用于图表和仪表板，以帮助您可视化系统的当前和历史状态。

Metrics are used for two primary purposes: health and insight.

指标有两个主要目的：健康和洞察力。

Understanding the health and state of your infrastructure, platform, and service is essential to keeping them available to users. Generally, these are emitted by the various components chosen to build services, and it's just a matter of setting up the right collection and storage infrastructure to be able to use them. Metrics from the simple (node CPU utilization) to the esoteric (garbage collection statistics) fall into this category.

了解基础设施、平台和服务的运行状况和状态对于让用户可以使用它们至关重要。通常，这些是由为构建服务而选择的各种组件发出的，只需设置正确的收集和存储基础设施即可使用它们。从简单（节点 CPU 利用率）到深奥（垃圾收集统计）的指标都属于这一类。

Metrics are also essential to understanding what is happening in the system to avoid interruptions to your services. From this perspective, a service can emit custom telemetry that precisely describes specific aspects of how the service is functioning and performing. This will require you to instrument the code itself, usually by including specific libraries, and specify an export destination.

指标对于了解系统中正在发生的事情以避免服务中断也是必不可少的。从这个角度来看，服务可以发出自定义遥测，精确描述服务如何运作和执行的特定方面。这将要求您检测代码本身，通常是通过包含特定的库，并指定导出目标。

### What are logs?

### 什么是日志？

Unlike metrics that represent numeric values that change over time, logs represent discrete events. Log entries contain both the log payload—the message emitted by a component of the service or the code—and often metadata, such as the timestamp, label, tag, or other identifiers. Therefore, this is by far the largest volume of data you need to store, and you should carefully consider your log ingestion and storage strategies as you look to take on increasing user traffic.

与表示随时间变化的数值的指标不同，日志表示离散事件。日志条目包含日志负载（由服务的组件或代码发出的消息）以及元数据，例如时间戳、标签、标记或其他标识符。因此，这是迄今为止您需要存储的最大数据量，当您希望承担不断增加的用户流量时，您应该仔细考虑您的日志摄取和存储策略。

### What are traces?

### 什么是痕迹？

Distributed tracing is a relatively new addition to the observability toolkit and is specifically relevant to microservice architectures to allow you to understand latency and how various backend service calls contribute to it. Ted Young published an [excellent article on the concept](https://opensource.com/article/18/5/distributed-tracing) that includes its origins with Google's [Dapper paper](https://research.google.com/pubs/pub36356.html) and subsequent evolution. This series will be specifically concerned with the various implementations available.

分布式跟踪是可观察性工具包的一个相对较新的补充，它与微服务架构特别相关，可让您了解延迟以及各种后端服务调用对延迟的影响。 Ted Young 发表了一篇[关于该概念的优秀文章](https://opensource.com/article/18/5/distributed-tracing)，其中包括其起源于 Google 的 [Dapper 论文](https://research.google.com/pubs/pub36356.html) 和随后的演变。本系列将特别关注各种可用的实现。

## Instrumentation 

## 仪表

Once you identify the signals you want to emit, store, and analyze, you need to instruct your system to create the signals and build a mechanism to store and analyze them. Instrumentation refers to those parts of your code that are used to generate metrics, logs, and traces. In this series, we'll discuss open source instrumentation options and introduce the basics of their use through hands-on tutorials.

一旦确定了要发出、存储和分析的信号，就需要指示系统创建信号并构建一种机制来存储和分析它们。检测是指用于生成指标、日志和跟踪的代码部分。在本系列中，我们将讨论开源检测选项，并通过动手教程介绍其使用的基础知识。

## Observability on Kubernetes

## Kubernetes 上的可观察性

Kubernetes is the dominant platform today for deploying and maintaining containers. As it rose to the top of the industry's consciousness, so did new technologies to provide effective observability tooling around it. Here is a short list of these essential technologies; they will be covered in greater detail in future articles in this series.

Kubernetes 是当今用于部署和维护容器的主要平台。随着它上升到行业意识的顶端，围绕它提供有效的可观察性工具的新技术也随之兴起。以下是这些基本技术的简短列表；在本系列的后续文章中将更详细地介绍它们。

### Metrics

### 指标

Once you select your preferred approach for instrumenting your service with metrics, the next decision is where to store those metrics and what set of services will support your effort to monitor your environment.

一旦您选择了使用指标检测服务的首选方法，下一个决定就是将这些指标存储在何处以及哪些服务集将支持您监控环境的工作。

#### Prometheus

#### 普罗米修斯

[Prometheus](https://prometheus.io/) is the best place to start when looking to monitor both your Kubernetes infrastructure and the services running in the cluster. It provides everything you'll need, including client instrumentation libraries, the [storage backend](https://prometheus.io/docs/prometheus/latest/storage/), a visualization UI, and an alerting framework. Running Prometheus also provides a wealth of infrastructure metrics right out of the box. It further provides [integrations](https://prometheus.io/docs/operating/integrations/#remote-endpoints-and-storage) with third-party providers for storage, although the data exchange is not bi-directional in every case , so be sure to read the documentation if you want to store metric data in multiple locations.

[Prometheus](https://prometheus.io/) 是监控 Kubernetes 基础设施和集群中运行的服务的最佳起点。它提供了您需要的一切，包括客户端检测库、[存储后端](https://prometheus.io/docs/prometheus/latest/storage/)、可视化 UI 和警报框架。运行 Prometheus 还提供了大量开箱即用的基础设施指标。它进一步向第三方提供商提供 [integrations](https://prometheus.io/docs/operating/integrations/#remote-endpoints-and-storage) 以进行存储，尽管数据交换并非在所有情况下都是双向的，因此如果您想在多个位置存储指标数据，请务必阅读文档。

Later in this series, I will walk through setting up Prometheus in a cluster for basic infrastructure monitoring and adding custom telemetry to an application using the Prometheus client libraries.

在本系列的后面部分，我将逐步介绍如何在集群中设置 Prometheus 以进行基本的基础设施监控，以及使用 Prometheus 客户端库向应用程序添加自定义遥测。

#### Graphite

#### 石墨

[Graphite](https://graphiteapp.org/) grew out of an in-house development effort at Orbitz and is now positioned as an enterprise-ready monitoring tool. It provides metrics storage and retrieval mechanisms, but no instrumentation capabilities. Therefore, you will still need to implement Prometheus or OpenCensus instrumentation to collect metrics. Later in this series, I will walk through setting up Graphite and sending metrics to it.

[Graphite](https://graphiteapp.org/) 源于 Orbitz 的内部开发工作，现在被定位为企业级监控工具。它提供指标存储和检索机制，但没有检测功能。因此，您仍然需要实施 Prometheus 或 OpenCensus 检测来收集指标。在本系列的后面，我将逐步介绍如何设置 Graphite 并向其发送指标。

#### InfluxDB

#### InfluxDB

[InfluxDB](https://www.influxdata.com/get-influxdb/) is another open source database purpose-built for storing and retrieving time-series metrics. Unlike Graphite, InfluxDB is supported by a company called InfluxData, which provides both the InfluxDB software and a cloud-hosted version called InfluxDB Cloud. Later in this series, I will walk through setting up InfluxDB in a cluster and sending metrics to it.

[InfluxDB](https://www.influxdata.com/get-influxdb/) 是另一个开源数据库，专门用于存储和检索时间序列指标。与 Graphite 不同，InfluxDB 由一家名为 InfluxData 的公司提供支持，该公司提供 InfluxDB 软件和名为 InfluxDB Cloud 的云托管版本。在本系列的后面，我将逐步介绍在集群中设置 InfluxDB 并向其发送指标。

#### OpenTSDB

#### OpenTSDB

[OpenTSDB](http://opentsdb.net/) is also an open source purpose-built time-series database. One of its advantages is the ability to use [HBase](https://hbase.apache.org/) as the storage layer, which allows integration with a cloud managed service like Google's Cloud Bigtable. Google has published a [reference guide](https://cloud.google.com/solutions/opentsdb-cloud-platform) on setting up OpenTSDB to monitor your Kubernetes cluster (assuming it's running in Google Kubernetes Engine, or GKE). Since it's a great introduction, I recommend following Google's tutorial if you're interested in learning more about OpenTSDB.

[OpenTSDB](http://opentsdb.net/) 也是一个开源的专用时间序列数据库。它的优势之一是能够使用 [HBase](https://hbase.apache.org/) 作为存储层，这允许与像谷歌的 Cloud Bigtable 这样的云管理服务集成。 Google 发布了一个 [参考指南](https://cloud.google.com/solutions/opentsdb-cloud-platform)，用于设置 OpenTSDB 以监控您的 Kubernetes 集群（假设它在 Google Kubernetes Engine 或 GKE 中运行)。由于这是一个很好的介绍，如果您有兴趣了解有关 OpenTSDB 的更多信息，我建议您遵循 Google 的教程。

#### OpenCensus 

[OpenCensus](https://opencensus.io/) is the open source version of the [Census library](https://opensource.googleblog.com/2018/03/how-google-uses-opencensus-internally.html) developed at Google. It provides both metric and tracing instrumentation capabilities and supports a number of backends to [export](https://opencensus.io/exporters/#exporters) the metrics to—including Prometheus! Note that OpenCensus does not monitor your infrastructure, and you will still need to determine the best approach if you choose to use OpenCensus for custom metric telemetry.

[OpenCensus](https://opencensus.io/) 是 [Census library](https://opensource.googleblog.com/2018/03/how-google-uses-opencensus-internally.html)的开源版本) 在 Google 开发。它提供指标和跟踪检测功能，并支持多个后端[导出](https://opencensus.io/exporters/#exporters) 指标到——包括 Prometheus！请注意，OpenCensus 不会监控您的基础设施，如果您选择使用 OpenCensus 进行自定义指标遥测，您仍然需要确定最佳方法。

We'll revisit this library later in this series, and I will walk through creating metrics in a service and exporting them to a backend.

我们将在本系列的后面部分重新访问这个库，我将逐步介绍在服务中创建指标并将它们导出到后端。

### Logging for observability

### 记录可观察性

If metrics provide "what" is happening, logging tells part of the story of "why." Here are some common options for consistently gathering and analyzing logs.

如果指标提供了正在发生的“什么”，日志记录说明了“为什么”的部分故事。以下是一些用于持续收集和分析日志的常用选项。

#### Collecting with fluentd

#### 使用 fluentd 收集

In the Kubernetes ecosystem, [fluentd](https://www.fluentd.org/) is the de-facto open source standard for collecting logs emitted in the cluster and forwarding them to a specified backend. You can use config maps to modify fluentd's behavior, and later in the series, I'll walk through deploying it in a cluster and modifying the associated config map to parse unstructured logs and convert them to structured for better and easier analysis. In the meantime, you can read my post " [Customizing Kubernetes logging (Part 1)](https://medium.com/google-cloud/customizing-kubernetes-logging-part-1-a1e5791dcda8)" on how to do that on GKE.

在 Kubernetes 生态系统中，[fluentd](https://www.fluentd.org/) 是事实上的开源标准，用于收集集群中发出的日志并将它们转发到指定的后端。您可以使用配置映射来修改 fluentd 的行为，在本系列的后面部分，我将逐步介绍如何将其部署到集群中并修改关联的配置映射以解析非结构化日志并将它们转换为结构化日志，以便更好、更轻松地进行分析。同时，您可以阅读我的帖子“[自定义 Kubernetes 日志记录（第 1 部分)](https://medium.com/google-cloud/customizing-kubernetes-logging-part-1-a1e5791dcda8)”，了解如何做到这一点在 GKE 上。

#### Storing and analyzing with ELK

#### 使用 ELK 进行存储和分析

The most common storage mechanism for logs is provided by [Elastic](https://www.elastic.co/) in the form of the "ELK" stack. As Elastic says:

> "'ELK' is the acronym for three open source projects: Elasticsearch, Logstash, and Kibana. Elasticsearch is a search and analytics engine. Logstash is a server‑side data processing pipeline that ingests data from multiple sources simultaneously, transforms it, and then sends it to a 'stash' like Elasticsearch. Kibana lets users visualize data with charts and graphs in Elasticsearch."

最常见的日志存储机制是[Elastic](https://www.elastic.co/)以“ELK”栈的形式提供的。正如弹性所说：

> "'ELK' 是三个开源项目的首字母缩写词：Elasticsearch、Logstash 和 Kibana。Elasticsearch 是一个搜索和分析引擎。Logstash 是一个服务器端数据处理管道，它可以同时从多个来源获取数据，对其进行转换，然后然后将其发送到 Elasticsearch 之类的“存储库”。Kibana 允许用户使用 Elasticsearch 中的图表和图形来可视化数据。”

Later in the series, I'll walk through setting up Elasticsearch, Kibana, and Logstash in

在本系列的后面部分，我将逐步介绍如何设置 Elasticsearch、Kibana 和 Logstash

a cluster to store and analyze logs being collected by fluentd.

一个集群，用于存储和分析 fluentd 收集的日志。

### Distributed traces and observability

### 分布式跟踪和可观察性

When asking "why" in analyzing service issues, logs can only provide the information that applications are designed to share with it. The way to go even deeper is to gather traces. As the [OpenTracing initiative](https://opentracing.io/docs/overview/what-is-tracing) says:

> "Distributed tracing, also called distributed request tracing, is a method used to profile and monitor applications, especially those built using a microservices architecture. Distributed tracing helps pinpoint where failures occur and what causes poor performance."

当在分析服务问题时询问“为什么”时，日志只能提供应用程序旨在与其共享的信息。更深入的方法是收集痕迹。正如 [OpenTracing Initiative](https://opentracing.io/docs/overview/what-is-tracing) 所说：

> “分布式跟踪，也称为分布式请求跟踪，是一种用于分析和监控应用程序的方法，尤其是那些使用微服务架构构建的应用程序。分布式跟踪有助于查明发生故障的位置以及导致性能不佳的原因。”

#### Istio 

The [Istio](http://istio.io/) open source service mesh provides multiple benefits for microservice architectures, including traffic control, security, and observability capabilities. It does not combine multiple spans into a single trace to assemble a full picture of what happens when a user call traverses a distributed system, but it can nevertheless be useful as an easy first step toward distributed tracing. It also provides other observability benefits—it's the easiest way to get ["golden signal"](https://landing.google.com/sre/sre-book/chapters/monitoring-distributed-systems/) metrics for each service, and it also adds logging for each request, which can be very useful for calculating error rates. You can read my post on [using it with Google's Stackdriver](https://medium.com/google-cloud/istio-and-stackdriver-59d157282258). I'll revisit it in this series and show how to install it in a cluster and configure it to export observability data to a backend.

[Istio](http://istio.io/) 开源服务网格为微服务架构提供了多种优势，包括流量控制、安全性和可观察性能力。它不会将多个跨度组合到单个跟踪中来组装用户调用遍历分布式系统时发生的情况的完整图片，但它仍然可以作为通往分布式跟踪的简单第一步很有用。它还提供了其他可观察性优势——这是获取每个服务的 [“黄金信号”](https://landing.google.com/sre/sre-book/chapters/monitoring-distributed-systems/) 指标的最简单方法，它还为每个请求添加日志记录，这对于计算错误率非常有用。您可以阅读我关于 [将其与 Google 的 Stackdriver 一起使用](https://medium.com/google-cloud/istio-and-stackdriver-59d157282258) 的帖子。我将在本系列中重新讨论它并展示如何将它安装在集群中并配置它以将可观察性数据导出到后端。

#### OpenCensus

I described [OpenCensus](http://opencensus.io/) in the Metrics section above, and that's one of the main reasons for choosing it for distributed tracing: Using a single library for both metrics and traces is a great option to reduce your instrumentation work—with the caveat that you must be working in a language that supports both the traces and stats exporters. I'll come back to OpenCensus and show how to get started instrumenting code for distributed tracing. Note that OpenCensus provides only the instrumentation library, and you'll still need to use a storage and visualization layer like Zipkin, Jaeger, Stackdriver (on GCP), or X-Ray (on AWS).

我在上面的指标部分描述了 [OpenCensus](http://opencensus.io/)，这是选择它进行分布式跟踪的主要原因之一：对指标和跟踪使用单个库是减少您的检测工作 - 需要注意的是，您必须使用一种支持跟踪和统计数据导出器的语言。我将回到 OpenCensus 并展示如何开始为分布式跟踪检测代码。请注意，OpenCensus 仅提供检测库，您仍需要使用存储和可视化层，例如 Zipkin、Jaeger、Stackdriver（在 GCP 上）或 X-Ray（在 AWS 上)。

#### Zipkin

[Zipkin](https://zipkin.io/) is a full, distributed tracing solution that includes instrumentation, storage, and visualization. It's a tried and true set of tools that's been around for years and has a strong user and developer community. It can also be used as a backend for other instrumentation options like OpenCensus. In a future tutorial, I'll show how to set up the Zipkin server and instrument your code.

[Zipkin](https://zipkin.io/) 是一个完整的分布式跟踪解决方案，包括检测、存储和可视化。这是一套久经考验的真正工具，已经存在多年并且拥有强大的用户和开发人员社区。它也可以用作 OpenCensus 等其他检测选项的后端。在以后的教程中，我将展示如何设置 Zipkin 服务器并检测您的代码。

#### Jaeger

#### 耶格

[Jaeger](https://www.jaegertracing.io/) is another open source tracing solution that includes all the components you'll need. It's a newer project that's being incubated at the Cloud Native Computing Foundation (CNCF). Whether you choose to use Zipkin or Jaeger may ultimately depend on your experience with them and their support for the language you're writing your service in. In this series, I'll walk through setting up Jaeger and instrumenting code for tracing.

[Jaeger](https://www.jaegertracing.io/) 是另一种开源跟踪解决方案，其中包含您需要的所有组件。这是一个较新的项目，正在云原生计算基金会 (CNCF) 孵化。您选择使用 Zipkin 还是 Jaeger 最终可能取决于您使用它们的经验以及它们对您编写服务所用语言的支持。在本系列中，我将逐步介绍设置 Jaeger 和检测代码以进行跟踪。

## Visualizing observability data

## 可视化可观察性数据

The final piece of the toolkit for using metrics is the visualization layer. There are basically two options here: the "native" visualization that your persistence layers enable (e.g., the Prometheus UI or Flux with InfluxDB) or a purpose-built visualization tool.

使用指标的工具包的最后一部分是可视化层。这里基本上有两个选项：持久层启用的“原生”可视化（例如，Prometheus UI 或 Flux with InfluxDB）或专门构建的可视化工具。

[Grafana](https://grafana.com/) is currently the de facto standard for open source visualization. I'll walk through setting it up and using it to visualize data from various backends later in this series.

[Grafana](https://grafana.com/) 目前是开源可视化的事实上的标准。在本系列的后面部分，我将介绍如何设置它并使用它来可视化来自各种后端的数据。

## Looking ahead

##  展望未来

Observability on Kubernetes has many parts and many options for each type of need. Metric, logging, and tracing instrumentation provide the bedrock of information needed to make decisions about services. Instrumenting, storing, and visualizing data are also essential. Future articles in this series will dive into all of these options with hands-on tutorials for each.

Kubernetes 上的可观察性有很多部分，针对每种类型的需求都有很多选择。指标、日志记录和跟踪检测提供了做出服务决策所需的信息基础。检测、存储和可视化数据也是必不可少的。本系列的后续文章将深入探讨所有这些选项，并为每个选项提供动手教程。

[![Coffee shop photo](https://opensource.com/sites/default/files/styles/teaser-wide/public/lead-images/coffee-shop-devops.png?itok=ewhDnS9x)](http://opensource.com/article/19/6/you-cant-buy-devops) 

