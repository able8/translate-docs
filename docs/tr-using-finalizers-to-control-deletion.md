# Using Finalizers to Control Deletion

# 使用终结器控制删除

Friday, May 14, 2021

**Authors:** Aaron Alpar (Kasten)

Deleting objects in Kubernetes can be challenging. You may think you’ve deleted something, only to find it still persists. While issuing a `kubectl delete` command and hoping for the best might work for day-to-day operations, understanding how Kubernetes `delete` commands operate will help you understand why some objects linger after deletion.

在 Kubernetes 中删除对象可能具有挑战性。你可能认为你已经删除了一些东西，却发现它仍然存在。虽然发出 `kubectl delete` 命令并希望最好的结果可能适用于日常操作，了解 Kubernetes `delete` 命令的操作方式将帮助您理解为什么某些对象在删除后仍然存在。

In this post, I’ll look at:

在这篇文章中，我将看看：

- What properties of a resource govern deletion
- How finalizers and owner references impact object deletion
- How the propagation policy can be used to change the order of deletions
- How deletion works, with examples



- 资源的哪些属性控制删除
- 终结器和所有者引用如何影响对象删除
- 如何使用传播策略来改变删除的顺序
- 删除的工作原理，举例说明

For simplicity, all examples will use ConfigMaps and basic shell commands to demonstrate the process. We’ll explore how the commands work and discuss repercussions and results from using them in practice.

为简单起见，所有示例都将使用 ConfigMaps 和基本的 shell 命令来演示该过程。我们将探索这些命令的工作原理，并讨论在实践中使用它们的影响和结果。

## The basic `delete`

## 基本的`delete`

Kubernetes has several different commands you can use that allow you to create, read, update, and delete objects. For the purpose of this blog post, we’ll focus on four `kubectl` commands: `create`, `get`, `patch`, and `delete`.

Kubernetes 有几个不同的命令，您可以使用它们来创建、读取、更新和删除对象。出于这篇博文的目的，我们将重点介绍四个 `kubectl` 命令：`create`、`get`、`patch` 和 `delete`。

Here are examples of the basic `kubectl delete` command:

以下是基本的“kubectl delete”命令示例：

```
kubectl create configmap mymap
configmap/mymap created
```


```
kubectl get configmap/mymap
NAME    DATA   AGE
mymap   0      12s
```


```
kubectl delete configmap/mymap
configmap "mymap" deleted
```


```
kubectl get configmap/mymap
Error from server (NotFound): configmaps "mymap" not found
```


Shell commands preceded by `$` are followed by their output. You can see that we begin with a `kubectl create configmap mymap`, which will create the empty configmap `mymap`. Next, we need to `get` the configmap to prove it exists. We can then delete that configmap. Attempting to `get` it again produces an HTTP 404 error, which means the configmap is not found.

以`$` 开头的Shell 命令后跟其输出。你可以看到我们从一个 `kubectl create configmap mymap` 开始，它将创建空的 configmap `mymap`。接下来，我们需要`get` configmap 来证明它存在。然后我们可以删除该配置映射。再次尝试“get”它会产生 HTTP 404 错误，这意味着找不到 configmap。

The state diagram for the basic `delete` command is very simple:

基本 `delete` 命令的状态图非常简单：

