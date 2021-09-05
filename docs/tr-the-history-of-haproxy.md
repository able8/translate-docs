# The History of HAProxy

# HAProxy 的历史

[Willy Tarreau](https://www.haproxy.com/blog/author/wtarreau/ "Posts by Willy Tarreau") [Willy Tarreau](https://www.haproxy.com/blog/author/wtarreau/ " Posts by Willy Tarreau") \| Nov 8, 2019 \| [LOAD BALANCING / ROUTING](https://www.haproxy.com/blog/category/load-balancing-routing/),[TECH](https://www.haproxy.com/blog/category/tech/) \| [1 comment](https://www.haproxy.com/blog/the-history-of-haproxy/#respond)

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/History-of-HAProxy.png)

Today, HAProxy is the world’s fastest and most widely used software load balancer. It’s responsible for providing high availability, security, and observability to some of the world’s highest trafficked websites. Many Site Reliability Engineers and Network/Systems Engineers alike consider it to be an essential part of their critical infrastructure. However, HAProxy’s origin story is one that has not been told and you may be curious about its roots and what drove it to be what it is today.

今天，HAProxy 是世界上速度最快、使用最广泛的软件负载均衡器。它负责为世界上一些访问量最高的网站提供高可用性、安全性和可观察性。许多站点可靠性工程师和网络/系统工程师都认为它是其关键基础设施的重要组成部分。然而，HAProxy 的起源故事是一个尚未被告知的故事，您可能对它的起源以及是什么促使它成为今天的样子感到好奇。

All of this started in 1999 when I needed a way to gauge how an application would perform when facing lots of clients with 28 Kbps modems instead of a single 100-Mbps link. I decided to reuse a proxy I’d written earlier for compression, called _Zprox_. I removed the compression code, hacked some bandwidth limitations, and quickly had a useful testing tool.

所有这一切都始于 1999 年，当时我需要一种方法来衡量应用程序在面对大量使用 28 Kbps 调制解调器而不是单个 100 Mbps 链路的客户端时的性能。我决定重用我之前为压缩编写的代理，称为 _Zprox_。我删除了压缩代码，破解了一些带宽限制，很快就有了一个有用的测试工具。

A year later, I was involved in another benchmarking project where the people in charge of operations were facing an interoperability issue: Their load generating tool emitted HTTP headers in a different syntax than the one expected by the server. I reused my hacked proxy to implement regex-based header rewriting so that the requests were transformed as they passed through. At the same time, I started a minimalistic configuration language, since the proxy needed to listen on multiple ports with different regular expressions. The `listen` and `server` keywords were born—and they’re still supported 19 years later in HAProxy 2.0.

一年后，我参与了另一个基准测试项目，其中负责运营的人员面临一个互操作性问题：他们的负载生成工具以不同于服务器预期的语法发送 HTTP 标头。我重用了我被黑的代理来实现基于正则表达式的标头重写，以便请求在通过时进行转换。同时，我开始使用简约的配置语言，因为代理需要使用不同的正则表达式来监听多个端口。 `listen` 和 `server` 关键字诞生了——19 年后，HAProxy 2.0 仍然支持它们。

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-1@4x.png)

Then in mid-2001, a heavy application was deployed. This application was known for occasionally exhibiting huge response times, which caused connections to pile up in front of it. This was a situation that the hardware load balancers in place were totally unable to deal with, and they started to repeatedly fail, bringing all other applications down with them. Since it was not possible to fix the application quickly, our short-term workaround consisted of deploying my proxy—which, by then, people were calling “Willy’s proxy”—to offload some of the traffic from the hardware load balancers.

然后在 2001 年年中，部署了一个繁重的应用程序。此应用程序以偶尔表现出巨大的响应时间而闻名，这会导致连接堆积在它面前。这是硬件负载平衡器完全无法处理的情况，它们开始反复失败，使所有其他应用程序也随之宕机。由于无法快速修复应用程序，我们的短期解决方法包括部署我的代理——那时人们称之为“威利的代理”——从硬件负载平衡器卸载一些流量。

