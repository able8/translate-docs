# Six Strategies for Application Deployment

# 应用部署六大策略

21 Nov 2017

There are a variety of techniques to deploy new applications to  production, so choosing the right strategy is an important decision,  weighing the options in terms of the impact of change on the system, and on the end-users.

有多种技术可以将新应用程序部署到生产环境中，因此选择正确的策略是一项重要决策，需要根据更改对系统和最终用户的影响来权衡各种选项。

In this post, we are going to talk about the following strategies:

- **Recreate**: Version A is terminated then version B is rolled out.
- **Ramped** (also known as rolling-update or incremental): Version B is slowly rolled out and replacing version A.
- **Blue/Green**: Version B is released alongside version A, then the traffic is switched to version B.
- **Canary**: Version B is released to a subset of users, then proceed to a full rollout.
- **A/B testing**: Version B is released to a subset of users under specific condition.
- **Shadow**: Version B receives real-world traffic alongside version A and doesn’t impact the response.

在这篇文章中，我们将讨论以下策略：

- **重新创建**：终止版本 A，然后推出版本 B。
- **Ramped**（也称为滚动更新或增量）：版本 B 缓慢推出并替换版本 A。
- **蓝/绿**：版本 B 与版本 A 一起发布，然后流量切换到版本 B。
- **Canary**：版本 B 发布给部分用户，然后进行全面部署。
- **A/B 测试**：版本 B 在特定条件下发布给部分用户。
- **Shadow**：版本 B 与版本 A 一起接收真实世界的流量，并且不会影响响应。

