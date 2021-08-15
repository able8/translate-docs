# What is an Operator

# 什么是运算符

## What is an Operator after all?

## 到底什么是运算符？

Operators are a design pattern made public in a 2016 CoreOS [blog](https://web.archive.org/web/20170129131616/https://coreos.com/blog/introducing-operators.html) post. The goal of an Operator is to put operational knowledge into  software. Previously this knowledge only resided in the minds of  administrators, various combinations of shell scripts or automation  software like Ansible. It was outside of your Kubernetes cluster and  hard to integrate. With Operators, CoreOS changed that.

运算符是 2016 年 CoreOS [博客](https://web.archive.org/web/20170129131616/https://coreos.com/blog/introducing-operators.html) 帖子中公开的一种设计模式。操作员的目标是将操作知识应用到软件中。以前，这些知识只存在于管理员、shell 脚本的各种组合或 Ansible 等自动化软件的脑海中。它在您的 Kubernetes 集群之外，很难集成。有了 Operator，CoreOS 改变了这一点。

Operators  implement and automate common Day-1 (installation, configuration, etc)  and Day-2 (re-configuration, update, backup, failover, restore, etc.)  activities in a piece of software running inside your Kubernetes  cluster, by integrating natively with Kubernetes concepts and APIs. We  call this a Kubernetes-native application. With Operators you can stop  treating an application as a collection of primitives like `Pods`, `Deployments`, `Services` or `ConfigMaps`, but instead as a single object that only exposes the knobs that make sense for the application.

操作员通过本地集成在 Kubernetes 集群内运行的软件中实施和自动化常见的第 1 天（安装、配置等）和第 2 天（重新配置、更新、备份、故障转移、恢复等）活动Kubernetes 概念和 API。我们称其为 Kubernetes 原生应用程序。使用 Operators，您可以不再将应用程序视为诸如 `Pods`、`Deployments`、`Services` 或 `ConfigMaps` 之类的原语集合，而是将其视为仅公开对应用程序有意义的旋钮的单个对象。

## How are Operators created?

## Operator 是如何创建的？

The premise of an Operator is to have it be a custom form of `Controllers`, a core concept of Kubernetes. A controller is basically a software loop that runs continuously on the Kubernetes master nodes. In these loops  the control logic looks at certain Kubernetes objects of interest. It  audits the desired state of these objects, expressed by the user,  compares that to what’s currently going on in the cluster and then does  anything in its power to reach the desired state.

Operator 的前提是让它成为 Kubernetes 的核心概念“Controllers”的自定义形式。控制器基本上是一个在 Kubernetes 主节点上持续运行的软件循环。在这些循环中，控制逻辑查看某些感兴趣的 Kubernetes 对象。它审核用户表达的这些对象的所需状态，将其与集群中当前发生的情况进行比较，然后尽其所能达到所需状态。

This declarative model is basically the way a user interacts with Kubernetes. Operators  apply this model at the level of entire applications. They are in effect application-specific controllers. This is possible with the ability to  define custom objects, called *Custom Resource Definitions* (CRD), which were introduced in Kubernetes 1.7. An Operator for a custom app would, for example, introduce a CRD called `FooBarApp`. This is basically treated like any other object in Kubernetes, e.g. a `Service`.

这种声明式模型基本上是用户与 Kubernetes 交互的方式。运营商在整个应用程序级别应用此模型。它们实际上是特定于应用程序的控制器。这可以通过定义自定义对象（称为*自定义资源定义*（CRD））的能力实现，这些对象是在 Kubernetes 1.7 中引入的。例如，自定义应用程序的 Operator 将引入名为“FooBarApp”的 CRD。这基本上被视为 Kubernetes 中的任何其他对象，例如一个“服务”。

The Operator itself is a piece of software running in a Pod on the cluster, interacting with the Kubernetes API server. That’s how it gets notified about the presence or modification of `FooBarApp` objects. That’s also when it will start running its loop to ensure that the  application service is actually available and configured in the way the  user expressed in the specification of `FooBarApp` objects. This is called a reconciliation loop ([example code](https://github.com/k8s-operatorhub/operator-sdk/blob/master/testdata/go/v3/memcached-operator/controllers/memcached_controller.go#L51- L137)). The application service may in turn be implemented with more basic objects like `Pods`, `Secrets` or `PersistentVolumes`, but carefully arranged and initialized, specific to the needs of this  application. Furthermore, the Operator could possibly introduce an  object of type`FooBarAppBackup` and create backups of the app as a result.

Operator 本身是一个运行在集群上的 Pod 中的软件，与 Kubernetes API 服务器交互。这就是它如何获得有关“FooBarApp”对象的存在或修改的通知。这也是它开始运行其循环以确保应用程序服务实际可用并以用户在“FooBarApp”对象规范中表达的方式进行配置的时间。这称为协调循环（[示例代码]（https://github.com/k8s-operatorhub/operator-sdk/blob/master/testdata/go/v3/memcached-operator/controllers/memcached_controller.go#L51- L137))。应用服务可以依次使用更基本的对象（如`Pods`、`Secrets` 或`PersistentVolumes`）来实现，但经过精心安排和初始化，专门针对此应用程序的需求。此外，操作员可能会引入一个类型为“FooBarAppBackup”的对象，并因此创建应用程序的备份。

## How do I start writing an Operator?

## 我如何开始编写 Operator？

Head over to [Jump Start Using the Operator SDK](https://operatorhub.io/getting-started) to find out how to write your own Operator with Go, Ansible or with Helm charts.

前往 [开始使用 Operator SDK](https://operatorhub.io/getting-started) 了解如何使用 Go、Ansible 或 Helm 图表编写您自己的 Operator。

### List your operator on OperatorHub.io
 [Submit your operator >](https://operatorhub.io/contribute)

### 在 OperatorHub.io 上列出您的运营商
[提交您的运营商 >](https://operatorhub.io/contribute)

The Operator Framework is an open source toolkit to manage Kubernetes  native applications, called Operators, in an effective, automated, and  scalable way.
 [Jump-start with the SDK](https://operatorhub.io/getting-started) 
Operator Framework 是一个开源工具包，用于以有效、自动化和可扩展的方式管理称为 Operators 的 Kubernetes 原生应用程序。
[快速启动 SDK](https://operatorhub.io/getting-started)
