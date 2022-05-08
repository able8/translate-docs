[Blog](https://blog.flant.com/)

- [Home](https://blog.flant.com/)
- [DevOps](https://blog.flant.com/category/devops/)
- [CI/CD](https://blog.flant.com/category/devops/ci-cd/)
- [Releases](https://blog.flant.com/category/releases/)
- [flant.com](https://flant.com)

DevOps as a Service from FlantDedicated 24x7x365 support [Discover the benefits](https://flant.com/services/devops-as-a-service?utm_source=blog.flant.com&utm_medium=banner)

![](https://blog.flant.com/wp-content/uploads/sites/2/2021/08/flant-kubernetes-deploy-best-practices.png)

6 August 2021

[Ilya Lesikov](https://github.com/ilya-lesikov), software engineer

- [DevOps](https://blog.flant.com/category/devops/)
- [#best practices](https://blog.flant.com/tag/best-practices/)
- [#favourites](https://blog.flant.com/tag/favourites/)
- [#Kubernetes](https://blog.flant.com/tag/kubernetes/)

# Best practices for deploying highly available apps in Kubernetes. Part 1

As you know, deploying a basic viable app configuration in Kubernetes is a breeze. On the other hand, trying to make your application as available and fault-tolerant as possible inevitably entails a great number of hurdles and pitfalls. In this article, we break down what we believe to be the most important rules when it comes to deploying high-availability applications in Kubernetes and sharing them in a concise way.

Note that we will not be using any features that aren’t available right out-of-the-box. What we will also not do is lock into specific CD solutions, and we will omit the issues of templating/generating Kubernetes manifests. In this article, we only discuss the final structure of Kubernetes manifests when deploying to the cluster.

## 1\. Number of replicas

You need at least two replicas for the application to be considered minimally available. But why, you may ask, is a single replica not enough? The problem is that many entities in Kubernetes (Node, Pod, ReplicaSet, etc.) are ephemeral, i.e. under certain conditions, they may be automatically deleted/recreated. Obviously, the Kubernetes cluster and applications running in it must account for that.

For example, when the autoscaler scales down your number of nodes, some of those nodes will be deleted, including the Pods running on them. If the sole instance of your application is running on one of the nodes to be deleted, you may find your application completely unavailable, though this is usually short-lived. In general, if you only have one replica of the application, any abnormal termination of it will result in downtime. In other words, you must have **at least two running replicas of the application.**

The more replicas there are, the milder of a decline there will be in your application’s computing capacity in the event that some replica fails. For example, suppose you have two replicas and one fails due to network issues on a node. The load that the application can handle will be cut in half (with only one of the two replicas available). Of course, the new replica will be scheduled on a new node, and the load capacity of the application will be fully restored. But until then, increasing the load can lead to service disruptions, which is why you **must have some replicas in reserve.**

_The above recommendations are relevant to cases in which there is no HorizontalPodAutoscaler used. The best alternative for applications that have more than a few replicas is to configure HorizontalPodAutoscaler and let it manage the number of replicas. We will focus on HorizontalPodAutoscaler in the next article._

## 2\. The update strategy

The default update strategy for Deployment entails a reduction of the number of old+new ReplicaSet Pods with a `Ready` status of 75% of their pre-update amount. Thus, during the update, the computing capacity of an application may drop to 75% of its regular level, and that may lead to a partial failure (degradation of the application’s performance). The `strategy.RollingUpdate.maxUnavailable` parameter allows you to configure the maximum percentage of Pods that can become unavailable during an update. Therefore, either make sure that your application runs smoothly even in the event that 25% of your Pods are unavailable or lower the `maxUnavailable` parameter. Note that the `maxUnavailable` parameter is rounded down.

There’s a little trick to the default update strategy ( `RollingUpdate`): the application will temporarily have not only a few replicas, but two different versions (the old one and the new one) running concurrently as well. Therefore, if running different replicas and different versions of the application side by side is unfeasible for some reason, then you can use `strategy.type: Recreate`. Under the `Recreate` strategy, all the existing Pods are killed before the new Pods are created. This results in a short-lived downtime.

_Other deployment strategies (blue-green, canary, etc.) can often provide a much better alternative to the RollingUpdate strategy. However, we are not taking them into account in this article since their implementation depends on the software used to deploy the application. That goes beyond the scope of this article (here is a_ [_great article_](https://www.weave.works/blog/kubernetes-deployment-strategies) _on the topic that we recommend and is well worth the read)._

## 3\. Uniform replicas distribution across nodes

It is very important that you distribute Pods of the application across different nodes if you have multiple replicas of the application. To do so, **you can instruct your scheduler to avoid starting multiple Pods of the same Deployment on the same node:**

```
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchLabels:
                  app: testapp
              topologyKey: kubernetes.io/hostname
```

It is better to use `preferredDuringSchedulingaffinity` instead of `requiredDuringScheduling`. The latter may render it impossible to start new Pods if the number of nodes required for the new Pods is larger than the number of nodes available. Still, the `requiredDuringScheduling` affinity might come in handy when the number of nodes and application replicas is known in advance and you need to be sure that two Pods will not end up on the same node.

## 4\. Priority

`priorityClassName` represents your Pod priority. The scheduler uses it to decide which Pods are to be scheduled first and which Pods should be evicted first if there is no space for Pods left on the nodes.

You will need to add several [PriorityClass](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass) type resources and map Pods to them using `priorityClassName`. Here is an example of how `PriorityClasses` may vary:

- _Cluster_. Priority > 10000. Cluster-critical components, such as kube-apiserver.
- _Daemonsets_. Priority: 10000. Usually, it is not advised for DaemonSet Pods to be evicted from cluster nodes and replaced by ordinary applications.
- _Production-high_. Priority: 9000. Stateful applications.
- _Production-medium._ Priority: 8000. Stateless applications.
- _Production-low_. Priority: 7000. Less critical applications.
- _Default_. Priority: 0. Non-production applications.

Setting priorities will help you to avoid sudden evictions of critical components. Also, critical applications will evict less important applications if there is a lack of node resources.

## 5\. Stopping processes in containers

The signal specified in `STOPSIGNAL` (usually, the `TERM` signal) is sent to the process to stop it. However, some applications cannot handle it properly and cannot manage to shut down gracefully. The same is true for applications running in Kubernetes.

For example, in order to shut down nginx properly, you will need a `preStop` hook like this:

```
lifecycle:
preStop:
    exec:
      command:
      - /bin/sh
      - -ec
      - |
        sleep 3
        nginx -s quit
```

A brief explanation for this listing:

1. `sleep 3` prevents race conditions that may be caused by an endpoint being deleted.
2. `nginx -s quit` shuts down nginx properly. This line isn’t required for more up-to-date images since the `STOPSIGNAL: SIGQUIT` parameter is set there by default.

_(You can learn more about graceful shutdowns for nginx bundled with PHP-FPM in_ [_our other article_](https://blog.flant.com/graceful-shutdown-in-kubernetes-is-not-always-trivial/) _.)_

The way `STOPSIGNAL` is handled depends on the application itself. In practice, for most applications, you have to Google the way `STOPSIGNAL` is handled. If the signal is not handled appropriately the `preStop` hook can help you solve the problem. Another option is to replace `STOPSIGNAL` with a signal that the application can handle properly (and permit it to shut down gracefully).

`terminationGracePeriodSeconds` is another crucial parameter important in shutting down the application. It specifies the time period for which the application is to shut down gracefully. If the application does not terminate within this time frame (30 seconds by default), it will receive a `KILL` signal. Thus, you will need to increase the terminationGracePeriodSeconds parameter if you think that running the `preStop` hook and/or shutting down the application at the `STOPSIGNAL` may take more than 30 seconds. For example, you may need to increase it if some requests from your web service clients take a long time to complete (e.g. requests that involve downloading large files).

It is worth noting that the `preStop` hook has a locking mechanism, i.e. `STOPSIGNAL` may be sent only after the `preStop` hook has finished running. At the same time, the `terminationGracePeriodSeconds` countdown _continues during the_ _preStop_ _hook execution_. All the hook-induced processes, as well as the processes running in the container, will be `KILL` ed after `terminationGracePeriodSeconds` is over.

Also, some applications have specific settings that set the deadline at which point the application must terminate (for example, the `--timeout` option in Sidekiq). Therefore, in each case, you have to make sure that if the application has this setting, it has a slightly lower value than that of `terminationGracePeriodSeconds`.

## 6\. Reserving resources

The scheduler uses a Pod’s `resources.requests` to decide which node to place the Pod on. For instance, a Pod cannot be scheduled on a Node that does not have enough free (i.e., _non-requested_) resources to cover that Pod’s resource requests. On the other hand, `resources.limits` allow you to limit Pods’ resource consumption that heavily exceeds their respective requests. A good tip is to **set limits equal to requests**. Setting limits at much higher than requests may lead to a situation when some of a node’s Pods not getting the requested resources. This may lead to the failure of other applications on the node (or even the node itself). Kubernetes assigns a [QoS class](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod) to each Pod based on its resource scheme. K8s then uses QoS classes to make decisions about which Pods should be evicted from the nodes.

Therefore, you **have to set both requests and limits for both the CPU and memory.** The only thing you can/should omit is the CPU limit if the [Linux kernel version is older than 5.4](https://engineering.indeedblog.com/blog/2019/12/cpu-throttling-regression-fix/) (in the case of EL7/CentOS7, the kernel version must be older than 3.10.0-1062.8.1.el7).

Furthermore, the memory consumption of some applications tends to grow in an unlimited fashion. A good example of that is Redis used for caching or an application that basically runs “on its own”. To limit their impact on other applications on the node, you can (and should) set limits for the amount of memory to be consumed. The only problem with that is the application will be `KILL` ed when this limit is reached. Applications cannot predict/handle this signal, and this will probably prevent them from shutting down correctly. That is why, in addition to Kubernetes limits, we **highly recommend using application-specific mechanisms for limiting memory consumption** so that it does not exceed (or come close to) the amount set in a Pod’s `limits.memory` parameter.

Here is a Redis configuration that can help you with this:

```
maxmemory 500mb   # if the amount of data exceeds 500 MB...
maxmemory-policy allkeys-lru   # ...Redis would delete rarely used keys
```

As for Sidekiq, you can use the [Sidekiq worker killer](https://github.com/klaxit/sidekiq-worker-killer):

```
require 'sidekiq/worker_killer'
Sidekiq.configure_server do |config|
config.server_middleware do |chain|
    # Terminate Sidekiq correctly when it consumes 500 MB
    chain.add Sidekiq::WorkerKiller, max_rss: 500
end
end
```

It is clear that in all these cases that `limits.memory` needs to be higher than the thresholds for triggering the above mechanisms.

_In the next article, we’ll discuss using VerticalPodAutoscaler to allocate resources automatically._

## 7\. Probes

In Kubernetes, probes (health checks) are used to determine whether it is possible to switch traffic to the application ( _readiness_ probes) and whether the application needs to be restarted ( _liveness_ probes). They play an important role in updating Deployments and starting new Pods in general.

First of all, we would like to provide a general recommendation for all probe types: **set a high value for the** **`timeoutSeconds` parameter**. A default value of one second is way too low, and it will have a critical impact on readinessProbe & livenessProbe. If `timeoutSeconds` is too low, an increase in the application response time (which often takes place simultaneously for all Pods due to Service load balancing) may either result in these Pods being removed from load balancing (in the case of a failed readiness probe) or, what’s worse, in cascading container restarts (in the case of a failed liveness probe).

### 7.1. Liveness probe

In practice, the liveness probe is not as widely used as you may have thought. Its purpose is to restart a container if, for example, the application is frozen. However, in real life, such app deadlocks are an exception rather than the rule. If the application demonstrates partial functionality for some reason (e.g., it cannot restore connection to a database after it has been broken), you have to fix that in the application, rather than “inventing” livenessProbe-based workarounds.

While you can use livenessProbe to check for these kinds of states, we recommend either **not using livenessProbe by default** or only performing some **basic liveness-testing, such as testing for the TCP connection (remember to set a high timeout value)**. This way, the application will be restarted in response to an apparent deadlock without risking falling into the trap of a loop of restarts (i.e. restarting it won’t help).

Risks related to a poorly configured livenessProbe are serious. In the most common cases, livenessProbe fails due to increased load on the application (it simply cannot make it within the time specified by the timeout parameter) or due to the state of external dependencies that are currently down being checked (directly or indirectly). In the latter case, all the containers will be restarted. In the best case scenario, this would result in nothing, but in the worst case, this would render the application  completely unavailable, probably long-term. Long-term total unavailability of an application (if it has a large number of replicas) may result if most Pods’ containers are restarted within a short time period. Some containers are likely to become READY faster than others, and the entire load will be distributed over this limited number of running containers. That will end up causing livenessProbe timeouts, which will trigger even more restarts.

Also, ensure that livenessProbe does not stop responding if your application has a limit on the number of established connections and that limit has been reached. Usually, you have to dedicate a separate application thread/process to livenessProbe to avoid such problems. For example, if your application has 11 threads (one thread per client), you can limit the number of clients to 10, ensuring that there is an idle thread available for livenessProbe.

And, of course, do not add any external dependency checks to your livenessProbe.

See [this article](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html) for more information on liveness probe issues and how to prevent them.

### 7.2. Readiness probe

The design of readinessProbe has turned out not to be very successful. readinessProbe combines two different functions:

- it finds out if an application is available during the container start;
- it checks if an application remains available after the container has been successfully started.

In practice, the first function is required in the vast majority of cases, while the second is only needed as often as the livenessProbe. The poorly configured readinessProbe can cause issues similar to those of livenessProbe. In the worst case, they can also end up causing long-term unavailability for the application.

When readinessProbe fails, the Pod ceases to receive traffic. In most cases, such behavior is of little use, since the traffic is usually distributed more or less evenly between the Pods. Thus, generally, readinessProbe either works everywhere or does not work on a large number of Pods at once. There are situations when such behavior can be useful. However, in my experience, that is for the most part under exceptional cases.

Still, readinessProbe comes with another crucial feature: it helps determine when a newly created container can handle the traffic so as not to forward load to an application that isn’t ready yet. This readinessProbe feature, au contraire, is necessary at all times.

In other words, one feature of readinessProbe is in high demand, while the other is not necessary at all. This dilemma was solved with the introduction of startupProbe. It first appeared in Kubernetes 1.16 becoming beta in v1.18 and stable in v1.20. Thus, you’re best off **using readinessProbe to check if an application is ready in Kubernetes versions below 1.18, but startupProbe – in Kubernetes versions 1.18 and up.** Then again, you can use readinessProbe in Kubernetes 1.18+ if you have any need to stop traffic to individual Pods after the application has been started.

### 7.3. Startup probe

startupProbe checks if an application in the container is ready. Then it marks the current Pod as ready to receive traffic or goes on updating/restarting the Deployment. Unlike readinessProbe, startupProbe stops working after the container has been started. We do not advise using startupProbe for checking external dependencies: its failure would trigger a container restart, which may eventually cause the Pod to go `CrashLoopBackOff`. In this state, the delay between attempts to restart a failed container can be as high as five minutes. It may lead to unnecessary downtime since, despite the application being _ready to be restarted_, the container continues to wait until the end of the `CrashLoopBackOff` period before trying to restart.

You should use startupProbe if your application receives traffic and your Kubernetes version is 1.18 or higher.

Also, note that increasing `failureThreshold` instead of setting `initialDelaySeconds` is the preferred method for configuring the probe. This will allow the container to become available as quickly as possible.

## 8\. Checking external dependencies

As you know, readinessProbe is often used for checking external dependencies (e.g. databases). While this approach has the right to exist, you’d be well advised to separate your means of checking for external dependencies and your means of checking whether the application in the Pod is running at full capacity (and cutting off the sending of traffic to it is a good idea as well).

You can use `initContainers` to check external dependencies before running the main containers’ startupProbe/readinessProbe. It’s pretty clear that in that case, you will no longer need to check external dependencies using readinessProbe. `initContainers` do not require changes to the application code. You do not need to embed additional tools to use them for checking external dependencies in the application containers. Usually, they are reasonably easy to implement:

```
      initContainers:
      - name: wait-postgres
        image: postgres:12.1-alpine
        command:
        - sh
        - -ec
        - |
          until (pg_isready -h example.org -p 5432 -U postgres); do
            sleep 1
          done
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 50m
            memory: 50Mi
      - name: wait-redis
        image: redis:6.0.10-alpine3.13
        command:
        - sh
        - -ec
        - |
          until (redis-cli -u redis://redis:6379/0 ping); do
            sleep 1
          done
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 50m
            memory: 50Mi
```

## Complete example

To sum it up, here is a complete example of the production-grade Deployment of a stateless application that comprises all the recommendations provided above.

_You will need Kubernetes 1.18 or higher and Ubuntu/Debian-based nodes with kernel version 5.4 or higher._

```
apiVersion: apps/v1
kind: Deployment
metadata:
name: testapp
spec:
replicas: 10
selector:
    matchLabels:
      app: testapp
template:
    metadata:
      labels:
        app: testapp
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchLabels:
                  app: testapp
              topologyKey: kubernetes.io/hostname
      priorityClassName: production-medium
      terminationGracePeriodSeconds: 40
      initContainers:
      - name: wait-postgres
        image: postgres:12.1-alpine
        command:
        - sh
        - -ec
        - |
          until (pg_isready -h example.org -p 5432 -U postgres); do
            sleep 1
          done
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 50m
            memory: 50Mi
      containers:
      - name: backend
        image: my-app-image:1.11.1
        command:
        - run
        - app
        - --trigger-graceful-shutdown-if-memory-usage-is-higher-than
        - 450Mi
        - --timeout-seconds-for-graceful-shutdown
        - 35s
        startupProbe:
          httpGet:
            path: /simple-startup-check-no-external-dependencies
            port: 80
          timeoutSeconds: 7
          failureThreshold: 12
        lifecycle:
          preStop:
            exec:
              ["sh", "-ec", "#command to shutdown gracefully if needed"]
        resources:
          requests:
            cpu: 200m
            memory: 500Mi
          limits:
            cpu: 200m
            memory: 500Mi
```

## In the next article

There are several other important topics that need to be addressed, such as `PodDisruptionBudget`, `HorizontalPodAutoscaler`, and `VerticalPodAutoscaler`. We will discuss them in Part 2 of this article **(UPDATE: [Part 2 is out!](https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-2/))** — subscribe to our blog below and/or follow [our Twitter](https://twitter.com/flant_com) not to miss it! Please share your best practices for deploying applications (or, if need be, you can correct/supplement the ones discussed above).

## Related posts:

- [Comparing Ingress controllers for Kubernetes](https://blog.flant.com/comparing-ingress-controllers-for-kubernetes/ "Comparing Ingress controllers for Kubernetes")
- [Migrating your app to Kubernetes: what to do with files?](https://blog.flant.com/migrating-your-app-to-kubernetes-what-to-do-with-files/ "Migrating your app to Kubernetes: what to do with files?")
- [How we enjoyed upgrading a bunch of Kubernetes clusters from v1.16 to v1.19](https://blog.flant.com/how-we-enjoyed-upgrading-kubernetes-clusters-from-v1-16-to-v1-19/ "How we enjoyed upgrading a bunch of Kubernetes clusters from v1.16 to v1.19")

## Share

© 2008–2022 [Flant Europe OÜ](http://flant.com)

- [Twitter](https://twitter.com/flant_com)
- [Facebook](https://www.facebook.com/flantcom)
- [LinkedIn](https://www.linkedin.com/company/flant)
- [GitHub](http://github.com/flant)
- [YouTube](https://www.youtube.com/c/Flant_com)

Subscribe me to your new tech posts:

✕



https://blog.flant.com/best-practices-for-deploying-highly-available-apps-in-kubernetes-part-1
