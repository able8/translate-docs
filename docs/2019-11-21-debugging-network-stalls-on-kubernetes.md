# Debugging network stalls on Kubernetes

We’ve talked about [Kubernetes](https://github.blog/2017-08-16-kubernetes-at-github/) before, and over the last couple of years it’s become the standard deployment pattern at GitHub. We now run a large portion of both internal and public-facing services on Kubernetes. As our Kubernetes clusters have grown, and our targets on the latency of our services have become more stringent, we began to notice that certain services running on Kubernetes in our environment were experiencing sporadic latency that couldn’t be attributed to the performance characteristics of the application itself.

Essentially, applications running on our Kubernetes clusters would observe seemingly random latency of up to and over 100ms on connections, which would cause downstream timeouts or retries. Services were expected to be able to respond to requests in well under 100ms, which wasn’t feasible when the connection itself was taking so long. Separately, we also observed very fast MySQL queries, which we expected to take a matter of milliseconds and that MySQL observed taking only milliseconds, were being observed taking 100ms or more from the perspective of the querying application.

The problem was initially narrowed down to communications that involved a Kubernetes node, even if the other side of a connection was outside Kubernetes. The most simple reproduction we had was a [Vegeta](https://github.com/tsenart/vegeta) benchmark that could be run from any internal host, targeting a Kubernetes service running on a node port, and would observe the sporadically high latency. In this post, we’ll walk through how we tracked down the underlying issue.

## Removing complexity to find the path at fault

Using an example reproduction, we wanted to narrow down the problem and remove layers of complexity. Initially, there were too many moving parts in the flow between Vegeta and pods running on Kubernetes to determine if this was a deeper network problem, so we needed to rule some out.

![vegeta to kube pod via NodePort](https://i1.wp.com/user-images.githubusercontent.com/349190/66730905-2ede2800-eea0-11e9-9091-5e68bb23467d.png?resize=816%2C210&ssl=1)

The client, Vegeta, creates a TCP connection to any kube-node in the cluster. Kubernetes runs in our data centers as an [overlay network](https://en.wikipedia.org/wiki/Overlay_network) (a network that runs on top of our existing datacenter network) that uses [IPIP](https://en.wikipedia.org/wiki/IP_in_IP) (which encapsulates the overlay network’s IP packet inside the datacenter’s IP packet). When a connection is made to that first kube-node, it performs stateful [Network Address Translation](https://en.wikipedia.org/wiki/Network_address_translation) (NAT) to convert the kube-node’s IP and port to an IP and port on the overlay network (specifically, of the pod running the application). On return, it undoes each of these steps. This is a complex system with a lot of state, and a lot of moving parts that are constantly updating and changing as services deploy and move around.

As part of running a `tcpdump` on the original Vegeta benchmark, we observed the latency during a TCP handshake (between SYN and SYN-ACK). To simplify some of the complexity of HTTP and Vegeta, we can use `hping3` to just “ping” with a SYN packet and see if we observe the latency in the response packet—then throw away the connection. We can filter it to only include packets over 100ms and get a simpler reproduction case than a full Layer 7 Vegeta benchmark or attack against the service. The following “pings” a kube-node using TCP SYN/SYN-ACK on the “node port” for the service (30927) with an interval of 10ms, filtered for slow responses.

```shell
theojulienne@shell ~ $ sudo hping3 172.16.47.27 -S -p 30927 -i u10000 | egrep --line-buffered 'rtt=[0-9]{3}\.'
len=46 ip=172.16.47.27 ttl=59 DF id=0 sport=30927 flags=SA seq=1485 win=29200 rtt=127.1 ms
len=46 ip=172.16.47.27 ttl=59 DF id=0 sport=30927 flags=SA seq=1486 win=29200 rtt=117.0 ms
len=46 ip=172.16.47.27 ttl=59 DF id=0 sport=30927 flags=SA seq=1487 win=29200 rtt=106.2 ms
len=46 ip=172.16.47.27 ttl=59 DF id=0 sport=30927 flags=SA seq=1488 win=29200 rtt=104.1 ms
```

Our first new observation from the sequence numbers and timings is that this isn’t a one-off, but is often grouped, like a backlog that eventually gets processed.

Next up, we want to narrow down which component(s) were potentially at fault. Is it the kube-proxy iptables NAT rules that are hundreds of rules long? Is it the IPIP tunnel and something on the network handling them poorly? One way to validate this is to test each step of the system. What happens if we remove the NAT and firewall logic and only use the IPIP part.

![hping 3 on the IPIP tunnel alone](https://i0.wp.com/user-images.githubusercontent.com/349190/66730912-37cef980-eea0-11e9-8048-984d6c5966a3.png?resize=570%2C210&ssl=1)

Linux thankfully lets you just talk directly to an overlay IP when you’re on a machine that’s part of the same network, so that’s pretty easy to do.

```shell
theojulienne@kube-node-client ~ $ sudo hping3 10.125.20.64 -S -i u10000 | egrep --line-buffered 'rtt=[0-9]{3}\.'
len=40 ip=10.125.20.64 ttl=64 DF id=0 sport=0 flags=RA seq=7346 win=0 rtt=127.3 ms
len=40 ip=10.125.20.64 ttl=64 DF id=0 sport=0 flags=RA seq=7347 win=0 rtt=117.3 ms
len=40 ip=10.125.20.64 ttl=64 DF id=0 sport=0 flags=RA seq=7348 win=0 rtt=107.2 ms
```

Based on our results, the problem still remains! That rules out iptables and NAT. Is it TCP that’s the problem? Let’s see what happens when we perform a normal ICMP ping.

```shell
theojulienne@kube-node-client ~ $ sudo hping3 10.125.20.64 --icmp -i u10000 | egrep --line-buffered 'rtt=[0-9]{3}\.'
len=28 ip=10.125.20.64 ttl=64 id=42594 icmp_seq=104 rtt=110.0 ms
len=28 ip=10.125.20.64 ttl=64 id=49448 icmp_seq=4022 rtt=141.3 ms
len=28 ip=10.125.20.64 ttl=64 id=49449 icmp_seq=4023 rtt=131.3 ms
len=28 ip=10.125.20.64 ttl=64 id=49450 icmp_seq=4024 rtt=121.2 ms
```

Our results show that the problem still exists. Is it the IPIP tunnel that’s causing the problem? Let’s simplify things further.

![hping 3 directly between hosts](https://i1.wp.com/user-images.githubusercontent.com/349190/66730916-3c93ad80-eea0-11e9-810f-0f32da529da0.png?resize=570%2C210&ssl=1)

Is it possible that it’s every packet between these two hosts?

```shell
theojulienne@kube-node-client ~ $ sudo hping3 172.16.47.27 --icmp -i u10000 | egrep --line-buffered 'rtt=[0-9]{3}\.'
len=46 ip=172.16.47.27 ttl=61 id=41127 icmp_seq=12564 rtt=140.9 ms
len=46 ip=172.16.47.27 ttl=61 id=41128 icmp_seq=12565 rtt=130.9 ms
len=46 ip=172.16.47.27 ttl=61 id=41129 icmp_seq=12566 rtt=120.8 ms
len=46 ip=172.16.47.27 ttl=61 id=41130 icmp_seq=12567 rtt=110.8 ms
```

Behind the complexity, it’s as simple as two kube-node hosts sending any packet, even ICMP pings, to each other. They’ll still see the latency, if the target host is a “bad” one (some are worse than others).

Now there’s one last thing to question: we clearly don’t observe this everywhere, so why is it just on kube-node servers? And does it occur when the kube-node is the sender or the receiver? Luckily, this is also pretty easy to narrow down by using a host outside Kubernetes as a sender, but with the same “known bad” target host (from a staff shell host to the same kube-node). We can observe this is still an issue in that direction.

```shell
theojulienne@shell ~ $ sudo hping3 172.16.47.27 -p 9876 -S -i u10000 | egrep --line-buffered 'rtt=[0-9]{3}\.'
len=46 ip=172.16.47.27 ttl=61 DF id=0 sport=9876 flags=RA seq=312 win=0 rtt=108.5 ms
len=46 ip=172.16.47.27 ttl=61 DF id=0 sport=9876 flags=RA seq=5903 win=0 rtt=119.4 ms
len=46 ip=172.16.47.27 ttl=61 DF id=0 sport=9876 flags=RA seq=6227 win=0 rtt=139.9 ms
len=46 ip=172.16.47.27 ttl=61 DF id=0 sport=9876 flags=RA seq=7929 win=0 rtt=131.2 ms
```

And then perform the same from the previous source kube-node to a staff shell host (which rules out the source host, since a ping has both an RX and TX component).

```shell
theojulienne@kube-node-client ~ $ sudo hping3 172.16.33.44 -p 9876 -S -i u10000 | egrep --line-buffered 'rtt=[0-9]{3}\.'
^C
--- 172.16.33.44 hping statistic ---
22352 packets transmitted, 22350 packets received, 1% packet loss
round-trip min/avg/max = 0.2/7.6/1010.6 ms
```

Looking into the packet captures of the latency we observed, we get some more information. Specifically, that the “sender” host (bottom) observes this timeout while the “receiver” host (top) does not—see the Delta column (in seconds).

![packet captures of the latency](https://i0.wp.com/user-images.githubusercontent.com/349190/66730925-44ebe880-eea0-11e9-9efd-a533d2531de4.png?resize=3714%2C800&ssl=1)

Additionally, by looking at the difference between the ordering of the packets (based on the sequence numbers) on the receiver side of the TCP and ICMP results above, we can observe that ICMP packets always arrive in the same sequence they were sent, but with uneven timing, while TCP packets are sometimes interleaved, but a subset of them stall. Notably, we observe that if you count the ports of the SYN packets, the ports are not in order on the receiver side, while they’re in order on the sender side.

There is a subtle difference between how modern server [NICs](https://en.wikipedia.org/wiki/Network_interface_controller)—like we have in our data centers—handle packets containing TCP vs ICMP. When a packet arrives, the NIC hashes the packet “per connection” and tries to divvy up the connections across receive queues, each (approximately) delegated to a given CPU core. For TCP, this hash includes both source and destination IP and port. In other words, each connection is hashed (potentially) differently. For ICMP, just the IP source and destination are hashed, since there are no ports.

Another new observation is that we can tell that ICMP observes stalls on all communications between the two hosts during this period from the sequence numbers in ICMP vs TCP, while TCP does not. This tells us that the RX queue hashing is likely in play, almost certainly indicating the stall is in processing RX packets, not in sending responses.

This rules out kube-node transmits, so we now know that it’s a stall in processing packets, and that it’s on the receive side on some kube-node servers.

## Deep dive into Linux kernel packet processing

To understand why the problem could be on the receiving side on some kube-node servers, let’s take a look at how the Linux kernel processes packets.

Going back to the simplest traditional implementation, the network card receives a packet and sends an [interrupt](https://en.wikipedia.org/wiki/Interrupt) to the Linux kernel stating that there’s a packet that should be handled. The kernel stops other work, switches context to the interrupt handler, processes the packet, then switches back to what it was doing.

![traditional interrupt-driven approach](https://i1.wp.com/user-images.githubusercontent.com/349190/66730927-4ae1c980-eea0-11e9-897c-75cd39dbe286.png?resize=743%2C272&ssl=1)

This context switching is slow, which may have been fine on a 10Mbit NIC in the 90s, but on modern servers where the NIC is 10G and at maximal line rate can bring in around 15 million packets per second, on a smaller server with eight cores that could mean the kernel is interrupted millions of times per second per core.

Instead of constantly handling interrupts, many years ago Linux added [NAPI](https://en.wikipedia.org/wiki/New_API), the networking API that modern drivers use for improved performance at high packet rates. At low rates, the kernel still accepts interrupts from the NIC in the method we mentioned. Once enough packets arrive and cross a threshold, it disables interrupts and instead begins polling the NIC and pulling off packets in batches. This processing is done in a “softirq”, or [software interrupt context](https://www.kernel.org/doc/htmldocs/kernel-hacking/basics-softirqs.html). This happens at the end of syscalls and hardware interrupts, which are times that the kernel (as opposed to userspace) is already running.

![NAPI polling in softirq at end of syscalls](https://i0.wp.com/user-images.githubusercontent.com/349190/66730932-4f0de700-eea0-11e9-927f-393690107ace.png?resize=740%2C293&ssl=1)

This is much faster, but brings up another problem. What happens if we have so many packets to process that we spend all our time processing packets from the NIC, but we never have time to let the userspace processes actually drain those queues (read from TCP connections, etc.)? Eventually the queues would fill up, and we’d start dropping packets. To try and make this fair, the kernel limits the amount of packets processed in a given softirq context to a certain budget. Once this budget is exceeded, it wakes up a separate thread called `ksoftirqd` (you’ll see one of these in `ps` for each core) which processes these softirqs outside of the normal syscall/interrupt path. This thread is scheduled using the standard process scheduler, which already tries to be fair.

![NAPI polling crossing threshold in schedule ksoftirqd](https://i0.wp.com/user-images.githubusercontent.com/349190/66730935-533a0480-eea0-11e9-8697-7fc52d2ea1fa.png?resize=841%2C358&ssl=1)

With an overview of the way the kernel is processing packets, we can see there is definitely opportunity for this processing to become stalled. If the time between softirq processing calls grows, packets could sit in the NIC RX queue for a while before being processed. This could be something deadlocking the CPU core, or it could be something slow preventing the kernel from running softirqs.

## Narrow down processing to a core or method 

At this point, it makes sense that this could happen, and we know we’re observing something that looks a lot like it. The next step is to confirm this theory, and if we do, understand what’s causing it.

Let’s revisit the slow round trip packets we saw before.

```shell
len=46 ip=172.16.53.32 ttl=61 id=29573 icmp_seq=1953 rtt=99.3 ms
len=46 ip=172.16.53.32 ttl=61 id=29574 icmp_seq=1954 rtt=89.3 ms
len=46 ip=172.16.53.32 ttl=61 id=29575 icmp_seq=1955 rtt=79.2 ms
len=46 ip=172.16.53.32 ttl=61 id=29576 icmp_seq=1956 rtt=69.1 ms
len=46 ip=172.16.53.32 ttl=61 id=29577 icmp_seq=1957 rtt=59.1 ms
```

As discussed previously, these ICMP packets are hashed to a single NIC RX queue and processed by a single CPU core. If we want to understand what the kernel is doing, it’s helpful to know where (cpu core) and how (softirq, ksoftirqd) it’s processing these packets so we can catch them in action.

Now it’s time to use the tools that allow live tracing of a running Linux kernel— [bcc](https://github.com/iovisor/bcc) is what was used here. This allows you to write small C programs that hook arbitrary functions in the kernel, and buffer events back to a userspace Python program which can summarize and return them to you. The “hook arbitrary functions in the kernel” is the difficult part, but it actually goes out of its way to be as safe as possible to use, because it’s designed for tracing exactly this type of production issue that you can’t simply reproduce in a testing or dev environment.

The plan here is simple: we know the kernel is processing those ICMP ping packets, so let’s hook the kernel function [icmp\_echo](https://github.com/torvalds/linux/blob/v4.19/net/ipv4/icmp.c#L925) which takes an incoming ICMP “echo request” packet and initiates sending the ICMP “echo response” reply. We can identify the packet using the incrementing icmp\_seq shown by `hping3` above.

The code for this [bcc script](https://gist.github.com/theojulienne/9d78a0cb68dbe56f19a2ae6316bc6846) looks complex, but breaking it down it’s not as scary as it sounds. The `icmp_echo` function is passed a `struct sk_buff *skb`, which is the packet containing the ICMP echo request. We can delve into this live and pull out the `echo.sequence` (which maps to the `icmp_seq` shown by `hping3` above), and send that back to userspace. Conveniently, we can also grab the current process name/id as well. This gives us results like the following, live as the kernel processes these packets.

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

With that running, we can now correlate back from the stalled packets observed with `hping3` to the process that’s handling it. A simple `grep` on that capture for the `icmp_seq` values with some context shows what happened before these packets were processed. The packets that line up with the above `hping3` icmp\_seq values have been marked along with the rtt’s we observed above (and what we’d have expected if <50ms rtt’s weren’t filtered out).

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

The results tells us a few things. First, these packets are being processed by `ksoftirqd/11` which conveniently tells us this particular pair of machines have their ICMP packets hashed to core 11 on the receiving side. We can also see that every time we see a stall, we always see some packets processed in `cadvisor`’s syscall softirq context, followed by `ksoftirqd` taking over and processing the backlog, exactly the number we’d expect to work through the backlog.

The fact that `cadvisor` is always running just prior to this immediately also implicates it in the problem. Ironically, [cadvisor](https://github.com/google/cadvisor) “analyzes resource usage and performance characteristics of running containers”, yet it’s triggering this performance problem. As with many things related to containers, it’s all relatively bleeding-edge tooling which can result in some somewhat expected corner cases of bad performance.

## What is cadvisor doing to stall things? [What is cadvisor doing to stall things?](http://github.blog\#what-is-cadvisor-doing-to-stall-things)

With the understanding of how the stall can happen, the process causing it, and the CPU core it’s happening on, we now have a pretty good idea of what this looks like. For the kernel to hard block and not schedule `ksoftirqd` earlier, and given we see packets processed under `cadvisor`’s softirq context, it’s likely that `cadvisor` is running a slow syscall which ends with the rest of the packets being processed.

![slow syscall causing stalled packet processing on NIC RX queue](https://i2.wp.com/user-images.githubusercontent.com/349190/66730943-5a611280-eea0-11e9-902a-454f482d4223.png?resize=841%2C345&ssl=1)

That’s a theory but how do we validate this is actually happening? One thing we can do is trace what’s running on the CPU core throughout this process, catch the point where the packets are overflowing budget and processed by ksoftirqd, then look back a bit to see what was running on the CPU core. Think of it like taking an x-ray of the CPU every few milliseconds. It would look something like this.

![tracing of cpu to catch bad syscall and preceding work](https://i2.wp.com/user-images.githubusercontent.com/349190/66730945-5e8d3000-eea0-11e9-89af-daea3a7bd3f6.png?resize=748%2C224&ssl=1)

Conveniently, this is something that’s already mostly supported. The [`perf record`](https://perf.wiki.kernel.org/index.php/Tutorial#Sampling_with_perf_record) tool samples a given CPU core at a certain frequency and can generate a call graph of the live system, including both userspace and the kernel. Taking that recording and manipulating it using a quick fork of a tool from [Brendan Gregg’s FlameGraph](https://github.com/brendangregg/FlameGraph) that retained stack trace ordering, we can get a one-line stack trace for each 1ms sample, then get a sample of the 100ms before `ksoftirqd` is in the trace.

```shell
# record 999 times a second, or every 1ms with some offset so not to align exactly with timers
sudo perf record -C 11 -g -F 999
# take that recording and make a simpler stack trace.
sudo perf script 2>/dev/null | ./FlameGraph/stackcollapse-perf-ordered.pl | grep ksoftir -B 100
```

This results in the following.

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

Each line is a trace of the CPU at a point in time. Each call down the stack is separated by `;` on that line. Looking at the middle of the lines we can see the syscall being called is `read(): .... ;do_syscall_64;sys_read; ...` So cadvisor is spending a lot of time in a `read()` syscall relating to `mem_cgroup_*` functions (the top of the call stack / end of line).

The call stack trace isn’t convenient to see what’s being read, so let’s use `strace` to see what cadvisor is doing and find 100ms-or-slower syscalls.

```shell
theojulienne@kube-node-bad ~ $ sudo strace -p 10137 -T -ff 2>&1 | egrep '<0\.[1-9]'
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

```shell
theojulienne@kube-node-bad ~ $ time cat /sys/fs/cgroup/memory/memory.stat >/dev/null

real    0m0.153s
user    0m0.000s
sys    0m0.152s
theojulienne@kube-node-bad ~ $
```

Since we can reproduce it, this indicates that it’s the kernel hitting a pathologically bad case.

## What causes this read to be so slow [What causes this read to be so slow](http://github.blog\#what-causes-this-read-to-be-so-slow)

At this point it’s much more simple to find similar issues reported by others. As it turns out, this has been reported to cadvisor as an [excessive CPU usage problem](https://github.com/google/cadvisor/issues/1774), it just hadn’t been observed that latency was also being introduced to the network stack randomly as well. In fact, some folks internally had noticed cadvisor was consuming more CPU than expected, but it didn’t seem to be causing an issue since our servers had plenty of CPU capacity, and so the CPU usage hadn’t yet been investigated.

The overview of the issue is that the memory cgroup is accounting for memory usage inside a namespace (container). When all processes in that cgroup exit, the memory cgroup is released by Docker. However, “memory” isn’t just process memory, and although processes memory usage itself is gone, it turns out the kernel also assigns cached content like dentries and inodes (directory and file metadata) that are cached to the memory cgroup. From that issue.

“zombie” cgroups: cgroups that have no processes and have been deleted but still have memory charged to them (in my case, from the dentry cache, but it could also be from page cache or tmpfs).

Rather than the kernel iterating over every page in the cache at cgroup release time, which could be very slow, they choose to wait for those pages to be reclaimed and then finally clean up the cgroup once all are reclaimed when memory is needed, lazily. In the meantime, the cgroup still needs to be counted during stats collection.

From a performance perspective, they are trading off time on a slow process by amortizing it over the reclamation of each page, opting to make the initial cleanup fast in return for leaving some cached memory around. That’s fine, when the kernel reclaims the last of the cached memory, the cgroup eventually gets cleaned up, so it’s not really a “leak”. Unfortunately the search that `memory.stat` performs, the way it’s implemented on the kernel version (4.9) we’re running on some servers, combined with the huge amount of memory on our servers, means it can take a significantly long time for the last of the cached data to be reclaimed and for the zombie cgroup to be cleaned up.

It turns out we had nodes that had such a large number of zombie cgroups that some had reads/stalls of over a second.

The workaround on that cadvisor issue, to immediately free the dentries/inodes cache systemwide, immediately stopped the read latency, and also the network latency stalls on the host, since the dropping of the cache included the cached pages in the “zombie” cgroups and so they were also freed. This isn’t a solution, but it does validate the cause of the issue.

As it turns out newer kernel releases (4.19+) have improved the performance of the `memory.stat` call and so this is no longer a problem after moving to that kernel. In the interim, we had existing tooling that was able to detect problems with nodes in our Kubernetes clusters and gracefully drain and reboot them, which we used to detect the cases of high enough latency that would cause issues, and treat them with a graceful reboot. This gave us breathing room while OS and kernel upgrades were rolled out to the remainder of the fleet.

## Wrapping up [Wrapping up](http://github.blog\#wrapping-up)

Since this problem manifested as NIC RX queues not being processed for hundreds of milliseconds, it was responsible for both high latency on short connections and latency observed mid-connection such as between MySQL query and response packets. Understanding and maintaining performance of our most foundational systems like Kubernetes is critical to the reliability and speed of all services that build on top of them. As we invest in and improve on this performance, every system we run benefits from those improvements.
