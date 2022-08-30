# Making the USE method of monitoring useful

https://www.infoworld.com/article/3638772/making-the-use-method-of-monitoring-useful.html


### When performance issues arise, checking the USE metrics—utilization, saturation, and errors—can help you identify system bottlenecks.


Errors happen. Things will go wrong. It’s not a matter of if — it’s a matter of when. But if we understand that fact ahead of time, we can take steps to prepare for the inevitable. Having a way to quickly identify contributing factors means we can address them faster. That translates into less downtime, which makes everyone happier.

However, knowing that you need to prepare for problems isn’t the same as having a strategy for how to identify them. If you want to be able to quickly and systematically rule things out, then you need to know what those things are, as well as their acceptable thresholds.

## **The USE method**

Think of the USE method like an emergency checklist for all your critical resources. For every resource in the list, check for one or more of:

- Utilization
- Saturation
- Errors

When performance issues arise, the USE method can help identify system bottlenecks.

First, let’s define a resource. In this case, a resource is any functional server component. This can be physical elements, such as disks, CPUs, network connections, or buses, as well as certain software components.

The three USE criteria can mean different things depending on context. Let’s define them for the USE method.

- **Utilization:** The average time that the resource was busy serving work. We usually represent utilization as a percentage over time.
- **Saturation:** The amount of work that a resource is unable to service. We usually represent this metric as a queue length.
- **Errors:** The number of error events. We usually represent errors as a total count.

It’s important to remember that utilization and saturation are time series metrics, so it may take some trial and error to find the optimal monitoring interval.

For example, a long time interval can display high saturation levels with low utilization levels. Shortening the time interval can reveal utilization spikes. You may want to dashboard a few different time intervals to get a clearer picture of performance trends.

The above example also illustrates the value of high-performance time series data stores. InfluxDB, for example, allows you to ingest high granularity data and to slice and dice in multiple ways. This capability allows you to simultaneously answer multiple questions about the same aspects of the system.

## Create a checklist

Think about all the different resources that your system uses and how you want to measure them. Some resources can cause bottlenecks in more than one way. For example, a network interconnect might have I/O issues as well as CPU performance issues. You want to be sure to create a separate item for each type of issue to make the identification process more thorough and faster.

The USE method works best on resources that experience performance degradation under heavy usage. It doesn’t work well on resources that utilize caching because caching improves resource performance under heavy usage.

## Build a monitoring system

As the last point about caching indicates, the USE method is not a cure-all. To get the most out of the USE method, combine it with other monitoring methods and processes. Be prepared to put a good chunk of time into planning and optimizing your monitoring system.

This is the approach we take at InfluxData:

1. Before we configure any dashboards we work to understand our thresholds, as measured by our service level indicators (SLIs). This is a critical step because it allows us to avoid alert fatigue by filtering out metrics outside the established thresholds. At the same time, having a threshold enables us to track issues as they emerge, rather than having them pop up out of the blue. The more we understand about our systems from the perspective of acceptable performance and expected scale, the more predictable our pain points become. In other words, we use data to try to figure out how to avoid throwing alerts in the first place.
2. We set up alerts so that we know about issues right away.
3. We built USE and RED dashboards and use them as input to our SLIs to see if the alert indicates a current or potential problem. These dashboards also function as a troubleshooting tool to pinpoint factors that may contribute to an outage or incident.
4. We use SLO dashboards to gauge availability and to help determine if we need to invest in more availability or features to correct the problem.
5. Finally, we created a wide range of custom dashboards that we use to investigate and diagnose issues if the USE/RED dashboards indicate a valid issue.

InfluxDataInfluxData

The goal is to catch problems early and solve them before they affect users or system performance. If nothing else, hopefully our system illustrates how to think about performance issues and the interconnectedness of different monitoring methods.

