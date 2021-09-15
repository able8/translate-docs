# Why and How to Use containerd from the Command Line

# 为什么以及如何从命令行使用 containerd

September 12, 2021

[containerd](https://github.com/containerd/containerd) is a [high-level container runtime](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-runtimes), _aka_ [container manager](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-management). To put it simply, it's a daemon that manages the complete container lifecycle on a single host: creates, starts, stops containers, pulls and stores images, configures mounts, networking, etc.

[containerd](https://github.com/containerd/containerd) 是一个[高级容器运行时](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-Beyond/#container-runtimes)、_aka_ [容器管理器](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-management)。简单来说，它是一个守护进程，在单个主机上管理完整的容器生命周期：创建、启动、停止容器、拉取和存储镜像、配置挂载、网络等。

containerd is designed to be easily embeddable into larger systems. [Docker uses containerd under the hood](https://www.docker.com/blog/what-is-containerd-runtime/) to run containers. [Kubernetes can use containerd via CRI](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#containerd) to manage containers on a single node. But smaller projects also can benefit from the ease of integrating with containerd - for instance, [faasd](https://github.com/openfaas/faasd) uses containerd (we need more **d**'s!) to spin up a full-fledged [Function-as-a-Service](https://en.wikipedia.org/wiki/Function_as_a_service) solution on a standalone server.

containerd 旨在轻松嵌入到更大的系统中。 [Docker 在幕后使用 containerd](https://www.docker.com/blog/what-is-containerd-runtime/) 来运行容器。 [Kubernetes 可以通过 CRI 使用 containerd](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#containerd) 在单个节点上管理容器。但是较小的项目也可以从与 containerd 集成的便利中受益——例如，[faasd](https://github.com/openfaas/faasd) 使用 containerd（我们需要更多 **d**'s！)来旋转在独立服务器上创建成熟的 [Function-as-a-Service](https://en.wikipedia.org/wiki/Function_as_a_service) 解决方案。

![Docker and Kubernetes use containerd](http://iximiuz.com/containerd-command-line-clients/docker-and-kubernetes-use-containerd-2000-opt.png)

However, using containerd programmatically is not the only option. It also can be used from the command line via one of the available clients. The resulting container UX may not be as comprehensive and user-friendly as the one provided by the `docker` client, but it still can be useful, for instance, for debugging or learning purposes.

但是，以编程方式使用 containerd 并不是唯一的选择。它还可以通过可用客户端之一从命令行使用。由此产生的容器 UX 可能不像 `docker` 客户端提供的那样全面和用户友好，但它仍然是有用的，例如，用于调试或学习目的。

![containerd command-line clients (ctr, nerdctl, crictl)](http://iximiuz.com/containerd-command-line-clients/containerd-command-line-clients-2000-opt.png)

## How to use containerd with ctr

## 如何使用containerd和ctr

[ctr](https://github.com/containerd/containerd/tree/e1ad7791077916aac9c1f4981ad350f0e3fce719/cmd/ctr) is a command-line client shipped as part of the containerd project. If you have containerd running on a machine, chances are the `ctr` binary is also there.

[ctr](https://github.com/containerd/containerd/tree/e1ad7791077916aac9c1f4981ad350f0e3fce719/cmd/ctr) 是作为 containerd 项目的一部分提供的命令行客户端。如果你在一台机器上运行了 containerd，那么 `ctr` 二进制文件很可能也在那里。

The `ctr` interface is [obviously] incompatible with Docker CLI and, at first sight, may look not so user-friendly. Apparently, its primary audience is containerd developers testing the daemon. However, since it's the closest thing to the actual containerd API, it can serve as a great exploration means - by examining the available commands, you can get a rough idea of what containerd can and cannot do.

`ctr` 界面 [显然] 与 Docker CLI 不兼容，乍一看，可能看起来不太用户友好。显然，它的主要受众是测试守护进程的容器开发人员。但是，由于它最接近实际的 containerd API，因此它可以作为一种很好的探索手段——通过检查可用命令，您可以大致了解 containerd 可以做什么和不能做什么。

`ctr` is also well-suitable for learning the capabilities of [low-level [OCI] container runtimes](http://iximiuz.com/en/posts/oci-containers/) since `ctr + containerd` is [much closer to actual containers](http://iximiuz.com/en/posts/implementing-container-runtime-shim/) than `docker + dockerd`.

`ctr` 也非常适合学习[低级 [OCI] 容器运行时](http://iximiuz.com/en/posts/oci-containers/) 的功能，因为 `ctr + containerd` [很多更接近实际容器](http://iximiuz.com/en/posts/implementing-container-runtime-shim/) 比`docker + dockerd`。

### Working with container images using ctr

### 使用 ctr 处理容器图像

When **pulling images**, the fully-qualified reference seems to be required, so you cannot omit the registry or the tag part:

**拉图片**时，好像需要全限定引用，所以不能省略registry或tag部分：

```bash
$ ctr images pull docker.io/library/nginx:1.21
$ ctr images pull docker.io/kennethreitz/httpbin:latest
$ ctr images pull docker.io/kennethreitz/httpbin:latest
$ ctr images pull quay.io/quay/redis:latest

```

To **list local images**, one can use:

要**列出本地图像**，可以使用：

```bash
$ ctr images ls

```

Surprisingly, containerd doesn't provide out-of-the-box image building support. However, containerd itself is often used to build images by higher-level tools.

令人惊讶的是，containerd 不提供开箱即用的镜像构建支持。但是，containerd 本身经常被更高级别的工具用来构建镜像。

Check out my investigation post on [what actually happens when you build an image](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/) to learn more about image building internals.

查看我关于 [构建映像时实际发生的情况](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/) 的调查帖子，了解有关映像构建的更多信息内件。

Instead of building images with `ctr`, you can **import existing images** built with `docker build` or other OCI-compatible software:

您可以**导入使用 docker build 或其他 OCI 兼容软件构建的现有图像**，而不是使用 `ctr` 构建图像：

```bash
$ docker build -t my-app .
$ docker save -o my-app.tar my-app

$ ctr images import my-app.tar

```

With `ctr`, you can also **mount images** for future exploration:

使用`ctr`，您还可以**挂载图像**以供将来探索：

```bash
$ mkdir /tmp/httpbin
$ ctr images mount docker.io/kennethreitz/httpbin:latest /tmp/httpbin

$ ls -l /tmp/httpbin/
total 80
drwxr-xr-x 2 root root 4096 Oct 18  2018 bin
drwxr-xr-x 2 root root 4096 Apr 24  2018 boot
drwxr-xr-x 4 root root 4096 Oct 18  2018 dev
drwxr-xr-x 1 root root 4096 Oct 24  2018 etc
drwxr-xr-x 2 root root 4096 Apr 24  2018 home
drwxr-xr-x 3 root root 4096 Oct 24  2018 httpbin
...

$ ctr images unmount /tmp/httpbin

```

To **remove images** using `ctr`, run:

要使用 `ctr` **删除图像**，请运行：

```bash
$ ctr images remove docker.io/library/nginx:1.21

```

### Working with containers using ctr

### 使用 ctr 处理容器

Having a local image, you can **run a container** with `ctr run <image-ref> <container-id>`. For instance:

拥有本地镜像，您可以使用 `ctr run <image-ref> <container-id>` **运行容器**。例如：

```bash
$ ctr run --rm -t docker.io/library/debian:latest cont1

```

Notice that unlike user-friendly `docker run` generating a unique container ID for you, with `ctr`, you must supply the unique container ID yourself. The `ctr run` command also supports only some of the familiar `docker run` flags: `--env`, `-t,--tty`, `-d,--detach`, `--rm`, etc . But no port publishing or automatic container restart with `--restart=always` out of the box.

请注意，与用户友好的 `docker run` 为您生成唯一的容器 ID 不同，使用 `ctr`，您必须自己提供唯一的容器 ID。 `ctr run` 命令也只支持一些熟悉的 `docker run` 标志：`--env`、`-t、--tty`、`-d、--detach`、`--rm` 等. 但是开箱即用的`--restart=always` 没有端口发布或容器自动重启。

Similarly to images, you can **list existing containers** with:

与图像类似，您可以**列出现有容器**：

```bash
$ ctr containers ls

```

Interesting that the `ctr run` command is actually a shortcut for `ctr container create` \+ `ctr task start`:

有趣的是，`ctr run` 命令实际上是`ctr container create` \+ `ctr task start` 的快捷方式：

```bash
$ ctr container create -t docker.io/library/nginx:latest nginx_1
$ ctr container ls
CONTAINER    IMAGE                              RUNTIME
nginx_1      docker.io/library/nginx:latest     io.containerd.runc.v2

$ ctr task ls
TASK    PID    STATUS        # Empty!

$ ctr task start -d nginx_1  # -d for --detach
$ ctr task list
TASK     PID      STATUS
nginx_1  10074    RUNNING

```

I like this separation of `container` and `task` subcommands because it reflects the often forgotten nature of OCI containers. Despite the common belief, containers aren't processes - [_containers are isolated and restricted execution environments_](http://iximiuz.com/en/posts/oci-containers/) for processes.

我喜欢这种`container` 和`task` 子命令的分离，因为它反映了OCI 容器经常被遗忘的特性。尽管普遍认为容器不是进程 - [_containers 是隔离且受限的执行环境_](http://iximiuz.com/en/posts/oci-containers/) 用于进程。

With `ctr task attach`, you can **reconnect to the stdio streams** of an existing task running inside of a container:

使用 `ctr task attach`，您可以**重新连接到容器内运行的现有任务的 stdio 流**：

```bash
$ ctr task attach nginx_1
2021/09/12 15:42:20 [notice] 1#1: using the "epoll" event method
2021/09/12 15:42:20 [notice] 1#1: nginx/1.21.3
2021/09/12 15:42:20 [notice] 1#1: built by gcc 8.3.0 (Debian 8.3.0-6)
2021/09/12 15:42:20 [notice] 1#1: OS: Linux 4.19.0-17-amd64
2021/09/12 15:42:20 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1024:1024
2021/09/12 15:42:20 [notice] 1#1: start worker processes
2021/09/12 15:42:20 [notice] 1#1: start worker process 31
...

```

Much like with `docker`, you can **execute a task in an existing container**:

与 `docker` 非常相似，您可以**在现有容器中执行任务**：

```bash
$ ctr task exec -t --exec-id bash_1 nginx_1 bash

# From inside the container:
$ root@host:/# curl 127.0.0.1:80
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
...

```

Before removing a container, all its tasks must be **stopped**:

在移除容器之前，必须**停止**其所有任务：

```bash
$ ctr task kill -9 nginx_1

```

Alternatively, you can **remove running tasks** using the `--force` flag:

或者，您可以使用 `--force` 标志**删除正在运行的任务**：

```bash
$ ctr task rm -f nginx_1

```

Finally, to **remove the container**, run:

最后，要**删除容器**，请运行：

```bash
$ ctr container rm nginx_1

```

## How to use containerd with nerdctl

## 如何在 nerdctl 中使用 containerd

[nerdctl](https://github.com/containerd/nerdctl) is a relatively new command-line client for containerd. Unlike `ctr`, `nerdctl` aims to be user-friendly and Docker-compatible. To some extent, `nerdctl + containerd` can seamlessly replace `docker + dockerd`. However, [this does not seem to be the goal of the project](https://medium.com/nttlabs/nerdctl-359311b32d0e): 

[nerdctl](https://github.com/containerd/nerdctl) 是一个相对较新的 containerd 命令行客户端。与 `ctr` 不同，`nerdctl` 旨在用户友好且与 Docker 兼容。在某种程度上，`nerdctl + containerd` 可以无缝替代`docker + dockerd`。然而，[这似乎不是项目的目标](https://medium.com/nttlabs/nerdctl-359311b32d0e)：

> The goal of `nerdctl` is to facilitate experimenting the cutting-edge features of containerd that are not present in Docker. Such features include, but not limited to, lazy-pulling (stargz) and encryption of images (ocicrypt). These features are expected to be eventually available in Docker as well, however, it is likely to take several months, or perhaps years, as Docker is currently designed to use only a small portion of the containerd subsystems. Refactoring Docker to use the entire containerd would be possible, but not straightforward. So we [ [NTT](https://www.rd.ntt/e/)] decided to create a new CLI that fully uses containerd, but we do not intend to complete with Docker. We have been contributing to Docker/Moby as well as containerd, and will continue to do so.

> `nerdctl` 的目标是促进试验 Docker 中不存在的 containerd 的尖端功能。此类功能包括但不限于延迟拉取 (stargz) 和图像加密 (ocicrypt)。这些功能预计最终也将在 Docker 中可用，但是，这可能需要几个月甚至几年的时间，因为 Docker 目前被设计为仅使用容器子系统的一小部分。重构 Docker 以使用整个 containerd 是可能的，但并不简单。所以我们 [[NTT](https://www.rd.ntt/e/)] 决定创建一个完全使用 containerd 的新 CLI，但我们不打算用 Docker 来完成。我们一直在为 Docker/Moby 以及 containerd 做出贡献，并将继续这样做。

From the basic usage standpoint, comparing to `ctr`, `nerdctl` supports:

从基本使用的角度来看，与 `ctr` 相比，`nerdctl` 支持：

- Image building with`nerdctl build`
- Container networking management
- Docker Compose with`nerdctl compose up`

- 使用`nerdctl build`构建图像
- 容器网络管理
- Docker Compose 与`nerdctl compose up`

And the coolest part about it is that `nerdctl` tries to provide the identical to `docker` (and `podman`) command-line UX. So, **if you are familiar with `docker` (or `podman`) CLI, you are already familiar with `nerdctl`.**

最酷的部分是`nerdctl` 试图提供与`docker`（和`podman`）相同的命令行用户体验。所以，**如果你熟悉 `docker`（或 `podman`）CLI，你就已经熟悉了 `nerdctl`。**

## How to use containerd with crictl

## 如何在 crictl 中使用 containerd

[crictl](https://github.com/kubernetes-sigs/cri-tools) is a command-line client for [[Kubernetes] CRI-compatible container runtimes](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/).

[crictl](https://github.com/kubernetes-sigs/cri-tools) 是 [[Kubernetes] CRI 兼容容器运行时的命令行客户端](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/)。

Click here to learn more about Kubernetes Container Runtime Interface (CRI).

单击此处了解有关 Kubernetes 容器运行时接口 (CRI) 的更多信息。

**[Kubernetes Container Runtime Interface (CRI)](https://github.com/kubernetes/cri-api/)** was introduced to make Kubernetes container runtime-agnostic. The Kubernetes node agent, _kubelet_, implementing the CRI client API, can use any container runtime implementing the CRI server API to manage containers and pods on its node.

**[Kubernetes 容器运行时接口 (CRI)](https://github.com/kubernetes/cri-api/)** 被引入以使 Kubernetes 容器运行时不可知。 Kubernetes 节点代理 _kubelet_ 实现了 CRI 客户端 API，可以使用任何实现 CRI 服务器 API 的容器运行时来管理其节点上的容器和 pod。

![Kubernetes CRI](http://iximiuz.com/containerd-command-line-clients/cri-2000-opt.png)

_Kubernetes CRI._



Since version 1.1, containerd comes with a built-in CRI plugin. Hence, containerd is a CRI-compatible container runtime. Therefore, it can be used with `crictl`.

从 1.1 版开始，containerd 带有一个内置的 CRI 插件。因此，containerd 是一个兼容 CRI 的容器运行时。因此，它可以与 `crictl` 一起使用。

`crictl` was created to inspect and debug container runtimes and applications on a Kubernetes node. [It supports the following operations](https://github.com/kubernetes-sigs/cri-tools/blob/98f3364b4b684966b27bf372412d805d7dcbcb10/docs/crictl.md):

`crictl` 被创建来检查和调试 Kubernetes 节点上的容器运行时和应用程序。  [支持以下操作](https://github.com/kubernetes-sigs/cri-tools/blob/98f3364b4b684966b27bf372412d805d7dcbcb10/docs/crictl.md)：

```bash
attach: Attach to a running container
create: Create a new container
exec: Run a command in a running container
version: Display runtime version information
images, image, img: List images
inspect: Display the status of one or more containers
inspecti: Return the status of one or more images
imagefsinfo: Return image filesystem info
inspectp: Display the status of one or more pods
logs: Fetch the logs of a container
port-forward: Forward local port to a pod
ps: List containers
pull: Pull an image from a registry
run: Run a new container inside a sandbox
runp: Run a new pod
rm: Remove one or more containers
rmi: Remove one or more images
rmp: Remove one or more pods
pods: List pods
start: Start one or more created containers
info: Display information of the container runtime
stop: Stop one or more running containers
stopp: Stop one or more running pods
update: Update one or more running containers
config: Get and set crictl client configuration options
stats: List container(s) resource usage statistics

```

The interesting part here is that with `crictl + containerd` bundle, one can learn how pods are actually implemented. [But this topic deserves its own blog post](http://iximiuz.com/en/newsletter/) 😉

这里有趣的部分是，通过 `crictl + containerd` 包，你可以了解 pod 是如何实际实现的。 [但这个话题值得拥有自己的博文](http://iximiuz.com/en/newsletter/) 😉

For more information on how to use `crictl` with containerd, check out [this document (part of the containerd project)](https://github.com/containerd/cri/blob/68b61297b59e38c1088db10fbd19807a4ffbad87/docs/crictl.md).

有关如何在 containerd 中使用 `crictl` 的更多信息，请查看 [本文档（containerd 项目的一部分）](https://github.com/containerd/cri/blob/68b61297b59e38c1088db10fbd19807a4ffbad87/docs/crictl.md)。

### More [Containers](http://iximiuz.com/en/categories/?category=Containers) Posts From This Blog

### 更多 [Containers](http://iximiuz.com/en/categories/?category=Containers) 来自此博客的帖子

- [Journey From Containerization to Orchestration and Beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)
- [Containers aren't Linux processes](http://iximiuz.com/en/posts/oci-containers/)
- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/) 

- [从容器化到编排及超越的旅程](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)
- [容器不是 Linux 进程](http://iximiuz.com/en/posts/oci-containers/)
- [并非每个容器内部都有操作系统](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [你不需要图像来运行容器](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)

- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [conman - [the] container manager: inception](http://iximiuz.com/en/posts/conman-the-container-manager-inception/)
- [Implementing Container Runtime Shim: runc](http://iximiuz.com/en/posts/implementing-container-runtime-shim/)
- [Implementing Container Runtime Shim: First Code](http://iximiuz.com/en/posts/implementing-container-runtime-shim-2/)
- [Implementing Container Runtime Shim: Interactive Containers](http://iximiuz.com/en/posts/implementing-container-runtime-shim-3/)

- [你需要容器来构建镜像](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [conman - [the] 容器管理器：inception](http://iximiuz.com/en/posts/conman-the-container-manager-inception/)
- [实现容器运行时 Shim: runc](http://iximiuz.com/en/posts/implementing-container-runtime-shim/)
- [实现容器运行时 Shim：第一个代码](http://iximiuz.com/en/posts/implementing-container-runtime-shim-2/)
- [实现容器运行时 Shim：交互式容器](http://iximiuz.com/en/posts/implementing-container-runtime-shim-3/)

[containerd,](javascript: void 0) [ctr,](javascript: void 0) [crictl,](javascript: void 0) [nerdctl](javascript: void 0)

[containerd,](javascript: void 0) [ctr,](javascript: void 0) [crictl,](javascript: void 0) [nerdctl](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

