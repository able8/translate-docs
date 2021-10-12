# To run or not to run a database on Kubernetes: What to consider

Solutions Architect, Google Cloud

July 3, 2019

#### Gartner Cloud DBMS MQ Report

Learn why Google Cloud was named a leader in the market.

Today, more and more applications are being deployed in containers on Kubernetes—so much so that we’ve heard Kubernetes called the Linux of the cloud. Despite all that growth on the application layer, the data layer hasn’t gotten as much traction with containerization. That’s not surprising, since containerized workloads inherently have to be resilient to restarts, scale-out, virtualization, and other constraints. So handling things like state (the database), availability to other layers of the application, and redundancy for a database can have very specific requirements. That makes it challenging to run a database in a distributed environment.

However, the data layer is getting more attention, since many developers want to treat data infrastructure the same as application stacks. Operators want to use the same tools for databases and applications, and get the same benefits as the application layer in the data layer: rapid spin-up and repeatability across environments. In this blog, we’ll explore when and what types of databases can be effectively run on Kubernetes.

Before we dive into the considerations for running a database on Kubernetes, let’s briefly review our options for running databases on [Google Cloud Platform](https://cloud.google.com/) (GCP) and what they’re best used for.

- **Fully managed databases**. This includes [Cloud Spanner](https://cloud.google.com/spanner/), [Cloud Bigtable](https://cloud.google.com/bigtable/) and [Cloud SQL](https://cloud.google.com/sql/), among [others](https://cloud.google.com/products/databases/). This is the low-ops choice, since Google Cloud handles many of the maintenance tasks, like backups, patching and scaling. As a developer or operator, you don’t need to mess with them. You just create a database, build your app, and let Google Cloud scale it for you. This also means you might not have access to the exact version of a database, extension, or the exact flavor of database that you want.

- **Do-it-yourself on a VM**. This might best be described as the full-ops option, where you take full responsibility for building your database, scaling it, managing reliability, setting up backups, and more. All of that can be a lot of work, but you have all the features and database flavors at your disposal.

- **Run it on Kubernetes**. Running a database on Kubernetes is closer to the full-ops option, but you do get some benefits in terms of the automation Kubernetes provides to keep the database application running. That said, it is important to remember that pods (the database application containers) are transient, so the likelihood of database application restarts or failovers is higher. Also, some of the more database-specific administrative tasks—backups, scaling, tuning, etc.—are different due to the added abstractions that come with containerization.


**Tips for running your database on Kubernetes** When choosing to go down the Kubernetes route, think about what database you will be running, and how well it will work given the trade-offs previously discussed. Since pods are mortal, the likelihood of failover events is higher than a traditionally hosted or fully managed database. It will be easier to run a database on Kubernetes if it includes concepts like sharding, failover elections and replication built into its DNA (for example, ElasticSearch, Cassandra, or MongoDB). Some open source projects provide [custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) and [operators](https://coreos.com/operators/) to help with managing the database.

Next, consider the function that database is performing in the context of your application and business. Databases that are storing more transient and caching layers are better fits for Kubernetes. Data layers of that type typically have more resilience built into the applications, making for a better overall experience.

Finally, be sure you understand the replication modes available in the database. Asynchronous modes of replication leave room for data loss, because transactions might be committed to the primary database but not to the secondary database(s). So, be sure to understand whether you might incur data loss, and how much of that is acceptable in the context of your application.

After evaluating all of those considerations, you’ll end up with a decision tree looking something like this:

![Tech Diag K8s Database Blog Flowchart.png](https://storage.googleapis.com/gweb-cloudblog-publish/images/Tech_Diag_K8s_Database_Bl.0734067713421317.max-1500x1500.png)

**How to deploy a database on Kubernetes** Now, let’s dive into more details on how to deploy a database on Kubernetes using StatefulSets. With a StatefulSet, your data can be stored on persistent volumes, decoupling the database application from the persistent storage, so when a pod (such as the database application) is recreated, all the data is still there. Additionally, when a pod is recreated in a StatefulSet, it keeps the same name, so you have a consistent endpoint to connect to. Persistent data and consistent naming are two of the largest benefits of StatefulSets. You can check out the Kubernetes [documentation](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/) for more details.

If you need to run a database that doesn’t perfectly fit the model of a Kubernetes-friendly database (such as MySQL or PostgreSQL), consider using Kubernetes Operators or projects that wrap those database with additional features. [Operators](https://coreos.com/operators/) will help you spin up those databases and perform database maintenance tasks like backups and replication. For MySQL in particular, take a look at the [Oracle MySQL Operator](https://github.com/oracle/mysql-operator) and [Crunchy Data](https://github.com/CrunchyData/postgres-operator) for PostgreSQL.

Operators use [custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) and controllers to expose application-specific operations through the Kubernetes API. For example, to perform a backup using Crunchy Data, simply execute `pgo backup [cluster_name]`. To add a Postgres replica, use `pgo scale cluster [cluster_name]`.

There are some other projects out there that you might explore, such as [Patroni](https://github.com/zalando/patroni) for PostgreSQL. These projects use Operators, but go one step further. They’ve built many tools around their respective databases to aid their operation inside of Kubernetes. They may include additional features like sharding, leader election, and failover functionality needed to successfully deploy MySQL or PostgreSQL in Kubernetes.

While running a database in Kubernetes is gaining traction, it is still far from an exact science. There is a lot of work being done in this area, so keep an eye out as technologies and tools evolve toward making running databases in Kubernetes much more the norm.

When you’re ready to get started, check out [GCP Marketplace](https://console.cloud.google.com/marketplace/browse?filter=category%3Adatabase&utm_source=blog&utm_medium=k8sdatabase) for easy-to-deploy SaaS, VM, and containerized database solutions and operators that can be deployed to GCP or Kubernetes clusters anywhere.

[Learn More\
\
**How to connect GKE to Cloud SQL, including MySQL, PostgreSQL, and SQL Server** \
\
Check out the documentation for connecting your GKE-based app to Cloud SQL.\
\
Read Article](https://cloud.google.com/sql/docs/mysql/connect-kubernetes-engine)
