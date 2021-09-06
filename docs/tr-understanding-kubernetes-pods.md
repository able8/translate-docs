# Understanding Kubernetes: Part 1-Pods

# 了解 Kubernetes：第 1 部分-Pods

Kubernetes, which is also known as 'k8s' is a portable, extendable open-source container management system (orchestrator) that automates the process of manually managing, deploying, and scaling the containerized applications like docker, containerd etc. This system was originally designed by Google; however, as of now, the Cloud Native Computing Foundation oversees all the arrangements. If you want to learn more about Kubernetes we have written a couple of blogs in the same topic,

Kubernetes，也称为“k8s”，是一种可移植、可扩展的开源容器管理系统（编排器），可自动执行手动管理、部署和扩展容器化应用程序（如 docker、containerd 等）的过程。该系统最初是设计的由谷歌提供；然而，截至目前，云原生计算基金会负责监督所有安排。如果您想了解更多关于 Kubernetes 的信息，我们已经写了几篇关于同一主题的博客，

- [What is Kubernetes?](https://goglides.io/what-is-kubernetes/529/)
- [Kubernetes !!! How it Works.](https://goglides.io/kubernetes-how-it-works/94/)


So, to jump in the depths about this popular orchestration system for the management of multi-container applications at a scale, we have come up with a brief description of its first part-Pods

因此，为了深入了解这个用于大规模管理多容器应用程序的流行编排系统，我们简要介绍了它的第一部分 - Pods

## So, What exactly are the Pods?

## 那么，Pod 到底是什么？

In simple words, the pods are the group of one or more containers that are deployed to a single node, which represents the process running on the cluster. Unlike other systems introduced in the past, what users may find different is how the system doesn’t directly run the containers. Instead of that, the Kubernetes combines one or more containers into a high-level structure, which are known as pods. All of the containers in the same pod share the same IP address, hostname, IPC(International Patent Classification), including other resources . The container can communicate with each other as long as they are in the same pod, they act as they are the integrated part of a single machine and at the same time maintaining some level of isolation from other pods. There is also the possibility of the individual applications within a Pod’s context may have some further sub-isolations part applied. In the matter of containers in different pods, the containers have a different IP address that varies from each pod and cannot communicate with each other without a special configuration.

简单来说，pods 是一组部署在单个节点上的一个或多个容器，代表在集群上运行的进程。与过去引入的其他系统不同，用户可能会发现不同的是系统不直接运行容器。取而代之的是，Kubernetes 将一个或多个容器组合成一个高层结构，称为 pods。同一个 pod 中的所有容器共享相同的 IP 地址、主机名、IPC（国际专利分类），包括其他资源.只要它们在同一个 Pod 中，容器就可以相互通信，它们就像是单个机器的集成部分一样，同时与其他 Pod 保持一定程度的隔离。 Pod 上下文中的单个应用程序也有可能应用一些进一步的子隔离部分。容器在不同的pod中的问题，每个pod的容器IP地址不同，没有特殊的配置是不能互相通信的。

Further, the Kubernetes uses the whole pod as the smallest deployable unit instead of a single container; there are also good reasons for it. As the pods represent the layer of abstractions, and it helps the system to wrap the container that should be managed as a single entity. The containers in the pod have access to the shared volumes (a directory that is accessible to all the containers that are running in a pod) that are used to mount each container’s filesystems. Generally, the pods are used as Kubernetes’s replicating units. In case the application of the user becomes very popular, and the load cannot be handled by a single pod, the users can configure the Kubernetes to deploy the new replicas of their pod to the cluster as per their need. Even when the load isn’t that heavy, it is best to have multiple copies of the pod running in the production system to ensure that the load is balanced and so that the system is failure resistance. So, what are the other uses of pods?

此外，Kubernetes 使用整个 pod 作为最小的可部署单元，而不是单个容器；这也有充分的理由。由于 Pod 代表抽象层，它有助于系统将应该作为单个实体管理的容器进行包装。 Pod 中的容器可以访问用于挂载每个容器的文件系统的共享卷（一个可供在 Pod 中运行的所有容器访问的目录）。通常，Pod 用作 Kubernetes 的复制单元。如果用户的应用变得非常流行，并且负载无法由单个 Pod 处理，用户可以根据需要配置 Kubernetes 将其 Pod 的新副本部署到集群中。即使负载不那么重，最好在生产系统中运行多个 pod 副本，以确保负载平衡，从而使系统具有抗故障能力。那么，pods 的其他用途是什么？

## An example of the pod:

## Pod 示例：

```
apiVersion: v1
kind: Pod
metadata:
name: example-app
labels:
    app: example-app
    version: v1
    role: backend
spec:
containers:
  - name: java
    image: companyname/java
    ports:
    - containerPort: 443
    volumeMounts:
    - mountPath: /volumes/logs
      name: logs
  - name: logger
    image: companyname/logger:v1.2.3
    ports:
    - containerPort: 9999
    volumeMounts:
    - mountPath: /logs
      name: logs
  - name: monitoring
    image: companyname/monitoring:v4.5.6
    ports:
    - containerPort: 1234
```


## Uses of Pods

## Pod 的使用

- The pods act as the mode of the pattern of several co-existing processes that form a single cohesive unit of service. They make the deployment of the application and its management easily accessible, providing the high-level of abstraction for the user and controller to run the program specific to requirements.

- They are used to abstract the network and storage from the underlying container.

- The pods make moving of containers around the cluster more easily accessible as the containers are managed in an organized way and labeled according to their nature. 

  

- Pod 充当多个共存进程模式的模式，这些进程形成一个单一的内聚服务单元。它们使应用程序的部署及其管理易于访问，为用户和控制器提供高级抽象以运行特定于需求的程序。

- 它们用于从底层容器中抽象出网络和存储。

- 由于容器以有组织的方式进行管理并根据其性质进行标记，因此 Pod 使得在集群中移动容器变得更加容易。

- The pods allow data sharing and communication between their constituents. The containers in a pod use the same IP and port space, so they can locate each other and communicate using the ‘localhost.’ The containers coordinate their usage of ports and perform together as needed by the application.

- Pods 允许其成员之间的数据共享和通信。 Pod 中的容器使用相同的 IP 和端口空间，因此它们可以相互定位并使用“localhost”进行通信。容器协调它们对端口的使用，并根据应用程序的需要一起执行。

## Types of Pods

## Pod 的类型

Customarily, the pods in the Kubernetes cluster can be used in two major ways:

通常情况下，Kubernetes 集群中的 pod 主要有两种使用方式：

**1\. Pods for a single container**

**1\. 单个容器的 Pod **

The single-container pod is the commonly used pod system in the Kubernetes. As the system of Kubernetes uses the pods as the miniature unit of the system and doesn’t directly deal with the containers, the one-container pod wraps the single container to interact with the system. So, these containers are suitable for applications that are composed of just a single container.

单容器 Pod 是 Kubernetes 中常用的 Pod 系统。由于Kubernetes的系统以pods作为系统的微型单元，并不直接与容器打交道，因此单容器pod将单个容器包裹起来与系统进行交互。因此，这些容器适用于仅由单个容器组成的应用程序。

**2\. Pods for multiple containers that need to work as a single cohesive unit**

**2\. 需要作为单个内聚单元工作的多个容器的 Pod**

In case the application runs with the multiple containers, the multi-container pods tightly couples containers that need to share the resources. The co-located containers in this type of pods act as the single cohesive unit sharing the volume and working as the need for the application. The significant purpose of a multi-container pod is to co-locate, co-manage, and support the helper process for the primary application. However, there are some downside like, if you have 3 containers inside same pod, let say 1,2, and 3. If container 1 needs scaling, we need to scale all 3. This is simply wasting memory and resource usages.

如果应用程序与多个容器一起运行，则多容器 Pod 将需要共享资源的容器紧密耦合。这种类型的 pod 中并置的容器充当单个内聚单元，共享卷并根据应用程序的需要工作。多容器 Pod 的重要目的是共同定位、共同管理和支持主要应用程序的辅助进程。但是，也有一些缺点，例如，如果同一个 Pod 中有 3 个容器，比如说 1、2 和 3。如果容器 1 需要扩展，我们需要扩展所有 3 个。这只是浪费内存和资源使用。

## Pods and Controllers

## Pod 和控制器

The Kubernetes pod works as a group of containers that are deployed together to interact with the system, on the other hand, the controller is the manager in the system that maintains the pods and makes sure that they set on a specified state. The controllers also play a very crucial role in the system as they are in charge of replacing the pods, in case they fail to perform. Similarly, the controllers are also responsible for running the required number of the pods’ replicas in the whole cluster in the situations if they get deleted or terminated.

Kubernetes pod 作为一组部署在一起与系统交互的容器工作，另一方面，控制器是系统中维护 pod 并确保它们设置为指定状态的管理器。控制器在系统中也起着非常重要的作用，因为它们负责更换吊舱，以防它们无法执行。同样，控制器还负责在整个集群中运行所需数量的 pod 副本，以防它们被删除或终止。

The desired state of the replication controller is customized according to the need of the application deployer. When the change is made in the properties of the controller to get it in the desired state, the systems are updated in order to meet the specified state. The controller can also be taken as the process supervisor that ensures the specified number of processes are running smoothly on a single server or on several servers. In cases, if there are too many pods running in the system, the controller stops the surplus pods, and if there are only a few numbers of pods, the controller starts the new ones.

复制控制器的期望状态是根据应用程序部署者的需要定制的。当控制器的属性发生变化以使其处于所需状态时，系统会更新以满足指定状态。控制器还可以作为进程监督者，确保指定数量的进程在单个服务器或多个服务器上顺利运行。在某些情况下，如果系统中运行的 Pod 过多，控制器会停止多余的 Pod，如果只有少数 Pod，控制器会启动新的 Pod。

## Pod Templates 

## Pod 模板

The Pod templates field is the component of the controller. The Pod templates contain a ‘pod specification’ field that determines how a pod should run, including how the containers should work together within a pod and the volume the pod should mount. The controller uses the Pod templates to create a pod and manage it’s stated as desired by the application. Each controller separated for a workload resource uses the templates as the desired state of workload object necessary to run the application. Also, the characteristics of a pod can be changed using the new template, all the pods running in the system are not bound to accept the new changes and the changes templates doesn't affect them, however, the new pods will reflect the characteristics mention in the new template. The pods won’t receive the templates updates directly, instead of that, the controller creates a new pod to match requirements set by the new templates.

Pod 模板字段是控制器的组件。 Pod 模板包含一个“Pod 规范”字段，用于确定 Pod 应如何运行，包括容器应如何在 Pod 内协同工作以及 Pod 应安装的卷。控制器使用 Pod 模板来创建一个 Pod 并根据应用程序的需要管理它。为工作负载资源分离的每个控制器使用模板作为运行应用程序所需的工作负载对象的所需状态。另外，Pod 的特性可以使用新的模板进行更改，系统中运行的所有 Pod 不一定接受新的更改，更改模板不会影响它们，但是，新的 Pod 会反映提到的特性在新模板中。 Pod 不会直接接收模板更新，相反，控制器会创建一个新的 Pod 来匹配新模板设置的要求。

```
apiVersion: batch/v1
kind: Job
metadata:
name: hello
spec:
template:
# This is the pod template
spec:
containers:
- name: hello
image: busybox
command: ['sh', '-c', 'echo "Hello, Kubernetes!"&& sleep 3600']
restartPolicy: OnFailure
```


## Creating pods

## 创建 Pod

The pods are created by using the controller,such as Deployment, which creates and destroys the replicas of the pods as per necessity. As the pods are transitory, they are no created directly; after creating the pods, they are scheduled to run specific to the needs in the user’s cluster. Further, the controllers also create and manage the pods to meet the specification as the pods are not able to repair or replace themselves when needed. These controllers also are important in updating the pods and changing the version of the application running in the container.

Pod 是通过使用控制器（例如 Deployment）创建的，它根据需要创建和销毁 Pod 的副本。由于 Pod 是暂时的，它们不是直接创建的；创建 Pod 后，它们将被安排运行特定于用户集群中的需求。此外，控制器还创建和管理吊舱以满足规范，因为吊舱无法在需要时自行修复或更换。这些控制器在更新 Pod 和更改容器中运行的应用程序版本方面也很重要。

The pods are labeled taking their name, role, nature, version, etc. into consideration. Then the pods are selected by the replication controllers depending on the similarities for a certain role, as per the need of the applications and also for more complex action depending on the demand of the user. The labeling system is exceptionally flexible by design, and the users can experiment with the practices that work best for them.

Pod 被标记为考虑了它们的名称、角色、性质、版本等。然后，复制控制器根据特定角色的相似性，根据应用程序的需要以及根据用户的需求进行更复杂的操作来选择 pod。标签系统的设计非常灵活，用户可以尝试最适合他们的做法。

Basically, there are thee way of creating a pod in the Kubernetes cluster:

基本上，在 Kubernetes 集群中创建 pod 的方法有以下几种：

**1\. Imperative Method**

**1\.命令式方法**

The pods can be created with the kubectl command that is directly applied to the Kubernetes cluster. For example, if the user wants to deploy the Nginx web server into the cluster, the user can directly create a pod.

可以使用直接应用于 Kubernetes 集群的 kubectl 命令创建 pod。例如，如果用户想将 Nginx web 服务器部署到集群中，用户可以直接创建一个 pod。

```
$ kubectl run --generator=run-pod/v1 nginx --image=nginx
pod/nginx created

$ kubectl get pods nginx
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          35s
```


**2\. Declarative Method** In this process, the pods are created by writing the manifests and using the `kubectl apply`. This manifest contains the metadata and the specification details w.r.t that help the user to create the object they want. This method is also quite helpful for updating the objects running in the Kubernetes cluster.

**2\.声明性方法** 在此过程中，通过编写清单并使用“kubectl apply”来创建 Pod。此清单包含元数据和规范详细信息 w.r.t，可帮助用户创建他们想要的对象。这种方法对于更新在 Kubernetes 集群中运行的对象也很有帮助。

First create an nginx-deployment file in declarative way:

首先以声明方式创建一个 nginx-deployment 文件：

`vim simple_deployment.yaml`

```
apiVersion: apps/v1
kind: Deployment
metadata:
name: nginx-deployment
spec:
selector:
    matchLabels:
      app: nginx-deployment
minReadySeconds: 5
template:
    metadata:
      labels:
        app: nginx-deployment
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```


Now, let’s create a pod using simple\_deployment.yaml file, we have just created and validate.

现在，让我们使用 simple\_deployment.yaml 文件创建一个 pod，我们刚刚创建并验证。

```
$ kubectl create -f simple_deployment.yaml
deployment.apps/nginx-deployment created

$ kubectl get deployment nginx-deployment
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
nginx-deployment   1/1     1            1           19s

$ kubectl get pods -l app=nginx-deployment
NAME                                READY   STATUS    RESTARTS   AGE
nginx-deployment-5684d7c768-t8xsc   1/1     Running   0          38s
```




**3\. Using an Rest API interface** The API server performs the authentication, authorization, and admission control of the clients. In the cluster, the API is exported on port 443, which can be accessed through the TLS connection. The self-signed certificate which is generated during the cluster configuration is available at ‘$USER/.kube/config’ on the client machine. A simple nginx pod:

**3\.使用 Rest API 接口** API 服务器执行客户端的身份验证、授权和准入控制。在集群中，API在443端口导出，可以通过TLS连接访问。在集群配置期间生成的自签名证书在客户端机器上的“$USER/.kube/config”中可用。一个简单的 nginx pod：

```
cat > nginx-pod.json <<EOF
{
"apiVersion": "apps/v1",
"kind": "Deployment",
"metadata": {
    "name": "nginx-deployment"
},
"spec": {
    "selector": {
      "matchLabels": {
        "app": "nginx"
      }
    },
    "minReadySeconds": 5,
    "template": {
      "metadata": {
        "labels": {
          "app": "nginx"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx:1.14.2",
            "ports": [
              {
                "containerPort": 80
              }
            ]
          }
        ]
      }
    }
}
}
EOF

```


## Lifecycle of Pods

## Pod 的生命周期

[![kubernetes-life-cycle-of-a-pod](https://i2.wp.com/goglides.io/wp-content/uploads/2020/06/kubernetes-life-cycle-of-a-pod.png?resize=1024%2C810&ssl=1)](https://i0.wp.com/github.com/pandeybk/content-goglides-io/blob/part1-pods-sachin/Understanding-Kubernetes-Part1-Pods/media/kubernetes-life-cycle-of-a-pod.png?ssl=1)

**Image courtesy:** [Joe Beda’s Blog.](https://blog.heptio.com/core-kubernetes-jazz-improv-over-orchestration-a7903ea92ca)

**图片提供：** [Joe Beda 的博客。](https://blog.heptio.com/core-kubernetes-jazz-improv-over-orchestration-a7903ea92ca)

The status of pods can determine in which phase of its lifecycle. The pods are not designed to run forever, and the terminated pods cannot be brought back. Pods don’t cease to exist unless the user himself or the controller deletes the pods. The pods in the system have a PodStaus API object which determines the current life phase of the pod. The pods themselves publish their phase to the PodStatus, and the user or controller can take the necessary actions accordingly. The lifecycle of the pods that are created using cronjobs or jobs can be classified into five phases:

Pod 的状态可以决定其生命周期的哪个阶段。 Pod 并非设计为永远运行，并且终止的 Pod 无法恢复。除非用户本人或控制器删除 Pod，否则 Pod 不会停止存在。系统中的 Pod 有一个 PodStaus API 对象，用于确定 Pod 的当前生命周期。 Pod 本身将其阶段发布到 PodStatus，用户或控制器可以相应地采取必要的操作。使用 cronjobs 或作业创建的 Pod 的生命周期可以分为五个阶段：

**1\. Pending**

**1\. 待办的**

This is the state of the pod when it has been created and already accepted by the cluster; however, all the containers inside the pod are not fully functional at the moment. This phase determines the missing container for the pod to be fully functional.

这是 pod 在创建并已被集群接受时的状态；但是，目前 pod 内的所有容器都不能完全正常工作。此阶段确定 pod 缺少的容器是否具有完整功能。

**2\. Running** The running state of the pod represents it’s successful bound to a node after the completion of creating all the necessary containers in it. In this phase, at least one container is running; it either in the process of starting or restarting.

**2\.运行** Pod 的运行状态表示在其中创建完所有必要的容器后，它已成功绑定到节点。在这个阶段，至少有一个容器在运行；它在启动或重新启动的过程中。

**3\. Succeeded**

**3\.成功**

The succeeded phase of the pod’s lifecycle determines the deletion of all the containers in a specific pod. Whether it be they are out of the use or need new replacements, the pod is emptied deleting all the containers. The terminated pods do not start again after being removed from the system.

Pod 生命周期的成功阶段决定删除特定 Pod 中的所有容器。无论它们是不再使用还是需要新的替换，pod 都会被清空，删除所有容器。从系统中删除后，终止的 Pod 不会再次启动。

**4\. Failed**

**4. 失败的**

The failed phase of the pods states the termination of all the containers in a pod, but at least one container has terminated in failure. Such cases can occur if a container exists with a non-zero status or the container was terminated by the system itself.

pods 的失败阶段表示一个 pod 中的所有容器都终止了，但至少有一个容器因失败而终止。如果容器以非零状态存在或容器被系统本身终止，则可能会发生这种情况。

**5\. Unknown**

**5. 未知**

The unknown phase is published by the pods in PodStatus when the state of that particular pod cannot be determined.

当无法确定特定 Pod 的状态时，未知阶段由 PodStatus 中的 Pod 发布。

Furthermore, to provide brief details about the condition of the pods and what is cause for that status of the pods, the PodStaus also contains an array named 'PodCondtions.' The PodCondition represents the condition of the pod, indicating the conditions that are the cause for the current phase of the pods. This is the helpful component of the system that aids the user in taking the appropriate actions or troubleshoot the pods taking the situation into consideration.

此外，为了提供有关 Pod 状态的简要详细信息以及导致 Pod 状态的原因，PodStaus 还包含一个名为“PodCondtions”的数组。PodCondition 表示 Pod 的状态，指示导致原因的条件对于 pod 的当前阶段。这是系统的有用组件，可帮助用户根据情况采取适当的操作或对 Pod 进行故障排除。

The ‘PodConditions’ are represented in the pod manifest as ‘conditions’; the field has a 'type' and 'status.' The 'field' contains the 'PodScheduled,' 'Ready,' 'Initialized,' and 'Unschedulable' condition of the whereas, the 'status' field determines 'type' field, including 'True,' 'False' or 'Unknown' statuses.

‘PodConditions’ 在 pod manifest 中表示为 ‘conditions’；该字段具有“type”和“status”。“field”包含“PodScheduled”、“Ready”、“Initialized”和“Unschedulable”条件，而“status”字段确定“type”字段，包括“真”、“假”或“未知”状态。

## Termination of pods 

## Pod 的终止

[![kubernetes-termination-of-a-pod](https://i1.wp.com/goglides.io/wp-content/uploads/2020/06/kubernetes-termination-of-a-pod.png?resize=1024%2C694&ssl=1)](https://i1.wp.com/github.com/pandeybk/content-goglides-io/blob/part1-pods-sachin/Understanding-Kubernetes-Part1-Pods/media/kubernetes-termination-of-a-pod.png?ssl=1)



**Image courtesy:** [Harshal Shah’s Blog.](https://dzone.com/articles/kubernetes-lifecycle-of-a-pod)

**图片提供：** [Harshal Shah 的博客。](https://dzone.com/articles/kubernetes-lifecycle-of-a-pod)

As the pods represent the processing running on the cluster, after their processes are complete, they can be gracefully terminated. Kubernetes provides the 30 seconds grace period for the deletion of the pods that are no longer required. The grace period can be configured according to the wish of the user by setting the ‘–grace-period’ while interacting with the cluster to request the termination.

由于 Pod 代表在集群上运行的处理，因此在它们的进程完成后，它们可以正常终止。 Kubernetes 为删除不再需要的 pod 提供了 30 秒的宽限期。宽限期可以根据用户的意愿通过在与集群交互以请求终止时设置“--grace-period”来配置。

**Here are the steps for the termination of a pod:**

**以下是终止 Pod 的步骤：**

**Step 1:** The users send commands or API calls to terminate the pod.

**第 1 步：** 用户发送命令或 API 调用来终止 Pod。

**Step 2:** The Kubernetes updates the Pod status with the time period beyond which the pod is considered ‘dead’(the termination request period plus the grace period set by the user).

**第 2 步：** Kubernetes 使用 Pod 被视为“死亡”的时间段（终止请求期加上用户设置的宽限期）更新 Pod 状态。

**Step 3:** The Kubernetes system marks the pod as ‘Terminating and stops sending traffics to the pod.

**第 3 步：** Kubernetes 系统将 Pod 标记为“终止并停止向 Pod 发送流量”。

**Step 4:** The termination request is sent to the main process ‘PID 1’ in each container of the pod, and the ‘grace period’ countdown commences. Then the containers start the graceful shutdown of the running application and exit it.

**第 4 步：** 终止请求被发送到 Pod 的每个容器中的主进程“PID 1”，“宽限期”倒计时开始。然后容器开始正常关闭正在运行的应用程序并退出它。

**Step 5:**  In case, if a container doesn’t terminate within the marked grace period, a ‘SIGKILL’ is sent by the system to terminate the container violently.

**第 5 步：** 如果容器未在标记的宽限期内终止，系统将发送“SIGKILL”以暴力终止容器。

**Step 6:**  The Kubernetes finishes deleting the pod on the API server on the Kubernetes Master and is no longer visible from the client. 

**第 6 步：** Kubernetes 在 Kubernetes Master 上的 API 服务器上完成删除 Pod，并且不再从客户端可见。

