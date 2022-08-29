# What are DORA (DevOps Research and Assessments) Metrics?

# 什么是 DORA（DevOps 研究和评估）指标？

https://www.splunk.com/en_us/data-insider/devops-research-and-assessment-metrics.html

DORA metrics are a framework of performance metrics that help DevOps teams understand how effectively they develop, deliver and maintain software. They identify elite, high, medium and low performing teams and provide a baseline to help organizations continuously improve their DevOps performance and achieve better business outcomes. DORA metrics were defined by Google Cloud’s DevOps Research and Assessments team based on six years of research into the DevOps practices of 31,000 engineering professionals.

DORA 指标是一个性能指标框架，可帮助 DevOps 团队了解他们开发、交付和维护软件的效率。他们识别精英、高、中和低绩效团队，并提供基准以帮助组织不断提高其 DevOps 绩效并实现更好的业务成果。 DORA 指标由 Google Cloud 的 DevOps 研究和评估团队根据对 31,000 名工程专业人员的 DevOps 实践进行了六年的研究来定义。

While DevOps and engineering leaders can often provide a gut-level assessment of their team’s performance, they struggle to quantify this value to the business — or to pinpoint where and how improvements can be made. DORA metrics can help by providing an objective way to measure and optimize software delivery performance and validate business value.

虽然 DevOps 和工程领导者通常可以对其团队的绩效进行直观的评估，但他们很难量化这种对业务的价值，或者确定可以在何处以及如何进行改进。 DORA 指标可以通过提供一种客观的方式来衡量和优化软件交付性能并验证业务价值来提供帮助。

In the following sections, we’ll look at the four specific DORA metrics, how software engineers can apply them to assess their performance and the benefits and challenges of implementing them. We’ll also look at how you can get started with DORA metrics.

在接下来的部分中，我们将研究四个特定的 DORA 指标，软件工程师如何应用它们来评估它们的性能以及实施它们的好处和挑战。我们还将了解如何开始使用 DORA 指标。

## Components of DORA Metrics

## DORA 指标的组成部分

What are the four key metrics in DORA?

The DORA framework uses the four key metrics outlined below to measure two core areas of DevOps: speed and stability. Deployment Frequency and Mean Lead Time for Changes measure DevOps speed, and Change Failure Rate and Time to Restore Service measure DevOps stability. Used together, these four DORA metrics provide a baseline of a DevOps team’s performance and clues about where it can be improved.

DORA 的四个关键指标是什么？

DORA 框架使用下面概述的四个关键指标来衡量 DevOps 的两个核心领域：速度和稳定性。变更的部署频率和平均提前期衡量 DevOps 的速度，而变更失败率和恢复服务的时间衡量 DevOps 的稳定性。这四个 DORA 指标一起使用，提供了 DevOps 团队绩效的基线以及可以改进的地方。

1. Deployment Frequency

   部署频率

Deployment frequency indicates how often an organization successfully deploys code to production or releases software to end users. DevOps’ goal of continuous development essentially requires that teams achieve multiple daily deployments; the deployment frequency metric provides them a clear picture of where they stand in relation to that goal.

部署频率表示组织成功地将代码部署到生产环境或向最终用户发布软件的频率。 DevOps 的持续开发目标本质上要求团队实现多个日常部署；部署频率指标让他们清楚地了解他们与该目标相关的位置。

The deployment frequency benchmarks are:

- **Elite**: Multiple deployments per day
- **High**: One deployment per week to one per month
- **Medium**: One deployment per month to one every six months
- **Low**: Fewer than one deployment every six months

部署频率基准是：

- **Elite**：每天多次部署
- **高**：每周一次部署到每月一次
- **中**：每月部署一次到每六个月部署一次
- **低**：每六个月部署少于一次

Organizations vary in how they define a successful deployment, and deployment frequency can even differ across teams within a single organization.

组织在定义成功部署的方式上各不相同，部署频率甚至可能因单个组织内的团队而异。

2. Mean Lead Time for Changes

   变更的平均提前期

Mean lead time for changes measures the average time between committing code and releasing that code into production. Measuring lead time is important because the shorter the lead time, the more quickly the team can receive feedback and release software improvements. Lead time is calculated by measuring how long it takes to complete each project from start to finish and averaging those times.

更改的平均提前期衡量的是从提交代码到将该代码发布到生产环境之间的平均时间。测量提前期很重要，因为提前期越短，团队收到反馈和发布软件改进的速度就越快。提前期是通过测量从开始到结束完成每个项目所需的时间并将这些时间平均来计算的。

