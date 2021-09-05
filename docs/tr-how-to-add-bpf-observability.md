## How To Add eBPF Observability To Your Product

## 如何为您的产品添加 eBPF 可观察性

03 Jul 2021

There's an arms race to add [eBPF](http://www.brendangregg.com/BPF) to commercial observability products, and in this post I'll describe how to quickly do that. This is also applicable for people adding it to their own in-house monitoring systems.

将 [eBPF](http://www.brendangregg.com/BPF) 添加到商业可观察性产品是一场军备竞赛，在这篇文章中，我将描述如何快速做到这一点。这也适用于将其添加到自己的内部监控系统中的人。

People like to show me their BPF observability products after they have prototyped or built them, but I often wish I had given them advice before they started. As the leader of BPF observability, it's advice I've been including in recent talks, and now I'm including it in this post.

人们喜欢在他们完成原型设计或构建之后向我展示他们的 BPF 可观察性产品，但我经常希望我在他们开始之前就给了他们建议。作为 BPF 可观察性的领导者，我在最近的谈话中一直包含它的建议，现在我将它包含在这篇文章中。

First, I know you're busy. You might not even like BPF. To be pragmatic, I'll describe how to spend the least effort to get the most value. Think of this as "version 1": A starting point that's pretty useful. Whether you follow this advice or not, at least please understand it to avoid later regrets and pain.

首先，我知道你很忙。你甚至可能不喜欢 BPF。为了务实，我将描述如何花费最少的努力获得最大的价值。将此视为“版本 1”：一个非常有用的起点。不管你是否遵循这个建议，至少请理解它，以免以后后悔和痛苦。

If you're using an open source monitoring platform, first check if it already has a BPF agent. This post assumes it doesn't, and you'll be adding something for the first time.

如果你使用的是开源监控平台，首先检查它是否已经有 BPF 代理。这篇文章假设它没有，你将第一次添加一些东西。

## 1\. Run your first tool

## 1\.运行你的第一个工具

Start by installing the [bcc](https://github.com/iovisor/bcc) or [bpftrace](https://github.com/iovisor/bpftrace) tools. E.g., bcc on Ubuntu:

首先安装 [bcc](https://github.com/iovisor/bcc) 或 [bpftrace](https://github.com/iovisor/bpftrace) 工具。

```
# apt-get install bpfcc-tools
```


Then try running a tool. E.g., to see process execution with timestamps using execsnoop(8):

然后尝试运行一个工具。例如，使用 execsnoop(8) 查看带有时间戳的进程执行：

```
# execsnoop-bpfcc -T
TIME     PCOMM            PID    PPID   RET ARGS
19:36:15 service          828567 6009     0 /usr/sbin/service --status-all
19:36:15 basename         828568 828567   0
19:36:15 basename         828569 828567   0 /usr/bin/basename /usr/sbin/service
19:36:15 env              828570 828567   0 /usr/bin/env -i LANG=en_AU.UTF-8 LANGUAGE=en_AU:en LC_CTYPE= LC_NUMERIC= LC_TIME= LC_COLLATE= LC_MONETARY= LC_MESSAGES= LC_PAPER= LC_NAME= LC_ADDRESS= LC_TELEPHONE= LC_MEASUREMENT= LC_IDENTIFICATION= LC_ALL= PATH=/opt/local/bin:/opt/local/sbin:/usr/local/git/bin:/home/bgregg/.local/bin:/home/bgregg/bin:/opt/local/bin:/opt/local/sbin:/ TERM=xterm-256color /etc/init.d/acpid
19:36:15 acpid            828570 828567   0 /etc/init.d/acpid status
19:36:15 run-parts        828571 828570   0 /usr/bin/run-parts --lsbsysinit --list /lib/lsb/init-functions.d
19:36:15 systemctl        828572 828570   0 /usr/bin/systemctl -p LoadState --value show acpid.service
19:36:15 readlink         828573 828570   0 /usr/bin/readlink -f /etc/init.d/acpid
[...]

```


While basic, I've solved many perf issues with this tool alone, including for misconfigured systems where a shell script is launching failing processes in a loop, and when some minor application is crashing and is restarting every few minutes but has not yet been noticed .

虽然是基本的，但我单独使用这个工具解决了许多性能问题，包括配置错误的系统，其中 shell 脚本在循环中启动失败的进程，以及当一些小应用程序崩溃并每隔几分钟重新启动但尚未注意到时.

## 2\. Add a tool to your product

## 2\.为您的产品添加工具

Now imagine adding execsnoop(8) to your product. You likely already have agents running on all your customer systems. Do they have a way to run a command and return the text output? Or run a command and send the output elsewhere for aggregation (S3, Hive, Druid, etc.)? There are so many options it's really your own preference based on your existing system and customer environments.

现在想象一下将 execsnoop(8) 添加到您的产品中。您可能已经在所有客户系统上运行了代理。他们有办法运行命令并返回文本输出吗？或者运行命令并将输出发送到其他地方进行聚合（S3、Hive、Druid 等）？有很多选项，这真的是您根据现有系统和客户环境的偏好。

When you add your first tool to your product, have it run it for a short duration such as 10 to 60 seconds.
I just noticed execsnoop(8) doesn't have a duration option yet, so in the interim you could wrap it with watch -s2 60 execsnoop-bpfcc.
If you want to run these tools 24x7, study overheads to understand the cost first. Low frequency events such as process execution should be negligible to capture.

当您将第一个工具添加到您的产品时，让它运行一小段时间，例如 10 到 60 秒。
我刚刚注意到 execsnoop(8) 还没有持续时间选项，所以在此期间你可以用 watch -s2 60 execsnoop-bpfcc 包装它。
如果您想 24x7 全天候运行这些工具，请先研究开销以了解成本。诸如流程执行之类的低频事件应该可以忽略不计。

Instead of bcc, you can also use the [bpftrace](https://github.com/iovisor/bpftrace) versions. These typically don't have canned options (-v, -l, etc.), but do have a json output mode. E.g.:

除了 bcc，您还可以使用 [bpftrace](https://github.com/iovisor/bpftrace) 版本。这些通常没有固定选项（-v、-l 等)，但有 json 输出模式。例如。：

```
# bpftrace -f json execsnoop.bt 
{"type": "attached_probes", "data": {"probes": 2}}
{"type": "printf", "data": "TIME(ms)   PID   ARGS\n"}
{"type": "printf", "data": "2737       849176 "}
{"type": "join", "data": "ls -F"}
{"type": "printf", "data": "5641       849178 "}
{"type": "join", "data": "date"}

```


This mode was added so that BPF observability products can be built on top of bpftrace.

添加此模式是为了可以在 bpftrace 之上构建 BPF 可观察性产品。

## 3\. Don't worry about dependencies 

## 3\. 不用担心依赖

I am indeed suggesting that you install bcc or bpftrace on your customer systems, and they currently have llvm dependencies. This can add up to tens of Mbytes, which can be a problem for some resource-constrained environments (embedded). We've been doing lots of work to fix this in the future. bcc has newer versions of the tools (libbpf-tools) that use [BTF and CO-RE](http://www.brendangregg.com/blog/2020-11-04/bpf-co-re-btf-libbpf.html) (and not Python) and will ultimately mean you can install 100-Kbyte binary versions of the tools with no dependencies.
bpftrace has a similar plan to produce a small dependency-less binary using the newer kernel features.

我确实建议您在客户系统上安装 bcc 或 bpftrace，它们目前具有 llvm 依赖项。这可以加起来高达数十 MB，这对于某些资源受限的环境（嵌入式）来说可能是一个问题。我们一直在做很多工作来解决这个问题。 bcc 有使用 [BTF 和 CO-RE] 的较新版本的工具 libbpf-tools（而不是 Python），最终意味着您可以安装 100 KB 的二进制版本的工具，而无需依赖项。
bpftrace 有一个类似的计划，即使用较新的内核功能生成一个小的无依赖二进制文件。

This does require at least Linux 5.8 to work well, and your customers may not run that for years. In the interim I'd suggest not worrying about the llvm dependencies for now since it will be fixed later.

这确实需要至少 Linux 5.8 才能正常运行，而您的客户可能多年都不会运行它。在此期间，我建议暂时不要担心 llvm 依赖项，因为稍后会修复它。

Note that not all Linux distributions have enabled CONFIG\_DEBUG\_INFO\_BTF=y, which is necessary for the future of BTF and CO-RE. Major distros have set it, such as in Ubuntu 20.10, Fedora 30, and RHEL 8.2. But if you know some of your customers are running something uncommon, please check and encourage them or the distro vendor to set CONFIG\_DEBUG\_INFO\_BTF=y and CONFIG\_DEBUG\_INFO\_BTF\_MODULES=y to avoid pain in the future.

请注意，并非所有 Linux 发行版都启用了 CONFIG\_DEBUG\_INFO\_BTF=y，这是 BTF 和 CO-RE 的未来所必需的。主要发行版都设置了它，例如在 Ubuntu 20.10、Fedora 30 和 RHEL 8.2 中。但是如果你知道你的一些客户正在运行一些不常见的东西，请检查并鼓励他们或发行版供应商设置 CONFIG\_DEBUG\_INFO\_BTF=y 和 CONFIG\_DEBUG\_INFO\_BTF\_MODULES=y 以避免痛苦未来。

## 4\. Version 1 dashboard

## 4. 版本 1 仪表板

Now you have one BPF observability tool in your product, it's time to add more. Here are the top ten tools you can run and present as a generic BPF observability dashboard, along with suggested visualizations:

现在您的产品中有一个 BPF 可观察性工具，是时候添加更多了。以下是您可以作为通用 BPF 可观察性仪表板运行和呈现的十大工具，以及建议的可视化：



This is based on my [bcc Tutorial](https://github.com/iovisor/bcc/blob/master/docs/tutorial.md), and many also exist in bpftrace. I chose these to find the most performance wins with the fewest tools.

这是基于我的 [bcc Tutorial](https://github.com/iovisor/bcc/blob/master/docs/tutorial.md)，很多也存在于 bpftrace 中。我选择这些是为了用最少的工具获得最大的性能优势。

Note that runqlat and profile can have noticable overheads, so I'd run these tools for between 10 and 60 seconds only and generate a report. Some are low enough overhead to be run 24x7 if desired (e.g., execsnoop, biolatency, tcplife, tcpretrans).

请注意，runqlat 和 profile 可能有明显的开销，因此我只会运行这些工具 10 到 60 秒并生成报告。如果需要，有些开销足够低，可以 24x7 运行（例如，execsnoop、biolatency、tcplife、tcpretrans）。

There is already documentation as man pages and example files in the bcc and bpftrace repositories that you can link to, to help your customers understand the tool output. Eg, here's the execsnoop(8) example files in [bcc](https://github.com/iovisor/bcc/blob/master/tools/execsnoop_example.txt) and [bpftrace](https://github.com/iovisor/bpftrace/blob/master/tools/execsnoop_example.txt).

您可以链接到 bcc 和 bpftrace 存储库中的手册页和示例文件等文档，以帮助您的客户了解工具输出。例如，这里是 [bcc](https://github.com/iovisor/bcc/blob/master/tools/execsnoop_example.txt) 和 [bpftrace](https://github.com/iovisor/bpftrace/blob/master/tools/execsnoop_example.txt)。

Once you have this all working, you have version 1!

一旦您完成所有这些工作，您就拥有了版本 1！

## Case study: Netflix

## 案例研究：Netflix

Netflix is building a new GUI that does this tool dashboard and more, based on the bpftrace versions of these tools. The architecture is:

Netflix 正在构建一个新的 GUI，它基于这些工具的 bpftrace 版本来执行此工具仪表板等。架构是：

![](http://www.brendangregg.com/blog/images/2021/flamecommander-bpf.png)

While the bpftrace binary is installed on all the target systems, the bpftrace tools (text files) live on a web server and are pushed out when needed. This means we can ensure we're always running the latest version of the tools by updating them in one place.

虽然 bpftrace 二进制文件安装在所有目标系统上，但 bpftrace 工具（文本文件）位于 Web 服务器上，并在需要时推出。这意味着我们可以通过在一个地方更新工具来确保我们始终运行最新版本的工具。

This is currently part of our FlameCommander UI, which also runs flame graphs across the cloud. Our previous BPF GUI was part of [Vector](https://github.com/Netflix/vector), and used bcc, but we've since deprecated that. We'll likely open source the new one at some point and have a post about it on the Netflix tech blog.

这是目前我们 FlameCommander UI 的一部分，它也在云中运行火焰图。我们之前的 BPF GUI 是 [Vector](https://github.com/Netflix/vector) 的一部分，并使用了 bcc，但我们已经弃用了。我们可能会在某个时候开放新的源代码，并在 Netflix 技术博客上发表一篇关于它的文章。

## Case study: Facebook

## 案例研究：Facebook

Facebook are advanced users of BPF, but deep details of how they run the tools fleet-wide aren't fully public. Based on the activity in bcc, and their development of the BTF and CO-RE technologies, I'd strongly suspect their solution is based on the bcc libbpf-tool versions.

Facebook 是 BPF 的高级用户，但他们如何在整个车队范围内运行工具的深入细节尚未完全公开。根据 bcc 中的活动以及他们对 BTF 和 CO-RE 技术的开发，我强烈怀疑他们的解决方案是基于 bcc libbpf-tool 版本。

## Think like a sysadmin, not like a programmer

## 像系统管理员一样思考，而不是像程序员一样思考

In summary, what we're doing here is installing the tools and building upon them, rather than rewriting everything from scratch. This is thinking like a sysadmin who installs and maintains software, and not like a programmer who codes everything. 

总之，我们在这里所做的是安装工具并在其上构建，而不是从头开始重写所有内容。这是像安装和维护软件的系统管理员一样思考，而不是像编码一切的程序员一样思考。

Install the [bcc](https://github.com/iovisor/bcc) or [bpftrace](https://github.com/iovisor/bpftrace) tools, add them to your observability product, and pull package updates as needed . That will be a quick and useful version 1. BPF up and running!

安装 [bcc](https://github.com/iovisor/bcc) 或 [bpftrace](https://github.com/iovisor/bpftrace) 工具，将它们添加到您的可观察性产品中，并根据需要拉取包更新.这将是一个快速且有用的版本 1。BPF 启动并运行！

I see people think like a programmer instead and feel they must start by learning bcc and BPF programming in depth. Then, having discovered everything is C or Python, some rewrite it all in a different language.

我看到人们像程序员一样思考，觉得他们必须从深入学习 bcc 和 BPF 编程开始。然后，在发现一切都是 C 或 Python 之后，有些人用不同的语言重写了它们。

First, learning bcc and BPF well takes weeks; Learning the subtleties and pitfalls of system tracing can take months or years. If you really want to do this and have the time, you certainly can (you'll probably wind up at tracing conferences and bumping into me: See you at Linux Plumber's or the Tracing Summit!) But if you're under some deadline to add BPF observability, try thinking like a sysadmin instead and just build upon the existing tools. That's the fast way. Think like a programmer later, if or when you have the time.

首先，学习 bcc 和 BPF 需要数周时间；了解系统跟踪的微妙之处和陷阱可能需要数月或数年的时间。如果你真的想这样做并且有时间，你当然可以（你可能会在跟踪会议上结束并撞到我：在 Linux Plumber 或 Tracing Summit 上见！）但是如果你在某个截止日期之前添加 BPF 可观察性，尝试像系统管理员一样思考，并在现有工具的基础上进行构建。这是最快的方法。如果您有时间，请稍后像程序员一样思考。

Second, the BPF software, especially certain kprobe-based tools, require ongoing maintenance. A tool may work on Linux 5.3 but break on 5.4, as a traced function was renamed or a new code path added. The BPF libraries are also evolving rapidly. I'd try not to rewrite any of these and build upon them, so you can just pull updated versions. In a previous blog post, [An Unbelievable Demo](http://www.brendangregg.com/blog/2021-06-04/an-unbelievable-demo.html), I talked about how something similar happened many years ago where old tracing tool versions were used without updates.

其次，BPF 软件，尤其是某些基于 kprobe 的工具，需要持续维护。一个工具可能在 Linux 5.3 上工作但在 5.4 上中断，因为跟踪的函数被重命名或添加了新的代码路径。 BPF 库也在迅速发展。我尽量不重写其中的任何一个并在它们的基础上进行构建，这样您就可以拉取更新的版本。在之前的博文 [An Unbelievable Demo](http://www.brendangregg.com/blog/2021-06-04/an-unbelievable-demo.html) 中，我谈到了多年前类似的事情是如何发生的使用旧的跟踪工具版本而没有更新。

For a more recent example, I wrote cachestat(8) while on vacation in 2014 for use on the Netflix cloud, which was a mix of Linux 3.2 and 3.13 at the time. BPF didn't exist on those versions, so I used basic Ftrace capabilities that were available on Linux 3.2. I described this approach as [brittle](http://www.brendangregg.com/blog/2014-12-31/linux-page-cache-hit-ratio.html) and a [sandcastle](https://github.com/brendangregg/perf-tools/blob/master/fs/cachestat) that would need maintenance as the kernel changed. It was later ported to BPF with kprobes, and has now been rewritten and included in commercial observability products. Unsurprisingly, I've heard it has problems on newer kernels. Note that I also wouldn't have even coded it this way had BPF been available on my target environment at the time. It really needs an overhaul. When I (or someone) does, anyone pulling updates from bcc will automatically get the fixed version, no effort. Those that have rewritten it will need to rewrite theirs. I fear they won't, and customers will be running a broken version of cachestat(8) for years.

举个最近的例子，我在 2014 年休假时编写了 cachestat(8) 用于 Netflix 云，当时它是 Linux 3.2 和 3.13 的混合版本。这些版本中不存在 BPF，因此我使用了 Linux 3.2 上可用的基本 Ftrace 功能。我将这种方法描述为 [脆弱](http://www.brendangregg.com/blog/2014-12-31/linux-page-cache-hit-ratio.html) 和 [sandcastle](https://github .com/brendangregg/perf-tools/blob/master/fs/cachestat），在内核更改时需要维护。它后来被 kprobes 移植到 BPF，现在已经被重写并包含在商业可观察性产品中。不出所料，我听说它在较新的内核上有问题。请注意，如果当时 BPF 在我的目标环境中可用，我什至不会以这种方式对其进行编码。它确实需要大修。当我（或某人)这样做时，任何从 bcc 提取更新的人都会自动获得固定版本，无需任何努力。那些已经重写它的人将需要重写他们的。我担心他们不会，而且客户会运行一个损坏的 cachestat(8) 版本多年。

The problems I'm describing are specific to BPF software and kernel tracing. As a different example, my flame graph software has been rewritten over a dozen times, and since it's a simple and finished algorithm I don't see a big problem with that. I prefer people help with the newer [d3 version](https://github.com/spiermar/d3-flame-graph), but if people do their own it's no big deal. You can code it and it'll work forever. That's not the case with the kprobe-based BPF tools, because they do need maintenance. I'd rewrite these tools using tracepoints instead, as they have a stable API which in theory would alleviate this issue, but tracepoints aren't always available where you need them.

我所描述的问题特定于 BPF 软件和内核跟踪。作为一个不同的例子，我的火焰图软件已经被重写了十几次，而且由于它是一个简单而完善的算法，我认为这没有什么大问题。我更喜欢人们帮助使用较新的 [d3 版本](https://github.com/spiermar/d3-flame-graph)，但如果人们自己动手，那没什么大不了的。您可以对其进行编码，它将永远有效。基于 kprobe 的 BPF 工具并非如此，因为它们确实需要维护。我会改用跟踪点重写这些工具，因为它们有一个稳定的 API，理论上可以缓解这个问题，但跟踪点并不总是在您需要它们的地方可用。

The BPF libraries and frameworks are also changing and evolving, most recently with the BTF and CO-RE support. This is something I hope people consider before choosing to rewrite them: Do you have a plan to rewrite all the updates as well, or will you end up stuck on an old port of the library? 

BPF 库和框架也在不断变化和发展，最近随着 BTF 和 CO-RE 的支持。这是我希望人们在选择重写它们之前考虑的事情：您是否也有重写所有更新的计划，或者您最终会停留在库的旧端口上吗？

What if you have a great idea for a _better_ BPF library or framework than what we're using in bcc and bpftrace? Talk to us, try it out, innovate. We're at the start of the BPF era and there's lots more to explore. But please understand what exists first, and understand the maintenance burden you are taking on. Your energies may be better spent creating something new, on top of what exists, than porting something old.

如果你有一个比我们在 bcc 和 bpftrace 中使用的更好的 BPF 库或框架的好主意怎么办？与我们交谈、尝试、创新。我们正处于 BPF 时代的开始，还有很多东西需要探索。但请先了解存在的是什么，并了解您所承担的维护负担。与移植旧事物相比，您的精力可能会更好地用于在现有事物的基础上创造新事物。

* * *

* * *

_You can comment here, but I can't guarantee your comment will remain here forever: I might switch comment systems at some point (eg, if disqus add advertisements)._

_你可以在这里评论，但我不能保证你的评论会永远留在这里：我可能会在某个时候切换评论系统（例如，如果disqus添加广告）。_

[comments powered by Disqus](http://disqus.com) 

[评论由 Disqus 提供](http://disqus.com)

