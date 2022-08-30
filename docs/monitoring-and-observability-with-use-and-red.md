# Monitoring and Observability With USE and RED

https://orangematter.solarwinds.com/2017/10/05/monitoring-and-observability-with-use-and-red/


October 5, 2017


Modern systems can emit thousands or millions of metrics, and modern monitoring tools can collect them all. Faced with such an abundance of data, it can be difficult to know where to start looking when trying to diagnose a problem. And when you’re not in diagnosis mode, but you just want to know whether there’s a problem at all, you might have the same difficulty. What are the most important KPIs coming from your systems?I’ve written extensively about this before, but this time I want to refer to acronyms—USE and RED respectively—that are easy to remember and provide good high-level guidance for system observability.

## The USE Method

USE is an acronym for [Utilization, Saturation, and Errors](http://www.brendangregg.com/usemethod.html). Brendan Gregg suggests using it to get started quickly when you’re diving into an unknown system: “I developed the USE method to teach others how to solve common performance issues quickly, without overlooking important areas. Like an emergency checklist in a flight manual, it is intended to be simple, straightforward, complete, and fast.”A summary of USE is “For every resource, check utilization, saturation, and errors.” What do those things mean? Brendan defines the terminology:

- Utilization: the average time the resource was busy servicing work
- Saturation: the degree to which the resource has extra work which it can't service, often queued
- Errors: the count of error events

This disambiguates utilization and saturation, making it clear utilization is “busy time %” and saturation is “backlog.” These terms are very different from things a person might confuse with them, such as “disk utilization” as an expression of how much disk space is left.

## The RED Method

I first saw this acronym in a [talk on monitoring microservices](https://www.slideshare.net/weaveworks/monitoring-microservices) in 2015. The acronym stands for Rate, Errors, and Duration. These are request-scoped, not resource-scoped as the USE method is. Duration is explicitly taken to mean _distributions_, not averages.

## USE and RED: Two Sides of the Same Coin

What may not be obvious is USE and RED complement one another. The USE method is an internal, service-centric view. The system or service’s workload is assumed, and USE directs attention to the _resources_ handling the workload. The goal is to understand how these resources behave in the presence of the load.The RED method, on the other hand, is about the _workload_ itself, and treats the service as a black box. It’s an externally-visible view of the behavior of the workload as serviced by the resources. I define _workload_ as a population of requests over a period of time. I’ve spoken and written extensively before about the importance of measuring the workload, since the system’s _raison d’etre_ is to do useful work.Taken together, RED and USE comprise minimally complete, maximally useful observability—a way to understand _both_aspects of a system: its users/customers and the work they request, and its resources/components and how they react to the workload. (I include users in the system. Users aren’t separate from the system; they’re an inextricable part of it.)I often refer to this duality as the "Zen of Performance," a holistic, unified system performance worldview I'm developing. It's work in progress!

## Mapping USE and RED to Standard Terminology

USE and RED are convenient, and part of the reason they’re so valuable is their atoms map directly to standard concepts that are core performance metrics:

- U = Utilization, as canonically defined
- S = Concurrency
- E = Error Rate, as a throughput metric
- R = Request Throughput, in requests per second
- E = Request Error Rate, as either a throughput metric or a fraction of overall throughput
- D = Latency, Residence Time, or Response Time; all three are widely used

To learn more about why these metrics are so fundamental to performance and observability, listen to Jon Moore’s talk on why [API admission control](https://www.youtube.com/watch?v=m64SWl9bfvk) should use concurrency instead of throughput. And, for further reading, consider my eBooks on [queuing theory](https://www.vividcortex.com/resources/queueing-theory) and the [Universal Scalability Law](https://www.vividcortex.com/resources/universal-scalability-law/). In conclusion, if you’re unsure which metrics are most useful for both monitoring and diagnosis, USE and RED are great places to start.
