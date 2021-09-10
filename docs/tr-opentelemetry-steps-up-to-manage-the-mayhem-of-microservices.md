# OpenTelemetry Steps up to Manage the Mayhem of Microservices

# OpenTelemetry 加强管理微服务的混乱

17 Apr 2020



![](https://cdn.thenewstack.io/media/2020/04/89962ad8-open-telemetry.jpg)

Work continues to make [OpenTelemetry](https://opentelemetry.io/about/) the standard set of vendor-neutral specifications and associated tools for capturing cloud native operational data.

继续努力使 [OpenTelemetry](https://opentelemetry.io/about/) 成为一套标准的供应商中立规范和相关工具，用于捕获云原生操作数据。

Bringing the work close to production-readiness the project has released a collector, as well as software development kits (SDKs) to support a number of different languages.

为了使工作接近生产就绪，该项目发布了一个收集器，以及支持多种不同语言的软件开发工具包 (SDK)。

For end-users, the new release candidate offers end-users “confidence that OpenTelemetry is safe to evaluate because it is now at a place where it's nearing feature completion, runtime stability, and API stability,” wrote [Liz Fong-Jones]( https://twitter.com/lizthegrey), principal developer advocate for observability software provider and project contributor [Honeycomb](https://www.honeycomb.io/), in an e-mail.

Liz Fong-Jones 写道，对于最终用户而言，新的候选版本让最终用户“相信 OpenTelemetry 可以安全地进行评估，因为它现在已接近功能完成、运行时稳定性和 API 稳定性”。可观察性软件提供商和项目贡献者 [Honeycomb](https://www.honeycomb.io/) 的主要开发者倡导者，在一封电子邮件中。

She notes that the release signals to monitoring software and services providers that it is time to write their own exporters for the project, and it provides “a smoother and less friction path for users to get critical telemetry data into their app.”

她指出，该版本向监控软件和服务提供商发出信号，表明是时候为该项目编写自己的导出器了，它为用户提供了“一条更顺畅、更少摩擦的路径，让用户将关键遥测数据输入到他们的应用程序中”。

The work sets the stage for making “high-quality telemetry a built-in feature for any cloud native software,” said [Ben Sigelman](https://www.linkedin.com/in/bensigelman/), CEO and co- founder of cloud native monitoring provider [Lightstep](https://lightstep.com/), one of the contributors to the project.

首席执行官兼联合创始人 [Ben Sigelman](https://www.linkedin.com/in/bensigelman/) 表示，这项工作为“高质量遥测成为任何云原生软件的内置功能”奠定了基础。云原生监控提供商 [Lightstep](https://lightstep.com/) 创始人，项目贡献者之一。

The work is particularly pertinent as concerns grow that running microservices-based architectures in large-scale environments [may be too unwieldy](https://thenewstack.io/kelsey-hightower-and-ben-sigelman-debate-microservices-vs-monoliths/), given the difficulties in debugging across multiple components. “The challenge with telemetry in cloud native is that there’s a lot of it, especially for tracing data,” Sigelman said. A [sandbox project](https://landscape.cncf.io/selected=open-telemetry) of the Cloud Native Computing Foundation, OpenTelemetry would be a key enabler in making microservices manageable (OpenTelemetry itself was a 2019 merger between two overlapping projects , [OpenTracing and OpenCensus](https://thenewstack.io/opentracing-opencensus-merge-into-a-single-new-project-opentelemetry/)).

随着人们越来越担心在大规模环境中运行基于微服务的架构 [可能太笨拙](https://thenewstack.io/kelsey-hightower-and-ben-sigelman-debate-microservices-vs-monoliths/)，考虑到跨多个组件进行调试的困难。 “云原生遥测的挑战在于它有很多，尤其是在跟踪数据方面，”西格尔曼说。作为云原生计算基金会的 [沙盒项目](https://landscape.cncf.io/selected=open-telemetry)，OpenTelemetry 将成为使微服务可管理的关键推动因素（OpenTelemetry 本身是两个重叠项目之间的 2019 年合并, [OpenTracing 和 OpenCensus](https://thenewstack.io/opentracing-opencensus-merge-into-a-single-new-project-opentelemetry/))。

An open-source collaborative project across 82 different companies, OpenTelemetry is building a set of libraries, agents, and other components to aid in the observation, management and debugging of microservices and distributed applications. Telemetry is built on the three pillars of metrics, logs, and tracing, providing the context needed to trace an individual transaction as it flows through multiple components

作为一个跨 82 家不同公司的开源协作项目，OpenTelemetry 正在构建一组库、代理和其他组件，以帮助观察、管理和调试微服务和分布式应用程序。遥测建立在指标、日志和跟踪三大支柱之上，提供了在单个事务流经多个组件时对其进行跟踪所需的上下文

The goal with OpenTelemetry is not to provide a platform for observability, but rather to provide a standard substrate to collect and convey operational data so it can be used in monitoring and observational platforms, either of the open source or commercial variety.

OpenTelemetry 的目标不是提供一个可观察性平台，而是提供一个标准基础来收集和传输操作数据，以便它可以用于监控和观察平台，无论是开源还是商业平台。

Historically, when an enterprise would purchase a package for systems monitoring, all the agents that would be attached to the resources would be specific to that provider’s implementation. If a customer wanted to change out, the applications and infrastructure would have to be entirely re-instrumented, Sigelman explained. By using the OpenTelemetry, users could instrument their systems once and pick the best and visualization and analysis products for their workloads, and not worry about lock-in.

从历史上看，当企业购买用于系统监控的软件包时，所有附加到资源的代理都将特定于该提供商的实施。 Sigelman 解释说，如果客户想要更换，则必须对应用程序和基础设施进行完全重新检测。通过使用 OpenTelemetry，用户可以对他们的系统进行一次检测，并为他们的工作负载选择最好的可视化和分析产品，而不必担心锁定。

In addition to Honeycomb and Lightstep, some of the largest vendors in the monitoring field, as well as the largest end-users are participating, including Google, Microsoft, Splunk, Postmates, and Uber.

除了Honeycomb和Lightstep，一些监控领域最大的厂商，以及最大的终端用户也在参与，包括谷歌、微软、Splunk、Postmates和Uber。

### New Collectors and SDK 

### 新的收集器和 SDK

The new collector is crucial, explained Honeycomb’s Fong-Jones, in that it narrows the minimum scope of what vendors must support in order to ingest telemetry. “It enables vendors to write one exporter in Golang and have _any_ [Jaeger](https://www.jaegertracing.io/),[Zipkin](https://zipkin.io/), or [OTLP](https://github.com/open-telemetry/oteps/blob/master/text/0035-opentelemetry-protocol.md) (OpenTelemetry's new telemetry protocol) producer transmit to that vendor via the collector.”

Honeycomb 的 Fong-Jones 解释说，新的收集器至关重要，因为它缩小了供应商必须支持的最小范围，以便摄取遥测数据。 “它使供应商能够在 Golang 中编写一个导出器并拥有 _any_ [Jaeger](https://www.jaegertracing.io/)、[Zipkin](https://zipkin.io/) 或 [OTLP](https: //github.com/open-telemetry/oteps/blob/master/text/0035-opentelemetry-protocol.md）（OpenTelemetry 的新遥测协议)生产者通过收集器传输给该供应商。”

![](https://cdn.thenewstack.io/media/2020/04/0d6104d3-opentelemetry.jpg)

OpenTelemetry Collector is the ‘swiss army knife’ for critical telemetry data that can be used downstream in a number of tools including Honeycomb for production system Observability and improved resilience.

OpenTelemetry Collector 是关键遥测数据的“瑞士军刀”，可在下游用于许多工具，包括用于生产系统可观察性和改进弹性的 Honeycomb。

The Honeycomb software itself, which originally supported both OpenCensus and OpenTracing, can now be streamlined by using only OpenTelemetry to format and deliver into Honeycomb.

最初支持 OpenCensus 和 OpenTracing 的 Honeycomb 软件本身现在可以通过仅使用 OpenTelemetry 进行格式化并交付到 Honeycomb 中进行简化。

The SDKs are also vital in that they allow developers to write instrumentation in their own favored languages, while still adhering to the OpenTelemetry API specifications, Fong-Jones explained. Languages currently supported include Erlang, GoLang, Java, JavaScript,  and Python.

Fong-Jones 解释说，SDK 也很重要，因为它们允许开发人员用他们自己喜欢的语言编写检测，同时仍然遵守 OpenTelemetry API 规范。当前支持的语言包括 Erlang、GoLang、Java、JavaScript 和 Python。

[Open Keynote: (Open)Telemetry Makes Observability Simple – Sarah Novotny & Liz Fong-Jones on YouTube.](https://www.youtube.com/watch?v=W_8MHdtrgZE&list=WL&index=2&t=0s)

[开放式主题演讲：（开放式）遥测使可观察性变得简单——YouTube 上的 Sarah Novotny 和 Liz Fong-Jones。](https://www.youtube.com/watch?v=W_8MHdtrgZE&list=WL&index=2&t=0s)

[Open OpenTelemetry: Overview & Backwards Compatibility of OpenTracing + OpenCensus – Steve Flanders on YouTube.](https://www.youtube.com/watch?v=UdRqts403G4&list=WL&index=3&t=12s)

[Open OpenTelemetry：OpenTracing + OpenCensus 的概述和向后兼容性 – YouTube 上的 Steve Flanders。](https://www.youtube.com/watch?v=UdRqts403G4&list=WL&index=3&t=12s)

[Open Beyond Getting Started: Using OpenTelemetry to Its Full Potential – Sergey Kanzhelev & Morgan McLean on YouTube.](https://www.youtube.com/watch?v=FlghuHDlQdM&list=WL&index=4&t=51s)

[Open Beyond Getting Started：充分发挥 OpenTelemetry 的潜力——YouTube 上的 Sergey Kanzhelev 和 Morgan McLean。](https://www.youtube.com/watch?v=FlghuHDlQdM&list=WL&index=4&t=51s)

Lightstep is a sponsor of The New Stack. 

Lightstep 是 The New Stack 的赞助商。

