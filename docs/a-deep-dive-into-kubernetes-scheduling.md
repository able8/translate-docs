# A Deep Dive into Kubernetes Scheduling

​                    February 8, 2021                

Kubernetes Scheduler is one of the core components of the Kubernetes control plane. It runs on the control plane, and its  default behavior assigns pods to nodes while balancing resource  utilization among them. When the pods are assigned to a new node, the  kubelet running on the node retrieves the pod definition from the  Kubernetes API. Then, the kubelet creates the resources and containers  according to the pod specification on the node. In other words, the  scheduler runs inside the control plane and distributes the workload to  the Kubernetes cluster. It’s possible that you’ve never checked  Kubernetes Scheduler’s logs or configuration parameters because, for the most part, the tool works well for the majority of development,  testing, and production cases.

This article will take a deep dive into Kubernetes Scheduler,  starting with an overview of scheduling in general and scheduling  eviction with affinity and taints. We’ll then discuss the scheduler’s  bottlenecks and the issues that you may run into in production. Finally, we’ll examine how to fine-tune the scheduler’s parameters to suit your  custom-tailored clusters.

## An Overview of Scheduling

Kubernetes scheduling is simply the process of assigning pods to the  matched nodes in a cluster. A scheduler watches for newly created pods  and finds the best node for their assignment. It chooses the optimal  node based on Kubernetes’ scheduling principles and your configuration  options.

The simplest configuration option is setting the **nodeName** field in `podspec` directly as follows:

```
apiVersion: v1

kind: Pod

metadata:

  name: nginx

spec:

  containers:

  - name: nginx

    image: nginx

  nodeName: node-0
```

The nginx pod above will run on node-01 by default. However, **nodeName** has many limitations that lead to non-functional pods, such as unknown node names in the cloud, out of resource nodes, and nodes with intermittent  network problems. For this reason, you should not use **nodeName** at any time other than during testing or development.

If you want to run your pods on a specific set of nodes, use **nodeSelector** to ensure that happens. You can define the **nodeSelector** field as a set of key-value pairs in ‘PodSpec’:

```
apiVersion: v1

kind: Pod

metadata:

  name: nginx

spec:

  containers:

  - name: nginx

    image: nginx

  nodeSelector:

    disktype: ssd
```

