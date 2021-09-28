# A Deep Dive into Kubernetes Scheduling

# 深入研究 Kubernetes 调度

                    February 8, 2021

2021 年 2 月 8 日

Kubernetes Scheduler is one of the core components of the Kubernetes control plane. It runs on the control plane, and its  default behavior assigns pods to nodes while balancing resource  utilization among them. When the pods are assigned to a new node, the  kubelet running on the node retrieves the pod definition from the  Kubernetes API. Then, the kubelet creates the resources and containers  according to the pod specification on the node. In other words, the  scheduler runs inside the control plane and distributes the workload to  the Kubernetes cluster. It’s possible that you’ve never checked  Kubernetes Scheduler’s logs or configuration parameters because, for the most part, the tool works well for the majority of development,  testing, and production cases.

Kubernetes Scheduler 是 Kubernetes 控制平面的核心组件之一。它运行在控制平面上，其默认行为将 pod 分配给节点，同时平衡它们之间的资源利用率。当 pod 被分配到一个新节点时，在该节点上运行的 kubelet 会从 Kubernetes API 中检索 pod 定义。然后，kubelet 根据节点上的 pod 规范创建资源和容器。换句话说，调度程序在控制平面内运行，并将工作负载分配给 Kubernetes 集群。您可能从未检查过 Kubernetes Scheduler 的日志或配置参数，因为在大多数情况下，该工具适用于大多数开发、测试和生产案例。

This article will take a deep dive into Kubernetes Scheduler,  starting with an overview of scheduling in general and scheduling  eviction with affinity and taints. We’ll then discuss the scheduler’s  bottlenecks and the issues that you may run into in production. Finally, we’ll examine how to fine-tune the scheduler’s parameters to suit your  custom-tailored clusters.

本文将深入探讨 Kubernetes 调度程序，首先概述一般调度以及具有关联性和污点的调度驱逐。然后，我们将讨论调度程序的瓶颈以及您在生产中可能遇到的问题。最后，我们将研究如何微调调度程序的参数以适合您定制的集群。

## An Overview of Scheduling

## 调度概述

Kubernetes scheduling is simply the process of assigning pods to the  matched nodes in a cluster. A scheduler watches for newly created pods  and finds the best node for their assignment. It chooses the optimal  node based on Kubernetes’ scheduling principles and your configuration  options.

Kubernetes 调度只是将 pod 分配给集群中匹配节点的过程。调度程序监视新创建的 pod，并为它们的分配找到最佳节点。它根据 Kubernetes 的调度原则和您的配置选项选择最佳节点。

The simplest configuration option is setting the **nodeName** field in `podspec` directly as follows:

最简单的配置选项是直接在 `podspec` 中设置 **nodeName** 字段，如下所示：

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

默认情况下，上面的 nginx pod 将在 node-01 上运行。但是，**nodeName** 有很多限制会导致 Pod 无法正常工作，例如云中的未知节点名称、资源节点不足以及具有间歇性网络问题的节点。因此，除了测试或开发期间，您不应在任何时候使用 **nodeName**。

If you want to run your pods on a specific set of nodes, use **nodeSelector** to ensure that happens. You can define the **nodeSelector** field as a set of key-value pairs in ‘PodSpec’:

如果您想在一组特定的节点上运行 Pod，请使用 **nodeSelector** 来确保发生这种情况。您可以将 **nodeSelector** 字段定义为“PodSpec”中的一组键值对：

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