The traffic would pass through the original load balancers working at Layer 4 and then be forwarded to a pair of the proxies, which took over inspecting headers and persisting connections to servers. The server persistence worked by checking whether a cookie was present, then searching for its value among a list of predefined servers, then selecting the matched server. If no cookie was found, the connection was forwarded back to the original load balancers which would contact a server and add the cookie. We estimated that this reduced the Layer 7 work being handled by the original load balancers to a tenth or even a hundredth.

流量将通过在第 4 层工作的原始负载平衡器，然后转发到一对代理，代理负责检查标头并保持与服务器的连接。服务器持久性的工作方式是检查 cookie 是否存在，然后在预定义的服务器列表中搜索其值，然后选择匹配的服务器。如果未找到 cookie，则将连接转发回原始负载均衡器，该负载均衡器将联系服务器并添加 cookie。我们估计这将原始负载均衡器处理的第 7 层工作减少到十分之一甚至百分之一。

The solution was quick enough to implement using the existing proxy code, as it was already capable of handling HTTP headers. It wouldn’t have to deal with any load-balancing mechanics and only had to detect server availability. Finally, its name had been coined: “HAProxy” for “high availability proxy”. The solution was accepted, quickly implemented, then deployed. Since it was not very stable by then it received updates often, but it never crashed so no user ever noticed the quirks. This was HAProxy 1.0, released in 2001.

该解决方案足够快，可以使用现有的代理代码实现，因为它已经能够处理 HTTP 标头。它不必处理任何负载平衡机制，只需检测服务器可用性。最后，它的名字被创造出来了：“HAProxy”代表“高可用性代理”。该解决方案被接受，迅速实施，然后部署。由于那时它不是很稳定，它经常收到更新，但它从未崩溃，所以没有用户注意到这些怪癖。这是 2001 年发布的 HAProxy 1.0。

## Discovering Its Strengths 

## 发现它的优势

Historically, the network team was only in charge of the network equipment and never had to deal with anything related to the HTTP layer or traffic logs. Nevertheless, they were the ones from whom everyone asked for help, since they had a network packet capture device and, honestly, because it was harder for them to prove the problem was not on their side. The infrastructure team was in charge of firewalls and, therefore, had visibility into firewall logs. The applications people became accustomed to considering that anything not working was probably due to a firewall configuration issue. If that wasn’t the case, then they’d ask the network people for a packet capture.

过去，网络团队只负责网络设备，从来不需要处理与 HTTP 层或流量日志相关的任何事情。尽管如此，他们是每个人都向他们寻求帮助的人，因为他们有网络数据包捕获设备，而且老实说，因为他们更难证明问题不在他们这边。基础架构团队负责防火墙，因此可以查看防火墙日志。人们习惯于认为任何不工作的应用程序都可能是由于防火墙配置问题。如果不是这样，那么他们会要求网络人员进行数据包捕获。

Needless to say, this process does not scale. It also doesn’t work at all for symptoms like “it doesn’t function well enough.” With HAProxy in the hands of the infrastructure team, everything was suddenly simplified: the logs it produced immediately indicated where a connection originated, where it was sent, how long it spent waiting for the server, and which side closed it first. Alone, this was so precise and helpful to each team that it quickly became the first place to look. It was highly trusted, as it significantly reduced the time needed to solve any trouble. It also explains the sophisticated indicators that are included in the default log format.

不用说，这个过程不能扩展。对于“功能不够好”等症状，它也根本不起作用。 HAProxy 掌握在基础架构团队的手中后，一切都突然变得简单了：它生成的日志立即表明连接的来源、发送的位置、等待服务器的时间以及哪一方先关闭它。单独来看，这对每个团队都非常精确和有帮助，以至于它很快成为第一个关注的地方。它受到高度信任，因为它大大减少了解决任何问题所需的时间。它还解释了默认日志格式中包含的复杂指标。

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-2@4x.png)

The development of logging in HAProxy was also spurred by HTTP interoperability being far from good in 2001. You simply couldn’t trust a server’s logs to visualize what was actually on the wire, and having to take a network trace required a lot of time. HAProxy quickly gained the ability to capture:

