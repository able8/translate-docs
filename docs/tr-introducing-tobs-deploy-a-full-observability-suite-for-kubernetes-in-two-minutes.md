# Introducing Tobs: Deploy a full observability suite for Kubernetes in two minutes

# Tobs 简介：两分钟内为 Kubernetes 部署一个完整的可观察性套件

Tobs will deploy a full observability suite into your Kubernetes cluster with only a single command, so you can start collecting, analyzing, and visualizing all of your metrics in only a few minutes. While the set of components and their configuration is highly opinionated by default, Tobs is also highly extensible and configurable. This makes it easy to get started, but also gives you the flexibility to make the suite your own.

Tob 只需一个命令即可将完整的可观察性套件部署到您的 Kubernetes 集群中，因此您可以在几分钟内开始收集、分析和可视化所有指标。虽然默认情况下组件集及其配置是高度自以为是的，但 Tob 也是高度可扩展和可配置的。这使您可以轻松入门，但也使您可以灵活地制作自己的套件。

Today’s dynamic microservices-based architectures require advanced monitoring and observability capabilities. Unlike the monolithic systems of the past, today’s systems are composed of many interconnected components orchestrated through a significant degree of automation. Troubleshooting problems with modern systems requires shifting through the monitoring data of many individual components and piecing together enough information to conduct root cause analysis.

当今基于动态微服务的架构需要先进的监控和可观察能力。与过去的单体系统不同，今天的系统由许多互连的组件组成，这些组件通过高度的自动化进行编排。对现代系统的问题进行故障排除需要转换许多单独组件的监控数据，并将足够的信息拼凑起来进行根本原因分析。

Observability tools help troubleshoot these problems and come in two main flavors: proprietary and open-source. Proprietary solutions such as DataDog and Splunk offer easy turn-key solutions for monitoring and observability. It is easy to get started with proprietary tools, and they may cover a great deal of the need for observability. But they tend to be rigid and inflexible, and users don’t want to be at the mercy of a SaaS provider's changing pricing structure.

可观察性工具有助于解决这些问题，主要有两种形式：专有和开源。 DataDog 和 Splunk 等专有解决方案为监控和可观察性提供了简单的交钥匙解决方案。使用专有工具很容易上手，它们可能涵盖了对可观察性的大量需求。但它们往往是僵化和不灵活的，用户不想受 SaaS 提供商不断变化的定价结构的摆布。

Increasingly, users opt for open-source solutions instead. Open-source observability involves bringing together multiple interconnected tools, each of which tend to do one thing well — eg, TimescaleDB for long-term data storage, Prometheus for collecting metrics, Fluentd for collecting logs, Grafana for visualization, Jaeger for tracing, etc . From a software engineering perspective, loosely coupled and highly customizable tools provide the flexibility we need.

越来越多的用户选择开源解决方案。开源可观察性涉及将多个相互关联的工具组合在一起，每个工具都倾向于做好一件事——例如，用于长期数据存储的 TimescaleDB、用于收集指标的 Prometheus、用于收集日志的 Fluentd、用于可视化的 Grafana、用于跟踪的 Jaeger 等. 从软件工程的角度来看，松耦合和高度可定制的工具提供了我们需要的灵活性。

However, setting up observability suites can be intimidating for developers. We have often seen new users get lost and overwhelmed when trying to implement observability into their systems for the first time. Luckily, infrastructure automation tools like Helm, Ansible, and Puppet can help, but they don’t help developers decide which tools to use or how to configure them together.

然而，设置可观察性套件对开发人员来说可能是令人生畏的。我们经常看到新用户在第一次尝试在他们的系统中实现可观察性时迷失方向并不知所措。幸运的是，像 Helm、Ansible 和 Puppet 这样的基础设施自动化工具可以提供帮助，但它们并不能帮助开发人员决定使用哪些工具或如何一起配置它们。

