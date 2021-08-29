# Maximizing Developer Effectiveness

# 最大限度地提高开发人员的效率

*Technology is constantly becoming smarter and more powerful. I often observe that as these technologies are introduced an organization’s productivity instead of improving has reduced. This is because the technology has increased complexities and cognitive overhead to the developer, reducing their effectiveness. In this article, the first of a series, I introduce a framework for maximizing developer effectiveness. Through research I have identified key developer feedback loops, including micro-feedback loops that developers do 200 times a day. These should be optimized so they are quick, simple and impactful for developers. I will examine how some organizations have used these feedback loops to improve overall effectiveness and productivity.*

*技术不断变得更智能、更强大。我经常观察到，随着这些技术的引入，组织的生产力不仅没有提高，反而降低了。这是因为该技术增加了开发人员的复杂性和认知开销，降低了他们的有效性。在本系列的第一篇文章中，我介绍了一个最大限度地提高开发人员效率的框架。通过研究，我确定了关键的开发人员反馈循环，包括开发人员每天执行 200 次的微反馈循环。这些应该被优化，以便它们对开发人员来说是快速、简单和有影响力的。我将研究一些组织如何使用这些反馈循环来提高整体效率和生产力。*

26 January 2021 From: https://martinfowler.com/articles/developer-effectiveness.html

------

