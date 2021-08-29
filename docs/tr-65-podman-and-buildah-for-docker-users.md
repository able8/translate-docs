# Podman and Buildah for Docker users

# Docker 用户的 Podman 和 Buildah

February 21, 2019 From: https://developers.redhat.com/blog/2019/02/21/podman-and-buildah-for-docker-users#

I was asked recently on Twitter to better explain [Podman](https://github.com/containers/libpod)and [Buildah](https://github.com/containers/libpod) for someone familiar with Docker. Though there are many blogs and  tutorials out there, which I will list later, we in the community have  not centralized an explanation of how Docker users move from Docker to  Podman and Buildah. Also what role does Buildah play? Is Podman  deficient in some way that we need both Podman and Buildah to replace  Docker?

最近在 Twitter 上有人要求我为熟悉 Docker 的人更好地解释 [Podman](https://github.com/containers/libpod)和 [Buildah](https://github.com/containers/libpod)。虽然有很多博客和教程，我将在后面列出，但我们社区中并没有集中解释 Docker 用户如何从 Docker 迁移到 Podman 和 Buildah。 Buildah 扮演什么角色？ Podman 是否在某些方面存在缺陷，以至于我们需要 Podman 和 Buildah 来取代 Docker？

This article answers those questions and shows how to migrate to Podman.

本文回答了这些问题，并展示了如何迁移到 Podman。

## How does Docker work?

## Docker 是如何工作的？

First, let’s be clear about how Docker works; that will help us to  understand the motivation for Podman and also for Buildah. If you are a  Docker user, you understand that there is a daemon process that must be  run to service all of your Docker commands. I can’t claim to understand  the motivation behind this but I imagine it seemed like a great idea, at the time, to do all the cool things that Docker does in one place and  also provide a useful API to that process for future evolution. In the  diagram below, we can see that the Docker daemon provides all the  functionality needed to:

首先，让我们清楚Docker是如何工作的；这将有助于我们了解 Podman 和 Buildah 的动机。如果您是 Docker 用户，您就会明白必须运行一个守护进程来为您的所有 Docker 命令提供服务。我不能声称理解这背后的动机，但我认为在当时，在一个地方完成 Docker 所做的所有很酷的事情，并为该过程提供有用的 API 以供未来发展似乎是一个好主意。在下图中，我们可以看到 Docker 守护进程提供了以下所需的所有功能：

- Pull and push images from an image registry
- Make copies of images in a local container storage and to add layers to those containers
- Commit containers and remove local container images from the host repository
- Ask the kernel to run a container with the right namespace and cgroup, etc.

- 从映像注册表中拉取和推送映像
- 在本地容器存储中制作镜像副本并向这些容器添加层
- 提交容器并从主机存储库中删除本地容器映像
- 要求内核运行具有正确命名空间和 cgroup 等的容器。

Essentially the Docker daemon does all the work with registries, images, containers, and the kernel. The Docker command-line interface  (CLI) asks the daemon to do this on your behalf.

本质上，Docker 守护进程使用注册表、镜像、容器和内核完成所有工作。 Docker 命令行界面 (CLI) 要求守护程序代表您执行此操作。

[![How does Docker Work -- Docker architecture overview](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig1.png)](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig1.png)

This article does not get into the detailed pros and cons of the  Docker daemon process. There is much to be said in favor of this  approach and I can see why, in the early days of Docker, it made a lot  of sense. Suffice it to say that there were several reasons why Docker  users were concerned about this approach as usage went up. To list a  few:

本文不深入探讨 Docker 守护进程的详细优缺点。有很多支持这种方法的说法，我明白为什么在 Docker 的早期，它很有意义。可以说，随着使用量的增加，Docker 用户对这种方法感到担忧有几个原因。列举几个：

- A single process could be a single point of failure.
- This process owned all the child processes (the running containers).
- If a failure occurred, then there were orphaned processes.
- Building containers led to security vulnerabilities.
- All Docker operations had to be conducted by a user (or users) with the same full root authority.

- 单个进程可能是单点故障。
- 此进程拥有所有子进程（正在运行的容器）。
- 如果发生故障，则存在孤立进程。
- 构建容器导致安全漏洞。
- 所有 Docker 操作都必须由具有相同完全 root 权限的用户（或多个用户）执行。

There are probably more. Whether these issues have been fixed or you disagree with this characterization is not something this article  is going to debate. We in the community believe that Podman has  addressed many of these problems. If you want to take advantage of  Podman’s improvements, then this article is for you.

可能还有更多。这些问题是否已得到解决，或者您不同意这种描述，本文不打算讨论。我们社区认为 Podman 已经解决了其中的许多问题。如果您想利用 Podman 的改进，那么本文适合您。

The Podman approach is simply to directly interact with the image  registry, with the container and image storage, and with the Linux  kernel through the runC container runtime process (not a daemon).

Podman 方法只是通过 runC 容器运行时进程（不是守护进程）直接与镜像注册表、容器和镜像存储以及 Linux 内核进行交互。

[![Podman architectural approach](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig2.png)](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig2.png)

Now that we’ve discussed some of the motivation it’s time to discuss  what that means for the user migrating to Podman. There are a few  things to unpack here and we’ll get into each one separately:

现在我们已经讨论了一些动机，是时候讨论这对用户迁移到 Podman 意味着什么了。这里有一些东西需要解压，我们将分别介绍：

- You install Podman instead of Docker. You do not need to start or manage a daemon process like the Docker daemon.
- The commands you are familiar with in Docker work the same for Podman.
- Podman stores its containers and images in a different place than Docker.
- Podman and Docker images are compatible.
- Podman does more than Docker for [Kubernetes](https://developers.redhat.com/topics/kubernetes/) environments.
- What is Buildah and why might I need it?

- 您安装 Podman 而不是 Docker。您不需要像 Docker 守护进程那样启动或管理守护进程。
- 您在 Docker 中熟悉的命令对 Podman 的工作方式相同。
- Podman 将其容器和图像存储在与 Docker 不同的位置。
- Podman 和 Docker 镜像兼容。
- Podman 为 [Kubernetes](https://developers.redhat.com/topics/kubernetes/) 环境所做的不仅仅是 Docker。
- Buildah 是什么，我为什么需要它？



## Installing Podman 

## 安装 Podman

If you are using Docker today, you can remove it when you decide to  make the switch. However, you may wish to keep Docker around while you  try out Podman. There are some useful [tutorials](https://github.com/containers/libpod/blob/master/docs/tutorials/podman_tutorial.md) and an awesome [demonstration](https://github.com/containers/Demos /tree/master/building/buildah_intro) that you may wish to run through first so you can understand the  transition more. One example in the demonstration requires Docker in  order to show compatibility.

如果您现在正在使用 Docker，则可以在决定进行切换时将其删除。但是，您可能希望在试用 Podman 时保留 Docker。有一些有用的 [教程](https://github.com/containers/libpod/blob/master/docs/tutorials/podman_tutorial.md)和一个很棒的 [演示](https://github.com/containers/Demos /tree/master/building/buildah_intro)，您可能希望首先运行它，以便您可以更多地了解转换。演示中的一个示例需要 Docker 以显示兼容性。

To install Podman on [Red Hat Enterprise Linux](https://developers.redhat.com/products/rhel/overview/) 7.6 or later, use the following; if you are using Fedora, then replace `yum` with `dnf`:

要在 [Red Hat Enterprise Linux](https://developers.redhat.com/products/rhel/overview/) 7.6 或更高版本上安装 Podman，请使用以下命令；如果你使用的是 Fedora，那么用 `dnf` 替换 `yum`：

```
# yum -y install podman
```


## Podman commands are the same as Docker’s

## Podman 命令与 Docker 相同

When building Podman, the goal was to make sure that Docker users  could easily adapt. So all the commands you are familiar with also exist with Podman. In fact, the claim is made that if you have existing  scripts that run Docker you can create a `docker` alias for `podman` and all your scripts should work (`alias docker=podman`). Try it. Of course, you should stop Docker first (`systemctl stop docker`). There is a package you can install called `podman-docker` that does this for conversion for you. It drops a script at `/usr/bin/docker` that executes Podman with the same arguments.

在构建 Podman 时，目标是确保 Docker 用户可以轻松适应。所以你熟悉的所有命令也存在于 Podman 中。事实上，有人声称如果你有运行 Docker 的现有脚本，你可以为 `podman` 创建一个 `docker` 别名，并且你的所有脚本都应该可以工作（`alias docker=podman`）。尝试一下。当然，你应该先停止Docker（`systemctl stop docker`）。您可以安装一个名为 `podman-docker` 的包，它可以为您进行转换。它在 `/usr/bin/docker` 放置一个脚本，该脚本使用相同的参数执行 Podman。

The commands you are familiar with—`pull`, `push`, `build`, `run`, `commit`, `tag`, etc.—all exist with Podman. See the [manual pages for Podman](https://github.com/containers/Demos/tree/master/building/buildah_intro) for more information. One notable difference is that Podman has added  some convenience flags to some commands. For example, Podman has added `--all` (`-a`) flags for `podman rm` and `podman rmi`. Many users will find that very helpful.

你熟悉的命令——`pull`、`push`、`build`、`run`、`commit`、`tag` 等——都存在于Podman 中。有关详细信息，请参阅 [Podman 手册页](https://github.com/containers/Demos/tree/master/building/buildah_intro)。一个显着的区别是Podman 为一些命令添加了一些便利标志。例如，Podman 为 `podman rm` 和 `podman rmi` 添加了 `--all` (`-a`) 标志。许多用户会发现这非常有帮助。

You can also run Podman from your normal non-root user in Podman 1.0  on Fedora. RHEL support is aimed for version 7.7 and 8.1 onwards. Enhancements in userspace security have made this possible. Running  Podman as a normal user means that Podman will, by default, store images and containers in the user’s home directory. This is explained in the  next section. For more information on how Podman runs as a non-root  user, please check out Dan Walsh's article: [How does rootless Podman work?](https://opensource.com/article/19/2/how-does-rootless-podman-work)

您还可以在 Fedora 上的 Podman 1.0 中从普通的非 root 用户运行 Podman。 RHEL 支持针对 7.7 和 8.1 版本以上。用户空间安全性的增强使这成为可能。以普通用户身份运行 Podman 意味着，默认情况下 Podman 会将图像和容器存储在用户的主目录中。这将在下一节中解释。有关 Podman 如何以非 root 用户身份运行的更多信息，请查看 Dan Walsh 的文章：无根 Podman 是如何工作的？

## Podman and container images

## Podman 和容器镜像

When you first type `podman images`, you might be  surprised that you don’t see any of the Docker images you’ve already  pulled down. This is because Podman’s local repository is in `/var/lib/containers` instead of `/var/lib/docker`. This isn’t an arbitrary change; this new storage structure is based on the Open Containers Initiative (OCI) standards.

当您第一次输入 `podman images` 时，您可能会惊讶于没有看到任何已经拉下的 Docker 镜像。这是因为 Podman 的本地存储库位于 `/var/lib/containers` 而不是 `/var/lib/docker`。这不是随意更改；这种新的存储结构基于开放容器倡议 (OCI) 标准。

In 2015, Docker, Red Hat, CoreOS, SUSE, Google, and other leaders in  the Linux containers industry created the Open Container Initiative in  order to provide an independent body to manage the standard  specifications for defining container images and the runtime. In order  to maintain that independence, the [containers/image](https://github.com/containers/image)and [containers/storage](https://github.com/containers/storage) projects were created on [ GitHub](https://github.com/containers). 

2015 年，Docker、Red Hat、CoreOS、SUSE、Google 和 Linux 容器行业的其他领导者创建了 Open Container Initiative，以提供一个独立的机构来管理定义容器镜像和运行时的标准规范。为了保持这种独立性，[containers/image](https://github.com/containers/image)和 [containers/storage](https://github.com/containers/storage) 项目是在 [ GitHub](https://github.com/containers)。

Since you can run `podman` without being root, there needs to be a separate place where `podman` can write images. Podman uses a repository in the user’s home directory: `~/.local/share/containers`. This avoids making `/var/lib/containers` world-writeable or other practices that might lead to potential  security problems. This also ensures that every user has separate sets  of containers and images and all can use Podman concurrently on the same host without stepping on each other. When users are finished with their work, they can push to a common registry to share their image with  others.

既然你不用root就可以运行`podman`，那么就需要一个单独的地方让`podman`可以写图片。 Podman 使用用户主目录中的存储库：`~/.local/share/containers`。这避免了使 `/var/lib/containers` 世界可写或其他可能导致潜在安全问题的做法。这也确保了每个用户都有独立的容器和镜像集，并且所有人都可以在同一台主机上同时使用 Podman，而不会相互干扰。当用户完成他们的工作时，他们可以推送到一个公共注册表以与他人共享他们的图像。

Docker users coming to Podman find that knowing these locations is useful for debugging and for the important `rm -rf /var/lib/containers`, when you just want to start over. However, once you start using Podman, you’ll probably start using the new `-all` option to `podman rm` and `podman rmi` instead.

来到 Podman 的 Docker 用户发现，当您只想重新开始时，了解这些位置对于调试和重要的 `rm -rf /var/lib/containers` 很有用。然而，一旦你开始使用 Podman，你可能会开始使用新的 `-all` 选项来代替 `podman rm` 和 `podman rmi`。

## Container images are compatible between Podman and other runtimes

## 容器镜像在 Podman 和其他运行时之间兼容

Despite the new locations for the local repositories, the images  created by Docker or Podman are compatible with the OCI standard. Podman can push to and pull from popular container registries like [Quay.io](https://quay.io/) and Docker hub, as well as private registries. For example, you can  pull the latest Fedora image from the Docker hub and run it using  Podman. Not specifying a registry means Podman will default to searching through registries listed in the `registries.conf` file, in the order in which they are listed. An unmodified `registries.conf` file means it will look in the Docker hub first.

尽管本地存储库有新位置，但 Docker 或 Podman 创建的映像与 OCI 标准兼容。 Podman 可以推送和拉取流行的容器注册表，如 [Quay.io](https://quay.io/) 和 Docker hub，以及私有注册表。例如，您可以从 Docker 中心拉取最新的 Fedora 镜像并使用 Podman 运行它。不指定注册表意味着 Podman 将默认搜索 `registries.conf` 文件中列出的注册表，按照它们的列出顺序。未修改的 `registries.conf` 文件意味着它将首先在 Docker hub 中查找。

```
$ podman pull fedora:latest
$ podman run -it fedora bash
```


Images pushed to an image registry by Docker can be pulled down and  run by Podman. For example, an image (myfedora) I created using  Docker and pushed to my [Quay.io repository](https://quay.io/repository/ipbabble)(ipbabble) using Docker can be pulled and run with Podman as follows:

由 Docker 推送到镜像注册表的镜像可以被 Podman 拉下并运行。例如，我使用 Docker 创建并推送到我的 [Quay.io 存储库](https://quay.io/repository/ipbabble)(ipbabble) 的镜像 (myfedora) 可以使用 Docker 拉取并使用 Podman 运行，如下所示：

```
$ podman pull quay.io/ipbabble/myfedora:latest
$ podman run -it myfedora bash
```


Podman provides capabilities in its command-line `push` and `pull` commands to gracefully move images from `/var/lib/docker `to` /var/lib/containers` and vice versa. For example:

Podman 在其命令行 `push` 和 `pull` 命令中提供了功能，可以优雅地将图像从 `/var/lib/docker `到` /var/lib/containers`，反之亦然。例如：

```
$ podman push myfedora docker-daemon:myfedora:latest
```


Obviously, leaving out the `docker-daemon` above will default to pushing to the Docker hub. Using `quay.io/myquayid/myfedora` will push the image to the Quay.io registry (where `myquayid` below is your personal Quay.io account):

显然，省略上面的 `docker-daemon` 将默认推送到 Docker 集线器。使用 `quay.io/myquayid/myfedora` 会将镜像推送到 Quay.io 注册表（下面的 `myquayid` 是您的个人 Quay.io 帐户）：

```
$ podman push myfedora quay.io/myquayid/myfedora:latest
```


If you are ready to remove Docker, you should shut down the daemon  and then remove the Docker package using your package manager. But  first, if you have images you created with Docker that you wish to keep, you should make sure those images are pushed to a registry so that you  can pull them down later. Or you can use Podman to pull each image (for  example, fedora) from the host’s Docker repository into Podman’s  OCI-based repository. With RHEL you can run the following:

如果您准备删除 Docker，您应该关闭守护进程，然后使用您的包管理器删除 Docker 包。但首先，如果您希望保留使用 Docker 创建的图像，则应确保将这些图像推送到注册表，以便稍后将其拉下。或者，您可以使用 Podman 将每个映像（例如，fedora）从主机的 Docker 存储库拉入 Podman 基于 OCI 的存储库。使用 RHEL，您可以运行以下命令：

```
# systemctl stop docker
# podman pull docker-daemon:fedora:latest
# yum -y remove docker  # optional
```




## Podman helps users move to Kubernetes

## Podman 帮助用户迁移到 Kubernetes

Podman provides some extra features that help developers and  operators in Kubernetes environments. There are extra commands provided  by Podman that are not available in Docker. If you are familiar with  Docker and are considering using Kubernetes/[OpenShift](http://openshift.com/) as your container platform, then Podman can help you.

Podman 提供了一些额外的功能来帮助 Kubernetes 环境中的开发人员和运营商。 Podman 提供了 Docker 中没有的额外命令。如果您熟悉 Docker 并且正在考虑使用 Kubernetes/[OpenShift](http://openshift.com/) 作为您的容器平台，那么 Podman 可以帮助您。

Podman can generate a Kubernetes YAML file based on a running container using `podman generate kube`. The command `podman pod` can be used to help debug running Kubernetes pods along with the  standard container commands. For more details on how Podman can help  you transition to Kubernetes, see the following article by Brent Baude: [Podman can now ease the transition to Kubernetes and CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/).

Podman 可以使用 `podman generate kube` 基于正在运行的容器生成 Kubernetes YAML 文件。命令 `podman pod` 可用于帮助调试正在运行的 Kubernetes pod 以及标准容器命令。有关 Podman 如何帮助您过渡到 Kubernetes 的更多详细信息，请参阅 Brent Baude 的以下文章：[Podman 现在可以轻松过渡到 Kubernetes 和 CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/)。



## What is Buildah and why would I use it?

## Buildah 是什么，我为什么要使用它？

Buildah actually came first. And maybe that’s why some Docker users  get a bit confused. Why do these Podman evangelists also talk about  Buildah? Does Podman not do builds? 

Buildah实际上是第一个。也许这就是为什么一些 Docker 用户会感到有点困惑的原因。为什么这些 Podman 布道者也谈论 Buildah？ Podman 不做构建吗？

Podman does do builds and for those familiar with Docker, the build  process is the same. You can either build using a Dockerfile using `podman build` or you can run a container and make lots of changes and then commit  those changes to a new image tag. Buildah can be described as a superset of commands related to creating and managing container images and,  therefore, it has much finer-grained control over images. Podman’s `build` command contains a subset of the Buildah functionality. It uses the same code as Buildah for building.

Podman 确实会进行构建，对于熟悉 Docker 的人来说，构建过程是相同的。您可以使用 Dockerfile 使用 `podman build` 进行构建，也可以运行容器并进行大量更改，然后将这些更改提交到新的镜像标签。 Buildah 可以被描述为与创建和管理容器镜像相关的命令的超集，因此，它对镜像有更细粒度的控制。 Podman 的 `build` 命令包含 Buildah 功能的一个子集。它使用与 Buildah 相同的代码进行构建。

The most powerful way to use Buildah is to write Bash scripts for  creating your images—in a similar way that you would write a Dockerfile.

使用 Buildah 的最强大的方法是编写 Bash 脚本来创建您的图像——与编写 Dockerfile 的方式类似。

I like to think of the evolution in the following way. When Kubernetes moved to [CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/) based on the OCI runtime specification, there was no need to run a  Docker daemon and, therefore, no need to install Docker on any host in  the Kubernetes cluster for running pods and containers. Kubernetes could call CRI-O and it could call runC directly. This, in turn, starts the  container processes. However, if we want to use the same Kubernetes  cluster to do builds, as in the case of OpenShift clusters, then we  needed a new tool to perform builds that would not require the Docker  daemon and subsequently require that Docker be installed. Such a tool,  based on the `containers/storage` and `containers/image` projects, would also eliminate the security risk of the open Docker daemon socket during builds, which concerned many users.

我喜欢以下列方式思考进化。当Kubernetes基于OCI运行时规范迁移到[CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/)时，不需要运行Docker守护进程，因此无需在 Kubernetes 集群中的任何主机上安装 Docker 来运行 pod 和容器。 Kubernetes 可以调用 CRI-O，也可以直接调用 runC。反过来，这会启动容器进程。但是，如果我们想使用相同的 Kubernetes 集群进行构建，就像在 OpenShift 集群的情况下一样，那么我们需要一个新工具来执行不需要 Docker 守护程序并随后需要安装 Docker 的构建。这样一个工具，基于`containers/storage`和`containers/image`项目，也将消除构建过程中打开Docker守护进程套接字的安全风险，这是许多用户关注的问题。

Buildah (named for fun because of Dan Walsh’s Boston accent when  pronouncing "builder") fit this bill. For more information on Buildah,  see [buildah.io](https://buildah.io/) and specifically see the blogs and tutorials sections.

Buildah（因 Dan Walsh 在发音为“builder”时的波士顿口音而得名）符合这个要求。有关 Buildah 的更多信息，请参阅 [buildah.io](https://buildah.io/)，具体请参阅博客和教程部分。

There are a couple of extra things practitioners need to understand about Buildah:

从业者需要了解关于 Buildah 的一些额外的事情：

1. It allows for finer control of creating image layers. This is a  feature that many container users have been asking for for a long time. Committing many changes to a single layer is desirable.
2. Buildah’s `run` command is not the same as Podman’s `run` command. Because Buildah is for building images, the `run` command is *essentially the same as the Dockerfile* `RUN` *command*. In fact, I remember the week this was made explicit. I was foolishly  complaining that some port or mount that I was trying wasn’t working as I expected it to. Dan ([@rhatdan](https://twitter.com/rhatdan)) weighed in and said that Buildah should not be supporting running  containers in that way. No port mapping. No volume mounting. Those flags were removed. Instead `buildah run` is for running specific commands in order to help build a container image, for example, `buildah run dnf -y install nginx`.
3. Buildah can build images from scratch, that is, images with nothing in them at all. Nothing. In fact, looking at the container storage  created as a result of a `buildah from scratch` command  yields an empty directory. This is useful for creating very lightweight  images that contain only the packages needed in order to run your  application. 

 1.它允许更精细地控制创建图像层。这是许多容器用户长期以来一直要求的功能。对单个层进行许多更改是可取的。
2. Buildah 的 `run` 命令与 Podman 的 `run` 命令不同。因为 Buildah 是用来构建镜像的，所以 `run` 命令与 Dockerfile* 本质上是一样的* `RUN` *command*。事实上，我记得这是明确的那个星期。我愚蠢地抱怨我正在尝试的某些端口或安装没有按我预期的那样工作。 Dan ([@rhatdan](https://twitter.com/rhatdan)) 表示，Buildah 不应该支持以这种方式运行容器。没有端口映射。没有卷安装。那些标志被移除了。相反，`buildah run` 用于运行特定命令以帮助构建容器映像，例如，`buildah run dnf -y install nginx`。
3. Buildah 可以从头开始构建镜像，也就是什么都没有的镜像。没有。事实上，查看由 `buildah from scratch` 命令创建的容器存储会产生一个空目录。这对于创建仅包含运行应用程序所需的包的非常轻量级的映像很有用。

A good example use case for a scratch build is to consider the  development images versus staging or production images of a Java  application. During development, a Java application container image may  require the Java compiler and Maven and other tools. But in production,  you may only require the Java runtime and your packages. And, by the  way, you also do not require a package manager such as DNF/YUM or even  Bash. Buildah is a powerful CLI for this use case. See the diagram  below. For more information, see [Building a Buildah Container Image for Kubernetes](https://buildah.io/blogs/2018/03/01/building-buildah-container-image-for-kubernetes.html)and also this [Buildah introduction demo](https://github.com/containers/Demos/tree/master/building/buildah_intro).

临时构建的一个很好的示例用例是考虑 Java 应用程序的开发映像与暂存或生产映像。在开发过程中，Java 应用容器镜像可能需要 Java 编译器和 Maven 等工具。但在生产中，您可能只需要 Java 运行时和您的包。而且，顺便说一句，您也不需要诸如 DNF/YUM 甚至 Bash 之类的包管理器。 Buildah 是用于此用例的强大 CLI。见下图。有关更多信息，请参阅 [为 Kubernetes 构建 Buildah 容器映像](https://buildah.io/blogs/2018/03/01/building-buildah-container-image-for-kubernetes.html)以及此 [Buildah介绍演示](https://github.com/containers/Demos/tree/master/building/buildah_intro)。

[![Buildah is a powerful CLI](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig3-1024x703.png)](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig3.png)

Getting back to the evolution story...Now that we had solved the  Kubernetes runtime issue with CRI-O and runC, and we had solved the  build problem with Buildah, there was still one reason why Docker was  still needed on a Kubernetes host: debugging. How can we debug container issues on a host if we don't have the tools to do it? We would need to  install Docker, and then we are back where we started with the Docker  daemon on the host. Podman solves this problem.

回到进化故事……既然我们已经用 CRI-O 和 runC 解决了 Kubernetes 运行时问题，并且我们已经用 Buildah 解决了构建问题，那么在 Kubernetes 主机上仍然需要 Docker 的原因还有一个：调试。如果我们没有工具，我们如何在主机上调试容器问题？我们需要安装 Docker，然后我们又回到了从主机上的 Docker 守护进程开始的地方。 Podman 解决了这个问题。

Podman becomes a tool that solves two problems. It allows operators  to examine containers and images with commands they are familiar with  using. And it also provides developers with the same tools. So Docker  users, developers, or operators, can move to Podman, do all the fun  tasks that they are familiar with from using Docker, and do much more.

Podman 成为解决两个问题的工具。它允许操作员使用他们熟悉的命令检查容器和图像。它还为开发人员提供了相同的工具。因此，Docker 用户、开发人员或操作员可以迁移到 Podman，执行他们使用 Docker 时熟悉的所有有趣任务，并做更多事情。

## Conclusion

##  结论

I hope this article has been useful and will help you migrate to using Podman (and Buildah) confidently and successfully.

我希望本文对您有所帮助，并能帮助您自信且成功地迁移到使用 Podman（和 Buildah）。

For more information:

想要查询更多的信息：

- [Podman.io](https://podman.io/)and [Buildah.io](https://buildah.io/) project web sites

- [Podman.io](https://podman.io/)和 [Buildah.io](https://buildah.io/) 项目网站

- github.com/containers

    - github.com/containers

   projects (get involved, get the source, see what's being developed):

   项目（参与，获取源代码，查看正在开发的内容）：

  - [libpod](https://github.com/containers/libpod)(Podman)
   - [buildah](https://github.com/containers/buildah)
   - [image](https://github.com/containers/image)(code for working with OCI container images)
   - [storage](https://github.com/containers/storage)(code for local image and container storage)



## Related Articles

##  相关文章

- [Containers without daemons: Podman and Buildah available in RHEL 7.6 and RHEL 8 Beta](https://developers.redhat.com/blog/2018/11/20/buildah-podman-containers-without-daemons/)
- [Podman: Managing pods and containers in a local container runtime](https://developers.redhat.com/blog/2019/01/15/podman-managing-containers-pods/)
- [Managing containerized system services with Podman](https://developers.redhat.com/blog/2018/11/29/managing-containerized-system-services-with-podman/)(Use systemd to manage your podman containers)
- [Building a Buildah Container Image for Kubernetes](https://buildah.io/blogs/2018/03/01/building-buildah-container-image-for-kubernetes.html)
- [Podman can now ease the transition to Kubernetes and CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/)
- [Security Considerations for Container Runtimes](https://developers.redhat.com/blog/2018/12/19/security-considerations-for-container-runtimes/)(Video of Dan Walsh's talk from KubeCon 2018)
- [IoT edge development and deployment with containers through OpenShift: Part 1](https://developers.redhat.com/blog/2019/01/31/iot-edge-development-and-deployment-with-containers-through-openshift-part-1/) (Building and testing ARM64 containers on OpenShift using podman, qemu, binfmt_misc, and Ansible)

  
- [没有守护进程的容器：Podman 和 Buildah 在 RHEL 7.6 和 RHEL 8 Beta 中可用](https://developers.redhat.com/blog/2018/11/20/buildah-podman-containers-without-daemons/)
- [Podman：在本地容器运行时管理 Pod 和容器](https://developers.redhat.com/blog/2019/01/15/podman-managing-containers-pods/)
- [Managing containerized system services with Podman](https://developers.redhat.com/blog/2018/11/29/managing-containerized-system-services-with-podman/)（使用systemd来管理你的podman容器）
- [为Kubernetes构建Buildah容器镜像](https://buildah.io/blogs/2018/03/01/building-buildah-container-image-for-kubernetes.html)
- [Podman 现在可以轻松过渡到 Kubernetes 和 CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/)
- [容器运行时的安全考虑](https://developers.redhat.com/blog/2018/12/19/security-container-runtimes/)（Dan Walsh 在 KubeCon 2018 上的演讲视频）
- [通过 OpenShift 使用容器进行物联网边缘开发和部署：第 1 部分](https://developers.redhat.com/blog/2019/01/31/iot-edge-development-and-deployment-with-containers-through-openshift-part-1/)（使用 podman、qemu、binfmt_misc 和 Ansible 在 OpenShift 上构建和测试 ARM64 容器）


*Last updated:              June 17, 2021* 

*上次更新时间：2021 年 6 月 17 日*
