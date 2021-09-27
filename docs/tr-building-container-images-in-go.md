# Building container images in Go

  # 在 Go 中构建容器镜像

 [Ahmet Alp Balkan](https://twitter.com/ahmetb)  published on        03 February 2021

[Ahmet Alp Balkan](https://twitter.com/ahmetb) 发布于 2021 年 2 月 3 日

In this article, I’ll explain how to build OCI container images without using Docker by building the layers and image manifests programmatically using the [go-containerregistry](https://github.com/google/go-containerregistry) module. As an example, I'll build a container image by adding some static website content on top of the [`nginx`](https://hub.docker.com/_/nginx) image and push it to a registry like ` gcr.io` using a Go program.

在本文中，我将通过使用 [go-containerregistry](https://github.com/google/go-containerregistry) 模块以编程方式构建层和图像清单来解释如何在不使用 Docker 的情况下构建 OCI 容器图像。例如，我将通过在 [`nginx`](https://hub.docker.com/_/nginx) 镜像之上添加一些静态网站内容来构建一个容器镜像，并将其推送到像 ` gcr.io` 使用 Go 程序。

The procedure will look like this:

该过程将如下所示：

1. Pull the `nginx` image from Docker Hub
2. Create new layer that deletes the default `/usr/share/nginx/html` directory.
3. Create new layer with static HTML contents and assets.
4. Append new layers to the image and tag.
5. Push the new image to the registry.

1. 从 Docker Hub 拉取 `nginx` 镜像
2.创建删除默认`/usr/share/nginx/html`目录的新层。
3. 使用静态 HTML 内容和资产创建新层。
4. 将新图层附加到图像和标签。
5. 将新镜像推送到注册表。

You can find the example code of this exercise [in this gist](https://gist.github.com/ahmetb/430baa4e8bb0b0f78abb1c34934cd0b6). Let’s dive in.

您可以在 [this gist](https://gist.github.com/ahmetb/430baa4e8bb0b0f78abb1c34934cd0b6) 中找到本练习的示例代码。让我们深入了解。

Download the [module](https://pkg.go.dev/github.com/google/go-containerregistry):

下载 [模块](https://pkg.go.dev/github.com/google/go-containerregistry)：

```sh
go get -u github.com/google/go-containerregistry
```

Pull an image reference. This method resolves `nginx` reference to the `index.docker.io/library/nginx:latest` and then negotiates anonymous credentials from Docker Hub, and returns a [`v1.Image`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1#Image) which is actually a [`remote.Image`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/remote#Image):

拉取图像参考。此方法解析对 `index.docker.io/library/nginx:latest` 的 `nginx` 引用，然后从 Docker Hub 协商匿名凭据，并返回 [`v1.Image`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1#Image) 这实际上是一个 [`remote.Image`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/remote#Image)：

```go
img, err := crane.Pull("nginx")
if err != nil {
    panic(err)
}
```

Now, let's create a layer that uses [whiteout files](https://github.com/opencontainers/image-spec/blob/79b036d80240ae530a8de15e1d21c7ab9292c693/layer.md#whiteouts) to remove the `/usr/share/nginx/html` directory that comes with the nginx image.

现在，让我们创建一个使用 [whiteout files](https://github.com/opencontainers/image-spec/blob/79b036d80240ae530a8de15e1d21c7ab9292c693/layer.md#whiteouts) 的层来删除`/usr/share/nginx/html` nginx 映像附带的目录。

To do that, we use a helper method that lets us create tarballs from a list of file names and in-memory byte slices. We need a file named `usr/share/nginx/.wh.html` in the tar file to clear this path in this layer:

为此，我们使用了一个辅助方法，该方法允许我们从文件名列表和内存中的字节片创建 tarball。我们需要在tar文件中有一个名为`usr/share/nginx/.wh.html`的文件来清除这一层的这个路径：

```go
deleteMap := map[string][]byte{
    "usr/share/nginx/.wh.html": []byte{},
}
deleteLayer, err := crane.Layer(deleteMap)
if err != nil {
    panic(err)
}
```

Now, we need to scan the directory tree that contains the static HTML files and assets that we want to add to this container image. We can again use the `crane.Layer` method, but that requires you to read all files into the memory.

现在，我们需要扫描包含要添加到此容器映像的静态 HTML 文件和资产的目录树。我们可以再次使用 `crane.Layer` 方法，但这需要您将所有文件读入内存。

Here, we can also shell out to the `tar` command to to create this tarball and print the results to the stdout (which we then read and pass to [`tarball.FromReader`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/tarball#LayerFromReader). This command would look something along the lines of:

在这里，我们还可以使用 `tar` 命令来创建这个 tarball 并将结果打印到标准输出（然后我们读取并传递给 [`tarball.FromReader`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/tarball#LayerFromReader)。这个命令看起来类似于：

```sh
tar -cf- DIR \
    --transform 's,^,usr/share/nginx/,'
    --owner=0 --group=0
```

Or we can natively build tarballs using `tar.Writer` and write the result into an in-memory buffer like we do in the [gist](https://gist.github.com/ahmetb/430baa4e8bb0b0f78abb1c34934cd0b6). Here, we scan the files in the directory tree using the `filepath.Walk` method, and we add directory and file entries into the tar archive. As a shortcut, I only implemented directories and regular files (symlinks etc are left as an exercise to the reader). Note that we also add a `usr/share/nginx/html` prefix to the file entries.

或者我们可以使用 `tar.Writer` 本地构建 tarball，并将结果写入内存缓冲区，就像我们在 [gist](https://gist.github.com/ahmetb/430baa4e8bb0b0f78abb1c34934cd0b6) 中所做的那样。在这里，我们使用 `filepath.Walk` 方法扫描目录树中的文件，并将目录和文件条目添加到 tar 存档中。作为一种快捷方式，我只实现了目录和常规文件（符号链接等留给读者作为练习)。请注意，我们还在文件条目中添加了一个 `usr/share/nginx/html` 前缀。

Then, we append these layers into a new image:

然后，我们将这些层附加到一个新图像中：

```go
newImg, err := mutate.AppendLayers(img, deleteLayer, addLayer)
if err != nil {
    panic(err)
}
```

This is also where you can change entrypoint and arguments of the image.

这也是您可以更改图像的入口点和参数的地方。

Then, we tag the image:

然后，我们标记图像：

```go
tag, err := name.NewTag("gcr.io/ahmetb-blog/blog:latest")
if err != nil {
    panic(err)
}
```

At this point we can either push the image to a remote registry (using local credential keychain and helpers), or load into a local Docker daemon for testing:

此时，我们可以将映像推送到远程注册表（使用本地凭据钥匙串和帮助程序），或者加载到本地 Docker 守护程序中进行测试：

```go
// for local testing, load into local docker engine
if s, err := daemon.Write(tag, newImg);err != nil {
    panic(err)
} else {
    fmt.Println("pushed "+s)
}

// push to remote registry
if err := crane.Push(newImg, tag.String());err != nil {
    panic(err)
} else {
    fmt.Println(s)
}
```

So that’s it. I hope this was a nice exercise to give you an idea what [go-containerregistry](https://github.com/google/go-containerregistry) can do for you. It has a lot more capabilities, such the [`mutate`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/mutate) package to modify manifests, rebase layers, flatten images. (Did you know tools like [`ko`](https://github.com/google/ko) and [`crane`](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane.md) are built using this Go module?)

就是这样了。我希望这是一个很好的练习，可以让您了解 [go-containerregistry](https://github.com/google/go-containerregistry) 可以为您做什么。它有更多的功能，例如 [`mutate`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/mutate) 包来修改清单、rebase 层、扁平化图片。 （你知道像 [`ko`](https://github.com/google/ko) 和 [`crane`](https://github.com/google/go-containerregistry/blob/main/cmd) 这样的工具吗？ /crane/doc/crane.md) 是使用这个 Go 模块构建的？)

Make sure to [star the repository](https://github.com/google/go-containerregistry) and follow maintainers [@jonjohnsonjr](https://twitter.com/jonjonsonjr), [@ImJasonH](https://twitter.com/imjasonh) and [@mattomata](https://twitter.com/mattomata) on Twitter to stay in the loop.

确保 [star the repository](https://github.com/google/go-containerregistry) 并关注维护者 [@jonjohnsonjr](https://twitter.com/jonjonsonjr), [@ImJasonH](https://twitter.com/imjasonh) 和 [@mattomata](https://twitter.com/mattomata) 在 Twitter 上保持循环。

# Ahmet Alp Balkan

  # 艾哈迈德阿尔卑斯巴尔干

 I am a software engineer at Google, working on      [Cloud Run](https://cloud.run),      and [Kubernetes](https://kubernetes.io). I focus on improving developer experiences and explaining      complex features in simple words. I’ve created developer      tools like [Krew](https://krew.dev) and      [kubectx](https://github.com/ahmetb/kubectx). You can [       follow me on Twitter](https://twitter.com/ahmetb). 

我是 Google 的一名软件工程师，负责 [Cloud Run](https://cloud.run) 和 [Kubernetes](https://kubernetes.io)。我专注于改善开发人员体验并用简单的语言解释复杂的功能。我创建了像[Krew](https://krew.dev) 和 [kubectx](https://github.com/ahmetb/kubectx) 这样的开发者工具。你可以 [在 Twitter 上关注我](https://twitter.com/ahmetb)。

