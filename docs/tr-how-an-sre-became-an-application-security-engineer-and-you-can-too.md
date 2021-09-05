# How an SRE became an Application Security Engineer (and you can too)

# SRE 如何成为应用安全工程师（你也可以）

I’ve had an ambition to become a security engineer for some time. I realized I found security really interesting early on in my career as an engineer in late 2015; it was a nice complement to infrastructure and networking, the other interests I quickly picked up. However, security roles are notoriously hard to fill but also notoriously hard to get. I assumed it would be something I’d pursue once I felt more secure in my skills as a site reliability engineer – meaning I might have waited forever.

一段时间以来，我一直渴望成为一名安全工程师。 2015 年底，我意识到在我作为工程师的职业生涯早期，我发现安全非常有趣；这是对基础设施和网络的一个很好的补充，我很快就找到了其他兴趣。然而，众所周知，安全角色很难填补，也很难得到。我认为，一旦我觉得自己作为站点可靠性工程师的技能更加安全，我就会追求它——这意味着我可能会一直等待。

Instead, early last fall, a friend reached out to me about an opening on her team. They wanted someone with an SRE background and security aspirations. Was I interested in pursuing it as a job, she asked, or was it just a professionally adjacent hobby?

取而代之的是，去年秋天早些时候，一位朋友向我询问了她团队的空缺职位。他们想要一个具有 SRE 背景和安全愿望的人。她问，我是否有兴趣将其作为一份工作来追求，还是只是一种与专业相关的爱好？

I had to sit and think about that.

我不得不坐下来想一想。

For about five seconds.

大约五秒钟。

I typed back a quick I VOLUNTEER AS TRIBUTE extremely casual and professional confirmation that I was indeed interested in security as a career, and then the process began in earnest.

我快速打回了一份“I VOLUNTEER AS TRIBUTE”，非常随意和专业地确认我确实对将安全作为职业感兴趣，然后这个过程就开始了。

But before we get to where I am now, let’s back up to how I got here – and what I brought to the interview that (in my opinion, at least) got me the job.

但在我们到达现在的位置之前，让我们回顾一下我是如何来到这里的——以及我在面试中带来的东西（至少在我看来）让我得到了这份工作。

## The earnest scribbler

## 认真的涂鸦者

I’ve taken to describing the last couple of months as the _Slumdog Millionaire_ portion of my career: it’s the part where everything I’ve done suddenly falls into place in a new and rather wonderful way.

我已经开始将过去几个月描述为我职业生涯中_贫民窟的百万富翁_部分：这是我所做的一切突然以一种新的、相当美妙的方式落实到位的部分。

I began my career as a writer and editor. It’s what I went to college for; rather shockingly, it’s how I supported myself for more than a decade. I’ve been a proofreader, a freelance writer, a mediocre book reviewer, a contract content editor, a lead editor in charge of style guides and training groups of other contractors, a mediocre marketing writer, and a content strategist. Toward the end of that time, I got a certificate in user-centered design from the University of Washington. I loved the work and loved content strategy, but it had become increasingly apparent over the previous couple of years that most writing that paid was meant to sell something, which wasn’t what I wanted to spend forty-plus hours a week doing. I could get by, but I knew I needed to change something, so I began paying closer attention to what I was actually really good at and what wasn’t a total slog within my jobs.

我以作家和编辑的身份开始了我的职业生涯。这就是我上大学的目的；令人震惊的是，这就是我十多年来养活自己的方式。我做过校对员、自由作家、平庸的书评人、合同内容编辑、负责其他承包商的风格指南和培训小组的主编、平庸的营销作家和内容策略师。在那段时间快结束的时候，我从华盛顿大学获得了以用户为中心的设计证书。我喜欢这份工作，也喜欢内容策略，但在过去的几年里，越来越明显的是，大多数付费写作都是为了卖东西，这不是我想要每周花四十多个小时做的事情。我可以过得去，但我知道我需要改变一些东西，所以我开始密切关注我真正擅长的事情以及我工作中不完全是麻烦的事情。

I learned to understand things and to be able to teach those same things to other people in language they could understand and refer to over and over. I became a prolific and effective documentation writer. I learned to navigate people and teams, most especially to get buy-in, because as a writer, sometimes half the job is convincing people that writing is worth paying for.

我学会了理解事物，并且能够用他们可以理解和反复引用的语言向其他人传授同样的东西。我成为了一名多产且高效的文档作家。我学会了为人和团队导航，最重要的是获得支持，因为作为一名作家，有时一半的工作是让人们相信写作是值得付出的。

