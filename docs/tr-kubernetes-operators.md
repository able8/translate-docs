# Kubernetes operators: Embedding operational expertise side by side with containerized applications

# Kubernetes operators：将运营专业知识与容器化应用程序并排嵌入

Kubernetes isn't complex, your business problem is. Learn how operators make it easy to run complex software at scale.

Kubernetes 并不复杂，您的业务问题才是。了解操作员如何使大规模运行复杂软件变得容易。

March 2, 2020 From: https://www.redhat.com/sysadmin/kubernetes-operators

**Why use operators?** 

**为什么要使用operators？**

At a high level, the [Kubernetes Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) makes it easy to run complex software at scale. With the DevOps movement, we learned to manage and monitor complex applications and infrastructure from centralized platforms (Chef Server, Puppet Enterprise, Ansible Tower, Nagios, etc). This centralized monitoring & automated remediation works great with relatively stable infrastructure components like bare metal servers, virtual machines, and network devices. However, containers change much quicker than traditional infrastructure and traditional applications. One might say this speed is a tenet of Cloud Native behavior.

在高层次上，[Kubernetes Operator 模式](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) 可以轻松地大规模运行复杂的软件。随着 DevOps 运动，我们学会了从集中式平台（Chef Server、Puppet Enterprise、Ansible Tower、Nagios 等)管理和监控复杂的应用程序和基础设施。这种集中监控和自动修复非常适用于相对稳定的基础架构组件，例如裸机服务器、虚拟机和网络设备。然而，容器的变化比传统基础设施和传统应用程序要快得多。有人可能会说这种速度是云原生行为的一个原则。

The ephemeral nature of containerized applications drives an immense amount of state change. In turn, this rapid state change challenges the ability of centralized monitoring and automation to keep up. Kubernetes can literally change the state of containers every few seconds. The solution is to bring the automation close to the applications - to deploy it within Kubernetes so that it has direct access to the application’s state at all times. The Kubernetes Operator pattern allows us to do just this - deploy automation side by side with the containerized application.

容器化应用程序的短暂性推动了大量的状态变化。反过来，这种快速的状态变化挑战了集中监控和自动化跟上的能力。 Kubernetes 可以每隔几秒改变一次容器的状态。解决方案是将自动化靠近应用程序——将其部署在 Kubernetes 中，以便它始终可以直接访问应用程序的状态。 Kubernetes Operator 模式允许我们做到这一点 - 与容器化应用程序并行部署自动化。

I often describe the Operator pattern as deploying a robot sysadmin next to the containerized application. Though, to truly understand how Operators work, we need to dive a bit deeper into the history of running services, and the art of making these services resilient. From here on, we will refer to this as operational knowledge, or operational excellence.

我经常将 Operator 模式描述为在容器化应用程序旁边部署机器人系统管理员。但是，要真正了解 Operator 的工作原理，我们需要更深入地了解运行服务的历史，以及使这些服务具有弹性的艺术。从这里开始，我们将其称为运营知识或卓越运营。

