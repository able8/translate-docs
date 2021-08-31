# Kubernetes Is Not Your Platform, It's Just the Foundation

# Kubernetes 不是你的平台，它只是基础

Mar 19, 2021

2021 年 3 月 19 日

### Key Takeaways

### 关键要点

- Kubernetes itself is not a platform but only the foundational element of an ecosystem not only of tools, and services, but also offering support as part of a compelling internal product.
- Platform teams should provide useful abstractions of Kubernetes complexities to reduce cognitive load on stream teams.
- The needs of platform users change and a platform team needs to ease their journey forward.
- Changes to the platform should meet the needs of stream teams, prompting a collaborative discovery phase with the platform team followed by stabilizing the new features (or service) before they can be consumed by other stream teams in a self-service fashion.
- A team-focused Kubernetes adoption requires an assessment of cognitive load and tradeoffs, clear platform and service definitions, and defined team interactions.

- Kubernetes 本身不是一个平台，而只是一个生态系统的基础元素，不仅包括工具和服务，而且作为引人注目的内部产品的一部分提供支持。
- 平台团队应提供有用的 Kubernetes 复杂性抽象，以减少流团队的认知负担。
- 平台用户的需求发生变化，平台团队需要简化他们的前进旅程。
- 平台的变化应满足流团队的需求，促使与平台团队进行协作发现阶段，然后稳定新功能（或服务），然后其他流团队以自助方式使用它们。
- 以团队为中心的 Kubernetes 采用需要评估认知负载和权衡、明确的平台和服务定义以及明确的团队交互。

I read a lot of articles and see presentations on impressive Kubernetes tools, automation, and different ways to use the technology but these often offer little context about the organization other than the name of the teams involved.

我阅读了很多文章，并看到了有关令人印象深刻的 Kubernetes 工具、自动化和使用该技术的不同方式的介绍，但除了所涉及团队的名称之外，这些通常很少提供有关组织的背景信息。

We can’t really judge the success of technology adoption, particularly Kubernetes, if we don't know who asked for it, who needs it, and who is implementing it in which ways. We need to know how teams are buying into this new offering — or not.

如果我们不知道谁要求它、谁需要它以及谁以何种方式实施它，我们就无法真正判断技术采用的成功与否，尤其是 Kubernetes。我们需要知道团队如何购买这种新产品——或者不购买。

