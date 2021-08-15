## Introducing Operators: Putting Operational Knowledge into Software

## Operators介绍：将操作知识应用到软件中

November 3, 2016 · By Brandon Philips From: https://web.archive.org/web/20170129131616/https://coreos.com/blog/introducing-operators.html

A Site Reliability Engineer (SRE) is a person that operates an  application by writing software. They are an engineer, a developer, who  knows how to develop software specifically for a particular application  domain. The resulting piece of software has an application's operational domain knowledge programmed into it.

站点可靠性工程师 (SRE) 是通过编写软件来操作应用程序的人员。他们是工程师、开发人员，知道如何专门为特定应用领域开发软件。由此产生的软件具有编程到其中的应用程序的操作领域知识。

Our team has been busy in the Kubernetes community designing and  implementing this concept to reliably create, configure, and manage  complex application instances atop Kubernetes.

我们的团队一直忙于 Kubernetes 社区设计和实现这一概念，以在 Kubernetes 上可靠地创建、配置和管理复杂的应用程序实例。

We call this new class of software Operators. An Operator is an  application-specific controller that extends the Kubernetes API to  create, configure, and manage instances of complex stateful applications on behalf of a Kubernetes user. It builds upon the basic Kubernetes  resource and controller concepts but includes domain or  application-specific knowledge to automate common tasks.

我们称这类新的软件Operators为。 Operator 是一个特定于应用程序的控制器，它扩展了 Kubernetes API 以代表 Kubernetes 用户创建、配置和管理复杂的有状态应用程序的实例。它建立在基本的 Kubernetes 资源和控制器概念之上，但包括域或特定于应用程序的知识来自动执行常见任务。

## Stateless is Easy, Stateful is Hard

## 无状态很容易，有状态很难

With Kubernetes, it is relatively easy to manage and scale web apps,  mobile backends, and API services right out of the box. Why? Because  these applications are generally stateless, so the basic Kubernetes  APIs, like Deployments, can scale and recover from failures without  additional knowledge.

使用 Kubernetes，开箱即用地管理和扩展 Web 应用程序、移动后端和 API 服务相对容易。为什么？因为这些应用程序通常是无状态的，所以基本的 Kubernetes API，比如部署，可以在没有额外知识的情况下扩展和从故障中恢复。

A larger challenge is managing stateful applications, like databases, caches, and monitoring systems. These systems require application  domain knowledge to correctly scale, upgrade, and reconfigure while  protecting against data loss or unavailability. We want this  application-specific operational knowledge encoded into software that  leverages the powerful Kubernetes abstractions to run and manage the  application correctly.

更大的挑战是管理有状态的应用程序，如数据库、缓存和监控系统。这些系统需要应用领域知识才能正确扩展、升级和重新配置，同时防止数据丢失或不可用。我们希望将这种特定于应用程序的操作知识编码到软件中，利用强大的 Kubernetes 抽象来正确运行和管理应用程序。

