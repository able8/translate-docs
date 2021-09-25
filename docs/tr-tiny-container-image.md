# Tiny Container Challenge: Building a 6kB Containerized HTTP Server!

# 微型容器挑战：构建 6kB 容器化 HTTP 服务器！

##### April 14, 2021

##### 2021 年 4 月 14 日

**TL;DR** I set out to build the smallest  container image that I could that was still able to do something useful. By taking advantage of multistage builds, the `scratch` base image, and a tiny assembly based http server, I was able to get it down to 6.32kB!

**TL;DR** 我着手构建最小的容器镜像，它仍然能够做一些有用的事情。通过利用多阶段构建、`scratch` 基础映像和基于小型程序集的 http 服务器，我能够将其减小到 6.32kB！

------

Table of Contents:

目录：

- [Bloated Containers](https://devopsdirective.com/posts/2021/04/tiny-container-image/#bloated-containers)
- [Challenge](https://devopsdirective.com/posts/2021/04/tiny-container-image/#challenge)
- [Naive Solution](https://devopsdirective.com/posts/2021/04/tiny-container-image/#naive-solution)
- [Smaller Base Image](https://devopsdirective.com/posts/2021/04/tiny-container-image/#smaller-base-image)
- [Compiled Languages](https://devopsdirective.com/posts/2021/04/tiny-container-image/#compiled-languages)
- [Multi-stage Builds](https://devopsdirective.com/posts/2021/04/tiny-container-image/#multi-stage-builds)
- [Static Compilation + Scratch image](https://devopsdirective.com/posts/2021/04/tiny-container-image/#static-compilation--scratch-image)
- [ASM for the Win!](https://devopsdirective.com/posts/2021/04/tiny-container-image/#asm-for-the-win)

- [膨胀的容器](https://devopsdirective.com/posts/2021/04/tiny-container-image/#bloated-containers)
- [挑战](https://devopsdirective.com/posts/2021/04/tiny-container-image/#challenge)
- [天真解决方案](https://devopsdirective.com/posts/2021/04/tiny-container-image/#naive-solution)
- [较小的基础图像](https://devopsdirective.com/posts/2021/04/tiny-container-image/#smaller-base-image)
- [编译语言](https://devopsdirective.com/posts/2021/04/tiny-container-image/#compiled-languages)
- [多阶段构建](https://devopsdirective.com/posts/2021/04/tiny-container-image/#multi-stage-builds)
- [静态编译 + 抓图](https://devopsdirective.com/posts/2021/04/tiny-container-image/#static-compilation--scratch-image)
- [ASM for the Win！](https://devopsdirective.com/posts/2021/04/tiny-container-image/#asm-for-the-win)

------

[![images/tiny-container-thumbnail.png](https://devopsdirective.com/posts/2021/04/tiny-container-image/images/tiny-container-thumbnail.png)](https://www.youtube.com/watch?v=VG8rZIE8ET8) If you prefer video format, check out the corresponding YouTube video here!

www.youtube.com/watch?v=VG8rZIE8ET8) 如果您更喜欢视频格式，请在此处查看相应的 YouTube 视频！

## Bloated Containers

## 臃肿的容器

Containers are often touted as a silver bullet to solve every  challenge associated with operating software. While I love containers, I often encounter container images in the wild that have a variety of  issues. One common issue is size, with images sometimes clocking in at  multiple GB!

容器通常被吹捧为解决与操作软件相关的所有挑战的灵丹妙药。虽然我喜欢容器，但我经常在野外遇到容器镜像有各种问题。一个常见问题是大小，图像有时会以多 GB 的速度记录！

Because of this, I decided to challenge myself and others to build the smallest possible image.

正因为如此，我决定挑战自己和其他人，以建立尽可能小的形象。

## Challenge

##  挑战

The rules are pretty simple:

规则非常简单：

- The container should serve the contents of a file over http on the port of your choosing
- No volume mounts are allowed (Also known as “The Marek Rule” 😜)

- 容器应该在您选择的端口上通过 http 提供文件内容
- 不允许安装卷（也称为“马立克规则”😜）

## Naive Solution

## 天真的解决方案

To get a baseline image size, we can use node.js to create a simple server `index.js`:

为了获得基线图像大小，我们可以使用 node.js 创建一个简单的服务器 `index.js`：

```js
const fs = require("fs");
const http = require('http');

const server = http.createServer((req, res) => {
  res.writeHead(200, { 'content-type': 'text/html' })
  fs.createReadStream('index.html').pipe(res)
})

server.listen(port, hostname, () => {
  console.log(`Server: http://0.0.0.0:8080/`);
});
```

and built it into an image starting the official node base image:

并将其构建为启动官方节点基础镜像的镜像：

```bash
FROM node:14
COPY ..
CMD ["node", "index.js"]
```

This clocked in at `943MB`! 😳

这计时为`943MB`！ 😳

## Smaller Base Image

## 较小的基础图像

One of the simplest and most obvious tactics for reducing image size  is to use a smaller base image. The official node image has a `slim` variant (still debian based, but with fewer dependencies preinstalled) and an `alpine` variant based on [Alpine Linux](https://alpinelinux.org/).

减小图像大小的最简单和最明显的策略之一是使用较小的基础图像。官方节点映像有一个 `slim` 变体（仍然基于 debian，但预装的依赖项较少）和一个基于 [Alpine Linux](https://alpinelinux.org/) 的 `alpine` 变体。

Using `node:14-slim` and `node:14-alpine` as the base image brought the image size down to `167MB` and `116MB` respectively.

使用 `node:14-slim` 和 `node:14-alpine` 作为基础图像，将图像大小分别降低到 `167MB` 和 `116MB`。

Because docker images are additive, which each layer building on the  next there isn’t much else we can do to make the node.js solution  smaller.

因为 docker 镜像是可叠加的，每一层都在下一层构建，所以我们没有什么其他办法可以使 node.js 解决方案更小。

## Compiled Languages

## 编译语言

To take things to the next level we can move to a compiled language  with many fewer runtime dependencies. There are a variety of options,  but for building web services, [golang](https://golang.org/) is a popular choice.

为了让事情更上一层楼，我们可以转向一种运行时依赖更少的编译语言。有多种选择，但对于构建 Web 服务，[golang](https://golang.org/) 是一种流行的选择。

I created a basic fileserver `server.go`:

我创建了一个基本的文件服务器`server.go`：

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fileServer := http.FileServer(http.Dir("./"))
    http.Handle("/", fileServer)
    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil);err != nil {
            log.Fatal(err)
    }
}
```

and built it into a container image using the official golang base image:

并使用官方 golang 基础镜像将其构建到容器镜像中：

```bash
FROM golang:1.14
COPY ..
RUN go build -o server .
CMD ["./server"]
```

Which clocked in at… `818MB`. 😡

其中时钟在......`818MB`。 😡

The issue here is that the golang base image has a lot of  dependencies installed that are useful when building go software but  aren’t needed to run the software.

这里的问题是 golang 基础镜像安装了很多依赖项，这些依赖项在构建 go 软件时很有用，但在运行软件时不需要。

## Multi-stage Builds 

## 多阶段构建

Docker has a feature called [multi-stage builds](https://docs.docker.com/develop/develop-images/multistage-build/) which make it easy to build the code in an environment with all of the  necessary dependencies , and then copy the resulting executable into a  different image.

Docker 有一项称为 [多阶段构建](https://docs.docker.com/develop/develop-images/multistage-build/) 的功能，它可以轻松地在具有所有必要依赖项的环境中构建代码，然后将生成的可执行文件复制到不同的映像中。

This is useful for a variety of reasons, but one of the most obvious is image size! By refactoring the dockerfile as follows:

出于多种原因，这很有用，但最明显的原因之一是图像大小！通过如下重构dockerfile：

```bash
### build stage ###
FROM golang:1.14-alpine AS builder
COPY ..
RUN go build -o server .

### run stage ###
FROM alpine:3.12
COPY --from=builder /go/server ./server
COPY index.html index.html
CMD ["./server"]
```

The resulting image is just `13.2MB`! 🙂

结果图像只有`13.2MB`！ 🙂

## Static Compilation + Scratch image

## 静态编译+抓图

13 MB isn’t too bad, but there are still a couple of tricks we can play to make it even smaller.

13 MB 还不错，但我们仍然可以使用一些技巧使其更小。

There is a base image called [scratch](https://hub.docker.com/_/scratch) which is explicitly empty and has zero size. Because the `scratch` has nothing inside, any image built from it must bring all of its necessary dependencies.

有一个名为 [scratch](https://hub.docker.com/_/scratch) 的基本图像，它明确为空且大小为零。因为 `scratch` 里面什么都没有，所以从它构建的任何图像都必须带来所有必要的依赖项。

To make this possible with our go based server, we need to add a few  flags to the compilation step in order to ensure the necessary libraries are statically linked into the executable:

为了在我们基于 go 的服务器上实现这一点，我们需要在编译步骤中添加一些标志，以确保必要的库静态链接到可执行文件中：

```bash
### build stage ###
FROM golang:1.14 as builder
COPY ..
RUN go build \
  -ldflags "-linkmode external -extldflags -static" \
  -a server.go

### run stage ###
FROM scratch
COPY --from=builder /go/server ./server
COPY index.html index.html
CMD ["./server"]
```

Specifically, we set the link mode to `external` and pass the `-static` flag to the external linker.

具体来说，我们将链接模式设置为 `external` 并将 `-static` 标志传递给外部链接器。

These two changes bring the image size to `8.65MB` 😀

这两个变化使图像大小达到了 `8.65MB` 😀

## ASM for the Win!

## ASM 为胜利！

An image less than 10MB, written in a language like Go is plenty  small for almost any circumstance… but we can go smaller! Github user [nemasu](https://github.com/nemasu) fully functional http server written in assembly on github named [assmttpd](https://github.com/nemasu/asmttpd).

一个小于 10MB 的图像，用 Go 之类的语言编写，几乎在任何情况下都足够小……但我们可以更小！ Github 用户 [nemasu](https://github.com/nemasu) 在 github 上用汇编编写的全功能 http 服务器名为 [assmttpd](https://github.com/nemasu/asmttpd)。

All that was required to containerize it was to install a few build  dependencies into the ubuntu base image before running the provided `make release` recipe:

容器化它所需要的只是在运行提供的 `make release` 配方之前将一些构建依赖项安装到 ubuntu 基础映像中：

```bash
### build stage ###
FROM ubuntu:18.04 as builder
RUN apt update
RUN apt install -y make yasm as31 nasm binutils
COPY ..
RUN make release

### run stage ###
FROM scratch
COPY --from=builder /asmttpd /asmttpd
COPY /web_root/index.html /web_root/index.html
CMD ["/asmttpd", "/web_root", "8080"]
```

The resulting `asmttpd` executable is then copied into the scratch image and invoked with the `CMD`. This resulted in an image size of just 6.34kB! 🥳

然后将生成的 `asmttpd` 可执行文件复制到暂存映像中并使用 `CMD` 调用。这导致图像大小仅为 6.34kB！ 🥳

![images/image-sizes.png](https://devopsdirective.com/posts/2021/04/tiny-container-image/images/image-sizes.png) The progression of container image sizes!

Hopefully you enjoyed this journey from our initial 943MB Node.js  image all the way to this tiny 6.34kB Assembly image and learned some  techniques you can apply to make your container images smaller in the  future. 

希望您喜欢从我们最初的 943MB Node.js 镜像一直到这个 6.34kB 的小程序集镜像的整个过程，并学习了一些可以在未来应用以缩小容器镜像的技术。

