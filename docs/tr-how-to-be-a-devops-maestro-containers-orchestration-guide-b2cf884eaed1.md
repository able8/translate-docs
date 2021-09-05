# How to be a DevOps maestro: containers orchestration guide

# 如何成为 DevOps 大师：容器编排指南

“All things are difficult before they are easy.” — Thomas Fuller

“凡事难于易。” — 托马斯·富勒

## Introduction

##  介绍

Nowadays software packaged as container images and ran in containers is an  industry-accepted standard of running and distributing applications. Some of the benefits of using containers are:

- A high degree of portability
- Consistent operations processes
- Scalability and resiliency mostly
- Enabling common tooling thanks to open standards

如今，打包为容器映像并在容器中运行的软件是运行和分发应用程序的行业公认标准。使用容器的一些好处是：

- 高度便携
- 一致的操作流程
- 主要是可扩展性和弹性
- 由于开放标准，启用通用工具

As with everything, there are no free lunches! The power that containers  provide comes with complexity. Running standalone containers instances  can be ok for limited size workloads or at a smaller scale. At a certain point, the complexity is too much to deal with manually.

和所有事情一样，天下没有免费的午餐！容器提供的强大功能伴随着复杂性。运行独立容器实例可以适用于有限大小的工作负载或较小的规模。在某个点上，手动处理的复杂性太大了。

Developing and running containerized workloads at scale requires a high degree of  automation and operational excellence. In other words, it requires **containers orchestration.**

大规模开发和运行容器化工作负载需要高度自动化和卓越运营。换句话说，它需要**容器编排。**

After reading this article you will:

- understand what containers orchestration is
- why is it a critical part of running containers
- how can it help make your life easier

阅读本文后，您将：

- 了解什么是容器编排
- 为什么它是运行容器的关键部分
- 它如何帮助您的生活更轻松

It can also help you explain to your peers or leadership why and when  containers orchestration is important. You will also learn about useful  tools and patterns.

它还可以帮助您向同行或领导解释为什么以及何时容器编排很重要。您还将了解有用的工具和模式。

## What is containers orchestration?

## 什么是容器编排？

Traditionally, containers orchestration is associated mostly with operational level  activities. There are however lots of activities happening parallelly or before expanded to containers orchestration. By including development  as well as building and deployment automation related tasks a new  definition can be applied: **Containerized workloads management.**

传统上，容器编排主要与操作级活动相关联。然而，有许多活动同时发生或在扩展到容器编排之前发生。通过包括开发以及构建和部署自动化相关任务，可以应用新定义：**容器化工作负载管理。**

> Containerized workloads management is a process of automating common tasks throughout the whole lifecycle of container images and containers. Containers  orhestration is is a process of automating common operational-level  tasks.

> 容器化工作负载管理是在容器镜像和容器的整个生命周期中自动化常见任务的过程。容器管理是自动化常见操作级任务的过程。

The below diagram shows how **containerized workloads management** involves development, security, operations and other disciplines in the following areas:

