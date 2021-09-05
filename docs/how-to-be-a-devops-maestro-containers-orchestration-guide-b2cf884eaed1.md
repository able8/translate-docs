# How to be a DevOps maestro: containers orchestration guide

“All things are difficult before they are easy.” — Thomas Fuller

## Introduction

Nowadays software packaged as container images and ran in containers is an  industry-accepted standard of running and distributing applications.  Some of the benefits of using containers are:

- A high degree of portability
- Consistent operations processes
- Scalability and resiliency mostly
- Enabling common tooling thanks to open standards

As with everything, there are no free lunches! The power that containers  provide comes with complexity. Running standalone containers instances  can be ok for limited size workloads or at a smaller scale. At a certain point, the complexity is too much to deal with manually.

Developing and running containerized workloads at scale requires a high degree of  automation and operational excellence. In other words, it requires **containers orchestration.**

After reading this article you will:

- understand what containers orchestration is
- why is it a critical part of running containers
- how can it help make your life easier

It can also help you explain to your peers or leadership why and when  containers orchestration is important. You will also learn about useful  tools and patterns.

## What is containers orchestration?

Traditionally, containers orchestration is associated mostly with operational level  activities. There are however lots of activities happening parallelly or before expanded to containers orchestration. By including development  as well as building and deployment automation related tasks a new  definition can be applied: **Containerized workloads management.**

> Containerized workloads management is a process of automating common tasks throughout the whole lifecycle of container images and containers. Containers  orhestration is is a process of automating common operational-level  tasks.

The below diagram shows how **containerized workloads management** involves development, security, operations and other disciplines in the following areas:

