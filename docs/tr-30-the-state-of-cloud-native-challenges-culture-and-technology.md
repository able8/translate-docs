# The State of Cloud Native: Challenges, Culture and Technology

# 云原生现状：挑战、文化和技术

Cloud native has taken the IT landscape by storm. But what is it? We sat down with Pini Reznik, CTO at Container Solutions and co-author of “Cloud Native Transformation: Practical Patterns for Innovation” to try and figure out what exactly Cloud native is, which specific technology pieces, processes and cultural dynamics need to come together to create Cloud native environments and the best way for organisations to forge into the Cloud native future.

云原生已经席卷了 IT 领域。但它是什么？我们与 Container Solutions 的 CTO、《云原生转型：创新的实用模式》的合著者 Pini Reznik 坐下来，试图弄清楚云原生到底是什么，哪些特定的技术片段、流程和文化动态需要结合在一起创建云原生环境以及组织进入云原生未来的最佳方式。

April 22, 2020

If you are an IT conference buff, you have noticed the significant uptick in the number of conferences with cloud native in their name. This is just one small indicator of the speed with which Cloud native has caught on. In Spite of all these conferences, however, Cloud native is still, very much a puzzle: Is it a tool? A set of tools? A set of processes? A way of architecting applications as microservices? Automating tests and integration? Pushing new features out every day? And what does culture have to do with it?

如果您是 IT 会议爱好者，您会注意到以云原生为名的会议数量显着增加。这只是云原生流行速度的一个小指标。然而，尽管有这些会议，云原生仍然是一个非常令人困惑的问题：它是一种工具吗？一套工具？一套流程？一种将应用程序架构为微服务的方法？自动化测试和集成？每天都推出新功能？文化与它有什么关系？