In [Team Topologies](https://teamtopologies.com/book), which I wrote with [Matthew Skelton](https://www.infoq.com/profile/Matthew-Skelton/), we talk about fundamental types of teams, their expected behaviors and purposes, and perhaps more importantly about the interactions among teams. Because that's what's going to drive adoption of technology. If we don't take that into consideration, we might fall into a trap and build technologies or services that no one needs or which have no clear use case.

在我与 [Matthew Skelton](https://www.infoq.com/profile/Matthew-Skelton/) 共同撰写的 [Team Topologies](https://teamtopologies.com/book) 中，我们讨论了基本类型团队、他们的预期行为和目的，也许更重要的是团队之间的互动。因为这将推动技术的采用。如果我们不考虑这一点，我们可能会陷入陷阱并构建没有人需要或没有明确用例的技术或服务。

#### Related Sponsored Content

#### 相关赞助内容

- ##### [Kubernetes Up & Running - Download the eBook (By O'Reilly)](http://www.infoq.com/infoq/url.action?i=e1b2b077-e265-40a5-8f12-593f0f6f6baa&t =f)

- ##### [Kubernetes 启动和运行 - 下载电子书（O'Reilly 着）](http://www.infoq.com/infoq/url.action?i=e1b2b077-e265-40a5-8f12-593f0f6f6baa&t =f)

#### Related Sponsor

#### 相关赞助商

[![](https://res.infoq.com/sponsorship/topic/22bdfa2a-7b5c-4be5-b60c-93141547ff5a/VMWareTanzuRSBLogo-1589361540634-1612367286968.png)](http://www.infoq.com/infoq/url.action?i=9e6eac17-4b86-410d-8cf9-9ce2e2a28863&t=f)

/url.action?i=9e6eac17-4b86-410d-8cf9-9ce2e2a28863&t=f)

### [**Free your apps. Simplify your ops.**](http://www.infoq.com/infoq/url.action?i=4e465f12-aac3-4817-8b96-89c289f60a82&t=f)

### [**释放您的应用程序。简化您的操作。**](http://www.infoq.com/infoq/url.action?i=4e465f12-aac3-4817-8b96-89c289f60a82&t=f)

Team cognitive load has a direct effect on platform success and team interaction patterns can help reduce unnecessary load on product teams. We’ll look at how platforms can minimize cognitive load and end with ideas on how to take a team-centric approach to Kubernetes adoption.

团队认知负荷对平台成功有直接影响，而团队交互模式有助于减少产品团队不必要的负担。我们将研究平台如何最大限度地减少认知负荷，并最终提出如何采用以团队为中心的方法来采用 Kubernetes。

## Is Kubernetes a platform?

## Kubernetes 是一个平台吗？

In 2019, Melanie Cebula, an infrastructure engineer at Airbnb, gave an excellent talk at Qcon. As DevOps lead editor for InfoQ, I wrote [the story about it](https://www.infoq.com/news/2019/03/airbnb-kubernetes-workflow/). In the first week online, this story had more than 23,000 page views. It was my first proper viral story, and I tried to understand why this one got so much attention. Yes, Airbnb helped, and Kubernetes is all the rage, but I think the main factor was that the story was about simplifying Kubernetes adoption at a large scale, for thousands of engineers.

2019 年，Airbnb 基础设施工程师 Melanie Cebula 在 Qcon 做了精彩演讲。作为 InfoQ 的 DevOps 主编，我写了 [关于它的故事](https://www.infoq.com/news/2019/03/airbnb-kubernetes-workflow/)。在上线的第一周，这个故事的页面浏览量就超过了 23,000。这是我第一个真正的病毒式故事，我试图理解为什么这个故事会受到如此多的关注。是的，Airbnb 提供了帮助，Kubernetes 风靡一时，但我认为主要因素是这个故事是关于为数千名工程师大规模简化 Kubernetes 的采用。

The point is that many developers and many engineers find it complicated and hard to adopt Kubernetes and change the way they work with the APIs and the artifacts they need to produce to use it effectively. 
关键是，许多开发人员和工程师发现采用 Kubernetes 并改变他们使用 API 的方式以及他们需要生产以有效使用它的工件的方式既复杂又困难。
The term “platform” has been overloaded with a lot of different meanings. Kubernetes is a platform in the sense that it helps us deal with the complexity of operating microservices. It helps provide better abstractions for deploying and running our services. That's all great, but there's a lot more going on. We need to think about how to size our hosts, how and when to create/destroy clusters, and how to update to new Kubernetes versions and who that will impact. We need to decide on how to isolate different environments or applications, with namespaces, clusters, or whatever it might be. Anyone who has worked with Kubernetes can add to this list: perhaps worrying about security, for example. A lot of things need to happen before we can use Kubernetes as a platform.

“平台”一词已经被赋予了许多不同的含义。 Kubernetes 是一个平台，它帮助我们处理操作微服务的复杂性。它有助于为部署和运行我们的服务提供更好的抽象。这一切都很棒，但还有很多事情要做。我们需要考虑如何调整我们的主机大小，如何以及何时创建/销毁集群，以及如何更新到新的 Kubernetes 版本以及这会影响谁。我们需要决定如何隔离不同的环境或应用程序，使用命名空间、集群或任何可能的东西。任何使用 Kubernetes 的人都可以添加到此列表中：例如，可能担心安全性。在我们可以使用 Kubernetes 作为平台之前，需要做很多事情。

One problem is that the boundaries between roles are often unclear in an organization. Who is the provider? Who is the owner responsible for doing all of this? Who is the consumer? What are the teams that will consume the platform? With blurry boundaries, it's complicated to understand who is responsible for what and how our decisions will affect other teams.

一个问题是，在一个组织中，角色之间的界限往往不明确。谁是提供者？谁负责做这一切？谁是消费者？哪些团队将使用该平台？由于边界模糊，很难理解谁负责什么以及我们的决定将如何影响其他团队。

[Evan Bottcher](https://www.linkedin.com/in/evanbottcher/) defines a digital platform as “a foundation of self-service APIs, tools, and services, but also knowledge and support, and everything arranged as a compelling internal product.” We know that self-service tools and APIs are important. They're critical in allowing a lot of teams to be more independent and to work autonomously. But I want to bring attention to his mention of “knowledge and support”. That implies that we have teams running the platform that are focused on helping product teams understand and adopt it, besides providing support when problems arise in the platform.

[Evan Bottcher](https://www.linkedin.com/in/evanbottcher/) 将数字平台定义为“自助式 API、工具和服务的基础，同时也是知识和支持的基础，一切安排为引人注目的内部产品。”我们知道自助服务工具和 API 很重要。它们对于让许多团队更加独立和自主工作至关重要。但我想提请注意他提到的“知识和支持”。这意味着我们拥有运行平台的团队，除了在平台出现问题时提供支持外，还专注于帮助产品团队理解和采用它。

Another key aspect is to understand the platform as “a compelling internal product”, not a mandatory platform with shared services that we impose on everyone else. As an industry, we've been doing that for a long time, and it simply doesn't work very well. It often creates more pain than the benefits it provides for the teams forced to use a platform that is supposed to be a silver bullet. And we know silver bullets don’t exist.

另一个关键方面是将平台理解为“引人注目的内部产品”，而不是我们强加给其他人的具有共享服务的强制性平台。作为一个行业，我们已经这样做了很长时间，但效果并不好。对于被迫使用应该是灵丹妙药的平台的团队来说，它通常会带来更多的痛苦。我们知道银弹不存在。

We have to think about the platform as a product. It's meant for our internal teams, but it's still a product. We're going to see what that implies in practice.

我们必须将平台视为一种产品。它适用于我们的内部团队，但它仍然是一个产品。我们将看看这在实践中意味着什么。

The key idea is that Kubernetes by itself is not a platform. It's a foundation. Yes, it provides all this great functionality — autoscaling, self-healing, service discovery, you name it — but a good product is more than just a set of features. We need to think about how easy it is to adopt, its reliability, and the support for the platform.

关键思想是 Kubernetes 本身并不是一个平台。这是一个基础。是的，它提供了所有这些强大的功能——自动扩展、自我修复、服务发现，你能想到的——但一个好的产品不仅仅是一组功能。我们需要考虑采用它的难易程度、可靠性以及对平台的支持。

A good platform, as Bottcher says, should create a path of least resistance. Essentially, the right thing to do should be the easiest thing to do with the platform. We can’t just say that whatever Kubernetes does is the right thing. The right thing depends on your context, on the teams that are going to consume the platform, which services they need the most, what kind of help they need to onboard, and so on.

正如 Bottcher 所说，一个好的平台应该创造一条阻力最小的路径。从本质上讲，正确的事情应该是平台上最容易做的事情。我们不能只说 Kubernetes 所做的一切都是正确的。正确的做法取决于您的背景、将要使用该平台的团队、他们最需要哪些服务、他们需要什么样的帮助，等等。

One of the hard things about platforms is that the needs of the internal teams are going to change, with respect to old and new customers. Teams that are consuming the platform are probably going to have more specific requests and requirements over time. At the same time, the platform must remain understandable and usable for new teams or new engineers adopting it. The needs and the technology ecosystem keep evolving.

平台的难点之一是内部团队的需求将发生变化，无论是新老客户。随着时间的推移，使用该平台的团队可能会有更具体的请求和要求。同时，该平台必须对采用它的新团队或新工程师保持可理解和可用。需求和技术生态系统不断发展。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/54image001-1615999814012.jpg)

**Figure 1: Avoid creating isolated teams when bringing Kubernetes into the organization.**

**图 1：在将 Kubernetes 引入组织时避免创建孤立的团队。**

## Team cognitive load

##团队认知负荷

Adopting Kubernetes is not a small change. It's more like an elephant charging into a room. We have to adopt this technology in a way that does not cause more pain than the benefits it brings. We don't want to end up with something like Figure 1, that resembles life before the DevOps movement, with isolated groups. It’s the Kubernetes team now rather than the operations team, but if we create an isolationist approach where one group makes decisions independent from the other group, we’re going to run into the same kinds of problems.

采用 Kubernetes 是一个不小的变化。这更像是一头大象冲进房间。我们必须以一种不会造成比带来的好处更多的痛苦的方式采用这项技术。我们不希望以图 1 那样的方式结束，它类似于 DevOps 运动之前的生活，具有孤立的群体。现在是 Kubernetes 团队而不是运营团队，但是如果我们创建一种孤立主义的方法，其中一个团队独立于另一个团队做出决策，我们将遇到同样的问题。

If we make platform decisions without considering the impact on consumers, we're going to increase pain in the form of their cognitive load, the amount of effort our teams must put out to understand and use the platform. 
如果我们在不考虑对消费者的影响的情况下做出平台决策，我们将增加他们认知负担的形式的痛苦，我们的团队必须为理解和使用平台付出的努力。
[John Sweller](https://en.wikipedia.org/wiki/John_Sweller) formulated cognitive load as “the total amount of mental effort being used in the working memory”. We can break this down into three types of cognitive load: intrinsic, extraneous, and germane.

[John Sweller](https://en.wikipedia.org/wiki/John_Sweller) 将认知负荷表述为“工作记忆中使用的脑力劳动总量”。我们可以将其分解为三种类型的认知负荷：内在的、外在的和密切的。

We can map these types of cognitive load in software delivery: intrinsic cognitive load is the skills I need to do my work; extraneous load is the mechanics, things that need to happen to deliver the value; and germane is the domain focus.

我们可以在软件交付中映射这些类型的认知负荷：内在认知负荷是我完成工作所需的技能；额外的负载是机制，需要发生以传递价值的事情；和germane 是领域焦点。

Intrinsic cognitive load, if I’m a Java developer, is knowing how to write classes in Java. If I don't know how to do that, this takes effort. I have to Google it or I have to try to remember it. But we know how to minimize intrinsic cognitive load. We can have classical training, pair programming, mentoring, code reviews, and all techniques that help people improve their skills.

如果我是一名 Java 开发人员，内在认知负荷是知道如何用 Java 编写类。如果我不知道该怎么做，这需要努力。我必须谷歌它或者我必须尝试记住它。但我们知道如何最大限度地减少内在认知负荷。我们可以进行经典培训、结对编程、指导、代码审查以及所有帮助人们提高技能的技术。

Extraneous cognitive load includes any task needed to deliver the work I'm doing to the customers or to production. Having to remember how to deploy this application, how to access a staging environment, or how to clean up test data are all things that are not directly related to the problem to solve but are things that I need to get done. Team topologies and platform teams in particular can minimize extraneous cognitive load. This is what we're going to explore throughout the rest of this article.

外部认知负荷包括将我正在做的工作交付给客户或生产所需的任何任务。必须记住如何部署此应用程序、如何访问暂存环境或如何清理测试数据，这些都与要解决的问题没有直接关系，但都是我需要完成的事情。团队拓扑和平台团队尤其可以最大限度地减少无关的认知负担。这就是我们将在本文的其余部分探索的内容。

Finally, germane cognitive load is knowledge in my business domain or problem space. For example, if I'm working in private banking, I need to know how bank transfers work. The point of minimizing extraneous cognitive load is to free up as much memory available for focus on the germane cognitive load.

最后，密切相关的认知负荷是我的业务领域或问题空间中的知识。例如，如果我在私人银行工作，我需要知道银行转账是如何运作的。最小化无关认知负荷的目的是释放尽可能多的可用记忆，以专注于密切相关的认知负荷。

Jo Pearce’s “Hacking Your Head” articles and [presentations](https://www.slideshare.net/JoPearce5/hacking-your-head-managing-information-overload-extended) go deeper into this.

Jo Pearce 的“Hacking Your Head”文章和[演示文稿](https://www.slideshare.net/JoPearce5/hacking-your-head-managing-information-overload-extended) 对此进行了更深入的探讨。

The general principle is to be mindful of the impact of your platform choices on your teams’ cognitive loads.

一般原则是注意平台选择对团队认知负荷的影响。

## Case studies

＃＃ 实例探究

I mentioned Airbnb engineer Melanie Cebula’s excellent talk at Qcon so let’s look at how that company reduced the cognitive load of their development teams.

我提到了 Airbnb 工程师 Melanie Cebula 在 Qcon 上的精彩演讲，让我们看看这家公司如何减少开发团队的认知负担。

“The best part of my day is when I update 10 different YAML files to deploy a one-line code change,” said no one, ever. Airbnb teams were feeling this sort of pain as they embarked on their Kubernetes journey. To reduce this cognitive load, they created a simple command-line tool, kube-gen, which allows the application teams or service teams to focus on a smaller set of configurations and details, which are specific to their own project or services. A team needs to configure only those settings like files or volumes specifically related to the germane aspect of their application. The kube-gen tool then generates the boilerplate code configuration for each environment, in their case: production, canary, and development environments. This makes it much easier for development teams to focus on the germane parts of their work.

“我一天中最棒的部分是更新 10 个不同的 YAML 文件以部署一行代码更改，”从来没有人说过。 Airbnb 团队在开始他们的 Kubernetes 之旅时感受到了这种痛苦。为了减少这种认知负担，他们创建了一个简单的命令行工具 kube-gen，它允许应用程序团队或服务团队专注于更小的配置和细节，这些配置和细节特定于他们自己的项目或服务。团队只需要配置那些与其应用程序密切相关的设置，如文件或卷。然后 kube-gen 工具为每个环境生成样板代码配置，在它们的情况下：生产、金丝雀和开发环境。这使得开发团队更容易专注于他们工作的相关部分。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/23image002-1615999815140.jpg)

**Figure 2: Airbnb uses kube-gen to simplify their Kubernetes ecosystem.**

**图 2：Airbnb 使用 kube-gen 来简化他们的 Kubernetes 生态系统。**

As Airbnb did, we essentially want to clarify the boundaries of the services provided by the platform and provide good abstractions so we reduce the cognitive load on each service team.

正如 Airbnb 所做的那样，我们本质上希望澄清平台提供的服务的边界并提供良好的抽象，从而减少每个服务团队的认知负担。

In Team Topologies, we talk about four types of teams, shown in Figure 3. Stream-aligned teams provide the end-customer value, they’re the heartbeat that delivers business value. The three other types of teams provide support and help reduce cognitive load. The platform team shields the details of the lower-level services that these teams need to use for deployment, monitoring, CI/CD, and other lifecycle supporting services.

在团队拓扑中，我们讨论了四种类型的团队，如图 3 所示。流对齐的团队提供最终客户价值，他们是提供业务价值的心跳。其他三种类型的团队提供支持并帮助减少认知负荷。平台团队屏蔽了这些团队需要用于部署、监控、CI/CD 和其他生命周期支持服务的较低级别服务的细节。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/44image003-1615999812928.jpg)

**Figure 3: Four types of teams.** 
**图 3：四种类型的团队。**
The stream-aligned team resembles what organizations variously call a product team, DevOps team, or build-and-run team — these teams have end-to-end ownership of the services that they deliver. They have runtime ownership, and they can take feedback from monitoring or live customer usage for improving the next iteration of their service or application. We call this a “stream-aligned team” for two reasons. First, “product” is an overloaded term; as our systems become more and more complex, the less precise it is to define a standalone product. Second, we want to acknowledge the different types of streams beyond just the business-value streams; it can be compliance, specific user personas, or whatever makes sense for aligning a team with the value it provides.

流对齐团队类似于组织对产品团队、DevOps 团队或构建和运行团队的不同称呼——这些团队对他们提供的服务拥有端到端的所有权。他们拥有运行时所有权，并且可以从监控或实时客户使用中获取反馈，以改进其服务或应用程序的下一次迭代。我们将其称为“流对齐团队”有两个原因。首先，“产品”是一个重载的术语；随着我们的系统变得越来越复杂，定义一个独立的产品就越不精确。其次，我们要承认不同类型的流，而不仅仅是业务价值流；它可以是合规性、特定的用户角色，或者任何使团队与其提供的价值保持一致的有意义的事情。

Another case study is uSwitch, which helps users in the UK compare utility providers and home services and makes it easy to switch between them.

另一个案例研究是 uSwitch，它帮助英国的用户比较公用事业提供商和家庭服务，并可以轻松地在它们之间进行切换。

A couple of years ago, [Paul Ingles](https://www.linkedin.com/in/pingles/), head of engineering at uSwitch, wrote “ [Convergence to Kubernetes](https://pingles.medium.com /convergence-to-kubernetes-137ffa7ea2bc)”, which brought together the adoption of Kubernetes technology, how it helps or hurts teams and their work, and data for meaningful analysis. Figure 4, from that article, measures all the different teams’ low-level AWS service calls at uSwitch.

几年前，uSwitch 的工程主管 [Paul Ingles](https://www.linkedin.com/in/pingles/) 写道：“[Convergence to Kubernetes](https://pingles.medium.com /convergence-to-kubernetes-137ffa7ea2bc)”，它汇集了 Kubernetes 技术的采用、它如何帮助或损害团队及其工作，以及用于有意义分析的数据。来自那篇文章的图 4 测量了所有不同团队在 uSwitch 上的低级别 AWS 服务调用。

When uSwitch started, every team was responsible for their own service, and they were as autonomous as possible. Teams were responsible for creating their own AWS accounts, security groups, networking, etc. uSwitch noticed that the number of calls to these services was increasing, correlated with a feeling that teams were getting slower at delivering new features and value for the business.

uSwitch刚开始的时候，每个团队都对自己的服务负责，并且尽可能的自治。团队负责创建自己的 AWS 账户、安全组、网络等。uSwitch 注意到对这些服务的调用数量正在增加，这与团队在为业务提供新功能和价值方面变得越来越慢的感觉相关。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/24image004-1615999812514.jpg)

**Figure 4: Use of low-level AWS services over time.**

**图 4：随着时间的推移使用低级别 AWS 服务。**

What uSwitch wanted to do by adopting Kubernetes was not only to bring in the technology but also to change the organizational structure and introduce an infrastructure platform team. This would address the increasing cognitive load generated by teams having to understand these different AWS services at a low level.

uSwitch 想要通过采用 Kubernetes 来做的不仅仅是引入技术，而是改变组织结构，引入一个基础设施平台团队。这将解决团队必须在低级别了解这些不同的 AWS 服务而产生的不断增加的认知负担。

That was a powerful idea. Once uSwitch introduced the platform, traffic directly through AWS decreased. The curves of calls in Figure 5 is a proxy for the cognitive load on the application teams. This concept aligns with a platform team’s purpose: to enable stream-aligned teams to work more autonomously with self-service capabilities and reduced extraneous cognitive load. This is a very different conceptual starting point from saying, "Well, we're going to put all shared services in a platform."

这是一个强大的想法。一旦 uSwitch 引入了该平台，直接通过 AWS 的流量就会减少。图 5 中的调用曲线代表了应用程序团队的认知负载。这个概念与平台团队的目的一致：使流对齐的团队能够通过自助服务功能更自主地工作并减少无关的认知负担。这是一个与说“好吧，我们将把所有共享服务放在一个平台中”的非常不同的概念起点。

Ingles also wrote that  they wanted to keep the principles they had in place before around team autonomy and teams working with minimal coordination, by providing a self-service infrastructure platform.

Ingles 还写道，他们希望通过提供自助服务基础设施平台来保持他们之前围绕团队自治和团队协作最少的原则。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/39image005-1615999814289.jpg)

**Figure 5: The number of low-level AWS service calls dropped after uSwitch introduced their Kubernetes-based platform.**

**图 5：在 uSwitch 推出其基于 Kubernetes 的平台后，低级别 AWS 服务调用的数量下降。**

## Treat the platform as a product

##将平台视为产品

We talk about treating the platform as a product — an internal product, but still a product. That means that we should think about its reliability, its fit for purpose, and the developer experience (DevEx) while using it.

我们谈论将平台视为产品——一个内部产品，但仍然是一个产品。这意味着我们应该在使用它时考虑它的可靠性、它的适用性以及开发人员体验 (DevEx)。

First, for the platform to be reliable, we need to have on-call support, because now the platform sits in the path of production. If our platform’s monitoring services fail or we run out of storage space for logs, for example, the customer-facing teams are going to suffer. They need someone who provides support, tells them what's going on, and estimates the expected time for a fix. It should also be easy to understand the status of the platform services and we should have clear, established communication channels between the platform and the stream teams in order to reduce cognitive load in communications. Finally, there needs to be coordination with potentially affected teams for any planned downtime. 
首先，要使平台可靠，我们需要有随叫随到的支持，因为现在平台处于生产路径中。例如，如果我们平台的监控服务出现故障或我们的日志存储空间不足，那么面向客户的团队就会受到影响。他们需要有人提供支持，告诉他们发生了什么，并估计修复的预期时间。平台服务的状态也应该很容易理解，我们应该在平台和流团队之间建立清晰的、已建立的沟通渠道，以减少沟通中的认知负担。最后，对于任何计划的停机时间，都需要与可能受影响的团队进行协调。
Secondly, having a platform that's fit for purpose means that we use techniques like prototyping and we get regular feedback from the internal customers. We use iterative practices like agile, pair programming, or TDD for faster delivery with higher quality. Importantly, we should focus on having fewer services of higher quality and availability rather than trying to build every service that we can imagine to be potentially useful. We need to focus on what teams really need and make sure those services are of high quality. This means we need very good product management to understand priorities, to establish a clear but flexible roadmap, and so on.

其次，拥有适合目的的平台意味着我们使用原型设计等技术，并定期从内部客户那里获得反馈。我们使用诸如敏捷、结对编程或 TDD 之类的迭代实践来以更高的质量更快地交付。重要的是，我们应该专注于拥有更少的更高质量和更高可用性的服务，而不是尝试构建我们认为可能有用的每一个服务。我们需要专注于团队真正需要的东西，并确保这些服务是高质量的。这意味着我们需要非常好的产品管理来了解优先级，建立清晰但灵活的路线图等等。

Finally, development teams and the platform should speak the same language in order to maximize the experience and usability. The platform should provide the services in a straightforward way. Sometimes, we might need to compromise. If development teams are not familiar with YAML, we might think of the low cost, low effort, and long-term gain of training all development teams in YAML. But these decisions are not always straightforward and we should never make them without considering the impact on the development teams or the consuming teams. We should provide the right levels of abstraction for our teams today, but the context may change in the future. We might adopt better or higher levels of abstraction, but we always should look at what makes sense given the current maturity and engineering practices of our teams.

最后，开发团队和平台应该使用同一种语言，以最大限度地提高体验和可用性。平台应以直接的方式提供服务。有时，我们可能需要妥协。如果开发团队不熟悉 YAML，我们可能会想到用 YAML 培训所有开发团队的低成本、低工作量和长期收益。但是这些决定并不总是直截了当的，我们不应该在不考虑对开发团队或消费团队的影响的情况下做出这些决定。我们今天应该为我们的团队提供正确的抽象级别，但上下文可能会在未来发生变化。我们可能会采用更好或更高级别的抽象，但鉴于我们团队当前的成熟度和工程实践，我们始终应该考虑什么是有意义的。

Kubernetes helped uSwitch establish these more application-focused abstractions for things like services, deployments, and ingress rather than the lower-level service abstractions that they were using before with AWS. It also helped them minimize coordination, which was another of the key principles.

Kubernetes 帮助 uSwitch 为服务、部署和入口等事物建立了这些更以应用程序为中心的抽象，而不是他们之前在 AWS 上使用的较低级别的服务抽象。它还帮助他们最大限度地减少协调，这是另一个关键原则。

I spoke with Ingles and with [Tom Booth](https://www.linkedin.com/in/thbooth/), at the time infrastructure lead at uSwitch, about what they did and how they did it.

我与 Ingles 和 [Tom Booth](https://www.linkedin.com/in/thbooth/) 进行了交谈，当时 uSwitch 的基础设施负责人，关于他们做了什么以及他们是如何做的。

Some of the things that the platform team helped the service teams with were providing dynamic database credentials and multi-cluster load balancing. They also made it easier for service teams to get alerts for their customer-facing services, define service-level objectives (SLOs), and make all that more visible. The platform make it easy for teams to configure and monitor their SLOs with dashboards — if an indicator drops below a threshold, notifications go out in Slack, as shown in Figure 6.

平台团队帮助服务团队完成的一些事情是提供动态数据库凭据和多集群负载平衡。它们还使服务团队能够更轻松地获得面向客户的服务的警报、定义服务级别目标 (SLO)，并使所有这些更加可见。该平台使团队可以使用仪表板轻松配置和监控他们的 SLO — 如果指标低于阈值，则通知会在 Slack 中发出，如图 6 所示。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/12image006-1615999813749.jpg)

**Figure 6: An example of a SLO threshold notification in Slack at uSwitch.**

**图 6：uSwitch 上 Slack 中的 SLO 阈值通知示例。**

Teams found it easy to adopt these new practices. uSwitch teams are familiar with YAML and can quickly configure these services and benefit from them.

团队发现采用这些新实践很容易。 uSwitch 团队熟悉 YAML，可以快速配置这些服务并从中受益。

## Achievements beyond the technical aspects

##超越技术方面的成就

uSwitch’s journey is fascinating beyond the technical achievements. They started this infrastructure adoption in 2018 with only a few services. They identified their first customer to be one team that was struggling without any centralized logging, metrics, or autoscaling. They recognized that growing services around these aspects would be a successful beginning.

uSwitch 的旅程令人着迷，超越了技术成就。他们于 2018 年开始采用这种基础设施，仅提供少量服务。他们确定他们的第一个客户是一个在没有任何集中日志记录、指标或自动缩放的情况下苦苦挣扎的团队。他们认识到围绕这些方面发展服务将是一个成功的开始。

In time, the platform team started to define their own SLAs and SLOs for the Kubernetes platform, serving as an example for the rest of the company and highlighting the improvements in performance, latency, and reliability. Other teams could observe and make an informed decision to adopt the platform services or not. Remember, it was never mandated but always optional.

随着时间的推移，平台团队开始为 Kubernetes 平台定义自己的 SLA 和 SLO，作为公司其他部门的榜样，并突出了性能、延迟和可靠性方面的改进。其他团队可以观察并做出明智的决定是否采用平台服务。请记住，它从未被强制要求，而始终是可选的。

Booth told me that uSwitch saw traffic increasing through the Kubernetes platform versus what was going directly through AWS and this gave them some idea of how much adoption was taking place. Later, the team addressed some critical cross-functional gaps in security, GDPR data privacy and handling of alerts and SLOs. 
Booth 告诉我，uSwitch 看到通过 Kubernetes 平台的流量与直接通过 AWS 的流量相比有所增加，这让他们对采用的情况有了一些了解。后来，该团队解决了安全性、GDPR 数据隐私以及警报和 SLO 处理方面的一些关键跨职能差距。
One team, one of the more advanced in both engineering terms and revenue generation, was already doing everything that the Kubernetes platform could provide. They had no significant motivation to adopt the platform — until they were sure that it provided the same functionality with the increased levels of reliability, performance, and so on. It no longer made sense for them to take care of all these infrastructure aspects on their own. The team switched to use the Kubernetes platform and increased its capacity to focus on the business aspects of the service. That was the “ultimate” prize for the platform team, to gain the adoption from the most advanced engineering team in the organization.

一个团队，在工程术语和创收方面更先进的团队之一，已经在做 Kubernetes 平台可以提供的一切。他们没有采用该平台的重大动机——直到他们确定它提供了相同的功能以及更高级别的可靠性、性能等。他们自己处理所有这些基础设施方面不再有意义。该团队转而使用 Kubernetes 平台，并增加了专注于服务业务方面的能力。这是平台团队的“终极”奖，获得组织中最先进的工程团队的采用。

## Four Key Metrics

## 四个关键指标

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/28image007-1615999814841.jpg)

**Figure 7: Four key metrics from the Accelerate book.**

**图 7：Accelerate 书中的四个关键指标。**

Having metrics can be quite useful. As the platform should be considered a product, we can look at product metrics — and the categories in Figure 7 come from the book Accelerate by Nicole Forsgren, Jez Humble, and Gene Kim. They write that high-performing teams do very well at these key metrics around lead time, deployment frequency, MTTR, and change failure rate. We can use this to help guide our own platform services delivery and operations.

拥有指标可能非常有用。由于平台应被视为产品，因此我们可以查看产品指标——图 7 中的类别来自 Nicole Forsgren、Jez Humble 和 Gene Kim 合着的《Accelerate》一书。他们写道，高绩效团队在交付周期、部署频率、MTTR 和变更失败率等关键指标方面做得非常好。我们可以使用它来帮助指导我们自己的平台服务交付和运营。

Besides the Accelerate metrics, user satisfaction is another useful and important measurement. If we're creating a product for users, we want to make sure it helps them do their job, that they're happy with it, and that they recommend it to others. There's a simple example from Twilio. Every quarter or so, their platform team surveys the engineering teams with some questions (Figure 8) on how well the platform helps them build, deliver, and run their service and how compelling it is to use.

除了 Accelerate 指标，用户满意度是另一个有用且重要的衡量指标。如果我们正在为用户创建产品，我们希望确保它可以帮助他们完成他们的工作，他们对此感到满意，并且他们将其推荐给其他人。 Twilio 有一个简单的例子。每个季度左右，他们的平台团队都会对工程团队进行调查，并提出一些问题（图 8），了解平台在多大程度上帮助他们构建、交付和运行他们的服务以及它的使用吸引力如何。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/11image008-1615999813204.jpg)

**Figure 8: The survey that Twiliio’s platform team sends out to engineering teams.**

**图 8：Twiliio 平台团队发给工程团队的调查。**

With this simple questionnaire, you can look at overall user satisfaction over time and see trends. It’s not just about the general level of satisfaction with the technical services, but also with the support from the platform team. Dissatisfaction may arise because the platform team was too busy to listen to feedback for a period of time, for example. It's not just about the technology. It's also about interactions among teams.

通过这个简单的调查问卷，您可以查看一段时间内的整体用户满意度并查看趋势。这不仅仅是对技术服务的总体满意度，还包括平台团队的支持。不满可能是因为平台团队太忙了一段时间没有听取反馈。这不仅仅是关于技术。它还与团队之间的互动有关。

Yet another important area to measure is around platform adoption and engagement. In the end, for the platform to be successful, it must be adopted. That means it’s serving its purpose. At the most basic level, we can look at how many teams in the organization are using the platform versus how many teams are not. We can also look at adoption per platform service, or even adoption of a particular service functionality. That will help understand the success of a service or feature. If we have a service that we expected to be easily adopted but many teams are lagging behind, we can look for what may have caused that.

另一个需要衡量的重要领域是平台采用和参与。最后，平台要想成功，就必须被采用。这意味着它正在为它的目的服务。在最基本的层面上，我们可以查看组织中有多少团队在使用该平台，而有多少团队没有使用。我们还可以查看每个平台服务的采用情况，甚至是特定服务功能的采用情况。这将有助于了解服务或功能的成功。如果我们有一项预计很容易被采用但许多团队落后的服务，我们可以寻找可能导致这种情况的原因。

Finally, measuring the reliability of the platform itself, as uSwitch did, is important as well. They had their own SLOs for the platform and this was available to all teams. Making sure we provide that information is quite important.

最后，如 uSwitch 所做的那样，衡量平台本身的可靠性也很重要。他们有自己的平台 SLO，所有团队都可以使用。确保我们提供这些信息非常重要。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/21image009-1615999815426.jpg)

**Figure 9: Useful metrics for your platform as a product.**

**图 9：作为产品的平台的有用指标。**

Figure 9 shows some examples of the metrics categories. Each organization with its own context might have different specific metrics, but the types and categories that we should be looking at should be more or less the same.

图 9 显示了度量类别的一些示例。每个具有自己背景的组织可能有不同的特定指标，但我们应该查看的类型和类别应该或多或少相同。

## Team interactions

##团队互动

The success of the platform team is the success of the stream-aligned teams. These two things go together. It's the same for other types of supporting teams. Team interactions are critical because the definition of success is no longer just about making technology available, it’s about helping consumers of that technology get the expected benefits in terms of speed, quality, and operability.

平台团队的成功就是流对齐团队的成功。这两件事是相辅相成的。其他类型的支持团队也是如此。团队互动至关重要，因为成功的定义不再只是让技术可用，而是要帮助该技术的消费者在速度、质量和可操作性方面获得预期的好处。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/8image010-1615999815955.jpg)

**Figure 10: Airbnb’s platform team provides an abstraction of the underlying Kubernetes architecture.** 
**图 10：Airbnb 的平台团队提供了底层 Kubernetes 架构的抽象。**
Airbnb effectively had a platform team (although internally called Infrastructure team) to abstract a lot of the underlying details of the Kubernetes platform, as shown in Figure 10. This clarifies the platform's boundaries for the development teams and exposes those teams to a much smaller cognitive load than telling them to use Kubernetes and to read the official documentation to understand how it works. That's a huge task that requires a lot of effort. This platform-team approach reduces cognitive load by providing more tightly focused services that meet our teams’ needs.

Airbnb 实际上有一个平台团队（虽然内部称为基础设施团队）来抽象 Kubernetes 平台的许多底层细节，如图 10 所示。这为开发团队澄清了平台的边界，并将这些团队暴露给一个更小的认知加载而不是告诉他们使用 Kubernetes 并阅读官方文档以了解其工作原理。这是一项艰巨的任务，需要付出很多努力。这种平台-团队方法通过提供更加专注的服务来满足我们团队的需求，从而减少了认知负担。

To accomplish this, we need adequate behaviors and interactions between teams. When we start a new platform service or change an existing one, then we expect strong collaboration between the platform team and the first stream-aligned teams that will use the new/changed service. Right from the beginning of the discovery period, a platform team should understand a streaming team’s needs, looking for the simplest solutions and interfaces that meet those needs. Once the stream-aligned team starta using the service, then the platform team’s focus should move on to supporting the service and providing up-to-date, easy to follow documentation for onboarding new users. The teams no longer have to collaborate as much and the platform team is focused on providing a good service. We call this interaction mode X-as-a-Service.

为了实现这一点，我们需要团队之间的适当行为和互动。当我们启动新的平台服务或更改现有的服务时，我们希望平台团队与将使用新/更改后的服务的第一个流对齐团队之间进行强有力的协作。从发现期开始，平台团队就应该了解流媒体团队的需求，寻找满足这些需求的最简单的解决方案和接口。一旦流对齐团队开始使用该服务，那么平台团队的重点应该转移到支持该服务并提供最新的、易于遵循的文档以供新用户使用。团队不再需要那么多协作，平台团队专注于提供良好的服务。我们称这种交互模式为 X-as-a-Service。

Note that this doesn't mean that the platform hides everything and the development teams are not allowed to understand what's going on behind the scenes. That's not the point. Everyone knows that it's a Kubernetes-based platform and we should not forbid teams from offering feedback or suggesting new tools or methods. We should actually promote that kind of engagement and discussion between stream-aligned teams and platform teams.

请注意，这并不意味着平台隐藏了一切，并且不允许开发团队了解幕后发生的事情。这不是重点。每个人都知道它是一个基于 Kubernetes 的平台，我们不应该禁止团队提供反馈或建议新的工具或方法。我们实际上应该促进流对齐团队和平台团队之间的这种参与和讨论。

For example, troubleshooting services in Kubernetes can be quite complicated. Figure 11 shows only the top half of a flow chart for diagnosing a deployment issue in Kubernetes. This is not something we want our engineering teams to have to go through every time there's a problem. Neither did Airbnb.

例如，Kubernetes 中的故障排除服务可能非常复杂。图 11 仅显示了诊断 Kubernetes 中部署问题的流程图的上半部分。这不是我们希望我们的工程团队每次出现问题时都必须经历的事情。 Airbnb也没有。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/18image011-1615999813488.jpg)

**Figure 11: Half of a Kubernetes deployment troubleshooting flow chart.**

**图 11：Kubernetes 部署故障排除流程图的一半。**

Airbnb initiated a discovery period during which teams closely collaborated to understand what kind of information they needed in order to diagnose a problem and what kinds of problems regularly occurred. This led to an agreement on what a troubleshooting service should look like. Eventually, the service became clear and stable enough to be consumed by all the stream-aligned teams.

Airbnb 启动了一个发现期，在此期间，团队密切合作以了解他们需要什么样的信息来诊断问题以及哪些类型的问题经常发生。这导致就故障排除服务的外观达成一致。最终，该服务变得足够清晰和稳定，可以被所有流对齐的团队使用。

Airbnb already provided these two services in their platform: kube-gen and kube-deploy. The new troubleshooting service, kube-diagnose, collected all the relevant logs, as well as all sorts of status checks and other useful data to simplify the lives of their development teams. Diagnosing got a lot easier, with teams focused on potential causes for problems rather than remembering where all the data was or which steps to get them.

Airbnb 已经在他们的平台中提供了这两项服务：kube-gen 和 kube-deploy。新的故障排除服务 kube-diagnose 收集了所有相关日志，以及各种状态检查和其他有用数据，以简化开发团队的生活。诊断变得更加容易，团队专注于问题的潜在原因，而不是记住所有数据的位置或获取它们的步骤。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/7image012-1615999814550.jpg)

**Figure 12: Services provided by the infrastructure (platform) team at Airbnb.**

**图 12：Airbnb 基础设施（平台）团队提供的服务。**

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/17image013-1615999815692.jpg)

**Figure 13: The cloud-native landscape.**

**图 13：云原生景观。**

Figure 13 shows just how broad the cloud-native landscape is. We don’t want stream-aligned teams to have to deal with that on their own. Part of the role of a platform team is to follow the technology lifecycle. We know how important that is. Let’s imagine there's a [CNCF](https://www.cncf.io/) tool that just graduated which would help reduce the amount of custom code (or get rid of an older, less reliable tool) in our internal platform today. If we can adopt this new tool in a transparent fashion that doesn’t leak into the platform service usage by stream-aligned teams, then we have an easier path forward in terms of keeping up with this ever-evolving landscape.

图 13 显示了云原生环境的广泛性。我们不希望流对齐的团队必须自己处理。平台团队的部分职责是跟踪技术生命周期。我们知道这有多重要。假设有一个 [CNCF](https://www.cncf.io/) 工具刚刚毕业，这将有助于减少我们今天内部平台中自定义代码的数量（或摆脱旧的、不太可靠的工具）。如果我们能够以透明的方式采用这种新工具，而不会泄漏到流对齐团队的平台服务使用中，那么我们就可以更轻松地跟上这个不断发展的格局。

If, on the contrary, a new technology we'd like to adopt in the platform implies a change in one or more service interfaces that means we need to consult the stream-aligned teams to understand the effort required of them and evaluate the trade- offs at play. Are we getting more benefit in the long run than the pain of migration/adaptation today?

相反，如果我们希望在平台中采用的新技术意味着一个或多个服务接口的更改，这意味着我们需要咨询流对齐团队以了解他们所需的工作并评估交易 -休赛期。从长远来看，我们是否获得了比今天迁移/适应的痛苦更多的好处？

In any case, the internal platform approach helps us make visible the evolution of the technology inside our platform.

无论如何，内部平台方法帮助我们使平台内部技术的演变可见。

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/7image014-1615999816275.jpg)

**Figure 14: uSwitch open-sourced Heimdall tool that their platform team created for internal use initially.**

**图 14：uSwitch 开源 Heimdall 工具，该工具是他们的平台团队最初为内部使用而创建的。**

The same goes for adopting open source solutions. For example, uSwitch open sourced the Heimdall application that provides the chat tool integration with the SLOs dashboard service. And there are many more open-source tools. Zalando, for example, has really cool stuff around cluster lifecycle management. The point being that if a piece of open-source technology makes sense for us, we can adopt it more easily if they sit under a certain level of abstraction in the platform. And if it is not a transparent change, we can always collaborate with stream-aligned teams to identify what would need to change in terms of their usage of the affected service(s).

采用开源解决方案也是如此。例如，uSwitch 开源了 Heimdall 应用程序，该应用程序提供聊天工具与 SLO 仪表板服务的集成。还有更多的开源工具。例如，Zalando 在集群生命周期管理方面有非常酷的东西。关键是，如果一项开源技术对我们有意义，如果它们位于平台中的某个抽象级别下，我们就可以更容易地采用它。如果这不是一个透明的变化，我们总是可以与流对齐的团队合作，以确定在他们对受影响服务的使用方面需要改变什么。

## Starting with team-centric approach for Kubernetes adoption

## 从以团队为中心的 Kubernetes 采用方法开始

There are three keys to a team-focused approach to Kubernetes adoption.

以团队为中心的 Kubernetes 采用方法有三个关键。

We start by assessing the cognitive load on our development teams or stream-aligned teams. Let’s determine how easy it is for them to understand the current platform based on Kubernetes. What abstractions do they need to know about? Is this easy for them or are they struggling? It's not an exact science, but by asking these questions we can get a feel for what problems they're facing and need help with. The Airbnb case study is important because it made it clear that a tool-based approach to Kubernetes adoption brings with it difficulties and anxiety if people have to use this all-new platform without proper support, DevEx, or collaboration to understand their real needs.

我们首先评估我们的开发团队或流对齐团队的认知负荷。让我们确定他们了解当前基于 Kubernetes 的平台有多容易。他们需要了解哪些抽象？这对他们来说容易还是他们在挣扎？这不是一门精确的科学，但通过提出这些问题，我们可以了解他们面临哪些问题并需要帮助。 Airbnb 案例研究很重要，因为它清楚地表明，如果人们不得不在没有适当支持、DevEx 或协作的情况下使用这个全新的平台来了解他们的真正需求，那么基于工具的 Kubernetes 采用方法会带来困难和焦虑。

Next, we should clarify our platform. This often is simple, but we don't always do it. List exactly all services we have in the platform, who is responsible for each, and all the other aspects of a digital platform that I’ve mentioned like responsibility for on-call support, communication mechanisms, etc. All this should be clear. We can start immediately by looking at the gaps between our actual Kubernetes implementation and an ideal digital platform, and addressing those.

接下来，我们应该澄清我们的平台。这通常很简单，但我们并不总是这样做。准确列出我们在平台中拥有的所有服务，谁负责每个服务，以及我提到的数字平台的所有其他方面，例如负责随叫随到的支持、通信机制等。 所有这些都应该很清楚。我们可以通过查看实际 Kubernetes 实施与理想数字平台之间的差距并解决这些差距来立即开始。

Finally, clarify the team interactions. Be more intentional about when we should collaborate and when should we expect to consume a service independently. Determine how we develop new services and who needs to be involved and for how long. We shouldn't  just say that we're going to collaborate and leave it as an open-ended interaction. Establish an expected duration for the collaboration, for example two weeks to understand what teams need from a new platform service and how the service interface should look like, before we’re actually building out such service. Do the necessary discovery first, then focus on functionality and reliability. At some point, the service can become “generally available” and interactions evolve to “X as a service”.

最后，明确团队互动。更有意识地了解何时应该协作以及何时应该独立使用服务。确定我们如何开发新服务以及需要参与的人员和参与时间。我们不应该只是说我们要合作，而将其作为开放式互动。在我们实际构建此类服务之前，确定协作的预期持续时间，例如两周，以了解团队需要从新平台服务中获得什么以及服务界面应该是什么样子。首先进行必要的发现，然后关注功能和可靠性。在某个时候，服务可以变得“普遍可用”，交互演变为“X 即服务”。

There are many good platform examples to look at from companies like  [Zalando](https://www.slideshare.net/try_except_/kubernetes-at-zalando-cncf-end-user-committee-presentation), [Twilio](https ://www.infoq.com/presentations/twilio-devops/), [Adidas](https://youtu.be/XwaRKcjkAAo), [Mercedes](https://speakerdeck.com/devopslx/2019-dot- 02-meetup-talk-devops-adoption-at-mercedes-benz-dot-io), etc. The common thread among them is a digital platform approach consisting not only of technical services but good support, on-call, high quality documentation , and all these things that make the platform easy for their teams to use and accelerate their capacity to deliver and operate their software more autonomously. I also wrote [an article for TechBeacon](https://techbeacon.com/enterprise-it/why-teams-fail-kubernetes-what-do-about-it) that goes a bit deeper into these ideas.

[Zalando](https://www.slideshare.net/try_except_/kubernetes-at-zalando-cncf-end-user-committee-presentation)、[Twilio](https ://www.infoq.com/presentations/twilio-devops/)、[阿迪达斯](https://youtu.be/XwaRKcjkAAo)、[梅赛德斯](https://speakerdeck.com/devopslx/2019-dot- 02-meetup-talk-devops-adoption-at-mercedes-benz-dot-io) 等。它们之间的共同点是数字平台方法，不仅包括技术服务，还包括良好的支持、随叫随到、高质量的文档，以及所有这些使他们的团队更容易使用该平台并提高他们更自主地交付和操作软件的能力的东西。我还写了 [一篇关于 TechBeacon 的文章](https://techbeacon.com/enterprise-it/why-teams-fail-kubernetes-what-do-about-it) 更深入地探讨了这些想法。

## About the Author

＃＃ 关于作者

**![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/4manuel-pais-1615999143952.jpg)Manuel Pais** is an independent IT organizational consultant and trainer, focused on team interactions, delivery practices, and accelerating flow. He is co-author of the book [Team Topologies: Organizing Business and Technology Teams for Fast Flow](https://teamtopologies.com/book). He helps organizations rethink their approach to software delivery, operations, and support via strategic assessments, practical workshops, and coaching. 
，专注于团队互动、交付实践和加速流程。他是 [团队拓扑：组织业务和技术团队以实现快速流程](https://teamtopologies.com/book) 一书的合著者。他通过战略评估、实践研讨会和辅导帮助组织重新思考他们的软件交付、运营和支持方法。
