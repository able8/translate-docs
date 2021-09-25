# Cloud-Native Is about Culture, Not Containers

# 云原生是关于文化，而不是容器

LikePrint [Bookmarks](http://www.infoq.com/showbookmarks.action)

LikePrint [书签](http://www.infoq.com/showbookmarks.action)

Mar 17, 2021

2021 年 3 月 17 日

### Key Takeaways

### 关键要点

- It is possible to be very cloud-native without a single microservice
- Before embarking on a cloud-native transformation, it’s important to be clear on what cloud-native means to your team and what the true problem being solved is
- The benefits of a microservices architecture will not be realised if releases involve heavy ceremony, are infrequent, and if all microservices have to be released at the same time
- Continuous integration and deployment is something you do, not a tool you buy
- Excessive governance chokes the speed out of cloud, but if you don’t pay enough attention to what is being consumed, there can be serious waste

- 没有单个微服务也可以成为非常云原生的
- 在开始云原生转型之前，重要的是要清楚云原生对您的团队意味着什么以及要解决的真正问题是什么
- 如果发布涉及繁重的仪式，不频繁，并且必须同时发布所有微服务，则微服务架构的好处将无法实现
- 持续集成和部署是您要做的事情，而不是您购买的工具
- 过度的治理会扼杀出云的速度，但如果你对正在消耗的东西不够关注，就会产生严重的浪费

At the QCon London last year, I provided a Cloud-Native session about Culture, not Containers. What had started me thinking about the role of culture in cloud native was a great [InfoQ article from Bilgin Ibryam](https://www.infoq.com/articles/microservices-post-kubernetes/). One of the things Bilgin did was define a cloud native architecture as lots of microservices, connected by smart pipes. I looked at that and thought it looked totally unlike applications I wrote, even though I thought I was writing cloud native applications. I'm part of the IBM Garage, helping clients get cloud native,, and yet I rarely used microservices in my apps. The apps I create mostly  looked nothing like Bilgin’s diagram. Does that mean I'm doing it wrong, or is maybe the definition of cloud native a bit complicated?

在去年的伦敦 QCon 大会上，我提供了一个关于文化而非容器的云原生会议。一篇很棒的 [InfoQ 文章来自 Bilgin Ibryam](https://www.infoq.com/articles/microservices-post-kubernetes/) 让我开始思考文化在云原生中的作用。 Bilgin 所做的一件事是将云原生架构定义为许多通过智能管道连接的微服务。我看着它，认为它看起来与我编写的应用程序完全不同，尽管我认为我正在编写云原生应用程序。我是 IBM Garage 的一员，帮助客户获得云原生，但我很少在我的应用程序中使用微服务。我创建的应用程序大多看起来与 Bilgin 的图表完全不同。这是否意味着我做错了，或者云原生的定义有点复杂？

I don’t want to single Bilgin out, since Bilgin’s article was called "Microservices in the post-Kubernetes Era," so it would be a bit ridiculous if he weren't talking about microservices a lot in that article. It’s also the case that almost all definitions of cloud native equate it to microservices. Everywhere I looked, I kept seeing the assumption that microservices equals native and Cloud-native equals microservices. Even the Cloud Native Computing Foundation used to define cloud native as all about microservices, and all about containers, with a bit of dynamic orchestration in there. Saying cloud native doesn’t always involve microservices, which puts me in this peculiar position because not only am I saying Bilgin is wrong, I'm saying the Cloud Native Computing Foundation is wrong - what did they ever know about Cloud-native? I'm sure I know way more than them, right?

我不想把 Bilgin 单独列出来，因为 Bilgin 的文章叫做“后 Kubernetes 时代的微服务”，所以如果他在那篇文章中不大量谈论微服务，那就有点可笑了。同样，几乎所有云原生的定义都将其等同于微服务。在我所看到的任何地方，我都看到了微服务等于原生和云原生等于微服务的假设。甚至云原生计算基金会过去也曾将云原生定义为关于微服务和容器的所有内容，其中包含一些动态编排。说云原生并不总是涉及微服务，这让我处于这个特殊的位置，因为我不仅说 Bilgin 错了，而且我说云原生计算基金会错了——他们对云原生了解多少？我相信我比他们知道的多，对吧？

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/10figure-1-1615824705148.jpg)

.jpg）

#### Related Sponsored Content

#### 相关赞助内容

#### Related Sponsor

#### 相关赞助商

**[![](https://assets.infoq.com/resources/en/ScrumRSBCoverScrum.png)](http://www.infoq.com/infoq/url.action?i=5e855cd4-7e93-4c47-844e-63e238096d43&t=f)**

-844e-63e238096d43&t=f)**

**Measure Value to Enable Improvement and Agility with Evidence-Based Management from Scrum.org - [Download the EBM Guide](http://www.infoq.com/infoq/url.action?i=4a892cb7-427b-4b9e-bb91-0a6d79444a4a&t=f).**

**衡量价值以通过 Scrum.org 的循证管理实现改进和敏捷性 - [下载 EBM 指南](http://www.infoq.com/infoq/url.action?i=4a892cb7-427b-4b9e-bb91-0a6d79444a4a&t=f).**

Well, obviously I don't. I'm on the wrong side of history on this one. I will admit that. (Although I'm on the wrong side of history, I notice that the CNCF have updated their definition of cloud native and, while microservices and containers are still there, they don't seem quite as mandatory as they used to be, so a little bit of history might be on my side!). Right or wrong, I'm still going to die on my little hill; that Cloud-native is about something much bigger than microservices. Microservices are one way of doing it. They're not the only way of doing it.

嗯，显然我没有。在这一点上，我站在历史的错误一边。我会承认这一点。 （虽然我站在历史的错误一边，但我注意到 CNCF 更新了他们对云原生的定义，虽然微服务和容器仍然存在，但它们似乎不像以前那样强制性了，所以一点点历史可能站在我这边！）。不管对错，我还是要死在我的小山上；云原生比微服务更重要。微服务是一种实现方式。它们不是唯一的方法。

In fact, you do see a range of definitions within our community. If you ask a bunch of people what Cloud-native means, some people will say “born on the cloud”. This was very much the original definition of Cloud-native back before microservices were even a thing. Some people will say it's microservices. 

事实上，您确实在我们的社区中看到了一系列定义。如果你问一堆人Cloud-native是什么意思，有人会说“生于云”。在微服务还没有出现之前，这在很大程度上是云原生的原始定义。有人会说是微服务。

Some people will say, “oh no, it's not just microservices, it's microservices on Kubernetes, and that's how you get Cloud-native”. This one I don’t like, because to me, Cloud-native shouldn't be about a technology choice. Sometimes I see Cloud-native used as a synonym for DevOps, because a lot of the cloud native principles and practices are similar to what devops teaches.

有些人会说，“哦，不，这不仅仅是微服务，它是 Kubernetes 上的微服务，这就是您获得云原生的方式”。我不喜欢这个，因为对我来说，云原生不应该是一种技术选择。有时我看到 Cloud-native 被用作 DevOps 的同义词，因为很多云原生的原则和实践与 DevOps 所教的很相似。

Sometimes I see Cloud-native used just as a way of saying “we’re developing modern software”: “We're going to use best practices; it's going to be observable; it's going to be robust; we’re going to release often and automate everything; in short, we're going to take everything we've learned over the last 20 years and develop software that way, and that's what makes it Cloud-native”. In this definition, cloud is just a given - of course it’s on cloud, because we’re developing this in 2021.

有时我看到云原生只是用来表达“我们正在开发现代软件”：“我们将使用最佳实践；这将是可观察的；它会很健壮；我们将经常发布并自动化一切；简而言之，我们将利用我们在过去 20 年中学到的一切，并以这种方式开发软件，这就是使其成为云原生的原因。”在这个定义中，云只是一个给定的——当然它是在云上，因为我们将在 2021 年开发它。

Sometimes I see Cloud-native used just to mean Cloud. We got so used to hearing Cloud-native that every time we talk about Cloud, we just feel like we have to tack a ‘-native’ on afterwards, but we're really just talking about Cloud. Finally, when people say Cloud-native, sometimes what they mean is idempotent. The problem with this is if you say Cloud-native means idempotent, everybody else goes, "What? What we really mean by “idempotent” is rerunnable? If I take it, shut it down, and then start it up again, there's no harm done. That's a fundamental requirement for services on the Cloud.

有时我看到 Cloud-native 只是用来表示 Cloud。我们已经习惯于听到 Cloud-native 以至于每次我们谈论 Cloud 时，我们只是觉得我们必须在之后加上一个“-native”，但我们实际上只是在谈论 Cloud。最后，当人们说云原生时，有时他们的意思是幂等的。问题在于，如果你说云原生意味着幂等，那么其他人都会说，“什么？我们所说的“幂等”的真正含义是可重新运行的吗？如果我接受它，关闭它，然后再次启动它，就没有已造成伤害。这是对云上服务的基本要求。

With all of these different definitions, is it any wonder we're not entirely sure what we're trying to do when we do Cloud-native?

有了所有这些不同的定义，难怪我们在做云原生时并不完全确定我们要做什么？

## Why?

##  为什么？

“What are we actually trying to achieve?” is an incredibly important question. When we're thinking about technology choices and technology styles, we want to be stepping back just from “I'm doing Cloud-native because that's what everybody else is doing” to thinking “what problem am I actually trying to solve?” To be fair to the CNCF, they had this “why” right on the front of their definition of Cloud-native. They said, "Cloud-native is about using microservices to build great products faster." We're not just using microservices because we want to; we're using microservices because they help us build great products faster.

“我们实际上想要达到什么目标？”是一个非常重要的问题。当我们考虑技术选择和技术风格时，我们希望从“我正在做云原生，因为这是其他人都在做的事情”到思考“我实际上想要解决什么问题？”为了对 CNCF 公平，他们将“为什么”放在了云原生定义的前面。他们说，“云原生是关于使用微服务更快地构建伟大的产品。”我们使用微服务不仅仅是因为我们想要；我们使用微服务是因为它们可以帮助我们更快地构建出色的产品。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/5figure-2-1615824705387.jpg)

.jpg）

This is where we step back to make sure we understand the problem we’re solving. Why couldn't we build great products faster before? It’s easy to skip this step, and I think all of us are guilty of this sometimes. Sometimes the problem that we're actually trying to solve is that everybody else is doing it, so we fear missing out unless we start doing it. Once we put it like that, FOMO isn’t a great decision criteria. Even worse, “my CV looks dull” definitely isn't the right reason to choose a technology.

这是我们退后一步以确保我们了解我们正在解决的问题的地方。为什么我们以前不能更快地构建伟大的产品？跳过这一步很容易，我认为我们所有人有时都会为此感到内疚。有时，我们实际上试图解决的问题是其他人都在做，所以我们害怕错过，除非我们开始做。一旦我们这样说，FOMO 就不是一个很好的决策标准。更糟糕的是，“我的简历看起来很乏味”绝对不是选择一项技术的正确理由。

## Why Cloud?

## 为什么是云？

I think to get to why we should be doing things in a Cloud-native way; we want to step back and say, "Why were we even doing things on the Cloud?" Here are the reasons:

我想了解为什么我们应该以云原生方式做事；我们想退后一步说，“为什么我们甚至在云端做事？”原因如下：

- **Cost**: Back when we first started putting things on the Cloud, price was the primary motivator. We said, "I've got this data center, I have to pay for the electricity, I have to pay people to maintain it. And I have to buy all the hardware. Why would I do that when I could use someone else's data center?" What creates a cost-saving between your own data center and someone else's data center is that your own data center has to stock up enough hardware for the maximum demand. That’s potentially a lot of capacity which is unused most of the time. If it's someone else's data center, you can pool resources. When demand is low, you won’t pay for the extra capacity.
- **Elasticity**: The reason Cloud saves you money is because of that elasticity. You can scale up; you can scale down. Of course, that's old news now. We all take elasticity for granted. 

- **成本**：当我们第一次开始将东西放在云端时，价格是主要的动力。我们说，“我有这个数据中心，我要付电费，我要付钱来维护它。我必须购买所有的硬件。当我可以使用别人的数据时，我为什么要这样做？中央？”在您自己的数据中心和其他人的数据中心之间节省成本的是，您自己的数据中心必须为最大需求储备足够的硬件。这可能是大部分时间未使用的大量容量。如果是别人的数据中心，你可以共享资源。当需求低时，您无需为额外的容量付费。
- **弹性**：云为您省钱的原因是因为这种弹性。你可以扩大规模；你可以缩小规模。当然，这已经是旧闻了。我们都认为弹性是理所当然的。

- **Speed**: The reason we're interested in Cloud now is because of the speed. Not necessarily the speed of the hardware, although some cloud hardware can be dazzlingly fast. The cloud is an excellent way to use GPUs, and it’s more or less the only way to use quantum computers. More generally, though, we can get something to market way, way faster via the cloud than we could when we had to print software onto CD-Roms and mail them out to people, or even when we had to stand instances up in our own data center.

- **速度**：我们现在对云感兴趣的原因是速度。不一定是硬件的速度，尽管一些云硬件可以快得令人眼花缭乱。云是使用 GPU 的绝佳方式，它或多或少是使用量子计算机的唯一方式。但是，更普遍的是，我们可以通过云以更快的方式将某些东西推向市场，这比我们不得不将软件打印到 CD-Roms 上并将它们邮寄给人们时，或者甚至当我们必须在我们自己的实例中站立时更快数据中心。

## 12 Factors

## 12 个因素

Cost savings, elasticity, and delivery speed are great, but we get all of that  just by being on the Cloud. Why do we need Cloud-native? The reason we need Cloud-native is that a lot of companies found they tried to go to the Cloud and they got electrocuted.

成本节约、弹性和交付速度都很棒，但我们只需在云端即可获得所有这些。为什么我们需要云原生？我们需要云原生的原因是，很多公司发现他们试图去云，结果触电了。

It turns out things need to be written differently and managed differently on the cloud. Articulating these differences led to the 12 factors. The 12 factors were a set of mandates for how you should write your Cloud application so that you didn't get electrocuted.

事实证明，事情需要在云上以不同方式编写和管理。阐明这些差异导致了 12 个因素。这 12 个因素是关于您应该如何编写云应用程序以免触电的一组指令。

You could say the 12 factors described how to write a cloud native application - but the 12 factors had absolutely nothing to do with microservices. They were all about how you managed your state. They were about how you managed your logs. The 12 factors helped applications become idempotent, but “the 12 factors” is catchier than “the idempotency factors”.

你可以说这 12 个因素描述了如何编写云原生应用程序 - 但 12 个因素与微服务完全无关。它们都是关于你如何管理你的状态。它们是关于您如何管理日志的。 12 个因素帮助应用程序变得幂等，但“12 个因素”比“幂等因素”更吸引人。

The 12 factors were published two years before Docker got to market. Docker containers revolutionised how the cloud was used. Containers are so good, it’s hard to overstate their importance. They solve many problems and create new architectural possibilities. Because containers are so it’s easy, it’s possible to distribute an application across many containers. Some companies are running single applications across 100, 200, 300, 400, or 500 distinct containers. Compared to that kind of engineering prowess, an application which is spread across a mere six containers seems a bit inadequate. In the face of so little complexity, it’s easy to think  “I must be doing it really wrong. I'm not as good a developer as them over there”.

这 12 个因素是在 Docker 上市前两年发布的。 Docker 容器彻底改变了云的使用方式。容器是如此之好，很难夸大它们的重要性。它们解决了许多问题并创造了新的架构可能性。由于容器非常简单，因此可以将应用程序分布在多个容器中。一些公司在 100、200、300、400 或 500 个不同的容器中运行单个应用程序。与那种工程实力相比，仅仅分布在六个容器中的应用程序似乎有点不够。面对这么小的复杂性，很容易想到“我一定做错了。我不像他们那边的开发人员那么好”。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/4figure-3-1615824705625.jpg)

.jpg）

Actually, no. It's not a competition to see how many containers you can have. Containers are great, but the number of containers you have should be tuned to your needs.

实际上，没有。看您可以拥有多少个容器不是比赛。容器很棒，但是您拥有的容器数量应该根据您的需要进行调整。

## Speed

##  速度

Let’s try and remember - what were your needs again? When we think about Cloud,  we usually want to be thinking about that speed. The reason we want lots of containers is that we want to get new things to market faster. If we have lots of containers and we're either shipping the exact same things to market or we're getting to market at the same speed, then all of a sudden, those containers are only a cost. They're not helping us, and we’re burning cycles managing the complexity that comes with scattering an application in tiny pieces all over the infrastructure. If we have this amazing architecture that allows us to respond to the market but we’re not responding, then that's a waste. If we have this architecture, that means we can go fast, but we're not going fast, then that's a waste as well.

让我们试着记住——你的需求又是什么？当我们想到 Cloud 时，我们通常想考虑这种速度。我们需要大量容器的原因是我们希望更快地将新事物推向市场。如果我们有很多集装箱，我们要么向市场运送完全相同的东西，要么以相同的速度进入市场，那么突然之间，这些集装箱只是成本。他们没有帮助我们，我们正在燃烧周期来管理将应用程序分散到整个基础架构中的小块所带来的复杂性。如果我们拥有这种令人惊叹的架构，让我们能够对市场做出反应，但我们却没有做出反应，那就是一种浪费。如果我们有这种架构，那意味着我们可以走得快，但我们走得不快，那也是一种浪费。

## How to fail at Cloud-Native 

## 如何在云原生上失败

Which brings me to how to fail at Cloud-native. For context, I'm a consultant. I'm a full-stack developer in the IBM Garage. We work with startups and with large companies, helping them get to the cloud and get the most out of cloud. As part of that, we help them solve interesting, tough, problems, and we help them do software faster than they’ve been able to to do it before. To make sure we’re really getting the most out of the cloud, we do a lean startup, extreme programming, design thinking,  DevOps; and cloud native. Because I’m a consultant, I see a lot of customers who are on the journey to the Cloud. Sometimes that goes well, and sometimes there are these pitfalls. Here  are some of the traps that I've seen smart clients fall into. So, What is Cloud-Native?

这让我想到了如何在云原生上失败。就上下文而言，我是一名顾问。我是 IBM Garage 的一名全栈开发人员。我们与初创公司和大公司合作，帮助他们进入云并充分利用云。作为其中的一部分，我们帮助他们解决有趣的、棘手的问题，我们帮助他们以比以前更快的速度开发软件。为了确保我们真正充分利用云，我们进行了精益创业、极限编程、设计思维、DevOps；和云原生。因为我是一名顾问，所以我看到很多客户正在走向云端。有时这很顺利，有时也有这些陷阱。以下是我见过的聪明客户陷入的一些陷阱。那么，什么是云原生？

One of the earliest traps is the magic morphing meaning. If I say Cloud-native and I mean one thing and you say Cloud-native and mean another thing, we’re going to have a problem communicating...

最早的陷阱之一是魔法变形的含义。如果我说云原生的意思是一回事，而你说云原生的意思是另一回事，那么我们的沟通就会出现问题......

Sometimes that doesn't really matter, but sometimes it makes a big difference. If one person thinks the goal is microservices and then the other person feels the goal is to have an idempotent system, uh oh. Or if part of an organisation wants to go to the Cloud because they think it's going to allow them to get to market faster, but another part is only going to the Cloud to deliver the exact same speed as before, but more cost-effectively, then we might have some conflict down the road.

有时这并不重要，但有时它会产生很大的不同。如果一个人认为目标是微服务，然后另一个人认为目标是拥有一个幂等系统，呃哦。或者，如果组织的一部分想要使用云，因为他们认为这将使他们能够更快地进入市场，但另一部分只是为了提供与以前完全相同的速度，但更具成本效益，那么我们可能会在路上发生一些冲突。

## Microservices Envy

## 微服务羡慕

Often one of the things that drives some of this confusion about goals is because we have a natural tendency to look at other people, doing fantastic things, and want to emulate them. We want to do those fantastic things ourselves without really thinking about our context and whether they're appropriate for us. One of our IBM Fellows has a heuristic when he goes in to talk to a client about microservices. He says, "If they start talking about Netflix and they just keep talking about Netflix, and they never mention coherence, and they never mention coupling, then probably they're not really doing it for the right reasons."

通常，造成这种对目标的混淆的原因之一是因为我们有一种自然的倾向，看别人，做奇妙的事情，并想效仿他们。我们想自己做这些奇妙的事情，而无需真正考虑我们的背景以及它们是否适合我们。我们的一位 IBM 研究员在与客户讨论微服务时有一种启发式方法。他说：“如果他们开始谈论 Netflix，然后一直在谈论 Netflix，却从不提及一致性，也从不提及耦合，那么他们这样做的理由可能并不正确。”

Sometimes we talk to clients, and they say, "Right, I want to modernize to microservices." Well, microservices are not a goal. No customer will look at your website and say, "Oh, microservices. That's nice." Customers are going to look at your website and judge it on whether it serves their needs, whether it’s easy and delightful, and, all of these other things. Microservices can be an excellent means to that end, but they're not a goal in themselves. I should also say: microservices are a means. They're not necessarily the only means to that goal.

有时我们与客户交谈，他们会说：“是的，我想对微服务进行现代化改造。”好吧，微服务不是目标。没有客户会看着你的网站说，“哦，微服务。这很好。”客户会查看您的网站，并根据它是否满足他们的需求、是否简单和令人愉快，以及所有其他方面来判断它。微服务可以成为实现这一目标的绝佳手段，但它们本身并不是目标。我还应该说：微服务是一种手段。它们不一定是实现该目标的唯一途径。

A colleague of mine in the IBM Garage had some conversations with a bank in Asia-Pacific. The bank was having problems responding to their customers, because their software was all old and heavy and calcified. They were also having a people problem because all of their COBOL developers were old and leaving the workforce. So the bank  knew they had to modernise. The main driver in this case wasn't the aging workforce, it was really competitiveness and agility. They were getting beaten by their competitors because they had this big estate of COBOL code and every change was expensive and slow. They said, "Well, to solve this problem, we need to get rid of all of our COBOL, and we need to switch to a modern microservices architecture." 

我在 IBM Garage 的一位同事与亚太地区的一家银行进行了一些对话。这家银行在响应客户方面遇到了问题，因为他们的软件都是旧的、笨重的、钙化的。他们还遇到了人员问题，因为他们所有的 COBOL 开发人员都很老并且都离开了劳动力队伍。所以银行知道他们必须现代化。在这种情况下，主要驱动因素不是劳动力老龄化，而是竞争力和敏捷性。他们被竞争对手打败了，因为他们拥有大量的 COBOL 代码，而且每次更改都既昂贵又缓慢。他们说，“嗯，要解决这个问题，我们需要摆脱所有的 COBOL，我们需要切换到现代微服务架构。”

So far, so good. We were just gearing up to jump in with some cloud native goodness, when the bank added that their release board only met twice a year. At this point, we wound back. It didn't matter how many microservices the bank’s shiny new architecture would have; those microservices were all going to be assembled up into a big monolith release package and deployed twice a year. That’s taking the overhead of microservices without the benefit. Since it’s not a competition to see how many containers you have, lots of containers and slow releases would be a stack in which absolutely no one won.

到现在为止还挺好。当银行补充说他们的发布委员会每年只召开两次会议时，我们正准备加入一些云原生优势。在这一点上，我们退缩了。银行闪亮的新架构将拥有多少微服务并不重要；这些微服务都将组装成一个大型单体发布包，并每年部署两次。这在没有好处的情况下占用了微服务的开销。既然看你有多少容器不是一场比赛，那么大量的容器和缓慢的发布将是一个绝对没有人赢的堆栈。

Not only would lots of microservices locked into a sluggish release cadence not be a win, it could be a bad loss. When organisations attempt microservices, they don’t always end up with a beautiful decoupled microservices architecture like the ones in the pictures. Instead, they  end up with a distributed monolith. This is like a normal monolith, but far worse. The reason that this is extra-scary bad is because a normal, non-distributed, monolith has things like compile-time checking for types and synchronous, guaranteed, internal communication. Running in a single process is going to hurt your scalability, but it means that you can't get bitten by the distributed computing fallacies. If you take that same application and then just smear it across the internet and don't put in any type checking or invest in error handling for network issues, you're not going to have a better customer experience; you're going to have a worse customer experience.

许多微服务被锁定在缓慢的发布节奏中不仅不会成功，而且可能是一个严重的损失。当组织尝试微服务时，他们并不总是像图片中那样以漂亮的解耦微服务架构告终。相反，他们最终得到了一个分布式单体。这就像一个普通的单体，但更糟糕。这是非常糟糕的原因是因为一个普通的、非分布式的、单体有诸如编译时检查类型和同步、有保证的内部通信之类的东西。在单个进程中运行会损害您的可扩展性，但这意味着您不会被分布式计算谬误所困扰。如果您使用相同的应用程序，然后只是在互联网上涂抹它，而不进行任何类型检查或投资于网络问题的错误处理，那么您将不会获得更好的客户体验；你会有更糟糕的客户体验。

There's a lot of contexts in which microservices are the wrong answer. If you're a small team, you don't need to have lots of autonomous teams because each independent team would be about a quarter of a person. Suppose you don't have any plans or any desire to release part of your application independently, then you won’t benefit from microservices’s independence.

在很多情况下，微服务都是错误的答案。如果你是一个小团队，你不需要有很多自治团队，因为每个独立团队大约有四分之一的人。假设您没有任何计划或希望独立发布部分应用程序，那么您将无法从微服务的独立性中受益。

In order to give security and reliable communication and discoverability between all of these components of your application that you've just smeared across a part of the Cloud, you're going to need something like a service mesh. You might be either quite advanced on the tech curve or a little bit new to that tech curve. You either don't know what a service mesh is, or you say, "I know all about what a service mesh is. So complicated, so overhyped. I don't need a service mesh. I'm just going to roll my own service mesh instead." This will not necessarily give you the outcome you hoped for. You will still end up with a  service mesh, but you have to maintain it, because you wrote it!

为了在应用程序的所有这些组件之间提供安全、可靠的通信和可发现性，这些组件刚刚在云的一部分中涂抹，您将需要诸如服务网格之类的东西。您可能在技术曲线上非常先进，或者对该技术曲线有点陌生。你要么不知道服务网格是什么，要么你说，“我知道服务网格是什么。太复杂了，太夸张了。我不需要服务网格。我只是要推出我的而是拥有自己的服务网格。”这不一定会给您带来希望的结果。你仍然会得到一个服务网格，但你必须维护它，因为它是你写的！

Another good reason not to do microservices is sometimes the domain model just doesn't have those natural fracture points that allow you to get nice neat microservices. In that case, it is totally reasonable to say, "You know what? I'm just going to leave it."

另一个不做微服务的好理由是，有时领域模型没有那些自然的断裂点，让你可以得到漂亮整洁的微服务。在这种情况下，完全有理由说，“你知道吗？我要离开它了。”

## Cloud-native spaghetti

## 云原生意大利面

If you don't step away from the blob, then you end up with the next problem, which is Cloud-native spaghetti. I always feel slightly panicked when I look at the communication diagram for the Netflix microservices. I'm sure they know what they're doing, and they've got it figured out, but to my eyes, it looks exactly like spaghetti. Making that work needs a lot of really solid engineering and specialised skills. If you don't have that specialisation, then you end up in a messy situation. 

如果您不离开 blob，那么您最终会遇到下一个问题，即云原生意大利面。看 Netflix 微服务的通信图时，总觉得有点慌。我敢肯定他们知道他们在做什么，而且他们已经弄清楚了，但在我看来，它看起来就像意大利面。完成这项工作需要很多非常扎实的工程和专业技能。如果你没有那个专业，那么你最终会陷入混乱的境地。

I was brought in to do some firefighting with a client who was struggling. They were developing a greenfield application, and so of course they’d chosen microservices, to be as modern as possible. One of the first things they said to me was "any time we change any code at all, something else breaks." This isn’t what’s supposed to happen with microservices. In fact, it’s the exact opposite of what we’ve all been told happens if we implement microservices. The dream of microservices is that they are decoupled. Sadly, decoupling doesn't come for free. It certainly doesn't magically happen just because you distributed things. All that happens when you distribute things is that you have two problems instead of one.

我被带去和一个苦苦挣扎的客户一起救火。他们正在开发一个全新的应用程序，因此他们当然选择了尽可能现代的微服务。他们对我说的第一件事就是“每当我们更改任何代码时，其他事情都会中断。”这不是微服务应该发生的事情。事实上，这与我们所知道的如果我们实施微服务会发生的情况完全相反。微服务的梦想就是解耦。可悲的是，解耦不是免费的。它当然不会因为你分发东西而神奇地发生。当你分发东西时发生的所有事情就是你有两个问题而不是一个。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/3figure-4-1615824703351.jpg)

.jpg）

**Cloud Native Spaghetti is still spaghetti.**

**云原生意大利面仍然是意大利面。**

One of the reasons my client’s code was so brittle and connected was that they had quite a complex object model, with around 20 classes and 70 fields in some of the classes. Handling that kind of complex object model in a microservices system is tough. In this case, they looked at their complex object model, and decided, "We know it's really bad to have common code between our microservices because then we're not decoupled. Instead, we're going to copy and paste this common object model across all of our six microservices. Because we cut and paste it rather than linking to it, we're decoupled." Well, no, you're not decoupled. If things break when one thing changes, whether the code is linked or copied, there’s coupling.

我客户的代码如此脆弱且相互关联的原因之一是他们有一个相当复杂的对象模型，其中一些类有大约 20 个类和 70 个字段。在微服务系统中处理这种复杂的对象模型是很困难的。在这种情况下，他们查看了他们的复杂对象模型，并决定，“我们知道在我们的微服务之间使用公共代码真的很糟糕，因为这样我们就不会解耦。相反，我们将复制和粘贴这个公共对象模型跨越我们所有的六个微服务。因为我们剪切和粘贴它而不是链接到它，所以我们解耦了。”嗯，不，你没有解耦。如果当一件事发生变化时事情就会中断，无论是链接还是复制代码，都存在耦合。

What was the ‘right’ thing to do in this case? In the ideal case, each microservice maps neatly to a domain, and they're quite distinct. If you have one big domain and lots of tiny microservices, then there's going to be a problem. The solution is to either decide the domain really is big and merge the microservices, or to do deeper domain modelling to try and untangle the object model into distinct bounded contexts.

在这种情况下，“正确”的做法是什么？在理想情况下，每个微服务都巧妙地映射到一个域，并且它们非常不同。如果您有一个大域和许多微小的微服务，那么就会出现问题。解决方案是确定域确实很大并合并微服务，或者进行更深入的域建模以尝试将对象模型解开为不同的有界上下文。

Even with the cleanest domain separation, in any system, there will always be some touch points between components - that’s what makes it a system. These touch points are easy to get wrong, even if they’re minimal, and especially if they’re hidden. Do you remember the Mars Climate Orbiter? Unlike the Perseverance, it was designed to orbit Mars from a safe distance, rather than land on it. Sadly, it strayed too close to Mars, got pulled in by Mars’s gravity, and crashed. The loss of the probe was sad, and the underlying reason was properly tragic. The Orbiter was controlled by two modules, one the probe, and one on earth. The probe module was semiautonomous, since the Orbiter was not visible from earth most of the time. About every three days the planets would align, it would come into view, and the team on earth would fine-tune its trajectory. I imagine the instructions were along the lines of "Oh, I think you need to shift a bit left and oh you're going to miss Mars if you don't go a bit right," except in numbers.

即使有最清晰的域分离，在任何系统中，组件之间总会有一些接触点——这就是它成为一个系统的原因。这些接触点很容易出错，即使它们很小，尤其是在它们被隐藏的情况下。你还记得火星气候轨道器吗？与毅力不同的是，它的设计目的是从安全距离绕火星运行，而不是降落在火星上。可悲的是，它离火星太近了，被火星的引力拉了进去，然后坠毁了。探测器的丢失令人悲伤，其根本原因是悲惨的。轨道器由两个模块控制，一个是探测器，一个在地球上。探测器模块是半自主的，因为大多数时候从地球上看不到轨道飞行器。大约每三天，行星就会对齐，它就会进入视野，地球上的团队就会微调它的轨迹。我想象这些指令是“哦，我认为你需要向左移动一点，哦，如果你不向右移动一点，你会错过火星”，除了数字。

The numbers were what led to the problem. The earth module and probe module were two different systems built by two different teams. The probe used imperial units, and the JPL ground team used metric. Even though the two systems seemed independent, there was a very significant point of coupling between them. Every time the ground team transmitted instructions, what they sent was interpreted in a way that no one expected.

数字是导致问题的原因。地球模块和探测器模块是由两个不同团队建造的两个不同系统。探测器使用英制单位，喷气推进实验室地面团队使用公制单位。尽管这两个系统看起来是独立的，但它们之间存在非常重要的耦合点。每次地面小队传送指令时，他们发出的指令都以一种出人意料的方式被解读。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/2figure-5-1615824704329.jpg)

.jpg）

**The moral of the story is that distributing the system did not help. Part of the system was on Mars, and part of the system was on Earth, and you can’t get more distributed than that.**

**这个故事的寓意是分发系统没有帮助。该系统的一部分在火星上，而该系统的一部分在地球上，您无法获得比这更多的分布。**

## Microservices need consumer-driven contact tests 

## 微服务需要消费者驱动的接触测试

In this case, the solution, the correct thing to do is be really clear about what the points of coupling are and what each side’s expectations are. A great way of doing this is consumer contract-driven tests. Contract tests aren’t yet widely used in our industry, despite being a clean solution to a big problem. I think part of the problem is that they can be a bit tricky to learn, which has slowed adoption. Cross-team negotiations about the tests can also be complicated - although if negotiation about a test is too hard, negotiation about the actual interaction parameters will be even harder. If you’re thinking of exploring contract testing, Spring Contract or Pact are good starting points. Which one is right for you depends on your context. Spring Contract is nicely integrated into the Spring ecosystem, whereas Pact is framework-agnostic and supports a huge range of languages, including Java and Javascript.

在这种情况下，解决方案，正确的做法是真正清楚耦合点是什么以及每一方的期望是什么。这样做的一个好方法是消费者合同驱动的测试。尽管合同测试是解决大问题的干净解决方案，但在我们的行业中尚未广泛使用。我认为部分问题在于它们学习起来有点棘手，这减缓了采用速度。关于测试的跨团队协商也可能很复杂——尽管如果关于测试的协商太难，关于实际交互参数的协商将更加困难。如果您正在考虑探索合约测试，Spring Contract 或 Pact 是不错的起点。哪一种适合您取决于您的上下文。 Spring Contract 很好地集成到 Spring 生态系统中，而 Pact 与框架无关，支持多种语言，包括 Java 和 Javascript。

Contract tests go well beyond what OpenAPI validation does, because it checks the semantics of the APIs, rather than just the syntax. It’s a much more helpful check than "well, the fields on each side have the same name, so we're good." It allows you to check, "is my behavior when I get these inputs the expected behavior? Are the assumptions I'm naming about that API over there still valid?" Those are things you need to check, because if they’re not true, things are going to get really bad.

契约测试远远超出 OpenAPI 验证的作用，因为它检查 API 的语义，而不仅仅是语法。这是一个比“好吧，每一侧的字段名称相同，所以我们很好”更有帮助的检查。它允许您检查，“当我获得这些输入时，我的行为是否符合预期行为？我在那里命名的关于该 API 的假设是否仍然有效？”这些是你需要检查的事情，因为如果它们不是真的，事情会变得非常糟糕。

Many companies are aware of this risk and are aware that there's an instability in the system when they're doing microservices. To have confidence that these things work together, they impose a UAT phase before releasing them. Before any microservice can be released, someone needs to spend several weeks testing it works properly in the broader system. With that kind of overhead, releasing isn’t going to be happening often. Then that leads us to the classic anti-pattern, which is not-actually-continuous continuous integration and continuous deployment, or I/D.

许多公司都意识到了这种风险，并意识到在进行微服务时系统存在不稳定性。为了确信这些东西可以协同工作，他们在发布之前强加了一个 UAT 阶段。在发布任何微服务之前，有人需要花费数周时间测试它在更广泛的系统中是否正常工作。有了这种开销，发布就不会经常发生了。那么这将我们引向经典的反模式，即实际上不是连续的持续集成和持续部署，或 I/D。

### Why Continuous Integration and Deployment isn’t

### 为什么持续集成和部署不是

I talk to a lot of customers, and they'll say, "We have a CI/CD." The ‘a’ sets off alarm bells, because CI/CD, should not be a tool you buy, put on a server, and admire, saying "there's CI/CD." CD/CD is something that you have to be doing. The letters stand for continuous integration and continuous deployment or delivery. Continuous in this context means “integrating really often” and “deploying really often,”, and if you're not doing that, then it's simply not continuous.

我与很多客户交谈，他们会说，“我们有 CI/CD。” ‘a’ 敲响了警钟，因为 CI/CD 不应该是你购买、放在服务器上并欣赏的工具，说“有 CI/CD”。 CD/CD 是您必须要做的事情。这些字母代表持续集成和持续部署或交付。在这种情况下，连续意味着“非常频繁地集成”和“非常频繁地部署”，如果你不这样做，那么它就不是连续的。

Sometimes I'll overhear comments  like “I'll merge my branch into our CI system next week”. This completely misses the point of the “C” in “CI”, which stands for continuous. If you merge once a week, that's not continuous. That's almost the opposite of continuous.

有时我会无意中听到诸如“下周我将我的分支合并到我们的 CI 系统中”之类的评论。这完全忽略了“CI”中“C”的意思，它代表连续。如果你每周合并一次，那不是连续的。这几乎与连续相反。

The “D” part can be even more of a struggle. If software is only deployed every six months, the CI/CD server may be useful, but no one is doing CD. There may be “D”, but everyone has forgotten the “C” part. 

“D”部分可能更像是一场斗争。如果软件每六个月才部署一次，CI/CD 服务器可能有用，但没有人在做 CD。可能有“D”，但大家都忘记了“C”部分。

How often is it actually reasonable to be pushing to main? How continuous is continuous have to be? Even I will admit that some strict definitions of continuous would be a ridiculous way to write software in a team. If you pushed to main every character, that is technically continuous, but it is going to cause mayhem in a team. If you integrate every commit and aim to commit several times an hour, that's probably a pretty good cadence. If you commit often and integrate every few commits, you're pushing several times a day, so that's also pretty good. If you're doing test-driven development, then integrating when you get a passing test is an excellent pattern. I'm a big advocate of trunk-based development. TBD has many benefits in terms of debugging, enabling opportunistic refactoring, and avoiding big surprises for colleagues. The technical definition of trunk-based development is that you need to be integrating at least once a day to count. I sometimes hear “once a day” described as the bar between “ok” and “just not continuous”. Once a week is getting really problematic.

推到 main 的频率实际上是多少？连续必须有多连续？甚至我也承认，在团队中编写软件的一些严格的连续定义将是一种荒谬的方式。如果你推动每个角色都成为主角，这在技术上是连续的，但这会在团队中造成混乱。如果您整合每个提交并打算每小时提交几次，这可能是一个非常好的节奏。如果你经常提交并且每隔几次提交就集成一次，那么你每天会推送几次，所以这也很好。如果您正在进行测试驱动的开发，那么当您通过测试时进行集成是一个很好的模式。我是基于主干开发的大力倡导者。 TBD 在调试、支持机会重构和避免给同事带来大的惊喜方面有很多好处。基于主干开发的技术定义是，您需要每天至少集成一次才能计算。我有时会听到“一天一次”被描述为“还可以”和“只是不连续”之间的界限。每周一次真的很成问题。

Once you get into one every month, it's terrible. When I joined IBM we used a build system and a code repository called CMVC. For context, this was about twenty years ago, and our whole industry was younger and more foolish. My first job in IBM was helping build the WebSphere Application Server. We had a big multi-site build, and the team met six days a week, including Saturdays,  to discuss any build failures. That call  had a lot of focus, and you did not want to be called up on the WebSphere build call. I’d just left university and knew nothing about software development in a team, so some of the senior developers took me under their wings. One piece of advice I still remember was that the way to avoid being on the WebSphere build call was to save up all of your changes on your local machine for six months and then push them all in a batch.

一旦你每个月进入一个，那就太可怕了。当我加入 IBM 时，我们使用了一个构建系统和一个名为 CMVC 的代码存储库。就上下文而言，这是大约二十年前，我们整个行业更年轻、更愚蠢。我在 IBM 的第一份工作是帮助构建 WebSphere Application Server。我们有一个大型的多站点构建，团队每周开会六天，包括周六，讨论任何构建失败。该调用有很多重点，您不希望在 WebSphere 构建调用中被调用。我刚离开大学，对团队中的软件开发一无所知，所以一些高级开发人员把我收在了他们的翅膀下。我仍然记得的一条建议是，避免参与 WebSphere 构建调用的方法是将所有更改保存在本地机器上六个月，然后将它们全部推送到批处理中。

At the item, I was little, and I thought, ok, that doesn't seem like quite the right advice, but I guess you know best. With hindsight, I realize the WebSphere build broke badly because people were saving their changes for six months before then trying to integrate with their colleagues. Obviously, that didn't work, and we changed how we did things.

在这个项目上，我还很小，我想，好吧，这似乎不是一个正确的建议，但我想你最了解。事后看来，我意识到 WebSphere 构建失败了，因为人们在尝试与同事集成之前将他们的更改保存了六个月。显然，那行不通，我们改变了做事的方式。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-6-1615824705942.jpg)

.jpg）

**How often should you integrate?**

**您应该多久集成一次？**

The next question, which is even harder, is how often you should release? Like with integration, there's a spectrum of reasonable options. You could release every push. Many tech companies do this. If you're deploying it once an iteration, you’re still in good company. Once a quarter is a bit sad. You could release once every two years. It seems absurdly slow now, but in the bad old days, this was the standard model in our industry.

下一个更难的问题是你应该多久发布一次？与集成一样，有一系列合理的选择。你可以释放每一次推动。许多科技公司都这样做。如果您在迭代中部署一次，那么您仍然处于良好的状态。一个季度一次有点难过。你可以每两年发布一次。现在看起来慢得离谱，但在过去的糟糕日子里，这是我们行业的标准模式。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-7-1615824704871.jpg)

