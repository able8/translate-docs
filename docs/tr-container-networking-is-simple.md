# Container Networking Is Simple!

# 容器网络很简单！

October 18, 2020 (Updated: August 4, 2021)

_**Just kidding, it's not... But fear not and read on!**_

_**开个玩笑，这不是......但不要害怕，继续阅读！**_

_You can find a Russian translation of this article [here](https://habr.com/ru/company/timeweb/blog/558612/)_.

_您可以在 [此处](https://habr.com/ru/company/timeweb/blog/558612/)_ 中找到本文的俄语翻译。

Working with containers always feels like magic. In a good way for those who understand the internals and in a terrifying - for those who don't. Luckily, we've been looking under the hood of the containerization technology for quite some time already and even managed to uncover that [containers are just isolated and restricted Linux processes](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/#container-is-just-a-processes), that [images aren't really needed to run containers](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/), and on the contrary - [to build an image we need to run some containers](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/).

使用容器总是感觉像魔术一样。对于那些了解内部结构的人来说这是一种很好的方式，而对于那些不了解内部结构的人来说则是一种可怕的方式。幸运的是，我们已经研究了容器化技术的幕后很长一段时间，甚至设法发现 容器只是隔离和受限制的 Linux 进程，[运行容器并不真正需要图像](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)，相反 - [要构建映像，我们需要运行一些容器](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)。

Now comes a time to tackle the container networking problem. Or, more precisely, a single-host container networking problem. In this article, we are going to answer the following questions:

现在是解决容器网络问题的时候了。或者，更准确地说，是单主机容器网络问题。在本文中，我们将回答以下问题：

- How to virtualize network resources to make containers think each of them has a dedicated network stack?
- How to turn containers into friendly neighbors, prevent them from interfering, and teach to communicate well?
- How to reach the outside world (e.g. the Internet) from inside the container?
- How to reach containers running on a machine from the outside world ( _aka_ port publishing)?

- 如何虚拟化网络资源，让容器认为它们每个都有一个专用的网络堆栈？
- 如何将集装箱变成友好的邻居，防止他们干扰，并教他们良好的沟通？
- 如何从容器内部访问外部世界（例如互联网）？
- 如何从外部世界（_aka_ 端口发布）访问机器上运行的容器？

While answering these questions, we'll setup a container networking from scratch using standard Linux tools. As a result, it'll become apparent that the single-host container networking is nothing more than a simple combination of the well-known Linux facilities:

在回答这些问题时，我们将使用标准 Linux 工具从头开始设置容器网络。因此，很明显，单主机容器网络只不过是众所周知的 Linux 工具的简单组合：

- network namespaces;
- virtual Ethernet devices (veth);
- virtual network switches (bridge);
- IP routing and network address translation (NAT).

- 网络命名空间；
- 虚拟以太网设备（veth）；
- 虚拟网络交换机（网桥）；
- IP 路由和网络地址转换 (NAT)。

And for better or worse, no code is required to make the networking magic happen...

不管是好是坏，不需要任何代码来实现网络魔法……

## Prerequisites

## 先决条件

Any decent Linux distribution would probably suffice. All the examples in the article have been made on a fresh _vagrant_ CentOS 8 virtual machine:

任何体面的 Linux 发行版可能就足够了。文章中的所有示例都是在全新的 _vagrant_ CentOS 8 虚拟机上制作的：

```bash
$ vagrant init centos/8
$ vagrant up
$ vagrant ssh

[vagrant@localhost ~]$ uname -a
Linux localhost.localdomain 4.18.0-147.3.1.el8_1.x86_64

```

For the sake of simplicity of the examples, in this article, we are not going to rely on any fully-fledged containerization solution (e.g. _docker_ or _podman_). Instead, we'll focus on the basic concepts and use the bare minimum tooling to achieve our learning goals.

为了示例的简单起见，在本文中，我们不会依赖任何成熟的容器化解决方案（例如 _docker_ 或 _podman_）。相反，我们将专注于基本概念并使用最少的工具来实现我们的学习目标。

## Isolating containers with network namespaces

## 使用网络命名空间隔离容器

What constitutes a Linux network stack? Well, obviously, the set of network devices. What else? Probably, the set of routing rules. And not to forget, the set of netfilter hooks, including defined by iptables rules.

什么构成了 Linux 网络堆栈？嗯，很明显，网络设备的集合。还有什么？可能是路由规则集。不要忘记，一组 netfilter 钩子，包括由 iptables 规则定义的。

We can quickly forge a non-comprehensive `inspect-net-stack.sh` script:

我们可以快速伪造一个不全面的 `inspect-net-stack.sh` 脚本：

```bash
#!/usr/bin/env bash

echo "> Network devices"
ip link

echo -e "\n> Route table"
ip route

echo -e "\n> Iptables rules"
iptables --list-rules

```

Before running it, let's taint the iptables rules a bit to make them recognizable:

在运行它之前，让我们稍微修改一下 iptables 规则以使其易于识别：

```bash
$ sudo iptables -N ROOT_NS

```

After that, execution of the inspect script on my machine produces the following output:

之后，在我的机器上执行检查脚本会产生以下输出：

```bash
$ sudo ./inspect-net-stack.sh
> Network devices
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP mode DEFAULT group default qlen 1000
    link/ether 52:54:00:e3:27:77 brd ff:ff:ff:ff:ff:ff

> Route table
default via 10.0.2.2 dev eth0 proto dhcp metric 100
10.0.2.0/24 dev eth0 proto kernel scope link src 10.0.2.15 metric 100

> Iptables rules
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-N ROOT_NS

```

We are interested in that output because we want to make sure that each of the containers we are going to create soon will get a separate network stack. 

我们对该输出感兴趣，因为我们想确保我们即将创建的每个容器都将获得一个单独的网络堆栈。

Well, you might have heard already, that one of the Linux namespaces used for containers isolation is called _network namespace_. From [`man ip-netns`](https://man7.org/linux/man-pages/man8/ip-netns.8.html), _"network namespace is logically another copy of the network stack, with its own routes, firewall rules, and network devices."_ For the sake of simplicity, this is the only namespace we're going to use in this article. Instead of creating fully-isolated containers, we'd rather restrict the scope to only the network stack.

好吧，您可能已经听说过，用于容器隔离的 Linux 命名空间之一称为 _network namespace_。来自 [`man ip-netns`](https://man7.org/linux/man-pages/man8/ip-netns.8.html)，_"network 命名空间在逻辑上是网络堆栈的另一个副本，其自己的路由、防火墙规则和网络设备。”_ 为简单起见，这是我们将在本文中使用的唯一命名空间。与其创建完全隔离的容器，不如将范围限制在网络堆栈上。

One of the ways to create a network namespace is the `ip` tool - part of the de facto standard [iproute2](https://en.wikipedia.org/wiki/Iproute2) collection:

创建网络命名空间的方法之一是 `ip` 工具 - 事实标准 [iproute2](https://en.wikipedia.org/wiki/Iproute2) 集合的一部分：

```bash
$ sudo ip netns add netns0
$ ip netns
netns0
```

How to start using the just created namespace? There is a lovely Linux command called `nsenter`. It enters one or more of the specified namespaces and then executes the given program:

如何开始使用刚刚创建的命名空间？有一个可爱的 Linux 命令叫做 `nsenter`。它进入一个或多个指定的命名空间，然后执行给定的程序：

```bash
$ sudo nsenter --net=/var/run/netns/netns0 bash
# The newly created bash process lives in netns0

$ sudo ./inspect-net-stack.sh
> Network devices
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

> Route table

> Iptables rules
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT

```

From the output above it's clear that the _bash_ process running inside `netns0` namespace sees a totally different network stack. There is no routing rules at all, no custom iptables chain, and only one loopback network device. So far, so good...

从上面的输出可以清楚地看出，在 `netns0` 命名空间内运行的 _bash_ 进程看到了一个完全不同的网络堆栈。完全没有路由规则，没有自定义iptables链，只有一个loopback网络设备。到现在为止还挺好...

![Linux network namespace visualized](http://iximiuz.com/container-networking-is-simple/network-namespace-4000-opt.png)

_Network namespace visualized._

_网络命名空间可视化。_

## Connecting containers to host with virtual Ethernet devices (veth)

## 使用虚拟以太网设备（veth）将容器连接到主机

A dedicated network stack would be not so useful if we could not communicate with it. Luckily, Linux provides a suitable facility for that - a virtual Ethernet device! From [`man veth`](https://man7.org/linux/man-pages/man4/veth.4.html), _"veth devices are virtual Ethernet devices. They can act as tunnels between network namespaces to create a bridge to a physical network device in another namespace, but can also be used as standalone network devices."_

如果我们不能与之通信，专用的网络堆栈将不会那么有用。幸运的是，Linux 提供了一个合适的工具——一个虚拟以太网设备！来自 [`man veth`](https://man7.org/linux/man-pages/man4/veth.4.html)，_"veth 设备是虚拟以太网设备。它们可以充当网络命名空间之间的隧道来创建连接到另一个命名空间中的物理网络设备的桥接器，但也可以用作独立的网络设备。”_

Virtual Ethernet devices always go in pairs. No worries, it'll be clear when we take a look at the creation command:

虚拟以太网设备总是成对使用。不用担心，当我们看一下创建命令时就会很清楚：

```bash
$ sudo ip link add veth0 type veth peer name ceth0
```

With this single command, we just created a pair of _interconnected_ virtual Ethernet devices. The names `veth0` and `ceth0` have been chosen arbitrarily:

使用这个单一命令，我们刚刚创建了一对 _interconnected_ 虚拟以太网设备。名称 `veth0` 和 `ceth0` 是任意选择的：

```bash
$ ip link
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP mode DEFAULT group default qlen 1000
    link/ether 52:54:00:e3:27:77 brd ff:ff:ff:ff:ff:ff
5: ceth0@veth0: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 66:2d:24:e3:49:3f brd ff:ff:ff:ff:ff:ff
6: veth0@ceth0: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 96:e8:de:1d:22:e0 brd ff:ff:ff:ff:ff:ff
```

Both `veth0` and `ceth0` after creation resides on the host's network stack (also called root network namespace). To connect the root namespace with the `netns0` namespace, we need to keep one of the devices in the root namespace and move another one into the `netns0`:

创建后的 `veth0` 和 `ceth0` 都驻留在主机的网络堆栈（也称为根网络命名空间）上。要将根命名空间与 `netns0` 命名空间连接起来，我们需要将其中一台设备保留在根命名空间中，并将另一台设备移动到 `netns0` 中：

```bash
$ sudo ip link set ceth0 netns netns0

# List all the devices to make sure one of them disappeared from the root stack
$ ip link
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP mode DEFAULT group default qlen 1000
    link/ether 52:54:00:e3:27:77 brd ff:ff:ff:ff:ff:ff
6: veth0@if5: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 96:e8:de:1d:22:e0 brd ff:ff:ff:ff:ff:ff link-netns netns0

```

Once we turn the devices on and assign proper IP addresses, any packet occurring on one of the devices will immediately pop up on its peer device connecting two namespaces. Let's start from the root namespace:

一旦我们打开设备并分配正确的 IP 地址，其中一个设备上发生的任何数据包都会立即在其连接两个命名空间的对等设备上弹出。让我们从根命名空间开始：

```bash
$ sudo ip link set veth0 up
$ sudo ip addr add 172.18.0.11/16 dev veth0

```

And continue with the `netns0`:

并继续使用`netns0`：

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ip link set lo up  # whoops
$ ip link set ceth0 up
$ ip addr add 172.18.0.10/16 dev ceth0
$ ip link
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
5: ceth0@if6: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether 66:2d:24:e3:49:3f brd ff:ff:ff:ff:ff:ff link-netnsid 0

```

![Connecting network namespaces via veth device](http://iximiuz.com/container-networking-is-simple/veth-4000-opt.png)

_Connecting network namespaces via veth device._

_通过 veth 设备连接网络命名空间。_

We are ready to check the connectivity:

我们准备好检查连通性：

```bash
# From `netns0`, ping root's veth0
$ ping -c 2 172.18.0.11
PING 172.18.0.11 (172.18.0.11) 56(84) bytes of data.
64 bytes from 172.18.0.11: icmp_seq=1 ttl=64 time=0.038 ms
64 bytes from 172.18.0.11: icmp_seq=2 ttl=64 time=0.040 ms

--- 172.18.0.11 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 58ms
rtt min/avg/max/mdev = 0.038/0.039/0.040/0.001 ms

# Leave `netns0`
$ exit

# From root namespace, ping ceth0
$ ping -c 2 172.18.0.10
PING 172.18.0.10 (172.18.0.10) 56(84) bytes of data.
64 bytes from 172.18.0.10: icmp_seq=1 ttl=64 time=0.073 ms
64 bytes from 172.18.0.10: icmp_seq=2 ttl=64 time=0.046 ms

--- 172.18.0.10 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 3ms
rtt min/avg/max/mdev = 0.046/0.059/0.073/0.015 ms

```

At the same time, if we try to reach any other addresses from the `netns0` namespace, we are not going to succeed:

同时，如果我们尝试从 `netns0` 命名空间访问任何其他地址，我们将不会成功：

```bash
# Inside root namespace
$ ip addr show dev eth0
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 52:54:00:e3:27:77 brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global dynamic noprefixroute eth0
       valid_lft 84057sec preferred_lft 84057sec
    inet6 fe80::5054:ff:fee3:2777/64 scope link
       valid_lft forever preferred_lft forever

# Remember this 10.0.2.15

$ sudo nsenter --net=/var/run/netns/netns0

# Try host's eth0
$ ping 10.0.2.15
connect: Network is unreachable

# Try something from the Internet
$ ping 8.8.8.8
connect: Network is unreachable

```

That's easy to explain, though. There is simply no route in the `netns0` routing table for such packets. The only entry there shows how to reach `172.18.0.0/16` network:

不过，这很容易解释。在`netns0` 路由表中根本没有针对此类数据包的路由。唯一的条目显示了如何访问“172.18.0.0/16”网络：

```bash
# From `netns0` namespace:
$ ip route
172.18.0.0/16 dev ceth0 proto kernel scope link src 172.18.0.10

```

Linux has a bunch of ways to populate the routing table. One of them is to extract routes from the directly attached network interfaces. Remember, the routing table in `netns0` was empty right after the namespace creation. But then we added the `ceth0` device there and assigned it an IP address `172.18.0.10/16`. Since we were using not a simple IP address, but a combination of the address and the netmask, the network stack managed to extract the routing information from it. Every packet destined to `172.18.0.0/16` network will be sent through `ceth0` device. But any other packets will be discarded. Similarly, there is a new route in the root namespace:

Linux 有很多方法来填充路由表。其中之一是从直接连接的网络接口中提取路由。请记住，在创建命名空间后，`netns0` 中的路由表是空的。但随后我们在那里添加了“ceth0”设备并为其分配了一个 IP 地址“172.18.0.10/16”。由于我们使用的不是简单的 IP 地址，而是地址和网络掩码的组合，因此网络堆栈设法从中提取路由信息。每个发往“172.18.0.0/16”网络的数据包都将通过“ceth0”设备发送。但是任何其他数据包都将被丢弃。同样，根命名空间中有一个新路由：

```bash
# From `root` namespace:
$ ip route
# ... omitted lines ...
172.18.0.0/16 dev veth0 proto kernel scope link src 172.18.0.11

```

At this point, we are ready to mark our very first question answered. **We know now how to isolate, virtualize, and connect Linux network stacks.**

此时，我们已准备好标记我们的第一个问题已回答。 **我们现在知道如何隔离、虚拟化和连接 Linux 网络堆栈。**

## Interconnecting containers with virtual network switch (bridge) 

## 用虚拟网络交换机（网桥）互连容器

The whole idea of containerization boils down to efficient resource sharing. I.e. it's uncommon to have a single container per machine. Instead, the goal is to run as many isolated processes in the shared environment as possible. So, what'd happen if we were to place multiple containers on the same host following the `veth` approach from above? Let's add the second _container_:

容器化的整个想法归结为有效的资源共享。 IE。每台机器有一个容器并不常见。相反，目标是在共享环境中运行尽可能多的隔离进程。那么，如果我们按照上面的“veth”方法将多个容器放在同一主机上会发生什么？让我们添加第二个 _container_：

```bash
# From root namespace
$ sudo ip netns add netns1
$ sudo ip link add veth1 type veth peer name ceth1
$ sudo ip link set ceth1 netns netns1
$ sudo ip link set veth1 up
$ sudo ip addr add 172.18.0.21/16 dev veth1

$ sudo nsenter --net=/var/run/netns/netns1
$ ip link set lo up
$ ip link set ceth1 up
$ ip addr add 172.18.0.20/16 dev ceth1

```

My favourite part, checking the connectivity:

我最喜欢的部分，检查连接：

```bash
# From `netns1` we cannot reach the root namespace!
$ ping -c 2 172.18.0.21
PING 172.18.0.21 (172.18.0.21) 56(84) bytes of data.
From 172.18.0.20 icmp_seq=1 Destination Host Unreachable
From 172.18.0.20 icmp_seq=2 Destination Host Unreachable

--- 172.18.0.21 ping statistics ---
2 packets transmitted, 0 received, +2 errors, 100% packet loss, time 55ms
pipe 2

# But there is a route!
$ ip route
172.18.0.0/16 dev ceth1 proto kernel scope link src 172.18.0.20

# Leaving `netns1`
$ exit

# From root namespace we cannot reach the `netns1`
$ ping -c 2 172.18.0.20
PING 172.18.0.20 (172.18.0.20) 56(84) bytes of data.
From 172.18.0.11 icmp_seq=1 Destination Host Unreachable
From 172.18.0.11 icmp_seq=2 Destination Host Unreachable

--- 172.18.0.20 ping statistics ---
2 packets transmitted, 0 received, +2 errors, 100% packet loss, time 23ms
pipe 2

# From `netns0` we CAN reach `veth1`
$ sudo nsenter --net=/var/run/netns/netns0
$ ping -c 2 172.18.0.21
PING 172.18.0.21 (172.18.0.21) 56(84) bytes of data.
64 bytes from 172.18.0.21: icmp_seq=1 ttl=64 time=0.037 ms
64 bytes from 172.18.0.21: icmp_seq=2 ttl=64 time=0.046 ms

--- 172.18.0.21 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 33ms
rtt min/avg/max/mdev = 0.037/0.041/0.046/0.007 ms

# But we still cannot reach `netns1`
$ ping -c 2 172.18.0.20
PING 172.18.0.20 (172.18.0.20) 56(84) bytes of data.
From 172.18.0.10 icmp_seq=1 Destination Host Unreachable
From 172.18.0.10 icmp_seq=2 Destination Host Unreachable

--- 172.18.0.20 ping statistics ---
2 packets transmitted, 0 received, +2 errors, 100% packet loss, time 63ms
pipe 2

```

Whoops! Something is wrong... `netns1` is stuck in limbo. For some reason, it cannot talk to the root and from the root namespace we cannot reach it out too. However, since both containers reside in the same IP network `172.18.0.0/16`, we now can talk to the host's `veth1` from the `netns0` container. Interesting...

哎呀！出了点问题...... `netns1` 陷入困境。出于某种原因，它无法与根通信，并且我们也无法从根命名空间访问它。但是，由于两个容器都位于同一个 IP 网络“172.18.0.0/16”中，我们现在可以从“netns0”容器中与主机的“veth1”通信。有趣的...

Well, it took me some time to figure it out, but apparently we are facing the clash of routes. Let's inspect the routing table in the root namespace:

好吧，我花了一些时间才弄明白，但显然我们正面临路线冲突。让我们检查根命名空间中的路由表：

```bash
$ ip route
# ... omitted lines ...
172.18.0.0/16 dev veth0 proto kernel scope link src 172.18.0.11
172.18.0.0/16 dev veth1 proto kernel scope link src 172.18.0.21

```

Even though after adding the second `veth` pair, root's network stack learned the new route `172.18.0.0/16 dev veth1 proto kernel scope link src 172.18.0.21`, there already was an existing route for exactly the same network. When the second container tries to ping `veth1` device, the first route is being selected breaking the connectivity. If we were to delete the first route `sudo ip route delete 172.18.0.0/16 dev veth0 proto kernel scope link src 172.18.0.11` and recheck the connectivity, the situation would turn into a mirrored case. I.e. the connectivity of the `netns1` would be restored, but `netns0` would be in limbo.

即使在添加第二个 `veth` 对后，root 的网络堆栈学习到了新路由 `172.18.0.0/16 dev veth1 proto kernel scope link src 172.18.0.21`，已经存在完全相同网络的现有路由。当第二个容器尝试 ping 'veth1' 设备时，第一个路由被选择破坏了连接。如果我们删除第一条路由`sudo ip route delete 172.18.0.0/16 dev veth0 proto kernel scope link src 172.18.0.11`并重新检查连通性，情况就会变成镜像情况。 IE。 `netns1` 的连接将恢复，但 `netns0` 将处于不确定状态。

![Connecting multiple network namespaces with a bridge](http://iximiuz.com/container-networking-is-simple/route-clash-4000-opt.png)

Well, I believe if we selected another IP network for `netns1`, everything would work. However, multiple containers sitting in one IP network is a legitimate use case. Thus, we need to adjust the `veth` approach somehow...

好吧，我相信如果我们为 `netns1` 选择另一个 IP 网络，一切都会正常。但是，位于一个 IP 网络中的多个容器是一种合法用例。因此，我们需要以某种方式调整 `veth` 方法......

Behold the Linux bridge - yet another virtualized network facility! The Linux bridge behaves like a network switch. It forwards packets between interfaces that are connected to it. And since it's a switch, it does it on the L2 (i.e. Ethernet) level. 

看看 Linux 网桥——又一个虚拟化的网络设施！ Linux 网桥的行为类似于网络交换机。它在连接到它的接口之间转发数据包。因为它是一个交换机，所以它在 L2（即以太网）级别上进行。

Let's try to play with our new toy. But first, we need to clean up the existing setup because some of the configurational changes we've made so far aren't really needed anymore. Removing network namespaces would suffice:

让我们试着玩我们的新玩具。但首先，我们需要清理现有设置，因为我们目前所做的一些配置更改实际上不再需要了。删除网络命名空间就足够了：

```bash
$ sudo ip netns delete netns0
$ sudo ip netns delete netns1

# But if you still have some leftovers...
$ sudo ip link delete veth0
$ sudo ip link delete ceth0
$ sudo ip link delete veth1
$ sudo ip link delete ceth1

```

Quickly re-create two containers. Notice, we don't assign any IP address to the new `veth0` and `veth1` devices:

快速重新创建两个容器。请注意，我们没有为新的 `veth0` 和 `veth1` 设备分配任何 IP 地址：

```bash
$ sudo ip netns add netns0
$ sudo ip link add veth0 type veth peer name ceth0
$ sudo ip link set veth0 up
$ sudo ip link set ceth0 netns netns0

$ sudo nsenter --net=/var/run/netns/netns0
$ ip link set lo up
$ ip link set ceth0 up
$ ip addr add 172.18.0.10/16 dev ceth0
$ exit

$ sudo ip netns add netns1
$ sudo ip link add veth1 type veth peer name ceth1
$ sudo ip link set veth1 up
$ sudo ip link set ceth1 netns netns1

$ sudo nsenter --net=/var/run/netns/netns1
$ ip link set lo up
$ ip link set ceth1 up
$ ip addr add 172.18.0.20/16 dev ceth1
$ exit

```

Make sure there is no new routes on the host:

确保主机上没有新路由：

```bash
$ ip route
default via 10.0.2.2 dev eth0 proto dhcp metric 100
10.0.2.0/24 dev eth0 proto kernel scope link src 10.0.2.15 metric 100

```

And finally, create the bridge interface:

最后，创建桥接口：

```bash
$ sudo ip link add br0 type bridge
$ sudo ip link set br0 up

```

Now, attach `veth0` and `veth1` ends to the bridge:

现在，将 `veth0` 和 `veth1` 端附加到桥上：

```bash
$ sudo ip link set veth0 master br0
$ sudo ip link set veth1 master br0
```

![Setting up routing between multiple network namespaces](http://iximiuz.com/container-networking-is-simple/bridge-4000-opt.png)

...and check the connectivity between containers:

...并检查容器之间的连接：

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ping -c 2 172.18.0.20
PING 172.18.0.20 (172.18.0.20) 56(84) bytes of data.
64 bytes from 172.18.0.20: icmp_seq=1 ttl=64 time=0.259 ms
64 bytes from 172.18.0.20: icmp_seq=2 ttl=64 time=0.051 ms

--- 172.18.0.20 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 2ms
rtt min/avg/max/mdev = 0.051/0.155/0.259/0.104 ms

```

```bash
$ sudo nsenter --net=/var/run/netns/netns1
$ ping -c 2 172.18.0.10
PING 172.18.0.10 (172.18.0.10) 56(84) bytes of data.
64 bytes from 172.18.0.10: icmp_seq=1 ttl=64 time=0.037 ms
64 bytes from 172.18.0.10: icmp_seq=2 ttl=64 time=0.089 ms

--- 172.18.0.10 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 36ms
rtt min/avg/max/mdev = 0.037/0.063/0.089/0.026 ms

```

Lovely! Everything works great. With this new approach, we haven't been configuring `veth0` and `veth1` at all. The only two IP addresses we assigned were on the `ceth0` and `ceth1` ends. But since both of them are on the same Ethernet segment (remember, we connected them to the virtual switch), there is connectivity on the L2 level:

迷人的！一切都很好。使用这种新方法，我们根本没有配置 `veth0` 和 `veth1`。我们分配的仅有的两个 IP 地址位于“ceth0”和“ceth1”端。但是由于它们都在同一个以太网段上（请记住，我们将它们连接到虚拟交换机），因此在 L2 级别存在连接：

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ip neigh
172.18.0.20 dev ceth0 lladdr 6e:9c:ae:02:60:de STALE
$ exit

$ sudo nsenter --net=/var/run/netns/netns1
$ ip neigh
172.18.0.10 dev ceth1 lladdr 66:f3:8c:75:09:29 STALE
$ exit

```

Congratulations, we learned how to **turn containers into friendly neighbors, prevent them from interfering, but keep the connectivity.**

恭喜，我们学会了如何**将容器变成友好的邻居，防止它们干扰，但保持连接。**

## Reaching out to the outside world (IP routing and masquerading)

## 与外界联系（IP 路由和伪装）

Our containers can talk to each other. But can they talk to the host, i.e. the root namespace?

我们的容器可以相互交谈。但是它们可以与主机（即根命名空间）通信吗？

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ping 10.0.2.15  # eth0 address
connect: Network is unreachable

```

That's kind of obvious, there is simply no route for that in `netns0`:

这很明显，在 `netns0` 中根本没有路径：

```bash
$ ip route
172.18.0.0/16 dev ceth0 proto kernel scope link src 172.18.0.10

```

The root namespace cannot talk to containers either:

根命名空间也不能与容器通信：

```bash
# Use exit to leave `netns0` first:
$ ping -c 2 172.18.0.10
PING 172.18.0.10 (172.18.0.10) 56(84) bytes of data.
From 213.51.1.123 icmp_seq=1 Destination Net Unreachable
From 213.51.1.123 icmp_seq=2 Destination Net Unreachable

--- 172.18.0.10 ping statistics ---
2 packets transmitted, 0 received, +2 errors, 100% packet loss, time 3ms

$ ping -c 2 172.18.0.20
PING 172.18.0.20 (172.18.0.20) 56(84) bytes of data.
From 213.51.1.123 icmp_seq=1 Destination Net Unreachable
From 213.51.1.123 icmp_seq=2 Destination Net Unreachable

--- 172.18.0.20 ping statistics ---
2 packets transmitted, 0 received, +2 errors, 100% packet loss, time 3ms

```

To establish the connectivity between the root and container namespaces, we need to assign the IP address to the bridge network interface:

为了在根命名空间和容器命名空间之间建立连接，我们需要为桥接网络接口分配 IP 地址：

```bash
$ sudo ip addr add 172.18.0.1/16 dev br0

```

Once we assigned the IP address to the bridge interface, we got a route on the host routing table:

一旦我们为网桥接口分配了 IP 地址，我们就在主机路由表上得到了一条路由：

```bash
$ ip route
# ... omitted lines ...
172.18.0.0/16 dev br0 proto kernel scope link src 172.18.0.1

$ ping -c 2 172.18.0.10
PING 172.18.0.10 (172.18.0.10) 56(84) bytes of data.
64 bytes from 172.18.0.10: icmp_seq=1 ttl=64 time=0.036 ms
64 bytes from 172.18.0.10: icmp_seq=2 ttl=64 time=0.049 ms

--- 172.18.0.10 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 11ms
rtt min/avg/max/mdev = 0.036/0.042/0.049/0.009 ms

$ ping -c 2 172.18.0.20
PING 172.18.0.20 (172.18.0.20) 56(84) bytes of data.
64 bytes from 172.18.0.20: icmp_seq=1 ttl=64 time=0.059 ms
64 bytes from 172.18.0.20: icmp_seq=2 ttl=64 time=0.056 ms

--- 172.18.0.20 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 4ms
rtt min/avg/max/mdev = 0.056/0.057/0.059/0.007 ms

```

The container probably also got an ability to ping the bridge interface, but they still cannot reach out to host's `eth0`. We need to add the default route for containers:

容器可能也有能力 ping 桥接接口，但它们仍然无法访问主机的“eth0”。我们需要为容器添加默认路由：

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ip route add default via 172.18.0.1
$ ping -c 2 10.0.2.15
PING 10.0.2.15 (10.0.2.15) 56(84) bytes of data.
64 bytes from 10.0.2.15: icmp_seq=1 ttl=64 time=0.036 ms
64 bytes from 10.0.2.15: icmp_seq=2 ttl=64 time=0.053 ms

--- 10.0.2.15 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 14ms
rtt min/avg/max/mdev = 0.036/0.044/0.053/0.010 ms

# And repeat the change for `netns1`

```

This change basically turned the host machine into a router and the bridge interface became the default gateway for the containers.

这一变化基本上将主机变成了路由器，桥接接口成为容器的默认网关。

![Using bridge as a gateway](http://iximiuz.com/container-networking-is-simple/router-4000-opt.png)

Perfect, we connected containers with the root namespace. Now, let's try to connect them to the outside world. By default, the packet forwarding (i.e. the router functionality) is disabled in Linux. We need to turn it on:

完美，我们将容器与根命名空间连接起来。现在，让我们尝试将它们连接到外部世界。默认情况下，Linux 中禁用数据包转发（即路由器功能）。我们需要开启它：

```bash
# In the root namespace
sudo bash -c 'echo 1 > /proc/sys/net/ipv4/ip_forward'

```

Again, my favourite part - checking the connectivity:

再次，我最喜欢的部分 - 检查连接：

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ping 8.8.8.8
# hangs indefinitely long for me...

```

Well, still doesn't work. What have we missed? If the container were to sends packets to the outside world, the destination server would not be able to send packets back to the container because the container's IP address is private. I.e. the routing rules for that particular IP are known only to the local network. And lots of the containers in the world share exactly the same private IP address `172.18.0.10`. The solution to this problem is called the [Network address translation (NAT)](https://en.wikipedia.org/wiki/Network_address_translation). Before going to the external network, packets originated by the containers will get their source IP addresses replaced with the host's external interface address. The host also will track all the existing mappings and on arrival, it'll be restoring the IP addresses before forwarding packets back to the containers. Sounds complicated, but I have good news for you! Thanks to [iptables](http://iximiuz.com/en/posts/laymans-iptables-101/) module, we need only a single command to make it happen:

嗯，还是不行。我们错过了什么？如果容器将数据包发送到外部世界，目标服务器将无法将数据包发送回容器，因为容器的 IP 地址是私有的。 IE。该特定 IP 的路由规则只有本地网络知道。世界上许多容器共享完全相同的私有 IP 地址“172.18.0.10”。此问题的解决方案称为 [网络地址转换 (NAT)](https://en.wikipedia.org/wiki/Network_address_translation)。在进入外部网络之前，由容器发起的数据包会将其源 IP 地址替换为主机的外部接口地址。主机还将跟踪所有现有映射，并在到达时恢复 IP 地址，然后将数据包转发回容器。听起来很复杂，但我有个好消息要告诉你！感谢 [iptables](http://iximiuz.com/en/posts/laymans-iptables-101/) 模块，我们只需要一个命令就可以实现：

```bash
$ sudo iptables -t nat -A POSTROUTING -s 172.18.0.0/16 !-o br0 -j MASQUERADE

```

The command is fairly simple. We are adding a new rule to the `nat` table of the `POSTROUTING` chain asking to masquerade all the packets originated in `172.18.0.0/16` network, but not by the bridge interface.

命令相当简单。我们正在向`POSTROUTING`链的`nat`表添加一个新规则，要求伪装所有源自`172.18.0.0/16`网络的数据包，但不是来自网桥接口。

Check the connectivity:

检查连通性：

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ping -c 2 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=61 time=43.2 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=61 time=36.8 ms

--- 8.8.8.8 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 2ms
rtt min/avg/max/mdev = 36.815/40.008/43.202/3.199 ms

```

Beware that we're following _by default - allow_ strategy which might be quite dangerous in a real-world setup. The host's default iptables policy is `ACCEPT` for every chain:

请注意，我们正在遵循 _by default - allow_ 策略，这在实际设置中可能非常危险。对于每个链，主机的默认 iptables 策略是 `ACCEPT`：

```bash
sudo iptables -S
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT

```

As a good example, Docker, instead, restricts everything by default and then enables routing for only known paths.

作为一个很好的例子，Docker 在默认情况下限制一切，然后只为已知路径启用路由。

Click here to see Docker iptables rules.

单击此处查看 Docker iptables 规则。

The following are the dumped rules generated by the Docker daemon on a _CentOS 8_ machine with single container exposed on port 5005:

以下是 _CentOS 8_ 机器上的 Docker 守护程序生成的转储规则，单个容器暴露在端口 5005 上：

```bash
$ sudo iptables -t filter --list-rules
-P INPUT ACCEPT
-P FORWARD DROP
-P OUTPUT ACCEPT
-N DOCKER
-N DOCKER-ISOLATION-STAGE-1
-N DOCKER-ISOLATION-STAGE-2
-N DOCKER-USER
-A FORWARD -j DOCKER-USER
-A FORWARD -j DOCKER-ISOLATION-STAGE-1
-A FORWARD -o docker0 -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
-A FORWARD -o docker0 -j DOCKER
-A FORWARD -i docker0 !-o docker0 -j ACCEPT
-A FORWARD -i docker0 -o docker0 -j ACCEPT
-A DOCKER -d 172.17.0.2/32 !-i docker0 -o docker0 -p tcp -m tcp --dport 5000 -j ACCEPT
-A DOCKER-ISOLATION-STAGE-1 -i docker0 !-o docker0 -j DOCKER-ISOLATION-STAGE-2
-A DOCKER-ISOLATION-STAGE-1 -j RETURN
-A DOCKER-ISOLATION-STAGE-2 -o docker0 -j DROP
-A DOCKER-ISOLATION-STAGE-2 -j RETURN
-A DOCKER-USER -j RETURN

$ sudo iptables -t nat --list-rules
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P POSTROUTING ACCEPT
-P OUTPUT ACCEPT
-N DOCKER
-A PREROUTING -m addrtype --dst-type LOCAL -j DOCKER
-A POSTROUTING -s 172.17.0.0/16 !-o docker0 -j MASQUERADE
-A POSTROUTING -s 172.17.0.2/32 -d 172.17.0.2/32 -p tcp -m tcp --dport 5000 -j MASQUERADE
-A OUTPUT !-d 127.0.0.0/8 -m addrtype --dst-type LOCAL -j DOCKER
-A DOCKER -i docker0 -j RETURN
-A DOCKER !-i docker0 -p tcp -m tcp --dport 5005 -j DNAT --to-destination 172.17.0.2:5000

$ sudo iptables -t mangle --list-rules
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT

$ sudo iptables -t raw --list-rules
-P PREROUTING ACCEPT
-P OUTPUT ACCEPT

```

## Letting the outside world reach out to containers (port publishing)

## 让外界接触到容器（端口发布）

It's a known practice to _publish_ container ports to some (or all) of the host's interfaces. But what does _port publishing_ really mean?

将容器端口_发布_到主机的某些（或全部）接口是一种众所周知的做法。但是 _port 发布_ 到底是什么意思？

Imagine we have a server running inside a container:

想象一下，我们有一个在容器内运行的服务器：

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ python3 -m http.server --bind 172.18.0.10 5000

```

If we try to send an HTTP request to this server process from the host, everything will work (well, there is a connectivity between root namespace and all the container interfaces, why wouldn't it?):

如果我们尝试从主机向此服务器进程发送 HTTP 请求，一切都会正常工作（好吧，根命名空间和所有容器接口之间存在连接，为什么不呢？）：

```bash
# From root namespace
$ curl 172.18.0.10:5000
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
# ... omitted lines ...

```

However, if we were to access this server from the outside world, what IP address would we use? The only IP address we might know is the host's external interface address `eth0`:

但是，如果我们从外部访问这个服务器，我们会使用什么 IP 地址？我们可能知道的唯一 IP 地址是主机的外部接口地址“eth0”：

```bash
$ curl 10.0.2.15:5000
curl: (7) Failed to connect to 10.0.2.15 port 5000: Connection refused

```

Thus, we need to find a way to forward any packets arriving at port 5000 on the host's `eth0` interface to `172.18.0.10:5000` destination. Or, in other words, we need to _publish_ the container's port 5000 on the host's `eth0` interface. iptables to the rescue!

因此，我们需要找到一种方法将任何到达主机“eth0”接口上端口 5000 的数据包转发到“172.18.0.10:5000”目的地。或者，换句话说，我们需要在主机的 eth0 接口上_发布_容器的端口 5000。 iptables 来救援！

```bash
# External traffic
sudo iptables -t nat -A PREROUTING -d 10.0.2.15 -p tcp -m tcp --dport 5000 -j DNAT --to-destination 172.18.0.10:5000

# Local traffic (since it doesn't pass the PREROUTING chain)
sudo iptables -t nat -A OUTPUT -d 10.0.2.15 -p tcp -m tcp --dport 5000 -j DNAT --to-destination 172.18.0.10:5000

```

Additionally, we need to enable [iptables intercepting traffic over bridged networks](https://github.com/omribahumi/libvirt_metadata_api/pull/4/files):

此外，我们需要启用 [iptables 拦截桥接网络上的流量](https://github.com/omribahumi/libvirt_metadata_api/pull/4/files)：

```bash
sudo modprobe br_netfilter

```

Testing time!

测试时间！

```bash
curl 10.0.2.15:5000
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
# ... omitted lines ...
```

## Understanding Docker network drivers

## 了解 Docker 网络驱动程序

Ok, sir, what can we do with all this useless knowledge? For instance, we could try to understand some of the [Docker network modes](https://docs.docker.com/network/#network-drivers)! 

好的，先生，我们可以用这些无用的知识做什么？例如，我们可以尝试了解一些[Docker 网络模式](https://docs.docker.com/network/#network-drivers)！

Let's start from the `--network host` mode. Try to compare the output of the following commands `ip link` and `sudo docker run -it --rm --network host alpine ip link`. Surprise, surprise, they are exactly the same! I.e. in the `host` mode, Docker simply doesn't use the network namespace isolation and containers work in the root network namespace and share the network stack with the host machine.

让我们从`--network host`模式开始。尝试比较以下命令 `ip link` 和 `sudo docker run -it --rm --network host alpine ip link` 的输出。惊喜，惊喜，它们一模一样！ IE。在`host` 模式下，Docker 根本不使用网络命名空间隔离，容器在根网络命名空间中工作并与主机共享网络堆栈。

The next mode to inspect is `--network none`. The output of the `sudo docker run -it --rm --network none alpine ip link` command shows only a single loopback network interface. It's very similar to our observations of the freshly created network namespace. I.e. before the point where we were adding any `veth` devices.

下一个要检查的模式是`--network none`。 `sudo docker run -it --rm --network none alpine ip link` 命令的输出仅显示一个环回网络接口。这与我们对新创建的网络命名空间的观察非常相似。 IE。在我们添加任何“veth”设备之前。

Last but not least, the `--network bridge` (the default) mode. Well, it's exactly what we've been trying to reproduce in this whole article. I encourage you to play with `ip` and `iptables` commands and inspect the network stack from the host and containers point of view.

最后但并非最不重要的是，`--network bridge`（默认）模式。嗯，这正是我们在整篇文章中一直试图重现的内容。我鼓励你使用 `ip` 和 `iptables` 命令，并从主机和容器的角度检查网络堆栈。

## Bonus: rootless containers and networking

## 奖励：无根容器和网络

One of the nice features of `podman` container manager is its focus on rootless containers. However, as you probably noticed, we used a lot of `sudo` escalations in this article. I.e. it's impossible to configure the network without root privileges. [Podman's approach](https://www.redhat.com/sysadmin/container-networking-podman) to rootful networking is very close to docker. But when it comes to rootless containers, podman relies on [slirp4netns](https://github.com/rootless-containers/slirp4netns) project:

`podman` 容器管理器的一个很好的特性是它专注于无根容器。但是，您可能已经注意到，我们在本文中使用了很多 `sudo` 升级。 IE。没有root权限就不可能配置网络。  [Podman 的做法](https://www.redhat.com/sysadmin/container-networking-podman) 对 rootful 网络非常接近 docker。但是说到无根容器，podman 依赖于 [slirp4netns](https://github.com/rootless-containers/slirp4netns) 项目：

> Starting with Linux 3.8, unprivileged users can create network\_namespaces(7) along with user\_namespaces(7). However, unprivileged network namespaces had not been very useful, because creating veth(4) pairs across the host and network namespaces still requires the root privileges. (i.e. No internet connection)

> 从 Linux 3.8 开始，非特权用户可以创建 network\_namespaces(7) 和 user\_namespaces(7)。然而，非特权网络命名空间并不是很有用，因为在主机和网络命名空间之间创建 veth(4) 对仍然需要 root 权限。 （即没有互联网连接）

> slirp4netns allows connecting a network namespace to the Internet in a completely unprivileged way, by connecting a TAP device in a network namespace to the usermode TCP/IP stack ("slirp").

> slirp4netns 允许以完全无特权的方式将网络命名空间连接到 Internet，方法是将网络命名空间中的 TAP 设备连接到用户模式 TCP/IP 堆栈（“slirp”）。

The rootless networking [is quite limited](https://www.redhat.com/sysadmin/container-networking-podman): "technically, the container itself does not have an IP address, because without root privileges, network device association cannot be achieved. Moreover, pinging from a rootless container does not work because it lacks the CAP\_NET\_RAW security capability that the ping command requires." But it's still better than no connectivity at all.

无根网络[相当有限](https://www.redhat.com/sysadmin/container-networking-podman)：“从技术上讲，容器本身没有IP地址，因为没有root权限，网络设备关联不能实现。此外，从无根容器 ping 不起作用，因为它缺少 ping 命令所需的 CAP\_NET\_RAW 安全功能。”但这仍然比完全没有连接要好。

## Conclusion

##  结论

The considered in this article approach to organizing container networking is only one of the possible ways (well, probably the most widely used one). There are many more other ways, implemented as official or 3rd party plugins, but all of them heavily rely on [Linux network virtualization facilities](https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking/). Thus, containerization can fairly be regarded as virtualization technology.

本文中考虑的组织容器网络的方法只是其中一种可能的方法（嗯，可能是使用最广泛的一种）。还有更多其他方式，作为官方或第 3 方插件实现，但它们都严重依赖 [Linux 网络虚拟化设施](https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking/)。因此，容器化可以说是一种虚拟化技术。

Make code, not war!

编写代码，而不是战争！

### References:

###  参考：

- [Introduction to Linux interfaces for virtual networking](https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking/)
- [ip(8) — Linux manual page](https://man7.org/linux/man-pages/man8/ip.8.html)
- ⭐[Linux Bridge - Part 1](https://hechao.li/2017/12/13/linux-bridge-part1/) by [Hechao Li](https://hechao.li/aboutme/) (and more techical [Part 2](https://hechao.li/2018/01/31/linux-bridge-part2/))
- [Anatomy of a Linux bridge](https://wiki.aalto.fi/download/attachments/70789083/linux_bridging_final.pdf)
- [A container networking overview](https://jvns.ca/blog/2016/12/22/container-networking/)
- [Demystifying container networking](https://blog.mbrt.dev/2017-10-01-demystifying-container-networking/)
- [Introducing Linux Network Namespaces](https://blog.scottlowe.org/2013/09/04/introducing-linux-network-namespaces/)
- 🎥[Container Networking From Scratch](https://www.youtube.com/watch?v=6v_BDHIgOY8&list=WL&index=2&t=0s)

- [Linux 虚拟网络接口介绍](https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking/)
- [ip(8) — Linux 手册页](https://man7.org/linux/man-pages/man8/ip.8.html)
- ⭐[Linux Bridge - Part 1](https://hechao.li/2017/12/13/linux-bridge-part1/) by [Hechao Li](https://hechao.li/aboutme/)（和更多技术[第2部分](https://hechao.li/2018/01/31/linux-bridge-part2/))
- [Linux 网桥剖析](https://wiki.aalto.fi/download/attachments/70789083/linux_bridging_final.pdf)
- [容器网络概述](https://jvns.ca/blog/2016/12/22/container-networking/)
- [揭秘容器网络](https://blog.mbrt.dev/2017-10-01-demystifying-container-networking/)
- [Linux 网络命名空间介绍](https://blog.scottlowe.org/2013/09/04/introducing-linux-network-namespaces/)
- 🎥[从头开始容器网络](https://www.youtube.com/watch?v=6v_BDHIgOY8&list=WL&index=2&t=0s)

### More [Networking](http://iximiuz.com/en/categories/?category=Networking) articles

### 更多[网络](http://iximiuz.com/en/categories/?category=Networking) 文章

- [How to Expose Multiple Containers On the Same Port](http://iximiuz.com/en/posts/multiple-containers-same-port-reverse-proxy/)
- [Computer Networking Introduction - Ethernet and IP (Heavily Illustrated)](http://iximiuz.com/en/posts/computer-networking-101/) 

- [如何在同一个端口上暴露多个容器](http://iximiuz.com/en/posts/multiple-containers-same-port-reverse-proxy/)
- [计算机网络介绍 - 以太网和 IP（大量插图）](http://iximiuz.com/en/posts/computer-networking-101/)

- [Bridge vs Switch: What I Learned From a Data Center Tour](http://iximiuz.com/en/posts/bridge-vs-switch/)
- [Illustrated introduction to Linux iptables](http://iximiuz.com/en/posts/laymans-iptables-101/)

- [Bridge vs Switch：我从数据中心之旅中学到的东西](http://iximiuz.com/en/posts/bridge-vs-switch/)
-  [Linux iptables 图解介绍](http://iximiuz.com/en/posts/laymans-iptables-101/)

### Other [Containers](http://iximiuz.com/en/categories/?category=Containers) articles

### 其他 [容器](http://iximiuz.com/en/categories/?category=Containers) 文章

- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)

- [并非每个容器内部都有操作系统](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [你不需要图像来运行容器](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [你需要容器来构建镜像](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)

[docker,](javascript: void 0) [linux,](javascript: void 0) [container](javascript: void 0)

[docker,](javascript: void 0) [linux,](javascript: void 0) [container](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

