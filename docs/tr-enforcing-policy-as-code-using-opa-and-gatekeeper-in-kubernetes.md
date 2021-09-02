## Enforcing Policy as Code using OPA and Gatekeeper in Kubernetes

## 在 Kubernetes 中使用 OPA 和 Gatekeeper 执行策略即代码

April 15, 2021

Enforcing organizational policies on a Kubernetes cluster allows you to be in control of the resources being deployed. For example, you can prevent deploying non-vetted pods to the production environment or disable usage of default passwords for databases. Doing so stops you from worrying about quarterly security reviews and a ton of issues in your backlog.

在 Kubernetes 集群上实施组织策略允许您控制正在部署的资源。例如，您可以阻止将未经审查的 pod 部署到生产环境或禁用数据库的默认密码。这样做可以让您不必担心季度安全审查和积压工作中的大量问题。

In this blog post, we will go through everything necessary to set up “OPA(Open Policy Agent)”/Gatekeeper as your Kubernetes admission webhook, which enables you to enforce policies on your Kubernetes cluster.

在这篇博文中，我们将介绍将“OPA（开放策略代理）”/Gatekeeper 设置为 Kubernetes 准入 webhook 所需的一切，它使您能够在 Kubernetes 集群上实施策略。

## Open Policy Agent (OPA)

## 开放策略代理（OPA）

Open Policy Agent is an open-source, general-purpose policy engine that enforces validation of objects during creation, updating, and deletion operations. OPA lets us enforce custom policies on Kubernetes objects without manually reconfiguring the Kubernetes API server. It will ensure that no Deployments, Jobs, Pods, etc are scheduled without being compliant with your Constraints and rules.

Open Policy Agent 是一个开源的通用策略引擎，它在创建、更新和删除操作期间强制执行对象验证。 OPA 允许我们在 Kubernetes 对象上实施自定义策略，而无需手动重新配置 Kubernetes API 服务器。它将确保不会在不符合您的约束和规则的情况下安排任何部署、作业、Pod 等。

