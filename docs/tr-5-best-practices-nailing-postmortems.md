# 5 Best Practices for Nailing Incident Retrospectives

# 5 个极好的事件回顾的最佳实践

December 26, 2019

Reading about incident retrospectives (or postmortems) is different from seeing them in action. Retrospectives are like snowflakes; no two will ever look the same. There isn’t a template that will work in every situation, but there are some best practices that can help. Here are five practices to take your retrospectives to a new level. Each practice has an example that proves the method

阅读事件回顾（或事后分析）与在行动中看到它们不同。回顾就像雪花；没有两个看起来是一样的。没有适用于所有情况的模板，但有一些最佳实践可以提供帮助。这里有五种做法可以将您的回顾展提升到一个新的水平。每个练习都有一个例子来证明方法

### Use visuals

### 使用视觉效果

As [Steve McGhee](https://www.blameless.com/improve-postmortem-with-sre-steve-mcghee/) says, “A 'what happened' narrative with graphs is the best textbook-let for teaching other engineers how to get better at progressing through future incidents.” Days, weeks, or even years after a retrospective is written, graphs still provide an engineer with a quick and in-depth explanation for what was happening during the incident.

正如 [Steve McGhee](https://www.blameless.com/improve-postmortem-with-sre-steve-mcghee/) 所说，“用图表讲述‘发生了什么’是教其他工程师的最佳教科书如何更好地应对未来的事件。”在编写回顾的几天、几周甚至几年之后，图表仍然为工程师提供了对事件期间发生的事情的快速而深入的解释。

In a [Cloudflare retrospective](https://blog.cloudflare.com/details-of-the-cloudflare-outage-on-july-2-2019/) authors use visuals to help readers understand the background for a DNS outage. The retrospective reads, “Unfortunately, last Tuesday’s update contained a regular expression that backtracked enormously and exhausted CPU used for HTTP/HTTPS serving. This brought down Cloudflare’s core proxying, CDN and WAF functionality. The following graph shows CPUs dedicated to serving HTTP/HTTPS traffic spiking to nearly 100% usage across the servers in our network.” A graph showing the CPU usage during the incident follows:

在 [Cloudflare 回顾](https://blog.cloudflare.com/details-of-the-cloudflare-outage-on-july-2-2019/) 中，作者使用视觉效果帮助读者了解 DNS 中断的背景。回顾中写道：“不幸的是，上周二的更新包含一个正则表达式，该表达式大量回溯并耗尽了用于 HTTP/HTTPS 服务的 CPU。这导致 Cloudflare 的核心代理、CDN 和 WAF 功能失效。下图显示了专用于服务 HTTP/HTTPS 流量的 CPU，在我们网络中的服务器上使用率飙升至近 100%。”显示事件期间 CPU 使用率的图表如下：

![](https://uploads-ssl.webflow.com/5ec0224560bd6a6ef89a51ae/5ec6cf7291d49faf237a835e_733A157F-B3E4-4BAC-A119-51F73CF73662.jpeg)

Visuals embedded within the retrospective benefit readers in two major ways. First, this allows new hires to visualize the problem. They can feel like they’re working through the incident with the engineers who mitigated it. Second, it allows engineers who may handle a similar issue to find the information they’re looking for faster.

嵌入在回顾中的视觉效果以两种主要方式使读者受益。首先，这允许新员工将问题形象化。他们会觉得他们正在与减轻事件的工程师一起解决事件。其次，它允许可能处理类似问题的工程师更快地找到他们正在寻找的信息。

### Be a historian

### 成为历史学家

Using timelines when writing retrospectives is very valuable. But, there’s an art to crafting them. As [Steve McGhee](https://www.blameless.com/improve-postmortem-with-sre-steve-mcghee/) says, “There is little utility to including the entire chat log of an incident. Instead, consider illustrating a timeline of the important inflection points (e.g. actions that turned the situation around). This may prove to be very helpful for troubleshooting future incidents.” Retrospective timelines need the perfect balance of information. Too much to sift through, and the retrospective will become cluttered. Too little and it’s vague.

在编写回顾时使用时间线是非常有价值的。但是，制作它们是一门艺术。正如 [Steve McGhee](https://www.blameless.com/improve-postmortem-with-sre-steve-mcghee/)所说，“包含事件的整个聊天日志几乎没有用处。相反，请考虑说明重要转折点的时间表（例如扭转局势的行动)。这可能对解决未来的事件非常有帮助。”追溯时间线需要信息的完美平衡。筛选太多，回顾会变得杂乱无章。太少了，而且很模糊。

Twilio's “ [Billing Incident Post-Mortem: Breakdown, Analysis and Root Cause](https://www.twilio.com/blog/2013/07/billing-incident-post-mortem-breakdown-analysis-and-root-cause.html),” shows this balance. What Twilio does well in this retrospective is clarity. In this particular incident, the timeline and contributing factor are separate. In the entry for 1:35 AM July 18, the timeline note reads, “We experienced a loss of network connectivity between all of our billing redis-slaves and our redis-master. This caused all redis-slaves to reconnect and request full synchronization with the master at the same time.”

Twilio 的“[计费事件事后分析：故障、分析和根本原因](https://www.twilio.com/blog/2013/07/billing-incident-post-mortem-breakdown-analysis-and-root-cause.html)”显示了这种平衡。 Twilio 在这次回顾中做得很好的是清晰度。在此特定事件中，时间线和促成因素是分开的。在 7 月 18 日凌晨 1 点 35 分的条目中，时间线注释写道：“我们的所有计费 redis-slaves 和 redis-master 之间的网络连接丢失。这导致所有 redis-slave 重新连接并同时请求与 master 完全同步。”

In the root cause analysis, the retrospective authors add details for this time stamp. They explain that the loss of network connectivity “caused all redis-slaves to reconnect and request full synchronization with the master at the same time." They also state how this affected the redis-master.

在根本原因分析中，回顾作者为此时间戳添加了详细信息。他们解释说，网络连接的丢失“导致所有 redis-slave 重新连接并同时请求与 master 完全同步。”他们还说明了这如何影响 redis-master。

The timeline entry is half the word count of the explanation in the analysis. Yet, it still relays the most crucial information. This allows engineers to reap the benefit of speed.

时间线条目是分析中解释字数的一半。然而，它仍然传递最关键的信息。这使工程师能够从速度中获益。

If the billing redis-slaves disconnect again, an engineer might treat this retrospective as a clue. When retrospective timelines include only the most important moments, an engineer can save time sifting through clutter.

如果计费 redis-slaves 再次断开连接，工程师可能会将这次回顾视为一个线索。当追溯时间线仅包括最重要的时刻时，工程师可以节省筛选杂乱的时间。

### Publish promptly 

### 及时发布

As the [Google SRE book](https://landing.google.com/sre/workbook/chapters/postmortem-culture/) says, “A prompt postmortem tends to be more accurate because information is fresh in the contributors’ minds. The people who were affected by the outage are waiting for an explanation and some demonstration that you have things under control. The longer you wait, the more they will fill the gap with the products of their imagination. That seldom works in your favor!”

正如 [Google SRE book](https://landing.google.com/sre/workbook/chapters/postmortem-culture/) 所说，“及时的事后分析往往更准确，因为信息在贡献者的脑海中是新鲜的。受停电影响的人们正在等待解释和一些证明您已控制一切的证明。你等待的时间越长，他们就越会用他们想象的产品来填补空白。这很少对你有利！”

Promptness has two main benefits. First, it allows the authors of the retrospective to report on the incident with a clear mind. Second, it soothes affected customers with less opportunity for churn.

及时性有两个主要好处。首先，它允许回顾的作者以清晰的头脑报告事件。其次，它可以安抚受影响的客户，减少流失的机会。

Google and other best-in-class companies like Uber practice what they preach. These companies often publish retrospectives within 48 hours. This discipline leads to more accurate retrospectives. After two months, will you remember the specifics of an incident, even after looking at the logs? It’s not likely.

谷歌和优步等其他一流公司践行他们的主张。这些公司通常会在 48 小时内发布回顾。这种纪律导致更准确的回顾。两个月后，即使查看了日志，您还会记得事件的具体细节吗？不太可能。

By publishing retrospectives within two days of mitigation, the information is fresher. This makes it more useful for teaching/onboarding and future reference.

通过在缓解后两天内发布回顾，信息更加新鲜。这使得它对教学/入职和未来参考更有用。

Prompt retrospectives are also crucial to foster a culture of transparency. If an incident affects your customers, it’s likely they’ll feel upset. In the case of an incident involving critical features, billing, or data breaches, customers will often be on edge waiting for an explanation. Some of your customers may even have SLAs set for the promptness of a retrospective. Waiting to publish only increases customer dissatisfaction. But, if teams communicate via a detailed, accurate retrospective, customers don’t have to remain anxious.

及时的回顾对于培养透明文化也至关重要。如果事件影响到您的客户，他们很可能会感到不安。如果发生涉及关键功能、计费或数据泄露的事件，客户通常会迫不及待地等待解释。您的一些客户甚至可能设置了 SLA 以加快回顾。等待发布只会增加客户的不满。但是，如果团队通过详细、准确的回顾进行沟通，客户就不必保持焦虑。

### Be blameless

### 无可指责

Blameless retrospectives are often referred to when talking about best practices. But, what does blameless culture actually look like? When writing blameless retrospectives, there are 3 important things to keep in mind.

在谈论最佳实践时，通常会提到无可指责的回顾。但是，无可指责的文化究竟是什么样子的呢？在编写无可指责的回顾时，需要牢记 3 件重要的事情。

- **People are not points of failure.** Pinning an incident on one person, or a group of people is counterproductive. It creates an environment where people are afraid to take risks, innovate, and problem solve. This leads to stagnancy and avoidance.
- **Everyone on the team is working with good intentions.** People make mistakes. It’ rare for a team member to cause problems on purpose. Everyone is doing what makes the most sense to them at the time to be helpful.
- **Failure will happen.** There’s no way around it. But, by having a good incident resolution and retrospective practice in place, failure can actually be beneficial. It uncovers areas to focus on to improve resiliency. As long as you learn from an incident, you’ve made progress.


- **人不是故障点。**将事件归咎于一个人或一群人会适得其反。它创造了一个人们害怕冒险、创新和解决问题的环境。这导致停滞和回避。
- **团队中的每个人都怀着良好的意愿工作。** 人们会犯错误。团队成员故意制造问题的情况很少见。每个人都在做当时对他们最有意义的事情来提供帮助。
- **失败会发生。** 没有办法解决。但是，通过良好的事件解决方案和追溯实践，失败实际上是有益的。它揭示了需要关注的领域，以提高弹性。只要你从事件中吸取教训，你就已经取得了进步。


Many teams choose to have a meeting after an incident to talk through what happened. Etsy created an introduction to this meeting that voices the 3 above points.

许多团队选择在事件发生后召开会议来讨论发生的事情。 Etsy 对这次会议进行了介绍，表达了上述 3 点。

In Etsy's [Debriefing Facilitation Guide](https://extfiles.etsy.com/DebriefingFacilitationGuide.pdf) it states, “The goal for our time together today is to recreate the event, talking through what happened for each person at each stage in order to create as robust a portrait as possible of what happened, and what the circumstances in play were at each juncture (when decisions were made, and actions were taken) that made it make sense for people to do what they did in the moment. If one of you gains an insight into the complexity of another person’s role, this was an hour well spent.” 

在 Etsy 的 [Debriefing Facilitation Guide](https://extfiles.etsy.com/DebriefingFacilitationGuide.pdf)中，它指出：“我们今天在一起的时间的目标是重新创建事件，讲述每个人在每个阶段发生的事情为了尽可能地描绘所发生的事情，以及在每个关键时刻（做出决定和采取行动时)的情况，这使得人们做他们当时所做的事情是有意义的。如果你们中的一个人洞察了另一个人角色的复杂性，那么这一个小时就花得很值了。”

[Sentry’s retrospective](https://blog.sentry.io/2016/06/14/security-incident-june-12-2016) from a security incident occurring July 12, 2016 demonstrates this. The retrospective uses the collective “we” pronoun to avoid naming people as problems. Additionally, it states “It’s been a valuable experience for our product team, albeit one we wish we could have avoided.” The point here is that this was a learning experience. Failure happened and will happen again. Sure, incidents are painful, but they’re one of the best ways to learn and become better.

[Sentry 的回顾](https://blog.sentry.io/2016/06/14/security-incident-june-12-2016) 来自 2016 年 7 月 12 日发生的安全事件证明了这一点。回顾使用集体代词“我们”来避免将人们称为问题。此外，它还指出“这对我们的产品团队来说是一次宝贵的经历，尽管我们希望我们可以避免这种经历。”这里的重点是，这是一次学习经历。失败已经发生，而且还会再次发生。当然，事件是痛苦的，但它们是学习和变得更好的最佳方式之一。

### Tell a story

###  讲一个故事

An incident is a story. To tell a story well, many components must work together.

一个事件就是一个故事。为了讲好一个故事，许多组件必须协同工作。

- Without enough background knowledge, this story loses depth and context.
- Without a plan to rectify outstanding action items, the story loses a resolution.
- Without a timeline dictating what happened, the story loses its plot.


- 如果没有足够的背景知识，这个故事就会失去深度和背景。
- 如果没有纠正未完成的行动项目的计划，故事就会失去解决方案。
- 没有时间表来说明发生了什么，故事就失去了情节。


Make sure that your retrospectives have all the necessary parts to create a compelling and helpful narrative.

确保您的回顾包含所有必要的部分，以创建引人入胜且有用的叙述。

In Travis CI’s retrospective on [high queue times on OSX builds](https://www.traviscistatus.com/incidents/khzk8bg4p9sy), the author begins by giving an overview of the incident. Next, is the background that explains its relevance to the incident. It states, “Understanding this separation of the creation/build run and the cleanup parts of the life-cycle becomes important in understanding what contributed to this incident.”

在 Travis CI 对 [OSX 构建的高排队时间](https://www.traviscistatus.com/incidents/khzk8bg4p9sy) 的回顾中，作者首先概述了该事件。接下来，是解释其与事件相关性的背景。它指出，“理解生命周期的创建/构建运行和清理部分的这种分离对于理解导致此事件的原因很重要。”

After the background, we get into the incident itself. The author walks us step by step through what happened, using timestamps to show us the duration. After sharing how the team mitigated the incident, the author explains what they intend to do going forward. They list three main objectives to strengthen infrastructure.

在背景之后，我们进入事件本身。作者一步一步地引导我们了解所发生的事情，使用时间戳向我们展示持续时间。在分享了团队如何缓解这一事件之后，作者解释了他们接下来打算做什么。他们列出了加强基础设施的三个主要目标。

The story closes with an excellent, blameless summary. “We always use problems like these as an opportunity for us to improve, and this will be no exception.”

这个故事以一个出色的、无可指责的总结结束。 “我们总是把这样的问题作为我们改进的机会，这次也不例外。”

By learning from example and applying it to your organizational context, your team can write better retrospectives. Retrospectives shouldn’t aren't only a checkbox item. They're a way to catalyze introspection and action to prevent further incidents. Again, there’s no one size fits all, but your team can apply any one (or all) of the above practices starting today.

通过从示例中学习并将其应用于您的组织环境，您的团队可以编写更好的回顾。回顾不应该只是一个复选框项目。它们是一种促进内省和采取行动以防止进一步事件发生的方式。同样，没有一刀切，但您的团队可以从今天开始应用上述任何一种（或全部）实践。

If you want more reading, check out

如果您想阅读更多内容，请查看

- This [example postmortem](https://landing.google.com/sre/sre-book/chapters/postmortem/) from Google
- [Building Reliability Through Culture with Veteran Google SRE, Steve McGhee](https://www.blameless.com/building-reliability-through-culture-sre-steve-mcghee/) 

- 这个[事后分析示例](https://landing.google.com/sre/sre-book/chapters/postmortem/) 来自 Google
- [与资深 Google SRE Steve McGhee 通过文化建立可靠性](https://www.blameless.com/building-reliability-through-culture-sre-steve-mcghee/)

