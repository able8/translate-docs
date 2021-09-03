# What’s new in Kubernetes 1.20?

# Kubernetes 1.20 有什么新变化？

on December 1, 2020 From: https://sysdig.com/blog/whats-new-kubernetes-1-20/

Table of contents 目录

**Kubernetes 1.20** is about to be released, and it comes packed with novelties! Where do we begin?

**Kubernetes 1.20** 即将发布，新鲜出炉！我们从哪里开始？

As we highlighted in the last release, enhancements [now have to move forward to stability or being deprecated](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1635). As a result, alpha features that have been around since the early times of Kubernetes, like [CronJobs](http://sysdig.com#19) and [Kubelet CRI support](http://sysdig.com#2040), are now getting the attention they deserve.

正如我们在上一个版本中强调的那样，增强功能 [现在必须转向稳定性或被弃用](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1635)。因此，自 Kubernetes 早期就存在的 alpha 功能，如 [CronJobs](http://sysdig.com#19) 和 [Kubelet CRI support](http://sysdig.com#2040)，现在得到了他们应得的关注。

Another noteworthy fact of this Kubernetes 1.20 release is that it brings 43 enhancements, up from 34 in 1.19. Of those 43 enhancements, 11 are graduating to Stable, 15 are completely new, and 17 are existing features that keep improving.

Kubernetes 1.20 版本的另一个值得注意的事实是它带来了 43 项增强功能，高于 1.19 中的 34 项。在这 43 个增强功能中，11 个升级到稳定版，15 个是全新的，17 个是不断改进的现有功能。

So many enhancements means that they are smaller in scope. Kubernetes 1.20 is a healthy house cleaning event with a lot of small user-friendly changes. For example, improvements in kube-apiserver to [work better in HA clusters](http://sysdig.com#1965) and [reboot more efficiently](http://sysdig.com#1904) after an upgrade. Or, being able to [gracefully shutdown nodes](http://sysdig.com#2000) so resources can be freed properly. It’s exciting to see small features like these paving the way for the big changes that are to come.

如此多的增强意味着它们的范围更小。 Kubernetes 1.20 是一个健康的房屋清洁活动，有很多用户友好的小变化。例如，kube-apiserver 的改进[在 HA 集群中更好地工作](http://sysdig.com#1965)和升级后[更有效地重新启动](http://sysdig.com#1904)。或者，能够[正常关闭节点](http://sysdig.com#2000) 以便正确释放资源。看到这样的小功能为即将到来的大变化铺平了道路，真是令人兴奋。

Here is the detailed list of what’s new in Kubernetes 1.20.

以下是 Kubernetes 1.20 新功能的详细列表。

## Kubernetes 1.20 – Editor’s pick:

These are the features that look most exciting to us in this release (ymmv):

### [\#1753 Kubernetes system components logs sanitization](http://sysdig.com\#1753)

## Kubernetes 1.20 – 编辑推荐：

这些是我们在此版本 (ymmv) 中最令人兴奋的功能：

### [\#1753 Kubernetes 系统组件日志清理](http://sysdig.com\#1753)

There have been some vulnerabilities lately related to credentials being [leaked into the log output](https://sysdig.com/blog/falco-cve-2020-8566-ceph/). It’s comforting knowing that they are being approached by keeping the big picture in mind, identifying the potential sources of leaks, and placing a redacting mechanism in place to cut those leaks.

最近出现了一些与凭据[泄露到日志输出中](https://sysdig.com/blog/falco-cve-2020-8566-ceph/) 相关的漏洞。知道他们正在通过牢记大局，确定泄漏的潜在来源并设置编辑机制来减少这些泄漏，这是令人欣慰的。

_[Vicente Herrera](https://twitter.com/Vicen_Herrera) – Cloud native security advocate at Sysdig_

_[Vicente Herrera](https://twitter.com/Vicen_Herrera) – Sysdig 的云原生安全倡导者

### [\#2047 CSIServiceAccountToken](http://sysdig.com\#2047)

### [\#2047 CSIServiceAccountToken](http://sysdig.com\#2047)

This enhancement represents the huge effort to improve the security around authentication and the handling of tokens that is taking place in Kubernetes. Take a look at [the Auth section](http://sysdig.com#auth) in this article to see what I mean. This particular feature makes access to volumes that requires authentication (like secret vaults) more secure and easier to set up.

这种增强代表了为提高围绕身份验证和 Kubernetes 中发生的令牌处理的安全性所做的巨大努力。看一下本文中的[身份验证部分](http://sysdig.com#auth)以了解我的意思。此特殊功能使访问需要身份验证的卷（如秘密保险库)更安全且更易于设置。

_[Víctor Jiménez](https://twitter.com/capitangolo/) – Content Marketing Engineer at Sysdig_

_[Víctor Jiménez](https://twitter.com/capitangolo/) – Sysdig 的内容营销工程师_

### [\#19 CronJobs](http://sysdig.com\#19) \+ [\#2040 Kubelet CRI support](http://sysdig.com\#2040)

You just need to look at the issue number (19) from CronJobs to know these are not new features. CronJobs have been around since 1.4 and CRI support since 1.5, and although used widely in production, neither are considered stable yet. It’s comforting seeing that features you depend on to run your production cluster aren’t alphas anymore.

您只需查看 CronJobs 的问题编号 (19) 即可知道这些不是新功能。 CronJobs 从 1.4 开始就已经存在，从 1.5 开始就支持 CRI，虽然在生产中广泛使用，但它们都被认为是稳定的。令人欣慰的是，您运行生产集群所依赖的功能不再是 alpha 版本。

_David de Torres – Integrations Engineer at Sysdig_

_David de Torres – Sysdig 的集成工程师_

### [\#2000 Graceful node shutdown](http://sysdig.com\#2000)

### [\#2000 节点正常关闭](http://sysdig.com\#2000)

A small feature? Yes. A huge demonstration of goodwill towards developers? That is also true. Being able to properly release resources when a node shuts down will avoid many weird behaviors, including the hairy ones that are tough to debug and troubleshoot.

一个小功能？是的。对开发的善意的巨大表现？这也是事实。能够在节点关闭时正确释放资源将避免许多奇怪的行为，包括难以调试和排除故障的毛茸茸的行为。

_Álvaro Iradier – Integrations Engineer at Sysdig__

### [\#1965 kube-apiserver identity](http://sysdig.com\#1965)

### [\#1965 kube-apiserver 身份](http://sysdig.com\#1965)

Having a unique identifier for each kube-apiserver instance is one of those features that usually goes unnoticed. However, knowing that it will enable better high availability features in future Kubernetes versions really hypes me up.

每个 kube-apiserver 实例都有一个唯一标识符是通常不被注意的特性之一。然而，知道它将在未来的 Kubernetes 版本中实现更好的高可用性功能真的让我大吃一惊。

_[Mateo Burillo](https://twitter.com/mateobur) – Product Manager at Sysdig_

_[Mateo Burillo](https://twitter.com/mateobur) – Sysdig 的产品经理_

### [\#1748 Expose metrics about resource requests and limits that represent the pod model](http://sysdig.com\#1748)

### [\#1748 公开有关代表 pod 模型的资源请求和限制的指标](http://sysdig.com\#1748)

Yes! More metrics are always welcome. Even more so when they will enable you to better plan the capacity of your cluster and help troubleshoot eviction problems.

是的！更多的指标总是受欢迎的。当它们使您能够更好地规划集群容量并帮助解决驱逐问题时，更是如此。

## Deprecations in Kubernetes 1.20

## Kubernetes 1.20 中的弃用

### [\#1558](https://github.com/kubernetes/enhancements/issues/1558) Streaming proxy redirects

### [\#1558](https://github.com/kubernetes/enhancements/issues/1558) 流代理重定向

**Stage:** Deprecation

**阶段：** 弃用

As part of deprecating `StreamingProxyRedirects`, the `--redirect-container-streaming` flag can no longer be enabled.

作为弃用 `StreamingProxyRedirects` 的一部分，不能再启用 `--redirect-container-streaming` 标志。

Both the `StreamingProxyRedirects` feature gate and the `--redirect-container-streaming` flag were marked as deprecated on 1.18. The feature gate will be disabled by default in 1.22, and both will be removed on 1.24.

`StreamingProxyRedirects` 特性门和 `--redirect-container-streaming` 标志在 1.18 中被标记为已弃用。默认情况下，功能门将在 1.22 中禁用，并且两者都将在 1.24 中删除。

You can read the rationale behind this deprecation [in the KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20191205-container-streaming-requests.md).

您可以阅读 [在 KEP 中](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20191205-container-streaming-requests.md)弃用背后的基本原理。

### [\#2067](https://github.com/kubernetes/enhancements/issues/2067) Rename the kubeadm master node-role label and taint

### [\#2067](https://github.com/kubernetes/enhancements/issues/2067) 重命名 kubeadm 主节点-角色标签和污点

**Stage:** Deprecation

**阶段：** 弃用

**Feature group:** cluster-lifecycle

**功能组：**集群生命周期

In an effort to move away from offensive wording, `node-role.kubernetes.io/master` is being renamed to `node-role.kubernetes.io/control-plane`.

为了摆脱令人反感的措辞，`node-role.kubernetes.io/master` 被重命名为 `node-role.kubernetes.io/control-plane`。

You can read more about the implications of this change [in the KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/kubeadm/2067-rename-master-label-taint).

您可以在 [KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/kubeadm/2067-rename-master-label-污点)。

### [\#1164](https://github.com/kubernetes/enhancements/issues/1164) Deprecate and remove SelfLink

### [\#1164](https://github.com/kubernetes/enhancements/issues/1164) 弃用并删除 SelfLink

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** api-machinery

**功能组：** api-machinery

The field `SelfLink` is present in every Kubernetes object and contains a URL representing the given object. This field does not provide any new information and its creation and maintenance has a performance impact.

“SelfLink”字段存在于每个 Kubernetes 对象中，并包含表示给定对象的 URL。此字段不提供任何新信息，其创建和维护会影响性能。

Initially deprecated in [Kubernetes 1.16](https://sysdig.com/blog/whats-new-kubernetes-1-16/#1164), the feature gate is now disabled by default and will finally be removed in Kubernetes 1.21.

最初在 [Kubernetes 1.16](https://sysdig.com/blog/whats-new-kubernetes-1-16/#1164) 中被弃用，特性门现在默认禁用，最终将在 Kubernetes 1.21 中删除。

## Kubernetes 1.20 API

## Kubernetes 1.20 API

### [\#1965](https://github.com/kubernetes/enhancements/issues/1965) kube-apiserver identity

### [\#1965](https://github.com/kubernetes/enhancements/issues/1965) kube-apiserver 身份

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** api-machinery

**功能组：** api-machinery

In order to better control which _kube-apiservers_ are alive in a high availability cluster, a new lease / heartbeat system has been implemented.

为了更好地控制高可用性集群中哪些 _kube-apiservers_ 处于活动状态，已经实施了新的租用/心跳系统。

Each `kube-apiserver` will assign a unique ID to itself in the format of `kube-apiserver-{UUID}`. These IDs will be stored in Lease objects that will be refreshed every 10 seconds (by default), and garbage collected if not renewed.

每个 `kube-apiserver` 都会以 `kube-apiserver-{UUID}` 的格式为自己分配一个唯一的 ID。这些 ID 将存储在 Lease 对象中，该对象将每 10 秒刷新一次（默认情况下），如果不更新，则会被垃圾收集。

This system is similar to the node heartbeat. In fact, it will be reusing the existing Kubelet heartbeat logic.

这个系统类似于节点心跳。事实上，它将重用现有的 Kubelet 心跳逻辑。

### [\#1904](https://github.com/kubernetes/enhancements/issues/1904) Efficient watch resumption after kube-apiserver reboot

### [\#1904](https://github.com/kubernetes/enhancements/issues/1904) kube-apiserver 重启后的高效手表恢复

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** api-machinery, storage

**功能组：** api-machinery，存储

From now on, kube-apiserver can initialize its _watch cache_ faster after a reboot. This feature can be enabled with the `EfficientWatchResumption` feature flag.

从现在开始，kube-apiserver 可以在重启后更快地初始化其 _watch cache_。可以使用“EfficientWatchResumption”功能标志启用此功能。

Clients can keep track of the changes in Kubernetes objects with a watch. To serve these requests, `kube-apiserver` keeps a watch cache. But after a reboot (e.g., during rolling upgrades) that cache is often outdated, so `kube-apiserver` needs to fetch the updated state from etcd.

客户端可以通过手表跟踪 Kubernetes 对象的变化。为了服务这些请求，`kube-apiserver` 保留了一个监视缓存。但是在重新启动后（例如，在滚动升级期间），缓存通常已经过时，因此 `kube-apiserver` 需要从 etcd 中获取更新的状态。

The updated implementation will leverage the `WithProgressNotify` option from etcd version 3.0. `WithProgressNotify` will enable getting the updated state for the objects every minute, and also once before shutting down. This way, when `kube-apiserver` restarts, its cache will already be fairly updated.

更新后的实现将利用 etcd 3.0 版中的 `WithProgressNotify` 选项。 `WithProgressNotify` 将能够每分钟获取对象的更新状态，并且在关闭之前也会获取一次。这样，当 `kube-apiserver` 重新启动时，它的缓存就已经相当更新了。

Check the full implementation details [in the KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1904-efficient-watch-resumption).

检查完整的实现细节 [在 KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1904-efficient-watch-resumption)。

### [\#1929](https://github.com/kubernetes/enhancements/issues/1929) Built-in API types defaults

### [\#1929](https://github.com/kubernetes/enhancements/issues/1929) 内置 API 类型默认值

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** api-machinery

**功能组：** api-machinery

When building custom resource definitions (CRDs) as Go structures, you will now be able to use the `// +default` marker to provide default values. These markers will be translated into OpenAPI `default` fields.

将自定义资源定义 (CRD) 构建为 Go 结构时，您现在可以使用 `// +default` 标记来提供默认值。这些标记将被转换为 OpenAPI `default` 字段。

```
type Foo struct {
// +default=32
Integer int
// +default="bar"
String string
// +default=["popcorn", "chips"]
StringList []string
}

```


### [\#1040](https://github.com/kubernetes/enhancements/issues/1040) Priority and fairness for API server requests

### [\#1040](https://github.com/kubernetes/enhancements/issues/1040) API 服务器请求的优先级和公平性

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** api-machinery 

**功能组：** api-machinery

The `APIPriorityAndFairness` feature gate enables a new max-in-flight request handler in the API server. By defining different types of requests with `FlowSchema` objects and assigning them resources with `RequestPriority` objects, you can ensure the Kubernetes API server will be responsive for admin and maintenance tasks during high loads.

`APIPriorityAndFairness` 特性门在 API 服务器中启用了一个新的 max-in-flight 请求处理程序。通过使用 FlowSchema 对象定义不同类型的请求并使用 RequestPriority 对象为它们分配资源，您可以确保 Kubernetes API 服务器在高负载期间能够响应管理和维护任务。

Read more on the release for 1.18 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1040) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1040) 系列中阅读有关 1.18 版本的更多信息。

## Auth in Kubernetes 1.20

## 在 Kubernetes 1.20 中进行身份验证

### [\#541](https://github.com/kubernetes/enhancements/issues/541) External client-go credential providers

### [\#541](https://github.com/kubernetes/enhancements/issues/541) 外部客户端凭证提供程序

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** auth

**功能组：** auth

This enhancement allows Go clients to authenticate using external credential providers, like Key Management Systems (KMS), Trusted Platform Modules (TPM), or Hardware Security Modules (HSM).

此增强功能允许 Go 客户端使用外部凭据提供程序进行身份验证，例如密钥管理系统 (KMS)、可信平台模块 (TPM) 或硬件安全模块 (HSM)。

Those devices are already used to authenticate against other services, are easier to rotate, and are more secure as they don’t exist as files on the disk.

这些设备已经用于针对其他服务进行身份验证，更容易轮换，并且更安全，因为它们不作为磁盘上的文件存在。

Initially introduced on Kubernetes 1.10, this feature finally reaches Beta status.

最初在 Kubernetes 1.10 上引入，此功能最终达到 Beta 状态。

### [\#542](https://github.com/kubernetes/enhancements/issues/542) TokenRequest API and Kubelet integration

### [\#542](https://github.com/kubernetes/enhancements/issues/542) TokenRequest API 和 Kubelet 集成

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** auth

**功能组：** auth

The current JSON Web Tokens (JWT) that workloads use to authenticate against the API have some security issues. This enhancement comprises the work to create a more secure API for JWT.

工作负载用于针对 API 进行身份验证的当前 JSON Web 令牌 (JWT) 存在一些安全问题。此增强包括为 JWT 创建更安全的 API 的工作。

In the new API, tokens can be bound to a specific workload, given a validity duration, and exist only while a given object exists.

在新的 API 中，令牌可以绑定到特定的工作负载，给定有效期限，并且仅在给定对象存在时才存在。

Check the full details [in the KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/1205-bound-service-account-tokens).

检查 [在 KEP 中] 的全部细节（https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/1205-bound-service-account-tokens）。

### [\#1393](https://github.com/kubernetes/enhancements/issues/1393) Provide OIDC discovery for service account token issuer

### [\#1393](https://github.com/kubernetes/enhancements/issues/1393) 为服务帐户令牌颁发者提供 OIDC 发现

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** auth

**功能组：** auth

Kubernetes service accounts (KSA) can currently use JSON Web Tokens (JWT) to authenticate against the Kubernetes API, using `kubectl --token <the_token_string> ` for example. This enhancement allows services outside the cluster to use these tokens as a general authentication method without overloading the API server.

Kubernetes 服务帐户 (KSA) 目前可以使用 JSON Web 令牌 (JWT) 来针对 Kubernetes API 进行身份验证，例如使用 `kubectl --token <the_token_string> `。此增强功能允许集群外的服务将这些令牌用作通用身份验证方法，而不会使 API 服务器过载。

Read more on the release for 1.18 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1393) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1393) 系列中阅读有关 1.18 版本的更多信息。

## Kubernetes 1.20 CLI

## Kubernetes 1.20 CLI

### [\#1020](https://github.com/kubernetes/enhancements/issues/1020) Moving kubectl package code to staging

### [\#1020](https://github.com/kubernetes/enhancements/issues/1020) 将 kubectl 包代码移动到 staging

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** cli

**功能组：** cli

Continuing the work done in [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1020), this internal restructuring of the `kubectl` code is the first step to move the kubectl binary into [its own repository](https://github.com/kubernetes/kubectl). This helped decouple kubectl from the kubernetes code base and made it easier for out-the-tree projects to reuse its code.

继续在 [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1020) 中完成的工作，对 `kubectl` 代码的这种内部重构是移动kubectl 二进制文件到 [它自己的存储库](https://github.com/kubernetes/kubectl)。这有助于将 kubectl 与 kubernetes 代码库分离，并使树外项目更容易重用其代码。

### [\#1441](https://github.com/kubernetes/enhancements/issues/1441) kubectl debug

### [\#1441](https://github.com/kubernetes/enhancements/issues/1441) kubectl 调试

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** cli

**功能组：** cli

There have been two major changes in `kubectl debug` since [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1441).

自 [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1441) 以来，`kubectl debug` 有两个主要变化。

You can now use `kubectl debug` instead of `kubectl alpha debug`.

您现在可以使用 `kubectl debug` 而不是 `kubectl alpha debug`。

Also, `kubectl debug` can now change container images when copying a pod for debugging, similar to how `kubectl set image` works.

此外，“kubectl debug”现在可以在复制 pod 进行调试时更改容器映像，类似于“kubectl set image”的工作方式。

## Cloud providers in Kubernetes 1.20

## Kubernetes 1.20 中的云提供商

### [\#2133](https://github.com/kubernetes/enhancements/issues/2133) Kubelet credential provider

### [\#2133](https://github.com/kubernetes/enhancements/issues/2133) Kubelet 凭证提供程序

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** cloud-provider

**功能组：** 云提供商

This enhancement comprises the work done to move cloud provider specific SDKs out-of-tree. In particular, it replaces in-tree container image registry credential providers with a new mechanism that is external and pluggable.

此增强功能包括将云提供商特定的 SDK 移出树外所做的工作。特别是，它用一种外部可插拔的新机制替换了树内容器映像注册表凭据提供程序。

You can get more information on how to build a credential provider [in the KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cloud-provider/20191004-out-of-tree-credential-providers.md).

您可以获得有关如何构建凭证提供程序的更多信息 [在 KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cloud-provider/20191004-out-of-tree-凭证提供者.md)。

### [\#667](https://github.com/kubernetes/enhancements/issues/667) Support out-of-tree Azure cloud provider

### [\#667](https://github.com/kubernetes/enhancements/issues/667) 支持树外 Azure 云提供商

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** cloud-provider

**功能组：** 云提供商

This enhancement contains the work done to move the Azure cloud provider out-of-tree while keeping feature parity with the existing code in kube-controller-manager.

此增强功能包含将 Azure 云提供商移出树外的工作，同时保持与 kube-controller-manager 中现有代码的功能相同。

## Kubernetes 1.20 instrumentation 

## Kubernetes 1.20 检测

### [\#1748](https://github.com/kubernetes/enhancements/issues/1748) Expose metrics about resource requests and limits that represent the pod model

### [\#1748](https://github.com/kubernetes/enhancements/issues/1748) 公开有关代表 pod 模型的资源请求和限制的指标

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** instrumentation

**功能组：**仪器

The `kube-scheduler` now exposes more metrics on the requested resources and the desired limits of all running pods. This will help cluster administrators better plan capacity and triage errors.

`kube-scheduler` 现在公开了更多关于请求资源的指标以及所有正在运行的 pod 的所需限制。这将帮助集群管理员更好地规划容量和分类错误。

The metrics are exposed at the HTTP endpoint `/metrics/resources` when you use the `--show-hidden-metrics-for-version=1.20` flag.

当您使用 `--show-hidden-metrics-for-version=1.20` 标志时，指标会在 HTTP 端点 `/metrics/resources` 处公开。

### [\#1753](https://github.com/kubernetes/enhancements/issues/1753) Kubernetes system components logs sanitization

### [\#1753](https://github.com/kubernetes/enhancements/issues/1753) Kubernetes 系统组件日志清理

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** instrumentation

**功能组：**仪器

This enhancement aims to avoid sensitive data like passwords and tokens from being leaked in the Kubernetes log.

此增强功能旨在避免密码和令牌等敏感数据在 Kubernetes 日志中泄露。

Sensitive fields have been tagged:

```
type ConfigMap struct {
Data map[string]string `json:"data,omitempty" datapolicy:"password,token,security-key"`
}

```


So that they can be redacted in the log output.

以便它们可以在日志输出中进行编辑。

The sanitization filter kicks in when the `--experimental-logging-sanitization` flag is used. However, be aware that the current implementation takes a noticeable performance hit, and that sensitive data can still be leaked by user workloads.

当使用 `--experimental-logging-sanitization` 标志时，清理过滤器就会启动。但是，请注意，当前的实现对性能造成了显着影响，敏感数据仍可能因用户工作负载而泄露。

### [\#1933](https://github.com/kubernetes/enhancements/issues/1933) Defend against logging secrets via static analysis

### [\#1933](https://github.com/kubernetes/enhancements/issues/1933) 通过静态分析防御记录机密

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** instrumentation

**功能组：**仪器

While preparing the previous enhancement, [#1753](http://sysdig.com#1753), _go-flow-levee_ was used to provide insight on where sensitive data is used.

在准备之前的增强功能 [#1753](http://sysdig.com#1753) 时，_go-flow-levee_ 用于提供有关使用敏感数据的位置的洞察力。

## Network in Kubernetes 1.20

## Kubernetes 1.20 中的网络

### [\#1435](https://github.com/kubernetes/enhancements/issues/1435) Support of mixed protocols in Services with type=LoadBalancer

### [\#1435](https://github.com/kubernetes/enhancements/issues/1435) 支持类型=LoadBalancer 的服务中的混合协议

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** network

**功能组：**网络

The current LoadBalancer Service implementation does not allow different protocols under the same port (UDP, TCP). The rationale behind this is to avoid negative surprises with the load balancer bills in some cloud implementations.

当前的 LoadBalancer Service 实现不允许在同一端口（UDP、TCP）下使用不同的协议。这背后的基本原理是为了避免某些云实施中的负载均衡器账单出现负面意外。

However, a user might want to serve both UDP and TCP requests for a DNS or SIP server on the same port.

但是，用户可能希望在同一端口上同时为 DNS 或 SIP 服务器提供 UDP 和 TCP 请求。

This enhancement comprises the work to remove this limitation, and investigate the effects on billing on different clouds.

此增强功能包括消除此限制的工作，并调查对不同云上计费的影响。

### [\#1864](https://github.com/kubernetes/enhancements/issues/1864) Optionally Disable NodePorts for Service Type=LoadBalancer

### [\#1864](https://github.com/kubernetes/enhancements/issues/1864) 可选择为 Service Type=LoadBalancer 禁用 NodePorts

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** network

**功能组：**网络

Some implementations of the LoadBalancer API do not consume the node ports automatically allocated by Kubernetes, like MetalLB or kube-router. However, the API requires a port to be defined and allocated.

LoadBalancer API 的一些实现不使用 Kubernetes 自动分配的节点端口，如 MetalLB 或 kube-router。但是，API 需要定义和分配端口。

As a result, the number of load balancers is limited to the number of available ports. Also, these allocated but unused ports are exposed, which impact compliance requirements.

因此，负载均衡器的数量受限于可用端口的数量。此外，这些已分配但未使用的端口会暴露出来，这会影响合规性要求。

To solve this, the field `allocateLoadBalancerNodePort` has been added to `Service.Spec`. When set to `true`, it behaves as usual; when set to `false`, it will stop allocating new node ports (it won’t deallocate the existing ones, however).

为了解决这个问题，在`Service.Spec`中添加了字段`allocateLoadBalancerNodePort`。当设置为 `true` 时，它的行为和往常一样；当设置为 `false` 时，它将停止分配新的节点端口（但是它不会释放现有的端口）。

### [\#1672](https://github.com/kubernetes/enhancements/issues/1672) Tracking terminating Endpoints

### [\#1672](https://github.com/kubernetes/enhancements/issues/1672) 跟踪终止端点

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** network

**功能组：**网络

When Kubelet starts a graceful shutdown, the shutting-down Pod is removed from Endpoints (and, if enabled, EndpointSlice). If a consumer of the EndpointSlice API wants to know what Pods are terminating, it will have to watch the Pods directly. This complicates things, and has some scalability implications.

当 Kubelet 开始正常关闭时，正在关闭的 Pod 将从 Endpoints（以及，如果启用，EndpointSlice）中删除。如果 EndpointSlice API 的使用者想知道哪些 Pod 正在终止，则必须直接观察 Pod。这使事情变得复杂，并且具有一些可扩展性影响。

With this enhancement, a new `Terminating` field has been added to the EndpointConditions struct. Thus, terminating pods can be kept in the EndpointSlice API.

通过此增强功能，EndpointConditions 结构中添加了一个新的“Terminating”字段。因此，终止 pod 可以保留在 EndpointSlice API 中。

In the long term, this will enable services to handle terminating pods more intelligently. For example, IPVS proxier could set the weight of an endpoint to `0` during termination, instead of guessing that information from the connection tracking table.

从长远来看，这将使服务能够更智能地处理终止 pod。例如，IPVS 代理可以在终止期间将端点的权重设置为“0”，而不是从连接跟踪表中猜测该信息。

### [\#614](https://github.com/kubernetes/enhancements/issues/614) SCTP support for Services, Pod, Endpoint, and NetworkPolicy

### [\#614](https://github.com/kubernetes/enhancements/issues/614) 对服务、Pod、端点和网络策略的 SCTP 支持

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** network

**功能组：**网络

SCTP is now supported as an additional protocol alongside TCP and UDP in Pod, Service, Endpoint, and NetworkPolicy. 

在 Pod、Service、Endpoint 和 NetworkPolicy 中，现在支持 SCTP 作为与 TCP 和 UDP 一起的附加协议。

Introduced in [Kubernetes 1.12](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#614httpsgithubcomkubernetesfeaturesissues614sctpsupportforservicespodendpointandnetworkpolicyalpha), this feature finally graduates to _Stable_.

在[Kubernetes 1.12](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#614httpsgithubcomkubernetesfeaturesissues614sctpsupportforservicespodendpointandnetworkpolicyalpha)中引入，这个特性终于毕业了_Stable_。

### [\#1507](https://github.com/kubernetes/enhancements/issues/1507) Adding AppProtocol to Services and Endpoints

### [\#1507](https://github.com/kubernetes/enhancements/issues/1507) 将 AppProtocol 添加到服务和端点

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** network

**功能组：**网络

The [EndpointSlice API](https://sysdig.com/blog/whats-new-kubernetes-1-16/#752) added a new AppProtocol field in Kubernetes 1.17 to allow application protocols to be specified for each port. This enhancement brings that field into the `ServicePort` and `EndpointPort` resources, replacing non-standard annotations that are causing a [bad user experience](https://github.com/kubernetes/kubernetes/issues/40244).

[EndpointSlice API](https://sysdig.com/blog/whats-new-kubernetes-1-16/#752) 在 Kubernetes 1.17 中添加了一个新的 AppProtocol 字段，允许为每个端口指定应用协议。此增强功能将该字段引入`ServicePort` 和`EndpointPort` 资源，替换导致[不良用户体验](https://github.com/kubernetes/kubernetes/issues/40244) 的非标准注释。

Initially introduced in [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1507), this enhancement now graduates to _Stable_.

最初在 [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1507) 中引入，此增强功能现已升级为 _Stable_。

### [\#752](https://github.com/kubernetes/enhancements/issues/752) EndpointSlice API

### [\#752](https://github.com/kubernetes/enhancements/issues/752) EndpointSlice API

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** network

**功能组：**网络

The new `EndpointSlice` API will split endpoints into several Endpoint Slice resources. This solves many problems in the current API that are related to big `Endpoints` objects. This new API is also designed to support other future features, like multiple IPs per pod.

新的“EndpointSlice”API 将端点拆分为多个 Endpoint Slice 资源。这解决了当前 API 中与大型“端点”对象相关的许多问题。这个新 API 还旨在支持其他未来功能，例如每个 pod 多个 IP。

In Kubernetes 1.20, `Topology` has been deprecated, and a new `NodeName` field has been added. Check [the roll out plan](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/0752-endpointslices#roll-out-plan) for more info.

在 Kubernetes 1.20 中，`Topology` 已被弃用，并添加了一个新的 `NodeName` 字段。查看 [推出计划](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/0752-endpointslices#roll-out-plan) 了解更多信息。

Read more on the release for 1.16 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#752) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-16/#752) 系列中阅读有关 1.16 版本的更多信息。

### [\#563](https://github.com/kubernetes/enhancements/issues/563) Add IPv4/IPv6 dual-stack support

### [\#563](https://github.com/kubernetes/enhancements/issues/563) 添加 IPv4/IPv6 双栈支持

**Stage:** Graduating to Alpha

**阶段：** Alpha 毕业

**Feature group:** network

**功能组：**网络

This feature summarizes the work done to natively support dual-stack mode in your cluster, so you can assign both IPv4 and IPv6 addresses to a given pod.

此功能总结了为在集群中原生支持双栈模式所做的工作，因此您可以将 IPv4 和 IPv6 地址分配给给定的 pod。

Read more on the release for 1.16 in the “ [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#563)” series.

在“[Kubernetes 的新功能](https://sysdig.com/blog/whats-new-kubernetes-1-16/#563)”系列中阅读有关 1.16 版本的更多信息。

A big overhaul of this feature has taken place in Kubernetes 1.20 with breaking changes, which is why this feature is still on Alpha.

在 Kubernetes 1.20 中对该功能进行了大修，并进行了重大更改，这就是该功能仍在 Alpha 上的原因。

The major user-facing change is the introduction of the `.spec.ipFamilyPolicy` field. It can be set to: `SingleStack`, `PreferDualStack`, and `RequireDualStack`.

面向用户的主要变化是引入了`.spec.ipFamilyPolicy` 字段。它可以设置为：`SingleStack`、`PreferDualStack` 和`RequireDualStack`。

Dual stack is a big project that has implications on many Kubernetes services, so expect new improvements and changes before it leaves the alpha stage.

双栈是一个对许多 Kubernetes 服务都有影响的大项目，所以在它离开 alpha 阶段之前期待新的改进和变化。

## Kubernetes 1.20 nodes

## Kubernetes 1.20 节点

### [\#1967](https://github.com/kubernetes/enhancements/issues/1967) Support to size memory-backed volumes

### [\#1967](https://github.com/kubernetes/enhancements/issues/1967) 支持调整内存支持的卷大小

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** node

**功能组：**节点

When a Pod defines a memory-backed empty dir volume, (e.g., tmpfs) not all hosts size this volume equally. For example, a Linux host sizes it to 50% of the memory on the host.

当 Pod 定义一个内存支持的空目录卷（例如，tmpfs）时，并非所有主机都将此卷的大小相同。例如，Linux 主机将其大小调整为主机内存的 50%。

This has implications at several levels. For example, the Pod is less portable, as it’s behavior depends on the host it’s being deployed on. Also, this memory usage is not transparent to the eviction heuristics.

这在几个层面上都有影响。例如，Pod 的可移植性较差，因为它的行为取决于部署它的主机。此外，这种内存使用对驱逐试探法来说并不透明。

This new enhancement, after enabling the `SizeMemoryBackedVolumes` feature gate, will size these volumes not only with the node allocatable memory in mind, but also with the pod allocatable memory and the `emptyDir.sizeLimit` field.

在启用“SizeMemoryBackedVolumes”功能门后，这项新增强功能不仅会考虑节点可分配内存，还会考虑 pod 可分配内存和“emptyDir.sizeLimit”字段来调整这些卷的大小。

### [\#1972](https://github.com/kubernetes/enhancements/issues/1972) Fixing Kubelet exec probe timeouts

### [\#1972](https://github.com/kubernetes/enhancements/issues/1972) 修复 Kubelet exec 探测超时

**Stage:** Stable

**阶段：**稳定

**Feature group:** node

**功能组：**节点

Now, exec probes in Kubelet will respect the `timeoutSeconds` field.

现在，Kubelet 中的 exec 探测将尊重 `timeoutSeconds` 字段。

Since this is a bugfix, this feature graduates directly to Stable. You can roll back to the old behavior by disabling the `ExecProbeTimeout` feature gate.

由于这是一个错误修正，此功能直接升级到稳定版。您可以通过禁用“ExecProbeTimeout”功能门来回滚到旧行为。

### [\#2000](https://github.com/kubernetes/enhancements/issues/2000) Graceful node shutdown

### [\#2000](https://github.com/kubernetes/enhancements/issues/2000) 优雅的节点关闭

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** node

**功能组：**节点

With the `GracefulNodeShutdown` feature gate enabled, Kubelet will try to gracefully terminate the pods running in the node when shutting down. The implementation works by listening to systemd inhibitor locks (for Linux).

启用 GracefulNodeShutdown 功能门后，Kubelet 将尝试在关闭时优雅地终止节点中运行的 Pod。该实现通过侦听 systemd 抑制锁（适用于 Linux）来工作。

The new Kubelet config setting `kubeletConfig.ShutdownGracePeriod` will dictate how much time Pods have to terminate gracefully. 

新的 Kubelet 配置设置 `kubeletConfig.ShutdownGracePeriod` 将规定 Pod 必须正常终止的时间。

This enhancement can mitigate issues in some workloads, making life easier for admins and developers. For example, a cold shutdown can corrupt open files, or leave resources in an undesired state.

此增强功能可以缓解某些工作负载中的问题，使管理员和开发人员的工作更轻松。例如，冷关机会损坏打开的文件，或使资源处于不希望的状态。

### [\#2053](https://github.com/kubernetes/enhancements/issues/2053) Add downward API support for hugepages

### [\#2053](https://github.com/kubernetes/enhancements/issues/2053) 为大页面添加向下 API 支持

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** node

**功能组：**节点

If enabled via the `DownwardAPIHugePages` feature gate, Pods will be able to fetch information on their hugepage requests and limits via the downward API. This keeps things consistent with other resources like cpu, memory, and ephemeral-storage.

如果通过 `DownwardAPIHugePages` 功能门启用，Pods 将能够通过向下 API 获取有关其大页面请求和限制的信息。这使事物与 CPU、内存和临时存储等其他资源保持一致。

### [\#585](https://github.com/kubernetes/enhancements/issues/585) RuntimeClass

### [\#585](https://github.com/kubernetes/enhancements/issues/585) RuntimeClass

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** node

**功能组：**节点

The `RuntimeClass` resource provides a mechanism for supporting multiple runtimes in a cluster (Docker, rkt, gVisor, etc.), and surfaces information about that container runtime to the control plane.

`RuntimeClass` 资源提供了一种机制来支持集群中的多个运行时（Docker、rkt、gVisor 等），并将有关该容器运行时的信息显示到控制平面。

For example:

```
apiVersion: v1
kind: Pod
metadata:
name: mypod
spec:
runtimeClassName: sandboxed
# ...

```


This YAML will instruct the Kubelet to use the `sandboxed` RuntimeClass to run this pod. This selector is required to dynamically adapt to multiple container runtime engines beyond Docker, like rkt or gVisor.

此 YAML 将指示 Kubelet 使用“沙盒”运行时类来运行此 pod。这个选择器需要动态适应 Docker 之外的多个容器运行时引擎，比如 rkt 或 gVisor。

Introduced on [Kubernetes 1.12](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#585httpsgithubcomkubernetesfeaturesissues585runtimeclassalpha), this feature finally graduates to Stable.

在[Kubernetes 1.12](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#585httpsgithubcomkubernetesfeaturesissues585runtimeclassalpha)上引入，这个特性最终升级到稳定。

### [\#606](https://github.com/kubernetes/enhancements/issues/606) Support 3rd party device monitoring plugins

### [\#606](https://github.com/kubernetes/enhancements/issues/606) 支持 3rd 方设备监控插件

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** node

**功能组：**节点

This feature allows the Kubelet to expose container bindings to third-party monitoring plugins.

此功能允许 Kubelet 将容器绑定公开给第三方监控插件。

With this implementation, administrators will be able to monitor the custom resource assignment to containers using a third-party Device Monitoring Agent (For example, percent GPU use per pod).

通过此实施，管理员将能够使用第三方设备监控代理（例如，每个 Pod 的 GPU 使用百分比）监控容器的自定义资源分配。

Read more on the release for 1.15 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-15/#606) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-15/#606) 系列中阅读有关 1.15 版本的更多信息。

### [\#757](https://github.com/kubernetes/enhancements/issues/757) PID limiting

### [\#757](https://github.com/kubernetes/enhancements/issues/757) PID 限制

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** node

**功能组：**节点

PIDs are a fundamental resource on any host. Administrators require mechanisms to ensure that user pods cannot induce pid exhaustion that may prevent host daemons (runtime, Kubelet, etc.) from running.

PID 是任何主机上的基本资源。管理员需要机制来确保用户 pod 不会导致 pid 耗尽，这可能会阻止主机守护程序（运行时、Kubelet 等）运行。

This feature allows for the configuration of a Kubelet to limit the number of PIDs a given pod can consume. Node-level support for PID limiting no longer requires setting the feature gate `SupportNodePidsLimit=true` explicitly.

此功能允许配置 Kubelet 以限制给定 pod 可以使用的 PID 数量。节点级对 PID 限制的支持不再需要显式设置特征门 `SupportNodePidsLimit=true`。

Read more on the release for 1.15 in the [What's new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-15/#757) series, and [in the Kubernetes blog](https://kubernetes.io/blog/2019/04/15/process-id-limiting-for-stability-improvements-in-kubernetes-1.14/).

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-15/#757) 系列和 [在 Kubernetes 博客中](https://kubernetes.io/blog/2019/04/15/process-id-limiting-for-stability-improvements-in-kubernetes-1.14/)。

### [\#950](https://github.com/kubernetes/enhancements/issues/950) Add pod-startup liveness-probe holdoff for slow-starting pods

### [\#950](https://github.com/kubernetes/enhancements/issues/950) 为缓慢启动的 pod 添加 pod-startup liveness-probe 延迟

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** node

**功能组：**节点

Probes allow Kubernetes to monitor the status of your applications. If a pod takes too long to start, those probes might think the pod is dead, causing a restart loop. This feature lets you define a `startupProbe` that will hold off all of the other probes until the pod finishes its startup. For example: _“Don’t test for liveness until a given HTTP endpoint is available_. _“_

探针允许 Kubernetes 监控应用程序的状态。如果 Pod 启动时间过长，这些探测器可能会认为 Pod 已死，从而导致重新启动循环。此功能允许您定义一个 `startupProbe`，它将阻止所有其他探测器，直到 pod 完成其启动。例如：_“在给定的 HTTP 端点可用之前不要测试活跃度_。 _“_

Read more on the release for 1.16 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#950) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-16/#950) 系列中阅读有关 1.16 版本的更多信息。

### [\#693](https://github.com/kubernetes/enhancements/issues/693) Node topology manager

### [\#693](https://github.com/kubernetes/enhancements/issues/693) 节点拓扑管理器

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** node

**功能组：**节点

Machine learning, scientific computing, and financial services are examples of systems that are computational intensive and require ultra-low latency. These kinds of workloads benefit from isolated processes to one CPU core rather than jumping between cores or sharing time with other processes. 

机器学习、科学计算和金融服务是计算密集型和需要超低延迟的系统的例子。这些类型的工作负载受益于隔离进程到一个 CPU 内核，而不是在内核之间跳转或与其他进程共享时间。

The node topology manager is a `kubelet` component that centralizes the coordination of hardware resource assignments. The current approach divides this task between several components (CPU manager, device manager, and CNI), which sometimes results in unoptimized allocations.

节点拓扑管理器是一个 `kubelet` 组件，它集中协调硬件资源分配。当前的方法将这个任务划分到多个组件（CPU 管理器、设备管理器和 CNI）之间，这有时会导致未优化的分配。

In Kubernetes 1.20, you can define the `scope` where topology hints should be collected on a `container`-by-container basis, or on a `pod`-by-pod basis.

在 Kubernetes 1.20 中，您可以定义“范围”，拓扑提示应在“容器”逐个容器的基础上或在“pod”逐个容器的基础上收集。

Read more on the release for 1.16 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#693) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-16/#693) 系列中阅读有关 1.16 版本的更多信息。

### [\#1797](https://github.com/kubernetes/enhancements/issues/1797) Allow users to set a pod’s hostname to its Fully Qualified Domain Name (FQDN)

### [\#1797](https://github.com/kubernetes/enhancements/issues/1797) 允许用户将 pod 的主机名设置为其完全限定域名 (FQDN)

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** node

**功能组：**节点

Now, it’s possible to set a pod’s hostname to its Fully Qualified Domain Name (FQDN), which increases the interoperability of Kubernetes with legacy applications.

现在，可以将 pod 的主机名设置为其完全限定域名 (FQDN)，这增加了 Kubernetes 与遗留应用程序的互操作性。

After setting `hostnameFQDN: true`, running `uname -n` inside a Pod returns `foo.test.bar.svc.cluster.local` instead of just `foo`.

设置 `hostnameFQDN: true` 后，在 Pod 内运行 `uname -n` 会返回 `foo.test.bar.svc.cluster.local` 而不仅仅是 `foo`。

This feature was introduced on [Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1797), and you can read more details [in the enhancement proposal](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/1797-configure-fqdn-as-hostname-for-pods/README.md).

这个特性是在[Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1797)上引入的，你可以阅读更多细节[在增强提案中](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/1797-configure-fqdn-as-hostname-for-pods/README.md)。

### [\#1867](https://github.com/kubernetes/enhancements/issues/1867) Kubelet feature: Disable AcceleratorUsage metrics

### [\#1867](https://github.com/kubernetes/enhancements/issues/1867) Kubelet 功能：禁用 AcceleratorUsage 指标

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** node

**功能组：**节点

With [#606 (Support third-party device monitoring plugins)](https://sysdig.com/blog/whats-new-kubernetes-1-15/#606) and the PodResources API about to enter GA, it isn' t expected for Kubelet to gather metrics anymore.

随着[#606（支持第三方设备监控插件）](https://sysdig.com/blog/whats-new-kubernetes-1-15/#606)和即将进入GA的PodResources API，它是预计 Kubelet 将不再收集指标。

Introduced on [Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1867), this enhancement summarizes the process to deprecate Kubelet collecting those Accelerator Metrics.

在 [Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1867) 上引入，此增强功能总结了弃用 Kubelet 收集这些加速器指标的过程。

### [\#2040](https://github.com/kubernetes/enhancements/issues/2040) Kubelet CRI support

### [\#2040](https://github.com/kubernetes/enhancements/issues/2040) Kubelet CRI 支持

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** node

**功能组：**节点

Kubernetes introduced support for the Container Runtime Interface (CRI) [as soon as Kubernetes 1.5](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/). It is a plugin interface that enables Kubelet to use a wide variety of container runtimes, without the need to recompile. This includes alternatives to Docker like CRI-O or containerd.

Kubernetes 引入了对容器运行时接口 (CRI) 的支持 [从 Kubernetes 1.5 开始](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/)。它是一个插件接口，使 Kubelet 能够使用各种容器运行时，而无需重新编译。这包括 Docker 的替代品，如 CRI-O 或 containerd。

Although the CRI API has been widely tested in production, it was still officially in Alpha.

尽管 CRI API 已经在生产中进行了广泛的测试，但它仍处于正式的 Alpha 阶段。

This enhancement comprises the work to identify and close the remaining gaps, so that CRI can finally be promoted to Stable.

此增强包括识别和缩小剩余差距的工作，以便最终可以将 CRI 提升到稳定。

## Kubernetes 1.20 storage

## Kubernetes 1.20 存储

### [\#2047](https://github.com/kubernetes/enhancements/issues/2047) CSIServiceAccountToken

### [\#2047](https://github.com/kubernetes/enhancements/issues/2047) CSIServiceAccountToken

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** storage

**功能组：** 存储

CSI drivers [can impersonate the pods](https://kubernetes-csi.github.io/docs/token-requests.html) that they mount volumes for. This feature increases security by providing the CSI drivers with only the permissions they need. However, in the current implementation, the drivers read the service account tokens directly from the filesystem.

CSI 驱动程序 [可以模拟 pod](https://kubernetes-csi.github.io/docs/token-requests.html)，它们为其安装卷。此功能通过仅向 CSI 驱动程序提供他们需要的权限来提高安全性。但是，在当前的实现中，驱动程序直接从文件系统读取服务帐户令牌。

Some downsides of this are that the drivers need permissions to read the filesystem, which might give access to more secrets than needed, and that the token is not guaranteed to be available (e.g., `automountServiceAccountToken=false`).

这样做的一些缺点是驱动程序需要读取文件系统的权限，这可能允许访问比所需更多的秘密，并且不能保证令牌可用（例如，`automountServiceAccountToken=false`）。

With this enhancement, CSI drivers will be able to request the service account tokens from Kubelet to the `NodePublishVolume` function. Kubelet will also be able to limit what tokens are available to which driver. And finally, the driver will be able to re-execute `NodePublishVolume` to remount the volume by setting `RequiresRepublish` to `true`.

通过此增强功能，CSI 驱动程序将能够从 Kubelet 向 `NodePublishVolume` 函数请求服务帐户令牌。 Kubelet 还将能够限制哪些令牌可供哪个驱动程序使用。最后，驱动程序将能够通过将 `RequiresRepublish` 设置为 `true` 来重新执行 `NodePublishVolume` 以重新安装卷。

This last feature will come in handy when the mounted volumes can expire and need a re-login. For example, a secrets vault.

当安装的卷可能过期并需要重新登录时，最后一个功能将派上用场。例如，秘密保险库。

Check the full details [in the KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1855-csi-driver-service-account-token/README.md).

检查 [在 KEP 中] 的全部细节（https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1855-csi-driver-service-account-token/README.md）。

### [\#177](https://github.com/kubernetes/enhancements/issues/177) Snapshot / restore volume support for Kubernetes

### [\#177](https://github.com/kubernetes/enhancements/issues/177) Kubernetes 的快照/恢复卷支持

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** storage 

**功能组：** 存储

In alpha, since [the 1.12 Kubernetes release](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#177httpsgithubcomkubernetesfeaturesissues177snapshotrestorevolumesupportforkubernetescrdexternalcontrolleralpha), this feature finally graduates to Stable.

在 alpha 版本中，自从 [1.12 Kubernetes 版本](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#177httpsgithubcomkubernetesfeaturesissues177snapshotrestorevolumesupportforkubernetescrdexternalcontrolleralpha)，这个功能终于升级到稳定版。

Similar to how API resources `PersistentVolume` and `PersistentVolumeClaim` are used to provision volumes for users and administrators, `VolumeSnapshotContent` and `VolumeSnapshot` API resources can be provided to create volume snapshots for users and administrators. Read more [about volume snapshots here](https://github.com/xing-yang/website/blob/f53fe7ed8bf0ee98a1c45eb67bc505e309fdce0e/content/en/docs/concepts/storage/volume-snapshots.md).

类似于 API 资源 `PersistentVolume` 和 `PersistentVolumeClaim` 用于为用户和管理员配置卷，可以提供 `VolumeSnapshotContent` 和 `VolumeSnapshot` API 资源来为用户和管理员创建卷快照。阅读更多 [关于卷快照](https://github.com/xing-yang/website/blob/f53fe7ed8bf0ee98a1c45eb67bc505e309fdce0e/content/en/docs/concepts/storage/volume-snapshots.md)。

Read more on the release for 1.16 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#693) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-16/#693) 系列中阅读有关 1.16 版本的更多信息。

### [\#695](https://github.com/kubernetes/enhancements/issues/695) Skip volume ownership change

### [\#695](https://github.com/kubernetes/enhancements/issues/695) 跳过卷所有权更改

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** storage

**功能组：** 存储

Before a volume is bind-mounted inside a container, all of its file permissions are changed to the provided `fsGroup` value. This ends up being a slow process on very large volumes, and also breaks some permission sensitive applications, like databases.

在将卷绑定安装到容器内之前，其所有文件权限都更改为提供的“fsGroup”值。这最终在非常大的卷上成为一个缓慢的过程，并且还会破坏一些对权限敏感的应用程序，例如数据库。

Introduced in [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#695), the new `FSGroupChangePolicy` field has been added [to control this behavior](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/20200120-skip-permission-change.md). If set to Always, it will maintain the current behavior. However, when set to OnRootMismatch, it will only change the volume permissions if the top level directory does not match the expected `fsGroup` value.

[Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#695)中引入了新的`FSGroupChangePolicy`字段[控制这种行为](https:///github.com/kubernetes/enhancements/blob/master/keps/sig-storage/20200120-skip-permission-change.md)。如果设置为 Always，它将保持当前行为。但是，当设置为 OnRootMismatch 时，如果顶级目录与预期的 `fsGroup` 值不匹配，它只会更改卷权限。

### [\#1682](https://github.com/kubernetes/enhancements/issues/1682) Allow CSI drivers to opt-in to volume ownership change

### [\#1682](https://github.com/kubernetes/enhancements/issues/1682) 允许 CSI 驱动程序选择加入卷所有权更改

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** storage

**功能组：** 存储

Not all volumes support the `fsGroup` operations mentioned in the previous enhancement ( [#695](http://sysdig.com#695)), like NFS. This can result in errors reported to the user.

并非所有卷都支持之前增强功能 ([#695](http://sysdig.com#695)) 中提到的 `fsGroup` 操作，例如 NFS。这可能会导致向用户报告错误。

This enhancement adds a new field called `CSIDriver.Spec.SupportsFSGroup` that allows the driver to define if it supports volume ownership modifications via fsGroup.

此增强功能添加了一个名为“CSIDriver.Spec.SupportsFSGroup”的新字段，允许驱动程序定义它是否支持通过 fsGroup 修改卷所有权。

Read more on the release for 1.19 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1682) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1682) 系列中阅读有关 1.19 版本的更多信息。

## Miscellaneous

##  各种各样的

### [\#1610](https://github.com/kubernetes/enhancements/issues/1610) Container resource based Pod autoscaling

### [\#1610](https://github.com/kubernetes/enhancements/issues/1610) 基于容器资源的 Pod 自动伸缩

**Stage:** Alpha

**阶段：** Alpha

**Feature group:** autoscaling

**功能组：** 自动缩放

The current Horizontal Pod Autoscaler (HPA) can scale workloads based on the resources used by their Pods. This is the aggregated usage from all of the containers in a Pod.

当前的 Horizontal Pod Autoscaler (HPA) 可以根据其 Pod 使用的资源扩展工作负载。这是 Pod 中所有容器的聚合使用量。

[This feature](https://github.com/kubernetes/enhancements/blob/master/keps/sig-autoscaling/0001-container-resource-autoscaling.md) allows the HPA to scale those workloads based on the resource usage of individual containers:

```
type: ContainerResource
containerResource:
name: cpu
container: application
target:
     type: Utilization
     averageUtilization: 60

```


### [\#1001](https://github.com/kubernetes/enhancements/issues/1001) Support CRI-ContainerD on Windows

### [\#1001](https://github.com/kubernetes/enhancements/issues/1001) 在 Windows 上支持 CRI-ContainerD

**Stage:** Graduating to Stable

**阶段：** 毕业至稳定

**Feature group:** windows

**功能组：** 窗口

ContainerD is an OCI-compliant runtime that works with Kubernetes and has support for the host container service (HCS v2) in Windows Server 2019. This enhancement introduces ContainerD 1.3 support in Windows as a Container Runtime Interface (CRI).

ContainerD 是符合 OCI 的运行时，可与 Kubernetes 配合使用，并支持 Windows Server 2019 中的主机容器服务 (HCS v2)。此增强功能在 Windows 中引入了 ContainerD 1.3 支持，作为容器运行时接口 (CRI)。

Read more on the release for 1.18 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1001) series.

在 [Kubernetes 的新特性](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1001) 系列中阅读有关 1.18 版本的更多信息。

### [\#19](https://github.com/kubernetes/enhancements/issues/19) CronJobs (previously ScheduledJobs)

### [\#19](https://github.com/kubernetes/enhancements/issues/19) CronJobs（以前的 ScheduledJobs)

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** apps

**功能组：** 应用

Introduced in Kubernetes 1.4 and in beta since 1.8, CronJobs are finally on the road to become Stable.

在 Kubernetes 1.4 和自 1.8 以来的测试版中引入，CronJobs 终于走上了稳定的道路。

CronJobs runs periodic tasks in a Kubernetes cluster, similar to cron on UNIX-like systems.

CronJobs 在 Kubernetes 集群中运行周期性任务，类似于类 UNIX 系统上的 cron。

A new, alternate implementation of CronJobs [has been created](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/19-Graduate-CronJob-to-Stable) to address the main issues of the current code without breaking the current functionality.

CronJobs 的新替代实现 [已创建](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/19-Graduate-CronJob-to-Stable) 以解决主要问题在不破坏当前功能的情况下查看当前代码。

This new implementation will focus on scalability, providing more status information and addressing the current open issues. 

这种新的实施将侧重于可扩展性，提供更多状态信息并解决当前未解决的问题。

You can start testing the new CronJobs by setting the `CronJobControllerV2` feature flag to `true`.

您可以通过将 `CronJobControllerV2` 功能标志设置为 `true` 来开始测试新的 CronJobs。

### [\#1258](https://github.com/kubernetes/enhancements/issues/1258) Add a configurable default constraint to PodTopologySpread

### [\#1258](https://github.com/kubernetes/enhancements/issues/1258) 向 PodTopologySpread 添加可配置的默认约束

**Stage:** Graduating to Beta

**阶段：** 进入测试阶段

**Feature group:** scheduling

**功能组：**调度

In order to take advantage of [even pod spreading](https://docs.google.com/document/d/1ZDSHeySKoYYKnP_86rj2d5GRvYWn3QPKOu1XDWOJAeI/edit#895), each pod needs its own spreading rules and this can be a tedious task.

为了利用 [偶数 pod 传播](https://docs.google.com/document/d/1ZDSHeySKoYYKnP_86rj2d5GRvYWn3QPKOu1XDWOJAeI/edit#895)，每个 pod 都需要自己的传播规则，这可能是一项繁琐的任务。

Introduced in [Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1258), this enhancement allows you to define global `defaultConstraints` that will be applied at cluster level to all of the pods that don't define their own `topologySpreadConstraints`.

在 [Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1258) 中引入，此增强功能允许您定义全局 `defaultConstraints`，它将在集群级别应用于所有没有定义自己的 `topologySpreadConstraints` 的 pod。

* * *

* * *

That’s all, folks! Exciting as always; get ready to upgrade your clusters if you are intending to use any of these features.

就是这样，伙计们！一如既往的精彩；如果您打算使用这些功能中的任何一个，请准备好升级您的集群。

If you liked this, you might want to check out our previous ‘What’s new in Kubernetes’ editions:

- [What’s new in Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/)
- [What’s new in Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/)
- [What’s new in Kubernetes 1.17](https://sysdig.com/blog/whats-new-kubernetes-1-17/)
- [What’s new in Kubernetes 1.16](https://sysdig.com/blog/whats-new-kubernetes-1-16/)
- [What’s new in Kubernetes 1.15](https://sysdig.com/blog/whats-new-kubernetes-1-15/)
- [What’s new in Kubernetes 1.14](https://sysdig.com/blog/whats-new-kubernetes-1-14/)
- [What’s new in Kubernetes 1.13](https://sysdig.com/blog/whats-new-in-kubernetes-1-13)
- [What’s new in Kubernetes 1.12](https://sysdig.com/blog/whats-new-in-kubernetes-1-12)

如果你喜欢这个，你可能想看看我们以前的“Kubernetes 的新特性”版本：

- [Kubernetes 1.19 的新变化](https://sysdig.com/blog/whats-new-kubernetes-1-19/)
- [Kubernetes 1.18 的新变化](https://sysdig.com/blog/whats-new-kubernetes-1-18/)
- [Kubernetes 1.17 的新变化](https://sysdig.com/blog/whats-new-kubernetes-1-17/)
- [Kubernetes 1.16 的新变化](https://sysdig.com/blog/whats-new-kubernetes-1-16/)
- [Kubernetes 1.15 的新变化](https://sysdig.com/blog/whats-new-kubernetes-1-15/)
- [Kubernetes 1.14 的新变化](https://sysdig.com/blog/whats-new-kubernetes-1-14/)
- [Kubernetes 1.13 的新变化](https://sysdig.com/blog/whats-new-in-kubernetes-1-13)
- [Kubernetes 1.12 的新变化](https://sysdig.com/blog/whats-new-in-kubernetes-1-12)

And if you enjoy keeping up to date with the Kubernetes ecosystem, [subscribe to our container newsletter](https://go.sysdig.com/container-newsletter-signup.html), a monthly email with the coolest stuff happening in the cloud-native ecosystem. 

如果您喜欢了解 Kubernetes 生态系统的最新动态，请[订阅我们的容器通讯](https://go.sysdig.com/container-newsletter-signup.html)，这是一封每月发送的电子邮件，其中包含最酷的内容云原生生态系统。