To answer these questions and others, we sat down with Pini Reznik, the co-author of a new book titled **“ [Cloud Native Transformation: Practical Patterns for Innovation](https://info.container-solutions.com/oreilly-cloud-native-transformation)”** for a one-on-one exchange.

为了回答这些问题和其他问题，我们与 Pini Reznik 坐下来，他是一本名为 [Cloud Native Transformation: Practical Patterns for Innovation](https://info.container-solutions.com/oreilly-cloud-native-transformation) 新书的合著者进行一对一交流。

Pini is Co-founder and CTO at [Container solutions](https://www.container-solutions.com/) and has been at the forefront of both shaping the Cloud native landscape as well as advocating for Cloud native as the next logical step in the evolution of the IT landscape.

Pini 是 [Container Solutions](https://www.container-solutions.com/) 的联合创始人兼首席技术官，并且一直处于塑造云原生格局以及倡导云原生作为下一个逻辑的前沿在 IT 格局的演变中迈出一步。

As the title suggests the book takes a much broader view of [Cloud native](http://www.replex.io/blog/a-comprehensive-guide-to-cloud-native-application-architecture-tools-culture-and-features) beyond simply the tools that make it up and its technical aspects. Cloud native according to Pini exists at the intersection of Cloud native tools, organizational change and culture. The book also introduces the concept of patterns: a common language for the Cloud native community to communicate about concrete pathways towards digital transformation. Below is the full text of the interview.

正如书名所示，这本书对 [云原生](http://www.replex.io/blog/a-comprehensive-guide-to-cloud-native-application-architecture-tools-culture-and-features)不仅仅是构成它的工具及其技术方面。根据 Pini 的说法，云原生存在于云原生工具、组织变革和文化的交叉点。本书还介绍了模式的概念：云原生社区的一种通用语言，用于就数字化转型的具体途径进行交流。以下是采访全文。

**Q: Give us your take on Cloud Native: What is it, and why do you think it's important?**

**问：请谈谈您对 Cloud Native 的看法：它是什么，为什么您认为它很重要？**

Pini Reznik: Most people think that Cloud Native is a set of technologies related to clouds, containers, Kubernetes, microservices, etc. Even the [CNCF](https://www.cncf.io/) (Cloud Native Computing Foundation), a non-profit organization leading the way of Cloud Native adoption, is mostly talking about the technical aspects, such as containers, orchestration, and microservices.

Pini Reznik：大多数人都认为Cloud Native是一组与云、容器、Kubernetes、微服务等相关的技术。即使是[CNCF](https://www.cncf.io/)（云原生计算基金会），一个引领 Cloud Native 采用方式的非营利组织，主要谈论技术方面，例如容器、编排和微服务。

We think Cloud Native is a much bigger thing. We think that, in addition to the aforementioned technical sides, it also includes organizational change, new and more flexible processes, and much more dynamic culture that will allow the company to utilize the benefits of all those modern technologies. Cloud Native requires a full organizational transformation.

我们认为 Cloud Native 是一件更大的事情。我们认为，除了上述技术方面，它还包括组织变革、新的和更灵活的流程，以及更有活力的文化，这将使公司能够利用所有这些现代技术的好处。 Cloud Native 需要全面的组织转型。

In our mind, Cloud Native is not about new technologies but about faster innovation that helps to learn what your customers really want and deliver an improved product on a daily or even hourly basis.

在我们看来，云原生不是关于新技术，而是关于更快的创新，它有助于了解您的客户真正想要什么，并每天甚至每小时交付改进的产品。

**Q: Who are “strangers” in the context of Cloud Native and why are they a threat to traditional enterprises?**

**问：谁是云原生背景下的“陌生人”，为什么会对传统企业构成威胁？**

Pini Reznik: In our book, [_Cloud Native Transformation_](https://info.container-solutions.com/oreilly-cloud-native-transformation), we tell the story of the fictional financial company WealthGrid. In that story, we're talking about three types of strangers or companies that come to WealthGrid’s market and start taking their market share:

Pini Reznik：在我们的书 [_Cloud Native Transformation_](https://info.container-solutions.com/oreilly-cloud-native-transformation) 中，我们讲述了虚构的金融公司 WealthGrid 的故事。在那个故事中，我们谈论的是三种类型的陌生人或公司来到 WealthGrid 的市场并开始占据他们的市场份额：

Number one, A tiny startup, in our case a U.K.-based challenger, Starling Bank. Starling could build a functional mobile bank in just one year with less than 20 people. This is an amazing speed that allowed them to access essentially unlimited VC funds and grow to 800 people, and over 1.5 million customers in just three years.

第一，一家小型创业公司，在我们的案例中是英国的挑战者 Starling Bank。 Starling 不到 20 人，只需一年时间，就能打造一家功能齐全的手机银行。这是一个惊人的速度，使他们能够获得基本上无限的风险投资基金，并在短短三年内增长到 800 人和超过 150 万客户。

Two, a large incumbent bank like the Dutch ING that sets itself a goal of “being a tech company with a banking license,” as its CEO Ralph Hamers says.

二是像荷兰 ING 这样的大型现有银行，其首席执行官拉尔夫·哈默斯 (Ralph Hamers) 表示，该银行为自己设定了“成为一家拥有银行牌照的科技公司”的目标。

And three, a large tech giant like Amazon, which is currently holding a banking license and may become the largest bank in the U.S. in the matter of a couple of years. 
第三，像亚马逊这样的大型科技巨头，目前持有银行牌照，可能在几年内成为美国最大的银行。

But an even more interesting one is the latest newcomer that just landed on us totally unexpectedly and immediately shuffled the cards: the COVID-19 virus. While the rest may come slowly over the course of the years and, step by step, grab parts of the market share, this stranger is much faster and way more dramatic in its effect. Everyone is losing and has to be very quick and decisive to survive.

但更有趣的是最新的新来者，它完全出乎意料地降临在我们身上并立即洗牌：COVID-19 病毒。虽然其余的可能会随着时间的推移慢慢到来，并逐步抢占部分市场份额，但这个陌生人的影响要快得多，而且效果也更显着。每个人都在失败，必须非常迅速和果​​断才能生存。

**Q: Which cultural and technological aspects have to come together for Cloud Native?**

**问：Cloud Native 必须将哪些文化和技术方面结合起来？**

Pini Reznik: The core principle of the [Cloud Native](http://www.replex.io/blog/a-comprehensive-guide-to-cloud-native-application-architecture-tools-culture-and-features) approach is fast experimentation and research through action. This is compared to more planned and predictable ways of the more traditional approach. Teams have to be very agile and just try new things instead of spending a lot of time and effort trying to answer all the possible questions prior to jumping into the projects.

Pini Reznik：[云原生](http://www.replex.io/blog/a-comprehensive-guide-to-cloud-native-application-architecture-tools-culture-and-features)方法的核心原理是通过行动进行快速实验和研究。这与更传统的方法的更有计划和可预测的方式相比。团队必须非常敏捷，只尝试新事物，而不是在投入项目之前花费大量时间和精力尝试回答所有可能的问题。

Unfortunately, we as humans are preconditioned to preserve the current state and protect the status quo. Cloud Native is about fast change in uncertain situations.

不幸的是，我们作为人类有先决条件来保持当前状态并保护现状。云原生是关于在不确定情况下的快速变化。

The most successful Cloud Native teams are capable of living with ambiguity and uncertainty.

最成功的云原生团队能够忍受模糊和不确定性。

**Q: The book dedicates an entire chapter to the human challenge of Cloud Native. What is that challenge, and why do you think cultural change is the toughest of all the Cloud Native elements?**

**问：这本书用一整章的篇幅介绍了 Cloud Native 的人类挑战。那是什么挑战，为什么你认为文化变革是所有云原生元素中最艰难的？**

Pini Reznik: Because those challenges are hidden during our normal work. Companies understand the technological and organizational challenges but rarely focus on the human side of the change. Also, technology is easily changeable but our biases are genetic; we as a human race have been carrying them for thousands of years. It’s easy to install Kubernetes, but difficult to ignore our cognitive biases.

Pini Reznik：因为这些挑战隐藏在我们的正常工作中。公司了解技术和组织方面的挑战，但很少关注变革的人为方面。此外，技术很容易改变，但我们的偏见是遗传的；作为人类，我们已经承载了它们数千年。安装 Kubernetes 很容易，但很难忽视我们的认知偏见。

**Q: The book also introduces the Cloud Native Maturity Matrix. Give us an overview of this tool.**

**问：本书还介绍了云原生成熟度矩阵。给我们一个这个工具的概述。**

Pini Reznik: The main goal of Cloud Native maturity is to show that the transition to Cloud Native has to be done in a holistic way. Technology and organization have to change together and support each other.

Pini Reznik：云原生成熟度的主要目标是表明必须以整体方式过渡到云原生。技术和组织必须一起改变并相互支持。

For example, when a company is moving to microservices but doesn’t introduce Continuous Delivery, it will lead to too much manual work and eventual low quality. The delivery team will be always too busy with manual deliveries.

例如，当一家公司转向微服务但不引入持续交付时，将导致过多的手动工作并最终导致质量低下。交付团队总是忙于手动交付。

But also, a strong hierarchical structure should change as in [Cloud Native](http://www.replex.io/blog/a-comprehensive-guide-to-cloud-native-application-architecture-tools-culture-and- features) world. Decisions cannot be taken by managers, as they cannot possibly know everything that's going on and make reasonable and timely decisions. But to allow this to happen, the teams have to have psychological safety to make mistakes and not be punished for them.

但是，强大的层次结构应该像 Cloud Native 世界。管理者无法做出决定，因为他们不可能了解正在发生的一切并做出合理和及时的决定。但要让这种情况发生，团队必须有心理安全感，才能犯错，而不是因此而受到惩罚。

So, everything has to be done together

所以，一切都必须一起做

**Q: What is the one element of Cloud Native that is absolutely crucial to transformation?**

**问：Cloud Native 中对转型绝对至关重要的一个要素是什么？**

Pini Reznik: Starting small and expanding over time. Transformation is mostly about change management, moving from known to unknown. If you jump all in, you almost certainly will make mistakes as they are natural in a new environment. Therefore it’s better to start with small experiments and invest increasingly more, following new uncovered information.

Pini Reznik：从小处开始，随着时间的推移不断扩大。转型主要是关于变更管理，从已知到未知。如果你全力以赴，你几乎肯定会犯错误，因为它们在新环境中很自然。因此，最好从小规模实验开始，并根据新发现的信息进行越来越多的投资。

There’s a pattern called [Gradually Raising the Stakes](https://www.cnpatterns.org/strategy-risk-reduction/gradually-raising-the-stakes) that applies here.

有一种称为 [逐渐提高赌注](https://www.cnpatterns.org/strategy-risk-reduction/gradually-raise-the-stakes) 的模式适用于此。

**Q: The book also introduces the concept of patterns. How would you define patterns and how do they help?**

**Q：这本书还介绍了模式的概念。您将如何定义模式以及它们如何提供帮助？**

Pini Reznik: Patterns were introduced by Christopher Alexander in his books _The Timeless Way of Building_, _A Pattern Language_, and _The Oregon Experiment_.

Pini Reznik：模式是 Christopher Alexander 在他的书_永恒的建筑方式_、_模式语言_和_俄勒冈实验_中介绍的。

The main idea behind the books is to encourage evolutionary piecemeal improvement over big-bang changes. Exactly the principle of starting small.

这些书背后的主要思想是鼓励对大爆炸式变化的渐进式渐进式改进。正是从小处着手的原则。

To achieve this, he had to create a common language for architects and the actual users of the buildings. Patterns are like words in this language— table, chair, sofa, etc. They are sort of mutually agreed on everyone, but still open for many ways to implement them. Language is a collection of the words on a topic and a design is like a story: “There is a square table with four chairs and a sofa in a room.” 

为了实现这一目标，他必须为建筑师和建筑物的实际用户创建一种通用语言。模式就像这种语言中的单词——桌子、椅子、沙发等。它们在某种程度上是每个人都同意的，但仍然可以通过多种方式实现它们。语言是关于一个主题的词语的集合，设计就像一个故事：“房间里有一张四把椅子和一张沙发的方桌。”

[Patterns](https://www.cnpatterns.org/patterns-library) and pattern languages create shared understanding between experts and users and allow them to build the future together, when experts bring their domain knowledge and users bring their own needs and understanding of their people and organizational culture.

[Patterns](https://www.cnpatterns.org/patterns-library) 和模式语言在专家和用户之间建立共识，让他们共同构建未来，当专家带来他们的领域知识，用户带来他们自己的需求和了解他们的员工和组织文化。

Without a common language, users can’t clearly explain what they want and experts tend to build the best possible solution they can imagine without considering the real needs of their own customer. As the philosopher Ludwig Wittgenstein said, “The limits of my language are the limits of my world.”

如果没有通用语言，用户就无法清楚地解释他们想要什么，而专家往往会在不考虑客户真实需求的情况下构建他们能想象到的最佳解决方案。正如哲学家路德维希·维特根斯坦 (Ludwig Wittgenstein) 所说：“我的语言的界限就是我的世界的界限。”

**Q: How do team composition, skills, and communication differ in a Cloud Native organization as compared to legacy Waterfall models?**

**问：与传统瀑布模型相比，云原生组织中的团队组成、技能和沟通有何不同？**

Pini Reznik: The team structure is pretty much similar to Agile cross-functional team structure.

Pini Reznik：团队结构与敏捷跨职能团队结构非常相似。

There are some additions, though.

不过，还有一些补充。

In the beginning of the transformation, a small [Core Team](https://www.cnpatterns.org/organization-culture/core-team) is the most effective. Their task is to break ground, gain knowledge and set up the basics of the future Cloud Native setup.

在转型之初，一个小的 核心团队是最有效的。他们的任务是破土动工、获取知识并为未来的云原生设置奠定基础。

Then there are typically three types of teams:

然后通常有三种类型的团队：

A [Platform Team](https://www.cnpatterns.org/organization-culture/platform-team) is the team that is building and maintaining the company cloud native platform that can be used by the Build-Run teams.

[平台团队](https://www.cnpatterns.org/organization-culture/platform-team) 是构建和维护可供 Build-Run 团队使用的公司云原生平台的团队。

[Build-Run teams or DevOps teams](https://www.cnpatterns.org/organization-culture/build-run-teams-cn-devops) are the development teams that are fully responsible also for the infrastructure for their products. They are doing so by using infrastructure as code principles. They are also fully responsible for the lifecycle of their product including it’s maintenance and continuous improvement.

[Build-Run 团队或 DevOps 团队](https://www.cnpatterns.org/organization-culture/build-run-teams-cn-devops) 是完全负责其产品基础架构的开发团队。他们通过使用基础设施作为代码原则来做到这一点。他们还对产品的生命周期负全部责任，包括维护和持续改进。

A [SRE team](https://www.cnpatterns.org/organization-culture/sre-team), or Site Reliability Engineering team, is helping everyone to continuously increase quality.

[SRE 团队](https://www.cnpatterns.org/organization-culture/sre-team)，或站点可靠性工程团队，正在帮助每个人不断提高质量。

**Q: What is, in your opinion, the best way for companies to approach Cloud Native?**

**问：在您看来，公司采用 Cloud Native 的最佳方式是什么？**

Pini Reznik: We’re typically approaching the transformation by doing a bit of thinking first. Finding the right people to support the initiative, define the business case, create a vision, set up a team, etc. This all can be done in a matter of a couple of weeks. Such initial investment in strategy may save a lot of trouble later on.

Pini Reznik：我们通常通过先做一些思考来接近转型。寻找合适的人来支持计划、定义业务案例、创建愿景、组建团队等。这一切都可以在几周内完成。这种对策略的初期投资，可能会省去很多后期的麻烦。

Then we go through a series of experiments, all the way to building a minimal viable product, or MVP, which should be done as fast as possible. Following that, a gradual transformation really starts. Teams are educated and onboarded one by one, to allow the Platform Team to continue improving the platform and providing effective support.

然后我们进行一系列实验，一直到构建最小可行产品或 MVP，这应该尽快完成。之后，一个逐渐的转变才真正开始。团队一一接受教育和入职，让平台团队继续改进平台并提供有效支持。

**Q: Based on your experience, what are some common challenges that organizations face in moving to a Cloud Native way of doing things?**

**问：根据您的经验，组织在转向云原生做事方式时面临哪些常见挑战？**

Pini Reznik: Starting too big.

Pini Reznik：开始太大了。

Only doing technical change, ignoring the organizational and cultural elements.

只做技术变革，忽略组织文化要素。

Another one is starting with “lift and shift,” simply moving their old databases etc. to new technology. This leads to poor results and costs more than expected. And although companies may think that they will continue the refactoring later, they would typically lose motivation after all the old stuff is in the cloud.

另一种是从“提升和转移”开始，简单地将他们的旧数据库等转移到新技术上。这会导致糟糕的结果和超出预期的成本。尽管公司可能认为他们稍后会继续重构，但在所有旧东西都在云中之后，他们通常会失去动力。

**Q: What would you identify as the three key takeaways from the book for CIOs and CTOs?**

**问：对于 CIO 和 CTO，您认为本书的三个关键要点是什么？**

Pini Reznik: Number one, slow innovation in today's fast-changing world is an existential threat.

Pini Reznik：第一，在当今瞬息万变的世界中，缓慢的创新是一种生存威胁。

Two, Cloud Native transformation has to be done in a holistic way including technology, processes and culture

二、云原生转型要综合技术、流程、文化

And three, the “start small” principle will save time and lead to higher quality.

第三，“从小处着手”的原则将节省时间并带来更高的质量。

**Q: What’s next? What do you think comes after Cloud Native?**

**问：接下来是什么？您认为 Cloud Native 之后会发生什么？**

Pini Reznik: Technology will of course continue changing and everything will become even smaller. Microservices will become functions and speed of delivery will be even faster.

Pini Reznik：技术当然会不断变化，一切都会变得更小。微服务将成为功能，交付速度将更快。

But a more interesting change for me is the shift from large software factories to complex supply chains producing large digital systems. Same way as normal containers not just changed transportation but allowed us to produce complex products in different parts of the world and then assemble them and distribute around the world at very low cost. 

但对我来说更有趣的变化是从大型软件工厂转向生产大型数字系统的复杂供应链。与普通集装箱相同的方式不仅改变了运输方式，而且使我们能够在世界不同地区生产复杂的产品，然后将它们组装起来并以非常低的成本在世界各地分销。

I predict similar changes to happen in the software world. Instead of creating end-user products, companies will gradually shift to production of components that could be used in a variety of other products. Those components will be consumed through APIs or as functions or though whatever next round of technological change will bring us.

我预测软件世界会发生类似的变化。公司将逐渐转向生产可用于各种其他产品的组件，而不是创建最终用户产品。这些组件将通过 API 或作为函数使用，或者无论下一轮技术变革将给我们带来什么。

_[Here](https://info.container-solutions.com/oreilly-cloud-native-transformation) is a free 75 page preview of **Cloud Native Transformation: Practical Patterns for Innovation.** If you want dig deeper into Kubernetes based Cloud native environments, download our comprehensive guide below:_ 
_[此处](https://info.container-solutions.com/oreilly-cloud-native-transformation) 是 **Cloud Native Transformation: Practical Patterns for Innovation.** 的 75 页免费预览。如果您想深入了解进入基于 Kubernetes 的云原生环境，在下面下载我们的综合指南：_
