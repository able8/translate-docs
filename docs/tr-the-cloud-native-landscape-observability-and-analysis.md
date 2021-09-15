# TheCloudNativeLandscape:ObservabilityandAnalysis

# TheCloudNativeLandscape：可观察性和分析

#### 17 May 2021 1:24pm,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack. io/author/jason-morgan/ "Posts by Jason Morgan")

#### 2021 年 5 月 17 日下午 1:24，作者：[Catherine Paganini](https://thenewstack.io/author/catherine-paganini/“Catherine Paganini 的帖子”) 和 [Jason Morgan](https://thenewstack.io/author/jason-morgan/“杰森摩根的帖子”)

This post is part of an ongoing series from [Cloud Native Computing Foundation Business Value Subcommittee](https://lists.cncf.io/g/cncf-business-value) co-chairs [Catherine Paganini](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) and [Jason Morgan](https://thenewstack.io/author/jason-morgan/) that focuses on explaining each category of the cloud native landscape to a non -technical audience as well as engineers just getting started with cloud native computing.

这篇文章是 [云原生计算基金会商业价值小组委员会](https://lists.cncf.io/g/cncf-business-value) 联合主席 [Catherine Paganini](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) 和 [Jason Morgan](https://thenewstack.io/author/jason-morgan/) 专注于向非- 刚开始使用云原生计算的技术受众和工程师。

We finally arrived at the last section of our [Cloud Native Computing Foundation](https://cncf.io/?utm_content=inline-mention)‘s Landscape series. If you missed our previous articles, we covered an [introduction](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/), and then the [provisioning](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/), [runtime](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/) , [orchestration and management layer](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/), and [platforms](https://thenewstack.io/the-cloud-native-landscape-platforms-explained/) each in a separate article. Today, we’ll discuss each category of the observability and analysis “column.”

终于到了我们 [云原生计算基础](https://cncf.io/?utm_content=inline-mention) 景观系列的最后一部分。如果您错过了我们之前的文章，我们介绍了 [介绍](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/)，然后是 [provisioning](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/), [运行时](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/)、[编排和管理层](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/)和[平台](https://thenewstack.io/the-cloud-native-landscape-platforms-explained/) 每个都在单独的文章中。今天，我们将讨论可观察性和分析“列”的每个类别。

Let’s start by defining observability and analysis. Observability is a system characteristic describing the degree to which a system can be understood from its external outputs. Measured by CPU time, memory, disk space, latency, errors, etc., computer systems can be more or less observable. Analysis, on the other hand, is an activity in which you look at this observable data and make sense of it.

让我们从定义可观察性和分析开始。可观察性是描述系统可以从其外部输出中被理解的程度的系统特性。通过 CPU 时间、内存、磁盘空间、延迟、错误等来衡量，计算机系统或多或少都可以被观察到。另一方面，分析是一项活动，您可以在其中查看这些可观察数据并理解它。

To ensure there is no service disruption, you’ll need to observe and analyze every aspect of your application so any anomaly gets detected and rectified right away. This is what this category is all about. It runs across and observes all layers which is why it’s on the side and not embedded in a specific layer.

为确保没有服务中断，您需要观察和分析应用程序的各个方面，以便立即检测和纠正任何异常。这就是这个类别的全部内容。它运行并观察所有层，这就是为什么它在侧面而不是嵌入到特定层中。

Tools in this category are broken down into logging, monitoring, tracing, and chaos engineering. Please note that the category name is somewhat misleading. While listed here, chaos engineering is rather a reliability than an observability or analysis tool.

此类别中的工具分为日志记录、监控、跟踪和混沌工程。请注意，类别名称有些误导。虽然在这里列出，混沌工程与其说是一种可观察性或分析工具，不如说是一种可靠性。

## **Logging**

## **日志记录**

### What It Is

###  这是什么

Applications emit a steady stream of log messages describing what they are doing at any given time. These log messages capture various events happening in the system such as failed or successful actions, audit information, or health events. Logging tools collect, store, and analyze these messages to track error reports and related data. Along with metrics and tracing, logging is one of the pillars of observability.

应用程序发出稳定的日志消息流，描述它们在任何给定时间所做的事情。这些日志消息捕获系统中发生的各种事件，例如失败或成功的操作、审计信息或健康事件。日志工具收集、存储和分析这些消息以跟踪错误报告和相关数据。与指标和跟踪一起，日志记录是可观察性的支柱之一。

### Problem It Addresses

### 它解决的问题

Collecting, storing, and analyzing logs is a crucial part of building a modern platform. Logging helps with, performs one, or all of those tasks. Some tools handle every aspect from collection to analysis while others focus on a single task like collection. All logging tools aim at helping organizations gain control over their log messages.

收集、存储和分析日志是构建现代平台的关键部分。日志记录有助于执行其中一项或所有任务。一些工具处理从收集到分析的各个方面，而另一些工具则专注于收集等单一任务。所有日志工具都旨在帮助组织控制其日志消息。

### How It Helps

### 它如何帮助

When collecting, storing, and analyzing application log messages, you’ll understand what an application was communicating at any given time. But note, logs represent messages that applications can deliberately emit, they won’t necessarily pinpoint the root cause of a given issue. That being said, collecting and retaining log messages over time is an extremely powerful capability and will help teams diagnose issues and meet regulatory and compliance requirements.

在收集、存储和分析应用程序日志消息时，您将了解应用程序在任何给定时间通信的内容。但请注意，日志表示应用程序可以故意发出的消息，它们不一定能查明给定问题的根本原因。话虽如此，随着时间的推移收集和保留日志消息是一项非常强大的功能，将帮助团队诊断问题并满足法规和合规性要求。

### Technical 101 

### 技术 101

While collecting, storing, and processing log messages is by no means a new problem, cloud native patterns and Kubernetes have caused significant changes in the way we handle logs. Traditional approaches to logging that were appropriate for virtual and physical machines, like writing logs to a file, are ill-suited to containerized applications, where the file system doesn’t outlast an application. In a cloud native environment log collection tools like Fluentd, run alongside application containers and collect messages directly from the applications. Messages are then forwarded on to a central log store to be aggregated and analyzed.

虽然收集、存储和处理日志消息绝不是一个新问题，但云原生模式和 Kubernetes 已经导致我们处理日志的方式发生了重大变化。适用于虚拟机和物理机的传统日志记录方法（例如将日志写入文件）不适合容器化应用程序，其中文件系统不会比应用程序更持久。在像 Fluentd 这样的云原生环境日志收集工具中，与应用程序容器一起运行并直接从应用程序收集消息。然后将消息转发到中央日志存储进行聚合和分析。

Fluentd is the only CNCF project in this space.

Fluentd 是该领域唯一的 CNCF 项目。

**Buzzwords****Popular Projects**

**流行语****热门项目**

- Logging

- 记录

- Fluentd & Fluentbit
- Elastic Logstash

- Fluentd 和 Fluentbit
- 弹性 Logstash

![logging ](https://cdn.thenewstack.io/media/2021/05/61f23a99-screen-shot-2021-05-13-at-8.17.44-am.png)

## Monitoring

## 监控

### What It Is

###  这是什么

Monitoring refers to instrumenting an app to collect, aggregate, and analyze logs and metrics to improve our understanding of its behavior. While logs describe specific events, metrics are a measurement of a system at a given point in time — they are two different things but are both necessary to get the full picture of your system’s health. Monitoring includes everything from watching disk space, CPU usage, and memory consumption on individual nodes to doing detailed synthetic transactions to see if a system or application is responding correctly and in a timely manner. There are a number of different approaches to monitoring systems and applications.

监控是指对应用程序进行检测以收集、汇总和分析日志和指标，以提高我们对其行为的理解。虽然日志描述了特定事件，但指标是在给定时间点对系统的度量——它们是两种不同的东西，但对于全面了解系统的健康状况都是必要的。监控包括从观察单个节点上的磁盘空间、CPU 使用率和内存消耗到执行详细的综合事务以查看系统或应用程序是否正确及时地响应的所有内容。有许多不同的方法来监控系统和应用程序。

### Problem It Addresses

### 它解决的问题

When running an application or platform, you want it to accomplish a specific task as designed and ensure it’s only accessed by authorized users. Monitoring allows you to know if it is working correctly, securely, cost-effectively, only accessed by authorized users, and/or any other characteristic you may be tracking.

在运行应用程序或平台时，您希望它按照设计完成特定任务，并确保它只能由授权用户访问。监控可让您了解它是否正常、安全、经济高效地工作，是否仅由授权用户访问，和/或您可能正在跟踪的任何其他特征。

### How It Helps

### 它如何帮助

Good monitoring allows operators to respond quickly, and potentially automatically when an incident arises. It provides insights into the current health of a system and watches for changes. Monitoring tracks everything from application health to user behavior and is an essential part of effectively running applications.

良好的监控使操作员能够快速响应，并且可能在发生事故时自动响应。它提供对系统当前健康状况的洞察并观察变化。监控跟踪从应用程序运行状况到用户行为的所有内容，是有效运行应用程序的重要组成部分。

### Technical 101

### 技术 101

Monitoring in a cloud native context is generally similar to monitoring traditional applications. You need to track metrics, logs, and events to understand the health of your applications. The main difference is that some of the managed objects are ephemeral, meaning they may not be long-lasting so tying your monitoring to auto-generated resource names won’t be a good long-term strategy. There are a number of CNCF projects in this space that largely revolve around Prometheus, the CNCF graduated project.

在云原生环境中进行监控通常类似于监控传统应用程序。您需要跟踪指标、日志和事件以了解应用程序的运行状况。主要区别在于某些托管对象是短暂的，这意味着它们可能不会持久，因此将您的监控与自动生成的资源名称联系起来不是一个好的长期策略。在这个领域有许多 CNCF 项目主要围绕着 CNCF 毕业项目 Prometheus 展开。

**Buzzwords**

**流行语**

**Popular Projects/Products**

**热门项目/产品**

- Monitoring
- Time series
- Alerting
- Metrics

- 监控
- 时间序列
- 警报
- 指标

- Prometheus
- Cortex
- Thanos
- Grafana

- 普罗米修斯
- 皮质
- 灭霸
- 格拉法纳

![monitoring ](https://cdn.thenewstack.io/media/2021/05/dc4db049-screen-shot-2021-05-13-at-8.16.16-am.png)

## Tracing

## 追踪

### What It Is

###  这是什么

In a microservices world, services are constantly communicating with each other over the network. Tracing, a specialized use of logging, allows you to trace the path of a request as it moves through a distributed system.

在微服务世界中，服务通过网络不断地相互通信。跟踪是日志记录的一种特殊用途，它允许您在请求通过分布式系统时跟踪请求的路径。

### Problem It Addresses

### 它解决的问题

Understanding how a microservice application behaves at any given point in time is an extremely challenging task. While many tools provide deep insights into service behavior, it can be difficult to tie an action of an individual service to the broader understanding of how the entire app behaves.

了解微服务应用程序在任何给定时间点的行为是一项极具挑战性的任务。虽然许多工具提供了对服务行为的深入洞察，但很难将单个服务的操作与对整个应用程序行为的更广泛理解联系起来。

### How It Helps

### 它如何帮助

Tracing solves this problem by adding a unique identifier to messages sent by the application. That unique identifier allows you to follow, or trace, individual transactions as they move through your system. You can use this information to see both the health of your application as well as to debug problematic microservices or activities.

跟踪通过向应用程序发送的消息添加唯一标识符来解决此问题。该唯一标识符允许您在单个交易通过您的系统时对其进行跟踪或跟踪。您可以使用此信息查看应用程序的运行状况以及调试有问题的微服务或活动。

### Technical 101 

### 技术 101

Tracing is a powerful debugging tool that allows you to troubleshoot and fine-tune the behavior of a distributed application. That power does come at a cost. Application code needs to be modified to emit tracing data and any spans need to be propagated by infrastructure components in the data path of your application. Specifically service meshes’ and their proxies. Jaeger and Open Tracing are CNCF projects in this space.

跟踪是一种强大的调试工具，可让您对分布式应用程序的行为进行故障排除和微调。这种力量是有代价的。需要修改应用程序代码以发出跟踪数据，并且任何跨度都需要由应用程序数据路径中的基础设施组件传播。特别是服务网格及其代理。 Jaeger 和 Open Tracing 是该领域的 CNCF 项目。

**Buzzwords**

**流行语**

**Popular Projects**

**热门项目**

- Span
- Tracing

- 跨度
- 追踪

- Jaeger
- OpenTracing

- 耶格
- 开放追踪

![Tracing](https://cdn.thenewstack.io/media/2021/05/ea2aaf06-screen-shot-2021-05-13-at-8.28.38-am.png)

## **Chaos Engineering**

## **混沌工程**

### What It Is

###  这是什么

Chaos engineering refers to the practice of intentionally introducing faults into a system in order to create more resilient applications and engineering teams. A chaos engineering tool will provide a controlled way to introduce faults and run specific experiments against a particular instance of an application.

混沌工程是指有意将故障引入系统以创建更具弹性的应用程序和工程团队的做法。混沌工程工具将提供一种受控方式来引入故障并针对应用程序的特定实例运行特定实验。

### Problem It Addresses

### 它解决的问题

Complex systems fail. They fail for a host of reasons, and the consequences in a distributed system are typically hard to understand. Chaos engineering is embraced by organizations that accept that failures will occur and, instead of trying to prevent failures, practice recovering from them. This is referred to as optimizing for [mean time to repair](https://en.wikipedia.org/wiki/Mean_time_to_repair), or MTTR.

复杂的系统失败。它们失败的原因有很多，分布式系统中的后果通常很难理解。接受混沌工程的组织接受失败会发生，而不是试图防止失败，而是练习从中恢复。这称为优化 [平均修复时间](https://en.wikipedia.org/wiki/Mean_time_to_repair) 或 MTTR。

**Side note**: The traditional approach to maintaining high availability for applications is referred to as optimizing for [mean time between failures](https://en.wikipedia.org/wiki/Mean_time_between_failures), or MTBF. You can observe this practice in organizations that use things like “change review boards” and long change freezes” to keep an application environment stable by restricting changes. The authors of [Accelerate](https://itrevolution.com/book/accelerate/) suggest that high-performing IT organizations achieve high availability by optimizing for mean time to recovery, or MTTR, instead.

**旁注**：保持应用程序高可用性的传统方法被称为优化 [平均故障间隔时间](https://en.wikipedia.org/wiki/Mean_time_between_failures) 或 MTBF。您可以在使用“变更审查委员会”和长期变更冻结等手段通过限制变更来保持应用程序环境稳定的组织中观察这种做法。 [Accelerate](https://itrevolution.com/book/accelerate/) 的作者建议高性能 IT 组织通过优化平均恢复时间或 MTTR 来实现高可用性。

### How It Helps

### 它如何帮助

In a cloud native world, applications must dynamically adjust to failures — a relatively new concept. That means, when something fails, the system doesn’t go down completely but gracefully degrades or recovers. Chaos engineering tools enable you to experiment on a software system in production to ensure they do that should a real failure occur.

在云原生世界中，应用程序必须动态适应故障——这是一个相对较新的概念。这意味着，当出现故障时，系统不会完全崩溃，而是会优雅地降级或恢复。混沌工程工具使您能够在生产中的软件系统上进行试验，以确保它们在发生真正的故障时也能这样做。

In short, you experiment with a system because you want to be confident that it can withstand turbulent and unexpected conditions. Instead of waiting for something to happen and find out, you put it through duress under controlled conditions to identify weaknesses and fix them before chance uncovers them.

简而言之，您对系统进行试验是因为您希望确信它能够承受动荡和意外情况。与其等待某事发生并发现，您可以在受控条件下对其施加压力，以找出弱点并在机会发现之前修复它们。

### Technical 101

### 技术 101

Chaos engineering tools and practices are critical to achieving high availability for your applications. Distributed systems are often too complex to be fully understood by any one engineer and no change process can fully predetermine the impact of changes on an environment. By introducing deliberate chaos engineering practices teams are able to practice, and automate, recovering from failures. Chaos Mesh and Litmus Chaos are CNCF tools in this space but there are many open source and proprietary options available.

混沌工程工具和实践对于实现应用程序的高可用性至关重要。分布式系统通常过于复杂，任何一位工程师都无法完全理解，而且没有任何变更过程可以完全预先确定变更对环境的影响。通过引入深思熟虑的混沌工程实践，团队能够进行实践和自动化，从故障中恢复。 Chaos Mesh 和 Litmus Chaos 是该领域的 CNCF 工具，但有许多开源和专有选项可用。

**Buzzwords**

**流行语**

**Popular Projects**

**热门项目**

- Chaos engineering

- 混沌工程

- Chaos Mesh
- Litmus Chaos

- 混沌网格
- 石蕊混沌

![Chaos engineering ](https://cdn.thenewstack.io/media/2021/05/5463852b-screen-shot-2021-05-13-at-8.32.22-am.png)

As we’ve seen, the observability and analysis layer is all about understanding the health of your system and ensuring it stays operational even under tough conditions. Logging tools capture event messages emitted by apps, monitoring watches logs and metrics, and tracing follows the path of individual requests. When combined, these tools ideally provide a 360-degree view of what’s going on within your system. Chaos engineering is a little different. It provides a safe way to verify the system can withstand unexpected events, basically ensuring it stays healthy.

正如我们所见，可观察性和分析层旨在了解系统的健康状况并确保其即使在恶劣条件下也能保持运行。日志工具捕获应用程序发出的事件消息，监控手表日志和指标，并跟踪各个请求的路径。结合使用这些工具时，理想情况下可以 360 度全方位了解您的系统中发生的情况。混沌工程有点不同。它提供了一种安全的方法来验证系统是否可以承受意外事件，基本上确保它保持健康。

_This finally concludes our series on the CNCF landscape. We certainly learned a lot while writing these articles and hope you did too._

_这终于结束了我们关于 CNCF 景观的系列。我们在写这些文章的过程中当然学到了很多东西，希望你也这样做。_

Feature image [via](https://pixabay.com/fr/photos/jumelles-%C3%A0-la-recherche-l-homme-1209011/) Pixabay.

特色图片 [via](https://pixabay.com/fr/photos/jumelles-%C3%A0-la-recherche-l-homme-1209011/)

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Real.

The New Stack 是 Insight Partners 的全资子公司。 TNS 所有者 Insight Partners 是以下公司的投资者：Real。

This post is part of a larger story we're telling about Observability.

这篇文章是我们正在讲述的关于可观察性的更大故事的一部分。

[Notify me when ebook is available](http://thenewstack.io/ebooks/observability/cloud-native-observability-for-devops-teams/)

[电子书可用时通知我](http://thenewstack.io/ebooks/observability/cloud-native-observability-for-devops-teams/)

[Notify me when ebook is available](http://thenewstack.io/ebooks/observability/cloud-native-observability-for-devops-teams/)

[电子书可用时通知我](http://thenewstack.io/ebooks/observability/cloud-native-observability-for-devops-teams/)

[Cloud Native Observability for DevOps Teams](https://thenewstack.io/tag/cloud-native-observability-for-devops-teams/)[Contributed](https://thenewstack.io/tag/contributed/) 

[DevOps 团队的云原生可观察性](https://thenewstack.io/tag/cloud-native-observability-for-devops-teams/)[贡献](https://thenewstack.io/tag/contributed/)