.jpg）

**How often should you deploy to production?**

**您应该多久部署到生产环境？**

What makes deploying to production every push possible that deploying is not the same as releasing. If our new code is too incomplete or too scary to actually show to users, we can still deploy it, but keep it hidden. We can have the code actually in the production code base, but nothing is wired to it. That's pretty safe. If we’re already a bit too entangled for that, we can use feature flags to flip function on and off. If we’re feeling more adventurous, we can do A/B or friends and family testing so only a tiny portion of users see our scary code. Canary deploys are another variation for pre-detecting nightmares, before they hit mainstream usage. 

是什么让每次推送都可以部署到生产环境中，部署与发布不同。如果我们的新代码太不完整或太吓人而无法实际显示给用户，我们仍然可以部署它，但将其隐藏。我们可以在生产代码库中实际拥有代码，但没有任何连接。那相当安全。如果我们对此已经有点纠结了，我们可以使用功能标志来打开和关闭功能。如果我们更喜欢冒险，我们可以进行 A/B 或朋友和家人测试，这样只有一小部分用户会看到我们可怕的代码。 Canary 部署是另一种预检测噩梦的变体，在它们进入主流使用之前。

Not releasing has two bad consequences. It lengthens feedback cycles, which can impact decision making and makes engineers feel sad. Economically, it also means there’s inventory (the working software) sat on the shelf, rather than getting out to customers. Lean principles tell us that having  inventory sat around, not generating returns,  is waste.

