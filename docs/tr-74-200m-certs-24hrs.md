# Preparing to Issue 200 Million Certificates in 24 Hours

#准备在24小时内发行2亿张证书

Feb 10, 2021 • Josh Aas

2021 年 2 月 10 日 • 乔什·阿斯

On a normal day Let’s Encrypt issues nearly [two million certificates](https://letsencrypt.org/stats/). When we think about what essential infrastructure for the Internet needs to be prepared for though, we’re not thinking about normal days. We want to be prepared to respond as best we can to the most difficult situations that might arise. In some of the worst scenarios, we might want to re-issue all of our certificates in a 24 hour period in order to avoid widespread disruptions. That means being prepared to issue 200 million certificates in a day, something no publicly trusted CA has ever done.

在平常的日子里，Let's Encrypt 发行近 [200 万个证书](https://letsencrypt.org/stats/)。但是，当我们考虑需要为 Internet 准备哪些基本基础设施时，我们并不是在考虑正常的日子。我们希望做好准备，尽可能应对可能出现的最困难的情况。在一些最糟糕的情况下，我们可能希望在 24 小时内重新颁发我们的所有证书，以避免大范围的中断。这意味着准备好在一天内颁发 2 亿张证书，这是公开信任的 CA 从未做过的事情。

We recently completed most of the work and investments needed to issue 200 million certificates in a day and we thought we’d let people know what was involved. All of this was made possible by our [sponsors](https://letsencrypt.org/sponsors/) and funders, including major hardware contributions from [Cisco](https://www.cisco.com/), [Thales] (https://www.thalesgroup.com/en), and [Fortinet](https://www.fortinet.com/).

我们最近完成了一天签发 2 亿张证书所需的大部分工作和投资，我们认为我们会让人们知道所涉及的内容。所有这一切都是由我们的 [赞助商](https://letsencrypt.org/sponsors/) 和资助者实现的，其中包括 [Cisco](https://www.cisco.com/)、[Thales] 的主要硬件贡献(https://www.thalesgroup.com/en) 和 [Fortinet](https://www.fortinet.com/)。

## Scenarios

##场景

Security and compliance events are a certainty in our industry. We obviously try to minimize them, but since we anticipate they will happen, we spend a lot of time preparing to respond in the best ways possible.

安全性和合规性事件在我们的行业中是确定无疑的。我们显然试图将它们最小化，但由于我们预计它们会发生，我们花了很多时间准备以尽可能最好的方式做出回应。

In February of 2020 [a bug affecting our compliance](https://community.letsencrypt.org/t/2020-02-29-caa-rechecking-bug/114591/) caused us to need to revoke and replace three million active certificates. That was approximately 2.6% of all active certificates.

2020 年 2 月 [一个影响我们合规性的错误](https://community.letsencrypt.org/t/2020-02-29-caa-rechecking-bug/114591/) 导致我们需要撤销和替换 300 万个活跃的证书。这大约占所有有效证书的 2.6%。

What if that bug had affected all of our certificates? That’s more than 150 million certificates covering more than 240 million domains. What if it had also been a more serious bug, requiring us to revoke and replace all certificates within 24 hours? That’s the kind of worst case scenario we need to be prepared for.

如果那个错误影响了我们所有的证书怎么办？这是超过 1.5 亿个证书，涵盖超过 2.4 亿个域。如果它也是一个更严重的错误，要求我们在 24 小时内撤销和替换所有证书怎么办？这是我们需要准备的最坏情况。

To make matters worse, during an incident it takes some time to evaluate the problem and make decisions, so we would be starting revocation and reissuance somewhat after the beginning of the 24-hour clock. That means we would actually have less time once the momentous decision has been made to revoke and replace 150 million or more certificates.

更糟糕的是，在事件发生期间，评估问题并做出决定需要一些时间，因此我们将在 24 小时制开始后的某个时间开始撤销和重新发布。这意味着一旦做出撤销和替换 1.5 亿或更多证书的重大决定，我们实际上将有更少的时间。

## Infrastructure Improvements

## 基础设施改进

After reviewing our systems, we determined that there were four primary bottlenecks that would prevent us from replacing 200 million certificates in a day. Database performance, internal networking speed, cryptographic signing module (HSM) performance, and bandwidth.

在审查了我们的系统后，我们确定有四个主要瓶颈会阻止我们在一天内更换 2 亿个证书。数据库性能、内部网络速度、加密签名模块 (HSM) 性能和带宽。

### Database Performance

### 数据库性能

Let’s Encrypt has a primary certificate authority database at the heart of the service we offer. This tracks the status of certificate issuance and the associated accounts. It is write heavy, with plenty of reads as well. At any given time a single database server is the writer, with some reads directed at identical machines with replicas. The single writer non-clustered strategy helps with consistency and reduces complexity, but it means that writes and some reads must operate within the performance constraints of a single machine.

Let's Encrypt 在我们提供的服务的核心有一个主要的证书颁发机构数据库。这会跟踪证书颁发和相关帐户的状态。它的写入量很大，读取量也很大。在任何给定时间，单个数据库服务器都是写入者，一些读取指向具有副本的相同机器。单写入器非集群策略有助于保持一致性并降低复杂性，但这意味着写入和某些读取必须在单台机器的性能限制内进行。

Our previous generation of database servers had dual Intel Xeon E5-2650 v4 CPUs, 24 physical cores in total. They had 1TB memory with 24 3.8TB SSDs connected via SATA in a RAID 10 configuration. These worked fine for daily issuance but would not be able to handle replacing all of our certificates in a single day.

我们的上一代数据库服务器具有双 Intel Xeon E5-2650 v4 CPU，总共 24 个物理内核。他们有 1TB 内存和 24 个 3.8TB SSD，通过 RAID 10 配置中的 SATA 连接。这些对于日常发行工作正常，但无法在一天内更换我们所有的证书。

We have replaced them with [a new generation of database server](https://letsencrypt.org/2021/01/21/next-gen-database-servers.html) from Dell featuring dual AMD EPYC 7542 CPUs, 64 physical cores in total. These machines have 2TB of faster RAM. Much faster CPUs and double the memory is great, but the really interesting thing about these machines is that the EPYC CPUs provide 128 PCIe4 lanes each. This means we could pack in 24 6.4TB NVME drives for massive I/O performance. There is no viable hardware RAID for NVME, so we’ve switched to ZFS to provide the data protection we need.

我们已将它们替换为来自戴尔的 [新一代数据库服务器](https://letsencrypt.org/2021/01/21/next-gen-database-servers.html)，具有双 AMD EPYC 7542 CPU、64 个物理内核总共。这些机器有 2TB 更快的 RAM。更快的 CPU 和双倍的内存很棒，但这些机器真正有趣的是，EPYC CPU 每个提供 128 个 PCIe4 通道。这意味着我们可以装入 24 个 6.4TB NVME 驱动器以实现大规模 I/O 性能。 NVME 没有可行的硬件 RAID，因此我们已切换到 ZFS 以提供我们需要的数据保护。

### Internal Networking

### 内部网络

Let’s Encrypt infrastructure was originally built with 1G copper networking. Between bonding multiple connections for 2G performance, and use of the very limited number of 10G ports available on our switches, it served us pretty well until 2020.

Let's Encrypt 基础设施最初是使用 1G 铜缆网络构建的。在绑定多个连接以实现 2G 性能和使用我们交换机上可用的非常有限数量的 10G 端口之间，它一直为我们服务到 2020 年。

By 2020 the volume of data we were moving around internally was too much to handle efficiently with 1G copper. Some normal operations were significantly slower than we’d like (e.g. backups, replicas), and during incidents the 1G network could cause significant delays in our response. 
到 2020 年，我们在内部移动的数据量过多，无法使用 1G 铜线进行有效处理。一些正常操作比我们希望的要慢得多（例如备份、副本），并且在发生事件时，1G 网络可能会导致我们的响应出现严重延迟。
We originally looked into upgrading to 10G, but learned that upgrading to 25G fiber wasn’t much more expensive. Cisco ended up generously donating most of the switches and equipment we needed for this upgrade, and after replacing a lot of server NICs Let’s Encrypt is now running on a 25G fiber network!

我们最初考虑升级到 10G，但了解到升级到 25G 光纤并不昂贵。思科最终慷慨地捐赠了我们此次升级所需的大部分交换机和设备，在更换了大量服务器网卡后，Let's Encrypt 现在运行在 25G 光纤网络上！

Funny story - back in 2014 Cisco had donated really nice 10G fiber switches to use in building the original Let’s Encrypt infrastructure. Our cabinets at the time had an unusually short depth though, and the 10G switches didn’t physically fit. We had to return them for physically smaller 1G switches. 1G seemed like plenty at the time. We have since moved to new cabinets with standard depth!

有趣的故事 - 早在 2014 年，思科就捐赠了非常好的 10G 光纤交换机，用于构建原始的 Let's Encrypt 基础设施。不过，我们当时的机柜深度异常短，而且 10G 交换机在物理上并不适合。我们不得不为物理上更小的 1G 交换机退回它们。 1G 在当时似乎足够了。从那以后，我们搬到了标准深度的新橱柜！

### HSM Performance

### HSM 性能

Each Let’s Encrypt data center has a pair of Luna HSMs that sign all certificates and their OCSP responses. If we want to revoke and reissue 200 million certificates, we need the Luna HSMs to perform the following cryptographic signing operations:

每个 Let's Encrypt 数据中心都有一对 Luna HSM，用于签署所有证书及其 OCSP 响应。如果我们要撤销和重新颁发 2 亿张证书，我们需要 Luna HSM 执行以下加密签名操作：

- 200 million OCSP response signatures for revocation
- 200 million certificate signatures for replacement certificates
- 200 million OCSP response signatures for the new certificates

- 2 亿个用于撤销的 OCSP 响应签名
- 用于替换证书的 2 亿个证书签名
- 新证书的 2 亿个 OCSP 响应签名

That means we need to perform 600 million cryptographic signatures in 24 hours or less, with some performance overhead to account for request clustering.

这意味着我们需要在 24 小时或更短的时间内执行 6 亿次加密签名，并有一些性能开销来解决请求集群的问题。

We need to assume that during incidents we might only have one data center online, so when we’re thinking about HSM performance we’re thinking about what we can do with just one pair. Our previous HSMs could perform approximately 1,100 signing operations each per second, 2,200 between the pair. That’s 190,080,000 signatures in 24 hours, working at full capacity. That isn’t enough.

我们需要假设在发生事件时，我们可能只有一个数据中心在线，因此当我们考虑 HSM 性能时，我们会考虑仅使用一对数据中心可以做什么。我们之前的 HSM 每秒可以执行大约 1,100 次签名操作，在对之间执行 2,200 次。这是 24 小时内 190,080,000 个签名，满负荷工作。这还不够。

In order to get to where we need to be Thales generously donated new [HSMs](https://cpl.thalesgroup.com/encryption/hardware-security-modules) with about 10x the performance - approximately 10,000 signing operations per second, 20,000 between the pair. That means we can now perform 864,000,000 signing operations in 24 hours from a single data center.

为了达到我们需要的目标，Thales 慷慨地捐赠了新的 [HSM](https://cpl.thalesgroup.com/encryption/hardware-security-modules)，性能提高了大约 10 倍——每秒大约 10,000 次签名操作，20,000对之间。这意味着我们现在可以在 24 小时内从单个数据中心执行 864,000,000 次签名操作。

### Bandwidth

### 带宽

Issuing a certificate is not particularly bandwidth intensive, but during an incident we typically use a lot more bandwidth for systems recovery and analysis. We move large volumes of logs and other files between data centers and the cloud for analysis and forensic purposes. We may need to synchronize large databases. We create copies and additional back-ups. If the connections in our data centers are slow, it slows down our response.

颁发证书并不是特别需要带宽，但在发生事件期间，我们通常会使用更多的带宽来进行系统恢复和分析。我们在数据中心和云之间移动大量日志和其他文件以进行分析和取证。我们可能需要同步大型数据库。我们创建副本和额外的备份。如果我们数据中心的连接速度很慢，就会减慢我们的响应速度。

After determining that data center connection speed could add significantly to our response time, we increased bandwidth accordingly. Fortinet helped provide hardware that helps us protect and manage these higher capacity connections.

在确定数据中心连接速度会显着增加我们的响应时间后，我们相应地增加了带宽。 Fortinet 帮助提供了帮助我们保护和管理这些更高容量连接的硬件。

## API Extension

## API 扩展

In order to get all those certificates replaced, we need an efficient and automated way to notify ACME clients that they should perform early renewal. Normally ACME clients renew their certificates when one third of their lifetime is remaining, and don’t contact our servers otherwise. We [published a draft extension to ACME](https://mailarchive.ietf.org/arch/msg/acme/b-RddSX8TdGYvO3f9c7Lzg6I2I4/) last year that describes a way for clients to regularly poll ACME servers to find out about early- renewal events. We plan to polish up that draft, implement, and collaborate with clients and large integrators to get it implemented on the client side.

为了替换所有这些证书，我们需要一种高效且自动化的方式来通知 ACME 客户他们应该提前更新。通常 ACME 客户端会在其生命周期剩余三分之一时更新其证书，否则不要联系我们的服务器。我们[发布了 ACME 的扩展草案](https://mailarchive.ietf.org/arch/msg/acme/b-RddSX8TdGYvO3f9c7Lzg6I2I4/) 去年，它描述了一种让客户端定期轮询 ACME 服务器以了解早期-更新事件。我们计划完善该草案，实施并与客户和大型集成商合作，以使其在客户端实施。

## Supporting Let’s Encrypt

## 支持让我们加密

We depend on contributions from our community of users and supporters in order to provide our services. If your company or organization would like to [sponsor](https://letsencrypt.org/become-a-sponsor/) Let’s Encrypt please email us at [sponsor@letsencrypt.org](mailto:sponsor@letsencrypt.org). We ask that you make an [individual contribution](https://letsencrypt.org/donate/) if it is within your means. 
我们依靠用户和支持者社区的贡献来提供我们的服务。如果您的公司或组织想要 [赞助](https://letsencrypt.org/become-a-sponsor/) Let's Encrypt，请发送电子邮件至 [sponsor@letsencrypt.org](mailto:sponsor@letsencrypt.org)。我们要求您在力所能及的情况下做出[个人贡献](https://letsencrypt.org/donate/)。
