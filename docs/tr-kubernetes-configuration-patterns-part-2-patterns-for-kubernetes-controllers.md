# Kubernetes configuration patterns, Part 2: Patterns for Kubernetes controllers

# Kubernetes 配置模式，第 2 部分：Kubernetes 控制器模式

May 5, 2021

2021 年 5 月 5 日

Senior Software Engineer

高级软件工程师

Table of contents:

目录：

This article is the second in a two-part article series on [Kubernetes](http://developers.redhat.com/topics/kubernetes) configuration patterns, which you can use to configure your Kubernetes applications and controllers. The first article [introduced patterns and antipatterns that use only Kubernetes primitives](http://developers.redhat.com/blog/2021/04/28/kubernetes-configuration-patterns-part-1-patterns-for-kubernetes-primitives/). Those simple patterns are applicable to any application. This second article describes more advanced patterns that require coding against the Kubernetes API, which is what a Kubernetes controller should use.

本文是关于 [Kubernetes](http://developers.redhat.com/topics/kubernetes) 配置模式的两部分系列文章中的第二篇，您可以使用它来配置 Kubernetes 应用程序和控制器。第一篇 [介绍仅使用Kubernetes原语的模式和反模式](http://developers.redhat.com/blog/2021/04/28/kubernetes-configuration-patterns-part-1-patterns-for-kubernetes-primitives/)。这些简单的模式适用于任何应用程序。第二篇文章描述了需要针对 Kubernetes API 进行编码的更高级模式，这是 Kubernetes 控制器应该使用的。

The patterns you will learn in this article are suitable for scenarios where the basic Kubernetes features are not enough. These patterns will help you when you can't mount a `ConfigMap` from another namespace into a `Pod`, can't reload the configuration without killing the `Pod`, and so on.

您将在本文中学习的模式适用于基本 Kubernetes 功能还不够的场景。当您无法从另一个命名空间将 `ConfigMap` 挂载到 `Pod` 中，无法在不杀死 `Pod` 的情况下重新加载配置时，这些模式将帮助你。

As in the first article, for simplicity, I've used only `Deployment` s in the example YAML files. However, the examples should work with other _PodSpecables_ (anything that describes a `PodSpec`) such as `DaemonSet` s and `ReplicaSet` s. I have also omitted fields like `image`, `imagePullPolicy`, and others in the example `Deployment` YAML.

在第一篇文章中，为了简单起见，我在示例 YAML 文件中仅使用了“部署”。但是，这些示例应该与其他 _PodSpecables_（任何描述 `PodSpec` 的东西）一起使用，例如 `DaemonSet` 和 `ReplicaSet`。我还在示例“Deployment”YAML 中省略了“image”、“imagePullPolicy”等字段。

## Configuration with central ConfigMaps

## 使用中央 ConfigMap 进行配置

It isn't possible to mount a `ConfigMap` to a `Deployment`, `Pod`, `ReplicaSet`, or other components, when they are in separate namespaces. In some cases though, you need a single central configuration store for components that run in different namespaces. The following `Deployment` illustrates the situation:

当它们位于单独的命名空间中时，无法将 `ConfigMap` 挂载到 `Deployment`、`Pod`、`ReplicaSet` 或其他组件。但在某些情况下，对于在不同命名空间中运行的组件，您需要一个单一的中央配置存储。下面的“部署”说明了这种情况：

```
apiVersion: v1
kind: ConfigMap
metadata:
    name: logcollector-config
    namespace: logcollector
data:
    buffer: "2048"
    target: "file"
---
apiVersion: apps/v1
kind: Deployment
metadata:
    name: logcollector-agent
    namespace: user-ns-1
spec:
    ...
    template:
        spec:
            containers:
            - name: server
---
apiVersion: apps/v1
kind: Deployment
metadata:
    name: logcollector-agent
    namespace: user-ns-2
spec:
    ...
    template:
        spec:
            containers:
            - name: server

```

In this example, imagine there is a central log collector system, as shown in Figure 1. A Kubernetes controller (not shown in the example) creates a separate log collector agent `Deployment` for each namespace. That is why the `ConfigMap` and `Deployment` namespaces are different.

在此示例中，假设有一个中央日志收集器系统，如图 1 所示。Kubernetes 控制器（示例中未显示）为每个命名空间创建一个单独的日志收集器代理“Deployment”。这就是 `ConfigMap` 和 `Deployment` 命名空间不同的原因。

[![In this Kubernetes controller pattern, different containers in different namespaces retrieve their configuration by reading from a centralized ConfigMap.](http://developers.redhat.com/sites/default/files/styles/article_floated/public/blog/2020/11/advanced_1.png?itok=ElnioakN)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_1.png)

2020/11/advanced_1.png?itok=ElnioakN)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_1.png)

Figure 1:

图1：

Figure 1: Different containers read from a centralized ConfigMap.

图 1：从集中式 ConfigMap 读取的不同容器。

In the log collector agents within each namespace, the agent application reads the central `ConfigMap` for configuration. Please note that the `ConfigMap` is not mounted to the container. The application needs to use the Kubernetes API to read the `ConfigMap`, as the following pseudocode shows:

在每个命名空间内的日志收集器代理中，代理应用程序读取中央`ConfigMap` 进行配置。请注意，`ConfigMap` 没有挂载到容器中。应用程序需要使用Kubernetes API来读取`ConfigMap`，如下伪代码所示：

```
package main

import ...

func main(){
if _, err := kubeclient.Get(ctx).CoreV1().ConfigMaps("logcollector").Get("logcollector-config", metav1.GetOptions{});err == nil {
    watchConfigmap("logcollector" ,"logcollector-config", func(configMap *v1.ConfigMap) {
      updateConfig(configMap)
    })
} else if apierrors.IsNotFound(err) {
      log.Fatal("Central ConfigMap 'logcollector-config' in namespace 'logcollector' does not exist")
} else {
      log.Fatal("Error reading central ConfigMap 'logcollector-config' in namespace 'logcollector'")
}
}

```

This pseudocode gets the `ConfigMap` and starts watching it. If there is any change to the [ConfigMap](http://developers.redhat.com#advanced_1), the function updates the configuration. It is not necessary to roll out a new `Deployment`. 

这个伪代码获取`ConfigMap`并开始观察它。如果 [ConfigMap](http://developers.redhat.com#advanced_1) 有任何更改，该函数会更新配置。没有必要推出新的“部署”。

As the `ConfigMap` is in another namespace, containers need additional permissions for getting, reading and watching it. System administrators should handle the role-based access control (RBAC) settings for the containers. These settings are not shown here for brevity.

由于 `ConfigMap` 在另一个命名空间中，容器需要额外的权限来获取、读取和查看它。系统管理员应处理容器的基于角色的访问控制 (RBAC) 设置。为简洁起见，此处未显示这些设置。

## Configuration with central and namespaced ConfigMaps

## 使用中央和命名空间的 ConfigMap 进行配置

The [Configuration with central ConfigMaps](http://developers.redhat.com#advanced_1) pattern lets you use a central configuration. In many cases, you also need to override some part of the configuration for each namespace.

[使用中央 ConfigMap 配置](http://developers.redhat.com#advanced_1) 模式可让您使用中央配置。在许多情况下，您还需要覆盖每个命名空间的某些部分配置。

This pattern uses a _namespaced configuration_ in addition to the central configuration. A separate watch is needed for the namespaced `ConfigMap` in the code, as shown in Figure 2.

除了中央配置之外，此模式还使用_命名空间配置_。代码中命名空间`ConfigMap` 需要一个单独的监视，如图 2 所示。

[![Alt](http://developers.redhat.com/sites/default/files/styles/article_floated/public/blog/2020/11/advanced_2.png?itok=CazTkG4O)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_2.png)

redhat.com/sites/default/files/blog/2020/11/advanced_2.png)

Figure 2:

图 2：

Figure 2: Each container takes its config from a centralized configuration and a local ConfigMap.

图 2：每个容器从集中配置和本地 ConfigMap 获取其配置。

Here is the `Deployment` that sets up the namespaces and configurations:

这是设置命名空间和配置的“部署”：

```
apiVersion: v1
kind: ConfigMap
metadata:
name: logcollector-config
namespace: logcollector
data:
buffer: "2048"
---
apiVersion: v1
kind: ConfigMap
metadata:
name: logcollector-config
namespace: user-ns-1
data:
buffer: "1024"
target: "file"
---
apiVersion: apps/v1
kind: Deployment
metadata:
labels:
    app: logcollector
name: logcollector-deployment
namespace: user-ns-1
spec:
...
template:
    spec:
      containers:
      - name: server
        env:
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace

```

For brevity, the `Deployment` in namespace `user-ns-2` is not shown in the preceding YAML. It is exactly the same as the `Deployment` in namespace `user-ns-1`.

为简洁起见，命名空间 `user-ns-2` 中的 `Deployment` 未显示在前面的 YAML 中。它与命名空间 `user-ns-1` 中的 `Deployment` 完全相同。

There are two `ConfigMap` s now. One is in the `logcollector` namespace that is considered the central configuration. The other is in the same namespace as the log collector instance in the `user-ns-1` namespace. The code for retrieving the configurations is:

现在有两个 `ConfigMap`。一个是在被认为是中央配置的`logcollector` 命名空间中。另一个与 `user-ns-1` 命名空间中的日志收集器实例位于相同的命名空间中。检索配置的代码是：

```
package main

import ...

func main(){
if namespace := os.Getenv("NAMESPACE");ns == "" {
    panic("Unable to determine application namespace")
}

var central *v1.ConfigMap
var local *v1.ConfigMap

if _, err := kubeclient.Get(ctx).CoreV1().ConfigMaps("logcollector").Get("logcollector-config", metav1.GetOptions{});err == nil {
    watchConfigmap("logcollector" ,"logcollector-config", func(configMap *v1.ConfigMap) {
      central = configMap
      config := mergeConfig(central, local)
      updateConfig(config)
    })
} else if apierrors.IsNotFound(err) {
    log.Fatal("Central ConfigMap 'logcollector-config' in namespace 'logcollector' does not exist")
} else {
    log.Fatal("Error reading central ConfigMap 'logcollector-config' in namespace 'logcollector'")
}

if _, err := kubeclient.Get(ctx).CoreV1().ConfigMaps(namespace).Get("logcollector-config", metav1.GetOptions{});err == nil {
    watchConfigmap(namespace ,"logcollector-config", func(configMap *v1.ConfigMap) {
      central = configMap
      config := mergeConfig(central, local)
      updateConfig(config)
    })
} else if apierrors.IsNotFound(err) {
    log.Infof("Local ConfigMap 'logcollector-config' in namespace '%s' does not exist", namespace)
} else {
    log.Fatalf("Error reading local ConfigMap 'logcollector-config' in namespace '%s'", namespace)
}
}

func mergeConfig(central, local *v1.ConfigMap) (...){
// merge 2 configmaps here and return the result
}

```

The Kubernetes controller application will check the central config first, then the namespaced config. The configuration in the namespaced `ConfigMap` will have precedence.

Kubernetes 控制器应用程序将首先检查中央配置，然后是命名空间配置。命名空间`ConfigMap` 中的配置将具有优先权。

## Notes about the pattern

## 关于模式的注意事项

It is okay to not have any local `ConfigMap` s. Unlike when the central `ConfigMap` is missing, the program will not stop execution if the local `ConfigMap` does not exist. 

没有任何本地 `ConfigMap` 是可以的。与缺少中央`ConfigMap` 时不同，如果本地`ConfigMap` 不存在，程序将不会停止执行。

The application needs to know the container namespace in which it is running, because the local `ConfigMap` is not mounted using the `ConfigMap` mounting and the application needs to issue a GET for the local `ConfigMap`. Therefore, the application needs to tell the Kubernetes API the namespace it is looking in for the `ConfigMap`.

应用程序需要知道它正在运行的容器命名空间，因为本地 `ConfigMap` 没有使用 `ConfigMap` 挂载，应用程序需要为本地 `ConfigMap` 发出 GET。因此，应用程序需要告诉 Kubernetes API 它正在寻找“ConfigMap”的命名空间。

It would also be possible to simply mount the namespaced `ConfigMap` into the log collector container. But in that case, you would lose the ability to reload the config without any restarts.

也可以简单地将命名空间的 `ConfigMap` 挂载到日志收集器容器中。但在这种情况下，您将无法重新加载配置而无需重新启动。

## Configuration with custom resources

## 配置自定义资源

You can use custom resources to extend the Kubernetes API. A custom resource is a very powerful concept, but too complicated to explain in depth here. A sample custom resource follows:

您可以使用自定义资源来扩展 Kubernetes API。自定义资源是一个非常强大的概念，但过于复杂，无法在此深入解释。示例自定义资源如下：

```
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
name: eventconsumers.example.com
spec:
group: example.com
    versions:
    - name: v1
...
scope: Namespaced
names:
    plural: eventconsumers
    singular: eventconsumer
    kind: EventConsumer
...

```

After applying this custom resource definition, you can create an `EventConsumer` resource on Kubernetes:

应用此自定义资源定义后，您可以在 Kubernetes 上创建一个 `EventConsumer` 资源：

```
apiVersion: example.com/v1
kind: EventConsumer
metadata:
name: consumer-1
namespace: user-ns-1
spec:
source: source1.foobar.com:443
---
apiVersion: example.com/v1
kind: EventConsumer
metadata:
name: consumer-2
namespace: user-ns-2
spec:
source: source2.foobar.com:443

```

Custom resources can be used as configurations to be read by the Kubernetes controller, as shown in Figures 3 and 4. Figure 3 shows that the controller reads the custom resources.

自定义资源可以作为Kubernetes控制器读取的配置，如图3和图4所示。图3显示了控制器读取自定义资源。

[![alt](http://developers.redhat.com/sites/default/files/styles/article_floated/public/blog/2020/11/advanced_3_a.png?itok=pxfgtNOq)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_3_a.png)

redhat.com/sites/default/files/blog/2020/11/advanced_3_a.png)

Figure 3:

图 3：

Figure 3: The controller reads each custom resource.

图 3：控制器读取每个自定义资源。

Figure 4 shows that after the controller reads and processes the custom resources, it creates an event consumer application instance per custom resource.

图 4 显示，在控制器读取和处理自定义资源后，它会为每个自定义资源创建一个事件使用者应用程序实例。

[![alt](http://developers.redhat.com/sites/default/files/styles/article_floated/public/blog/2020/11/advanced_3_b.png?itok=7OdGHTYE)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_3_b.png)

redhat.com/sites/default/files/blog/2020/11/advanced_3_b.png)

Figure 4:

图 4：

Figure 4: The controller creates event consumer application instances for each custom resource.

图 4：控制器为每个自定义资源创建事件使用者应用程序实例。

In this case, the `EventConsumer` custom resource provides a configuration for the event consumption applications that are created by the event consumer Kubernetes controller.

在这种情况下，`EventConsumer` 自定义资源为事件消费者 Kubernetes 控制器创建的事件消费应用程序提供配置。

Note the `scope: Namespaced` key and value in the custom resource definition. This tells Kubernetes that the custom resources created for this definition will live in a namespace.

请注意自定义资源定义中的 `scope: Namespaced` 键和值。这告诉 Kubernetes 为此定义创建的自定义资源将存在于命名空间中。

A strong API for fields in the config is a nice feature. With rules defined in the custom resource definition, you can use a Kubernetes validation without any hassle.

配置中字段的强大 API 是一个不错的功能。使用自定义资源定义中定义的规则，您可以轻松使用 Kubernetes 验证。

## Drawbacks to configuration with custom resources

## 使用自定义资源进行配置的缺点

A custom resource is less flexible than a `ConfigMap`. It is possible to add data in any shape in a `ConfigMap`. Also, it is not necessary to register a `ConfigMap` in Kubernetes; you just create it. You must register custom resources by creating a `CustomResourceDefinition`. The information you add in a custom resource should match the shape defined in the `CustomResourceDefinition`. As stated earlier, a strong API is often considered a best practice. Even though `ConfigMap` s are more flexible, using well-defined shapes for configuration is a good idea whenever possible.

自定义资源不如 `ConfigMap` 灵活。可以在“ConfigMap”中添加任何形状的数据。另外，也不需要在 Kubernetes 中注册一个 `ConfigMap`；您只需创建它。您必须通过创建“CustomResourceDefinition”来注册自定义资源。您在自定义资源中添加的信息应与“CustomResourceDefinition”中定义的形状相匹配。如前所述，强大的 API 通常被认为是最佳实践。尽管 ConfigMap 更灵活，但尽可能使用定义明确的形状进行配置是一个好主意。

Additionally, not everything can be specified strongly in advance. Imagine a program that can connect to multiple databases. The database clients will have different configurations and their shapes will be different. It will not be possible to simply define fields in the custom resource to pass them as a whole to the database client. It is not a good idea to create a custom resource that has a super large shape definition, because it can be in multiple shapes.

此外，并非所有内容都可以提前明确指定。想象一个可以连接到多个数据库的程序。数据库客户端会有不同的配置，它们的形状也会不同。不可能简单地在自定义资源中定义字段以将它们作为一个整体传递给数据库客户端。创建具有超大形状定义的自定义资源不是一个好主意，因为它可以是多种形状。

Additionally, the custom resource definition must be maintained carefully, and you cannot simply delete fields or update fields in an incompatible way without releasing a new version of the custom resource. There is no problem with adding new fields.

此外，必须谨慎维护自定义资源定义，您不能在不发布新版本自定义资源的情况下以不兼容的方式简单地删除字段或更新字段。添加新字段没有问题。

## Configuration with custom resources, falling back to ConfigMaps 

## 使用自定义资源配置，回退到 ConfigMaps

This pattern is a mix of the [Configuration with custom resources](http://developers.redhat.com#advanced_3) pattern and [Configuration with central and namespaced ConfigMaps](http://developers.redhat.com#advanced_2) patterns .

此模式是 [具有自定义资源的配置](http://developers.redhat.com#advanced_3) 模式和 [具有中央和命名空间的 ConfigMaps 的配置](http://developers.redhat.com#advanced_2) 模式的混合.

In this pattern, configuration options that cannot be used in every scenario are placed in custom resources. Options that could be held in common by different custom resources are placed in `ConfigMap` s so that they can be shared. This pattern is shown in Figures 5 and 6. Figure 5 shows that the controller is watching and reading the user-namespaced `ConfigMap` and the custom resources. However, it relies first on the custom resources, falling back to `ConfigMap` for values missing in the custom resources.

在此模式中，无法在每个场景中使用的配置选项都放置在自定义资源中。可以由不同自定义资源共同持有的选项放置在 `ConfigMap` 中，以便它们可以共享。这种模式如图 5 和图 6 所示。图 5 显示控制器正在监视和读取用户命名空间的“ConfigMap”和自定义资源。但是，它首先依赖于自定义资源，而对于自定义资源中缺失的值，则回退到“ConfigMap”。

[![alt](http://developers.redhat.com/sites/default/files/styles/article_floated/public/blog/2020/11/advanced_4_a.png?itok=oD14StZ_)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_4_a.png)

redhat.com/sites/default/files/blog/2020/11/advanced_4_a.png)

Figure 5:

图 5：

Figure 5: The controller creates the application landscape based on custom resources and ConfigMaps.

图 5：控制器基于自定义资源和 ConfigMap 创建应用程序环境。

Figure 6 shows the application landscape created by the controller for the configuration given in Figure 5.

图 6 显示了控制器为图 5 中给出的配置创建的应用程序环境。

[![alt](http://developers.redhat.com/sites/default/files/styles/article_floated/public/blog/2020/11/advanced_4_b.png?itok=Ax6jgcYC)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_4_b.png)

redhat.com/sites/default/files/blog/2020/11/advanced_4_b.png)

Figure 6:

图 6：

Figure 6: The configuration relies first on custom resources, falling back to ConfigMaps.

图 6：配置首先依赖于自定义资源，然后回退到 ConfigMaps。

Here is the `Deployment` that sets up the configurations, starting with the custom resource definition:

这是设置配置的“部署”，从自定义资源定义开始：

```
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
name: eventconsumers.example.com
spec:
group: example.com
versions:
  - name: v1
    ...
scope: Namespaced
names:
    plural: eventconsumers
    singular: eventconsumer
    kind: EventConsumer
    ...
---
apiVersion: example.com/v1
kind: EventConsumer
metadata:
name: consumer-1
namespace: user-ns-1
spec:
source: source1.foobar.com:443
---
apiVersion: example.com/v1
kind: EventConsumer
metadata:
name: consumer-2
namespace: user-ns-2
spec:
source: source2.foobar.com:443
privateKey: |
    -----BEGIN PRIVATE KEY-----
    foobarbazfoobar
    barfoobarbazfoo
    bazfoobarfoobar
    -----END PRIVATE KEY-----

```

And here is the `ConfigMap`:

这是`ConfigMap`：

```
apiVersion: v1
kind: ConfigMap
metadata:
name: eventconsumer-config
namespace: eventconsumer
data:
buffer: "2048"
privateKey: |
    -----BEGIN PRIVATE KEY-----
    MIICdQIBADANBgk
    nC45zqIvd1QXloq
    bokumO0HhjqI12a
    -----END PRIVATE KEY-----
---
apiVersion: v1
kind: ConfigMap
metadata:
name: eventconsumer-config
namespace: user-ns-1
data:
buffer: "4096"

```

In the example, `EventConsumer` `consumer-1` connects to  the `source1.foobar.com:8443` URL to consume events. Note, though, that `consumer-2` is connecting to another URL.

在示例中，`EventConsumer` `consumer-1` 连接到 `source1.foobar.com:8443` URL 以使用事件。但是请注意，`consumer-2` 正在连接到另一个 URL。

The central `ConfigMap` in `eventconsumer-config` contains the buffer size as well as the private key. These will be used by the event consumer controllers unless there are overrides in local `ConfigMap` s or in the custom resources.

`eventconsumer-config` 中的中央 `ConfigMap` 包含缓冲区大小和私钥。除非在本地 `ConfigMap` s 或自定义资源中有覆盖，否则这些将由事件使用者控制器使用。

In the `user-ns-1` namespace, a local `ConfigMap` overrides the buffer size for the `EventConsumer` in that namespace. No local `ConfigMap` exists for `consumer-2` and no buffer size override exists in the custom resource, so the default buffer size in the central `ConfigMap` will be used for that consumer.

在`user-ns-1` 命名空间中，本地`ConfigMap` 会覆盖该命名空间中`EventConsumer` 的缓冲区大小。 “consumer-2”不存在本地“ConfigMap”，并且自定义资源中不存在缓冲区大小覆盖，因此中央“ConfigMap”中的默认缓冲区大小将用于该消费者。

In the `consumer-2` custom resource, there’s a `privateKey` defined, so the one in the central `ConfigMap` is not used.

在 `consumer-2` 自定义资源中，定义了一个 `privateKey`，因此没有使用中央 `ConfigMap` 中的那个。

## Notes about the pattern

## 关于模式的注意事项

This pattern helps with sharing common information for custom resources. The buffer size or the private key could have been hardcoded in the application. Instead, in this case, even the defaults are configurable. 

此模式有助于共享自定义资源的通用信息。缓冲区大小或私钥可能已在应用程序中硬编码。相反，在这种情况下，即使是默认值也是可配置的。

As explained in the [Configuration with central and namespaced ConfigMaps](http://developers.redhat.com#advanced_2) pattern, a better approach could be creating a separate central configuration custom resource that is cluster-scoped and that contains the shared configuration . That might then additionally bring the namespaced configuration custom resources, and so on. If there are many options and if they are complex to build and validate, using a separate custom resource might be better than using a central configuration custom resource, because it can leverage the OpenAPI validation in Kubernetes.

如 [使用中央和命名空间 ConfigMap 配置](http://developers.redhat.com#advanced_2) 模式中所述，更好的方法可能是创建一个单独的中央配置自定义资源，该资源是集群范围的并包含共享配置.这可能会额外带来命名空间配置自定义资源，等等。如果有很多选项并且构建和验证起来很复杂，那么使用单独的自定义资源可能比使用中央配置自定义资源更好，因为它可以利用 Kubernetes 中的 OpenAPI 验证。

## References to ConfigMaps in custom resources

## 对自定义资源中的 ConfigMap 的引用

The previous pattern, [Configuration with custom resources, falling back to ConfigMaps](http://developers.redhat.com#advanced_4), helps with sharing the configuration for multiple custom resources. But it does not allow you to specify different shared configurations for custom resources in the same namespace.

之前的模式[使用自定义资源配置，回退到 ConfigMaps](http://developers.redhat.com#advanced_4)，有助于共享多个自定义资源的配置。但是它不允许您为同一命名空间中的自定义资源指定不同的共享配置。

We can remedy that issue with this pattern, which keeps the long and complicated configurations in the `ConfigMap` s, as shown in Figures 7 and 8. Storing the configurations in a custom resource specification is not desirable.

我们可以使用这种模式来解决这个问题，它将长而复杂的配置保留在 `ConfigMap` 中，如图 7 和图 8 所示。将配置存储在自定义资源规范中是不可取的。

[![alt](http://developers.redhat.com/sites/default/files/styles/article_floated/public/blog/2020/11/advanced_5_a.png?itok=6UI9aVFQ)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_5_a.png)

redhat.com/sites/default/files/blog/2020/11/advanced_5_a.png)

Figure 7:

图 7：

Figure 7: Custom resources refer to a ConfigMap for part of their configuration.

图 7：自定义资源引用 ConfigMap 以获取其部分配置。

Figure 8 shows the application landscape created by the controller for the configuration given in the above figure.

图 8 显示了控制器为上图中给出的配置创建的应用程序环境。

[![alt](http://developers.redhat.com/sites/default/files/styles/article_floated/public/blog/2020/11/advanced_5_b.png?itok=xlfnbzrb)](https://developers.redhat.com/sites/default/files/blog/2020/11/advanced_5_b.png)

redhat.com/sites/default/files/blog/2020/11/advanced_5_b.png)

Figure 8:

图 8：

Figure 8: The controller uses custom resources referring to a ConfigMap to create the application landscape.

图 8：控制器使用引用 ConfigMap 的自定义资源来创建应用程序环境。

Here is the `Deployment` that sets up the configurations, starting with the custom resource definition:

这是设置配置的“部署”，从自定义资源定义开始：

```
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
name: eventconsumers.example.com
spec:
group: example.com
versions:
  - name: v1
    ...
scope: Namespaced
    names:
    plural: eventconsumers
    singular: eventconsumer
    kind: EventConsumer
    ...
---
apiVersion: example.com/v1
kind: EventConsumer
metadata:
name: consumer-1-a
spec:
source: source1a.foobar.com:443
connectionConfig:
ref: throttled
---
apiVersion: example.com/v1
kind: EventConsumer
metadata:
name: consumer-1-b
spec:
source: source1b.foobar.com:443
connectionConfig:
ref: throttled
---
apiVersion: example.com/v1
kind: EventConsumer
metadata:
name: consumer-2
spec:
source: source2.foobar.com:443
connectionConfig:
ref: unlimited

```

And here is the `ConfigMap`:

这是`ConfigMap`：

```
apiVersion: v1
kind: ConfigMap
metadata:
name: throttled
data:
lotsOfConfig: here
---
apiVersion: v1
kind: ConfigMap
metadata:
name: unlimited
data:
otherTypeOfConfig: here

```

In this pattern, the custom resource contains a reference to the `ConfigMap`. The `ConfigMap` can be referenced in multiple custom resources in the same namespace. It is also possible to reference different `ConfigMap` s in custom resources that are in the same namespace.

在此模式中，自定义资源包含对“ConfigMap”的引用。 `ConfigMap` 可以在同一命名空间中的多个自定义资源中引用。也可以在同一命名空间中的自定义资源中引用不同的 `ConfigMap`。

Furthermore, the configuration in the `ConfigMap` can even be referenced in custom resources of different types. For example, the TLS settings could be kept in a `ConfigMap` and referenced in both a Kafka client component custom resource and a CouchDB client component custom resource.

此外，`ConfigMap` 中的配置甚至可以在不同类型的自定义资源中引用。例如，TLS 设置可以保存在“ConfigMap”中，并在 Kafka 客户端组件自定义资源和 CouchDB 客户端组件自定义资源中引用。

## Conclusion 

##  结论

Patterns in the [first article](http://developers.redhat.com/blog/2021/04/28/kubernetes-configuration-patterns-part-1-patterns-for-kubernetes-primitives/) and this one complete this article series on Kubernetes configuration patterns. I personally have seen more configuration patterns used in Kubernetes controllers in the projects I work on, sometimes even very strange ones. The handpicked patterns in this article are the best I've found. They are flexible enough to let you build on top of them to solve your configuration problems.

[第一篇文章](http://developers.redhat.com/blog/2021/04/28/kubernetes-configuration-patterns-part-1-patterns-for-kubernetes-primitives/)中的模式和这个完成了这个关于 Kubernetes 配置模式的文章系列。我个人在我从事的项目中看到了更多 Kubernetes 控制器中使用的配置模式，有时甚至是非常奇怪的配置模式。本文中精心挑选的模式是我发现的最好的模式。它们足够灵活，可以让您在它们之上构建以解决您的配置问题。

_Last updated:
May 4, 2021_ 

_最近更新时间：
2021 年 5 月 4 日_

