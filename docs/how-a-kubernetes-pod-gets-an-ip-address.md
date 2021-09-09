# How a Kubernetes Pod Gets an IP Address

Aug 21, 20209 min read[Kubernetes](http://ronaknathani.com/category/kubernetes/)

One of the core requirements of the
[Kubernetes networking model](https://kubernetes.io/docs/concepts/cluster-administration/networking/#the-kubernetes-network-model) is that every pod should get its own IP address and that every pod in the cluster should be able to talk to it using this IP address. There are several network providers (flannel, calico, canal, etc.) that implement this networking model.

As I started working on Kubernetes, it wasn’t completely clear to me how every pod is assigned an IP address. I understood how various components worked independently, however, it wasn’t clear how these components fit together. For instance, I understood what CNI plugins were, however, I didn’t know how they were invoked. So, I wanted to write this post to share what I have learned about various networking components and how they are stitched together in a kubernetes cluster for every pod to receive an IP address.

There are various ways of setting up networking in kubernetes and various options for a container runtime. For this post, I will use
[Flannel](https://github.com/coreos/flannel) as the network provider and
[Containerd](https://github.com/containerd/containerd) as the container runtime. Also, I am going to assume that you know how container networking works and only share a very brief overview below for context.

## Some Background Concepts

### Container Networking: A Very Brief Overview

There are some really good posts explaining how container networking works. For context, I will go over a very high level overview here with a single approach that involves linux bridge networking and packet encapsulation. I am skipping details here as container networking deserves a blog post of itself. Some of the posts that I have found to be very educational in this space are
[linked in the references below](http://ronaknathani.com#container-networking).

#### Containers on the same host

One of the ways containers running on the same host can talk to each other via their IP addresses is through a linux bridge. In the kubernetes (and docker) world, a
[veth (virtual ethernet)](https://man7.org/linux/man-pages/man4/veth.4.html) device is created to achieve this. One end of this veth device is inserted into the container network namespace and the other end is connected to a
[linux bridge](https://wiki.archlinux.org/index.php/Network_bridge) on the host network. All containers on the same host have one end of this veth pair connected to the linux bridge and they can talk to each other using their IP addresses via the bridge. The linux bridge is also assigned an IP address and it acts as a gateway for egress traffic from pods destined to different nodes.
![bridge networking](http://ronaknathani.com/bridge-networking.png)

#### Containers on different hosts

One of the ways containers running on different hosts can talk to each other via their IP addresses is by using packet encapsulation. Flannel supports this through
[vxlan](https://vincent.bernat.ch/en/blog/2017-vxlan-linux) which wraps the original packet inside a UDP packet and sends it to the destination.

In a kubernetes cluster, flannel creates a vxlan device and some route table entries on each of the nodes. Every packet that’s destined for a container on a different host goes through the vxlan device and is encapsulated in a UDP packet. On the destination, the encapsulated packet is retrieved and the packet is routed through to the destined pod.
![flannel networking](http://ronaknathani.com/flannel-networking.png)

_NOTE: This is just one of the ways how networking between containers can be configured._

### What Is CRI?

[CRI (Container Runtime Interface)](https://github.com/kubernetes/cri-api) is a plugin interface that allows kubelet to use different container runtimes. Various container runtimes implement the CRI API and this allows users to use the container runtime of their choice in their kubernetes installation.

### What is CNI?

[CNI project](https://github.com/containernetworking/cni) includes a
[spec](https://github.com/containernetworking/cni/blob/master/SPEC.md) to provide a generic plugin-based networking solution for linux containers. It also consists of various
[plugins](https://github.com/containernetworking/plugins) which perform different functions in configuring the pod network. A CNI plugin is an executable that follows the CNI spec and we’ll discuss some plugins in the post below.

## Assigning Subnets To Nodes For Pod IP Addresses

If all pods are required to have an IP address, it’s important to ensure that all pods across the entire cluster have a unique IP address. This is achieved by assigning each node a unique subnet from which pods are assigned IP addresses on that node.

### Node IPAM Controller

When `nodeipam` is passed as an option to the
[kube-controller-manager’s](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/) `--controllers` command line flag, it allocates each node a dedicated subnet (podCIDR) from the cluster CIDR (IP range for the cluster network). Since these podCIDRs are disjoint subnets, it allows assigning each pod a unique IP address.

A kubernetes node is assigned a podCIDR when the node first registers with the cluster. To change the podCIDR allocated to nodes in a cluster, nodes need to be de-registered and then re-registered with any configuration changes first applied to the kubernetes control plane. `podCIDR` for a node can be listed using the following command.

```
$ kubectl get no <nodeName> -o json | jq '.spec.podCIDR'
10.244.0.0/24

```

## Kubelet, Container Runtime and CNI Plugins - how it’s all stitched together

When a pod is scheduled on a node, a lot of things happen to start up a pod. In this section, I’ll only focus on the interactions that relate to configuring network for the pod.

Once a pod is scheduled on the node, the following interactions result in configuring the network and starting the application container.
![kubelet-cri-cni-flowchart](http://ronaknathani.com/kubelet-cri-cni-flowchart.png)

Ref:
[Containerd cri plugin architecture](https://github.com/containerd/cri/blob/v1.11.1/docs/architecture.md)

## Interactions between Container Runtime and CNI Plugins

Every network provider has a CNI plugin which is invoked by the container runtime to configure network for a pod as it’s started. With containerd as the container runtime,
[Containerd CRI plugin](https://github.com/containerd/cri) invokes the CNI plugin. Every network provider also has an agent that’s installed on each of the kubernetes node to configure pod networking. When the network provider agent is installed, it either ships with the CNI config or it creates one on the node which is then used by the CRI plugin to figure out which CNI plugin to call.

The location for the CNI config file is configurable and the default value is `/etc/cni/net.d/<config-file>`. CNI plugins need to be shipped on every node by the cluster administrators. The location for CNI plugins is configurable as well and the default value is `/opt/cni/bin`.

In case of containerd as the container runtime, path for CNI configuration and CNI plugin binaries can be specified under `[plugins."io.containerd.grpc.v1.cri".cni]` section of the
[containerd config](https://github.com/containerd/cri/blob/master/docs/config.md).

Since we are referring to Flannel as the network provider here, I’ll talk a little about how Flannel is set up. Flanneld is the Flannel daemon and is typically installed on a kubernetes cluster as a daemonset with `install-cni` as an
[init container](https://github.com/coreos/flannel/blob/master/Documentation/kube-flannel.yml#L172). The `install-cni` container creates the
[CNI configuration file](https://gist.github.com/ronaknnathani/957a56210bd4fbd8e11120273c6b4ede) \- `/etc/cni/net.d/10-flannel.conflist` \- on each node. Flanneld creates a vxlan device, fetches networking metadata from the apiserver and watches for updates on pods. As pods are created, it distributes routes for all pods across the entire cluster and these routes allow pods to connect to each other via their IP addresses. For details on how flannel works, I recommend the
[linked references below](http://ronaknathani.com#how-flannel-works).

The interactions between Containerd CRI Plugin and CNI plugins can be visualized as follows:
![kubelet-cri-cni-interactions](http://ronaknathani.com/kubelet-cri-cni-interactions.png)

As described above, kubelet calls the Containerd CRI plugin in order to create a pod and Containerd CRI plugin calls the CNI plugin to configure network for the pod. The network provider CNI plugin calls other base CNI plugins to configure the network. The interactions between CNI plugins are described below.

### Interactions Between CNI Plugins

There are various CNI plugins that help configure networking between containers on a host. For this post, we will refer to 3 plugins.

#### Flannel CNI Plugin

When using Flannel as the network provider, the Containerd CRI plugin invokes the
[Flannel CNI plugin](https://github.com/containernetworking/plugins/tree/master/plugins/meta/flannel) using the CNI configuration file - `/etc/cni/net.d/10-flannel.conflist`.

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

```
FLANNEL_NETWORK=10.244.0.0/16
FLANNEL_SUBNET=10.244.0.1/24
FLANNEL_MTU=1450
FLANNEL_IPMASQ=false

```

The Flannel CNI plugin uses the information in `/run/flannel/subnet.env` to configure and invoke the bridge CNI plugin.

#### Bridge CNI Plugin

Flannel CNI plugin calls the Bridge CNI plugin with the following configuration:

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
[Bridge CNI plugin](https://github.com/containernetworking/plugins/tree/master/plugins/main/bridge) is invoked for the first time, it creates a linux bridge with the `"name": "cni0"` specified in the config file. For every pod, it then creates a veth pair - one end of the pair is in the container’s network namespace and the other end is connected to the linux bridge on the host network. With Bridge CNI plugin, all containers on a host are connected to the linux bridge on the host network.

After configuring the veth pair, Bridge plugin invokes the host-local IPAM CNI plugin. Which IPAM plugin to use can be configured in the CNI config CRI plugin uses to call the flannel CNI plugin.

#### Host-local IPAM CNI plugins

The Bridge CNI plugin calls the
[host-local IPAM CNI plugin](https://github.com/containernetworking/plugins/tree/master/plugins/ipam/host-local) with the following configuration:

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

Host-local IPAM (IP Address Management) plugin returns an IP address for the container from the `subnet` and stores the allocated IP locally on the host under the directory specified under `dataDir` \- `/var/lib/cni/networks/<network-name=cni0>/<ip>`. `/var/lib/cni/networks/<network-name=cni0>/<ip>` file contains the container ID to which the IP is assigned.

When invoked, the host-local IPAM plugin returns the following payload

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

Kube-controller-manager assigns a podCIDR to each node. Pods on a node are assigned an IP address from the subnet value in podCIDR. Because podCIDRs across all nodes are disjoint subnets, it allows assigning each pod a unique IP address.

Kubernetes cluster administrator configures and installs kubelet, container runtime, network provider agent and distributes CNI plugins on each node. When network provider agent starts, it generates a CNI config. When a pod is scheduled on a node, kubelet calls the CRI plugin to create the pod. In containerd’s case, Containerd CRI plugin then calls the CNI plugin specified in the CNI config to configure the pod network. And all of this results in a pod getting an IP address.

* * *

It took me a while to understand all the interactions and the details involved. I hope this helped you in improving your understanding of how kubernetes works. If you think I got something wrong, please let me know via
[twitter](https://twitter.com/RonakNathani) or email me at
[hello@ronaknathani.com](mailto:hello@ronaknathani.com). If you’d like to discuss something in this post or anything else, feel free to reach out. I’d love to hear from you!

* * *

## References

### Container Networking

- [A container networking overview](https://jvns.ca/blog/2016/12/22/container-networking/)
- [Demystifying container networking](https://blog.mbrt.dev/2017-10-01-demystifying-container-networking/)

### How Flannel Works

- [Flannel Networking Demistify](https://msazure.club/flannel-networking-demystify/)
- [Kubernetes With Flannel - Understanding The Networking](https://medium.com/@anilkreddyr/kubernetes-with-flannel-understanding-the-networking-part-2-78b53e5364c7)

### CRI and CNI

- [CRI Plugin Architecture](https://github.com/containerd/cri/blob/v1.11.1/docs/architecture.md)
- [CNI Spec](https://github.com/containernetworking/cni/blob/master/SPEC.md)
- [CNI Plugins](https://github.com/containernetworking/plugins)

[cni](http://ronaknathani.com/tag/cni/) [flannel](http://ronaknathani.com/tag/flannel/) [kubernetes](http://ronaknathani.com/tag/kubernetes/) [networking](http://ronaknathani.com/tag/networking/)
