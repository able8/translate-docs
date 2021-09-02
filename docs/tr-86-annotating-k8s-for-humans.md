# Annotating Kubernetes Services for Humans

# 为 Kubernetes 注释服务

Tuesday, April 20, 2021

Have you ever been asked to troubleshoot a failing Kubernetes service and struggled to find basic information about the service such as the  source repository and owner?

您是否曾被要求对失败的 Kubernetes 服务进行故障排除，并且难以找到有关该服务的基本信息，例如源存储库和所有者？

One of the problems as Kubernetes applications grow is the  proliferation of services. As the number of services grows, developers  start to specialize working with specific services. When it comes to  troubleshooting, however, developers need to be able to find the source, understand the service and dependencies, and chat with the owning team  for any service.

随着 Kubernetes 应用程序的增长，问题之一是服务的激增。随着服务数量的增长，开发人员开始专门处理特定服务。但是，在进行故障排除时，开发人员需要能够找到源，了解服务和依赖项，并与拥有任何服务的团队聊天。

## Human service discovery

## 人工服务发现

Troubleshooting always begins with information gathering. While much attention has been paid to centralizing machine data (e.g., logs, metrics), much less  attention has been given to the human aspect of service discovery. Who  owns a particular service? What Slack channel does the team work on? Where is the source for the service? What issues are currently known and being tracked?

故障排除总是从信息收集开始。虽然对集中机器数据（例如日志、指标）给予了很多关注，但对服务发现的人为方面的关注却少得多。谁拥有特定服务？团队在哪个 Slack 频道上工作？服务的来源在哪里？目前已知并正在跟踪哪些问题？

## Kubernetes annotations

## Kubernetes 注释

