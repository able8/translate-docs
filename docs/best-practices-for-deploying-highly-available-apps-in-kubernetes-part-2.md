[Blog](https://blog.flant.com/)

- [Home](https://blog.flant.com/)
- [DevOps](https://blog.flant.com/category/devops/)
- [CI/CD](https://blog.flant.com/category/devops/ci-cd/)
- [Releases](https://blog.flant.com/category/releases/)
- [flant.com](https://flant.com)

DevOps as a Service from FlantDedicated 24x7x365 support [Discover the benefits](https://flant.com/services/devops-as-a-service?utm_source=blog.flant.com&utm_medium=banner)

![](https://blog.flant.com/wp-content/uploads/sites/2/2021/08/flant-kubernetes-deploy-best-practices.png)

16 August 2021

[Sergey Sizov](https://github.com/SpecialForce3331), software engineer

- [DevOps](https://blog.flant.com/category/devops/)
- [#best practices](https://blog.flant.com/tag/best-practices/)
- [#Kubernetes](https://blog.flant.com/tag/kubernetes/)

# Best practices for deploying highly available apps in Kubernetes. Part 2

In [Part One](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/), we shared some recommendations for a number of Kubernetes mechanisms that facilitate the deployment of highly available applications. We discussed aspects of the scheduler operation, update strategies, priorities, probes, etc. This second part discusses the remaining three crucial topics: PodDisruptionBudget, HorizontalPodAutoscaler, and VerticalPodAutoscaler (the numbering continues from [Part One](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/)).

## 9\. PodDisruptionBudget

The [PodDisruptionBudget](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#pod-disruption-budgets) (PDB) mechanism is a must-have for applications running in production. It provides you the means to specify a maximum limit to the number of application Pods that can be unavailable simultaneously. In [Part One](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/), we discussed some methods that are instrumental in avoiding potentially risky situations: running several application replicas, specifying `podAntiAffinity` (to prevent several Pods from being assigned to the same node), etc.

However, you may encounter a situation in which more than one K8s node becomes unavailable at the same time. For example, suppose you decide to switch instances to more powerful ones. There may be other reasons besides that, but that’s beyond the scope of this article. The thing is that several nodes get removed at the same time. “But that’s Kubernetes!” you might say. ”Everything is ephemeral here! The Pods will get moved to other nodes, so what’s the deal?” Well, let’s have a look.

Suppose the application has three replicas. The load is evenly distributed between them, while the Pods are distributed across the nodes. In this case, the application will continue to run even if one of the replicas fails. However, the failure of two replicas will result in a service degradation: one single Pod simply cannot handle the entire load on its own. The clients will start getting 5XX errors. (Of course, you can set a rate limit in the nginx container; in that case, the error will be _429 Too Many Requests_. Still, the service will degrade nevertheless).

And that’s where PodDisruptionBudget comes to help. Let’s take a look at its manifest:

```
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
name: app-pdb
spec:
maxUnavailable: 1
selector:
    matchLabels:
      app: app
```

The manifest is pretty straightforward; you are probably familiar with most of its fields, `maxUnavailable` being the most interesting among them. This field sets the maximum number of Pods that can be simultaneously unavailable. This can be either an absolute number or a percentage.

Suppose the PDB is configured for the application. What will happen in the event that, for some reason, two or more nodes start evicting application Pods? The above PDB only allows for one Pod to be evicted at a time. Therefore, the second node will wait until the number of replicas reverts to the pre-eviction level, and only then will the second replica be evicted.

As an alternative, you can also set a `minAvailable` parameter. For example:

```
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
name: app-pdb
spec:
minAvailable: 80%
selector:
    matchLabels:
      app: app
```

This parameter ensures that at least 80% of replicas are available in the cluster at all times. Thus, only 20% of replicas can be evicted if necessary. `minAvailable` can be either an absolute number or a percentage.

But there is a catch: there have to be enough nodes in the cluster that satisfy the `podAntiAffinity` criteria. Otherwise, you may encounter a situation in which a replica gets evicted, but the scheduler cannot re-deploy it due to a lack of suitable nodes. As a result, draining a node will take forever to complete, and will get you two application replicas instead of three. Granted, you can invoke `kubectl describe` for a _Pending_ Pod to see what’s going on and eliminate the problem. But still, it is better to prevent this kind of situation from happening.

To summarize, **always configure the PDB for critical components of your system**.

## 10\. HorizontalPodAutoscaler

Let’s consider another situation: what happens if an application has an unexpected load that is significantly higher than usual? Yes, you can scale up the cluster manually, but that is not the method we use.

That is where [HorizontalPodAutoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) (HPA) comes in. With HPA, you can choose a metric and use it as a trigger for scaling the cluster up/down automatically, depending on the metric’s value. Imagine that on one quiet night your cluster suddenly gets blasted with a massive uptick in traffic, say, Reddit users have found out about your service. The CPU load (or some other Pod metric) increases, hits the threshold, and then HPA comes into play. It scales up the cluster, thus distributing the load between a larger number of Pods.

Thanks to that, all the incoming requests are processed successfully. Just as important, after the load returns to the average level, **HPA scales the cluster down to reduce infrastructure costs and save money**. Sounds great, doesn’t it?

Let’s see how exactly HPA calculates the number of replicas to be added. Here is the formula from the documentation:

`desiredReplicas = ceil[currentReplicas * ( currentMetricValue / desiredMetricValue )]` ``

Now suppose that:

- the current number of replicas is 3;
- the current metric value is 100;
- the metric threshold value is 60;

In this case, the resulting number is `3 * ( 100 / 60 )`, i.e. “about” 5 replicas (HPA will round the result up). Thus, the application will gain two more replicas. But that is not the end of the story: HPA will continue to calculate the number of replicas required (using the formula above) to scale down the cluster if the load decreases.

And that brings us to the most exciting part. What metric should you use? The first thing that comes to mind is one of the primary metrics, such as CPU or Memory utilization. And that will work if your CPU and Memory consumption is directly proportional to the incoming load. But what if the Pods are handling different requests? Some requests require many CPU cycles, others may consume a lot of memory, and still others only demand minimum resources.

Let’s take a look, for example, at the RabbitMQ queue and the instances processing it. Suppose there are ten messages in the queue. Monitoring shows that messages are being dequeued (as per RabbitMQ’s terminology) steadily and regularly. That is, we feel that ten messages in the queue on average is okay. But then the load suddenly increases, and the queue grows to 100 messages. However, the workers’ CPU and Memory consumption stays the same: they are steadily processing the queue, leaving about 80-90 messages in it.

But **what if we use a custom metric** that describes the number of messages in the queue? Let’s configure our custom metric as follows:

- the current number of replicas is 3;
- the current metric value is 80;
- the metric threshold value is 15.

Thus, `3 * ( 80 / 15 ) = 16`. In this case, HPA can increase the number of workers to 16, and they quickly process all the messages in the queue (at which point HPA will decrease their number again). However, all the required infrastructure must be ready to accommodate this number of Pods. That is, they must fit on the existing nodes, or new nodes must be provisioned by the infrastructure provider (cloud provider) in the case that [Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) is used. In other words, we are back to planning cluster resources.

Now let’s take a look at some manifests:

```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
name: php-apache
spec:
scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: php-apache
minReplicas: 1
maxReplicas: 10
targetCPUUtilizationPercentage: 50
```

This one is simple. As soon as the CPU load reaches 50%, HPA starts scaling the number of replicas to a maximum of 10.

Here is a more interesting one:

```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
name: worker
spec:
scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: worker
minReplicas: 1
maxReplicas: 10
metrics:
  - type: External
    external:
      metric:
        name: queue_messages
      target:
        type: AverageValue
        averageValue: 15
```

Note that in this example, HPA uses the [custom metric](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-custom-metrics). It will base its scaling decisions on the size of the queue ( `queue_messages` metric). Given that the average number of messages in the queue is 10, we set the threshold to 15. This way, you can manage the number of replicas more accurately. As you can see, the custom metric enables more accurate cluster autoscaling than, say, a CPU-based metric.

### Additional features

The HPA configuration options are pretty diverse. For example, you can combine different metrics. In the manifest below, CPU utilization and queue size are used to trigger scaling decisions.

```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
name: worker
spec:
scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: worker
minReplicas: 1
maxReplicas: 10
metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 50
  - type: External
    external:
      metric:
        name: queue_messages
      target:
        type: AverageValue
        averageValue: 15
```

What calculation algorithm does HPA apply? Well, it uses the highest calculated number of replicas regardless of the metric exploited. For example, if the calculation based on the CPU metric shows that 5 replicas need to be added while the queue size-based metric gives only 3 Pods, HPA will use the larger value and add 5 Pods.

With the release of Kubernetes 1.18, you now [have the ability](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-configurable-scaling-behavior) to define `scaleUp` and `scaleDown` policies. For example:

```
behavior:
scaleDown:
    stabilizationWindowSeconds: 60
    policies:
    - type: Percent
      value: 5
      periodSeconds: 20
   - type: Pods
      value: 5
      periodSeconds: 60
    selectPolicy: Min
scaleUp:
    stabilizationWindowSeconds: 0
    policies:
    - type: Percent
      value: 100
      periodSeconds: 10
```

As you can see in the manifest above, it features two sections. The first one ( `scaleDown`) defines the scaling down parameters while the second ( `scaleUp`) is used for scaling up. Each section features the `stabilizationWindowSeconds`. This helps prevent what is referred to as “flapping” (or unnecessary scaling) as the number of replicas continues to oscillate. This parameter essentially serves as a timeout after the number of replicas is changed.

Now let’s talk about the policies. The `scaleDown` policy allows you to specify the percent of Pods ( `type: Percent`) to scale down over a specific period of time. If the load features a cyclical pattern, what you have to do is decrease the percentage and increase the duration period. In that case, as the load decreases, HPA will not kill a large number of Pods at once (according to its formula) but will do so gradually instead. Furthermore, you can set the maximum number of Pods ( `type: Pods`) that HPA is allowed to kill over the specified time period.

Note the `selectPolicy: Min` parameter. What that means is HPA uses the policy that affects the minimum number of Pods. Thus, HPA will choose a percent value if it (5% in the example above) is less than the numeric alternative (5 Pods in the example above). Conversely, the `selectPolicy: Max` policy will have the opposite effect.

Similar parameters are used in the `scaleUp` section. Note that in most situations, the cluster must scale up (almost) instantly since even a slight delay can affect users and their experience. For that reason, `stabilizationWindowSeconds` is set to 0 in this section. If the load has a cyclical pattern, HPA can increase the replica count to `maxReplicas` (as defined in the HPA manifest) if necessary. Our policy allows HPA to add up to 100% to the currently running replicas every 10 seconds ( `periodSeconds: 10`).

Finally, you can set the `selectPolicy` parameter to `Disabled` to turn off scaling in the given direction:

```
behavior:
scaleDown:
    selectPolicy: Disabled
```

Most of the time that policies are used is when HPA does not work as expected. **Policies provide flexibility but render the manifest harder to grasp.**

Recently, HPA [became capable](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#container-resource-metrics) to track the resource usage of individual containers across a set of Pods (introduced as an alpha feature in Kubernetes 1.20).

### HPA: summary

Let us conclude this section with an example of the complete HPA manifest:

```
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
name: worker
spec:
scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: worker
minReplicas: 1
maxReplicas: 10
metrics:
  - type: External
    external:
      metric:
        name: queue_messages
      target:
        type: AverageValue
        averageValue: 15
behavior:
    scaleDown:
      stabilizationWindowSeconds: 60
      policies:
      - type: Percent
        value: 5
        periodSeconds: 20
     - type: Pods
        value: 5
        periodSeconds: 60
      selectPolicy: Min
    scaleUp:
      stabilizationWindowSeconds: 0
      policies:
      - type: Percent
        value: 100
        periodSeconds: 10
```

Please note this example is provided for informational purposes only. You will need to adapt it to suit the specifics of your own operation.

Horizontal Pod Autoscaler resume: **HPA is perfect for production environments. But you have to be careful and forward-thinking when choosing metrics for HPA.** A mistaken metric or an incorrect threshold will result in either a waste of resources (from unnecessary replicas) or service degradation (if the number of replicas is not enough). Closely monitor the behavior of the application and test it until you’ve achieved the right balance.

## 11\. VerticalPodAutoscaler

[VPA](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) analyzes the resource requirements of the containers and sets (if the corresponding mode is enabled) their limits and requests.

Suppose you have deployed a new app version with some new functions and it turns out that, say, the imported library is a huge resource eater, or the code isn’t very well optimized. In other words, the application resource requirements have increased. You failed to notice this during testing (since it is hard to load the application in the same way as in production).

And, of course, the relevant requests and limits had been set for the app before an update begins. And now the application reaches the memory limit, and its Pod gets killed due to OOM. VPA can prevent this! At first glance, **VPA looks like a great tool that should be used whenever and wherever possible. But in real life that isn’t always necessarily the case**, and you have to bear in mind the finer details involved.

The main problem (it isn’t solved yet) is that the Pod needs to be restarted for resource changes to take effect. In the future, VPA will modify them without restarting the Pod, but for now, it simply isn’t capable of doing that. But no need to worry. That isn’t a big deal if you have a “well-written” application that is always ready for redeployment (say, it has a large number of replicas; its PodAntiAffinity, PodDistruptionBudget, HorizontalPodAutoscaler are carefully configured; etc.). In that case, you (probably) won’t even notice the VPA activity.

Sadly, there are other less pleasant scenarios that may occur like: the application not taking redeployment very well, the number of replicas being limited due to a lack of nodes, our application running as a StatefulSet, etc. In the worst-case scenario, the Pods’ resource consumption grows due to an increased load, HPA starts to scale up the cluster, and then, suddenly, VPA proceeds to modify the resource parameters and restarts the Pods. As a result, this high load gets distributed across the rest of the Pods. Some of them may crash, rendering things even worse and resulting in a chain reaction of failure.

That is why having a profound understanding of various VPA operating modes is important. Let’s start with the simplest one — “Off”.

**Off mode**

All this mode does is calculate the resource consumption of Pods and make recommendations. Looking ahead, I would like to note that at Flant **we use this mode in the majority of cases** (and we recommend it). But first, let’s look at a few examples.

Some basic manifests follow below:

```
apiVersion: autoscaling.k8s.io/v1
kind: VerticalPodAutoscaler
metadata:
name: my-app-vpa
spec:
targetRef:
    apiVersion: "apps/v1"
    kind: Deployment
    name: my-app
updatePolicy:
    updateMode: "Recreate"
    containerPolicies:
      - containerName: "*"
        minAllowed:
          cpu: 100m
          memory: 250Mi
        maxAllowed:
          cpu: 1
          memory: 500Mi
        controlledResources: ["cpu", "memory"]
        controlledValues: RequestsAndLimits
```

We will not go into detail about this manifest’s parameters: [this article](https://povilasv.me/vertical-pod-autoscaling-the-definitive-guide/) provides a detailed description of the features and aspects of VPA. In short, we specify the VPA target ( `targetRef`) and select the update policy. Additionally, we specify the upper and lower limits for the resources VPA can use. The primary focus is on the `updateMode` field. In “Recreate” or “Auto” mode, VPA will recreate Pods with all consequences (until an above-mentioned patch for in-place Pod resource parameters update becomes available). Since we don’t want it, we use the “Off” mode:

```
apiVersion: autoscaling.k8s.io/v1
kind: VerticalPodAutoscaler
metadata:
name: my-app-vpa
spec:
targetRef:
    apiVersion: "apps/v1"
    kind: Deployment
    name: my-app
updatePolicy:
    updateMode: "Off"   # !!!
resourcePolicy:
    containerPolicies:
      - containerName: "*"
        controlledResources: ["cpu", "memory"]
```

VPA starts collecting metrics. You can use the `kubectl describe vpa` command to see the recommendations (just let VPA run for a few minutes):

```
Recommendation:
    Container Recommendations:
      Container Name:  nginx
      Lower Bound:
        Cpu:     25m
        Memory:  52428800
      Target:
        Cpu:     25m
        Memory:  52428800
      Uncapped Target:
        Cpu:     25m
        Memory:  52428800
      Upper Bound:
        Cpu:     25m
        Memory:  52428800
```

The VPA recommendations will be more accurate after a couple of days (a week, a month, etc.) of running. And then is the perfect time to adjust limits in the application manifest. That way, you can avoid OOM kills due to a lack of resources and save on infrastructure (if initial requests/limits are too high).

Now, let’s talk about some of the details of using VPA.

### Other VPA modes

Note that in “Initial” Mode, VPA assigns resources when Pods are started and never changes them later. Thus, VPA will set low requests/limits for newly created Pods if the load was relatively low over the past week. It may lead to problems if the load suddenly increases because the requests/limits will be much lower than what is required for such a load. This mode may come in handy if your load is uniformly distributed and grows in a linear fashion.

In “Auto” mode, VPA recreates the Pods. Thus, the application must handle the restart properly. If it cannot shutdown gracefully (i.e. by closing the existing connections correctly and so on), you will most likely catch some avoidable 5XX errors. Using Auto mode with a StatefulSet is rarely advisable: imagine VPA attempting to add PostgreSQL resources to production…

As for the dev environment, you can freely experiment to find the level of resources to use (later) in production that is acceptable to you. Suppose you want to use VPA in the “Initial” mode and we have Redis in the cluster using the `maxmemory` parameter. You will most likely need to change it to adjust it to your needs. The problem is Redis doesn’t care about the limits at the cgroups level. In other words, you are risking a lot if `maxmemory` is, say, 2GB while your Pod’s memory is capped at 1GB. But how can you set `maxmemory` to be the same as the limit? Well, there is a way! You can use the VPA-recommended values:

```
apiVersion: apps/v1
kind: Deployment
metadata:
name: redis
labels:
    app: redis
spec:
replicas: 1
selector:
    matchLabels:
      app: redis
template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
      - name: redis
        image: redis:6.2.1
        ports:
        - containerPort: 6379
        resources:
           requests:
             memory: "100Mi"
             cpu: "256m"
           limits:
             memory: "100Mi"
             cpu: "256m"
        env:
          - name: MY_MEM_REQUEST
            valueFrom:
              resourceFieldRef:
                containerName: app
                resource: requests.memory
          - name: MY_MEM_LIMIT
            valueFrom:
              resourceFieldRef:
                containerName: app
                resource: limits.memory
```

You can use environment variables to obtain the memory limit (and subtract, say, 10% from that for application needs) and set the resulting value as `maxmemory`. You will probably have to do something about the init container that uses `sed` to process the Redis config since the default Redis container image does not support passing `maxmemory` using an environment variable. Nevertheless, this solution is quite functional.

Finally, I would like to turn your attention to the fact that VPA evicts the DaemonSet Pods all at once, en masse. We are currently working on a [patch](https://github.com/kubernetes/kubernetes/pull/98307) that fixes this.

### Final VPA recommendations

“Off” mode is suitable for the majority of cases.

You can experiment with “Auto” and “Initial” modes in the dev environment.

**Only use VPA in production if you have already accumulated recommendations and tested them thoroughly**. In addition, you have to clearly understand what you are doing and why you are doing it.

In the meantime, we are eagerly anticipating in-place (restart-free) updates for Pod resources.

Note that there are some limitations associated with joint use of HPA and VPA. For instance, VPA should not be used together with HPA if the CPU- or Memory-based metric is used as a trigger. The reason is that when the threshold is reached, VPA increases resource requests/limits while HPA adds new replicas. Consequently, the load will drop off sharply, and the process will go in reverse, resulting in “flapping”. The [official documentation](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#known-limitations) sheds more light on the existing limitations.

## Conclusion

This concludes our review of best practices for deploying HA applications in Kubernetes. Please share your thoughts and suggestions!

## Related posts:

- [Migrating your app to Kubernetes: what to do with files?](https://blog.flant.com/migrating-your-app-to-kubernetes-what-to-do-with-files/ "Migrating your app to Kubernetes: what to do with files?")
- [ConfigMaps in Kubernetes: how they work and what you should remember](https://blog.flant.com/configmaps-in-kubernetes-how-they-work-and-what-you-should-remember/ "ConfigMaps in Kubernetes: how they work and what you should remember")
- [Best practices for deploying highly available apps in Kubernetes. Part 1](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1/ "Best practices for deploying highly available apps in Kubernetes. Part 1")

## Share

© 2008–2022 [Flant Europe OÜ](http://flant.com)

- [Twitter](https://twitter.com/flant_com)
- [Facebook](https://www.facebook.com/flantcom)
- [LinkedIn](https://www.linkedin.com/company/flant)
- [GitHub](http://github.com/flant)
- [YouTube](https://www.youtube.com/c/Flant_com)

Subscribe me to your new tech posts:

✕



https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-2