- cookies, to diagnose stickiness issues
- headers, to diagnose their issues
- source ports, to correlate with firewalls
- timers, to detect timeouts and application server issues
- termination codes, to analyze failed sessions

2001 年 HTTP 互操作性还很差，也推动了 HAProxy 日志记录的发展。你根本无法相信服务器的日志来可视化实际在线情况，并且必须进行网络跟踪需要大量时间。 HAProxy 很快就获得了捕获：

- cookie，用于诊断粘性问题
- 标题，诊断他们的问题
- 源端口，与防火墙相关联
- 计时器，用于检测超时和应用程序服务器问题
- 终止代码，分析失败的会话

Initially, HAProxy didn’t modify anything in the request or response. However, the architecture—two layers of load balancers with HAProxy kicking the connection back to the first layer if no cookie was present—was complex and difficult to configure. So, eventually HAProxy had to learn about load balancing, which meant adding a simple round-robin scheduler, simple health checks, and cookie insertion. At that point, the architecture became cleaner: one layer of L3/L4 load balancers, a layer of HAProxy load balancers handling all the L7 work, then the application servers. HAProxy 1.1 was born, with a codebase three times as large as 1.0.

最初，HAProxy 没有修改请求或响应中的任何内容。然而，如果没有 cookie 存在，两层负载均衡器和 HAProxy 将连接踢回到第一层，架构复杂且难以配置。因此，最终 HAProxy 必须了解负载平衡，这意味着添加一个简单的循环调度程序、简单的健康检查和 cookie 插入。那时，架构变得更清晰：一层 L3/L4 负载均衡器，一层处理所有 L7 工作的 HAProxy 负载均衡器，然后是应用服务器。 HAProxy 1.1 诞生了，其代码库是 1.0 的三倍。

## A Decisive Coding Standard

## 决定性的编码标准

With the new responsibilities, performance started to become a concern.

有了新的职责，绩效开始成为一个问题。

HAProxy was mostly running on old Linux 2.4 systems made of 733 MHz, single-core, Pentium III processors and on a few Solaris 8 systems made of 400 MHz UltraSparc processors. Both were often equipped with 128 MB of RAM, in the best case! On such machines, every byte of memory and every CPU cycle is important. Great care was taken to never use more than necessary of either.

HAProxy 主要运行在由 733 MHz 单核 Pentium III 处理器组成的旧 Linux 2.4 系统和由 400 MHz UltraSparc 处理器组成的少数 Solaris 8 系统上。在最好的情况下，两者通常都配备 128 MB 的 RAM！在这样的机器上，每个内存字节和每个 CPU 周期都很重要。非常小心，不要使用超过必要的任何一种。

Structure fields were only initialized once and never reinitialized if their state was known. There was little error checking when it was known that the design would not allow certain conditions to happen. Content was parsed in one round, taking care to never read the same byte twice. Such an approach requires clear communication: Often the comments in the code explaining why an operation could be avoided were much longer than the code that was avoided!

结构字段只初始化一次，如果它们的状态已知，则永远不会重新初始化。当知道设计不允许某些条件发生时，几乎没有错误检查。内容在一轮中解析，注意不要两次读取相同的字节。这种方法需要清晰的沟通：通常代码中解释为什么可以避免操作的注释比被避免的代码长得多！

Even if a few of these design rules had to be revisited as the code became more modular, they still seeded a certain philosophy in the project that resources aren’t free and developers are not allowed to waste them. We had a mantra: “Keep the developer busy to make the code lazy”.

即使随着代码变得更加模块化，必须重新审视其中的一些设计规则，它们仍然在项目中播种了某种理念，即资源不是免费的，不允许开发人员浪费它们。我们有一句口头禅：“让开发人员忙着让代码变得懒惰”。

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-3@4x.png)



