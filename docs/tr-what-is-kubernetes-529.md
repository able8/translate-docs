# What is Kubernetes?

# 什么是 Kubernetes？

If you know what is Kubernetes you may be interested in our previous blog “ [How Kubernetes Works](https://goglides.io/kubernetes-how-it-works/94/)“.

如果您知道什么是 Kubernetes，您可能会对我们之前的博客“[Kubernetes 的工作原理](https://goglides.io/kubernetes-how-it-works/94/)”感兴趣。

Designed and developed by Google, Kubernetes is an open-sourced container-orchestration system that helps in the automation of application deployment, scaling containerized applications, and monitoring such applications. Now maintained by the Cloud Native Computing Foundation, Kubernetes seeks to provide a platform for the automated deployment, scaling, and the operations of application containers across a cluster of hosts.

Kubernetes 由 Google 设计和开发，是一个开源容器编排系统，有助于实现应用程序部署的自动化、扩展容器化应用程序和监控此类应用程序。现在由云原生计算基金会维护，Kubernetes 旨在为跨主机集群的应用程序容器的自动化部署、扩展和操作提供一个平台。

What Kubernetes does is that it provides you with a framework that allows you to run distributed systems smoothly and effectively. It looks after the scaling, looks after any application failovers, and immediately provides a deployment pattern.

Kubernetes 所做的是它为您提供了一个框架，让您可以平稳有效地运行分布式系统。它负责扩展，负责任何应用程序故障转移，并立即提供部署模式。

The development of Kubernetes was influenced by Google’s work on its Borg System. And following the release of Kubernetes v1.0, Google partnered with the Linux Foundation and formed the Cloud Native Computing Foundation.

Kubernetes 的发展受到 Google 在其 Borg 系统上的工作的影响。随着 Kubernetes v1.0 的发布，谷歌与 Linux 基金会合作，成立了云原生计算基金会。

## Basic Kubernetes Terminologies

## Kubernetes 基本术语

Before we learn what Kubernetes is and how it operates, let us first understand a few application program interfaces that are a part of Kubernetes.

在我们了解 Kubernetes 是什么以及它是如何运行之前，让我们先了解一些作为 Kubernetes 一部分的应用程序接口。

1. **Node:** A node like a virtual machine is a machine entity or physical hardware. It acts as a worker instance within the Kubernetes Cluster. The nodes are responsible for performing the assigned tasks. Nodes are the physical infrastructure that allows your application to run on the Virtual Machine’s server.
2. **Pods:** Pods are a group of one or more containers that share network and storage with a Kubernetes configuration, which allows you to move containers within the cluster more efficiently. Each particular pod is assigned with an IP address, a hostname, IPC, which enable it to be accessed by other pods within a cluster.
3. **Service:** Service helps in separating the work definitions from the pods. Services are a sort of static pointers to the pods—services help in assuring that each pod receives an IP address.
4. **Replica Set:** The primary function of a replica set is to ensure that a given number of pods are running at a particular time. Replica Set is also a vital aspect of the Kubernetes autoscaling functionality. If a pod crashes or dies, the Replica Set makes sure a new one is created. In case there are too many pods, it cut down on them, and if there are few, then it creates new pods. Replica sets also make it much easier to target a particular set of pods.
5. **Kubelet:** A Kubelet is present inside each of the Nodes. Each node has a Kubelet. The primary responsibility of the Kubelet is to manage the nodes. It reads the container manifests and ensures that the particular containers are up and running.
6. **Kubectl:** A Kubectl is simply a command-line configuration interface for Kubernetes. It consists of an extensive list of available commands that can be used to manage Kubernetes.

1. **节点：** 像虚拟机这样的节点是机器实体或物理硬件。它充当 Kubernetes 集群中的工作实例。节点负责执行分配的任务。节点是允许您的应用程序在虚拟机服务器上运行的物理基础设施。
2. **Pods：** Pods 是一组一个或多个容器，它们通过 Kubernetes 配置共享网络和存储，这使您可以更有效地在集群内移动容器。每个特定的 pod 都分配有一个 IP 地址、一个主机名、IPC，使其能够被集群中的其他 pod 访问。
3. **Service:** Service 有助于将工作定义与 Pod 分开。服务是一种指向 Pod 的静态指针——服务有助于确保每个 Pod 都收到一个 IP 地址。
4. **副本集：**副本集的主要功能是确保给定数量的 Pod 在特定时间运行。副本集也是 Kubernetes 自动缩放功能的一个重要方面。如果 Pod 崩溃或死亡，副本集会确保创建一个新的。如果 Pod 太多，它会减少它们，如果很少，它会创建新的 Pod。副本集还可以更轻松地定位特定的 pod 集。
5. **Kubelet：** 每个节点内都有一个 Kubelet。每个节点都有一个 Kubelet。 Kubelet 的主要职责是管理节点。它读取容器清单并确保特定容器已启动并正在运行。
6. **Kubectl:** Kubectl 只是 Kubernetes 的命令行配置界面。它包含可用于管理 Kubernetes 的大量可用命令列表。

## How Does Kubernetes Work?

## Kubernetes 是如何工作的？

When you first deploy Kubernetes, you get a Cluster. A Cluster consists of two parts, i.e., The Control Plane and Compute Machines. The control plane includes Master Nodes, and the Compute Machines consist of Worker Nodes. The worker nodes are responsible for running the Pods, which are made up of Containers. The Control Plane is responsible for managing the worker nodes as well as the pods within the cluster.

首次部署 Kubernetes 时，您会获得一个集群。集群由两部分组成，即控制平面和计算机。控制平面包括主节点，计算机由工作节点组成。工作节点负责运行由容器组成的 Pod。控制平面负责管理集群中的工作节点和 pod。

![kubernetes-diagram](https://i0.wp.com/goglides.io/wp-content/uploads/2020/06/01-kubernetes-diagram-1.png?resize=1024%2C636&ssl=1)kubernetes-diagram

-图表

[Source](https://www.redhat.com/cms/managed-files/kubernetes-diagram-2-824x437.png) 

[来源](https://www.redhat.com/cms/managed-files/kubernetes-diagram-2-824x437.png)

You can see the image above, Kubernetes operates on top of the operating system and then interacts with the pods of the container that are running on the nodes. The operator relays the command to the master node, which then distributes it to the secondary nodes. The node that is best suited for the particular task will be automatically decided, and the resources required for the task will be allocated to the particular pod required to fulfill the assigned work.

您可以看到上图，Kubernetes 在操作系统之上运行，然后与运行在节点上的容器的 Pod 进行交互。操作员将命令中继到主节点，然后将其分发到辅助节点。最适合特定任务的节点将被自动决定，任务所需的资源将分配给完成指定工作所需的特定 pod。

## What Services Does Kubernetes Provide You?

## Kubernetes 为您提供哪些服务？

Some of the most basic features that Kubernetes provides you are:

Kubernetes 为您提供的一些最基本的功能是：

1. Maintain and adapt containers across a wide range of hosts.
2. Self-Heal applications with Auto-Placement, Restart, Scaling, and Replication.
3. Makes sure the applications run as you intended it to run.
4. It helps you control the updates and application deployment.
5. Mount and add storage systems of your choice
6. Efficient use of hardware for maximum utilization of resources.



1. 在各种主机上维护和调整容器。
2. 具有自动放置、重启、扩展和复制功能的自我修复应用程序。
3. 确保应用程序按照您的预期运行。
4. 它可以帮助您控制更新和应用程序部署。
5. 安装和添加您选择的存储系统
6. 有效利用硬件，最大限度地利用资源。

## Kubernetes Components & Architecture

## Kubernetes 组件和架构

Before studying the Kubernetes Components, let’s take a brief look at the principles and design that define Kubernetes.

在学习 Kubernetes Components 之前，让我们先简单了解一下定义 Kubernetes 的原则和设计。

The Kubernetes Cluster design is based on three simple principles, i.e.

Kubernetes 集群设计基于三个简单的原则，即

1. Easy to use: The Cluster should be operable using minimum commands.

2. Extendable: It should be customizable.

3. Secure: Probably the most important, The Kubernetes Cluster should always follow the latest up-to-date security practices and principles.


1. 易于使用：集群应该可以使用最少的命令进行操作。
2. 可扩展：它应该是可定制的。
3. 安全：可能是最重要的，Kubernetes 集群应该始终遵循最新的安全实践和原则。

Kubernetes consists of many components that all communicate with each other through the API server. As we talked about before, the Kubernetes cluster consists of The Control Plane, which holds the Master Node, and the Compute Machine, which holds the Worker Node.

Kubernetes 由许多组件组成，这些组件都通过 API 服务器相互通信。正如我们之前谈到的，Kubernetes 集群由控制平面组成，控制平面包含主节点，计算机器包含工作节点。

![Components of Kubernetes](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/02-components-of-kubernetes.png?resize=1024%2C605&ssl=1)Componentsof Kubernetes

的 Kubernetes

[Source:](https://d33wubrfki0l68.cloudfront.net/7016517375d10c702489167e704dcb99e570df85/7bb53/images/docs/components-of-kubernetes.png)

[来源：](https://d33wubrfki0l68.cloudfront.net/7016517375d10c702489167e704dcb99e570df85/7bb53/images/docs/components-of-kubernetes.png)

The above image shows the components inside the Kubernetes Control Plane and the components inside of the Compute Machine.

上图显示了 Kubernetes Control Plane 内部的组件和 Compute Machine 内部的组件。

Now let’s take a closer look at the components.

现在让我们仔细看看这些组件。

## Control Plane Components

## 控制平面组件

The Control Plane is the central part of the Kubernetes cluster. Inside the Control Plane is the Master Node, which holds all the components that control the Kubernetes Cluster. It makes the entire decision about the cluster. It also detects as well as response to any cluster events.

控制平面是 Kubernetes 集群的核心部分。控制平面内是主节点，它包含控制 Kubernetes 集群的所有组件。它做出关于集群的整个决定。它还检测并响应任何集群事件。

1. **kube-apiserver:** This is the front end of the Kubernetes Control Plane. It handles the internal as well as external requests. The kube-apiserver has been specifically designed to scale horizontally, which means that it is scaling up by deploying more instances. It also helps to communicate with all the other components inside the cluster.
2. **etcd:** The primary function of ‘etcd’ is to store the configuration information. The nodes can then use the information within the cluster. As the etcd contains sensitive information, it can only be accessed by the Kubernetes API server. It is a distributed key-value store responsible for implementing locks within the cluster.
3. **kube-scheduler:** A vital component of the Kubernetes, a kube-scheduler, is responsible for the distribution of the container and the workloads across multiple nodes. In simple terms, it allocates newly formed pods to the available nodes. The kube-scheduler is responsible for finding newly formed containers and assigning them to the available nodes.
4. **kube-controller-manager:** This is a component of the Master Node that is responsible for running the controller. The kube-controller-manager contains several controller functions compiled into a single binary to save time and reduce complexity.



1. **kube-apiserver:** 这是 Kubernetes 控制平面的前端。它处理内部和外部请求。 kube-apiserver 专门设计用于水平扩展，这意味着它通过部署更多实例来扩展。它还有助于与集群内的所有其他组件进行通信。
2. **etcd:** ‘etcd’的主要功能是存储配置信息。然后节点可以使用集群内的信息。由于 etcd 包含敏感信息，因此只能由 Kubernetes API 服务器访问。它是一个分布式键值存储，负责在集群内实现锁。
3. **kube-scheduler：** Kubernetes 的一个重要组件，kube-scheduler，负责跨多个节点分发容器和工作负载。简单来说，它将新形成的 Pod 分配给可用节点。 kube-scheduler 负责寻找新形成的容器并将它们分配给可用节点。
4. **kube-controller-manager:** 这是主节点的一个组件，负责运行控制器。 kube-controller-manager 包含多个控制器函数，它们被编译成一个二进制文件，以节省时间并降低复杂性。

The controller functions include:

控制器功能包括：

- **Node controller:** It notices and responds whenever a particular node goes down.
- **Replication controller:** It is responsible for maintaining the correct number of pods.
- **Endpoints controller:** It is responsible for populating the endpoint objects.
- **Service accounts & token controller:** It is responsible for creating default accounts and API access tokens for the new namespaces. 

- **节点控制器：** 它会在特定节点出现故障时发出通知并做出响应。
- **复制控制器：** 它负责维护正确数量的 Pod。
- **端点控制器：** 它负责填充端点对象。
- **服务帐户和令牌控制器：** 它负责为新命名空间创建默认帐户和 API 访问令牌。

1. **Cloud-controller-manager:** The cloud-controller-manager allows you to interact and link your cluster to your cloud provider’s API. It means that it allows you to talk to your cloud providers. The cloud-controller-manager also combines several different control functions into one.

1. **Cloud-controller-manager：** cloud-controller-manager 允许您交互并将集群链接到云提供商的 API。这意味着它允许您与您的云提供商交谈。云控制器管理器还将几种不同的控制功能合二为一。

The controller functions include:

控制器功能包括：

- **Node Controller:** It is responsible for checking whether or not the node has been deleted from the cloud provider after it stops responding.
- **Route Controller:** It is responsible for setting up routes in the underlying cloud infrastructure.
- **Service Controller:** It is responsible for creating, updating, and deleting the cloud provider load balancers.

- **节点控制器：** 负责检查节点停止响应后是否已从云提供商中删除。
- **Route Controller:** 负责在底层云基础架构中设置路由。
- **服务控制器：** 负责创建、更新和删除云提供商负载均衡器。

## Worker Node Components

## 工作节点组件

The Master Node controls the worker node. The node components will run on every node while maintaining running pods as well as providing the Kubernetes runtime environment.

主节点控制工作节点。节点组件将在每个节点上运行，同时维护正在运行的 Pod 并提供 Kubernetes 运行时环境。

1. **kubelet:** We talked briefly about what kubelet is. A kubelet is present inside of the Node. Each node has a kubelet. The primary responsibility of the Kubelet is to manage the nodes. It reads the container manifests and ensures that the particular containers are up and running.
2. **kube-proxy:** A kube-proxy is a network proxy service that runs on each of the nodes within a cluster. It also helps the external host by making services available. It also assists in forwarding all the requests to the correct containers and is also capable of adequately performing first load balancing. A kube-proxy managed the pods within a node creates new containers’ health checkups, etc.
3. **container runtime:** A container runtime is responsible for running the containers. Kubernetes supports several different container runtimes, such as Docker, CRI-O, containers as well as any and all implementations of the Kubernetes Container Runtime Interface.



1. **kubelet：** 我们简要介绍了 kubelet 是什么。一个 kubelet 存在于节点内部。每个节点都有一个 kubelet。 Kubelet 的主要职责是管理节点。它读取容器清单并确保特定容器已启动并正在运行。
2. **kube-proxy：** kube-proxy 是一种网络代理服务，运行在集群内的每个节点上。它还通过提供服务来帮助外部主机。它还有助于将所有请求转发到正确的容器，并且还能够充分执行首次负载平衡。 kube-proxy 管理节点内的 pod，创建新容器的健康检查等。
3. **容器运行时：** 容器运行时负责运行容器。 Kubernetes 支持多种不同的容器运行时，例如 Docker、CRI-O、容器以及 Kubernetes 容器运行时接口的任何和所有实现。

## Kubernetes & Docker



Docker is a tool that is used for building, distributing as well as running Docker containers. Docker provides the user with its clustering tools that can be used to arrange and schedule the containers on the clusters. Likewise, Kubernetes is a more extensive container orchestration system for Docker. It is meant for the effective coordination of the cluster of nodes efficiently. Although Kubernetes and Docker are different types of technologies, they coordinate well with one another and help in the management and the deployment of containers in a distributed architecture. When using Kubernetes with Docker, the automated system requests the Docker to do things such as launch a specified container, start and stop the containers, etc., in cases when it would have been done manually by an admin for all the containers.

Docker 是一种用于构建、分发和运行 Docker 容器的工具。 Docker 为用户提供了其集群工具，可用于在集群上安排和调度容器。同样，Kubernetes 是更广泛的 Docker 容器编排系统。它旨在有效地有效协调节点集群。虽然Kubernetes和Docker是不同类型的技术，但它们相互协调，有助于分布式架构中容器的管理和部署。当将 Kubernetes 与 Docker 一起使用时，自动化系统会请求 Docker 执行诸如启动指定容器、启动和停止容器等操作，以防管理员为所有容器手动完成。

## Why Do You Need Kubernetes?

## 为什么需要 Kubernetes？

To meet the demands of the ever-changing business world, you and your team must be able to adapt to the situation. You will be required to build new applications and services at a rapid pace.

为了满足不断变化的商业世界的需求，您和您的团队必须能够适应这种情况。您将需要快速构建新的应用程序和服务。

An application requires multiple containers. The use of Kubernetes will allow you to build applications and services that will span across several different containers. It will also allow you to schedule those containers across clusters while also scaling those containers and managing the health of the containers over time. It will provide you with information and metrics regarding your containers and clusters.

一个应用程序需要多个容器。 Kubernetes 的使用将允许您构建跨多个不同容器的应用程序和服务。它还允许您跨集群调度这些容器，同时扩展这些容器并随着时间的推移管理容器的健康状况。它将为您提供有关容器和集群的信息和指标。

The use of Kubernetes will assist in conserving resources more effectively and efficiently. It monitors all the clusters and makes decisions on where to effectively launch the containers with the available resources currently being utilized on the nodes. And if an application goes down, Kubernetes recovers it automatically.

Kubernetes 的使用将有助于更有效地节约资源。它监控所有集群，并决定在何处有效启动容器，并使用当前在节点上使用的可用资源。如果应用程序出现故障，Kubernetes 会自动恢复它。

## Installing & Getting Started with Kubernetes 

## 安装和开始使用 Kubernetes

The first thing that is required is for you to install Kubernetes on your hardware or one of the major cloud providers. Installing Kubernetes is quite tricky because there are many components involved. But there are plenty of tools, both open-source and paid solutions, in the market place, which makes the installation process more comfortable. There are many ways how you can install a Kubernetes cluster. For this tutorial, we will be using Minikube. For that, you will first need to start and run two things, i.e., Kubectl and Minikube.

您需要做的第一件事是在您的硬件或主要云提供商之一上安装 Kubernetes。安装 Kubernetes 非常棘手，因为涉及到许多组件。但是市场上有很多工具，包括开源和付费解决方案，这使得安装过程更加舒适。您可以通过多种方式安装 Kubernetes 集群。在本教程中，我们将使用 Minikube。为此，您首先需要启动并运行两个东西，即 Kubectl 和 Minikube。

Kubectl is a command-line interface (CLI) tool that will allow you to interact with the cluster.

Kubectl 是一个命令行界面 (CLI) 工具，可让您与集群进行交互。

Minikube is a binary that will deploy the cluster locally on your development machine.

Minikube 是一个二进制文件，它将在您的开发机器上本地部署集群。

With these tools, you can now start arranging your containerized applications to the cluster in a few short minutes. After you have installed Minikube, you can now run a single node cluster within your local machine.

使用这些工具，您现在可以在几分钟内开始将容器化应用程序安排到集群中。安装 Minikube 后，您现在可以在本地机器中运行单节点集群。

To start the Minikube cluster, you can use the code.

要启动 Minikube 集群，您可以使用代码。

```
minikube start
```


![minikube start command](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/minikube-start.png?resize=1024%2C264&ssl=1)minikube start

Likewise, if you want to interact with the Kubernetes Cluster, you will need to install Kubectl CLI. Once you installed it, you can use the following codes to interact with the Kubernetes cluster.

同样，如果您想与 Kubernetes 集群交互，则需要安装 Kubectl CLI。安装后，您可以使用以下代码与 Kubernetes 集群进行交互。

```
kubectl config get-contexts
```


![kubectl-config-get-contexts](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-config-get-contexts.png?resize=1024%2C136&ssl=1)

```
kubectl config set-contexts <context-name>
```


![kubectl-config-set-context-minikube](https://i0.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-config-set-context-minikube.png?resize=1024%2C46&ssl=1)



```
kubectl config current-context
```


![kubectl-config-current-context](https://i0.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-config-current-context.png?resize=982%2C76&ssl=1)



```
Kubectl config delete-context <context-name>
```


By now, you will have a local single Node Kubernetes cluster up and running on your machine. Now, if you are looking to deploy your first containerized application, then you run the following codes to Minikube. In this demonstration, we will be deploying a simple Hello World with an exposed endpoint on the Minikube IP address.

到现在为止，您将在您的机器上启动并运行一个本地单节点 Kubernetes 集群。现在，如果您想部署您的第一个容器化应用程序，那么您可以将以下代码运行到 Minikube。在本演示中，我们将在 Minikube IP 地址上部署一个带有公开端点的简单 Hello World。

> Start with,

> 开始，

```
kubectl create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.10
```


![kubectl-create-deployment-hello-minikube](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-create-deployment-hello-minikube.png?resize=1024%2C56&ssl=1)



You will see that your deployment was successful. To view the deployment, you can use :

您将看到您的部署成功。要查看部署，您可以使用：

```
kubectl get deployments
```


![kubectl-get-deployment](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-get-deployment.png?resize=882%2C96&ssl=1)

After the deployment process, a Kubernetes Pod will have been created. To view the pods, you can use:

在部署过程之后，将创建一个 Kubernetes Pod。要查看 pod，您可以使用：

```
kubectl get pods
```


![kubectl-get-pods](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-get-pods.png?resize=996%2C92&ssl=1)

You will then need to expose the pod as a Kubernetes service before being able to hit the Hello World with an HTTP request from outside your cluster. To do so, you can use:

然后，您需要将 pod 公开为 Kubernetes 服务，然后才能使用来自集群外部的 HTTP 请求访问 Hello World。为此，您可以使用：

```
kubectl expose deployment hello-minikube --type=NodePort --port=8080
```


![kubectl-expose-deployment-hello-minikube](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-expose-deployment-hello-minikube.png?resize=1024%2C65&ssl=1)

The exposure will create a Kubernetes service. And to view the Service, you can use:

曝光将创建一个 Kubernetes 服务。要查看服务，您可以使用：

```
kubectl get services
```


![kubectl-get-services](https://i0.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-get-services.png?resize=1024%2C111&ssl=1)

If you want to find out the URL that was used to access your containerized application, you can use:

如果您想找出用于访问容器化应用程序的 URL，您可以使用：

```
minikube service hello-minikube --url
```


![minikube-serviec-hello-minikube-url](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/minikube-serviec-hello-minikube-url.png?resize=1024%2C104&ssl=1)



To test whether or not your exposed Service has reached the pod, you can curl the response from your terminal. To do that, you can use:

要测试您暴露的 Service 是否已到达 pod，您可以从终端 curl 响应。为此，您可以使用：

```
curl http://<minikube-ip>:<port>
```


![curl-url-command](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/curl-url-command.png?resize=1004%2C802&ssl=1)



The HTTP request has made via your Kubernetes Service. And once you check your logs, you will see the following,

HTTP 请求是通过您的 Kubernetes 服务发出的。一旦你检查你的日志，你会看到以下内容，

![kubectl-logs-pods](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/kubectl-logs-pods.png?resize=1024%2C268&ssl=1)

Now once you have followed the above steps, you should have a functioning Kubernetes pod and deployment that is running a simple Hello World application.

现在，一旦您按照上述步骤进行操作，您应该拥有一个运行简单的 Hello World 应用程序的 Kubernetes pod 和部署。

Likewise, if you are looking to start using Kubernetes, you can start by trying to build a Kubernetes Cluster. You can start by taking a look at different Managed Kubernetes offering provided by some top cloud providers. 

同样，如果您打算开始使用 Kubernetes，您可以从尝试构建 Kubernetes 集群开始。您可以首先查看一些顶级云提供商提供的不同托管 Kubernetes 产品。

