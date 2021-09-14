# Container Networking Is Simple!

October 18, 2020 (Updated: August 4, 2021)

_**Just kidding, it's not... But fear not and read on!**_

_You can find a Russian translation of this article [here](https://habr.com/ru/company/timeweb/blog/558612/)_.

Working with containers always feels like magic. In a good way for those who understand the internals and in a terrifying - for those who don't. Luckily, we've been looking under the hood of the containerization technology for quite some time already and even managed to uncover that [containers are just isolated and restricted Linux processes](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/#container-is-just-a-processes), that [images aren't really needed to run containers](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/), and on the contrary - [to build an image we need to run some containers](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/).

Now comes a time to tackle the container networking problem. Or, more precisely, a single-host container networking problem. In this article, we are going to answer the following questions:

- How to virtualize network resources to make containers think each of them has a dedicated network stack?
- How to turn containers into friendly neighbors, prevent them from interfering, and teach to communicate well?
- How to reach the outside world (e.g. the Internet) from inside the container?
- How to reach containers running on a machine from the outside world ( _aka_ port publishing)?

While answering these questions, we'll setup a container networking from scratch using standard Linux tools. As a result, it'll become apparent that the single-host container networking is nothing more than a simple combination of the well-known Linux facilities:

- network namespaces;
- virtual Ethernet devices (veth);
- virtual network switches (bridge);
- IP routing and network address translation (NAT).

And for better or worse, no code is required to make the networking magic happen...

## Prerequisites

Any decent Linux distribution would probably suffice. All the examples in the article have been made on a fresh _vagrant_ CentOS 8 virtual machine:

```bash
$ vagrant init centos/8
$ vagrant up
$ vagrant ssh

[vagrant@localhost ~]$ uname -a
Linux localhost.localdomain 4.18.0-147.3.1.el8_1.x86_64

```

For the sake of simplicity of the examples, in this article, we are not going to rely on any fully-fledged containerization solution (e.g. _docker_ or _podman_). Instead, we'll focus on the basic concepts and use the bare minimum tooling to achieve our learning goals.

## Isolating containers with network namespaces

What constitutes a Linux network stack? Well, obviously, the set of network devices. What else? Probably, the set of routing rules. And not to forget, the set of netfilter hooks, including defined by iptables rules.

We can quickly forge a non-comprehensive `inspect-net-stack.sh` script:

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

```bash
$ sudo iptables -N ROOT_NS

```

After that, execution of the inspect script on my machine produces the following output:

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

Well, you might have heard already, that one of the Linux namespaces used for containers isolation is called _network namespace_. From [`man ip-netns`](https://man7.org/linux/man-pages/man8/ip-netns.8.html), _"network namespace is logically another copy of the network stack, with its own routes, firewall rules, and network devices."_ For the sake of simplicity, this is the only namespace we're going to use in this article. Instead of creating fully-isolated containers, we'd rather restrict the scope to only the network stack.

One of the ways to create a network namespace is the `ip` tool - part of the de facto standard [iproute2](https://en.wikipedia.org/wiki/Iproute2) collection:

```bash
$ sudo ip netns add netns0
$ ip netns
netns0

```

How to start using the just created namespace? There is a lovely Linux command called `nsenter`. It enters one or more of the specified namespaces and then executes the given program:

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

![Linux network namespace visualized](http://iximiuz.com/container-networking-is-simple/network-namespace-4000-opt.png)

_Network namespace visualized._

## Connecting containers to host with virtual Ethernet devices (veth)

A dedicated network stack would be not so useful if we could not communicate with it. Luckily, Linux provides a suitable facility for that - a virtual Ethernet device! From [`man veth`](https://man7.org/linux/man-pages/man4/veth.4.html), _"veth devices are virtual Ethernet devices. They can act as tunnels between network namespaces to create a bridge to a physical network device in another namespace, but can also be used as standalone network devices."_

Virtual Ethernet devices always go in pairs. No worries, it'll be clear when we take a look at the creation command:

```bash
$ sudo ip link add veth0 type veth peer name ceth0

```

With this single command, we just created a pair of _interconnected_ virtual Ethernet devices. The names `veth0` and `ceth0` have been chosen arbitrarily:

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

```bash
$ sudo ip link set veth0 up
$ sudo ip addr add 172.18.0.11/16 dev veth0

```

And continue with the `netns0`:

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

We are ready to check the connectivity:

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

```bash
# From `netns0` namespace:
$ ip route
172.18.0.0/16 dev ceth0 proto kernel scope link src 172.18.0.10

```

Linux has a bunch of ways to populate the routing table. One of them is to extract routes from the directly attached network interfaces. Remember, the routing table in `netns0` was empty right after the namespace creation. But then we added the `ceth0` device there and assigned it an IP address `172.18.0.10/16`. Since we were using not a simple IP address, but a combination of the address and the netmask, the network stack managed to extract the routing information from it. Every packet destined to `172.18.0.0/16` network will be sent through `ceth0` device. But any other packets will be discarded. Similarly, there is a new route in the root namespace:

```bash
# From `root` namespace:
$ ip route
# ... omitted lines ...
172.18.0.0/16 dev veth0 proto kernel scope link src 172.18.0.11

```

At this point, we are ready to mark our very first question answered. **We know now how to isolate, virtualize, and connect Linux network stacks.**

## Interconnecting containers with virtual network switch (bridge)

The whole idea of containerization boils down to efficient resource sharing. I.e. it's uncommon to have a single container per machine. Instead, the goal is to run as many isolated processes in the shared environment as possible. So, what'd happen if we were to place multiple containers on the same host following the `veth` approach from above? Let's add the second _container_:

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

Well, it took me some time to figure it out, but apparently we are facing the clash of routes. Let's inspect the routing table in the root namespace:

```bash
$ ip route
# ... omitted lines ...
172.18.0.0/16 dev veth0 proto kernel scope link src 172.18.0.11
172.18.0.0/16 dev veth1 proto kernel scope link src 172.18.0.21

```

Even though after adding the second `veth` pair, root's network stack learned the new route `172.18.0.0/16 dev veth1 proto kernel scope link src 172.18.0.21`, there already was an existing route for exactly the same network. When the second container tries to ping `veth1` device, the first route is being selected breaking the connectivity. If we were to delete the first route `sudo ip route delete 172.18.0.0/16 dev veth0 proto kernel scope link src 172.18.0.11` and recheck the connectivity, the situation would turn into a mirrored case. I.e. the connectivity of the `netns1` would be restored, but `netns0` would be in limbo.

![Connecting multiple network namespaces with a bridge](http://iximiuz.com/container-networking-is-simple/route-clash-4000-opt.png)

Well, I believe if we selected another IP network for `netns1`, everything would work. However, multiple containers sitting in one IP network is a legitimate use case. Thus, we need to adjust the `veth` approach somehow...

Behold the Linux bridge - yet another virtualized network facility! The Linux bridge behaves like a network switch. It forwards packets between interfaces that are connected to it. And since it's a switch, it does it on the L2 (i.e. Ethernet) level.

Let's try to play with our new toy. But first, we need to clean up the existing setup because some of the configurational changes we've made so far aren't really needed anymore. Removing network namespaces would suffice:

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

```bash
$ ip route
default via 10.0.2.2 dev eth0 proto dhcp metric 100
10.0.2.0/24 dev eth0 proto kernel scope link src 10.0.2.15 metric 100

```

And finally, create the bridge interface:

```bash
$ sudo ip link add br0 type bridge
$ sudo ip link set br0 up

```

Now, attach `veth0` and `veth1` ends to the bridge:

```bash
$ sudo ip link set veth0 master br0
$ sudo ip link set veth1 master br0

```

![Setting up routing between multiple network namespaces](http://iximiuz.com/container-networking-is-simple/bridge-4000-opt.png)

...and check the connectivity between containers:

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

## Reaching out to the outside world (IP routing and masquerading)

Our containers can talk to each other. But can they talk to the host, i.e. the root namespace?

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ping 10.0.2.15  # eth0 address
connect: Network is unreachable

```

That's kind of obvious, there is simply no route for that in `netns0`:

```bash
$ ip route
172.18.0.0/16 dev ceth0 proto kernel scope link src 172.18.0.10

```

The root namespace cannot talk to containers either:

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

```bash
$ sudo ip addr add 172.18.0.1/16 dev br0

```

Once we assigned the IP address to the bridge interface, we got a route on the host routing table:

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

![Using bridge as a gateway](http://iximiuz.com/container-networking-is-simple/router-4000-opt.png)

Perfect, we connected containers with the root namespace. Now, let's try to connect them to the outside world. By default, the packet forwarding (i.e. the router functionality) is disabled in Linux. We need to turn it on:

```bash
# In the root namespace
sudo bash -c 'echo 1 > /proc/sys/net/ipv4/ip_forward'

```

Again, my favourite part - checking the connectivity:

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ ping 8.8.8.8
# hangs indefinitely long for me...

```

Well, still doesn't work. What have we missed? If the container were to sends packets to the outside world, the destination server would not be able to send packets back to the container because the container's IP address is private. I.e. the routing rules for that particular IP are known only to the local network. And lots of the containers in the world share exactly the same private IP address `172.18.0.10`. The solution to this problem is called the [Network address translation (NAT)](https://en.wikipedia.org/wiki/Network_address_translation). Before going to the external network, packets originated by the containers will get their source IP addresses replaced with the host's external interface address. The host also will track all the existing mappings and on arrival, it'll be restoring the IP addresses before forwarding packets back to the containers. Sounds complicated, but I have good news for you! Thanks to [iptables](http://iximiuz.com/en/posts/laymans-iptables-101/) module, we need only a single command to make it happen:

```bash
$ sudo iptables -t nat -A POSTROUTING -s 172.18.0.0/16 ! -o br0 -j MASQUERADE

```

The command is fairly simple. We are adding a new rule to the `nat` table of the `POSTROUTING` chain asking to masquerade all the packets originated in `172.18.0.0/16` network, but not by the bridge interface.

Check the connectivity:

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

```bash
sudo iptables -S
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT

```

As a good example, Docker, instead, restricts everything by default and then enables routing for only known paths.

Click here to see Docker iptables rules.

The following are the dumped rules generated by the Docker daemon on a _CentOS 8_ machine with single container exposed on port 5005:

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
-A FORWARD -i docker0 ! -o docker0 -j ACCEPT
-A FORWARD -i docker0 -o docker0 -j ACCEPT
-A DOCKER -d 172.17.0.2/32 ! -i docker0 -o docker0 -p tcp -m tcp --dport 5000 -j ACCEPT
-A DOCKER-ISOLATION-STAGE-1 -i docker0 ! -o docker0 -j DOCKER-ISOLATION-STAGE-2
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
-A POSTROUTING -s 172.17.0.0/16 ! -o docker0 -j MASQUERADE
-A POSTROUTING -s 172.17.0.2/32 -d 172.17.0.2/32 -p tcp -m tcp --dport 5000 -j MASQUERADE
-A OUTPUT ! -d 127.0.0.0/8 -m addrtype --dst-type LOCAL -j DOCKER
-A DOCKER -i docker0 -j RETURN
-A DOCKER ! -i docker0 -p tcp -m tcp --dport 5005 -j DNAT --to-destination 172.17.0.2:5000

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

It's a known practice to _publish_ container ports to some (or all) of the host's interfaces. But what does _port publishing_ really mean?

Imagine we have a server running inside a container:

```bash
$ sudo nsenter --net=/var/run/netns/netns0
$ python3 -m http.server --bind 172.18.0.10 5000

```

If we try to send an HTTP request to this server process from the host, everything will work (well, there is a connectivity between root namespace and all the container interfaces, why wouldn't it?):

```bash
# From root namespace
$ curl 172.18.0.10:5000
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
# ... omitted lines ...

```

However, if we were to access this server from the outside world, what IP address would we use? The only IP address we might know is the host's external interface address `eth0`:

```bash
$ curl 10.0.2.15:5000
curl: (7) Failed to connect to 10.0.2.15 port 5000: Connection refused

```

Thus, we need to find a way to forward any packets arriving at port 5000 on the host's `eth0` interface to `172.18.0.10:5000` destination. Or, in other words, we need to _publish_ the container's port 5000 on the host's `eth0` interface. iptables to the rescue!

```bash
# External traffic
sudo iptables -t nat -A PREROUTING -d 10.0.2.15 -p tcp -m tcp --dport 5000 -j DNAT --to-destination 172.18.0.10:5000

# Local traffic (since it doesn't pass the PREROUTING chain)
sudo iptables -t nat -A OUTPUT -d 10.0.2.15 -p tcp -m tcp --dport 5000 -j DNAT --to-destination 172.18.0.10:5000

```

Additionally, we need to enable [iptables intercepting traffic over bridged networks](https://github.com/omribahumi/libvirt_metadata_api/pull/4/files):

```bash
sudo modprobe br_netfilter

```

Testing time!

```bash
curl 10.0.2.15:5000
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
# ... omitted lines ...

```

## Understanding Docker network drivers

Ok, sir, what can we do with all this useless knowledge? For instance, we could try to understand some of the [Docker network modes](https://docs.docker.com/network/#network-drivers)!

Let's start from the `--network host` mode. Try to compare the output of the following commands `ip link` and `sudo docker run -it --rm --network host alpine ip link`. Surprise, surprise, they are exactly the same! I.e. in the `host` mode, Docker simply doesn't use the network namespace isolation and containers work in the root network namespace and share the network stack with the host machine.

The next mode to inspect is `--network none`. The output of the `sudo docker run -it --rm --network none alpine ip link` command shows only a single loopback network interface. It's very similar to our observations of the freshly created network namespace. I.e. before the point where we were adding any `veth` devices.

Last but not least, the `--network bridge` (the default) mode. Well, it's exactly what we've been trying to reproduce in this whole article. I encourage you to play with `ip` and `iptables` commands and inspect the network stack from the host and containers point of view.

## Bonus: rootless containers and networking

One of the nice features of `podman` container manager is its focus on rootless containers. However, as you probably noticed, we used a lot of `sudo` escalations in this article. I.e. it's impossible to configure the network without root privileges. [Podman's approach](https://www.redhat.com/sysadmin/container-networking-podman) to rootful networking is very close to docker. But when it comes to rootless containers, podman relies on [slirp4netns](https://github.com/rootless-containers/slirp4netns) project:

> Starting with Linux 3.8, unprivileged users can create network\_namespaces(7) along with user\_namespaces(7). However, unprivileged network namespaces had not been very useful, because creating veth(4) pairs across the host and network namespaces still requires the root privileges. (i.e. No internet connection)

> slirp4netns allows connecting a network namespace to the Internet in a completely unprivileged way, by connecting a TAP device in a network namespace to the usermode TCP/IP stack ("slirp").

The rootless networking [is quite limited](https://www.redhat.com/sysadmin/container-networking-podman): "technically, the container itself does not have an IP address, because without root privileges, network device association cannot be achieved. Moreover, pinging from a rootless container does not work because it lacks the CAP\_NET\_RAW security capability that the ping command requires." But it's still better than no connectivity at all.

## Conclusion

The considered in this article approach to organizing container networking is only one of the possible ways (well, probably the most widely used one). There are many more other ways, implemented as official or 3rd party plugins, but all of them heavily rely on [Linux network virtualization facilities](https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking/). Thus, containerization can fairly be regarded as virtualization technology.

Make code, not war!

### References:

- [Introduction to Linux interfaces for virtual networking](https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking/)
- [ip(8) — Linux manual page](https://man7.org/linux/man-pages/man8/ip.8.html)
- ⭐[Linux Bridge - Part 1](https://hechao.li/2017/12/13/linux-bridge-part1/) by [Hechao Li](https://hechao.li/aboutme/) (and more techical [Part 2](https://hechao.li/2018/01/31/linux-bridge-part2/))
- [Anatomy of a Linux bridge](https://wiki.aalto.fi/download/attachments/70789083/linux_bridging_final.pdf)
- [A container networking overview](https://jvns.ca/blog/2016/12/22/container-networking/)
- [Demystifying container networking](https://blog.mbrt.dev/2017-10-01-demystifying-container-networking/)
- [Introducing Linux Network Namespaces](https://blog.scottlowe.org/2013/09/04/introducing-linux-network-namespaces/)
- 🎥[Container Networking From Scratch](https://www.youtube.com/watch?v=6v_BDHIgOY8&list=WL&index=2&t=0s)

### More [Networking](http://iximiuz.com/en/categories/?category=Networking) articles

- [How to Expose Multiple Containers On the Same Port](http://iximiuz.com/en/posts/multiple-containers-same-port-reverse-proxy/)
- [Computer Networking Introduction - Ethernet and IP (Heavily Illustrated)](http://iximiuz.com/en/posts/computer-networking-101/)
- [Bridge vs Switch: What I Learned From a Data Center Tour](http://iximiuz.com/en/posts/bridge-vs-switch/)
- [Illustrated introduction to Linux iptables](http://iximiuz.com/en/posts/laymans-iptables-101/)

### Other [Containers](http://iximiuz.com/en/categories/?category=Containers) articles

- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)

[docker,](javascript: void 0) [linux,](javascript: void 0) [container](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