[Tim Cochran](https://www.linkedin.com/in/timcochran)

[蒂姆·科克伦](https://www.linkedin.com/in/timcochran)

Tim Cochran is a Technical Director for the US East Market at Thoughtworks. Tim has over 19 years of experience leading work across start-ups and large enterprises in various domains such as retail, financial services, and government. He advises organizations on technology strategy and making the right technology investments to enable digital transformation goals. He is a vocal advocate for the developer experience and passionate about using data-driven approaches to improve it.

Tim Cochran 是 Thoughtworks 美国东部市场的技术总监。 Tim 拥有超过 19 年在零售、金融服务和政府等各个领域的初创企业和大型企业领导工作的经验。他就技术战略和进行正确的技术投资为组织提供建议，以实现数字化转型目标。他是开发人员体验的积极倡导者，并热衷于使用数据驱动的方法来改进它。

[PRODUCTIVITY](https://martinfowler.com/tags/productivity.html)

[生产力](https://martinfowler.com/tags/productivity.html)

[COLLABORATION](https://martinfowler.com/tags/collaboration.html)

[协作](https://martinfowler.com/tags/collaboration.html)

## CONTENTS

##  内容

- [Day in the life in a highly effective environment](https://martinfowler.com/articles/developer-effectiveness.html#DayInTheLifeInAHighlyEffectiveEnvironment)
- [Day in the life in a low effective environment](https://martinfowler.com/articles/developer-effectiveness.html#DayInTheLifeInALowEffectiveEnvironment)
- [Developer effectiveness](https://martinfowler.com/articles/developer-effectiveness.html#DeveloperEffectiveness)
- [Case Study: Spotify](https://martinfowler.com/articles/developer-effectiveness.html#CaseStudySpotify)
- [How to get started?](https://martinfowler.com/articles/developer-effectiveness.html#HowToGetStarted)
- [Feedback Loops](https://martinfowler.com/articles/developer-effectiveness.html#FeedbackLoops)
- [Introducing micro feedback loops](https://martinfowler.com/articles/developer-effectiveness.html#IntroducingMicroFeedbackLoops)
- [Organizational Effectiveness](https://martinfowler.com/articles/developer-effectiveness.html#OrganizationalEffectiveness)
- [Case Study: Etsy](https://martinfowler.com/articles/developer-effectiveness.html#CaseStudyEtsy)
- [Conclusion](https://martinfowler.com/articles/developer-effectiveness.html#Conclusion)

- [高效环境中的一天](https://martinfowler.com/articles/developer-effectiveness.html#DayInTheLifeInAHighlyEffectiveEnvironment)
- [低效率环境中的一天](https://martinfowler.com/articles/developer-effectiveness.html#DayInTheLifeInALowEffectiveEnvironment)
- [开发效率](https://martinfowler.com/articles/developer-effectiveness.html#DeveloperEffectiveness)
- [案例研究：Spotify](https://martinfowler.com/articles/developer-effectiveness.html#CaseStudySpotify)
- [如何开始？](https://martinfowler.com/articles/developer-effectiveness.html#HowToGetStarted)
- [反馈循环](https://martinfowler.com/articles/developer-effectiveness.html#FeedbackLoops)
- [引入微反馈循环](https://martinfowler.com/articles/developer-effectiveness.html#IntroducingMicroFeedbackLoops)
- [组织效率](https://martinfowler.com/articles/developer-effectiveness.html#OrganizationalEffectiveness)
- [案例研究：Etsy](https://martinfowler.com/articles/developer-effectiveness.html#CaseStudyEtsy)
- [结论](https://martinfowler.com/articles/developer-effectiveness.html#Conclusion)

------

I often help engineering organizations that are in the midst of a transformation. This is typically both a technology transformation and a cultural transformation. For example, these organizations might be attempting to break a core monolithic system into microservices, so that they can have independent teams and adopt a DevOps approach. They also might want to improve their agile and product techniques to respond faster to feedback and signals in the market.

我经常帮助处于转型中的工程组织。这通常既是技术转型，也是文化转型。例如，这些组织可能试图将核心单体系统分解为微服务，以便他们可以拥有独立的团队并采用 DevOps 方法。他们还可能希望提高他们的敏捷性和产品技术，以更快地响应市场中的反馈和信号。

Over and over, these efforts have failed at some point in the transformation journey. Managers are unhappy with delays and budget overruns, while technologists struggle to resolve roadblocks from every direction. Productivity is too low. The teams are paralyzed with a myriad of dependencies, cognitive overload, and a lack of knowledge in the new tools/processes. The promises that were made to executive leadership about the latest technology are not coming to fruition quickly enough.

一次又一次，这些努力在转型过程中的某个时刻失败了。管理人员对延迟和预算超支感到不满，而技术人员则难以从各个方向解决障碍。生产力太低。团队因无数的依赖关系、认知过载以及对新工具/流程缺乏知识而陷入瘫痪。就最新技术向行政领导层作出的承诺并没有很快实现。

There is a stark contrast of approach between companies that have high and low developer effectiveness 

开发人员效率高和低的公司之间的方法形成鲜明对比

When we look into these scenarios, a primary reason for the problems is that the engineering organization has neglected to provide developers with an effective working environment. While transforming, they have introduced too many new processes, too many new tools and new technologies, which has led to increased complexity and added friction in their everyday tasks.

当我们研究这些场景时，问题的一个主要原因是工程组织忽略了为开发人员提供有效的工作环境。在转型的过程中，他们引入了太多的新流程、太多的新工具和新技术，这导致了他们日常工作的复杂性和摩擦增加。

I work with various types of companies. These could be enterprises which are just at the beginning of their digital transformation or are halfway there, and companies which have adopted a [DevOps culture](https://martinfowler.com/bliki/DevOpsCulture.html) culture from the very beginning. I have found there is a stark contrast of approach between companies that have high and low developer effectiveness.

我与各种类型的公司合作。这些可能是刚刚开始或已完成数字化转型的企业，也可能是从一开始就采用 [DevOps 文化](https://martinfowler.com/bliki/DevOpsCulture.html) 文化的公司。我发现开发人员效率高和低的公司之间的方法形成鲜明对比。

The easiest way to explain is via a *developer day in the life:*

最简单的解释方法是通过生活中的*开发人员日：*

## Day in the life in a highly effective environment

## 在高效环境中的一天

*The developer:*

*开发商：*

- checks the team project management tool and then attends standup where she is clear about what she has to work on.
- notes that the development environment has been automatically updated with libraries matching development and production, and the CI/CD pipelines are green.
- pulls down the latest code, makes an incremental code change that is quickly validated by deploying to a local environment and by running unit tests.
- depends on another team’s business capabilities for her feature. She is able to find documentation and the API spec through a developer portal. She still has some queries, so she jumps into the team’s Slack room and quickly gets some help from another developer who is doing support.
- focuses on her task for a few hours without any interruptions.
- takes a break, gets coffee, takes a walk, plays some ping pong with colleagues.
- commits the code change, which then passes through a number of automated checks before being deployed to production. Releases the change gradually to users in production, while monitoring business and operational metrics.

- 检查团队项目管理工具，然后参加站立会议，她清楚自己必须做什么。
- 注意到开发环境已经自动更新了与开发和生产匹配的库，并且 CI/CD 管道是绿色的。
- 提取最新代码，进行增量代码更改，通过部署到本地环境和运行单元测试来快速验证。
- 她的功能取决于另一个团队的业务能力。她能够通过开发人员门户找到文档和 API 规范。她还有一些疑问，所以她跳进了团队的 Slack 房间，并迅速从另一位提供支持的开发人员那里得到了一些帮助。
- 专注于她的任务几个小时，不受任何干扰。
- 休息一下，喝杯咖啡，散步，和同事打乒乓球。
- 提交代码更改，然后在部署到生产之前通过一些自动检查。逐步向生产中的用户发布更改，同时监控业务和运营指标。

*The developer is able to make incremental progress in a day, and goes home happy.*

*开发者一天渐进式的进步，开心回家。*

## Day in the life in a low effective environment

## 生活在低效率环境中的一天

*The developer:*

*开发商：*

- starts the day having to deal immediately with a number of alerts for problems in production.
- checks a number of logging and monitoring systems to find the error report as there are no aggregated logs across systems.
- works with operations on the phone and determines that the alerts are false positives.
- has to wait for a response from architecture, security and governance groups for a previous feature she had completed.
- has a day broken up with many meetings, many of which are status meetings
- notes that a previous feature has been approved by reviewers, she moves it into another branch that kicks off a long nightly E2E test suite that is almost always red, managed by a siloed QA team.
- depends on another team's API, but she cannot find current documentation. So instead she talks to a project manager on the other team, trying to get a query. The ticket to find an answer will take a few days, so this is blocking her current task.

- 新的一天必须立即处理生产中出现的问题的一些警报。
- 检查多个日志记录和监控系统以查找错误报告，因为没有跨系统的聚合日志。
- 与手机上的操作一起使用并确定警报是误报。
- 必须等待架构、安全和治理组对她已完成的先前功能的响应。
- 一天有很多会议，其中很多是状态会议
- 注意到先前的功能已被审阅者批准，她将其移至另一个分支，该分支启动了一个漫长的夜间 E2E 测试套件，该套件几乎总是红色的，由一个孤立的 QA 团队管理。
- 取决于另一个团队的 API，但她找不到当前的文档。因此，她转而与另一个团队的项目经理交谈，试图得到一个查询。找到答案的票需要几天时间，所以这阻碍了她当前的任务。

*We could go on. But ultimately the developer doesn’t achieve much, leaves frustrated and unmotivated*

*我们可以继续。但最终开发人员没有取得多大成就，感到沮丧和没有动力*

## Developer effectiveness

## 开发人员效率

What does being effective mean? As a developer, it is delivering the maximum value to your customers. It is being able to apply your energy and innovation in the best ways towards the company’s goals. An effective environment makes it easy to put useful, high-quality software into production; and to operate it so that the developers do not have to deal with unnecessary complexities, frivolous churn or long delays — freeing them to concentrate on value-adding tasks. 

有效是什么意思？作为开发人员，它正在为您的客户提供最大价值。它能够以最佳方式将您的精力和创新应用于实现公司目标。一个有效的环境可以很容易地将有用的、高质量的软件投入生产；并对其进行操作，以便开发人员不必处理不必要的复杂性、琐碎的流失或长时间的延迟——使他们能够专注于增值任务。

In the example illustrating a low effective environment, everything takes longer than it should. As a developer, your day is made up of endless blockers and bureaucracy. It is not just one thing; it is many. This is akin to death by a 1,000 cuts. Slowly, productivity is destroyed by small inefficiencies, which have compounding effects. The feeling of inefficiency spreads throughout the organization beyond just engineering. Engineers end up feeling helpless; they are unproductive. And worse they accept it, the way of working becomes an accepted routine defining how development is done. The developers experience a [learned helplessness](https://www.psychologytoday.com/us/basics/learned-helplessness).

在说明低效率环境的示例中，一切都需要比预期更长的时间。作为开发人员，您的一天充满了无休止的阻碍和官僚主义。这不仅仅是一件事；它很多。这类似于削减 1,000 次。慢慢地，生产力被小的低效率破坏，这会产生复合效应。效率低下的感觉在整个组织中蔓延，而不仅仅是工程。工程师最终感到无助；他们没有生产力。更糟糕的是，他们接受了它，工作方式变成了一种公认的惯例，定义了如何进行开发。开发人员体验到 [习得性无助](https://www.psychologytoday.com/us/basics/learned-helpless)。

Whereas in the organization that provides a highly effective environment, there is a feeling of momentum; everything is easy and efficient, and developers encounter little friction. They spend more time creating value. It is this frictionless environment, and the culture that supports it by fostering the desire and ability to constantly improve, that is the hardest thing for companies to create when they are doing a digital transformation.

而在提供高效环境的组织中，有一种动力感；一切都简单高效，开发人员遇到的摩擦很少。他们花更多的时间创造价值。正是这种无摩擦的环境，以及通过培养不断改进的愿望和能力来支持它的文化，这是公司在进行数字化转型时最难创造的东西。

Being productive motivates developers. Without the friction, they have time to think creatively and apply themselves

高效能激励开发人员。没有摩擦，他们有时间创造性地思考和应用自己

Organizations look for ways to measure developer productivity. The common anti-pattern is to look at lines of code, feature output or to put too much focus on trying to spot the underperforming developers. It is better to turn the conversation around to focus on how the organization is providing an effective engineering environment. Being productive motivates developers. Without the friction, they have time to think creatively and apply themselves. If organizations do not do this, then in my experience the best engineers will leave. There is no reason for a developer to work in an ineffective environment when lots of great innovative digital companies are looking to hire strong technical talent.

组织正在寻找衡量开发人员生产力的方法。常见的反模式是查看代码行、功能输出或过度关注试图发现表现不佳的开发人员。最好将话题转向关注组织如何提供有效的工程环境。高效能激励开发人员。没有摩擦，他们有时间创造性地思考和应用自己。如果组织不这样做，那么根据我的经验，最好的工程师将离开。当许多伟大的创新数字公司都在寻找强大的技术人才时，开发人员没有理由在低效的环境中工作。

Let's look at an example of a company that has optimized developer effectiveness.

让我们看一个优化开发人员效率的公司的例子。

## Case Study: Spotify

## 案例研究：Spotify

Spotify conducted user research among their engineers to better understand developer effectiveness. Through this research, they uncovered two key findings:

Spotify 在他们的工程师中进行了用户研究，以更好地了解开发人员的效率。通过这项研究，他们发现了两个关键发现：

1. Fragmentation in the internal tooling. Spotify’s internal infrastructure and tooling was built as small isolated “islands” leading to context switching and cognitive load for engineers.
2. Poor discoverability. Spotify had no central place to find technical information. As information was spread all over, engineers did not even know where to start looking for information.

1. 内部工具碎片化。 Spotify 的内部基础设施和工具被构建为小的孤立“孤岛”，导致工程师的上下文切换和认知负载。
2. 可发现性差。 Spotify 没有找到技术信息的中心位置。由于信息遍布各地，工程师甚至不知道从哪里开始寻找信息。

Spotify's developer experience team describes these problems as a negative flywheel; a vicious cycle where developers are presented with too many unknowns, forcing them to make many decisions in isolation, which in turn compounds fragmentation and duplication of efforts, and ultimately erodes the end-to-end delivery time of products.

Spotify 的开发者体验团队将这些问题描述为负面飞轮；一个恶性循环，开发人员面临太多未知数，迫使他们孤立地做出许多决定，这反过来又加剧了碎片化和重复工作，最终侵蚀了产品的端到端交付时间。

![img](https://martinfowler.com/articles/developer-effectiveness/negative-flywheel.png)

Figure 1: Spotify's negative flywheel

图 1：Spotify 的负飞轮

To mitigate these complexities, they developed [Backstage](https://backstage.io/), an Open Source developer portal with a plugin architecture to help expose all infrastructure products in one place, offering a coherent developer experience and a starting point for engineers to find the information they need.

为了减轻这些复杂性，他们开发了 [Backstage](https://backstage.io/)，这是一个具有插件架构的开源开发人员门户，可帮助在一个地方公开所有基础设施产品，提供连贯的开发人员体验和起点工程师找到他们需要的信息。

## How to get started? 

## 如何开始？

What I am describing in the highly effective environment is what it feels like to work in a company that has fully embraced a DevOps culture, continuous delivery and product thinking. Very sensibly, most companies are on a journey towards achieving this environment. They have read [Accelerate](https://www.amazon.com/gp/product/B07B9F83WM/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=B07B9F83WM&linkCode=as2&tag=martinfowlerc-20)and the [State of DevOps report ](https://martinfowler.com/bliki/StateOfDevOpsReport.html). They know what type of organization they are striving to build. The four key metrics (lead time, deployment frequency, MTTR and change fail percentage) are great measures of DevOps performance.

我在高效环境中描述的是在一家完全接受 DevOps 文化、持续交付和产品思维的公司工作的感觉。非常明智的是，大多数公司都在努力实现这种环境。他们已阅读 [加速](https://www.amazon.com/gp/product/B07B9F83WM/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=B07B9F83WM&linkCode=as2&tag=martinfowlerc-20的 DevO 报告)和](https://martinfowler.com/bliki/StateOfDevOpsReport.html)。他们知道他们正在努力建立什么类型的组织。四个关键指标（提前期、部署频率、MTTR 和变更失败百分比）是 DevOps 性能的重要衡量标准。

One way to look at the DevOps measures is that they are [lagging indicators](https://www.investopedia.com/ask/answers/what-are-leading-lagging-and-coincident-indicators/). They are useful measurements to understand where you are, and to indicate when there is work to be done to figure out what tangible things the company should do to get better. Ideally, we want to identify leading lower level metrics of effectiveness that are more actionable. There is a correlation to the higher level metrics. It will ladder up. This should also be combined with other sources of research such as surveys on developer satisfaction.

看待 DevOps 措施的一种方式是它们是 [滞后指标](https://www.investopedia.com/ask/answers/what-are-leading-lagging-and-coincident-indicators/)。它们是有用的衡量标准，可帮助您了解您所处的位置，并指示何时需要完成工作以弄清楚公司应该做哪些切实的事情才能变得更好。理想情况下，我们希望确定更具可操作性的领先的较低级别的有效性指标。与更高级别的指标存在相关性。它会爬上去。这还应与其他研究来源相结合，例如对开发人员满意度的调查。

There is an overwhelming amount of good advice, practices, tools, and processes that you should use to improve. It is very hard to know what to do. My research has shown that there are a number of key developer feedback loops. I recommend focusing on optimizing these loops, making them fast and simple. Measure the length of the feedback loop, the constraints, and the resulting outcome. When new tools and techniques are introduced, these metrics can clearly show the degree to which developer effectiveness is improved or at least isn't worse.

您应该使用大量好的建议、实践、工具和流程进行改进。很难知道该怎么做。我的研究表明，有许多关键的开发人员反馈循环。我建议重点优化这些循环，使它们快速而简单。测量反馈回路的长度、约束条件和产生的结果。当引入新的工具和技术时，这些指标可以清楚地显示开发人员效率提高或至少没有变差的程度。

## Feedback Loops

##  反馈回路

The key loops I have identified are:

我确定的关键循环是：

| Feedback Loop                                       | Low Effectiveness | High Effectiveness                            |
| :-------------------------------------------------- | :---------------- |:-------------------------------------------- |
| Validate a local code change works                  | 2 mins            | 5-15 seconds (depending on tech choice)       |
| Find root cause for defect                          | 4-7 days          | 1 day                                         |
| Validate component integrates with other components | 3 days - 2 weeks  | 2 hours                                       |
| Validate a change meets non-functional requirements | 3 months          | 1 day - 1 week (depending on scope of change) |
| Become productive on new team                       | 2 months          | 4 weeks                                       |
| Get answers to an internal technical query          | 1-2 weeks         | 30 mins                                       |
| Launch a new service in production                  | 2-4 months        | 3 days                                        |
| Validate a change was useful to the customer        | 6 months or never | 1 - 4 weeks (depending on scope of change)    |

|反馈回路 |低效|高效率 |
| - | |验证本地代码更改是否有效 | 2 分钟 | 5-15 秒（取决于技术选择）|
|找出缺陷的根本原因 | 4-7 天 | 1 天 |
|验证组件与其他组件集成 | 3 天 - 2 周 | 2 小时 |
|验证更改是否满足非功能性需求 | 3 个月 | 1 天 - 1 周（取决于更改范围）|
|在新团队中变得富有成效 | 2 个月 | 4 周 |
|获取内部技术查询的答案 | 1-2 周 | 30 分钟 |
|在生产中推出新服务 | 2-4 个月 | 3 天 |
|验证更改对客户有用 | 6 个月或永远1 - 4 周（取决于变更范围）|

The metrics are based on what I have observed is possible. Not every company needs every feedback loop to be in the high effectiveness bucket, but they provide concrete goals to guide decision-making. Engineering organizations should conduct research within their specific context to figure out what cycles and metrics are important technology strategy.

这些指标基于我观察到的可能性。并非每个公司都需要每个反馈循环都在高效桶中，但它们提供了指导决策的具体目标。工程组织应在其特定背景下进行研究，以确定哪些周期和指标是重要的技术战略。

It is useful to look at what techniques have been applied to optimize the feedback loops and the journey that companies have taken to get there. Those case studies can provide many ideas to apply in your own organization.

查看应用了哪些技术来优化反馈循环以及公司为达到目标而采取的旅程是很有用的。这些案例研究可以提供许多可以应用于您自己组织的想法。

![img](https://martinfowler.com/articles/developer-effectiveness/feedback-loops.png)

Figure 2: Feedback Loops during feature development

图 2：功能开发期间的反馈循环

The diagram above shows a simplified representation of how developers use feedback loops during development. You can see that the developer validates their work is meeting the specifications and expected standards at multiple points along the way. The key observations to note are:

上图显示了开发人员如何在开发过程中使用反馈循环的简化表示。您可以看到开发人员在整个过程中的多个点验证他们的工作是否符合规范和预期标准。需要注意的主要观察是：

- Developers will run the feedback loops more often if they are shorter. 

- 如果反馈循环较短，开发人员将更频繁地运行反馈循环。

- Developers will run more often and take action on the result, if they are seen as valuable to the developer and not purely bureaucratic overhead.
- Getting validation earlier and more often reduces the rework later on.
- Feedback loops that are simple to interpret results, reduce back and forth communications and cognitive overhead.

- 开发人员将更频繁地运行并根据结果采取行动，如果他们被视为对开发人员有价值而不是纯粹的官僚开销。
- 更早、更频繁地获得验证减少了以后的返工。
- 易于解释结果的反馈循环，减少来回通信和认知开销。

When organizations fail to achieve these results, the problems are quickly compounded. There is a great deal of wasted effort for the developers. Embodied in the time spent waiting, searching, or trying to understand results. It adds up, causing significant delays in product development, which will manifest as lower scores in the four key metrics (particularly deployment frequency and lead time).

当组织未能实现这些结果时，问题会迅速复杂化。开发人员浪费了大量精力。体现在等待、搜索或试图理解结果所花费的时间。它加起来会导致产品开发的显着延迟，这将表现为四个关键指标（尤其是部署频率和交付周期）的得分较低。

## Introducing micro feedback loops

## 引入微反馈循环

From what I have observed, you have to nail the basics, the things that developers do 10, 100 or 200 times a day. I call them micro-feedback loops. This could be running a unit test while fixing a bug. It could be seeing a code change reflected in your local environment or development environments. It could be refreshing data in your environment. Developers, if empowered, will naturally optimize, but often I find the micro-feedback loops have been neglected. These loops are intentionally short, so you end up dealing with some very small time increments.

根据我的观察，您必须掌握基础知识，即开发人员每天要做 10、100 或 200 次的事情。我称它们为微反馈循环。这可能是在修复错误的同时运行单元测试。它可能会看到反映在您的本地环境或开发环境中的代码更改。它可能会刷新您环境中的数据。开发人员如果被授权，自然会优化，但我经常发现微反馈循环被忽略了。这些循环故意很短，因此您最终会处理一些非常小的时间增量。

![img](https://martinfowler.com/articles/developer-effectiveness/micro-feedback-loops-image-only.png)

Figure 3: Micro-feedback loops compound to affect larger feedback loops.

图 3：微反馈回路复合以影响更大的反馈回路。

It is hard to explain to management why we have to focus on such small problems. Why do we have to invest time to optimize a compile stage with a two minute runtime to instead take only 15 seconds? This might be a lot of work, perhaps requiring a system to be decoupled into independent components. It is much easier to understand optimizing something that is taking two days as something worth taking on.

很难向管理层解释为什么我们必须专注于这些小问题。为什么我们必须投入时间来优化运行时间为 2 分钟的编译阶段，而不是只需要 15 秒？这可能需要大量工作，可能需要将系统解耦为独立组件。更容易理解优化需要两天时间的事情是值得做的事情。

Those two minutes can add up quickly, and could top 100 minutes a day. These small pauses are opportunities to lose context and focus. They are long enough for a developer to get distracted, decide to open an email or go and get a coffee, so that now they are distracted and out of their state of flow, there is [research](https://www.ics.uci.edu/~gmark/chi08-mark.pdf) that indicates it can take up to 23 minutes to get back into the state of flow and return to high productivity. I am not suggesting that engineers should not take breaks and clear their head occasionally! But they should do that intentionally, not enforced by the environment.

这两分钟可以很快加起来，每天最多可以达到 100 分钟。这些小的停顿是失去上下文和注意力的机会。它们足够长的时间让开发人员分心，决定打开电子邮件或去喝杯咖啡，所以现在他们分心并脱离了他们的流动状态，有 [研究](https://www.ics.uci.edu/~gmark/chi08-mark.pdf) 表示可能需要长达 23 分钟才能重新进入流动状态并恢复到高生产力。我并不是建议工程师不应该偶尔休息一下并清理他们的头脑！但他们应该有意识地这样做，而不是由环境强制执行。

In reality, developers will compensate by filling these moments of inactivity with useful things. They might have two tasks going on and toggle between them. They might slow their compile frequency by batching up changes. In my research both of these will lead to a delay in integration of code, and development time.

实际上，开发人员将通过用有用的东西填充这些不活动的时刻来进行补偿。他们可能有两个任务正在进行并在它们之间切换。他们可能会通过批量更改来降低编译频率。在我的研究中，这两者都会导致代码集成和开发时间的延迟。

How far do you take optimizing? When is enough? Imagine that now we have that change down to 15 seconds, but we think we can get it to three seconds. Is that worth the investment? It depends on how difficult it is to make that change and the impact it will bring. If you can develop a tool or capability that will speed up 10 teams then it might be worth it. This is where platform thinking, rather than optimizing for individual teams, comes into play.

你在优化上有多远？什么时候够？想象一下，现在我们将更改缩短到 15 秒，但我们认为可以将其缩短到 3 秒。那值得投资吗？这取决于进行这种改变的难度以及它将带来的影响。如果您可以开发一种工具或功能来加速 10 个团队，那么这可能是值得的。这就是平台思维，而不是针对单个团队进行优化，发挥作用的地方。

Distributed systems are a particular challenge. There are many valid reasons for splitting systems into different deployable units (usually microservices). However, distributed systems also make many things difficult (see [Microservice Prerequisites](https://martinfowler.com/bliki/MicroservicePrerequisites.html)), including developer effectiveness. Sometimes teams might optimize for team autonomy or for runtime performance, but they sacrifice developer effectiveness because they do not invest in maintaining fast feedback loops. This is a very common situation my company runs into.

分布式系统是一个特殊的挑战。将系统拆分为不同的可部署单元（通常是微服务）有很多正当理由。然而，分布式系统也使很多事情变得困难（参见[微服务先决条件](https://martinfowler.com/bliki/MicroservicePrerequisites.html)），包括开发人员效率。有时，团队可能会针对团队自治或运行时性能进行优化，但他们会牺牲开发人员的效率，因为他们没有投资于维护快速反馈循环。这是我公司遇到的一种非常普遍的情况。

## Organizational Effectiveness

## 组织效率

Highly effective organizations have designed their engineering organization to optimize for effectiveness and feedback loops. Leadership over time creates a culture that leads to empowering developers to make incremental improvements to these feedback loops.

高效的组织设计了他们的工程组织来优化效率和反馈循环。随着时间的推移，领导力会创造一种文化，使开发人员能够对这些反馈循环进行渐进式改进。

It starts with a recognition by leadership that technology — and removing friction from development teams — is vital to the business. 

首先是领导层认识到技术——以及消除开发团队之间的摩擦——对业务至关重要。

It starts with a recognition by leadership that technology — and removing friction from development teams — is vital to the business. This is manifested in a number of ways.

首先是领导层认识到技术——以及消除开发团队之间的摩擦——对业务至关重要。这体现在许多方面。

Technical leaders continually measure and re-examine effectiveness. Highly effective organizations have created a framework to make data-driven decisions by tracking the four key metrics and other data points important to their context. This culture starts at the executive level and is communicated to the rest of the organization.

技术领导者不断衡量和重新检查有效性。高效的组织已经创建了一个框架，通过跟踪四个关键指标和其他对其上下文很重要的数据点来制定数据驱动的决策。这种文化始于行政级别，并传达给组织的其他部门。

In addition to the metrics, they create an open forum to listen to the individual contributors that work in the environment day to day. Developers will know the problems they face and will have many ideas on how to solve them.

除了指标之外，他们还创建了一个开放论坛来听取每天在环境中工作的个人贡献者的意见。开发人员会知道他们面临的问题，并且会对如何解决这些问题有很多想法。

Based on this information, engineering managers can decide on priorities for investments. Large problems may require correspondingly large programs of modernization to address a poor developer experience. But often it is more about empowering teams to make continuous improvement.

根据这些信息，工程经理可以决定投资的优先级。大问题可能需要相应的大型现代化程序来解决糟糕的开发人员体验。但通常更多的是授权团队进行持续改进。

A key principle is to embrace the **developer experience**. It is common to see a program of work of a team focused on this. Developer experience means technical capabilities should be built with the same approaches used for end-user product development, applying the same research, prioritization, outcome-based thinking, and consumer feedback mechanisms.

一个关键原则是拥抱**开发人员体验**。经常看到一个团队的工作计划专注于此。开发人员经验意味着技术能力应该使用与最终用户产品开发相同的方法来构建，应用相同的研究、优先级、基于结果的思维和消费者反馈机制。

To motivate developers, highly effective organizations franchise; that means developers should have the ability to improve their day to day lives. They have a policy for teams to make incremental technical improvements and manage technical debt. There should be a healthy data-backed discussion between developers and product management. Highly effective organizations also provide the ability for developers to innovate; when their teams have clear goals and a clear idea of bottlenecks, developers can be creative in solving problems. These organizations also create ways for the best ideas to "bubble to the top", and then double down, using data to evaluate what is best.

激励开发者，高效组织特许经营；这意味着开发人员应该有能力改善他们的日常生活。他们为团队制定了一项政策，以进行渐进式技术改进和管理技术债务。开发人员和产品管理人员之间应该有一个健康的数据支持的讨论。高效的组织还为开发人员提供创新能力；当他们的团队有明确的目标和明确的瓶颈概念时，开发人员可以创造性地解决问题。这些组织还创造了让最佳想法“冒泡到顶部”的方法，然后加倍努力，使用数据来评估什么是最好的。

After **commitment**, **measurement** and **empowerment** comes **scaling**.

在**承诺**之后，**衡量**和**授权**来**缩放**。

At a certain organizational size, there is a need to create efficiency through economies of scale. Organizations do this by applying [platform thinking](https://martinfowler.com/articles/talk-about-platforms.html) - creating an internal platform specifically focused on improving effectiveness. They invest in engineering teams that build technical capabilities to improve developer effectiveness. They regard other development teams as their consumers, and the services they provide are treated like products. The teams have technical product managers and success metrics related to how they are impacting the consuming teams. For example, a platform capability team focused on observability creates monitoring, logging, alerting, and tracing capabilities so that teams can easily monitor their service health and debug problems in their product.

在一定的组织规模下，需要通过规模经济来创造效率。组织通过应用[平台思维](https://martinfowler.com/articles/talk-about-platforms.html) 来做到这一点——创建一个专门专注于提高效率的内部平台。他们投资于构建技术能力以提高开发人员效率的工程团队。他们将其他开发团队视为他们的消费者，他们提供的服务被视为产品。这些团队有技术产品经理和与他们如何影响消费团队相关的成功指标。例如，专注于可观察性的平台能力团队创建了监控、日志记录、警报和跟踪功能，以便团队可以轻松监控其服务健康状况并调试产品中的问题。

The need for governance is still a priority. However, this does not need to be seen as a dirty word since its application is very different in highly effective organizations. They move away from centralized processes to a lightweight approach. This is about setting and communicating the guardrails, and then nudging teams in the right direction rather than lengthy approval process approaches. Governance can have a critical role in effectiveness when it is implemented via:

治理的需要仍然是一个优先事项。然而，这并不需要被视为一个肮脏的词，因为它在高效组织中的应用非常不同。他们从集中式流程转向轻量级方法。这是关于设置和沟通护栏，然后推动团队朝着正确的方向前进，而不是冗长的审批流程方法。当通过以下方式实施时，治理可以在有效性方面发挥关键作用：

- Clear engineering goals
- Specifying ways that teams and services communicate with each other
- Encouraging useful peer review
- Baking best practices into platform capabilities
- Automating control via [architecture fitness functions](https://evolutionaryarchitecture.com/)

- 明确的工程目标
- 指定团队和服务相互沟通的方式
- 鼓励有用的同行评审
- 将最佳实践融入平台功能
- 通过 [架构适应度函数] 自动控制(https://evolutionaryarchitecture.com/)

Essentially, effective organizations shorten the governance feedback loop. I will be expanding on this in a future article.

本质上，有效的组织缩短了治理反馈循环。我将在以后的文章中对此进行扩展。

## Case Study: Etsy 

## 案例研究：Etsy

Etsy was one of the pioneers of the DevOps movement. Its leaders have worked to embed developer effectiveness into their culture, with the belief that moving quickly is both a technical and a business strategy. They actively measure their ability to put valuable products into production quickly and safely, and will adjust their technical investments to fix any blockers or slowness.

Etsy 是 DevOps 运动的先驱之一。它的领导者一直致力于将开发人员的效率嵌入到他们的文化中，相信快速行动既是一种技术战略，也是一种商业战略。他们积极衡量自己快速、安全地将有价值的产品投入生产的能力，并将调整他们的技术投资以解决任何阻碍或缓慢的问题。

One of Etsy’s key metrics is lead time, which is measured, monitored, and displayed in real-time throughout their offices. When lead time reaches above a certain key threshold, the release engineering team will work to lower it to a reasonable level. Their CTO, Mike Fisher, talks about Etsy engineers being “fearless” to move forward quickly, having a safety net to try new things.

Etsy 的关键指标之一是交货时间，它在整个办公室中实时测量、监控和显示。当提前期达到某个关键阈值以上时，发布工程团队将努力将其降低到一个合理的水平。他们的 CTO Mike Fisher 谈到 Etsy 工程师“无所畏惧”地快速前进，拥有尝试新事物的安全网。

Deploying software fast is only half of the story. To be truly effective that software has to be valuable to consumers. Etsy does this by taking a data-driven approach, with each feature having measurable KPIs.

快速部署软件只是故事的一半。为了真正有效，软件必须对消费者有价值。 Etsy 通过采用数据驱动的方法来做到这一点，每个功能都有可衡量的 KPI。

Code changes go through a set of checks, so that developers have confidence the change meets Etsy’s SLAs for performance, availability, failure rate, etc. Once the change is in production, Etsy’s experimentation platform is able to capture user behavior metrics. Teams use the metrics to iterate on products, optimizing for the associated KPIs and user satisfaction. If eventually the change is proven not to be valuable, it would be cleaned up, thereby avoiding technical debt.

代码更改经过一组检查，以便开发人员确信更改符合 Etsy 的性能、可用性、故障率等 SLA。一旦更改投入生产，Etsy 的实验平台就能够捕获用户行为指标。团队使用这些指标来迭代产品，优化相关的 KPI 和用户满意度。如果最终证明更改没有价值，则将对其进行清理，从而避免技术债务。

Etsy has a current initiative that prioritizes the developer experience. It has four key pillars:

Etsy 目前有一项优先考虑开发人员体验的计划。它有四个关键支柱：

### 4 Pillars of Developer Experience

### 开发人员体验的 4 个支柱

**Help me craft products** ensures we have the right abstractions, libraries, and scaffolding for product engineers to do their work.

**帮我制作产品** 确保我们有正确的抽象、库和脚手架供产品工程师完成他们的工作。

**Help me develop**, test, and deploy focuses on product engineers, specifically the development environments themselves (IDEs, linters), unit/integration test patterns/runners, and the deployment tooling and processes.

**帮助我开发**、测试和部署侧重于产品工程师，特别是开发环境本身（IDE、linter）、单元/集成测试模式/运行程序以及部署工具和流程。

**Help me build with data** focuses on data scientists and machine learning engineers, making sure the entire data engineering ecosystem is set up in a way that is intuitive, easy to test, and easy to deploy.

**Help me build with data** 专注于数据科学家和机器学习工程师，确保以直观、易于测试和易于部署的方式建立整个数据工程生态系统。

**Help me reduce toil** focuses on the on-call engineers, to make sure we build production systems with the appropriate levels of automation, have runbooks and documentation that is easily accessible and current, and we track and prioritize reducing toil-y activities.

**帮助我减少工作量**专注于待命工程师，以确保我们构建具有适当自动化水平的生产系统，拥有易于访问和最新的运行手册和文档，我们跟踪并优先减少工作量活动。

This policy represents a true commitment from Etsy’s leadership to their developers. They continually verify their effectiveness by tracking metrics including the 4 key metrics, and conducting monthly surveys with developers to capture net promoter scores (NPS).

该政策代表了 Etsy 领导层对其开发人员的真正承诺。他们通过跟踪包括 4 个关键指标在内的指标，并与开发人员进行月度调查以获取净推荐值 (NPS)，从而不断验证其有效性。

## Conclusion

##  结论

The beginning of this article speaks to the importance of developer effectiveness and its impact on developer happiness and productivity. I focused on the outcomes developers aim to achieve and not just the tools and techniques. In pushing this examination further we see a series of feedback loops that developers often use while developing a product.

本文的开头谈到了开发人员效率的重要性及其对开发人员幸福感和生产力的影响。我专注于开发人员旨在实现的结果，而不仅仅是工具和技术。在进一步推动这项检查的过程中，我们看到了开发人员在开发产品时经常使用的一系列反馈循环。

I also spoke to how inefficiencies in micro-feedback loops compound to affect larger indicators such the four key metrics and product development speed. As well as highlighted the importance of developer experience as a principal and how platform thinking will help maximize efficiency and effectiveness at scale.

我还谈到了微反馈循环中的低效率如何影响更大的指标，例如四个关键指标和产品开发速度。并强调了开发人员经验作为主体的重要性以及平台思维将如何帮助最大限度地提高效率和有效性。

In the coming series, a deeper dive into developer effectiveness and the individual feedback loops will be further explained through case studies. These will provide concrete details in how organizations have been able to achieve these numbers and the resulting outcomes. In addition to describing the organizational structures and processes that enable these optimizations at a local and a global level.

在接下来的系列中，将通过案例研究更深入地了解开发人员的效率和个人反馈循环。这些将提供有关组织如何能够实现这些数字和结果的具体细节。除了描述在本地和全球层面实现这些优化的组织结构和流程。

The next article will start with the smallest micro-feedback loops.

下一篇文章将从最小的微反馈循环开始。

------

## Acknowledgements

## 致谢

Thanks to Pia Nilsson at Spotify and Keyur Govande at Etsy for collaborating on case studies about their work.

感谢 Spotify 的 Pia Nilsson 和 Etsy 的 Keyur Govande 合作进行关于他们工作的案例研究。

Many thanks to Martin Fowler for his support.

非常感谢 Martin Fowler 的支持。

Thanks to the ThoughtWorkers whose amazing work this article references.

感谢本文引用的 ThoughtWorkers 出色的工作。

Thanks to my colleagues Cassie Shum and Carl Nygard for their feedback and help with research. This article wouldn’t have been possible without Ryan Murray’s ideas about platforms thinking.

感谢我的同事 Cassie Shum 和 Carl Nygard 的反馈和研究帮助。如果没有 Ryan Murray 关于平台思维的想法，这篇文章就不可能完成。

Thanks to Mike McCormack and Gareth Morgan for editorial review. 

感谢 Mike McCormack 和 Gareth Morgan 的编辑审核。

