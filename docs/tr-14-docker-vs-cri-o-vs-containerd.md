# Docker vs CRI-O vs Containerd

January 8, 2021 From: https://computingforgeeks.com/docker-vs-cri-o-vs-containerd/

## Introduction

The sphere of containers is like a labyrinthine forest cover. The many  branching tunnels and jargon on top of jargon it is characterized with  can sooner or later lead you to a familiar destination that we have all  been to. The chilly destination of confusion. This post will do its best to try to clear the thicket and let in some sunshine to light a few  paths here and there so that you can continue with your container  journey with some smile added on your face. We are going to look at the  differences that exist among Docker, CRI-O, and containerd.  After doing a bit of reading, here is some information about Docker, CRI-O, and containerd.

容器的球体就像迷宫般的森林覆盖物。许多分支隧道和行话加上行话的特点，迟早会带你到一个我们都去过的熟悉的目的地。混乱的寒冷目的地。这篇文章将尽最大努力清理灌木丛，让一些阳光照亮这里和那里的几条小路，这样你就可以继续你的集装箱之旅，脸上带着微笑。我们将看看 Docker、CRI-O 和 containerd 之间存在的差异。
在阅读了一些之后，这里有一些关于 Docker、CRI-O 和 containerd 的信息。

## Docker

Before version 1.11, the implementation of Docker was a monolithic daemon. The monolith did everything as one package such as downloading container  images, launching container processes, exposing a remote API, and acting as a log collection daemon, all in a centralized process running as  root (Source:[ Coreos](https://coreos .com/rkt/docs/latest/rkt-vs-other-projects.html#rkt-vs-docker)).

在 1.11 版本之前，Docker 的实现是一个单体守护进程。单体作为一个包完成了所有事情，例如下载容器镜像、启动容器进程、公开远程 API 以及充当日志收集守护进程，所有这些都在一个以 root 身份运行的集中进程中。

Such a centralized architecture has some benefits when it comes to  deployment but unearths other fundamental problems. An example is that  it does not follow best practices for Unix process and privilege  separation. Moreover, the monolithic implementation makes Docker  difficult to properly integrate with Linux init systems such as upstart  and systemd (Source: [Coreos](https://coreos.com/rkt/docs/latest/rkt-vs-other-projects.html #rkt-vs-docker)).

这种集中式架构在部署方面有一些好处，但会发现其他基本问题。一个例子是它没有遵循 Unix 进程和权限分离的最佳实践。此外，单体实现使得 Docker 难以与 upstart 和 systemd 等 Linux init 系统正确集成（来源：[Coreos](https://coreos.com/rkt/docs/latest/rkt-vs-other-projects.html #rkt-vs-docker))。

This led to the splitting of Docker into different parts as quoted in the below opening paragraph when Docker 1.11 was launched.

这导致在 Docker 1.11 发布时将 Docker 拆分为不同的部分，如下面的开篇段落所引用。

“*We are excited to introduce Docker Engine 1.11, our first release built on runC ™ and containerd ™. With this release, Docker is the first to ship a runtime based on OCI technology, demonstrating the progress the team  has made since donating our industry-standard container format and  runtime under the Linux Foundation in June of 2015*” (Source [Docker]( https://www.docker.com/blog/docker-engine-1-11-runc/)).

“*我们很高兴推出 Docker Engine 1.11，这是我们基于 runC™ 和 containerd™ 构建的第一个版本。在这个版本中，Docker 是第一个发布基于 OCI 技术的运行时，展示了自 2015 年 6 月在 Linux 基金会下捐赠我们的行业标准容器格式和运行时以来团队取得的进步*”（来源 [Docker]（ https://www.docker.com/blog/docker-engine-1-11-runc/))。

According to Docker, splitting it up into focused independent tools mean more  focused maintainers and ultimately better quality software. You can get more information and details about [OCI technology here](https://www.opencontainers.org/). The figure below illustrates the new architecture of Docker 1.11 after building it on **runC ™** and **containerd ™**.

根据 Docker 的说法，将其拆分为专注的独立工具意味着更专注的维护人员和最终质量更高的软件。您可以在此处获取有关 [OCI 技术](https://www.opencontainers.org/) 的更多信息和详细信息。下图展示了 Docker 1.11 在 **runC™** 和 **containerd™** 上构建后的新架构。

![Docker1.11](https://computingforgeeks.com/wp-content/uploads/2019/12/Docker1.11-1024x590.png?ezimgfmt=rs:696x401/rscb23/ng:webp/ngcb23)

Since then, containerd now handles the execution of containers which was  previously done by docker daemon itself. This is the exact flow:

A user runs commands from Docker-CLI **>** Docker-CLI talks to the Docker daemon(dockerd) **>** Docker daemon(dockerd) listens for requests and manages the lifecycle of the container via containerd which it contacts **>** containerd takes the request and starts a container through runC and does all the container life-cylces within the host.

从那时起，containerd 现在处理容器的执行，以前由 docker daemon 本身完成。这是确切的流程：

用户从 Docker-CLI 运行命令 **>** Docker-CLI 与 Docker 守护进程（dockerd）对话 **>** Docker 守护进程（dockerd）监听请求并通过它联系的 containerd 管理容器的生命周期 **>** containerd 接受请求并通过 runC 启动一个容器，并在主机内执行所有容器生命周期。

Note: ***runc,*** in brief, is a CLI tool for spawning and running containers according to the OCI specification.

注意：***runc,*** 简而言之，是一个 CLI 工具，用于根据 OCI 规范生成和运行容器。

## Containerd

Now let us shift our focus and get to understand what containerd is all  about. From a high level stand point, containerd is a daemon that  controls **runC**. From [containerd website](https://containerd.io/), “*containerd manages the complete container lifecycle of its host system, from image transfer and storage to container execution and supervision to  low-level storage to network attachments and beyond*". It is also known as a container engine. 

现在让我们转移注意力并了解 containerd 的全部内容。从高层次的角度来看，containerd 是一个控制 **runC** 的守护进程。来自 [containerd 网站](https://containerd.io/)，“*containerd 管理其主机系统的完整容器生命周期，从图像传输和存储到容器执行和监督到低级存储到网络附件等等* “。它也被称为容器引擎。

containerd helps abstract away syscalls or Operating-System specific functionality to run containers on Linux, Windows or any other Operating System. It  provides a client layer that any other platform such as Docker or  Kubernetes can build on top of without ever caring to sneak into the  kernel level details. It should be noted that in Kubernetes containerd  can be used as CRI runtime. We will tackle CRI in a jiffy. These are  what you get by leveraging on containerd:

- You get push and pull functionality
 - Image management APIs to create, execute, and manage containers and their tasks
 - Snapshot management.
 - You get all of that without ever scratching your head with the underlying OS details (Source: [Docker](https://www.docker.com/blog/what-is-containerd-runtime/))

containerd 有助于抽象出系统调用或操作系统特定的功能，以在 Linux、Windows 或任何其他操作系统上运行容器。它提供了一个客户端层，任何其他平台（如 Docker 或 Kubernetes）都可以在其上构建，而无需关心内核级别的细节。需要注意的是，在 Kubernetes 中，containerd 可以用作 CRI 运行时。我们将很快解决 CRI。这些是您利用 containerd 获得的结果：

- 您可以获得推拉功能
- 用于创建、执行和管理容器及其任务的图像管理 API
- 快照管理。
- 您无需担心底层操作系统的详细信息即可获得所有这些（来源：[Docker](https://www.docker.com/blog/what-is-containerd-runtime/)）

## CRI-O

And now onto CRI-O. Before we delve into CRI-O, lets immerse our heads briefly into the pool of CRI, that is **Container Runtime Interface**. CRI is a plugin interface that gives kubelet the ability to use  different OCI-compliant container runtimes (such as runc – the most  widely used but there are others such as crun, railcar, and  katacontainers), without needing to recompile Kubernetes. Kubelet, as  you know, is a cluster node agent used to create pods and start  containers (Source: [RedHat](https://www.redhat.com/en/blog/introducing-cri-o-10)).

现在到 CRI-O。在我们深入研究 CRI-O 之前，让我们先简单地沉浸在 CRI 池中，即 **Container Runtime Interface**。 CRI 是一个插件接口，它使 kubelet 能够使用不同的 OCI 兼容容器运行时（例如 runc - 使用最广泛，但还有其他的，例如 crun、railcar 和 katacontainers），而无需重新编译 Kubernetes。如您所知，Kubelet 是一个集群节点代理，用于创建 Pod 和启动容器（来源：[RedHat](https://www.redhat.com/en/blog/introducing-cri-o-10)）。

To understand the need for CRI, it will be wise to know the pain points  that Kubernetes experienced prior to it. Kubernetes was previously bound to a specific container runtime which spawned a lot of maintenance  overhead for the upstream Kubernetes community. Moreover, vendors  building solutions on top of the Kubernetes experienced the same  overhead. This necessitated the development of CRI to make Kubernetes  container runtime-agnostic by decoupling it from various runtimes.

要了解对 CRI 的需求，明智的做法是了解 Kubernetes 在此之前经历的痛点。 Kubernetes 以前绑定到特定的容器运行时，这为上游 Kubernetes 社区带来了大量维护开销。此外，在 Kubernetes 之上构建解决方案的供应商也经历了相同的开销。这需要开发 CRI，通过将其与各种运行时解耦来使 Kubernetes 容器与运行时无关。

Since the plugin is already built, CRI-O project was begun to provide a  lightweight runtime specifically for Kubernetes. CRI-O makes it possible for Kubernetes to run containers directly without much tools and code.

由于插件已经构建好了，CRI-O 项目开始专门为 Kubernetes 提供一个轻量级的运行时。 CRI-O 使 Kubernetes 无需太多工具和代码即可直接运行容器。

### Components of CRI-O

#### The following are the components of CRI-O

- OCI compatible runtime – Default is **runC**, other OCI compliant are supported as well e.g Kata Containers.
 - containers/storage – Library used for managing layers and creating root file-systems for the containers in a pod.
 - containers/image – Library is used for pulling images from registries.
 - networking (CNI) – Used for setting up networking for the pods. Flannel, Weave and OpenShift-SDN CNI plugins have been tested.
 - container monitoring (conmon) – Utility within CRI-O that is used to monitor the containers.
 - security is provided by several core Linux capabilities

#### 以下是CRI-O的组成部分

- OCI 兼容运行时 - 默认为 **runC**，其他 OCI 兼容也受支持，例如 Kata Containers。
- 容器/存储 - 用于管理层和为 pod 中的容器创建根文件系统的库。
- 容器/图像 - 库用于从注册表中提取图像。
- 网络 (CNI) – 用于为 Pod 设置网络。 Flannel、Weave 和 OpenShift-SDN CNI 插件已经过测试。
- 容器监控 (conmon) – CRI-O 中用于监控容器的实用程序。
- 安全性由几个核心 Linux 功能提供

Source: [CRI-O](https://cri-o.io/)

The screenshot below illustrates the whole Kubernetes and CRI-O process. From it, it can also be observed that CRI-O lies under the category of a container engine.

下面的截图说明了整个 Kubernetes 和 CRI-O 过程。从中也可以看出，CRI-O 属于容器引擎的范畴。

![Kubernetes](https://computingforgeeks.com/wp-content/uploads/2019/12/Kubernetes-1024x620.png?ezimgfmt=rs:696x421/rscb23/ng:webp/ngcb23)

## Conclusion

## 结论

Even though there are many products in the container sphere, a closer look  at most of the solutions provided efforts to fix an issue here or there. Docker, CRI-O, and containerd all have their own spaces and can all  benefit Kubernetes in launching and maintaining pods. What can be  observed is that the three depend on **runC** at the lowest level to handle the running of containers. We hope the post was informative as beneficial as you had wished.

尽管容器领域有很多产品，但仔细研究大多数解决方案，可以解决这里或那里的问题。 Docker、CRI-O 和 containerd 都有自己的空间，都可以在启动和维护 Pod 时使 Kubernetes 受益。可以观察到的是，三者在最低层依赖**runC**来处理容器的运行。我们希望这篇文章内容丰富，如您所愿。

Before you leave, take a look at related articles and guides below:

在您离开之前，请查看以下相关文章和指南：

[Easily Manage Multiple Kubernetes Clusters with kubectl & kubectx](https://computingforgeeks.com/manage-multiple-kubernetes-clusters-with-kubectl-kubectx/)

[使用 kubectl & kubectx 轻松管理多个 Kubernetes 集群](https://computingforgeeks.com/manage-multiple-kubernetes-clusters-with-kubectl-kubectx/)

[How To Deploy Lightweight Kubernetes Cluster in 5 minutes with K3s](https://computingforgeeks.com/how-to-deploy-lightweight-kubernetes-cluster-with-k3s/)

### Your support is our everlasting motivation, that cup of coffee is what keeps us going! 
### 您的支持是我们永恒的动力，那杯咖啡是我们前进的动力！

