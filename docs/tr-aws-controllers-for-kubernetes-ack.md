# Introducing the AWS Controllers for Kubernetes (ACK)

# 介绍适用于 Kubernetes 的 AWS 控制器 (ACK)

by Jay Pipes, Michael Hausenblas, and Nathan Taber \| on
19 AUG 2020

AWS Controllers for Kubernetes (ACK) is a new tool that lets you directly manage AWS services from Kubernetes. ACK makes it simple to build scalable and highly-available Kubernetes applications that utilize AWS services.

AWS Controllers for Kubernetes (ACK) 是一种新工具，可让您直接从 Kubernetes 管理 AWS 服务。 ACK 使构建可扩展且高度可用的 Kubernetes 应用程序变得简单，这些应用程序利用 AWS 服务。

Today, ACK is available as a **developer preview** on [GitHub](https://github.com/aws-controllers-k8s/community).

今天，ACK 在 [GitHub](https://github.com/aws-controllers-k8s/community) 上作为**开发者预览**提供。

In this post we will give you a brief introduction to the history of the ACK project, show you how ACK works, and how you can start to use the ACK or contribute.

在这篇文章中，我们将向您简要介绍 ACK 项目的历史，向您展示 ACK 的工作原理，以及如何开始使用 ACK 或做出贡献。

## How did we get here?

##  我们是怎么来到这里的？

In late 2018, [Chris Hein](https://twitter.com/christopherhein) [introduced](https://aws.amazon.com/blogs/opensource/aws-service-operator-kubernetes-available/) the AWS Service Operator as an experimental personal project. We reviewed the feedback from the community and internal stakeholders and [decided](https://github.com/aws/containers-roadmap/issues/456) to relaunch as a first-tier open source project. In this process, we renamed the project to AWS Controllers for Kubernetes (ACK). The tenets we put forward are:

2018 年底，[Chris Hein](https://twitter.com/christopherhein)[介绍](https://aws.amazon.com/blogs/opensource/aws-service-operator-kubernetes-available/) AWS服务运营商作为一个实验性的个人项目。我们审查了社区和内部利益相关者的反馈，并[决定](https://github.com/aws/containers-roadmap/issues/456) 作为一级开源项目重新启动。在此过程中，我们将项目重命名为 AWS Controllers for Kubernetes (ACK)。我们提出的宗旨是：

- ACK is a community-driven project, based on a governance model defining roles and responsibilities.
- ACK is optimized for production usage with full test coverage including performance and scalability test suites.
- ACK strives to be the only code base exposing AWS services via a Kubernetes operator.

- ACK 是一个社区驱动的项目，基于定义角色和职责的治理模型。
- ACK 针对生产使用进行了优化，具有完整的测试范围，包括性能和可扩展性测试套件。
- ACK 努力成为唯一通过 Kubernetes 运营商公开 AWS 服务的代码库。

Over the past year, we have significantly evolved the project's [design](https://github.com/aws/aws-controllers-k8s/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Adesign), continued the discussion with internal stakeholders (more in a moment why this is important), and reviewed related projects in the space. A special shout-out in this context to the [Crossplane](https://crossplane.io) project which does an awesome job for cross-cloud use cases and deservedly became an CNCF project in the meantime.

在过去的一年里，我们显着改进了项目的[设计](https://github.com/aws/aws-controllers-k8s/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Adesign)，继续与内部利益相关者进行讨论（稍后会详细说明为什么这很重要)，并审查了该领域的相关项目。在这种情况下，特别感谢 [Crossplane](https://crossplane.io) 项目，该项目在跨云用例方面做得非常出色，同时当之无愧地成为了 CNCF 项目。

The new ACK project continues the spirit of the original AWS Service Operator, but with a few updates:

新的 ACK 项目延续了原始 AWS 服务运营商的精神，但有一些更新：

- AWS cloud resources are managed directly through the AWS APIs instead of CloudFormation. This allows Kubernetes to be the single ‘source of truth’ for a resources desired state.
- Code for the controllers and custom resource definitions is automatically generated from the AWS Go SDK, with human editing and approval. This allows us to support more services with less manual work and keep the project up-to-date with the latest innovations.
- This is an official project built and maintained by the AWS Kubernetes team. We plan to continue investing in this project in conjunction with our colleagues across AWS.

- AWS 云资源直接通过 AWS API 而不是 CloudFormation 进行管理。这使得 Kubernetes 成为资源所需状态的唯一“真实来源”。
- 控制器和自定义资源定义的代码是从 AWS Go SDK 自动生成的，需要人工编辑和批准。这使我们能够以更少的手动工作支持更多服务，并使项目与最新的创新保持同步。
- 这是一个由 AWS Kubernetes 团队构建和维护的官方项目。我们计划与我们在 AWS 的同事一起继续投资这个项目。

## How ACK works

## ACK 是如何工作的

Our goal with ACK to provide a consistent Kubernetes interface for AWS, regardless of the AWS service API. One example of this is ensuring field names and identifiers are normalized and tags are handled the same way across AWS resources.

我们与 ACK 的目标是为 AWS 提供一致的 Kubernetes 接口，而不管 AWS 服务 API。这方面的一个例子是确保字段名称和标识符被规范化，并且标签在 AWS 资源中以相同的方式处理。

![](https://d2908q01vomqb2.cloudfront.net/fe2ef495a1152561572949784c16bf23abb28057/2020/08/14/feature-img-1024x862.png)

As depicted above, from a high level, the ACK workflow is as follows:

如上所述，从高层次来看，ACK 工作流程如下：

1. We, as in “the project team lead by the authors”, generate and maintain a collection of artifacts (binaries, container images, Helm charts, etc.). These artifacts are automatically derived from the AWS services APIs and represent the business logic of how to manage AWS resources from within Kubernetes.
2. As a cluster admin you select one or more ACK controllers you want to install and configure for a cluster you’re responsible.
3. As an application developer, you create (Kubernetes)[custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) representing AWS resources.
4. The respective ACK controller (installed in the step 2.) manages said custom resources and with it the underlying AWS resources. Based on the custom resource defined in the step 3., the controller creates, updates, or deletes the underlying AWS resources using the AWS APIs.

1. 我们作为“作者领导的项目团队”，生成并维护一个工件集合（二进制文件、容器镜像、Helm 图表等）。这些构件自动派生自 AWS 服务 API，代表了如何从 Kubernetes 中管理 AWS 资源的业务逻辑。
2. 作为集群管理员，您可以选择一个或多个要为您负责的集群安装和配置的 ACK 控制器。
3. 作为应用程序开发人员，您创建 (Kubernetes)[自定义资源](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) 代表 AWS 资源。
4. 相应的 ACK 控制器（在步骤 2 中安装）管理所述自定义资源以及底层 AWS 资源。根据步骤 3. 中定义的自定义资源，控制器使用 AWS API 创建、更新或删除底层 AWS 资源。

Let’s now have a closer look at the workflow, using a concrete example.

现在让我们用一个具体的例子来仔细看看工作流程。

### 1\. Generation of artifacts 

### 1. 工件的产生

[Artifacts generation](https://aws-controllers-k8s.github.io/community/dev-docs/code-generation/) creates the required code components that allow you to manage AWS services using Kubernetes. We took an multi-phased approach, yielding hybrid custom+controller-runtimes:

[工件生成](https://aws-controllers-k8s.github.io/community/dev-docs/code-generation/) 创建所需的代码组件，让您可以使用 Kubernetes 管理 AWS 服务。我们采用了多阶段方法，产生了混合自定义+控制器运行时：

- First, we consume model information from a canonical source of truth about AWS services. We settled on the source of truth as [the model files](https://github.com/aws/aws-sdk-go/tree/master/models/apis) from the `aws/aws-sdk-go` repository . AWS SDKs are regularly updated with all API changes so this is an accurate source of information and closely tracks the service API avilability. In this phase we generate files containing code that exposes Go types for objects and interfaces found there.
- After generating Kubernetes API type definitions for the top-level resources exposed by the AWS API, we then need to generate the interface implementations enabling those top-level resources and type definitions to be used by the Kubernetes runtime package
- Next, we generate the custom resource definition (CRD) configuration files, one for each top-level resource identified in earlier steps. Then, we generate the implementation of the ACK controller for the target service. Along with these controller implementations in Go, this steps also outputs a set of Kubernetes manifests for the`Deployment` and the `ClusterRoleBinding` of the `Role` for the next step.
- Finally, we generate the Kubernetes manifests for a Kubernetes`Role` for the Kubernetes `Deployment` running the respective ACK service controllers. Abiding the least privileges principle, this `Role` needs to be equipped with the exact permissions to read and write custom resources of the `Kind` that said service controller manages.

- 首先，我们使用来自有关 AWS 服务的规范真实来源的模型信息。我们将真实来源确定为来自 `aws/aws-sdk-go` 存储库的[模型文件](https://github.com/aws/aws-sdk-go/tree/master/models/apis) . AWS SDK 会定期更新所有 API 更改，因此这是一个准确的信息来源，并密切跟踪服务 API 的可用性。在这个阶段，我们生成包含代码的文件，这些代码公开了在那里找到的对象和接口的 Go 类型。
- 在为 AWS API 公开的顶级资源生成 Kubernetes API 类型定义后，我们需要生成接口实现，使这些顶级资源和类型定义能够被 Kubernetes 运行时包使用
- 接下来，我们生成自定义资源定义 (CRD) 配置文件，一个用于前面步骤中确定的每个顶级资源。然后，我们为目标服务生成 ACK 控制器的实现。除了 Go 中的这些控制器实现之外，此步骤还为下一步的“角色”的“部署”和“ClusterRoleBinding”输出一组 Kubernetes 清单。
- 最后，我们为运行相应 ACK 服务控制器的 Kubernetes `Deployment` 生成 Kubernetes`Role` 的 Kubernetes 清单。遵循最小权限原则，此“Role”需要具备读取和写入所述服务控制器管理的“Kind”的自定义资源的确切权限。

The above artifacts—Go code, container images, Kubernetes manifests for CRDs, roles, deployments, etc.—represent the business logic of how to manage AWS resources from within Kubernetes and are the responsibility of AWS service teams to create and maintain along with input from the community.

上述工件——Go 代码、容器镜像、CRD 的 Kubernetes 清单、角色、部署等——代表了如何从 Kubernetes 中管理 AWS 资源的业务逻辑，并且是 AWS 服务团队创建和维护输入的责任来自社区。

### 2\. Installation of custom resources & controllers

### 2. 安装自定义资源和控制器

To use ACK in a cluster you install the desired AWS service controller(s), considering that:

要在集群中使用 ACK，您需要安装所需的 AWS 服务控制器，考虑到：

- You set the respective Kubernetes Role-based Access Control ( [RBAC](https://rbac.dev/)) permissions for ACK custom resources. Note that each ACK service controller runs in its own pod and you can and should enforce existing IAM controls including permissions boundaries or service control policies to define who has access to which resources, transitively defining them via RBAC.
- You associate an AWS account ID to a Kubernetes namespace. Consequently that mean every ACK custom resource must be namespaced (no cluster-wide custom resources).

- 您为 ACK 自定义资源设置相应的 Kubernetes 基于角色的访问控制 ([RBAC](https://rbac.dev/)) 权限。请注意，每个 ACK 服务控制器都在其自己的 pod 中运行，您可以而且应该强制执行现有的 IAM 控制，包括权限边界或服务控制策略，以定义谁有权访问哪些资源，并通过 RBAC 传递地定义它们。
- 您将 AWS 账户 ID 关联到 Kubernetes 命名空间。因此，这意味着每个 ACK 自定义资源都必须命名空间（没有集群范围的自定义资源）。

As per the [AWS Shared Responsibility Model](https://aws.amazon.com/compliance/shared-responsibility-model/), in the context of the cluster administration, you are responsible for regularly upgrading the ACK service controllers as well as applying security patches as they are made available.

根据 [AWS 责任共担模型](https://aws.amazon.com/compliance/shared-responsibility-model/)，在集群管理的上下文中，您还负责定期升级 ACK 服务控制器就像应用可用的安全补丁一样。

### 3\. Starting AWS resources from Kubernetes

### 3. 从 Kubernetes 启动 AWS 资源

As an application developer, you create a [namespaced](https://kubernetes.io/docs/tasks/administer-cluster/namespaces-walkthrough/) custom resource in one of your ACK-enabled clusters. For example, let’s say you want ACK to create an Amazon Elastic Container Registry (ECR) repository, you’d define and subsequently apply something like:

作为应用程序开发人员，您在启用 ACK 的集群之一中创建一个 [命名空间](https://kubernetes.io/docs/tasks/administer-cluster/namespaces-walkthrough/) 自定义资源。例如，假设您希望 ACK 创建一个 Amazon Elastic Container Registry (ECR) 存储库，您将定义并随后应用如下内容：

```lang-yaml
apiVersion: "ecr.services.k8s.aws/v1alpha1"
kind: Repository
metadata:
    name: "my-ecr-repo"
spec:
    repositoryName: "encrypted-repo-managed-by-ack"
    encryptionConfiguration:
        encryptionType: AES256
    tags:
    - key: "is-encrypted"
      value: "true"

```

### 4\. Managing of AWS resources from Kubernetes 

### 4. 从 Kubernetes 管理 AWS 资源

ACK service controllers installed by cluster admins can create, update, or delete AWS resources, based on the intent found in the custom resource defined in the previous step, by developers. This means that in an ACK-enabled target cluster, the respective AWS resource (in our example case the ECR repo) will be created, with you having access to it, once you apply the custom resource.

由集群管理员安装的 ACK 服务控制器可以根据开发人员在上一步定义的自定义资源中找到的意图创建、更新或删除 AWS 资源。这意味着在启用 ACK 的目标集群中，将创建相应的 AWS 资源（在我们的示例中为 ECR 存储库），一旦您应用自定义资源，您就可以访问它。

Let’s focus a little bit more on the creation and management of AWS resources from Kubernetes, since this is how most users will interact with ACK. In our example, we will create an S3 bucket from our cluster.

让我们更多地关注从 Kubernetes 创建和管理 AWS 资源，因为这是大多数用户与 ACK 交互的方式。在我们的示例中，我们将从集群创建一个 S3 存储桶。

## Walkthrough: managing an S3 bucket with ACK

## 演练：使用 ACK 管理 S3 存储桶

In the following, we want to use ACK to manage an S3 bucket for us. Given that this is a developer preview, we’re using the [testing instructions as per contributor docs](https://aws-controllers-k8s.github.io/community/dev-docs/testing/). In this context, we’re use [kind](https://kind.sigs.k8s.io/) to do local end-to-end testing with Docker as its only dependency.

下面，我们要使用ACK来为我们管理一个S3 bucket。鉴于这是开发人员预览版，我们正在使用 [根据贡献者文档的测试说明](https://aws-controllers-k8s.github.io/community/dev-docs/testing/)。在这种情况下，我们使用[kind](https://kind.sigs.k8s.io/) 以 Docker 作为其唯一依赖项进行本地端到端测试。

Building the container image for the S3 service controller, creating the cluster and deploying all the resources with the `make kind-test -s SERVICE=s3` command will likely take some 45min the first time around (cold caches). Once that’s done, you can have a look at the ACK setup:

为 S3 服务控制器构建容器镜像，创建集群并使用“make kind-test -s SERVICE=s3”命令部署所有资源，第一次可能需要大约 45 分钟（冷缓存）。完成后，您可以查看 ACK 设置：

```lang-bash
$ kubectl -n ack-system get all
NAME                                     READY   STATUS    RESTARTS   AGE
pod/ack-s3-controller-86d9cf5cd7-z7l42   1/1     Running   0          10m

NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/ack-s3-controller   1/1     1            1           10m

NAME                                           DESIRED   CURRENT   READY   AGE
replicaset.apps/ack-s3-controller-86d9cf5cd7   1         1         1       10
```

Further, we’d expect the S3 CRD to be installed and available in the test cluster, and indeed:

此外，我们希望在测试集群中安装并使用 S3 CRD，并且确实：

```lang-bash
$ kubectl get crd
NAME                          CREATED AT
buckets.s3.services.k8s.aws   2020-08-17T06:15:22Z
```

Based on the CRD we would further expect to find an S3 bucket custom resource:

基于 CRD，我们将进一步期望找到一个 S3 存储桶自定义资源：

```lang-bash
$ kubectl get buckets
NAME AGE
ack-test-smoke-s3   3m8s
```

Let’s now have a cluster look at the S3 bucket customer resource (note: what is shown here is the automatically generated custom resource from the integration test, edited for readability):

现在让我们看一下 S3 存储桶客户资源的集群（注意：此处显示的是集成测试中自动生成的自定义资源，为便于阅读而进行了编辑）：

```lang-bash
$ kubectl get buckets ack-test-smoke-s3 -o yaml
apiVersion: s3.services.k8s.aws/v1alpha1
kind: Bucket
metadata:
name: ack-test-smoke-s3
namespace: default
spec:
name: ack-test-smoke-s3
```

Taking all together, the above setup looks as follows:

综上所述，上述设置如下所示：

![](https://d2908q01vomqb2.cloudfront.net/fe2ef495a1152561572949784c16bf23abb28057/2020/08/17/s3-example-1024x743.png)

OK, with this hands-on completed you should now have an idea how ACK works. Let us now turn our attention to how you can get engaged and what’s up next.

好的，完成这个动手操作后，您现在应该了解 ACK 的工作原理了。现在让我们将注意力转移到如何参与以及接下来会发生什么。

## Next steps

##  下一步

We are super excited that as of today, ACK is available as a [developer preview,](https://github.com/aws/aws-controllers-k8s/issues/22) supporting the following AWS services:

我们非常高兴，截至今天，ACK 可作为 [开发人员预览](https://github.com/aws/aws-controllers-k8s/issues/22) 支持以下 AWS 服务：

- Amazon API Gateway V2
- Amazon DynamoDB
- Amazon ECR
- Amazon S3
- Amazon SNS
- Amazon SQS

- 亚马逊 API 网关 V2
- 亚马逊 DynamoDB
- 亚马逊ECR
- 亚马逊 S3
- 亚马逊社交网络
- 亚马逊 SQS

You can get started with installing and using ACK with [our documentation](https://aws-controllers-k8s.github.io/community/). Note that developer preview means that the [end-user facing install](https://aws-controllers-k8s.github.io/community/user-docs/wip/) mechanisms are not yet in place.

您可以通过 [我们的文档](https://aws-controllers-k8s.github.io/community/) 开始安装和使用 ACK。请注意，开发人员预览意味着 [面向最终用户的安装](https://aws-controllers-k8s.github.io/community/user-docs/wip/) 机制尚未到位。

### **Future services**

### **未来的服务**

In time, we expect to support and enable as many AWS services as possible, and onboard them with full support to ACK. Specifically, in the coming months we plan to focus on:

随着时间的推移，我们希望支持和启用尽可能多的 AWS 服务，并为它们提供对 ACK 的全面支持。具体来说，在接下来的几个月里，我们计划重点关注：

- Amazon Relational Database Service (RDS), track via the[RDS label](https://github.com/aws/aws-controllers-k8s/labels/RDS).
- Amazon ElastiCache offers fully managed Redis and Memcached, track via the[Elasticache label](https://github.com/aws/aws-controllers-k8s/labels/Elasticache).

- 亚马逊关系数据库服务 (RDS)，通过[RDS 标签](https://github.com/aws/aws-controllers-k8s/labels/RDS) 进行跟踪。
- Amazon ElastiCache 提供完全托管的 Redis 和 Memcached，通过 [Elasticache 标签](https://github.com/aws/aws-controllers-k8s/labels/Elasticache) 进行跟踪。

In addition to RDS and ElastiCache, we are also considering support for Amazon Elastic Kubernetes Service (EKS) as well as Amazon Managed Streaming for Apache Kafka (MSK).

除了 RDS 和 ElastiCache，我们还在考虑支持 Amazon Elastic Kubernetes Service (EKS) 以及 Amazon Managed Streaming for Apache Kafka (MSK)。

### **Upcoming features**

### **即将推出的功能**

Important features we’re working on and which should be available as well within the next couple of weeks and months include: 

我们正在开发的重要功能以及应该在接下来的几周和几个月内可用的功能包括：

- Enabling [cross-account](https://github.com/aws-controllers-k8s/community/blob/main/docs/design/proposals/carm/cross-account-resource-management.md) resource management.
- Native application[secrets integration](https://github.com/aws-controllers-k8s/community/blob/main/docs/design/proposals/secrets/secrets.md).

- 启用[跨账户](https://github.com/aws-controllers-k8s/community/blob/main/docs/design/proposals/carm/cross-account-resource-management.md) 资源管理。
- 本机应用程序[秘密集成](https://github.com/aws-controllers-k8s/community/blob/main/docs/design/proposals/secrets/secrets.md)。

### **Help us build**

### **帮助我们构建**

ACK is still a new project, and we’re looking for input from you to help guide our development. Conceptually, we are interested in your feedback about:

ACK 仍然是一个新项目，我们正在寻找您的意见来帮助指导我们的开发。从概念上讲，我们对您对以下方面的反馈感兴趣：

- The expected behavior of[destructive operations](https://github.com/aws/aws-controllers-k8s/issues/82) in ACK.
- Whether or not ACK should be able to[adopt AWS resources](https://github.com/aws/aws-controllers-k8s/issues/41).

- ACK 中[破坏性操作](https://github.com/aws/aws-controllers-k8s/issues/82) 的预期行为。
- ACK 是否应该能够[采用 AWS 资源](https://github.com/aws/aws-controllers-k8s/issues/41)。

Starting today, during the AWS Container Day as well as during the entire KubeCon EU 2020 we will have a number of opportunities to discuss the current state of ACK and the next steps. Fire up your clusters, we can’t wait to see what you build!

从今天开始，在 AWS 容器日以及整个 KubeCon EU 2020 期间，我们将有很多机会讨论 ACK 的当前状态和后续步骤。启动您的集群，我们迫不及待地想看看您构建了什么！

TAGS:
[AWS services](https://aws.amazon.com/blogs/containers/tag/aws-services/), [custom controller](https://aws.amazon.com/blogs/containers/tag/custom-controller/), [custom resources](https://aws.amazon.com/blogs/containers/tag/custom-resources/), [Kubernetes](https://aws.amazon.com/blogs/containers/tag/kubernetes/), [operator](https://aws.amazon.com/blogs/containers/tag/operator/)

标签：
[AWS 服务](https://aws.amazon.com/blogs/containers/tag/aws-services/),[自定义控制器](https://aws.amazon.com/blogs/containers/tag/custom-控制器/), [自定义资源](https://aws.amazon.com/blogs/containers/tag/custom-resources/), [Kubernetes](https://aws.amazon.com/blogs/containers/tag/kubernetes/), [运营商](https://aws.amazon.com/blogs/containers/tag/operator/)

### Jay Pipes

Jay is a Principal Open Source Engineer at Amazon Web Services working on cloud-native technologies in the EKS team focused on open source contribution in the Kubernetes ecosystem.

Jay 是 Amazon Web Services 的首席开源工程师，在 EKS 团队中从事云原生技术研究，专注于 Kubernetes 生态系统中的开源贡献。

### Michael Hausenblas

Michael works in the AWS open source observability service team where he is a Solution Engineering Lead and owns the AWS Distro for OpenTelemetry (ADOT) from the product side.

Michael 在 AWS 开源可观察性服务团队工作，他是该团队的解决方案工程主管，并从产品方面拥有 AWS Distro for OpenTelemetry (ADOT)。

### Nathan Taber

Nathan is a Principal Product Manager for Amazon EKS. When he’s not writing and creating, he loves to sail, row, and roam the Pacific Northwest with his Goldendoodles, Emma & Leo.

Nathan 是 Amazon EKS 的首席产品经理。当他不写作和创作时，他喜欢和他的 Goldendoodles、Emma 和 Leo 一起在太平洋西北部航行、划船和漫游。

## Comments

##  注释

[View Comments](https://commenting.awsblogs.com/embed-1.0.html?disqus_shortname=aws-containers-blog&disqus_identifier=2842&disqus_title=Introducing+the+AWS+Controllers+for+Kubernetes+%28ACK%29&disqus_url=https://aws.amazon.com/blogs/containers/aws-controllers-for-kubernetes-ack/)

[查看评论](https://commenting.awsblogs.com/embed-1.0.html?disqus_shortname=aws-containers-blog&disqus_identifier=2842&disqus_title=Introducing+the+AWS+Controllers+for+Kubernetes+%28ACK%29&disqus_url=https://aws.amazon.com/blogs/containers/aws-controllers-for-kubernetes-ack/)

### Resources

###  资源

### Learn About AWS

### 了解 AWS

### Resources for AWS

### AWS 资源

https://aws.amazon.com/blogs/containers/aws-controllers-for-kubernetes-ack