Let’s take a look at each strategy and see which strategy would fit  best for a particular use case. For the sake of simplicity, we used [Kubernetes](https://kubernetes.io) and tested the example against [Minikube](https://github.com/kubernetes/minikube). Examples of configuration and step-by-step approaches on for each strategy [can be found in this git repository](https://github.com/ContainerSolutions/k8s-deployment-strategies).

让我们来看看每种策略，看看哪种策略最适合特定用例。为简单起见，我们使用了 [Kubernetes](https://kubernetes.io) 并针对 [Minikube](https://github.com/kubernetes/minikube) 测试了示例。每个策略的配置示例和分步方法 [可以在这个 git 存储库中找到](https://github.com/ContainerSolutions/k8s-deployment-strategies)。

## Recreate

## 重新创建

The recreate strategy is a dummy deployment which consists of  shutting down version A then deploying version B after version A is  turned off. This technique implies downtime of the service that depends  on both shutdown and boot duration of the application.

重新创建策略是一种虚拟部署，包括关闭版本 A，然后在关闭版本 A 后部署版本 B。这种技术意味着服务的停机时间取决于应用程序的关闭和启动持续时间。

![img](https://storage.googleapis.com/cdn.thenewstack.io/media/2017/11/c42fa239-recreate.gif)

Pros:

- Easy to setup.
- Application state entirely renewed.

优点：

- 易于设置。
- 应用程序状态完全更新。

Cons:

- High impact on the user, expect downtime that depends on both shutdown and boot duration of the application.

缺点：

- 对用户的影响很大，预计停机时间取决于应用程序的关闭和启动持续时间。

## Ramped

## 斜坡

The ramped deployment strategy consists of slowly rolling out a version of an application by replacing instances one after the other  until all the instances are rolled out. It usually follows the following process: with a pool of version A behind a load balancer, one instance  of version B is deployed. When the service is ready to accept traffic,  the instance is added to the pool. Then, one instance of version A is  removed from the pool and shut down.

渐进式部署策略包括通过一个接一个地替换实例来缓慢地推出应用程序的一个版本，直到推出所有实例。它通常遵循以下过程：在负载均衡器后面使用版本 A 的池，部署版本 B 的一个实例。当服务准备好接受流量时，实例被添加到池中。然后，从池中删除版本 A 的一个实例并关闭。

Depending on the system taking care of the ramped deployment, you can tweak the following parameters to increase the deployment time:

- Parallelism, max batch size: Number of concurrent instances to roll out.
- Max surge: How many instances to add in addition of the current amount.
- Max unavailable: Number of unavailable instances during the rolling update procedure.

根据负责加速部署的系统，您可以调整以下参数以增加部署时间：

- 并行性，最大批量大小：要推出的并发实例数。
- 最大激增：在当前数量的基础上要添加多少个实例。
- 最大不可用：滚动更新过程中不可用实例的数量。

![img](https://storage.googleapis.com/cdn.thenewstack.io/media/2017/11/5bddc931-ramped.gif)

Pros:

- Easy to set up.
- Version is slowly released across instances.
- Convenient for stateful applications that can handle rebalancing of the data.

优点：

- 易于设置。
- 版本跨实例缓慢发布。
- 方便可以处理数据重新平衡的有状态应用程序。

Cons:

- Rollout/rollback can take time.
- Supporting multiple APIs is hard.
- No control over traffic.

缺点：

- 推出/回滚可能需要时间。
- 支持多个 API 很困难。
- 无法控制交通。

## Blue/Green

##  蓝绿

The blue/green deployment strategy differs from a ramped deployment,  version B (green) is deployed alongside version A (blue) with exactly  the same amount of instances. After testing that the new version meets  all the requirements the traffic is switched from version A to version B at the load balancer level.

蓝/绿部署策略与斜坡部署不同，版本 B（绿色）与版本 A（蓝色）一起部署，实例数量完全相同。在测试新版本满足所有要求后，流量在负载均衡器级别从版本 A 切换到版本 B。

![img](https://storage.googleapis.com/cdn.thenewstack.io/media/2017/11/73a2824d-blue-green.gif)

Pros:

- Instant rollout/rollback.
- Avoid versioning issue, the entire application state is changed in one go.

优点：

- 即时推出/回滚。
- 避免版本控制问题，整个应用程序状态一次性更改。

Cons:

- Expensive as it requires double the resources.
- Proper test of the entire platform should be done before releasing to production.
- Handling stateful applications can be hard.

缺点：

- 昂贵，因为它需要双倍的资源。
- 在发布到生产环境之前，应对整个平台进行适当的测试。
- 处理有状态的应用程序可能很困难。

## Canary

## 金丝雀

A canary deployment consists of gradually shifting production traffic from version A to version B. Usually the traffic is split based on  weight. For example, 90 percent of the requests go to version A, 10  percent go to version B.

金丝雀部署包括逐渐将生产流量从版本 A 转移到版本 B。通常根据权重拆分流量。例如，90% 的请求转到版本 A，10% 转到版本 B。

This technique is mostly used when the tests are lacking or not  reliable or if there is little confidence about the stability of the new release on the platform.

当测试缺乏或不可靠或者对平台上新版本的稳定性没有信心时，通常会使用此技术。

![img](https://storage.googleapis.com/cdn.thenewstack.io/media/2017/11/a6324354-canary.gif)

Pros:

- Version released for a subset of users.
- Convenient for error rate and performance monitoring.
- Fast rollback.

优点：

- 为部分用户发布的版本。
- 便于错误率和性能监控。
- 快速回滚。

Con:

- Slow rollout.

缺点：

- 缓慢推出。

## A/B testing 

## A/B 测试

A/B testing deployments consists of routing a subset of users to a  new functionality under specific conditions. It is usually a technique  for making business decisions based on statistics, rather than a  deployment strategy. However, it is related and can be implemented by adding extra functionality to a canary deployment so we will briefly discuss it here.

A/B 测试部署包括在特定条件下将一部分用户路由到新功能。它通常是一种基于统计信息而非部署策略做出业务决策的技术。但是，它是相关的，可以通过向金丝雀部署添加额外功能来实现，因此我们将在此处简要讨论它。

This technique is widely used to test conversion of a given feature and only roll-out the version that converts the most.

此技术广泛用于测试给定功能的转换，并仅推出转换最多的版本。

Here is a list of conditions that can be used to distribute traffic amongst the versions:

- By browser cookie
- Query parameters
- Geolocalisation
- Technology support: browser version, screen size, operating system, etc.
- Language

以下是可用于在版本之间分配流量的条件列表：

- 通过浏览器cookie
- 查询参数
- 地理定位
- 技术支持：浏览器版本、屏幕尺寸、操作系统等。
- 语

![img](https://storage.googleapis.com/cdn.thenewstack.io/media/2017/11/5deeea9c-a-b.gif)

Pros:

- Several versions run in parallel.
- Full control over the traffic distribution.

优点：

- 多个版本并行运行。
- 完全控制流量分布。

Cons:

- Requires intelligent load balancer.
- Hard to troubleshoot errors for a given session, distributed tracing becomes mandatory.

缺点：

- 需要智能负载平衡器。
- 很难对给定会话的错误进行故障排除，分布式跟踪成为强制性的。

## Shadow

##  阴影

A shadow deployment consists of releasing version B alongside version A, fork version A’s incoming requests and send them to version B as  well without impacting production traffic. This is particularly useful  to test production load on a new feature. A rollout of the application  is triggered when stability and performance meet the requirements.

影子部署包括与版本 A 一起发布版本 B，分叉版本 A 的传入请求并将它们发送到版本 B，而不会影响生产流量。这对于测试新功能的生产负载特别有用。当稳定性和性能满足要求时，将触发应用程序的推出。

This technique is fairly complex to setup and needs special  requirements, especially with egress traffic. For example, given a  shopping cart platform, if you want to shadow test the payment service  you can end-up having customers paying twice for their order. In this  case, you can solve it by creating a mocking service that replicates the response from the provider.

这种技术设置起来相当复杂，需要特殊要求，尤其是出口流量。例如，给定一个购物车平台，如果您想对支付服务进行影子测试，您最终可以让客户为他们的订单支付两次费用。在这种情况下，您可以通过创建一个复制来自提供者的响应的模拟服务来解决它。

![img](https://storage.googleapis.com/cdn.thenewstack.io/media/2017/11/fdd947f8-shadow.gif)

Pros:

- Performance testing of the application with production traffic.
- No impact on the user.
- No rollout until the stability and performance of the application meet the requirements.

优点：

- 使用生产流量对应用程序进行性能测试。
- 对用户没有影响。
- 在应用程序的稳定性和性能满足要求之前不会推出。

Cons:

- Expensive as it requires double the resources.
- Not a true user test and can be misleading.
- Complex to setup.
- Requires mocking service for certain cases.

缺点：

- 昂贵，因为它需要双倍的资源。
- 不是真正的用户测试，可能会产生误导。
- 设置复杂。
- 在某些情况下需要模拟服务。

## To Sum Up

##  总结

There are multiple ways to deploy a new version of an application and it really depends on the needs and budget. When releasing to  development/staging environments, a recreate or ramped deployment is  usually a good choice. When it comes to production, a ramped or  blue/green deployment is usually a good fit, but proper testing of the  new platform is necessary.

有多种方法可以部署应用程序的新版本，这实际上取决于需求和预算。当发布到开发/暂存环境时，重新创建或升级部署通常是一个不错的选择。在生产方面，斜坡或蓝/绿部署通常是一个很好的选择，但对新平台进行适当的测试是必要的。

Blue/green and shadow strategies have more impact on the budget as it requires double resource capacity. If the application lacks in tests or if there is little confidence about the impact/stability of the  software, then a canary, a/ b testing or shadow release can be used. If  your business requires testing of a new feature amongst a specific pool  of users that can be filtered depending on some parameters like  geolocation, language, operating system or browser features, then you  may want to use the a/b testing technique.

蓝/绿和影子战略对预算的影响更大，因为它需要双倍的资源能力。如果应用程序缺乏测试，或者对软件的影响/稳定性没有信心，那么可以使用金丝雀、a/b 测试或影子发布。如果您的企业需要在特定用户群中测试新功能，并且可以根据地理位置、语言、操作系统或浏览器功能等一些参数进行过滤，那么您可能需要使用 a/b 测试技术。

Last but not least, a shadow release is complex and requires extra  work to mock egress traffic which is mandatory when calling external  dependencies with mutable actions (email, bank, etc.). However, this  technique can be useful when migrating to a new database technology and  use shadow traffic to monitor system performance under load.

最后但并非最不重要的一点是，影子发布很复杂，需要额外的工作来模拟出口流量，这在使用可变操作（电子邮件、银行等）调用外部依赖项时是必需的。但是，当迁移到新的数据库技术并使用影子流量来监视负载下的系统性能时，此技术非常有用。

Below is a diagram to help you choose the right strategy:

[![img](https://storage.googleapis.com/cdn.thenewstack.io/media/2017/11/9e09392d-k8s_deployment_strategies.png)](https://storage.googleapis.com/cdn.thenewstack.io/media/2017/11/9e09392d-k8s_deployment_strategies.png)

下图可帮助您选择正确的策略：


Depending on the Cloud provider or platform, the following docs can be a good start to understand deployment:

- [Amazon Web Services](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html)
- [Docker Swarm](https://docs.docker.com/engine/swarm/swarm-tutorial/rolling-update/)
- [Google Cloud](https://cloud.google.com/sdk/gcloud/reference/beta/compute/instance-groups/managed/rolling-action/replace)
- [Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/update-intro/)

根据云提供商或平台的不同，以下文档可能是了解部署的好的开始：

- [亚马逊网络服务](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html)
- [Docker Swarm](https://docs.docker.com/engine/swarm/swarm-tutorial/rolling-update/)
- [谷歌云](https://cloud.google.com/sdk/gcloud/reference/beta/compute/instance-groups/managed/rolling-action/replace)
- [Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/update-intro/)

I hope this was useful, if you have any questions/feedback feel free to comment below. 

我希望这是有用的，如果您有任何问题/反馈，请随时在下面发表评论。
