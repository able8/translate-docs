# One year using Kubernetes in production: Lessons learned

# 在生产中使用 Kubernetes 一年：经验教训

[Paul Bakker](http://techbeacon.com/contributors/paul-bakker)
Software architect, Netflix

![A wheel of containers](http://techbeacon.scdn7.secure.raxcdn.com/sites/default/files/styles/article_hero_image/public/field/image/fineas-anton-141927.jpg?itok=QrFQCRoA)

In early 2015, after years of running deployments on Amazon EC2, my team at [Luminis Technologies](http://luminis-technologies.com/) was tasked with building a new deployment platform for all our development teams. The AWS-based setup had worked very well for [deploying new releases](https://content.microfocus.com/continuous-delivery-release-automation-tb/effective-product-release?lx=-DC2cJ&custom_url=continuous-delivery-release-automation-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOkQAM) over the years, but the deployment setup, with custom scripts and tooling to automate deployments, wasn't very easy for teams outside of operations to use—especially small teams that didn't have the resources to learn all of the details about these scripts and tools. The main issue was that there was no “unit-of-deployment,” and without one, there was a gap between development and operations. The [containerization trend](http://techbeacon.com/essential-guide-software-containers-application-development) was clearly going to change that.

2015 年初，在 Amazon EC2 上运行部署多年后，我在 [Luminis Technologies](http://luminis-technologies.com/) 的团队的任务是为我们所有的开发团队构建一个新的部署平台。基于 AWS 的设置非常适合 [部署新版本](https://content.microfocus.com/continuous-delivery-release-automation-tb/effective-product-release?lx=-DC2cJ&custom_url=continuous-delivery-release-automation-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOkQAM)多年来，但是部署设置，带有自定义脚本和工具来自动化部署，对于运营之外的团队来说并不是很容易使用——尤其是小团队没有资源来了解有关这些脚本和工具的所有详细信息。主要问题是没有“部署单元”，如果没有，则开发和运营之间存在差距。 [容器化趋势](http://techbeacon.com/essential-guide-software-containers-application-development) 显然会改变这种状况。

If you haven't bought in to the production readiness of Docker and Kubernetes yet, read about how my team became early adopters. We have now been running Kubernetes in production for over a year.

如果您还没有了解 Docker 和 Kubernetes 的生产准备情况，请阅读我的团队如何成为早期采用者。我们现在已经在生产中运行 Kubernetes 一年多了。

## Starting out with containers and container orchestration tools

## 从容器和容器编排工具开始

I now believe [containers are the deployment format of the future](http://techbeacon.com/state-containers-5-things-you-need-know-now). They make it much easier to package an application with its required infrastructure. While tools such as [Docker](http://techbeacon.com/docker-just-first-killer-app-container-revolution) provide the actual containers, we also need tools to take care of things such as replication and failovers, as well as APIs to automate deployments to multiple machines.

我现在相信[容器是未来的部署格式](http://techbeacon.com/state-containers-5-things-you-need-know-now)。它们使将应用程序与其所需的基础结构打包在一起变得更加容易。虽然诸如[Docker](http://techbeacon.com/docker-just-first-killer-app-container-revolution) 之类的工具提供了实际的容器，但我们还需要工具来处理诸如复制和故障转移之类的事情，以及用于自动部署到多台机器的 API。

The state of [clustering tools such as Kubernetes and Docker Swarm](http://techbeacon.com/scaling-containers-essential-guide-container-clusters) was very immature in early 2015, with only early alpha versions available. We still tried using them and started with Docker Swarm.

[Kubernetes 和 Docker Swarm 等集群工具](http://techbeacon.com/scaling-containers-essential-guide-container-clusters) 的状态在 2015 年初非常不成熟，只有早期的 alpha 版本可用。我们仍然尝试使用它们，并从 Docker Swarm 开始。

At first we used it to handle networking on our own with the ambassador pattern and a bunch of scripts to automate the deployments. How hard could it possibly be? That was our first hard lesson: **Container clustering, networking, and deployment automation are actually _very hard problems to solve_.**

起初，我们使用大使模式和一堆脚本来自动处理我们自己的网络部署。可能有多难？那是我们的第一堂课：**容器集群、网络和部署自动化实际上是_非常难解决的问题_。**

We realized this quickly enough and decided to bet on another one of the available tools. [Kubernetes](http://techbeacon.com/does-kubernetes-make-containers-ready-prime-time) seemed to be the best choice, since it was being backed by Google, Red Hat, Core OS, and other groups that clearly know about running large-scale deployments.

我们很快意识到这一点，并决定押注另一种可用工具。 [Kubernetes](http://techbeacon.com/does-kubernetes-make-containers-ready-prime-time) 似乎是最好的选择，因为它得到了 Google、Red Hat、Core OS 和其他团体的支持清楚地了解运行大规模部署。

## Load balancing with Kubernetes

## Kubernetes 负载均衡

When working with Kubernetes, you have to become familiar with concepts such as [pods](http://kubernetes.io/docs/user-guide/pods/), [services](http://kubernetes.io/docs/user-guide/services/), and [replication controllers](http://kubernetes.io/docs/user-guide/replication-controller/). If you're not already familiar with these concepts, there are some excellent resources available to get up to speed. The Kubernetes [documentation](http://kubernetes.io/docs/whatisk8s/) is a great place to start, since it has several guides for beginners.

在使用 Kubernetes 时，您必须熟悉诸如 [pods](http://kubernetes.io/docs/user-guide/pods/)、[services](http://kubernetes.io/docs/)等概念user-guide/services/) 和 [复制控制器](http://kubernetes.io/docs/user-guide/replication-controller/)。如果您还不熟悉这些概念，有一些优秀的资源可以帮助您快速上手。 Kubernetes [文档](http://kubernetes.io/docs/whatisk8s/) 是一个很好的起点，因为它为初学者提供了几个指南。

Once we had a Kubernetes cluster up and running, we could deploy an application using [kubectl](http://kubernetes.io/docs/user-guide/kubectl-overview/), the Kubernetes CLI, but we quickly found that kubectl wasn't sufficient when we wanted to automate deployments. But first, we had another problem to solve: How to access the deployed application from the Internet? 

一旦我们启动并运行了 Kubernetes 集群，我们就可以使用 Kubernetes CLI [kubectl](http://kubernetes.io/docs/user-guide/kubectl-overview/) 部署应用程序，但我们很快发现 kubectl当我们想要自动化部署时，这还不够。但首先，我们有另一个问题需要解决：如何从 Internet 访问已部署的应用程序？

The service in front of the deployment has an IP address, but this address only exists within the Kubernetes cluster. This means the service isn’t available to the Internet at all! When running on Google Cloud Engine, Kubernetes can automatically configure a load balancer to access the application. If you’re not on GCE (like us), you need to do a little extra legwork to get load balancing working.

部署前面的服务有一个IP地址，但是这个地址只存在于Kubernetes集群中。这意味着该服务根本无法在 Internet 上使用！在 Google Cloud Engine 上运行时，Kubernetes 可以自动配置负载均衡器来访问应用程序。如果你不在 GCE 上（像我们一样），你需要做一些额外的工作来让负载平衡工作。

It’s possible to expose a service directly on a host machine port—and this is how a lot of people get started—but we found that it voids a lot of Kubernetes' benefits. If we rely on ports in our host machines, we will get into port conflicts when deploying multiple applications. It also makes it much harder to scale the cluster or replace host machines.

可以直接在主机端口上公开服务——这就是很多人开始的方式——但我们发现它失去了 Kubernetes 的许多好处。如果我们依赖主机中的端口，我们将在部署多个应用程序时遇到端口冲突。它还使得扩展集群或更换主机变得更加困难。

### A two-step load-balancer setup

### 两步负载平衡器设置

We found that a much better approach is to configure a load balancer such as HAProxy or NGINX in front of the Kubernetes cluster. We started running our Kubernetes clusters inside a VPN on AWS and using an AWS Elastic Load Balancer to route external web traffic to an internal HAProxy cluster. HAProxy is configured with a “back end” for each Kubernetes service, which proxies traffic to individual pods.

我们发现更好的方法是在 Kubernetes 集群前配置负载均衡器，例如 HAProxy 或 NGINX。我们开始在 AWS 上的 VPN 内运行我们的 Kubernetes 集群，并使用 AWS Elastic Load Balancer 将外部 Web 流量路由到内部 HAProxy 集群。 HAProxy 为每个 Kubernetes 服务配置了一个“后端”，它将流量代理到各个 pod。

This two-step load-balancer setup is mostly in response AWS ELB's fairly limited configuration options. One of the limitations is that it can’t handle multiple vhosts. This is the reason we’re using HAProxy as well. Just using HAProxy (without an ELB) could also work, but you would have to work around dynamic AWS IP addresses on the DNS level.

这种两步式负载平衡器设置主要是为了响应 AWS ELB 相当有限的配置选项。限制之一是它无法处理多个虚拟主机。这也是我们使用 HAProxy 的原因。仅使用 HAProxy（没有 ELB）也可以工作，但您必须在 DNS 级别处理动态 AWS IP 地址。

![](http://techbeacon.com/sites/default/files/cb0_bxnyzdujb4p0ap7melynm59fdzguvyihqto8y75iwecdozx_vunkwkkn0u0hposnmhfu7ea4fpnmk0nlkklrdvya9chymspi6dyt_3safakd0wtf1bysfcz-8oc5f1vqo_1q.jpeg)

_Figure 1: Diagram of our two-step process for load balancing_

_图 1：我们的负载均衡两步流程示意图_

In any case, we needed a mechanism to dynamically reconfigure the load balancer (HAProxy, in our case) when new Kubernetes services are created.

在任何情况下，我们都需要一种机制来在创建新的 Kubernetes 服务时动态重新配置负载均衡器（在我们的例子中为 HAProxy）。

The Kubernetes community is currently working on a feature called [ingress](http://kubernetes.io/docs/user-guide/ingress/). It will make it possible to configure an external load balancer directly from Kubernetes. Currently, this feature isn’t really usable yet because it’s simply not finished. Last year, we used the API and a small open-source tool to configure load balancing instead.

Kubernetes 社区目前正在开发一项名为 [ingress](http://kubernetes.io/docs/user-guide/ingress/) 的功能。这将使直接从 Kubernetes 配置外部负载均衡器成为可能。目前，这个功能还不能真正使用，因为它还没有完成。去年，我们使用 API 和一个小型开源工具来配置负载平衡。

### Configuring load balancing

### 配置负载均衡

First, we needed a place to store load-balancer configurations. They could be stored anywhere, but because we already had [etcd](https://github.com/coreos/etcd) available, we decided to store the load-balancer configurations there. We use a tool called [confd](http://www.confd.io/) to watch configuration changes in etcd and generate a new HAProxy configuration file based on a template. When a new service is added to Kubernetes, we add a new configuration to etcd, which results in a new configuration file for HAProxy.

首先，我们需要一个地方来存储负载均衡器配置。它们可以存储在任何地方，但因为我们已经有 [etcd](https://github.com/coreos/etcd) 可用，我们决定将负载均衡器配置存储在那里。我们使用一个名为 [confd](http://www.confd.io/) 的工具来观察 etcd 中的配置变化，并基于模板生成一个新的 HAProxy 配置文件。当一个新的服务被添加到 Kubernetes 时，我们会向 etcd 添加一个新的配置，这会产生一个新的 HAProxy 配置文件。

### Kubernetes: Maturing the right way

### Kubernetes：以正确的方式成熟

There are still plenty of unsolved problems in Kubernetes, just as there are in load balancing generally. Many of these issues are recognized by the community, and there are design documents that discuss new features that can solve some of them. But coming up with solutions that work for everyone requires time, which means some of these features can take quite a while before they land in a release. **This is a good thing**, because it would be harmful in the long term to take shortcuts when designing new functionality.

Kubernetes 中还有很多未解决的问题，就像一般的负载均衡一样。其中许多问题已得到社区的认可，并且有一些设计文档讨论了可以解决其中一些问题的新功能。但是提出适用于每个人的解决方案需要时间，这意味着其中一些功能在发布之前可能需要很长时间。 **这是一件好事**，因为从长远来看，在设计新功能时走捷径是有害的。

This doesn’t mean Kubernetes is limited today. Using the API, it’s possible to make Kubernetes do pretty much everything you need it to if you want to start using it today. Once more features land in Kubernetes itself, we can replace custom solutions with standard ones.

这并不意味着 Kubernetes 在今天是有限的。如果您想从今天开始使用它，那么使用 API 可以让 Kubernetes 完成您需要它做的几乎所有事情。一旦 Kubernetes 本身具有更多功能，我们就可以用标准解决方案替换自定义解决方案。

After we developed our custom solution for load balancing, our next challenge was implementing an essential deployment technique for us: [Blue-green deployments](http://martinfowler.com/bliki/BlueGreenDeployment.html).

在我们开发了负载平衡的自定义解决方案之后，我们的下一个挑战是为我们实现一种基本的部署技术：[蓝绿部署](http://martinfowler.com/bliki/BlueGreenDeployment.html)。

## Blue-green deployments in Kubernetes 

## Kubernetes 中的蓝绿部署

A blue-green deployment is one without any downtime. In contrast to rolling updates, a blue-green deployment works by starting a cluster of replicas running the new version while all the old replicas are still serving all the live requests. Only when the new set of replicas is completely up and running is the load-balancer configuration changed to switch the load to the new version. A benefit of this approach is that there’s always only one version of the application running, reducing the complexity of handling multiple concurrent versions. Blue-green deployments also work better when the number of replicas is fairly small.

蓝绿部署是一种没有任何停机时间的部署。与滚动更新相比，蓝绿部署的工作原理是启动运行新版本的副本集群，同时所有旧副本仍为所有实时请求提供服务。只有当新的副本集完全启动并运行时，负载平衡器配置才会更改以将负载切换到新版本。这种方法的一个好处是始终只有一个版本的应用程序在运行，从而降低了处理多个并发版本的复杂性。当副本数量相当少时，蓝绿部署也能更好地工作。

![](http://techbeacon.com/sites/default/files/cgvferf0lrgrc4vst4myz0lnxezt4os7vtjroeek5icvxjrhtvnag-tf9wqjdutntq2mn1_igaixy1mgzswbodceyql0s6sid-e-omrjddembbfz5-con_ggxlj-5x7ltyqdo9fr.jpeg)

_Figure 2: Our blue-green deployments with Kubernetes_

_图 2：我们使用 Kubernetes 的蓝绿部署_

Figure 2 shows a component “Deployer” that orchestrates the deployment. This component can easily be created by your own team because we [open-sourced](https://bitbucket.org/amdatulabs/amdatu-kubernetes-deployer) our implementation under the Apache License as part of the [Amdatu](https://bitbucket.org/account/user/amdatulabs/projects/INFRA) umbrella project. It also comes with a web UI to configure deployments.

图 2 显示了一个用于编排部署的组件“Deployer”。该组件可以由您自己的团队轻松创建，因为我们 [开源](https://bitbucket.org/amdatulabs/amdatu-kubernetes-deployer) 作为 [Amdatu] (https: //bitbucket.org/account/user/amdatulabs/projects/INFRA) 伞形项目。它还带有一个用于配置部署的网络用户界面。

An important aspect of this mechanism is the health checking it performs on the pods before reconfiguring the load balancer. We wanted each component that was deployed to provide a health check. Now we typically add a health check that's available on HTTP to each application component.

这种机制的一个重要方面是它在重新配置负载均衡器之前对 Pod 执行的健康检查。我们希望部署的每个组件都提供健康检查。现在，我们通常会向每个应用程序组件添加 HTTP 上可用的健康检查。

## Making the deployments automatic

## 使部署自动化

With the Deployer in place, we were able to hook up deployments to a build pipeline. Our build server can, after a successful build, push a new Docker image to a registry such as Docker Hub. Then the build server can invoke the Deployer to automatically deploy the new version to a test environment. The same image can be promoted to production by triggering the Deployer on the production environment.

部署器到位后，我们能够将部署连接到构建管道。我们的构建服务器可以在成功构建后，将新的 Docker 镜像推送到诸如 Docker Hub 之类的注册中心。然后构建服务器可以调用 Deployer 自动将新版本部署到测试环境。可以通过在生产环境中触发 Deployer 将相同的映像提升到生产环境。

![](http://techbeacon.com/sites/default/files/wrhr6wcfhbbn13ggjyxbmjybw8jwsifh9arcfysryisfvllzrkdm4kvubzhxfspc7vlejehm-ub4q_0mpylovpjc0gyu2tgn5wbpjizkclq8cz602uhemgincy-niphtffu3dezp.jpeg)

_Figure 3: Our automated container deployment pipeline_

_图 3：我们的自动化容器部署管道_

## Know your resource constraints

## 了解您的资源限制

Knowing our [resource constraints](http://kubernetes.io/docs/user-guide/compute-resources/) was critical when we started using Kubernetes. You can configure resource requests and CPU/memory limits on each pod. You can also control resource guarantees and bursting limits.

当我们开始使用 Kubernetes 时，了解我们的 [资源限制](http://kubernetes.io/docs/user-guide/compute-resources/) 至关重要。您可以在每个 Pod 上配置资源请求和 CPU/内存限制。您还可以控制资源保证和突发限制。

These settings are extremely important for running multiple containers together efficiently. If we didn't set these settings correctly, containers would often crash because they couldn't allocate enough memory.

这些设置对于高效地一起运行多个容器非常重要。如果我们没有正确设置这些设置，容器经常会因为无法分配足够的内存而崩溃。

**Start early with setting and testing constraints**. Without constraints, everything will still run fine, but you'll get a big, unpleasant surprise when you put any serious load on one of the containers.

**尽早开始设置和测试约束**。没有约束，一切仍然会正常运行，但是当您在其中一个容器上放置任何严重负载时，您会得到一个很大的、令人不快的惊喜。

## How we monitored Kubernetes

## 我们如何监控 Kubernetes

When we had Kubernetes mostly set up, we quickly realized that monitoring and logging would be crucial in this new dynamic environment. Logging into a server to look a log files just doesn’t work anymore when you're dealing with a large number of replicas and nodes. As soon as you start using Kubernetes, you should also have a plan to build centralized logging and monitoring.

当我们基本设置 Kubernetes 时，我们很快意识到监控和日志记录在这个新的动态环境中至关重要。当您处理大量副本和节点时，登录服务器查看日志文件不再有效。一旦您开始使用 Kubernetes，您还应该有一个计划来构建集中式日志记录和监控。

### Logging 

### 日志记录

There are plenty of open-source tools available for logging. We decided to use [Graylog](https://www.graylog.org/)—an excellent tool for logging—and [Apache Kafka, a messaging system](http://techbeacon.com/what-apache-kafka-why-it-so-popular-should-you-use-it) to collect and digest logs from our containers. The containers send logs to Kafka, and Kafka hands them off to Graylog for indexing. We chose to make the application components send logs to Kafka themselves so that we could stream logs in an easy-to-index format. [Alternatively](https://www.loggly.com/blog/top-5-docker-logging-methods-to-fit-your-container-deployment-strategy/), there are tools that retrieve logs from outside the container and forward them to a logging solution.

有很多开源工具可用于日志记录。我们决定使用[Graylog](https://www.graylog.org/)——一个优秀的日志工具——和[ApacheKafka，一个消息系统](http://techbeacon.com/what-apache-kafka-为什么它如此受欢迎，你应该使用它)从我们的容器中收集和消化日志。容器将日志发送到 Kafka，然后 Kafka 将它们交给 Graylog 进行索引。我们选择让应用程序组件自己将日志发送到 Kafka，以便我们可以以易于索引的格式流式传输日志。 [或者](https://www.loggly.com/blog/top-5-docker-logging-methods-to-fit-your-container-deployment-strategy/)，有一些工具可以从容器外检索日志并将它们转发到日志记录解决方案。

### Monitoring

### 监控

Kubernetes does an excellent job of recovering when there's an error. When pods crash for any reason, Kubernetes will restart them. When Kubernetes is running replicated, end users probably won't even notice a problem. **Kubernetes recovery works so well that we have had situations where our containers would crash multiple times a day because of a memory leak, without anyone (including ourselves) noticing it.**

当出现错误时，Kubernetes 在恢复方面做得非常出色。当 Pod 因任何原因崩溃时，Kubernetes 将重新启动它们。当 Kubernetes 运行复制时，最终用户可能甚至不会注意到问题。 **Kubernetes 恢复效果非常好，以至于我们的容器一天会因为内存泄漏而崩溃多次，而没有人（包括我们自己）注意到这一点。**

Although this is great from the perspective of Kubernetes, you probably still want to know whenever there’s a problem. We use a custom health-check dashboard that monitors the Kubernetes nodes, individual pods—using application-specific health checks—and other services such as data stores. To implement a dashboard such as this, the Kubernetes API proves to be extremely valuable again.

尽管从 Kubernetes 的角度来看这很好，但您可能仍然想知道何时出现问题。我们使用自定义运行状况检查仪表板来监控 Kubernetes 节点、单个 pod（使用特定于应用程序的运行状况检查）以及其他服务（例如数据存储）。为了实现这样的仪表板，Kubernetes API 再次被证明是非常有价值的。

We also thought it was important to measure load, throughput, application errors, and other stats. Again, the open-source space has a lot to offer. Our application components post metrics to an [InfluxDB](https://influxdata.com/) time-series store. We also use [Heapster](https://github.com/kubernetes/heapster) to gather Kubernetes metrics. The metrics stored in InfluxDB are visualized in [Grafana](http://grafana.org/), an open-source dashboard tool. There are a lot of alternatives to the InfluxDB/Grafana stack, and any one of them will provide a lot of value toward keeping track of how things are running.

我们还认为衡量负载、吞吐量、应用程序错误和其他统计数据很重要。同样，开源空间有很多东西可以提供。我们的应用程序组件将指标发布到 [InfluxDB](https://influxdata.com/) 时间序列存储。我们还使用 [Heapster](https://github.com/kubernetes/heapster) 来收集 Kubernetes 指标。 InfluxDB 中存储的指标在开源仪表板工具 [Grafana](http://grafana.org/) 中可视化。 InfluxDB/Grafana 堆栈有很多替代方案，其中任何一个都将为跟踪事物的运行方式提供很多价值。

## Data stores and Kubernetes

## 数据存储和 Kubernetes

A question that many new Kubernetes users ask is “How should I handle my data stores with Kubernetes?”

许多 Kubernetes 新用户问的一个问题是“我应该如何使用 Kubernetes 处理我的数据存储？”

When running a data store such as MongoDB or MySQL, you most likely want the data to be persistent. Out of the box, containers lose their data when they restart. This is fine for stateless components, but not for a persistent data store. Kubernetes has the concept of [volumes](http://kubernetes.io/docs/user-guide/volumes/) to work with persistent data.

在运行 MongoDB 或 MySQL 等数据存储时，您很可能希望数据是持久的。开箱即用，容器在重新启动时会丢失数据。这适用于无状态组件，但不适用于持久数据存储。 Kubernetes 有 [volumes](http://kubernetes.io/docs/user-guide/volumes/) 的概念来处理持久数据。

A volume can be backed by a variety of implementations, including files on the host machines, AWS Elastic Block Store (EBS), and [nfs](https://en.wikipedia.org/wiki/Network_File_System). When we were researching the question of persistent data, this provided a good answer, but it wasn't an answer for our running data stores yet.

卷可以由多种实现支持，包括主机上的文件、AWS Elastic Block Store (EBS) 和 [nfs](https://en.wikipedia.org/wiki/Network_File_System)。当我们研究持久化数据的问题时，这提供了一个很好的答案，但它还不是我们正在运行的数据存储的答案。

### Replication issues

### 复制问题

In most deployments, the data stores also run replicated. Mongo typically runs in a Replica Set, and MySQL could be running in primary/replica mode. This introduces a few problems. First of all, it’s important that each node in the data store’s cluster is backed by a different volume. Writing to the same volume will lead to data corruption. Another issue is that most data stores require precise configuration to get the clustering up and running; auto discovery and configuration of nodes is not common.

在大多数部署中，数据存储也以复制方式运行。 Mongo 通常在副本集中运行，而 MySQL 可以在主/副本模式下运行。这引入了一些问题。首先，重要的是数据存储集群中的每个节点都由不同的卷支持。写入同一个卷会导致数据损坏。另一个问题是大多数数据存储需要精确配置才能启动和运行集群；节点的自动发现和配置并不常见。

At the same time, a machine that runs a data store is often specifically tuned for that type of workload. Higher IOPS could be one example. Scaling (adding/removing nodes) is an expensive operation for most data stores as well. All these things don’t match very well with the dynamic nature of Kubernetes deployments.

同时，运行数据存储的机器通常专门针对该类型的工作负载进行调整。更高的 IOPS 就是一个例子。缩放（添加/删除节点）对于大多数数据存储也是一项昂贵的操作。所有这些都与 Kubernetes 部署的动态特性不太匹配。

### The decision not to use Kubernetes for running data stores in production

### 决定不使用 Kubernetes 在生产中运行数据存储

This brings us to a situation where we found that the benefits of running a data store inside Kubernetes are limited. The dynamics that Kubernetes give us can’t really be used. The setup is also much more complex than most Kubernetes deployments. 

这让我们发现在 Kubernetes 中运行数据存储的好处是有限的。 Kubernetes 为我们提供的动态无法真正使用。设置也比大多数 Kubernetes 部署复杂得多。

Because of this, **we are not running our production data stores inside Kubernetes**. Instead, we set up these clusters manually on different hosts, with all the tuning necessary to optimize the data store in question. Our applications running inside Kubernetes just connect to the data store cluster like normal. The important lesson is that **you don’t have to run everything in Kubernetes once you have Kubernetes**. Besides data stores and our HAProxy servers, everything else does run in Kubernetes, though, including our monitoring and logging solutions.

因此，**我们不在 Kubernetes 内运行我们的生产数据存储**。相反，我们在不同的主机上手动设置这些集群，并进行优化相关数据存储所需的所有调整。我们在 Kubernetes 中运行的应用程序只是像往常一样连接到数据存储集群。重要的一课是**一旦拥有 Kubernetes，您就不必在 Kubernetes 中运行所有内容**。不过，除了数据存储和我们的 HAProxy 服务器之外，其他一切都在 Kubernetes 中运行，包括我们的监控和日志记录解决方案。

## Why we're excited about our next year with Kubernetes

## 为什么我们对 Kubernetes 的明年感到兴奋

Looking at our deployments today, Kubernetes is absolutely fantastic. The Kubernetes API is a great tool when it comes to automating a deployment pipeline. Deployments are not only more reliable, but also much faster, because we’re no longer dealing with VMs. Our builds and deployments have become more reliable because it’s easier to test and ship containers.

看看我们今天的部署，Kubernetes 绝对是太棒了。在自动化部署管道方面，Kubernetes API 是一个很好的工具。部署不仅更可靠，而且速度更快，因为我们不再处理虚拟机。我们的构建和部署变得更加可靠，因为它更容易测试和运输容器。

We see now that this new way of deployment was necessary to keep up with other development teams around the industry that are pushing out deployments much more often and lowering their overhead for doing so.

我们现在看到，这种新的部署方式对于跟上行业中其他开发团队的步伐是必要的，这些团队正在更频繁地推出部署并降低他们这样做的开销。

### Cost calculation

### 成本计算

Looking at costs, there are two sides to the story. To run Kubernetes, an etcd cluster is required, as well as a master node. While these are not necessarily expensive components to run, this overhead can be relatively expensive when it comes to very small deployments. For these types of deployments, it’s probably best to use a hosted solution such as Google's Container Service.

从成本来看，故事有两个方面。要运行 Kubernetes，需要一个 etcd 集群以及一个主节点。虽然这些组件运行起来不一定很昂贵，但对于非常小的部署来说，这种开销可能相对昂贵。对于这些类型的部署，最好使用托管解决方案，例如 Google 的容器服务。

For larger deployments, it’s easy to save a lot on server costs. The overhead of running etcd and a master node aren’t significant in these deployments. Kubernetes makes it very easy to run many containers on the same hosts, making maximum use of the available resources. This reduces the number of required servers, which directly saves you money. When running Kubernetes sounds great, but the ops side of running such a cluster seems less attractive, there are a number of hosted services to look at, including Cloud RTI, which is what my team is working on.

对于大型部署，可以轻松节省大量服务器成本。在这些部署中，运行 etcd 和主节点的开销并不大。 Kubernetes 使在同一主机上运行多个容器变得非常容易，从而最大限度地利用了可用资源。这减少了所需服务器的数量，从而直接为您节省了资金。运行 Kubernetes 听起来不错，但运行这样一个集群的运维方面似乎不那么有吸引力，有许多托管服务可供查看，包括 Cloud RTI，这是我的团队正在研究的内容。

## A bright future for Kubernetes

## Kubernetes 的美好未来

Running Kubernetes in a pre-released version was challenging, and keeping up with (breaking) new releases was almost impossible at times. Development of Kubernetes has been happening at light-speed in the past year, and the community has grown into a legitimate powerhouse of dev talent. It’s hard to believe how much progress has been made in just over a year.

在预发布版本中运行 Kubernetes 具有挑战性，有时几乎不可能跟上（突破）新版本。在过去的一年里，Kubernetes 的发展一直在以光速进行，社区已经成长为一个合法的开发人才强国。很难相信在短短一年多的时间里取得了如此大的进步。

#### Keep learning

####  保持学习

- **Choose the right ESM tool** for your needs. Get up to speed with the our [Buyer's Guide to Enterprise Service Management Tools](https://content.microfocus.com/l/buyers-guide-esm-tools-special-report-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOTQA2)

- **根据您的需要选择合适的 ESM 工具**。快速了解我们的 [企业服务管理工具购买者指南](https://content.microfocus.com/l/buyers-guide-esm-tools-special-report-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOTQA2)

- **What will the next generation of enterprise service management tools look like?** TechBeacon's [Guide to Optimizing Enterprise Service Management](https://content.microfocus.com/l/enterprise-service-management-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOTQA2) offers the insights.

- **下一代企业服务管理工具会是什么样子？** TechBeacon 的 [优化企业服务管理指南](https://content.microfocus.com/l/enterprise-service-management-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOTQA2) 提供了见解。

- **Discover more** about [IT Operations Monitoring with TechBeacon's Guide](https://content.microfocus.com/l/it-operations-monitoring-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOTQA2).

- **发现更多**关于[使用 TechBeacon 指南进行 IT 运营监控](https://content.microfocus.com/l/it-operations-monitoring-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOTQA2)。

- What's **the best way to get your robotic process automation project off the ground?** Find out how to [choose the right tools—and the right project.](https://content.microfocus.com/l/robotic-process-automation-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOTQA2)

- **让您的机器人流程自动化项目起步的最佳方式是什么？** 了解如何[选择正确的工具和正确的项目。](https://content.microfocus.com/l/robotic-process-automation-tb?utm_source=techbeacon&utm_medium=referral&utm_campaign=7014J000000dVOTQA2)

- **Ready to advance up the IT career ladder?** [TechBeacon's Careers Topic Center](https://content.microfocus.com/careers-tb-tc?utm_campaign=7014J000000dVOTQA2) provides expert advice you need to prepare for your next move.

- **准备好提升 IT 职业阶梯了吗？** [TechBeacon 的职业主题中心](https://content.microfocus.com/careers-tb-tc?utm_campaign=7014J000000dVOTQA2) 提供您需要准备的专家建议下一步行动。

Read more articles about: [Enterprise IT](http://techbeacon.com/enterprise-it), [ITOps](https://techbeacon.com/categories/it-ops) 

阅读有关以下内容的更多文章：[企业 IT](http://techbeacon.com/enterprise-it)、[ITOps](https://techbeacon.com/categories/it-ops)

