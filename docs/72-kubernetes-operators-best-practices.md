# Kubernetes Operators Best Practices
June 11, 2019 Raffaele Spazzoli

## Introduction

Kubernetes Operators are processes connecting to the master API and watching for events, typically on a limited number of resource types.

When a relevant event occurs, the operator reacts and performs a specific action. This may be limited to interacting with the master API only, but will often involve performing some action on some other systems (this could be either in cluster or off cluster resources).

![](https://assets.openshift.com/hubfs/Imported_Blog_Media/rafop1.png)

Operators are implemented as a collection of controllers where each controller watches a specific resource type. When a relevant event occurs on a watched resource a **reconcile cycle** is started.

During the reconcile cycle, the controller has the responsibility to check that current state matches the desired state described by the watched resource. Interestingly, by design, the event is not passed to the reconcile cycle, which is then forced to consider the whole state of the instance that was referenced by the event.

This approach is referred to as [level-based, as opposed to edge-based](http://venkateshabbarapu.blogspot.com/2013/03/edge-triggered-vs-level-triggered.html). Deriving from electronic circuit design, level-based triggering is the idea of receiving an event (an interrupt for example) and reacting to a state, while edge-based triggering is the idea of receiving an event and reacting to a state variation.

Level-based triggering, while arguably less efficient because it forces to re-evaluate the entire state as opposed to just what changed, is considered more suitable in complex and unreliable environments where signals can be lost or retransmitted multiple times.

This design choice influences how we write controller’s code.

Also relevant to this discussion is an understanding of the lifecycle of an API request.  The following diagram provides a high level summary:

![](https://assets.openshift.com/hubfs/Imported_Blog_Media/rafop2.png)

When a request is made to the API server, especially for create and delete requests, the request goes through the above phases. Notice that it is possible to specify webhooks to perform [mutations and validations](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#experimenting-with-admission-webhooks). If the operator introduces a new [custom resource definition](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) (CRD), we may have to also define those webhooks. Normally, the operator process would also implement the webhook endpoint by listening on a port.

This document presents a set of best practices to keep in mind when designing and developing operators using the [Operator SDK](https://github.com/operator-framework/operator-sdk).

If your operator introduces a new CRD, the Operator SDK will assist in scaffolding it. To make sure your CRD conforms to the Kubernetes best practices for extending the API, follow [these conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#api-conventions).

All the best practices mentioned in this article are portrayed in an example available at the [operator-utils](https://github.com/redhat-cop/operator-utils) repository. This repository is also a library which you can import in your operator, giving you some useful utilities for writing your own operators.

Finally this set of best practices for writing operators represents my personal view and should not be considered an official list of best practices from Red Hat.

# Creating watches

As we said, controllers watch events on resources. This is done through the abstraction of watches.

A watch is a mechanism to receive events of a certain type (either a core type or a CRD). A watch is normally created by specifying the following:

1. The resource type to watch.
2. A handler. The handler maps the events on the watched type to one or more instances for which the reconcile cycle is called. Watched type and instance type do not have to be the same.
3. A predicate. The predicate is a set of functions that can be customized to filter only the events we are interested in.

The diagram below captures these contexts:

![](https://assets.openshift.com/hubfs/Imported_Blog_Media/rafop4.png)

In general, opening multiple watches on the same kind is acceptable because the watches are multiplexed.

You should also try to filter events as much as possible. Here, for example, is a predicate that filters events on secrets. Here we are interested only in events on secrets of type _kubernetes.io/tls_ which have a certain annotation:

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

```
err = c.Watch(&amp;source.Kind{Type: &amp;examplev1alpha1.MyControlledType{}}, &amp;handler.EnqueueRequestForOwner{})
```

A less common situation is where an event is multicast to several destination resources. Consider the case of a controller that injects TLS secrets into routes based on an annotation. Multiple routes in the same namespace can point to the same secret. If the secret changes we need to update all the routes. So, we would need to create a watch on the secret type and the handler would look as follows:

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

The reconciliation cycle is where the framework gives us back control after a watch has passed up an event. As explained before, at this point we don’t have the information about the type of event because we are working on level-based triggers.

Below is a model of what a common reconciliation cycle for a controller that manages a CRD could look like. As with every model, this is not going to completely reflect any particular use case, but I hope you find it useful to help think about the problem one needs to solve when writing an operator.

![](https://assets.openshift.com/hubfs/Imported_Blog_Media/rafop3.png)

As we can see from the diagram the main steps are:

1. Retrieve the interested CR instance.
2. Manage the instance validity. We don’t want to try to do anything on an instance that does not carry valid values.
3. Manage instance initialization. If some values of the instance are not initialized, this section will take care of it.
4. Manage instance deletion. If the instance is being deleted and we need to do some specific clean up, this is where we manage it.
5. Manage controller business logic. If the above steps all pass we can finally manage and execute the reconciliation logic for this particular instance. This will be very controller specific.

In the rest of this section you can find some more in depth considerations on each of these steps.

## Resource Validation

Two types of validation exist: Syntactic validation and semantic validation.

- **Syntactic validation** happens by defining OpenAPI validation rules.
- **Semantic Validation** can be done by creating a ValidatingAdmissionConfiguration.

**Note**: it is not possible to validate a CR within a controller. Once the CR is accepted by the API server it will be stored in etcd. Once it is in etcd, the controller that owns it cannot do anything to reject it and if the CR is not valid, trying to use/process it will result in an error.

**Recommendation**: because we cannot guarantee that a ValidatingAdmissionConfiguration will be created and or working, we should also validate the CR from within the controller and if they are not valid avoid creating an endless error-loop (see also: [error management](http://cloud.redhat.com#error)).

### Syntactic Validation

OpenAPI validation rules can be added as described [here](https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html).

**Recommendation**: model as much of the custom resource as possible of your validation as syntactic validation. Syntactic validation is relatively straightforward and prevents badly formed CRs from being stored in etcd, so it should be used as much as possible.

### Semantic Validation

Semantic validation is about making sure that fields have sensible values and that the entire resource record is meaningful. Semantic validation business logic depends on the concept that the CR represents and must be coded by the operator developer.

If semantic validation is required by the given CR, then the operator should expose a webhook and [ValidatingAdmissionConfiguration](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/) should be created as part of the operator deployment.

The following limitations currently exist:

1. In OpenShift 3.11, ValidatingAdmissionConfigurations are in tech preview (they are supported from 4.1 on).
2. The Operator SDK has no support for scaffolding webhooks. This can be worked around using [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder), for example: kubebuilder webhook --group crew --version v1 --kind FirstMate --type=mutating --operations=create,update

### Validating a resource in the controller

It is better to reject an invalid CR rather than to accept it in etcd and then manage the error condition. That said, there could be situations in which the ValidatingAdmissionConfiguration is not deployed or not available at all. I think it is still a good practice to do semantic validation in the controller code. Code should be structured in such a way that you can share the same validation routine between the  ValidatingAdmissionConfiguration and the controller.

The code of the controller calling the validation method should look like this:

```
if ok, err := r.IsValid(instance); !ok {<br/>
    return r.ManageError(instance, err)<br/>
}

```

Note that if the validation fails, we manage this error as described in the [error management section](http://cloud.redhat.com#error).

The _IsValid_ function will look something like:

```
func (r *ReconcileMyCRD) IsValid(obj metav1.Object) (bool, error) {<br/>
    mycrd, ok := obj.(*examplev1alpha1.MyCRD)<br/>
// validation logic<br/>
}

```

## Resource Initialization

One of the nice conventional features of Kubernetes resources is that only the needed fields of a resource are to be initialized by the user and the others can be omitted. This is the point of view of the user, but from the point of view of the coder and anyone debugging what is happening with a resource it is actually better to have all the fields initialized.This allows writing code without always checking if a field is defined, and allows for easy troubleshooting of error situations.In order to initialize resources there are two options:

1. Define an initialization method in the controller.
2. Define a MutatingAdmissionConfiguration (the procedure issimilar to the ValidatingAdmissionConfiguration).

**Recommendation:** define an initialization method in the controller. The code should look like this sample:

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

## Resource Finalization

If resources are not owned by the CR controlled by your operator but action needs to be taken when that CR is deleted, you must use a [finalizer](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/#finalizers).

Finalizers provide a mechanism to inform the Kubernetes control plane that an action needs to take place before the standard Kubernetes garbage collection logic can be performed.

One or more finalizers can be set on resources. Each controller should manage its own finalizer and ignore others if present.

This is the pseudo code algorithm to manage finalizers:

1. If needed, add finalizers during the initialization method.
2. If the resource is being deleted, check if the finalizer owned by this controller is present.
1. If not, return
2. If yes, execute the cleanup logic
      1. If successful, update the CR by removing the finalizer.
      2. If failure decide whether to retry or give up and likely leave garbage (in some situations this can be acceptable).

If your clean-up logic requires creating additional resources, do keep in mind that additional resources cannot be created in a namespace that is being deleted. A To-be-deleted namespace will trigger a delete of all in the included resources including your CR with the finalizer.

See an example of the code here:

```
if util.IsBeingDeleted(instance) {<br/>
    if !util.HasFinalizer(instance, controllerName) {<br/>
        return reconcile.Result{}, nil<br/>
    }<br/>
    err := r.manageCleanUpLogic(instance)<br/>
    if err != nil {<br/>
        log.Error(err, "unable to delete instance", "instance", instance)<br/>
        return r.ManageError(instance, err)<br/>
    }<br/>
    util.RemoveFinalizer(instance, controllerName)<br/>
    err = r.GetClient().Update(context.TODO(), instance)<br/>
    if err != nil {<br/>
        log.Error(err, "unable to update instance", "instance", instance)<br/>
        return r.ManageError(instance, err)<br/>
    }<br/>
    return reconcile.Result{}, nil<br/>
}
```

## Resource Ownership

Resource ownership is a native concept in Kubernetes that determines how resources are deleted. When a resource is deleted and it owns other resources the children resources will be, by default, also deleted (you can disable this behavior, by setting cascade=false).

This behavior is instrumental to guarantee correct [garbage collection](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/) of resources especially when resources control other resources in a multilevel hierarchy (think deployment-> repilcaset->pod).

**Recommendation**: If your controller creates resources and these resource’s lifecycle is tied to a resource (either core or CR), then you should set this resource as the owner of those resources. This can be done as follows:

```
controllerutil.SetControllerReference(owner, obj, r.GetScheme())
```

Some other rules around ownership are the following:

1. The owner object must be in the same namespace as the owned object.
2. A namespaced resource can own a cluster level resource. We have to be careful here. An object can have a list of owners. If multiple namespaced objects own the same cluster-level object then each should claim ownership without overwriting the others’ ownership (the above API takes care of that).
3. A cluster level resource cannot own a namespaced resource.
4. A cluster level object can own another cluster level object.

## Managing status

Status is a standard section of a resource. Status is obviously used to report on the status of the resource. In this document we are going to use status to report on the outcome of the last execution of the reconcile cycle. You can use status to add more information.

Under normal circumstances, If we were updating our resource every time we execute the reconcile cycle, this would trigger an update event which in turn would trigger a reconcile cycle in an endless loop.

For this reason, **Status** should be modeled as a subresource as explained here.

This way when we can update the status of our resources without increasing the _ResourceGeneration_ metadata field. We can update the status with this command:

```
err = r.Status().Update(context.Background(), instance)
```

Now we need to write a predicate for our watch (see the section about watches for more details on these concepts) that will discard updates that did not increase the _ResourceGeneration_, this can be done using the [GenerationChangePredicate](https://github.com/operator-framework/operator-sdk/blob/master/pkg/predicate/predicate.go#L27)

If you recall, if we are using a finalizer, the finalizer should be set up at initialization time. If the finalizer is the only item that is being initialized, since it is a portion of the metadata field, the _ResourceGeneration_ will not be incremented. To account for that use case, the following  is a modified version of the predicate:

```
<span>type</span> <span>resourceGenerationOrFinalizerChangedPredicate</span> <span>struct</span> {<br/>	predicate.<span>Funcs</span><br/>}<br/><br/><span>// Update implements default UpdateEvent filter for validating resource version change</span><br/><span>func</span> (<span>resourceGenerationOrFinalizerChangedPredicate</span>) <span>Update</span>(<span>e</span> event.<span>UpdateEvent</span>) <span>bool</span> {<br/>	<span>if</span> <span>e</span>.<span>MetaNew</span>.<span>GetGeneration</span>() <span>==</span> <span>e</span>.<span>MetaOld</span>.<span>GetGeneration</span>() <span>&&</span> <span>reflect</span>.<span>DeepEqual</span>(<span>e</span>.<span>MetaNew</span>.<span>GetFinalizers</span>(), <span>e</span>.<span>MetaOld</span>.<span>GetFinalizers</span>()) {<br/>		<span>return</span> <span>false</span><br/>	}<br/>	<span>return</span> <span>true</span><br/>}
```

Now assuming your status is as follows:

```
<span>type</span> <span>MyCRStatus</span> <span>struct</span> {<br/>	<span>// +kubebuilder:validation:Enum=Success,Failure</span><br/>	<span>Status</span>     <span>string</span>      <span>`json:"status,omitempty"`</span><br/>	<span>LastUpdate</span> metav1.<span>Time</span> <span>`json:"lastUpdate,omitempty"`</span><br/>	<span>Reason</span>     <span>string</span>      <span>`json:"reason,omitempty"`</span><br/>}
```

You can write a function to manage the successful execution of a reconciliation cycle:

```
<span><br/>func</span> (<span>r</span> <span>*</span><span>ReconcilerBase</span>) <span>ManageSuccess</span>(<span>obj</span> metav1.<span>Object</span>) (reconcile.<span>Result</span>, <span>error</span>) {<br/>	<span>runtimeObj</span>, <span>ok</span> <span>:=</span> (<span>obj</span>).(runtime.<span>Object</span>)<br/>	<span>if</span> <span>!</span><span>ok</span> {<br/>		<span>log</span>.<span>Error</span>(<span>errors</span>.<span>New</span>(<span>"not a runtime.Object"</span>), <span>"passed object was not a runtime.Object"</span>, <span>"object"</span>, <span>obj</span>)<br/>		<span>return</span> reconcile.<span>Result</span>{}, <span>nil</span><br/>	}<br/>	<span>if</span> <span>reconcileStatusAware</span>, <span>updateStatus</span> <span>:=</span> (<span>obj</span>).(apis.<span>ReconcileStatusAware</span>); <span>updateStatus</span> {<br/>		<span>status</span> <span>:=</span> apis.<span>ReconcileStatus</span>{<br/>			<span>LastUpdate</span>: <span>metav1</span>.<span>Now</span>(),<br/>			<span>Reason</span>:     <span>""</span>,<br/>			<span>Status</span>:     <span>"Success"</span>,<br/>		}<br/>		<span>reconcileStatusAware</span>.<span>SetReconcileStatus</span>(<span>status</span>)<br/>		<span>err</span> <span>:=</span> <span>r</span>.<span>GetClient</span>().<span>Status</span>().<span>Update</span>(<span>context</span>.<span>Background</span>(), <span>runtimeObj</span>)<br/>		<span>if</span> <span>err</span> <span>!=</span> <span>nil</span> {<br/>			<span>log</span>.<span>Error</span>(<span>err</span>, <span>"unable to update status"</span>)<br/>			<span>return</span> reconcile.<span>Result</span>{<br/>				<span>RequeueAfter</span>: <span>time</span>.<span>Second</span>,<br/>				<span>Requeue</span>:      <span>true</span>,<br/>			}, <span>nil</span><br/>		}<br/>	} <span>else</span> {<br/>		<span>log</span>.<span>Info</span>(<span>"object is not RecocileStatusAware, not setting status"</span>)<br/>	}<br/>	<span>return</span> reconcile.<span>Result</span>{}, <span>nil</span><br/>}<br/>
```

## Managing errors

If a controller enters an error condition and returns an error in the reconcile method, the error will be logged by the operator to standard output and a reconciliation event will be immediately rescheduled (the default scheduler should actually detect if the same error appears over and over again, and increase the scheduling time, but in my experience, this does not occur). If the error condition is permanent, this will generate an eternal error loop situation. Furthermore, this error condition will not be visible by the user.

There are two ways to notify the user of an error and they can be both used at the same time:

1. Return the error in the status of the object.
2. Generate an [event](https://kubernetes.io/blog/2018/01/reporting-errors-using-kubernetes-events/) describing the error.

Also, if you believe the error might resolve itself, you should reschedule a reconciliation cycle after a certain period of time. Often, the period of time is increased exponentially so that at every iteration the reconciliation event is scheduled farther in the future (for example twice the amount of time every time).

We are going to build on top of status management to handle error conditions:

```
func (r *ReconcilerBase) ManageError(obj metav1.Object, issue error) (reconcile.Result, error) {<br/>
    runtimeObj, ok := (obj).(runtime.Object)<br/>
    if !ok {<br/>
        log.Error(errors.New("not a runtime.Object"), "passed object was not a runtime.Object", "object", obj)<br/>
        return reconcile.Result{}, nil<br/>
    }<br/>
    var retryInterval time.Duration<br/>
    r.GetRecorder().Event(runtimeObj, "Warning", "ProcessingError", issue.Error())<br/>
    if reconcileStatusAware, updateStatus := (obj).(apis.ReconcileStatusAware); updateStatus {<br/>
        lastUpdate := reconcileStatusAware.GetReconcileStatus().LastUpdate.Time<br/>
        lastStatus := reconcileStatusAware.GetReconcileStatus().Status<br/>
        status := apis.ReconcileStatus{<br/>
            LastUpdate: metav1.Now(),<br/>
            Reason:     issue.Error(),<br/>
            Status:     "Failure",<br/>
        }<br/>
        reconcileStatusAware.SetReconcileStatus(status)<br/>
        err := r.GetClient().Status().Update(context.Background(), runtimeObj)<br/>
        if err != nil {<br/>
            log.Error(err, "unable to update status")<br/>
            return reconcile.Result{<br/>
                RequeueAfter: time.Second,<br/>
                Requeue:      true,<br/>
            }, nil<br/>
        }<br/>
        if lastUpdate.IsZero() || lastStatus == "Success" {<br/>
            retryInterval = time.Second<br/>
        } else {<br/>
            retryInterval = status.LastUpdate.Sub(lastUpdate).Round(time.Second)<br/>
        }<br/>
    } else {<br/>
        log.Info("object is not RecocileStatusAware, not setting status")<br/>
        retryInterval = time.Second<br/>
    }<br/>
    return reconcile.Result{<br/>
        RequeueAfter: time.Duration(math.Min(float64(retryInterval.Nanoseconds()*2), float64(time.Hour.Nanoseconds()*6))),<br/>
        Requeue:      true,<br/>
    }, nil<br/>
}

```

Notice that this function immediately sends an event, then it updates the status with the error condition. Finally, a calculation is made as to when to reschedule the next attempt. The algorithm tries to double the time every loop, up to a maximum of six hours.

Six hours is a good cap, because events last about six hours, so this should make sure that there is always an active event describing the current error condition.

# Conclusion

The practices presented in this blog deal with the most common concerns of designing a Kubernetes operator, and should allow you to write an operator that you feel confident putting in production. Very likely, this is just the beginning and it is easy to foresee that more frameworks and tools will come to life to help writing operators.

In this blog we saw many code fragments that should be immediately reusable, but for a more comprehensive example one can refer to the [operator-utils](https://github.com/redhat-cop/operator-utils) repository.