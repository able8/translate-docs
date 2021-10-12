# Helm Charts Tutorial: The Kubernetes Package Manager Explained

# Helm Charts 教程：解释 Kubernetes 包管理器

![Sebastian Sigl](http://www.freecodecamp.org/news/content/images/size/w100/2020/01/profile_photo_edited.png)[SebastianSigl](http://www.freecodecamp.org/news/author/sesigl/)

/作者/sesigl/)

![Helm Charts Tutorial: The Kubernetes Package Manager Explained](http://www.freecodecamp.org/news/content/images/size/w2000/2020/12/helm-blog-logo.jpg)

There are different ways of running production services at a high scale. One popular solution for running containers in production is Kubernetes. But interacting with Kubernetes directly comes with some caveats.

有多种方式可以大规模运行生产服务。在生产中运行容器的一种流行解决方案是 Kubernetes。但是直接与 Kubernetes 交互会带来一些警告。

Helm tries to solve some of the challenges with useful features that increase productivity and reduce maintenance efforts of complex deployments.

Helm 试图通过有用的功能来解决一些挑战，这些功能可以提高生产力并减少复杂部署的维护工作。

In this post you will learn:

在这篇文章中，您将了解到：

- What Helm is
- The most common use-cases of Helm
- How to configure and deploy a publicly available Helm package
- How to deploy a custom application using Helm

- 什么是头盔
- Helm 最常见的用例
- 如何配置和部署公开可用的 Helm 包
- 如何使用 Helm 部署自定义应用程序

Every code example in this post requires a Kubernetes cluster. The easiest way to get a cluster to play with is to install Docker and activate its Kubernetes cluster feature. Also, you need to [install kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and [Helm](https://helm.sh/docs/intro/install/) to interact with your cluster.

本文中的每个代码示例都需要一个 Kubernetes 集群。使用集群的最简单方法是安装 Docker 并激活其 Kubernetes 集群功能。此外，您需要[安装 kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) 和 [Helm](https://helm.sh/docs/intro/install/) 以与您的集群交互。

_Please note: When you try the examples, be patient. If you are too fast then the containers are not ready. It might take a few minutes until the containers can receive requests._

_请注意：当您尝试示例时，请耐心等待。如果你太快，那么容器还没有准备好。容器可能需要几分钟才能收到请求。_

## What is Helm?

## 什么是头盔？

Helm calls itself ”The Kubernetes package manager”. It is a command-line tool that enables you to create and use so-called Helm Charts.

Helm 称自己为“Kubernetes 包管理器”。它是一个命令行工具，可让您创建和使用所谓的 Helm Charts。

A Helm Chart is a collection of templates and settings that describe a set of Kubernetes resources. Its power spans from managing a single node definition to a highly scalable multi-node cluster.

Helm Chart 是描述一组 Kubernetes 资源的模板和设置的集合。它的强大功能涵盖从管理单节点定义到高度可扩展的多节点集群。

The architecture of Helm has changed over the last years. The current version of Helm communicates directly to your Kubernetes cluster via Rest. If you read something about Tiller in the context of Helm, then you're reading an old article. Tiller was removed in Helm 3.

Helm 的架构在过去几年中发生了变化。当前版本的 Helm 通过 Rest 直接与您的 Kubernetes 集群通信。如果您在 Helm 的上下文中阅读了有关 Tiller 的内容，那么您正在阅读一篇旧文章。 Tiller 在 Helm 3 中被移除。

Helm itself is stateful. When a Helm Chart gets installed, the defined resources are getting deployed and meta-information is stored in Kubernetes secrets.

Helm 本身是有状态的。安装 Helm Chart 后，将部署定义的资源，并将元信息存储在 Kubernetes 机密中。

## How to Deploy a Simple Helm Application

## 如何部署一个简单的 Helm 应用程序

Let’s get our hands dirty and make sure Helm is ready to use.

让我们亲自动手并确保 Helm 已准备好使用。

First, we need to be connected to a Kubernetes cluster. In this example, I will concentrate on a Kubernetes cluster that comes with your Docker setup. So if you use some other Kubernetes cluster, configurations and outputs might differ.

首先，我们需要连接到 Kubernetes 集群。在这个例子中，我将专注于你的 Docker 设置附带的 Kubernetes 集群。因此，如果您使用其他 Kubernetes 集群，配置和输出可能会有所不同。

```shell
$ kubectl config use-context docker-desktop

Switched to context "docker-desktop".

$ kubectl get node

NAME             STATUS   ROLES    AGE   VERSION
docker-desktop   Ready    master   20d   v1.19.3
```

Let’s deploy an Apache webserver using Helm. As a first step, we need to tell Helm what location to search by adding a Helm repository:

让我们使用 Helm 部署一个 Apache 网络服务器。作为第一步，我们需要通过添加 Helm 存储库来告诉 Helm 要搜索的位置：

```shell
$ helm repo add bitnami https://charts.bitnami.com/bitnami
```

Let’s install the actual container:

让我们安装实际的容器：

```shell
$ helm install my-apache bitnami/apache --version 8.0.2
```

After a few minutes your deployment is ready. We can check the state of the containers using kubectl:

几分钟后，您的部署已准备就绪。我们可以使用 kubectl 检查容器的状态：

```shell
$ kubectl get pods

NAME                               READY   STATUS    RESTARTS   AGE
my-apahe-apache-589b8df6bd-q6m2n   1/1     Running   0          2m27s
```

Now, open [http://localhost](http://localhost) to see the default Apache exposed website locally. Also, Helm can show us information about current deployments:

现在，打开 [http://localhost](http://localhost) 以在本地查看默认的 Apache 公开网站。此外，Helm 可以向我们显示有关当前部署的信息：

```shell
$ helm list

NAME         REVISION    STATUS      CHART           VERSION
my-apache    1        deployed      apache-8.0.2    2.4.46
```

### How to Upgrade a Helm Application

### 如何升级 Helm 应用程序

We can upgrade our deployed application to a new version like this:

我们可以像这样将部署的应用程序升级到新版本：

```shell
$ helm upgrade my-apache bitnami/apache --version 8.0.3

$ helm list

NAME         REVISION    STATUS      CHART          VERSION
my-apache    2           deployed    apache-8.0.3    2.4.46
```

The column Revision indicates that this is the 2nd version we've deployed.

Revision 列表示这是我们部署的第二个版本。

### How to Rollback a Helm Application

### 如何回滚 Helm 应用程序

So let’s try to rollback to the first deployed version:

所以让我们尝试回滚到第一个部署的版本：

```shell
$ helm rollback my-apache 1

Rollback was a success!Happy Helming!

$ helm list

NAME         REVISION    STATUS      CHART        VERSION
my-apache    3           deployed    apache-8.0.2    2.4.46
```

This is a very powerful feature that allows you to roll back changes in production quickly.

这是一项非常强大的功能，可让您快速回滚生产中的更改。

I mentioned that Helm stores deployment information in secrets – here they are:

我提到 Helm 将部署信息存储在秘密中——它们是：

```shell
$ kubectl get secret

NAME                    TYPE                         DATA   AGE
default-token-nc4hn               kubernetes.io/sat        3      20d
sh.helm.release.v1.my-apache.v1   helm.sh/release.v1        1      1m
sh.helm.release.v1.my-apache.v2   helm.sh/release.v1        1      1m
sh.helm.release.v1.my-apache.v3   helm.sh/release.v1        1      1m
```

### How to Remove a Deployed Helm Application

### 如何删除已部署的 Helm 应用程序

Let’s clean up our Kubernetes by removing the my-apache release:

让我们通过删除 my-apache 版本来清理我们的 Kubernetes：

```shell
$ helm delete my-apache

release "my-apache" uninstalled
```

Helm gives you a very convenient way of managing a set of applications that enables you to deploy, upgrade, rollback and delete.

Helm 为您提供了一种非常方便的方法来管理一组应用程序，使您能够部署、升级、回滚和删除。

Now, we are ready to use more advanced Helm features that will boost your productivity!

现在，我们已准备好使用更高级的 Helm 功能来提高您的工作效率！

## How to Access Production-Ready Helm Charts

## 如何访问生产就绪的 Helm Charts

You can search public hubs for Charts that enable you to quickly deploy your desired application with a customizable configuration.

您可以在公共中心搜索图表，使您能够使用可自定义的配置快速部署所需的应用程序。

A Helm Chart doesn't just contain a static set of definitions. Helm comes with capabilities to hook into any lifecycle state of a Kubernetes deployment. This means during the installation or upgrade of an application, various actions can be executed like creating a database update before updating the actual database.

Helm Chart 不仅仅包含一组静态定义。 Helm 具有挂钩 Kubernetes 部署的任何生命周期状态的功能。这意味着在安装或升级应用程序期间，可以执行各种操作，例如在更新实际数据库之前创建数据库更新。

This powerful definition of Helm Charts lets you share and improve an executable description of a deployment setup that spans from initial installation and version upgrades to rollback capabilities.

Helm Charts 的这个强大定义让您可以共享和改进部署设置的可执行描述，该描述涵盖从初始安装和版本升级到回滚功能。

Helm might be heavy for a simple container like a single node web server, but it’s very useful for more complex applications. For example it works great for a distributed system like Kafka or Cassandra that usually runs on multiple distributed nodes on different datacenters.

Helm 对于像单节点 Web 服务器这样的简单容器来说可能很笨重，但对于更复杂的应用程序非常有用。例如，它非常适用于像 Kafka 或 Cassandra 这样的分布式系统，这些系统通常在不同数据中心的多个分布式节点上运行。

We've already leveraged Helm to deploy a single Apache container. Now, we will deploy a production-ready WordPress application that contains:

我们已经利用 Helm 部署了单个 Apache 容器。现在，我们将部署一个生产就绪的 WordPress 应用程序，其中包含：

- Containers that serve WordPress,
- Instances of MariaDB for persistence and
- Prometheus sidecar containers for each WordPress container to expose health metrics.

- 为 WordPress 提供服务的容器，
- MariaDB 实例用于持久性和
- 每个 WordPress 容器的 Prometheus sidecar 容器以公开健康指标。

Before we deploy, it’s recommended to increase your Docker limits to at least 4GB of memory.

在我们部署之前，建议将您的 Docker 限制增加到至少 4GB 的内存。

Setting everything up sounds like a job that would take weeks. To make it resilient and scale, probably a job that would take months. In these areas, Helm Charts can really shine. Due to the growing community, there might already be a Helm Chart that we can use.

设置一切听起来像是一项需要数周时间的工作。为了使其具有弹性和规模，可能需要几个月的时间。在这些领域，Helm Charts 可以真正发挥作用。由于社区不断壮大，我们可能已经可以使用 Helm Chart。

### How to Deploy WordPress and MariaDB

### 如何部署 WordPress 和 MariaDB

There are different public hubs for Helm Charts. One of them is [artifacthub.io](https://artifacthub.io). We can search for “WordPress” and find an interesting [WordPress Chart](https://artifacthub.io/packages/helm/bitnami/wordpress).

Helm Charts 有不同的公共中心。其中之一是 [artifacthub.io](https://artifacthub.io)。我们可以搜索“WordPress”，找到一个有趣的[WordPress图表](https://artifacthub.io/packages/helm/bitnami/wordpress)。

On the right side, there is an install button. If you click it, you get clear instructions about what to do:

在右侧，有一个安装按钮。如果您单击它，您将获得有关如何操作的明确说明：

```shell
$ helm repo add bitnami https://charts.bitnami.com/bitnami

$ helm install my-wordpress bitnami/wordpress --version 10.1.4
```

You will also see some instructions that tell you how to access the admin interface and the admin password after installation.

您还将看到一些说明，告诉您安装后如何访问管理界面和管理密码。

Here is how you can get and decode the password for the **admin** user on Mac OS:

以下是在 Mac OS 上获取和解码 **admin** 用户密码的方法：

```shell
$ echo Username: user
$ echo Password: $(kubectl get secret --namespace default my-wordpress-3 -o jsonpath="{.data.wordpress-password}" | base64 --decode)

Username: user
Password: sZCa14VNXe
```

On windows, you can get the password for the **user** user in the powershell:

在 Windows 上，您可以在 powershell 中获取 **user** 用户的密码：

```powershell
$pw=kubectl get secret --namespace default my-wordpress -o jsonpath="{.data.wordpress-password}"
[System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($pw))
```

Our local development will be available at: [http://localhost](http://localhost).

我们的本地开发将在以下位置提供：[http://localhost](http://localhost)。

Our admin interface will be available at: [https://localhost/admin](https://localhost/admin%20).

我们的管理界面将位于：[https://localhost/admin](https://localhost/admin%20)。

So we have everything to run it locally. But in production, we want to scale some parts of it to serve more and more visitors. We can scale the number of WordPress services. We also want to expose some health metrics like the usage of our CPU and memory.

所以我们拥有在本地运行它的一切。但是在生产中，我们希望扩展其中的某些部分以服务越来越多的访问者。我们可以扩展 WordPress 服务的数量。我们还想公开一些健康指标，例如 CPU 和内存的使用情况。

We can [download the example configuration for production](https://raw.githubusercontent.com/bitnami/charts/master/bitnami/wordpress/values-production.yaml) from the maintainer of the WordPress Chart. The most important changes are:

我们可以从 WordPress Chart 的维护者处[下载用于生产的示例配置](https://raw.githubusercontent.com/bitnami/charts/master/bitnami/wordpress/values-production.yaml)。最重要的变化是：

```yaml
### Start 3 WordPress instances that will all receive
### requests from our visitors.A load-balancer will distribute calls
### to all containers evenly.
replicaCount: 3

### start a sidecar container that will expose metrics for your wordpress container
metrics:
enabled: true
image:
    registry: docker.io
    repository: bitnami/apache-exporter
    tag: 0.8.0-debian-10-r243
```

Let’s stop the default application:

让我们停止默认应用程序：

```shell
$ helm delete my-wordpress

release "my-wordpress" uninstalled
```

### How to Start a Multi-instance WordPress and MariaDB Deployment 

### 如何启动多实例 WordPress 和 MariaDB 部署

Deploy a new release using the production values:

使用生产值部署新版本：

```shell
$ helm install my-wordpress-prod bitnami/wordpress --version 10.1.4 -f values-production.yaml
```

This time, we have more containers running:

这一次，我们有更多的容器在运行：

```shell
$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
my-wordpress-prod-5c9776c976-4bs6f   2/2     Running   0          103s
my-wordpress-prod-5c9776c976-9ssmr   2/2     Running   0          103s
my-wordpress-prod-5c9776c976-sfq84   2/2     Running   0          103s
my-wordpress-prod-mariadb-0          1/1     Running   0          103s

```

We see 4 lines: 1 line for MariaDB, and 3 lines for our actual WordPress pods.

我们看到 4 行：1 行用于 MariaDB，3 行用于我们的实际 WordPress pod。

A pod in Kubernetes is a group of containers. Each group contains 2 containers, one for WordPress and one for an exporter for Prometheus that exposes valuable metrics in a special format.

Kubernetes 中的 Pod 是一组容器。每个组包含 2 个容器，一个用于 WordPress，一个用于 Prometheus 的导出器，以特殊格式公开有价值的指标。

As in the default setup, we can [open localhost](http://localhost) and play with our WordPress application.

在默认设置中，我们可以[打开 localhost](http://localhost) 并使用我们的 WordPress 应用程序。

### How to Access Exposed Health Metrics

### 如何访问公开的健康指标

We can check the exposed health metrics by proxying to one of the running pods:

我们可以通过代理到正在运行的 Pod 之一来检查公开的健康指标：

```shell
kubectl port-forward my-wordpress-prod-5c9776c976-sfq84 9117:9117
```

_Make sure to replace the pod-id with your own pod ID when you execute the port-forward command._

_执行port-forward命令时，请务必将pod-id替换为自己的pod ID。_

Now, we are connected to port 9117 of the WordPress Prometheus exporter and map the port to our local port 9117. Open [http://localhost:9117](http://localhost:9117/metrics) to check the output.

现在，我们连接到 WordPress Prometheus 导出器的端口 9117 并将端口映射到我们的本地端口 9117。打开 [http://localhost:9117](http://localhost:9117/metrics) 以检查输出。

If you are not used to the Prometheus format, it might be a little bit confusing in the beginning. But it’s actually pretty easy to read. Each line without _‘#’_ contains a metric key and a value behind it:

如果你不习惯 Prometheus 格式，一开始可能会有点混乱。但它实际上很容易阅读。没有 _‘#’_ 的每一行都包含一个度量键和它后面的值：

```prometheus
apache_cpuload 1.2766
process_resident_memory_bytes 1.6441344e+07
```

If you are not used to such metrics, don't worry – you will get used to them quickly. You can Google each of the keys and find out what it means. After some time, you will identify what metrics are the most valuable for you and how they behave as soon as your containers receive more and more production traffic.

如果您不习惯这些指标，请不要担心 - 您会很快适应它们。你可以谷歌每个键并找出它的含义。一段时间后，您将确定哪些指标对您最有价值，以及一旦您的容器收到越来越多的生产流量，它们的行为方式。

Let’s tidy up our setup by:

让我们整理一下我们的设置：

```shell
$ helm delete my-wordpress-prod

release "my-wordpress-prod" uninstalled
```

We touched on a lot of deployment areas and features. We deployed multiple WordPress instances and scaled it up to more containers for production. You could even go one step further and activate auto-scaling. Check out the documentation of the Helm Chart and play around with it!

我们触及了许多部署领域和功能。我们部署了多个 WordPress 实例并将其扩展到更多容器用于生产。您甚至可以更进一步并激活自动缩放。查看 Helm Chart 的文档并使用它！

### MariaDB Helm Chart

### MariaDB 掌舵图

The persistence of the helm Chart for WordPress depends on MariaDB. It builds on another [Helm Chart for MariaDB](https://artifacthub.io/packages/helm/bitnami/mariadb) that you can configure and scale to your needs by, for example, starting multiple replicas.

WordPress 掌舵图表的持久性取决于 MariaDB。它建立在另一个 [Helm Chart for MariaDB](https://artifacthub.io/packages/helm/bitnami/mariadb) 之上，您可以通过例如启动多个副本来配置和扩展以满足您的需求。

The possibilities that you have when running containers in production using Kubernetes are enormous. The definition of the WordPress Chart is publicly available.

在使用 Kubernetes 在生产中运行容器时，您拥有的可能性是巨大的。 WordPress 图表的定义是公开的。

In the next section, we will create our own Helm Chart with a basic application to understand the fundamentals of creating a Helm Chart and to make a static container deployment more dynamic.

在下一节中，我们将使用一个基本应用程序创建我们自己的 Helm Chart，以了解创建 Helm Chart 的基础知识并使静态容器部署更加动态。

## How to Create a Template for Custom Applications

## 如何为自定义应用程序创建模板

Helm adds a lot more flexibility to your Kubernetes deployment files. Kubernetes deployment files are static by their nature. This means, adjustments like

Helm 为您的 Kubernetes 部署文件增加了更多的灵活性。 Kubernetes 部署文件本质上是静态的。这意味着，调整如

- desired container count,
- environment variables or
- CPU and memory limit

- 所需的容器数量，
- 环境变量或
- CPU 和内存限制

are not adjustable by using plain Kubernetes deployment files. Either you solve this by duplicating configuration files or you put placeholders in your Kubernetes deployment files that are replaced at deploy-time.

不能通过使用普通的 Kubernetes 部署文件进行调整。您要么通过复制配置文件来解决此问题，要么将占位符放在部署时替换的 Kubernetes 部署文件中。

Both of these solutions require some additional work and will not scale well if you deploy a lot of applications with different variations.

这两种解决方案都需要一些额外的工作，如果您部署大量具有不同变体的应用程序，则无法很好地扩展。

But for sure, there is a smarter solution that is based on Helm that contains a lot of handy features from the Helm community. Let’s create a custom Chart for a blogging engine, this time for a NodeJS based blog called [ghost blog](https://ghost.org/).

但可以肯定的是，有一个基于 Helm 的更智能的解决方案，其中包含来自 Helm 社区的许多方便的功能。让我们为博客引擎创建一个自定义图表，这次是为一个名为 [ghost blog](https://ghost.org/) 的基于 NodeJS 的博客。

### How to Start a Ghost Blog Using Docker

### 如何使用 Docker 启动 Ghost 博客

A simple instance can be started using pure Docker:

可以使用纯 Docker 启动一个简单的实例：

```shell
docker run --rm -p 2368:2368 --name my-ghost ghost
```

Our blog is available at: [http://localhost:2368](http://localhost:2368/).

我们的博客位于：[http://localhost:2368](http://localhost:2368/)。

Let's stop the instance to be able to launch another one using Kubernetes:

让我们停止实例，以便能够使用 Kubernetes 启动另一个实例：

```shell
$ docker rm -f my-ghost

my-ghost
```

Now, we want to deploy the ghost blog with 2 instances in our Kubernetes cluster. Let’s set up a plain deployment first:

现在，我们要在 Kubernetes 集群中部署带有 2 个实例的 ghost 博客。让我们先设置一个简单的部署：

```yaml
# file 'application/deployment.yaml'

apiVersion: apps/v1
kind: Deployment
metadata:
name: ghost-app
spec:
selector:
    matchLabels:
      app: ghost-app
replicas: 2
template:
    metadata:
      labels:
        app: ghost-app
    spec:
      containers:
        - name: ghost-app
          image: ghost

          ports:
            - containerPort: 2368
```

and put a load balancer before it to be able to access our container and to distribute the traffic to both containers:

并在它之前放置一个负载均衡器，以便能够访问我们的容器并将流量分配给两个容器：

```yaml
# file 'application/service.yaml'

apiVersion: v1
kind: Service
metadata:
name: my-service-for-ghost-app
spec:
type: LoadBalancer
selector:
    app: ghost-app
ports:
    - protocol: TCP
      port: 80
      targetPort: 2368
```

We can now deploy both resources using kubectl:

我们现在可以使用 kubectl 部署这两种资源：

```
$ kubectl apply -f ./appplication/deployment.yaml -f ./appplication/service.yaml

deployment.apps/ghost-app created
service/my-service-for-ghost-app created
```

The ghost application is now available via [http://localhost](http://localhost). Let's again stop the application:

现在可以通过 [http://localhost](http://localhost) 使用 ghost 应用程序。让我们再次停止应用程序：

```
$ kubectl delete -f ./appplication/deployment.yaml -f ./appplication/service.yaml

deployment.apps/ghost-app delete
service/my-service-for-ghost-app delete
```

So far so good, it works with plain Kubernetes. But what if we need different settings for different environments?

到目前为止一切顺利，它适用于普通的 Kubernetes。但是如果我们需要针对不同的环境进行不同的设置呢？

Imagine that we want to deploy it to multiple data centers in different stages (non-prod, prod). You will end up duplicating your Kubernetes files over and over again. It will be hell for maintenance. Instead of scripting a lot, we can leverage Helm.

想象一下，我们想将它部署到不同阶段（非生产、生产）的多个数据中心。您最终将一遍又一遍地复制 Kubernetes 文件。这将是维护的地狱。我们可以利用 Helm，而不是编写大量脚本。

Let’s create a new Helm Chart from scratch:

让我们从头开始创建一个新的 Helm Chart：

```shell
$ helm create my-ghost-app

Creating my-ghost-app
```

Helm created a bunch of files for you that are usually important for a production-ready service in Kubernetes. To concentrate on the most important parts, we can remove a lot of the created files. Let’s go through the only required files for this example.

Helm 为您创建了一堆文件，这些文件对于 Kubernetes 中的生产就绪服务通常很重要。为了专注于最重要的部分，我们可以删除很多创建的文件。让我们来看看这个例子中唯一需要的文件。

We need a project file that is called Chart.yaml:

我们需要一个名为 Chart.yaml 的项目文件：

```yaml
# Chart.yaml

apiVersion: v2
name: my-ghost-app
description: A Helm chart for Kubernetes
type: application
version: 0.1.0
appVersion: 1.16.0
```

The deployment template file:

部署模板文件：

```yaml
# templates/deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
name: ghost-app
spec:
selector:
    matchLabels:
      app: ghost-app
replicas: {{ .Values.replicaCount }}
template:
    metadata:
      labels:
        app: ghost-app
    spec:
      containers:
        - name: ghost-app
          image: ghost
          ports:
            - containerPort: 2368
          env:
            - name: url
              {{- if .Values.prodUrlSchema }}
              value: http://{{ .Values.baseUrl }}
              {{- else }}
              value: http://{{ .Values.datacenter }}.non-prod.{{ .Values.baseUrl }}
              {{- end }}
```

It looks very similar to our plain Kubernetes file. Here, you can see different placeholders for the replica count, and an if-else condition for the environment variable called url. In the following files, we will see all the values defined.

它看起来与我们的普通 Kubernetes 文件非常相似。在这里，您可以看到副本计数的不同占位符，以及名为 url 的环境变量的 if-else 条件。在以下文件中，我们将看到定义的所有值。

The service template file:

服务模板文件：

```yaml
# templates/service.yaml

apiVersion: v1
kind: Service
metadata:
name: my-service-for-my-webapp
spec:
type: LoadBalancer
selector:
    app: ghost-app
ports:
    - protocol: TCP
      port: 80
      targetPort: 2368
```

Our Service configuration is completely static.

我们的服务配置是完全静态的。

The values for the templates are the last missing parts of our Helm Chart. Most importantly, there is a default values file required called values.yaml:

模板的值是 Helm Chart 中最后缺失的部分。最重要的是，需要一个名为 values.yaml 的默认值文件：

```yaml
# values.yaml

replicaCount: 1
prodUrlSchema: false
datacenter: us-east
baseUrl: myapp.org
```

A Helm Chart should be able to run just by using the default values. Before you proceed, make sure that you have deleted:

Helm Chart 应该能够仅通过使用默认值来运行。在继续之前，请确保您已删除：

- my-ghost-app/templates/tests/test-connection.yaml
- my-ghost-app/templates/serviceaccount.yaml
- my-ghost-app/templates/ingress.yaml
- my-ghost-app/templates/hpa.yaml
- my-ghost-app/templates/NOTES.txt.

- my-ghost-app/templates/tests/test-connection.yaml
- my-ghost-app/templates/serviceaccount.yaml
- my-ghost-app/templates/ingress.yaml
- my-ghost-app/templates/hpa.yaml
- my-ghost-app/templates/NOTES.txt。

We can get the final output that would be sent to Kubernetes by executing a “dry-run”:

我们可以通过执行“试运行”获得将发送到 Kubernetes 的最终输出：

```shell
$ helm template --debug my-ghost-app

install.go:159: [debug] Original chart version: ""
install.go:176: [debug] CHART PATH: /helm/my-ghost-app

---
# Source: my-ghost-app/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
name: my-service-for-my-webapp
spec:
type: LoadBalancer
selector:
    app: my-example-app
ports:
    - protocol: TCP
      port: 80
      targetPort: 2368
---
# Source: my-ghost-app/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
name: ghost-app
spec:
selector:
    matchLabels:
      app: ghost-app
replicas: 1
template:
    metadata:
      labels:
        app: ghost-app
    spec:
      containers:
        - name: ghost-app
          image: ghost
          ports:
            - containerPort: 2368
          env:
            - name: url
              value: us-east.non-prod.myapp.org

```

Helm inserted all the values and also set url to _`us-east.non-prod.myapp.org`_ because in the _`values.yaml`_, `prodUrlSchema` is set to false and the datacenter is set to us -east.

Helm 插入了所有值并将 url 设置为 _`us-east.non-prod.myapp.org`_ 因为在 _`values.yaml`_ 中，`prodUrlSchema` 设置为 false，数据中心设置为我们-东。

To get some more flexibility, we can define some override value files. Let’s define one for each datacenter:

为了获得更大的灵活性，我们可以定义一些覆盖值文件。让我们为每个数据中心定义一个：

```yaml
# values.us-east.yaml
datacenter: us-east
```

```yaml
# values.us-west.yaml
datacenter: us-west
```

and one for each stage:

每个阶段一个：

```yaml
# values.nonprod.yaml
replicaCount: 1
prodUrlSchema: false
```

```yaml
# values.prod.yaml
replicaCount: 3
prodUrlSchema: true
```

We can now use Helm to combine them as we want and check the result again:

我们现在可以根据需要使用 Helm 组合它们并再次检查结果：

```shell
$ helm template --debug my-ghost-app -f my-ghost-app/values.nonprod.yaml  -f my-ghost-app/values.us-east.yaml

install.go:159: [debug] Original chart version: ""
install.go:176: [debug] CHART PATH: /helm/my-ghost-app

---
# Source: my-ghost-app/templates/service.yaml
# templates/service.yaml

apiVersion: v1
kind: Service
metadata:
name: my-service-for-my-webapp
spec:
type: LoadBalancer
selector:
    app: my-example-app
ports:
    - protocol: TCP
      port: 80
      targetPort: 2368
---
# Source: my-ghost-app/templates/deployment.yaml
# templates/deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
name: ghost-app
spec:
selector:
    matchLabels:
      app: ghost-app
replicas: 1
template:
    metadata:
      labels:
        app: ghost-app
    spec:
      containers:
        - name: ghost-app
          image: ghost
          ports:
            - containerPort: 2368
          env:
            - name: url
              value: http://us-east.non-prod.myapp.org

```

And for sure, we can do a final deployment:

可以肯定的是，我们可以进行最终部署：

```shell
$ helm install -f my-ghost-app/values.prod.yaml my-ghost-prod ./my-ghost-app/

NAME: my-ghost-prod
LAST DEPLOYED: Mon Dec 21 00:09:17 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

As before, our ghost blog is available via [http://localhost](http://localhost).

和以前一样，我们的幽灵博客可以通过 [http://localhost](http://localhost) 访问。

We can delete this deployment and deploy the application with us-east and non prod settings like this:

我们可以删除此部署并使用 us-east 和 non prod 设置部署应用程序，如下所示：

```shell
$ helm delete my-ghost-prod
release "my-ghost-prod" uninstalled

$ helm install -f my-ghost-app/values.nonprod.yaml -f my-ghost-app/values.us-east.yaml my-ghost-nonprod ./my-ghost-app
```

We finally clean up our Kubernetes deployment via Helm:

我们最终通过 Helm 清理了我们的 Kubernetes 部署：

```shell
$ helm delete my-ghost-nonprod
```

So we can combine multiple override value files as we want. We can automate deployments in a flexible way that we need for many use-cases of deployment pipelines.

所以我们可以根据需要组合多个覆盖值文件。我们可以以灵活的方式自动化部署，这是许多部署管道用例所需要的。

Especially for companies, this means defining Chart Skeletons once to ensure the required criteria are fulfilled. Later, you can copy them and adjust them to the needs of your application.

特别是对于公司而言，这意味着定义图表骨架一次以确保满足所需的标准。稍后，您可以复制它们并根据应用程序的需要进行调整。

## Conclusion

##  结论

The power of a great templating engine and the possibility of executing releases, upgrades, and rollbacks makes Helm great. On top of that comes the publicly available Helm Chart Hub that contains thousands of production-ready templates. This makes Helm a must-have tool in your toolbox if you work with Kubernetes on a bigger scale!

强大的模板引擎的强大功能以及执行发布、升级和回滚的可能性使 Helm 变得很棒。最重要的是公开可用的 Helm Chart Hub，其中包含数以千计的生产就绪模板。如果您在更大范围内使用 Kubernetes，这使得 Helm 成为您工具箱中的必备工具！

I hope you enjoyed this hands-on tutorial. Motivate yourself to Google around, check out other examples, deploy containers, connect them, and use them.

我希望你喜欢这个动手教程。激励自己使用 Google，查看其他示例，部署容器，连接它们并使用它们。

You will learn many cool features in the future that enable you to ship your application to production in an effortless, reusable and scalable way.

将来您将学习许多很酷的功能，这些功能使您能够以轻松、可重用和可扩展的方式将应用程序交付到生产环境中。

As always, I appreciate any feedback and comments. I hope you enjoyed the article. If you like it and feel the need for a round of applause, [follow me on Twitter](https://twitter.com/journerist). I work at eBay Kleinanzeigen, one of the biggest classified companies globally. By the way, [we are hiring](https://jobs.ebayclassifiedsgroup.com/ebay-kleinanzeigen)!

与往常一样，我感谢任何反馈和评论。我希望你喜欢这篇文章。如果您喜欢它并觉得需要掌声，请[在 Twitter 上关注我](https://twitter.com/journerist)。我在 eBay Kleinanzeigen 工作，这是全球最大的分类公司之一。顺便说一下，[我们正在招聘](https://jobs.ebayclassifiedsgroup.com/ebay-kleinanzeigen)！

References:

参考：

- [https://helm.sh/docs/chart\_template\_guide/getting\_started/](https://helm.sh/docs/chart_template_guide/getting_started/)

- [https://helm.sh/docs/chart\_template\_guide/getting\_started/](https://helm.sh/docs/chart_template_guide/getting_started/)

* * *

* * *

![Sebastian Sigl](http://www.freecodecamp.org/news/content/images/size/w100/2020/01/profile_photo_edited.png)[SebastianSigl](http://www.freecodecamp.org/news/author/sesigl/)

/作者/sesigl/)

Software engineer who loves writing software and teaching people. Since 2018, I work for eBay Kleinanzeigen that is nowadays part of Adevinta. 

喜欢编写软件和教人的软件工程师。自 2018 年以来，我为 eBay Kleinanzeigen 工作，该公司现在是 Adevinta 的一部分。

