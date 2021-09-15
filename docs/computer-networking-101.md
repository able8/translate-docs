# Computer Networking Introduction - Ethernet and IP (Heavily Illustrated)

## Computer Networking Basics For Developers

March 21, 2021 (Updated: August 7, 2021)

[Networking,](http://iximiuz.com/en/categories/?category=Networking) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

As a software engineer, I need to deal with networking every now and then - be it configuring a [SOHO network](https://en.wikipedia.org/wiki/Small_office/home_office), setting up [container networking](http://iximiuz.com/en/posts/container-networking-is-simple/), or troubleshooting connectivity between servers in a data center. The domain is pretty broad, and the terminology can get quite confusing quickly. This article is my layman's attempt to sort the basic things out with the minimum words and maximum drawings. The primary focus will be on the Data link layer (OSI L2) of wired networks where the [Ethernet](https://en.wikipedia.org/wiki/Ethernet) is the king nowadays. But I'll slightly touch upon its neighboring layers too.

## What is LAN?

[**LAN (Local Area Network)**](https://en.wikipedia.org/wiki/Local_area_network) \- [broadly] a computer network that interconnects computers within a **limited area** such as a residence, school, office building, or data center. A LAN is not limited to a single [IP subnetwork](http://iximiuz.com#L3-segment). Much like any [WAN](https://en.wikipedia.org/wiki/Wide_area_network), a LAN can consist of multiple IP networks communicating via routers. The main determinant of a LAN is the locality (i.e. proximity) of the participants, not the L3 topology.

## What is Network Link?

**Network link** \- a physical and logical network component used to interconnect [any kind of] nodes in the network. All the nodes of a single network link use the same [link-layer protocol](https://en.wikipedia.org/wiki/Link_layer). Examples: a bunch of computers connected to a network switch (Ethernet); a bunch of smartphones connected to a [Wi-Fi access point](https://en.wikipedia.org/wiki/Wireless_access_point) (non-Ethernet).

## What is Network Segment?

[**Network segment**](https://en.wikipedia.org/wiki/Network_segment) \- [broadly] a portion of a computer network. The actual definition of a segment is technology-specific (see below).

## What is L1 Segment?

[**L1 segment**](https://en.wikipedia.org/wiki/Network_segment#Ethernet) ( _aka_ **physical segment**, _aka_ **Ethernet segment**) \- a _network segment_ formed by an electrical (or optical) connection between networked devices using a shared medium. Nodes on a single L1 segment have a common [physical layer](https://en.wikipedia.org/wiki/Physical_layer).

In the early days of the Ethernet, a bunch of computers connected to a [**shared coaxial cable**](https://en.wikipedia.org/wiki/Ethernet_over_coax) was forming a physical segment (so-called [bus topology](https://en.wikipedia.org/wiki/Bus_network)). A coaxial cable served as a shared medium between **multiple nodes**. Everything sent by one of the nodes was seen by all other nodes of the segment. Thus, the nodes were forming a single [**broadcast domain**](http://iximiuz.com#broadcast-domain) (this is 👌). Since multiple nodes could be transmitting frames simultaneously over a single cable, _collisions_ were likely to occur. Hence, an L1 segment was forming a single [**collision domain**](http://iximiuz.com#collision-domain) (this is 👎).

![Ethernet as it started, 100 000 years ago.](http://iximiuz.com/computer-networking-101/l1-coaxial-cable-2000-opt.png)

_Ethernet as it started, 100 000 years ago._

As an evolution of Ethernet technology, [**twisted-pair cables**](https://en.wikipedia.org/wiki/Ethernet_over_twisted_pair) **connected to a common** [**repeater hub**](https://en.wikipedia.org/wiki/Ethernet_hub) replaced the shared coaxial cable (so-called [star topology](https://en.wikipedia.org/wiki/Star_network)). When a node on one of the hub's ports was transmitting frames, they were retransmitted from all the other ports of the hub. The retransmission of frames was _as-is_, i.e. no modification or filtration of frames was happening (hubs were pretty dumb devices). All the nodes connected to the hub still were forming a single L1 segment (hence, a single [**broadcast domain**](http://iximiuz.com#broadcast-domain) 👌, hence a single [**collision domain**](http://iximiuz.com#collision-domain) 👎).

![Evolution of Ethernet, 500 A.D.](http://iximiuz.com/computer-networking-101/l1-repeater-hub-2000-opt.png)

_Evolution of Ethernet, 500 A.D._

**Both coaxial and hub-based approaches are obsolete now.**

In the modern days, the star topology is prevailing. However, hubs have been replaced by more advanced [network switch devices (aka _bridges_)](https://en.wikipedia.org/wiki/Network_switch). An L1 segment de facto was reduced to a **single point-to-point link** between an end-node and a switch (or a switch and another switch). Since there are only two nodes on a physical link the potential collision domain became very small. In reality, most of the modern wiring is full-duplex, so collisions cannot simply occur at all 🎉 Curious, what happened to the [broadcast domain](http://iximiuz.com#broadcast-domain)? Then keep reading!

![Ethernet via network switch, present day.](http://iximiuz.com/computer-networking-101/l1-network-switch-2000-opt.png)

_Ethernet via network switch, present day._

**_Disclaimer: In this article, the terms switch and bridge are used interchangeably. However, modern networking hardware is slightly more complex. So, whenever you see the word "bridge" here read it as "multi-port bridge". And every time you see a "switch", assume a "Layer 2 switch" only. These two things are more or less the same. Check out [Bridge vs Switch: What I Learned From a Data Center Tour](http://iximiuz.com/en/posts/bridge-vs-switch/) for more details._**

Another example of the contemporary L1 segment is a point-to-point connection between two end-nodes via a [patch](https://en.wikipedia.org/wiki/Patch_cable) or [crossover](https://en.wikipedia.org/wiki/Ethernet_crossover_cable) cable.

![Two computers connected with a patch cable](http://iximiuz.com/computer-networking-101/l1-patch-cable-2000-opt.png)

## What is Collision Domain?

[**Collision domain**](https://en.wikipedia.org/wiki/Collision_domain) \- a network segment connected by a shared medium or through hubs where simultaneous data transmissions collide with one another. Hence, the bigger a collision domain the worse. Nowadays, collision domains are common in wireless (i.e. non-Ethernet) networks, while back in the day they were common in Ethernet networks (see [What is L1 Segment](http://iximiuz.com#L1-segment)).

In the Ethernet world, **network switches ( _aka_ bridges) form borders of collision domains**.

![Three collision domains separated by a bridge](http://iximiuz.com/computer-networking-101/collision-domain-2000-opt.png)

_Three collision domains separated by a bridge (rather dated setup)._

## What is L2 Segment?

**L2 segment** \- multiple L1 segments interconnected using a shared switch ( _aka_ bridge) _or_ (somewhat recursively) multiple L2 segments merged into a bigger L2 segment by an upper-layer switch [\*](http://iximiuz.com#remark-xlan) where **nodes can communicate with each other using their [L2 addresses (MAC)](https://en.wikipedia.org/wiki/MAC_address) or by broadcasting frames.**

![L2 segment examples](http://iximiuz.com/computer-networking-101/l2-examples-2000-opt.png)

_L2 segment examples._

\* _Other implementations of L2 segments are possible, see [VLAN](http://iximiuz.com#VLAN) and [VXLAN](http://iximiuz.com#VXLAN) sections._

![Layer 2 Ethernet frame - super simple format](http://iximiuz.com/computer-networking-101/l2-ethernet-frame-2000-opt.png)

_Layer 2 Ethernet frame - super simple format._

That's where things get interesting.

If L1 segments are about the physical connectivity of nodes, L2 segments are rather about logical connectivity. The 1:1 and 1:All addressing provided by the [Data link layer](https://en.wikipedia.org/wiki/Data_link_layer) is vital for higher-layer protocols (ARP, IP, etc) implementations. See labs in the following sections.

![Node sends frame using destination MAC address](http://iximiuz.com/computer-networking-101/l2-send-to-mac-2000-opt.png)

_Node sends frame using destination MAC address._

![Node broadcasts frame](http://iximiuz.com/computer-networking-101/l2-send-broadcast-2000-opt.png)

_Node broadcasts frame._

## What is Broadcast Domain?

[**Broadcast domain**](https://en.wikipedia.org/wiki/Broadcast_domain) \- all the nodes of a single L2 segment, i.e. the nodes that can reach each other using a broadcast L2 address ( `ff:ff:ff:ff:ff:ff`). See the IP networks section to understand how L2 broadcast domains are used by higher layers.

In the early days of the Ethernet, [collision domains](http://iximiuz.com#collision-domain) and broadcast domains were formed by physically interconnected nodes. I.e. a typical broadcast domain would consist of all the nodes of an L1 segment and such a broadcast domain would be equal to the underlying collision domain. But if the collision domain was a misfortunate byproduct of the direct interconnection of nodes, the broadcast capabilities of such interconnection came in handy. So, the historical fight with collisions didn't affect the borders of broadcast domains.

![Collision domain vs. Broadcast domain](http://iximiuz.com/computer-networking-101/broadcast-domain-l1-l2-2000-opt.png)

With the invention of _transparent bridges_, it became possible to extend broadcast domains without extending collision domains by _bridging_ multiple L1 segments using network switches. Nowadays, hierarchical topologies of interconnected switches are used to form multi-thousand hosts broadcast domains.

![Broadcast domain examples](http://iximiuz.com/computer-networking-101/broadcast-domain-l2-only-2000-opt.png)

Normally, [L3 routers](https://en.wikipedia.org/wiki/Router_(computing)) form borders of broadcast domains. However, [VLAN](http://iximiuz.com#VLAN) can be configured to split a single L2 segment into multiple non-intersecting L2 segments, hence - broadcast domains.

Check out the lab on [how to use a Linux bridge (virtual network switch) to extend broadcast domains](http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/).

## VLAN

[**VLAN**](https://en.wikipedia.org/wiki/Virtual_LAN) \- [broadly] any broadcast domain that is partitioned and isolated at the data link layer (L2). Technically speaking, VLAN is a mechanism of _tagging_ Ethernet frames of a single L2 segment with some integer IDs (so-called VIDs).

![Layer 2 Ethernet Frame VLAN tagging](http://iximiuz.com/computer-networking-101/l2-ethernet-frame-vlan-tagging-2000-opt.png)

Frames with different IDs logically belong to different networks. This creates the appearance and functionality of network traffic that is physically on a single network segment but acts as if it is split between separate network segments. VLANs can keep network applications separate despite being connected to the same physical (or virtual) network.

![Two Virtual LANs on a single bridge](http://iximiuz.com/computer-networking-101/vlan-2000-opt.png)

_Two Virtual LANs on a single bridge._

Check out the lab on [how to set up a simple VLAN using a Linux bridge](http://iximiuz.com/en/posts/networking-lab-simple-vlan/).

VLAN technology can be seen as inverse to bridging. Bridges merge multiple L2 segments (and broadcast domains) into one bigger L2 segment. VLANs split a single L2 segment (potentially formed by bridging multiple smaller L2 segments) into multiple non-intersecting L2 segments (and broadcast domains).

## What is L3 Segment?

[**L3 segment**](https://en.wikipedia.org/wiki/Network_segment#IP) \- same as [IP subnetwork](https://en.wikipedia.org/wiki/Subnetwork) (e.g. 192.168.0/24 or 172.18.0.0/16).

Notice, that up to this point we haven't been talking about IP (L3) addressing. Communication within a single L2 segment required only MAC (L2) addresses. We know, that when a node emits a frame with a certain destination MAC address, it'll be delivered by the underlying L2 networking means to the destination node. Additionally, any node can emit a broadcast frame with the destination MAC `ff:ff:ff:ff:ff:ff` and it'll be delivered to all the nodes of its L2 segment. But how a node can reach another node (on the same L3 segment) by its IP address?

#### How to send IP packet

First and foremost, IP packets are sent wrapped into Ethernet frames (assuming the Layer 2 protocol in use is Ethernet, of course). I.e. IP protocol data units (packets) are encapsulated in the Ethernet protocol data units (frames).

![IP packet inside Ethernet frame](http://iximiuz.com/computer-networking-101/ip-packet-in-eth-frame-2000-opt.png)

Thus, the task of sending an IP packet within an L3 segment boils down to sending an Ethernet frame with the IP packet inside to the L2 segment's node that owns that destination IP. Hence, the sending node needs to learn the receiving node's MAC address first. So, some sort of L3 (IP) to L2 (MAC) address translation mechanism is required. This is usually done by a Neighbor Discovery Protocol (ARP for IPv4 and NDP for IPv6) that relies on L2 broadcast capabilities.

When the IP to MAC translation is not known, the transmitting node sends a broadcast L2 frame with a query like "Who has IP 192.168.38.12?" expecting to get back a point-to-point L2 response from the owner of that IP. Such response will obviously contain the MAC address of the node possessing the requested IP. Once the destination MAC address is known for the sender node, it just wraps an IP packet into an L2 frame destined to that MAC address. **Thus, an L3 segment heavily relies on the underlying L2 segment capabilities.**

#### L3 to L2 segments relationship

There is an interesting relationship between L3 and L2 segment borders. It's pretty common to have a 1:1 mapping of L3 and L2 segments. However, technically nothing prevents us from having multiple L3 segments over a single L2 broadcast domain.

![L3 to L2 segments relationship](http://iximiuz.com/computer-networking-101/l3-to-l2-2000-opt.png)

If stricter isolation is required, VLANs can be configured to split the L2 segment into multiple non-intersecting broadcast domains.

![Two VLANs on the same switch](http://iximiuz.com/computer-networking-101/l3-over-vlans-2000-opt.png)

Interesting that in some (rather exceptional) cases, a single L3 segment can be configured over multiple L2 segments interconnected via a router. The technique is called [Proxy ARP](https://en.wikipedia.org/wiki/Proxy_ARP) and it's documented in (rather dated) [RFC 1027](https://tools.ietf.org/html/rfc1027).

See ["L3 to L2 Segments Mapping"](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/) and ["Proxy ARP"](http://iximiuz.com/en/posts/networking-lab-proxy-arp/) labs.

#### Crossing L3 borders

Communication between any two different L3 segments always requires a router.

When a node wants to send an IP packet to a node that resides in another L3 segment (IP subnet), it needs to send that packet to a gateway router. Since nodes can talk directly only with other nodes from the same L2 segment, one of the router's interfaces has to reside on the sender's L2 segment. The IP address of the router can be obtained from the routing table every node supposedly should get configured. So, the packet sending procedure is pretty the same as above, but instead of directing an Ethernet frame with the wrapped IP packet to the final destination's MAC address (which can hardly be known to the sender), the node passes it to the router. Routers are usually connected to multiple network segments. So, when a router gets such a frame, it unwraps it and resends the underlying IP packet using one of its other interfaces. **I.e. a next-hop router of every router has to reside on one of the L2 segments the router is directly connected to.**

## VXLAN

[**VXLAN**](https://en.wikipedia.org/wiki/Virtual_Extensible_LAN) \- another network virtualization technology, somewhat similar to VLAN but much more powerful.

To some extent, VLAN can be considered as an overlay network. I.e. VLAN allows one to create multiple _virtual segments on top of an existing network segment_. However, there are some significant limitations. VLAN assumes that there is already a broadcast domain underneath, so it'll split it into multiple sub-domains by tagging frames. Additionally, there cannot be more than 4096 VLANs sharing the same underlying L2 segment. There is simply 12 spare bits to encode the VLAN ID field in the Ethernet frame format.

VXLAN technology also creates virtual broadcast domains out of an existing network. So, it's also sort of an overlay networking. However, it does so in a completely different fashion. Instead of relying on the underlying L2 segment capabilities, VXLAN assumes that all the participating nodes have an L3 (i.e. IP) connectivity. On every VXLAN node, outgoing Ethernet frames are captured then, wrapped into UDP datagrams (encapsulated), and sent over an L3 network to the destination VXLAN node. On arrival, Ethernet frames are extracted from UDP packets (decapsulated) and injected into the destination's network interface. This technique is called _tunneling_. As a result, VXLAN nodes create a virtual L2 segment, hence an L2 broadcast domain.

![VXLAN frame encapsulated in UDP packet](http://iximiuz.com/computer-networking-101/vxlan-encapsulation-2000-opt.png)

Of course, nothing prevents us from putting all the VXLAN nodes in a single L3/L2 segment. So, then VXLAN would be just a way to overcome the limitation of VLAN on the number of networks per segment. However, usually, VXLAN is used over multiple interconnected L3 segments.

I'd imagine that most of the real-world VXLANs probably reside in one or few tightly connected data centers. However, since VXLAN requires only IP to IP connectivity of the participating nodes, it essentially allows one to turn arbitrary internetwork nodes into a virtualized L2 segment. While impractical, such a virtual L2 segment can be spanning multiple WANs or even a part of the Internet. Mind-blowing 🤯

![VXLAN example](http://iximiuz.com/computer-networking-101/xvlan-2000-opt.png)

From some perspective, VXLAN can be even seen as inverse to VLAN. VLAN splits a single L2 segment (and broadcast domain) into multiple non-intersecting segments that can be used then to set up multiple L3 segments. VXLAN on the contrary can combine multiple L3 segments into one [virtual] L2 segment.

Check out the lab on [how to set up a simple VXLAN](http://iximiuz.com/en/posts/networking-lab-simple-vxlan/).

## Further Reading

- [Bridge vs Switch: What I Learned From a Data Center Tour](http://iximiuz.com/en/posts/bridge-vs-switch/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [Networking Lab - Ethernet Broadcast Domains](http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/)
- [Networking Lab - Simple VLAN](http://iximiuz.com/en/posts/networking-lab-simple-vlan/)
- [Networking Lab - L3 to L2 Segments Mapping](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/)
- [Networking Lab - Proxy ARP](http://iximiuz.com/en/posts/networking-lab-proxy-arp/)
- [Networking Lab - Simple VXLAN](http://iximiuz.com/en/posts/networking-lab-simple-vxlan/)

[Ethernet,](javascript: void 0) [MAC,](javascript: void 0) [IP,](javascript: void 0) [switch,](javascript: void 0) [bridge,](javascript: void 0) [vlan,](javascript: void 0) [vxlan](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

