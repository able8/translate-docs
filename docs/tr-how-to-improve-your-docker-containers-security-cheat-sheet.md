# How to improve your Docker containers security [cheat sheet included]

# 如何提高 Docker 容器的安全性 [包括备忘单]

Containers are no security devices. That's why we've curated a set of easily actionable recommendations to improve your Docker containers security. Check out the one-page cheat sheet.

容器不是安全设备。这就是为什么我们策划了一组易于操作的建议来提高您的 Docker 容器安全性。查看一页备忘单。

## Thomas Segura

## 托马斯塞古拉

Thomas' passion for tech and open-source led him to join GitGuardian as technical content writer. He focuses on clarifying the transformative changes that cybersecurity and software are going through.

Thomas 对技术和开源的热情促使他加入 GitGuardian，担任技术内容作家。他专注于阐明网络安全和软件正在经历的变革。

#### [Thomas Segura](http://blog.gitguardian.com/author/thomas/)

#### [托马斯塞古拉](http://blog.gitguardian.com/author/thomas/)

30 Jul 2021• 7 min read

2021 年 7 月 30 日• 阅读 7 分钟

###### Table of contents

#### ##  目录

Docker containers have been an essential part of the developer's toolbox for several years now, allowing them to build, distribute and deploy their applications in a standardized way.

多年来，Docker 容器一直是开发人员工具箱的重要组成部分，使他们能够以标准化的方式构建、分发和部署他们的应用程序。

