# The Cloud Native Landscape: The Provisioning Layer Explained

#### 20 Aug 2020 12:13pm,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack.io/author/jason-morgan/ "Posts by Jason Morgan")

_This post is part of an ongoing series from [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/) and Jason Morgan that focuses on explaining each category of the cloud native landscape. Both are co-organizers of the [Kubernetes Community Days DC](https://kubernetescommunitydays.org/events/2021-washington-dc/) and the [DC Kubernetes meetup group](https://www.meetup.com/All-Things-Kubernetes-k8s-DC/)._

Catherine is Head of Marketing at Buoyant, the creator of Linkerd. A marketing leader turned cloud native evangelist, Catherine is passionate about educating business leaders on the new stack and the critical flexibility it provides.](https://www.linkedin.com/in/catherinepaganini/en/)

In our [introduction to the cloud native landscape](https://thenewstack.io/an-introduction-to-the-cloud-native-landscape/), we provided a high-level overview of the [Cloud Native Computing Foundation](https://www.cncf.io/)‘s cloud native ecosystem. This article is the first in a series that examines each layer at the time. Non-technical readers will learn what the tools in each category are, what problem they solve, and how they address it. We also added a short technical 101 section for those engineers who are just getting started with cloud native.

The first layer in the cloud native landscape is provisioning. This encompasses the tools that are used to create and harden the foundation on which cloud native apps are built, including how the infrastructure is created, managed, and configured — automatically — as well as scanning, signing, and storing container images. The layer also extends to security with tools that enable policies to be set and enforced, authentication and authorization to be built into apps and platforms, and the handling of secrets distribution.

![provisioning layer](https://cdn.thenewstack.io/media/2020/08/5b39e1f1-screen-shot-2020-08-06-at-9.42.29-am.png)

When looking at the [cloud native landscape](https://landscape.cncf.io), you’ll note a few distinctions:

- Projects in large boxes are CNCF-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- Projects in small white boxes are open source projects.
- Products in gray boxes are proprietary products.

Please note that even during the time of this writing, we saw new projects becoming part of the CNCF so always refer to the actual landscape — things are moving fast!

Ok, let’s have a look at each category of the provisioning layer, the role it plays, and how these technologies help adapt applications to a new cloud native reality.

## Automation and Configuration

### What It Is

Automation and configuration tools speed up the creation and configuration of compute resources (virtual machines, networks, firewall rules, load balancers, etc.). These tools may handle different parts of the provisioning process or try to control everything end to end. Most provide the ability to integrate with other projects and products in the space.

### Problem it addresses

Traditionally, IT processes relied on lengthy and labor-intensive manual release cycles, typically between three to six months. Those cycles came with lots of human processes and controls that slowed down changes to production environments. These slow-release cycles and static environments aren’t compatible with cloud native development. To deliver on rapid development cycles, infrastructure must be provisioned dynamically and without human intervention.

### How it helps

Tools of this category allow engineers to build computing environments without human intervention. By codifying the environment setup it becomes reproducible with the click of a button. While manual setup is error-prone, once codified, environment creation matches the exact desired state — a huge advantage.

While tools may take different approaches, they all aim at reducing the required work to provision resources through automation.

### Technical 101

As we move from old-style human-driven provisioning to the new on-demand scaling model demanded by the cloud we find that the patterns and tools we used before no longer meet our needs. Your organization can’t afford to maintain a large 24×7 staff whose job it is to create, configure, and manage servers. Automated tools like Terraform reduce the level of effort required to scale tens of servers and associated network and up to hundreds of firewall rules. Tools like Puppet, Chef, and Ansible provision and/or configure these new servers and applications programmatically as they are spun up and allow them to be consumed by developers.

Some tools interact directly with the infrastructure APIs provided by platforms like AWS or vSphere, while others focus on configuring the individual machines to make them part of a Kubernetes cluster. Many, like Chef and Terraform, can interoperate to provision and configure the environment. Others, like OpenStack, exist to provide an Infrastructure-as-a-Service (IaaS) environment that other tools could consume. Fundamentally, you’ll need one or more tools in this space as part of laying down the computing environment, CPU, memory, storage, and networking, for your Kubernetes clusters. You’ll also need a subset of these to create and manage the Kubernetes clusters themselves.

At the time of this writing, there are three CNCF projects in this space: KubeEdge, a Sandbox CNCF project, as well as Kubespray and Kops (the latter two are Kubernetes subprojects and belong thus to the CNCF although they aren’t yet listed on the landscape). Most of the tools in this category offer an open source as well as a paid version.

BuzzwordsPopular Projects/Products

- Infrastructure-as-Code (IaC)
- Automation
- Declarative Configuration

- Chef
- Puppet
- Ansible
- Terraform

![automation and config](https://cdn.thenewstack.io/media/2020/08/df17f0e2-screen-shot-2020-08-06-at-9.52.34-am.png)

## Container Registry

### What It Is

Before defining container registries, let’s first discuss three tightly related concepts:

1. A container is a set of compute constraints used to execute a process. Processes launched within containers are tricked to believe they are running on their own dedicated computer vs. a machine shared with other processes (similar to virtual machines). In short, containers allow you to run your code in a controlled fashion no matter where it is.
2. An image is the set of archive files needed to run a container and its process. You could see it as a form of template on which you can create an unlimited number of containers.
3. A repository, or just repo, is a space to store images.

Back to container registries. Container registries are specialized web applications to categorize and store repositories.

In summary, images contain the information needed to execute a program (within a container) and are stored in repositories which in turn are categorized and grouped in registries. Tools that build, run, and manage containers need access to those images. Access is provided by referencing to the registry (the path to access the image).

![Registry, repo, containers](https://cdn.thenewstack.io/media/2020/08/b348c280-screen-shot-2020-08-06-at-9.59.11-am.png)

### Problem It Addresses

Cloud native applications are packaged and run as containers. Container registries store and provide these container images.

### How It Helps

By centrally storing all container images in one place, they are easily accessible for any developer working on that app.

### Technical 101

Container registry tools exist to either store and distribute images or to enhance an existing registry in some way. Fundamentally, a registry is a kind of web API that allows container engines to store and retrieve images. Many provide interfaces to allow container scanning or signing tools to enhance the security of the images they store. Some specialize in distributing or duplicating images in a particularly efficient manner. Any environment using containers will need to use one or more registries.

Tools in this space can provide integrations to scan, sign, and inspect the images they store. At the time of this writing Dragonfly and Harbor are CNCF projects and Harbor recently gained the distinction of [being the first](https://goharbor.io/blog/harbor-2.0/) OCI compliant registry. Each major cloud provider provides its own hosted registry and many other registries can be deployed standalone or directly into your Kubernetes cluster via tools like Helm.

BuzzwordsPopular Projects/Products

- Container
- OCI Image
- Registry

- Docker Hub
- Harbor
- Hosted registries from AWS, Azure, and GCP
- Artifactory

![Container registries](https://cdn.thenewstack.io/media/2020/08/ee64a02c-screen-shot-2020-08-06-at-10.02.34-am.png)

## Security and Compliance

### What It Is

Cloud native applications are designed to be rapidly iterated on. Think of the continuous flow of updates your iPhone apps get — everyday, they are evolving, presumably getting better. In order to release code on a regular cadence, we must ensure that our code and our operating environment are secure and only accessed by authorized engineers. Tools and projects in this section represent some of the things needed to create and run modern applications in a secure fashion.

### Problem It Addresses

These tools and projects help you harden, monitor, and enforce security for your platforms and applications. From the container to your Kubernetes environment, they enable you to set policies (for compliance), get insights into existing vulnerabilities, catch misconfigurations, and harden the containers and clusters.

### How It Helps

In order to securely run containers they must be scanned for known vulnerabilities and signed to ensure they haven’t been tampered with. Kubernetes itself defaults to extremely permissive access control settings that are unsuitable for production use. Furthermore, Kubernetes clusters are an attractive target to anyone looking to attack your systems. The tools and projects in this space help harden the cluster and provide tooling to detect when the system is behaving abnormally.

### Technical 101

In order to operate securely in a dynamic and rapidly evolving environment we must treat security as part of the platform and application development lifecycle. The tools in this space are an extremely varied group and seek to solve different portions of the problem. Most of the tooling falls into one of the following categories:

- Audit and compliance
- Path to production hardening tools
  - Code scanning
  - Vulnerability scanning
  - Image signing
- Policy creation and enforcement
- Network layer security

Some of these tools and projects will rarely be used directly, like Trivy, Claire, and Notary which are leveraged by registries or other scanning tools. Others are key hardening components of a modern application platform, like Falco or Open Policy Agent (OPA).

There are a number of mature vendors providing solutions in this space, as well as startups founded explicitly on bringing Kubernetes native frameworks to market. At the time of this writing Falco, Notary/TUF, and OPA are the only CNCF projects in this space.

BuzzwordsPopular Projects/Products

- Image scanning
- Image signing
- Policy enforcement
- Audit

- OPA
- Falco
- Sonobuoy

![security and compliance](https://cdn.thenewstack.io/media/2020/08/d60f1b0b-screen-shot-2020-08-06-at-10.06.50-am.png)

## Key (and Identity) Management

### What It Is

Before we go into key management, let’s first define cryptographic keys. A key is a string of characters used to encrypt or sign data. Like a physical key, it locks (encrypts) data so that only someone with the right key can unlock (decrypt) it.

As applications and operations adapt to a new cloud native world, security tools are evolving to meet new security needs. The tools and projects in this category cover everything from how to securely store passwords and other secrets (sensitive data such as API keys, encryption keys, etc.) to how to safely eliminate passwords and secrets from your microservices environment.

### Problem It Addresses

Cloud native environments are highly dynamic calling for secret distribution that is on-demand, entirely programmatic (no humans in the loop), and automated. Applications must also know if a given request comes from a valid source (authentication) and if that request has the right to do whatever it’s trying to do (authorization). This is commonly referred to as AuthN and AuthZ.

### How It Helps

Each tool or project takes a different approach but they all provide a way to either securely distribute secrets and keys, or they provide a service or specification related to authentication, authorization, or both.

### Technical 101

Tools in this category can be grouped into two sets: While some tools focus on key generation, storage, management and rotation, the other group focuses on single sign-on and identity management. Vault, for instance, is a rather generic key management tool allowing you to manage different types of keys. Keycloak, on the other hand, is an identity broker which can be used to manage access keys for different services.

At the time of this writing SPIFFE/SPIRE are the only CNCF projects in this space, and most tools offer an open source as well as paid version.

BuzzwordsPopular Projects

- AuthN and AuthZ
- Identity
- Access
- Secrets

- Vault
- Spiffe
- OAuth2

![Key management](https://cdn.thenewstack.io/media/2020/08/62e5a458-screen-shot-2020-08-06-at-10.13.56-am.png)

As we’ve seen the provisioning layer focuses on building the foundation of your cloud native platforms and applications, with tools handling everything from infrastructure provisioning to container registries to security. This piece is intended to be the first in a series of articles detailing the cloud native landscape. In the next article we’ll focus on the runtime layer and explore cloud native storage, container runtime, and networking.

_A very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate. Also, a big thanks to [Oleg Chunihkin](https://www.linkedin.com/in/olegch/) for all his input early on in this project._

The Cloud Native Computing Foundation is a sponsor of The New Stack.

Feature image by [torstensimon](https://pixabay.com/users/torstensimon-5039407/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=2211588) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=2211588).

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Docker.
