# Cloud Native Backups, Disaster Recovery and Migrations on Kubernetes

28 Jul 2020 Murat Karslioglu

Murat is an infrastructure architect with experience in storage, distributed systems and enterprise infrastructure development. He is VP of Product at MayaData, as well as one of the maintainers of the CNCF project _OpenEBS_ and author of _Kubernetes - A Complete DevOps Cookbook_.](https://www.linkedin.com/in/muratkarslioglu/)

Data is increasingly being run on Kubernetes, and increasingly in a way that supports cloud native architectures. In [a recent Cloud Native Computing Foundation webinar](https://www.cncf.io/webinars/kubernetes-for-storage-an-overview/), [Kiran Mova](https://www.linkedin.com/in/kiranmova/), the leader of the CNCF OpenEBS project and co-founder of MayaData, explained that loosely coupled architectures enable loosely coupled teams. The benefit being that loosely-coupled teams are unblocked; they’re able to iterate as and when they can, free from the scourge of meetings, long release cycles, change control boards and more.

One implication of this loosely-coupled, per workload approach — called **Container Attached Storage** — is that backups have to change as well. They need to change because the governance model is changing (small teams are in charge now) and because the fundamental architecture of loosely coupled workloads is different.

So what changes when data is deployed as an enormous number of loosely coupled workloads, as opposed to as a few centrally managed databases?

## **Challenges to Traditional Backups**

Traditionally backups could be performed bottom-up, for example backing up all the data on a particular volume, Kubernetes node, or virtual machine.

But what happens when instead of one or two workloads on a server or an ESX host, you have 110 or more pods; and each of these pods runs a workload, which is being moved by Kubernetes to other nodes as needed?

The answer is that if you can only backup at the cluster level, then your granularity is only that of a cluster. However, the small teams running workloads in your environment don’t manage clusters; they manage workloads. So in the event that they need their workload restored, they have to ask someone else to recover for them an entire Kubernetes cluster. And then they need to sort through this cluster for the particular data and application that they need. This breaks the control and autonomy of these small teams, which slows everyone down. It is a shared dependency and it puts your data at risk.

**Sponsor Note**

![sponsor logo](https://cdn.thenewstack.io/media/2020/07/2769a255-mayadata@2x.png)

MayaData delivers data agility. MayaData sponsors two Cloud Native Computing Foundation (CNCF) projects, OpenEBS – the leading open-source container attached storage solution – and Litmus – the leading Kubernetes native chaos engineering project. Well-known users of MayaData products include the CNCF itself, Bloomberg, Comcast, Arista, Orange, Intuit, and others.

## **Cloud Native Per Workload Backups, Disaster Recovery and Migration**

Contrast this to the experience of small teams running their workloads on a Container Attached Storage enabled Kubernetes environment, with data protection included via Kubera from MayaData. In this case, each small team remains in control of their workloads. If they feel that their workload needs to be cloned because they want to develop against a read-only copy of their production data, they can do that (assuming Kubernetes RBAC controls allows it). And if this small team feels that their particular use of PostgreSQL suggests they should have a cross AZ deployment for the primary DB cluster and also streaming or near real-time backups to another cloud, they can do that as well. They can also author a common pattern for production PostgreSQL for their environment and have that storage class become the default for all their deploys; and even share that pattern with others in their organization.

The autonomy of small teams to make decisions and iterate quickly is fundamental to agility. Loosely coupled architectures also require the separation of the state of workloads from infrastructure. Your “per workload backup” should be able to restore or migrate workloads to a different cloud platform, even though that other cloud platform has a different storage platform – otherwise, you are tightly coupled to the underlying platform.

Additionally, in many organizations, the platform teams need to be able to see the overall behavior of the platform. In our experience, some of the most important responsibilities of the platform teams include the protection of important data, capacity planning as data usage grows, and the management of one or more underlying cloud or hardware environments. Kubera gives these platform teams the tools they need to manage the resilience, capacity, and even the cost and performance of the overall environment.

## **Learn More**

Kubera from MayaData is available with a free forever tier. All you need to have running is a Kubernetes cluster — Kubera will connect to it and give you preconfigured off cluster logging and alerting for your stateful workloads, reporting and visualization, storage provisioning and more — in addition to the per workload backups, disaster recovery and migration mentioned in this article.

Kubera now includes integration with Cloudian Hyperstore at no additional costs — and Kubera starts at $49 per user per month. [Try it now — we are looking forward to your feedback](https://account.mayadata.io/signup).

_We will be discussing and demonstrating the new requirements of per workload Kubernetes backups, as well as data management and protection, backup, recovery and cloud migration, on August 6 with our partner, the leader in enterprise-class object storage, Cloudian._ _[Register now to learn more](https://go.mayadata.io/data-protection-for-kubernetes-webinar)._

Feature image via [Pixabay](https://pixabay.com/photos/squirrel-feeding-nuts-nager-garden-4382005/).

_At this time, The New Stack does not allow comments directly on this website. We invite all readers who wish to discuss a story to visit us on [Twitter](https://twitter.com/thenewstack) or [Facebook](https://www.facebook.com/thenewstack/). We also welcome your news tips and feedback via email: [feedback@thenewstack.io](mailto:feedback@thenewstack.io)._

[Contributed](https://thenewstack.io/tag/contributed/) [Sponsored](https://thenewstack.io/tag/sponsored/)

![](https://cdn.thenewstack.io/static/img/The-New-Stack-Updates-Logo.svg)
