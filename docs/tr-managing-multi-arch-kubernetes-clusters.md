# Getting up and running with multi-arch Kubernetes clusters

# 启动并运行多架构 Kubernetes 集群

*February 20, 2021*



The world of ARM processors has been getting very interesting over the last few years. Until fairly recently, for most people, ARM CPUs were reserved for their phone or maybe a [Raspberry Pi running their home DNS](https://pi-hole.net/). However now the Raspberry Pi 4 has a pretty decent quad-core CPU and up to 8GB RAM, Apple have blown away the industry with the M1 chips and AWS have launched [Graviton2](https://aws.amazon.com/ec2/graviton/) instances which depending on who you ask have 20-40% better price/performance than the Intel equivalents.

过去几年，ARM 处理器的世界变得非常有趣。直到最近，对于大多数人来说，ARM CPU 都是为他们的手机或 [Raspberry Pi 运行他们的家庭 DNS](https://pi-hole.net/) 保留的。然而，现在树莓派 4 拥有相当不错的四核 CPU 和高达 8GB 的 RAM，Apple 凭借 M1 芯片震撼了整个行业，AWS 推出了 [Graviton2](https://aws.amazon.com/ec2/graviton/) 实例，这取决于您询问的对象，其性价比比英特尔同类产品高 20-40%。

Despite all this, until recently I wasn’t convinced trying to use arm64 nodes in a production Kubernetes cluster was worth the effort. However the combination of all these events has caused a lot of the software to catch up and support for arm64 in Kubernetes distributions from Amazon EKS to k3s is now excellent! It is likely that a good proportion of the container images you’re using in production now support “multi-arch” and will run on both arm64 and amd64 machines.

尽管如此，直到最近我才相信在生产 Kubernetes 集群中尝试使用 arm64 节点是值得的。然而，所有这些事件的结合导致许多软件迎头赶上，并且从 Amazon EKS 到 k3s 的 Kubernetes 发行版中对 arm64 的支持现在非常出色！您在生产中使用的大部分容器镜像现在可能支持“多架构”，并且可以在 arm64 和 amd64 机器上运行。

## Container Images

## 容器镜像

Let’s start with the container images, as that’s what you’ll need to tackle first if you want to start using arm64 in your environments. Historically if you wanted to support multiple CPU architectures for your container image you would have to have two separate tags e.g. myimage:1.2.3-amd64 and myimage:1.2.3-arm64. This was both a pain for people building the images and for people trying to use them; if you're deploying a Helm Chart for example you would have to override the image tags if you wanted to run on ARM and there was no hope for you if you wanted to use mixture of arm64 and amd64 nodes in a cluster and have the same pods seamlessly schedule on either (without some extra tooling).

让我们从容器镜像开始，因为如果您想在您的环境中开始使用 arm64，这就是您首先需要解决的问题。从历史上看，如果您想为容器映像支持多个 CPU 架构，则必须有两个单独的标签，例如myimage:1.2.3-amd64 和 myimage:1.2.3-arm64。这对构建图像的人和试图使用它们的人来说都是一种痛苦；例如，如果您正在部署 Helm Chart，如果您想在 ARM 上运行，则必须覆盖图像标签，如果您想在集群中混合使用 arm64 和 amd64 节点并具有相同的Pod 可以无缝地安排在任何一个上（没有一些额外的工具）。

This problem has gone away with Manifest Lists in the [Image Manifest V2 spec](https://docs.docker.com/registry/spec/manifest-v2-2/). This allows you to specify a list of container images for a number of different architectures in a single “Manifest List”. In newer container runtime versions (like Docker) this means if you do a `docker pull nginx:alpine` on your Raspberry Pi you'll get an image for arm64 and on your Intel laptop you'll get an amd64 image without any further effort . Previously you would have got the godforsaken “exec format error”.

[Image Manifest V2 规范](https://docs.docker.com/registry/spec/manifest-v2-2/) 中的清单列表已经解决了这个问题。这允许您在单个“清单列表”中为多个不同架构指定容器镜像列表。在较新的容器运行时版本（如 Docker)中，这意味着如果您在 Raspberry Pi 上执行 `docker pull nginx:alpine`，您将获得 arm64 的图像，而在您的英特尔笔记本电脑上，您将获得 amd64 图像而无需任何进一步的努力.以前，您会遇到该死的“exec 格式错误”。

You might think that this means we’ve reached multi-arch nirvana and everything will “just work” now, but unfortunately this is not the case. If you’re using an existing public image, you will need to make sure that it is using a manifest list and supports both amd64 and arm64. Some registries make this really easy, such as Docker Hub, where you’ll get a nice list of architectures on the Tags tab.

你可能认为这意味着我们已经达到了多架构的必杀技，现在一切都将“正常工作”，但不幸的是，事实并非如此。如果您使用现有的公共映像，则需要确保它使用清单列表并支持 amd64 和 arm64。一些注册中心让这一切变得非常简单，例如 Docker Hub，您将在标签选项卡上获得一个不错的架构列表。

However some don’t make it obvious at all! The easiest way I've found so far to determine if an image is multi-arch, is to use the experimental [docker manifest command](https://docs.docker.com/engine/reference/commandline/manifest/) in the latest Docker versions. As it says in the documentation you will have to enable experimental features in `~/.docker/config.json` then you will be able to run a command like:

然而，有些根本不明显！到目前为止，我发现确定图像是否为多架构的最简单方法是使用实验性的 [docker manifest 命令](https://docs.docker.com/engine/reference/commandline/manifest/)最新的 Docker 版本。正如文档中所说，您必须在`~/.docker/config.json` 中启用实验性功能，然后您将能够运行如下命令：

```
$ docker manifest inspect nginx:alpine
{
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.list.v2+json",
   "manifests": [
      {
         "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
         "size": 1568,
         "digest": "sha256:9a39c77d9ea3a9ddc41535f875b7610a0f121df3c2496c16f2a3a5fcb0e43e4f",
         "platform": {
            "architecture": "amd64",
            "os": "linux"
         }
      },
      {
         "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
         "size": 1568,
         "digest": "sha256:22d2c4a5220232818a0fe7a5d3651c846bc3e7d2ff8dbfc2f665c717f0e43a69",
         "platform": {
            "architecture": "arm64",
            "os": "linux",
            "variant": "v8"
         }
      },
```

As you can see the nginx:alpine image is a “manifest list” rather than a plain old manifest and supports the two architectures we’re after. Great! However distroless-java is still single architecture:

正如你所看到的 nginx:alpine 图像是一个“清单列表”而不是一个普通的旧清单，并支持我们所追求的两种架构。伟大的！然而 distroless-java 仍然是单一架构：

```
$ docker manifest inspect gcr.io/distroless/java
{
    "schemaVersion": 2,
    "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
    "config": {
        "mediaType": "application/vnd.docker.container.image.v1+json",
        "size": 1164,
        "digest": "sha256:85cdcf63cad1cfe5373c68f78f21f0c6349fee87fbb40bc9a9dc7d560f52438b"
    },
    "layers": [
        {
            "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "size": 643695,
            "digest": "sha256:9e4425256ce4503b2a009683be02372ee51d411e9cc3547919e064fee4970eab"
        },
...
```

I won’t go into building these multi-architecture manifest lists here, as it very much depends on which tool you’re using. If you’re using Docker to build your images you can use the [docker buildx](https://docs.docker.com/buildx/working-with-buildx/) experimental feature. We're currently using the [jib-maven-plugin](https://github.com/GoogleContainerTools/jib/tree/master/jib-maven-plugin#extended-usage) for our Java-based apps, which has recently added the `platforms` feature. There are many other ways to do it and GitHub actions you can use, so it’s not too hard, and you don’t need an arm64 machine to build an arm64-compatible image anymore thanks to integration with QEMU.

我不会在这里构建这些多架构清单列表，因为这在很大程度上取决于您使用的工具。如果您使用 Docker 构建镜像，您可以使用 [docker buildx](https://docs.docker.com/buildx/working-with-buildx/) 实验功能。我们目前正在为基于 Java 的应用程序使用 [jib-maven-plugin](https://github.com/GoogleContainerTools/jib/tree/master/jib-maven-plugin#extended-usage)，最近添加了“平台”功能。还有很多其他方法可以做到这一点，并且您可以使用 GitHub 操作，所以这并不太难，而且由于与 QEMU 集成，您不再需要 arm64 机器来构建与 arm64 兼容的映像。

## Kubernetes

Getting arm64 nodes in your cluster is as simple as just creating an extra autoscaling group (or one per Availability Zone) in AWS or using [k3sup](https://github.com/alexellis/k3sup) to join your Raspberry Pi 4 to your k3s cluster. I'll try to keep this post generic to however you deploy Kubernetes, but as we're using EKS in production, [here is the documentation](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html#arm-ami) on getting the right AMI for your architecture. I also have a k3s cluster at home with two old Intel laptops and a Raspberry Pi 4, and this guide works exactly the same on that.

在您的集群中获取 arm64 节点就像在 AWS 中创建一个额外的自动缩放组（或每个可用区一个）或使用 [k3sup](https://github.com/alexellis/k3sup) 加入您的 Raspberry Pi 4 一样简单您的 k3s 集群。我会尽量使这篇文章适用于您部署 Kubernetes 的任何方式，但由于我们在生产中使用 EKS，[这里是文档](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html#arm-ami) 获取适合您的架构的 AMI。我家里还有一个 k3s 集群，里面有两台旧的 Intel 笔记本电脑和一台 Raspberry Pi 4，本指南的工作原理完全相同。

Before you add arm64 nodes to your cluster you must consider whether you want to specifically exclude incompatible Pods from running on these nodes with [node affinity rules](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes-using-node-affinity/), or if you want to exclude all Pods by default and specifically allow the ones you know work on arm64 with [taints and tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/). I went with the latter on our EKS clusters because the majority of our workloads aren’t yet multi-arch, but on my home Raspberry Pi cluster I went with the former.

在将 arm64 节点添加到集群之前，您必须考虑是否要使用 [节点关联规则](https://kubernetes.io/docs/tasks/configure-pod-container/assign) 专门排除不兼容的 Pod 在这些节点上运行-pods-nodes-using-node-affinity/)，或者如果您想默认排除所有 Pod 并特别允许您知道的那些在 arm64 上使用 [污点和容忍](https://kubernetes.io/docs/概念/调度驱逐/污点和容忍/)。我在我们的 EKS 集群上选择了后者，因为我们的大部分工作负载还不是多架构的，但在我的家庭 Raspberry Pi 集群上，我选择了前者。

### Node Affinity

### 节点亲和力

The node affinity rule option requires no configuration of the nodes themselves because recent Kubernetes versions have a standard `kubernetes.io/arch` node label; this should be set to arm64 or amd64. However it will require a lot of effort if you have a large number of workloads which don’t yet work on arm64.

节点关联规则选项不需要对节点本身进行配置，因为最近的 Kubernetes 版本有一个标准的 `kubernetes.io/arch` 节点标签；这应该设置为 arm64 或 amd64。但是，如果您有大量尚未在 arm64 上工作的工作负载，则需要付出很多努力。

```
$ kubectl describe no myarmnode
Name:               myarmnode
Roles:              control-plane,etcd,master
Labels:             kubernetes.io/arch=arm64
                    kubernetes.io/hostname=myarmnode
                    kubernetes.io/os=linux
```

All you need to do is set up a [node affinity rule](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes-using-node-affinity/) in the PodSpec of any Pods which don't have multi-arch images like this:

您需要做的就是在 PodSpec 的任何没有像这样的多架构图像的 Pod：

```
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
        - matchExpressions:
          - key: "kubernetes.io/arch"
            operator: In
            values: ["amd64"]
```

You could also do the opposite and use the `NotIn` operator and values set as `["arm64"]`.

你也可以做相反的事情，使用 `NotIn` 运算符和设置为 `["arm64"]` 的值。

### Taints and Tolerations

### 污点和容忍

The option I went with in our EKS clusters is to set up a `NoSchedule` “taint” on all arm64 nodes which we then “tolerate” on any pods we know to work on arm64. I've done this in the user-data script in the Launch Template used by our Graviton2 Auto Scaling Groups through the `--kubelet-extra-args` flag of [bootstrap.sh](https://aws.amazon.com/blogs/opensource/improvements-eks-worker-node-provisioning/); the extra arg you need to pass to the kubelet is `--register-with-taints="arch=arm64:NoSchedule"`. You can also just use this command after your nodes are registered to the cluster: `kubectl taint no myarmnode arch=arm64:NoSchedule`.

我在我们的 EKS 集群中采用的选项是在所有 arm64 节点上设置一个“NoSchedule”“污点”，然后我们“容忍”我们知道在 arm64 上工作的任何 pod。我已经通过 [bootstrap.sh](https://aws.amazon.com/blogs/opensource/improvements-eks-worker-node-provisioning/);您需要传递给 kubelet 的额外参数是 `--register-with-taints="arch=arm64:NoSchedule"`。你也可以在你的节点注册到集群后使用这个命令：`kubectl taint no myarmnode arch=arm64:NoSchedule`。

Once all your arm64 nodes are tainted, only Pods with the right [tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) will be scheduled to run on them. As with affinity rules, they are specified in the PodSpec of your Pod. In our case the toleration needed will be:

一旦所有 arm64 节点都被污染，只有具有正确 [tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) 的 Pod 才会被安排在它们上运行。与关联规则一样，它们在 Pod 的 PodSpec 中指定。在我们的例子中，所需的容忍度是：

```
tolerations:
- key: "arch"
  operator: "Equal"
  value: "arm64"
  effect: "NoSchedule"
```

You might see some system-level DaemonSets like kube-proxy already have tolerations like this, which will also work:

你可能会看到一些系统级的 DaemonSets 像 kube-proxy 已经有这样的容忍度，它也可以工作：

```
tolerations:
- effect: NoSchedule
  operator: Exists
```

If you get issues with certain DaemonSets not being scheduled on your arm nodes, even though they have the right tolerations, check that the affinity rules don’t exclude nodes by the `kubernetes.io/arch` label.

如果您遇到某些 DaemonSet 未在您的 arm 节点上调度的问题，即使它们具有正确的容忍度，请检查关联规则是否没有通过 `kubernetes.io/arch` 标签排除节点。

### Finding multi-arch images

### 查找多架构图像

The first Pods I had to make sure were running on our arm64 nodes were the DaemonSets for things like fluentd (our logging agent) or jaeger-agent (our tracing agent). Unfortunately neither of these images were multi-arch which would have been a blocker for running our workloads on arm64 nodes. However as is often the case in the Open Source world, somebody had already had the same problem and a bit of searching GitHub showed up an open Pull Request or Issue with links to un-official images; [this in the case of fluentd](https://github.com/kokuwaio/helm-charts/issues/12) and [this in the case of jaeger-agent](https://github.com/querycap/jaeger). This is course not as good as the project providing official multi-arch images (like everything except kube-state-metrics in the [kube-prometheus-stack Helm Chart](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)) but in more and more cases official multi-arch images are available and this will only get better over time.

我必须确保在我们的 arm64 节点上运行的第一个 Pod 是用于 fluentd（我们的日志代理）或 jaeger-agent（我们的跟踪代理）之类的 DaemonSet。不幸的是，这些镜像都不是多架构的，这会阻碍我们在 arm64 节点上运行我们的工作负载。然而，正如在开源世界中经常发生的情况一样，有人已经遇到了同样的问题，在 GitHub 上进行了一些搜索，结果显示了一个开放的 Pull Request 或 Issue，其中包含指向非官方图像的链接； [在 fluentd 的情况下](https://github.com/kokuwaio/helm-charts/issues/12) 和 [在 jaeger-agent 的情况下](https://github.com/querycap/jaeger)。这当然不如提供官方多架构图像的项目（就像[kube-prometheus-stack Helm Chart](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack))，但在越来越多的情况下，官方多架构图像可用，而且随着时间的推移，这只会变得更好。

### Cluster autoscaler

### 集群自动缩放器

After I’d got a number of our workloads set up to run on either arm64 or amd64 nodes, I of course wanted to run as many arm64 nodes as possible due to the lower cost. I did this using the [Priority based expander for cluster-autoscaler](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/expander/priority/readme.md) by setting the Auto Scaling Groups for arm64 nodes to a higher priority than the rest of our groups. You could also set up `preferredDuringSchedulingIgnoredDuringExecution` affinity rules on arm64 compatible Pods but we found the Cluster Autoscaler configuration to be sufficient.

在我将许多工作负载设置为在 arm64 或 amd64 节点上运行后，我当然希望运行尽可能多的 arm64 节点，因为成本较低。我使用 [Priority based expander for cluster-autoscaler](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/expander/priority/readme.md) 通过设置 Auto Scaling Groups for arm64 节点的优先级高于我们其他组。您还可以在 arm64 兼容的 Pod 上设置 `preferredDuringSchedulingIgnoredDuringExecution` 关联规则，但我们发现 Cluster Autoscaler 配置就足够了。

### Bash one-liner for good measure

### Bash one-liner 以获得良好的测量

This will (on Linux at least) list out all the arm64 compatible images running in your Kubernetes cluster:

这将（至少在 Linux 上）列出在您的 Kubernetes 集群中运行的所有 arm64 兼容镜像：

```
kubectl get po -A -o yaml |grep 'image:' |cut -f2- -d':' |sed 's/^[[:space:]]*//g' |grep '/' |sort -u |xargs -I{} bash -c "docker manifest inspect {} | grep -q arm64 && echo {}"
```

Change `grep -q` to `grep -vq` to invert the logic and return images which won’t work on arm64 nodes.

将 `grep -q` 更改为 `grep -vq` 以反转逻辑并返回在 arm64 节点上不起作用的图像。

This would definitely be nicer using [yq](https://github.com/mikefarah/yq) rather than parsing YAML with grep, cut and sed…but I know most people don’t have yq installed.

使用 [yq](https://github.com/mikefarah/yq) 肯定会比使用 grep、cut 和 sed 解析 YAML 更好……但我知道大多数人没有安装 yq。

## Final thoughts

##  最后的想法

It is still a bit of effort to add arm64 nodes to your Kubernetes clusters. However for us it was worth the effort for the cost savings and it’s always exciting to be an early adopter. If anyone has any questions on how I've got things set up or would like to share different approaches I'm [on Twitter](https://twitter.com/cablespaghetti) and love to hear from other people working on similar challenges ! 

将 arm64 节点添加到您的 Kubernetes 集群仍然需要一些努力。然而，对我们来说，为节省成本而付出的努力是值得的，成为早期采用者总是令人兴奋的。如果有人对我如何设置事情有任何疑问，或者想分享不同的方法，我 [在 Twitter](https://twitter.com/cablespaghetti) 上，并很乐意听取其他人的类似挑战！

