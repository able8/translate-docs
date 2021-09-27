## Why and How of Kubernetes Ingress (and Networking)

[Saaras Inc.](http://getenroute.io/author/saaras-inc.)August 07, 2021

![Why and How of Kubernetes Ingress (and Networking)](https://getenroute.io/img/EnrouteIngressDetail.jpeg)

Services running in Kubernetes are not accessible on public or private cloud. This is how Kubernetes is designed considering service security in mind.

Securely allowing access to a service outside the cluster requires some understanding of how networking is setup and the different requirements driving the networking choices.

We briefly start by exploring what is expected from a kubernetes cluster when it comes to service isolation, service scaling and service delivery. Once the high level requirements are laid out, it is easier to understand the significance of different constructs and abstractions.

We conclude by contrasting the advantages of using an Ingress to run a layer of L7 policy (or proxies) in front of a service running inside Kubernetes.

### Understanding the scheme of Kubernetes Networking

Understanding Kubernetes Ingress is key to running microservices and securely accessing those services. This article is an attempt to demystify how Kubernetes networking is setup.

We look at networking when a service is created, the different Kubernetes artifacts created, the networking machinery required to meet different requirements.

We also describe the significance of different types of IPs like External-IP, Node-IP, Cluster-IP, Pod-IP and describe how traffic passes through each one of them.

Starting with cluster networking requirements provides us an opportunity of why networking is setup the way it is.

#### Cluster Networking Requirements

**Cluster networking** in Kubernetes has several **requirements**

- **security and isolation of service(s)**
- **connection, networking and IP allocation for pods**
- **setup networking to build a cluster abstraction out of several physical nodes**
- **load balancing of traffic across multiple instances of a service**
- **controlling external access to a service**
- **working with Kubernetes networking in public and private cloud environments.**

To understand these different aspects of Kubernetes networking, we start by describing what happens when a service is created in a pod all the way to accessing that service in public and private cloud.

We highlight the need for Ingress and how it fits into the overall Kubernetes networking model.

#### Network isolation of Service running in Kubernetes Pods

Let us consider a simple Kubernetes cluster with two nodes

![](https://d33wubrfki0l68.cloudfront.net/9ba692de6b1b8123ba6657013707b7c33ad44229/35bb5/img/k8s-cluster-no-svc.png)

Kubernetes orchestrates containers or pods (which is a group of containers). **When Kubernetes creates a pod, it is run in its own isolated network (using network namespace).**

The diagram below shows two pods created on each node.

![](https://d33wubrfki0l68.cloudfront.net/9db66d843ff1651a60729bbb73357e8e8dc66837/66bf6/img/k8s-pod-ip.png)

What does this mean for a service? A service runs inside a pod in the pod’s network. **An IP address allocated on this pod network (for the service) isn’t accessible outside the pod.**

So how do you access this service?

#### Making Service in Pod accessible to host network stack

Kubernetes builds an abstraction of a cluster on top of multiple physical nodes or machines. The physical nodes have their own network stack. A pod created by Kubernetes creates an isolated network stack for the services that run inside the pod.

To reach this service (or IP address inside the pod), there needs to be routing/bridging that creates a path between the pod network and the host network. **Container Networking Interface or CNI sets up the networking associated with creating a traffic path between the node and pod.** Popular examples of CNI are calico, cilium, flannel, etc.

![](https://d33wubrfki0l68.cloudfront.net/c24a73d5ffe4bae6a90b11575322c7cffe5a1b75/6e56a/img/k8s-pod-ip-with-cni.png)

When Kubernetes creates a pod, it invokes the CNI callbacks. These callbacks result in the invocation of CNI provider services to setup IP addresses for the pod and connecting the pod network with the host network.

#### Making a Service accessible across Node boundaries

A service resides in a pod or several pods. Each of these pods can reside on one physical node or multiple physical nodes. As an example, say a service is spread across two pods that reside on two physical nodes.

When traffic is destined to this service (spread on two pods across two nodes), how does Kubernetes load balance traffic across them?

Kubernetes uses an abstraction of a Cluster IP. **Any traffic destined to Cluster IP is load-balanced across pods (in which the service runs)**.

To load balance to service instances in pods, networking is setup to reach the service in these pods. These pods may be running on different physical Nodes of the cluster. Wiring up a Cluster IP for a service ensures, that traffic sent to Cluster IP can be sent to all pods that run the service; regardless of the physical Node on which pods runs.

![](https://d33wubrfki0l68.cloudfront.net/e51cdd0b4cbafc371435491ad8bc092e147e2772/13c47/img/k8s-cluster-ip-service.png)

Implementation and realization of Cluster IP is achieved by kube-proxy component and mechanisms like iptables, ipvs or user space traffic steering.

#### Accessing a service from outside the cluster

Traffic destined to a ClusterIP is load-balanced across pods that may span multiple physical nodes. But the ClusterIP is only accessible from the nodes in the cluster. Or, put another way, **networking in Kubernetes ensures that external access to ClusterIP is restricted.**

Accessing the ClusterIP outside the cluster needs an explicit declaration to make it accessible outside the Nodes in a Kubernetes Cluster. This is `NodePort`

**A `NodePort` in Kubernetes wires up an Node IP (and port) with ClusterIP.**

Defining a `NodePort` provides an IP address on the local network. **Traffic sent to this NodePort IP (and port) is then routed to ClusterIP and eventually load balanced to the pods (and services).**

![](https://d33wubrfki0l68.cloudfront.net/b75cd368ad5928594d04f675c403d94958d4536a/c6dba/img/k8s-node-port.png)

#### Accessing service in Kubernetes on a public cloud

A `NodePort` makes a service accessible outside the cluster, but the IP address is only available locally. A `LoadBalancer` service is a way to associate a public IP (or DNS) with the `NodePort` service.

When a `LoadBalancer` type of service is created in a Kubernetes cluster, it allocates a public IP and sets up the load balancer on the cloud provider (like AWS, GCP, OCI, Azure etc.). **The cloud load balancer is configured to pipe traffic sent to the external IP to the `NodePort` service.**

![](https://d33wubrfki0l68.cloudfront.net/665c1caa1fb4dad7707ea9ab1ad0debe7b33e755/8b20c/img/k8s-external-ip-service.png)

#### Accessing service in Kubernetes on a private cloud

When running in a private cloud, creating a `LoadBalancer` type of service needs a Kubernetes controller that can provision a load balancer. One such implementation is [MetalLb](https://metallb.universe.tf/), which allocates an IP to route external traffic inside a cluster.

### Accessing Service on a Public Cloud with or without Ingress

There are a couple of ways to access a service running inside a Kubernetes Cluster on a public cloud. On a public cloud, when a service is of type `LoadBalancer` an External IP is allocated for external access.

- **A service can be directly declared as a `LoadBalancer` type.**

- **Alternatively, the Ingress service that controls and configures a proxy can be declared of type `LoadBalancer`. Routes and policies can then be created on this Ingress service, to route external traffic to the destination service.**


A proxy like Envoy/Nginx/HAProxy can receive all external traffic entering a cluster by running it as a service and defining this service of type `LoadBalancer`. These proxies can be configured using L7 routing and security rules. A collection of such rules forms the Ingress rules.

#### Without Ingress - Directly accessing a service by making it a service of type `LoadBalancer`

**When a service is declared as a type `LoadBalancer`, it directly receives traffic from the external load balancer.** In the diagram below, the service `helloenroute` service is declared of type `LoadBalancer`. It directly receives traffic from the external load balancer.

![](https://d33wubrfki0l68.cloudfront.net/81664a4db818bf4814a6e1e0fd6916a10632cd64/22e67/img/helloenroutedirectaccess.png)

#### With Ingress - Putting the Service behind a Proxy that is externally accessible through a `LoadBalancer`

A layer of L7 proxies can be placed before the service to apply L7 routing and policies. To achieve this, an Ingress Controller is needed.

**An Ingress Controller is a service inside a Kubernetes cluster that is configured as type `LoadBalancer` to receive external traffic. The Ingress Controller uses the defined L7 routing rules and L7 policy to route traffic to the service.**

In the example below, `helloenroute` service receives traffic from the EnRoute Ingress Controller which receives traffic from the external load balancer.

![](https://d33wubrfki0l68.cloudfront.net/72423798806754ee35c5981f16cd2e08e1a838b4/a21b1/img/enrouteingresslayer.png)

### Advantages of Using EnRoute Ingress Controller Proxy

There are several distinct advantages of running an Ingress Controller and enforcing policies at Ingress.

- Ingress provides a portable mechanism to enforce policy inside the Kubernetes Cluster. Policies enforced inside a cluster are easier to port across clouds.
- Multiple proxies can be horizontally scaled using Kubernetes service scaling. Elasticity of L7 fabric makes it easier to operate and scale it
- L7 policies can be hosted along with services inside the cluster with cluster-native state storage
- Keeping L7 policy closer to services simplifies policy enforcement and troubleshooting of services and APIs.

#### Plugins for fine-grained traffic control

EnRoute uses Envoy as the underlying proxy to provide the L7 ingress function. EnRoute has a modular architecture that closely reflects Envoy’s extensible model.

[Plugins/Filters](http://getenroute.io/features/) can be defined at the route level or service level to enforce L7 policies at Ingress. EnRoute provides advanced rate-limiting plugin in the community edition completely free without any limits. [Clocking your APIs and micro-services](http://getenroute.io/blog/why-every-api-needs-a-clock/) using deep L7 state is a critical need and EnRoutes flexible rate-limiting funtion provides immense flexibility to match a variety of rate-limiting use-cases.

[EnRoute Enterprise](http://getenroute.io/features/) includes support and enterprise plugins that help secure traffic at Kubernetes Ingress.

Tip: To quickly skim through the article, read through the bold text to get an overall understanding.
