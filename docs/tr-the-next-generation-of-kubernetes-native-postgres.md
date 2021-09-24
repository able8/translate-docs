# The Next Generation of Kubernetes Native Postgres

# 下一代 Kubernetes Native Postgres

July 07, 2021 [Jonathan S. Katz](https://blog.crunchydata.com/blog/author/jonathan-s-katz)

2021 年 7 月 7 日 [Jonathan S. Katz](https://blog.crunchydata.com/blog/author/jonathan-s-katz)

We're excited to announce the release of [PGO](https://github.com/CrunchyData/postgres-operator) 5.0, the open source [Postgres Operator](https://github.com/CrunchyData/postgres-operator) from [Crunchy Data](https://www.crunchydata.com/). While I'm very excited for you to [try out PGO 5.0](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/) and provide feedback, I also want to provide some background on this release.

我们很高兴地宣布发布 [PGO](https://github.com/CrunchyData/postgres-operator) 5.0，开源 [Postgres Operator](https://github.com/CrunchyData/postgres-operator) 来自 [Crunchy Data](https://www.crunchydata.com/)。虽然我很高兴您[试用 PGO 5.0](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/) 并提供反馈，但我也想提供一些关于此的背景释放。

When I joined Crunchy Data back in 2018, I had heard of [Kubernetes](https://kubernetes.io/) through my various open source activities, but I did not know much about it. I learned that we had been running Postgres on [Kubernetes](https://blog.crunchydata.com/blog/creating-a-postgresql-cluster-using-helm-for-kubernetes) and [OpenShift](http://blog.crunchydata.com/blog/advanced-crunchy-containers-for-postgresql) in production environments for years. This included the release of [one of the first Kubernetes Operators](http://blog.crunchydata.com/blog/postgres-operator-for-kubernetes)! It was quite remarkable to see how Crunchy lead the way in cloud native Postgres. I still remember how excited I was when I [got my first Postgres cluster up and running on Kubernetes](http://blog.crunchydata.com/blog/get-started-runnning-postgresql-on-kubernetes)!

早在 2018 年加入 Crunchy Data 时，我就通过我的各种开源活动听说过 [Kubernetes](https://kubernetes.io/)，但我对此并不了解。我了解到我们一直在[Kubernetes](https://blog.crunchydata.com/blog/creating-a-postgresql-cluster-using-helm-for-kubernetes) 和 [OpenShift](http://blog.crunchydata.com/blog/advanced-crunchy-containers-for-postgresql)在生产环境中使用多年。这包括发布 [首批 Kubernetes 运营商之一](http://blog.crunchydata.com/blog/postgres-operator-for-kubernetes)！看到 Crunchy 如何在云原生 Postgres 中处于领先地位，这是非常了不起的。我仍然记得当我 [在 Kubernetes 上启动并运行我的第一个 Postgres 集群](http://blog.crunchydata.com/blog/get-started-runnning-postgresql-on-kubernetes) 时我是多么兴奋！

Many things have changed in the cloud native world over the past three years. When I first started giving talks on the topic, I was answering questions like, "Can I run a database in a container?" coupled with "Should I run a database in a container?" The conversation has now shifted. My colleague Paul Laurence wrote an [excellent article](http://blog.crunchydata.com/blog/using-kubernetes-chances-are-you-need-a-database) capturing the current discourse. The question is no longer, "Should I run a database in Kubernetes" but " [Which database should I run in Kubernetes](http://blog.crunchydata.com/blog/using-kubernetes-chances-are-you-need-a-database)?" (Postgres!).

在过去三年中，云原生世界发生了许多变化。当我第一次开始就这个话题发表演讲时，我正在回答诸如“我可以在容器中运行数据库吗？”之类的问题。再加上“我应该在容器中运行数据库吗？”现在谈话已经转移了。我的同事 Paul Laurence 写了一篇 [优秀文章](http://blog.crunchydata.com/blog/using-kubernetes-chances-are-you-need-a-database) 捕捉了当前的话语。问题不再是“我应该在 Kubernetes 中运行数据库吗”而是“[我应该在 Kubernetes 中运行哪个数据库](http://blog.crunchydata.com/blog/using-kubernetes-chances-are-you-need -a-数据库）？” （Postgres！)。

Along with this shift in discussion is a shift in expectation for how databases should work on Kubernetes. To do that, we need to understand the difference between an imperative workflow and a declarative workflow.

随着讨论的这种转变，对数据库应该如何在 Kubernetes 上工作的期望也发生了转变。为此，我们需要了解命令式工作流和声明式工作流之间的区别。

## Cloud Native Declarative Postgres

## 云原生声明性 Postgres

[PGO](https://github.com/CrunchyData/postgres-operator), the open source [Postgres Operator from Crunchy Data](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes/), was initially designed for running cloud native Postgres using an imperative workflow. Many operations required the use of a command-line utility called "pgo". For convenience, "pgo" follows  the conventions of the Kubernetes " [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/)" command line tool. "pgo" includes many conveniences for managing Postgres on Kubernetes, including:

[PGO](https://github.com/CrunchyData/postgres-operator)，开源的 [来自 Crunchy Data 的 Postgres Operator](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes/)，最初设计用于使用命令式工作流运行云原生 Postgres。许多操作需要使用名为“pgo”的命令行实用程序。为方便起见，“pgo”遵循 Kubernetes“[kubectl](https://kubernetes.io/docs/reference/kubectl/overview/)”命令行工具的约定。 “pgo”包括许多在 Kubernetes 上管理 Postgres 的便利，包括：

- Creating databases:pgo create cluster hippo
- Maintenance operations:pgo update cluster hippo --memory=2Gi
- Backups:pgo backup hippo
- Cloning clusters:pgo create cluster rhino --restore-from=hippo

- 创建数据库：pgo 创建集群河马
- 维护操作：pgo update cluster hippo --memory=2Gi
- 备份：pgo 备份河马
- 克隆集群：pgo create cluster rhino --restore-from=hippo

and more.

和更多。

A declarative workflow is one where you describe what you want, and your application "makes it happen." An example of this is SQL: you write a query, e.g.:

声明式工作流是您描述您想要什么，并且您的应用程序“让它发生”的工作流。这方面的一个例子是 SQL：您编写一个查询，例如：

```
<span>SELECT</span> <span>*</span> <span>FROM</span> animals <span>WHERE</span> animal_type <span>=</span> <span>'hippo'</span>;
```

You are asking your database "Find me all the animals that are hippos". You don't care how the database actually does this: it could search over an index or perform a sequential scan. You know that for your end result, you want all the hippos.

你在问你的数据库“找到我所有的河马动物”。您并不关心数据库实际上是如何执行此操作的：它可以搜索索引或执行顺序扫描。你知道为了你的最终结果，你想要所有的河马。

This is a very powerful concept: instead of building towards your end result, you describe what you want it to be. A well-designed declarative engine can optimize your experience using the software.

这是一个非常强大的概念：不是朝着最终结果构建，而是描述你想要的结果。精心设计的声明式引擎可以优化您使用软件的体验。

This was the founding principle on which we built the new version of our Kubernetes Operator for Postgres.

这是我们为 Postgres 构建新版本 Kubernetes Operator 的基本原则。

## Designing the Next Generation Postgres Operator 

## 设计下一代 Postgres 运算符

We had several goals while designing the next generation of [Postgres Operator](https://github.com/CrunchyData/postgres-operator). We wanted to make it both [easy to get started running Postgres on Kubernetes with PGO](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/) and easy to manage "Day 2" type operations like cluster resizing. We also wanted to ensure our Postgres databases could withstand the rigors of running in a Kubernetes environment.

在设计下一代 [Postgres Operator](https://github.com/CrunchyData/postgres-operator) 时，我们有几个目标。我们想让它[轻松开始使用 PGO 在 Kubernetes 上运行 Postgres](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/) 和易于管理的“第 2 天”类型的操作像集群调整大小。我们还希望确保我们的 Postgres 数据库能够承受在 Kubernetes 环境中运行的严苛考验。

While building the new PGO, we focused on creating a seamless, declarative user experience that included the production-ready Postgres features introduced in prior releases. We also leveraged modern Kubernetes features like [server-side apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/). We architected PGO so that it works behind the scenes maintaining the Postgres environment that you requested.

在构建新的 PGO 时，我们专注于创建无缝的、声明性的用户体验，其中包括先前版本中引入的生产就绪 Postgres 功能。我们还利用了现代 Kubernetes 功能，例如 [服务器端应用](https://kubernetes.io/docs/reference/using-api/server-side-apply/)。我们对 PGO 进行了架构设计，以便它在维护您请求的 Postgres 环境的幕后工作。

For example, let's say you accidentally delete your connection pooler Deployment. Before this release, [pgMonitor](https://github.com/CrunchyData/pgmonitor) would send you alerts indicating your database is disconnected. After investigating, you realize that you need to manually recreate your pgBouncer deployment. With the updated declarative approach, PGO now automatically heals your environment. When PGO detects the missing Deployment, it will instantaneously recreate your connection pooler and make it seem like nothing happened!

例如，假设您不小心删除了连接池部署。在此版本之前，[pgMonitor](https://github.com/CrunchyData/pgmonitor) 会向您发送警报，指示您的数据库已断开连接。调查后，您意识到需要手动重新创建 pgBouncer 部署。通过更新的声明式方法，PGO 现在可以自动修复您的环境。当 PGO 检测到丢失的 Deployment 时，它会立即重新创建您的连接池，并使它看起来好像什么也没发生！

Think of it as the Postgres Operator automatically managing your operations for you.

可以将其视为 Postgres Operator 自动为您管理操作。

## Deploying Postgres the GitOps Way

## 以 GitOps 方式部署 Postgres

There is an ongoing push to use [infrastructure as code](https://en.wikipedia.org/wiki/Infrastructure_as_code) ([IAC](https://en.wikipedia.org/wiki/Infrastructure_as_code)) methodologies such as [GitOps](https://www.gitops.tech/) with Kubernetes. Because of this, we wanted PGO to work well with tools such as [Kustomize](https://kustomize.io/),[Helm](https://helm.sh/), and [OLM](https://olm.operatorframework.io/) throughout the lifetime of a Postgres cluster.

一直在推动使用 [基础设施即代码](https://en.wikipedia.org/wiki/Infrastructure_as_code)([IAC](https://en.wikipedia.org/wiki/Infrastructure_as_code)) 方法，例如[GitOps](https://www.gitops.tech/) 与 Kubernetes。正因为如此，我们希望 PGO 能够与 [Kustomize](https://kustomize.io/)、[Helm](https://helm.sh/) 和 [OLM](https://olm.operatorframework.io/) 贯穿 Postgres 集群的整个生命周期。

We took Kelsey Hightower's " [Stop scripting and start shipping](https://twitter.com/kelseyhightower/status/953638870888849408)" to heart. We made a simple, one-step operation to create your Postgres database and connect your application in a safe and secure way. We also allowed for everything to be modifiable through a declarative workflow. We fine-tuned the experience by running lots of tests through [ArgoCD](https://argoproj.github.io/argo-cd/). We also created a series of [examples for deploying Postgres clusters in various Kubernetes scenarios](https://github.com/CrunchyData/postgres-operator-examples).

我们将 Kelsey Hightower 的“[停止编写脚本并开始发货](https://twitter.com/kelseyhightower/status/953638870888849408)”铭记于心。我们做了一个简单的一步操作来创建您的 Postgres 数据库并以安全可靠的方式连接您的应用程序。我们还允许通过声明性工作流程修改所有内容。我们通过 [ArgoCD](https://argoproj.github.io/argo-cd/) 运行大量测试来微调体验。我们还创建了一系列 [在各种 Kubernetes 场景中部署 Postgres 集群的示例](https://github.com/CrunchyData/postgres-operator-examples)。

Part of managing your infrastructure using GitOps involves deploying applications with minimal disruptions. In Kubernetes, there are certain modifications that can cause downtime. For example, updating to a newer Postgres bugfix release causes Kubernetes to create new Pods. We designed PGO to make sure that when these necessary Day 2 actions occur, they are performed with minimal to zero downtime.

使用 GitOps 管理基础设施的一部分涉及以最小的中断部署应用程序。在 Kubernetes 中，某些修改可能会导致停机。例如，更新到较新的 Postgres 错误修复版本会导致 Kubernetes 创建新的 Pod。我们设计了 PGO 以确保当这些必要的第 2 天操作发生时，它们在最少到零停机时间的情况下执行。

## The Next Generation Cloud Native Postgres is Now

## 下一代云原生 Postgres 现已发布

I've been around open source for a long time, and through the years I've seen how open source innovation begets innovation. A decade ago, Postgres [added support for the JSON data type](http://blog.crunchydata.com/blog/better-json-in-postgres-with-postgresql-14). Adding JSON spurred many advances throughout the application development and relational database landscape!

我已经接触开源很长时间了，多年来我看到了开源创新是如何引发创新的。十年前，Postgres [增加了对 JSON 数据类型的支持](http://blog.crunchydata.com/blog/better-json-in-postgres-with-postgresql-14)。添加 JSON 推动了整个应用程序开发和关系数据库领域的许多进步！

I'm seeing something similar now occurring with Kubernetes. Our team built a powerful, declarative Postgres Operator thanks to advances in the core Kubernetes ecosystem. We can now run production cloud native Postgres by maintaining a few YAML files.

我现在看到 Kubernetes 发生了类似的事情。由于核心 Kubernetes 生态系统的进步，我们的团队构建了一个强大的、声明式的 Postgres Operator。我们现在可以通过维护一些 YAML 文件来运行生产云原生 Postgres。

Even with all these advances in cloud native technology, what excites me is that we're only starting to scratch the surface of what we can do. The good news is that the next generation of cloud native Postgres is here, and it's [open source](https://github.com/CrunchyData/postgres-operator). 

尽管云原生技术取得了所有这些进步，但让我兴奋的是我们才刚刚开始触及我们可以做的事情的皮毛。好消息是下一代云原生 Postgres 来了，它是[开源](https://github.com/CrunchyData/postgres-operator)。

[Crunchy Postgres for Kubernetes](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes/) 5.0 is built on the redesigned [PGO](https://github.com/CrunchyData/postgres-operator), the open source [Postgres Operator](https://github.com/CrunchyData/postgres-operator). PGO 5.0 combines many years experience running Postgres on Kubernetes with a robust operations feature set to deliver a modern cloud native Postgres distribution. You can check out our [full press release](http://blog.crunchydata.com/news/next-generation-crunchy-postgres-for-kubernetes-released) that also includes info on our PGO 5.0 webinar.

[Crunchy Postgres for Kubernetes](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes/) 5.0 建立在重新设计的 [PGO](https://github.com/CrunchyData/postgres-operator)，开源 [Postgres Operator](https://github.com/CrunchyData/postgres-operator)。 PGO 5.0 将多年在 Kubernetes 上运行 Postgres 的经验与强大的操作功能集相结合，以提供现代云原生 Postgres 发行版。您可以查看我们的 [完整新闻稿](http://blog.crunchydata.com/news/next-generation-crunchy-postgres-for-kubernetes-released)，其中还包括我们 PGO 5.0 网络研讨会的信息。

We cordially invite you to [try out PGO](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/), take it for a test drive, and let us know what you think. 

我们诚挚地邀请您[试用 PGO](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/)，试用它，并告诉我们您的想法。