不释放有两个不好的后果。它延长了反馈周期，这会影响决策并让工程师感到难过。从经济上讲，这也意味着货架上有库存（工作软件），而不是发给客户。精益原则告诉我们，闲置库存而不产生回报是浪费。

Then the conversation is,  why can't we release this? What's stopping more frequent deployments? Many organisations fear their microservices, and they want to do integration testing of the whole assembly, usually manual integration testing. One customer, with about 60 microservices, wanted to ensure that there was no possibility that some bright spark of an engineer could release one microservice without releasing the other 59 microservices. To enforce this, they had one single pipeline for all of the microservices in a big batch. This obviously is not the value proposition of microservices, which is that they are independently deployable. Sadly, it was the way that they felt safest to do it.

然后对话是，为什么我们不能发布这个？是什么阻止了更频繁的部署？许多组织害怕他们的微服务，他们想对整个程序集进行集成测试，通常是手动集成测试。一位拥有大约 60 个微服务的客户希望确保不可能出现工程师的一些闪光点在不发布其他 59 个微服务的情况下发布一个微服务。为了实施这一点，他们为大批量的所有微服务建立了一个单一的管道。这显然不是微服务的价值主张，即它们是可独立部署的。可悲的是，这是他们认为最安全的方式。

We also see a reluctance actually to deliver because of concerns about quality and completeness. Of course, these aren't ridiculous. You don't want to anger your customers. On the other hand, as Reid Hoffman said, if you're not embarrassed by your first release, it was too late. There is a value in continuous improvement, and there is value in getting things being used.

