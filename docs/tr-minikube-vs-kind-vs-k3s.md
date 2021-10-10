# Minikube vs. kind vs. k3s - What should I use?

# Minikube vs. kind vs. k3s - 我应该使用什么？

http://brennerm.github.io/posts/minikube-vs-kind-vs-k3s.html

December 5, 2019

2019 年 12 月 5 日

These days there are a few tools that claim to (partially) replace a fully fledged Kubernetes cluster. Using them allows e.g. every developer to have their own local cluster instance running to play around with it, deploy their application or execute tests against applications running in K8s during CI/CD. In this post we’ll have a look at three of them, compare their pros and cons and identify use cases for each of them.

如今，有一些工具声称（部分）取代了成熟的 Kubernetes 集群。使用它们允许例如每个开发人员都可以运行自己的本地集群实例来使用它，部署他们的应用程序或在 CI/CD 期间针对在 K8s 中运行的应用程序执行测试。在这篇文章中，我们将看看其中的三个，比较它们的优缺点并确定每个用例的用例。

## [minikube](http://brennerm.github.io\#minikube)

minikube is a Kubernetes SIGs project and has been started more than three years ago. It takes the approach of spawning a VM that is essentially a single node K8s cluster. Due to the support for a bunch of hypervisors it can be used on all of the major operating systems. This also allows you to create multiple instances in parallel.

minikube 是 Kubernetes SIG 的一个项目，已经启动三年多了。它采用生成虚拟机的方法，该虚拟机本质上是一个单节点 K8s 集群。由于支持一堆管理程序，它可以在所有主要操作系统上使用。这也允许您并行创建多个实例。

