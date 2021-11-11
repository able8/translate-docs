# To run or not to run a database on Kubernetes: What to consider

# 在 Kubernetes 上运行或不运行数据库：需要考虑的因素

Solutions Architect, Google Cloud

解决方案架构师，谷歌云

July 3, 2019 

#### Gartner Cloud DBMS MQ Report

#### Gartner Cloud DBMS MQ 报告

Learn why Google Cloud was named a leader in the market.

了解 Google Cloud 为何被评为市场领导者。

Today, more and more applications are being deployed in containers on Kubernetes—so much so that we’ve heard Kubernetes called the Linux of the cloud. Despite all that growth on the application layer, the data layer hasn’t gotten as much traction with containerization. That’s not surprising, since containerized workloads inherently have to be resilient to restarts, scale-out, virtualization, and other constraints. So handling things like state (the database), availability to other layers of the application, and redundancy for a database can have very specific requirements. That makes it challenging to run a database in a distributed environment.

今天，越来越多的应用程序被部署在 Kubernetes 上的容器中——以至于我们听说 Kubernetes 被称为云中的 Linux。尽管应用层有如此多的增长，但数据层并没有受到容器化的太大影响。这并不奇怪，因为容器化的工作负载本质上必须对重启、横向扩展、虚拟化和其他限制具有弹性。因此，处理诸如状态（数据库）、应用程序其他层的可用性以及数据库冗余之类的事情可能有非常具体的要求。这使得在分布式环境中运行数据库变得具有挑战性。

However, the data layer is getting more attention, since many developers want to treat data infrastructure the same as application stacks. Operators want to use the same tools for databases and applications, and get the same benefits as the application layer in the data layer: rapid spin-up and repeatability across environments. In this blog, we’ll explore when and what types of databases can be effectively run on Kubernetes.

然而，数据层越来越受到关注，因为许多开发人员希望将数据基础设施与应用程序堆栈一样对待。运营商希望对数据库和应用程序使用相同的工具，并在数据层获得与应用程序层相同的好处：快速启动和跨环境的可重复性。在这篇博客中，我们将探讨什么时候以及什么类型的数据库可以在 Kubernetes 上有效运行。

Before we dive into the considerations for running a database on Kubernetes, let's briefly review our options for running databases on [Google Cloud Platform](https://cloud.google.com/)(GCP) and what they're best used for .

