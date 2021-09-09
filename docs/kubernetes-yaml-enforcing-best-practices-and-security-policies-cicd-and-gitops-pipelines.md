# Kubernetes YAML: Enforcing best practices and security policies in CI/CD and GitOps pipelines

Using [Kubernetes](https://thechief.io/c/cloudplex/how-get-started-kubernetes/) is a synonym for manipulating YAML.

Even if the YAML community describes it as a human-readable language, it can sometimes be tricky to read it, especially in the context of Kubernetes when you are manipulating complex deployments, services, ingresses, and other resources. This may lead to [incoherences and problems](https://thechief.io/c/editorial/4-kubernetes-failure-stories-learn/) that may take some time to solve.

* * *

Before applying your YAML configuration to a [Kubernetes](https://thechief.io/c/cloudplex/21-resources-and-tutorials-learn-kubernetes/) cluster, it is evident and crucial to check if it's valid. Reading through your YAML configuration is the obvious way to do it. However, it's neither the fastest nor the easiest way to do it. On the other hand, if you are automating your deployments or testing, you will need to implement this step as part of your automated workflows.

Fortunately, the [Kubernetes](https://thechief.io/c/cloudplex/kubernetes-busy-developer/) and [Cloud Native](https://thechief.io/c/editorial/the-cios-guide-to-kubernetes-and-eventual-transition-to-cloud-native-development/) developer community has developed (and still) tools that make validating easy and especially automated. These tools can be extremely helpful when implemented in [CI/CD](https://thechief.io/c/cloudplex/top-10-kubernetes-cicd-tools/) or continuous GitOps pipelines. In addition to that, these tools help you in enforcing best practices, as well as applying policies and compliance requirements.

Let's discover some of these tools.

## [KubeLinter](https://github.com/stackrox/kube-linter)

KubeLinter offers the automated analysis of [Kubernetes](https://thechief.io/c/cloudplex/kubernetes-101/) YAML files and HELM charts before deployment.

It integrates [security-as-a-code](https://thechief.io/c/news/stackrox-announces-release-kubelinter/) in DevOps and other related processes. It helps ensure that the [Kubernetes](https://thechief.io/c/metricfire/cicd-pipelines-kubernetes-applications/) configuration is correct and the automatic enforcement of [security](https://thechief.io/c/editorial/takeaways-state-containers-and-kubernetes-security/) policies for Kubernetes applications.

Users can build [security](https://thechief.io/c/news/kubernetes-security-specialist-certification-now-available/) in the configuration as code in the application development process itself. It helps validate that the [Kubernetes](https://thechief.io/c/metricfire/aws-ecs-vs-kubernetes/) configuration is according to the security best practices.

KubeLinter users can integrate this tool to automate the process of carrying out configuration checks and error identification.

## [Kubeval](https://github.com/instrumenta/kubeval/)

Kubeval is another tool to validate your [Kubernetes](https://thechief.io/c/metricfire/kubernetes-networking-101/) configuration files. It supports multiple [Kubernetes](https://thechief.io/c/news/challenges-enterprise-kubernetes/) versions and works with YAML or JSON [configuration](https://thechief.io/c/news/red-hat-announces-integration-ansible-automation-openshift-kubernetes/) files.

This tool can be used in [CI/CD pipelines](https://thechief.io/c/metricfire/cicd-pipelines-kubernetes-applications/) as well as local development and [testing](https://thechief.io/c/cloudplex/kubernetes-distributed-performance-testing-using-locust/). It has the advantage of being a single binary that you can [download](https://www.kubeval.com/installation/) and run smoothly without any installation or configuration effort.

For a Linux machine, the installation is as easy as:

```
wget https://github.com/instrumenta/kubeval/releases/latest/download/kubeval-linux-amd64.tar.gz
tar xf kubeval-linux-amd64.tar.gz
sudo cp kubeval /usr/local/bin
```

## [kube-score](https://github.com/zegl/kube-score)

kube-score is a tool that performs static code analysis of your [Kubernetes](https://thechief.io/c/abvijaykumar/kubernetes-operators-realize-dream-zero-touch-ops/) object definitions.

The output is a list of recommendations of what you can improve to make your application more secure and resilient.

For the following input:

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

```
apps/v1/Deployment statefulset-test-1                                         💥
    [CRITICAL] Container Resources
        · foobar -> CPU limit is not set
            Resource limits are recommended to avoid resource DDOS. Set resources.limits.cpu
        · foobar -> Memory limit is not set
            Resource limits are recommended to avoid resource DDOS. Set resources.limits.memory
        · foobar -> CPU request is not set
            Resource requests are recommended to make sure that the application can start and run without
            crashing. Set resources.requests.cpu
        · foobar -> Memory request is not set
            Resource requests are recommended to make sure that the application can start and run without
            crashing. Set resources.requests.memory
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
            scheduled on the same node. This increases availability in case the node becomes unavailable.
policy/v1beta1/PodDisruptionBudget app-budget                                 ✅
```

## [**config-lint**](https://github.com/stelligent/config-lint)

This is a command-line tool to validate configuration files using rules specified in YAML. The configuration files can be one of several formats: [Terraform](https://thechief.io/c/codersociety/creating-and-manageing-ha-aks-kubernetes-cluster-azure-thanks-terraform/), JSON, YAML, with support for Kubernetes.

There are built-in rules provided for Terraform, and custom files can be used for other formats.

This tool, developed by Stelligent, has the advantage of being customizable since you can develop custom rules. In a blog post explaining how to use this tool, Dan Miller states:

> In addition to validating against our built-in ruleset, config-lint provides a simple and powerful way to add custom rules for any JSON, YAML, Terraform, or Kubernetes configuration. Custom rules allow for checks against the unique requirements of a use case.

You can install config-int on your Linux machine by executing:

```
curl -L https://github.com/stelligent/config-lint/releases/latest/download/config-lint_Linux_x86_64.tar.gz | tar xz -C /usr/local/bin config-lint
chmod +rx /usr/local/bin/config-lint
```

or use using Homebrew if you are using a macOS:

```
brew tap stelligent/tap
brew install config-lint
```

## [Copper](https://github.com/cloud66-oss/copper)

Copper, sponsored by Cloud 66, is a tool to validate your configuration files.

According to the Copper development team, Copper's mission can be summarized in 2 main goals:

1. Let developers make changes to Kubernetes configuration files as needed instead of restricting what they can do with another layer of APIs on top.
2. Make sure all configuration files applied to the Kubernetes clusters are tested and adhere to the relevant infrastructure policies.

As described in the official documentation, these are some of the things Copper checks in configuration files committed into Git repository:

- latest is not used as the image tag for any Deployment.
- Image versions are not changed for important components (like databases) except for minor versions and patches.
- Load balancer IP address is not changed in Service configuration by mistake.
- Any fixed IP address used is within a valid range.
- No secret is committed into the configuration repository
- Certain images come from our trusted repositories and not public ones.

Copper can be used in automated CI/CD and GitOps pipelines using its DSL. For example, to check if your configuration is not including the latest image, you can write the following check:

```
rule NoLatest ensure {  // use of latest as image tag is not allowed
fetch("$.spec.template.spec.containers..image")
.as(:image)
.pick(:tag)
.contains("latest") == false
}
```

Save this file to a .cop file and run Copper CLI:

```
$ copper check --rules my_rules.cop --file deploy.yml
```

## [Conftest](https://github.com/open-policy-agent/conftest/)

Conftest helps developers write tests against structured configuration data.

Using Conftest, you can write tests for your Kubernetes configuration, Tekton pipeline definitions, Terraform code, Serverless configs, or any other config files.

Here's a quick example from the documentation. Create a new policy file and save it under policy/deployment.rego:

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

```
$ conftest test deployment.yaml
FAIL - deployment.yaml - Containers must not run as root
FAIL - deployment.yaml - Containers must provide app label for pod selectors

2 tests, 0 passed, 0 warnings, 2 failures, 0 exceptions
```

Conftest supports multiple formats:

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

## [Polaris](https://github.com/FairwindsOps/polaris)

Polaris, developed by Fairwinds, runs a variety of checks to ensure that Kubernetes pods and controllers are configured using best practices, helping you avoid problems in the future.

Developers can run Polaris in a few different modes:

- As a[dashboard](https://github.com/FairwindsOps/polaris#dashboard), so you can audit what's running inside your cluster.
- As a[validating webhook](https://github.com/FairwindsOps/polaris#webhook), so you can automatically reject workloads that don't adhere to your organization's policies.
- As a[command-line tool](https://github.com/FairwindsOps/polaris#cli), so you can test local YAML files, e.g. as part of a CI/CD process.

To use this tool as a dashboard, simply run the following command then open http://localhost:8080 to see it up and running.

```
kubectl apply -f https://github.com/FairwindsOps/polaris/releases/latest/download/dashboard.yaml
kubectl port-forward --namespace polaris svc/polaris-dashboard 8080:80
```

As described in the Git repository, Polaris dashboard is a way to get a simple visual overview of the current state of your Kubernetes workloads as well as a roadmap for what can be improved. The dashboard provides a cluster wide overview as well as breaking out results by category, namespace, and workload.

[![Polaris Dashboard](https://static.thechief.io/prod/images/polaris-dashboard.width-1024.format-webp-lossless.webp)](https://static.thechief.io/prod/images/polaris-dashboard.width-1024.format-webp-lossless.webp)


Polaris Dashboard




Polaris supports a wide range of validations covering a number of Kubernetes best practices. Here's a [sample configuration file](https://github.com/FairwindsOps/polaris/blob/master/examples/config-full.yaml) that includes all currently supported checks. The [default configuration](https://github.com/FairwindsOps/polaris/blob/master/examples/config.yaml) contains a number of those checks.

Compared to the above tools, the dashboard is a major advantage that makes a difference.

## [KubeLibrary](https://github.com/devopsspiral/KubeLibrary/)

KubeLibrary is different from the tools that we enumerated previously as it's mainly focused on testing but it can be helpful in many cases. Read on to know how.

You can read more about this tool in an article that Nils Balkow-Tychsen (Lead QA Engineer at [Humanitec](https://humanitec.com/)) contributed to The Chief I/O: [KubeLibrary: Testing Kubernetes with RobotFramework](https://thechief.io/c/humanitec/kubelibrary-testing-kubernetes-robotframework/)

These are the most important points discussed in the article:

KubeLibrary is a wrapper for the Python Kubernetes Client. It enables you to assert the status of various objects in your Kubernetes Clusters.

As the library can be integrated with any RobotFramework test suite, it is ideal for verifying the testability of your System-under-Test by asserting the status of your nodes, deployments, pods, config maps, and others Kubernetes objects before running any end to end tests.

As KubeLibrary is based on the [official python kubernetes client](https://github.com/kubernetes-client/python/blob/master/kubernetes?ref=thechiefio), you can connect to your Kubernetes cluster while executing any Kubernetes API command.

Being part of the broader RobotFramework Library, all code is wrapped into keywords that can be used in test cases defined in ATDD (Acceptance Test Driven Development) or in BDD (Behavioral Driven Development) syntax.

There are [many different examples available within the GitHub repository of the KubeLibrary](https://github.com/devopsspiral/KubeLibrary/tree/master/testcases?ref=thechiefio). This is a quick example:

Let’s say you want to make sure that all pods in a certain namespace are running and use a specific image version.

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

After you have confirmed that your pods are running you can continue with any application-level testing.

The KubeLibrary can be also used for checking nodes, jobs, config maps, persistent volume claims, services, and other Kubernetes objects. Current keyword documentation is available on GitHub ( [link](https://github.com/devopsspiral/KubeLibrary/blob/master/docs/KubeLibrary.html?ref=thechiefio)).

If you want to learn more about KubeLibrary, you can join Humanitec's next [webinar](https://humanitec.com/webinars/test-automation-in-continuous-deployment?ref=thechiefio)!
