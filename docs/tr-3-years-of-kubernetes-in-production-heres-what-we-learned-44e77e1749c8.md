# 3 Years of Kubernetes in Production–Here’s What We Learned

# Kubernetes 在生产环境中的 3 年——这是我们学到的东西

## Key takeaways from our Kubernetes journal

## Kubernetes 日志中的关键要点

[Sep 9, 2020](https://betterprogramming.pub/3-years-of-kubernetes-in-production-heres-what-we-learned-44e77e1749c8?source=post_page-----44e77e1749c8--------------------------------) · 7 min read

We started out building our first Kubernetes cluster in 2017, version  1.9.4. We had two clusters, one that ran on bare-metal RHEL VMs, and  another that ran on AWS EC2.

我们于 2017 年开始构建我们的第一个 Kubernetes 集群，版本 1.9.4。我们有两个集群，一个在裸机 RHEL VM 上运行，另一个在 AWS EC2 上运行。

Today, our Kubernetes infrastructure fleet consists of over 400 virtual  machines spread across multiple data-centres. The platform hosts  highly-available mission-critical software applications and systems, to  manage a massive live network with nearly four million active devices.

今天，我们的 Kubernetes 基础设施群由分布在多个数据中心的 400 多台虚拟机组成。该平台托管高度可用的关键任务软件应用程序和系统，以管理具有近 400 万台活动设备的大型实时网络。

Kubernetes eventually made our lives easier, but the journey was a hard one, a  paradigm shift. There was a complete transformation in not just our  skillset and tools, but also our design and thinking. We had to embrace  multiple new technologies and invest massively to upscale and upskill  our teams and infrastructure.

Kubernetes 最终让我们的生活变得更轻松，但旅程是艰难的，是范式转变。不仅我们的技能和工具发生了彻底的转变，我们的设计和思维也发生了彻底的转变。我们必须采用多种新技术并进行大量投资，以提升我们的团队和基础设施的规模和技能。

Looking back, after three years of running Kubernetes in production, here are key lessons from our journal.

回顾过去，在生产环境中运行 Kubernetes 三年后，以下是我们日志中的重要经验教训。

[How a Faulty Visa System is Killing the Technology MarketWoes of immigrant software engineers on work visasmedium.com](https://medium.com/digital-diplomacy/how-a-faulty-visa-system-is-killing-the-technology-market-adcd8588d880)

[错误的签证系统如何扼杀技术市场上的移民软件工程师工作visamedium.com的困境](https://medium.com/digital-diploacy/how-a-faulty-visa-system-is-killing-the-技术-市场-adcd8588d880)

# 1. The Curious Case of Java Apps

# 1. Java 应用程序的奇特案例

When it comes to microservices and containerization, engineers tend to steer clear of using Java, primarily due to its notorious memory management. However, things have changed now and Java’s container compatibility has  improved over the years. After all, ubiquitous systems like `Apache Kafka`and `Elasticsearch`run on Java.

当谈到微服务和容器化时，工程师倾向于避开使用 Java，主要是因为它臭名昭著的内存管理。然而，现在情况发生了变化，多年来 Java 的容器兼容性有所提高。毕竟，像 Apache Kafka 和 Elasticsearch 这样无处不在的系统是在 Java 上运行的。

Back in 2017–18, we had a few apps that ran on Java version 8. These often  struggled to understand container environments like Docker and crashed  from heap memory issues and unusual garbage collection trends. We  learned that these were [caused by JVM's inability ](https://developers.redhat.com/blog/2017/03/14/java-inside-docker/)to honor Linux `cgroups`and `namespaces` , that are at the core of containerization technology.

早在 2017-18 年，我们有一些应用程序运行在 Java 版本 8 上。这些应用程序通常难以理解 Docker 等容器环境，并因堆内存问题和不寻常的垃圾收集趋势而崩溃。我们了解到这些是 [由 JVM 无法实现](https://developers.redhat.com/blog/2017/03/14/java-inside-docker/) 来尊重 Linux `cgroups` 和 `namespaces`，它们是容器化技术的核心。

However, since then, Oracle has been continuously improving Java’s compatibility in the container world. Even Java 8's subsequent patches introduced  experimental JVM flags to tackle these problems, `XX:+UnlockExperimentalVMOptions` and `XX:+UseCGroupMemoryLimitForHeap`

但是，此后，Oracle 一直在不断提高 Java 在容器世界中的兼容性。甚至 Java 8 的后续补丁也引入了实验性 JVM 标志来解决这些问题，`XX:+UnlockExperimentalVMOptions` 和 `XX:+UseCGroupMemoryLimitForHeap`

But despite all the improvements, there is no denying that Java still has a bad reputation for hogging memory and its slow startup time compared to its peers like Python or Go. Its primarily caused by JVM’s memory  management and class-loader.

但是，尽管进行了所有改进，但不可否认的是，与 Python 或 Go 等同类产品相比，Java 仍然因占用内存和启动时间缓慢而声名狼藉。它主要是由 JVM 的内存管理和类加载器引起的。

Today, if we *have* to choose Java, we ensure that it’s version 11 or above. And our  Kubernetes memory limits are set to 1GB on top of JVM max heap memory (`-Xmx`) for headroom. That is, if JVM uses 8GB for heap memory, our Kubernetes  resources limits for the app would be 9GB. With that, life has been  better.

今天，如果我们*必须*选择 Java，我们确保它是 11 或更高版本。我们的 Kubernetes 内存限制在 JVM 最大堆内存 (`-Xmx`) 之上设置为 1GB 以增加空间。也就是说，如果 JVM 使用 8GB 的堆内存，我们对应用程序的 Kubernetes 资源限制将为 9GB。有了它，生活变得更好了。

[Why Java Is DyingWhat does the future hold for Java?medium.com](https://medium.com/better-programming/why-java-is-dying-b02b5fd44db9)

[Java 为何消亡 Java 的未来如何？medium.com](https://medium.com/better-programming/why-java-is-dying-b02b5fd44db9)

# 2. Kubernetes Lifecycle Upgrades

# 2. Kubernetes 生命周期升级

Kubernetes lifecycle management such as upgrades or enhancements is cumbersome, especially if you've built your own cluster on [bare metal or VMs](https://platform9.com/blog/where-to-install-kubernetes-bare-metal-vs-vms-vs-cloud/). For upgrades, we’ve realized that the easiest way is to build a new  cluster with the latest version and transition workloads from old to  new. The effort and the planning that goes into in-place node upgrades  are just not worth it.

Kubernetes 生命周期管理（例如升级或增强）很麻烦，尤其是如果您在 [裸机或 VM] 上构建了自己的集群（https://platform9.com/blog/where-to-install-kubernetes-bare-metal- vs-vms-vs-cloud/)。对于升级，我们意识到最简单的方法是构建一个具有最新版本的新集群，并将工作负载从旧的过渡到新的。就地节点升级所付出的努力和规划是不值得的。

Kubernetes has multiple moving parts that need to align with an upgrade. From  Docker to CNI plugins like Calico or Flannel, you need to carefully  piece it all together for it to work. Although projects like Kubespray,  Kubeone, Kops, and Kubeaws make it easier, they all come with  shortcomings. 

Kubernetes 有多个需要与升级保持一致的移动部件。从 Docker 到 Calico 或 Flannel 等 CNI 插件，您需要小心地将它们拼凑在一起才能工作。尽管 Kubespray、Kubeone、Kops 和 Kubeaws 之类的项目使它变得更容易，但它们都有缺点。

We built our clusters using Kubespray on RHEL VMs. Kubespray was great, it had playbooks for building, adding and removing new nodes, upgrading  version, and pretty much everything we needed for operating Kubernetes  in production. But however, the upgrade playbook came with a disclaimer  that prevented us from skipping even minor versions. So one would have  to go through all intermediate versions to reach the target version.

我们在 RHEL 虚拟机上使用 Kubespray 构建了我们的集群。 Kubespray 很棒，它有用于构建、添加和删除新节点、升级版本以及我们在生产中运行 Kubernetes 所需的几乎所有内容的剧本。但是，升级手册附带了一个免责声明，阻止我们跳过即使是次要版本。因此，必须经过所有中间版本才能达到目标版本。

The takeaway is that if you plan to use Kubernetes or are already using  one, think about lifecycle activities and how your solution addresses  that. It’s relatively easier to build and run the cluster, but lifecycle maintenance is a whole new game with multiple moving parts.

要点是，如果您计划使用 Kubernetes 或已经在使用 Kubernetes，请考虑生命周期活动以及您的解决方案如何解决这个问题。构建和运行集群相对容易，但生命周期维护是一个全新的游戏，有多个活动部分。

# 3. Build and Deployment

# 3. 构建和部署

Be prepared to redesign your entire build and deployment pipelines. Our  build process and deployment had to go through a complete transformation for the Kubernetes world. There was a lot of restructuring in not just  Jenkins pipelines, but using new tools like Helm, strategizing new git  flows and builds, tagging docker images, and versioning helm deployment  charts.

准备好重新设计整个构建和部署管道。我们的构建过程和部署必须经过针对 Kubernetes 世界的完整转型。不仅在 Jenkins 管道中进行了大量重组，而且在使用 Helm 等新工具、制定新的 git 流程和构建策略、标记 docker 镜像以及版本控制 helm 部署图表方面进行了大量重组。

You would need a strategy to maintain not just code, but Kubernetes  deployment files, Docker files, Docker images, Helm charts, and design a way to link it all together.

您不仅需要维护代码，还需要维护 Kubernetes 部署文件、Docker 文件、Docker 映像、Helm 图表，并设计一种将它们链接在一起的方法。

After several iterations, we settled on the following design.

经过多次迭代，我们确定了以下设计。

- Application code and its helm charts reside in separate git repositories. This allows us to version them separately. ([semantic versioning](https://semver.org/))
- We then save a map of chart version with the app version and use that for tracking a release. So for example, `app-1.2.0`deployed with `charts-1.1.0` . If only Helm values file were to change, then only the patch version of the chart would change. (e.g. from `1.1.0` to `1.1.1`). All these versions were dictated by release notes in each repository, `RELEASE.txt`.
- System apps like Apache Kafka or Redis whose code we did not build or modify,  worked differently. That is, we did not have two git repositories as  Docker tag was simply part of Helm chart’s versioning. If we ever  changed the docker tag for an upgrade, we would bump up the major  version in the chart’s tag.

- 应用程序代码及其舵图位于单独的 git 存储库中。这允许我们分别对它们进行版本控制。 ([语义版本控制](https://semver.org/))
- 然后我们用应用程序版本保存图表版本的地图，并使用它来跟踪发布。例如，`app-1.2.0` 与 `charts-1.1.0` 一起部署。如果仅更改 Helm 值文件，则仅更改图表的补丁版本。 （例如，从“1.1.0”到“1.1.1”）。所有这些版本都由每个存储库中的发行说明决定，“RELEASE.txt”。
- 像 Apache Kafka 或 Redis 这样的系统应用程序，我们没有构建或修改它们的代码，它们的工作方式不同。也就是说，我们没有两个 git 存储库，因为 Docker 标签只是 Helm chart 版本控制的一部分。如果我们为了升级而更改了 docker 标签，我们会在图表的标签中提升主要版本。

# 4. Liveliness and Readiness Probes (the Double-Edged Sword)

# 4. 活力和准备探针（双刃剑）

Kubernetes’ liveliness and readiness probes are excellent features to combat system problems autonomously. They can restart containers on failures and  divert traffic away from unhealthy instances. But in certain failure  conditions, these probes can become a double-edged sword and affect your application’s startup and recovery, particularly, stateful applications like messaging platforms or databases.

Kubernetes 的活跃度和就绪度探测器是自主解决系统问题的出色功能。他们可以在出现故障时重新启动容器，并将流量从不健康的实例转移。但在某些故障情况下，这些探测可能会成为一把双刃剑，并影响应用程序的启动和恢复，尤其是消息平台或数据库等有状态应用程序。

Our Kafka system was a victim of this. We ran a `3 Broker 3 Zookeeper` stateful set with a `replicationFactor of 3 `and a `minInSyncReplica of 2.` The issue occurred when Kafka started up after accidental system  failures or crashes. This caused it to run additional scripts during  startup to fix corrupted indices, which took anywhere between 10 to 30  mins depending on the severity. Because of this added time, the  liveliness probes would constantly fail, issuing a kill signal to Kafka  to restart. This prevented Kafka from ever fixing up the indices and  starting up altogether.

我们的 Kafka 系统就是这种情况的受害者。我们运行了一个 `3 Broker 3 Zookeeper` 状态集，`replicationFactor 为 3`，`minInSyncReplica 为 2`。当 Kafka 在意外系统故障或崩溃后启动时出现问题。这导致它在启动期间运行额外的脚本来修复损坏的索引，根据严重程度，这需要 10 到 30 分钟之间的任何时间。由于这个额外的时间，活跃度探测器会不断失败，向 Kafka 发出终止信号以重新启动。这阻止了 Kafka 修复索引并完全启动。

The only solution is to configure `**initialDelaySeconds**` in liveliness probe settings to delay probe evaluations after the  container startup. But the problem, of course, is that its hard to put a number to this. Some recoveries even take even an hour, and we need to  provide enough headroom to account for this. But the more you increase `**initialDelaySeconds**` , the slower your resilience, as it would take longer for Kubernetes to restart your container during startup failures.

唯一的解决方案是在活跃度探测器设置中配置`**initialDelaySeconds**`，以在容器启动后延迟探测器评估。但是，当然，问题是很难给出一个数字。有些恢复甚至需要一个小时，我们需要提供足够的空间来解决这个问题。但是，`**initialDelaySeconds**` 增加得越多，你的弹性就越慢，因为在启动失败期间 Kubernetes 需要更长的时间来重启你的容器。

So the middle ground is to assess a value for the `**initialDelaySeconds**` field such that it better balances between the resilience you seek in  Kubernetes and the time taken by the app to successfully start in all  fault conditions (disk failures, network failures, system crashes, etc.) 

因此，中间立场是评估 `**initialDelaySeconds**` 字段的值，以便更好地平衡您在 Kubernetes 中寻求的弹性和应用程序在所有故障情况下（磁盘故障、网络故障）成功启动所花费的时间故障、系统崩溃等）

> ***Update\****: If you are on the last few latest releases,* [*Kubernetes has introduced a third probe-type called, 'Startup Probe,' to tackle this problem*](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)*. It is available in* `*alpha from 1.16* `*and* `*beta from 1.18*` *onwards.*
>
> *A startup probe disables readiness and liveliness checks until the  container has started up, making sure the application’s startup isn’t  interrupted.*

> ***更新\****：如果您使用的是最近几个最新版本，* [*Kubernetes 引入了第三种探针类型，称为“启动探针”，以解决此问题*](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)*。它可用于* `*alpha from 1.16* `* and* `*beta from 1.18*` *onwards.*
>
> *启动探测器禁用准备和活跃度检查，直到容器启动，确保应用程序的启动不被中断。*

# 5. Exposing External IPs

# 5. 暴露外部 IP

We learned that exposing services using static external IP takes a huge  toll on your kernel’s connection tracking mechanism. It simply breaks  down at scale unless planned thoroughly.

我们了解到，使用静态外部 IP 公开服务会对内核的连接跟踪机制造成巨大损失。除非经过周密的计划，否则它只会大规模崩溃。

Our cluster runs on `Calico for CNI `and `BGP` as our routing protocol inside Kubernetes and also to peer with edge routers. For Kubeproxy, we use `IP Tables`mode. We host a massive service in our Kubernetes exposed via external IP  that handles millions of connections every day. Because of all the SNAT  and masquerading that comes from software-defined networks, Kubernetes  needs a mechanism to track all these logical flows. To achieve this, it  uses the Kernel’s `Conntrack and netfilter` tools to manage these external connections to the static IP, which then translates to internal service IP and then to your pod IP. This is all  done through the`conntrack` table and IP Tables.

我们的集群在 CNI 的 Calico 和 BGP 上运行，作为我们在 Kubernetes 内部的路由协议，并与边缘路由器对等。对于 Kubeproxy，我们使用 `IP Tables` 模式。我们在通过外部 IP 公开的 Kubernetes 中托管了大量服务，每天处理数百万个连接。由于所有来自软件定义网络的 SNAT 和伪装，Kubernetes 需要一种机制来跟踪所有这些逻辑流。为此，它使用内核的“Conntrack 和 netfilter”工具来管理这些与静态 IP 的外部连接，然后将其转换为内部服务 IP，然后转换为您的 pod IP。这都是通过`conntrack` 表和IP 表完成的。

This `conntrack` table, however, has its limits. Once you hit the limit, your Kubernetes cluster (the OS Kernel underneath) will not be able to accept new  connections any more. On RHEL, you can check it this way.

然而，这个 `conntrack` 表有其局限性。一旦达到限制，您的 Kubernetes 集群（下面的操作系统内核）将无法再接受新连接。在 RHEL 上，您可以通过这种方式进行检查。

```
$  sysctl net.netfilter.nf_conntrack_count net.netfilter.nf_conntrack_maxnet.netfilter.nf_conntrack_count = 167012
net.netfilter.nf_conntrack_max = 262144
```

Some ways to work around this is to peer multiple nodes with edge routers so that the incoming connection to your static IP is sprayed across your  cluster. So if your cluster has a large fleet of machines, cumulatively  you could have a large `conntrack` table to handle massive incoming connections.

解决此问题的一些方法是将多个节点与边缘路由器对等，以便将传入的静态 IP 连接发送到整个集群。因此，如果您的集群拥有大量机器，那么累积起来您可以拥有一个大的“conntrack”表来处理大量传入连接。

Back when we started in 2017, this completely threw us off, but recently, a  detailed study on this was published by Calico in 2019, aptly titled “[Why conntrack is no longer your friend](https://www.projectcalico.org/when-linux-conntrack-is-no-longer-your-friend/).”

当我们在 2017 年开始时，这完全让我们失望，但最近，Calico 在 2019 年发表了一篇关于此的详细研究，标题恰如其分地题为“[为什么 conntrack 不再是你的朋友](https://www.projectcalico.org)/when-linux-conntrack-is-no-longer-your-friend/)。”

[4 Simple Kubernetes Terminal Customizations to Boost Your ProductivityThis is what I use for managing large-scale Kubernetes clusters in productionmedium.com](https://medium.com/better-programming/4-simple-kubernetes-terminal-customizations-to-boost-your-productivity-deda60a19924)

[4 个简单的 Kubernetes 终端定制来提高你的生产力这是我在 productionmedium.com 管理大型 Kubernetes 集群时使用的](https://medium.com/better-programming/4-simple-kubernetes-terminal-customizations-to-提高你的生产力-deda60a19924)

# Do you absolutely need Kubernetes?

# 你绝对需要Kubernetes吗？

Three years in and we still continue to discover and learn something new  every day. It is a complex platform with its own set of challenges,  particularly, the overhead in building and maintaining the environment. It will change your design, thinking, architecture, and will require  upskilling and upscaling your teams to meet the transformation.

三年过去了，我们仍然每天都在继续发现和学习新的东西。这是一个复杂的平台，有其自身的一系列挑战，尤其是构建和维护环境的开销。它将改变您的设计、思维、架构，并且需要提升您的团队的技能和规模以适应转型。

However, if you are on the cloud and are able to use Kubernetes as a “service,”  it can relieve you from most of that overhead that comes with platform  maintenance like “How do I expand my internal network CIDR?” or “How do I upgrade my Kubernetes version?”

但是，如果您在云上并且能够将 Kubernetes 用作“服务”，那么它可以将您从平台维护带来的大部分开销中解放出来，例如“如何扩展我的内部网络 CIDR？”或“如何升级我的 Kubernetes 版本？”

Today, we’ve realized that the first question you need to ask yourself is “Do you *absolutely* need Kubernetes?” This can help assess the problem you have and how significantly or not, Kubernetes addresses it.

今天，我们意识到您需要问自己的第一个问题是“您*绝对*需要 Kubernetes 吗？”这可以帮助评估您遇到的问题以及 Kubernetes 解决它的重要性与否。

Kubernetes transformation is not cheap. The price you pay for it must really  justify ‘your’ use case and how it leverages the platform. If it does,  then Kubernetes can immensely boost your productivity.

Kubernetes 转型并不便宜。你为它付出的代价必须真正证明“你的”用例以及它如何利用平台。如果是这样，那么 Kubernetes 可以极大地提高您的生产力。

> Remember, technology for the sake of technology is meaningless.

> 请记住，为了技术而技术是没有意义的。

[Better Programming](https://betterprogramming.pub/?source=post_sidebar--------------------------post_sidebar-----------)

Advice for programmers. 

给程序员的建议。

