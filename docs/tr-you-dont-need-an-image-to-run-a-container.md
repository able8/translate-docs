# How to Run a Container Without an Image

# 如何在没有镜像的情况下运行容器

May 16, 2020 (Updated: August 7, 2021)

[Containers,](http://iximiuz.com/en/categories/?category=Containers) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

As we already know, containers are just [isolated and restricted Linux processes](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/#container-is-just-a-processes). We also learned that it's fairly simple to [create a container with a single executable file inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/#container-without-distro-inside) starting from [_scratch_](https://hub.docker.com/_/scratch) image (ie without putting a full Linux distribution in there). This time we will go even further and demonstrate that containers don't require [images](https://github.com/opencontainers/image-spec) at all. And after that, we will try to justify the actual need for images and their place in the [_containerverse_](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/) .

正如我们已经知道的，容器只是[隔离和受限的 Linux 进程](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/#container-is-just -a-进程）。我们还了解到[创建一个包含单个可执行文件的容器]相当简单(http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/#container -without-distro-inside) 从 [_scratch_](https://hub.docker.com/_/scratch) 镜像开始（即没有在其中放置完整的 Linux 发行版)。这次我们将更进一步，证明容器根本不需要 [images](https://github.com/opencontainers/image-spec)。之后，我们将尝试证明对图像的实际需求及其在[_containerverse_](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/) 中的位置.

### How to Run Containers without Docker and... Images

### 如何在没有 Docker 和...图像的情况下运行容器

You may have heard that Docker uses a tool called [_runc_](https://github.com/opencontainers/runc) to run containers. Well, to be more accurate, Docker depends on a lower-level piece of software called [_containerd_](https://github.com/containerd/containerd) which in turn relies on a [standardized container runtime](https://github.com/opencontainers/runtime-spec) implementation. And in the wild, most of the time _runc_ plays the role of such a component.

您可能听说过 Docker 使用一个名为 [_runc_](https://github.com/opencontainers/runc) 的工具来运行容器。好吧，更准确地说，Docker 依赖于一个名为 [_containerd_](https://github.com/containerd/containerd) 的低级软件，而后者又依赖于 [标准化容器运行时](https://github.com/opencontainers/runtime-spec) 实现。在野外，大多数时候 _runc_ 扮演这样一个组件的角色。

Let's take a closer look at [_runc_](https://github.com/opencontainers/runc). Surprisingly or not, it's **a command-line tool** for running containers. Following the instructions from its README file, we can come up with the sequence of commands needed to create and then start a container.

让我们仔细看看 [_runc_](https://github.com/opencontainers/runc)。不管是否令人惊讶，它是用于运行容器的**命令行工具**。按照 README 文件中的说明，我们可以提出创建和启动容器所需的命令序列。

First, we need to create a container configuration file:

首先，我们需要创建一个容器配置文件：

```bash
$ mkdir my_container && cd my_container

$ runc spec

$ file config.json
> config.json: ASCII text

```

The content of `config.json` on my machine looks as follows:

我机器上`config.json`的内容如下：

```json
{
    "ociVersion": "1.0.1-dev",
    "process": {
        "terminal": true,
        "user": {
            "uid": 0,
            "gid": 0
        },
        "args": [
            "sh"
        ],
        "cwd": "/",
        "env": [ ... ],
        "capabilities": { ... },
        "rlimits": [ ... ]
    },
    "root": {
        "path": "rootfs",
        "readonly": true
    },
    "hostname": "runc",
    "mounts": [ ... ],
    "linux": {
        "namespaces": [
            { "type": "pid" },
            { "type": "network" },
            { "type": "ipc" },
            { "type": "uts" },
            { "type": "mount" }
        ]
    }
}

```

It's not so hard to notice, that this file describes the program to be launched inside of the container ( `process.args[0]`), its environment ( `env`, `capabilities`, `root` filesystem, etc), and the container itself ( `mounts`, `namespaces`, etc).

不难注意到，这个文件描述了要在容器内启动的程序（`process.args[0]`）、它的环境（`env`、`capabilities`、`root` 文件系统等），以及容器本身（`mounts`、`namespaces` 等）。

Now, we need to prepare a root filesystem for our container. To keep it simple, we will just start with an empty folder and put [a single (statically-linked) executable](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/#container-without-distro-inside) in there:

现在，我们需要为我们的容器准备一个根文件系统。为简单起见，我们将从一个空文件夹开始，然后放置 [一个（静态链接的）可执行文件](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/#container-without-distro-inside)在那里：

```bash
$ cat <<EOF > main.go
package main

import "fmt"
import "os"

func main() {
    fmt.Println(os.Hostname())
}
EOF

$ GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o printme
$ file printme
> printme: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, stripped

$ mkdir rootfs
$ mv printme rootfs/

```

We also need to modify the `config.json` accordingly by setting `process.args[0]` to `/printme`, disabling terminal support, and changing the hostname to something more pleasing to the eye.

我们还需要相应地修改 `config.json`，将 `process.args[0]` 设置为 `/printme`，禁用终端支持，并将主机名更改为更美观的内容。

Finally, we can create the container using the adjusted `config.json` and start it:

最后，我们可以使用调整后的 `config.json` 创建容器并启动它：

```bash
$ sudo runc create mycont1
$ sudo runc start mycont1
> HAL-9000 <nil>

```

That was it! We ran a process inside of a container and from its point of view, the hostname was `HAL-9000` while running `/usr/bin/hostname` on this machine outside of a container gives `localhost.localdomain`. Pretty simple, right?

就是这样！我们在容器内运行了一个进程，从它的角度来看，主机名是“HAL-9000”，而在容器外的这台机器上运行“/usr/bin/hostname”时，主机名是“localhost.localdomain”。很简单，对吧？

_Click here to see how to run a container with a full Linux distro inside._ 

_单击此处查看如何运行包含完整 Linux 发行版的容器。_

The procedure is exactly the same as above. The only difference is that we need to have a root filesystem of the desired Linux distribution beforehand.

过程与上述完全相同。唯一的区别是我们需要事先拥有所需 Linux 发行版的根文件系统。

```bash
$ mkdir my_container2 && cd my_container2
$ wget https://github.com/alpinelinux/docker-alpine/raw/c5510d5b1d2546d133f7b0938690c3c1e2cd9549/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz

$ mkdir rootfs
$ tar -C rootfs -xvf alpine-minirootfs-3.11.6-x86_64.tar.gz

$ runc spec
$ sudo runc run mycont2  # `run` is a combination of `create` and `start`
/# hostname
> runc

```

**Well, I hope you noticed, that we haven't used any image-related facilities so far.** That's because we don't really need images to create and/or run containers with _runc_ or any other OCI-compliant runtime . And just to make it clear, when I say [_images_](https://github.com/opencontainers/image-spec) here I mean these layered beasts we build using _Dockerfiles_, push and pull from the registries like [Docker Hub](https://hub.docker.com/), and base other images on.

**好吧，我希望你注意到，到目前为止我们还没有使用任何与图像相关的工具。**那是因为我们真的不需要图像来使用 _runc_ 或任何其他 OCI 兼容的运行时来创建和/或运行容器.只是为了说明清楚，当我在这里说 [_images_](https://github.com/opencontainers/image-spec) 时，我的意思是我们使用 _Dockerfiles_ 构建的这些分层野兽，从 [Docker Hub] 之类的注册表中推送和拉取，并基于其他图像。

Reflecting on the exercise from above, we can say that _runc_ needs just a regular filesystem directory with at least one executable file inside and a `config.json`. This combination is called [a bundle](https://github.com/opencontainers/runtime-spec/blob/master/bundle.md).

回顾上面的练习，我们可以说 _runc_ 只需要一个常规文件系统目录，里面至少有一个可执行文件和一个 `config.json`。这种组合称为 [a bundle](https://github.com/opencontainers/runtime-spec/blob/master/bundle.md)。

The runtime doesn't know anything about the [storage driver](https://docs.docker.com/storage/storagedriver/select-storage-driver/) used to mount this folder, it's not even aware of its potential layered structure , thanks to [union mount](https://en.wikipedia.org/wiki/Union_mount) abstracting it away. Of course, it has nothing to do with the image building procedure and all those `Dockerfile` directives. The content of a bundle is also out of scope for the runtime, i.e. you can put a full Linux distro in there if you will, or keep it as lightweight as a single statically linked executable.

运行时对用于挂载此文件夹的 [storage driver](https://docs.docker.com/storage/storagedriver/select-storage-driver/) 一无所知，甚至不知道其潜在的分层结构，感谢 [union mount](https://en.wikipedia.org/wiki/Union_mount) 将其抽象化。当然，它与镜像构建过程和所有这些 `Dockerfile` 指令无关。捆绑包的内容也超出了运行时的范围，也就是说，如果愿意，您可以将完整的 Linux 发行版放入其中，或者使其像单个静态链接的可执行文件一样轻量级。

Any compliant container runtime should adhere to the [OCI Runtime Specification](https://github.com/opencontainers/runtime-spec) and surprisingly or not the specification never [mentions](https://github.com/opencontainers/runtime-spec/search?q=image&unscoped_q=image) images! Even though the [OCI Image Format specification](https://github.com/opencontainers/image-spec) is its sibling project. It's all about the strict separation of concerns.

任何兼容的容器运行时都应遵守 [OCI 运行时规范](https://github.com/opencontainers/runtime-spec) 并且令人惊讶的是，该规范从不[提及](https://github.com/opencontainers/runtime-spec/search?q=image&unscoped_q=image) 图像！尽管 [OCI 图像格式规范](https://github.com/opencontainers/image-spec) 是它的兄弟项目。这完全是关于关注点的严格分离。

Docker (or any other container engine like _containerd_, or _podman_) takes an image and converts it to an OCI bundle before invoking the lower-level container runtime like _runc_. I.e. it's an engine who is supposed to unpack an image, extract environment variables, exposed ports, etc. that were set on the building stage; prepare the `config.json` file using these details; mount filesystem layers from the image, and finally pass it over to the OCI runtime in a form of a bundle.

Docker（或任何其他容器引擎，如 _containerd_ 或 _podman_）在调用较低级别的容器运行时（如 _runc_）之前，会获取一个映像并将其转换为 OCI 包。 IE。它是一个引擎，应该解压缩映像，提取在构建阶段设置的环境变量、暴露的端口等；使用这些详细信息准备 `config.json` 文件；从映像挂载文件系统层，最后以包的形式将其传递给 OCI 运行时。

**So, why do we need all this hassle? Wouldn't it be much simpler if we just used bundles in the first place?** Apparently, this would simplify the containers' creation and starting logic. This could also positively affect the performance of the container runtime by offloading the overhead of the union mounts from the filesystem I/O.

**那么，为什么我们需要这些麻烦？如果我们一开始就使用 bundles 不是更简单吗？** 显然，这将简化容器的创建和启动逻辑。通过从文件系统 I/O 中卸载联合挂载的开销，这也可能对容器运行时的性能产生积极影响。

> Layered image structure combined with an overlay mount actually makes your container runtime slower.
>
> — Ivan Velichko (@iximiuz) [May 10, 2020](https://twitter.com/iximiuz/status/1259569913183731714?ref_src=twsrc%5Etfw)

> 分层镜像结构与覆盖挂载相结合实际上会使您的容器运行时变慢。
>
> — Ivan Velichko (@iximiuz) [2020 年 5 月 10 日](https://twitter.com/iximiuz/status/1259569913183731714?ref_src=twsrc%5Etfw)

### Why Do You Need Container Images

### 为什么需要容器镜像

**The actual reasons why we have images so widespread come from the image building and storage sides.** The layered image structure and immutability of the layers are the key contributors to the extremely powerful building and storage techniques that made Docker so popular back in the days. 

**我们拥有镜像如此广泛的实际原因来自镜像构建和存储方面。** 分层镜像结构和层的不变性是极其强大的构建和存储技术的关键贡献者，这些技术使 Docker 在过去如此受欢迎那些日子。

Imagine for a second, that we use raw bundles to store and distribute containers. In that case, if you needed to create two new custom bundles both based on the base `ubuntu` bundle, you'd have to copy over all the files from the `ubuntu` bundle into two new folders. And if you needed to create a third- or fourth-level bundles, you'd need to apply the same procedure over and over again multiplying the number of duplicate files on your disk. Moving all these files around would also add a significant build time overhead.

想象一下，我们使用原始包来存储和分发容器。在这种情况下，如果您需要创建两个都基于基础 `ubuntu` 包的新自定义包，则必须将 `ubuntu` 包中的所有文件复制到两个新文件夹中。如果您需要创建三级或四级包，则需要一遍又一遍地应用相同的过程，以增加磁盘上重复文件的数量。移动所有这些文件也会增加大量的构建时间开销。

Of course, we need to somehow avoid this inefficiency and the layered storage seems like a great solution. If we pack the `ubuntu` bundle in a tar archive and compute its checksum ( `sha256`), we will always be able to refer to it later on using [content-addressable](https://en.wikipedia.org/wiki/Content-addressable_storage) techniques. I.e. every new image derived from the base `ubuntu` image will be referring only the hash of the uniquely presented base image and will not include its actual content.

当然，我们需要以某种方式避免这种低效率，分层存储似乎是一个很好的解决方案。如果我们将 `ubuntu` 包打包到 tar 存档中并计算其校验和（`sha256`），我们将始终能够在以后使用 [content-addressable](https://en.wikipedia.org/wiki/Content-addressable_storage) 技术。 IE。从基础`ubuntu` 图像派生的每个新图像将仅引用唯一呈现的基础图像的哈希值，而不包括其实际内容。

![Images layers form a directed acyclic graph.](http://iximiuz.com/you-dont-need-an-image-to-run-a-container/images-dag.png)

_Images form a directed acyclic graph._

_图像形成有向无环图。_

The layered image structure solves the problem of the efficient building (by eradicating the need for copying files, hence decreasing the build time) and storage (by eradicating file duplication, hence decreasing the space requirements). But what if we want to run hundreds of instances of exactly the same container on one Linux host? We may need to have hundreds of copies of the same bundle and this may occupy the full disk. Union mount to the rescue! Thanks to the immutability of the layers, we need to unpack the image only once and then mount the underlying layers as many times as we need. The only mutable layer will be the topmost (empty) layer corresponding to the bundle folder of every container.

分层镜像结构解决了高效构建（通过消除复制文件的需要，从而减少构建时间）和存储（通过消除文件重复，从而减少空间需求）的问题。但是如果我们想在一台 Linux 主机上运行数百个完全相同的容器实例怎么办？我们可能需要有数百个同一个包的副本，这可能会占用整个磁盘。联盟坐骑来救援！由于层的不变性，我们只需要解压一次镜像，然后根据需要多次挂载底层。唯一的可变层将是对应于每个容器的 bundle 文件夹的最顶层（空）层。

Obviously, the pros of having images outweighed the negative impact of the layered mounts and the need for a slightly more complex container management procedures.

显然，拥有映像的优点超过了分层安装的负面影响以及对稍微复杂的容器管理程序的需求。

### Wrapping up

###  包起来

So, the moral may have sounded something like that: while the images aren't really needed to run containers, they have made the usage of the containers simple and handy and it contributed a lot to the containers' popularization.

所以，寓意可能是这样的：虽然运行容器并不是真正需要图像，但它们使容器的使用变得简单方便，并且为容器的普及做出了很大贡献。

Make code, not war!

编写代码，而不是战争！

### Related articles

###  相关文章

- [Containers aren't Linux processes](http://iximiuz.com/en/posts/oci-containers/)
- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [Implementing Container Runtime Shim: runc](http://iximiuz.com/en/posts/implementing-container-runtime-shim/)
- [A journey from containerization to orchestration and beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)

- [容器不是 Linux 进程](http://iximiuz.com/en/posts/oci-containers/)
- [并非每个容器内部都有操作系统](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [你需要容器来构建镜像](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [容器网络很简单！](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [实现容器运行时 Shim: runc](http://iximiuz.com/en/posts/implementing-container-runtime-shim/)
- [从容器化到编排及其他的旅程](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)

[docker,](javascript: void 0) [linux,](javascript: void 0) [container](javascript: void 0)

[docker,](javascript: void 0) [linux,](javascript: void 0) [container](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

