# How to Successfully Hand Over Systems

# 如何成功移交系统

April 20th, 2021

In a product company, changes are inevitable so as to best support the strategy and the vision. Often during such a change, new teams are formed and other ones are restructured. While there are many challenges to be solved during a big change, there’s one in particular that’s often overlooked: system ownership.

在产品公司中，为了最好地支持战略和愿景，变化是不可避免的。通常在这种变化期间，会形成新的团队并重组其他团队。虽然在重大变革期间有许多挑战需要解决，但有一个特别容易被忽视的挑战：系统所有权。

Who will take ownership of the systems that were owned by a team that doesn’t exist anymore or that are better suited to be owned by another team? It’s in everyone’s interest that the ownership be given to a team familiar with the system’s domain, so that they can continue the maintenance and evolution.

谁将拥有已不存在的团队或更适合由另一个团队拥有的系统的所有权？将所有权交给熟悉系统领域的团队符合每个人的利益，以便他们可以继续维护和演进。

Regardless of when the system handover will happen, how it’s executed is important, since the cost of failure can be high, and that could result in an outage or a significant amount of unplanned work.

无论系统切换何时发生，它的执行方式都很重要，因为失败的成本可能很高，这可能会导致中断或大量计划外工作。

Having experienced not-so-successful handovers — some of which took place over the course of a one-hour meeting — I was inspired to create a guideline that will help other teams do handovers differently. At the same time, my colleague Antonio N. went through an ownership change with his team. This also had few mishaps, so we joined forces to write a proposal document for doing system handovers at SoundCloud.

经历了不那么成功的交接——其中一些发生在一个小时的会议过程中——我受到启发，创建了一个指导方针，以帮助其他团队以不同的方式进行交接。与此同时，我的同事 Antonio N. 与他的团队进行了所有权变更。这也没有什么意外，所以我们联合起来写了一份提案文件，用于在 SoundCloud 进行系统移交。

We used the RFC approach to gather input, experiences, and opinions from the entire organization. It was welcomed with enthusiasm and since then has been used multiple times.

我们使用 RFC 方法来收集整个组织的输入、经验和意见。它受到热烈欢迎，从那时起已被多次使用。

I’m posting the complete guideline below not only to document what we did at SoundCloud, but also in hopes of providing a template for other companies to use when faced with a similar scenario.

我在下面发布完整的指南，不仅是为了记录我们在 SoundCloud 所做的事情，而且是希望为其他公司提供一个模板，以便在面临类似情况时使用。

## Guideline for Internal System Handovers

## 内部系统移交指南

The guideline is a list of questions, tasks, and actions that the involved parties should consider as part of the system handover. The topics listed can be covered in different ways: through documentation, meetings, pairing sessions, workshops, tasks, PR reviews, etc. **The goal is to help the new team understand the what, why, and how of the system, and to empower them to maintain, change, and improve it.**

该指南是相关方在系统切换过程中应考虑的问题、任务和操作的列表。列出的主题可以通过不同的方式涵盖：通过文档、会议、配对会议、研讨会、任务、公关审查等。 **目标是帮助新团队了解系统的内容、原因和方式，以及使他们能够维护、改变和改进它。**

### General Recommendations

### 一般建议

As the system ownership change is a process itself, we recommend that it’s driven by the new team with the help and support of the previous system owner. Both teams should collaborate on the planning and execution of the tasks.

由于系统所有权变更本身就是一个过程，我们建议它由新团队在前任系统所有者的帮助和支持下推动。两个团队应该在任务的计划和执行上进行协作。

There will be some new documents and artifacts produced as an outcome of the system ownership change. We recommend storing them in the system’s repo when possible, or else including a link in the repo (e.g. from the README file). This will help with both onboarding new team members and potential ownership changes in the future.

作为系统所有权更改的结果，将产生一些新文档和工件。我们建议尽可能将它们存储在系统的 repo 中，或者在 repo 中包含一个链接（例如来自 README 文件）。这将有助于新团队成员的入职和未来潜在的所有权变更。

### Why Are We Changing the Ownership?

### 为什么我们要更改所有权？

Help everyone involved in the handover understand why there’s a system ownership change. This impacts the team engagement in the handover process.