在我们深入探讨在 Kubernetes 上运行数据库的注意事项之前，让我们简要回顾一下我们在 [Google Cloud Platform](https://cloud.google.com/)(GCP) 上运行数据库的选项以及它们的最佳用途.

- **Fully managed databases**. This includes [Cloud Spanner](https://cloud.google.com/spanner/), [Cloud Bigtable](https://cloud.google.com/bigtable/) and [Cloud SQL](https://cloud.google.com/sql/), among [others](https://cloud.google.com/products/databases/). This is the low-ops choice, since Google Cloud handles many of the maintenance tasks, like backups, patching and scaling. As a developer or operator, you don’t need to mess with them. You just create a database, build your app, and let Google Cloud scale it for you. This also means you might not have access to the exact version of a database, extension, or the exact flavor of database that you want.

- **完全托管的数据库**。这包括 [Cloud Spanner](https://cloud.google.com/spanner/)、[CloudBigtable](https://cloud.google.com/bigtable/) 和 [Cloud SQL](https://cloud.google.com/sql/)，以及 [其他](https://cloud.google.com/products/databases/)。这是低操作选择，因为 Google Cloud 处理许多维护任务，例如备份、修补和扩展。作为开发人员或运营商，您无需与他们混为一谈。您只需创建一个数据库，构建您的应用，然后让 Google Cloud 为您扩展它。这也意味着您可能无法访问您想要的数据库的确切版本、扩展或数据库的确切风格。

- **Do-it-yourself on a VM**. This might best be described as the full-ops option, where you take full responsibility for building your database, scaling it, managing reliability, setting up backups, and more. All of that can be a lot of work, but you have all the features and database flavors at your disposal.

- **在虚拟机上自己动手**。这可能最好被描述为 full-ops 选项，您全权负责构建数据库、扩展它、管理可靠性、设置备份等。所有这些都可能需要大量工作，但您可以使用所有功能和数据库风格。

- **Run it on Kubernetes**. Running a database on Kubernetes is closer to the full-ops option, but you do get some benefits in terms of the automation Kubernetes provides to keep the database application running. That said, it is important to remember that pods (the database application containers) are transient, so the likelihood of database application restarts or failovers is higher. Also, some of the more database-specific administrative tasks—backups, scaling, tuning, etc.—are different due to the added abstractions that come with containerization. 

- **在 Kubernetes 上运行它**。在 Kubernetes 上运行数据库更接近于 full-ops 选项，但您确实从 Kubernetes 提供的自动化方面获得了一些好处，以保持数据库应用程序运行。也就是说，重要的是要记住 pod（数据库应用程序容器）是瞬态的，因此数据库应用程序重新启动或故障转移的可能性更高。此外，一些特定于数据库的管理任务——备份、扩展、调整等——由于容器化带来的附加抽象而有所不同。

**Tips for running your database on Kubernetes** When choosing to go down the Kubernetes route, think about what database you will be running, and how well it will work given the trade-offs previously discussed. Since pods are mortal, the likelihood of failover events is higher than a traditionally hosted or fully managed database. It will be easier to run a database on Kubernetes if it includes concepts like sharding, failover elections and replication built into its DNA (for example, ElasticSearch, Cassandra, or MongoDB). Some open source projects provide [custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) and [operators](https://coreos.com/operators/) to help with managing the database.

**在 Kubernetes 上运行数据库的技巧** 在选择 Kubernetes 路线时，请考虑您将运行的数据库，以及考虑到前面讨论的权衡，它的工作情况。由于 Pod 是致命的，因此故障转移事件的可能性高于传统托管或完全托管的数据库。如果 Kubernetes 的 DNA 中包含分片、故障转移选举和复制等概念（例如 ElasticSearch、Cassandra 或 MongoDB），那么在 Kubernetes 上运行数据库会更容易。一些开源项目提供了 [自定义资源](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)和 [operators](https://coreos.com/operators/) 来帮助管理数据库。

Next, consider the function that database is performing in the context of your application and business. Databases that are storing more transient and caching layers are better fits for Kubernetes. Data layers of that type typically have more resilience built into the applications, making for a better overall experience.

接下来，考虑数据库在您的应用程序和业务上下文中执行的功能。存储更多瞬态和缓存层的数据库更适合 Kubernetes。这种类型的数据层通常具有更强的内置于应用程序的弹性，从而提供更好的整体体验。

Finally, be sure you understand the replication modes available in the database. Asynchronous modes of replication leave room for data loss, because transactions might be committed to the primary database but not to the secondary database(s). So, be sure to understand whether you might incur data loss, and how much of that is acceptable in the context of your application.

最后，确保您了解数据库中可用的复制模式。异步复制模式为数据丢失留下了空间，因为事务可能会提交到主数据库而不是辅助数据库。因此，请务必了解您是否可能会导致数据丢失，以及在您的应用程序上下文中，其中多少是可以接受的。

After evaluating all of those considerations, you’ll end up with a decision tree looking something like this:

在评估所有这些考虑因素后，您最终会得到一个如下所示的决策树：

![Tech Diag K8s Database Blog Flowchart.png](https://storage.googleapis.com/gweb-cloudblog-publish/images/Tech_Diag_K8s_Database_Bl.0734067713421317.max-1500x1500.png)

**How to deploy a database on Kubernetes** Now, let’s dive into more details on how to deploy a database on Kubernetes using StatefulSets. With a StatefulSet, your data can be stored on persistent volumes, decoupling the database application from the persistent storage, so when a pod (such as the database application) is recreated, all the data is still there. Additionally, when a pod is recreated in a StatefulSet, it keeps the same name, so you have a consistent endpoint to connect to. Persistent data and consistent naming are two of the largest benefits of StatefulSets. You can check out the Kubernetes [documentation](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/) for more details.

**如何在 Kubernetes 上部署数据库** 现在，让我们深入了解如何使用 StatefulSets 在 Kubernetes 上部署数据库的更多细节。使用 StatefulSet，您的数据可以存储在持久卷上，将数据库应用程序与持久存储解耦，因此当重新创建 pod（例如数据库应用程序）时，所有数据仍然存在。此外，当在 StatefulSet 中重新创建 pod 时，它会保留相同的名称，因此您可以连接到一致的端点。持久数据和一致命名是 StatefulSets 的两个最大好处。您可以查看 Kubernetes [文档](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/) 了解更多详细信息。

If you need to run a database that doesn’t perfectly fit the model of a Kubernetes-friendly database (such as MySQL or PostgreSQL), consider using Kubernetes Operators or projects that wrap those database with additional features. [Operators](https://coreos.com/operators/) will help you spin up those databases and perform database maintenance tasks like backups and replication. For MySQL in particular, take a look at the [Oracle MySQL Operator](https://github.com/oracle/mysql-operator) and [Crunchy Data](https://github.com/CrunchyData/postgres-operator) for PostgreSQL.

如果您需要运行的数据库不完全适合 Kubernetes 友好型数据库（例如 MySQL 或 PostgreSQL）的模型，请考虑使用 Kubernetes Operators 或使用附加功能包装这些数据库的项目。 [Operators](https://coreos.com/operators/) 将帮助您启动这些数据库并执行数据库维护任务，例如备份和复制。特别是对于 MySQL，请查看 [Oracle MySQL Operator](https://github.com/oracle/mysql-operator) 和 [Crunchy Data](https://github.com/CrunchyData/postgres-operator)对于 PostgreSQL。

Operators use [custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) and controllers to expose application-specific operations through the Kubernetes API. For example, to perform a backup using Crunchy Data, simply execute `pgo backup [cluster_name]`. To add a Postgres replica, use `pgo scale cluster [cluster_name]`.

Operator 使用 [自定义资源](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) 和控制器通过 Kubernetes API 公开特定于应用程序的操作。例如，要使用 Crunchy Data 执行备份，只需执行 `pgo backup [cluster_name]`。要添加 Postgres 副本，请使用 `pgo scale cluster [cluster_name]`。

There are some other projects out there that you might explore, such as [Patroni](https://github.com/zalando/patroni) for PostgreSQL. These projects use Operators, but go one step further. They’ve built many tools around their respective databases to aid their operation inside of Kubernetes. They may include additional features like sharding, leader election, and failover functionality needed to successfully deploy MySQL or PostgreSQL in Kubernetes.

您可以探索其他一些项目，例如用于 PostgreSQL 的 [Patroni](https://github.com/zalando/patroni)。这些项目使用 Operators，但更进一步。他们围绕各自的数据库构建了许多工具，以帮助他们在 Kubernetes 内部进行操作。它们可能包括在 Kubernetes 中成功部署 MySQL 或 PostgreSQL 所需的附加功能，如分片、领导者选举和故障转移功能。

While running a database in Kubernetes is gaining traction, it is still far from an exact science. There is a lot of work being done in this area, so keep an eye out as technologies and tools evolve toward making running databases in Kubernetes much more the norm. 

虽然在 Kubernetes 中运行数据库越来越受到关注，但它仍远非一门精确的科学。这方面有很多工作要做，因此请密切关注技术和工具的发展，使在 Kubernetes 中运行数据库更加规范。

When you're ready to get started, check out [GCP Marketplace](https://console.cloud.google.com/marketplace/browse?filter=category%3Adatabase&utm_source=blog&utm_medium=k8sdatabase) for easy-to-deploy SaaS, VM, and containerized database solutions and operators that can be deployed to GCP or Kubernetes clusters anywhere.

当您准备好开始时，请查看 [GCP Marketplace](https://console.cloud.google.com/marketplace/browse?filter=category%3Adatabase&utm_source=blog&utm_medium=k8sdatabase) 以获得易于部署的 SaaS，虚拟机，以及可部署到 GCP 或 Kubernetes 集群的容器化数据库解决方案和操作员。

[Learn More\
\
**How to connect GKE to Cloud SQL, including MySQL, PostgreSQL, and SQL Server** \
\
Check out the documentation for connecting your GKE-based app to Cloud SQL.\
\
Read Article](https://cloud.google.com/sql/docs/mysql/connect-kubernetes-engine) 

[了解更多\
\
**如何将 GKE 连接到 Cloud SQL，包括 MySQL、PostgreSQL 和 SQL Server** \
\
查看将基于 GKE 的应用连接到 Cloud SQL 的文档。\
\
阅读文章](https://cloud.google.com/sql/docs/mysql/connect-kubernetes-engine)

