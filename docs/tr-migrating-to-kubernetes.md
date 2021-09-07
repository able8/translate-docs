# Migrating to Kubernetes

# 迁移到 Kubernetes

March 30, 2020

2020 年 3 月 30 日

The reasons to move to Kubernetes are many and compelling. This post doesn’t make the case that you should migrate, but assumes you have already decided that you want to. When you’re clear on what you want to do and why you want to do it, the questions of “When?” and “How?” become your focus. What follows centers on the question of how to approach making Kubernetes the platform on which your workloads thrive.

迁移到 Kubernetes 的原因有很多而且非常引人注目。这篇文章并没有说明您应该迁移，而是假设您已经决定要迁移。当您清楚自己想做什么以及为什么要这样做时，就会出现“何时？”的问题。如何？”成为你的焦点。接下来的重点是如何让 Kubernetes 成为您的工作负载蓬勃发展的平台。

![Migrating to Kubernetes, captained by Sensu mascot Lizy](http://images.ctfassets.net/w1bd7cq683kz/uaexd5dpIUsM56gpKpOZp/3ef2131b76d60b37bc32ef2664e5d1ca/Picture1-2.png)

How to migrate depends, to an extent, on what you want to migrate from. The primary consideration is whether your existing infrastructure runs workloads in containers. If so, you’re already off to a quick start because you won’t have the containerization step to complete. Otherwise, you have a clear place to start.

如何迁移在一定程度上取决于您要迁移的对象。主要考虑因素是您现有的基础设施是否在容器中运行工作负载。如果是这样，您就可以快速开始了，因为您无需完成容器化步骤。否则，您将有一个明确的起点。

Further, your considerations will vary if you are running via a serverless solution or other cloud platforms as services. These platforms provide value by making some decisions for you. In some cases and some ways, your solution is specific to the platform. Migrating from infrastructure as a service or your on-premises infrastructure has its own set of concerns. Moving from a different container orchestration engine to Kubernetes is another possibility with some unique characteristics.

此外，如果您通过无服务器解决方案或其他云平台作为服务运行，您的考虑会有所不同。这些平台通过为您做出一些决定来提供价值。在某些情况下，您的解决方案特定于平台。从基础设施即服务或本地基础设施迁移有其自身的一系列问题。从不同的容器编排引擎迁移到 Kubernetes 是另一种具有一些独特特征的可能性。

Although where you’re coming from will account for some differences in how you proceed, there are common concerns you’ll need to consider in your journey. This post touches briefly on differences in migrating from differing platforms and focuses mostly on the decisions you’ll need to make in any case.

尽管您来自哪里会导致您的行进方式存在一些差异，但您在旅途中需要考虑一些常见问题。这篇文章简要介绍了从不同平台迁移的差异，并主要关注您在任何情况下都需要做出的决定。

## Kubernetes: What and why

## Kubernetes：是什么以及为什么

Kubernetes is the open source container-orchestration system for automating the deployment, scaling, and management of containerized applications. While we won't go into the details of how Kubernetes works here, you can learn more by checking out these posts from the Sensu team: [Kubernetes 101](https://sensu.io/blog/kubernetes-101) and [ How Kubernetes works](https://sensu.io/blog/how-kubernetes-works). We’ll assume from here on out that you’ve decided you’re ready to migrate (or at the very least, you’re ready to start reasoning about it).

Kubernetes 是一个开源容器编排系统，用于自动化容器化应用程序的部署、扩展和管理。虽然我们不会在这里详细介绍 Kubernetes 的工作原理，但您可以通过查看 Sensu 团队的这些帖子了解更多信息：[Kubernetes 101](https://sensu.io/blog/kubernetes-101) 和 [ Kubernetes 的工作原理](https://sensu.io/blog/how-kubernetes-works)。从现在开始，我们将假设您已决定已准备好迁移（或者至少，您已准备好开始考虑迁移)。

Your approach to using Kubernetes depends on your knowledge of what it provides. It also depends on why you want to use it. The ideas and opinions in this post are general guidelines. They don’t substitute for your judgment in making sure your decisions serve your purposes (you know your applications better than we do, after all!). We suggest you think carefully about migration in your own context.

您使用 Kubernetes 的方法取决于您对它提供的内容的了解。这也取决于您为什么要使用它。这篇文章中的想法和意见是一般准则。在确保您的决定符合您的目的时，它们并不能代替您的判断（毕竟，您比我们更了解您的应用程序！）。我们建议您根据自己的情况仔细考虑迁移。

## Table stakes: Containerization

## 赌注：容器化

If your application workloads aren’t already running in containers, step 0 in your migration will be to change that, and you’ll likely want to use Docker. Docker isn’t the only way to containerize, but it’s usually the obvious choice given that it’s intuitive, well-supported, and in broad use.

如果您的应用程序工作负载尚未在容器中运行，则迁移中的第 0 步将改变这一点，您可能希望使用 Docker。 Docker 不是容器化的唯一方法，但鉴于它直观、支持良好且用途广泛，它通常是显而易见的选择。

You use Docker to create images that support your applications with all of their dependencies. This is best done using an automated [continuous integration](https://www.atlassian.com/continuous-delivery/continuous-integration) pipeline that includes pushing the versioned images to a [Docker registry](https://docs.docker.com/registry/introduction/) (which can be private). With these images in a registry, you’re ready to move into using Kubernetes.

您可以使用 Docker 创建支持您的应用程序及其所有依赖项的映像。这最好使用自动化的 [持续集成](https://www.atlassian.com/continuous-delivery/continuous-integration) 管道来完成，该管道包括将版本化图像推送到 [Docker 注册表](https://docs.docker.com/registry/introduction/)（可以是私有的)。使用注册表中的这些映像，您就可以开始使用 Kubernetes。

Because this post presents Kubernetes migration rather than Docker fundamentals, there won’t be more here about this step. There are many other [great resources for getting familiar with Docker](https://docs.docker.com/get-started/).

因为这篇文章介绍的是 Kubernetes 迁移而不是 Docker 基础知识，所以这里不会有更多关于这一步的内容。还有许多其他 [用于熟悉 Docker 的重要资源](https://docs.docker.com/get-started/)。

## There and back again: A journey to Kubernetes

## 来来回回：Kubernetes 之旅

With the pieces of your application running in containers, it’s time to think about your strategy for migration.

随着应用程序的各个部分在容器中运行，是时候考虑迁移策略了。

Your primary concerns in architecting your application for Kubernetes execution are redundancy, resiliency, services, networking, monitoring, and persistent state.

在为 Kubernetes 执行构建应用程序时，您主要关心的是冗余、弹性、服务、网络、监控和持久状态。

### Redundancy 

### 冗余

Before migrating containerized application workloads onto your Kubernetes cluster, you need to have a cluster. Questions to answer include where to run this cluster and capacity planning for nodes in the cluster, including how many and how powerful.

在将容器化应用程序工作负载迁移到 Kubernetes 集群之前，您需要有一个集群。要回答的问题包括在何处运行此集群以及集群中节点的容量规划，包括数量和功能。

These are infrastructure questions.

这些是基础设施问题。

Using a cloud provider with a Kubernetes offering is usually the best choice on the where-to-run question. Which cloud provider and capacity planning on nodes are topics for other posts.

将云提供商与 Kubernetes 产品一起使用通常是解决在哪里运行问题的最佳选择。哪些云提供商和节点上的容量规划是其他帖子的主题。

Beyond the infrastructure level, you’ll need to plan for redundancy at the application level. How many replicas do you need for a given pod? You will answer such questions by telling Kubernetes how to run your pods via deployments and/or StatefulSets.

除了基础设施级别之外，您还需要在应用程序级别规划冗余。一个给定的 pod 需要多少个副本？您将通过告诉 Kubernetes 如何通过部署和/或 StatefulSet 运行您的 pod 来回答这些问题。

### Resiliency

### 弹性

Kubernetes takes care of monitoring the health of pods to make sure they are in operating condition. There are characteristics of containers that can indicate unhealthy pods, but Kubernetes needs your help knowing if your application is healthy in a container. To enable Kubernetes to take down pods that aren’t performing, replace them with new ones, and know when the new ones are ready, you need to contribute. You tell it how to know the state of your pods with [liveness, readiness, and startup probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).

Kubernetes 负责监控 Pod 的健康状况，以确保它们处于运行状态。容器的一些特征可以指示不健康的 pod，但 Kubernetes 需要您的帮助，了解您的应用程序在容器中是否健康。为了让 Kubernetes 能够删除不执行的 pod，用新的替换它们，并知道新的何时准备就绪，您需要做出贡献。你告诉它如何使用 [liveness, readiness, and startup probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)。

### Networking

###  联网

Having put your processes into containers doesn’t isolate them. You still have system components that need to communicate with other system components and with the outside world.

将您的流程放入容器并不会隔离它们。您仍然拥有需要与其他系统组件和外部世界进行通信的系统组件。

Ideally, you want most of the communication within your cluster to be either between containers within a pod or via [services](https://kubernetes.io/docs/concepts/services-networking/service/) for communications crossing pods.

理想情况下，您希望集群内的大部分通信在 pod 内的容器之间进行，或者通过 [services](https://kubernetes.io/docs/concepts/services-networking/service/) 进行跨 pod 的通信。

You have [several options for exposing your services to the outside world](https://medium.com/google-cloud/kubernetes-nodeport-vs-loadbalancer-vs-ingress-when-should-i-use-what-922f010849e0) and these can be confusing. With Kubernetes services, you specify a service type. Some of the types are directly exposed to the outside world. The LoadBalancer service type makes use of the hosting platform to set up a load balancer directly exposing your service. For example, LoadBalancer services in clusters in Amazon Web Services use [Elastic Load Balancers](https://aws.amazon.com/elasticloadbalancing/).

你有[向外界公开你的服务的几种选择](https://medium.com/google-cloud/kubernetes-nodeport-vs-loadbalancer-vs-ingress-when-should-i-use-what-922f010849e0) 而这些可能会令人困惑。使用 Kubernetes 服务，您可以指定服务类型。其中一些类型直接暴露于外部世界。 LoadBalancer 服务类型利用托管平台来设置直接公开您的服务的负载均衡器。例如，Amazon Web Services 中集群中的 LoadBalancer 服务使用 [Elastic Load Balancers](https://aws.amazon.com/elasticloadbalancing/)。

Instead of exposing a service directly by using a type, [ingress resources](https://kubernetes.io/docs/concepts/services-networking/ingress/) can be set up to control access to services or other resources. Ingress gives you more control over how you expose your application but can require more thought and add complexity.

可以设置 [ingress resources](https://kubernetes.io/docs/concepts/services-networking/ingress/) 来控制对服务或其他资源的访问，而不是通过使用类型直接暴露服务。 Ingress 使您可以更好地控制如何公开您的应用程序，但可能需要更多思考并增加复杂性。

### Monitoring

### 监控

Some operators, even before moving to Kubernetes, have embraced knowing that there are multiple system components with many replicas and that traditional monitoring of individual pieces isn’t going to cut it. There are new challenges with bringing new instances online and offline frequently (not to mention the fact that those instances are distributed). For a deeper dive into those challenges as well as popular patterns for [monitoring Kubernetes, check out this series from Sensu CTO Sean Porter](https://sensu.io/resources/whitepaper/integrations-for-cloud-monitoring-and-alerting#data-sources).

一些运营商，甚至在转向 Kubernetes 之前，就已经接受了，知道有多个系统组件和许多副本，并且传统的单个部分监控不会削减它。频繁地使新实例上线和下线会带来新的挑战（更不用说这些实例是分布式的了）。要更深入地了解这些挑战以及[监控 Kubernetes 的流行模式，请查看 Sensu 首席技术官 Sean Porter 的这个系列](https://sensu.io/resources/whitepaper/integrations-for-cloud-monitoring-and-警报#数据源)。

### Persistent state

### 持久状态

Migrating databases can be the most challenging part of moving systems. The first and most obvious hurdle is the moving of data. Large databases can be hard to migrate. More fundamental, though, is deciding where the data should go. 

迁移数据库可能是移动系统中最具挑战性的部分。第一个也是最明显的障碍是数据的移动。大型数据库可能难以迁移。不过，更重要的是决定数据应该去哪里。

Running your application in Kubernetes doesn’t necessarily mean your databases have to live in Kubernetes. The major cloud providers have offerings for running your database as a managed cloud service. Generally, that’s a pretty good way to offload much of the work required in maintaining databases and works well with pods in Kubernetes simply connecting to managed databases. Cloud providers offer custom cloud-native databases, but also services for running managed instances of well-known databases like SQL Server, Oracle, MySQL, PostgreSQL, and more. The most straightforward way to migrate databases is usually to move to something you’re already using, potentially just in a different place.

在 Kubernetes 中运行您的应用程序并不一定意味着您的数据库必须存在于 Kubernetes 中。主要的云提供商提供将数据库作为托管云服务运行的产品。通常，这是一种很好的方式来卸载维护数据库所需的大部分工作，并且只需连接到托管数据库即可与 Kubernetes 中的 Pod 配合使用。云提供商提供自定义云原生数据库，但也提供用于运行 SQL Server、Oracle、MySQL、PostgreSQL 等知名数据库的托管实例的服务。迁移数据库最直接的方法通常是迁移到您已经在使用的地方，可能只是在不同的地方。

Alternatively, virtual private networks (VPNs) can be used to connect to on-premises databases if you want to avoid migrating data altogether.

或者，如果您想完全避免迁移数据，可以使用虚拟专用网络 (VPN) 连接到本地数据库。

Databases in Kubernetes are also an option, though. This gives you more control in exchange for you needing to manage your persistent storage and database software containers in your cluster. Kubernetes does support attaching persistent storage volumes to containers.

不过，Kubernetes 中的数据库也是一种选择。这为您提供了更多控制权，以换取您需要管理集群中的持久存储和数据库软件容器。 Kubernetes 确实支持将持久存储卷附加到容器。

## Migrating from self-controlled infrastructure

## 从自控基础设施迁移

Deploying your systems on-premises is certainly not the same thing as running in a public cloud via infrastructure as a service. For the purposes of this post, though, consider them as close enough to the same to treat them together. With infrastructure as a service, you execute in the cloud with maximum control over how and where your workloads run. You have control over machines. Configuring, maintaining, and supporting physical hardware on-premises is obviously different, but that’s not a focus of an application-centric view.

在本地部署系统肯定不同于通过基础设施即服务在公共云中运行。但是，出于本文的目的，将它们视为足够接近相同的对象以将它们放在一起。借助基础设施即服务，您可以在云中执行，最大程度地控制工作负载的运行方式和位置。您可以控制机器。在本地配置、维护和支持物理硬件显然不同，但这不是以应用程序为中心的视图的重点。

### Toward immutable deployment units

### 面向不可变部署单元

If you have full control over your operating environments, you have the liberty to customize hardware (virtual or otherwise) to suit your needs. This enables useful scenarios and can help in troubleshooting but comes with great peril. Servers tweaked to get things running may have undocumented configurations. This makes the setup hard to repeat and servers hard to replace. Applications and support can behave nondeterministically. When you move to Kubernetes (or any platform as a service), you need to get accustomed to immutable runtime environments.

如果您可以完全控制您的操作环境，您可以自由定制硬件（虚拟或其他）以满足您的需求。这可以实现有用的场景并有助于故障排除，但会带来很大的危险。为使事情运行而进行调整的服务器可能具有未记录的配置。这使得设置难以重复且服务器难以更换。应用程序和支持的行为可能具有不确定性。当您转向 Kubernetes（或任何平台即服务）时，您需要习惯不可变的运行时环境。

Technically, you can mutate the state of running containers in Kubernetes pods, but it’s generally a bad idea. This is because Kubernetes replaces pods to maintain the desired state such that these changes in a container will be lost. Thus, treating deployed containers as immutable is a good idea. Further, enforcing policies via Kubernetes to make filesystems in containers read-only is supported and generally advisable.

从技术上讲，您可以改变 Kubernetes pod 中运行容器的状态，但这通常是一个坏主意。这是因为 Kubernetes 替换了 pod 以维持所需的状态，这样容器中的这些更改就会丢失。因此，将部署的容器视为不可变是一个好主意。此外，支持通过 Kubernetes 执行策略以使容器中的文件系统只读，并且通常是可取的。

This means teams need to shift their view of environment infrastructure as something set up by an operations role to something constructed from repeatable scripts and manifests.

这意味着团队需要将他们对环境基础设施的看法转变为由运营角色设置的事物，转变为由可重复的脚本和清单构建的事物。

### Planning the migration

### 规划迁移

Chances are if you’re running on-premises currently, you’re going to be containerizing first before moving anywhere. While containerizing, you’ll need to think through the application dependencies and how to put together images that serve your applications. At the same time, you’ll plan out manifests for how to get the containers to work together in a Kubernetes cluster.

如果您目前在本地运行，您可能会在移动到任何地方之前先进行容器化。在容器化时，您需要仔细考虑应用程序的依赖关系以及如何将服务于您的应用程序的图像放在一起。同时，您将规划如何让容器在 Kubernetes 集群中协同工作的清单。

Database migration is most challenging when moving from on-premises. Cloud providers do have tools that can help in moving the information.

从本地迁移时，数据库迁移最具挑战性。云提供商确实拥有可以帮助移动信息的工具。

## Migrating from cloud platforms as services

## 从云平台迁移为服务

If you’re using services specific to your cloud provider(s), you have some decisions to make. For example, you can continue to use proprietary databases like CosmosDB or DynamoDB if you construct your cluster in the same cloud as your existing application. Given the ubiquity of Kubernetes and that it’s not specific to any provider, you might want to consider moving away from proprietary dependencies. There is no right answer, only questions to make sure you consider.

如果您使用特定于您的云提供商的服务，您需要做出一些决定。例如，如果您在与现有应用程序相同的云中构建集群，则可以继续使用专有数据库，如 CosmosDB 或 DynamoDB。鉴于 Kubernetes 无处不在，而且它并不特定于任何提供商，您可能需要考虑摆脱专有依赖项。没有正确的答案，只有确保您考虑的问题。

## Onward to migration 

## 继续迁移

As you move forward in reaping the benefits of the production-grade container orchestration platform that is Kubernetes, remember that your system is yours and your context is unique. Every decision you make needs to serve your purposes so that you can better serve your users and your organization as a whole. Happy migrating!

当您继续享受生产级容器编排平台 Kubernetes 的好处时，请记住您的系统是您的，您的上下文是独一无二的。您做出的每个决定都需要服务于您的目的，以便您可以更好地为您的用户和整个组织服务。迁移快乐！

_Want to learn more? Watch Sensu CEO and co-founder Caleb Hailey's [webinar on filling the gaps in your Kubernetes observability strategy](https://sensu.io/resources/webinar/the-top-7-most-useful-kubernetes-apis-for-comprehensive-cloud-native-obsevability), in which he goes over the 7 most useful APIs for cloud-native observability, demonstrating how to get more context into what's going on with your K8s clusters._ 

_想了解更多？观看 Sensu 首席执行官兼联合创始人 Caleb Hailey 的 [关于填补 Kubernetes 可观察性战略空白的网络研讨会](https://sensu.io/resources/webinar/the-top-7-most-useful-kubernetes-apis-for-综合云原生可观察性)，其中介绍了用于云原生可观察性的 7 个最有用的 API，演示了如何获得更多关于 K8s 集群正在发生的事情的背景信息。_

