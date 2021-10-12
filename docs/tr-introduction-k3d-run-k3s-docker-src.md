# Introduction to k3d: Run K3s in Docker

# k3d 介绍：在 Docker 中运行 K3s

March 1, 2021
\|
By:
[Thorsten Klein](https://www.suse.com/c/author/thorsten-klein-95googlemail-com/ "View all posts by Thorsten Klein")

2021 年 3 月 1 日
\|
经过：
[Thorsten Klein](https://www.suse.com/c/author/thorsten-klein-95googlemail-com/“查看 Thorsten Klein 的所有帖子”)

In this blog post, we’re going to talk about k3d, a tool that allows you to run throwaway Kubernetes clusters anywhere you have Docker installed. I’ve anticipated your questions — so let’s go!

在这篇博文中，我们将讨论 k3d，这是一种工具，可让您在安装了 Docker 的任何地方运行一次性 Kubernetes 集群。我已经预料到你的问题了——所以我们走吧！

## What is k3d?

## 什么是k3d？

k3d is a small program made for running a [K3s](https://k3s.io) cluster in Docker. K3s is a lightweight, CNCF-certified Kubernetes distribution and Sandbox project. Designed for low-resource environments, K3s is distributed as a single binary that uses under 512MB of RAM. To learn more about K3s, head over to [the documentation](https://rancher.com/docs/k3s/latest/en/) or check out this [blog post](https://rancher.com/blog/2019/2019-02-26-introducing-k3s-the-lightweight-kubernetes-distribution-built-for-the-edge/) or [video.](https://www.youtube.com/watch?v=hMr3prm9gDM)

k3d 是一个小程序，用于在 Docker 中运行 [K3s](https://k3s.io) 集群。 K3s 是一个轻量级的、经过 CNCF 认证的 Kubernetes 发行版和沙盒项目。 K3s 专为低资源环境设计，作为单个二进制文件分发，使用的 RAM 低于 512MB。要了解有关 K3s 的更多信息，请转到 [文档](https://rancher.com/docs/k3s/latest/en/) 或查看此 [博客文章](https://rancher.com/blog/2019/2019-02-26-introducing-k3s-the-lightweight-kubernetes-distribution-built-for-the-edge/)或[视频。](https://www.youtube.com/watch?v=hMr3prm9gDM)

k3d uses a [Docker image](https://hub.docker.com/r/rancher/k3s/tags) built from the [K3s repository](https://github.com/rancher/k3s) to spin up multiple K3s nodes in Docker containers on any machine with Docker installed. That way, a single physical (or virtual) machine (let’s call it Docker Host) can run multiple K3s clusters, with multiple server and agent nodes each, simultaneously.

k3d 使用从 [K3s 存储库](https://github.com/rancher/k3s) 构建的 [Docker 映像](https://hub.docker.com/r/rancher/k3s/tags) 来启动多个任何安装了 Docker 的机器上的 Docker 容器中的 K3s 节点。这样，单个物理（或虚拟）机器（我们称之为 Docker 主机)可以同时运行多个 K3s 集群，每个集群具有多个服务器和代理节点。

## What Can k3d Do?

## k3d 能做什么？

As of k3d version v4.0.0, released in January 2021, k3d’s abilities boil down to the following features:

截至 2021 年 1 月发布的 k3d v4.0.0 版本，k3d 的能力归结为以下特性：

- create/stop/start/delete/grow/shrink K3s clusters (and individual nodes)
   - via command line flags
   - via configuration file
- manage and interact with container registries that can be used with the cluster
- manage Kubeconfigs for the clusters
- import images from your local Docker daemon into the container runtime running in the cluster

- 创建/停止/启动/删除/增长/缩小 K3s 集群（和单个节点）
  - 通过命令行标志
  - 通过配置文件
- 管理可与集群一起使用的容器注册表并与之交互
- 管理集群的 Kubeconfigs
- 将本地 Docker 守护进程中的图像导入集群中运行的容器运行时

Obviously, there’s way more to it and you can tweak everything in great detail.

显然，它还有更多方法，您可以非常详细地调整所有内容。

## What is k3d Used for?

## k3d 用于什么？

The main use case for k3d is local development on Kubernetes with little hassle and resource usage. The intention behind the initial development of k3d was to provide developers with an easy tool that allowed them to run a lightweight Kubernetes cluster on their development machine, giving them fast iteration times in a production-like environment (as opposed to running docker-compose locally vs. Kubernetes in production).

k3d 的主要用例是在 Kubernetes 上进行本地开发，几乎没有麻烦和资源使用。 k3d 最初开发背后的意图是为开发人员提供一个简单的工具，使他们能够在他们的开发机器上运行轻量级 Kubernetes 集群，让他们在类似生产的环境中快速迭代（而不是在本地运行 docker-compose与生产中的 Kubernetes 相比）。

Over time, k3d also evolved into a tool used by operations to test some Kubernetes (or, specifically K3s) features in an isolated environment. For example, with k3d you can easily create multi-node clusters, deploy something on top of it, simply stop a node and see how Kubernetes reacts and possibly reschedules your app to other nodes.

随着时间的推移，k3d 也演变成运维人员用来在隔离环境中测试某些 Kubernetes（或者，特别是 K3s）功能的工具。例如，使用 k3d，您可以轻松创建多节点集群，在其上部署一些东西，只需停止一个节点并查看 Kubernetes 的反应，并可能将您的应用程序重新安排到其他节点。

Additionally, you can use k3d in your continuous integration system to quickly spin up a cluster, deploy your test stack on top of it and run integration tests. Once you’re finished, you can simply decommission the cluster as a whole. No need to worry about proper cleanups and possible leftovers.

此外，您可以在持续集成系统中使用 k3d 来快速启动集群，在其上部署测试堆栈并运行集成测试。完成后，您可以简单地将集群作为一个整体退役。无需担心适当的清理和可能的剩菜。

We also provide a `k3d-dind` image (similar to dreams within dreams in the movie Inception, we've got containers within containers within containers.) With that, you can create a docker-in-docker environment where you run k3d, which spawns a K3s cluster in Docker. That means that you only have a single container (k3d-dind) running on your Docker host, which in turn runs a whole K3s/Kubernetes cluster inside.

我们还提供了一个 `k3d-dind` 映像（类似于电影《盗梦空间》中的梦中梦，我们在容器中的容器中有容器。）这样，您就可以创建一个运行 k3d 的 docker-in-docker 环境，它在 Docker 中生成了一个 K3s 集群。这意味着您只有一个容器 (k3d-dind) 在 Docker 主机上运行，而后者又在其中运行整个 K3s/Kubernetes 集群。

## How Do I Use k3d?

## 我如何使用 k3d？

1. [Install k3d](https://k3d.io/#installation) (and[kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), if you want to use it )

1. [安装 k3d](https://k3d.io/#installation)（和[kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)，如果你想使用它)

**Note**: to follow along with this post, use at least k3d v4.1.1
2. Try one of the following examples or use the[documentation](https://k3d.io/usage/commands/) or the CLI help text to find your own way ( `k3d [command] --help`)

**注意**：要跟随这篇文章，至少使用 k3d v4.1.1
2. 尝试以下示例之一或使用[文档](https://k3d.io/usage/commands/) 或 CLI 帮助文本找到自己的方法 (`k3d [command] --help`)

### The “Simple” Way

###“简单”的方式

```
k3d cluster create
```

This single command spawns a K3s cluster with two containers: A Kubernetes control-plane node (server) and a load balancer (serverlb) in front of it. It puts both of them in a dedicated Docker network and exposes the Kubernetes API on a randomly chosen free port on the Docker host. It also creates a named Docker volume in the background as a preparation for image imports.

这个单一的命令产生了一个 K3s 集群，它有两个容器：一个 Kubernetes 控制平面节点（服务器）和一个负载均衡器（serverlb）在它前面。它将它们都放在一个专用的 Docker 网络中，并在 Docker 主机上随机选择的空闲端口上公开 Kubernetes API。它还在后台创建一个命名的 Docker 卷，作为映像导入的准备。

By default, if you don’t provide a name argument, the cluster will be named `k3s-default`

默认情况下，如果您不提供 name 参数，则集群将命名为 `k3s-default`

and the containers will show up as `k3d-<-role>-<#>`, so in this case `k3d-k3s-default-serverlb` and `k3d-k3s-default-server-0`.

并且容器将显示为 `k3d-<-role>-<#>`，所以在这种情况下，`k3d-k3s-default-serverlb` 和 `k3d-k3s-default-server-0`。

k3d waits until everything is ready, pulls the Kubeconfig from the cluster and merges it with your default Kubeconfig (usually it’s in `$HOME/.kube/config` or whatever path your `KUBECONFIG` environment variable points to).

k3d 等待一切准备就绪，从集群中拉取 Kubeconfig 并将其与您的默认 Kubeconfig 合并（通常它在 `$HOME/.kube/config` 或您的 `KUBECONFIG` 环境变量指向的任何路径中）。

No worries, you can tweak that behavior as well.

不用担心，您也可以调整该行为。

Check out what you’ve just created using `kubectl` to show you the nodes: `kubectl get nodes`.

查看您刚刚使用 `kubectl` 创建的内容以显示节点：`kubectl get nodes`。

k3d also gives you some commands to list your creations: `k3d cluster|node|registry list`.

k3d 还提供了一些命令来列出您的作品：`k3d cluster|node|registry list`。

### The “Simple but Sophisticated” Way

###“简单但复杂”的方式

`k3d cluster create mycluster --api-port 127.0.0.1:6445 --servers 3 --agents 2 --volume '/home/me/mycode:/code@agent[*]' --port '8080:80@ loadbalancer'`

`k3d 集群创建 mycluster --api-port 127.0.0.1:6445 --servers 3 --agents 2 --volume '/home/me/mycode:/code@agent[*]' --port '8080:80@负载均衡器'`

`` This single command spawns a K3s cluster with six containers:

`` 这个单一的命令产生了一个包含六个容器的 K3s 集群：

-  1 load balancer
-  3 servers (control-plane nodes)
-  2 agents (formerly worker nodes)

- 1 个负载平衡器
- 3 个服务器（控制平面节点）
- 2 个代理（以前的工作节点）

With the `--api-port 127.0.0.1:6445`, you tell k3d to map the Kubernetes API Port ( `6443` internally) to `127.0.0.1`/localhost’s port `6445`. That means that you will have this connection string in your Kubeconfig: `server: https://127.0.0.1:6445` to connect to this cluster.

使用`--api-port 127.0.0.1:6445`，你告诉k3d将Kubernetes API端口（内部`6443`）映射到`127.0.0.1`/localhost的端口`6445`。这意味着您将在 Kubeconfig 中有此连接字符串：`server: https://127.0.0.1:6445` 以连接到此集群。

This port will be mapped from the load balancer to your host system. From there, requests will be proxied to your server nodes, effectively simulating a production setup, where server nodes also can go down and you would want to failover to another server.

此端口将从负载平衡器映射到您的主机系统。从那里，请求将被代理到您的服务器节点，有效地模拟生产设置，其中服务器节点也可能出现故障并且您希望故障转移到另一台服务器。

The `--volume /home/me/mycode:/code@agent[*]` bind mounts your local directory `/home/me/mycode` to the path `/code` inside all ( `[*]` of your agent nodes). Replace `*` with an index (here: 0 or 1) to only mount it into one of them.

`--volume /home/me/mycode:/code@agent[*]` 绑定将你的本地目录 `/home/me/mycode` 挂载到路径 `/code` 里面（你的代理节点）。将 `*` 替换为索引（此处：0 或 1）以仅将其装入其中之一。

The specification telling k3d which nodes it should mount the volume to is called “node filter” and it’s also used for other flags, like the `--port` flag for port mappings.

告诉 k3d 应该将卷挂载到哪些节点的规范称为“节点过滤器”，它也用于其他标志，例如用于端口映射的 `--port` 标志。

That said, `--port '8080:80@loadbalancer'` maps your local host’s port `8080` to port `80` on the load balancer (serverlb), which can be used to forward HTTP ingress traffic to your cluster. For example, you can now deploy a web app into the cluster (Deployment), which is exposed (Service) externally via an Ingress such as `myapp.k3d.localhost`.

也就是说，`--port '8080:80@loadbalancer'` 将本地主机的端口 8080 映射到负载均衡器 (serverlb) 上的端口 80，可用于将 HTTP 入口流量转发到您的集群。例如，您现在可以将 Web 应用程序部署到集群（部署）中，该应用程序通过诸如 `myapp.k3d.localhost` 之类的 Ingress 对外公开（服务）。

Then (provided that everything is set up to resolve that domain to your local host IP), you can point your browser to `http://myapp.k3d.localhost:8080` to access your app. Traffic then flows from your host through the Docker bridge interface to the load balancer. From there, it’s proxied to the cluster, where it passes via Ingress and Service to your application Pod.

然后（假设一切都设置为将该域解析为您的本地主机 IP），您可以将浏览器指向 `http://myapp.k3d.localhost:8080` 以访问您的应用程序。然后流量从您的主机通过 Docker 网桥接口流向负载均衡器。从那里，它被代理到集群，在那里它通过 Ingress 和 Service 传递到您的应用程序 Pod。

**Note**: You have to have some mechanism set up to route to resolve `myapp.k3d.localhost` to your local host IP ( `127.0.0.1`). The most common way is using entries of the form `127.0.0.1 myapp.k3d.localhost` in your `/etc/hosts` file ( `C:\Windows\System32\drivers\etc\hosts` on Windows). However, this does not allow for wildcard entries ( `*.localhost`), so it may become a bit cumbersome after a while, so you may want to have a look at tools like [`dnsmasq` (MacOS/UNIX](https://en.wikipedia.org/wiki/Dnsmasq)) or [`Acrylic` (Windows)](https://stackoverflow.com/a/9695861/6450189) to ease the burden. 

**注意**：你必须设置一些机制来将 `myapp.k3d.localhost` 解析到你的本地主机 IP (`127.0.0.1`)。最常见的方法是在你的 `/etc/hosts` 文件（Windows 上的 `C:\Windows\System32\drivers\etc\hosts`）中使用形式为 `127.0.0.1 myapp.k3d.localhost` 的条目。但是，这不允许通配符条目（`*.localhost`），所以一段时间后它可能会变得有点麻烦，所以你可能想看看像 [`dnsmasq` (MacOS/UNIX](https://en.wikipedia.org/wiki/Dnsmasq)) 或 [`Acrylic` (Windows)](https://stackoverflow.com/a/9695861/6450189) 以减轻负担。

**Tip**: You can install the package `<a href="https://manpages.debian.org/unstable/libnss-myhostname/nss-myhostname.8.en.html" target="_blank" rel= "noopener" data-index="40">libnss-myhostname</a>` on some systems (at least Linux operating systems including SUSE Linux and openSUSE), to auto-resolve `*.localhost` domains to `127.0.0.1 `, which means you don't have to fiddle around with eg `/etc/hosts`, if you prefer to test via Ingress, where you need to set a domain.

**提示**：您可以安装包`<a href="https://manpages.debian.org/unstable/libnss-myhostname/nss-myhostname.8.en.html" target="_blank" rel= "noopener" data-index="40">libnss-myhostname</a>` 在某些系统上（至少是 Linux 操作系统，包括 SUSE Linux 和 openSUSE），将 `*.localhost` 域自动解析为 `127.0.0.1 `，这意味着您不必摆弄例如`/etc/hosts`，如果您更喜欢通过 Ingress 进行测试，则需要在其中设置域。

One interesting thing to note here: if you create more than one server node, K3s will be given the `--cluster-init` flag, which means that it swaps its internal datastore (by default that’s SQLite) for etcd.

这里需要注意一件有趣的事情：如果你创建了多个服务器节点，K3s 将被赋予 `--cluster-init` 标志，这意味着它会将其内部数据存储（默认为 SQLite）交换为 etcd。

### The “Configuration as Code” Way

###“配置即代码”方式

As of k3d v4.0.0 (January 2021), we support config files to configure everything as code that you’d previously do via command line flags (and soon possibly even more than that).

从 k3d v4.0.0（2021 年 1 月）开始，我们支持配置文件将所有内容配置为您以前通过命令行标志执行的代码（很快甚至可能更多）。

As of this writing, the JSON-Schema used to validate the configuration file can be [found in the repository](https://github.com/rancher/k3d/blob/092f26a4e27eaf9d3a5bc32b249f897f448bc1ce/pkg/config/v1alpha2/schema.json) .

在撰写本文时，用于验证配置文件的 JSON-Schema 可以[在存储库中找到](https://github.com/rancher/k3d/blob/092f26a4e27eaf9d3a5bc32b249f897f448bc1ce/pkg/config/v1alpha2/schema.json) .

Here’s an example config file:

这是一个示例配置文件：

```
# k3d configuration file, saved as e.g./home/me/myk3dcluster.yaml
apiVersion: k3d.io/v1alpha2 # this will change in the future as we make everything more stable
kind: Simple # internally, we also have a Cluster config, which is not yet available externally
name: mycluster # name that you want to give to your cluster (will still be prefixed with `k3d-`)
servers: 1 # same as `--servers 1`
agents: 2 # same as `--agents 2`
kubeAPI: # same as `--api-port 127.0.0.1:6445`
hostIP: "127.0.0.1"
hostPort: "6445"
ports:
 - port: 8080:80 # same as `--port 8080:80@loadbalancer
nodeFilters:
 - loadbalancer
options:
k3d: # k3d runtime settings
wait: true # wait for cluster to be usable before returining;same as `--wait` (default: true)
timeout: "60s" # wait timeout before aborting;same as `--timeout 60s`
k3s: # options passed on to K3s itself
extraServerArgs: # additional arguments passed to the `k3s server` command
     - --tls-san=my.host.domain
extraAgentArgs: [] # addditional arguments passed to the `k3s agent` command
kubeconfig:
updateDefaultKubeconfig: true # add new cluster to your default Kubeconfig;same as `--kubeconfig-update-default` (default: true)
switchCurrentContext: true # also set current-context to the new cluster's context;same as `--kubeconfig-switch-context` (default: true)
```

Assuming that we saved this as `/home/me/myk3dcluster.yaml`, we can use it to configure a new cluster:

假设我们将其保存为 `/home/me/myk3dcluster.yaml`，我们可以使用它来配置一个新集群：

`k3d cluster create --config /home/me/myk3dcluster.yaml`

`k3d 集群创建 --config /home/me/myk3dcluster.yaml`

Note that you can still set additional arguments or flags, which will then take precedence (or will be merged) with whatever you have defined in the config file.

请注意，您仍然可以设置其他参数或标志，这些参数或标志将优先（或将合并）您在配置文件中定义的任何内容。

## What More Can I Do with k3d?

## 我还能用 k3d 做什么？

You can use k3d in even more ways, including:

您可以通过更多方式使用 k3d，包括：

- Create a cluster together with a k3d-managed container**registry**
- Use the cluster for fast development with**hot code reloading**
- Use k3d in combination with other development tools like`Tilt` or `Skaffold` ``
   - both can leverage the power of importing images via`k3d image import`
   - both can alternatively make use of a k3d-managed registry to speed up your development loop
- Use k3d in your CI system (we have a[PoC for that](https://github.com/iwilltry42/k3d-demo/blob/main/.drone.yml))
- Integrate it in your vscode workflow using the awesome new[community-maintained vscode extension](https://github.com/inercia/vscode-k3d)
- Use it to[set up K3s high availability](https://rancher.com/blog/2020/set-up-k3s-high-availability-using-k3d)

- 与 k3d 管理的容器一起创建集群**注册表**
- 使用集群快速开发**热代码重载**
- 将 k3d 与其他开发工具如`Tilt` 或`Skaffold` 结合使用``
  - 两者都可以利用通过`k3d image import`导入图像的能力
  - 两者都可以选择使用 k3d 管理的注册表来加速您的开发循环
- 在你的 CI 系统中使用 k3d（我们有一个[PoC](https://github.com/iwilltry42/k3d-demo/blob/main/.drone.yml))
- 使用令人敬畏的新 [社区维护的 vscode 扩展](https://github.com/inercia/vscode-k3d) 将其集成到您的 vscode 工作流程中
- 用它来[设置K3s高可用](https://rancher.com/blog/2020/set-up-k3s-high-availability-using-k3d)

You can try all of these yourself by using prepared scripts in [this demo repository](https://github.com/iwilltry42/k3d-demo) or watch us showing them off in [one of our meetups](https://www.youtube.com/watch?v=d9JRb4fk5ag&feature=youtu.be). 

您可以使用 [此演示存储库](https://github.com/iwilltry42/k3d-demo) 中准备好的脚本自行尝试所有这些，或者观看我们在 [我们的一次聚会](https://www.youtube.com/watch?v=d9JRb4fk5ag&feature=youtu.be)。

Other than that, remember that k3d is a [community-driven](https://github.com/rancher/k3d/blob/main/CONTRIBUTING.md) project, so we're always happy to hear from you on [Issues ](https://github.com/rancher/k3d/issues), [Pull Requests](https://github.com/rancher/k3d/issues), [Discussions](https://github.com/rancher/k3d/discussions) and [Slack Chats](https://slack.rancher.io/)!

除此之外，请记住 k3d 是一个 [社区驱动](https://github.com/rancher/k3d/blob/main/CONTRIBUTING.md) 项目，所以我们总是很高兴收到您关于 [问题](https://github.com/rancher/k3d/issues)、[拉取请求](https://github.com/rancher/k3d/issues)、[讨论](https://github.com/rancher/k3d/discussions) 和 [Slack Chats](https://slack.rancher.io/)！

Ready to give k3d a try? [Start by downloading K3s](https://github.com/k3s-io/k3s/releases).

准备好尝试 k3d 了吗？ [从下载 K3s 开始](https://github.com/k3s-io/k3s/releases)。

_Thorsten Klein is a DevOps Engineer at trivago and a freelance software engineer at SUSE. Thorsten is the maintener of k3d. Find him on [Twitter](https://twitter.com/iwilltry42) or visit his [website](https://iwilltry42.dev)._ 

_Thorsten Klein 是 trivago 的 DevOps 工程师和 SUSE 的自由软件工程师。 Thorsten 是 k3d 的维护者。在 [Twitter](https://twitter.com/iwilltry42) 上找到他或访问他的 [网站](https://iwilltry42.dev)。_

