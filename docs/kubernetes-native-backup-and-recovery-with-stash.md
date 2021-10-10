# Tutorial: Kubernetes-Native Backup and Recovery With Stash

![Tutorial: Kubernetes-Native Backup and Recovery With Stash](https://appfleet-com.cdn.ampproject.org/i/s/appfleet.com/blog/content/images/size/w1200/2020/07/91-Tutorial-Kubernetes-Native-Backup-and-Recovery-With-Stash.png)

* * *

## Intro

Having a proper backup recovery plan is vital to any organization's IT operation. However, when you begin to distribute workloads across data centers and regions, that process begins to become more and more complex. Container orchestration platforms such as Kubernetes have begun to ease this burden and enabled the management of distributed workloads in areas that were previously very challenging.

In this post, we are going to introduce you to a Kubernetes-native tool for taking backups of your disks, helping with the crucial recovery plan. **Stash is a [Restic](https://restic.net/) Operator that accelerates the task of backing up and recovering your Kubernetes infrastructure**. You can read more about the Operator Framework via [this blog post](https://appfleet.com/blog/first-steps-with-the-kubernetes-operator/).

## How does Stash work?

Using Stash, you can backup Kubernetes volumes mounted in following types of workloads:

- Deployment
- DaemonSet
- ReplicaSet
- ReplicationController
- StatefulSet

At the heart of Stash is a Kubernetes [controller](https://book.kubebuilder.io/basics/what_is_a_controller.html) which uses [Custom Resource Definition (CRD)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) to specify targets and behaviors of the backup and restore process in a Kubernetes native way. A simplified architecture of Stash is shown below:

![](https://appfleet-com.cdn.ampproject.org/i/s/appfleet.com/blog/content/images/2020/03/stash_architecture.svg)

## Installing Stash

### Using Helm 3

Stash can be installed via [Helm](https://helm.sh/) using the [chart](https://github.com/stashed/installer/tree/v0.9.0-rc.6/charts/stash) from [AppsCode Charts Repository](https://github.com/appscode/charts). To install the chart with the release name `stash-operator`:

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

If you prefer to not use Helm, you can generate YAMLs from Stash chart and deploy using `kubectl`:

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/stash --version v0.9.0-rc.6
NAME            CHART VERSION APP VERSION DESCRIPTION
appscode/stash  v0.9.0-rc.6    v0.9.0-rc.6  Stash by AppsCode - Backup your Kubernetes Volumes

$ helm template stash-operator appscode/stash \
  --version v0.9.0-rc.6 \
  --namespace kube-system \
  --no-hooks | kubectl apply -f -
```

### **Installing on GKE Cluster**

If you are installing Stash on a GKE cluster, you will need cluster admin permissions to install Stash operator. Run the following command to grant admin permission to the cluster.

```console
$ kubectl create clusterrolebinding "cluster-admin-$(whoami)" \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"

```

In addition, if your GKE cluster is a [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters), you will need to either add an additional firewall rule that allows master nodes access port `8443/tcp` on worker nodes, or change the existing rule that allows access to ports `443/tcp` and `10250/tcp` to also allow access to port `8443/tcp`. The procedure to add or modify firewall rules is described in the official GKE documentation for private clusters mentioned above.

### Verify installation

To check if Stash operator pods have started, run the following command:

```console
$ kubectl get pods --all-namespaces -l app=stash --watch

NAMESPACE     NAME                              READY     STATUS    RESTARTS   AGE
kube-system   stash-operator-859d6bdb56-m9br5   2/2       Running   2          5s

```

Once the operator pods are running, you can cancel the above command by typing `Ctrl+C`.

Now, to confirm CRD groups have been registered by the operator, run the following command:

```console
$ kubectl get crd -l app=stash

NAME                                 AGE
recoveries.stash.appscode.com        5s
repositories.stash.appscode.com      5s
restics.stash.appscode.com           5s

```

With this, you are ready to take your first backup using Stash.

## Configuring Auto Backup for Database

To keep everything isolated, we are going to use a separate namespace called `demo` throughout this tutorial.

```console
$ kubectl create ns demo
namespace/demo created
```

### Prepare Backup Blueprint

We are going to use [GCS Backend](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/gcs) to store the backed up data. You can use any supported backend you prefer. You just have to configure Storage Secret and `spec.backend` section of `BackupBlueprint` to match with your backend. Visit [here](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/overview) to learn which backends are supported by Stash and how to configure them.

For GCS backend, if the bucket does not exist, Stash needs `Storage Object Admin` role permissions to create the bucket. For more details, please check the following [guide](https://stash.run/docs/v0.9.0-rc.6/guides/latest/backends/gcs).

**Create Storage Secret:**

At first, let’s create a Storage Secret for the GCS backend,

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

Next, we have to create a `BackupBlueprint` CRD with a blueprint for `Repository` and `BackupConfiguration` object.

Below is the YAML of the `BackupBlueprint` object that we are going to create:

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

Let’s create the `BackupBlueprint` that we have shown above.

```console
$ kubectl apply -f https://github.com/stashed/docs/raw/v0.9.0-rc.6/docs/examples/guides/latest/auto-backup/database/backupblueprint.yaml
backupblueprint.stash.appscode.com/postgres-backup-blueprint created

```

With this, automatic backup is configured for PostgreSQL database. We just have to add an annotation to the `AppBinding` of the targeted database.

**Required Annotation for Auto-Backup Database:**

You have to add the following annotation to the `AppBinding` CRD of the targeted database to enable backup for it:

```yaml
stash.appscode.com/backup-blueprint: <BackupBlueprint name>

```

This annotation specifies the name of the `BackupBlueprint` object where a blueprint for `Repository` and `BackupConfiguration` has been defined.

### Prepare Databases

Next, we are going to deploy two sample PostgreSQL databases of two different versions using KubeDB. We are going to backup these two databases using auto-backup.

**Deploy First PostgreSQL Sample:**

Below is the YAML of the first `Postgres` CRD:

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

```console
$ kubectl apply -f https://github.com/stashed/docs/raw/v0.9.0-rc.6/docs/examples/guides/latest/auto-backup/database/sample-postgres-1.yaml
postgres.kubedb.com/sample-postgres-1 created

```

KubeDB will deploy a PostgreSQL database according to the above specification and it will create the necessary secrets and services to access the database. It will also create an `AppBinding` CRD that holds the necessary information to connect with the database.

Verify that an `AppBinding` has been created for this PostgreSQL sample:

```console
$ kubectl get appbinding -n demo
NAME                AGE
sample-postgres-1   47s

```

If you view the YAML of this `AppBinding`, you will see it holds service and secret information. Stash uses this information to connect with the database.

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

Below is the YAML of the second `Postgres` object:

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

```console
$ kubectl apply -f https://github.com/stashed/docs/raw/v0.9.0-rc.6/docs/examples/guides/latest/auto-backup/database/sample-postgres-2.yaml
postgres.kubedb.com/sample-postgres-2 created

```

Verify that an `AppBinding` has been created for this PostgreSQL database:

```console
$ kubectl get appbinding -n demo
NAME                AGE
sample-postgres-1   2m49s
sample-postgres-2   10s

```

Here, we can see `AppBinding` `sample-postgres-2` has been created for our second PostgreSQL sample.

## **Backup**

Next, we are going to add auto-backup specific annotation to the `AppBinding` of our desired database. Stash watches for `AppBinding` CRD. Once it finds an `AppBinding` with auto-backup annotation, it will create a `Repository` and a `BackupConfiguration` CRD according to respective `BackupBlueprint`. Then, rest of the backup process will proceed as normal database backup as described [here](https://stash.run/docs/v0.9.0-rc.6/guides/latest/addons/overview).

### **Backup First PostgreSQL Sample**

Let’s backup our first PostgreSQL sample using auto-backup.

**Add Annotations:**

At first, add the auto-backup specific annotation to the AppBinding `sample-postgres-1`:

```console
$ kubectl annotate appbinding sample-postgres-1 -n demo --overwrite \
stash.appscode.com/backup-blueprint=postgres-backup-blueprint

```

Verify that the annotation has been added successfully:

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

**Verify Repository:**

Verify that the `Repository` has been created successfully by the following command:

```console
$ kubectl get repository -n demo
NAME                         INTEGRITY   SIZE   SNAPSHOT-COUNT   LAST-SUCCESSFUL-BACKUP   AGE
postgres-sample-postgres-1                                                                2m23s

```

If we view the YAML of this `Repository`, we are going to see that the variables `${TARGET_NAMESPACE}`, `${TARGET_APP_RESOURCE}` and `${TARGET_NAME}` has been replaced by `demo`, `postgres` and `sample-postgres-1` respectively.

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

Verify that the `BackupConfiguration` CRD has been created by the following command:

```console
$ kubectl get backupconfiguration -n demo
NAME                         TASK                   SCHEDULE      PAUSED   AGE
postgres-sample-postgres-1   postgres-backup-11.2   */5 * * * *            3m39s

```

Notice the `TASK` field. It denotes that this backup will be performed using `postgres-backup-11.2` task. We had specified `postgres-backup-${TARGET_APP_VERSION}` as task name in the `BackupBlueprint`. Here, the variable `${TARGET_APP_VERSION}` has been substituted by the database version.

Let’s check the YAML of this `BackupConfiguration`.

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

**Wait for BackupSession:**

Now, wait for the next backup schedule. Run the following command to watch `BackupSession` CRD:

```console
$ watch -n 1 kubectl get backupsession -n demo -l=stash.appscode.com/backup-configuration=postgres-sample-postgres-1

Every 1.0s: kubectl get backupsession -n demo -l=stash.appscode.com/backup-configuration=postgres-sample-postgres-1  workstation: Thu Aug  1 20:35:43 2019

NAME                                    INVOKER-TYPE          INVOKER-NAME                 PHASE       AGE
postgres-sample-postgres-1-1564670101   BackupConfiguration   postgres-sample-postgres-1   Succeeded   42s

```

Note: Backup CronJob creates `BackupSession` CRD with the following label `stash.appscode.com/backup-configuration=<BackupConfiguration crd name>`. We can use this label to watch only the `BackupSession` of our desired `BackupConfiguration`.

**Verify Backup:**

When backup session is completed, Stash will update the respective `Repository` to reflect the latest state of backed up data.

Run the following command to check if a snapshot has been sent to the backend:

```console
$ kubectl get repository -n demo postgres-sample-postgres-1
NAME                         INTEGRITY   SIZE        SNAPSHOT-COUNT   LAST-SUCCESSFUL-BACKUP   AGE
postgres-sample-postgres-1   true        1.324 KiB   1                73s                      6m7s

```

If we navigate to `stash-backup/demo/postgres/sample-postgres-1` directory of our GCS bucket, we are going to see that the snapshot has been stored there.

![](https://appfleet-com.cdn.ampproject.org/i/s/appfleet.com/blog/content/images/2020/03/sample_postgres_1.png)

### **Backup Second Sample PostgreSQL**

Now, lets backup our second PostgreSQL sample using the same `BackupBlueprint` we have used to backup the first PostgreSQL sample.

**Add Annotations:**

Add the auto backup specific annotation to AppBinding `sample-postgres-2`.

```console
$ kubectl annotate appbinding sample-postgres-2 -n demo --overwrite \
stash.appscode.com/backup-blueprint=postgres-backup-blueprint

```

**Verify Repository:**

Verify that the `Repository` has been created successfully by the following command:

```console
$ kubectl get repository -n demo
NAME                         INTEGRITY   SIZE        SNAPSHOT-COUNT   LAST-SUCCESSFUL-BACKUP   AGE
postgres-sample-postgres-1   true        1.324 KiB   1                2m3s                     6m57s
postgres-sample-postgres-2                                                                     15s

```

Here, repository `postgres-sample-postgres-2` has been created for the second PostgreSQL sample.

If we view the YAML of this `Repository`, we will see that the variables `${TARGET_NAMESPACE}`, `${TARGET_APP_RESOURCE}` and `${TARGET_NAME}` have been replaced by `demo`, `postgres` and `sample-postgres-2` respectively.

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

Verify that the `BackupConfiguration` CRD has been created by the following command:

```console
$ kubectl get backupconfiguration -n demo
NAME                         TASK                   SCHEDULE      PAUSED   AGE
postgres-sample-postgres-1   postgres-backup-11.2   */5 * * * *            7m52s
postgres-sample-postgres-2   postgres-backup-10.6   */5 * * * *            70s

```

Again, notice the `TASK` field. This time, `${TARGET_APP_VERSION}` has been replaced with `10.6` which is the database version of our second sample.

**Wait for BackupSession:**

Now, wait for the next backup schedule. Run the following command to watch `BackupSession` CRD:

```console
$ watch -n 1 kubectl get backupsession -n demo -l=stash.appscode.com/backup-configuration=postgres-sample-postgres-2
Every 1.0s: kubectl get backupsession -n demo -l=stash.appscode.com/backup-configuration=postgres-sample-postgres-2  workstation: Thu Aug  1 20:55:40 2019

NAME                                    INVOKER-TYPE          INVOKER-NAME                 PHASE       AGE
postgres-sample-postgres-2-1564671303   BackupConfiguration   postgres-sample-postgres-2   Succeeded   37s

```

**Verify Backup:**

Run the following command to check if a snapshot has been sent to the backend:

```console
$ kubectl get repository -n demo postgres-sample-postgres-2
NAME                         INTEGRITY   SIZE        SNAPSHOT-COUNT   LAST-SUCCESSFUL-BACKUP   AGE
postgres-sample-postgres-2   true        1.324 KiB   1                52s                      19m

```

If we navigate to `stash-backup/demo/postgres/sample-postgres-2` directory of our GCS bucket, we are going to see that the snapshot has been stored there.

![](https://appfleet-com.cdn.ampproject.org/i/s/appfleet.com/blog/content/images/2020/03/sample_postgres_2.png)

## **Cleanup**

To cleanup the Kubernetes resources created by this tutorial, run:

```console
kubectl delete -n demo pg/sample-postgres-1
kubectl delete -n demo pg/sample-postgres-2

kubectl delete -n demo repository/postgres-sample-postgres-1
kubectl delete -n demo repository/postgres-sample-postgres-2

kubectl delete -n demo backupblueprint/postgres-backup-blueprint
```

## Final thoughts

You've now gotten a deep dive into setting up a Kubernetes-native disaster recovery and backup solution with Stash. You can find a lot of really helpful information on their documentation site [here](https://stash.run/). I hope you gained some educational knowledge from this post and will stay tuned for future tutorials!
