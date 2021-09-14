# Does Container Image Have an OS Inside

May 7, 2020 (Updated: August 7, 2021)

[Containers,](http://iximiuz.com/en/categories/?category=Containers) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

_Not every container has an operating system inside, but every one of them needs your Linux kernel._

_**Disclaimer 1:** before going any further it's important to understand the difference between a kernel, an operating system, and a distribution._

_1. [Linux kernel](https://en.wikipedia.org/wiki/Linux_kernel) is the core part of the Linux operating system. It's what originally Linus wrote.

2. [Linux operating system](https://en.wikipedia.org/wiki/Linux) is a combination of the kernel and a user-land (libraries, GNU utilities, config files, etc)._ - _[Linux distribution](https://en.wikipedia.org/wiki/Linux_distribution) is a particular version of the Linux operating system like Debian, CentOS, or Alpine._

_**Disclaimer 2:** the title of this article should have sounded like ["Does container image have a whole Linux distribution inside"](https://www.reddit.com/r/programming/comments/gfbaor/not_every_container_has_an_operating_system_inside/fpusvxu?utm_source=share&utm_medium=web2x), but I personally find this wording a bit boring_ 🤪

### Does a Container have an Operating System inside?

The majority of Docker examples out there explicitly or implicitly rely on some flavor of the Linux operating system running inside a container. I tried to quickly compile a list of the most prominent samples:

Running an interactive shell in the `debian jessie` distribution:

```bash
$ docker run -it debian:jessie

```

Running an `nginx` web-sever in a container and examine its config using `cat` utility:

```bash
$ docker run -d -P --name nginx nginx:latest
$ docker exec -it nginx cat /etc/nginx/nginx.conf

```

Building an image based on Alpine Linux:

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

For the newcomers learning the containerization through hands-on experience, this may lead to a _false_ impression that containers are somewhat indistinguishable from full-fledged operating systems and that they are always based on well-known and wide-spread Linux distributions like `debian`, `centos`, or `alpine`.

At the same time, approaching the containerization topic from the theoretical side ( [1](https://techcrunch.com/2016/10/16/wtf-is-a-container/), [2](https://www.docker.com/resources/what-container), [3](https://cloud.google.com/containers)) may lead to a rather opposite impression that containers (unlike the traditional virtual machines) are supposed to pack only the application (i.e. your code) and its dependencies (i.e. some libraries) instead of a trying to ship a full operating system.

As it usually happens, the truth lies somewhere in between both statements. From the implementation standpoint, **a container is indeed just a process (or a bunch of processes) running on the Linux host**. The container process is isolated ( [namespaces](https://docs.docker.com/engine/security/security/#kernel-namespaces)) from the rest of the system and restricted from both the resource consumption ( [cgroups](https://docs.docker.com/engine/security/security/#control-groups)) and security ( [capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html), [AppArmor](https://docs.docker.com/engine/security/apparmor/), [Seccomp](https://docs.docker.com/engine/security/seccomp/)) standpoints. But in the end, this is still a regular process, same as any other process on the host system.

> OCI/Docker containers thread:
>
> Containers are simply isolated and restricted Linux processes. [#Docker](https://twitter.com/hashtag/Docker?src=hash&ref_src=twsrc%5Etfw) [#containers](https://twitter.com/hashtag/containers?src=hash&ref_src=twsrc%5Etfw) [#linux](https://twitter.com/hashtag/linux?src=hash&ref_src=twsrc%5Etfw)
>
> — Ivan Velichko (@iximiuz) [May 10, 2020](https://twitter.com/iximiuz/status/1259569908385529859?ref_src=twsrc%5Etfw)

Just run `docker run -d nginx` and conduct your own investigation:

![ps axf output (excerpt)](http://iximiuz.com/not-every-container-has-an-operating-system-inside/nginx-ps2.png)

_`ps axf` output (excerpt)_

![systemctl status output (excerpt)](http://iximiuz.com/not-every-container-has-an-operating-system-inside/nginx-cgroups2.png)

_`systemctl status` output (excerpt)_

![sudo lsns output](http://iximiuz.com/not-every-container-has-an-operating-system-inside/nginx-lsns2.png)

_`sudo lsns` output_

Well, if a container is just a regular Linux process, we could try to run a single executable file inside of a container. I.e. instead of putting our application into a fully-featured Linux distribution, we will try to build a container image consisting of a folder with a single file inside. Upon the launch, this folder will become a root folder for the containerized environment.

### Create a Container from scratch (with a single executable binary inside)

If you have _Go_ installed on your system, you can utilize its handy cross-compilation abilities:

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello from OS-less container (Go edition)")
}

```

Build the program from above using:

```bash
$ GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o hello
$ file hello
> hello: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked

```

_Click here to see how to compile a similar C program._

```c
// main.c
#include <stdio.h>

int main() {
    printf("Hello from OS-less container (C edition)\n");
}

```

Compile it using the following builder container:

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

![Inspecting hello image with dive tool (screenshot)](http://iximiuz.com/not-every-container-has-an-operating-system-inside/dive-hello2.png)

_`dive hello`_

This exercise is roughly what the Docker's [hello-world](https://github.com/docker-library/hello-world) example does. There are two key moments here. First, we based our image on a so-called [`scratch`](https://hub.docker.com/_/scratch) image. This is just an empty image, i.e. the building starts from the empty folder and then just copies the executable file `hello` into it. Second, we used a statically linked binary file. I.e. there is no dependency on some shared libraries from the system. So, a bare Linux kernel is enough to execute it.

Now, what if we inspect the [`nginx`](https://hub.docker.com/_/nginx) image which we used at the beginning of this article?

![Inspecting nginx image with dive tool (screenshot)](http://iximiuz.com/not-every-container-has-an-operating-system-inside/dive-nginx2.png)

_`dive nginx`_

Well, the directory tree looks like a root filesystem of some Linux distribution. If we take a look at the [corresponding Dockerfile](https://github.com/nginxinc/docker-nginx/blob/594ce7a8bc26c85af88495ac94d5cd0096b306f7/mainline/buster/Dockerfile) we can notice that `nginx` image is based on [`debian`](https://hub.docker.com/_/debian):

```dockerfile
FROM debian:buster-slim

LABEL maintainer="NGINX Docker Maintainers <docker-maint@nginx.com>"

ENV NGINX_VERSION   1.17.10
ENV NJS_VERSION     0.3.9
ENV PKG_RELEASE     1~buster

...

```

And if we dive deeper and examine `debian:buster-slim` [Dockerfile](https://github.com/debuerreotype/docker-debian-artifacts/blob/d6ff3e75eeae3ea012c30bce9054336d99d1a20a/buster/slim/Dockerfile) we will see that it just copies [a root filesystem](https://github.com/debuerreotype/docker-debian-artifacts/blob/d6ff3e75eeae3ea012c30bce9054336d99d1a20a/buster/slim/rootfs.tar.xz) to an empty folder:

```dockerfile
FROM scratch
ADD rootfs.tar.xz /
CMD ["bash"]

```

[Combining Debian's user-land with the host's kernel](http://iximiuz.com/en/posts/from-docker-container-to-bootable-linux-disk-image/) containers start resembling fully-featured operating systems. With `nginx` image we can use the shell to interact with the container:

![Interactive shell with running nginx container.](http://iximiuz.com/not-every-container-has-an-operating-system-inside/nginx-exec-bash2.png)

_Interactive shell with running nginx container._

Can we do the same for our slim `hello` container? Obviously not, there is no `bash` executable inside:

![Demonstraiting that hello container doesn't have bash inside.](http://iximiuz.com/not-every-container-has-an-operating-system-inside/hello-run-bash2.png)

_`hello` container doesn't have `bash` inside._

### Wrapping up

So, what should be the conclusion here? The [virtualization capabilities](https://en.wikipedia.org/wiki/OS-level_virtualization) of containers turned out to be so powerful that people started packing fully-featured user-lands like `debian` (or more lightweight alternatives like `alpine` or `busybox`) into containers. By virtue of this ability:

- We can play with various Linux distribution using a simple`docker run -it fedora bash`.
- We can use OS commands including package managers like`yum` or `apt` while building our images.
- We can interact with running containers using various OS utilities.

But with great power comes great responsibility. Huge containers carrying lots of unnecessary tools slow down deployments and increase the surface of potential cyberattacks.

Make code, not war!

### Related articles

- [Containers aren't Linux processes](http://iximiuz.com/en/posts/oci-containers/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)
- [A journey from containerization to orchestration and beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)
- [From Docker Container to Bootable Linux Disk Image](http://iximiuz.com/en/posts/from-docker-container-to-bootable-linux-disk-image/)

[docker,](javascript: void 0) [linux,](javascript: void 0) [container](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

