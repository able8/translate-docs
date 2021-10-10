# Implementing SLI/SLO based Continuous Delivery Quality Gates using Prometheus

[Apr 7, 2020](https://medium.com/keptn/implementing-sli-slo-based-continuous-delivery-quality-gates-using-prometheus-9e17ec18ca36?source=post_page-----9e17ec18ca36--------------------------------) · 6 min read

[Google’s Book on Site Reliability Engineering](https://landing.google.com/sre/books/) has been a catalyst to have SRE’s, Performance and DevOps Engineers, or Cloud Operations incorporate the concept of

- **Service Level Indicators (SLIs)**, e.g: 95th percentile of your service’s response time
- **Service Level Objectives (SLOs)**, e.g: response time must not exceed 200ms during peak load

These concepts are great in production monitoring to ensure your organization meets your **Service Level Agreements (SLAs)** with your business. What we have seen is that the same concepts start “*shifting-left*” into the continuous delivery pipeline. To be more concrete: organizations start evaluating *every* build against their SLOs and let the result act as a *Quality Gate* before a build is promoted to production. Here is an example of 4 SLIs, their  SLOs and how each build gets a total score based on the individual SLO  results — determining whether the build is good enough to pass the  quality gate:

![img](https://miro.medium.com/freeze/max/60/1*P8f5NfJ3QVj6r8CvBsf8Sg.gif?q=20)

![img](https://miro.medium.com/max/1400/1*P8f5NfJ3QVj6r8CvBsf8Sg.gif)

SLO-based quality gate evaluation of 4 subsequent builds

When deploying workloads on Kubernetes, [Prometheus](https://prometheus.io/) comes to mind for capturing the values defined in your SLIs. The  missing component now is an implementation that automatically retrieves  all SLIs from a data source such as Prometheus, validates it against the SLOs and calculates an overall score that we can use for a quality  gate.

This is where the open-source project [Keptn](https://keptn.sh) comes in: an event-based control plane for continuous delivery and  automated operations. At its core it uses SLIs and SLOs to enforce  quality gates between delivery stages, to validate blue/green  deployments and canary releases and to auto-remediate problems in  production environments.

In this article we will focus on using *Keptn for Continuous Delivery* with Prometheus-based SLIs to evaluate quality gates. If you have  existing delivery pipelines and you just want to integrate quality gates then have a look at an example we have here: [Integrating Keptn Quality Gates with GitLab](https://www.youtube.com/watch?v=0JAGg6oC4UA). If you are interested in automated operations and self-healing have a look at the [tutorials on the Keptn Website](https://keptn.sh/docs/).

But now let’s get started with building our quality gates based on Prometheus with Keptn!

**Setting up and configuring Prometheus with Keptn**

Keptn can be used to setup and configure Prometheus if not already running in your cluster. The setup process is described in full detail on the [Keptn website](https://keptn.sh/docs/), and basically consists of these steps:

1. Download the [Keptn CLI from Github](https://github.com/keptn/keptn/releases) or initiate the download via: `curl -sL https://get.keptn.sh | sudo -E bash`
2. Install Keptn in your Kubernetes cluster via the Keptn CLI: `keptn install --platform=[gke|aks|eks|openshift|...]`
3. Create a project for your application that you want to have your quality gates set up:
   `keptn create project sockshop --shipyard=./sockshop.yaml`
4. Onboard your services you want to manage with Keptn:
   `keptn onboard service shoppingcart --project=sockshop --chart=./shoppingcart`
5. Configure Prometheus for your services: `keptn configure monitoring prometheus --project=sockshop --service=shoppingcart`

Now that we have installed Keptn and configured Prometheus, it is time to define our SLIs and SLOs for the quality gates.

**Setting up quality gates**

Defintions of quality gates in Keptn are based on the idea of having *service-level objectives* (SLOs) consisting of multiple *service level indicators* (SLIs) that define arbitrary metrics you can query from data providers such as Prometheus. Keptn abstracts from the actual API calls *how* to retrieve the values of the SLIs and focusses on the definition of *what* the SLO should evaluate. Let me illustrate this on an example.

In the following snippet I’ve created a quality gate that can be  automatically evaluated by Keptn and can be reused for different stages  and different microservices as it does not hold any particular  information about environments or services. The quality gates itself  holds two objectives: 1) the `reponse_time_p95` which is the the 95th percentile of the reponse time for the service and 2) the `error_rate` of the service. For the response time, the objective allows only a maximum increase of 25% to the previuous runs(see the `number_of_comparison_results` in the top of the quality gate) as well as an absolute threshold of  300ms. If both criteria are satisfied, this objective is fulfilled and  getting the full score. However, I’ve also specified a warning criteria  which means that a warning is issued (if within the boundaries of 600ms  in this case) and the objective will be evaluated with half the score.

The second objective targets the `error_rate` which has to be lower or equal to 5 for the objective to be fulfilled.  Since I consider the error rate as highly crucial I’ve defined this as a *key SLI* meaning that if this criteria is not met,  the quality gate will fail in its whole. In this sense I have control  which metrics are absolutely critical for my business and if quality  criteria upon them are not met the rollout of the service has to be  stopped.

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

Going into more technical detail, these are the actual Prometheus queries  that have been executed, but again, using Keptn there is no need in  being a Prometheus query expert:

- **response_time_p95**: `histogram_quantile(0.95,  sum(rate(http_response_time_milliseconds_bucket{job='<service>-<project>-<stage>-canary'}[<test_duration_in_seconds>s])) by (le))`
- **response_time_p50**: `histogram_quantile(0.50,  sum(rate(http_response_time_milliseconds_bucket{job='<service>-<project>-<stage>-canary'}[<test_duration_in_seconds>s])) by (le))`
- **error_rate**: `sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary",status!~'2..'}[<test_duration_in_seconds>s]))/sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary"}[<test_duration_in_seconds>s]))`
- **throughput**: `sum(rate(http_requests_total{job="<service>-<project>-<stage>-canary"}[<test_duration_in_seconds>s]))`

**Evaluation of Quality Gates**

Let us now take a look how Keptn evaluates the quality gates: Once Keptn is triggered either by the user or a CI/CD tool, it reaches out to the SLI provider, which in our case is Prometheus and all SLIs defined in the  quality gates will be queried for a given timeframe. This timeframe can  be user-defined or in case Keptn also triggers a test execution the time span of the test runs will be used. Next Keptn will evaluate the  metrics either against the absolute thresholds or to relatively to  previuous runs. Once the score is generated Keptn returns the results in the format of a [Cloud Event](https://cloudevents.io/) (an open-source specification for describing event data in a common way, part of the [CNCF](https://landscape.cncf.io/)). This allows that results can also be processed by external tools. Even [Integrations for Slack, MS Teams, etc](https://github.com/keptn-contrib/notification-service) to get notified about the deployment validation are available.

![img](https://miro.medium.com/freeze/max/60/1*GkcM5XiPGv2N5EJ-s-TFIQ.gif?q=20)

![img](https://miro.medium.com/max/1400/1*GkcM5XiPGv2N5EJ-s-TFIQ.gif)

Evaluation process of a quality gate

All evaluation runs and results are also visualized in the [Keptn’s Bridge](https://keptn.sh/docs/0.6.0/reference/keptnsbridge/) as you can see in the following screenshot. In our simple example  below, I have triggered 4 runs in total with run 1 and 2 passing the  quality gate while number 3 did not pass the quality gate (indicated in  red) and build number 4 was acceptable but raised a warning for the  response time that did not fully met the quality check.

![img](https://miro.medium.com/max/60/1*eaXihSENyhkcXHz1C6PaEQ.png?q=20)

![img](https://miro.medium.com/max/1400/1*eaXihSENyhkcXHz1C6PaEQ.png)

Evaluation of a quality gate

**How to get started!**

Now it is time for you to implement SLO-based continuous delivery quality  gates based on your Prometheus data. It is easy to get started with  Keptn — head over to [Keptn.sh](https://keptn.sh) and get the latest release, install it in your Kubernetes cluster, and  define your quality gates based on your SLOs! We are eager to hear from  you — shoot us a message in our [Keptn Slack](https://join.slack.com/t/keptn/shared_invite/zt-716xqbhz-w2xUeCpC1AgMMweGOYlPrA) or let us know via [Twitter](https://twitter.com/keptnProject) how your quality gates are improving the quality of your software to your end-users!

[keptn](https://medium.com/keptn?source=post_sidebar--------------------------post_sidebar-----------)

keptn — Cloud-native application life-cycle orchestration

- [Continuous Delivery](https://medium.com/keptn/tagged/continuous-delivery)
- [Prometheus](https://medium.com/keptn/tagged/prometheus)
- [Slo](https://medium.com/keptn/tagged/slo)
- [Sli](https://medium.com/keptn/tagged/sli)
- [Progressive Delivery](https://medium.com/keptn/tagged/progressive-delivery)