我们还看到，由于对质量和完整性的担忧，人们实际上不愿交付。当然，这些都不是开玩笑的。您不想激怒您的客户。另一方面，正如 Reid Hoffman 所说，如果你不为你的第一个版本感到尴尬，那就太晚了。持续改进是有价值的，让东西被使用也有价值。

If releases are infrequent and monolithic,  you've got these beautiful microservices architecture that allows you to go faster, and yet you're going really slow. This is bad business, and it’s bad engineering.

如果发布不频繁且单一，那么您将拥有这些漂亮的微服务架构，可以让您运行得更快，但您的运行速度却非常缓慢。这是糟糕的生意，也是糟糕的工程。

Let’s assume you go for the frequent deploys. All of the things which protect your users from half-baked features, like the automated testing, the feature flags, the A/B testing, the SRE, need substantial automation. Often when I start working with a customer, we have a question about testing, and they say, "Oh, our tests aren't automated." What that really means is that they don't actually know if the code works at any particular point. They hope it works, and it might have worked last time they checked, but we don't have any way of knowing whether it works right now without running manual tests.

假设您进行频繁部署。所有保护用户免受半生不熟的功能影响的东西，比如自动化测试、功能标志、A/B 测试、SRE，都需要大量的自动化。通常，当我开始与客户合作时，我们对测试有疑问，他们会说：“哦，我们的测试不是自动化的。”这真正意味着他们实际上并不知道代码是否在任何特定点工作。他们希望它有效，而且上次检查时可能有效，但是如果不运行手动测试，我们无法知道它现在是否有效。

