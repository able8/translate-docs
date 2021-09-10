# The Future of Ops Jobs


# 运维工作的未来


Aug 17, 2020  12 Minute Read



Infrastructure, ops, devops, systems engineering, sysadmin, infraops, SRE, platform engineering. As long as I’ve been doing computers, these terms have been effectively synonymous. If I wanted to tell someone what my job was, I could throw out any one of them and expect to be understood.

基础设施、运维、开发运维、系统工程、系统管理员、基础设施、SRE、平台工程。只要我从事计算机工作，这些术语就一直是有效的同义词。如果我想告诉别人我的工作是什么，我可以扔掉他们中的任何一个并期望被理解。

Every tech company had a software engineering team that built software and an operations team that built infrastructure for running that software. At some point in the past decade they would have renamed this operations team to “devops” or “SRE”, but whatever the name, that was my team and those were my people.

每家科技公司都有一个构建软件的软件工程团队和一个构建运行该软件的基础设施的运营团队。在过去十年的某个时候，他们会将这个运营团队重命名为“devops”或“SRE”，但不管叫什么名字，那是我的团队，那些是我的人。

But unless you’re an infrastructure company, infrastructure is not your mission. Which means that every second you devote to infrastructure work — and every engineer you devote to infrastructure problems — is a distraction from your core goals.

但除非您是一家基础设施公司，否则基础设施不是您的使命。这意味着你投入基础设施工作的每一秒——以及你投入基础设施问题的每一位工程师——都会分散你对核心目标的注意力。

What’s more, it’s a distraction that builds on itself. The more time and energy you spend on infrastructure, the more your focus gets scattered, and the more you deprive yourself of the time and energy that ought to be devoted to the problems your business exists to solve.

更重要的是，这是一种建立在自身之上的分心。您在基础设施上花费的时间和精力越多，您的注意力就越分散，您就越浪费时间和精力来解决您的业务存在的问题。

This isn’t exactly new. Infrastructure and operations have always been a distraction from your core business problems. It used to be the case that every company had to grow internal expertise in hardware, datacenters, networking, operating systems, config management, and so on up the tech stack until reaching their core business problems. Infrastructure and ops have always been a distraction, but until recently, a necessary one. A means to an end.

这并不新鲜。基础设施和运营一直会分散您对核心业务问题的注意力。过去的情况是，每家公司都必须在硬件、数据中心、网络、操作系统、配置管理等方面发展内部专业知识，直到解决他们的核心业务问题。基础设施和运营一直是令人分心的事情，但直到最近，还是必要的。一个达到目的的手段。

So what’s changed? These days, you increasingly have a choice. Sure, you _can_ build all that internal expertise, but every day more and more of it is being served up via API on a silver platter.

那么有什么变化呢？如今，您越来越多地有了选择。当然，您_可以_建立所有内部专业知识，但每天都有越来越多的专业知识通过 API 在银盘上提供。

Does this mean that operations is no longer important, no longer necessary? Far from it. Operations and operability are more important than ever. Let’s take a look at what’s happening to the ops profession at a high level, the emerging challenges we face, and the impact this’ll all have on our careers.

这是否意味着操作不再重要、不再必要？离得很远。运营和可操作性比以往任何时候都更加重要。让我们来看看运维职业在高层次上发生了什么、我们面临的新挑战，以及这一切对我们职业生涯的影响。

### **Changes afoot**

### **变化正在进行中**

Beyond the broader move into cloud services, here are some of the major shifts on our horizon.

除了更广泛地转向云服务之外，以下是我们地平线上的一些重大转变。

##### **From monolith to microservices**

##### **从单体到微服务**

Much has been written about the operational demands of a microservices architecture. Now that functions calling other functions involves a network hop, operational concerns are an unavoidable part of debugging even the most trivial problems. Microservices change the game from “building code” to “building systems”, which pushes more and more code writing into the realm of operations.

已经有很多关于微服务架构的操作需求的文章。现在函数调用其他函数涉及网络跳跃，操作问题是调试甚至最微不足道的问题不可避免的一部分。微服务将游戏从“构建代码”转变为“构建系统”，这将越来越多的代码编写推向了运维领域。

##### **From monitoring to observability**

##### **从监控到可观察性**

Metric-based tools like Prometheus and DataDog are infrastructure monitoring tools, and quite good ones at that. When you’re responsible for infrastructure, the questions you care about are of aggregates and trends, thresholds and capacity. Monitoring tools are the right tool for the job, because that’s how you understand whether your infrastructure is healthy, and what actions to take to make it or keep it healthy.

