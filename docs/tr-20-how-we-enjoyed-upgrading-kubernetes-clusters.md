# How we enjoyed upgrading a bunch of Kubernetes clusters from v1.16 to v1.19

# 我们如何享受将一堆 Kubernetes 集群从 v1.16 升级到 v1.19

From: https://blog.flant.com/how-we-enjoyed-upgrading-kubernetes-clusters-from-v1-16-to-v1-19/

At the beginning of December 2020, we at Flant maintained about  150 clusters running Kubernetes 1.16. All these clusters have varying  degrees of load. Some of them are the high-load production clusters,  while others are intended for the development and testing of new  features. These clusters run on top of different infrastructure  solutions, starting with cloud providers such as AWS, Azure, GCP,  various OpenStack/vSphere installations, and ending with bare-metal  servers.

2020 年 12 月初，我们在 Flant 维护了大约 150 个运行 Kubernetes 1.16 的集群。所有这些集群都有不同程度的负载。其中一些是高负载生产集群，而另一些则用于新功能的开发和测试。这些集群运行在不同的基础架构解决方案之上，从 AWS、Azure、GCP、各种 OpenStack/vSphere 安装等云提供商开始，到裸机服务器结束。

The clusters are managed by [Deckhouse](https://deckhouse.io/) — the tool developed at Flant that will be released as an Open Source  project this May. It is a single instrument for creating clusters and  an interface for managing all cluster components on all supported infrastructure types. To do this, Deckhouse consists of various  subsystems. For example, there is the *candi* (**C**luster **AND** **I**nfrastructure) subsystem. It is of particular interest to us in this article since it  manages the Kubernetes control-plane, configures nodes, and creates a  viable, up-to-date cluster.

集群由 [Deckhouse](https://deckhouse.io/) 管理——这是 Flant 开发的工具，将于今年 5 月作为开源项目发布*。它是用于创建集群的单一工具和用于管理所有支持的基础架构类型上的所有集群组件的界面。为此，Deckhouse 由各种子系统组成。例如，有 *candi* (**C**luster **AND** **I**nfrastructure) 子系统。我们在本文中特别感兴趣，因为它管理 Kubernetes 控制平面、配置节点并创建可行的最新集群。

** Currently, Deckhouse is available via our* [*Managed Kubernetes*](https://flant.com/services/managed-kubernetes-as-a-service) *service only. Its first public Open Source version will arrive next month. You can follow the* [*project’s Twitter*](https://twitter.com/deckhouseio) *to stay posted: the official announcement will be there.*

** 目前，Deckhouse 可通过我们的* [*Managed Kubernetes*](https://flant.com/services/managed-kubernetes-as-a-service) *服务获得。它的第一个公开开源版本将于下个月推出。你可以关注* [*项目的推特*](https://twitter.com/deckhouseio) *保持发布：官方公告将在那里。*

So, why did we get stuck with v1.16 if Kubernetes 1.17, 1.18, and  even a patch for version 1.19 were available for quite some time? The  thing is that the previous cluster upgrade (from 1.15 to 1.16) wasn’t so smooth. Or, more precisely, it went way too hard for the following  reasons:

那么，如果 Kubernetes 1.17、1.18 甚至 1.19 版本的补丁可用很长时间，为什么我们会卡在 v1.16 上？问题是之前的集群升级（从 1.15 到 1.16）并不是那么顺利。或者，更准确地说，由于以下原因，它变得太难了：

- [Kubernetes 1.16](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md) finally [got rid of](https://kubernetes.io/blog/2019/07 /18/api-deprecations-in-1-16/) some deprecated API versions. The most painful was the abandonment of the old versions of the API controllers: DaemonSet (`extensions/v1beta1` and `apps/v1beta2`), Deployment (`extensions/v1beta1`, `apps/v1beta1` and `apps/v1beta2`) and StatefulSet (`apps/v1beta1` and `apps/v1beta2`). Before switching to 1.16, we had to make sure that all API versions  were updated in the Helm charts, or the deployment of a Helm release  would have failed after the upgrade. Since applications are deployed  into the clusters and both our DevOps teams and the customers’  developers are engaged in writing Helm charts, we first had to notify  all parties involved about the problem. For this, we have implemented a  module in Deckhouse that checks all the latest installed Helm releases  in the cluster (to figure out if the old API versions are used) and  provides metrics with that info. Then Prometheus and Alertmanager join  in with their alerts.
 - To switch from 1.15 to 1.16, we had to restart the  containers on the node. So we had to drain it first, which, in the case  of many stateful applications, forced us to perform all the  manipulations at the agreed time and required the special attention of  engineers.

- [Kubernetes 1.16](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md) 终于[摆脱](https://kubernetes.io/blog/2019/07)一些已弃用的 API 版本。最痛苦的是放弃旧版本的 API 控制器：DaemonSet（`extensions/v1beta1` 和 `apps/v1beta2`）、Deployment（`extensions/v1beta1`、`apps/v1beta1` 和 `apps/v1beta2`）和 StatefulSet（`apps/v1beta1` 和 `apps/v1beta2`）。在切换到 1.16 之前，我们必须确保 Helm 图表中的所有 API 版本都已更新，否则升级后 Helm 版本的部署将失败。由于应用程序已部署到集群中，并且我们的 DevOps 团队和客户的开发人员都在参与编写 Helm 图表，因此我们首先必须将问题通知所有相关方。为此，我们在 Deckhouse 中实现了一个模块，用于检查集群中所有最新安装的 Helm 版本（以确定是否使用了旧的 API 版本）并提供包含该信息的指标。然后 Prometheus 和 Alertmanager 加入他们的警报。
- 要从 1.15 切换到 1.16，我们必须重新启动节点上的容器。所以我们不得不先把它排空，在许多有状态应用程序的情况下，这迫使我们在约定的时间执行所有操作，并需要工程师的特别关注。

These two factors slowed down the update process considerably. On the one hand, we had to persuade customers to remake the charts and release updates, and on the other hand, we were supposed to update the clusters carefully and in a half-manual mode. We have done much to convince all  cluster users that the need for an upgrade is real, valid and that the  approach “don’t touch it as long as it works” can play tricks  eventually. In this article, I will try to convince you of that as well.

这两个因素大大减慢了更新过程。一方面，我们不得不说服客户重新制作图表并发布更新，另一方面，我们应该以半手动模式仔细地更新集群。我们已经做了很多工作来说服所有集群用户升级的需求是真实的、有效的，并且“只要有效就不要碰它”的方法最终会耍花招。在这篇文章中，我也会试着说服你。

## Why even bother with a Kubernetes upgrade? 
## 为什么还要费心升级 Kubernetes？
Software aging is the most obvious reason for a Kubernetes upgrade. The thing is that[ only the three latest minor versions](https://kubernetes.io/docs/setup/release/version-skew-policy/) are supported by Kubernetes developers. Thus, with version 1.19  released, our current version 1.16 left the list of supported versions. On the other hand, with this ([v1.19](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md)) release, Kubernetes developers have taken into account the grim picture painted by the statistic and increased the support period to a year. Those stats indicated that the majority of existing K8s installations  were obsolete. (The[ survey](https://github.com/youngnick/enhancements/blob/one-year-support-window/keps/sig-release/20200122-kubernetes-yearly-support-period.md#motivation) conducted in early 2019 showed that only 50-60% of K8s clusters are running the supported Kubernetes version.)

软件老化是 Kubernetes 升级的最明显原因。问题是[只有三个最新的小版本](https://kubernetes.io/docs/setup/release/version-skew-policy/) 被 Kubernetes 开发者支持。因此，随着 1.19 版的发布，我们当前的 1.16 版离开了支持的版本列表。另一方面，随着这个 ([v1.19](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md)) 的发布，Kubernetes 开发者已经考虑到了严峻的形势通过统计绘制并将支持期增加到一年。这些统计数据表明，大多数现有的 K8s 安装已经过时。 （[调查](https://github.com/youngnick/enhancements/blob/one-year-support-window/keps/sig-release/20200122-kubernetes-yearly-support-period.md#motivation)在2019 年初表明，只有 50-60% 的 K8s 集群在运行受支持的 Kubernetes 版本。）

![img](https://blog.flant.com/wp-content/uploads/sites/2/2021/04/kubernetes-versions-in-production-2019-survey.png)At the time of the survey, the current K8s version was 1.13. Thus, all  users of Kubernetes 1.9 and 1.10 were running releases that were no  longer supported.

当前的 K8s 版本是 1.13。因此，Kubernetes 1.9 和 1.10 的所有用户都在运行不再受支持的版本。

***NB**. The issue of Kubernetes upgrades is widely  discussed in the community of Ops engineers. For example, Platform9  surveys in 2019 ([#1](https://platform9.com/blog/six-kubernetes-takeaways-for-it-ops-teams-from-the-2019-gartner-infrastructure-operations-cloud -strategies-conference/),[ # 2](https://platform9.com/blog/six-enterprise-kubernetes-takeaways-from-kubecon-2019-san-diego/)) showed that the upgrade was one of the top three challenges when maintaining Kubernetes. Actually, the Internet is full of[ failure](https://deploy.live/blog/the-shipwreck-of-gke-cluster-upgrade/)[ stories](https://www.jetstack.io/blog/ gke-webhook-outage/),[ webinars](https://www.youtube.com/watch?v=i24KuG5bGkg), etc., on the topic.*

***注意**。 Kubernetes 升级的问题在运维工程师社区中被广泛讨论。例如，Platform9 在 2019 年的调查（[#1](https://platform9.com/blog/six-kubernetes-takeaways-for-it-ops-teams-from-the-2019-gartner-infrastructure-operations-cloud -strategies-conference/),[#2](https://platform9.com/blog/six-enterprise-kubernetes-takeaways-from-kubecon-2019-san-diego/)) 表明升级是其中之一维护 Kubernetes 的三大挑战。其实网上到处都是[失败](https://deploy.live/blog/the-shipwreck-of-gke-cluster-upgrade/)[故事](https://www.jetstack.io/blog/ gke-webhook-outage/),[ 网络研讨会](https://www.youtube.com/watch?v=i24KuG5bGkg) 等，关于这个主题。*

But let’s get back to version 1.16. It had several issues that we  were forced to fix via various workarounds. Probably, most of our  readers did not encounter these issues. Still, we maintain a large  number of clusters (with thousands of nodes), so we regularly had to  deal with the consequences of those rare errors. By December, we  invented many tricky components both in Kubernetes and system units,  e.g.:

但是让我们回到 1.16 版。它有几个问题，我们被迫通过各种变通办法解决。可能我们的大多数读者都没有遇到过这些问题。尽管如此，我们仍维护着大量集群（具有数千个节点），因此我们必须定期处理这些罕见错误的后果。到 12 月，我们在 Kubernetes 和系统单元中发明了许多棘手的组件，例如：

- Quite regularly, alerts started to fire up, saying that some random cluster nodes are `NotReady`. It was caused by the[ issue](https://github.com/kubernetes/kubernetes/issues/87615) that was finally fixed in Kubernetes 1.19. However, the fix does not allow any backporting:![img](https://blog.flant.com/wp-content/uploads/sites/2/2021/04/k8s-issue-nodes-notready.png)
    To sleep blissfully at night, we deployed a systemd unit to all cluster nodes and named it … `kubelet-face-slapper`. It monitored the kubelet logs and restarted the kubelet in case of an `use of closed network connection` error. If you look at the history of the issue, you can see that people from all over the world had to apply similar workarounds. 
- 很规律地，警报开始出现，说一些随机的集群节点是“NotReady”。这是由最终在 Kubernetes 1.19 中修复的[问题](https://github.com/kubernetes/kubernetes/issues/87615) 引起的。为了晚上睡个好觉，我们在所有集群节点上部署了一个 systemd 单元，并将其命名为……`kubelet-face-slapper`。它监视 kubelet 日志并在出现“使用关闭的网络连接”错误时重新启动 kubelet。如果您查看该问题的历史，您会发现来自世界各地的人们不得不应用类似的解决方法。
- Occasionally, we have been noticing some strange  problems with the kube-scheduler. In our installations, metrics are  collected solely over HTTPS using a Prometheus client certificate. However, Prometheus has randomly stopped receiving scheduler metrics  because of the kube-scheduler that could not correctly process data of  the client certificate. We did not find any related issues (probably,  collecting metrics via HTTPS is not such a common practice). Still, the  code base has changed significantly between versions 1.16 and 1.19 (a  lot of refactoring and bug fixes took place), which is why we were sure  that the upgrade would solve this problem. Meanwhile, we ran a special  component on master nodes in each cluster as a temporary solution. It  simulated Prometheus scraping of metrics and, in the case of an error,  restarted the kube-scheduler. We named this component in the same  fashion — `kube-scheduler-face-slapper`.
 - Sometimes, something even more horrible happened. You could get many problems in the clusters because the kube-proxy has  crashed when accessing the kube-apiserver. It was caused by the lack of  health checks for HTTP/2 connections for all clients used in Kubernetes. The frozen kube-proxy has induced network problems (for obvious  reasons) that could end up in downtime. The[ fix](https://github.com/kubernetes/kubernetes/pull/95981) was released for version 1.20 and was backported to K8s 1.19 only. (By  the way, the same fix solved the problems with the kubelet freezes.)  Also, periodically, the kubectl could freeze when performing some  lengthy operations, so you had to always keep in mind that timeouts must be set.
 - We used Docker as a Container Runtime in clusters, which created problems with the Pods being stuck in a `Terminating` state regularly. These problems were caused by the widely-known [bug](https://github.com/kubernetes/kubernetes/issues/84214).

- 有时，我们会注意到 kube-scheduler 存在一些奇怪的问题。在我们的安装中，使用 Prometheus 客户端证书仅通过 HTTPS 收集指标。但是，由于 kube-scheduler 无法正确处理客户端证书的数据，Prometheus 随机停止接收调度程序指标。我们没有发现任何相关问题（可能，通过 HTTPS 收集指标并不是一种常见的做法）。尽管如此，1.16 和 1.19 版本之间的代码库已经发生了显着变化（进行了大量重构和错误修复），这就是为什么我们确信升级会解决这个问题。同时，我们在每个集群的主节点上运行了一个特殊的组件作为临时解决方案。它模拟 Prometheus 抓取指标，并在出现错误的情况下重新启动 kube-scheduler。我们以同样的方式命名这个组件——`kube-scheduler-face-slapper`。
- 有时，更可怕的事情发生了。您可能会在集群中遇到许多问题，因为访问 kube-apiserver 时 kube-proxy 已崩溃。这是由于 Kubernetes 中使用的所有客户端都缺乏对 HTTP/2 连接的健康检查造成的。冻结的 kube-proxy 引发了可能导致停机的网络问题（出于显而易见的原因）。 [修复](https://github.com/kubernetes/kubernetes/pull/95981) 已针对 1.20 版发布，仅向后移植到 K8s 1.19。 （顺便说一下，同样的修复解决了 kubelet 冻结的问题。）此外，定期执行一些冗长的操作时，kubectl 可能会冻结，因此您必须始终牢记必须设置超时。
- 我们在集群中使用 Docker 作为容器运行时，这会导致 Pod 经常陷入“终止”状态的问题。这些问题是由众所周知的 [bug](https://github.com/kubernetes/kubernetes/issues/84214) 引起的。

Also, there were other annoying problems — such as[ mount errors](https://github.com/kubernetes/kubernetes/pull/89629) when restarting containers if `subpath`s are used. We did not start inventing another *whatever-face-slapper* and decided that it is finally the time to upgrade to v1.19. Especially since, by this time, almost all our clusters were ready for the  upgrade.

此外，还有其他烦人的问题——例如，如果使用“子路径”，则在重新启动容器时会出现[挂载错误](https://github.com/kubernetes/kubernetes/pull/89629)。我们没有开始发明另一个 *whatever-face-slapper* 并决定现在是升级到 v1.19 的时候了。尤其是因为此时，我们几乎所有的集群都已准备好进行升级。

## How does the Kubernetes upgrade work?

## Kubernetes 升级是如何工作的？

Earlier, we mentioned Deckhouse and the *candi* subsystem  responsible for upgrading the control-plane and cluster nodes (among  other things). Basically, it has a slightly modified `kubeadm` inside. Thus, structurally, the upgrade process is similar to the one described in the[ Kubernetes documentation](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/) on upgrading clusters managed by `kubeadm `.

早些时候，我们提到了 Deckhouse 和负责升级控制平面和集群节点（除其他外）的 *candi* 子系统。基本上，它内部有一个稍微修改过的 `kubeadm`。因此，在结构上，升级过程类似于[Kubernetes 文档](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/) 中描述的升级由`kubeadm 管理的集群的过程`.

Steps for upgrading from 1.16 to 1.19 are as follows:

从1.16升级到1.19的步骤如下：

- Updating the control plane to version 1.17;
 - Updating the kubelet on nodes to version 1.17;
 - Updating the control plane to version 1.18;
 - Updating the kubelet on nodes to version 1.18;
 - Updating the control plane to version 1.19;
 - Updating the kubelet on nodes to version 1.19.

- 将控制平面更新至 1.17 版；
- 将节点上的 kubelet 更新到 1.17 版；
- 将控制平面更新至 1.18 版；
- 将节点上的 kubelet 更新到 1.18 版；
- 将控制平面更新至 1.19 版；
- 将节点上的 kubelet 更新到 1.19 版。

Deckhouse performs these steps automatically. For this, each cluster  has a Secret with the cluster configuration YAML file of the following  format:

Deckhouse 会自动执行这些步骤。为此，每个集群都有一个 Secret，其中包含以下格式的集群配置 YAML 文件：

```
 apiVersion: deckhouse.io/v1alpha1
 cloud:
   prefix: k-aksenov
   provider: OpenStack
 clusterDomain: cluster.local
 clusterType: Cloud
 kind: ClusterConfiguration
 kubernetesVersion: "1.16"
 podSubnetCIDR: 10.111.0.0/16
 podSubnetNodeCIDRPrefix: "24"
 serviceSubnetCIDR: 10.222.0.0/16
```


To fire up an upgrade, you just have to change the `kubernetesVersion` to the desired value (you can skip all the interim versions and get right to v1.19). There are two modules in the *candi* subsystem that are responsible for managing the control-plane and nodes.

要启动升级，您只需将 `kubernetesVersion` 更改为所需的值（您可以跳过所有临时版本并直接使用 v1.19）。 *candi* 子系统中有两个模块负责管理控制平面和节点。

The **control-plane-manager** automatically monitors this YAML file for changes.

- The current Kubernetes version is calculated based on the version of the control-plane and the cluster nodes. For example, if all nodes are running the kubelet 1.16 and all the control-plane  components have the same version (1.16), you can start an upgrade to  version 1.17. This process continues until the current version matches  the target one.
 - Also, the control-plane-manager makes sure that  control-plane components are upgraded sequentially on each master. For  this, through a dedicated manager, we have implemented an algorithm for  requesting and granting permission to upgrade.

**control-plane-manager** 自动监控这个 YAML 文件的变化。

- 当前Kubernetes版本是根据控制平面和集群节点的版本计算的。例如，如果所有节点都运行 kubelet 1.16，并且所有控制平面组件的版本都相同（1.16），则可以开始升级到 1.17 版本。这个过程一直持续到当前版本与目标版本匹配。
- 此外，控制平面管理器确保控制平面组件在每个主机上按顺序升级。为此，我们通过专门的管理器实现了一种用于请求和授予升级权限的算法。

**Node-manager** manages nodes and updates the kubelet: 

- Each node in the cluster belongs to some `NodeGroup`. As soon as the node-manager determines that the control-plane version  has been successfully updated on all nodes, it proceeds to update the  kubelet version. If the node upgrade does not involve downtime, it is  considered safe and is performed automatically. In this case, upgrading  the kubelet no longer requires the restart of the containers and,  therefore, recognized as safe.
 - Node-manager also has a mechanism for automatically  granting permissions to upgrade a node. It guarantees that only one node within a NodeGroup can be upgraded at a time. At the same time,  NodeGroup nodes are upgraded only if the required number of nodes is  equal to the current number of nodes in the `Ready` state, meaning no nodes are being ordered. *(Obviously, this only applies to cloud clusters where there is an automatic ordering of new VMs.)*

**Node-manager** 管理节点并更新 kubelet：

- 集群中的每个节点都属于某个 `NodeGroup`。一旦节点管理器确定控制平面版本已在所有节点上成功更新，它就会继续更新 kubelet 版本。如果节点升级不涉及停机，则认为是安全的并自动执行。在这种情况下，升级 kubelet 不再需要重启容器，因此被认为是安全的。
- 节点管理器还具有自动授予升级节点权限的机制。它保证一次只能升级一个 NodeGroup 中的一个节点。同时，NodeGroup 节点仅在所需节点数等于当前处于“就绪”状态的节点数时才会升级，这意味着没有节点被排序。 *（显然，这仅适用于新虚拟机自动排序的云集群。）*

## Our experience with upgrading Kubernetes to 1.19

## 我们将 Kubernetes 升级到 1.19 的经验

There were several reasons for upgrading directly to 1.19 and bypassing versions 1.17 and 1.18.

直接升级到 1.19 并绕过 1.17 和 1.18 版本有几个原因。

The main reason was that we didn’t want to **delay the upgrade process**. Each upgrade cycle involves coordination with the cluster users and  requires considerable effort from engineers that control the upgrade  process. This way, the risk of lagging behind the upstream persisted. But we wanted to upgrade all our clusters to Kubernetes 1.19 by February 2021. So, there was a strong desire to skip ahead to the latest  Kubernetes version in our clusters with one well-coordinated effort —  especially given that version 1.20 was around the corner ( it was  released on December 8, 2020).

主要原因是我们不想**延迟升级过程**。每个升级周期都涉及与集群用户的协调，并且需要控制升级过程的工程师付出相当大的努力。这样一来，上游落后的风险依然存在。但是我们希望在 2021 年 2 月之前将所有集群升级到 Kubernetes 1.19。因此，我们强烈希望通过协调一致的努力跳到集群中的最新 Kubernetes 版本——尤其是考虑到 1.20 版本即将推出（它于 2020 年 12 月 8 日发布）。

As mentioned above, we have clusters that run on instances  provisioned by various cloud providers. Thus, we use various components  specific to each cloud provider. [Container Storage Interface](https://kubernetes-csi.github.io/docs/) manages disks in our clusters while [Cloud Controller Manager](https://kubernetes.io/docs/concepts/architecture/cloud-controller /) interacts with APIs of cloud providers. Testing the operability of  these components for each Kubernetes version is very resource-intensive. And in our “skip ahead” case, we just **save time and efforts that otherwise would be spent on interim versions**.

如上所述，我们有在各种云提供商提供的实例上运行的集群。因此，我们使用特定于每个云提供商的各种组件。 [容器存储接口](https://kubernetes-csi.github.io/docs/) 管理我们集群中的磁盘，而[云控制器管理器](https://kubernetes.io/docs/concepts/architecture/cloud-controller /) 与云提供商的 API 交互。为每个 Kubernetes 版本测试这些组件的可操作性非常耗费资源。在我们的“跳过”案例中，我们只是**节省了原本会花在临时版本上的时间和精力**。

### The process

### 过程

So, we conducted full compatibility testing of components with  version 1.19 and decided to skip all interim versions. Since the  standard upgrade goes sequentially, we temporarily disabled the  components described above to avoid possible conflicts with  control-plane versions 1.17 and 1.18 in the cloud clusters.

因此，我们对 1.19 版本的组件进行了全面的兼容性测试，并决定跳过所有过渡版本。由于标准升级是按顺序进行的，我们暂时禁用了上述组件，以避免与云集群中的控制平面版本 1.17 和 1.18 可能发生冲突。

The upgrade duration depended on the number of worker / control-plane nodes and took between 20 to 40 minutes. During this period, ordering  nodes, deleting them, and any operations with disks were not available  in the cloud clusters. At the same time, nodes running in the cluster  and the mounted disks continued to work properly. It was the only  obvious disadvantage of the upgrade, and we had to accept it. Because of it, we decided to upgrade most clusters at night, when the load is low.

升级持续时间取决于工作节点/控制平面节点的数量，需要 20 到 40 分钟。在此期间，云集群中无法对节点进行排序、删除和任何磁盘操作。同时，集群中运行的节点和挂载的磁盘继续正常工作。这是升级唯一明显的缺点，我们不得不接受。因此，我们决定在晚上负载较低的时候升级大多数集群。

We first ran the upgrade on internal dev clusters several times and  then proceeded to upgrade customers’ clusters. Once again, we did it  slowly and carefully, starting with dev clusters since they can tolerate a little downtime.

我们首先在内部开发集群上运行了几次升级，然后继续升级客户的集群。再一次，我们从开发集群开始，缓慢而谨慎地进行，因为它们可以容忍一点停机时间。

### The first upgrades and the first pain

###第一次升级和第一次痛苦

The upgrade of low-load clusters went smoothly. However, we encountered a problem attempting to upgrade the high-load ones:[ Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) was still active and tried to request additional nodes. Deckhouse uses the[ Machine Controller Manager](https://github.com/gardener/machine-controller-manager) to order nodes. Since Flant engineers[ contribute actively](https://twitter.com/flant_com/status/1356948464748404736) to this project, we were sure that it is fully compatible with all  Kubernetes versions and did not disable it. But with Cloud Controller  Manager disabled, the nodes could not transition to the `Ready` state. Why?

低负载集群升级顺利。但是，我们在尝试升级高负载时遇到了问题：[ Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) 仍然处于活动状态并尝试请求其他节点。 Deckhouse 使用[机器控制器管理器](https://github.com/gardener/machine-controller-manager) 对节点进行排序。由于 Flant 工程师[积极贡献](https://twitter.com/flant_com/status/1356948464748404736) 到这个项目，我们确信它与所有 Kubernetes 版本完全兼容并且没有禁用它。但是在禁用 Cloud Controller Manager 的情况下，节点无法转换到“就绪”状态。为什么？

As we mentioned earlier, the node-manager has a built-in protection  mechanism: nodes are upgraded one at a time and only if all NodeGroup  nodes are `Ready`. It turns out we made things more difficult for ourselves by blocking node upgrades when the cluster is in the  process of scaling. 

正如我们之前提到的，节点管理器有一个内置的保护机制：一次升级一个节点，并且只有当所有 NodeGroup 节点都“就绪”时。事实证明，当集群处于扩展过程中时，我们通过阻止节点升级使事情变得更加困难。

Thus, we had to manually “push” the blocked clusters through the  upgrade process by changing NodeGroup nodes’ number and deleting the  newly ordered nodes that could not transition to the `Ready`  state. And, of course, we quickly found a universal solution — disabling the Machine Controller Manager and Cluster Autoscaler (just like other  components that were already disabled in cloud clusters).

因此，我们必须通过更改 NodeGroup 节点的数量并删除无法转换为“就绪”状态的新排序节点来手动“推送”阻塞的集群，以完成升级过程。当然，我们很快找到了一个通用的解决方案——禁用机器控制器管理器和集群自动缩放器（就像其他已经在云集群中禁用的组件一样）。

### Full-scale upgrade

###全面升级

By the end of December 2020, about 50 Kubernetes clusters were  upgraded to 1.19. Up to this point, we dared to upgrade only a few  clusters (mostly 2-3, but no more than 5) simultaneously. With the  growing confidence in the stability and correctness of the process, we  decided to run the full-scale upgrade of around 100 remaining clusters. We wanted to upgrade almost all of them to Kubernetes 1.19 by the end of January. The clusters were divided into two groups, with each  containing about 50 clusters.

到 2020 年 12 月底，约有 50 个 Kubernetes 集群升级到 1.19。到目前为止，我们只敢同时升级几个集群（主要是 2-3 个，但不超过 5 个）。随着对流程稳定性和正确性的信心不断增强，我们决定对剩余的大约 100 个集群进行全面升级。我们希望在 1 月底之前将几乎所有这些都升级到 Kubernetes 1.19。这些集群被分为两组，每组包含大约 50 个集群。

In the process of upgrading the first group, we encountered only one problem. One of the cluster nodes was persistently `NotReady` while the kubelet was getting the following errors:

在升级第一组的过程中，我们只遇到了一个问题。当 kubelet 收到以下错误时，其中一个集群节点一直处于“NotReady”状态：

```
 Failed to initialize CSINodeInfo: error updating CSINode annotation
```


Unfortunately, we didn’t have a chance to debug this issue: the  cluster was running on a bare-metal machine that was fully loaded. So  the failure of a single node started to impact the performance of the  application. We found a quick[ fix](https://github.com/kubernetes/kubernetes/issues/86094#issuecomment-564113477) that might help you as well if you will ever encounter a similar issue. Again, we would like to emphasize that this problem occurred only  once(!) on our 150 clusters during the upgrade. In other words, it is  pretty rare.

不幸的是，我们没有机会调试这个问题：集群在一台满载的裸机上运行。因此单个节点的故障开始影响应用程序的性能。我们找到了一个快速[修复](https://github.com/kubernetes/kubernetes/issues/86094#issuecomment-564113477)，如果您遇到类似的问题，它也可能对您有所帮助。再次强调，此问题在升级期间仅在我们的 150 个集群上发生一次（！）。换句话说，这是非常罕见的。

To meet the deadlines we set for ourselves, we scheduled the upgrade  of the second group of clusters for midnight of January 28. This group  included the production clusters with the highest load: their downtime  usually results in writing post-mortems, violating SLAs, and penalties. [Our CTO](https://www.linkedin.com/in/distol/) has personally supervised the upgrade process. Luckily, everything went smoothly: no unexpected problems, no manual intervention was required.

为了满足我们为自己设定的最后期限，我们计划在 1 月 28 日午夜升级第二组集群。该组包括负载最高的生产集群：它们的停机时间通常会导致编写事后分析、违反 SLA 和处罚。 [我们的 CTO](https://www.linkedin.com/in/distol/) 亲自监督了升级过程。幸运的是，一切都很顺利：没有意外问题，也不需要人工干预。

### The last pain

###最后的痛苦

However, as it turned out later, there was a problem related to the  upgrade. Its consequences manifested only after a couple of days, though (and, as is usually the case, in one of the most loaded and important  apps).

但是，后来发现，升级出现了问题。然而，它的后果仅在几天后显现（并且，通常情况下，在加载最多和最重要的应用程序之一中）。

During the upgrade, the kube-apiserver is being restarted several  times. Historically, the part of our clusters running under Deckhouse  uses flannel. The tricky part is that all this time, flannel had the  problem mentioned above: the Go client failure due to the lack of health checks for HTTP/2 connections when accessing the Kubernetes API. As a  result, errors like the one shown below appeared in the flannel logs:

升级过程中，kube-apiserver 会多次重启。从历史上看，我们在 Deckhouse 下运行的集群部分使用法兰绒。棘手的部分是，一直以来，flannel 都存在上述问题：由于在访问 Kubernetes API 时缺乏对 HTTP/2 连接的健康检查，导致 Go 客户端失败。结果，法兰绒日志中出现了如下所示的错误：

```
 E0202 21:52:35.791600       1 reflector.go:201] github.com/coreos/flannel/subnet/kube/kube.go:310: Failed to list *v1.Node: Get https://192.168.0.1:443/ api/v1/nodes?resourceVersion=0: dial tcp 192.168.0.1:443: connect: connection refused
```


They caused CNI to operate incorrectly on these nodes, which led to  5XX errors when accessing the services. Restarting the flannel fixed the problem. However, to solve it once and for all, you need to update the  flannel.

它们导致 CNI 在这些节点上运行不正确，从而导致访问服务时出现 5XX 错误。重新启动法兰绒解决了问题。但是，要一劳永逸地解决它，您需要更新法兰绒。

Fortunately, the[ pull request](https://github.com/flannel-io/flannel/pull/1384) opened on December 15, 2020, bumped the version of client-go. At the time of writing, the fix is available in[ v0. 13. 1-rc2](https://github.com/flannel-io/flannel/releases/tag/v0.13.1-rc2), and we plan to update the flannel version in all clusters running under Deckhouse.

幸运的是，2020 年 12 月 15 日开通的[pull request](https://github.com/flannel-io/flannel/pull/1384)，升级了 client-go 的版本。在撰写本文时，修复程序可在 [ v0. 13. 1-rc2](https://github.com/flannel-io/flannel/releases/tag/v0.13.1-rc2)，我们计划在Deckhouse下运行的所有集群中更新flannel版本。

## So, where is the gain?

## 那么，收益在哪里？

Now all our 150+ clusters are happily running Kubernetes 1.19 (not a  single v1.16 cluster is left). The upgrade went quite smoothly. It was a pleasure to watch how automation solves all the problems and how the `Version` value in the `kubectl get nodes` output changes almost simultaneously for thousands of nodes within an  hour or so. That was fascinating — we really enjoyed the way the  Kubernetes upgrade went.

现在我们所有 150 多个集群都在愉快地运行 Kubernetes 1.19（没有一个 v1.16 集群留下）。升级还算顺利。很高兴看到自动化如何解决所有问题，以及“kubectl get nodes”输出中的“Version”值如何在一小时左右内几乎同时更改数千个节点。这很有趣——我们真的很喜欢 Kubernetes 升级的方式。

The support for Kubernetes 1.20 in Deckhouse has been ready for a  while, so we are going to upgrade our clusters to the latest K8s version ASAP (especially now, when v1.21 has already landed…) using the method  described in this article. And we will definitely share any nuances or  problems encountered during the upcoming upgrade, so stay tuned!

Deckhouse 对 Kubernetes 1.20 的支持已经准备好一段时间了，所以我们将使用本文描述的方法尽快将我们的集群升级到最新的 K8s 版本（尤其是现在，当 v1.21 已经登陆时......）。我们一定会分享在即将到来的升级过程中遇到的任何细微差别或问题，敬请期待！

## Related posts:

## 相关文章：

- [Go? Bash! Meet the shell-operator ](https://blog.flant.com/go-bash-meet-the-shell-
- [shell-operator v1.0.0: the long-awaited release of our tool to create Kubernetes operators ](https://blog.flant.com/shell-operator-v1-release-for-kubernetes-operators/)
- [Making the most out of Helm templates ](https://blog.flant.com/advanced-helm-templating/) 

