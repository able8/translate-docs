# Container Runtimes Part 3: High-Level Runtimes

# 容器运行时第 3 部分：高级运行时

Oct 30, 2018 From: https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes

This is the third part in a four-part series on container runtimes. It's been a while since [part 1](https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r), but in that post I gave an overview of container runtimes and discussed the differences between low-level and high-level runtimes. In [part 2](https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level-contai) I went into detail on low-level container runtimes and built a simple low-level runtime.

这是关于容器运行时的四部分系列中的第三部分。距离 [第 1 部分](https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r) 已经有一段时间了，但在那篇文章中，我概述了容器运行时并讨论了低级和高级运行时之间的差异。在 [第 2 部分](https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level-contai) 中，我详细介绍了低级容器运行时并构建了一个简单的低级容器级别运行时。*

High-level runtimes are higher up the stack than low-level runtimes. While low-level runtimes are responsible for the mechanics of actually  running a container, high-level runtimes are responsible for transport  and management of container images, unpacking the image, and passing off to the low-level runtime to run the container. Typically, high-level  runtimes provide a daemon application and an API that remote  applications can use to logically run containers and monitor them but  they sit on top of and delegate to low-level runtimes or other  high-level runtimes for the actual work.

高级运行时比低级运行时在堆栈上更高。虽然低级运行时负责实际运行容器的机制，但高级运行时负责容器镜像的传输和管理、解包镜像以及传递给低级运行时来运行容器。通常，高级运行时提供一个守护程序应用程序和一个 API，远程应用程序可以使用它们来逻辑运行容器并监视它们，但它们位于低级运行时或其他高级运行时之上并委托给低级运行时或其他高级运行时进行实际工作。

High-level runtimes can also provide features that sound low-level,  but are used across individual containers on a machine. For example, one feature might be the management of network namespaces, and allowing  containers to join another container's network namespace.

高级运行时还可以提供听起来低级的功能，但在机器上的各个容器中使用。例如，一个功能可能是网络命名空间的管理，并允许容器加入另一个容器的网络命名空间。

Here's a conceptual diagram to understand how the components fit together:

这是一个概念图，用于了解组件如何组合在一起：

