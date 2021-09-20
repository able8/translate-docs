# Go Operator Tutorial

# Go Operator 教程

From: https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/

An in-depth walkthough of building and running a Go-based operator.

构建和运行基于 Go 的 Operator 的深入演练。

## Prerequisites

- Go through the [installation guide][install-guide].
 - User authorized with `cluster-admin` permissions.
 - An accessible image registry for various operator images (ex. [hub.docker.com](https://hub.docker.com/signup),
 [quay.io](https://quay.io/)) and be logged in in your command line environment.
   - `example.com` is used as the registry Docker Hub namespace in these examples.
   Replace it with another value if using a different registry or namespace.
   - [Authentication and certificates][image-reg-config] if the registry is private or uses a custom CA.

## 先决条件

- 浏览[安装指南][安装指南]。
- 拥有 `cluster-admin` 权限的用户。
- 各种Operators图像的可访问图像注册表（例如 [hub.docker.com](https://hub.docker.com/signup），
  [quay.io](https://quay.io/)) ) 并在您的命令行环境中登录。
  - `example.com` 在这些示例中用作注册表 Docker Hub 命名空间。
  如果使用不同的注册表或命名空间，请将其替换为另一个值。
  - [身份验证和证书][image-reg-config] 如果注册表是私有的或使用自定义 CA。

## Overview

## 概述

We will create a sample project to let you know how it works and this sample will:
- Create a Memcached Deployment if it doesn't exist
 - Ensure that the Deployment size is the same as specified by the Memcached CR spec
 - Update the Memcached CR status using the status writer with the names of the CR's pods

我们将创建一个示例项目，让您了解它是如何工作的，这个示例将：

- 如果它不存在，则创建一个 Memcached 部署
- 确保部署大小与 Memcached CR 规范指定的相同
- 使用状态编写器和 CR 的 Pod 名称更新 Memcached CR 状态

## Create a new project

## 创建一个新项目

Use the CLI to create a new memcached-operator project:

使用 CLI 创建一个新的 memcached-operator 项目：

```sh
 mkdir -p $HOME/projects/memcached-operator
 cd $HOME/projects/memcached-operator
 # we'll use a domain of example.com
 # so all API groups will be <group>.example.com
 operator-sdk init --domain example.com --repo github.com/example/memcached-operator
```


To learn about the project directory structure, see [Kubebuilder project layout][kubebuilder_layout_doc] doc.

要了解项目目录结构，请参阅 [Kubebuilder 项目布局][kubebuilder_layout_doc] 文档。

#### A note on dependency management

#### 关于依赖管理的说明

`operator-sdk init` generates a `go.mod` file to be used with [Go modules][go_mod_wiki]. The `--repo=<path>` flag is required when creating a project outside of `$GOPATH/src`, as scaffolded files require a valid module path. Ensure you [activate module support][activate_modules] by running `export GO111MODULE=on` before using the SDK.

`operator-sdk init` 生成一个 `go.mod` 文件，用于 [Go modules][go_mod_wiki]。在 `$GOPATH/src` 之外创建项目时需要 `--repo=<path>` 标志，因为脚手架文件需要有效的模块路径。在使用 SDK 之前，通过运行 `export GO111MODULE=on` 确保你 [激活模块支持][activate_modules]。

### Manager

The main program for the operator `main.go` initializes and runs the [Manager][manager_go_doc].

See the [Kubebuilder entrypoint doc][kubebuilder_entrypoint_doc] for more details on how the manager registers the Scheme for the custom resource API definitions, and sets up and runs controllers and webhooks.


请参阅 [Kubebuilder 入口点文档][kubebuilder_entrypoint_doc]，了解有关管理器如何为自定义资源 API 定义注册 Scheme 以及设置和运行控制器和 webhooks 的更多详细信息。


The Manager can restrict the namespace that all controllers will watch for resources:

Manager 可以限制所有控制器将监视资源的命名空间：

```Go
 mgr, err := ctrl.NewManager(cfg, manager.Options{Namespace: namespace})
```


By default this will be the namespace that the operator is running in. To watch all namespaces leave the namespace option empty:

默认情况下，这将是Operators运行的名称空间。要查看所有名称空间，请将名称空间选项留空：

```Go
 mgr, err := ctrl.NewManager(cfg, manager.Options{Namespace: ""})
```


Read the [operator scope][operator_scope] documentation on how to run your operator as namespace-scoped vs cluster-scoped.

阅读 [operator scope][operator_scope] 文档，了解如何将操作符作为命名空间范围和集群范围运行。

## Create a new API and Controller

## 创建一个新的 API 和控制器

Create a new Custom Resource Definition (CRD) API with group `cache` version `v1alpha1` and Kind Memcached.
 When prompted, enter yes `y` for creating both the resource and controller.

使用组“缓存”版本“v1alpha1”和种类 Memcached 创建一个新的自定义资源定义 (CRD) API。出现提示时，输入 yes `y` 以创建资源和控制器。

```console
 $ operator-sdk create api --group cache --version v1alpha1 --kind Memcached --resource --controller
 Writing scaffold for you to edit...
 api/v1alpha1/memcached_types.go
 controllers/memcached_controller.go
 ...
```


This will scaffold the Memcached resource API at `api/v1alpha1/memcached_types.go` and the controller at `controllers/memcached_controller.go`.

这将在 `api/v1alpha1/memcached_types.go` 和 `controllers/memcached_controller.go` 上搭建 Memcached 资源 API 和控制器。

**Note:** This guide will cover the default case of a single group API. If you would like to support Multi-Group APIs see the [Single Group to Multi-Group][multigroup-kubebuilder-doc] doc.

**注意：** 本指南将涵盖单个组 API 的默认情况。如果您想支持多组 API，请参阅 [Single Group to Multi-Group][multigroup-kubebuilder-doc] 文档。

#### Understanding Kubernetes APIs

#### 了解 Kubernetes API

For an in-depth explanation of Kubernetes APIs and the group-version-kind model, check out these [kubebuilder docs][kb-doc-gkvs].

有关 Kubernetes API 和 group-version-kind 模型的深入解释，请查看这些 [kubebuilder docs][kb-doc-gkvs]。

In general, it's recommended to have one controller responsible for manage each API created for the project to
 properly follow the design goals set by [controller-runtime][controller-runtime].

一般来说，建议让一个控制器负责管理为项目创建的每个 API 正确遵循 [controller-runtime][controller-runtime] 设定的设计目标。

### Define the API

### 定义API

To begin, we will represent our API by defining the `Memcached` type, which will have a `MemcachedSpec.Size` field to set the quantity of memcached instances (CRs) to be deployed, and a `MemcachedStatus.Nodes` field to store a CR's Pod names.

首先，我们将通过定义`Memcached` 类型来表示我们的API，该类型将有一个`MemcachedSpec.Size` 字段来设置要部署的memcached 实例（CR）的数量，以及一个`MemcachedStatus.Nodes` 字段来存储CR 的 Pod 名称。

**Note** The Node field is just to illustrate an example of a Status field. In real cases, it would be recommended to use [Conditions][conditionals].

**注意** 节点字段仅用于说明状态字段的示例。在实际情况下，建议使用 [Conditions][conditionals]。

Define the API for the Memcached Custom Resource(CR) by modifying the Go type definitions at `api/v1alpha1/memcached_types.go` to have the following spec and status:

通过修改 `api/v1alpha1/memcached_types.go` 中的 Go 类型定义来定义 Memcached 自定义资源 (CR) 的 API，使其具有以下规范和状态：

```Go
 // MemcachedSpec defines the desired state of Memcached
 type MemcachedSpec struct {
 //+kubebuilder:validation:Minimum=0
 // Size is the size of the memcached deployment
 Size int32 `json:"size"`
 }

 // MemcachedStatus defines the observed state of Memcached
 type MemcachedStatus struct { 

// Nodes are the names of the memcached pods
 Nodes []string `json:"nodes"`
 }
```

Add the `+kubebuilder:subresource:status` [marker][status_marker] to add a [status subresource][status_subresource] to the CRD manifest so that the controller can update the CR status without changing the rest of the CR object:

添加 `+kubebuilder:subresource:status` [marker][status_marker] 以将 [status subresource][status_subresource] 添加到 CRD 清单，以便控制器可以更新 CR 状态，而无需更改 CR 对象的其余部分：

```Go
 // Memcached is the Schema for the memcacheds API
 //+kubebuilder:subresource:status
 type Memcached struct {
 metav1.TypeMeta   `json:",inline"`
 metav1.ObjectMeta `json:"metadata,omitempty"`

 Spec   MemcachedSpec   `json:"spec,omitempty"`
 Status MemcachedStatus `json:"status,omitempty"`
 }
```


After modifying the `*_types.go` file always run the following command to update the generated code for that resource type:

修改 `*_types.go` 文件后，始终运行以下命令来更新该资源类型的生成代码：

```sh
 make generate
```


The above makefile target will invoke the [controller-gen][controller_tools] utility to update the `api/v1alpha1/zz_generated.deepcopy.go` file to ensure our API's Go type definitons implement the `runtime.Object` interface that all Kind types must implement.

上述 makefile 目标将调用 [controller-gen][controller_tools] 实用程序来更新 `api/v1alpha1/zz_generated.deepcopy.go` 文件，以确保我们 API 的 Go 类型定义实现所有 Kind 类型的 `runtime.Object` 接口必须执行。

### Generating CRD manifests

### 生成 CRD 清单

Once the API is defined with spec/status fields and CRD validation markers, the CRD manifests can be generated and updated with the following command:

使用规范/状态字段和 CRD 验证标记定义 API 后，可以使用以下命令生成和更新 CRD 清单：

```sh
 make manifests
```


This makefile target will invoke [controller-gen][controller_tools] to generate the CRD manifests at `config/crd/bases/cache.example.com_memcacheds.yaml`.

此 makefile 目标将调用 [controller-gen][controller_tools] 以在 `config/crd/bases/cache.example.com_memcacheds.yaml` 中生成 CRD 清单。

### OpenAPI validation

### OpenAPI 验证

OpenAPI validation defined in a CRD ensures CRs are validated based on a set of declarative rules. All CRDs should have validation.  See the [OpenAPI valiation][openapi-validation] doc for details.

CRD 中定义的 OpenAPI 验证可确保根据一组声明性规则验证 CR。所有 CRD 都应该有验证。有关详细信息，请参阅 [OpenAPI 验证][openapi-validation] 文档。

## Implement the Controller

## 实现控制器

For this example replace the generated controller file `controllers/memcached_controller.go` with the example [`memcached_controller.go`][memcached_controller] implementation.

对于此示例，将生成的控制器文件 `controllers/memcached_controller.go` 替换为示例 [`memcached_controller.go`][memcached_controller] 实现。

**Note**: The next two subsections explain how the controller watches resources and how the reconcile loop is triggered.
 If you'd like to skip this section, head to the [deploy](#run-the-operator) section to see how to run the operator.

**注意**：接下来的两小节解释了控制器如何监视资源以及如何触发协调循环。如果您想跳过此部分，请前往 [deploy](#run-the-operator) 部分查看如何运行该 Operator 。

### Resources watched by the Controller

### 控制器监视的资源

The `SetupWithManager()` function in `controllers/memcached_controller.go` specifies how the controller is built to watch a CR and other resources that are owned and managed by that controller.

`controllers/memcached_controller.go` 中的 `SetupWithManager()` 函数指定了如何构建控制器来监视 CR 以及由该控制器拥有和管理的其他资源。

```Go
 import (
 ...
 appsv1 "k8s.io/api/apps/v1"
 ...
 )

 func (r *MemcachedReconciler) SetupWithManager(mgr ctrl.Manager) error {
 return ctrl.NewControllerManagedBy(mgr).
 For(&cachev1alpha1.Memcached{}).
 Owns(&appsv1.Deployment{}).
 Complete(r)
 }
```


The `NewControllerManagedBy()` provides a controller builder that allows various controller configurations.

`NewControllerManagedBy()` 提供了一个控制器构建器，允许各种控制器配置。

`For(&cachev1alpha1.Memcached{})` specifies the Memcached type as the primary resource to watch. For each Memcached type Add/Update/Delete event the reconcile loop will be sent a reconcile `Request` (a namespace/name key) for that Memcached object.

`For(&cachev1alpha1.Memcached{})` 指定 Memcached 类型作为要观察的主要资源。对于每个 Memcached 类型的 Add/Update/Delete 事件，协调循环将为该 Memcached 对象发送一个协调`Request`（命名空间/名称键）。

`Owns(&appsv1.Deployment{})` specifies the Deployments type as the secondary resource to watch. For each Deployment type Add/Update/Delete event, the event handler will map each event to a reconcile `Request` for the owner of the Deployment. Which in this case is the Memcached object for which the Deployment was created.

`Owns(&appsv1.Deployment{})` 将 Deployments 类型指定为要监视的辅助资源。对于每个部署类型的添加/更新/删除事件，事件处理程序会将每个事件映射到部署所有者的协调“请求”。在这种情况下，这是为其创建部署的 Memcached 对象。

### Controller Configurations

### 控制器配置

There are a number of other useful configurations that can be made when initialzing a controller. For more details on these configurations consult the upstream [builder][builder_godocs] and [controller][controller_godocs] godocs.

在初始化控制器时，可以进行许多其他有用的配置。有关这些配置的更多详细信息，请参阅上游 [builder][builder_godocs] 和 [controller][controller_godocs] godocs。

- Set the max number of concurrent Reconciles for the controller via the [`MaxConcurrentReconciles`][controller_options]  option. Defaults to 1.
   ```Go
   func (r *MemcachedReconciler) SetupWithManager(mgr ctrl.Manager) error {
     return ctrl.NewControllerManagedBy(mgr).
       For(&cachev1alpha1.Memcached{}).
       Owns(&appsv1.Deployment{}).
       WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
       Complete(r)
   }
   ```

- Filter watch events using [predicates][event_filtering]
 - Choose the type of [EventHandler][event_handler_godocs] to change how a watch event will translate to reconcile requests for the reconcile loop. For operator relationships that are more complex than primary and secondary resources, the [`EnqueueRequestsFromMapFunc`][enqueue_requests_from_map_func] handler can be used to transform a watch event into an arbitrary set of reconcile requests.

- 使用 [predicates][event_filtering] 过滤观察事件
- 选择 [EventHandler][event_handler_godocs] 的类型以更改监视事件将如何转换为协调循环的协调请求。对于比主要和次要资源更复杂的 Operator 关系，[`EnqueueRequestsFromMapFunc`][enqueue_requests_from_map_func] 处理程序可用于将监视事件转换为任意一组协调请求。

### Reconcile loop

### 协调循环

The reconcile function is responsible for enforcing the desired CR state on the actual state of the system. It runs each time an event occurs on a watched CR or resource, and will return some value depending on whether those states match or not. 

协调功能负责在系统的实际状态上强制执行所需的 CR 状态。每次在受监视的 CR 或资源上发生事件时，它都会运行，并根据这些状态是否匹配返回一些值。

In this way, every Controller has a Reconciler object with a `Reconcile()` method that implements the reconcile loop. The reconcile loop is passed the [`Request`][request-go-doc] argument which is a Namespace/Name key used to lookup the primary resource object, Memcached, from the cache:

这样，每个 Controller 都有一个 Reconciler 对象，该对象带有一个实现协调循环的“Reconcile()”方法。协调循环传递 [`Request`][request-go-doc] 参数，该参数是用于从缓存中查找主要资源对象 Memcached 的命名空间/名称键：

```Go
 import (
 ctrl "sigs.k8s.io/controller-runtime"

 cachev1alpha1 "github.com/example/memcached-operator/api/v1alpha1"
 ...
 )

 func (r *MemcachedReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
   // Lookup the Memcached instance for this reconcile request
   memcached := &cachev1alpha1.Memcached{}
   err := r.Get(ctx, req.NamespacedName, memcached)
   ...
 }
```


For a guide on Reconcilers, Clients, and interacting with resource Events, see the [Client API doc][doc_client_api].

有关协调器、客户端和与资源事件交互的指南，请参阅 [客户端 API 文档][doc_client_api]。

The following are a few possible return options for a Reconciler:

以下是 Reconciler 的一些可能的返回选项：

- With the error:
   ```go
   return ctrl.Result{}, err
   ```

- Without an error:
   ```go
   return ctrl.Result{Requeue: true}, nil
   ```

- Therefore, to stop the Reconcile, use:
   ```go
   return ctrl.Result{}, nil
   ```

- Reconcile again after X time:
   ```go
    return ctrl.Result{RequeueAfter: nextRun.Sub(r.Now())}, nil
   ```


For more details, check the Reconcile and its [Reconcile godoc][reconcile-godoc].

有关更多详细信息，请查看 Reconcile 及其 [Reconcile godoc][reconcile-godoc]。

### Specify permissions and generate RBAC manifests

### 指定权限并生成 RBAC 清单

The controller needs certain [RBAC][rbac-k8s-doc] permissions to interact with the resources it manages. These are specified via [RBAC markers][rbac_markers] like the following:

控制器需要某些 [RBAC][rback-k8s-doc] 权限才能与其管理的资源进行交互。这些是通过 [RBAC 标记][rbac_markers] 指定的，如下所示：

```Go
 //+kubebuilder:rbac:groups=cache.example.com,resources=memcacheds,verbs=get;list;watch;create;update;patch;delete
 //+kubebuilder:rbac:groups=cache.example.com,resources=memcacheds/status,verbs=get;update;patch
 //+kubebuilder:rbac:groups=cache.example.com,resources=memcacheds/finalizers,verbs=update
 //+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
 //+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;

 func (r *MemcachedReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
   ...
 }
```


The `ClusterRole` manifest at `config/rbac/role.yaml` is generated from the above markers via controller-gen with the following command:

`config/rbac/role.yaml` 中的 `ClusterRole` 清单是通过 controller-gen 使用以下命令从上述标记生成的：

```sh
 make manifests
```


## Configure the operator's image registry

## 配置 Operator 的镜像注册中心

All that remains is to build and push the operator image to the desired image registry.

剩下的就是构建Operators映像并将其推送到所需的映像注册表。

Before building the operator image, ensure the generated Dockerfile references  the base image you want. You can change the default "runner" image `gcr.io/distroless/static:nonroot`  by replacing its tag with another, for example `alpine:latest`, and removing  the `USER 65532:65532` directive.

在构建Operators镜像之前，确保生成的 Dockerfile 引用你想要的基本图像。您可以更改默认的“runner”图像`gcr.io/distroless/static:nonroot` 通过用另一个标签替换它的标签，例如`alpine:latest`，并删除
`USER 65532:65532` 指令。

Your Makefile composes image tags either from values written at project initialization or from the CLI.
 In particular, `IMAGE_TAG_BASE` lets you define a common image registry, namespace, and partial name
 for all your image tags. Update this to another registry and/or namespace if the current value is incorrect.
 Afterwards you can update the `IMG` variable definition like so:

您的 Makefile 由在项目初始化时写入的值或来自 CLI 的值组成图像标签。特别是，`IMAGE_TAG_BASE` 允许你定义一个通用的镜像注册表、命名空间和部分名称为您的所有图片标签。如果当前值不正确，请将其更新到另一个注册表和/或命名空间。之后你可以像这样更新`IMG`变量定义：

```diff
 -IMG ?= controller:latest
 +IMG ?= $(IMAGE_TAG_BASE):$(VERSION)
```


Once done, you do not have to set `IMG` or any other image variable in the CLI. The following command will
 build and push an operator image tagged as `example.com/memcached-operator:v0.0.1` to Docker Hub:

完成后，您不必在 CLI 中设置 `IMG` 或任何其他图像变量。以下命令将
构建一个标记为 `example.com/memcached-operator:v0.0.1` 的操作符镜像并将其推送到 Docker Hub：

```console
 make docker-build docker-push
```

## Run the Operator

## 运行 Operator 

There are three ways to run the operator:

- As Go program outside a cluster
 - As a Deployment inside a Kubernetes cluster
 - Managed by the [Operator Lifecycle Manager (OLM)][doc-olm] in [bundle][quickstart-bundle] format

操作符的运行方式有以下三种：

- 作为集群外的 Go 程序
- 作为 Kubernetes 集群内的部署
- 由 [Operator Lifecycle Manager (OLM)][doc-olm] 以 [bundle][quickstart-bundle] 格式管理

### 1. Run locally outside the cluster

### 1. 在集群外本地运行

The following steps will show how to deploy the operator on the Cluster. However, to run locally for development purposes and outside of a Cluster use the target `make install run`.

以下步骤将展示如何在集群上部署Operators。但是，为了开发目的而在本地和集群之外运行，请使用目标“make install run”。

### 2. Run as a Deployment inside the cluster

### 2. 作为集群内的部署运行

By default, a new namespace is created with name `<project-name>-system`, ex. `memcached-operator-system`, and will be used for the deployment.

Run the following to deploy the operator. This will also install the RBAC manifests from `config/rbac`.

默认情况下，会创建一个名为 `<project-name>-system` 的新命名空间，例如。 `memcached-operator-system`，将用于部署。运行以下命令以部署Operators。这也将从`config/rbac`安装RBAC清单。

```sh
 make deploy
```

Verify that the memcached-operator is up and running:

验证 memcached-operator 是否已启动并正在运行：

```console
 $ kubectl get deployment -n memcached-operator-system
 NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE
 memcached-operator-controller-manager   1/1     1            1           8m
```


### 3. Deploy your Operator with OLM

### 3. 使用 OLM 部署您的 Operator

First, install [OLM][doc-olm]:

首先，安装 [OLM][doc-olm]：

```sh
 operator-sdk olm install
```


Bundle your operator, then build and push the bundle image. The `bundle` target generates a [bundle][doc-bundle]  in the `bundle` directory containing manifests and metadata defining your operator.  `bundle-build` and `bundle-push` build and push a bundle image defined by `bundle.Dockerfile`.

捆绑您的 Operator ，然后构建并推送捆绑镜像。 `bundle` 目标生成一个 [bundle ][doc-bundle]在包含定义您的Operators的清单和元数据的 `bundle` 目录中。`bundle-build` 和 `bundle-push` 构建和推送由 `bundle.Dockerfile` 定义的包镜像。

```sh
 make bundle bundle-build bundle-push
```


Finally, run your bundle. If your bundle image is hosted in a registry that is private and/or  has a custom CA, these [configuration steps][image-reg-config] must be complete.

最后，运行你的包。如果您的捆绑映像托管在私有和/或有自定义CA，这些[配置步骤](image-reg-config) 必须完成。

```sh
 operator-sdk run bundle example.com/memcached-operator-bundle:v0.0.1
```


Check out the [docs][tutorial-bundle] for a deep dive into `operator-sdk`'s OLM integration.


查看 [docs][tutorial-bundle] 以深入了解 `operator-sdk` 的 OLM 集成。


## Create a Memcached CR

## 创建一个 Memcached CR

Update the sample Memcached CR manifest at `config/samples/cache_v1alpha1_memcached.yaml` and define the `spec` as the following:

更新 `config/samples/cache_v1alpha1_memcached.yaml` 中的示例 Memcached CR 清单，并将 `spec` 定义如下：

```YAML
 apiVersion: cache.example.com/v1alpha1
 kind: Memcached
 metadata:
   name: memcached-sample
 spec:
   size: 3
```


Create the CR:

创建 CR：

```sh
 kubectl apply -f config/samples/cache_v1alpha1_memcached.yaml
```


Ensure that the memcached operator creates the deployment for the sample CR with the correct size:

确保 memcached Operators为示例 CR 创建具有正确大小的部署：

```console
 $ kubectl get deployment
 NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE
 memcached-sample                        3/3     3            3           1m
```


Check the pods and CR status to confirm the status is updated with the memcached pod names:

检查 pod 和 CR 状态以确认使用 memcached pod 名称更新状态：

```console
 $ kubectl get pods
 NAME                                  READY     STATUS    RESTARTS   AGE
 memcached-sample-6fd7c98d8-7dqdr      1/1       Running   0          1m
 memcached-sample-6fd7c98d8-g5k7v      1/1       Running   0          1m
 memcached-sample-6fd7c98d8-m7vn7      1/1       Running   0          1m
```


```console
 $ kubectl get memcached/memcached-sample -o yaml
 apiVersion: cache.example.com/v1alpha1
 kind: Memcached
 metadata:
   clusterName: ""
   creationTimestamp: 2018-03-31T22:51:08Z
   generation: 0
   name: memcached-sample
   namespace: default
   resourceVersion: "245453"
   selfLink: /apis/cache.example.com/v1alpha1/namespaces/default/memcacheds/memcached-sample
   uid: 0026cc97-3536-11e8-bd83-0800274106a1
 spec:
   size: 3
 status:
   nodes:
   - memcached-sample-6fd7c98d8-7dqdr
   - memcached-sample-6fd7c98d8-g5k7v
   - memcached-sample-6fd7c98d8-m7vn7
```


### Update the size

### 更新大小

Update `config/samples/cache_v1alpha1_memcached.yaml` to change the `spec.size` field in the Memcached CR from 3 to 5:

更新 `config/samples/cache_v1alpha1_memcached.yaml` 以将 Memcached CR 中的 `spec.size` 字段从 3 更改为 5：

```sh
 kubectl patch memcached memcached-sample -p '{"spec":{"size": 5}}' --type=merge
```


Confirm that the operator changes the deployment size:

确认 Operator 更改部署大小：

```console
 $ kubectl get deployment
 NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE
 memcached-sample                        5/5     5            5           3m
```


### Cleanup

### 清理

Run the following to delete all deployed resources:

运行以下命令删除所有部署的资源：

```sh
 kubectl delete -f config/samples/cache_v1alpha1_memcached.yaml
 make undeploy
```

## Next steps

## 下一步

Next, check out the following:
 1. Validating and mutating [admission webhooks][create_a_webhook].
 1. Operator packaging and distribution with [OLM][olm-integration].
 1. The [advanced topics][advanced-topics] doc for more use cases and under-the-hood details. 

接下来，检查以下内容：

1. 验证和变异[admission webhooks][create_a_webhook]。
1. Operator 打包和分发[OLM][olm-integration]。
1. [高级主题][advanced-topics]文档，了解更多用例和幕后细节。

