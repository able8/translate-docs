# 10 Anti-Patterns for Kubernetes Deployments

# Kubernetes 部署的 10 个反模式

## Common practices in Kubernetes deployments that have better solutions

## Kubernetes 部署中的常见做法，有更好的解决方案

[Aug 28, 2020](https://betterprogramming.pub/10-antipatterns-for-kubernetes-deployments-e97ce1199f2d?source=post_page-----e97ce1199f2d--------------------------------) · 22 min read

As container adoption and usage continues to rise, [Kubernetes](https://kubernetes.io/)(K8s) has become the leading platform for container orchestration. It's an open-source project with tens of thousands of contributors from over 315 companies with the intention of remaining [extensible and cloud-agnostic](https://kubernetes.io/blog/2019/04/17/the-future-of-cloud-providers-in-kubernetes/), and it's the foundation of every major cloud provider.

随着容器采用和使用的不断增加，[Kubernetes](https://kubernetes.io/)(K8s) 已成为容器编排的领先平台。这是一个开源项目，拥有来自超过 315 家公司的数万名贡献者，旨在保持 [可扩展且与云无关](https://kubernetes.io/blog/2019/04/17/the-future-of-cloud-providers-in-kubernetes/)，它是每个主要云提供商的基础。

When you have containers running in production, you want your production  environment to be as stable and resilient as possible to avert disaster  (think every online Black Friday shopping experience). When a container  goes down, another one needs to spin up to take its place, no matter  what time of day — or into the wee hours of the night — it is. Kubernetes provides a framework for running distributed systems  resiliently, from scaling to failover to load balancing and more. And  there are many tools that integrate with Kubernetes to help meet your  needs.

当您在生产中运行容器时，您希望生产环境尽可能稳定和有弹性以避免灾难（想想每个在线黑色星期五购物体验）。当一个容器发生故障时，另一个容器需要旋转以取代它的位置，无论是一天中的什么时间——或者是深夜——它都是如此。 Kubernetes 提供了一个用于弹性运行分布式系统的框架，从扩展到故障转移到负载平衡等等。并且有许多与 Kubernetes 集成的工具可以帮助满足您的需求。

Best practices evolve with time, so it’s always good to continuously  research and experiment for better ways for Kubernetes development. As  it is still a young technology, we are always looking to improve our  understanding and use of it.

最佳实践会随着时间的推移而发展，因此不断研究和试验更好的 Kubernetes 开发方法总是好的。由于它仍然是一项年轻的技术，我们一直在寻求提高对它的理解和使用。

In this article, we’ll be examining ten common practices in Kubernetes  deployments that have better solutions at a high level. I will not Go into depth on the best practices since custom implementation might vary  among users.

在本文中，我们将研究 Kubernetes 部署中的 10 种常见做法，这些做法在较高层面上具有更好的解决方案。我不会深入探讨最佳实践，因为自定义实现可能因用户而异。

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



1. 将配置文件放在 Docker 镜像中/旁边
2. 不使用 Helm 或其他类型的模板
3. 以特定顺序部署事物。 （应用程序不应因为依赖项尚未准备好而崩溃。）
4. 在没有设置内存和/或 CPU 限制的情况下部署 Pod
5. 在生产中的容器中拉取“latest”标签
6. 通过杀死 Pod 来部署新的更新/修复，以便它们在重启过程中拉取新的 Docker 镜像
7. 在同一集群中混合生产和非生产工作负载。
8. 不使用蓝/绿或金丝雀进行关键任务部署。 （Kubernetes 的默认滚动更新并不总是足够的。）
9. 没有适当的指标来了解部署是否成功。 （您的健康检查需要应用程序支持。）
10. 云供应商锁定：将自己锁定在 IaaS 提供商的 Kubernetes 或无服务器计算服务中

# Ten Kubernetes Anti-Patterns

# 十个 Kubernetes 反模式

## **1. Putting the configuration file inside/alongside the Docker image**

## **1。将配置文件放在 Docker 镜像内/旁边**

This Kubernetes anti-pattern is related to a Docker anti-pattern (see anti-patterns 5 and 8 in [this article](https://codefresh.io/containers/docker-anti-patterns/)). Containers give developers a way to use a single image, essentially in  the production environment, through the entire software lifecycle, from  dev/QA to staging to production.

此 Kubernetes 反模式与 Docker 反模式相关（请参阅 [本文](https://codefresh.io/containers/docker-anti-patterns/) 中的反模式 5 和 8)。容器为开发人员提供了一种使用单个镜像的方法，主要是在生产环境中，贯穿整个软件生命周期，从开发/质量保证到过渡到生产。

However, a common practice is to give each phase in the lifecycle its own image, each built with different artifacts specific to its environment (QA,  staging, or production). But now you’re no longer deploying what you’ve  tested.

然而，一种常见的做法是为生命周期中的每个阶段提供自己的映像，每个阶段都使用特定于其环境（QA、暂存或生产）的不同工件构建。但是现在您不再部署您测试过的内容。

![img](https://miro.medium.com/max/1400/0*8JxuyxQ75CJ0rHQl)

*Don't hardcode your configuration at build time (from* [*https://codefresh.io/containers/docker-anti-patterns/*](https://codefresh.io/containers/docker-anti-patterns/)*)* 

*不要在构建时对您的配置进行硬编码（来自* [*https://codefresh.io/containers/docker-anti-patterns/*](https://codefresh.io/containers/docker-anti-patterns/)*)*

The best practice here is to externalize general-purpose configuration in [ConfigMaps](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/), while sensitive information (like API keys and secrets) can be stored  in the Secrets resource (which has Base64 encoding but otherwise works  the same as ConfigMaps). ConfigMaps can be [mounted as volumes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) or passed in as [environment variables](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/), but [Secrets](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) should be mounted as volumes. I mention ConfigMaps and Secrets because  they are native Kubernetes resources and don’t require integrations, but they can be limiting. There are other solutions available like [ZooKeeper](https://zookeeper.apache.org/) and [Consul by HashiCorp](https://www.consul.io/) for configmaps, or [Vault by HashiCorp](https://www.vaultproject.io/), [Keywhiz](https://square.github.io/keywhiz/),[Confidant](https://lyft.github.io/confidant/), etc, for secrets, that might better fit your needs.

这里的最佳实践是将 [ConfigMaps](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) 中的通用配置外部化，而敏感信息（如 API 密钥和secrets）可以存储在 Secrets 资源中（它具有 Base64 编码，但在其他方面与 ConfigMaps 相同)。 ConfigMap 可以[挂载为卷](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) 或作为[环境变量](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/)，但是 [Secrets](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) 应该挂载为卷。我提到 ConfigMaps 和 Secrets 是因为它们是本机 Kubernetes 资源并且不需要集成，但它们可能会受到限制。还有其他可用的解决方案，如 [ZooKeeper](https://zookeeper.apache.org/) 和 [Consul by HashiCorp](https://www.consul.io/) 用于配置映射，或 [Vault by HashiCorp](https://www.vaultproject.io/)、[Keywhiz](https://square.github.io/keywhiz/)、[Confidant](https://lyft.github.io/confidant/) 等，用于秘密，这可能更适合您的需求。

When you’ve decoupled your configuration from your application, you no  longer need to recompile the application when you need to update the  configuration — and it can be updated while the app is running. Your  applications fetch the configuration during runtime instead of during  the build. More importantly, you’re using the same source code in all  the phases of the software lifecycle.

当您将配置与应用程序分离后，您不再需要在需要更新配置时重新编译应用程序——并且可以在应用程序运行时进行更新。您的应用程序在运行时而不是在构建期间获取配置。更重要的是，您在软件生命周期的所有阶段都使用相同的源代码。

![img](https://miro.medium.com/max/1400/0*TWGDk6_OXrd0Ktvz)

*Load configuration during runtime (from* [*https://codefresh.io/containers/docker-anti-patterns/*](https://codefresh.io/containers/docker-anti-patterns/)*)*

*在运行时加载配置（来自* [*https://codefresh.io/containers/docker-anti-patterns/*](https://codefresh.io/containers/docker-anti-patterns/)*)*

## **2. Not using Helm or other kinds of templating**

## **2。不使用 Helm 或其他类型的模板**

You can manage Kubernetes deployments by directly [updating YAML](https://stackoverflow.com/questions/48191853/how-to-update-a-deployment-via-editing-yml-file). When rolling out a new version of code, you will probably have to update one or more of the following:

您可以通过直接[更新 YAML](https://stackoverflow.com/questions/48191853/how-to-update-a-deployment-via-editing-yml-file) 来管理 Kubernetes 部署。在推出新版本的代码时，您可能需要更新以下一项或多项：

- Docker image name
- Docker image tag
- Number of replicas
- Service labels
- Pods
- Configmaps, etc

- Docker 镜像名称
- Docker 镜像标签
- 副本数量
- 服务标签
- Pods
- 配置映射等

This can get tedious if you’re managing multiple clusters and applying the  same updates across your development, staging, and production  environments. You are basically modifying the same files with minor  modifications across all your deployments. It’s a lot of copy-and-paste, or search-and-replace, while also staying aware of the environment for  which your deployment YAML is intended. There are a lot of opportunities for mistakes during this process:

如果您要管理多个集群并在开发、暂存和生产环境中应用相同的更新，这可能会变得乏味。您基本上是在对所有部署进行细微修改后修改相同的文件。这是大量的复制和粘贴或搜索和替换，同时还要注意部署 YAML 的目标环境。在这个过程中有很多出错的机会：

- Typos (wrong version numbers, misspelling image names, etc.)
- Modifying YAML with the wrong update (for example, connecting to the wrong database)
- Missing a resource to update, etc.

- 错别字（错误的版本号、拼写错误的图像名称等）
- 使用错误的更新修改 YAML（例如，连接到错误的数据库）
- 缺少要更新的资源等。


There might be a number of things you might need to change in the YAML, and  if you’re not paying close attention, one YAML could be easily mistaken  for another deployment’s YAML.

您可能需要在 YAML 中更改许多内容，如果您不密切注意，一个 YAML 很容易被误认为是另一个部署的 YAML。

[Templating](https://codefresh.io/docs/docs/deploy-to-kubernetes/kubernetes-templating/) helps streamline the installation and management of Kubernetes  applications. Since Kubernetes doesn’t provide a native templating  mechanism, we have to look elsewhere for this type of management. 

[Templating](https://codefresh.io/docs/docs/deploy-to-kubernetes/kubernetes-templating/) 有助于简化 Kubernetes 应用程序的安装和管理。由于 Kubernetes 不提供原生模板机制，我们必须在别处寻找这种类型的管理。

[Helm](https://helm.sh/) was the first package manager available (2015). It was proclaimed to be “Homebrew for Kubernetes” and evolved to include templating  capabilities. Helm packages its resources via [*charts*](https://v2.helm.sh/docs/developing_charts/), where a chart is a collection of files describing a related set of  Kubernetes resources. There are 1,400+ publicly available charts in the [chart repository](https://hub.helm.sh/), (you can also use `helm search hub [keyword] [flags]`), basically reusable recipes for installing, upgrading, and uninstalling  things on Kubernetes. With Helm charts, you can modify the `values.yaml` file to set the modifications you need for your Kubernetes deployments, and you can have a different Helm chart for each environment. So if you have a QA, staging, and production environment, you only have to manage three Helm charts instead of modifying each YAML in each deployment in  each environment.

[Helm](https://helm.sh/) 是第一个可用的包管理器 (2015)。它被宣布为“Kubernetes 的 Homebrew”，并演变为包括模板功能。 Helm 通过 [*charts*](https://v2.helm.sh/docs/developing_charts/) 打包其资源，其中图表是描述一组相关 Kubernetes 资源的文件集合。 [chart repository](https://hub.helm.sh/)中有1,400多个公开可用的chart，（你也可以使用`helm search hub [keyword] [flags]`)，基本上是可重用的安装recipe，升级和卸载 Kubernetes 上的东西。使用 Helm 图表，您可以修改 `values.yaml` 文件以设置您的 Kubernetes 部署所需的修改，并且您可以为每个环境使用不同的 Helm 图表。所以如果你有 QA、staging 和生产环境，你只需要管理三个 Helm chart，而不是在每个环境的每个部署中修改每个 YAML。

Another advantage we get with Helm is that it’s easy to roll back to a previous revision with [Helm rollbacks](https://helm.sh/docs/helm/helm_rollback/) if something goes wrong with:

我们使用 Helm 获得的另一个优点是，如果出现问题，可以使用 [Helm rollbacks](https://helm.sh/docs/helm/helm_rollback/) 轻松回滚到以前的修订：

`helm rollback <RELEASE> [REVISION] [flags]` .

If you want to roll back to the immediate prior version, you can use:

如果要回滚到之前的版本，可以使用：

`helm rollback <RELEASE> 0` .

So we’d see something like:

所以我们会看到类似的东西：

```
$ helm upgrade — install — wait — timeout 20 demo demo/
$ helm upgrade — install — wait — timeout 20 — set
readinessPath=/fail demo demo/
$ helm rollback — wait — timeout 20 demo 1Rollback was a success.
```

And the Helm chart history tracks it nicely:

Helm 图表历史记录很好地跟踪了它：

```
$ helm history demo
REVISION STATUS DESCRIPTION
1 SUPERSEDED Install complete
2 SUPERSEDED Upgrade “demo” failed: timed out waiting for the condition
3 DEPLOYED Rollback to 1
```

Google’s [Kustomize](https://kustomize.io/) is a popular alternative and can be [used in addition to Helm](https://helm.sh/docs/topics/advanced/).

Google 的 [Kustomize](https://kustomize.io/) 是一种流行的替代方案，可以 [与 Helm 一起使用](https://helm.sh/docs/topics/advanced/)。

## **3. Deploying things in a specific order**

## **3。按特定顺序部署事物**

Applications shouldn’t crash because a dependency isn’t ready. In traditional  development, there is a specific order to the startup and stop tasks  when bringing up applications. It’s important not to bring this mindset  into container orchestration. With Kubernetes, Docker, etc., these  components start concurrently, making it impossible to define a startup  order. Even when the application is up and running, its dependencies  could fail or be migrated, leading to further issues. The Kubernetes  reality is also riddled with myriad points of potential communication  failures where dependencies can’t be reached, during which a pod might  crash or a service might become unavailable. Network latency, like a  weak signal or interrupted network connection, is a common culprit for  communication failure.

应用程序不应该因为依赖项还没有准备好而崩溃。在传统开发中，在启动应用程序时，启动和停止任务有特定的顺序。重要的是不要将这种思维方式带入容器编排中。使用Kubernetes、Docker等，这些组件是并发启动的，无法定义启动顺序。即使应用程序启动并运行，其依赖项也可能会失败或迁移，从而导致更多问题。 Kubernetes 的现实还充斥着无数潜在的通信故障点，在这些故障点无法到达依赖项，在此期间 pod 可能崩溃或服务可能变得不可用。网络延迟，如信号弱或网络连接中断，是通信失败的常见罪魁祸首。

For simplicity’s sake, let’s examine a hypothetical shopping application  that has two services: an inventory database and a storefront UI. Before the application can launch, the back-end service has to start, meet all its checks, and start running. Then the front-end service can start,  meet its checks, and start running.

为简单起见，让我们检查一个假设的购物应用程序，它有两个服务：库存数据库和店面 UI。在应用程序可以启动之前，后端服务必须启动、满足所有检查并开始运行。然后前端服务可以启动，满足其检查，并开始运行。

Let’s say we’ve forced the deployment order with the `kubectl wait `command, something like:

假设我们已经使用 `kubectl wait` 命令强制部署顺序，例如：

```
kubectl wait — for=condition=Ready pod/serviceA
```

But when the condition is never met, the next deployment can’t proceed and the process breaks.

但是当条件永远不满足时，下一次部署就无法进行，流程就会中断。

This is a simplistic flow of what a deployment order might look like:

这是一个简单的部署顺序流程：

![img](https://miro.medium.com/max/1260/0*JXePbmNwNkB0g_OB)

*This process cannot move forward until the previous step is complete* 

*在上一步完成之前，此过程无法继续进行*

Since Kubernetes is [self-healing](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy). The standard approach is to let all the services in an application  start concurrently and let the containers crash and restart until they  are all up and running. I have service A and B starting independently  (as a decoupled, stateless cloud-native application should), but for the sake of the user experience, perhaps I could tell the UI (service B) to display a pretty loading message until service A is ready, but the  actual starting up of service B shouldn't be affected by service A.

由于 Kubernetes 是[自我修复](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy)。标准方法是让应用程序中的所有服务同时启动，让容器崩溃并重新启动，直到它们全部启动并运行。我有服务 A 和 B 独立启动（作为一个解耦的、无状态的云原生应用程序应该），但为了用户体验，也许我可以告诉 UI（服务 B)显示一个漂亮的加载消息，直到服务 A ready，但服务 B 的实际启动不应受到服务 A 的影响。

![img](https://miro.medium.com/max/1384/0*ZqfJHyOsWoigya7o)

*Now when the pod crashes, Kubernetes restarts the service until everything  is up and running. If you are stuck in CrashLoopBackOff, it’s worth  checking your code, configuration, or for resource contention.*

*现在当 pod 崩溃时，Kubernetes 会重新启动服务，直到一切正常运行。如果您陷入 CrashLoopBackOff，则值得检查您的代码、配置或资源争用情况。*

Of course, we need to do more than simply rely on self-healing. We need to implement solutions that will handle failures, which are inevitable and will happen. We should anticipate they will happen and lay down  frameworks to respond in a way that helps us avoid downtime and/or data  loss.

当然，我们需要做的不仅仅是依靠自我修复。我们需要实施能够处理故障的解决方案，这些故障是不可避免的并且会发生。我们应该预测它们会发生并制定框架以帮助我们避免停机和/或数据丢失的方式做出响应。

In my hypothetical shopping app, my storefront UI (service B) needs the  inventory (service A) in order to give the user a complete experience. So when there’s a partial failure, like if service A wasn’t available  for a short time or crashed, etc., the system should still be able to  recover from the issue.

在我假设的购物应用程序中，我的店面 UI（服务 B）需要库存（服务 A）才能为用户提供完整的体验。所以当出现部分故障时，比如服务 A 短时间内不可用或崩溃等，系统应该仍然能够从问题中恢复。

Transient faults like these are an ever-present possibility, so to minimize their effects we can implement a [Retry pattern](https://docs.microsoft.com/en-us/azure/architecture/patterns/retry). Retry patterns help improve application stability with strategies like:

像这样的瞬态故障是一种永远存在的可能性，因此为了尽量减少它们的影响，我们可以实施 [重试模式](https://docs.microsoft.com/en-us/azure/architecture/patterns/retry)。重试模式有助于通过以下策略提高应用程序稳定性：

- **Cancel
   **If the fault isn’t transient or if the process is unlikely to be  successful on repeated attempts, then the application should cancel the  operation and report an exception — e.g., authentication failure. Invalid credentials should never work!
- **Retry
   **If the fault is unusual or rare, it could be due to uncommon situations  (e.g., network packet corruption). The application should retry the  request immediately because the same failure is unlikely to reoccur.
- **Retry after delay
   **If the fault is caused by common occurrences like connectivity or busy  failures, it’s best to let any work backlog or traffic clear up before  trying again. The application should wait before retrying the request.
- You could also implement your retry pattern with an [exponential backoff](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/implement-resilient-applications/implement-retries-exponential-backoff) ( exponentially increasing the wait time and setting a maximum retry count).

- **取消
  **如果故障不是暂时的，或者如果重复尝试过程不太可能成功，那么应用程序应该取消操作并报告异常——例如，身份验证失败。无效的凭据不应该工作！
- **重试
  **如果故障异常或罕见，则可能是由于不常见的情况（例如，网络数据包损坏）。应用程序应立即重试请求，因为不太可能再次发生相同的失败。
- **延迟后重试
  **如果故障是由连接或繁忙故障等常见事件引起的，最好在重试之前让任何工作积压或流量清理干净。应用程序在重试请求之前应该等待。
- 您还可以使用 [指数退避](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/implement-resilient-applications/implement-retries-exponential-backoff)实现您的重试模式（以指数方式增加等待时间并设置最大重试次数)。

Implementing a [circuit-breaking pattern](https://istio.io/latest/docs/tasks/traffic-management/circuit-breaking/) is also an important strategy when creating resilient microservice  applications. Like how a circuit breaker in your house will  automatically switch to protect you from extensive damage due to excess  current or short-circuiting, the circuit-breaking pattern provides you a method of writing applications while limiting the impact of unexpected  faults that might take longer to fix, like partial loss of connectivity, or complete failure of a service. In these situations where retrying  won’t work, the application should be able to accept that the failure  has occurred and respond accordingly.

在创建弹性微服务应用程序时，实现[断路器模式](https://istio.io/latest/docs/tasks/traffic-management/circuit-breaking/)也是一个重要的策略。就像您家中的断路器如何自动切换以保护您免受过电流或短路造成的广泛损坏一样，断路器模式为您提供了一种编写应用程序的方法，同时限制了可能需要更长时间才能完成的意外故障的影响修复，例如部分连接丢失或服务完全失败。在这些重试不起作用的情况下，应用程序应该能够接受失败的发生并做出相应的响应。

## **4. Deploying pods without set memory and/or CPU limits**

## **4。在没有设置内存和/或 CPU 限制的情况下部署 Pod**

Resource allocation varies depending on the service, and it can be difficult to  predict what resources a container might require for optimal performance without testing implementation. One service could require a fixed CPU  and memory consumption profile, while another service’s consumption  profile could be dynamic. 

资源分配因服务而异，并且在不测试实现的情况下很难预测容器可能需要哪些资源才能实现最佳性能。一项服务可能需要固定的 CPU 和内存消耗配置文件，而另一项服务的消耗配置文件可能是动态的。

When you deploy pods without careful consideration of memory and CPU limits, this can lead to scenarios of resource contention and unstable  environments. If a container does not have a memory or CPU limit, then  the scheduler sees its memory utilization (and CPU utilization) as zero, so an unlimited number of pods can be scheduled on any node. This can  result in the overcommitment of resources and possible node and kubelet  crashes.

当您部署 Pod 时没有仔细考虑内存和 CPU 限制，这可能会导致资源争用和环境不稳定的情况。如果容器没有内存或 CPU 限制，则调度程序会将其内存利用率（和 CPU 利用率）视为零，因此可以在任何节点上调度无限数量的 Pod。这可能导致资源过度使用以及可能的节点和 kubelet 崩溃。

When the memory limit is not specified for a container, there are a couple  of scenarios that could apply (these also apply to CPU):

当没有为容器指定内存限制时，有几种可能适用的场景（这些也适用于 CPU）：

1. There is no upper bound on the amount of memory a container can use. Thus,  the container could use all of the available memory on its node,  possibly invoking the OOM (out of memory) Killer. An OOM Kill situation  has a greater chance of occurring for a container with no resource  limits.
2. The default memory limit of the namespace (in which the container is  running) is assigned to the container. The cluster administrators can  use a [LimitRange](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#limitrange-v1-core) to specify a default value for the memory limit.
3. 容器可以使用的内存量没有上限。因此，容器可以使用其节点上的所有可用内存，可能会调用 OOM（内存不足）Killer。对于没有资源限制的容器，OOM Kill 情况发生的可能性更大。
4. 命名空间（容器在其中运行）的默认内存限制被分配给容器。集群管理员可以使用 [LimitRange](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#limitrange-v1-core) 来指定内存限制的默认值。

Declaring memory and CPU limits for the containers in your cluster allows you to  make efficient use of the resources available on your cluster’s nodes. This helps the kube-scheduler determine on which node the pod should  reside for most efficient hardware utilization.

为集群中的容器声明内存和 CPU 限制可让您有效利用集群节点上的可用资源。这有助于 kube-scheduler 确定 pod 应该驻留在哪个节点上，以便最有效地利用硬件。

When setting the memory and CPU limits for a container, you should take care not to request more resources than the limit. For pods that have more  than one container, the aggregate resource requests must not exceed the  set limit(s) — otherwise, the pod will never be scheduled.

为容器设置内存和 CPU 限制时，应注意不要请求超过限制的资源。对于具有多个容器的 pod，聚合资源请求不得超过设置的限制——否则，pod 将永远不会被调度。

![img](https://miro.medium.com/max/1400/0*jLGrEtn42OPuqv-K)

*The resource request must not exceed the limit*

*资源请求不得超过限制*

Setting memory and CPU requests below their limits accomplishes two things:

将内存和 CPU 请求设置为低于其限制可以完成两件事：

1. The pod can make use of memory/CPU when it is available, leading to bursts of activity.
2. During a burst, the pod is limited to a reasonable amount of memory/CPU.

1. Pod 可以在内存/CPU 可用时使用它，从而导致活动的爆发。
2. 在爆发期间，pod 被限制在合理数量的内存/CPU 内。

The best practice is to keep the CPU request at one core or below, and then use [ReplicaSets](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) to scale it out, which gives the system flexibility and reliability.

最佳做法是将 CPU 请求保持在一个核心或以下，然后使用 [ReplicaSets](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) 将其横向扩展，这给系统灵活性和可靠性。

What happens when you have different teams competing for resources when  deploying containers in the same cluster? If the process exceeds the  memory limit, then it will be terminated, while if it exceeds the CPU  limit, the process will be throttled (resulting in worse performance).

在同一集群中部署容器时，如果不同团队争夺资源，会发生什么？如果进程超过内存限制，那么它会被终止，而如果超过 CPU 限制，进程将被限制（导致性能变差）。

You can control resource limits via [resource quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/) and [LimitRange](https://kubernetes.io/docs/concepts/policy/limit-range/) in the [namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) settings. These settings help account for containers deployments without limits or with high resource requests.

您可以通过 [resource quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/) 和 [LimitRange](https://kubernetes.io/docs/concepts/policy/limit-range/) 在 [namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) 设置中。这些设置有助于无限制或具有高资源请求的容器部署。

Setting hard resource limits might not be the best choice for your needs. Another option is to use the recommendation mode in the [Vertical Pod autoscaler ](https://cloud.google.com/kubernetes-engine/docs/concepts/verticalpodautoscaler)resource.

设置硬资源限制可能不是满足您需求的最佳选择。另一种选择是使用 [Vertical Pod autoscaler ](https://cloud.google.com/kubernetes-engine/docs/concepts/verticalpodautoscaler) 资源中的推荐模式。

## **5. Pulling the **`latest `**tag in containers in production** 

## **5。在生产中的容器中提取“**最新” `**标签**

[Using the ](https://kubernetes.io/docs/concepts/containers/images/#image-names)`latest`[ tag](https://kubernetes.io/docs/concepts/containers/images/#image-names) is considered bad practice, especially in production. Pods unexpectedly crash for all sorts of reasons, so they can pull down images at any  time. Unfortunately, the `latest `tag is not very descriptive when it comes to determining when the build  broke. What version of the image was running? When was the last time it  was working? This is especially bad in production since you need to be  able to get things back up and running with minimal downtime.

[使用](https://kubernetes.io/docs/concepts/containers/images/#image-names)`latest`[标签](https://kubernetes.io/docs/concepts/containers/images/#image-names) 被认为是不好的做法，尤其是在生产中。 Pod 因各种原因意外崩溃，因此它们可以随时下拉图像。不幸的是，在确定构建何时中断时，`latest` 标签不是很具有描述性。运行的是哪个版本的映像？上一次工作是什么时候？这在生产中尤其糟糕，因为您需要能够以最少的停机时间恢复并运行。

![img](https://miro.medium.com/max/1400/0*9paNbPUoky-4F3uX)

*You shouldn’t use the* `*latest* `*tag in production.*

*你不应该在生产中使用*`*latest*`*标签。*

By default, the `imagePullPolicy` is set to `Always` and will always pull down the image when it restarts. If you don’t specify a tag, Kubernetes will default to `latest`. However, a deployment will only be updated in the event of a crash  (when the pod pulls down the image on restart) or if the deployment  pod's template (`.spec.template`) is [changed](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#updating-a-deployment). See [this forum discussion](https://discuss.kubernetes.io/t/use-latest-image-tag-to-update-a-deployment/2929) for an example of `latest` not working as intended in development .

默认情况下，`imagePullPolicy` 设置为 `Always`，并且在重新启动时将始终下拉图像。如果不指定标签，Kubernetes 将默认为“latest”。但是，部署只会在发生崩溃（当 pod 在重启时拉下图像时）或部署 pod 的模板（`.spec.template`）[已更改](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#updating-a-deployment)。请参阅 [此论坛讨论](https://discuss.kubernetes.io/t/use-latest-image-tag-to-update-a-deployment/2929) 以了解“最新”在开发中无法按预期工作的示例.

Even if you’ve changed the `imagePullPolicy` to another value than `Always`, your pod will still pull an image if it needs to restart (whether it’s  because of a crash or deliberate reboot). If you use versioning and set  the `imagePullPolicy` with a meaningful tag, like v1.4.0, then you can [roll back](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#rolling-back-to-a-previous-revision) to the most recent stable version and more easily troubleshoot when and where something went wrong in your code. You can read more about best  practices for versioning in the [Semantic Versioning Specification](https://semver.org/) and [GCP Best Practices](https://semver.org/).

即使您将 `imagePullPolicy` 更改为除“Always”以外的其他值，如果需要重新启动（无论是由于崩溃还是故意重启），您的 pod 仍会拉取映像。如果您使用版本控制并使用有意义的标签设置 `imagePullPolicy`，例如 v1.4.0，那么您可以[回滚](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#rolling-back-to-a-previous-revision) 到最新的稳定版本，更轻松地解决代码中出现问题的时间和地点。您可以在 [语义版本控制规范](https://semver.org/) 和 [GCP 最佳实践](https://semver.org/) 中阅读有关版本控制最佳实践的更多信息。

In addition to using specific and meaningful Docker tags, you should also  remember that containers are stateless and immutable. They are also  meant to be ephemeral (and you should store any data outside containers  in persistent storage). Once you spin up a container, you should not  modify it: no patches, no updates, no configuration changes. When you  need to update a configuration, you should deploy a new container with  the updated config.

除了使用特定且有意义的 Docker 标签外，您还应该记住容器是无状态且不可变的。它们也意味着是短暂的（并且您应该将容器外的任何数据存储在持久存储中）。一旦你启动了一个容器，你就不应该修改它：没有补丁，没有更新，没有配置更改。当您需要更新配置时，您应该使用更新后的配置部署一个新容器。

![img](https://miro.medium.com/max/1400/0*1dTSSbrBOir4Kw3_)

[*Docker immutability, taken from Best Practices for Operating Containers*](https://cloud.google.com/solutions/best-practices-for-operating-containers)*.*

[*Docker 不变性，摘自操作容器的最佳实践*](https://cloud.google.com/solutions/best-practices-for-operating-containers)*.*

This immutability allows for safer and repeatable deployments. You can also  more easily roll back if you need to redeploy the old image. By keeping  your Docker images and container immutable, you are able to deploy the  same container image in every single environment. See Anti-pattern 1 to  read about externalizing your configuration data to keep your images  immutable.

这种不变性允许更安全和可重复的部署。如果您需要重新部署旧映像，您还可以更轻松地回滚。通过保持 Docker 镜像和容器不可变，您可以在每个环境中部署相同的容器镜像。请参阅反模式 1 以了解如何将配置数据外部化以保持图像不可变。

![img](https://miro.medium.com/max/1400/0*dtF2W7jkk_-D-bok)

*We can roll back to the previous stable version while we troubleshoot.*

*我们可以在排除故障的同时回滚到以前的稳定版本。*

## **6. Deploying new updates/fixes by killing pods so they pull the new Docker images during the restart process**

## **6。通过杀死 pod 来部署新的更新/修复，以便它们在重启过程中提取新的 Docker 镜像**

Like relying on the latest tag to pull updates, relying on killing pods to  roll out new updates is bad practice since you’re not versioning your  code. If you are killing pods to pull updated Docker images in  production, don’t do it. Once a version has been released in production, it should never be overwritten. If something breaks, then you won’t  know where or when things went wrong and how far back to go when you  need to roll back the code while you troubleshoot. 

就像依靠最新的标签来拉取更新一样，依靠杀死 pod 来推出新的更新是一种不好的做法，因为你没有对代码进行版本控制。如果您正在杀死 pod 以在生产中拉取更新的 Docker 镜像，请不要这样做。一旦一个版本在生产中发布，它就不应该被覆盖。如果出现问题，那么您将不知道哪里或何时出了问题，以及在进行故障排除时需要回滚代码的时间。

Another problem is that restarting the container to pull a new Docker image  doesn’t always work. “A Deployment’s rollout is triggered if and only if the Deployment’s Pod template (that is, `.spec.template`) is changed, for example if the labels or container images of the  template are updated. Other updates, such as scaling the Deployment, do  not trigger a rollout.”

另一个问题是重新启动容器以拉取新的 Docker 镜像并不总是有效。 “当且仅当 Deployment 的 Pod 模板（即`.spec.template`）发生更改时，才会触发 Deployment 的推出，例如，如果模板的标签或容器映像已更新。其他更新，例如扩展部署，不会触发推出。”

<iframe src="https://betterprogramming.pub/media/023ced411f5076fa92be0e796b90be38" allowfullscreen="" title="deployment-1.yaml" class="tuv lc aj" scrolling="auto" width="680" height=" 479" frameborder="0"></iframe>

<iframe src="https://betterprogramming.pub/media/023ced411f5076fa92be0e796b90be38" allowfullscreen="" title="deployment-1.yaml" class="tuv lc aj" scrolling="auto" width="680" height="第479话

You have to modify the `.spec.template` to trigger a deployment.

你必须修改 `.spec.template` 来触发部署。

The correct way to update your pods to pull new Docker images is to [version](https://semver.org/) (or increment for fixes/patches) your code and then modify the deployment spec to reflect a meaningful tag (not `latest`, see Anti-pattern 5) for further discussion on that, but something like  v1.4.0 for a new release or v1.4.1 for a patch). Kubernetes will then  trigger an upgrade with zero downtime.

更新 pod 以拉取新 Docker 镜像的正确方法是 [version](https://semver.org/)（或增加修复/补丁）您的代码，然后修改部署规范以反映有意义的标签（不是`latest`，请参阅反模式5) 以进一步讨论，但类似于 v1.4.0 用于新版本或 v1.4.1 用于补丁)。然后 Kubernetes 将触发零停机升级。

1. Kubernetes starts a new pod with the new image.
2. Waits for health checks to pass.
3. Deletes the old pod.

1. Kubernetes 使用新镜像启动一个新 pod。
2. 等待健康检查通过。
3. 删除旧的 pod。

## **7. Mixing both production and non-production workloads in the same cluster**

## **7。在同一集群中混合生产和非生产工作负载**

Kubernetes supports a [namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) feature, which enables users to manage different environments (virtual  clusters) within the same physical cluster. Namespaces can be seen as a  cost-effective way of managing different environments on a single  physical cluster. For example, you could run staging and production  environments in the same cluster and save resources and money. However,  there’s a big gap between running Kubernetes in development and running  Kubernetes in production.

Kubernetes 支持 [namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)特性，它使用户能够在同一个物理集群内管理不同的环境（虚拟集群)。命名空间可以被视为在单个物理集群上管理不同环境的一种经济高效的方式。例如，您可以在同一个集群中运行临时环境和生产环境，从而节省资源和资金。但是，在开发中运行 Kubernetes 与在生产中运行 Kubernetes 之间存在很大差距。

There are a lot of factors to consider when you mix your production and  non-production workloads on the same cluster. For one, you would have to consider resource limits to make sure the performance of your  production environment isn't compromised (a common practice one might  see is setting no quota on the production namespace and a quota on any  non-production namespace(s) ).

在同一集群上混合生产和非生产工作负载时，需要考虑很多因素。一方面，您必须考虑资源限制以确保您的生产环境的性能不会受到影响（人们可能会看到的一种常见做法是在生产命名空间上不设置配额，而在任何非生产命名空间上设置配额） ）。

You would also need to consider isolation. Developers require a lot more  access and permissions than in production, which you would want locked  down as much as possible. While namespaces are hidden from each other,  they are not fully isolated by default. That means your apps in a dev  namespace could call apps in test, staging, or production (or vice  versa), which is not considered good practice. Of course, you could use  NetworkPolicies to set rules to isolate the namespaces.

您还需要考虑隔离。开发人员需要比生产中更多的访问和权限，您希望尽可能地锁定它们。虽然命名空间彼此隐藏，但默认情况下它们并未完全隔离。这意味着 dev 命名空间中的应用程序可以调用测试、暂存或生产中的应用程序（反之亦然），这被认为是不好的做法。当然，您可以使用 NetworkPolicies 设置规则来隔离命名空间。

However, thoroughly testing resource limits, performance, security, and  reliability is time-consuming, so running production workloads in the  same cluster as non-production workloads is not advised. Rather than  mixing production and non-production workloads in the same cluster, use  separate clusters for development/test/production — you’ll have better  isolation and security that way. You should also automate as much as you can for CI/CD and promotion to reduce the chance for human error. Your  production environment needs to be as solid as possible.

但是，彻底测试资源限制、性能、安全性和可靠性非常耗时，因此不建议在与非生产工作负载相同的集群中运行生产工作负载。与其在同一个集群中混合生产和非生产工作负载，不如使用单独的集群进行开发/测试/生产——这样您将获得更好的隔离和安全性。您还应该尽可能多地自动化 CI/CD 和推广，以减少人为错误的机会。您的生产环境需要尽可能可靠。

## **8. Not using blue/green or canaries for mission-critical deployments**

## **8。不使用蓝/绿或金丝雀进行关键任务部署**

Many modern applications have frequent deployments, ranging from several  changes within a month to multiple deployments in a single day. This is  certainly achievable with microservice architecture since the different  components can be developed, managed, and released on different cycles  as long as they work together to perform seamlessly. And of course,  keeping applications up 24/7 is obviously important when rolling out  updates.

许多现代应用程序都有频繁的部署，从一个月内的多次更改到一天内的多次部署。这当然可以通过微服务架构实现，因为不同的组件可以在不同的周期内开发、管理和发布，只要它们协同工作以无缝执行。当然，在推出更新时，保持应用程序 24/7 全天候运行显然很重要。

The default rolling update of Kubernetes is not always enough. A common  strategy to perform updates is to use the default Kubernetes [rolling update](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#rolling-update-deployment) feature:

Kubernetes 的默认滚动更新并不总是足够的。执行更新的常见策略是使用默认的 Kubernetes [滚动更新](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#rolling-update-deployment) 功能：

```
.spec.strategy.type==RollingUpdate
```

where you can set the `maxUnavailable` (percentage or number of pods unavailable) and `maxSurge` fields (optional) to control the rolling update process. When  implemented properly, rolling updates allow a gradual update with zero  downtime as the pods are incrementally updated. Here’s an [example](https://medium.com/platformer-blog/enable-rolling-updates-in-kubernetes-with-zero-downtime-31d7ec388c81) of how one team updated their applications with zero downtime with rolling updates.

您可以在其中设置 `maxUnavailable`（不可用的 pod 的百分比或数量）和 `maxSurge` 字段（可选）来控制滚动更新过程。如果实施得当，滚动更新可以实现零停机时间的渐进更新，因为 pod 是增量更新的。这是一个[示例](https://medium.com/platformer-blog/enable-rolling-updates-in-kubernetes-with-zero-downtime-31d7ec388c81)，说明一个团队如何通过滚动更新以零停机时间更新他们的应用程序。

However, once you’ve updated your deployment to the next version, it’s not  always easy to go back. You should have a plan in place to roll it back  in case it breaks in production. When your pod is updated to the next  version, the deployment will create a new ReplicaSet. While Kubernetes  will store previous ReplicaSets (by default, it’s ten, but you could  change that with `spec.revisionHistoryLimit`). The ReplicaSets are saved under names such as `app6ff34b8374` in random order, and you won’t find a reference to the ReplicaSets in the deployment app YAML. You could find it with:

但是，一旦您将部署更新到下一个版本，返回并不总是那么容易。您应该制定计划以在生产中断时将其回滚。当您的 pod 更新到下一个版本时，部署将创建一个新的 ReplicaSet。虽然 Kubernetes 将存储以前的 ReplicaSet（默认情况下，它是十个，但您可以使用 `spec.revisionHistoryLimit` 更改它）。 ReplicaSet 以随机顺序保存在诸如“app6ff34b8374”之类的名称下，并且您不会在部署应用程序 YAML 中找到对 ReplicaSet 的引用。你可以找到它：

```
ReplicaSet.metatada.annotation
```

and inspect the revision with:

并检查修订：

```
kubectl get replicaset app-6ff88c4474 -o yaml
```

to find the revision number. This gets complicated because the rollout  history doesn’t keep a log unless you leave a note in the YAML resource  (which you could do with the `—record` flag:

找到修订号。这变得很复杂，因为除非您在 YAML 资源中留下注释（您可以使用 `-record` 标志来完成），否则推出历史不会保留日志：

```
$kubectl rollout history deployment/appREVISION CHANGE-CAUSE1 kubectl create — filename=deployment.yaml — record=true
2 kubectl apply — filename=deployment.yaml —record=true
```

When you have dozens, hundreds, or even thousands of deployments all going  through updates simultaneously, it’s difficult to keep track of them all at once. And if your stored revisions all contain the same regression,  then your production environment is not going to be in good shape! You  can read more in detail about using rolling updates [in this article](https://learnk8s.io/kubernetes-rollbacks#:~:text=In Kubernetes%2C rolling updates are,bring newer Pod in incrementally.&text=You have a Service and,three replicas on version 1.0.).

当您有数十个、数百个甚至数千个部署同时进行更新时，很难同时跟踪它们。如果您存储的修订都包含相同的回归，那么您的生产环境将不会处于良好状态！您可以阅读有关使用滚动更新的更多详细信息 [在本文中](https://learnk8s.io/kubernetes-rollbacks#:~:text=在 Kubernetes%2C 中，滚动更新是，以增量方式引入更新的 Pod。&text=You在版本 1.0 上有一个服务和三个副本。)。

Some other problems are:

其他一些问题是：

- Not all applications are capable of concurrently running multiple versions.
- Your cluster could run out of resources in the middle of the update, which could break the whole process.

- 并非所有应用程序都能够同时运行多个版本。
- 您的集群可能会在更新过程中耗尽资源，这可能会中断整个过程。

These are all very frustrating and stressful issues to run into when in a production environment.

这些都是在生产环境中遇到的非常令人沮丧和压力很大的问题。

Alternative ways to more reliably update deployments include:

更可靠地更新部署的替代方法包括：

**Blue/green (red/black) deployment**
With blue/green, a full set of both the old and new instances exist  simultaneously. Blue is the live version, and the new version is  deployed to the green replica. When the green environment has passed its tests and verifications, a load balancer simply flips the traffic to  the green, which becomes the blue environment, and the old version  becomes the green version. Since we have two full versions being  maintained, performing a rollback is simple — all you need to do is  switch back the load balancer.

**蓝/绿（红/黑）部署**
使用蓝/绿，一整套旧的和新的实例同时存在。蓝色是live版本，新版本部署到绿色副本。当绿色环境通过其测试和验证时，负载均衡器只需将流量翻转到绿色，绿色就变成蓝色环境，旧版本变成绿色版本。由于我们维护了两个完整版本，因此执行回滚很简单——您需要做的就是切换回负载均衡器。

![img](https://miro.medium.com/max/1228/0*LQzr9w-ADoabf8rJ)

*The load balancer flips between blue and green to set the active version. From* [*Continuous Deployment Strategies with Kubernetes*](https://codefresh.io/kubernetes-tutorial/continuous-deployment-strategies-kubernetes-2/)*.*

*负载均衡器在蓝色和绿色之间切换以设置活动版本。来自* [*Kubernetes 的持续部署策略*](https://codefresh.io/kubernetes-tutorial/continuous-deployment-strategies-kubernetes-2/)*.*

Additional advantages include:

其他优势包括：

- Since we never deploy directly to production, it’s pretty low-stress when we change green to blue.
- Traffic redirection occurs immediately, so there’s no downtime.
- There can be extensive testing done to reflect actual production prior to the switch. (As stated before, a development environment is very different  from production.)

- 由于我们从不直接部署到生产环境，所以当我们将绿色变为蓝色时压力非常低。
- 流量重定向立即发生，因此没有停机时间。
- 可以进行大量测试以反映切换前的实际生产情况。 （如前所述，开发环境与生产环境非常不同。）

Kubernetes does not include blue/green deployments as one of its native toolings. You can read more about how to implement blue/green into your CI/CD  automation [in this tutorial](https://codefresh.io/kubernetes-tutorial/fully-automated-blue-green-deployments-kubernetes-codefresh/) .

Kubernetes 不包括蓝/绿部署作为其原生工具之一。你可以阅读更多关于如何在你的 CI/CD 自动化中实现蓝/绿 [在本教程中](https://codefresh.io/kubernetes-tutorial/fully-automated-blue-green-deployments-kubernetes-codefresh/) .

### Canary Releases 

###  金丝雀发布

**Canary** releases allow us to test for potential problems and meet key metrics  before impacting the entire production system/user base. We “test in  production” by deploying directly to the production environment, but  only to a small subset of users. You can choose routing to be  percentage-based or driven by region/user location, the type of client,  and billing properties. Even when deploying to a small subset, it’s  important to carefully monitor application performance and measure  errors — these metrics define a quality threshold. If the application  behaves as expected, we start transferring more of the new version  instances to support more traffic.

金丝雀版本允许我们在影响整个生产系统/用户群之前测试潜在问题并满足关键指标。我们通过直接部署到生产环境来“在生产中测试”，但只部署到一小部分用户。您可以选择基于百分比或由区域/用户位置、客户端类型和计费属性驱动的路由。即使在部署到一个小子集时，仔细监控应用程序性能和测量错误也很重要——这些指标定义了一个质量阈值。如果应用程序按预期运行，我们将开始传输更多新版本实例以支持更多流量。

![img](https://miro.medium.com/max/1400/0*pzALOq2I9-ws85X9)

*The load balancer gradually releases the new version into production. From* [*Continuous Deployment Strategies with Kubernetes*](https://codefresh.io/kubernetes-tutorial/continuous-deployment-strategies-kubernetes-2/)*.*

*负载均衡器逐步将新版本发布到生产环境中。来自* [*Kubernetes 的持续部署策略*](https://codefresh.io/kubernetes-tutorial/continuous-deployment-strategies-kubernetes-2/)*.*

Other advantages include:

其他优势包括：

- Observability
- Ability to test on production traffic (getting a true production-like experience in development is hard)
- Ability to release a version to a small subset of users and get real feedback before a larger release
- Fail fast. Since we deploy straight into production, we can fail fast (i.e., revert immediately) if it breaks, and it affects only a subset rather  than the whole community.

- 可观察性
- 能够测试生产流量（在开发中获得真正的类似生产的体验很难）
- 能够向一小部分用户发布一个版本，并在更大的发布之前获得真实的反馈
- 快速失败。由于我们直接部署到生产环境中，如果它发生故障，我们可以快速失败（即立即恢复），并且它只影响一个子集而不是整个社区。

## **9. Not having metrics in place to understand if a deployment was successful or not**

## **9。没有指标来了解部署是否成功**

Your health checks need application support.

您的健康检查需要应用程序支持。

You can leverage Kubernetes to accomplish many tasks in container orchestration:

您可以利用 Kubernetes 来完成容器编排中的许多任务：

- Controlling resource consumption by an application or team (namespace, CPU/mem,  limits) stopping an app from consuming too many resources
- Load balancing across different app instances, moving application instances  from one host to another if there is a resource shortage or if the host  dies
- Self-healing — restarting containers if they crash
- Automatically leveraging additional resources if a new host is added to the cluster
- And more

- 控制应用程序或团队的资源消耗（命名空间、CPU/内存、限制）阻止应用程序消耗过多资源
- 跨不同应用程序实例的负载平衡，如果资源短缺或主机死机，则将应用程序实例从一台主机移动到另一台主机
- 自我修复——如果容器崩溃，则重新启动容器
- 如果向集群添加新主机，则自动利用额外资源
- 和更多

So sometimes it’s easy to forget about metrics and monitoring. However, a  successful deployment is not the end of your ops work. It’s better to be proactive and prepare for unexpected surprises. There are still a lot  more layers to monitor, and the dynamic nature of K8s makes it tough to  troubleshoot. For example, if you’re not closely watching your resource  available, the automatic rescheduling of pods could cause capacity  issues, and your app might crash or never deploy. This would be  especially unfortunate in production, as you wouldn’t know unless  someone filed a bug report or if you happened to check on it. Eep!

因此，有时很容易忘记指标和监控。但是，成功的部署并不是您的运维工作的结束。最好积极主动并为意外惊喜做好准备。还有很多层需要监控，而 K8s 的动态特性使得故障排除变得困难。例如，如果您没有密切关注可用资源，Pod 的自动重新调度可能会导致容量问题，并且您的应用程序可能会崩溃或永远不会部署。这在生产中尤其不幸，因为除非有人提交了错误报告或者您碰巧检查了它，否则您不会知道。哎呀！

Monitoring presents its own set of challenges: There are a lot of layers to watch, and there's a need to “[maintain a reasonably low maintenance burden on the engineers](https://landing.google.com/sre/sre-book/chapters/practical-alerting/).” When an application running on Kubernetes hits a snag, there are many  logs, data, and components to investigate, especially when there are  multiple microservices involved with the issue versus in traditional  monolithic architecture, where everything is output to a few logs. 

监控提出了自己的一系列挑战：有很多层需要观察，并且需要“[保持工程师的合理低维护负担](https://landing.google.com/sre/sre-book/章节/实用警报/)。”当运行在 Kubernetes 上的应用程序遇到问题时，有许多日志、数据和组件需要调查，尤其是当问题涉及多个微服务时，与传统的单体架构相比，所有内容都输出到几个日志中。

Insights on your application behavior, like how an application performs, helps  you continuously improve. You also need a pretty holistic view of the  containers, pods, services, and the cluster as a whole. If you can  identify how an application is using its resources, then you can use  Kubernetes to better detect and remove bottlenecks. To get a full view  of the application, you would need to use an application performance  monitoring solution like [Prometheus](https://prometheus.io/),[Grafana](https://grafana.com/), [New Relic](https://newrelic.com/), or [Cisco AppDynamics](https://www.appdynamics.com/appd-campaigns/?utm_source=adwords&utm_medium=ppc&utm_campaign=brand&gclid=CjwKCAjwydP5BRBREiwA-qrCGo92C606PzpGx6nOZdhkIs8WcxHadyb-gYDCeUfofm3hBSgeTTAW8BoC7DQQAvD_BwE), among many others.

洞察您的应用程序行为，例如应用程序的执行方式，可帮助您不断改进。您还需要对容器、pod、服务和整个集群有一个非常全面的了解。如果您可以确定应用程序如何使用其资源，那么您可以使用 Kubernetes 更好地检测和消除瓶颈。要全面了解应用程序，您需要使用应用程序性能监控解决方案，例如 [Prometheus](https://prometheus.io/)、[Grafana](https://grafana.com/)、[New Relic](https://newrelic.com/)，或[思科AppDynamics](https://www.appdynamics.com/appd-campaigns/?utm_source=adwords&utm_medium=ppc&utm_campaign=brand&gclid=CjwKCAjwydP5BRBREiwA-qrCGo92C606PzpGx6nOZdhkIs8WcxHadyb-gYDCeUfofm3hBSgeTTAW8BoC7DQQAvD_BwE)，除很多其他的。

Whether or not you decide to use a monitoring solution, these are the key  metrics that the Kubernetes documentation recommends you track closely:

无论您是否决定使用监控解决方案，以下是 Kubernetes 文档建议您密切跟踪的关键指标：

- Running pods and their deployments
- Resource metrics: CPU, memory usage, disk I/O
- Container-native metrics
- Application metrics

- 运行 Pod 及其部署
- 资源指标：CPU、内存使用率、磁盘 I/O
- 容器原生指标
- 应用指标

## **10. Cloud vendor lock-in: Locking yourself into an IaaS provider’s Kubernetes or serverless computing services**

## **10。云供应商锁定：将自己锁定在 IaaS 提供商的 Kubernetes 或无服务器计算服务中**

There are multiple types of lock-ins (Martin Fowler wrote a great [article](https://martinfowler.com/articles/oss-lockin.html), if you want to read more), but vendor lock-in negates the primary value of deploying to the cloud: container flexibility. It’s true that  choosing the right cloud provider is not an easy decision. Each provider has its own interfaces, open APIs, and proprietary specifications and  standards. Additionally, one provider might suit your needs better than  the others only for your business needs to unexpectedly change.

有多种类型的锁定（Martin Fowler 写了一篇很棒的[文章](https://martinfowler.com/articles/oss-lockin.html)，如果你想阅读更多)，但供应商锁定否定了部署到云的主要价值：容器灵活性。诚然，选择合适的云提供商并非易事。每个提供商都有自己的接口、开放的 API 以及专有的规范和标准。此外，一个提供商可能比其他提供商更适合您的需求，但仅当您的业务需求发生意外变化时。

Fortunately, containers are platform-agnostic and portable, and all the major  providers have a Kubernetes foundation, which is cloud-agnostic. You  don’t have to re-architect or rewrite your application code when you  need to move workloads between clouds, so you shouldn’t need to lock  yourself into a cloud provider because you can’t “lift and shift.”

幸运的是，容器与平台无关且可移植，并且所有主要供应商都有一个与云无关的 Kubernetes 基础。当您需要在云之间移动工作负载时，您不必重新架构或重写应用程序代码，因此您不需要将自己锁定在云提供商中，因为您无法“提升和转移”。

Here is a list of things you should consider to ensure you can be flexible to prevent or minimize vendor lock-in.

以下是您应该考虑的事项列表，以确保您可以灵活地防止或最大限度地减少供应商锁定。

**First,** [**housekeeping**](https://www.techrepublic.com/article/5-ways-to-avoid-vendor-lock-in/)**: Read the fine print**

**首先，** [**家政**](https://www.techrepublic.com/article/5-ways-to-avoid-vendor-lock-in/)**：阅读细则**

Negotiate entry and exit strategies. Many vendors make it easy to start — and get you hooked. This might include incentives like free trials or credits,  but these costs could rapidly increase as you scale up.

协商进入和退出策略。许多供应商让您轻松上手 — 并让您着迷。这可能包括免费试用或积分等激励措施，但随着规模扩大，这些成本可能会迅速增加。

Check for things like auto-renewal, early termination fees, and if the  provider will help with things like deconversion when migrating to  another vendor and SLAs associated with exit.

检查诸如自动续订、提前终止费用之类的内容，以及提供商是否会在迁移到另一个供应商和与退出相关的 SLA 等问题上提供帮助。

**Architect/design your applications such that they can run on any cloud**

**架构/设计您的应用程序，使其可以在任何云上运行**

If you’re already developing for the cloud and using cloud-native  principles, then most likely your application code should be easy to  lift and shift. It’s the things surrounding the code that potentially  lock you into a cloud vendor. For example, you could:

如果您已经在为云开发并使用云原生原则，那么您的应用程序代码很可能应该易于提升和转移。围绕代码的事情可能会将您锁定在云供应商中。例如，您可以：

- Check that your services and features (like databases, APIs, etc) used by your application are portable.
- Check if your deployment and provisioning scripts are cloud-specific. Some  clouds have their own native or recommended automation tools that may  not translate easily to other providers. There are many tools that can  be used to assist with cloud infrastructure automation and are  compatible with many of the major cloud providers, like [Puppet](https://puppet.com/), [Ansible](https://www.ansible.com/), and [Chef](https://www.chef.io/), to name a few. This [blog](https://www.ibm.com/cloud/blog/chef-ansible-puppet-terraform) has a handy chart that compares characteristics of common tools. 

- 检查您的应用程序使用的服务和功能（如数据库、API 等）是否可移植。
- 检查您的部署和供应脚本是否特定于云。一些云有自己的原生或推荐的自动化工具，这些工具可能无法轻松转换到其他提供商。有许多工具可用于协助实现云基础设施自动化，并且与许多主要的云提供商兼容，例如 [Puppet](https://puppet.com/)、[Ansible](https://www.ansible.com/) 和 [Chef](https://www.chef.io/)，仅举几例。这个[博客](https://www.ibm.com/cloud/blog/chef-ansible-puppet-terraform) 有一个方便的图表，比较了常用工具的特征。

- Check if your DevOps environment, which typically includes Git and CI/CD, can run in any cloud. For example, many clouds have their own specific  CI/CD tools, like [IBM Cloud Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery/pipeline_about.html#deliverypipeline_about), [Azure CI/CD]( https://docs.microsoft.com/en-us/azure/architecture/solution-ideas/articles/azure-devops-continuous-integration-and-continuous-deployment-for-azure-web-apps), or [AWS Pipelines](https://aws.amazon.com/getting-started/projects/set-up-ci-cd-pipeline/), that might require extra work to port over to another cloud vendor. Instead, you could use something like [Codefresh](https://codefresh.io/?utm_source=google&utm_medium=cpc&utm_campaign=brand-search&utm_term=codefresh&gclid=CjwKCAjwmrn5BRB2EiwAZgL9ogt93s-yRXqtsuBb65KqDJftz-6biYdFJ1LPdwMlV6y9pLPlRy4yxhoCr10QAvD_BwE), full CI/CD solutions that have great support for Docker and Kubernetes  and integrates with many other popular tools. There are also myriad  other solutions, some CI or CD, or both, like [GitLab](https://about.gitlab.com/blog/2019/07/15/finding-the-right-ci-cd/),[Bamboo](https://www.atlassian.com/software/bamboo), [Jenkins](https://www.jenkins.io/),[Travis](https://www.jenkins.io/) , etc.
- Check if your testing process will need to be changed between providers.

- 检查您的 DevOps 环境（通常包括 Git 和 CI/CD）是否可以在任何云中运行。例如，许多云都有自己特定的 CI/CD 工具，例如 [IBM Cloud Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery/pipeline_about.html#deliverypipeline_about)、[Azure CI/CD]( https://docs.microsoft.com/en-us/azure/architecture/solution-ideas/articles/azure-devops-continuous-integration-and-continuous-deployment-for-azure-web-apps），或[AWS管道](https://aws.amazon.com/getting-started/projects/set-up-ci-cd-pipeline/)，这可能需要额外的工作才能移植到另一个云供应商。相反，你可以使用类似[Codefresh（https://codefresh.io/?utm_source=google&utm_medium=cpc&utm_campaign=brand-search&utm_term=codefresh&gclid=CjwKCAjwmrn5BRB2EiwAZgL9ogt93s-yRXqtsuBb65KqDJftz-6biYdFJ1LPdwMlV6y9pLPlRy4yxhoCr10QAvD_BwE)，具有对码头工人的大力支持完整的CI / CD解决方案和 Kubernetes 并与许多其他流行工具集成。还有无数其他解决方案，一些 CI 或 CD，或两者兼而有之，例如 [GitLab](https://about.gitlab.com/blog/2019/07/15/finding-the-right-ci-cd/)，[Bamboo](https://www.atlassian.com/software/bamboo)、[Jenkins](https://www.jenkins.io/)、[Travis](https://www.jenkins.io/) ， 等等。
- 检查您的测试过程是否需要在提供者之间更改。

**You could also choose to follow a** [**multicloud strategy**](https://www.techrepublic.com/article/multicloud-the-smart-persons-guide/)

**您也可以选择遵循** [**多云策略**](https://www.techrepublic.com/article/multicloud-the-smart-persons-guide/)

With a multicloud strategy, you can pick and choose services from different  cloud providers that best the type of application(s) you are hoping to  deliver. When you plan a multicloud deployment, you should keep  interoperability in careful consideration.

借助多云战略，您可以从不同的云提供商处挑选最适合您希望交付的应用程序类型的服务。当您计划多云部署时，您应该仔细考虑互操作性。

# Summary

#  概括

Kubernetes is really popular, but it’s hard to get started with, and there are a  lot of practices in traditional development that don’t translate to  cloud-native development.

Kubernetes 真的很受欢迎，但是很难上手，而且传统开发中有很多实践并没有转化为云原生开发。

In this article, we’ve looked at:

在本文中，我们研究了：

1. Putting the configuration file inside/alongside the Docker image: **Externalise your configuration data. You can use ConfigMaps and Secrets or something similar.**
2. Not using Helm or other kinds of templating: **Use Helm or Kustomize to streamline your container orchestration and reduce human error.**
3. Deploying things in a specific order: **Applications shouldn’t crash because a dependency isn’t ready. Utilize Kubernetes’s  self-healing mechanism and implement retries and circuit breakers.**
4. Deploying pods without set memory and/or CPU limits: **You should consider setting memory and CPU limits to reduce the risk of  resource contention, especially when sharing the cluster with others.**
5. Pulling the `latest` tag in containers in production: **Never use** `**latest**`**. Always use something meaningful, like v1.4.0/according to** [**Semantic Versioning Specification**](https://semver.org/)**, and employ immutable Docker images.**
6. Deploying new updates/fixes by killing pods so they pull the new Docker images during the restart process: **Version your code so you can better manage your releases.**
7. Mixing both production and non-production workloads in the same cluster: **Run your production and non-production workloads in separate clusters if  you can. This reduces risk to your production environment from resource  contention and accidental environment cross-over.**
8. Not using blue/green or canaries for mission-critical deployments (the  default rolling update of Kubernetes is not always enough): **You should consider blue/green deployment or canary releases for less  stress in production and more meaningful production results. **
9. Not having metrics in place to understand if a deployment was successful or not (your health checks need application support): **You should make sure to monitor your deployment to avoid any surprises. You could use a tool like Prometheus, Grafana, New Relic, or Cisco  AppDynamics to help you gain better insights on your deployments.** 



1. 将配置文件放在 Docker 映像内部/旁边：**外部化您的配置数据。您可以使用 ConfigMaps 和 Secrets 或类似的东西。**
2. 不使用 Helm 或其他类型的模板：**使用 Helm 或 Kustomize 来简化您的容器编排并减少人为错误。**
3. 以特定顺序部署事物：**应用程序不应因为依赖项尚未准备好而崩溃。利用 Kubernetes 的自愈机制并实现重试和断路器。**
4. 在没有设置内存和/或 CPU 限制的情况下部署 Pod：**您应该考虑设置内存和 CPU 限制以降低资源争用的风险，尤其是在与他人共享集群时。**
5. 在生产中的容器中拉取 `latest` 标签：**永远不要使用** `**latest**`**。总是使用一些有意义的东西，比如 v1.4.0/according to** [**Semantic Versioning Specification**](https://semver.org/)**，并使用不可变的 Docker 镜像。**
6. 通过杀死 pod 来部署新的更新/修复，以便它们在重启过程中提取新的 Docker 镜像：**版本化您的代码，以便您可以更好地管理您的发布。**
7. 在同一个集群中混合生产和非生产工作负载：**如果可以，在不同的集群中运行您的生产和非生产工作负载。这降低了资源争用和意外环境交叉对您的生产环境造成的风险。**
8. 不使用蓝/绿或金丝雀进行关键任务部署（Kubernetes 的默认滚动更新并不总是足够）：**您应该考虑蓝/绿部署或金丝雀版本以减少生产压力和更有意义的生产结果。 **
9. 没有适当的指标来了解部署是否成功（您的健康检查需要应用程序支持）：**您应该确保监控您的部署以避免任何意外。您可以使用 Prometheus、Grafana、New Relic 或 Cisco AppDynamics 等工具来帮助您更好地了解您的部署。**
10. Cloud vendor lock-in: Locking yourself into an IaaS provider’s Kubernetes or serverless computing services: **Your business needs could change at any time. You shouldn’t unintentionally  lock yourself into a cloud provider since you can easily lift and shift  cloud-native applications.**
11. 云供应商锁定：将自己锁定在 IaaS 提供商的 Kubernetes 或无服务器计算服务中：**您的业务需求随时可能发生变化。您不应无意中将自己锁定在云提供商中，因为您可以轻松提升和转移云原生应用程序。**

Thanks for reading!

谢谢阅读！

[Better Programming](https://betterprogramming.pub/?source=post_sidebar--------------------------post_sidebar-----------)

Advice for programmers. 

给程序员的建议。