From a user perspective minikube is a very beginner friendly tool. You start the cluster using `minikube start`, wait a few minutes and your `kubectl` is ready to go. To specify a Kubernetes version you can use the `--kubernetes-version` flag. A list of supported versions can be found [here](https://minikube.sigs.k8s.io/docs/reference/configuration/kubernetes/).

从用户的角度来看，minikube 是一个非常适合初学者的工具。你使用 `minikube start` 启动集群，等待几分钟，你的 `kubectl` 就准备好了。要指定 Kubernetes 版本，您可以使用 `--kubernetes-version` 标志。可在 [此处](https://minikube.sigs.k8s.io/docs/reference/configuration/kubernetes/) 找到支持版本的列表。

If you are new to Kubernetes the first class support for its dashboard that minikube offers may help you. With a simple `minikube dashboard` the application will open up giving you a nice overview of everything that is going on in your cluster. This is being achieved by [minikube's addon system](https://minikube.sigs.k8s.io/docs/tasks/addons/) that helps you integrating things like, [Helm](https://helm.sh/) , [Nvidia GPUs](https://developer.nvidia.com/kubernetes-gpu) and an [image registry](https://docs.docker.com/registry/) with your cluster.

如果您是 Kubernetes 的新手，minikube 为其仪表板提供的一流支持可能会对您有所帮助。使用简单的“minikube 仪表板”，应用程序将打开，让您对集群中发生的一切有一个很好的概览。这是通过 [minikube 的插件系统](https://minikube.sigs.k8s.io/docs/tasks/addons/) 实现的，它可以帮助您集成诸如 [Helm](https://helm.sh/) , [Nvidia GPU](https://developer.nvidia.com/kubernetes-gpu) 和 [图像注册表](https://docs.docker.com/registry/) 与您的集群。

## [kind](http://brennerm.github.io\#kind)

Kind is another Kubernetes SIGs project but is quite different compared to minikube. As the name suggests it moves the cluster into Docker containers. This leads to a significantly faster startup speed compared to spawning VM.

Kind 是另一个 Kubernetes SIG 项目，但与 minikube 相比有很大不同。顾名思义，它将集群移动到 Docker 容器中。与生成 VM 相比，这导致启动速度明显更快。

Creating a cluster is very similar to minikube’s approach. Executing `kind create cluster`, playing the waiting game and afterwards you are good to go. By using different names ( `--name`) kind allows you to create multiple instances in parallel.

创建集群与 minikube 的方法非常相似。执行`kind create cluster`，玩等待游戏，然后你就可以开始了。通过使用不同的名称（`--name`） kind 允许您并行创建多个实例。

One feature that I personally enjoy is the ability to load my local images directly into the cluster. This saves me a few extra steps of setting up a registry and pushing my image each and every time I want to try out my changes. With a simple `kind load docker-image my-app:latest` the image is available for use in my cluster. Very nice!

我个人喜欢的一项功能是能够将我的本地图像直接加载到集群中。这为我节省了设置注册表和每次我想尝试更改时推送我的图像的一些额外步骤。通过简单的“kind load docker-image my-app:latest”，该图像可用于我的集群。非常好！

If you are looking for a way to programmatically create a Kubernetes cluster, kind kindly (you have been for waiting for this, don’t you :P) publishes its Go packages that are used under the hood. If you want to get to know more have a look at the [GoDocs](https://godoc.org/sigs.k8s.io/kind/pkg/cluster) and check out how [KUDO uses kind for their integration tests](https://github.com/kudobuilder/kudo/blob/f7b09025f5c2faf5492624facc1dc4c5c7a5ccad/pkg/test/harness.go#L105).

如果您正在寻找一种以编程方式创建 Kubernetes 集群的方法，请善待（您一直在等待，不是吗:P）发布其在后台使用的 Go 包。如果您想了解更多信息，请查看 [GoDocs](https://godoc.org/sigs.k8s.io/kind/pkg/cluster) 并查看 KUDO 如何使用 kind 进行集成测试。

## [k3s](http://brennerm.github.io\#k3s)

K3s is a minified version of Kubernetes developed by [Rancher Labs](https://rancher.com/). By removing dispensable features (legacy, alpha, non-default, in-tree plugins) and using lightweight components (e.g. sqlite3 instead of etcd3) they achieved a significant downsizing. This results in a single binary with a size of around 60 MB.

K3s 是 [Rancher Labs](https://rancher.com/) 开发的 Kubernetes 的缩小版。通过删除可有可无的功能（遗留、alpha、非默认、树内插件）并使用轻量级组件（例如 sqlite3 而不是 etcd3)，他们实现了显着的缩小。这会生成一个大小约为 60 MB 的单个二进制文件。

The application is split into the K3s server and the agent. The former acts as a manager while the latter is responsible for handling the actual workload. I discourage you from running them on your workstation as this leads to some clutter in your local filesystem. Instead put k3s in a container (e.g. by using [rancher/k3s](https://hub.docker.com/r/rancher/k3s)) which also allows you to easily run several independent instances. 

应用分为K3s服务器和代理。前者充当经理，而后者负责处理实际工作量。我不鼓励您在工作站上运行它们，因为这会导致本地文件系统中出现一些混乱。而是将 k3s 放在容器中（例如通过使用 [rancher/k3s](https://hub.docker.com/r/rancher/k3s))，这也允许您轻松运行多个独立实例。

One feature that stands out is called [auto deployment](https://rancher.com/docs/k3s/latest/en/configuration/#auto-deploying-manifests). It allows you to deploy your Kubernetes manifests and Helm charts by putting them in a specific directory. K3s watches for changes and takes care of applying them without any further interaction. This is especially useful for CI pipelines and IoT devices (both target use cases of K3s). Just create/update your configuration and K3s makes sure to keep your deployments up to date.

突出的一项功能称为[自动部署](https://rancher.com/docs/k3s/latest/en/configuration/#auto-deploying-manifests)。它允许您通过将 Kubernetes 清单和 Helm 图表放在特定目录中来部署它们。 K3s 监视更改并负责应用它们，无需任何进一步交互。这对于 CI 管道和 IoT 设备（两者都是 K3s 的目标用例)特别有用。只需创建/更新您的配置，K3s 就会确保您的部署保持最新。

## [Summary](http://brennerm.github.io\#summary)

## [总结](http://brennerm.github.io\#summary)

I was a long time minikube user as there where simply no alternatives (at least I never heard of one) and to be honest…it does a pretty good job at being a local Kubernetes development environment. You create the cluster, wait a few minutes and you are good to go. However for my use cases (mostly playing around with tools that run on K8s) I could fully replace it with kind due to the quicker setup time. If you are working in an environment with a tight resource pool or need an even quicker startup time, K3s is definitely a tool you should consider.

我是 minikube 的长期用户，因为那里根本没有替代品（至少我从未听说过），老实说……它在成为本地 Kubernetes 开发环境方面做得非常好。您创建集群，等待几分钟，您就可以开始了。然而，对于我的用例（主要是使用在 K8s 上运行的工具），由于设置时间更快，我可以完全用 kind 替换它。如果您在资源池紧张的环境中工作，或者需要更快的启动时间，K3s 绝对是您应该考虑的工具。

All in all these three tools are doing the job while using different approaches and focusing on different use cases. I hope you got a better understanding on how they work and which is the best candidate for solving your upcoming issue. Feel free to share your experience and let me know about use cases you are realizing with minikube, kind or k3s at [@\_\_brennerm](https://twitter.com/__brennerm).

总而言之，这三个工具在使用不同的方法并专注于不同的用例时都在完成这项工作。我希望您对它们的工作方式有更好的了解，以及哪个是解决即将出现的问题的最佳人选。请随时分享您的经验，并在 [@\_\_brennerm](https://twitter.com/__brennerm) 上告诉我您使用 minikube、kind 或 k3s 实现的用例。

Below you can find a table that lists a few key facts of each tool.

您可以在下面找到一个表格，其中列出了每个工具的一些关键事实。

* * *

## Comments

##  注释

If you have questions or want to give feedback feel free to contact me. 

如果您有任何疑问或想提供反馈，请随时与我联系。

