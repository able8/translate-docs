# Service Ownership: What It Really Means and How to Achieve It

# 服务所有权：它的真正含义以及如何实现它

Years ago, end-to-end software development involved dividing tasks based on where they fell in the system life cycle. One team wrote the code. Then another team deployed it to production. And yet another team monitored and maintained the service. This led to a lot of friction, needless handoffs, and bottlenecks.

多年前，端到端软件开发涉及根据任务在系统生命周期中的位置来划分任务。一个团队编写了代码。然后另一个团队将其部署到生产环境中。另一个团队负责监控和维护服务。这导致了很多摩擦、不必要的交接和瓶颈。

Then DevOps came to the rescue, promising to reduce handoffs and improve operations. But without the missing ingredient of service ownership, DevOps won’t be as powerful as it could be.

然后 DevOps 出手相救，承诺减少交接并改善运营。但是，如果缺少服务所有权的要素，DevOps 将不会像它应有的那样强大。

Now what does it mean to own a service? And how can we support this organizational pattern within our teams? While sharing ideas on how to [take back DevOps](http://www.opslevel.com/blog/taking-back-devops/), we talked a bit about service ownership. As a follow-up, we’ll now dive deeper into service ownership and how we can spread it across our teams.

现在拥有服务意味着什么？我们如何在团队中支持这种组织模式？在分享如何[收回 DevOps](http://www.opslevel.com/blog/taking-back-devops/) 的想法时，我们谈到了服务所有权。作为后续行动，我们现在将更深入地研究服务所有权以及如何将其分散到我们的团队中。

> Without the missing ingredient of **service ownership**,
>
> DevOps won’t be as **powerful** as it could be.

> 没有**服务所有权**的缺失成分，
>
> DevOps 不会像它应该的那样**强大**。

## What Does Service Ownership Mean?

## 服务所有权是什么意思？

To understand what service ownership provides, we’ll first want to understand the pain it solves.

要了解服务所有权提供了什么，我们首先要了解它解决的痛点。

### Without Service Ownership

### 没有服务所有权

To begin, let’s go back to the model mentioned in the intro where different teams were responsible for building, deploying, and managing a service.

首先，让我们回到介绍中提到的模型，其中不同的团队负责构建、部署和管理服务。

As you can imagine, these [silos](https://en.wikipedia.org/wiki/Silo_(software)) led to friction, handoffs, and bottlenecks.

可以想象，这些 [筒仓](https://en.wikipedia.org/wiki/Silo_(software)) 导致了摩擦、交接和瓶颈。

First, natural friction developed between development and ops teams. Development teams wanted to move features quickly, while operations teams wanted to slow the rate of change to ensure they weren’t buried in incidents and maintenance work. And release management played in the middle, trying to make both sides happy while following proper release protocol.

首先，开发和运维团队之间产生了自然摩擦。开发团队希望快速移动功能，而运营团队希望减慢更改速度，以确保它们不会被事件和维护工作所淹没。发布管理在中间发挥作用，在遵循适当的发布协议的同时，试图让双方都满意。

Second, handoffs between teams resulted in no one taking ownership of problems or improvements to the service’s overall life cycle. If a piece of the life cycle wasn’t theirs to own and improve, teams weren’t encouraged to look for solutions or improvements in that space.

其次，团队之间的交接导致没有人负责服务的整个生命周期的问题或改进。如果生命周期的一部分不属于他们自己和改进，则不鼓励团队在该领域寻找解决方案或改进。

For example, for many teams, observability that improved troubleshooting was always an afterthought. Furthermore, engineers didn’t always think about making their code easier to debug because they weren’t the ones who had to do that at three in the morning when the pager went off. And operations, the folks who could really use improved troubleshooting capabilities, didn’t have the autonomy to improve things themselves.

例如，对于许多团队来说，改进故障排除的可观察性总是事后才想到的。此外，工程师并不总是考虑让他们的代码更容易调试，因为他们不是那些在寻呼机响起的凌晨三点必须这样做的人。而运营，真正可以使用改进的故障排除功能的人，没有自主权来改进自己的事情。

If all that wasn’t painful enough, bottlenecks formed because of these silos. Development teams waited on release management to ship the code. Release management was waiting on approval from someone to deploy the code. And once the code deployed, operations teams waited on development teams to fix bugs or improve reliability. And those items never came quickly enough for the ops team.

如果这一切还不够痛苦，瓶颈就会因为这些孤岛而形成。开发团队等待发布管理人员交付代码。发布管理正在等待某人批准部署代码。一旦代码部署完毕，运营团队就会等待开发团队修复错误或提高可靠性。对于运维团队来说，这些项目从来没有来得足够快。

So how does service ownership help? Service ownership takes the responsibility of those three teams and combines it into one team’s responsibility. One team holds the responsibility for the product and not just a piece of the pipeline.

那么服务所有权有什么帮助呢？服务所有权将这三个团队的责任合并为一个团队的责任。一个团队负责产品，而不仅仅是管道的一部分。

Perhaps you’ve heard the phrase “build it, ship it, run it.” That’s the two-second description of DevOps and service ownership. To expand on that, let’s break those components down.

也许您听说过“构建它、运送它、运行它”这句话。这是对 DevOps 和服务所有权的两秒钟描述。为了扩展这一点，让我们分解这些组件。

> One team holds the **responsibility** for the product and not just a **piece** of the pipeline.
>

> 一个团队对产品负责 不仅仅是管道的**部分**。
>

### Build It

### 构建它

The phrase “build it” seems simple enough. Our engineering teams all build software. So does this look the same for teams that have service ownership and those that don’t?

短语“构建它”似乎很简单。我们的工程团队都在构建软件。那么对于拥有服务所有权的团队和没有服务所有权的团队来说，这看起来是一样的吗？

Not necessarily. When a team utilizes service ownership, they also gain autonomy and influence over how to build their product.

不必要。当团队利用服务所有权时，他们还获得了对如何构建产品的自主权和影响力。

In fact, when teams own the whole service, you’ll find that they appreciate the value of building services with appropriate documentation, thorough tests, and observability practices more than when separate teams operate or run the services.

事实上，当团队拥有整个服务时，您会发现他们比单独的团队操作或运行服务更欣赏构建具有适当文档、全面测试和可观察性实践的服务的价值。

### Ship It

###  装运它

Next, when a team owns a service, they own the responsibility of shipping it to production. 

接下来，当一个团队拥有一项服务时，他们有责任将它运送到生产环境中。

This doesn't mean that every team needs a [CI/CD](https://www.redhat.com/en/topics/devops/what-is-ci-cd) guru, with every team re-creating the basics of continuous deployment and automated compliance checks. It could be that the team uses common frameworks, patterns, and libraries provided by others.

这并不意味着每个团队都需要一个 [CI/CD](https://www.redhat.com/en/topics/devops/what-is-ci-cd) 大师，每个团队都需要重新创建基础知识持续部署和自动化合规性检查。可能是团队使用了其他人提供的通用框架、模式和库。

However, the team does own the responsibility of getting code in production efficiently. They should also understand the pipelines and fix things when they go wrong. Overall, the goal is not to throw problems over the wall at the release management team, but to take ownership of shipping the product to production.

但是，团队确实有责任有效地将代码投入生产。他们还应该了解管道并在出错时进行修复。总体而言，我们的目标不是将问题抛给发布管理团队，而是拥有将产品运送到生产环境的所有权。

### Run It

###  运行

Here, the team that writes the service is also the team that’s best equipped to run the service. They know how to resolve issues quickly and make changes that will prevent similar future issues.

在这里，编写服务的团队也是最有能力运行服务的团队。他们知道如何快速解决问题并进行更改以防止将来出现类似的问题。

When we run our own applications, we also realize the importance of operationalizing the product. So we’ll write code in a way that makes it more maintainable, observable, and debuggable. Because we’ll be the ones who have to fix it when it breaks.

当我们运行自己的应用程序时，我们也意识到可操作化产品的重要性。因此，我们将以一种更易于维护、可观察和可调试的方式编写代码。因为当它损坏时，我们将成为必须修复它的人。

## Getting Started

##  入门

Next, what first steps can you take toward driving service ownership in your organization? In short, we need to move away from telling teams what to do or how to do it and move toward looking at outcomes necessary for customer satisfaction. Furthermore, when multiple teams contribute to an initiative, we need to point them all toward the same outcome. Then we must encourage them to collaborate and work through integrations and dependencies together, always pointing toward the same goal.

接下来，您可以采取哪些第一步来推动组织中的服务所有权？简而言之，我们需要摆脱告诉团队该做什么或如何做，而转向关注客户满意度所需的结果。此外，当多个团队为一项计划做出贡献时，我们需要将它们指向相同的结果。然后，我们必须鼓励他们协作并共同解决集成和依赖关系，始终指向同一个目标。

> When **multiple teams** contribute to an initiative, we need
>
> to point them all toward the **same outcome**.

> 当**多个团队**为一项计划做出贡献时，我们需要
>
> 将它们都指向**相同的结果**。

So how can we do that?

那我们怎么做呢？

First, share expectations around service ownership clearly. Talk about where you see the organization is at today and what you’re working toward. And mention the benefits that not only the customer or organization receives, but also the benefits that the development teams should expect regarding reduced bottlenecks. Then ask for feedback and help in getting there. And be clear that all the teams own the problem and will all contribute to the solution.

首先，明确分享对服务所有权的期望。谈论您认为该组织目前所处的位置以及您正在努力实现的目标。并提及不仅客户或组织获得的好处，而且开发团队应该期望减少瓶颈的好处。然后寻求反馈并帮助实现目标。并明确所有团队都拥有问题，并且都将为解决方案做出贡献。

Once people align on expectations, look at gaps in tools, skills, and knowledge. What do your teams need to take ownership of their service? Do they need education regarding security so that they can prioritize concerns properly? Or do they need tools to automate scans or reduce unnecessary toil?

一旦人们与期望保持一致，请查看工具、技能和知识方面的差距。您的团队需要什么来掌控他们的服务？他们是否需要进行安全方面的教育，以便他们可以适当地优先考虑问题？或者他们是否需要工具来自动化扫描或减少不必要的工作？

Assess skills and seed teams with people who can build up release management and operational knowledge.

与可以建立发布管理和操作知识的人员一起评估技能和种子团队。

**Service Ownership, Solved**

**服务所有权，已解决**

Say goodbye to stale spreadsheets or wikis! See **OpsLevel** in action to accelerate your Service Ownership journey.

告别陈旧的电子表格或维基！查看 **OpsLevel** 的实际操作，以加快您的服务所有权之旅。

## Common Mistakes in Service Ownership

## 服务所有权的常见错误

Now that we know from a high level what service ownership looks like, let’s consider some common mistakes in the implementation of the pattern.

现在我们从高层次了解了服务所有权是什么样的，让我们考虑一下在模式实现中的一些常见错误。

### Accountability Without Autonomy

### 没有自主权的问责制

If the team owns the service, they should have the autonomy and accountability to build, ship, and run it as they see fit. But autonomy and accountability require more than lip service. If you tell teams they’re accountable but then don’t let them have control over their own backlog, you’ve probably given the team accountability without autonomy.

如果团队拥有该服务，他们就应该拥有自主权和责任感来按照他们认为合适的方式构建、发布和运行它。但是自主权和问责制需要的不仅仅是口头上的服务。如果你告诉团队他们要负责，但又不让他们控制自己的积压工作，那么你很可能在没有自主权的情况下给了团队责任。

For example, management may tell teams to “drop everything” and focus on one specific task or feature. And this task must be completed quickly, to the detriment of all others. This indicates that the team doesn’t have the autonomy to prioritize tasks or assess needs.

例如，管理层可能会告诉团队“放下一切”并专注于一项特定任务或功能。而这项任务必须迅速完成，这对所有其他人都是不利的。这表明团队没有自主权来确定任务的优先级或评估需求。

Service owners need to have the autonomy to prioritize the reduction of their own operational pain and toil, and balance these needs against feature enhancements and innovation. Additionally, the taking on of any new technical debt - usually through shortcuts due to strict time restraints - should be something the team takes on knowingly and willingly, and with an associated plan for dealing with it. 

服务所有者需要有自主权来优先考虑减少他们自己的运营痛苦和辛劳，并在这些需求与功能增强和创新之间取得平衡。此外，承担任何新的技术债务——通常是由于严格的时间限制而通过捷径——应该是团队有意和自愿承担的事情，并有相关的处理计划。

Therefore, instead of demanding specific work be done in a specific order, have the team focus on a target or outcome. For example, if a team focuses on helping the customer with a specific problem, they’re better equipped to prioritize tasks related to features, security, and tech debt.

因此，与其要求按特定顺序完成特定工作，不如让团队专注于一个目标或结果。例如，如果一个团队专注于帮助客户解决特定问题，他们就更有能力优先考虑与功能、安全性和技术债务相关的任务。

> Service owners need to have the **autonomy** to prioritize the reduction of their own operational pain and toil, and **balance** these needs against feature enhancements and innovation.

> 服务所有者需要拥有 **自主性** 来优先考虑减少他们自己的运营痛苦和辛苦，并**平衡**这些需求与功能增强和创新。

### Lack of Support or Resources

### 缺乏支持或资源

If you take an existing software team that has relied on other teams to ship and run their product and tell them that this is all now their responsibility, you’ll lose trust. And they’ll lose sleep.

如果你让一个依赖其他团队来交付和运行他们的产品的现有软件团队告诉他们现在这一切都是他们的责任，你就会失去信任。他们会失眠。

Instead, work to introduce change gradually. Start with awareness, letting the development team understand what it takes to truly own their product. And seed the team with people from operations or release management so they have someone there to help with the journey and expertise.

相反，努力逐步引入变化。从意识开始，让开发团队了解真正拥有他们的产品需要什么。并将来自运营或发布管理人员的人员作为团队种子，以便他们在那里有人帮助完成旅程和专业知识。

### Unclear Expectations

### 不明确的期望

When teams rely on each other in a complex distributed ecosystem, unclear expectations can hurt trust and collaboration between teams. If our team has a dependency on another team but ownership and performance expectations aren’t clear, this can result in blame games and finger-pointing. Suppose one service is down or degraded due to a failure in a dependency. In that case, we need a clear line of sight into escalation processes, expectation on service-level objectives (SLOs), and an understanding of what that dependency team offers.

当团队在复杂的分布式生态系统中相互依赖时，不明确的期望会损害团队之间的信任和协作。如果我们的团队依赖另一个团队，但所有权和绩效期望不明确，这可能会导致相互指责和指责。假设一项服务由于依赖项失败而关闭或降级。在这种情况下，我们需要清楚地了解升级流程、对服务级别目标 (SLO) 的期望，并了解依赖团队提供的服务。

Services should have clearly defined objectives, contact information, and access to metrics and status all in one place. If your teams struggle with finding out what to expect from another service, who to contact, or what state the service is in, consider tools like OpsLevel to provide a better service ownership experience to your organization.

服务应该在一个地方具有明确定义的目标、联系信息以及对指标和状态的访问。如果您的团队在寻找对其他服务的期望、与谁联系或服务处于什么状态方面苦苦挣扎，请考虑使用 OpsLevel 之类的工具为您的组织提供更好的服务所有权体验。

> When teams rely on each other in a **complex distributed ecosystem**,
>
> unclear expectations can **hurt trust** and collaboration between teams.

> 当团队在**复杂的分布式生态系统**中相互依赖时，
>
> 不明确的期望会**损害团队之间的信任**和协作。

### Invisible Ownership

### 无形的所有权

Sometimes we rely on word of mouth or team-specific methods of identifying service owners. Or disparate spreadsheets and wikis hold out-of-date information that’s difficult to find. Sometimes service ownership info hides away in hard-to-find architecture documents or data stores that lead you on a treasure hunt.

有时，我们依靠口耳相传或特定于团队的方法来识别服务所有者。或者，不同的电子表格和维基包含难以找到的过时信息。有时，服务所有权信息隐藏在难以找到的架构文档或数据存储中，引导您进行寻宝。

This results in frustration during incident response or even integration discussions.

这会导致在事件响应甚至集成讨论期间感到沮丧。

To reduce the complexity, provide teams with one central and easy-to-use place to get the information they need. For this common problem, OpsLevel can help provide visibility through their [service catalog](http://www.opslevel.com/landing/microservice-catalog/).

为了降低复杂性，为团队提供一个集中且易于使用的地方来获取他们需要的信息。对于这个常见问题，OpsLevel 可以通过他们的 [服务目录](http://www.opslevel.com/landing/microservice-catalog/) 帮助提供可见性。

Wherever the information lives, make sure that everyone knows where to find it and that the data stays up to date.

无论信息在哪里，请确保每个人都知道在哪里可以找到它，并且数据保持最新。

## Wrapping It Up

## 包装它

Service ownership provides many benefits, like encouraging teams to incorporate operational excellence early in the life cycle. Additionally, it gives the team autonomy to do the right thing and satisfy customer needs based on their knowledge and expertise.

服务所有权提供了许多好处，例如鼓励团队在生命周期的早期整合卓越运营。此外，它还赋予团队自主权，可以根据他们的知识和专业知识做正确的事情并满足客户需求。

To make service ownership easy to see, consider OpsLevel and how you can improve visibility and expectations. [Request a demo](http://www.opslevel.com/request-demo/) to see how OpsLevel can drive ownership today.

为了使服务所有权易于查看，请考虑 OpsLevel 以及如何提高可见性和期望。 [请求演示](http://www.opslevel.com/request-demo/) 了解 OpsLevel 如何推动今天的所有权。

This post was written by Sylvia Fronczak. [Sylvia](https://sylviafronczak.com/) is a software developer and SRE manager that has worked in various industries with various software methodologies. She’s currently focused on design practices that the whole team can own, understand, and evolve over time. 

这篇文章是由 Sylvia Fronczak 撰写的。 [Sylvia](https://sylviafronczak.com/) 是一名软件开发人员和 SRE 经理，曾在各个行业使用各种软件方法。她目前专注于整个团队可以拥有、理解和随着时间发展的设计实践。