The thing is, regressions happen. Even if all the engineers are the most perfect engineers, there’s an outside world which is less perfect. Systems they depend on might behave unexpectedly. If a dependency update changes behavior, something will break even if nobody did anything wrong. That brings us back to “we can't ship because we don't have confidence in the quality”. Well, let's fix the confidence in the quality, and then we can ship.

问题是，回归发生了。即使所有的工程师都是最完美的工程师，也有一个不那么完美的外部世界。他们依赖的系统可能会出现意外行为。如果依赖项更新改变了行为，即使没有人做错任何事情，某些事情也会中断。这让我们回到“我们无法发货，因为我们对质量没有信心”。好吧，让我们修复对质量的信心，然后我们就可以发货了。

I talked about contract testing. That is cheap and easy and can be done at a unit test level, but of course, you do also need automated integration tests. You don't want to be relying on manual integration tests or they become a bottleneck. 

我谈到了合同测试。这既便宜又容易，并且可以在单元测试级别完成，但是当然，您还需要自动化集成测试。您不想依赖手动集成测试，否则它们会成为瓶颈。

“CI/CD” seems to have replaced “build” in our vocabularies, but in both cases, it is one of the most valuable things that you have as an engineering organization. It should be your friend, and it should be this pervasive presence everywhere. Sometimes the way the build works is that it's off on a Jenkins system somewhere. Someone who is a bit diligent goes and checks the web page every now and then and notices it's red and goes and tells their colleagues, and then eventually someone fixes the issue. What's much better is just a passive build indicator that everybody can see without opening up a separate page for. If the monitor goes red, it's really obvious, that something changed, and easy to look at the most recent change. A traffic light works if you have one project. If you've got microservices, you're probably going to need something like a set of tiles. Even if you don't have microservices, you're probably going to have several projects, so you need something a bit more complete than a traffic light, even though the traffic lights are cute.

