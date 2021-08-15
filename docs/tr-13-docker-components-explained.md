# Docker components explained

# Docker 组件解释

April 27, 2018 · 5 min · Alexander Holbreich From: https://alexander.holbreich.org/docker-components-explained/

![img](https://alexander.holbreich.org/images/headers/container-ship.jpg)

It’s all started with a pressure of splitting the monolithic implementation of Docker and  [Moby Project](https://alexander.holbreich.org/docker-moby/) as result. Now Docker consist of several components on every particular machine. Confusion happens when people are talking about these  different components of the [Docker](https://alexander.holbreich.org/tag/docker). So let’s improve the situation…

这一切都始于拆分 Docker 和 [Moby 项目](https://alexander.holbreich.org/docker-moby/) 的单体实现的压力。现在 Docker 由每台特定机器上的几个组件组成。当人们谈论 [Docker](https://alexander.holbreich.org/tag/docker) 的这些不同组件时，会发生混淆。所以让我们改善这种情况......

## Docker CLI (docker)

```
 /usr/bin/docker
```


Docker is used as a reference to the whole set of docker tools and at the beginning, it was a monolith. But now `docker-cli` is only responsible for user-friendly communication with docker.

Docker 被用作整个 Docker 工具集的参考，一开始，它是一个整体。但是现在`docker-cli`只负责与docker的用户友好通信。

So the command’s like `docker build ...` `docker run ...` are handled by Docker CLI and result in the invocation of **dockerd** API.

所以像 `docker build ...` `docker run ...` 这样的命令由 Docker CLI 处理并导致调用 **dockerd** API。

## Dockerd

```
 /usr/bin/dockerd
```


The Docker daemon - **dockerd** listens for Docker API requests and manages the host’s Container life-cycles by utilizing **contanerd**

Docker 守护进程 - **dockerd** 侦听 Docker API 请求并利用 **contanerd** 管理主机的容器生命周期

`dockerd` can listen for Docker Engine API requests via  three different types of Socket: unix, tcp, and fd. By default, a unix  domain socket is created at `/var/run/docker.sock`, requiring either root permission, or docker group membership. On [Systemd](https://alexander.holbreich.org/tag/systemd/) based systems, you can communicate with the daemon via Systemd socket activation, use `dockerd -H fd://`.

`dockerd` 可以通过三种不同类型的 Socket 侦听 Docker Engine API 请求：unix、tcp 和 fd。默认情况下，unix 域套接字在 `/var/run/docker.sock` 创建，需要 root 权限或 docker 组成员身份。在基于 [Systemd](https://alexander.holbreich.org/tag/systemd/) 的系统上，您可以通过 Systemd 套接字激活与守护进程通信，使用 `dockerd -H fd://`。

There are many [configuration options](https://docs.docker.com/engine/reference/commandline/dockerd/) for the daemon, which are worth checking if you work with docker (dockerd). My impression is that `dockerd` is here to serve all the  features of Docker (or Docker EE) platform, while actual container  life-cycle management is “outsourced” to **containerd**.

守护进程有很多[配置选项](https://docs.docker.com/engine/reference/commandline/dockerd/)，如果你使用docker(dockerd)，值得检查一下。我的印象是，`dockerd` 是为了服务 Docker（或 Docker EE）平台的所有功能，而实际的容器生命周期管理“外包”给 **containerd**。

## Containerd

```
 /usr/bin/docker-containerd
```

`containerd` was introduced in Docker 1.11 and since then took the main responsibility of managing containers life-cycle. `containerd` is the *executor for containers*, but has a wider scope than *just executing* containers. So it also takes care of:

- Image push and pull
 - Managing storage
 - Of course executing of Containers by calling **runc** with the right parameters to run containers…
 - Managing of network primitives for interfaces
 - Management of network namespaces containers to join existing namespaces

`containerd` 是在 Docker 1.11 中引入的，从那时起主要负责管理容器生命周期。 `containerd` 是容器的*执行器*，但比*仅执行*容器的范围更广。因此，它还负责：

- 图像推拉
- 管理存储
- 当然，通过使用正确的参数调用 **runc** 来运行容器来执行容器......
- 管理接口的网络原语
- 管理网络命名空间容器以加入现有命名空间

*`containerd` fully leverages the \*OCI runtime specification\*[1](https://alexander.holbreich.org/docker-components-explained/#fn:1), image format specifications, and OCI reference implementation (runc )*. Because of its massive adoption, containerd is the industry standard  for implementing OCI. It is currently available for Linux and Windows.

*`containerd` 充分利用了 \*OCI 运行时规范\*[1](https://alexander.holbreich.org/docker-components-explained/#fn:1)、图像格式规范和 OCI 参考实现 (runc )*。由于其广泛采用，containerd 是实施 OCI 的行业标准。它目前可用于 Linux 和 Windows。

![img](https://github.com/containerd/containerd/raw/master/design/architecture.png)

As shown in the picture above, `contained` includes a  daemon exposing gRPC API over a local UNIX socket. The API is a  low-level one designed for higher layers to wrap and extend. `Containerd` uses `RunC` to run containers according to the OCI specification.

如上图所示，`contained` 包括一个通过本地 UNIX 套接字公开 gRPC API 的守护进程。 API 是一种低级 API，专为更高层的包装和扩展而设计。 `Containerd` 使用 `RunC` 根据 OCI 规范运行容器。

`containerd` is based on the Docker Engine’s core container runtime to benefit from its maturity and existing contributors, however, `containerd` is designed to be embedded into a larger system, rather than being used directly by developers or end-users.

`containerd` 基于 Docker 引擎的核心容器运行时，以从其成熟度和现有贡献者中受益，然而，`containerd` 旨在嵌入到更大的系统中，而不是由开发人员或最终用户直接使用。

Well, now other vendors can use containers without having to deal with docker-related parts. let’s go through some subsystems of `containerd`…

好吧，现在其他供应商可以使用容器而无需处理与 docker 相关的部分。让我们来看看`containerd`的一些子系统......

### RunC

`/usr/bin/docker-runc` runc (OCI runtime) can be seen as component of containerd.

`/usr/bin/docker-runc` runc（OCI运行时）可以看作是containerd的组件。

`runc` is a command-line client for running applications packaged according to the OCI format and is a compliant implementation of the OCI spec. Containers are configured using bundles. A bundle for a container is a directory that includes a specification file named “config.json” and a root filesystem. The root filesystem contains the contents of the container. Assuming you have an OCI bundle you can execute the container:

`runc` 是一个命令行客户端，用于运行根据 OCI 格式打包的应用程序，并且是 OCI 规范的兼容实现。容器是使用捆绑包配置的。容器的包是一个包含名为“config.json”的规范文件和根文件系统的目录。根文件系统包含容器的内容。假设你有一个 OCI 包，你可以执行容器

```bash
 # run as root
 cd /mycontainer
 runc run mycontainerid
```


### containerd-ctr

`/usr/bin/docker-containerd-ctr` (docker-)containerd-ctr  it’s barebone CLI (ctr) designed specifically  for development and debugging purposes for direct communication with `containerd`. It’s included in the releases of `containerd`. By that less interesting for docker users.

`/usr/bin/docker-containerd-ctr` (docker-)containerd-ctr 它是准系统 CLI (ctr)，专为开发和调试目的而设计，用于与 `containerd` 直接通信。它包含在 `containerd` 的版本中。对于 docker 用户来说，这不太有趣。

### containerd-shim

```
 /usr/bin/docker-containerd-shim
```

*The shim allows for daemon-less containers*. [According to Michael Crosby](https://groups.google.com/forum/#!topic/docker-dev/zaZFlvIx1_k) it basically sits as the parent of the container’s process to facilitate a few things.

- First it allows the runtimes, i.e. `runc`, to exit after it starts the container. This way we don’t have to have the long-running runtime processes for containers.
 - Second it keeps the STDIO and other fds open for the container in the case `containerd` and/or docker both die. If the shim was not running then the parent  site of the pipes or the TTY master would be closed and the container  would exit.
 - Finally it allows the container’s exit status to be reported back to a higher level tool like docker without having the be the actual parent of the container’s process and do a wait.

*垫片允许无守护进程的容器*。 [根据 Michael Crosby 的说法](https://groups.google.com/forum/#!topic/docker-dev/zaZFlvIx1_k) 它基本上作为容器进程的父进程来促进一些事情。

- 首先它允许运行时，即`runc`，在它启动容器后退出。这样我们就不必为容器提供长时间运行的运行时进程。
- 其次，在 `containerd` 和/或 docker 都死掉的情况下，它会保持 STDIO 和其他 fds 为容器打开。如果 shim 没有运行，则管道的父站点或 TTY 主站将关闭，容器将退出。
- 最后，它允许将容器的退出状态报告回更高级的工具，如 docker，而无需成为容器进程的实际父进程并执行 wait。

## How it all works together

## 这一切是如何协同工作的

We can do an experiment. First we check what *Docker processes* are running right after Docker installation.

我们可以做一个实验。首先，我们检查安装 Docker 后正在运行的 *Docker 进程 *。

```bash
 ps fxa | grep docker -A 3
 # prints:
 2239 ? Ssl    0:27 /usr/bin/dockerd -H fd://
 2397 ? Ssl    0:18  \_ docker-containerd -l unix:///var/run/docker/libcontainerd/docker-containerd.sock ...
 ...
```


well at this point we see that dockerd is started and containerd is running as a child process too. Like described, `dockerd` needs `containerd` ;) Now let’s run one simple container, that executes for a minute and then exits.

好吧，此时我们看到 dockerd 已启动，并且 containerd 也作为子进程运行。如上所述，`dockerd` 需要 `containerd` ;) 现在让我们运行一个简单的容器，它执行一分钟然后退出。

```
 docker run -d alpine sleep 60
```


Now we should see it in the process list in the next 60 seconds. Let’s check again:

现在我们应该在接下来的 60 秒内在进程列表中看到它。我们再检查一下：

```bash
 ps fxa | grep dockerd -A 3
 #prints
  2239 ? Ssl    0:28 /usr/bin/dockerd -H fd://
  2397 ? Ssl    0:19  \_ docker-containerd -l unix:///var/run/docker/libcontainerd/docker-containerd.sock ...
 15476 ? Sl     0:00      \_ docker-containerd-shim 3da7... /var/run/docker/libcontainerd/3da7.. docker-runc
 15494 ? Ss     0:00          \_ sleep 60
```


Now we see the whole process chain:

现在我们看到整个流程链：

**dockerd** –> **containerd** –> **containerd-shim** –> “sleep 60” (desired process in the container).

We do not see `runc` in the chain, we know `containerd-shim` takes over after `runc` has started the container. Also, we know that theoretically `containerd-shim` can survive the crash of `containerd`. But in the current docker version, it’s [not activated by default](https://docs.docker.com/config/containers/live-restore/#enable-live-restore).

我们没有在链中看到 `runc`，我们知道 `containerd-shim` 在 `runc` 启动容器后接管。此外，我们知道理论上 `containerd-shim` 可以在 `containerd` 崩溃后幸存下来。但是在当前的 docker 版本中，它[默认未激活](https://docs.docker.com/config/containers/live-restore/#enable-live-restore)。

However, it’s a pretty long chain with possible disadvantages that such chains might have.

然而，这是一个相当长的链，可能有这样的链可能具有的缺点。

### How it all works in Kubernetes

### 在 Kubernetes 中是如何工作的

You might imagine that [Kubernetes](https://alexander.holbreich.org/tag/k8n/) do not need Docker-specific parts. As of now, it’s exactly the case…

您可能会认为 [Kubernetes](https://alexander.holbreich.org/tag/k8n/) 不需要 Docker 特定的部分。截至目前，情况确实如此……

![img](https://github.com/containerd/cri/raw/master/docs/cri.png)

Kubernetes “speaks” with `contanerd` directly as depicted in the picture. If interested, check [how it was in between](https://kubernetes.io/blog/2017/11/containerd-container-runtime-options-kubernetes).

Kubernetes 直接与“contanerd”“对话”，如图所示。如果有兴趣，请查看 [中间情况](https://kubernetes.io/blog/2017/11/containerd-container-runtime-options-kubernetes)。

I hope this might help all Docker users. Give me a hint if something is not precise.

我希望这可以帮助所有 Docker 用户。如果有什么不准确的，请给我一个提示。

------

1. The [OCI Runtime Specification](https://github.com/opencontainers/runtime-spec) outlines how to run a containers “filesystem bundle” that is unpacked  on disk. At a high level, an OCI implementation would download an OCI  Image ([OCI Image Specification](https://github.com/opencontainers/image-spec)) then unpack that image into an OCI Runtime filesystem bundle. At this  point, the OCI Runtime Bundle would be run by an OCI Runtime. [↩︎](https://alexander.holbreich.org/docker-components-explained/#fnref:1)

1. [OCI 运行时规范](https://github.com/opencontainers/runtime-spec) 概述了如何运行在磁盘上解压缩的容器“文件系统包”。在高层次上，OCI 实现将下载 OCI 图像（[OCI 图像规范]（https://github.com/opencontainers/image-spec）），然后将该图像解压到 OCI 运行时文件系统包中。此时，OCI Runtime Bundle 将由 OCI Runtime 运行。 [↩︎](https://alexander.holbreich.org/docker-components-explained/#fnref:1)


