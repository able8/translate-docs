# Gracefully Shutting Down Pods in a Kubernetes Cluster

# 优雅地关闭 Kubernetes 集群中的 Pod

Gracefully Stopping Containers in Kubernetes

在 Kubernetes 中优雅地停止容器

This is part 2 of [our journey](https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33) to implementing a zero downtime update of our Kubernetes cluster. In [part 1 of the series](https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33), we laid out the problem and the challenges of naively draining our nodes in the cluster. In this post, we will cover how to tackle one of  those problems: gracefully shutting down the Pods.

这是 [我们的旅程](https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33) 的第 2 部分，以实现我们的 Kubernetes 集群的零停机更新。在 [系列的第 1 部分](https://blog.gruntwork.io/zero-downtime-server-updates-for-your-kubernetes-cluster-902009df5b33) 中，我们列出了天真地耗尽我们的资源的问题和挑战集群中的节点。在这篇文章中，我们将介绍如何解决其中一个问题：优雅地关闭 Pod。

## Pod Eviction Lifecycle

## Pod 驱逐生命周期

By default, `kubectl drain` will evict pods in a way to honor the pod lifecycle. What this means in practice is that it will respect the following flow:

默认情况下，`kubectl drain` 将以一种尊重 Pod 生命周期的方式驱逐 Pod。这在实践中意味着它将遵循以下流程：

- `drain` will issue a request to delete the pods on the target node to the control plane. This will subsequently notify the `kubelet` on the target node to start shutting down the pods.
 - The `kubelet` on the node will invoke the `preStop` hook in the pod.
 - Once the `preStop` hook completes, the `kubelet` on the node will issue the `TERM` signal to the running application in the containers of the pod.
 - The `kubelet` on the node will wait for up to the grace period (specified on the pod, or passed in from the command line; defaults to 30 seconds) for the  containers to shut down, before forcibly killing the process (with ` SIGKILL`). Note that this grace period includes the time to execute the `preStop` hook.

- `drain` 将向控制平面发出删除目标节点上的 pod 的请求。这将随后通知目标节点上的 `kubelet` 开始关闭 pod。
- 节点上的 `kubelet` 将调用 pod 中的 `preStop` 钩子。
- 一旦 `preStop` 钩子完成，节点上的 `kubelet` 将向 pod 容器中正在运行的应用程序发出 `TERM` 信号。
- 节点上的 `kubelet` 将等待容器关闭的宽限期（在 pod 上指定，或从命令行传入；默认为 30 秒），然后强制终止进程（使用 ` SIGKILL`）。请注意，此宽限期包括执行 `preStop` 钩子的时间。

Based on this flow, you can leverage `preStop` hooks and signal handling in your application pods to gracefully  shutdown your application so that it can "clean up" before it is  ultimately terminated. For example, if you have a worker process  streaming tasks from a queue, you can have your application trap the `TERM` signal to indicate that the application should stop accepting new work, and stop running after all current work has finished. Or, if you are  running an application which can't be modified to trap the `TERM` signal (a third party app for example), then you can use the `preStop` hook to implement the custom API that the service provides for graceful shut down of the application.

基于此流程，您可以利用应用程序 pod 中的 `preStop` 钩子和信号处理来正常关闭应用程序，以便它可以在最终终止之前“清理”。例如，如果您有一个工作进程从队列中流式传输任务，您可以让您的应用程序捕获“TERM”信号以指示应用程序应停止接受新工作，并在所有当前工作完成后停止运行。或者，如果您正在运行一个无法修改以捕获“TERM”信号的应用程序（例如第三方应用程序），那么您可以使用“preStop”钩子来实现该服务为优雅而提供的自定义 API关闭应用程序。

In our example, Nginx does not gracefully handle the `TERM` signal by default, causing existing requests being serviced to fail. Therefore, we will instead rely on a `preStop` hook to gracefully stop Nginx. We will modify our resource to add a `lifecycle` clause to the container spec. The `lifecycle` clause looks like this:

在我们的示例中，默认情况下 Nginx 不会优雅地处理“TERM”信号，从而导致正在处理的现有请求失败。因此，我们将改为依靠 `preStop` 钩子来优雅地停止 Nginx。我们将修改我们的资源以在容器规范中添加一个 `lifecycle` 子句。 `lifecycle` 子句如下所示：

```
 lifecycle:
   preStop:
     exec:
       command: [
         # Gracefully shutdown nginx
         "/usr/sbin/nginx", "-s", "quit"
       ]
 ```

 
With this config, the shutdown sequence will issue the command `/usr/sbin/nginx -s quit` before sending `SIGTERM` to the Nginx process in the container. Note that since the command will gracefully stop the Nginx process and the pod, the `TERM` signal essentially becomes a noop.

使用此配置，关闭序列将在将“SIGTERM”发送到容器中的 Nginx 进程之前发出命令“/usr/sbin/nginx -s quit”。请注意，由于该命令将优雅地停止 Nginx 进程和 pod，因此“TERM”信号本质上变成了一个 noop。

This should be nested under the Nginx container spec. When we include this, the full config for the `Deployment` looks as follows:

这应该嵌套在 Nginx 容器规范下。当我们包含它时，“部署”的完整配置如下所示：

```
 ---
 apiVersion: apps/v1
 kind: Deployment
 metadata:
   name: nginx-deployment
   labels:
     app: nginx
 spec:
   replicas: 2
   selector:
     matchLabels:
       app: nginx
   template:
     metadata:
       labels:
         app: nginx
     spec:
       containers:
       - name: nginx
         image: nginx:1.15
         ports:
         - containerPort: 80
         lifecycle:
           preStop:
             exec:
               command: [
                 # Gracefully shutdown nginx
                 "/usr/sbin/nginx", "-s", "quit"
               ]
 ```

 
## Continuous Traffic After Shutdown

## 关机后持续流量

The graceful shutdown of the Pod ensures Nginx is stopped in a way to  service the existing traffic before shutting down. However, you may  observe that despite best intentions, the Nginx container continues to  receive traffic after shutting down, causing downtime in your service.

Pod 的正常关闭可确保 Nginx 在关闭之前以一种为现有流量提供服务的方式停止。但是，您可能会观察到，尽管出于最好的意图，Nginx 容器在关闭后仍会继续接收流量，从而导致您的服务停机。

To see how this can be problematic, let’s walk through an example with our sample deployment. For the sake of this example, we will assume that  the node had received traffic from a client. This will spawn a worker  thread in the application to service the request. We will indicate this  thread with the circle on the pod container.

要了解这会如​​何产生问题，让我们通过示例部署来演练一个示例。在本示例中，我们假设节点已收到来自客户端的流量。这将在应用程序中产生一个工作线程来为请求提供服务。我们将用 pod 容器上的圆圈表示该线程。

![img](https://miro.medium.com/max/1400/1*1rT_u8Csg87rvbMmxa0e6w.png)

Suppose that at this point, a cluster operator decides to perform maintenance  on Node 1. As part of this, the operator runs the command `kubectl drain node-1` , causing the `kubelet` process on the node to execute the `preStop` hook, starting a graceful shutdown of the Nginx process:

假设此时，集群运营商决定在节点 1 上执行维护。作为其中的一部分，运营商运行命令 `kubectl drain node-1`，导致节点上的 `kubelet` 进程执行 `preStop`钩子，开始优雅地关闭 Nginx 进程：

![img](https://miro.medium.com/max/1400/1*blRUsNTnS0n6nCc6ZVbUHw.png) 

Because nginx is still servicing the original request, it does not immediately  terminate. However, when nginx starts a graceful shutdown, it errors and rejects additional traffic that comes to it.

由于 nginx 仍在为原始请求提供服务，因此它不会立即终止。然而，当 nginx 开始正常关闭时，它会出错并拒绝额外的流量。

At this point, suppose a new server request comes into our service. Since  the pod is still registered with the service, the pod can still receive  the traffic. If it does, this will return an error because the nginx  server is shutting down:

此时，假设一个新的服务器请求进入我们的服务。由于 pod 仍然在服务中注册，因此 pod 仍然可以接收流量。如果是这样，这将返回一个错误，因为 nginx 服务器正在关闭：

![img](https://miro.medium.com/max/1400/1*XmLMv3Df8lAci85HvGoPSA.png)

To complete the sequence, eventually nginx will finish processing the  original request, which will terminate the pod and the node will finish  draining:

为了完成序列，最终 nginx 将完成对原始请求的处理，这将终止 pod，节点将完成排空：

![img](https://miro.medium.com/max/1400/1*jO1zMky9e7sCoej7d3-JvQ.png)

![img](https://miro.medium.com/max/1400/1*jfdYZ0PGg3oGn5Wdak4XsQ.png)

In this example, when the application pod receives the traffic after the  shutdown sequence is initiate, the first client will receive a response  from the server. However, the second client receives an error, which  will be perceived downtime.

在这个例子中，当应用程序 pod 在关闭序列启动后收到流量时，第一个客户端将收到来自服务器的响应。但是，第二个客户端收到错误，这将被视为停机。

So why does this happen? And how do you mitigate potential downtime for  clients that end up connecting to the server during a shutdown sequence? In [the next part of our series](https://blog.gruntwork.io/delaying-shutdown-to-wait-for-pod-deletion-propagation-445f779a8304), we will cover the pod eviction lifecylce in more details and describe how you can introduce a delay in the `preStop` hook to mitigate the effects of continuous traffic from the `Service` .

那么为什么会发生这种情况呢？以及如何减少在关闭序列期间最终连接到服务器的客户端的潜在停机时间？在 [我们系列的下一部分](https://blog.gruntwork.io/delaying-shutdown-to-wait-for-pod-deletion-propagation-445f779a8304) 中，我们将更详细地介绍 pod eviction lifecylce 和描述如何在 `preStop` 钩子中引入延迟以减轻来自 `Service` 的持续流量的影响。

*To get a fully implemented version of zero downtime Kubernetes cluster updates on AWS and more, check out* [*Gruntwork.io*](http://gruntwork.io).

*要在 AWS 上获得完全实施的零停机 Kubernetes 集群更新版本等，请查看* [*Gruntwork.io*](http://gruntwork.io)。

The Gruntwork Blog 
Gruntwork 博客
