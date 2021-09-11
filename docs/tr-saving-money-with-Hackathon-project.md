## How A Tiny Go Microservice Coded In Hackathon Is Saving Us Thousands

## 在 Hackathon 中编写的微型 Go 微服务如何为我们节省数千美元

10 Apr 2018

In the past few weeks we've rolled out a Go microservice called "KIB" to production, which reduced a huge portion of the infrastructure necessary for [Movio Cinema](https://movio.co/en/movio-cinema/) 's core tools: Group Builder and Campaign Execution. In doing this, we saved considerable AWS bills, after-hours and office-hours support bills, significantly simplified our architecture and made our product 80% faster on average. We wrote KIB on a Hackathon.

在过去的几周里，我们推出了一个名为“KIB”的 Go 微服务到生产中，这减少了 [Movio Cinema](https://movio.co/en/movio-cinema/) 所需的大部分基础设施的核心工具：Group Builder 和Campaign Execution。通过这样做，我们节省了可观的 AWS 账单、下班后和办公时间的支持账单，显着简化了我们的架构并使我们的产品平均速度提高了 80%。我们在 Hackathon 上编写了 KIB。

Despite the fact that the domain-specific components of this post don't apply to every dev team, the success of applying the [guiding principles](https://go-proverbs.github.io/) of simplicity and pragmatism driving our decision-making process felt like a story worth sharing.

尽管本文中特定领域的组件并不适用于每个开发团队，但成功应用 [指导原则](https://go-proverbs.github.io/) 的简单和实用主义推动了我们的决策过程感觉就像一个值得分享的故事。

## Group Builder

## 组构建器

Movio Cinema's core functionality is to send targeted marketing campaigns to a cinema chain's loyalty members. If you're a member of a big cinema chain’s loyalty program, wherever you are in the world, chances are the emails you're getting from it come from us.

Movio Cinema 的核心功能是向连锁影院的忠诚会员发送有针对性的营销活动。如果您是大型连锁影院忠诚度计划的成员，无论您身在何处，您收到的电子邮件都可能来自我们。

Group Builder facilitates the segmentation of loyalty members by the construction of groups of filters of varying complexity over the cinema's loyalty member base.

Group Builder 通过在电影院的忠诚会员基础上构建不同复杂度的过滤器组来促进忠诚会员的细分。

[Marketing teams around the world](https://movio.co/en/blog/category/case-studies/)  build the perfect audiences for their marketing campaigns using this tool.

[世界各地的营销团队](https://movio.co/en/blog/category/case-studies/) 使用此工具为其营销活动建立完美的受众。

## The Group Builder algorithm

## Group Builder 算法

By using this algorithm, an arbitrarily complex set of filters can be solved with a single SQL query. Note that it's domain-agnostic; you can use this strategy for filtering any set of elements.

通过使用该算法，可以使用单个 SQL 查询解决任意复杂的过滤器集。请注意，它与域无关；您可以使用此策略过滤任何元素集。

### Constraints

### 约束

- A Group Builder group can have any number of filters.
- Filters can be grouped together in an arbitrary number of sub-groups.
- Strictly for UI/UX reasons, sub-groups can be nested only once (i.e. up to sub-sub-groups).
- Filters and groups operate against each other using[Algebra of Sets](https://en.wikipedia.org/wiki/Algebra_of_sets).
- Each filter can either include or exclude set elements, set elements being loyalty members. The UI shows filters as green when they include members, and red when it excludes members.

- Group Builder 组可以有任意数量的过滤器。
- 过滤器可以组合成任意数量的子组。
- 严格来说，出于 UI/UX 的原因，子组只能嵌套一次（即最多嵌套到子子组）。
- 过滤器和组使用[集合代数](https://en.wikipedia.org/wiki/Algebra_of_sets) 相互操作。
- 每个过滤器可以包括或排除集合元素，集合元素是忠诚度成员。 UI 在过滤器包含成员时显示为绿色，在排除成员时显示为红色。

Here's the SQL strategy

这是 SQL 策略

This would be the final query for the Group Builder UI image above; note that the 3 filter subqueries represent the 3 filters shown in the image:

这将是对上述 Group Builder UI 图像的最终查询；请注意，3 个过滤器子查询代表图像中显示的 3 个过滤器：

```sql
SELECT
t.loyaltyMemberID,
MAX(CASE WHEN t.filter = 1 THEN 1 ELSE 0 END) AS f1,
MAX(CASE WHEN t.filter = 2 THEN 1 ELSE 0 END) AS f2,
MAX(CASE WHEN t.filter = 3 THEN 1 ELSE 0 END) AS f3

FROM (
       (
                       SELECT loyaltyMemberID, 1 AS filter FROM … // Gender
       ) UNION ALL (
                       SELECT loyaltyMemberID, 2 AS filter FROM … // Age
       ) UNION ALL (
                       SELECT loyaltyMemberID, 3 AS filter FROM … // Censor
       )
) AS t
GROUP BY t.loyaltyMemberID
HAVING
       f1 = 1 // = 1 means “Include”
AND (
       f2 != 1 // != 1 means “exclude”
OR
       f3 = 1
) // Parens equate to subgroups in the UI
```


Each sub-query in the UNION section selects the set of loyalty members after applying each filter in the group. The result set of the UNION (before the GROUP BY) will have one row per member per filter. The GROUP BY together with the AGGREGATE FUNCTIONs in the main SELECT section provide a very simple way to replicate the set algebra specified in the Group Builder UI, which you can see cleanly separated in the HAVING section.

UNION 部分中的每个子查询在应用组中的每个过滤器后选择忠诚度成员集。 UNION 的结果集（在 GROUP BY 之前）每个过滤器的每个成员将有一行。 GROUP BY 与主 SELECT 部分中的 AGGREGATE FUNCTION 一起提供了一种非常简单的方法来复制 Group Builder UI 中指定的集合代数，您可以在 HAVING 部分中看到它们完全分开。

## The limits of MySQL

## MySQL 的限制

The Group Builder algorithm worked really well in the beginning, but after some customers reached more than 5 million members (and 100 million sales transactions) the queries simply became way too slow to be able to provide timely feedback.

Group Builder 算法一开始运行得非常好，但是在一些客户达到超过 500 万会员（和 1 亿销售交易）之后，查询变得太慢而无法提供及时的反馈。

We needed an option that was fast but didn't require re-architecting our whole product. That option was [InfiniDB](https://en.wikipedia.org/wiki/InfiniDB). This was 2014.

我们需要一个快速但不需要重新设计我们整个产品的选项。该选项是 [InfiniDB](https://en.wikipedia.org/wiki/InfiniDB)。这是 2014 年。

## InfiniDB: a magic drop-in replacement 

## InfiniDB：神奇的替代品

InfiniDB was a MySQL almost-drop-in replacement that stored data in [columnar format](https://en.wikipedia.org/wiki/Column-oriented_DBMS). As such, and given our queries were rarely involving many fields, our InfiniDB implementation was a big success. We didn't stop using MySQL; instead we replicated data to InfiniDB in near real-time.

InfiniDB 几乎是 MySQL 的替代品，它以 [列格式](https://en.wikipedia.org/wiki/Column-oriented_DBMS) 存储数据。因此，鉴于我们的查询很少涉及许多领域，我们的 InfiniDB 实现取得了巨大成功。我们并没有停止使用 MySQL；相反，我们几乎实时地将数据复制到 InfiniDB。

Super-slow groups were fast again! We rolled out our InfiniDB implementation to all big customers and saved the day.

超慢组又快了！我们向所有大客户推出了我们的 InfiniDB 实现并挽救了这一天。

## The cons of our implementation of InfiniDB

## 我们实现 InfiniDB 的缺点

Despite the success of our solution, it wasn't without significant costs:

尽管我们的解决方案取得了成功，但并非没有重大成本：

- The InfiniDB instances [required a lot of memory and CPU](https://github.com/infinidb/infinidb/issues/21) to function properly, so we had to put them on [i3.2xlarge](https://aws.amazon.com/ec2/instance-types/i3/) EC2 instances. This was [very expensive](https://www.ec2instances.info/?cost_duration=annually&selected=i3.2xlarge) (~$7k per annum), considering we didn't charge extra for InfiniDB-powered Movio Cinema consoles.
- InfiniDB didn't have a master-slave scheme for replication (except for[this](https://dba.stackexchange.com/questions/174133/infinidb-replication-and-failover)), so we had to build our own custom one. Unfortunately, table name-swapping sometimes got some tables corrupted and we had to re-bootstrap the whole schema to come back up online, a process that could take several hours.
- We had a[recurring bug](https://github.com/infinidb/infinidb/issues/3) where some complex count queries would just return 0 results on InfiniDB, but the same query on MySQL wouldn't. It was never resolved.
- Even though InfiniDB was much faster than MySQL, we still saw some slow running groups for every big customer every week throughout 2017.
- The company behind InfiniDB sadly[had to shutdown](http://infinidb.co/forum/important-announcement-infinidb) and was eventually [incorporated into MariaDB](https://mariadb.com/about-us/newsroom/press-releases/open-source-leader-mariadb-rockets-analytics-market) (now [MariaDB ColumnStore](https://mariadb.com/products/technology/columnstore)); this meant no updates nor bugfixes.
- InfiniDB had no substantial community, so it was really hard to troubleshoot any issues. There were only[120 questions](https://stackoverflow.com/search?q=infinidb) on StackOverflow at the time of this writing.

- InfiniDB 实例[需要大量内存和 CPU](https://github.com/infinidb/infinidb/issues/21) 才能正常运行，因此我们不得不将它们放在 [i3.2xlarge](https://aws.amazon.com/ec2/instance-types/i3/) EC2 实例。这是 [非常昂贵](https://www.ec2instances.info/?cost_duration=annually&selected=i3.2xlarge)（每年约 7000 美元)，考虑到我们没有为 InfiniDB 驱动的 Movio Cinema 控制台收取额外费用。
- InfiniDB 没有用于复制的主从方案（除了 [this](https://dba.stackexchange.com/questions/174133/infinidb-replication-and-failover))，所以我们必须构建我们的自己定制的。不幸的是，表名交换有时会损坏一些表，我们必须重新引导整个模式才能恢复在线状态，这个过程可能需要几个小时。
- 我们有一个 [recurring bug](https://github.com/infinidb/infinidb/issues/3)，其中一些复杂的计数查询在 InfiniDB 上只会返回 0 结果，但在 MySQL 上的相同查询不会。一直没有解决。
- 尽管 InfiniDB 比 MySQL 快得多，但我们仍然在整个 2017 年每周看到每个大客户的一些运行缓慢的组。
- InfiniDB 背后的公司遗憾地[不得不关闭](http://infinidb.co/forum/important-announcement-infinidb) 并最终[并入 MariaDB](https://mariadb.com/about-us/newsroom /press-releases/open-source-leader-mariadb-rockets-analytics-market）（现在是 [MariaDB ColumnStore](https://mariadb.com/products/technology/columnstore))；这意味着没有更新或错误修正。
- InfiniDB 没有实质性社区，因此很难解决任何问题。在撰写本文时，StackOverflow 上只有 [120 个问题](https://stackoverflow.com/search?q=infinidb)。

## Research time

## 研究时间

During the winter of 2017 we had an upcoming [Hackathon](https://movio.co/en/blog/Hackathon-August-2017/), so three of us decided to do a little research on how to create an alternative solution for Group Builder.

2017 年冬天，我们有一个即将到来的 [Hackathon](https://movio.co/en/blog/Hackathon-August-2017/)，所以我们三个人决定研究一下如何创建替代解决方案对于 Group Builder。

We found that, even in slow-running groups, most of the filters were relatively simple and fast queries, and retrieving the resulting set of members was incredibly fast as well (~one million per second). The slowness was mostly in the final aggregation step, which could produce billions of intermediate rows on groups with many filters.

我们发现，即使在运行缓慢的组中，大多数过滤器都是相对简单和快速的查询，并且检索结果成员集的速度也非常快（约每秒 100 万次）。缓慢主要出现在最后的聚合步骤中，这可能会在具有许多过滤器的组上产生数十亿个中间行。

However, a few filters were slow no matter what. Even on fully indexed query execution plans they still had to scan over half a billion rows.

但是，无论如何，一些过滤器都很慢。即使在完全索引的查询执行计划中，他们仍然需要扫描超过 50 亿行。

How could we circumvent these two issues with a simple solution achievable in a day?

我们怎样才能用一天内可以实现的简单解决方案来规避这两个问题呢？

## Hackathon idea: the KIB project

## 黑客马拉松的想法：KIB 项目

Our solution was:

我们的解决方案是：

1) Reviewing all slow filters for quick wins (e.g. adding/changing indexes, reworking queries)

1) 检查所有慢速过滤器以获得快速成功（例如添加/更改索引、重新处理查询）

2) Running every filter as a separate query against MySQL concurrently, and doing the aggregation programmatically using sparse bitsets

2）同时运行每个过滤器作为针对 MySQL 的单独查询，并使用稀疏位集以编程方式进行聚合

3) Caching filter results for a number of minutes to minimise the time recalculating long-running queries, given the repetitive usage pattern shown by our customers

3) 考虑到客户显示的重复使用模式，将过滤器结果缓存几分钟，以最大限度地减少重新计算长时间运行的查询的时间

After the Hackathon, we quickly added two planned features that covered outstanding problematic cases:

在 Hackathon 之后，我们迅速添加了两个涵盖突出问题案例的计划功能：

1) Refreshing caches automatically, to make most frequently used filters and slow-running ones very quick at all times.

1) 自动刷新缓存，使最常用的过滤器和运行缓慢的过滤器始终非常快速。

2) Pre-caching on startup based on usage history.

2) 根据使用历史在启动时预缓存。

We packed these features into a tiny Go microservice called KIB and shipped it onto a c4.xlarge EC2 instance (<$3k per annum). While before we had one InfiniDB instance on a i3.2xlarge for each customer, in this case we put all customers on the same single c4.xlarge instance. 

我们将这些功能打包到一个名为 KIB 的小型 Go 微服务中，并将其发送到 c4.xlarge EC2 实例（每年 <3,000 美元）。之前我们在 i3.2xlarge 上为每个客户设置一个 InfiniDB 实例，而在本例中，我们将所有客户放在同一个 c4.xlarge 实例上。

We did add a second instance for fault tolerance, and it was a fortunate decision, because that very week our EC2 instance died. Thanks to our [Kubernetes cluster implementation](https://movio.co/en/blog/6-months-kubernetes-production/), the KIB instance quickly restarted on the second healthy node and no customers were impacted. In contrast, when our InfiniDB nodes died, re-bootstrapping would sometimes take hours. Note, however, that this was not an intrinsic problem of InfiniDB itself, but of our custom replication implementation of it.

我们确实添加了第二个容错实例，这是一个幸运的决定，因为就在那个星期，我们的 EC2 实例死了。感谢我们的 [Kubernetes 集群实现](https://movio.co/en/blog/6-months-kubernetes-production/)，KIB 实例在第二个健康节点上快速重启，没有客户受到影响。相比之下，当我们的 InfiniDB 节点死亡时，重新引导有时需要几个小时。但是请注意，这不是 InfiniDB 本身的内在问题，而是我们对其的自定义复制实现的内在问题。

## Why Go?

The long answer to that question is described in [this blogpost](https://movio.co/en/blog/migrate-Scala-to-Go/). In short, we've been coding in Go for about a year, and it has noticeably reduced our estimations and workload, made our code much simpler, which in turn has indirectly improved our software quality, and in general has made our work more rewarding .

[这篇博文](https://movio.co/en/blog/migrate-Scala-to-Go/) 中描述了该问题的详细答案。简而言之，我们用 Go 编码大约一年了，它显着减少了我们的估计和工作量，使我们的代码更简单，这反过来又间接提高了我们的软件质量，总的来说，使我们的工作更有收获.

Here are the key bits of the KIB Go code with only a few details elided:

以下是 KIB Go 代码的关键部分，仅省略了一些细节：

An endpoint request represents a Group Builder group to be run, and the group is represented as a tree of SQL queries. It looks somewhat like this:

端点请求表示要运行的 Group Builder 组，该组表示为 SQL 查询树。它看起来有点像这样：

```go
type request struct {
N SQLNode `json:"node"`
}

type SQLNode struct {
SQL        string    `json:"sql"`
Operator   string    `json:"operator"`
Nodes      []SQLNode `json:"nodes"`
}
```


This is the function that takes the request and resolves every SQL in the tree to a bitset:

这是接受请求并将树中的每个 SQL 解析为位集的函数：

```go
func (r request) toBitsets(solver *solver) (*bitsetWr, error) {
var b = bitsetWr{new(tree)}
b.root.fill(r.treeRoot, solver) // runs all SQL concurrently
solver.wait() // waits for all goroutines to finish
return &b, solver.err
}
```


This is the inner function that deals with the tree structure; running “solve” on every node:

这是处理树结构的内部函数；在每个节点上运行“solve”：

```go
func (t *tree) fill(n SQLNode, solver *solver) {
if len(n.Nodes) == 0 { // if leaf
t.op = n.Operator
solver.solve(n.SQL, t) // runs SQL;fills bitset
return
}
t.nodes = make([]*tree, len(n.Nodes)) // if group
for i, nn := range n.Nodes {
t.nodes[i] = &tree{}
t.nodes[i].fill(nn, solver)
}
}
```


And this is the inner function that runs the SQL (or loads from cache) concurrently:

这是并发运行 SQL（或从缓存加载）的内部函数：

```go
func (r *solver) solve(sql string, b *tree) {
r.wg.Add(1)
go func(b *tree) { // returns immediately
defer r.wg.Done()
res, err := r.cacheMaybe(sql) // runs SQL or uses cache
if err != nil {
r.err = err
return
}
b.bitset = res
}(b)
}
```


There are about 50 more lines for solving the algebra and 25 more for the basic caching, but these three snippets are a representative example of what the KIB code looks like.

大约还有 50 多行代码用于解决代数问题，另外还有 25 行用于基本缓存，但这三个片段是 KIB 代码外观的代表性示例。

The snippets show some intermediate concepts: tree traversal, green threads, bitset operations and caching. While not that uncommon, what I've rarely seen in practice is an implementation of all these things solving a real business problem within a day's work. We don't think we would have been able to pull this off in any other language. I'll explain why we think so in the conclusion.

这些片段显示了一些中间概念：树遍历、绿色线程、位集操作和缓存。虽然并不少见，但我在实践中很少看到在一天的工作中实现所有这些事情来解决真正的业务问题。我们认为我们无法用任何其他语言实现这一点。我将在结论中解释我们为什么这么认为。

## Results

##  结果

The results of our project were extremely gratifying.

我们项目的结果非常令人满意。

In financial terms, we saved $31.000 yearly in AWS bills and $4.000 yearly in after-hours support.

在财务方面，我们每年在 AWS 账单上节省了 31.000 美元，在下班后支持上每年节省了 4.000 美元。

In-office support & maintenance probably costed more than the previous ones combined in the last year, but these are not easily quantifiable in dollar value.

办公室内支持和维护的成本可能比去年的前几项加起来还要多，但这些都不容易用美元价值来量化。

Note that our InfiniDB-based solution has undergone yearly rewrites (we've had to rewrite our InfiniDB solution in 2015 and again in 2016 since our original 2014 implementation).

请注意，我们基于 InfiniDB 的解决方案每年都会进行重写（自 2014 年最初实施以来，我们不得不在 2015 年和 2016 年再次重写我们的 InfiniDB 解决方案）。

By my calculation, we will save more than $50,000 this year and a similar amount every year going forward.

根据我的计算，我们今年将节省超过 50,000 美元，以后每年都会节省类似的金额。

In terms of speed, here's a week-on-week table of results for 9 of our biggest customers (we've replaced the actual names with their countries of origin):

在速度方面，以下是我们 9 个最大客户的每周结果表（我们已将实际名称替换为他们的原籍国）：

For some customers, this upgrade meant not waiting for Group Builder anymore, as it was the case for this USA customer:

对于某些客户而言，此升级意味着不再等待 Group Builder，就像这位美国客户的情况一样：

Globally, it was an 81% average improvement in group run times; an average constructed from every single Group Builder group ran on those weeks, not from a sample. That really exceeded our own expectations. 

在全球范围内，小组运行时间平均提高了 81%；由每个单独的 Group Builder 组构建的平均值在这些周内运行，而不是来自样本。这确实超出了我们自己的预期。

For us devs, replacing our complex custom replication implementation of InfiniDB that kept us up at night every other week with such a tiny and simple Go microservice we built on a Hackathon is the greatest gift.

对于我们开发人员来说，用我们在 Hackathon 上构建的如此微小而简单的 Go 微服务替换我们复杂的 InfiniDB 自定义复制实现，这让我们每隔一周都夜不能寐。

## Conclusions

## 结论

Throughout the Hackathon, we spent the day researching, designing and coding, and no more than half an hour debugging. We didn't have dependency issues. We didn't have issues understanding each other's code, because it looked exactly the same and just as simple. We fully understood the tools we were using, because we have been using the same bare bone building blocks for a year. We didn't struggle with debugging, because the only bugs we had were silly mistakes that were either found early by the IDE's linter or clearly explained by the compiler with error messages written for humans. All of this was possible thanks to the Go language.

在整个 Hackathon 中，我们花了一天的时间研究、设计和编码，调试时间不超过半小时。我们没有依赖问题。我们在理解彼此的代码方面没有问题，因为它看起来完全一样，而且同样简单。我们完全了解我们使用的工具，因为我们已经使用相同的基本构建块一年了。我们没有在调试中挣扎，因为我们遇到的唯一错误是愚蠢的错误，这些错误要么被 IDE 的 linter 早期发现，要么被编译器清楚地解释为为人类编写的错误消息。多亏了 Go 语言，这一切才成为可能。

KIB is in production today on all major customers, and it's barely using a few hundred MB of RAM per customer and sharing 4 vCPUs among all of them. Even though we aggressively parallelised SQL query execution and bitset operations, we had no issues at all related to this: no "too many connections" to MySQL, no bugs related to concurrency, etc; not even while writing it. We did have one pointer-related bug and one silly tree traversal edge case; we're human.

KIB 现已在所有主要客户的生产中投入使用，每个客户几乎不使用几百 MB 的 RAM，并在所有客户之间共享 4 个 vCPU。即使我们积极地并行化 SQL 查询执行和 bitset 操作，我们也没有任何与此相关的问题：没有与 MySQL 的“太多连接”，没有与并发相关的错误等；甚至在写的时候也不行。我们确实有一个与指针相关的错误和一个愚蠢的树遍历边缘情况；我们是人。

What's the lesson learned here? My lesson learned is the power and value of simplicity. Even within the simplicity of Go, we went with the simplest (not the easiest) possible subset of it: we didn't use channels, interfaces, panics, named returns, iotas, mutexes (we did have one WaitGroup for the one set of goroutines). We only used goroutines and pointers where necessary.

这里有什么教训？我学到的教训是简单的力量和价值。即使在 Go 的简单性中，我们也使用了它的最简单（不是最简单）的可能子集：我们没有使用通道、接口、恐慌、命名返回、iota、互斥（我们确实有一个 WaitGroup 用于一组协程）。我们只在必要时使用 goroutines 和指针。

Thanks for reading this blogpost. KISS!

感谢您阅读这篇博文。吻！

[Mariano Gappa](http://movio.co/blog/author/mariano/)

Find me on [github](https://github.com/marianogappa) and [Twitter](https://twitter.com/MarianoGappa) 

在 [github](https://github.com/marianogappa) 和 [Twitter](https://twitter.com/MarianoGappa) 上找到我

