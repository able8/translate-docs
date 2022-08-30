# Making the USE method of monitoring useful

# 使监控的USE方法有用

https://www.infoworld.com/article/3638772/making-the-use-method-of-monitoring-useful.html

### When performance issues arise, checking the USE metrics—utilization, saturation, and errors—can help you identify system bottlenecks.

### 当出现性能问题时，检查 USE 指标（利用率、饱和度和错误）可以帮助您识别系统瓶颈。

Errors happen. Things will go wrong. It’s not a matter of if — it’s a matter of when. But if we understand that fact ahead of time, we can take steps to prepare for the inevitable. Having a way to quickly identify contributing factors means we can address them faster. That translates into less downtime, which makes everyone happier.

发生错误。事情会出错。这不是是否的问题——而是何时的问题。但是，如果我们提前了解这一事实，我们就可以采取措施为不可避免的情况做好准备。有一种方法可以快速识别促成因素，这意味着我们可以更快地解决它们。这意味着更少的停机时间，这让每个人都更快乐。

However, knowing that you need to prepare for problems isn’t the same as having a strategy for how to identify them. If you want to be able to quickly and systematically rule things out, then you need to know what those things are, as well as their acceptable thresholds.

但是，知道您需要为问题做好准备并不等于制定了识别问题的策略。如果您希望能够快速系统地排除某些事情，那么您需要知道这些事情是什么，以及它们的可接受阈值。

## **The USE method**

## **使用方法**

Think of the USE method like an emergency checklist for all your critical resources. For every resource in the list, check for one or more of:

将 USE 方法想象为所有关键资源的紧急清单。对于列表中的每个资源，请检查以下一项或多项：

- Utilization
- Saturation
- Errors

- 利用
- 饱和度
- 错误

When performance issues arise, the USE method can help identify system bottlenecks.

当出现性能问题时，USE 方法可以帮助识别系统瓶颈。

First, let’s define a resource. In this case, a resource is any functional server component. This can be physical elements, such as disks, CPUs, network connections, or buses, as well as certain software components.

首先，让我们定义一个资源。在这种情况下，资源是任何功能性服务器组件。这可以是物理元素，例如磁盘、CPU、网络连接或总线，以及某些软件组件。

The three USE criteria can mean different things depending on context. Let’s define them for the USE method.

这三个 USE 标准可能意味着不同的事情，具体取决于上下文。让我们为 USE 方法定义它们。

- **Utilization:** The average time that the resource was busy serving work. We usually represent utilization as a percentage over time.
- **Saturation:** The amount of work that a resource is unable to service. We usually represent this metric as a queue length.
- **Errors:** The number of error events. We usually represent errors as a total count.

- **利用率：**资源忙于服务工作的平均时间。我们通常将利用率表示为一段时间内的百分比。
- **饱和度：**资源无法服务的工作量。我们通常将此度量表示为队列长度。
- **错误：**错误事件的数量。我们通常将错误表示为总计数。

It’s important to remember that utilization and saturation are time series metrics, so it may take some trial and error to find the optimal monitoring interval.

请务必记住，利用率和饱和度是时间序列指标，因此可能需要反复试验才能找到最佳监控间隔。

For example, a long time interval can display high saturation levels with low utilization levels. Shortening the time interval can reveal utilization spikes. You may want to dashboard a few different time intervals to get a clearer picture of performance trends.

例如，较长的时间间隔可以显示高饱和度水平和低利用率水平。缩短时间间隔可以揭示利用率峰值。您可能需要仪表板几个不同的时间间隔，以更清楚地了解性能趋势。

The above example also illustrates the value of high-performance time series data stores. InfluxDB, for example, allows you to ingest high granularity data and to slice and dice in multiple ways. This capability allows you to simultaneously answer multiple questions about the same aspects of the system.

上面的示例还说明了高性能时间序列数据存储的价值。例如，InfluxDB 允许您摄取高粒度数据并以多种方式进行切片和切块。此功能允许您同时回答有关系统相同方面的多个问题。

