# How eBPF Turns Linux into a Programmable Kernel

#### 8 Oct 2020 6:00am,   by [Joab Jackson](https://thenewstack.io/author/joab/ "Posts by Joab Jackson")

![](https://cdn.thenewstack.io/media/2020/10/1d303ea9-sculpture-3170012_640.jpg)

The Linux kernel could see a radical shift in how it operates, given the full promise of the [Extended Berkeley Packet Filter](https://ebpf.io/) (eBPF), argued [Daniel Borkmann](http://borkmann.ch/), Linux kernel engineer for [Cilium](https://cilium.io/), in a [technical session](https://www.youtube.com/watch?v=99jUcLt3rSk) during the recent [KubeCon + CloudNativeCon EU](https://thenewstack.io/kubecon-eu-cloud-native-developers-now-an-army-6-5-million-strong/) virtual conference.

Although originally targeted [for superior in-kernel monitoring](https://thenewstack.io/linux-technology-for-the-new-year-ebpf/), this memory-mapped extension of the original BPF can run any sandboxed programs within the kernel space, without changing kernel source code or loading modules. This represents a radically new, and potentially faster and safer way to use the Linux kernel. In effect, eBPF provides a way for developers to add their own programs into the kernel itself.

Borkmann even predicts that the Linux kernel may morph into an “eBPF-powered microkernel,” a tiny core kernel with minimal built-in capabilities. All other functions, including end-user custom ones, would be eBPF functions.

One day, for instance, Kubernetes would ship with a set of custom kernel extensions, based on the underlying workload, be it in the data center or in a small edge system. Using eBPF today, Cilium offers for Kubernetes a Container Networking Interface (CNI) that can enforce network traffic rules, offer load balancing, and arguably by Cilium, serve as a speedier alternative KubeProxy. The [Hubble](https://cilium.io/blog/2019/11/19/announcing-hubble/) project [builds on this work](https://www.youtube.com/watch?v=8WCbGSCyDSo&t=1382s) by offering Kubernetes observability, again using eBPF.

![](https://cdn.thenewstack.io/media/2020/10/2d244dfe-ebpf-1024x579.png)

When the upgrade was first offered up to the kernel keepers in 2013, it was rejected for being a big patch bomb, Borkmann recalled. With this in mind, the eBPF designers stepped back and started incrementally, and significantly, upgrading the kernel’s existing BPF. They heavily extended the instruction set, swapped out the interpreter for a new one altogether, and added in a verifier to ensure the code is correct for the kernel. As a result, eBPF is now full a virtual machine within the kernel itself.

Today, every packet that goes through Facebook is handled by eBPF. Cloudflare also deploys it to great extent, Borkmann said. Netflix has [started using eBPF](https://netflixtechblog.com/extending-vector-with-ebpf-to-inspect-host-and-container-performance-5da3af4c584b) to run its programs securely in the kernel, with Netflix Kernel Engineer [Brendan Gregg](http://www.brendangregg.com/) calling it “ [a fundamental change to a 50-year old kernel model](http://www.brendangregg.com/blog/2019-12-02/bpf-a-new-type-of-software.html).” And building upon eBPF, Microsoft Azure Chief Technology Officer [Mark Russinovich](https://www.linkedin.com/in/markrussinovich/) is [writing a version](https://twitter.com/markrussinovich/status/1283039153920368651) of his vaunted Sysmon monitoring program for Linux.

With [eBPF](https://man7.org/linux/man-pages/man2/bpf.2.html), developers write the code in a subset of C, which is compiled into BPF bytecode to run on the BPF virtual machine. After safety-checking the code, a just-in-time compiler converts, the bytecode into architecture-specific machine code, explained software engineer [Liz Rice](https://www.lizrice.com/), In [her eBPF talk at Dockercon 2019](https://www.youtube.com/watch?v=4SiWL5tULnQ&t=520s). Another project, the BPF Compiler Collection (BCC), is developing a language wrapper for Python, to make it easier to code into BPF.

Successfully compiled machine code is then attached to a kernel’s code path, which, when traversed, executes any attached eBPF programs. State can be accessible to all users through a shared memory map.

Kicking off an eBPF function requires an event of some sort, the arrival of a network packet, a data point from a tracepoint, or a signal from an application. An advantage eBPF has over a standard user-space function is that it can make logical decisions based on the contents of a packet or other bit of data flowing through the kernel, Rice pointed out. Each function is limited to 4,096 instructions, though larger functions can be created by chaining smaller ones together.

[Open eBPF and Kubernetes: Little Helper Minions for Scaling Microservices – Daniel Borkmann, Cilium on YouTube.](https://www.youtube.com/watch?v=99jUcLt3rSk)

[Open eBPF Superpowers on YouTube.](https://www.youtube.com/watch?v=4SiWL5tULnQ&t=520s)

[Open Hubble – eBPF Based Observability for Kubernetes – Sebastian Wicki, Isovalent on YouTube.](https://www.youtube.com/watch?v=8WCbGSCyDSo&t=1382s)

The Cloud Native Computing Foundation is a sponsor of The New Stack.

Feature image par [Couleur](https://pixabay.com/fr/users/couleur-1195798/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3170012) de [Pixabay.](https://pixabay.com/fr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=3170012)
