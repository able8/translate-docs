# What Is a Standard Container (2021 edition)

# 什么是标准容器（2021 版）

September 5, 2021

[Containers](http://iximiuz.com/en/categories/?category=Containers)

[容器](http://iximiuz.com/en/categories/?category=Containers)

**TL;DR** Per [OCI Runtime Specification](https://github.com/opencontainers/runtime-spec):

**TL;DR** 根据 [OCI 运行时规范](https://github.com/opencontainers/runtime-spec)：

- Containers are isolated and restricted boxes for running processes 📦
- Containers pack an app and all its dependencies (including OS libs) together
- Containers are for portability - any compliant runtime can run_standard_ containers
- Containers can be implemented using Linux, Windows, and other OS-es
- Virtual Machines also can be used as_standard_ containers 🤐

- 容器是用于运行进程的隔离和受限的盒子📦
- 容器将应用程序及其所有依赖项（包括操作系统库）打包在一起
- 容器是为了可移植性 - 任何兼容的运行时都可以运行_标准_容器
- 容器可以使用 Linux、Windows 和其他操作系统实现
- 虚拟机也可以用作标准容器🤐

There are many ways to create containers, especially on Linux and alike. Besides the super widespread Docker implementation, you may have heard about [LXC](https://github.com/lxc/lxc), [systemd-nspawn](https://www.linux.org/docs/man1/systemd-nspawn.html), or maybe even [OpenVZ](https://en.wikipedia.org/wiki/OpenVZ).

有很多方法可以创建容器，尤其是在 Linux 等上。除了超级广泛的 Docker 实现，你可能听说过 [LXC](https://github.com/lxc/lxc)、[systemd-nspawn](https://www.linux.org/docs/man1/systemd-nspawn.html)，或者甚至 [OpenVZ](https://en.wikipedia.org/wiki/OpenVZ)。

The general concept of the container is quite vague. What's true and what's not often depends on the context, but the context itself isn't always given explicitly. For instance, there is a common saying that [containers are Linux processes](https://www.redhat.com/en/blog/containers-are-linux) or that [containers aren't Virtual Machines](https://docs.microsoft.com/en-us/virtualization/windowscontainers/about/containers-vs-vm). However, the first statement is just an oversimplified attempt to explain Linux containers. And the second statement simply isn't always true.

容器的一般概念是相当模糊的。什么是真的，什么不是，通常取决于上下文，但上下文本身并不总是明确给出。例如，有一句俗语说[容器是Linux进程](https://www.redhat.com/en/blog/containers-are-linux)或[容器不是虚拟机](https://www.redhat.com/en/blog/containers-are-linux)。然而，第一个陈述只是对解释 Linux 容器的过于简单化的尝试。第二个陈述并不总是正确的。

In this article, I'm not trying to review all possible ways of creating containers. Instead, the article is an analysis of the [OCI Runtime Specification](https://github.com/opencontainers/runtime-spec). The spec turned out to be an insightful read! For instance, it gives a definition of the _standard container_ (and no, it's not a process) and sheds some light on _when Virtual Machines can be considered containers_.

在本文中，我不是要回顾创建容器的所有可能方法。相反，文章是对 [OCI 运行时规范](https://github.com/opencontainers/runtime-spec)的分析。结果证明该规范是一个有见地的阅读！例如，它给出了_标准容器_的定义（不，它不是一个进程)并阐明了_何时可以将虚拟机视为容器_。

Containers, as they brought to us by Docker and Podman, are OCI-compliant. Today we even use the terms _container_, _Docker container_, and _Linux container_ interchangeably. However, this is just one type of OCI-compliant container. So, let's take a closer look at the [OCI Runtime Specification](https://github.com/opencontainers/runtime-spec).

Docker 和 Podman 带给我们的容器是符合 OCI 的。今天，我们甚至可以互换使用术语 _container_、_Docker 容器_ 和 _Linux 容器_。然而，这只是一种符合 OCI 的容器。那么，让我们仔细看看 [OCI 运行时规范](https://github.com/opencontainers/runtime-spec)。

## What is Open Container Initiative (OCI)

## 什么是开放容器计划（OCI）

Open Container Initiative (OCI) is an open governance structure that was established in 2015 by Docker and other prominent players of the container industry to _[express the purpose of creating open industry standards around container formats and runtimes](https://opencontainers.org/)_. In other words, _[OCI develops specifications for standards on Operating System process and application containers.](https://github.com/opencontainers/runtime-spec)_

开放容器倡议 (OCI) 是一个开放的治理结构，由 Docker 和其他容器行业的知名参与者于 2015 年建立，目的是_[表达围绕容器格式和运行时创建开放行业标准的目的](https://opencontainers.组织/)_。换句话说，_[OCI为操作系统进程和应用程序容器的标准制定规范。](https://github.com/opencontainers/runtime-spec)_

Also think it's too many fancy words for such a small paragraph?

还觉得这么小的一段话太多花哨的词吗？

Here is how I understand it. By 2015 Docker already gained quite some popularity, but there were other competing projects implementing their own containers like [rkt](https://github.com/rkt/rkt) and [lmctfy](https://github.com/google/lmctfy). Apparently, the OCI was established _to standardize the way of doing containers_. De facto, it made the Docker's container implementation a standard one, but some non-Docker parts were incorporated too.

这是我的理解。到 2015 年 Docker 已经获得了相当大的知名度，但还有其他竞争项目实现了自己的容器，如 [rkt](https://github.com/rkt/rkt) 和 [lmctfy](https://github.com/google/lmctfy)。显然，OCI 的建立是为了规范容器的处理方式。事实上，它使 Docker 的容器实现成为一个标准的实现，但也包含了一些非 Docker 部分。

## What is an OCI Container

## 什么是 OCI 容器

So, [how does OCI define a _Container_ nowadays](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/glossary.md#container)?

那么，[现在 OCI 如何定义 _Container_](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/glossary.md#container)？

> A Standard Container is an environment for executing processes with configurable isolation and resource limitations.

> 标准容器是用于执行具有可配置隔离和资源限制的进程的环境。

[Why do we even need containers?](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/principles.md#the-5-principles-of-standard-containers)

[为什么我们甚至需要容器？](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/principles.md#the-5-principles-of-standard-containers)

> [To] define a unit of software delivery ... The **goal of a Standard Container is to encapsulate a software component and all its dependencies** in a format that is self-describing and portable, so that **any compliant runtime can run it** without extra dependencies, regardless of the underlying machine and the contents of the container.

> [To] 定义一个软件交付单元......标准容器的**目标是以自描述和可移植的格式封装软件组件及其所有依赖项**，以便**任何兼容运行时可以在没有额外依赖的情况下运行它**，而不管底层机器和容器的内容。

Ok, and [what can we do with containers?](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md) 

好的，[我们可以用容器做什么？](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md)

> [Containers] can be created, started, and stopped using standard container tools; copied and snapshotted using standard filesystem tools; and downloaded and uploaded using standard network tools.

> [Containers] 可以使用标准容器工具创建、启动和停止；使用标准文件系统工具复制和快照；并使用标准网络工具下载和上传。

![Containers work the same way on developer's laptop, CI/CD servers, and Kubernetes clusters running in the cloud.](http://iximiuz.com/oci-containers/container-2000-opt.png)

_Operations on containers that OCI runtimes must support: [Create](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md#create), [Start](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md#start), [Kill](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md#kill), [Delete]( https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md#delete), and [Query State](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md#query-state)._

_OCI 运行时必须支持的容器操作：[创建](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md#create)，[开始](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/runtime.md#start), [杀死](https://github.com/opencontainers/runtime-spec/blob/20a2d97827e0812ebe1515f73736c6a0c/runtime.md#start)，和[查询状态](https://github.com/opencontainers/runtime-spec/blob/2820a7370c7820c78c720c78c720c78c780c78c15c6a0c/runtime.md#delete)._

Well, makes sense. But... **a container cannot be a process** then! In accordance with the OCI Runtime Spec, it's more like _an isolated and restricted box_ for running one or more processes inside.

嗯，有道理。但是...... **容器不能是进程** 那么！根据 OCI 运行时规范，它更像是_一个隔离且受限制的盒子_，用于在内部运行一个或多个进程。

## Linux Container vs. Other Containers

## Linux 容器与其他容器

Apart from the container's operations and lifecycle, the [OCI Runtime Spec also specifies the container's configuration and execution environment](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/spec.md#abstract).

除了容器的操作和生命周期， [OCI Runtime Spec 还规定了容器的配置和执行环境](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/spec.md#abstract)。

Per the OCI Runtime Spec, to create a container, one needs to provide a runtime with a so-called [_filesystem bundle_](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/bundle.md) that consists of a mandatory [`config.json`](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config.md) file and an optional folder holding the future container's root filesystem.

根据 OCI 运行时规范，要创建容器，需要为运行时提供所谓的 [_filesystem bundle_](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/thatbundle.md)由一个强制性的[`config.json`](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config.md) 文件和一个包含未来容器根文件系统的可选文件夹组成。

_**Off-topic**: A bundle is usually obtained by unpacking a container image, but images aren't a part of the Runtime Spec. Instead, they are subject to the dedicated [OCI Image Specification](https://github.com/opencontainers/image-spec)._

_**Off-topic**：通常通过解包容器镜像获得包，但镜像不是运行时规范的一部分。相反，它们受专用的 [OCI 图像规范](https://github.com/opencontainers/image-spec)._

`config.json` contains data necessary to implement standard operations against the container (Create, Start, Query State, Kill, and Delete). But things start getting really interesting when it comes to the actual structure of the `config.json` file.

`config.json` 包含对容器实施标准操作（创建、启动、查询状态、终止和删除）所需的数据。但是当涉及到 `config.json` 文件的实际结构时，事情开始变得非常有趣。

The configuration consists of the _common_ and [_platform-specific_](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config.md#platform-specific-configuration) sections. The common section includes `ociVersion`, `root` filesystem path within the bundle, additional `mounts` beyond the `root`, a `process` to start in the container, a `user`, and a `hostname`. Hm... but where are the famous namespaces and cgroups?

配置由 _common_ 和 [_platform-specific_](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config.md#platform-specific-configuration) 部分组成。公共部分包括`ociVersion`、包内的`root` 文件系统路径、`root` 之外的附加`mounts`、在容器中启动的`process`、`user` 和`hostname`。嗯...但是著名的命名空间和 cgroup 在哪里？

By the time of writing this article, OCI Runtime Spec defines containers for the following platforms: [Linux](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-linux.md), [Solaris](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-solaris.md), [Windows](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-windows.md), [z/OS](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-zos.md), and [Virtual Machine](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-vm.md).

在撰写本文时，OCI 运行时规范为以下平台定义了容器：[Linux](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-linux.md)、[Solaris](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-solaris.md), [Windows](https://github.com/opencontainers/runtime-spec/blob/202986ec2a7e0812ebe1515f73736c6a0c / windows.md)、[z/OS](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-zos.md)和[虚拟机](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-vm.md)。

_Wait, what?! VMs are Containers??!_ 🤯 

_等等，什么？！ VM 是容器？？！_ 🤯

In particular, the [Linux-specific section](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-linux.md) brings in (among other things) pid, network, mount, ipc, uts, and user namespaces, control groups, and seccomp. In contrast, the [Windows-specific](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-windows.md) section comes with its own isolation and restriction mechanisms provided by the [Windows Host Compute Service (HCS)](https://docs.microsoft.com/en-us/virtualization/community/team-blog/2017/20170127-introducing-the-host-compute-service-hcs).

特别是，[Linux 特定部分](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-linux.md)引入了（除其他外)pid、network、mount、ipc、 uts、用户命名空间、控制组和 seccomp。相比之下，[Windows 特定](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-windows.md) 部分带有自己的隔离和限制机制，由 [Windows Host Compute服务 (HCS)](https://docs.microsoft.com/en-us/virtualization/community/team-blog/2017/20170127-introducing-the-host-compute-service-hcs)。

![OCI Runtime config.json consists of the common and platform-specific parts](http://iximiuz.com/oci-containers/config.json-2000-opt.png)

Thus, **only Linux containers rely on namespaces and cgroups. However, not all standard containers are Linux.**

因此，**只有 Linux 容器依赖于命名空间和 cgroup。但是，并非所有标准容器都是 Linux。**

## Virtual Machines vs. Containers

## 虚拟机与容器

The most widely-used OCI runtimes are [runc](https://github.com/opencontainers/runc) and [crun](https://github.com/containers/crun). Unsurprisingly, both implement Linux containers. But as we just saw, OCI Runtime Spec mentions Windows, Solaris, and other containers. And what's even more intriguing for me, it defines VM containers!

最广泛使用的 OCI 运行时是 [runc](https://github.com/opencontainers/runc) 和 [crun](https://github.com/containers/crun)。不出所料，两者都实现了 Linux 容器。但正如我们刚刚看到的，OCI 运行时规范提到了 Windows、Solaris 和其他容器。对我来说更有趣的是，它定义了 VM 容器！

_Aren't containers were meant to replace VMs as a more lightweight implementation of the same [execution environment](https://en.wikipedia.org/wiki/Computing_platform) abstraction?_

_容器不是要取代虚拟机作为相同[执行环境](https://en.wikipedia.org/wiki/Computing_platform) 抽象的更轻量级实现吗？_

Anyways, let's take a closer look at VM containers.

不管怎样，让我们仔细看看 VM 容器。

Clearly, they are not backed by Linux namespaces and cgroups. Instead, the [Virtual-machine-specific Container Configuration](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-vm.md) mentions a hypervisor, a kernel, and a VM image. So, the isolation is achieved by virtualizing some hardware (hypervisor) and then booting a full-fledged OS (kernel + image) on top of it. The resulting environment is our box, i.e., a container.

显然，它们不受 Linux 命名空间和 cgroup 的支持。相反，[特定于虚拟机的容器配置](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/config-vm.md) 提到了管理程序、内核和 VM 映像。因此，隔离是通过虚拟化某些硬件（管理程序）然后在其上启动成熟的操作系统（内核 + 映像)来实现的。由此产生的环境是我们的盒子，即一个容器。

![Linux containers vs. VM containers.](http://iximiuz.com/oci-containers/linux-container-vs-vm-container-2000-opt.png)

Notice that the VM image mentioned by the OCI Runtime Spec has nothing to do with the traditional container image that is used to create a bundle. The bundle root filesystem is mounted into a VM container separately.

请注意，OCI 运行时规范中提到的 VM 映像与用于创建包的传统容器映像无关。捆绑根文件系统单独挂载到 VM 容器中。

Thus, VM-based containers is a thing!

因此，基于 VM 的容器是一回事！

However, [the only non-deprecated implementation of OCI VM containers](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/implementations.md#runtime-virtual-machine), Kata containers, [has the following in its FAQ](https://katacontainers.io/learn/):

然而，[OCI VM 容器的唯一未弃用的实现](https://github.com/opencontainers/runtime-spec/blob/20a2d9782986ec2a7e0812ebe1515f73736c6a0c/implementations.md#runtime-virtual-machine)，Kata容器在其常见问题解答中关注：

> Kata Containers is still in its formational stages, but the technical basis for the project - Clear Containers and runV - are used globally at enterprise scale by organizations like JD.com, China's largest ecommerce company (by revenue).

> Kata Containers 仍处于形成阶段，但该项目的技术基础——Clear Containers 和 runV——被中国最大的电子商务公司京东等组织在全球企业范围内使用（按收入计算）。

That is, the good old Linux containers remain the default production choice. So, containers are still just [boxed] processes. 

也就是说，好的旧 Linux 容器仍然是默认的生产选择。因此，容器仍然只是 [盒装] 进程。

**UPD:** Adel Zaalouk ( [@ZaNetworker](https://twitter.com/ZaNetworker)) kindly pointed me to the [OpenShift Sandboxed Containers project](https://docs.openshift.com/container-platform/4.8/sandboxed_containers/understanding-sandboxed-containers.html). It's an attempt to make Kubernetes Open Shift workloads more secure. Long story short, it uses Kata Containers to run Kubernetes Open Shift Pods inside lightweight Virtual Machines. And it's already in the technology preview mode. Here is a [nice intro](https://cloud.redhat.com/blog/the-dawn-of-openshift-sandboxed-containers-overview) and the [coolest diagram ever showing _how_ in great detail](https://raw.githubusercontent.com/kata-containers/kata-containers/9bbaa66f3973a91e79752a4237d3ac79e80ab47c/docs/design/arch-images/katacontainers-e2e-with-bg.jpg) (and [its interactive counterpart](https://www.thinglink.com/card/1401236075678007299)). Do have this addition in mind while reading the following section :)

**UPD：** Adel Zaalouk ( [@ZaNetworker](https://twitter.com/ZaNetworker)) 亲切地向我指出了 [OpenShift Sandboxed Containers 项目](https://docs.openshift.com/container-platform /4.8/sandboxed_containers/understanding-sandboxed-containers.html）。它试图使 Kubernetes Open Shift 工作负载更加安全。长话短说，它使用 Kata Containers 在轻量级虚拟机中运行 Kubernetes Open Shift Pod。而且它已经处于技术预览模式。这是一个[不错的介绍](https://cloud.redhat.com/blog/the-dawn-of-openshift-sandboxed-containers-overview)和[有史以来最酷的图表详细地展示了_how_](https://raw.githubusercontent.com/kata-containers/kata-containers/9bbaa66f3973a91e79752a4237d3ac79e80ab47c/docs/design/arch-images/katacontainers-e2e-with-bg.jpg)（和[其互动对应物](https://www..com/card/1401236075678007299))。在阅读以下部分时，请记住这一点:)

## MicroVMs vs. Containers

## MicroVM 与容器

One of the coolest parts of Linux containers is that they are much more lightweight than Virtual Machines. The startup time is under a second, and there is almost no space and runtime overhead. However, their strongest part is their weakness as well. The Linux containers are so fast because they are regular Linux processes. So, they are as secure as the underlying Linux host. Thus, Linux containers are only good for trusted workloads.

Linux 容器最酷的部分之一是它们比虚拟机轻得多。启动时间不到一秒，几乎没有空间和运行时开销。然而，他们最强大的部分也是他们的弱点。 Linux 容器之所以如此之快，是因为它们是常规的 Linux 进程。因此，它们与底层 Linux 主机一样安全。因此，Linux 容器仅适用于受信任的工作负载。

Since shared infrastructure becomes more and more common, the need for stronger isolation remains. Serverless/FaaS computing is probably one of the most prominent examples. By running code in AWS Lambda or alike, you just don't deal with the _server_ abstraction anymore. Hence, there is no need for virtual machines or containers for development teams. But from the platform provider standpoint, using Linux containers to run workloads of different customers on the same host would be a security nightmare. Instead, functions need to be run in something as lightweight as Linux containers and as secure as Virtual Machines.

由于共享基础设施变得越来越普遍，因此仍然需要更强大的隔离。无服务器/FaaS 计算可能是最突出的例子之一。通过在 AWS Lambda 或类似环境中运行代码，您只需不再处理 _server_ 抽象。因此，开发团队不需要虚拟机或容器。但从平台提供商的角度来看，使用 Linux 容器在同一主机上运行不同客户的工作负载将是一场安全噩梦。相反，函数需要在像 Linux 容器一样轻量级和像虚拟机一样安全的东西中运行。

[AWS Firecracker to the rescue!](https://aws.amazon.com/blogs/aws/firecracker-lightweight-virtualization-for-serverless-computing/)

[AWS Firecracker 来救援！](https://aws.amazon.com/blogs/aws/firecracker-lightweight-virtualization-for-serverless-computing/)

> The main component of Firecracker is a virtual machine monitor (VMM) that uses the Linux Kernel Virtual Machine (KVM) to create and run microVMs. Firecracker has a minimalist design. It excludes unnecessary devices and guest-facing functionality to reduce the memory footprint and attack surface area of each microVM. This improves security, decreases the startup time, and increases hardware utilization. Firecracker has also been integrated in container runtimes, for example Kata Containers and Weaveworks Ignite.

> Firecracker 的主要组件是一个虚拟机监视器 (VMM)，它使用 Linux 内核虚拟机 (KVM) 来创建和运行 microVM。 Firecracker 拥有简约的设计。它排除了不必要的设备和面向访客的功能，以减少每个 microVM 的内存占用和攻击面区域。这提高了安全性，减少了启动时间，并提高了硬件利用率。 Firecracker 也已集成到容器运行时中，例如 Kata Containers 和 Weaveworks Ignite。

But surprisingly or not, Firecracker is not OCI-compliant runtime on itself... However, there seems to be a way to put an OCI runtime into a Firecracker microVM and get the best of all worlds - portability of containers, lightness of Firecracker microVMs , and full isolation from the host operating system. I'm definitely going to have a closer look at this option and I'll share my finding in the future posts.

但无论是否令人惊讶，Firecracker 本身并不是 OCI 兼容的运行时......然而，似乎有一种方法可以将 OCI 运行时放入 Firecracker 微虚拟机中并获得所有领域的最佳效果 - 容器的可移植性、Firecracker 微虚拟机的轻便性，并与主机操作系统完全隔离。我肯定会仔细研究这个选项，我会在以后的帖子中分享我的发现。

_**UPD:** Check out this [awesome drawing](https://raw.githubusercontent.com/kata-containers/kata-containers/9bbaa66f3973a91e79752a4237d3ac79e80ab47c/docs/design/arch-images/katacontainers-e2e-with-bg.jpg) showing how Kata Containers does it. In particular, Kata Containers can use Firecracker as VMM._

_**UPD:** 看看这个[很棒的图画](https://raw.githubusercontent.com/kata-containers/kata-containers/9bbaa66f3973a91e79752a4237d3ac79e80ab47c/docs/design/arch-images/katacontainers-bge.jpg) 展示了 Kata Containers 是如何做到的。特别是 Kata Containers 可以使用 Firecracker 作为 VMM。_

Another interesting project in the area of secure containers is Google's [gVisor](https://github.com/google/gvisor):

安全容器领域另一个有趣的项目是 Google 的 [gVisor](https://github.com/google/gvisor)：

> gVisor is an application kernel, written in Go, that implements a substantial portion of the Linux system surface. It includes an Open Container Initiative (OCI) runtime called runsc that provides an isolation boundary between the application and the host kernel. The runsc runtime integrates with Docker and Kubernetes, making it simple to run sandboxed containers.

> gVisor 是一个用 Go 编写的应用程序内核，它实现了 Linux 系统表面的很大一部分。它包括一个名为 runc 的开放容器计划 (OCI) 运行时，它在应用程序和主机内核之间提供隔离边界。 runc 运行时与 Docker 和 Kubernetes 集成，使运行沙盒容器变得简单。

Unlike Firecracker, gVisor provides an OCI-complaint runtime. But there is no full-fledged hypervisor like KVM for gVisor-backed containers. Instead, it emulates the kernel in the user-space. Sounds pretty cool, but the runtime overhead is likely to be noticeable.

与 Firecracker 不同，gVisor 提供了一个 OCI-complaint 运行时。但是对于 gVisor 支持的容器，没有像 KVM 这样成熟的虚拟机管理程序。相反，它模拟用户空间中的内核。听起来很酷，但运行时开销可能很明显。

## Instead of conclusion 

## 而不是结论

To summarize, containers aren't just slightly more isolated and restricted Linux processes. Instead, they are standardized execution environments improving workload portability. Linux containers are the most widespread form of containers nowadays, but the need for more secure containers is growing. The OCI Runtime Spec defines the VM-backed containers, and the Kata project makes them real. So, it's an exciting time to explore the containerverse!

总而言之，容器不仅仅是稍微孤立和受限的 Linux 进程。相反，它们是标准化的执行环境，可提高工作负载的可移植性。 Linux 容器是当今最普遍的容器形式，但对更安全的容器的需求正在增长。 OCI 运行时规范定义了 VM 支持的容器，而 Kata 项目使它们成为现实。所以，现在是探索容器世界的激动人心的时刻！

### Further reading

### 进一步阅读

- [A journey from containerization to orchestration and beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)
- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [Implementing Container Runtime Shim: runc](http://iximiuz.com/en/posts/implementing-container-runtime-shim/)

- [从容器化到编排及其他的旅程](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)
- [并非每个容器内部都有操作系统](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [你不需要图像来运行容器](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [你需要容器来构建镜像](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [实现容器运行时 Shim: runc](http://iximiuz.com/en/posts/implementing-container-runtime-shim/)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