![img](https://miro.medium.com/max/2000/1*HFSSx9WG3GXu8cozP75jMA.png)

Source: Author

下图显示了**容器化工作负载管理**如何涉及以下领域的开发、安全、运营和其他学科：

来源：作者

In this article, we will dig deeper into the “traditional” container orchestration side of Operations.

在本文中，我们将深入探讨 Operations 的“传统”容器编排方面。

## Example Scenario

## 示例场景

A sample scenario will help illustrate and understand all the principles.

示例场景将有助于说明和理解所有原则。

Imagine that you need to create and deploy a traditional n-tier app with a web  frontend, simple REST API, and document database as a persistence store. The setup will run on-prem and in a public cloud. We expect increased  traffic over time with a steady usage pattern. At some point our app  will undergo internal compliance and security audits, resulting in  additional requirements.

想象一下，您需要使用 Web 前端、简单的 REST API 和作为持久性存储的文档数据库来创建和部署传统的 n 层应用程序。该设置将在本地和公共云中运行。我们预计流量会随着时间的推移而增加，使用模式稳定。在某些时候，我们的应用程序将接受内部合规性和安全性审核，从而产生额外的要求。

## What you see is not all there is to it

## 你所看到的并不是全部

A typical tutorial or blog post showing how to get started with running  something in a container barely scratches the surface. Once a certain  scale is reached, orchestration becomes critical and unfortunately, most of it is not visible to the business.

展示如何开始在容器中运行某些东西的典型教程或博客文章几乎没有触及表面。一旦达到一定规模，编排就变得至关重要，不幸的是，大部分业务对业务不可见。

![img](https://miro.medium.com/max/2000/1*stRWa92BLo-6wplO46VbLw.jpeg)

Image by [Josep Monter Martinez](https://pixabay.com/users/josepmonter-1007570/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1321692) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1321692)

Let’s dig deeper into each task from the “Operations” area. The goal is to  show what problems containers orchestration tries to solve. We will look at patterns and paradigms as well as tools that help address different  areas of orchestration.

让我们从“操作”区域深入挖掘每项任务。目标是展示容器编排试图解决的问题。我们将研究有助于解决不同编排领域的模式和范例以及工具。

> **Provision and Deploy**

> **配置和部署**

Without tools, **provisioning** would typically mean sshing into a virtual machine and installing a  bunch of tools to get the nodes ready for receiving workloads. In our  example, it means a separate server for the frontend app, a separate one for middleware and another one for the database. 

如果没有工具，**配置** 通常意味着 SSH 到虚拟机并安装一堆工具以使节点准备好接收工作负载。在我们的示例中，它意味着前端应用程序的单独服务器，中间件的单独服务器和数据库的另一台服务器。

**A deployment** is an act of pulling images on the server and spinning up new  containers. Let’s assume that our application runs one container for  each component, so 3 containers in total. Without orchestration doing so would mean setting up a connection to a remote docker host and issuing  imperative commands like `docker run etc.`

**部署**是在服务器上拉取镜像并启动新容器的行为。假设我们的应用程序为每个组件运行一个容器，所以总共有 3 个容器。如果没有编排，这样做将意味着建立到远程 docker 主机的连接并发出命令式命令，如“docker run 等”。

How does this step look like with container orchestration tools and processes?

对于容器编排工具和流程，此步骤如何？

To enable smart **provisioning** we could take advantage of two patterns:

- [Infrastructure as Code](https://en.wikipedia.org/wiki/Infrastructure_as_code)
- [GitOps](https://github.com/gitops-working-group/gitops-working-group#gitops-principles)
- [Declarative Programming](https://en.wikipedia.org/wiki/Declarative_programming)

为了启用智能 **provisioning**，我们可以利用两种模式：

- [基础设施即代码](https://en.wikipedia.org/wiki/Infrastructure_as_code)
- [GitOps](https://github.com/gitops-working-group/gitops-working-group#gitops-principles)
- [声明式编程](https://en.wikipedia.org/wiki/Declarative_programming)

Many tools work great in this space. Highlighting a few here.

许多工具在这个领域工作得很好。这里重点介绍几个。

[Terraform by HashiCorp -Terraform is an open-source infrastructure as a code software tool that provides…www.terraform.io](https://www.terraform.io/)

[Terraform by HashiCorp-Terraform 是一个开源基础设施作为代码软件工具，它提供……www.terraform.io](https://www.terraform.io/)

[CrossplaneCrossplane brings Kubernetes-styled declarative and API-driven configuration and management to any piece of…crossplane.io](https://crossplane.io/)

[CrossplaneCrossplane 将 Kubernetes 风格的声明式和 API 驱动的配置和管理带到任何...crossplane.io](https://crossplane.io/)

[Pulumi - Modern Infrastructure as CodeAll architectures welcome Choose from over 50 cloud providers, including public, private, and hybrid architectures…www.pulumi.com](https://www.pulumi.com/)

[Pulumi - Modern Infrastructure as Code 欢迎所有架构从 50 多个云提供商中选择，包括公共、私有和混合架构...www.pulumi.com](https://www.pulumi.com/)

Deployment should be fully automated, especially at scale. A pattern to use here  would also be GitOps. Example tools in this space are:

[Flux enables continuous delivery of container images, using version control for each step to ensure deployment is…www.weave.works](https://www.weave.works/oss/flux/)

部署应该是完全自动化的，尤其是在规模上。此处使用的模式也将是 GitOps。此空间中的示例工具有：

[Flux 实现容器镜像的持续交付，对每个步骤使用版本控制以确保部署是...www.weave.works](https://www.weave.works/oss/flux/)

[Argo CD - Declarative GitOps CD for KubernetesArgo CD is a declarative, GitOps continuous delivery tool for Kubernetes. Application definitions, configurations, and…argoproj.github.io](https://argoproj.github.io/argo-cd/)

[Argo CD - Declarative GitOps CD for KubernetesArgo CD 是一个用于 Kubernetes 的声明性 GitOps 持续交付工具。应用程序定义、配置和...argoproj.github.io](https://argoproj.github.io/argo-cd/)

> Schedule on a correct node

> 在正确的节点上调度

Our workloads are containerized and ready, we can deploy them, but wait  where exactly? In a high scale environment, there are typically multiple compute nodes, moreover, the health of the nodes and available  resources can change over time. Without any tooling, we would have to  make decisions based on resources utilization, network latency, OS  version, etc. Due to the manual effort required, this decision would  typically be done once, not taking advantage of [bin packing](https://en.wikipedia.org/wiki/Bin_packing_problem).

我们的工作负载已经容器化并准备就绪，我们可以部署它们，但究竟在哪里等待？在大规模环境中，通常有多个计算节点，而且节点的健康状况和可用资源会随着时间的推移而发生变化。在没有任何工具的情况下，我们将不得不根据资源利用率、网络延迟、操作系统版本等做出决定。由于需要手动工作，这个决定通常只做一次，而不是利用 [bin Packaging](https://en.wikipedia.org/wiki/Bin_packing_problem)。

Kubernetes comes with a build-in scheduler controller, that takes care of exactly  this. Not only the scheduling is done once but is also evaluated and  adjusted dynamically.

Kubernetes 带有一个内置的调度程序控制器，它负责处理这个问题。不仅调度完成一次，而且还动态评估和调整。

Tools that help in this area: Kubernetes

在这方面有帮助的工具：Kubernetes

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[生产级容器编排Kubernetes，也称为K8s，是一个开源系统，用于自动化部署、扩展和管理……kubernetes.io](https://kubernetes.io/)

> Allocate Resources

> 分配资源

Each container needs compute resources, CPU, RAM, storage. Without  orchestration, there is no easy way of limiting how much resources a  container can consume. We are basically left with system monitoring  tools and retroactively fixing issues.

每个容器都需要计算资源、CPU、RAM、存储。如果没有编排，就没有简单的方法来限制容器可以消耗的资源量。我们基本上只剩下系统监控工具和追溯修复问题。

Kubernetes comes with a build-in mechanism to declare how much resources should be allocated to containers (requested resources). It also supports  resources limits that can be consumed by containers resources limits)

Kubernetes 带有一个内置机制来声明应该将多少资源分配给容器（请求的资源）。还支持容器资源限制可以消耗的资源限制）

Tools that help in this area: Kubernetes

在这方面有帮助的工具：Kubernetes

[Production-Grade Container Orchestration Kubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[生产级容器编排Kubernetes，也称为K8s，是一个开源系统，用于自动化部署、扩展和管理……kubernetes.io](https://kubernetes.io/)

> Conditionally auto-scale

> 有条件地自动缩放

Our app is a success and utilization is growing. Thankfully we have followed [The Twelve-Factor App](https://12factor.net/) design methodology and our frontend and middleware can [horizontally scale](https://en.wikipedia.org/wiki/Scalability#Horizontal_(scale_out)_and_vertical_scaling_(scale_up)) without a problem.

我们的应用程序是成功的，并且利用率正在增长。值得庆幸的是，我们遵循了 [The Twelve-Factor App](https://12factor.net/) 设计方法，我们的前端和中间件可以[水平扩展](https://en.wikipedia.org/wiki/Scalability#Horizontal_(scale_out)_and_vertical_scaling_(scale_up)) 没有问题。

Without orchestration, horizontal scaling of our frontend app in its simplest form involves:

- spinning up a new container
- creating a load balancer either as another container or a standalone installation (for example nginx)
- updating nginx configuration pointing to new IP addresses of the containers to the load balancer
- repeat if we want to scale up to additional containers 

在没有编排的情况下，以最简单的形式水平扩展我们的前端应用程序包括：

- 启动一个新容器
- 创建负载均衡器作为另一个容器或独立安装（例如 nginx）
- 将指向容器新 IP 地址的 nginx 配置更新到负载均衡器
- 如果我们想扩展到其他容器，请重复

Of course, if the demand drops, we should remove excessive containers. With a big enough scale, this is a full-time job!

当然，如果需求下降，我们应该移除过多的容器。规模够大，这是一份全职工作！

Kubernetes provides a build-in mechanism called [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) for scaling up and down based on resources utilization. What if we  wanted to scale based on completely arbitrary metrics or API calls? No  problem, KEDA provides a rich event-driven auto-scaling mechanism.

Kubernetes 提供了一种称为 [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) 的内置机制，用于根据资源利用率进行伸缩。如果我们想根据完全任意的指标或 API 调用进行扩展怎么办？没问题，KEDA 提供了丰富的事件驱动的自动缩放机制。

What if our database needs a more “beefy” server/node to support all the  requests coming from the middleware? Kubernetes got us covered, we can  create a new node pool for our database workloads and use a combination  of [taints and tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) to automatically migrate our database pods to a better-suited node.

如果我们的数据库需要一个更“强大”的服务器/节点来支持来自中间件的所有请求怎么办？ Kubernetes 为我们提供了保障，我们可以为我们的数据库工作负载创建一个新的节点池，并结合使用 [taints 和 tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/)自动将我们的数据库 pod 迁移到更适合的节点。

Tools that help in this area: Kubernetes, KEDA

在这方面有帮助的工具：Kubernetes、KEDA

[Production-Grade Container Orchestration Kubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[生产级容器编排Kubernetes，也称为K8s，是一个开源系统，用于自动化部署、扩展和管理……kubernetes.io](https://kubernetes.io/)

[KEDAKEDA is a Kubernetes-based Event Driven Autoscaler. With KEDA, you can drive the scaling of any container in Kubernetes…keda.sh](https://keda.sh/)

[KEDAKEDA 是一个基于 Kubernetes 的事件驱动自动缩放器。使用 KEDA，您可以驱动 Kubernetes 中任何容器的扩展……keda.sh](https://keda.sh/)

> Load balance and route traffic

> 负载平衡和路由流量

Let’s imagine, that we have added a new route to our frontend app and need to load balance it from publicly facing domain to appropriate pod and  container. Without tooling, it involves a manual reconfiguration of load balancer which can get very messy very fast.

让我们想象一下，我们已经向前端应用程序添加了一条新路由，并且需要将它从面向公众的域负载平衡到适当的 pod 和容器。如果没有工具，它需要手动重新配置负载平衡器，这会很快变得非常混乱。

Kubernetes provides native primitives (services, ingresses) to helps us automate load balancing tasks.

Kubernetes 提供了原生原语（服务、入口）来帮助我们自动执行负载平衡任务。

Although the DNS plugin is not part of “core” Kubernetes, it is almost always  installed on a cluster. DNS provides service discovery, no more fiddling with IP addresses in config files, localhost is enough!

尽管 DNS 插件不是“核心”Kubernetes 的一部分，但它几乎总是安装在集群上。 DNS 提供服务发现，无需再在配置文件中摆弄 IP 地址，localhost 就足够了！

Tools that help in this area: Kubernetes with CoreDNS plugin

在这方面有帮助的工具：带有 CoreDNS 插件的 Kubernetes

[CoreDNS: DNS and Service DiscoveryCoreDNS chains plugins. Each plugin performs a DNS function, such as Kubernetes service discovery, prometheus metrics…coredns.io](https://coredns.io/)

[CoreDNS：DNS 和服务发现CoreDNS 链插件。每个插件都执行一个 DNS 功能，例如 Kubernetes 服务发现、prometheus 指标……coredns.io](https://coredns.io/)

> Ensure high availability

> 确保高可用性

Mission-critical systems, such as our example app, must adhere to SLAs (Service Level  Agreements) defining the uptime of a service. [High availability](https://en.wikipedia.org/wiki/High_availability) means ensuring that the system must be up and available for a specified duration of time. For example “four nines” of availability (99.99% or  time) translates to 52 minutes per year or 4.38 minutes per month of  unavailability.

关键任务系统，例如我们的示例应用程序，必须遵守定义服务正常运行时间的 SLA（服务级别协议）。 [高可用性](https://en.wikipedia.org/wiki/High_availability) 意味着确保系统必须在指定的时间内启动并可用。例如，可用性的“四个 9”（99.99% 或时间)转化为每年 52 分钟或每月 4.38 分钟的不可用性。

Without any tools, it is very hard to ensure high availability.

没有任何工具，很难保证高可用。

Kubernetes comes with a powerful mechanism of “reconciliation loops” which are  controllers running on the cluster and constantly monitoring the  infrastructure and applications for any drifts from the desired state. Once such drift is detected, Kubernetes will automatically try to  restore balance by making sure desired state corresponds to the actual  state.

Kubernetes 带有一个强大的“协调循环”机制，它是在集群上运行的控制器，并不断监控基础设施和应用程序是否偏离所需状态。一旦检测到这种漂移，Kubernetes 将通过确保期望状态与实际状态对应来自动尝试恢复平衡。

There are two paradigms we can take advantage of here:

- [GitOps](https://github.com/gitops-working-group/gitops-working-group#gitops-principles)
- [Declarative Programming](https://en.wikipedia.org/wiki/Declarative_programming)

我们可以在这里利用两种范式：

- [GitOps](https://github.com/gitops-working-group/gitops-working-group#gitops-principles)
- [声明式编程](https://en.wikipedia.org/wiki/Declarative_programming)

Tools that help in this area: Kubernetes

在这方面有帮助的工具：Kubernetes

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[生产级容器编排Kubernetes，也称为K8s，是一个开源系统，用于自动化部署、扩展和管理……kubernetes.io](https://kubernetes.io/)

> Monitor and health check

> 监控和健康检查

Observability is a great example of a decentralized system design with a centralized  control plane. To observe our containers we will need highly  decentralized components operating on various levels of the stack  (kernel syscalls, system-wide events, application-specific monitoring  endpoints, etc). 

可观察性是具有集中控制平面的分散系统设计的一个很好的例子。为了观察我们的容器，我们需要在堆栈的各个级别（内核系统调用、系统范围的事件、特定于应用程序的监控端点等）上运行的高度分散的组件。

All those pieces of data are collected and presented in a meaningful way  either to humans or automated processes. Kubernetes can enable  cluster-wide [audit policy](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/#audit-policy) and collect data on various levels.

所有这些数据都以有意义的方式收集并呈现给人类或自动化流程。 Kubernetes 可以启用集群范围的 [审计策略](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/#audit-policy) 并收集各个级别的数据。

There are a lot of tools that can support this process, such as Prometheus, Graphana, Loki, FluentD, Jaeger, just to name a few.

有很多工具可以支持这个过程，比如 Prometheus、Graphana、Loki、FluentD、Jaeger 等等。

To enable monitoring on an application level, Kubernetes comes with a concept of [liveness, readiness and startup probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).

为了实现应用层面的监控，Kubernetes 自带了一个概念[liveness, readiness and startup probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)。

Cillium and Falco deserve a special mention. These tools can work with [eBPF](https://ebpf.io/) probes to provide deep level observability. Falco specializes in  security-related monitoring whereas Cillium does the same for  networking.

Cillium 和 Falco 值得特别一提。这些工具可以与 [eBPF](https://ebpf.io/) 探针配合使用以提供深层可观察性。 Falco 专注于与安全相关的监控，而 Cillium 在网络方面做同样的事情。

Tools that help in this area: Kubernetes, Prometheus, Jaeger, Grafana, Fluentd, Cillium, Falco

在这方面有帮助的工具：Kubernetes、Prometheus、Jaeger、Grafana、Fluentd、Cillium、Falco

[FalcoStrengthen container security The flexible rules engine allows you to describe any type of host or container behavior…falco.org](https://falco.org/)

[Falco加强容器安全灵活的规则引擎允许您描述任何类型的主机或容器行为……falco.org](https://falco.org/)

[Cilium - Linux Native, API-Aware Networking and Security for ContainersTraditional firewalls limit their inspection to the IP and TCP layers. Cilium uses eBPF to accelerate getting data in…cilium.io](https://cilium.io/)

[Cilium - Linux 原生、API 感知网络和容器安全性传统防火墙将其检查限制在 IP 和 TCP 层。 Cilium 使用 eBPF 加速在…cilium.io 中获取数据](https://cilium.io/)

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[生产级容器编排Kubernetes，也称为K8s，是一个开源系统，用于自动化部署、扩展和管理……kubernetes.io](https://kubernetes.io/)

[Prometheus - Monitoring system & time series databasePower your metrics and alerting with a leadingopen-source monitoring solution. © Prometheus Authors 2014-2021 |…prometheus.io](https://prometheus.io/)

[Prometheus - 监控系统和时间序列数据库通过领先的开源监控解决方案为您的指标和警报提供支持。 © Prometheus Authors 2014-2021 |…prometheus.io](https://prometheus.io/)

[Jaeger: open source, end-to-end distributed tracingMonitor and troubleshoot transactions in complex distributed systemswww.jaegertracing.io](https://www.jaegertracing.io/)

 [Jaeger：开源、端到端的分布式追踪监控并解决复杂分布式系统中的事务www.jaegertracing.io](https://www.jaegertracing.io/)

[Grafana: The open observability platformGrafana is the open source analytics & monitoring solution for every database.grafana.com](https://grafana.com/)

[Grafana：开放的可观察性平台Grafana 是每个database.grafana.com 的开源分析和监控解决方案](https://grafana.com/)

[Fluentd | Open Source Data CollectorFluentd is an open source data collector for unified logging layer.www.fluentd.org](https://www.fluentd.org/)

[流利|开源数据收集器Fluentd是统一日志层的开源数据收集器。www.fluentd.org](https://www.fluentd.org/)

> Configure

> 配置

There are many ways of providing configuration to containerized workloads. Reading environmental variables, config files, external endpoints such  as databases or HTTP/gRPC services etc.

有很多方法可以为容器化工作负载提供配置。读取环境变量、配置文件、外部端点（如数据库或 HTTP/gRPC 服务等）。

Kubernetes provides consistent configuration management via the following primitives:

- Config Maps
- Secrets
- Volumes

Kubernetes 通过以下原语提供一致的配置管理：

- 配置地图
- 秘密
- 卷

All those configuration options can be “mounted” to pods in various ways; providing a consistent model of configuration management.

所有这些配置选项都可以通过各种方式“挂载”到 Pod；提供一致的配置管理模型。

Tools that help in this area: Kubernetes

在这方面有帮助的工具：Kubernetes

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[生产级容器编排Kubernetes，也称为K8s，是一个开源系统，用于自动化部署、扩展和管理……kubernetes.io](https://kubernetes.io/)

> Security and secure communication

> 安全和安全的通信

Last, but definitely not least, how to ensure secure communication between  containers? Kubernetes without any additional settings or changes is  very permissive. Thankfully, there are tools and standards to help us  harden and secure the cluster.

最后但同样重要的是，如何确保容器之间的安全通信？没有任何额外设置或更改的 Kubernetes 是非常宽松的。值得庆幸的是，有一些工具和标准可以帮助我们加固和保护集群。

Let’s imagine that our application is being audited and a security  recommendation was made to ensure that our database container cannot  talk to the frontend container. As long as we don’t scale anything, it  is possible to do with docker-compose and networking settings, but it is a manual process.

假设我们的应用程序正在接受审计，并提出了安全建议以确保我们的数据库容器无法与前端容器通信。只要我们不扩展任何东西，就可以使用 docker-compose 和网络设置，但这是一个手动过程。

Kubernetes comes with a built-in mechanism of [network policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/) that regulate which pods can communicate with each other.

Kubernetes 带有一个内置的 [网络策略](https://kubernetes.io/docs/concepts/services-networking/network-policies/) 机制，用于管理哪些 Pod 可以相互通信。

This is great, but after a while, another team deployed a service to the  cluster. This service provides sensitive lookup data, we have to ensure  that the communication between our pods and the service is encrypted.

这很好，但过了一会儿，另一个团队向集群部署了一个服务。该服务提供敏感的查找数据，我们必须确保我们的 Pod 和服务之间的通信是加密的。

Here we can use tools from the Service Mesh category. Those tools provide  enhanced observability as well as secure communication capabilities. Istio and LinkerD for example. 

在这里，我们可以使用 Service Mesh 类别中的工具。这些工具提供了增强的可观察性以及安全的通信功能。例如 Istio 和 LinkerD。

Maybe we want to prevent pods from using the default public Docker Hub registry for docker images? Kubernetes comes with [admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) and Image Policy Webhook is one of them. If we want to create more  fine-grained policies, tools such as Open Policy Agent(OPA) or Kyverno  will be helpful.

也许我们想阻止 pod 使用默认的公共 Docker Hub 注册表来存储 docker 镜像？ Kubernetes 带有 [准入控制器](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/)，Image Policy Webhook 就是其中之一。如果我们想创建更细粒度的策略，Open Policy Agent (OPA) 或 Kyverno 等工具会很有帮助。

Tools that help in this area: Kubernetes, Istio, LinkerD, OPA, Kyverno

在这方面有帮助的工具：Kubernetes、Istio、LinkerD、OPA、Kyverno

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[生产级容器编排Kubernetes，也称为K8s，是一个开源系统，用于自动化部署、扩展和管理……kubernetes.io](https://kubernetes.io/)

[Open Policy AgentStop using a different policy language, policy model, and policy API for every product and service you use. Use OPA for…www.openpolicyagent.org](https://www.openpolicyagent.org/)

[打开策略代理停止为您使用的每个产品和服务使用不同的策略语言、策略模型和策略 API。将 OPA 用于……www.openpolicyagent.org](https://www.openpolicyagent.org/)

[KyvernoKyverno is a policy engine designed for Kubernetes. With Kyverno, policies are managed as Kubernetes resources and no…kyverno.io](https://kyverno.io/)

[KyvernoKyverno 是为 Kubernetes 设计的策略引擎。使用 Kyverno，策略作为 Kubernetes 资源进行管理，而不是……kyverno.io](https://kyverno.io/)

[IstioA service mesh for observability, security in depth, and management that speeds deployment cycles.istio.io](https://istio.io/)

[Istio 用于可观察性、深度安全性和可加速部署周期的管理的服务网格.istio.io](https://istio.io/)

[The world's lightest, fastest service mesh.Linkerd adds critical security, observability, and reliability to your Kubernetes stack, without any code changes.linkerd.io](https://linkerd.io/)

[世界上最轻、最快的服务网格。Linkerd 为您的 Kubernetes 堆栈增加了关键的安全性、可观察性和可靠性，无需任何代码更改.linkerd.io](https://linkerd.io/)

## Conclusion

##  结论

Containers orchestration captures the “day-2” operational tasks that sooner or  later we will have to deal with. Fortunately, there are plenty of tools  and standards to help us with this task. The tools mentioned in this  article are examples and mostly ones I’m personally familiar with, but  there are a lot of alternatives. If you are interested in learning more  about those, check out the ever-growing [CNCF landscape](https://landscape.cncf.io/?category=application-definition-image-build&fullscreen=yes&grouping=category).

容器编排捕获了我们迟早要处理的“第 2 天”操作任务。幸运的是，有很多工具和标准可以帮助我们完成这项任务。本文中提到的工具是示例，主要是我个人熟悉的工具，但还有很多替代方案。如果您有兴趣了解更多相关信息，请查看不断增长的 [CNCF 景观](https://landscape.cncf.io/?category=application-definition-image-build&fullscreen=yes&grouping=category)。

[CNCF Cloud Native Interactive LandscapeThe Cloud Native Trail Map ( png, pdf) is CNCF's recommended path through the cloud native landscape. The cloud native…landscape.cncf.io](https://landscape.cncf.io/?category=application-definition-image-build&fullscreen=yes&grouping=category)

 [CNCF Cloud Native Interactive Landscape Cloud Native Trail Map (png, pdf) 是 CNCF 推荐的穿越云原生景观的路径。云原生...landscape.cncf.io](https://landscape.cncf.io/?category=application-definition-image-build&fullscreen=yes&grouping=category)

I hope that this short guide helped you understand what container  orchestration is, why it is important and how you can start thinking  about it in your organization.

我希望这个简短的指南可以帮助您了解什么是容器编排、为什么它很重要以及如何在您的组织中开始考虑它。

[ITNEXT](https://itnext.io/?source=post_sidebar--------------------------post_sidebar-----------)

ITNEXT is a platform for IT developers & software engineers… 

ITNEXT 是一个面向 IT 开发人员和软件工程师的平台……

