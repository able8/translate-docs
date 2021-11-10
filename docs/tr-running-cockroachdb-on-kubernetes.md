# Running CockroachDB on Kubernetes

# 在 Kubernetes 上运行 CockroachDB

- on Mar 3, 2021

  


![Running CockroachDB on Kubernetes](https://crl2020.imgix.net/img/crl-solutions-diagram-kubernetes_1020x475@2X.jpg?auto=format,compress&q=60&w=1185)

_Since this post was originally published in 2017, StatefulSets have become common and allow a wide array of stateful workloads to run on Kubernetes. In this post,  we'll quickly walk through the history of StatefulSets, and how they fit with [CockroachDB and Kubernetes](https://www.cockroachlabs.com/product/kubernetes/), before jumping into a tutorial for running CockroachDB on Kubernetes._

_自从这篇文章最初于 2017 年发布以来，StatefulSets 已经变得很普遍，并允许在 Kubernetes 上运行各种有状态的工作负载。在这篇文章中，我们将快速浏览 StatefulSets 的历史，以及它们如何与 [CockroachDB 和 Kubernetes](https://www.cockroachlabs.com/product/kubernetes/) 配合，然后再进入运行 CockroachDB 的教程在 Kubernetes 上。_

Managing resilience, scale, and ease of operations in a containerized world is largely what Kubernetes is all about—and one of the reasons platform adoption has doubled since 2017.  And as container orchestration continues to become a dominant DevOps paradigm, the ecosystem has continued to mature with better tools for replication, management, and monitoring of our workloads.

在容器化世界中管理弹性、规模和易于操作在很大程度上是 Kubernetes 的全部——这也是自 2017 年以来平台采用率翻番的原因之一。随着容器编排继续成为主要的 DevOps 范式，生态系统也在继续使用更好的工具来复制、管理和监控我们的工作负载。

And as Kubernetes grows, so does CockroachDB as we’ve recently simplified some of the day 2 operations associated with our distributed database with our Kubernetes Operator. Ultimately, however, our overall goal in the cloud-native community is singular: ease the deployment of stateful workloads on Kubernetes.

随着 Kubernetes 的发展，CockroachDB 也在发展，因为我们最近使用 Kubernetes Operator 简化了与分布式数据库相关的一些第 2 天操作。然而，最终，我们在云原生社区的总体目标是单一的：简化有状态工作负载在 Kubernetes 上的部署。

## Bringing State to Kubernetes

## 将状态带入 Kubernetes

CockroachDB helps solve for stateful, database-dependent applications through replication of data across independent database nodes in a way that will survive any failure (just in case our name didn’t make total sense). CockroachDB, combined with Kubernetes’ built-in scale out, survivability and replication strategies, can give you the speed and simplicity of orchestration without sacrificing the high availability and correctness you expect from critical stateful databases.

CockroachDB 通过跨独立数据库节点复制数据来帮助解决有状态的、依赖于数据库的应用程序，这种方式可以在任何故障中幸存下来（以防万一我们的名字没有完全意义）。 CockroachDB 与 Kubernetes 的内置横向扩展、生存能力和复制策略相结合，可以为您提供编排的速度和简单性，而不会牺牲您对关键状态数据库的高可用性和正确性的期望。

## How do CockroachDB + Kubernetes = Retained State?

## CockroachDB + Kubernetes = 保留状态如何？

While Kubernetes is fairly straightforward for use with stateless services, management and surviving state has been a challenge.

虽然 Kubernetes 与无状态服务一起使用相当简单，但管理和生存状态一直是一个挑战。

Why? You can’t simply swap out nodes as they depend on data in pod-mounted storage. And rolling back doesn’t work for databases either.

为什么？您不能简单地换出节点，因为它们依赖于安装在 pod 上的存储中的数据。回滚也不适用于数据库。

Some best practices have evolved to workaround the challenge of deploying data-driven apps on K8s:

一些最佳实践已经演变为解决在 K8s 上部署数据驱动应用程序的挑战：

- **Run the database outside of Kubernetes:** This creates lots of extra work, adds redundant tooling, and can actually invalidate the value you were looking to gain from K8s.
- **Use a DBaaS:** This, too, limits the value of scale and resilience of Kubernetes and limits your choice to only those options provided by your Cloud Service Provider

- **在 Kubernetes 之外运行数据库：** 这会产生大量额外工作，添加冗余工具，并且实际上可能会使您希望从 K8s 中获得的价值失效。
- **使用 DBaaS：** 这也限制了 Kubernetes 的规模和弹性的价值，并将您的选择限制为仅由您的云服务提供商提供的选项

To keep up with the demands of modern, data-driven apps, the Kubernetes community developed  a native way to manage state, via StatefulSets.

为了跟上现代数据驱动应用程序的需求，Kubernetes 社区开发了一种通过 StatefulSets 管理状态的本机方式。

- StatefulSets assign a unique ID that keeps application and database containers connected through automation.
- _Note: The use of “Unique ID” is a bit tricky here. Each resource in kubernetes gets UID to identify it but that UID will change whenever the resource is updated. [StatefulSets assign pod identity](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-identity) which is persistent across pod generations but is separate from UIDs._ 

- StatefulSets 分配一个唯一的 ID，使应用程序和数据库容器通过自动化保持连接。
- _注意：这里使用“唯一ID”有点棘手。 kubernetes 中的每个资源都会获得 UID 来识别它，但是只要资源更新，UID 就会改变。 [StatefulSets 分配 pod 身份](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-identity)，它在 pod 代中是持久的，但与 UID 分开。_

StatefulSets are ideal for CockroachDB because the UID means it doesn’t get treated as a new node in a Kubernetes cluster, cutting way back on the amount of data replication required to keep data available. This is key to efficiently supporting [fast distributed transactions](https://www.cockroachlabs.com/docs/stable/architecture/life-of-a-distributed-transaction.html) and our [consensus protocol](https://www.cockroachlabs.com/docs/stable/architecture/replication-layer.html). For a real life example of a CockroachDB running on Kubernetes to retain state check out this [Pure Storage case study](https://resources.cockroachlabs.com/case-study/pure-storage-pso-case-study).

StatefulSets 是 CockroachDB 的理想选择，因为 UID 意味着它不会被视为 Kubernetes 集群中的新节点，从而减少了保持数据可用所需的数据复制量。这是有效支持[快速分布式事务](https://www.cockroachlabs.com/docs/stable/architecture/life-of-a-distributed-transaction.html)和我们的[共识协议](https://www.cockroachlabs.com/docs/stable/architecture/replication-layer.html)。有关在 Kubernetes 上运行的 CockroachDB 以保留状态的真实示例，请查看此 [Pure Storage 案例研究](https://resources.cockroachlabs.com/case-study/pure-storage-pso-case-study)。

## Step By Step Kubernetes Tutorial

## Kubernetes 分步教程

Step One: Building Your Kubernetes Cluster

第一步：构建您的 Kubernetes 集群

The year is 2021.  There are lots of ways to get your Kubernetes cluster up and running. For this walkthrough we’ll use [GKE](https://cloud.google.com/kubernetes-engine). If you’re interested in other paths, we have resources for:

今年是 2021 年。有很多方法可以让您的 Kubernetes 集群启动并运行。在本演练中，我们将使用 [GKE](https://cloud.google.com/kubernetes-engine)。如果您对其他路径感兴趣，我们有以下资源：

- Running[a cluster locally on Minikube](https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html)
- Setting a[cluster up on AWS](http://kubernetes.io/docs/stable/getting-started-guides/kops/)

- 运行[在 Minikube 本地集群](https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html)
- 在 AWS 上设置 [集群](http://kubernetes.io/docs/stable/getting-started-guides/kops/)

With the Google Cloud CLI installed, create the cluster by running

安装 Google Cloud CLI 后，通过运行创建集群

`gcloud container clusters create cockroachdb-cluster`

`gcloud 容器集群创建 cockroachdb-cluster`

### Step Two: Spinning up CockroachDB

### 第二步：启动 CockroachDB

Just like most Kubernetes deployment configuration work, CockroachDB config is managed by a YAML file like the one below. We’ve added comments to help provide some context for what’s going on.

就像大多数 Kubernetes 部署配置工作一样，CockroachDB 配置由一个 YAML 文件管理，如下所示。我们添加了评论，以帮助为正在发生的事情提供一些背景信息。

- Start by copying the file from our[GitHub repository](https://github.com/cockroachdb/cockroach/blob/master/cloud/kubernetes/cockroachdb-statefulset.yaml) into a file named 'cockroachdb-statefulset.yaml' . This file defines the resources to be created, in this case including the StatefulSet object that will spin up the CockroachDB containers and then attach them to persistent volumes.
- You’ll then need to create the resources shown below. If you're using Minikube, you may need[to manually provision persistent volumes](https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html#step-2-start-the-cockroachdb-cluster).

- 首先将文件从我们的 [GitHub 存储库](https://github.com/cockroachdb/cockroach/blob/master/cloud/kubernetes/cockroachdb-statefulset.yaml) 复制到名为“cockroachdb-statefulset.yaml”的文件中.该文件定义了要创建的资源，在本例中包括 StatefulSet 对象，该对象将启动 CockroachDB 容器，然后将它们附加到持久卷。
- 然后您需要创建如下所示的资源。如果您使用 Minikube，您可能需要[手动配置持久卷](https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html#step-2-start-the- cockroachdb 集群)。

You should soon see 3 replicas running in your cluster along with a couple of services. At first, only some of the replicas may show because they haven’t all yet started. This is normal, as StatefulSets create the replicas one-by-one, starting with the first.

您应该很快就会看到 3 个副本以及几个服务在您的集群中运行。起初，可能只显示部分副本，因为它们尚未全部启动。这是正常的，因为 StatefulSet 从第一个开始一个一个地创建副本。

```
$ kubectl create -f cockroachdb-statefulset.yaml

service "cockroachdb-public" created

service "cockroachdb" created

poddisruptionbudget "cockroachdb-budget" created

statefulset "cockroachdb" created

$ kubectl get services

cockroachdb          None         <none>        26257/TCP,8080/TCP   4s

cockroachdb-public   10.0.0.85    <none>        26257/TCP,8080/TCP   4s

kubernetes 10.0.0.1 <none> 443/TCP 1h

$ kubectl get pods

NAME            READY     STATUS    RESTARTS   AGE

cockroachdb-0   1/1       Running   0          29s

cockroachdb-1   0/1       Running   0          9s

$ kubectl get pods

NAME            READY     STATUS    RESTARTS   AGE

cockroachdb-0   1/1       Running   0          1m

cockroachdb-1   1/1       Running   0          41s

cockroachdb-2   1/1       Running   0          21s

```

#### Lifting the Hood

#### 掀开引擎盖

If you’re curious to see what’s happening inside the cluster, check the logs for one of the pods by running `kubectl logs cockroachdb-0`.

如果您想知道集群内部发生了什么，请通过运行 `kubectl logs cockroachdb-0` 检查其中一个 Pod 的日志。

### Step Three: Using the CockroachDB cluster

### 第三步：使用 CockroachDB 集群

If all has gone to plan, you now have a cluster up and running. Congratulations!

如果一切都按计划进行，您现在已经启动并运行了一个集群。恭喜！

To open a SQL shell within the Kubernetes cluster, you can run a one-off interactive pod like this, using the `cockroachdb-public` hostname to access the CockroachDB cluster. Kubernetes will then automatically load-balance connections to that hostname across the healthy CockroachDB instances.

要在 Kubernetes 集群中打开 SQL shell，您可以像这样运行一次性交互式 pod，使用 `cockroachdb-public` 主机名访问 CockroachDB 集群。然后，Kubernetes 将自动在健康的 CockroachDB 实例中对与该主机名的连接进行负载平衡。

```
$ kubectl run cockroachdb -it --image=cockroachdb/cockroach --rm --restart=Never -- sql --insecure --host=cockroachdb-public

Waiting for pod default/cockroachdb to be running, status is Pending, pod ready: false

Hit enter for command prompt

root@cockroachdb-public:26257> CREATE DATABASE bank;

CREATE DATABASE

root@cockroachdb-public:26257> CREATE TABLE bank.accounts (id INT PRIMARY KEY, balance DECIMAL);

CREATE TABLE

root@cockroachdb-public:26257> INSERT INTO bank.accounts VALUES (1234, 10000.50);

INSERT 1

root@cockroachdb-public:26257> SELECT * FROM bank.accounts;

+------+---------+
|id  |balance |
+------+---------+
|1234 |10000.5 |
+------+---------+

(1 row)

```

### Step Four: Accessing the CockroachDB Console

### 第四步：访问 CockroachDB 控制台

To get more information into cluster behavior and health, you can pull up the CockroachDB Console by port-forwarding from your local machine to one of the pods as shown below:

要获取有关集群行为和健康状况的更多信息，您可以通过从本地机器到 pod 之一的端口转发来拉出 CockroachDB 控制台，如下所示：

If you want to see information about how the cluster is doing, you can try pulling up the CockroachDB admin UI by port-forwarding from your local machine to one of the pods:

如果您想查看有关集群运行情况的信息，您可以尝试通过从本地机器到其中一个 pod 的端口转发来拉出 CockroachDB 管理 UI：

`kubectl port-forward cockroachdb-0 8080`

`kubectl port-forward cockroachdb-0 8080`

You should now be able to access the admin UI by visiting [http://localhost:8080/](http://localhost:8080/) in your web browser:

您现在应该可以通过在 Web 浏览器中访问 [http://localhost:8080/](http://localhost:8080/) 来访问管理 UI：

![CockroachDB on Kubernetes - DB Console Screen](https://crl2020.imgix.net/img/cockroachdb-console-_k8s.png?auto=format,compress&max-w=700)

### Step Five: Simulating node failure

### 第五步：模拟节点故障

We talked about DB survivability earlier. Now you can test it for yourself. What happens when a pod goes bad or gets deleted?

我们之前讨论过数据库的生存性。现在您可以自己测试一下。当 Pod 变坏或被删除时会发生什么？

- To test the resiliency of the cluster, try killing some of the containers by running a command like`kubectl delete pod cockroachdb-3`. This must be done from a different terminal while you’re still accessing the cluster from your SQL shell.
- If you get a “bad connection” error from deleting the same instance your shell was communicating with, simply retry the query.

- 要测试集群的弹性，请尝试通过运行类似 `kubectl delete pod cockroachdb-3` 的命令杀死一些容器。当您仍然从 SQL shell 访问集群时，这必须从不同的终端完成。
- 如果您在删除与您的 shell 通信的同一个实例时遇到“错误连接”错误，只需重试查询即可。

The container will now be recreated for you by the StatefulSet controller, just as it would happen in the event of a real production failure.

现在，StatefulSet 控制器将为您重新创建容器，就像发生真正的生产故障时一样。

If you’re up for testing the durability of the cluster data, you can try deleting all the pods at once and ensuring they start up properly again from their persistent volumes. To do this, you can run `kubectl delete pod –selector app=cockroachdb`, which deletes all pods that have the [label](http://kubernetes.io/docs/stable/user-guide/labels/) `app =cockroachdb.`  This includes the pods from our StatefulSet.

如果您准备测试集群数据的持久性，您可以尝试一次删除所有 pod，并确保它们从持久卷中再次正确启动。为此，您可以运行 `kubectl delete pod –selector app=cockroachdb`，它会删除所有具有 [label](http://kubernetes.io/docs/stable/user-guide/labels/) `app =cockroachdb.` 这包括来自我们 StatefulSet 的 pod。

Just like during setup, it might take some time for them all to come back up again. But once they are up and running again, you’ll be able to get the same data back from the SQL queries you’re making in the shell.

就像在设置过程中一样，它们可能需要一些时间才能再次恢复。但是一旦它们再次启动并运行，您将能够从您在 shell 中进行的 SQL 查询中获取相同的数据。

### Step Six: Scaling the CockroachDB cluster

### 第六步：扩展 CockroachDB 集群

[Before removing nodes from your cluster, you must first tell CockroachDB to decommission them](https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html#remove-nodes). (This lets nodes finish in-flight requests, rejects any new requests, and transfer all range replicas and range leases off the nodes.

[在从集群中删除节点之前，您必须先告诉 CockroachDB 将它们退役](https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html#remove-nodes)。 （这允许节点完成正在进行的请求，拒绝任何新请求，并将所有范围副本和范围租约转移到节点上。

Now that the nodes are decommissioned you can scale your Kubernetes cluster by simply adding or subtracting replicas by resizing the StatefulSet as shown below:

现在节点已退役，您可以通过调整 StatefulSet 的大小来简单地添加或减少副本来扩展您的 Kubernetes 集群，如下所示：

`kubectl scale statefulset cockroachdb --replicas=4`

`kubectl scale statefulset cockroachdb --replicas=4`

Step Seven: Shutting the CockroachDB cluster down

第七步：关闭 CockroachDB 集群

Once you’re done, a single command will clean up all the resources we’ve created during our oh-so-brief Kubernetes tutorial. The [labels](http://kubernetes.io/docs/stable/user-guide/labels/) we added to the resources do all the work.

完成后，一个命令将清理我们在非常简短的 Kubernetes 教程中创建的所有资源。我们添加到资源中的 [标签](http://kubernetes.io/docs/stable/user-guide/labels/) 完成了所有工作。

`kubectl delete statefulsets,pods,persistentvolumes,persistentvolumeclaims,services,poddisruptionbudget -l app=cockroachdb`

`kubectl delete statefulsets,pods,persistentvolumes,persistentvolumeclaims,services,poddisruptionbudget -l app=cockroachdb`

You can also shut down your entire Kubernetes cluster by running:

您还可以通过运行以下命令关闭整个 Kubernetes 集群：

`gcloud container clusters delete cockroachdb-cluster`

`gcloud 容器集群删除 cockroachdb-cluster`

### Where to go from here

###  从这往哪儿走

### Now that you’ve mastered the basics, what next?

### 现在您已经掌握了基础知识，接下来要做什么？

- Writing applications that use the CockroachDB cluster via one of the [many supported client libraries](https://www.cockroachlabs.com/docs/stable/install-client-drivers.html).
- Modifying how the cluster is initialized [to use certificates for encryption between nodes](https://www.cockroachlabs.com/docs/stable/secure-a-cluster.html).
- Setting up a cluster in a cloud or bare metal environment using a different kind of PersistentVolume, rather than on Container Engine.
- [Setting up Prometheus](https://coreos.com/blog/prometheus-and-kubernetes-up-and-running.html) to monitor CockroachDB within the cluster, taking advantage of the annotations we put on the CockroachDB StatefulSet.
- Contributing feature requests, issues, or improvements to [CockroachDB](https://github.com/cockroachdb/cockroach), either for the Kubernetes documentation or for the core database itself!
- Check out more[tutorials and tech talks about Kubernetes](https://www.cockroachlabs.com/kubernetes-bootcamp/)

- 通过 [许多受支持的客户端库] (https://www.cockroachlabs.com/docs/stable/install-client-drivers.html) 之一编写使用 CockroachDB 集群的应用程序。
- 修改集群的初始化方式[使用证书在节点之间进行加密](https://www.cockroachlabs.com/docs/stable/secure-a-cluster.html)。
- 使用不同类型的 PersistentVolume 在云或裸机环境中设置集群，而不是在容器引擎上。
- [设置 Prometheus](https://coreos.com/blog/prometheus-and-kubernetes-up-and-running.html) 监控集群内的 CockroachDB，利用我们在 CockroachDB StatefulSet 上的注释。
- 为 [CockroachDB](https://github.com/cockroachdb/cockroach) 贡献功能请求、问题或改进，无论是针对 Kubernetes 文档还是针对核心数据库本身！
- 查看更多[关于 Kubernetes 的教程和技术讨论](https://www.cockroachlabs.com/kubernetes-bootcamp/)

### References 

###  参考

More information and up-to-date configuration files for running CockroachDB on Kubernetes can be found [in our documentation](https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html).

有关在 Kubernetes 上运行 CockroachDB 的更多信息和最新配置文件，可以在 [我们的文档](https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html) 中找到。

- [deployment](https://www.cockroachlabs.com/tags/deployment/)
- [Kubernetes](https://www.cockroachlabs.com/tags/kubernetes/)
- [containers](https://www.cockroachlabs.com/tags/containers/)
- [cloud deployment](https://www.cockroachlabs.com/tags/cloud-deployment/)
- [engineering](https://www.cockroachlabs.com/tags/engineering/) 

- [部署](https://www.cockroachlabs.com/tags/deployment/)
- [Kubernetes](https://www.cockroachlabs.com/tags/kubernetes/)
- [容器](https://www.cockroachlabs.com/tags/containers/)
- [云部署](https://www.cockroachlabs.com/tags/cloud-deployment/)
- [工程](https://www.cockroachlabs.com/tags/engineering/)

