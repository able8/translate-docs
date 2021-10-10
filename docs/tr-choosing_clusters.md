# Choosing a Local Dev Cluster

# 选择本地开发集群

How do you run Kubernetes locally?

如何在本地运行 Kubernetes？

There are lots of Kubernetes dev solutions out there. The choices can be overwhelming. We’re here to help you figure out the right one for you.

有很多 Kubernetes 开发解决方案。选择可能是压倒性的。我们在这里帮助您找到适合您的产品。

Beginner Level:

初级水平：

- [Kind](https://docs.tilt.dev/choosing_clusters.html#kind)
- [Docker for Desktop](https://docs.tilt.dev/choosing_clusters.html#docker-for-desktop)
- [Microk8s](https://docs.tilt.dev/choosing_clusters.html#microk8s)


Intermediate Level:

中级水平：

- [Minikube](https://docs.tilt.dev/choosing_clusters.html#minikube)
- [k3d](https://docs.tilt.dev/choosing_clusters.html#k3d)


Advanced Level:

先进的水平：

- [Amazon Elastic Kubernetes Service](https://docs.tilt.dev/choosing_clusters.html#remote)
- [Azure Kubernetes Service](https://docs.tilt.dev/choosing_clusters.html#remote)
- [Google Kubernetes Engine](https://docs.tilt.dev/choosing_clusters.html#remote)
- [Custom Clusters](https://docs.tilt.dev/choosing_clusters.html#custom-clusters)


------

## [Kind](https://docs.tilt.dev/choosing_clusters.html#kind)

[Kind](https://kind.sigs.k8s.io/) runs Kubernetes inside a Docker container.

[Kind](https://kind.sigs.k8s.io/) 在 Docker 容器内运行 Kubernetes。

The Kubernetes team uses Kind to test Kubernetes itself. But its fast startup time also makes it a good solution for local dev. Use `ctlptl` to set up Kind with a registry:

Kubernetes 团队使用 Kind 来测试 Kubernetes 本身。但它的快速启动时间也使它成为本地开发人员的一个很好的解决方案。使用 `ctlptl` 使用注册表设置 Kind：

[**Kind Setup**](https://github.com/tilt-dev/ctlptl#kind-with-a-built-in-registry-at-a-random-port)

### [Pros](https://docs.tilt.dev/choosing_clusters.html#pros)

### [优点](https://docs.tilt.dev/choosing_clusters.html#pros)

- Creating a new cluster is fast (~20 seconds). Deleting a cluster is even faster.
- Much more robust than Docker for Mac. Uses containerd instead of docker-shim. Short-lived clusters tend to be more reliable.
- Supports a local image registry (with our [Kind setup tool](https://github.com/tilt-dev/ctlptl#kind-with-a-built-in-registry-at-a-random-port)) . Pushing images is fast. No fiddling with image registry auth credentials.
- Can run in [most CI environments](https://github.com/kind-ci/examples) (TravisCI, CircleCI, etc.)

- 创建新集群很快（约 20 秒）。删除集群甚至更快。
- 比 Docker for Mac 强大得多。使用 containerd 而不是 docker-shim。短期集群往往更可靠。
- 支持本地镜像注册表（使用我们的 [Kind 设置工具](https://github.com/tilt-dev/ctlptl#kind-with-a-built-in-registry-at-a-random-port)) .推送图像很快。无需摆弄图像注册表身份验证凭据。
- 可以在[大多数 CI 环境](https://github.com/kind-ci/examples) 中运行（TravisCI、CircleCI 等)

### [Cons](https://docs.tilt.dev/choosing_clusters.html#cons)

### [缺点](https://docs.tilt.dev/choosing_clusters.html#cons)

- The local registry setup is still new, and changing rapidly. You need to be using Tilt v0.12.0+
- If Tilt can’t find the registry, it will use the much slower `kind load` to load images. (This con is mitigated if you use Kind with a local registry, as described in the instructions linked above.)

- 本地注册表设置仍然是新的，并且变化很快。您需要使用 Tilt v0.12.0+
- 如果 Tilt 找不到注册表，它将使用慢得多的 `kind load` 来加载图像。 （如上面链接的说明中所述，如果您将 Kind 与本地注册表一起使用，则会减轻此问题。）

------

## [Docker for Desktop](https://docs.tilt.dev/choosing_clusters.html#docker-for-desktop)

## [桌面Docker](https://docs.tilt.dev/choosing_clusters.html#docker-for-desktop)

Docker for Desktop is the easiest to get started with if you’re on MacOS.

如果您使用的是 MacOS，桌面版 Docker 是最容易上手的。

In the Docker For Mac preferences, click [Enable Kubernetes](https://docs.docker.com/docker-for-mac/#kubernetes)

在 Docker For Mac 首选项中，单击 [启用 Kubernetes](https://docs.docker.com/docker-for-mac/#kubernetes)

### [Pros](https://docs.tilt.dev/choosing_clusters.html#pros-1)

### [优点](https://docs.tilt.dev/choosing_clusters.html#pros-1)

- Widely used and supported.
- Nothing else to install.
- Built images are immediately available in-cluster. No pushing and pulling from image registries.

- 广泛使用和支持。
- 没有其他要安装的。
- 构建的图像可立即在集群中使用。无需从映像注册表中推送和拉取。

### [Cons](https://docs.tilt.dev/choosing_clusters.html#cons-1)

### [缺点](https://docs.tilt.dev/choosing_clusters.html#cons-1)

- If Kubernetes breaks, it’s easier to reset the whole thing than debug it.
- Much more resource-intensive because it uses docker-shim as the container runtime.
- Different defaults than a prod cluster and difficult to customize.
- Not available on Linux.

- 如果 Kubernetes 出现故障，重置整个事情比调试更容易。
- 更多资源密集型，因为它使用 docker-shim 作为容器运行时。
- 与 prod 集群不同的默认值并且难以自定义。
- 在 Linux 上不可用。

------

## [MicroK8s](https://docs.tilt.dev/choosing_clusters.html#microk8s)

## [MicroK8s](https://docs.tilt.dev/choosing_clusters.html#microk8s)

[Microk8s](https://microk8s.io) is what we recommend most often for Ubuntu users.

[Microk8s](https://microk8s.io) 是我们最常推荐给 Ubuntu 用户的。

Install:

```
sudo snap install microk8s --classic && \
sudo microk8s.enable dns && \
sudo microk8s.enable registry
```

Make microk8s your local Kubernetes cluster:

将 microk8s 设为您的本地 Kubernetes 集群：

```
sudo microk8s.kubectl config view --flatten > ~/.kube/microk8s-config && \
KUBECONFIG=~/.kube/microk8s-config:~/.kube/config kubectl config view --flatten > ~/.kube/temp-config && \
mv ~/.kube/temp-config ~/.kube/config && \
kubectl config use-context microk8s
```

### [Pros](https://docs.tilt.dev/choosing_clusters.html#pros-2)

### [优点](https://docs.tilt.dev/choosing_clusters.html#pros-2)

- No virtual machine overhead on Linux
- Ships with plugins that make common configs as easy as `microk8s.enable`
- Supports a local image registry with `microk8s.enable registry`. Pushing images is fast. No fiddling with image registry auth credentials.

- Linux 上没有虚拟机开销
- 附带插件，使常见的配置像 `microk8s.enable` 一样简单
- 支持带有“microk8s.enable registry”的本地镜像注册表。推送图像很快。无需摆弄图像注册表身份验证凭据。

### [Cons](https://docs.tilt.dev/choosing_clusters.html#cons-2)

### [缺点](https://docs.tilt.dev/choosing_clusters.html#cons-2)

- Resetting the cluster is slow and error-prone. 

- 重置集群速度慢且容易出错。

- Optimized for Ubuntu. In theory, works on any Linux that supports [Snap](https://snapcraft.io/) and on MacOS/Windows with [Multipass](https://multipass.run/), but it's not as stable on those platforms .

- 针对 Ubuntu 进行了优化。理论上，适用于任何支持 [Snap](https://snapcraft.io/) 的 Linux 和具有 [Multipass](https://multipass.run/) 的 MacOS/Windows，但在这些平台上并不稳定.

------

## [Minikube](https://docs.tilt.dev/choosing_clusters.html#minikube)

[Minikube](https://github.com/kubernetes/minikube) is what we recommend when you’re willing to pay some overhead for a more high-fidelity cluster.

[Minikube](https://github.com/kubernetes/minikube) 当您愿意为更高保真度的集群支付一些开销时，我们推荐您这样做。

Minikube has tons of options for customizing the cluster. You can choose between a VM and a Docker container for running a machine, choose from different container runtimes, and more.

Minikube 有大量用于自定义集群的选项。您可以在 VM 和 Docker 容器之间进行选择以运行机器，从不同的容器运行时中进行选择等等。

Follow these instructions to set up Minikube for use with Tilt:

按照以下说明设置 Minikube 以与 Tilt 一起使用：

[**Minikube Setup**](https://github.com/tilt-dev/ctlptl#minikube-with-a-built-in-registry)

### [Pros](https://docs.tilt.dev/choosing_clusters.html#pros-3)

### [优点](https://docs.tilt.dev/choosing_clusters.html#pros-3)

- The most full-featured local Kubernetes solution
- Can easily run different Kubernetes versions, container runtimes, and controllers
- Supports a local image registry (with our [cluster setup tool](https://github.com/tilt-dev/ctlptl#minikube-with-a-built-in-registry)). Pushing images is fast. No fiddling with image registry auth credentials.

- 功能最全的本地Kubernetes解决方案
- 可以轻松运行不同的 Kubernetes 版本、容器运行时和控制器
- 支持本地镜像注册表（使用我们的 [集群设置工具](https://github.com/tilt-dev/ctlptl#minikube-with-a-built-in-registry))。推送图像很快。无需摆弄图像注册表身份验证凭据。

### [Cons](https://docs.tilt.dev/choosing_clusters.html#cons-3)

### [缺点](https://docs.tilt.dev/choosing_clusters.html#cons-3)

- We often see engineers struggle to set it up the first time, getting lost in a maze of VM drivers that they’re unfamiliar with.
- Minikube gives you lots of options, but many of them are difficult to use or require manual setup.
- Beware if using a VM instead of a Docker container to run your cluster. You’ll usually want to shut down Minikube when you’re finished because of the VM’s drain on your resources.

- 我们经常看到工程师在第一次设置时很费劲，迷失在他们不熟悉的 VM 驱动程序的迷宫中。
- Minikube 为您提供了很多选项，但其中很多都难以使用或需要手动设置。
- 如果使用 VM 而不是 Docker 容器来运行集群，请注意。由于 VM 耗尽了您的资源，您通常希望在完成后关闭 Minikube。

------

## [k3d](https://docs.tilt.dev/choosing_clusters.html#k3d)

[k3d](https://github.com/rancher/k3d) runs [k3s](https://k3s.io/), a lightweight Kubernetes distro, inside a Docker container.

[k3d](https://github.com/rancher/k3d) 在 Docker 容器内运行 [k3s](https://k3s.io/)，这是一个轻量级的 Kubernetes 发行版。

k3s is fully compliant with “full” Kubernetes, but has a lot of optional and legacy features removed.

k3s 完全符合“完整”Kubernetes，但删除了许多可选和遗留功能。

### [Pros](https://docs.tilt.dev/choosing_clusters.html#pros-4)

### [优点](https://docs.tilt.dev/choosing_clusters.html#pros-4)

- Extremely fast to start up (less than 5 seconds on most machines)
- k3d has [a built-in local registry](https://k3d.io/usage/guides/registries/#using-a-local-registry) that’s explicitly designed to work well with Tilt. Start k3d with the local registry to make pushing and pulling images fast.

- 启动速度极快（在大多数机器上不到 5 秒）
- k3d 有 [一个内置的本地注册表](https://k3d.io/usage/guides/registries/#using-a-local-registry)，它明确设计为与 Tilt 配合良好。使用本地注册表启动 k3d，可以快速推送和拉取镜像。

### [Cons](https://docs.tilt.dev/choosing_clusters.html#cons-4)

### [缺点](https://docs.tilt.dev/choosing_clusters.html#cons-4)

- The least widely used. That’s not *necessarily* bad. Just be aware that there’s less documentation on its pitfalls. Tools (including the Tilt team!) tend to be slower to add support for it.

- 最不广泛使用。这不是*必然*坏的。请注意，关于其陷阱的文档较少。工具（包括 Tilt 团队！）添加对它的支持的速度往往较慢。

------

## [Remote](https://docs.tilt.dev/choosing_clusters.html#remote)

## [远程](https://docs.tilt.dev/choosing_clusters.html#remote)

### [(EKS, AKS, GKE, and ](https://docs.tilt.dev/choosing_clusters.html#eks-aks-gke-and-custom-clusters)[custom clusters](https://medium.com/@cfatechblog/bare-metal-k8s-clustering-at-chick-fil-a-scale-7b0607bd3541))

### [(EKS、AKS、GKE 和 ](https://docs.tilt.dev/choosing_clusters.html#eks-aks-gke-and-custom-clusters)[自定义集群](https://medium.com/@cfatechblog/bare-metal-k8s-clustering-at-chick-fil-a-scale-7b0607bd3541))

By default, Tilt will not let you develop against a remote cluster.

默认情况下，Tilt 不允许您针对远程集群进行开发。

If you start Tilt while you have `kubectl` configured to talk to a remote cluster, you will get an error. You have to explicitly allow the cluster  using [allow_k8s_contexts](https://docs.tilt.dev/api.html#api.allow_k8s_contexts):

如果在将 `kubectl` 配置为与远程集群通信时启动 Tilt，您将收到错误消息。您必须使用 [allow_k8s_contexts](https://docs.tilt.dev/api.html#api.allow_k8s_contexts) 明确允许集群：

```
allow_k8s_contexts('my-cluster-name')
```

If your team connects to many remote dev clusters, a common approach is to disable the check entirely and add your own validation:

如果您的团队连接到许多远程开发集群，常见的方法是完全禁用检查并添加您自己的验证：

```
allow_k8s_contexts(k8s_context())
local('./validate-dev-cluster.sh')
```

We only recommend remote clusters for large engineering teams where a dedicated dev infrastructure team can maintain your dev cluster.

我们只为大型工程团队推荐远程集群，其中专门的开发基础架构团队可以维护您的开发集群。

Or if you need to debug something that only reproduces in a complete cluster.

或者，如果您需要调试仅在完整集群中重现的内容。

### [Pros](https://docs.tilt.dev/choosing_clusters.html#pros-5)

### [优点](https://docs.tilt.dev/choosing_clusters.html#pros-5)

- Can customize to your heart’s desire
- Share common services (e.g., a dev database) across developers
- Use a cheap laptop and the most expensive cloud instance you can buy for development

- 可以根据您的心愿定制
- 在开发人员之间共享公共服务（例如，开发数据库）
- 使用便宜的笔记本电脑和您可以购买的最昂贵的云实例进行开发

### [Cons](https://docs.tilt.dev/choosing_clusters.html#cons-5)

### [缺点](https://docs.tilt.dev/choosing_clusters.html#cons-5)

- Need to use a remote image registry. Make sure you have Tilt’s [live_update](https://docs.tilt.dev/live_update_tutorial.html) set up!
- Need to set up namespaces and access control so that each dev has their own sandbox
- If the cluster needs to be reset, we hope you’re good friends with your DevOps team

- 需要使用远程映像注册表。确保您已设置 Tilt 的 [live_update](https://docs.tilt.dev/live_update_tutorial.html)！
- 需要设置命名空间和访问控制，以便每个开发人员都有自己的沙箱
- 如果集群需要重置，我们希望您与您的 DevOps 团队成为好朋友

------

## [Custom Clusters](https://docs.tilt.dev/choosing_clusters.html#custom-clusters) 

## [自定义集群](https://docs.tilt.dev/choosing_clusters.html#custom-clusters)

If you’re rolling your own Kubernetes dev cluster, and want it to work with Tilt, there are two things you need to do.

如果您正在滚动自己的 Kubernetes 开发集群，并希望它与 Tilt 一起使用，您需要做两件事。

- Tilt needs to recognize the cluster as a dev cluster.
- Tilt needs to be able to discover any in-cluster registry.

- Tilt 需要将集群识别为开发集群。
- Tilt 需要能够发现任何集群内注册表。

### [Allowing the Cluster](https://docs.tilt.dev/choosing_clusters.html#allowing-the-cluster)

### [允许集群](https://docs.tilt.dev/choosing_clusters.html#allowing-the-cluster)

Users have to explicitly allow the cluster with this line in their Tiltfile:

用户必须在其 Tiltfile 中明确允许使用此行的集群：

```
allow_k8s_contexts('my-cluster-name')
```

If your cluster is a dev-only cluster that you think Tilt should recognize automatically, we accept PRs to allow the cluster in Tilt. Here’s an example:

如果您的集群是您认为 Tilt 应自动识别的仅开发集群，我们接受 PR 以允许该集群在 Tilt 中使用。下面是一个例子：

[Recognize Red Hat CodeReady Containers as a local cluster](https://github.com/tilt-dev/tilt/pull/3242)

 [将 Red Hat CodeReady Containers 识别为本地集群](https://github.com/tilt-dev/tilt/pull/3242)

### [Discovering the Registry](https://docs.tilt.dev/choosing_clusters.html#discovering-the-registry)

### [发现注册表](https://docs.tilt.dev/choosing_clusters.html#discovering-the-registry)

A local registry is often the fastest way to speed up your dev experience.

本地注册表通常是加快开发体验的最快方式。

Every cluster sets up this registry slightly differently.

每个集群设置此注册表的方式略有不同。

Tilt-team is currently collaborating with the Kubernetes community on protocols for discovery, so that multi-service development tools like Tilt will auto-configure when a local registry is present.

Tilt-team 目前正在与 Kubernetes 社区就发现协议进行合作，以便在存在本地注册表时，像 Tilt 这样的多服务开发工具将自动配置。

Tilt currently supports two generic protocols for discovering your cluster’s local registry, so you don’t have to do any configuration yourself. The Kubernetes standard protocol is a [KEP](https://github.com/kubernetes/enhancements/tree/master/keps#kubernetes-enhancement-proposals-keps) that has been vetted by the Kubernetes community. The annotation-based protocol is used in legacy Tilt scripts.

Tilt 目前支持两种通用协议来发现集群的本地注册表，因此您无需自己进行任何配置。 Kubernetes 标准协议是经过 Kubernetes 社区审核的 [KEP](https://github.com/kubernetes/enhancements/tree/master/keps#kubernetes-enhancement-proposals-keps)。旧版 Tilt 脚本中使用了基于注释的协议。

You can configure the registry manually in your Tiltfile if these options fail. We’ve documented all the options below.

如果这些选项失败，您可以在 Tiltfile 中手动配置注册表。我们已经记录了以下所有选项。

#### Kubernetes Standard Registry Discovery

#### Kubernetes 标准注册表发现

The standard protocol uses configmaps in the `kube-public` namespace of your cluster.

标准协议在集群的 `kube-public` 命名空间中使用配置映射。

If you have a local registry running at `localhost:5000`, apply the following config map to your cluster:

如果您有一个在“localhost:5000”上运行的本地注册表，请将以下配置映射应用于您的集群：

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:5000"
```

Tilt will automatically detect your local registry, and will push and pull images from it.

Tilt 将自动检测您的本地注册表，并从中推送和拉取图像。

For more details on how to use this configmap, see

有关如何使用此配置映射的更多详细信息，请参阅

- [The Kubernetes Enhancement Proposal](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry)
- [A sample implementation in Go](https://github.com/tilt-dev/localregistry-go)

- [Kubernetes 增强提案](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry)
- [Go 中的示例实现](https://github.com/tilt-dev/localregistry-go)

We’re working with local development cluster teams to ensure that clusters support this protocol when they have a built-in registry.

我们正在与本地开发集群团队合作，以确保集群在具有内置注册表时支持此协议。

#### Legacy Annotation-based Registry Discovery

#### 传统的基于注解的注册表发现

To discover the registry, Tilt reads two annotatons from the node of your Kubernetes cluster:

为了发现注册表，Tilt 从 Kubernetes 集群的节点读取两个注释：

- `tilt.dev/registry`: The host of the registry, as seen by your local machine.
- `tilt.dev/registry-from-cluster`: The host of the registry, as seen by your cluster. If omitted, Tilt will assume that the host is the same as `tilt.dev/registry`.

- `tilt.dev/registry`：注册表的主机，如本地机器所见。
- `tilt.dev/registry-from-cluster`：注册表的主机，如您的集群所见。如果省略，Tilt 将假定主机与`tilt.dev/registry` 相同。

Our cluster-specific setup scripts often have a shell script snippet like:

我们特定于集群的设置脚本通常有一个 shell 脚本片段，例如：

```
nodes=$(kubectl get nodes -o go-template --template='{{range .items}}{{printf "%s\n" .metadata.name}}{{end}}')
if [ !-z $nodes ];then
  for node in $nodes;do
    kubectl annotate node "${node}" \
        tilt.dev/registry=localhost:5000 \
        tilt.dev/registry-from-cluster=registry:5000
  done
fi
```

to help Tilt find the registry.

帮助 Tilt 找到注册表。

#### Manual Configuration

#### 手动配置

You can manually configure the registry in your Tiltfile with [`default_registry`](https://docs.tilt.dev/api.html#api.default_registry).

您可以使用 [`default_registry`](https://docs.tilt.dev/api.html#api.default_registry) 在 Tiltfile 中手动配置注册表。

```
default_registry('gcr.io/my-personal-registry')
```

Because the Tiltfile is scriptable, you can configure this to fit your team’s conventions:

由于 Tiltfile 是可编写脚本的，因此您可以对其进行配置以符合您团队的约定：

```
reg = os.environ.get('MY_PERSONAL_REGISTRY', '')
if reg:
  default_registry(reg)
```

