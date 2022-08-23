# Difference Between multi-cluster, multi-master, multi-tenant & federated Kubernetes

Last updated June 9, 2021 https://platform9.com/blog/difference-between-multi-cluster-multi-master-multi-tenant-federated-kubernetes

Just as there are many parts that comprise Kubernetes, there are multiple ways to go about [Kubernetes deployment](https://platform9.com/docs/deploy-kubernetes-the-ultimate-guide/). The best approach for your needs depends on your team’s technical expertise, your infrastructure availability (or lack thereof), your capital expenditure and ROI goals, and more.

There are multiple ways to scale a Kubernetes environment to meet the needs of your organization’s requirements and manage the infrastructure as it grows. While they have similar-sounding names, they have very different applications in the real world. Read on to see how multi-tenant, multi-cluster, multi-master, and Federation fit into the mix.

## Summary of Deployment Models

The deployment models for Kubernetes range from single-server acting as master and worker with a single tenant, all the way to multiple multi-node clusters across multiple data centers, with federation enabled for some applications.

In this article we will focus on the more common deployment models and how they are different, as opposed to how they are similar. Because they all use the same base product, they have most functionality in common.

Some key differences include single-server models which are typically optimized for the use of developers, and not meant for production. Whereas production systems recommend a minimum of multi-master (with an odd number of nodes) and often end up with multi-cluster, and Federation becomes a very real consideration when the deployments need to scale to thousands of nodes and across regions.

## Single-tenant

A simple deployment, often using a single server, is perfect for a developer or QA team member who is trying to ensure containers are being built properly, and all scripts are Kubernetes compatible. These single-server deployments use a mini distribution such as minikube, or a limited deployment of a full suite like Platform9 Managed Kubernetes Free Tier (PMKFT).

These single-server single-tenant deployments run everything, not only on the same host but in the same name space. This is NOT a good practice for production, as it exposes application, deployment, and pod-specific configuration items – from secrets, to resource items like storage volumes, to all containers on the node. In addition, the network is wide open, so there are no restrictions on container interactions, which is against both network and security best practices.

[![](https://platform9.com/wp-content/uploads/2020/01/Single-Server-Single-Tenant-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Single-Server-Single-Tenant.jpeg)

## Multi-tenant

Once you passed the basic stage of just having Kubernetes running – whether it was a mini distribution or a full Kubernetes install with only the basic default networking enabled – you have, no doubt, run into the issue where services can access whatever they want, with no control. This makes enforcing any kind of security model or application flow nearly impossible, as developers can simply elect to ignore it and call anything they want.

Multi-tenant is the term for when a program, like a Kubernetes cluster, is configured to allow multiple areas within its purview to operate in isolation. This is critical when any kind of data is used, especially if it falls under financial or privacy regulations. The additional benefits are that multiple development teams, QA staff, or other entities can work in parallel within the same cluster with their own quotas and security profiles, without having any negative impact on the other teams also working in the cluster.

Within Kubernetes, this is primarily accomplished at the network layer for deployed pods that can only see pods within their name space by default, and can be granted access via network policies to get access to other namespaces and the applications they are running.

Additional technology, called Service Meshes, can be added to environments that are multi-tenant. The purpose is to increase network resilience between pods, enforce authentication and authorization policies, and to increase compliance to defined network policies. Service Meshes can also be leveraged for inter-cluster communications, but that is beyond what most deployments currently use them for.

[![](https://platform9.com/wp-content/uploads/2020/01/Single-Server-Multi-tenant-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Single-Server-Multi-tenant.jpeg)

## Multi-master

When Kubernetes is going to be used to support any kind of real workload, it is best to start down the path of having clusters be multi-master. By having more than one master server, the Kubernetes cluster can continue to be administered even in the event that not only a container like the api server fails, but if an entire server crashes taking out instances of all the operators and other controllers that are deployed and required to maintain a steady state across all the worker nodes.

Special consideration needs to be paid to etcd, which is the most commonly-used datastore for Kubernetes clusters. As etcd requires a quorum to support read-write operations, it needs to have a minimum of three nodes to be highly available and recommends having an odd number that will increase based on the number of worker nodes in the cluster. By the time you get to about 100 worker nodes, you’ll be having five to seven etcd servers to handle the load. Some Kubernetes distributions have etcd share the same servers as the master control plane containers, which works great as long as the individual nodes are sized appropriately.

[![](https://platform9.com/wp-content/uploads/2020/01/Multi-Master-Multi-Tenant-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Multi-Master-Multi-Tenant.jpeg)

## Multi-cluster

Multi-cluster environments are typically configured as multi-master and multi-tenant, as well. If you are big enough to need multiple clusters, there are probably high-availability and isolation requirements driven by at least a few of the groups that are actively building and deploying containers within your hybrid cloud infrastructure.

A multi-cluster deployment is simply deploying multiple clusters across one or more data centers, ideally with a single management interface to simplify life; although, that is not always the case. Companies like Platform9 offer products with multi-cluster management capabilities that can work with infrastructure provided by all the biggest cloud providers, in addition to on-premises infrastructure like OpenStack and VMware.

Reasons you would get into a multi-cluster scenario is separation of development and production; or, since it is not recommended to have a Kubernetes cluster that spans more than a single datacenter, having one or more regions in use across single- or multiple-cloud providers.

[![](https://platform9.com/wp-content/uploads/2020/01/Multi-Cluster-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Multi-Cluster.jpeg)

## Federation

Kubernetes Federation is the cherry on top of a multi-cluster deployment as it allows parts of the configuration to be synchronized across specific namespaces in the clusters that are federated. Where this becomes invaluable is when you have the same application and services spread across multiple datacenters or cloud regions, for reasons like capacity and avoiding a single point of failure.

By having these applications managed through Federation, the CI/CD or another deployment process can issue a single Kubernetes Deployment. It can roll out an application update across a truly global infrastructure while coordinating between the member clusters to ensure the chosen deployment model is being adhered to regardless of what cloud provider or region it is located in. This could be a hard cutover to the new application, canary deployments, or even just a gradual cutover leveraging a service mesh so there are no interrupted user sessions.

[![](https://platform9.com/wp-content/uploads/2020/01/Federated-Clusters-300x279.jpeg)](https://platform9.com/wp-content/uploads/2020/01/Federated-Clusters.jpeg)

## Conclusion

Kubernetes has a deployment model to meet the needs of any organization regardless of the scale they operate at, or their individual requirements. Within many organizations, multiple different deployment models are often used to support different phases of an application’s lifecycle and the needs of different business units. From the CTO office just needing a wide open playground to try things, to the shopping cart requiring truly isolated services to be PCI compliant.

- [Author](http://platform9.com#abh_about)
- [Recent Posts](http://platform9.com#abh_posts)


Platform9 is the world’s #1 open distributed cloud service, offering the power of the public cloud on infrastructure of customers’ choice — powered by Kubernetes and cloud-native technologies.Public clouds are walled gardens, and DIY is difficult and time consuming. Platform9 offers a third option — an open and faster option — enabling a better way to go cloud-native. Platform9’s service powers 40K+ nodes across private, public, and edge clouds.Innovative enterprises like Juniper, Kingfisher Plc, Mavenir, Redfin, and Cloudera achieve 4x faster time-to-market, up to 90% reduction in operational costs, and 99.9% uptime. Platform9 is an inclusive, globally distributed company backed by leading investors.


Latest posts by Platform9 ( [see all](https://platform9.com/blog/author/platform9/))