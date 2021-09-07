# Kubernetes Finalizers

# 终结者

Finalizers are namespaced keys that tell Kubernetes to wait until specific conditions are met before it fully deletes resources marked for deletion. Finalizers alert [controllers](https://kubernetes.io/docs/concepts/architecture/controller/) to clean up resources the deleted object owned.

终结器是命名空间键，告诉 Kubernetes 在完全删除标记为删除的资源之前等待特定条件满足。终结器提醒 [controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 清理已删除对象拥有的资源。

When you tell Kubernetes to delete an object that has finalizers specified for it, the Kubernetes API marks the object for deletion, putting it into a read-only state. The target object remains in a terminating state while the control plane, or other components, take the actions defined by the finalizers. After these actions are complete, the controller removes the relevant finalizers from the target object. When the `metadata.finalizers` field is empty, Kubernetes considers the deletion complete.

当您告诉 Kubernetes 删除为其指定了终结器的对象时，Kubernetes API 会将对象标记为删除，将其置于只读状态。当控制平面或其他组件执行终结器定义的操作时，目标对象保持终止状态。这些操作完成后，控制器从目标对象中删除相关的终结器。当 `metadata.finalizers` 字段为空时，Kubernetes 认为删除完成。

You can use finalizers to control [garbage collection](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/) of resources. For example, you can define a finalizer to clean up related resources or infrastructure before the controller deletes the target resource.

您可以使用终结器来控制资源的 [垃圾收集](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/)。例如，您可以定义一个终结器来在控制器删除目标资源之前清理相关资源或基础设施。

You can use finalizers to control [garbage collection](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/) of resources by alerting [controllers](https://kubernetes.io/docs/concepts/architecture/controller/) to perform specific cleanup tasks before deleting the target resource.

您可以使用终结器通过警告 [controllers](https://kubernetes.io/docs/) 来控制资源的 [垃圾收集](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/)概念/架构/控制器/)在删除目标资源之前执行特定的清理任务。

Finalizers don't usually specify the code to execute. Instead, they are typically lists of keys on a specific resource similar to annotations. Kubernetes specifies some finalizers automatically, but you can also specify your own.

终结器通常不指定要执行的代码。相反，它们通常是特定资源上的键列表，类似于注释。 Kubernetes 会自动指定一些终结器，但您也可以指定自己的终结器。

## How finalizers work

## 终结器如何工作

When you create a resource using a manifest file, you can specify finalizers in the `metadata.finalizers` field. When you attempt to delete the resource, the controller that manages it notices the values in the `finalizers` field and does the following:

当您使用清单文件创建资源时，您可以在 `metadata.finalizers` 字段中指定终结器。当您尝试删除资源时，管理它的控制器会注意到 `finalizers` 字段中的值并执行以下操作：

- Modifies the object to add a `metadata.deletionTimestamp` field with the time you started the deletion.
- Marks the object as read-only until its `metadata.finalizers` field is empty.

- 修改对象以在您开始删除时添加 `metadata.deletionTimestamp` 字段。
- 将对象标记为只读，直到其 `metadata.finalizers` 字段为空。

The controller then attempts to satisfy the requirements of the finalizers specified for that resource. Each time a finalizer condition is satisfied, the controller removes that key from the resource's `finalizers` field. When the field is empty, garbage collection continues. You can also use finalizers to prevent deletion of unmanaged resources.

然后控制器尝试满足为该资源指定的终结器的要求。每次满足终结器条件时，控制器都会从资源的“终结器”字段中删除该键。当该字段为空时，垃圾收集继续。您还可以使用终结器来防止删除非托管资源。

A common example of a finalizer is `kubernetes.io/pv-protection`, which prevents accidental deletion of `PersistentVolume` objects. When a `PersistentVolume` object is in use by a Pod, Kubernetes adds the `pv-protection` finalizer. If you try to delete the `PersistentVolume`, it enters a `Terminating` status, but the controller can't delete it because the finalizer exists. When the Pod stops using the `PersistentVolume`, Kubernetes clears the `pv-protection` finalizer, and the controller deletes the volume.

终结器的一个常见示例是 `kubernetes.io/pv-protection`，它可以防止意外删除 `PersistentVolume` 对象。当 Pod 正在使用“PersistentVolume”对象时，Kubernetes 会添加“pv-protection”终结器。如果您尝试删除 `PersistentVolume`，它会进入 `Terminating` 状态，但控制器无法删除它，因为终结器存在。当 Pod 停止使用 `PersistentVolume` 时，Kubernetes 会清除 `pv-protection` 终结器，并且控制器删除该卷。

## Owner references, labels, and finalizers

## 所有者引用、标签和终结器

Like [labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels), [owner references](https://kubernetes.io/concepts/overview/working-with-objects/owners-dependents/) describe the relationships between objects in Kubernetes, but are used for a different purpose. When a [controller](https://kubernetes.io/docs/concepts/architecture/controller/) manages objects like Pods, it uses labels to track changes to groups of related objects. For example, when a [Job](https://kubernetes.io/docs/concepts/workloads/controllers/job/) creates one or more Pods, the Job controller applies labels to those pods and tracks changes to any Pods in the cluster with the same label.

像[标签](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels)，[所有者引用](https://kubernetes.io/concepts/overview/working-with-objects) 描述了 Kubernetes 中对象之间的关系，但用于不同的目的。当 [controller](https://kubernetes.io/docs/concepts/architecture/controller/) 管理 Pod 等对象时，它会使用标签来跟踪相关对象组的更改。例如，当一个 [Job](https://kubernetes.io/docs/concepts/workloads/controllers/job/) 创建一个或多个 Pod 时，Job 控制器会为这些 Pod 应用标签并跟踪其中任何 Pod 的变化。具有相同标签的集群。

The Job controller also adds *owner references* to those Pods, pointing at the Job that created the Pods. If you delete the Job while these Pods are running, Kubernetes uses the owner references (not labels) to determine which Pods in the cluster need cleanup.

作业控制器还向这些 Pod 添加*所有者引用*，指向创建 Pod 的作业。如果您在这些 Pod 运行时删除作业，Kubernetes 将使用所有者引用（而不是标签）来确定集群中的哪些 Pod 需要清理。

Kubernetes also processes finalizers when it identifies owner references on a resource targeted for deletion. 

Kubernetes 在确定要删除的资源上的所有者引用时也会处理终结器。

In some situations, finalizers can block the deletion of dependent objects, which can cause the targeted owner object to remain in a read-only state for longer than expected without being fully deleted. In these situations, you should check finalizers and owner references on the target owner and dependent objects to troubleshoot the cause.

在某些情况下，终结器可以阻止依赖对象的删除，这可能导致目标所有者对象在未完全删除的情况下保持只读状态的时间比预期的长。在这些情况下，您应该检查目标所有者和从属对象上的终结器和所有者引用以排除原因。

> **Note:** In cases where objects are stuck in a deleting state, try to avoid manually removing finalizers to allow deletion to continue. Finalizers are usually added to resources for a reason, so forcefully removing them can lead to issues in your cluster.

> **注意：** 在对象卡在删除状态的情况下，尽量避免手动删除终结器以允许删除继续。终结器通常出于某种原因添加到资源中，因此强行删除它们可能会导致集群出现问题。

## What's next

##  下一步是什么

Read [Using Finalizers to Control Deletion](https://kubernetes.io/blog/2021/05/14/using-finalizers-to-control-deletion/) on the Kubernetes blo 

在 Kubernetes blo 上阅读 [Using Finalizers to Control Deletion](https://kubernetes.io/blog/2021/05/14/using-finalizers-to-control-deletion/)

