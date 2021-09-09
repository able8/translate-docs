# Why I've Been Merging Microservices Back Into The Monolith At InVision

# 为什么我一直在 InVision 将微服务合并回单体应用

December 21, 2020

2020 年 12 月 21 日

If you follow me on Twitter, you may notice that every [now (1)](https://twitter.com/BenNadel/status/1097596366321303552) [and (2)](https://twitter.com/BenNadel/status/1321530362909044737) [then (3)](https://twitter.com/BenNadel/status/1338904547159379970) I post a celebratory tweet about merging one of our microservices **back into the monolith** at [InVision](http ://www.bennadel.com/invision/co-founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision is the digital product design platform used to make the world's best customer experiences."). My tweets are usually accompanied by a Thanos GIF in which Thanos is returning the last Infinity Stone to the Infinity Gauntlet. I find this GIF quite fitting as the reuniting of the stones gives Thanos immense power; much in the same way that _reuniting the microservices_ give _me and my team_ power. I've been asked several times as to why it is that I am killing-off my microservices. So, I wanted to share a bit more insight about this particular journey in the world of web application development.

如果你在 Twitter 上关注我，你可能会注意到每个 [now (1)](https://twitter.com/BenNadel/status/1097596366321303552) [和 (2)](https://twitter.com/BenNadel/status/1321530362909044737) [then (3)](https://twitter.com/BenNadel/status/1338904547159379970) 我在 [InVision] 上发布了一条关于将我们的一个微服务 ** 重新整合到整体**中的庆祝推文（http ://www.bennadel.com/invision/co- Founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision 是数字产品设计平台，用于制作世界上最好的客户体验。”)。我的推文通常伴随着灭霸 GIF，其中灭霸将最后一颗无限宝石归还给无限手套。我发现这个 GIF 非常合适，因为石头的重聚赋予了灭霸巨大的力量；与_重新联合微服务_赋予_我和我的团队_权力的方式大致相同。我曾多次被问到为什么我要终止我的微服务。因此，我想分享更多有关 Web 应用程序开发世界中这段特殊旅程的见解。

![Thanos returning the last stone to the Infinity Gauntlet](https://bennadel-cdn.com/resources/uploads/2020/thanos-inserting-the-last-inifinity-stone.gif)

### I Am Not "Anti-Microservices"

### 我不是“反微服务”

To be very clear, I wanted to start this post off by stating unequivocally that I am **not anti-microservices**. My merging of services back into the monolith is not some crusade to get microservices out of my life. This quest is intended to **"right size" the monolith**. What I am doing is _solving a pain-point_ for my team. If it weren't reducing friction, I wouldn't spend **so much time (and opportunity cost)** lifting, shifting, and refactoring old code.

明确地说，我想通过明确声明我**不反对微服务**来开始这篇文章。我将服务合并回单体应用并不是为了让微服务脱离我的生活。这个任务的目的是**“合适的大小”巨石**。我正在做的是为我的团队_解决一个痛点_。如果不是减少摩擦，我就不会花费 ** 这么多时间（和机会成本）** 提升、转移和重构旧代码。

![Tweet highlight: 3 weeks and about 40 JIRA tickets worth of effort.](https://bennadel-cdn.com/resources/uploads/2020/merging-microservices-effort-is-not-free.png)

Every time I do this, I run the risk of introducing new bugs and breaking the user experience. Merging microservices back into the monolith, while sometimes exhilarating, it _always terrifying_; and, represents a _Master Class_ in planning, risk reduction, and testing. Again, if it weren't worth doing, I wouldn't be doing it.

每次我这样做时，我都会冒着引入新错误和破坏用户体验的风险。将微服务合并回整体，虽然有时令人振奋，但_总是令人恐惧_；并且，代表了规划、风险降低和测试方面的_大师班_。再说一次，如果不值得做，我就不会做。

### Microservices Solve Both Technical _and_ People Problems

### 微服务解决了技术_和_人员问题

In order to understand why I am destroying some microservices, it's important to understand why microservices get created in the first place. Microservices solve two types of problems: **Technical problems** and **People problems**.

为了理解我为什么要破坏一些微服务，首先要理解为什么要创建微服务。微服务解决两类问题：**技术问题**和**人员问题**。

A **Technical problem** is one in which an aspect of the application is putting an undue burden on the infrastructure; which, in turn, is likely causing a poor user experience (UX). For example, image processing requires a lot of CPU. If this CPU load becomes too great, it could start starving the rest of the application of processing resources. This could affect system latency. And, if it gets bad enough, it could start affecting _system availability_.

**技术问题**是指应用程序的某个方面给基础设施带来了不适当的负担；反过来，这可能会导致糟糕的用户体验 (UX)。例如，图像处理需要大量的 CPU。如果这个 CPU 负载变得太大，它可能会开始耗尽处理资源的其余应用程序。这可能会影响系统延迟。而且，如果它变得足够糟糕，它可能会开始影响_系统可用性_。

A **People problem**, on the other hand, has little to do with the application at all and everything to do with how your team is organized. The more people you have working in any given part of the application, the slower and more error-prone development and deployment becomes. For example, if you have 30 engineers all competing to "Continuously Deploy" (CD) the same service, you're going to get a lot of queuing; which means, a lot of engineers that _could otherwise be shipping product_ are actually sitting around waiting for their turn to deploy.

另一方面，**人员问题** 与应用程序几乎没有关系，而与您的团队组织方式有关。在应用程序的任何给定部分工作的人越多，开发和部署就越慢，越容易出错。例如，如果您有 30 名工程师都在竞争“持续部署”（CD）相同的服务，那么您将需要大量排队；这意味着，许多 _本可以交付产品_的工程师实际上正在等待轮到他们进行部署。

### Early InVision Microservices Mostly Solved "People" Problems 

### 早期的 InVision 微服务主要解决了“人”问题

[InVision](http://www.bennadel.com/invision/co-founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision is the digital product design platform used to make the world's best customer experiences.") has been a monolithic system since its onset 8-years ago when **3 engineers** were working on it. As the company began to grow and gain traction, the _number of systems_ barely increased while the _size of the engineering team_ began to grow rapidly. In a few years, we had _dozens_ of engineers - both back-end and front-end - all working on the same codebase and all deploying to the same service queue.

[InVision](http://www.bennadel.com/invision/co- Founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision是数字产品设计平台用于提供世界上最好的客户体验。”)自 8 年前 **3 名工程师**正在开发以来，它一直是一个整体系统。随着公司开始成长并获得吸引力，_系统数量_几乎没有增加，而工程团队的_规模_开始迅速增长。几年后，我们有 _数十名工程师 - 后端和前端 - 都在同一个代码库上工作，并且都部署到同一个服务队列。

As I mentioned above, having a lot of people all working in the same place can become very problematic. Not only were the various teams all competing for the same deployment resources, it meant that every time an "Incident" was declared, several teams' code had to get rolled-back; and, _no team could deploy_ while an incident was being managed. As you can imagine, this was causing a _lot of friction_ across the organization, both for the engineering team and for the product team.

正如我上面提到的，让很多人都在同一个地方工作会变得很成问题。不仅各个团队都在争夺相同的部署资源，这意味着每次宣布“事件”时，必须回滚几个团队的代码；并且，在管理事件时_没有团队可以部署_。可以想象，这在整个组织中造成了_很多摩擦_，包括工程团队和产品团队。

And so, "microservices" were born to solve the **"People problem"**. A select group of engineers started drawing boundaries around parts of the application that they felt corresponded to _team boundaries_. This was done so that teams could work more independently, deploy independently, and ship more product. Early InVision microservices had _almost nothing to do_ with solving technical problems.

因此，“微服务”的诞生是为了解决**“人的问题”**。一组选定的工程师开始围绕他们认为与_团队边界_对应的应用程序部分绘制边界。这样做是为了让团队可以更独立地工作、独立部署并交付更多产品。早期的 InVision 微服务_几乎与解决技术问题无关。

### Conway's Law Is Good If Your Boundaries Are Good

### 康威定律是好的，如果你的边界是好的

If you work with microservices, you've undoubtedly heard of ["Conway's Law"](https://en.wikipedia.org/wiki/Conway%27s_law), introduced by [Melvin Conway](https://www.melconway.com/) in 1967. It states:

如果您使用微服务，您无疑听说过 ["Conway's Law"](https://en.wikipedia.org/wiki/Conway%27s_law)，由 [Melvin Conway](https://www.melconway)介绍.com/) 于 1967 年。它指出：

> Any organization that designs a system (defined broadly) will produce a design whose structure is a copy of the organization's communication structure.

> 任何设计系统（广义上的定义）的组织都会产生一个设计，其结构是该组织通信结构的副本。

This law is often illustrated with a "compiler" example:

这个定律通常用一个“编译器”的例子来说明：

> If you have four groups working on a compiler, you'll get a 4-pass compiler.

> 如果你有四个小组在研究一个编译器，你会得到一个 4-pass 编译器。

The idea here being that the solution is "optimized" around team structures (and team communication overhead) and not necessarily designed to solve any particular technical or performance issues.

这里的想法是解决方案围绕团队结构（和团队沟通开销）进行“优化”，不一定旨在解决任何特定的技术或性能问题。

In the world _before microservices_, Conway's Law was generally discussed in a negative light. As in, Conway's Law represented poor planning and organization of your application. But, in a _post-microservices_ world, Conway's Law is given much more latitude. Because, as it turns out, if you can break your system up into a set of **independent services** with **cohesive boundaries**, you can ship more product with fewer bugs because you've created teams that are much more focused working on set of services that entail a narrower set of responsibilities.

在世界_微服务出现之前_，康威定律的讨论通常是负面的。如在，康威定律代表您的应用程序的规划和组织不善。但是，在_后微服务_世界中，康威定律被赋予了更大的自由度。因为，事实证明，如果你能将你的系统分解成一组具有 **内聚边界**的**独立服务**，你就可以以更少的错误交付更多的产品，因为你已经创建了更多的团队专注于一组承担更窄职责的服务。

Of course, the _benefits_ of Conway's Law depend heavily on _where you draw boundaries_; and, how those boundaries **evolve over time**. And this is where me and my team - the **Rainbow Team** \- come into the picture.

当然，康威定律的_好处_在很大程度上取决于_划定界限的位置_；以及，这些界限如何**随着时间的推移而演变**。这就是我和我的团队 - **Rainbow Team** \- 出现的地方。

Over the years, [InVision](http://www.bennadel.com/invision/co-founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision is the digital product design platform used to make the world's best customer experiences.") has had to evolve from both an organizational and an infrastructure standpoint. What this means is that, _under the hood_, there is an older "legacy" platform and a growing "modern" platform. As more of our teams migrate to the "modern" platform, the services for which those teams were responsible need to get handed-off to the remaining "legacy" teams.

多年来，[InVision](http://www.bennadel.com/invision/co-Founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom"InVision用于提供世界上最佳客户体验的数字产品设计平台。”)必须从组织和基础设施的角度发展。这意味着，_在幕后_，有一个较旧的“遗留”平台和一个不断增长的“现代”平台。随着我们越来越多的团队迁移到“现代”平台，这些团队负责的服务需要移交给剩余的“遗留”团队。

Today - in 2020 - **my team is the legacy team**. My team has slowly but steadily become responsible for more and more services. Which means: fewer people but more repositories, more programming languages, more databases, more monitoring dashboards, more error logs, and more late-night pages. 

今天 - 2020 年 - **我的团队是传统团队**。我的团队慢慢但稳步地负责越来越多的服务。这意味着：更少的人但更多的存储库、更多的编程语言、更多的数据库、更多的监控仪表板、更多的错误日志和更多的深夜页面。

In short, all the _benefits_ of Conway's Law for the organization have become **liabilities** over time for my "legacy" team. And so, we've been trying to "right size" our domain of responsibility, bringing balance back to Conway's Law. Or, in other words, we're trying to alter **our service boundaries** to match **our team boundary**. Which means, **merging microservices _back_ into the monolith**.

简而言之，随着时间的推移，康威定律对组织的所有_好处_都已成为我的“遗留”团队的**负债**。因此，我们一直在努力“调整”我们的职责范围，将平衡带回康威定律。或者，换句话说，我们试图改变**我们的服务边界**以匹配**我们的团队边界**。这意味着，**将微服务 _back_ 合并到单体应用中**。

### Microservices Are Not "Micro", They Are "Right Sized"

### 微服务不是“微”，而是“大小合适”

Perhaps the worst thing that's ever happened to the microservices architecture is the term, "micro". Micro is a _meaningless_ but heavily loaded term that's practically dripping with historical connotations and human bias. A far more helpful term would have been, "right sized". Microservices were never intended to be "small services", they were intended to be "right sized services."

也许微服务架构中发生过的最糟糕的事情就是“微”这个词。 Micro 是一个_毫无意义的_但内容繁重的术语，实际上充满了历史内涵和人为偏见。一个更有用的术语是“大小合适”。微服务从来就不是“小服务”，而是“大小合适的服务”。

"Micro" is apropos of nothing; it means nothing; it entails nothing. "Right sized", on the other hand, entails that a service has been **appropriately designed** to meet its requirements: it is responsible for the "right amount" of functionality. And, what's "right" is not a static notion - it is dependent on the team, its skill-set, the state of the organization, the calculated return-on-investment (ROI), the cost of ownership, and the moment in time in which that service is operating.

“微”是指无；这不代表任何意思;它没有任何意义。另一方面，“合适的规模”意味着服务已经**适当地设计**以满足其要求：它负责“适量”的功能。而且，什么是“正确的”并不是一个静态的概念——它取决于团队、其技能组合、组织的状态、计算的投资回报率 (ROI)、拥有成本以及进入的时刻。该服务运行的时间。

For my team, "right sized" means fewer repositories, fewer deployment queues, fewer languages, and fewer operational dashboards. For my _rather small_ team, "right sized" is more about "People" than it is about "Technology". So, in the same way that [InVision](http://www.bennadel.com/invision/co-founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision is the digital product design platform used to make the world's best customer experiences.") originally introduced microservices to solve "People problems", my team is now destroying those very same microservices in order to solve "People problems".

对于我的团队来说，“合适的规模”意味着更少的存储库、更少的部署队列、更少的语言和更少的操作仪表板。对于我的_相当小的_团队来说，“规模合适”更多的是关于“人”而不是“技术”。所以，就像 [InVision](http://www.bennadel.com/invision/co- Founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom “InVision 是用于打造世界上最佳客户体验的数字产品设计平台。”)最初引入微服务是为了解决“人的问题”，我的团队现在正在摧毁那些完全相同的微服务以解决“人的问题”。

The gesture is the same, the manifestation is different.

姿态相同，表现不同。

I am extremely proud of my team and our efforts on the legacy platform. We are small band of warriors; but we accomplish quite a lot with what we have. I attribute this success to our deep knowledge of the legacy platform; our **aggressive pragmatism**; and, our continued efforts to design a system that _speaks to our abilities_ rather than an attempt to expand our abilities to match our system demands. That might sound narrow-minded; but, it is the only approach that is tenable _for our team and its resources in this moment in time_.

我为我的团队和我们在传统平台上的努力感到非常自豪。我们是一小群战士；但是我们用我们拥有的东西完成了很多。我将这一成功归功于我们对传统平台的深入了解；我们的**积极务实**；并且，我们不断努力设计一个系统来_表达我们的能力_，而不是试图扩展我们的能力以满足我们的系统需求。这听起来可能很狭隘。但是，这是目前唯一可行的方法_对于我们的团队及其资源而言_。

### Epilogue: Most Technology Doesn't Have to "Scale Independently"

### 结语：大多数技术不必“独立扩展”

One of the arguments in favor of creating independent services is the idea that those services can then "scale independently". Meaning, you can be more targeted in how you provision servers and databases to meet service demands. So, rather than creating massive services to scale only a portion of the functionality, you can leave some services small while _independently_ scaling-up other services.

支持创建独立服务的论据之一是这些服务可以“独立扩展”的想法。这意味着，您可以更有针对性地配置服务器和数据库以满足服务需求。因此，与其创建大量服务来仅扩展一部分功能，您可以将一些服务保留为小规模，同时_独立地_扩展其他服务。

Of all the reasons as to why independent services are a "Good Thing", this one gets used very often but is, in my (very limited) opinion, usually irrelevant. Unless a piece of functionality is **CPU bound** or **IO bound** or **Memory bound**, independent scalability is probably not the _"ility"_ you have to worry about. Much of the time, your servers are _waiting for things to do_; adding "more HTTP route handlers" to an application is not going to suddenly drain it of all of its resources. 

在所有关于为什么独立服务是“好事”的原因中，这个经常被使用，但在我（非常有限的）看来，通常是无关紧要的。除非某项功能是**CPU 绑定** 或**IO 绑定** 或**内存绑定**，否则独立的可扩展性可能不是您需要担心的_“能力”_。大部分时间，您的服务器都在_等待要做的事情_；向应用程序添加“更多 HTTP 路由处理程序”不会突然耗尽其所有资源。

If I could go back and **redo our early microservice attempts**, I would 100% start by focusing on all the "CPU bound" functionality first: image processing and resizing, thumbnail generation, PDF exporting, PDF importing, file versioning with `rdiff`, ZIP archive generation. I would have broken teams out along those boundaries, and have them create "pure" services that dealt with nothing but Inputs and Outputs (ie, no "integration databases", no "shared file systems") such that every other service could consume them while maintaining loose-coupling.

如果我可以回去**重做我们早期的微服务尝试**，我将 100% 首先关注所有“CPU 绑定”功能：图像处理和调整大小、缩略图生成、PDF 导出、PDF 导入、文件版本控制`rdiff`，ZIP 存档生成。我会沿着这些边界拆分团队，让他们创建“纯”服务，只处理输入和输出（即，没有“集成数据库”，没有“共享文件系统”），这样其他所有服务都可以使用它们同时保持松耦合。

I'm not saying this would have solved all our problems - after all, we had more "people" problems than we did "technology" problems; but, it would have solved some more of the "right" problems, which may have made life a bit easier in the long-run.

我并不是说这会解决我们所有的问题——毕竟，我们有更多的“人”问题而不是“技术”问题；但是，它本来可以解决更多“正确”的问题，从长远来看，这可能会使生活更轻松一些。

### Epilogue: Microservices _Also_ Have a Dollars-And-Cents Cost

### 尾声：微服务 _Also_ 也有美元和美分的成本

Service don't run in the abstract: they run on servers and talk to databases and report metrics and generate log entries. All of that has a very real dollars-and-cents cost. So while your "lambda function" doesn't cost you money when you're not using it, your "microservices" most certainly do. Especially when you consider the redundancy that you need to maintain in order to create a "highly available" system.

服务不是抽象地运行：它们在服务器上运行并与数据库对话并报告指标并生成日志条目。所有这些都有非常真实的美元和美分成本。因此，虽然您的“lambda 函数”在您不使用时不会花钱，但您的“微服务”肯定会这样做。尤其是当您考虑为创建“高可用”系统而需要维护的冗余时。

My team's merging of microservices back into the monolith has had an actual impact on the bottom-line of the business (in a good way). It's not massive - we're only talking about a few small services; but, it's not zero either. So, in additional to all of the "People" benefits we get from merging the systems together, we _also get_ a dollars-and-cents benefit as well. 

我的团队将微服务合并回单体应用对业务的底线产生了实际影响（以一种好的方式）。它并不庞大——我们只是在谈论一些小的服务；但是，它也不是零。因此，除了我们通过将系统合并在一起获得的所有“人员”利益之外，我们还_还获得了_美元和美分的好处。

