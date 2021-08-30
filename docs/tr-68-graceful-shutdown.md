# Graceful shutdown and zero downtime deployments in Kubernetes

# Kubernetes 中的优雅关闭和零停机部署

Published in August 2020

2020 年 8 月发布

**TL;DR:** *In this article, you will learn how to prevent broken connections when a  Pod starts up or shuts down. You will also learn how to shut down  long-running tasks gracefully.*

**TL;DR:** *在本文中，您将了解如何在 Pod 启动或关闭时防止连接断开。您还将学习如何优雅地关闭长时间运行的任务。*

![Graceful shutdown and zero downtime deployments in Kubernetes](https://learnk8s.io/a/55d21503055aaf9ef8a04d5e595ed505.png)

> You can [download this handy diagram as a PDF here](https://learnk8s.io/a/graceful-shutdown-and-zero-downtime-deployments-in-kubernetes/graceful-shutdown.pdf).

> 您可以[在此处以 PDF 格式下载这个方便的图表](https://learnk8s.io/a/graceful-shutdown-and-zero-downtime-deployments-in-kubernetes/graceful-shutdown.pdf)。

In Kubernetes, creating and deleting Pods is one of the most common tasks.

在 Kubernetes 中，创建和删除 Pod 是最常见的任务之一。

Pods are created when you execute a rolling update, scale deployments, for every new release, for every job and cron job, etc.

当您为每个新版本、每个作业和 cron 作业等执行滚动更新、扩展部署时，会创建 Pod。

But Pods are also deleted and recreated after evictions — when you mark a node as unschedulable for example.

但是 Pod 也会在驱逐后被删除和重新创建——例如，当您将节点标记为不可调度时。

*If the nature of those Pods is so ephemeral, what happens when a Pod is in the middle of responding to a request but it's told to shut down?*

*如果这些 Pod 的性质如此短暂，当 Pod 正在响应请求但被告知关闭时会发生什么？*

*Is the request completed before shutdown?*

*请求是否在关机前完成？*

*What about subsequent requests, are those redirected somewhere else?*

*后续的请求呢，是否重定向到其他地方？*

Before discussing what happens when a Pod is deleted, it's necessary to talk to about what happens when a Pod is created.

在讨论删除 Pod 时会发生什么之前，有必要先谈谈创建 Pod 时会发生什么。

Let's assume you want to create the following Pod in your cluster:

假设您要在集群中创建以下 Pod：

pod.yaml

pod.yaml

```
apiVersion: v1
 kind: Pod
 metadata:
   name: my-pod
 spec:
   containers:
     - name: web
       image: nginx
       ports:
         - name: web
           containerPort: 80
 ```

 
You can submit the YAML definition to the cluster with:

您可以使用以下命令将 YAML 定义提交到集群：

bash

猛击

```
kubectl apply -f pod.yaml
 ```

 
As soon as you enter the command, kubectl submits the Pod definition to the Kubernetes API.

只要您输入命令，kubectl 就会将 Pod 定义提交给 Kubernetes API。

*This is where the journey begins.*

*这是旅程开始的地方。*

## Saving the state of the cluster in the database

##在数据库中保存集群的状态

The Pod definition is received and inspected by the API and subsequently stored in the database — etcd.

Pod 定义由 API 接收和检查，然后存储在数据库中 - etcd。

The Pod is also added to [the Scheduler's queue.](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/#scheduling-cycle-binding-cycle)

Pod 也被添加到 [调度程序的队列。](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/#scheduling-cycle-binding-cycle)

The Scheduler:

调度器：

1. inspects the definition
2. collects details about the workload such as CPU and memory requests and then
3. decides which Node is best suited to run it [(through a process called Filters and Predicates).](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/#extension-points)

1.检查定义
2. 收集有关工作负载的详细信息，例如 CPU 和内存请求，然后
3.决定哪个节点最适合运行它[（通过一个叫做过滤器和谓词的过程）。](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/#extension-points)

At the end of the process:

在过程结束时：

- The Pod is marked as *Scheduled* in etcd.
- The Pod has a Node assigned to it.
- The state of the Pod is stored in etcd.

- Pod 在 etcd 中被标记为 *Scheduled*。
- Pod 分配了一个节点。
- Pod 的状态存储在 etcd 中。

**But the Pod still does not exist.**

**但 Pod 仍然不存在。**

- ![When you submit a Pod with kubectl apply -f the YAML is sent to the Kubernetes API.](https://learnk8s.io/a/54a28f4c41dfd3abb594848af5f71eaf.svg)

  1/3 When you submit a Pod with `kubectl apply -f` the YAML is sent to the Kubernetes API. Next

1/3 当您使用 `kubectl apply -f` 提交 Pod 时，YAML 将发送到 Kubernetes API。下一个

The previous tasks happened in the control plane, and the state is stored in the database.

之前的任务发生在控制平面，状态存储在数据库中。

*So who is creating the Pod in your Nodes?*

*那么谁在你的节点中创建 Pod？*

## The kubelet — the Kubernetes agent

## kubelet — Kubernetes 代理

**It's the kubelet's job to poll the control plane for updates.**

**轮询控制平面以获取更新是 kubelet 的工作。**

You can imagine the kubelet relentlessly asking to the master node: *"I look after the worker Node 1, is there any new Pod for me?".*

你可以想象 kubelet 无情地问主节点：*“我照顾工作节点 1，有没有新的 Pod 给我？”。*

When there is a Pod, the kubelet creates it.

当有 Pod 时，kubelet 会创建它。

*Sort of.*

*有点。*

The kubelet doesn't create the Pod by itself. Instead, it delegates the work to three other components:

kubelet 不会自己创建 Pod。相反，它将工作委托给其他三个组件：

1. **The Container Runtime Interface (CRI)** — the component that creates the containers for the Pod.
2. **The Container Network Interface (CNI)** — the component that connects the containers to the cluster network and assigns IP addresses.
3. **The Container Storage Interface (CSI)** — the component that mounts volumes in your containers.

1. **容器运行时接口 (CRI)** — 为 Pod 创建容器的组件。
2. **容器网络接口 (CNI)** — 将容器连接到集群网络并分配 IP 地址的组件。
3. **容器存储接口 (CSI)** — 在容器中安装卷的组件。

In most cases, the Container Runtime Interface (CRI) is doing a similar job to:

在大多数情况下，容器运行时接口 (CRI) 的作用类似于：

bash

猛击

```
docker run -d <my-container-image>
 ```

 
The Container Networking Interface (CNI) is a bit more interesting because it is in charge of:

容器网络接口（CNI）更有趣，因为它负责：

1. Generating a valid IP address for the Pod.
2. Connecting the container to the rest of the network.

1. 为 Pod 生成一个有效的 IP 地址。
2. 将容器连接到网络的其余部分。

As you can imagine, there are several ways to connect the container to the network and assign a valid IP address (you could choose between IPv4 or IPv6 or maybe assign multiple IP addresses).

可以想象，有多种方法可以将容器连接到网络并分配有效的 IP 地址（您可以在 IPv4 或 IPv6 之间进行选择，也可以分配多个 IP 地址）。

As an example, [Docker creates virtual ethernet pairs and attaches it to a bridge](https://archive.shivam.dev/docker-networking-explained/), whereas [the AWS-CNI connects the Pods directly to the rest of the Virtual Private Cloud (VPC).](https://itnext.io/kubernetes-is-hard-why-eks-makes-it-easier-for-network-and-security-architects-ea6d8b2ca965)

例如，[Docker 创建虚拟以太网对并将其附加到网桥](https://archive.shivam.dev/docker-networking-explained/)，而 [AWS-CNI 将 Pod 直接连接到其余的虚拟私有云 (VPC)。](https://itnext.io/kubernetes-is-hard-why-eks-makes-it-easier-for-network-and-security-architects-ea6d8b2ca965)

When the Container Network Interface finishes its job, the Pod is connected  to the rest of the network and has a valid IP address assigned.

当容器网络接口完成其工作时，Pod 连接到网络的其余部分并分配了一个有效的 IP 地址。

*There's only one issue.*

*只有一个问题。*

**The kubelet knows about the IP address (because it invoked the Container Network Interface), but the control plane does not.** 
**kubelet 知道 IP 地址（因为它调用了容器网络接口），但控制平面不知道。**
No one told the master node that the Pod has an IP address assigned and it's ready to receive traffic.

没有人告诉主节点 Pod 已分配了 IP 地址，并且已准备好接收流量。

As far the control plane is concerned, the Pod is still being created.

就控制平面而言，Pod 仍在创建中。

It's the job of **the kubelet to collect all the details of the Pod such as the IP address and report them back to the control plane.**

**kubelet 的工作是收集 Pod 的所有详细信息，例如 IP 地址，并将它们报告回控制平面。**

You can imagine that inspecting etcd would reveal not just where the Pod is running, but also its IP address.

您可以想象，检查 etcd 不仅会揭示 Pod 的运行位置，还会揭示其 IP 地址。

- ![The Kubelet polls the control plane for updates.](https://learnk8s.io/a/b8bdb7bf659fbea3d2949930093d56b1.svg)

  1/5 The Kubelet polls the control plane for updates. Next

1/5 Kubelet 轮询控制平面以获取更新。下一个

If the Pod isn't part of any Service, this is the end of the journey.

如果 Pod 不是任何服务的一部分，这就是旅程的终点​​。

The Pod is created and ready to use.

Pod 已创建并可以使用。

*When the Pod is part of the Service, there are a few more steps needed.*

*当 Pod 是服务的一部分时，还需要执行一些步骤。*

## Pods and Services

## Pod 和服务

When you create a Service, there are usually two pieces of information that you should pay attention to:

创建 Service 时，通常需要注意两条信息：

1. The selector, which is used to specify the Pods that will receive the traffic.
2. The `targetPort` — the port used by the Pods to receive traffic.

1. 选择器，用于指定将接收流量的 Pod。
2. `targetPort` — Pod 用于接收流量的端口。

A typical YAML definition for the Service looks like this:

服务的典型 YAML 定义如下所示：

service.yaml

服务.yaml

```
apiVersion: v1
 kind: Service
 metadata:
   name: my-service
 spec:
   ports:
   - port: 80
     targetPort: 3000
   selector:
     name: app
 ```

 
When you submit the Service to the cluster with `kubectl apply`, Kubernetes finds all the Pods that have the same label as the selector (`name: app`) and collects their IP addresses — but only if they passed the [Readiness probe] (https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-tcp-liveness-probe).

当您使用 `kubectl apply` 将 Service 提交到集群时，Kubernetes 会找到与选择器（`name: app`）具有相同标签的所有 Pod 并收集它们的 IP 地址——但前提是它们通过了 [Readiness probe] （https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-tcp-liveness-probe）。

Then, for each IP address, it concatenates the IP address and the port.

然后，对于每个 IP 地址，它连接 IP 地址和端口。

If the IP address is `10.0.0.3` and the `targetPort` is `3000`, Kubernetes concatenates the two values and calls them an endpoint.

如果 IP 地址是 `10.0.0.3` 并且 `targetPort` 是 `3000`，Kubernetes 会连接这两个值并将它们称为端点。

```
IP address + port = endpoint
 ---------------------------------
10.0.0.3   + 3000 = 10.0.0.3:3000
 ```

 
The endpoints are stored in etcd in another object called Endpoint.

端点存储在另一个名为 Endpoint 的对象中的 etcd 中。

*Confused?*

*使困惑？*

Kubernetes refers to:

Kubernetes 指的是：

- endpoint (in this article and the Learnk8s material this is referred to as a lowercase `e` endpoint) is the IP address + port pair (`10.0.0.3:3000`).
- Endpoint (in this article and the Learnk8s material this is referred to as an uppercase `E` Endpoint) is a collection of endpoints.

- 端点（在本文和 Learnk8s 材料中称为小写的“e”端点）是 IP 地址 + 端口对（“10.0.0.3:3000”）。
- 端点（在本文和 Learnk8s 材料中，这被称为大写的“E”端点）是端点的集合。

The Endpoint object is a real object in Kubernetes and for every Service Kubernetes automatically creates an Endpoint object.

Endpoint 对象是 Kubernetes 中的一个真实对象，Kubernetes 会为每个服务自动创建一个 Endpoint 对象。

You can verify that with:

您可以通过以下方式验证：

bash

猛击

```
kubectl get services,endpoints
 NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)
 service/my-service-1   ClusterIP   10.105.17.65   <none>        80/TCP
 service/my-service-2   ClusterIP   10.96.0.1      <none>        443/TCP

 NAME                     ENDPOINTS
 endpoints/my-service-1   172.17.0.6:80,172.17.0.7:80
 endpoints/my-service-2   192.168.99.100:8443
 ```

 
The Endpoint collects all the IP addresses and ports from the Pods.

Endpoint 从 Pod 收集所有 IP 地址和端口。

*But not just one time.*

*但不止一次。*

The Endpoint object is refreshed with a new list of endpoints when:

在以下情况下，使用新的端点列表刷新 Endpoint 对象：

1. A Pod is created.
2. A Pod is deleted.
3. A label is modified on the Pod.

1. 创建了一个 Pod。
2. Pod 被删除。
3. Pod 上修改了一个标签。

So you can imagine that every time you create a Pod and after the kubelet  posts its IP address to the master Node, Kubernetes updates all the  endpoints to reflect the change:

因此，您可以想象，每次创建 Pod 时，在 kubelet 将其 IP 地址发布到主节点后，Kubernetes 都会更新所有端点以反映更改：

bash

猛击

```
kubectl get services,endpoints
 NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)
 service/my-service-1   ClusterIP   10.105.17.65   <none>        80/TCP
 service/my-service-2   ClusterIP   10.96.0.1      <none>        443/TCP

 NAME                     ENDPOINTS
 endpoints/my-service-1   172.17.0.6:80,172.17.0.7:80,172.17.0.8:80
 endpoints/my-service-2   192.168.99.100:8443
 ```

 
Great, the endpoint is stored in the control plane, and the Endpoint object was updated.

太好了，端点存储在控制平面中，并且端点对象已更新。

- ![In this picture, there's a single Pod deployed in your cluster.The Pod belongs to a Service. If you were to inspect etcd, you would find the Pod's details as well as Service.](https://learnk8s.io/a/5ec899bd0f7067f1e01bb1accacb35ac.svg)

Pod 属于一个服务。如果您要检查 etcd，您会找到 Pod 的详细信息以及服务。](https://learnk8s.io/a/5ec899bd0f7067f1e01bb1accacb35ac.svg)

  1/8

1/8

  In this picture, there's a single Pod deployed in your cluster. The Pod  belongs to a Service. If you were to inspect etcd, you would find the  Pod's details as well as Service.

在这张图片中，您的集群中部署了一个 Pod。 Pod 属于一个服务。如果您要检查 etcd，您会发现 Pod 的详细信息以及服务。

  Next

下一个

*Are you ready to start using your Pod?*

*您准备好开始使用您的 Pod 了吗？*

**There's more.**

**还有更多。**

A lot more!

多很多！

## Consuming endpoints in Kubernetes

## 在 Kubernetes 中使用端点

**Endpoints are used by several components in Kubernetes.**

**端点由 Kubernetes 中的多个组件使用。**

Kube-proxy uses the endpoints to set up iptables rules on the Nodes.

Kube-proxy 使用端点在节点上设置 iptables 规则。

So every time there is a change to an Endpoint (the object), kube-proxy  retrieves the new list of IP addresses and ports and write the new  iptables rules.

因此，每次端点（对象）发生更改时，kube-proxy 都会检索新的 IP 地址和端口列表并编写新的 iptables 规则。

- ![Let's consider this three-node cluster with two Pods and no Services.The state of the Pods is stored in etcd.](https://learnk8s.io/a/50e4746cbbda956a4550077f2954dd7a.svg)

Pod 的状态存储在 etcd 中。](https://learnk8s.io/a/50e4746cbbda956a4550077f2954dd7a.svg)

  1/6 
1/6
Let's consider this three-node cluster with two Pods and no Services. The state of the Pods is stored in etcd.

让我们考虑这个具有两个 Pod 且没有服务的三节点集群。 Pod 的状态存储在 etcd 中。

  Next

下一个

The Ingress controller uses the same list of endpoints.

Ingress 控制器使用相同的端点列表。

The Ingress controller is that component in the cluster that routes external traffic into the cluster.

Ingress 控制器是集群中将外部流量路由到集群的组件。

When you set up an Ingress manifest you usually specify the Service as the destination:

当您设置 Ingress 清单时，您通常将服务指定为目的地：

ingress.yaml

入口.yaml

```
apiVersion: networking.k8s.io/v1
 kind: Ingress
 metadata:
   name: my-ingress
 spec:
   rules:
   - http:
       paths:
       - backend:
           service:
             name: my-service
             port:
               number: 80
         path: /
         pathType: Prefix
 ```

 
*In reality, the traffic is not routed to the Service.*

*实际上，流量不会路由到服务。*

Instead, the Ingress controller sets up a subscription to be notified every time the endpoints for that Service change.

取而代之的是，入口控制器设置一个订阅，以便在该服务的端点每次更改时收到通知。

**The Ingress routes the traffic directly to the Pods skipping the Service.**

**Ingress 将流量直接路由到跳过 Service 的 Pod。**

As you can imagine, every time there is a change to an Endpoint (the  object), the Ingress retrieves the new list of IP addresses and ports  and reconfigures the controller to include the new Pods.

可以想象，每次 Endpoint（对象）发生更改时，Ingress 都会检索新的 IP 地址和端口列表，并重新配置控制器以包含新的 Pod。

- ![In this picture, there's an Ingress controller with a Deployment with two replicas and a Service.](https://learnk8s.io/a/175161d4127b2daf602b2a240190771d.svg)

  1/9

1/9

  In this picture, there's an Ingress controller with a Deployment with two replicas and a Service.

在这张图片中，有一个 Ingress 控制器和一个带有两个副本和一个服务的部署。

  Next

下一个

There are more examples of Kubernetes components that subscribe to changes to endpoints.

还有更多订阅端点更改的 Kubernetes 组件示例。

CoreDNS, the DNS component in the cluster, is another example.

集群中的 DNS 组件 CoreDNS 是另一个例子。

If you use [Services of type Headless](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services), CoreDNS will have to subscribe to changes to the endpoints and reconfigure itself every time an endpoint is added or removed.

如果您使用 [Headless 类型的服务](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services)，CoreDNS 将不得不订阅端点的更改并在每次端点被添加或删除。

The same endpoints are consumed by service meshes such as Istio or Linkerd, [by cloud providers to create Services of `type:LoadBalancer`](https://thebsdbox.co.uk/2020/03/18/Creating-a-Kubernetes -cloud-doesn-t-required-boiling-the-ocean/) and countless operators.

相同的端点被 Istio 或 Linkerd 等服务网格使用，[由云提供商创建`type:LoadBalancer` 的服务](https://thebsdbox.co.uk/2020/03/18/Creating-a-Kubernetes -cloud-doesn-t-required-boiling-the-ocean/) 和无数的运营商。

You must remember that several components subscribe to change to endpoints  and they might receive notifications about endpoint updates at different times.

您必须记住，多个组件订阅端点的更改，并且它们可能会在不同时间收到有关端点更新的通知。

*Is it enough, or is there something happening after you create a Pod?*

*是否足够，或者在创建 Pod 后发生了什么？*

**This time you're done!**

**这次你完成了！**

A quick recap on what happens when you create a Pod:

快速回顾一下创建 Pod 时会发生什么：

1. The Pod is stored in etcd.
2. The scheduler assigns a Node. It writes the node in etcd.
3. The kubelet is notified of a new and scheduled Pod.
4. The kubelet delegates creating the container to the Container Runtime Interface (CRI).
5. The kubelet delegates attaching the container to the Container Network Interface (CNI).
6. The kubelet delegates mounting volumes in the container to the Container Storage Interface (CSI).
7. The Container Network Interface assigns an IP address.
8. The kubelet reports the IP address to the control plane.
9. The IP address is stored in etcd.

1. Pod 存储在 etcd 中。
2.调度器分配一个节点。它将节点写入 etcd。
3. kubelet 被通知有一个新的和预定的 Pod。
4. kubelet 将容器的创建委托给容器运行时接口 (CRI)。
5. kubelet 委托将容器附加到容器网络接口 (CNI)。
6. kubelet 将容器中的挂载卷委托给容器存储接口 (CSI)。
7.容器网络接口分配一个IP地址。
8. kubelet 将 IP 地址上报给控制平面。
9、IP地址存储在etcd中。

And if your Pod belongs to a Service:

如果您的 Pod 属于服务：

1. The kubelet waits for a successful Readiness probe.
2. All relevant Endpoints (objects) are notified of the change.
3. The Endpoints add a new endpoint (IP address + port pair) to their list.
4. Kube-proxy is notified of the Endpoint change. Kube-proxy updates the iptables rules on every node.
5. The Ingress controller is notified of the Endpoint change. The controller routes traffic to the new IP addresses.
6. CoreDNS is notified of the Endpoint change. If the Service is of type Headless, the DNS entry is updated.
7. The cloud provider is notified of the Endpoint change. If the Service is of `type: LoadBalancer`, the new endpoint are configured as part of the load balancer pool.
8. Any service mesh installed in the cluster is notified of the Endpoint change.
9. Any other operator subscribed to Endpoint changes is notified too.

1. kubelet 等待成功的 Readiness 探测。
2. 通知所有相关端点（对象）更改。
3. 端点将新端点（IP 地址 + 端口对）添加到其列表中。
4. Kube-proxy 收到 Endpoint 变化的通知。 Kube-proxy 更新每个节点上的 iptables 规则。
5. 入口控制器收到端点变化的通知。控制器将流量路由到新的 IP 地址。
6. CoreDNS 收到 Endpoint 更改的通知。如果服务的类型为 Headless，则更新 DNS 条目。
7. 将端点更改通知云提供商。如果 Service 的类型为“类型：LoadBalancer”，则新端点将配置为负载均衡器池的一部分。
8. 集群中安装的任何服务网格都会收到端点更改的通知。
9. 订阅端点更改的任何其他运营商也会收到通知。

*Such a long list for what is surprisingly a common task — creating a Pod.*

*对于一项令人惊讶的常见任务 - 创建 Pod 而言，这是一个很长的清单。*

The Pod is *Running*. It is time to discuss what happens when you delete it.

Pod 正在*运行*。现在是讨论删除它时会发生什么的时候了。

## Deleting a Pod

## 删除一个 Pod

You might have guessed it already, but when the Pod is deleted, you have to follow the same steps but in reverse.

您可能已经猜到了，但是当 Pod 被删除时，您必须遵循相同的步骤，但相反。

First, the endpoint should be removed from the Endpoint (object).

首先，端点应该从端点（对象）中移除。

This time the Readiness probe is ignored, and the endpoint is removed immediately from the control plane.

这次准备就绪探测被忽略，端点立即从控制平面中删除。

That, in turn, fires off all the events to kube-proxy, Ingress controller, DNS, service mesh, etc.

这反过来将所有事件触发到 kube-proxy、Ingress 控制器、DNS、服务网格等。

Those components will update their internal state and stop routing traffic to the IP address.

这些组件将更新其内部状态并停止将流量路由到 IP 地址。

Since the components might be busy doing something else, **there is no guarantee on how long it will take to remove the IP address from their internal state.**

由于组件可能正忙于做其他事情，**无法保证从其内部状态中删除 IP 地址需要多长时间。**

For some, it could take less than a second; for others, it could take more. 
对于某些人来说，可能只需要不到一秒钟；对于其他人，可能需要更多。
- ![If you're deleting a Pod with kubectl delete pod, the command reaches the Kubernetes API first.](https://learnk8s.io/a/336567c1d80c8853cb6900bdf6fd30d9.svg)

  1/5

1/5

  If you're deleting a Pod with `kubectl delete pod`, the command reaches the Kubernetes API first.

如果您使用 `kubectl delete pod` 删除 Pod，该命令首先到达 Kubernetes API。

  Next

下一个

At the same time, the status of the Pod in etcd is changed to *Terminating*.

同时，将 Pod 在 etcd 中的状态更改为 *Terminating*。

The kubelet is notified of the change and delegates:

kubelet 会收到更改和委托的通知：

1. Unmounting any volumes from the container to the Container Storage Interface (CSI).
2. Detaching the container from the network and releasing the IP address to the Container Network Interface (CNI).
3. Destroying the container to the Container Runtime Interface (CRI).

1. 将任何卷从容器卸载到容器存储接口 (CSI)。
2. 将容器从网络中分离，并将 IP 地址释放到容器网络接口（CNI）。
3. 销毁容器到容器运行时接口（CRI）。

In other words, Kubernetes follows precisely the same steps to create a Pod but in reverse.

换句话说，Kubernetes 遵循与创建 Pod 完全相同的步骤，但相反。

- ![If you're deleting a Pod with kubectl delete pod, the command reaches the Kubernetes API first.](https://learnk8s.io/a/dbfd8be2cd6fbc3984dbb12cfece608d.svg)

  1/3

1/3

  If you're deleting a Pod with `kubectl delete pod`, the command reaches the Kubernetes API first.

如果您使用 `kubectl delete pod` 删除 Pod，该命令首先到达 Kubernetes API。

  Next

下一个

However, there is a subtle but essential difference.

然而，有一个微妙但本质的区别。

**When you terminate a Pod, removing the endpoint and the signal to the kubelet are issued at the same time.**

**当你终止一个 Pod 时，删除端点和发送到 kubelet 的信号是同时发出的。**

When you create a Pod for the first time, Kubernetes waits for the kubelet  to report the IP address and then kicks off the endpoint propagation.

当您第一次创建 Pod 时，Kubernetes 会等待 kubelet 报告 IP 地址，然后开始端点传播。

**However, when you delete a Pod, the events start in parallel.**

**但是，当您删除 Pod 时，事件会并行启动。**

And this could cause quite a few race conditions.

这可能会导致相当多的竞争条件。

*What if the Pod is deleted before the endpoint is propagated?*

*如果在传播端点之前删除 Pod 会怎样？*

- ![Deleting the endpoint and deleting the Pod happen at the same time.](https://learnk8s.io/a/e17a49eb08f03f2c2ff02b91409314fb.svg)

  1/3

1/3

  Deleting the endpoint and deleting the Pod happen at the same time.

删除端点和删除 Pod 是同时发生的。

  Next

下一个

## Graceful shutdown

## 优雅关机

When a Pod is terminated before the endpoint is removed from kube-proxy or  the Ingress controller, you might experience downtime.

当 Pod 在端点从 kube-proxy 或 Ingress 控制器中删除之前终止时，您可能会遇到停机时间。

And, if you think about it, it makes sense.

而且，如果你仔细想想，这是有道理的。

Kubernetes is still routing traffic to the IP address, but the Pod is no longer there.

Kubernetes 仍在将流量路由到 IP 地址，但 Pod 不再存在。

The Ingress controller, kube-proxy, CoreDNS, etc. didn't have enough time to remove the IP address from their internal state.

Ingress 控制器、kube-proxy、CoreDNS 等没有足够的时间从其内部状态中删除 IP 地址。

Ideally, Kubernetes should wait for all components in the cluster to have an  updated list of endpoints before the Pod is deleted.

理想情况下，Kubernetes 应该在删除 Pod 之前等待集群中的所有组件都有更新的端点列表。

*But Kubernetes doesn't work like that.*

*但 Kubernetes 不是这样工作的。*

Kubernetes offers robust primitives to distribute the endpoints (i.e. the Endpoint object and more advanced abstractions such as [Endpoint Slices](https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/)).

Kubernetes 提供了强大的原语来分发端点（即端点对象和更高级的抽象，例如 [端点切片](https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/)）。

However, Kubernetes does not verify that the components that subscribe to  endpoints changes are up-to-date with the state of the cluster.

但是，Kubernetes 不会验证订阅端点更改的组件是否与集群状态保持同步。

*So what can you do avoid this race conditions and make sure that the Pod is deleted after the endpoint is propagated?*

*那么你能做些什么来避免这种竞争条件并确保在端点传播后删除 Pod？*

**You should wait.**

**你应该等待。**

**When the Pod is about to be deleted, it receives a SIGTERM signal.**

**当 Pod 即将被删除时，它会收到一个 SIGTERM 信号。**

Your application can capture that signal and start shutting down.

您的应用程序可以捕获该信号并开始关闭。

Since it's unlikely that the endpoint is immediately deleted from all components in Kubernetes, you could:

由于不太可能立即从 Kubernetes 的所有组件中删除端点，您可以：

1. Wait a bit longer before exiting.
2. Still process incoming traffic, despite the SIGTERM.
3. Finally, close existing long-lived connections (perhaps a database connection or WebSockets).
4. Close the process.

1. 稍等片刻再退出。
2. 仍然处理传入的流量，尽管有 SIGTERM。
3. 最后，关闭现有的长期连接（可能是数据库连接或 WebSockets）。
4. 关闭进程。

*How long should you wait?*

*你要等多久？*

**By default, Kubernetes will send the SIGTERM signal and wait 30 seconds before force killing the process.**

**默认情况下，Kubernetes 会发送 SIGTERM 信号并等待 30 秒，然后再强制终止进程。**

So you could use the first 15 seconds to continue operating as nothing happened.

因此，您可以使用前 15 秒继续操作，因为什么都没有发生。

Hopefully, the interval should be enough to propagate the endpoint removal to kube-proxy, Ingress controller, CoreDNS, etc.

希望间隔应该足以将端点删除传播到 kube-proxy、Ingress 控制器、CoreDNS 等。

And, as a consequence, less and less traffic will reach your Pod until it stops.

因此，到达您的 Pod 的流量会越来越少，直到它停止。

After the 15 seconds, it's safe to close your connection with the database  (or any persistent connections) and terminate the process.

15 秒后，关闭与数据库的连接（或任何持久连接）并终止进程是安全的。

If you think you need more time, you can stop the process at 20 or 25 seconds.

如果您认为需要更多时间，可以在 20 或 25 秒后停止该过程。

However, you should remember that Kubernetes will forcefully kill the process after 30 seconds [(unless you change the `terminationGracePeriodSeconds` in your Pod definition).](https://kubernetes.io/docs/concepts/containers/container-lifecycle- hooks/#hook-handler-execution)

但是，您应该记住 Kubernetes 会在 30 秒后强行终止进程 [（除非您更改 Pod 定义中的 `terminationGracePeriodSeconds`）。](https://kubernetes.io/docs/concepts/containers/container-lifecycle-钩子/#hook-handler-execution)

*What if you can't change the code to wait longer?*

*如果您无法更改代码以等待更长时间怎么办？*

You could invoke a script to wait for a fixed amount of time and then let the app exit.

您可以调用一个脚本等待一段固定的时间，然后让应用程序退出。

Before the SIGTERM is invoked, Kubernetes exposes a `preStop` hook in the Pod.

在调用 SIGTERM 之前，Kubernetes 在 Pod 中公开了一个 `preStop` 钩子。

You could set the `preStop` to hook to wait for 15 seconds.

您可以将 `preStop` 设置为挂钩以等待 15 秒。

Let's have a look at an example:

让我们看一个例子：

pod.yaml

pod.yaml

```
apiVersion: v1
 kind: Pod
 metadata:
   name: my-pod
 spec:
   containers:
     - name: web
       image: nginx
       ports:
         - name: web
           containerPort: 80
       lifecycle:
         preStop:
           exec:
             command: ["sleep", "15"]
 ``` 


The `preStop` hook is one of the [Pod LifeCycle hooks](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/).

`preStop` 钩子是 [Pod LifeCycle 钩子](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/) 之一。

*Is a 15 seconds delay the recommended amount?*

*延迟 15 秒是推荐量吗？*

It depends, but it could be a sensible way to start testing.

这取决于，但这可能是开始测试的明智方法。

Here's a recap of what options you have:

以下是您有哪些选择的概述：

- ![You already know that, when a Pod is deleted, the kubelet is notified of the change.](https://learnk8s.io/a/bdaa1da0be0fa3e9fe022cbf2b22bd1d.svg)

  1/5

1/5

  You already know that, when a Pod is deleted, the kubelet is notified of the change.

您已经知道，当 Pod 被删除时，kubelet 会收到更改通知。

  Next

下一个

## Grace periods and rolling updates

## 宽限期和滚动更新

Graceful shutdown applies to Pods being deleted.

正常关闭适用于被删除的 Pod。

*But what if you don't delete the Pods?*

*但是如果您不删除 Pod 呢？*

Even if you don't, Kubernetes deletes Pods all the times.

即使您不这样做，Kubernetes 也会一直删除 Pod。

In particular, Kubernetes creates and deletes Pods every time you deploy a newer version of your application.

特别是，每当您部署新版本的应用程序时，Kubernetes 都会创建和删除 Pod。

When you change the image in your Deployment, Kubernetes rolls out the change incrementally.

当您更改 Deployment 中的映像时，Kubernetes 会逐步推出更改。

pod.yaml

pod.yaml

```
apiVersion: apps/v1
 kind: Deployment
 metadata:
   name: app
 spec:
   replicas: 3
   selector:
     matchLabels:
       name: app
   template:
     metadata:
       labels:
         name: app
     spec:
       containers:
       - name: app
         # image: nginx:1.18 OLD
         image: nginx:1.19
         ports:
           - containerPort: 3000
 ```

 
If you have three replicas and as soon as you submit the new YAML resources Kubernetes:

如果您有三个副本，并且一旦您提交新的 YAML 资源 Kubernetes：

- Creates a Pod with the new container image.
- Destroys an existing Pod.
- Waits for the Pod to be ready.

- 使用新的容器镜像创建一个 Pod。
- 摧毁一个现有的 Pod。
- 等待 Pod 准备就绪。

And it repeats the steps above until all the Pods are migrated to the newer version.

并重复上述步骤，直到所有 Pod 都迁移到较新版本。

Kubernetes repeats each cycle only after the new Pod is ready to receive traffic  (in other words, it passes the Readiness check).

只有在新 Pod 准备好接收流量（换句话说，它通过了就绪检查）之后，Kubernetes 才会重复每个周期。

*Does Kubernetes wait for the Pod to be deleted before moving to the next one?*

*Kubernetes 是否会在移动到下一个之前等待 Pod 被删除？*

**No.**

**不。**

If you have 10 Pods and the Pod takes 2 seconds to be ready and 20 to shut down this is what happens:

如果您有 10 个 Pod，并且 Pod 需要 2 秒才能准备好，需要 20 秒才能关闭，则会发生以下情况：

1. The first Pod is created, and a previous Pod is terminated.
2. The new Pod takes 2 seconds to be ready after that Kubernetes creates a new one.
3. In the meantime, the Pod being terminated stays terminating for 20 seconds

1.第一个Pod被创建，前一个Pod被终止。
2. Kubernetes 创建一个新 Pod 后，新 Pod 需要 2 秒才能准备好。
3. 同时，被终止的 Pod 保持终止状态 20 秒

After 20 seconds, all new Pods are live (10 Pods, *Ready* after 2 seconds) and all 10 the previous Pods are terminating (the first *Terminated* Pod is about to exit).

20 秒后，所有新 Pod 都处于活动状态（10 个 Pod，2 秒后 *Ready*）并且所有 10 个之前的 Pod 都将终止（第一个 *Terminated* Pod 即将退出）。

In total, you have double the amount of Pods for a short amount of time (10 *Running*, 10 *Terminating*).

总的来说，您在短时间内拥有双倍数量的 Pod（10 *运行*，10 *终止*）。

![Rolling update and graceful shutdown](https://learnk8s.io/a/7934a67ff44a183254acf81a763e6c2f.svg)

The longer the graceful period compared to the Readiness probe, the more Pods you will have *Running* (and *Terminating*) at the same time.

与 Readiness 探针相比，宽限期越长，您同时*运行*（和*终止*）的 Pod 就越多。

*Is it bad?*

*不好吗？*

Not necessarily since you're careful not dropping connections.

不一定，因为您小心不要断开连接。

## Terminating long-running tasks

## 终止长时间运行的任务

*And what about long-running jobs?*

*长期运行的工作呢？*

*If you are transcoding a large video, is there any way to delay stopping the Pod?*

*如果您正在转码大型视频，是否有任何方法可以延迟停止 Pod？*

Imagine you have a Deployment with three replicas.

假设您有一个包含三个副本的 Deployment。

Each replica is assigned a video to transcode, and the task could take several hours to complete.

每个副本都分配了一个视频进行转码，该任务可能需要几个小时才能完成。

When you trigger a rolling update, the Pod has 30 seconds to complete the task before it's killed.

当您触发滚动更新时，Pod 在被杀死之前有 30 秒的时间来完成任务。

*How can you avoid delaying shutting down the Pod?*

*如何避免延迟关闭 Pod？*

You could increase the `terminationGracePeriodSeconds` to a couple of hours.

您可以将“terminationGracePeriodSeconds”增加到几个小时。

**However, the endpoint of the Pod is unreachable at that point.**

**但是，此时 Pod 的端点无法访问。**

![Unreachable Pod](https://learnk8s.io/a/f6777b220dd43ad627b02073c2ffde30.svg)

If you expose metrics to monitor your Pod, your instrumentation won't be able to reach your Pod.

如果您公开指标来监控您的 Pod，您的检测将无法访问您的 Pod。

*Why?*

*为什么？*

**Tools such as Prometheus rely on Endpoints to scrape Pods in your cluster.**

**Prometheus 等工具依赖 Endpoints 来抓取集群中的 Pod。**

However, as soon as you delete the Pod, the endpoint deletion is propagated in the cluster — even to Prometheus!

但是，一旦您删除 Pod，端点删除就会在集群中传播——甚至传播到 Prometheus！

**Instead of increasing the grace period, you should consider creating a new Deployment for every new release.**

**与其增加宽限期，不如考虑为每个新版本创建一个新部署。**

When you create a brand new Deployment, the existing Deployment is left untouched.

当您创建一个全新的 Deployment 时，现有的 Deployment 保持不变。

The long-running jobs can continue processing the video as usual.

长时间运行的作业可以像往常一样继续处理视频。

Once they are done, you can delete them manually.

完成后，您可以手动删除它们。

If you wish to delete them automatically, you might want to set up an  autoscaler that can scale your deployment to zero replicas when they run out of tasks.

如果您希望自动删除它们，您可能需要设置一个自动缩放器，当它们用完任务时，它可以将您的部署扩展到零个副本。

An example of such Pod autoscaler is [Osiris — a general-purpose, scale-to-zero component for Kubernetes.](https://github.com/deislabs/osiris)

此类 Pod 自动缩放器的一个示例是 [Osiris — Kubernetes 的通用、缩放到零的组件。](https://github.com/deislabs/osiris)

The technique is sometimes referred to as **Rainbow Deployments** and is useful every time you have to keep the previous Pods *Running* for longer than the grace period.

该技术有时被称为 **Rainbow 部署**，每当您必须让之前的 Pod 保持*运行*超过宽限期时，它都会很有用。

*Another excellent example is WebSockets.*

*另一个很好的例子是 WebSockets。*

If you are streaming real-time updates to your users, you might not want  to terminate the WebSockets every time there is a release.

如果您正在向用户流式传输实时更新，您可能不想在每次发布时终止 WebSocket。

If you are frequently releasing during the day, that could lead to several interruptions to real-time feeds.

如果您在白天频繁发布，这可能会导致实时提要的多次中断。

**Creating a new Deployment for every release is a less obvious but better choice.** 
**为每个版本创建一个新的 Deployment 是一个不太明显但更好的选择。**
Existing users can continue streaming updates while the most recent Deployment serves the new users.

现有用户可以继续流式传输更新，而最近的部署为新用户提供服务。

As a user disconnects from old Pods, you can gradually decrease the replicas and retire past Deployments.

当用户与旧 Pod 断开连接时，您可以逐渐减少副本并退出过去的部署。

## Summary

＃＃ 概括

You should pay attention to Pods being deleted from your cluster since their IP address might be still used to route traffic.

您应该注意 Pod 从集群中删除，因为它们的 IP 地址可能仍用于路由流量。

Instead of immediately shutting down your Pods, you should consider waiting a little bit longer in your application or set up a `preStop` hook.

与其立即关闭你的 Pod，你应该考虑在你的应用程序中等待更长时间，或者设置一个 `preStop` 钩子。

The Pod should be removed only after all the endpoints in the cluster are  propagated and removed from kube-proxy, Ingress controllers, CoreDNS,  etc.

只有在集群中的所有端点都从 kube-proxy、Ingress 控制器、CoreDNS 等中传播和删除后，才应删除 Pod。

If your Pods run long-lived  tasks such as transcoding videos or serving real-time updates with  WebSockets, you should consider using rainbow deployments.

如果您的 Pod 运行长期任务，例如转码视频或使用 WebSockets 提供实时更新，您应该考虑使用彩虹部署。

In rainbow deployments, you create a new Deployment for every release and  delete the previous one when the connection (or the tasks) drained.

在彩虹部署中，您为每个版本创建一个新部署，并在连接（或任务）耗尽时删除前一个部署。

You can manually remove the older deployments as soon as the long-running task is completed.

您可以在长时间运行的任务完成后立即手动删除较旧的部署。

Or you could automatically scale your deployment to zero replicas to automate the process. 
或者您可以自动将您的部署扩展到零副本以自动化该过程。