![State diagram for delete](https://d33wubrfki0l68.cloudfront.net/884a50ad3b41914a7a94f4c4c696753c399e450a/264d8/images/blog/2021-05-14-using-finalizers-to-control-deletion/state-diagram-delete.png)

State diagram for delete

删除状态图

Although this operation is straightforward, other factors may interfere with the deletion, including finalizers and owner references.

尽管此操作很简单，但其他因素可能会干扰删除，包括终结器和所有者引用。

## Understanding Finalizers

## 理解终结器

When it comes to understanding resource deletion in Kubernetes, knowledge of how finalizers work is helpful and can help you understand why some objects don’t get deleted.

在了解 Kubernetes 中的资源删除时，了解终结器的工作原理很有帮助，可以帮助您理解为什么某些对象没有被删除。

Finalizers are keys on resources that signal pre-delete operations. They control the garbage collection on resources, and are designed to alert controllers what cleanup operations to perform prior to removing a resource. However, they don’t necessarily name code that should be executed; finalizers on resources are basically just lists of keys much like annotations. Like annotations, they can be manipulated.

终结器是发出预删除操作信号的资源上的键。它们控制资源的垃圾收集，旨在提醒控制器在删除资源之前执行哪些清理操作。但是，它们不一定指定应该执行的代码；资源的终结器基本上只是键列表，就像注释一样。像注释一样，它们可以被操纵。

Some common finalizers you’ve likely encountered are:

您可能遇到的一些常见终结器是：

- `kubernetes.io/pv-protection`
- `kubernetes.io/pvc-protection`

  


The finalizers above are used on volumes to prevent accidental deletion. Similarly, some finalizers can be used to prevent deletion of any resource but are not managed by any controller.

上述终结器用于卷以防止意外删除。同样，一些终结器可用于防止删除任何资源，但不受任何控制器管理。

Below with a custom configmap, which has no properties but contains a finalizer:

下面是一个自定义配置映射，它没有属性但包含一个终结器：

```
cat <<EOF |kubectl create -f -
apiVersion: v1
kind: ConfigMap
metadata:
name: mymap
finalizers:
  - kubernetes
EOF

```


The configmap resource controller doesn't understand what to do with the `kubernetes` finalizer key. I term these “dead” finalizers for configmaps as it is normally used on namespaces. Here’s what happen upon attempting to delete the configmap:

configmap 资源控制器不知道如何处理 `kubernetes` 终结器键。我将这些“死”终结器称为 configmaps，因为它通常用于命名空间。以下是尝试删除 configmap 时发生的情况：

```
kubectl delete configmap/mymap &
configmap "mymap" deleted
jobs
[1]+  Running kubectl delete configmap/mymap

```


Kubernetes will report back that the object has been deleted, however, it hasn’t been deleted in a traditional sense. Rather, it’s in the process of deletion. When we attempt to `get` that object again, we discover the object has been modified to include the deletion timestamp.

Kubernetes 会报告该对象已被删除，但它不是传统意义上的删除。相反，它正在删除过程中。当我们再次尝试“获取”该对象时，我们发现该对象已被修改为包含删除时间戳。

```
kubectl get configmap/mymap -o yaml
apiVersion: v1
kind: ConfigMap
metadata:
creationTimestamp: "2020-10-22T21:30:18Z"
deletionGracePeriodSeconds: 0
deletionTimestamp: "2020-10-22T21:30:34Z"
finalizers:
  - kubernetes
name: mymap
namespace: default
resourceVersion: "311456"
selfLink: /api/v1/namespaces/default/configmaps/mymap
uid: 93a37fed-23e3-45e8-b6ee-b2521db81638

```




In short, what’s happened is that the object was updated, not deleted. That’s because Kubernetes saw that the object contained finalizers and put it into a read-only state. The deletion timestamp signals that the object can only be read, with the exception of removing the finalizer key updates. In other words, the deletion will not be complete until we edit the object and remove the finalizer.

简而言之，发生的事情是对象被更新了，而不是被删除了。那是因为 Kubernetes 看到该对象包含终结器并将其置于只读状态。删除时间戳表示该对象只能被读取，但删除终结器键更新除外。换句话说，在我们编辑对象并移除终结器之前，删除不会完成。

Here's a demonstration of using the `patch` command to remove finalizers. If we want to delete an object, we can simply patch it on the command line to remove the finalizers. In this way, the deletion that was running in the background will complete and the object will be deleted. When we attempt to `get` that configmap, it will be gone.

这是使用 `patch` 命令删除终结器的演示。如果我们想删除一个对象，我们可以简单地在命令行上修补它以删除终结器。这样，在后台运行的删除将完成，对象将被删除。当我们尝试`get`那个configmap时，它就会消失。

```
kubectl patch configmap/mymap \
    --type json \
    --patch='[ { "op": "remove", "path": "/metadata/finalizers" } ]'
configmap/mymap patched
[1]+  Done  kubectl delete configmap/mymap

kubectl get configmap/mymap -o yaml
Error from server (NotFound): configmaps "mymap" not found
```


Here's a state diagram for finalization:

这是最终确定的状态图：

![State diagram for finalize](https://d33wubrfki0l68.cloudfront.net/2921aff96caba07229c862903fea89cbab9ad5a6/8e8fd/images/blog/2021-05-14-using-finalizers-to-control-deletion/state-diagram-finalize.png)

State diagram for finalize

完成状态图

So, if you attempt to delete an object that has a finalizer on it, it will remain in finalization until the controller has removed the finalizer keys or the finalizers are removed using Kubectl. Once that finalizer list is empty, the object can actually be reclaimed by Kubernetes and put into a queue to be deleted from the registry.

因此，如果您尝试删除具有终结器的对象，它将保持终结状态，直到控制器删除终结器键或使用 Kubectl 删除终结器。一旦终结器列表为空，该对象实际上可以被 Kubernetes 回收并放入一个队列中以从注册表中删除。

## Owner References

## 所有者参考

Owner references describe how groups of objects are related. They are properties on resources that specify the relationship to one another, so entire trees of resources can be deleted.

所有者引用描述了对象组之间的关系。它们是资源的属性，用于指定彼此之间的关系，因此可以删除整个资源树。

Finalizer rules are processed when there are owner references. An owner reference consists of a name and a UID. Owner references link resources within the same namespace, and it also needs a UID for that reference to work. Pods typically have owner references to the owning replica set. So, when deployments or stateful sets are deleted, then the child replica sets and pods are deleted in the process.

当有所有者引用时处理终结器规则。所有者引用由名称和 UID 组成。所有者引用同一命名空间内的链接资源，它还需要一个 UID 才能使该引用起作用。 Pod 通常具有对拥有副本集的所有者引用。因此，当部署或有状态集被删除时，子副本集和 pod 也会在此过程中被删除。

Here are some examples of owner references and how they work. In the first example, we create a parent object first, then the child. The result is a very simple configmap that contains an owner reference to its parent:

以下是所有者引用及其工作原理的一些示例。在第一个例子中，我们首先创建一个父对象，然后是子对象。结果是一个非常简单的配置映射，其中包含对其父级的所有者引用：

```
cat <<EOF |kubectl create -f -
apiVersion: v1
kind: ConfigMap
metadata:
name: mymap-parent
EOF
CM_UID=$(kubectl get configmap mymap-parent -o jsonpath="{.metadata.uid}")

cat <<EOF |kubectl create -f -
apiVersion: v1
kind: ConfigMap
metadata:
name: mymap-child
ownerReferences:
  - apiVersion: v1
    kind: ConfigMap
    name: mymap-parent
    uid: $CM_UID
EOF

```


Deleting the child object when an owner reference is involved does not delete the parent:

在涉及所有者引用时删除子对象不会删除父对象：

```
kubectl get configmap
NAME           DATA   AGE
mymap-child    0      12m4s
mymap-parent   0      12m4s

kubectl delete configmap/mymap-child
configmap "mymap-child" deleted

kubectl get configmap
NAME           DATA   AGE
mymap-parent   0      12m10s

```


In this example, we re-created the parent-child configmaps from above. Now, when deleting from the parent (instead of the child) with an owner reference from the child to the parent, when we `get` the configmaps, none are in the namespace:

在这个例子中，我们从上面重新创建了父子配置映射。现在，当使用从子到父的所有者引用从父（而不是子）中删除时，当我们“获取”配置映射时，命名空间中没有：

```
kubectl get configmap
NAME           DATA   AGE
mymap-child    0      10m2s
mymap-parent   0      10m2s

kubectl delete configmap/mymap-parent
configmap "mymap-parent" deleted

kubectl get configmap
No resources found in default namespace.

```


To sum things up, when there's an override owner reference from a child to a parent, deleting the parent deletes the children automatically. This is called `cascade`. The default for cascade is `true`, however, you can use the --cascade=false option for `kubectl delete` to delete an object and orphan its children.

总而言之，当存在从子级到父级的覆盖所有者引用时，删除父级会自动删除子级。这称为“级联”。级联的默认值为“true”，但是，您可以使用“kubectl delete”的 --cascade=false 选项来删除对象并孤立其子对象。

In the following example, there is a parent and a child. Notice the owner references are still included. If I delete the parent using --cascade=false, the parent is deleted but the child still exists:

在以下示例中，有一个父级和一个子级。请注意，仍然包含所有者引用。如果我使用 --cascade=false 删除父级，则删除父级但子级仍然存在：

```
kubectl get configmap
NAME           DATA   AGE
mymap-child    0      13m8s
mymap-parent   0      13m8s

kubectl delete --cascade=false configmap/mymap-parent
configmap "mymap-parent" deleted

kubectl get configmap
NAME          DATA   AGE
mymap-child   0      13m21s

```




The --cascade option links to the propagation policy in the API, which allows you to change the order in which objects are deleted within a tree. In the following example uses API access to craft a custom delete API call with the background propagation policy:

--cascade 选项链接到 API 中的传播策略，它允许您更改对象在树中删除的顺序。在以下示例中，使用 API 访问权来制作具有后台传播策略的自定义删除 API 调用：

```
kubectl proxy --port=8080 &
Starting to serve on 127.0.0.1:8080

curl -X DELETE \
localhost:8080/api/v1/namespaces/default/configmaps/mymap-parent \
  -d '{ "kind":"DeleteOptions", "apiVersion":"v1", "propagationPolicy":"Background" }' \
  -H "Content-Type: application/json"
{
"kind": "Status",
"apiVersion": "v1",
"metadata": {},
"status": "Success",
"details": { ... }
}

```


Note that the propagation policy cannot be specified on the command line using kubectl. You have to specify it using a custom API call. Simply create a proxy, so you have access to the API server from the client, and execute a `curl` command with just a URL to execute that `delete` command.

请注意，不能使用 kubectl 在命令行上指定传播策略。您必须使用自定义 API 调用来指定它。只需创建一个代理，这样您就可以从客户端访问 API 服务器，并执行一个带有 URL 的 `curl` 命令来执行该 `delete` 命令。

There are three different options for the propagation policy:

传播策略有三种不同的选项：

- `Foreground`: Children are deleted before the parent (post-order)
- `Background`: Parent is deleted before the children (pre-order)
- `Orphan`: Owner references are ignored



- `Foreground`：在父级之前删除子级（后序）
- 后台：在孩子之前删除父母（预购）
- `Orphan`：忽略所有者引用

Keep in mind that when you delete an object and owner references have been specified, finalizers will be honored in the process. This can result in trees of objects persisting, and you end up with a partial deletion. At that point, you have to look at any existing owner references on your objects, as well as any finalizers, to understand what’s happening.

请记住，当您删除对象并指定所有者引用时，终结器将在此过程中受到尊重。这可能导致对象树持续存在，最终导致部分删除。此时，您必须查看对象上的任何现有所有者引用以及任何终结器，以了解发生了什么。

## Forcing a Deletion of a Namespace

## 强制删除命名空间

There's one situation that may require forcing finalization for a namespace. If you've deleted a namespace and you've cleaned out all of the objects under it, but the namespace still exists, deletion can be forced by updating the namespace subresource, `finalize`. This informs the namespace controller that it needs to remove the finalizer from the namespace and perform any cleanup:

有一种情况可能需要强制完成命名空间。如果您删除了一个命名空间并清除了它下的所有对象，但该命名空间仍然存在，则可以通过更新命名空间子资源“finalize”来强制删除。这会通知命名空间控制器它需要从命名空间中删除终结器并执行任何清理：

```
cat <<EOF |curl -X PUT \
localhost:8080/api/v1/namespaces/test/finalize \
  -H "Content-Type: application/json" \
  --data-binary @-
{
"kind": "Namespace",
"apiVersion": "v1",
"metadata": {
    "name": "test"
},
"spec": {
    "finalizers": null
}
}
EOF

```


This should be done with caution as it may delete the namespace only and leave orphan objects within the, now non-exiting, namespace - a confusing state for Kubernetes. If this happens, the namespace can be re-created manually and sometimes the orphaned objects will re-appear under the just-created namespace which will allow manual cleanup and recovery.

应谨慎执行此操作，因为它可能仅删除命名空间并将孤立对象留在现在不存在的命名空间中——这对 Kubernetes 来说是一种令人困惑的状态。如果发生这种情况，可以手动重新创建命名空间，有时孤立对象将重新出现在刚刚创建的命名空间下，这将允许手动清理和恢复。

## Key Takeaways

## 关键要点

As these examples demonstrate, finalizers can get in the way of deleting resources in Kubernetes, especially when there are parent-child relationships between objects. Often, there is a reason for adding a finalizer into the code, so you should always investigate before manually deleting it. Owner references allow you to specify and remove trees of resources, although finalizers will be honored in the process. Finally, the propagation policy can be used to specify the order of deletion via a custom API call, giving you control over how objects are deleted. Now that you know a little more about how deletions work in Kubernetes, we recommend you try it out on your own, using a test cluster.

正如这些示例所示，终结器可能会妨碍删除 Kubernetes 中的资源，尤其是当对象之间存在父子关系时。通常，在代码中添加终结器是有原因的，因此您应该始终在手动删除它之前进行调查。所有者引用允许您指定和删除资源树，尽管在此过程中会使用终结器。最后，传播策略可用于通过自定义 API 调用指定删除顺序，让您可以控制对象的删除方式。既然您对 Kubernetes 中的删除工作原理有了更多的了解，我们建议您使用测试集群自行尝试一下。

- [← Previous](http://kubernetes.io/blog/2021/04/23/kubernetes-release-1.21-metrics-stability-ga/)
[Next →](http://kubernetes.io/blog/2021/06/21/writing-a-controller-for-pod-labels/) 

- [← 上一页](http://kubernetes.io/blog/2021/04/23/kubernetes-release-1.21-metrics-stability-ga/)
[下一步→](http://kubernetes.io/blog/2021/06/21/writing-a-controller-for-pod-labels/)

