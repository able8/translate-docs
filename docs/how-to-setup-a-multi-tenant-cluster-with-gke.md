## How to setup a multi-tenant cluster with GKE

https://cloudsolutions.academy/solution/how-to-setup-a-multi-tenant-cluster-with-gke/


One of the best practices around development environment is to have one large Kubernetes cluster in a multi-tenant mode for your developers. This brings in cost saving specially when you have many small teams divided along the lines of microservices. Before we dive into understanding how we can setup a cluster in a multi-tenant mode, lets understand pros and cons of not having a single cluster for development:

**Pros**

- Dedicated cluster per developer, assumes more control
- Complete isolation, no dependency on other developer resources

**Cons**

- Can be expensive
- End up with too many clusters (usually each per developer)
- More likely the cluster remains under utilised
- With too many clusters, maintenance becomes overhead

So what are the pros and cons of having single shared cluster:

**Pros**

- Less expensive and more efficient
- Easy integration with common enterprise grade services like logging and monitoring
- Maintenance is easy as its one cluster to manage

**Cons**

- Calls for proper user on-boarding and resource management
- Chances of developer conflict in terms of usage

As part of the solution, you will perform the following steps:

1. Partition cluster based on development teams (using namespaces)
2. Provide access control to team in that namespace
3. Allocate resources to the team namespace
4. Monitor the resource utilisation

> The article assumes you have basic knowledge of configuring Google Cloud project and fair understanding of Google Kubernetes Engine (GKE) service.

## Partition cluster based on development teams (using namespaces)

As a first step, you will setup a decent sized cluster depending on the size of your development workloads. For this use case, you can create a cluster with four nodes each with 1 CPU and about 4 GB RAM. Let’s call the cluster ‘dev-cluster’.

We will assume an ERP application that has many modules like Accounts, E-commerce, HR, SCM, Analytics etc. For the target application, we will currently focus on developing accounts and HR modules. So we setup two teams: Team Accounts and Team HR. Let’s setup a shared development cluster for our teams. You will create separate namespaces per team.

By default, kubernetes cluster will have a default namespace. For our development team, we will create two new namespaces viz. team-accounts and team-hr.

```

kubectl create namespace team-accounts
kubectl create namespace team-hr

```

Now you will create a sample application in each of these namespaces. You will deploy a simple nginx pod in each namespace. The below command will deploy a nginx pod in team-accounts and team-hr namespaces.

```

kubectl run app-acct --image=nginx --namespace=team-accounts
kubectl run app-hr --image=nginx --namespace=team-hr

```

Now that the namespaces and its application are ready, let’s setup an access control using Google IAM and Kubernetes based RBAC

## Provide access control to team in that namespace

You will first create a service account that will act as a ‘developer’ for the team and assign ‘cluster viewer’ IAM role to it. The said role will have access to cluster and namespaced resources.

```

gcloud iam service-accounts create team-accounts-developer \
--description="Developer for team accounts" \
--display-name="accts-developer"
```

```

gcloud projects add-iam-policy-binding ${GOOGLE_CLOUD_PROJECT} \
--member=serviceAccount:team-accounts-developer@${GOOGLE_CLOUD_PROJECT}.iam.gserviceaccount.com  \
--role=roles/container.clusterViewer
```

As a next step, you will create a Kubernetes role that will have basic CRUD operations permission for pods and deployments. You will bind this role to the already created IAM service account. You can create the role using below YAML (developer-role.yaml)

```

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
namespace: team-accounts
name: developer
rules:
- apiGroups: [""]
resources: ["pods", "services", "serviceaccounts"]
verbs: ["update", "create", "delete", "get", "watch", "list"]
- apiGroups:["apps"]
resources: ["deployments"]
verbs: ["update", "create", "delete", "get", "watch", "list"]

```

```

kubectl create -f developer-role.yaml

```

```

kubectl create rolebinding team-accounts-developer -n team-accounts \
--role=developer --user=team-accounts-developer@${GOOGLE_CLOUD_PROJECT}.iam.gserviceaccount.com

```

To test the access, you have to download the json key of the created service account and login as that service account. Once you are logged with the ‘team-accounts-developer’ service account, try to access the application in team-accounts namespace.

```

kubectl get pods --namespace=team-accounts

```

You will able to view the pods in the team-accounts namespace.

Now try to access the pods in team-hr namespace:

```
kubectl get pods --namespace=team-hr
```

It will throw the error ‘access forbidden’ as the service account has the ‘developer’ role that is only associated with the team-accounts namespace.

## Allocate resources to the team namespace

When you have a cluster shared across your development team, it becomes necessary to allocate specific compute resources for each team so that resources are not over utilised by one team. You will use ResourceQuota Kubernetes object to allocate compute resources for your team namespace. A resource quota can be configured on object counts like setting limits for number of pods, services, total sum of storage space or sum of compute resources like cpu, memory etc.

For this use case you will configure a ResourceQuota with the cap of 2 cpu and 8 Gi memory. Apply the below yaml (quota.yaml)

```

apiVersion: v1
kind: ResourceQuota
metadata:
name: compute-quota
namespace: team-accounts
spec:
hard:
    limits.cpu: "4"
    limits.memory: "12Gi"
    requests.cpu: "2"
    requests.memory: "8Gi"

```

The above code indicates you can have max quota of upto 4 cpus and 16G memory. In this way you can restrict compute usage depending on the size of workloads per team in a namespace. You now go ahead and allocate resource to your pods in the team-accounts namespace within the above specified limits.

## Monitor the resource utilisation

When you use multi-tenant cluster, you may have different configurations and resource quota set for each team of developers. Resource utilisation may change as we proceed with the development and adding more use cases. It is therefore helpful to understand the usage pattern of each team in a namespace via monitoring. With Google Cloud, you can make use of monitoring dashboard for GKE to view compute utilisation based on namespace.

In the Google cloud console, navigate to Operations->Monitoring->Dashboard. This will create the workspace for your project. Once the workspace is created, you can select GKE from the Dashboard Overview page. You can then view CPU and memory utilisation for your namespaces. You can monitor the usage pattern of all the teams in each namepace. You can also make use of Metrics Explorer to view utilisation based on a specific metric on a particular cluster resource.

In summary, with the advent of cloud and various tools it has become easy to provision cluster per developer but one has to take an informed decision when it comes to setting up development environment. One large multi-tenant cluster can significantly reduce cost, also at the same time gives you an opportunity to setup a robust on-boarding process for your developers. Moreover monitoring and dashboard can further enable you to check resource usage for each namespace in the cluster thereby giving you a bigger picture of your project workloads.
