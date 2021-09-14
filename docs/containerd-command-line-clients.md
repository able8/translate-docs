# Why and How to Use containerd from the Command Line

September 12, 2021

[Containers](http://iximiuz.com/en/categories/?category=Containers)

[containerd](https://github.com/containerd/containerd) is a [high-level container runtime](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-runtimes), _aka_ [container manager](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-management). To put it simply, it's a daemon that manages the complete container lifecycle on a single host: creates, starts, stops containers, pulls and stores images, configures mounts, networking, etc.

containerd is designed to be easily embeddable into larger systems. [Docker uses containerd under the hood](https://www.docker.com/blog/what-is-containerd-runtime/) to run containers. [Kubernetes can use containerd via CRI](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#containerd) to manage containers on a single node. But smaller projects also can benefit from the ease of integrating with containerd - for instance, [faasd](https://github.com/openfaas/faasd) uses containerd (we need more **d**'s!) to spin up a full-fledged [Function-as-a-Service](https://en.wikipedia.org/wiki/Function_as_a_service) solution on a standalone server.

![Docker and Kubernetes use containerd](http://iximiuz.com/containerd-command-line-clients/docker-and-kubernetes-use-containerd-2000-opt.png)

However, using containerd programmatically is not the only option. It also can be used from the command line via one of the available clients. The resulting container UX may not be as comprehensive and user-friendly as the one provided by the `docker` client, but it still can be useful, for instance, for debugging or learning purposes.

![containerd command-line clients (ctr, nerdctl, crictl)](http://iximiuz.com/containerd-command-line-clients/containerd-command-line-clients-2000-opt.png)

## How to use containerd with ctr

[ctr](https://github.com/containerd/containerd/tree/e1ad7791077916aac9c1f4981ad350f0e3fce719/cmd/ctr) is a command-line client shipped as part of the containerd project. If you have containerd running on a machine, chances are the `ctr` binary is also there.

The `ctr` interface is [obviously] incompatible with Docker CLI and, at first sight, may look not so user-friendly. Apparently, its primary audience is containerd developers testing the daemon. However, since it's the closest thing to the actual containerd API, it can serve as a great exploration means - by examining the available commands, you can get a rough idea of what containerd can and cannot do.

`ctr` is also well-suitable for learning the capabilities of [low-level [OCI] container runtimes](http://iximiuz.com/en/posts/oci-containers/) since `ctr + containerd` is [much closer to actual containers](http://iximiuz.com/en/posts/implementing-container-runtime-shim/) than `docker + dockerd`.

### Working with container images using ctr

When **pulling images**, the fully-qualified reference seems to be required, so you cannot omit the registry or the tag part:

```bash
$ ctr images pull docker.io/library/nginx:1.21
$ ctr images pull docker.io/kennethreitz/httpbin:latest
$ ctr images pull docker.io/kennethreitz/httpbin:latest
$ ctr images pull quay.io/quay/redis:latest

```

To **list local images**, one can use:

```bash
$ ctr images ls

```

Surprisingly, containerd doesn't provide out-of-the-box image building support. However, containerd itself is often used to build images by higher-level tools.

Check out my investigation post on [what actually happens when you build an image](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/) to learn more about image building internals.

Instead of building images with `ctr`, you can **import existing images** built with `docker build` or other OCI-compatible software:

```bash
$ docker build -t my-app .
$ docker save -o my-app.tar my-app

$ ctr images import my-app.tar

```

With `ctr`, you can also **mount images** for future exploration:

```bash
$ mkdir /tmp/httpbin
$ ctr images mount docker.io/kennethreitz/httpbin:latest /tmp/httpbin

$ ls -l /tmp/httpbin/
total 80
drwxr-xr-x 2 root root 4096 Oct 18  2018 bin
drwxr-xr-x 2 root root 4096 Apr 24  2018 boot
drwxr-xr-x 4 root root 4096 Oct 18  2018 dev
drwxr-xr-x 1 root root 4096 Oct 24  2018 etc
drwxr-xr-x 2 root root 4096 Apr 24  2018 home
drwxr-xr-x 3 root root 4096 Oct 24  2018 httpbin
...

$ ctr images unmount /tmp/httpbin

```

To **remove images** using `ctr`, run:

```bash
$ ctr images remove docker.io/library/nginx:1.21

```

### Working with containers using ctr

Having a local image, you can **run a container** with `ctr run <image-ref> <container-id>`. For instance:

```bash
$ ctr run --rm -t docker.io/library/debian:latest cont1

```

Notice that unlike user-friendly `docker run` generating a unique container ID for you, with `ctr`, you must supply the unique container ID yourself. The `ctr run` command also supports only some of the familiar `docker run` flags: `--env`, `-t,--tty`, `-d,--detach`, `--rm`, etc. But no port publishing or automatic container restart with `--restart=always` out of the box.

Similarly to images, you can **list existing containers** with:

```bash
$ ctr containers ls

```

Interesting that the `ctr run` command is actually a shortcut for `ctr container create` \+ `ctr task start`:

```bash
$ ctr container create -t docker.io/library/nginx:latest nginx_1
$ ctr container ls
CONTAINER    IMAGE                              RUNTIME
nginx_1      docker.io/library/nginx:latest     io.containerd.runc.v2

$ ctr task ls
TASK    PID    STATUS        # Empty!

$ ctr task start -d nginx_1  # -d for --detach
$ ctr task list
TASK     PID      STATUS
nginx_1  10074    RUNNING

```

I like this separation of `container` and `task` subcommands because it reflects the often forgotten nature of OCI containers. Despite the common belief, containers aren't processes - [_containers are isolated and restricted execution environments_](http://iximiuz.com/en/posts/oci-containers/) for processes.

With `ctr task attach`, you can **reconnect to the stdio streams** of an existing task running inside of a container:

```bash
$ ctr task attach nginx_1
2021/09/12 15:42:20 [notice] 1#1: using the "epoll" event method
2021/09/12 15:42:20 [notice] 1#1: nginx/1.21.3
2021/09/12 15:42:20 [notice] 1#1: built by gcc 8.3.0 (Debian 8.3.0-6)
2021/09/12 15:42:20 [notice] 1#1: OS: Linux 4.19.0-17-amd64
2021/09/12 15:42:20 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1024:1024
2021/09/12 15:42:20 [notice] 1#1: start worker processes
2021/09/12 15:42:20 [notice] 1#1: start worker process 31
...

```

Much like with `docker`, you can **execute a task in an existing container**:

```bash
$ ctr task exec -t --exec-id bash_1 nginx_1 bash

# From inside the container:
$ root@host:/# curl 127.0.0.1:80
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
...

```

Before removing a container, all its tasks must be **stopped**:

```bash
$ ctr task kill -9 nginx_1

```

Alternatively, you can **remove running tasks** using the `--force` flag:

```bash
$ ctr task rm -f nginx_1

```

Finally, to **remove the container**, run:

```bash
$ ctr container rm nginx_1

```

## How to use containerd with nerdctl

[nerdctl](https://github.com/containerd/nerdctl) is a relatively new command-line client for containerd. Unlike `ctr`, `nerdctl` aims to be user-friendly and Docker-compatible. To some extent, `nerdctl + containerd` can seamlessly replace `docker + dockerd`. However, [this does not seem to be the goal of the project](https://medium.com/nttlabs/nerdctl-359311b32d0e):

> The goal of `nerdctl` is to facilitate experimenting the cutting-edge features of containerd that are not present in Docker. Such features include, but not limited to, lazy-pulling (stargz) and encryption of images (ocicrypt). These features are expected to be eventually available in Docker as well, however, it is likely to take several months, or perhaps years, as Docker is currently designed to use only a small portion of the containerd subsystems. Refactoring Docker to use the entire containerd would be possible, but not straightforward. So we [ [NTT](https://www.rd.ntt/e/)] decided to create a new CLI that fully uses containerd, but we do not intend to complete with Docker. We have been contributing to Docker/Moby as well as containerd, and will continue to do so.

From the basic usage standpoint, comparing to `ctr`, `nerdctl` supports:

- Image building with`nerdctl build`
- Container networking management
- Docker Compose with`nerdctl compose up`

And the coolest part about it is that `nerdctl` tries to provide the identical to `docker` (and `podman`) command-line UX. So, **if you are familiar with `docker` (or `podman`) CLI, you are already familiar with `nerdctl`.**

## How to use containerd with crictl

[crictl](https://github.com/kubernetes-sigs/cri-tools) is a command-line client for [[Kubernetes] CRI-compatible container runtimes](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/).

Click here to learn more about Kubernetes Container Runtime Interface (CRI).

**[Kubernetes Container Runtime Interface (CRI)](https://github.com/kubernetes/cri-api/)** was introduced to make Kubernetes container runtime-agnostic. The Kubernetes node agent, _kubelet_, implementing the CRI client API, can use any container runtime implementing the CRI server API to manage containers and pods on its node.

![Kubernetes CRI](http://iximiuz.com/containerd-command-line-clients/cri-2000-opt.png)

_Kubernetes CRI._

Since version 1.1, containerd comes with a built-in CRI plugin. Hence, containerd is a CRI-compatible container runtime. Therefore, it can be used with `crictl`.

`crictl` was created to inspect and debug container runtimes and applications on a Kubernetes node. [It supports the following operations](https://github.com/kubernetes-sigs/cri-tools/blob/98f3364b4b684966b27bf372412d805d7dcbcb10/docs/crictl.md):

```bash
attach: Attach to a running container
create: Create a new container
exec: Run a command in a running container
version: Display runtime version information
images, image, img: List images
inspect: Display the status of one or more containers
inspecti: Return the status of one or more images
imagefsinfo: Return image filesystem info
inspectp: Display the status of one or more pods
logs: Fetch the logs of a container
port-forward: Forward local port to a pod
ps: List containers
pull: Pull an image from a registry
run: Run a new container inside a sandbox
runp: Run a new pod
rm: Remove one or more containers
rmi: Remove one or more images
rmp: Remove one or more pods
pods: List pods
start: Start one or more created containers
info: Display information of the container runtime
stop: Stop one or more running containers
stopp: Stop one or more running pods
update: Update one or more running containers
config: Get and set crictl client configuration options
stats: List container(s) resource usage statistics

```

The interesting part here is that with `crictl + containerd` bundle, one can learn how pods are actually implemented. [But this topic deserves its own blog post](http://iximiuz.com/en/newsletter/) 😉

For more information on how to use `crictl` with containerd, check out [this document (part of the containerd project)](https://github.com/containerd/cri/blob/68b61297b59e38c1088db10fbd19807a4ffbad87/docs/crictl.md).

### More [Containers](http://iximiuz.com/en/categories/?category=Containers) Posts From This Blog

- [Journey From Containerization to Orchestration and Beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)
- [Containers aren't Linux processes](http://iximiuz.com/en/posts/oci-containers/)
- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [conman - [the] container manager: inception](http://iximiuz.com/en/posts/conman-the-container-manager-inception/)
- [Implementing Container Runtime Shim: runc](http://iximiuz.com/en/posts/implementing-container-runtime-shim/)
- [Implementing Container Runtime Shim: First Code](http://iximiuz.com/en/posts/implementing-container-runtime-shim-2/)
- [Implementing Container Runtime Shim: Interactive Containers](http://iximiuz.com/en/posts/implementing-container-runtime-shim-3/)

[containerd,](javascript: void 0) [ctr,](javascript: void 0) [crictl,](javascript: void 0) [nerdctl](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

