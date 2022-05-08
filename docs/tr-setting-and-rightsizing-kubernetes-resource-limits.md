# Setting and Rightsizing Kubernetes Resource Limits \| Best Practices

# 设置和调整 Kubernetes 资源限制 \|最佳实践

September 7, 2021

2021 年 9 月 7 日

Part of managing a Kubernetes cluster is making sure your clusters aren’t using too many resources. Let’s walk through the concepts of setting and rightsizing resource limits for Kubernetes.

管理 Kubernetes 集群的一部分是确保您的集群没有使用太多资源。让我们了解一下为 Kubernetes 设置和调整资源限制的概念。

Overview

概述

1. [What Are Resource Limits, and Why Do They Matter?](http://www.containiq.com#1)
2. [Implementing Rightsizing](http://www.containiq.com#2)
3. [Calculate Resource Limits in 3 Steps](http://www.containiq.com#3)
4. [Using ContainIQ for Rightsizing & Resource Limits](http://www.containiq.com#4)
5. [How to Set Resource Limits](http://www.containiq.com#5)
6. [What Are the Next Steps?](http://www.containiq.com#6)
7. [Conclusion](http://www.containiq.com#7)



1. [什么是资源限制，它们为什么重要？](http://www.containiq.com#1)
2. [实施大小调整](http://www.containiq.com#2)
3. [3步计算资源限制](http://www.containiq.com#3)
4. [使用 ContainIQ 调整大小和资源限制](http://www.containiq.com#4)
5. [如何设置资源限制](http://www.containiq.com#5)
6. [下一步是什么？](http://www.containiq.com#6)
7. [结论](http://www.containiq.com#7)

Managing a Kubernetes cluster is like sitting in front of a large DJ mixing console. You have an almost overwhelming number of knobs and sliders in front of you to tune the sound to perfection. It can seem challenging to even know where to begin—a feeling that Kubernetes engineers probably know well.

管理 Kubernetes 集群就像坐在大型 DJ 混音台前。您面前有几乎压倒性的旋钮和滑块来将声音调至完美。甚至知道从哪里开始似乎都具有挑战性——Kubernetes 工程师可能很清楚这种感觉。

However, a [Kubernetes cluster](https://www.containiq.com/post/kubernetes-cluster) without resource limits could conceivably lead to the most issues. So setting resource limits is a logical place to start.

但是，没有资源限制的 [Kubernetes 集群](https://www.containiq.com/post/kubernetes-cluster) 可能会导致大多数问题。因此，设置资源限制是一个合乎逻辑的起点。

That being said, your fine-tuning challenges are just beginning. To set Kubernetes resources limits correctly, you have to be methodical and take care to find the correct values. Set them too high, and you may negatively affect nodes of your clusters. Set the value too low, and you negatively impact your application performances.

话虽如此，您的微调挑战才刚刚开始。要正确设置 Kubernetes 资源限制，您必须有条不紊并注意找到正确的值。将它们设置得太高，您可能会对集群的节点产生负面影响。将该值设置得太低，会对应用程序性能产生负面影响。

Fortunately, this guide will walk you through all you need to know to properly tune resource limits and keep both your cluster and your applications healthy.

幸运的是，本指南将引导您了解正确调整资源限制并保持集群和应用程序健康所需的所有知识。

## What Are Resource Limits, and Why Do They Matter?

## 什么是资源限制，它们为什么重要？

First of all, in Kubernetes, resources limits come in pairs with resource requests:

首先，在 Kubernetes 中，资源限制与资源请求成对出现：

- **Resource Requests:** The amount of CPU or memory allocated to your container. A Pod resource request is equal to the sum of its container resource requests. When scheduling Pods, Kubernetes will guarantee that this amount of resource is available for your Pod to run.
- **Resource Limits:** The level at which Kubernetes will start taking action against a container going above the limit. Kubernetes will kill a container consuming too much memory or throttle a container using too much CPU.

- **资源请求：**分配给容器的 CPU 或内存量。 Pod 资源请求等于其容器资源请求的总和。在调度 Pod 时，Kubernetes 将保证此数量的资源可供您的 Pod 运行。
- **资源限制：** Kubernetes 将开始对超出限制的容器采取行动的级别。 Kubernetes 会杀死一个消耗过多内存的容器或限制一个使用过多 CPU 的容器。

If you set resource limits but no resource request, Kubernetes implicitly sets memory and CPU requests equal to the limit. This behavior is excellent as a first step toward getting your Kubernetes cluster under control. This is often referred to as the _conservative approach_, where resources allocated to a container are at maximum.

如果设置了资源限制但没有资源请求，Kubernetes 会隐式设置内存和 CPU 请求等于限制。这种行为非常适合作为控制 Kubernetes 集群的第一步。这通常被称为_保守方法_，其中分配给容器的资源最多。

### Have You Set Requests and Limits on All Your Containers?

### 您是否对所有容器设置了请求和限制？

If not, the [Kubernetes Scheduler](https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/) will assign any Pods without request and limits randomly. With limits set, you will avoid most of the following problems:

如果没有，[Kubernetes 调度程序](https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/) 将随机分配任何没有请求和限制的 Pod。设置限制后，您将避免以下大多数问题：

- **Out of Memory (OOM) issues.** A node could die of memory starvation, affecting cluster stability. For instance, an application with [a memory leak could cause an OOM issue](https://www.containiq.com/post/oomkilled-troubleshooting-kubernetes-memory-requests-and-limits).
- **CPU Starvation.** Your applications will get slower because they must share a limited amount of CPU. An application consuming an excessive amount of CPU could affect all applications on the same node.
- **Pod eviction.** When a node lacks resources, it [starts the eviction process](https://www.containiq.com/post/kubernetes-pod-evictions) and terminates Pods, starting with Pods without resource requests .
- **Financial waste.** Assuming your cluster is fine without requests and limits, this means that you are most likely overprovisioning. In other words, you’re spending money on resources you never use.

- **内存不足 (OOM) 问题。** 节点可能死于内存不足，影响集群稳定性。例如，具有 [内存泄漏可能导致 OOM 问题](https://www.containiq.com/post/oomkilled-troubleshooting-kubernetes-memory-requests-and-limits) 的应用程序。
- **CPU 饥饿。**您的应用程序将变得更慢，因为它们必须共享有限数量的 CPU。消耗过多 CPU 的应用程序可能会影响同一节点上的所有应用程序。
- **Pod eviction.** 当一个节点缺乏资源时，它[启动驱逐过程](https://www.containiq.com/post/kubernetes-pod-evictions) 并终止Pods，从没有资源请求的Pods开始.
- **财务浪费。**假设您的集群在没有请求和限制的情况下很好，这意味着您很可能过度配置。换句话说，你把钱花在了你从未使用过的资源上。

The first step to prevent these issues is to set limits for all containers.

防止这些问题的第一步是为所有容器设置限制。

Imagine your node as a big box with a width (CPU) and length (memory), and you need to fill the big box with smaller boxes (Pods). Not setting requests and limits is similar to playing the game without knowing the width and length of the smaller boxes. 

把你的节点想象成一个有宽度（CPU）和长度（内存）的大盒子，你需要用更小的盒子（Pods）来填充这个大盒子。不设置请求和限制类似于在不知道小盒子的宽度和长度的情况下玩游戏。

Assuming you only set limits, Kubernetes will check nodes and find a fit for your Pod. Pod assignments will become more coherent. Also, you will have a clear picture of the situation on each node, as illustrated below.

假设您只设置限制，Kubernetes 将检查节点并找到适合您的 Pod。 Pod 分配将变得更加连贯。此外，您将清楚地了解每个节点的情况，如下图所示。



From there, you can start observing your application more closely and begin optimizing resource requests. You will maximize resource utilizations, making your system more cost-efficient.

从那里，您可以开始更密切地观察您的应用程序并开始优化资源请求。您将最大限度地利用资源，使您的系统更具成本效益。

## Implementing Rightsizing

## 实施大小调整

What do you need to calculate resource limits? The first thing is metrics!

你需要什么来计算资源限制？首先是指标！

Resource limits are calculated based on historical data. Kubernetes doesn’t come out of the box with sufficient tools to gather memory and CPU, so here’s a small list of options:

资源限制是根据历史数据计算的。 Kubernetes 没有开箱即用的工具来收集内存和 CPU，所以这里有一小部分选项：

- The[Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server) collects and aggregates metrics. However, you won’t get far with it because it only gives you the current value of a metric.
- [Prometheus](https://prometheus.io/) is a popular solution to monitor anything.
- ContainIQ is a monitoring solution tailored for Kubernetes.

- [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server) 收集和汇总指标。但是，您不会走得太远，因为它只为您提供指标的当前值。
- [Prometheus](https://prometheus.io/) 是监控任何事物的流行解决方案。
- ContainIQ 是为 Kubernetes 量身定制的监控解决方案。

The next thing you need is to identify Pods without limits.

接下来你需要做的是无限制地识别 Pod。

```c
kubectl get pods --all-namespaces -o go-template-file=./get-pods-without-limits.gotemplate
kubectl get pods --all-namespaces -o go-template-file=./get-pods-without-limits.gotemplate
```



The following  go-template iterates through all Pods and containers and outputs a container with metrics.

下面的 go-template 遍历所有 Pod 和容器，并输出一个带有指标的容器。

```c hljs
{{- range .items -}}
{{ $name := .metadata.name }}
{{- range .spec.containers -}}
{{- if not .resources.limits -}}
{{"no limits for: "}}{{$name}}{{"/"}}{{.name}}{{"\n"}}
{{- end -}}
{{- end -}}
{{- end -}}
```

Now you have a list of Pods to work on.

现在你有了一个要处理的 Pod 列表。

![Pods without resource limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6137caa585b79a0e1667afb1_MsKPWHS.png)

Pods without resource limits

没有资源限制的 Pod


Finally, you need to revise Kubernetes deployment, [StatefulSet](https://www.containiq.com/post/kubernetes-statefulsets), or [DaemonSets](https://www.containiq.com/post/using-kubernetes-daemonsets-effectively) for each Pod you found and include resource limits. Unfortunately, when it comes to resource limits, there is no magic formula that fits all cases.

最后，您需要修改 Kubernetes 部署，[StatefulSet](https://www.containiq.com/post/kubernetes-statefulsets) 或 [DaemonSets](https://www.containiq.com/post/using-kubernetes-daemonsets-effectively) 为您找到的每个 Pod 提供资源限制。不幸的是，当谈到资源限制时，没有适合所有情况的神奇公式。

In this activity, you need to use your infrastructure knowledge and challenge the numbers you see. Let me explain—CPU and memory usage is affected by the type of application you are using. For example:

在本活动中，您需要使用您的基础架构知识并挑战您看到的数字。让我解释一下——CPU 和内存使用受您使用的应用程序类型的影响。例如：

- **Microservices** tackle HTTP requests, consuming resources based on the traffic. You often have many instances of a microservice and have good toleration in terms of failure. Tight resource limits are acceptable and help prevent abnormal behavior.
- **Databases** tend to consume increasingly more memory over time. On top of that, you have no tolerance for failure; tight resource limits are not an option.
- **ETL/ELT** tends to consume resources by burst, but memory and CPU usage are mostly static. Use resource limits to prevent unexpected bursts of resource usage.

- **微服务**处理 HTTP 请求，根据流量消耗资源。您通常有许多微服务实例，并且在故障方面具有良好的容忍度。严格的资源限制是可以接受的，有助于防止异常行为。
- **数据库**随着时间的推移会消耗越来越多的内存。最重要的是，你不能容忍失败；严格的资源限制不是一种选择。
- **ETL/ELT**倾向于突发消耗资源，但内存和CPU使用大多是静态的。使用资源限制来防止资源使用的意外爆发。

Setting up limits leads to fine-tuning other settings of your cluster, such as node size or [auto-scaling](https://www.containiq.com/post/kubernetes-autoscaling). We’ll come back to that later on. For now, let’s focus on a general strategy for calculating limits.

设置限制会导致微调集群的其他设置，例如节点大小或 [auto-scaling](https://www.containiq.com/post/kubernetes-autoscaling)。我们稍后再谈。现在，让我们专注于计算限制的一般策略。

## Calculate Resource Limits in 3 Steps

## 分 3 步计算资源限制

To carefully analyze the metrics, you need to take it one step at a time. This approach consists of three phases that each lead to a different strategy. Start with the most aggressive, challenge the result, and move on to a more conservative option if necessary.

要仔细分析指标，您需要一步一步来。这种方法由三个阶段组成，每个阶段都会导致不同的策略。从最激进的开始，挑战结果，并在必要时转向更保守的选择。

_You must consider CPU and memory independently and apply different strategies based on your conclusion at each phase._

_您必须独立考虑CPU和内存，并根据您在每个阶段的结论应用不同的策略。_

### 1\. Check Memory or CPU 

### 1. 检查内存或 CPU

In the first phase, look at the ninety-ninth percentile of memory or CPU. This _aggressive approach_ aims to reduce issues by forcing Kubernetes to take action against outliers. If you set limits with this value, your application will be affected 1 percent of the time. The container CPU will be throttled and never reach that value again.

在第一阶段，查看内存或 CPU 的百分之九十九。这种“激进的方法”旨在通过强制 Kubernetes 对异常值采取行动来减少问题。如果您使用此值设置限制，您的应用程序将有 1% 的时间受到影响。容器 CPU 将受到限制，并且永远不会再次达到该值。

The aggressive approach is often good for CPU limits because the consequences are relatively acceptable and help you better manage resources. Concerning memory, the ninety-ninth percentile may be problematic; containers restart if the limit is reached.

激进的方法通常有利于 CPU 限制，因为其后果相对可以接受并帮助您更好地管理资源。关于记忆，百分之九十九可能有问题；如果达到限制，容器将重新启动。

At this point, you should weigh the consequences and conclude if the ninety-ninth percentile makes sense for you (as a side note, you should investigate deeper as to why applications sometimes reach the set limit). Maybe the ninety-ninth percentile is too restrictive because your application is not yet at its max utilization. In that case, move on to the second strategy to set limits.

此时，您应该权衡后果并得出结论，如果 99 个百分位对您有意义（作为旁注，您应该更深入地调查为什么应用程序有时会达到设定的限制）。可能 99% 的限制过于严格，因为您的应用程序尚未达到最大利用率。在这种情况下，请继续使用第二个策略来设置限制。

### 2\. Check Memory or CPU in a Given Period

### 2. 检查给定时间段内的内存或 CPU

In the second phase, you’ll look at the max CPU or memory in a given period. If you set a limit with this value, in theory, no application will be affected. However, it prevents your applications from moving past that limit and keeps your cluster under control.

在第二阶段，您将查看给定时间段内的最大 CPU 或内存。如果你用这个值设置一个限制，理论上不会影响任何应用程序。但是，它可以防止您的应用程序超出该限制并控制您的集群。

Once again, you should challenge the value you found. Is the max far from the ninety-ninth percentile found earlier? Are your applications under a lot of traffic? Do you expect your application to handle more load?

再一次，你应该挑战你发现的价值。最大值与之前发现的 99 个百分位数相去甚远吗？您的应用程序是否有大量流量？您是否希望您的应用程序能够处理更多负载？

At this point, the decision branches out again with two paths. The max makes sense, and applications should be stopped or throttled if they reach that limit. On the other hand, if the maximum is much greater than the ninety-ninth percentile (outlier), or you know you need more room in the future, move to the final option.

在这一点上，决策再次分支出两条路径。最大值是有意义的，如果应用程序达到该限制，则应停止或限制它们。另一方面，如果最大值远大于第 99 个百分位数（离群值），或者您知道将来需要更多空间，请转到最后一个选项。

### 3\. Find a Compromise

### 3. 找到妥协

The last stage of this three-step process is to find a compromise based on the maximum by adding or subtracting a coefficient (ie., max + 20%). If you ever reach this point, you should consider performing load tests to characterize your application performances and resource usage better.

这个三步过程的最后阶段是通过添加或减去一个系数（即，最大值 + 20%）来找到基于最大值的折衷方案。如果您达到这一点，您应该考虑执行负载测试以更好地描述您的应用程序性能和资源使用情况。

Repeat this process for each of your applications without limits.

无限制地对每个应用程序重复此过程。

## Using ContainIQ for Rightsizing & Resource Limits

## 使用 ContainIQ 调整大小和资源限制

As discussed above, Kubernetes, by default, does not come with tooling to view pod and node-level metrics over time. Engineering teams often look to third-party tools to deliver both real-time and historical views of cluster performance. ContainIQ, a [Kubernetes monitoring platform](https://www.containiq.com/kubernetes-monitoring), can help engineering teams monitor, track, and alert on core metrics. These tools are quite helpful when it comes to implementing rightsizing and setting resource limits.

如上所述，默认情况下，Kubernetes 不提供随时间查看 pod 和节点级别指标的工具。工程团队经常使用第三方工具来提供集群性能的实时和历史视图。 ContainIQ 是一个 [Kubernetes 监控平台](https://www.containiq.com/kubernetes-monitoring)，可以帮助工程团队对核心指标进行监控、跟踪和警报。这些工具在实现大小调整和设置资源限制时非常有用。

### Implementing Rightsizing

### 实施大小调整

Rightsizing nodes is tricky without historical data. Engineering teams often allocate too much and are left with unused resources. This brings peace of mind, but with added cost and likely waste.

如果没有历史数据，调整节点大小是很棘手的。工程团队经常分配过多，留下未使用的资源。这让您高枕无忧，但会增加成本并可能造成浪费。

Fortunately, ContainIQ collects and stores both real-time and historical node metric data. By clicking on a node, users can view node CPU and memory on the Nodes dashboard:

幸运的是，ContainIQ 收集并存储实时和历史节点度量数据。通过单击一个节点，用户可以在 Nodes 仪表板上查看节点 CPU 和内存：

‍

‍

![ContainIQ Node Conditions](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6241c38b07b0a12908732f56_rightsize_nodes.png)

Nodes are color-coded based on relative usage. And users can use the See Pods tab to see CPU Requests and Limits by pod. Below, users are able to view historical node metrics by node as well as the average across all nodes:

节点根据相对使用情况进行颜色编码。用户可以使用 See Pods 选项卡按 pod 查看 CPU 请求和限制。下面，用户可以按节点查看历史节点指标以及所有节点的平均值：

‍

‍

![ContainIQ Node Limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6241c3d50dcc9ddcfbfa5ca0_node_rightsize2.png)

Using the Show Limits toggle, users are able to view the node’s capacity or the amount allocatable alongside historical performance. For example, in the screenshot above, you can see that the given node is continually hitting its CPU allocatable limit. And it might make sense to either increase the size of the node or set up the cluster autoscaler to account for these spikes.

使用显示限制切换，用户可以查看节点的容量或可分配的数量以及历史性能。例如，在上面的屏幕截图中，您可以看到给定节点不断达到其 CPU 可分配限制。增加节点的大小或设置集群自动缩放器来解决这些峰值可能是有意义的。

Users are able to more accurately size nodes, which comes with a number of benefits for end-user performance and likely cost savings to the organization.

用户能够更准确地调整节点大小，这为最终用户性能带来了许多好处，并可能为组织节省成本。

### Setting Resource Limits 

### 设置资源限制

Gathering the right data is an important first step in setting good resource limits. With ContainIQ, users are able to view the CPU and memory for every pod by simply clicking on the given pod:

收集正确的数据是设置良好资源限制的重要第一步。使用 ContainIQ，用户只需单击给定的 pod，即可查看每个 pod 的 CPU 和内存：

‍

‍

![ContainIQ Pod Conditions](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6241c3f87d17744bf1168bee_rightsize_pods.png)

In addition, users are able to view in real-time the statuses and recent events for particular pods. Users are also able to view the historical CPU and memory overtime for particular pods, as well as the average across all pods:

此外，用户可以实时查看特定 pod 的状态和最近发生的事件。用户还可以查看特定 pod 的历史 CPU 和内存超时，以及所有 pod 的平均值：

‍

‍

![ContainIQ Pod Limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6241c428b23ac908d97217aa_pod_rightsize2.png)

Having this data is invaluable when calculating the necessary resource limits and allows for a more precise calculation. The dropdown allows users to easily change between time periods ranging from hourly to weekly. Using the Show Limits toggle, users are able to see the historical data alongside the limits set at that point in time. If no limits are currently set for the given pod or pods the Show Limits toggle won’t return any data.

在计算必要的资源限制时，拥有这些数据是非常宝贵的，并且可以进行更精确的计算。下拉列表允许用户在从每小时到每周的时间段之间轻松更改。使用显示限制切换，用户可以查看历史数据以及当时设置的限制。如果当前没有为给定的 pod 设置限制，则 Show Limits 切换不会返回任何数据。

When deciding on your resource limits for a given pod or pods, users can use ContainIQ to make accurate limits backed by historical data. And in the future, users can set alerts when a limit is exceeded for a given pod or pods.

在决定给定 pod 的资源限制时，用户可以使用 ContainIQ 制定由历史数据支持的准确限制。将来，用户可以在给定的一个或多个 pod 超出限制时设置警报。

You can sign up for ContainIQ [here](https://www.containiq.com/pricing?utm_source=blog&utm_medium=rightsizing&utm_campaign=plug&utm_content=click-here), or [book a demo](https://www.containiq.com/book-a-demo?utm_source=blog&utm_medium=rightsizing&utm_campaign=plug&utm_content=book-a-demo) to learn more.

您可以在 [此处](https://www.containiq.com/pricing?utm_source=blog&utm_medium=rightsizing&utm_campaign=plug&utm_content=click-here) 或 [预订演示](https://www.containiq.com/book-a-demo?utm_source=blog&utm_medium=rightsizing&utm_campaign=plug&utm_content=book-a-demo) 了解更多信息。

## How to Set Resource Limits

## 如何设置资源限制

Once you have calculated CPU and memory limits, it’s time for you to update your configuration and redeploy your application. Limits are defined at the container level as follows:

一旦计算出 CPU 和内存限制，就该更新配置并重新部署应用程序了。限制在容器级别定义如下：

```yaml
---
apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
kind: Deployment
metadata:
name: frontend
labels:
app: guestbook
spec:
replicas: 3
template:
metadata:
labels:
     app: guestbook
     tier: frontend
spec:
containers:
    - name: php-redis
     image: gcr.io/google-samples/gb-frontend:v4
     # Here is the section you define requests
     resources:
      limits:
       cpu: 100m # the CPU limit is define in milicore (m)
       memory: 100Mi # the Memory is define in Mebibytes (Mi)
     env:
     - name: GET_HOSTS_FROM
      value: dns
     ports:
     - containerPort: 80
```

## What Are the Next Steps?

##  什么是下一个步骤？

I warned you earlier that setting limits isn’t the end of the story. It’s the first step along the way toward refining your Kubernetes cluster management. Considering your resource limits will lead you to reconsidering other settings for your cluster. That’s why I like the metaphor of the mixing console when we think about administering Kubernetes clusters.

我早些时候警告过你，设定限制并不是故事的结局。这是完善 Kubernetes 集群管理的第一步。考虑您的资源限制将导致您重新考虑集群的其他设置。这就是为什么当我们考虑管理 Kubernetes 集群时，我喜欢混合控制台的比喻。

In the following diagram, you can see the path of upcoming tasks you’ll need to identify and optimize for efficiency.

在下图中，您可以看到需要识别和优化以提高效率的即将执行任务的路径。

![What to do after setting requests and limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6233710959c0cb5765ec51f0_K8_resource-last-step.png)

To understand the diagram, consider a box as an activity or a setting related to Kubernetes. An arrow linking boxes together represents related tasks and the direction of impact. For instance _Request & Limits_ impacts _node sizing_.

要理解该图，请将框视为与 Kubernetes 相关的活动或设置。将方框连接在一起的箭头表示相关任务和影响方向。例如_请求和限制_影响_节点大小_。

Other labels in the diagram include:

图中的其他标签包括：

- **Resource Quota:** When dealing with many teams and applications, setting limits can be daunting. Fortunately Kubernetes administrators can set limits at the namespace level.
- **Node Sizing:** Choosing node size is essential to optimize resource usage. But it relies on understanding application needs.
- **Pods and Custer Auto-Scaling**: Resource usage is greatly affected by how many Pods are available to handle the load.
- **Readiness & Liveness**: Properly managing Pod lifecycle can prevent many issues. [For example](https://www.containiq.com/post/kubernetes-readiness-probe), if a Pod is consuming too many resources, it may not be ready to receive traffic.

- **资源配额：** 在处理许多团队和应用程序时，设置限制可能会令人生畏。幸运的是，Kubernetes 管理员可以在命名空间级别设置限制。
- **节点大小：** 选择节点大小对于优化资源使用至关重要。但它依赖于理解应用程序的需求。
- **Pods 和 Custer Auto-Scaling**：资源使用很大程度上受可用于处理负载的 Pod 数量的影响。
- **Readiness & Liveness**：正确管理 Pod 生命周期可以防止许多问题。 [例如](https://www.containiq.com/post/kubernetes-readiness-probe)，如果一个 Pod 消耗的资源太多，它可能还没有准备好接收流量。

## Conclusion

##  结论

You started your journey in managing a Kubernetes cluster by learning about requests and limits and quickly learned that establishing resource limits for all your containers is a solid first step. 

您通过了解请求和限制开始了管理 Kubernetes 集群的旅程，并很快了解到为所有容器建立资源限制是坚实的第一步。

Gathering metrics is essential for calculating resource limits. If you get to the point where you struggle to gather and visualize metrics, ContainIQ may be a good solution. It’s tailored for Kubernetes and can make [monitoring your cluster](https://www.containiq.com/kubernetes-monitoring) much easier.

收集指标对于计算资源限制至关重要。如果您到了难以收集和可视化指标的地步，ContainIQ 可能是一个很好的解决方案。它是为 Kubernetes 量身定制的，可以让 [监控您的集群](https://www.containiq.com/kubernetes-monitoring) 变得更加容易。

Alexandre is a DevOps Engineer at Vosker where he specializes in Complex Systems Engineering and is a Management Specialist. He has embraced DevOps culture since he started his career by contributing to the digital transformation of a leading financial institution in Canada. His passion is the DevOps Revolution and Industrial Engineering. He loves that he has sufficient hindsight to get the best of both worlds. Alexandre has a Master of Applied Science (MASc) in Industrial Engineering from Concordia University.

Alexandre 是 Vosker 的一名 DevOps 工程师，他专攻复杂系统工程，并且是一名管理专家。自从他的职业生涯开始以来，他通过为加拿大一家领先金融机构的数字化转型做出贡献，从而接受了 DevOps 文化。他的热情是 DevOps 革命和工业工程。他喜欢自己有足够的后见之明，可以两全其美。 Alexandre 拥有康考迪亚大学工业工程应用科学硕士学位 (MASc)。

https://www.containiq.com/post/setting-and-rightsizing-kubernetes-resource-limits

