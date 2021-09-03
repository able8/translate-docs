# DevOps, SRE, and Platform Engineering

# DevOps、SRE 和平台工程

[August 1, 2021](https://iximiuz.com/en/posts/devops-sre-and-platform-engineering/#) From: https://iximiuz.com/en/posts/devops-sre-and-platform-engineering/

I'm a Software Engineer at heart, SRE at day job, and tech storyteller at night.

我本质上是一名软件工程师，白天工作是 SRE，晚上是技术讲故事的人。

[Subscribe to my monthly newsletter](https://iximiuz.com/en/newsletter/) or [follow me on Twitter](https://iximiuz.com/en/newsletter/#twitter) for quality content on Containers, Kubernetes, Cloud Native stack, and Programming!

[订阅我的每月通讯](https://iximiuz.com/en/newsletter/) 或 [在 Twitter 上关注我](https://iximiuz.com/en/newsletter/#twitter) 以获取有关 Containers 的优质内容， Kubernetes、云原生堆栈和编程！

I compiled this thread on Twitter, and all of a sudden, it got quite  some attention. So here, I'll try to elaborate on the topic a bit more. Maybe it would be helpful for someone trying to make a career decision  or just improve general understanding of the most hyped titles in the  industry.

我在 Twitter 上编译了这个帖子，突然之间，它引起了相当多的关注。所以在这里，我将尝试详细说明这个主题。也许这对试图做出职业决定或只是提高对行业中最受炒作的头衔的一般理解的人会有所帮助。

DevOps, SRE, and Platform Engineering (thread)
Sharing my understanding of things after working in this domain for about two years.
Starting from the clearest one.
Dev - this is about application development, aka business logic. The only one that makes money for a company.

DevOps、SRE 和平台工程（线程）
在这个领域工作了大约两年后，分享我对事物的理解。
从最清晰的开始。
开发 - 这是关于应用程序开发，也就是业务逻辑。唯一为一家公司赚钱的人。

During my career, I used to work in teams and companies where as a  developer, I would push code to a repository and just hope that it would work well when some mythical system administrator would eventually take it to production. I also was in setups where I would need to provision bare-metal servers on Monday, figure out the deployment strategy on  Tuesday, write some business logic on Wednesday, roll it out myself on Thursday, and firefight a production incident on Friday. And all this  without even being aware of the existence of fancy titles like DevOps or SRE engineer.

在我的职业生涯中，我曾经在团队和公司工作，作为开发人员，我会将代码推送到存储库，只是希望当某个神秘的系统管理员最终将其投入生产时，它会运行良好。我还需要在星期一提供裸机服务器，在星期二制定部署策略，在星期三编写一些业务逻辑，在星期四自己推出它，并在星期五解决生产事故。所有这一切甚至都没有意识到像 DevOps 或 SRE 工程师这样的花哨头衔的存在。

But then people around me started talking DevOps and SRE, comparing them with each other, and compiling [awesome lists](https://github.com/dastergon/awesome-sre) of [resources](https://github.com/AcalephStorage/awesome-devops). New job opportunities began emerging, and I quickly jumped into the SRE train. So, below is my experience of being involved in all things SRE  and Platform Engineering from the former Software Developer standpoint. And yeah, I think it's applicable primarily for companies where the  product is some sort of a web-facing service. This is the kind of  company I spent ten years working for. People doing embedded software or implementing databases probably live in totally different realities.

但后来我周围的人开始谈论 DevOps 和 SRE，将它们相互比较，并编译 [resources](https://github.com) 的 [awesome list](https://github.com/dastergon/awesome-sre)/AcalephStorage/awesome-devops)。新的工作机会开始出现，我很快就跳上了 SRE 的火车。因此，以下是我从前软件开发人员的角度参与所有 SRE 和平台工程的经验。是的，我认为它主要适用于产品是某种面向网络的服务的公司。这就是我工作了十年的那种公司。从事嵌入式软件或实施数据库的人们可能生活在完全不同的现实中。

## What is Development

## 什么是开发

This one is the simplest to explain. Development - is about application programming, i.e., writing the business logic of your main product. This is the only activity among the three ones being discussed  here that directly makes money for the company.

这是最简单的解释。开发 - 是关于应用程序编程，即编写主要产品的业务逻辑。这是这里讨论的三个活动中唯一直接为公司赚钱的活动。

> The only one that makes for a company is of course sales, everything else is expenditure :)

> 对一家公司来说，唯一能赚钱的当然是销售，其他的都是支出 :)

IMO, development is super hot! As a developer, you quickly start  thinking that you are the most important person around. Without your  code, there is nothing. But apparently, just writing code often isn't  enough. The code needs to be delivered to production and executed there.

IMO，开发超级火爆！作为开发人员，您很快就会开始认为自己是周围最重要的人。没有你的代码，什么都没有。但显然，仅仅经常编写代码是不够的。代码需要交付到生产环境并在那里执行。

I'd been carrying the Software Developer (or Software Engineer) title since the very beginning of my career in 2011. And I still remember the pain quite vividly - I always wished to have control over deploying my  code. And I rarely had it. Instead, there would be some obscure  procedure when someone, usually not even your senior colleague, would  have access to production servers and deploy the code there for you. So, if after pushing the changes to the repository, you got unlucky enough  to notice a bug only on the live version of your service, you'd need to  beg for an extra rollout. It most definitely sucked.

自 2011 年职业生涯开始以来，我一直拥有软件开发人员（或软件工程师）的头衔。我仍然清楚地记得那种痛苦——我一直希望能够控制我的代码部署。我很少有它。相反，当有人（通常甚至不是您的高级同事）可以访问生产服务器并在那里为您部署代码时，将会有一些模糊的过程。因此，如果在将更改推送到存储库后，您不幸发现仅在服务的实时版本上存在错误，则需要请求额外的部署。它绝对很烂。

## What is DevOps 

## 什么是 DevOps

I'll not even try to quote the official definition here. Instead,  I'll share the first-hand experience. For me, DevOps was a cultural  shift giving development teams more control over shipping code to  production. The implementation could vary. I've been in setups where  developers would just have `sudo` on production servers. But probably the most common approach is to provide development teams with some sort of CI/CD pipelines.

我什至不会在这里引用官方定义。相反，我将分享第一手经验。对我来说，DevOps 是一种文化转变，让开发团队可以更好地控制将代码交付到生产环境。实施可能会有所不同。我一直在开发人员在生产服务器上只有`sudo`的设置。但最常见的方法可能是为开发团队提供某种 CI/CD 管道。

In an ideal GitOps world, developers would still be just pushing code to repositories. However, there would be a magical button somewhere at  the team's disposal that would put the new version on live or maybe even provision a new piece of infrastructure to cover the new requirements.

在理想的 GitOps 世界中，开发人员仍然只是将代码推送到存储库。然而，在团队可以使用的某个地方会有一个神奇的按钮，可以让新版本上线，甚至可能提供一个新的基础设施来满足新的要求。

The original idea of DevOps is probably much broader than just that. But from what I see in the job descriptions, what I hear from recruiters trying to hunt me for a DevOps position, and what I managed to gather  from my fellow colleagues carrying the DevOps engineer title, most of  the time, it's about creating an efficient way to deploy stuff produced  by Development. In more advanced setups, DevOps may also be concerned  with other things improving the Development velocity. But DevOps itself  is never concerned with the actual application business logic.

DevOps 的最初想法可能远不止于此。但从我在工作描述中看到的，我从招聘人员那里听到的试图寻找我担任 DevOps 职位的消息，以及我设法从拥有 DevOps 工程师头衔的同事那里收集到的信息，大多数情况下，这是关于创建一个高效的部署开发产生的东西的方式。在更高级的设置中，DevOps 还可能关注其他提高开发速度的事情。但是 DevOps 本身从不关心实际的应用程序业务逻辑。

## What is SRE

## 什么是SRE

There is a [excellent series of books by Google](https://sre.google/books/) explaining the idea of the Site Reliability Engineering and, what's  even more important for me, sharing some real tech practices conducted  by Google SREs. In particular, it says that SRE is just one of the ways  to implement the DevOps culture - `class SRE implements DevOps {}`.

有一个 [Google 出品的优秀丛书](https://sre.google/books/) 解释了站点可靠性工程的思想，对我来说更重要的是，分享了一些由 Google SRE 进行的真实技术实践。特别是，它说 SRE 只是实现 DevOps 文化的方法之一——`class SRE 实现 DevOps {}`。

This explanation didn't really help me much. But what was even more  puzzling, subconsciously, I always felt excited while reading SRE job  descriptions and got bored quickly by the DevOps ones... So, there was  clearly a difference but, for a long time, I couldn't distill it.

这个解释并没有真正帮助我。但更令人费解的是，潜意识里，我总是在阅读 SRE 职位描述时感到兴奋，并且很快就对 DevOps 感到厌烦……所以，明显有区别，但很长一段时间，我无法提炼它。

Of course, that's just about my personal preferences, but whenever  someone mentions configuring a CI/CD pipeline, I always got depressed. And the DevOps job descriptions nowadays are full of such  responsibilities. Don't get me wrong, CI/CD pipelines are amazing! I'm  always glad when I have a chance to use one. But setting them up isn't a thing I enjoy the most. On the contrary, when someone asks me to jump  in and take a look at a bleeding production, be it chasing a bug, a  memory leak, or performance degradation, I'm always more than just happy to help.

当然，这只是我个人的喜好，但是每当有人提到配置 CI/CD 管道时，我总是感到沮丧。而如今的 DevOps 工作描述中充满了这样的职责。不要误会我的意思，CI/CD 管道很棒！当我有机会使用它时，我总是很高兴。但是设置它们并不是我最喜欢的事情。相反，当有人让我参与并查看正在流血的产品时，无论是在寻找错误、内存泄漏还是性能下降，我总是乐于提供帮助。

Developing code and shipping it to production still doesn't give you  the full picture. Someone needs to keep the production alive and  healthy! And that's how I see the place of SRE in my model of the world.

开发代码并将其交付到生产环境仍然不能为您提供全貌。有人需要让生产保持活力和健康！这就是我如何看待 SRE 在我的世界模型中的地位。

Google's SRE book focuses on monitoring and alerting, defining SLOs  of your services and tracking error budgets, incident response and  postmortems. These are the things one would need to apply to make the  production reliable. Facebook has a famous Production Engineer role, but it's pretty hard to distinguish it from a typical SRE role, judging  only by the job description.

Google 的 SRE 书籍侧重于监控和警报、定义服务的 SLO 以及跟踪错误预算、事件响应和事后分析。这些是人们需要应用才能使生产可靠的事情。 Facebook 有一个著名的生产工程师角色，但很难将其与典型的 SRE 角色区分开来，仅从职位描述来看。

Here is also a great tweet that kind of confirms my feeling that the primary focus of SRE is production.

这也是一条很棒的推文，它证实了我的感觉，即 SRE 的主要重点是生产。

My very simplified answer when someone says what is the difference between SRE and DevOps.

当有人说 SRE 和 DevOps 之间有什么区别时，我的回答非常简单。

* SRE = focused primarily on production
* DevOps = focused primarily on CI/CD and developer velocity

* SRE = 主要专注于生产
* DevOps = 主要关注 CI/CD 和开发人员速度

And one more:

I like it! My typical answer is:
SRE works from Production backward. DevOps works from development forward. Somewhere in the middle, they meet.

还有一个：

我喜欢！我的典型回答是：
SRE 从生产向后工作。 DevOps 从开发开始工作。在中间的某个地方，他们相遇了。

So, DevOps keeps production fresh. SRE keeps production healthy.

因此，DevOps 使生产保持新鲜。 SRE 保持生产健康。

## What is Platform Engineering 

## 什么是平台工程

When I used to be the only engineer in a startup, a decent part of my job was to turn some generic resources I'd rent from the infrastructure provider into something more tailored for the company's needs. So, I  had a bunch of scripts to provision a new server, some understanding of  how to provide network connectivity between our servers in different  data centers, some skills to replicate the production setup on staging,  and maybe even write one or two daemons to help me with log collection. I didn't really understand it, but these things constituted our Platform.

当我曾经是一家初创公司的唯一工程师时，我工作的一个体面部分是将我从基础设施提供商那里租用的一些通用资源转化为更适合公司需求的东西。所以，我有一堆脚本来配置一个新服务器，对如何在不同数据中心的服务器之间提供网络连接有一些了解，一些在登台时复制生产设置的技能，甚至可能编写一两个守护进程来帮助我用日志收集。我不是很理解，但这些东西构成了我们的平台。

Joining a much bigger company and starting consuming infra-related  resources brought me to a realization that there is a third area of  focus that might be quite close to DevOps and SRE. It's called Platform  Engineering.

加入一家更大的公司并开始使用与基础设施相关的资源，这让我意识到还有第三个重点领域可能与 DevOps 和 SRE 非常接近。它被称为平台工程。

From my understanding, Platform Engineering focuses on developing an  ecosystem that can be efficiently used from the Dev, Ops, and SRE  standpoints.

根据我的理解，平台工程专注于开发一个可以从 Dev、Ops 和 SRE 的角度有效使用的生态系统。

There might be quite some code writing in Platform Engineering. Or,  it could be mostly about configuring things. But again, it's not about  the primary business logic of your product - it's about making some  basic infrastructure more suitable for the day-to-day needs.
Platform Engineering - this is about infrastructure development.
PE focuses on creating a platform that can be efficiently used from the Dev, Ops, and SRE standpoints.
There is plenty of actual code writing in PE, but again, it's not about the primary business logic.

平台工程中可能有相当多的代码编写。或者，它可能主要是关于配置事物。但同样，这与您产品的主要业务逻辑无关——而是关于使一些基本基础设施更适合日常需求。
平台工程 - 这是关于基础设施开发。
PE 专注于创建一个可以从 Dev、Ops 和 SRE 的角度有效使用的平台。
在 PE 中有大量实际代码编写，但同样，这与主要业务逻辑无关。

To be honest, I don't see a contradiction between my way of seeing  Platform Engineering and the explanation from this tweet. Development  needs infrastructure to run the code. So, if Platform Engineering is  about enabling others to do whatever they want to do, at least in part,  it should be concerned with infrastructure development.

老实说，我认为我看待平台工程的方式与这条推文的解释之间没有矛盾。开发需要基础设施来运行代码。因此，如果平台工程的目的是让其他人做他们想做的任何事情，至少在一定程度上，它应该关注基础设施开发。

I have a feeling that in a bigger setup, when a company would have  thousands of bare-metal servers in its own data centers, a Platform  Engineering would start from managing this fleet of machines. So, some  sort of inventory software might need to be installed or even developed  internally. Installing operating systems and basic packages on the  servers being provisioned would probably also fall into the Platform  Engineering responsibility.

我有一种感觉，在更大的设置中，当一家公司在自己的数据中心拥有数千台裸机服务器时，平台工程师将从管理这组机器开始。因此，可能需要在内部安装甚至开发某种库存软件。在被配置的服务器上安装操作系统和基本包也可能属于平台工程的职责。

Luckily, clouds made Platform Engineering operating on much higher  layers. All the basic fleet management tasks are already solved for you. And even orchestration of your workloads is solved by projects like  Kubernetes or AWS ECS. However, the solution is quite generic, while  your teams are likely to deploy pretty similar microservices. So,  providing them with a default project template that would be integrated  with the company's metrics and logs collection subsystems would make  things moving much faster.

幸运的是，云使平台工程在更高的层上运行。所有基本的车队管理任务都已为您解决。甚至工作负载的编排也由 Kubernetes 或 AWS ECS 等项目解决。但是，该解决方案非常通用，而您的团队可能会部署非常相似的微服务。因此，为他们提供一个默认的项目模板，该模板将与公司的指标和日志收集子系统集成在一起，这将使事情进展得更快。

## What's about titles?

## 标题是什么？

So far, I was deliberately avoiding talking about roles and titles. Development, Operations, SRE, and Platform Engineering for me are about  areas of focus. And to a much lesser extent about titles. One person can be a Dev this week, then an Ops on the next week, and an SRE on the  week after.

到目前为止，我故意避免谈论角色和头衔。对我来说，开发、运营、SRE 和平台工程都是重点领域。而关于标题的程度要小得多。一个人可以在本周成为 Dev，然后在下周成为 Ops，在下周成为 SRE。

From my experience, the separation between Dev, Ops, SRE, and PE  becomes more apparent when the company size gets bigger. A bigger  company size usually means more specialists and fewer generalists. That's how you end up with dedicated SRE teams and a Platform  Engineering department. But of course, it's not a strict rule. For  instance, with my SRE title, I spent like a year doing all things true  SRE (SLO, monitoring, alerting, incident response) and then transitioned into Platform Engineering, where I do more infra development than  traditional SRE. YMMV.

根据我的经验，当公司规模变大时，Dev、Ops、SRE 和 PE 之间的分离变得更加明显。更大的公司规模通常意味着更多的专家和更少的通才。这就是您最终拥有专门的 SRE 团队和平台工程部门的方式。但当然，这不是一个严格的规则。例如，凭借我的 SRE 头衔，我花了大约一年的时间做所有真正的 SRE（SLO、监控、警报、事件响应），然后过渡到平台工程，在那里我比传统的 SRE 做更多的基础设施开发。天啊。

## Where Security goes?

## 安全去哪儿了？

> Awesome , but where the security team gets involved from DevOps and SRE prospective. 

> 很棒，但是安全团队从 DevOps 和 SRE 的角度参与进来。

That's a very good question! But I don't have a simple answer. For  me, a reasonable approach is to make security a cross-cutting theme in  all Dev, Ops, SRE, and PE. Different security concerns can be addressed  on different layers using different tools. For instance, Development  could be concerned with preventing SQL injections while Platform folks  could harden the networking by configuring some fancy cilium policies.

这是一个很好的问题！但我没有一个简单的答案。对我来说，一个合理的方法是让安全成为所有 Dev、Ops、SRE 和 PE 的交叉主题。可以使用不同的工具在不同的层上解决不同的安全问题。例如，开发人员可能会关注防止 SQL 注入，而平台人员可以通过配置一些奇特的 cilium 策略来强化网络。

## Instead of conclusion

## 而不是结论

Don't forget, all the things above are IMO 😉 

别忘了，以上都是IMO 😉
