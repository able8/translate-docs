# 10 Best Practices Worth Implementing to Adopt Kubernetes

# 10 个值得实施以采用 Kubernetes 的最佳实践

September 25, 2020



We already know that [Kubernetes](https://kubernetes.io/) is the No. 1 orchestration platform for container-based applications, automating the deployment and scaling of these apps and streamlining maintenance operations. However, Kubernetes comes with its own complexity challenges. So how can an enterprise take advantage of containerization to tackle complexity and not end up with even more complexity? This article provides some of the best practices that you can implement to adopt Kubernetes.

我们已经知道 [Kubernetes](https://kubernetes.io/) 是基于容器的应用程序的第一编排平台，可自动部署和扩展这些应用程序并简化维护操作。然而，Kubernetes 也有其自身的复杂性挑战。那么，企业如何才能利用容器化来解决复杂性，而不是最终变得更加复杂呢？本文提供了一些您可以实施以采用 Kubernetes 的最佳实践。

## Keep a Tab on Policies

## 密切关注政策

Define appropriate policies for cluster access controls, service access controls, resource utilization controls and secret access controls. By default, containers run with unbounded compute resources on a Kubernetes cluster. To limit or restrict you must implement appropriate policies.

为集群访问控制、服务访问控制、资源利用控制和秘密访问控制定义适当的策略。默认情况下，容器在 Kubernetes 集群上以无限的计算资源运行。要限制或限制，您必须实施适当的政策。

- Use NetworkPolicy resources labels to select pods and define rules that specify what traffic is allowed to the selected pods.
- Kubernetes scheduler has default limits on the number of volumes that can be attached to a Node. To define the maximum number of volumes that can be attached to a Node for various cloud providers, use Node-specific Volume Limits.
- To enforce constraints on resource usage, use Limit Range option for appropriate resource in the namespace.
- To limit aggregate resource consumption per namespace, use the below Resource Quotas:
   - Compute Resource Quota
   - Storage Resource Quota
   - Object Count Quota
   - Limits the number of resources based on scope defined in Quota Scopes option
   - Requests vs Limits – Each container can specify a request and a limit value for either CPU or memory
   - Quota and cluster capacity – Expressed in absolute units
   - Limit Priority Class consumption by default – For example, restrict usage of certain high-priority pods
- To allow/deny fine-grained permissions, use RBAC (role-based access control) and rules can be defined to allow/deny fine-grained permissions.
- To define and control security aspects of Pods, use Pod Security Policy (available from v1.15) to enable fine-grained authorization of pod creation and updates.
   - Running of privileged containers
   - Usage of host namespaces
   - Usage of host networking and ports
   - Usage of volume types
   - Usage of the host filesystem
   - Restricting escalation to root privileges
   - The user and group IDs of the container
   - AppArmor or seccomp or sysctl profile used by containers
- Use any of the tools such as[Open Policy Agent Gatekeeper](https://www.upnxtblog.com/index.php/2019/12/09/implementing-policies-in-kubernetes/) policy engine to manage, author the policies.



## Manage Resources Wisely

## 明智地管理资源

Use resource utilization (resource quota) guidelines to ensure the containerized applications co-exist without being eliminated due to resource violations at runtime. To enforce constraints on resource usage, use _Limit Range option_ for appropriate resources in the namespace.

使用资源利用率（资源配额）准则来确保容器化应用程序共存，而不会因运行时的资源冲突而被淘汰。要强制限制资源使用，请对命名空间中的适当资源使用 _Limit Range option_。

To limit aggregate resource consumption per namespace, use the below Resource Quotas:

要限制每个命名空间的聚合资源消耗，请使用以下资源配额：

- Compute Resource Quota
- Storage Resource Quota
- Object Count Quota
- Limits the number of resources based on scope defined in Quota Scopes option
- Requests vs Limits – Each container can specify a request and a limit value for either CPU or memory.
- Quota and cluster capacity – Expressed in absolute units
- Limit Priority Class consumption by default – For example, restrict usage of certain high priority pods

- 计算资源配额
- 存储资源配额
- 对象计数配额
- 根据配额范围选项中定义的范围限制资源数量
- 请求与限制——每个容器可以为 CPU 或内存指定一个请求和一个限制值。
- 配额和集群容量 – 以绝对单位表示
- 默认限制优先级消费 - 例如，限制某些高优先级 pod 的使用

## Focus on Comprehensive Observability of the Cluster

## 关注集群的综合可观察性

Currently, the Kubernetes ecosystem provides two add-ons for aggregating and reporting monitoring data from your cluster: **(1) Metrics Server and (2) kube-state-metrics.** 

目前，Kubernetes 生态系统提供了两个附加组件来聚合和报告来自集群的监控数据：**(1) Metrics Server 和 (2) kube-state-metrics。**

[**Metrics**](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/metrics-server.md) **Server is** a cluster add-on that collects resource usage data from each node and provides aggregated metrics through [the Metrics API](https://github.com/kubernetes/metrics). **kube-state-metrics** service provides additional cluster information that Metrics Server does not.

[**指标**](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/metrics-server.md) **服务器是**一个集群附加组件从每个节点收集资源使用数据，并通过 [Metrics API](https://github.com/kubernetes/metrics) 提供聚合指标。 **kube-state-metrics** 服务提供了 Metrics Server 不提供的额外集群信息。

Below are the key metrics and alerts that are required to monitor your Kubernetes cluster.

以下是监控 Kubernetes 集群所需的关键指标和警报。

**What to monitor? **
Monitor the aggregated resources usage across all nodes in your cluster.

监控集群中所有节点的聚合资源使用情况。

- Node status
- Desired pods
- Current pods
- Available pods
- Unavailable pods
- Node status
- Desired vs. current pods
- Available and unavailable pods



**Node resources**

For each of the node monitor :

**节点资源**对于每个节点监视器：

- Memory requests
- Memory limits
- Allocatable memory
- Memory utilization
- CPU requests
- CPU limits
- Allocatable CPU
- CPU utilization
- Disk utilization


If the node’s CPU or memory usage drops below a desired threshold.

如果节点的 CPU 或内存使用率低于所需阈值。

- Memory limits per pod vs. memory utilization per pod
- Memory utilization
- Memory requests per node vs. allocatable memory per node
- Disk utilization
- CPU requests per node vs. allocatable CPU per node
- CPU limits per pod vs. CPU utilization per pod
- CPU utilization



**Missing pod** Health and availability of your pod deployments.

**缺少 Pod** Pod 部署的运行状况和可用性。

- Available pods
- Unavailable pods


If the number of available pods for a deployment falls below the number of pods you specified when you created the deployment .

如果部署的可用 Pod 数量低于您在创建部署时指定的 Pod 数量。



**Pods that are not running**

If a pod isn't running or even scheduled, there could be an issue with either the pod or the cluster, or with your entire Kubernetes deployment.

**未运行的 Pod **

如果 Pod 未运行或什至未安排，则可能存在问题Pod 或集群，或整个 Kubernetes 部署。

Alerts should be based on the status of your pods (“Failed,” ”Pending,” or “Unknown” for the period of time you specify).

警报应基于 Pod 的状态（在您指定的时间段内为“失败”、“待处理”或“未知”）。



**Container restarts** 

Container restarts could happen when you're hitting a memory limit (ex.Out of Memory kills) in your containers.

**容器重新启动**

当您达到内存限制时可能会发生容器重新启动(ex.Out of Memory kills) 在您的容器中。

Also, there could be an issue with either the container itself or its host.

此外，容器本身或其主机可能存在问题。

Kubernetes automatically restarts containers,  but setting up an alert will give you an immediate notification later you can analyze and set the proper limits.

Kubernetes 会自动重启容器，但设置警报后会立即通知您，您可以分析并设置适当的限制。

**Container resource usage**

Monitor container resource usage for containers in case you're hitting resource limits, spikes in resource consumption, Alerts to check if container CPU and memory usage and on limits are based on thresholds.

监控容器的容器资源使用情况，以防您达到资源限制、峰值资源消耗，检查容器 CPU 和内存使用情况以及限制是否基于阈值的警报。

**Storage volumes **Monitor storage to:

**容器资源使用情况** **存储量**监控存储以：

- Ensure your application has enough disk space so pods don’t run out of space
- Volume usage and adjust either the amount of data generated by the application or the size of the volume according to usage

- 确保您的应用程序有足够的磁盘空间，以免 Pod 空间不足
- 卷使用量并根据使用情况调整应用程序生成的数据量或卷的大小

Alerts to check if available bytes, capacity crosses your thresholds.

检查可用字节、容量是否超过阈值的警报。

Identify persistent volumes and apply a different alert threshold or notification for these volumes, which likely hold important application data.

识别持久卷并为这些可能包含重要应用程序数据的卷应用不同的警报阈值或通知。

**Control Plane – Etcd** Monitor etcd for the below parameters:

**控制平面 - Etcd **监控以下参数的 etcd：

- Leader existence and change rate
- Committed, applied, pending, and failed proposals
- gRPC performance

- 领导者的存在和变化率
- 提交、应用、待定和失败的提案
- gRPC 性能

Alerts to check if any pending or failed proposals or reach inappropriate thresholds.

检查是否有任何未决或失败的提案或达到不适当的阈值的警报。



**Control Plane – API Server** Monitor the API server for below parameters:

**控制平面 – API 服务器**监控 API 服务器的以下参数：

- Rate / number of HTTP requests
- Rate/number of apiserver requests

- HTTP请求的速率/数量
- apiserver 请求的速率/数量

Alerts to check if the rate or number of HTTP requests crosses a desired threshold.

检查 HTTP 请求的速率或数量是否超过所需阈值的警报。



**Control Plane – Scheduler** Monitor the scheduler for the below parameters:

**控制平面 – 调度程序 **监控调度程序的以下参数：

- Rate, number, and latency of HTTP requests
- Scheduling latency
- Scheduling attempts by result
- End-to-end scheduling latency (sum of scheduling)

- HTTP 请求的速率、数量和延迟
- 调度延迟
- 按结果安排尝试
- 端到端调度延迟（调度总和）

Alerts to check if the rate or number of HTTP requests crosses a desired threshold.

检查 HTTP 请求的速率或数量是否超过所需阈值的警报。



**Control Plane – Controller Manager**

Monitor the scheduler for the below parameters:

**控制平面 - 控制器管理器**

监控以下参数的调度程序：

- Work queue depth
- Number of retries handled by the work queue 

- 工作队列深度
- 工作队列处理的重试次数

Alerts to check if requests to the work queue exceed a maximum threshold.

用于检查对工作队列的请求是否超过最大阈值的警报。



**Kubernetes Events**

Collecting events from Kubernetes and from the container engine (such as Docker) allows you to see how pod creation, destruction, starting or stopping affects the performance of your infrastructure. Any failure or exception should need to be alerted.

**Kubernetes 事件**

从 Kubernetes 和容器引擎（例如 Docker）收集事件允许您查看 pod 创建、销毁、启动或停止如何影响性能您的基础设施。任何故障或异常都应该得到警告。

Consider integrating with any of the commercial monitoring tools to consume probe-generated metrics and platform-generated metrics to have comprehensive observability of the cluster.

考虑与任何商业监控工具集成以使用探针生成的指标和平台生成的指标，以全面了解集群。

## Container Security Management Must Be Part of Your DevOps Pipeline

## 容器安全管理必须是您的 DevOps 管道的一部分

Continuous security must be included as part of the DevOps pipeline to ensure containers are well-managed. Use any of the below static analysis tools to identify vulnerabilities in application containers while building images for containers:

必须将持续安全性作为 DevOps 管道的一部分包含在内，以确保容器得到良好管理。在为容器构建映像时，使用以下任何一种静态分析工具来识别应用程序容器中的漏洞：

- [Clair](https://github.com/quay/clair)
- [Trivy](https://github.com/aquasecurity/trivy)
- [kube-bench](https://github.com/aquasecurity/kube-bench)
- [Falco](https://github.com/falcosecurity/falco)
- [Notary](https://github.com/theupdateframework/notary)


## Audit and Compliance Your Cluster Routinely

## 定期审核和合规您的集群

Routinely audit the platform for Kubernetes patch levels, secret stores, compliance against the security vulnerabilities, encryption of secret stores, storage volumes, cluster policies, role binding policies, RBAC and user management controls.

定期审核平台的 Kubernetes 补丁级别、秘密存储、安全漏洞合规性、秘密存储加密、存储卷、集群策略、角色绑定策略、RBAC 和用户管理控制。

## Chaos Test Your Cluster

## Chaos 测试你的集群

Proactively chaos tests your platform to ensure the robustness of the cluster. It also helps to test the stability of the containerized applications and the impact of crashing these containers. A wide range of open source and commercial tools can be used, some of which are listed below:

主动混沌测试您的平台，以确保集群的健壮性。它还有助于测试容器化应用程序的稳定性以及这些容器崩溃的影响。可以使用范围广泛的开源和商业工具，其中一些如下所列：

- [Chaosblade](https://github.com/chaosblade-io/chaosblade)
- [Chaos Mesh](https://github.com/pingcap/chaos-mesh)
- [PowerfulSeal](https://github.com/bloomberg/powerfulseal)
- [chaoskube](https://github.com/linki/chaoskube)
- [Chaos Toolkit](https://github.com/chaostoolkit/chaostoolkit)
- [Litmus](https://github.com/litmuschaos/litmus)


## Archive and Back Up Your Cluster

## 归档和备份您的集群

Kubernetes uses etcd as its internal metadata management store to manage the objects across clusters. It is necessary to define a backup strategy for etcd and any other dependent persistent stores used within the Kubernetes clusters.

Kubernetes 使用 etcd 作为其内部元数据管理存储来管理跨集群的对象。有必要为 etcd 和 Kubernetes 集群中使用的任何其他依赖持久存储定义备份策略。

Use [Velero](https://www.upnxtblog.com/index.php/2019/12/16/how-to-back-up-and-restore-your-kubernetes-cluster-resources-and-persistent-volumes/) or any other open source tools to backup **Kubernetes resources** and **application data** so that in cases of **disaster** it can reduce recovery time.

使用 [Velero](https://www.upnxtblog.com/index.php/2019/12/16/how-to-back-up-and-restore-your-kubernetes-cluster-resources-and-persistent-volumes/) 或任何其他开源工具来备份 **Kubernetes 资源** 和 **应用程序数据**，以便在 **灾难** 的情况下可以减少恢复时间。

## Manage Your Deployment Manifests

## 管理您的部署清单

Kubernetes follows declaration-based management; hence, every object or resource or instruction is described only through YAML declarative manifests. It is necessary to leverage SCM tools or create custom utilities to manage these manifests.

Kubernetes 遵循基于声明的管理；因此，每个对象或资源或指令仅通过 YAML 声明性清单进行描述。有必要利用 SCM 工具或创建自定义实用程序来管理这些清单。

## Continuous Deployment of Services

## 服务的持续部署

kubectl style of deployments would not be possible in a large-scale production setup. Instead, you have to use some of the established open source frameworks. [**Helm**](https://helm.sh/), for example, is specifically built for Kubernetes to manage seamless deployments via the CI/CD pipeline.

kubectl 风格的部署在大规模生产设置中是不可能的。相反，您必须使用一些已建立的开源框架。例如，[**Helm**](https://helm.sh/) 是专门为 Kubernetes 构建的，用于通过 CI/CD 管道管理无缝部署。

Helm uses _Charts_ that define the set of Kubernetes resources that together define an application. You can think of Charts as packages of pre-configured Kubernetes resources. Charts help you to define, install and upgrade even the most complex Kubernetes application. These charts can describe a single resource, such as a Redis pod or a full-stack of a web application: HTTP servers, databases and caches.

Helm 使用 _Charts_ 来定义一组共同定义应用程序的 Kubernetes 资源。您可以将 Charts 视为预先配置的 Kubernetes 资源包。 Charts 可帮助您定义、安装和升级最复杂的 Kubernetes 应用程序。这些图表可以描述单个资源，例如 Redis pod 或 Web 应用程序的完整堆栈：HTTP 服务器、数据库和缓存。

In the recent release of Helm, releases will be managed inside of Kubernetes using [Release Objects](https://helm.sh/docs/chart_template_guide/builtin_objects/) and Kubernetes Secrets. All modifications such as installing, upgrading and downgrading releases will end in having a new version of that Kubernetes Secret.

在最近发布的 Helm 中，将使用 [发布对象](https://helm.sh/docs/chart_template_guide/builtin_objects/) 和 Kubernetes Secrets 在 Kubernetes 内部管理发布。所有修改（例如安装、升级和降级版本)都将以该 Kubernetes Secret 的新版本结束。

## Use Service Mesh 

## 使用服务网格

[Service mesh](https://www.upnxtblog.com/index.php/2018/12/17/what-is-service-mesh-why-do-we-need-it-linkered-tutorial/) offers consistent discovery, security, tracing, monitoring and failure handling without the need for a shared asset such as an API gateway. So, if you have service mesh on your cluster, you can achieve all the below items without making changes to your application code:

[服务网格](https://www.upnxtblog.com/index.php/2018/12/17/what-is-service-mesh-why-do-we-need-it-linkered-tutorial/)提供一致的无需共享资产（例如 API 网关)即可发现、安全、跟踪、监控和故障处理。因此，如果您的集群上有服务网格，您可以在不更改应用程序代码的情况下实现以下所有项目：

- Automatic load balancing
- Fine-grained control of traffic behavior with routing rules, retries, failovers, etc.
- Pluggable policy layer
- Configuration API supporting access controls, rate limits and quotas
- Service discovery
- Service monitoring with automatic metrics, logs, and traces for all traffic
- Secure service to service communication



- 自动负载平衡
- 通过路由规则、重试、故障转移等对流量行为进行细粒度控制。
- 可插拔策略层
- 支持访问控制、速率限制和配额的配置 API
- 服务发现
- 使用自动指标、日志和跟踪所有流量的服务监控
- 安全的服务到服务通信

Currently, service mesh is being offered by [Linkerd](https://github.com/linkerd/linkerd2),[Istio](https://github.com/istio/istio) and [Conduit](http://www.conduit.io/) providers.

目前，服务网格由 [Linkerd](https://github.com/linkerd/linkerd2)、[Istio](https://github.com/istio/istio) 和 [Conduit](http://www.conduit.io/) 提供商。

It is necessary to choose an appropriate service mesh that is compatible with the Kubernetes cluster as well as the underlying infrastructure.

有必要选择一个与 Kubernetes 集群以及底层基础设施兼容的合适的服务网格。

## Conclusion

##  结论

This article covers the key best practices that you can implement for Kubernetes adoption. However, operating Kubernetes clusters is not without its challenges.

本文介绍了您可以为采用 Kubernetes 实施的关键最佳实践。然而，运行 Kubernetes 集群并非没有挑战。



### _Related_

###  _有关的_

- [← 2nd Watch Adds Hybrid Anthos Practice to IT Services Portfolio](https://containerjournal.com/topics/container-management/2nd-watch-adds-hybrid-anthos-practice-to-it-services-portfolio/)
- [Microsoft to Bring AKS Service to HCI Platforms →](https://containerjournal.com/topics/container-management/microsoft-to-bring-aks-service-to-hci-platforms/) 

- [← 2nd Watch 将混合 Anthos 实践添加到 IT 服务组合](https://containerjournal.com/topics/container-management/2nd-watch-adds-hybrid-anthos-practice-to-it-services-portfolio/)
- [微软将 AKS 服务引入 HCI 平台 →](https://containerjournal.com/topics/container-management/microsoft-to-bring-aks-service-to-hci-platforms/)