Mean lead time for changes benchmarks are:

- **Elite**: Less than one per hour
- **High**: Between one day and one week
- **Medium**: Between one month and six months
- **Low**: More than six months

变更基准的平均提前期为：

- **精英**：每小时少于一个
- **高**：一天到一周之间
- **中**：在 1 个月到 6 个月之间
- **低**：超过六个月

An organization’s particular cultural processes — such as separate test teams or shared test environments — can impact lead time and slow a team’s performance.

组织的特定文化流程（例如单独的测试团队或共享的测试环境）会影响交付周期并降低团队的绩效。

3. Change Failure Rate

   更改失败率

Change failure rate is the percentage of deployments causing a failure in production that require an immediate fix, such as service degradation or an outage. A low change failure rate is desirable because the more time a team spends addressing failures, the less time it has to deliver new features and customer value. This metric is usually calculated by counting how many times a deployment results in failure and dividing that by the total number of deployments to get an average. You can calculate this metric as follows:

变更失败率是导致生产失败且需要立即修复（例如服务降级或中断）的部署百分比。较低的变更失败率是可取的，因为团队花在解决故障上的时间越多，交付新功能和客户价值的时间就越少。该指标通常通过计算部署导致失败的次数并将其除以部署总数得到平均值来计算。您可以按如下方式计算此指标：

_(deployment failures / total deployments) x 100_

_（部署失败/总部署）x 100_

Change failure rate benchmarks are:

- **Elite**: 0-15%
- **High**: 16-30%
- **Medium**: 16-30%
- **Low**: 16-30%

变更失败率基准是：

- **精英**：0-15%
- **高**：16-30%
- **中等**：16-30%
- **低**：16-30%

4. Time to Restore Service 

   恢复服务的时间

The time to restore service metric, sometimes called mean time to recover or mean time to repair (MTTR), measures how quickly a team can restore service when a failure impacts customers. A failure can be anything from a bug in production to an unplanned outage.

恢复服务的时间指标，有时称为平均恢复时间或平均修复时间 (MTTR)，用于衡量当故障影响客户时团队恢复服务的速度。故障可以是任何东西，从生产中的错误到计划外的中断。

Time to Restore Service benchmarks are:

- **Elite**: Less than one hour
- **High**: Less than one day
- **Medium**: Between one day and one week
- **Low**: More than six months

时间恢复服务基准是：

- **精英**：不到一小时
- **高**：不到一天
- **中**：一天到一周之间
- **低**：超过六个月