帮助所有参与移交的人了解系统所有权变更的原因。这会影响团队在移交过程中的参与。

What’s best is to document the reasoning and to add additional information to the history of the system. In turn, this can reveal different things, such as if the ownership changed or underwent restructuring multiple times in a short period of time, or if the current organization isn't set up to own such a system or if the system doesn't belong to any team. Uncovering this information helps us ask important questions, such as is the system too complex, is it not in any team’s domain, or is it even needed anymore?

最好的办法是记录推理并在系统历史中添加额外的信息。反过来，这可以揭示不同的事情，例如所有权是否在短时间内多次更改或进行了重组，或者当前组织是否设置为拥有这样的系统，或者该系统是否不属于对任何球队。发现这些信息有助于我们提出重要问题，例如系统是否过于复杂，是否不在任何团队的领域内，或者是否不再需要它？

### What Does the System Do? What Problem Does It Solve? What Is the Vision?

### 系统是做什么的？它解决什么问题？什么是愿景？

Here, we’re looking to understand the system from a product perspective. It’s helpful to know some history about the system, how it evolved, and what its vision for the future is.

在这里，我们希望从产品的角度来理解系统。了解该系统的一些历史、它是如何演变的以及它对未来的愿景是有帮助的。

This could be a product session where product managers are involved. As an outcome of this session, it’d be nice to have a document to help onboard new people to the system.

这可能是产品经理参与的产品会议。作为本次会议的结果，最好有一份文件来帮助新人加入系统。

### High-Level Architecture of the System

### 系统的高级架构

Get an overview of the system, the main components, and their interfaces. A more detailed diagram will probably lead to more detailed discussions. It’s best to have an online diagram so that it’s easily available for future reference.

了解系统、主要组件及其接口。更详细的图表可能会导致更详细的讨论。最好有一个在线图表，以便将来可以轻松获取。

### Use, Availability, and Criticality of the System 

### 系统的使用、可用性和关键性

Get familiar with who’s using the system, what the criticality of the system is, and what it means if the system isn’t available. This is an opportunity to look into the available runbooks, metrics, and monitoring.

熟悉谁在使用系统，系统的重要性是什么，以及如果系统不可用意味着什么。这是一个查看可用运行手册、指标和监控的机会。

