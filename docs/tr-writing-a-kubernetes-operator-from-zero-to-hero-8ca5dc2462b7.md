# Writing a Kubernetes Operator: From Zero to Hero

# 编写 Kubernetes Operator：从零到英雄

[Apr 22·11 min read](https://anupamgogoi.medium.com/writing-a-kubernetes-operator-from-zero-to-hero-8ca5dc2462b7?source=post_page-----8ca5dc2462b7--------------------------------)

# Introduction

#  介绍

In this article, I am going to explain in detail how to create your own  Kubernetes operator from zero. Operators are software extensions that make use of the custom resources (or kind in k8s paradigm) to manage the applications. To know more about operators please read the official [documentation](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

在本文中，我将详细解释如何从零创建您自己的 Kubernetes 算子。 Operators 是利用自定义资源（或 k8s 范式中的种类）来管理应用程序的软件扩展。如需了解更多关于运营商的信息，请阅读官方[文档](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)。

In dummy language, let's say we have a HelloApp application. To deploy the HelloApp application we will create the below k8s resource.

在虚拟语言中，假设我们有一个 HelloApp 应用程序。为了部署 HelloApp 应用程序，我们将创建以下 k8s 资源。

![img](https://miro.medium.com/max/1400/1*qDKXx5geD_Zh6ed3c6n3lg.png)

Note the **kind: HelloApp** is our custom resource definition (CRD) and the code that handles this CRD is  our custom Operator (or Controller). And the focus of this article is to create this Operator (or Controller) from zero.

注意 **kind:HelloApp** 是我们的自定义资源定义（CRD），处理这个 CRD 的代码是我们的自定义 Operator（或 Controller）。而本文的重点是从零开始创建这个 Operator（或 Controller）。

# The Myth

#  神话

Before I started creating an Operator, I thought that Go was the only language that was meant for writing Operators. But it's a myth.

在我开始创建 Operator 之前，我认为 Go 是唯一用于编写 Operator 的语言。但这是一个神话。

> You also implement an Operator (that is, a Controller) using any language / runtime that can act as a [client for the Kubernetes API](https://kubernetes.io/docs/reference/using-api/client-libraries/).

> 您还可以使用任何可以充当 [Kubernetes API 客户端](https://kubernetes.io/docs/reference/using-api/client-libraries) 的语言/运行时来实现 Operator（即控制器) /)。

Check this screenshot from k8s documentation itself.

从 k8s 文档本身检查此屏幕截图。

![img](https://miro.medium.com/max/1400/1*lY77GbOX-zxm0m7jRWg9sg.png)

Kubernetes itself is a massive monitoring system. All its functionalities can be  accessed through the APIs (the API Server). So, if we can write an application that can access (client) to the k8s API Server, our  application can do the necessary actions to become an Operator. So,  language does not matter to implement an Operator. However, the Golang, being the native language of the k8s runtime and due to its vast amount of libraries present for Operator implementation, it's the most  preferred one for building operators.

Kubernetes 本身就是一个庞大的监控系统。它的所有功能都可以通过 API（API 服务器）访问。因此，如果我们可以编写一个可以访问（客户端）k8s API Server 的应用程序，我们的应用程序就可以执行必要的操作来成为 Operator。因此，语言与实现 Operator 无关。然而，Go lang 是 k8s 运行时的本地语言，并且由于其为 Operator 实现提供了大量库，因此它是构建 Operator 的首选语言。

# Softwares needed for this tutorial

# 本教程需要的软件

The following software will be required for this tutorial.

本教程需要以下软件。

1. [Go lang (1.16)](https://golang.org/)
2. [Operator SDK (1.5)](https://sdk.operatorframework.io/)
3. [Kind](https://kind.sigs.k8s.io/)
4. Visual Studio Code with Go plugin


# Tutorial Flow

# 教程流程

I will divide the article into the following parts to make it modular.

我将文章分为以下几个部分，使其模块化。

> Part1: Creating the Operator Project
>
> Part2: Implement the Operator Logic
>
> Part 3: Generating CRDs
>
> Part 4: Installing the CRDs
>
> Part 5: Running the Operator outside the Cluster
>
> Part 6: Debugging the Operator outside the Cluster
>
> Part7: Running the Operator inside the Cluster

> 第 1 部分：创建 Operator 项目
>
> 第 2 部分：实现运算符逻辑
>
> 第 3 部分：生成 CRD
>
> 第 4 部分：安装 CRD
>
> 第 5 部分：在集群外运行 Operator
>
> 第 6 部分：在集群外调试 Operator
>
> Part7：在集群内运行Operator

# Part 1: Creating the Operator Project

# 第 1 部分：创建 Operator 项目

We are going to use the [Operator SDK](https://sdk.operatorframework.io/) to create the project structure. At the moment of writing the article, I was using Operator SDK version 1.5. To know in detail how to create a  project using the SDK please read this official [documentation](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/). For the sake of brevity, I will explain the basic steps only to create the project.

我们将使用 [Operator SDK](https://sdk.operatorframework.io/) 创建项目结构。在撰写本文时，我使用的是 Operator SDK 1.5 版。要详细了解如何使用 SDK 创建项目，请阅读此官方[文档](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/)。为简洁起见，我将仅解释创建项目的基本步骤。

```
$ mkdir demo-operator
$ cd demo-operator
$ operator-sdk init --domain anupam.com --repo github.com/anupamgogoi/demo-operator
```

## Create the API and the Controller

## 创建 API 和控制器

```
$ operator-sdk create api --group apps --version v1 --kind HelloApp --resource --controller
```

With the above commands, our minimal project structure is ready to work on.

使用上述命令，我们的最小项目结构就可以开始工作了。

![img](https://miro.medium.com/max/1400/1*sGWUZRdlEqcUimldGuY1pA.png)

A brief notes on the important files created.

关于创建的重要文件的简要说明。

1. It's the Makefile with all the necessary commands we need to generate the artifacts for the operator. Execute **make help** and it will show all the available commands available to execute.
2. It's the central point of entry to the operator that contains the main  function. Also, the central point of entry for debugging the operator in the local cluster.
3. The controller. The main logic of the operator goes here.
4. The structure for our custom resource. 



1. 它是包含我们为操作员生成工件所需的所有必要命令的 Makefile。执行 **make help** 它将显示所有可用于执行的命令。
2. 它是包含主函数的操作符的中心入口点。此外，在本地集群中调试操作员的中心入口点。
3. 控制器。运算符的主要逻辑在这里。
4. 我们自定义资源的结构。
5. The group and version info we specified during the creation of the  operator. To know more about the group and version please read this  official [documentation](https://book.kubebuilder.io/cronjob-tutorial/gvks.html).
6. 我们在创建操作符时指定的组和版本信息。要了解有关组和版本的更多信息，请阅读此官方 [文档](https://book.kubebuilder.io/cronjob-tutorial/gvks.html)。

# Part 2: Implement the Operator Logic

# 第 2 部分：实现运算符逻辑

Our operator logic is very simple. When the below CRD is applied, the  operator (controller) should create a deployment of the kind HelloApp  with the number of pods specified in the **spec.size**

我们的操作符逻辑非常简单。应用以下 CRD 时，操作员（控制器）应使用 **spec.size** 中指定的 pod 数量创建 HelloApp 类型的部署

![img](https://miro.medium.com/max/1400/1*RUMDCno-FLs6J1a7yJWefw.png)

The complete source code can be found in this [repository](https://github.com/anupamgogoi-wso2/go-apps/tree/master/demo-operator). So our Spec has only two fields namely image and size. To add them to the spec, edit the file [[4](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/api/v1/helloapp_types.go)] showed in Part 1 of the article.

完整的源代码可以在这个 [repository](https://github.com/anupamgogoi-wso2/go-apps/tree/master/demo-operator) 中找到。所以我们的 Spec 只有两个字段，即 image 和 size。要将它们添加到规范中，请编辑文件 [[4](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/api/v1/helloapp_types.go)]文章的第 1 部分。

![img](https://miro.medium.com/max/1400/1*Z_L1kaVOMmQZyrhLg4Ah1A.png)

The controller logic can be found [here](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/controllers/helloapp_controller.go). It's doing no magic. First, it checks if there is deployment for the  HelloApp and if not it tries to create the deployment. Eventually, it  checks if the number of pods is the same as we expect in the **spec.size** and if not create more pods. That's it. To know more about the API reference, please read this [documentation](https://sdk.operatorframework.io/docs/building-operators/golang/references/).

控制器逻辑可以在这里找到。它没有任何魔法。首先，它检查是否有 HelloApp 部署，如果没有，则尝试创建部署。最终，它会检查 pod 的数量是否与我们在 **spec.size** 中预期的相同，如果不是，则创建更多的 pod。就是这样。要了解有关 API 参考的更多信息，请阅读此 [文档](https://sdk.operatorframework.io/docs/building-operators/golang/references/)。

![img](https://miro.medium.com/max/1400/1*LI5TWiVThcTS6ksGd_Jwbw.png)

The controllers implement the **Reconciler** interface which exposes the **Reconcile** method. The Reconcile function is called for cluster events like CRUD  operations and thus it can compare the actual status of the resource  (kind) with the expected status (spec) and then if necessary reconcile  it.

控制器实现 **Reconciler** 接口，该接口公开 **Reconcile** 方法。 Reconcile 函数被调用用于像 CRUD 操作这样的集群事件，因此它可以将资源的实际状态（种类）与预期状态（规格）进行比较，然后在必要时进行协调。

![img](https://miro.medium.com/max/1400/1*dOsI0JFt5Yw6acKRBpap7g.png)

## main.go

When the main.go is generated by the Operator-SDK it creates some extra  stuff. For simplicity, I removed the extra stuff and left only what is  necessary to run the operator. Take a look at the [main.go](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/main.go) file.

当 main.go 由 Operator-SDK 生成时，它会创建一些额外的东西。为简单起见，我删除了额外的东西，只留下运行操作符所需的东西。查看 [main.go](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/main.go) 文件。

In the main.go file, there are the below codes. This is the most important part of the code.

在 main.go 文件中，有以下代码。这是代码中最重要的部分。

![img](https://miro.medium.com/max/1400/1*GigTt5ZG3oJ0xMiSwDINnQ.png)

**ctrl.GetConfigOrDie()** will try to read the k8s cluster config from **~/.kube/config** file and based on that it will get the connection information. Here is how a **~/.kube/config** looks like.

**ctrl.GetConfigOrDie()** 将尝试从 **~/.kube/config** 文件中读取 k8s 集群配置，并基于此获取连接信息。这是 **~/.kube/config** 的样子。

![img](https://miro.medium.com/max/1400/1*J2NFkiQmF97OBJBSqKJ2pA.png)

You can see that the config file has information such as server IP, certificate, etc. This is the most important thing. The **GetConfigOrDie()** method will read this information and based on that **ctrl.NewManager()** will create the manager for our controller. The rest of it is just calling  the APIs of the K8s API Server. That's the magic. Just take a look into  the main.go file and the things will be crystal clear.

可以看到配置文件有服务器IP、证书等信息，这是最重要的。 **GetConfigOrDie()** 方法将读取此信息，并基于该信息 **ctrl.NewManager()** 将为我们的控制器创建管理器。剩下的只是调用 K8s API Server 的 API。这就是魔法。只需查看 main.go 文件，事情就会一清二楚。

Even the **kubectl** CLI uses the API calls to the k8s API server. Just execute the below command in the console and you can verify.

甚至 **kubectl** CLI 也使用对 k8s API 服务器的 API 调用。只需在控制台中执行以下命令即可验证。

```
$ kubectl get nodes --v=8
```

![img](https://miro.medium.com/max/1400/1*DKbXjD64lQXoJZ956UhP4w.png)

# Part 3: Generating CRDs

# 第 3 部分：生成 CRD

At this point, the logic of our HelloApp controller is ready. Now, we will need to generate the CRD for it. Go to the root of the project **demo-operator** and execute the below command.

至此，我们的HelloApp控制器的逻辑就准备好了。现在，我们需要为它生成 CRD。转到项目 **demo-operator** 的根目录并执行以下命令。

```
$ make manifests
```

It will generate the CRDs for us in this location.

它将在此位置为我们生成 CRD。

**~/demo-operator/config/crd/bases**

![img](https://miro.medium.com/max/1400/1*PVL5lMXY5_4keLicVLdfBw.png)

Also, in the **~/demo-operator/config/samples** directory, it will generate a sample for us.

另外，在 **~/demo-operator/config/samples** 目录中，它会为我们生成一个示例。

![img](https://miro.medium.com/max/1400/1*VG3vTE8YXinFKlEF9rOtig.png)

# Part 4: Installing the CRDs

# 第 4 部分：安装 CRD

Before you can run the operator you need a local cluster. Use [kind](https://kind.sigs.k8s.io/) to create a cluster in your local host.

在您可以运行操作员之前，您需要一个本地集群。使用 [kind](https://kind.sigs.k8s.io/) 在本地主机中创建集群。

```
$ kind create cluster --name k8s
```

There are two ways to install the CRDs in the cluster. Simply run the below command,

有两种方法可以在集群中安装 CRD。只需运行以下命令，

```
$ make install
```

Or navigate to the **~/demo-operator/config/crd/bases** and execute the below command.

或者导航到 **~/demo-operator/config/crd/bases** 并执行以下命令。

```
$ kubectl apply -f .
```

Both do the same task.

两者都做同样的任务。

# Part 5: Running the Operator outside the Cluster 

# 第 5 部分：在集群外运行 Operator

It's the simplest way to Run & Debug your operator logic. Also, to  discover the internals of the k8s, it's the best option to start with.

这是运行和调试操作员逻辑的最简单方法。此外，要发现 k8s 的内部结构，这是开始的最佳选择。

```
$ cd demo-operator
$ go run main.go
```

That's it. You should be able to see the below output.

就是这样。您应该能够看到以下输出。

![img](https://miro.medium.com/max/1400/1*dfrD8hziYHVMwSXtW6847g.png)

Let's deploy our custom resource. Let's navigate to the samples directory and apply the resource.

让我们部署我们的自定义资源。让我们导航到示例目录并应用资源。

```
$ kubectl create ns test
$ kubectl apply -f apps_v1_helloapp.yaml -n test
```

Let's check what is created inside the **test** namespace.

让我们检查在 **test** 命名空间中创建了什么。

```
$ kubectl get all -n test
```

![img](https://miro.medium.com/max/1206/1*Hd51oIHZ9msVCRVInbBUOw.png)

Now the most enticing part.

现在是最诱人的部分。

```
$ kubectl get HelloApp -n test
```

![img](https://miro.medium.com/max/846/1*oXXvc1EUc-LBAonILo4-gA.png)

You can see that our custom resource is there. It's cool.

您可以看到我们的自定义资源就在那里。这个很酷。

# Part 6: Debugging the Operator outside the Cluster

# 第 6 部分：在集群外调试 Operator

This is the most exciting part of the article. We can debug each line of our operator which is an immense source of knowledge for its development.

这是本文最令人兴奋的部分。我们可以调试操作符的每一行，这是其开发的巨大知识来源。

To debug the operator, first of all, the CRDs must be installed in the cluster as shown in

要调试算子，首先必须在集群中安装 CRD，如图所示

> Part 4: Installing the CRDs

> 第 4 部分：安装 CRD

Make sure that your local k8s cluster is up & running. When you install  the Go plugin in your Visual Studio Code it installs the debugger also. So, simply click the Run → Start Debugging option of the VS Code, and  the debugging configuration will be automatically done for you.

确保您的本地 k8s 集群已启动并正在运行。当您在 Visual Studio Code 中安装 Go 插件时，它也会安装调试器。因此，只需单击 VS Code 的 Run → Start Debugging 选项，就会自动为您完成调试配置。

![img](https://miro.medium.com/max/1400/1*qfS0-oAbAnnJb6wzBPWC9Q.png)

Then put breakpoints where you wish and there you go.

然后在您希望的位置放置断点，然后就可以了。

![img](https://miro.medium.com/max/1400/1*1oxvNtT-92sktOC7TtXUFw.png)

Now, open a terminal and browse to the **~/demo-operator/config/samples**

现在，打开终端并浏览到 **~/demo-operator/config/samples**

directory and deploy the CRD.

目录并部署 CRD。

```
$ kubectl apply -f apps_v1_helloapp.yaml -n test
```

When it's created the first time, it will trigger the Reconcile loop and the program will stop at the breakpoints in the Reconcile function as shown in the above diagram. Also, for any Update or Delete, the Reconcile  function will be called. You can play with it as much as you wish until  you discover the internals of k8s.

第一次创建时，会触发 Reconcile 循环，程序会在 Reconcile 函数中的断点处停止，如上图所示。此外，对于任何更新或删除，将调用协调函数。在您发现 k8s 的内部结构之前，您可以随心所欲地使用它。

# Part7: Running the Operator inside the Cluster

# Part7：在集群内运行 Operator

A custom operator is nothing but a bunch of configuration files (YAML)  and a docker image of the operator itself. The minimal configuration  files we will need for our custom operator are:

自定义运算符只不过是一堆配置文件 (YAML) 和运算符本身的 docker 映像。我们自定义操作符所需的最少配置文件是：

1. A config file to create a Namespace for the operator.
2. A config file to create a Service Account.
3. A config file to define the Roles that our operator needs to invoke the APIs of the k8s server.
4. A config file to bind the Roles to the service account defined in step 2.
5. Finally, a Deployment config file to deploy the operator itself.

1. 用于为操作员创建命名空间的配置文件。
2. 用于创建服务帐户的配置文件。
3. 一个配置文件，用于定义我们的操作员调用 k8s 服务器的 API 所需的角色。
4. 将角色绑定到步骤 2 中定义的服务帐户的配置文件。
5. 最后，一个部署配置文件来部署操作员本身。

## Generating Roles for the Operator

## 为操作员生成角色

![img](https://miro.medium.com/max/864/1*0rX2UaGaGv0wpUquPxmKkQ.png)

Under the folder **~/demo-operator/config/rbac,** you can see there are lots of config files. But, we won't need all of them right now. But what roles we need?

在**~/demo-operator/config/rbac文件夹下，可以看到很多配置文件。但是，我们现在不需要所有这些。但是我们需要什么角色呢？

Let's dig into our controller [code](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/controllers/helloapp_controller.go).

让我们深入研究我们的控制器 [代码](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/controllers/helloapp_controller.go)。

![img](https://miro.medium.com/max/1400/1*hV0NcelmEvG7T40MIgX2Sg.png)

As shown in the above diagram, you can see that,

如上图所示，你可以看到，

1. At this point, we are doing a Get operation to check if there is a HelloApp resource present. So, we need permission to do the **Get** operation on **resource** HelloApp that belongs to the API group **apps.anupam.com**

1. 此时我们在做Get操作，检查是否有HelloApp资源存在。因此，我们需要在属于 API 组 **apps.anupam.com** 的 **resource** HelloApp 上执行 **Get** 操作的权限

2. Similarly, at this point we are doing a **Get** operation to check if there is a **Deployment** resource. Note that the Deployment resource of k8s belongs to the **apps** API group.

2. 同样的，此时我们正在做一个**Get**操作来检查是否有**Deployment**资源。注意 k8s 的 Deployment 资源属于 **apps** API 组。

3. At this point we are doing an **Update** operation on resource **Deployment** that belongs to the **apps** API group.

3. 此时，我们正在对属于 **apps** API 组的资源 **Deployment** 执行 **Update** 操作。

Now, what are all these comments that start with a + sign?

现在，所有这些以 + 号开头的注释是什么？

```
//+kubebuilder:rbac:groups=apps.anupam.com,resources=helloapps,verbs=get;list;watch;create;update;patch;delete//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete//+kubebuilder:rbac:groups=apps.anupam.com,resources=helloapps/status,verbs=get;update;patch//+kubebuilder:rbac:groups=apps.anupam.com,resources=helloapps/finalizers,verbs=update
```

They are called markers. You can read more about k8s markers in this [guide](https://book.kubebuilder.io/reference/markers.html). The [controller-gen](https://book.kubebuilder.io/reference/controller-gen.html) CLI uses these markers to generate all the artifacts (CRDs, RBAC, etc) for us.

它们被称为标记。您可以在本 [指南](https://book.kubebuilder.io/reference/markers.html) 中阅读有关 k8s 标记的更多信息。 [controller-gen](https://book.kubebuilder.io/reference/controller-gen.html) CLI 使用这些标记为我们生成所有工件（CRD、RBAC 等)。

Once these markers are added, just execute the **make manifests** command. It will generate/update all necessary stuff like CRDs, roles, role-bindings, etc.

添加这些标记后，只需执行 **make manifests** 命令。它将生成/更新所有必要的东西，如 CRD、角色、角色绑定等。

## Generating the Docker Image for the Operator

## 为 Operator 生成 Docker 镜像

The custom operator will be shipped nothing more than a docker image and all its artifacts to deploy it in a k8s cluster.

自定义操作符将只提供一个 docker 镜像及其所有工件，以将其部署到 k8s 集群中。

If you check the project structure, there is already a [Dockerfile](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/Dockerfile) generated for us by the Operator-SDK. We don't have to do anything manually.

查看项目结构，已经有Operator-SDK为我们生成的[Dockerfile](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/Dockerfile)。我们无需手动执行任何操作。

![img](https://miro.medium.com/max/1400/1*OMm10j-8Y5HGJN2ccf_BCA.png)

Now, go to the root of the project and execute the following command to build the docker image.

现在，转到项目的根目录并执行以下命令来构建 docker 镜像。

```
$ make docker-build IMG=anupamgogoi/demo-operator:latest
```

Please put your own docker repository to push the image. Once the image is  generated, push it to the Docker registry by executing the following  command.

请放置您自己的docker仓库来推送镜像。生成镜像后，通过执行以下命令将其推送到 Docker 注册表。

```
make docker-push IMG=anupamgogoi/demo-operator:latest
```

We are done.

我们完了。

## Finalize the shipment package

## 完成装运包裹

For simplicity, I will manually create a folder called [dist](https://github.com/anupamgogoi-wso2/go-apps/tree/master/demo-operator/dist) in the same project and add the 5 files as described in the very first paragraph of this section.

为简单起见，我将在同一个项目中手动创建一个名为 [dist](https://github.com/anupamgogoi-wso2/go-apps/tree/master/demo-operator/dist) 的文件夹，并将 5 个文件添加为在本节的第一段中进行了描述。

![img](https://miro.medium.com/max/934/1*pVZtGpChiz3Mr6eZuQnF7A.png)

Create the **1-Namespace.yaml** with a namespace that you prefer. The operator will be installed in this namespace. You can copy files 2, 3, and 4 from the **rbac** folder itself. Copy file 5 from the **crd/bases** folder. And finally, create the 6-Controller.yaml to deploy the operator. Just  make sure to change the namespaces in the config files. Here are the  complete [files](https://github.com/anupamgogoi-wso2/go-apps/tree/master/demo-operator/dist).

使用您喜欢的命名空间创建 **1-Namespace.yaml**。运算符将安装在此命名空间中。您可以从 **rbac** 文件夹本身复制文件 2、3 和 4。从 **crd/bases** 文件夹复制文件 5。最后，创建 6-Controller.yaml 来部署操作员。只需确保更改配置文件中的命名空间即可。这是完整的[文件](https://github.com/anupamgogoi-wso2/go-apps/tree/master/demo-operator/dist)。

![img](https://miro.medium.com/max/814/1*NJ0mwQT92XItG-0nkhqA-w.png)

We are done.

我们完了。

## Run the Operator inside the cluster

## 在集群内运行 Operator

I have a 3 node k8s cluster created using CentOS VM. To learn about creating a k8s cluster please read this [article](https://dzone.com/articles/create-a-kubernetes-cluster-with-centos). I will just download these config files to my k8s master node and deploy the operator. Let's do it.

我有一个使用 CentOS VM 创建的 3 节点 k8s 集群。要了解如何创建 k8s 集群，请阅读这篇 [文章](https://dzone.com/articles/create-a-kubernetes-cluster-with-centos)。我将这些配置文件下载到我的 k8s 主节点并部署操作员。我们开始做吧。

I have copied the **dist** directory to the **master node** of my cluster and I will simply **kubectl apply** all the config files.

我已将 **dist** 目录复制到集群的 **master node** ，我将简单地 **kubectl apply** 所有配置文件。

![img](https://miro.medium.com/max/1400/1*En0Hq-_D-Nnx0btUJrryvQ.png)

The namespace I specified to deploy the operator was **demo-operator-system.** Let's check if the operator was created in it.

我指定用来部署 operator 的命名空间是 **demo-operator-system。** 让我们检查是否在其中创建了 operator。

![img](https://miro.medium.com/max/1400/1*FDkcTUJITFFhI0IsXm9KCg.png)

Cool!. The custom operator is deployed and also the CRD. Note that the name of the CRD is **helloapps.apps.anupam.com** as specified in the config file [5-apps.anupam.com_helloapps.yaml](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/dist/5-apps.anupam.com_helloapps.yaml).

凉爽的！。部署了自定义操作符和 CRD。请注意，CRD 的名称是 **helloapps.apps.anupam.com**，如配置文件 [5-apps.anupam.com_helloapps.yaml](https://github.com/anupamgogoi-wso2/go-apps/blob/master/demo-operator/dist/5-apps.anupam.com_helloapps.yaml)。

Now we are good to create our custom resource or kind i.e the HelloApp. Let's open one more terminal of the master node to check logs of the  custom operator while in another terminal we deploy the HelloApp  resource.

现在我们可以创建我们的自定义资源或种类，即 HelloApp。让我们再打开一个主节点的终端来查看自定义算子的日志，而在另一个终端我们部署 HelloApp 资源。

![img](https://miro.medium.com/max/1400/1*1A29p9rvKEQLx7VMtw6elQ.png)

You can see in the above screenshot that as soon as the HelloApp custom  resource is deployed (lower terminal) the logs were displayed in the  first terminal (upper). This is the best way to debug the operator while deploying inside a k8s cluster.

您可以在上面的屏幕截图中看到，一旦部署了 HelloApp 自定义资源（下部终端），日志就会显示在第一个终端（上部）中。这是在 k8s 集群中部署时调试 operator 的最佳方式。

Now let's verify if the deployment for the custom resource HelloApp was created.

现在让我们验证是否创建了自定义资源 HelloApp 的部署。

```
$ kubectl get all -n test
```

Voilà!!

瞧！！

![img](https://miro.medium.com/max/1400/1*L-1W3lP5QxQ23I5AgIZpEQ.png)

The deployment is created and it, in turn, created the Pod as specified in the spec (size=1).

部署被创建，然后它按照规范（大小=1）中的指定创建了 Pod。

Let's access this application just inside the cluster. It can be done by calling the application in its Pod IP.

让我们在集群内部访问这个应用程序。这可以通过在其 Pod IP 中调用应用程序来完成。

![img](https://miro.medium.com/max/1400/1*p3YmTxFZ6aMqmuuxQvowtg.png)

That's it. We have got our response from the application.

就是这样。我们已收到申请的回复。

# Conclusion 

#  结论

In this article, I tried to explain how to create a very simple k8s custom operator from zero using the Operator-SDK. But, please feel free to  create your custom operator using any language of your choice. Operator  implementation is not limited only to Go lang. Do give it a try!

在本文中，我试图解释如何使用 Operator-SDK 从零开始创建一个非常简单的 k8s 自定义运算符。但是，请随意使用您选择的任何语言创建您的自定义运算符。运算符的实现不仅限于 Go lang。试一试吧！

Thanks for reading. 

谢谢阅读。

