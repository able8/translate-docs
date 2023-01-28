# Config Connector overview

https://cloud.google.com/config-connector/docs/overview

Config Connector is an [open source](https://github.com/GoogleCloudPlatform/k8s-config-connector) Kubernetes addon that allows you to manage Google Cloud resources through Kubernetes.

Many cloud-native development teams work with a mix of configuration systems, APIs, and tools to manage their infrastructure. This mix is often difficult to understand, leading to reduced velocity and expensive mistakes. Config Connector provides a method to configure many [Google Cloud services and resources](https://cloud.google.com/config-connector/docs/reference/resources) using Kubernetes tooling and APIs.

With Config Connector, your environments can leverage how Kubernetes manages [Resources](https://cloud.google.com/config-connector/docs/concepts/resources#managing_resources_with_kubernetes_objects) including:

- RBAC for access control.
- Events for visibility.
- Single source of configuration and desired state management for reduced complexity.
- Eventual consistency for loosely coupling dependencies.

You can manage your Google Cloud infrastructure the same way you manage your Kubernetes applications, reducing the complexity and cognitive load for developers.

## How Config Connector works

Config Connector provides a collection of Kubernetes [Custom Resource Definitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)  (CRDs) and controllers. The Config Connector CRDs allow Kubernetes to create and manage Google Cloud resources when you configure and apply Objects to your cluster.

For Config Connector CRDs to function correctly, Config Connector deploys Pods to your nodes that have elevated RBAC permissions, such as the ability to create, delete, get, and list CustomResourceDefinitions (CRDs). These permissions are required for Config Connector to create and reconcile Kubernetes resources.

To get started, [install Config Connector](https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall) and [create your first resource](https://cloud.google.com/config-connector/docs/how-to/getting-started). Config Connector's [controllers](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#custom-controllers)  eventually reconcile your environment with your desired state.

### Customizing Config Connector's behavior

Config Connector provides additional functionality beyond creating resources. For example, you can manage [existing Google Cloud resources](https://cloud.google.com/config-connector/docs/how-to/managing-deleting-resources#acquiring_an_existing_resource), and use [Kubernetes Secrets](https://cloud.google.com/config-connector/docs/how-to/secrets) to provide sensitive data, such as passwords, to your resources. For more information, see the list of [how-to guides](https://cloud.google.com/config-connector/docs/how-to).

In addition, you can learn more about how Config Connector uses Kubernetes constructs to manage [Resources](https://cloud.google.com/config-connector/docs/concepts/resources) and see the Google Cloud [resources](https://cloud.google.com/config-connector/docs/reference/resources) Config Connector can manage.

## What's next

- [Install Config Connector](https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall).
- [Get started](https://cloud.google.com/config-connector/docs/how-to/getting-started) by creating your first resource.
- Learn about best practices for common cloud applications by exploring [Cloud Foundation Toolkit's Config Connector solutions](https://github.com/GoogleCloudPlatform/cloud-foundation-toolkit/tree/master/config-connector/solutions).
- Explore Config Connector [source code](https://github.com/GoogleCloudPlatform/k8s-config-connector). Config Connector is fully open sourced on GitHub.



​                    











​                         