Prometheus 和 DataDog 等基于度量的工具是基础设施监控工具，在这方面非常出色。当您负责基础设施时，您关心的问题是聚合和趋势、阈值和容量。监控工具是适合这项工作的工具，因为这样您才能了解您的基础设施是否健康，以及采取哪些措施来使其保持健康。

Observability tools, on the other hand, are for the people writing and shipping code to users every day, and trying to inspect and understand behavior at the nexus of users, production, and code. Observability tools preserve the full context of the request. This allows you to slice and dice and tease out new correlations, as well as view events in a waterfall by time (“tracing”). Observability is how you connect the dots between software and real business impact, and between your engineers’ experience and your users’ experience.

另一方面，可观察性工具供人们每天编写代码并将其发送给用户，并尝试检查和了解用户、生产和代码之间的行为。可观察性工具保留请求的完整上下文。这允许您切片和切块并梳理出新的相关性，以及按时间查看瀑布中的事件（“跟踪”）。可观察性是您将软件与实际业务影响之间、工程师体验与用户体验之间的点联系起来的方式。

_\*Caution: many monitoring tools are trying to rebrand themselves as observability tools without first building the necessary functionality. To tell the difference,_ [_see here_](https://www.honeycomb.io/blog/so-you-want-to-build-an-observability-tool/) _._

_\*警告：许多监控工具试图在没有首先构建必要功能的情况下将自己重新命名为可观察性工具。为了区分，_ [_see here_](https://www.honeycomb.io/blog/so-you-want-to-build-an-observability-tool/)_._

##### **From magic autoinstrumentation to instrumenting with intent** 

##### **从神奇的自动仪器到有意图的仪器**

Instrumentation is just another form of commenting and documenting your code. There are tools that promise to do it automatically for you, but they aren't great at capturing intent.. Auto-instrumentation can tell you loads of ancillary details, but they will _never_ let you divine the business value intended by the engineer who built it. So suck it up and instrument your code.

检测只是注释和记录代码的另一种形式。有一些工具承诺会自动为您执行此操作，但它们在捕捉意图方面并不出色。自动检测可以告诉您大量的辅助细节，但它们永远不会让您预测构建工程师的预期业务价值它。所以把它搞定并检测你的代码。

### **Ops is dead. Long live ops.**

### **Ops 已死。行动万岁。**

Ops teams are going the way of the dodo bird, yet operability, resiliency, and reliability have never been more important. The role of operations engineers is changing fast, and the role is bifurcating along the question of infrastructure. In the future, people who would formerly have called themselves “operations engineers” (or devops engineers) will get to choose between a role that emphasizes _building infrastructure software as a service_ and a role that uses their infrastructure expertise to help teams of engineers ship software more effectively and efficiently…generally by building as little infrastructure as possible.

运营团队正在走渡渡鸟的道路，但可操作性、弹性和可靠性从未像现在这样重要。运营工程师的角色正在快速变化，并且角色正在沿着基础设施问题分叉。未来，以前自称为“运维工程师”（或 DevOps 工程师）的人将在强调_构建基础架构软件即服务_的角色和利用其基础架构专业知识帮助工程师团队交付软件的角色之间做出选择更有效和高效……通常是通过尽可能少地构建基础设施。

If your heart truly beats for working on infrastructure problems, you’re in luck! — there are more of those than ever. Go join an infrastructure company. Try one of the many companies — AWS, Azure, all the many developer tooling companies — whose mission consists of building infrastructure software, of being the best in the world at infrastructure, and selling that expertise to other companies. There are roles for software engineers who enjoy building infrastructure solutions as a service, and there are even specialist ops roles for running and operating that infrastructure at scale, or administering those data services at scale. Whether you are a developer or not, working alone or in a team, [Azure DevOps training](https://acloudguru.com/course/introduction-to-azure-devops) can help you organize the way you plan, create and deliver software.

如果您真的为解决基础设施问题而心跳加速，那么您很幸运！ ——比以往任何时候都多。去加入一家基础设施公司。尝试众多公司中的一家——AWS、Azure，以及众多开发工具公司——其使命包括构建基础设施软件，成为世界上最好的基础设施，并将这些专业知识出售给其他公司。喜欢将基础架构解决方案构建为服务的软件工程师有一些角色，甚至还有专业操作角色来大规模运行和操作该基础架构，或大规模管理这些数据服务。无论您是开发人员与否，单独工作还是团队合作，[Azure DevOps 培训](https://acloudguru.com/course/introduction-to-azure-devops) 都可以帮助您组织计划、创建和交付软件。

Otherwise, embrace the fact that your job consists of _building systems to enable teams of engineers to ship software that creates core business value_, which means home brewing as little infrastructure as possible. And what’s left?

否则，接受这样一个事实，即你的工作包括_构建系统，使工程师团队能够交付创造核心业务价值的软件_，这意味着尽可能少地自制基础设施。还剩下什么？

##### **Vendor engineering**

##### **供应商工程**

Effectively outsourcing components of your infrastructure and weaving them together into a seamless whole involves a great deal of architectural skill and domain expertise. This skill set is both rare and incredibly undervalued, especially considering how pervasive the need for it is. Think about it. If you work at a large company, dealing with other internal teams should feel like dealing with vendors. And if you work at a small company, dealing with other vendors should feel like dealing with other teams.

有效地外包基础设施的组件并将它们编织成一个无缝的整体，需要大量的架构技能和领域专业知识。这种技能集既罕见又被低估，尤其是考虑到对它的需求是多么普遍。想想看。如果你在一家大公司工作，与其他内部团队打交道就像与供应商打交道。如果你在一家小公司工作，与其他供应商打交道就像与其他团队打交道。

Anyone who wants a long and vibrant career in SRE leadership would do well to invest some energy into areas like these:

任何想要在 SRE 领导领域拥有漫长而充满活力的职业生涯的人，最好在以下领域投入一些精力：

- Learn to evaluate vendors and their products effectively. Ask piercing, probing questions to gauge compatibility and fit. Determine which areas of friction you can live with and which are dealbreakers.
- Learn to calculate and quantify the cost of your and your team’s time and labor. Be ruthless about shedding as much labor as possible in order to focus on your core business.
- Learn to manage the true cost of ownership, and to advocate and educate internally for the right solution, particularly by managing up to execs and finance folks.

- 学习有效评估供应商及其产品。提出穿孔、探索性问题以衡量兼容性和适合性。确定您可以忍受哪些摩擦领域，哪些是交易破坏者。
- 学习计算和量化您和您的团队的时间和劳动力成本。为了专注于您的核心业务，尽量减少劳动力。
- 学习管理真正的拥有成本，并在内部倡导和教育正确的解决方案，特别是通过管理高管和财务人员。

##### **Product engineering**

##### **产品工程**

One of the great tragedies of infrastructure is how thoroughly most of us managed to evade the past 20+ years of lessons in managing products and learning how to work with designers. It’s no wonder most infrastructure tools require endless laborious trainings and certifications. They simply weren’t built like modern products for humans.

基础设施的一大悲剧是我们大多数人如何彻底地逃避过去 20 多年管理产品和学习如何与设计师合作的经验教训。难怪大多数基础设施工具都需要无休止的艰苦培训和认证。它们根本不是像现代人类产品那样建造的。

- I recommend a crash course. Embed yourself within a B2B or B2C feature delivery team for a spell. Learn their rhythms, learn their language, soak up some of their instincts. You’ll need them to balance and blend your instincts for architectural correctness, scaling patterns, and firefighting. 

- 我推荐速成课程。将自己融入 B2B 或 B2C 功能交付团队中。学习他们的节奏，学习他们的语言，吸收他们的一些本能。您将需要它们来平衡和融合您对架构正确性、缩放模式和消防的直觉。

- You don’t have to become an expert in shipping features. But you should learn the elements of relationship-building the way a good product manager does. And you must learn enough about the product lifecycle that you can help debug and divert teams whose roadmaps are hopelessly intertwined and whose roadmaps are grinding to a halt.

- 您不必成为运输功能方面的专家。但是你应该像一个优秀的产品经理那样学习建立关系的要素。并且您必须对产品生命周期有足够的了解，以便您可以帮助调试和转移路线图无可救药地交织在一起且路线图陷入停顿的团队。

##### **Sociotechnical systems engineering**

##### **社会技术系统工程**

The irreducible core of the SRE/devops skill set increasingly revolves around crafting and curating efficient, effective sociotechnical feedback loops that enable and empower engineers to ship code — to move swiftly, with confidence. Your job is not to say “no” or throw up roadblocks, it’s to figure out how to help them get to yes.

SRE/devops 技能集的不可简化的核心越来越多地围绕着制作和策划高效、有效的社会技术反馈循环，使工程师能够并授权他们交付代码——快速、充满信心地行动。你的工作不是说“不”或设置障碍，而是想办法帮助他们达成“是”。

- Start with embracing releases. Lean hard into the deploy pipeline. The safest diff is the smallest diff, and you should ship automatically and compulsively. Optimize tests, CI/CD, etc so that deploys happen automatically upon merge to main, so that a single mergeset gets deployed at a time, there are no human gates, and everything goes live automatically within a few minutes of a developer committing their code . This is your holy grail, and most teams are nowhere near there.
- Design and optimize on-call rotations that load balance the effort fairly and sustainably, and won’t burn people out. Apply the appropriate amount of pressure on management to devote enough time to reliability and fixing things versus just shipping new features. Hook up the feedback loops so that the people who are getting alerted are the ones empowered and motivated to fix the problems that are paging them. Ideally, you should page the person who made the change, every time.
- Foster a culture of ownership and accountability while promulgating blamelessness throughout the org. Welcome engineers into production, and help them navigate production waters happily and successfully.

- 从拥抱发布开始。努力进入部署管道。最安全的 diff 是最小的 diff，您应该自动且强制地发货。优化测试、CI/CD 等，以便在合并到 main 时自动进行部署，以便一次部署单个合并集，没有人工门，并且一切都在开发人员提交代码后的几分钟内自动上线.这是您的圣杯，而大多数团队都离那里不远。
- 设计和优化随叫随到的轮换，以公平和可持续地负载平衡工作，并且不会让人们筋疲力尽。对管理施加适当的压力，以投入足够的时间来提高可靠性和修复问题，而不是仅仅发布新功能。连接反馈循环，以便收到警报的人有能力并有动力解决正在传唤他们的问题。理想情况下，您应该每次都呼叫进行更改的人。
- 培养主人翁精神和责任感的文化，同时在整个组织中宣扬无可指责。欢迎工程师投入生产，帮助他们愉快、成功地在生产水域航行。

##### **Managing the portfolio of technical investments.**

##### **管理技术投资组合。**

- Operability is the longest term investment / primary source of technical debt, so no one is better positioned to help evaluate and amortize those risks than ops engineers. It is effectively free to write code, compared to the gargantuan resources it takes to run that code and tend to it over the years.
- Get excellent at [migrations](https://acloudguru.com/blog/business/what-is-cloud-migration). Leave no trailing, stale remnants of systems behind to support — those are a terrible drain on the team. Surface this energy drain to decision-makers instead of letting it silently fester.
- Hold the line against writing any more code than is absolutely necessary. Or adding any more tools than are necessary. Your line is, “what is the maintenance plan for this tool?”
- Educate and influence. Lobby for the primacy of operability. Take an interest in job ladders and leveling documents. No one should be promoted to senior engineering levels unless they write and support operable services.

- 可操作性是最长期的投资/技术债务的主要来源，因此没有人比运维工程师更能帮助评估和分摊这些风险。与运行该代码并多年来倾向于它的庞大资源相比，编写代码实际上是免费的。
- 在 [migrations](https://acloudguru.com/blog/business/what-is-cloud-migration) 方面表现出色。不留下任何落后的、陈旧的系统残余来支持——这些对团队来说是一种可怕的消耗。将这种能量流失告诉决策者，而不是让它无声无息地恶化。
- 坚决反对编写任何超过绝对必要的代码。或者添加任何不必要的工具。你的台词是，“这个工具的维护计划是什么？”
- 教育和影响。可操作性至上的大厅。对工作阶梯和水平文件感兴趣。除非他们编写和支持可操作的服务，否则任何人都不应该被提升到高级工程级别。

This world is changing fast, and these changes are accelerating. Ops is everybody’s job now. Many engineers have no idea what this means, and have absorbed the lingering cultural artifacts of terror. It’s our job to fix the terror we ops folks instilled. We must find ways to reward curiosity, not punish it.

这个世界变化很快，而且这些变化正在加速。运营现在是每个人的工作。许多工程师不知道这意味着什么，并吸收了挥之不去的恐怖文物。我们的工作是解决我们操作人员灌输的恐惧。我们必须找到奖励好奇心的方法，而不是惩罚它。

There’s never been a better time to [develop cloud skills](https://acloudguru.com/solutions/individuals) and level up your career.

现在是[发展云技能](https://acloudguru.com/solutions/individuals) 和提升职业生涯的最佳时机。

_Charity Majors is the CTO at Honeycomb and a former product engineering manager at Facebook._

_Charity Majors 是 Honeycomb 的 CTO 和 Facebook 的前产品工程经理。_

## Recommended

##  受到推崇的

Get more insights, news, and assorted awesomeness around all things cloud learning.

获取有关云学习所有事物的更多见解、新闻和各种精彩内容。