This gain in traction has been, not surprisingly, accompanied by **a surge in security issues** related to containerization technology. Indeed, containers also represent a standardized surface for attackers. They can easily exploit [misconfigurations](https://blog.gitguardian.com/hunting-for-secrets-in-docker-hub/) and **escape from within containers to the host machine**.

毫无疑问，这种吸引力的增加伴随着与容器化技术相关的**安全问题**激增。事实上，容器也代表了攻击者的标准化表面。他们可以轻松利用 [错误配置](https://blog.gitguardian.com/hunting-for-secrets-in-docker-hub/) 并**从容器内逃逸到主机**。

Furthermore, the word “container” is often misunderstood, as many developers tend to **associate the concept of isolation with a false sense of security**, believing that this technology is inherently safe.

此外，“容器”这个词经常被误解，因为许多开发人员倾向于**将隔离的概念与虚假的安全感**联系起来，认为这项技术本质上是安全的。

The key here is that containers **don’t have any security dimension by default**. Their security completely depends on:

这里的关键是容器**默认没有任何安全维度**。它们的安全性完全取决于：

- the supporting infrastructure (OS and platform)
- their embedded software components
- their runtime configuration

- 支持基础设施（操作系统和平台）
- 他们的嵌入式软件组件
- 他们的运行时配置

Container security represents a broad topic, but the good news is that many best practices are low-hanging fruits one can harvest to **quickly reduce the attack surface** of their deployments.

容器安全是一个广泛的话题，但好消息是，许多最佳实践都是唾手可得的成果，可以**快速减少部署的攻击面**。

> That's why we curated a set of the best recommendations regarding Docker containers configuration at build and runtime. Check out the one-page cheat sheet below.

> 这就是为什么我们策划了一组关于构建和运行时 Docker 容器配置的最佳建议。查看下面的一页备忘单。

[![](https://res.cloudinary.com/da8kiytlc/image/upload/c_scale,w_500/v1627655008/Cheatsheets/Docker-Security-Cheatsheet_hp8lh3.png)](https://res.cloudinary.com/da8kiytlc/image/upload/v1627655008/Cheatsheets/Docker-Security-Cheatsheet_hp8lh3.pdf) [Download the Docker security cheatsheet](https://res.cloudinary.com/da8kiytlc/image/upload/v1627655008/Cheatsheets/Docker-Security-Cheatsheet_hp8lh3.pdf)

/image/upload/v1627655008/Cheatsheets/Docker-Security-Cheatsheet_hp8lh3.pdf) [下载 Docker 安全备忘单](https://res.cloudinary.com/da8kiytlc/image/upload/v1627655008/Cheatsheets/Docker-HPSecurity-Cheatsheet_.pdf)

_Note: in a managed environment like Kubernetes, most of these settings can be overridden by a Security Context or other higher-level security rules._ [_See more_](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/)

_注意：在像 Kubernetes 这样的托管环境中，这些设置中的大部分都可以被安全上下文或其他更高级别的安全规则覆盖。_ [_查看更多_](https://kubernetes.io/docs/tasks/configure-pod-容器/安全上下文/)

## Build Configuration

## 构建配置

### Check your images

### 检查您的图像

Carefully choose your base image when you `docker pull image:tag`

docker pull image:tag 时请仔细选择基础镜像

You should always prefer using a **trusted image**, preferably from the [Docker Official Images](https://docs.docker.com/docker-hub/official_repos/), in order to mitigate supply chain attacks.

您应该始终更喜欢使用 **trusted image**，最好来自 [Docker Official Images](https://docs.docker.com/docker-hub/official_repos/)，以减轻供应链攻击。

If you need to choose a base distro, [Alpine Linux is recommended](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) since it is one of the lightest available, ensuring the attack surface is reduced.

如果您需要选择基础发行版，[推荐使用 Alpine Linux](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) 因为它是最轻量级的发行版之一，可确保攻击面降低了。

> Do I need to use the latest or a fixed tag version?

> 我需要使用最新的还是固定的标签版本？

First, you should understand that Docker tagging works from less to more specific, that’s the reason why :

首先，您应该了解 Docker 标记从较少到更加具体的工作，这就是原因：

```
python:3.9.6-alpine3.14

python:3.9.6-alpine

python:3.9-alpine

python:alpine

```

all refer to the same image (at the moment of writing) for example.

例如，所有都指的是相同的图像（在撰写本文时）。

By being very specific and pinning down a version, you are shielding yourself from any future breaking change. On the other hand, using the latest version ensures that more vulnerabilities are patched. **This is a tradeoff**, but pinning to a stable release is what is generally recommended.

通过非常具体并确定一个版本，您可以保护自己免受未来任何重大更改的影响。另一方面，使用最新版本可确保修补更多漏洞。 **这是一种权衡**，但通常建议固定到稳定版本。

Considering that, we would pick `python:3.9-alpine` here.

考虑到这一点，我们会在这里选择 `python:3.9-alpine`。

> Note: the same applies to packages installed during the build process of your image.

> 注意：这同样适用于在你的镜像构建过程中安装的包。

### Always use an unprivileged user

### 始终使用非特权用户

By default, the process inside a container is **run as root** (id=0).

默认情况下，容器内的进程**以 root 身份运行** (id=0)。

To enforce the principle of least privilege, you should set a default user. For this you have two options:

为了执行最小权限原则，您应该设置一个默认用户。为此，您有两个选择：

- Either specify an arbitrary user ID that won’t exist in the running container, with the`-u` option:

- 使用`-u`选项指定一个在运行容器中不存在的任意用户ID：

`docker run -u 4000 <image>`

`docker run -u 4000 <图像>`

> Note: if you later need to mount a filesystem, you should match the user ID you are using to the host user in order to access the files.

> 注意：如果您以后需要挂载文件系统，您应该将您使用的用户 ID 与主机用户匹配，以便访问文件。

- Or anticipate by creating a default user in your Dockerfile:

- 或者通过在 Dockerfile 中创建默认用户来预测：

```
FROM <base image>

RUN addgroup -S appgroup \
&& adduser -S appuser -G appgroup

USER appuser

... <rest of Dockerfile> ...

```

> Note: you would need to check what tool is used to create users and groups in your base image.

> 注意：您需要检查使用什么工具在您的基础映像中创建用户和组。

### Use a separate User ID namespace

### 使用单独的用户 ID 命名空间

By default, the Docker daemon uses the host’s user ID namespace. Consequently, any success in privilege escalation inside a container would also mean root access both to the host and to other containers.

默认情况下，Docker 守护进程使用主机的用户 ID 命名空间。因此，容器内权限提升的任何成功也意味着对主机和其他容器的 root 访问。

To mitigate this risk, you should configure your host and the Docker daemon to use a separate namespace with the `--userns-remap` option. [See more](https://docs.docker.com/engine/security/userns-remap/#prerequisites)

为了降低这种风险，您应该将主机和 Docker 守护程序配置为使用带有 `--userns-remap` 选项的单独命名空间。 [查看更多](https://docs.docker.com/engine/security/userns-remap/#prerequisites)

### Handle environment variables with care

### 小心处理环境变量

You should **never include sensitive information** in plaintext in an ENV directive: they are simply not a safe place to store any bit of information you don’t want to be present in the final layer. For example, if you thought that unsetting an environment variable like this:

**永远不要在 ENV 指令中以明文形式包含敏感信息**：它们根本不是一个安全的地方来存储任何你不想出现在最后一层的信息。例如，如果您认为像这样取消设置环境变量：

```
ENV $VAR
RUN unset $VAR

```

Was safe, you are wrong! `$VAR` will still be present in the containers and could be dumped anytime!

是安全的，你错了！ `$VAR` 仍将存在于容器中，并且可以随时倾倒！

To prevent runtime read access, use a single RUN command to set and unset the variable in a single layer (don't forget the variable **can still be extracted** from the image).

为防止运行时读取访问，请使用单个 RUN 命令在单个层中设置和取消设置变量（不要忘记变量 ** 仍然可以从图像中提取**）。

```
RUN export ADMIN_USER="admin" \
    && ... \
    && unset ADMIN_USER

```

More idiomatically, **use the ARG directive** (ARG values are not available after the image is built).

更惯用的是，**使用 ARG 指令**（构建图像后 ARG 值不可用）。

Unfortunately, **secrets are too often hardcoded into docker images’ layers**, that’s the reason we developed a [scanning tool](https://github.com/GitGuardian/ggshield) leveraging GitGuardian secrets engine to find them:

不幸的是，**秘密经常被硬编码到 docker 图像的层中**，这就是我们开发 [扫描工具](https://github.com/GitGuardian/ggshield) 利用 GitGuardian 秘密引擎来查找它们的原因：

`ggshield scan docker <image>`

`ggshield 扫描泊坞窗 <图像>`

More on scanning images for vulnerabilities later.

稍后将详细介绍扫描图像中的漏洞。

### Don’t expose the Docker daemon socket

### 不要暴露 Docker 守护进程套接字

Unless you are very confident with what you are doing, never expose the UNIX socket that Docker is listening to: `/var/run/docker.sock`

除非你对自己正在做的事情非常有信心，否则永远不要暴露 Docker 正在侦听的 UNIX 套接字：`/var/run/docker.sock`

This is the primary entry point for the Docker API. Giving someone access to it is equivalent to giving unrestricted root access to your host.

这是 Docker API 的主要入口点。授予某人访问权限等同于授予对您的主机的无限制 root 访问权限。

You should never expose it to other containers:

你永远不应该将它暴露给其他容器：

`-v /var/run/docker.sock://var/run/docker.sock`

`-v /var/run/docker.sock://var/run/docker.sock`

## Privileges, capabilities and shared resources

## 权限、能力和共享资源

First, your container should **never be running as privileged**, otherwise, it would be allowed to have all the root capabilities on the host machine.

首先，您的容器应该**永远不要以特权方式运行**，否则，它将被允许在主机上拥有所有 root 权限。

To be even safer, it is recommended to explicitly forbid the possibility to add new privileges after a container has been created with the option `--security-opt=no-new-privileges`.

为了更安全，建议使用选项`--security-opt=no-new-privileges`明确禁止在创建容器后添加新权限的可能性。

Second, **capabilities** are a Linux mechanism used by Docker to turn the binary “root/non-root” dichotomy into a fine-grained access control system: your containers are run with a default set of enabled capabilities, which you most probably don't all need.

其次，**capabilities** 是 Docker 使用的一种 Linux 机制，用于将二进制“root/non-root”二分法转变为细粒度的访问控制系统：您的容器以一组默认启用的功能运行，您最可能不需要。

It's recommended to **drop all default capabilities** and only add them individually: see the list of default capabilities

建议**删除所有默认功能**，只单独添加它们：查看默认功能列表

for instance, a web server would probably only need the NET\_BIND\_SERVICE to bind to a port under 1024 (like port 80).

例如，Web 服务器可能只需要 NET\_BIND\_SERVICE 来绑定到 1024 下的端口（如端口 80）。

Third, **don’t share the sensitive parts** of the host filesystem :

第三，**不要共享主机文件系统的敏感部分**：

- root (/),
- device (/dev)
- process (/proc)
- virtual (/sys) mount points.

- 根 （/），
- 设备 (/dev)
- 进程 (/proc)
- 虚拟 (/sys) 挂载点。

If you need access to host devices, be careful to selectively enable the access options with the`[r|w|m]` flags (read, write, and use mknod).

如果您需要访问主机设备，请小心使用`[r|w|m]` 标志（读、写和使用 mknod）有选择地启用访问选项。

### Use Control Groups to limit access to resources

### 使用控制组来限制对资源的访问

Control Groups are the mechanism used to control access to CPU, memory, disk I/O for each container.

控制组是用于控制每个容器对 CPU、内存、磁盘 I/O 的访问的机制。

By default, a container is associated with a dedicated `cgroup`, but if the option `--cgroup-parent` is present, you are putting the host resources **at risk of a DoS attack**, because you are allowing shared resources between the host and the container.

默认情况下，一个容器与一个专用的 `cgroup` 相关联，但是如果选项 `--cgroup-parent` 存在，则您将主机资源**置于 DoS 攻击的风险中**，因为您允许共享主机和容器之间的资源。

In the same idea, it is recommended to specify memory and CPU usage by using options like

出于同样的想法，建议使用以下选项指定内存和 CPU 使用率

```
--memory=”400m”
--memory-swap=”1g”

--cpus=0.5
--restart=on-failure:5
--ulimit nofile=5
--ulimit nproc=5

```

[See more on resources constraints](https://docs.docker.com/config/containers/resource_constraints/)

[查看更多资源限制](https://docs.docker.com/config/containers/resource_constraints/)

## Filesystem

##  文件系统

### Only allow read access to the root filesystem

### 只允许读访问根文件系统

Containers should be ephemeral and thus mostly stateless. That’s why you can often limit the mounted filesystem to be read-only.

容器应该是短暂的，因此大多是无状态的。这就是为什么您通常可以将挂载的文件系统限制为只读的原因。

`docker run --read-only <image>`

`docker run --read-only <image>`

### Use a temporary filesystem for non-persistent data

### 为非持久性数据使用临时文件系统

If you need only temporary storage, use the appropriate option

如果您只需要临时存储，请使用适当的选项

`docker run --read-only --tmpfs /tmp:rw ,noexec,nosuid <image>`

`docker run --read-only --tmpfs /tmp:rw ,noexec,nosuid <image>`

### Use a filesystem for persistent data 

### 使用文件系统保存持久数据

If you need to share data with the host filesystem or other containers, you have two options:

如果您需要与主机文件系统或其他容器共享数据，您有两种选择：

- Create a bind mount with a limited useable disk space ( `--mount type=bind,o=size`)
- Create a bind volume for a dedicated partition ( `--mount type=volume`)

- 创建一个可用磁盘空间有限的绑定挂载（`--mount type=bind,o=size`）
- 为专用分区创建一个绑定卷（`--mount type=volume`）

In either case, if the shared data doesn’t need to be modified by the container, use the read-only option.

无论哪种情况，如果容器不需要修改共享数据，请使用只读选项。

`docker run -v <volume-name>:/path/in/container:ro <image>`

`docker run -v <volume-name>:/path/in/container:ro <image>`

or

或者

`docker run --mount source=<volume-name>,destination=/path/in/container,readonly <image>`

`docker run --mount source=<volume-name>,destination=/path/in/container,readonly <image>`

## Networking

##  联网

### Don’t use Docker’s default bridge docker0

### 不要使用Docker的默认网桥docker0

`docker0` is a network bridge that is created on start to separate the host network from the container network.

`docker0` 是一个在启动时创建的网桥，用于将主机网络与容器网络分开。

When a container is created, Docker connects it to the `docker0` network by default. Therefore, all containers are connected to `docker0` and are able to communicate with each other.

创建容器时，Docker 默认将其连接到 `docker0` 网络。因此，所有容器都连接到 docker0 并且能够相互通信。

You should disable this default connection of all the containers by specifying the option `--bridge=none` and instead, **create a dedicated network for every connection** with the command:

您应该通过指定选项`--bridge=none` 来禁用所有容器的默认连接，而是使用以下命令**为每个连接创建一个专用网络**：

`docker network create <network_name>`

`docker 网络创建 <network_name>`

And then use it to access the host network interface

然后用它来访问主机网络接口

`docker run --network=<network_name>`

`docker run --network=<network_name>`

![Docker networking simple example](https://blog.gitguardian.com/content/images/2021/07/Docker-networking.png)Docker networking simple example

For example, to create a web server talking to a database (started in another container), the best practice would be to create a bridge network `WEB` in order to route incoming traffic from the host network interface and use another bridge `DB` only used to connect the database and the web containers.

例如，要创建一个与数据库通信的 Web 服务器（在另一个容器中启动），最佳实践是创建一个桥接网络“WEB”，以便路由来自主机网络接口的传入流量并使用另一个桥接“DB”仅用于连接数据库和 Web 容器。

### Don’t share the host’s network namespace

### 不要共享主机的网络命名空间

Same idea, isolate the host's network interface: the `--network` host option should not be used.

同样的想法，隔离主机的网络接口：不应使用 `--network` 主机选项。

## Logging

## 记录

The default logging level is INFO, but you can specify another one with the option:

默认日志级别为 INFO，但您可以使用以下选项指定另一个级别：

`--log-level="debug"|"info"|"warn"|"error"|"fatal"`

`--log-level="debug"|"info"|"warn"|"error"|"fatal"`

What is less known is the log export capacity of Docker: if your containerized app produces event logs, you can redirect `STDERR` and `STDOUT` streams to an external logging service for decoupling using the option `--log-driver=<logging_driver >`

Docker 的日志导出能力鲜为人知：如果您的容器化应用程序生成事件日志，您可以将 `STDERR` 和 `STDOUT` 流重定向到外部日志服务以使用选项 `--log-driver=<logging_driver 进行解耦>`

You can also enable dual logging to preserve docker access to logs while using an external service. If your app uses dedicated files (often written under `/var/log`), you can still redirect these streams: [see the official documentation](https://docs.docker.com/config/containers/logging/configure/)

您还可以启用双日志记录以在使用外部服务时保留 docker 对日志的访问。如果您的应用程序使用专用文件（通常写在`/var/log` 下），您仍然可以重定向这些流：[参见官方文档](https://docs.docker.com/config/containers/logging/configure/)

## Scan for vulnerabilities & secrets

## 扫描漏洞和秘密

Last but not least, I hope it is now clear that your containers are only going to be as safe as the software they are running. To make sure your images are vulnerability-free, you need to perform a scan for known vulnerabilities.

最后但并非最不重要的一点是，我希望现在您的容器将与它们运行的软件一样安全。为确保您的图像没有漏洞，您需要对已知漏洞执行扫描。

Many tools are available for different use-case and in different forms:

许多工具可用于不同的用例和不同的形式：

Scanning for vulnerabilities:

漏洞扫描：

- Free options:
   - [Clair](https://github.com/quay/clair)
   - [Trivy](https://github.com/aquasecurity/trivy)
   - [Docker Bench for Security](https://github.com/docker/docker-bench-security)
- Commercial:
   - [Snyk](https://github.com/snyk/snyk) (open source and free option available)
   - [Anchore](https://github.com/anchore/anchore-engine) (open source and free option available)
   - [JFrog XRay](https://jfrog.com/fr/xray/)
   - [Qualys](https://qualysguard.qg2.apps.qualys.com/cs/help/vuln_scans/docker_images.htm)

- 免费选项：
  - [克莱尔](https://github.com/quay/clair)
  - [琐事](https://github.com/aquasecurity/trivy)
  - [Docker Bench for Security](https://github.com/docker/docker-bench-security)
- 商业的：
  - [Snyk](https://github.com/snyk/snyk)（提供开源和免费选项)
  - [锚点](https://github.com/anchore/anchore-engine)（提供开源和免费选项)
  - [JFrog XRay](https://jfrog.com/fr/xray/)
  - [Qualys](https://qualysguard.qg2.apps.qualys.com/cs/help/vuln_scans/docker_images.htm)

Scanning for secrets:

秘密扫描：

- [ggshield](https://github.com/GitGuardian/ggshield) (open source and free option available)
- [SecretScanner](https://github.com/deepfence/SecretScanner)(free)

- [ggshield](https://github.com/GitGuardian/ggshield)（提供开源和免费选项)
- [SecretScanner](https://github.com/deepfence/SecretScanner)（免费)

If you are interested in other cheat sheets about security we have put together these for you:

如果您对其他有关安全性的备忘单感兴趣，我们为您整理了这些：

- Best practices for[managing and storing secrets including API keys and other credentials](https://blog.gitguardian.com/secrets-api-management/)
- [Rewriting your git history, removing files permanently](https://blog.gitguardian.com/rewriting-git-history-cheatsheet/)

- [管理和存储机密，包括 API 密钥和其他凭据]的最佳实践（https://blog.gitguardian.com/secrets-api-management/）
- [重写你的 git 历史，永久删除文件](https://blog.gitguardian.com/rewriting-git-history-cheatsheet/)

Explore related articles by category [Cheat sheets](http://blog.gitguardian.com/tag/cheat-sheets/ "Cheat sheets")

按类别浏览相关文章 [备忘单](http://blog.gitguardian.com/tag/cheat-sheets/“备忘单”)

### More in [Cheat sheets](http://blog.gitguardian.com/tag/cheat-sheets/)

### 更多在 [备忘单](http://blog.gitguardian.com/tag/cheat-sheets/)

- #### [Rewriting your git history, removing files permanently - [cheat sheet included]](http://blog.gitguardian.com/rewriting-git-history-cheatsheet/) 

- #### [重写你的 git 历史，永久删除文件 - [包含备忘单]](http://blog.gitguardian.com/rewriting-git-history-cheatsheet/)