![four-key-metrics-of-dora](http://www.splunk.com/content/dam/splunk2/images/data-insider/dora-metrics/four-key-metrics-of-dora.jpg)



DORA uses four main metrics to measure two core areas of DevOps: speed and stability.

What is a DORA (DevOps Research and Assessments) survey?

DORA 使用四个主要指标来衡量 DevOps 的两个核心领域：速度和稳定性。

什么是 DORA（DevOps 研究和评估）调查？

A DORA survey is a simple way to collect information around the four DORA metrics and measure the current state of an organization’s software delivery performance. Google Cloud’s DevOps Research and Assessments team offers an official survey called the [DORA DevOps Quick Check](https://www.devops-research.com/quickcheck.html). You simply answer five multiple-choice questions and your results are compared to other organizations, providing a top-level view of which DevOps capabilities your organization should focus on to improve.

DORA 调查是收集有关 DORA 四个指标的信息并衡量组织软件交付绩效的当前状态的一种简单方法。 Google Cloud 的 DevOps 研究和评估团队提供了一项名为 [DORA DevOps 快速检查](https://www.devops-research.com/quickcheck.html) 的官方调查。您只需回答五个多项选择题，然后将您的结果与其他组织进行比较，从而提供您的组织应重点改进哪些 DevOps 功能的顶级视图。

While a DORA survey can provide generalized guidance, many organizations additionally enlist the help of third-party vendors to conduct personalized assessments. These more closely examine a company’s culture, practices, technology and processes to identify specific ways to improve its DevOps team’s productivity.

虽然 DORA 调查可以提供一般性指导，但许多组织还寻求第三方供应商的帮助来进行个性化评估。这些更仔细地检查公司的文化、实践、技术和流程，以确定提高其 DevOps 团队生产力的具体方法。

## Benefits of DORA Metrics

## DORA 指标的好处

What are some applications/use cases of DORA metrics?

DORA 指标有哪些应用/用例？

Companies in virtually any industry can use DORA metrics to measure and improve their software development and delivery performance. A mobile game developer, for example, could use DORA metrics to understand and optimize their response when a game goes offline, minimizing customer dissatisfaction and preserving revenue. A finance company might communicate the positive business impact of DevOps to business stakeholders by translating DORA metrics into dollars saved through increased productivity or decreased downtime.

几乎任何行业的公司都可以使用 DORA 指标来衡量和改进他们的软件开发和交付绩效。例如，移动游戏开发商可以使用 DORA 指标来了解和优化他们在游戏下线时的响应，从而最大限度地减少客户不满并保持收入。金融公司可以通过将 DORA 指标转化为通过提高生产力或减少停机时间而节省的资金，将 DevOps 的积极业务影响传达给业务利益相关者。

What are the benefits and challenges of DORA metrics?

DORA metrics are a useful tool for quantifying your organization’s software delivery performance and how it compares to that of other companies in your industry. This can lead to:

DORA 指标的好处和挑战是什么？

DORA 指标是一种有用的工具，可用于量化您组织的软件交付绩效以及与您所在行业的其他公司相比如何。这可能导致：

- **Better decision making**: Consistently tracking DORA metrics helps organizations use hard data to understand the current state of their development process and make decisions about how to improve it. Because DORA helps DevOps teams focus on specific metrics rather than monitoring everything, teams can more easily identify where bottlenecks exist in their process, focus efforts on resolving them and validate the results. The net result is faster, high-quality delivery driven by data rather than gut instinct.
- **Greater value**: Elite and high performing teams can be confident they are delivering value to their customers and positively impacting the business. DORA metrics also give lower-performing teams a clear view of where their process may be stalled and help them identify a path for delivering greater value.
- **Continuous improvement**: DevOps teams can baseline their performance with DORA metrics and discover what habits, policies, processes, technologies and other factors are impeding their productivity. The four key metrics provide a path to setting goals to optimize the team’s performance and determine the most effective ways to do so.

- **更好的决策**：持续跟踪 DORA 指标有助于组织使用硬数据来了解其开发过程的当前状态，并就如何改进它做出决策。因为 DORA 帮助 DevOps 团队专注于特定指标而不是监控所有内容，所以团队可以更轻松地识别流程中存在的瓶颈，集中精力解决这些瓶颈并验证结果。最终结果是由数据而非直觉驱动的更快、高质量的交付。
- **更大的价值**：精英和高绩效团队可以确信他们正在为客户创造价值并对业务产生积极影响。 DORA 指标还可以让表现不佳的团队清楚地了解他们的流程可能在哪里停滞，并帮助他们确定提供更大价值的途径。
- **持续改进**：DevOps 团队可以使用 DORA 指标衡量他们的绩效，并发现哪些习惯、政策、流程、技术和其他因素阻碍了他们的生产力。这四个关键指标提供了设定目标以优化团队绩效并确定最有效的方法的途径。

![benefits-of-dora-metrics](http://www.splunk.com/content/dam/splunk2/images/data-insider/dora-metrics/benefits-of-dora-metrics.jpg)

DORA metrics can lead to better decision making, greater value and continuous improvement. 

DORA 指标可以带来更好的决策、更大的价值和持续改进。

Even though DORA metrics provide a starting point for evaluating your software delivery performance, they can also present some challenges. Metrics can vary widely between organizations, which can cause difficulties when accurately assessing the performance of the organization as a whole and comparing your organization’s performance against another’s. Each metric typically also relies on collecting information from multiple tools and applications. Determining your Time to Restore Service, for example, may require collecting data from PagerDuty, GitHub and Jira. Variations in tools used from team to team can further complicate collecting and consolidating this data.

尽管 DORA 指标为评估您的软件交付性能提供了一个起点，但它们也可能带来一些挑战。不同组织之间的指标差异很大，这可能会在准确评估整个组织的绩效并将您的组织的绩效与其他组织的绩效进行比较时造成困难。每个指标通常还依赖于从多个工具和应用程序中收集信息。例如，确定恢复服务的时间可能需要从 PagerDuty、GitHub 和 Jira 收集数据。团队之间使用的工具的差异会使收集和整合这些数据变得更加复杂。

## Measuring With DORA Metrics

## 使用 DORA 指标进行测量

How do you measure DevOps success with DORA?

您如何使用 DORA 衡量 DevOps 的成功？

DORA uses the four key metrics to identify elite, high, medium, and low performing teams. The [State of DevOps Report](https://www.devops-research.com/research.html#reports) has shown that elite performers have 208 times more frequent code deployments, 106 times faster lead time from commit to deploy, 2,604 times faster time to recover from incidents and 7 times lower change failure rate than low performers. Elite performing teams are also twice as likely to meet or exceed their organizational performance goals.

DORA 使用四个关键指标来识别精英、高、中、低绩效团队。 [DevOps 状态报告](https://www.devops-research.com/research.html#reports) 显示，精英执行者的代码部署频率高出 208 倍，从提交到部署的交付周期快 106 倍，2,604与低绩效者相比，从事件中恢复的时间要快 1 倍，变更失败率要低 7 倍。精英绩效团队达到或超过其组织绩效目标的可能性也增加了一倍。

How do you measure and improve [MTTR](https://www.splunk.com/en_us/data-insider/what-is-mean-time-to-repair.html#:~:text=MTTR%20is%20calculated%20by%20dividing,MTTR%20would%20be%20two%20hours.) with DORA?

您如何衡量和改进 [MTTR](https://www.splunk.com/en_us/data-insider/what-is-mean-time-to-repair.html#:~:text=MTTR%20is%20calculated%20by%20除法，MTTR%20would%20be%20two%20hours.) 与 DORA？

MTTR is calculated by dividing the total downtime in a defined period by the total number of failures. For example, if a system fails three times in a day and each failure results in one hour of downtime, the MTTR would be 20 minutes.

MTTR 的计算方法是将定义期间内的总停机时间除以故障总数。例如，如果一个系统一天发生 3 次故障，每次故障导致一小时的停机时间，则 MTTR 将为 20 分钟。

_MTTR = 60 min / 3 failures = 20 minutes_

_MTTR = 60 分钟 / 3 次失败 = 20 分钟_

MTTR begins the moment a failure is detected and ends when service is restored for end users — encompassing diagnostic time, repair time, testing and all other activities.

MTTR 从检测到故障的那一刻开始，在为最终用户恢复服务时结束——包括诊断时间、修复时间、测试和所有其他活动。

In DORA, MTTR is one measure of the stability of an organization’s continuous development process and is commonly used to evaluate how quickly teams can address failures in the continuous delivery pipeline. A low MTTR indicates that a team can quickly diagnose and correct problems and that any failures will have a reduced business impact. A high MTTR indicates that a team’s incident response is slow or ineffective and any failure could result in a significant service interruption.

在 DORA 中，MTTR 是衡量组织持续开发过程稳定性的一种方法，通常用于评估团队在持续交付管道中解决故障的速度。低 MTTR 表明团队可以快速诊断和纠正问题，并且任何故障都会减少对业务的影响。高 MTTR 表明团队的事件响应缓慢或无效，任何故障都可能导致严重的服务中断。

While there’s no magic bullet for improving MTTR, response time can be reduced by following some best practices:

虽然改善 MTTR 没有灵丹妙药，但可以通过遵循一些最佳实践来缩短响应时间：

- **Understand your incidents**: Understanding the nature of your incidents and failures is the first step in reducing MTTR. Modern enterprise software can help provide a consolidated view of your siloed data, producing a reliable MTTR metric that reveals insights into its contributing factors.
- **Make sure to monitor**: The earlier you can identify a problem, the better your chances of resolving it before it impacts users. A modern monitoring solution provides a continuous stream of real-time data about your system’s performance in a single, easy-to-digest dashboard interface and will alert you to any issues.
- **Create an [incident-management](https://www.splunk.com/en_us/data-insider/what-is-incident-management.html) action plan**: Generally, companies favor one of two approaches : ad hoc responses are often necessary for smaller, resource-strapped companies, while large enterprises favor more rigid procedures and protocols. Whatever shape your plan takes, make sure it clearly outlines whom to notify when an incident occurs, how to document it and what steps to take to address it.
- **Automate your incident-management system**: While a simple phone call may work when a low-priority incident occurs during business hours, you need to make sure all your bases are covered when a major incident strikes, particularly during off hours . An automated incident-management system that can send alerts by phone, text and email to all designated responders at once is critical for mounting a quick response. 

- **了解您的事件**：了解您的事件和故障的性质是减少 MTTR 的第一步。现代企业软件可以帮助提供孤立数据的统一视图，生成可靠的 MTTR 指标，揭示对其影响因素的洞察。
- **确保监控**：越早发现问题，在影响用户之前解决问题的机会就越大。现代监控解决方案在单一、易于理解的仪表板界面中提供有关系统性能的连续实时数据流，并会提醒您任何问题。
- **创建 [事件管理](https://www.splunk.com/en_us/data-insider/what-is-incident-management.html) 行动计划**：通常，公司倾向于两种方法之一：对于资源匮乏的小公司来说，临时响应通常是必要的，而大企业则倾向于更严格的程序和协议。无论您的计划采取何种形式，请确保它清楚地概述了在事件发生时通知谁、如何记录它以及采取哪些步骤来解决它。
- **自动化您的事件管理系统**：虽然在工作时间发生低优先级事件时，一个简单的电话可能会起作用，但您需要确保在发生重大事件时覆盖您的所有基地，尤其是在非工作时间.一个可以通过电话、短信和电子邮件同时向所有指定响应者发送警报的自动化事件管理系统对于快速响应至关重要。

- **Cross-train team members for different response roles**: Having only one person knowledgeable about each system or technology is risky. What if that system goes down when they’re on vacation? With multiple engineers each versed in several relevant functions and responsibilities, your team is better positioned to respond effectively — no matter who’s on-call.
- **Take advantage of AI**: [AIOps](https://www.splunk.com/en_us/data-insider/ai-for-it-operations-aiops.html) help [DevOps](https://www.splunk.com/en_us/data-insider/what-is-devops-and-why-is-it-important.html) teams better respond to production failures and lower their MTTR. Specifically, AIOps can help detect issues before they impact users — prioritizing incidents by criticality, correlating and contextualizing related incidents, alerting and escalating incidents to the appropriate response team members, and automating remediation to resolve incidents.
- **Have a follow-up procedure**: Once an incident is resolved, it’s important to follow up with all the key team members to determine how and why the problem occurred and strategize how to prevent it from happening again. In DevOps, this typically takes the form of a blameless post-incident review, which analyzes both the technical and human factors of their response efforts that can be improved. Ultimately, this results in improved incident response and lower MTTR, and also more innovative ideas and better applications.

- **针对不同响应角色对团队成员进行交叉培训**：只有一个人了解每个系统或技术是有风险的。如果该系统在他们度假时出现故障怎么办？拥有多名工程师，每个工程师都精通多项相关职能和职责，您的团队可以更好地做出有效响应——无论谁待命。
- **利用 AI**：[AIOps](https://www.splunk.com/en_us/data-insider/ai-for-it-operations-aiops.html) 帮助 [DevOps](https:///www.splunk.com/en_us/data-insider/what-is-devops-and-why-is-it-important.html) 团队更好地应对生产故障并降低他们的 MTTR。具体来说，AIOps 可以帮助在问题影响用户之前检测到问题——按重要性对事件进行优先级排序、关联和上下文化相关事件、向相应的响应团队成员发出警报和上报事件，以及自动修复以解决事件。
- **有一个跟进程序**：一旦事件得到解决，重要的是要跟进所有关键团队成员，以确定问题发生的方式和原因，并制定如何防止再次发生的策略。在 DevOps 中，这通常采用无可指责的事后审查的形式，该审查分析了可以改进的响应工作的技术和人为因素。最终，这会改善事件响应并降低 MTTR，并带来更多创新想法和更好的应用。

What is DORA in Agile?

敏捷中的 DORA 是什么？

In Agile, DORA metrics are used to improve the productivity of DevOps teams and the speed and stability of the software delivery process. DORA supports Agile’s goal of delivering customer value faster with fewer impediments by helping identify bottlenecks. DORA metrics also provide a mechanism to measure delivery performance so teams can continuously evaluate practices and results and quickly respond to changes. In this way, DORA metrics drive data-backed decisions to foster continuous improvement.

在敏捷中，DORA 指标用于提高 DevOps 团队的生产力以及软件交付过程的速度和稳定性。 DORA 通过帮助识别瓶颈来支持敏捷的目标，即以更少的障碍更快地交付客户价值。 DORA 指标还提供了一种衡量交付绩效的机制，因此团队可以持续评估实践和结果并快速响应变化。通过这种方式，DORA 指标推动数据支持的决策，以促进持续改进。

What are flow metrics?

什么是流量指标？

Flow metrics are a framework for measuring how much value is being delivered by a product value stream and the rate at which it is delivered from start to finish. While traditional performance metrics focus on specific processes and tasks, flow metrics measure the end-to-end flow of business and its results. This helps organizations see where obstructions exist in the value stream that are preventing desired outcomes.

流量指标是一个框架，用于衡量产品价值流交付了多少价值以及它从开始到结束的交付速度。传统的绩效指标侧重于特定的流程和任务，而流程指标衡量的是端到端的业务流程及其结果。这有助于组织了解价值流中存在阻碍预期结果的障碍。

There are four primary flow metrics for measuring value streams:

- **Flow velocity** measures the number of flow items completed over a period to determine if value is accelerating.
- **Flow time** measures how much time has elapsed between the start and finish of a flow item to gauge time to market.
- **Flow efficiency** measures the ratio of active time to total flow time to identify waste in the value stream.
- **Flow load** measures the number of flow items in a value stream to identify over- and under-utilization of value streams.

衡量价值流的主要流量指标有四个：

- **流动速度**测量一段时间内完成的流动项目的数量，以确定价值是否正在加速。
- **流程时间**衡量流程项目开始和结束之间经过的时间，以衡量上市时间。
- **流动效率**衡量活动时间与总流动时间的比率，以识别价值流中的浪费。
- **流量负载**衡量价值流中的流量项目数量，以确定价值流的过度利用和利用不足。

Flow metrics help organizations see what flows across their entire software delivery process from both a customer and business perspective, regardless of what software delivery methodologies it uses. This provides a clearer view of how their software delivery impacts business results.

流程度量可帮助组织从客户和业务的角度了解其整个软件交付过程中的流程，无论其使用何种软件交付方法。这让他们更清楚地了解他们的软件交付如何影响业务结果。

## Getting Started With DORA Metrics

## DORA 指标入门

What are some popular DORA metrics tools?

有哪些流行的 DORA 指标工具？

Once you automate DORA metrics tracking, you can begin improving your software delivery performance. Severalv engineering metrics trackers capture the four key DORA metrics, including:

一旦您自动化 DORA 指标跟踪，您就可以开始提高您的软件交付性能。几个工程指标跟踪器捕获四个关键的 DORA 指标，包括：

- Faros
- Haystack
- LinearB
- Sleuth
- Velocity by Code Climate

When considering a metric tracker, it’s important to make sure it integrates with key software delivery systems including CI/CD, issue tracking and monitoring tools. It should also display metrics clearly in easily digestible formats so teams can quickly extract insights, identify trends and draw conclusions from the data.

在考虑指标跟踪器时，确保它与关键软件交付系统集成非常重要，包括 CI/CD、问题跟踪和监控工具。它还应该以易于理解的格式清楚地显示指标，以便团队可以快速提取见解、识别趋势并从数据中得出结论。

How do you get started with DORA metrics? 

您如何开始使用 DORA 指标？

To get started with DORA metrics, start collecting data. There are many data collection and visualization solutions on the market, including those mentioned above. The easiest place to start, however, is with Google's [Four Keys](https://cloud.google.com/blog/products/devops-sre/using-the-four-keys-to-measure-your-devops-performance) open source project, which it created to help DevOps teams generate DORA metrics. Four Keys is an ETL pipeline that ingests data from Github or a Gitlab repository through Google Cloud services and into Google DataStudio. The data is then aggregated and compiled into a dashboard with data visualizations of the four key DORA metrics, which DevOps teams can use to track their progress over time.

要开始使用 DORA 指标，请开始收集数据。市场上有许多数据收集和可视化解决方案，包括上面提到的那些。然而，最简单的起点是 Google 的 [四键](https://cloud.google.com/blog/products/devops-sre/using-the-four-keys-to-measure-your-devops-性能)开源项目，它创建该项目是为了帮助 DevOps 团队生成 DORA 指标。 Four Keys 是一个 ETL 管道，它通过 Google Cloud 服务从 Github 或 Gitlab 存储库中提取数据并输入到 Google DataStudio。然后将数据汇总并编译到仪表板中，其中包含四个关键 DORA 指标的数据可视化，DevOps 团队可以使用这些数据来跟踪他们的进度。

## The Bottom line: DORA metrics are the key to getting better business value from your software delivery

## 底线：DORA 指标是从软件交付中获得更好商业价值的关键

Data-backed decisions are essential for driving better software delivery performance. DORA metrics give you an accurate assessment of your DevOps team’s productivity and the effectiveness of your software delivery practices and processes. Every DevOps team should strive to align software development with their organization’s business goals. Implementing DORA metrics is the first step. 

数据支持的决策对于推动更好的软件交付性能至关重要。 DORA 指标可让您准确评估 DevOps 团队的生产力以及软件交付实践和流程的有效性。每个 DevOps 团队都应努力使软件开发与其组织的业务目标保持一致。实施 DORA 指标是第一步。

