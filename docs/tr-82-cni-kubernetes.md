# A brief overview of the Container Network Interface (CNI) in Kubernetes

# Kubernetes 容器网络接口（CNI）的简要概述

Understand where the CNI fits into the Kubernetes architecture.

了解 CNI 在何处适合 Kubernetes 架构。

April 29, 2021 From: https://www.redhat.com/sysadmin/cni-kubernetes

If you have worked with [Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes?intcmp=701f20000012ngPAAQ) (K8s) and tried to learn some of its inner workings, either on the job or in a  training course, you must have learned a bit about [Container Network Interface](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/) (CNI). This article de-mystifies what CNI means and does.

如果您使用过 [Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes?intcmp=701f20000012ngPAAQ)(K8s) 并尝试了解其内部工作原理，无论是在在工作或培训课程中，你一定对 [容器网络接口](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/)（CNI)。本文揭开了 CNI 的含义和作用的神秘面纱。

## What is CNI?

## 什么是CNI？

A CNI plugin is responsible for inserting a network interface into  the container network namespace (eg, one end of a virtual ethernet  (veth) pair) and making any necessary changes on the host (eg,  attaching the other end of the veth into a bridge ). It then assigns an  IP address to the interface and sets up the routes consistent with the  IP Address Management section by invoking the appropriate IP Address  Management (IPAM) plugin.

CNI 插件负责将网络接口插入容器网络命名空间（例如，虚拟以太网（veth）对的一端）并在主机上进行任何必要的更改（例如，将 veth 的另一端连接到网桥）。然后它会为接口分配一个 IP 地址，并通过调用适当的 IP 地址管理 (IPAM) 插件来设置与 IP 地址管理部分一致的路由。

***[ You might also like: [A sysadmin's guide to basic Kubernetes components](https://www.redhat.com/sysadmin/kubernetes-components) ]***

***[您可能还喜欢：[系统管理员基本 Kubernetes 组件指南](https://www.redhat.com/sysadmin/kubernetes-components) ]***

## Where does CNI fit in?

## CNI 适合哪里？

CNI is used by container runtimes, such as Kubernetes (as shown below), as well as Podman, CRI-O, Mesos, and others.

CNI 被容器运行时使用，例如 Kubernetes（如下所示），以及 Podman、CRI-O、Mesos 等。

The container/pod initially has no network interface. The container runtime calls the CNI plugin with verbs such as **ADD**, **DEL**, **CHECK**, etc. ADD creates a new network interface for the container, and details of what is to be added are passed to CNI via JSON payload.

容器/pod 最初没有网络接口。容器运行时使用 **ADD**、**DEL**、**CHECK** 等动词调用 CNI 插件。 ADD 为容器创建一个新的网络接口，要添加的细节是通过 JSON 负载传递给 CNI。

## What does the CNI project consist of?

## CNI 项目由什么组成？

1. CNI specifications - Documents what the configuration format is  when you call the CNI plugin, what it should do with that information,  and the result that plugin should return.
2. Set of reference and example plugins - These can help you  understand how to write a new plugin or how existing plugins might work. They are cloud-agnostic. These are limited functionality plugins and  just for reference.



1. CNI 规范 - 记录调用 CNI 插件时的配置格式，它应该如何处理该信息，以及插件应该返回的结果。
2. 参考和示例插件集 - 这些可以帮助您了解如何编写新插件或现有插件如何工作。它们与云无关。这些是功能有限的插件，仅供参考。

The CNI specifications

CNI规范

- Vendor-neutral specifications
- Used by Mesos, CloudFoundry, Podman, CRI-O
- Defines basic execution flow and configuration format
- Attempt to keep things simple and backward compatible

- 供应商中立的规格
- 被 Mesos、CloudFoundry、Podman、CRI-O 使用
- 定义基本的执行流程和配置格式
- 尝试保持简单和向后兼容

## Execution flow of the CNI plugins

## CNI 插件的执行流程

1. When the container runtime expects to perform network operations on a container, it (like the kubelet in the case of K8s) calls the CNI  plugin with the desired command.
2. The container runtime also provides related network configuration and container-specific data to the plugin.
3. The CNI plugin performs the required operations and reports the result.



1. 当容器运行时期望在容器上执行网络操作时，它（如 K8s 中的 kubelet）使用所需命令调用 CNI 插件。
2. 容器运行时还为插件提供相关的网络配置和容器特定的数据。
3. CNI 插件执行所需的操作并报告结果。

CNI is called twice by K8s (kubelet) to set up **loopback** and **eth0** interfaces for a pod.

K8s (kubelet) 两次调用 CNI 来为 pod 设置 **loopback** 和 **eth0** 接口。

**Note**: CNI plugins are executable and support ADD, DEL, CHECK, VERSION commands, as discussed above.

**注意**：CNI 插件是可执行的，支持 ADD、DEL、CHECK、VERSION 命令，如上所述。

## Why are there multiple plugins?

## 为什么有多个插件？

CNI provides the specifications for various plugins. And as you know, networking is a complex topic with a variety of user needs. Hence,  there are multiple CNI plugins that do things differently to satisfy  various use cases.

CNI 提供了各种插件的规范。如您所知，网络是一个复杂的主题，具有多种用户需求。因此，有多个 CNI 插件以不同的方式执行不同的操作以满足各种用例。

## Wrap up

##  总结

There are many aspects of container orchestration and management with Kubernetes. You just learned a bit about the usage of CNI for networking within Kubernetes. For more, see the [CNI project page](https://github.com/containernetworking/cni) on GitHub. 

使用 Kubernetes 进行容器编排和管理有很多方面。您刚刚了解了在 Kubernetes 中使用 CNI 进行网络连接的一些知识。更多内容请查看 GitHub 上的 [CNI 项目页面](https://github.com/containernetworking/cni)。


