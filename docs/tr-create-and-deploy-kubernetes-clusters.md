# How to Create a Kubernetes Cluster Locally - Simple Tutorial

# 如何在本地创建 Kubernetes 集群 - 简单教程

## How to create a Kubernetes cluster locally and deploy simple front-end apps that communicate with Kubernetes

## 如何在本地创建Kubernetes集群并部署与Kubernetes通信的简单前端应用

June 17, 2020

2020 年 6 月 17 日

_As a software engineer at Capital One, I get to explore cutting edge technologies every day in my work. I have worked with Docker and Docker Swarm and I always wanted to learn Kubernetes. However, I kept postponing it. Finally, I was able to dive into it and thought, “Why not create an application with Kubernetes and write about it while my mind is fresh!” That way others - and not just myself - can benefit from what I learned._

_作为 Capital One 的一名软件工程师，我每天都在工作中探索前沿技术。我曾与 Docker 和 Docker Swarm 合作过，我一直想学习 Kubernetes。然而，我一直在推迟它。最后，我能够深入研究它并想，“为什么不使用 Kubernetes 创建一个应用程序并在我头脑清醒的时候写下它！”这样其他人——而不仅仅是我自己——可以从我学到的东西中受益。_

## Introduction

##  介绍

Today we are going to create a Kubernetes cluster and deploy a simple React JS app which generates a random number by calling an Express JS app. We are going to orchestrate the whole process by using Kubernetes. We will first dockerize our front end/back end apps. Then using Kubernetes we will deploy the pods (React front end app/Express back end app) and access them via Kubernetes services.

今天我们将创建一个 Kubernetes 集群并部署一个简单的 React JS 应用程序，它通过调用一个 Express JS 应用程序来生成一个随机数。我们将使用 Kubernetes 来编排整个过程。我们将首先dockerize我们的前端/后端应用程序。然后使用 Kubernetes，我们将部署 pod（React 前端应用程序/Express 后端应用程序）并通过 Kubernetes 服务访问它们。

## What is a Kubernetes Cluster?

## 什么是 Kubernetes 集群？

Like any cluster you provide a set of nodes to Kubernetes. You tell kubernetes how to deploy containers in the cluster. How much memory or processing units each container gets and how they interact with each other.

像任何集群一样，您向 Kubernetes 提供一组节点。你告诉 kubernetes 如何在集群中部署容器。每个容器获得多少内存或处理单元以及它们如何相互交互。

## Kubernetes Cluster Diagram

## Kubernetes 集群图

The below diagram depicts what we are going to achieve today, where we will create a Kubernetes cluster and services/deployments for our front and back ends.

下图描述了我们今天要实现的目标，我们将为前端和后端创建一个 Kubernetes 集群和服务/部署。

## Prerequisites

## 先决条件