OPA was designed to let us write policies over arbitrary JSON/YAML. You can essentially use OPA to enforce policies on any tool that takes JSON/YAML as input, such as Kubernetes, Terraform, CI/CD pipelines. For information on OPA and its use cases, please refer to the [official documentation](https://www.openpolicyagent.org/docs/latest/). We will focus on OPA’s Kubernetes [admission controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) use case with Gatekeeper.

OPA 旨在让我们通过任意 JSON/YAML 编写策略。您基本上可以使用 OPA 对任何以 JSON/YAML 作为输入的工具实施策略，例如 Kubernetes、Terraform、CI/CD 管道。有关 OPA 及其用例的信息，请参阅[官方文档](https://www.openpolicyagent.org/docs/latest/)。我们将重点关注 OPA 的 Kubernetes [准入控制器](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) 使用 Gatekeeper 的用例。

## Gatekeeper

## 看门人

Gatekeeper is a customizable admission webhook for Kubernetes that dynamically enforces policies executed by the OPA. Gatekeeper uses [CustomResourceDefinitions](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/) internally and allows us to define **ConstraintTemplates** and **Constraints** to enforce policies on Kubernetes resources such as Pods, Deployments, Jobs.

Gatekeeper 是一个可定制的 Kubernetes 准入 webhook，它动态地强制执行由 OPA 执行的策略。 Gatekeeper 在内部使用 [CustomResourceDefinitions](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/) 并允许我们定义 **ConstraintTemplates** 和 **Constraints**对 Kubernetes 资源（例如 Pod、部署、作业)实施策略。

OPA/Gatekeeper uses its own declarative language called Rego, a query language. You define rules in Rego which, if invalid or returned a false expression, will trigger a constraint violation and blocks the ongoing process of creating/updating/deleting the resource.

OPA/Gatekeeper 使用自己的声明性语言 Rego，这是一种查询语言。您在 Rego 中定义规则，如果这些规则无效或返回错误表达式，将触发约束违规并阻止正在进行的创建/更新/删除资源的过程。

### Prerequisites

### 先决条件

- kubectl CLI
- helm CLI
- A running Kubernetes cluster. If you don’t have a cluster running we would recommend using [kubespray](https://github.com/kubernetes-sigs/kubespray).

- 正在运行的 Kubernetes 集群。如果您没有运行集群，我们建议您使用 [kubespray](https://github.com/kubernetes-sigs/kubespray)。

To install gatekeeper, run the following helm commands.

要安装gatekeeper，请运行以下 helm 命令。

```
helm repo add gatekeeper https://open-policy-agent.github.io/gatekeeper/charts

helm install gatekeeper/gatekeeper --generate-name
```


> **_NOTE:_** To check whether it is installed, the simplest way is to see if Gatekeeper controller manager and Gatekeeper audit deployments and services are deployed in the gatekeeper-system namespace.

> **_NOTE:_** 要检查是否安装，最简单的方法是查看Gatekeeper 控制器管理器和Gatekeeper 审计部署和服务是否部署在gatekeeper-system 命名空间中。

Now, before we go into the demo, you should know about constraint templates and constraints. Every policy in the Gatekeeper consists of two manifests, Constraint template, and Constraint.

现在，在我们进入演示之前，您应该了解约束模板和约束。 Gatekeeper 中的每个策略都包含两个清单：约束模板和约束。

To define a Constraint, you need to create a Constraint Template that allows people to declare new Constraints

要定义约束，您需要创建一个约束模板，允许人们声明新的约束

## ConstraintTemplate

## 约束模板

A ConstraintTemplate consists of both the Rego logic that enforces the Constraint and the schema for the Constraint, which includes the schema of the CRD and the parameters that can be passed into a Constraint.

ConstraintTemplate 由强制执行约束的 Rego 逻辑和约束的架构组成，后者包括 CRD 的架构和可以传递到约束的参数。

## Constraint

## 约束

Constraint is an object that says on which resources are the policies applicable, and also what parameters are to be queried and checked to see if they are available in the resource manifest the user is trying to apply in your Kubernetes cluster. Simply put, it is a declaration that its author wants the system to meet a given set of requirements.

Constraint 是一个对象，它说明哪些资源适用于策略，以及哪些参数将被查询和检查以查看它们在用户尝试在您的 Kubernetes 集群中应用的资源清单中是否可用。简而言之，它是作者希望系统满足一组给定要求的声明。

## Working example

## 工作示例

Here is a ConstraintTemplate CRD that requires certain labels to be present on the objects that are being created/deployed in a cluster.

这是一个 ConstraintTemplate CRD，它要求在集群中创建/部署的对象上存在某些标签。

```
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
name: k8srequiredlabels
spec:
crd:
spec:
     names:
       kind: K8sRequiredLabels
     validation:
       # Schema for the `parameters` field
       openAPIV3Schema:
         properties:
           labels:
             type: array
             items: string
targets:
   - target: admission.k8s.gatekeeper.sh
     rego: |
       package k8srequiredlabels
       violation[{"msg": msg, "details": {"missing_labels": missing}}] {
         provided := {label |input.review.object.metadata.labels[label]}
         required := {label |label := input.parameters.labels[_]}
         missing := required - provided
         count(missing) > 0
         msg := sprintf("\n\nDENIED. \nReason: Our org policy mandates the following labels: \nYou must provide these labels: %v", [missing])
       }
```




Once ConstraintTemplate Object is deployed, you can now enforce your polices by creating Constraints. Here is a Constraint that requires labels “zone”,“stage”,“status”.

部署 ConstraintTemplate 对象后，您现在可以通过创建约束来实施您的策略。这是一个需要标签“区域”、“阶段”、“状态”的约束。

> **_NOTE:_** This policy does not consider the values of the labels, it only checks whether the resources that are being deployed have the required labels attached to it or not. If the provided labels does not match the required labels, the constraint denies the objects creation.

> **_注意：_** 此策略不考虑标签的值，它只检查正在部署的资源是否附加了所需的标签。如果提供的标签与所需标签不匹配，则约束拒绝对象创建。

```
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredLabels
metadata:
name: label-check
spec:
match:
kinds:
     - apiGroups: [""]
       kinds: ["Namespace", "Pod"]
excludedNamespaces:
   - kube-system
   - kube-public
   - kube-node-lease
   - gatekeeper-system
parameters:
labels:
     - "zone"
     - "stage"
     - "status"
```


You can copy and save these as `constrainttemplate.yml` and `constraints.yml` and run the following commands to apply the objects.

您可以将这些复制并保存为 `constrainttemplate.yml` 和 `constraints.yml`，然后运行以下命令来应用这些对象。

```
kubectl apply -f constrainttemplate.yml
kubectl apply -f constraints.yml
```


Now, once applied, This constraint checks if every pod or namespace that is being created has the specified labels attached to it. You can test it by creating a pod or a namespace with and without labels and see the difference for yourselves. Here is an output of a pod that is being deployed without any labels.

现在，一旦应用，此约束将检查正在创建的每个 pod 或命名空间是否附加了指定的标签。您可以通过创建一个带标签和不带标签的 pod 或命名空间来测试它，并亲自查看差异。这是正在部署的 pod 的输出，没有任何标签。

```
$ kubectl run nginx --image=nginx --generator=run-pod/v1
Error from server ([denied by label-check]

DENIED.
Reason: Our org policy mandates the following labels:
You must provide these labels: {"stage", "status", "zone"}): admission webhook "validation.gatekeeper.sh" denied the request: [denied by label-check]

DENIED.
Reason: Our org policy mandates the following labels:
You must provide these labels: {"stage", "status", "zone"}
```


The policy we set using the Constraint, denied the pod creation without required labels.

我们使用约束设置的策略拒绝在没有必需标签的情况下创建 pod。

Here is an output of a pod that is being deployed adhering to the policy of having required labels **_stage_**, **_status_**, **_zone_**.

这是正在部署的 pod 的输出，该 pod 遵循具有必需标签 **_stage_**、**_status_**、**_zone_** 的策略。

```
$ kubectl run nginx --image=nginx --generator=run-pod/v1 -l zone=us-east-1,stage=dev,status=ready
pod/nginx created
```


#### Note

####  笔记

You as an admin need to explicitly decide your organization specific policies and then turn your policies as code using Rego.

作为管理员，您需要明确决定您组织的特定政策，然后使用 Rego 将您的政策转化为代码。

## Community-provided libraries of policies

## 社区提供的政策库

You don’t have to write policies on your own at the beginning of your journey, OPA and Gatekeeper both have excellent community libraries. You can have a look, fork them, and use them in your organization from here, [OPA](https://github.com/open-policy-agent/library), and [Gatekeeper](https://github.com/open-policy-agent/gatekeeper-library) libraries.

您不必在旅程开始时自己编写策略，OPA 和 Gatekeeper 都有出色的社区库。您可以从这里查看、分叉它们并在您的组织中使用它们，[OPA](https://github.com/open-policy-agent/library) 和 [Gatekeeper](https://github.com)。com/open-policy-agent/gatekeeper-library) 库。

A glimpse on getting started with the Gatekeeper library, the library provided in the bonus assumes that you have prior knowledge or have already played with the Gatekeeper. If you have no prior knowledge, I would recommend going through this blog post and the [official documentation](https://open-policy-agent.github.io/gatekeeper/website/docs/) of Gatekeeper again, understand how Rego queries allow and deny objects.

一瞥 Gatekeeper 库的入门，奖金中提供的库假设您有先验知识或已经玩过 Gatekeeper。如果你没有任何先验知识，我建议你再次阅读这篇博文和Gatekeeper的[官方文档](https://open-policy-agent.github.io/gatekeeper/website/docs/)，了解Rego如何查询允许和拒绝对象。

## Summary

##  概括

In this blog post, we have done everything needed to enforce policies using Gatekeeper in your Kubernetes cluster and we went through how Gatekeeper reduces Policy Compliance burden. So you can rest assured that common configuration errors or security issues will be automatically prevented. We are sure that this tutorial will help with your journey of running a Kubernetes cluster complying with your organizational policies.

在这篇博文中，我们已经完成了在您的 Kubernetes 集群中使用 Gatekeeper 执行策略所需的一切，我们还介绍了 Gatekeeper 如何减轻策略合规性负担。因此您可以放心，常见的配置错误或安全问题将被自动阻止。我们确信本教程将帮助您运行符合组织策略的 Kubernetes 集群。

## Read more of our **engineering blog posts**

## 阅读更多我们的**工程博文**

This blog post is part of our engineering blog post series. Experience and expertise, straight from our engineering team. Always with a focus on technical, hands-on HOWTO content with copy-pasteable code or CLI commands.

这篇博文是我们工程博文系列的一部分。经验和专业知识，直接来自我们的工程团队。始终关注具有可复制粘贴代码或 CLI 命令的技术性、实践性 HOWTO 内容。

Would you like to read more content like this? Click the button below and see the other blog posts in this series!

你想阅读更多这样的内容吗？单击下面的按钮，查看本系列中的其他博客文章！

[Read more engineering blog posts](https://elastisys.com/category/tech-post/engineering/?utm_source=website&utm_medium=cta&utm_campaign=engineering) 

[阅读更多工程博文](https://elastisys.com/category/tech-post/engineering/?utm_source=website&utm_medium=cta&utm_campaign=engineering)

Tags: [Blog](https://elastisys.com/tag/blog/),[devops](https://elastisys.com/tag/devops/), [devsecops](https://elastisys.com/tag/devsecops/), [gatekeeper](https://elastisys.com/tag/gatekeeper/), [open policy agent](https://elastisys.com/tag/open-policy-agent/), [policy as code](https://elastisys.com/tag/policy-as-code/) 

标签：[博客](https://elastisys.com/tag/blog/)、[devops](https://elastisys.com/tag/devops/)、[devsecops](https://elastisys.com/tag/devsecops/), [gatekeeper](https://elastisys.com/tag/gatekeeper/),[开放策略代理](https://elastisys.com/tag/open-policy-agent/), [policy作为代码](https://elastisys.com/tag/policy-as-code/)