![img](https://miro.medium.com/max/2000/1*HFSSx9WG3GXu8cozP75jMA.png)

Source: Author

In this article, we will dig deeper into the “traditional” container orchestration side of Operations.

## Example Scenario

A sample scenario will help illustrate and understand all the principles.

Imagine that you need to create and deploy a traditional n-tier app with a web  frontend, simple REST API, and document database as a persistence store. The setup will run on-prem and in a public cloud. We expect increased  traffic over time with a steady usage pattern. At some point our app  will undergo internal compliance and security audits, resulting in  additional requirements.

## What you see is not all there is to it

A typical tutorial or blog post showing how to get started with running  something in a container barely scratches the surface. Once a certain  scale is reached, orchestration becomes critical and unfortunately, most of it is not visible to the business.

![img](https://miro.medium.com/max/2000/1*stRWa92BLo-6wplO46VbLw.jpeg)

Image by [Josep Monter Martinez](https://pixabay.com/users/josepmonter-1007570/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1321692) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1321692)

Let’s dig deeper into each task from the “Operations” area. The goal is to  show what problems containers orchestration tries to solve. We will look at patterns and paradigms as well as tools that help address different  areas of orchestration.

> **Provision and Deploy**

Without tools, **provisioning** would typically mean sshing into a virtual machine and installing a  bunch of tools to get the nodes ready for receiving workloads. In our  example, it means a separate server for the frontend app, a separate one for middleware and another one for the database.

**A deployment** is an act of pulling images on the server and spinning up new  containers. Let’s assume that our application runs one container for  each component, so 3 containers in total. Without orchestration doing so would mean setting up a connection to a remote docker host and issuing  imperative commands like `docker run etc.`

How does this step look like with container orchestration tools and processes?

To enable smart **provisioning** we could take advantage of two patterns:

- [Infrastructure as Code](https://en.wikipedia.org/wiki/Infrastructure_as_code)
- [GitOps](https://github.com/gitops-working-group/gitops-working-group#gitops-principles)
- [Declarative Programming](https://en.wikipedia.org/wiki/Declarative_programming)

Many tools work great in this space. Highlighting a few here.

[Terraform by HashiCorpBlog post Read the 1.0 launch blog post-Terraform is an open-source infrastructure as a code software tool that provides…www.terraform.io](https://www.terraform.io/)

[CrossplaneCrossplane brings Kubernetes-styled declarative and API-driven configuration and management to any piece of…crossplane.io](https://crossplane.io/)

[Pulumi - Modern Infrastructure as CodeAll architectures welcome Choose from over 50 cloud providers, including public, private, and hybrid architectures…www.pulumi.com](https://www.pulumi.com/)

Deployment should be fully automated, especially at scale. A pattern to use here  would also be GitOps. Example tools in this space are:

[FluxFlux enables continuous delivery of container images, using version control for each step to ensure deployment is…www.weave.works](https://www.weave.works/oss/flux/)

[Argo CD - Declarative GitOps CD for KubernetesArgo CD is a declarative, GitOps continuous delivery tool for Kubernetes. Application definitions, configurations, and…argoproj.github.io](https://argoproj.github.io/argo-cd/)

> Schedule on a correct node

Our workloads are containerized and ready, we can deploy them, but wait  where exactly? In a high scale environment, there are typically multiple compute nodes, moreover, the health of the nodes and available  resources can change over time. Without any tooling, we would have to  make decisions based on resources utilization, network latency, OS  version, etc. Due to the manual effort required, this decision would  typically be done once, not taking advantage of [bin packing](https://en.wikipedia.org/wiki/Bin_packing_problem).

Kubernetes comes with a build-in scheduler controller, that takes care of exactly  this. Not only the scheduling is done once but is also evaluated and  adjusted dynamically.

Tools that help in this area: Kubernetes

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

> Allocate Resources

Each container needs compute resources, CPU, RAM, storage. Without  orchestration, there is no easy way of limiting how much resources a  container can consume. We are basically left with system monitoring  tools and retroactively fixing issues.

Kubernetes comes with a build-in mechanism to declare how much resources should be allocated to containers (requested resources). It also supports  resources limits that can be consumed by containers resources limits)

Tools that help in this area: Kubernetes

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

> Conditionally auto-scale

Our app is a success and utilization is growing. Thankfully we have followed [The Twelve-Factor App](https://12factor.net/) design methodology and our frontend and middleware can [horizontally scale](https://en.wikipedia.org/wiki/Scalability#Horizontal_(scale_out)_and_vertical_scaling_(scale_up)) without a problem.

Without orchestration, horizontal scaling of our frontend app in its simplest form involves:

- spinning up a new container
- creating a load balancer either as another container or a standalone installation (for example nginx)
- updating nginx configuration pointing to new IP addresses of the containers to the load balancer
- repeat if we want to scale up to additional containers

Of course, if the demand drops, we should remove excessive containers. With a big enough scale, this is a full-time job!

Kubernetes provides a build-in mechanism called [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) for scaling up and down based on resources utilization. What if we  wanted to scale based on completely arbitrary metrics or API calls? No  problem, KEDA provides a rich event-driven auto-scaling mechanism.

What if our database needs a more “beefy” server/node to support all the  requests coming from the middleware? Kubernetes got us covered, we can  create a new node pool for our database workloads and use a combination  of [taints and tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) to automatically migrate our database pods to a better-suited node.

Tools that help in this area: Kubernetes, KEDA

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[KEDAKEDA is a Kubernetes-based Event Driven Autoscaler. With KEDA, you can drive the scaling of any container in Kubernetes…keda.sh](https://keda.sh/)

> Load balance and route traffic

Let’s imagine, that we have added a new route to our frontend app and need to load balance it from publicly facing domain to appropriate pod and  container. Without tooling, it involves a manual reconfiguration of load balancer which can get very messy very fast.

Kubernetes provides native primitives (services, ingresses) to helps us automate load balancing tasks.

Although the DNS plugin is not part of “core” Kubernetes, it is almost always  installed on a cluster. DNS provides service discovery, no more fiddling with IP addresses in config files, localhost is enough!

Tools that help in this area: Kubernetes with CoreDNS plugin

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[CoreDNS: DNS and Service DiscoveryCoreDNS chains plugins. Each plugin performs a DNS function, such as Kubernetes service discovery, prometheus metrics…coredns.io](https://coredns.io/)

> Ensure high availability

Mission-critical systems, such as our example app, must adhere to SLAs (Service Level  Agreements) defining the uptime of a service. [High availability](https://en.wikipedia.org/wiki/High_availability) means ensuring that the system must be up and available for a specified duration of time. For example “four nines” of availability (99.99% or  time) translates to 52 minutes per year or 4.38 minutes per month of  unavailability.

Without any tools, it is very hard to ensure high availability.

Kubernetes comes with a powerful mechanism of “reconciliation loops” which are  controllers running on the cluster and constantly monitoring the  infrastructure and applications for any drifts from the desired state.  Once such drift is detected, Kubernetes will automatically try to  restore balance by making sure desired state corresponds to the actual  state.

There are two paradigms we can take advantage of here:

- [GitOps](https://github.com/gitops-working-group/gitops-working-group#gitops-principles)
- [Declarative Programming](https://en.wikipedia.org/wiki/Declarative_programming)

Tools that help in this area: Kubernetes

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

> Monitor and health check

Observability is a great example of a decentralized system design with a centralized  control plane. To observe our containers we will need highly  decentralized components operating on various levels of the stack  (kernel syscalls, system-wide events, application-specific monitoring  endpoints, etc).

All those pieces of data are collected and presented in a meaningful way  either to humans or automated processes. Kubernetes can enable  cluster-wide [audit policy](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/#audit-policy) and collect data on various levels.

There are a lot of tools that can support this process, such as Prometheus, Graphana, Loki, FluentD, Jaeger, just to name a few.

To enable monitoring on an application level, Kubernetes comes with a concept of [liveness, readiness and startup probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).

Cillium and Falco deserve a special mention. These tools can work with [eBPF](https://ebpf.io/) probes to provide deep level observability. Falco specializes in  security-related monitoring whereas Cillium does the same for  networking.

Tools that help in this area: Kubernetes, Prometheus, Jaeger, Grafana, Fluentd, Cillium, Falco

[FalcoStrengthen container security The flexible rules engine allows you to describe any type of host or container behavior…falco.org](https://falco.org/)

[Cilium - Linux Native, API-Aware Networking and Security for ContainersTraditional firewalls limit their inspection to the IP and TCP layers. Cilium uses eBPF to accelerate getting data in…cilium.io](https://cilium.io/)

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[Prometheus - Monitoring system & time series databasePower your metrics and alerting with a leadingopen-source monitoring solution. © Prometheus Authors 2014-2021 |…prometheus.io](https://prometheus.io/)

[Jaeger: open source, end-to-end distributed tracingMonitor and troubleshoot transactions in complex distributed systemswww.jaegertracing.io](https://www.jaegertracing.io/)

[Grafana: The open observability platformGrafana is the open source analytics & monitoring solution for every database.grafana.com](https://grafana.com/)

[Fluentd | Open Source Data CollectorFluentd is an open source data collector for unified logging layer.www.fluentd.org](https://www.fluentd.org/)

> Configure

There are many ways of providing configuration to containerized workloads.  Reading environmental variables, config files, external endpoints such  as databases or HTTP/gRPC services etc.

Kubernetes provides consistent configuration management via the following primitives:

- Config Maps
- Secrets
- Volumes

All those configuration options can be “mounted” to pods in various ways;  providing a consistent model of configuration management.

Tools that help in this area: Kubernetes

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

> Security and secure communication

Last, but definitely not least, how to ensure secure communication between  containers? Kubernetes without any additional settings or changes is  very permissive. Thankfully, there are tools and standards to help us  harden and secure the cluster.

Let’s imagine that our application is being audited and a security  recommendation was made to ensure that our database container cannot  talk to the frontend container. As long as we don’t scale anything, it  is possible to do with docker-compose and networking settings, but it is a manual process.

Kubernetes comes with a built-in mechanism of [network policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/) that regulate which pods can communicate with each other.

This is great, but after a while, another team deployed a service to the  cluster. This service provides sensitive lookup data, we have to ensure  that the communication between our pods and the service is encrypted.

Here we can use tools from the Service Mesh category. Those tools provide  enhanced observability as well as secure communication capabilities.  Istio and LinkerD for example.

Maybe we want to prevent pods from using the default public Docker Hub registry for docker images? Kubernetes comes with [admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) and Image Policy Webhook is one of them. If we want to create more  fine-grained policies, tools such as Open Policy Agent(OPA) or Kyverno  will be helpful.

Tools that help in this area: Kubernetes, Istio, LinkerD, OPA, Kyverno

[Production-Grade Container OrchestrationKubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of…kubernetes.io](https://kubernetes.io/)

[Open Policy AgentStop using a different policy language, policy model, and policy API for every product and service you use. Use OPA for…www.openpolicyagent.org](https://www.openpolicyagent.org/)

[KyvernoKyverno is a policy engine designed for Kubernetes. With Kyverno, policies are managed as Kubernetes resources and no…kyverno.io](https://kyverno.io/)

[IstioA service mesh for observability, security in depth, and management that speeds deployment cycles.istio.io](https://istio.io/)

[The world's lightest, fastest service mesh.Linkerd adds critical security, observability, and reliability to your Kubernetes stack, without any code changes.linkerd.io](https://linkerd.io/)

## Conclusion

Containers orchestration captures the “day-2” operational tasks that sooner or  later we will have to deal with. Fortunately, there are plenty of tools  and standards to help us with this task. The tools mentioned in this  article are examples and mostly ones I’m personally familiar with, but  there are a lot of alternatives. If you are interested in learning more  about those, check out the ever-growing [CNCF landscape](https://landscape.cncf.io/?category=application-definition-image-build&fullscreen=yes&grouping=category).

[CNCF Cloud Native Interactive LandscapeThe Cloud Native Trail Map ( png, pdf) is CNCF's recommended path through the cloud native landscape. The cloud native…landscape.cncf.io](https://landscape.cncf.io/?category=application-definition-image-build&fullscreen=yes&grouping=category)

I hope that this short guide helped you understand what container  orchestration is, why it is important and how you can start thinking  about it in your organization.

[ITNEXT](https://itnext.io/?source=post_sidebar--------------------------post_sidebar-----------)

ITNEXT is a platform for IT developers & software engineers…
