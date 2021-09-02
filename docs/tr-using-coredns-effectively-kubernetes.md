# How to use CoreDNS Effectively with Kubernetes

# 如何在 Kubernetes 中有效地使用 CoreDNS

May 5th, 2021 From: https://www.infracloud.io/blogs/using-coredns-effectively-kubernetes/

### Background story

### 背景故事

We were increasing HTTP requests for one of our applications, hosted on the Kubernetes cluster, which resulted in a spike of 5xx errors.
The application is a GraphQL server calling a lot of external APIs and then returning an aggregated response.
Our initial response was to increase the number of replicas for the application to see if it improves the performance and reduces errors.
As we drilled down further with the application developer, found most of the failures were related to DNS resolution.
That’s where we started drilling down DNS resolution in Kubernetes.

我们增加了对托管在 Kubernetes 集群上的一个应用程序的 HTTP 请求，这导致了 5xx 错误的激增。
该应用程序是一个 GraphQL 服务器，调用大量外部 API，然后返回聚合响应。
我们最初的反应是增加应用程序的副本数量，看看它是否能提高性能并减少错误。
随着我们与应用程序开发人员进一步深入研究，发现大多数失败都与 DNS 解析有关。
这就是我们开始在 Kubernetes 中深入研究 DNS 解析的地方。

This post highlights our learnings related to using CoreDNS effectively with Kubernetes, as we did deep dive in process of configuring and troubleshooting.

这篇文章重点介绍了我们在 Kubernetes 中有效使用 CoreDNS 的相关知识，因为我们在配置和故障排除过程中进行了深入研究。

### CoreDNS Metrics

### CoreDNS 指标

DNS server stores record in its database and answers domain name query using the database.
If the DNS server doesn’t have this data, it tries to find for a solution from other DNS servers.

DNS 服务器在其数据库中存储记录并使用数据库回答域名查询。
如果 DNS 服务器没有此数据，它会尝试从其他 DNS 服务器中寻找解决方案。

