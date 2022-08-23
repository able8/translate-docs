## How to setup a multi-tenant cluster with GKE

## 如何使用 GKE 设置多租户集群

https://cloudsolutions.academy/solution/how-to-setup-a-multi-tenant-cluster-with-gke/

One of the best practices around development environment is to have one large Kubernetes cluster in a multi-tenant mode for your developers. This brings in cost saving specially when you have many small teams divided along the lines of microservices. Before we dive into understanding how we can setup a cluster in a multi-tenant mode, lets understand pros and cons of not having a single cluster for development:

围绕开发环境的最佳实践之一是为您的开发人员在多租户模式下拥有一个大型 Kubernetes 集群。当您有许多按照微服务划分的小团队时，这尤其会节省成本。在深入了解如何在多租户模式下设置集群之前，让我们先了解一下没有单个集群进行开发的利弊：

**Pros**

**优点**

- Dedicated cluster per developer, assumes more control
- Complete isolation, no dependency on other developer resources

- 每个开发人员专用集群，承担更多控制权
- 完全隔离，不依赖其他开发者资源

**Cons**

**缺点**

- Can be expensive
- End up with too many clusters (usually each per developer)
- More likely the cluster remains under utilised
- With too many clusters, maintenance becomes overhead

- 可能很贵
- 以太多集群结束（通常每个开发人员每个集群）
- 集群更有可能仍未得到充分利用
- 集群过多，维护成为开销

So what are the pros and cons of having single shared cluster:

那么拥有单个共享集群的优缺点是什么：

**Pros**

**优点**

- Less expensive and more efficient
- Easy integration with common enterprise grade services like logging and monitoring
- Maintenance is easy as its one cluster to manage

- 更便宜，更高效
- 与常见的企业级服务（如日志记录和监控）轻松集成
- 维护很容易，因为它的一个集群可以管理

**Cons**

**缺点**

- Calls for proper user on-boarding and resource management
- Chances of developer conflict in terms of usage

- 要求适当的用户入职和资源管理
- 开发人员在使用方面发生冲突的可能性

As part of the solution, you will perform the following steps:

作为解决方案的一部分，您将执行以下步骤：

1. Partition cluster based on development teams (using namespaces)
2. Provide access control to team in that namespace
3. Allocate resources to the team namespace
4. Monitor the resource utilisation

1. 基于开发团队的分区集群（使用命名空间）
2. 为该命名空间中的团队提供访问控制
3. 为团队命名空间分配资源
4. 监控资源利用率

> The article assumes you have basic knowledge of configuring Google Cloud project and fair understanding of Google Kubernetes Engine (GKE) service.

> 本文假设您具备配置 Google Cloud 项目的基本知识，并对 Google Kubernetes Engine (GKE) 服务有一定的了解。

## Partition cluster based on development teams (using namespaces)

## 基于开发团队的分区集群（使用命名空间）

As a first step, you will setup a decent sized cluster depending on the size of your development workloads. For this use case, you can create a cluster with four nodes each with 1 CPU and about 4 GB RAM. Let’s call the cluster ‘dev-cluster’.

作为第一步，您将根据开发工作负载的大小设置一个合适大小的集群。对于此用例，您可以创建一个包含四个节点的集群，每个节点具有 1 个 CPU 和大约 4 GB RAM。我们将集群称为“dev-cluster”。

We will assume an ERP application that has many modules like Accounts, E-commerce, HR, SCM, Analytics etc. For the target application, we will currently focus on developing accounts and HR modules. So we setup two teams: Team Accounts and Team HR. Let’s setup a shared development cluster for our teams. You will create separate namespaces per team.

我们将假设一个 ERP 应用程序具有许多模块，如 Accounts、E-commerce、HR、SCM、Analytics 等。对于目标应用程序，我们目前将重点开发帐户和 HR 模块。所以我们建立了两个团队：团队客户和团队人力资源。让我们为我们的团队设置一个共享开发集群。您将为每个团队创建单独的命名空间。

By default, kubernetes cluster will have a default namespace. For our development team, we will create two new namespaces viz. team-accounts and team-hr.

默认情况下，kubernetes 集群会有一个默认的命名空间。对于我们的开发团队，我们将创建两个新的命名空间，即。团队帐户和团队小时。

Now you will create a sample application in each of these namespaces. You will deploy a simple nginx pod in each namespace. The below command will deploy a nginx pod in team-accounts and team-hr namespaces.

现在，您将在每个命名空间中创建一个示例应用程序。您将在每个命名空间中部署一个简单的 nginx pod。下面的命令将在 team-accounts 和 team-hr 命名空间中部署一个 nginx pod。

Now that the namespaces and its application are ready, let’s setup an access control using Google IAM and Kubernetes based RBAC.

现在命名空间及其应用程序已准备就绪，让我们使用基于 Google IAM 和 Kubernetes 的 RBAC 设置访问控制

## Provide access control to team in that namespace

## 为该命名空间中的团队提供访问控制