Toward the end of that period of struggle, I hung out with some software engineers for the first time. I discovered – and I say this with infinite gentleness – that they weren’t any smarter than I am. Between that and the rising resources for people looking to get into programming without a computer science degree, I decided it was now or never, and I needed to make the leap. If all else failed, I could go back to being a frustrated writer – just one who at least _tried_ to do something else.

在那段挣扎期即将结束时，我第一次与一些软件工程师一起出去玩。我发现——我以无限温和的语气说——他们并不比我更聪明。在这与希望在没有计算机科学学位的情况下进入编程领域的人们不断增加的资源之间，我决定现在或永远不会，我需要实现飞跃。如果一切都失败了，我可以重新做一个沮丧的作家——只是一个至少_尝试_做其他事情的人。

## Engineering part one: consulting

## 工程第一部分：咨询

I landed in San Francisco in 2015 to go to code school, and I stayed when I got my first job. I worked at a consultancy for three years, doing a lot of work in healthcare and govtech. I was thrown in the deep end as an infrastructure engineer and essentially a sysadmin, which was incredibly difficult and also incredibly formative to the kind of engineer I’d become. I also spent six months as a full-stack engineer. 

我于 2015 年登陆旧金山去上代码学校，当我找到第一份工作时我就留下来了。我在一家咨询公司工作了三年，在医疗保健和政府科技领域做了很多工作。作为一名基础设施工程师和系统管理员，我陷入了深渊，这对我成为的工程师来说非常困难，而且也非常重要。我还花了六个月的时间担任全栈工程师。

Code school taught me how to code, to connect different layers of the stack to each other, and how to begin researching complex problems. At my first job, I added networking and AWS, Terraform, Bash, my love of writing CLI tools, and automation. I learned more about navigating bureaucratic nightmares, how to run teams effectively, and how to facilitate good meetings and retros. I also learned that I don’t like writing Javascript very much. Between that and realizing how much I enjoyed working with AWS, I decided my next role would be as a site reliability engineer or something like it.

代码学校教我如何编码，如何将堆栈的不同层相互连接，以及如何开始研究复杂的问题。在我的第一份工作中，我添加了网络和 AWS、Terraform、Bash、我对编写 CLI 工具的热爱以及自动化。我学到了更多关于如何应对官僚主义的噩梦、如何有效地运营团队以及如何促进良好的会议和回顾。我还了解到我非常不喜欢编写 Javascript。在此期间，意识到我非常喜欢与 AWS 合作，我决定我的下一个角色是站点可靠性工程师或类似的职位。

## Interlude one: in which security becomes very interesting

## 插曲一：其中安全变得非常有趣