![](https://www.redhat.com/sysadmin/sites/default/files/styles/embed_large/public/2020-02/Kube%201.png?itok=MB_IOTbL)




Traditionally, when new software was deployed, we also deployed a real, human Sysadmin to care and feed the application. This care and feeding included tasks like installation, upgrades, backups, & restores, troubleshooting, and return to service. If a service failed, we paged this Sysadmin, they logged into a server, would troubleshoot the application, and fix what was broken. To track this work, they would document their progress in a ticketing system.

传统上，在部署新软件时，我们还部署了一个真正的人类系统管理员来照顾和提供应用程序。这种照顾和喂养包括安装、升级、备份和恢复、故障排除和恢复服务等任务。如果服务失败，我们会呼叫该系统管理员，他们会登录服务器，对应用程序进行故障排除，并修复损坏的部分。为了跟踪这项工作，他们将在票务系统中记录他们的进度。

As we moved into the world of containers (about six years ago at the time of this writing), we updated the packaging format of the application, but we continued to deploy a real human Sysadmin to care and feed the application. This simplified installation, and upgrades, but did little for data backups/restores and break/fix work.

随着我们进入容器世界（大约在撰写本文时的六年前），我们更新了应用程序的打包格式，但我们继续部署真正的人类系统管理员来照顾和提供应用程序。这种简化的安装和升级，但对数据备份/恢复和中断/修复工作几乎没有作用。

With the advent of the Kubernetes Operator pattern, we deploy the software application in a container and we also deploy a robot Sysadmin in the same environment to care and feed the application. We refer to these robot Sysadmins as Operators. Not only can they perform installation, upgrades, backups, and restores, but they can also perform more complex tasks like recovering a database when the tables become corrupted. This takes us into another world of automation and reliability.

随着 Kubernetes Operator 模式的出现，我们将软件应用程序部署在一个容器中，我们还在同一环境中部署了一个机器人系统管理员来照顾和馈送应用程序。我们将这些机器人系统管理员称为操作员。他们不仅可以执行安装、升级、备份和还原，还可以执行更复杂的任务，例如在表损坏时恢复数据库。这将我们带入了另一个自动化和可靠性的世界。

But, why does this work philosophically? You might still struggle to understand how this is different than traditional clustering, or having a monitoring system participate in “self-healing” (an old buzzword now). Let’s explain...

但是，为什么这在哲学上有效？您可能仍然很难理解这与传统集群有何不同，或者让监控系统参与“自我修复”（现在是一个古老的流行词）。我们来解释一下……

**A brief history of operational knowledge**

**操作知识简史**

Operators embed operational knowledge, or operational excellence in the infrastructure, close to the service. To better explain this, let’s walk through the history of operating servers & services.

operators将运营知识或卓越运营嵌入到服务附近的基础设施中。为了更好地解释这一点，让我们回顾一下运营服务器和服务的历史。

Image

图片

 

- **Operational excellence 1.0** \- One computer, multiple administrators. Back in years past, before most users of containers remember, there were large computers cared for, managed and operated by multiple systems administrators. This remained true all the way into the late 1990s with large Mainframes and Unix systems managed by multiple systems administrators. That’s right, there were multiple human beings assigned to one computer. These systems administrators automated operations on these systems, fixed services if they were broken, added/removed users, and even wrote device drivers themselves. The administrators were highly technical, and would be considered software engineers by modern definitions. They performed all software related tasks, maintaining uptime, reliability, and operational excellence. The cost was high, but the quality was also very high.

- **卓越运营 1.0** \- 一台电脑，多个管理员。多年前，在大多数容器用户还记得之前，有多个系统管理员照顾、管理和操作大型计算机。这种情况一直持续到 1990 年代后期，大型主机和 Unix 系统由多个系统管理员管理。没错，一台电脑分配了多个人。这些系统管理员在这些系统上实现自动化操作，修复损坏的服务，添加/删除用户，甚至自己编写设备驱动程序。管理员技术含量很高，按照现代定义将被视为软件工程师。他们执行所有与软件相关的任务，保持正常运行时间、可靠性和卓越运营。成本很高，但质量也很高。

- **Operational Excellence 2.0** \- One administrator, multiple computers. With the advent of Linux, and perhaps more so Windows, the number of servers outgrew the number of administrators. There became less and less large, multi-user systems. At the same time there became more and more single service systems - DNS Servers, Web Servers, Mail Servers, etc. A single administrator could manage multiple servers remotely, and we often measured productivity by comparing the number of servers per administrator. Administrators still retained intimate knowledge of each service they managed, sometimes using [Runbooks](https://en.wikipedia.org/wiki/Runbook) to document common tasks. If a service or server failed, an administrator would work on it remotely thereby maintaining a high level of quality, while at the same time being responsible for a higher total number of servers.

- **Operational Excellence 2.0** \- 一位管理员，多台计算机。随着 Linux 的出现，也许更多的是 Windows，服务器的数量超过了管理员的数量。大型多用户系统变得越来越少。与此同时，出现了越来越多的单一服务系统——DNS 服务器、Web 服务器、邮件服务器等。一个管理员可以远程管理多个服务器，我们经常通过比较每个管理员的服务器数量来衡量生产力。管理员仍然对他们管理的每项服务保持深入了解，有时使用 [Runbooks](https://en.wikipedia.org/wiki/Runbook) 来记录常见任务。如果服务或服务器出现故障，管理员将远程处理它，从而保持高水平的质量，同时负责更多的服务器总数。

- **Operational Excellence 3.0** \- One service, multiple computers. As a sort of operational dead end, there was a renewed focus on the quality of the service by leveraging clustering and automatic recovery on cheaper hardware. Databases, web servers, NFS, DNS SAMBA servers and more were clustered for resilience. Tools like Veritas Clustering, Veritas File System, Pacemaker, GFS, and GPFS became popular. For this to work properly, operational knowledge of how to start and recover a service had to be configured in the clustering software. This required a solid understanding of how the service worked (detection of failures, how to restart, etc). With clustering software and configuration, nodes could be treated as capacity (N+1 means one extra server, N+2 means two extra servers). Bringing operational knowledge close to the service allowed for automated recovery, but building new services or clusters, much less decommissioning them, could take days or even weeks because each service had to be designed and maintained separately. 

- **Operational Excellence 3.0** \- 一项服务，多台计算机。作为一种运营的死胡同，通过在更便宜的硬件上利用集群和自动恢复，重新关注服务质量。数据库、Web 服务器、NFS、DNS SAMBA 服务器等都进行了集群以实现弹性。 Veritas Clustering、Veritas File System、Pacemaker、GFS 和 GPFS 等工具开始流行。为了使其正常工作，必须在集群软件中配置有关如何启动和恢复服务的操作知识。这需要对服务的工作方式有深入的了解（检测故障、如何重新启动等）。使用集群软件和配置，节点可以被视为容量（N+1 表示多一台服务器，N+2 表示多两台服务器）。将运营知识引入服务允许自动恢复，但构建新服务或集群，更不用说停用它们，可能需要数天甚至数周的时间，因为每个服务都必须单独设计和维护。

- **Operational Excellence 4.0** \- With this iteration, we moved the logic for recovering services back away from the service itself, and put it in the monitoring systems and load balancers. We embedded configuration in DNS, and started to use configuration management to maintain things. This created a fundamental tension in a lot of IT organizations. Server administrators would embed the logic for recovering services in the monitoring system. For example, if the monitoring system saw a problem with a web server it could ssh into the server and restart Apache. There were several major challenges with this paradigm. First, having configuration in many different places created a lot of complexity (see also: [Why you don't have to be afraid of Kubernetes](https://opensource.com/article/19/10/kubernetes-complex-business-problem)). Second, storage, network, and virtualization administrators didn’t want automation logging in and provisioning/deprovisioning services, so achieving truly cloud-native architectures was difficult.

- **Operational Excellence 4.0** \- 通过这次迭代，我们将恢复服务的逻辑从服务本身移回，并将其放入监控系统和负载平衡器中。我们将配置嵌入到 DNS 中，并开始使用配置管理来维护事物。这在许多 IT 组织中造成了根本性的紧张。服务器管理员将在监控系统中嵌入恢复服务的逻辑。例如，如果监控系统发现 Web 服务器出现问题，它可以通过 ssh 进入服务器并重新启动 Apache。这种范式存在几个主要挑战。首先，在许多不同的地方进行配置会带来很多复杂性（另请参阅：[为什么你不必害怕 Kubernetes](https://opensource.com/article/19/10/kubernetes-complex-business-问题）)。其次，存储、网络和虚拟化管理员不希望自动登录和配置/取消配置服务，因此实现真正的云原生架构很困难。

- **Operational Excellence 5.0** \- While centralized monitoring and automation can be abused with containers to achieve a [Desired State/Actual State](https://kubernetes.io/docs/concepts/#overview) model similar to Kubernetes , it's not nearly as elegant. Using a manifest (YAML or JSON), Kubernetes enables the powerful desired state. For example, once an administrator defines that they want three copies of a web server running, Kubernetes will maintain exactly three copies running. If one of the containerized web servers dies, Kubernetes will restart another one because it sees that the actual state doesn’t match the desired state. This gives you application definition and recovery in one file format. But, how do you manage more complex tasks like scanning corrupted database tables, upgrades, schema changes, or rebalancing data between volumes? That’s what operators do. It moves this logic close to the service bringing the best of Operational Excellence 3.0 and 4.0 together in one place. This same logic can be applied to the Kubernetes platform itself (see also: [Red Hat OpenShift Container Platform 4 now defaults to CRI-O as underlying container engine](https://www.redhat.com/en/blog/red-hat-openshift-container-platform-4-now-defaults-cri-o-underlying-container-engine)).


- **Operational Excellence 5.0** \- 虽然集中监控和自动化可以与容器一起滥用以实现类似于 Kubernetes 的 [Desired State/Actual State](https://kubernetes.io/docs/concepts/#overview) 模型，它几乎没有那么优雅。使用清单（YAML 或 JSON)，Kubernetes 可以实现强大的所需状态。例如，一旦管理员定义他们想要运行 Web 服务器的三个副本，Kubernetes 将保持运行的三个副本。如果其中一个容器化 Web 服务器死机，Kubernetes 将重新启动另一个，因为它发现实际状态与所需状态不匹配。这为您提供了一种文件格式的应用程序定义和恢复。但是，您如何管理更复杂的任务，例如扫描损坏的数据库表、升级、架构更改或在卷之间重新平衡数据？这就是operators所做的。它使这一逻辑更接近于将卓越运营 3.0 和 4.0 的最佳之处集中在一个地方的服务。同样的逻辑可以应用于 Kubernetes 平台本身（另请参阅：[Red Hat OpenShift Container Platform 4 现在默认使用 CRI-O 作为底层容器引擎](https://www.redhat.com/en/blog/red-hat-openshift-container-platform-4-now-defaults-cri-o-underlying-container-engine))。


**Who should build operators**

**谁应该建立operators**

This leads us to the important question of who should be building operators? Well, the short answer is, it depends.

这将我们引向了一个重要的问题，即谁应该是建筑operators？嗯，简短的回答是，这取决于。

If you are a developer building Java web services, Python web services, or really any web services then you shouldn’t have to write an Operator. Your web services should be relatively stateless and as long as you focus on learning to use [readiness and liveness checks](https://opensource.com/article/19/10/kubernetes-complex-business-problem), Kubernetes will manage the desired state for you. That said, you might use pre-built Operators to manage a PostgreSQL, MongoDB, or MariaDB instance. Check out all of the services you can consume from [Operator Hub](https://operatorhub.io/). Welcome to the Kubernetes ecosystem.

如果您是构建 Java Web 服务、Python Web 服务或任何 Web 服务的开发人员，那么您不必编写 Operator。你的 web 服务应该是相对无状态的，只要你专注于学习使用 [readiness and liveness checks](https://opensource.com/article/19/10/kubernetes-complex-business-problem)，Kubernetes 将管理你想要的状态。也就是说，您可以使用预构建的 Operator 来管理 PostgreSQL、MongoDB 或 MariaDB 实例。从 [Operator Hub](https://operatorhub.io/) 查看您可以使用的所有服务。欢迎来到 Kubernetes 生态系统。

If you are an administrator, practicing DevOps, and you are building complex services for others to consume, then you may very well need to think about writing operators. You will almost certainly be consuming Operators and upgrading them, so check out the [Operator Lifecycle Manager](https://github.com/operator-framework/operator-lifecycle-manager) (Built into OpenShift 3.X and 4.X ) and [Operator SDK.](https://github.com/operator-framework/operator-sdk) 

如果您是一名管理员，正在实践 DevOps，并且您正在构建复杂的服务供他人使用，那么您可能非常需要考虑编写operators。您几乎肯定会使用 Operator 并升级它们，因此请查看 [Operator Lifecycle Manager](https://github.com/operator-framework/operator-lifecycle-manager)（内置于 OpenShift 3.X 和 4.X ) 和 [Operator SDK.](https://github.com/operator-framework/operator-sdk)

If you’re a developer working on a data management, networking, monitoring, storage, security, or GPU solution for Kubernetes/OpenShift then you are the most likely to need to write an operator. I would suggest looking at the [Operator Framework](https://github.com/operator-framework/getting-started), [Operator SDK](https://github.com/operator-framework/operator-sdk), [Operator Lifecycle Manager](https://github.com/operator-framework/operator-lifecycle-manager), and [Operator Metering](https://github.com/operator-framework/operator-metering). You might also like to look at the Red Hat container certification program called [Partner Connect](https://connect.redhat.com/).

如果您是为 Kubernetes/OpenShift 开发数据管理、网络、监控、存储、安全或 GPU 解决方案的开发人员，那么您最有可能需要编写operators。我建议查看 [Operator Framework](https://github.com/operator-framework/getting-started)、[OperatorSDK](https://github.com/operator-framework/operator-sdk)， [Operator Lifecycle Manager](https://github.com/operator-framework/operator-lifecycle-manager) 和 [Operator Metering](https://github.com/operator-framework/operator-metering)。您可能还想查看名为 [Partner Connect](https://connect.redhat.com/) 的 Red Hat 容器认证计划。

**Conclusion**

**结论**

As I have said before, Kubernetes isn’t complex, your business problem is. Kubernetes not only redefines the operating system ( [OpenShift is the New Enterprise Linux](https://www.linkedin.com/pulse/openshift-new-enterprise-linux-daniel-riek/) and [Sorry, Linux. Kubernetes is now the OS that matters](https://www.infoworld.com/article/3322120/sorry-linux-kubernetes-is-now-the-os-that-matters.html)), it redefines the entire operational paradigm. Surely, the history of operational excellence could be divided into any number of paradigm shifts, but I have attempted to break it down with a model that is cognitively digestible to help operations teams, software vendors (ISVs), and even developers better understand why Operators are important.

正如我之前所说，Kubernetes 并不复杂，您的业务问题很复杂。 Kubernetes 不仅重新定义了操作系统（[OpenShift 是新的企业 Linux](https://www.linkedin.com/pulse/openshift-new-enterprise-linux-daniel-riek/) 和 [抱歉，Linux.Kubernetes 是现在是重要的操作系统](https://www.infoworld.com/article/3322120/sorry-linux-kubernetes-is-now-the-os-that-matters.html))，它重新定义了整个操作范式。当然，卓越运营的历史可以分为任意数量的范式转变，但我试图用一个在认知上易于理解的模型来分解它，以帮助运营团队、软件供应商(ISV) 甚至开发人员更好地理解为什么operators是重要的。

This fifth generation of operational excellence, using the Kubernetes Operator pattern brings the automation close to the application giving it access to state in near real-time. With Operators deployed side by side within Kubernetes, response, management and recovery of applications can happen at a speed that just isn’t possible with human beings and ticket systems. An added benefit is the ability to provision multiple copies of an application with a single command, or more importantly deprovision them. This ability to provision, deprovision, recover, and upgrade is the fundamental difference between cloud-native and traditional applications.

这第五代卓越运营，使用 Kubernetes Operator 模式使自动化接近应用程序，使其能够近乎实时地访问状态。随着 Operator 在 Kubernetes 中并行部署，应用程序的响应、管理和恢复可以以人类和票务系统无法实现的速度进行。一个额外的好处是能够使用单个命令配置应用程序的多个副本，或者更重要的是取消配置它们。这种供应、取消供应、恢复和升级的能力是云原生应用程序和传统应用程序之间的根本区别。

**Credits**

人员

I want to give special thanks to Daniel Riek who presented this concept of Operational Excellence at FOSDEM 20 in Brussels, Belgium last week. If you didn’t have an opportunity to attend his talk, I recommend you watch it when the video goes live. Until then, see this interview with him: [How Containers and Kubernetes re-defined the GNU/Linux Operating System. A Greybeard's Worst Nightmare](https://fosdem.org/2020/interviews/daniel-riek/)

我要特别感谢 Daniel Riek，他上周在比利时布鲁塞尔的 FOSDEM 20 上提出了这一卓越运营概念。如果你没有机会参加他的演讲，我建议你在视频直播时观看。在此之前，请参阅对他的采访：[容器和 Kubernetes 如何重新定义 GNU/Linux 操作系统。灰胡子最糟糕的噩梦](https://fosdem.org/2020/interviews/daniel-riek/)

## Scott McCarty

## 斯科特麦卡蒂

At Red Hat, Scott McCarty is a technical product manager for the container subsystem team, which enables key product capabilities in OpenShift Container Platform and Red Hat Enterprise Linux. Focus areas include container runtimes, tools, and images.
[More about me](http://www.redhat.com/sysadmin/users/scott-mccarty)

在 Red Hat，Scott McCarty 是容器子系统团队的技术产品经理，负责在 OpenShift Container Platform 和 Red Hat Enterprise Linux 中实现关键产品功能。重点领域包括容器运行时、工具和图像。
[更多关于我](http://www.redhat.com/sysadmin/users/scott-mccarty)