在我们的词汇表中，“CI/CD”似乎已经取代了“build”，但在这两种情况下，它都是作为工程组织的最有价值的东西之一。它应该是你的朋友，它应该无处不在。有时构建的工作方式是它在某个 Jenkins 系统上关闭。有点勤奋的人时不时地去检查一下网页，发现它是红色的，然后去告诉他们的同事，然后最终有人解决了这个问题。更好的是只是一个被动的构建指示器，每个人都可以看到而无需打开单独的页面。如果显示器变红，那真的很明显，有些东西发生了变化，并且很容易查看最近的变化。如果您有一个项目，则交通灯会起作用。如果你有微服务，你可能需要像一组磁贴这样的东西。即使你没有微服务，你也可能会有几个项目，所以你需要比红绿灯更完整的东西，即使红绿灯很可爱。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-8-1615824703841.jpg)

.jpg）

**“We don’t know when the build is broken”**

**“我们不知道构建何时被破坏”**

If you invest in your build monitoring, then you end up with the broken window situation. I've arrived at customers, and the first thing I've done is I've looked at the build, and I said, "Oh, this build seems to be broken." They've said, "Yeah, it's been broken for a few weeks." At that point, I knew I had a lot of work to do!

如果您对构建监控进行投资，那么您最终会遇到窗口损坏的情况。我到达客户那里，我做的第一件事就是查看构建，然后我说，“哦，这个构建似乎坏了。”他们说，“是的，它已经坏了几个星期了。”那一刻，我知道我有很多工作要做！

Why is a perma-broken build bad? It means you can't do the automated integration testing because nothing is making it out of the build. In fact, you can’t even do manual integration testing, so inter-service compatibility could be deteriorating and no one would know.

为什么永久损坏的构建不好？这意味着您无法进行自动化集成测试，因为构建中没有任何内容。实际上，您甚至无法进行手动集成测试，因此服务间兼容性可能会恶化，而没有人会知道。

New regressions go undetected, because the build is already red. Perhaps worst of all, it creates a culture so that when one of the other builds goes red, people aren't that worried, because it's more of the same: "Now we've got two red. Perhaps we could get the whole set , and then it would match if we got them all red." Well, no, that's not how it should be.

新的回归没有被发现，因为构建已经是红色的。也许最糟糕的是，它创造了一种文化，当其他版本中的一个变红时，人们不会那么担心，因为它更相似：“现在我们有两个红色。也许我们可以得到整个系列，然后如果我们把它们都变成红色就会匹配。”嗯，不，这不应该是这样。

## The locked-down totally rigid, inflexible uncloudy Cloud

## 锁定的完全僵化，不灵活的无云云

These are all challenges which happen at the team level. They’re about how we as engineers manage ourselves and our code. But of course, particularly once you get to an organization of a certain size, you end up with another set of challenges, which is what the organization does with the Cloud. I have noticed that some organisations love to take the Cloud, and turn it into a locked down, totally rigid, flexible, un-cloudy Cloud.

这些都是在团队层面发生的挑战。它们是关于我们作为工程师如何管理自己和我们的代码。但是，当然，尤其是当您进入某个规模的组织时，您最终会遇到另一组挑战，这就是该组织对云所做的事情。我注意到一些组织喜欢使用云，并将其变成一个锁定的、完全僵化的、灵活的、无云的云。

