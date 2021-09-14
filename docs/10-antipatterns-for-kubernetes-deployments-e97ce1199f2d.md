# 10 Anti-Patterns for Kubernetes Deployments

## Common practices in Kubernetes deployments that have better solutions

[Aug 28, 2020](https://betterprogramming.pub/10-antipatterns-for-kubernetes-deployments-e97ce1199f2d?source=post_page-----e97ce1199f2d--------------------------------) · 22 min read

As container adoption and usage continues to rise, [Kubernetes](https://kubernetes.io/) (K8s) has become the leading platform for container orchestration. It’s an open-source project with tens of thousands of contributors from over 315 companies with the intention of remaining [extensible and cloud-agnostic](https://kubernetes.io/blog/2019/04/17/the-future-of-cloud-providers-in-kubernetes/), and it’s the foundation of every major cloud provider.

When you have containers running in production, you want your production  environment to be as stable and resilient as possible to avert disaster  (think every online Black Friday shopping experience). When a container  goes down, another one needs to spin up to take its place, no matter  what time of day — or into the wee hours of the night — it is.  Kubernetes provides a framework for running distributed systems  resiliently, from scaling to failover to load balancing and more. And  there are many tools that integrate with Kubernetes to help meet your  needs.

Best practices evolve with time, so it’s always good to continuously  research and experiment for better ways for Kubernetes development. As  it is still a young technology, we are always looking to improve our  understanding and use of it.

In this article, we’ll be examining ten common practices in Kubernetes  deployments that have better solutions at a high level. I will not go  into depth on the best practices since custom implementation might vary  among users.

1. Putting the configuration file inside/alongside the Docker image
2. Not using Helm or other kinds of templating
3. Deploying things in a specific order. (Applications shouldn’t crash because a dependency isn’t ready.)
4. Deploying pods without set memory and/or CPU limits
5. Pulling the `latest` tag in containers in production
6. Deploying new updates/fixes by killing pods so they pull the new Docker images during the restart process
7. Mixing both production and non-production workloads in the same cluster.
8. Not using blue/green or canaries for mission-critical deployments. (The  default rolling update of Kubernetes is not always enough.)
9. Not having metrics in place to understand if a deployment was successful or not. (Your health checks need application support.)
10. Cloud vendor lock-in: locking yourself into an IaaS provider’s Kubernetes or serverless computing services

# Ten Kubernetes Anti-Patterns

## **1. Putting the configuration file inside/alongside the Docker image**

This Kubernetes anti-pattern is related to a Docker anti-pattern (see anti-patterns 5 and 8 in [this article](https://codefresh.io/containers/docker-anti-patterns/)). Containers give developers a way to use a single image, essentially in  the production environment, through the entire software lifecycle, from  dev/QA to staging to production.

However, a common practice is to give each phase in the lifecycle its own image, each built with different artifacts specific to its environment (QA,  staging, or production). But now you’re no longer deploying what you’ve  tested.


![img](https://miro.medium.com/max/1400/0*8JxuyxQ75CJ0rHQl)

*Don’t hardcode your configuration at build time (from* [*https://codefresh.io/containers/docker-anti-patterns/*](https://codefresh.io/containers/docker-anti-patterns/)*)*

The best practice here is to externalize general-purpose configuration in [ConfigMaps](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/), while sensitive information (like API keys and secrets) can be stored  in the Secrets resource (which has Base64 encoding but otherwise works  the same as ConfigMaps). ConfigMaps can be [mounted as volumes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) or passed in as [environment variables](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/), but [Secrets](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) should be mounted as volumes. I mention ConfigMaps and Secrets because  they are native Kubernetes resources and don’t require integrations, but they can be limiting. There are other solutions available like [ZooKeeper](https://zookeeper.apache.org/) and [Consul by HashiCorp](https://www.consul.io/) for configmaps, or [Vault by HashiCorp](https://www.vaultproject.io/), [Keywhiz](https://square.github.io/keywhiz/), [Confidant](https://lyft.github.io/confidant/), etc, for secrets, that might better fit your needs.

When you’ve decoupled your configuration from your application, you no  longer need to recompile the application when you need to update the  configuration — and it can be updated while the app is running. Your  applications fetch the configuration during runtime instead of during  the build. More importantly, you’re using the same source code in all  the phases of the software lifecycle.


![img](https://miro.medium.com/max/1400/0*TWGDk6_OXrd0Ktvz)

*Load configuration during runtime (from* [*https://codefresh.io/containers/docker-anti-patterns/*](https://codefresh.io/containers/docker-anti-patterns/)*)*

## **2. Not using Helm or other kinds of templating**

You can manage Kubernetes deployments by directly [updating YAML](https://stackoverflow.com/questions/48191853/how-to-update-a-deployment-via-editing-yml-file). When rolling out a new version of code, you will probably have to update one or more of the following:

- Docker image name
- Docker image tag
- Number of replicas
- Service labels
- Pods
- Configmaps, etc

This can get tedious if you’re managing multiple clusters and applying the  same updates across your development, staging, and production  environments. You are basically modifying the same files with minor  modifications across all your deployments. It’s a lot of copy-and-paste, or search-and-replace, while also staying aware of the environment for  which your deployment YAML is intended. There are a lot of opportunities for mistakes during this process:

- Typos (wrong version numbers, misspelling image names, etc.)
- Modifying YAML with the wrong update (for example, connecting to the wrong database)
- Missing a resource to update, etc.

<iframe src="https://betterprogramming.pub/media/e3e8e1d8b404d7279f48d69344700e6c" allowfullscreen="" title="deployment.yaml" class="t u v lc aj" scrolling="auto" width="680" height="479" frameborder="0"></iframe>

There might be a number of things you might need to change in the YAML, and  if you’re not paying close attention, one YAML could be easily mistaken  for another deployment’s YAML.

[Templating](https://codefresh.io/docs/docs/deploy-to-kubernetes/kubernetes-templating/) helps streamline the installation and management of Kubernetes  applications. Since Kubernetes doesn’t provide a native templating  mechanism, we have to look elsewhere for this type of management.

[Helm](https://helm.sh/) was the first package manager available (2015). It was proclaimed to be “Homebrew for Kubernetes” and evolved to include templating  capabilities. Helm packages its resources via [*charts*](https://v2.helm.sh/docs/developing_charts/), where a chart is a collection of files describing a related set of  Kubernetes resources. There are 1,400+ publicly available charts in the [chart repository](https://hub.helm.sh/), (you can also use `helm search hub [keyword] [flags]`), basically reusable recipes for installing, upgrading, and uninstalling  things on Kubernetes. With Helm charts, you can modify the `values.yaml` file to set the modifications you need for your Kubernetes deployments, and you can have a different Helm chart for each environment. So if you have a QA, staging, and production environment, you only have to manage three Helm charts instead of modifying each YAML in each deployment in  each environment.

Another advantage we get with Helm is that it’s easy to roll back to a previous revision with [Helm rollbacks](https://helm.sh/docs/helm/helm_rollback/) if something goes wrong with:

`helm rollback <RELEASE> [REVISION] [flags]` .

If you want to roll back to the immediate prior version, you can use:

`helm rollback <RELEASE> 0` .

So we’d see something like:

```
$ helm upgrade — install — wait — timeout 20 demo demo/
$ helm upgrade — install — wait — timeout 20 — set
readinessPath=/fail demo demo/
$ helm rollback — wait — timeout 20 demo 1Rollback was a success.
```

And the Helm chart history tracks it nicely:

```
$ helm history demo
REVISION STATUS DESCRIPTION
1 SUPERSEDED Install complete
2 SUPERSEDED Upgrade “demo” failed: timed out waiting for the condition
3 DEPLOYED Rollback to 1
```

Google’s [Kustomize](https://kustomize.io/) is a popular alternative and can be [used in addition to Helm](https://helm.sh/docs/topics/advanced/).

## **3. Deploying things in a specific order**

Applications shouldn’t crash because a dependency isn’t ready. In traditional  development, there is a specific order to the startup and stop tasks  when bringing up applications. It’s important not to bring this mindset  into container orchestration. With Kubernetes, Docker, etc., these  components start concurrently, making it impossible to define a startup  order. Even when the application is up and running, its dependencies  could fail or be migrated, leading to further issues. The Kubernetes  reality is also riddled with myriad points of potential communication  failures where dependencies can’t be reached, during which a pod might  crash or a service might become unavailable. Network latency, like a  weak signal or interrupted network connection, is a common culprit for  communication failure.

For simplicity’s sake, let’s examine a hypothetical shopping application  that has two services: an inventory database and a storefront UI. Before the application can launch, the back-end service has to start, meet all its checks, and start running. Then the front-end service can start,  meet its checks, and start running.

Let’s say we’ve forced the deployment order with the `kubectl wait `command, something like:

```
kubectl wait — for=condition=Ready pod/serviceA
```

But when the condition is never met, the next deployment can’t proceed and the process breaks.

This is a simplistic flow of what a deployment order might look like:


![img](https://miro.medium.com/max/1260/0*JXePbmNwNkB0g_OB)

*This process cannot move forward until the previous step is complete*

Since Kubernetes is [self-healing](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy). The standard approach is to let all the services in an application  start concurrently and let the containers crash and restart until they  are all up and running. I have service A and B starting independently  (as a decoupled, stateless cloud-native application should), but for the sake of the user experience, perhaps I could tell the UI (service B) to display a pretty loading message until service A is ready, but the  actual starting up of service B shouldn’t be affected by service A.


![img](https://miro.medium.com/max/1384/0*ZqfJHyOsWoigya7o)

*Now when the pod crashes, Kubernetes restarts the service until everything  is up and running. If you are stuck in CrashLoopBackOff, it’s worth  checking your code, configuration, or for resource contention.*

Of course, we need to do more than simply rely on self-healing. We need to implement solutions that will handle failures, which are inevitable and will happen. We should anticipate they will happen and lay down  frameworks to respond in a way that helps us avoid downtime and/or data  loss.

In my hypothetical shopping app, my storefront UI (service B) needs the  inventory (service A) in order to give the user a complete experience.  So when there’s a partial failure, like if service A wasn’t available  for a short time or crashed, etc., the system should still be able to  recover from the issue.

Transient faults like these are an ever-present possibility, so to minimize their effects we can implement a [Retry pattern](https://docs.microsoft.com/en-us/azure/architecture/patterns/retry). Retry patterns help improve application stability with strategies like:

- **Cancel
  **If the fault isn’t transient or if the process is unlikely to be  successful on repeated attempts, then the application should cancel the  operation and report an exception — e.g., authentication failure.  Invalid credentials should never work!
- **Retry
  **If the fault is unusual or rare, it could be due to uncommon situations  (e.g., network packet corruption). The application should retry the  request immediately because the same failure is unlikely to reoccur.
- **Retry after delay
  **If the fault is caused by common occurrences like connectivity or busy  failures, it’s best to let any work backlog or traffic clear up before  trying again. The application should wait before retrying the request.
- You could also implement your retry pattern with an [exponential backoff](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/implement-resilient-applications/implement-retries-exponential-backoff) (exponentially increasing the wait time and setting a maximum retry count).

Implementing a [circuit-breaking pattern](https://istio.io/latest/docs/tasks/traffic-management/circuit-breaking/) is also an important strategy when creating resilient microservice  applications. Like how a circuit breaker in your house will  automatically switch to protect you from extensive damage due to excess  current or short-circuiting, the circuit-breaking pattern provides you a method of writing applications while limiting the impact of unexpected  faults that might take longer to fix, like partial loss of connectivity, or complete failure of a service. In these situations where retrying  won’t work, the application should be able to accept that the failure  has occurred and respond accordingly.

## **4. Deploying pods without set memory and/or CPU limits**

Resource allocation varies depending on the service, and it can be difficult to  predict what resources a container might require for optimal performance without testing implementation. One service could require a fixed CPU  and memory consumption profile, while another service’s consumption  profile could be dynamic.

When you deploy pods without careful consideration of memory and CPU limits, this can lead to scenarios of resource contention and unstable  environments. If a container does not have a memory or CPU limit, then  the scheduler sees its memory utilization (and CPU utilization) as zero, so an unlimited number of pods can be scheduled on any node. This can  result in the overcommitment of resources and possible node and kubelet  crashes.

When the memory limit is not specified for a container, there are a couple  of scenarios that could apply (these also apply to CPU):

1. There is no upper bound on the amount of memory a container can use. Thus,  the container could use all of the available memory on its node,  possibly invoking the OOM (out of memory) Killer. An OOM Kill situation  has a greater chance of occurring for a container with no resource  limits.
2. The default memory limit of the namespace (in which the container is  running) is assigned to the container. The cluster administrators can  use a [LimitRange](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#limitrange-v1-core) to specify a default value for the memory limit.

Declaring memory and CPU limits for the containers in your cluster allows you to  make efficient use of the resources available on your cluster’s nodes.  This helps the kube-scheduler determine on which node the pod should  reside for most efficient hardware utilization.

When setting the memory and CPU limits for a container, you should take care not to request more resources than the limit. For pods that have more  than one container, the aggregate resource requests must not exceed the  set limit(s) — otherwise, the pod will never be scheduled.


![img](https://miro.medium.com/max/1400/0*jLGrEtn42OPuqv-K)

*The resource request must not exceed the limit*

Setting memory and CPU requests below their limits accomplishes two things:

1. The pod can make use of memory/CPU when it is available, leading to bursts of activity.
2. During a burst, the pod is limited to a reasonable amount of memory/CPU.

The best practice is to keep the CPU request at one core or below, and then use [ReplicaSets](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) to scale it out, which gives the system flexibility and reliability.

What happens when you have different teams competing for resources when  deploying containers in the same cluster? If the process exceeds the  memory limit, then it will be terminated, while if it exceeds the CPU  limit, the process will be throttled (resulting in worse performance).

You can control resource limits via [resource quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/) and [LimitRange](https://kubernetes.io/docs/concepts/policy/limit-range/) in the [namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) settings. These settings help account for containers deployments without limits or with high resource requests.

Setting hard resource limits might not be the best choice for your needs.  Another option is to use the recommendation mode in the [Vertical Pod autoscaler ](https://cloud.google.com/kubernetes-engine/docs/concepts/verticalpodautoscaler)resource.

## **5. Pulling the ‘**`**latest'** `**tag in containers in production**

[Using the ](https://kubernetes.io/docs/concepts/containers/images/#image-names)`latest`[ tag](https://kubernetes.io/docs/concepts/containers/images/#image-names) is considered bad practice, especially in production. Pods unexpectedly crash for all sorts of reasons, so they can pull down images at any  time. Unfortunately, the `latest `tag is not very descriptive when it comes to determining when the build  broke. What version of the image was running? When was the last time it  was working? This is especially bad in production since you need to be  able to get things back up and running with minimal downtime.


![img](https://miro.medium.com/max/1400/0*9paNbPUoky-4F3uX)

*You shouldn’t use the* `*latest* `*tag in production.*

By default, the `imagePullPolicy` is set to `Always` and will always pull down the image when it restarts. If you don’t specify a tag, Kubernetes will default to `latest`. However, a deployment will only be updated in the event of a crash  (when the pod pulls down the image on restart) or if the deployment  pod’s template (`.spec.template`) is [changed](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#updating-a-deployment). See [this forum discussion](https://discuss.kubernetes.io/t/use-latest-image-tag-to-update-a-deployment/2929) for an example of `latest` not working as intended in development.

Even if you’ve changed the `imagePullPolicy` to another value than `Always`, your pod will still pull an image if it needs to restart (whether it’s  because of a crash or deliberate reboot). If you use versioning and set  the `imagePullPolicy` with a meaningful tag, like v1.4.0, then you can [roll back](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#rolling-back-to-a-previous-revision) to the most recent stable version and more easily troubleshoot when and where something went wrong in your code. You can read more about best  practices for versioning in the [Semantic Versioning Specification](https://semver.org/) and [GCP Best Practices](https://semver.org/).

In addition to using specific and meaningful Docker tags, you should also  remember that containers are stateless and immutable. They are also  meant to be ephemeral (and you should store any data outside containers  in persistent storage). Once you spin up a container, you should not  modify it: no patches, no updates, no configuration changes. When you  need to update a configuration, you should deploy a new container with  the updated config.


![img](https://miro.medium.com/max/1400/0*1dTSSbrBOir4Kw3_)

[*Docker immutability, taken from Best Practices for Operating Containers*](https://cloud.google.com/solutions/best-practices-for-operating-containers)*.*

This immutability allows for safer and repeatable deployments. You can also  more easily roll back if you need to redeploy the old image. By keeping  your Docker images and container immutable, you are able to deploy the  same container image in every single environment. See Anti-pattern 1 to  read about externalizing your configuration data to keep your images  immutable.


![img](https://miro.medium.com/max/1400/0*dtF2W7jkk_-D-bok)

*We can roll back to the previous stable version while we troubleshoot.*

## **6. Deploying new updates/fixes by killing pods so they pull the new Docker images during the restart process**

Like relying on the latest tag to pull updates, relying on killing pods to  roll out new updates is bad practice since you’re not versioning your  code. If you are killing pods to pull updated Docker images in  production, don’t do it. Once a version has been released in production, it should never be overwritten. If something breaks, then you won’t  know where or when things went wrong and how far back to go when you  need to roll back the code while you troubleshoot.

Another problem is that restarting the container to pull a new Docker image  doesn’t always work. “A Deployment’s rollout is triggered if and only if the Deployment’s Pod template (that is, `.spec.template`) is changed, for example if the labels or container images of the  template are updated. Other updates, such as scaling the Deployment, do  not trigger a rollout.”

<iframe src="https://betterprogramming.pub/media/023ced411f5076fa92be0e796b90be38" allowfullscreen="" title="deployment-1.yaml" class="t u v lc aj" scrolling="auto" width="680" height="479" frameborder="0"></iframe>

You have to modify the `.spec.template` to trigger a deployment.

The correct way to update your pods to pull new Docker images is to [version](https://semver.org/) (or increment for fixes/patches) your code and then modify the deployment spec to reflect a meaningful tag (not `latest`, see Anti-pattern 5) for further discussion on that, but something like  v1.4.0 for a new release or v1.4.1 for a patch). Kubernetes will then  trigger an upgrade with zero downtime.

1. Kubernetes starts a new pod with the new image.
2. Waits for health checks to pass.
3. Deletes the old pod.

## **7. Mixing both production and non-production workloads in the same cluster**

Kubernetes supports a [namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) feature, which enables users to manage different environments (virtual  clusters) within the same physical cluster. Namespaces can be seen as a  cost-effective way of managing different environments on a single  physical cluster. For example, you could run staging and production  environments in the same cluster and save resources and money. However,  there’s a big gap between running Kubernetes in development and running  Kubernetes in production.

There are a lot of factors to consider when you mix your production and  non-production workloads on the same cluster. For one, you would have to consider resource limits to make sure the performance of your  production environment isn’t compromised (a common practice one might  see is setting no quota on the production namespace and a quota on any  non-production namespace(s)).

You would also need to consider isolation. Developers require a lot more  access and permissions than in production, which you would want locked  down as much as possible. While namespaces are hidden from each other,  they are not fully isolated by default. That means your apps in a dev  namespace could call apps in test, staging, or production (or vice  versa), which is not considered good practice. Of course, you could use  NetworkPolicies to set rules to isolate the namespaces.

However, thoroughly testing resource limits, performance, security, and  reliability is time-consuming, so running production workloads in the  same cluster as non-production workloads is not advised. Rather than  mixing production and non-production workloads in the same cluster, use  separate clusters for development/test/production — you’ll have better  isolation and security that way. You should also automate as much as you can for CI/CD and promotion to reduce the chance for human error. Your  production environment needs to be as solid as possible.

## **8. Not using blue/green or canaries for mission-critical deployments**

Many modern applications have frequent deployments, ranging from several  changes within a month to multiple deployments in a single day. This is  certainly achievable with microservice architecture since the different  components can be developed, managed, and released on different cycles  as long as they work together to perform seamlessly. And of course,  keeping applications up 24/7 is obviously important when rolling out  updates.

The default rolling update of Kubernetes is not always enough. A common  strategy to perform updates is to use the default Kubernetes [rolling update](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#rolling-update-deployment) feature:

```
.spec.strategy.type==RollingUpdate
```

where you can set the `maxUnavailable` (percentage or number of pods unavailable) and `maxSurge` fields (optional) to control the rolling update process. When  implemented properly, rolling updates allow a gradual update with zero  downtime as the pods are incrementally updated. Here’s an [example](https://medium.com/platformer-blog/enable-rolling-updates-in-kubernetes-with-zero-downtime-31d7ec388c81) of how one team updated their applications with zero downtime with rolling updates.

However, once you’ve updated your deployment to the next version, it’s not  always easy to go back. You should have a plan in place to roll it back  in case it breaks in production. When your pod is updated to the next  version, the deployment will create a new ReplicaSet. While Kubernetes  will store previous ReplicaSets (by default, it’s ten, but you could  change that with `spec.revisionHistoryLimit`). The ReplicaSets are saved under names such as `app6ff34b8374` in random order, and you won’t find a reference to the ReplicaSets in the deployment app YAML. You could find it with:

```
ReplicaSet.metatada.annotation
```

and inspect the revision with:

```
kubectl get replicaset app-6ff88c4474 -o yaml
```

to find the revision number. This gets complicated because the rollout  history doesn’t keep a log unless you leave a note in the YAML resource  (which you could do with the `— record` flag:

```
$kubectl rollout history deployment/appREVISION CHANGE-CAUSE1 kubectl create — filename=deployment.yaml — record=true
2 kubectl apply — filename=deployment.yaml — record=true
```

When you have dozens, hundreds, or even thousands of deployments all going  through updates simultaneously, it’s difficult to keep track of them all at once. And if your stored revisions all contain the same regression,  then your production environment is not going to be in good shape! You  can read more in detail about using rolling updates [in this article](https://learnk8s.io/kubernetes-rollbacks#:~:text=In Kubernetes%2C rolling updates are,bring newer Pod in incrementally.&text=You have a Service and,three replicas on version 1.0.).

Some other problems are:

- Not all applications are capable of concurrently running multiple versions.
- Your cluster could run out of resources in the middle of the update, which could break the whole process.

These are all very frustrating and stressful issues to run into when in a production environment.

Alternative ways to more reliably update deployments include:

**Blue/green (red/black) deployment**
With blue/green, a full set of both the old and new instances exist  simultaneously. Blue is the live version, and the new version is  deployed to the green replica. When the green environment has passed its tests and verifications, a load balancer simply flips the traffic to  the green, which becomes the blue environment, and the old version  becomes the green version. Since we have two full versions being  maintained, performing a rollback is simple — all you need to do is  switch back the load balancer.


![img](https://miro.medium.com/max/1228/0*LQzr9w-ADoabf8rJ)

*The load balancer flips between blue and green to set the active version. From* [*Continuous Deployment Strategies with Kubernetes*](https://codefresh.io/kubernetes-tutorial/continuous-deployment-strategies-kubernetes-2/)*.*

Additional advantages include:

- Since we never deploy directly to production, it’s pretty low-stress when we change green to blue.
- Traffic redirection occurs immediately, so there’s no downtime.
- There can be extensive testing done to reflect actual production prior to the switch. (As stated before, a development environment is very different  from production.)

Kubernetes does not include blue/green deployments as one of its native toolings.  You can read more about how to implement blue/green into your CI/CD  automation [in this tutorial](https://codefresh.io/kubernetes-tutorial/fully-automated-blue-green-deployments-kubernetes-codefresh/).

**Canary Releases
**Canary releases allow us to test for potential problems and meet key metrics  before impacting the entire production system/user base. We “test in  production” by deploying directly to the production environment, but  only to a small subset of users. You can choose routing to be  percentage-based or driven by region/user location, the type of client,  and billing properties. Even when deploying to a small subset, it’s  important to carefully monitor application performance and measure  errors — these metrics define a quality threshold. If the application  behaves as expected, we start transferring more of the new version  instances to support more traffic.


![img](https://miro.medium.com/max/1400/0*pzALOq2I9-ws85X9)

*The load balancer gradually releases the new version into production. From* [*Continuous Deployment Strategies with Kubernetes*](https://codefresh.io/kubernetes-tutorial/continuous-deployment-strategies-kubernetes-2/)*.*

Other advantages include:

- Observability
- Ability to test on production traffic (getting a true production-like experience in development is hard)
- Ability to release a version to a small subset of users and get real feedback before a larger release
- Fail fast. Since we deploy straight into production, we can fail fast (i.e., revert immediately) if it breaks, and it affects only a subset rather  than the whole community.

## **9. Not having metrics in place to understand if a deployment was successful or not**

Your health checks need application support.

You can leverage Kubernetes to accomplish many tasks in container orchestration:

- Controlling resource consumption by an application or team (namespace, CPU/mem,  limits) stopping an app from consuming too many resources
- Load balancing across different app instances, moving application instances  from one host to another if there is a resource shortage or if the host  dies
- Self-healing — restarting containers if they crash
- Automatically leveraging additional resources if a new host is added to the cluster
- And more

So sometimes it’s easy to forget about metrics and monitoring. However, a  successful deployment is not the end of your ops work. It’s better to be proactive and prepare for unexpected surprises. There are still a lot  more layers to monitor, and the dynamic nature of K8s makes it tough to  troubleshoot. For example, if you’re not closely watching your resource  available, the automatic rescheduling of pods could cause capacity  issues, and your app might crash or never deploy. This would be  especially unfortunate in production, as you wouldn’t know unless  someone filed a bug report or if you happened to check on it. Eep!

Monitoring presents its own set of challenges: There are a lot of layers to watch, and there’s a need to “[maintain a reasonably low maintenance burden on the engineers](https://landing.google.com/sre/sre-book/chapters/practical-alerting/).” When an application running on Kubernetes hits a snag, there are many  logs, data, and components to investigate, especially when there are  multiple microservices involved with the issue versus in traditional  monolithic architecture, where everything is output to a few logs.

Insights on your application behavior, like how an application performs, helps  you continuously improve. You also need a pretty holistic view of the  containers, pods, services, and the cluster as a whole. If you can  identify how an application is using its resources, then you can use  Kubernetes to better detect and remove bottlenecks. To get a full view  of the application, you would need to use an application performance  monitoring solution like [Prometheus](https://prometheus.io/), [Grafana](https://grafana.com/), [New Relic](https://newrelic.com/), or [Cisco AppDynamics](https://www.appdynamics.com/appd-campaigns/?utm_source=adwords&utm_medium=ppc&utm_campaign=brand&gclid=CjwKCAjwydP5BRBREiwA-qrCGo92C606PzpGx6nOZdhkIs8WcxHadyb-gYDCeUfofm3hBSgeTTAW8BoC7DQQAvD_BwE), among many others.

Whether or not you decide to use a monitoring solution, these are the key  metrics that the Kubernetes documentation recommends you track closely:

- Running pods and their deployments
- Resource metrics: CPU, memory usage, disk I/O
- Container-native metrics
- Application metrics

## **10. Cloud vendor lock-in: Locking yourself into an IaaS provider’s Kubernetes or serverless computing services**

There are multiple types of lock-ins (Martin Fowler wrote a great [article](https://martinfowler.com/articles/oss-lockin.html), if you want to read more), but vendor lock-in negates the primary value of deploying to the cloud: container flexibility. It’s true that  choosing the right cloud provider is not an easy decision. Each provider has its own interfaces, open APIs, and proprietary specifications and  standards. Additionally, one provider might suit your needs better than  the others only for your business needs to unexpectedly change.

Fortunately, containers are platform-agnostic and portable, and all the major  providers have a Kubernetes foundation, which is cloud-agnostic. You  don’t have to re-architect or rewrite your application code when you  need to move workloads between clouds, so you shouldn’t need to lock  yourself into a cloud provider because you can’t “lift and shift.”

Here is a list of things you should consider to ensure you can be flexible to prevent or minimize vendor lock-in.

**First,** [**housekeeping**](https://www.techrepublic.com/article/5-ways-to-avoid-vendor-lock-in/)**: Read the fine print**

Negotiate entry and exit strategies. Many vendors make it easy to start — and get you hooked. This might include incentives like free trials or credits,  but these costs could rapidly increase as you scale up.

Check for things like auto-renewal, early termination fees, and if the  provider will help with things like deconversion when migrating to  another vendor and SLAs associated with exit.

**Architect/design your applications such that they can run on any cloud**

If you’re already developing for the cloud and using cloud-native  principles, then most likely your application code should be easy to  lift and shift. It’s the things surrounding the code that potentially  lock you into a cloud vendor. For example, you could:

- Check that your services and features (like databases, APIs, etc) used by your application are portable.
- Check if your deployment and provisioning scripts are cloud-specific. Some  clouds have their own native or recommended automation tools that may  not translate easily to other providers. There are many tools that can  be used to assist with cloud infrastructure automation and are  compatible with many of the major cloud providers, like [Puppet](https://puppet.com/), [Ansible](https://www.ansible.com/), and [Chef](https://www.chef.io/), to name a few. This [blog](https://www.ibm.com/cloud/blog/chef-ansible-puppet-terraform) has a handy chart that compares characteristics of common tools.
- Check if your DevOps environment, which typically includes Git and CI/CD, can run in any cloud. For example, many clouds have their own specific  CI/CD tools, like [IBM Cloud Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery/pipeline_about.html#deliverypipeline_about), [Azure CI/CD](https://docs.microsoft.com/en-us/azure/architecture/solution-ideas/articles/azure-devops-continuous-integration-and-continuous-deployment-for-azure-web-apps), or [AWS Pipelines](https://aws.amazon.com/getting-started/projects/set-up-ci-cd-pipeline/), that might require extra work to port over to another cloud vendor. Instead, you could use something like [Codefresh](https://codefresh.io/?utm_source=google&utm_medium=cpc&utm_campaign=brand-search&utm_term=codefresh&gclid=CjwKCAjwmrn5BRB2EiwAZgL9ogt93s-yRXqtsuBb65KqDJftz-6biYdFJ1LPdwMlV6y9pLPlRy4yxhoCr10QAvD_BwE), full CI/CD solutions that have great support for Docker and Kubernetes  and integrates with many other popular tools. There are also myriad  other solutions, some CI or CD, or both, like [GitLab](https://about.gitlab.com/blog/2019/07/15/finding-the-right-ci-cd/), [Bamboo](https://www.atlassian.com/software/bamboo), [Jenkins](https://www.jenkins.io/), [Travis](https://www.jenkins.io/), etc.
- Check if your testing process will need to be changed between providers.

**You could also choose to follow a** [**multicloud strategy**](https://www.techrepublic.com/article/multicloud-the-smart-persons-guide/)

With a multicloud strategy, you can pick and choose services from different  cloud providers that best the type of application(s) you are hoping to  deliver. When you plan a multicloud deployment, you should keep  interoperability in careful consideration.

# Summary

Kubernetes is really popular, but it’s hard to get started with, and there are a  lot of practices in traditional development that don’t translate to  cloud-native development.

In this article, we’ve looked at:

1. Putting the configuration file inside/alongside the Docker image: **Externalise your configuration data. You can use ConfigMaps and Secrets or something similar.**
2. Not using Helm or other kinds of templating: **Use Helm or Kustomize to streamline your container orchestration and reduce human error.**
3. Deploying things in a specific order: **Applications shouldn’t crash because a dependency isn’t ready. Utilize Kubernetes’s  self-healing mechanism and implement retries and circuit breakers.**
4. Deploying pods without set memory and/or CPU limits: **You should consider setting memory and CPU limits to reduce the risk of  resource contention, especially when sharing the cluster with others.**
5. Pulling the `latest` tag in containers in production: **Never use** `**latest**`**. Always use something meaningful, like v1.4.0/according to** [**Semantic Versioning Specification**](https://semver.org/)**, and employ immutable Docker images.**
6. Deploying new updates/fixes by killing pods so they pull the new Docker images during the restart process: **Version your code so you can better manage your releases.**
7. Mixing both production and non-production workloads in the same cluster: **Run your production and non-production workloads in separate clusters if  you can. This reduces risk to your production environment from resource  contention and accidental environment cross-over.**
8. Not using blue/green or canaries for mission-critical deployments (the  default rolling update of Kubernetes is not always enough): **You should consider blue/green deployment or canary releases for less  stress in production and more meaningful production results.**
9. Not having metrics in place to understand if a deployment was successful or not (your health checks need application support): **You should make sure to monitor your deployment to avoid any surprises. You could use a tool like Prometheus, Grafana, New Relic, or Cisco  AppDynamics to help you gain better insights on your deployments.**
10. Cloud vendor lock-in: Locking yourself into an IaaS provider’s Kubernetes or serverless computing services: **Your business needs could change at any time. You shouldn’t unintentionally  lock yourself into a cloud provider since you can easily lift and shift  cloud-native applications.**

Thanks for reading!

[Better Programming](https://betterprogramming.pub/?source=post_sidebar--------------------------post_sidebar-----------)

Advice for programmers.
