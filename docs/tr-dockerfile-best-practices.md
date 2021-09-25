# Top 20 Dockerfile best practices

# 前 20 个 Dockerfile 最佳实践

                     on March 9, 2021

2021 年 3 月 9 日

Learn how to **prevent security issues** and **optimize** containerized applications by applying a quick set of **Dockerfile best practices** in your image builds.

通过在您的映像构建中应用一组快速的 **Dockerfile 最佳实践**，了解如何**防止安全问题** 和**优化**容器化应用程序。

If you are familiar with containerized applications and microservices, you might have realized that your services might be *micro*; but detecting vulnerabilities, investigating security issues, and  reporting and fixing them after the deployment is making your management overhead *macro*.

如果您熟悉容器化应用程序和微服务，您可能已经意识到您的服务可能是*微*；但是检测漏洞、调查安全问题并在部署后报告和修复它们会增加您的管理开销*宏*。

Much of this overhead can be prevented by **shifting left security**, tackling potential problems as soon as possible in your development workflow. *We recently covered in this blog how [image scanning best practices](https://sysdig.com/blog/image-scanning-best-practices/) helps you shift left security.*

通过**左移安全**，尽快解决开发工作流程中的潜在问题，可以防止大部分开销。 *我们最近在此博客中介绍了 [图像扫描最佳实践](https://sysdig.com/blog/image-scanning-best-practices/) 如何帮助您转移安全性。*

A **well crafted Dockerfile** will avoid the need for  privileged containers, exposing unnecessary ports, unused packages,  leaked credentials, etc., or anything that can be used for an attack. Getting rid of the known risks in advance will help reduce your security management and operational overhead.

**精心设计的 Dockerfile** 将避免需要特权容器、暴露不必要的端口、未使用的包、泄露的凭据等，或任何可用于攻击的内容。提前消除已知风险将有助于减少您的安全管理和运营开销。

Following the best practices, patterns, and recommendations for the  tools you use will help you avoid common errors and pitfalls.

遵循您使用的工具的最佳实践、模式和建议将帮助您避免常见错误和陷阱。

![img](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-01.png)

This article dives into a curated **list of Docker security best practices** that are focused on writing Dockerfiles and container security, but also cover other related topics, like image optimization:

本文深入探讨了一个精选的 **Docker 安全最佳实践列表**，重点关注编写 Dockerfile 和容器安全，但也涵盖其他相关主题，例如图像优化：

1. Avoid unnecessary privileges

    1.避免不必要的特权

   1. [Avoid running containers as root](https://sysdig.com/blog/dockerfile-best-practices/#1-1).
    2. [Don’t bind to a specific UID](https://sysdig.com/blog/dockerfile-best-practices/#1-2).
    3. [Make executables owned by root and not writable](https://sysdig.com/blog/dockerfile-best-practices/#1-3).

1. [避免以root身份运行容器](https://sysdig.com/blog/dockerfile-best-practices/#1-1)。
   2. [不要绑定到特定的 UID](https://sysdig.com/blog/dockerfile-best-practices/#1-2)。
   3. [使可执行文件归root所有且不可写](https://sysdig.com/blog/dockerfile-best-practices/#1-3)。

2. Reduce attack surface

    2. 减少攻击面

   1. [Leverage multistage builds](https://sysdig.com/blog/dockerfile-best-practices/#2-1).
    2. [Use distroless images, or build your own from scratch](https://sysdig.com/blog/dockerfile-best-practices/#2-2).
    3. [Update your images frequently](https://sysdig.com/blog/dockerfile-best-practices/#2-3).
    4. [Watch out for exposed ports](https://sysdig.com/blog/dockerfile-best-practices/#2-4).

1. [利用多阶段构建](https://sysdig.com/blog/dockerfile-best-practices/#2-1)。
   2. [使用 distroless 镜像，或从头开始构建自己的镜像](https://sysdig.com/blog/dockerfile-best-practices/#2-2)。
   3. [经常更新你的图片](https://sysdig.com/blog/dockerfile-best-practices/#2-3)。
   4. [注意暴露端口](https://sysdig.com/blog/dockerfile-best-practices/#2-4)。

3. Prevent confidential data leaks

    3. 防止机密数据泄露

   1. [Never put secrets or credentials in Dockerfile instructions](https://sysdig.com/blog/dockerfile-best-practices/#3-1).
    2. [Prefer COPY over ADD](https://sysdig.com/blog/dockerfile-best-practices/#3-2).
    3. [Be aware of the Docker context, and use .dockerignore](https://sysdig.com/blog/dockerfile-best-practices/#3-3).

1. [永远不要在 Dockerfile 指令中放置机密或凭据](https://sysdig.com/blog/dockerfile-best-practices/#3-1)。
   2. [优先复制而不是添加](https://sysdig.com/blog/dockerfile-best-practices/#3-2)。
   3. [注意Docker上下文，使用.dockerignore](https://sysdig.com/blog/dockerfile-best-practices/#3-3)。

4. Others

    4. 其他

   1. [Reduce the number of layers, and order them intelligently](https://sysdig.com/blog/dockerfile-best-practices/#4-1).
    2. [Add metadata and labels](https://sysdig.com/blog/dockerfile-best-practices/#4-2).
    3. [Leverage linters to automatize checks](https://sysdig.com/blog/dockerfile-best-practices/#4-3).
    4. [Scan your images locally during development](https://sysdig.com/blog/dockerfile-best-practices/#4-4).

1. [减少层数，智能排序](https://sysdig.com/blog/dockerfile-best-practices/#4-1)。
   2. [添加元数据和标签](https://sysdig.com/blog/dockerfile-best-practices/#4-2)。
   3. [利用 linters 自动化检查](https://sysdig.com/blog/dockerfile-best-practices/#4-3)。
   4. [开发时在本地扫描你的图片](https://sysdig.com/blog/dockerfile-best-practices/#4-4)。

5. Beyond image building

    5. 超越形象建设

   1. [Protect the docker socket and TCP connections](https://sysdig.com/blog/dockerfile-best-practices/#5-1).
    2. [Sign your images, and verify them on runtime](https://sysdig.com/blog/dockerfile-best-practices/#5-2).
    3. [Avoid tag mutability](https://sysdig.com/blog/dockerfile-best-practices/#5-3).
    4. [Don’t run your environment as root](https://sysdig.com/blog/dockerfile-best-practices/#5-4).
    5. [Include a health check](https://sysdig.com/blog/dockerfile-best-practices/#5-5).
    6. [Restrict your application capabilities](https://sysdig.com/blog/dockerfile-best-practices/#5-6).

1. [保护docker socket和TCP连接](https://sysdig.com/blog/dockerfile-best-practices/#5-1)。
   2. [签署您的图像，并在运行时验证它们](https://sysdig.com/blog/dockerfile-best-practices/#5-2)。
   3. [避免标签可变性](https://sysdig.com/blog/dockerfile-best-practices/#5-3)。
   4. [不要以root身份运行你的环境](https://sysdig.com/blog/dockerfile-best-practices/#5-4)。
   5. [包括健康检查](https://sysdig.com/blog/dockerfile-best-practices/#5-5)。
   6. [限制你的应用能力](https://sysdig.com/blog/dockerfile-best-practices/#5-6)。

We have grouped our selected set of Dockerfile best practices by topic. Please remember that Dockerfile best practices are **just a piece in the whole development process**. We include a closing section pointing to related container image  security and shifting left security resources to apply before and after  the image building.

我们按主题对选定的 Dockerfile 最佳实践集进行了分组。请记住，Dockerfile 最佳实践**只是整个开发过程中的一个片段**。我们包括一个结束部分，指向相关的容器镜像安全和转移左侧安全资源以在镜像构建之前和之后应用。

![img](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-02-local-development.png)

## #1 Avoid unnecessary privileges 

## #1 避免不必要的特权

These tips follow the principle of least privilege so your service or  application only has access to the resources and information necessary  to perform its purpose.

这些提示遵循最小权限原则，因此您的服务或应用程序只能访问执行其目的所需的资源和信息。

### #1.1 Rootless containers

### #1.1 无根容器

Our [recent report highlighted that 58% of images](https://sysdig.com/blog/sysdig-2021-container-security-usage-report/) are running the container entrypoint as **root (UID 0)** . However, it is a Dockerfile best practice to avoid doing that. There  are very few use cases where the container needs to execute as **root**, so don’t forget to include the *USER* instruction to change the default effective UID to a non-root user.

我们的 [最近的报告强调 58% 的图像](https://sysdig.com/blog/sysdig-2021-container-security-usage-report/) 以 **root (UID 0)** 身份运行容器入口点.但是，避免这样做是 Dockerfile 的最佳实践。容器需要以 **root** 身份执行的用例很少，所以不要忘记包含 *USER* 指令以将默认有效 UID 更改为非 root 用户。

Furthermore, your execution environment might block containers running  as root by default (i.e., Openshift requires additional  SecurityContextConstraints).

此外，您的执行环境可能会默认阻止以 root 身份运行的容器（即 Openshift 需要额外的 SecurityContextConstraints）。

Running as non-root might require a couple of additional steps in your Dockerfile, as now you will need to:

以非 root 用户身份运行可能需要在 Dockerfile 中执行几个额外的步骤，因为现在您需要：

- Make sure the user specified in the *USER* instruction exists inside the container.
- Provide appropriate file system permissions in the locations where the process will be reading or writing.

- 确保 *USER* 指令中指定的用户存在于容器内。
- 在进程将要读取或写入的位置提供适当的文件系统权限。

```
FROM alpine:3.12
# Create user and set ownership and permissions as required
RUN adduser -D myuser && chown -R myuser /myapp-data
# ... copy application files
USER myuser
ENTRYPOINT ["/myapp"]
```

You might see containers that start as root and then use [gosu](https://github.com/tianon/gosu) or [su-exec](https://github.com/ncopa/su-exec) to drop to a standard user.

您可能会看到以 root 身份启动然后使用 [gosu](https://github.com/tianon/gosu) 或 [su-exec](https://github.com/ncopa/su-exec) 删除的容器给标准用户。

Also, if a container needs to run a very specific command as root, it may rely on [sudo](https://www.sudo.ws/).

此外，如果容器需要以 root 身份运行非常具体的命令，它可能依赖于 [sudo](https://www.sudo.ws/)。

While these two alternatives are better than running as root, it might not work in restricted environments like Openshift.

虽然这两个替代方案比以 root 身份运行要好，但它可能不适用于 Openshift 等受限环境。

### #1.2 Don’t bind to a specific UID

### #1.2 不要绑定到特定的 UID

Run the container as a non-root user, but don’t make that user UID a requirement.

以非 root 用户身份运行容器，但不要要求该用户 UID。

Why?

为什么？

- Openshift, by default, will use random UIDs when running containers.
- Forcing a specific UID (i.e., the first standard user with `UID 1000`) requires adjusting the permissions of any bind mount, like a host  folder for data persistence. Alternatively, if you run the container (`-u` option in docker) with the host UID, it might break the service when trying to read or write from folders within the container.

- 默认情况下，Openshift 在运行容器时将使用随机 UID。
- 强制使用特定的 UID（即具有“UID 1000”的第一个标准用户）需要调整任何绑定安装的权限，例如用于数据持久性的主机文件夹。或者，如果您使用主机 UID 运行容器（docker 中的“-u”选项），则在尝试从容器内的文件夹读取或写入时可能会中断服务。

```
...
RUN mkdir /myapp-tmp-dir && chown -R myuser /myapp-tmp-dir
USER myuser
ENTRYPOINT ["/myapp"]
```

This container will have trouble if running with an UID different than `myuser`, as the application won’t be able to write in `/myapp-tmp-dir` folder.

如果使用不同于 `myuser` 的 UID 运行，此容器将出现问题，因为应用程序将无法写入 `/myapp-tmp-dir` 文件夹。

Don’t use a hardcoded path only writable by `myuser`. Instead, write temporary data to `/tmp` (where any user can write, thanks to the sticky bit permissions). Make  resources world readable (i.e., 0644 instead of 0640), and ensure that  everything works if the UID is changed.

不要使用只能由 myuser 写入的硬编码路径。相反，将临时数据写入`/tmp`（由于粘滞位权限，任何用户都可以写入）。使资源世界可读（即，0644 而不是 0640），并确保在 UID 更改时一切正常。

```
...
USER myuser
ENV APP_TMP_DATA=/tmp
ENTRYPOINT ["/myapp"]
```

In this example our application will use the path in `APP_TMP_DATA` environment variable. The default value `/tmp` will allow the application to execute as any UID and still write temporary data to `/tmp`. Having the path as a configurable environment variable is not always  necessary, but it will make things easier when setting up and mounting  volumes for persistence.

在此示例中，我们的应用程序将使用“APP_TMP_DATA”环境变量中的路径。默认值 `/tmp` 将允许应用程序作为任何 UID 执行，并且仍然将临时数据写入 `/tmp`。将路径作为可配置的环境变量并不总是必要的，但在设置和安装卷以实现持久性时，它会使事情变得更容易。

### #1.3 Make executables owned by root and not writable

### #1.3 使 root 拥有的可执行文件不可写

It is a Dockerfile best practice for every executable in a container to  be owned by the root user, even if it is executed by a non-root user and should not be world-writable.

容器中的每个可执行文件都由 root 用户拥有是 Dockerfile 的最佳实践，即使它由非 root 用户执行并且不应该是全局可写的。

This will block the executing user from modifying existing binaries or  scripts, which could enable different attacks. By following this best  practice, you’re[ effectively enforcing container immutability](https://cloud.google.com/solutions/best-practices-for-operating-containers#immutability). Immutable containers do not update their code automatically at runtime  and, in this way, you can prevent your running application from being  accidentally or maliciously modified.

这将阻止执行用户修改现有的二进制文件或脚本，这可能会导致不同的攻击。通过遵循此最佳实践，您可以[有效地实施容器不变性](https://cloud.google.com/solutions/best-practices-for-operating-containers#immutability)。不可变容器不会在运行时自动更新其代码，通过这种方式，您可以防止正在运行的应用程序被意外或恶意修改。

To follow this best practice, try to avoid:

要遵循此最佳实践，请尽量避免：

```
...
WORKDIR $APP_HOME
COPY --chown=app:app app-files/ /app
USER app
ENTRYPOINT /app/my-app-entrypoint.sh
```

Most of the time, you can just drop the `--chown app:app` option (or `RUN chown ... `commands). The *app* user only needs execution permissions on the file, **not ownership**.

大多数情况下，您可以删除 `--chown app:app` 选项（或 `RUN chown ... `commands）。 *app* 用户只需要文件的执行权限，**不需要所有权**。

## #2 Reduce attack surface

## #2 减少攻击面

It is a Dockerfile best practice to **keep the images minimal**. 

**保持图像最小**是 Dockerfile 的最佳实践。

Avoid including unnecessary packages or exposing ports to reduce the  attack surface. The more components you include inside a container, the  more exposed your system will be and the harder it is to maintain,  especially for components not under your control.

避免包含不必要的包或暴露端口以减少攻击面。您在容器中包含的组件越多，您的系统暴露的就越多，维护起来就越困难，尤其是对于不受您控制的组件。

### #2.1 Multistage builds

### #2.1 多阶段构建

Make use of [multistage building](https://docs.docker.com/develop/develop-images/multistage-build/) features to have reproducible builds inside containers.

利用 [多阶段构建](https://docs.docker.com/develop/develop-images/multistage-build/) 功能在容器内进行可重现的构建。

In a multistage build, you create an intermediate container – or stage – with all the required tools to compile or produce your final artifacts  (i.e., the final executable). Then, you copy **only the resulting** artifacts to the final image, without additional development dependencies, temporary build files, etc.

在多阶段构建中，您创建一个中间容器（或阶段），其中包含编译或生成最终工件（即最终可执行文件）所需的所有工具。然后，您将**仅生成的** 工件复制到最终映像，而无需额外的开发依赖项、临时构建文件等。

A well crafted multistage build includes only the minimal required  binaries and dependencies in the final image, and not build tools or  intermediate files. This reduces the attack surface, decreasing  vulnerabilities.

精心设计的多阶段构建仅包含最终映像中所需的最少二进制文件和依赖项，而不包含构建工具或中间文件。这减少了攻击面，减少了漏洞。

It is safer, and it also reduces image size.

它更安全，并且还减小了图像大小。

For a go application, an example of a multistage build would look like this:

对于 go 应用程序，多阶段构建的示例如下所示：

```
#This is the "builder" stage
FROM golang:1.15 as builder
WORKDIR /my-go-app
COPY app-src .
RUN GOOS=linux GOARCH=amd64 go build ./cmd/app-service
#This is the final stage, and we copy artifacts from "builder"
FROM gcr.io/distroless/static-debian10
COPY --from=builder /my-go-app/app-service /bin/app-service
ENTRYPOINT ["/bin/app-service"]
```

With those Dockerfile instructions, we create a *builder* stage using the golang:1.15 container, which includes all of the go toolchain.

通过这些 Dockerfile 指令，我们使用 golang:1.15 容器创建了一个 *builder* 阶段，其中包括所有 go 工具链。

```
FROM golang:1.15 as builder
```

We can copy the source code in there and build.

我们可以在那里复制源代码并构建。

```
WORKDIR /my-go-app
COPY app-src .
RUN GOOS=linux GOARCH=amd64 go build ./cmd/app-service
```

Then, we define another stage based on a Debian distroless image (see next tip).

然后，我们根据 Debian distroless 映像定义另一个阶段（请参阅下一个技巧）。

```
FROM gcr.io/distroless/static-debian10
```

`COPY` the resulting executable from the *builder* stage using the `--from=builder flag`.

使用`--from=builder 标志`从*builder* 阶段`COPY` 生成的可执行文件。

```
COPY --from=builder /my-go-app/app-service /bin/app-service
```

The final image will contain only the minimal set of libraries from distroless/static-debian-10 image and your app executable.

最终映像将仅包含来自 distroless/static-debian-10 映像和您的应用程序可执行文件的最小库集。

No build toolchain, no source code.

没有构建工具链，没有源代码。

We recommend you check this [NodeJS application example](https://github.com/Coffee-WIP/coffeewip-website/blob/master/Dockerfile) or this [efficient Python with Django multi-stage build](https://blog.ploetzli.ch/2020/efficient-multi-stage-build-django-docker/).

我们建议您查看此 [NodeJS 应用程序示例](https://github.com/Coffee-WIP/coffeewip-website/blob/master/Dockerfile) 或此 [带 Django 多阶段构建的高效 Python](https://blog.ploetzli.ch/2020/efficient-multi-stage-build-django-docker/)。

### #2.2 Distroless, from scratch

### #2.2 Distroless，从头开始

Use the minimal required base container to follow Dockerfile best practices.

使用所需的最少基础容器来遵循 Dockerfile 最佳实践。

Ideally, we would create containers from scratch, but only binaries that are 100% static will work.

理想情况下，我们会从头开始创建容器，但只有 100% 静态的二进制文件才能工作。

[Distroless](https://github.com/GoogleContainerTools/distroless) are a nice alternative. These are designed to contain only the minimal  set of libraries required to run Go, Python, or other frameworks.

[Distroless](https://github.com/GoogleContainerTools/distroless) 是一个不错的选择。它们旨在仅包含运行 Go、Python 或其他框架所需的最少库集。

For example, if you were to base a container in a generic `ubuntu:xenial` image:

例如，如果您要在通用的 `ubuntu:xenial` 映像中创建一个容器：

```
FROM ubuntu:xenial-20210114
```

You would include more than 100 vulnerabilities, as detected by [Sysdig inline scanner](https://sysdig.com/products/secure/image-scanning/), related to the large amount of packages that you are including and probably neither need nor ever use:

您将包含 100 多个漏洞，这些漏洞由 [Sysdig 内联扫描仪](https://sysdig.com/products/secure/image-scanning/) 检测到，与您包含的大量软件包相关，并且可能都不需要也不曾使用：

```
❯ docker run -v /var/run/docker.sock:/var/run/docker.sock --rm quay.io/sysdig/secure-inline-scan:2 image-ubuntu -k $SYSDIG_SECURE_TOKEN --storage-typedocker-daemon
Inspecting image from Docker daemon -- distroless-1:latest
  Full image:  docker.io/library/image-ubuntu
  Full tag:    localbuild/distroless-1:latest
…
Analyzing image…
Analysis complete!
...
Evaluation results
 - warn dockerfile:instruction Dockerfile directive 'HEALTHCHECK' not found, matching condition 'not_exists' check
 - warn dockerfile:instruction Dockerfile directive 'USER' not found, matching condition 'not_exists' check
 - warn files:suid_or_guid_set SUID or SGID found set on file /bin/mount.Mode: 0o104755
 - warn files:suid_or_guid_set SUID or SGID found set on file /bin/su.Mode: 0o104755
 - warn files:suid_or_guid_set SUID or SGID found set on file /bin/umount.Mode: 0o104755
 - warn files:suid_or_guid_set SUID or SGID found set on file /sbin/pam_extrausers_chkpwd.Mode: 0o102755
 - warn files:suid_or_guid_set SUID or SGID found set on file /sbin/unix_chkpwd.Mode: 0o102755
 - warn files:suid_or_guid_set SUID or SGID found set on file /usr/bin/chage.Mode: 0o102755
…
Vulnerabilities report
   Vulnerability    Severity Package                                  Type     Fix version      URL
 - CVE-2019-18276   Low      bash-4.3-14ubuntu1.4                     dpkg     None             http://people.ubuntu.com/~ubuntu-security/cve/CVE-2019-18276
 - CVE-2016-2781    Low      coreutils-8.25-2ubuntu3~16.04            dpkg     None             http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2781
 - CVE-2017-8283    Negligible dpkg-1.18.4ubuntu1.6                     dpkg     None             http://people.ubuntu.com/~ubuntu-security/cve/CVE-2017-8283
 - CVE-2020-13844   Medium   gcc-5-base-5.4.0-6ubuntu1~16.04.12       dpkg     None             http://people.ubuntu.com/~ubuntu-security/cve/CVE-2020-13844
...
 - CVE-2018-20839   Medium   systemd-sysv-229-4ubuntu21.29            dpkg     None             http://people.ubuntu.com/~ubuntu-security/cve/CVE-2018-20839
 - CVE-2016-5011    Low      util-linux-2.27.1-6ubuntu3.10            dpkg     None             http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-5011
```

Do you need the *gcc* compiler or *systemd SysV* compatibility in your container? Most likely, you don’t. The same goes for *dpkg* or bash.

您的容器中是否需要 *gcc* 编译器或 *systemd SysV* 兼容性？很可能，你没有。 *dpkg* 或 bash 也是如此。

If you base your image on [gcr.io/distroless/base-debian10](https://github.com/GoogleContainerTools/distroless/tree/master/base):

如果您的图像基于 [gcr.io/distroless/base-debian10](https://github.com/GoogleContainerTools/distroless/tree/master/base)：

```
FROM gcr.io/distroless/base-debian10
```

Then it will only contain a basic set of packages, including just required libraries like *glibc*, *libssl,* and *openssl*.

然后它将只包含一组基本的包，包括只需要的库，如 *glibc*、*libssl、* 和 *openssl*。

For statically compiled applications like Go that don’t require *libc*, you can even go with the slimmer:

对于像 Go 这样不需要 *libc* 的静态编译应用程序，你甚至可以使用更精简的：

```
FROM gcr.io/distroless/static-debian10
```

### #2.3 Use trusted base images

### #2.3 使用受信任的基础镜像

Carefully choose the base for your images (the `FROM` instruction).

仔细选择图像的基础（`FROM` 指令）。

Building on top of untrusted or unmaintained images will inherit all of  the problems and vulnerabilities from that image into your containers.

建立在不受信任或不受维护的镜像之上会将所有问题和漏洞从该镜像继承到您的容器中。

Follow these Dockerfile best practices to select your base images:

遵循这些 Dockerfile 最佳实践来选择您的基础镜像：

- You should prefer *verified* and *official* **images from trusted repositories** and providers over images built by unknown users.
- When using custom images, check for the image source and the Dockerfile, and **build your own base image**. There is no guarantee that an image published in a public registry is  really built from the given Dockerfile. Neither is assurance that it is  kept up to date.
- Sometimes the *official* images might not be the **better fit**, in regards to security and minimalism. For example, comparing the [official node image](https://hub.docker.com/_/node) with the [bitnami/node](https://hub.docker.com/r/bitnami/node/) image, the latter offers customized versions on top of a minideb  distribution. They are frequently updated with the latest bug fixes,  signed with *Docker Content Trust*, and pass a [security scan for tracking known vulnerabilities](https://quay.io/repository/bitnami/node?tab=tags).

- 您应该更喜欢*已验证* 和*官方* **来自受信任存储库** 和提供者的图像，而不是由未知用户构建的图像。
- 使用自定义镜像时，请检查镜像源和 Dockerfile，并**构建您自己的基础镜像**。无法保证在公共注册表中发布的映像确实是从给定的 Dockerfile 构建的。也不能保证它是最新的。
- 在安全性和极简主义方面，有时*官方*图像可能不是**更合适**。比如比较 [官方节点镜像](https://hub.docker.com/_/node)和[bitnami/node](https://hub.docker.com/r/bitnami/node/)图像，后者在 minideb 发行版之上提供定制版本。它们经常更新最新的错误修复，使用 *Docker Content Trust* 签名，并通过 [用于跟踪已知漏洞的安全扫描](https://quay.io/repository/bitnami/node?tab=tags)。

![Diagram of the layer tree of an image.If one layer is compromised, all the following layers will probably be compromised as well](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-03-image-layer-vulnerabilities-inherited.png)

如果某一层被攻破，接下来的所有层都可能被攻破](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-03-image-layer-vulnerabilities-inherited.png)

### #2.4 Update your images frequently

### #2.4 经常更新你的图片

Use base images that are frequently updated, and rebuild yours on top of them.

使用经常更新的基础镜像，并在它们之上重建你的镜像。

As new security vulnerabilities are discovered continuously, it is a  general security best practice to stick to the latest security patches.

随着新的安全漏洞不断被发现，坚持使用最新的安全补丁是一种通用的安全最佳实践。

There is no need to always go to the latest version, which might contain breaking changes, but define a versioning strategy:

无需总是使用最新版本，其中可能包含重大更改，但需要定义版本控制策略：

- **Stick to stable** or long-term support versions, which deliver security fixes soon and often.
- **Plan in advance**. Be ready to drop old versions and migrate before your base image version reaches the end of its life  and stops receiving updates.
- Also, **rebuild your own images periodically** and with a similar strategy to get the latest packages from the base  distro, Node, Golang, Python, etc. Most package or dependency managers,  like [npm](https://docs.npmjs.com/cli/v6/configuring-npm/package-json#dependencies) or [go mod](https://golang.org/ref/mod), will offer ways to specify version ranges to keep up with latest security updates.

- **坚持稳定**或长期支持版本，这些版本会很快并经常提供安全修复。
- **提前计划**。准备好在您的基本映像版本达到其生命周期结束并停止接收更新之前删除旧版本并迁移。
- 此外，**定期重建你自己的镜像**并使用类似的策略从基础发行版、Node、Golang、Python 等获取最新的包。大多数包或依赖项管理器，如 [npm](https://docs.npmjs.com/cli/v6/configuring-npm/package-json#dependencies) 或 [go mod](https://golang.org/ref/mod)，将提供指定版本范围的方法以跟上最新的安全更新。

### #2.5 Exposed ports

### #2.5 暴露端口

Every opened port in your container is an open door to your system. Expose only the ports that your application needs and avoid exposing  ports like SSH (22).

容器中每个打开的端口都是通往系统的大门。仅公开您的应用程序需要的端口，并避免公开 SSH (22) 等端口。

Please note that even though the Dockerfile offers the [EXPOSE command](https://docs.docker.com/engine/reference/builder/#expose), this command is only informational and for documentation purposes. Exposing the port does not automatically allow connections for all  EXPOSED ports when running the container (unless you use `docker run --publish-all`). You need to specify the published ports at runtime, when executing the container.

请注意，尽管 Dockerfile 提供了 [EXPOSE 命令](https://docs.docker.com/engine/reference/builder/#expose)，但此命令仅用于提供信息并用于文档目的。运行容器时，公开端口不会自动允许所有 EXPOSED 端口的连接（除非您使用 `docker run --publish-all`)。执行容器时，您需要在运行时指定发布的端口。

Use EXPOSE to flag and document only the required ports in the  Dockerfile, and then stick to those ports when publishing or exposing in execution.

使用 EXPOSE 仅标记和记录 Dockerfile 中所需的端口，然后在发布或执行时坚持这些端口。

## #3 Prevent confidential data leaks

## #3 防止机密数据泄露

Be really careful about your confidential data when dealing with containers.

处理容器时要非常小心您的机密数据。

The following Dockerfile best practices will provide some advice on  handling credentials for containers, and how to avoid accidentally  leaking undesired files or information.

以下 Dockerfile 最佳实践将提供一些关于处理容器凭据的建议，以及如何避免意外泄漏不需要的文件或信息。

### #3.1 Credentials and confidentiality

### #3.1 凭证和机密性

Never put any secret or credentials in the Dockerfile instructions  (environment variables, args, or hard coded into any command). 

切勿在 Dockerfile 指令（环境变量、参数或硬编码到任何命令中）中放入任何机密或凭据。

Be extra careful with files that get copied into the container. Even if a file is removed in a later instruction in the Dockerfile, it can still  be accessed on the previous layers as it is not really removed, only  “hidden” in the final filesystem. So, when building your images, follow  these practices:

对复制到容器中的文件要格外小心。即使文件在 Dockerfile 的后续指令中被删除，它仍然可以在前面的层上访问，因为它并没有真正被删除，只是“隐藏”在最终文件系统中。因此，在构建图像时，请遵循以下做法：

- If the application supports **configuration via environment variables**, use them to set the secrets on execution (-e option in docker run), or use [Docker secrets](https://docs.docker.com/engine/swarm/secrets/), [Kubernetes secrets](https://kubernetes.io/docs/concepts/configuration/secret/) to provide the values as environment variables.
- **Use configuration files** and [bind mount](https://docs.docker.com/storage/bind-mounts/) the configuration files in docker, or [mount them from a Kubernetes secret](https://kubernetes.io/docs/concepts/storage/volumes/#secret).

- 如果应用程序支持**通过环境变量配置**，使用它们来设置执行时的秘密（docker run 中的 -e 选项），或使用 [Docker secrets](https://docs.docker.com/engine/swarm/secrets/), [Kubernetes secrets](https://kubernetes.io/docs/concepts/configuration/secret/) 提供值作为环境变量。
- **使用配置文件**和[绑定挂载](https://docs.docker.com/storage/bind-mounts/)docker中的配置文件，或者[从Kubernetes秘密挂载它们](https://kubernetes.io/docs/concepts/storage/volumes/#secret)。

Also, **your images shouldn’t contain confidential information** or configuration values that tie them to some specific environment (i.e., production, staging, etc.).

此外，**您的图像不应包含机密信息**或将它们与某些特定环境（即生产、暂存等）联系起来的配置值。

Instead, allow the image to be customized by **injecting the values on runtime**, especially secrets. You should only include configuration files with safe or dummy values inside, as an example.

相反，允许通过**在运行时注入值**来自定义图像，尤其是机密。例如，您应该只包含其中包含安全值或虚拟值的配置文件。

### #3.2 ADD, COPY

### #3.2 添加，复制

Both the ADD and COPY instructions provide similar functions in a Dockerfile. However, COPY is more explicit.

ADD 和 COPY 指令在 Dockerfile 中提供类似的功能。但是，COPY 更为明确。

Use COPY unless you really need the ADD functionality, like to add files from an URL or from a tar file. COPY is more predictable and less error prone.

除非您确实需要 ADD 功能，否则请使用 COPY，例如从 URL 或从 tar 文件添加文件。 COPY 更具可预测性且不易出错。

In some cases it is preferred to use the RUN instruction over ADD to download a package using *curl* or *wget*, extract it, and then remove the original file in a single step, reducing the number of layers.

在某些情况下，最好使用 RUN 指令而不是 ADD 来下载使用 *curl* 或 *wget* 的包，解压缩它，然后一步删除原始文件，减少层数。

Multistage builds also solve this problem and help you follow Dockerfile best practices, allowing you to copy only the final extracted files  from a previous stage.

多阶段构建也解决了这个问题，并帮助您遵循 Dockerfile 最佳实践，允许您仅复制从前一阶段最终提取的文件。

### #3.3 Build context and dockerignore

### #3.3 构建上下文和 dockerignore

Here is a typical execution of a build using docker, with a default *Dockerfile*, and the context in the current folder:

这是使用 docker 执行构建的典型执行，具有默认的 *Dockerfile*，以及当前文件夹中的上下文：

```
docker build -t myimage .
```

**Beware!**

**谨防！**

The “.” parameter is the build context. Using “.” as context is  dangerous as you can copy confidential or unnecessary files into the  container, like configuration files, credentials, backups, lock files,  temporary files, sources, subfolders, dotfiles, etc.

这 ”。”参数是构建上下文。使用 ”。”因为上下文是危险的，因为您可以将机密或不必要的文件复制到容器中，例如配置文件、凭据、备份、锁定文件、临时文件、源、子文件夹、点文件等。

Imagine that you have the following command inside the Dockerfile:

假设您在 Dockerfile 中有以下命令：

```
COPY ./my-app
```

This would copy **everything** inside the build context, which for the “.” example, includes the Dockerfile itself.

这将复制构建上下文中的**所有内容**，对于“.”例如，包括 Dockerfile 本身。

It would be Dockerfile best practices to create a subfolder containing  the files that need to be copied inside the container, use it as the  build context, and when possible, be explicit for the COPY instructions  (avoid wildcards). For example:

Dockerfile 的最佳实践是创建一个包含需要在容器内复制的文件的子文件夹，将其用作构建上下文，并在可能的情况下明确 COPY 指令（避免使用通配符）。例如：

```
docker build -t myimage files/
```

Also, create a `.dockerignore` file to explicitly exclude files and directories.

此外，创建一个 `.dockerignore` 文件以明确排除文件和目录。

Even if you are extra careful with the `COPY` instructions,  all of the build context is sent to the docker daemon before starting  the image build. That means having a smaller and restricted build  context will make your builds faster.

即使您对 `COPY` 指令格外小心，在开始镜像构建之前，所有构建上下文都会发送到 docker 守护进程。这意味着拥有更小且受限制的构建上下文将使您的构建更快。

Put your build context in its own folder and use `.dockerignore` to reduce it as much as possible.

将你的构建上下文放在它自己的文件夹中，并使用 `.dockerignore` 尽可能地减少它。

## #4 Others

## #4 其他

### #4.1 Layer sanity

### #4.1 层的完整性

Remember that order in the Dockerfile instructions is very important.

请记住，Dockerfile 指令中的顺序非常重要。

Since RUN, COPY, ADD, and other instructions will create a new container layer, grouping multiple commands together will reduce the number of  layers.

由于 RUN、COPY、ADD 和其他指令将创建一个新的容器层，将多个命令组合在一起将减少层数。

For example, instead of:

例如，而不是：

```
FROM ubuntu
RUN apt-get install -y wget
RUN wget https://…/downloadedfile.tar
RUN tar xvzf downloadedfile.tar
RUN rm downloadedfile.tar
RUN apt-get remove wget
```

It would be a Dockerfile best practice to do:

这将是 Dockerfile 的最佳实践：

```
FROM ubuntu
RUN apt-get install wget && wget https://…/downloadedfile.tar && tar xvzf downloadedfile.tar && rm downloadedfile.tar && apt-get remove wget
```

Also, place the commands that are less likely to change, and easier to cache, first.

此外，首先放置不太可能更改且更易于缓存的命令。

Instead of:

代替：

```
FROM ubuntu
COPY source/* .
RUN apt-get install nodejs
ENTRYPOINT ["/usr/bin/node", "/main.js"]
```

It would be better to do:

最好这样做：

```
FROM ubuntu
RUN apt-get install nodejs
COPY source/* .
ENTRYPOINT ["/usr/bin/node", "/main.js"]
```

The `nodejs` package is less likely to change than our application source. 

`nodejs` 包比我们的应用程序源不太可能改变。

Please remember that executing a `rm` command removes the  file on the next layer, but it is still available and can be accessed,  as the final image filesystem is composed from all the previous layers.

请记住，执行 `rm` 命令会删除下一层的文件，但它仍然可用并且可以访问，因为最终的图像文件系统由所有先前的层组成。

So **don’t copy confidential files and then remove them**, they will be not visible in the final container filesystem but still be easily accessible.

所以**不要复制机密文件然后删除它们**，它们在最终容器文件系统中将不可见，但仍然可以轻松访问。

### #4.2 Metadata labels

### #4.2 元数据标签

It is a Dockerfile best practice to include metadata labels when building your image.

在构建映像时包含元数据标签是 Dockerfile 的最佳实践。

Labels will help in image management, like including the application  version, a link to the website, how to contact the maintainer, and more.

标签将有助于图像管理，例如包括应用程序版本、网站链接、如何联系维护者等等。

You can take a look at the [predefined annotations from the OCI image spec](https://github.com/opencontainers/image-spec/blob/master/annotations.md), which deprecate the previous [Label schema standard draft] (http://label-schema.org/rc1/).

你可以看看[来自OCI图像规范的预定义注释](https://github.com/opencontainers/image-spec/blob/master/annotations.md)，它弃用了之前的[标签模式标准草案](http://label-schema.org/rc1/)。

### #4.3 Linting

### #4.3 Linting

Tools like [Haskell Dockerfile Linter (hadolint)](https://github.com/hadolint/hadolint) can detect bad practices in your Dockerfile, and even expose issues inside the shell commands executed by the `RUN` instruction.

像 [Haskell Dockerfile Linter (hadolint)](https://github.com/hadolint/hadolint) 之类的工具可以检测 Dockerfile 中的不良做法，甚至可以暴露由 `RUN` 指令执行的 shell 命令中的问题。

Consider incorporating such a tool in your CI pipelines.

考虑在您的 CI 管道中加入这样一个工具。

Image scanners are also capable of detecting bad practices via  customizable rules, and report them along with image vulnerabilities:

图像扫描仪还能够通过可定制的规则检测不良做法，并将其与图像漏洞一起报告：

![Image scanning policies in Sysdig Secure.You can create gates to check for misconfigurations in the Dockerfile.](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-04-image-scanning-policies.png)

您可以创建门来检查 Dockerfile 中的错误配置。](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-04-image-scanning-policies.png)

Some of the misconfigurations you can detect are images running as root, exposed ports, usage of the `ADD` instruction, hardcoded secrets, or discouraged `RUN` commands.

您可以检测到的一些错误配置包括以 root 身份运行的映像、暴露的端口、“ADD”指令的使用、硬编码的秘密或不鼓励的“RUN”命令。

### #4.4 Locally scan images during development

### #4.4 在开发过程中本地扫描图像

Image scanning is another way of detecting potential problems before running your containers. In order to follow the [image scanning best practices](https://sysdig.com/blog/image-scanning-best-practices/), you should perform the scanning at different stages of the image life  cycle, in addition to when the image is already pushed to a container  registry.

图像扫描是在运行容器之前检测潜在问题的另一种方法。为了遵循[图像扫描最佳实践](https://sysdig.com/blog/image-scanning-best-practices/)，您应该在图像生命周期的不同阶段执行扫描，除了何时映像已推送到容器注册表。

It is a security best practice to apply the “shift left security”  paradigm by directly scanning your images, as soon as they are built, in your CI pipelines before pushing to the registry.

应用“左移安全”范式是一种安全最佳实践，即在镜像构建后立即在 CI 管道中直接扫描镜像，然后再推送到注册表。

This also includes **in the developer computer**, *using the [Sysdig inline scanner](https://docs.sysdig.com/en/integrate-with-ci-cd-tools.html), which provides different integrations with CI/CD tools like [Jenkins](https://plugins.jenkins.io/sysdig-secure/), [Github actions](https://sysdig.com/blog/image-scanning-github-actions/) , and more.*

这还包括**在开发者计算机中**，*使用 [Sysdig 内联扫描仪](https://docs.sysdig.com/en/integrate-with-ci-cd-tools.html)，它提供了不同的集成使用 CI/CD 工具，如 [Jenkins](https://plugins.jenkins.io/sysdig-secure/)、[Github操作](https://sysdig.com/blog/image-scanning-github-actions/) ， 和更多。*

And remember, a scanned image might be “safe” now. But as it ages and  new vulnerabilities are discovered, it might become dangerous.

请记住，扫描图像现在可能是“安全的”。但随着它的老化和新的漏洞被发现，它可能会变得危险。

Periodically **reevaluate for new vulnerabilities**.

定期**重新评估新漏洞**。

![Diagram of a vulnerability timeline.A vulnerability exists in software before they are detected. You may deploy images that are vulnerable, the vulnerability may be discovered later. This means your were vulnerable all the time, so you need to continuosly scan your images running in production to protect from these newly discovered vulnerabilities.](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-05-runtime-vulnerability-timeline.png)

在检测到漏洞之前，软件中存在漏洞。您可能会部署易受攻击的映像，该漏洞可能会在以后被发现。这意味着您一直很容易受到攻击，因此您需要不断扫描生产中运行的图像以防止这些新发现的漏洞。](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-05-runtime-vulnerability-timeline.png)

## #5 Beyond image building

## #5 超越图像构建

So far, we have focused on the image building process and discussed tips for creating optimal Dockerfiles. But let’s not forget about some  additional pre-checks and what comes after building your image: running  it.

到目前为止，我们一直专注于镜像构建过程，并讨论了创建最佳 Dockerfile 的技巧。但我们不要忘记一些额外的预检查以及构建映像后的内容：运行它。

![img](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-06-container-image-lifecycle.png)

### #5.1 Docker port socket and TCP protection

### #5.1 Docker 端口套接字和 TCP 保护

The docker socket is a big privileged door into your host system that, [as seen recently, can be used for intrusion and malicious software usage](https://sysdig.com/blog/mitigating-weave-scope/). Make sure your `/var/run/docker.sock` has the correct permissions, and if docker is exposed via TCP (which is not recommended at all), make sure [it is properly protected](https://docs.docker.com/engine/security/https/).

docker 套接字是进入您的主机系统的一扇大特权门，[正如最近所见，可用于入侵和恶意软件使用](https://sysdig.com/blog/mitigating-weave-scope/)。确保您的 `/var/run/docker.sock` 具有正确的权限，如果 docker 通过 TCP 公开（完全不推荐），请确保 [它受到适当保护](https://docs.docker.com/engine/security/https/)。

### #5.2 Sign images and verify signatures 

### #5.2 对图像进行签名并验证签名

It is one of the Dockerfile best practices to use [docker content trust](https://docs.docker.com/engine/security/trust/), Docker notary, Harbor notary, or similar tools to **digitally sign your images ** and then **verify them on runtime**.

使用 [docker content trust](https://docs.docker.com/engine/security/trust/)、Docker notary、Harbor notary 或类似工具对您的图像进行**数字签名是 Dockerfile 最佳实践之一** 然后 ** 在运行时验证它们**。

Enabling signature verification is different on each runtime. For example, in docker this is done with the `DOCKER_CONTENT_TRUST` environment variable:
`export DOCKER_CONTENT_TRUST=1`

在每个运行时启用签名验证是不同的。例如，在 docker 中，这是通过 `DOCKER_CONTENT_TRUST` 环境变量完成的：
`导出DOCKER_CONTENT_TRUST=1`

### #5.3 Tag mutability

### #5.3 标签可变性

In container land, tags are a volatile reference to a concrete image version in a specific point in time.

在容器领域，标签是在特定时间点对具体图像版本的易失性引用。

![If you use mutant tags, you might be scanning one version, but deploying another one.](https://sysdig.com/wp-content/uploads/Dockerfile-best-practices-07-tag_mutability_pointers.png)

Tags can change unexpectedly, and at any moment. See our “[Attack of the mutant tags](https://sysdig.com/blog/toctou-tag-mutability/)” to learn more.

标签可能会随时发生意外变化。请参阅我们的“[突变标签的攻击](https://sysdig.com/blog/toctou-tag-mutability/)”了解更多信息。

### #5.4 Run as non root

### #5.4 以非 root 身份运行

Previously, we talked about using a non-root user when building a  container. The USER instruction will set the default user for the  container, but the orchestrator or runtime environment (i.e., docker  run, kubernetes, etc.) has the last word on who is the running container effective user.

之前，我们谈到在构建容器时使用非 root 用户。 USER 指令将为容器设置默认用户，但编排器或运行时环境（即 docker run、kubernetes 等）对谁是正在运行的容器有效用户拥有最终决定权。

Really **avoid running your environment as root**.

真的**避免以 root 身份运行您的环境**。

Openshift and some Kubernetes clusters will apply restrictive policies  by default, preventing root containers from running. Avoid the  temptation of running as root to circumvent permission or ownership  issues, and **fix the real problem** instead.

Openshift 和一些 Kubernetes 集群默认会应用限制性策略，阻止根容器运行。避免以 root 身份运行以规避权限或所有权问题的诱惑，而是**解决真正的问题**。

### #5.5 Include health / liveness checks

### #5.5 包括健康/活性检查

When using plain Docker or Docker Swarm, [include a HEALTHCHECK instruction](https://docs.docker.com/engine/reference/builder/#healthcheck) in your Dockerfile whenever possible. This is critical for long running or persistent services in order to ensure they are healthy, and manage  restarting the service otherwise.

使用普通 Docker 或 Docker Swarm 时，请尽可能在 Dockerfile 中 [包括 HEALTHCHECK 指令](https://docs.docker.com/engine/reference/builder/#healthcheck)。这对于长时间运行或持久的服务至关重要，以确保它们健康，否则管理重新启动服务。

If running your images in Kubernetes, [use livenessProbe configuration](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) inside the container definitions, as the docker HEALTHCHECK instruction won't be applied.

如果在 Kubernetes 中运行你的镜像，[使用 livenessProbe 配置](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) 在容器定义中，作为 docker HEALTHCHECK 指令不会被应用。

### #5.6 Drop capabilities

### #5.6 删除功能

Also in execution, you can **restrict the application capabilities** to the minimal required set using `--cap-drop flag` in Docker or[ securityContext.capabilities.drop in Kubernetes](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-capabilities-for-a-container). That way, in case your container is compromised, the range of action available to an attacker is limited.

同样在执行过程中，您可以在 Docker 中使用 `--cap-drop 标志` 或 [Kubernetes 中的 securityContext.capabilities.drop](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-capabilities-for-a-container)。这样，万一您的容器遭到破坏，攻击者可采取的行动范围就会受到限制。

Also, see more information on how to apply AppArmor and Seccomp as additional mechanisms to restrict container privileges:

此外，请参阅有关如何应用 AppArmor 和 Seccomp 作为限制容器权限的附加机制的更多信息：

- AppArmor in [Docker](https://docs.docker.com/engine/security/apparmor/) or [Kubernetes](https://sysdig.com/blog/manage-apparmor-profiles-in-kubernetes-with-kube-apparmor-manager/)
- Seccomp in [Docker](https://docs.docker.com/engine/security/seccomp/) or [Kubernetes](https://kubernetes.io/docs/tutorials/clusters/seccomp/).

- [Docker](https://docs.docker.com/engine/security/apparmor/) 或 [Kubernetes](https://sysdig.com/blog/manage-apparmor-profiles-in-kubernetes-with 中的 AppArmor -kube-apparmor-manager/)
- [Docker](https://docs.docker.com/engine/security/seccomp/) 或 [Kubernetes](https://kubernetes.io/docs/tutorials/clusters/seccomp/) 中的 Seccomp。

## Conclusion

##  结论

We have seen that container image security is a complex and critical  topic that simply cannot be ignored until it explodes with terrible  consequences.

我们已经看到，容器镜像安全是一个复杂而关键的话题，在它爆发并带来可怕后果之前，不能忽视它。

**Prevention and shifting security left is essential** for improving your security posture and reducing the management overhead.

**预防和转移安全性对于改善您的安全状况和减少管理开销至关重要**。

This set of recommendations, focused on Dockerfiles best practices, will help you in this mission.

这组建议侧重于 Dockerfiles 最佳实践，将帮助您完成这项任务。

*If you want to go a step further, check also our [12 container image scanning best practices](https://sysdig.com/blog/image-scanning-best-practices/) article, to help you shift left security. *

*如果您想更进一步，请查看我们的 [12 容器镜像扫描最佳实践](https://sysdig.com/blog/image-scanning-best-practices/) 文章，以帮助您左移安全。 *

The [image scanning feature in Sysdig Secure](https://sysdig.com/products/secure/image-scanning/) will help you follow these Dockerfile best practices. It will help you  shift left security by checking for vulnerabilities and  misconfigurations, allowing you to act before threats are deployed. You’ll be set in only a few minutes. [Try it today!](https://sysdig.com/company/free-trial-platform/) 

[Sysdig Secure 中的图像扫描功能](https://sysdig.com/products/secure/image-scanning/) 将帮助您遵循这些 Dockerfile 最佳实践。它将通过检查漏洞和错误配置来帮助您转移安全性，让您在部署威胁之前采取行动。您只需几分钟即可完成设置。 [今天就试试！](https://sysdig.com/company/free-trial-platform/)

