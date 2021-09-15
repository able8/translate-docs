# How Docker Build Works Internally

# Docker Build 如何在内部工作

May 25, 2020 (Updated: August 7, 2021)

[Containers,](http://iximiuz.com/en/categories/?category=Containers) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

[容器，](http://iximiuz.com/en/categories/?category=Containers)[Linux/Unix](http://iximiuz.com/en/categories/?category=Linux/Unix)

_You need containers to build images. Yes, you've heard it right. Not another way around._

_你需要容器来构建镜像。是的，你没听错。不是另一种方式。_

For people who found their way to containers through Docker (well, most of us I believe) it may seem like _images_ are of somewhat primary nature. We've been taught to start from a _Dockerfile_, build an image using that file, and only then run a container from that image. Alternatively, we could run a container specifying an image from a registry, yet the main idea persists - an image comes first, and only then the container.

对于通过 Docker 找到容器的人（好吧，我相信我们大多数人）似乎 _images_ 似乎具有某种主要性质。我们被教导从 _Dockerfile_ 开始，使用该文件构建一个映像，然后才从该映像运行一个容器。或者，我们可以运行一个容器，从注册表中指定一个图像，但主要思想仍然存在——先有图像，然后才是容器。

**But what if I tell you that the actual workflow is reverse?** Even when you are building your very first image using Docker, podman, or buildah, you are already, albeit implicitly, running containers under the hood!

**但是，如果我告诉你实际的工作流程是相反的呢？** 即使你正在使用 Docker、podman 或 buildah 构建你的第一个镜像，你已经（虽然隐含地）在后台运行容器！

## How container images are created

## 容器镜像是如何创建的

Let's avoid any unfoundedness and take a closer look at the image building procedure. The easiest way to spot this behavior is to build a simple image using the following _Dockerfile_:

让我们避免任何毫无根据，并仔细研究图像构建过程。发现这种行为的最简单方法是使用以下 _Dockerfile_ 构建一个简单的图像：

```dockerfile
FROM debian:latest

RUN sleep 2 && apt-get update
RUN sleep 2 && apt-get install -y uwsgi
RUN sleep 2 && apt-get install -y python3

COPY some_file /
```

While building the image, try running `docker stats -a` in another terminal:

在构建映像时，尝试在另一个终端中运行 `docker stats -a`：

_Running `docker build` and `docker stats` in different terminals._

_在不同的终端中运行`docker build`和`docker stats`。_

Huh, we haven't been launching any containers ourselves, nevertheless, `docker stats` shows that there were 3 of them 🙈 But how come? 

嗯，我们自己还没有启动任何容器，不过，`docker stats` 显示有 3 个容器🙈 但是怎么会呢？

Simplifying a bit, [images](https://github.com/opencontainers/image-spec) can be seen as archives with a filesystem inside. Additionally, they may also contain some configurational data like a default command to be executed when a container starts, exposed ports, etc, but we will be mostly focusing on the filesystem part. Luckily, we already know, that technically [images aren't required to run containers](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/) . Unlike virtual machines, containers are just [isolated and restricted processes](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/) on your Linux host. They do form an isolated execution environment, including the personalized root filesystem, but the bare minimum to start a container is just a folder with a single executable file inside. So, when we are starting a container from an image, the image gets unpacked and its content is provided to the [container runtime](https://github.com/opencontainers/runtime-spec) in a form of a [filesystem bundle ](https://github.com/opencontainers/runtime-spec/blob/44341cdd36f6fee6ddd73e602f9e3eca1466052f/bundle.md), ie a regular directory containing the future root filesystem files and some configs (all those layers you may have started thinking about are abstracted away by a [union mount](https://en.wikipedia.org/wiki/Union_mount) driver like [overlay fs](https://dev.to/napicella/how-are-docker-images-built-a-look-into-the-linux-overlay-file-systems-and-the-oci-specification-175n)). Thus, if you don't have the image but you do need the `alpine` Linux distro as the execution environment, you always can grab Alpine's rootfs ( [2.6 MB](https://github.com/alpinelinux/docker-alpine/raw/c5510d5b1d2546d133f7b0938690c3c1e2cd9549/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz)) and put it to a regular folder on your disk, then mix in your application files, feed it to the container runtime and call it a day.

稍微简化一下，[图像](https://github.com/opencontainers/image-spec)可以被视为内部有文件系统的档案。此外，它们还可能包含一些配置数据，例如在容器启动时要执行的默认命令、暴露的端口等，但我们将主要关注文件系统部分。幸运的是，我们已经知道，从技术上讲[运行容器不需要图像](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/) .与虚拟机不同，容器只是 Linux 主机上的[隔离和受限进程](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)。它们确实形成了一个隔离的执行环境，包括个性化的根文件系统，但启动容器的最低要求只是一个包含单个可执行文件的文件夹。所以，当我们从一个镜像启动一个容器时，镜像被解包，其内容以[文件系统包的形式提供给 容器运行时，即包含未来根文件系统文件和一些配置的常规目录（您可能已经开始考虑的所有那些层都是抽象的通过 [union mount](https://en.wikipedia.org/wiki/Union_mount) 驱动程序，例如 [overlay fs](https://dev.to/napicella/how-are-docker-images-built-a-look-into-the-linux-overlay-file-systems-and-the-oci-specification-175n))。因此，如果您没有映像但确实需要 `alpine` Linux 发行版作为执行环境，您总是可以获取 Alpine 的 rootfs ( [2.6 MB](https://github.com/alpinelinux/docker-alpine/raw/c5510d5b1d2546d133f7b0938690c3c1e2cd9549/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz)) 并将其放入磁盘上的常规文件夹中，然后将其混合到您的应用程序容器文件中，并提供给它日。

However, to unleash the full power of containers, we need handy image building facilities. Historically, Dockerfiles have been serving this purpose. Any Dockerfile must have the `FROM` instruction at the very beginning. This instruction specifies the base image while the rest of the Dockerfile describes the difference between the base and the derived (i.e. current) images.

然而，要释放容器的全部力量，我们需要方便的图像构建工具。从历史上看，Dockerfiles 一直在为这个目的服务。任何 Dockerfile 的开头都必须有“FROM”指令。该指令指定了基础镜像，而 Dockerfile 的其余部分描述了基础镜像和派生（即当前）镜像之间的区别。

The most basic container image is a so-called [`scratch`](https://hub.docker.com/_/scratch) image. It corresponds to an empty folder and the `FROM scratch` instruction in a Dockerfile means _noop_.

最基本的容器镜像就是所谓的 [`scratch`](https://hub.docker.com/_/scratch) 镜像。它对应于一个空文件夹，Dockerfile 中的“FROM scratch”指令表示 _noop_。

Now, let's take a look at the beloved `alpine` image:

现在，让我们来看看心爱的 `alpine` 图像：

```dockerfile
# https://github.com/alpinelinux/docker-alpine/blob/v3.11/x86_64/Dockerfile

FROM scratch
ADD alpine-minirootfs-3.11.6-x86_64.tar.gz /
CMD ["/bin/sh"]
```

I.e. to make the Alpine Linux distro image we just need to copy its root filesystem to an empty folder ( _scratch_ image) and that's it! Well, I bet Dockerfiles you've seen so far a rarely that trivial. More often than not, we need to utilize distro's facilities to prepare the file system of the future container and one of the most common examples is probably when we need to pre-install some external packages using `yum`, `apt`, or ` apk`:

IE。要制作 Alpine Linux 发行版映像，我们只需要将其根文件系统复制到一个空文件夹（_scratch_ 映像），就是这样！好吧，我敢打赌，到目前为止，您所看到的 Dockerfiles 很少有那么简单。通常情况下，我们需要利用发行版的工具来准备未来容器的文件系统，最常见的例子之一可能是我们需要使用 `yum`、`apt` 或 `预先安装一些外部包。 apk`：

```dockerfile
FROM debian:latest

RUN apt-get install -y ca-certificates

```

But how can we have `apt` running if we are building this image, say, on a Fedora host? Containers to the rescue! Every time, Docker (or buildah, or podman, etc) encounters a `RUN` instruction in the Dockerfile it actually fires a new container! The bundle for this container is formed by the base image plus all the changes made by preceding instructions from the Dockerfile (if any). When the execution of the `RUN` step completes, all the changes made to the container's filesystem (so-called _diff_) become a new _layer_ in the image being built and the process repeats starting from the next Dockerfile instruction.

但是，如果我们在 Fedora 主机上构建此映像，我们如何让 `apt` 运行？容器来救援！每次 Docker（或 buildah，或 podman 等）在 Dockerfile 中遇到 `RUN` 指令时，它实际上都会触发一个新容器！该容器的包由基础镜像加上 Dockerfile（如果有）的前面指令所做的所有更改构成。当 `RUN` 步骤的执行完成时，对容器文件系统（所谓的 _diff_）所做的所有更改都成为正在构建的映像中的一个新 _layer_，并且该过程从下一个 Dockerfile 指令开始重复。

![Building image using Dockerfile](http://iximiuz.com/you-need-containers-to-build-an-image/kdpv.png)

_Building image using Dockerfile_

_使用 Dockerfile 构建镜像_

Getting back to the original example from the beginning of this article, a mind-reader could have noticed that the number of containers we've seen in the second terminal corresponded exactly to the number of `RUN` instructions in the Dockerfile. For people with a solid understanding of the internal kitchen of containers it may sound obvious. However, for the rest of us possessing rather hands-on containers experience, the Dockerfile-based (and immensely popularized) workflow may instead obscure some things.

回到本文开头的原始示例，读心者可能已经注意到，我们在第二个终端中看到的容器数量与 Dockerfile 中“RUN”指令的数量完全对应。对于对容器的内部厨房有深入了解的人来说，这听起来很明显。然而，对于我们其他拥有相当动手容器经验的人来说，基于 Dockerfile（并且非常流行）的工作流程可能会掩盖一些事情。

Luckily, even though Dockerfiles are a de facto standard to describe images, it's not the only existing way. Thus, when using Docker, we can [`commit`](https://docs.docker.com/engine/reference/commandline/commit/) any running container to produce a new image. All the changes made to the container's filesystem by any command run inside of it since its start will form the topmost layer of the image created by the _commit_, while the base will be taken from the image used to create the said container. Although, if we decide to use this method we may face some reproducibility problems.

幸运的是，尽管 Dockerfiles 是描述图像的事实上的标准，但它并不是唯一存在的方式。因此，在使用 Docker 时，我们可以 [`commit`](https://docs.docker.com/engine/reference/commandline/commit/) 任何正在运行的容器来生成新的镜像。自启动以来，在容器内部运行的任何命令对容器文件系统所做的所有更改都将形成由 _commit_ 创建的映像的最顶层，而基础将取自用于创建所述容器的映像。虽然，如果我们决定使用这种方法，我们可能会面临一些可重复性问题。

## How to build image without Dockerfile

## 如何在没有 Dockerfile 的情况下构建镜像

Interesting enough, that some of the novel image building tools consider Dockerfile not as an advantage, but as a limitation. For instance, _buildah_ [promotes an alternative command-line image building procedure](https://www.redhat.com/sysadmin/building-buildah):

有趣的是，一些新颖的镜像构建工具并不认为 Dockerfile 是一种优势，而是一种限制。例如，_buildah_ [促进替代命令行映像构建过程](https://www.redhat.com/sysadmin/building-buildah)：

```bash
# Start building an image FROM fedora
$ buildah from fedora
> Getting image source signatures
> Copying blob 4c69497db035 done
> Copying config adfbfa4a11 done
> Writing manifest to image destination
> Storing signatures
> fedora-working-container  # <-- name of the newly started container!

# Examine running containers
$ buildah ps
> CONTAINER ID  BUILDER  IMAGE ID     IMAGE NAME                       CONTAINER NAME
> 2aa8fb539d69     *     adfbfa4a115a docker.io/library/fedora:latest  fedora-working-container

# Same as using ENV instruction in Dockerfile
$ buildah config --env MY_VAR="foobar" fedora-working-container

# Same as RUN in Dockerfile
$ buildah run fedora-working-container -- yum install python3
> ... installing packages

# Finally, make a layer (or an image)
$ buildah commit fedora-working-container

```

We can choose between interactive command-line image building or putting all these instructions to a _shell_ script, but regardless of the actual choice, _buildah_'s approach makes the need for running builder containers obvious. Skipping the rant about the pros and cons of this building technique, I just want to notice that the nature of the image building might have been much more apparent if _buildah_'s approach would predate Dockerfiles. 

我们可以在交互式命令行映像构建或将所有这些指令放入 _shell_ 脚本之间进行选择，但无论实际选择如何，_buildah_ 的方法都明显需要运行构建器容器。跳过关于这种构建技术的利弊的咆哮，我只想注意到，如果 _buildah_ 的方法早于 Dockerfiles，那么镜像构建的性质可能会更加明显。

Finalizing, let's take a brief look at two other prominent tools - Google's [kaniko](https://github.com/GoogleContainerTools/kaniko) and Uber's [makisu](https://github.com/uber/makisu). They try to tackle the image building problem from a slightly different angle. These tools don't really run containers while building images. Instead, they directly modify the local filesystem while following image building instructions 🤯 I.e. if you accidentally start such a tool on your laptop, it's highly likely that your host OS will be wiped and replaced with the rootfs of the image. So, beware. Apparently, these tools are supposed to be executed fully inside of an already existing container. This solves some security problems bypassing the need for elevating privileges of the builder process. Nevertheless, while the technique itself is very different from the traditional Docker's or buildah's approaches, the containers are still there. The main difference is that they have been moved out of the scope of the build tool.

最后，让我们简要介绍一下另外两个突出的工具——谷歌的 [kaniko](https://github.com/GoogleContainerTools/kaniko) 和 Uber 的 [makisu](https://github.com/uber/makisu)。他们试图从稍微不同的角度解决图像构建问题。这些工具在构建镜像时并没有真正运行容器。相反，他们直接修改本地文件系统，同时遵循图像构建说明🤯，即如果您不小心在笔记本电脑上启动了这样的工具，则您的主机操作系统很可能会被擦除并替换为映像的 rootfs。所以，当心。显然，这些工具应该完全在已经存在的容器内执行。这解决了一些安全问题，绕过了提升构建器进程权限的需要。尽管如此，虽然该技术本身与传统的 Docker 或 buildah 的方法有很大不同，但容器仍然存在。主要区别在于它们已移出构建工具的范围。

## Instead of conclusion

## 而不是结论

The concept of container images turned out to be very handy. The layered image structure in conjunction with union mounts like overlay fs made the storage and usage of images immensely efficient. The declarative Dockerfile-based approach enabled the reproducible and cachable building of artifacts. This allowed the idea of container images to become so wide-spread that sometime it may seem like it's an inalienable and archetypal part of the containerverse. However, as we saw in the article, from the implementation standpoint, containers are independent of images. Instead, most of the time we need containers to build images, not vice-a-verse.

容器镜像的概念非常方便。分层图像结构与诸如overlay fs 之类的联合挂载相结合，使得图像的存储和使用非常高效。基于 Dockerfile 的声明性方法支持可重现和可缓存的工件构建。这使得容器镜像的想法变得如此广泛，以至于有时它看起来像是容器世界中不可分割的原型部分。然而，正如我们在文章中看到的，从实现的角度来看，容器独立于图像。相反，大多数时候我们需要容器来构建镜像，而不是反之亦然。

Make code, not war!

编写代码，而不是战争！

### Appendix:  Image Building Tools

### 附录：图像构建工具

- [Docker](https://github.com/docker/docker-ce)
- [Podman](https://github.com/containers/libpod) & [Buildah](https://github.com/containers/buildah) / [intro](https://www.giantswarm.io/blog/building-container-images-with-podman-and-buildah)
- [BuildKit](https://github.com/moby/buildkit) / [intro](https://www.giantswarm.io/blog/container-image-building-with-buildkit)
- [img](https://github.com/genuinetools/img) / [intro](https://www.giantswarm.io/blog/building-container-images-with-img)
- [kaniko](https://github.com/GoogleContainerTools/kaniko) / [intro](https://www.giantswarm.io/blog/container-image-building-with-kaniko)
- [makisu](https://github.com/uber/makisu) / [intro](https://www.giantswarm.io/blog/container-image-building-with-makisu)

- [Docker](https://github.com/docker/docker-ce)
- [Podman](https://github.com/containers/libpod) & [Buildah](https://github.com/containers/buildah) / [介绍](https://www.giantswarm.io/blog/building-container-images-with-podman-and-buildah)
- [BuildKit](https://github.com/moby/buildkit) / [介绍](https://www.giantswarm.io/blog/container-image-building-with-buildkit)
- [img](https://github.com/genuinetools/img) / [介绍](https://www.giantswarm.io/blog/building-container-images-with-img)
- [kaniko](https://github.com/GoogleContainerTools/kaniko) / [介绍](https://www.giantswarm.io/blog/container-image-building-with-kaniko)
- [makisu](https://github.com/uber/makisu) / [介绍](https://www.giantswarm.io/blog/container-image-building-with-makisu)

### Related articles

###  相关文章

- [Containers aren't Linux processes](http://iximiuz.com/en/posts/oci-containers/)
- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [A journey from containerization to orchestration and beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)

- [容器不是 Linux 进程](http://iximiuz.com/en/posts/oci-containers/)
- [并非每个容器内部都有操作系统](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [你不需要图像来运行容器](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [容器网络很简单！](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [从容器化到编排及其他的旅程](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)

[linux,](javascript: void 0) [container,](javascript: void 0) [docker,](javascript: void 0) [podman,](javascript: void 0) [buildah,](javascript: void 0) [kaniko,](javascript: void 0) [makisu](javascript: void 0)

[linux,](javascript: void 0) [container,](javascript: void 0) [docker,](javascript: void 0) [podman,](javascript: void 0) [buildah,](javascript: void 0) [kaniko,](javascript: void 0) [makisu](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

