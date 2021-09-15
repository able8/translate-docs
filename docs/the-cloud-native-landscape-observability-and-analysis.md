# TheCloudNativeLandscape:ObservabilityandAnalysis

#### 17 May 2021 1:24pm,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack.io/author/jason-morgan/ "Posts by Jason Morgan")


This post is part of an ongoing series from [Cloud Native Computing Foundation Business Value Subcommittee](https://lists.cncf.io/g/cncf-business-value) co-chairs [Catherine Paganini](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) and [Jason Morgan](https://thenewstack.io/author/jason-morgan/) that focuses on explaining each category of the cloud native landscape to a non-technical audience as well as engineers just getting started with cloud native computing.



We finally arrived at the last section of our [Cloud Native Computing Foundation](https://cncf.io/?utm_content=inline-mention)‘s Landscape series. If you missed our previous articles, we covered an [introduction](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/), and then the [provisioning](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/), [runtime](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/), [orchestration and management layer](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/), and [platforms](https://thenewstack.io/the-cloud-native-landscape-platforms-explained/) each in a separate article. Today, we’ll discuss each category of the observability and analysis “column.”

Let’s start by defining observability and analysis. Observability is a system characteristic describing the degree to which a system can be understood from its external outputs. Measured by CPU time, memory, disk space, latency, errors, etc., computer systems can be more or less observable. Analysis, on the other hand, is an activity in which you look at this observable data and make sense of it.


To ensure there is no service disruption, you’ll need to observe and analyze every aspect of your application so any anomaly gets detected and rectified right away. This is what this category is all about. It runs across and observes all layers which is why it’s on the side and not embedded in a specific layer.

Tools in this category are broken down into logging, monitoring, tracing, and chaos engineering. Please note that the category name is somewhat misleading. While listed here, chaos engineering is rather a reliability than an observability or analysis tool.

## **Logging**

### What It Is

Applications emit a steady stream of log messages describing what they are doing at any given time. These log messages capture various events happening in the system such as failed or successful actions, audit information, or health events. Logging tools collect, store, and analyze these messages to track error reports and related data. Along with metrics and tracing, logging is one of the pillars of observability.

### Problem It Addresses

Collecting, storing, and analyzing logs is a crucial part of building a modern platform. Logging helps with, performs one, or all of those tasks. Some tools handle every aspect from collection to analysis while others focus on a single task like collection. All logging tools aim at helping organizations gain control over their log messages.

### How It Helps

When collecting, storing, and analyzing application log messages, you’ll understand what an application was communicating at any given time. But note, logs represent messages that applications can deliberately emit, they won’t necessarily pinpoint the root cause of a given issue. That being said, collecting and retaining log messages over time is an extremely powerful capability and will help teams diagnose issues and meet regulatory and compliance requirements.

### Technical 101

While collecting, storing, and processing log messages is by no means a new problem, cloud native patterns and Kubernetes have caused significant changes in the way we handle logs. Traditional approaches to logging that were appropriate for virtual and physical machines, like writing logs to a file, are ill-suited to containerized applications, where the file system doesn’t outlast an application. In a cloud native environment log collection tools like Fluentd, run alongside application containers and collect messages directly from the applications. Messages are then forwarded on to a central log store to be aggregated and analyzed.

Fluentd is the only CNCF project in this space.

**Buzzwords****Popular Projects**

- Logging

- Fluentd & Fluentbit
- Elastic Logstash

![logging ](https://cdn.thenewstack.io/media/2021/05/61f23a99-screen-shot-2021-05-13-at-8.17.44-am.png)

## Monitoring

### What It Is

Monitoring refers to instrumenting an app to collect, aggregate, and analyze logs and metrics to improve our understanding of its behavior. While logs describe specific events, metrics are a measurement of a system at a given point in time — they are two different things but are both necessary to get the full picture of your system’s health. Monitoring includes everything from watching disk space, CPU usage, and memory consumption on individual nodes to doing detailed synthetic transactions to see if a system or application is responding correctly and in a timely manner. There are a number of different approaches to monitoring systems and applications.

### Problem It Addresses

When running an application or platform, you want it to accomplish a specific task as designed and ensure it’s only accessed by authorized users. Monitoring allows you to know if it is working correctly, securely, cost-effectively, only accessed by authorized users, and/or any other characteristic you may be tracking.

### How It Helps

Good monitoring allows operators to respond quickly, and potentially automatically when an incident arises. It provides insights into the current health of a system and watches for changes. Monitoring tracks everything from application health to user behavior and is an essential part of effectively running applications.

### Technical 101

Monitoring in a cloud native context is generally similar to monitoring traditional applications. You need to track metrics, logs, and events to understand the health of your applications. The main difference is that some of the managed objects are ephemeral, meaning they may not be long-lasting so tying your monitoring to auto-generated resource names won’t be a good long-term strategy. There are a number of CNCF projects in this space that largely revolve around Prometheus, the CNCF graduated project.

**Buzzwords**

**Popular Projects/Products**

- Monitoring
- Time series
- Alerting
- Metrics

- Prometheus
- Cortex
- Thanos
- Grafana

![monitoring ](https://cdn.thenewstack.io/media/2021/05/dc4db049-screen-shot-2021-05-13-at-8.16.16-am.png)

## Tracing

### What It Is

In a microservices world, services are constantly communicating with each other over the network. Tracing, a specialized use of logging, allows you to trace the path of a request as it moves through a distributed system.

### Problem It Addresses

Understanding how a microservice application behaves at any given point in time is an extremely challenging task. While many tools provide deep insights into service behavior, it can be difficult to tie an action of an individual service to the broader understanding of how the entire app behaves.

### How It Helps

Tracing solves this problem by adding a unique identifier to messages sent by the application. That unique identifier allows you to follow, or trace, individual transactions as they move through your system. You can use this information to see both the health of your application as well as to debug problematic microservices or activities.

### Technical 101

Tracing is a powerful debugging tool that allows you to troubleshoot and fine-tune the behavior of a distributed application. That power does come at a cost. Application code needs to be modified to emit tracing data and any spans need to be propagated by infrastructure components in the data path of your application. Specifically service meshes’ and their proxies. Jaeger and Open Tracing are CNCF projects in this space.

**Buzzwords**

**Popular Projects**

- Span
- Tracing

- Jaeger
- OpenTracing

![Tracing](https://cdn.thenewstack.io/media/2021/05/ea2aaf06-screen-shot-2021-05-13-at-8.28.38-am.png)

## **Chaos Engineering**

### What It Is

Chaos engineering refers to the practice of intentionally introducing faults into a system in order to create more resilient applications and engineering teams. A chaos engineering tool will provide a controlled way to introduce faults and run specific experiments against a particular instance of an application.

### Problem It Addresses

Complex systems fail. They fail for a host of reasons, and the consequences in a distributed system are typically hard to understand. Chaos engineering is embraced by organizations that accept that failures will occur and, instead of trying to prevent failures, practice recovering from them. This is referred to as optimizing for [mean time to repair](https://en.wikipedia.org/wiki/Mean_time_to_repair), or MTTR.

**Side note**: The traditional approach to maintaining high availability for applications is referred to as optimizing for [mean time between failures](https://en.wikipedia.org/wiki/Mean_time_between_failures), or MTBF. You can observe this practice in organizations that use things like “change review boards” and long change freezes” to keep an application environment stable by restricting changes. The authors of [Accelerate](https://itrevolution.com/book/accelerate/) suggest that high-performing IT organizations achieve high availability by optimizing for mean time to recovery, or MTTR, instead.

### How It Helps

In a cloud native world, applications must dynamically adjust to failures — a relatively new concept. That means, when something fails, the system doesn’t go down completely but gracefully degrades or recovers. Chaos engineering tools enable you to experiment on a software system in production to ensure they do that should a real failure occur.

In short, you experiment with a system because you want to be confident that it can withstand turbulent and unexpected conditions. Instead of waiting for something to happen and find out, you put it through duress under controlled conditions to identify weaknesses and fix them before chance uncovers them.

### Technical 101

Chaos engineering tools and practices are critical to achieving high availability for your applications. Distributed systems are often too complex to be fully understood by any one engineer and no change process can fully predetermine the impact of changes on an environment. By introducing deliberate chaos engineering practices teams are able to practice, and automate, recovering from failures. Chaos Mesh and Litmus Chaos are CNCF tools in this space but there are many open source and proprietary options available.

**Buzzwords**

**Popular Projects**

- Chaos engineering

- Chaos Mesh
- Litmus Chaos

![Chaos engineering ](https://cdn.thenewstack.io/media/2021/05/5463852b-screen-shot-2021-05-13-at-8.32.22-am.png)

As we’ve seen, the observability and analysis layer is all about understanding the health of your system and ensuring it stays operational even under tough conditions. Logging tools capture event messages emitted by apps, monitoring watches logs and metrics, and tracing follows the path of individual requests. When combined, these tools ideally provide a 360-degree view of what’s going on within your system. Chaos engineering is a little different. It provides a safe way to verify the system can withstand unexpected events, basically ensuring it stays healthy.

_This finally concludes our series on the CNCF landscape. We certainly learned a lot while writing these articles and hope you did too._

Feature image [via](https://pixabay.com/fr/photos/jumelles-%C3%A0-la-recherche-l-homme-1209011/) Pixabay.


The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Real.

This post is part of a larger story we're telling about Observability.

[Notify me when ebook is available](http://thenewstack.io/ebooks/observability/cloud-native-observability-for-devops-teams/)

[Notify me when ebook is available](http://thenewstack.io/ebooks/observability/cloud-native-observability-for-devops-teams/)

[Cloud Native Observability for DevOps Teams](https://thenewstack.io/tag/cloud-native-observability-for-devops-teams/) [Contributed](https://thenewstack.io/tag/contributed/)
