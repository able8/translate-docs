# Bridge vs. Switch: What I Learned From a Data Center Tour

# Bridge vs. Switch：我从数据中心之旅中学到的东西

March 28, 2021 (Updated: August 4, 2021)

2021 年 3 月 28 日（更新：2021 年 8 月 4 日）

[Networking,](http://iximiuz.com/en/categories/?category=Networking) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

[网络](http://iximiuz.com/en/categories/?category=Networking)[Linux/Unix](http://iximiuz.com/en/categories/?category=Linux/Unix)

_Many thanks to [Victor Nagoryanskii](https://twitter.com/PA8MM) for helping me with the materials for this article._

_非常感谢 [Victor Nagoryanskii](https://twitter.com/PA8MM) 帮助我编写本文的材料。_

The difference between these two networking devices has been an unsolvable mystery to me for quite some time. For a while, I used to use the words **_"bridge"_** and **_"switch"_** interchangeably. But after getting more into networking, I started noticing that some people tend to see them as rather different devices... So, maybe I've been totally wrong? Maybe saying _"bridge aka switch"_ is way too inaccurate?

很长一段时间以来，这两种网络设备之间的区别对我来说一直是一个无法解开的谜。有一段时间，我习惯交替使用 **_"bridge"_** 和 **_"switch"_** 这两个词。但是在深入了解网络之后，我开始注意到有些人倾向于将它们视为完全不同的设备......所以，也许我完全错了？也许说 _"bridge aka switch"_ 太不准确了？

Let's try to figure it out!

让我们试着弄清楚吧！

![How network switch works](http://iximiuz.com/bridge-vs-switch/l1-network-switch-2000-opt.png)

## Switch == Bridge

## 开关 == 桥接

There is a nice book called ["Understanding Linux Network Internals"](https://www.oreilly.com/library/view/understanding-linux-network/0596002556/). It has a whole chapter on bridges. Among other things, it states that **_a bridge is the same as a switch_**.

有一本不错的书叫[“了解 Linux 网络内部原理”](https://www.oreilly.com/library/view/understanding-linux-network/0596002556/)。它有一整章关于桥梁。除其他事项外，它指出 **_a 网桥与 switch_** 相同。

![Bridges versus Switches (excerpt from Understanding Linux Network Internals)](http://iximiuz.com/bridge-vs-switch/bridge-vs-switches-quote-2000-opt.png)

_Quote from the "Understanding Linux Network Internals" book_

_引自《了解 Linux 网络内部原理》一书_

My take from this book was that:

我对这本书的看法是：

- By saying "bridge" we refer to a_logical function_ of a device.
- By saying "switch" we refer to the actual physical device performing that function.

- 通过说“桥”，我们指的是设备的_逻辑功能_。
- 通过说“开关”，我们指的是执行该功能的实际物理设备。

So, what is this _function_?

那么，这个_function_是什么？

_Bridges **transparently** combine network nodes into Layer 2 segments creating Layer 2 broadcast domains. Nodes of a single segment can exchange [data link layer](https://en.wikipedia.org/wiki/Data_link_layer) frames with each other using either unicast (MAC) or broadcast addresses._

_Bridges **透明地**将网络节点组合成第 2 层网段，创建第 2 层广播域。单个网段的节点可以使用单播 (MAC) 或广播地址相互交换 [数据链路层](https://en.wikipedia.org/wiki/Data_link_layer) 帧。_

![Bridge network function](http://iximiuz.com/bridge-vs-switch/bridge-function-2000-opt.png)

_Bridge network function._

_桥接网络功能。_

_**NB1:** Bridges can combine not just end-nodes, but also sub-segments. I.e., one can connect a bridge to a bridge, for instance, doubling the max segment size._

_**NB1:** 网桥不仅可以组合端节点，还可以组合子网段。即，可以将网桥连接到网桥，例如，将最大段大小加倍。_

_**NB2:** Bridges perform their function transparently. From the network participants' standpoint, bridges don't exist. Nodes just send frames to each other. The task of the bridge is to learn which node is reachable from which port and forward the frames accordingly._

_**NB2:** 网桥透明地执行其功能。从网络参与者的角度来看，桥梁是不存在的。节点只是互相发送帧。网桥的任务是了解从哪个端口可以到达哪个节点并相应地转发帧。_

## Switch != Bridge

## 开关 != 桥

This finding reassured me for a moment, so I felt brave enough to use the _"bridge aka switch"_ statement in [my recent write-up on low-level networking fundamentals](http://iximiuz.com/en/posts/computer-networking-101/). I was very lucky the day I published it, so the article got quite some attention and feedback. And [some of it was about my bridge-switch assumption](https://news.ycombinator.com/item?id=26577216). Apparently, not everyone was happy about it…

这一发现让我安心了片刻，所以我感到足够勇敢，可以在 [我最近关于低级网络基础知识的文章](http://iximiuz.com/en/posts) 中使用 _"bridge aka switch"_ 语句/computer-networking-101/)。发表的那天我很幸运，所以这篇文章得到了相当多的关注和反馈。 [其中一些是关于我的桥接器假设](https://news.ycombinator.com/item?id=26577216)。显然，并不是每个人都对此感到高兴……

But you know why I absolutely love the Learning in Public idea? Because of the thing that happened next. A [practicing Network Engineer](https://twitter.com/PA8MM) reached out to me and kindly offered a virtual excursion into a data center. That sounded like an awesome opportunity for me to take a look at the true enterprise-grade network and ask some questions to a real networking guru.

但是你知道为什么我绝对喜欢公共学习的想法吗？因为接下来发生的事情。一位 [执业网络工程师](https://twitter.com/PA8MM) 与我联系，并亲切地提供了一次进入数据中心的虚拟游览。对我来说，这听起来像是一个绝佳的机会，可以了解真正的企业级网络并向真正的网络专家提出一些问题。

Below are my learnings from that trip.

以下是我从那次旅行中学到的东西。

## Historical discourse

## 历史话语

Seems like there is a long history behind the evolution of networking devices. Supposedly, bridges started as two-port devices combining two L2 (or L1?) network segments into a bigger L2 segment. Hence, the name _bridge_. Then they turned into multi-port devices, but the name stayed. Regardless of the number of the ports such devices were unnoticeable from the network participant standpoint - bridges always perform their function transparently. 

似乎网络设备的发展背后有着悠久的历史。据说，网桥最初是将两个 L2（或 L1？）网段组合成一个更大的 L2 网段的双端口设备。因此，名称_bridge_。然后他们变成了多端口设备，但这个名字仍然存在。无论端口数量如何，从网络参与者的角度来看，此类设备都不会引起注意 - 网桥始终透明地执行其功能。

Then, the hardware evolved even further and nowadays, it's hard to find pure hardware bridges, especially in a serious context of use (see the next section). However, the need for the logical bridge function stayed unchanged. In the real world, this function is often performed by _switches_. In the virtual world - by _Linux bridge_ virtual device.

然后，硬件进一步发展，如今，很难找到纯粹的硬件桥接器，尤其是在严重的使用环境中（请参阅下一节）。但是，对逻辑桥接功能的需求保持不变。在现实世界中，此功能通常由 _switches_ 执行。在虚拟世界中——由_Linux桥接_虚拟设备。

## Switch >= Bridge

## 开关 >= 桥

What helped me to understand the difference is that I started thinking in terms of the following logical networking devices:

帮助我理解差异的是我开始考虑以下逻辑网络设备：

- **Bridge** \- transparent Layer 2 device performing frame forwarding on a single L2 segment.
- **Router** \- non-transparent Layer 3 device performing IP packet forwarding between multiple L3 segments.

- **桥接** \- 透明的第 2 层设备在单个 L2 段上执行帧转发。
- **路由器** \- 在多个 L3 网段之间执行 IP 数据包转发的非透明第 3 层设备。

![Bridge vs. Router](http://iximiuz.com/bridge-vs-switch/bridge-vs-router-2000-opt.png)

So, how a reasonably big LAN can be organized? Theoretically, it's possible to interconnect multiple bridges to extend a broadcast domain up to hundreds (or even thousands) of nodes.

那么，如何组织一个相当大的局域网呢？理论上，可以将多个网桥互连以将广播域扩展到数百（甚至数千）个节点。

But, it turned out that giant broadcast domains are pretty hard to manage. Apparently, in the real world, huge domains often lead to huge outages.

但是，事实证明，巨大的广播域很难管理。显然，在现实世界中，巨大的域往往会导致巨大的中断。

Thus, instead of having an immense data-center-wide broadcast domain, it's better to have smaller isolated L2 segments that can talk to each other through... L3 routers!

因此，与其拥有一个巨大的数据中心范围的广播域，不如拥有更小的隔离 L2 网段，它们可以通过…… L3 路由器相互通信！

Getting, back to switches...

得到，回到开关......

Modern network switches used in data centers are pretty advanced devices. They can work as bridges. Or, as routers. Or... some of their ports can work in a bridge mode, and some others - in a router mode. I'd speculate a bit and say that it's likely we're talking here about [_multilayer switches_](https://en.wikipedia.org/wiki/Multilayer_switch).

数据中心中使用的现代网络交换机是非常先进的设备。它们可以充当桥梁。或者，作为路由器。或者...他们的一些端口可以在桥接模式下工作，而另一些端口可以在路由器模式下工作。我推测一下，我们很可能在这里谈论的是 [_multilayer switch_](https://en.wikipedia.org/wiki/Multilayer_switch)。

In a setup I was introduced to, all servers on a single rack belong to the same _/24 IP subnet_ and are connected to a single switch.

在我介绍的设置中，单个机架上的所有服务器都属于同一个 _/24 IP 子网_并连接到单个交换机。

**_NB:_** _Facebook seems to have a pretty similar DC networking architecture. I didn't research much, but this [Load Balancer article](https://engineering.fb.com/2018/05/22/open-source/open-sourcing-katran-a-scalable-network-load-balancer/) indirectly confirms it._

**_NB:_** _Facebook 似乎有一个非常相似的 DC 网络架构。我没有研究太多，但是这篇[负载均衡器文章](https://engineering.fb.com/2018/05/22/open-source/open-sourcing-katran-a-scalable-network-load-balancer/) 间接证实了这一点。_

Such a switch is called _top-of-rack (TOR)_.

这种交换机称为 _top-of-rack (TOR)_。

For these servers, the TOR switch behaves as a canonical transparent [multi-port] bridge. I.e. as a pure Layer 2 device. And a rack forms equally sized L2 and L3 segments.

对于这些服务器，TOR 交换机充当规范的透明 [多端口] 网桥。 IE。作为纯第 2 层设备。一个机架形成相同大小的 L2 和 L3 段。

![Server rack with a top-of-rack switch](http://iximiuz.com/bridge-vs-switch/single-rack-2000-opt.png)

However, the L3 segments formed by individual racks need to be joined into internetwork. For that, one of the remaining ports of the TOR switch is configured to work in L3 mode. Unlike the L2 ports, it means that this port is an addressable network node with an IP address. It's then connected to a higher-layer switch.

但是，由单个机架形成的 L3 段需要加入到互联网络中。为此，TOR 交换机的其余端口之一配置为在 L3 模式下工作。与 L2 端口不同，这意味着该端口是具有 IP 地址的可寻址网络节点。然后它连接到更高层的交换机。

![Simplified hierarchical internetworking model](http://iximiuz.com/bridge-vs-switch/multiple-racks-2000-opt.png)

_Simplified [hierarchical internetworking model](https://en.wikipedia.org/wiki/Hierarchical_internetworking_model)._

_简化的[分层互联模型](https://en.wikipedia.org/wiki/Hierarchical_internetworking_model)._

_Disclaimer: actually, every TOR switch is connected to at least two higher-layer switches. First of all, to provide some redundancy in the case of hardware failure. But also if these multiple physical connections are combined into a single logical link, it can increase the resulting throughput. However, for the sake of simplicity, let's omit this part._

_免责声明：实际上，每个 TOR 交换机都连接到至少两个更高层的交换机。首先，在硬件故障的情况下提供一些冗余。但是，如果将这些多个物理连接组合成单个逻辑链路，也可以增加产生的吞吐量。不过，为了简单起见，我们省略这部分。_

These _higher-layer switches_ on the diagram above are called _distribution layer_ switches. In the network I was looking at, switches on this layer work as pure L3 routers. Hence, each port of a distribution-layer switch is a full-fledged Layer 3 device with an IP address assigned to it. 

上图中的这些_高层交换机_称为_分布层_交换机。在我查看的网络中，这一层上的交换机作为纯 L3 路由器工作。因此，分布层交换机的每个端口都是一个成熟的第 3 层设备，并为其分配了 IP 地址。

TOR-switches can be thought of as compound devices with a bridge (with tens of ports) and a router (with just 2 ports) inside. Distribution layer switches can be thought of as multi-port L3 routers. Like truly _multi_-port. 48- or 64-port routers! The router on a TOR-switch knows only two routes - to the rack's subnet and the default one, pointing to its distribution layer switch. And every distribution layer switch knows a lot of routes. Every rack connected to it has its own /24 IP subnet and this switch works as a border router for tens such subnets.

TOR 交换机可以被认为是内部有一个网桥（有几十个端口）和一个路由器（只有 2 个端口）的复合设备。分布层交换机可以被认为是多端口 L3 路由器。就像真正的_multi_-port。 48 或 64 端口路由器！ TOR 交换机上的路由器只知道两条路由 - 到机架的子网和默认路由，指向其分布层交换机。每个分布层交换机都知道很多路由。连接到它的每个机架都有自己的 /24 IP 子网，该交换机用作数十个此类子网的边界路由器。

**However, physically, there is no difference between the first and the second types of switches!** They all look the same, but just configured differently.

**但是，在物理上，第一种和第二种类型的交换机之间没有区别！**它们看起来都一样，只是配置不同。

In addition to these 48 (or 64) ports, switches have one or two [_out-of-band_](https://en.wikipedia.org/wiki/Out-of-band_data) ports.

除了这 48 个（或 64 个）端口之外，交换机还有一两个 [_out-of-band_](https://en.wikipedia.org/wiki/Out-of-band_data) 端口。

![Picture of a real switch (hardware)](http://iximiuz.com/bridge-vs-switch/real-switch-2000-opt.png)

Regardless of the mode of other ports, the _out-of-band_ ports always work in L3 mode. One can log in on a switch using an IP address of an _out-of-band_ port. This is needed to manage switches because configuring switches through the normal ports would be simply dangerous. Imagine, you messed up with some commands and blocked yourself out of the switch?

不管其他端口的模式如何，带外端口始终工作在 L3 模式。可以使用 _out-of-band_ 端口的 IP 地址登录交换机。这是管理交换机所必需的，因为通过普通端口配置交换机非常危险。想象一下，你搞砸了一些命令，把自己挡在了开关之外？

Curious what would it look like when you _ssh_ to a switch? Surprise, surprise! It's Linux! Or FreeBSD. Or a proprietary Unix-like OS. I.e. one can configure a switch via a traditional ssh session using widely-known _iproute2_ tools such as `ip` and/or `bridge`.

好奇当你 _ssh_ 切换到交换机时会是什么样子？惊喜，惊喜！是Linux！或 FreeBSD。或者专有的类 Unix 操作系统。 IE。可以使用广为人知的 _iproute2_ 工具（例如“ip”和/或“bridge”）通过传统的 ssh 会话配置交换机。

From the management standpoint, every port on a switch looks like a traditional network device. One can enslave some ports to a _virtual Linux bridge_, assign some other ports IP addresses, configure packet forwarding between ports, set a route table, etc. So, you can think of a switch as of a Linux server with many-many network ports . And it's totally up to you how to configure them. All the traditional Linux capabilities are at your disposal. But of course, from the hardware standpoint, switches are highly-optimized packet processing devices.

从管理的角度来看，交换机上的每个端口看起来都像一个传统的网络设备。可以将一些端口奴役到_虚拟 Linux 网桥_，分配一些其他端口的 IP 地址，配置端口之间的数据包转发，设置路由表等。因此，您可以将交换机视为具有许多网络端口的 Linux 服务器.如何配置它们完全取决于您。您可以使用所有传统的 Linux 功能。但当然，从硬件的角度来看，交换机是高度优化的数据包处理设备。

![Managed network switch](http://iximiuz.com/bridge-vs-switch/managed-switch-2000-opt.png)

So, what's up with broadcast domains? Can we have a broadcast domain spanning multiple racks? Sure! [VXLAN](http://iximiuz.com/en/posts/computer-networking-101/#VXLAN) to the rescue! Up until this point, I was describing a physical network setup of a data center. But one can configure any sort of an overlay network on top of it making it tailored for the end-users use cases.

那么，广播域是怎么回事？我们可以拥有一个跨越多个机架的广播域吗？当然！ [VXLAN](http://iximiuz.com/en/posts/computer-networking-101/#VXLAN) 来救援！到目前为止，我在描述数据中心的物理网络设置。但是可以在其上配置任何类型的覆盖网络，使其适合最终用户用例。

![VXLAN over multiple L3 networks](http://iximiuz.com/bridge-vs-switch/xvlan-2000-opt.png)

## Instead of conclusion

## 而不是结论

Actually, it was quite exciting for me to see the real production network in action. Things that I learned on that journey make a lot of sense to me. However, I'm pretty sure that it's not the only possible setup. So, I'd appreciate it if you share your experience here or [on twitter](https://twitter.com/iximiuz).

实际上，看到真实的制作网络在运行对我来说是非常令人兴奋的。我在那次旅程中学到的东西对我来说很有意义。但是，我很确定这不是唯一可能的设置。因此，如果您在这里或 [在 twitter](https://twitter.com/iximiuz) 上分享您的经验，我将不胜感激。

Have fun!

玩得开心！

## Further Reading

## 进一步阅读

- [Computer Networking Introduction - Ethernet and IP (Heavily Illustrated)](http://iximiuz.com/en/posts/computer-networking-101/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [Networking Lab - Ethernet Broadcast Domains](http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/)
- [Networking Lab - Simple VLAN](http://iximiuz.com/en/posts/networking-lab-simple-vlan/)
- [Networking Lab - L3 to L2 Segments Mapping](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/)
- [Networking Lab - Proxy ARP](http://iximiuz.com/en/posts/networking-lab-proxy-arp/)
- [Networking Lab - Simple VXLAN](http://iximiuz.com/en/posts/networking-lab-simple-vxlan/)

- [计算机网络介绍 - 以太网和 IP（大量插图）](http://iximiuz.com/en/posts/computer-networking-101/)
- [容器网络很简单！](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [网络实验室 - 以太网广播域](http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/)
- [网络实验室 - 简单 VLAN](http://iximiuz.com/en/posts/networking-lab-simple-vlan/)
- [网络实验室 - L3 到 L2 段映射](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/)
- [网络实验室 - 代理 ARP](http://iximiuz.com/en/posts/networking-lab-proxy-arp/)
- [网络实验室 - 简单 VXLAN](http://iximiuz.com/en/posts/networking-lab-simple-vxlan/)

[bridge,](javascript: void 0) [router,](javascript: void 0) [switch,](javascript: void 0) [Ethernet,](javascript: void 0) [IP](javascript: void 0)

[bridge,](javascript: void 0) [router,](javascript: void 0) [switch,](javascript: void 0) [以太网,](javascript: void 0) [IP](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

