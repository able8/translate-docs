# Introducing Tobs: Deploy a full observability suite for Kubernetes in two minutes

Tobs will deploy a full observability suite into your Kubernetes cluster with only a single command, so you can start collecting, analyzing, and visualizing all of your metrics in only a few minutes. While the set of components and their configuration is highly opinionated by default, Tobs is also highly extensible and configurable. This makes it easy to get started, but also gives you the flexibility to make the suite your own.

Today’s dynamic microservices-based architectures require advanced monitoring and observability capabilities. Unlike the monolithic systems of the past, today’s systems are composed of many interconnected components orchestrated through a significant degree of automation. Troubleshooting problems with modern systems requires shifting through the monitoring data of many individual components and piecing together enough information to conduct root cause analysis.

Observability tools help troubleshoot these problems and come in two main flavors: proprietary and open-source. Proprietary solutions such as DataDog and Splunk offer easy turn-key solutions for monitoring and observability. It is easy to get started with proprietary tools, and they may cover a great deal of the need for observability. But they tend to be rigid and inflexible, and users don’t want to be at the mercy of a SaaS provider's changing pricing structure.

Increasingly, users opt for open-source solutions instead. Open-source observability involves bringing together multiple interconnected tools, each of which tend to do one thing well — e.g., TimescaleDB for long-term data storage, Prometheus for collecting metrics, Fluentd for collecting logs, Grafana for visualization, Jaeger for tracing, etc. From a software engineering perspective, loosely coupled and highly customizable tools provide the flexibility we need.

However, setting up observability suites can be intimidating for developers. We have often seen new users get lost and overwhelmed when trying to implement observability into their systems for the first time. Luckily, infrastructure automation tools like Helm, Ansible, and Puppet can help, but they don’t help developers decide which tools to use or how to configure them together.

This is why we created Tobs, an open-source tool that leverages Helm and enables users to deploy an observability suite into any Kubernetes cluster with a simple command line instruction.

For example, the following command will install and configure **TimescaleDB, Promscale, Prometheus, Grafana, PromLens, Kube-State-Metrics, and Prometheus-Node-Exporter**:

```
tobs install
```