You will first create a service account that will act as a ‘developer’ for the team and assign ‘cluster viewer’ IAM role to it. The said role will have access to cluster and namespaced resources.

您将首先创建一个服务帐户，该帐户将充当团队的“开发人员”，并为其分配“集群查看器”IAM 角色。所述角色将有权访问集群和命名空间资源。

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

下一步，您将创建一个 Kubernetes 角色，该角色将具有对 pod 和部署的基本 CRUD 操作权限。您将将此角色绑定到已创建的 IAM 服务账户。您可以使用以下 YAML (developer-role.yaml) 创建角色

To test the access, you have to download the json key of the created service account and login as that service account. Once you are logged with the ‘team-accounts-developer’ service account, try to access the application in team-accounts namespace.

要测试访问，您必须下载创建的服务帐户的 json 密钥并以该服务帐户的身份登录。使用“team-accounts-developer”服务帐户登录后，尝试访问 team-accounts 命名空间中的应用程序。

You will able to view the pods in the team-accounts namespace.

您将能够查看 team-accounts 命名空间中的 pod。

Now try to access the pods in team-hr namespace:

现在尝试访问 team-hr 命名空间中的 pod：

```
kubectl get pods --namespace=team-hr
```

It will throw the error ‘access forbidden’ as the service account has the ‘developer’ role that is only associated with the team-accounts namespace.

它将抛出错误“禁止访问”，因为服务帐户具有仅与团队帐户命名空间相关联的“开发者”角色。

## Allocate resources to the team namespace

## 将资源分配给团队命名空间

When you have a cluster shared across your development team, it becomes necessary to allocate specific compute resources for each team so that resources are not over utilised by one team. You will use ResourceQuota Kubernetes object to allocate compute resources for your team namespace. A resource quota can be configured on object counts like setting limits for number of pods, services, total sum of storage space or sum of compute resources like cpu, memory etc.

当您的开发团队共享一个集群时，有必要为每个团队分配特定的计算资源，这样资源就不会被一个团队过度利用。您将使用 ResourceQuota Kubernetes 对象为您的团队命名空间分配计算资源。可以根据对象计数配置资源配额，例如设置 pod 数量、服务、存储空间总和或计算资源总和（如 cpu、内存等）的限制。

For this use case you will configure a ResourceQuota with the cap of 2 cpu and 8 Gi memory. Apply the below yaml (quota.yaml)

对于这个用例，您将配置一个 ResourceQuota，其上限为 2 cpu 和 8 Gi 内存。应用以下 yaml (quota.yaml)

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

上面的代码表明您最多可以有 4 个 CPU 和 16G 内存的最大配额。通过这种方式，您可以根据命名空间中每个团队的工作负载大小来限制计算使用量。您现在继续在上面指定的限制内将资源分配给 team-accounts 命名空间中的 pod。

## Monitor the resource utilisation

## 监控资源利用率

When you use multi-tenant cluster, you may have different configurations and resource quota set for each team of developers. Resource utilisation may change as we proceed with the development and adding more use cases. It is therefore helpful to understand the usage pattern of each team in a namespace via monitoring. With Google Cloud, you can make use of monitoring dashboard for GKE to view compute utilisation based on namespace.

当您使用多租户集群时，您可能会为每个开发团队设置不同的配置和资源配额。随着我们继续开发和添加更多用例，资源利用率可能会发生变化。因此，通过监控了解命名空间中每个团队的使用模式很有帮助。借助 Google Cloud，您可以使用 GKE 的监控仪表板来查看基于命名空间的计算利用率。

In the Google cloud console, navigate to Operations->Monitoring->Dashboard. This will create the workspace for your project. Once the workspace is created, you can select GKE from the Dashboard Overview page. You can then view CPU and memory utilisation for your namespaces. You can monitor the usage pattern of all the teams in each namepace. You can also make use of Metrics Explorer to view utilisation based on a specific metric on a particular cluster resource.

在谷歌云控制台中，导航到操作->监控->仪表板。这将为您的项目创建工作区。创建工作区后，您可以从仪表板概览页面中选择 GKE。然后，您可以查看命名空间的 CPU 和内存利用率。您可以监控每个命名空间中所有团队的使用模式。您还可以使用 Metrics Explorer 根据特定集群资源的特定指标查看利用率。

In summary, with the advent of cloud and various tools it has become easy to provision cluster per developer but one has to take an informed decision when it comes to setting up development environment. One large multi-tenant cluster can significantly reduce cost, also at the same time gives you an opportunity to setup a robust on-boarding process for your developers. Moreover monitoring and dashboard can further enable you to check resource usage for each namespace in the cluster thereby giving you a bigger picture of your project workloads. 

总之，随着云和各种工具的出现，为每个开发人员配置集群变得很容易，但在设置开发环境时必须做出明智的决定。一个大型多租户集群可以显着降低成本，同时让您有机会为您的开发人员设置一个强大的入职流程。此外，监控和仪表板可以进一步让您检查集群中每个命名空间的资源使用情况，从而让您更全面地了解项目工作负载。