Not everything was perfectly designed. For example, we discovered, the hard way (i.e. in production), that under Solaris `select()` was limited to only 256 file descriptors by default; This led us to rush the implementation of the `nbproc` mechanism the same day, in order to start multiple processes, and it was considered a satisfying solution; So, it was only much later that we added support for `poll()`. Nowadays, as people use `nbproc` to improve their scalability, some wonder why certain states are not synchronized between the processes. It’s significant that what they are using was absolutely not designed for the same purpose at all, and that the performance benefit they value is just the byproduct of a nearly two-decades old workaround for a system limitation.

并非一切都经过完美设计。例如，我们发现了一个困难的方法（即在生产中），在 Solaris `select()` 下默认限制为只有 256 个文件描述符；这导致我们在当天就匆忙实现了 `nbproc` 机制，以便启动多个进程，这被认为是一个令人满意的解决方案；所以，直到很久以后我们才添加了对 `poll()` 的支持。现在，当人们使用 `nbproc` 来提高他们的可扩展性时，有些人想知道为什么某些状态在进程之间不同步。重要的是，他们所使用的东西绝对不是为相同目的而设计的，他们所重视的性能优势只是近两年前针对系统限制的解决方法的副产品。

By version 1.2, the project had started to attract outside users. Some hosting providers and large infrastructure providers needed a load balancer for specific services and they found that HAProxy was extremely fast and unbreakable compared to what they’d tried until then. The scope of use cases exploded: faster forwarding from servers to clients for video providers, detailed statistics for hosting providers who had to report to their customers, a CLI to make it easier to get statistics.

到 1.2 版本时，该项目已开始吸引外部用户。一些托管服务提供商和大型基础设施提供商需要针对特定服务的负载均衡器，他们发现 HAProxy 与他们之前尝试过的相比非常快速且牢不可破。用例的范围激增：视频提供商从服务器到客户端的更快转发，必须向客户报告的托管提供商的详细统计数据，使获取统计数据更容易的 CLI。

## Pushing Killer Performance

## 推动杀手级表现

Changes continued. It became obvious that backwards compatibility was very important so that users could upgrade without breaking a sweat. More log indicators were needed, and more stats. Then, a new trend of low-performance servers appeared, some limited to only a single connection, others to just a few. Since HAProxy’s connections were cheap, it casually fit the role of being a connection concentrator. The `maxconn` setting was born. Connections could be queued in HAProxy, which acted as a buffer for the servers.

变化还在继续。很明显，向后兼容性非常重要，这样用户就可以毫不费力地升级。需要更多的日志指标和更多的统计数据。然后，出现了低性能服务器的新趋势，有些仅限于单个连接，有些仅限于几个。由于 HAProxy 的连接很便宜，所以它随便扮演了一个连接集中器的角色。 `maxconn` 设置诞生了。连接可以在 HAProxy 中排队，它充当服务器的缓冲区。

This was a killer feature that protected servers from dying under load. Even better, by reducing the amount of resources a server had to use, it was possible to operate much faster. We held a number of impressive demos alongside popular application servers and showed how placing HAProxy in front sped them up, sometimes by an order of magnitude.

这是一个杀手级功能，可以保护服务器免于在负载下死亡。更好的是，通过减少服务器必须使用的资源量，可以更快地运行。我们与流行的应用程序服务器一起举办了许多令人印象深刻的演示，并展示了将 HAProxy 放在前面如何加快它们的速度，有时会提高一个数量级。

The rapid growth of high-bandwidth websites (i.e. gigabit level) made it mandatory to further think in terms of layers, granting lower layers autonomy so that higher ones could be called less often. This update arrived in HAProxy 1.3.16 and was really the key to pushing HAProxy to where it needed to be to deal with 10 Gbps links. This was thanks to TCP splicing, which, when supported by the network card, allowed data to be proxied between two sides while using only a single copy in memory. By 1.3.17 in 2009, we were saturating a 10G link with video-like traffic using roughly 25% of a Core 2 Duo machine. At the time, it was very difficult to find a motherboard capable of dealing with such bandwidth on the PCIe bridge! So, we were a bit ahead of the times.

