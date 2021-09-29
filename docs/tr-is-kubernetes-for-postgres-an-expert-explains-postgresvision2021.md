### Is Kubernetes for Postgres? An expert explains

### Kubernetes 是用于 Postgres 的吗？专家解释

by [Patrick Nelson](https://siliconangle.com/author/patricknelson/)

A database’s level of simplicity is the most important consideration for assessing when it’s worth shifting it to an advanced, sophisticated infrastructure platform, according to a Kubernetes application development specialist who works closely with PostgreSQL, an open-source relational database management system.

一位与开源关系数据库管理系统 PostgreSQL 密切合作的 Kubernetes 应用程序开发专家表示，在评估何时值得将其转移到先进、复杂的基础设施平台时，数据库的简单程度是最重要的考虑因素。

“The more commodifiable a particular database instance is, the better candidate it is to move,” said [Josh Berkus](https://www.linkedin.com/in/josh-berkus-1792412/))(pictured), Kubernetes community manager at Red Hat Inc.

“特定数据库实例的商品化程度越高，它就越适合移动，”[Josh Berkus](https://www.linkedin.com/in/josh-berkus-1792412/))（如图)，Kubernetes Red Hat Inc. 社区经理

Berkus, however, is not simply talking about levels of complication that might hinder the migration, but that there’s less gain to be made by doing the move under certain circumstances. “The real advantage of moving stuff to Kubernetes is your ability to automate things,” he explained.

然而，Berkus 并不是简单地谈论可能阻碍迁移的复杂程度，而是在某些情况下进行迁移所带来的收益较少。 “将东西转移到 Kubernetes 的真正优势在于你能够实现自动化，”他解释道。

Berkus spoke with [Dave Vellante](https://twitter.com/dvellante), host of theCUBE, SiliconANGLE Media's livestreaming studio, during [Postgres Vision 2021](https://www.thecube.net/postgres-vision-2021). They discussed just how appropriate Kubernetes is for Postgres data. _( Disclosure below.)_

Berkus 在 [Postgres Vision 2021](https://www.thecube.net/postgres-vision-2021) 期间与 SiliconANGLE Media 直播工作室 theCUBE 的主持人 [Dave Vellante](https://twitter.com/dvellante)进行了交谈)。他们讨论了 Kubernetes 对 Postgres 数据的适用性。 _(\\* 下面披露。)_

### What databases to migrate

### 要迁移哪些数据库

Massive company-wide databases, where the entire business operation is run using one elaborate database would be an example of a less-suitable candidate, according to Berkus.

Berkus 表示，大型公司范围的数据库，其中整个业务运营都使用一个精心设计的数据库运行，这将是不太合适的候选人的一个例子。

“To the extent that you can describe [a] particular database, what it does, who needs to use it, what’s in it in a simple one pager, then that’s probably a really good candidate for hosting on Kubernetes,” Berkus said.

“就你可以描述 [a] 特定数据库、它的作用、谁需要使用它、其中包含什么在一个简单的单页程序中的程度而言，那么这可能是在 Kubernetes 上托管的一个非常好的候选者，”Berkus 说。

Less appropriate are databases that are so complicated it gets hard, or even impossible, to explain the inputs and outputs.

不太合适的数据库太复杂以至于很难甚至不可能解释输入和输出。

“I’ve worked with people who have one big database, where the database is three terabytes in size. It powers their reporting system and their customer’s system and the web portal and everything else in one database,” Berkus said. “That’s the one that’s really going to be a hard call and that you might, in fact, never physically migrate to Kubernetes.”

“我曾与拥有一个大数据库的人一起工作，其中数据库大小为 3 TB。它为他们的报告系统、客户系统、门户网站以及一个数据库中的所有其他内容提供支持，”伯库斯说。 “这真的是一个艰难的决定，事实上，你可能永远不会物理迁移到 Kubernetes。”

fEqually, databases that aren’t going to be taking advantage of automation are less optimal candidates. Berkus advises that one should assess whether application workflow and team organization can handle the new setup. If that’s in place, and particularly if development is unified, along with an infra team that owns everything, “then those people are going to be a really good candidate for moving that stack to Kubernetes.”

同样，不会利用自动化的数据库不是最佳候选者。 Berkus 建议应该评估应用程序工作流和团队组织是否可以处理新设置。如果这一点到位，特别是如果开发是统一的，以及拥有一切的基础设施团队，“那么这些人将成为将该堆栈迁移到 Kubernetes 的非常好的候选人。”

But even if all the considerations don’t fall into place, all is not lost. Berkus believes users should take a look at Service Catalog in Kubernetes if they’re dead set on the migration. That’s a way of exposing an external service in Kubernetes and making it appear a Kubernetes service.

但即使所有的考虑都没有落实到位，也不会失去一切。 Berkus 认为，如果用户在迁移上死心塌地，应该看看 Kubernetes 中的服务目录。这是在 Kubernetes 中公开外部服务并使其看起来像 Kubernetes 服务的一种方式。

“That’s what I tend to do with those kinds of [complicated] databases,” he said.

“这就是我倾向于用那些 [复杂] 数据库做的事情，”他说。

Exceptions to the aforementioned do exist, however. Big data infrastructure on Postgres is a contender for Kubernetes, according to Berkus.

但是，上述情况确实存在例外情况。根据 Berkus 的说法，Postgres 上的大数据基础设施是 Kubernetes 的竞争者。

“A lot of modern data analysis and data-mining platforms are built on top of Postgres,” he said. “Part of how they do their work is they actually run a bunch of little Postgres instances that they federate together.” Kubernetes then can be the tool that lets you to manage “all of those little Postgres instances.”

“许多现代数据分析和数据挖掘平台都建立在 Postgres 之上，”他说。 “他们工作的部分方式是他们实际上运行了一堆他们联合在一起的小 Postgres 实例。”然后，Kubernetes 可以成为让您管理“所有这些小 Postgres 实例”的工具。

Don’t expect substantial performance differences by moving to Kubernetes if you’re already on cloud or network storage and have databases that can share hardware systems, according to the Postgres Kubernetes expert.

根据 Postgres Kubernetes 专家的说法，如果您已经在云或网络存储上并且拥有可以共享硬件系统的数据库，则不要指望迁移到 Kubernetes 会产生实质性的性能差异。

“Your whole point of moving to Kubernetes in general is going to be: take advantage of the automation,” Berkus concluded. 

“总体而言，迁移到 Kubernetes 的重点将是：利用自动化，”Berkus 总结道。

Watch the complete video interview below, and be sure to check out more of SiliconANGLE’s and theCUBE’s coverage of [Postgres Vision 2021](https://www.thecube.net/postgres-vision-2021). _( Disclosure: TheCUBE is a paid media partner for the Postgres Vision event. Neither EnterpriseDB Corp., the sponsor for theCUBE’s event coverage, nor other sponsors have editorial control over content on theCUBE or SiliconANGLE.)_ 

观看下面的完整视频采访，一定要查看更多 SiliconANGLE 和 theCUBE 对 [Postgres Vision 2021](https://www.thecube.net/postgres-vision-2021) 的报道。 _（\\* 披露：TheCUBE 是 Postgres Vision 活动的付费媒体合作伙伴。CUBE 活动报道的赞助商 EnterpriseDB Corp. 和其他赞助商都没有对 CUBE 或 SiliconANGLE 上的内容进行编辑控制。)_

