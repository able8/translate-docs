# Kubernetes Performance Trouble Spots: Airbnb’sTake

#### 7 Jan 2020

Now that organizations are starting to rely on Kubernetes and containers in general, performance becomes a major focus point for admins, particularly for public-facing high-use services, such as [Airbnb](https://www.airbnb.com/). Engineers from the company shared some lessons learned on this topic at KubeCon+CloudNativeCon North America 2019.

In their talk, “ [Did Kubernetes Make My p95s Worse?](https://www.youtube.com/watch?v=QXApVwRBeys&t=3s)” Airbnb software engineers [Stephen Chan](https://www.linkedin.com/in/stephenyehengchan/), who is at the company’s compute infrastructure team, and [Jian Cheung](https://www.linkedin.com/in/jian-cheung/), who works in the service orchestration teams, discussed the performance gotcha they’ve witnessed working with the open source container orchestration engine.

Since 2018, the online housing marketplace has been the process of moving its services residing directly on AWS EC2 instance to its own Kubernetes-managed containers, presently about 1,000 services in all. As a result, Airbnb developers are quick to ask service orchestration team, “Why is my pod so slow?” The company runs [Amazon Linux 2](https://aws.amazon.com/amazon-linux-2/) for minion instances, Ubuntu images, the Flannel/Calico integration [Canal](https://github.com/projectcalico/canal) for the container networking, and K8s [NodePort](https://kubernetes.io/docs/concepts/services-networking/service/) to interface with the company’s service discovery mechanism.

At the conference, the engineers shared some performance issues they’ve encountered, as well as potential solutions. Their overall message was clear: When dealing with a complex Kubernetes-based infrastructure, performance tuning must be done all across the stack, including host, cluster, container, networking, and even with the underlying applications.

## Those Noisy Neighbors Again

Why do some pods have more latency than their peers in a cluster? One of the first culprits to check may be a neighboring pod, one that may be hogging all the CPU and networking resources for a heavy workload, the duo advised. Sometimes, the noise is purely accidental, when one hungry service meant to stay in staging was moved to a production cluster instead. But there are also control knobs that must be set as well: Airbnb made a choice early on not to enact CPU limits on its services, which limit the amount of resources that a service would take from its host CPU, this proved to be a bad idea, and the company has since set the resource limits from Kubernetes.

[![](https://cdn.thenewstack.io/media/2019/11/6809bec7-airbnb-noisu_neighbors.jpg)](https://static.sched.com/hosted_files/kccncna19/e1/%5Bkubecon%202019%5D%20Did%20Kubernetes%20make%20my%20p95s%20worse.pdf)

The “noisy neighbor problem” is not new to Kubernetes. It was first encountered, and mitigated against, when multiple virtual machines were first packed into servers, and a VM with a CPU-hungry app would hog all the resources, to the detriments of others.

Kubernetes has tools to prevent this from happening, though they can be tricky to use and can lead to what Cheung called “fine-grained hotspots” that are very difficult to pinpoint. Kubernetes uses the Linux kernel’s [CFS Bandwidth Control](https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt), which allots CPU time in microseconds to pre-defined groups. This can lead to throttling issues: A node can look slow, even when there is not a lot else happening on that CPU. If you set a CFS quota of 100 milliseconds of processor time for an application that requests 10 CPUs, it can use all 10 CPUs and burn up its quota within 20 milliseconds and will be throttled for the remaining 80 milliseconds, spiking the legacy levels for that app (The Linux kernel [has subsequently addressed this issue](https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt) with a patch).

“It’s hard not to take some performance hits from multitenancy. Before applications were running on their own dedicated boxes, but now they are sharing all their resources from other strange applications,” Cheung said.

The Kubernetes community has developed some fixes, including the ability to [make CFS quota periods configurable](https://github.com/kubernetes/kubernetes/pull/63437). There is a pull request that [disables CPU quotas](https://github.com/kubernetes/kubernetes/pull/75682) for pods requiring guaranteed quality of service (which you’d think wouldn’t be throttled [but they were](https://github.com/kubernetes/kubernetes/issues/70585)).

**Sponsor Note**

Portworx is the leading provider of persistent storage for containers and is used in production by healthcare, global manufacturing and telecom members of the Fortune Global 500 and other great companies. Learn about Portworx solutions for Kubernetes storage, DCOS storage and more at portworx.com.

Autoscaling can be another sneaky culprit of performance lag. At Airbnb, one service jumped from 600 pods to 1,000 due to heavy demand. To the scheduler, the  CPU utilization was fine, running at about 50% overall. At least one host, however, was crammed with 18 identical service pods by the scheduler. In other words, the overall CPU usage was fine, but some nodes were nonetheless starved from the lack of attention from the CPU.

The K8s scheduler has a set of rules about where to place images, based on a number of rules, such as spreading the workloads across as many nodes as possible, or abiding by a preferred or required affinity to a particular node. One rule, however, is that if an image was already downloaded to one minion node, the pods would more likely schedule to that node. In this case, one gigantic image was downloaded to one node, where all the other pods all piled up on that node.

“The scheduler can work against you in some pathological cases,” Chan said. The company will look at ways at limiting the number of pods that could run per node.

## Write Once Run Anywhere

The applications and underlying dependencies can also hamper performance. Take Java, for instance, Cheung said. One Airbnb development team noticed that a Java application, jumped from a 30 millisecond response time to over 100 milliseconds in the 95th percentile latency (or taking place on 95% of the servers). This happened, however, only when interacting with a database through a driver.  The unusual thing was the application worked fine before it ran on Kubernetes. The culprit turned out to be how the Java Virtual Machine (JVM) handled multi-CPU nodes. A single JVM within a pod on a 36-node cluster would see 36 CPUs. Great. But put two more JVMs, each on its own pod, on that node, and they will _all_ see 36 CPUs. Naturally, bottlenecks would ensue.

The team found that the problem with earlier versions of Java, which were [not aware of containers](https://bugs.openjdk.java.net/browse/JDK-8146115), Cheung noted. Java would auto-tune itself by how many CPUs it thought it had, and in container environments, this could adversely affect how thread pools were handled.

The issue has since been fixed in Java 8u191+, though the lesson here is one to keep in mind: “Languages and apps can have deeper dependencies on the underlying systems that they run on,” Cheung said. Another lesson learned was that it is useful to have a baseline to compare Kubernetes performance against that of running without K8s, as this application did.

Cheung and Chan discussed some other potential trouble spots, such as load balance issues stemming from IPtables, and general slowness stemming from DNS (Domain Name Server) misconfigurations.

Overall, Kubernetes is only one component of a complex cloud native stack, they reminded the audience, and so user performance may vary.

“Set the expectation that small performance differences will happen,” Chan said.

[KubeCon+CloudNativeCon](https://www.cncf.io/kubecon-cloudnativecon-events/) is a sponsor of The New Stack.