Kubernetes annotations are designed to solve exactly this problem. Oft-overlooked, Kubernetes annotations are designed to add metadata to Kubernetes  objects. The Kubernetes documentation says annotations can “attach  arbitrary non-identifying metadata to objects.” This means that  annotations should be used for attaching metadata that is external to  Kubernetes (ie, metadata that Kubernetes won't use to identify  objects. As such, annotations can contain any type of data. This is a  contrast to labels, which are designed for uses internal to Kubernetes.  As such, label structure and values are [constrained](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set) so they can be efficiently used by Kubernetes.

Kubernetes 注释旨在解决这个问题。经常被忽视的 Kubernetes 注释旨在向 Kubernetes 对象添加元数据。 Kubernetes 文档说注释可以“将任意的非识别元数据附加到对象上”。这意味着注释应该用于附加 Kubernetes 外部的元数据（即 Kubernetes 不会用来标识对象的元数据。因此，注释可以包含任何类型的数据。这与标签形成对比，标签被设计为用于 Kubernetes 内部的用途。因此，标签结构和值是 [约束](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set) 所以Kubernetes 可以有效地使用它们。

## Kubernetes annotations in action

## Kubernetes 注释实践

Here is an example. Imagine you have a Kubernetes service for quoting, called the quote service. You can do the following:

```
kubectl annotate service quote a8r.io/owner=”@sally”
```


In this example, we've just added an annotation called `a8r.io/owner` with the value of @sally. Now, we can use `kubectl describe` to get the information.

在这个例子中，我们刚刚添加了一个名为 `a8r.io/owner` 的注解，其值为 @sally。现在，我们可以使用 `kubectl describe` 来获取信息。

```
Name:              quote
Namespace:         default
Labels:            <none>
Annotations:       a8r.io/owner: @sally
Selector:          app=quote
Type:              ClusterIP
IP:                10.109.142.131
Port:              http  80/TCP
TargetPort:        8080/TCP
Endpoints:         <none>
Session Affinity:  None
Events:            <none>
```


If you’re practicing GitOps (and you should be!) you’ll want to code these values directly into your Kubernetes manifest, e.g.,

如果您正在练习 GitOps（您应该这样做！），您将希望将这些值直接编码到您的 Kubernetes 清单中，例如，

```yaml
apiVersion: v1
kind: Service
metadata:
  name: quote
  annotations:
    a8r.io/owner: “@sally”
spec:
  ports:
  - name: http
    port: 80
    targetPort: 8080
  selector:
    app: quote
```


## A Convention for Annotations

## 注释约定

Adopting a common convention for annotations ensures consistency and  understandability. Typically, you’ll want to attach the annotation to  the service object, as services are the high-level resource that maps  most clearly to a team’s responsibility. Namespacing your annotations is also very important. Here is one set of conventions, documented at [a8r.io](https://a8r.io), and reproduced below:

对注释采用通用约定可确保一致性和可理解性。通常，您需要将注释附加到服务对象，因为服务是最清楚地映射到团队职责的高级资源。给注释命名空间也很重要。这是一组约定，记录在 [a8r.io](https://a8r.io)，并转载如下：

| Annotation             | Description                                                  |
| ---------------------- |------------------------------------------------------------ |
| `a8r.io/description`   | Unstructured text description of the service for humans. |
| `a8r.io/owner`         | SSO username (GitHub), email address (linked to GitHub account), or unstructured owner description. |
| `a8r.io/chat`          | Slack channel, or link to external chat system. |
| `a8r.io/bugs`          | Link to external bug tracker. |
| `a8r.io/logs`          | Link to external log viewer. |
| `a8r.io/documentation` | Link to external project documentation. |
| `a8r.io/repository`    | Link to external VCS repository. |
| `a8r.io/support`       | Link to external support center. |
| `a8r.io/runbook`       | Link to external project runbook. |
| `a8r.io/incidents`     | Link to external incident dashboard. | 

## Visualizing annotations: Service Catalogs

## 可视化注释：服务目录

As the number of microservices and annotations proliferate, running `kubectl describe` can get tedious. Moreover, using `kubectl describe` requires every developer to have some direct access to the Kubernetes  cluster. Over the past few years, service catalogs have gained greater  visibility in the Kubernetes ecosystem. Popularized by tools such as [Shopify's ServicesDB](https://shopify.engineering/scaling-mobile-development-by-treating-apps-as-services) and [Spotify's System Z](https://dzone.com/articles/modeling-microservices-at-spotify-with-petter-mari), service catalogs are internally-facing developer portals that present critical information about microservices.

随着微服务和注释的数量激增，运行 `kubectl describe` 可能会变得乏味。此外，使用 `kubectl describe` 要求每个开发人员都可以直接访问 Kubernetes 集群。在过去几年中，服务目录在 Kubernetes 生态系统中获得了更高的知名度。通过 [Shopify 的 ServicesDB](https://shopify.engineering/scaling-mobile-development-by-treating-apps-as-services) 和 [Spotify 的 System Z](https://dzone.com/articles)等工具普及/modeling-microservices-at-spotify-with-petter-mari)，服务目录是面向内部的开发人员门户，提供有关微服务的关键信息。

Note that these service catalogs should not be confused with the [Kubernetes Service Catalog project](https://svc-cat.io/). Built on the Open Service Broker API, the Kubernetes Service Catalog  enables Kubernetes operators to plug in different services (e.g.,  databases) to their cluster.

请注意，这些服务目录不应与 [Kubernetes 服务目录项目](https://svc-cat.io/) 混淆。 Kubernetes Service Catalog 建立在 Open Service Broker API 之上，使 Kubernetes 运营商能够将不同的服务（例如，数据库)插入到他们的集群中。

## Annotate your services now and thank yourself later

## 现在注释您的服务，稍后感谢您自己

Much like implementing observability within microservice systems, you often  don’t realize that you need human service discovery until it’s too late. Don't wait until something is on fire in production to start wishing  you had implemented better metrics and also documented how to get in  touch with the part of your organization that looks after it.

就像在微服务系统中实现可观察性一样，您通常不会意识到需要人工服务发现，直到为时已晚。不要等到生产中发生某些事情时才开始希望您已经实施了更好的指标，并记录了如何与您的组织中负责管理它的部分取得联系。

There's enormous benefits to building an effective “version 0” service: a [*dancing skeleton*](https://containerjournal.com/topics/container-management/dancing-skeleton-apis-and-microservices/) application with a thin slice of complete functionality that can be  deployed to production with a minimal yet effective continuous delivery  pipeline.

构建有效的“版本 0”服务有巨大的好处：[*dancing skeleton*](https://containerjournal.com/topics/container-management/dancing-skeleton-apis-and-microservices/) 应用程序可以通过最小但有效的持续交付管道部署到生产中的完整功能片段。

Adding service annotations should be an essential part of your  “version 0” for all of your services. Add them now, and you’ll thank  yourself later. 

添加服务注释应该是所有服务的“版本 0”的重要组成部分。现在添加它们，稍后您会感谢自己。


