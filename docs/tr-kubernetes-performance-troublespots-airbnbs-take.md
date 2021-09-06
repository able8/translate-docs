# Kubernetes Performance Trouble Spots: Airbnb’sTake

# Kubernetes 性能问题：Airbnb'sTake

#### 7 Jan 2020

#### 

Now that organizations are starting to rely on Kubernetes and containers in general, performance becomes a major focus point for admins, particularly for public-facing high-use services, such as [Airbnb](https://www.airbnb.com/) . Engineers from the company shared some lessons learned on this topic at KubeCon+CloudNativeCon North America 2019.

现在组织开始普遍依赖 Kubernetes 和容器，性能成为管理员的主要关注点，尤其是面向公众的高使用服务，例如 [Airbnb](https://www.airbnb.com/) .该公司的工程师在 KubeCon+CloudNativeCon North America 2019 上分享了有关该主题的一些经验教训。

In their talk, “ [Did Kubernetes Make My p95s Worse?](https://www.youtube.com/watch?v=QXApVwRBeys&t=3s)” Airbnb software engineers [Stephen Chan](https://www.linkedin.com/in/stephenyehengchan/), who is at the company's compute infrastructure team, and [Jian Cheung](https://www.linkedin.com/in/jian-cheung/), who works in the service orchestration teams, discussed the performance gotcha they've witnessed working with the open source container orchestration engine.

在他们的演讲中，“[Kubernetes 是否让我的 p95s 变得更糟？](https://www.youtube.com/watch?v=QXApVwRBeys&t=3s)”Airbnb 软件工程师 [Stephen Chan](https://www.linkedin.com/in/stephenyehengchan/)，他在公司的计算基础设施团队，和[Jian Cheung](https://www.linkedin.com/in/jian-cheung/)，在服务编排团队工作，讨论他们目睹了使用开源容器编排引擎的性能问题。

Since 2018, the online housing marketplace has been the process of moving its services residing directly on AWS EC2 instance to its own Kubernetes-managed containers, presently about 1,000 services in all. As a result, Airbnb developers are quick to ask service orchestration team, “Why is my pod so slow?” The company runs [Amazon Linux 2](https://aws.amazon.com/amazon-linux-2/) for minion instances, Ubuntu images, the Flannel/Calico integration [Canal](https://github.com/projectcalico/canal) for the container networking, and K8s [NodePort](https://kubernetes.io/docs/concepts/services-networking/service/) to interface with the company's service discovery mechanism.

自 2018 年以来，在线住房市场一直在将其直接驻留在 AWS EC2 实例上的服务转移到其自己的 Kubernetes 管理的容器中，目前总共约有 1,000 项服务。因此，Airbnb 开发人员很快就问服务编排团队：“为什么我的 pod 这么慢？”该公司运行 [Amazon Linux 2](https://aws.amazon.com/amazon-linux-2/) 用于 minion 实例、Ubuntu 映像、Flannel/Calico 集成 [Canal](https://github.com/projectcalico/canal) 用于容器网络，以及 K8s [NodePort](https://kubernetes.io/docs/concepts/services-networking/service/) 与公司的服务发现机制接口。

At the conference, the engineers shared some performance issues they’ve encountered, as well as potential solutions. Their overall message was clear: When dealing with a complex Kubernetes-based infrastructure, performance tuning must be done all across the stack, including host, cluster, container, networking, and even with the underlying applications.

在会议上，工程师们分享了他们遇到的一些性能问题，以及可能的解决方案。他们的总体信息很明确：在处理复杂的基于 Kubernetes 的基础设施时，必须在整个堆栈中进行性能调优，包括主机、集群、容器、网络，甚至底层应用程序。

## Those Noisy Neighbors Again

## 又是那些吵闹的邻居

Why do some pods have more latency than their peers in a cluster? One of the first culprits to check may be a neighboring pod, one that may be hogging all the CPU and networking resources for a heavy workload, the duo advised. Sometimes, the noise is purely accidental, when one hungry service meant to stay in staging was moved to a production cluster instead. But there are also control knobs that must be set as well: Airbnb made a choice early on not to enact CPU limits on its services, which limit the amount of resources that a service would take from its host CPU, this proved to be a bad idea, and the company has since set the resource limits from Kubernetes.

为什么某些 Pod 比集群中的对等节点具有更多延迟？两人建议说，首先要检查的罪魁祸首之一可能是相邻的 Pod，它可能会占用所有 CPU 和网络资源以应对繁重的工作负载。有时，噪音纯粹是偶然的，当一个本应保持暂存状态的饥饿服务被转移到生产集群时。但也有一些控制旋钮也必须设置：Airbnb 很早就做出了不对其服务实施 CPU 限制的选择，这限制了服务从其主机 CPU 获取的资源量，事实证明这是一个坏主意想法，该公司此后设置了 Kubernetes 的资源限制。

[![](https://cdn.thenewstack.io/media/2019/11/6809bec7-airbnb-noisu_neighbors.jpg)](https://static.sched.com/hosted_files/kccncna19/e1/%5Bkubecon%202019%5D%20Did%20Kubernetes%20make%20my%20p95s%20worse.pdf)



The “noisy neighbor problem” is not new to Kubernetes. It was first encountered, and mitigated against, when multiple virtual machines were first packed into servers, and a VM with a CPU-hungry app would hog all the resources, to the detriments of others. 

“嘈杂的邻居问题”对 Kubernetes 来说并不陌生。当多个虚拟机首次打包到服务器中时，首次遇到并缓解了这种情况，并且具有 CPU 饥渴的应用程序的虚拟机会占用所有资源，从而损害其他人的利益。

Kubernetes has tools to prevent this from happening, though they can be tricky to use and can lead to what Cheung called “fine-grained hotspots” that are very difficult to pinpoint. Kubernetes uses the Linux kernel’s [CFS Bandwidth Control](https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt), which allots CPU time in microseconds to pre-defined groups. This can lead to throttling issues: A node can look slow, even when there is not a lot else happening on that CPU. If you set a CFS quota of 100 milliseconds of processor time for an application that requests 10 CPUs, it can use all 10 CPUs and burn up its quota within 20 milliseconds and will be throttled for the remaining 80 milliseconds, spiking the legacy levels for that app (The Linux kernel [has subsequently addressed this issue](https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt) with a patch).

Kubernetes 有防止这种情况发生的工具，尽管它们使用起来可能很棘手，并且可能导致 Cheung 所说的难以确定的“细粒度热点”。 Kubernetes 使用 Linux 内核的 [CFS 带宽控制](https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt)，它将 CPU 时间以微秒为单位分配给预定义的组。这可能会导致节流问题：即使该 CPU 上没有发生很多其他事情，节点也可能看起来很慢。如果您为请求 10 个 CPU 的应用程序设置 100 毫秒的 CFS 处理器时间配额，它可以使用所有 10 个 CPU 并在 20 毫秒内耗尽其配额，并将在剩余的 80 毫秒内受到限制，从而提高传统级别应用程序（Linux 内核 [随后解决了这个问题](https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt)和补丁)。

“It’s hard not to take some performance hits from multitenancy. Before applications were running on their own dedicated boxes, but now they are sharing all their resources from other strange applications,” Cheung said.

“很难不从多租户中减少一些性能。以前，应用程序运行在自己的专用机器上，但现在它们正在共享来自其他奇怪应用程序的所有资源，”张说。

The Kubernetes community has developed some fixes, including the ability to [make CFS quota periods configurable](https://github.com/kubernetes/kubernetes/pull/63437). There is a pull request that [disables CPU quotas](https://github.com/kubernetes/kubernetes/pull/75682) for pods requiring guaranteed quality of service (which you'd think wouldn't be throttled [but they were ](https://github.com/kubernetes/kubernetes/issues/70585)).

Kubernetes 社区已经开发了一些修复程序，包括能够[使 CFS 配额期限可配置](https://github.com/kubernetes/kubernetes/pull/63437)。有一个拉取请求 [禁用 CPU 配额](https://github.com/kubernetes/kubernetes/pull/75682) 用于需要保证服务质量的 pod（您认为不会受到限制 [但他们](https://github.com/kubernetes/kubernetes/issues/70585))。

**Sponsor Note**

**赞助商备注**

Portworx is the leading provider of persistent storage for containers and is used in production by healthcare, global manufacturing and telecom members of the Fortune Global 500 and other great companies. Learn about Portworx solutions for Kubernetes storage, DCOS storage and more at portworx.com.

Portworx 是容器持久存储的领先供应商，被财富全球 500 强的医疗保健、全球制造和电信成员以及其他伟大公司用于生产。在 portworx.com 上了解用于 Kubernetes 存储、DCOS 存储等的 Portworx 解决方案。

Autoscaling can be another sneaky culprit of performance lag. At Airbnb, one service jumped from 600 pods to 1,000 due to heavy demand. To the scheduler, the  CPU utilization was fine, running at about 50% overall. At least one host, however, was crammed with 18 identical service pods by the scheduler. In other words, the overall CPU usage was fine, but some nodes were nonetheless starved from the lack of attention from the CPU.

自动缩放可能是性能滞后的另一个狡猾的罪魁祸首。在 Airbnb，由于需求旺盛，一项服务从 600 个吊舱跃升至 1000 个。对于调度程序，CPU 利用率很好，总体运行在 50% 左右。然而，至少一台主机被调度程序塞满了 18 个相同的服务 pod。换句话说，整体 CPU 使用率很好，但由于 CPU 缺乏关注，一些节点仍然处于饥饿状态。

The K8s scheduler has a set of rules about where to place images, based on a number of rules, such as spreading the workloads across as many nodes as possible, or abiding by a preferred or required affinity to a particular node. One rule, however, is that if an image was already downloaded to one minion node, the pods would more likely schedule to that node. In this case, one gigantic image was downloaded to one node, where all the other pods all piled up on that node.

K8s 调度程序有一组关于放置图像的位置的规则，基于许多规则，例如将工作负载分散到尽可能多的节点上，或者遵守对特定节点的首选或要求的亲和性。然而，一个规则是，如果图像已经下载到一个 minion 节点，则 pod 更有可能调度到该节点。在这种情况下，一个巨大的图像被下载到一个节点，所有其他 pod 都堆积在该节点上。

“The scheduler can work against you in some pathological cases,” Chan said. The company will look at ways at limiting the number of pods that could run per node.

“在某些病理情况下，调度程序可能会对您不利，”陈说。该公司将研究限制每个节点可以运行的 Pod 数量的方法。

## Write Once Run Anywhere

## 一次编写，随处运行

The applications and underlying dependencies can also hamper performance. Take Java, for instance, Cheung said. One Airbnb development team noticed that a Java application, jumped from a 30 millisecond response time to over 100 milliseconds in the 95th percentile latency (or taking place on 95% of the servers). This happened, however, only when interacting with a database through a driver. The unusual thing was the application worked fine before it ran on Kubernetes. The culprit turned out to be how the Java Virtual Machine (JVM) handled multi-CPU nodes. A single JVM within a pod on a 36-node cluster would see 36 CPUs. Great. But put two more JVMs, each on its own pod, on that node, and they will _all_ see 36 CPUs. Naturally, bottlenecks would ensue.

应用程序和底层依赖项也会影响性能。 Cheung 说，以 Java 为例。一个 Airbnb 开发团队注意到 Java 应用程序的响应时间从 30 毫秒跃升至第 95 个百分位延迟（或发生在 95% 的服务器上）超过 100 毫秒。但是，只有在通过驱动程序与数据库交互时才会发生这种情况。不寻常的是，该应用程序在运行在 Kubernetes 上之前运行良好。罪魁祸首原来是 Java 虚拟机 (JVM) 如何处理多 CPU 节点。一个 36 节点集群上的 pod 中的单个 JVM 将看到 36 个 CPU。伟大的。但是在该节点上再放置两个 JVM，每个 JVM 都在自己的 pod 上，它们将_all_看到 36 个 CPU。自然会出现瓶颈。

The team found that the problem with earlier versions of Java, which were [not aware of containers](https://bugs.openjdk.java.net/browse/JDK-8146115), Cheung noted. Java would auto-tune itself by how many CPUs it thought it had, and in container environments, this could adversely affect how thread pools were handled. 

Cheung 指出，该团队发现早期版本的 Java 存在问题，这些版本 [不知道容器](https://bugs.openjdk.java.net/browse/JDK-8146115)。 Java 会根据它认为拥有的 CPU 数量自动调整自身，而在容器环境中，这可能会对线程池的处理方式产生不利影响。

The issue has since been fixed in Java 8u191+, though the lesson here is one to keep in mind: “Languages and apps can have deeper dependencies on the underlying systems that they run on,” Cheung said. Another lesson learned was that it is useful to have a baseline to compare Kubernetes performance against that of running without K8s, as this application did.

此问题已在 Java 8u191+ 中得到解决，尽管这里要记住的教训是：“语言和应用程序可能对它们运行的底层系统有更深层次的依赖，”Cheung 说。另一个教训是，像这个应用程序一样，有一个基线来比较 Kubernetes 的性能与没有 K8s 的运行的性能是很有用的。

Cheung and Chan discussed some other potential trouble spots, such as load balance issues stemming from IPtables, and general slowness stemming from DNS (Domain Name Server) misconfigurations.

Cheung 和 Chan 讨论了其他一些潜在的问题点，例如源自 IPtables 的负载平衡问题，以及源自 DNS（域名服务器）错误配置的普遍缓慢。

Overall, Kubernetes is only one component of a complex cloud native stack, they reminded the audience, and so user performance may vary.

总体而言，Kubernetes 只是复杂云原生堆栈的一个组件，他们提醒观众，因此用户性能可能会有所不同。

“Set the expectation that small performance differences will happen,” Chan said.

“设定会发生微小性能差异的预期，”陈说。

[KubeCon+CloudNativeCon](https://www.cncf.io/kubecon-cloudnativecon-events/) is a sponsor of The New Stack. 

[KubeCon+CloudNativeCon](https://www.cncf.io/kubecon-cloudnativecon-events/) 是 The New Stack 的赞助商。

