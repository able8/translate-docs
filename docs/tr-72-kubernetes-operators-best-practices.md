# Kubernetes Operators Best Practices

# Kubernetes Operators 最佳实践

2019 年 6 月 11 日 From https://cloud.redhat.com/blog/kubernetes-operators-best-practices

## Introduction

## 介绍

Kubernetes Operators are processes connecting to the master API and watching for events, typically on a limited number of resource types.

Kubernetes Operator 是连接到主 API 并监视事件的进程，通常在有限数量的资源类型上。

When a relevant event occurs, the operator reacts and performs a specific action. This may be limited to interacting with the master API only, but will often involve performing some action on some other systems (this could be either in cluster or off cluster resources).

当相关事件发生时，操作员会做出反应并执行特定操作。这可能仅限于与主 API 交互，但通常会涉及在其他一些系统上执行某些操作（这可能在集群中或集群外资源）。

![](https://assets.openshift.com/hubfs/Imported_Blog_Media/rafop1.png)

Operators are implemented as a collection of controllers where each controller watches a specific resource type. When a relevant event occurs on a watched resource a **reconcile cycle** is started.

Operator 被实现为控制器的集合，其中每个控制器监视特定的资源类型。当相关事件发生在受监视的资源上时，**协调循环** 开始。

During the reconcile cycle, the controller has the responsibility to check that current state matches the desired state described by the watched resource. Interestingly, by design, the event is not passed to the reconcile cycle, which is then forced to consider the whole state of the instance that was referenced by the event.

在协调周期中，控制器有责任检查当前状态是否与被监视资源描述的期望状态相匹配。有趣的是，按照设计，事件不会传递到协调循环，然后它被迫考虑事件引用的实例的整个状态。

This approach is referred to as [level-based, as opposed to edge-based](http://venkateshabbarapu.blogspot.com/2013/03/edge-triggered-vs-level-triggered.html). Deriving from electronic circuit design, level-based triggering is the idea of receiving an event (an interrupt for example) and reacting to a state, while edge-based triggering is the idea of receiving an event and reacting to a state variation.

这种方法被称为[基于电平触发，而不是基于边缘](http://venkateshabbarapu.blogspot.com/2013/03/edge-triggered-vs-level-triggered.html)。源自电子电路设计，基于电平的触发是接收事件（例如中断）并对状态做出反应的想法，而基于边缘的触发是接收事件并对状态变化做出反应的想法。

Level-based triggering, while arguably less efficient because it forces to re-evaluate the entire state as opposed to just what changed, is considered more suitable in complex and unreliable environments where signals can be lost or retransmitted multiple times.

基于电平的触发虽然效率较低，因为它强制重新评估整个状态而不是仅仅改变状态，但被认为更适合于信号可能丢失或多次重新传输的复杂和不可靠的环境。

This design choice influences how we write controller’s code.

这种设计选择会影响我们编写控制器代码的方式。

Also relevant to this discussion is an understanding of the lifecycle of an API request. The following diagram provides a high level summary:

与此讨论相关的还有对 API 请求生命周期的理解。下图提供了一个高层次的总结：

![](https://assets.openshift.com/hubfs/Imported_Blog_Media/rafop2.png)

When a request is made to the API server, especially for create and delete requests, the request goes through the above phases. Notice that it is possible to specify webhooks to perform [mutations and validations](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#experimenting-with-admission-webhooks). If the operator introduces a new [custom resource definition](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) (CRD), we may have to also define those webhooks. Normally, the operator process would also implement the webhook endpoint by listening on a port.

当向 API 服务器发出请求时，尤其是创建和删除请求时，请求会经历上述阶段。请注意，可以指定 webhooks 来执行 [突变和验证](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#experimenting-with-admission-webhooks)。如果Operators引入了新的 自定义资源，我们可能还需要定义那些 webhooks。通常，操作员进程也会通过侦听端口来实现 webhook 端点。

This document presents a set of best practices to keep in mind when designing and developing operators using the [Operator SDK](https://github.com/operator-framework/operator-sdk).

本文档介绍了在使用 [Operator SDK](https://github.com/operator-framework/operator-sdk) 设计和开发运算符时要牢记的一组最佳实践。

If your operator introduces a new CRD, the Operator SDK will assist in scaffolding it. To make sure your CRD conforms to the Kubernetes best practices for extending the API, follow [these conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions. md#api-conventions).

如果您的Operators引入了新的 CRD，Operators SDK 将协助搭建它。为确保您的 CRD 符合 Kubernetes 扩展 API 的最佳实践，请遵循 [这些约定](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions。 md#api-约定）。

All the best practices mentioned in this article are portrayed in an example available at the [operator-utils](https://github.com/redhat-cop/operator-utils) repository. This repository is also a library which you can import in your operator, giving you some useful utilities for writing your own operators.

本文中提到的所有最佳实践都在 [operator-utils](https://github.com/redhat-c​​op/operator-utils) 存储库中提供的示例中进行了描述。这个存储库也是一个你可以导入到你的操作符中的库，为你提供一些有用的实用程序来编写你自己的操作符。

Finally this set of best practices for writing operators represents my personal view and should not be considered an official list of best practices from Red Hat.

最后，这组编写操作符的最佳实践代表了我的个人观点，不应被视为 Red Hat 的官方最佳实践列表。

# Creating watches

# 创建 watch 

As we said, controllers watch events on resources. This is done through the abstraction of watches.

正如我们所说，控制器在资源上观察事件。这是通过 watch 的抽象来完成的。

A watch is a mechanism to receive events of a certain type (either a core type or a CRD). A watch is normally created by specifying the following:

1. The resource type to watch.
2. A handler. The handler maps the events on the watched type to one or more instances for which the reconcile cycle is called. Watched type and instance type do not have to be the same.
3. A predicate. The predicate is a set of functions that can be customized to filter only the events we are interested in.

watch 是一种接收某种类型（核心类型或 CRD）事件的机制。通常通过指定以下内容来创建监视：

1. 要观看的资源类型。
2. 处理程序。处理程序将被监视类型上的事件映射到一个或多个调用协调周期的实例。监视类型和实例类型不必相同。
3. 断言。断言是一组可以自定义的函数，可以只过滤我们感兴趣的事件。

The diagram below captures these contexts:

下图捕获了这些上下文：

![](https://assets.openshift.com/hubfs/Imported_Blog_Media/rafop4.png)

In general, opening multiple watches on the same kind is acceptable because the watches are multiplexed. 
一般情况下，在同一种情况下打开多个 watch 是可以接受的，因为观察者是多路复用的。

You should also try to filter events as much as possible. Here, for example, is a predicate that filters events on secrets. Here we are interested only in events on secrets of type _kubernetes.io/tls_ which have a certain annotation:

您还应该尝试尽可能多地过滤事件。例如，这里是一个过滤秘密事件的断言。在这里，我们只对 _kubernetes.io/tls_ 类型的机密事件感兴趣，这些事件具有特定的注释：

```
isAnnotatedSecret := predicate.Funcs{
     UpdateFunc: func(e event.UpdateEvent) bool {
         oldSecret, ok := e.ObjectOld.(*corev1.Secret)
         if !ok {
             return false
         }
         newSecret, ok := e.ObjectNew.(*corev1.Secret)
         if !ok {
             return false
         }
         if newSecret.Type != util.TLSSecret {
             return false
         }
         oldValue, _ := e.MetaOld.GetAnnotations()[certInfoAnnotation]
         newValue, _ := e.MetaNew.GetAnnotations()[certInfoAnnotation]
         old := oldValue == "true"
         new := newValue == "true"
         // if the content has changed we trigger if the annotation is there
         if !reflect.DeepEqual(newSecret.Data[util.Cert], oldSecret.Data[util.Cert]) ||
             !reflect.DeepEqual(newSecret.Data[util.CA], oldSecret.Data[util.CA]) {
             return new
         }
         // otherwise we trigger if the annotation has changed
         return old != new
     },
     CreateFunc: func(e event.CreateEvent) bool {
         secret, ok := e.Object.(*corev1.Secret)
         if !ok {
             return false
         }
         if secret.Type != util.TLSSecret {
             return false
         }
         value, _ := e.Meta.GetAnnotations()[certInfoAnnotation]
         return value == "true"
     },
 }
```

A very common pattern is to observe events on the resources that we create (and we own) and to schedule a reconcile cycle on the CR that owns those resources, to do so you can use the _EnqueueRequestForOwner_ handler. This can be done as follows:

一个非常常见的模式是观察我们创建（并且我们拥有）的资源上的事件，并在拥有这些资源的 CR 上安排协调周期，为此您可以使用 _EnqueueRequestForOwner_ 处理程序。这可以按如下方式完成：

```
err = c.Watch(&source.Kind{Type: &examplev1alpha1.MyControlledType{}}, &handler.EnqueueRequestForOwner{})
```

A less common situation is where an event is multicast to several destination resources. Consider the case of a controller that injects TLS secrets into routes based on an annotation. Multiple routes in the same namespace can point to the same secret. If the secret changes we need to update all the routes. So, we would need to create a watch on the secret type and the handler would look as follows:

一种不太常见的情况是将事件多播到多个目标资源。考虑控制器根据注释将 TLS 秘密注入路由的情况。同一个命名空间中的多个路由可以指向同一个秘密。如果秘密发生变化，我们需要更新所有路由。因此，我们需要在秘密类型上创建一个监视，处理程序将如下所示：

```
type enqueueRequestForReferecingRoutes struct {
         client.Client
 }

 // trigger a router reconcile event for those routes that reference this secret
 func (e *enqueueRequestForReferecingRoutes) Create(evt event.CreateEvent, q workqueue.RateLimitingInterface) {
         routes, _ := matchSecret(e.Client, types.NamespacedName{
                 Name:      evt.Meta.GetName(),
                 Namespace: evt.Meta.GetNamespace(),
         })
         for _, route := range routes {
                 q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
                         Namespace: route.GetNamespace(),
                         Name:      route.GetName(),
                 }})
         }
 }

 // Update implements EventHandler
 // trigger a router reconcile event for those routes that reference this secret
 func (e *enqueueRequestForReferecingRoutes) Update(evt event.UpdateEvent, q workqueue.RateLimitingInterface) {
         routes, _ := matchSecret(e.Client, types.NamespacedName{
                 Name:      evt.MetaNew.GetName(),
                 Namespace: evt.MetaNew.GetNamespace(),
         })
         for _, route := range routes {
                 q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
                         Namespace: route.GetNamespace(),
                         Name:      route.GetName(),
                 }})
         }
 }

```


# Resource Reconciliation Cycle

# 资源协调周期

The reconciliation cycle is where the framework gives us back control after a watch has passed up an event. As explained before, at this point we don’t have the information about the type of event because we are working on level-based triggers.

协调周期是在监视传递事件后框架将控制权交给我们的地方。如前所述，此时我们没有关于事件类型的信息，因为我们正在研究基于级别的触发器。

Below is a model of what a common reconciliation cycle for a controller that manages a CRD could look like. As with every model, this is not going to completely reflect any particular use case, but I hope you find it useful to help think about the problem one needs to solve when writing an operator.

下面是管理 CRD 的控制器的常见协调周期的模型。与每个模型一样，这不会完全反映任何特定用例，但我希望您发现它有助于思考编写运算符时需要解决的问题。

![](https://assets.openshift.com/hubfs/Imported_Blog_Media/rafop3.png)

As we can see from the diagram the main steps are:

1. Retrieve the interested CR instance.
2. Manage the instance validity. We don’t want to try to do anything on an instance that does not carry valid values.
3. Manage instance initialization. If some values of the instance are not initialized, this section will take care of it.
4. Manage instance deletion. If the instance is being deleted and we need to do some specific clean up, this is where we manage it. 

从图中我们可以看出，主要步骤是：

1. 检索感兴趣的 CR 实例。
2. 管理实例有效性。我们不想尝试在不携带有效值的实例上做任何事情。
3. 管理实例初始化。如果实例的某些值未初始化，本节将处理它。
4. 管理实例删除。如果实例被删除，我们需要做一些特定的清理，这就是我们管理它的地方。
5. Manage controller business logic. If the above steps all pass we can finally manage and execute the reconciliation logic for this particular instance. This will be very controller specific.
6. 管理控制器业务逻辑。如果上述步骤全部通过，我们最终可以管理和执行此特定实例的对帐逻辑。这将是非常特定于控制器的。

In the rest of this section you can find some more in depth considerations on each of these steps.

在本节的其余部分，您可以找到有关每个步骤的更深入的注意事项。

## Resource Validation

## 资源验证

Two types of validation exist: Syntactic validation and semantic validation.

- **Syntactic validation** happens by defining OpenAPI validation rules.
- **Semantic Validation** can be done by creating a ValidatingAdmissionConfiguration.

存在两种类型的验证：句法验证和语义验证。

- **语法验证**通过定义 OpenAPI 验证规则进行。
- **语义验证**可以通过创建一个 ValidatingAdmissionConfiguration 来完成。

**Note**: it is not possible to validate a CR within a controller. Once the CR is accepted by the API server it will be stored in etcd. Once it is in etcd, the controller that owns it cannot do anything to reject it and if the CR is not valid, trying to use/process it will result in an error.

**注意**：无法在控制器内验证 CR。一旦 CR 被 API 服务器接受，它将存储在 etcd 中。一旦它在 etcd 中，拥有它的控制器就不能做任何事情来拒绝它，如果 CR 无效，尝试使用/处理它会导致错误。

**Recommendation**: because we cannot guarantee that a ValidatingAdmissionConfiguration will be created and or working, we should also validate the CR from within the controller and if they are not valid avoid creating an endless error-loop (see also: [error management ](http://cloud.redhat.com#error)).

**建议**：因为我们不能保证 ValidatingAdmissionConfiguration 将被创建和/或工作，我们还应该从控制器内部验证 CR，如果它们无效，请避免创建一个无限的错误循环（另请参阅：[错误管理](http://cloud.redhat.com#error))。

### Syntactic Validation

### 语法验证

OpenAPI validation rules can be added as described [here](https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html).

可以按照 [此处](https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html) 的描述添加 OpenAPI 验证规则。

**Recommendation**: model as much of the custom resource as possible of your validation as syntactic validation. Syntactic validation is relatively straightforward and prevents badly formed CRs from being stored in etcd, so it should be used as much as possible.

**建议**：将尽可能多的自定义资源建模为语法验证。语法验证相对简单，可以防止格式错误的 CR 存储在 etcd 中，因此应尽可能使用它。

### Semantic Validation

### 语义验证

Semantic validation is about making sure that fields have sensible values and that the entire resource record is meaningful. Semantic validation business logic depends on the concept that the CR represents and must be coded by the operator developer.

语义验证是为了确保字段具有合理的值以及整个资源记​​录是有意义的。语义验证业务逻辑取决于 CR 所代表的概念，并且必须由操作员开发人员进行编码。

If semantic validation is required by the given CR, then the operator should expose a webhook and [ValidatingAdmissionConfiguration](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/) should be created as part of the operator deployment.

如果给定的 CR 需要语义验证，那么操作员应该公开一个 webhook 并且应该创建 [ValidatingAdmissionConfiguration](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/)作为Operators部署的一部分。

The following limitations currently exist:

1. In OpenShift 3.11, ValidatingAdmissionConfigurations are in tech preview (they are supported from 4.1 on).
2. The Operator SDK has no support for scaffolding webhooks. This can be worked around using [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder), for example: kubebuilder webhook --group crew --version v1 --kind FirstMate --type=mutating --operations =create,update

目前存在以下限制：

1. 在 OpenShift 3.11 中，ValidatingAdmissionConfigurations 处于技术预览中（从 4.1 开始支持它们）。
2. Operator SDK 不支持脚手架 webhook。这可以使用 [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) 解决，例如：kubebuilder webhook --group team --version v1 --kind FirstMate --type=mutating --operations =创建，更新

### Validating a resource in the controller

### 验证控制器中的资源

It is better to reject an invalid CR rather than to accept it in etcd and then manage the error condition. That said, there could be situations in which the ValidatingAdmissionConfiguration is not deployed or not available at all. I think it is still a good practice to do semantic validation in the controller code. Code should be structured in such a way that you can share the same validation routine between the  ValidatingAdmissionConfiguration and the controller.

最好拒绝一个无效的 CR，而不是在 etcd 中接受它然后管理错误条件。也就是说，可能存在 ValidatingAdmissionConfiguration 未部署或根本不可用的情况。我认为在控制器代码中进行语义验证仍然是一个很好的做法。代码的结构应使您可以在 ValidatingAdmissionConfiguration 和控制器之间共享相同的验证例程。

The code of the controller calling the validation method should look like this:

调用验证方法的控制器代码应如下所示：

```
if ok, err := r.IsValid(instance); !ok {<br/>
     return r.ManageError(instance, err)<br/>
 }
```

Note that if the validation fails, we manage this error as described in the [error management section](http://cloud.redhat.com#error).

请注意，如果验证失败，我们会按照 [错误管理部分](http://cloud.redhat.com#error) 中的描述管理此错误。

The _IsValid_ function will look something like:

_IsValid_ 函数将类似于：

```
func (r *ReconcileMyCRD) IsValid(obj metav1.Object) (bool, error) {<br/>
     mycrd, ok := obj.(*examplev1alpha1.MyCRD)<br/>
 // validation logic<br/>
}
```


## Resource Initialization

## 资源初始化

One of the nice conventional features of Kubernetes resources is that only the needed fields of a resource are to be initialized by the user and the others can be omitted. This is the point of view of the user, but from the point of view of the coder and anyone debugging what is happening with a resource it is actually better to have all the fields initialized.This allows writing code without always checking if a field is defined, and allows for easy troubleshooting of error situations.In order to initialize resources there are two options:

1. Define an initialization method in the controller.
2. Define a MutatingAdmissionConfiguration (the procedure issimilar to the ValidatingAdmissionConfiguration).

Kubernetes 资源的一个很好的传统特性是，用户只需要初始化资源所需的字段，其他字段可以省略。这是用户的观点，但从编码人员和任何调试资源正在发生的事情的人的角度来看，实际上最好初始化所有字段。这允许编写代码而不必总是检查字段是否正确定义，并允许轻松排除错误情况。为了初始化资源，有两个选项：

1. 在控制器中定义一个初始化方法。
2. 定义一个 MutatingAdmissionConfiguration（过程类似于ValidatingAdmissionConfiguration）。

**Recommendation:** define an initialization method in the controller. The code should look like this sample:

**建议：** 在控制器中定义一个初始化方法。代码应类似于此示例：

```
if ok := r.IsInitialized(instance); !ok {<br/>
     err := r.GetClient().Update(context.TODO(), instance)<br/>
     if err != nil {<br/>
         log.Error(err, "unable to update instance", "instance", instance)<br/>
         return r.ManageError(instance, err)<br/>
     }<br/>
     return reconcile.Result{}, nil<br/>
 }
```

Notice that if the result is true, we update the instance and then we return. This will trigger another immediate reconcile cycle. This second time the initialize method will return false, and the logic will continue to the next phase.

请注意，如果结果为真，我们更新实例然后返回。这将触发另一个立即协调循环。第二次 initialize 方法将返回 false，逻辑将继续到下一阶段。

## Resource Finalization

## 资源终结

If resources are not owned by the CR controlled by your operator but action needs to be taken when that CR is deleted, you must use a [finalizer](https://kubernetes.io/docs/tasks/access-kubernetes-api/ custom-resources/custom-resource-definitions/#finalizers).

如果资源不属于您的Operators控制的 CR，但需要在删除该 CR 时采取行动，您必须使用 finalizer。

Finalizers provide a mechanism to inform the Kubernetes control plane that an action needs to take place before the standard Kubernetes garbage collection logic can be performed.

终结器提供了一种机制来通知 Kubernetes 控制平面在标准 Kubernetes 垃圾收集逻辑可以执行之前需要发生一个操作。

One or more finalizers can be set on resources. Each controller should manage its own finalizer and ignore others if present.

可以在资源上设置一个或多个终结器。每个控制器都应该管理自己的终结器并忽略其他控制器（如果存在）。

This is the pseudo code algorithm to manage finalizers:

1. If needed, add finalizers during the initialization method.

2. If the resource is being deleted, check if the finalizer owned by this controller is present.
3. If not, return
4. If yes, execute the cleanup logic
   1. If successful, update the CR by removing the finalizer.
   2. If failure decide whether to retry or give up and likely leave garbage (in some situations this can be acceptable).

这是管理终结器的伪代码算法：

1. 如果需要，在初始化方法期间添加终结器。

2. 如果正在删除资源，请检查该控制器拥有的终结器是否存在。
3. 如果没有，返回
4. 如果是，则执行清理逻辑
   1. 如果成功，则通过移除终结器来更新 CR。
      2 .  如果失败决定是重试还是放弃并可能留下垃圾（在某些情况下这是可以接受的）。

If your clean-up logic requires creating additional resources, do keep in mind that additional resources cannot be created in a namespace that is being deleted. A To-be-deleted namespace will trigger a delete of all in the included resources including your CR with the finalizer.

如果您的清理逻辑需要创建额外的资源，请记住不能在被删除的命名空间中创建额外的资源。 To-be-deleted 命名空间将触发删除所有包含的资源，包括带有终结器的 CR。

See an example of the code here:

在此处查看代码示例：

```
if util.IsBeingDeleted(instance) {

    if !util.HasFinalizer(instance, controllerName) {

        return reconcile.Result{}, nil

    }

    err := r.manageCleanUpLogic(instance)

    if err != nil {

        log.Error(err, "unable to delete instance", "instance", instance)

        return r.ManageError(instance, err)

    }

    util.RemoveFinalizer(instance, controllerName)

    err = r.GetClient().Update(context.TODO(), instance)

    if err != nil {

        log.Error(err, "unable to update instance", "instance", instance)

        return r.ManageError(instance, err)

    }

    return reconcile.Result{}, nil

}
```


## Resource Ownership

## 资源所有权

Resource ownership is a native concept in Kubernetes that determines how resources are deleted. When a resource is deleted and it owns other resources the children resources will be, by default, also deleted (you can disable this behavior, by setting cascade=false).

资源所有权是 Kubernetes 中的一个原生概念，它决定了如何删除资源。当一个资源被删除并且它拥有其他资源时，默认情况下，子资源也将被删除（您可以通过设置cascade=false 禁用此行为）。

This behavior is instrumental to guarantee correct [garbage collection](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/) of resources especially when resources control other resources in a multilevel hierarchy (think deployment-> repilcaset->pod).

这种行为有助于保证资源的正确[垃圾收集](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/)，尤其是当资源控制多级层次结构中的其他资源时（想想部署-> repilcaset->pod）。

**Recommendation**: If your controller creates resources and these resource’s lifecycle is tied to a resource (either core or CR), then you should set this resource as the owner of those resources. This can be done as follows:

**建议**：如果您的控制器创建资源并且这些资源的生命周期与资源（核心或 CR）相关联，那么您应该将此资源设置为这些资源的所有者。这可以按如下方式完成：

```
controllerutil.SetControllerReference(owner, obj, r.GetScheme())
```

Some other rules around ownership are the following:

关于所有权的其他一些规则如下：

1. The owner object must be in the same namespace as the owned object.
2. A namespaced resource can own a cluster level resource. We have to be careful here. An object can have a list of owners. If multiple namespaced objects own the same cluster-level object then each should claim ownership without overwriting the others’ ownership (the above API takes care of that).
3. A cluster level resource cannot own a namespaced resource.
4. A cluster level object can own another cluster level object.

5. 拥有者对象必须与拥有者在同一个命名空间中。
6. 命名空间资源可以拥有集群级别的资源。我们在这里必须小心。一个对象可以有一个所有者列表。如果多个命名空间对象拥有相同的集群级对象，那么每个对象都应该声明所有权而不覆盖其他人的所有权（上面的 API 会处理这个问题）。
7. 集群级资源不能拥有命名空间资源。
8. 一个集群级别的对象可以拥有另一个集群级别的对象。

## Managing status

## 管理状态

Status is a standard section of a resource. Status is obviously used to report on the status of the resource. In this document we are going to use status to report on the outcome of the last execution of the reconcile cycle. You can use status to add more information.

状态是资源的标准部分。状态显然用于报告资源的状态。在本文档中，我们将使用 status 报告最后一次执行协调周期的结果。您可以使用状态添加更多信息。

Under normal circumstances, If we were updating our resource every time we execute the reconcile cycle, this would trigger an update event which in turn would trigger a reconcile cycle in an endless loop.

在正常情况下，如果我们每次执行协调循环时都更新我们的资源，这将触发一个更新事件，进而触发一个无限循环的协调循环。

For this reason, **Status** should be modeled as a subresource as explained here.

出于这个原因，**状态** 应该被建模为一个子资源，如此处所述。

This way when we can update the status of our resources without increasing the _ResourceGeneration_ metadata field. We can update the status with this command:

这样我们就可以在不增加 _ResourceGeneration_ 元数据字段的情况下更新资源的状态。我们可以使用以下命令更新状态：

```
err = r.Status().Update(context.Background(), instance)
```


Now we need to write a predicate for our watch (see the section about watches for more details on these concepts) that will discard updates that did not increase the _ResourceGeneration_, this can be done using the [GenerationChangePredicate](https://github. com/operator-framework/operator-sdk/blob/master/pkg/predicate/predicate.go#L27)

现在我们需要为我们的 watch 编写一个断言（有关这些概念的更多详细信息，请参阅关于 watch 的部分），它将丢弃没有增加 _ResourceGeneration_ 的更新，这可以使用 [GenerationChangePredicate](https://github.com) 来完成。 com/operator-framework/operator-sdk/blob/master/pkg/predicate/predicate.go#L27)

If you recall, if we are using a finalizer, the finalizer should be set up at initialization time. If the finalizer is the only item that is being initialized, since it is a portion of the metadata field, the _ResourceGeneration_ will not be incremented. To account for that use case, the following  is a modified version of the predicate:

如果你还记得，如果我们使用终结器，终结器应该在初始化时设置。如果终结器是唯一被初始化的项目，因为它是元数据字段的一部分，_ResourceGeneration_ 将不会增加。考虑到该用例，以下是断言的修改版本：

```
type resourceGenerationOrFinalizerChangedPredicate struct {	predicate.Funcs}// Update implements default UpdateEvent filter for validating resource version changefunc (resourceGenerationOrFinalizerChangedPredicate) Update(e event.UpdateEvent) bool {	if e.MetaNew.GetGeneration() == e.MetaOld.GetGeneration() && reflect.DeepEqual(e.MetaNew.GetFinalizers(), e.MetaOld.GetFinalizers()) {		return false	}	return true}
```


Now assuming your status is as follows:

现在假设您的状态如下：

```
type MyCRStatus struct {	// +kubebuilder:validation:Enum=Success,Failure	Status     string      `json:"status,omitempty"`	LastUpdate metav1.Time `json:"lastUpdate,omitempty"`	Reason     string      `json:"reason,omitempty"`}
```


You can write a function to manage the successful execution of a reconciliation cycle:

您可以编写一个函数来管理对帐周期的成功执行：

``` 
func (r *ReconcilerBase) ManageSuccess(obj metav1.Object) (reconcile.Result, error) {	runtimeObj, ok := (obj).(runtime.Object)	if !ok {		log.Error(errors.New("not a runtime.Object"), "passed object was not a runtime.Object", "object", obj)		return reconcile.Result{}, nil	}	if reconcileStatusAware, updateStatus := (obj).(apis.ReconcileStatusAware); updateStatus {		status := apis.ReconcileStatus{			LastUpdate: metav1.Now(),			Reason:     "",			Status:     "Success",		}		reconcileStatusAware.SetReconcileStatus(status)		err := r.GetClient().Status().Update(context.Background(), runtimeObj)		if err != nil {			log.Error(err, "unable to update status")			return reconcile.Result{				RequeueAfter: time.Second,				Requeue:      true,			}, nil		}	} else {		log.Info("object is not RecocileStatusAware, not setting status")	}	return reconcile.Result{}, nil}
```

## Managing errors

## 管理错误

If a controller enters an error condition and returns an error in the reconcile method, the error will be logged by the operator to standard output and a reconciliation event will be immediately rescheduled (the default scheduler should actually detect if the same error appears over and over again, and increase the scheduling time, but in my experience, this does not occur). If the error condition is permanent, this will generate an eternal error loop situation. Furthermore, this error condition will not be visible by the user.

如果控制器进入错误条件并在 reconcile 方法中返回错误，则操作员会将错误记录到标准输出中，并且将立即重新安排协调事件（默认调度程序实际上应该检测相同的错误是否一遍又一遍地出现再次，并增加调度时间，但根据我的经验，这不会发生）。如果错误条件是永久性的，这将产生一个永恒的错误循环情况。此外，用户将看不到此错误情况。

There are two ways to notify the user of an error and they can be both used at the same time:

有两种方法可以通知用户错误，并且可以同时使用它们：

1. Return the error in the status of the object.
2. Generate an [event](https://kubernetes.io/blog/2018/01/reporting-errors-using-kubernetes-events/) describing the error.

3. 返回对象状态的错误。
4. 生成描述错误的 [event](https://kubernetes.io/blog/2018/01/reporting-errors-using-kubernetes-events/)。

Also, if you believe the error might resolve itself, you should reschedule a reconciliation cycle after a certain period of time. Often, the period of time is increased exponentially so that at every iteration the reconciliation event is scheduled farther in the future (for example twice the amount of time every time).

此外，如果您认为错误可能会自行解决，则应在一段时间后重新安排对帐周期。通常，时间段以指数方式增加，以便在每次迭代时，协调事件被安排在更远的未来（例如，每次的时间量的两倍）。

We are going to build on top of status management to handle error conditions:

我们将建立在状态管理之上来处理错误情况：

```
func (r *ReconcilerBase) ManageError(obj metav1.Object, issue error) (reconcile.Result, error) {    runtimeObj, ok := (obj).(runtime.Object)    if !ok {        log.Error(errors.New("not a runtime.Object"), "passed object was not a runtime.Object", "object", obj)        return reconcile.Result{}, nil    }    var retryInterval time.Duration    r.GetRecorder().Event(runtimeObj, "Warning", "ProcessingError", issue.Error())    if reconcileStatusAware, updateStatus := (obj).(apis.ReconcileStatusAware); updateStatus {        lastUpdate := reconcileStatusAware.GetReconcileStatus().LastUpdate.Time        lastStatus := reconcileStatusAware.GetReconcileStatus().Status        status := apis.ReconcileStatus{            LastUpdate: metav1.Now(),            Reason:     issue.Error(),            Status:     "Failure",        }        reconcileStatusAware.SetReconcileStatus(status)        err := r.GetClient().Status().Update(context.Background(), runtimeObj)        if err != nil {            log.Error(err, "unable to update status")            return reconcile.Result{                RequeueAfter: time.Second,                Requeue:      true,            }, nil        }        if lastUpdate.IsZero() || lastStatus == "Success" {            retryInterval = time.Second        } else {            retryInterval = status.LastUpdate.Sub(lastUpdate).Round(time.Second)        }    } else {        log.Info("object is not RecocileStatusAware, not setting status")        retryInterval = time.Second    }    return reconcile.Result{        RequeueAfter: time.Duration(math.Min(float64(retryInterval.Nanoseconds()*2), float64(time.Hour.Nanoseconds()*6))),        Requeue:      true,    }, nil}
```

Notice that this function immediately sends an event, then it updates the status with the error condition. Finally, a calculation is made as to when to reschedule the next attempt. The algorithm tries to double the time every loop, up to a maximum of six hours.

请注意，此函数会立即发送一个事件，然后使用错误条件更新状态。最后，计算何时重新安排下一次尝试。该算法尝试将每个循环的时间加倍，最多可达 6 小时。

Six hours is a good cap, because events last about six hours, so this should make sure that there is always an active event describing the current error condition.

六个小时是一个很好的上限，因为事件持续大约六个小时，所以这应该确保总是有一个描述当前错误情况的活动事件。

# Conclusion

## 结论

The practices presented in this blog deal with the most common concerns of designing a Kubernetes operator, and should allow you to write an operator that you feel confident putting in production. Very likely, this is just the beginning and it is easy to foresee that more frameworks and tools will come to life to help writing operators.

本博客中介绍的实践处理了设计 Kubernetes 运算符时最常见的问题，并且应该允许您编写一个您有信心投入生产的运算符。很可能，这只是一个开始，很容易预见更多的框架和工具将会出现来帮助编写运算符。

In this blog we saw many code fragments that should be immediately reusable, but for a more comprehensive example one can refer to the [operator-utils](https://github.com/redhat-cop/operator-utils) repository. 
在这篇博客中，我们看到了许多应该可以立即重用的代码片段，但更全面的示例可以参考 [operator-utils](https://github.com/redhat-c​​op/operator-utils) 存储库。