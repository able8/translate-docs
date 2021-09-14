# Bridge vs. Switch: What I Learned From a Data Center Tour

March 28, 2021 (Updated: August 4, 2021)

[Networking,](http://iximiuz.com/en/categories/?category=Networking) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

_Many thanks to [Victor Nagoryanskii](https://twitter.com/PA8MM) for helping me with the materials for this article._

The difference between these two networking devices has been an unsolvable mystery to me for quite some time. For a while, I used to use the words **_"bridge"_** and **_"switch"_** interchangeably. But after getting more into networking, I started noticing that some people tend to see them as rather different devices... So, maybe I've been totally wrong? Maybe saying _"bridge aka switch"_ is way too inaccurate?

Let's try to figure it out!

![How network switch works](http://iximiuz.com/bridge-vs-switch/l1-network-switch-2000-opt.png)

## Switch == Bridge

There is a nice book called ["Understanding Linux Network Internals"](https://www.oreilly.com/library/view/understanding-linux-network/0596002556/). It has a whole chapter on bridges. Among other things, it states that **_a bridge is the same as a switch_**.

![Bridges versus Switches (excerpt from Understanding Linux Network Internals)](http://iximiuz.com/bridge-vs-switch/bridge-vs-switches-quote-2000-opt.png)

_Quote from the "Understanding Linux Network Internals" book_

My take from this book was that:

- By saying "bridge" we refer to a_logical function_ of a device.
- By saying "switch" we refer to the actual physical device performing that function.

So, what is this _function_?

_Bridges **transparently** combine network nodes into Layer 2 segments creating Layer 2 broadcast domains. Nodes of a single segment can exchange [data link layer](https://en.wikipedia.org/wiki/Data_link_layer) frames with each other using either unicast (MAC) or broadcast addresses._

![Bridge network function](http://iximiuz.com/bridge-vs-switch/bridge-function-2000-opt.png)

_Bridge network function._

_**NB1:** Bridges can combine not just end-nodes, but also sub-segments. I.e., one can connect a bridge to a bridge, for instance, doubling the max segment size._

_**NB2:** Bridges perform their function transparently. From the network participants' standpoint, bridges don't exist. Nodes just send frames to each other. The task of the bridge is to learn which node is reachable from which port and forward the frames accordingly._

## Switch != Bridge

This finding reassured me for a moment, so I felt brave enough to use the _"bridge aka switch"_ statement in [my recent write-up on low-level networking fundamentals](http://iximiuz.com/en/posts/computer-networking-101/). I was very lucky the day I published it, so the article got quite some attention and feedback. And [some of it was about my bridge-switch assumption](https://news.ycombinator.com/item?id=26577216). Apparently, not everyone was happy about it…

But you know why I absolutely love the Learning in Public idea? Because of the thing that happened next. A [practicing Network Engineer](https://twitter.com/PA8MM) reached out to me and kindly offered a virtual excursion into a data center. That sounded like an awesome opportunity for me to take a look at the true enterprise-grade network and ask some questions to a real networking guru.

Below are my learnings from that trip.

## Historical discourse

Seems like there is a long history behind the evolution of networking devices. Supposedly, bridges started as two-port devices combining two L2 (or L1?) network segments into a bigger L2 segment. Hence, the name _bridge_. Then they turned into multi-port devices, but the name stayed. Regardless of the number of the ports such devices were unnoticeable from the network participant standpoint - bridges always perform their function transparently.

Then, the hardware evolved even further and nowadays, it's hard to find pure hardware bridges, especially in a serious context of use (see the next section). However, the need for the logical bridge function stayed unchanged. In the real world, this function is often performed by _switches_. In the virtual world - by _Linux bridge_ virtual device.

## Switch >= Bridge

What helped me to understand the difference is that I started thinking in terms of the following logical networking devices:

- **Bridge** \- transparent Layer 2 device performing frame forwarding on a single L2 segment.
- **Router** \- non-transparent Layer 3 device performing IP packet forwarding between multiple L3 segments.

![Bridge vs. Router](http://iximiuz.com/bridge-vs-switch/bridge-vs-router-2000-opt.png)

So, how a reasonably big LAN can be organized? Theoretically, it's possible to interconnect multiple bridges to extend a broadcast domain up to hundreds (or even thousands) of nodes.

But, it turned out that giant broadcast domains are pretty hard to manage. Apparently, in the real world, huge domains often lead to huge outages.

Thus, instead of having an immense data-center-wide broadcast domain, it's better to have smaller isolated L2 segments that can talk to each other through... L3 routers!

Getting, back to switches...

Modern network switches used in data centers are pretty advanced devices. They can work as bridges. Or, as routers. Or... some of their ports can work in a bridge mode, and some others - in a router mode. I'd speculate a bit and say that it's likely we're talking here about [_multilayer switches_](https://en.wikipedia.org/wiki/Multilayer_switch).

In a setup I was introduced to, all servers on a single rack belong to the same _/24 IP subnet_ and are connected to a single switch.

**_NB:_** _Facebook seems to have a pretty similar DC networking architecture. I didn't research much, but this [Load Balancer article](https://engineering.fb.com/2018/05/22/open-source/open-sourcing-katran-a-scalable-network-load-balancer/) indirectly confirms it._

Such a switch is called _top-of-rack (TOR)_.

For these servers, the TOR switch behaves as a canonical transparent [multi-port] bridge. I.e. as a pure Layer 2 device. And a rack forms equally sized L2 and L3 segments.

![Server rack with a top-of-rack switch](http://iximiuz.com/bridge-vs-switch/single-rack-2000-opt.png)

However, the L3 segments formed by individual racks need to be joined into internetwork. For that, one of the remaining ports of the TOR switch is configured to work in L3 mode. Unlike the L2 ports, it means that this port is an addressable network node with an IP address. It's then connected to a higher-layer switch.

![Simplified hierarchical internetworking model](http://iximiuz.com/bridge-vs-switch/multiple-racks-2000-opt.png)

_Simplified [hierarchical internetworking model](https://en.wikipedia.org/wiki/Hierarchical_internetworking_model)._

_Disclaimer: actually, every TOR switch is connected to at least two higher-layer switches. First of all, to provide some redundancy in the case of hardware failure. But also if these multiple physical connections are combined into a single logical link, it can increase the resulting throughput. However, for the sake of simplicity, let's omit this part._

These _higher-layer switches_ on the diagram above are called _distribution layer_ switches. In the network I was looking at, switches on this layer work as pure L3 routers. Hence, each port of a distribution-layer switch is a full-fledged Layer 3 device with an IP address assigned to it.

TOR-switches can be thought of as compound devices with a bridge (with tens of ports) and a router (with just 2 ports) inside. Distribution layer switches can be thought of as multi-port L3 routers. Like truly _multi_-port. 48- or 64-port routers! The router on a TOR-switch knows only two routes - to the rack's subnet and the default one, pointing to its distribution layer switch. And every distribution layer switch knows a lot of routes. Every rack connected to it has its own /24 IP subnet and this switch works as a border router for tens such subnets.

**However, physically, there is no difference between the first and the second types of switches!** They all look the same, but just configured differently.

In addition to these 48 (or 64) ports, switches have one or two [_out-of-band_](https://en.wikipedia.org/wiki/Out-of-band_data) ports.

![Picture of a real switch (hardware)](http://iximiuz.com/bridge-vs-switch/real-switch-2000-opt.png)

Regardless of the mode of other ports, the _out-of-band_ ports always work in L3 mode. One can log in on a switch using an IP address of an _out-of-band_ port. This is needed to manage switches because configuring switches through the normal ports would be simply dangerous. Imagine, you messed up with some commands and blocked yourself out of the switch?

Curious what would it look like when you _ssh_ to a switch? Surprise, surprise! It's Linux! Or FreeBSD. Or a proprietary Unix-like OS. I.e. one can configure a switch via a traditional ssh session using widely-known _iproute2_ tools such as `ip` and/or `bridge`.

From the management standpoint, every port on a switch looks like a traditional network device. One can enslave some ports to a _virtual Linux bridge_, assign some other ports IP addresses, configure packet forwarding between ports, set a route table, etc. So, you can think of a switch as of a Linux server with many-many network ports. And it's totally up to you how to configure them. All the traditional Linux capabilities are at your disposal. But of course, from the hardware standpoint, switches are highly-optimized packet processing devices.

![Managed network switch](http://iximiuz.com/bridge-vs-switch/managed-switch-2000-opt.png)

So, what's up with broadcast domains? Can we have a broadcast domain spanning multiple racks? Sure! [VXLAN](http://iximiuz.com/en/posts/computer-networking-101/#VXLAN) to the rescue! Up until this point, I was describing a physical network setup of a data center. But one can configure any sort of an overlay network on top of it making it tailored for the end-users use cases.

![VXLAN over multiple L3 networks](http://iximiuz.com/bridge-vs-switch/xvlan-2000-opt.png)

## Instead of conclusion

Actually, it was quite exciting for me to see the real production network in action. Things that I learned on that journey make a lot of sense to me. However, I'm pretty sure that it's not the only possible setup. So, I'd appreciate it if you share your experience here or [on twitter](https://twitter.com/iximiuz).

Have fun!

## Further Reading

- [Computer Networking Introduction - Ethernet and IP (Heavily Illustrated)](http://iximiuz.com/en/posts/computer-networking-101/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [Networking Lab - Ethernet Broadcast Domains](http://iximiuz.com/en/posts/networking-lab-ethernet-broadcast-domains/)
- [Networking Lab - Simple VLAN](http://iximiuz.com/en/posts/networking-lab-simple-vlan/)
- [Networking Lab - L3 to L2 Segments Mapping](http://iximiuz.com/en/posts/networking-lab-l3-to-l2-segments-mapping/)
- [Networking Lab - Proxy ARP](http://iximiuz.com/en/posts/networking-lab-proxy-arp/)
- [Networking Lab - Simple VXLAN](http://iximiuz.com/en/posts/networking-lab-simple-vxlan/)

[bridge,](javascript: void 0) [router,](javascript: void 0) [switch,](javascript: void 0) [Ethernet,](javascript: void 0) [IP](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

