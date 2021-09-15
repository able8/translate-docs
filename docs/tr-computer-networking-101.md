# Computer Networking Introduction - Ethernet and IP (Heavily Illustrated)

# 计算机网络介绍 - 以太网和 IP（大量插图）

## Computer Networking Basics For Developers

## 开发人员的计算机网络基础知识

March 21, 2021 (Updated: August 7, 2021)

[Networking,](http://iximiuz.com/en/categories/?category=Networking) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)



As a software engineer, I need to deal with networking every now and then - be it configuring a [SOHO network](https://en.wikipedia.org/wiki/Small_office/home_office), setting up [container networking](http://iximiuz.com/en/posts/container-networking-is-simple/), or troubleshooting connectivity between servers in a data center. The domain is pretty broad, and the terminology can get quite confusing quickly. This article is my layman's attempt to sort the basic things out with the minimum words and maximum drawings. The primary focus will be on the Data link layer (OSI L2) of wired networks where the [Ethernet](https://en.wikipedia.org/wiki/Ethernet) is the king nowadays. But I'll slightly touch upon its neighboring layers too.

作为一名软件工程师，我需要时不时地处理网络问题 - 无论是配置 [SOHO 网络](https://en.wikipedia.org/wiki/Small_office/home_office)，设置 [容器网络](http://iximiuz.com/en/posts/container-networking-is-simple/)，或对数据中心中服务器之间的连接进行故障排除。该领域非常广泛，术语很快就会变得非常混乱。这篇文章是我门外汉的尝试，用最少的文字和最多的图来梳理基本的东西。主要关注点是有线网络的数据链路层 (OSI L2)，其中 [以太网](https://en.wikipedia.org/wiki/Ethernet) 是当今的王者。但我也会稍微触及它的相邻层。

## What is LAN?

## 什么是局域网？

[**LAN (Local Area Network)**](https://en.wikipedia.org/wiki/Local_area_network) \- [broadly] a computer network that interconnects computers within a **limited area** such as a residence , school, office building, or data center. A LAN is not limited to a single [IP subnetwork](http://iximiuz.com#L3-segment). Much like any [WAN](https://en.wikipedia.org/wiki/Wide_area_network), a LAN can consist of multiple IP networks communicating via routers. The main determinant of a LAN is the locality (i.e. proximity) of the participants, not the L3 topology.

[**LAN（局域网）**](https://en.wikipedia.org/wiki/Local_area_network) \- [广义上] 将 **有限区域** 内的计算机互连的计算机网络，例如住宅、学校、办公楼或数据中心。 LAN 不限于单个 [IP 子网](http://iximiuz.com#L3-segment)。就像任何[WAN](https://en.wikipedia.org/wiki/Wide_area_network) 一样，LAN 可以由多个通过路由器通信的 IP 网络组成。 LAN 的主要决定因素是参与者的位置（即接近度)，而不是 L3 拓扑。

## What is Network Link?

## 什么是网络链接？

**Network link** \- a physical and logical network component used to interconnect [any kind of] nodes in the network. All the nodes of a single network link use the same [link-layer protocol](https://en.wikipedia.org/wiki/Link_layer). Examples: a bunch of computers connected to a network switch (Ethernet); a bunch of smartphones connected to a [Wi-Fi access point](https://en.wikipedia.org/wiki/Wireless_access_point)(non-Ethernet).

**网络链接** \- 用于互连网络中[任何种类] 节点的物理和逻辑网络组件。单个网络链路的所有节点使用相同的[链路层协议](https://en.wikipedia.org/wiki/Link_layer)。示例：连接到网络交换机（以太网）的一堆计算机；一堆连接到 [Wi-Fi 接入点](https://en.wikipedia.org/wiki/Wireless_access_point)（非以太网)的智能手机。

## What is Network Segment?

## 什么是网段？

[**Network segment**](https://en.wikipedia.org/wiki/Network_segment) \- [broadly] a portion of a computer network. The actual definition of a segment is technology-specific (see below).

[**网络段**](https://en.wikipedia.org/wiki/Network_segment) \- [广义上] 计算机网络的一部分。细分市场的实际定义是特定于技术的（见下文)。

## What is L1 Segment?

## 什么是L1段？

[**L1 segment**](https://en.wikipedia.org/wiki/Network_segment#Ethernet) ( _aka_ **physical segment**, _aka_ **Ethernet segment**) \- a _network segment_ formed by an electrical (or optical) connection between networked devices using a shared medium. Nodes on a single L1 segment have a common [physical layer](https://en.wikipedia.org/wiki/Physical_layer).

[**L1 段**](https://en.wikipedia.org/wiki/Network_segment#Ethernet)（_aka_ **物理段**，_aka_ **以太网段**）\- 一个_网络段_由使用共享介质的联网设备之间的电（或光)连接。单个 L1 网段上的节点具有公共 [物理层](https://en.wikipedia.org/wiki/Physical_layer)。

In the early days of the Ethernet, a bunch of computers connected to a [**shared coaxial cable**](https://en.wikipedia.org/wiki/Ethernet_over_coax) was forming a physical segment (so-called [bus topology](https://en.wikipedia.org/wiki/Bus_network)). A coaxial cable served as a shared medium between **multiple nodes**. Everything sent by one of the nodes was seen by all other nodes of the segment. Thus, the nodes were forming a single [**broadcast domain**](http://iximiuz.com#broadcast-domain) (this is 👌). Since multiple nodes could be transmitting frames simultaneously over a single cable, _collisions_ were likely to occur. Hence, an L1 segment was forming a single [**collision domain**](http://iximiuz.com#collision-domain) (this is 👎).

在以太网的早期，连接到[**共享同轴电缆**](https://en.wikipedia.org/wiki/Ethernet_over_coax)的一堆计算机正在形成一个物理段（所谓的[总线拓扑](https://en.wikipedia.org/wiki/Bus_network))。同轴电缆用作**多个节点**之间的共享介质。由一个节点发送的所有内容都被该段的所有其他节点看到。因此，节点形成了一个单一的[**广播域**](http://iximiuz.com#broadcast-domain)（这是👌）。由于多个节点可以通过单根电缆同时传输帧，因此可能会发生_冲突_。因此，一个 L1 段形成了一个单一的 [**collision domain**](http://iximiuz.com#collision-domain)（这是👎)。

![Ethernet as it started, 100 000 years ago.](http://iximiuz.com/computer-networking-101/l1-coaxial-cable-2000-opt.png)

_Ethernet as it started, 100 000 years ago._ 

_10 万年前开始的以太网。_

As an evolution of Ethernet technology, [**twisted-pair cables**](https://en.wikipedia.org/wiki/Ethernet_over_twisted_pair) **connected to a common** [**repeater hub**](https://en.wikipedia.org/wiki/Ethernet_hub) replaced the shared coaxial cable (so-called [star topology](https://en.wikipedia.org/wiki/Star_network)). When a node on one of the hub's ports was transmitting frames, they were retransmitted from all the other ports of the hub. The retransmission of frames was _as-is_, i.e. no modification or filtration of frames was happening (hubs were pretty dumb devices). All the nodes connected to the hub still were forming a single L1 segment (hence, a single [**broadcast domain**](http://iximiuz.com#broadcast-domain) 👌, hence a single [**collision domain **](http://iximiuz.com#collision-domain)👎).

作为以太网技术的演进，[**双绞线**](https://en.wikipedia.org/wiki/Ethernet_over_twisted_pair) **连接到公共** [**中继集线器**](https://en.wikipedia.org/wiki/Ethernet_over_twisted_pair)取代了共享同轴电缆（所谓的[星型拓扑](https://en.wikipedia.org/wiki/Star_network）)。当集线器端口之一上的节点正在传输帧时，它们会从集线器的所有其他端口重新传输。帧的重传是 _as-is_，即没有发生帧的修改或过滤（集线器是非常愚蠢的设备)。所有连接到集线器的节点仍然形成一个单一的 L1 段（因此，一个单一的 [**broadcast domain**](http://iximiuz.com#broadcast-domain) 👌，因此一个单一的 [**collision domain] **](http://iximiuz.com#collision-domain)👎)。

![Evolution of Ethernet, 500 A.D.](http://iximiuz.com/computer-networking-101/l1-repeater-hub-2000-opt.png)

_Evolution of Ethernet, 500 A.D._

_以太网的演变，公元 500 年_

**Both coaxial and hub-based approaches are obsolete now.**

**同轴和基于集线器的方法现在都已过时。**

In the modern days, the star topology is prevailing. However, hubs have been replaced by more advanced [network switch devices (aka _bridges_)](https://en.wikipedia.org/wiki/Network_switch). An L1 segment de facto was reduced to a **single point-to-point link** between an end-node and a switch (or a switch and another switch). Since there are only two nodes on a physical link the potential collision domain became very small. In reality, most of the modern wiring is full-duplex, so collisions cannot simply occur at all 🎉 Curious, what happened to the [broadcast domain](http://iximiuz.com#broadcast-domain)? Then keep reading!

在现代，星形拓扑占主导地位。但是，集线器已被更先进的 [网络交换设备（又名 _bridges_）](https://en.wikipedia.org/wiki/Network_switch) 所取代。事实上，L1 段被简化为终端节点和交换机（或交换机和另一个交换机)之间的**单点对点链路**。由于物理链路上只有两个节点，潜在的冲突域变得非常小。在现实中，现代布线大多是全双工的，所以根本不可能发生冲突 🎉 好奇，[广播域](http://iximiuz.com#broadcast-domain) 发生了什么？然后继续阅读！

![Ethernet via network switch, present day.](http://iximiuz.com/computer-networking-101/l1-network-switch-2000-opt.png)

_Ethernet via network switch, present day._

_通过网络交换机的以太网，现在。_

**_Disclaimer: In this article, the terms switch and bridge are used interchangeably. However, modern networking hardware is slightly more complex. So, whenever you see the word "bridge" here read it as "multi-port bridge". And every time you see a "switch", assume a "Layer 2 switch" only. These two things are more or less the same. Check out [Bridge vs Switch: What I Learned From a Data Center Tour](http://iximiuz.com/en/posts/bridge-vs-switch/) for more details._**

**_免责声明：在本文中，术语开关和桥接可互换使用。然而，现代网络硬件稍微复杂一些。所以，每当你在这里看到“桥”这个词时，就把它读成“多端口桥”。并且每次您看到“开关”时，仅假设是“第 2 层开关”。这两件事或多或少是一样的。查看 [Bridge vs Switch：我从数据中心之旅中学到的东西](http://iximiuz.com/en/posts/bridge-vs-switch/) 了解更多详情。_**

Another example of the contemporary L1 segment is a point-to-point connection between two end-nodes via a [patch](https://en.wikipedia.org/wiki/Patch_cable) or [crossover](https://en.wikipedia.org/wiki/Ethernet_crossover_cable) cable.

当代 L1 段的另一个示例是通过 [patch](https://en.wikipedia.org/wiki/Patch_cable) 或 [crossover](https://en.wikipedia.org/wiki/Ethernet_crossover_cable) 电缆。

![Two computers connected with a patch cable](http://iximiuz.com/computer-networking-101/l1-patch-cable-2000-opt.png)

## What is Collision Domain?

## 什么是碰撞域？

[**Collision domain**](https://en.wikipedia.org/wiki/Collision_domain) \- a network segment connected by a shared medium or through hubs where simultaneous data transmissions collide with one another. Hence, the bigger a collision domain the worse. Nowadays, collision domains are common in wireless (i.e. non-Ethernet) networks, while back in the day they were common in Ethernet networks (see [What is L1 Segment](http://iximiuz.com#L1-segment)).

[**冲突域**](https://en.wikipedia.org/wiki/Collision_domain) \- 通过共享介质或通过集线器连接的网段，其中同时数据传输相互冲突。因此，冲突域越大越糟糕。如今，冲突域在无线（即非以太网）网络中很常见，而在过去，它们在以太网中很常见（请参阅 [什么是 L1 段](http://iximiuz.com#L1-segment))。

In the Ethernet world, **network switches ( _aka_ bridges) form borders of collision domains**.

在以太网世界中，**网络交换机（_aka_网桥）形成冲突域的边界**。

![Three collision domains separated by a bridge](http://iximiuz.com/computer-networking-101/collision-domain-2000-opt.png)

_Three collision domains separated by a bridge (rather dated setup)._

_由桥隔开的三个冲突域（相当陈旧的设置）。_

## What is L2 Segment?

## 什么是L2段？

**L2 segment** \- multiple L1 segments interconnected using a shared switch ( _aka_ bridge) _or_ (somewhat recursively) multiple L2 segments merged into a bigger L2 segment by an upper-layer switch [\*](http://iximiuz.com#remark-xlan) where **nodes can communicate with each other using their [L2 addresses (MAC)](https://en.wikipedia.org/wiki/MAC_address) or by broadcasting frames.**

**L2 段** \- 使用共享交换机（_aka_ 桥）互连的多个 L1 段 _或_（有点递归）多个 L2 段通过上层交换机合并为更大的 L2 段 [\*](http://iximiuz.com#remark-xlan)，其中**节点可以使用它们的 [L2 地址 (MAC)](https://en.wikipedia.org/wiki/MAC_address) 或通过广播帧相互通信。**

![L2 segment examples](http://iximiuz.com/computer-networking-101/l2-examples-2000-opt.png)

_L2 segment examples._

_L2 段示例。_

\* _Other implementations of L2 segments are possible, see [VLAN](http://iximiuz.com#VLAN) and [VXLAN](http://iximiuz.com#VXLAN) sections._

\* _L2 段的其他实现是可能的，请参阅 [VLAN](http://iximiuz.com#VLAN) 和 [VXLAN](http://iximiuz.com#VXLAN) 部分。_

![Layer 2 Ethernet frame - super simple format](http://iximiuz.com/computer-networking-101/l2-ethernet-frame-2000-opt.png)

_Layer 2 Ethernet frame - super simple format._

_第 2 层以太网帧 - 超级简单的格式。_

That's where things get interesting. 

这就是事情变得有趣的地方。

If L1 segments are about the physical connectivity of nodes, L2 segments are rather about logical connectivity. The 1:1 and 1:All addressing provided by the [Data link layer](https://en.wikipedia.org/wiki/Data_link_layer) is vital for higher-layer protocols (ARP, IP, etc) implementations. See labs in the following sections.

如果 L1 段是关于节点的物理连接，L2 段则是关于逻辑连接。 [数据链路层](https://en.wikipedia.org/wiki/Data_link_layer) 提供的 1:1 和 1:All 寻址对于高层协议（ARP、IP 等)实现至关重要。请参阅以下部分中的实验室。

![Node sends frame using destination MAC address](http://iximiuz.com/computer-networking-101/l2-send-to-mac-2000-opt.png)

_Node sends frame using destination MAC address._

_节点使用目标 MAC 地址发送帧。_

![Node broadcasts frame](http://iximiuz.com/computer-networking-101/l2-send-broadcast-2000-opt.png)

_Node broadcasts frame._

_节点广播帧。_

## What is Broadcast Domain?

## 什么是广播域？

[**Broadcast domain**](https://en.wikipedia.org/wiki/Broadcast_domain) \- all the nodes of a single L2 segment, ie the nodes that can reach each other using a broadcast L2 address ( `ff :ff:ff:ff:ff:ff`). See the IP networks section to understand how L2 broadcast domains are used by higher layers.

[**广播域**](https://en.wikipedia.org/wiki/Broadcast_domain) \- 单个 L2 段的所有节点，即可以使用广播 L2 地址（`ff :ff:ff:ff:ff:ff`)。请参阅 IP 网络部分以了解更高层如何使用 L2 广播域。

In the early days of the Ethernet, [collision domains](http://iximiuz.com#collision-domain) and broadcast domains were formed by physically interconnected nodes. I.e. a typical broadcast domain would consist of all the nodes of an L1 segment and such a broadcast domain would be equal to the underlying collision domain. But if the collision domain was a misfortunate byproduct of the direct interconnection of nodes, the broadcast capabilities of such interconnection came in handy. So, the historical fight with collisions didn't affect the borders of broadcast domains.

在以太网的早期，[碰撞域](http://iximiuz.com#collision-domain)和广播域由物理互连的节点组成。 IE。典型的广播域由 L1 段的所有节点组成，这样的广播域等于底层冲突域。但是，如果冲突域是节点直接互连的不幸副产品，那么这种互连的广播功能就派上用场了。因此，与冲突的历史斗争并没有影响广播域的边界。

![Collision domain vs. Broadcast domain](http://iximiuz.com/computer-networking-101/broadcast-domain-l1-l2-2000-opt.png)

With the invention of _transparent bridges_, it became possible to extend broadcast domains without extending collision domains by _bridging_ multiple L1 segments using network switches. Nowadays, hierarchical topologies of interconnected switches are used to form multi-thousand hosts broadcast domains.

随着_透明网桥_的发明，通过使用网络交换机_桥接_多个L1网段来扩展广播域而不扩展冲突域成为可能。如今，互连交换机的分层拓扑用于形成数千个主机广播域。

![Broadcast domain examples](http://iximiuz.com/computer-networking-101/broadcast-domain-l2-only-2000-opt.png)

Normally, [L3 routers](https://en.wikipedia.org/wiki/Router_(computing)) form borders of broadcast domains. However, [VLAN](http://iximiuz.com#VLAN) can be configured to split a single L2 segment into multiple non-intersecting L2 segments, hence - broadcast domains.

通常，[L3 路由器](https://en.wikipedia.org/wiki/Router_(computing))形成广播域的边界。但是，[VLAN](http://iximiuz.com#VLAN) 可以配置为将单个 L2 段拆分为多个不相交的 L2 段，因此 - 广播域。

Check out the lab on [how to use a Linux bridge (virtual network switch) to extend broadcast domains](http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/).

查看有关 [如何使用 Linux 网桥（虚拟网络交换机）扩展广播域] (http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/) 的实验室。

## VLAN

## VLAN

[**VLAN**](https://en.wikipedia.org/wiki/Virtual_LAN) \- [broadly] any broadcast domain that is partitioned and isolated at the data link layer (L2). Technically speaking, VLAN is a mechanism of _tagging_ Ethernet frames of a single L2 segment with some integer IDs (so-called VIDs).

[**VLAN**](https://en.wikipedia.org/wiki/Virtual_LAN) \- [广义上] 在数据链路层 (L2) 分区和隔离的任何广播域。从技术上讲，VLAN 是一种用一些整数 ID（所谓的 VID)标记单个 L2 网段的以太网帧的机制。

![Layer 2 Ethernet Frame VLAN tagging](http://iximiuz.com/computer-networking-101/l2-ethernet-frame-vlan-tagging-2000-opt.png)

Frames with different IDs logically belong to different networks. This creates the appearance and functionality of network traffic that is physically on a single network segment but acts as if it is split between separate network segments. VLANs can keep network applications separate despite being connected to the same physical (or virtual) network.

具有不同 ID 的帧在逻辑上属于不同的网络。这会创建物理上位于单个网段上的网络流量的外观和功能，但其行为就像在不同的网段之间拆分一样。尽管 VLAN 连接到同一个物理（或虚拟）网络，但 VLAN 可以使网络应用程序保持独立。

![Two Virtual LANs on a single bridge](http://iximiuz.com/computer-networking-101/vlan-2000-opt.png)

_Two Virtual LANs on a single bridge._

_单个网桥上的两个虚拟 LAN。_

Check out the lab on [how to set up a simple VLAN using a Linux bridge](http://iximiuz.com/en/posts/networking-lab-simple-vlan/).

查看有关[如何使用 Linux 网桥设置简单 VLAN](http://iximiuz.com/en/posts/networking-lab-simple-vlan/) 的实验室。

VLAN technology can be seen as inverse to bridging. Bridges merge multiple L2 segments (and broadcast domains) into one bigger L2 segment. VLANs split a single L2 segment (potentially formed by bridging multiple smaller L2 segments) into multiple non-intersecting L2 segments (and broadcast domains).

VLAN 技术可以看作是桥接的反面。网桥将多个 L2 段（和广播域）合并为一个更大的 L2 段。 VLAN 将单个 L2 段（可能通过桥接多个较小的 L2 段形成）拆分为多个不相交的 L2 段（和广播域）。

## What is L3 Segment?

## 什么是L3段？

[**L3 segment**](https://en.wikipedia.org/wiki/Network_segment#IP) \- same as [IP subnetwork](https://en.wikipedia.org/wiki/Subnetwork) (eg 192.168.0/24 or 172.18.0.0/16). 

[**L3 段**](https://en.wikipedia.org/wiki/Network_segment#IP) \- 与 [IP 子网](https://en.wikipedia.org/wiki/Subnetwork) 相同（例如192.168.0/24 或 172.18.0.0/16)。

Notice, that up to this point we haven't been talking about IP (L3) addressing. Communication within a single L2 segment required only MAC (L2) addresses. We know, that when a node emits a frame with a certain destination MAC address, it'll be delivered by the underlying L2 networking means to the destination node. Additionally, any node can emit a broadcast frame with the destination MAC `ff:ff:ff:ff:ff:ff` and it'll be delivered to all the nodes of its L2 segment. But how a node can reach another node (on the same L3 segment) by its IP address?

请注意，到目前为止，我们还没有讨论 IP (L3) 寻址。单个 L2 网段内的通信仅需要 MAC (L2) 地址。我们知道，当节点发出具有特定目标 MAC 地址的帧时，它将通过底层 L2 网络方式传递到目标节点。此外，任何节点都可以使用目标 MAC `ff:ff:ff:ff:ff:ff` 发出广播帧，并将其传送到其 L2 段的所有节点。但是一个节点如何通过其 IP 地址到达另一个节点（在同一 L3 网段上）？

#### How to send IP packet

#### 如何发送IP数据包

First and foremost, IP packets are sent wrapped into Ethernet frames (assuming the Layer 2 protocol in use is Ethernet, of course). I.e. IP protocol data units (packets) are encapsulated in the Ethernet protocol data units (frames).

首先，IP 数据包被包裹在以太网帧中发送（当然，假设使用的第 2 层协议是以太网）。 IE。 IP 协议数据单元（数据包）被封装在以太网协议数据单元（帧）中。

![IP packet inside Ethernet frame](http://iximiuz.com/computer-networking-101/ip-packet-in-eth-frame-2000-opt.png)

Thus, the task of sending an IP packet within an L3 segment boils down to sending an Ethernet frame with the IP packet inside to the L2 segment's node that owns that destination IP. Hence, the sending node needs to learn the receiving node's MAC address first. So, some sort of L3 (IP) to L2 (MAC) address translation mechanism is required. This is usually done by a Neighbor Discovery Protocol (ARP for IPv4 and NDP for IPv6) that relies on L2 broadcast capabilities.

因此，在 L3 网段内发送 IP 数据包的任务归结为将带有 IP 数据包的以太网帧发送到拥有该目标 IP 的 L2 网段节点。因此，发送节点需要首先了解接收节点的 MAC 地址。因此，需要某种 L3 (IP) 到 L2 (MAC) 地址转换机制。这通常由依赖 L2 广播功能的邻居发现协议（用于 IPv4 的 ARP 和用于 IPv6 的 NDP）完成。

When the IP to MAC translation is not known, the transmitting node sends a broadcast L2 frame with a query like "Who has IP 192.168.38.12?" expecting to get back a point-to-point L2 response from the owner of that IP. Such response will obviously contain the MAC address of the node possessing the requested IP. Once the destination MAC address is known for the sender node, it just wraps an IP packet into an L2 frame destined to that MAC address. **Thus, an L3 segment heavily relies on the underlying L2 segment capabilities.**

当 IP 到 MAC 的转换未知时，发送节点会发送一个广播 L2 帧，其中包含类似“谁拥有 IP 192.168.38.12？”的查询。希望从该 IP 的所有者那里得到点对点的 L2 响应。此类响应显然将包含拥有所请求 IP 的节点的 MAC 地址。一旦知道发送方节点的目标 MAC 地址，它只需将 IP 数据包包装到以该 MAC 地址为目标的 L2 帧中。 **因此，L3 段严重依赖底层 L2 段功能。**

#### L3 to L2 segments relationship

#### L3 到 L2 段关系

There is an interesting relationship between L3 and L2 segment borders. It's pretty common to have a 1:1 mapping of L3 and L2 segments. However, technically nothing prevents us from having multiple L3 segments over a single L2 broadcast domain.

L3 和 L2 段边界之间存在有趣的关系。 L3 和 L2 段的 1:1 映射是很常见的。然而，从技术上讲，没有什么能阻止我们在单个 L2 广播域上拥有多个 L3 段。

![L3 to L2 segments relationship](http://iximiuz.com/computer-networking-101/l3-to-l2-2000-opt.png)

If stricter isolation is required, VLANs can be configured to split the L2 segment into multiple non-intersecting broadcast domains.

如果需要更严格的隔离，可以配置 VLAN 将 L2 段拆分为多个不相交的广播域。

![Two VLANs on the same switch](http://iximiuz.com/computer-networking-101/l3-over-vlans-2000-opt.png)

Interesting that in some (rather exceptional) cases, a single L3 segment can be configured over multiple L2 segments interconnected via a router. The technique is called [Proxy ARP](https://en.wikipedia.org/wiki/Proxy_ARP) and it's documented in (rather dated) [RFC 1027](https://tools.ietf.org/html/rfc1027) .

有趣的是，在某些（相当特殊的）情况下，可以在通过路由器互连的多个 L2 段上配置单个 L3 段。该技术称为 [代理 ARP](https://en.wikipedia.org/wiki/Proxy_ARP)，它记录在（相当过时)[RFC1027](https://tools.ietf.org/html/rfc1027) .

See ["L3 to L2 Segments Mapping"](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/) and ["Proxy ARP"](http://iximiuz.com/en/posts/networking-lab-proxy-arp/) labs.

参见 ["L3 到 L2 段映射"](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/) 和 ["代理 ARP"](http://iximiuz.com/en/posts/networking-lab-proxy-arp/) 实验室。

#### Crossing L3 borders

#### 跨越 L3 边界

Communication between any two different L3 segments always requires a router. 

任何两个不同的 L3 网段之间的通信始终需要路由器。

When a node wants to send an IP packet to a node that resides in another L3 segment (IP subnet), it needs to send that packet to a gateway router. Since nodes can talk directly only with other nodes from the same L2 segment, one of the router's interfaces has to reside on the sender's L2 segment. The IP address of the router can be obtained from the routing table every node supposedly should get configured. So, the packet sending procedure is pretty the same as above, but instead of directing an Ethernet frame with the wrapped IP packet to the final destination's MAC address (which can hardly be known to the sender), the node passes it to the router. Routers are usually connected to multiple network segments. So, when a router gets such a frame, it unwraps it and resends the underlying IP packet using one of its other interfaces. **I.e. a next-hop router of every router has to reside on one of the L2 segments the router is directly connected to.**

当一个节点想要向位于另一个 L3 网段（IP 子网）中的节点发送 IP 数据包时，它需要将该数据包发送到网关路由器。由于节点只能与来自同一 L2 段的其他节点直接对话，因此路由器的接口之一必须驻留在发送方的 L2 段上。路由器的 IP 地址可以从路由表中获得，每个节点都应该配置。因此，数据包发送过程与上面的非常相似，但节点不是将带有包装的 IP 数据包的以太网帧定向到最终目的地的 MAC 地址（发送者几乎不知道），而是将其传递给路由器。路由器通常连接到多个网段。因此，当路由器获得这样的帧时，它会解开它并使用其其他接口之一重新发送底层 IP 数据包。 **IE。每个路由器的下一跳路由器必须驻留在路由器直接连接到的 L2 网段之一上。**

## VXLAN

## 虚拟局域网

[**VXLAN**](https://en.wikipedia.org/wiki/Virtual_Extensible_LAN) \- another network virtualization technology, somewhat similar to VLAN but much more powerful.

[**VXLAN**](https://en.wikipedia.org/wiki/Virtual_Extensible_LAN) \- 另一种网络虚拟化技术，有点类似于 VLAN，但功能更强大。

To some extent, VLAN can be considered as an overlay network. I.e. VLAN allows one to create multiple _virtual segments on top of an existing network segment_. However, there are some significant limitations. VLAN assumes that there is already a broadcast domain underneath, so it'll split it into multiple sub-domains by tagging frames. Additionally, there cannot be more than 4096 VLANs sharing the same underlying L2 segment. There is simply 12 spare bits to encode the VLAN ID field in the Ethernet frame format.

在某种程度上，VLAN 可以被认为是一个覆盖网络。 IE。 VLAN 允许在现有网段之上创建多个_虚拟网段_。但是，存在一些重大限制。 VLAN 假设下面已经有一个广播域，因此它会通过标记帧将其拆分为多个子域。此外，共享同一底层 L2 网段的 VLAN 不能超过 4096 个。只有 12 个备用位用于对以太网帧格式中的 VLAN ID 字段进行编码。

VXLAN technology also creates virtual broadcast domains out of an existing network. So, it's also sort of an overlay networking. However, it does so in a completely different fashion. Instead of relying on the underlying L2 segment capabilities, VXLAN assumes that all the participating nodes have an L3 (i.e. IP) connectivity. On every VXLAN node, outgoing Ethernet frames are captured then, wrapped into UDP datagrams (encapsulated), and sent over an L3 network to the destination VXLAN node. On arrival, Ethernet frames are extracted from UDP packets (decapsulated) and injected into the destination's network interface. This technique is called _tunneling_. As a result, VXLAN nodes create a virtual L2 segment, hence an L2 broadcast domain.

VXLAN 技术还可以在现有网络之外创建虚拟广播域。所以，它也是一种覆盖网络。然而，它以完全不同的方式这样做。 VXLAN 不依赖底层 L2 网段功能，而是假设所有参与节点都具有 L3（即 IP）连接。在每个 VXLAN 节点上，传出的以太网帧被捕获，然后包装成 UDP 数据报（封装），并通过 L3 网络发送到目标 VXLAN 节点。到达时，以太网帧从 UDP 数据包中提取（解封装）并注入目的地的网络接口。这种技术称为_隧道_。因此，VXLAN 节点创建了一个虚拟的 L2 段，因此是一个 L2 广播域。

![VXLAN frame encapsulated in UDP packet](http://iximiuz.com/computer-networking-101/vxlan-encapsulation-2000-opt.png)

Of course, nothing prevents us from putting all the VXLAN nodes in a single L3/L2 segment. So, then VXLAN would be just a way to overcome the limitation of VLAN on the number of networks per segment. However, usually, VXLAN is used over multiple interconnected L3 segments.

当然，没有什么能阻止我们将所有 VXLAN 节点放在一个 L3/L2 网段中。因此，VXLAN 将只是克服 VLAN 对每段网络数量限制的一种方法。但是，VXLAN 通常用于多个互连的 L3 网段。

I'd imagine that most of the real-world VXLANs probably reside in one or few tightly connected data centers. However, since VXLAN requires only IP to IP connectivity of the participating nodes, it essentially allows one to turn arbitrary internetwork nodes into a virtualized L2 segment. While impractical, such a virtual L2 segment can be spanning multiple WANs or even a part of the Internet. Mind-blowing 🤯

我想大多数现实世界的 VXLAN 可能驻留在一个或几个紧密连接的数据中心。然而，由于 VXLAN 只需要参与节点的 IP 到 IP 连接，它本质上允许将任意网络节点变成虚拟化的 L2 网段。虽然不切实际，但这样的虚拟 L2 网段可以跨越多个 WAN，甚至是 Internet 的一部分。心动🤯

![VXLAN example](http://iximiuz.com/computer-networking-101/xvlan-2000-opt.png)

From some perspective, VXLAN can be even seen as inverse to VLAN. VLAN splits a single L2 segment (and broadcast domain) into multiple non-intersecting segments that can be used then to set up multiple L3 segments. VXLAN on the contrary can combine multiple L3 segments into one [virtual] L2 segment.

从某种角度来看，VXLAN 甚至可以看作是 VLAN 的反面。 VLAN 将单个 L2 段（和广播域）拆分为多个不相交的段，然后可以使用这些段来设置多个 L3 段。相反，VXLAN 可以将多个 L3 段组合成一个 [虚拟] L2 段。

Check out the lab on [how to set up a simple VXLAN](http://iximiuz.com/en/posts/networking-lab-simple-vxlan/).

查看有关 [如何设置简单的 VXLAN] (http://iximiuz.com/en/posts/networking-lab-simple-vxlan/) 的实验室。

## Further Reading

## 进一步阅读

- [Bridge vs Switch: What I Learned From a Data Center Tour](http://iximiuz.com/en/posts/bridge-vs-switch/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [Networking Lab - Ethernet Broadcast Domains](http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/)
- [Networking Lab - Simple VLAN](http://iximiuz.com/en/posts/networking-lab-simple-vlan/)
- [Networking Lab - L3 to L2 Segments Mapping](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/)
- [Networking Lab - Proxy ARP](http://iximiuz.com/en/posts/networking-lab-proxy-arp/)
- [Networking Lab - Simple VXLAN](http://iximiuz.com/en/posts/networking-lab-simple-vxlan/) 

- [Bridge vs Switch：我从数据中心之旅中学到的东西](http://iximiuz.com/en/posts/bridge-vs-switch/)
- [容器网络很简单！](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [网络实验室 - 以太网广播域](http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/)
- [网络实验室 - 简单 VLAN](http://iximiuz.com/en/posts/networking-lab-simple-vlan/)
- [网络实验室 - L3 到 L2 段映射](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/)
- [网络实验室 - 代理 ARP](http://iximiuz.com/en/posts/networking-lab-proxy-arp/)
- [网络实验室 - 简单 VXLAN](http://iximiuz.com/en/posts/networking-lab-simple-vxlan/)

[Ethernet,](javascript: void 0) [MAC,](javascript: void 0) [IP,](javascript: void 0) [switch,](javascript: void 0) [bridge,](javascript: void 0) [vlan,](javascript: void 0) [vxlan](javascript: void 0)

[以太网,](javascript: void 0) [MAC,](javascript: void 0) [IP,](javascript: void 0) [switch,](javascript: void 0) [bridge,](javascript: void 0) [vlan,](javascript: void 0) [vxlan](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