对于上面的 nginx pod，Kubernetes Scheduler 会找到一个带有 disktype: ssd 标签的节点。当然，节点可以有额外的标签，Kubernetes 已经为节点 kubernetes.io/arch 或 kubernetes.io/os 填充了通用标签。您可以在 [Kubernetes 参考文档](https://kubernetes.io/docs/reference/kubernetes-api/labels-annotations-taints/) 中查看完整的标签列表。

The use of `nodeSelector` efficiently constrains pods to  run on nodes with specific labels. However, its use is only constrained  with labels and their values. There are two more comprehensive features  in Kubernetes to express more complicated scheduling requirements: **node affinity**, to mark pods to attract them to a set of nodes; and **taints and tolerations**, to mark nodes to repel a set of pods. These features are discussed below.

`nodeSelector` 的使用有效地限制了 pod 在具有特定标签的节点上运行。但是，它的使用仅受标签及其值的限制。 Kubernetes 中有两个更全面的特性来表达更复杂的调度需求：**节点亲和**，标记 Pod 以将它们吸引到一组节点；和 **taints 和 tolerations**，用于标记节点以排斥一组 pod。下面将讨论这些功能。

## **Node Affinity**

## **节点亲和力**

Node affinity is a set of constraints defined on pods that determine  which nodes are eligible for scheduling. It’s possible to define hard  and soft requirements for the pods’ node assignments using affinity  rules. For instance, you can configure a pod to run only the nodes with  GPUs and preferably with NVIDIA_TESLA_V100 for your deep learning  workload. The scheduler evaluates the rules and tries to find a suitable node within the defined constraints. Like `nodeSelectors`, node affinity rules work with the node labels; however, they are more powerful than `nodeSelectors`.

节点亲和性是在 pod 上定义的一组约束，用于确定哪些节点有资格进行调度。可以使用关联规则为 Pod 的节点分配定义硬要求和软要求。例如，您可以将 pod 配置为仅运行具有 GPU 的节点，最好使用 NVIDIA_TESLA_V100 来运行深度学习工作负载。调度程序评估规则并尝试在定义的约束内找到合适的节点。与 `nodeSelectors` 一样，节点关联规则与节点标签一起工作；然而，它们比`nodeSelectors` 更强大。

There are four affinity rules you can add to `podspec`:

您可以将四个关联规则添加到 `podspec`：

- **requiredDuringSchedulingIgnoredDuringExecution**
- **requiredDuringSchedulingRequiredDuringExecution**
- **preferredDuringSchedulingIgnoredDuringExecution**
- **preferredDuringSchedulingRequiredDuringExecution** 

- **requiredDuringSchedulingIgnoredDuringExecution**
- **requiredDuringSchedulingRequiredDuringExecution**
- **preferredDuringSchedulingIgnoredDuringExecution**
- **preferredDuringSchedulingRequiredDuringExecution**

These four rules consist of two criteria: required or preferred, and  two stages: Scheduling and Execution. Rules starting with required  describe hard requirements that must be met. Rules beginning with  preferred are soft requirements that will be enforced but not  guaranteed. The Scheduling stage refers to the first assignment of the  pod to the nodes. The Execution stage applies to situations where node  labels change after the scheduling assignment.

这四个规则由两个标准组成：必需或首选，以及两个阶段：调度和执行。以 required 开头的规则描述了必须满足的硬性要求。以首选开头的规则是将强制执行但不保证的软要求。调度阶段是指第一次将 pod 分配给节点。执行阶段适用于调度分配后节点标签发生变化的情况。

If a rule is stated as **IgnoredDuringExecution**, the scheduler will not check its validity after the first assignment. However, if the rule is specified with **RequiredDuringExecution**, the scheduler will always ensure the rule’s validity by moving the pod to a suitable node.

如果规则声明为 **IgnoredDuringExecution**，则调度程序将不会在第一次分配后检查其有效性。但是，如果使用 **RequiredDuringExecution** 指定规则，则调度程序将始终通过将 pod 移动到合适的节点来确保规则的有效性。

Check out the following example to help you grasp these affinities:

查看以下示例以帮助您掌握这些相似性：

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

上面的 nginx pod 有一个节点关联规则，表明 Kubernetes Scheduler 应该只将 pod 放置到 us-east 区域中的节点。第二条规则表示应首选 us-east-1 或 us-east-2 区域。

Using affinity rules, you can make Kubernetes scheduling decisions work for your custom requirements.

使用关联规则，您可以使 Kubernetes 调度决策适合您的自定义需求。

## **Taints and Tolerations**

## **污点和容忍**

Not all Kubernetes nodes are the same in a cluster. It’s possible to  have nodes with special hardware, such as GPU, disk, or network  capabilities. Similarly, you may need to dedicate some nodes for  testing, data protection, or user groups. Taints can be added to the  nodes to repel pods, as in the following example:

并非所有 Kubernetes 节点在集群中都是相同的。可能有具有特殊硬件的节点，例如 GPU、磁盘或网络功能。同样，您可能需要将一些节点专用于测试、数据保护或用户组。可以将污点添加到节点以排斥 Pod，如下例所示：

```
kubectl taint nodes node1 test-environment=true:NoSchedule
```

With taint `test-environment=true:NoSchedule`, Kubernetes Scheduler will not assign any pod unless it has matching toleration in the`podspec`:

使用 taint `test-environment=true:NoSchedule`，Kubernetes Scheduler 不会分配任何 pod，除非它在 `podspec` 中有匹配的容忍度：

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

污点和容忍度一起工作，使 Kubernetes 调度程序专用于一些节点并仅分配特定的 pod。

## Scheduling Bottlenecks

## 调度瓶颈

Although Kubernetes Scheduler is designed to select the best node,  the “best node” can change after the pods start running. As a result,  there are potential issues with pods’ resource usages and their node  assignments over the long run.

尽管 Kubernetes Scheduler 旨在选择最佳节点，但“最佳节点”可能会在 Pod 开始运行后发生变化。因此，从长远来看，Pod 的资源使用及其节点分配存在潜在问题。

### Resource Requests and Limits: “Noisy Neighbors”

### 资源请求和限制：“吵闹的邻居”

“Noisy neighbors” are not specific to Kubernetes. Any multitenant  system is a potential residency for them. For example, let’s assume you  have two pods, A and B, running on the same node. If pod B tries to  create noise by consuming all of the CPU or memory, pod A will have  problems. Luckily, setting resource [requests and limits for containers](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/%23how-pods-with-resource-requests-are-scheduled) keeps neighbors under control. Kubernetes will ensure that the containers are scheduled for their requested resources and don’t consume more than  their resource limits. If you’re running Kubernetes in production,  you’ll want to set resource requests and limits to have a reliable  system.

“嘈杂的邻居”并不是 Kubernetes 特有的。任何多租户系统都是他们的潜在驻地。例如，假设您有两个 Pod，A 和 B，在同一个节点上运行。如果 pod B 试图通过消耗所有 CPU 或内存来制造噪音，那么 pod A 就会出现问题。幸运的是，设置资源 [容器的请求和限制](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/%23how-pods-with-resource-requests-are-scheduled) 使邻居保持在控制。 Kubernetes 将确保容器为其请求的资源进行调度，并且不会消耗超过其资源限制。如果您在生产中运行 Kubernetes，您需要设置资源请求和限制以拥有一个可靠的系统。

### Out of Resources for System Processes 

### 系统进程资源不足

Kubernetes nodes are mostly virtual machines connected to the  Kubernetes control plane. Thus, the nodes also have their own operating  systems and related processes running on them. If Kubernetes workloads  consume all of the resources, the nodes will not operate and will run  into problems. You need to set reserved resources in kubelet with the  flags **–system-reserved** to prevent this.

Kubernetes 节点大多是连接到 Kubernetes 控制平面的虚拟机。因此，节点也有自己的操作系统和运行在其上的相关进程。如果 Kubernetes 工作负载消耗了所有资源，节点将无法运行并会遇到问题。您需要在 kubelet 中使用标志 **–system-reserved** 设置保留资源以防止这种情况发生。

### Preempted or Scheduled Pods

### 抢占或预定的 Pod

If Kubernetes Scheduler cannot schedule a pod to an available node,  it can preempt (evict) some pods from nodes to allocate resources. If  you see pods move around the cluster without specific reasons for doing  so, consider defining them with priority classes. Similarly, if your  pods are not scheduled and are waiting for other pods, you need to check their priority classes.

如果 Kubernetes Scheduler 无法将 Pod 调度到可用节点，它可以从节点中抢占（驱逐）一些 Pod 以分配资源。如果您看到 pod 在没有特定原因的情况下在集群中移动，请考虑使用优先级类定义它们。同样，如果您的 pod 没有被调度并且正在等待其他 pod，您需要检查它们的优先级类。

In the following example, the priority class will not preempt any other pods for itself:

在以下示例中，优先级类不会为自己抢占任何其他 Pod：

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

您可以通过这种方式为 podspec 中的 pod 分配优先级类：

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

## 调度框架

Kubernetes Scheduler has a pluggable scheduling framework  architecture that allows you to add a new set of plugins to the  framework. The plugins implement the Plugin API and are compiled into  the scheduler. In this section, we’ll discuss the workflow, extension  points, and Plugin API of the scheduling framework.

Kubernetes Scheduler 具有可插拔的调度框架架构，允许您向框架添加一组新插件。插件实现插件 API 并编译到调度程序中。在本节中，我们将讨论调度框架的工作流、扩展点和插件 API。

### Workflow and Extension Points

### 工作流和扩展点

Scheduling a pod consists of two phases: the scheduling cycle and the binding cycle. In the scheduling cycle, the scheduler finds a feasible  node. Then, in the binding process, the decision is applied to the  cluster.

调度 Pod 包括两个阶段：调度周期和绑定周期。在调度周期中，调度器寻找一个可行的节点。然后，在绑定过程中，将决策应用于集群。

The diagram below illustrates the flow of phases and extension points:

下图说明了阶段和扩展点的流程：

![img](https://cdn.thenewstack.io/media/2020/11/9d4f9780-screen-shot-2020-11-17-at-11.41.00-am-1024x541.png)

Figure 1: Scheduling workflow (Source: [Kubernetes documentation](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/))

图 1：调度工作流（来源：[Kubernetes 文档](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/))

The following points in the workflow are open to plugin extension:

工作流中的以下几点对插件扩展开放：

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

- **QueueSort**：对队列中的 Pod 进行排序
- **PreFilter**：检查 Pod 调度周期的前置条件
- **Filter**：过滤掉不适合pod的节点
- **PostFilter**：如果没有为 pod 找到可行节点则运行
- **PreScore**：运行预评分任务以生成评分插件的可共享状态
- **Score**：通过调用每个评分插件对过滤的节点进行排名
- **NormalizeScore**：组合分数并计算节点的最终排名
- **Reserve**：在绑定周期前选择节点作为保留
- **Permit**：批准或拒绝调度周期结果
- **PreBind**：执行任何先决条件工作，例如配置网络卷
- **Bind**：将 Pod 分配给 Kubernetes API 中的节点
- **PostBind**：通知绑定周期结果

Plugin extensions implement the Plugin API and are a part of Kubernetes Scheduler. You can check the interface in the [Kubernetes repository](https://github.com/kubernetes/kubernetes/blob/master/pkg/scheduler/framework/v1alpha1/interface.go). A plugin is expected to register itself with its name as follows:

插件扩展实现了插件 API，是 Kubernetes 调度器的一部分。可以在 [Kubernetes 仓库](https://github.com/kubernetes/kubernetes/blob/master/pkg/scheduler/framework/v1alpha1/interface.go)查看接口。插件应该用它的名字注册自己，如下所示：

```
// Plugin is the parent type for all the scheduling framework plugins.

type Plugin interface {

Name() string

}
```

The plugin also implements the related extension points, as shown below:

该插件还实现了相关的扩展点，如下图：

```
// QueueSortPlugin is an interface that must be implemented by "QueueSort" plugins.
// These plugins are used to sort pods in the scheduling queue.Only one queue sort plugin may be enabled at a time.

type QueueSortPlugin interface {

 Plugin

 // Less are used to sort pods in the scheduling queue.

 Less(*QueuedPodInfo, *QueuedPodInfo) bool
}
```

## Scheduler Performance Tuning 

## 调度器性能调优

Kubernetes Scheduler has a workflow to find and bind a feasible node  for the pods. The scheduler’s workload increases exponentially when the  number of nodes in the cluster is very high. In large clusters, it can  take a long time to find the best node. To fine-tune the scheduler’s  performance, find a compromise between latency and accuracy.

Kubernetes Scheduler 有一个工作流来为 pod 查找和绑定一个可行的节点。当集群中的节点数量非常多时，调度程序的工作负载呈指数增长。在大型集群中，可能需要很长时间才能找到最佳节点。要微调调度程序的性能，请在延迟和准确性之间找到折衷。

The **percentageOfNodesToScore** sets a limit to the  number of nodes to calculate their scores. By default, Kubernetes sets a linear threshold between 50% for a 100-node cluster and 10% for a  5000-node cluster. The minimum of the default value is 5%. It ensures  that at least 5% of the nodes in your cluster are considered for  scheduling.

**percentageOfNodesToScore** 对计算其分数的节点数量设置了限制。默认情况下，Kubernetes 将线性阈值设置为 100 节点集群的 50% 和 5000 节点集群的 10% 之间。默认值的最小值为 5%。它确保您的集群中至少有 5% 的节点被考虑用于调度。

In the example below, you can see how to set the threshold manually by performance tuning the kube-scheduler:

在下面的示例中，您可以看到如何通过对 kube-scheduler 进行性能调优来手动设置阈值：

```
apiVersion: kubescheduler.config.k8s.io/v1alpha1

kind: KubeSchedulerConfiguration

algorithmSource:

  provider: DefaultProvider

percentageOfNodesToScore: 50
```

It’s a good idea to change the percentage if you have a massive  cluster and your Kubernetes workloads don’t tolerate the latency caused  by Kubernetes Scheduler.

如果您有一个庞大的集群并且您的 Kubernetes 工作负载不能容忍 Kubernetes Scheduler 引起的延迟，那么更改百分比是个好主意。

## Summary

##  概括

This article covered all aspects of scheduling in Kubernetes. We  started with configurations for pods and nodes, including node  selectors, affinity rules, taints, and tolerations. We then covered  Kubernetes Scheduler’s framework, its extension points, and its API, as  well as the resource-related bottlenecks that can occur. Finally, we  presented the tool’s performance tuning settings. Although Kubernetes  Scheduler works out of the box to assign pods to nodes, it is essential  to know its dynamics and configure them for a reliable, production-grade Kubernetes setup. 

本文涵盖了 Kubernetes 中调度的所有方面。我们从 pod 和节点的配置开始，包括节点选择器、关联规则、污点和容忍度。然后我们介绍了 Kubernetes Scheduler 的框架、扩展点和 API，以及可能出现的与资源相关的瓶颈。最后，我们介绍了该工具的性能调整设置。尽管 Kubernetes Scheduler 可以开箱即用地将 pod 分配给节点，但了解其动态并将其配置为可靠的生产级 Kubernetes 设置至关重要。

