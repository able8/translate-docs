# The Problem With Agile Scrum (And Why We Use Kanban Instead)

# 敏捷 Scrum 的问题（以及我们为什么使用看板）

Discover why Bill Buckley, CloudZero VPE, chose Kanban for CloudZero's engineering team and the benefits Kanban provides for software engineering.

了解 CloudZero VPE 的 Bill Buckley 为什么为 CloudZero 的工程团队选择看板以及看板为软件工程提供的好处。

April 15, 2021

Every engineering team has their own approach when it comes to development methodologies. Most teams have embraced popular frameworks, [Agile Scrum](https://www.cprime.com/resources/what-is-agile-what-is-scrum/) seems to be the most popular, both putting their own spin on it and choosing the parts that work for them.

每个工程团队在开发方法方面都有自己的方法。大多数团队都采用了流行的框架，[Agile Scrum](https://www.cprime.com/resources/what-is-agile-what-is-scrum/) 似乎是最受欢迎的，两者都有自己的想法并选择适合他们的部件。

Despite any differences, we’re all out to achieve the same goal. We want a process that scales with our organizations and results in happy teams, high velocity, and quality software.

尽管存在任何差异，但我们都全力以赴实现相同的目标。我们想要一个与我们的组织一起扩展的流程，并产生快乐的团队、高速度和高质量的软件。

At CloudZero, we base our development process on the [Kanban methodology](https://www.atlassian.com/agile/kanban). I know from speaking with peers that it’s not a particularly common approach, maybe 1 in 10 engineering leaders I come across use it. Despite not being as popular, we’ve found it to be incredibly successful and well-aligned with the outcomes all engineering leaders want.

在 CloudZero，我们的开发流程基于 [看板方法论](https://www.atlassian.com/agile/kanban)。我从与同行的交谈中了解到，这并不是一种特别常见的方法，我遇到的工程领导者中可能有十分之一使用它。尽管不那么受欢迎，但我们发现它非常成功，并且与所有工程领导者想要的结果非常吻合。

Since Kanban is a lesser adopted framework, I thought I’d share a little bit about why we like it and how I think it can be a great way for engineering organizations to approach development work. Hopefully, you take away some ideas you can apply to your team, whether you decide to fully embrace Kanban or not.

由于看板是一个较少采用的框架，我想我想分享一下我们为什么喜欢它以及我认为它如何成为工程组织处理开发工作的好方法。希望您带走一些可以应用于团队的想法，无论您是否决定完全接受看板。

## A Brief History Of Development Methodologies (And My Experience With Them)

## 开发方法简史（以及我对它们的经验）

Before we jump in, I want to provide a little bit of context about how I arrived at my decision to adopt Kanban and ultimately why I landed on it as my framework of choice.

在我们开始之前，我想提供一些背景信息，说明我是如何决定采用看板的，以及最终为什么我选择看板作为我的选择框架。

I started my career as a software developer at EMC. As you’d expect from a large company a few decades ago, our processes were about as “old school waterfall” as it gets. I remember working on requirement documents for a single driver that would end up in a larger system, optimistically expected to be shipped two years later.

我的职业生涯始于 EMC 的软件开发人员。正如您在几十年前对一家大公司所期望的那样，我们的流程几乎就像“老派瀑布”一样。我记得为一个最终会出现在更大系统中的单个驱动程序编写需求文档，乐观地预计将在两年后发布。

When Agile was popularized, it was a massive shift and there’s no question that it was revolutionary. The Agile manifesto helped bring a completely new way of looking at how engineers could work together to build complex systems and results in more autonomy, quality, and velocity across the field.

当敏捷普及时，这是一个巨大的转变，毫无疑问它是革命性的。敏捷宣言帮助带来了一种全新的方式来看待工程师如何协同工作来构建复杂的系统，并在整个领域实现更高的自主性、质量和速度。

As Agile has spread, Agile Scrum has emerged as one of the most popular approaches to codifying a process around the Agile tenets. Scrum involves planning in sprints, which usually involves committing to a certain amount of work for a predetermined segment of time (usually two weeks).

随着敏捷的传播，敏捷 Scrum 已成为围绕敏捷原则编写流程的最流行方法之一。 Scrum 涉及冲刺计划，这通常涉及在预定的时间段（通常为两周）内承诺一定量的工作。

It’s not a _bad_ methodology and it works for many teams, but in my experience, it has some limitations.

这不是一种_坏_方法，它适用于许多团队，但根据我的经验，它有一些局限性。

## Software Isn’t Exactly Like Manufacturing

## 软件并不完全像制造业

Many of these development methodologies are modeled after manufacturing concepts.

许多这些开发方法都是根据制造概念建模的。

If you've ever read (or heard about) [The Phoenix Project](https://www.amazon.com/Phoenix-Project-DevOps-Helping-Business/dp/0988262592), you probably remember that hero of the book , Bill Palmer, visits a manufacturing plant. Once he sees how auto parts are made, he “sees the light” and saves his company by implementing the same processes back at his office to build software.

如果您曾经读过（或听说过）[凤凰计划](https://www.amazon.com/Phoenix-Project-DevOps-Helping-Business/dp/0988262592)，您可能还记得书中的那个英雄, Bill Palmer, 参观制造工厂。一旦他看到汽车零部件是如何制造的，他就会“看到曙光”，并通过在他的办公室实施相同的流程来构建软件来拯救他的公司。

While it’s a great story and an entertaining parable, the metaphor can be dangerous.

虽然这是一个伟大的故事和有趣的寓言，但这个比喻可能很危险。

Software development isn’t linear, nor is it about producing something repeatedly. Every software project is different and teams need to consider how new software will be integrated with existing products.

软件开发不是线性的，也不是重复生产某些东西。每个软件项目都是不同的，团队需要考虑如何将新软件与现有产品集成。

Furthermore, once your product or feature is produced, you can’t forget about it like an auto part you’ve shipped off to a store. It becomes a living and breathing application that behaves differently as users interact with it, and requires ongoing care and feeding.

此外，一旦您的产品或功能被生产出来，您就不能像运送到商店的汽车零件一样忘记它。它变成了一个活生生的应用程序，当用户与其交互时，它的行为会有所不同，并且需要持续的照顾和喂养。

My belief is that teams do better with a method that embraces the inherent complexity and uncertainty of software development, rather than one that introduces arbitrary time constraints.

我的信念是，团队使用包含软件开发固有复杂性和不确定性的方法会做得更好，而不是引入任意时间限制的方法。

## The Problem With Sprints 

## Sprint 的问题

The Scrum methodology has some real weaknesses, especially when you consider that we need to control higher levels of uncertainty and variability than many other fields. It's undoubtedly better than the old Waterfall days but has many traps that are easy to fall into.

Scrum 方法有一些真正的弱点，尤其是当您考虑到我们需要控制比许多其他领域更高水平的不确定性和可变性时。毫无疑问，它比旧的瀑布时代要好，但有许多容易掉入的陷阱。

The intention of sprints is to help teams break down work into smaller pieces and think through how long something will take, as well as predict blockers and dependencies.

冲刺的目的是帮助团队将工作分解成更小的部分，并考虑完成某件事需要多长时间，并预测阻碍因素和依赖关系。

However, I’ve found that engineers often spend a lot of time trying to get their work to fit into specific sized time boxes when in reality we know that estimating software projects is an imprecise exercise.

然而，我发现工程师经常花费大量时间试图让他们的工作适合特定大小的时间框，而实际上我们知道估算软件项目是一种不精确的练习。

This imprecision can lead to frustration when actual timelines don’t match the planned ones. Engineers feel like they’re racing to meet artificial deadlines and might produce rushed work when quality is more important.

当实际时间表与计划的时间表不符时，这种不精确会导致沮丧。工程师们觉得他们正在赶在人为的最后期限前完成任务，并且在质量更重要时可能会匆忙完成工作。

It also leads to organizations focusing on [metrics](http://www.cloudzero.com/blog/devops-metrics) that are not aligned with making your company successful such as the percentage of stories in a sprint successfully closed.

它还导致组织专注于 [指标](http://www.cloudzero.com/blog/devops-metrics)，而这些指标与使您的公司成功不一致，例如冲刺中成功结束的故事百分比。

## Kanban Focuses On Flow

## 看板关注流程

Kanban is also a concept adopted from traditional manufacturing, but it’s not as focused on breaking work into ‘sprints’ or estimating the ‘size’ of any particular project. A basic Kanban board might focus on visualizing the flow of work through three states: **Requested, In Progress, and Done.**

看板也是从传统制造业中采用的概念，但它并不专注于将工作分解为“冲刺”或估计任何特定项目的“规模”。一个基本的看板可能专注于通过三个状态可视化工作流程：**请求、进行中和完成。**

At CloudZero, we think it’s helpful to think about our pipeline like, well, a pipeline.

在 CloudZero，我们认为将我们的管道视为管道是有帮助的。

Instead of two (or three, or four) week sprints, I’ve found that teams can work with more agency when they're focused on small batches of work that they can continuously focus on completing to high quality. Since agency and autonomy usually result in higher morale, I’ve also witnessed this method keeps engineers happier.

我发现，当团队专注于可以持续专注于高质量完成的小批量工作时，他们可以与更多的机构合作，而不是两个（或三个或四个）周的冲刺。由于代理和自治通常会带来更高的士气，我也亲眼目睹了这种方法让工程师更快乐。

Using this method also helps keep us focused on the most important problems for our end customers, and then deciding what the next highest leverage piece of work is whenever we’ve completed that first mission.

使用这种方法还有助于让我们专注于对最终客户而言最重要的问题，然后在我们完成第一个任务时决定下一个影响最大的工作是什么。

## Kanban Can Improve Metrics

## 看板可以改进指标

One of the reasons I like the Kanban method is that it helps individuals and teams focus on metrics that really matter to the business. Engineers are especially data and process-driven, so anytime you impose certain KPIs on them, they will usually try to meet them.

我喜欢看板方法的原因之一是它可以帮助个人和团队专注于对业务真正重要的指标。工程师特别是数据和流程驱动的，因此无论何时您将某些 KPI 强加给他们，他们通常都会尝试满足它们。

But it only really works for the business when the metrics and KPIs we use to evaluate performance are actually tied to the results we want — happier customers and higher profits.

但是，只有当我们用来评估绩效的指标和 KPI 实际上与我们想要的结果（更快乐的客户和更高的利润）相关联时，它才真正适用于业务。

When you run a Scrum-based project management system and everything is broken down into two-week sprints, then developers can become very focused on whether or not they finished everything in their sprint.

当你运行一个基于 Scrum 的项目管理系统并且一切都被分解成两周的冲刺时，开发人员就会非常关注他们是否在冲刺中完成了所有的事情。

With the Kanban model, we pay much more attention to cycle time, which we think is a more relevant metric than whether or not we finished all the projected work in a given timebox. Cycle time represents our ability to deliver quality results to problems that our end users have, which is a much more valuable measurement than how accurately we can estimate how many stories fit inside of a sprint.

使用看板模型，我们更加关注周期时间，我们认为这是一个比我们是否在给定的时间范围内完成所有预计工作更相关的指标。周期时间代表了我们为最终用户遇到的问题提供高质量结果的能力，这是一个比我们如何准确地估计冲刺中适合多少故事更有价值的衡量标准。

Engineers are craftspeople who use their knowledge and experience to come up with novel, innovative solutions to customer problems. Kanban allows them to focus on doing a great job on one project at a time and focus on moving that project forward and delivering value to the end customer.

工程师是工匠，他们利用他们的知识和经验为客户问题提出新颖、创新的解决方案。看板使他们能够专注于一次出色地完成一个项目，并专注于推进该项目并为最终客户提供价值。

The Kanban system doesn’t force them to break up projects into arbitrary chunks or impose metrics that aren’t actually aligned with the business goal.

看板系统不会强迫他们将项目分解成任意的块或强加与业务目标实际上不相符的指标。

Ultimately, that makes the engineers happier and the business more productive, which should be every software company’s goal. 

最终，这会让工程师更快乐，业务更高效，这应该是每个软件公司的目标。


