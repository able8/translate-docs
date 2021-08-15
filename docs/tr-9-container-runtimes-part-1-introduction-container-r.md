# Container Runtimes Part 1: An Introduction to Container Runtimes

# 容器运行时第 1 部分：容器运行时简介

>  From https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r

One of the terms you hear a lot when dealing with containers is "container runtime". "Container runtime" can have different meanings to different people so it's no wonder that it's such a confusing and  vaguely understood term, even within the container community.

在处理容器时，您经常听到的术语之一是“容器运行时”。 “容器运行时”对不同的人可能有不同的含义，因此即使在容器社区内，它也是一个如此令人困惑和模糊理解的术语也就不足为奇了。

This post is the first in a series that will be in four parts:
 1. Part 1: Intro to Container Runtimes: why are they so confusing?
 2. Part 2: Deep Dive into Low-Level Runtimes
 3. Part 3: Deep Dive into High-Level Runtimes
 4. Part 4: Kubernetes Runtimes and the CRI

这篇文章是系列文章的第一篇，该系列文章将分为四个部分：
1. 第 1 部分：容器运行时介绍：为什么它们如此令人困惑？
2. 第 2 部分：深入研究低级运行时
3. 第 3 部分：深入研究高级运行时
4. 第 4 部分：Kubernetes 运行时和 CRI

This post will explain what container runtimes are and why there is  so much confusion. I will then dive into different types of container  runtimes, what they do, and how they are different from each other.

这篇文章将解释什么是容器运行时以及为什么会有这么多的混乱。然后，我将深入探讨不同类型的容器运行时、它们的作用以及它们之间的不同之处。

Traditionally, a computer programmer might know "runtime" as either  the lifecycle phase when a program is running, or the specific  implementation of a language that supports its execution. An example  might be the Java HotSpot runtime. This latter meaning is the closest to "container runtime". A container runtime is responsible for all the  parts of running a container that isn't actually running the program  itself. As we will see throughout this series, runtimes implement  varying levels of features, but running a container is actually all  that's required to call something a container runtime.

传统上，计算机程序员可能将“运行时”理解为程序运行时的生命周期阶段，或者支持其执行的语言的特定实现。一个例子可能是 Java HotSpot 运行时。后一种含义最接近“容器运行时”。容器运行时负责运行容器的所有部分，这些部分实际上并未运行程序本身。正如我们将在本系列中看到的，运行时实现不同级别的功能，但运行容器实际上就是调用容器运行时所需的全部内容。