CoreDNS became the [default DNS service](https://kubernetes.io/blog/2018/12/03/kubernetes-1-13-release-announcement/#coredns-is-now-the-default-dns-server-for-kubernetes) for Kubernetes 1.13+ onwards.
Nowadays, when you are using a managed Kubernetes cluster or you are self-managing a cluster for your application workloads, you often focus on tweaking your application but not much on the services provided by Kubernetes or how you are leveraging them.
DNS resolution is the basic requirement of any application so you need to ensure it’s working properly.
We would suggest looking at [dns-debugging-resolution](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/) troubleshooting guide and ensure your CoreDNS is configured and running properly.

CoreDNS 成为 [默认 DNS 服务](https://kubernetes.io/blog/2018/12/03/kubernetes-1-13-release-announcement/#coredns-is-now-the-default-dns-server-for-kubernetes) 用于 Kubernetes 1.13+。
如今，当您使用托管 Kubernetes 集群或为应用程序工作负载自行管理集群时，您通常会专注于调整应用程序，而不太关注 Kubernetes 提供的服务或如何利用它们。
DNS 解析是任何应用程序的基本要求，因此您需要确保它正常工作。
我们建议您查看 [dns-debugging-resolution](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/) 故障排除指南，并确保您的 CoreDNS 已配置并正常运行。

By default, when you provision a cluster you should always have a dashboard to observe for key CoreDNS metrics.
For getting CoreDNS metrics you should have [Prometheus plugin](https://coredns.io/plugins/metrics/) enabled as part of the CoreDNS config.

默认情况下，当您配置集群时，您应该始终有一个仪表板来观察关键的 CoreDNS 指标。
要获取 CoreDNS 指标，您应该启用 [Prometheus 插件](https://coredns.io/plugins/metrics/) 作为 CoreDNS 配置的一部分。

Below sample config using `prometheus` plugin to enable metrics collection from CoreDNS instance.

下面的示例配置使用 `prometheus` 插件从 CoreDNS 实例启用指标收集。

```
.:53 {
    errors
    health
    kubernetes cluster.local in-addr.arpa ip6.arpa {
      pods verified
      fallthrough in-addr.arpa ip6.arpa
    }
    prometheus :9153
    forward ./etc/resolv.conf
    cache 30
    loop
    reload
    loadbalance
}

```


Following are the key coreDNS metrics, we would suggest to have in your dashboard:
If you are using Prometheus, DataDog, Kibana etc, you may find ready to use dashboard template from community/provider.

以下是关键的 coreDNS 指标，我们建议您在您的仪表板中添加这些指标：
如果您正在使用 Prometheus、DataDog、Kibana 等，您可能会发现可以使用来自社区/提供商的仪表板模板。

- **Cache Hit percentage:** Percentage of requests responded using CoreDNS cache
- **DNS requests latency**  - CoreDNS: Time taken by CoreDNS to process DNS request
   - Upstream server: Time taken to process DNS request forwarded to upstream
- **Number of requests forwarded to upstream servers**
- **[Error codes](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6) for requests**  - NXDomain: Non-Existent Domain
   - FormErr: Format Error in DNS request
   - ServFail: Server Failure
   - NoError: No Error, successfully processed request
- **CoreDNS resource usage:** Different resources consumed by server such as memory, CPU etc.

- **缓存命中百分比：** 使用 CoreDNS 缓存响应的请求百分比
- **DNS 请求延迟** - CoreDNS：CoreDNS 处理 DNS 请求所用的时间
  - 上游服务器：处理转发到上游的 DNS 请求所需的时间
- **转发到上游服务器的请求数量**
- **[错误代码](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6) 用于请求** - NXDomain：不存在的域
  - FormErr：DNS 请求格式错误
  - ServFail：服务器故障
  - NoError：无错误，成功处理请求
- **CoreDNS 资源使用情况：** 服务器消耗的不同资源，如内存、CPU 等。

We were using DataDog for specific application monitoring. Following is just a sample dashboard I built with DataDog for my analysis.

我们使用 DataDog 进行特定的应用程序监控。以下只是我使用 DataDog 构建的用于分析的示例仪表板。

![datadog-coredns-dashboard](https://d33wubrfki0l68.cloudfront.net/26c0fc0512f17a51d420ed77746900ff5d0c278e/b52b1/assets/img/blog/using-coredns-effectively-with-kubernetes/dd-dashboard.png)

### How to reduce CoreDNS errors?

### 如何减少 CoreDNS 错误？

As we started drilling down more into how the application is making requests to CoreDNS, we observed most of the outbound requests happening through the application to an external API server.

当我们开始深入研究应用程序如何向 CoreDNS 发出请求时，我们观察到大多数出站请求通过应用程序发生到外部 API 服务器。

This is typically how resolv.conf looks in the application deployment pod.

这通常是 resolv.conf 在应用程序部署 pod 中的外观。

```
nameserver 10.100.0.10
search kube-namespace.svc.cluster.local svc.cluster.local cluster.local us-west-2.compute.internal
options ndots:5

```


If you understand how Kubernetes tries to resolve an FQDN, it tries to DNS lookup at different levels.

如果您了解 Kubernetes 如何尝试解析 FQDN，它会尝试在不同级别进行 DNS 查找。

Considering the above DNS config, when the DNS resolver sends a query to the CoreDNS server, it tries to search the domain considering the search path. 

考虑到上述 DNS 配置，当 DNS 解析器向 CoreDNS 服务器发送查询时，它会尝试考虑搜索路径来搜索域。

If we are looking for a domain boktube.io, it would make the following queries and it receives a successful response in the last query.

如果我们正在寻找域 boktube.io，它将进行以下查询，并在最后一个查询中收到成功响应。

```
botkube.io.kube-namespace.svc.cluster.local  <= NXDomain
botkube.io.svc.cluster.local <= NXDomain
boktube.io.cluster.local <= NXDomain
botkube.io.us-west-2.compute.internal <= NXDomain
botkube.io <= NoERROR

```


As we were making too many external lookups, we were getting a lot of NXDomain responses for DNS searches.
To optimize this we customized `spec.template.spec.dnsConfig` in the Deployment object.
This is how change looks like:

```
     dnsPolicy: ClusterFirst
     dnsConfig:
       options:
       - name: ndots
         value: "1"
```


With the above change, resolve.conf on pods changed.
The search was being performed only for an external domain.
This reduced number of queries to DNS servers.
This also helped in reducing 5xx errors for an application.
You can notice the difference in the NXDomain response count in the following graph.

通过上述更改，pod 上的 resolve.conf 发生了变化。
仅针对外部域执行搜索。
这减少了对 DNS 服务器的查询次数。
这也有助于减少应用程序的 5xx 错误。
您可以在下图中注意到 NXDomain 响应计数的差异。

![coredns-rcode-nxdomain-reduction](https://d33wubrfki0l68.cloudfront.net/59e1f460286518c33b5841818c0eb7dae16992cc/15811/assets/img/blog/using-coredns-effectively-with-kubernetes/coredns-rcode-nxdomain.png)

A better solution for this problem is [Node Level Cache](https://kubernetes.io/docs/tasks/administer-cluster/nodelocaldns/) which is introduced Kubernetes 1.18+.

这个问题更好的解决方案是 [Node Level Cache](https://kubernetes.io/docs/tasks/administer-cluster/nodelocaldns/)，它是 Kubernetes 1.18+ 引入的。

### How to customize CoreDNS to your needs?

### 如何根据您的需要自定义 CoreDNS？

We can customize CoreDNS by using plugins.
Kubernetes supports a different kind of workload and the standard CoreDNS config may not fit all your needs.
CoreDNS has a couple of in-tree and external plugins.
Depending on the kind of workloads you are running on your cluster, let’s say applications intercommunicating with each other or standalone apps that are interacting outside your Kubernetes cluster, the kind of FQDN you are trying to resolve might vary.
We should try to adjust the knobs of CoreDNS accordingly.
Suppose you are running Kubernetes in a particular public/private cloud and most of the DNS backed applications are in the same cloud.
In that case, CoreDNS also provides particular cloud-related or generic plugins which can be used to extend DNS zone records.

我们可以使用插件自定义 CoreDNS。
Kubernetes 支持不同类型的工作负载，标准 CoreDNS 配置可能无法满足您的所有需求。
CoreDNS 有几个树内和外部插件。
根据您在集群上运行的工作负载类型，假设应用程序相互通信或在 Kubernetes 集群外交互的独立应用程序，您尝试解决的 FQDN 类型可能会有所不同。
我们应该尝试相应地调整 CoreDNS 的旋钮。
假设您在特定的公共/私有云中运行 Kubernetes，并且大多数 DNS 支持的应用程序都在同一云中。
在这种情况下，CoreDNS 还提供特定的云相关或通用插件，可用于扩展 DNS 区域记录。

If you are interested in customizing DNS behaviour for your needs, we would suggest going through the book “Learning CoreDNS” by Cricket Liu and John Belamaric.
Book provides a detailed overview of different CoreDNS plugins with their use cases.
It also covers CoreDNS + Kubernetes integration in depth.

如果您有兴趣根据自己的需要自定义 DNS 行为，我们建议您阅读 Cricket Liu 和 John Belamaric 合着的“Learning CoreDNS”一书。
本书详细概述了不同的 CoreDNS 插件及其用例。
它还深入介绍了 CoreDNS + Kubernetes 集成。

If you are running an appropriate number of instances of CoreDNS in your Kubernetes cluster or not is one of the critical factors to decide.
It’s recommended to run a least two instances of the CoreDNS servers for a better guarantee of DNS requests being served.
Depending upon the number of requests being served, nature of requests, number of workloads running on the cluster, and size of cluster you may need to add extra instances of CoreDNS or configure HPA (Horizontal Pod Autoscaler) for your cluster.
Factors like the number of requests being served, nature of requests, number of workloads running on the cluster and cluster size should help you in deciding the number of CoreDNS instances.
You may need to add extra instances of CoreDNS or configure HPA (Horizontal Pod Autoscaler) for your cluster.

是否在 Kubernetes 集群中运行适当数量的 CoreDNS 实例是决定的关键因素之一。
建议至少运行两个 CoreDNS 服务器实例，以更好地保证 DNS 请求得到服务。
根据服务的请求数量、请求的性质、集群上运行的工作负载数量以及集群的大小，您可能需要为集群添加额外的 CoreDNS 实例或配置 HPA（水平 Pod 自动扩展）。
服务的请求数量、请求的性质、集群上运行的工作负载数量和集群大小等因素应该可以帮助您决定 CoreDNS 实例的数量。
您可能需要为集群添加额外的 CoreDNS 实例或配置 HPA（Horizontal Pod Autoscaler）。

### Summary

###  概括

This blog post tries to highlight the importance of the DNS request cycle in Kubernetes, and many times you would end up in a situation where you start with “it’s not DNS” but end up “it’s always DNS !”.
So be careful of these landmines.

这篇博文试图强调 DNS 请求周期在 Kubernetes 中的重要性，很多时候你会遇到这样的情况：从“它不是 DNS”开始，但最终“它总是 DNS ！”。
所以要小心这些地雷。

Enjoyed the article? Let’s start a conversation on [Twitter](https://twitter.com/sanketsudake) and share your “It’s always DNS” stories and how it was resolved.

喜欢这篇文章吗？让我们在 [Twitter](https://twitter.com/sanketsudake) 上开始对话并分享您的“始终是 DNS”的故事以及它是如何解决的。

Love Cloud Native? We do too ❤️ 

喜欢云原生？我们也一样❤️


