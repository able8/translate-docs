# Contributing to the Development Guide

# 为开发指南做贡献

A new contributor describes the experience of writing and submitting changes to the Kubernetes Development Guide.

一位新的贡献者描述了编写和提交 Kubernetes 开发指南更改的经验。

By **Erik L. Arneson** \|
Monday, September 28, 2020

作者：**Erik L. Arneson** \|
2020 年 9 月 28 日，星期一

When most people think of contributing to an open source project, I suspect they probably think of
contributing code changes, new features, and bug fixes. As a software engineer and a long-time open
source user and contributor, that’s certainly what I thought. Although I have written a good quantity
of documentation in different workflows, the massive size of the Kubernetes community was a new kind
of “client.” I just didn’t know what to expect when Google asked my compatriots and me at
[Lion’s Way](https://lionswaycontent.com/) to make much-needed updates to the Kubernetes Development Guide.

当大多数人想到为开源项目做贡献时，我怀疑他们可能会想到
贡献代码更改、新功能和错误修复。作为一名软件工程师和长期开放的
源用户和贡献者，这当然是我的想法。虽然我已经写了很多
不同工作流中的文档，Kubernetes 社区的庞大规模是一种新类型
的“客户”。当 Google 询问我和我的同胞时，我只是不知道会发生什么
[Lion's Way](https://lionswaycontent.com/) 对 Kubernetes 开发指南进行急需的更新。

## The Delights of Working With a Community

## 与社区合作的乐趣

As professional writers, we are used to being hired to write very specific pieces. We specialize in
marketing, training, and documentation for technical services and products, which can range anywhere from relatively fluffy marketing emails to deeply technical white papers targeted at IT and developers. With
this kind of professional service, every deliverable tends to have a measurable return on investment.
I knew this metric wouldn’t be present when working on open source documentation, but I couldn’t
predict how it would change my relationship with the project.

作为专业作家，我们习惯于受雇撰写非常具体的作品。我们专注于
技术服务和产品的营销、培训和文档，范围可以从相对蓬松的营销电子邮件到针对 IT 和开发人员的深入技术白皮书。和
这种专业的服务，每一个可交付成果往往都有可衡量的投资回报。
我知道在处理开源文档时不会出现这个指标，但我不能
预测它将如何改变我与项目的关系。

One of the primary traits of the relationship between our writing and our traditional clients is that we
always have one or two primary points of contact inside a company. These contacts are responsible
for reviewing our writing and making sure it matches the voice of the company and targets the
audience they’re looking for. It can be stressful – which is why I’m so glad that my writing
partner, eagle-eyed reviewer, and bloodthirsty editor [Joel](https://twitter.com/JoelByronBarker)
handles most of the client contact.

我们的写作与传统客户之间关系的主要特征之一是我们
在公司内部始终拥有一两个主要联系人。这些联系人负责
审查我们的写作并确保它符合公司的声音并针对
他们正在寻找的观众。可能会有压力——这就是为什么我很高兴我的写作
合伙人，鹰眼审稿人，嗜血编辑 [Joel](https://twitter.com/JoelByronBarker)
处理大部分客户联系。

I was surprised and delighted that all of the stress of client contact went out the window when
working with the Kubernetes community.

当客户接触的所有压力消失时，我感到惊讶和高兴
与 Kubernetes 社区合作。

“How delicate do I have to be? What if I screw up? What if I make a developer angry? What if I make
enemies?” These were all questions that raced through my mind and made me feel like I was
approaching a field of eggshells when I first joined the `#sig-contribex` channel on the Kubernetes
Slack and announced that I would be working on the
[Development Guide](https://git.k8s.io/community/contributors/devel/development.md).

“我需要多娇气？如果我搞砸了怎么办？如果我让开发人员生气怎么办？如果我做
敌人？”这些都是我脑海中闪过的问题，让我觉得我是
当我第一次加入 Kubernetes 上的“#sig-contribex”频道时，我正在接近一个蛋壳领域
Slack 并宣布我将致力于
 [开发指南](https://git.k8s.io/community/contributors/devel/development.md)。

![](http://www.kubernetes.dev/blog/2020/09/28/contributing-to-the-development-guide/jorge-castro-code-of-conduct_hu5bc3c30874931ced96ecf71d135c93d2_143155_800x450_fit_q75_catmullrom.jpg)

"The Kubernetes Code of Conduct is in effect, so please be excellent to each other." — Jorge
Castro, SIG ContribEx co-chair

“Kubernetes 行为准则已经生效，所以请彼此善待。” - 乔治
Castro，SIG ContribEx 联合主席

My fears were unfounded. Immediately, I felt welcome. I like to think this isn’t just because I was
working on a much needed task, but rather because the Kubernetes community is filled
with friendly, welcoming people. During the weekly SIG ContribEx meetings, our reports on progress
with the Development Guide were included immediately. In addition, the leader of the meeting would
always stress that the [Kubernetes Code of Conduct](https://www.kubernetes.dev/community/code-of-conduct/) was in
effect, and that we should, like Bill and Ted, be excellent to each other.

我的恐惧是没有根据的。顿时，我感到宾至如归。我喜欢认为这不仅仅是因为我
从事一项急需的任务，而是因为 Kubernetes 社区已满
与友好、热情的人在一起。在每周 SIG ContribEx 会议期间，我们的进度报告
与开发指南被立即包括在内。此外，会议负责人将
始终强调 [Kubernetes 行为准则](https://www.kubernetes.dev/community/code-of-conduct/) 在
效果，而且我们应该像比尔和泰德一样，对彼此都很好。

## This Doesn’t Mean It’s All Easy

## 这并不意味着一切都很简单

The Development Guide needed a pretty serious overhaul. When we got our hands on it, it was already
packed with information and lots of steps for new developers to go through, but it was getting dusty
with age and neglect. Documentation can really require a global look, not just point fixes.
As a result, I ended up submitting a gargantuan pull request to the
[Community repo](https://github.com/kubernetes/community): 267 additions and 88 deletions.

开发指南需要进行非常认真的修改。当我们拿到手时，它已经
为新开发人员提供了大量信息和许多步骤，但它变得尘土飞扬
随着年龄和忽视。文档确实需要全局查看，而不仅仅是点修复。
结果，我最终向
[社区仓库](https://github.com/kubernetes/community)：添加 267 条，删除 88 条。

The life cycle of a pull request requires a certain number of Kubernetes organization members to review and approve changes
before they can be merged. This is a great practice, as it keeps both documentation and code in
pretty good shape, but it can be tough to cajole the right people into taking the time for such a hefty
review. As a result, that massive PR took 26 days from my first submission to final merge. But in
the end, [it was successful](https://github.com/kubernetes/community/pull/5003). 

拉取请求的生命周期需要一定数量的 Kubernetes 组织成员来审查和批准更改
在它们合并之前。这是一个很好的做法，因为它同时保留了文档和代码
身材不错，但要说服合适的人花时间做这么大的事情可能很难
审查。结果，从我第一次提交到最终合并，这个庞大的 PR 花了 26 天。但在
最后，[成功](https://github.com/kubernetes/community/pull/5003)。

Since Kubernetes is a pretty fast-moving project, and since developers typically aren’t really
excited about writing documentation, I also ran into the problem that sometimes, the secret jewels
that describe the workings of a Kubernetes subsystem are buried deep within the [labyrinthine mind of\
a brilliant engineer](https://github.com/amwat), and not in plain English in a Markdown file. I ran headlong into this issue
when it came time to update the getting started documentation for end-to-end (e2e) testing.

由于 Kubernetes 是一个快速发展的项目，而且由于开发人员通常并不是真正的
对编写文档感到兴奋，我也遇到了问题，有时，秘密珠宝
描述 Kubernetes 子系统工作的那些东西被深深地埋在了[迷宫般的头脑中]
一位出色的工程师](https://github.com/amwat)，而不是在 Markdown 文件中使用简单的英语。我一头扎进这个问题
当需要更新端到端 (e2e) 测试的入门文档时。

This portion of my journey took me out of documentation-writing territory and into the role of a
brand new user of some unfinished software. I ended up working with one of the developers of the new
[`kubetest2` framework](https://github.com/kubernetes-sigs/kubetest2) to document the latest process of
getting up-and-running for e2e testing, but it required a lot of head scratching on my part. You can
judge the results for yourself by checking out my
[completed pull request](https://github.com/kubernetes/community/pull/5045).

我旅程的这一部分使我脱离了文档编写领域，并成为了
一些未完成软件的全新用户。我最终与新的开发人员之一合作
[`kubetest2` 框架](https://github.com/kubernetes-sigs/kubetest2) 记录最新进程
开始并运行 e2e 测试，但这需要我费心费力。你可以
通过查看我的结果为自己判断结果
[完成拉取请求](https://github.com/kubernetes/community/pull/5045)。

## Nobody Is the Boss, and Everybody Gives Feedback

## 没有人是老板，每个人都提供反馈

But while I secretly expected chaos, the process of contributing to the Kubernetes Development Guide
and interacting with the amazing Kubernetes community went incredibly smoothly. There was no
contention. I made no enemies. Everybody was incredibly friendly and welcoming. It was _enjoyable_.

但是虽然我暗自预料会出现混乱，但为 Kubernetes 开发指南做出贡献的过程
与令人惊叹的 Kubernetes 社区的互动非常顺利。没有
争执。我没有树敌。每个人都非常友好和热情。这是_令人愉快的_。

With an open source project, there is no one boss. The Kubernetes project, which approaches being
gargantuan, is split into many different special interest groups (SIGs), working groups, and
communities. Each has its own regularly scheduled meetings, assigned duties, and elected
chairpersons. My work intersected with the efforts of both SIG ContribEx (who watch over and seek to
improve the contributor experience) and SIG Testing (who are in charge of testing). Both of these
SIGs proved easy to work with, eager for contributions, and populated with incredibly friendly and
welcoming people.

对于开源项目，没有一个老板。 Kubernetes 项目，它接近于
庞大的，分为许多不同的特殊兴趣小组 (SIG)、工作组和
社区。每个人都有自己定期安排的会议、分配的职责和选举产生的
主席。我的工作与 SIG ContribEx（他们监督并寻求
改善贡献者体验）和 SIG 测试（负责测试）。这两个
事实证明，SIG 易于使用，渴望做出贡献，并且拥有令人难以置信的友好和
欢迎人们。

In an active, living project like Kubernetes, documentation continues to need maintenance, revision,
and testing alongside the code base. The Development Guide will continue to be crucial to onboarding
new contributors to the Kubernetes code base, and as our efforts have shown, it is important that
this guide keeps pace with the evolution of the Kubernetes project.

在像 Kubernetes 这样活跃的、有生命力的项目中，文档继续需要维护、修订、
并与代码库一起进行测试。开发指南将继续对入职至关重要
Kubernetes 代码库的新贡献者，正如我们的努力所表明的那样，重要的是
本指南与 Kubernetes 项目的发展保持同步。

Joel and I really enjoy interacting with the Kubernetes community and contributing to
the Development Guide. I really look forward to continuing to not only contributing more, but to
continuing to build the new friendships I’ve made in this vast open source community over the past
few months.

Joel 和我真的很喜欢与 Kubernetes 社区互动并为
开发指南。我真的很期待继续不仅贡献更多，而且
继续建立我过去在这个庞大的开源社区中建立的新友谊
几个月。

- [← Previous](http://www.kubernetes.dev/blog/2020/08/24/announcing-the-contributor-website/)
[Next →](http://www.kubernetes.dev/blog/2021/05/14/contributor-stories-series/)

- [← 上一页](http://www.kubernetes.dev/blog/2020/08/24/annoucing-the-contributor-website/)
[下一步→](http://www.kubernetes.dev/blog/2021/05/14/contributor-stories-series/)

© 2021 The Kubernetes Authors All Rights Reserved 

© 2021 The Kubernetes Authors 版权所有