## Create a checklist

## 创建一个清单

Think about all the different resources that your system uses and how you want to measure them. Some resources can cause bottlenecks in more than one way. For example, a network interconnect might have I/O issues as well as CPU performance issues. You want to be sure to create a separate item for each type of issue to make the identification process more thorough and faster.

考虑一下您的系统使用的所有不同资源以及您希望如何衡量它们。某些资源可能以不止一种方式造成瓶颈。例如，网络互连可能存在 I/O 问题以及 CPU 性能问题。您要确保为每种类型的问题创建一个单独的项目，以使识别过程更加彻底和快速。

The USE method works best on resources that experience performance degradation under heavy usage. It doesn’t work well on resources that utilize caching because caching improves resource performance under heavy usage.

USE 方法最适用于在大量使用下性能下降的资源。它不适用于利用缓存的资源，因为缓存可以提高资源在大量使用下的性能。

## Build a monitoring system

## 搭建监控系统

As the last point about caching indicates, the USE method is not a cure-all. To get the most out of the USE method, combine it with other monitoring methods and processes. Be prepared to put a good chunk of time into planning and optimizing your monitoring system.

正如关于缓存的最后一点表明的那样，USE 方法并不是万能的。要充分利用 USE 方法，请将其与其他监控方法和流程结合使用。准备好花大量时间来规划和优化您的监控系统。

This is the approach we take at InfluxData: 

这是我们在 InfluxData 采用的方法：

1. Before we configure any dashboards we work to understand our thresholds, as measured by our service level indicators (SLIs). This is a critical step because it allows us to avoid alert fatigue by filtering out metrics outside the established thresholds. At the same time, having a threshold enables us to track issues as they emerge, rather than having them pop up out of the blue. The more we understand about our systems from the perspective of acceptable performance and expected scale, the more predictable our pain points become. In other words, we use data to try to figure out how to avoid throwing alerts in the first place.
2. We set up alerts so that we know about issues right away.
3. We built USE and RED dashboards and use them as input to our SLIs to see if the alert indicates a current or potential problem. These dashboards also function as a troubleshooting tool to pinpoint factors that may contribute to an outage or incident.
4. We use SLO dashboards to gauge availability and to help determine if we need to invest in more availability or features to correct the problem.
5. Finally, we created a wide range of custom dashboards that we use to investigate and diagnose issues if the USE/RED dashboards indicate a valid issue.



1. 在我们配置任何仪表板之前，我们会努力了解我们的阈值，由我们的服务水平指标 (SLI) 衡量。这是一个关键步骤，因为它允许我们通过过滤掉既定阈值之外的指标来避免警报疲劳。同时，设置阈值使我们能够在问题出现时对其进行跟踪，而不是让它们突然出现。我们从可接受的性能和预期规模的角度对我们的系统了解得越多，我们的痛点就越容易预测。换句话说，我们首先使用数据来试图弄清楚如何避免引发警报。
2. 我们设置警报，以便我们立即了解问题。
3. 我们构建了 USE 和 RED 仪表板，并将它们用作 SLI 的输入，以查看警报是否指示当前或潜在问题。这些仪表板还可以用作故障排除工具，以查明可能导致中断或事件的因素。
4. 我们使用 SLO 仪表板来衡量可用性并帮助确定我们是否需要投资更多的可用性或功能来解决问题。
5. 最后，我们创建了广泛的自定义仪表板，如果 USE/RED 仪表板表明存在有效问题，我们将使用这些仪表板来调查和诊断问题。



The goal is to catch problems early and solve them before they affect users or system performance. If nothing else, hopefully our system illustrates how to think about performance issues and the interconnectedness of different monitoring methods. 

目标是尽早发现问题并在它们影响用户或系统性能之前解决它们。如果不出意外，希望我们的系统能够说明如何考虑性能问题以及不同监控方法的相互联系。

