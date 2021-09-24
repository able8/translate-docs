# Incident Management Goes to the Olympics

# 事件管理进入奥运会

A look at outages and disruptions to the IT systems that power the Olympics, from 1996 to today.

看看从 1996 年到今天，为奥运会提供动力的 IT 系统的中断和中断。

Quentin Rousseau

昆汀·卢梭

August 05 2021 · 5 min read

2021 年 8 月 5 日 · 5 分钟阅读

* * *

* * *

![](https://rootly.io/rails/active_storage/blobs/redirect/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBBdWtPIiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--918010be537329a1f4a366b8bda030dc0f866f4c/2020-04-03-media-thumbnail.jpeg)

A lot of things can go wrong during the Olympics. [Broken legs](https://www.usatoday.com/story/sports/columnist/nancy-armour/2021/07/24/french-gymnast-samir-ait-said-broke-leg-rio-olympics-power/8080783002/), [food poisoning](https://www.neogen.com/neocenter/blog/foodborne-illness-spreads-at-site-of-winter-olympics/) and, of course, pandemics can throw a wrench in the years of careful planning that athletes and organizers put into the Games.

奥运会期间很多事情都可能出错。 [断腿](https://www.usatoday.com/story/sports/columnist/nancy-armour/2021/07/24/french-gymnast-samir-ait-said-broke-leg-rio-olympics-power/8080783002/)、[食物中毒](https://www.neogen.com/neocenter/blog/foodborne-illness-spreads-at-site-of-winter-olympics/)，当然，流行病会引发多年来，运动员和组织者对奥运会进行了精心策划。

Here’s another common, but often-overlooked, source of disruption at the Olympics: IT failures. Disruptions to the IT infrastructure that powers the Olympics and makes them viewable by audiences across the globe are more frequent than you may think. It’s only thanks to the work of world-class SREs that these problems are remediated before they exert a serious impact on spectators and athletes.

这是奥运会上另一个常见但经常被忽视的干扰源：IT 故障。为奥运会提供动力并使全球观众都能看到的 IT 基础设施中断比您想象的要频繁。多亏了世界级 SRE 的工作，这些问题才得以解决，才对观众和运动员造成严重影响。

## Digitizing the Olympics: Lessons from 1996

## 奥运会数字化：1996 年的教训

The 1996 Olympic Games in Atlanta occurred relatively early in the history of modern computing. No one at the time had heard of a smartphone, and email remained a novelty for many folks.

1996 年亚特兰大奥运会在现代计算史上发生得相对较早。当时没有人听说过智能手机，而电子邮件对许多人来说仍然是新鲜事物。

Nonetheless, the Olympic Committee and partner businesses seized the event as an opportunity to highlight the new opportunities afforded by digital technology. Organizers “promised the most technologically sophisticated Games ever,” [according to the New York Times](https://www.nytimes.com/1996/07/22/business/olympics-stung-by-technology-s-false-starts.html).

尽管如此，奥委会和合作伙伴企业还是抓住了这次活动的机会，以强调数字技术提供的新机遇。 [据纽约时报报道](https://www.nytimes.com/1996/07/22/business/olympics-stung-by-technology-s-false-开始.html)。

Unfortunately, reality didn’t fully live up to the promises. Phone systems sporadically failed, broadcast streams were disrupted and in at least one case, the flashy electronic system that organizers had built for recording event results registered inaccurate scores.

不幸的是，现实并没有完全兑现承诺。电话系统偶尔出现故障，广播流中断，至少在一种情况下，组织者为记录活动结果而建立的华丽电子系统记录了不准确的分数。

There was also a temporary blackout inside one of the Olympic stadiums, although “the problem was caused by a technician who pulled the wrong switch” rather than an IT failure, according to the Times.

据“泰晤士报”报道，其中一个奥林匹克体育场内也出现了临时停电，尽管“问题是由一名技术人员拉错了开关”而不是 IT 故障引起的。

Ultimately, none of these incidents turned into show-stopping disruptions. But they did give Olympic organizers their first hands-on experience with the types of problems that IT teams need to manage to deliver a digital event of global proportions, paving the way for fully digitized games in the new millennium.

最终，这些事件都没有变成引人注目的中断。但他们确实让奥运会组织者第一次亲身体验了 IT 团队在交付全球规模的数字赛事时需要设法解决的各种问题，为新千年的完全数字化比赛铺平了道路。

## Spotty service in Athens

## 雅典的服务参差不齐

IT systems had been much improved, but not perfected, by the time of the 2004 Olympics in Athens.

到 2004 年雅典奥运会时，IT 系统已经有了很大改进，但还没有完善。

While there were no reports of major Internet-related outages for those Games, which occurred at a time when the was becoming a major sports viewing platform, athletes and spectators did [report serious issues with phone service](https://siouxcityjournal.com/sports/olympics/as-thousands-arrive-athens-wrestles-to-avoid-phone-outages-communications-blackouts/article_3df2fecb-3ca8-5129-b5e3-245c0fdbf173.html). Some attendees were unable to make calls for periods of as long as ten hours, according to media reports from the time.

虽然没有关于这些奥运会与互联网相关的重大中断的报告，这发生在它成为主要体育观看平台的时候，运动员和观众确实[报告了电话服务的严重问题](https://siouxcityjournal.com/sports/olympics/as-thousands-arrive-athens-wrestles-to-avoid-phone-outages-communication-blackouts/article_3df2fecb-3ca8-5129-b5e3-245c0fdbf173.html)。据当时的媒体报道，一些与会者无法拨打电话长达十个小时。

The problem seemed to stem from simple exhaustion of local phone infrastructure. Although Greece’s telecommunications provider had invested in significant infrastructure expansions in anticipation of the Games, they didn’t turn out to be capable of handling all of the demand.

问题似乎源于本地电话基础设施的简单耗尽。尽管希腊的电信供应商在奥运会前投资了大量的基础设施扩建，但结果证明他们并不能满足所有的需求。

The takeaway here for SREs is straightforward enough: When performing capacity planning, assume the worst. Design systems to handle twice as much demand as you actually expect, and plan for some of your infrastructure to go offline occasionally.

SRE 的要点非常简单：在执行容量规划时，假设最坏的情况。设计系统以处理两倍于您实际预期的需求，并计划您的某些基础设施偶尔会脱机。

## Malware strikes in South Korea

## 恶意软件在韩国罢工

If you know anything about cyberattacks, it’s that they have become steadily more common and disruptive over the past decade. 

如果您对网络攻击一无所知，那就是它们在过去十年中变得越来越普遍和具有破坏性。

That fact was reflected in the 2018 Winter Olympics in South Korea, where a [malware attack brought core IT systems offline](https://www.wired.com/story/untold-story-2018-olympics-destroyer-cyberattack/) right in the middle of the opening ceremony. The Games’s website went offline, Internet broadcasts were disrupted and some spectators were not able to attend the ceremony because they couldn’t print their tickets.

这一事实反映在 2018 年韩国冬季奥运会上，[恶意软件攻击使核心 IT 系统脱机](https://www.wired.com/story/untold-story-2018-olympics-destroyer-cyberattack/)就在开幕式的中间。奥运会网站下线，互联网广播中断，一些观众因无法打印门票而无法参加仪式。

To their credit, the engineers overseeing the Games, who had run fire drills to prepare for cyberattacks in the lead-up to the event, resolved service in a matter of hours. They also prevented the incident from escalating into a power outage, which was [reportedly the goal of the attackers](https://www.nytimes.com/2018/02/12/technology/winter-olympic-games-hack.html).

值得称赞的是，负责监督奥运会的工程师在赛事前进行了消防演习以准备网络攻击，在几个小时内解决了服务问题。他们还阻止了事件升级为停电，这是 [据报道攻击者的目标](https://www.nytimes.com/2018/02/12/technology/winter-olympic-games-hack.html)。

The Olympic IT team was able to do this despite having virtually no understanding at first of how the malware, called Olympic Destroyer, worked. It wasn’t until several days after the attack that analysts began unraveling the origins of the worm, which seemed deliberately designed to send security researchers on a wild goose chase as they tried to analyze the code and identify its source.

尽管一开始几乎不了解名为 Olympic Destroyer 的恶意软件的工作原理，但 Olympic IT 团队还是能够做到这一点。直到攻击发生几天后，分析人员才开始揭开蠕虫的起源，这似乎是故意设计的，目的是让安全研究人员在试图分析代码并确定其来源时大惊小怪。

The lesson for SREs: Preparation is golden. You can never know exactly what’s going to hit your systems, and in many cases, you won’t be able to identify root causes until well after you’re knee-deep in an outage. Nonetheless, by performing dry-runs and developing the right playbooks, you’ll position yourself to react effectively even in response to attacks of mystifying complexity.

SRE 的教训：准备是金。您永远无法确切地知道什么会影响您的系统，而且在许多情况下，直到您完全陷入停电之后，您才能确定根本原因。尽管如此，通过执行试运行和制定正确的剧本，即使在应对神秘复杂的攻击时，您也可以有效地做出反应。

## DNS outage strikes the Games

## DNS 中断影响了奥运会

From an IT perspective, the Olympic Games currently taking place in Tokyo have gone pretty smoothly, despite the challenges caused by the pandemic that postponed them.

从 IT 角度来看，尽管大流行带来的挑战推迟了奥运会，但目前在东京举行的奥运会进展顺利。

Nonetheless, a [temporary disruption to the Games's website and app](https://www.nbcnewyork.com/news/widespread-outage-disrupts-major-retail-financial-travel-websites-worldwide/3168859/) just as Olympic events were getting underway raised early fears that things would not proceed so well. The incident, which also affected the websites of a variety of major retailers, stemmed from a problem with Akamai’s DNS network, which the company attributed to a bad software update.

尽管如此，[暂时中断奥运会网站和应用程序](https://www.nbcnewyork.com/news/widespread-outage-disrupts-major-retail-financial-travel-websites-worldwide/3168859/)就像奥运会一样事件正在发生，人们早期担心事情不会进展得如此顺利。该事件还影响了多家主要零售商的网站，其根源在于 Akamai 的 DNS 网络问题，该公司将其归因于软件更新不当。

Akamai didn’t release further details, but from the looks of things, this was an SRE 101 type of incident. Presumably, a bug somewhere in a software release eluded testing routines and made it into production.

Akamai 没有透露更多细节，但从表面上看，这是一起 SRE 101 类型的事件。据推测，软件版本中某处的错误避开了测试例程并将其投入生产。

The good news is that Akamai resolved the incident in about an hour. Did they perform a rollback or redirect traffic to backup infrastructure? We’ll probably never know, but what is clear is that they had a plan in place for responding quickly to one of the most common sources of IT disruptions: A bad application update. Thanks to their preparation, most Olympic viewers never even knew that an outage had occurred.

好消息是 Akamai 在大约一个小时内解决了该事件。他们是否执行了回滚或将流量重定向到备份基础设施？我们可能永远不会知道，但很明显的是，他们制定了一个计划来快速响应 IT 中断的最常见来源之一：错误的应用程序更新。由于他们的准备，大多数奥运观众甚至都不知道发生了停电。

## Conclusion

##  结论

Although most Games over the past two decades have witnessed some disruptions to their IT infrastructure, the teams responsible for managing reliability for the Olympics deserve a lot of credit. To date, no show-stopping outage has taken place. That’s a pretty good reliability scorecard when you’re dealing with the systems behind the largest, most-watched sporting event in the world. 

尽管过去 20 年的大多数奥运会都见证了其 IT 基础架构的一些中断，但负责管理奥运会可靠性的团队值得称赞。迄今为止，还没有发生停止表演的停电事件。当您处理世界上最大、最受关注的体育赛事背后的系统时，这是一个非常好的可靠性记分卡。