How do you make a Cloud un-cloudy? You say, "Well, I know you could go fast, and I know all of your automation support is going fast, but we have a process. We have an architecture review board, and it meets rather infrequently." It will meet a month after the project is ready to ship, or in the worst case, it will meet a month after the project has shipped. We're going through the motions even though the thing has shipped. The architecture will be reviewed on paper after it’s already been validated in the field, which is silly.

你如何让云变得不多云？你说，“嗯，我知道你可以做得很快，我知道你所有的自动化支持都在快速进行，但我们有一个流程。我们有一个架构审查委员会，它很少会面。”它会在项目准备发货后一个月开会，或者在最坏的情况下，它会在项目发货后一个月开会。即使这东西已经发货，我们也正在处理这些议案。该架构在经过现场验证后将在纸上进行审查，这很愚蠢。

Someone told me a story once. A client came to them with a complaint that some provisioning software IBM had sold them didn’t work. What had happened was we'd promised that our nifty provisioning software would allow them to create virtual machines in ten minutes. This was several years ago, when “a VM in ten minutes” was advanced and cool. We promised them it would be wonderful. 

曾经有人给我讲过一个故事。一位客户向他们投诉，称 IBM 出售给他们的某些配置软件无法正常工作。发生的事情是我们承诺我们的漂亮配置软件将允许他们在 10 分钟内创建虚拟机。那是几年前，当时“十分钟一个 VM”是先进而酷的。我们向他们保证这会很棒。

When the client got it installed and started using it, they did not find it wonderful. They’d thought they were going to get a 10-minute provision time, but what they were seeing is that it took them three months to provision a Cloud instance. They came back to us, and they said, "your software is totally broken. You mis-sold it. Look, it's taking three months." We were puzzled by this, so we went in and did some investigation. It turns out what had happened was they had put an 84-step pre-approval process in place to get one of those instances.

当客户安装它并开始使用它时，他们并不觉得它很棒。他们原以为他们将获得 10 分钟的配置时间，但他们看到的是，他们花了三个月的时间来配置一个云实例。他们回来找我们，他们说：“你的软件完全坏了。你卖错了。看，这需要三个月的时间。”我们对此感到困惑，所以我们进去做了一些调查。事实证明，他们已经制定了一个 84 步的预批准流程来获得其中一个实例。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-9-1615824704594.jpg)

.jpg）

**“This provisioning software is broken”**

**“此配置软件已损坏”**

The technology was there, but the culture wasn't there, so the technology didn't work. This is sad. We take this cloud, it's a beautiful cloud, it has all these fantastic properties, it makes everything really easy, and then another part of the organization says, "Oh, that's a bit scary. We wouldn't want people actually to be able to do things. Let's put it in a cage!" That old-style paperwork-heavy governance is just not going to work – as well as being really annoying to everyone. It's not going to give the results. What’s worse, it's not actually going to make things more secure. It's probably going to make them less secure. It's definitely going to make things slower and cost money. We shouldn't be doing it.

技术在那里，但文化不在那里，所以技术不起作用。这真是难过;这真是伤心。我们拿了这朵云，它是一朵美丽的云，它具有所有这些奇妙的特性，它使一切变得非常简单，然后组织的另一部分说，“哦，这有点可怕。我们不希望人们实际上能够做事。我们把它关在笼子里吧！”那种老式的文书工作繁重的治理是行不通的——而且真的让每个人都讨厌。它不会给出结果。更糟糕的是，它实际上不会让事情变得更安全。这可能会使它们变得不那么安全。这肯定会让事情变慢并花钱。我们不应该这样做。

I talked to another client, a large automotive company, and they were having a real problem with their Cloud provisioning. It was taking a really long time to get instances. They thought, "The way we're going to fix this is we're going to move from Provider A to Provider B." That might have worked, except the slowness was actually with their internal procurement. Switching providers would bypass their established procurement processes, so it might speed things up for a while, but eventually, their governance team were going to notice that the new provider, and impose controls. Once that happened, they would put the regulation in place, and the status quo would be restored. They would have had all the cost of changing but not actually any of the benefits. It's a bit like— I'm sorry to say I have sometimes been tempted to do this— if you're looking at your stove, and you decide, "Oh, that oven is filthy. Cleaning it will be hard, so I' m going to move house, so I don't have to clean the oven." But then, of course, the same thing happens in the other house, and the new oven gets dirty. You need a more sustainable process than just switching providers to try to outfox your own procurement.

我与另一家客户，一家大型汽车公司进行了交谈，他们在云配置方面遇到了真正的问题。获取实例需要很长时间。他们认为，“我们要解决这个问题的方法是从提供商 A 转移到提供商 B。”这可能会奏效，除了缓慢实际上是由于他们的内部采购。更换供应商会绕过他们既定的采购流程，所以它可能会加快速度一段时间，但最终，他们的治理团队会注意到新的供应商，并实施控制。一旦发生这种情况，他们就会制定法规，并恢复现状。他们会承担改变的所有成本，但实际上没有任何好处。这有点像——我很抱歉地说我有时很想这样做——如果你看着你的炉子，你决定，“哦，那个烤箱很脏。清洁它会很困难，所以我”我要搬家了，所以我不用清理烤箱了。”但是，当然，同样的事情发生在另一所房子里，新烤箱变脏了。您需要一个更具可持续性的流程，而不仅仅是更换供应商以试图超越自己的采购。

If the developers are the only ones changing, if the developers are the only ones going Cloud-native, then it's just not going to work. That doesn’t mean a developer-driven free-for-all is the right model. If there isn’t some governance around it, then Cloud can become a mystery money pit. Many of us have had the problem of looking at a cloud bill and thinking "Hmm. Yeah, that large, and I don't understand where it's all going or who's doing it."

如果开发人员是唯一的改变者，如果开发人员是唯一的云原生者，那么它就是行不通的。这并不意味着开发人员驱动的免费模式是正确的模式。如果周围没有一些治理，那么 Cloud 可能会成为一个神秘的钱坑。我们中的许多人都遇到过这样的问题：看着云账单并想“嗯。是的，这么大，我不明白这一切都去哪里了，也不知道是谁在做的。”

It is so easy to provision hardware with the Cloud, but that doesn't mean the hardware is free. Someone still has to pay for it. Hardware being easy to provision also doesn’t guarantee the hardware is useful.

使用云配置硬件非常容易，但这并不意味着硬件是免费的。仍然需要有人为此付出代价。易于配置的硬件也不能保证硬件是有用的。

When I was first learning Kubernetes, I tried it out, of course. I created a cluster, but then I got side-tracked, because I had too much work in progress. After two months, I came back to my cluster and discovered this cluster was about £1000 a month … and it was completely value-free. That’s so wasteful I still cringe thinking about it. 

当我第一次学习 Kubernetes 时，我当然尝试过。我创建了一个集群，但后来我偏离了方向，因为我有太多的工作要做。两个月后，我回到我的集群，发现这个集群每月大约 1000 英镑……而且它完全没有价值。这太浪费了，我现在想想都害怕。

A lot of what our technology allows us to do is to make things efficient. Peter Drucker, the great management consultant, said “There is nothing so useless as doing efficiently that which should not be done at all.” Efficiently creating Kubernetes clusters with no value, that's not good. As well as being expensive, there's an ecological impact. Having a Kubernetes cluster consuming £1000 worth of electricity to do nothing is not very good for the planet.

我们的技术允许我们做的很多事情就是让事情变得高效。伟大的管理顾问彼得德鲁克说：“没有什么比有效地做根本不应该做的事情更无用的了。”高效地创建没有价值的 Kubernetes 集群，那不好。除了昂贵之外，还有生态影响。让一个 Kubernetes 集群消耗价值 1000 英镑的电力而无所事事，对地球来说并不是很好。

For many  of the problems I’ve described, what initially seems like a technology problem is actually a people problem. I think this one is a little bit different, because this one seems like a people problem and is actually a technology problem. This is an area where tooling actually can help. For example, tools can help us manage waste by detecting unused servers and helping us trace servers back to originators. The tooling for this isn’t there yet, but it’s getting more mature.

对于我描述的许多问题，最初看起来是技术问题实际上是人的问题。我认为这个有点不同，因为这看起来像是一个人的问题，实际上是一个技术问题。这是工具实际上可以提供帮助的领域。例如，工具可以通过检测未使用的服务器并帮助我们将服务器追溯到发起者来帮助我们管理浪费。用于此的工具尚不存在，但它变得越来越成熟。

## Cloud to manage your Cloud

## 云来管理您的云

This cloud-management tooling ends up being on the Cloud, so you end up in the recursion situation to have some Cloud to manage your clouds. My company has a multi-cloud manager that will look at your workloads, figure out the shape of the workload, what the most optimum provider you could have it on is financially, and then make that move automatically. I expect we'll probably start to see more and more software like this where it's looking at it and saying, "By the way, I can tell that there's actually no traffic to his Kubernetes cluster that's been sat there for two months. Why don 't you go have some words with Holly?"

