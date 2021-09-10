# Who Needs a Dashboard? Why the Kubernetes Command Line Is Not Enough

# 谁需要仪表板？为什么 Kubernetes 命令行不够用

#### 16 Oct 2020 12:26pm,   by [Emily Omier](https://thenewstack.io/author/emily-omier/ "Posts by Emily Omier")



![](https://cdn.thenewstack.io/media/2020/10/4556424b-benjamin-child-igqmknl6lne-unsplash.jpg)

Developers don’t use them, but enterprises won’t buy a product without them. What’s the deal with dashboards? And why do they matter for Kubernetes?

开发人员不使用它们，但企业不会购买没有它们的产品。与仪表板有什么关系？为什么它们对 Kubernetes 很重要？

Dashboards make things easier, and for that reason are sometimes looked down on by engineers, would think that the best way to interact with Kubernetes is through the command line. But dashboards are extremely important to technology buyers because they address some of the realities about using Kubernetes: That not everyone in their organization will be Kubernetes experts, that a central team will need some control and that instantaneous access to information is essential during an incident.

仪表板使事情变得更容易，因此有时被工程师看不起，会认为与 Kubernetes 交互的最佳方式是通过命令行。但是仪表板对于技术购买者来说非常重要，因为它们解决了使用 Kubernetes 的一些现实问题：并非组织中的每个人都是 Kubernetes 专家，中央团队需要一些控制，并且在发生事件时即时访问信息是必不可少的。

## **Development vs. Operations**

## **开发与运营**

The way that people use dashboards — if they use them at all — depends on their skill sets and what they are trying to accomplish. “If I'm using Kubernetes, most likely I will use the command line, that will be easiest and fastest,” explained [Idit Levine](https://www.linkedin.com/in/iditlevine/), founder and CEO of API infrastructure software provider [Solo.io](https://www.solo.io/). This goes, she says, for most developers working on creating and deploying applications. In most cases, dashboards aren’t particularly helpful at this stage.

人们使用仪表板的方式——如果他们使用它们——取决于他们的技能和他们想要完成的任务。 “如果我使用 Kubernetes，我很可能会使用命令行，这将是最简单和最快的，”创始人兼首席执行官 [Idit Levine](https://www.linkedin.com/in/iditlevine/) 解释说API 基础设施软件提供商 [Solo.io](https://www.solo.io/)。她说，这适用于大多数致力于创建和部署应用程序的开发人员。在大多数情况下，仪表板在这个阶段并不是特别有用。

But that all changes once the application is in production. Dashboards make it easier to see a large amount of data — they also make it easier to see when something is wrong. “We mainly see if for debugging,” Levine said. “You want to see when something is wrong. You want to see a red dot.”

但是一旦应用程序投入生产，这一切都会改变。仪表板可以更轻松地查看大量数据——它们还可以更轻松地查看出现问题时的情况。 “我们主要看是否用于调试，”莱文说。 “你想看看什么时候出了问题。你想看到一个红点。”

Dashboards make it easier for humans to process data quickly. Human-readable data is especially important as part of the operational story — if you want to recover from incidents quickly, every second that an operator is running CLI commands instead of seeing a red dot or green dot on the dashboard is critical.

仪表板使人们可以更轻松地快速处理数据。作为操作故事的一部分，人类可读的数据尤其重要——如果您想快速从事件中恢复，操作员运行 CLI 命令而不是在仪表板上看到红点或绿点的每一秒都是至关重要的。

## **Complexity**

## **复杂性**

Kubernetes is complex, and relying on the command line and raw data feeds can become so overwhelming that even an experienced Kubernetes engineer could easily miss something important. “There's so much data being thrown off a Kubernetes cluster that a dashboard to ingest that data, filter it, prioritize it and put it in front of a user in a sensible way, that's easy to just quickly grok, you know, is there a problem that I need to focus on, is there something that doesn't look right?” explained [Robert Brennan](https://www.linkedin.com/in/robert-a-brennan/), director of open source software at Kubernetes service provider [Fairwinds](https://www.fairwinds.com/) .

Kubernetes 很复杂，对命令行和原始数据源的依赖会变得如此不堪重负，以至于即使是经验丰富的 Kubernetes 工程师也很容易错过一些重要的东西。 “有太多的数据从 Kubernetes 集群中被抛出，以至于一个仪表板来摄取这些数据、过滤它、确定优先级并以一种明智的方式把它放在用户面前，这很容易快速理解，你知道，有没有我需要关注的问题，是不是有什么地方不对劲？”解释 [Robert Brennan](https://www.linkedin.com/in/robert-a-brennan/)，Kubernetes 服务提供商 [Fairwinds](https://www.fairwinds.com/) 开源软件总监.

Dashboards can become particularly important as Kubernetes scales. The more clusters there are to manage, the more applications are running, the harder it is to extract the information from raw data feeds or from the command line.

随着 Kubernetes 的扩展，仪表板会变得尤为重要。需要管理的集群越多，运行的应用程序越多，从原始数据源或命令行中提取信息就越困难。

“Beyond surfacing data in a pretty UI, dashboards can also take on things like alerting and policy configuration, so you get a Slack message every time a particular issue arises,” Brennan said. Particularly when talking about the needs of large organizations running mission-critical applications on Kubernetes, dashboards can provide the centralized control and alerting that are essential to long-term success with Kubernetes. 

“除了在漂亮的 UI 中显示数据之外，仪表板还可以执行警报和策略配置等操作，因此每次出现特定问题时，您都会收到 Slack 消息，”布伦南说。特别是在讨论在 Kubernetes 上运行关键任务应用程序的大型组织的需求时，仪表板可以提供集中控制和警报，这对于 Kubernetes 的长期成功至关重要。

“When we built our Kubernetes solution, we started with understanding the Kubernetes structure natively and then creating a set of dashboards that allow you to unpeel the onion in a thoughtful way,” explained [Kalyan Ramanathan](https://www.linkedin.com/in/kalyanramanathan/), vice president of product marketing at [observability](https://www.sumologic.com/) company Sumo Logic. Users need to be able to immediately pinpoint the relevant information, especially in a debugging scenario. “There are seven layers of Kubernetes, and I can see metrics and logs across all seven layers. But what am I supposed to do with this?” Ramanathan said. The goal of most Kubernetes dashboards is to connect the dots for users, so not only information but also the logical next course of action is obvious.

[Kalyan Ramanathan](https://www.linkedin.com) 解释说：“当我们构建 Kubernetes 解决方案时，我们首先从本机了解 Kubernetes 结构，然后创建一组仪表板，让您能够以深思熟虑的方式剥洋葱皮。” com/in/kalyanramanathan/)，[可观察性](https://www.sumologic.com/) 公司 Sumo Logic 的产品营销副总裁。用户需要能够立即查明相关信息，尤其是在调试场景中。 “Kubernetes 有七层，我可以查看所有七层的指标和日志。可我该怎么办呢？”拉马纳坦说。大多数 Kubernetes 仪表板的目标是为用户连接各个点，因此不仅信息而且合乎逻辑的下一步行动是显而易见的。

Dashboards can also provide a way to correlate signals, making it easier to understand relationships between the different logging metrics and the different parts of the system and ultimately easier to troubleshoot during incidents.

仪表板还可以提供一种关联信号的方法，从而更容易理解不同日志记录指标与系统不同部分之间的关系，并最终更容易在事件期间进行故障排除。

## **More than Pure Ops**

## **不仅仅是纯粹的操作**

One of the challenges with a technology as complex as Kubernetes is that one dashboard isn’t always enough. One dashboard might be enough for an incident response use case (though it often isn’t, and dashboard proliferation is its own problem), but organizations also need a way to track things like security and cost. It’s hard to spot the one container running as root, for example, without some kind of dashboard that is collecting that data and alerting users. The same goes for something like costs: Cloud costs are notoriously byzantine. Without a dedicated dashboard to sort through the information, the chances of anyone truly understanding why the bill doubled last month are pretty low.

像 Kubernetes 这样复杂的技术面临的挑战之一是，一个仪表板并不总是足够的。对于事件响应用例来说，一个仪表板可能就足够了（尽管通常不是这样，仪表板扩散是其自身的问题），但组织还需要一种方法来跟踪安全性和成本等事项。例如，如果没有某种仪表板来收集数据并提醒用户，就很难发现以 root 身份运行的一个容器。成本之类的事情也是如此：众所周知，云成本是拜占庭式的。如果没有专门的仪表板来对信息进行分类，那么任何人真正理解为什么上个月账单翻了一番的机会非常低。

Whether it’s related to security, cost or something else, dashboards are also a way to create and enforce policies or put the “guardrails” on developers. Without dashboards, getting centralized control of Kubernetes is near impossible, leaving individual developers responsible for correct configurations every time. This doesn’t matter much to a developer creating a pet project on the side, but it matters a lot to an organization in a regulated industry with a team of 800 developers with varying skill levels and hundreds of microservices.

无论是与安全性、成本还是其他方面有关，仪表板也是创建和执行策略或为开发人员设置“护栏”的一种方式。没有仪表板，对 Kubernetes 进行集中控制几乎是不可能的，每个开发人员每次都要负责正确的配置。这对于创建宠物项目的开发人员来说无关紧要，但对于受监管行业中的组织而言却非常重要，该组织拥有 800 名不同技能水平的开发人员和数百个微服务团队。

## **Maturity Markers**

## **成熟度标记**

So what do dashboards ultimately say about a project, and what do they mean for the community and for adoption in general? “What I see in my market is that a lot of the time if there is a dashboard it shows about the maturity of the product,” Levine said. “If you have a beautiful UI, usually, in this market, then it’s not just an open source project that two people are trying to spin up. If you’re thinking about the user interface, it means there is a certain amount of maturity.”

那么仪表板最终对一个项目有什么意义，它们对社区和一般采用意味着什么？ “我在我的市场上看到的是，很多时候如果有一个仪表板，它会显示产品的成熟度，”莱文说。 “如果你有一个漂亮的用户界面，通常在这个市场上，那么它不仅仅是一个两个人试图启动的开源项目。如果你在考虑用户界面，这意味着有一定的成熟度。”

It also is a requirement for widespread adoption. “When I started the company and it was just open source, I was talking to a lot of customers and showing them demos with the command line,” Levine said about Gloo, the open source ingress controller Solo.io produces. “To be honest, it was hard to impress them. Then we took exactly the same project and we put on a beautiful UI. Suddenly, I gave them the same demo and everyone said whoa, yeah, we want to proceed.”

这也是广泛采用的要求。 “当我创办这家公司时，它只是开源的，我正在与很多客户交谈，并使用命令行向他们展示演示，”莱文谈到 Solo.io 生产的开源入口控制器 Gloo 时说。 “老实说，很难给他们留下深刻印象。然后我们采用了完全相同的项目，并设计了一个漂亮的 UI。突然，我给了他们同样的演示，每个人都说哇，是的，我们要继续。”

Feature image by [Benjamin Child](https://unsplash.com/@bchild311?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/dashboard?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText). 

[Benjamin Child](https://unsplash.com/@bchild311?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) 的专题图片 [Unsplash](https://unsplash.com/s/photos/dashboard?utm_source=unsplash&utm_medium=推荐&utm_content=creditCopyText)。

