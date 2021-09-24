# 8 Use Cases for Kubernetes over VPN: Unlocking Multicloud Flexibility

[Jun 29 2021](http://itnext.io/8-use-cases-for-kubernetes-over-vpn-unlocking-multicloud-flexibility-3958dab2219f?source=post_page-----3958dab2219f--------------------------------)·8 min read

This is a higher-level take on another story I wrote about [how to run Kubernetes across multiple clouds](http://itnext.io/how-to-deploy-a-single-kubernetes-cluster-across-multiple-clouds-using-k3s-and-wireguard-a5ae176a6e81). Before explaining _how_, I should have answered _why?_ So, let’s start with the problem.

# **What’s the Problem with Multicloud?**

Kubernetes has become the de facto cloud native computing standard in a short period of time. Released in 2014, there is a Kubernetes distribution for every platform, and most enterprises have a Kubernetes strategy.

This has introduced a set of problems that enterprises are still learning to solve. Chief among these problems is **inter-cloud operability**, or, the ability to manage clusters and cluster-based applications across a variety of environments.

Typical solution architectures focus on deploying multiple clusters into various environments, and then coordinating the management of both infrastructure and applications across these environments.

In many scenarios where environments are strictly segmented, such solutions are necessary. However, these architectures introduce problems of their own. They require additional complex software which must be learned and operated. They require an operations team to manage many clusters simultaneously. They also require substantial resource overhead, because at least one full cluster is required per environment.

# How can we solve this?

Enterprises typically overlook an alternative approach to these problems: the mesh VPN. This is not to be confused with a “service mesh,” an unrelated Kubernetes concept. The mesh VPN creates a “virtual” cloud environment for a cluster, which may be composed of many physical environments.

With a mesh VPN, enterprises can manage a **single cluster** across **multiple environments**. This method is relatively simple and has several key benefits:

- Cloud burst into new environments
- Save on resource overhead
- No new skills or tools required

A mesh VPN also provides an added benefit to Kubernetes clusters outside of time and savings: **security**.

A Kubernetes cluster enabled with a mesh VPN has encrypted traffic between all nodes, and enables new patterns for secure access from and to the cluster.

# What is a Mesh VPN?

![](https://miro.medium.com/max/900/0*URDLYJlaaebN6Mlu?q=20)

A Mesh VPN is a virtual private network where every computer has a direct connection to every other computer over a private IP address.

In the context of Kubernetes, it is a virtual subnet where you deploy your worker nodes, which does not require those nodes to be deployed in the same location.

This enables platform teams to deploy clusters with nodes placed in arbitrary environments. For instance, you can have an on-prem data center which can “cloud burst” into a public cloud environment if more resources are needed. From the cluster’s perspective, it is just one normal Kubernetes cluster. It has no idea its nodes are placed in different locations. Simple management is enabled via node selectors.

An enterprise could use any one of the available mesh VPNs, including Nebula, Tailscale, Twingate, Netmaker, and others. However, it is critically important to use a mesh VPN based on **kernel WireGuard**.

![](https://miro.medium.com/max/900/0*fQukEPouNwGQxYtI?q=20)

VPN’s have historically caused significant latency, and can lead to a 30% reduction in bandwidth or more. Additionally, they are usually complex and heavy. WireGuard, a breakthrough VPN technology, eliminates these problems. It achieves near-parity in speed to the same network without WireGuard, and has a very simple implementation relative to older technologies like OpenVPN and IPSec.

Only a few of the latest mesh VPN’s (such as Netmaker) utilize kernel WireGuard in order to maximize speed.

# What are the limitations?

Before going into the use cases, let’s cover a few limitations. There are reasons this is not a widely used pattern today, though they are due to some misconceptions which can be solved.

The largest misconception is that Kubernetes cannot run with high latency: it _can,_ but _etcd_ cannot. This means that your master nodes either must be co-located, or use a database other than etcd. For instance, SQL with K3S has no such problem. MicroK8s also runs with Dqlite, a distributed version of SQL which is tolerant to latency.

In addition, operators are used to the latency cost of running with traditional VPN’s like IPSec and OpenVPN. They may not know how fast WireGuard is.

There are three reasonable factors to consider that should dissuade some companies from running with this strategy: bandwidth pricing, corporate firewalls, and application-level latency.

If you are using a cloud provider with very high egress data charges, you may want to run an analysis of what those charges will look like, since your nodes will be sending data back and forth between clouds. Still, this price may be lower than the price of running duplicated infra in each environment. DigitalOcean’s bandwidth pricing is likely much lower than running extra infrastructure.

If you are in a corporate environment with heavy firewall restrictions between environments, requiring layers of permissions to run data between them, this may also create a challenge in setting up this topology. Still, it may be worth considering if you already need to run such integrations.

Finally, there are many cases where your applications cannot tolerate high latency. In these cases, you can solve the problem with simple affinity rules and node selectors: Per location, give nodes a label, and then set up affinity rules on your cluster so that pods in applications will schedule to one group or the other by default.

There are some cases where it makes more sense to run duplicated infrastructure to avoid dealing with potential inter-cloud issues, but there are less of these cases than you might think. Still, it is always important to consider the potential costs and hazards involved.

# What are the use cases?

Assuming a company is not heavily impacted by the above limitations, a mesh VPN enables many valuable use cases, which both decrease cost and complexity while enabling powerful deployment patterns. We discuss these different use cases below.

![](https://miro.medium.com/max/600/0*83CJvB0Hk0c0R_K2?q=20)

**Case 0: Regular Cluster**

If you are planning to deploy a cluster, it is always worth considering a WireGuard kernel-based mesh VPN underneath. You will have a normal functioning cluster with negligible performance differences. However, you enable your cluster to run with enhanced topologies in the future, with the added benefit of encrypted inter-node traffic.

**Case 1: Wide Nodes**

In this scenario, imagine a cluster from Case 0, but now you want to deploy an application that is integrated with an Azure service. You simply deploy a VM in Azure, add the VM to your mesh, and install the node. You now have the below topology. Migrating an application between the two clouds is as simple as changing the node selector.

![](https://miro.medium.com/max/900/0*gE3w6w2gvxGteEGV?q=20)

![](https://miro.medium.com/max/600/0*re8_QyXLE_YRNEzi?q=20)

**Case 2: Cloud Bursting**

In this scenario, an enterprise has an on-prem data center they are using for Kubernetes. They have deployed the mesh VPN underneath their cluster, which enables them to add nodes arbitrarily from a cloud environment (DigitalOcean, AWS, etc.). This can be very useful for cases where you need to scale your application quickly but have limited resources locally.

![](https://miro.medium.com/max/600/0*8Bp-71SJQDTCidHx?q=20)

**Case 3: Cloud Control**

In this scenario, an enterprise has substantial on-prem resources, but does not want to manage the control plane locally. Master node failure can be catastrophic and they’d rather use cloud resources. Because of this, the enterprise utilizes a cloud-based control plane but uses on-prem workers. This can bring down cloud costs substantially.

This scenario can also be applied to using **cloud-based storage nodes**, another critical component where failure is not an option.

**Case 4: Distributed/Edge Nodes**

In this scenario, worker nodes are all placed in different environments. This use case can also be applied to edge environments, where you have one central control plane pushing out applications (perhaps via daemonset) to all of the edge nodes.

![](https://miro.medium.com/max/900/0*UTVZCDwpWwaYP4c7?q=20)

**Case 5: Distributed Clusters**

This scenario is more complex and requires a deeper dive, but with the right architecture, the entire cluster may be run across arbitrary clouds, including the master nodes. The main limitation here is etcd, which does not tolerate high latency. Such a pattern may require a non-etcd cluster database.

**Case 6: Connected Clusters**

In this scenario, clusters are connected over the mesh VPN and can communicate with each other directly, meaning a microservice-based application can be split between the two clusters. This may be used in conjunction with multi-cluster management tools.

![](https://miro.medium.com/max/900/0*krXPHPaHhzMFTWyh?q=20)

**Case 7: Secure, private access (inbound)**

In this scenario, an enterprise only wants trusted users or applications to access cluster resources. Here, the mesh VPN acts similarly to a corporate VPN, where one must attach to the VPN in order to gain access. This can also be performed at a lower level to allow outside access to the service and pod networks within the Kubernetes cluster, a useful feature for development and operations teams.

![](https://miro.medium.com/max/900/0*ptp8f6Gi_3e_8rBF?q=20)

**Case 8: Secure, private access (outbound)**

In this scenario, a Kubernetes-based application needs secure access to a non-Kubernetes application such as a SQL database. The database VM is added to the VPN and the application now has secure, instant access (see above). This method can be used to add any arbitrary non-Kubernetes application to the Kubernetes network and allow pods to use the service.

# **Options for a Mesh VPN**

Several options exist for a mesh VPN today, including Tailscale, Netmaker, Kilo, Nebula, and others. Some of the above use cases become less feasible based on the choice of VPN. In addition, only Netmaker and Kilo offer kernel WireGuard as an option, a key consideration to minimize latency.

A company can also “roll their own” mesh using WireGuard directly, but this becomes very difficult to manage at scale and requires significant manual intervention.

GRAVITL has designed [Netmaker](https://gravitl.com/netmaker) to handle all of the above use cases while being largely automated and based on kernel WireGuard. In addition, since it is a generalized mesh VPN, it can be used for intergrating external services. That said, every company should make their own determination on which network virtualization tool is most suited to their environments.

# **Conclusion**

We have discussed the current state of multi-cloud Kubernetes clusters, the need for a mesh VPN within Kubernetes, and the different use cases it enables.

It is recommended that platform owners consider deploying a mesh VPN underneath their clusters, even if they do not have a current-state use case, as it does not significantly impact cluster performance, and allows for quick enablement of the above topologies when needed.

For more information about mesh VPN’s or [Netmaker](https://gravitl.com/netmaker), contact GRAVITL at [info@gravitl.com](mailto:info@gravitl.com), or visit our website at [https://gravitl.com/book](https://gravitl.com/book) .

[**ITNEXT**](https://itnext.io/?source=post_sidebar--------------------------post_sidebar-----------)

ITNEXT is a platform for IT developers & software engineers…