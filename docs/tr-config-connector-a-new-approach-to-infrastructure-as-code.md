# Config Connector, a new approach to Infrastructure as Code

# Config Connector，一种新的基础架构即代码方法

Dec 07, 2022

2022 年 12 月 7 日

Infrastructure as Code (IaC) helps “cloud native” companies manage their infrastructure based on the principles of software engineering. A wide range of IaC tools and frameworks facilitate in updating the cloud infrastructure. Config Connector is the latest member of this family and brings a new approach based on the power of Kubernetes. In this blog post we outline how it works compared to other tools.

基础架构即代码 (IaC) 可帮助“云原生”公司根据软件工程原则管理其基础架构。范围广泛的 IaC 工具和框架有助于更新云基础设施。 Config Connector 是这个家族的最新成员，带来了一种基于 Kubernetes 强大功能的新方法。在这篇博文中，我们概述了它与其他工具相比的工作原理。

As companies expand their infrastructure, creating and enforcing consistent configurations and security policies across a growing environment becomes difficult and creates friction. Infrastructure as Code (IaC) helps solve this by automating through code the configuration and provisioning of resources, so that human error is eliminated, time is saved, and every step is fully documented.

随着公司扩展其基础架构，在不断发展的环境中创建和实施一致的配置和安全策略变得困难并产生摩擦。基础设施即代码 (IaC) 通过代码自动配置和配置资源来帮助解决这个问题，从而消除人为错误，节省时间，并且每个步骤都有完整记录。

IaC applies software engineering practices to infrastructure and brings the same benefits to infrastructure :

IaC 将软件工程实践应用于基础设施，并为基础设施带来同样的好处：

1\. Automate: Commit, version, trace, deploy, and collaborate, just like source code.

2\. Declarative: Specify the desired state of infrastructure, not updates

3\. Roll back: Roll out and roll back changes just like a regular application

4\. Validate: Assess desired state vs. current state infrastructure

5\. Scale: Build reusable infrastructure blocks across an organization



1. 自动化：提交、版本、跟踪、部署和协作，就像源代码一样。
2. 声明式：指定所需的基础设施状态，而不是更新
3. 回滚：像常规应用程序一样推出和回滚更改
4. 验证：评估期望状态与当前状态基础设施
5. 规模：在整个组织中构建可重用的基础设施块



## Iac Tool landscape

### Iac 工具景观

Over the years there’s been an explosion in infrastructure platforms and application frameworks that form the foundation of “cloud native.” The most popular are listed in the table below.

多年来，构成“云原生”基础的基础设施平台和应用程序框架呈爆炸式增长。下表列出了最受欢迎的。

### Config Connector

### 配置连接器

Tools like Terraform and Pulumi let admins declare infrastructure in code. But code does not establish a strong contract between desired and current state, and every time code is modified or refactored, a procedural or imperative approach (think: plan/apply step) is required to revalidate the state.

Terraform 和 Pulumi 等工具让管理员可以在代码中声明基础设施。但是代码并没有在期望状态和当前状态之间建立牢固的契约，并且每次修改或重构代码时，都需要一种程序性或命令性方法（想想：计划/应用步骤）来重新验证状态。

Bring in Kubernetes.

引入 Kubernetes。

Controllers are the core of Kubernetes. It’s a controller’s job to ensure that, for any given object, the actual state of the world matches the desired state in the object.

控制器是 Kubernetes 的核心。控制器的工作是确保对于任何给定对象，世界的实际状态与对象中的所需状态相匹配。

