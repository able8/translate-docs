# Container Runtimes Part 2: Anatomy of a Low-Level Container Runtime

# 容器运行时第 2 部分：底层容器运行时剖析

*Feb 26, 2018* From https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level-contai

This is the second in a four-part series on container runtimes. In [part 1](https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r), I gave an overview of container runtimes and discussed the differences  between low-level and high -level runtimes. In this post I will go into  detail on low-level container runtimes.

这是关于容器运行时的四部分系列中的第二部分。在[第 1 部分](https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r) 中，我概述了容器运行时并讨论了低级和高级之间的区别级运行时。在这篇文章中，我将详细介绍低级容器运行时。

Low-level runtimes have a limited feature set and typically perform  the low-level tasks for running a container. Most developers shouldn't  use them for their day-to-day work. Low-level runtimes are usually  implemented as simple tools or libraries that developers of higher level runtimes and tools can use for the low-level features. While most  developers won't use low-level runtimes directly, it's good to know how  they work for troubleshooting and debugging purposes.

低级运行时具有有限的功能集，通常执行用于运行容器的低级任务。大多数开发人员不应该在日常工作中使用它们。低级运行时通常作为简单的工具或库实现，高级运行时和工具的开发人员可以将其用于低级功能。虽然大多数开发人员不会直接使用低级运行时，但最好了解它们如何用于故障排除和调试目的。

