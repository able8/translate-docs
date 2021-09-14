# How Docker Build Works Internally

May 25, 2020 (Updated: August 7, 2021)

[Containers,](http://iximiuz.com/en/categories/?category=Containers) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

_You need containers to build images. Yes, you've heard it right. Not another way around._

For people who found their way to containers through Docker (well, most of us I believe) it may seem like _images_ are of somewhat primary nature. We've been taught to start from a _Dockerfile_, build an image using that file, and only then run a container from that image. Alternatively, we could run a container specifying an image from a registry, yet the main idea persists - an image comes first, and only then the container.

**But what if I tell you that the actual workflow is reverse?** Even when you are building your very first image using Docker, podman, or buildah, you are already, albeit implicitly, running containers under the hood!

## How container images are created

Let's avoid any unfoundedness and take a closer look at the image building procedure. The easiest way to spot this behavior is to build a simple image using the following _Dockerfile_:

```dockerfile
FROM debian:latest

RUN sleep 2 && apt-get update
RUN sleep 2 && apt-get install -y uwsgi
RUN sleep 2 && apt-get install -y python3

COPY some_file /

```

While building the image, try running `docker stats -a` in another terminal:

<img src="/you-need-containers-to-build-an-image/docker-build.gif" title="Your browser does not support the  tag" alt="Running docker build and docker stats in different terminals.

">



_Running `docker build` and `docker stats` in different terminals._

Huh, we haven't been launching any containers ourselves, nevertheless, `docker stats` shows that there were 3 of them 🙈 But how come?

Simplifying a bit, [images](https://github.com/opencontainers/image-spec) can be seen as archives with a filesystem inside. Additionally, they may also contain some configurational data like a default command to be executed when a container starts, exposed ports, etc, but we will be mostly focusing on the filesystem part. Luckily, we already know, that technically [images aren't required to run containers](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/). Unlike virtual machines, containers are just [isolated and restricted processes](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/) on your Linux host. They do form an isolated execution environment, including the personalized root filesystem, but the bare minimum to start a container is just a folder with a single executable file inside. So, when we are starting a container from an image, the image gets unpacked and its content is provided to the [container runtime](https://github.com/opencontainers/runtime-spec) in a form of a [filesystem bundle](https://github.com/opencontainers/runtime-spec/blob/44341cdd36f6fee6ddd73e602f9e3eca1466052f/bundle.md), i.e. a regular directory containing the future root filesystem files and some configs (all those layers you may have started thinking about are abstracted away by a [union mount](https://en.wikipedia.org/wiki/Union_mount) driver like [overlay fs](https://dev.to/napicella/how-are-docker-images-built-a-look-into-the-linux-overlay-file-systems-and-the-oci-specification-175n)). Thus, if you don't have the image but you do need the `alpine` Linux distro as the execution environment, you always can grab Alpine's rootfs ( [2.6 MB](https://github.com/alpinelinux/docker-alpine/raw/c5510d5b1d2546d133f7b0938690c3c1e2cd9549/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz)) and put it to a regular folder on your disk, then mix in your application files, feed it to the container runtime and call it a day.

However, to unleash the full power of containers, we need handy image building facilities. Historically, Dockerfiles have been serving this purpose. Any Dockerfile must have the `FROM` instruction at the very beginning. This instruction specifies the base image while the rest of the Dockerfile describes the difference between the base and the derived (i.e. current) images.

The most basic container image is a so-called [`scratch`](https://hub.docker.com/_/scratch) image. It corresponds to an empty folder and the `FROM scratch` instruction in a Dockerfile means _noop_.

Now, let's take a look at the beloved `alpine` image:

```dockerfile
# https://github.com/alpinelinux/docker-alpine/blob/v3.11/x86_64/Dockerfile

FROM scratch
ADD alpine-minirootfs-3.11.6-x86_64.tar.gz /
CMD ["/bin/sh"]

```

I.e. to make the Alpine Linux distro image we just need to copy its root filesystem to an empty folder ( _scratch_ image) and that's it! Well, I bet Dockerfiles you've seen so far a rarely that trivial. More often than not, we need to utilize distro's facilities to prepare the file system of the future container and one of the most common examples is probably when we need to pre-install some external packages using `yum`, `apt`, or `apk`:

```dockerfile
FROM debian:latest

RUN apt-get install -y ca-certificates

```

But how can we have `apt` running if we are building this image, say, on a Fedora host? Containers to the rescue! Every time, Docker (or buildah, or podman, etc) encounters a `RUN` instruction in the Dockerfile it actually fires a new container! The bundle for this container is formed by the base image plus all the changes made by preceding instructions from the Dockerfile (if any). When the execution of the `RUN` step completes, all the changes made to the container's filesystem (so-called _diff_) become a new _layer_ in the image being built and the process repeats starting from the next Dockerfile instruction.

![Building image using Dockerfile](http://iximiuz.com/you-need-containers-to-build-an-image/kdpv.png)

_Building image using Dockerfile_

Getting back to the original example from the beginning of this article, a mind-reader could have noticed that the number of containers we've seen in the second terminal corresponded exactly to the number of `RUN` instructions in the Dockerfile. For people with a solid understanding of the internal kitchen of containers it may sound obvious. However, for the rest of us possessing rather hands-on containers experience, the Dockerfile-based (and immensely popularized) workflow may instead obscure some things.

Luckily, even though Dockerfiles are a de facto standard to describe images, it's not the only existing way. Thus, when using Docker, we can [`commit`](https://docs.docker.com/engine/reference/commandline/commit/) any running container to produce a new image. All the changes made to the container's filesystem by any command run inside of it since its start will form the topmost layer of the image created by the _commit_, while the base will be taken from the image used to create the said container. Although, if we decide to use this method we may face some reproducibility problems.

## How to build image without Dockerfile

Interesting enough, that some of the novel image building tools consider Dockerfile not as an advantage, but as a limitation. For instance, _buildah_ [promotes an alternative command-line image building procedure](https://www.redhat.com/sysadmin/building-buildah):

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

Finalizing, let's take a brief look at two other prominent tools - Google's [kaniko](https://github.com/GoogleContainerTools/kaniko) and Uber's [makisu](https://github.com/uber/makisu). They try to tackle the image building problem from a slightly different angle. These tools don't really run containers while building images. Instead, they directly modify the local filesystem while following image building instructions 🤯 I.e. if you accidentally start such a tool on your laptop, it's highly likely that your host OS will be wiped and replaced with the rootfs of the image. So, beware. Apparently, these tools are supposed to be executed fully inside of an already existing container. This solves some security problems bypassing the need for elevating privileges of the builder process. Nevertheless, while the technique itself is very different from the traditional Docker's or buildah's approaches, the containers are still there. The main difference is that they have been moved out of the scope of the build tool.

## Instead of conclusion

The concept of container images turned out to be very handy. The layered image structure in conjunction with union mounts like overlay fs made the storage and usage of images immensely efficient. The declarative Dockerfile-based approach enabled the reproducible and cachable building of artifacts. This allowed the idea of container images to become so wide-spread that sometime it may seem like it's an inalienable and archetypal part of the containerverse. However, as we saw in the article, from the implementation standpoint, containers are independent of images. Instead, most of the time we need containers to build images, not vice-a-verse.

Make code, not war!

### Appendix:  Image Building Tools

- [Docker](https://github.com/docker/docker-ce)
- [Podman](https://github.com/containers/libpod) & [Buildah](https://github.com/containers/buildah) / [intro](https://www.giantswarm.io/blog/building-container-images-with-podman-and-buildah)
- [BuildKit](https://github.com/moby/buildkit) / [intro](https://www.giantswarm.io/blog/container-image-building-with-buildkit)
- [img](https://github.com/genuinetools/img) / [intro](https://www.giantswarm.io/blog/building-container-images-with-img)
- [kaniko](https://github.com/GoogleContainerTools/kaniko) / [intro](https://www.giantswarm.io/blog/container-image-building-with-kaniko)
- [makisu](https://github.com/uber/makisu) / [intro](https://www.giantswarm.io/blog/container-image-building-with-makisu)

### Related articles

- [Containers aren't Linux processes](http://iximiuz.com/en/posts/oci-containers/)
- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [A journey from containerization to orchestration and beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)

[linux,](javascript: void 0) [container,](javascript: void 0) [docker,](javascript: void 0) [podman,](javascript: void 0) [buildah,](javascript: void 0) [kaniko,](javascript: void 0) [makisu](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

