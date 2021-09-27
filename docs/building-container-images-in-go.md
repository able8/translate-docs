# Building container images in Go

 [Ahmet Alp Balkan](https://twitter.com/ahmetb)  published on        03 February 2021   

In this article, I’ll explain how to build OCI container images without using Docker by building the layers and image manifests programmatically using the [go-containerregistry](https://github.com/google/go-containerregistry) module. As an example, I’ll build a container image by adding some static website content on top of the [`nginx`](https://hub.docker.com/_/nginx) image and push it to a registry like `gcr.io` using a Go program.

The procedure will look like this:

1. Pull the `nginx` image from Docker Hub
2. Create new layer that deletes the default `/usr/share/nginx/html` directory.
3. Create new layer with static HTML contents and assets.
4. Append new layers to the image and tag.
5. Push the new image to the registry.

You can find the example code of this exercise [in this gist](https://gist.github.com/ahmetb/430baa4e8bb0b0f78abb1c34934cd0b6). Let’s dive in.

Download the [module](https://pkg.go.dev/github.com/google/go-containerregistry):

```sh
go get -u github.com/google/go-containerregistry
```

Pull an image reference. This method resolves `nginx` reference to the `index.docker.io/library/nginx:latest` and then negotiates anonymous credentials from Docker Hub, and returns a [`v1.Image`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1#Image) which is actually a [`remote.Image`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/remote#Image):

```go
img, err := crane.Pull("nginx")
if err != nil {
    panic(err)
}
```

Now, let’s create a layer that uses [whiteout files](https://github.com/opencontainers/image-spec/blob/79b036d80240ae530a8de15e1d21c7ab9292c693/layer.md#whiteouts) to remove the `/usr/share/nginx/html` directory that comes with the nginx image.

To do that, we use a helper method that lets us create tarballs from a list of file names and in-memory byte slices. We need a file named `usr/share/nginx/.wh.html` in the tar file to clear this path in this layer:

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

Here, we can also shell out to the `tar` command to to create this tarball and print the results to the stdout (which we then read and pass to [`tarball.FromReader`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/tarball#LayerFromReader). This command would look something along the lines of:

```sh
tar -cf- DIR \
    --transform 's,^,usr/share/nginx/,'
    --owner=0 --group=0
```

Or we can natively build tarballs using `tar.Writer` and write the result into an in-memory buffer like we do in the [gist](https://gist.github.com/ahmetb/430baa4e8bb0b0f78abb1c34934cd0b6). Here, we scan the files in the directory tree using the `filepath.Walk` method, and we add directory and file entries into the tar archive. As a shortcut, I only implemented directories and regular files (symlinks etc are left as an exercise to the reader). Note that we also add a `usr/share/nginx/html` prefix to the file entries.

Then, we append these layers into a new image:

```go
newImg, err := mutate.AppendLayers(img, deleteLayer, addLayer)
if err != nil {
    panic(err)
}
```

This is also where you can change entrypoint and arguments of the image.

Then, we tag the image:

```go
tag, err := name.NewTag("gcr.io/ahmetb-blog/blog:latest")
if err != nil {
    panic(err)
}
```

At this point we can either push the image to a remote registry (using local credential keychain and helpers), or load into a local Docker daemon for testing:

```go
// for local testing, load into local docker engine
if s, err := daemon.Write(tag, newImg); err != nil {
    panic(err)
} else {
    fmt.Println("pushed "+s)
}

// push to remote registry
if err := crane.Push(newImg, tag.String()); err != nil {
    panic(err)
} else {
    fmt.Println(s)
}
```

So that’s it. I hope this was a nice exercise to give you an idea what [go-containerregistry](https://github.com/google/go-containerregistry) can do for you. It has a lot more capabilities, such the [`mutate`](https://pkg.go.dev/github.com/google/go-containerregistry/pkg/v1/mutate) package to modify manifests, rebase layers, flatten images. (Did you know tools like [`ko`](https://github.com/google/ko) and [`crane`](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane.md) are built using this Go module?)

Make sure to [star the repository](https://github.com/google/go-containerregistry) and follow maintainers [@jonjohnsonjr](https://twitter.com/jonjonsonjr), [@ImJasonH](https://twitter.com/imjasonh) and [@mattomata](https://twitter.com/mattomata) on Twitter to stay in the loop.

# Ahmet Alp Balkan

 I am a software engineer at Google, working on      [Cloud Run](https://cloud.run),      and [Kubernetes](https://kubernetes.io).      I focus on improving developer experiences and explaining      complex features in simple words. I’ve created developer      tools like [Krew](https://krew.dev) and      [kubectx](https://github.com/ahmetb/kubectx).      You can [       follow me on Twitter](https://twitter.com/ahmetb).  

