# Using GitOps to Self-Manage Postgres in Kubernetes

February 01, 2021 [Jonathan S. Katz](https://blog.crunchydata.com/blog/author/jonathan-s-katz)

[PostgreSQL](https://blog.crunchydata.com/blog/topic/postgresql) [Kubernetes](https://blog.crunchydata.com/blog/topic/kubernetes) [PostgreSQL Operator](https://blog.crunchydata.com/blog/topic/postgresql-operator) [GitOps](https://blog.crunchydata.com/blog/topic/gitops)

" [GitOps](https://www.gitops.tech/)" is a term that I've been seeing come up more and more. The concept was first put forward by the team at [Weaveworks](https://www.weave.works/technologies/gitops/) as a way to consolidate thought around deploying applications. In essence: your deployment topology lives in your git repository. You can update your deployment information by adding a new commit. Likewise, if you need to revert your system's state, you can rollback to the commit that you want to represent your production environment. Any changes to your deployment topology should be reconciled in your production environment.

A lot of the conversations around GitOps came around the [Postgres Operator](https://github.com/CrunchyData/postgres-operator) for [Kubernetes](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes/) and how to apply these principles. Platforms like Kubernetes make it relatively seamless to apply the ideas of GitOps to the stateless pieces of applications (e.g., your web application). More work needs to be done with a stateful service such as [PostgreSQL](https://www.postgresql.org/), as each stateful application can have unique requirements. Let's take something like replication: [PostgreSQL's replication system](https://www.postgresql.org/docs/current/high-availability.html) is both configured and managed differently than in other database systems, and any GitOps management tool, such as [Helm](https://helm.sh/) or [Kustomize](https://kustomize.io/), would have to account for this.

This is the beauty of the [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/). It provides a generic framework for developers to allow for GitOps style management of a stateful application in a way that can be "simply" configured. I say "simply" because there is still a lot to consider when running a complex stateful application like a database in a production environment, but an Operator can make the overall management and application of changes easier.

There are many ways to support GitOps workflows in current versions of the PostgreSQL Operator, from Kubernetes [YAML files](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/#create-a-postgresql-cluster) to [Helm charts](https://github.com/CrunchyData/postgres-operator/tree/master/examples/helm) to [Kustomize manifests](https://github.com/CrunchyData/postgres-operator/tree/master/examples/kustomize/createcluster).

Let's work through a GitOps style workflow. The examples below will assume that you have [installed the Postgres Operator](https://access.crunchydata.com/documentation/postgres-operator/latest/quickstart/).

## Deploying a HA PostgreSQL Cluster

For the first example, let's create a [HA PostgreSQL](https://access.crunchydata.com/documentation/postgres-operator/latest/architecture/high-availability/) cluster called hippo using a Kubernetes YAML file. In your command-line environment, set the following environmental variables to where you want your PostgreSQL cluster deployed.

For this example, the namespace uses the same one created as part of the [quickstart](https://access.crunchydata.com/documentation/postgres-operator/latest/quickstart/). You can copy and paste the below into your environment to try the example out. You can also modify the [custom resource](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/) manifest to match your specific environment. Note that if your environment does not support environmental variables, you can find/replace the below values in your manifest file.

From your command line, execute the following:

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

```
kubectl -n "${cluster_namespace}" get pods  --selector="pg-cluster=${pgo_cluster_name},pgo-pg-database"
NAME                                          READY   STATUS      RESTARTS   AGE
hippo-585bb4f797-f76bl                        1/1     Running     0          3m32s
hippo-gftf-7f55674d78-gsjx7                   1/1     Running     0          2m53s
```

Let's log into our newly provisioned database cluster. In our example above, we created a database named hippo along with a user called hippo. The Postgres Operator creates user credentials for several defaults users, including any users we specify. To get the credentials, assuming you still have the environmental variables set from the earlier step, you can use the following command:

```
kubectl -n "${cluster_namespace}" get secrets "${pgo_cluster_name}-${pgo_cluster_name}-secret" -o jsonpath="{.data.password}" | base64 -d
```

For convenience for connecting to the cluster, we can store the user's password directly in an environmental variable that psql recognizes:

```
export PGPASSWORD=$(kubectl -n jkatz get secrets "${pgo_cluster_name}-${pgo_cluster_name}-secret" -o jsonpath="{.data.password}" | base64 -d)
```

In a separate terminal window, open up a port-forward:

```
export cluster_namespace=pgo
export pgo_cluster_name=hippo
kubectl port-forward -n "${cluster_namespace}" "svc/${pgo_cluster_name}" 5432:5432
```

Now, back in the original window, you can now connect to the cluster:

```
psql -h localhost -U "${pgo_cluster_name}" "${pgo_cluster_name}"
psql (13.1)
Type "help" for help.

hippo=>
```

Success!

## Adding More Resources to a Deployed Cluster

Part of the GitOps principle is the ability to modify (or version) a configuration file and have the changes reflected in your environment. This is typical of "Day 2" type of operations, such as requiring more memory / CPU resources as the workload on a database increases.

Let's say we want to raise our memory limits to 2Gi and our CPU limit to 2.0 cores. Open up the file you created in the previous step, and add the following block to the spec:

```
limits:
memory: 2Gi
cpu: 2.0
```

Your file should look similar to this:

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

```
kubectl apply -f "${pgo_cluster_name}-pgcluster.yaml"
```

Upon detecting the change, the Postgres Operator uses a [rolling update strategy](https://access.crunchydata.com/documentation/postgres-operator/latest/architecture/high-availability/#rolling-updates) to apply the resource changes to each PostgreSQL instance in a way that minimizes downtime. Wait a few moments for the changes to be apple, and then run a describe on the Pods to see the changes (output truncated):

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

For a full list of attributes that can be updated, please see the [custom resources](https://access.crunchydata.com/documentation/postgres-operator/latest/custom-resources/) section of the [Postgres Operator documentation](https://access.crunchydata.com/documentation/postgres-operator/).

## Evolution of GitOps and Stateful Services

When coupled with a tool like the [PostgreSQL Operator](https://github.com/CrunchyData/postgres-operator), GitOps principles can be extended to work for stateful services. Applying a GitOps mindset to managing PostgreSQL workloads can help make it easier to manage production-grade Postgres instances on Kubernetes. GitOps principles can certainly make it easier for deploying a range of PostgreSQL topologies, from single instances to multi-zone, fault tolerant clusters.

Upcoming posts will look at how some of the other Kubernetes toolsets can make it even easier to work with the PostgreSQL Operator in a GitOps manner.


Like what you're reading? Stay informed by subscribing for our newsletter!


## Newsletter

Like what you're reading? Stay informed by subscribing for our newsletter!