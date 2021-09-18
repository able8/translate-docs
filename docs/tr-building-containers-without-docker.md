# Building containers without Docker

# 不用 Docker 构建容器

25 January 2020

2020 年 1 月 25 日

In this post I'll outline several ways to build containers without the need for Docker itself. I'll use [OpenFaaS](https://github.com/openfaas/) as the case-study, which uses OCI-format container images for its workloads. The easiest way to think about OpenFaaS is as a CaaS platform for [Kubernetes](https://kubernetes.io) which can run microservices, and add in FaaS and event-driven tooling for free.

在这篇文章中，我将概述几种无需 Docker 本身即可构建容器的方法。我将使用 [OpenFaaS](https://github.com/openfaas/) 作为案例研究，它使用 OCI 格式的容器映像来处理其工作负载。考虑 OpenFaaS 的最简单方法是作为 [Kubernetes](https://kubernetes.io) 的 CaaS 平台，它可以运行微服务，并免费添加 FaaS 和事件驱动工具。

See also [OpenFaaS.com](https://openfaas.com/)

另见 [OpenFaaS.com](https://openfaas.com/)

The first option in the post will show how to use the built-in buildkit option for Docker's CLI, then [buildkit](https://github.com/moby/buildkit) stand-alone (on Linux only), followed by Google's container builder, [Kaniko](https://github.com/GoogleContainerTools/kaniko).

帖子中的第一个选项将展示如何使用 Docker 的 CLI 的内置 buildkit 选项，然后是 [buildkit](https://github.com/moby/buildkit) 独立（仅在 Linux 上)，然后是 Google 的容器构建器，[Kaniko](https://github.com/GoogleContainerTools/kaniko)。

This post covers tooling which can build an image from a Dockerfile, and so anything which limits the user to only Java (jib) or Go (ko) for instance is out of scope. I'll then wrap things up and let you know how to get in touch with suggestions, feedback and your own stories around wants and needs in container tooling.

这篇文章涵盖了可以从 Dockerfile 构建映像的工具，因此任何将用户限制为仅使用 Java (jib) 或 Go (ko) 的东西都超出了范围。然后我会总结一下，让你知道如何就容器工具的需求和需求与建议、反馈和你自己的故事取得联系。

## So what's wrong with Docker?

## 那么 Docker 有什么问题呢？

Nothing as such, Docker runs well on armhf, arm64, and on `x86_64`. The main Docker CLI has become a lot more than build/ship/run, and also lugs around several years of baggage, it now comes bundled with Docker Swarm and EE features.

并非如此，Docker 在 armhf、arm64 和 `x86_64` 上运行良好。主要的 Docker CLI 已经不仅仅是构建/交付/运行，而且还拖着几年的包袱，它现在捆绑了 Docker Swarm 和 EE 功能。

> Update for Nov 2020: anyone using Docker's set of official base-images should also read: [Preparing for the Docker Hub Rate Limits](https://inlets.dev/blog/2020/10/29/preparing-docker-hub-rate-limits.html)

> 2020 年 11 月更新：任何使用 Docker 官方基础镜像集的人还应该阅读：[为 Docker Hub 速率限制做准备](https://inlets.dev/blog/2020/10/29/preparing-docker-hub-rate-limits.html)

### Alternatives to Docker

### Docker 的替代品

There are a few efforts that attempt to strip "docker" back to its component pieces, the original UX we all fell in love with:

有一些努力试图将“docker”剥离回其组成部分，即我们都爱上的原始用户体验：

- [Docker](https://github.com/docker/docker) \- docker itself now uses containerd to run containers, and has support for enabling buildkit to do highly efficient, caching builds.

- [Docker](https://github.com/docker/docker) \- docker 本身现在使用 containerd 来运行容器，并且支持启用 buildkit 来执行高效的缓存构建。

- [Podman](https://podman.io/) and [buildah](https://github.com/containers/buildah) combination - RedHat / IBM's effort, which uses their own OSS toolchain to generate OCI images. Podman is marketed as being daemonless and rootless, but still ends up having to mount overlay filesystems and use a UNIX socket.

- [Podman](https://podman.io/) 和 [buildah](https://github.com/containers/buildah) 组合 - RedHat / IBM 的努力，它使用他们自己的 OSS 工具链来生成 OCI 图像。 Podman 被标榜为无守护进程和无根，但最终仍必须挂载覆盖文件系统并使用 UNIX 套接字。

- [pouch](https://github.com/alibaba/pouch) \- from Alibaba, pouch is billed as "An Efficient Enterprise-class Container Engine". It uses containerd just like Docker, and supports both container-level isolation with [runc](https://github.com/opencontainers/runc) and "lightweight VMs" such as [runV](https://github.com/hyperhq/runv). There's also more of a [focus on image distribution and strong isolation](https://github.com/alibaba/pouch/blob/master/docs/architecture.md).

- [pouch](https://github.com/alibaba/pouch) \- 来自阿里巴巴，pouch 被誉为“高效的企业级容器引擎”。它像 Docker 一样使用 containerd，并支持容器级隔离 [runc](https://github.com/opencontainers/runc) 和“轻量级 VM”，例如 [runV](https://github.com/hyperhq/runv)。还有更多的[专注于图像分发和强隔离](https://github.com/alibaba/pouch/blob/master/docs/architecture.md)。

- Stand-alone buildkit - buildkit was started by [Tõnis Tiigi](https://twitter.com/tonistiigi?lang=en) from Docker Inc as a brand new container builder with caching and concurrency in mind. buildkit currently only runs as a daemon, but you will hear people claim otherwise. They are forking the daemon and then killing it after a build.

- 独立构建工具包——构建工具包是由 Docker 公司的 [Tõnis Tiigi](https://twitter.com/tonistiigi?lang=en) 作为一个全新的容器构建器启动的，它考虑了缓存和并发性。 buildkit 目前仅作为守护进程运行，但您会听到人们声称并非如此。他们分叉守护进程，然后在构建后杀死它。

- [img](https://github.com/genuinetools/img) \- img was written by [Jess Frazelle](https://github.com/jessfraz) and is often quoted in these sorts of guides and is a wrapper for buildkit. That said, I haven't seen traction with it compared to the other options mentioned. The project was quite active [until late 2018 and has only received a few patches since](https://github.com/genuinetools/img/commits/master). img claims to be daemonless, but it uses buildkit so is probably doing some trickery there. I hear that `img` gives a better UX than buildkit's own CLI `buildctr`, but it should also be noted that img is only released for `x86_64` and there are no binaries for armhf / arm64.

- [img](https://github.com/genuinetools/img) \- img 是由 [Jess Frazelle](https://github.com/jessfraz) 编写的，经常在这些指南中被引用，是一个buildkit 的包装器。也就是说，与提到的其他选项相比，我还没有看到它的吸引力。该项目非常活跃 [直到 2018 年底，此后只收到了几个补丁](https://github.com/genuinetools/img/commits/master)。 img 声称是无守护进程的，但它使用 buildkit，所以可能在那里做了一些诡计。我听说 `img` 提供了比 buildkit 自己的 CLI `buildctr` 更好的用户体验，但还应该注意的是，img 仅针对 `x86_64` 发布，并且没有针对 armhf / arm64 的二进制文件。

> An alternative to `img` would be `k3c` which also includes a runtime component and plans to support ARM architectures.

> `img` 的替代方案是 `k3c`，它也包含一个运行时组件并计划支持 ARM 架构。

- [k3c](https://github.com/ibuildthecloud/k3c) \- Rancher's latest experiment which uses containerd and buildkit to re-create the original, classic, vanilla, lite experience of the original Docker version. 

- [k3c](https://github.com/ibuildthecloud/k3c) \- Rancher 的最新实验，使用 containerd 和 buildkit 重新创建原始 Docker 版本的原始、经典、vanilla、lite 体验。

Out of all the options, I think that I like k3c the most, but it is very nascient and bundles everything into one binary which is likely to conflict with other software, at present it runs its own embedded containerd and buildkit binaries.

在所有选项中，我认为我最喜欢 k3c，但它非常幼稚，将所有内容捆绑到一个二进制文件中，这可能会与其他软件冲突，目前它运行自己的嵌入式 containerd 和 buildkit 二进制文件。

> Note: If you're a RedHat customer and paying for support, then you really should use their entire toolchain to get the best value for your money. I checked out some of the examples and saw one that used my "classic" blog post on multi-stage builds. See for yourself which style you prefer [the buildah example](https://github.com/containers/buildah/blob/master/demos/buildah_multi_stage.sh) vs. [Dockerfile example](https://blog.alexellis.io/mutli-stage-docker-builds).

> 注意：如果您是 RedHat 客户并支付支持费用，那么您真的应该使用他们的整个工具链来获得最大的价值。我查看了一些示例，并看到了一个使用我关于多阶段构建的“经典”博客文章的示例。亲眼看看你更喜欢哪种风格 [buildah 示例](https://github.com/containers/buildah/blob/master/demos/buildah_multi_stage.sh) 与 [Dockerfile 示例](https://blog.alexellis.io/mutli-stage-docker-builds)。

So since we are focusing on the "build" piece here and want to look at relativelt stable options, I'm going to look at:

因此，由于我们在这里专注于“构建”部分并希望查看相对稳定的选项，因此我将查看：

- buildkit in Docker,
- buildkit stand-alone
- and kaniko.

- Docker 中的 buildkit，
- buildkit 独立
- 和卡尼科。

All of the above and more are now possible since the OpenFaaS CLI can output a standard "build context" that any builder can work with.

由于 OpenFaaS CLI 可以输出任何构建器都可以使用的标准“构建上下文”，因此上述所有和更多功能现在都成为可能。

## Build a test app

## 构建一个测试应用程序

Let's start with a Golang HTTP middleware, this is a cross between a function and a microservice and shows off how versatile OpenFaaS can be.

让我们从 Golang HTTP 中间件开始，这是一个函数和微服务之间的交叉，展示了 OpenFaaS 的多功能性。

```sh
faas-cli template store pull golang-middleware

faas-cli new --lang golang-middleware \
build-test --prefix=alexellis2

```

- `--lang` specifies the build template
- `build-test` is the name of the function
- `--prefix` is the Docker Hub username to use for pushing up our OCI image

- `--lang` 指定构建模板
- `build-test` 是函数名
- `--prefix` 是用于推送我们的 OCI 镜像的 Docker Hub 用户名

We'll get the following created:

我们将创建以下内容：

```
./
├── build-test
│   └── handler.go
└── build-test.yml

1 directory, 2 files

```

The handler looks like this, and is easy to modify. Additional dependencies can be added through vendoring or [Go modules](https://blog.golang.org/using-go-modules).

处理程序看起来像这样，并且很容易修改。可以通过 vendoring 或 [Go modules](https://blog.golang.org/using-go-modules) 添加其他依赖项。

```golang
package function

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
    var input []byte

    if r.Body != nil {
        defer r.Body.Close()

        body, _ := ioutil.ReadAll(r.Body)

        input = body
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("Hello world, input was: %s", string(input))))
}

```

### Build the normal way

### 以正常方式构建

The normal way to build this app would be:

构建此应用程序的正常方法是：

```sh
faas-cli build -f build-test.yml

```

A local cache of the template and Dockerfile is also available at `./template/golang-middleware/Dockerfile`

模板和 Dockerfile 的本地缓存也可在`./template/golang-middleware/Dockerfile` 获得

There are three images that are pulled in for this template:

为该模板提取了三个图像：

```
FROM openfaas/of-watchdog:0.7.3 as watchdog
FROM golang:1.13-alpine3.11 as build
FROM alpine:3.12

```

With the traditional builder, each of the images will be pulled in sequentially.

使用传统的构建器，每个图像将被依次拉入。

The wait a few moments and you're done, we now have that image in our local library.

稍等片刻，您就完成了，我们现在在本地图书馆中有该图像。

We can also push it up to a registry with `faas-cli push -f build-test.yml`.

我们还可以使用“faas-cli push -f build-test.yml”将其推送到注册表。

![seq](http://blog.alexellis.io/content/images/2020/01/seq.png)

### Build with Buildkit and Docker

### 使用 Buildkit 和 Docker 构建

This is the easiest change of all to make, and gives a fast build too.

这是所有更改中最简单的更改，并且也提供了快速构建。

```sh
DOCKER_BUILDKIT=1 faas-cli build -f build-test.yml

```

We'll see that with this approach, the Docker daemon automatically switches out its builder for buildkit.

我们将看到，使用这种方法，Docker 守护程序会自动为 buildkit 切换其构建器。

Buildkit offers a number of advantages:

Buildkit 提供了许多优点：

- More sophisticated caching
- Running later instructions first, when possible - i.e. downloading the "runtime" image, before the build in the "sdk" layer is even completed
- Super fast when building a second time

- 更复杂的缓存
- 在可能的情况下，首先运行后面的指令 - 即在“sdk”层中的构建完成之前下载“运行时”映像
- 第二次构建时超快

With buildkit, all of the base images can be pulled in to our local library at once, since the FROM (download) commands are not executed sequentially.

使用 buildkit，所有基础镜像都可以一次拉入我们的本地库，因为 FROM（下载）命令不是按顺序执行的。

```
FROM openfaas/of-watchdog:0.7.3 as watchdog
FROM golang:1.13-alpine3.11 as build
FROM alpine:3.11

```

This option works even on a Mac, since buildkit is proxied via the Docker daemon running in the VM.

这个选项甚至适用于 Mac，因为 buildkit 是通过在 VM 中运行的 Docker 守护进程代理的。

![dkit](http://blog.alexellis.io/content/images/2020/01/dkit.png)

### Build with Buildkit standalone

### 使用 Buildkit 独立构建

To build with Buildkit in a stand-alone setup we need to run buildkit separately on a Linux host, so we can't use a Mac.

要在独立设置中使用 Buildkit 进行构建，我们需要在 Linux 主机上单独运行 buildkit，因此我们不能使用 Mac。

`faas-cli build` would normally execute or fork `docker`, because the command is just a wrapper. So to bypass this behaviour we should write out a build context, that's possible via the following command:

`faas-cli build` 通常会执行或派生 `docker`，因为该命令只是一个包装器。因此，为了绕过这种行为，我们应该写出一个构建上下文，这可以通过以下命令实现：

```sh
faas-cli build -f build-test.yml --shrinkwrap

[0] > Building build-test.
Clearing temporary build folder: ./build/build-test/
Preparing ./build-test/ ./build/build-test//function
Building: alexellis2/build-test:latest with golang-middleware template.Please wait..
build-test shrink-wrapped to ./build/build-test/
[0] < Building build-test done in 0.00s.
[0] Worker done.

Total build time: 0.00

```

Our context is now available in the `./build/build-test/` folder with our function code and the template with its entrypoint and Dockerfile.

我们的上下文现在在 `./build/build-test/` 文件夹中可用，其中包含我们的函数代码和模板及其入口点和 Dockerfile。

```sh
./build/build-test/
├── Dockerfile
├── function
│   └── handler.go
├── go.mod
├── main.go
└── template.yml

1 directory, 5 files

```

Now we need to run buildkit, we can build from source, or grab upstream binaries.

现在我们需要运行 buildkit，我们可以从源代码构建，或者获取上游二进制文件。

```sh
curl -sSLf https://github.com/moby/buildkit/releases/download/v0.6.3/buildkit-v0.6.3.linux-amd64.tar.gz |sudo tar -xz -C /usr/local/bin/ --strip-components=1

```

If you checkout the releases page, you'll also find buildkit available for armhf and arm64, which is great for multi-arch.

如果您查看发布页面，您还会发现可用于 armhf 和 arm64 的 buildkit，这非常适合多架构。

Run the buildkit daemon in a new window:

在新窗口中运行 buildkit 守护进程：

```sh
sudo buildkitd
WARN[0000] using host network as the default
INFO[0000] found worker "l1ltft74h0ek1718gitwghjxy", labels=map[org.mobyproject.buildkit.worker.executor:oci org.mobyproject.buildkit.worker.hostname:nuc org.mobyproject.buildkit.worker.snapshotter:overlayfs], platforms=[linux/amd64 linux/386]
WARN[0000] skipping containerd worker, as "/run/containerd/containerd.sock" does not exist
INFO[0000] found 1 workers, default="l1ltft74h0ek1718gitwghjxy"
WARN[0000] currently, only the default worker can be used.
INFO[0000] running server on /run/buildkit/buildkitd.sock

```

Now let's start a build, passing in the shrink-wrapped location as the build-context. The command we want is `buildctl`, buildctl is a client for the daemon and will configure how to build the image and what to do when it's done, such as exporting a tar, ignoring the build or pushing it to a registry.

现在让我们开始构建，将收缩包装的位置作为构建上下文传递。我们想要的命令是`buildctl`，buildctl 是守护进程的客户端，它将配置如何构建映像以及完成后要做什么，例如导出 tar、忽略构建或将其推送到注册表。

```sh
buildctl build --help
NAME:
buildctl build - build

USAGE:

To build and push an image using Dockerfile:
    $ buildctl build --frontend dockerfile.v0 --opt target=foo --opt build-arg:foo=bar --local context=.--local dockerfile=.--output type=image,name=docker.io/username/image,push=true

OPTIONS:
   --output value, -o value  Define exports for build result, e.g.--output type=image,name=docker.io/username/image,push=true
   --progress value          Set type of progress (auto, plain, tty).Use plain to show container output (default: "auto")
   --trace value             Path to trace file.Defaults to no tracing.
   --local value             Allow build access to the local directory
   --frontend value          Define frontend used for build
   --opt value               Define custom options for frontend, e.g.--opt target=foo --opt build-arg:foo=bar
   --no-cache                Disable cache for all the vertices
   --export-cache value      Export build cache, e.g.--export-cache type=registry,ref=example.com/foo/bar, or --export-cache type=local,dest=path/to/dir
   --import-cache value      Import build cache, e.g.--import-cache type=registry,ref=example.com/foo/bar, or --import-cache type=local,src=path/to/dir
   --secret value            Secret value exposed to the build.Format id=secretname,src=filepath
   --allow value             Allow extra privileged entitlement, e.g.network.host, security.insecure
   --ssh value               Allow forwarding SSH agent to the builder.Format default|<id>[=<socket>|<key>[,<key>]]

```

Here's what I ran to get the equivalent of the Docker command with the `DOCKER_BUILDKIT` override:

这是我运行以获取等效于带有“DOCKER_BUILDKIT”覆盖的 Docker 命令的内容：

```
sudo -E buildctl build --frontend dockerfile.v0 \
 --local context=./build/build-test/ \
 --local dockerfile=./build/build-test/ \
 --output type=image,name=docker.io/alexellis2/build-test:latest,push=true

```

Before running this command, you'll need to run `docker login`, or to create $HOME/.docker/config.json\` with a valid set of unencrypted credentials.

在运行此命令之前，您需要运行 `docker login`，或者使用一组有效的未加密凭据创建 $HOME/.docker/config.json\`。

You'll see a nice ASCII animation for this build.

您将看到此版本的漂亮 ASCII 动画。

![buildkit-stand-alone](http://blog.alexellis.io/content/images/2020/01/buildkit-stand-alone.png)

### Build with `img` and buildkit

### 使用 `img` 和 buildkit 构建

Since I've never used `img` and haven't really heard of it being used a lot with teams vs the more common options I thought I'd give it a shot.

因为我从来没有使用过 `img`，也没有真正听说过它在团队中被大量使用，而不是更常见的选项，我想我会试一试。

First impressions are that multi-arch is not a priority and given the age of the project, may be unlikely to land. There is no binary for armhf or ARM64.

第一印象是多拱不是优先考虑的项目，考虑到项目的年龄，可能不太可能落地。 armhf 或 ARM64 没有二进制文件。

For `x86_64` the latest version is `v0.5.7` from 7 May 2019, built with Go 1.11, with Go 1.13 being the current release:

对于 `x86_64`，最新版本是 2019 年 5 月 7 日的 `v0.5.7`，使用 Go 1.11 构建，当前版本为 Go 1.13：

```sh
sudo curl -fSL "https://github.com/genuinetools/img/releases/download/v0.5.7/img-linux-amd64" -o "/usr/local/bin/img" \
    && sudo chmod a+x "/usr/local/bin/img"

```

The build options look like a subset of `buildctl`:

构建选项看起来像是`buildctl`的一个子集：

```sh
img build --help
Usage: img build [OPTIONS] PATH

Build an image from a Dockerfile.

Flags:

  -b, --backend  backend for snapshots ([auto native overlayfs]) (default: auto)
  --build-arg    Set build-time variables (default: [])
  -d, --debug    enable debug logging (default: false)
  -f, --file     Name of the Dockerfile (Default is 'PATH/Dockerfile') (default: <none>)
  --label        Set metadata for an image (default: [])
  --no-cache     Do not use cache when building the image (default: false)
  --no-console   Use non-console progress UI (default: false)
  --platform     Set platforms for which the image should be built (default: [])
  -s, --state    directory to hold the global state (default: /home/alex/.local/share/img)
  -t, --tag      Name and optionally a tag in the 'name:tag' format (default: [])
  --target       Set the target build stage to build (default: <none>)

```

Here's what we need to do a build:

这是我们需要做的构建：

```sh
sudo img build -f ./build/build-test/Dockerfile -t alexellis2/build-test:latest ./build/build-test/

```

Now for one reason or another, `img` actually failed to do a successful build. It may be due to some of the optimizations to attempt to run as non-root.

现在由于某种原因，`img` 实际上未能成功构建。这可能是由于某些优化尝试以非 root 用户身份运行。

![fail-build](http://blog.alexellis.io/content/images/2020/01/fail-build.png)

```sh
fatal error: unexpected signal during runtime execution
[signal SIGSEGV: segmentation violation code=0x1 addr=0xe5 pc=0x7f84d067c420]

runtime stack:
runtime.throw(0xfa127f, 0x2a)
    /home/travis/.gimme/versions/go1.11.10.linux.amd64/src/runtime/panic.go:608 +0x72
runtime.sigpanic()
    /home/travis/.gimme/versions/go1.11.10.linux.amd64/src/runtime/signal_unix.go:374 +0x2f2

goroutine 529 [syscall]:
runtime.cgocall(0xc9d980, 0xc00072d7d8, 0x29)
    /home/travis/.gimme/versions/go1.11.10.linux.amd64/src/runtime/cgocall.go:128 +0x5e fp=0xc00072d7a0 sp=0xc00072d768 pc=0x4039ee
os/user._Cfunc_mygetgrgid_r(0x2a, 0xc000232260, 0x7f84a40008c0, 0x400, 0xc0004ba198, 0xc000000000)

```

There seemed to be [three similar issues](https://github.com/genuinetools/img/issues/272) open.

似乎有 [三个类似问题](https://github.com/genuinetools/img/issues/272) 打开。

### Build with Kaniko

### 使用 Kaniko 构建

Kaniko is Google's container builder which aims to sandbox container builds. You can use it as a one-shot container, or as a stand-alone binary.

Kaniko 是谷歌的容器构建器，旨在沙盒容器构建。您可以将其用作一次性容器，也可以用作独立的二进制文件。

I took a look at the build [in this blog post](https://blog.alexellis.io/quick-look-at-google-kaniko/)

我查看了构建 [在这篇博文中](https://blog.alexellis.io/quick-look-at-google-kaniko/)

```sh
docker run -v $PWD/build/build-test:/workspace \
 -v ~/.docker/config.json:/kaniko/config.json \
 --env DOCKER_CONFIG=/kaniko \
gcr.io/kaniko-project/executor:latest \
 -d alexellis2/build-test:latest

```

- The flag`-d` specifies where the image should be pushed after a successful build.
- The`-v` flag is bind-mounting the current directory into the Kaniko container, it also adds your `config.json` file for pushing to a remote registry.

- 标志`-d` 指定成功构建后应将图像推送到何处。
- `-v` 标志将当前目录绑定安装到 Kaniko 容器中，它还添加了用于推送到远程注册表的 `config.json` 文件。

![kaniko](http://blog.alexellis.io/content/images/2020/01/kaniko.png)

There is some support for caching in Kaniko, but it needs manual management and preservation since Kaniko runs in a one-shot mode, rather than daemonized like Buildkit.

Kaniko 对缓存有一些支持，但它需要手动管理和保存，因为 Kaniko 以一次性模式运行，而不是像 Buildkit 那样守护进程。

### Summing up the options

### 总结选项

- Docker - traditional builder

- Docker - 传统构建器

Installing Docker can be heavy-weight and add more than expected to your system. The builder is the oldest and slowest, but gets the job done. Watch out for the networking bridge installed by Docker, it can conflict with other private networks using the same private IP range.

安装 Docker 可能是重量级的，并且会为您的系统添加超出预期的内容。建造者是最古老和最慢的，但可以完成工作。注意 Docker 安装的网桥，它可能与使用相同私有 IP 范围的其他私有网络发生冲突。

- Docker - with buildkit

- Docker - 使用 buildkit

This is the fastest option with the least amount of churn or change. It's simply enabled by prefixing the command `DOCKER_BUILDKIT=1`

这是最快的选择，最少的流失或变化。只需在命令“DOCKER_BUILDKIT=1”前加上前缀即可启用

- Stand-alone buildkit

- 独立构建工具包

This option is great for in-cluster builds, or a system that doesn't need Docker such as a CI box or runner. It does need a Linux host and there's no good experience for using it on MacOS, perhaps by running an additional VM or host and accessing over TCP?

此选项非常适合集群内构建，或不需要 Docker 的系统，例如 CI 盒或运行器。它确实需要一个 Linux 主机，但在 MacOS 上使用它没有很好的经验，也许是通过运行额外的 VM 或主机并通过 TCP 访问？

I also wanted to include a presentation by [Akihiro Suda]( [https://twitter.com/@](https://twitter.com/@) _AkihiroSuda_ /), a buildkit maintainer from NTT, Japan. This information is around 2 years old but provides another high-level overview from the landscape in 2018 [Comparing Next-Generation\
\
Container Image Building Tools](https://events19.linuxfoundation.org/wp-content/uploads/2017/11/Comparing-Next-Generation-Container-Image-Building-Tools-OSS-Akihiro-Suda.pdf)

我还想包括 [Akihiro Suda]( [https://twitter.com/@](https://twitter.com/@) _AkihiroSuda_ /) 的演讲，他是来自日本 NTT 的 buildkit 维护者。此信息已有大约 2 年的历史，但提供了 2018 年景观的另一个高级概述 [比较下一代\
\
容器镜像构建工具](https://events19.linuxfoundation.org/wp-content/uploads/2017/11/Comparing-Next-Generation-Container-Image-Building-Tools-OSS-Akihiro-Suda.pdf)

This is the best option for [faasd users](https://github.com/alexellis/faasd), where users rely only on containerd and CNI, rather than Docker or Kubernetes.

这是 [faasd 用户](https://github.com/alexellis/faasd) 的最佳选择，其中用户仅依赖 containerd 和 CNI，而不是 Docker 或 Kubernetes。

- Kaniko

- 卡尼科

The way we used Kaniko still required Docker to be installed, but provided another option.

我们使用 Kaniko 的方式仍然需要安装 Docker，但提供了另一种选择。

## Wrapping up

##  包起来

You can either use your normal container builder with OpenFaaS, or `faas-cli build --shrinkwrap` and pass the build-context along to your preferred tooling.

您可以使用带有 OpenFaaS 的普通容器构建器，或者使用 `faas-cli build --shrinkwrap` 并将构建上下文传递给您首选的工具。

Here's examples for the following tools for building OpenFaaS containers:

以下是用于构建 OpenFaaS 容器的以下工具的示例：

- [Google Cloud Build](https://www.openfaas.com/blog/openfaas-cloudrun/)
- [GitHub Actions](https://lucasroesler.com/2019/09/action-packed-functions/)
- [Jenkins](https://docs.openfaas.com/reference/cicd/jenkins/) and
- [GitLab CI](https://docs.openfaas.com/reference/cicd/gitlab/). 

- [谷歌云构建](https://www.openfaas.com/blog/openfaas-cloudrun/)
- [GitHub 操作](https://lucasroesler.com/2019/09/action-packed-functions/)
- [詹金斯](https://docs.openfaas.com/reference/cicd/jenkins/) 和
- [GitLab CI](https://docs.openfaas.com/reference/cicd/gitlab/)。

In [OpenFaaS Cloud](https://docs.openfaas.com/openfaas-cloud/architecture/). we provide a complete hands-off CI/CD experience using the shrinkwrap approach outlined in this post and the buildkit daemon. For all other users I would recommend using Docker, or Docker with buildkit.

在 [OpenFaaS Cloud](https://docs.openfaas.com/openfaas-cloud/architecture/) 中。我们使用本文中概述的收缩包装方法和 buildkit 守护程序提供完整的 CI/CD 体验。对于所有其他用户，我建议使用 Docker，或 Docker 和 buildkit。

You can [build your own self-hosted OpenFaaS Cloud](https://www.openfaas.com/blog/ofc-private-cloud/) environment with GitHub or GitLab integration.

您可以[构建您自己的自托管 OpenFaaS 云](https://www.openfaas.com/blog/ofc-private-cloud/) 环境与 GitHub 或 GitLab 集成。

For [faasd users](https://github.com/openfaas/faasd), you only have containerd installed on your host instead of `docker`, so the best option for you is to download buildkit.

对于 [faasd 用户](https://github.com/openfaas/faasd)，您的主机上只安装了 containerd 而不是 docker，因此您最好的选择是下载 buildkit。

If you are interested in what Serverless Functions are and what they can do for you, why not checkout my new eBook and video workshop on Gumroad?

如果您对无服务器功能是什么以及它们能为您做什么感兴趣，为什么不在 Gumroad 上查看我的新电子书和视频研讨会？

- [Checkout Serverless For Everyone](https://gumroad.com/l/serverless-for-everyone-else)

- [Checkout Serverless For Everyone](https://gumroad.com/l/serverless-for-everyone-else)

We did miss out one of the important parts of the workflow in this post, the deployment. Any OCI container can be deployed to the OpenFaaS control-plane on top of Kubernetes as long as its [conforms to the serverless workload definition](https://docs.openfaas.com/reference/workloads/). If you'd like to see the full experience of build, push and deploy, check out the [OpenFaaS workshop](https://github.com/openfaas/workshop/).

我们确实错过了这篇文章中工作流程的重要部分之一，即部署。任何 OCI 容器都可以部署到 Kubernetes 之上的 OpenFaaS 控制平面，只要其 [符合无服务器工作负载定义](https://docs.openfaas.com/reference/workloads/)。如果您想了解构建、推送和部署的完整体验，请查看 [OpenFaaS 研讨会](https://github.com/openfaas/workshop/)。

## Wrapping up

##  包起来

### Get help with Cloud Native, Docker, Go, CI & CD, or Kubernetes

### 获取有关云原生、Docker、Go、CI 和 CD 或 Kubernetes 的帮助

Could you use some help with a difficult problem, an external view on a new idea or project? Perhaps you would like to build a technology proof of concept before investing more? Get in touch via [alex@openfaas.com](mailto:alex@openfaas.com) or book a session with me on [calendly.com/alexellis](https://calendly.com/alexellis/). 

您能否在解决难题时使用一些帮助，或对新想法或项目的外部看法？也许您想在投资更多之前建立一个技术概念证明？通过 [alex@openfaas.com](mailto:alex@openfaas.com) 与我联系或在 [calendly.com/alexellis](https://calendly.com/alexellis/) 上与我预约。

