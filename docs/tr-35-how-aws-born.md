# How was the world's largest cloud vendor, AWS, born?

# 全球最大的云供应商 AWS 是如何诞生的？

2021-08-25 19:53:59 HKT From: https://inf.news/en/tech/6de7b65ded3cf076349d96cb57ee3faa.html

In 2000, when the Internet bubble burst, I worked at Amazon. At that  time, the capital market dried up, and we burned one billion US dollars  every year. Among them, the biggest overhead is the data center, which is full of expensive Sun servers. In order to reduce costs, we spent a  year "going to Sun" and replacing it with HP/Linux. This laid the  foundation for AWS!

2000 年，当互联网泡沫破灭时，我在亚马逊工作。那时资本市场枯竭，我们每年烧掉10亿美元。其中，开销最大的是数据中心，里面堆满了昂贵的Sun服务器。为了降低成本，我们花了一年时间“去 Sun”，换上HP/Linux。这为 AWS 奠定了基础！

###  Expensive IT infrastructure

### 昂贵的 IT 基础设施

In 1999, I entered Amazon. In the first week of work, I met Scott McCliny (Sun  founder and CEO) in the elevator, who was about to go to Bezos' office. At that time, Sun was one of the most valuable companies in the world,  with its highest market value exceeding US$200 billion and 50,000  employees worldwide.

1999年，我进入亚马逊。在工作的第一周，我在电梯里遇到了斯科特麦克林尼（Sun 创始人兼 CEO），他正准备去贝索斯的办公室。当时，Sun 是世界上最有价值的公司之一，其最高市值超过 2000 亿美元，在全球拥有 50,000 名员工。

Although Sun's products are not cheap, for  example, a workstation costs tens of thousands of dollars and a server  costs 100,000 dollars, but its products are not worried about sales,  because buying Sun's products is like buying IBM's things , "No one will be fired ."

Sun的产品虽然不便宜，比如一台工作站几万块，一台服务器10万块，但它的产品并不担心销量，因为买Sun的产品就像买IBM的东西，“没人会被炒鱿鱼” 。

Our motto is "Rapid Growth". For this reason, the  stability of the website is very important - every second of downtime  means lost sales, so we spend a lot of money to maintain the normal  operation of the website.

我们的座右铭是“快速增长”。为此，网站的稳定性非常重要——每一秒的宕机都意味着销售的损失，所以我们花费了大量的资金来维护网站的正常运行。

And Sun servers are the most reliable. Although its proprietary stack is expensive and tricky, all Internet companies use them.

而 Sun 服务器是最可靠的。尽管其专有堆栈昂贵且棘手，但所有互联网公司都在使用它们。

### The .Com bubble burst and Amazon began to "go to Sun"

### .Com 泡沫破灭，亚马逊开始“去 Sun”

During the "go to Sun" transition period, product development  stagnated, and we frozen all new features for a year many. Before us,  there is a huge to-do list, but nothing can be released until the  migration to Linux is complete.

在“去 Sun”过渡期间，产品开发停滞不前，我们将所有新功能冻结了一年多。在我们面前，有一个巨大的待办事项清单，但在迁移到 Linux 完成之前，什么都不能发布。

I clearly remember that in an  all-staff meeting, a vice president of engineering quickly showed a  picture of a snake swallowing a mouse.

我清楚地记得，在一次全体员工会议上，一位工程副总裁迅速展示了一张蛇吞下老鼠的图片。

This happened at the same  time as the slowdown in revenue growth—and further exacerbated the  slowdown, because we also had to raise prices to slow the company's  burnout rate. This is a vicious circle. We have no time or money. At this time, in a few more quarters, Amazon will go bankrupt.

这与收入增长放缓同时发生——并进一步加剧了放缓，因为我们还不得不提高价格以减缓公司的倦怠率。这是一个恶性循环。我们没有时间或金钱。这时候，再过几个季度，亚马逊就会破产。

However, once we start migrating to Linux, there is no turning back. Everyone is mobilized to refactor our code base, replace servers, and  prepare to switch. If successful, infrastructure costs will drop by more than 80%. If it fails, the website will collapse and the company will  die.

然而，一旦我们开始迁移到 Linux，就没有回头路了。每个人都被动员起来重构我们的代码库，更换服务器，并准备切换。如果成功，基础设施成本将下降80%以上。如果失败，网站将崩溃，公司将死亡。

In the end, we completed the migration in a timely and  smooth manner. This is a huge achievement for the entire engineering  team. The website continued to function without crashing. Capital  expenditures dropped drastically overnight. Suddenly, we have an  infinitely scalable infrastructure.

最终，我们及时、顺利地完成了迁移。这对整个工程团队来说是一个巨大的成就。该网站继续运行而没有崩溃。资本支出在一夜之间急剧下降。突然之间，我们拥有了一个无限可扩展的基础设施。

###  Starting from renting servers, AWS was born

### 从租用服务器开始，AWS 诞生了

Then, more interesting things happened. As a retailer, we have been  facing huge seasonal problems. In November and December every year,  traffic and revenue will soar. Bezos began to think, we have a surplus  of 46 weeks a year, why not rent it to other companies?

然后，更有趣的事情发生了。作为零售商，我们一直面临着巨大的季节性问题。每年11月和12月，流量和收入都会猛增。贝索斯开始思考，我们一年有 46 周的盈余，为什么不租给其他公司呢？

At about  the same time, he also became interested in decoupling internal  dependencies, so that the construction of the team could not be  restricted by other teams. The architectural changes required to  implement this loosely coupled model became AWS's API primitives.

大约在同一时间，他也开始对解耦内部依赖产生兴趣，让团队的建设不受其他团队的限制。实现这种松散耦合模型所需的架构更改成为 AWS 的 API 原语。

These are basic insights from AWS. I remember that Bezos demonstrated  this idea formed in the context of the power grid in a plenary meeting. In 1900, companies had to build their own generators to open stores. But why did enterprises in 2000 still have to build their own data centers?

这些是来自 AWS 的基本见解。我记得贝索斯在一次全体会议上展示了这个在电网背景下形成的想法。 1900 年，公司不得不建造自己的发电机来开设商店。但是为什么2000年的企业还是要自建数据中心呢？

Without AWS, cloud infrastructure will eventually appear (just as there is no Tesla, so will electric cars), but how long will the time be  delayed, and what is the opportunity cost? After AWS drastically reduced the cost of starting a company, innovation exploded and the modern  venture capital ecosystem was born.

没有AWS，云基础设施最终会出现（就像没有特斯拉，电动汽车也会出现），但时间会延迟多久，机会成本是多少？在 AWS 大幅降低创办公司的成本之后，创新爆发了，现代风险投资生态系统诞生了。

From 2000 to 2003, Amazon  almost went bankrupt. But without this crisis, the company is unlikely  to make this difficult decision to move to a brand new architecture. Without this shift, AWS may never appear. Find opportunities in the  crisis!

从 2000 年到 2003 年，亚马逊几乎破产。但如果没有这场危机，该公司不太可能做出迁移到全新架构的艰难决定。如果没有这种转变，AWS 可能永远不会出现。在危机中寻找机会！

Also: Amazon has recently spent several years going to Oracle, and few people have done so. Doing difficult things requires  strength, and strength is exercised by doing difficult things. The best  companies see every challenge as an opportunity and incorporate this  philosophy into their culture. 

另外：亚马逊最近花了几年时间去甲骨文，很少有人这样做。做困难的事情需要力量，做困难的事情锻炼力量。最好的公司将每一次挑战都视为机遇，并将这种理念融入到他们的文化中。

