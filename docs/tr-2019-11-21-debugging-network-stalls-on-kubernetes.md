# Debugging network stalls on Kubernetes

# 在 Kubernetes 上调试网络停顿

We’ve talked about [Kubernetes](https://github.blog/2017-08-16-kubernetes-at-github/) before, and over the last couple of years it’s become the standard deployment pattern at GitHub. We now run a large portion of both internal and public-facing services on Kubernetes. As our Kubernetes clusters have grown, and our targets on the latency of our services have become more stringent, we began to notice that certain services running on Kubernetes in our environment were experiencing sporadic latency that couldn't be attributed to the performance characteristics of the application itself.

我们之前讨论过 [Kubernetes](https://github.blog/2017-08-16-kubernetes-at-github/)，在过去的几年里，它已经成为 GitHub 的标准部署模式。我们现在在 Kubernetes 上运行大部分内部和面向公众的服务。随着我们的 Kubernetes 集群的增长，我们对服务延迟的目标变得更加严格，我们开始注意到在我们的环境中运行在 Kubernetes 上的某些服务正在经历零星的延迟，这不能归因于应用程序本身。

Essentially, applications running on our Kubernetes clusters would observe seemingly random latency of up to and over 100ms on connections, which would cause downstream timeouts or retries. Services were expected to be able to respond to requests in well under 100ms, which wasn’t feasible when the connection itself was taking so long. Separately, we also observed very fast MySQL queries, which we expected to take a matter of milliseconds and that MySQL observed taking only milliseconds, were being observed taking 100ms or more from the perspective of the querying application.

本质上，在我们的 Kubernetes 集群上运行的应用程序会观察到看似随机的连接延迟高达 100 毫秒以上，这会导致下游超时或重试。预计服务能够在不到 100 毫秒的时间内响应请求，当连接本身需要这么长时间时，这是不可行的。另外，我们还观察到非常快的 MySQL 查询，我们预计这需要几毫秒，而 MySQL 观察到的仅需要几毫秒，从查询应用程序的角度来看，它需要 100 毫秒或更长时间。

The problem was initially narrowed down to communications that involved a Kubernetes node, even if the other side of a connection was outside Kubernetes. The most simple reproduction we had was a [Vegeta](https://github.com/tsenart/vegeta) benchmark that could be run from any internal host, targeting a Kubernetes service running on a node port, and would observe the sporadically high latency. In this post, we’ll walk through how we tracked down the underlying issue.

问题最初被缩小到涉及 Kubernetes 节点的通信，即使连接的另一端在 Kubernetes 之外。我们拥有的最简单的复制是 [Vegeta](https://github.com/tsenart/vegeta) 基准测试，它可以从任何内部主机运行，针对在节点端口上运行的 Kubernetes 服务，并且会观察到零星的高潜伏。在这篇文章中，我们将介绍我们如何追踪潜在问题。

## Removing complexity to find the path at fault

## 去除复杂性以找到故障路径

Using an example reproduction, we wanted to narrow down the problem and remove layers of complexity. Initially, there were too many moving parts in the flow between Vegeta and pods running on Kubernetes to determine if this was a deeper network problem, so we needed to rule some out.

使用示例复制，我们希望缩小问题的范围并消除复杂性。最初，在 Vegeta 和在 Kubernetes 上运行的 Pod 之间的流程中有太多的移动部分，无法确定这是否是更深层次的网络问题，因此我们需要排除一些。

![vegeta to kube pod via NodePort](https://i1.wp.com/user-images.githubusercontent.com/349190/66730905-2ede2800-eea0-11e9-9091-5e68bb23467d.png?resize=816%2C210&ssl=1)

1)

The client, Vegeta, creates a TCP connection to any kube-node in the cluster. Kubernetes runs in our data centers as an [overlay network](https://en.wikipedia.org/wiki/Overlay_network) (a network that runs on top of our existing datacenter network) that uses [IPIP](https://en.wikipedia.org/wiki/IP_in_IP) (which encapsulates the overlay network's IP packet inside the datacenter's IP packet). When a connection is made to that first kube-node, it performs stateful [Network Address Translation](https://en.wikipedia.org/wiki/Network_address_translation)(NAT) to convert the kube-node's IP and port to an IP and port on the overlay network (specifically, of the pod running the application). On return, it undoes each of these steps. This is a complex system with a lot of state, and a lot of moving parts that are constantly updating and changing as services deploy and move around.

客户端 Vegeta 创建到集群中任何 kube 节点的 TCP 连接。 Kubernetes 作为 [覆盖网络](https://en.wikipedia.org/wiki/Overlay_network)（运行在我们现有数据中心网络之上的网络）在我们的数据中心内运行，该网络使用 [IPIP](https://en.wikipedia.org/wiki/IP_in_IP)（将覆盖网络的 IP 数据包封装在数据中心的 IP 数据包中）。当与第一个 kube-node 建立连接时，它会执行有状态的 [网络地址转换](https://en.wikipedia.org/wiki/Network_address_translation)(NAT) 以将 kube-node 的 IP 和端口转换为 IP和覆盖网络上的端口（特别是运行应用程序的 pod)。返回时，它会撤消这些步骤中的每一个。这是一个复杂的系统，有很多状态，以及随着服务部署和移动而不断更新和变化的许多活动部分。

As part of running a `tcpdump` on the original Vegeta benchmark, we observed the latency during a TCP handshake (between SYN and SYN-ACK). To simplify some of the complexity of HTTP and Vegeta, we can use `hping3` to just “ping” with a SYN packet and see if we observe the latency in the response packet—then throw away the connection. We can filter it to only include packets over 100ms and get a simpler reproduction case than a full Layer 7 Vegeta benchmark or attack against the service. The following “pings” a kube-node using TCP SYN/SYN-ACK on the “node port” for the service (30927) with an interval of 10ms, filtered for slow responses.

作为在原始 Vegeta 基准测试上运行 `tcpdump` 的一部分，我们观察了 TCP 握手期间（SYN 和 SYN-ACK 之间）的延迟。为了简化 HTTP 和 Vegeta 的一些复杂性，我们可以使用 `hping3` 来“ping”一个 SYN 数据包，看看我们是否观察到响应数据包中的延迟——然后丢弃连接。我们可以过滤它以仅包含超过 100 毫秒的数据包，并获得比完整的第 7 层 Vegeta 基准测试或针对服务的攻击更简单的复制案例。下面使用 TCP SYN/SYN-ACK 在服务 (30927) 的“节点端口”上“ping”一个 kube 节点，间隔为 10 毫秒，过滤慢响应。

```shell
theojulienne@shell ~ $ sudo hping3 172.16.47.27 -S -p 30927 -i u10000 |egrep --line-buffered 'rtt=[0-9]{3}\.'
len=46 ip=172.16.47.27 ttl=59 DF id=0 sport=30927 flags=SA seq=1485 win=29200 rtt=127.1 ms
len=46 ip=172.16.47.27 ttl=59 DF id=0 sport=30927 flags=SA seq=1486 win=29200 rtt=117.0 ms
len=46 ip=172.16.47.27 ttl=59 DF id=0 sport=30927 flags=SA seq=1487 win=29200 rtt=106.2 ms
len=46 ip=172.16.47.27 ttl=59 DF id=0 sport=30927 flags=SA seq=1488 win=29200 rtt=104.1 ms
```


Our first new observation from the sequence numbers and timings is that this isn’t a one-off, but is often grouped, like a backlog that eventually gets processed.

我们对序列号和时间的第一个新观察是，这不是一次性的，而是经常分组的，就像最终得到处理的积压。

Next up, we want to narrow down which component(s) were potentially at fault. Is it the kube-proxy iptables NAT rules that are hundreds of rules long? Is it the IPIP tunnel and something on the network handling them poorly? One way to validate this is to test each step of the system. What happens if we remove the NAT and firewall logic and only use the IPIP part.

接下来，我们要缩小哪些组件可能有故障的范围。是数百条规则的 kube-proxy iptables NAT 规则吗？是 IPIP 隧道和网络上的某些东西处理它们很差吗？验证这一点的一种方法是测试系统的每个步骤。如果我们移除 NAT 和防火墙逻辑而只使用 IPIP 部分会发生什么。

![hping 3 on the IPIP tunnel alone](https://i0.wp.com/user-images.githubusercontent.com/349190/66730912-37cef980-eea0-11e9-8048-984d6c5966a3.png?resize=570%2C210&ssl=1)

Linux thankfully lets you just talk directly to an overlay IP when you’re on a machine that’s part of the same network, so that’s pretty easy to do.

幸运的是，当您在同一网络中的机器上时，Linux 允许您直接与覆盖 IP 对话，因此这很容易做到。

```shell
theojulienne@kube-node-client ~ $ sudo hping3 10.125.20.64 -S -i u10000 |egrep --line-buffered 'rtt=[0-9]{3}\.'
len=40 ip=10.125.20.64 ttl=64 DF id=0 sport=0 flags=RA seq=7346 win=0 rtt=127.3 ms
len=40 ip=10.125.20.64 ttl=64 DF id=0 sport=0 flags=RA seq=7347 win=0 rtt=117.3 ms
len=40 ip=10.125.20.64 ttl=64 DF id=0 sport=0 flags=RA seq=7348 win=0 rtt=107.2 ms
```


Based on our results, the problem still remains! That rules out iptables and NAT. Is it TCP that’s the problem? Let’s see what happens when we perform a normal ICMP ping.

根据我们的结果，问题仍然存在！这排除了 iptables 和 NAT。这是 TCP 的问题吗？让我们看看当我们执行正常的 ICMP ping 时会发生什么。

```shell
theojulienne@kube-node-client ~ $ sudo hping3 10.125.20.64 --icmp -i u10000 |egrep --line-buffered 'rtt=[0-9]{3}\.'
len=28 ip=10.125.20.64 ttl=64 id=42594 icmp_seq=104 rtt=110.0 ms
len=28 ip=10.125.20.64 ttl=64 id=49448 icmp_seq=4022 rtt=141.3 ms
len=28 ip=10.125.20.64 ttl=64 id=49449 icmp_seq=4023 rtt=131.3 ms
len=28 ip=10.125.20.64 ttl=64 id=49450 icmp_seq=4024 rtt=121.2 ms
```


Our results show that the problem still exists. Is it the IPIP tunnel that’s causing the problem? Let’s simplify things further.

我们的结果表明问题仍然存在。是 IPIP 隧道导致了问题吗？让我们进一步简化事情。

![hping 3 directly between hosts](https://i1.wp.com/user-images.githubusercontent.com/349190/66730916-3c93ad80-eea0-11e9-810f-0f32da529da0.png?resize=570%2C210&ssl=1)



Is it possible that it’s every packet between these two hosts?

是否有可能是这两个主机之间的每个数据包？

```shell
theojulienne@kube-node-client ~ $ sudo hping3 172.16.47.27 --icmp -i u10000 |egrep --line-buffered 'rtt=[0-9]{3}\.'
len=46 ip=172.16.47.27 ttl=61 id=41127 icmp_seq=12564 rtt=140.9 ms
len=46 ip=172.16.47.27 ttl=61 id=41128 icmp_seq=12565 rtt=130.9 ms
len=46 ip=172.16.47.27 ttl=61 id=41129 icmp_seq=12566 rtt=120.8 ms
len=46 ip=172.16.47.27 ttl=61 id=41130 icmp_seq=12567 rtt=110.8 ms
```


Behind the complexity, it’s as simple as two kube-node hosts sending any packet, even ICMP pings, to each other. They’ll still see the latency, if the target host is a “bad” one (some are worse than others).

在复杂性背后，它就像两个 kube-node 主机相互发送任何数据包，甚至 ICMP ping 一样简单。如果目标主机是“坏”主机（有些比其他主机更糟糕），他们仍然会看到延迟。

Now there’s one last thing to question: we clearly don’t observe this everywhere, so why is it just on kube-node servers? And does it occur when the kube-node is the sender or the receiver? Luckily, this is also pretty easy to narrow down by using a host outside Kubernetes as a sender, but with the same “known bad” target host (from a staff shell host to the same kube-node). We can observe this is still an issue in that direction.

现在还有最后一件事需要质疑：我们显然不会在任何地方观察到这一点，那么为什么只在 kube-node 服务器上观察到呢？当 kube-node 是发送者还是接收者时，它会发生吗？幸运的是，通过使用 Kubernetes 外部的主机作为发送方，但使用相同的“已知坏”目标主机（从员工 shell 主机到同一个 kube 节点），这也很容易缩小范围。我们可以观察到这仍然是那个方向的问题。

```shell
theojulienne@shell ~ $ sudo hping3 172.16.47.27 -p 9876 -S -i u10000 |egrep --line-buffered 'rtt=[0-9]{3}\.'
len=46 ip=172.16.47.27 ttl=61 DF id=0 sport=9876 flags=RA seq=312 win=0 rtt=108.5 ms
len=46 ip=172.16.47.27 ttl=61 DF id=0 sport=9876 flags=RA seq=5903 win=0 rtt=119.4 ms
len=46 ip=172.16.47.27 ttl=61 DF id=0 sport=9876 flags=RA seq=6227 win=0 rtt=139.9 ms
len=46 ip=172.16.47.27 ttl=61 DF id=0 sport=9876 flags=RA seq=7929 win=0 rtt=131.2 ms
```


And then perform the same from the previous source kube-node to a staff shell host (which rules out the source host, since a ping has both an RX and TX component).

然后从之前的源 kube-node 到人员 shell 主机（排除源主机，因为 ping 具有 RX 和 TX 组件）执行相同的操作。

```shell
theojulienne@kube-node-client ~ $ sudo hping3 172.16.33.44 -p 9876 -S -i u10000 |egrep --line-buffered 'rtt=[0-9]{3}\.'
^C
--- 172.16.33.44 hping statistic ---
22352 packets transmitted, 22350 packets received, 1% packet loss
round-trip min/avg/max = 0.2/7.6/1010.6 ms
```




Looking into the packet captures of the latency we observed, we get some more information. Specifically, that the “sender” host (bottom) observes this timeout while the “receiver” host (top) does not—see the Delta column (in seconds).

查看我们观察到的延迟的数据包捕获，我们获得了更多信息。具体来说，“发送方”主机（底部）观察到这个超时，而“接收方”主机（顶部）没有——参见增量列（以秒为单位）。

![packet captures of the latency](https://i0.wp.com/user-images.githubusercontent.com/349190/66730925-44ebe880-eea0-11e9-9efd-a533d2531de4.png?resize=3714%2C800&ssl=1)



Additionally, by looking at the difference between the ordering of the packets (based on the sequence numbers) on the receiver side of the TCP and ICMP results above, we can observe that ICMP packets always arrive in the same sequence they were sent, but with uneven timing, while TCP packets are sometimes interleaved, but a subset of them stall. Notably, we observe that if you count the ports of the SYN packets, the ports are not in order on the receiver side, while they’re in order on the sender side.

此外，通过查看上述 TCP 和 ICMP 结果的接收方的数据包排序（基于序列号）之间的差异，我们可以观察到 ICMP 数据包总是以它们发送的相同顺序到达，但是时间不均匀，而 TCP 数据包有时会交错，但其中的一个子集会停滞。值得注意的是，我们观察到，如果您计算 SYN 数据包的端口数，接收方的端口不是有序的，而发送方的端口是有序的。

There is a subtle difference between how modern server [NICs](https://en.wikipedia.org/wiki/Network_interface_controller)—like we have in our data centers—handle packets containing TCP vs ICMP. When a packet arrives, the NIC hashes the packet “per connection” and tries to divvy up the connections across receive queues, each (approximately) delegated to a given CPU core. For TCP, this hash includes both source and destination IP and port. In other words, each connection is hashed (potentially) differently. For ICMP, just the IP source and destination are hashed, since there are no ports.

现代服务器 [NIC](https://en.wikipedia.org/wiki/Network_interface_controller)——就像我们在我们的数据中心——处理包含 TCP 和 ICMP 的数据包之间存在细微的差异。当数据包到达时，NIC 对数据包“每个连接”进行散列，并尝试在接收队列之间分配连接，每个（大约）委派给给定的 CPU 内核。对于 TCP，此散列包括源和目标 IP 和端口。换句话说，每个连接的散列（可能)不同。对于 ICMP，只有 IP 源和目标被散列，因为没有端口。

Another new observation is that we can tell that ICMP observes stalls on all communications between the two hosts during this period from the sequence numbers in ICMP vs TCP, while TCP does not. This tells us that the RX queue hashing is likely in play, almost certainly indicating the stall is in processing RX packets, not in sending responses.

另一个新观察结果是，我们可以从 ICMP 与 TCP 中的序列号看出 ICMP 在此期间观察到两台主机之间所有通信的停顿，而 TCP 则没有。这告诉我们 RX 队列散列可能在起作用，几乎可以肯定表明停顿是在处理 RX 数据包，而不是发送响应。

This rules out kube-node transmits, so we now know that it’s a stall in processing packets, and that it’s on the receive side on some kube-node servers.

这排除了 kube-node 传输，所以我们现在知道它是处理数据包的一个停顿，并且它在一些 kube-node 服务器的接收端。

## Deep dive into Linux kernel packet processing

## 深入了解 Linux 内核数据包处理

To understand why the problem could be on the receiving side on some kube-node servers, let’s take a look at how the Linux kernel processes packets.

要了解为什么某些 kube-node 服务器上的接收端可能会出现问题，让我们来看看 Linux 内核如何处理数据包。

Going back to the simplest traditional implementation, the network card receives a packet and sends an [interrupt](https://en.wikipedia.org/wiki/Interrupt) to the Linux kernel stating that there’s a packet that should be handled. The kernel stops other work, switches context to the interrupt handler, processes the packet, then switches back to what it was doing.

回到最简单的传统实现，网卡接收一个数据包并向 Linux 内核发送一个 [interrupt](https://en.wikipedia.org/wiki/Interrupt)，说明有一个数据包应该被处理。内核停止其他工作，将上下文切换到中断处理程序，处理数据包，然后切换回它正在做的事情。

![traditional interrupt-driven approach](https://i1.wp.com/user-images.githubusercontent.com/349190/66730927-4ae1c980-eea0-11e9-897c-75cd39dbe286.png?resize=743%2C272&ssl=1)



This context switching is slow, which may have been fine on a 10Mbit NIC in the 90s, but on modern servers where the NIC is 10G and at maximal line rate can bring in around 15 million packets per second, on a smaller server with eight cores that could mean the kernel is interrupted millions of times per second per core. 

这种上下文切换速度很慢，这在 90 年代的 10Mbit NIC 上可能很好，但在 NIC 为 10G 且以最大线速运行的现代服务器上，在具有 8 个内核的较小服务器上，每秒可以带来大约 1500 万个数据包这可能意味着内核每内核每秒被中断数百万次。

Instead of constantly handling interrupts, many years ago Linux added [NAPI](https://en.wikipedia.org/wiki/New_API), the networking API that modern drivers use for improved performance at high packet rates. At low rates, the kernel still accepts interrupts from the NIC in the method we mentioned. Once enough packets arrive and cross a threshold, it disables interrupts and instead begins polling the NIC and pulling off packets in batches. This processing is done in a “softirq”, or [software interrupt context](https://www.kernel.org/doc/htmldocs/kernel-hacking/basics-softirqs.html). This happens at the end of syscalls and hardware interrupts, which are times that the kernel (as opposed to userspace) is already running.

许多年前，Linux 并没有不断处理中断，而是添加了 [NAPI](https://en.wikipedia.org/wiki/New_API)，现代驱动程序使用网络 API 来提高高数据包速率下的性能。在低速率下，内核仍然以我们提到的方法接受来自 NIC 的中断。一旦足够的数据包到达并超过阈值，它就会禁用中断，而是开始轮询 NIC 并批量提取数据包。此处理在“softirq”或[软件中断上下文](https://www.kernel.org/doc/htmldocs/kernel-hacking/basics-softirqs.html)中完成。这发生在系统调用和硬件中断结束时，此时内核（而不是用户空间)已经在运行。

![NAPI polling in softirq at end of syscalls](https://i0.wp.com/user-images.githubusercontent.com/349190/66730932-4f0de700-eea0-11e9-927f-393690107ace.png?resize=740%2C293&ssl=1)



This is much faster, but brings up another problem. What happens if we have so many packets to process that we spend all our time processing packets from the NIC, but we never have time to let the userspace processes actually drain those queues (read from TCP connections, etc.)? Eventually the queues would fill up, and we’d start dropping packets. To try and make this fair, the kernel limits the amount of packets processed in a given softirq context to a certain budget. Once this budget is exceeded, it wakes up a separate thread called `ksoftirqd` (you’ll see one of these in `ps` for each core) which processes these softirqs outside of the normal syscall/interrupt path. This thread is scheduled using the standard process scheduler, which already tries to be fair.

这要快得多，但带来了另一个问题。如果我们有太多的数据包要处理，以至于我们将所有时间都花在处理来自 NIC 的数据包上，但我们从来没有时间让用户空间进程实际排空这些队列（从 TCP 连接读取等），会发生什么？最终队列会填满，我们将开始丢弃数据包。为了尽量做到公平，内核将在给定软中断上下文中处理的数据包数量限制在一定的预算内。一旦超出这个预算，它就会唤醒一个名为 `ksoftirqd` 的单独线程（你会在 `ps` 中看到每个内核中的一个），它在正常的系统调用/中断路径之外处理这些软中断。该线程使用标准进程调度程序进行调度，该调度程序已经尝试公平。

![NAPI polling crossing threshold in schedule ksoftirqd](https://i0.wp.com/user-images.githubusercontent.com/349190/66730935-533a0480-eea0-11e9-8697-7fc52d2ea1fa.png?resize=841%2C358&ssl=1)



With an overview of the way the kernel is processing packets, we can see there is definitely opportunity for this processing to become stalled. If the time between softirq processing calls grows, packets could sit in the NIC RX queue for a while before being processed. This could be something deadlocking the CPU core, or it could be something slow preventing the kernel from running softirqs.

通过对内核处理数据包的方式的概述，我们可以看到这种处理肯定有机会停止。如果软中断处理调用之间的时间增加，数据包可能会在处理之前在 NIC RX 队列中停留一段时间。这可能是 CPU 内核死锁的原因，也可能是阻止内核运行软中断的原因。

## Narrow down processing to a core or method

## 将处理范围缩小到核心或方法

At this point, it makes sense that this could happen, and we know we’re observing something that looks a lot like it. The next step is to confirm this theory, and if we do, understand what’s causing it.

在这一点上，这可能发生是有道理的，我们知道我们正在观察看起来很像它的东西。下一步是确认这个理论，如果我们这样做了，了解是什么导致了它。

Let’s revisit the slow round trip packets we saw before.

让我们重新审视我们之前看到的慢速往返数据包。

```shell
len=46 ip=172.16.53.32 ttl=61 id=29573 icmp_seq=1953 rtt=99.3 ms
len=46 ip=172.16.53.32 ttl=61 id=29574 icmp_seq=1954 rtt=89.3 ms
len=46 ip=172.16.53.32 ttl=61 id=29575 icmp_seq=1955 rtt=79.2 ms
len=46 ip=172.16.53.32 ttl=61 id=29576 icmp_seq=1956 rtt=69.1 ms
len=46 ip=172.16.53.32 ttl=61 id=29577 icmp_seq=1957 rtt=59.1 ms
```


As discussed previously, these ICMP packets are hashed to a single NIC RX queue and processed by a single CPU core. If we want to understand what the kernel is doing, it’s helpful to know where (cpu core) and how (softirq, ksoftirqd) it’s processing these packets so we can catch them in action.

如前所述，这些 ICMP 数据包被散列到单个 NIC RX 队列并由单个 CPU 内核处理。如果我们想了解内核在做什么，了解它在哪里（cpu 内核）和如何（softirq，ksoftirqd）处理这些数据包是有帮助的，这样我们就可以在活动中捕获它们。

Now it’s time to use the tools that allow live tracing of a running Linux kernel— [bcc](https://github.com/iovisor/bcc) is what was used here. This allows you to write small C programs that hook arbitrary functions in the kernel, and buffer events back to a userspace Python program which can summarize and return them to you. The “hook arbitrary functions in the kernel” is the difficult part, but it actually goes out of its way to be as safe as possible to use, because it's designed for tracing exactly this type of production issue that you can't simply reproduce in a testing or dev environment. 

现在是时候使用允许实时跟踪正在运行的 Linux 内核的工具了——这里使用的是 [bcc](https://github.com/iovisor/bcc)。这允许您编写小型 C 程序来挂钩内核中的任意函数，并将事件缓冲回用户空间 Python 程序，该程序可以汇总并将它们返回给您。 “钩子内核中的任意函数”是困难的部分，但它实际上竭尽全力尽可能安全地使用，因为它旨在准确跟踪这种类型的生产问题，您不能简单地在其中重现测试或开发环境。

The plan here is simple: we know the kernel is processing those ICMP ping packets, so let's hook the kernel function [icmp\_echo](https://github.com/torvalds/linux/blob/v4.19/net/ipv4/icmp.c#L925) which takes an incoming ICMP “echo request” packet and initiates sending the ICMP “echo response” reply. We can identify the packet using the incrementing icmp\_seq shown by `hping3` above.

这里的计划很简单：我们知道内核正在处理那些 ICMP ping 数据包，所以让我们挂钩内核函数 [icmp\_echo](https://github.com/torvalds/linux/blob/v4.19/net/ipv4/icmp.c#L925) 接收传入的 ICMP“echo request”数据包并开始发送 ICMP“echo response”回复。我们可以使用上面“hping3”所示的递增 icmp\_seq 来识别数据包。

The code for this [bcc script](https://gist.github.com/theojulienne/9d78a0cb68dbe56f19a2ae6316bc6846) looks complex, but breaking it down it’s not as scary as it sounds. The `icmp_echo` function is passed a `struct sk_buff *skb`, which is the packet containing the ICMP echo request. We can delve into this live and pull out the `echo.sequence` (which maps to the `icmp_seq` shown by `hping3` above), and send that back to userspace. Conveniently, we can also grab the current process name/id as well. This gives us results like the following, live as the kernel processes these packets.

这个 [bcc 脚本](https://gist.github.com/theojulienne/9d78a0cb68dbe56f19a2ae6316bc6846) 的代码看起来很复杂，但分解它并不像听起来那么可怕。 `icmp_echo` 函数被传递一个 `struct sk_buff *skb`，它是包含 ICMP 回显请求的数据包。我们可以深入研究这个 live 并拉出 `echo.sequence`（它映射到上面 `hping3` 显示的 `icmp_seq`)，并将其发送回用户空间。方便的是，我们也可以获取当前进程名称/id。这为我们提供了如下结果，在内核处理这些数据包时有效。

```shell
TGID    PID     PROCESS NAME    ICMP_SEQ
0       0       swapper/11      770
0       0       swapper/11      771
0       0       swapper/11      772
0       0       swapper/11      773
0       0       swapper/11      774
20041   20086   prometheus      775
0       0       swapper/11      776
0       0       swapper/11      777
0       0       swapper/11      778
4512    4542   spokes-report-s  779
```


One thing to note about this process name is that in a post-syscall `softirq` context, you see the process that made the syscall show as the “process”, even though really it’s the kernel processing it safely within the kernel context.

关于这个进程名称需要注意的一件事是，在系统调用后的“softirq”上下文中，您会看到使系统调用显示为“进程”的进程，即使实际上是内核在内核上下文中安全地处理它。

With that running, we can now correlate back from the stalled packets observed with `hping3` to the process that’s handling it. A simple `grep` on that capture for the `icmp_seq` values with some context shows what happened before these packets were processed. The packets that line up with the above `hping3` icmp\_seq values have been marked along with the rtt’s we observed above (and what we’d have expected if <50ms rtt’s weren’t filtered out).

运行后，我们现在可以将使用 `hping3` 观察到的停滞数据包与处理它的进程相关联。带有一些上下文的“icmp_seq”值捕获的简单“grep”显示了在处理这些数据包之前发生了什么。与上述 `hping3` icmp\_seq 值对齐的数据包已与我们在上面观察到的 rtt 一起标记（如果 <50ms rtt 没有被过滤掉，我们会期待什么）。

```shell
TGID    PID     PROCESS NAME    ICMP_SEQ ** RTT
--
10137   10436   cadvisor        1951
10137   10436   cadvisor        1952
76      76      ksoftirqd/11    1953 ** 99ms
76      76      ksoftirqd/11    1954 ** 89ms
76      76      ksoftirqd/11    1955 ** 79ms
76      76      ksoftirqd/11    1956 ** 69ms
76      76      ksoftirqd/11    1957 ** 59ms
76      76      ksoftirqd/11    1958 ** (49ms)
76      76      ksoftirqd/11    1959 ** (39ms)
76      76      ksoftirqd/11    1960 ** (29ms)
76      76      ksoftirqd/11    1961 ** (19ms)
76      76      ksoftirqd/11    1962 ** (9ms)
--
10137   10436   cadvisor        2068
10137   10436   cadvisor        2069
76      76      ksoftirqd/11    2070 ** 75ms
76      76      ksoftirqd/11    2071 ** 65ms
76      76      ksoftirqd/11    2072 ** 55ms
76      76      ksoftirqd/11    2073 ** (45ms)
76      76      ksoftirqd/11    2074 ** (35ms)
76      76      ksoftirqd/11    2075 ** (25ms)
76      76      ksoftirqd/11    2076 ** (15ms)
76      76      ksoftirqd/11    2077 ** (5ms)
```




The results tells us a few things. First, these packets are being processed by `ksoftirqd/11` which conveniently tells us this particular pair of machines have their ICMP packets hashed to core 11 on the receiving side. We can also see that every time we see a stall, we always see some packets processed in `cadvisor`'s syscall softirq context, followed by `ksoftirqd` taking over and processing the backlog, exactly the number we'd expect to work through the backlog.

结果告诉我们一些事情。首先，这些数据包正在由 `ksoftirqd/11` 处理，它方便地告诉我们这对特定的机器将它们的 ICMP 数据包散列到接收端的核心 11。我们还可以看到，每次看到停顿时，我们总会看到一些数据包在 `cadvisor` 的 syscall softirq 上下文中处理，然后是 `ksoftirqd` 接管并处理积压，正是我们期望处理的数量积压。

The fact that `cadvisor` is always running just prior to this immediately also implicates it in the problem. Ironically, [cadvisor](https://github.com/google/cadvisor) “analyzes resource usage and performance characteristics of running containers”, yet it’s triggering this performance problem. As with many things related to containers, it’s all relatively bleeding-edge tooling which can result in some somewhat expected corner cases of bad performance.

`cadvisor` 总是在此之前立即运行这一事实也将其与问题联系起来。具有讽刺意味的是，[cadvisor](https://github.com/google/cadvisor)“分析了运行容器的资源使用情况和性能特征”，却引发了这个性能问题。与容器相关的许多事情一样，它都是相对前沿的工具，可能会导致一些预期的性能不佳的极端情况。

## What is cadvisor doing to stall things?

## cadvisor 正在做什么来拖延事情？

With the understanding of how the stall can happen, the process causing it, and the CPU core it’s happening on, we now have a pretty good idea of what this looks like. For the kernel to hard block and not schedule `ksoftirqd` earlier, and given we see packets processed under `cadvisor`'s softirq context, it's likely that `cadvisor` is running a slow syscall which ends with the rest of the packets being processed .

了解了停顿如何发生、导致停顿的过程以及发生停顿的 CPU 内核后，我们现在对停顿的情况有了一个很好的了解。对于内核硬阻塞而不是提前调度 `ksoftirqd`，并且鉴于我们看到在 `cadvisor` 的软中断上下文下处理的数据包，很可能 `cadvisor` 正在运行一个缓慢的系统调用，该系统调用以处理其余数据包结束.

![slow syscall causing stalled packet processing on NIC RX queue](https://i2.wp.com/user-images.githubusercontent.com/349190/66730943-5a611280-eea0-11e9-902a-454f482d4223.png?resize=841%2C345&ssl=1)

841%2C345&ssl=1)

That’s a theory but how do we validate this is actually happening? One thing we can do is trace what’s running on the CPU core throughout this process, catch the point where the packets are overflowing budget and processed by ksoftirqd, then look back a bit to see what was running on the CPU core. Think of it like taking an x-ray of the CPU every few milliseconds. It would look something like this.

这是一个理论，但我们如何验证这是否真的发生了？我们可以做的一件事是在整个过程中跟踪 CPU 内核上运行的内容，捕捉数据包超出预算并由 ksoftirqd 处理的点，然后回头看看 CPU 内核上运行的内容。可以把它想象成每隔几毫秒拍摄一次 CPU 的 X 光片。它看起来像这样。

![tracing of cpu to catch bad syscall and preceding work](https://i2.wp.com/user-images.githubusercontent.com/349190/66730945-5e8d3000-eea0-11e9-89af-daea3a7bd3f6.png?resize=748%2C224&ssl=1)



Conveniently, this is something that’s already mostly supported. The [`perf record`](https://perf.wiki.kernel.org/index.php/Tutorial#Sampling_with_perf_record) tool samples a given CPU core at a certain frequency and can generate a call graph of the live system, including both userspace and the kernel. Taking that recording and manipulating it using a quick fork of a tool from [Brendan Gregg's FlameGraph](https://github.com/brendangregg/FlameGraph) that retained stack trace ordering, we can get a one-line stack trace for each 1ms sample, then get a sample of the 100ms before `ksoftirqd` is in the trace.

方便的是，这是已经得到大部分支持的东西。 [`perf record`](https://perf.wiki.kernel.org/index.php/Tutorial#Sampling_with_perf_record)工具以一定频率对给定的CPU内核进行采样，可以生成实时系统的调用图，包括用户空间和内核。使用保留堆栈跟踪顺序的 [Brendan Gregg's FlameGraph](https://github.com/brendangregg/FlameGraph) 工具的快速分叉记录并操作它，我们可以获得每 1 毫秒的一行堆栈跟踪样本，然后在“ksoftirqd”进入跟踪之前获取 100 毫秒的样本。

```shell
# record 999 times a second, or every 1ms with some offset so not to align exactly with timers
sudo perf record -C 11 -g -F 999
# take that recording and make a simpler stack trace.
sudo perf script 2>/dev/null |./FlameGraph/stackcollapse-perf-ordered.pl |grep ksoftir -B 100
```


This results in the following.

这导致以下结果。

```shell
(hundreds of traces that look similar)

cadvisor;[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];entry_SYSCALL_64_after_swapgs;do_syscall_64;sys_read;vfs_read;seq_read;memcg_stat_show;mem_cgroup_nr_lru_pages;mem_cgroup_node_nr_lru_pages
cadvisor;[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];entry_SYSCALL_64_after_swapgs;do_syscall_64;sys_read;vfs_read;seq_read;memcg_stat_show;mem_cgroup_nr_lru_pages;mem_cgroup_node_nr_lru_pages
cadvisor;[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];entry_SYSCALL_64_after_swapgs;do_syscall_64;sys_read;vfs_read;seq_read;memcg_stat_show;mem_cgroup_iter
cadvisor;[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];entry_SYSCALL_64_after_swapgs;do_syscall_64;sys_read;vfs_read;seq_read;memcg_stat_show;mem_cgroup_nr_lru_pages;mem_cgroup_node_nr_lru_pages
cadvisor;[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];[cadvisor];entry_SYSCALL_64_after_swapgs;do_syscall_64;sys_read;vfs_read;seq_read;memcg_stat_show;mem_cgroup_nr_lru_pages;mem_cgroup_node_nr_lru_pages
ksoftirqd/11;ret_from_fork;kthread;kthread;smpboot_thread_fn;smpboot_thread_fn;run_ksoftirqd;__do_softirq;net_rx_action;ixgbe_poll;ixgbe_clean_rx_irq;napi_gro_receive;netif_receive_skb_internal;inet_gro_receive;bond_handle_frame;__netif_receive_skb_core;ip_rcv_finish;ip_rcv;ip_forward_finish;ip_forward;ip_finish_output;nf_iterate;ip_output;ip_finish_output2;__dev_queue_xmit;dev_hard_start_xmit;ipip_tunnel_xmit;ip_tunnel_xmit;iptunnel_xmit;ip_local_out;dst_output;__ip_local_out;nf_hook_slow;nf_iterate;nf_conntrack_in;generic_packet;ipt_do_table;set_match_v4;ip_set_test;hash_net4_kadt;ixgbe_xmit_frame_ring;swiotlb_dma_mapping_error;hash_net4_test
ksoftirqd/11;ret_from_fork;kthread;kthread;smpboot_thread_fn;smpboot_thread_fn;run_ksoftirqd;__do_softirq;net_rx_action;gro_cell_poll;napi_gro_receive;netif_receive_skb_internal;inet_gro_receive;__netif_receive_skb_core;ip_rcv_finish;ip_rcv;ip_forward_finish;ip_forward;ip_finish_output;nf_iterate;ip_output;ip_finish_output2;__dev_queue_xmit;dev_hard_start_xmit;dev_queue_xmit_nit;packet_rcv;tpacket_rcv;sch_direct_xmit;validate_xmit_skb_list;validate_xmit_skb;netif_skb_features;ixgbe_xmit_frame_ring;swiotlb_dma_mapping_error;__dev_queue_xmit;dev_hard_start_xmit;__bpf_prog_run;__bpf_prog_run
```




There’s a lot there, but looking through it you can see it’s the cadvisor-then-ksoftirqd pattern we saw from the ICMP tracer above. What does it mean?

那里有很多，但仔细看，你可以看到它是我们从上面的 ICMP 跟踪器中看到的 cadvisor-then-ksoftirqd 模式。这是什么意思？

Each line is a trace of the CPU at a point in time. Each call down the stack is separated by `;` on that line. Looking at the middle of the lines we can see the syscall being called is `read(): .... ;do_syscall_64;sys_read; ...` So cadvisor is spending a lot of time in a `read()` syscall relating to `mem_cgroup_*` functions (the top of the call stack / end of line).

每一行都是 CPU 在某个时间点的轨迹。堆栈中的每个调用都在该行以`;` 分隔。查看行的中间，我们可以看到被调用的系统调用是 `read(): .... ;do_syscall_64;sys_read; ...` 所以 cadvisor 在与 `mem_cgroup_*` 函数（调用堆栈的顶部/行尾）相关的 `read()` 系统调用中花费了大量时间。

The call stack trace isn’t convenient to see what’s being read, so let’s use `strace` to see what cadvisor is doing and find 100ms-or-slower syscalls.

调用堆栈跟踪不方便查看正在读取的内容，因此让我们使用 `strace` 来查看 cadvisor 正在做什么并找到 100 毫秒或更慢的系统调用。

```shell
theojulienne@kube-node-bad ~ $ sudo strace -p 10137 -T -ff 2>&1 |egrep '<0\.[1-9]'
[pid 10436] <... futex resumed> )       = 0 <0.156784>
[pid 10432] <... futex resumed> )       = 0 <0.258285>
[pid 10137] <... futex resumed> )       = 0 <0.678382>
[pid 10384] <... futex resumed> )       = 0 <0.762328>
[pid 10436] <... read resumed> "cache 154234880\nrss 507904\nrss_h"..., 4096) = 658 <0.179438>
[pid 10384] <... futex resumed> )       = 0 <0.104614>
[pid 10436] <... futex resumed> )       = 0 <0.175936>
[pid 10436] <... read resumed> "cache 0\nrss 0\nrss_huge 0\nmapped_"..., 4096) = 577 <0.228091>
[pid 10427] <... read resumed> "cache 0\nrss 0\nrss_huge 0\nmapped_"..., 4096) = 577 <0.207334>
[pid 10411] <... epoll_ctl resumed> )   = 0 <0.118113>
[pid 10382] <... pselect6 resumed> )    = 0 (Timeout) <0.117717>
[pid 10436] <... read resumed> "cache 154234880\nrss 507904\nrss_h"..., 4096) = 660 <0.159891>
[pid 10417] <... futex resumed> )       = 0 <0.917495>
[pid 10436] <... futex resumed> )       = 0 <0.208172>
[pid 10417] <... futex resumed> )       = 0 <0.190763>
[pid 10417] <... read resumed> "cache 0\nrss 0\nrss_huge 0\nmapped_"..., 4096) = 576 <0.154442>
```


Sure enough, we see the slow `read()` calls. From the content being read and `mem_cgroup` context above, these `read()` calls are to a `memory.stat` file which shows the memory usage and limits of a cgroup (the resource isolation technology used by Docker). cadvisor is polling this file to get resource utilization details for the containers. Let’s see if it’s the kernel or cadvisor that’s doing something unexpected by attempting the read ourselves.

果然，我们看到了缓慢的 `read()` 调用。从正在读取的内容和上面的 `mem_cgroup` 上下文来看，这些 `read()` 调用是到一个 `memory.stat` 文件，该文件显示了一个 cgroup 的内存使用和限制（Docker 使用的资源隔离技术）。 cadvisor 正在轮询此文件以获取容器的资源利用率详细信息。让我们通过自己尝试读取来看看是内核还是 cadvisor 做了一些意想不到的事情。

```shell
theojulienne@kube-node-bad ~ $ time cat /sys/fs/cgroup/memory/memory.stat >/dev/null

real    0m0.153s
user    0m0.000s
sys    0m0.152s
theojulienne@kube-node-bad ~ $
```


Since we can reproduce it, this indicates that it’s the kernel hitting a pathologically bad case.

由于我们可以重现它，这表明内核遇到了病态的坏情况。

## What causes this read to be so slow [What causes this read to be so slow](http://github.blog\#what-causes-this-read-to-be-so-slow)

## 是什么导致这个读取这么慢[是什么导致这个读取这么慢](http://github.blog\#what-causes-this-read-to-be-so-slow)

At this point it’s much more simple to find similar issues reported by others. As it turns out, this has been reported to cadvisor as an [excessive CPU usage problem](https://github.com/google/cadvisor/issues/1774), it just hadn't been observed that latency was also being introduced to the network stack randomly as well. In fact, some folks internally had noticed cadvisor was consuming more CPU than expected, but it didn’t seem to be causing an issue since our servers had plenty of CPU capacity, and so the CPU usage hadn’t yet been investigated.

在这一点上，找到其他人报告的类似问题要简单得多。事实证明，这已作为 [CPU 使用率过高问题](https://github.com/google/cadvisor/issues/1774) 报告给 cadvisor，只是没有观察到延迟也被引入也随机发送到网络堆栈。事实上，内部有些人已经注意到 cadvisor 消耗的 CPU 比预期的要多，但它似乎没有引起问题，因为我们的服务器有足够的 CPU 容量，因此尚未调查 CPU 使用率。

The overview of the issue is that the memory cgroup is accounting for memory usage inside a namespace (container). When all processes in that cgroup exit, the memory cgroup is released by Docker. However, “memory” isn’t just process memory, and although processes memory usage itself is gone, it turns out the kernel also assigns cached content like dentries and inodes (directory and file metadata) that are cached to the memory cgroup. From that issue.

该问题的概述是内存 cgroup 正在考虑命名空间（容器）内的内存使用情况。当该 cgroup 中的所有进程退出时，内存 cgroup 由 Docker 释放。然而，“内存”不仅仅是进程内存，虽然进程内存使用本身已经消失，但事实证明内核还会分配缓存内容，如缓存到内存 cgroup 的 dentries 和 inode（目录和文件元数据）。从那个问题。

“zombie” cgroups: cgroups that have no processes and have been deleted but still have memory charged to them (in my case, from the dentry cache, but it could also be from page cache or tmpfs). 

“zombie” cgroups：没有进程且已被删除但仍为它们分配内存的 cgroups（在我的情况下，来自 dentry 缓存，但也可能来自页面缓存或 tmpfs）。

Rather than the kernel iterating over every page in the cache at cgroup release time, which could be very slow, they choose to wait for those pages to be reclaimed and then finally clean up the cgroup once all are reclaimed when memory is needed, lazily. In the meantime, the cgroup still needs to be counted during stats collection.

不是内核在 cgroup 释放时迭代缓存中的每个页面，这可能会非常慢，而是选择等待这些页面被回收，然后在需要内存时懒惰地在所有页面被回收后最终清理 cgroup。同时，在统计数据收集过程中仍然需要计算 cgroup。

From a performance perspective, they are trading off time on a slow process by amortizing it over the reclamation of each page, opting to make the initial cleanup fast in return for leaving some cached memory around. That’s fine, when the kernel reclaims the last of the cached memory, the cgroup eventually gets cleaned up, so it’s not really a “leak”. Unfortunately the search that `memory.stat` performs, the way it's implemented on the kernel version (4.9) we're running on some servers, combined with the huge amount of memory on our servers, means it can take a significantly long time for the last of the cached data to be reclaimed and for the zombie cgroup to be cleaned up.

从性能的角度来看，他们通过在每个页面的回收上分摊时间来折衷一个缓慢的过程，选择快速进行初始清理以换取一些缓存内存。没关系，当内核回收最后的缓存内存时，cgroup 最终会被清理干净，所以这并不是真正的“泄漏”。不幸的是，`memory.stat` 执行的搜索，它在我们在某些服务器上运行的内核版本 (4.9) 上的实现方式，再加上我们服务器上的大量内存，意味着它可能需要很长时间要回收的最后一个缓存数据并清理僵尸 cgroup。

It turns out we had nodes that had such a large number of zombie cgroups that some had reads/stalls of over a second.

事实证明，我们的节点拥有如此多的僵尸 cgroup，以至于有些节点的读取/停顿时间超过一秒。

The workaround on that cadvisor issue, to immediately free the dentries/inodes cache systemwide, immediately stopped the read latency, and also the network latency stalls on the host, since the dropping of the cache included the cached pages in the “zombie” cgroups and so they were also freed. This isn’t a solution, but it does validate the cause of the issue.

该 cadvisor 问题的解决方法是立即释放系统范围内的 dentries/inode 缓存，立即停止读取延迟，并且主机上的网络延迟也停止，因为缓存的删除包括“僵尸”cgroup 中的缓存页面和所以他们也被释放了。这不是解决方案，但它确实验证了问题的原因。

As it turns out newer kernel releases (4.19+) have improved the performance of the `memory.stat` call and so this is no longer a problem after moving to that kernel. In the interim, we had existing tooling that was able to detect problems with nodes in our Kubernetes clusters and gracefully drain and reboot them, which we used to detect the cases of high enough latency that would cause issues, and treat them with a graceful reboot . This gave us breathing room while OS and kernel upgrades were rolled out to the remainder of the fleet.

事实证明，较新的内核版本 (4.19+) 提高了 `memory.stat` 调用的性能，因此在移动到该内核后这不再是问题。在此期间，我们现有的工具能够检测 Kubernetes 集群中节点的问题，并优雅地排空和重启它们，我们用来检测会导致问题的足够高的延迟情况，并通过优雅的重启来处理它们.这给了我们喘息的空间，同时操作系统和内核升级被推出到机队的其余部分。

## Wrapping up [Wrapping up](http://github.blog\#wrapping-up)

## 总结 [总结](http://github.blog\#wrapping-up)

Since this problem manifested as NIC RX queues not being processed for hundreds of milliseconds, it was responsible for both high latency on short connections and latency observed mid-connection such as between MySQL query and response packets. Understanding and maintaining performance of our most foundational systems like Kubernetes is critical to the reliability and speed of all services that build on top of them. As we invest in and improve on this performance, every system we run benefits from those improvements. 

由于此问题表现为数百毫秒内未处理 NIC RX 队列，因此它负责短连接的高延迟和连接中间（例如 MySQL 查询和响应数据包之间）观察到的延迟。了解和维护我们最基本的系统（如 Kubernetes）的性能对于建立在它们之上的所有服务的可靠性和速度至关重要。随着我们对这种性能进行投资和改进，我们运行的每个系统都会从这些改进中受益。