If you're not super familiar with containers, check out these links first and come back:
 - [What even is a container: namespaces and cgroups](https://jvns.ca/blog/2016/10/10/what-even-is-a-container/)
 - [Cgroups, namespaces, and beyond: what are containers made from?](https://www.youtube.com/watch?v=sK5i-N34im8)

如果您对容器不是很熟悉，请先查看这些链接，然后再回来：
- [什么是容器：命名空间和 cgroups](https://jvns.ca/blog/2016/10/10/what-even-is-a-container/)
- [Cgroups、命名空间等：容器是由什么制成的？](https://www.youtube.com/watch?v=sK5i-N34im8)

## Why are Container Runtimes so Confusing?

## 为什么容器运行时如此令人困惑？

Docker was released in 2013 and solved many of the problems that  developers had running containers end-to-end. It had all these things:
 - A container image format
 - A method for building container images (Dockerfile/docker build)
 - A way to manage container images (docker images, docker rm, etc.)
 - A way to manage instances of containers (docker ps, docker rm , etc.)
 - A way to share container images (docker push/pull)
 - A way to run containers (docker run)

Docker 于 2013 年发布，解决了开发人员在端到端运行容器时遇到的许多问题。它有所有这些东西：
- 一种容器镜像格式
- 一种构建容器镜像的方法（Dockerfile/docker build）
- 一种管理容器实例的方法（docker ps、docker rm 等）
- 一种共享容器镜像的方法（docker push/pull）
- 一种运行容器的方法（docker run）

At the time, Docker was a monolithic system. However, none of these  features were really dependent on each other. Each of these could be  implemented in smaller and more focused tools that could be used  together. Each of the tools could work together by using a common  format, a container standard.

当时，Docker 是一个单体系统。但是，这些功能中没有一个是真正相互依赖的。这些中的每一个都可以在可以一起使用的更小、更集中的工具中实现。每个工具都可以通过使用一种通用格式、一种容器标准来协同工作。

Because of that, Docker, Google, CoreOS, and other vendors created the [Open Container Initiative (OCI)](https://www.opencontainers.org/). They then broke out their code for running containers as a tool and library called [runc](https://github.com/opencontainers/runc) and donated it to OCI as a reference implementation of the [OCI runtime specification](https: //github.com/opencontainers/runtime-spec).

因此，Docker、Google、CoreOS 和其他供应商创建了[开放容器计划 (OCI)](https://www.opencontainers.org/)。然后，他们将用于运行容器的代码作为一个名为 [runc](https://github.com/opencontainers/runc) 的工具和库分解出来，并将其捐赠给 OCI 作为 [OCI 运行时规范](https: //github.com/opencontainers/runtime-spec)。

It was initially confusing what Docker had contributed to OCI. What  they contributed was a standard way to "run" containers but nothing  more. They didn't include the image format or registry push/pull  formats. When you run a Docker container, these are the steps Docker  actually goes through:
 1. Download the image
 2. Unpack the image into a "bundle". This flattens the layers into a single filesystem.
 3. Run the container from the bundle

最初，人们对 Docker 对 OCI 的贡献感到困惑。他们贡献的是一种“运行”容器的标准方式，仅此而已。它们不包括镜像格式或注册表推/拉格式。当你运行一个 Docker 容器时，这些是 Docker 实际经历的步骤：

1. 下载镜像

1. 将镜像解压成一个“包”。这将层扁平化为单个文件系统。

3. 从 bundle 中运行容器

What Docker standardized was only #3. Until that was clarified,  everyone had thought of a container runtime as supporting all of the  features Docker supported. Eventually, Docker folks clarified that the [original spec](https://github.com/opencontainers/runtime-spec/commit/77d44b10d5df53ee63f0768cd0a29ef49bad56b6#diff-b84a8d65d8ed53f4794cd2db7e8ea731R45) stated that only the "running the container" part that made up the  runtime. This is a disconnect that continues even today, and makes  "container runtimes" such a confusing topic. I'll hopefully show that  neither side is totally wrong and I'll use the term pretty broadly in  this blog post.

Docker 标准化的只是#3。在澄清之前，每个人都认为容器运行时支持 Docker 支持的所有功能。最终，Docker 人员澄清了 [原始规范](https://github.com/opencontainers/runtime-spec/commit/77d44b10d5df53ee63f0768cd0a29ef49bad56b6#diff-b84a8d65d8ed53f4794cd2db7e8ea731R45) 。即使在今天，这种脱节仍在继续，并使“容器运行时”成为一个令人困惑的话题。我希望表明双方都没有完全错误，我将在这篇博文中广泛使用该术语。

## Low-Level and High-Level Container Runtimes 
## 低级和高级容器运行时
When folks think of container runtimes, a list of examples might come to mind; runc, lxc, lmctfy, Docker (containerd), rkt, cri-o. Each of  these is built for different situations and implements different  features. Some, like containerd and cri-o, actually use runc to run the  container but implement image management and APIs on top. You can think  of these features -- which include image transport, image management,  image unpacking, and APIs -- as high-level features as compared to  runc's low-level implementation.

当人们想到容器运行时，可能会想到一系列示例； runc、lxc、lmctfy、Docker（容器）、rkt、cri-o。这些中的每一个都是为不同的情况而构建的，并实现了不同的功能。有些，如 containerd 和 cri-o，实际上使用 runc 来运行容器，但在顶部实现镜像管理和 API。与 runc 的低级实现相比，您可以将这些功能（包括镜像传输、镜像管理、镜像解包和 API）视为高级功能。

With that in mind you can see that the container runtime space is  fairly complicated. Each runtime covers different parts of this  low-level to high-level spectrum. Here is a very subjective diagram:

考虑到这一点，您可以看到容器运行时空间相当复杂。每个运行时都涵盖了这个低级到高级频谱的不同部分。这是一个非常主观的图表：

![img](https://storage.googleapis.com/static.ianlewis.org/prod/img/768/runtimes.png)

So for practical purposes, actual container runtimes that focus on  just running containers are usually referred to as "low-level container  runtimes". Runtimes that support more high-level features, like image  management and gRPC/Web APIs, are usually referred to as "high-level  container tools", "high-level container runtimes" or usually just  "container runtimes". I'll refer to them as "high-level container  runtimes". It's important to note that low-level runtimes and high-level runtimes are fundamentally different things that solve different  problems.

因此，出于实际目的，专注于运行容器的实际容器运行时通常被称为“低级容器运行时”。支持更多高级功能的运行时，如镜像管理和 gRPC/Web API，通常被称为“高级容器工具”、“高级容器运行时”或通常简称为“容器运行时”。我将它们称为“高级容器运行时”。重要的是要注意，低级运行时和高级运行时是解决不同问题的根本不同的东西。

Containers are implemented using [Linux namespaces](https://en.wikipedia.org/wiki/Linux_namespaces) and [cgroups](https://en.wikipedia.org/wiki/Cgroups). Namespaces let you virtualize system resources, like the file system or networking, for each container. Cgroups provide a way to limit the  amount of resources like CPU and memory that each container can use. At  the lowest level, container runtimes are responsible for setting up  these namespaces and cgroups for containers, and then running commands  inside those namespaces and cgroups. Low-level runtimes support using  these operating system features.

容器是使用 [Linux 命名空间](https://en.wikipedia.org/wiki/Linux_namespaces) 和 [cgroups](https://en.wikipedia.org/wiki/Cgroups) 实现的。命名空间让您可以为每个容器虚拟化系统资源，例如文件系统或网络。 Cgroups 提供了一种方法来限制每个容器可以使用的资源量，例如 CPU 和内存。在最低级别，容器运行时负责为容器设置这些命名空间和 cgroup，然后在这些命名空间和 cgroup 中运行命令。低级运行时支持使用这些操作系统功能。

Typically, developers who want to run apps in containers will need  more than just the features that low-level runtimes provide. They need  APIs and features around image formats, image management, and sharing  images. These features are provided by high-level runtimes. Low-level  runtimes just don't provide enough features for this everyday use. For  that reason the only folks that will actually use low-level runtimes  would be developers who implement higher level runtimes, and tools for  containers.

通常，想要在容器中运行应用程序的开发人员需要的不仅仅是低级运行时提供的功能。他们需要有关镜像格式、镜像管理和共享镜像的 API 和功能。这些功能由高级运行时提供。低级运行时无法为日常使用提供足够的功能。出于这个原因，真正使用低级运行时的人将是实现更高级别运行时和容器工具的开发人员。

Developers who implement low-level runtimes will say that higher  level runtimes like containerd and cri-o are not actually container  runtimes, as from their perspective they outsource the implementation of running a container to runc. But, from the user's perspective, they are a singular component that provides the ability to run containers. One  implementation can be swapped out for another, so it still makes sense  to call it a runtime from that perspective. Even though containerd and  cri-o both use runc, they are very different projects that have very  different feature support.

实现低级运行时的开发人员会说，像 containerd 和 cri-o 这样的高级运行时实际上并不是容器运行时，因为从他们的角度来看，他们将运行容器的实现外包给了 runc。但是，从用户的角度来看，它们是提供运行容器能力的单一组件。一种实现可以换成另一种实现，因此从这个角度来看，将其称为运行时仍然是有意义的。尽管 containerd 和 cri-o 都使用 runc，但它们是非常不同的项目，具有非常不同的功能支持。

## 'Til Next Time

## '直到下一次

I hope that helped explain container runtimes and why they are so hard to understand. Feel free to leave me comments below or [on Twitter](https://twitter.com/IanMLewis) and let me know what about container runtimes was hardest for you to understand.

我希望这有助于解释容器运行时以及为什么它们如此难以理解。请随时在下面或 [在 Twitter](https://twitter.com/IanMLewis) 上给我留言，让我知道容器运行时最难理解的是什么。

In the next post I'll do a deep dive into low-level container  runtimes. In that post I'll talk about exactly what low-level container  runtimes do. I'll talk about popular low-level runtimes like runc and  rkt, as well as unpopular-but-important ones like lmctfy. I'll even walk through how to implement a simple low-level runtime. Be sure to add my  RSS feed or [follow me on Twitter](https://twitter.com/IanMLewis) to get notified when the next blog post comes out.

在下一篇文章中，我将深入探讨低级容器运行时。在那篇文章中，我将确切地讨论低级容器运行时的作用。我将讨论流行的低级运行时，例如 runc 和 rkt，以及不受欢迎但很重要的运行时，例如 lmctfy。我什至会介绍如何实现一个简单的低级运行时。请务必添加我的 RSS 提要或 [在 Twitter 上关注我](https://twitter.com/IanMLewis)，以便在下一篇博文发布时收到通知。

> Update: Please continue on and check out [Container Runtimes Part 2: Anatomy of a Low-Level Container Runtime](https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level -contai)

> 更新：请继续阅读 [Container Runtimes Part 2: Anatomy of a Low-Level Container Runtime](https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level - 包含）

Until then, you can get more involved with the Kubernetes community via these channels:
 - Post and answer questions on [Stack Overflow](http://stackoverflow.com/questions/tagged/kubernetes)
 - Follow [@Kubernetesio](https://twitter.com/kubernetesio) on Twitter
 - Join the Kubernetes[ Slack](http://slack.k8s.io/) and chat with us. (I'm ianlewis so say Hi!)
 - Contribute to the Kubernetes project on[ GitHub](https://github.com/kubernetes/kubernetes) 
> Thanks to [Sandeep Dinesh](https://twitter.com/SandeepDinesh), [Mark Mandel](https://twitter.com/neurotic), [Craig Box](https://twitter.com/craigbox) , [Maya Kaczorowski](https://twitter.com/mayakaczorowski), and Joe Burnett for reviewing drafts of this post.

> 感谢 [Sandeep Dinesh](https://twitter.com/SandeepDinesh)、[Mark Mandel](https://twitter.com/neurotic)、[Craig Box](https://twitter.com/craigbox) , [Maya Kaczorowski](https://twitter.com/mayakaczorowski) 和 Joe Burnett 审阅了这篇文章的草稿。

------


