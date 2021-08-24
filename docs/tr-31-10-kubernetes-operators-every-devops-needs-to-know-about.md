# 10 Kubernetes Operators Every DevOps needs to know about

# 10 个 Kubernetes Operators 每个 DevOps 都需要了解

Kubernetes Operators are built for specific applications and make it easier to create, configure, manage and operate those applications on Kubernetes. In this blog post we dig into the mechanics of Kubernetes operators and outline 10 operators that every DevOps needs to know about.

Kubernetes Operator 是为特定应用程序构建的，可以更轻松地在 Kubernetes 上创建、配置、管理和操作这些应用程序。在这篇博文中，我们深入研究了 Kubernetes 操作符的机制，并概述了每个 DevOps 都需要了解的 10 个操作符。

February 17, 2020

Kubernetes and containers have made software applications more [portable, scalable and helped improve resource utilisation](https://www.replex.io/blog/virtualization-vs-containers-scalability-portability-and-resource-utilization). For DevOps however Kubernetes has a much broader appeal: the ability to configure, manage and operate containerised microservices at scale. Kubernetes allows them to bake in a degree of automation into the creation, deployment, scaling and configuration of these applications that significantly reduces the management overhead and the probability of mistakes happening.

Kubernetes 和容器使软件应用程序更加[便携、可扩展并有助于提高资源利用率](https://www.replex.io/blog/virtualization-vs-containers-scalability-portability-and-resource-utilization)。然而，对于 DevOps 而言，Kubernetes 具有更广泛的吸引力：大规模配置、管理和操作容器化微服务的能力。 Kubernetes 允许他们在这些应用程序的创建、部署、扩展和配置中实现一定程度的自动化，从而显着降低管理开销和发生错误的可能性。

There is one caveat to all this however: It doesn’t work so well for [stateful applications](https://whatis.techtarget.com/definition/stateful-app). Deploying, scaling, operating and configuring stateful applications and building in automation requires a lot more input from DevOps in the shape of application specific domain knowledge.

然而，所有这一切都有一个警告：它对于[有状态应用程序](https://whatis.techtarget.com/definition/stateful-app) 效果不佳。部署、扩展、操作和配置有状态应用程序以及构建自动化需要 DevOps 以应用程序特定领域知识的形式提供更多输入。

## What are Kubernetes Operators?

## 什么是 Kubernetes Operator？

Enter Kubernetes operators. Operators are built for specific applications that make it easier to create, configure and manage those applications on Kubernetes. Most operators also extend across the entire application lifecycle making it easier to perform operational tasks like scaling, upgrading, backup and recovery of complex stateful applications. Since they use and extend the Kubernetes API, they are tightly integrated in the Kubernetes framework.

输入 Kubernetes 运算符。 Operator 是为特定应用程序构建的，可以更轻松地在 Kubernetes 上创建、配置和管理这些应用程序。大多数运维人员还扩展了整个应用程序生命周期，从而更轻松地执行复杂的有状态应用程序的扩展、升级、备份和恢复等操作任务。由于它们使用并扩展了 Kubernetes API，因此它们紧密集成在 Kubernetes 框架中。

Here is a list of some of the most common functions that [Kubernetes operators](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) perform:

- Install applications with the required configurations and number of application instances.
- Scale applications in or out
- Initiate upgrades, automated backups and failure recovery
- Perform any other administrative task that can be presented as code

以下是 [Kubernetes Operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) 执行的一些最常见功能的列表：
- 安装具有所需配置和应用程序实例数量的应用程序。
- 扩展或扩展应用程序
- 启动升级、自动备份和故障恢复
- 执行可以作为代码呈现的任何其他管理任务

## How do Kubernetes Operators work?

## Kubernetes Operator 是如何工作的？

Operators leverage the extensibility and modularity of Kubernetes to help automate administrative and operational tasks involved in creating, configuring and managing Kubernetes applications.

运维人员利用 Kubernetes 的可扩展性和模块化来帮助自动化涉及创建、配置和管理 Kubernetes 应用程序的管理和操作任务。

Operators build on the concepts of [custom Kubernetes controllers](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#custom-controllers) (CRDs) and custom resources and allow DevOps to incorporate operational knowledge into how applications are managed on Kubernetes. They act on CRDs to ensure the actual state of the cluster matches that defined in the CRDs.

Operator 基于 [自定义 Kubernetes 控制器](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#custom-controllers) (CRD) 和自定义资源的概念，并允许 DevOps将操作知识融入如何在 Kubernetes 上管理应用程序。它们作用于 CRD 以确保集群的实际状态与 CRD 中定义的相匹配。

The [prometheus](https://github.com/coreos/prometheus-operator) operator from CoreOs is a great example. It is deployed as a custom Kubernetes controller that watches the Kubernetes API for four custom resource definitions: Prometheus, ServiceMonitor, PrometheusRule and AlertManager. Once deployed the Prometheus operator installs and configures a full Prometheus stack that includes Prometheus servers, Alertmanager, Grafana, Host node\_exporter and kube-state-metrics. DevOps can then easily scale the number of individual replicas of each component, make configuration changes, update alerting rules or automatically monitor new services.

CoreOs 的 [prometheus](https://github.com/coreos/prometheus-operator) 操作符就是一个很好的例子。它部署为自定义 Kubernetes 控制器，该控制器监视 Kubernetes API 以获取四个自定义资源定义：Prometheus、ServiceMonitor、PrometheusRule 和 AlertManager。部署后，Prometheus operator 将安装并配置完整的 Prometheus 堆栈，其中包括 Prometheus 服务器、Alertmanager、Grafana、Host node\_exporter 和 kube-state-metrics。然后，DevOps 可以轻松扩展每个组件的单个副本的数量、更改配置、更新警报规则或自动监控新服务。

Now that we have covered the concept of Kubernetes operators let’s outline some useful operators that every DevOps should know about.

现在我们已经介绍了 Kubernetes 操作符的概念，让我们概述一些每个 DevOps 都应该知道的有用的操作符。

## Which Kubernetes Operators should you know about

## 您应该了解哪些 Kubernetes Operator

### RBAC Manager Operator 
### RBAC 管理器操作员
The [RBAC Manager](https://github.com/reactiveops/rbac-manager) is a Kubernetes operator from Fairwinds that aims to make RBAC on Kubernetes easier to setup, configure and manage. Kubernetes authorisation is often tedious and repetitive, requires lots of manual configuration and is hard to scale. RBAC manager significantly reduces the configuration involved in managing RBAC and creating, deleting or updating role bindings, cluster role bindings and service accounts. It serves as a single source of truth for understanding [RBAC](https://kubernetes.io/docs/reference/access-authn-authz/authorization/) state by summarising role bindings across multiple namespaces in a single RBAC Definitions file.

[RBAC Manager](https://github.com/reactiveops/rbac-manager) 是 Fairwinds 的 Kubernetes 操作员，旨在使 Kubernetes 上的 RBAC 更易于设置、配置和管理。 Kubernetes 授权通常繁琐且重复，需要大量手动配置且难以扩展。 RBAC 管理器显着减少了管理 RBAC 以及创建、删除或更新角色绑定、集群角色绑定和服务帐户所涉及的配置。它通过在单个 RBAC 定义文件中汇总跨多个命名空间的角色绑定，作为理解 [RBAC](https://kubernetes.io/docs/reference/access-authn-authz/authorization/) 状态的单一事实来源。

### MongoDB Enterprise Kubernetes Operator

### MongoDB 企业 Kubernetes Operator

The [MongoDB](https://docs.mongodb.com/manual/reference/operator/) operator helps DevOps standardize the process of creating MongoDB clusters at scale and makes it repeatable. The operator can be configured to take over typical administrative tasks involved in spinning up and managing MongoDB clusters including provisioning storage and compute, configuring network connections and setting up users. The Kubernetes operator also integrates with other MongoDB management tools like MongoDB Ops Manager and MongoDB Cloud Manager to provide backup, monitoring and performance optimisation.

[MongoDB](https://docs.mongodb.com/manual/reference/operator/) 操作符帮助 DevOps 标准化大规模创建 MongoDB 集群的过程并使其可重复。操作员可以配置为接管启动和管理 MongoDB 集群所涉及的典型管理任务，包括配置存储和计算、配置网络连接和设置用户。 Kubernetes Operator还与其他 MongoDB 管理工具（如 MongoDB Ops Manager 和 MongoDB Cloud Manager）集成，以提供备份、监控和性能优化。

### HPA Kubernetes Operator

The HPA operator from Banzai cloud is another useful operator that makes it easier to add pod autoscaling features to Helm charts. It watches for Kubernetes deployments or StatefulSets and automatically creates, deletes or updates Horizontal Pod Autoscalers (HPAs) based on annotations defined in the config. [HPA’s Github](https://github.com/banzaicloud/hpa-operator) page provides Kafka as an example. The Helm chart for Kafka does not define any HPAs for the cluster which means that deploying it will not bring up any HPAs as part of the Kafka deployment. To ensure HPAs are deployed as part of the Helm chart, DevOps can add annotations for min and maxReplicas. Once added the HPA operator will spin up the desired number of HPA replicas based on the annotations. The HPA operator also takes Prometheus based custom metrics exposed by Kube Metrics Adapter.

Banzai 云中的 HPA 算子是另一个有用的算子，它可以更轻松地向 Helm 图表添加 pod 自动缩放功能。它监视 Kubernetes 部署或 StatefulSet，并根据配置中定义的注释自动创建、删除或更新 Horizo​​ntal Pod Autoscalers (HPA)。 [HPA 的 Github](https://github.com/banzaicloud/hpa-operator) 页面以 Kafka 为例。 Kafka 的 Helm 图表没有为集群定义任何 HPA，这意味着部署它不会将任何 HPA 作为 Kafka 部署的一部分。为确保 HPA 作为 Helm 图表的一部分进行部署，DevOps 可以为 min 和 maxReplicas 添加注释。添加后，HPA 操作员将根据注释启动所需数量的 HPA 副本。 HPA 运算符还采用 Kube Metrics Adapter 公开的基于 Prometheus 的自定义指标。

### Cert-manager Kubernetes Operator

### 证书管理器 Kubernetes 操作员

Cert-manager from Jetstack is a Kubernetes operator that aims to automate the [management and issuance](https://github.com/jetstack/cert-manager) of TLS certificates. DevOps can use this operator to automate recurring tasks like ensuring certificates are valid and up to date and renewal. Once deployed Cert-manager runs as a Kubernetes deployment. DevOps can configure a list of certificates and certificate issuers as Kubernetes CRDs. Once configured certificates can be requested on the fly by referring to one of the configured issuers.

Jetstack 的 Cert-manager 是一个 Kubernetes Operator，旨在自动化 TLS 证书的[管理和颁发](https://github.com/jetstack/cert-manager)。 DevOps 可以使用此运算符自动执行重复性任务，例如确保证书有效、最新和更新。部署后，Cert-manager 作为 Kubernetes 部署运行。 DevOps 可以将证书和证书颁发者列表配置为 Kubernetes CRD。可以通过参考已配置的颁发者之一来即时请求配置的证书。

### ArgoCD Operator

### ArgoCD 操作员

The [ArgoCD operator](https://github.com/jmckind/argocd-operator) manages the complete life cycle for ArgoCD and its components. ArgoCD is one of the highest rated continuous delivery tools in the CNCF landscape and is specifically targeted towards Kubernetes. The operator makes it easy to configure and install ArgoCD, as well as making it easier to upgrade, backup, restore and scale ArgoCD components. The operator does this by watching for three Kubernetes CRDs including ArgoCD, which defines the desired state for an ArgoCD cluster and ArgoCDExport which defines the desired state for export and recovery of ArgoCD components.

[ArgoCD operator](https://github.com/jmckind/argocd-operator) 管理 ArgoCD 及其组件的完整生命周期。 ArgoCD 是 CNCF 领域中评价最高的持续交付工具之一，专门针对 Kubernetes。操作员可以轻松配置和安装 ArgoCD，也可以更轻松地升级、备份、恢复和扩展 ArgoCD 组件。操作员通过观察三个 Kubernetes CRD 来做到这一点，包括定义 ArgoCD 集群所需状态的 ArgoCD 和定义 ArgoCD 组件导出和恢复所需状态的 ArgoCDExport。

### **Istio Operator** 
[Istio](https://istio.io/) has emerged as the go-to service mesh tool to manage, orchestrate, secure and monitor communications across microservices deployed on Kubernetes. The [Istio operator](https://istio.io/blog/2019/introducing-istio-operator/) makes it easier to install, upgrade and troubleshoot Istio. Installation requires only istioctl as a prerequisite, small customizations are easier to make since they don’t require API changes, and version specific upgrade hooks can be easily implemented. Installing Istio using the operator also ensures that all API fields are validated. The operator API supports all 6 [built-in installation config profiles](https://istio.io/docs/setup/additional-setup/config-profiles/) including default, demo, minimal and remote. DevOps and SREs can start off with any one of these and make configuration changes further along to tailor the service mesh to their specific needs.

[Istio](https://istio.io/) 已成为管理、编排、保护和监控部署在 Kubernetes 上的微服务之间的通信的首选服务网格工具。 [Istio operator](https://istio.io/blog/2019/introducing-istio-operator/) 使 Istio 的安装、升级和故障排除变得更加容易。安装只需要 istioctl 作为先决条件，小的自定义更容易进行，因为它们不需要 API 更改，并且可以轻松实现特定于版本的升级挂钩。使用操作符安装 Istio 还可以确保所有 API 字段都经过验证。 operator API 支持所有 6 个 [内置安装配置文件](https://istio.io/docs/setup/additional-setup/config-profiles/)，包括默认、演示、最小和远程。 DevOps 和 SRE 可以从其中任何一个开始，然后进一步更改配置以根据他们的特定需求定制服务网格。

### **Etcd Operator**

[Etcd](https://coreos.com/blog/introducing-the-etcd-operator.html) serves as the primary data store for all cluster data on Kubernetes and as such is a critical component of each cluster. Managing and configuring etcd clusters on Kubernetes is a time-consuming task and requires hands-on expertise. Ensuring high availability, monitoring and disaster recovery add additional complexities. The [etcd operator](https://coreos.com/operators/etcd/docs/latest/) helps DevOps and SRE simplify these tasks by making it easier to create, configure and manage etcd clusters on Kubernetes. Teams can easily spin up multiple highly available etcd instances without having to specify detailed configuration settings, modify cluster spec to resize clusters, configure automated backup policies for disaster recovery, and initiate graceful upgrades without downtime.

[Etcd](https://coreos.com/blog/introducing-the-etcd-operator.html) 作为 Kubernetes 上所有集群数据的主要数据存储，因此是每个集群的关键组件。在 Kubernetes 上管理和配置 etcd 集群是一项耗时的任务，需要实践经验。确保高可用性、监控和灾难恢复会增加额外的复杂性。 [etcd operator](https://coreos.com/operators/etcd/docs/latest/) 通过更轻松地在 Kubernetes 上创建、配置和管理 etcd 集群来帮助 DevOps 和 SRE 简化这些任务。团队可以轻松启动多个高度可用的 etcd 实例，而无需指定详细的配置设置、修改集群规范以调整集群大小、配置用于灾难恢复的自动备份策略，以及在不停机的情况下启动优雅升级。

### **Elastic Cloud on Kubernetes (Elastic Kubernetes Operator)**

[Elastic cloud on Kubernetes (ECK)](https://github.com/elastic/cloud-on-k8s) is the official Kubernetes operator from elastic.co and aims to provide a seamless experience for deploying, managing and operating the entire elastic stack on Kubernetes. In addition to making it easier to deploy elastic search and Kibana on Kubernetes, it also simplifies critical operations tasks including managing and monitoring multiple clusters, initiating graceful upgrades, scaling both cluster capacity and local storage, making configuration changes, and backups. The default ECK distribution is free and open-source with built-in features including frozen indices for dense storage, Kibana Spaces, Canvas and Elastic Maps, and also supports monitoring of Kubernetes logs and infrastructure.

[Elastic cloud on Kubernetes (ECK)](https://github.com/elastic/cloud-on-k8s) 是elastic.co 的官方Kubernetes Operator，旨在为部署、管理和运营整个Kubernetes 提供无缝体验。 Kubernetes 上的弹性堆栈。除了可以更轻松地在 Kubernetes 上部署弹性搜索和 Kibana 之外，它还简化了关键操作任务，包括管理和监控多个集群、启动优雅升级、扩展集群容量和本地存储、进行配置更改和备份。默认的 ECK 发行版是免费和开源的，具有内置功能，包括用于密集存储的冻结索引、Kibana Spaces、Canvas 和 Elastic Maps，还支持监控 Kubernetes 日志和基础设施。

### **Grafana Operator**

The [Grafana operator](https://github.com/integr8ly/grafana-operator), offered by RedHat, simplifies the process of creating, configuring and managing Grafana instances on Kubernetes. In addition to helping deploy Grafana it also supports making Grafana available via ingress, automated dashboard and data source discovery, and installation of dashboard dependencies. The operator can be [installed](https://github.com/integr8ly/grafana-operator/blob/master/documentation/deploy_grafana.md) using either Ansible or manually via kubectl commands and creating a custom resource. Once installed the operator watches for dashboard definitions in either its own namespace or all namespaces, depending on the flag passed during deployment, discovers dashboards, adds error messages to the status field of the dashboard in case of any invalid json, and automatically [installs](https://github.com/integr8ly/grafana-operator/blob/v3.2.0/documentation/dashboards.md#plugins) any plugins specified. DevOps and SREs can then add data sources in the GrafanaDataSource as well as add extra configuration files.

RedHat 提供的 [Grafana operator](https://github.com/integr8ly/grafana-operator) 简化了在 Kubernetes 上创建、配置和管理 Grafana 实例的过程。除了帮助部署 Grafana 之外，它还支持通过入口、自动化仪表板和数据源发现以及仪表板依赖项的安装使 Grafana 可用。可以使用 Ansible 或手动通过 kubectl 命令[安装](https://github.com/integr8ly/grafana-operator/blob/master/documentation/deploy_grafana.md) 并创建自定义资源。安装后，操作员会在其自己的命名空间或所有命名空间中监视仪表板定义，具体取决于部署期间传递的标志，发现仪表板，在出现任何无效 json 的情况下向仪表板的状态字段添加错误消息，并自动 [安装](https://github.com/integr8ly/grafana-operator/blob/v3.2.0/documentation/dashboards.md#plugins) 指定的任何插件。 DevOps 和 SRE 然后可以在 GrafanaDataSource 中添加数据源以及添加额外的配置文件。

### **Jaeger Kubernetes Operator** 
The [Jaeger](https://github.com/jaegertracing/jaeger-operator) Kubernetes Operator helps deploy, manage and configure Jaeger instances on Kubernetes. When installing DevOps and SREs can specify [configuration](https://www.jaegertracing.io/docs/1.17/operator/#configuring-the-custom-resource) options for jaeger including storage options, deriving dependencies, injecting Jaeger agent sidecars and UI configurations among others. Once installed, the operator can be used to create a Jaeger instance and associate it with a deployment strategy. DevOps and SREs can choose any one of [three](https://www.jaegertracing.io/docs/1.17/operator/#deployment-strategies) supported deployment strategies: allInOne, production and streaming. When using the Production strategy, the operator will spin up a more scalable and highly available environment as well as deploy each of the backend components separately. The allInOne strategy is meant for testing and development purposes while the streaming strategy augments the production strategy and provides streaming between the collector and backend storage.

[Jaeger](https://github.com/jaegertracing/jaeger-operator) Kubernetes Operator 帮助在 Kubernetes 上部署、管理和配置 Jaeger 实例。安装 DevOps 和 SRE 时可以为 jaeger 指定 [configuration](https://www.jaegertracing.io/docs/1.17/operator/#configuring-the-custom-resource) 选项，包括存储选项、导出依赖项、注入 Jaeger 代理 sidecars和 UI 配置等。安装后，操作员可用于创建 Jaeger 实例并将其与部署策略相关联。 DevOps 和 SRE 可以选择 [三种](https://www.jaegertracing.io/docs/1.17/operator/#deployment-strategies) 支持的部署策略中的任何一种：allInOne、生产和流媒体。使用生产策略时，Operator将启动一个更具可扩展性和高度可用的环境，并分别部署每个后端组件。 allInOne 策略用于测试和开发目的，而流策略增强了生产策略并在收集器和后端存储之间提供流。

