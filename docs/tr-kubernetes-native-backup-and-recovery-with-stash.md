# Tutorial: Kubernetes-Native Backup and Recovery With Stash

# 教程：使用 Stash 进行 Kubernetes 原生备份和恢复

![Tutorial: Kubernetes-Native Backup and Recovery With Stash](https://appfleet-com.cdn.ampproject.org/i/s/appfleet.com/blog/content/images/size/w1200/2020/07/91-Tutorial-Kubernetes-Native-Backup-and-Recovery-With-Stash.png)

91-Tutorial-Kubernetes-Native-Backup-and-Recovery-With-Stash.png)

* * *

* * *

## Intro

##  介绍

Having a proper backup recovery plan is vital to any organization's IT operation. However, when you begin to distribute workloads across data centers and regions, that process begins to become more and more complex. Container orchestration platforms such as Kubernetes have begun to ease this burden and enabled the management of distributed workloads in areas that were previously very challenging.

拥有适当的备份恢复计划对于任何组织的 IT 运营都至关重要。但是，当您开始跨数据中心和区域分配工作负载时，该过程开始变得越来越复杂。 Kubernetes 等容器编排平台已经开始减轻这种负担，并在以前非常具有挑战性的领域实现分布式工作负载的管理。

In this post, we are going to introduce you to a Kubernetes-native tool for taking backups of your disks, helping with the crucial recovery plan. **Stash is a [Restic](https://restic.net/) Operator that accelerates the task of backing up and recovering your Kubernetes infrastructure**. You can read more about the Operator Framework via [this blog post](https://appfleet.com/blog/first-steps-with-the-kubernetes-operator/).

在这篇文章中，我们将向您介绍一个 Kubernetes 原生工具，用于备份您的磁盘，帮助您制定关键的恢复计划。 **Stash 是一个 [Restic](https://restic.net/) Operator，可加速备份和恢复 Kubernetes 基础设施的任务**。您可以通过 [这篇博文](https://appfleet.com/blog/first-steps-with-the-kubernetes-operator/) 阅读有关 Operator Framework 的更多信息。

## How does Stash work?

## Stash 是如何工作的？

Using Stash, you can backup Kubernetes volumes mounted in following types of workloads:

使用 Stash，您可以备份安装在以下类型工作负载中的 Kubernetes 卷：

- Deployment
- DaemonSet
- ReplicaSet
- ReplicationController
- StatefulSet

- 部署
- 守护进程集
- 副本集
- 复制控制器
- 状态集

At the heart of Stash is a Kubernetes [controller](https://book.kubebuilder.io/basics/what_is_a_controller.html) which uses [Custom Resource Definition (CRD)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) to specify targets and behaviors of the backup and restore process in a Kubernetes native way. A simplified architecture of Stash is shown below:

Stash 的核心是一个 Kubernetes [控制器](https://book.kubebuilder.io/basics/what_is_a_controller.html)，它使用 [自定义资源定义 (CRD)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) 以 Kubernetes 原生方式指定备份和恢复过程的目标和行为。 Stash 的简化架构如下所示：

![](https://appfleet-com.cdn.ampproject.org/i/s/appfleet.com/blog/content/images/2020/03/stash_architecture.svg)

## Installing Stash

## 安装 Stash

### Using Helm 3

### 使用 Helm 3

Stash can be installed via [Helm](https://helm.sh/) using the [chart](https://github.com/stashed/installer/tree/v0.9.0-rc.6/charts/stash) from [AppsCode Charts Repository](https://github.com/appscode/charts). To install the chart with the release name `stash-operator`:

可以使用 [chart](https://github.com/stashed/installer/tree/v0.9.0-rc.6/charts/stash) 通过 [Helm](https://helm.sh/) 安装 Stash来自 [AppsCode Charts Repository](https://github.com/appscode/charts)。要安装发行版名称为 `stash-operator` 的图表：

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/stash --version v0.9.0-rc.6
NAME            CHART          VERSION      APP VERSION DESCRIPTION
appscode/stash  v0.9.0-rc.6    v0.9.0-rc.6  Stash by AppsCode - Backup your Kubernetes Volumes

$ helm install stash-operator appscode/stash \
  --version v0.9.0-rc.6 \
  --namespace kube-system
```

### Using YAML

### 使用 YAML

If you prefer to not use Helm, you can generate YAMLs from Stash chart and deploy using `kubectl`:

如果你不想使用 Helm，你可以从 Stash chart 生成 YAMLs 并使用 `kubectl` 进行部署：

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/stash --version v0.9.0-rc.6
NAME            CHART VERSION APP VERSION DESCRIPTION
appscode/stash  v0.9.0-rc.6    v0.9.0-rc.6  Stash by AppsCode - Backup your Kubernetes Volumes

$ helm template stash-operator appscode/stash \
  --version v0.9.0-rc.6 \
  --namespace kube-system \
  --no-hooks |kubectl apply -f -
```

### **Installing on GKE Cluster**

### **在 GKE 集群上安装**

If you are installing Stash on a GKE cluster, you will need cluster admin permissions to install Stash operator. Run the following command to grant admin permission to the cluster.

如果您要在 GKE 集群上安装 Stash，则需要集群管理员权限才能安装 Stash 操作员。运行以下命令向集群授予管理员权限。

```console
$ kubectl create clusterrolebinding "cluster-admin-$(whoami)" \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"

```

In addition, if your GKE cluster is a [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters), you will need to either add an additional firewall rule that allows master nodes access port `8443/tcp` on worker nodes, or change the existing rule that allows access to ports `443/tcp` and `10250/tcp` to also allow access to port `8443/tcp`. The procedure to add or modify firewall rules is described in the official GKE documentation for private clusters mentioned above.

此外，如果您的 GKE 集群是 [私有集群](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters)，则您需要添加额外的防火墙规则允许主节点访问工作节点上的端口“8443/tcp”，或更改允许访问端口“443/tcp”和“10250/tcp”的现有规则也允许访问端口“8443/tcp”。上面提到的私有集群的官方 GKE 文档中描述了添加或修改防火墙规则的过程。

### Verify installation

### 验证安装

To check if Stash operator pods have started, run the following command:

要检查 Stash 操作员 Pod 是否已启动，请运行以下命令：

```console
$ kubectl get pods --all-namespaces -l app=stash --watch

NAMESPACE     NAME                              READY     STATUS    RESTARTS   AGE
kube-system   stash-operator-859d6bdb56-m9br5   2/2       Running   2          5s

```

Once the operator pods are running, you can cancel the above command by typing `Ctrl+C`. 

一旦操作员 Pod 运行，您可以通过键入“Ctrl+C”来取消上述命令。

Now, to confirm CRD groups have been registered by the operator, run the following command:

现在，要确认操作员已注册 CRD 组，请运行以下命令：

```console
$ kubectl get crd -l app=stash

NAME AGE
recoveries.stash.appscode.com        5s
repositories.stash.appscode.com      5s
restics.stash.appscode.com           5s

```

With this, you are ready to take your first backup using Stash.

有了这个，您就可以使用 Stash 进行第一次备份了。

## Configuring Auto Backup for Database

## 为数据库配置自动备份

To keep everything isolated, we are going to use a separate namespace called `demo` throughout this tutorial.

为了保持一切隔离，我们将在本教程中使用一个名为“demo”的单独命名空间。

```console
$ kubectl create ns demo
namespace/demo created
```

### Prepare Backup Blueprint

### 准备备份蓝图

We are going to use [GCS Backend](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/gcs) to store the backed up data. You can use any supported backend you prefer. You just have to configure Storage Secret and `spec.backend` section of `BackupBlueprint` to match with your backend. Visit [here](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/overview) to learn which backends are supported by Stash and how to configure them.

我们将使用 [GCS Backend](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/gcs) 来存储备份的数据。您可以使用您喜欢的任何受支持的后端。你只需要配置 Storage Secret 和 `BackupBlueprint` 的 `spec.backend` 部分来匹配你的后端。访问 [此处](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/overview) 了解 Stash 支持哪些后端以及如何配置它们。

For GCS backend, if the bucket does not exist, Stash needs `Storage Object Admin` role permissions to create the bucket. For more details, please check the following [guide](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/gcs).

对于GCS后端，如果bucket不存在，Stash需要`Storage Object Admin`角色权限来创建bucket。更多详情，请查看以下[指南](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/gcs)。

**Create Storage Secret:**

**创建存储密钥：**

At first, let’s create a Storage Secret for the GCS backend,

首先，让我们为 GCS 后端创建一个 Storage Secret，

```console
$ echo -n 'changeit' > RESTIC_PASSWORD
$ echo -n '<your-project-id>' > GOOGLE_PROJECT_ID
$ mv downloaded-sa-json.key > GOOGLE_SERVICE_ACCOUNT_JSON_KEY
$ kubectl create secret generic -n demo gcs-secret \
    --from-file=./RESTIC_PASSWORD \
    --from-file=./GOOGLE_PROJECT_ID \
    --from-file=./GOOGLE_SERVICE_ACCOUNT_JSON_KEY
secret/gcs-secret created

```

**Create BackupBlueprint:**

**创建备份蓝图：**

Next, we have to create a `BackupBlueprint` CRD with a blueprint for `Repository` and `BackupConfiguration` object.

接下来，我们必须创建一个带有“Repository”和“BackupConfiguration”对象蓝图的“BackupBlueprint”CRD。

Below is the YAML of the `BackupBlueprint` object that we are going to create:

下面是我们将要创建的 `BackupBlueprint` 对象的 YAML：

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: BackupBlueprint
metadata:
name: postgres-backup-blueprint
spec:
# ============== Blueprint for Repository ==========================
backend:
    gcs:
      bucket: appscode-qa
      prefix: stash-backup/${TARGET_NAMESPACE}/${TARGET_APP_RESOURCE}/${TARGET_NAME}
    storageSecretName: gcs-secret
# ============== Blueprint for BackupConfiguration =================
task:
    name: postgres-backup-${TARGET_APP_VERSION}
schedule: "*/5 * * * *"
retentionPolicy:
    name: 'keep-last-5'
    keepLast: 5
    prune: true
```

Note that we have used few variables (format: `${<variable name>}`) in the `spec.backend.gcs.prefix` field. Stash will substitute these variables with values from the respective target. To learn which variables you can use in the `prefix` field, please visit [here](https://stash.run/docs/v0.9.0-rc.6/concepts/crds/backupblueprint#repository-blueprint).

请注意，我们在`spec.backend.gcs.prefix` 字段中使用了很少的变量（格式：`${<variable name>}`）。 Stash 将用来自相应目标的值替换这些变量。要了解您可以在 `prefix` 字段中使用哪些变量，请访问 [此处](https://stash.run/docs/v0.9.0-rc.6/concepts/crds/backupblueprint#repository-blueprint)。

Let’s create the `BackupBlueprint` that we have shown above.

让我们创建上面显示的“BackupBlueprint”。

```console
$ kubectl apply -f https://github.com/stashed/docs/raw/v0.9.0-rc.6/docs/examples/guides/latest/auto-backup/database/backupblueprint.yaml
backupblueprint.stash.appscode.com/postgres-backup-blueprint created

```

With this, automatic backup is configured for PostgreSQL database. We just have to add an annotation to the `AppBinding` of the targeted database.

这样，PostgreSQL 数据库就配置了自动备份。我们只需要在目标数据库的“AppBinding”中添加一个注解。

**Required Annotation for Auto-Backup Database:**

**自动备份数据库所需的注释：**

You have to add the following annotation to the `AppBinding` CRD of the targeted database to enable backup for it:

您必须将以下注释添加到目标数据库的“AppBinding”CRD 以为其启用备份：

```yaml
stash.appscode.com/backup-blueprint: <BackupBlueprint name>

```

This annotation specifies the name of the `BackupBlueprint` object where a blueprint for `Repository` and `BackupConfiguration` has been defined.

此注释指定了“BackupBlueprint”对象的名称，其中定义了“Repository”和“BackupConfiguration”的蓝图。

### Prepare Databases

### 准备数据库

Next, we are going to deploy two sample PostgreSQL databases of two different versions using KubeDB. We are going to backup these two databases using auto-backup.

接下来，我们将使用 KubeDB 部署两个不同版本的两个示例 PostgreSQL 数据库。我们将使用自动备份来备份这两个数据库。

**Deploy First PostgreSQL Sample:**

**部署第一个 PostgreSQL 示例：**

Below is the YAML of the first `Postgres` CRD:

下面是第一个 `Postgres` CRD 的 YAML：

```yaml
apiVersion: kubedb.com/v1alpha1
kind: Postgres
metadata:
name: sample-postgres-1
namespace: demo
spec:
version: "11.2"
storageType: Durable
storage:
    storageClassName: "standard"
    accessModes:
    - ReadWriteOnce
    resources:
      requests:
        storage: 1Gi
terminationPolicy: Delete

```

Let’s create the `Postgres` we have shown above:

让我们创建上面显示的`Postgres`：

```console
$ kubectl apply -f https://github.com/stashed/docs/raw/v0.9.0-rc.6/docs/examples/guides/latest/auto-backup/database/sample-postgres-1.yaml
postgres.kubedb.com/sample-postgres-1 created

```

KubeDB will deploy a PostgreSQL database according to the above specification and it will create the necessary secrets and services to access the database. It will also create an `AppBinding` CRD that holds the necessary information to connect with the database.

KubeDB 将根据上述规范部署 PostgreSQL 数据库，并创建访问数据库所需的机密和服务。它还将创建一个“AppBinding”CRD，其中包含与数据库连接所需的信息。

Verify that an `AppBinding` has been created for this PostgreSQL sample:

验证是否已为此 PostgreSQL 示例创建了“AppBinding”：

```console
$ kubectl get appbinding -n demo
NAME AGE
sample-postgres-1   47s

```

If you view the YAML of this `AppBinding`, you will see it holds service and secret information. Stash uses this information to connect with the database.

如果您查看此 `AppBinding` 的 YAML，您将看到它包含服务和机密信息。 Stash 使用此信息连接数据库。

```console
$ kubectl get appbinding -n demo sample-postgres-1 -o yaml

```

```yaml
apiVersion: appcatalog.appscode.com/v1alpha1
kind: AppBinding
metadata:
name: sample-postgres-1
namespace: demo
...
spec:
clientConfig:
    service:
      name: sample-postgres-1
      path: /
      port: 5432
      query: sslmode=disable
      scheme: postgresql
secret:
    name: sample-postgres-1-auth
secretTransforms:
  - renameKey:
      from: POSTGRES_USER
      to: username
  - renameKey:
      from: POSTGRES_PASSWORD
      to: password
type: kubedb.com/postgres
version: "11.2"
```

**Deploy Second PostgreSQL Sample:**

**部署第二个 PostgreSQL 示例：**

Below is the YAML of the second `Postgres` object:

下面是第二个 `Postgres` 对象的 YAML：

```yaml
apiVersion: kubedb.com/v1alpha1
kind: Postgres
metadata:
name: sample-postgres-2
namespace: demo
spec:
version: "10.6-v2"
storageType: Durable
storage:
    storageClassName: "standard"
    accessModes:
    - ReadWriteOnce
    resources:
      requests:
        storage: 1Gi
terminationPolicy: Delete

```

Let’s create the `Postgres` we have shown above.

让我们创建上面显示的“Postgres”。

```console
$ kubectl apply -f https://github.com/stashed/docs/raw/v0.9.0-rc.6/docs/examples/guides/latest/auto-backup/database/sample-postgres-2.yaml
postgres.kubedb.com/sample-postgres-2 created

```

Verify that an `AppBinding` has been created for this PostgreSQL database:

验证是否已为此 PostgreSQL 数据库创建了“AppBinding”：

```console
$ kubectl get appbinding -n demo
NAME AGE
sample-postgres-1   2m49s
sample-postgres-2   10s

```

Here, we can see `AppBinding` `sample-postgres-2` has been created for our second PostgreSQL sample.

在这里，我们可以看到已经为我们的第二个 PostgreSQL 示例创建了 `AppBinding` `sample-postgres-2`。

## **Backup**

## **备份**

Next, we are going to add auto-backup specific annotation to the `AppBinding` of our desired database. Stash watches for `AppBinding` CRD. Once it finds an `AppBinding` with auto-backup annotation, it will create a `Repository` and a `BackupConfiguration` CRD according to respective `BackupBlueprint`. Then, rest of the backup process will proceed as normal database backup as described [here](https://stash.run/docs/v0.9.0-rc.6/guides/latest/addons/overview).

接下来，我们将向所需数据库的“AppBinding”添加特定于自动备份的注释。 Stash 监视`AppBinding` CRD。一旦找到带有自动备份注释的“AppBinding”，它就会根据各自的“BackupBlueprint”创建一个“Repository”和一个“BackupConfiguration”CRD。然后，其余的备份过程将按照 [此处](https://stash.run/docs/v0.9.0-rc.6/guides/latest/addons/overview) 所述的正常数据库备份进行。

### **Backup First PostgreSQL Sample**

### **备份第一个 PostgreSQL 示例**

Let’s backup our first PostgreSQL sample using auto-backup.

让我们使用自动备份来备份我们的第一个 PostgreSQL 样本。

**Add Annotations:**

**添加注释：**

At first, add the auto-backup specific annotation to the AppBinding `sample-postgres-1`:

首先，将特定于自动备份的注释添加到 AppBinding `sample-postgres-1` 中：

```console
$ kubectl annotate appbinding sample-postgres-1 -n demo --overwrite \
stash.appscode.com/backup-blueprint=postgres-backup-blueprint

```

Verify that the annotation has been added successfully:

验证注解是否添加成功：

```console
$ kubectl get appbinding -n demo sample-postgres-1 -o yaml

```

```yaml
apiVersion: appcatalog.appscode.com/v1alpha1
kind: AppBinding
metadata:
annotations:
    stash.appscode.com/backup-blueprint: postgres-backup-blueprint
name: sample-postgres-1
namespace: demo
...
spec:
clientConfig:
    service:
      name: sample-postgres-1
      path: /
      port: 5432
      query: sslmode=disable
      scheme: postgresql
secret:
    name: sample-postgres-1-auth
secretTransforms:
  - renameKey:
      from: POSTGRES_USER
      to: username
  - renameKey:
      from: POSTGRES_PASSWORD
      to: password
type: kubedb.com/postgres
version: "11.2"

```

Following this, Stash will create a `Repository` and a `BackupConfiguration` CRD according to the blueprint.

在此之后，Stash 将根据蓝图创建一个 `Repository` 和一个 `BackupConfiguration` CRD。

**Verify Repository:**

**验证存储库：**

Verify that the `Repository` has been created successfully by the following command:

通过以下命令验证“Repository”是否已成功创建：

```console
$ kubectl get repository -n demo
NAME                         INTEGRITY   SIZE   SNAPSHOT-COUNT   LAST-SUCCESSFUL-BACKUP   AGE
postgres-sample-postgres-1                                                                2m23s

```

If we view the YAML of this `Repository`, we are going to see that the variables `${TARGET_NAMESPACE}`, `${TARGET_APP_RESOURCE}` and `${TARGET_NAME}` has been replaced by `demo`, `postgres` and `sample-postgres-1` respectively.

如果我们查看这个 `Repository` 的 YAML，我们将看到变量 `${TARGET_NAMESPACE}`、`${TARGET_APP_RESOURCE}` 和 `${TARGET_NAME}` 已经被 `demo`、`postgres` 替换和 `sample-postgres-1` 分别。

```console
$ kubectl get repository -n demo postgres-sample-postgres-1 -o yaml

```

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: Repository
metadata:
creationTimestamp: "2019-08-01T13:54:48Z"
finalizers:
  - stash
generation: 1
name: postgres-sample-postgres-1
namespace: demo
resourceVersion: "50171"
selfLink: /apis/stash.appscode.com/v1beta1/namespaces/demo/repositories/postgres-sample-postgres-1
uid: ed49dde4-b463-11e9-a6a0-080027aded7e
spec:
backend:
    gcs:
      bucket: appscode-qa
      prefix: stash-backup/demo/postgres/sample-postgres-1
    storageSecretName: gcs-secret

```

**Verify BackupConfiguration:**

**验证备份配置：**

Verify that the `BackupConfiguration` CRD has been created by the following command:

验证是否已通过以下命令创建了 `BackupConfiguration` CRD：

```console
$ kubectl get backupconfiguration -n demo
NAME                         TASK                   SCHEDULE      PAUSED   AGE
postgres-sample-postgres-1   postgres-backup-11.2   */5 * * * *            3m39s

```

Notice the `TASK` field. It denotes that this backup will be performed using `postgres-backup-11.2` task. We had specified `postgres-backup-${TARGET_APP_VERSION}` as task name in the `BackupBlueprint`. Here, the variable `${TARGET_APP_VERSION}` has been substituted by the database version.

注意“任务”字段。它表示此备份将使用 `postgres-backup-11.2` 任务执行。我们在“BackupBlueprint”中指定了“postgres-backup-${TARGET_APP_VERSION}”作为任务名称。此处，变量`${TARGET_APP_VERSION}` 已被数据库版本替换。

Let’s check the YAML of this `BackupConfiguration`.

让我们检查一下这个 `BackupConfiguration` 的 YAML。

```console
$ kubectl get backupconfiguration -n demo postgres-sample-postgres-1 -o yaml

```

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: BackupConfiguration
metadata:
creationTimestamp: "2019-08-01T13:54:48Z"
finalizers:
  - stash.appscode.com
generation: 1
name: postgres-sample-postgres-1
namespace: demo
ownerReferences:
  - apiVersion: v1
    blockOwnerDeletion: false
    kind: AppBinding
    name: sample-postgres-1
    uid: a799156e-b463-11e9-a6a0-080027aded7e
resourceVersion: "50170"
selfLink: /apis/stash.appscode.com/v1beta1/namespaces/demo/backupconfigurations/postgres-sample-postgres-1
uid: ed4bd257-b463-11e9-a6a0-080027aded7e
spec:
repository:
    name: postgres-sample-postgres-1
retentionPolicy:
    keepLast: 5
    name: keep-last-5
    prune: true
runtimeSettings: {}
schedule: '*/5 * * * *'
target:
    ref:
      apiVersion: v1
      kind: AppBinding
      name: sample-postgres-1
task:
    name: postgres-backup-11.2
tempDir: {}

```

Notice that the `spec.target.ref` is pointing to the AppBinding `sample-postgres-1` that we have just annotated with auto-backup annotation.

请注意，`spec.target.ref` 指向我们刚刚使用自动备份注释进行注释的 AppBinding `sample-postgres-1`。

**Wait for BackupSession:**

**等待备份会话：**

Now, wait for the next backup schedule. Run the following command to watch `BackupSession` CRD:

现在，等待下一个备份计划。运行以下命令来观察`BackupSession` CRD：

```console
$ watch -n 1 kubectl get backupsession -n demo -l=stash.appscode.com/backup-configuration=postgres-sample-postgres-1

Every 1.0s: kubectl get backupsession -n demo -l=stash.appscode.com/backup-configuration=postgres-sample-postgres-1  workstation: Thu Aug  1 20:35:43 2019

NAME                                    INVOKER-TYPE          INVOKER-NAME                 PHASE       AGE
postgres-sample-postgres-1-1564670101   BackupConfiguration   postgres-sample-postgres-1   Succeeded   42s

```

Note: Backup CronJob creates `BackupSession` CRD with the following label `stash.appscode.com/backup-configuration=<BackupConfiguration crd name>`. We can use this label to watch only the `BackupSession` of our desired `BackupConfiguration`.

注意：Backup CronJob 创建了带有以下标签的 `BackupSession` CRD，标签为 `stash.appscode.com/backup-configuration=<BackupConfiguration crd name>`。我们可以使用此标签仅查看所需的“BackupConfiguration”的“BackupSession”。

**Verify Backup:**

**验证备份：**

When backup session is completed, Stash will update the respective `Repository` to reflect the latest state of backed up data.

备份会话完成后，Stash 将更新相应的“存储库”以反映备份数据的最新状态。

Run the following command to check if a snapshot has been sent to the backend:

运行以下命令检查快照是否已发送到后端：

```console
$ kubectl get repository -n demo postgres-sample-postgres-1
NAME                         INTEGRITY   SIZE        SNAPSHOT-COUNT   LAST-SUCCESSFUL-BACKUP   AGE
postgres-sample-postgres-1   true        1.324 KiB   1                73s                      6m7s

```

If we navigate to `stash-backup/demo/postgres/sample-postgres-1` directory of our GCS bucket, we are going to see that the snapshot has been stored there.

如果我们导航到 GCS 存储桶的 `stash-backup/demo/postgres/sample-postgres-1` 目录，我们将看到快照已存储在那里。

![](https://appfleet-com.cdn.ampproject.org/i/s/appfleet.com/blog/content/images/2020/03/sample_postgres_1.png)

### **Backup Second Sample PostgreSQL**

### **备份第二个示例 PostgreSQL**

Now, lets backup our second PostgreSQL sample using the same `BackupBlueprint` we have used to backup the first PostgreSQL sample.

现在，让我们使用我们用来备份第一个 PostgreSQL 样本的相同 `BackupBlueprint` 来备份我们的第二个 PostgreSQL 样本。

**Add Annotations:**

**添加注释：**

Add the auto backup specific annotation to AppBinding `sample-postgres-2`.

将特定于自动备份的注释添加到 AppBinding `sample-postgres-2`。

```console
$ kubectl annotate appbinding sample-postgres-2 -n demo --overwrite \
stash.appscode.com/backup-blueprint=postgres-backup-blueprint

```

**Verify Repository:**

**验证存储库：**

Verify that the `Repository` has been created successfully by the following command:

通过以下命令验证“Repository”是否已成功创建：

```console
$ kubectl get repository -n demo
NAME                         INTEGRITY   SIZE        SNAPSHOT-COUNT   LAST-SUCCESSFUL-BACKUP   AGE
postgres-sample-postgres-1   true        1.324 KiB   1                2m3s                     6m57s
postgres-sample-postgres-2                                                                     15s

```

Here, repository `postgres-sample-postgres-2` has been created for the second PostgreSQL sample. 

在这里，已为第二个 PostgreSQL 示例创建了存储库 `postgres-sample-postgres-2`。

If we view the YAML of this `Repository`, we will see that the variables `${TARGET_NAMESPACE}`, `${TARGET_APP_RESOURCE}` and `${TARGET_NAME}` have been replaced by `demo`, `postgres` and ` sample-postgres-2` respectively.

如果我们查看这个 `Repository` 的 YAML，我们会看到变量 `${TARGET_NAMESPACE}`、`${TARGET_APP_RESOURCE}` 和 `${TARGET_NAME}` 已经被替换为 `demo`、`postgres` 和 ` sample-postgres-2` 分别。

```console
$ kubectl get repository -n demo postgres-sample-postgres-2 -o yaml

```

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: Repository
metadata:
creationTimestamp: "2019-08-01T14:37:22Z"
finalizers:
  - stash
generation: 1
name: postgres-sample-postgres-2
namespace: demo
resourceVersion: "56103"
selfLink: /apis/stash.appscode.com/v1beta1/namespaces/demo/repositories/postgres-sample-postgres-2
uid: df58523c-b469-11e9-a6a0-080027aded7e
spec:
backend:
    gcs:
      bucket: appscode-qa
      prefix: stash-backup/demo/postgres/sample-postgres-2
    storageSecretName: gcs-secret

```

**Verify BackupConfiguration:**

**验证备份配置：**

Verify that the `BackupConfiguration` CRD has been created by the following command:

验证是否已通过以下命令创建了 `BackupConfiguration` CRD：

```console
$ kubectl get backupconfiguration -n demo
NAME                         TASK                   SCHEDULE      PAUSED   AGE
postgres-sample-postgres-1   postgres-backup-11.2   */5 * * * *            7m52s
postgres-sample-postgres-2   postgres-backup-10.6   */5 * * * *            70s

```

Again, notice the `TASK` field. This time, `${TARGET_APP_VERSION}` has been replaced with `10.6` which is the database version of our second sample.

再次注意“TASK”字段。这一次，`${TARGET_APP_VERSION}` 被替换为 `10.6`，这是我们第二个示例的数据库版本。

**Wait for BackupSession:**

**等待备份会话：**

Now, wait for the next backup schedule. Run the following command to watch `BackupSession` CRD:

现在，等待下一个备份计划。运行以下命令来观察`BackupSession` CRD：

```console
$ watch -n 1 kubectl get backupsession -n demo -l=stash.appscode.com/backup-configuration=postgres-sample-postgres-2
Every 1.0s: kubectl get backupsession -n demo -l=stash.appscode.com/backup-configuration=postgres-sample-postgres-2  workstation: Thu Aug  1 20:55:40 2019

NAME                                    INVOKER-TYPE          INVOKER-NAME                 PHASE       AGE
postgres-sample-postgres-2-1564671303   BackupConfiguration   postgres-sample-postgres-2   Succeeded   37s

```

**Verify Backup:**

**验证备份：**

Run the following command to check if a snapshot has been sent to the backend:

运行以下命令检查快照是否已发送到后端：

```console
$ kubectl get repository -n demo postgres-sample-postgres-2
NAME                         INTEGRITY   SIZE        SNAPSHOT-COUNT   LAST-SUCCESSFUL-BACKUP   AGE
postgres-sample-postgres-2   true        1.324 KiB   1                52s                      19m

```

If we navigate to `stash-backup/demo/postgres/sample-postgres-2` directory of our GCS bucket, we are going to see that the snapshot has been stored there.

如果我们导航到 GCS 存储桶的 `stash-backup/demo/postgres/sample-postgres-2` 目录，我们将看到快照已存储在那里。

![](https://appfleet-com.cdn.ampproject.org/i/s/appfleet.com/blog/content/images/2020/03/sample_postgres_2.png)

## **Cleanup**

##  **清理**

To cleanup the Kubernetes resources created by this tutorial, run:

要清理本教程创建的 Kubernetes 资源，请运行：

```console
kubectl delete -n demo pg/sample-postgres-1
kubectl delete -n demo pg/sample-postgres-2

kubectl delete -n demo repository/postgres-sample-postgres-1
kubectl delete -n demo repository/postgres-sample-postgres-2

kubectl delete -n demo backupblueprint/postgres-backup-blueprint
```

## Final thoughts

##  最后的想法

You've now gotten a deep dive into setting up a Kubernetes-native disaster recovery and backup solution with Stash. You can find a lot of really helpful information on their documentation site [here](https://stash.run/). I hope you gained some educational knowledge from this post and will stay tuned for future tutorials! 

您现在已经深入了解如何使用 Stash 设置 Kubernetes 原生灾难恢复和备份解决方案。您可以在他们的文档站点 [此处](https://stash.run/) 上找到许多非常有用的信息。我希望你从这篇文章中获得了一些教育知识，并会继续关注未来的教程！

