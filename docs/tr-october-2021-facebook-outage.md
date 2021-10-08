# Understanding How Facebook Disappeared from the Internet

# 了解 Facebook 如何从互联网上消失

2021/10/05

![img](https://blog.cloudflare.com/content/images/2021/10/Understanding-how-Facebook-disappeared-from-the-Internet-header-on-blog--1-.png)

“Facebook can't be down, can it?”, we thought, for a second.

“Facebook 不会宕机吧？”我们想了一会儿。

Today at 15:51 UTC, we opened an internal incident entitled "Facebook DNS  lookup returning SERVFAIL" because we were worried that something was  wrong with our DNS resolver [1.1.1.1](https://developers.cloudflare.com/warp-client/). But as we were about to post on our [public status](https://www.cloudflarestatus.com/) page we realized something else more serious was going on.

今天 15:51 UTC，我们打开了一个名为“Facebook DNS 查找返回 SERVFAIL”的内部事件，因为我们担心我们的 DNS 解析器 [1.1.1.1](https://developers.cloudflare.com/warp-客户/)。但是当我们即将在我们的 [public status](https://www.cloudflarestatus.com/) 页面上发帖时，我们意识到正在发生其他更严重的事情。

Social media quickly burst into flames, reporting what our engineers rapidly  confirmed too. Facebook and its affiliated services WhatsApp and  Instagram were, in fact, all down. Their DNS names stopped resolving,  and their infrastructure IPs were unreachable. It was as if someone had  "pulled the cables" from their data centers all at once and disconnected them from the Internet.

社交媒体迅速火了起来，我们的工程师也迅速证实了这一点。事实上，Facebook 及其附属服务 WhatsApp 和 Instagram 都已关闭。他们的 DNS 名称停止解析，他们的基础设施 IP 无法访问。就好像有人一下子从他们的数据中心“拔掉了电缆”并断开了它们与互联网的连接。

This wasn't a DNS issue itself, but failing DNS was the first symptom we'd seen of a larger Facebook outage.

这本身不是 DNS 问题，但 DNS 失败是我们看到的 Facebook 更大中断的第一个症状。

How's that even possible?

这怎么可能？

### Update from Facebook

### 来自 Facebook 的更新

Facebook has now [published a blog post](https://engineering.fb.com/2021/10/04/networking-traffic/outage/) giving some details of what happened internally. Externally, we saw the BGP and DNS problems outlined in this post but the problem actually  began with a configuration change that affected the entire internal  backbone. That cascaded into Facebook and other properties disappearing  and staff internal to Facebook having difficulty getting service going  again.

Facebook 现在[发布了一篇博文](https://engineering.fb.com/2021/10/04/networking-traffic/outage/) 提供了内部发生的一些细节。在外部，我们看到了这篇文章中概述的 BGP 和 DNS 问题，但问题实际上始于影响整个内部骨干网的配置更改。这导致 Facebook 和其他资产消失，Facebook 内部员工难以再次提供服务。

Facebook posted [a further blog post](https://engineering.fb.com/2021/10/05/networking-traffic/outage-details/) with a lot more detail about what happened. You can read that post for the inside view and this post for the outside view.

Facebook 发布了 [另一篇博文](https://engineering.fb.com/2021/10/05/networking-traffic/outage-details/)，详细介绍了所发生的事情。您可以阅读该帖子的内部视图和此帖子的外部视图。

Now on to what we saw from the outside.

现在来看看我们从外面看到的。

### Meet BGP

### 认识BGP

[BGP](https://www.cloudflare.com/learning/security/glossary/what-is-bgp/) stands for Border Gateway Protocol. It's a mechanism to exchange  routing information between autonomous systems (AS) on the Internet. The big routers that make the Internet work have huge, constantly updated  lists of the possible routes that can be used to deliver every network  packet to their final destinations. Without BGP, the Internet routers  wouldn't know what to do, and the Internet wouldn't work.

[BGP](https://www.cloudflare.com/learning/security/glossary/what-is-bgp/) 代表边界网关协议。它是一种在 Internet 上的自治系统 (AS) 之间交换路由信息的机制。使 Internet 工作的大型路由器拥有庞大的、不断更新的可能路由列表，可用于将每个网络数据包传送到其最终目的地。如果没有 BGP，Internet 路由器将不知道该做什么，Internet 将无法工作。

The  Internet is literally a network of networks, and it’s bound together by  BGP. BGP allows one network (say Facebook) to advertise its presence to  other networks that form the Internet. As we write Facebook is not  advertising its presence, ISPs and other networks can’t find Facebook’s  network and so it is unavailable.

互联网实际上是一个网络的网络，它由 BGP 绑定在一起。 BGP 允许一个网络（例如 Facebook）向构成 Internet 的其他网络通告其存在。正如我们写的那样，Facebook 没有宣传其存在，ISP 和其他网络无法找到 Facebook 的网络，因此它不可用。

The individual networks each  have an ASN: an Autonomous System Number. An Autonomous System (AS) is  an individual network with a unified internal routing policy. An AS can  originate prefixes (say that they control a group of IP addresses), as  well as transit prefixes (say they know how to reach specific groups of  IP addresses).

每个单独的网络都有一个 ASN：自治系统编号。自治系统 (AS) 是具有统一内部路由策略的单个网络。 AS 可以产生前缀（比如它们控制一组 IP 地址），以及传输前缀（比如它们知道如何到达特定的 IP 地址组）。

Cloudflare's ASN is [AS13335](https://www.peeringdb.com/asn/13335). Every ASN needs to announce its prefix routes to the Internet using  BGP; otherwise, no one will know how to connect and where to find us.

Cloudflare 的 ASN 是 [AS13335](https://www.peeringdb.com/asn/13335)。每个 ASN 都需要使用 BGP 向 Internet 通告其前缀路由；否则，没有人会知道如何连接以及在哪里找到我们。

Our [learning center](https://www.cloudflare.com/learning/) has a good overview of what [BGP](https://www.cloudflare.com/learning/security/glossary/what-is-bgp/) and [ASNs](https://www.cloudflare.com/learning/network-layer/what-is-an-autonomous-system/) are and how they work.

我们的 [学习中心](https://www.cloudflare.com/learning/) 对 [BGP](https://www.cloudflare.com/learning/security/glossary/what-is-bgp)有一个很好的概述/) 和 [ASN](https://www.cloudflare.com/learning/network-layer/what-is-an-autonomous-system/) 是以及它们是如何工作的。

In this simplified diagram, you can see six autonomous systems on the  Internet and two possible routes that one packet can use to go from  Start to End. AS1 → AS2 → AS3 being the fastest, and AS1 → AS6 → AS5 →  AS4 → AS3 being the slowest, but that can be used if the first fails.

在此简化图中，您可以看到 Internet 上的六个自治系统以及一个数据包可用于从 Start 到 End 的两个可能路由。 AS1 → AS2 → AS3 最快，AS1 → AS6 → AS5 → AS4 → AS3 最慢，但如果第一个失败，则可以使用。

![img](https://blog.cloudflare.com/content/images/2021/10/image5-10.png)

At 15:58 UTC we noticed that Facebook had stopped announcing the routes to their DNS prefixes. That meant that, at least, Facebook’s DNS servers  were unavailable. Because of this Cloudflare’s 1.1.1.1 DNS resolver  could no longer respond to queries asking for the IP address of  facebook.com.

在 UTC 时间 15:58，我们注意到 Facebook 已停止公布其 DNS 前缀的路由。这意味着，至少，Facebook 的 DNS 服务器不可用。因此，Cloudflare 的 1.1.1.1 DNS 解析器无法再响应请求 facebook.com 的 IP 地址的查询。

```
route-views>show ip bgp 185.89.218.0/23
% Network not in table
route-views>

route-views>show ip bgp 129.134.30.0/23
% Network not in table
route-views>
```

Meanwhile, other Facebook IP addresses remained routed but weren’t  particularly useful since without DNS Facebook and related services were effectively unavailable:

与此同时，其他 Facebook IP 地址仍然被路由但并不是特别有用，因为没有 DNS Facebook 和相关服务实际上不可用：

```
route-views>show ip bgp 129.134.30.0
BGP routing table entry for 129.134.0.0/17, version 1025798334
Paths: (24 available, best #14, table default)
  Not advertised to any peer
  Refresh Epoch 2
  3303 6453 32934
    217.192.89.50 from 217.192.89.50 (138.187.128.158)
      Origin IGP, localpref 100, valid, external
      Community: 3303:1004 3303:1006 3303:3075 6453:3000 6453:3400 6453:3402
      path 7FE1408ED9C8 RPKI State not found
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 1
route-views>
```

We keep track of all the BGP updates and announcements we see in our  global network. At our scale, the data we collect gives us a view of how the Internet is connected and where the traffic is meant to flow from  and to everywhere on the planet.

我们会跟踪我们在全球网络中看到的所有 BGP 更新和公告。在我们的规模上，我们收集的数据使我们能够了解互联网的连接方式以及流量将在地球上的任何地方流动。

A BGP UPDATE message informs a  router of any changes you’ve made to a prefix advertisement or entirely  withdraws the prefix. We can clearly see this in the number of updates  we received from Facebook when checking our time-series BGP database. Normally this chart is fairly quiet: Facebook doesn’t make a lot of  changes to its network minute to minute.

BGP UPDATE 消息会通知路由器您对前缀广告所做的任何更改或完全撤销前缀。在检查我们的时间序列 BGP 数据库时，我们可以从 Facebook 收到的更新数量中清楚地看到这一点。通常这张图表相当安静：Facebook 不会每分钟对其网络进行大量更改。

But at around 15:40 UTC we saw a peak of routing changes from Facebook. That’s when the trouble began.

但是在 UTC 时间 15:40 左右，我们看到了 Facebook 的路由更改高峰。这时候麻烦就开始了。

![img](https://blog.cloudflare.com/content/images/2021/10/image4-11.png)

If we split this view by routes announcements and withdrawals, we get an  even better idea of what happened. Routes were withdrawn, Facebook’s DNS servers went offline, and one minute after the problem occurred,  Cloudflare engineers were in a room wondering why 1.1.1.1 couldn’t  resolve facebook.com and worrying that it was somehow a fault with our  systems.

如果我们按照路线公告和撤回路线来划分这个观点，我们就能更好地了解发生了什么。路由被撤回，Facebook 的 DNS 服务器离线，问题发生一分钟后，Cloudflare 工程师在一个房间里想知道为什么 1.1.1.1 无法解析 facebook.com 并担心它在某种程度上是我们系统的故障。

![img](https://blog.cloudflare.com/content/images/2021/10/image3-9.png)

With those withdrawals, Facebook and its sites had effectively disconnected themselves from the Internet.

通过这些提款，Facebook 及其网站实际上已经与互联网断开了连接。

### DNS gets affected

### DNS 受到影响

As a direct consequence of this, DNS resolvers all over the world stopped resolving their domain names.

其直接后果是，世界各地的 DNS 解析器停止解析其域名。

```
➜  ~ dig @1.1.1.1 facebook.com
;;->>HEADER<<- opcode: QUERY, status: SERVFAIL, id: 31322
;facebook.com.IN A
➜  ~ dig @1.1.1.1 whatsapp.com
;;->>HEADER<<- opcode: QUERY, status: SERVFAIL, id: 31322
;whatsapp.com.IN A
➜  ~ dig @8.8.8.8 facebook.com
;;->>HEADER<<- opcode: QUERY, status: SERVFAIL, id: 31322
;facebook.com.IN A
➜  ~ dig @8.8.8.8 whatsapp.com
;;->>HEADER<<- opcode: QUERY, status: SERVFAIL, id: 31322
;whatsapp.com.IN A
```

This happens because DNS, like many other systems on the Internet, also has its routing mechanism. When someone types the https://facebook.com URL in the browser, the DNS resolver, responsible for translating  domain names into actual IP addresses to connect to, first checks if it  has something in its cache and uses it. If not, it tries to grab the  answer from the domain nameservers, typically hosted by the entity that owns it.

发生这种情况是因为 DNS 与 Internet 上的许多其他系统一样，也有其路由机制。当有人在浏览器中输入 https://facebook.com URL 时，DNS 解析器负责将域名转换为要连接的实际 IP 地址，首先检查它的缓存中是否有内容并使用它。如果没有，它会尝试从域名服务器中获取答案，域名服务器通常由拥有它的实体托管。

If the nameservers are unreachable or fail to respond  because of some other reason, then a SERVFAIL is returned, and the browser issues an error to the user.

如果名称服务器无法访问或由于某些其他原因无法响应，则返回 SERVFAIL，并且浏览器会向用户发出错误消息。

Again, our learning center provides a [good explanation](https://www.cloudflare.com/learning/dns/what-is-dns/) on how DNS works.

同样，我们的学习中心就 DNS 的工作原理提供了 [很好的解释](https://www.cloudflare.com/learning/dns/what-is-dns/)。

![img](https://blog.cloudflare.com/content/images/2021/10/image8-8.png)

Due to Facebook stopping announcing their DNS prefix routes through BGP,  our and everyone else's DNS resolvers had no way to connect to their  nameservers. Consequently, 1.1.1.1, 8.8.8.8, and other major public DNS  resolvers started issuing (and caching) SERVFAIL responses.

由于 Facebook 停止通过 BGP 宣布他们的 DNS 前缀路由，我们和其他所有人的 DNS 解析器无法连接到他们的名称服务器。因此，1.1.1.1、8.8.8.8 和其他主要的公共 DNS 解析器开始发出（和缓存）SERVFAIL 响应。

But  that's not all. Now human behavior and application logic kicks in and causes another exponential effect. A tsunami of additional DNS traffic follows. 

但这还不是全部。现在人类行为和应用程序逻辑开始发挥作用并导致另一个指数效应。随之而来的是额外的 DNS 流量海啸。

This happened in part because apps won't accept an error  for an answer and start retrying, sometimes aggressively, and in part because end-users also won't take an error for an answer and start  reloading the pages, or killing and relaunching their apps, sometimes  also aggressively.

发生这种情况的部分原因是应用程序不会接受答案错误并开始重试，有时是积极的，部分原因是最终用户也不会接受答案错误并开始重新加载页面，或杀死并重新启动他们的应用程序，有时也很积极。

This is the traffic increase (in number of requests) that we saw on 1.1.1.1:

这是我们在 1.1.1.1 上看到的流量增加（请求数量）：

![img](https://blog.cloudflare.com/content/images/2021/10/image6-9.png)

So now, because Facebook and their sites are so big, we have DNS resolvers worldwide handling 30x more queries than usual and potentially causing  latency and timeout issues to other platforms.

所以现在，由于 Facebook 及其网站如此之大，我们在全球拥有 DNS 解析器，可以处理比平时多 30 倍的查询，并可能导致其他平台出现延迟和超时问题。

Fortunately, 1.1.1.1 was built to be Free, Private, Fast (as the independent DNS monitor [DNSPerf](https://www.dnsperf.com/#!dns-resolvers) can attest), and scalable, and we were able to keep servicing our users with minimal impact.

幸运的是，1.1.1.1 被构建为免费、私有、快速（正如独立 DNS 监控器 [DNSPerf](https://www.dnsperf.com/#!dns-resolvers)所证明的那样)和可扩展性，我们能够以最小的影响继续为我们的用户提供服务。

The vast majority of our DNS requests kept resolving in under 10ms. At the  same time, a minimal fraction of p95 and p99 percentiles saw increased response times, probably due to expired TTLs having to resort to the  Facebook nameservers and timeout. The 10 seconds DNS timeout limit is  well known amongst engineers.

我们绝大多数 DNS 请求的解析时间都在 10 毫秒以内。同时，p95 和 p99 百分位数的一小部分响应时间增加，可能是由于过期的 TTL 不得不求助于 Facebook 名称服务器和超时。 10 秒 DNS 超时限制在工程师中是众所周知的。

![img](https://blog.cloudflare.com/content/images/2021/10/image2-11.png)

### Impacting other services

### 影响其他服务

People look for alternatives and want to know more or discuss what’s going on. When Facebook became unreachable, we started seeing increased DNS  queries to Twitter, Signal and other messaging and social media  platforms.

人们寻找替代方案并想了解更多或讨论正在发生的事情。当 Facebook 变得无法访问时，我们开始看到对 Twitter、Signal 和其他消息传递和社交媒体平台的 DNS 查询增加。

![img](https://blog.cloudflare.com/content/images/2021/10/image1-12.png)

We can also see another side effect of this unreachability in our WARP  traffic to and from Facebook's affected ASN 32934. This chart shows how  traffic changed from 15:45 UTC to 16:45 UTC compared with three hours  before in each country. All over the world WARP traffic to and from  Facebook’s network simply disappeared.

我们还可以看到这种不可达性对 Facebook 受影响的 ASN 32934 和来自受影响的 ASN 32934 的 WARP 流量造成的另一个副作用。该图表显示了每个国家/地区从 15:45 UTC 到 16:45 UTC 的流量变化情况。全世界进出 Facebook 网络的 WARP 流量都消失了。

![img](https://blog.cloudflare.com/content/images/2021/10/image7-6.png)

### The Internet

###  互联网

Today's events are a gentle reminder that the Internet is a very complex and  interdependent system of millions of systems and protocols working together. That trust, standardization, and cooperation between entities  are at the center of making it work for almost five billion active users worldwide.

今天的事件温和地提醒人们，互联网是一个非常复杂且相互依赖的系统，由数百万个系统和协议协同工作。实体之间的信任、标准化和合作是使其为全球近 50 亿活跃用户服务的核心。

### Update

###  更新

At around 21:00 UTC we saw renewed BGP activity from Facebook's network which peaked at 21:17 UTC.

在 UTC 时间 21:00 左右，我们看到来自 Facebook 网络的更新的 BGP 活动在 UTC 时间 21:17 达到峰值。

![img](https://blog.cloudflare.com/content/images/2021/10/unnamed-3-3.png)

This chart shows the availability of the DNS name 'facebook.com' on  Cloudflare's DNS resolver 1.1.1.1. It stopped being available at around  15:50 UTC and returned at 21:20 UTC.

此图表显示了 DNS 名称“facebook.com”在 Cloudflare 的 DNS 解析器 1.1.1.1 上的可用性。它在 15:50 UTC 左右停止可用，并在 21:20 UTC 返回。

![img](https://blog.cloudflare.com/content/images/2021/10/unnamed-4.png)

Undoubtedly Facebook, WhatsApp and Instagram services will take further time to  come online but as of 21:28 UTC Facebook appears to be reconnected to the global Internet and DNS working again. 

毫无疑问，Facebook、WhatsApp 和 Instagram 服务将需要更多时间才能上线，但截至 UTC 时间 21 点 28 分，Facebook 似乎已重新连接到全球互联网，DNS 再次正常工作。

