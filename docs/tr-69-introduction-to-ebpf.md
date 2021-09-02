# [Introduction to eBPF](https://oswalt.dev/2021/01/introduction-to-ebpf/)

# [eBPF 简介](https://oswalt.dev/2021/01/introduction-to-ebpf/)

January 19, 2021 16-minute read

If you’re paying much attention at all to the systems or cloud-native communities, you have **certainly** heard about eBPF. It has dominated several conference schedules for the last few years, and it even has [its own conference](https://ebpf.io/summit-2020/) now! However, unlike a lot of hyped-up technology buzzwords, this one’s momentum doesn’t seem to be unwarranted, or even ahead of the curve:

如果您非常关注系统或云原生社区，那么您**肯定**听说过 eBPF。它在过去几年中主导了几个会议日程，现在甚至有了[自己的会议]（https://ebpf.io/summit-2020/）！然而，与许多炒作的技术流行语不同，这个势头似乎并非毫无根据，甚至是领先于曲线：

> “BPF is not someone’s academic paper. BPF is not a “proof of concept”. BPF is running in production”
> – [Brendan Gregg](https://www.youtube.com/watch?v=ZYBXZFKPS28)

> “BPF 不是某人的学术论文。 BPF 不是“概念证明”。 BPF 正在生产中运行”
> – [布伦丹格雷格](https://www.youtube.com/watch?v=ZYBXZFKPS28)

I’d characterize the community around eBPF to still be fairly small, and highly pragmatic. After researching this topic, I think it’s an important technology to be aware of, but how it applies to you depends wildly on your skills and your role as a technologist.

我认为 eBPF 周围的社区仍然相当小，而且非常务实。在研究了这个主题之后，我认为这是一项需要注意的重要技术，但它如何适用于您在很大程度上取决于您的技能和您作为技术专家的角色。

eBPF itself as a topic is **huge**, and there is a multitude of resources out there that dive into the specifics, many of which I’ll link to below. Please check out those external resources - I couldn’t possibly improve on them. I’m writing this post primarily as a way to kick off my eBPF learning journey, but also to pursue answers to the following questions:

eBPF 本身作为一个主题是**巨大的**，并且有大量资源可以深入了解具体细节，我将在下面链接到其中的许多资源。请查看这些外部资源 - 我不可能改进它们。我写这篇文章主要是为了开启我的 eBPF 学习之旅，同时也是为了寻求以下问题的答案：

- What problem was eBPF created to solve, and why is it a better model than some of the existing alternative approaches?

- 创建 eBPF 是为了解决什么问题，为什么它比一些现有的替代方法更好？

- How does eBPF work, and how can we use it to build applications that fit within this new architecture?

- eBPF 是如何工作的，我们如何使用它来构建适合这种新架构的应用程序？

- Who should care about eBPF, and for what reasons?


- 谁应该关心 eBPF，出于什么原因？


## Why We Use Operating Systems

## 我们为什么使用操作系统

Operating systems give us a lot to build on top of. The tradeoff is that you have to play by their rules. If for instance, you want to send some packets over the network, you don’t access the NIC directly; rather you have to work through the API that sits on top of the underlying kernel, using a mechanism typically referred to as “syscalls”. The applications you use every day, even the very browser you’re using to read this post right now is leveraging a similar API. This creates a clear demarcation between the space where our apps run (userspace) and the space where the operating system internals and hardware interactions take place (kernel space). In fact, operating systems often partition system memory so that the two are kept separate from each other.

操作系统为我们提供了很多构建的基础。权衡是你必须遵守他们的规则。例如，如果你想通过网络发送一些数据包，你不直接访问 NIC；相反，您必须使用位于底层内核之上的 API，使用通常称为“系统调用”的机制。你每天使用的应用程序，甚至你现在用来阅读这篇文章的浏览器都在利用类似的 API。这在我们的应用程序运行的空间（用户空间）和操作系统内部和硬件交互发生的空间（内核空间）之间创建了清晰的界限。事实上，操作系统经常对系统内存进行分区，以便将两者彼此分开。

This paradigm is extremely helpful - it’s not safe or feasible to expect every application in the universe to be able to work with every type of hardware in existence, and also coexist with each other without crashing the system. However, this comes with some tradeoffs. The two big ones are as follows:

这种范式非常有用 - 期望宇宙中的每个应用程序都能够与现有的每种类型的硬件一起工作，并且彼此共存而不会使系统崩溃，这是不安全或不可行的。但是，这需要进行一些权衡。两个大的如下：

1. When your application requests something from the kernel, often a chunk of data in kernel space needs to be copied into userspace. We need to do this because operating systems strictly partition regions of memory used for the kernel, so it’s not possible to simply give a userspace program a pointer to some region of kernel memory. This is commonly referred to as "[crossing the user/kernel boundary](https://developer.apple.com/library/archive/documentation/Darwin/Conceptual/KernelProgramming/boundaries/boundaries.html)", and because of this copy operation, operations like these can have pretty significant performance implications.

1. 当你的应用程序向内核请求某些东西时，通常需要将内核空间中的一大块数据复制到用户空间中。我们需要这样做是因为操作系统严格划分了用于内核的内存区域，所以不可能简单地给用户空间程序一个指向内核内存某个区域的指针。这通常被称为“[跨越用户/内核边界](https://developer.apple.com/library/archive/documentation/Darwin/Conceptual/KernelProgramming/boundaries/boundaries.html)”，因此复制操作，像这样的操作可能会对性能产生相当大的影响。

1. Like all abstractions, the programmatic interface provided by an operating system via syscalls results in a certain loss of context. You cannot freely access resources in kernel space, you can only work with what you’re given. Similarly, you cannot just walk into a restaurant’s kitchen and start making food - you have to order from the menu, and this keeps a large chunk of the experience out of your hands.

1. 与所有抽象一样，操作系统通过系统调用提供的编程接口会导致一定程度的上下文丢失。你不能自由地访问内核空间中的资源，你只能使用你得到的东西。同样，您不能只是走进餐厅的厨房并开始制作食物 - 您必须从菜单中订购，这会让您无法获得大量体验。


> For the sake of this article, we’ll be focusing only on Linux, but many of the same traits discussed here apply to other operating systems. 
> 就本文而言，我们将只关注 Linux，但此处讨论的许多相同特征也适用于其他操作系统。

Linux is optimized for being an end-device, and working with large packets, not so much for packet forwarding/decisions and lots of small packets. This is a big reason for Linux’s reputation as a very poorly performing substitute for the big ASIC-driven machines that most IT shops and service providers are accustomed to. One reason for this is the penalty of crossing the userspace boundary. For traffic that cannot be handled purely by the kernel, the entire packet must be copied into the memory space available to a userspace process, which means each packet incurs a performance cost, and the more packets you have to receive, make decisions on, and forward, the more this cost becomes apparent.

Linux 针对终端设备进行了优化，可以处理大数据包，而不是数据包转发/决策和大量小数据包。这是 Linux 作为大多数 IT 商店和服务提供商习惯的大型 ASIC 驱动机器的性能非常差的替代品而享有盛誉的一个重要原因。原因之一是跨越用户空间边界的惩罚。对于不能完全由内核处理的流量，必须将整个数据包复制到用户空间进程可用的内存空间中，这意味着每个数据包都会产生性能成本，并且您必须接收、做出决策和处理的数据包越多越往前，这种成本就越明显。

[![](https://oswalt.dev/assets/2021/01/memcopy.png)](https://oswalt.dev/assets/2021/01/memcopy.png)

The attempt to use Linux as a full-blown networking platform, therefore, has resulted in three distinct approaches, each of which has some significant tradeoffs:

因此，尝试将 Linux 作为成熟的网络平台使用，导致了三种不同的方法，每种方法都有一些重要的权衡：

1. **Traditional Approach** \- an application that sits in userspace, leverages syscalls, etc to work with the kernel. This is the “traditional” approach because you're using each part of the system the way it was intended, but as we've discussed, for some use cases like network forwarding, there can be severe performance and loss-of-context penalties . A lot of higher-level network service applications like load balancers and intrusion prevention systems work this way.

1. **传统方法** \- 位于用户空间中的应用程序，利用系统调用等与内核一起工作。这是“传统”方法，因为您正在按照预期的方式使用系统的每个部分，但正如我们所讨论的，对于网络转发等某些用例，可能会出现严重的性能和上下文丢失损失.许多更高级别的网络服务应用程序，如负载平衡器和入侵防御系统，都以这种方式工作。

2. **Do Everything in [Userspace](https://arxiv.org/pdf/1901.10664.pdf)** \- rely on the kernel for nothing, and reinvent all the wheels (aka “have fun writing your own TCP stack”). This is the approach taken by [Intel’s DPDK](https://www.linuxjournal.com/content/userspace-networking-dpdk).

2. **Do Everything in [Userspace](https://arxiv.org/pdf/1901.10664.pdf)** \- 无所事事地依赖内核，重新发明所有轮子（也就是“编写自己的 TCP 玩得开心”堆”）。这是[英特尔的 DPDK](https://www.linuxjournal.com/content/userspace-networking-dpdk) 采用的方法。

3. **Kernel Modules** \- add the functionality you need to the kernel by writing a module. There’s no userspace boundary crossing here, and you have access to everything you need, all the way down to the hardware. This is the approach taken by [Juniper Contrail](https://github.com/tungstenfabric/tf-vrouter) as an example. Unfortunately, when working purely in kernel space as you are with kernel modules, there's [no hard boundary](https://www.tedinski.com/2019/01/15/system-boundaries-and-the-linux-kernel. html) or API guarantees, like there is with the userspace boundary. Also, the blast radius for when things go wrong is system-wide; you can cause security issues, or crash the entire system. Finally, each new Linux kernel version brings the very real possibility of breaking changes, so that a kernel module must be refactored to be compatible. So this approach is extremely fragile.


3. **内核模块** \- 通过编写模块将您需要的功能添加到内核中。这里没有用户空间边界跨越，你可以访问你需要的一切，一直到硬件。这是以 [Juniper Contrail](https://github.com/tungstenfabric/tf-vrouter) 为例的方法。不幸的是，当您像使用内核模块一样纯粹在内核空间中工作时，[没有硬边界](https://www.tedinski.com/2019/01/15/system-boundaries-and-the-linux-kernel。 html) 或 API 保证，就像用户空间边界一样。此外，出现问题时的爆炸半径是系统范围的；您可能会导致安全问题，或使整个系统崩溃。最后，每个新的 Linux 内核版本都带来了非常真实的破坏性更改的可能性，因此必须重构内核模块才能兼容。所以这种方法极其脆弱。


eBPF gives us a fourth option. But before we get there, we need to discuss a little history.

eBPF 为我们提供了第四种选择。但在我们到达那里之前，我们需要讨论一点历史。

## Berkeley Packet Filter

## 伯克利数据包过滤器

The origin story of eBPF lies in it’s first, more narrowly-focused implementation known at the time as Berkeley Packet Filter. This was a new architecture that allowed systems to filter packets **well** before they were further received by the kernel - or worse - copied into userspace. In addition, this made use of a new register-based virtual machine, with a very limited set of instructions, that could be provided by a userspace process. This meant that utilities like `tcpdump` could compile very simple filtering programs (influenced by more human-readable filter strings provided by the user via command-line) that ran entirely in kernel space, without the need for kernel modules.

eBPF 的起源故事在于它的第一个更狭窄的实现，当时称为伯克利数据包过滤器。这是一种新的架构，允许系统在数据包被内核进一步接收之前过滤**好**，或者更糟的是 - 复制到用户空间。此外，这利用了一个新的基于寄存器的虚拟机，其指令集非常有限，可以由用户空间进程提供。这意味着像 `tcpdump` 这样的实用程序可以编译完全在内核空间中运行的非常简单的过滤程序（受用户通过命令行提供的更具可读性的过滤器字符串的影响），而无需内核模块。

The tradeoff here was that you can “program” the kernel to filter certain packets without having to consult a userspace process for these decisions - all without using kernel modules - but the tools you have to make decisions are fairly limited to the instructions available to you in the BPF virtual machine. This limitation is done to ensure the safety and stability of the kernel. You get the performance of running entirely in kernel space, but you have to play by the rules of the BPF compiler.

这里的权衡是，您可以“编程”内核以过滤某些数据包，而无需咨询用户空间进程来做出这些决定——所有这些都不需要使用内核模块——但是您必须做出决定的工具仅限于您可用的指令在 BPF 虚拟机中。进行此限制是为了确保内核的安全性和稳定性。您可以获得完全在内核空间中运行的性能，但您必须遵守 BPF 编译器的规则。

A fun exercise is to run `tcpdump` with the `-d` flag to see the resulting BPF program that is compiled on your behalf, and sent to the kernel for execution.

一个有趣的练习是运行带有 `-d` 标志的 `tcpdump`，以查看代表你编译的结果 BPF 程序，并发送到内核执行。

```fallback
~$ tcpdump -i wlp2s0 -d 'tcp'

 (000) ldh      [12]
 (001) jeq      #0x86dd          jt 2	jf 7
 (002) ldb      [20]
 (003) jeq      #0x6             jt 10	jf 4
 (004) jeq      #0x2c            jt 5	jf 11
 (005) ldb      [54]
 (006) jeq      #0x6             jt 10	jf 11
 (007) jeq      #0x800           jt 8	jf 11
 (008) ldb      [23]
 (009) jeq      #0x6             jt 10	jf 11
 (010) ret      #262144
 (011) ret      #0

```


I highly recommend reading the [original BPF paper](https://www.tcpdump.org/papers/bpf-usenix93.pdf), as it discusses these ideas quite well.

我强烈建议阅读 [原始 BPF 论文](https://www.tcpdump.org/papers/bpf-usenix93.pdf)，因为它很好地讨论了这些想法。

## Enter: “eBPF” 
## 输入：“eBPF”
The “ [elevator pitch](https://ebpf.io/what-is-ebpf/#why-ebpf)” you'll hear about eBPF (“extended BPF”) is that it does to the kernel what Javascript did for HTML. The Linux kernel is mostly presented “as is” - in that you have a surface area you can interact with via syscalls, and that’s about it. This may be fine for some applications, but for others, this is a pretty significant constraint that warranted some of the workarounds we discussed in the sections above. eBPF expands on the original architecture behind classic BPF by providing a set of tools for writing fairly generic programs that are run purely in kernel space, safely, and with a suite of other tools to complete the picture.

您将听到的关于 eBPF（“扩展 BPF”）的“[elevator pitch](https://ebpf.io/what-is-ebpf/#why-ebpf)”是它对内核的作用与 Javascript 所做的一样HTML。 Linux 内核大多按“原样”呈现——因为您有一个可以通过系统调用与之交互的表面区域，仅此而已。这对于某些应用程序可能没问题，但对于其他应用程序，这是一个非常重要的约束，需要我们在上面的部分中讨论的一些变通方法。 eBPF 扩展了经典 BPF 背后的原始架构，提供了一组工具来编写完全在内核空间中安全运行的相当通用的程序，并使用一套其他工具来完成这幅图。

eBPF is the result of many improvements and enhancements to BPF, stretching back to 2014, turning it into a much more capable and robust method of programming the kernel. There are still constraints in place to ensure the safety of the kernel, but it is far more robust than the classic BPF implementation. eBPF enables a much wider variety of use cases ranging from high-performance networking to system observability, storage troubleshooting, and more. It retains all of the advantages of the original idea behind BPF; it makes the kernel more programmable without the inherent performance penalties, instability, and insecurity of other approaches. However, it does so with a greatly expanded set of tools, resources, and instructions, which makes it applicable to a much wider set of use cases.

eBPF 是对 BPF 进行许多改进和增强的结果，可以追溯到 2014 年，将其变成了一种更强大、更强大的内核编程方法。仍然存在一些约束来确保内核的安全，但它比经典的 BPF 实现要健壮得多。 eBPF 支持更广泛的用例，从高性能网络到系统可观察性、存储故障排除等。它保留了 BPF 背后原始思想的所有优点；它使内核更具可编程性，而没有其他方法固有的性能损失、不稳定性和不安全性。但是，它使用大量扩展的工具、资源和说明来实现，这使其适用于更广泛的用例集。

You’d be forgiven for being confused about the difference between eBPF and BPF when referring to the current state of the art. Indeed, eBPF’s most prominent advocates like Brendan Gregg have been trying hard to clarify this:

当提到当前的技术状态时，你会因为 eBPF 和 BPF 之间的区别而感到困惑，这是情有可原的。事实上，像 Brendan Gregg 这样的 eBPF 最杰出的倡导者一直在努力澄清这一点：

> “It’s eBPF but really just BPF, which stands for Berkeley Packet Filter, although today it has little to do with Berkeley, packets, or filtering. Thus, BPF should be regarded now as a technology name rather than as an acronym.”
>
> – “Footnote from ‘BPF Performance Tools’ on page 18 by Brendan Gregg

> “它是 eBPF 但实际上只是 BPF，它代表伯克利数据包过滤器，尽管今天它与伯克利、数据包或过滤无关。因此，BPF 现在应该被视为一个技术名称而不是首字母缩写词。”
>
> – “Brendan Gregg 第 18 页上‘BPF 性能工具’的脚注

I’ve often seen eBPF and BPF used interchangeably, and in fact, when classic/legacy BPF is meant (which is rare), I’ve seen it described as “cBPF”. So I would recommend you use the same convention. These days, BPF and eBPF are the same; they both refer to BPF in the modern sense, with all of the aforementioned enhancements.

我经常看到 eBPF 和 BPF 互换使用，事实上，当指的是经典/传统 BPF 时（这种情况很少见），我看到它被描述为“cBPF”。所以我建议你使用相同的约定。现在，BPF 和 eBPF 是一样的；它们都指现代意义上的 BPF，具有上述所有增强功能。

[![Credit: https://ebpf.io](https://oswalt.dev/assets/2021/01/loader-ebpf.png)](https://oswalt.dev/assets/2021/01/loader-ebpf.png "Credit: https://ebpf.io")



When I think about eBPF, I like to break it down to the following three topic areas. Together, these provide a really valuable set of tools for anyone looking to add functionality on top of an existing Linux kernel.

当我考虑 eBPF 时，我喜欢将其分解为以下三个主题领域。总之，这些为希望在现有 Linux 内核之上添加功能的任何人提供了一组非常有价值的工具。

1. **The eBPF Toolchain** \- these are a series of software utilities (many of which are open source) as well as existing in-kernel components that allow you to write, compile, and deploy eBPF programs that are guaranteed to be as safe and secure as possible.

1. **eBPF 工具链** \- 这些是一系列软件实用程序（其中许多是开源的）以及现有的内核组件，它们允许您编写、编译和部署 eBPF 程序，这些程序保证尽可能安全可靠。

2. **Probes** \- specify event-driven rules for triggering a given eBPF program. Just like Javascript programs can react to things like a user clicking a button, eBPF can react to kernel events, and fire custom programs to respond to them _quickly_, and without involving a userspace process at all.

2. **Probes** \- 指定用于触发给定 eBPF 程序的事件驱动规则。就像 Javascript 程序可以对用户点击按钮之类的事情做出反应一样，eBPF 可以对内核事件做出反应，并触发自定义程序来_快速_响应它们，并且根本不涉及用户空间进程。

3. **Maps** \- share data between kernel and user space safely. This provides a place for eBPF programs to store data and get on with the rest of their operations, while userspace programs can access this data in parallel. You can think of it as a sort of in-kernel key/value store that is accessible from both sides of the boundary.


3. **Maps** \- 在内核和用户空间之间安全地共享数据。这为 eBPF 程序提供了一个存储数据并继续执行其余操作的地方，而用户空间程序可以并行访问这些数据。您可以将其视为一种可从边界两侧访问的内核内键/值存储。


There’s certainly more to talk about with eBPF, and each of these subjects probably warrants multiple blog posts on their own. However, I think this gives the best 10,000' view of what eBPF allows you to do. Write safe, performant kernel-space programs, run those programs only when needed, and get the relevant data out into userspace so you can do something meaningful with it.

eBPF 当然还有更多要讨论的内容，并且这些主题中的每一个都可能需要单独发表多篇博文。然而，我认为这提供了 eBPF 允许你做的最好的 10,000' 视图。编写安全、高性能的内核空间程序，仅在需要时运行这些程序，并将相关数据导出到用户空间中，以便您可以用它做一些有意义的事情。

## Diving a Bit Deeper

## 深入一点

With the high-level out of the way, I’d like to dive a little deeper into exactly what’s going on with BPF programs themselves. I’m planning to address some of the other topics like the toolchain, and the various other resources available **to** a BPF program in other posts. These are crucial things to know to understand the full value of eBPF, but for now, let’s just focus on what exactly BPF programs **are**. 
有了高层，我想更深入地了解 BPF 程序本身到底发生了什么。我计划在其他帖子中讨论一些其他主题，例如工具链，以及可用于 **BPF 程序的各种其他资源。这些是了解 eBPF 全部价值的关键，但现在，让我们只关注 BPF 程序 ** 是什么**。
The number one link to bookmark when learning eBPF is [ebpf.io](https://ebpf.io/what-is-ebpf/). However, if you are interested in more of the “in the weeds” aspects of eBPF, you'll quickly discover the link to the [eBPF and XDP Reference](https://docs.cilium.io/en/stable/bpf /) within the documentation for Cilium.

学习 eBPF 时书签的第一个链接是 [ebpf.io](https://ebpf.io/what-is-ebpf/)。但是，如果您对 eBPF 的更多“杂草丛生”方面感兴趣，您会很快发现 [eBPF 和 XDP 参考](https://docs.cilium.io/en/stable/bpf) 的链接/) 在 Cilium 的文档中。

> [Cilium](https://cilium.io/) is an open-source networking and security platform for Kubernetes that is powered by eBPF, and it just so happens I'm using this as the primary CNI in the cluster that powers [NRE Labs](https://nrelabs.io). It’s one of the most prominent eBPF-powered platforms out there, so a lot of their documentation is also good reading for learning eBPF, and related technologies like XDP which I’ll cover in a future post.

> [Cilium](https://cilium.io/) 是 Kubernetes 的开源网络和安全平台，由 eBPF 提供支持，碰巧我使用它作为支持集群的主要 CNI [NRE 实验室](https://nrelabs.io)。它是最突出的 eBPF 驱动平台之一，因此他们的许多文档也是学习 eBPF 和 XDP 等相关技术的好读物，我将在以后的文章中介绍。

At its core, BPF can be thought of as a simplified general-purpose [instruction set](https://docs.cilium.io/en/stable/bpf/#instruction-set). It is similar to something like the x86 instruction set in that it is a series of instructions and registers that can be used to execute programs. However, it is **highly simplified**. How simplified, you ask?

在其核心，BPF 可以被认为是一个简化的通用[指令集](https://docs.cilium.io/en/stable/bpf/#instruction-set)。它类似于 x86 指令集，因为它是一系列可用于执行程序的指令和寄存器。但是，它是**高度简化的**。你问有多简单？

[![](https://oswalt.dev/assets/2021/01/bpf_x86_comparison.png)](https://www.youtube.com/watch?v=Qhm1Zn_BNi4)

So, BPF is sort of a [“general-purpose execution engine”](https://www.youtube.com/watch?v=Qhm1Zn_BNi4), but not quite a fully generic instruction set (when compared to something like x86\ _64). It is [highly constrained and simplified](https://events.static.linuxfound.org/sites/events/files/slides/bpf_collabsummit_2015feb20.pdf) to maintain the promise of safe and stable code within the kernel (this is in addition to a series of other verifications that take place at a higher level).

因此，BPF 是一种 [“通用执行引擎”](https://www.youtube.com/watch?v=Qhm1Zn_BNi4)，但不是完全通用的指令集（与 x86 之类的指令集相比\ _64）。 [高度约束和简化](https://events.static.linuxfound.org/sites/events/files/slides/bpf_collabsummit_2015feb20.pdf) 维护内核中安全和稳定代码的承诺（这是另外到更高级别进行的一系列其他验证）。

However, there are two other pieces to be aware of that makes this part of the BPF story really powerful. First, you are not required to write BPF programs from scratch by slinging opcodes yourself. There exist plenty of tools that can take higher-level languages, and compile them down into BPF bytecode. [`bcc`](https://github.com/iovisor/bcc) allows you to write BPF programs in C, and even front-end those programs using high-level languages like Python. There are other libraries in Go and Rust (which we’ll explore in a future post) that give you a lot of flexibility in working with BPF. So, in terms of writing BPF programs, you have a lot of options.

然而，还有另外两个部分需要注意，这使得 BPF 故事的这一部分非常强大。首先，您不需要通过自己吊索操作码从头开始编写 BPF 程序。有很多工具可以采用高级语言，并将它们编译成 BPF 字节码。 [`bcc`](https://github.com/iovisor/bcc) 允许您用 C 编写 BPF 程序，甚至可以使用 Python 等高级语言对这些程序进行前端处理。 Go 和 Rust 中还有其他库（我们将在以后的文章中探讨）为您在使用 BPF 时提供了很大的灵活性。因此，在编写 BPF 程序方面，您有很多选择。

The other thing to consider is that once a BPF program is compiled into BPF bytecode, it needs to be executed. Of course, the processor in your computer doesn’t know how to speak BPF, so at some point, this needs to be translated to native machine instructions. This is the purpose of the “ [JIT compiler](https://docs.cilium.io/en/stable/bpf/#jit)”. Because of this, the BPF program can perform just as well as any other code running in the kernel (including modules), but with the added benefit of having gone through a pipeline of tools that ensure it is safe.

另一件要考虑的事情是，一旦 BPF 程序被编译成 BPF 字节码，就需要执行它。当然，您计算机中的处理器不知道如何使用 BPF，因此在某些时候，需要将其转换为本地机器指令。这就是“[JIT 编译器](https://docs.cilium.io/en/stable/bpf/#jit)”的目的。正因为如此，BPF 程序可以像内核（包括模块）中运行的任何其他代码一样执行，但具有通过确保其安全的工具管道的额外好处。

So, while there are a lot more components that play a role in getting a BPF program from inception to execution, it generally flows through these three states: source code, then BPF bytecode, and finally, native machine instructions.

因此，虽然有更多的组件在让 BPF 程序从开始到执行中发挥作用，但它通常流经这三个状态：源代码，然后是 BPF 字节码，最后是本机机器指令。

[![](https://oswalt.dev/assets/2021/01/bpf-program-lifecycle.png)](https://oswalt.dev/assets/2021/01/bpf-program-lifecycle.png)



## Who Cares?

## 谁在乎？

If you’re not a kernel developer (statistically speaking, most of you), you may be asking why you should care about eBPF. This is a valid concern, but I’d like to raise a few important points.

如果您不是内核开发人员（从统计上讲，你们中的大多数人），您可能会问为什么要关心 eBPF。这是一个合理的担忧，但我想提出一些重要的观点。

I think a takeaway for the average infrastructure operator is that eBPF is one of those technologies that has significant trickle-down effects for your operations. You should care about the role eBPF plays in your infrastructure software in the same way you should care about where your food comes from. You don’t have to be a farmer to have a vested interest in knowing more about that supply chain and using that knowledge to drive decision-making. 

我认为对于普通基础设施操作员来说，eBPF 是对您的运营具有显着涓滴效应的技术之一。你应该关心 eBPF 在你的基础设施软件中扮演的角色，就像关心食物的来源一样。您不必是农民，就可以在了解更多有关该供应链的信息并利用这些知识来推动决策方面拥有既得利益。

Think about [how bugs in Linux are fixed today](https://ebpf.io/what-is-ebpf#ebpfs-impact-on-the-linux-kernel), and the weeks, months, or sometimes **years ** it takes for those fixes to make their way into your shop. Bugs like these have to get fixed in the kernel, which then has to get released, and then we have to wait for our operating system vendor to release a version of their own that uses this upgraded kernel (which for some vendors can take a long time). Once that’s all done, we **then** have to actually upgrade the thing, which has all kinds of operational considerations (2 AM Saturday maintenance windows anyone?).

想想 [今天 Linux 中的错误是如何修复的](https://ebpf.io/what-is-ebpf#ebpfs-impact-on-the-linux-kernel)，以及数周、数月或有时是**年** 这些修复程序需要进入您的商店。像这样的错误必须在内核中修复，然后必须发布，然后我们必须等待我们的操作系统供应商发布他们自己的使用此升级内核的版本（对于某些供应商来说可能需要很长时间）时间）。一旦这一切都完成了，我们**那么**必须真正升级这个东西，它有各种操作考虑（周六凌晨 2 点维护窗口有人吗？）。

This kind of operational pain is felt equally by sysadmins who are deeply familiar with the distribution running on their server cluster but also to network engineers or storage admins who aren’t aware that their switch or cluster is running Linux under the covers. With BPF, if there’s a bug, this can get fixed in the BPF program itself, and in most cases, upgraded live without bringing anything down, or waiting for a new kernel or operating system release. This is because BPF allows your vendors to insert their own kernel functionality without having to include it in the mainstream kernel, or maintain fragile kernel modules.

非常熟悉在其服务器集群上运行的发行版的系统管理员以及不知道他们的交换机或集群在幕后运行 Linux 的网络工程师或存储管理员同样会感受到这种操作痛苦。使用 BPF，如果有错误，可以在 BPF 程序本身中修复，并且在大多数情况下，实时升级而不会导致任何故障，或等待新内核或操作系统发布。这是因为 BPF 允许您的供应商插入他们自己的内核功能，而不必将其包含在主流内核中，或维护脆弱的内核模块。

eBPF also enables a significant ability to simplify the software that’s running on our systems. Ivan Pepelnjak and Thomas Graf discussed eBPF on [an episode of “Software Gone Wild”](https://blog.ipspace.net/2016/10/fast-linux-packet-forwarding-with.html), and Thomas raised an interesting point:

eBPF 还能够显着简化在我们系统上运行的软件。 Ivan Pepelnjak 和 Thomas Graf 在 [“Software Gone Wild”的一集](https://blog.ipspace.net/2016/10/fast-linux-packet-forwarding-with.html) 上讨论了 eBPF，Thomas 提出了一个有趣的一点：

> “The nice thing about using eBPF for container networking is that since you’re creating the forwarding logic as a program on the fly, you can remove a whole lot of complexity because you’re only adding the code that you need. If a container doesn’t need v4, don’t include that logic.”

> “将 eBPF 用于容器网络的好处在于，由于您将转发逻辑创建为动态程序，因此您可以消除很多复杂性，因为您只需添加所需的代码。如果容器不需要 v4，则不要包含该逻辑。”

That said, I do think eBPF is more directly applicable to a software-savvy infrastructure professional when it comes to troubleshooting and observability. Clearly, some platforms like Cilium have been able to leverage eBPF to build full-blown networking solutions, but you don’t have to have that kind of big idea to get some value out of eBPF. You can use eBPF to add a little bit of instrumentation to your system to expose some metrics that are meaningful to you. See [tcplife](http://www.brendangregg.com/blog/2016-11-30/linux-bcc-tcplife.html) as an example. [bpftrace](https://github.com/iovisor/bpftrace) allows you to write very simple BPF programs for tracing and observability **very** quickly. Having the ability to write some quick tooling for yourself is powerful, and the eBPF tool-chain makes the kernel more accessible for those that just want a quick solution to a problem they’re having.

也就是说，在故障排除和可观察性方面，我确实认为 eBPF 更直接适用于精通软件的基础设施专业人员。很明显，像 Cilium 这样的一些平台已经能够利用 eBPF 来构建成熟的网络解决方案，但你不必有那种大创意就可以从 eBPF 中获得一些价值。您可以使用 eBPF 向您的系统添加一点检测，以公开一些对您有意义的指标。以[tcplife](http://www.brendangregg.com/blog/2016-11-30/linux-bcc-tcplife.html)为例。 [bpftrace](https://github.com/iovisor/bpftrace) 允许您编写非常简单的 BPF 程序来**非常**快速地进行跟踪和可观察性。拥有为自己编写一些快速工具的能力是很强大的，而且 eBPF 工具链使内核对于那些只想快速解决他们遇到的问题的人来说更容易访问。

So, while the vast majority of those reading this post may not need to dive into the deepest weeds of eBPF, I believe it's important for everyone to be aware of it, understand the problems it solves, and consider this new architecture when making technology decisions going forward. I like to explain it like this: even network operators who have never designed an ASIC before are still expected to know at least a little bit about how they work; it only aids in their understanding of the networks they manage. Minimally, I view eBPF the same way. So, whether you’re rack-mounting switches or allocating memory on the heap, eBPF will be a part of your supply chain, and you owe it to yourself to be informed. This is really about making the kernel programmable - which is such a fundamental change, I think we’ll be experiencing the repercussions on this for years to come.

因此，虽然绝大多数阅读这篇文章的人可能不需要深入研究 eBPF 的最深处，但我相信每个人都必须意识到它，了解它解决的问题，并在做出技术决策时考虑这种新架构往前走。我喜欢这样解释：即使是以前从未设计过 ASIC 的网络运营商，仍然希望至少了解一点它们的工作原理；它只会帮助他们了解他们管理的网络。至少，我以同样的方式看待 eBPF。因此，无论您是机架安装交换机还是在堆上分配内存，eBPF 都将成为您供应链的一部分，您应该了解情况。这实际上是关于使内核可编程 - 这是一个如此根本性的变化，我认为我们将在未来几年经历这方面的影响。

## Conclusion

## 结论

Honestly, I’m excited to dig into eBPF. The architecture is very pleasing to me, not only because of the speed and safety aspects but because [event-driven architectures](https://oswalt.dev/2016/12/introduction-to-stackstorm/) have been a part of my upbringing for a while and it's the way my brain works. eBPF also happens to be spiritually aligned with what I've come to appreciate from other technologies like the [Rust compiler](https://oswalt.dev/2020/03/getting-rusty/), which goes a long way (sometimes painfully) to ensure your software is safe while still keeping performance as high as possible.

老实说，我很高兴能够深入研究 eBPF。该架构让我非常满意，不仅因为速度和安全方面，还因为 [事件驱动架构](https://oswalt.dev/2016/12/introduction-to-stackstorm/) 已经成为我的成长经历了一段时间，这就是我大脑的工作方式。 eBPF 也恰好与我从 [Rust 编译器](https://oswalt.dev/2020/03/getting-rusty/) 等其他技术中所欣赏到的东西在精神上保持一致，这有很长的路要走（有时痛苦）以确保您的软件安全，同时仍保持尽可能高的性能。

I’ll be following up with some posts focused on some specifics, but hopefully, this suffices as a high-level introduction to eBPF and the problems it aims to solve.

我将跟进一些专注于某些细节的帖子，但希望这足以作为对 eBPF 及其旨在解决的问题的高级介绍。

## Additional Resources

## 其他资源

Some other super useful resources I didn’t link to above:

我上面没有链接到的其他一些超级有用的资源：

- [eBPF Summit Videos](https://ebpf.io/summit-2020/) (I learned a LOT watching these videos)
- [eBPF - a new type of software - Brendan Gregg](https://www.youtube.com/watch?v=7pmXdG8-7WU)
- [BPF at Facebook - Alexei Starovoitov](https://www.youtube.com/watch?v=ZYBXZFKPS28) 
- [eBPF 峰会视频](https://ebpf.io/summit-2020/)（观看这些视频我学到了很多）
- [eBPF - 一种新型软件 - Brendan Gregg](https://www.youtube.com/watch?v=7pmXdG8-7WU)
- [Facebook 上的 BPF - Alexei Starovoitov](https://www.youtube.com/watch?v=ZYBXZFKPS28)
- [eBPF Superpowers - Liz Rice](https://www.youtube.com/watch?v=4SiWL5tULnQ)
- [Github collection of awesome eBPF resources](https://github.com/zoidbergwill/awesome-ebpf)
- [LWN - a thorough introduction to eBPF](https://lwn.net/Articles/740157/)
- [A Brief Introduction to XDP and eBPF](http://blogs.igalia.com/dpino/2019/01/07/a-brief-introduction-to-xdp-and-ebpf/)
- [BPF, eBPF, XDP and Bpfilter… What are These Things and What do They Mean for the Enterprise?](https://www.netronome.com/blog/bpf-ebpf-xdp-and-bpfilter-what-are -these-things-and-what-do-they-mean-enterprise/) 
- [eBPF Superpowers - Liz Rice](https://www.youtube.com/watch?v=4SiWL5tULnQ)
- [Github 收集了很棒的 eBPF 资源](https://github.com/zoidbergwill/awesome-ebpf)
- [LWN - eBPF 的全面介绍](https://lwn.net/Articles/740157/)
- [XDP和eBPF简介](http://blogs.igalia.com/dpino/2019/01/07/a-brief-introduction-to-xdp-and-ebpf/)
- [BPF、eBPF、XDP 和 Bpfilter……这些东西是什么，它们对企业意味着什么？](https://www.netronome.com/blog/bpf-ebpf-xdp-and-bpfilter-what-are -这些事情和他们的意思是企业/)


