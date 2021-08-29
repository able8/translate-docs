# Podman and Buildah for Docker users

February 21, 2019 From: https://developers.redhat.com/blog/2019/02/21/podman-and-buildah-for-docker-users#

I was asked recently on Twitter to better explain [Podman](https://github.com/containers/libpod) and [Buildah](https://github.com/containers/libpod) for someone familiar with Docker.  Though there are many blogs and  tutorials out there, which I will list later, we in the community have  not centralized an explanation of how Docker users move from Docker to  Podman and Buildah.  Also what role does Buildah play? Is Podman  deficient in some way that we need both Podman and Buildah to replace  Docker?

This article answers those questions and shows how to migrate to Podman.

## How does Docker work?

First, let’s be clear about how Docker works; that will help us to  understand the motivation for Podman and also for Buildah. If you are a  Docker user, you understand that there is a daemon process that must be  run to service all of your Docker commands. I can’t claim to understand  the motivation behind this but I imagine it seemed like a great idea, at the time, to do all the cool things that Docker does in one place and  also provide a useful API to that process for future evolution. In the  diagram below, we can see that the Docker daemon provides all the  functionality needed to:

- Pull and push images from an image registry
- Make copies of images in a local container storage and to add layers to those containers
- Commit containers and remove local container images from the host repository
- Ask the kernel to run a container with the right namespace and cgroup, etc.

Essentially the Docker daemon does all the work with registries, images, containers, and the kernel.  The Docker command-line interface  (CLI) asks the daemon to do this on your behalf.

[![How does Docker Work -- Docker architecture overview](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig1.png)](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig1.png)

This article does not get into the detailed pros and cons of the  Docker daemon process.  There is much to be said in favor of this  approach and I can see why, in the early days of Docker, it made a lot  of sense.  Suffice it to say that there were several reasons why Docker  users were concerned about this approach as usage went up. To list a  few:

- A single process could be a single point of failure.
- This process owned all the child processes (the running containers).
- If a failure occurred, then there were orphaned processes.
- Building containers led to security vulnerabilities.
- All Docker operations had to be conducted by a user (or users) with the same full root authority.

There are probably more. Whether these issues have been fixed or you disagree with this characterization is not something this article  is going to debate. We in the community believe that Podman has  addressed many of these problems. If you want to take advantage of  Podman’s improvements, then this article is for you.

The Podman approach is simply to directly interact with the image  registry, with the container and image storage, and with the Linux  kernel through the runC container runtime process (not a daemon).

[![Podman architectural approach](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig2.png)](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig2.png)

Now that we’ve discussed some of the motivation it’s time to discuss  what that means for the user migrating to Podman.  There are a few  things to unpack here and we’ll get into each one separately:

- You install Podman instead of Docker. You do not need to start or manage a daemon process like the Docker daemon.
- The commands you are familiar with in Docker work the same for Podman.
- Podman stores its containers and images in a different place than Docker.
- Podman and Docker images are compatible.
- Podman does more than Docker for [Kubernetes](https://developers.redhat.com/topics/kubernetes/) environments.
- What is Buildah and why might I need it?



## Installing Podman

If you are using Docker today, you can remove it when you decide to  make the switch. However, you may wish to keep Docker around while you  try out Podman. There are some useful [tutorials](https://github.com/containers/libpod/blob/master/docs/tutorials/podman_tutorial.md) and an awesome [demonstration](https://github.com/containers/Demos/tree/master/building/buildah_intro) that you may wish to run through first so you can understand the  transition more. One example in the demonstration requires Docker in  order to show compatibility.

To install Podman on [Red Hat Enterprise Linux](https://developers.redhat.com/products/rhel/overview/) 7.6 or later, use the following; if you are using Fedora, then replace `yum` with `dnf`:

```
# yum -y install podman
```



## Podman commands are the same as Docker’s

When building Podman, the goal was to make sure that Docker users  could easily adapt. So all the commands you are familiar with also exist with Podman. In fact, the claim is made that if you have existing  scripts that run Docker you can create a `docker` alias for `podman` and all your scripts should work (`alias docker=podman`). Try it. Of course, you should stop Docker first (`systemctl stop docker`). There is a package you can install called `podman-docker` that does this for conversion for you. It drops a script at `/usr/bin/docker` that executes Podman with the same arguments.

The commands you are familiar with—`pull`, `push`, `build`, `run`, `commit`, `tag`, etc.—all exist with Podman.  See the [manual pages for Podman](https://github.com/containers/Demos/tree/master/building/buildah_intro) for more information. One notable difference is that Podman has added  some convenience flags to some commands. For example, Podman has added `--all` (`-a`) flags for `podman rm` and `podman rmi`. Many users will find that very helpful.

You can also run Podman from your normal non-root user in Podman 1.0  on Fedora. RHEL support is aimed for version 7.7 and 8.1 onwards.  Enhancements in userspace security have made this possible. Running  Podman as a normal user means that Podman will, by default, store images and containers in the user’s home directory. This is explained in the  next section. For more information on how Podman runs as a non-root  user, please check out Dan Walsh’s article: [How does rootless Podman work?](https://opensource.com/article/19/2/how-does-rootless-podman-work)

[![img](https://developers.redhat.com/blog/wp-content/uploads/2019/10/Java-Maven-Che-1024x314.png)](https://che.openshift.io/f?url=https://raw.githubusercontent.com/redhat-developer/devfile/master/getting-started/java-maven/devfile.yaml/?sc_cid=7013a000002D1quAAC)



## Podman and container images

When you first type `podman images`, you might be  surprised that you don’t see any of the Docker images you’ve already  pulled down. This is because Podman’s local repository is in `/var/lib/containers` instead of `/var/lib/docker`.  This isn’t an arbitrary change; this new storage structure is based on the Open Containers Initiative (OCI) standards.

In 2015, Docker, Red Hat, CoreOS, SUSE, Google, and other leaders in  the Linux containers industry created the Open Container Initiative in  order to provide an independent body to manage the standard  specifications for defining container images and the runtime. In order  to maintain that independence, the [containers/image](https://github.com/containers/image) and [containers/storage](https://github.com/containers/storage) projects were created on [GitHub](https://github.com/containers).

Since you can run `podman` without being root, there needs to be a separate place where `podman` can write images. Podman uses a repository in the user’s home directory: `~/.local/share/containers`. This avoids making `/var/lib/containers` world-writeable or other practices that might lead to potential  security problems. This also ensures that every user has separate sets  of containers and images and all can use Podman concurrently on the same host without stepping on each other. When users are finished with their work, they can push to a common registry to share their image with  others.

Docker users coming to Podman find that knowing these locations is useful for debugging and for the important `rm -rf /var/lib/containers`, when you just want to start over.  However, once you start using Podman, you’ll probably start using the new `-all` option to `podman rm` and `podman rmi` instead.



## Container images are compatible between Podman and other runtimes

Despite the new locations for the local repositories, the images  created by Docker or Podman are compatible with the OCI standard. Podman can push to and pull from popular container registries like [Quay.io](https://quay.io/) and Docker hub, as well as private registries. For example, you can  pull the latest Fedora image from the Docker hub and run it using  Podman. Not specifying a registry means Podman will default to searching through registries listed in the `registries.conf` file, in the order in which they are listed. An unmodified `registries.conf` file means it will look in the Docker hub first.

```
$ podman pull fedora:latest
$ podman run -it fedora bash
```

Images pushed to an image registry by Docker can be pulled down and  run by Podman. For example, an image (myfedora) I created using  Docker and pushed to my [Quay.io repository](https://quay.io/repository/ipbabble) (ipbabble) using Docker can be pulled and run with Podman as follows:

```
$ podman pull quay.io/ipbabble/myfedora:latest
$ podman run -it myfedora bash
```

Podman provides capabilities in its command-line `push` and `pull` commands to gracefully move images from `/var/lib/docker `to` /var/lib/containers` and vice versa.  For example:

```
$ podman push myfedora docker-daemon:myfedora:latest
```

Obviously, leaving out the `docker-daemon` above will default to pushing to the Docker hub.  Using `quay.io/myquayid/myfedora` will push the image to the Quay.io registry (where `myquayid` below is your personal Quay.io account):

```
$ podman push myfedora quay.io/myquayid/myfedora:latest
```

If you are ready to remove Docker, you should shut down the daemon  and then remove the Docker package using your package manager. But  first, if you have images you created with Docker that you wish to keep, you should make sure those images are pushed to a registry so that you  can pull them down later. Or you can use Podman to pull each image (for  example, fedora) from the host’s Docker repository into Podman’s  OCI-based repository. With RHEL you can run the following:

```
# systemctl stop docker
# podman pull docker-daemon:fedora:latest
# yum -y remove docker  # optional
```



## Podman helps users move to Kubernetes

Podman provides some extra features that help developers and  operators in Kubernetes environments. There are extra commands provided  by Podman that are not available in Docker. If you are familiar with  Docker and are considering using Kubernetes/[OpenShift](http://openshift.com/) as your container platform, then Podman can help you.

Podman can generate a Kubernetes YAML file based on a running container using `podman generate kube`. The command `podman pod` can be used to help debug running Kubernetes pods along with the  standard container commands.  For more details on how Podman can help  you transition to Kubernetes, see the following article by Brent Baude: [Podman can now ease the transition to Kubernetes and CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/).



## What is Buildah and why would I use it?

Buildah actually came first. And maybe that’s why some Docker users  get a bit confused. Why do these Podman evangelists also talk about  Buildah? Does Podman not do builds?

Podman does do builds and for those familiar with Docker, the build  process is the same. You can either build using a Dockerfile using `podman build` or you can run a container and make lots of changes and then commit  those changes to a new image tag. Buildah can be described as a superset of commands related to creating and managing container images and,  therefore, it has much finer-grained control over images. Podman’s `build` command contains a subset of the Buildah functionality. It uses the same code as Buildah for building.

The most powerful way to use Buildah is to write Bash scripts for  creating your images—in a similar way that you would write a Dockerfile.

I like to think of the evolution in the following way. When Kubernetes moved to [CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/) based on the OCI runtime specification, there was no need to run a  Docker daemon and, therefore, no need to install Docker on any host in  the Kubernetes cluster for running pods and containers. Kubernetes could call CRI-O and it could call runC directly. This, in turn, starts the  container processes. However, if we want to use the same Kubernetes  cluster to do builds, as in the case of OpenShift clusters, then we  needed a new tool to perform builds that would not require the Docker  daemon and subsequently require that Docker be installed. Such a tool,  based on the `containers/storage` and `containers/image` projects, would also eliminate the security risk of the open Docker daemon socket during builds, which concerned many users.

Buildah (named for fun because of Dan Walsh’s Boston accent when  pronouncing "builder") fit this bill. For more information on Buildah,  see [buildah.io](https://buildah.io/) and specifically see the blogs and tutorials sections.

There are a couple of extra things practitioners need to understand about Buildah:

1. It allows for finer control of creating image layers. This is a  feature that many container users have been asking for for a long time.  Committing many changes to a single layer is desirable.
2. Buildah’s `run` command is not the same as Podman’s `run` command.  Because Buildah is for building images, the `run` command is *essentially the same as the Dockerfile* `RUN` *command*. In fact, I remember the week this was made explicit. I was foolishly  complaining that some port or mount that I was trying wasn’t working as I expected it to.  Dan ([@rhatdan](https://twitter.com/rhatdan)) weighed in and said that Buildah should not be supporting running  containers in that way. No port mapping. No volume mounting. Those flags were removed.  Instead `buildah run` is for running specific commands in order to help build a container image, for example, `buildah run dnf -y install nginx`.
3. Buildah can build images from scratch, that is, images with nothing in them at all. Nothing. In fact, looking at the container storage  created as a result of a `buildah from scratch` command  yields an empty directory. This is useful for creating very lightweight  images that contain only the packages needed in order to run your  application.

A good example use case for a scratch build is to consider the  development images versus staging or production images of a Java  application. During development, a Java application container image may  require the Java compiler and Maven and other tools. But in production,  you may only require the Java runtime and your packages. And, by the  way, you also do not require a package manager such as DNF/YUM or even  Bash. Buildah is a powerful CLI for this use case. See the diagram  below. For more information, see [Building a Buildah Container Image for Kubernetes](https://buildah.io/blogs/2018/03/01/building-buildah-container-image-for-kubernetes.html) and also this [Buildah introduction demo](https://github.com/containers/Demos/tree/master/building/buildah_intro).

[![Buildah is a powerful CLI](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig3-1024x703.png)](https://developers.redhat.com/blog/wp-content/uploads/2019/02/fig3.png)

Getting back to the evolution story...Now that we had solved the  Kubernetes runtime issue with CRI-O and runC, and we had solved the  build problem with Buildah, there was still one reason why Docker was  still needed on a Kubernetes host: debugging. How can we debug container issues on a host if we don't have the tools to do it? We would need to  install Docker, and then we are back where we started with the Docker  daemon on the host. Podman solves this problem.

Podman becomes a tool that solves two problems. It allows operators  to examine containers and images with commands they are familiar with  using. And it also provides developers with the same tools. So Docker  users, developers, or operators, can move to Podman, do all the fun  tasks that they are familiar with from using Docker, and do much more.



## Conclusion

I hope this article has been useful and will help you migrate to using Podman (and Buildah) confidently and successfully.

For more information:

- [Podman.io](https://podman.io/) and [Buildah.io](https://buildah.io/) project web sites

- github.com/containers

   projects (get involved, get the source, see what's being developed): 

  - [libpod](https://github.com/containers/libpod) (Podman)
  - [buildah](https://github.com/containers/buildah)
  - [image](https://github.com/containers/image) (code for working with OCI container images)
  - [storage](https://github.com/containers/storage) (code for local image and container storage)



## Related Articles

- [Containers without daemons: Podman and Buildah available in RHEL 7.6 and RHEL 8 Beta](https://developers.redhat.com/blog/2018/11/20/buildah-podman-containers-without-daemons/)
- [Podman: Managing pods and containers in a local container runtime](https://developers.redhat.com/blog/2019/01/15/podman-managing-containers-pods/)
- [Managing containerized system services with Podman](https://developers.redhat.com/blog/2018/11/29/managing-containerized-system-services-with-podman/) (Use systemd to manage your podman containers)
- [Building a Buildah Container Image for Kubernetes](https://buildah.io/blogs/2018/03/01/building-buildah-container-image-for-kubernetes.html)
- [Podman can now ease the transition to Kubernetes and CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/)
- [Security Considerations for Container Runtimes](https://developers.redhat.com/blog/2018/12/19/security-considerations-for-container-runtimes/) (Video of Dan Walsh's talk from KubeCon 2018)
- [IoT edge development and deployment with containers through OpenShift: Part 1](https://developers.redhat.com/blog/2019/01/31/iot-edge-development-and-deployment-with-containers-through-openshift-part-1/) (Building and testing ARM64 containers on OpenShift using podman, qemu, binfmt_misc, and Ansible)

 

*Last updated:              June 17, 2021*