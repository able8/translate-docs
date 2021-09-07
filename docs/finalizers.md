# Finalizers

Finalizers are namespaced keys that tell Kubernetes to wait until specific conditions are met before it fully deletes resources marked for deletion. Finalizers alert [controllers](https://kubernetes.io/docs/concepts/architecture/controller/) to clean up resources the deleted object owned.

When you tell Kubernetes to delete an object that has finalizers specified for it, the Kubernetes API marks the object for deletion, putting it into a read-only state. The target object remains in a terminating state while the control plane, or other components, take the actions defined by the finalizers. After these actions are complete, the controller removes the relevant finalizers from the target object. When the `metadata.finalizers` field is empty, Kubernetes considers the deletion complete.

You can use finalizers to control [garbage collection](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/) of resources. For example, you can define a finalizer to clean up related resources or infrastructure before the controller deletes the target resource.

You can use finalizers to control [garbage collection](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/) of resources by alerting [controllers](https://kubernetes.io/docs/concepts/architecture/controller/) to perform specific cleanup tasks before deleting the target resource.

Finalizers don't usually specify the code to execute. Instead, they are typically lists of keys on a specific resource similar to annotations. Kubernetes specifies some finalizers automatically, but you can also specify your own.

## How finalizers work

When you create a resource using a manifest file, you can specify finalizers in the `metadata.finalizers` field. When you attempt to delete the resource, the controller that manages it notices the values in the `finalizers` field and does the following:

- Modifies the object to add a `metadata.deletionTimestamp` field with the time you started the deletion.
- Marks the object as read-only until its `metadata.finalizers` field is empty.

The controller then attempts to satisfy the requirements of the finalizers specified for that resource. Each time a finalizer condition is satisfied, the controller removes that key from the resource's `finalizers` field. When the field is empty, garbage collection continues. You can also use finalizers to prevent deletion of unmanaged resources.

A common example of a finalizer is `kubernetes.io/pv-protection`, which prevents accidental deletion of `PersistentVolume` objects. When a `PersistentVolume` object is in use by a Pod, Kubernetes adds the `pv-protection` finalizer. If you try to delete the `PersistentVolume`, it enters a `Terminating` status, but the controller can't delete it because the finalizer exists. When the Pod stops using the `PersistentVolume`, Kubernetes clears the `pv-protection` finalizer, and the controller deletes the volume.

## Owner references, labels, and finalizers

Like [labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels), [owner references](https://kubernetes.io/concepts/overview/working-with-objects/owners-dependents/) describe the relationships between objects in Kubernetes, but are used for a different purpose. When a [controller](https://kubernetes.io/docs/concepts/architecture/controller/) manages objects like Pods, it uses labels to track changes to groups of related objects. For example, when a [Job](https://kubernetes.io/docs/concepts/workloads/controllers/job/) creates one or more Pods, the Job controller applies labels to those pods and tracks changes to any Pods in the cluster with the same label.

The Job controller also adds *owner references* to those Pods, pointing at the Job that created the Pods. If you delete the Job while these Pods are running, Kubernetes uses the owner references (not labels) to determine which Pods in the cluster need cleanup.

Kubernetes also processes finalizers when it identifies owner references on a resource targeted for deletion.

In some situations, finalizers can block the deletion of dependent objects, which can cause the targeted owner object to remain in a read-only state for longer than expected without being fully deleted. In these situations, you should check finalizers and owner references on the target owner and dependent objects to troubleshoot the cause.

> **Note:** In cases where objects are stuck in a deleting state, try to avoid manually removing finalizers to allow deletion to continue. Finalizers are usually added to resources for a reason, so forcefully removing them can lead to issues in your cluster.

## What's next

Read [Using Finalizers to Control Deletion](https://kubernetes.io/blog/2021/05/14/using-finalizers-to-control-deletion/) on the Kubernetes blo
