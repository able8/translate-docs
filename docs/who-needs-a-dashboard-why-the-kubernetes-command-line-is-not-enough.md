# Who Needs a Dashboard? Why the Kubernetes Command Line Is Not Enough

#### 16 Oct 2020 12:26pm,   by [Emily Omier](https://thenewstack.io/author/emily-omier/ "Posts by Emily Omier")

![](https://cdn.thenewstack.io/media/2020/10/4556424b-benjamin-child-igqmknl6lne-unsplash.jpg)

Developers don’t use them, but enterprises won’t buy a product without them. What’s the deal with dashboards? And why do they matter for Kubernetes?

Dashboards make things easier, and for that reason are sometimes looked down on by engineers, would think that the best way to interact with Kubernetes is through the command line. But dashboards are extremely important to technology buyers because they address some of the realities about using Kubernetes: That not everyone in their organization will be Kubernetes experts, that a central team will need some control and that instantaneous access to information is essential during an incident.

## **Development vs. Operations**

The way that people use dashboards — if they use them at all — depends on their skill sets and what they are trying to accomplish. “If I’m using Kubernetes, most likely I will use the command line, that will be easiest and fastest,” explained [Idit Levine](https://www.linkedin.com/in/iditlevine/), founder and CEO of API infrastructure software provider [Solo.io](https://www.solo.io/). This goes, she says, for most developers working on creating and deploying applications. In most cases, dashboards aren’t particularly helpful at this stage.

But that all changes once the application is in production. Dashboards make it easier to see a large amount of data — they also make it easier to see when something is wrong. “We mainly see if for debugging,” Levine said. “You want to see when something is wrong. You want to see a red dot.”

Dashboards make it easier for humans to process data quickly. Human-readable data is especially important as part of the operational story — if you want to recover from incidents quickly, every second that an operator is running CLI commands instead of seeing a red dot or green dot on the dashboard is critical.

## **Complexity**

Kubernetes is complex, and relying on the command line and raw data feeds can become so overwhelming that even an experienced Kubernetes engineer could easily miss something important. “There’s so much data being thrown off a Kubernetes cluster that a dashboard to ingest that data, filter it, prioritize it and put it in front of a user in a sensible way, that’s easy to just quickly grok, you know, is there a problem that I need to focus on, is there something that doesn’t look right?” explained [Robert Brennan](https://www.linkedin.com/in/robert-a-brennan/), director of open source software at Kubernetes service provider [Fairwinds](https://www.fairwinds.com/).

Dashboards can become particularly important as Kubernetes scales. The more clusters there are to manage, the more applications are running, the harder it is to extract the information from raw data feeds or from the command line.

“Beyond surfacing data in a pretty UI, dashboards can also take on things like alerting and policy configuration, so you get a Slack message every time a particular issue arises,” Brennan said. Particularly when talking about the needs of large organizations running mission-critical applications on Kubernetes, dashboards can provide the centralized control and alerting that are essential to long-term success with Kubernetes.

“When we built our Kubernetes solution, we started with understanding the Kubernetes structure natively and then creating a set of dashboards that allow you to unpeel the onion in a thoughtful way,” explained [Kalyan Ramanathan](https://www.linkedin.com/in/kalyanramanathan/), vice president of product marketing at [observability](https://www.sumologic.com/) company Sumo Logic. Users need to be able to immediately pinpoint the relevant information, especially in a debugging scenario. “There are seven layers of Kubernetes, and I can see metrics and logs across all seven layers. But what am I supposed to do with this?” Ramanathan said. The goal of most Kubernetes dashboards is to connect the dots for users, so not only information but also the logical next course of action is obvious.

Dashboards can also provide a way to correlate signals, making it easier to understand relationships between the different logging metrics and the different parts of the system and ultimately easier to troubleshoot during incidents.

## **More than Pure Ops**

One of the challenges with a technology as complex as Kubernetes is that one dashboard isn’t always enough. One dashboard might be enough for an incident response use case (though it often isn’t, and dashboard proliferation is its own problem), but organizations also need a way to track things like security and cost. It’s hard to spot the one container running as root, for example, without some kind of dashboard that is collecting that data and alerting users. The same goes for something like costs: Cloud costs are notoriously byzantine. Without a dedicated dashboard to sort through the information, the chances of anyone truly understanding why the bill doubled last month are pretty low.

Whether it’s related to security, cost or something else, dashboards are also a way to create and enforce policies or put the “guardrails” on developers. Without dashboards, getting centralized control of Kubernetes is near impossible, leaving individual developers responsible for correct configurations every time. This doesn’t matter much to a developer creating a pet project on the side, but it matters a lot to an organization in a regulated industry with a team of 800 developers with varying skill levels and hundreds of microservices.

## **Maturity Markers**

So what do dashboards ultimately say about a project, and what do they mean for the community and for adoption in general? “What I see in my market is that a lot of the time if there is a dashboard it shows about the maturity of the product,” Levine said. “If you have a beautiful UI, usually, in this market, then it’s not just an open source project that two people are trying to spin up. If you’re thinking about the user interface, it means there is a certain amount of maturity.”

It also is a requirement for widespread adoption. “When I started the company and it was just open source, I was talking to a lot of customers and showing them demos with the command line,” Levine said about Gloo, the open source ingress controller Solo.io produces. “To be honest, it was hard to impress them. Then we took exactly the same project and we put on a beautiful UI. Suddenly, I gave them the same demo and everyone said whoa, yeah, we want to proceed.”

Feature image by [Benjamin Child](https://unsplash.com/@bchild311?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/dashboard?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText).
