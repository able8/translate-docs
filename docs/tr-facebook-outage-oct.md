##       A Tale of DNS & BGP: The Facebook Outage, October 2021

## DNS 和 BGP 的故事：Facebook 中断，2021 年 10 月

![img](https://riskledger-website-media-uploads.s3-eu-west-1.amazonaws.com/chrome-dns-error-facebook.png)

This is a short tale. One of despair and intrigue, as we realise the  fragility of the global internet that we all so love and adore. On  Monday 4th October 2021, people stopped scrolling through Facebook, they gave up posting selfies to Instagram, they ceased texting on WhatsApp,  and Facebook employees abandoned doing any work. For all Facebook-owned  websites were down, thanks to a couple of three letter acronyms: DNS and BGP.

这是一个简短的故事。绝望和阴谋之一，因为我们意识到我们都如此热爱和崇拜的全球互联网的脆弱性。 2021 年 10 月 4 日星期一，人们停止滚动浏览 Facebook，他们放弃在 Instagram 上发自拍，他们停止在 WhatsApp 上发短信，Facebook 员工也放弃了做任何工作。由于三个字母的首字母缩写词：DNS 和 BGP，所有 Facebook 拥有的网站都关闭了。

While Facebook have not released any key details nor a postmortem,  the internet is built on a series of open source standards and protocols that allow us to remotely inspect some of the fallout.

虽然 Facebook 没有发布任何关键细节或事后分析，但互联网是建立在一系列开源标准和协议之上的，这些标准和协议允许我们远程检查一些后果。

When you type `facebook.com`, `riskledger.com`  or any other name into your web browser of choice, an invisible  background process begins. This starts with the Domain Name System, or  DNS for short. It is often said to be the phone book of the internet. It is responsible for converting long series of arcane numbers, useful  only to machines, to something memorable.

当您在您选择的网络浏览器中输入`facebook.com`、`riskledger.com` 或任何其他名称时，一个不可见的后台进程就会开始。这从域名系统开始，简称 DNS。人们常说它是互联网的电话簿。它负责将一长串仅对机器有用的神秘数字转换为令人难忘的东西。

The DNS system is contacted by your web browser asking simply, "what  is the internet address for facebook.com?". Now if one device was  answering every DNS question, it would quickly become overwhelmed, so  over the years a distributed hierarchy formed. This starts with a DNS  implementation built right into your device's operating system, to  devices called "recursers" operated by your ISP and sometimes large  organizations such as [Google](https://dns.google/) or [CloudFlare](https://1.1.1.1/), and even scaling up to central "root" devices that know the whereabouts of every .com, .co.uk and more.

您的网络浏览器会联系 DNS 系统，询问“facebook.com 的互联网地址是什么？”。现在，如果一台设备回答每个 DNS 问题，它很快就会不堪重负，因此多年来形成了分布式层次结构。首先是在您的设备操作系统中内置 DNS 实现，到由您的 ISP 运营的称为“递归器”的设备，有时是大型组织，例如 [Google](https://dns.google/) 或 [CloudFlare](https://1.1.1.1/)，甚至扩展到知道每个 .com、.co.uk 等位置的中央“根”设备。

Facebook in this case, operates a set of intermediary DNS servers  that are responsible for everything between your ISP's recursers and the roots. These are responsible for facebook.com, instagram.com,  whatsapp.com and everything else they operate. These servers are not  responding. This is what we see above from our web browser when it hints to us the error is `DNS_PROBE_FINISHED_NXDOMAIN`. Our web browser tried it's best to work out the internet address for Facebook, but it didn't get a reply.

在这种情况下，Facebook 运营着一组中间 DNS 服务器，这些服务器负责您的 ISP 的递归器和根之间的所有事情。这些负责 facebook.com、instagram.com、whatsapp.com 以及他们运营的所有其他内容。这些服务器没有响应。这是我们从上面的 Web 浏览器中看到的，它提示我们错误是“DNS_PROBE_FINISHED_NXDOMAIN”。我们的网络浏览器尽最大努力计算出 Facebook 的互联网地址，但没有得到回复。

While DNS may allow our computers to rapidly translate from names  useful to humans to numbers useful for computers, how does your device,  with it's own internet address, traverse the global internet to reach  Facebook?

虽然 DNS 可以让我们的计算机快速从对人类有用的名称转换为对计算机有用的数字，但您的设备如何通过自己的互联网地址遍历全球互联网以到达 Facebook？

No two devices on the internet are directly connected. At home, you  will have your residential ISP. In a datacentre, there will be multiple commercial ISPs. Given two internet addresses, how do the two  communicate? This is where routing comes in.

互联网上没有两台设备直接连接。在家里，您将拥有您的住宅 ISP。在一个数据中心，将有多个商业 ISP。给定两个互联网地址，两者如何通信？这就是路由的用武之地。

![multiple internet routers](https://riskledger-website-media-uploads.s3-eu-west-1.amazonaws.com/bgp-routing.png)

Routing is the system by which a path between two devices is  calculated and established, potentially traversing dozens of ISP  networks in the process. Given there must be tens-of-thousands of ISPs,  how do we ensure they all can speak to each other?

路由是计算和建立两个设备之间路径的系统，在此过程中可能会遍历数十个 ISP 网络。鉴于必须有数以万计的 ISP，我们如何确保它们都可以相互通话？

We establish a standard! In 1989 no less, the Border Gateway  Protocol, or BGP for short, was accepted by the internet community as [RFC 1105](https://datatracker.ietf.org/doc/html/rfc1105). The document laid out a protocol by which routers operated by different companies and ISPs could exchange routing information which each other.

我们建立一个标准！ 1989 年，边界网关协议（简称 BGP）被互联网社区接受为 [RFC 1105](https://datatracker.ietf.org/doc/html/rfc1105)。该文件制定了一个协议，不同公司和 ISP 运营的路由器可以通过该协议相互交换路由信息。

So what does this have to do with Facebook? We discussed earlier how  our web browsers were not receiving a response to their DNS questions,  this was however not a result of DNS itself, it was a result of  Facebook's routers ceasing to speak BGP with the rest of the internet,  and ultimately all of our ISP's routers stopped knowing where to send  our DNS requests, or any traffic to Facebook for that matter.

那么这与 Facebook 有什么关系呢？我们之前讨论过我们的 Web 浏览器如何没有收到对他们的 DNS 问题的响应，但这不是 DNS 本身的结果，而是 Facebook 的路由器停止与互联网的其余部分通信 BGP 的结果，最终我们所有的ISP 的路由器不再知道将我们的 DNS 请求或任何流量发送到 Facebook 的位置。

While we can speculate all day as to exactly what issue occurred on  the 4th of October in Facebook's infrastructure, it is however fun to  arm-chair investigate using open source information, and send [#hugops](https://www.pagerduty.com/blog/hugops-in-practice/) to the Facebook's network operations team! 

虽然我们可以一整天都在推测 10 月 4 日在 Facebook 的基础设施中究竟发生了什么问题，但使用开源信息进行调查并发送 [#hugops](https://www.pagerduty.com/blog/hugops-in-practice/) 给 Facebook 的网络运营团队！