For the nginx pod above, Kubernetes Scheduler will find a node with  the disktype: ssd label. Of course, the node can have additional labels, and Kubernetes already populates common labels to the nodes  kubernetes.io/arch or kubernetes.io/os. You can check the complete list  of tags in the [Kubernetes reference documentation](https://kubernetes.io/docs/reference/kubernetes-api/labels-annotations-taints/).

The use of `nodeSelector` efficiently constrains pods to  run on nodes with specific labels. However, its use is only constrained  with labels and their values. There are two more comprehensive features  in Kubernetes to express more complicated scheduling requirements: **node affinity**, to mark pods to attract them to a set of nodes; and **taints and tolerations**, to mark nodes to repel a set of pods. These features are discussed below.

## **Node Affinity**

Node affinity is a set of constraints defined on pods that determine  which nodes are eligible for scheduling. It’s possible to define hard  and soft requirements for the pods’ node assignments using affinity  rules. For instance, you can configure a pod to run only the nodes with  GPUs and preferably with NVIDIA_TESLA_V100 for your deep learning  workload. The scheduler evaluates the rules and tries to find a suitable node within the defined constraints. Like `nodeSelectors`, node affinity rules work with the node labels; however, they are more powerful than `nodeSelectors`.

There are four affinity rules you can add to `podspec`:

- **requiredDuringSchedulingIgnoredDuringExecution**
- **requiredDuringSchedulingRequiredDuringExecution**
- **preferredDuringSchedulingIgnoredDuringExecution**
- **preferredDuringSchedulingRequiredDuringExecution**

 

These four rules consist of two criteria: required or preferred, and  two stages: Scheduling and Execution. Rules starting with required  describe hard requirements that must be met. Rules beginning with  preferred are soft requirements that will be enforced but not  guaranteed. The Scheduling stage refers to the first assignment of the  pod to the nodes. The Execution stage applies to situations where node  labels change after the scheduling assignment.

If a rule is stated as **IgnoredDuringExecution**, the scheduler will not check its validity after the first assignment. However, if the rule is specified with **RequiredDuringExecution**, the scheduler will always ensure the rule’s validity by moving the pod to a suitable node.

Check out the following example to help you grasp these affinities:

```
apiVersion: v1
kind: Pod

metadata:
  name: nginx

spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: topology.kubernetes.io/region
            operator: In
            values:
            - us-east

      preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 1
        preference:
          matchExpressions:
          - key: topology.kubernetes.io/zone
            operator: In
            values:
            - us-east-1
            - us-east-2

  containers:
  - name: nginx
    image: nginx
```

The nginx pod above has a node affinity rule indicating that  Kubernetes Scheduler should only place the pod to a node in the us-east  region. The second rule indicates that the us-east-1 or us-east-2 zones  should be preferred.

Using affinity rules, you can make Kubernetes scheduling decisions work for your custom requirements.

## **Taints and Tolerations**

Not all Kubernetes nodes are the same in a cluster. It’s possible to  have nodes with special hardware, such as GPU, disk, or network  capabilities. Similarly, you may need to dedicate some nodes for  testing, data protection, or user groups. Taints can be added to the  nodes to repel pods, as in the following example:

```
kubectl taint nodes node1 test-environment=true:NoSchedule
```

With taint `test-environment=true:NoSchedule`, Kubernetes Scheduler will not assign any pod unless it has matching toleration in the`podspec`:

```
apiVersion: v1
kind: Pod

metadata:
  name: nginx

spec:
  containers:
  - name: nginx
    image: nginx

  tolerations:
  - key: "test-environment"
    operator: "Exists"
    effect: "NoSchedule"
```

Taints and tolerations work together to make Kubernetes Scheduler dedicate some nodes and assign only specific pods.

## Scheduling Bottlenecks

Although Kubernetes Scheduler is designed to select the best node,  the “best node” can change after the pods start running. As a result,  there are potential issues with pods’ resource usages and their node  assignments over the long run.

### Resource Requests and Limits: “Noisy Neighbors”

“Noisy neighbors” are not specific to Kubernetes. Any multitenant  system is a potential residency for them. For example, let’s assume you  have two pods, A and B, running on the same node. If pod B tries to  create noise by consuming all of the CPU or memory, pod A will have  problems. Luckily, setting resource [requests and limits for containers](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/%23how-pods-with-resource-requests-are-scheduled) keeps neighbors under control. Kubernetes will ensure that the containers are scheduled for their requested resources and don’t consume more than  their resource limits. If you’re running Kubernetes in production,  you’ll want to set resource requests and limits to have a reliable  system.

### Out of Resources for System Processes

Kubernetes nodes are mostly virtual machines connected to the  Kubernetes control plane. Thus, the nodes also have their own operating  systems and related processes running on them. If Kubernetes workloads  consume all of the resources, the nodes will not operate and will run  into problems. You need to set reserved resources in kubelet with the  flags **–system-reserved** to prevent this.

### Preempted or Scheduled Pods

If Kubernetes Scheduler cannot schedule a pod to an available node,  it can preempt (evict) some pods from nodes to allocate resources. If  you see pods move around the cluster without specific reasons for doing  so, consider defining them with priority classes. Similarly, if your  pods are not scheduled and are waiting for other pods, you need to check their priority classes.

In the following example, the priority class will not preempt any other pods for itself:

```
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass

metadata:
  name: high-priority-nonpreempting

value: 100000
preemptionPolicy: Never
globalDefault: false

description: "This priority class will not preempt other pods."
```

You can assign priority classes to the pods in their podspec this way:

```
apiVersion: v1
kind: Pod

metadata:
  name: nginx

spec:
  containers:
  - name: nginx
    image: nginx

priorityClassName: high-priority-nonpreempting
```

## Scheduling Framework

Kubernetes Scheduler has a pluggable scheduling framework  architecture that allows you to add a new set of plugins to the  framework. The plugins implement the Plugin API and are compiled into  the scheduler. In this section, we’ll discuss the workflow, extension  points, and Plugin API of the scheduling framework.

### Workflow and Extension Points

Scheduling a pod consists of two phases: the scheduling cycle and the binding cycle. In the scheduling cycle, the scheduler finds a feasible  node. Then, in the binding process, the decision is applied to the  cluster.

The diagram below illustrates the flow of phases and extension points:

![img](https://cdn.thenewstack.io/media/2020/11/9d4f9780-screen-shot-2020-11-17-at-11.41.00-am-1024x541.png)

Figure 1: Scheduling workflow (Source: [Kubernetes documentation](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/))

The following points in the workflow are open to plugin extension:

- **QueueSort**: Sort the pods in the queue
- **PreFilter**: Check the preconditions of the pods for scheduling cycle
- **Filter**: Filter the nodes that are not suitable for the pod
- **PostFilter**: Run if there are no feasible nodes found for the pod
- **PreScore**: Run prescoring tasks to generate shareable state for scoring plugins
- **Score**: Rank the filtered nodes by calling each scoring plugins
- **NormalizeScore**: Combine the scores and compute a final ranking of the nodes
- **Reserve**: Choose the node as reserved before the binding cycle
- **Permit**: Approve or deny the scheduling cycle result
- **PreBind**: Perform any prerequisite work, such as provisioning a network volume
- **Bind**: Assign the pods to the nodes in Kubernetes API
- **PostBind**: Inform the result of the binding cycle

Plugin extensions implement the Plugin API and are a part of Kubernetes Scheduler. You can check the interface in the [Kubernetes repository](https://github.com/kubernetes/kubernetes/blob/master/pkg/scheduler/framework/v1alpha1/interface.go). A plugin is expected to register itself with its name as follows:

```
// Plugin is the parent type for all the scheduling framework plugins.

type Plugin interface {

Name() string

}
```

The plugin also implements the related extension points, as shown below:

```
// QueueSortPlugin is an interface that must be implemented by "QueueSort" plugins.
// These plugins are used to sort pods in the scheduling queue. Only one queue sort plugin may be enabled at a time.

type QueueSortPlugin interface {

 Plugin

 // Less are used to sort pods in the scheduling queue.

 Less(*QueuedPodInfo, *QueuedPodInfo) bool
}
```

## Scheduler Performance Tuning

Kubernetes Scheduler has a workflow to find and bind a feasible node  for the pods. The scheduler’s workload increases exponentially when the  number of nodes in the cluster is very high. In large clusters, it can  take a long time to find the best node. To fine-tune the scheduler’s  performance, find a compromise between latency and accuracy.

The **percentageOfNodesToScore** sets a limit to the  number of nodes to calculate their scores. By default, Kubernetes sets a linear threshold between 50% for a 100-node cluster and 10% for a  5000-node cluster. The minimum of the default value is 5%. It ensures  that at least 5% of the nodes in your cluster are considered for  scheduling.

In the example below, you can see how to set the threshold manually by performance tuning the kube-scheduler:

```
apiVersion: kubescheduler.config.k8s.io/v1alpha1

kind: KubeSchedulerConfiguration

algorithmSource:

  provider: DefaultProvider

percentageOfNodesToScore: 50
```

It’s a good idea to change the percentage if you have a massive  cluster and your Kubernetes workloads don’t tolerate the latency caused  by Kubernetes Scheduler.

## Summary

This article covered all aspects of scheduling in Kubernetes. We  started with configurations for pods and nodes, including node  selectors, affinity rules, taints, and tolerations. We then covered  Kubernetes Scheduler’s framework, its extension points, and its API, as  well as the resource-related bottlenecks that can occur. Finally, we  presented the tool’s performance tuning settings. Although Kubernetes  Scheduler works out of the box to assign pods to nodes, it is essential  to know its dynamics and configure them for a reliable, production-grade Kubernetes setup.
