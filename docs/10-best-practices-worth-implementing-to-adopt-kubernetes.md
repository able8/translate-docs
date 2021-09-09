# 10 Best Practices Worth Implementing to Adopt Kubernetes

September 25, 2020

We already know that [Kubernetes](https://kubernetes.io/) is the No. 1 orchestration platform for container-based applications, automating the deployment and scaling of these apps and streamlining maintenance operations. However, Kubernetes comes with its own complexity challenges. So how can an enterprise take advantage of containerization to tackle complexity and not end up with even more complexity? This article provides some of the best practices that you can implement to adopt Kubernetes.

## Keep a Tab on Policies

Define appropriate policies for cluster access controls, service access controls, resource utilization controls and secret access controls. By default, containers run with unbounded compute resources on a Kubernetes cluster. To limit or restrict you must implement appropriate policies.

- Use_NetworkPolicy_ resources labels to select pods and define rules that specify what traffic is allowed to the selected pods.
- Kubernetes scheduler has default limits on the number of volumes that can be attached to a Node. To define the maximum number of volumes that can be attached to a Node for various cloud providers, use_Node-specific Volume Limits_.
- To enforce constraints on resource usage, use_Limit Range option_ for appropriate resource in the namespace.
- To limit aggregate resource consumption per namespace, use the below Resource Quotas:
  - Compute Resource Quota
  - Storage Resource Quota
  - Object Count Quota
  - Limits the number of resources based on scope defined in Quota Scopes option
  - Requests vs Limits – Each container can specify a request and a limit value for either CPU or memory
  - Quota and cluster capacity – Expressed in absolute units
  - Limit Priority Class consumption by default – For example, restrict usage of certain high-priority pods
- To allow/deny fine-grained permissions, use RBAC (role-based access control) and rules can be defined to allow/deny fine-grained permissions.
- To define and control security aspects of Pods, use Pod Security Policy (available from v1.15) to enable fine-grained authorization of pod creation and updates.
  - Running of privileged containers
  - Usage of host namespaces
  - Usage of host networking and ports
  - Usage of volume types
  - Usage of the host filesystem
  - Restricting escalation to root privileges
  - The user and group IDs of the container
  - AppArmor or seccomp or sysctl profile used by containers
- Use any of the tools such as[Open Policy Agent Gatekeeper](https://www.upnxtblog.com/index.php/2019/12/09/implementing-policies-in-kubernetes/) policy engine to manage, author the policies.

## Manage Resources Wisely

Use resource utilization (resource quota) guidelines to ensure the containerized applications co-exist without being eliminated due to resource violations at runtime. To enforce constraints on resource usage, use _Limit Range option_ for appropriate resources in the namespace.

To limit aggregate resource consumption per namespace, use the below Resource Quotas:

- Compute Resource Quota
- Storage Resource Quota
- Object Count Quota
- Limits the number of resources based on scope defined in Quota Scopes option
- Requests vs Limits – Each container can specify a request and a limit value for either CPU or memory.
- Quota and cluster capacity – Expressed in absolute units
- Limit Priority Class consumption by default – For example, restrict usage of certain high priority pods

## Focus on Comprehensive Observability of the Cluster

Currently, the Kubernetes ecosystem provides two add-ons for aggregating and reporting monitoring data from your cluster: **(1) Metrics Server and (2) kube-state-metrics.**

[**Metrics**](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/metrics-server.md) **Server is** a cluster add-on that collects resource usage data from each node and provides aggregated metrics through [the Metrics API](https://github.com/kubernetes/metrics). **kube-state-metrics** service provides additional cluster information that Metrics Server does not.

Below are the key metrics and alerts that are required to monitor your Kubernetes cluster.

**What to monitor?****Metrics to monitor****Alert Criteria****Cluster state**Monitor the aggregated resources usage across all nodes in your cluster.

- Node status
- Desired pods
- Current pods
- Available pods
- Unavailable pods

- Node status
- Desired vs. current pods
- Available and unavailable pods

**Node resources**For each of the node monitor :

- Memory requests
- Memory limits
- Allocatable memory
- Memory utilization
- CPU requests
- CPU limits
- Allocatable CPU
- CPU utilization
- Disk utilization

If the node’s CPU or memory usage drops below a desired threshold.

- Memory limits per pod vs. memory utilization per pod
- Memory utilization
- Memory requests per node vs. allocatable memory per node
- Disk utilization
- CPU requests per node vs. allocatable CPU per node
- CPU limits per pod vs. CPU utilization per pod
- CPU utilization

**Missing pod**Health and availability of your pod deployments.

- Available pods
- Unavailable pods

If the number of available pods for a deployment falls below the number of pods you specified when you created the deployment.**Pods that are not running**If a pod isn’t running or even scheduled, there could be an issue with either the pod or the cluster, or with your entire Kubernetes deployment.

- Pod status

Alerts should be based on the status of your pods (“Failed,” ”Pending,” or “Unknown” for the period of time you specify).**Container restarts**Container restarts could happen when you’re hitting a memory limit (ex.Out of Memory kills) in your containers.

Also, there could be an issue with either the container itself or its host.

Kubernetes automatically restarts containers,  but setting up an alert will give you an immediate notification later you can analyze and set the proper limits.**Container resource usage**Monitor container resource usage for containers in case you’re hitting resource limits, spikes in resource consumption, Alerts to check if container CPU and memory usage and on limits are based on thresholds.**Storage volumes**Monitor storage to:

- Ensure your application has enough disk space so pods don’t run out of space
- Volume usage and adjust either the amount of data generated by the application or the size of the volume according to usage

Alerts to check if available bytes, capacity crosses your thresholds.

Identify persistent volumes and apply a different alert threshold or notification for these volumes, which likely hold important application data.

**Control Plane – Etcd**Monitor etcd for the below parameters:

- Leader existence and change rate
- Committed, applied, pending, and failed proposals
- gRPC performance

Alerts to check if any pending or failed proposals or reach inappropriate thresholds.**Control Plane – API Server**Monitor the API server for below parameters:

- Rate / number of HTTP requests
- Rate/number of apiserver requests

Alerts to check if the rate or number of HTTP requests crosses a desired threshold.**Control Plane – Scheduler**Monitor the scheduler for the below parameters:

- Rate, number, and latency of HTTP requests
- Scheduling latency
- Scheduling attempts by result
- End-to-end scheduling latency (sum of scheduling)

Alerts to check if the rate or number of HTTP requests crosses a desired threshold.**Control Plane – Controller Manager**Monitor the scheduler for the below parameters:

- Work queue depth
- Number of retries handled by the work queue

Alerts to check if requests to the work queue exceed a maximum threshold.**Kubernetes Events**Collecting events from Kubernetes and from the container engine (such as Docker) allows you to see how pod creation, destruction, starting or stopping affects the performance of your infrastructure.Any failure or exception should need to be alerted.

Consider integrating with any of the commercial monitoring tools to consume probe-generated metrics and platform-generated metrics to have comprehensive observability of the cluster.

## Container Security Management Must Be Part of Your DevOps Pipeline

Continuous security must be included as part of the DevOps pipeline to ensure containers are well-managed. Use any of the below static analysis tools to identify vulnerabilities in application containers while building images for containers:

- [Clair](https://github.com/quay/clair)
- [Trivy](https://github.com/aquasecurity/trivy)
- [kube-bench](https://github.com/aquasecurity/kube-bench)
- [Falco](https://github.com/falcosecurity/falco)
- [Notary](https://github.com/theupdateframework/notary)

## Audit and Compliance Your Cluster Routinely

Routinely audit the platform for Kubernetes patch levels, secret stores, compliance against the security vulnerabilities, encryption of secret stores, storage volumes, cluster policies, role binding policies, RBAC and user management controls.

## Chaos Test Your Cluster

Proactively chaos tests your platform to ensure the robustness of the cluster. It also helps to test the stability of the containerized applications and the impact of crashing these containers. A wide range of open source and commercial tools can be used, some of which are listed below:

- [Chaosblade](https://github.com/chaosblade-io/chaosblade)
- [Chaos Mesh](https://github.com/pingcap/chaos-mesh)
- [PowerfulSeal](https://github.com/bloomberg/powerfulseal)
- [chaoskube](https://github.com/linki/chaoskube)
- [Chaos Toolkit](https://github.com/chaostoolkit/chaostoolkit)
- [Litmus](https://github.com/litmuschaos/litmus)

## Archive and Back Up Your Cluster

Kubernetes uses etcd as its internal metadata management store to manage the objects across clusters. It is necessary to define a backup strategy for etcd and any other dependent persistent stores used within the Kubernetes clusters.

Use [Velero](https://www.upnxtblog.com/index.php/2019/12/16/how-to-back-up-and-restore-your-kubernetes-cluster-resources-and-persistent-volumes/) or any other open source tools to backup **Kubernetes resources** and **application data** so that in cases of **disaster** it can reduce recovery time.

## Manage Your Deployment Manifests

Kubernetes follows declaration-based management; hence, every object or resource or instruction is described only through YAML declarative manifests. It is necessary to leverage SCM tools or create custom utilities to manage these manifests.

## Continuous Deployment of Services

kubectl style of deployments would not be possible in a large-scale production setup. Instead, you have to use some of the established open source frameworks. [**Helm**](https://helm.sh/), for example, is specifically built for Kubernetes to manage seamless deployments via the CI/CD pipeline.

Helm uses _Charts_ that define the set of Kubernetes resources that together define an application. You can think of Charts as packages of pre-configured Kubernetes resources. Charts help you to define, install and upgrade even the most complex Kubernetes application. These charts can describe a single resource, such as a Redis pod or a full-stack of a web application: HTTP servers, databases and caches.

In the recent release of Helm, releases will be managed inside of Kubernetes using [Release Objects](https://helm.sh/docs/chart_template_guide/builtin_objects/) and Kubernetes Secrets. All modifications such as installing, upgrading and downgrading releases will end in having a new version of that Kubernetes Secret.

## Use Service Mesh

[Service mesh](https://www.upnxtblog.com/index.php/2018/12/17/what-is-service-mesh-why-do-we-need-it-linkered-tutorial/) offers consistent discovery, security, tracing, monitoring and failure handling without the need for a shared asset such as an API gateway. So, if you have service mesh on your cluster, you can achieve all the below items without making changes to your application code:

- Automatic load balancing
- Fine-grained control of traffic behavior with routing rules, retries, failovers, etc.
- Pluggable policy layer
- Configuration API supporting access controls, rate limits and quotas
- Service discovery
- Service monitoring with automatic metrics, logs, and traces for all traffic
- Secure service to service communication

Currently, service mesh is being offered by [Linkerd](https://github.com/linkerd/linkerd2), [Istio](https://github.com/istio/istio) and [Conduit](http://www.conduit.io/) providers.

It is necessary to choose an appropriate service mesh that is compatible with the Kubernetes cluster as well as the underlying infrastructure.

## Conclusion

This article covers the key best practices that you can implement for Kubernetes adoption. However, operating Kubernetes clusters is not without its challenges.

- [Click to share on Twitter (Opens in new window)](https://containerjournal.com/topics/container-management/10-best-practices-worth-implementing-to-adopt-kubernetes/?share=twitter "Click to share on Twitter")
- [Click to share on Facebook (Opens in new window)](https://containerjournal.com/topics/container-management/10-best-practices-worth-implementing-to-adopt-kubernetes/?share=facebook "Click to share on Facebook")
- [Click to share on LinkedIn (Opens in new window)](https://containerjournal.com/topics/container-management/10-best-practices-worth-implementing-to-adopt-kubernetes/?share=linkedin "Click to share on LinkedIn")
- [Click to share on Reddit (Opens in new window)](https://containerjournal.com/topics/container-management/10-best-practices-worth-implementing-to-adopt-kubernetes/?share=reddit "Click to share on Reddit")

### _Related_

- [← 2nd Watch Adds Hybrid Anthos Practice to IT Services Portfolio](https://containerjournal.com/topics/container-management/2nd-watch-adds-hybrid-anthos-practice-to-it-services-portfolio/)
- [Microsoft to Bring AKS Service to HCI Platforms →](https://containerjournal.com/topics/container-management/microsoft-to-bring-aks-service-to-hci-platforms/)