An Operator is software that encodes this domain knowledge and extends the Kubernetes API through the [*third party resources*](https://web.archive.org/web/20170129131616/http://kubernetes.io/docs/user- guide/thirdpartyresources/) mechanism, enabling users to create, configure, and manage  applications. Like Kubernetes's built-in resources, an Operator doesn't  manage just a single instance of the application, but multiple instances across the cluster.

Operator 是对这些领域知识进行编码并通过 [*第三方资源*](https://web.archive.org/web/20170129131616/http://kubernetes.io/docs/user- guide/thirdpartyresources/) 机制，使用户能够创建、配置和管理应用程序。与 Kubernetes 的内置资源一样，Operator 不仅管理应用程序的单个实例，还管理集群中的多个实例。

![img](https://web.archive.org/web/20170129131616/https://coreos.com/assets/blog/2016-11-03-introducing-the-etcd-operator/Overview-etcd.png)

To demonstrate the Operator concept in running code, we have two concrete examples to announce as open source projects today:

1. The [*etcd Operator*](https://web.archive.org/web/20170129131616/https://coreos.com/blog/introducing-the-etcd-operator.html) creates, configures, and manages etcd clusters. etcd is a reliable,  distributed key-value store introduced by CoreOS for sustaining the most critical data in a distributed system, and is the primary configuration datastore of Kubernetes itself.
 2. The [*Prometheus Operator*](https://web.archive.org/web/20170129131616/https://coreos.com/blog/the-prometheus-operator.html) creates, configures, and manages Prometheus monitoring instances. Prometheus is a powerful monitoring, metrics, and alerting tool, and a  Cloud Native Computing Foundation (CNCF) project supported by the CoreOS team.

为了演示运行代码中的 Operator 概念，我们今天宣布两个具体的示例为开源项目：
1. [*etcd Operator*](https://web.archive.org/web/20170129131616/https://coreos.com/blog/introducing-the-etcd-operator.html)创建、配置和管理etcd 集群。 etcd 是 CoreOS 引入的一种可靠的分布式键值存储，用于维护分布式系统中最关键的数据，并且是 Kubernetes 本身的主要配置数据存储。
2. [*Prometheus Operator*](https://web.archive.org/web/20170129131616/https://coreos.com/blog/the-prometheus-operator.html) 创建、配置和管理 Prometheus 监控实例。 Prometheus 是一个强大的监控、指标和警报工具，是 CoreOS 团队支持的云原生计算基金会 (CNCF) 项目。

## How is an Operator Built? 
## Operator 是如何构建的？
Operators build upon two central Kubernetes concepts: Resources and Controllers. As an example, the built-in [*ReplicaSet*](https://web.archive.org/web/20170129131616/http://kubernetes.io/docs/user-guide/replicasets/) resource lets users set a desired number number of Pods to run, and  controllers inside Kubernetes ensure the desired state set in the  ReplicaSet resource remains true by creating or removing running Pods. There are many fundamental controllers and resources in Kubernetes that  work in this manner, including [Services](https://web.archive.org/web/20170129131616/http://kubernetes.io/docs/user-guide/services/ ), [Deployments](https://web.archive.org/web/20170129131616/http://kubernetes.io/docs/user-guide/deployments/), and [Daemon Sets](https://web. archive.org/web/20170129131616/http://kubernetes.io/docs/admin/daemons/).

Operator 建立在两个核心 Kubernetes 概念之上：资源和控制器。例如，内置的 [*ReplicaSet*](https://web.archive.org/web/20170129131616/http://kubernetes.io/docs/user-guide/replicasets/) 资源允许用户设置需要运行的 Pod 数量，Kubernetes 内部的控制器通过创建或删除正在运行的 Pod 来确保 ReplicaSet 资源中设置的所需状态保持为真。 Kubernetes 中有许多基本控制器和资源以这种方式工作，包括 [服务](https://web.archive.org/web/20170129131616/http://kubernetes.io/docs/user-guide/services/ )、[部署](https://web.archive.org/web/20170129131616/http://kubernetes.io/docs/user-guide/deployments/)和[守护进程集](https://web. archive.org/web/20170129131616/http://kubernetes.io/docs/admin/daemons/）。

[![img](https://web.archive.org/web/20170129131616im_/https://coreos.com/assets/blog/2016-11-03-introducing-operators/RS-before.svg)]( https://web.archive.org/web/20170129131616/https://coreos.com/assets/blog/2016-11-03-introducing-operators/RS-before.png)

Example 1a: A single pod is running, and the user updates the desired Pod count to 3.

示例 1a：单个 Pod 正在运行，并且用户将所需的 Pod 数量更新为 3。

[![img](https://web.archive.org/web/20170129131616im_/https://coreos.com/assets/blog/2016-11-03-introducing-operators/RS-scaled.svg)]( https://web.archive.org/web/20170129131616/https://coreos.com/assets/blog/2016-11-03-introducing-operators/RS-scaled.png)

Example 1b: A few moments later and controllers inside of Kubernetes have created new Pods to meet the user's request.

示例 1b：片刻之后，Kubernetes 内部的控制器创建了新的 Pod 来满足用户的请求。

An Operator builds upon the basic Kubernetes resource and controller  concepts and adds a set of knowledge or configuration that allows the  Operator to execute common application tasks. For example, when scaling  an etcd cluster manually, a user has to perform a number of steps:  create a DNS name for the new etcd member, launch the new etcd instance, and then use the etcd administrative tools (`etcdctl member add`) to tell the existing cluster about this new member. Instead with the *etcd Operator* a user can simply increase the etcd cluster size field by 1.

Operator 建立在基本的 Kubernetes 资源和控制器概念之上，并添加了一组允许 Operator 执行常见应用程序任务的知识或配置。例如，手动扩展 etcd 集群时，用户必须执行多个步骤：为新的 etcd 成员创建 DNS 名称，启动新的 etcd 实例，然后使用 etcd 管理工具（`etcdctl member add`）告诉现有集群这个新成员。相反，使用 *etcd Operator*，用户可以简单地将 etcd 集群大小字段增加 1。

[![img](https://web.archive.org/web/20170129131616im_/https://coreos.com/assets/blog/2016-11-03-introducing-operators/Operator-scale.svg)]( https://web.archive.org/web/20170129131616/https://coreos.com/assets/blog/2016-11-03-introducing-operators/Operator-scale.png)

Example 2: A backup is triggered by a user with kubectl

示例 2：备份由使用 kubectl 的用户触发

Other examples of complex administrative tasks that an Operator might handle include safe coordination of application upgrades, configuration of backups to offsite storage, service discovery via native Kubernetes  APIs, application TLS certificate configuration, and disaster recovery.

Operator 可能处理的复杂管理任务的其他示例包括应用程序升级的安全协调、异地存储的备份配置、通过本机 Kubernetes API 的服务发现、应用程序 TLS 证书配置和灾难恢复。

## How can you create an Operator?

## 如何创建 Operator？

Operators, by their nature, are application-specific, so the hard  work is going to be encoding all of the application operational domain  knowledge into a reasonable configuration resource and control loop. There are some common patterns that we have found while building  operators that we think are important for any application:

Operators，就其本质而言，是特定于应用程序的，因此艰苦的工作将是将所有应用程序操作领域知识编码到合理的配置资源和控制循环中。在构建我们认为对任何应用程序都很重要的 Operators 时，我们发现了一些常见的模式：

1. Operators should install as a single deployment e.g. `kubectl create -f https://coreos.com/operators/etcd/latest/deployment.yaml` and take no additional action once installed.
 2. Operators should create a new third party type when installed into  Kubernetes. A user will create new application instance using this type.
 3. Operators should leverage built-in Kubernetes primitives like  Services and Replica Sets when possible to leverage well-tested and  well-understood code.
 4. Operators should be backwards compatible and always understand previous versions of resources a user has created.
 5. Operators should be designed so application instances continue to run unaffected if the Operator is stopped or removed.
 6. Operators should give users the ability to declare a desired version  and orchestrate application upgrades based on the desired version. Not  upgrading software is a common source of operational bugs and security issues and Operators can  help users more confidently address this burden.
 7. Operators should be tested against a "Chaos Monkey" test suite that  simulates potential failures of Pods, configuration, and networking.



1. Operator 应安装为单个部署，例如`kubectl create -f https://coreos.com/operators/etcd/latest/deployment.yaml` 安装后不执行任何其他操作。
2. 安装到 Kubernetes 时，Operator 应该创建一个新的第三方类型。用户将使用此类型创建新的应用程序实例。
3. Operators应尽可能利用内置的 Kubernetes 原语，如服务和副本集，以利用经过良好测试和易于理解的代码。
4. Operators应该向后兼容，并始终了解用户创建的资源的先前版本。
5. Operators 的设计应确保如果 Operator 被停止或删除，应用程序实例将继续运行而不受影响。
6. Operators 应赋予用户声明所需版本并根据所需版本编排应用程序升级的能力。不升级软件是操作错误和安全问题的常见来源，Operators可以帮助用户更自信地解决这一负担。
7. Operator 应针对“Chaos Monkey”测试套件进行测试，该套件模拟 Pod、配置和网络的潜在故障。

## The Future of Operators

## Operators的未来

The etcd Operator and Prometheus Operator introduced by CoreOS today  showcase the power of the Kubernetes platform. For the last year, we  have worked alongside the wider Kubernetes community, laser-focused on  making Kubernetes stable, secure, easy to manage, and quick to install. 

CoreOS 今天推出的 etcd Operator 和 Prometheus Operator 展示了 Kubernetes 平台的强大功能。去年，我们与更广泛的 Kubernetes 社区合作，专注于使 Kubernetes 稳定、安全、易于管理和快速安装。

Now, as the foundation for Kubernetes has been laid, our new focus is the system to be built on top: software that extends Kubernetes with  new capabilities. We envision a future where users install Postgres  Operators, Cassandra Operators, or Redis Operators on their Kubernetes  clusters, and operate scalable instances of these programs as easily  they deploy replicas of their stateless web applications today.

现在，随着 Kubernetes 的基础已经奠定，我们的新重点是构建在其上的系统：使用新功能扩展 Kubernetes 的软件。我们设想了一个未来，用户可以在他们的 Kubernetes 集群上安装 Postgres Operators、Cassandra Operators 或 Redis Operators，并像他们今天部署无状态 Web 应用程序的副本一样轻松地操作这些程序的可扩展实例。

To learn more, dive into the GitHub repos, discuss on our [community](https://web.archive.org/web/20170129131616/https://coreos.com/community/) channels, or come talk with the CoreOS team [at KubeCon](https://web.archive.org/web/20170129131616/https://tectonic.com/blog/kubecon-preview.html) on Tuesday, November 8. Don't miss [my keynote on Tuesday, November 8 at 5:25 pm PT](https://web.archive.org/web/20170129131616/https://cnkc16.sched.org/event/8g4I), where I'll cover Operators and other Kubernetes topics.

要了解更多信息，请深入 GitHub 存储库，在我们的 [社区](https://web.archive.org/web/20170129131616/https://coreos.com/community/) 频道上讨论，或与 CoreOS 交谈团队 [at KubeCon](https://web.archive.org/web/20170129131616/https://tectonic.com/blog/kubecon-preview.html) 于 11 月 8 日星期二。不要错过 [我的主题演讲11 月 8 日，星期二，下午 5:25 PT](https://web.archive.org/web/20170129131616/https://cnkc16.sched.org/event/8g4I)，我将在这里介绍 Operator 和其他 Kubernetes 主题。

## FAQ

## 常问问题

**Q: How is this different than StatefulSets (previously PetSets)?**

**问：这与 StatefulSets（以前的 PetSets）有什么不同？**

A: StatefulSets are designed to enable support in Kubernetes for  applications that require the cluster to give them "stateful resources"  like static IPs and storage. Applications that need this more stateful  deployment model still need Operator automation to alert and act on  failure, backup, or reconfigure. So, an Operator for applications  needing these deployment properties could use StatefulSets instead of  leveraging ReplicaSets or Deployments.

答：StatefulSet 旨在支持 Kubernetes 中需要集群为其提供“有状态资源”（如静态 IP 和存储）的应用程序。需要这种更有状态的部署模型的应用程序仍然需要Operators自动化来警告和应对故障、备份或重新配置。因此，需要这些部署属性的应用程序的 Operator 可以使用 StatefulSets 而不是利用 ReplicaSets 或 Deployments。

**Q: How is this different from configuration management like Puppet or Chef?**

**问：这与 Puppet 或 Chef 等配置管理有何不同？**

A: Containers and Kubernetes are the big differentiation that make  Operators possible. With these two technologies deploying new software,  coordinating distributed configuration, and checking on multi-host  system state is consistent and easy using Kubernetes APIs. Operators  glue these primitives together in a useful way for application  consumers; it isn't just about configuration but the entire, live,  application state.

答：容器和 Kubernetes 是使 Operator 成为可能的最大区别。通过这两种技术，部署新软件、协调分布式配置和检查多主机系统状态是一致且易于使用 Kubernetes API 的。操作符以一种对应用程序消费者有用的方式将这些原语粘合在一起；它不仅与配置有关，而且与整个、实时的应用程序状态有关。

**Q: How is this different than Helm?**

**问：这与 Helm 有何不同？**

A: Helm is a tool for packaging multiple Kubernetes resources into a  single package. The concept of packaging up multiple applications  together and using Operators that actively manage applications are  complementary. For example, traefik is a load balancer that can use etcd as its backend database. You could create a Helm Chart that deploys a  traefik Deployment and etcd cluster instance together. The etcd cluster  would then be deployed and managed by the etcd Operator.

A：Helm 是一个将多个 Kubernetes 资源打包成一个包的工具。将多个应用程序打包在一起并使用 Operator 来主动管理应用程序的概念是互补的。例如，traefik 是一个负载均衡器，可以使用 etcd 作为其后端数据库。您可以创建一个 Helm Chart，将 traefik 部署和 etcd 集群实例部署在一起。然后 etcd 集群将由 etcd Operator 部署和管理。

**Q: What if someone is new to Kubernetes? What does this mean?**

**问：如果有人不熟悉 Kubernetes 怎么办？这是什么意思？**

A: This shouldn't change anything for new users except make it easier for them to deploy complex applications like etcd, Prometheus, and  others in the future. Our recommended onboarding path for Kubernetes is  still [minikube](https://web.archive.org/web/20170129131616/https://github.com/kubernetes/minikube), [kubectl run](https://web. archive.org/web/20170129131616/http://kubernetes.io/docs/user-guide/kubectl/kubectl_run/), and then maybe start playing with the Prometheus Operator to monitor the app you deployed with `kubectl run`.

答：这对新用户来说不会有任何改变，只是让他们在未来更容易部署复杂的应用程序，如 etcd、Prometheus 和其他应用程序。我们推荐的 Kubernetes 入门路径仍然是 [minikube](https://web.archive.org/web/20170129131616/https://github.com/kubernetes/minikube)、[kubectl run](https://web. archive.org/web/20170129131616/http://kubernetes.io/docs/user-guide/kubectl/kubectl_run/)，然后可能开始使用 Prometheus Operator 来监控您使用 `kubectl run` 部署的应用程序。

**Q: Is the code available for etcd Operator and Prometheus Operator today?**

**问：今天的 etcd Operator 和 Prometheus Operator 代码可用吗？**

A: Yes! They can be found on GitHub at [https://github.com/coreos/etcd-operator](https://web.archive.org/web/20170129131616/https://github.com/coreos/etcd-operator ) and [https://github.com/coreos/prometheus-operator](https://web.archive.org/web/20170129131616/https://github.com/coreos/prometheus-operator).

答：是的！它们可以在 GitHub 上找到 [https://github.com/coreos/etcd-operator](https://web.archive.org/web/20170129131616/https://github.com/coreos/etcd-operator ) 和 [https://github.com/coreos/prometheus-operator](https://web.archive.org/web/20170129131616/https://github.com/coreos/prometheus-operator)。

**Q: Do you have plans for other Operators?**

**问：你有其他Operators的计划吗？**

A: Yes, that is likely in the future. We would also love to see new  Operators get built by the community as well. Let us know what other  Operators you would like to see built next.

A：是的，这在未来很有可能。我们也很乐意看到社区也构建新的 Operator。让我们知道您接下来希望看到哪些其他 Operator。

**Q: How do Operators help secure a cluster?**

**问：Operators 如何帮助保护集群？**

A: Not upgrading software is a common source of operational bugs and  security issues and Operators can help users more confidently address  the burden of doing a correct upgrade.

答：不升级软件是操作错误和安全问题的常见来源，Operators可以帮助用户更自信地解决正确升级的负担。

**Q: Can Operators help with disaster recovery?**

**问：Operators可以帮助灾难恢复吗？**

A: Operators can make it easy to periodically back up application  state and recover previous state from the backup. A feature we hope will become common with Operators is easily enabling users to deploy new  instances from backups. 

A：Operators可以很容易地定期备份应用程序状态并从备份中恢复以前的状态。我们希望 Operators 能够普及的一项功能是让用户能够轻松地从备份部署新实例。
