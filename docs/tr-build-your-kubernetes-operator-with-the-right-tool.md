# Build Your Kubernetes Operator with the Right Tool

# 使用正确的工具构建您的 Kubernetes Operator

[Rafal Leszko](https://hazelcast.com/blog/author/rafal-leszko/ "Posts by Rafal Leszko")
November 23, 2020

[Rafal Leszko](https://hazelcast.com/blog/author/rafal-leszko/“Rafal Leszko 的帖子”)
2020 年 11 月 23 日

You want to build a Kubernetes Operator for your software. Which tool to choose from? Operator SDK with Helm, Ansible, or Go? Or maybe start from scratch with Python, Java, or any other programming language? In this blog post, I discuss different approaches to writing Kubernetes Operators and list each solution’s pros and cons. All that to help you decide which tool is the right one for you!

您想为您的软件构建一个 Kubernetes Operator。选择哪种工具？带 Helm、Ansible 或 Go 的 Operator SDK？或者从头开始使用 Python、Java 或任何其他编程语言？在这篇博文中，我讨论了编写 Kubernetes Operator 的不同方法，并列出了每种解决方案的优缺点。所有这些都可以帮助您决定哪种工具适合您！

You can find the source code for this blog post [here](https://github.com/leszko/build-your-operator).

您可以在 [此处](https://github.com/leszko/build-your-operator) 中找到这篇博文的源代码。

![](https://hazelcast.com/wp-content/themes/hazelcast/assets/images/placeholder.jpg)

## Introduction

##  介绍

[Kubernetes Operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) is an **application** that **watches** a custom Kubernetes **resource** and performs **some operations** upon its changes.

[Kubernetes Operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) 是一个**应用程序**，它**监视**自定义 Kubernetes **资源** 并执行**一些操作**。

This definition is very generic because the operators themselves can do a great variety of things. To make it more digestible, let’s focus on one example that we will use throughout this blog post. This example will be a Hazelcast Operator used to create and scale a Hazelcast cluster. Let’s imagine that a user wants to manage the Hazelcast cluster via a custom Kubernetes resource `hazelcast.yaml`.

这个定义非常通用，因为操作符本身可以做各种各样的事情。为了使它更易于理解，让我们关注一个我们将在整个博客文章中使用的示例。此示例将是用于创建和扩展 Hazelcast 集群的 Hazelcast Operator。假设用户想要通过自定义 Kubernetes 资源“hazelcast.yaml”来管理 Hazelcast 集群。

```
apiVersion: hazelcast.my.domain/v1
kind: Hazelcast
metadata:
name: hazelcast-sample
spec:
size: 1
```

Upon applying this declarative configuration, a user wants to see a Hazelcast cluster with one Hazelcast member created.

应用此声明性配置后，用户希望查看创建了一个 Hazelcast 成员的 Hazelcast 集群。

```
$ kubectl apply -f hazelcast.yaml
```

Then, a user wants to be able to modify this resource configuration (i.e., change its size to 3), apply it again, and the Hazelcast cluster should resize automatically.

然后，用户希望能够修改此资源配置（即，将其大小更改为 3），再次应用它，Hazelcast 集群应自动调整大小。

![](https://hazelcast.com/wp-content/themes/hazelcast/assets/images/placeholder.jpg)

Hazelcast Operator is the **application** that watches the Hazelcast **resource** and interacts with Kubernetes API to **create**(or **update**) deployment with the given number of Hazelcast pod replicas.

Hazelcast Operator 是监视 Hazelcast **资源** 并与 Kubernetes API 交互以使用给定数量的 Hazelcast pod 副本**创建**（或**更新**）部署的**应用程序**。

This example is pretty simple, but in real life, a change in `hazelcast.yaml` could result in more complex operations. For example, cleaning data in the Hazelcast cluster, upgrading the Hazelcast version, sending metrics/logs to an external system, setting up WAN geo-replication, or creating some additional Kubernetes resources. There is no limit here; the point is that an operator observes your resource and performs any operation you want.

这个例子非常简单，但在现实生活中，`hazelcast.yaml` 的变化可能会导致更复杂的操作。例如，清理 Hazelcast 集群中的数据、升级 Hazelcast 版本、将指标/日志发送到外部系统、设置 WAN 异地复制或创建一些额外的 Kubernetes 资源。这里没有限制；关键是操作员会观察您的资源并执行您想要的任何操作。

## Operator Similarities

## 运算符相似性

Since an operator is simply an **application**, technically, you can write it in **any programming language or framework**. What’s more, Kubernetes exposes a REST API, so you can really use any language to watch for Kubernetes events and to interact with the Kubernetes cluster. Nevertheless, no matter what implementation method you choose, operators have some similarities in how you build, install, and use them.

由于运算符只是一个**应用程序**，因此从技术上讲，您可以用**任何编程语言或框架** 编写它。更重要的是，Kubernetes 公开了一个 REST API，因此您可以真正使用任何语言来监视 Kubernetes 事件并与 Kubernetes 集群进行交互。尽管如此，无论您选择哪种实现方法，操作员在构建、安装和使用它们的方式上都有一些相似之处。

To create, install, and use an operator, you always have to:

要创建、安装和使用运算符，您始终必须：

1. Implement the operator logic (in your preferred language/framework)
2. Dockerize your operator and push it to the Docker registry

1. 实现操作符逻辑（使用您的首选语言/框架）
2. Dockerize 你的操作符并将其推送到 Docker 注册表

```
$ docker build -t <user>/hazelcast-operator .&& docker push <user>/hazelcast-operator
```

3. Create CRD (Custom Resource Definition), which defines your custom resource

3.创建CRD（Custom Resource Definition），它定义了你的自定义资源

```
$ kubectl apply -f hazelast.crd.yaml
```

4. Create RBAC (Role and Role Binding) to allow an operator to interact with Kubernetes API

4. 创建 RBAC（角色和角色绑定）以允许操作员与 Kubernetes API 交互

```
$ kubectl apply -f role.yaml
$ kubectl apply -f role_binding.yaml
```

5. Deploy the operator into your Kubernetes cluster

5. 将操作员部署到您的 Kubernetes 集群中

```
$ kubectl apply -f operator.yaml
```

6. Create your custom resource

6. 创建您的自定义资源

```
$ kubectl apply -f hazelcast.yaml
```

Note that only the first point looks differently depending on the operator tool. All other steps are always exactly the same. That is why in the next sections we focus on implementing Hazelcast Operator using different techniques and for points 2-6, you can use the following source files: [hazelcast.crd.yaml](https://github.com/leszko/build-your-operator/blob/main/hazelcast.crd.yaml), [role.yaml](https://github.com/leszko/build-your-operator/blob/main/role.yaml), [role\ _binding.yaml](https://github.com/leszko/build-your-operator/blob/main/role_binding.yaml), [operator.yaml](https://github.com/leszko/build-your-operator/blob/main/operator.yaml), [hazelcast.yaml](https://github.com/leszko/build-your-operator/blob/main/hazelcast.yaml).

请注意，根据操作员工具的不同，只有第一点看起来有所不同。所有其他步骤始终完全相同。这就是为什么在接下来的部分我们专注于使用不同的技术实现 Hazelcast Operator 的原因，对于第 2-6 点，您可以使用以下源文件：[hazelcast.crd.yaml](https://github.com/leszko/build-your-operator/blob/main/hazelcast.crd.yaml), [role.yaml](https://github.com/leszko/build-your-operator/blob/main/role.yaml), [role\ _binding.yaml](https://github.com/leszko/build-your-operator/blob/main/role_binding.yaml), [operator.yaml](https://github.com/leszko/build-your-operator/blob/main/operator.yaml), [hazelcast.yaml](https://github.com/leszko/build-your-operator/blob/main/hazelcast.yaml)。

## Tools for building Operators

## 构建 Operator 的工具

Let’s start with the most popular tool for building operators, [Operator SDK](https://sdk.operatorframework.io/). It offers 3 approaches: Helm, Ansible, Go. Then, we’ll take a look into other operator frameworks, to close with the bare programming language implementation.

让我们从最流行的构建操作员工具开始，[Operator SDK](https://sdk.operatorframework.io/)。它提供了 3 种方法：Helm、Ansible、Go。然后，我们将研究其他运算符框架，以结束裸露的编程语言实现。

#### Operator SDK: Helm

#### 运营商 SDK：Helm

[Helm](https://helm.sh/) is a package manager for Kubernetes. It allows you to create a set of templated Kubernetes configurations, package them into a Helm chart, and then render using parameters defined in `values.yaml`. Helm is very simple to use because creating a Helm chart requires no more knowledge than defining standard Kubernetes configuration files.

[Helm](https://helm.sh/) 是 Kubernetes 的包管理器。它允许您创建一组模板化的 Kubernetes 配置，将它们打包到 Helm 图表中，然后使用在 `values.yaml` 中定义的参数进行渲染。 Helm 使用起来非常简单，因为创建 Helm 图表不需要比定义标准 Kubernetes 配置文件更多的知识。

In our example, to create a [Helm-based Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/operator-sdk-helm), we need first to [create the Hazelcast Helm Chart](https://github.com/leszko/build-your-operator/tree/main/operator-sdk-helm/chart).

在我们的示例中，要创建 [基于 Helm 的 Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/operator-sdk-helm)，我们首先需要 [创建 Hazelcast掌舵图](https://github.com/leszko/build-your-operator/tree/main/operator-sdk-helm/chart)。

```
$ helm create chart
```

Then, in `chart/templates` we can create `deployment.yaml`.

然后，在 `chart/templates` 中，我们可以创建 `deployment.yaml`。

```
apiVersion: apps/v1
kind: Deployment
metadata:
name: {{ include "hazelcast.fullname" .}}
spec:
replicas: {{ .Values.size }}
selector:
    matchLabels:
      app: hazelcast
template:
    metadata:
      labels:
        app: hazelcast
    spec:
      containers:
        - name: hazelcast
          image: "hazelcast/hazelcast:4.1"
```

Our only user parameter is the cluster size and we put it into `chart/values.yaml`.

我们唯一的用户参数是集群大小，我们将其放入 `chart/values.yaml`。

```
size: 1
```

Hazelcast Helm Chart is ready. Now we can generate an operator from it.

Hazelcast Helm Chart 已准备就绪。现在我们可以从中生成一个运算符。

```
$ operator-sdk init --plugins=helm
$ operator-sdk create api --group=hazelcast --version=v1 --helm-chart=./chart
```

That’s it! The operator is ready. If you want to build, install, and use it, execute the common steps for all operators.

就是这样！操作员准备好了。如果要构建、安装和使用它，请执行所有操作员的通用步骤。

```
$ docker build -t leszko/hazelcast-operator .&& docker push leszko/hazelcast-operator
$ make install                                 # create Hazelcast CRD
$ make deploy IMG=leszko/hazelcast-operator    # create operator RBAC and install operator deployment
$ kubectl apply -f config/samples/hazelcast_v1_hazelcast.yaml # create Hazelcast resource
```

Note that you need to change `leszko` to your Docker Hub account name and make your Docker registry public.

请注意，您需要将 `leszko` 更改为您的 Docker Hub 帐户名并公开您的 Docker 注册表。

A few **comments about the “Operator SDK: Helm”** approach:

关于“Operator SDK: Helm”** 方法的一些**评论：

- Operator implementation is **declarative** and therefore **simple; it** requires no more knowledge than the standard Kubernetes configurations
- If you**already have a Helm chart** for your software, then creating an operator requires **no work** at all
- Your operator functionality is**limited to the features available in Helm**
- All operator configuration files (hazelcast.crd.yaml, role.yaml, role\_binding.yaml, operator.yaml hazelcast.yaml) are**automatically generated**, so you don’t need to maintain them separately

- 运算符实现是**声明性**，因此**简单；它**不需要比标准 Kubernetes 配置更多的知识
- 如果您**已经为您的软件准备了 Helm 图表**，那么创建操作符**无需工作**
- 您的操作员功能**仅限于 Helm 中可用的功能**
- 所有operator配置文件（hazelcast.crd.yaml、role.yaml、role\_binding.yaml、operator.yaml hazelcast.yaml）都是**自动生成**的，不需要单独维护

#### Operator SDK: Ansible

#### 运营商 SDK：Ansible

[Ansible](https://www.ansible.com/) is a very powerful tool for IT automation. Its nature is declarative and, thanks to a number of plugins, you can write simple YAML files to perform complex DevOps tasks. Ansible has a plugin called `community.kubernetes.k8s` dedicated to interacting with Kubernetes API. What’s more, Operator SDK supports generating operators from Ansible roles.

[Ansible](https://www.ansible.com/) 是一个非常强大的 IT 自动化工具。它的性质是声明性的，并且由于有许多插件，您可以编写简单的 YAML 文件来执行复杂的 DevOps 任务。 Ansible 有一个名为“community.kubernetes.k8s”的插件，专门用于与 Kubernetes API 交互。更重要的是，Operator SDK 支持从 Ansible 角色生成操作符。

To create an [Ansible-based Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/operator-sdk-ansible), we need first to scaffold the Ansible operator project.

要创建一个[基于 Ansible 的 Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/operator-sdk-ansible)，我们首先需要搭建 Ansible operator 项目。

```
$ operator-sdk init --plugins=ansible
$ operator-sdk create api --group hazelcast --version v1 --kind Hazelcast --generate-role
```

Then, we can add the operator logic into [roles/hazelcast/tasks/main.yml](https://github.com/leszko/build-your-operator/blob/main/operator-sdk-ansible/roles/hazelcast/tasks/main.yml).

然后，我们可以将运算符逻辑添加到 [roles/hazelcast/tasks/main.yml](https://github.com/leszko/build-your-operator/blob/main/operator-sdk-ansible/roles/hazelcast/tasks/main.yml)。

```
---
- name: start hazelcast
community.kubernetes.k8s:
    definition:
      kind: Deployment
      apiVersion: apps/v1
      metadata:
        name: hazelcast
        namespace: '{{ ansible_operator_meta.namespace }}'
      spec:
        replicas: "{{size}}"
        selector:
          matchLabels:
            app: hazelcast
        template:
          metadata:
            labels:
              app: hazelcast
          spec:
            containers:
            - name: hazelcast
              image: "hazelcast/hazelcast:4.1"
```

Finally, we can add operator default parameters into [roles/hazelcast/defaults/main.yml](https://github.com/leszko/build-your-operator/blob/main/operator-sdk-ansible/roles/hazelcast/defaults/main.yml).

最后，我们可以将操作员默认参数添加到 [roles/hazelcast/defaults/main.yml](https://github.com/leszko/build-your-operator/blob/main/operator-sdk-ansible/roles/hazelcast/defaults/main.yml)。

```
---
size: 1
```

Note that all the implementation logic looks very similar to the standard Kubernetes configuration and therefore, similar to the Helm-based operator. A significant difference is, however, that now all the configuration is interpreted by the `community.kubernetes.k8s` plugin, only later passed to Kubernetes API, while the Helm configuration was a direct Kubernetes configuration.

请注意，所有实现逻辑看起来都非常类似于标准的 Kubernetes 配置，因此类似于基于 Helm 的操作符。然而，一个显着的区别是，现在所有配置都由 `community.kubernetes.k8s` 插件解释，后来才传递给 Kubernetes API，而 Helm 配置是直接的 Kubernetes 配置。

Steps to install, build, and use an Ansible-based operator are the same as for the Helm-based operator.

安装、构建和使用基于 Ansible 的 operator 的步骤与基于 Helm 的 operator 的步骤相同。

```
$ docker build -t leszko/hazelcast-operator .&& docker push leszko/hazelcast-operator
$ make install                                 # create Hazelcast CRD
$ make deploy IMG=leszko/hazelcast-operator    # create operator RBAC and install operator deployment
$ kubectl apply -f config/samples/hazelcast_v1_hazelcast.yaml # create Hazelcast resource
```

A few **comments about the “Operator SDK: Ansible”** approach:

关于“Operator SDK: Ansible”** 方法的一些**评论：

- Ansible allows implementing operator in a**declarative** form which is concise and human-readable
- Ansible operator is**similar to the pure Kubernetes configuration** but executed via the `community.kubernetes.k8s` plugin
- Ansible is a very**powerful tool** and it lets you express almost any logic you may want
- Similar to Helm-based operator, all configuration files (hazelcast.crd.yaml, role.yaml, role\_binding.yaml, operator.yaml hazelcast.yaml) are**automatically generated**, so no need for any additional maintenance

- Ansible 允许以简洁易读的**声明性**形式实现运算符
- Ansible 操作符**类似于纯 Kubernetes 配置**，但通过 `community.kubernetes.k8s` 插件执行
- Ansible 是一个非常**强大的工具**，它可以让你表达几乎任何你想要的逻辑
- 与基于 Helm 的 operator 类似，所有配置文件（hazelcast.crd.yaml、role.yaml、role\_binding.yaml、operator.yaml hazelcast.yaml）都是**自动生成**，因此无需任何额外维护

#### Operator SDK: Go

#### 运营商 SDK：去

[Go](https://golang.org/) is a general-purpose programming language, so you can technically write any operator logic you could ever imagine. What’s more, the Kubernetes environment itself is written in Go, so the Kubernetes client library (interacting with Kubernetes API) is second to none. Operator SDK (with embedded [Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder)) supports implementing operators in Go, so you get a lot of scaffolding and code generation for free.

[Go](https://golang.org/) 是一种通用编程语言，因此您可以在技术上编写您能想象到的任何运算符逻辑。更重要的是，Kubernetes 环境本身是用 Go 编写的，因此 Kubernetes 客户端库（与 Kubernetes API 交互）是首屈一指的。 Operator SDK（内嵌[Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder))支持在Go中实现Operator，免费获得大量的脚手架和代码生成。

To create a [Go-based Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/operator-sdk-go), we need first to execute a few commands to scaffold the project .

要创建一个[基于Go的Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/operator-sdk-go)，我们首先需要执行一些命令来搭建项目.

```
$ operator-sdk init --repo=github.com/leszko/hazelcast-operator
$ operator-sdk create api --version v1 --group=hazelcast --kind Hazelcast --resource=true --controller=true
```

Then, we are ready to implement the operator logic inside the function `Reconcile()` of the file [controllers/hazelcast\_controller.go](https://github.com/leszko/build-your-operator/blob/main/operator-sdk-go/controllers/hazelcast_controller.go).

然后，我们准备在文件 [controllers/hazelcast\_controller.go](https://github.com/leszko/build-your-operator/blob/main/operator-sdk-go/controllers/hazelcast_controller.go)。

```
// +kubebuilder:rbac:groups=hazelcast.my.domain,resources=hazelcasts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=hazelcast.my.domain,resources=hazelcasts/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch

func (r *HazelcastReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
    ctx := context.Background()
    log := r.Log.WithValues("hazelcast", req.NamespacedName)

    // Fetch the Hazelcast instance
    hazelcast := &hazelcastv1.Hazelcast{}
    err := r.Get(ctx, req.NamespacedName, hazelcast)
    if err != nil {
        if errors.IsNotFound(err) {
            log.Info("Hazelcast resource not found. Ignoring since object must be deleted")
            return ctrl.Result{}, nil
        }
        log.Error(err, "Failed to get Hazelcast")
        return ctrl.Result{}, err
    }

    // Check if the deployment already exists, if not create a new one
    found := &appsv1.Deployment{}
    err = r.Get(ctx, types.NamespacedName{Name: hazelcast.Name, Namespace: hazelcast.Namespace}, found)
    if err != nil && errors.IsNotFound(err) {
        // Define a new deployment
        dep := r.deploymentForHazelcast(hazelcast)
        log.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
        err = r.Create(ctx, dep)
        if err != nil {
            log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
            return ctrl.Result{}, err
        }
        return ctrl.Result{Requeue: true}, nil
    } else if err != nil {
        log.Error(err, "Failed to get Deployment")
        return ctrl.Result{}, err
    }

    // Ensure the deployment size is the same as the spec
    size := hazelcast.Spec.Size
    if *found.Spec.Replicas != size {
        found.Spec.Replicas = &size
        err = r.Update(ctx, found)
        if err != nil {
            log.Error(err, "Failed to update Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
            return ctrl.Result{}, err
        }
        return ctrl.Result{Requeue: true}, nil
    }

    // Update the Hazelcast status with the pod names
    // List the pods for this hazelcast's deployment
    podList := &corev1.PodList{}
    listOpts := []client.ListOption{
        client.InNamespace(hazelcast.Namespace),
        client.MatchingLabels(labelsForHazelcast(hazelcast.Name)),
    }
    if err = r.List(ctx, podList, listOpts...);err != nil {
        log.Error(err, "Failed to list pods", "Hazelcast.Namespace", hazelcast.Namespace, "Hazelcast.Name", hazelcast.Name)
        return ctrl.Result{}, err
    }
    podNames := getPodNames(podList.Items)

    // Update status.Nodes if needed
    if !reflect.DeepEqual(podNames, hazelcast.Status.Nodes) {
        hazelcast.Status.Nodes = podNames
        err := r.Status().Update(ctx, hazelcast)
        if err != nil {
            log.Error(err, "Failed to update Hazelcast status")
            return ctrl.Result{}, err
        }
    }

    return ctrl.Result{}, nil
}

// deploymentForHazelcast returns a hazelcast Deployment object
func (r *HazelcastReconciler) deploymentForHazelcast(m *hazelcastv1.Hazelcast) *appsv1.Deployment {
    ls := labelsForHazelcast(m.Name)
    replicas := m.Spec.Size

    dep := &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name:      m.Name,
            Namespace: m.Namespace,
        },
        Spec: appsv1.DeploymentSpec{
            Replicas: &replicas,
            Selector: &metav1.LabelSelector{
                MatchLabels: ls,
            },
            Template: corev1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: ls,
                },
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{{
                        Image: "hazelcast/hazelcast:4.1",
                        Name:  "hazelcast",
                    }},
                },
            },
        },
    }
    // Set Hazelcast instance as the owner and controller
    ctrl.SetControllerReference(m, dep, r.Scheme)
    return dep
}

// labelsForHazelcast returns the labels for selecting the resources
// belonging to the given hazelcast CR name.
func labelsForHazelcast(name string) map[string]string {
    return map[string]string{"app": "hazelcast", "hazelcast_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
    var podNames []string
    for _, pod := range pods {
        podNames = append(podNames, pod.Name)
    }
    return podNames
}
```

We also need to add the size field to the Hazelcast resource structure in [api/v1/hazelcast\_types.go](https://github.com/leszko/build-your-operator/blob/main/operator-sdk-go/api/v1/hazelcast_types.go).

我们还需要在[api/v1/hazelcast\_types.go](https://github.com/leszko/build-your-operator/blob/main/operator-sdk-go/api/v1/hazelcast_types.go)。

```
type HazelcastSpec struct {
    Size int32 `json:"size,omitempty"`
}
```

You can already see that we had to write **way more code** and that this code is **much more complex** than the previous Operator SDK solutions. That’s all because we came from the **declarative** Kubernetes configurations to the **imperative** programming language. That means that now it’s not enough to change the size in the configuration, but we need to provide the code flow (with the proper error handling). That’s definitely more difficult! On the other hand, programming language gives you the flexibility to program anything you want. The declarative Kubernetes configuration no longer limits you and your operator can perform any logic you could ever imagine.

您已经看到我们必须编写**更多的代码**，并且此代码比之前的 Operator SDK 解决方案**复杂得多**。这一切都是因为我们从 **声明性** Kubernetes 配置到**命令式**编程语言。这意味着现在仅更改配置中的大小是不够的，但我们需要提供代码流（具有适当的错误处理）。那肯定更难！另一方面，编程语言使您可以灵活地编写任何您想要的程序。声明式 Kubernetes 配置不再限制您，您的操作员可以执行您能想象到的任何逻辑。

One thing to note is the list of comments above the `Reconcile()` function. They are used by Operator SDK to generate `role.yaml` for the operator.

需要注意的一件事是`Reconcile()` 函数上方的注释列表。 Operator SDK 使用它们为操作员生成 `role.yaml`。

Steps to install, build, and use a Go-based operator are the same as for any other Operator SDK operator.

安装、构建和使用基于 Go 的运算符的步骤与任何其他 Operator SDK 运算符的步骤相同。

```
$ docker build -t leszko/hazelcast-operator .&& docker push leszko/hazelcast-operator
$ make install                                 # create Hazelcast CRD
$ make deploy IMG=leszko/hazelcast-operator    # create operator RBAC and install operator deployment
$ kubectl apply -f config/samples/hazelcast_v1_hazelcast.yaml # create Hazelcast resource
```

A few **comments about the “Operator SDK: Go”** approach:

关于“Operator SDK: Go”** 方法的一些**评论：

- You implement your operator in an**imperative** code, which requires more work and caution
- Go language is**well integrated with Kubernetes**
- Writing an operator in the real programming language as Go means**no limits on the functionality you want to implement**
- Operator SDK helps to**scaffold** the Go operator project as well as **generating boilerplate configuration files** (hazelcast.crd.yaml, role.yaml, role\_binding.yaml, operator.yaml, hazelcast.yaml)

- 你在**命令**代码中实现你的操作符，这需要更多的工作和谨慎
- Go 语言**与 Kubernetes 完美集成**
- 用真正的编程语言编写一个操作符，就像 Go 一样意味着**对你想要实现的功能没有限制**
- Operator SDK 有助于**脚手架** Go 运算符项目以及**生成样板配置文件**（hazelcast.crd.yaml、role.yaml、role\_binding.yaml、operator.yaml、hazelcast.yaml）

#### Operator Framework: KOPF

#### 运营商框架：KOPF

Operator SDK is the most popular tool for creating operators, but it’s not the only one. You can find some other interesting solutions for most programming languages, for example, [Java Operator SDK](https://github.com/java-operator-sdk/java-operator-sdk) or [Kubernetes Operator Pythonic Framework (KOPF) ](https://kopf.readthedocs.io/). The former one gained some traction because its implementation was started inside the Zalando company.

Operator SDK 是最流行的用于创建 Operator 的工具，但它不是唯一的。您可以为大多数编程语言找到一些其他有趣的解决方案，例如，[Java Operator SDK](https://github.com/java-operator-sdk/java-operator-sdk) 或 [Kubernetes Operator Pythonic Framework (KOPF) ](https://kopf.readthedocs.io/)。前者获得了一些关注，因为它的实施是在 Zalando 公司内部开始的。

To start creating the [KOPF-based Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/kopf-python), we first need to manually prepare all the boilerplate files: [Dockerfile ](https://github.com/leszko/build-your-operator/blob/main/kopf-python/Dockerfile) and [hazelcast.crd.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/hazelcast.crd.yaml), [role.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/role.yaml), [role\_binding.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/role_binding.yaml), [operator.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/operator.yaml), and [hazelcast.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/hazelcast.yaml). Then, we’re ready to write the operator logic in the [operator.py](https://github.com/leszko/build-your-operator/blob/main/kopf-python/operator.py) file.

要开始创建 [基于 KOPF 的 Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/kopf-python)，我们首先需要手动准备所有样板文件：[Dockerfile](https://github.com/leszko/build-your-operator/blob/main/kopf-python/Dockerfile) 和 [hazelcast.crd.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/hazelcast.crd.yaml), [role.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/role.yaml), [role\_binding.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/role_binding.yaml),[operator.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/role_binding.yaml) github.com/leszko/build-your-operator/blob/main/kopf-python/operator.yaml) 和 [hazelcast.yaml](https://github.com/leszko/build-your-operator/blob/main/kopf-python/hazelcast.yaml)。然后，我们准备在 [operator.py](https://github.com/leszko/build-your-operator/blob/main/kopf-python/operator.py) 文件中编写运算符逻辑。

```
import kopf
import pykube
import yaml

@kopf.on.create('hazelcast.my.domain', 'v1', 'hazelcasts')
def create_fn(spec, **kwargs):
    doc = create_deployment(spec)
    kopf.adopt(doc)

    api = pykube.HTTPClient(pykube.KubeConfig.from_env())
    deployment = pykube.Deployment(api, doc)
    deployment.create()

    api.session.close()

    return {'children': [deployment.metadata['uid']]}

@kopf.on.update('hazelcast.my.domain', 'v1', 'hazelcasts')
def update_fn(spec, **kwargs):
    api = pykube.HTTPClient(pykube.KubeConfig.from_env())
    deployment = pykube.Deployment.objects(api).get(name="hazelcast")
    deployment.replicas = spec.get('size', 1)
    deployment.update()

    api.session.close()

    return {'children': [deployment.metadata['uid']]}

def create_deployment(spec):
    return yaml.safe_load(f"""
        apiVersion: apps/v1
        kind: Deployment
        metadata:
          name: hazelcast
        spec:
          replicas: {spec.get('size', 1)}
          selector:
            matchLabels:
              app: hazelcast
          template:
            metadata:
              labels:
                app: hazelcast
            spec:
              containers:
                - name: hazelcast
                  image: "hazelcast/hazelcast:4.1"
    """)

```

Python itself is quite a concise language and, thanks to using Python decorators, the code looks short, clean, and tidy. However, similar to Go-based implementation, you need to cover each operation (create and update) separately because we write the imperative code. One difference compared to Go is that the Python Kubernetes client is not as well integrated with Kubernetes as the Go Kubernetes client. Python uses YAML Kubernetes configurations and manipulates them, while Go operates on Kubernetes structures.

Python 本身是一种非常简洁的语言，并且由于使用了 Python 装饰器，代码看起来简短、干净和整洁。但是，类似于基于 Go 的实现，您需要单独涵盖每个操作（创建和更新），因为我们编写命令式代码。与 Go 相比的一个区别是 Python Kubernetes 客户端与 Kubernetes 的集成不如 Go Kubernetes 客户端。 Python 使用 YAML Kubernetes 配置并操作它们，而 Go 在 Kubernetes 结构上运行。

Steps to install, build, and use a KOPF-based operator are similar to what we saw before.

安装、构建和使用基于 KOPF 的运算符的步骤与我们之前看到的类似。

```
$ docker build -t leszko/hazelcast-operator:kopf .&& docker push leszko/hazelcast-operator:kopf
$ kubectl apply -f hazelcast.crd.yaml # create Hazelcast CRD
$ kubectl apply -f role.yaml          # create operator RBAC
$ kubectl apply -f role_binding.yaml  # create operator RBAC
$ kubectl apply -f operator.yaml      # install operator
$ kubectl apply -f hazelcast.yaml     # create Hazelcast resource
```

A few **comments about the “Operator Framework”** approach:

关于“Operator Framework”** 方法的一些**评论：

- Operator Frameworks for specific languages are**less developed and popular** than Operator SDK
- While Operator SDK provides project**scaffolding** and **boilerplate code generation,** Operator Frameworks usually leave this work to a developer
- **Kubernetes clients** for Python, Java, or other languages are always **slightly worse** choice than the native Go Kubernetes client
- Programming in any general-purpose language like Python or Java means that there is**no limit on the functionary your operator provides**

- 特定语言的 Operator 框架比 Operator SDK 开发和流行度低**
- 虽然 Operator SDK 提供了项目**脚手架**和 **样板代码生成，** Operator Frameworks 通常将这项工作留给开发人员
- **Python、Java 或其他语言的 **Kubernetes 客户端**始终是**比原生 Go Kubernetes 客户端稍差**的选择
- 使用任何通用语言（如 Python 或 Java）编程意味着**对您的操作员提供的功能没有限制**

#### Bare Programming Language: Java

#### 裸编程语言：Java

Operators are nothing more than dockerized applications, so technically you can write them in any programming language. Such a solution means, however, that you’re on your own. Nothing helps you in generating all the boilerplate code or configurations. On the other hand, you can choose the language you already use for other projects in your enterprise, decreasing the learning curve. A blog post [Writing a Kubernetes Operator in Java](https://www.instana.com/blog/writing-a-kubernetes-operator-in-java-part-1/) describes how to implement an operator with Java, using Quarkus to increase the performance by building Docker native images. Let’s take a similar approach and create a [Java-based Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/java).

Operator 只不过是 dockerized 应用程序，因此从技术上讲，您可以用任何编程语言编写它们。然而，这样的解决方案意味着您必须靠自己。没有什么可以帮助您生成所有样板代码或配置。另一方面，您可以选择已用于企业中其他项目的语言，从而缩短学习曲线。一篇博文 [Writing a Kubernetes Operator in Java](https://www.instana.com/blog/writing-a-kubernetes-operator-in-java-part-1/) 描述了如何用 Java 实现一个运算符，使用 Quarkus 通过构建 Docker 原生镜像来提高性能。让我们采用类似的方法并创建一个 [基于 Java 的 Hazelcast Operator](https://github.com/leszko/build-your-operator/tree/main/java)。

We need first to manually prepare all the boilerplate files: [hazelcast.crd.yaml](https://github.com/leszko/build-your-operator/blob/main/java/hazelcast.crd.yaml), [role .yaml](https://github.com/leszko/build-your-operator/blob/main/java/role.yaml), [role\_binding.yaml](https://github.com/leszko/build-your-operator/blob/main/java/role_binding.yaml), [operator.yaml](https://github.com/leszko/build-your-operator/blob/main/java/operator.yaml), and [hazelcast.yaml](https://github.com/leszko/build-your-operator/blob/main/java/hazelcast.yaml). We can then scaffold a Quarkus project using quarkus-maven-plugin or just clone the [source repository for this blog post](https://github.com/leszko/build-your-operator). Finally, we can implement the operator logic in a few classes.

我们首先需要手动准备所有样板文件：[hazelcast.crd.yaml](https://github.com/leszko/build-your-operator/blob/main/java/hazelcast.crd.yaml), [role .yaml](https://github.com/leszko/build-your-operator/blob/main/java/role.yaml), [role\_binding.yaml](https://github.com/leszko/build-your-operator/blob/main/java/role_binding.yaml)、[operator.yaml](https://github.com/leszko/build-your-operator/blob/main/java/operator.yaml)和[hazelcast.yaml](https://github.com/leszko/build-your-operator/blob/main/java/hazelcast.yaml)。然后我们可以使用 quarkus-maven-plugin 来搭建 Quarkus 项目，或者只是克隆 [这篇博文的源代码库](https://github.com/leszko/build-your-operator)。最后，我们可以在几个类中实现运算符逻辑。

```
public class ClientProvider {

    @Produces
    @Singleton
    @Named("namespace")
    private String findNamespace() throws IOException {
        return new String(Files.readAllBytes(Paths.get("/var/run/secrets/kubernetes.io/serviceaccount/namespace")));
    }

    @Produces
    @Singleton
    KubernetesClient newClient(@Named("namespace") String namespace) {
        return new DefaultKubernetesClient().inNamespace(namespace);
    }

    @Produces
    @Singleton
    NonNamespaceOperation<HazelcastResource, HazelcastResourceList, HazelcastResourceDoneable, Resource<HazelcastResource, HazelcastResourceDoneable>> makeCustomResourceClient(
            KubernetesClient defaultClient, @Named("namespace") String namespace) {

        KubernetesDeserializer.registerCustomKind("hazelcast.my.domain/v1", "Hazelcast", HazelcastResource.class);

        CustomResourceDefinition crd = defaultClient
                .customResourceDefinitions()
                .list()
                .getItems()
                .stream()
                .filter(d -> "hazelcasts.hazelcast.my.domain".equals(d.getMetadata().getName()))
                .findAny()
                .orElseThrow(
                        () -> new RuntimeException(
                                "Deployment error: Custom resource definition \"hazelcasts.hazelcast.my.domain\" not found."));

        return defaultClient
                .customResources(crd, HazelcastResource.class, HazelcastResourceList.class, HazelcastResourceDoneable.class)
                .inNamespace(namespace);
    }
}

@ApplicationScoped
public class DeploymentInstaller {

    @Inject
    private KubernetesClient client;

    @Inject
    private HazelcastResourceCache cache;

    void onStartup(@Observes StartupEvent _ev) {
        new Thread(this::runWatch).start();
    }

    private void runWatch() {
        cache.listThenWatch(this::handleEvent);
    }

    private void handleEvent(Watcher.Action action, String uid) {
        try {
            HazelcastResource resource = cache.get(uid);
            if (resource == null) {
                return;
            }

            Predicate ownerRefMatches = deployments -> deployments.getMetadata().getOwnerReferences().stream()
                    .anyMatch(ownerReference -> ownerReference.getUid().equals(uid));

            List hazelcastDeployments = client.apps().deployments().list().getItems().stream()
                    .filter(ownerRefMatches)
                    .collect(toList());

            if (hazelcastDeployments.isEmpty()) {
                client.apps().deployments().create(newDeployment(resource));
            } else {
                for (Deployment deployment : hazelcastDeployments) {
                    setSize(deployment, resource);
                    client.apps().deployments().createOrReplace(deployment);
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
            System.exit(-1);
        }
    }

    private Deployment newDeployment(HazelcastResource resource) {
        Deployment deployment = client.apps().deployments().load(getClass().getResourceAsStream("/deployment.yaml")).get();
        setSize(deployment, resource);
        deployment.getMetadata().getOwnerReferences().get(0).setUid(resource.getMetadata().getUid());
        deployment.getMetadata().getOwnerReferences().get(0).setName(resource.getMetadata().getName());
        return deployment;
    }

    private void setSize(Deployment deployment, HazelcastResource resource) {
        deployment.getSpec().setReplicas(resource.getSpec().getSize());
    }
}

@ApplicationScoped
public class HazelcastResourceCache {

    private final Map<String, HazelcastResource> cache = new ConcurrentHashMap<>();

    @Inject
    private NonNamespaceOperation<HazelcastResource, HazelcastResourceList, HazelcastResourceDoneable, Resource<HazelcastResource, HazelcastResourceDoneable>> crClient;

    private Executor executor = Executors.newSingleThreadExecutor();

    public HazelcastResource get(String uid) {
        return cache.get(uid);
    }

    public void listThenWatch(BiConsumer<Watcher.Action, String> callback) {

        try {
            // list
            crClient
                    .list()
                    .getItems()
                    .forEach(resource -> {
                                cache.put(resource.getMetadata().getUid(), resource);
                                String uid = resource.getMetadata().getUid();
                                executor.execute(() -> callback.accept(Watcher.Action.ADDED, uid));
                            }
                    );

            // watch
            crClient.watch(new Watcher() {
                @Override
                public void eventReceived(Action action, HazelcastResource resource) {
                    try {
                        String uid = resource.getMetadata().getUid();
                        if (cache.containsKey(uid)) {
                            int knownResourceVersion = Integer.parseInt(cache.get(uid).getMetadata().getResourceVersion());
                            int receivedResourceVersion = Integer.parseInt(resource.getMetadata().getResourceVersion());
                            if (knownResourceVersion > receivedResourceVersion) {
                                return;
                            }
                        }
                        System.out.println("received " + action + " for resource " + resource);
                        if (action == Action.ADDED || action == Action.MODIFIED) {
                            cache.put(uid, resource);
                        } else if (action == Action.DELETED) {
                            cache.remove(uid);
                        } else {
                            System.err.println("Received unexpected " + action + " event for " + resource);
                            System.exit(-1);
                        }
                        executor.execute(() -> callback.accept(action, uid));
                    } catch (Exception e) {
                        e.printStackTrace();
                        System.exit(-1);
                    }
                }

                @Override
                public void onClose(KubernetesClientException cause) {
                    cause.printStackTrace();
                    System.exit(-1);
                }
            });
        } catch (Exception e) {
            e.printStackTrace();
            System.exit(-1);
        }
    }
}

```

Additionally, we need to add the Kubernetes configuration file `src/main/resources/deployment.yaml`, used in the Java code.

此外，我们还需要添加 Java 代码中使用的 Kubernetes 配置文件 `src/main/resources/deployment.yaml`。

```
apiVersion: apps/v1
kind: Deployment
metadata:
name: hazelcast
ownerReferences:
    - apiVersion: apps/v1
      kind: Hazelcast
      name: placeholder
      uid: placeholder
spec:
replicas: 1
selector:
    matchLabels:
      app: hazelcast
template:
    metadata:
      labels:
        app: hazelcast
    spec:
      containers:
        - name: hazelcast
          image: "hazelcast/hazelcast:4.1"
```

Apart from the code above, we need to add additional Java boilerplate classes: [HazelcastResource](https://github.com/leszko/build-your-operator/blob/main/java/src/main/java/com/hazelcast/operator/cr/HazelcastResource.java), [HazelcastResourceDoneable](https://github.com/leszko/build-your-operator/blob/main/java/src/main/java/com/hazelcast/operator/cr/HazelcastResourceDoneable.java), [HazelcastResourceList](https://github.com/leszko/build-your-operator/blob/main/java/src/main/java/com/hazelcast/operator/cr/HazelcastResourceList.java),[HazelcastResourceSpec](https://github.com/leszko/build-your-operator/blob/main/java/src/main/java/com/hazelcast/operator/cr/HazelcastResourceSpec.java). Yes… the combination of using Java and not using any Operator Framework must result in a lot of code. A lot of code to write and a lot of code to maintain. Java is verbose; Java Kubernetes client is verbose. That’s it. For details of the code above, I recommend reading [Writing a Kubernetes Operator in Java](https://www.instana.com/blog/writing-a-kubernetes-operator-in-java-part-1/). You can also check there the tweaks you need to make to build the Docker native image.

除了上面的代码，我们还需要添加额外的 Java 样板类：[HazelcastResource](https://github.com/leszko/build-your-operator/blob/main/java/src/main/java/com/hazelcast/operator/cr/HazelcastResource.java), [HazelcastResourceDoneable](https://github.com/leszko/build-your-operator/blob/main/java/src/main/java/com/hazelcast/operator/cr/HazelcastResourceDoneable.java), [HazelcastResourceList](https://github.com/leszko/build-your-operator/blob/main/java/src/main/java/com/hazelcast/operator/cr/HazelcastResourceList.java),[HazelcastResourceSpec](https://github.com/leszko/build-your-operator/blob/main/java/src/main/java/com/hazelcast/operator/cr/HazelcastResourceSpec.java)。是的……使用 Java 和不使用任何 Operator Framework 的组合必须导致大量代码。需要编写大量代码，需要维护大量代码。 Java 很冗长； Java Kubernetes 客户端很冗长。就是这样。关于上面代码的细节，我推荐阅读[Writing a Kubernetes Operator in Java](https://www.instana.com/blog/writing-a-kubernetes-operator-in-java-part-1/)。您还可以在此处查看构建 Docker 本机映像所需的调整。

We can build the application in two ways: Java Docker image or Docker native image. The second approach is better for the performance but requires a few tweaks in code, so for the purpose of this blog post, let’s just build a standard Docker image (and then install and use it),

我们可以通过两种方式构建应用程序：Java Docker 镜像或 Docker 原生镜像。第二种方法对性能更好，但需要对代码进行一些调整，因此出于本博文的目的，让我们构建一个标准的 Docker 镜像（然后安装和使用它），

```
$ mvn package
$ docker build -f src/main/docker/Dockerfile.jvm -t leszko/hazelcast-operator:java .&& docker push leszko/hazelcast-operator:java
$ kubectl apply -f hazelcast.crd.yaml # create Hazelcast CRD
$ kubectl apply -f role.yaml          # create operator RBAC
$ kubectl apply -f role_binding.yaml  # create operator RBAC
$ kubectl apply -f operator.yaml      # install operator
$ kubectl apply -f hazelcast.yaml     # create Hazelcast resource
```

A few **comments about the “Bare Programming Language”** approach:

关于“裸编程语言”** 方法的一些**评论：

- Creating an operator from scratch means**writing more code**
- There is **no limit on the logic** you want to deliver
- Before starting to write an operator in your preferred language, you can check if it provides a**good and mature Kubernetes client library**
- The only good reason to write an operator from scratch is using**a single programming language** inside your project/organization/enterprise

- 从头开始创建运算符意味着**编写更多代码**
- **您要交付的逻辑**没有限制
- 在开始用你喜欢的语言编写一个 operator 之前，你可以检查它是否提供了**好的和成熟的 Kubernetes 客户端库**
- 从头开始编写运算符的唯一理由是在您的项目/组织/企业中使用**单一的编程语言**

## Summary

##  概括

You can look at the code snippets above and decide which operator implementation is the right one for you. However, that’s just a part of the story. Let me give you an example. I’m a Java developer, so for me using pure Java is the simplest approach. Still, I would never choose Java for writing an operator. Why? Most developers do not write operators in Java. So, I’d be alone in it! Alone with bugs, alone with new features, alone with my questions on StackOverflow. And programming is a collaborative work!

您可以查看上面的代码片段并决定哪种运算符实现适合您。然而，这只是故事的一部分。让我给你举个例子。我是一名 Java 开发人员，所以对我来说使用纯 Java 是最简单的方法。尽管如此，我永远不会选择 Java 来编写运算符。为什么？大多数开发人员不会用 Java 编写运算符。所以，我会一个人在里面！单独处理错误，单独处理新功能，单独处理我在 StackOverflow 上的问题。编程是一项协作工作！

Then, what operator tool do others use? Let’s look at the data from [OperatorHub.io](https://operatorhub.io/).

那么，其他人使用什么操作员工具？我们来看看[OperatorHub.io](https://operatorhub.io/)的数据。

![](https://hazelcast.com/wp-content/themes/hazelcast/assets/images/placeholder.jpg)

Go-based operators are by far the most popular. You may find the data slightly biased because the operators published at OperatorHub are only those operators that are built and distributed for others, so you won’t find any internal operators there. But still, if you decide to develop your operator in Go, you’re in good company!

迄今为止，基于 Go 的运算符是最受欢迎的。您可能会发现数据略有偏差，因为在 OperatorHub 上发布的算子只是为他人构建和分发的算子，因此您不会在那里找到任何内部算子。但是，如果您决定在 Go 中开发您的运营商，那么您就是一个好伙伴！

So, which operator implementation is the right one for you? The choice is yours, but let me give you some hints.

那么，哪种运算符实现适合您？选择权在你，但让我给你一些提示。

- **Hint 1**: If you **already have a Helm chart** for your software and you don't need any complex [capability levels](https://operatorframework.io/operator-capabilities/) =\ > Operator SDK: Helm 

- **提示 1**：如果您 ** 已经为您的软件准备了 Helm 图表**，并且您不需要任何复杂的 [能力级别](https://operatorframework.io/operator-capabilities/) =\ > 运营商 SDK：Helm

- **Hint 2**: If you want to **create your operator quickly** and you don't need any complex [capability levels](https://operatorframework.io/operator-capabilities/) =\> Operator SDK: Helm
- **Hint 3:** If you want **complex features** or/and be flexible about any future implementations => Operator SDK: Go
- **Hint 4**: If you want to keep a **single programming language in your organization**
   - If a popular Operator Framework exists for your language or/and you want to contribute to it => Operator Framework
   - If no popular Operator Framework exists for your programming language => Bare Programming Language
- **Hint 5**: If **none of the above** =\> Operator SDK: Go

- **提示 2**：如果您想 **快速创建您的运营商** 并且您不需要任何复杂的[能力级别](https://operatorframework.io/operator-capabilities/) =\> 运营商SDK：头盔
- **提示 3：** 如果您想要 **复杂的功能** 或/并且对任何未来的实现保持灵活 => Operator SDK：Go
- **提示 4**：如果您想在组织中保留 **单一的编程语言**
  - 如果您的语言存在流行的 Operator Framework 或/并且您想为其做出贡献 => Operator Framework
  - 如果您的编程语言不存在流行的 Operator Framework => 裸编程语言
- **提示 5**：如果 **以上都不是** =\> Operator SDK: Go

#### Relevant Resources

#### 相关资源

#### See the Hazelcast Platform in Action (EMEA)

#### 查看 Hazelcast 平台在行动 (EMEA)

Nov 3, 2021 \| **10:00am Europe/London** 

2021 年 11 月 3 日 \| **上午 10:00 欧洲/伦敦**