Config Connector extends the [Kubernetes Resource Model](https://github.com/kubernetes/design-proposals-archive/blob/main/architecture/resource-management.md) with [Custom Resource Definitions (CRD)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) for GCP services and resources.

Config Connector 使用 [自定义资源定义 (CRD)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) 用于 GCP 服务和资源。

When you install Config Connector on a Kubernetes cluster, a CRD is defined for every service and resource in GCP :

在 Kubernetes 集群上安装 Config Connector 时，会为 GCP 中的每个服务和资源定义一个 CRD：

```
➜ kubectl get crd --namespace cnrm-system
NAME                                                                               CREATED AT
accesscontextmanageraccesslevels.accesscontextmanager.cnrm.cloud.google.com        2022-11-29T07:41:51Z
accesscontextmanageraccesspolicies.accesscontextmanager.cnrm.cloud.google.com      2022-11-29T07:41:51Z
accesscontextmanagerserviceperimeters.accesscontextmanager.cnrm.cloud.google.com   2022-11-29T07:41:51Z
apigeeenvironments.apigee.cnrm.cloud.google.com                                    2022-11-29T07:41:51Z
apigeeorganizations.apigee.cnrm.cloud.google.com                                   2022-11-29T07:41:51Z
artifactregistryrepositories.artifactregistry.cnrm.cloud.google.com                2022-11-29T07:41:51Z
backendconfigs.cloud.google.com                                                    2022-11-25T13:06:26Z
bigquerydatasets.bigquery.cnrm.cloud.google.com                                    2022-11-29T07:41:51Z
bigqueryjobs.bigquery.cnrm.cloud.google.com                                        2022-11-29T07:41:52Z
bigquerytables.bigquery.cnrm.cloud.google.com                                      2022-11-29T07:41:52Z
bigtableappprofiles.bigtable.cnrm.cloud.google.com                                 2022-11-29T07:41:52Z
bigtablegcpolicies.bigtable.cnrm.cloud.google.com                                  2022-11-29T07:41:52Z
bigtableinstances.bigtable.cnrm.cloud.google.com                                   2022-11-29T07:41:52Z
bigtabletables.bigtable.cnrm.cloud.google.com                                      2022-11-29T07:41:52Z
billingbudgetsbudgets.billingbudgets.cnrm.cloud.google.com                         2022-11-29T07:41:52Z
binaryauthorizationattestors.binaryauthorization.cnrm.cloud.google.com             2022-11-29T07:41:52Z
binaryauthorizationpolicies.binaryauthorization.cnrm.cloud.google.com              2022-11-29T07:41:52Z
capacityrequests.internal.autoscaling.gke.io                                       2022-11-25T13:05:36Z
cloudbuildtriggers.cloudbuild.cnrm.cloud.google.com                                2022-11-29T07:41:52Z
...

```

Admins can now define the desired state of the infrastructure as objects in the Kubernetes `etcd` database using the Config Connector CRDs. The Config Connector `operator` defines a controller for every CRD that will reconcile the actual state of the infrastructure with the desired state of the objects in the Kubernetes `etcd` database as defined by the admins.

管理员现在可以使用 Config Connector CRD 将基础架构的所需状态定义为 Kubernetes `etcd` 数据库中的对象。 Config Connector `operator` 为每个 CRD 定义了一个控制器，它将基础设施的实际状态与管理员定义的 Kubernetes `etcd` 数据库中对象的期望状态相协调。

Config Connector `operator` translates desired declarative state to imperative API calls.

Config Connector `operator` 将所需的声明状态转换为命令式 API 调用。

![Config Connector architecture](https://binx.io/wp-content/uploads/2022/12/config-connector-a-new-approach-to-infrastructure-as-code-configconnector-architecture.png)

### Example

###  例子

Let’s say we want to create a `pubsub` topic.

假设我们要创建一个 `pubsub` 主题。

First we need to enable the `pubsub api` on our project. Based on the [ServiceUsage CRD](https://cloud.google.com/config-connector/docs/reference/resource-docs/serviceusage/service), we can create a YAML file that declares that the service is enabled :

首先，我们需要在我们的项目上启用 `pubsub api`。基于 [ServiceUsage CRD](https://cloud.google.com/config-connector/docs/reference/resource-docs/serviceusage/service)，我们可以创建一个声明服务已启用的 YAML 文件：

```
apiVersion: serviceusage.cnrm.cloud.google.com/v1beta1
kind: Service
metadata:
name: pubsub.googleapis.com
spec:
projectRef:
    external: projects/my-project-id

```

Store this YAML file as `pubsub-service.yaml` and apply it to your Kubernetes cluster:

将此 YAML 文件存储为“pubsub-service.yaml”并将其应用于您的 Kubernetes 集群：

```
kubectl apply -f pubsub-service.yaml

```

Now that the `pubsub` service is enabled, let’s create a topic. Based on the [PubSubTopic CRD](https://cloud.google.com/config-connector/docs/reference/resource-docs/pubsub/pubsubtopic), below YAML file `pubsub-topic.yaml` declares a new `pubsub ` topic :

现在启用了 pubsub 服务，让我们创建一个主题。基于 [PubSubTopic CRD](https://cloud.google.com/config-connector/docs/reference/resource-docs/pubsub/pubsubtopic)，YAML 文件“pubsub-topic.yaml”下方声明了一个新的“pubsub” `主题：

```
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
annotations:
    cnrm.cloud.google.com/project-id: my-project-id
labels:
    managed: configconnector
name: cc-managed-topic

```

And apply the file again :

并再次应用该文件：

```
kubectl apply -f pubsub-topic.yaml

```

Head over to your GCP console and verify that a `pubsub` topic called `cc-managed-topic` is actually created. Congratulations, you have declared your configuration as data !

转到您的 GCP 控制台并验证是否实际创建了名为 cc-managed-topic 的 pubsub 主题。恭喜，您已将配置声明为数据！

But now for the beauty of Kubernetes and continuous reconciliation …

但是现在为了 Kubernetes 的美丽和持续的协调......

Delete the `pubsub` topic manually from your project. The `pubsubtopic-controller` will detect that the actual state is no longer inline with the desired state and will start taking remediation action. Wait a couple of seconds and … Kubernetes will recreate the topic you manually deleted ! And there was no procedural/imperative step required.

从您的项目中手动删除 `pubsub` 主题。 `pubsubtopic-controller` 将检测到实际状态不再符合所需状态，并将开始采取补救措施。等几秒钟……Kubernetes 将重新创建您手动删除的主题！并且不需要任何程序性/强制性步骤。

Check the events on the `pubsub` object to verify what happened :

检查 `pubsub` 对象上的事件以验证发生了什么：

![Kubernetes CRD reconcile events](https://binx.io/wp-content/uploads/2022/12/config-connector-a-new-approach-to-infrastructure-as-code-reconcile.png)

## Conclusion

##  结论

With Google [Config Connector](https://cloud.google.com/config-connector/docs/overview) we can move to a truly declarative approach for infrastructure using Configuration as Data by harnessing the power of Kubernetes.

借助 Google [Config Connector](https://cloud.google.com/config-connector/docs/overview)，我们可以利用 Kubernetes 的强大功能，使用配置即数据，转向真正的声明式基础设施方法。

While Config Connector was released by Google for GCP, we see an adoption of the same principles by other Cloud providers. Microsoft Azure released [Azure Service Operator](https://github.com/Azure/azure-service-operator) and AWS is building around [AWS Controllers for Kubernetes (ACK)](https://github.com/aws-controllers-k8s/community).z

虽然 Config Connector 是由 Google 为 GCP 发布的，但我们看到其他云提供商也采用了相同的原则。 Microsoft Azure 发布了 [Azure Service Operator](https://github.com/Azure/azure-service-operator)，AWS 正在围绕 [AWS Controllers for Kubernetes (ACK)](https://github.com/aws-控制器-k8s/社区)。

And then there is [Crossplane](https://crossplane.io/) that aims to bring a universal control plane to enable platform teams to assemble infrastructure from multiple vendors.

然后是 [Crossplane](https://crossplane.io/)，旨在带来一个通用的控制平面，使平台团队能够组装来自多个供应商的基础设施。

Is this the new industry trend ? Let’s see what the future brings.

这是新的行业趋势吗？让我们看看未来会带来什么。

https://binx.io/2022/12/07/config-connector-a-new-approach-to-infrastructure-as-code

