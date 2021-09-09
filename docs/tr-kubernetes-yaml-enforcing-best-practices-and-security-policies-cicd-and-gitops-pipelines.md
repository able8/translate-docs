# Kubernetes YAML: Enforcing best practices and security policies in CI/CD and GitOps pipelines

# Kubernetes YAML：在 CI/CD 和 GitOps 管道中实施最佳实践和安全策略

Using [Kubernetes](https://thechief.io/c/cloudplex/how-get-started-kubernetes/) is a synonym for manipulating YAML.

使用 [Kubernetes](https://thechief.io/c/cloudplex/how-get-started-kubernetes/) 是操作 YAML 的同义词。

Even if the YAML community describes it as a human-readable language, it can sometimes be tricky to read it, especially in the context of Kubernetes when you are manipulating complex deployments, services, ingresses, and other resources. This may lead to [incoherences and problems](https://thechief.io/c/editorial/4-kubernetes-failure-stories-learn/) that may take some time to solve.

即使 YAML 社区将其描述为一种人类可读的语言，阅读它有时也很棘手，尤其是在 Kubernetes 的上下文中，当您操作复杂的部署、服务、入口和其他资源时。这可能会导致 [不连贯和问题](https://thechief.io/c/editorial/4-kubernetes-failure-stories-learn/) 可能需要一些时间来解决。

* * *

* * *

Before applying your YAML configuration to a [Kubernetes](https://thechief.io/c/cloudplex/21-resources-and-tutorials-learn-kubernetes/) cluster, it is evident and crucial to check if it's valid. Reading through your YAML configuration is the obvious way to do it. However, it's neither the fastest nor the easiest way to do it. On the other hand, if you are automating your deployments or testing, you will need to implement this step as part of your automated workflows.

在将您的 YAML 配置应用于 [Kubernetes](https://thechief.io/c/cloudplex/21-resources-and-tutorials-learn-kubernetes/) 集群之前，检查它是否有效很明显且至关重要。通读您的 YAML 配置是显而易见的方法。但是，这既不是最快也不是最简单的方法。另一方面，如果您要自动化部署或测试，则需要将此步骤作为自动化工作流的一部分来实施。

Fortunately, the [Kubernetes](https://thechief.io/c/cloudplex/kubernetes-busy-developer/) and [Cloud Native](https://thechief.io/c/editorial/the-cios-guide-to-kubernetes-and-eventual-transition-to-cloud-native-development/) developer community has developed (and still) tools that make validating easy and especially automated. These tools can be extremely helpful when implemented in [CI/CD](https://thechief.io/c/cloudplex/top-10-kubernetes-cicd-tools/) or continuous GitOps pipelines. In addition to that, these tools help you in enforcing best practices, as well as applying policies and compliance requirements.

幸运的是，[Kubernetes](https://thechief.io/c/cloudplex/kubernetes-busy-developer/) 和 [Cloud Native](https://thechief.io/c/editorial/the-cios-guide-to-kubernetes-and-eventual-transition-to-cloud-native-development/) 开发者社区已经开发（并且仍然)使验证变得容易且特别自动化的工具。这些工具在 [CI/CD](https://thechief.io/c/cloudplex/top-10-kubernetes-cicd-tools/) 或连续 GitOps 管道中实施时非常有用。除此之外，这些工具还可以帮助您实施最佳实践，以及应用策略和合规性要求。

Let's discover some of these tools.

让我们发现其中的一些工具。

## [KubeLinter](https://github.com/stackrox/kube-linter)

## [KubeLinter](https://github.com/stackrox/kube-linter)

KubeLinter offers the automated analysis of [Kubernetes](https://thechief.io/c/cloudplex/kubernetes-101/) YAML files and HELM charts before deployment.

KubeLinter 提供了在部署前对 [Kubernetes](https://thechief.io/c/cloudplex/kubernetes-101/) YAML 文件和 HELM 图表的自动分析。

It integrates [security-as-a-code](https://thechief.io/c/news/stackrox-announces-release-kubelinter/) in DevOps and other related processes. It helps ensure that the [Kubernetes](https://thechief.io/c/metricfire/cicd-pipelines-kubernetes-applications/) configuration is correct and the automatic enforcement of [security](https://thechief.io/c/editorial/takeaways-state-containers-and-kubernetes-security/) policies for Kubernetes applications.

它在 DevOps 和其他相关流程中集成了 [security-as-a-code](https://thechief.io/c/news/stackrox-announces-release-kubelinter/)。它有助于确保[Kubernetes](https://thechief.io/c/metricfire/cicd-pipelines-kubernetes-applications/) 配置正确并自动执行 [security](https://thechief.io/c/editor/takeaways-state-containers-and-kubernetes-security/) Kubernetes 应用程序的策略。

Users can build [security](https://thechief.io/c/news/kubernetes-security-specialist-certification-now-available/) in the configuration as code in the application development process itself. It helps validate that the [Kubernetes](https://thechief.io/c/metricfire/aws-ecs-vs-kubernetes/) configuration is according to the security best practices.

用户可以在应用程序开发过程中以代码的形式在配置中构建[安全](https://thechief.io/c/news/kubernetes-security-specialist-certification-now-available/)。它有助于验证[Kubernetes](https://thechief.io/c/metricfire/aws-ecs-vs-kubernetes/) 配置是否符合安全最佳实践。

KubeLinter users can integrate this tool to automate the process of carrying out configuration checks and error identification.

KubeLinter 用户可以集成这个工具来自动化执行配置检查和错误识别的过程。

## [Kubeval](https://github.com/instrumenta/kubeval/)

## [Kubeval](https://github.com/instrumenta/kubeval/)

Kubeval is another tool to validate your [Kubernetes](https://thechief.io/c/metricfire/kubernetes-networking-101/) configuration files. It supports multiple [Kubernetes](https://thechief.io/c/news/challenges-enterprise-kubernetes/) versions and works with YAML or JSON [configuration](https://thechief.io/c/news/red-hat-announces-integration-ansible-automation-openshift-kubernetes/) files.

Kubeval 是另一种验证 [Kubernetes](https://thechief.io/c/metricfire/kubernetes-networking-101/) 配置文件的工具。它支持多个 [Kubernetes](https://thechief.io/c/news/challenges-enterprise-kubernetes/) 版本并与 YAML 或 JSON [配置](https://thechief.io/c/news/red-hat-announces-integration-ansible-automation-openshift-kubernetes/) 文件。

This tool can be used in [CI/CD pipelines](https://thechief.io/c/metricfire/cicd-pipelines-kubernetes-applications/) as well as local development and [testing](https://thechief.io/c/cloudplex/kubernetes-distributed-performance-testing-using-locust/). It has the advantage of being a single binary that you can [download](https://www.kubeval.com/installation/) and run smoothly without any installation or configuration effort.

此工具可用于 [CI/CD 管道](https://thechief.io/c/metricfire/cicd-pipelines-kubernetes-applications/) 以及本地开发和 [测试](https://thechief.io/c)。io/c/cloudplex/kubernetes-distributed-performance-testing-using-locust/)。它的优点是作为单个二进制文件，您可以 [下载](https://www.kubeval.com/installation/) 并顺利运行，无需任何安装或配置工作。

For a Linux machine, the installation is as easy as:

对于 Linux 机器，安装非常简单：

```
wget https://github.com/instrumenta/kubeval/releases/latest/download/kubeval-linux-amd64.tar.gz
tar xf kubeval-linux-amd64.tar.gz
sudo cp kubeval /usr/local/bin
```


## [kube-score](https://github.com/zegl/kube-score)

## [kube-score](https://github.com/zegl/kube-score)

kube-score is a tool that performs static code analysis of your [Kubernetes](https://thechief.io/c/abvijaykumar/kubernetes-operators-realize-dream-zero-touch-ops/) object definitions.

kube-score 是一个对你的 [Kubernetes](https://thechief.io/c/abvijaykumar/kubernetes-operators-realize-dream-zero-touch-ops/) 对象定义进行静态代码分析的工具。

The output is a list of recommendations of what you can improve to make your application more secure and resilient.

输出是您可以改进哪些方面的建议列表，以使您的应用程序更安全和更有弹性。

For the following input:

对于以下输入：

```
apiVersion: apps/v1
kind: Deployment
metadata:
name: statefulset-test-1
spec:
template:
    metadata:
      labels:
        app: foo
    spec:
      containers:
      - name: foobar
        image: foo:bar
---
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
name: app-budget
spec:
minAvailable: 2
selector:
    matchLabels:
      app: not-foo
```




Kube-score will output the following list of recommendations:

Kube-score 将输出以下推荐列表：

```
apps/v1/Deployment statefulset-test-1                                         💥
    [CRITICAL] Container Resources
        · foobar -> CPU limit is not set
            Resource limits are recommended to avoid resource DDOS.Set resources.limits.cpu
        · foobar -> Memory limit is not set
            Resource limits are recommended to avoid resource DDOS.Set resources.limits.memory
        · foobar -> CPU request is not set
            Resource requests are recommended to make sure that the application can start and run without
            crashing.Set resources.requests.cpu
        · foobar -> Memory request is not set
            Resource requests are recommended to make sure that the application can start and run without
            crashing.Set resources.requests.memory
    [CRITICAL] Container Image Pull Policy
        · foobar -> ImagePullPolicy is not set to Always
            It's recommended to always set the ImagePullPolicy to Always, to make sure that the
            imagePullSecrets are always correct, and to always get the image you want.
    [CRITICAL] Pod NetworkPolicy
        · The pod does not have a matching NetworkPolicy
            Create a NetworkPolicy that targets this pod to control who/what can communicate with this pod.
            Note, this feature needs to be supported by the CNI implementation used in the Kubernetes cluster
            to have an effect.
    [CRITICAL] Container Security Context
        · foobar -> Container has no configured security context
            Set securityContext to run the container in a more secure context.
    [CRITICAL] Deployment has PodDisruptionBudget
        · No matching PodDisruptionBudget was found
            It's recommended to define a PodDisruptionBudget to avoid unexpected downtime during Kubernetes
            maintenance operations, such as when draining a node.
    [WARNING] Deployment has host PodAntiAffinity
        · Deployment does not have a host podAntiAffinity set
            It's recommended to set a podAntiAffinity that stops multiple pods from a deployment from being
            scheduled on the same node.This increases availability in case the node becomes unavailable.
policy/v1beta1/PodDisruptionBudget app-budget                                 ✅
```


## [**config-lint**](https://github.com/stelligent/config-lint)

## [**config-lint**](https://github.com/stelligent/config-lint)

This is a command-line tool to validate configuration files using rules specified in YAML. The configuration files can be one of several formats: [Terraform](https://thechief.io/c/codersociety/creating-and-manageing-ha-aks-kubernetes-cluster-azure-thanks-terraform/), JSON, YAML, with support for Kubernetes.

这是一个命令行工具，用于使用 YAML 中指定的规则验证配置文件。配置文件可以是以下几种格式之一：[Terraform](https://thechief.io/c/codersociety/creating-and-manageing-ha-aks-kubernetes-cluster-azure-thanks-terraform/)、JSON、 YAML，支持 Kubernetes。

There are built-in rules provided for Terraform, and custom files can be used for other formats.

Terraform 提供了内置规则，自定义文件可用于其他格式。

This tool, developed by Stelligent, has the advantage of being customizable since you can develop custom rules. In a blog post explaining how to use this tool, Dan Miller states:

这个由 Stelligent 开发的工具具有可定制的优势，因为您可以开发自定义规则。在解释如何使用此工具的博客文章中，Dan Miller 指出：

> In addition to validating against our built-in ruleset, config-lint provides a simple and powerful way to add custom rules for any JSON, YAML, Terraform, or Kubernetes configuration. Custom rules allow for checks against the unique requirements of a use case.

> 除了针对我们的内置规则集进行验证之外，config-lint 还提供了一种简单而强大的方法来为任何 JSON、YAML、Terraform 或 Kubernetes 配置添加自定义规则。自定义规则允许根据用例的独特要求进行检查。

You can install config-int on your Linux machine by executing:

您可以通过执行以下命令在 Linux 机器上安装 config-int：

```
curl -L https://github.com/stelligent/config-lint/releases/latest/download/config-lint_Linux_x86_64.tar.gz |tar xz -C /usr/local/bin config-lint
chmod +rx /usr/local/bin/config-lint
```


or use using Homebrew if you are using a macOS:

或者如果您使用的是 macOS，请使用 Homebrew：

```
brew tap stelligent/tap
brew install config-lint
```


## [Copper](https://github.com/cloud66-oss/copper)

## [铜](https://github.com/cloud66-oss/copper)

Copper, sponsored by Cloud 66, is a tool to validate your configuration files.

Copper 由 Cloud 66 赞助，是一种验证配置文件的工具。

According to the Copper development team, Copper's mission can be summarized in 2 main goals:

根据 Copper 开发团队的说法，Copper 的使命可以概括为两个主要目标：

1. Let developers make changes to Kubernetes configuration files as needed instead of restricting what they can do with another layer of APIs on top.
2. Make sure all configuration files applied to the Kubernetes clusters are tested and adhere to the relevant infrastructure policies.

1. 让开发人员根据需要对 Kubernetes 配置文件进行更改，而不是限制他们在顶部的另一层 API 上可以做什么。
2. 确保应用到 Kubernetes 集群的所有配置文件都经过测试并遵守相关的基础架构策略。

As described in the official documentation, these are some of the things Copper checks in configuration files committed into Git repository:

如官方文档中所述，以下是 Copper 在提交到 Git 存储库的配置文件中检查的一些内容：

- latest is not used as the image tag for any Deployment. 

- 最新不用作任何部署的图像标记。

- Image versions are not changed for important components (like databases) except for minor versions and patches.
- Load balancer IP address is not changed in Service configuration by mistake.
- Any fixed IP address used is within a valid range.
- No secret is committed into the configuration repository
- Certain images come from our trusted repositories and not public ones.

- 除了次要版本和补丁外，不会更改重要组件（如数据库）的映像版本。
- 负载均衡器 IP 地址未在服务配置中错误更改。
- 使用的任何固定 IP 地址都在有效范围内。
- 没有秘密提交到配置库
- 某些图像来自我们受信任的存储库，而不是公共存储库。

Copper can be used in automated CI/CD and GitOps pipelines using its DSL. For example, to check if your configuration is not including the latest image, you can write the following check:

Copper 可以使用其 DSL 用于自动化 CI/CD 和 GitOps 管道。例如，要检查您的配置是否不包括最新映像，您可以编写以下检查：

```
rule NoLatest ensure {  // use of latest as image tag is not allowed
fetch("$.spec.template.spec.containers..image")
.as(:image)
.pick(:tag)
.contains("latest") == false
}
```


Save this file to a .cop file and run Copper CLI:

将此文件保存到 .cop 文件并运行 Copper CLI：

```
$ copper check --rules my_rules.cop --file deploy.yml
```


## [Conftest](https://github.com/open-policy-agent/conftest/)

## [Conftest](https://github.com/open-policy-agent/conftest/)

Conftest helps developers write tests against structured configuration data.

Conftest 帮助开发人员针对结构化配置数据编写测试。

Using Conftest, you can write tests for your Kubernetes configuration, Tekton pipeline definitions, Terraform code, Serverless configs, or any other config files.

使用 Conftest，您可以为 Kubernetes 配置、Tekton 管道定义、Terraform 代码、无服务器配置或任何其他配置文件编写测试。

Here's a quick example from the documentation. Create a new policy file and save it under policy/deployment.rego:

这是文档中的一个快速示例。创建一个新的策略文件并将其保存在 policy/deployment.rego 下：

```
package main

deny[msg] {
input.kind == "Deployment"
not input.spec.template.spec.securityContext.runAsNonRoot

msg := "Containers must not run as root"
}

deny[msg] {
input.kind == "Deployment"
not input.spec.selector.matchLabels.app

msg := "Containers must provide app label for pod selectors"
}
```


Now run Conftest to test your Deployment file (deployment.yaml) like so:

现在运行 Conftest 来测试你的部署文件（deployment.yaml），如下所示：

```
$ conftest test deployment.yaml
FAIL - deployment.yaml - Containers must not run as root
FAIL - deployment.yaml - Containers must provide app label for pod selectors

2 tests, 0 passed, 0 warnings, 2 failures, 0 exceptions
```


Conftest supports multiple formats:

Conftest 支持多种格式：

- YAML
- JSON
- INI
- TOML
- HOCON
- HCL
- HCL 2
- CUE
- Dockerfile
- EDN
- VCL
- XML
- Jsonnet

- YAML
- JSON
- INI
- TOML
- 浩康
- 盐酸
- 盐酸 2
- 提示
- Dockerfile
- EDN
- VCL
- XML
- Jsonnet

## [Polaris](https://github.com/FairwindsOps/polaris)

## [北极星](https://github.com/FairwindsOps/polaris)

Polaris, developed by Fairwinds, runs a variety of checks to ensure that Kubernetes pods and controllers are configured using best practices, helping you avoid problems in the future.

由 Fairwinds 开发的 Polaris 运行各种检查以确保使用最佳实践配置 Kubernetes pod 和控制器，帮助您避免将来出现问题。

Developers can run Polaris in a few different modes:

开发人员可以在几种不同的模式下运行 Polaris：

- As a[dashboard](https://github.com/FairwindsOps/polaris#dashboard), so you can audit what's running inside your cluster.
- As a[validating webhook](https://github.com/FairwindsOps/polaris#webhook), so you can automatically reject workloads that don't adhere to your organization's policies.
- As a[command-line tool](https://github.com/FairwindsOps/polaris#cli), so you can test local YAML files, e.g. as part of a CI/CD process.

- 作为 [dashboard](https://github.com/FairwindsOps/polaris#dashboard)，您可以审核集群内运行的内容。
- 作为[验证 webhook](https://github.com/FairwindsOps/polaris#webhook)，您可以自动拒绝不符合组织策略的工作负载。
- 作为[命令行工具](https://github.com/FairwindsOps/polaris#cli)，您可以测试本地YAML文件，例如作为 CI/CD 过程的一部分。

To use this tool as a dashboard, simply run the following command then open http://localhost:8080 to see it up and running.

要将这个工具用作仪表板，只需运行以下命令，然后打开 http://localhost:8080 以查看它的启动和运行情况。

```
kubectl apply -f https://github.com/FairwindsOps/polaris/releases/latest/download/dashboard.yaml
kubectl port-forward --namespace polaris svc/polaris-dashboard 8080:80
```


As described in the Git repository, Polaris dashboard is a way to get a simple visual overview of the current state of your Kubernetes workloads as well as a roadmap for what can be improved. The dashboard provides a cluster wide overview as well as breaking out results by category, namespace, and workload.

正如 Git 存储库中所述，Polaris 仪表板是一种获取 Kubernetes 工作负载当前状态的简单直观概览以及可以改进的路线图的方法。仪表板提供集群范围的概览，并按类别、命名空间和工作负载细分结果。

[![Polaris Dashboard](https://static.thechief.io/prod/images/polaris-dashboard.width-1024.format-webp-lossless.webp)](https://static.thechief.io/prod/images/polaris-dashboard.width-1024.format-webp-lossless.webp)


/images/polaris-dashboard.width-1024.format-webp-lossless.webp)


Polaris Dashboard




北极星仪表板




Polaris supports a wide range of validations covering a number of Kubernetes best practices. Here's a [sample configuration file](https://github.com/FairwindsOps/polaris/blob/master/examples/config-full.yaml) that includes all currently supported checks. The [default configuration](https://github.com/FairwindsOps/polaris/blob/master/examples/config.yaml) contains a number of those checks.

Polaris 支持广泛的验证，涵盖了许多 Kubernetes 最佳实践。这是一个 [示例配置文件](https://github.com/FairwindsOps/polaris/blob/master/examples/config-full.yaml)，其中包括所有当前支持的检查。[默认配置](https://github.com/FairwindsOps/polaris/blob/master/examples/config.yaml) 包含许多这些检查。

Compared to the above tools, the dashboard is a major advantage that makes a difference.

与上述工具相比，仪表板是一个与众不同的主要优势。

## [KubeLibrary](https://github.com/devopsspiral/KubeLibrary/)

## [KubeLibrary](https://github.com/devopsspiral/KubeLibrary/)

KubeLibrary is different from the tools that we enumerated previously as it's mainly focused on testing but it can be helpful in many cases. Read on to know how. 

KubeLibrary 与我们之前列举的工具不同，因为它主要专注于测试，但在许多情况下它会有所帮助。请继续阅读以了解如何操作。

You can read more about this tool in an article that Nils Balkow-Tychsen (Lead QA Engineer at [Humanitec](https://humanitec.com/)) contributed to The Chief I/O: [KubeLibrary: Testing Kubernetes with RobotFramework] (https://thechief.io/c/humanitec/kubelibrary-testing-kubernetes-robotframework/)

您可以在 Nils Balkow-Tychsen（[Humanitec](https://humanitec.com/) 的首席 QA 工程师）为首席 I/O 贡献的文章中阅读有关此工具的更多信息：[KubeLibrary：使用 RobotFramework 测试 Kubernetes] （https://thechief.io/c/humanitec/kubelibrary-testing-kubernetes-robotframework/)

These are the most important points discussed in the article:

这些是文章中讨论的最重要的观点：

KubeLibrary is a wrapper for the Python Kubernetes Client. It enables you to assert the status of various objects in your Kubernetes Clusters.

KubeLibrary 是 Python Kubernetes 客户端的包装器。它使您能够断言 Kubernetes 集群中各种对象的状态。

As the library can be integrated with any RobotFramework test suite, it is ideal for verifying the testability of your System-under-Test by asserting the status of your nodes, deployments, pods, config maps, and others Kubernetes objects before running any end to end tests.

由于该库可以与任何 RobotFramework 测试套件集成，因此它非常适合通过在运行任何终端之前断言节点、部署、pod、配置映射和其他 Kubernetes 对象的状态来验证被测系统的可测试性。结束测试。

As KubeLibrary is based on the [official python kubernetes client](https://github.com/kubernetes-client/python/blob/master/kubernetes?ref=thechiefio), you can connect to your Kubernetes cluster while executing any Kubernetes API command.

由于 KubeLibrary 基于 [official python kubernetes client](https://github.com/kubernetes-client/python/blob/master/kubernetes?ref=thechiefio)，您可以在执行任何 Kubernetes API 的同时连接到您的 Kubernetes 集群命令。

Being part of the broader RobotFramework Library, all code is wrapped into keywords that can be used in test cases defined in ATDD (Acceptance Test Driven Development) or in BDD (Behavioral Driven Development) syntax.

作为更广泛的 RobotFramework 库的一部分，所有代码都封装在关键字中，这些关键字可用于以 ATDD（验收测试驱动开发）或 BDD（行为驱动开发）语法定义的测试用例。

There are [many different examples available within the GitHub repository of the KubeLibrary](https://github.com/devopsspiral/KubeLibrary/tree/master/testcases?ref=thechiefio). This is a quick example:

[KubeLibrary 的 GitHub 存储库中提供了许多不同的示例](https://github.com/devopsspiral/KubeLibrary/tree/master/testcases?ref=thechiefio)。这是一个快速示例：

Let’s say you want to make sure that all pods in a certain namespace are running and use a specific image version.

假设您想确保某个命名空间中的所有 pod 都在运行并使用特定的镜像版本。

```
# This is an example test case for the Robot Framework KubeLibrary
# https://github.com/devopsspiral/KubeLibrary

*** Settings ***
Library           KubeLibrary    None    True    False

*** Variables ***
${POD_NAME_PATTERN}       my-pod-name
${NAMESPACE}              my-namespace
${IMAGE_NAME}             my-image:1.0.0
${TIMEOUT}                2min
${RETRY_INTERVAL}         5sec

*** Test Cases ***
Pods are running with correct image
    Given waited for pods matching "${POD_NAME_PATTERN}" in namespace "${NAMESPACE}" to be running
    When getting pods matching "${POD_NAME_PATTERN}" in namespace "${NAMESPACE}"
    Then all pods containers are using "${IMAGE_NAME}" image

*** Keywords ***
waited for pods matching "${POD_NAME_PATTERN}" in namespace "${NAMESPACE}" to be running
    Wait Until Keyword Succeeds    ${TIMEOUT}    ${RETRY_INTERVAL}
    ...  pod "${POD_NAME_PATTERN}" status in namespace "${NAMESPACE}" is running

pod "${POD_NAME_PATTERN}" status in namespace "${NAMESPACE}" is running
    @{namespace_pods}=    Get Pod Names in Namespace  ${POD_NAME_PATTERN}    ${NAMESPACE}
    ${num_of_pods}=    Get Length    ${namespace_pods}
    Should Be True    ${num_of_pods} >= 1    No pods matching "${POD_NAME_PATTERN}" found
    FOR    ${pod}    IN    @{namespace_pods}
        ${status}=    Get Pod Status in Namespace    ${pod}    ${NAMESPACE}
        Should Be True     '${status}'=='Running'
    END

getting pods matching "${POD_NAME_PATTERN}" in namespace "${NAMESPACE}"
    @{namespace_pods}=    Get Pods in Namespace  ${POD_NAME_PATTERN}    ${NAMESPACE}
    Set Test Variable    ${namespace_pods}

all pods containers are using "${container_image}" image
    @{containers}=    Filter Pods Containers By Name    ${namespace_pods}    .*
    @{containers_images}=    Filter Containers Images    ${containers}
    FOR    ${item}    IN    @{containers_images}
        Should Be Equal As Strings    ${item}    ${container_image}
    END
```


If you run this test in your Kubernetes cluster, it would check for pods matching the name pattern in ${POD\_NAME\_PATTERN} in a namespace ${NAMESPACE}. It will search for running pods for 2 minutes. Once the pods are confirmed running it will confirm the image they are using.

如果您在 Kubernetes 集群中运行此测试，它将检查命名空间 ${NAMESPACE} 中 ${POD\_NAME\_PATTERN} 中的名称模式是否匹配的 Pod。它将搜索正在运行的 Pod 2 分钟。确认 Pod 正在运行后，它将确认它们正在使用的映像。

After you have confirmed that your pods are running you can continue with any application-level testing.

在确认 Pod 正在运行后，您可以继续进行任何应用程序级别的测试。

The KubeLibrary can be also used for checking nodes, jobs, config maps, persistent volume claims, services, and other Kubernetes objects. Current keyword documentation is available on GitHub ( [link](https://github.com/devopsspiral/KubeLibrary/blob/master/docs/KubeLibrary.html?ref=thechiefio)).

KubeLibrary 还可用于检查节点、作业、配置映射、持久卷声明、服务和其他 Kubernetes 对象。当前的关键字文档可在 GitHub 上找到（[link](https://github.com/devopsspiral/KubeLibrary/blob/master/docs/KubeLibrary.html?ref=thechiefio))。

If you want to learn more about KubeLibrary, you can join Humanitec's next [webinar](https://humanitec.com/webinars/test-automation-in-continuous-deployment?ref=thechiefio)! 

如果您想进一步了解 KubeLibrary，可以加入 Humanitec 的下一次 [网络研讨会](https://humanitec.com/webinars/test-automation-in-continuous-deployment?ref=thechiefio)！

