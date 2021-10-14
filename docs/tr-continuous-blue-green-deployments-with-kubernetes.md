# Continuous Blue-Green Deployments With Kubernetes

# 使用 Kubernetes 进行持续蓝绿部署

8 Sep 2020 · [Software Engineering](https://semaphoreci.com/category/engineering)

You can also download this article as a PDF and read it on the couch.

您还可以将本文下载为 PDF 文件并在沙发上阅读。

[Download](https://semaphoreci.com/resources/download-blue-green-kubernetes)

[下载](https://semaphoreci.com/resources/download-blue-green-kubernetes)

_Do you know what airplanes, rockets, submarines, and blue-green deployments have in common? They all go to great lengths to prevent failures. And they do that using redundancy._

_你知道飞机、火箭、潜艇、蓝绿部署有什么共同点吗？他们都竭尽全力防止失败。他们使用冗余来做到这一点。_

We’ve talked before about the generalities of blue-green deployments in [another post](https://semaphoreci.com/blog/blue-green-deployment). Today, I’d like to get into the gory details and see how we can create a CI/CD pipeline that deploys a Kubernetes application using the blue-green methodology.

我们之前在 [另一篇文章](https://semaphoreci.com/blog/blue-green-deployment) 中讨论了蓝绿部署的一般性。今天，我想深入了解细节，看看我们如何创建 CI/CD 管道，使用蓝绿方法部署 Kubernetes 应用程序。

The gist of blue-green deployments is to have two identical environments, conventionally called blue and green, to do continuous, risk-free updates. This way, users access one while the other receives updates.

蓝绿部署的要点是拥有两个相同的环境，通常称为蓝色和绿色，以进行持续、无风险的更新。这样，用户访问一个，而另一个接收更新。

![Blue-green deployments at glance](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/bg-overview-1024x776.png)Blue-green deployments at a glance

Blue and green take turns. On each cycle, we deploy new versions into the idle environment, test them, and finally switch routes so all users can start using it. With this method, we get three benefits:

蓝色和绿色轮流。在每个周期中，我们将新版本部署到空闲环境中，对其进行测试，最后切换路由，以便所有用户都可以开始使用它。使用这种方法，我们可以获得三个好处：

- We test in a real production environment.
- Users don’t experience any downtime.
- We can rollback in an instant in case there is trouble.

- 我们在真实的生产环境中进行测试。
- 用户不会遇到任何停机时间。
- 万一出现问题，我们可以立即回滚。

## Why Kubernetes?

## 为什么是 Kubernetes？

The way we manage the infrastructure around blue-green deployments depends on the technology we’re running. If we’re using bare metal servers, one system will be idle most of the time. In practice, however, it’s a lot more common and cost-effective to provision resources on-demand in the cloud using _infrastructure as code_ (IaC). For instance, we can start virtual machines and spin up containers, configure networks and services before we start the deployment. Once users have been switched to the new version, the old environment can be torn down.

我们围绕蓝绿部署管理基础设施的方式取决于我们正在运行的技术。如果我们使用裸机服务器，一个系统将在大部分时间处于空闲状态。然而，在实践中，使用 _infrastructure as code_ (IaC) 在云中按需供应资源更为常见且更具成本效益。例如，我们可以在开始部署之前启动虚拟机并启动容器、配置网络和服务。一旦用户切换到新版本，旧环境就可以拆除。

Here is where Kubernetes enters the picture. Kubernetes is an orchestration platform that’s perfect for blue-green deployments. We can, for instance, use the platform to dynamically create the green environment, deploy the application, switch over the user’s traffic, and finally delete the blue environment. Kubernetes lets us manage the whole blue-green process using one tool.

这就是 Kubernetes 进入画面的地方。 Kubernetes 是一个非常适合蓝绿部署的编排平台。比如我们可以通过平台动态创建绿色环境，部署应用，切换用户流量，最后删除蓝色环境。 Kubernetes 让我们可以使用一种工具管理整个蓝绿过程。

If you are an absolute beginner in Kubernetes or if you'd like a refresher, grab a copy of our free eBook [CI/CD with Docker and Kubernetes](https://semaphoreci.com/resources/cicd-docker-kubernetes) . It is a great way to get started with Docker and Kubernetes.

如果您是 Kubernetes 的绝对初学者，或者您想复习一下，请获取我们的免费电子书 [使用 Docker 和 Kubernetes 的 CI/CD](https://semaphoreci.com/resources/cicd-docker-kubernetes) 的副本.这是开始使用 Docker 和 Kubernetes 的好方法。

## Blue-green Deployments with Kubernetes

## 使用 Kubernetes 进行蓝绿部署

Let’s see Kubernetes blue-green deployments in action. Imagine we have version `v1` of awesome application called `myapp`, and that is currently running in blue. In Kubernetes, we run applications with deployments and pods.

让我们看看 Kubernetes 蓝绿部署的实际效果。想象一下，我们有一个名为“myapp”的很棒的应用程序的“v1”版本，它目前以蓝色运行。在 Kubernetes 中，我们运行带有部署和 Pod 的应用程序。

![V1 deployment running in blue](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/bg1-1024x544.png)v1 deployment running in blue

Sometime later, we have the next version ( `v2`) ready to go. So we create a brand-new production environment called green. As it turns out, in Kubernetes we only have to declare a new deployment, and the platform takes care of the rest. Users are not yet aware of the change as the blue environment keeps on working unaffected. They won’t see any change until we switch traffic over from blue to green.

稍后，我们准备好了下一个版本（`v2`）。所以我们创造了一个全新的生产环境，叫做绿色。事实证明，在 Kubernetes 中，我们只需要声明一个新的部署，平台会负责其余的工作。用户还没有意识到这一变化，因为蓝色环境继续工作而不受影响。在我们将流量从蓝色切换到绿色之前，他们不会看到任何变化。

![A new deployment is created to run V2](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/bg2-1024x937.png)A new deployment is created to run v2

It’s said that only developers that like to live dangerously test in production. But here we all have the chance to do that without risks. We can test green at leisure on the same Kubernetes cluster where blue is running.

据说只有喜欢在生产中进行危险测试的开发人员。但在这里，我们都有机会在没有风险的情况下做到这一点。我们可以在运行 blue 的同一个 Kubernetes 集群上空闲时测试 green。

![v2 is active on green, v1 is on stand-by on blue](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/bg3-1024x937.png)v2 is active on green,v1 is on stand-by on blue

v1 处于蓝色待机状态

Once we have moved the users from blue to green and happy with the result, we can delete blue to free up resources.

一旦我们将用户从蓝色移动到绿色并对结果感到满意，我们就可以删除蓝色以释放资源。

![blue deployment is gone](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/bg4-1024x544.png)Blue deployment is gone

As you can imagine, blue-green deployments are complex. We have to grapple with two deployments at once and manage the network. Fortunately, Kubernetes makes things a lot easier. Even so, we should strive to automate the release cycle as much as possible. In this tutorial, I'm going to show you how to use Semaphore [Continuous Integration](https://semaphoreci.com/continuous-integration)(CI) and [Continuous Delivery](https://semaphoreci.com/cicd) (CD) to test and release any project.

可以想象，蓝绿部署是复杂的。我们必须同时处理两个部署并管理网络。幸运的是，Kubernetes 使事情变得容易多了。即便如此，我们还是应该努力使发布周期尽可能自动化。在本教程中，我将向您展示如何使用信号量 [Continuous Integration](https://semaphoreci.com/continuous-integration)(CI) 和 [Continuous Delivery](https://semaphoreci.com/cicd) (CD) 测试和发布任何项目。

## Getting Ready

## 准备

I’ll try to avoid going into cloud vendor-specific details. Technically, Kubernetes’ behavior itself doesn’t depend on the provider, so this guide should work the same for any application on Google Cloud, AWS, Azure, or any other cloud. However, the commands used to connect to the Kubernetes cluster will change.

我会尽量避免涉及特定于云供应商的细节。从技术上讲，Kubernetes 的行为本身不依赖于提供者，因此本指南应该适用于 Google Cloud、AWS、Azure 或任何其他云上的任何应用程序。但是，用于连接到 Kubernetes 集群的命令将发生变化。

That being said, we’ll need to define some common ground. You’ll need:

话虽如此，我们需要定义一些共同点。你需要：

- A Kubernetes cluster running with [Istio](https://istio.io/). Istio is a service mesh that adds many features to Kubernetes.
- A Docker registry to store the container images. We’ll use[Docker Hub](https://hub.docker.com/) because it’s the default. Most cloud providers also offer private registries that may be more convenient for you.
- An application and its associated Dockerfile. We’ll use the[semaphore-demo-cicd-kubernetes](https://github.com/semaphoreci-demos/semaphore-demo-cicd-kubernetes) demo project. You’re welcome to **fork** it and play with it.
- The[kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) CLI and your cluster's [kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/) (plus any other tools needed to manage it).

- 使用 [Istio](https://istio.io/) 运行的 Kubernetes 集群。 Istio 是一个服务网格，它为 Kubernetes 增加了许多功能。
- 用于存储容器映像的 Docker 注册表。我们将使用 [Docker Hub](https://hub.docker.com/) 因为它是默认设置。大多数云提供商还提供可能对您更方便的私有注册表。
- 应用程序及其关联的 Dockerfile。我们将使用 [semaphore-demo-cicd-kubernetes](https://github.com/semaphoreci-demos/semaphore-demo-cicd-kubernetes) 演示项目。欢迎您 ** fork ** 它并使用它。
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) CLI 和集群的 [kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)（以及管理它所需的任何其他工具)。

### Preparing the Manifests

### 准备清单

In Kubernetes, we use manifests to describe what we want and let the platform figure out the rest.

在 Kubernetes 中，我们使用清单来描述我们想要什么，让平台来解决剩下的问题。

I’ve split manifests into three parts:

我将清单分为三个部分：

- **Gateway**: the entry point for the application. Accepts HTTP requests.
- **Routing**: describes the routes that send requests to blue and green.
- **Deployments**: describes the pods that run the application. We’ll have a deployment for each color.

- **网关**：应用程序的入口点。接受 HTTP 请求。
- **Routing**：描述了向蓝绿发送请求的路由。
- **部署**：描述运行应用程序的 Pod。我们将为每种颜色进行部署。

To get started, create a directory called `manifests` in the root of your project. Place the files as shown in the next section. These manifests were designed to work with our demo project, so you may need to adjust them for your use case.

首先，在项目的根目录中创建一个名为“manifests”的目录。放置文件，如下一节所示。这些清单旨在与我们的演示项目配合使用，因此您可能需要针对您的用例调整它们。

```
manifests/
├── deployment.yml
├── gateway.yml
├── route-test.yml
├── route.yml
└── service.yml

```

Once you are done editing the manifests, push them to the repo:

完成清单编辑后，将它们推送到存储库：

```bash
$ git add manifests/*
$ git commit -m "add Kubernetes manifests"
$ git push origin master
```

### Istio Manifests

### Istio 清单

Istio is an open-source service-mesh platform designed to run on top of products such as Kubernetes and Consul. This service is a popular choice for running microservice applications because it facilitates communication and provides security. Compared with native Kubernetes controllers, Istio’s service mesh gives us more control and flexibility. We’ll rely on Istio for handling all network traffic.

Istio 是一个开源服务网格平台，旨在运行在 Kubernetes 和 Consul 等产品之上。此服务是运行微服务应用程序的流行选择，因为它促进了通信并提供了安全性。与原生 Kubernetes 控制器相比，Istio 的服务网格为我们提供了更多的控制和灵活性。我们将依靠 Istio 来处理所有网络流量。

**Gateway**

**网关**

An [Istio Ingress](https://istio.io/latest/docs/tasks/traffic-management/ingress/ingress-control/) gateway is a resource that processes traffic entering the cluster. Istio gateways describe a load balancer at the edge of the service mesh. We can use them to encrypt connections and expose ports to the Internet.

[Istio Ingress](https://istio.io/latest/docs/tasks/traffic-management/ingress/ingress-control/) 网关是一种处理进入集群的流量的资源。 Istio 网关描述了服务网格边缘的负载均衡器。我们可以使用它们来加密连接并向 Internet 公开端口。

The following gateway accepts HTTP (port 80) connections from all hosts.

以下网关接受来自所有主机的 HTTP（端口 80）连接。

```yaml
# manifests/gateway.yml

apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
name: myapp-gateway
spec:
selector:
    istio: ingressgateway
servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*"
```

**Routing**

**路由**

Before going into the routing let me tell you the variable convention we’ll use from now on:

在进入路由之前，让我告诉你我们将从现在开始使用的变量约定：

- `$COLOR_ACTIVE` is the live, in-production deployment. We’ll route all user connections into it by default. This environment runs the old/stable version.
- `$COLOR_TEST` is where we deploy the new version and where we run the tests. It temporarily exists while we’re making a deployment.

- `$COLOR_ACTIVE` 是实时的、生产中的部署。默认情况下，我们会将所有用户连接路由到它。此环境运行旧/稳定版本。
- `$COLOR_TEST` 是我们部署新版本和运行测试的地方。它在我们进行部署时暂时存在。

When one variable is `green`, the other is `blue` and vice-versa.

当一个变量为“green”时，另一个变量为“blue”，反之亦然。

We’ll need two Istio resources to route inbound traffic: 

我们需要两个 Istio 资源来路由入站流量：

- [VirtualService](https://istio.io/latest/docs/reference/config/networking/virtual-service/): binds to the Istio gateway and uses rules to decide where to route the requests.
- [DestinationRule](https://istio.io/latest/docs/reference/config/networking/destination-rule/): maps VirtualService rules with deployments using labels.

- [VirtualService](https://istio.io/latest/docs/reference/config/networking/virtual-service/)：绑定到 Istio 网关并使用规则来决定将请求路由到哪里。
- [DestinationRule](https://istio.io/latest/docs/reference/config/networking/destination-rule/)：使用标签将 VirtualService 规则与部署映射。

**Steady State Routes**

**稳定的状态路线**

The following manifest describes the controllers that send all traffic from the gateway to the `$COLOR_ACTIVE` deployment:

以下清单描述了将所有流量从网关发送到 `$COLOR_ACTIVE` 部署的控制器：

```yaml
# manifests/route.yml

apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
name: myapp-blue-green
spec:
hosts:
  - "*"
gateways:
  - myapp-gateway
http:
  - name: myapp-default
    route:
    - destination:
        host: myapp
        subset: $COLOR_ACTIVE

---

apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
name: myapp-blue-green
spec:
host: myapp
subsets:
  - name: $COLOR_ACTIVE
    labels:
      color: $COLOR_ACTIVE
```

**Test Routes**

**测试路线**

During deployment things are different. We’ll need to split traffic in two. We want regular users to maintain their access to the old version, running in `$COLOR_ACTIVE`. At the same time, we’d like to have a special route for us to run some tests on the new version, running in `$COLOR_TEST`, before making the switch. Istio has several [routing options](https://istio.io/latest/docs/reference/config/networking/virtual-service/) to achieve this. For HTTP traffic, I find that the easiest one is to use a cookie.

在部署期间，事情是不同的。我们需要将流量分成两部分。我们希望普通用户保持对旧版本的访问，运行在 `$COLOR_ACTIVE` 中。同时，我们希望有一个特殊的路线让我们在进行切换之前在新版本上运行一些测试，在 `$COLOR_TEST` 中运行。 Istio 有几个 [路由选项](https://istio.io/latest/docs/reference/config/networking/virtual-service/) 来实现这一点。对于 HTTP 流量，我发现最简单的方法是使用 cookie。

The next manifest describes a VirtualService and DestinationRule that routes requests with a cookie having `test=true` to `$COLOR_TEST`. The rest of the traffic, that is, requests without the cookie, goes to `$COLOR_ACTIVE`. We’ll call this manifest `manifests/route-test.yml`.

下一个清单描述了一个 VirtualService 和 DestinationRule，它使用具有 `test=true` 的 cookie 将请求路由到 `$COLOR_TEST`。其余的流量，即没有 cookie 的请求，将转到 `$COLOR_ACTIVE`。我们将这个清单命名为“manifests/route-test.yml”。

```yaml
# manifests/route-test.yml

apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
name: myapp-blue-green
spec:
hosts:
  - "*"
gateways:
  - myapp-gateway
http:
  - name: myapp-test
    match:
    - headers:
        cookie:
          regex: "^(.*?;)?(test=true)(;.*)?$"
    route:
    - destination:
        host: myapp
        subset: $COLOR_TEST

  - name: myapp-default
    route:
    - destination:
        host: myapp
        subset: $COLOR_ACTIVE

---

apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
name: myapp-blue-green
spec:
host: myapp
subsets:
  - name: blue
    labels:
      color: blue
  - name: green
    labels:
      color: green
```

### Kubernetes Manifests

### Kubernetes 清单

The remaining manifests describe native Kubernetes resources. They do not depend on Istio. These are the last manifests that complete the Kubernetes setup.

其余清单描述了原生 Kubernetes 资源。它们不依赖于 Istio。这些是完成 Kubernetes 设置的最后一个清单。

**Deployments & Service**

**部署和服务**

First, we’ll define a [deployment](https://semaphoreci.com/blog/kubernetes-deployment), which creates the application’s pods. It includes a [readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-liveness-http-request) and some database connection variables.

首先，我们将定义一个 [deployment](https://semaphoreci.com/blog/kubernetes-deployment)，它会创建应用程序的 pod。它包括一个 [readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-liveness-http-request) 和一些数据库连接变量。

```yaml
# manifests/deployment.yml

apiVersion: apps/v1
kind: Deployment
metadata:
name: myapp-$COLOR_TEST
labels:
    app: myapp
    color: $COLOR_TEST
spec:
replicas: 1
selector:
    matchLabels:
      app: myapp
      color: $COLOR_TEST
strategy:
    type: Recreate
template:
    metadata:
      labels:
        app: myapp
        color: $COLOR_TEST
    spec:
      imagePullSecrets:
      - name: dockerhub
      containers:
      - name: myapp
        image: $DOCKER_USERNAME/myapp:$SEMAPHORE_WORKFLOW_ID
        ports:
        - containerPort: 3000
        readinessProbe:
          httpGet:
            path: /ready
            port: 3000
        env:
        - name: NODE_ENV
          value: "production"
        - name: DB_HOST
          value: "$DB_HOST"
        - name: DB_PORT
          value: "$DB_PORT"
        - name: DB_SCHEMA
          value: "$DB_SCHEMA"
        - name: DB_USER
          value: "$DB_USER"
        - name: DB_PASSWORD
          value: "$DB_PASSWORD"
```

Second, you’ll need to store your docker registry credentials in the Kubernetes cluster. To do this run the following command. You only have to do this once:

其次，您需要将 docker 注册表凭据存储在 Kubernetes 集群中。为此，请运行以下命令。你只需要做一次：

```bash
$ kubectl create secret docker-registry dockerhub \
--docker-server=docker.io \
--docker-username=YOUR_DOCKER_HUB_USERNAME \
--docker-password=YOUR_DOCKER_HUB_PASSWORD

secret/dockerhub created
```

Finally, we’ll use a service to get a stable IP and hostname for the application. This [service](https://kubernetes.io/docs/concepts/services-networking/service/) targets the application pods labeled as `app = myapp`.

最后，我们将使用服务为应用程序获取稳定的 IP 和主机名。此 [service](https://kubernetes.io/docs/concepts/services-networking/service/) 以标记为“app = myapp”的应用程序 pod 为目标。

```yaml
# manifests/service.yml

apiVersion: v1
kind: Service
metadata:
name: myapp
labels:
    app: myapp
spec:
selector:
    app: myapp
ports:
  - port: 3000
    name: http
```

## Setting Up Your Continuous Integration Pipelines

## 设置您的持续集成管道

I will assume that you already have working continuous integration and delivery pipelines configured in Semaphore, which should [build and test](https://semaphoreci.com/blog/build-stage) your Docker images. The only requirement is that, at some point, it pushes the image into the registry of your choice.

我将假设您已经在 Semaphore 中配置了工作持续集成和交付管道，这应该[构建和测试](https://semaphoreci.com/blog/build-stage)您的 Docker 镜像。唯一的要求是，在某个时候，它会将映像推送到您选择的注册表中。

If you need help setting up your pipelines, you can find detailed step-by-step instructions in our free eBook [CI/CD with Docker and Kubernetes](https://semaphoreci.com/resources/cicd-docker-kubernetes). We also have detailed tutorials on dockerizing applications:

如果您在设置管道方面需要帮助，可以在我们的免费电子书 [带有 Docker 和 Kubernetes 的 CI/CD](https://semaphoreci.com/resources/cicd-docker-kubernetes) 中找到详细的分步说明。我们还有关于 dockerizing 应用程序的详细教程：

- [Dockerizing a Ruby on Rails Application](https://semaphoreci.com/community/tutorials/dockerizing-a-ruby-on-rails-application)
- [Dockerizing a Node.js Web Application](https://semaphoreci.com/community/tutorials/dockerizing-a-node-js-web-application)
- [Dockerizing a Python Django Web Application](https://semaphoreci.com/community/tutorials/dockerizing-a-python-django-web-application)
- [Dockerizing a PHP Application](https://semaphoreci.com/community/tutorials/dockerizing-a-php-application)
- [Continuous Integration with Deno](https://semaphoreci.com/blog/continuous-integration-with-deno)

- [Dockerizing Ruby on Rails 应用程序](https://semaphoreci.com/community/tutorials/dockerizing-a-ruby-on-rails-application)
- [Dockerizing Node.js Web 应用程序](https://semaphoreci.com/community/tutorials/dockerizing-a-node-js-web-application)
- [Dockerizing Python Django Web 应用程序](https://semaphoreci.com/community/tutorials/dockerizing-a-python-django-web-application)
- [Dockerizing PHP 应用程序](https://semaphoreci.com/community/tutorials/dockerizing-a-php-application)
- [与 Deno 的持续集成](https://semaphoreci.com/blog/continuous-integration-with-deno)

## How to Organize Releases

## 如何组织发布

There are probably as many ways of making releases as there are developers, after all, that’s our thing. But I think it’s a pretty safe bet that [Git tags](https://git-scm.com/book/en/v2/Git-Basics-Tagging) will be somehow involved. So, let’s mark releases using color-coded Git tags.

发布版本的方式可能和开发人员一样多，毕竟这是我们的事情。但我认为 [Git 标签](https://git-scm.com/book/en/v2/Git-Basics-Tagging) 会以某种方式参与其中是一个非常安全的赌注。因此，让我们使用颜色编码的 Git 标签来标记版本。

### Step 1: Decide Which Pipeline Should Start

### 第 1 步：决定应该启动哪个管道

Suppose we want to deploy a new version of the application into green. The old version is currently running on blue. The deployment starts once the continuous integration pipeline is done building the docker image.

假设我们要将应用程序的新版本部署为绿色。旧版本目前以蓝色运行。一旦持续集成管道完成构建 docker 镜像，部署就会开始。

The first step is to decide which deployment pipeline to start: the green or the blue. As I said, we’ll make the decision based on how the release was tagged. We’ll use a regular expression to find if the tag contains either blue or green, and activate the appropriate pipeline. If there are no tags, or they don’t match a color, nothing happens.

第一步是决定启动哪个部署管道：绿色或蓝色。正如我所说，我们将根据发布的标记方式做出决定。我们将使用正则表达式来查找标签是否包含蓝色或绿色，并激活相应的管道。如果没有标签，或者它们与颜色不匹配，则不会发生任何事情。

Imagine that we tagged our release as `v2.0-green`. Since it matches “green”, the green deployment pipeline is activated.

想象一下，我们将我们的版本标记为“v2.0-green”。由于它匹配“green”，绿色部署管道被激活。

### Step 2: Deploy

### 第 2 步：部署

The second step is to make the deployment. Here, we create the green pods with the new version of the Docker image. Additionally, we create a test route in the VirtualService.

第二步是进行部署。在这里，我们使用新版本的 Docker 镜像创建绿色 Pod。此外，我们在 VirtualService 中创建了一个测试路由。

### Step 3: Test the Deployment

### 第 3 步：测试部署

The main benefit of blue-green deployments is that we can [test](https://semaphoreci.com/blog/automated-testing-cicd) the application in a real production setting. The third step is to run tests on the new deployment.

蓝绿部署的主要好处是我们可以在真实的生产环境中 [测试](https://semaphoreci.com/blog/automated-testing-cicd) 应用程序。第三步是在新部署上运行测试。

To run the tests, we make HTTP requests using a cookie. The test route sends them to the green deployment.

为了运行测试，我们使用 cookie 发出 HTTP 请求。测试路由将它们发送到绿色部署。

### Step 4: Go Live

### 第 4 步：上线

If all tests pass, the fourth step is to change the default route, so all users access the new version. We do this by updating the VirtualService default route.

如果所有测试都通过，第四步是更改默认路由，让所有用户访问新版本。我们通过更新 VirtualService 默认路由来做到这一点。

Once we changed the default route to green, all users access the new deployment. Meanwhile, the old version is still running, nothing has changed on blue, other than it is no longer receiving any traffic.

一旦我们将默认路由更改为绿色，所有用户都可以访问新部署。同时，旧版本仍在运行，蓝色没有任何变化，除了不再接收任何流量。

### Step 5: Cleanup or Rollback

### 第 5 步：清理或回滚

At this point, the deployment is mostly complete. The only thing left is to do a cleanup. We can either delete the old environment or rollback. This will be the only manual step in the workflow. If for any reason, we are not satisfied with the new version, doing a rollback is easy, we just need to change the VirtualService route back to the blue.

至此，部署基本完成。唯一剩下的就是做一个清理工作。我们可以删除旧环境或回滚。这将是工作流程中唯一的手动步骤。如果因为任何原因，我们对新版本不满意，回滚很容易，我们只需要将 VirtualService 路由改回蓝色即可。

On the other hand, if the new version works perfectly, we can start the cleanup pipeline that deletes the blue deployment. This step releases the computing resources back to the cluster.

另一方面，如果新版本运行良好，我们可以启动删除蓝色部署的清理管道。此步骤将计算资源释放回集群。

### Deployment Flowchart

### 部署流程图

We can represent the sequence more visually using a flowchart.

我们可以使用流程图更直观地表示序列。

![deployment flowchart](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/flowchart-1024x543.png)Deployment flowchart

## Connecting Semaphore with Kubernetes 

## 将信号量与 Kubernetes 连接

Semaphore needs access to the Kubernetes cluster to make the deployment. To do this, follow these steps to create a [secret](https://docs.semaphoreci.com/essentials/using-secrets/):

信号量需要访问 Kubernetes 集群才能进行部署。为此，请按照以下步骤创建 [secret](https://docs.semaphoreci.com/essentials/using-secrets/)：

- Click on the account badge on the top-right corner and enter the**Settings** section.

- 点击右上角的帐户徽章并进入**设置**部分。

![blue-green](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/settings.png)

- Click on**Secrets**, then **New Secret**.

- 点击**Secrets**，然后点击**New Secret**。

![secrets](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/new-secret.png)

- Create a secret called “kubeconfig”.
- Upload your kubeconfig or any other files needed to connect.

- 创建一个名为“kubeconfig”的秘密。
- 上传您的 kubeconfig 或连接所需的任何其他文件。

![create a secret](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/kubeconfig.png)

The details depend on where the cluster is hosted; some cloud providers let you download the kubeconfig directly. Others require additional steps like installing and running dedicated CLIs. If you have any trouble setting this up, check out the links at the end of this post, you’ll find examples with various providers.

详细信息取决于集群的托管位置；一些云提供商允许您直接下载 kubeconfig。其他人需要额外的步骤，如安装和运行专用 CLI。如果您在设置时遇到任何问题，请查看本文末尾的链接，您会找到各种提供商的示例。

Next, create two sets of environment variables for your application:

接下来，为您的应用程序创建两组环境变量：

- Create a new secret called “env-blue”.
- Add all the environment variables the application needs, for instance, the database connection parameters.
- Click on**Save Changes**.

- 创建一个名为“env-blue”的新机密。
- 添加应用程序需要的所有环境变量，例如数据库连接参数。
- 单击**保存更改**。

![create a secret 2](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/env-blue.png)

Repeat the steps to create the green secret.

重复这些步骤以创建绿色秘密。

![create the green secret](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/env-green.png)

## Blue Pipelines

## 蓝色管道

It looks like we’re finally ready to start a deployment. We’ll do blue first.

看起来我们终于准备好开始部署了。我们先做蓝色。

To begin, open the Workflow Editor.

首先，打开工作流编辑器。

![workflow editor](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/builder.png)

- Select the branch.
- Press**Add First Promotion**.

- 选择分支。
- 按**添加第一个促销**。

![select the branch](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/promotion1.png)

- Change the name to “Deploy to Blue”.
- Check the**Enable automatic promotion** option.

- 将名称更改为“部署到蓝色”。
- 勾选**启用自动升级**选项。

![deploy to blue](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/blue1-1024x436.png)Deploy to blue promotion

- Type the following conditions to follow releases tagged like`v1.2.3-blue`:

- 输入以下条件以关注标记为“v1.2.3-blue”的版本：

```
result = 'passed' AND tag =~ '^v.*-blue$'
```

### Adding a Sanity Check

### 添加健全性检查

We’ll use the first block to do a sanity check. We have to make sure that we’re not deploying the new version into a live environment, which would be disastrous. This could happen if we Git-tagged the wrong color by mistake. The sanity check consists of retrieving the VirtualService default route and verifying that IS NOT blue.

我们将使用第一个块进行完整性检查。我们必须确保我们不会将新版本部署到实时环境中，这将是灾难性的。如果我们错误地在 Git 上标记了错误的颜色，就会发生这种情况。健全性检查包括检索 VirtualService 默认路由并验证 IS NOT 蓝色。

- Open the**prologue** section. Place here any commands that you need to connect with your cluster.
- Open the**environment variables** section and create a variable named `COLOR_TEST = blue`.
- Import the`kubeconfig` secret.
- Type the following command (sorry, it’s rather long):

- 打开**序幕**部分。在此处放置与集群连接所需的任何命令。
- 打开**环境变量** 部分并创建一个名为`COLOR_TEST = blue` 的变量。
- 导入`kubeconfig` 秘密。
- 输入以下命令（抱歉，有点长）：

```bash
if kubectl get virtualservice myapp-blue-green;then VSERVICE_DEFAULT=$(kubectl get virtualservice myapp-blue-green -o json | jq -r '.spec.http[-1].route[0].destination.subset');echo "Default route goes to $VSERVICE_DEFAULT";test "$VSERVICE_DEFAULT" != "$COLOR_TEST";fi
```

![sanity check block](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/blue0-1024x545.png)Sanity check block

### Adding a Deployment Block

### 添加部署块

Beyond this point, we can assume blue is not active and that it’s safe to deploy.

除此之外，我们可以假设 blue 不活跃并且可以安全部署。

- Create a second block called “Deploy”.
- Import the`kubeconfig`, `dockerhub` and `env-blue` secrets.
- Configure the same**prologue** and **environment variables** as before.
- Type the following commands in the job, which creates the gateway, routes, service, deployment:

- 创建名为“部署”的第二个块。
- 导入 `kubeconfig`、`dockerhub` 和 `env-blue` 秘密。
- 配置与以前相同的**序言**和**环境变量**。
- 在作业中键入以下命令，这将创建网关、路由、服务、部署：

```bash
checkout

# service & deployment
kubectl apply -f manifests/service.yml
envsubst < manifests/deployment.yml |tee _deployment.yml
kubectl apply -f _deployment.yml
kubectl rollout status -f _deployment.yml --timeout=120s

# routes & gateway
kubectl apply -f manifests/gateway.yml
envsubst < manifests/route-test.yml |tee _route.yml
kubectl apply -f _route.yml

# place any other setup/initialization commands, for instance...
kubectl exec -it -c myapp $(kubectl get pod -l app=myapp,color=$COLOR_TEST -o name | head -n 1) -- npm run migrate
```

![deploy block](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/blue2-1024x686.png)Deploy block

### Adding Smoke Tests

### 添加烟雾测试

- Create a third block called “Smoke Tests”.
- Repeat the **prologue** and **environment variables** from the last block.
- Import the **kubeconfig** secret. 

- 创建名为“烟雾测试”的第三个块。
- 从最后一个块重复**序言**和**环境变量**。
- 导入**kubeconfig** 秘密。

- Add your test scripts. You can run them inside a running pod with`kubectl exec`:

- 添加您的测试脚本。您可以使用 `kubectl exec` 在正在运行的 pod 中运行它们：

```bash
kubectl exec -it -c myapp $(kubectl get pod -l app=myapp,color=$COLOR_TEST -o name | head -n 1) -- npm run ping
```

- You may add more test jobs. As an example, the following command uses curl and a cookie to access blue directly:

- 您可以添加更多测试作业。例如，以下命令使用 curl 和 cookie 直接访问 blue：

```bash
export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export URL="http://${INGRESS_HOST}:${INGRESS_PORT}"
echo "Ingress gateway is $URL"
export TEST_VALUE=$(curl --cookie 'test=true' $URL/ready | jq -r '.ready')
test "$TEST_VALUE" = "true"
```

![smoke tests block](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/blue3-1024x667.png)Smoke tests block

### Activating the Blue Route

### 激活蓝色路线

The only thing left is to change the route.

剩下的就是改变路线了。

- Create a new**promotion** called “Activate Blue”

- 创建一个名为“Activate Blue”的新**促销**

![activate blue promotion](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/blue-promotion-prod-1024x267.png)Activate blue promotion

- Set the`COLOR_ACTIVE` variable to `blue`.
- Configure the**prologue** commands exactly as you did earlier.
- Import the**kubeconfig** secret.
- Type the following commands in the job:

- 将`COLOR_ACTIVE` 变量设置为`blue`。
- 完全像之前一样配置 **prologue** 命令。
- 导入**kubeconfig** 秘密。
- 在作业中键入以下命令：

```bash
checkout
envsubst < manifests/route.yml |tee _route.yml
kubectl apply -f _route.yml
```

![activate blue route block](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/blue-switch-1024x519.png)Activate blue route block

## Cleanup Pipelines

## 清理管道

At this point in the workflow, either the upgrade was a success and everything went smoothly, or we’re not satisfied and wish to go back.

在工作流的这一点上，要么升级成功且一切顺利，要么我们不满意并希望返回。

### Decommission Pipeline

### 退役管道

The decommission pipeline deletes the green deployment to free cluster resources:

退役管道删除绿色部署以释放集群资源：

- Create a new promotion: “Decommission Green”
- Add the**prologue** and kubectl **secret**.
- Set the**environment variable** `COLOR_DECOMISSION = green`.
- Type the following commands in the job:

- 创建一个新的促销活动：“退役绿色”
- 添加 **prologue** 和 kubectl **secret**。
- 设置**环境变量**`COLOR_DECOMISSION = green`。
- 在作业中键入以下命令：

```bash
if kubectl get deployment/myapp-$COLOR_DECOMISSION 2>/dev/null;then kubectl delete deployment/myapp-$COLOR_DECOMISSION;fi
```

![decomission promotion](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/dc-green-1-1024x267.png)Decomission promotion

![decomission block](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/dc-green-2-1024x505.png)Decomission block

### Rollback Pipeline

### 回滚管道

Activating the rollback pipeline is like pressing CTRL+Z; it routes all traffic back to green, effectively undoing the upgrade.

激活回滚管道就像按CTRL+Z；它将所有流量路由回绿色，有效地取消升级。

- Create a new promotion: “Rollback to Green”
- Add the**prologue** and kubectl **secret**.
- Set the**environment variable** `COLOR_ACTIVE = green`.
- Type the following commands in the job:

- 创建一个新的促销活动：“Rollback to Green”
- 添加 **prologue** 和 kubectl **secret**。
- 设置**环境变量**`COLOR_ACTIVE = green`。
- 在作业中键入以下命令：

```bash
checkout
envsubst < manifests/route.yml |tee _route.yml
kubectl apply -f _route.yml
```

![rollback promotion](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/rb-green-1.png)Rollback promotion

![rollback block](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/rb-green-2-1024x509.png)Rollback block

Congratulations! Your blue pipelines are ready.

恭喜！您的蓝色管道已准备就绪。

## Green Pipelines

## 绿色管道

We’re halfway done. Now we have to do everything again for green, but in reverse–that’s “Deploy to Green”, “Activate Green”, “Decommission Blue” and “Rollback to Blue”.

我们已经完成了一半。现在我们必须为绿色再次做所有事情，但反过来——那就是“部署到绿色”、“激活绿色”、“退役蓝色”和“回滚到蓝色”。

Go back to the first pipeline in the workflow and create a second promotion for the green branch. Keep in mind the following:

返回工作流中的第一个管道并为绿色分支创建第二个促销。请记住以下几点：

- Green pipelines are the mirror image of blue.
- The automatic promotion condition is:

- 绿色管道是蓝色的镜像。
- 自动升级条件为：

```
result = 'passed' AND tag =~ '^v.*-green$'
```

- Import`env-green` instead of `env-blue` in secrets.
- Reverse the values of`$COLOR_TEST`, `$COLOR_ACTIVE`, and `$COLOR_DECOMISSION`: replace blue with green, and green with blue.

- 在秘密中导入 `env-green` 而不是 `env-blue`。
- 反转`$COLOR_TEST`、`$COLOR_ACTIVE` 和`$COLOR_DECOMISSION` 的值：将蓝色替换为绿色，将绿色替换为蓝色。

![green deploy pipelines](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/green-staging-1024x272.png)Create the green deploy pipelines

![create the green activate & cleanup pipelines](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/green-prod-1024x222.png)Create the green activate & cleanup pipelines

When you’re done, save your work with **Run the Workflow** \> **Start**.

完成后，使用 **Run the Workflow** \> **Start** 保存您的工作。

## Deploying to Green 

## 部署到绿色

Let’s do this. Imagine we want to release `v1.0` of our awesome application. It’s green’s turn to be in production. Run these commands to start the green release:

我们开工吧。想象一下，我们想要发布我们很棒的应用程序的“v1.0”。轮到绿色投入生产了。运行这些命令以启动绿色版本：

1. Get the latest revision from GitHub.

1. 从 GitHub 获取最新版本。

```bash
$ git pull origin setup-semaphore
$ git checkout setup-semaphore
```

1. Create a release according to the naming convention.

1. 根据命名约定创建一个版本。

```bash
$ git tag -a v1.0-green -m "release v1.0 to green"
$ git commit -m "releasing v1.0"
$ git push origin v1.0-green
```

1. Semaphore picks up the commit and begins working. When the deploy pipeline stops, click on**promote** to switch traffic to green:

1. 信号量接收提交并开始工作。当部署管道停止时，点击**promote** 将流量切换为绿色：

![switch users to green](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/green-promote1-1024x211.png)Switch users to green

![green route active](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/green-complete1-1024x227.png)Green route active

1. If everything goes as planned, the happy path is to decommission blue:

1.如果一切按计划进行，快乐的道路是退役蓝色：

![green deployment complete](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/green-decomm1-1024x288.png)Green deployment complete

You can check the status of your deployment with the following commands:

您可以使用以下命令检查部署状态：

```bash
$ kubectl get deployments
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
myapp-green   1/1     1            1           72m

$ kubectl get pods
NAME                           READY   STATUS    RESTARTS   AGE
myapp-green-664d56548d-5rm24   1/1     Running   0          72m

$ kubectl get virtualservice,destinationrules
NAME                                                  GATEWAYS          HOSTS   AGE
virtualservice.networking.istio.io/myapp-blue-green   [myapp-gateway]   [*]     70m

NAME                                                   HOST    AGE
destinationrule.networking.istio.io/myapp-blue-green   myapp   70m

$ kubectl get gateway
NAME AGE
myapp-gateway   75m
```

To view the active route run: `kubectl describe virtualservice/myapp-blue-green`

查看活动路由运行：`kubectl describe virtualservice/myapp-blue-green`

![active route](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/get-vs-1024x845.png)Active route

## Handling Simultaneous Deployments

## 处理同时部署

There is a possible edge case in our process: if we make two simultaneous releases, we could find ourselves deploying a different version than the one we intended.

在我们的过程中可能存在一个边缘情况：如果我们同时发布两个版本，我们可能会发现自己部署的版本与我们预期的版本不同。

To prevent concurrency side-effects, we can set up [pipeline queues](https://docs.semaphoreci.com/essentials/pipeline-queues/). Pipeline queues let us force pipelines to run sequentially.

为了防止并发副作用，我们可以设置[管道队列](https://docs.semaphoreci.com/essentials/pipeline-queues/)。管道队列让我们强制管道按顺序运行。

You’ll have to edit the pipeline YAML directly to change this setting as the workflow editor doesn’t yet have the option. The pipeline files are located on the `.semaphore` folder, at the root of the project.

您必须直接编辑管道 YAML 才能更改此设置，因为工作流编辑器尚无该选项。管道文件位于项目根目录下的“.semaphore”文件夹中。

![blue-green deployment](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/yaml-path.png)

First, do a `git pull` to ensure you’re working on the latest revision:

首先，执行“git pull”以确保您正在处理最新版本：

```bash
$ git pull origin setup-semaphore
```

Next, add the following lines to the deployment pipeline files. You should change eight files in total: “Deploy to Blue/Green”, “Activate Blue/Green”, and the decommission and the rollback pipelines.

接下来，将以下行添加到部署管道文件中。您总共应该更改八个文件：“Deploy to Blue/Green”、“Activate Blue/Green”以及退役和回滚管道。

```yaml
queue:
name: shared kubernetes deployment
scope: organization
```

This option puts the pipelines an organization-wide queue called “shared kubernetes deployment”, ensuring that pipelines belonging to that queue always run sequentially. You can also set up conditions for placing pipelines in different queues. For further details, read about the [queue](https://docs.semaphoreci.com/reference/pipeline-yaml-reference/#queue) property.

此选项将管道置于一个名为“共享 kubernetes 部署”的组织范围的队列中，确保属于该队列的管道始终按顺序运行。您还可以设置将管道放置在不同队列中的条件。有关更多详细信息，请阅读 [queue](https://docs.semaphoreci.com/reference/pipeline-yaml-reference/#queue) 属性。

Push the changes to update the pipelines:

推送更改以更新管道：

```bash
$ git add .semaphore/*
$ git commit -m "setup shared deployment queue"
$ git push origin setup-semaphore
```

## Adding More Sanity Checks

## 添加更多健全性检查

Can you imagine what would happen if someone presses the promote button in a stale workflow by mistake? It’s difficult to say for sure, but the consequences aren’t likely to be good.

您能想象如果有人在陈旧的工作流程中错误地按下了升级按钮会发生什么吗？很难说肯定，但后果不太可能是好的。

To minimize the impact of human errors, you can add checks before every command that affects the cluster. We’ve already done some of that when we created the “Blue/Green not active” block in the deployment pipeline.

为了尽量减少人为错误的影响，您可以在影响集群的每个命令之前添加检查。当我们在部署管道中创建“Blue/Green not active”块时，我们已经完成了其中的一些工作。

The trick is to add labels to the Kubernetes resources and use `kubectl get` to validate their value before changing things. For example, you may add the Semaphore workflow id into the deployment manifest:

诀窍是为 Kubernetes 资源添加标签，并在更改内容之前使用 `kubectl get` 来验证它们的值。例如，您可以将信号量工作流 ID 添加到部署清单中：

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
name: myapp-$COLOR_TEST
labels:
    app: myapp
    color: $COLOR_TEST
    workflow: $SEMAPHORE_WORKFLOW_ID

# ... rest of the manifest ...
```

And test that the workflow label on the cluster is valid before changing the route:

并在更改路由之前测试集群上的工作流标签是否有效：

```bash
test "$(kubectl get deployments -l app=myapp,color=$COLOR_DEPLOY,workflow=$SEMAPHORE_WORKFLOW_ID -o=jsonpath={.items..metadata.name})" = "myapp-${COLOR_DEPLOY}"
```

Because any command that exits non-zero status stops the pipeline, this effectively prevents anyone from activating an invalid route. You may also add failsafe checks to verify that the pods are on the correct version and to validate the deployment. The more sanity tests you add, the more robust the process becomes.

因为任何退出非零状态的命令都会停止管道，这有效地防止了任何人激活无效路由。您还可以添加故障安全检查以验证 pod 是否在正确的版本上并验证部署。您添加的健全性测试越多，流程就越稳健。

## Trying the Blue Pipeline

## 尝试蓝色管道

Let’s be thorough and try blue by simulating a second release.

让我们通过模拟第二个版本来彻底尝试蓝色。

1. Start the process by tagging as blue.

1. 通过标记为蓝色开始该过程。

```bash
$ git tag -a v2.0-blue -m "release v2.0 to blue"
$ git commit -m "releasing v2.0"
$ git push origin v2.0-blue
```

1. This time, the blue pipelines are activated.
2. Wait for it to stop, press**promote**.

1. 这一次，蓝色管道被激活。
2. 等待它停止，按**promote**。

![blue](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/blue-promote1-1024x165.png)Switch users to blue

1. You can now remove green or try a rollback.

1.您现在可以删除绿色或尝试回滚。

![blue route is active](https://wpblog.semaphoreci.com/wp-content/uploads/2020/09/blue-complete1-1024x223.png)Blue route is now active

![+1](https://github.githubassets.com/images/icons/emoji/unicode/1f44d.png) You just finished your first blue-green cycle, way to go!![smiley](https://github.githubassets.com/images/icons/emoji/unicode/1f603.png)

## Conclusion

##  结论

Upgrading is always a risky business. No matter how much testing we do, there’s still a chance for something to go wrong. But with a few carefully placed tests and a robust CI/CD workflow, we can avoid a lot of headaches.

升级始终是一项有风险的业务。无论我们进行多少测试，仍有可能出现问题。但是通过一些精心放置的测试和强大的 CI/CD 工作流程，我们可以避免很多麻烦。

With a few modifications, you can adapt these pipelines to any application and cloud. By all means, play with them, swap parts as required, or experiment with different routing strategies. You can even use a setup like this to do [canary releases](https://semaphoreci.com/blog/what-is-canary-deployment).

通过一些修改，您可以使这些管道适应任何应用程序和云。无论如何，与他们一起玩，根据需要交换零件，或尝试不同的路由策略。您甚至可以使用这样的设置来执行 [canary 版本](https://semaphoreci.com/blog/what-is-canary-deployment)。

We have a lot of great resources to help you with your Docker and Kubernetes learning:

我们有很多很棒的资源可以帮助您学习 Docker 和 Kubernetes：

- [What is Blue-Green Deployment?](https://semaphoreci.com/blog/blue-green-deployment)
- [CI/CD with Docker and Kubernetes](https://semaphoreci.com/resources/cicd-docker-kubernetes): Learn Docker, Kubernetes, and CI/CD principles with this free eBook.
- [How to Release Faster with Continuous Delivery for Google Kubernetes](https://semaphoreci.com/blog/continuous-delivery-google-kubernetes): Shows deployment for Google Kubernetes Engine.
- [CI/CD for Microservices on DigitalOcean Kubernetes](https://semaphoreci.com/blog/cicd-microservices-digitalocean-kubernetes): Deploy to DigitalOcean Kubernetes.
- [Continuous Integration and Delivery to AWS Kubernetes](https://semaphoreci.com/blog/continuous-integration-delivery-aws-eks-kubernetes): Learn how to use AWS Elastic Kubernetes Service.

- [什么是蓝绿部署？](https://semaphoreci.com/blog/blue-green-deployment)
- [使用 Docker 和 Kubernetes 进行 CI/CD](https://semaphoreci.com/resources/cicd-docker-kubernetes)：通过这本免费电子书了解 Docker、Kubernetes 和 CI/CD 原则。
- [如何通过 Google Kubernetes 的持续交付更快地发布](https://semaphoreci.com/blog/continuous-delivery-google-kubernetes)：展示了 Google Kubernetes Engine 的部署。
- [DigitalOcean Kubernetes 上微服务的 CI/CD](https://semaphoreci.com/blog/cicd-microservices-digitalocean-kubernetes)：部署到 DigitalOcean Kubernetes。
- [持续集成和交付到 AWS Kubernetes](https://semaphoreci.com/blog/continuous-integration-delivery-aws-eks-kubernetes)：了解如何使用 AWS Elastic Kubernetes 服务。

Thanks for reading!

谢谢阅读！

You can also download this article as a PDF and read it on the couch.

您还可以将本文下载为 PDF 文件并在沙发上阅读。

[Download](https://semaphoreci.com/resources/download-blue-green-kubernetes) 

[下载](https://semaphoreci.com/resources/download-blue-green-kubernetes)

