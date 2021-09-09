# How eBPF Turns Linux into a Programmable Kernel

# eBPF 如何将 Linux 变成可编程内核

#### 8 Oct 2020 6:00am,   by [Joab Jackson](https://thenewstack.io/author/joab/ "Posts by Joab Jackson")

#### 2020 年 10 月 8 日上午 6:00，作者：[Joab Jackson](https://thenewstack.io/author/joab/“Joab Jackson 的帖子”)

![](https://cdn.thenewstack.io/media/2020/10/1d303ea9-sculpture-3170012_640.jpg)

The Linux kernel could see a radical shift in how it operates, given the full promise of the [Extended Berkeley Packet Filter](https://ebpf.io/)(eBPF), argued [Daniel Borkmann](http://borkmann.ch/), Linux kernel engineer for [Cilium](https://cilium.io/), in a [technical session](https://www.youtube.com/watch?v=99jUcLt3rSk) during the recent [ KubeCon + CloudNativeCon EU](https://thenewstack.io/kubecon-eu-cloud-native-developers-now-an-army-6-5-million-strong/) virtual conference.

[Daniel Borkmann](http://borkmann.ch/)，[Cilium](https://cilium.io/) 的 Linux 内核工程师，在 [技术会议](https://www.youtube.com/watch?v=99jUcLt3rSk) 期间 [ KubeCon + CloudNativeCon EU](https://thenewstack.io/kubecon-eu-cloud-native-developers-now-an-army-6-5-million-strong/) 虚拟会议。

Although originally targeted [for superior in-kernel monitoring](https://thenewstack.io/linux-technology-for-the-new-year-ebpf/), this memory-mapped extension of the original BPF can run any sandboxed programs within the kernel space, without changing kernel source code or loading modules. This represents a radically new, and potentially faster and safer way to use the Linux kernel. In effect, eBPF provides a way for developers to add their own programs into the kernel itself.

尽管最初针对 [用于高级内核监控](https://thenewstack.io/linux-technology-for-the-new-year-ebpf/)，但原始 BPF 的这种内存映射扩展可以运行任何沙盒程序在内核空间内，无需更改内核源代码或加载模块。这代表了一种全新的、可能更快、更安全的使用 Linux 内核的方式。实际上，eBPF 为开发人员提供了一种将他们自己的程序添加到内核本身的方法。

Borkmann even predicts that the Linux kernel may morph into an “eBPF-powered microkernel,” a tiny core kernel with minimal built-in capabilities. All other functions, including end-user custom ones, would be eBPF functions.

Borkmann 甚至预测 Linux 内核可能会演变成一个“eBPF 驱动的微内核”，一个具有最少内置功能的微小核心内核。所有其他函数，包括最终用户自定义的函数，都将是 eBPF 函数。

One day, for instance, Kubernetes would ship with a set of custom kernel extensions, based on the underlying workload, be it in the data center or in a small edge system. Using eBPF today, Cilium offers for Kubernetes a Container Networking Interface (CNI) that can enforce network traffic rules, offer load balancing, and arguably by Cilium, serve as a speedier alternative KubeProxy. The [Hubble](https://cilium.io/blog/2019/11/19/announcing-hubble/) project [builds on this work](https://www.youtube.com/watch?v=8WCbGSCyDSo&t=1382s) by offering Kubernetes observability, again using eBPF.

例如，有一天，Kubernetes 会根据底层工作负载（无论是在数据中心还是在小型边缘系统中）附带一组自定义内核扩展。今天使用 eBPF，Cilium 为 Kubernetes 提供了一个容器网络接口 (CNI)，它可以强制执行网络流量规则，提供负载平衡，并且可以说是由 Cilium 用作更快的替代 KubeProxy。 [Hubble](https://cilium.io/blog/2019/11/19/annoucing-hubble/) 项目 [建立在这项工作的基础上](https://www.youtube.com/watch?v=8WCbGSCyDSo&t=1382s)通过提供 Kubernetes 可观察性，再次使用 eBPF。

![](https://cdn.thenewstack.io/media/2020/10/2d244dfe-ebpf-1024x579.png)

When the upgrade was first offered up to the kernel keepers in 2013, it was rejected for being a big patch bomb, Borkmann recalled. With this in mind, the eBPF designers stepped back and started incrementally, and significantly, upgrading the kernel’s existing BPF. They heavily extended the instruction set, swapped out the interpreter for a new one altogether, and added in a verifier to ensure the code is correct for the kernel. As a result, eBPF is now full a virtual machine within the kernel itself.

Borkmann 回忆说，当 2013 年首次向内核管理员提供升级时，它被拒绝了，因为它是一个大补丁炸弹。考虑到这一点，eBPF 设计者退后一步，开始逐步升级内核现有的 BPF。他们大量扩展了指令集，将解释器完全替换为一个新的解释器，并添加了一个验证器以确保代码对内核来说是正确的。因此，eBPF 现在完全是内核内部的虚拟机。

Today, every packet that goes through Facebook is handled by eBPF. Cloudflare also deploys it to great extent, Borkmann said. Netflix has [started using eBPF](https://netflixtechblog.com/extending-vector-with-ebpf-to-inspect-host-and-container-performance-5da3af4c584b) to run its programs securely in the kernel, with Netflix Kernel Engineer [Brendan Gregg](http://www.brendangregg.com/) calling it “ [a fundamental change to a 50-year old kernel model](http://www.brendangregg.com/blog/2019-12-02/bpf-a-new-type-of-software.html).” And building upon eBPF, Microsoft Azure Chief Technology Officer [Mark Russinovich](https://www.linkedin.com/in/markrussinovich/) is [writing a version](https://twitter.com/markrussinovich/status/1283039153920368651) of his vaunted Sysmon monitoring program for Linux. 

今天，通过 Facebook 的每个数据包都由 eBPF 处理。博克曼说，Cloudflare 也在很大程度上部署了它。 Netflix 已经[开始使用 eBPF](https://netflixtechblog.com/extending-vector-with-ebpf-to-inspect-host-and-container-performance-5da3af4c584b) 在内核中安全地运行其程序，使用 Netflix Kernel工程师 [Brendan Gregg](http://www.brendangregg.com/) 称其为“[对 50 年历史内核模型的根本改变](http://www.brendangregg.com/blog/2019-12-02/bpf-a-new-type-of-software.html)。”基于 eBPF，Microsoft Azure 首席技术官 [Mark Russinovich](https://www.linkedin.com/in/markrussinovich/) 正在 [编写版本](https://twitter.com/markrussinovich/status/1283039153920368651) 他吹嘘的 Linux Sysmon 监控程序。

With [eBPF](https://man7.org/linux/man-pages/man2/bpf.2.html), developers write the code in a subset of C, which is compiled into BPF bytecode to run on the BPF virtual machine. After safety-checking the code, a just-in-time compiler converts, the bytecode into architecture-specific machine code, explained software engineer [Liz Rice](https://www.lizrice.com/), In [her eBPF talk at Dockercon 2019](https://www.youtube.com/watch?v=4SiWL5tULnQ&t=520s). Another project, the BPF Compiler Collection (BCC), is developing a language wrapper for Python, to make it easier to code into BPF.

使用[eBPF](https://man7.org/linux/man-pages/man2/bpf.2.html)，开发者用C的一个子集编写代码，编译成BPF字节码在BPF虚拟机上运行机器。软件工程师 [Liz Rice](https://www.lizrice.com/) 在 [她的 eBPF 演讲中解释说，在对代码进行安全检查后，即时编译器将字节码转换为特定于架构的机器代码在 Dockercon 2019](https://www.youtube.com/watch?v=4SiWL5tULnQ&t=520s)。另一个项目，BPF 编译器集合 (BCC)，正在为 Python 开发一个语言包装器，以便更容易地编码到 BPF 中。

Successfully compiled machine code is then attached to a kernel’s code path, which, when traversed, executes any attached eBPF programs. State can be accessible to all users through a shared memory map.

成功编译的机器代码然后附加到内核的代码路径，当遍历时，执行任何附加的 eBPF 程序。所有用户都可以通过共享内存映射访问状态。

Kicking off an eBPF function requires an event of some sort, the arrival of a network packet, a data point from a tracepoint, or a signal from an application. An advantage eBPF has over a standard user-space function is that it can make logical decisions based on the contents of a packet or other bit of data flowing through the kernel, Rice pointed out. Each function is limited to 4,096 instructions, though larger functions can be created by chaining smaller ones together.

启动 eBPF 功能需要某种事件、网络数据包的到达、来自跟踪点的数据点或来自应用程序的信号。赖斯指出，eBPF 与标准用户空间函数相比的一个优势是，它可以根据数据包的内容或流经内核的其他数据位做出逻辑决策。每个函数限制为 4,096 条指令，但可以通过将较小的函数链接在一起来创建更大的函数。

[Open eBPF and Kubernetes: Little Helper Minions for Scaling Microservices – Daniel Borkmann, Cilium on YouTube.](https://www.youtube.com/watch?v=99jUcLt3rSk)

[开放 eBPF 和 Kubernetes：用于扩展微服务的小助手 Minions – Daniel Borkmann，YouTube 上的 Cilium。](https://www.youtube.com/watch?v=99jUcLt3rSk)

[Open eBPF Superpowers on YouTube.](https://www.youtube.com/watch?v=4SiWL5tULnQ&t=520s)

[在 YouTube 上打开 eBPF Superpowers。](https://www.youtube.com/watch?v=4SiWL5tULnQ&t=520s)

[Open Hubble – eBPF Based Observability for Kubernetes – Sebastian Wicki, Isovalent on YouTube.](https://www.youtube.com/watch?v=8WCbGSCyDSo&t=1382s)

[Open Hubble – 基于 eBPF 的 Kubernetes 可观察性 – Sebastian Wicki，YouTube 上的 Isovalent。](https://www.youtube.com/watch?v=8WCbGSCyDSo&t=1382s)

The Cloud Native Computing Foundation is a sponsor of The New Stack.

云原生计算基金会是 The New Stack 的赞助商。

Feature image par [Couleur](https://pixabay.com/fr/users/couleur-1195798/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3170012) de [Pixabay.](https://pixabay.com/fr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3170012) 

特征图像标准 [Couleur](https://pixabay.com/fr/users/couleur-1195798/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3170012) de [Pixabay.](https/pixabay.comfr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3170012)

