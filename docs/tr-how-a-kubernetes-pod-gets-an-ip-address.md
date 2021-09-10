# How a Kubernetes Pod Gets an IP Address

# Kubernetes Pod 如何获取 IP 地址

Aug 21, 20209 min read [Kubernetes](http://ronaknathani.com/category/kubernetes/)

One of the core requirements of the [Kubernetes networking model](https://kubernetes.io/docs/concepts/cluster-administration/networking/#the-kubernetes-network-model) is that every pod should get its own IP address and that every pod in the cluster should be able to talk to it using this IP address. There are several network providers (flannel, calico, canal, etc.) that implement this networking model.
[Kubernetes 网络模型](https://kubernetes.io/docs/concepts/cluster-administration/networking/#the-kubernetes-network-model) 的核心要求之一是每个 pod 都应该有自己的 IP 地址，集群中的每个 pod应该能够使用这个 IP 地址与它交谈。有几个网络供应商（法兰绒、印花布、运河等)实现了这种网络模型。

As I started working on Kubernetes, it wasn’t completely clear to me how every pod is assigned an IP address. I understood how various components worked independently, however, it wasn’t clear how these components fit together. For instance, I understood what CNI plugins were, however, I didn’t know how they were invoked. So, I wanted to write this post to share what I have learned about various networking components and how they are stitched together in a kubernetes cluster for every pod to receive an IP address.

当我开始研究 Kubernetes 时，我并不完全清楚每个 pod 是如何分配 IP 地址的。我了解各种组件如何独立工作，但是，不清楚这些组件如何组合在一起。比如我知道什么是CNI插件，但是不知道是怎么调用的。所以，我想写这篇文章来分享我对各种网络组件的了解，以及它们如何在 kubernetes 集群中拼接在一起，以便每个 pod 接收 IP 地址。

There are various ways of setting up networking in kubernetes and various options for a container runtime. For this post, I will use [Flannel](https://github.com/coreos/flannel) as the network provider and [Containerd](https://github.com/containerd/containerd) as the container runtime. Also, I am going to assume that you know how container networking works and only share a very brief overview below for context.

在 kubernetes 中有多种设置网络的方法以及容器运行时的各种选项。对于这篇文章，我将使用[Flannel](https://github.com/coreos/flannel) 作为网络提供者和[Containerd](https://github.com/containerd/containerd) 作为容器运行时。此外，我将假设您知道容器网络的工作原理，并且仅在下面分享一个非常简短的概述以了解上下文。

## Some Background Concepts

## 一些背景概念

### Container Networking: A Very Brief Overview

### 容器网络：非常简要的概述

There are some really good posts explaining how container networking works. For context, I will go over a very high level overview here with a single approach that involves linux bridge networking and packet encapsulation. I am skipping details here as container networking deserves a blog post of itself. Some of the posts that I have found to be very educational in this space are
[linked in the references below](http://ronaknathani.com#container-networking).

有一些非常好的帖子解释了容器网络的工作原理。对于上下文，我将在这里使用涉及 linux 网桥网络和数据包封装的单一方法进行非常高级的概述。我在这里跳过细节，因为容器网络值得写一篇博文。我发现在这个领域非常有教育意义的一些帖子是
[在下面的参考文献中链接](http://ronaknathani.com#container-networking)。

#### Containers on the same host

#### 容器在同一主机上

One of the ways containers running on the same host can talk to each other via their IP addresses is through a linux bridge. In the kubernetes (and docker) world, a [veth (virtual ethernet)](https://man7.org/linux/man-pages/man4/veth.4.html) device is created to achieve this. One end of this veth device is inserted into the container network namespace and the other end is connected to a [linux bridge](https://wiki.archlinux.org/index.php/Network_bridge) on the host network. All containers on the same host have one end of this veth pair connected to the linux bridge and they can talk to each other using their IP addresses via the bridge. The linux bridge is also assigned an IP address and it acts as a gateway for egress traffic from pods destined to different nodes.
![bridge networking](http://ronaknathani.com/bridge-networking.png)

#### Containers on different hosts

运行在同一主机上的容器可以通过它们的 IP 地址相互通信的一种方式是通过 linux 网桥。在 kubernetes（和 docker）世界中，[veth（虚拟以太网）](https://man7.org/linux/man-pages/man4/veth.4.html) 设备就是为了实现这一点而创建的。这个 veth 设备的一端插入容器网络命名空间，另一端连接到一个[linux bridge](https://wiki.archlinux.org/index.php/Network_bridge) 在主机网络上。同一主机上的所有容器都将这个 veth 对的一端连接到 linux 网桥，它们可以通过网桥使用 IP 地址相互通信。 linux 网桥也被分配了一个 IP 地址，它充当从 pod 到不同节点的出口流量的网关。
#### 不同主机上的容器

One of the ways containers running on different hosts can talk to each other via their IP addresses is by using packet encapsulation. Flannel supports this through [vxlan](https://vincent.bernat.ch/en/blog/2017-vxlan-linux) which wraps the original packet inside a UDP packet and sends it to the destination.

运行在不同主机上的容器可以通过它们的 IP 地址相互通信的一种方式是使用数据包封装。法兰绒支持这一点
[vxlan](https://vincent.bernat.ch/en/blog/2017-vxlan-linux) 将原始数据包包装在 UDP 数据包中并将其发送到目的地。

In a kubernetes cluster, flannel creates a vxlan device and some route table entries on each of the nodes. Every packet that’s destined for a container on a different host goes through the vxlan device and is encapsulated in a UDP packet. On the destination, the encapsulated packet is retrieved and the packet is routed through to the destined pod.
![flannel networking](http://ronaknathani.com/flannel-networking.png)

_NOTE: This is just one of the ways how networking between containers can be configured._

在 kubernetes 集群中，flannel 在每个节点上创建一个 vxlan 设备和一些路由表条目。每个发往不同主机上的容器的数据包都经过 vxlan 设备并封装在 UDP 数据包中。在目的地，检索封装的数据包，并将数据包路由到目的地 pod。
_注意：这只是如何配置容器之间的网络的一种方式。_

### What Is CRI?

### 什么是 CRI？

[CRI (Container Runtime Interface)](https://github.com/kubernetes/cri-api) is a plugin interface that allows kubelet to use different container runtimes. Various container runtimes implement the CRI API and this allows users to use the container runtime of their choice in their kubernetes installation.

[CRI (Container Runtime Interface)](https://github.com/kubernetes/cri-api) 是一个插件接口，允许 kubelet 使用不同的容器运行时。各种容器运行时实现了 CRI API，这允许用户在他们的 kubernetes 安装中使用他们选择的容器运行时。

### What is CNI?

### 什么是CNI？

[CNI project](https://github.com/containernetworking/cni) includes a [spec](https://github.com/containernetworking/cni/blob/master/SPEC.md) to provide a generic plugin-based networking solution for linux containers. It also consists of various [plugins](https://github.com/containernetworking/plugins) which perform different functions in configuring the pod network. A CNI plugin is an executable that follows the CNI spec and we’ll discuss some plugins in the post below.

[CNI 项目](https://github.com/containernetworking/cni) 包括一个 [spec](https://github.com/containernetworking/cni/blob/master/SPEC.md) 为 linux 容器提供基于插件的通用网络解决方案。它还包括各种[plugins](https://github.com/containernetworking/plugins) 在配置 pod 网络时执行不同的功能。 CNI 插件是一个遵循 CNI 规范的可执行文件，我们将在下面的帖子中讨论一些插件。

## Assigning Subnets To Nodes For Pod IP Addresses

## 为 Pod IP 地址的节点分配子网

If all pods are required to have an IP address, it’s important to ensure that all pods across the entire cluster have a unique IP address. This is achieved by assigning each node a unique subnet from which pods are assigned IP addresses on that node.

如果所有 pod 都需要有一个 IP 地址，那么确保整个集群中的所有 pod 都有一个唯一的 IP 地址很重要。这是通过为每个节点分配一个唯一的子网来实现的，pod 从该子网分配到该节点上的 IP 地址。

### Node IPAM Controller

### 节点 IPAM 控制器

When `nodeipam` is passed as an option to the [kube-controller-manager's](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/) `--controllers` command line flag, it allocates each node a dedicated subnet (podCIDR) from the cluster CIDR (IP range for the cluster network). Since these podCIDRs are disjoint subnets, it allows assigning each pod a unique IP address.

当 `nodeipam` 作为选项传递给 [kube-controller-manager's](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/) `--controllers` 命令行标志，它为每个节点分配一个专用的来自集群 CIDR（集群网络的 IP 范围)的子网 (podCIDR)。由于这些 podCIDR 是不相交的子网，它允许为每个 pod 分配一个唯一的 IP 地址。

A kubernetes node is assigned a podCIDR when the node first registers with the cluster. To change the podCIDR allocated to nodes in a cluster, nodes need to be de-registered and then re-registered with any configuration changes first applied to the kubernetes control plane. `podCIDR` for a node can be listed using the following command.

当 kubernetes 节点首次向集群注册时，会为其分配一个 podCIDR。要更改分配给集群中节点的 podCIDR，需要取消注册节点，然后使用首先应用于 kubernetes 控制平面的任何配置更改重新注册。可以使用以下命令列出节点的 `podCIDR`。

```
$ kubectl get no <nodeName> -o json |jq '.spec.podCIDR'
10.244.0.0/24

```


## Kubelet, Container Runtime and CNI Plugins - how it’s all stitched together

## Kubelet、Container Runtime 和 CNI 插件——它们是如何拼接在一起的

When a pod is scheduled on a node, a lot of things happen to start up a pod. In this section, I’ll only focus on the interactions that relate to configuring network for the pod.

当一个 Pod 被调度到一个节点上时，会发生很多事情来启动一个 Pod。在本节中，我将只关注与为 pod 配置网络相关的交互。

Once a pod is scheduled on the node, the following interactions result in configuring the network and starting the application container.
![kubelet-cri-cni-flowchart](http://ronaknathani.com/kubelet-cri-cni-flowchart.png)

Ref:
[Containerd cri plugin architecture](https://github.com/containerd/cri/blob/v1.11.1/docs/architecture.md)

在节点上调度 Pod 后，以下交互将导致配置网络并启动应用程序容器。
参考：
 [Containerd cri 插件架构](https://github.com/containerd/cri/blob/v1.11.1/docs/architecture.md)

## Interactions between Container Runtime and CNI Plugins

## 容器运行时和 CNI 插件之间的交互

Every network provider has a CNI plugin which is invoked by the container runtime to configure network for a pod as it’s started. With containerd as the container runtime, [Containerd CRI plugin](https://github.com/containerd/cri) invokes the CNI plugin. Every network provider also has an agent that’s installed on each of the kubernetes node to configure pod networking. When the network provider agent is installed, it either ships with the CNI config or it creates one on the node which is then used by the CRI plugin to figure out which CNI plugin to call.

每个网络提供者都有一个 CNI 插件，容器运行时会调用该插件来为 pod 启动时配置网络。使用 containerd 作为容器运行时， [Containerd CRI插件](https://github.com/containerd/cri)调用CNI插件。每个网络提供商也有一个安装在每个 kubernetes 节点上的代理来配置 pod 网络。安装网络提供程序代理后，它要么随 CNI 配置一起提供，要么在节点上创建一个，然后 CRI 插件使用它来确定要调用哪个 CNI 插件。

The location for the CNI config file is configurable and the default value is `/etc/cni/net.d/<config-file>`. CNI plugins need to be shipped on every node by the cluster administrators. The location for CNI plugins is configurable as well and the default value is `/opt/cni/bin`.

CNI 配置文件的位置是可配置的，默认值为`/etc/cni/net.d/<config-file>`。 CNI 插件需要由集群管理员在每个节点上提供。 CNI 插件的位置也是可配置的，默认值为 `/opt/cni/bin`。

In case of containerd as the container runtime, path for CNI configuration and CNI plugin binaries can be specified under `[plugins."io.containerd.grpc.v1.cri".cni]` section of the [containerd config](https://github.com/containerd/cri/blob/master/docs/config.md).

如果使用 containerd 作为容器运行时，可以在 `[plugins."io.containerd.grpc.v1.cri".cni]` 部分指定 CNI 配置和 CNI 插件二进制文件的路径[容器配置](https://github.com/containerd/cri/blob/master/docs/config.md)。

Since we are referring to Flannel as the network provider here, I’ll talk a little about how Flannel is set up. Flanneld is the Flannel daemon and is typically installed on a kubernetes cluster as a daemonset with `install-cni` as an [init container](https://github.com/coreos/flannel/blob/master/Documentation/kube-flannel.yml#L172). The `install-cni` container creates the [CNI configuration file](https://gist.github.com/ronaknnathani/957a56210bd4fbd8e11120273c6b4ede) \- `/etc/cni/net.d/10-flannel.conflist` \- on each node. Flanneld creates a vxlan device, fetches networking metadata from the apiserver and watches for updates on pods. As pods are created, it distributes routes for all pods across the entire cluster and these routes allow pods to connect to each other via their IP addresses. For details on how flannel works, I recommend the
[linked references below](http://ronaknathani.com#how-flannel-works).

由于我们在这里将 Flannel 称为网络提供商，因此我将稍微谈谈 Flannel 是如何设置的。 Flanneld 是 Flannel 守护进程，通常作为守护进程安装在 kubernetes 集群上，使用 `install-cni` 作为
[初始化容器](https://github.com/coreos/flannel/blob/master/Documentation/kube-flannel.yml#L172)。 `install-cni` 容器创建[CNI 配置文件](https://gist.github.com/ronaknnathani/957a56210bd4fbd8e11120273c6b4ede) \- `/etc/cni/net.d/10-flannel.conflist` \- 在每个节点上。 Flanneld 创建一个 vxlan 设备，从 apiserver 获取网络元数据并监视 pod 上的更新。创建 Pod 时，它会为整个集群中的所有 Pod 分配路由，这些路由允许 Pod 通过其 IP 地址相互连接。有关法兰绒工作原理的详细信息，我推荐
[以下链接参考](http://ronaknathani.com#how-flannel-works)。

The interactions between Containerd CRI Plugin and CNI plugins can be visualized as follows:
![kubelet-cri-cni-interactions](http://ronaknathani.com/kubelet-cri-cni-interactions.png)

Containerd CRI Plugin 和 CNI 插件之间的交互可视化如下：


As described above, kubelet calls the Containerd CRI plugin in order to create a pod and Containerd CRI plugin calls the CNI plugin to configure network for the pod. The network provider CNI plugin calls other base CNI plugins to configure the network. The interactions between CNI plugins are described below.

如上所述，kubelet 调用 Containerd CRI 插件以创建 pod，Containerd CRI 插件调用 CNI 插件为 pod 配置网络。网络提供者 CNI 插件调用其他基础 CNI 插件来配置网络。 CNI 插件之间的交互描述如下。

### Interactions Between CNI Plugins

### CNI 插件之间的交互

There are various CNI plugins that help configure networking between containers on a host. For this post, we will refer to 3 plugins.

有各种 CNI 插件可以帮助配置主机上容器之间的网络。对于这篇文章，我们将参考 3 个插件。

#### Flannel CNI Plugin

#### 法兰绒CNI插件

When using Flannel as the network provider, the Containerd CRI plugin invokes the
[Flannel CNI plugin](https://github.com/containernetworking/plugins/tree/master/plugins/meta/flannel) using the CNI configuration file - `/etc/cni/net.d/10-flannel.conflist` .

当使用 Flannel 作为网络提供者时，Containerd CRI 插件会调用
[Flannel CNI 插件](https://github.com/containernetworking/plugins/tree/master/plugins/meta/flannel) 使用CNI 配置文件-`/etc/cni/net.d/10-flannel.conflist` .

```
$ cat /etc/cni/net.d/10-flannel.conflist
{
"name": "cni0",
"plugins": [
    {
      "type": "flannel",
      "delegate": {
         "ipMasq": false,
        "hairpinMode": true,
        "isDefaultGateway": true
      }
    }
]
}

```


The Fannel CNI plugin works in conjunction with Flanneld. When Flanneld starts up, it fetches the podCIDR and other network related details from the apiserver and stores them in a file - `/run/flannel/subnet.env`.

Fannel CNI 插件与 Flanneld 配合使用。当 Flanneld 启动时，它会从 apiserver 获取 podCIDR 和其他网络相关的详细信息，并将它们存储在一个文件中 - `/run/flannel/subnet.env`。

```
FLANNEL_NETWORK=10.244.0.0/16
FLANNEL_SUBNET=10.244.0.1/24
FLANNEL_MTU=1450
FLANNEL_IPMASQ=false

```


The Flannel CNI plugin uses the information in `/run/flannel/subnet.env` to configure and invoke the bridge CNI plugin.

Flannel CNI 插件使用 `/run/flannel/subnet.env` 中的信息来配置和调用桥接 CNI 插件。

#### Bridge CNI Plugin

#### 桥接 CNI 插件

Flannel CNI plugin calls the Bridge CNI plugin with the following configuration:

Flannel CNI 插件调用 Bridge CNI 插件，配置如下：

```
{
"name": "cni0",
"type": "bridge",
"mtu": 1450,
"ipMasq": false,
"isGateway": true,
"ipam": {
    "type": "host-local",
    "subnet": "10.244.0.0/24"
}
}

```


When
[Bridge CNI plugin](https://github.com/containernetworking/plugins/tree/master/plugins/main/bridge) is invoked for the first time, it creates a linux bridge with the `"name": "cni0" ` specified in the config file. For every pod, it then creates a veth pair - one end of the pair is in the container’s network namespace and the other end is connected to the linux bridge on the host network. With Bridge CNI plugin, all containers on a host are connected to the linux bridge on the host network.

什么时候
[Bridge CNI 插件](https://github.com/containernetworking/plugins/tree/master/plugins/main/bridge) 被第一次调用，它创建了一个名为 `"name": "cni0" 的 linux 桥` 在配置文件中指定。对于每个 pod，它然后创建一个 veth 对——该对的一端在容器的网络命名空间中，另一端连接到主机网络上的 linux 网桥。使用 Bridge CNI 插件，主机上的所有容器都连接到主机网络上的 linux 网桥。

After configuring the veth pair, Bridge plugin invokes the host-local IPAM CNI plugin. Which IPAM plugin to use can be configured in the CNI config CRI plugin uses to call the flannel CNI plugin.

配置 veth 对后，Bridge 插件调用主机本地 IPAM CNI 插件。使用哪个 IPAM 插件可以在 CNI 配置 CRI 插件用于调用 flannel CNI 插件中进行配置。

#### Host-local IPAM CNI plugins

#### 主机本地 IPAM CNI 插件

The Bridge CNI plugin calls the
[host-local IPAM CNI plugin](https://github.com/containernetworking/plugins/tree/master/plugins/ipam/host-local) with the following configuration:

Bridge CNI 插件调用
[host-local IPAM CNI plugin](https://github.com/containernetworking/plugins/tree/master/plugins/ipam/host-local) 配置如下：

```
{
"name": "cni0",
"ipam": {
    "type": "host-local",
    "subnet": "10.244.0.0/24",
    "dataDir": "/var/lib/cni/networks"
}
}

```


Host-local IPAM (IP Address Management) plugin returns an IP address for the container from the `subnet` and stores the allocated IP locally on the host under the directory specified under `dataDir` \- `/var/lib/cni/networks /<network-name=cni0>/<ip>`. `/var/lib/cni/networks/<network-name=cni0>/<ip>` file contains the container ID to which the IP is assigned.

主机本地IPAM（IP地址管理）插件从`subnet`返回容器的IP地址，并将分配的IP本地存储在主机上`dataDir`下指定的目录下\- `/var/lib/cni/networks /<网络名称=cni0>/<ip>`。 `/var/lib/cni/networks/<network-name=cni0>/<ip>` 文件包含分配了 IP 的容器 ID。

When invoked, the host-local IPAM plugin returns the following payload

调用时，主机本地 IPAM 插件返回以下有效负载

```
{
"ip4": {
    "ip": "10.244.4.2",
    "gateway": "10.244.4.3"
},
"dns": {}
}

```


## Summary

##  概括

Kube-controller-manager assigns a podCIDR to each node. Pods on a node are assigned an IP address from the subnet value in podCIDR. Because podCIDRs across all nodes are disjoint subnets, it allows assigning each pod a unique IP address.

Kube-controller-manager 为每个节点分配一个 podCIDR。节点上的 Pod 从 podCIDR 中的子网值中分配了一个 IP 地址。因为跨所有节点的 podCIDR 是不相交的子网，所以它允许为每个 pod 分配一个唯一的 IP 地址。

Kubernetes cluster administrator configures and installs kubelet, container runtime, network provider agent and distributes CNI plugins on each node. When network provider agent starts, it generates a CNI config. When a pod is scheduled on a node, kubelet calls the CRI plugin to create the pod. In containerd’s case, Containerd CRI plugin then calls the CNI plugin specified in the CNI config to configure the pod network. And all of this results in a pod getting an IP address.

Kubernetes 集群管理员配置和安装 kubelet、容器运行时、网络提供程序代理并在每个节点上分发 CNI 插件。当网络提供程序代理启动时，它会生成一个 CNI 配置。当一个 pod 被调度到一个节点上时，kubelet 会调用 CRI 插件来创建 pod。在 containerd 的情况下，Containerd CRI 插件然后调用 CNI 配置中指定的 CNI 插件来配置 pod 网络。所有这些都会导致 pod 获得 IP 地址。

* * *

* * *

It took me a while to understand all the interactions and the details involved. I hope this helped you in improving your understanding of how kubernetes works. If you think I got something wrong, please let me know via
[twitter](https://twitter.com/RonakNathani) or email me at 

我花了一段时间才了解所有的交互和所涉及的细节。我希望这有助于您提高对 kubernetes 工作原理的理解。如果您认为我做错了什么，请通过以下方式告诉我
[twitter](https://twitter.com/RonakNathani) 或给我发电子邮件

[hello@ronaknathani.com](mailto:hello@ronaknathani.com). If you’d like to discuss something in this post or anything else, feel free to reach out. I’d love to hear from you!

[hello@ronaknathani.com](mailto:hello@ronaknathani.com)。如果您想在这篇文章或其他任何内容中讨论某些内容，请随时与我们联系。我很想听听你的意见！

* * *

* * *

## References

##  参考

### Container Networking

### 容器网络

- [A container networking overview](https://jvns.ca/blog/2016/12/22/container-networking/)
- [Demystifying container networking](https://blog.mbrt.dev/2017-10-01-demystifying-container-networking/)

- [容器网络概述](https://jvns.ca/blog/2016/12/22/container-networking/)
- [揭秘容器网络](https://blog.mbrt.dev/2017-10-01-demystifying-container-networking/)

### How Flannel Works

### 法兰绒的工作原理

- [Flannel Networking Demistify](https://msazure.club/flannel-networking-demystify/)
- [Kubernetes With Flannel - Understanding The Networking](https://medium.com/@anilkreddyr/kubernetes-with-flannel-understanding-the-networking-part-2-78b53e5364c7)

- [Flannel Networking Demistify](https://msazure.club/flannel-networking-demystify/)
- [Kubernetes with Flannel - 理解网络](https://medium.com/@anilkreddyr/kubernetes-with-flannel-understanding-the-networking-part-2-78b53e5364c7)

### CRI and CNI

### CRI 和 CNI

- [CRI Plugin Architecture](https://github.com/containerd/cri/blob/v1.11.1/docs/architecture.md)
- [CNI Spec](https://github.com/containernetworking/cni/blob/master/SPEC.md)
- [CNI Plugins](https://github.com/containernetworking/plugins)

- [CRI 插件架构](https://github.com/containerd/cri/blob/v1.11.1/docs/architecture.md)
- [CNI 规范](https://github.com/containernetworking/cni/blob/master/SPEC.md)
- [CNI 插件](https://github.com/containernetworking/plugins)

[cni](http://ronaknathani.com/tag/cni/)[flannel](http://ronaknathani.com/tag/flannel/) [kubernetes](http://ronaknathani.com/tag/kubernetes/) [networking](http://ronaknathani.com/tag/networking/) 

[cni](http://ronaknathani.com/tag/cni/)[flannel](http://ronaknathani.com/tag/flannel/) [kubernetes](http://ronaknathani.com/tag/kubernetes/) [网络](http://ronaknathani.com/tag/networking/)

