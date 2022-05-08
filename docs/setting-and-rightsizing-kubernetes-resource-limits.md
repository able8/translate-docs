# Setting and Rightsizing Kubernetes Resource Limits \| Best Practices

September 7, 2021

Part of managing a Kubernetes cluster is making sure your clusters aren’t using too many resources. Let’s walk through the concepts of setting and rightsizing resource limits for Kubernetes.



Overview

1. [What Are Resource Limits, and Why Do They Matter?](http://www.containiq.com#1)
2. [Implementing Rightsizing](http://www.containiq.com#2)
3. [Calculate Resource Limits in 3 Steps](http://www.containiq.com#3)
4. [Using ContainIQ for Rightsizing & Resource Limits](http://www.containiq.com#4)
5. [How to Set Resource Limits](http://www.containiq.com#5)
6. [What Are the Next Steps?](http://www.containiq.com#6)
7. [Conclusion](http://www.containiq.com#7)



Managing a Kubernetes cluster is like sitting in front of a large DJ mixing console. You have an almost overwhelming number of knobs and sliders in front of you to tune the sound to perfection. It can seem challenging to even know where to begin—a feeling that Kubernetes engineers probably know well.

However, a [Kubernetes cluster](https://www.containiq.com/post/kubernetes-cluster) without resource limits could conceivably lead to the most issues. So setting resource limits is a logical place to start.

That being said, your fine-tuning challenges are just beginning. To set Kubernetes resources limits correctly, you have to be methodical and take care to find the correct values. Set them too high, and you may negatively affect nodes of your clusters. Set the value too low, and you negatively impact your application performances.

Fortunately, this guide will walk you through all you need to know to properly tune resource limits and keep both your cluster and your applications healthy.

## What Are Resource Limits, and Why Do They Matter?

First of all, in Kubernetes, resources limits come in pairs with resource requests:

- **Resource Requests:** The amount of CPU or memory allocated to your container. A Pod resource request is equal to the sum of its container resource requests. When scheduling Pods, Kubernetes will guarantee that this amount of resource is available for your Pod to run.
- **Resource Limits:** The level at which Kubernetes will start taking action against a container going above the limit. Kubernetes will kill a container consuming too much memory or throttle a container using too much CPU.

If you set resource limits but no resource request, Kubernetes implicitly sets memory and CPU requests equal to the limit. This behavior is excellent as a first step toward getting your Kubernetes cluster under control. This is often referred to as the _conservative approach_, where resources allocated to a container are at maximum.

### Have You Set Requests and Limits on All Your Containers?

If not, the [Kubernetes Scheduler](https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/) will assign any Pods without request and limits randomly. With limits set, you will avoid most of the following problems:

- **Out of Memory (OOM) issues.** A node could die of memory starvation, affecting cluster stability. For instance, an application with [a memory leak could cause an OOM issue](https://www.containiq.com/post/oomkilled-troubleshooting-kubernetes-memory-requests-and-limits).
- **CPU Starvation.** Your applications will get slower because they must share a limited amount of CPU. An application consuming an excessive amount of CPU could affect all applications on the same node.
- **Pod eviction.** When a node lacks resources, it [starts the eviction process](https://www.containiq.com/post/kubernetes-pod-evictions) and terminates Pods, starting with Pods without resource requests.
- **Financial waste.** Assuming your cluster is fine without requests and limits, this means that you are most likely overprovisioning. In other words, you’re spending money on resources you never use.

The first step to prevent these issues is to set limits for all containers.

Imagine your node as a big box with a width (CPU) and length (memory), and you need to fill the big box with smaller boxes (Pods). Not setting requests and limits is similar to playing the game without knowing the width and length of the smaller boxes.

Assuming you only set limits, Kubernetes will check nodes and find a fit for your Pod. Pod assignments will become more coherent. Also, you will have a clear picture of the situation on each node, as illustrated below.

![Kubernetes Cluster with resource limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6233716e5697daf095b21a9f_K8_resource-limits.png)

From there, you can start observing your application more closely and begin optimizing resource requests. You will maximize resource utilizations, making your system more cost-efficient.

## Implementing Rightsizing

What do you need to calculate resource limits? The first thing is metrics!

Resource limits are calculated based on historical data. Kubernetes doesn’t come out of the box with sufficient tools to gather memory and CPU, so here’s a small list of options:

- The[Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server) collects and aggregates metrics. However, you won’t get far with it because it only gives you the current value of a metric.
- [Prometheus](https://prometheus.io/) is a popular solution to monitor anything.
- ContainIQ is a monitoring solution tailored for Kubernetes.

The next thing you need is to identify Pods without limits.

An easy solution is to use <terminal inline>kubectl<terminal inline> and a <terminal inline>go-template<terminal inline>.<terminal inline>go-template<terminal inline> allows you to manipulate data provided by a <terminal inline>kubectl<terminal inline> command and output the desired information.

```c hljs

kubectl get pods --all-namespaces -o go-template-file=./get-pods-without-limits.gotemplate

```

The following <terminal inline>go-template<terminal inline> iterates through all Pods and containers and outputs a container with metrics.

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

![Pods without resource limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6137caa585b79a0e1667afb1_MsKPWHS.png)

Pods without resource limits

‍

Finally, you need to revise Kubernetes deployment, [StatefulSet](https://www.containiq.com/post/kubernetes-statefulsets), or [DaemonSets](https://www.containiq.com/post/using-kubernetes-daemonsets-effectively) for each Pod you found and include resource limits. Unfortunately, when it comes to resource limits, there is no magic formula that fits all cases.

In this activity, you need to use your infrastructure knowledge and challenge the numbers you see. Let me explain—CPU and memory usage is affected by the type of application you are using. For example:

- **Microservices** tackle HTTP requests, consuming resources based on the traffic. You often have many instances of a microservice and have good toleration in terms of failure. Tight resource limits are acceptable and help prevent abnormal behavior.
- **Databases** tend to consume increasingly more memory over time. On top of that, you have no tolerance for failure; tight resource limits are not an option.
- **ETL/ELT** tends to consume resources by burst, but memory and CPU usage are mostly static. Use resource limits to prevent unexpected bursts of resource usage.

Setting up limits leads to fine-tuning other settings of your cluster, such as node size or [auto-scaling](https://www.containiq.com/post/kubernetes-autoscaling). We’ll come back to that later on. For now, let’s focus on a general strategy for calculating limits.

## Calculate Resource Limits in 3 Steps

To carefully analyze the metrics, you need to take it one step at a time. This approach consists of three phases that each lead to a different strategy. Start with the most aggressive, challenge the result, and move on to a more conservative option if necessary.

_You must consider CPU and memory independently and apply different strategies based on your conclusion at each phase._

![3-step approach for calculating resource limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/62337146a566668b8758eb18_K8_calculate-resource-limits.png)

### 1\. Check Memory or CPU

In the first phase, look at the ninety-ninth percentile of memory or CPU. This _aggressive approach_ aims to reduce issues by forcing Kubernetes to take action against outliers. If you set limits with this value, your application will be affected 1 percent of the time. The container CPU will be throttled and never reach that value again.

The aggressive approach is often good for CPU limits because the consequences are relatively acceptable and help you better manage resources. Concerning memory, the ninety-ninth percentile may be problematic; containers restart if the limit is reached.

At this point, you should weigh the consequences and conclude if the ninety-ninth percentile makes sense for you (as a side note, you should investigate deeper as to why applications sometimes reach the set limit). Maybe the ninety-ninth percentile is too restrictive because your application is not yet at its max utilization. In that case, move on to the second strategy to set limits.

### 2\. Check Memory or CPU in a Given Period

In the second phase, you’ll look at the max CPU or memory in a given period. If you set a limit with this value, in theory, no application will be affected. However, it prevents your applications from moving past that limit and keeps your cluster under control.

Once again, you should challenge the value you found. Is the max far from the ninety-ninth percentile found earlier? Are your applications under a lot of traffic? Do you expect your application to handle more load?

At this point, the decision branches out again with two paths. The max makes sense, and applications should be stopped or throttled if they reach that limit. On the other hand, if the maximum is much greater than the ninety-ninth percentile (outlier), or you know you need more room in the future, move to the final option.

### 3\. Find a Compromise

The last stage of this three-step process is to find a compromise based on the maximum by adding or subtracting a coefficient (ie., max + 20%). If you ever reach this point, you should consider performing load tests to characterize your application performances and resource usage better.

Repeat this process for each of your applications without limits.

## Using ContainIQ for Rightsizing & Resource Limits

As discussed above, Kubernetes, by default, does not come with tooling to view pod and node-level metrics over time. Engineering teams often look to third-party tools to deliver both real-time and historical views of cluster performance. ContainIQ, a [Kubernetes monitoring platform](https://www.containiq.com/kubernetes-monitoring), can help engineering teams monitor, track, and alert on core metrics. These tools are quite helpful when it comes to implementing rightsizing and setting resource limits.

### Implementing Rightsizing

Rightsizing nodes is tricky without historical data. Engineering teams often allocate too much and are left with unused resources. This brings peace of mind, but with added cost and likely waste.

Fortunately, ContainIQ collects and stores both real-time and historical node metric data. By clicking on a node, users can view node CPU and memory on the Nodes dashboard:

‍

![ContainIQ Node Conditions](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6241c38b07b0a12908732f56_rightsize_nodes.png)

Nodes are color-coded based on relative usage. And users can use the See Pods tab to see CPU Requests and Limits by pod. Below, users are able to view historical node metrics by node as well as the average across all nodes:

‍

![ContainIQ Node Limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6241c3d50dcc9ddcfbfa5ca0_node_rightsize2.png)

Using the Show Limits toggle, users are able to view the node’s capacity or the amount allocatable alongside historical performance. For example, in the screenshot above, you can see that the given node is continually hitting its CPU allocatable limit. And it might make sense to either increase the size of the node or set up the cluster autoscaler to account for these spikes.

Users are able to more accurately size nodes, which comes with a number of benefits for end-user performance and likely cost savings to the organization.

### Setting Resource Limits

Gathering the right data is an important first step in setting good resource limits. With ContainIQ, users are able to view the CPU and memory for every pod by simply clicking on the given pod:

‍

![ContainIQ Pod Conditions](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6241c3f87d17744bf1168bee_rightsize_pods.png)

In addition, users are able to view in real-time the statuses and recent events for particular pods. Users are also able to view the historical CPU and memory overtime for particular pods, as well as the average across all pods:

‍

![ContainIQ Pod Limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6241c428b23ac908d97217aa_pod_rightsize2.png)

Having this data is invaluable when calculating the necessary resource limits and allows for a more precise calculation. The dropdown allows users to easily change between time periods ranging from hourly to weekly. Using the Show Limits toggle, users are able to see the historical data alongside the limits set at that point in time. If no limits are currently set for the given pod or pods the Show Limits toggle won’t return any data.

When deciding on your resource limits for a given pod or pods, users can use ContainIQ to make accurate limits backed by historical data. And in the future, users can set alerts when a limit is exceeded for a given pod or pods.

You can sign up for ContainIQ [here](https://www.containiq.com/pricing?utm_source=blog&utm_medium=rightsizing&utm_campaign=plug&utm_content=click-here), or [book a demo](https://www.containiq.com/book-a-demo?utm_source=blog&utm_medium=rightsizing&utm_campaign=plug&utm_content=book-a-demo) to learn more.

## How to Set Resource Limits

Once you have calculated CPU and memory limits, it’s time for you to update your configuration and redeploy your application. Limits are defined at the container level as follows:

```yaml hljs

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

I warned you earlier that setting limits isn’t the end of the story. It’s the first step along the way toward refining your Kubernetes cluster management. Considering your resource limits will lead you to reconsidering other settings for your cluster. That’s why I like the metaphor of the mixing console when we think about administering Kubernetes clusters.

In the following diagram, you can see the path of upcoming tasks you’ll need to identify and optimize for efficiency.

![What to do after setting requests and limits](https://assets.website-files.com/5fbfbba70f3f813561ef7b9f/6233710959c0cb5765ec51f0_K8_resource-last-step.png)

To understand the diagram, consider a box as an activity or a setting related to Kubernetes. An arrow linking boxes together represents related tasks and the direction of impact. For instance _Request & Limits_ impacts _node sizing_.

Other labels in the diagram include:

- **Resource Quota:** When dealing with many teams and applications, setting limits can be daunting. Fortunately Kubernetes administrators can set limits at the namespace level.
- **Node Sizing:** Choosing node size is essential to optimize resource usage. But it relies on understanding application needs.
- **Pods and Custer Auto-Scaling**: Resource usage is greatly affected by how many Pods are available to handle the load.
- **Readiness & Liveness**: Properly managing Pod lifecycle can prevent many issues. [For example](https://www.containiq.com/post/kubernetes-readiness-probe), if a Pod is consuming too many resources, it may not be ready to receive traffic.

## Conclusion

You started your journey in managing a Kubernetes cluster by learning about requests and limits and quickly learned that establishing resource limits for all your containers is a solid first step.

Gathering metrics is essential for calculating resource limits. If you get to the point where you struggle to gather and visualize metrics, ContainIQ may be a good solution. It’s tailored for Kubernetes and can make [monitoring your cluster](https://www.containiq.com/kubernetes-monitoring) much easier.

Alexandre is a DevOps Engineer at Vosker where he specializes in Complex Systems Engineering and is a Management Specialist. He has embraced DevOps culture since he started his career by contributing to the digital transformation of a leading financial institution in Canada. His passion is the DevOps Revolution and Industrial Engineering. He loves that he has sufficient hindsight to get the best of both worlds. Alexandre has a Master of Applied Science (MASc) in Industrial Engineering from Concordia University.

[READ MORE](https://containiq.com/author/alexandre-couedelo)

## Related Articles

Looking to learn more? The below posts may be helpful for you to learn more about Kubernetes and our company.

https://www.containiq.com/post/setting-and-rightsizing-kubernetes-resource-limits