1. [Docker for Desktop](https://www.docker.com/products/docker-desktop) (latest version)
2. [Kubernetes](https://kubernetes.io/docs/setup/learning-environment/minikube/) or [Docker Kubernetes](https://www.docker.com/products/kubernetes)
3. Node & NPM (Only if you want to run the applications stand alone)
4. YAML

1. [桌面Docker](https://www.docker.com/products/docker-desktop)（最新版本)
2. [Kubernetes](https://kubernetes.io/docs/setup/learning-environment/minikube/) 或 [Docker Kubernetes](https://www.docker.com/products/kubernetes)
3. Node & NPM（仅当您想独立运行应用程序时）
4. YAML

**_NOTE:_** _This tutorial requires basic/working knowledge on Docker, Node, & NPM._

**_注意：_** _本教程需要有关 Docker、Node 和 NPM 的基本/工作知识。_

Let’s dive in without any further ado.

让我们毫不费力地深入研究。

## Getting Started With Kubernetes

## Kubernetes 入门

#### **Installing Kubernetes**

#### **安装Kubernetes**

_If you have **Docker Desktop**, go to **preferences**, go to the **Kubernetes tab,** and click **Enable Kubernetes**_ _. It may take a while to spin up Kubernetes on to your machine, so go make a coffee while it does its magic. ☕_

_如果您有 **Docker Desktop**，请转到 **preferences**，转到 **Kubernetes 选项卡，** 并单击 **Enable Kubernetes**_ _。将 Kubernetes 启动到您的机器上可能需要一段时间，所以在它发挥其魔力的时候去冲杯咖啡吧。 ☕_

To verify if Kubernetes is running, type the below two commands:

要验证 Kubernetes 是否正在运行，请键入以下两个命令：

```
kubectl version

Outputs:

Client Version: version.Info{Major:"1", Minor:"15", GitVersion:"v1.15.0", GitCommit:"e8462b5b5dc2584fdcd18e6bcfe9f1e4d970a529", GitTreeState:"clean", BuildDate:"2019-06-19T16:40:16Z", GoVersion:"go1.12.5", Compiler:"gc", Platform:"darwin/amd64"}

Server Version: version.Info{Major:"1", Minor:"14", GitVersion:"v1.14.8", GitCommit:"211047e9a1922595eaa3a1127ed365e9299a6c23", GitTreeState:"clean", BuildDate:"2019-10-15T12:02:12Z", GoVersion:"go1.12.10", Compiler:"gc", Platform:"linux/amd64"}

kubectl cluster-info

Outputs:
Kubernetes master is running at https://kubernetes.docker.internal:6443

KubeDNS is running at https://kubernetes.docker.internal:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
```

## Bundling Application As Docker Images

## 将应用程序捆绑为 Docker 镜像

#### The full code for this can be found on [GitHub](https://github.com/chiku11/react-k8)

#### 完整代码可以在 [GitHub](https://github.com/chiku11/react-k8) 上找到

Please download the project. The project has two sub folders in it.

请下载项目。该项目中有两个子文件夹。

- Client -> React based app
- Server -> Express based app

- 客户端 -> 基于 React 的应用
- 服务器 -> 基于 Express 的应用程序

Please follow the below steps to setup the project and start the Kubernetes cluster:

请按照以下步骤设置项目并启动 Kubernetes 集群：

#### **Client**

####  **客户**

- _cd client_
- _npm install_
- _npm run build_
- **_docker build -t frontend:1.0 ._**
- **_kubectl apply -f frontend.deploy.yml_**
- **_kubectl apply -f frontend.service.yml_**

- _cd 客户端_
- _npm 安装_
- _npm 运行构建_
- **_docker build -t 前端：1.0 ._**
- **_kubectl apply -f frontend.deploy.yml_**
- **_kubectl apply -f frontend.service.yml_**

#### **Backend**

#### **后端**

- _cd server_
- _npm install_
- **_docker build -t backend:1.0 ._**
- **_kubectl apply -f backend.deploy.yml_**
- **_kubectl apply -f backend.service.yml_**

- _cd 服务器_
- _npm 安装_
- **_docker build -t 后端：1.0 ._**
- **_kubectl apply -f backend.deploy.yml_**
- **_kubectl apply -f backend.service.yml_**

Go to the browser, type **localhost**and hit enter, you should see the application loaded.

转到浏览器，输入 **localhost** 并按 Enter，您应该会看到应用程序已加载。

## Bringing It Down Further

## 进一步降低

#### **Kubernetes Pods**

#### **Kubernetes Pods**

Pods are the smallest deployable units of computing that can be created and managed in Kubernetes.

Pod 是可以在 Kubernetes 中创建和管理的最小的可部署计算单元。

```
kubectl run nginx-frontend --image=frontend:1.0
```

The above creates a pod which hosts the front end container. You can’t access it yet since a host port isn’t exposed for the container. We will uncover those later.

以上创建了一个托管前端容器的 pod。您还不能访问它，因为没有为容器公开主机端口。我们稍后会揭开这些。

A pod can host multiple containers as well. Instead of creating via command line let’s create via a YAML file.

一个 Pod 也可以托管多个容器。我们不是通过命令行创建，而是通过 YAML 文件创建。

```
# To create a pod with multiple containers
kubectl apply -f app.pod.yml

# To see container status
kubectl get pod/mymulticontainerapp
Outputs:
NAME                  READY   STATUS    RESTARTS   AGE
mymulticontainerapp   2/2     Running   0          9m

# To see more details about the container
kubectl describe pod/mymulticontainerapp
```

To access the container from the host machine you have to do a port forwarding, i.e attaching host port to container port. You can do that via the below command:

要从主机访问容器，您必须进行端口转发，即将主机端口连接到容器端口。您可以通过以下命令执行此操作：

```
kubectl port-forward pod/mymulticontainerapp 9999:80 3000:3000

Outputs:
Forwarding from 127.0.0.1:9999 -> 80
Forwarding from [::1]:9999 -> 80
Forwarding from 127.0.0.1:3000 -> 3000
Forwarding from [::1]:3000 -> 3000
```

The above command exposes the host port 9999 to container port 80 and host port 3000 to container port 3000 which is what our front end/back end containers are listening on. Any request to host 9999 will be forwarded to container on port 80. Any request to the host 3000 will be forwarded to container on port 3000.

上面的命令将主机端口 9999 暴露给容器端口 80，将主机端口 3000 暴露给容器端口 3000，这是我们的前端/后端容器正在监听的。对主机 9999 的任何请求都将转发到端口 80 上的容器。对主机 3000 的任何请求都将转发到端口 3000 上的容器。

Go to the browser and open 127.0.0.1:9999 or localhost:9999, it should load the front end app.

转到浏览器并打开 127.0.0.1:9999 或 localhost:9999，它应该加载前端应用程序。

To delete a pod you can do with below command:

要删除 pod，您可以使用以下命令：

```
k delete pod/mymulticontainerapp
```

To inspect a pod you can use the below command:

要检查 Pod，您可以使用以下命令：

```
Front End:
kubectl exec mymulticontainerapp -c myfrontendapp -it /bin/sh
ls /usr/share/nginx/html
exit

Back End:
kubectl exec mymulticontainerapp -c mybackendapp -it /bin/sh
ls
exit
```

The above commands open a shell inside a container to interact.

上述命令在容器内打开一个 shell 进行交互。

Before we jump on to the next topic -deployments - it’s crucial that you understand metadata in a YAML file. Metadata is data about the containers. You can add labels and key-values to the metadata. You can also use the labels or key-values as selectors to identify a pod which will be used by deployments/services later.

在我们进入下一个主题 - 部署之前 - 了解 YAML 文件中的元数据至关重要。元数据是关于容器的数据。您可以向元数据添加标签和键值。您还可以使用标签或键值作为选择器来标识稍后将被部署/服务使用的 pod。

Let’s get all the pods labeled as **mymulticontainerapp**:

让我们将所有 pod 标记为 **mymulticontainerapp**：

```
kubectl get pods --selector=name=mymulticontainerapp
```

## Deployments

## 部署

You describe a desired state in a deployment, and the Deployment [Controller](https://kubernetes.io/docs/concepts/architecture/controller/) changes the actual state to the desired state at a controlled rate. You can define deployments to create new ReplicaSets, or to remove existing deployments and adopt all their resources with new deployments.

您在部署中描述了所需的状态，部署 [Controller](https://kubernetes.io/docs/concepts/architecture/controller/) 以受控速率将实际状态更改为所需状态。您可以定义部署以创建新的 ReplicaSet，或删除现有部署并在新部署中采用其所有资源。

It should look pretty similar to the previous YAML you saw. The salient point to note here is _kind -_ it’s a deployment rather than pod.

它看起来应该与您之前看到的 YAML 非常相似。这里要注意的重点是 _kind -_ 这是一个部署而不是 pod。

**Selector:** It is used to identify existing pods by labels, such as their metadata. If there is a pod already running with such a label, it will be part of this deployment.

**Selector:** 它用于通过标签识别现有的 pod，例如它们的元数据。如果已经有一个带有这样标签的 pod 正在运行，它将成为此部署的一部分。

**Replicas:** It specifies how many pods within this container you wish to create. If you say two, it will have two pods running the back end.

**副本：** 它指定您希望在此容器中创建多少个 pod。如果你说两个，它将有两个 Pod 运行后端。

Let’s say replicas is set to two and there are four pods already running with the label **app: node-backend**. It will terminate two of those pods to meet the requirement of two replicas. It will also scale up the pods by two if replicas is set to six.

假设副本数设置为 2，并且有四个 pod 已经在运行，标签为 **app: node-backend**。它将终止其中两个 pod 以满足两个副本的要求。如果副本数设置为 6，它还会将 pod 扩展两个。

**Spec:** It specifies the details about the container such as image name, container port, CPU/memory limits for the container inside the pod, etc.

**Spec:** 它指定了容器的详细信息，例如镜像名称、容器端口、pod 内容器的 CPU/内存限制等。

**Template:** It specifies the tags to be used for the newly created pods via the spec.

**模板：** 它通过规范指定用于新创建的 pod 的标签。

**\\*\\*\\***

**\\*\\***\\***

To create a deployment run the below command:

要创建部署，请运行以下命令：

```
kubectl apply -f backend.deploy.yml

Output:
deployment.apps/node-backend created
```

To see all the deployments:

要查看所有部署：

```
k get deployments

Output:
NAME           READY   UP-TO-DATE   AVAILABLE   AGE
node-backend   2/2     2            2           46s
```

To access the containers created during deployment you can use port-forward.

要访问在部署期间创建的容器，您可以使用 port-forward。

```
kubectl port-forward deployment/node-backend 3000:3000

You can access by going to localhost:3000/random
```

## Kubernetes Services

## Kubernetes 服务

An abstract way to expose an application running on a set of [pods](https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/) is as a network service.

公开在一组 [pods](https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/) 上运行的应用程序的抽象方法是作为网络服务。

Things to notice here - the kind is Service and the type is LoadBalancer. This creates a LoadBalancer on host port 3000 and proxies the request to container on 3000.

这里需要注意的事情 - 类型是服务，类型是 LoadBalancer。这会在主机端口 3000 上创建一个 LoadBalancer，并将请求代理到 3000 上的容器。

What is LoadBalancer balancing here? If you look closely we specified a selector. The selector searches for pods with label **app: node-background**and any request sent to the host on 3000 will be load balances among those pods. Here we have two replicas and the request will be load balanced to these two pods.

什么是 LoadBalancer 平衡？如果仔细观察，我们指定了一个选择器。选择器搜索带有标签 **app: node-background** 的 pod，发送到 3000 上的主机的任何请求都将在这些 pod 之间进行负载平衡。这里我们有两个副本，请求将负载均衡到这两个 pod。

To create a service run the below command:

要创建服务，请运行以下命令：

```
kubectl apply -f backend.service.yml.
```

To see all the services run the below command:

要查看所有服务，请运行以下命令：

```
k get service
```

This tells you the service name and what type of service it is. Here if you see there is a service called backend and of type LoadBalancer, which we just created using the service yml.

这会告诉您服务名称以及它是什么类型的服务。如果您看到这里有一个名为 backend 且类型为 LoadBalancer 的服务，我们刚刚使用服务 yml 创建了该服务。

You can do the same steps for frontend as well:

您也可以对前端执行相同的步骤：

```
kubectl apply -f backend.deploy.yml
kubectl apply -f backend.service.yml
```

## Wasn't That Easy?

## 不是那么容易吗？

We have looked at how to create a Kubernetets cluster and deploy a simple front end app that communicates with a backend app in Kubernetes. Hope this helps your understanding of the basics of Kubernetes. Kubernetes is a vast ocean and we just touched a tiny drop of it. But I hope this helps pave the path for your future Kubernetes learning!

我们已经了解了如何创建 Kubernetets 集群并部署一个简单的前端应用程序，该应用程序与 Kubernetes 中的后端应用程序进行通信。希望这有助于您了解 Kubernetes 的基础知识。 Kubernetes 是一片广阔的海洋，我们只是触及了其中的一小部分。但我希望这有助于为您未来的 Kubernetes 学习铺平道路！

**Useful Links:**

**有用的链接：**

[https://kubernetes.io/docs/concepts/](https://kubernetes.io/docs/concepts/)

[https://kubernetes.io/docs/concepts/](https://kubernetes.io/docs/concepts/)

[https://docs.docker.com/](https://docs.docker.com/)

[https://docs.docker.com/](https://docs.docker.com/)

* * *

* * *

**Srikant Vavilapalli**, Senior Software Engineer

**Srikant Vavilapalli**，高级软件工程师

Expert in designing/developing highly resilient applications on cloud.

在云上设计/开发高弹性应用程序的专家。

* * *

* * *

_DISCLOSURE STATEMENT: © 2020 Capital One. Opinions are those of the individual author. Unless noted otherwise in this post, Capital One is not affiliated with, nor endorsed by, any of the companies mentioned. All trademarks and other intellectual property used or displayed are property of their respective owners._

_披露声明：© 2020 资本一号。意见是个别作者的意见。除非本文另有说明，Capital One 不隶属于上述任何公司，也不受其认可。使用或展示的所有商标和其他知识产权均为其各自所有者的财产。_

### Cloud Container Adoption Report

### 云容器采用报告

Learn why 86% of tech leaders are prioritizing containers for more applications.

了解为什么 86% 的技术领导者优先考虑容器用于更多应用程序。

[Download Report](https://www.capitalone.com/tech/cloud-container-adoption-report/"")

[下载报告](https://www.capitalone.com/tech/cloud-container-adoption-report/"")

June 17, 2020

2020 年 6 月 17 日

#### Related Content

####  相关内容

Software Engineering [**Policy Enabled Kubernetes with Open Policy Agent**](http://www.capitalone.com/tech/software-engineering/policy-enabled-kubernetes-with-open-policy-agent)

软件工程 [**Policy Enabled Kubernetes with Open Policy Agent**](http://www.capitalone.com/tech/software-engineering/policy-enabled-kubernetes-with-open-policy-agent)

article \|January 11, 2019 

文章 \|2019 年 1 月 11 日

