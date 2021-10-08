# Facebook outage More details about the October 4 outage

# Facebook 中断 有关 10 月 4 日中断的更多详细信息

POSTED ON OCTOBER 5, 2021 TO [Networking & Traffic](https://engineering.fb.com/category/networking-traffic/)

2021 年 10 月 5 日发布到 [网络和流量](https://engineering.fb.com/category/networking-traffic/)

![More details about the Oct. 4 Facebook outage](https://engineering.fb.com/wp-content/uploads/2021/10/datahall.jpg)

By [Santosh Janardhan](https://engineering.fb.com/author/santosh-janardhan/ "Posts by Santosh Janardhan")

作者：[Santosh Janardhan](https://engineering.fb.com/author/santosh-janardhan/“Santosh Janardhan 的帖子”)

Now that our platforms are up and running as usual after yesterday’s outage, I thought it would be worth sharing a little more detail on what happened and why — and most importantly, how we’re learning from it.

既然我们的平台在昨天的停电后照常运行，我认为有必要分享更多关于发生的事情和原因的细节——最重要的是，我们如何从中学习。

This outage was triggered by the system that manages our global backbone network capacity. The backbone is the network Facebook has built to connect all our computing facilities together, which consists of tens of thousands of miles of fiber-optic cables crossing the globe and linking all our data centers.

这次中断是由管理我们全球骨干网络容量的系统触发的。骨干网是 Facebook 为将我们所有的计算设施连接在一起而建立的网络，它由数万英里的光纤电缆组成，跨越全球并连接我们所有的数据中心。

Those data centers come in different forms. Some are massive buildings that house millions of machines that store data and run the heavy computational loads that keep our platforms running, and others are smaller facilities that connect our backbone network to the broader internet and the people using our platforms.

这些数据中心有不同的形式。有些是巨大的建筑物，里面有数百万台机器来存储数据并运行保持我们平台运行的繁重计算负载，而另一些则是将我们的骨干网络连接到更广泛的互联网和使用我们平台的人们的小型设施。

When you open one of our apps and load up your feed or messages, the app’s request for data travels from your device to the nearest facility, which then communicates directly over our backbone network to a larger data center. That’s where the information needed by your app gets retrieved and processed, and sent back over the network to your phone.

当您打开我们的一个应用程序并加载您的提要或消息时，该应用程序对数据的请求会从您的设备传输到最近的设施，然后直接通过我们的骨干网络与更大的数据中心通信。这就是您的应用程序所需的信息被检索和处理的地方，并通过网络发送回您的手机。

The data traffic between all these computing facilities is managed by routers, which figure out where to send all the incoming and outgoing data. And in the extensive day-to-day work of maintaining this infrastructure, our engineers often need to take part of the backbone offline for maintenance — perhaps repairing a fiber line, adding more capacity, or updating the software on the router itself.

所有这些计算设施之间的数据流量由路由器管理，路由器确定将所有传入和传出数据发送到何处。在维护这一基础设施的大量日常工作中，我们的工程师经常需要让部分主干离线进行维护——可能是修复光纤线路、增加更多容量或更新路由器本身的软件。

This was the source of yesterday’s outage. During one of these routine maintenance jobs, a command was issued with the intention to assess the availability of global backbone capacity, which unintentionally took down all the connections in our backbone network, effectively disconnecting Facebook data centers globally. Our systems are designed to audit commands like these to prevent mistakes like this, but a bug in that audit tool prevented it from properly stopping the command.

这就是昨天停电的根源。在其中一项日常维护工作中，发布了一项命令，旨在评估全球骨干网容量的可用性，这无意中切断了我们骨干网络中的所有连接，从而有效地断开了 Facebook 全球数据中心的连接。我们的系统旨在审核此类命令以防止出现此类错误，但该审核工具中的错误使其无法正确停止命令。

This change caused a complete disconnection of our server connections between our data centers and the internet. And that total loss of connection caused a second issue that made things worse.

这一变化导致我们的数据中心和互联网之间的服务器连接完全断开。完全失去连接导致了第二个问题，使事情变得更糟。

One of the jobs performed by our smaller facilities is to respond to DNS queries. DNS is the address book of the internet, enabling the simple web names we type into browsers to be translated into specific server IP addresses. Those translation queries are answered by our authoritative name servers that occupy well known IP addresses themselves, which in turn are advertised to the rest of the internet via another protocol called the border gateway protocol (BGP).

我们较小的设施执行的工作之一是响应 DNS 查询。 DNS 是互联网的地址簿，可以将我们在浏览器中输入的简单网络名称转换为特定的服务器 IP 地址。这些转换查询由我们的权威名称服务器回答，这些服务器本身占用众所周知的 IP 地址，这些服务器又通过另一种称为边界网关协议 (BGP) 的协议向互联网的其余部分通告。

To ensure reliable operation, our DNS servers disable those BGP advertisements if they themselves can not speak to our data centers, since this is an indication of an unhealthy network connection. In the recent outage the entire backbone was removed from operation,  making these locations declare themselves unhealthy and withdraw those BGP advertisements. The end result was that our DNS servers became unreachable even though they were still operational. This made it impossible for the rest of the internet to find our servers.

为确保可靠运行，如果我们的 DNS 服务器本身无法与我们的数据中心通信，则会禁用这些 BGP 广告，因为这表明网络连接不健康。在最近的中断中，整个主干网都停止运行，使这些位置宣布自己不健康并撤回那些 BGP 广告。最终结果是我们的 DNS 服务器无法访问，即使它们仍在运行。这使得互联网的其余部分无法找到我们的服务器。

All of this happened very fast. And as our engineers worked to figure out what was happening and why, they faced two large obstacles: first, it was not possible to access our data centers through our normal means because their networks were down, and second, the total loss of DNS broke many of the internal tools we'd normally use to investigate and resolve outages like this. 

所有这一切都发生得非常快。当我们的工程师努力弄清楚发生了什么以及为什么会发生时，他们遇到了两个巨大的障碍：首先，无法通过我们的正常方式访问我们的数据中心，因为他们的网络出现故障，其次，DNS 完全丢失了我们通常用来调查和解决此类中断的许多内部工具。

Our primary and out-of-band network access was down, so we sent engineers onsite to the data centers to have them debug the issue and restart the systems. But this took time, because these facilities are designed with high levels of physical and system security in mind. They’re hard to get into, and once you’re inside, the hardware and routers are designed to be difficult to modify even when you have physical access to them. So it took extra time to activate the secure access protocols needed to get people onsite and able to work on the servers. Only then could we confirm the issue and bring our backbone back online.

我们的主要和带外网络访问出现故障，因此我们派工程师到现场现场让他们调试问题并重新启动系统。但这需要时间，因为这些设施的设计考虑到了高水平的物理和系统安全性。它们很难进入，一旦您进入内部，即使您可以物理访问它们，硬件和路由器的设计也很难修改。因此，需要额外的时间来激活让人们到现场并能够在服务器上工作所需的安全访问协议。只有这样我们才能确认问题并使我们的主干重新上线。

Once our backbone network connectivity was restored across our data center regions, everything came back up with it. But the problem was not over — we knew that flipping our services back on all at once could potentially cause a new round of crashes due to a surge in traffic. Individual data centers were reporting dips in power usage in the range of tens of megawatts, and suddenly reversing such a dip in power consumption could put everything from electrical systems to caches at risk.

一旦我们的主干网络连接在我们的数据中心区域中恢复，一切都恢复了。但问题还没有结束——我们知道，一次将我们的服务全部重新启动可能会由于流量激增而导致新一轮的崩溃。个别数据中心报告称，电力使用量下降了数十兆瓦，突然扭转这种电力消耗下降趋势可能会使从电气系统到缓存的一切都面临风险。

Helpfully, this is an event we’re well prepared for thanks to the “storm” drills we’ve been running for a long time now. In a storm exercise, we simulate a major system failure by taking a service, data center, or entire region offline, stress testing all the infrastructure and software involved. Experience from these drills gave us the confidence and experience to bring things back online and carefully manage the increasing loads. In the end, our services came back up relatively quickly without any further systemwide failures. And while we’ve never previously run a storm that simulated our global backbone being taken offline, we’ll certainly be looking for ways to simulate events like this moving forward.

有益的是，由于我们已经进行了很长时间的“风暴”演习，我们已经为这次活动做好了充分的准备。在风暴演练中，我们通过使服务、数据中心或整个区域脱机，对所有涉及的基础设施和软件进行压力测试来模拟主要系统故障。这些演习的经验给了我们信心和经验，让我们重新上线并仔细管理不断增加的负载。最后，我们的服务相对较快地恢复了，而没有任何系统范围的进一步故障。虽然我们之前从未经历过模拟我们的全球主干离线的风暴，但我们肯定会寻找方法来模拟类似这样的事件向前发展。

Every failure like this is an opportunity to learn and get better, and there’s plenty for us to learn from this one. After every issue, small and large, we do an extensive review process to understand how we can make our systems more resilient. That process is already underway.

每一次像这样的失败都是学习和变得更好的机会，我们可以从这次失败中学到很多东西。在每一个问题之后，无论大小，我们都会进行广泛的审查，以了解如何使我们的系统更具弹性。这个过程已经在进行中。

We’ve done extensive work hardening our systems to prevent unauthorized access, and it was interesting to see how that hardening slowed us down as we tried to recover from an outage caused not by malicious activity, but an error of our own making. I believe a tradeoff like this is worth it — greatly increased day-to-day security vs. a slower recovery from a hopefully rare event like this. From here on out, our job is to strengthen our testing, drills, and overall resilience to make sure events like this happen as rarely as possible.

我们已经做了大量的工作来加固我们的系统以防止未经授权的访问，当我们试图从不是由恶意活动而是我们自己造成的错误引起的中断中恢复时，看到这种加固如何减慢我们的速度很有趣。我相信这样的权衡是值得的——大大提高了日常安全性，而不是从像这样的罕见事件中恢复得更慢。从现在开始，我们的工作是加强我们的测试、演练和整体弹性，以确保此类事件尽可能少发生。

### Read More in Networking & Traffic

### 阅读更多网络和流量

[View All](https://engineering.fb.com/category/networking-traffic/)

[查看全部](https://engineering.fb.com/category/networking-traffic/)

AUG 9, 2021 

2021 年 8 月 9 日

