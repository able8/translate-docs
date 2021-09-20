# Does Container Image Have an OS Inside

# Container Image 里面有操作系统吗

May 7, 2020 (Updated: August 7, 2021)

[Containers,](http://iximiuz.com/en/categories/?category=Containers) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

[容器，](http://iximiuz.com/en/categories/?category=Containers)[Linux/Unix](http://iximiuz.com/en/categories/?category=Linux/Unix)

Not every container has an operating system inside, but every one of them needs your Linux kernel.

不是每个容器内部都有操作系统，但每个容器都需要你的 Linux 内核。

**Disclaimer 1:** before going any further it's important to understand the difference between a kernel, an operating system, and a distribution._

**免责声明 1：** 在进一步讨论之前，了解内核、操作系统和发行版之间的区别很重要。

1. [Linux kernel](https://en.wikipedia.org/wiki/Linux_kernel) is the core part of the Linux operating system. It's what originally Linus wrote.

1.  [Linux内核](https://en.wikipedia.org/wiki/Linux_kernel)是Linux操作系统的核心部分。这就是 Linus 最初写的。

2. [Linux operating system](https://en.wikipedia.org/wiki/Linux) is a combination of the kernel and a user-land (libraries, GNU utilities, config files, etc)._ - _[Linux distribution](https://en.wikipedia.org/wiki/Linux_distribution) is a particular version of the Linux operating system like Debian, CentOS, or Alpine._

2. [Linux 操作系统](https://en.wikipedia.org/wiki/Linux) 是内核和用户空间（库、GNU 实用程序、配置文件等)的组合。_ - _[Linux发行版](https://en.wikipedia.org/wiki/Linux_distribution) 是 Linux 操作系统的特定版本，如 Debian、CentOS 或 Alpine。_

_**Disclaimer 2:** the title of this article should have sounded like ["Does container image have a whole Linux distribution inside"](https://www.reddit.com/r/programming/comments/gfbaor/not_every_container_has_an_operating_system_inside/fpusvxu?utm_source=share&utm_medium=web2x), but I personally find this wording a bit boring_ 🤪

_**免责声明 2：** 这篇文章的标题应该听起来像 [“容器镜像是否有完整的 Linux 发行版”](https://www.reddit.com/r/programming/comments/gfbaor/not_every_container_has_an_operating_system_inside/fpusvxu?utm_source=share&utm_medium=web2x)，但我个人觉得这个用词有点无聊_🤪

### Does a Container have an Operating System inside?

### 容器内部是否有操作系统？

The majority of Docker examples out there explicitly or implicitly rely on some flavor of the Linux operating system running inside a container. I tried to quickly compile a list of the most prominent samples:

大多数 Docker 示例都显式或隐式地依赖于在容器内运行的 Linux 操作系统的某种风格。我试图快速编制一份最突出的样本列表：

Running an interactive shell in the `debian jessie` distribution:

在 debian jessie 发行版中运行交互式 shell：

```bash
$ docker run -it debian:jessie
```

Running an `nginx` web-sever in a container and examine its config using `cat` utility:

在容器中运行 nginx 网络服务器并使用 cat 实用程序检查其配置：

```bash
$ docker run -d -P --name nginx nginx:latest
$ docker exec -it nginx cat /etc/nginx/nginx.conf
```

Building an image based on Alpine Linux:

基于 Alpine Linux 构建镜像：

```bash
$ cat <<EOF > Dockerfile
FROM alpine:3.7
RUN apk add --no-cache mysql-client
ENTRYPOINT ["mysql"]
EOF

$ docker build -t mysql-alpine .
$ docker run mysql-alpine

```

And so forth and so on...

等等等等……

For the newcomers learning the containerization through hands-on experience, this may lead to a _false_ impression that containers are somewhat indistinguishable from full-fledged operating systems and that they are always based on well-known and wide-spread Linux distributions like `debian` , `centos`, or `alpine`.

对于通过实践经验学习容器化的新手来说，这可能会导致一种_错误_的印象，即容器与成熟的操作系统在某种程度上难以区分，并且它们总是基于知名且广泛使用的 Linux 发行版，例如`debian` 、`centos` 或 `alpine`。

At the same time, approaching the containerization topic from the theoretical side ( [1](https://techcrunch.com/2016/10/16/wtf-is-a-container/), [2](https://www.docker.com/resources/what-container), [3](https://cloud.google.com/containers)) may lead to a rather opposite impression that containers (unlike the traditional virtual machines) are supposed to pack only the application (ie your code) and its dependencies (ie some libraries) instead of a trying to ship a full operating system.

同时，从理论层面接近容器化话题（[1](https://techcrunch.com/2016/10/16/wtf-is-a-container/)、[2](https://www.docker.com/resources/what-container), [3](https://cloud.google.com/containers))可能会导致相反的印象，即容器（与传统虚拟机不同）应该打包只有应用程序（即您的代码）及其依赖项（即某些库)，而不是尝试发布完整的操作系统。

As it usually happens, the truth lies somewhere in between both statements. From the implementation standpoint, **a container is indeed just a process (or a bunch of processes) running on the Linux host**. The container process is isolated ( [namespaces](https://docs.docker.com/engine/security/security/#kernel-namespaces)) from the rest of the system and restricted from both the resource consumption ([cgroups]( https://docs.docker.com/engine/security/security/#control-groups)) and security ( [capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html), [AppArmor](https://docs.docker.com/engine/security/apparmor/),[Seccomp](https://docs.docker.com/engine/security/seccomp/)) standpoints. But in the end, this is still a regular process, same as any other process on the host system.

正如通常发生的那样，真相介于两者之间。从实现的角度来看，**容器确实只是在 Linux 主机上运行的一个进程（或一堆进程）**。容器进程与系统的其余部分隔离 ( [namespaces](https://docs.docker.com/engine/security/security/#kernel-namespaces)) 并限制资源消耗 ( [cgroups]( https://docs.docker.com/engine/security/security/#control-groups))和安全性（[功能](http://man7.org/linux/man-pages/man7/capabilities.7.html)、[AppArmor](https://docs.docker.com/engine/security/apparmor/)、[Seccomp](https://docs.docker.com/engine/security/seccomp/)) 的立场。但最终，这仍然是一个常规进程，与主机系统上的任何其他进程一样。

> OCI/Docker containers thread:
>
> Containers are simply isolated and restricted Linux processes. [#Docker](https://twitter.com/hashtag/Docker?src=hash&ref_src=twsrc%5Etfw)[#containers](https://twitter.com/hashtag/containers?src=hash&ref_src=twsrc%5Etfw) [#linux](https://twitter.com/hashtag/linux?src=hash&ref_src=twsrc%5Etfw)
>
> — Ivan Velichko (@iximiuz) [May 10, 2020](https://twitter.com/iximiuz/status/1259569908385529859?ref_src=twsrc%5Etfw)

> OCI/Docker 容器线程：
>
> 容器只是隔离和受限的 Linux 进程。 [#Docker](https://twitter.com/hashtag/Docker?src=hash&ref_src=twsrc%5Etfw)[#containers](https://twitter.com/hashtag/containers?src=hash&ref_src=twsrc%5Etfw) [#linux](https://twitter.com/hashtag/linux?src=hash&ref_src=twsrc%5Etfw)
>
> — Ivan Velichko (@iximiuz) [2020 年 5 月 10 日](https://twitter.com/iximiuz/status/1259569908385529859?ref_src=twsrc%5Etfw)

Just run `docker run -d nginx` and conduct your own investigation:

只需运行`docker run -d nginx`并进行你自己的调查：

![ps axf output (excerpt)](http://iximiuz.com/not-every-container-has-an-operating-system-inside/nginx-ps2.png)

_`ps axf` output (excerpt)_ 

_`ps axf` 输出（摘录）_

![systemctl status output (excerpt)](http://iximiuz.com/not-every-container-has-an-operating-system-inside/nginx-cgroups2.png)

_`systemctl status` output (excerpt)_

_`systemctl status` 输出（摘录）_

![sudo lsns output](http://iximiuz.com/not-every-container-has-an-operating-system-inside/nginx-lsns2.png)

_`sudo lsns` output_

Well, if a container is just a regular Linux process, we could try to run a single executable file inside of a container. I.e. instead of putting our application into a fully-featured Linux distribution, we will try to build a container image consisting of a folder with a single file inside. Upon the launch, this folder will become a root folder for the containerized environment.

好吧，如果容器只是一个普通的 Linux 进程，我们可以尝试在容器内运行单个可执行文件。 IE。我们不会将我们的应用程序放入功能齐全的 Linux 发行版中，而是尝试构建一个容器映像，其中包含一个文件夹，其中包含一个文件。启动后，此文件夹将成为容器化环境的根文件夹。

### Create a Container from scratch (with a single executable binary inside)

### 从头开始创建一个容器（里面有一个可执行的二进制文件）

If you have _Go_ installed on your system, you can utilize its handy cross-compilation abilities:

如果您的系统上安装了 _Go_，您可以利用其方便的交叉编译功能：

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello from OS-less container (Go edition)")
}

```

Build the program from above using:

使用以下方法从上面构建程序：

```bash
$ GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o hello
$ file hello
> hello: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked

```

_Click here to see how to compile a similar C program._

_单击此处查看如何编译类似的 C 程序。_

```c
// main.c
#include <stdio.h>

int main() {
    printf("Hello from OS-less container (C edition)\n");
}

```

Compile it using the following builder container:

使用以下构建器容器编译它：

```dockerfile
# Dockerfile.builder
FROM gcc:4.9
COPY main.c /main.c
CMD ["gcc", "-std=c99", "-static", "-o", "/out/hello", "/main.c"]

```

```bash
$ docker build -t builder -f Dockerfile.builder .
$ docker run -v `pwd`:/out builder
$ file hello
> hello: ELF 64-bit LSB executable, x86-64, version 1 (GNU/Linux), statically linked, for GNU/Linux 2.6.32

```

Finally, let's build the target container using the following trivial Dockerfile:

最后，让我们使用以下简单的 Dockerfile 构建目标容器：

```dockerfile
FROM scratch
COPY hello /
CMD ["/hello"]

```

```bash
$ docker build -t hello .
$ docker run hello
> Hello from OS-less container (Go edition)

```

If we now inspect the `hello` image with the wonderful [`dive`](https://github.com/wagoodman/dive) tool, we will notice that it consists of a directory with the single executable file in it:

如果我们现在使用美妙的 [`dive`](https://github.com/wagoodman/dive) 工具检查 `hello` 图像，我们会注意到它包含一个目录，其中包含单个可执行文件：

![Inspecting hello image with dive tool (screenshot)](http://iximiuz.com/not-every-container-has-an-operating-system-inside/dive-hello2.png)

_`dive hello`_

This exercise is roughly what the Docker's [hello-world](https://github.com/docker-library/hello-world) example does. There are two key moments here. First, we based our image on a so-called [`scratch`](https://hub.docker.com/_/scratch) image. This is just an empty image, i.e. the building starts from the empty folder and then just copies the executable file `hello` into it. Second, we used a statically linked binary file. I.e. there is no dependency on some shared libraries from the system. So, a bare Linux kernel is enough to execute it.

这个练习大致就是 Docker 的 [hello-world](https://github.com/docker-library/hello-world) 示例所做的。这里有两个关键时刻。首先，我们的图像基于所谓的 [`scratch`](https://hub.docker.com/_/scratch) 图像。这只是一个空图像，即建筑物从空文件夹开始，然后将可执行文件 `hello` 复制到其中。其次，我们使用了一个静态链接的二进制文件。 IE。不依赖于系统中的某些共享库。因此，一个裸 Linux 内核就足以执行它。

Now, what if we inspect the [`nginx`](https://hub.docker.com/_/nginx) image which we used at the beginning of this article?

现在，如果我们检查本文开头使用的 [`nginx`](https://hub.docker.com/_/nginx) 镜像会怎样？

![Inspecting nginx image with dive tool (screenshot)](http://iximiuz.com/not-every-container-has-an-operating-system-inside/dive-nginx2.png)

_`dive nginx`_

Well, the directory tree looks like a root filesystem of some Linux distribution. If we take a look at the [corresponding Dockerfile](https://github.com/nginxinc/docker-nginx/blob/594ce7a8bc26c85af88495ac94d5cd0096b306f7/mainline/buster/Dockerfile) we can notice that `nginx` image is based on [`debian `](https://hub.docker.com/_/debian):

嗯，目录树看起来像一些 Linux 发行版的根文件系统。如果我们查看[相应的 Dockerfile](https://github.com/nginxinc/docker-nginx/blob/594ce7a8bc26c85af88495ac94d5cd0096b306f7/mainline/buster/Dockerfile)，我们可以注意到 `nginx` 镜像基于 [`debian `](https://hub.docker.com/_/debian)：

```dockerfile
FROM debian:buster-slim

LABEL maintainer="NGINX Docker Maintainers <docker-maint@nginx.com>"

ENV NGINX_VERSION   1.17.10
ENV NJS_VERSION     0.3.9
ENV PKG_RELEASE     1~buster

...

```

And if we dive deeper and examine `debian:buster-slim` [Dockerfile](https://github.com/debuerreotype/docker-debian-artifacts/blob/d6ff3e75eeae3ea012c30bce9054336d99d1a20a/buster/slim/Dockerfile) we will see that it just copies [a root filesystem](https://github.com/debuerreotype/docker-debian-artifacts/blob/d6ff3e75eeae3ea012c30bce9054336d99d1a20a/buster/slim/rootfs.tar.xz) to an empty folder:

如果我们深入研究`debian:buster-slim` [Dockerfile](https://github.com/debuerreotype/docker-debian-artifacts/blob/d6ff3e75eeae3ea012c30bce9054336d99d1a20a/buster/slim/Dockerfile)，我们会看到它将[根文件系统](https://github.com/debuerreotype/docker-debian-artifacts/blob/d6ff3e75eeae3ea012c30bce9054336d99d1a20a/buster/slim/rootfs.tar.xz) 复制到一个空文件夹：

```dockerfile
FROM scratch
ADD rootfs.tar.xz /
CMD ["bash"]

```

[Combining Debian's user-land with the host's kernel](http://iximiuz.com/en/posts/from-docker-container-to-bootable-linux-disk-image/) containers start resembling fully-featured operating systems. With `nginx` image we can use the shell to interact with the container:

[将 Debian 的用户空间与主机的内核相结合](http://iximiuz.com/en/posts/from-docker-container-to-bootable-linux-disk-image/) 容器开始类似于全功能的操作系统。使用 nginx 镜像，我们可以使用 shell 与容器进行交互：

![Interactive shell with running nginx container.](http://iximiuz.com/not-every-container-has-an-operating-system-inside/nginx-exec-bash2.png)

_Interactive shell with running nginx container._

_带有运行 nginx 容器的交互式 shell。_

Can we do the same for our slim `hello` container? Obviously not, there is no `bash` executable inside:

我们可以为我们纤细的 `hello` 容器做同样的事情吗？显然不是，里面没有 `bash` 可执行文件：

![Demonstraiting that hello container doesn't have bash inside.](http://iximiuz.com/not-every-container-has-an-operating-system-inside/hello-run-bash2.png)

_`hello` container doesn't have `bash` inside._

_`hello` 容器里面没有 `bash`。_

### Wrapping up

###  总结

So, what should be the conclusion here? The [virtualization capabilities](https://en.wikipedia.org/wiki/OS-level_virtualization) of containers turned out to be so powerful that people started packing fully-featured user-lands like `debian` (or more lightweight alternatives like `alpine` or `busybox`) into containers. By virtue of this ability:

那么，这里的结论应该是什么？容器的[虚拟化功能](https://en.wikipedia.org/wiki/OS-level_virtualization) 变得如此强大，以至于人们开始打包功能齐全的用户空间，如`debian`（或更轻量级的替代品，如`alpine` 或 `busybox`) 放入容器中。凭借这种能力：

- We can play with various Linux distribution using a simple`docker run -it fedora bash`.
- We can use OS commands including package managers like`yum` or `apt` while building our images.
- We can interact with running containers using various OS utilities.

- 我们可以使用简单的`docker run -it fedora bash` 来玩各种 Linux 发行版。
- 我们可以在构建镜像时使用操作系统命令，包括包管理器，如 `yum` 或 `apt`。
- 我们可以使用各种操作系统实用程序与正在运行的容器进行交互。

But with great power comes great responsibility. Huge containers carrying lots of unnecessary tools slow down deployments and increase the surface of potential cyberattacks.

但权力越大，责任越大。携带大量不必要工具的巨大容器会减慢部署速度并增加潜在网络攻击的可能性。

Make code, not war!

编写代码，而不是战争！

### Related articles

###  相关文章

- [Containers aren't Linux processes](http://iximiuz.com/en/posts/oci-containers/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [A journey from containerization to orchestration and beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)
- [From Docker Container to Bootable Linux Disk Image](http://iximiuz.com/en/posts/from-docker-container-to-bootable-linux-disk-image/)

- [容器不是 Linux 进程](http://iximiuz.com/en/posts/oci-containers/)
- [你不需要图像来运行容器](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [你需要容器来构建镜像](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [容器网络很简单！](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [从容器化到编排及其他的旅程](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)
- [从 Docker 容器到可启动 Linux 磁盘映像](http://iximiuz.com/en/posts/from-docker-container-to-bootable-linux-disk-image/)

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

