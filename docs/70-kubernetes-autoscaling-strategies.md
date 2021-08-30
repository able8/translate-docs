# Architecting Kubernetes clusters — choosing the best autoscaling strategy

Published in June 2021

![Architecting Kubernetes clusters — choosing the best autoscaling strategy](https://learnk8s.io/a/0498c77a1b646af0f3fc42a03f0171fb.svg)

*TL;DR: Scaling pods and nodes in a Kubernetes cluster could take several  minutes with the default settings. Learn how to size your cluster nodes, configure the Horizontal and Cluster Autoscaler, and overprovision your cluster for faster scaling.*

**Table of content:**

- [When autoscaling pods goes wrong](https://learnk8s.io/kubernetes-autoscaling-strategies#when-autoscaling-pods-goes-wrong)
- [How the Cluster Autoscaler works in Kubernetes](https://learnk8s.io/kubernetes-autoscaling-strategies#how-the-cluster-autoscaler-works-in-kubernetes)
- [Exploring pod autoscaling lead time](https://learnk8s.io/kubernetes-autoscaling-strategies#exploring-pod-autoscaling-lead-time)
- [Choosing the optimal instance size for a Kubernetes node](https://learnk8s.io/kubernetes-autoscaling-strategies#choosing-the-optimal-instance-size-for-a-kubernetes-node)
- [Overprovisioning nodes in your Kubernetes cluster](https://learnk8s.io/kubernetes-autoscaling-strategies#overprovisioning-nodes-in-your-kubernetes-cluster)
- [Selecting the correct memory and CPU requests for your Pods](https://learnk8s.io/kubernetes-autoscaling-strategies#selecting-the-correct-memory-and-cpu-requests-for-your-pods)
- [What about downscaling a cluster?](https://learnk8s.io/kubernetes-autoscaling-strategies#what-about-downscaling-a-cluster-)
- [Why not autoscaling based on memory or CPU?](https://learnk8s.io/kubernetes-autoscaling-strategies#why-not-autoscaling-based-on-memory-or-cpu-)

In Kubernetes, several things are referred to as "autoscaling", including:

- [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/).
- [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler).
- [Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).

Those autoscalers belong to different categories because they address other concerns.

The **Horizontal Pod Autoscaler (HPA)** is designed to increase the replicas in your deployments.

As your application receives more traffic, you could have the autoscaler  adjusting the number of replicas to handle more requests.

- ![The Horizontal Pod Autoscaler (HPA) inspects metrics such as memory and CPU at a regular interval.](https://learnk8s.io/a/b39ab8cff1b9a9f443aa9ef798c9dffb.svg)

  1/2

  The Horizontal Pod Autoscaler (HPA) inspects metrics such as memory and CPU at a regular interval.

  Next 

The **Vertical Pod Autoscaler (VPA)** is useful when you can't create more copies of your Pods, but you still need to handle more traffic.

As an example, you can't scale a database (easily) only by adding more Pods.

A database might require sharding or configuring read-only replicas.

But you can make a database handle more connections by increasing the memory and CPU available to it.

That's precisely the purpose of the vertical autoscaler — increasing the size of the Pod.

- ![You can't only increase the number of replicas to scale a database in Kubernetes.](https://learnk8s.io/a/41e2f5317f5dc17ddd0fcb9846ceeae1.svg)

  1/2

  You can't only increase the number of replicas to scale a database in Kubernetes.

  Next 

Lastly, the **Cluster Autoscaler (CA)**.

When your cluster runs low on resources, the Cluster Autoscaler provision a new compute unit and adds it to the cluster.

If there are too many empty nodes, the cluster autoscaler will remove them to reduce costs.

- ![When you scale your pods in Kubernetes, you might run out of space.](https://learnk8s.io/a/7389f3827b5ba289475c2f03893b0d9e.svg)

  1/3

  When you scale your pods in Kubernetes, you might run out of space.

  Next 

While these components all "autoscale" something, they are entirely unrelated to each other.

They all address very different use cases and use other concepts and mechanisms.

And they are developed in separate projects and can be used independently from each other.

**However, scaling your cluster requires fine-tuning the setting of the autoscalers so that they work in concert.**

Let's have a look at an example.

## When autoscaling pods goes wrong

Imagine having an application that requires and uses 1.5GB of memory and 0.25 vCPU at all times.

You provisioned a cluster with a single node of 8GB and 2 vCPU — it should  be able to fit four pods perfectly (and have a little bit of extra space left).

![A single node cluster with 8GB of memory and 2 vCPU](https://learnk8s.io/a/458f663479dc10d60ba31648fbb225fb.svg)

You deploy a single Pod and set up:

1. An **Horizontal Pod Autoscaler** adds a replica every 10 incoming requests (i.e. if you have 40 concurrent requests, it should scale to 4 replicas).
2. A **Cluster Autoscaler** to create more nodes when resources are low.

> The Horizontal Pod Autoscaler can scale the replicas in your deployment  using Custom Metrics such as the queries per second (QPS) from an  Ingress controller.

You start driving traffic 30 concurrent requests to your cluster and observe the following:

1. The **Horizontal Pod Autoscaler** starts scaling the Pods.
2. Two more Pods are created.
3. The **Cluster Autoscaler** doesn't trigger — no new node is created in the cluster.

It makes sense since there's enough space for one more Pod in that node.

![Scaling three replicas in a single node](https://learnk8s.io/a/2af2246616bdd048e4ef698bb4c9648c.svg)

You further increase the traffic to 40 concurrent requests and observe again:

1. The **Horizontal Pod Autoscaler** creates one more Pod.
2. The Pod is pending and cannot be deployed.
3. The **Cluster Autoscaler** triggers creating a new node.
4. The node is provisioned in 4 minutes. After that, the pending Pod is deployed.

- ![When you scale to four replicas, the fourth replicas isn't deployed in the first node. Instead, it stays "Pending".](https://learnk8s.io/a/1ffab2618c899dd9652e55ea8c973404.svg)

  1/2 When you scale to four replicas, the fourth replicas isn't deployed in the first node. Instead, it stays *"Pending"*.  Next 

*Why is the fourth Pod not deployed in the first node?*

Pods deployed in your Kubernetes cluster consume resources such as memory, CPU and storage.

However, on the same node, **the operating system and the kubelet require memory and CPU too.**

In a Kubernetes worker node's memory and CPU are divided into:

1. **Resources needed to run the operating system and system daemons** such as SSH, systemd, etc.
2. **Resources necessary to run Kubernetes agents** such as the Kubelet, the container runtime, [node problem detector](https://github.com/kubernetes/node-problem-detector), etc.
3. **Resources available to Pods.**
4. **Resources reserved to the [eviction threshold](https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/#eviction-thresholds)**.

![Resources in a Kubernetes cluster are consumed by Pods, the operating system, kubelet and eviction threshold](https://learnk8s.io/a/bbaadb833f5689978cfdd0d2fcd9b7ac.svg)

As you can guess, [all of those quotas are customisable](https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/#eviction-thresholds), but you need to account for them.

In an 8GB and 2 vCPU virtual machine, you can expect:

- 100MB of memory and 0.1 vCPU to be reserved for the operating system.
- 1.8GB of memory and 0.07 vCPU to be reserved for the Kubelet.
- 100MB of memory for the eviction threshold.

**The remaining ~6GB of memory and 1.83 vCPU are usable by the Pods.**

If your cluster runs a DeamonSet such as kube-proxy, you should further reduce the available memory and CPU.

Considering kube-proxy has requests of 128MB and 0.1 vCPU, only ~5.9GB of memory and 1.73 vCPU are available to run Pods.

Running a CNI like Flannel and a log collector such as Fluentd will further reduce your resource footprint.

After accounting for all the extra resources, you have space left for only three pods.

```
OS                  100MB, 0.1 vCPU   +
Kubelet             1.8GB, 0.07 vCPU  +
Eviction threshold  100MB, 0 vCPU     +
Daemonsets          128MB, 0.1 vCPU   +
======================================
Used                2.1GB, 0.27 vCPU

======================================
Available to Pods   5.9GB, 1.73 vCPU

Pod requests        1.5GB, 0.25 vCPU
======================================
Total (4 Pods)        6GB, 1vCPU
```

The fourth stays "Pending" unless it can be deployed on another node.

*Since the Cluster Autoscaler knows that there's no space for a fourth Pod, why doesn't it provision a new node?*

*Why does it wait for the Pod to be pending before it triggers creating a node?*

## How the Cluster Autoscaler works in Kubernetes

**The Cluster Autoscaler doesn't look at memory or CPU available when it triggers the autoscaling.**

Instead, the Cluster Autoscaler reacts to events and checks for any unschedulable Pods every 10 seconds.

A pod is unschedulable when the scheduler is unable to find a node that can accommodate it.

For example, when a Pod requests 1 vCPU but the cluster has only 0.5 vCPU  available, the scheduler marks the Pod as unschedulable.

**That's when the Cluster Autoscaler initiates creating a new node.**

The Cluster Autoscaler scans the current cluster and [checks if any of the unschedulable pods would fit on in a new node.](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders)

If you have a cluster with several node types (often also referred to as  node groups or node pools), the Cluster Autoscaler will pick one of them using the following strategies:

- **Random** — picks a node type at random. This is the default strategy.
- **Most pods** — selects the node group that would schedule the most pods.
- **Least waste** — selects the node group with the least idle CPU after scale-up.
- **Price** — select the node group that will cost the least (only works for GCP at the moment).
- **Priority** — selects the node group with the highest priority (and you manually assign priorities).

Once the node type is identified, the Cluster Autoscaler will call the relevant API to provision a new compute resource.

If you're using AWS, the Cluster Autoscaler will provision a new EC2 instance.

On Azure, it will create a new Virtual Machine and on GCP, a new Compute Engine.

It may take some time before the created nodes appear in Kubernetes.

Once the compute resource is ready, the [node is initialised](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet-tls-bootstrapping/) and added to the cluster where unscheduled Pods can be deployed.

**Unfortunately, provisioning new nodes is usually slow.**

It might take several minutes to provision a new compute unit.

*But let's dive into the numbers.*

## Exploring pod autoscaling lead time

The time it takes to create a new Pod on a new Node is determined by four major factors:

1. **Horizontal Pod Autoscaler reaction time.**
2. **Cluster Autoscaler reaction time.**
3. **Node provisioning time.**
4. **Pod creation time.**

By default, [pods' CPU and memory usage is scraped by kubelet every 10 seconds.](https://github.com/kubernetes/kubernetes/blob/2da8d1c18fb9406bd8bb9a51da58d5f8108cb8f7/pkg/kubelet/kubelet.go#L1855)

[Every minute, the Metrics Server will aggregate those metrics](https://github.com/kubernetes-sigs/metrics-server/blob/master/FAQ.md#how-often-metrics-are-scraped) and expose them to the rest of the Kubernetes API.

The Horizontal Pod Autoscaler controller is in charge of checking the metrics and deciding to scale up or down your replicas.

By default, the [Horizontal Pod Autoscaler checks Pods metrics every 15 seconds.](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#how-does-the-horizontal-pod-autoscaler-work)

![The Horizontal Pod Autoscaler can take up to 1 minute and a half to trigger the autoscaling](https://learnk8s.io/a/f147f5a02119f3320d03c689397f3b48.svg)

[The Cluster Autoscaler checks for unschedulable Pods in the cluster every 10 seconds.](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#how-does-scale-up-work)

Once one or more Pods are detected, it will run an algorithm to decide:

1. **How many nodes** are necessary to deploy all pending Pods.
2. **What type of node group** should be created.

The entire process should take:

- **No more than 30 seconds** on clusters with **less than 100 nodes** with up to 30 pods each. The average latency should be about 5 seconds.
- **No more than 60 seconds** on cluster with **100 to 1000 nodes**. The average latency should be about 15 seconds.

![The Cluster Autoscaler takes 30 seconds to decide to create a small node](https://learnk8s.io/a/54ed073d5b8e6c852ba0a7cb19dedec0.svg)

Then, there's the Node provisioning time, which depends mainly on the cloud provider.

**It's pretty standard for a new compute resource to be provisioned in 3 to 5 minutes.**

![Creating a virtual machine on a cloud provider could take several minutes](https://learnk8s.io/a/fdaa1ad973d840c63de2430ae6efdab6.svg)

Lastly, the Pod has to be created by the container runtime.

Launching a container shouldn't take more than few milliseconds, but **downloading the container image could take several seconds.**

If you're not caching your container images, downloading an image from the container registry could take from a couple of seconds up to a minute,  depending on the size and number of layers.

![Downloading a container image could take time and affect scaling](https://learnk8s.io/a/8c0b05d9016ee14a1a6f9952c9d3aa8a.svg)

So the total timing for trigger the autoscaling when there is no space in the current cluster is:

1. The Horizontal Pod Autoscaler might take up to 1m30s to increase the number of replicas.
2. The Cluster Autoscaler should take less than 30 seconds for a cluster with  less than 100 nodes and less than a minute for a cluster with more than  100 nodes.
3. The cloud provider might take 3 to 5 minutes to create the computer resource.
4. The container runtime could take up to 30 seconds to download the container image.

In the worse case, with a small cluster, you have:

```
HPA delay:          1m30s +
CA delay:           0m30s +
Cloud provider:     4m    +
Container runtime:  0m30s +
=========================
Total               6m30s
```

With a cluster with more than 100 nodes, the total delay could be up to 7 minutes.

*Are you happy to wait for 7 minutes before you have more Pods to handle a sudden surge in traffic?*

*How can you tune the autoscaling to reduce the 7 minutes scaling time if you need a new node?*

You could change:

- The refresh time for the Horizontal Pod Autoscaler (controlled by the `--horizontal-pod-autoscaler-sync-period` flag, default is 15 seconds).
- The interval for metrics scraping in the Metrics Server (controlled by the `metric-resolution` flag, default 60 seconds).
- How frequently the cluster autoscaler scans for unscheduled Pods (controlled by the `scan-interval` flag, default 10 seconds).
- How you cache the image on the local node ([with a tool such as kube-fledged](https://github.com/senthilrch/kube-fledged)).

But even if you were to tune those settings to a tiny number, you will still be limited by the cloud provider provisioning time.

*So, how could you fix that?*

Since you can't change the provisioning time, you will need a workaround this time.

You could try two things:

1. **Avoid creating new nodes,** if possible.
2. **Creating nodes proactively** so that they are already provisioned when you need them.

*Let's have a look at the options one at a time.*

## Choosing the optimal instance size for a Kubernetes node

**Choosing the right instance type for your cluster has dramatic consequences on your scaling strategy.**

*Consider the following scenario.*

You have an application that requests 1GB of memory and 0.1 vCPU.

You provision a node that has 4GB of memory and 1 vCPU.

After reserving memory and CPU for the kubelet, operating system and eviction threshold, you are left with ~2.5GB of memory and 0.7 vCPU that can be  used for running Pods.

![Choosing a smaller instance can affect scaling](https://learnk8s.io/a/22fe8bec231d15671512cf797ff3d1cc.svg)

Your node has space for only two Pods.

Every time you scale your replicas, you are likely to incur in up to 7  minutes delay (the lead time to trigger the Horizontal Pod Autoscaler,  Cluster Autoscaler and provisioning the compute resource on the cloud  provider).

*Let's have a look at what happens if you decide to use a 64GB memory and 16 vCPU node instead.*

After reserving memory and CPU for the kubelet, operating system and eviction threshold, you are left with ~58.32GB of memory and 15.8 vCPU that can  be used for running Pods.

**The available space can host 58 Pods, and you are likely to need a new node only when you have more than 58 replicas.**

![Choosing a larger instance can affect scaling](https://learnk8s.io/a/691e6c33c3d1db3bb5906512552c5bac.svg)

Also, every time a node is added to the cluster, several pods can be deployed.

There is less chance to trigger *again* the Cluster Autoscaler (and provisioning new compute units on the cloud provider).

Choosing large instance types also has another benefit.

**The ratio between resource reserved for Kubelet, operating system and  eviction threshold and available resources to run Pods is greater.**

Have a look at this graph that pictures the memory available to pods.

![Memory available to pods based on different instance types](https://learnk8s.io/a/6829f5509ca3044e1d3166a1f6218eac.svg)

As the instance size increase, you can notice that (in proportion) the resources available to pods increase.

In other words, you are utilising your resources more efficiently than having two instances of half of the size.

*Should you select the biggest instance all the time?*

**There's a peak in efficiency dictated by how many Pods you can have on the node.**

Some cloud providers limit the number of Pods to 110 (i.e. GKE). Others have limits dictated by the underlying network on a per-instance basis (i.e. AWS).

> [You can inspect the limits from most cloud providers here.](https://docs.google.com/spreadsheets/d/1RPpyDOLFmcgxMCpABDzrsBYWpPYCIBuvAoUQLwOGoQw/edit#gid=907731238)

**And choosing a larger instance type is not always a good option.**

You should also consider:

1. **Blast radius** — if you have only a few nodes, then the impact of a failing node is bigger than if you have many nodes.
2. **Autoscaling is less cost-effective** as the next increment is a (very) large node.

Assuming you have selected the right instance type for your cluster, you might  still face a delay in provisioning the new compute unit.

*How can you work around that?*

*What if instead of creating a new node when it's time to scale, you create the same node ahead of time?*

## Overprovisioning nodes in your Kubernetes cluster

If you can afford to have a spare node available at all times, you could:

1. Create a node and leave it empty.
2. As soon as there's a Pod in the empty node, you create another empty node.

**In other words, you teach the autoscaler always to have a spare empty node if you need to scale.**

- ![When you decide to overprovision a cluster, a node is always empty and ready to deploy Pods.](https://learnk8s.io/a/dfe3abd74845075ae2f9fc9b4cf896c3.svg)

  1/3 When you decide to overprovision a cluster, a node is always empty and ready to deploy Pods.  Next 

**It's a trade-off: you incur an extra cost (one empty compute unit available at all times), but you gain in speed.**

With this strategy, you can scale your fleet much quicker.

*But there's bad and good news.*

The bad news is that the Cluster Autoscaler doesn't have this functionality built-in.

**It cannot be configured to be proactive, and there is no flag to "always provision an empty node".**

The good news is that you can still fake it.

*Let me explain.*

**You could run a Deployment with enough requests to reserve an entire node.**

You could think about this pod as a placeholder — it is meant to reserve space, not use any resource.

As soon as a real Pod is created, you could evict the placeholder and deploy the Pod.

- ![In an overprovisioned cluster you have a Pod as a placeholder with low priority.](https://learnk8s.io/a/5297ee233777a99ef552f288299ffb80.svg)

  1/3

  In an overprovisioned cluster you have a Pod as a placeholder with low priority.

  Next 

Notice how this time, you still have to wait 5 minutes for the node to be  added to the cluster, but you can keep using the current node.

In the meantime, a new node is provisioned in the background.

*How can you achieve that?*

**Overprovisioning can be configured using deployment running a pod that sleeps forever.**

overprovision.yaml

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: overprovisioning
spec:
  replicas: 1
  selector:
    matchLabels:
      run: overprovisioning
  template:
    metadata:
      labels:
        run: overprovisioning
    spec:
      containers:
        - name: pause
          image: k8s.gcr.io/pause
          resources:
            requests:
              cpu: '1739m'
              memory: '5.9G'
```

**You should pay extra attention to the memory and CPU requests.**

The scheduler uses those values to decide where to deploy a Pod.

In this particular case, they are used to reserve the space.

You could provision a single large pod that has roughly the requests matching the available node resources.

**Please make sure that you account for resources consumed by the kubelet, operating system, kube-proxy, etc.**

If your node instance is 2 vCPU and 8GB of memory and the available space  for pods is 1.73 vCPU and ~5.9GB of memory, your pause pod should match  the latter.

![Sizing sleep pod in overprovisioned clusters](https://learnk8s.io/a/2866bf31d1616ce105778a5f4a5820a9.svg)

To make sure that the Pod is evicted as soon as a real Pod is created, you can use [Priorities and Preemptions.](https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/)

**Pod Priority indicates the importance of a Pod relative to other Pods.**

When a Pod cannot be scheduled, the scheduler tries to preempt (evict) lower priority Pods to schedule the Pending pod.

You can configure Pod Priorities in your cluster with a PodPriorityClass:

priority.yaml

```
apiVersion: scheduling.k8s.io/v1beta1
kind: PriorityClass
metadata:
  name: overprovisioning
value: -1
globalDefault: false
description: 'Priority class used by overprovisioning.'
```

Since the default priority for a Pod is `0` and the `overprovisioning` PriorityClass has a value of `-1`, those Pods are the first to be evicted when the cluster runs out of space.

PriorityClass also has two optional fields: `globalDefault` and `description`.

- The `description` is a human-readable memo of what the PriorityClass is about.
- The `globalDefault` field indicates that the value of this PriorityClass should be used for Pods without a `priorityClassName`. Only one PriorityClass with `globalDefault` set to `true` can exist in the system.

You can assign the priority to your sleep Pod with:

overprovision.yaml

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: overprovisioning
spec:
  replicas: 1
  selector:
    matchLabels:
      run: overprovisioning
  template:
    metadata:
      labels:
        run: overprovisioning
    spec:
      priorityClassName: overprovisioning
      containers:
        - name: reserve-resources
          image: k8s.gcr.io/pause
          resources:
            requests:
              cpu: '1739m'
              memory: '5.9G'
```

*The setup is complete!*

When there are not enough resources in the cluster, the pause pod is preempted, and new pods take their place.

Since the pause pod become unschedulable, it forces the Cluster Autoscaler to add more nodes to the cluster.

*Now that you're ready to overprovision your cluster, it's worth having a look at optimising your applications for scaling.*

## Selecting the correct memory and CPU requests for your Pods

**The cluster autoscaler makes scaling decisions based on the presence of pending pods.**

The Kubernetes scheduler assigns (or not) a Pod to a Node based on its memory and CPU requests.

Hence, it's essential to set the correct requests on your workloads, or you  might be triggering your autoscaler too late (or too early).

*Let's have a look at an example.*

You decide to profile an application, and you found out that:

- Under average load, the application consumes 512MB of memory and 0.25 vCPU.
- At peak, the application should consume up to 4GB of memory and 1 vCPU.

![Setting the right memory and CPU requests](https://learnk8s.io/a/d44981930b1e0d843e696dd67f3b3b43.svg)

The limit for your container is 4GB of memory and 1 vCPU.

*However, what about the requests?*

The scheduler uses the Pod's memory and CPU requests to select the best node before creating the Pod.

So you could:

1. **Set requests lower than the actual average usage.**
2. Be conservative and **assign requests closer to the limit.**
3. **Set requests to match the actual limits.**

- ![You could assign requests that are lower than the average app consumption.](https://learnk8s.io/a/27a58b9c37d2e1ba3386f02c85c110ab.svg)

  1/3

  You could assign requests that are **lower** than the average app consumption.

  Next 

**Defining requests lower than the actual usage is problematic since your nodes will be often overcommitted.**

As an example, you can assign 256MB of memory as a memory request.

The Kubernetes scheduler can fit twice as many Pods for each node.

However, Pods use twice as much memory in practice and start competing for  resources (CPU) and being evicted (not enough memory on the Node).

![Overcommitting nodes](https://learnk8s.io/a/22fdf4559d43e563b3b9b0472ea68969.svg)

**Overcommitting nodes can lead to excessive evictions, more work for the kubelet and a lot of rescheduling.**

*What happens if you set the request to the same value of the limit?*

You can set request and limits to the same values.

In Kubernetes, this is often referred to as [Guaranteed Quality of Service class](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/#qos-classes) and refers to the fact that it's improbable that the pod will be terminated and evicted.

The Kubernetes scheduler will reserve the entire CPU and memory for the Pod on the assigned node.

**Pods with Guaranteed Quality of Service are stable but also inefficient.**

If your app uses 512MB of memory on average, but you reserve 4GB for it, you have 3.5GB unused most of the time.

![Overcommitting nodes](https://learnk8s.io/a/3661626fe6a72a79770b9f8e2139e015.svg)

*Is it worth it?*

If you want extra stability, yes.

If you want efficiency, you might want to lower the requests and leave a gap between those and the limit.

This is often referred to as **Burstable Quality of Service class** and refers to the fact that the Pod baseline consumption can occasionally burst into using more memory and CPU.

When your requests match the app's actual usage, the scheduler will pack your pods in your nodes efficiently.

**Occasionally, the app might require more memory or CPU.**

1. If there are resources in the Node, the app will use them before returning to the baseline consumption.
2. If the node is low on resources, the pod will compete for resources (CPU), and the kubelet might try to evict the Pod (memory).

*Should you use Guaranteed or Burstable quality of Service?*

*It depends.*

1. **Use Guaranteed Quality of Service (requests equal to limits) when you want to minimise rescheduling and evictions for the Pod.** An excellent example is a Pod for a database.
2. **Use Burstable Quality of Service (requests to match actual average usage)  when you want to optimise your cluster and use the resources wisely.** If you have a web application or a REST API, you might want to use a Burstable Quality of Service.

*How do you select the correct requests and limits values?*

**You should profile the application and measure memory and CPU consumption when idle, under load and at peak.**

A more straightforward strategy consists of deploying the Vertical Pod Autoscaler and wait for it to suggest the correct values.

The Vertical Pod Autoscaler collects the data from the Pod and applies a regression model to extrapolate requests and limits.

[You can learn more about how to do that in this article.](https://learnk8s.io/setting-cpu-memory-limits-requests)

## What about downscaling a cluster?

**Every 10 seconds, the Cluster Autoscaler decides to remove a node only when the request utilization falls below 50%.**

In other words, for all the pods on the same node, it sums the CPU and memory requests.

If they are lower than half of the node's capacity, the Cluster Autoscaler will consider the current node for downscaling.

> It's worth noting that the Cluster Autoscaler does not consider actual CPU  and memory usage or limits and instead only looks at resource requests.

Before the node is removed, the Cluster Autoscaler executes:

- [Pods checks](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-types-of-pods-can-prevent-ca-from-removing-a-node) to make sure that the Pods can be moved to other nodes.
- [Nodes checks](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#i-have-a-couple-of-nodes-with-low-utilization-but-they-are-not-scaled-down-why) to prevent nodes from being destroyed prematurely.

If the checks pass, the Cluster Autoscaler will remove the node from the cluster.

## Why not autoscaling based on memory or CPU?

**CPU or memory-based cluster autoscalers don't care about pods when scaling up and down.**

Imagine having a cluster with a single node and setting up the autoscaler to  add a new node with the CPU reaches 80% of the total capacity.

You decide to create a Deployment with 3 replicas.

The combined resource usage for the three pods reaches 85% of the CPU.

A new node is provisioned.

*What if you don't need any more pods?*

You have a full node idling — not great.

**Usage of these type of autoscalers with Kubernetes is discouraged.**

## Summary

Defining and implementing a successful scaling strategy in Kubernetes requires you to master several subjects:

- Allocatable resources in Kubernetes nodes.
- Fine-tuning refresh intervals for Metrics Server, Horizontal Pod Autoscaler and Cluster Autoscalers.
- Architecting cluster and node instance sizes.
- Container image caching.
- Application benchmarking and profiling.

But with the proper monitoring tool, you can iteratively test your scaling  strategy and tune the speed and costs of your cluster.
