# The Cloud Native Landscape: Platforms Explained

#### 17 Mar 2021 12:08pm,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack.io/author/jason-morgan/ "Posts by Jason Morgan")

![](https://cdn.thenewstack.io/media/2021/03/2244406d-background-1126047_640.jpg)

_This post is part of an ongoing series from the Cloud Native Computing Foundation’s_ [_Business Value Subcommittee_](https://lists.cncf.io/g/cncf-business-value) _co-chairs_ [_Catherine Paganini_](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) _and_ [_Jason Morgan_](https://thenewstack.io/author/jason-morgan/) _that focuses on explaining each category of the [cloud native landscape](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/) to a non-technical audience as well as engineers just getting started with cloud native. See also installments on the layers for [application definition development](https://thenewstack.io/the-cloud-native-landscape-the-application-definition-and-development-layer/), the [runtime](https://thenewstack.io/the-cloud-native-landscape-the-runtime-layer-explained/), the [orchestration and management](https://thenewstack.io/the-cloud-native-landscape-the-orchestration-and-management-layer/), and the [provisioning.](https://thenewstack.io/the-cloud-native-landscape-the-provisioning-layer-explained/)_


There isn’t anything inherently new in these platforms. Everything they do can be done by one of the tools in these layers or the observability and analysis column. You could certainly build your own platform, and in fact, many organizations do. However, configuring and fine-tuning the different modules reliably and securely while ensuring that all technologies are always updated and vulnerabilities patched is no easy task—you’ll need a dedicated team to build and maintain it. If you don’t have the required bandwidth and/or know-how, your team is likely better off with a platform. For some organizations, especially those with small engineering teams, platforms are the only way to adopt a cloud native approach.

You’ll probably notice, all platforms [revolve around Kubernetes](https://thenewstack.io/primer-how-kubernetes-came-to-be-what-it-is-and-why-you-should-care/). That’s because Kubernetes, is at the core of the cloud native stack.

### Sidenote

When looking at the [Cloud Native Landscape](https://landscape.cncf.io/), you’ll note a few distinctions:

- Projects in large boxes are Cloud Native Computing Foundation-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- Projects in small white boxes are open source projects.
- Products in gray boxes are proprietary.

Please note that even during the time of this writing, we saw new projects becoming part of the Cloud Native Computing Foundation (CNCF) so always refer to the actual landscape — things are moving fast!

## Kubernetes Distributions

### What It Is

A distribution, or distro, is when a vendor takes core Kubernetes — that’s the unmodified, open source code (although some modify it) — and packages it for redistribution. Usually, that entails finding and validating the Kubernetes software and providing a mechanism handling cluster installation and upgrades. Many Kubernetes distributions include other proprietary or open source applications.

### What’s the Problem They Solve

[Open source Kubernetes](https://github.com/kubernetes/kubernetes) doesn’t specify a particular installation tool and leaves many setup configuration choices to the user. Additionally, there is limited support for issues as they arise beyond community resources like [Community Forums](https://discuss.kubernetes.io/), [StackOverflow](http://stackoverflow.com/questions/tagged/kubernetes), [Slack](https://slack.k8s.io/) or Discord.

While using Kubernetes has gotten easier over time, it can be challenging to find and use the open source installers. Users need to understand what versions to use, where to get them, and if a particular component is compatible with another. They also need to decide what software will be deployed to their clusters and what settings to use to ensure their platforms are secure, stable, and performant. All this requires deep Kubernetes expertise that may not be readily available in-house.

### How It Helps

Kubernetes distributions provide a trusted and reliable way to install Kubernetes and provide opinionated defaults that create a better and more secure operating environment. A Kubernetes distribution gives vendors and projects the control and predictability they need to provide support for a customer as they go through the lifecycle of deploying, maintaining, and upgrading their Kubernetes clusters.

That predictability enables distribution providers to support users when they have production issues. Distributions also often provide a tested and supported upgrade path that allows users to keep their Kubernetes clusters up to date. Additionally, distributions often provide software to deploy on top of Kubernetes that makes it easier to use.

Distributions significantly ease and speed up Kubernetes adoption. Since the expertise needed to configure and fine-tune the clusters is coded into the platform, organizations can get up and running with cloud native tools without having to hire additional engineers with specialized expertise.

### Technical 101

If you’ve installed Kubernetes, you’ve likely used something like kubeadm to get your cluster up and running. Even then, you probably had to decide on a CNI (Container Network Interface), install, and configure it. Then, you might have added some storage classes, a tool to handle log messages, maybe an ingress controller, and the list goes on. A Kubernetes distribution will automate some, or all, of that setup. It will also ship with configuration settings based on its own interpretation of best practice or an intelligent default. Additionally, most distributions will come with some extensions or add-ons bundled and tested to ensure you can get going with your new cluster as quickly as possible.

Let’s take [Kublr](https://kublr.com/) as an example. With Kubernetes at its core, this platform bundles technologies from mainly three layers: provisioning, runtime, orchestration and management, and observability and analysis. All modules are preconfigured with a few options to choose from and ready to go.Different platforms have different focal points. In the case of Kublr, the focus is more on the operations side, while other platforms may focus more on developer tooling.

There are a lot of options in this category. As of this writing, [k3s](https://k3s.io) is the only CNCF project. There are a lot of great open source and commercial options available, including, Microk8s from Canonical, k3s, Tanzu Kubernetes Grid from [VMware](https://tanzu.vmware.com?utm_content=inline-mention), Docker Enterprise from [Mirantis](https://www.mirantis.com/?utm_content=inline-mention), Rancher from Suse, and of course [Red Hat](https://www.openshift.com/try?utm_content=inline-mention)’s Openshift. We didn’t have time to mention even close to half the Kubernetes distributions, and we encourage you to think carefully about your needs when you begin evaluating distributions.

![Kubernetes distributions](https://cdn.thenewstack.io/media/2021/03/342aedc7-screen-shot-2021-03-16-at-9.18.35-am.png)

## Hosted Kubernetes

### What It Is

Hosted Kubernetes is a service offered by infrastructure providers like [Amazon Web Services](https://aws.amazon.com/?utm_content=inline-mention) (AWS), DigitalOcean, Azure, or Google, allowing customers to spin up a Kubernetes cluster on-demand. The cloud provider takes responsibility for managing part of the Kubernetes cluster, usually called the control plane. They are similar to distributions but managed by the cloud provider on _their_ infrastructure.

### What’s the Problem They Solve

Hosted Kubernetes allows teams to get started with Kubernetes without knowing or doing anything beyond setting up an account with a cloud vendor. It solves four of the five Ws of getting started with Kubernetes. Who (manages it): your cloud provider; what: their hosted Kubernetes offering; when: now; and where: on the cloud providers infrastructure. The why is up to you.

### How It Helps

Since the provider takes care of all management details, hosted Kubernetes is the easiest way to get started with cloud native. All users have to do, is develop their apps and deploy them on the hosted Kubernetes services — it’s incredibly convenient. The hosted offering allows users to spin up a Kubernetes cluster and get started right away,\* while taking some responsibility for the cluster availability. It’s worth noting that with the extra convenience of these services comes some reduced flexibility. The offering is bound to the cloud provider, and Kubernetes users don’t have access to the Kubernetes control plane, so some configuration options are limited.

_\\* Slight exception for EKS from AWS as it also requires users to take some additional steps to prepare their clusters._

### Technical 101

Hosted Kubernetes are on-demand Kubernetes clusters provided by a vendor, usually an infrastructure hosting provider. The vendor takes responsibility for provisioning the cluster and managing the Kubernetes control plane. Again, the notable exception is EKS, where individual node provisioning is left up to the client.

Hosted Kubernetes allows an organization to quickly provision new clusters and reduce their operational risk by outsourcing infrastructure component management to another organization. The main trade-offs are that you’ll likely be charged for the control plane management (GKE ran into a [bit of controversy](https://www.theregister.com/2020/03/05/google_reintroduces_management_fee_for_kubernetes_clusters/) around price changes last year) and that you’ll be limited in what you can do. Managed clusters provide stricter limits on configuring your Kubernetes cluster than DIY Kubernetes clusters.

There are numerous vendors and projects in this space and, at the time of this writing, no CNCF projects.

![Hosted Kubernetes](https://cdn.thenewstack.io/media/2021/03/a061e08f-screen-shot-2021-03-16-at-9.21.08-am.png)

## Kubernetes Installer

### What It Is

Kubernetes installers help install Kubernetes on a machine. They automate the Kubernetes installation and configuration process and may even help with upgrades. Kubernetes installers are often coupled with or used by Kubernetes distributions or hosted Kubernetes offerings.

### What’s the Problem They Solve

Similar to Kubernetes distributions, Kubernetes installers simplify getting started with Kubernetes. [Open source Kubernetes](https://github.com/kubernetes/kubernetes) relies on installers like kubeadm, which, as of this writing, is part of the Certified Kubernetes Administrator certification test to get Kubernetes clusters up and running.

### How It Helps

Kubernetes installers ease the Kubernetes installation process. Like distributions, they provide a vetted source for the source code and version. They also often ship with opinionated Kubernetes environment configurations. Kubernetes installers like-kind (Kubernetes in Docker) allow you to get a Kubernetes cluster with a single command.

### Technical 101

Whether you’re installing Kubernetes locally on Docker, spinning up and provisioning new virtual machines, or preparing new physical servers, you’re going to need a tool to handle all the preparation of various Kubernetes components (unless you’re looking to do it [the hard way](https://github.com/kelseyhightower/kubernetes-the-hard-way)).

Kubernetes installers simplify that process. Some handle spinning up nodes and others merely configure nodes you’ve already provisioned. They all offer various levels of automation and are each suited for different use cases. When getting started with an installer, start by understanding your needs, then pick an installer that addresses them. At the time of this writing, kubeadm is considered so fundamental to the Kubernetes ecosystem that it’s included as part of the CKA, certified Kubernetes administrator exam. Minikube, kind, kops, and kubespray are all CNCF-owned Kubernetes installer projects.

![Kubernetes installer](https://cdn.thenewstack.io/media/2021/03/2891980c-screen-shot-2021-03-16-at-9.24.16-am.png)

## PaaS / Container Service

### What It Is

A [platform as a service](https://en.wikipedia.org/wiki/Platform_as_a_service), or PaaS, is an environment that allows users to run applications without necessarily understanding or knowing about the underlying compute resources. PaaS and container services in this category are mechanisms to either host a PaaS for developers or host services they can use.

### What’s the Problem They Solve

In this series, we’ve talked a lot about the tools and technologies around “cloud native.” A PaaS attempts to connect many of the technologies found in this landscape in a way that provides direct value to developers. It answers the following questions: how will I run applications in various environments and, once running, how will my team and users interact with them?

### How It Helps

PaaS provides opinions and choices around how to piece together the various open and closed source tools needed to run applications. Many offerings include tools that handle PaaS installation and upgrades and the mechanisms to convert application code into a running application. Additionally, PaaS handle the runtime needs of application instances, including on-demand scaling of individual components and visibility into performance and log messages of individual apps.

### Technical 101

Organizations are adopting cloud native technologies to achieve specific business or organizational objectives. A PaaS provides a quicker path to value than building a custom application platform. Tools like Heroku or Cloud Foundry Application Runtime help organizations get up and running with new applications quickly. They excel at providing the tools needed to run [12 factor](https://12factor.net/) or cloud native applications.

Any PaaS comes with its own set of trade-offs and restrictions. Most only work with a subset of languages or application types and the opinions and decisions baked into them may or may not be a good fit for your needs. Stateless applications tend to do very well in a PaaS but stateful applications like databases usually don’t. There are currently no CNCF projects in this space but most of the offerings are open source and Cloud Foundry is managed by the Cloud Foundry Foundation.

![Kubernetes PaaS / Container as a Service](https://cdn.thenewstack.io/media/2021/03/a12a4c10-screen-shot-2021-03-16-at-9.26.50-am.png)

## Conclusion

As we’ve seen there are multiple tools that help ease Kubernetes adoption. From Kubernetes distributions and hosted Kubernetes to more barebones installers or PaaS, they all take some of the installation and configuration burden and pre-package it for you. Each solution comes with its own “flavor.” Vendor opinions about what’s important and appropriate are built into the solution.

Before adopting any of these, you’ll need to do some research to identify the best solution for your particular use case. Will you likely encounter advanced Kubernetes scenarios where you’ll need control over the control plane? Then, hosted solutions are likely not a good fit. Do you have a small team that manages “standard” workloads and need to offload as many operational tasks as possible? Then, hosted solutions may be a great fit. Is portability important? What about production-readiness? There are multiple aspects to consider. There is no “one best tool,” but there certainly is an optimal tool for your use case. Hopefully, this article will help you narrow your search down to the right “bucket.”

This concludes the platform “column” of the CNCF landscape. Next, we’ll tackle the last article of this series, the observability and analysis “column.”

_As always, a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Mirantis, Docker, Bit.

Amazon Web Services, CNCF, Mirantis, Red Hat and VMware are sponsors of The New Stack.


Image par [kalhh](https://pixabay.com/fr/users/kalhh-86169/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1126047) de [Pixabay](https://pixabay.com/fr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1126047).

