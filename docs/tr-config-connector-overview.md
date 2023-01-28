# Config Connector overview

# 配置连接器概述

https://cloud.google.com/config-connector/docs/overview

Config Connector is an [open source](https://github.com/GoogleCloudPlatform/k8s-config-connector) Kubernetes addon that allows you to manage Google Cloud resources through Kubernetes.

Config Connector 是一个[开源](https://github.com/GoogleCloudPlatform/k8s-config-connector) Kubernetes 插件，可让您通过 Kubernetes 管理 Google Cloud 资源。

Many cloud-native development teams work with a mix of configuration systems, APIs, and tools to manage their infrastructure. This mix is often difficult to understand, leading to reduced velocity and expensive mistakes. Config Connector provides a method to configure many [Google Cloud services and resources](https://cloud.google.com/config-connector/docs/reference/resources) using Kubernetes tooling and APIs.

许多云原生开发团队混合使用配置系统、API 和工具来管理他们的基础设施。这种混合通常难以理解，导致速度降低和代价高昂的错误。 Config Connector 提供了一种使用 Kubernetes 工具和 API 配置许多 [Google Cloud 服务和资源](https://cloud.google.com/config-connector/docs/reference/resources) 的方法。

With Config Connector, your environments can leverage how Kubernetes manages [Resources](https://cloud.google.com/config-connector/docs/concepts/resources#managing_resources_with_kubernetes_objects) including:

借助 Config Connector，您的环境可以利用 Kubernetes 管理[资源](https://cloud.google.com/config-connector/docs/concepts/resources#managing_resources_with_kubernetes_objects) 的方式，包括：

- RBAC for access control.
- Events for visibility.
- Single source of configuration and desired state management for reduced complexity.
- Eventual consistency for loosely coupling dependencies.

- 用于访问控制的 RBAC。
- 可见性事件。
- 用于降低复杂性的配置和所需状态管理的单一来源。
- 松耦合依赖的最终一致性。

You can manage your Google Cloud infrastructure the same way you manage your Kubernetes applications, reducing the complexity and cognitive load for developers.

您可以像管理 Kubernetes 应用程序一样管理 Google Cloud 基础架构，从而降低开发人员的复杂性和认知负担。

## How Config Connector works

## 配置连接器如何工作

Config Connector provides a collection of Kubernetes [Custom Resource Definitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)  (CRDs) and controllers. The Config Connector CRDs allow Kubernetes to create and manage Google Cloud resources when you configure and apply Objects to your cluster.

Config Connector 提供了 Kubernetes [自定义资源定义](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)(CRD) 和控制器的集合。 Config Connector CRD 允许 Kubernetes 在您配置对象并将其应用于集群时创建和管理 Google Cloud 资源。

For Config Connector CRDs to function correctly, Config Connector deploys Pods to your nodes that have elevated RBAC permissions, such as the ability to create, delete, get, and list CustomResourceDefinitions (CRDs). These permissions are required for Config Connector to create and reconcile Kubernetes resources.

为了使 Config Connector CRD 正常运行，Config Connector 将 Pod 部署到具有提升的 RBAC 权限的节点，例如创建、删除、获取和列出 CustomResourceDefinitions (CRD) 的能力。 Config Connector 需要这些权限才能创建和协调 Kubernetes 资源。

To get started, [install Config Connector](https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall) and [create your first resource](https://cloud.google.com/config-connector/docs/how-to/getting-started). Config Connector's [controllers](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#custom-controllers)  eventually reconcile your environment with your desired state.

要开始使用，请[安装 Config Connector](https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall) 并[创建您的第一个资源](https://cloud.google.com/config-connector/docs/how-to/getting-started)。 Config Connector 的 [controllers](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#custom-controllers) 最终将您的环境与您想要的状态相协调。

### Customizing Config Connector's behavior

### 自定义配置连接器的行为

Config Connector provides additional functionality beyond creating resources. For example, you can manage [existing Google Cloud resources](https://cloud.google.com/config-connector/docs/how-to/managing-deleting-resources#acquiring_an_existing_resource), and use [Kubernetes Secrets](https://cloud.google.com/config-connector/docs/how-to/secrets) to provide sensitive data, such as passwords, to your resources. For more information, see the list of [how-to guides](https://cloud.google.com/config-connector/docs/how-to).

Config Connector 提供了创建资源之外的附加功能。例如，您可以管理[现有的谷歌云资源](https://cloud.google.com/config-connector/docs/how-to/managing-deleting-resources#acquiring_an_existing_resource)，并使用[Kubernetes Secrets](https://cloud.google.com/config-connector/docs/how-to/secrets)为您的资源提供敏感数据，例如密码。有关详细信息，请参阅[操作指南](https://cloud.google.com/config-connector/docs/how-to) 列表。

In addition, you can learn more about how Config Connector uses Kubernetes constructs to manage [Resources](https://cloud.google.com/config-connector/docs/concepts/resources) and see the Google Cloud [resources](https://cloud.google.com/config-connector/docs/reference/resources) Config Connector can manage.

此外，您可以详细了解 Config Connector 如何使用 Kubernetes 构造来管理 [资源](https://cloud.google.com/config-connector/docs/concepts/resources) 并查看 Google Cloud [资源](https://cloud.google.com/config-connector/docs/reference/resources) Config Connector 可以管理。

## What's next

##  下一步是什么

- [Install Config Connector](https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall).
- [Get started](https://cloud.google.com/config-connector/docs/how-to/getting-started) by creating your first resource.
- Learn about best practices for common cloud applications by exploring [Cloud Foundation Toolkit's Config Connector solutions](https://github.com/GoogleCloudPlatform/cloud-foundation-toolkit/tree/master/config-connector/solutions).
- Explore Config Connector [source code](https://github.com/GoogleCloudPlatform/k8s-config-connector). Config Connector is fully open sourced on GitHub.

- [安装配置连接器](https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall)。
- [开始](https://cloud.google.com/config-connector/docs/how-to/getting-started) 创建您的第一个资源。
- 通过探索 [Cloud Foundation Toolkit 的 Config Connector 解决方案](https://github.com/GoogleCloudPlatform/cloud-foundation-toolkit/tree/master/config-connector/solutions) 了解常见云应用程序的最佳实践。
- 探索配置连接器 [源代码](https://github.com/GoogleCloudPlatform/k8s-config-connector)。 Config Connector 在 GitHub 上完全开源。

