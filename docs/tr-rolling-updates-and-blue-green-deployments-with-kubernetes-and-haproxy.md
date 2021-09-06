# Rolling Updates and Blue-Green Deployments with Kubernetes and HAProxy

# 使用 Kubernetes 和 HAProxy 滚动更新和蓝绿部署

Feb 11, 2020 From: https://www.haproxy.com/blog/rolling-updates-and-blue-green-deployments-with-kubernetes-and-haproxy/

_The HAProxy Kubernetes Ingress Controller supports two popular deployment patterns for updating applications in Kubernetes: rolling updates and blue-green deployments._

_HAProxy Kubernetes Ingress Controller 支持两种流行的部署模式来更新 Kubernetes 中的应用程序：滚动更新和蓝绿部署。_

_This is the second post in a [series about HAProxy's role](https://www.haproxy.com/blog/building-blocks-of-a-modern-proxy/) in building a modern systems architecture that relies on cloud-native technology such-as Docker containers and Kubernetes. Containers have revolutionized how software is deployed, allowing the microservice pattern to flourish and enabling self-healing, autoscaling applications. HAProxy is an intelligent load balancer that adds high performance, observability, security, and many other features to the mix._

_这是 [关于 HAProxy 角色的系列](https://www.haproxy.com/blog/building-blocks-of-a-modern-proxy/) 在构建依赖于云的现代系统架构中的第二篇文章-本地技术，例如 Docker 容器和 Kubernetes。容器彻底改变了软件的部署方式，使微服务模式蓬勃发展，并支持自我修复、自动扩展的应用程序。 HAProxy 是一个智能负载均衡器，它增加了高性能、可观察性、安全性和许多其他特性。_

**Learn more by registering for our webinar** [**_“HAProxy Skills Lab: Deployment Patterns in Kubernetes Using the HAProxy Kubernetes Ingress Controller_ _”._**](https://www.haproxy.com/blog/webinar-haproxy-skills-lab-deployment-patterns-in-kubernetes-using-the-haproxy-kubernetes-ingress-controller/)

**通过注册我们的网络研讨会了解更多信息** [**_“HAProxy Skills Lab: Deployment Patterns in Kubernetes Using the HAProxy Kubernetes Ingress Controller_”._**](https://www.haproxy.com/blog/网络研讨会-haproxy-skills-lab-deployment-patterns-in-kubernetes-using-the-haproxy-kubernetes-ingress-controller/)

So, you have deployed your application to Kubernetes and it’s running flawlessly. The next important question is, how should you deploy the next version of it safely? How can you replace the existing pods without disrupting traffic? Furthermore, how is it affected by routing traffic through the HAProxy Kubernetes Ingress Controller?

因此，您已将应用程序部署到 Kubernetes，并且它运行完美。下一个重要问题是，您应该如何安全地部署它的下一个版本？如何在不中断流量的情况下替换现有的 pod？此外，通过 HAProxy Kubernetes Ingress Controller 路由流量如何影响它？

Kubernetes accommodates a wide range of deployment methods. We’ll cover two that guarantee a safe rollout while keeping the ability to revert if necessary:

Kubernetes 支持多种部署方法。我们将介绍两个保证安全推出的同时保持必要时恢复的能力：

- _Rolling updates_ have first-class support in Kubernetes and allow you to phase in a new version gradually;
- _Blue-green deployments_ avoid having two versions at play at the same time by swapping one set of pods for another.

- _滚动更新_在Kubernetes中拥有一流的支持，并允许您逐步进入新版本；
- _蓝绿部署_通过将一组 Pod 交换为另一组 Pod 来避免同时运行两个版本。

The HAProxy Kubernetes Ingress Controller is powered by the world’s fastest and most widely used software load balancer. Known to provide the utmost performance, observability, and security, it is the most efficient way to route traffic into a Kubernetes cluster. It automatically detects changes within your Kubernetes infrastructure and ensures accurate distribution of traffic to healthy pods. Its design prevents downtime even when there are rapid configuration changes. It supports both deployment patterns and reliably exposes the correct pods to clients.

HAProxy Kubernetes 入口控制器由世界上最快、使用最广泛的软件负载均衡器提供支持。众所周知，它可以提供最高的性能、可观察性和安全性，是将流量路由到 Kubernetes 集群的最有效方式。它会自动检测 Kubernetes 基础设施中的变化，并确保将流量准确分配到健康的 Pod。即使配置发生快速变化，其设计也能防止停机。它支持两种部署模式，并可靠地向客户端公开正确的 pod。

## Deploy the HAProxy Kubernetes Ingress Controller

## 部署 HAProxy Kubernetes Ingress Controller

In this blog post, I use [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) to start up a simple Kubernetes cluster on my workstation. Minikube requires a hypervisor, such as VirtualBox, to be installed. Once it’s up and running, you will be able to expose services running inside the Kubernetes cluster at the IP address **192.168.99.100**.

在这篇博文中，我使用 [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) 在我的工作站上启动一个简单的 Kubernetes 集群。 Minikube 需要安装虚拟机管理程序，例如 VirtualBox。一旦它启动并运行，您将能够在 IP 地址 **192.168.99.100** 处公开在 Kubernetes 集群内运行的服务。

After installing and starting Minikube, deploy the [HAProxy Kubernetes Ingress Controller](https://www.haproxy.com/documentation/hapee/2-0r1/traffic-management/kubernetes-ingress-controller/), which is responsible for routing traffic into your Kubernetes cluster. You can either install the open-source version or the Enterprise version, which is built upon HAProxy Enterprise. It adds features such as a Web Application Firewall, which is essential for stopping application-layer attacks.

Minikube安装启动后，部署 [HAProxy Kubernetes Ingress Controller](https://www.haproxy.com/documentation/hapee/2-0r1/traffic-management/kubernetes-ingress-controller/)，负责路由进入您的 Kubernetes 集群的流量。您可以安装开源版本或基于 HAProxy Enterprise 构建的 Enterprise 版本。它添加了 Web 应用防火墙等功能，这对于阻止应用层攻击至关重要。

By default, the Ingress Controller assumes that you want to configure SSL. If you prefer to try things without SSL, then [download its YAML file](https://raw.githubusercontent.com/haproxytech/kubernetes-ingress/master/deploy/haproxy-ingress.yaml) and modify its `ConfigMap` so that `ssl-redirect` is _OFF_.

默认情况下，入口控制器假定您要配置 SSL。如果你更喜欢尝试没有 SSL 的东西，那么[下载它的 YAML 文件](https://raw.githubusercontent.com/haproxytech/kubernetes-ingress/master/deploy/haproxy-ingress.yaml) 并修改它的`ConfigMap` `ssl-redirect` 是 _OFF_。

haproxy-ingress.yaml



## Rolling Updates 

## 滚动更新

A **rolling update** offers a way to deploy the new version of your application gradually across your cluster. It replaces pods during several phases. For example, you may replace 25% of the pods during the first phase, then another 25% during the next, and so on until all are upgraded. Since the pods are not replaced all at once, this means that both versions will be live, at least for a short time, during the rollout.

**滚动更新** 提供了一种在整个集群中逐步部署应用程序新版本的方法。它在几个阶段替换 pod。例如，您可以在第一阶段更换 25% 的 pod，然后在下一个阶段更换 25%，依此类推，直到全部升级。由于 pod 不会一次全部更换，这意味着两个版本将在推出期间至少在短时间内上线。

**Did You Know?** Because a rolling update creates the potential for two versions of your application to be deployed simultaneously, make sure that any upstream databases and services are compatible with both versions.

**您知道吗？** 因为滚动更新可能会同时部署两个版本的应用程序，所以请确保任何上游数据库和服务都与这两个版本兼容。

This deployment model enjoys first-class support in Kubernetes with baked-in YAML configuration options. Here’s how it works:

此部署模型在 Kubernetes 中享有一流的支持，具有内置的 YAML 配置选项。这是它的工作原理：

1. Version 1 of your application is already deployed.
2. Push version 2 of your application to your container image repository.
3. Update the version number in the Deployment object’s definition.
4. Apply the change with`kubectl`.
5. Kubernetes staggers the rollout of the new version across your pods.
6. The HAProxy Kubernetes Ingress Controller detects when the new pods are live. It automatically updates its proxy configuration, routing traffic away from the old pods and towards the new ones.

1. 您的应用程序的版本 1 已经部署。
2. 将应用程序的第 2 版推送到您的容器映像存储库。
3. 更新部署对象定义中的版本号。
4. 使用 `kubectl` 应用更改。
5. Kubernetes 在您的 pod 中交错推出新版本。
6. HAProxy Kubernetes Ingress Controller 检测新 Pod 何时上线。它会自动更新其代理配置，将流量从旧 Pod 路由到新 Pod。

A rolling update dodges downtime by replacing existing pods incrementally. If the new pods introduce an error that stops them from starting up, Kubernetes will pause the rollout. Also, a rolling update ensures that some pods are always up, so there’s no downtime. Kubernetes keeps a minimum number of pods running during the rollout. However, this requires that you’ve added a _readiness check_ to your pods so that Kubernetes knows when they are truly ready to receive traffic.

滚动更新通过逐步替换现有 pod 来避免停机。如果新的 pod 引入一个错误，导致它们无法启动，Kubernetes 将暂停推出。此外，滚动更新可确保某些 Pod 始终处于运行状态，因此不会出现停机时间。 Kubernetes 在部署期间保持最少数量的 pod 运行。但是，这需要您向 Pod 添加 _readiness check_，以便 Kubernetes 知道它们何时真正准备好接收流量。

### Deploy the Original Application

### 部署原始应用程序

Kubernetes enables rolling updates by default. An update begins when you change your `Deployment` resource’s YAML file and then use `kubectl` apply. Consider the following definition, which deploys version 1 of an application. Note that I am using the [errm/versions](https://hub.docker.com/r/errm/versions) Docker image because it displays the version of the application when you browse to its webpage, which makes it easy to see which version you're running.

Kubernetes 默认启用滚动更新。当您更改 `Deployment` 资源的 YAML 文件然后使用 `kubectl` apply 时，更新开始。考虑以下定义，它部署应用程序的第 1 版。请注意，我使用的是 [errm/versions](https://hub.docker.com/r/errm/versions) Docker 映像，因为它会在您浏览应用程序的网页时显示该应用程序的版本，这使得您可以轻松地看看你运行的是哪个版本。

app.yaml



The `readinessProbe` section tells Kubernetes to send an HTTP request to the application five seconds after it has started, and then every five seconds thereafter. No traffic is sent to the pod until a successful response is returned. This is key to preventing downtime.

`readinessProbe` 部分告诉 Kubernetes 在应用程序启动 5 秒后向应用程序发送 HTTP 请求，然后每 5 秒发送一次。在返回成功响应之前，不会向 Pod 发送流量。这是防止停机的关键。

**Did You Know?** Consider tagging your container images with version numbers, rather than using a tag like _latest_. This allows you to keep track of the versions that are deployed and manage the release of new versions.

**您知道吗？** 考虑使用版本号标记您的容器映像，而不是使用 _latest_ 之类的标记。这允许您跟踪部署的版本并管理新版本的发布。

Next, define a `Service` object that will categorize the pods into a single group that the Ingress Controller will watch:

接下来，定义一个“Service”对象，该对象将 Pod 分类为 Ingress Controller 将监视的单个组：

app-service.yaml



Next, define an `Ingress` object. This configures how the HAProxy Ingress Controller will route traffic to the pods:

接下来，定义一个 Ingress 对象。这将配置 HAProxy 入口控制器如何将流量路由到 pod：

ingress.yaml



Use `kubectl apply` to deploy the pods, service and ingress:

使用 `kubectl apply` 来部署 pods、服务和入口：

Version 1 of your application is now deployed. Run the following command to see which port the HAProxy Kubernetes Ingress Controller has mapped to port 80:

您的应用程序的第 1 版现已部署。运行以下命令以查看 HAProxy Kubernetes Ingress Controller 已将哪个端口映射到端口 80：

You can then see that the application is exposed on port 31179. You can see it by visiting the Minikube IP address **http://192.168.99.100:31179** in your browser.

然后您可以看到应用程序暴露在端口 31179 上。您可以通过在浏览器中访问 Minikube IP 地址 **http://192.168.99.100:31179** 来查看它。

![Version 1 web page](https://cdn.haproxy.com/wp-content/uploads/2020/02/version1.png)

Version 1 web page

版本 1 网页

Let’s see how to upgrade it to version 2 next.

接下来让我们看看如何将其升级到版本 2。

### Upgrade Using a Rolling Update

### 使用滚动更新升级

After you have pushed a new version of your application to your container repository, trigger a rolling update by increasing the version number set on the `Deployment` definition’s `spec.template.spec.containers.image` property. This tells Kubernetes that the current, desired version of your application has changed. In our example, since we’re using a prebaked image, there’s already a version 2 set up in the Docker Hub repository.

将应用程序的新版本推送到容器存储库后，通过增加在 `Deployment` 定义的 `spec.template.spec.containers.image` 属性上设置的版本号来触发滚动更新。这告诉 Kubernetes 您的应用程序当前所需的版本已更改。在我们的示例中，由于我们使用的是预烘焙镜像，因此 Docker Hub 存储库中已经设置了版本 2。

app.yaml



Then, use `kubectl apply` to start the rollout:

然后，使用 `kubectl apply` 开始部署：

You can check the status of the rollout by using the `kubectl rollout status` command: 

您可以使用 `kubectl rollout status` 命令检查 rollout 的状态：

Once completed, you can access the application again at the same URL, **http://192.168.99.100:31179**. It shows you a new web page signifying that version 2 has been deployed.

完成后，您可以通过同一 URL **http://192.168.99.100:31179** 再次访问该应用程序。它会向您显示一个新网页，表明版本 2 已部署。

![Version 2 web page](https://cdn.haproxy.com/wp-content/uploads/2020/02/version2.png)

Version 2 web page



If you decide that the new version is faulty, you can revert to the previous one by using the `kubectl rollout undo` command, like this:

如果您认为新版本有问题，您可以使用 `kubectl rollout undo` 命令恢复到以前的版本，如下所示：

The HAProxy Kubernetes Ingress Controller detects pod changes quickly and can switch back and forth between versions without dropping connections. Rolling updates aren’t the only way to accomplish highly-available services, though. In the next section, you’ll learn about blue-green deployments, which update all pods simultaneously.

HAProxy Kubernetes Ingress Controller 可以快速检测 pod 更改，并且可以在版本之间来回切换而不会断开连接。不过，滚动更新并不是实现高可用性服务的唯一方法。在下一节中，您将了解蓝绿部署，它同时更新所有 pod。

## Blue-Green Deployments

## 蓝绿部署

A **blue-green deployment** lets you replace an existing version of your application across all pods at once. The name, blue-green, [was coined](https://gitlab.com/snippets/1846041) in the book _Continuous Delivery_ by Jez Humble and David Farley. Here’s how it works:

**蓝绿部署** 允许您一次在所有 pod 中替换应用程序的现有版本。 Jez Humble 和 David Farley 在 _Continuous Delivery_ 一书中[创造](https://gitlab.com/snippets/1846041) 这个名字，蓝绿色。这是它的工作原理：

1. Version 1 of your application is already deployed.
2. Push version 2 of your application to your container image repository.
3. Deploy version 2 of your application to a new group of pods. Both versions 1 and 2 pods are now running in parallel. However, only version 1 is exposed to external clients.
4. Run internal testing on version 2 and make sure it is ready to go live.
5. Flip a switch and the ingress controller in front of your clusters stops routing traffic to the version 1 pods and starts routing it to the version 2 pods.

1. 您的应用程序的版本 1 已经部署。
2. 将应用程序的第 2 版推送到您的容器映像存储库。
3. 将应用程序的第 2 版部署到一组新的 pod。版本 1 和 2 pod 现在并行运行。但是，只有版本 1 暴露给外部客户端。
4. 在版本 2 上运行内部测试并确保它已准备好上线。
5. 翻转开关，集群前面的入口控制器停止将流量路由到版本 1 的 Pod，并开始将其路由到版本 2 的 Pod。

This deployment pattern has a few advantages over a rolling update. For one, at no time are there ever two versions of your application accessible to external clients at the same time. So, all users will receive the same client-side Javascript files and be routed to a version of the application that supports the API calls within those files. It also simplifies upstream dependencies, such as database schemas.

与滚动更新相比，这种部署模式有一些优势。一方面，任何时候都不会有两个版本的应用程序可供外部客户端同时访问。因此，所有用户都将收到相同的客户端 Javascript 文件，并被路由到支持这些文件中 API 调用的应用程序版本。它还简化了上游依赖项，例如数据库模式。

Another advantage is that it gives you time to test the new version in a production environment before it goes live. You control how long to wait before making the switch. Meanwhile, you can verify that the application and its dependencies function normally.

另一个优点是它让您有时间在新版本上线之前在生产环境中对其进行测试。您可以控制在进行切换之前等待的时间。同时，您可以验证应用程序及其依赖项是否正常运行。

On the other hand, a blue-green deployment is all-or-nothing. Unlike a rolling update, you aren’t able to gradually roll out the new version. All users will receive the update at the same time, although existing sessions will be allowed to finish their work on the old instances. So, the stakes are a bit higher that everything should work, once you do initiate the change. It also requires allocating more server resources, since you will need to run two copies of every pod.

另一方面，蓝绿部署是全有或全无。与滚动更新不同，您无法逐步推出新版本。所有用户将同时收到更新，但允许现有会话在旧实例上完成他们的工作。因此，一旦您开始更改，一切都应该起作用的风险要高一些。它还需要分配更多的服务器资源，因为您需要运行每个 Pod 的两个副本。

Luckily, the rollback procedure is just as easy: You simply flip the switch again and the previous version is swapped back into place. That’s because the old version is still running on the old pods. It is simply that traffic is no longer being routed to them. When you’re confident that the new version is here to stay, you can decommission those pods.

幸运的是，回滚过程同样简单：您只需再次拨动开关，先前的版本就会换回原位。那是因为旧版本仍在旧 Pod 上运行。只是流量不再被路由到他们。当您确信新版本会继续存在时，您可以停用这些 pod。

You’ll need to set up your original application in a slightly different way when you expect to use a blue-green deployment. There is more emphasis on using Kubernetes metadata labels, which will become clear in the next section.

当您希望使用蓝绿部署时，您需要以稍微不同的方式设置您的原始应用程序。更强调使用 Kubernetes 元数据标签，这将在下一节中变得清晰。

### Deploy the Original Application

### 部署原始应用程序

Consider the following definition, which deploys version 1 of your application. Note its `spec.selector` section, which specifies a label called _version_:

考虑以下定义，它部署应用程序的第 1 版。注意它的 `spec.selector` 部分，它指定了一个名为 _version_ 的标签：

app-v1.yaml



A `Deployment` object defines a `spec.selector` section that matches the `spec.template.metadata` section. This is how a Deployment tags pods and keeps track of them. This is the key to setting up a blue-green deployment. By using different labels, you can deploy multiple versions of the same application. Here, the `spec.selector.matchLabels` property is set to _run=app,version=0.0.1_. The version should match the version tag of your Docker image, for convenience and simplicity.

`Deployment` 对象定义了一个与 `spec.template.metadata` 部分匹配的 `spec.selector` 部分。这就是 Deployment 标记 pod 并跟踪它们的方式。这是设置蓝绿部署的关键。通过使用不同的标签，您可以部署同一应用程序的多个版本。这里，`spec.selector.matchLabels` 属性设置为 _run=app,version=0.0.1_。为方便和简单起见，版本应与 Docker 映像的版本标记匹配。

The following Service definition targets that same selector:

以下服务定义针对同一个选择器：

app-service-bg.yaml



Next, use the following `Ingress` definition to expose the version 1 pods to the world. It registers a route with the HAProxy Kubernetes Ingress Controller:

接下来，使用以下 `Ingress` 定义向世界公开版本 1 的 pod。它向 HAProxy Kubernetes Ingress Controller 注册一个路由：

ingress.yaml



Apply everything using `kubectl`: 

使用 `kubectl` 应用所有内容：

At this point, you can access the application at the HTTP port exposed by the Ingress Controller: **http://192.168.99.100:31179**. Now, let’s see how to use a blue-green deployment to upgrade the version.

此时可以在Ingress Controller暴露的HTTP端口访问应用：**http://192.168.99.100:31179**。现在，让我们看看如何使用蓝绿部署来升级版本。

### Upgrade Using a Blue-green Deployment

### 使用蓝绿部署升级

Now that the _blue_ version (ie version 1) is released, create a _green_ version of your `Deployment` object that will deploy version 2. The YAML will be the same, except that you increase the value of the _version_ label, as well as the Docker image tag. Also note that the name of the deployment is changed from _app-blue_ to _app-green_, since you cannot have two Deployments with the same name that target different pods.

现在 _blue_ 版本（即版本 1）已发布，创建一个 _green_ 版本的 `Deployment` 对象将部署版本 2。 YAML 将是相同的，除了增加 _version_ 标签的值，以及Docker 镜像标签。另请注意，部署的名称从 _app-blue_ 更改为 _app-green_，因为您不能有两个针对不同 Pod 的具有相同名称的部署。

app-v2.yaml



Apply it with `kubectl`:

用 `kubectl` 应用它：

At this point, both blue (version 1) and green (version 2) are deployed. Only the blue instance is receiving traffic, though. To make the switch, update your `Service` definition’s _version_ selector so that it points to the new version:

此时，蓝色（版本 1）和绿色（版本 2）都已部署。但是，只有蓝色实例正在接收流量。要进行切换，请更新您的“Service”定义的 _version_ 选择器，使其指向新版本：

app-service.yaml



Apply it with `kubectl`:

用 `kubectl` 应用它：

Check the application again and you will see that the new version is live. If you need to roll back to the earlier version, simply change the `Service` definition’s selector back and reapply it. The HAProxy Kubernetes Ingress Controller detects these changes almost instantly and you can swap back and forth to your heart’s content. There’s no downtime during the cutover. Established TCP connections will finish normally on the instance where they began.

再次检查应用程序，您将看到新版本已上线。如果您需要回滚到早期版本，只需将 `Service` 定义的选择器更改回并重新应用它。 HAProxy Kubernetes 入口控制器几乎立即检测到这些变化，您可以来回交换您的心意。切换期间没有停机时间。已建立的 TCP 连接将在它们开始的实例上正常结束。

### Testing the New Pods

### 测试新的 Pod

You can also test the new version before it’s released by registering a different ingress route that exposes the application at a new URL path. First, create another `Service` definition called _test-service_:

您还可以在新版本发布之前通过注册不同的入口路由来测试新版本，该路由在新的 URL 路径中公开应用程序。首先，创建另一个名为 _test-service_ 的 `Service` 定义：

test-service.yaml



Note that we are including the path-rewrite annotation, which rewrites the URL **/test** to **/** before it reaches the pod. Then, add a new route to your existing `Ingress` object that exposes this service at the URL path **/test**, as shown:

请注意，我们包含了 path-rewrite 注释，它在 URL **/test** 到达 pod 之前将其重写为 **/**。然后，向现有的 Ingress 对象添加一条新路由，该路由在 URL 路径 **/test** 处公开此服务，如下所示：

ingress.yaml



This lets you check your application by visiting **/test** in your browser.

这让您可以通过在浏览器中访问 **/test** 来检查您的应用程序。

## Conclusion

##  结论

The HAProxy Kubernetes Ingress Controller is powered by the legendary HAProxy. Known to provide the utmost performance, observability, and security, it features many benefits including SSL termination, rate limiting, and IP whitelisting. When you deploy the ingress controller into your cluster, it’s important to consider how your applications will be upgraded later. Two popular methods are rolling updates and blue-green deployments.

HAProxy Kubernetes 入口控制器由传奇的 HAProxy 提供支持。它以提供最佳性能、可观察性和安全性而著称，具有许多优点，包括 SSL 终止、速率限制和 IP 白名单。当您将入口控制器部署到集群中时，重要的是要考虑您的应用程序以后将如何升级。两种流行的方法是滚动更新和蓝绿部署。

Rolling updates allow you to phase in a new version gradually and it has first-class support in Kubernetes. Blue-green deployments avoid the complexity of having two versions at play at the same time and give you a chance to test the change before going live. In either case, the HAProxy Kubernetes Ingress Controller detects these changes quickly and maintains uptime throughout.

滚动更新允许您逐步引入新版本，并且它在 Kubernetes 中具有一流的支持。蓝绿部署避免了同时运行两个版本的复杂性，并让您有机会在上线之前测试更改。在任何一种情况下，HAProxy Kubernetes 入口控制器都会快速检测到这些变化并始终保持正常运行时间。

If you enjoyed this post and want to see more like it, subscribe to this blog! You can also follow us on [Twitter](https://twitter.com/haproxy) and join the conversation on [Slack](https://slack.haproxy.com/).

如果您喜欢这篇文章并希望看到更多类似的文章，请订阅此博客！您也可以在 [Twitter](https://twitter.com/haproxy) 上关注我们并在 [Slack](https://slack.haproxy.com/) 上加入对话。

The Enterprise version of the ingress controller combines HAProxy, the world’s fastest and most widely used open-source software load balancer and application delivery controller, with enterprise-class features, services and premium support. [Contact us](https://www.haproxy.com/contact-us/) to learn more and sign up for a [free trial](https://www.haproxy.com/downloads/hapee-trial/) . 

入口控制器的企业版将 HAProxy（世界上最快、使用最广泛的开源软件负载平衡器和应用程序交付控制器）与企业级功能、服务和高级支持相结合。 [联系我们](https://www.haproxy.com/contact-us/)了解更多信息并注册[免费试用](https://www.haproxy.com/downloads/hapee-trial/) .

