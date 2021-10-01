# Writing a Controller for Pod Labels

# 为 Pod 标签编写控制器

Monday, June 21, 2021

[Operators](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) are proving to be an excellent solution to running stateful distributed applications in Kubernetes. Open source tools like the [Operator SDK](https://sdk.operatorframework.io/) provide ways to build reliable and maintainable operators, making it easier to extend Kubernetes and implement custom scheduling.

[Operators](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) 被证明是在 Kubernetes 中运行有状态分布式应用程序的绝佳解决方案。 [Operator SDK](https://sdk.operatorframework.io/) 等开源工具提供了构建可靠且可维护的操作符的方法，使扩展 Kubernetes 和实现自定义调度变得更加容易。

Kubernetes operators run complex software inside your cluster. The open source community has already built [many operators](https://operatorhub.io/) for distributed applications like Prometheus, Elasticsearch, or Argo CD. Even outside of open source, operators can help to bring new functionality to your Kubernetes cluster.

Kubernetes 操作员在您的集群内运行复杂的软件。开源社区已经为 Prometheus、Elasticsearch 或 Argo CD 等分布式应用程序构建了 [许多运算符](https://operatorhub.io/)。即使在开源之外，运维人员也可以帮助为您的 Kubernetes 集群带来新功能。

An operator is a set of [custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) and a set of [controllers](https://kubernetes.io/docs/concepts/architecture/controller/). A controller watches for changes to specific resources in the Kubernetes API and reacts by creating, updating, or deleting resources.

一个操作符是一组[自定义资源](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)和一组[控制器](https://kubernetes.io/docs/concepts/architecture/controller/)。控制器监视 Kubernetes API 中特定资源的更改，并通过创建、更新或删除资源来做出反应。

The Operator SDK is best suited for building fully-featured operators. Nonetheless, you can use it to write a single controller. This post will walk you through writing a Kubernetes controller in Go that will add a `pod-name` label to pods that have a specific annotation.

Operator SDK 最适合构建功能齐全的 Operator。尽管如此，您可以使用它来编写单个控制器。这篇博文将引导您使用 Go 编写一个 Kubernetes 控制器，该控制器将为具有特定注释的 pod 添加一个 `pod-name` 标签。

## Why do we need a controller for this?

## 为什么我们需要一个控制器？

I recently worked on a project where we needed to create a Service that routed traffic to a specific Pod in a ReplicaSet. The problem is that a Service can only select pods by label, and all pods in a ReplicaSet have the same labels. There are two ways to solve this problem:

我最近参与了一个项目，我们需要创建一个服务，将流量路由到 ReplicaSet 中的特定 Pod。问题是一个Service只能通过标签来选择pod，一个ReplicaSet中的所有pod都拥有相同的标签。有两种方法可以解决这个问题：

1. Create a Service without a selector and manage the Endpoints or EndpointSlices for that Service directly. We would need to write a custom controller to insert our Pod's IP address into those resources.
2. Add a label to the Pod with a unique value. We could then use this label in our Service's selector. Again, we would need to write a custom controller to add this label.

1. 创建一个没有选择器的服务并直接管理该服务的端点或端点切片。我们需要编写一个自定义控制器来将 Pod 的 IP 地址插入这些资源中。
2. 给 Pod 添加一个具有唯一值的标签。然后我们可以在我们的服务选择器中使用这个标签。同样，我们需要编写一个自定义控制器来添加此标签。

A controller is a control loop that tracks one or more Kubernetes resource types. The controller from option n°2 above only needs to track pods, which makes it simpler to implement. This is the option we are going to walk through by writing a Kubernetes controller that adds a `pod-name` label to our pods.

控制器是跟踪一种或多种 Kubernetes 资源类型的控制循环。上面选项 n°2 中的控制器只需要跟踪 Pod，这使得实现更简单。这是我们将要通过编写一个 Kubernetes 控制器来为我们的 pod 添加一个 `pod-name` 标签的选项。

StatefulSets [do this natively](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-name-label) by adding a `pod-name` label to each Pod in the set. But what if we don't want to or can't use StatefulSets?

StatefulSets [本地执行此操作](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-name-label) 通过向集合中的每个 Pod 添加一个 `pod-name` 标签。但是如果我们不想或不能使用 StatefulSet 呢？

We rarely create pods directly; most often, we use a Deployment, ReplicaSet, or another high-level resource. We can specify labels to add to each Pod in the PodSpec, but not with dynamic values, so no way to replicate a StatefulSet's `pod-name` label.

我们很少直接创建 pod；大多数情况下，我们使用 Deployment、ReplicaSet 或其他高级资源。我们可以指定标签添加到 PodSpec 中的每个 Pod，但不能使用动态值，因此无法复制 StatefulSet 的 `pod-name` 标签。

We tried using a [mutating admission webhook](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#mutatingadmissionwebhook). When anyone creates a Pod, the webhook patches the Pod with a label containing the Pod's name. Disappointingly, this does not work: not all pods have a name before being created. For instance, when the ReplicaSet controller creates a Pod, it sends a `namePrefix` to the Kubernetes API server and not a `name`. The API server generates a unique name before persisting the new Pod to etcd, but only after calling our admission webhook. So in most cases, we can't know a Pod's name with a mutating webhook.

我们尝试使用 [mutating Admission webhook](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#mutatingadmissionwebhook)。当任何人创建 Pod 时，webhook 会使用包含 Pod 名称的标签修补 Pod。令人失望的是，这不起作用：并非所有 Pod 在创建之前都有名称。例如，当 ReplicaSet 控制器创建一个 Pod 时，它会向 Kubernetes API 服务器发送一个 `namePrefix` 而不是 `name`。 API 服务器在将新 Pod 持久化到 etcd 之前生成一个唯一名称，但仅在调用我们的准入 webhook 之后。所以在大多数情况下，我们无法通过变异的 webhook 知道 Pod 的名称。

Once a Pod exists in the Kubernetes API, it is mostly immutable, but we can still add a label. We can even do so from the command line:

一旦 Pod 存在于 Kubernetes API 中，它基本上是不可变的，但我们仍然可以添加标签。我们甚至可以从命令行执行此操作：

```bash
kubectl label my-pod my-label-key=my-label-value
```

We need to watch for changes to any pods in the Kubernetes API and add the label we want. Rather than do this manually, we are going to write a controller that does it for us.

我们需要观察 Kubernetes API 中任何 Pod 的变化并添加我们想要的标签。我们将编写一个控制器来为我们执行此操作，而不是手动执行此操作。

## Bootstrapping a controller with the Operator SDK

## 使用 Operator SDK 引导控制器

A controller is a reconciliation loop that reads the desired state of a resource from the Kubernetes API and takes action to bring the cluster's actual state closer to the desired state. 

控制器是一个协调循环，它从 Kubernetes API 读取资源的期望状态，并采取行动使集群的实际状态更接近期望状态。

In order to write this controller as quickly as possible, we are going to use the Operator SDK. If you don't have it installed, follow the [official documentation](https://sdk.operatorframework.io/docs/installation/).

为了尽快编写此控制器，我们将使用 Operator SDK。如果您没有安装，请按照[官方文档](https://sdk.operatorframework.io/docs/installation/)进行操作。

```terminal
$ operator-sdk version
operator-sdk version: "v1.4.2", commit: "4b083393be65589358b3e0416573df04f4ae8d9b", kubernetes version: "v1.19.4", go version: "go1.15.8", GOOS: "darwin", GOARCH: "amd64"
```

Let's create a new directory to write our controller in:

让我们创建一个新目录来写入我们的控制器：

```bash
mkdir label-operator && cd label-operator
```

Next, let's initialize a new operator, to which we will add a single controller. To do this, you will need to specify a domain and a repository. The domain serves as a prefix for the group your custom Kubernetes resources will belong to. Because we are not going to be defining custom resources, the domain does not matter. The repository is going to be the name of the Go module we are going to write. By convention, this is the repository where you will be storing your code.

接下来，让我们初始化一个新的操作符，我们将向其中添加一个控制器。为此，您需要指定域和存储库。该域用作您的自定义 Kubernetes 资源所属的组的前缀。因为我们不会定义自定义资源，所以域无关紧要。存储库将是我们将要编写的 Go 模块的名称。按照惯例，这是您将存储代码的存储库。

As an example, here is the command I ran:

例如，这是我运行的命令：

```bash
# Feel free to change the domain and repo values.
operator-sdk init --domain=padok.fr --repo=github.com/busser/label-operator
```

Next, we need a create a new controller. This controller will handle pods and not a custom resource, so no need to generate the resource code. Let's run this command to scaffold the code we need:

接下来，我们需要创建一个新的控制器。该控制器将处理 pods 而不是自定义资源，因此无需生成资源代码。让我们运行这个命令来搭建我们需要的代码：

```bash
operator-sdk create api --group=core --version=v1 --kind=Pod --controller=true --resource=false
```

We now have a new file: `controllers/pod_controller.go`. This file contains a `PodReconciler` type with two methods that we need to implement. The first is `Reconcile`, and it looks like this for now:

我们现在有一个新文件：`controllers/pod_controller.go`。该文件包含一个 `PodReconciler` 类型，其中包含我们需要实现的两个方法。第一个是 `Reconcile`，它现在看起来像这样：

```go
func (r *PodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    _ = r.Log.WithValues("pod", req.NamespacedName)

    // your logic here

    return ctrl.Result{}, nil
}
```

The `Reconcile` method is called whenever a Pod is created, updated, or deleted. The name and namespace of the Pod are in the `ctrl.Request` the method receives as a parameter.

每当创建、更新或删除 Pod 时，都会调用“Reconcile”方法。 Pod 的名称和命名空间位于该方法作为参数接收的 `ctrl.Request` 中。

The second method is `SetupWithManager` and for now it looks like this:

第二种方法是“SetupWithManager”，现在它看起来像这样：

```go
func (r *PodReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        // Uncomment the following line adding a pointer to an instance of the controlled resource as an argument
        // For().
        Complete(r)
}
```

The `SetupWithManager` method is called when the operator starts. It serves to tell the operator framework what types our `PodReconciler` needs to watch. To use the same `Pod` type used by Kubernetes internally, we need to import some of its code. All of the Kubernetes source code is open source, so you can import any part you like in your own Go code. You can find a complete list of available packages in the Kubernetes source code or [here on pkg.go.dev](https://pkg.go.dev/k8s.io/api). To use pods, we need the `k8s.io/api/core/v1` package.

`SetupWithManager` 方法在操作员启动时被调用。它用于告诉 operator 框架我们的 `PodReconciler` 需要观察什么类型。要使用 Kubernetes 内部使用的相同 `Pod` 类型，我们需要导入它的一些代码。所有 Kubernetes 源代码都是开源的，因此您可以在自己的 Go 代码中导入您喜欢的任何部分。您可以在 Kubernetes 源代码或 [pkg.go.dev 上的此处](https://pkg.go.dev/k8s.io/api) 中找到可用软件包的完整列表。要使用 pod，我们需要 `k8s.io/api/core/v1` 包。

```go
package controllers

import (
    // other imports...
    corev1 "k8s.io/api/core/v1"
    // other imports...
)
```

Lets use the `Pod` type in `SetupWithManager` to tell the operator framework we want to watch pods:

让我们使用 `SetupWithManager` 中的 `Pod` 类型来告诉 operator 框架我们想要观看 pods：

```go
func (r *PodReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&corev1.Pod{}).
        Complete(r)
}
```

Before moving on, we should set the RBAC permissions our controller needs. Above the `Reconcile` method, we have some default permissions:

在继续之前，我们应该设置我们的控制器需要的 RBAC 权限。在 `Reconcile` 方法上方，我们有一些默认权限：

```go
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=pods/finalizers,verbs=update
```

We don't need all of those. Our controller will never interact with a Pod's status or its finalizers. It only needs to read and update pods. Lets remove the unnecessary permissions and keep only what we need:

我们不需要所有这些。我们的控制器永远不会与 Pod 的状态或其终结器交互。它只需要读取和更新 pod。让我们删除不必要的权限并只保留我们需要的权限：

```go
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;update;patch
```

We are now ready to write our controller's reconciliation logic.

我们现在准备编写控制器的协调逻辑。

## Implementing reconciliation

## 实现和解

Here is what we want our `Reconcile` method to do:

这是我们希望我们的 `Reconcile` 方法执行的操作：

1. Use the Pod's name and namespace from the `ctrl.Request` to fetch the Pod from the Kubernetes API. 

1. 使用 `ctrl.Request` 中的 Pod 名称和命名空间从 Kubernetes API 中获取 Pod。

2. If the Pod has an `add-pod-name-label` annotation, add a `pod-name` label to the Pod; if the annotation is missing, don't add the label.
3. Update the Pod in the Kubernetes API to persist the changes made.

2.如果Pod有`add-pod-name-label`注解，则给Pod添加`pod-name`标签；如果缺少注释，请不要添加标签。
3. 更新 Kubernetes API 中的 Pod 以保留所做的更改。

Lets define some constants for the annotation and label:

让我们为注释和标签定义一些常量：

```go
const (
    addPodNameLabelAnnotation = "padok.fr/add-pod-name-label"
    podNameLabel              = "padok.fr/pod-name"
)
```

The first step in our reconciliation function is to fetch the Pod we are working on from the Kubernetes API:

我们协调功能的第一步是从 Kubernetes API 获取我们正在处理的 Pod：

```go
// Reconcile handles a reconciliation request for a Pod.
// If the Pod has the addPodNameLabelAnnotation annotation, then Reconcile
// will make sure the podNameLabel label is present with the correct value.
// If the annotation is absent, then Reconcile will make sure the label is too.
func (r *PodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    log := r.Log.WithValues("pod", req.NamespacedName)

    /*
        Step 0: Fetch the Pod from the Kubernetes API.
    */

    var pod corev1.Pod
    if err := r.Get(ctx, req.NamespacedName, &pod);err != nil {
        log.Error(err, "unable to fetch Pod")
        return ctrl.Result{}, err
    }

    return ctrl.Result{}, nil
}
```

Our `Reconcile` method will be called when a Pod is created, updated, or deleted. In the deletion case, our call to `r.Get` will return a specific error. Let's import the package that defines this error:

我们的 `Reconcile` 方法将在创建、更新或删除 Pod 时被调用。在删除的情况下，我们对 `r.Get` 的调用将返回一个特定的错误。让我们导入定义此错误的包：

```go
package controllers

import (
    // other imports...
    apierrors "k8s.io/apimachinery/pkg/api/errors"
    // other imports...
)
```

We can now handle this specific error and — since our controller does not care about deleted pods — explicitly ignore it:

我们现在可以处理这个特定的错误，并且——因为我们的控制器不关心删除的 pod——明确地忽略它：

```go
     /*
        Step 0: Fetch the Pod from the Kubernetes API.
    */

    var pod corev1.Pod
    if err := r.Get(ctx, req.NamespacedName, &pod);err != nil {
        if apierrors.IsNotFound(err) {
            // we'll ignore not-found errors, since we can get them on deleted requests.
            return ctrl.Result{}, nil
        }
        log.Error(err, "unable to fetch Pod")
        return ctrl.Result{}, err
    }
```

Next, lets edit our Pod so that our dynamic label is present if and only if our annotation is present:

接下来，让我们编辑我们的 Pod，以便当且仅当我们的注释存在时我们的动态标签才存在：

```go
     /*
        Step 1: Add or remove the label.
    */

    labelShouldBePresent := pod.Annotations[addPodNameLabelAnnotation] == "true"
    labelIsPresent := pod.Labels[podNameLabel] == pod.Name

    if labelShouldBePresent == labelIsPresent {
        // The desired state and actual state of the Pod are the same.
        // No further action is required by the operator at this moment.
        log.Info("no update required")
        return ctrl.Result{}, nil
    }

    if labelShouldBePresent {
        // If the label should be set but is not, set it.
        if pod.Labels == nil {
            pod.Labels = make(map[string]string)
        }
        pod.Labels[podNameLabel] = pod.Name
        log.Info("adding label")
    } else {
        // If the label should not be set but is, remove it.
        delete(pod.Labels, podNameLabel)
        log.Info("removing label")
    }
```

Finally, let's push our updated Pod to the Kubernetes API:

最后，让我们将更新后的 Pod 推送到 Kubernetes API：

```go
     /*
        Step 2: Update the Pod in the Kubernetes API.
    */

    if err := r.Update(ctx, &pod);err != nil {
        log.Error(err, "unable to update Pod")
        return ctrl.Result{}, err
    }
```

When writing our updated Pod to the Kubernetes API, there is a risk that the Pod has been updated or deleted since we first read it. When writing a Kubernetes controller, we should keep in mind that we are not the only actors in the cluster. When this happens, the best thing to do is start the reconciliation from scratch, by requeuing the event. Lets do exactly that:

将更新后的 Pod 写入 Kubernetes API 时，存在自我们第一次读取 Pod 以来已更新或删除的风险。在编写 Kubernetes 控制器时，我们应该记住，我们不是集群中唯一的参与者。发生这种情况时，最好的办法是重新排队事件，从头开始协调。让我们这样做：

```go
     /*
        Step 2: Update the Pod in the Kubernetes API.
    */

    if err := r.Update(ctx, &pod);err != nil {
        if apierrors.IsConflict(err) {
            // The Pod has been updated since we read it.
            // Requeue the Pod to try to reconciliate again.
            return ctrl.Result{Requeue: true}, nil
        }
        if apierrors.IsNotFound(err) {
            // The Pod has been deleted since we read it.
            // Requeue the Pod to try to reconciliate again.
            return ctrl.Result{Requeue: true}, nil
        }
        log.Error(err, "unable to update Pod")
        return ctrl.Result{}, err
    }
```

Let's remember to return successfully at the end of the method:

让我们记住在方法结束时成功返回：

```go
     return ctrl.Result{}, nil
}
```

And that's it! We are now ready to run the controller on our cluster.

就是这样！我们现在准备在我们的集群上运行控制器。

## Run the controller on your cluster

## 在集群上运行控制器

To run our controller on your cluster, we need to run the operator. For that, all you will need is `kubectl`. If you don't have a Kubernetes cluster at hand, I recommend you start one locally with [KinD (Kubernetes in Docker)](https://kind.sigs.k8s.io/docs/user/quick-start/#installation).

要在您的集群上运行我们的控制器，我们需要运行操作员。为此，您只需要 `kubectl`。如果您手头没有 Kubernetes 集群，我建议您使用 [KinD (Kubernetes in Docker)](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)在本地启动一个)。

All it takes to run the operator from your machine is this command:

从您的机器运行操作员只需以下命令：

```bash
make run
```

After a few seconds, you should see the operator's logs. Notice that our controller's `Reconcile` method was called for all pods already running in the cluster.

几秒钟后，您应该会看到操作员的日志。请注意，我们的控制器的 `Reconcile` 方法已被集群中已经运行的所有 pod 调用。

Let's keep the operator running and, in another terminal, create a new Pod:

让我们保持操作符运行，并在另一个终端中创建一个新的 Pod：

```bash
kubectl run --image=nginx my-nginx
```

The operator should quickly print some logs, indicating that it reacted to the Pod's creation and subsequent changes in status:

操作员应该快速打印一些日志，表明它对 Pod 的创建和随后的状态变化做出了反应：

```text
INFO    controllers.Pod no update required  {"pod": "default/my-nginx"}
INFO    controllers.Pod no update required  {"pod": "default/my-nginx"}
INFO    controllers.Pod no update required  {"pod": "default/my-nginx"}
INFO    controllers.Pod no update required  {"pod": "default/my-nginx"}
```

Lets check the Pod's labels:

让我们检查 Pod 的标签：

```terminal
$ kubectl get pod my-nginx --show-labels
NAME       READY   STATUS    RESTARTS   AGE   LABELS
my-nginx   1/1     Running   0          11m   run=my-nginx
```

Let's add an annotation to the Pod so that our controller knows to add our dynamic label to it:

让我们向 Pod 添加一个注解，以便我们的控制器知道将我们的动态标签添加到它：

```bash
kubectl annotate pod my-nginx padok.fr/add-pod-name-label=true
```

Notice that the controller immediately reacted and produced a new line in its logs:

请注意，控制器立即做出反应并在其日志中生成了一个新行：

```text
INFO    controllers.Pod adding label    {"pod": "default/my-nginx"}
$ kubectl get pod my-nginx --show-labels
NAME       READY   STATUS    RESTARTS   AGE   LABELS
my-nginx   1/1     Running   0          13m   padok.fr/pod-name=my-nginx,run=my-nginx
```

Bravo! You just successfully wrote a Kubernetes controller capable of adding labels with dynamic values to resources in your cluster.

太棒了！您刚刚成功编写了一个 Kubernetes 控制器，该控制器能够为集群中的资源添加带有动态值的标签。

Controllers and operators, both big and small, can be an important part of your Kubernetes journey. Writing operators is easier now than it has ever been. The possibilities are endless.

控制器和操作员，无论大小，都可以成为 Kubernetes 旅程的重要组成部分。现在编写运算符比以往任何时候都容易。可能性是无止境。

## What next?[ ](https://kubernetes.io/blog/2021/06/21/writing-a-controller-for-pod-labels/#what-next)

## 下一步是什么？[ ](https://kubernetes.io/blog/2021/06/21/writing-a-controller-for-pod-labels/#what-next)

If you want to go further, I recommend starting by deploying your controller or operator inside a cluster. The `Makefile` generated by the Operator SDK will do most of the work.

如果您想更进一步，我建议首先在集群中部署您的控制器或操作员。 Operator SDK 生成的`Makefile` 将完成大部分工作。

When deploying an operator to production, it is always a good idea to implement robust testing. The first step in that direction is to write unit tests. [This documentation](https://sdk.operatorframework.io/docs/building-operators/golang/testing/) will guide you in writing tests for your operator. I wrote tests for the operator we just wrote; you can find all of my code in [this GitHub repository](https://github.com/busser/label-operator).

在将操作员部署到生产环境时，实施稳健的测试始终是一个好主意。朝着这个方向迈出的第一步是编写单元测试。 [本文档](https://sdk.operatorframework.io/docs/building-operators/golang/testing/) 将指导您为您的运营商编写测试。我为我们刚刚编写的运算符编写了测试；你可以在 [this GitHub repository](https://github.com/busser/label-operator) 中找到我的所有代码。

## How to learn more?

## 如何学习更多？

The [Operator SDK documentation](https://sdk.operatorframework.io/docs/) goes into detail on how you can go further and implement more complex operators.

[Operator SDK 文档](https://sdk.operatorframework.io/docs/) 详细介绍了如何进一步实现更复杂的运算符。

When modeling a more complex use-case, a single controller acting on built-in Kubernetes types may not be enough. You may need to build a more complex operator with [Custom Resource Definitions (CRDs)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) and multiple controllers. The Operator SDK is a great tool to help you do this.

在对更复杂的用例进行建模时，对内置 Kubernetes 类型起作用的单个控制器可能是不够的。您可能需要使用 [自定义资源定义 (CRD)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) 和多个控制器构建更复杂的操作符。 Operator SDK 是一个很好的工具，可以帮助您做到这一点。

If you want to discuss building an operator, join the [#kubernetes-operator](https://kubernetes.slack.com/messages/kubernetes-operators) channel in the [Kubernetes Slack workspace](https://slack.k8s.io/)! 

如果您想讨论构建一个操作员，请加入 [Kubernetes Slack 工作区](https://slack.k8s) 中的 [#kubernetes-operator](https://kubernetes.slack.com/messages/kubernetes-operators)频道.io/)！