高带宽网站（即千兆级别）的快速增长使得必须进一步考虑层级，授予较低层自主权，从而减少调用较高层的频率。此更新在 HAProxy 1.3.16 中发布，它确实是将 HAProxy 推到需要处理 10 Gbps 链接的位置的关键。这要归功于 TCP 拼接，当它得到网卡支持时，允许数据在两侧之间进行代理，而只使用内存中的一个副本。到 2009 年的 1.3.17，我们使用 Core 2 Duo 机器的大约 25% 的视频流量使 10G 链路饱和。当时在PCIe桥上很难找到能够处理这样带宽的主板！所以，我们有点领先于时代。

The rise of chat sites opened a new challenge: 64,000 ports weren’t enough to accommodate the number of connections to the servers. It became necessary to implement strategies to connect up to 500,000 clients. By then, Linux processes were limited to 1-million file descriptors. We added explicit [source port ranges](https://www.haproxy.com/documentation/hapee/1-9r1/onepage/#source%20(Server%20and%20default-server%20options)) and interface bindings, which allowed us to expand the number of used file descriptors by distributing connections across several servers.

聊天网站的兴起带来了新的挑战：64,000 个端口不足以容纳与服务器的连接数量。有必要实施连接多达 500,000 个客户端的策略。到那时，Linux 进程仅限于 100 万个文件描述符。我们添加了明确的 [源端口范围](https://www.haproxy.com/documentation/hapee/1-9r1/onepage/#source%20(Server%20and%20default-server%20options)) 和接口绑定，其中允许我们通过在多个服务器之间分配连接来扩展使用的文件描述符的数量。

The internal scheduler, which is responsible for running tasks such as processing requests and responses, executing ACLs, and purging data from stick tables, started to become a real problem with its simplistic linked list-based design. It was okay when all tasks had the same timeouts, which decide when the tasks should run, but was lackluster when multiple proxies (i.e. `frontend` and `listen` sections) defined different timeouts. 

内部调度器负责运行诸如处理请求和响应、执行 ACL 以及从棒表中清除数据等任务，其基于简单链表的设计开始成为一个真正的问题。当所有任务都有相同的超时时间是可以的，这决定了任务应该何时运行，但是当多个代理（即`frontend` 和 `listen` 部分）定义了不同的超时时就显得乏味了。

There had been a workaround, which was to limit the number of distinct timeouts within the configuration in order to save the scheduler from walking too far when inserting tasks. Fortunately, the parallel research on Elastic Binary Trees started to provide solid results and the time to use them for the scheduler soon came. Suddenly, it wasn’t a problem any more to have many timeouts. Watch Andjelko Iharos’ tell the story in his QCon 2019 presentation, [EBtree – Design for a Scheduler and Use (Almost) Everywhere](https://www.infoq.com/presentations/ebtree-design/).

有一个解决方法，即限制配置中不同超时的数量，以防止调度程序在插入任务时走得太远。幸运的是，对弹性二叉树的并行研究开始提供可靠的结果，很快就到了将它们用于调度程序的时候了。突然间，多次超时不再是问题。观看 Andjelko Iharos 在他的 QCon 2019 演讲中讲述的故事，[EBtree – 为调度程序设计和使用（几乎）无处不在](https://www.infoq.com/presentations/ebtree-design/)。

The scheduler was further optimized from the principle that timeouts are almost never met, so we only need to advance them in the tree, but rarely postpone them. For example, we close an inactive connection after a timeout period, but postpone that if the connection is still active. This led to a split of one queue into two: one for the timers and another for the runnable tasks. That way, the tasks are never removed from the timers list until they are truly expired. All of these updates made HAProxy capable of surpassing one million concurrent connections.

调度器从几乎永远不会遇到超时的原则进一步优化，所以我们只需要在树中推进它们，很少推迟它们。例如，我们在超时后关闭不活动的连接，但如果连接仍处于活动状态，则推迟关闭。这导致一个队列分为两个：一个用于计时器，另一个用于可运行任务。这样，任务永远不会从计时器列表中删除，直到它们真正过期。所有这些更新使 HAProxy 能够超过一百万个并发连接。

## The Keep-Alive Bottleneck

## 保持活力的瓶颈

Our next performance bottleneck was the lack of Keep-Alive. Sites had started to be loaded with many resources, such as images, scripts, and stylesheets that caused many requests; Others contacted the server often for dynamic auto-completion. For HAProxy 1.4, we focused on HTTP Keep-Alive and it turned out to be much more complicated than anticipated, since it required HAProxy to close one side of the proxied connection while keeping the other side open. This required that the _close_ event no longer being forwarded, which led to numerous bugs referred to as “CLOSE\_WAIT”.

我们的下一个性能瓶颈是缺少 Keep-Alive。站点开始加载许多资源，例如图像、脚本和样式表，这些资源引起了大量请求；其他人经常联系服务器以进行动态自动完成。对于 HAProxy 1.4，我们专注于 HTTP Keep-Alive，结果证明它比预期复杂得多，因为它需要 HAProxy 关闭代理连接的一侧，同时保持另一侧打开。这要求不再转发 _close_ 事件，这导致了许多被称为“CLOSE\_WAIT”的错误。

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-4@4x.png)

Also, keeping connections open with the client forced us to be even more careful about memory usage. At the time, IPv6 support was considered crucial and each connection had a pair of addresses—source and destination—eating 32 extra bytes per side. Even worse, supporting UNIX domain sockets pushed that overhead to 256 bytes per side. A great amount of effort was made to reduce the size of the data structures, carving off parts that weren’t always needed to save memory. HAProxy was mostly used with 8 kB static buffers, which meant that for one million connections, it required 16 GB of memory, or one buffer per direction.

此外，保持与客户端的连接打开迫使我们更加小心内存使用。当时，IPv6 支持被认为是至关重要的，每个连接都有一对地址——源和目标——每边消耗 32 个额外字节。更糟糕的是，支持 UNIX 域套接字将开销推到了每边 256 字节。为减少数据结构的大小做出了大量努力，切掉了并不总是需要节省内存的部分。 HAProxy 主要与 8 kB 静态缓冲区一起使用，这意味着对于 100 万个连接，它需要 16 GB 内存，或每个方向一个缓冲区。

HAProxy 1.4 was released with Keep-Alive only on the client side, leaving server-side Keep-Alive for 1.5. The server side required many changes, and there were a lot of other contributions arriving at the same time that competed for attention. Then, SSL was implemented, requiring an entire redesign of the connection layer. In the end, although it was estimated that server-side Keep-Alive would require six months of development, it actually took 4 ½ years. After that experience, we settled on scheduling releases at fixed dates so that it would be impossible for features to delay a release.

HAProxy 1.4 仅在客户端随 Keep-Alive 一起发布，而服务器端 Keep-Alive 则用于 1.5。服务器端需要进行许多更改，同时还有许多其他贡献争夺注意力。然后，实施了 SSL，需要对连接层进行整体重新设计。最后，虽然估计服务器端Keep-Alive需要6个月的开发时间，但实际上花了4.5年。在那次经历之后，我们决定在固定日期安排发布，这样功能就不可能延迟发布。

## SSL, Compression and Dynamic Changes

## SSL、压缩和动态变化

Although version 1.5 brought much needed features like SSL and compression, the memory usage went through the roof. In version 1.6, work was done to allocate buffers dynamically and release them as soon as possible. Meanwhile, the standard compression library, Zlib, used 256 kB of memory per connection. We decided to replace it with an in-house, stateless implementation that, although it compresses slightly less, only requires 8 bytes and is 3x faster. We also introduced server-side connection multiplexing to reduce memory usage, improve performance, and save server resources.

尽管 1.5 版带来了 SSL 和压缩等急需的功能，但内存使用量却达到了顶峰。在 1.6 版本中，已经完成了动态分配缓冲区并尽快释放它们的工作。同时，标准压缩库 Zlib 每个连接使用 256 kB 内存。我们决定用一个内部的、无状态的实现来替换它，虽然它压缩得稍微少一些，但只需要 8 个字节，速度提高了 3 倍。我们还引入了服务器端连接多路复用，以减少内存使用、提高性能并节省服务器资源。

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-5@4x.png)



Then, the code began to evolve towards more dynamic use cases. We added DNS resolution for configuring server addresses; Version 1.7 brought the ability to change almost all settings of a `server` line at runtime by using the [HAProxy Runtime API](https://www.haproxy.com/blog/dynamic-configuration-haproxy-runtime-api/) . We also introduced filters that hook into the life of a transaction and provide a way to offload heavy, content-transformation operations to external processes. This we called the [Stream Processing Offload Engine](https://www.haproxy.com/blog/extending-haproxy-with-the-stream-processing-offload-engine/)(SPOE).

然后，代码开始朝着更动态的用例发展。我们添加了用于配置服务器地址的 DNS 解析； 1.7 版通过使用 [HAProxy 运行时 API](https://www.haproxy.com/blog/dynamic-configuration-haproxy-runtime-api/) 带来了在运行时更改“服务器”行的几乎所有设置的能力.我们还引入了与事务生命周期挂钩的过滤器，并提供了一种将繁重的内容转换操作卸载到外部进程的方法。我们将其称为 [流处理卸载引擎](https://www.haproxy.com/blog/extending-haproxy-with-the-stream-processing-offload-engine/)(SPOE)。

Version 1.8 brought multithreading. We expected the first version to not be exceptionally fast, but actually it was way better than anticipated. Not only did it not lose performance (which couldn’t be taken for granted when adding threads in a two-decades old product), but it scaled up HAProxy to be ~2.5x faster on four threads. Adding multithreading broke many old assumptions that applied when you only need to initialize something once. A lot of legacy code was replaced or updated at that point.

1.8 版带来了多线程。我们预计第一个版本不会特别快，但实际上它比预期的要好得多。它不仅没有损失性能（在 20 年前的旧产品中添加线程时不能想当然），而且在四个线程上将 HAProxy 扩展到约 2.5 倍。添加多线程打破了许多旧的假设，这些假设只需要初始化一次。许多遗留代码在那时被替换或更新。

HTTP/2 to the client required support for multiplexing and that led to another redesign of the connection layers. A new _mux_ layer was introduced, which was bundled with its own set of surprises and challenges.

客户端的 HTTP/2 需要支持多路复用，这导致了连接层的另一次重新设计。引入了一个新的 _mux_ 层，它捆绑了自己的一系列惊喜和挑战。

With version 1.9, our focus was squarely on maturing features we’d added in 1.8. The scalability of the multithreading feature was improved upon by using a number of atomic operations, some lockfree operations, and other tunings. However, it was known that when the code ran in parallel, the HTTP/2 implementation, which relied on being converted to HTTP/1 before being forwarded to the server, was severely handicapped by the many HTTP header manipulations. This was especially true on the response side when the server was fast enough to send the body with the headers, thus requiring moving this body for each header adjustment.

在 1.9 版中，我们的重点完全放在了我们在 1.8 中添加的成熟功能上。通过使用许多原子操作、一些无锁操作和其他调整，多线程功能的可伸缩性得到了改进。然而，众所周知，当代码并行运行时，依赖于在转发到服务器之前转换为 HTTP/1 的 HTTP/2 实现受到许多 HTTP 标头操作的严重阻碍。当服务器足够快以发送带有标头的正文时，在响应端尤其如此，因此需要为每个标头调整移动此正文。

So, a new internal HTTP representation was created, called HTX, to allow headers to be manipulated out of order without doing expensive `memmove()`. It would also enable transport of both HTTP/1 and HTTP/2 semantics from end to end.

因此，创建了一个新的内部 HTTP 表示，称为 HTX，以允许在不执行昂贵的“memmove()”的情况下乱序操作标头。它还可以实现 HTTP/1 和 HTTP/2 语义的端到端传输。

We’d reached a turning point, leaving behind a decades-long era wherein everyone knew all of the code. Now, only a subset of people could deal with their respective, highly advanced parts. This is also a proof of excellence for certain areas, since it has become difficult to find people who can improve them.

我们已经到了一个转折点，留下了一个长达数十年的时代，每个人都知道所有代码。现在，只有一部分人可以处理他们各自的高度先进的部分。这也是某些领域卓越的证明，因为很难找到可以改进它们的人。

## A New Major Version

## 一个新的主要版本

Version 2.0 did not radically change the technical layers of 1.9, as it mostly focused on adding new features so that end users could benefit from all the work done in 1.9 that, until then, was mostly invisible…and the improvements continue. Each new feature often comes with a number of prerequisites that are adapted to existing features, until it becomes obvious that some of them hinder performance, resource usage, or even code maintainability. At that point, the code in question is redesigned to better address the shortcomings created by the initial feature.

2.0 版本并没有从根本上改变 1.9 的技术层，因为它主要专注于添加新功能，以便最终用户可以从 1.9 中完成的所有工作中受益，在此之前，这些工作大部分是不可见的……并且改进仍在继续。每个新功能通常都带有一些适用于现有功能的先决条件，直到很明显其中一些阻碍了性能、资源使用甚至代码可维护性。届时，相关代码将被重新设计，以更好地解决初始功能造成的缺点。

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-6@4x.png)

One example of this increasing level of optimization is the file-descriptor cache that was created in 1.5 as a generalization of the _speculative epoll_ polling mechanism to all pollers. It resulted in a nice simplification of the code and a performance boost on all operating systems. On the other hand, with the introduction of threads in 1.8, it has become a limiting factor. With the layer changes in 2.0, it is no longer necessary. In 2.1 (still in development) the cache has already been removed and, confirming our suspicions, resulted in a 20% higher peak connection rate in certain situations.

这种不断提高的优化级别的一个例子是在 1.5 中创建的文件描述符缓存，作为对所有轮询器的 _speculative epoll_ 轮询机制的推广。它大大简化了代码并提高了所有操作系统的性能。另一方面，随着 1.8 中线程的引入，它已成为一个限制因素。随着2.0层的变化，不再需要了。在 2.1（仍在开发中）中，缓存已经被移除，这证实了我们的怀疑，在某些情况下导致峰值连接率提高了 20%。

## Conclusion 

##  结论

From a vantage outside of the project, one might think that HAProxy is always faster than it used to be. After all, we do always say, “this is faster”. The reality is more nuanced: Each new release accelerates something and introduces a set of features that have little or no impact on the vast majority of setups. However, in Darwinian fashion, modern architecture designs gain influence and what used to be a corner case becomes the norm. Then, the next version focuses on that use case and addresses it. So, each new release of HAProxy is not exactly faster than the previous one; It’s faster than the previous one for a new use case. That’s why it’s important to stay up to date.

从项目之外的角度来看，人们可能会认为 HAProxy 总是比以前更快。毕竟，我们总是说，“这更快”。现实更加微妙：每个新版本都会加速某些事情并引入一组对绝大多数设置几乎没有影响的功能。然而，在达尔文的时尚中，现代建筑设计获得了影响，曾经的角落案例成为常态。然后，下一个版本专注于该用例并解决它。因此，HAProxy 的每个新版本并不比前一个版本快；对于新用例，它比前一个更快。这就是为什么保持最新状态很重要。

While a number of people admit that 1.4 and 1.5 are still very present in their infrastructure, even now in 2019—1.4 is unmaintained and version 1.5 was released five years ago—if their environment isn’t changing, this is fine. Yet, the web adapts very quickly and if your environment keeps up with the evolution, the benefit of having a component like HAProxy, which is always growing with the times, is hard to miss.

虽然许多人承认 1.4 和 1.5 仍然存在于他们的基础设施中，但即使是在 2019 年的现在——1.4 没有维护，5 年前发布了 1.5 版——如果他们的环境没有改变，这很好。然而，网络的适应速度非常快，如果您的环境跟上发展的步伐，那么拥有像 HAProxy 这样始终与时俱进的组件的好处是不容错过的。

What could be said is that HAProxy is always more scalable. It always adapts to faster hardware and to more demanding environments. It never quits progressing. 

可以说 HAProxy 总是更具可扩展性。它始终适应更快的硬件和更苛刻的环境。它永远不会停止前进。

