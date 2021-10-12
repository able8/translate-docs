# Auto-labeling Kubernetes resources with Kyverno

# 使用 Kyverno 自动标记 Kubernetes 资源

[CNCF Member Blog Post](http://www.cncf.io/lf-author-category/member/ "See more content from Member")

 [CNCF会员博文](http://www.cncf.io/lf-author-category/member/“查看更多会员内容”)

Posted on
December 30, 2020 By Anubhav Sharma

发表于
2020 年 12 月 30 日 作者 Anubhav Sharma

_Guest post originally published on [Nirmata’s blog](https://nirmata.com/2020/10/30/auto-labeling-kubernetes-resources-with-kyverno/) by Anubhav Sharma, VP, Business Development & Customer Success at Nirmata_

_客帖最初发表在 [Nirmata 的博客](https://nirmata.com/2020/10/30/auto-labeling-kubernetes-resources-with-kyverno/) 上，作者是 Nirmata 业务发展和客户成功副总裁 Anubhav Sharma_

## Introduction

##  介绍

As Kubernetes has become the foundational building block for enterprises to go cloud-native, the last couple of years have seen many solutions that have simplified the cluster creation process. But the Day-2 operations around Kubernetes still remains a complex endeavor, slowing down adoption and increasing the operational costs. Kubernetes’ complexity and skills gaps still remain the biggest factors that are in the way of Enterprise’s adoption of Kubernetes.

由于 Kubernetes 已成为企业走向云原生的基础构建块，过去几年出现了许多简化集群创建过程的解决方案。但是围绕 Kubernetes 的第 2 天操作仍然是一项复杂的工作，它减缓了采用速度并增加了运营成本。 Kubernetes 的复杂性和技能差距仍然是阻碍企业采用 Kubernetes 的最大因素。

Many of the Day-2 operations use cases include requirements for the central platform team to deliver secure and compliant environments to developers as efficiently as possible with necessary services and best practices preconfigured. Some examples of such use cases include configuring environments with Kubernetes best practices like resource quotas, network policy and pod security. This requires tools that can assess the environments as they are created and then configure them in compliance with the standard defined by the central platform team.

许多第 2 天运营用例包括要求中央平台团队尽可能高效地向开发人员提供安全且合规的环境，并预先配置必要的服务和最佳实践。此类用例的一些示例包括使用 Kubernetes 最佳实践（如资源配额、网络策略和 pod 安全性）配置环境。这需要能够在创建环境时评估环境的工具，然后根据中央平台团队定义的标准对其进行配置。

## **Kyverno: A Flexible Ops Tool for K8s**

## **Kyverno：适用于 K8s 的灵活运维工具**

Kubernetes provides powerful constructs like admission control webhooks that can be leveraged for the purposes of validating and mutating resources. Nirmata’s [Kyverno](https://kyverno.io/) was designed specifically to address these types of use cases using the declarative paradigm. Kyverno is an open-source policy engine that was designed for Kubernetes, It provides users with familiar constructs to write custom rules and easily implement to validate, mutate, and generate new resources as needed.

Kubernetes 提供了强大的结构，如准入控制 webhook，可用于验证和改变资源的目的。 Nirmata 的 [Kyverno](https://kyverno.io/) 专门设计用于使用声明式范例解决这些类型的用例。 Kyverno 是一个专为 Kubernetes 设计的开源策略引擎，它为用户提供了熟悉的构造来编写自定义规则并轻松实现以根据需要验证、变异和生成新资源。

Managing Kubernetes at scales requires the following best practices and applying standardization across configurations. One such pattern is to use Kubernetes labels. In Kubernetes, every resource can have one or more labels and Kubernetes makes it easy to find and manage the resources using labels.

大规模管理 Kubernetes 需要以下最佳实践并跨配置应用标准化。其中一种模式是使用 Kubernetes 标签。在 Kubernetes 中，每个资源都可以有一个或多个标签，Kubernetes 可以使用标签轻松查找和管理资源。

A very common use case for Day-2 operations is managing labels across namespaces and pods so that use cases like certificate updates, self-service logging/monitoring, backups etc. can be easily implemented by other Kubernetes controllers and operators.

Day-2 操作的一个非常常见的用例是跨命名空间和 pod 管理标签，以便其他 Kubernetes 控制器和操作员可以轻松实现证书更新、自助日志记录/监控、备份等用例。

## **Auto-Labeling Namespaces**

## **自动标记命名空间**

Below is an example of how to implement namespace labeling upon creation in a Kubernetes cluster using Kyverno.

下面是一个示例，说明如何使用 Kyverno 在 Kubernetes 集群中创建时实现命名空间标记。

Install Kyverno in your cluster:

在您的集群中安装 Kyverno：

```
kubectl create -f https://github.com/kyverno/kyverno/raw/master/definitions/install.yaml
```

Detailed installation instructions are available [here](https://kyverno.io/docs/installation/).

[此处](https://kyverno.io/docs/installation/) 提供了详细的安装说明。

Here is a sample Kyverno policy that adds labels to namespaces  –

这是一个向命名空间添加标签的 Kyverno 策略示例 –

```
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
name: add-labels
spec:
background: false
rules:
  - name: add-ns-label
    match:
      resources:
        kinds:
        - Namespace
    exclude:
      clusterroles: ["cluster-admin"]
    mutate:
      patchStrategicMerge:
        metadata:
          labels:
            kyverno/user: "{{ request.userInfo.username }}"
            +(kyverno/network): "default"
```

The policy inserts a label \`kyverno/user\`with the value of the user making the API request to create the namespace. The policy also inserts a label \`kyverno/network\`, but only if one is not already specified by the user. This simple policy demonstrates some powerful features in Kyverno like [variable substitution](https://kyverno.io/docs/writing-policies/writing-policies-variables/) and [conditional anchors](https://kyverno.io/docs/writing-policies/writing-policies-validate/).

该策略插入一个标签 \`kyverno/user\`，其中包含发出 API 请求以创建命名空间的用户的值。该策略还会插入一个标签 \`kyverno/network\`，但前提是用户尚未指定标签。这个简单的策略展示了 Kyverno 中的一些强大功能，例如 [变量替换](https://kyverno.io/docs/writing-policies/writing-policies-variables/) 和 [conditional anchors](https://kyverno.io/docs/writing-policies/writing-policies-validate/)。

Once the policy is configured in your cluster, create a new namespace and verify the labels have been added to the namespace automatically.

在集群中配置策略后，创建一个新的命名空间并验证标签是否已自动添加到命名空间。

Create a new namespace:

创建一个新的命名空间：

```
kubectl create ns test

```

View the namespace:

查看命名空间：

```
kubectl get ns test -o yaml
```

This should show a namespace similar to:

这应该显示类似于以下内容的命名空间：

```
apiVersion: v1
kind: Namespace
metadata:
labels:
    kyverno/network: default
    kyverno/user: docker-for-desktop
```

Now, what if you want to make sure that users cannot update a specific label? 

现在，如果您想确保用户无法更新特定标签怎么办？

Kyverno makes that easy to do as well! Here is a policy that prevents the update of the \`kyverno/network\` label:

Kyverno 也让这一切变得简单！这是一个阻止更新 \`kyverno/network\` 标签的策略：

```
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
name: protect-label
spec:
validationFailureAction: enforce
background: false
rules:
  - name: block-updates-for-label
    match:
      resources:
        kinds:
           - Namespace
validate:
    message: Updating label `kyverno/network` is not allowed
    deny:
    - key: "{{ request.operation }}"
      operator: "EQUALS"
      value: UPDATE
```

## **Summary**

##  **概括**

Managing Kubernetes configurations can be complex, and policy engines provide standardization, automated validation, and the ability to mutate and generate configurations.

管理 Kubernetes 配置可能很复杂，策略引擎提供标准化、自动化验证以及变异和生成配置的能力。

Kyverno is an open-source policy engine designed for Kubernetes. It has a minimal learning curve and provides tremendous flexibility for Kubernetes administrators to solve Day-2 operations challenges using Kubernetes’ powerful declarative management capabilities and native tools.

Kyverno 是一个为 Kubernetes 设计的开源策略引擎。它具有最小的学习曲线，并为 Kubernetes 管理员使用 Kubernetes 强大的声明式管理功能和本机工具解决第 2 天运营挑战提供了极大的灵活性。

Learn what else Kyverno can do at [https://kyverno.io](https://kyverno.io/).

在 [https://kyverno.io](https://kyverno.io/) 上了解 Kyverno 还能做什么。

Share this post 

分享这个帖子