这个云管理工具最终在云上，所以你最终处于递归情况，需要一些云来管理你的云。我的公司有一个多云管理器，可以查看您的工作负载，确定工作负载的形状，您可以拥有的最佳提供商在财务上是什么，然后自动采取行动。我预计我们可能会开始看到越来越多这样的软件，它正在查看它并说：“顺便说一句，我可以说他的 Kubernetes 集群实际上没有流量，已经在那里坐了两个月。为什么不呢？你不去跟霍莉说几句话吗？”

## Microservices Ops Mayhem

## 微服务运营混乱

Managing cloud costs is getting more complex, and this reflects a more general thing, which is that cloud ops is getting more complex. We're using more and more cloud providers. There are more and more Cloud instances springing up. We've got clusters everywhere, so how on earth do we do ops for this? This is where SRE ( Site Reliability Engineering) comes in.

管理云成本变得越来越复杂，这反映了一个更普遍的事情，那就是云运营变得越来越复杂。我们正在使用越来越多的云提供商。越来越多的云实例如雨后春笋般涌现。我们到处都有集群，那么我们到底如何为此进行操作呢？这就是 SRE（站点可靠性工程）的用武之地。

Site reliability engineering aims to make ops more reproducible and less tedious, in order to make services more reliable. One of the ways it does this is by automating everything, which I think is an admirable goal. The more we automate things like releases, the more we can do them, which is good for both engineers and consumers. The ultimate goal should be that releases aren’t an event; they’re business as usual.

站点可靠性工程旨在使操作更具可重复性且不那么乏味，从而使服务更可靠。它做到这一点的方法之一是自动化一切，我认为这是一个令人钦佩的目标。我们越是自动化发布之类的事情，我们就能做的越多，这对工程师和消费者都有好处。最终目标应该是发布不是事件；他们照常营业。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-10-1615824706223.jpg)

.jpg）

**Make releases deeply boring.**

**让发布变得非常无聊。**

What enables that boring-ness is that we have confidence in the recoverability, and it's the SRE who gives us confidence in the recoverability.

导致这种无聊的原因是我们对可恢复性有信心，而 SRE 让我们对可恢复性充满信心。

I've got another sad space story, this time from the Soviet Union. In the 80s, an engineer wanted to make an update to the code on a Soviet space probe called Phobos. At this time, it was machine code, all 0s and 1s, and all written by hand. Obviously, you don’t want to do a live update of a spacecraft hurtling around the earth with hand-written machine code, without some checks. Before any push, code would be passed through a validator, which was the equivalent of a linter for the machine code.

我还有另一个悲伤的太空故事，这次来自苏联。在 80 年代，一位工程师想要更新苏联太空探测器 Phobos 上的代码。这时候是机器码，全是0和1，都是手写的。显然，您不想在没有一些检查的情况下使用手写机器代码对绕地球飞驰的航天器进行实时更新。在进行任何推送之前，代码将通过验证器，这相当于机器代码的 linter。

This worked well, until the automated checker was broken, when changes needed to be made. An engineer said, "Oh, but I really want to do this change. I'll just bypass the automated checks and just do the push of my code to the space probe because, of course, my code is perfect." So they did a live update of a spacecraft hurtling around the earth with hand-written machine code, without checks. What could possibly go wrong? 

这很有效，直到自动检查器被破坏，需要进行更改时。一位工程师说，“哦，但我真的很想做这个改变。我会绕过自动检查，只是将我的代码推送到空间探测器，因为，当然，我的代码是完美的。”因此，他们使用手写机器代码对绕地球飞驰的航天器进行了实时更新，无需检查。什么可能出错？

What happened was a very subtle bug. Things seemed to be working fine. Unfortunately, the engineer had forgotten one zero on one of the instructions. This changed the instruction from the intended instruction to one which stopped the probe’s charging fins rotating. The Phobos had fins which turned to orient towards the sun so that it could collect solar power, no matter which way it was facing. Everything worked great for about two days, until the battery went flat. Once the probe ran out of power, there was nothing that they could do to revive it because the entire thing was dead.

发生的事情是一个非常微妙的错误。事情似乎运作良好。不幸的是，工程师在其中一条指令上忘记了一个零。这将指令从预期指令更改为停止探针充电鳍旋转的指令。火卫一的鳍转向朝向太阳，这样它就可以收集太阳能，无论它面向哪个方向。一切正常运行了大约两天，直到电池没电了。一旦探测器电量耗尽，他们就无法恢复它，因为整个东西都死了。

That is an example of a system that is completely unrecoverable. Once it is dead, you are never getting that back. You can't just do something and recover it to a clean copy of the space probe code, because it's up in space.

这是一个完全不可恢复的系统的例子。一旦它死了，你就再也找不回来了。您不能只是做某事并将其恢复为空间探测代码的干净副本，因为它在空间中。

Systems like this are truly unrecoverable. Many of us believe that all of our systems are almost as unrecoverable as the space probe, but in fact, very few systems are.

像这样的系统确实无法恢复。我们中的许多人认为我们所有的系统几乎都像太空探测器一样不可恢复，但实际上，很少有系统是这样的。

Where we really want to be is at the top end of this spectrum, where we can be back in milliseconds, with no data loss. If anything goes wrong, it's just, “ping, it's fixed”. That's really hard to get to, but there are a whole bunch of intermediate points that are realistic goals.

我们真正想要的地方是在这个范围的顶端，我们可以在几毫秒内返回，而不会丢失数据。如果出现任何问题，它只是“ping，它已修复”。这真的很难达到，但有一大堆中间点是现实的目标。

If we're fast in recovering, but data is lost, that's not so good, but we can live with that. If we have handoffs and manual intervention, then that will be a lot slower for the recovery. When we're thinking about deploying frequently and deploying with great boredom - we want to be confident that we're at that upper end. The way we get there, handoffs bad, automation, good.

如果我们的恢复速度很快，但数据丢失了，那不太好，但我们可以接受。如果我们有交接和人工干预，那么恢复速度会慢很多。当我们考虑频繁部署并且无聊地部署时 - 我们希望确信我们处于那个高端。我们到达那里的方式，交接不好，自动化，好。

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-11-1615824704082.jpg)

.jpg）

## Ways to succeed at Cloud Native

## 在 Cloud Native 上取得成功的方法

This article has included a whole bunch of miserable stories about things that I've seen that can go wrong. I don’t want to leave you with an impression that everything goes wrong all the time because a lot of the time, things do go really right. Cloud native is a wonderful way of developing software, which can feel better for teams, lower costs, and make happier users. As engineers, we can spend less time on the toil and the drudgery and more time on the things that we actually want to be doing… and we can get to market faster.

这篇文章包含了一大堆关于我所见过的可能出错的悲惨故事。我不想给你留下一切都会出错的印象，因为很多时候，事情确实是正确的。云原生是一种很好的软件开发方式，可以让团队感觉更好，降低成本，让用户更快乐。作为工程师，我们可以将更少的时间花在辛劳和苦力上，而将更多时间花在我们真正想做的事情上……而且我们可以更快地进入市场。

To get to that happy state, we have to have alignment across the organization. We can't have one group that says microservices, one group saying fast, and one group saying old-style governance. That's almost certainly not going to work, and there will be a lot of grumpy engineers and aggrieved finance officers. Instead, an organisation should agree, at a holistic level, what it’s trying to achieve. Once that goal is agreed, it should optimize for feedback, ensuring that feedback loops as short as possible, because that’s sound engineering.

为了达到这种快乐状态，我们必须在整个组织中保持一致。我们不能让一组说微服务，一组说快速，一组说旧式治理。这几乎肯定是行不通的，而且会有很多脾气暴躁的工程师和愤愤不平的财务人员。相反，一个组织应该在整体层面上就它试图实现的目标达成一致。一旦该目标达成一致，它应该针对反馈进行优化，确保反馈循环尽可能短，因为这是合理的工程。

## About the Author

##  关于作者

Holly Cummins** is an innovation leader in IBM Corporate Strategy, and spent several years as a consultant in the IBM Garage. As part of the Garage, she delivers technology-enabled innovation to clients across various industries, from banking to catering to retail to NGOs. Holly is an Oracle Java Champion, IBM Q Ambassador, and JavaOne Rock Star. She co-authored Manning's Enterprise OSGi in Action. 

Holly Cummins** 是 IBM 企业战略的创新领导者，曾在 IBM Garage 担任顾问多年。作为 Garage 的一部分，她为各行各业的客户提供技术支持的创新，从银行业到餐饮、零售再到非政府组织。 Holly 是 Oracle Java Champion、IBM Q 大使和 JavaOne Rock Star。她与人合着了 Manning 的 Enterprise OSGi in Action。

