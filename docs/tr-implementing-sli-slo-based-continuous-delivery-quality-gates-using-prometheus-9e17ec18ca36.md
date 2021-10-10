# Implementing SLI/SLO based Continuous Delivery Quality Gates using Prometheus

# 使用 Prometheus 实现基于 SLI/SLO 的持续交付质量门

[Apr 7, 2020](https://medium.com/keptn/implementing-sli-slo-based-continuous-delivery-quality-gates-using-prometheus-9e17ec18ca36?source=post_page-----9e17ec18ca36--------------------------------) · 6 min read

[Google’s Book on Site Reliability Engineering](https://landing.google.com/sre/books/) has been a catalyst to have SRE’s, Performance and DevOps Engineers, or Cloud Operations incorporate the concept of

[谷歌关于站点可靠性工程的书](https://landing.google.com/sre/books/) 已经成为让 SRE、性能和 DevOps 工程师或云运营整合概念的催化剂

- **Service Level Indicators (SLIs)**, e.g: 95th percentile of your service’s response time
- **Service Level Objectives (SLOs)**, e.g: response time must not exceed 200ms during peak load

- **服务水平指标 (SLI)**，例如：服务响应时间的第 95 个百分位
- **服务水平目标 (SLO)**，例如：峰值负载期间的响应时间不得超过 200 毫秒

These concepts are great in production monitoring to ensure your organization meets your **Service Level Agreements (SLAs)** with your business. What we have seen is that the same concepts start “*shifting-left*” into the continuous delivery pipeline. To be more concrete: organizations start evaluating *every* build against their SLOs and let the result act as a *Quality Gate* before a build is promoted to production. Here is an example of 4 SLIs, their  SLOs and how each build gets a total score based on the individual SLO  results — determining whether the build is good enough to pass the  quality gate:

这些概念非常适用于生产监控，可确保您的组织满足您的业务**服务水平协议 (SLA)**。我们看到的是，相同的概念开始“*左移*”进入持续交付管道。更具体地说：组织开始根据他们的 SLO 评估*每个*构建，并在构建被提升到生产之前让结果充当*质量门*。以下是 4 个 SLI、它们的 SLO 以及每个构建如何根据各个 SLO 结果获得总分的示例——确定构建是否足够好以通过质量门：

![img](https://miro.medium.com/freeze/max/60/1*P8f5NfJ3QVj6r8CvBsf8Sg.gif?q=20)

![img](https://miro.medium.com/max/1400/1*P8f5NfJ3QVj6r8CvBsf8Sg.gif)

SLO-based quality gate evaluation of 4 subsequent builds

4 个后续构建的基于 SLO 的质量门评估

When deploying workloads on Kubernetes, [Prometheus](https://prometheus.io/) comes to mind for capturing the values defined in your SLIs. The  missing component now is an implementation that automatically retrieves  all SLIs from a data source such as Prometheus, validates it against the SLOs and calculates an overall score that we can use for a quality  gate.

在 Kubernetes 上部署工作负载时，会想到 [Prometheus](https://prometheus.io/) 来捕获 SLI 中定义的值。现在缺少的组件是一个实现，它自动从 Prometheus 等数据源检索所有 SLI，根据 SLO 对其进行验证，并计算我们可用于质量门限的总体分数。

This is where the open-source project [Keptn](https://keptn.sh) comes in: an event-based control plane for continuous delivery and  automated operations. At its core it uses SLIs and SLOs to enforce  quality gates between delivery stages, to validate blue/green  deployments and canary releases and to auto-remediate problems in  production environments.

这就是开源项目 [Keptn](https://keptn.sh) 的用武之地：用于持续交付和自动化操作的基于事件的控制平面。它的核心是使用 SLI 和 SLO 在交付阶段之间实施质量关口，验证蓝/绿部署和金丝雀版本，并自动修复生产环境中的问题。

In this article we will focus on using *Keptn for Continuous Delivery* with Prometheus-based SLIs to evaluate quality gates. If you have  existing delivery pipelines and you just want to integrate quality gates then have a look at an example we have here: [Integrating Keptn Quality Gates with GitLab](https://www.youtube.com/watch?v=0JAGg6oC4UA) . If you are interested in automated operations and self-healing have a look at the [tutorials on the Keptn Website](https://keptn.sh/docs/).

在本文中，我们将重点介绍使用 *Keptn for Continuous Delivery* 和基于 Prometheus 的 SLI 来评估质量门。如果您有现有的交付管道并且只想集成质量门，请查看我们这里的示例：[将 Keptn 质量门与 GitLab 集成](https://www.youtube.com/watch?v=0JAGg6oC4UA) .如果您对自动化操作和自我修复感兴趣，请查看 [Keptn 网站上的教程](https://keptn.sh/docs/)。

But now let’s get started with building our quality gates based on Prometheus with Keptn!

但是现在让我们开始基于 Prometheus 和 Keptn 构建我们的质量门！

**Setting up and configuring Prometheus with Keptn**

**使用 Keptn 设置和配置 Prometheus**

Keptn can be used to setup and configure Prometheus if not already running in your cluster. The setup process is described in full detail on the [Keptn website](https://keptn.sh/docs/), and basically consists of these steps:

如果 Prometheus 尚未在集群中运行，则 Keptn 可用于设置和配置 Prometheus。 [Keptn 网站](https://keptn.sh/docs/) 上详细描述了设置过程，基本上包括以下步骤：

1. Download the [Keptn CLI from Github](https://github.com/keptn/keptn/releases) or initiate the download via: `curl -sL https://get.keptn.sh | sudo -E bash`
2. Install Keptn in your Kubernetes cluster via the Keptn CLI: `keptn install --platform=[gke|aks|eks|openshift|...]`
3. Create a project for your application that you want to have your quality gates set up:
    `keptn create project sockshop --shipyard=./sockshop.yaml`
4. Onboard your services you want to manage with Keptn:
    `keptn onboard service shoppingcart --project=sockshop --chart=./shoppingcart`
5. Configure Prometheus for your services: `keptn configure monitoring prometheus --project=sockshop --service=shoppingcart`

1. 从 Github 下载 [Keptn CLI](https://github.com/keptn/keptn/releases) 或通过以下方式启动下载：`curl -sL https://get.keptn.sh |须藤 -E bash`
2. 通过 Keptn CLI 在您的 Kubernetes 集群中安装 Keptn：`keptn install --platform=[gke|aks|eks|openshift|...]`
3. 为您要设置质量门的应用程序创建一个项目：
   `keptn 创建项目sockshop --shipyard=./sockshop.yaml`
4. 加入您想要使用 Keptn 管理的服务：
   `保持板载服务购物车 --project=sockshop --chart=./shoppingcart`
5. 为你的服务配置 Prometheus：`keptn configure monitoring prometheus --project=sockshop --service=shoppingcart`

Now that we have installed Keptn and configured Prometheus, it is time to define our SLIs and SLOs for the quality gates.

现在我们已经安装了 Keptn 并配置了 Prometheus，是时候为质量门定义我们的 SLI 和 SLO。

**Setting up quality gates** 

**设置质量门**

Defintions of quality gates in Keptn are based on the idea of having *service-level objectives* (SLOs) consisting of multiple *service level indicators* (SLIs) that define arbitrary metrics you can query from data providers such as Prometheus. Keptn abstracts from the actual API calls *how* to retrieve the values of the SLIs and focusses on the definition of *what* the SLO should evaluate. Let me illustrate this on an example.

Keptn 中质量门的定义基于具有*服务水平目标* (SLO) 的想法，该目标由多个 *服务水平指标* (SLI) 组成，这些指标定义了您可以从 Prometheus 等数据提供者查询的任意指标。 Keptn 从实际的 API 调用中抽象出*如何* 来检索 SLI 的值，并专注于 SLO 应该评估的*什么*的定义。让我用一个例子来说明这一点。

In the following snippet I’ve created a quality gate that can be  automatically evaluated by Keptn and can be reused for different stages  and different microservices as it does not hold any particular  information about environments or services. The quality gates itself  holds two objectives: 1) the `reponse_time_p95` which is the the 95th percentile of the reponse time for the service and 2) the `error_rate` of the service. For the response time, the objective allows only a maximum increase of 25% to the previuous runs(see the `number_of_comparison_results` in the top of the quality gate) as well as an absolute threshold of  300ms. If both criteria are satisfied, this objective is fulfilled and  getting the full score. However, I’ve also specified a warning criteria  which means that a warning is issued (if within the boundaries of 600ms  in this case) and the objective will be evaluated with half the score.

在下面的代码片段中，我创建了一个质量门，它可以由 Keptn 自动评估，并且可以在不同阶段和不同微服务中重用，因为它不包含有关环境或服务的任何特定信息。质量门本身有两个目标：1）“reponse_time_p95”，它是服务响应时间的第 95 个百分位；2）服务的“error_rate”。对于响应时间，目标仅允许比之前运行的最大增加 25%（参见质量门顶部的“number_of_comparison_results”）以及 300 毫秒的绝对阈值。如果这两个标准都满足，则该目标已实现并获得满分。但是，我还指定了一个警告标准，这意味着发出警告（如果在这种情况下在 600 毫秒的范围内）并且目标将用一半的分数进行评估。

The second objective targets the `error_rate` which has to be lower or equal to 5 for the objective to be fulfilled. Since I consider the error rate as highly crucial I’ve defined this as a *key SLI* meaning that if this criteria is not met,  the quality gate will fail in its whole. In this sense I have control  which metrics are absolutely critical for my business and if quality  criteria upon them are not met the rollout of the service has to be  stopped.

第二个目标的目标是“error_rate”，它必须低于或等于 5 才能实现目标。由于我认为错误率非常重要，因此我将其定义为 *key SLI*，这意味着如果不满足此标准，则质量门将完全失败。从这个意义上说，我可以控制哪些指标对我的业务绝对重要，如果不符合质量标准，则必须停止推出服务。

```
---
spec_version: '0.1.1'
comparison:
  compare_with: "several_results"
  number_of_comparison_results: 3
  include_result_with_score: "pass"
  aggregate_function: avg
objectives:
  - sli: response_time_p95
    pass:
      - criteria:
          - "<=+25%"    # relative values
          - "<300"      # absolute values
    warning:
      - criteria:
          - "<=600"
  - sli: error_rate
    pass:
      - criteria:
         - "<=5"
    key_sli: true       # if not met, evaluation fails
  - sli: throughput
  - sli: response_time_p50
total_score:
  pass: "90%"
  warning: "75%"
```

Well, how can this quality gate be evaluated? There are no API calls to  Prometheus nor is a service to evaluate defined? The answer lies in the  combination of a GitOps approach and Keptn built-in functionalities:  first, following a GitOps approach, all its configuration files are  version-controlled and managed in a Git repository. Each application  (consisting of multiple microservices) has its own repository and  folders and branches are used to distinguish between different  microservices in different environments (e.g., dev, hardening,  production). Second, Keptn has a built-in library for the most common  service-level indicators (SLIs) such as response time, failure rate and  throughput that is also easily extendible. Having this library allows us to write quality gates based on SLOs — without necessarily being an  expert in Prometheus queries. Furthermore, it allows to switch the  underlying monitoring tool and gather data from a different source, but  still keeping the same quality gates in place.

那么，如何评估这个质量门呢？没有对 Prometheus 的 API 调用，也没有定义要评估的服务？答案在于 GitOps 方法和 Keptn 内置功能的组合：首先，遵循 GitOps 方法，其所有配置文件都在 Git 存储库中进行版本控制和管理。每个应用程序（由多个微服务组成）都有自己的存储库和文件夹和分支，用于区分不同环境（例如，开发、强化、生产）中的不同微服务。其次，Keptn 为最常见的服务级别指标 (SLI) 提供了一个内置库，例如响应时间、故障率和吞吐量，这些指标也很容易扩展。有了这个库，我们就可以编写基于 SLO 的质量门——不一定是 Prometheus 查询方面的专家。此外，它允许切换底层监控工具并从不同来源收集数据，但仍保持相同的质量门限。

Going into more technical detail, these are the actual Prometheus queries  that have been executed, but again, using Keptn there is no need in  being a Prometheus query expert:

进入更多技术细节，这些是已执行的实际 Prometheus 查询，但同样，使用 Keptn 无需成为 Prometheus 查询专家：

- **response_time_p95**: `histogram_quantile(0.95,  sum(rate(http_response_time_milliseconds_bucket{job='<service>-<project>-<stage>-canary'}[<test_duration_in_seconds>s])) by (le))`
- **response_time_p50**: `histogram_quantile(0.50,  sum(rate(http_response_time_milliseconds_bucket{job='<service>-<project>-<stage>-canary'}[<test_duration_in_seconds>s])) by (le))`
- **error_rate**: `sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary",status!~'2..'}[<test_duration_in_seconds>s]))/ sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary"}[<test_duration_in_seconds>s]))`
- **throughput**: `sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary"}[<test_duration_in_seconds>s]))`

- **response_time_p95**：`histogram_quantile(0.95, sum(rate(http_response_time_milliseconds_bucket{job='<service>-<project>-<stage>-canary'}[<test_duration_in_seconds>s])) by (le))`
- **response_time_p50**：`histogram_quantile(0.50, sum(rate(http_response_time_milliseconds_bucket{job='<service>-<project>-<stage>-canary'}[<test_duration_in_seconds>s])) by (le))`
- **error_rate**: `sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary",status!~'2..'}[<test_duration_in_seconds>s]))/ sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary"}[<test_duration_in_seconds>s]))`
- **吞吐量**：`sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary"}[<test_duration_in_seconds>s]))`

**Evaluation of Quality Gates** 

**质量门的评估**

Let us now take a look how Keptn evaluates the quality gates: Once Keptn is triggered either by the user or a CI/CD tool, it reaches out to the SLI provider, which in our case is Prometheus and all SLIs defined in the  quality gates will be queried for a given timeframe. This timeframe can  be user-defined or in case Keptn also triggers a test execution the time span of the test runs will be used. Next Keptn will evaluate the  metrics either against the absolute thresholds or to relatively to  previuous runs. Once the score is generated Keptn returns the results in the format of a [Cloud Event](https://cloudevents.io/) (an open-source specification for describing event data in a common way, part of the [CNCF]( https://landscape.cncf.io/)). This allows that results can also be processed by external tools. Even [Integrations for Slack, MS Teams, etc](https://github.com/keptn-contrib/notification-service) to get notified about the deployment validation are available.

现在让我们看看 Keptn 如何评估质量门：一旦 Keptn 被用户或 CI/CD 工具触发，它就会联系 SLI 提供者，在我们的例子中是 Prometheus 和质量门中定义的所有 SLI将在给定的时间范围内查询。此时间范围可以是用户定义的，或者如果 Keptn 也触发测试执行，则将使用测试运行的时间跨度。 Next Keptn 将根据绝对阈值或相对于先前的运行来评估指标。生成分数后，Keptn 以 [Cloud Event](https://cloudevents.io/)（一种以通用方式描述事件数据的开源规范，[CNCF]的一部分)的格式返回结果https://landscape.cncf.io/))。这允许结果也可以由外部工具处理。甚至 [Slack、MS Teams 等的集成](https://github.com/keptn-contrib/notification-service) 也可以获得有关部署验证的通知。

![img](https://miro.medium.com/freeze/max/60/1*GkcM5XiPGv2N5EJ-s-TFIQ.gif?q=20)

![img](https://miro.medium.com/max/1400/1*GkcM5XiPGv2N5EJ-s-TFIQ.gif)

Evaluation process of a quality gate

质量门的评价过程

All evaluation runs and results are also visualized in the [Keptn’s Bridge](https://keptn.sh/docs/0.6.0/reference/keptnsbridge/) as you can see in the following screenshot. In our simple example  below, I have triggered 4 runs in total with run 1 and 2 passing the  quality gate while number 3 did not pass the quality gate (indicated in  red) and build number 4 was acceptable but raised a warning for the  response time that did not fully met the quality check.

所有评估运行和结果也在 [Keptn's Bridge](https://keptn.sh/docs/0.6.0/reference/keptnsbridge/) 中可视化，如下面的屏幕截图所示。在我们下面的简单示例中，我总共触发了 4 次运行，其中运行 1 和 2 通过了质量门，而编号 3 没有通过质量门（以红色表示)并且构建编号 4 是可以接受的，但对响应时间提出了警告不完全符合质量检查。

![img](https://miro.medium.com/max/60/1*eaXihSENyhkcXHz1C6PaEQ.png?q=20)

![img](https://miro.medium.com/max/1400/1*eaXihSENyhkcXHz1C6PaEQ.png)

Evaluation of a quality gate

质量门的评估

**How to get started!**

**如何开始！**

Now it is time for you to implement SLO-based continuous delivery quality  gates based on your Prometheus data. It is easy to get started with  Keptn — head over to [Keptn.sh](https://keptn.sh) and get the latest release, install it in your Kubernetes cluster, and  define your quality gates based on your SLOs! We are eager to hear from  you — shoot us a message in our [Keptn Slack](https://join.slack.com/t/keptn/shared_invite/zt-716xqbhz-w2xUeCpC1AgMMweGOYlPrA) or let us know via [Twitter]( https://twitter.com/keptnProject) how your quality gates are improving the quality of your software to your end-users!

现在是时候根据 Prometheus 数据实施基于 SLO 的持续交付质量门限了。开始使用 Keptn 很容易——前往 [Keptn.sh](https://keptn.sh) 并获取最新版本，将其安装在您的 Kubernetes 集群中，并根据您的 SLO 定义您的质量门！我们渴望收到您的来信 - 在我们的 [Keptn Slack](https://join.slack.com/t/keptn/shared_invite/zt-716xqbhz-w2xUeCpC1AgMMweGOYlPrA) 中给我们留言或通过 [Twitter]( https://twitter.com/keptnProject)您的质量门如何提高最终用户的软件质量！

[keptn](https://medium.com/keptn?source=post_sidebar--------------------------post_sidebar-----------)

keptn — Cloud-native application life-cycle orchestration

keepn — 云原生应用生命周期编排

- [Continuous Delivery](https://medium.com/keptn/tagged/continuous-delivery)
- [Prometheus](https://medium.com/keptn/tagged/prometheus)
- [Slo](https://medium.com/keptn/tagged/slo)
- [Sli](https://medium.com/keptn/tagged/sli)
- [Progressive Delivery](https://medium.com/keptn/tagged/progressive-delivery) 

- [持续交付](https://medium.com/keptn/tagged/continuous-delivery)
- [普罗米修斯](https://medium.com/keptn/tagged/prometheus)
- [Slo](https://medium.com/keptn/tagged/slo)
- [Sli](https://medium.com/keptn/tagged/sli)
- [渐进式交付](https://medium.com/keptn/tagged/progressive-delivery)