This is why we created Tobs, an open-source tool that leverages Helm and enables users to deploy an observability suite into any Kubernetes cluster with a simple command line instruction.

这就是我们创建 Tob 的原因，这是一个利用 Helm 的开源工具，使用户能够使用简单的命令行指令将可观察性套件部署到任何 Kubernetes 集群中。

For example, the following command will install and configure **TimescaleDB, Promscale, Prometheus, Grafana, PromLens, Kube-State-Metrics, and Prometheus-Node-Exporter**:

例如，以下命令将安装和配置 **TimescaleDB、Promscale、Prometheus、Grafana、PromLens、Kube-State-Metrics 和 Prometheus-Node-Exporter**：

```
tobs install
```

![](https://lh4.googleusercontent.com/bitVUwkieXnPX_BErI-AlXm4XoAUZjgO46x-jTcMjOHW8e5WDwGdIf39chNRbIcNUgpq95FwDHCQek1StTBhaXwyn7H6ct39BeOOvYT0bBCZgZwNP1HR-IB7vhjVFPHd9b9nI4k6)

Tobs allows users to get started with observability more easily: previously users had to first decide which tools they needed, how to set up and configure every tool individually, figure out how to connect and secure them, etc. — all before collecting any data or seeing any benefits.

Tob 允许用户更轻松地开始使用可观察性：以前用户必须首先决定他们需要哪些工具，如何单独设置和配置每个工具，弄清楚如何连接和保护它们等等——所有这些都是在收集任何数据或看到任何好处。

With the ability to deploy an entire suite at once with a few commands, users get the power of several curated tools working together to collect and analyze their cluster right away. Users can then start seeing the benefits quickly and can customize their suite as their needs evolve. And, because modern systems emit relentless streams of observability data, we built Tobs around the rock-solid foundation of TimescaleDB, a petabyte-scale relational database for Prometheus data.

通过使用几个命令一次部署整个套件的能力，用户可以获得多个精选工具协同工作的强大功能，以立即收集和分析他们的集群。然后，用户可以开始快速看到好处，并可以随着需求的发展定制他们的套件。而且，由于现代系统会不断发出可观察性数据流，我们围绕 TimescaleDB 坚如磐石的基础构建了 Tob，TimescaleDB 是一个 PB 级的 Prometheus 数据关系数据库。

**Currently, Tobs can install the following tools to collect, store, and visualize metrics:**

**目前，Tobs 可以安装以下工具来收集、存储和可视化指标：**

- [TimescaleDB](https://www.timescale.com) stores and analyzes observability data over the long-term.
- [Promscale](https://www.timescale.com/promscale) provides the power of PromQL for data stored in TimescaleDB.
- [Prometheus](https://www.prometheus.io) collects metric data across your cluster.
- [Grafana](https://www.grafana.com) visualizes your data through customizable graphs and dashboards.
- [PromLens](https://promlens.com/) helps you build the PromQL queries you need to understand your system.
- [Kube-State-Metrics](https://github.com/kubernetes/kube-state-metrics) exposes metrics about the Kubernetes cluster itself. 

- [TimescaleDB](https://www.timescale.com) 长期存储和分析可观察性数据。
- [Promscale](https://www.timescale.com/promscale) 为存储在 TimescaleDB 中的数据提供 PromQL 的强大功能。
- [Prometheus](https://www.prometheus.io) 收集整个集群的指标数据。
- [Grafana](https://www.grafana.com) 通过可定制的图形和仪表板将您的数据可视化。
- [PromLens](https://promlens.com/) 可帮助您构建了解系统所需的 PromQL 查询。
- [Kube-State-Metrics](https://github.com/kubernetes/kube-state-metrics) 公开了关于 Kubernetes 集群本身的指标。

- [Prometheus-Node-Exporter](https://github.com/prometheus/node_exporter) exposes metric data for nodes in your Kubernetes cluster.

- [Prometheus-Node-Exporter](https://github.com/prometheus/node_exporter) 公开 Kubernetes 集群中节点的指标数据。

![A flowchart displaying the components that make up Tobs](https://lh6.googleusercontent.com/iPU1TSotUSatjdsQrEkZgqfm0eL91o50jkjmmmjjCKQ-eS_yNsnydfrjsTuMCcSEOghv36JaSAsptyaMbCD4eiPcdbyBlOE7Pk5LP5P1Zmn8GBSBF1QFT9KeweF7N1SnN2YrQ816)The open-source components that make up Tobs

Tobs is a one-stop-shop for all your Kubernetes monitoring needs. We aim to expand the suite with support for logs and traces in the future (and we always [welcome community assistance](https://github.com/timescale/tobs/issues) with adding more components).

Tobs 是满足您所有 Kubernetes 监控需求的一站式商店。我们的目标是在未来扩展套件以支持日志和跟踪（我们总是 [欢迎社区帮助](https://github.com/timescale/tobs/issues)添加更多组件)。

Because we're developers ourselves, we wanted to make Tobs as easy as possible to operate, while still allowing full customization. We recommend using the CLI, which makes it easy to install, upgrade, customize, and maintain the suite with just a few commands. For example, you can install the entire suite in two minutes (really). If you need to more tightly integrate with a more complex Helm setup, you can use our Helm chart without the CLI as a sub-chart.

因为我们自己是开发人员，所以我们希望让 Tob 尽可能易于操作，同时仍然允许完全自定义。我们建议使用 CLI，只需几个命令即可轻松安装、升级、自定义和维护套件。例如，您可以在两分钟内安装整个套件（真的）。如果您需要与更复杂的 Helm 设置更紧密地集成，您可以使用我们的 Helm 图表而不使用 CLI 作为子图表。

A Tobs deployment will install the components listed above, connect, and secure them. By leveraging Prometheus service discovery, this deployment will then discover the components in your Kubernetes cluster that are emitting observability data and will start automatically collecting and storing it.

Tob 部署将安装上面列出的组件，连接并保护它们。通过利用 Prometheus 服务发现，此部署将发现 Kubernetes 集群中发出可观察性数据的组件，并将开始自动收集和存储这些数据。

- [Download Tobs](https://github.com/timescale/tobs) today from our GitHub repository
- [Read the Tobs docs](https://github.com/timescale/tobs)
- [Contribute to the Tobs community](https://github.com/timescale/tobs/issues) and provide your input to the project

- 今天从我们的 GitHub 存储库 [下载 Tob](https://github.com/timescale/tobs)
- [阅读 Tobs 文档](https://github.com/timescale/tobs)
- [为 Tobs 社区做贡献](https://github.com/timescale/tobs/issues) 并为项目提供您的意见

Read on for more information about how to use Tobs and the kinds of problems it can solve for you.

请继续阅读有关如何使用 Tob 以及它可以为您解决的问题类型的更多信息。

## Diving into Tobs

## 潜入托布斯

Tobs is a CLI tool designed to install all the observability stack components you need for monitoring in your Kubernetes cluster and provide complete lifecycle support for your monitoring stack. It abstracts all the actions for the observability stack with a single command.

Tobs 是一个 CLI 工具，旨在安装您在 Kubernetes 集群中进行监控所需的所有可观察性堆栈组件，并为您的监控堆栈提供完整的生命周期支持。它使用单个命令抽象了可观察性堆栈的所有操作。

You’ve already seen how easy it is to use Tobs to install a monitoring stack:

您已经看到使用 Tob 安装监控堆栈是多么容易：

```
tobs install
```

This will install **TimescaleDB, Promscale, Prometheus, Grafana, PromLens, Kube-State-Metrics, and Prometheus-Node-Exporter.**

这将安装 **TimescaleDB、Promscale、Prometheus、Grafana、PromLens、Kube-State-Metrics 和 Prometheus-Node-Exporter。**

All the components deployed will be configured to connect with the other components. Also, Kubernetes dashboards are pre-configured in the Grafana UI. You can visualize all the observability data you can obtain from your Kubernetes cluster.

部署的所有组件都将配置为与其他组件连接。此外，Kubernetes 仪表板已在 Grafana UI 中预先配置。您可以可视化您可以从 Kubernetes 集群中获取的所有可观察性数据。

![](https://lh4.googleusercontent.com/upmOMr1pwM9A2EF6EPpDiizdg4uAtUrDOo5hG_XfTenZTZTa14wuURQ-ZbD6H5pHb92BT-7nOPXV-Z60T3szhN3UW8cLcT3ECJ3gZBavecipYQoDLDcVGCcSD3203VwuEo0kZQBo)

### Upgrading

### 升级

To upgrade all the components in Tobs, simply run:

要升级 Tob 中的所有组件，只需运行：

```
tobs upgrade

```

### Forwarding Ports

### 转发端口

To access the TimescaleDB, Prometheus, Promscale, PromLens, or Grafana components running in the cluster on your local machine you can port-forward Tobs.

要访问本地机器上集群中运行的 TimescaleDB、Prometheus、Promscale、PromLens 或 Grafana 组件，您可以端口转发 Tob。

```
tobs port-forward
```

Then, simply access the component on its port on localhost.

然后，只需在本地主机上的端口上访问该组件。

### Resetting passwords

### 重置密码

Tobs allows you to reset passwords for various components. For example, to reset the password for Grafana, you run:

Tob 允许您重置各种组件的密码。例如，要重置 Grafana 的密码，请运行：

```
tobs grafana change-password <new password>
```

### Configuring Metric Retention

### 配置指标保留

Tobs allows you to set retention policies on a global basis and per metric basis. For example, to configure the retention policy to 2 days for `go_threads` metric, you run:

Tob 允许您在全局和每个指标的基础上设置保留策略。例如，要将“go_threads”指标的保留策略配置为 2 天，请运行：

```
tobs metrics retention set go_threads 2
```

### Volume Expansion

### 卷扩展

Tobs offers you an easy way to expand persistent volumes claims (PVC’s) for TimescaleDB and Prometheus. For example, to expand Prometheus storage and TimescaleDB storage, you can run:

Tobs 为您提供了一种简单的方法来扩展 TimescaleDB 和 Prometheus 的持久卷声明 (PVC)。例如，要扩展 Prometheus 存储和 TimescaleDB 存储，您可以运行：

```
tobs volume expand --timescaleDB-storage 175Gi --timescaleDB-wal 25Gi --prometheus-storage 15Gi
```

### Integrating Tobs with an external TimescaleDB

### 将 Tob 与外部 TimescaleDB 集成

You can connect Tobs to an external TimescaleDB (i.e., to an existing or external instance of TimescaleDB outside the k8s cluster) by providing the DB URI.

您可以通过提供 DB URI 将 Tob 连接到外部 TimescaleDB（即连接到 k8s 集群外部的 TimescaleDB 的现有或外部实例）。

This will skip deploying of TimescaleDB during the Tobs installation and connects the rest of the stack to the provided DB URI.

这将在 Tobs 安装期间跳过 TimescaleDB 的部署，并将堆栈的其余部分连接到提供的 DB URI。

```
tobs install --external-db-uri postgres://some_user:some_password@cloud.timescale.com:5432/tsdb?sslmode=prefer
```

## Components installed by Tobs

## 组件由 Tobs 安装

Let’s dive into each component that Tobs installs and configures on your behalf.

让我们深入了解 Tobs 代表您安装和配置的每个组件。

### **TimescaleDB** 

### **时间刻度数据库**

We use TimescaleDB for long-term storage of metric data. Long-term storage provides the ability to perform post-hoc analysis on metric data over long-periods of time. Such data analysis can be used for capacity planning, identifying slow-moving regressions, trend analysis, auditing, and more.

我们使用 TimescaleDB 来长期存储指标数据。长期存储提供了对长期指标数据执行事后分析的能力。此类数据分析可用于容量规划、识别缓慢回归、趋势分析、审计等。

We picked TimescaleDB as opposed to other systems because it is unique in the ability to perform analytics using SQL. This allows data to be used for much richer analysis than other stores. TimescaleDB also supports high-cardinality and is built on top of PostgreSQL, ensuring good reliability and durability of data as well as support for a wide-array of high-availability options.

我们选择 TimescaleDB 而不是其他系统，因为它在使用 SQL 执行分析的能力方面是独一无二的。这允许将数据用于比其他商店更丰富的分析。 TimescaleDB 还支持高基数，并建立在 PostgreSQL 之上，确保数据的良好可靠性和持久性，并支持广泛的高可用性选项。

Tobs also stores Grafana’s configuration data in TimescaleDB. This allows the Grafana deployment itself to be stateless, easing backup and reliability concerns.

Tob 还将 Grafana 的配置数据存储在 TimescaleDB 中。这允许 Grafana 部署本身是无状态的，从而缓解备份和可靠性问题。

[Learn more about TimescaleDB](https://www.timescale.com/products) and [try it for free today](https://www.timescale.com/timescale-signup).

[了解有关 TimescaleDB 的更多信息](https://www.timescale.com/products) 和 [立即免费试用](https://www.timescale.com/timescale-signup)。

### **Promscale**

### **舞会**

When deploying TimescaleDB as long-term storage, Promscale provides the translation layer between Prometheus and the database. In particular, it allows the Prometheus server to store and retrieve metrics from TimescaleDB, and allows users to use PromQL on Promscale and Prometheus. (Plainly stated, Promscale is the obvious choice when connecting Prometheus and TimescaleDB.)

在部署 TimescaleDB 作为长期存储时，Promscale 提供了 Prometheus 和数据库之间的翻译层。特别是，它允许 Prometheus 服务器从 TimescaleDB 存储和检索指标，并允许用户在 Promscale 和 Prometheus 上使用 PromQL。 （简单地说，Promscale 是连接 Prometheus 和 TimescaleDB 的明显选择。）

[Learn more about Promscale](https://www.timescale.com/promscale) and [read our blog post](https://blog.timescale.com/blog/promscale-analytical-platform-long-term-store-for-prometheus-combined-sql-promql-postgresql/) for more information.

[了解有关 Promscale 的更多信息](https://www.timescale.com/promscale) 和 [阅读我们的博客文章](https://blog.timescale.com/blog/promscale-analytical-platform-long-term-store-for-prometheus-combined-sql-promql-postgresql/) 了解更多信息。

### **Prometheus**

### **普罗米修斯**

Prometheus is an open-source systems monitoring and altering stack. It has become the de-facto standard in metric monitoring and is the basis of standards such as OpenMetrics. It allows you to monitor and understand how your infrastructure and applications are performing. Service discovery allows Prometheus to automagically discover components within your Kubernetes cluster that are already emitting metrics.

Prometheus 是一个开源系统监控和更改堆栈。它已成为指标监控的事实上的标准，并且是 OpenMetrics 等标准的基础。它允许您监控和了解您的基础设施和应用程序的执行情况。服务发现允许 Prometheus 自动发现 Kubernetes 集群中已经发出指标的组件。

[Learn more about Prometheus](https://github.com/prometheus/prometheus).

[了解有关 Prometheus 的更多信息](https://github.com/prometheus/prometheus)。

### Grafana

### 格拉法纳

Grafana is a popular visualization tool to create and view rich dashboards based on metrics.

Grafana 是一种流行的可视化工具，用于根据指标创建和查看丰富的仪表板。

To help users gain insights into their cluster right away and see value, Tobs deploys grafana with pre-built dashboards to monitor Kubernetes.

为了帮助用户立即深入了解他们的集群并看到价值，Tobs 部署了带有预构建仪表板的 grafana 来监控 Kubernetes。

[Learn more about Grafana](https://github.com/grafana/grafana).

[了解有关 Grafana 的更多信息](https://github.com/grafana/grafana)。

### **PromLens**

### **PromLens**

A tool to help users build PromQL queries with ease. PromLens is a PromQL query builder that helps you build, understand, and fix your queries much more effectively.

帮助用户轻松构建 PromQL 查询的工具。 PromLens 是 PromQL 查询构建器，可帮助您更有效地构建、理解和修复查询。

As with any query language, PromQL can be challenging to learn. PromLens makes it easier and is thus an invaluable tool for users who are new to Prometheus and observability.

与任何查询语言一样，PromQL 可能很难学习。 PromLens 使它变得更容易，因此对于 Prometheus 和可观察性的新用户来说是一个非常宝贵的工具。

[Learn more about PromLens](https://promlens.com/).

[了解有关 PromLens 的更多信息](https://promlens.com/)。

### **Kube-State-Metrics**

### **Kube 状态指标**

Kube-state-metrics exports the metrics related to Kubernetes resources, e.g., the status and count of Kubernetes resources, with visibility of the desired resources and the current resources, as well as the trends in your cluster.

Kube-state-metrics 导出与 Kubernetes 资源相关的指标，例如 Kubernetes 资源的状态和数量，以及所需资源和当前资源的可见性，以及集群中的趋势。

[Learn more about Kube-state-metrics](https://github.com/kubernetes/kube-state-metrics).

[了解有关 Kube-state-metrics 的更多信息](https://github.com/kubernetes/kube-state-metrics)。

### **Node-Exporter**

### **节点导出器**

Node-Exporter is deployed to export node related metrics (e.g.d, cpu, memory) from the Kubernetes cluster.

Node-Exporter 用于从 Kubernetes 集群导出节点相关的指标（例如 d、cpu、内存）。

[Learn more about Node-Exporter](https://github.com/prometheus/node_exporter).

[了解有关 Node-Exporter 的更多信息](https://github.com/prometheus/node_exporter)。

### **Missing something?**

###  **遗漏了什么？**

Is your favorite tool not included yet? [Create an issue and let us know](https://github.com/timescale/tobs/issues), or better yet, [submit a pull request](https://github.com/timescale/tobs).

你最喜欢的工具还没有包括在内吗？ [创建问题并让我们知道](https://github.com/timescale/tobs/issues)，或者更好的是，[提交拉取请求](https://github.com/timescale/tobs)。

## Conclusion

##  结论

Observability is increasingly critical in today’s world of complex microservices architecture.

在当今复杂的微服务架构世界中，可观察性越来越重要。

Proprietary solutions are easy to get started, but can be inflexible and costly in the long-run. Open-source solutions are complex to configure and get started with , but can be fully customized and cost-effective once implemented.

专有解决方案易于上手，但从长远来看可能不够灵活且成本高昂。开源解决方案的配置和入门很复杂，但一旦实施就可以完全定制并且具有成本效益。

We built Tobs to make open-source systems accessible to everyone. Tobs is the easy to use, open-source tool for deploying an observability suite into any Kubernetes cluster. With a simple command-line instruction, you can get up and running in under two minutes.

我们构建 Tob 是为了让每个人都可以访问开源系统。 Tobs 是一种易于使用的开源工具，用于将可观察性套件部署到任何 Kubernetes 集群中。使用简单的命令行指令，您可以在两分钟内启动并运行。

By reducing the startup time, we hope to spark even greater adoption of open-source observability solutions. 

通过减少启动时间，我们希望激发更多采用开源可观察性解决方案。

