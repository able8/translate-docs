# Using GitOps to Self-Manage Postgres in Kubernetes

# 使用 GitOps 在 Kubernetes 中自我管理 Postgres

February 01, 2021 [Jonathan S. Katz](https://blog.crunchydata.com/blog/author/jonathan-s-katz)

2021 年 2 月 1 日 [Jonathan S. Katz](https://blog.crunchydata.com/blog/author/jonathan-s-katz)

[PostgreSQL](https://blog.crunchydata.com/blog/topic/postgresql)[Kubernetes](https://blog.crunchydata.com/blog/topic/kubernetes) [PostgreSQL Operator](https://blog.crunchydata.com/blog/topic/postgresql-operator) [GitOps](https://blog.crunchydata.com/blog/topic/gitops)

[PostgreSQL](https://blog.crunchydata.com/blog/topic/postgresql)[Kubernetes](https://blog.crunchydata.com/blog/topic/kubernetes) [PostgreSQL Operator](https://blog.crunchydata.com/blog/topic/postgresql-operator) [GitOps](https://blog.crunchydata.com/blog/topic/gitops)

" [GitOps](https://www.gitops.tech/)" is a term that I've been seeing come up more and more. The concept was first put forward by the team at [Weaveworks](https://www.weave.works/technologies/gitops/) as a way to consolidate thought around deploying applications. In essence: your deployment topology lives in your git repository. You can update your deployment information by adding a new commit. Likewise, if you need to revert your system's state, you can rollback to the commit that you want to represent your production environment. Any changes to your deployment topology should be reconciled in your production environment.

“ [GitOps](https://www.gitops.tech/)”是我看到越来越多的术语。这个概念最初是由[Weaveworks](https://www.weave.works/technologies/gitops/) 的团队提出的，作为一种整合部署应用程序思想的方式。本质上：您的部署拓扑存在于您的 git 存储库中。您可以通过添加新提交来更新部署信息。同样，如果您需要恢复系统状态，您可以回滚到您想要代表您的生产环境的提交。对部署拓扑的任何更改都应在您的生产环境中进行协调。

A lot of the conversations around GitOps came around the [Postgres Operator](https://github.com/CrunchyData/postgres-operator) for [Kubernetes](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes/) and how to apply these principles. Platforms like Kubernetes make it relatively seamless to apply the ideas of GitOps to the stateless pieces of applications (e.g., your web application). More work needs to be done with a stateful service such as [PostgreSQL](https://www.postgresql.org/), as each stateful application can have unique requirements. Let's take something like replication: [PostgreSQL's replication system](https://www.postgresql.org/docs/current/high-availability.html) is both configured and managed differently than in other database systems, and any GitOps management tool, such as [Helm](https://helm.sh/) or [Kustomize](https://kustomize.io/), would have to account for this.

很多关于 GitOps 的讨论都围绕着 [Kubernetes](https://www.crunchydata.com/products/crunchy-postgresql-) 的 [Postgres Operator](https://github.com/CrunchyData/postgres-operator)展开for-kubernetes/) 以及如何应用这些原则。像 Kubernetes 这样的平台可以相对无缝地将 GitOps 的思想应用于无状态的应用程序（例如，您的 Web 应用程序）。需要使用有状态服务（例如 [PostgreSQL](https://www.postgresql.org/))来完成更多工作，因为每个有状态应用程序都可能有独特的要求。让我们以复制为例：[PostgreSQL的复制系统](https://www.postgresql.org/docs/current/high-availability.html) 的配置和管理都与其他数据库系统以及任何 GitOps 管理工具不同，例如 [Helm](https://helm.sh/) 或 [Kustomize](https://kustomize.io/)，就必须考虑到这一点。

This is the beauty of the [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/). It provides a generic framework for developers to allow for GitOps style management of a stateful application in a way that can be "simply" configured. I say "simply" because there is still a lot to consider when running a complex stateful application like a database in a production environment, but an Operator can make the overall management and application of changes easier.

这就是 [Operator 模式](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) 的美妙之处。它为开发人员提供了一个通用框架，允许以一种可以“简单”配置的方式对有状态应用程序进行 GitOps 风格的管理。我说“简单”是因为在生产环境中运行像数据库这样复杂的有状态应用程序时仍然需要考虑很多，但是 Operator 可以使更改的整体管理和应用程序变得更容易。

There are many ways to support GitOps workflows in current versions of the PostgreSQL Operator, from Kubernetes [YAML files](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/#create-a-postgresql-cluster) to [Helm charts](https://github.com/CrunchyData/postgres-operator/tree/master/examples/helm) to [Kustomize manifests](https://github.com/CrunchyData/postgres-operator/tree/master/examples/kustomize/createcluster).

在当前版本的 PostgreSQL Operator 中支持 GitOps 工作流的方法有很多，来自 Kubernetes [YAML 文件](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/#create-a-postgresql-cluster) 到 [Helm 图表](https://github.com/CrunchyData/postgres-operator/tree/master/examples/helm) 到 [Kustomize manifests](https://github.com/CrunchyData/postgres-operator/tree/master/examples/kustomize/createcluster)。

Let's work through a GitOps style workflow. The examples below will assume that you have [installed the Postgres Operator](https://access.crunchydata.com/documentation/postgres-operator/latest/quickstart/).

让我们完成一个 GitOps 风格的工作流程。下面的示例假设您已经[安装了 Postgres Operator](https://access.crunchydata.com/documentation/postgres-operator/latest/quickstart/)。

## Deploying a HA PostgreSQL Cluster

## 部署一个 HA PostgreSQL 集群

For the first example, let's create a [HA PostgreSQL](https://access.crunchydata.com/documentation/postgres-operator/latest/architecture/high-availability/) cluster called hippo using a Kubernetes YAML file. In your command-line environment, set the following environmental variables to where you want your PostgreSQL cluster deployed. 

对于第一个示例，让我们使用 Kubernetes YAML 文件创建一个名为 hippo 的 [HA PostgreSQL](https://access.crunchydata.com/documentation/postgres-operator/latest/architecture/high-availability/) 集群。在您的命令行环境中，将以下环境变量设置为您希望部署 PostgreSQL 集群的位置。

For this example, the namespace uses the same one created as part of the [quickstart](https://access.crunchydata.com/documentation/postgres-operator/latest/quickstart/). You can copy and paste the below into your environment to try the example out. You can also modify the [custom resource](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/) manifest to match your specific environment. Note that if your environment does not support environmental variables, you can find/replace the below values in your manifest file.

对于此示例，命名空间使用作为 [quickstart](https://access.crunchydata.com/documentation/postgres-operator/latest/quickstart/) 的一部分创建的名称空间。您可以将以下内容复制并粘贴到您的环境中以试用该示例。您还可以修改 [自定义资源](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/) 清单以匹配您的特定环境。请注意，如果您的环境不支持环境变量，您可以在清单文件中找到/替换以下值。

From your command line, execute the following:

从您的命令行，执行以下命令：

```
# this variable is the namespace the cluster is being deployed into
export cluster_namespace=pgo
# this variable is the name of the cluster being created
export pgo_cluster_name=hippo
# this variable sets the default disk size
export cluster_disk_size=5Gi

cat <<-EOF > "${pgo_cluster_name}-pgcluster.yaml"
apiVersion: crunchydata.com/v1
kind: Pgcluster
metadata:
annotations:
    current-primary: ${pgo_cluster_name}
labels:
    crunchy-pgha-scope: ${pgo_cluster_name}
    deployment-name: ${pgo_cluster_name}
    name: ${pgo_cluster_name}
    pg-cluster: ${pgo_cluster_name}
    pgo-version: 4.6.0
name: ${pgo_cluster_name}
namespace: ${cluster_namespace}
spec:
BackrestStorage:
    accessmode: ReadWriteOnce
    size: ${cluster_disk_size}
    storagetype: dynamic
PrimaryStorage:
    accessmode: ReadWriteOnce
    name: ${pgo_cluster_name}
    size: ${cluster_disk_size}
    storagetype: dynamic
ReplicaStorage:
    accessmode: ReadWriteOnce
    size: ${cluster_disk_size}
    storagetype: dynamic
ccpimage: crunchy-postgres-ha
ccpimageprefix: registry.developers.crunchydata.com/crunchydata
ccpimagetag: centos8-13.1-4.6.0
clustername: ${pgo_cluster_name}
database: ${pgo_cluster_name}
exporterport: "9187"
limits: {}
name: ${pgo_cluster_name}
namespace: ${cluster_namespace}
pgDataSource: {}
pgbadgerport: "10000"
pgoimageprefix: registry.developers.crunchydata.com/crunchydata
podAntiAffinity:
    default: preferred
    pgBackRest: preferred
    pgBouncer: preferred
port: "5432"
replicas: "1"
user: ${pgo_cluster_name}
userlabels:
    pgo-version: 4.6.0
EOF

kubectl apply -f "${pgo_cluster_name}-pgcluster.yaml"
```

This manifest tells the PostgreSQL Operator to create a high availability PostgreSQL cluster. It will take a few moments to get everything provisioned. You can check on the status using kubectl get pods or using the [pgo test](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorial/create-cluster/) command if you have installed the PostgreSQL Operator client. For example, the below demonstrates that there are two PostgreSQL instances in the hippo cluster available:

此清单告诉 PostgreSQL Operator 创建高可用性 PostgreSQL 集群。需要一些时间来配置所有内容。如果您已经安装了 PostgreSQL，您可以使用 kubectl get pods 或使用 [pgo test](https://access.crunchydata.com/documentation/postgres-operator/latest/tutorial/create-cluster/) 命令检查状态运营商客户端。例如，下图演示了 hippo 集群中有两个 PostgreSQL 实例可用：

```
kubectl -n "${cluster_namespace}" get pods  --selector="pg-cluster=${pgo_cluster_name},pgo-pg-database"
NAME                                          READY   STATUS      RESTARTS   AGE
hippo-585bb4f797-f76bl                        1/1     Running     0          3m32s
hippo-gftf-7f55674d78-gsjx7                   1/1     Running     0          2m53s
```

Let's log into our newly provisioned database cluster. In our example above, we created a database named hippo along with a user called hippo. The Postgres Operator creates user credentials for several defaults users, including any users we specify. To get the credentials, assuming you still have the environmental variables set from the earlier step, you can use the following command:

让我们登录到我们新配置的数据库集群。在上面的示例中，我们创建了一个名为 hippo 的数据库以及一个名为 hippo 的用户。 Postgres Operator 为多个默认用户创建用户凭据，包括我们指定的任何用户。要获取凭据，假设您仍然具有前面步骤中设置的环境变量，则可以使用以下命令：

```
kubectl -n "${cluster_namespace}" get secrets "${pgo_cluster_name}-${pgo_cluster_name}-secret" -o jsonpath="{.data.password}" |base64 -d
```

For convenience for connecting to the cluster, we can store the user's password directly in an environmental variable that psql recognizes:

为了方便连接集群，我们可以直接将用户的密码保存在一个psql可以识别的环境变量中：

```
export PGPASSWORD=$(kubectl -n jkatz get secrets "${pgo_cluster_name}-${pgo_cluster_name}-secret" -o jsonpath="{.data.password}" | base64 -d)
```

In a separate terminal window, open up a port-forward:

在一个单独的终端窗口中，打开一个端口转发：

```
export cluster_namespace=pgo
export pgo_cluster_name=hippo
kubectl port-forward -n "${cluster_namespace}" "svc/${pgo_cluster_name}" 5432:5432
```

Now, back in the original window, you can now connect to the cluster:

现在，回到原始窗口，您现在可以连接到集群：

```
psql -h localhost -U "${pgo_cluster_name}" "${pgo_cluster_name}"
psql (13.1)
Type "help" for help.

hippo=>
```

Success!

成功！

## Adding More Resources to a Deployed Cluster 

## 向已部署的集群添加更多资源

Part of the GitOps principle is the ability to modify (or version) a configuration file and have the changes reflected in your environment. This is typical of "Day 2" type of operations, such as requiring more memory / CPU resources as the workload on a database increases.

GitOps 原则的一部分是能够修改（或版本）配置文件并将更改反映在您的环境中。这是典型的“第 2 天”类型的操作，例如随着数据库工作负载的增加需要更多的内存/CPU 资源。

Let's say we want to raise our memory limits to 2Gi and our CPU limit to 2.0 cores. Open up the file you created in the previous step, and add the following block to the spec:

假设我们想将内存限制提高到 2Gi，将 CPU 限制提高到 2.0 核。打开您在上一步中创建的文件，并将以下块添加到规范中：

```
limits:
memory: 2Gi
cpu: 2.0
```

Your file should look similar to this:

您的文件应类似于以下内容：

```
apiVersion: crunchydata.com/v1
kind: Pgcluster
metadata:
annotations:
    current-primary: hippo
labels:
    crunchy-pgha-scope: hippo
    deployment-name: hippo
    name: hippo
    pg-cluster: hippo
    pgo-version: 4.6.0
name: hippo
namespace: pgo
spec:
BackrestStorage:
    accessmode: ReadWriteOnce
    size: 5Gi
    storagetype: dynamic
PrimaryStorage:
    accessmode: ReadWriteOnce
    name: hippo
    size: 5Gi
    storagetype: dynamic
ReplicaStorage:
    accessmode: ReadWriteOnce
    size: 5Gi
    storagetype: dynamic
ccpimage: crunchy-postgres-ha
ccpimageprefix: registry.developers.crunchydata.com/crunchydata
ccpimagetag: centos8-13.1-4.6.0
clustername: hippo
database: hippo
exporterport: "9187"
limits:
    cpu: 2.0
    memory: 2Gi
name: hippo
namespace: pgo
pgDataSource: {}
pgbadgerport: "10000"
pgoimageprefix: registry.developers.crunchydata.com/crunchydata
podAntiAffinity:
    default: preferred
    pgBackRest: preferred
    pgBouncer: preferred
port: "5432"
replicas: "1"
user: hippo
userlabels:
    pgo-version: 4.6.0
```

After saving your changes, apply the updates to your Kubernetes environment:

保存更改后，将更新应用到 Kubernetes 环境：

```
kubectl apply -f "${pgo_cluster_name}-pgcluster.yaml"
```

Upon detecting the change, the Postgres Operator uses a [rolling update strategy](https://access.crunchydata.com/documentation/postgres-operator/latest/architecture/high-availability/#rolling-updates) to apply the resource changes to each PostgreSQL instance in a way that minimizes downtime. Wait a few moments for the changes to be apple, and then run a describe on the Pods to see the changes (output truncated):

检测到更改后，Postgres Operator 使用 [滚动更新策略](https://access.crunchydata.com/documentation/postgres-operator/latest/architecture/high-availability/#rolling-updates) 来应用资源更改以最小化停机时间的方式发送到每个 PostgreSQL 实例。稍等片刻，更改为苹果，然后在 Pod 上运行描述以查看更改（输出被截断)：

```
kubectl -n jkatz describe pods --selector="pg-cluster=${pgo_cluster_name},pgo-pg-database"
Name:         hippo-585bb4f797-f76bl
Namespace:    pgo
Containers:
database:
    Limits:
      cpu:     1000m
      memory:  2Gi
    Requests:
      cpu:      1000m
      memory:   2Gi

Name:         hippo-gftf-7f55674d78-gsjx7
Namespace:    pgo
Containers:
database:
    Limits:
      cpu:     1000m
      memory:  2Gi
    Requests:
      cpu:      1000m
      memory:   2Gi
```

Excellent! And using GitOps principles, if I wanted to revert these changes, I could do so by removing the "limits" clause and the Postgres Operator would update the cluster as such.

优秀！并且使用 GitOps 原则，如果我想恢复这些更改，我可以通过删除“限制”子句来实现，Postgres Operator 将更新集群。

For a full list of attributes that can be updated, please see the [custom resources](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/) section of the [Postgres Operator documentation] (https://access.crunchydata.com/documentation/postgres-operator/).

有关可以更新的完整属性列表，请参阅 [Postgres Operator 文档] 的 [自定义资源](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/)部分（https://access.crunchydata.com/documentation/postgres-operator/)。

## Evolution of GitOps and Stateful Services

## GitOps 和有状态服务的演变

When coupled with a tool like the [PostgreSQL Operator](https://github.com/CrunchyData/postgres-operator), GitOps principles can be extended to work for stateful services. Applying a GitOps mindset to managing PostgreSQL workloads can help make it easier to manage production-grade Postgres instances on Kubernetes. GitOps principles can certainly make it easier for deploying a range of PostgreSQL topologies, from single instances to multi-zone, fault tolerant clusters.

当与 [PostgreSQL Operator](https://github.com/CrunchyData/postgres-operator) 等工具结合使用时，GitOps 原则可以扩展为适用于有状态服务。将 GitOps 思维方式应用于管理 PostgreSQL 工作负载有助于更轻松地管理 Kubernetes 上的生产级 Postgres 实例。 GitOps 原则当然可以更轻松地部署一系列 PostgreSQL 拓扑，从单实例到多区域、容错集群。

Upcoming posts will look at how some of the other Kubernetes toolsets can make it even easier to work with the PostgreSQL Operator in a GitOps manner.

即将发布的文章将着眼于其他一些 Kubernetes 工具集如何以 GitOps 方式更轻松地使用 PostgreSQL Operator。

Like what you're reading? Stay informed by subscribing for our newsletter!

比如你在读什么？订阅我们的时事通讯，随时了解最新信息！

## Newsletter

## 时事通讯

Like what you're reading? Stay informed by subscribing for our newsletter! 

比如你在读什么？订阅我们的时事通讯，随时了解最新信息！

