# Kubernetes vs Docker: Understanding Containers in 2021

# 2021 年了解 Kubernetes vs Docker

> From https://semaphoreci.com/blog/kubernetes-vs-docker

[Tomas Fernandez](https://semaphoreci.com/author/tfernandez)[Twitter](https://twitter.com/TomFernBlog)   3 Feb 2021 ·  [Software Engineering](https://semaphoreci.com/category/ engineering)

A few weeks ago, the Kubernetes development team announced that they are [deprecating Docker](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#deprecation). This piece of news made the rounds through tech communities and social  networks alike. Will Kubernetes clusters break, and if so, how will we  run our applications? What should we do now? Today, we’ll examine all  these questions and more.

几周前，Kubernetes 开发团队宣布他们正在 [弃用 Docker](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#deprecation) 这条新闻在科技社区和社交网络上广为流传。 Kubernetes 集群会不会崩溃，如果会，我们将如何运行我们的应用程序？我们现在该做什么？今天，我们将研究所有这些问题以及更多问题。

Let’s start from the top. If you're already familiar with Docker and Kubernetes and want to get to the juicy parts, skip to *[how does the Dockershim deprecation impact you?](https://semaphoreci.com/blog/kubernetes-vs-docker#h -how-does-kubernetes-deprecating-docker-impact-you)*

让我们从顶部开始。如果您已经熟悉 Docker 和 Kubernetes 并想了解有趣的部分，请跳至 *[Dockershim 弃用对您有何影响？](https://semaphoreci.com/blog/kubernetes-vs-docker#h -how-does-kubernetes-deprecating-docker-impact-you)*

## What is a Container?

## 什么是容器？

Even though Docker is used as a synonym for containers, the reality  is that they have existed long before Docker was a thing. Unix and Linux have had containers in some form or another since the late 70s, when [chroot](https://man7.org/linux/man-pages/man2/chroot.2.html) was introduced. Chroot allowed system admins to run programs in a  kind-but-not-really-isolated filesystem. Later, the idea was refined and enhanced into container engines such as [FreeBSD Jails](https://docs-dev.freebsd.org/en/books/handbook/jails), [OpenVZ](https://openvz.org ), or [Linux Containers (LXC)](https://linuxcontainers.org/).

尽管 Docker 被用作容器的同义词，但事实是它们早在 Docker 出现之前就已经存在。自从 70 年代末引入 [chroot](https://man7.org/linux/man-pages/man2/chroot.2.html) 以来，Unix 和 Linux 已经有了某种形式的容器。 Chroot 允许系统管理员在一个友好但并非真正隔离的文件系统中运行程序。后来，这个想法被提炼并增强到容器引擎中，例如[FreeBSD Jails](https://docs-dev.freebsd.org/en/books/handbook/jails)、[OpenVZ](https://openvz.org ) 或 [Linux 容器 (LXC)](https://linuxcontainers.org/)。

But what are containers?

但是什么是容器？

A container is a logical partition where we can run applications  isolated from the rest of the system. Each application gets its own  private network and a virtual filesystem that is not shared with other  containers or the host.

容器是一个逻辑分区，我们可以在其中运行与系统其余部分隔离的应用程序。每个应用程序都有自己的私有网络和一个不与其他容器或主机共享的虚拟文件系统。

![Containers](https://wpblog.semaphoreci.com/wp-content/uploads/2021/01/container-arch.png)

Running containerized applications is a lot more convenient than  installing and configuring software. For one thing, containers are  portable; we can build in one server with the confidence that it will  work in any server. Another advantage is that we can run multiple copies of the same program simultaneously without conflict or overlap,  something really hard to do otherwise.

运行容器化应用程序比安装和配置软件要方便得多。一方面，容器是便携的；我们可以在一台服务器中构建，并确信它可以在任何服务器上运行。另一个优点是我们可以同时运行同一程序的多个副本而不会发生冲突或重叠，否则很难做到这一点。

However, for all this to work, we need a *container runtime*, a piece of software capable of running containers.

然而，为了让所有这些工作，我们需要一个*容器运行时*，一个能够运行容器的软件。

## What is Docker?

## Docker 是什么？

Docker is the most popular container runtime — by a long shot. It  shouldn’t be surprising, as it brought the concept of containers into  the mainstream, which in turn inspired the creation of platforms like  Kubernetes.

Docker 是最受欢迎的容器运行时——从长远来看。这不足为奇，因为它将容器的概念带入了主流，进而激发了 Kubernetes 等平台的创建。

Before Docker, running containers was indeed possible, but it was  hard work. Docker made things simple because it’s a complete tech stack  that can:

- Manage container lifecycle.
 - Proxy requests to and from the containers.
 - Monitor and log container activity.
 - Mount shared directories.
 - Set resource limits on containers.
 - Build images. The `Dockerfile` is the de-facto format for building container images.
 - Push and pull images from registries.

在 Docker 之前，运行容器确实是可能的，但这是一项艰苦的工作。 Docker 使事情变得简单，因为它是一个完整的技术堆栈，可以：

- 管理容器生命周期。
- 代理请求进出容器。
- 监控和记录容器活动。
- 挂载共享目录。
- 对容器设置资源限制。
- 构建图像。 `Dockerfile` 是构建容器镜像的事实上的格式。
- 从注册表中推送和拉取镜像。

In its first iterations, Docker used Linux Containers (LXC) as the runtime backend. As the project evolved, LXC was replaced by [containerd](https://containerd.io/), Docker’s own implementation. A modern Docker installation is divided into two services: `containerd`, responsible for managing containers, and `dockerd`, which does all the rest.

在第一次迭代中，Docker 使用 Linux Containers (LXC) 作为运行时后端。随着项目的发展，LXC 被 Docker 自己的实现 [containerd](https://containerd.io/) 取代。现代 Docker 安装分为两个服务：`containerd`，负责管理容器，以及 `dockerd`，完成其余所有工作。

![Docker 引擎](https://wpblog.semaphoreci.com/wp-content/uploads/2021/01/docker-arch.png)

## What is Kubernetes?

## 什么是Kubernetes？

Kubernetes takes the idea of containers and turns it up a notch. Instead of running containerized applications in a single server,  Kubernetes distributes them across a cluster of machines. Applications  running in Kubernetes look and behave like a single unit, even though,  in reality, they may consist of an arrangement of loosely-coupled  containers.

Kubernetes 采用了容器的概念并将其提升了一个档次。 Kubernetes 不是在单个服务器中运行容器化应用程序，而是将它们分布在一组机器上。在 Kubernetes 中运行的应用程序看起来和行为就像一个单一的单元，尽管实际上它们可能由松散耦合的容器排列组成。

Kubernetes adds distributed computing features on top of containers:

- **Pods**: pods are logical groups of containers that share resources like memory, CPU, storage, and network.
 - **Auto-scaling**: Kubernetes can automatically adapt to changing workloads by starting and stopping pods as needed.
 - **Self-healing**: containers are monitored and restarted on failure.
 - **Load-balancing**: requests are distributed over the healthy available pods. 
 - **Rollouts**: Kubernetes supports automated rollouts and rollbacks. Making otherwise complex procedures like [Canary](https://semaphoreci.com/blog/what-is-canary-deployment) and [Blue-Green](https://semaphoreci.com/blog/continuous-blue-green- deployments-with-kubernetes) releases trivial.

Kubernetes 在容器之上增加了分布式计算特性：

- **Pods**：Pods 是共享内存、CPU、存储和网络等资源的逻辑容器组。

- **自动缩放**：Kubernetes 可以通过根据需要启动和停止 Pod 来自动适应不断变化的工作负载。

- **自我修复**：容器被监控并在失败时重新启动。

- **负载平衡**：请求分布在健康的可用 Pod 上。

- **Rollouts**：Kubernetes 支持自动推出和回滚。制作其他复杂的程序，如 [Canary](https://semaphoreci.com/blog/what-is-canary-deployment) 和 [Blue-Green](https://semaphoreci.com/blog/continuous-blue-green-deployments-with-kubernetes)发布。

  

We can think of Kubernetes’ architecture as a combination of two planes:

- The **control plane** is the coordinating brain of the cluster. It has a *controller* that manages nodes and services, a *scheduler* that assigns pods to the nodes, and the *API service*, which handles communication. Configuration and state are stored on a highly-available database called *etcd*.
 - The **worker nodes** are the machines that run the containers. Each worker node runs a few components like the *kubelet* agent, a network proxy, and the container runtime. The default container runtime up to Kubernetes version v1.20 was Docker.

我们可以将 Kubernetes 的架构视为两个平面的组合：

- **控制平面**是集群的协调大脑。它有一个管理节点和服务的*控制器*，一个为节点分配 pod 的*调度程序*，以及处理通信的 *API 服务*。配置和状态存储在名为 *etcd* 的高可用性数据库中。
- **工作节点**是运行容器的机器。每个工作节点运行一些组件，如 *kubelet* 代理、网络代理和容器运行时。 Kubernetes 版本 v1.20 之前的默认容器运行时是 Docker。

![control plane](https://wpblog.semaphoreci.com/wp-content/uploads/2021/01/k8s-arch.png)

## Container Formats

## 容器格式

Before starting a container, we need to either build or download a *container image*, which is a filesystem packed with everything the application needs:  code, binaries, configuration files, libraries, and dependencies.

在启动容器之前，我们需要构建或下载*容器映像*，这是一个包含应用程序所需一切的文件系统：代码、二进制文件、配置文件、库和依赖项。

The rise in popularity of containers showed the need for an open  image standard. As a result, Docker Inc and CoreOS established the [Open Container Initiative](https://opencontainers.org/) (OCI) in 2015, with the mission of producing vendor-neutral formats. The result of this effort was the creation of two standards:

- An image specification that defines the image binary format.
 - A [runtime specification](https://github.com/opencontainers/runtime-spec) that describes how to unpack and run a container. OCI maintains a reference implementation called [runc](https://github.com/opencontainers/runc). Both containerd and CRI-O use runc in the background to spawn containers.

容器的流行表明需要一个开放的镜像标准。因此，Docker Inc 和 CoreOS 在 2015 年成立了 [Open Container Initiative](https://opencontainers.org/) (OCI)，其使命是生产供应商中立的格式。这项努力的结果是创建了两个标准：

- 定义镜像二进制格式的镜像规范。
- [运行时规范](https://github.com/opencontainers/runtime-spec)，描述了如何解包和运行容器。 OCI 维护一个名为 [runc](https://github.com/opencontainers/runc) 的参考实现。 containerd 和 CRI-O 都在后台使用 runc 来生成容器。

The OCI standard brought interoperability among different container  solutions. As a result, images built in one system can run in any other  compliant stack.

OCI 标准带来了不同容器解决方案之间的互操作性。因此，在一个系统中构建的映像可以在任何其他兼容堆栈中运行。

![Kubernetes vs. Docker](https://wpblog.semaphoreci.com/wp-content/uploads/2021/01/oci-interoperability.png)

## Docker Vs. Kubernetes

## Docker 对比。 Kubernetes

Here is where things get a bit more technical. I said that each Kubernetes worker node needs a container runtime. In its first [original design](https://github.com/kubernetes/kubernetes/blob/release-0.4/DESIGN.md), Docker was inseparable from Kubernetes because it was the only runtime supported.

这是事情变得更加技术化的地方。我说过每个 Kubernetes 工作节点都需要一个容器运行时。在其第一个 [原始设计](https://github.com/kubernetes/kubernetes/blob/release-0.4/DESIGN.md) 中，Docker 与 Kubernetes 密不可分，因为它是唯一受支持的运行时。

Docker, however, was never designed to run inside Kubernetes. Realizing this problem, the Kubernetes developers eventually implemented an API called *Container Runtime Interface* (CRI). This  interface allows us to choose among different container runtimes, making the platform more flexible and less dependent on Docker.

然而，Docker 从未设计为在 Kubernetes 内部运行。意识到这个问题后，Kubernetes 开发人员最终实现了一个名为 *Container Runtime Interface* (CRI) 的 API。该接口允许我们在不同的容器运行时中进行选择，从而使平台更加灵活并减少对 Docker 的依赖。

![Container runtime](https://wpblog.semaphoreci.com/wp-content/uploads/2021/01/cri.png)

This change introduced a new difficulty for the Kubernetes team since Docker doesn’t know about or support the CRI. Hence, at the same time  the API was introduced, they had to write an adaptor called *Dockershim* to translate CRI messages into Docker-specific commands.

这一变化给 Kubernetes 团队带来了新的困难，因为 Docker 不知道或不支持 CRI。因此，在引入 API 的同时，他们必须编写一个名为 *Dockershim* 的适配器来将 CRI 消息转换为 Docker 特定的命令。

## The Dockershim Deprecation

## Dockershim 弃用

While Docker was the first and only supported engine for a time, it was never on the long-term plans. [Kubernetes version 1.20 deprecates Dockershim](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#deprecation), kicking off the transition away from Docker.Once the transition is done, the stack gets significantly smaller. It goes from this:

虽然 Docker 曾经是第一个也是唯一受支持的引擎，但它从来没有出现在长期计划中。 [Kubernetes 1.20 版弃用 Dockershim](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#deprecation)，开始从 Docker 过渡。一旦转换完成，堆栈就会明显变小。它来自于：

![Dockershim 弃用](https://wpblog.semaphoreci.com/wp-content/uploads/2021/01/kubelet-dockershim.png)

To this:

对此：

![Dockershim 弃用](https://wpblog.semaphoreci.com/wp-content/uploads/2021/01/kubelet-containerd.png)

The result is less bloat and fewer dependencies needed on each of the worker nodes.

结果是每个工作节点所需的臃肿和依赖更少。

So, why the change?

那么，为什么要改变？

Simply put, Docker is heavy. We get better performance with a lightweight container runtime like containerd or [CRI-O](https://cri-o.io/). As a recent example, Google [benchmarks](https://kubernetes.io/blog/2018/05/24/kubernetes-containerd-integration-goes-ga/) have shown that containerd consumes less memory and CPU, and that pods start in less time than on Docker.

简单地说，Docker 很重。我们使用轻量级容器运行时（如 containerd 或 [CRI-O](https://cri-o.io/)）获得更好的性能。举个最近的例子，谷歌 [benchmarks](https://kubernetes.io/blog/2018/05/24/kubernetes-containerd-integration-goes-ga/) 表明 containerd 消耗更少的内存和 CPU，并且 pods比在 Docker 上启动的时间更短。

Besides, in some ways Docker itself can be considered [technical debt](https://www.tariqislam.com/posts/kubernetes-docker-dep/). What Kubernetes needs from Docker is, in fact, the container runtime:  containerd. The rest, at least as far as Kubernetes is concerned, is  overhead.

此外，Docker 本身在某些方面可以被视为 [技术债务](https://www.tariqislam.com/posts/kubernetes-docker-dep/)。实际上，Kubernetes 需要 Docker 的是容器运行时：containerd。其余的，至少就 Kubernetes 而言，是开销。

## How Does Kubernetes Deprecating Docker Impact You? 
## Kubernetes 弃用 Docker 对您有何影响？
Things are not as dramatic as they sound. Let’s preface this whole  section by saying the only thing that changes in v1.20 is that you’ll  get a deprecation warning, only if you’re running Docker. That’s all.

事情并不像听起来那么戏剧化。让我们在这整个部分的开头说 v1.20 中唯一改变的是你会得到一个弃用警告，只有当你运行 Docker 时。就这样。

**Can I still use Docker for development?**

**我还可以使用 Docker 进行开发吗？**

Yes, you absolutely can, now and in the foreseeable future. You see,  Docker doesn’t run Docker-specific images; it runs OCI-compliant  containers. As long as Docker continues using this format, Kubernetes  will keep accepting them.

是的，您现在和在可预见的未来绝对可以。你看，Docker 不运行 Docker 特定的镜像；它运行符合 OCI 的容器。只要 Docker 继续使用这种格式，Kubernetes 就会继续接受它们。

**Can I Still Package My Production Apps With Docker?**

**我还可以使用 Docker 打包我的生产应用程序吗？**

Yes, for the same reasons as in the previous question. Applications  packaged with Docker will continue to run — no change there. Thus, you  can still build and test containers with the tools you know and love. You don't need to change your [CI/CD pipelines](https://semaphoreci.com/blog/cicd-pipeline) or switch to other image registries, Docker-produced images will continue to work in your cluster just as they always have.

是的，原因与上一个问题相同。使用 Docker 打包的应用程序将继续运行——那里没有变化。因此，您仍然可以使用您熟悉和喜爱的工具构建和测试容器。您无需更改 [CI/CD 管道](https://semaphoreci.com/blog/cicd-pipeline) 或切换到其他镜像注册表，Docker 生成的镜像将继续在您的集群中正常工作一直有。

**What Do I Need to Change?**

**我需要改变什么？**

Right now, nothing. If your cluster uses Docker as a runtime, you’ll  get a deprecation warning after upgrading to v1.20. But the change is a  clear signal from the Kubernetes community about the direction they want to take. It’s time to start planning for the future.

现在，什么都没有。如果您的集群使用 Docker 作为运行时，您将在升级到 v1.20 后收到弃用警告。但这一变化是 Kubernetes 社区关于他们想要采取的方向的明确信号。是时候开始规划未来了。

**When Is the Change Going to Happen?**

**什么时候会发生变化？**

The plan is to have all Docker dependencies completely removed by [v1.23 in late 2021](https://kubernetes.io/blog/2020/12/02/dont-panic-kubernetes-and-docker/#so-why-the-confusion-and-what-is-everyone-freaking-out-about).

计划是在 [v1.23 在 2021 年末](https://kubernetes.io/blog/2020/12/02/dont-panic-kubernetes-and-docker/#so-why) 完全删除所有 Docker 依赖项）。

**When Dockershim Goes Away, What Will Happen?**

**当 Dockershim 消失时，会发生什么？**

At that point, Kubernetes cluster admins will be forced to switch to a CRI-compliant container runtime.

届时，Kubernetes 集群管理员将被迫切换到符合 CRI 标准的容器运行时。

**If you are an end-user** not a lot changes for you. Unless you are running some kind of [node customizations](https://kubernetes.io/docs/concepts/extend-kubernetes/extend-cluster/), you probably won’t have to do anything special. Only [test](https://semaphoreci.com/blog/automated-testing-cicd) that your applications work with the new container runtime.

**如果您是最终用户** 对您来说变化不大。除非您正在运行某种[节点定制](https://kubernetes.io/docs/concepts/extend-kubernetes/extend-cluster/)，否则您可能不需要做任何特别的事情。只有 [测试](https://semaphoreci.com/blog/automated-testing-cicd) 您的应用程序可以使用新的容器运行时。

These are some of the things that will cause problems or break after upgrading to v1.23:

这些是升级到 v1.23 后会出现问题或崩溃的一些事情：

- Using Docker-specific logging and monitoring. That is, parsing docker messages from a log or polling the Docker API.
 - Using Docker optimizations.
 - Running scripts that rely on `docker` CLI.
 - Running Docker commands in privileged pods. For instance: to build images with `docker build`. See projects like [kaniko](https://github.com/GoogleContainerTools/kaniko) for alternative solutions.
 - Using Docker-in-Docker setups.
 - Running Windows containers. Containerd does work in Windows, but its support  level is not yet up to par with Docker’s. The objective is to have a  stable containerd release for Windows by [containerd version 1.20](https://github.com/kubernetes/enhancements/issues/1001).

- 使用 Docker 特定的日志记录和监控。也就是说，从日志中解析 docker 消息或轮询 Docker API。
- 使用 Docker 优化。
- 运行依赖于`docker` CLI 的脚本。
- 在特权 Pod 中运行 Docker 命令。例如：使用`docker build`构建镜像。有关替代解决方案，请参阅 [kaniko](https://github.com/GoogleContainerTools/kaniko) 等项目。
- 使用 Docker-in-Docker 设置。
- 运行 Windows 容器。 Containerd 确实可以在 Windows 中运行，但它的支持水平还没有达到 Docker 的水平。目标是通过 [containerd version 1.20](https://github.com/kubernetes/enhancements/issues/1001) 为 Windows 发布一个稳定的 containerd 版本。

**If you’re using a managed cluster** on a cloud  provider like AWS EKS, Google GKE, or Azure AKS, check that your cluster uses a supported runtime before Docker support goes away. Some cloud  vendors are a few versions behind, so you may have more time to plan. So, check with your provider. To give an example, Google Cloud announced they are changing the default runtime from Docker to containerd for all newly-created worker nodes, but you can still opt-in for Docker.

**如果您在 AWS EKS、Google GKE 或 Azure AKS 等云提供商上使用托管集群**，请在 Docker 支持消失之前检查您的集群是否使用受支持的运行时。一些云供应商的版本落后，因此您可能有更多时间进行计划。因此，请咨询您的提供商。举个例子，谷歌云宣布他们将所有新创建的工作节点的默认运行时从 Docker 更改为 containerd，但您仍然可以选择加入 Docker。

**If you run your own cluster**: in addition to checking the points mentioned above, you will need to evaluate moving to another container runtime that is fully compatible with CRI. The Kubernetes  docs explain the steps in detail:

- [Switching to containerd](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#containerd)
 - [Switching to CRI-O](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#cri-o)

**如果您运行自己的集群**：除了检查上述要点之外，您还需要评估迁移到另一个与 CRI 完全兼容的容器运行时。 Kubernetes 文档详细解释了这些步骤：

- [切换到 containerd](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#containerd)
- [切换到 CRI-O](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#cri-o)

Alternatively, if you want to keep on using Docker past version 1.23, follow the [cri-dockerd](https://github.com/Mirantis/cri-dockerd) project, which [plans to keep Docker](https:// www.mirantis.com/blog/mirantis-to-take-over-support-of-kubernetes-dockershim-2/) as a viable runtime alternative.

或者，如果您想在 1.23 版本之后继续使用 Docker，请遵循 [cri-dockerd](https://github.com/Mirantis/cri-dockerd) 项目，该项目[计划保留 Docker](https:// www.mirantis.com/blog/mirantis-to-take-over-support-of-kubernetes-dockershim-2/) 作为可行的运行时替代方案。

## Conclusion

## 结论

Kubernetes is growing, but the change doesn’t need to be a traumatic  experience. Most users won’t have to take any action. For those who do,  there’s still time to test and plan.

Kubernetes 正在增长，但这种变化并不一定是一种创伤性的经历。大多数用户不必采取任何行动。对于那些这样做的人，还有时间进行测试和计划。

To continue learning about Docker and Kubernetes, read these next:

要继续了解 Docker 和 Kubernetes，请阅读以下内容：

- [Download our free book: CI/CD with Docker and Kubernetes](https://semaphoreci.com/resources/cicd-docker-kubernetes)
 - [Powerful CI/CD for Docker and Kubernetes](https://semaphoreci.com/product/docker)
 - [A Step-by-Step Guide to Continuous Deployment on Kubernetes](https://semaphoreci.com/blog/guide-continuous-deployment-kubernetes)
 - [Docker Image Size – Does It Matter?](https://semaphoreci.com/blog/2018/03/14/docker-image-size.html) 

