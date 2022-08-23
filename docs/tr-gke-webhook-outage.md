# How a simple admission webhook lead to a cluster outage

# 一个简单的 admission webhook 如何导致集群中断

5/Sep 2019 https://www.jetstack.io/blog/gke-webhook-outage/

Jetstack often works with customers to provision multi-tenant platforms on Kubernetes. Sometimes special requirements arise that we cannot control with stock Kubernetes configuration. In order to implement such requirements, we’ve recently started making use of the [Open Policy Agent](https://www.openpolicyagent.org/) project as an admission controller to enforce custom policies.

Jetstack 经常与客户合作，在 Kubernetes 上配置多租户平台。有时会出现我们无法通过库存 Kubernetes 配置控制的特殊要求。为了实现这些要求，我们最近开始使用 [Open Policy Agent](https://www.openpolicyagent.org/) 项目作为准入控制器来执行自定义策略。

This post is a write up of an incident caused by misconfiguration of this integration.

这篇文章是对由于此集成的错误配置引起的事件的记录。

## The Incident

## 事件

We were in the process of upgrading the master for a development cluster used by many teams to test their apps during the working day. This was a regional cluster running in europe-west1 on Google Kubernetes Engine (GKE).

我们正在为一个开发集群升级主服务器，许多团队在工作日使用它来测试他们的应用程序。这是在 Google Kubernetes Engine (GKE) 上运行在 europe-west1 的区域集群。

Teams had been warned that the upgrade was taking place though no API downtime was expected. We had previously performed the same upgrade on another pre-production environment earlier that day.

团队已被警告升级正在进行，但预计不会出现 API 停机时间。当天早些时候，我们之前在另一个预生产环境中执行了相同的升级。

We began the upgrade via our GKE Terraform pipeline. When performing the master upgrade the operation did not complete before the Terraform timeout (which we had set to 20 minutes). This was the first sign that something was wrong though the cluster was still showing as upgrading in the GKE console.

我们通过 GKE Terraform 管道开始升级。执行主升级时，操作未在 Terraform 超时（我们设置为 20 分钟）之前完成。尽管集群在 GKE 控制台中仍显示为升级，但这是出现问题的第一个迹象。

Re-running the pipeline resulted in the following error:

重新运行管道导致以下错误：

```
google_container_cluster.cluster: Error waiting for updating GKE master version:
All cluster resources were brought up, but the cluster API is reporting that:
component "kube-apiserver" from endpoint "gke-..." is unhealthy
```

During this time the API server was intermittently timing out and teams were not able to deploy their applications.

在此期间，API 服务器间歇性地超时，团队无法部署他们的应用程序。

While we began to investigate the issue, all the nodes started to be destroyed and recreated in an endless loop. This lead to an indiscriminate loss of service for all tenants.

当我们开始调查这个问题时，所有节点都开始被销毁并在无限循环中重新创建。这会导致所有租户不分青红皂白地失去服务。

## Identifying the Root Cause

## 确定根本原因

With the help of Google Support we identified the sequence of events that lead to the outage.

在 Google 支持的帮助下，我们确定了导致中断的事件顺序。

1. GKE completed the upgrade of one master instance, and started to receive all API server traffic as the following masters were upgraded.
2. During the upgrade of the second master instance, the API server was unable to run [`PostStartHook`](https://github.com/kubernetes/kubernetes/blob/e09f5c40b55c91f681a46ee17f9bc447eeacee57/pkg/master/client_ca_hook.go#L43) for [`ca-registration`](https://github.com/kubernetes/kubernetes/blob/e09f5c40b55c91f681a46ee17f9bc447eeacee57/pkg/master/client_ca_hook.go#L121).
3. While running this hook, the API server attempted to update a`ConfigMap` called `extension-apiserver-authentication` in `kube-system`. This operation timed out as the backend for the validating Open Policy Agent (OPA) webhook we had configured was not responding.
4. This operation must complete for a master to pass a health check, because it continuously failed the second master entered a crash loop and halted the upgrade.



1. GKE完成了一个master实例的升级，随着以下master的升级，开始接收所有的API server流量。
2. 第二个master实例升级过程中，API服务器无法运行[`PostStartHook`](https://github.com/kubernetes/kubernetes/blob/e09f5c40b55c91f681a46ee17f9bc447eeacee57/pkg/master/client_ca_hook.go#L43)对于[`ca-registration`](https://github.com/kubernetes/kubernetes/blob/e09f5c40b55c91f681a46ee17f9bc447eeacee57/pkg/master/client_ca_hook.go#L121)。
3. 在运行此钩子时，API 服务器尝试更新 `kube-system` 中名为 `extension-apiserver-authentication` 的`ConfigMap`。此操作超时，因为我们配置的验证开放策略代理 (OPA) webhook 的后端没有响应。
4. 这个操作必须完成，master 才能通过健康检查，因为它连续失败，第二个 master 进入崩溃循环并停止升级。

This had the knock on effect of intermittent API downtime which caused kubelets to be unable to report node health. This in turn caused GKE node auto-repair to recreate nodes as a means of repair. This feature is explained in the [documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-repair):

这产生了间歇性 API 停机的连锁反应，导致 kubelet 无法报告节点运行状况。这反过来又导致 GKE 节点自动修复以重新创建节点作为修复手段。此功能在 [文档](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-repair) 中有说明：

> An unhealthy status can mean: A node does not report any status at all over the given time threshold (approximately 10 minutes).

> 不健康状态可能意味着：节点在给定时间阈值（大约 10 分钟）内根本不报告任何状态。

## Resolution

解决

Once we had identified the webhook as the cause of the issue, with intermittent API server access, we were able to delete this `ValidatingAdmissionWebhook` resource to restore the cluster service.

一旦我们确定 Webhook 是问题的原因，并且 API 服务器访问间歇性，我们就可以删除此 ValidatingAdmissionWebhook 资源以恢复集群服务。

Since then, we have configured our `ValidatingAdmissionWebhook` for OPA to only monitor the namespaces where policy is applicable and development teams have access. We have also only enabled the webhook for `Ingress` and `Service` resources, the only resources validated by our policy.

从那时起，我们为 OPA 配置了“ValidatingAdmissionWebhook”，以仅监控适用于政策且开发团队有权访问的命名空间。我们还只为“Ingress”和“Service”资源启用了 webhook，这是我们政策验证的唯一资源。

Since we first deployed OPA, the documentation has been [updated](https://github.com/open-policy-agent/opa/pull/1435) to reflect this change.

自从我们首次部署 OPA 以来，文档已经 [更新](https://github.com/open-policy-agent/opa/pull/1435) 以反映这一变化。

We have also added a liveness probe to ensure OPA is restarted when it becomes unresponsive and have [updated the documentation](https://github.com/open-policy-agent/opa/pull/1605).

我们还添加了一个活性探针，以确保 OPA 在无响应时重新启动，并 [更新了文档](https://github.com/open-policy-agent/opa/pull/1605)。

We also considered but decided against disabling [GKE node auto-repair](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-repair).

我们也考虑过但决定不禁用 [GKE 节点 自动修复](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-repair)。

## Takeaways 

## 要点

If we had alerting on API server response times we might have noticed these increase across the board for all `CREATE` and `UPDATE` requests after deploying the OPA-backed webhook initially.

如果我们有关于 API 服务器响应时间的警报，我们可能会注意到在最初部署 OPA 支持的 webhook 后，所有“CREATE”和“UPDATE”请求的这些增加。

It also hammers home the value of configuring probes for all workloads. In hindsight, deploying OPA was so deceptively simple we didn’t reach for a [Helm chart](https://github.com/helm/charts/tree/master/stable/opa) when we maybe should have. This chart encodes a number of adjustments beyond the basic config in the tutorial - including a `livenessProbe` for the admission controller containers.

它还强调了为所有工作负载配置探针的价值。事后看来，部署 OPA 非常简单，以至于我们没有达到 [Helm 图表](https://github.com/helm/charts/tree/master/stable/opa)，而我们本应该有。此图表编码了教程中基本配置之外的许多调整 - 包括准入控制器容器的“livenessProbe”。

We were not the first to run into this issue, an upstream issue remains open [here](https://github.com/kubernetes/kubernetes/issues/54522) to improve the functionality here - we’ll be following this one.

我们不是第一个遇到这个问题的人，上游问题仍然开放 [这里](https://github.com/kubernetes/kubernetes/issues/54522) 以改进这里的功能 - 我们将关注这个问题。

------

#### Interested in learning about Kubernetes failure cases **without** disrupting your dev teams? Check out Flightdeck, available as part of [Jetstack Subscription](http://www.jetstack.io/subscription/). 

#### 有兴趣了解 Kubernetes 故障案例**而不**干扰您的开发团队？查看作为 [Jetstack 订阅](http://www.jetstack.io/subscription/) 的一部分提供的 Flightdeck。