As I explained in part 1, containers are implemented using [Linux namespaces](https://en.wikipedia.org/wiki/Linux_namespaces) and [cgroups](https://en.wikipedia.org/wiki/Cgroups). Namespaces let you virtualize system resources, like the file system or networking for each container. On the other hand, cgroups provide a way to limit the amount of resources, such as CPU and memory, that each  container can use. At their core, low-level container runtimes are  responsible for setting up these namespaces and cgroups for containers,  and then running commands inside those namespaces and cgroups. Most  container runtimes implement more features, but those are the essential  bits.

正如我在第 1 部分中所解释的，容器是使用 [Linux 命名空间](https://en.wikipedia.org/wiki/Linux_namespaces) 和 [cgroups](https://en.wikipedia.org/wiki/Cgroups) 实现的。命名空间让您可以虚拟化系统资源，例如每个容器的文件系统或网络。另一方面，cgroup 提供了一种方法来限制每个容器可以使用的资源量，例如 CPU 和内存。在其核心，低级容器运行时负责为容器设置这些命名空间和 cgroup，然后在这些命名空间和 cgroup 中运行命令。大多数容器运行时实现了更多功能，但这些是必不可少的部分。

Be sure to check out the amazing talk ["Building a container from scratch in Go"](https://www.youtube.com/watch?v=Utf-A4rODH8) by Liz Rice. Her talk is a great introduction to how low-level  container runtimes are implemented. Liz goes through many of these  steps, but the most trivial runtime you can imagine that you could still call a "container runtime" might do something like the following:
 - Create cgroup
 - Run command(s) in cgroup
 - [Unshare](http://man7.org/linux/man-pages/man2/unshare.2.html) to move to its own namespaces
 - Clean up cgroup after command completes (namespaces are deleted automatically when not referenced by a running process)

请务必查看 Liz Rice 的精彩演讲 [“在 Go 中从头开始构建容器”](https://www.youtube.com/watch?v=Utf-A4rODH8)。她的演讲很好地介绍了如何实现低级容器运行时。 Liz 经历了其中的许多步骤，但是您可以想象的最简单的运行时仍然可以调用“容器运行时”可能会执行以下操作：
- 创建 cgroup
- 在 cgroup 中运行命令
- [Unshare](http://man7.org/linux/man-pages/man2/unshare.2.html) 移动到自己的命名空间
- 命令完成后清理 cgroup（命名空间在没有被正在运行的进程引用时自动删除）

A robust low-level container runtime, however, would do a lot more,  like allow for setting resource limits on the cgroup, setting up a root  filesystem, and chrooting the container's process to the root file  system.

然而，一个强大的低级容器运行时会做更多的事情，比如允许在 cgroup 上设置资源限制、设置根文件系统以及将容器的进程 chroot 到根文件系统。

## Building a Sample Runtime

## 构建示例运行时

Let's walk through running a simple ad hoc runtime to set up a container. We can perform the steps using the standard Linux [cgcreate](https://linux.die.net/man/1/cgcreate), [cgset](https://linux.die.net/man/1/cgset) , [cgexec](https://linux.die.net/man/1/cgexec), [chroot](http://man7.org/linux/man-pages/man2/chroot.2.html) and [ unshare](http://man7.org/linux/man-pages/man1/unshare.1.html) commands. You'll need to run most of the commands below as root.

让我们逐步运行一个简单的临时运行时来设置容器。我们可以使用标准 Linux [cgcreate](https://linux.die.net/man/1/cgcreate), [cgset](https://linux.die.net/man/1/cgset) 执行这些步骤, [cgexec](https://linux.die.net/man/1/cgexec), [chroot](http://man7.org/linux/man-pages/man2/chroot.2.html) 和 [ unshare](http://man7.org/linux/man-pages/man1/unshare.1.html) 命令。您需要以 root 身份运行以下大部分命令。

First let's set up a root filesystem for our container. We'll use the busybox Docker container as our base. Here we create a temporary  directory and extract busybox into it. Most of these commands need to be run as root.

首先让我们为我们的容器设置一个根文件系统。我们将使用 busybox Docker 容器作为我们的基础。这里我们创建一个临时目录并将busybox解压到其中。大多数这些命令需要以 root 身份运行。

```
 $ CID=$(docker create busybox)
 $ ROOTFS=$(mktemp -d)
 $ docker export $CID | tar -xf - -C $ROOTFS
```


Now let's create our cgroup and set restrictions on the memory and  CPU. Memory limits are set in bytes. Here we are setting the limit to  100MB.

现在让我们创建我们的 cgroup 并设置内存和 CPU 的限制。内存限制以字节为单位设置。在这里，我们将限制设置为 100MB。

```
 $ UUID=$(uuidgen)
 $ cgcreate -g cpu,memory:$UUID
 $ cgset -r memory.limit_in_bytes=100000000 $UUID
 $ cgset -r cpu.shares=512 $UUID
```


CPU usage can be restricted in one of two ways. Here we set our CPU  limit using CPU "shares". Shares are an amount of CPU time relative to  other processes running at the same time. Containers running by  themselves can use the whole CPU, but if other containers are running,  they can use a proportional amount of CPU to their CPU shares.

可以通过以下两种方式之一限制 CPU 使用率。在这里，我们使用 CPU“份额”设置我们的 CPU 限制。份额是相对于同时运行的其他进程的 CPU 时间量。自己运行的容器可以使用整个 CPU，但如果其他容器正在运行，它们可以使用与其 CPU 份额成比例的 CPU。

CPU limits based on CPU cores are a bit more complicated. They let  you set hard limits on the amount of CPU cores that a container can use. Limiting CPU cores requires you set two options on the cgroup, `cfs_period_us` and `cfs_quota_us`. `cfs_period_us` specifies how often CPU usage is checked and `cfs_quota_us` specifies the amount of time that a task can run on one core in one period. Both are specified in microseconds. 

基于 CPU 内核的 CPU 限制要复杂一些。它们让您可以对容器可以使用的 CPU 内核数量设置硬限制。限制 CPU 内核需要您在 cgroup 上设置两个选项，`cfs_period_us` 和 `cfs_quota_us`。 `cfs_period_us` 指定检查 CPU 使用率的频率，`cfs_quota_us` 指定任务可以在一个周期内在一个内核上运行的时间量。两者都以微秒为单位指定。

For instance, if we wanted to limit our container to two cores we  could specify a period of one second and a quota of two seconds (one  second is 1,000,000 microseconds) and this would effectively allow our  process to use two cores during a one-second period. [This article](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/resource_management_guide/sec-cpu) explains this concept in depth.

例如，如果我们想将容器限制为两个内核，我们可以指定一秒的时间段和两秒的配额（一秒是 1,000,000 微秒），这将有效地允许我们的进程在一秒内使用两个内核时期。[This article](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/resource_management_guide/sec-cpu) 深入解释了这个概念。

```
 $ cgset -r cpu.cfs_period_us=1000000 $UUID
 $ cgset -r cpu.cfs_quota_us=2000000 $UUID
```


Next we can execute a command in the container. This will execute the command within the cgroup we created, unshare the specified namespaces, set the hostname, and chroot to our filesystem.

接下来我们可以在容器中执行一个命令。这将在我们创建的 cgroup 中执行命令，取消共享指定的命名空间，设置主机名和 chroot 到我们的文件系统。

```
 $ cgexec -g cpu,memory:$UUID \
 >     unshare -uinpUrf --mount-proc \
 >     sh -c "/bin/hostname $UUID && chroot $ROOTFS /bin/sh"
 / # echo "Hello from in a container"
 Hello from in a container
 / # exit
```


Finally, after our command has completed, we can clean up by deleting the cgroup and temporary directory that we created.

最后，在我们的命令完成后，我们可以通过删除我们创建的 cgroup 和临时目录来进行清理。

```
 $ cgdelete -r -g cpu,memory:$UUID
 $ rm -r $ROOTFS
```


To further demonstrate how this works, I wrote a simple runtime in bash called [execc](https://github.com/ianlewis/execc). It supports mount, user, pid, ipc, uts, and network namespaces; setting memory limits; setting CPU limits by number of cores; mounting the proc file system; and running the container in its own root file system.

为了进一步演示这是如何工作的，我在 bash 中编写了一个简单的运行时，名为 [execc](https://github.com/ianlewis/execc)。支持mount、user、pid、ipc、uts、网络命名空间；设置内存限制；按内核数设置 CPU 限制；挂载proc文件系统；并在其自己的根文件系统中运行容器。

## Examples of Low-Level Container Runtimes

## 低级容器运行时示例

In order to better understand low-level container runtimes it's  useful to look at some examples. These runtimes implement different  features and emphasize different aspects of containerization.

为了更好地理解低级容器运行时，查看一些示例很有用。这些运行时实现了不同的功能并强调容器化的不同方面。

### lmctfy

Though not in wide use, one container runtime of note is [lmctfy](https://github.com/google/lmctfy). lmctfy is a project by Google, based on the internal container runtime that [Borg](https://research.google.com/pubs/pub43438.html) uses. One of its most interesting features is that it supports  container hierarchies that use cgroup hierarchies via the container  names. For example, a root container called "busybox" could create  sub-containers under the name "busybox/sub1" or "busybox/sub2" where the names form a kind of path structure. As a result each sub-container can have its own cgroups that are then limited by the parent container's  cgroup. This is inspired by Borg and gives containers in lmctfy the  ability to run sub-task containers under a pre-allocated set of  resources on a server, and thus achieve more stringent SLOs than could  be provided by the runtime itself.

虽然没有被广泛使用，但一个值得注意的容器运行时是 [lmctfy](https://github.com/google/lmctfy)。 lmctfy 是 Google 的一个项目，基于 [Borg](https://research.google.com/pubs/pub43438.html) 使用的内部容器运行时。它最有趣的特性之一是它支持通过容器名称使用 cgroup 层次结构的容器层次结构。例如，名为“busybox”的根容器可以在名称“busybox/sub1”或“busybox/sub2”下创建子容器，其中名称形成一种路径结构。因此，每个子容器都可以拥有自己的 cgroup，然后受父容器的 cgroup 限制。这受到 Borg 的启发，使 lmctfy 中的容器能够在服务器上预先分配的一组资源下运行子任务容器，从而实现比运行时本身所能提供的更严格的 SLO。

While lmctfy provides some interesting features and ideas, other  runtimes were more usable so Google decided it would be better for the  community to focus worked on Docker's libcontainer instead of lmctfy.

虽然 lmctfy 提供了一些有趣的功能和想法，但其他运行时更有用，因此 Google 决定社区最好专注于 Docker 的 libcontainer 而不是 lmctfy。

### runc

runc is currently the most widely used container runtime. It was  originally developed as part of Docker and was later extracted out as a  separate tool and library.

runc 是目前使用最广泛的容器运行时。它最初是作为 Docker 的一部分开发的，后来被提取出来作为一个单独的工具和库。

Internally, runc runs containers similarly to how I described it  above, but runc implements the OCI runtime spec. That means that it runs containers from a specific "OCI bundle" format. The format of the  bundle has a config.json file for some configuration and a root file  system for the container. You can find out more by reading the [OCI runtime spec](https://github.com/opencontainers/runtime-spec) on GitHub. You can learn how to install runc from the [runc GitHub project](https://github.com/opencontainers/runc).

在内部，runc 运行容器的方式与我上面描述的类似，但 runc 实现了 OCI 运行时规范。这意味着它从特定的“OCI 包”格式运行容器。捆绑包的格式有一个用于某些配置的 config.json 文件和用于容器的根文件系统。您可以通过阅读 GitHub 上的 [OCI 运行时规范](https://github.com/opencontainers/runtime-spec) 了解更多信息。您可以从 [runc GitHub 项目](https://github.com/opencontainers/runc) 了解如何安装 runc。

First create the root filesystem. Here we'll use busybox again.

首先创建根文件系统。这里我们将再次使用busybox。

```
 $ mkdir rootfs
 $ docker export $(docker create busybox) | tar -xf - -C rootfs
```


Next create a config.json file.

接下来创建一个 config.json 文件。

```
 $ runc spec
```


This command creates a template config.json for our container. It should look something like this:

此命令为我们的容器创建模板 config.json。它应该是这样的：

```
 $ cat config.json
 {
         "ociVersion": "1.0.0",
         "process": {
                 "terminal": true,
                 "user": {
                         "uid": 0,
                         "gid": 0
                 },
                 "args": [
                         "sh"
                 ],
                 "env": [
                         "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                         "TERM=xterm"
                 ],
                 "cwd": "/",
                 "capabilities": {
 ...
```


By default it runs the sh command in a container with a root  filesystem at ./rootfs. Since that's exactly the setup we want we can  just go ahead and run the container.

默认情况下，它在一个根文件系统位于 ./rootfs 的容器中运行 sh 命令。由于这正是我们想要的设置，我们可以继续运行容器。

```
 $ sudo runc run mycontainerid
 / # echo "Hello from in a container"
 Hello from in a container
```


## rkt 
rkt is a popular alternative to Docker/runc that was developed by  CoreOS. rkt is a bit hard to categorize because it provides all the  features that other low-level runtimes like runc provide, but also  provides features typical of high-level runtimes. Here I'll describe the low-level features of rkt and leave the high-level features for the  next post.

rkt 是 CoreOS 开发的 Docker/runc 的流行替代品。 rkt 有点难以分类，因为它提供了 runc 等其他低级运行时提供的所有功能，但也提供了高级运行时的典型功能。在这里我将描述rkt的低级特性，而将高级特性留给下一篇文章。

rkt originally used the [Application Container](https://coreos.com/rkt/docs/latest/app-container.html) (appc) standard, which was developed as an open alternative standard to alternative to Docker's container format. Appc never achieved  widespread adoption as a container format and appc is no longer being  actively developed bit achieved its goals to ensure open standards are  available to the community. Instead of appc, rkt will use OCI container  formats in the future.

rkt 最初使用 [Application Container](https://coreos.com/rkt/docs/latest/app-container.html) (appc) 标准，该标准被开发为一种开放的替代标准，以替代 Docker 的容器格式。 Appc 从未作为一种容器格式被广泛采用，并且 appc 也不再被积极开发，它的目标已经实现，以确保社区可以使用开放标准。 rkt 将来会使用 OCI 容器格式，而不是 appc。

Application Container Image (ACI) is the image format for Appc. Images are a tar.gz containing a manifest directory and a rootfs  directory for the root filesystem. You can read more about ACI [here](https://github.com/appc/spec/blob/master/spec/aci.md).

应用程序容器映像 (ACI) 是 Appc 的映像格式。镜像是包含清单目录和根文件系统的 rootfs 目录的 tar.gz。您可以在 [此处](https://github.com/appc/spec/blob/master/spec/aci.md) 阅读有关 ACI 的更多信息。

You can build a container image using the `acbuild` tool. You can use acbuild in shell scripts that can be executed much like Dockerfiles are run.

您可以使用 `acbuild` 工具构建容器映像。您可以在 shell 脚本中使用 acbuild，这些脚本可以像运行 Dockerfile 一样执行。

```
 acbuild begin
 acbuild set-name example.com/hello
 acbuild dep add quay.io/coreos/alpine-sh
 acbuild copy hello /bin/hello
 acbuild set-exec /bin/hello
 acbuild port add www tcp 5000
 acbuild label add version 0.0.1
 acbuild label add arch amd64
 acbuild label add os linux
 acbuild annotation add authors "Carly Container <carly@example.com>"
 acbuild write hello-0.0.1-linux-amd64.aci
 acbuild end
```


## Adiós!

## 再见！

I hope this helped you get an idea of what low-level container  runtimes are. While most users of containers will use higher-level  runtimes, it's good to know how containers are working underneath the  covers for troubleshooting issues and debugging.

我希望这能帮助您了解什么是低级容器运行时。虽然容器的大多数用户将使用更高级别的运行时，但了解容器如何在幕后工作以解决问题和调试是很好的。

In the next post I'll move up the stack and talk about high-level  container runtimes. I'll talk about what high-level runtimes provide and how they are much better for app developers who want to use containers. I'll also talk about popular runtimes like Docker and rkt's high-level  features. Be sure to add my RSS feed or [follow me on Twitter](https://twitter.com/IanMLewis) to get notified when the next blog post comes out.

在下一篇文章中，我将向上移动堆栈并讨论高级容器运行时。我将讨论高级运行时提供了什么，以及它们如何为想要使用容器的应用程序开发人员提供更好的服务。我还将讨论流行的运行时，如 Docker 和 rkt 的高级功能。请务必添加我的 RSS 提要或 [在 Twitter 上关注我](https://twitter.com/IanMLewis)，以便在下一篇博文发布时收到通知。

> Update: Please continue on and check out [Container Runtimes Part 3: High-Level Runtimes](https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes)

> 更新：请继续阅读 [Container Runtimes Part 3: High-Level Runtimes](https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes)

Until then, you can get more involved with the Kubernetes community via these channels:
 - Post and answer questions on [Stack Overflow](http://stackoverflow.com/questions/tagged/kubernetes)
 - Follow [@Kubernetesio](https://twitter.com/kubernetesio) on Twitter
 - Join the Kubernetes[ Slack](http://slack.k8s.io/) and chat with us. (I'm ianlewis so say Hi!)
 - Contribute to the Kubernetes project on[ GitHub](https://github.com/kubernetes/kubernetes)

在此之前，您可以通过以下渠道更多地参与 Kubernetes 社区：
- 在 [Stack Overflow](http://stackoverflow.com/questions/tagged/kubernetes) 上发布和回答问题
- 在 Twitter 上关注 [@Kubernetesio](https://twitter.com/kubernetesio)
- 加入 Kubernetes[Slack](http://slack.k8s.io/) 和我们聊天。 （我是 ianlewis 所以打个招呼吧！）
- 为 [GitHub](https://github.com/kubernetes/kubernetes) 上的 Kubernetes 项目做出贡献

> Thanks to [Craig Box](https://twitter.com/craigbox), Jack Wilbur, Philip Mallory, [David Gageot](https://twitter.com/dgageot), Jonathan MacMillan, and [Maya Kaczorowski]( https://twitter.com/MayaKaczorowski) for reviewing drafts of this post.*

> 感谢 [Craig Box](https://twitter.com/craigbox)、Jack Wilbur、Philip Mallory、[David Gageot](https://twitter.com/dgageot)、Jonathan MacMillan 和 [Maya Kaczorowski]( https://twitter.com/MayaKaczorowski) 用于审阅这篇文章的草稿。*

------