I began to think I might have something to offer the security world at the predecessor to [Day of Shecurity](https://www.dayofshecurity.com/) in 2017. I was interested enough in the subject to sign up, but thinking I might have relevant skills was a different matter. Generally, I explained my interest in security as a complement to my regular job. “Ops is about building things,” I’d say. “Security tells me how you can break them, so I can learn to build them better.”

2017年的[Day of Shecurity](https://www.dayofshecurity.com/)的前身，我开始想我可能有一些东西可以提供给安全世界。我对这个主题有足够的兴趣报名，但在想我可能有相关技能是另一回事。一般来说，我将我对安全的兴趣解释为对我的常规工作的补充。 “运营是关于构建事物，”我会说。 “安全性告诉我如何破解它们，这样我才能学会更好地构建它们。”

One session that day was a CTF of sorts, with three flags to find in the vulnerable test network we were exploring. Navigating to them required using command line tools, the ability to grep, and feeling comfortable with flags and documentation. I won two of the three and won two Amazon gift cards. I bought the official Golang book, cat litter, and the sparkly boots I still wear at least a couple of times a week, which I saw on a cool woman in the Lookout office that day. I’ve thought of them as my security boots ever since and have stomped through DEF CON, DoS, and job interviews in them. I use them as a reminder of that feeling: oh, wait, I might actually have something to offer in this field.

那天的一个会议是各种 CTF，在我们正在探索的易受攻击的测试网络中可以找到三个标志。导航到它们需要使用命令行工具、grep 的能力以及对标志和文档感到满意。我赢了三张中的两张，赢了两张亚马逊礼品卡。我买了 Golang 的官方书、猫砂和闪亮的靴子，我仍然每周至少穿几次，那天我在 Lookout 办公室的一个很酷的女人身上看到了这些。从那时起，我就将它们视为我的安全靴，并在其中踩过 DEF CON、DoS 和工作面试。我用它们来提醒那种感觉：哦，等等，我实际上可能在这个领域提供一些东西。

## Engineering part two: SRE life

## 工程第二部分：SRE 生命

I spent 13 months as an SRE, a job I was thrilled to get (and still thrilled I landed). I got to dig deeper into the skills I’d gotten at the last job, as well as spending long days with Elasticsearch, becoming friends with Ansible, and learning another flavor of Linux. My company outsourced their security work to an outside firm, and I made a point of studying what they did: reading the emails they sent back to bug bounty seekers, responding to the small incidents that popped up here and there, and carrying out mitigations of issues in the infrastructure reviews they made for us.

我做了 13 个月的 SRE，这份工作让我很兴奋（并且仍然很高兴我找到了）。我必须更深入地挖掘我在上一份工作中获得的技能，并在 Elasticsearch 度过了漫长的日子，与 Ansible 成为朋友，并学习了另一种风格的 Linux。我的公司将他们的安全工作外包给了一家外部公司，我重点研究了他们的工作：阅读他们发回给漏洞赏金寻求者的电子邮件，对这里和那里突然出现的小事件做出回应，并采取缓解措施他们为我们所做的基础设施审查中的问题。

I also grabbed the security-adjacent opportunities that came up, doing a lot of work on our AWS IAM policies and co-creating the company educational material on phishing avoidance. I learned about secret storage and rotation, artful construction of security groups for our servers, and how to best communicate policies like password manager use to people with lots of different technical backgrounds.

我还抓住了出现的与安全相关的机会，在我们的 AWS IAM 策略方面做了大量工作，并共同创建了公司关于避免网络钓鱼的教育材料。我了解了秘密存储和轮换、为我们的服务器巧妙构建安全组，以及如何最好地向具有许多不同技术背景的人传达密码管理器使用等策略。

## Engineering part three: present day

## 工程第三部分：现在

Early last fall, a friend reached out saying that her team at Salesforce wanted someone with an SRE background and an interest in learning security. We’d gone to the same coding school, though not at the same time. We actually met at a [WISP](https://wisporg.com) event. She placed first in the handcuff escape competition; I placed second. We stayed in touch. She invited me over to a certain very tall San Francisco building to talk to her and her manager about the role, and so the process began. 

去年秋天早些时候，一位朋友伸出手说她在 Salesforce 的团队想要一个具有 SRE 背景并且对学习安全感兴趣的人。我们去了同一所编程学校，尽管不是同时。我们实际上是在 [WISP](https://wisporg.com) 活动中认识的。她在逃脱手铐比赛中获得第一名；我排名第二。我们保持联系。她邀请我到旧金山某栋很高的大楼去和她和她的经理谈谈这个角色，于是这个过程就开始了。

My team does software reviews, which can involve black box pen testing (where we don't see the code), code reviews, consulting on responsible data use for possible software options and the expansion of existing tools we use, and being a resource for other teams. We’re a friendlier face of security, which is the only kind of security I’m really interested in being a part of. We also work directly with outside software companies to improve their security practices if they don’t pass our initial review, so I’ll get the chance to help other engineers be better, which is one of my favorite things.

我的团队进行软件审查，这可能涉及黑盒笔测试（我们看不到代码的地方）、代码审查、对可能的软件选项的负责任数据使用的咨询以及我们使用的现有工具的扩展，以及作为资源其他球队。我们是一个更友好的安全面孔，这是我真正有兴趣参与的唯一一种安全。如果他们没有通过我们的初步审查，我们还直接与外部软件公司合作以改进他们的安全实践，所以我将有机会帮助其他工程师做得更好，这是我最喜欢的事情之一。

As of this writing, I still spend most of my days on training: learning to write and read Apex, doing secure code reviews of increasing complexity, and figuring out who does what in a security org with more than a thousand people. Coming into a very large company requires a ton of building context, and fortunately, I get the space to figure it all out.

在撰写本文时，我的大部分时间仍然花在培训上：学习编写和阅读 Apex，进行日益复杂的安全代码审查，并弄清楚谁在拥有一千多人的安全组织中做什么。进入一家非常大的公司需要大量的建筑背景，幸运的是，我有足够的空间来弄清楚这一切。

## Skills, revisited

## 技能，重温

So now you have an idea of what I learned and brought to the process of applying to this job. I recognized fairly quickly before that ops and security have a lot of things in common – that is, beyond a reputation for being risk-averse and more than a little curmudgeonly.

所以现在您对我所学到的知识以及在申请这份工作的过程中所学到的有所了解。在此之前，我很快就意识到运维和安全有很多共同点——也就是说，除了以规避风险而闻名之外，还有一点脾气暴躁。

There are skills that are essential for both, including:

两者都需要一些技能，包括：

- Networking, AWS, and port hygiene
- Coding, especially scripting; Bash and Python are great choices
- Command line abilities
- Looking up running processes and changing their state
- Reading and manipulating logs

- 网络、AWS 和端口卫生
- 编码，尤其是脚本； Bash 和 Python 是不错的选择
- 命令行能力
- 查找正在运行的进程并更改它们的状态
- 读取和操作日志

The skills that less explicitly in demand but that I’ve found to be really useful include:

需求不那么明确但我发现非常有用的技能包括：

- Communication, both written and verbal
- Documentation creation and maintenance
- Teaching
- A UX-centered approach

- 沟通，包括书面和口头
- 文档创建和维护
- 教学
- 以用户体验为中心的方法

Let me explain what I mean by that last one. As I said before, I have some education in UX principles in practices, and I’ve done official UX exercises as part of jobs. I’m still able to, if needed. The part of it I use most often, though, is what I’ve come to think of as a UX approach to the everyday.

让我解释一下我说的最后一个是什么意思。正如我之前所说，我在实践中接受了一些 UX 原则方面的教育，并且作为工作的一部分，我已经完成了官方的 UX 练习。如果需要，我仍然可以。不过，我最常使用的部分是我认为的日常 UX 方法。

What I mean by that is the ability to come into a situation with someone and assume that you don’t understand their motivations, previous actions, or context, and then to work deliberately to build those by asking questions. The key part is remembering that, even if someone is doing something you don’t think makes sense, they most likely have reasons for it, and you can only discover those by asking them.

我的意思是，有能力与某人一起进入某种情境，并假设您不了解他们的动机、以前的行为或背景，然后通过提问来刻意地建立这些。关键是要记住，即使有人在做你认为不合理的事情，他们很可能有理由这样做，而你只能通过询问他们来发现这些。

This is at the center of how I approach all of my work, and it seems to be distinctive – when I left my last job, a senior engineer pulled me aside and gave me the nicest compliment about how he'd learned from me by watching me do exactly that approach for the year we'd worked together. He told me how different it was from how he worked and that he’d learned from me. It was a very nice sendoff.

这是我处理所有工作的方式的中心，它似乎很独特——当我离开上一份工作时，一位高级工程师把我拉到一边，并称赞我他是如何通过观看向我学习的在我们一起工作的那一年，我就是这样做的。他告诉我这与他的工作方式有多么不同，而且他是从我那里学到的。这是一个非常好的送别。

## Interlude two: my accidental security education

## 插曲二：我的意外安全教育

Here’s something I only realized afterward, which I alluded to earlier: I’ve done a LOT of security learning since becoming an engineer. I just didn’t fully realize what I’d been doing because I thought I was just having fun.

这是我后来才意识到的事情，我之前提到过：自从成为工程师以来，我已经进行了很多安全学习。我只是没有完全意识到我在做什么，因为我以为我只是在玩。

So I did none of these things with interview preparation in mind. The closest I came was thinking, “Oh, I see how this might be useful for the kinds of jobs I might want later, but I’m definitely not pursuing that job right now.” Well! Maybe you can be more deliberate and aware than I was.

所以我做这些事情时都没有考虑到面试准备。最接近我的想法是，“哦，我明白这对我以后可能想要的工作类型有什么用，但我现在绝对不会追求那份工作。”好！也许你可以比我更深思熟虑。

These are the things I did that ended up being really helpful, when it came to prepare officially for a security interview, over the last four years:

在过去的四年里，在正式准备安全面试时，这些是我所做的最终非常有帮助的事情：

- Going to DEF CON four times
- Going to Day of Shecurity three times
- Being a beta student for a[friend’s security education startup](https://www.goldhatsecurity.com/) for an eight-part course all about writing secure code
- Attempting CTFs (though I’m still not super proficient at this yet)
- Talking security with my ops coworkers, who all have opinions and stories
- Volunteering for AWS IAM work whenever it came up as a task
- Classes at the[Bradfield School of Computer Science](https://bradfieldcs.com/) in computer architecture and networking (try to get a company to pay for this) 

- 四次参加 DEF CON
- 三次去Shecurity日
- 成为 [朋友的安全教育初创公司](https://www.goldhatsecurity.com/) 的 Beta 学生，学习有关编写安全代码的八部分课程
- 尝试 CTF（虽然我还不是很精通）
- 与我的运维同事讨论安全问题，他们都有意见和故事
- 自愿为 AWS IAM 工作，每当它作为一项任务出现时
- 在[布拉德菲尔德计算机科学学院](https://bradfieldcs.com/)的计算机体系结构和网络课程（尝试让公司为此付费)

Every one of these things gave me something that either helped me feel more adept while interviewing or something I mentioned specifically when discussing things and answering questions. Four years is a lot of time to pursue something casually, especially since I usually went to an event every month or two.

这些东西中的每一件都给了我一些东西，让我在面试时感觉更熟练，或者在讨论事情和回答问题时特别提到的东西。四年时间很长，可以随便追求一些东西，尤其是因为我通常每两个月都会参加一次活动。

I’ve also benefited a lot from different industry newsletters, especially these:

我也从不同的行业通讯中受益匪浅，尤其是这些：

- [The Cloud Security Reading List](https://cloudseclist.com/)
- [Julia Evans’s weekly emailed engineering comics](https://jvns.ca/newsletter/)
- [Devops Weekly](https://www.devopsweekly.com/)
- [SRE Weekly](https://sreweekly.com/)
- [Crypto-Gram from Bruce Schneier](https://www.schneier.com/crypto-gram/subscribe.html)

- [云安全阅读清单](https://cloudseclist.com/)
- [Julia Evans 每周通过电子邮件发送的工程漫画](https://jvns.ca/newsletter/)
- [Devops Weekly](https://www.devopsweekly.com/)
- [SRE Weekly](https://sreweekly.com/)
- [来自布鲁斯·施奈尔的 Crypto-Gram](https://www.schneier.com/crypto-gram/subscribe.html)

Many of these are ops-centric, but all of them have provided something as I was working toward shifting jobs. Very few issues and problems exist in only a single discipline, and these digests have been really useful for seeing the regular intersections between things I knew and things I wanted to know more about.

其中许多都是以运营为中心的，但在我努力换工作时，所有这些都提供了一些东西。很少有问题只存在于一个学科中，这些摘要对于查看我所知道的事物和我想了解的事物之间的常规交叉点非常有用。

## Interview preparation, done deliberately

## 面试准备，刻意做好

I officially applied for the job a month or so after that fateful informational coffee. I applied while I was out of town for three weeks being a maid of honor in my best friend’s wedding, meaning I didn’t get to do much until I was home and had slept for a couple of days.

在那次决定性的信息咖啡之后一个月左右，我正式申请了这份工作。我在出城三周的时候申请成为我最好朋友婚礼的伴娘，这意味着我直到回家睡了几天才能做很多事情。

Once my brain worked again, I made a wishlist of everything I wanted to be able to talk confidently about. Then I prioritized it. Then I began working through everything I could. I touched on about half of it.

一旦我的大脑再次运转起来，我就列出了我想要能够自信地谈论的所有事情的愿望清单。然后我优先考虑它。然后我开始尽我所能。我触及了大约一半。

I studied for about a week and a half, a couple hours at a time. I focused on three main things:

我学习了大约一个半星期，一次几个小时。我主要关注三件事：

- [Exercism](https://exercism.io/), primarily in Python
- The OWASP top ten from 2013 and 2017
- Blog posts that crossed my current discipline and the one I aspired to

- [Exercism](https://exercism.io/)，主要使用 Python
- 2013 年和 2017 年 OWASP 前十名
- 跨越我目前的学科和我渴望的学科的博客文章

The Exercism work was because I never feel like I code as much as I’d like in my jobs, and I feel more confident in technical settings when I feel more fluent in code. The OWASP reading was a mix of official resources, their cheat sheets, and other people’s writing about them; reading different perspectives is part of how I wrap my head around things like this. And the blog posts were for broader context and also to get more conversant about the intersection between my existing skills and the role I was aspiring to. The Capital One breach was really useful for this, because it happened due to misconfigured AWS IAM permissions.

Exercism 的工作是因为我从来没有觉得我的代码像我在工作中想要的那样多，而且当我对代码更流利时，我对技术设置更有信心。 OWASP 阅读材料混合了官方资源、他们的备忘单和其他人关于他们的文章；阅读不同的观点是我如何围绕这样的事情进行思考的一部分。博客文章是为了更广泛的背景，也是为了更熟悉我现有的技能和我渴望担任的角色之间的交叉点。 Capital One 漏洞对此非常有用，因为它是由于 AWS IAM 权限配置错误而发生的。

This is the list I made, ordered by priority. The ones in italics are the ones I addressed to my satisfaction.

这是我制作的列表，按优先级排序。斜体字是我写的令我满意的那些。

- _Python [Exercism](https://exercism.io/)(80%)_
- Dash of Bash Exercism (20%)
- Practice using ops-related Python libraries (request, others???)
- Get a handle on ten core automation-related bash commands
- Bash loops practice
- _DNS, record types_
- _Hack this Site or something similar for pen testing_
- _Read up on Linux privilege escalation_
- _OWASP reading_
- _DNS tunneling_
- _Read over notes from the Day of Shecurity 2019 threat modeling workshop_
- [Katie Murphy’s blog](https://localhost.network/)
- flAWS s3 thing
- Jenkins security issues
- _CircleCI breach_
- Common CI security issues
- Common AWS security issues
- Hacker 101
- _Something something appsec resource_
- Infrastructure principles blog posts
- Security exploits for DNS TXT records

- _Python [运动](https://exercism.io/)(80%)_
- Bash 练习冲刺 (20%)
- 练习使用与 ops 相关的 Python 库（请求，其他？？？）
- 掌握十个与自动化相关的核心 bash 命令
- Bash 循环练习
- _DNS，记录类型_
- _黑掉这个网站或类似的东西以进行渗透测试_
- _阅读 Linux 权限提升_
- _OWASP阅读_
- _DNS 隧道_
- _阅读来自 Shecurity 2019 年威胁建模研讨会的笔记_
- [凯蒂墨菲的博客](https://localhost.network/)
- 错误的 s3 东西
- 詹金斯安全问题
- _CircleCI 违规_
- 常见的 CI 安全问题
- 常见的 AWS 安全问题
- 黑客 101
- _东西 appsec 资源_
- 基础设施原则博客文章
- DNS TXT 记录的安全漏洞

And here, with dates and links, is exactly what I did to study in the week and a half leading up to the interview.

在这里，有日期和链接，正是我在面试前一周半的时间里所做的研究。

**28 October**

**10 月 28 日**

[Cracking Websites with Cross Site Scripting – Computerphile](https://www.youtube.com/watch?v=L5l9lSnNMxg)

[使用跨站脚本破解网站 - Computerphile](https://www.youtube.com/watch?v=L5l9lSnNMxg)

[Hacking Websites with SQL Injection – Computerphile](https://www.youtube.com/watch?v=_jKylhJtPmI)

[使用 SQL 注入攻击网站 - Computerphile](https://www.youtube.com/watch?v=_jKylhJtPmI)

2.5 easy Exercism Python problems

2.5 简单的Exercism Python 问题

**30 October**

**10 月 30 日**

Two easy Exercism Python problems

两个简单的 Exercism Python 问题

[Security Incident on 8/31/2019 – Details and FAQs](https://support.circleci.com/hc/en-us/articles/360034852194-Security-Incident-on-8-31-2019-Details-and-FAQs)

[2019 年 8 月 31 日的安全事件 – 详细信息和常见问题解答](https://support.circleci.com/hc/en-us/articles/360034852194-Security-Incident-on-8-31-2019-Details-and-常见问题)

Three [Hack This Site](https://www.hackthissite.org/) exercises

三 [Hack This Site](https://www.hackthissite.org/) 练习

**31 October**

**10 月 31 日**

[DNS Tunneling: how DNS can be (ab)used by malicious actors](https://unit42.paloaltonetworks.com/dns-tunneling-how-dns-can-be-abused-by-malicious-actors/)

[DNS 隧道：恶意行为者如何 (ab) 使用 DNS](https://unit42.paloaltonetworks.com/dns-tunneling-how-dns-can-be-abused-by-malicious-actors/)

Two easy Exercism problems

两个简单的锻炼问题

**3 November**

**11 月 3 日**

[How NOT to Store Passwords! – Computerphile](https://www.youtube.com/watch?v=8ZtInClXe1Q)

[如何不存储密码！ – Computerphile](https://www.youtube.com/watch?v=8ZtInClXe1Q)

Socket coding in Python with a friend

与朋友在 Python 中进行套接字编码

**4 November** 

**11 月 4 日**

[A Technical Analysis of the Capital One Hack](https://blog.cloudsploit.com/a-technical-analysis-of-the-capital-one-hack-a9b43d7c8aea)

[Capital One Hack 技术分析](https://blog.cloudsploit.com/a-technical-analysis-of-the-capital-one-hack-a9b43d7c8aea)

[How GCHQ Classifies Computer Security – Computerphile](https://www.youtube.com/watch?v=iesgXoOBLZM)

[GCHQ 如何分类计算机安全 – Computerphile](https://www.youtube.com/watch?v=iesgXoOBLZM)

[Basic Linux Privilege Escalation](https://blog.g0tmi1k.com/2011/08/basic-linux-privilege-escalation/)

 [基本Linux提权](https://blog.g0tmi1k.com/2011/08/basic-linux-privilege-escalation/)

Two easy Exercism problems

两个简单的锻炼问题

**5 November**

**11 月 5 日**

1.5 Exercisms

1.5 练习

[The Book of Secret Knowledge](https://github.com/trimstray/the-book-of-secret-knowledge)

 [秘识之书](https://github.com/trimstray/the-book-of-secret-knowledge)

Read about [Scapy](https://scapy.net/) for Python

阅读有关 Python 的 [Scapy](https://scapy.net/)

**6 November**

**11 月 6 日**

Read [OWASP stuff](https://owasp.org/www-project-top-ten/) and made notes, including the [2017 writeup](https://www.owasp.org/images/7/72/OWASP_Top_10-2017_%28en%29.pdf.pdf)

阅读 [OWASP 内容](https://owasp.org/www-project-top-ten/) 并做笔记，包括 [2017 writeup](https://www.owasp.org/images/7/72/OWASP_Top_10-2017_%28en%29.pdf.pdf)

[Bash For Loop Examples](https://www.cyberciti.biz/faq/bash-for-loop/)

[Bash For 循环示例](https://www.cyberciti.biz/faq/bash-for-loop/)

[Every Linux Geek Needs To Know Sed and Awk. Here’s Why…](https://www.makeuseof.com/tag/sed-awk-learn/)

[每个 Linux Geek 都需要了解 Sed 和 Awk。这就是为什么...](https://www.makeuseof.com/tag/sed-awk-learn/)

**7 November**

**11 月 7 日**

An easy Exercism

简单的锻炼

Recited OWASP stuff to Sean

向肖恩背诵 OWASP 的内容

Sean is my boyfriend. One of the kindest things he does for me is that he lets me explain technical things to him until I’m able to explain them to non-engineers again. I do this pretty regularly, because it's really important to me to be able to teach people without a lengthy engineering background, and I did it during interview preparation because I know how easy it is to obscure a lack of understanding with jargon, and I didn 't want to do that. Having someone who lets me do this is perhaps the other thing I didn’t realize would be as helpful as it has been; we started doing it because he wanted to know what I did at work, and I realized that it helped make me a better communicator and engineer. May you all have someone as patient as he is to help you translate engineerspeak to human language on the regular.

肖恩是我的男朋友。他为我做的最友善的事情之一就是让我向他解释技术问题，直到我能够再次向非工程师解释它们为止。我经常这样做，因为对我来说能够教没有冗长的工程背景的人真的很重要，而且我在面试准备期间这样做了，因为我知道用行话来掩盖缺乏理解是多么容易，而且我没有不想那样做。让我做这件事的人也许是我没有意识到的另一件事会如此有帮助；我们开始这样做是因为他想知道我在工作中做了什么，我意识到这有助于让我成为更好的沟通者和工程师。愿你们都有一个像他一样有耐心的人，帮助您定期将工程师的语言翻译成人类语言。

So that was how I spent my preparation time. Next: the interview.

所以这就是我如何度过我的准备时间。下一个：面试。

## A series of conversations, across from the tallest tower

## 最高塔对面的一系列对话

For reasons I’m sure you can guess, I can’t give you the most specific play-by-play of the interview process. However, I got permission to give you a higher-level view of it that I hope will still be illuminating.

出于我相信你能猜到的原因，我不能给你最具体的面试过程。但是，我得到了许可，可以让您对它有一个更高层次的看法，我希望它仍然具有启发性。

My interview was a bit bespoke, because they were more accustomed to hiring people who had already been pen testers or security researchers. Because of that, in addition to proving that I knew a few things about spotting insecure code and thinking through vulnerabilities, I also talked to their DevOps architect about ops things, including opinions on infrastructure as code and the creation and socialization of development environments. (We also found that we take a similarly dim view of senior engineers who bully junior engineers.) I talked about securing a server when several different types of users would need to reach it in different ways. And yes, I talked some about the OWASP top ten.

我的面试有点定制，因为他们更习惯于雇用已经做过渗透测试员或安全研究人员的人。因此，除了证明我知道一些关于发现不安全代码和思考漏洞的知识之外，我还与他们的 DevOps 架构师讨论了运维方面的事情，包括对基础设施即代码的看法以及开发环境的创建和社会化。 （我们还发现，我们对欺负初级工程师的高级工程师也持同样的看法。）我谈到了在几种不同类型的用户需要以不同方式访问服务器时保护服务器的安全。是的，我谈到了 OWASP 前十名。

My bar for a “good interview” is whether the things we talked about or did were directly relevant to the needs and responsibilities of the job, and that was absolutely the case here. The only whiteboarding I did was when I volunteered to do so, drawing out network diagrams when I realized my hand gestures were not up to conveying the complexity of what we were discussing. Everything else felt collaborative, casual, and built to help me explain the things I knew about without feeling all the uncertainty that badly designed interviews can evoke.

我对“好的面试”的标准是我们谈论或做的事情是否与工作的需求和职责直接相关，这里绝对是这种情况。我做的唯一白板是当我自愿这样做时，当我意识到我的手势不能传达我们正在讨论的复杂性时绘制网络图。其他一切都让人感觉协作、随意，并且旨在帮助我解释我所知道的事情，而不会感受到设计糟糕的采访可能引起的所有不确定性。

## Getting ready for your own security path

## 为您自己的安全路径做好准备

My goal in writing this post (based on a talk I did for Secure Diversity on 28 January 2020, which I will link to when the video is up) was to give the extremely specific information about how I got the job that I've always been thirsty for but often found lacking in “how I got here” talks for these kinds of roles. I hope I managed that; when I proposed the talk, I was very grateful to my past self for keeping such fastidious notes.

我写这篇文章的目标（基于我在 2020 年 1 月 28 日为 Secure Diversity 所做的演讲，我将在视频播放时链接到该演讲）是提供有关我如何得到我一直以来的工作的极其具体的信息一直渴望但经常发现缺乏关于这些角色的“我是如何来到这里的”谈话。我希望我做到了；当我提出这个演讲时，我非常感谢过去的自己留下了如此挑剔的笔记。

However, I also want to leave you with some more general ideas of how to shape your current career to more effectively get to the security role I presume you’re seeking. 

但是，我还想为您提供一些关于如何塑造您当前的职业以更有效地担任我认为您正在寻求的安全角色的更一般的想法。

Find a couple security-essential skills you already know something about and dive deeply into them. I have a lot to say about IAM stuff, in AWS and Jenkins and general principle of least privilege stuff, so that’s been something I’ve really focused on when trying to convey my skills to other people. Find what you’re doing that already applies to the role you want, and get conversational. Keep up on news stories relevant to those skills. This part shouldn’t be that hard, because these skills should be interesting to you. If they aren’t, choose different skills to focus on.

找到一些您已经了解的安全基本技能，并深入研究它们。我有很多关于 IAM 的东西，在 AWS 和 Jenkins 以及最小特权的一般原则，所以当我试图向其他人传达我的技能时，这是我真正关注的事情。找到你正在做的已经适用于你想要的角色的事情，并进行对话。了解与这些技能相关的新闻报道。这部分不应该那么难，因为这些技能对你来说应该很有趣。如果不是，请选择不同的技能来关注。

While you’re doing this learning, make sure the people in your professional life know what you’re doing. This can be your manager, but it can also be online communities, coworkers you keep in touch with as you all move companies, and anyone else you can speak computer or security with. Don’t labor in obscurity; share links, mention things you’ve learned, and throw bait out to find other people interested in the same things.

当你在学习时，确保你职业生涯中的人知道你在做什么。这可以是您的经理，但也可以是在线社区、您在搬家时保持联系的同事，以及您可以与之交谈计算机或安全的任何其他人。不要在默默无闻中工作；分享链接，提及你学到的东西，并抛出诱饵寻找其他对相同事物感兴趣的人。

Build that community further by going to meetups and workshops. When I think about living outside the Bay Area (which of course I do, because it's a beloved hobby of just about everyone who lives around here), one of the things that would be hardest to give up is all the free education that's available almost every night of the week. [Day of Shecurity](https://www.dayofshecurity.com/), Secure Diversity, [OWASP](https://www.meetup.com/Bay-Area-OWASP/) in SF and the south bay, [ NCC meetups](https://www.meetup.com/NCCOpenForumSF/), and there are so many more. Go to the thing, learn the thing, and read about the thing after.

通过参加聚会和研讨会进一步建立该社区。当我想到住在湾区以外时（我当然会这样做，因为几乎每个住在这里的人都喜欢它），最难放弃的一件事就是几乎可以提供的所有免费教育一周的每个晚上。 [Shecurity 日](https://www.dayofshecurity.com/)，SecureDiversity，[OWASP](https://www.meetup.com/Bay-Area-OWASP/) 在旧金山和南湾，[ NCC 聚会](https://www.meetup.com/NCCOpenForumSF/)，还有很多。去事物，学习事物，然后阅读事物。

Finally, remember that security needs you. Like all of tech, security is better when there are a lot of different kinds of people working out how to make things and fix things. Please hang in there and keep trying.

最后，请记住安全需要您。像所有技术一样，当有很多不同类型的人在研究如何制造和修复事物时，安全性会更好。请坚持下去并继续努力。

And good luck. <3

还有祝你好运。 <3

Posted on [January 27, 2020July 24, 2021](https://breanneboland.com/blog/2020/01/27/how-an-sre-became-an-application-security-engineer-and-you-can-too/) 

发表于 [2020 年 1 月 27 日 2021 年 7 月 24 日](https://breanneboland.com/blog/2020/01/27/how-an-sre-became-an-application-security-engineer-and-you-can-也/)

