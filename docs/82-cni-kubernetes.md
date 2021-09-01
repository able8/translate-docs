# A brief overview of the Container Network Interface (CNI) in Kubernetes     

Understand where the CNI fits into the Kubernetes architecture.

April 29, 2021 From: https://www.redhat.com/sysadmin/cni-kubernetes

If you have worked with [Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes?intcmp=701f20000012ngPAAQ) (K8s) and tried to learn some of its inner workings, either on the job or in a  training course, you must have learned a bit about [Container Network Interface](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/) (CNI). This article de-mystifies what CNI means and does.

## What is CNI?

A CNI plugin is responsible for inserting a network interface into  the container network namespace (e.g., one end of a virtual ethernet  (veth) pair) and making any necessary changes on the host (e.g.,  attaching the other end of the veth into a bridge). It then assigns an  IP address to the interface and sets up the routes consistent with the  IP Address Management section by invoking the appropriate IP Address  Management (IPAM) plugin.

***[ You might also like: [A sysadmin's guide to basic Kubernetes components](https://www.redhat.com/sysadmin/kubernetes-components) ]***

## Where does CNI fit in?

CNI is used by container runtimes, such as Kubernetes (as shown below), as well as Podman, CRI-O, Mesos, and others.

The container/pod initially has no network interface. The container runtime calls the CNI plugin with verbs such as **ADD**, **DEL**, **CHECK**, etc. ADD creates a new network interface for the container, and details of what is to be added are passed to CNI via JSON payload.

## What does the CNI project consist of?

1. CNI specifications - Documents what the configuration format is  when you call the CNI plugin, what it should do with that information,  and the result that plugin should return.
2. Set of reference and example plugins - These can help you  understand how to write a new plugin or how existing plugins might work. They are cloud-agnostic. These are limited functionality plugins and  just for reference.

The CNI specifications

- Vendor-neutral specifications
- Used by Mesos, CloudFoundry, Podman, CRI-O
- Defines basic execution flow and configuration format
- Attempt to keep things simple and backward compatible

## Execution flow of the CNI plugins

1. When the container runtime expects to perform network operations on a container, it (like the kubelet in the case of K8s) calls the CNI  plugin with the desired command.
2. The container runtime also provides related network configuration and container-specific data to the plugin.
3. The CNI plugin performs the required operations and reports the result.

CNI is called twice by K8s (kubelet) to set up **loopback** and **eth0** interfaces for a pod.

**Note**: CNI plugins are executable and support ADD, DEL, CHECK, VERSION commands, as discussed above.

## Why are there multiple plugins?

CNI provides the specifications for various plugins. And as you know, networking is a complex topic with a variety of user needs. Hence,  there are multiple CNI plugins that do things differently to satisfy  various use cases.

***[ Learn the basics of using Kubernetes in this [free cheat sheet](https://developers.redhat.com/promotions/kubernetes-cheatsheet?intcmp=701f20000012ngPAAQ). ]*** 

## Wrap up

There are many aspects of container orchestration and management with Kubernetes. You just learned a bit about the usage of CNI for  networking within Kubernetes. For more, see the [CNI project page](https://github.com/containernetworking/cni) on GitHub.