In cases when the previous and the new team owning the system are not part of the same [on-call rotation](https://developers.soundcloud.com/blog/building-a-healthy-on-call-culture), there **must** be an additional session where the system is introduced and explained to the engineers in the rotation group, so that all of them can respond to incidents related to the system. This helps prevent bigger outages and maintain the healthy on-call culture.

如果拥有该系统的前任和新团队不属于同一 [on-call 轮换](https://developers.soundcloud.com/blog/building-a-healthy-on-call-culture)， **必须**有一个额外的会议，向轮换组的工程师介绍和解释系统，以便他们都能对与系统相关的事件做出反应。这有助于防止更大的中断并保持健康的随叫随到文化。

### Maintenance

###  维护

As part of the maintenance of our system, we have daily, weekly, and monthly tasks that need to be executed by the team. In this section, we need to identify what those tasks are and what their periodicity is (if needed).

作为系统维护的一部分，我们有需要团队执行的每日、每周和每月任务。在本节中，我们需要确定这些任务是什么以及它们的周期性（如果需要）。

### Data Storage Overview

### 数据存储概述

Most of the systems have their own data storage. When taking ownership of a system, the new team takes ownership of that data and the infrastructure that comes with it. This is to get an overview of what, how, and where that data is stored and/or purged.

大多数系统都有自己的数据存储。在获得系统所有权时，新团队将获得该数据及其附带的基础设施的所有权。这是为了了解存储和/或清除数据的内容、方式和位置。

### Batch (Offline) Jobs Overview

### 批处理（离线）作业概述

Some systems are using data that’s an outcome of a batch job. Also, many systems have batch jobs that produce datasets to be consumed for analytics or reporting. Get an overview of the batch jobs, their usage, outcomes, and maintenance.

一些系统正在使用批处理作业的结果数据。此外，许多系统都有批处理作业，可以生成用于分析或报告的数据集。了解批处理作业、它们的使用、结果和维护。

### Decision History

### 决策历史

This is helpful to understand the architecture choices and the evolution of the system, as well as to learn about any constraints the system might have. It’s best if this is documented.

这有助于了解体系结构选择和系统的演变，以及了解系统可能具有的任何约束。最好有记录。

We recommend using [Lightweight Architecture Decisions Records](https://www.thoughtworks.com/de/radar/techniques/lightweight-architecture-decision-records), which has the following format: [Context, Decision, and Consequences]( https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions).

我们推荐使用[轻量级架构决策记录](https://www.thoughtworks.com/de/radar/techniques/lightweight-architecture-decision-records)，其格式如下：[Context, Decision, and Consequences]( https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)。

### Tech Debt

### 技术债务

Make sure you’re aware of the existing tech debt and you understand its implications. This is best if it’s documented.

确保您了解现有的技术债务并了解其影响。如果有记录，这是最好的。

### Known Bugs

### 已知错误

When ownership changes, user-facing bugs will be reported to the new team; however, many of them might already be known. Make sure you understand them and why they’re present. This helps not only by decreasing the time to investigate, but also by providing a good service to our users.

当所有权发生变化时，将向新团队报告面向用户的错误；然而，其中许多可能已经为人所知。确保您了解它们以及它们存在的原因。这不仅有助于减少调查时间，还有助于为我们的用户提供良好的服务。

### To Dos

### 待办事项

In addition to the above, here are some tasks (not in order and not complete, since they are SoundCloud specific) that can ease the ownership change:

- **Project ownership update on GitHub**
- **Grant permissions to the new team**
- **Update offline jobs configuration**
- **Local development**
   - Can the engineers build the project locally?
   - Does the system have integration tests? If so, do they run in a local environment? Does the new owner need any additional information to run them?
- **Deployment**
   - How is the system deployed?
   - Is there a CI/CD?
- **Monitoring and Alerting**
   - Check if the monitoring graphs need to be updated
   - Update the system-related runbooks
   - In case the runbook location changes, please reflect that change in the corresponding alerts to avoid broken links from the alerts
   - Add the system to a corresponding on-call group and have a knowledge transfer session
   - Update alerts
   - Update PagerDuty
   - Our suggestion is to update the on-call rotation at the end, once the team has gained sufficient knowledge and confidence in the system

除了上述任务之外，还有一些任务（不按顺序且不完整，因为它们是 SoundCloud 特定的）可以缓解所有权变更：

- **GitHub 上的项目所有权更新**
- **授予新团队权限**
- **更新离线作业配置**
- **本地开发**
  - 工程师可以在本地建造项目吗？
  - 系统是否有集成测试？如果是这样，它们是否在本地环境中运行？新所有者是否需要任何其他信息来运行它们？
- **部署**
  - 系统是如何部署的？
  - 有 CI/CD 吗？
- **监控和警报**
  - 检查监控图是否需要更新
  - 更新系统相关的运行手册
  - 如果 Runbook 位置发生变化，请在相应的警报中反映该变化，以避免警报中的链接断开
  - 将系统加入对应的on-call组，进行知识传授
  - 更新警报
  - 更新 PagerDuty
  - 我们的建议是在团队对系统获得足够的知识和信心后，在最后更新 on-call 轮换

💡 The list of things to do is quite long. Take your time with each task and don’t rush. Use the help of the previous team, and pay attention to details and to the alerts.

💡要做的事情清单很长。花点时间完成每项任务，不要着急。使用前一个团队的帮助，并注意细节和警报。

💡 You can use your project management tool (e.g. JIRA) to track the progress of the handover. That will help both involved teams stay up to date on the status of the handover, the next steps, and when it will be completed.

💡 您可以使用您的项目管理工具（例如 JIRA）来跟踪移交的进度。这将有助于两个相关团队及时了解移交的状态、后续步骤以及何时完成。

💡 If you’ve discovered other helpful tasks or topics, please update the guideline with them.

💡 如果您发现了其他有用的任务或主题，请与他们一起更新指南。

## Usage

##  用法

As the name suggests, the above document is a guideline, and it’s up to the parties involved in the system ownership change to decide **if they’re going to use it and how they’re going to use it**. 

顾名思义，上述文件是一个指南，由参与系统所有权变更的各方决定**他们是否要使用它以及他们将如何使用它**。

In most cases, ownership change is a collaborative process that enables the new owners to be motivated and have a solid understanding of the system to maintain and continue evolving it. In some cases, it can happen that there’s no one in the company that previously contributed to the system. However, even then, this guideline can help the team keep the focus on topics that are important to know, and not only on the codebase.

在大多数情况下，所有权变更是一个协作过程，它使新所有者能够受到激励并对系统有深入的了解，以维护和继续发展它。在某些情况下，公司中可能没有人以前为系统做出过贡献。然而，即便如此，该指南仍可以帮助团队将注意力集中在需要了解的重要主题上，而不仅仅是代码库。

The most important thing is to not be judgmental of the choices others made and understand that, at that time, it was the best decision. For example, instead of making statements like “You could have used A instead of B,” or “You could have done it like this,” or even the harsher “That is wrong!” or “That is a huge mistake!”, try to be curious and ask open-ended questions like “What made you use A?” or similar.

最重要的是不要评判别人所做的选择，并了解当时这是最好的决定。例如，与其说“你本可以用 A 而不是 B”或“你可以这样做”，甚至更严厉的“那是错的！”或“这是一个巨大的错误！”，试着保持好奇并提出开放式问题，例如“是什么让你使用 A？”或类似。

I’d also recommend that the team taking ownership takes the time to go through each of the topics and gain understanding and knowledge — not only from an engineering perspective, but from the perspective of the product. One might think they can copy the guideline and fill in the sections, and that writing everything they know of and handing it over to the new team will complete the transfer. I would argue that this isn’t the intention, and I don’t believe it will have the same positive impact that can be seen when doing this collaboratively and dedicating time to exploration.

我还建议拥有所有权的团队花时间浏览每个主题并获得理解和知识——不仅从工程角度，而且从产品的角度。有人可能认为他们可以复制指南并填写部分，写下他们知道的所有内容并将其交给新团队将完成转移。我认为这不是我们的意图，我不相信它会产生与合作进行并投入时间进行探索时可以看到的相同的积极影响。

Additionally, the guideline is meant to be a live document, updated as teams are learning through the process.

此外，该指南是一个实时文档，随着团队在整个过程中学习而更新。

## Side Effects / Other Impacts

## 副作用/其他影响

This guideline should inspire teams to have useful and up-to-date documentation; a README on how to contribute, test, and run locally; and high-level architecture diagrams. This helps not only when changing ownership, but also when onboarding new team members.

该指南应激励团队拥有有用且最新的文档；关于如何在本地贡献、测试和运行的自述文件；和高级架构图。这不仅有助于更改所有权，而且有助于新团队成员入职。

Furthermore, it’s important to embrace the use of architectural decision records and help to reason about them in the future.

此外，重要的是要接受架构决策记录的使用，并帮助将来对其进行推理。

## Summary

##  概括

This guideline exists to help engineering managers, product managers, and teams acknowledge that system ownership change is a process that should be well planned and done at a time that works best for everyone involved. It’s a process that requires effort and has its cost. However, it can inspire the organization to nurture a healthy engineering culture with a high bus factor and systems that are easy to maintain, evolve, and reason about.

本指南旨在帮助工程经理、产品经理和团队承认系统所有权变更是一个应该精心计划并在最适合所有相关人员的时间完成的过程。这是一个需要付出努力并付出代价的过程。但是，它可以激励组织培养具有高总线系数和易于维护、发展和推理的系统的健康工程文化。

- [Project Management](http://developers.soundcloud.com/blog/category/project%20management)
- [Engineering Management](http://developers.soundcloud.com/blog/category/engineering%20management)
- [Operations](http://developers.soundcloud.com/blog/category/operations) 

- [项目管理](http://developers.soundcloud.com/blog/category/project%20management)
- [工程管理](http://developers.soundcloud.com/blog/category/engineering%20management)
- [运营](http://developers.soundcloud.com/blog/category/operations)