![](https://lh4.googleusercontent.com/bitVUwkieXnPX_BErI-AlXm4XoAUZjgO46x-jTcMjOHW8e5WDwGdIf39chNRbIcNUgpq95FwDHCQek1StTBhaXwyn7H6ct39BeOOvYT0bBCZgZwNP1HR-IB7vhjVFPHd9b9nI4k6)

Tobs allows users to get started with observability more easily: previously users had to first decide which tools they needed, how to set up and configure every tool individually, figure out how to connect and secure them, etc. — all before collecting any data or seeing any benefits.

With the ability to deploy an entire suite at once with a few commands, users get the power of several curated tools working together to collect and analyze their cluster right away. Users can then start seeing the benefits quickly and can customize their suite as their needs evolve. And, because modern systems emit relentless streams of observability data, we built Tobs around the rock-solid foundation of TimescaleDB, a petabyte-scale relational database for Prometheus data.

**Currently, Tobs can install the following tools to collect, store, and visualize metrics:**

- [TimescaleDB](https://www.timescale.com) stores and analyzes observability data over the long-term.
- [Promscale](https://www.timescale.com/promscale) provides the power of PromQL for data stored in TimescaleDB.
- [Prometheus](https://www.prometheus.io) collects metric data across your cluster.
- [Grafana](https://www.grafana.com) visualizes your data through customizable graphs and dashboards.
- [PromLens](https://promlens.com/) helps you build the PromQL queries you need to understand your system.
- [Kube-State-Metrics](https://github.com/kubernetes/kube-state-metrics) exposes metrics about the Kubernetes cluster itself.
- [Prometheus-Node-Exporter](https://github.com/prometheus/node_exporter) exposes metric data for nodes in your Kubernetes cluster.

![A flowchart displaying the components that make up Tobs](https://lh6.googleusercontent.com/iPU1TSotUSatjdsQrEkZgqfm0eL91o50jkjmmmjjCKQ-eS_yNsnydfrjsTuMCcSEOghv36JaSAsptyaMbCD4eiPcdbyBlOE7Pk5LP5P1Zmn8GBSBF1QFT9KeweF7N1SnN2YrQ816)The open-source components that make up Tobs

Tobs is a one-stop-shop for all your Kubernetes monitoring needs. We aim to expand the suite with support for logs and traces in the future (and we always [welcome community assistance](https://github.com/timescale/tobs/issues) with adding more components).

Because we're developers ourselves, we wanted to make Tobs as easy as possible to operate, while still allowing full customization. We recommend using the CLI, which makes it easy to install, upgrade, customize, and maintain the suite with just a few commands. For example, you can install the entire suite in two minutes (really). If you need to more tightly integrate with a more complex Helm setup, you can use our Helm chart without the CLI as a sub-chart.

A Tobs deployment will install the components listed above, connect, and secure them. By leveraging Prometheus service discovery, this deployment will then discover the components in your Kubernetes cluster that are emitting observability data and will start automatically collecting and storing it.

- [Download Tobs](https://github.com/timescale/tobs) today from our GitHub repository
- [Read the Tobs docs](https://github.com/timescale/tobs)
- [Contribute to the Tobs community](https://github.com/timescale/tobs/issues) and provide your input to the project

Read on for more information about how to use Tobs and the kinds of problems it can solve for you.

## Diving into Tobs

Tobs is a CLI tool designed to install all the observability stack components you need for monitoring in your Kubernetes cluster and provide complete lifecycle support for your monitoring stack. It abstracts all the actions for the observability stack with a single command.

You’ve already seen how easy it is to use Tobs to install a monitoring stack:

```
tobs install
```

This will install **TimescaleDB, Promscale, Prometheus, Grafana, PromLens, Kube-State-Metrics, and Prometheus-Node-Exporter.**

All the components deployed will be configured to connect with the other components. Also, Kubernetes dashboards are pre-configured in the Grafana UI. You can visualize all the observability data you can obtain from your Kubernetes cluster.

![](https://lh4.googleusercontent.com/upmOMr1pwM9A2EF6EPpDiizdg4uAtUrDOo5hG_XfTenZTZTa14wuURQ-ZbD6H5pHb92BT-7nOPXV-Z60T3szhN3UW8cLcT3ECJ3gZBavecipYQoDLDcVGCcSD3203VwuEo0kZQBo)

### Upgrading

To upgrade all the components in Tobs, simply run:

```
tobs upgrade

```

### Forwarding Ports

To access the TimescaleDB, Prometheus, Promscale, PromLens, or Grafana components running in the cluster on your local machine you can port-forward Tobs.

```
tobs port-forward
```

Then, simply access the component on its port on localhost.

### Resetting passwords

Tobs allows you to reset passwords for various components. For example, to reset the password for Grafana, you run:

```
tobs grafana change-password <new password>
```

### Configuring Metric Retention

Tobs allows you to set retention policies on a global basis and per metric basis. For example, to configure the retention policy to 2 days for `go_threads` metric, you run:

```
tobs metrics retention set go_threads 2
```

### Volume Expansion

Tobs offers you an easy way to expand persistent volumes claims (PVC’s) for TimescaleDB and Prometheus. For example, to expand Prometheus storage and TimescaleDB storage, you can run:

```
tobs volume expand --timescaleDB-storage 175Gi --timescaleDB-wal 25Gi --prometheus-storage 15Gi
```

### Integrating Tobs with an external TimescaleDB

You can connect Tobs to an external TimescaleDB (i.e., to an existing or external instance of TimescaleDB outside the k8s cluster) by providing the DB URI.

This will skip deploying of TimescaleDB during the Tobs installation and connects the rest of the stack to the provided DB URI.

```
tobs install --external-db-uri postgres://some_user:some_password@cloud.timescale.com:5432/tsdb?sslmode=prefer
```

## Components installed by Tobs

Let’s dive into each component that Tobs installs and configures on your behalf.

### **TimescaleDB**

We use TimescaleDB for long-term storage of metric data. Long-term storage provides the ability to perform post-hoc analysis on metric data over long-periods of time. Such data analysis can be used for capacity planning, identifying slow-moving regressions, trend analysis, auditing, and more.

We picked TimescaleDB as opposed to other systems because it is unique in the ability to perform analytics using SQL. This allows data to be used for much richer analysis than other stores. TimescaleDB also supports high-cardinality and is built on top of PostgreSQL, ensuring good reliability and durability of data as well as support for a wide-array of high-availability options.

Tobs also stores Grafana’s configuration data in TimescaleDB. This allows the Grafana deployment itself to be stateless, easing backup and reliability concerns.

[Learn more about TimescaleDB](https://www.timescale.com/products) and [try it for free today](https://www.timescale.com/timescale-signup).

### **Promscale**

When deploying TimescaleDB as long-term storage, Promscale provides the translation layer between Prometheus and the database. In particular, it allows the Prometheus server to store and retrieve metrics from TimescaleDB, and allows users to use PromQL on Promscale and Prometheus. (Plainly stated, Promscale is the obvious choice when connecting Prometheus and TimescaleDB.)

[Learn more about Promscale](https://www.timescale.com/promscale) and [read our blog post](https://blog.timescale.com/blog/promscale-analytical-platform-long-term-store-for-prometheus-combined-sql-promql-postgresql/) for more information.

### **Prometheus**

Prometheus is an open-source systems monitoring and altering stack. It has become the de-facto standard in metric monitoring and is the basis of standards such as OpenMetrics. It allows you to monitor and understand how your infrastructure and applications are performing. Service discovery allows Prometheus to automagically discover components within your Kubernetes cluster that are already emitting metrics.

[Learn more about Prometheus](https://github.com/prometheus/prometheus).

### Grafana

Grafana is a popular visualization tool to create and view rich dashboards based on metrics.

To help users gain insights into their cluster right away and see value, Tobs deploys grafana with pre-built dashboards to monitor Kubernetes.

[Learn more about Grafana](https://github.com/grafana/grafana).

### **PromLens**

A tool to help users build PromQL queries with ease. PromLens is a PromQL query builder that helps you build, understand, and fix your queries much more effectively.

As with any query language, PromQL can be challenging to learn. PromLens makes it easier and is thus an invaluable tool for users who are new to Prometheus and observability.

[Learn more about PromLens](https://promlens.com/).

### **Kube-State-Metrics**

Kube-state-metrics exports the metrics related to Kubernetes resources, e.g., the status and count of Kubernetes resources, with visibility of the desired resources and the current resources, as well as the trends in your cluster.

[Learn more about Kube-state-metrics](https://github.com/kubernetes/kube-state-metrics).

### **Node-Exporter**

Node-Exporter is deployed to export node related metrics (e.g.d, cpu, memory) from the Kubernetes cluster.

[Learn more about Node-Exporter](https://github.com/prometheus/node_exporter).

### **Missing something?**

Is your favorite tool not included yet? [Create an issue and let us know](https://github.com/timescale/tobs/issues), or better yet, [submit a pull request](https://github.com/timescale/tobs).

## Conclusion

Observability is increasingly critical in today’s world of complex microservices architecture.

Proprietary solutions are easy to get started, but can be inflexible and costly in the long-run. Open-source solutions are complex to configure and get started with , but can be fully customized and cost-effective once implemented.

We built Tobs to make open-source systems accessible to everyone. Tobs is the easy to use, open-source tool for deploying an observability suite into any Kubernetes cluster. With a simple command-line instruction, you can get up and running in under two minutes.

By reducing the startup time, we hope to spark even greater adoption of open-source observability solutions.