![运行时架构图](https://storage.googleapis.com/static.ianlewis.org/prod/img/771/runtime-architecture.png)

# Examples of High-Level Runtimes

# 高级运行时示例

To better understand high-level runtimes, it’s helpful to look at a  few examples. Like low-level runtimes, each runtime implements different features.

为了更好地理解高级运行时，看几个例子会很有帮助。与低级运行时一样，每个运行时实现不同的功能。

## Docker

## 码头工人

Docker is one of the first open source container runtimes. It was  developed by the platform-as-a-service company dotCloud, and was used to run their users' web applications in containers.

Docker 是最早的开源容器运行时之一。它由平台即服务公司 dotCloud 开发，用于在容器中运行用户的 Web 应用程序。

Docker is a container runtime that incorporates building, packaging,  sharing, and running containers. Docker has a client/server architecture and was originally built as a monolithic daemon, `dockerd`, and the `docker` client application. The daemon provided most of the logic of building  containers, managing the images, and running containers, along with an  API. The command line client could be run to send commands and to get  information from the daemon.

Docker 是一个容器运行时，包括构建、打包、共享和运行容器。 Docker 有一个客户端/服务器架构，最初是作为一个单一的守护进程构建的，`dockerd` 和 `docker` 客户端应用程序。该守护进程提供了构建容器、管理图像和运行容器的大部分逻辑，以及一个 API。可以运行命令行客户端来发送命令并从守护程序获取信息。

It was the first popular runtime to incorporate all of the features  needed during the lifecycle of building and running containers.

它是第一个包含构建和运行容器生命周期所需的所有功能的流行运行时。

Docker originally implemented both high-level and low-level runtime  features, but those pieces have since been broken out into separate  projects as runc and containerd. Docker now consists of the `dockerd` daemon, and the `docker-containerd` daemon along with `docker-runc`. `docker-containerd` and `docker-runc` are just Docker packaged versions of vanilla `containerd` and `runc`.

Docker 最初实现了高级和低级运行时功能，但后来这些部分被分解为单独的项目，如 runc 和 containerd。 Docker 现在由`dockerd` 守护进程、`docker-containerd` 守护进程和`docker-runc` 组成。 `docker-containerd` 和 `docker-runc` 只是 vanilla `containerd` 和 `runc` 的 Docker 打包版本。

![Docker架构图](https://storage.googleapis.com/static.ianlewis.org/prod/img/771/docker.png)

`dockerd` provides features such as building images, and dockerd uses `docker-containerd` to provide features such as image management and running containers. For instance, Docker's build step is actually just some logic that  interprets a Dockerfile, runs the necessary commands in a container  using `containerd`, and saves the resulting container file system as an image.

`dockerd` 提供构建镜像等功能，dockerd 使用 `docker-containerd` 提供镜像管理、运行容器等功能。例如，Docker 的构建步骤实际上只是一些解释 Dockerfile 的逻辑，使用 `containerd` 在容器中运行必要的命令，并将生成的容器文件系统保存为映像。

## containerd

[containerd](https://containerd.io/) is a high-level  runtime that was split off from Docker. Like runc, which was broken off  as the low-level runtime piece, containerd was broken off as the  high-level runtime piece of Docker. `containerd` implements  downloading images, managing them, and running containers from images. When it needs to run a container it unpacks the image into an OCI  runtime bundle and shells out to `runc` to run it.

[containerd](https://containerd.io/) 是一个从 Docker 中分离出来的高级运行时。就像 runc 是被拆分出的低级运行时部分一样，containerd 是被拆分出的 Docker 的高级运行时部分。 `containerd` 实现了下载镜像、管理镜像以及从镜像运行容器。当它需要运行一个容器时，它会将镜像解包到一个 OCI 运行时包中，并通过“runc”来运行它。

Containerd also provides an API and client application that can be  used to interact with it. The containerd command line client is `ctr`.`ctr` can be used to tell `containerd` to pull a container image:

Containerd 还提供了可用于与其交互的 API 和客户端应用程序。 containerd 命令行客户端是 `ctr`。

ctr` 可用于告诉 `containerd` 拉取容器镜像：

```
 $ sudo ctr images pull docker.io/library/redis:latest
```


List the images you have:

列出您拥有的图像：

```
 $ sudo ctr images list
```


Run a container based on an image:

运行基于镜像的容器：

```
 $ sudo ctr container create docker.io/library/redis:latest redis
```


List the running containers:

列出正在运行的容器：

```
 $ sudo ctr container list
```


Stop the container:

停止容器：

```
 $ sudo ctr container delete redis
```


These commands are similar to how a user interacts with Docker. However, in contrast with Docker, containerd is focused solely on  running containers, so it does not provide a mechanism for building  containers. Docker was focused on end-user and developer use cases,  whereas containerd is focused on operational use cases, such as running  containers on servers. Tasks such as building container images are left  to other tools.

这些命令类似于用户与 Docker 交互的方式。但是，与 Docker 相比，containerd 只专注于运行容器，因此它没有提供构建容器的机制。 Docker 专注于最终用户和开发人员用例，而 containerd 专注于操作用例，例如在服务器上运行容器。诸如构建容器镜像之类的任务留给了其他工具。

## rkt

In the previous post, I mentioned that `rkt` is a runtime  that has both low-level and high-level features. For instance, much like Docker, rkt allows you to build container images, fetch and manage  container images in a local repository, and run them all from a single  command. `rkt` stops short of Docker's functionality, however, in that it doesn't provide a long-running daemon and remote API.

在上一篇文章中，我提到 rkt 是一个同时具有低级和高级功能的运行时。例如，与 Docker 非常相似，rkt 允许您构建容器镜像，在本地存储库中获取和管理容器镜像，并通过单个命令运行它们。然而，`rkt` 没有提供 Docker 的功能，因为它不提供长期运行的守护进程和远程 API。

You can fetch remote images: 您可以获取远程图像：

```
 $ sudo rkt fetch coreos.com/etcd:v3.3.10
```


You can then list the images installed locally: 然后，您可以列出本地安装的映像：

```
 $ sudo rkt image list
 ID                      NAME                                    SIZE    IMPORT TIME     LAST USED
 sha512-07738c36c639     coreos.com/rkt/stage1-fly:1.30.0        44MiB   2 minutes ago   2 minutes ago
 sha512-51ea8f513d06     coreos.com/oem-gce:1855.5.0             591MiB  2 minutes ago   2 minutes ago
 sha512-2ba519594e47     coreos.com/etcd:v3.3.10                 69MiB   25 seconds ago  24 seconds ago
```


And delete images: 并删除图像：

```
 $ sudo rkt image rm coreos.com/etcd:v3.3.10
 successfully removed aci for image: "sha512-2ba519594e4783330ae14e7691caabfb839b5f57c0384310a7ad5fa2966d85e3"
 rm: 1 image(s) successfully removed
```


Though rkt doesn't seem to be actively developed very much anymore it is an interesting tool and an important part of the history of  container technology.

虽然 rkt 似乎不再被积极开发，但它是一个有趣的工具，也是容器技术历史的重要组成部分。

## Onward, Upward

## 向前，向上

In the next post I'll move up the stack and talk about runtimes from  the perspective of Kubernetes and how they work. Be sure to add [my RSS feed](https://www.ianlewis.org/feed/enfeed) or follow me on Twitter to get notified when the next blog post comes out. 

在下一篇文章中，我将向上移动堆栈并从 Kubernetes 的角度讨论运行时及其工作原理。请务必添加 [我的 RSS 提要](https://www.ianlewis.org/feed/enfeed) 或在 Twitter 上关注我，以便在下一篇博文发布时收到通知。

