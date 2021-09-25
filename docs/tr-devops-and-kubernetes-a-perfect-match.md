# DevOps and Kubernetes: A Perfect Match?

# DevOps 和 Kubernetes：完美匹配？

April 5, 2021

2021 年 4 月 5 日

DevOps is a software development strategy that combines development and operations teams into a single unit. Kubernetes is an open source orchestration platform designed to help you manage container deployments at scale. On the surface, it is not entirely clear where these two meet, why and whether this union produces the desired results.

DevOps 是一种软件开发策略，它将开发和运营团队合并为一个单元。 Kubernetes 是一个开源编排平台，旨在帮助您大规模管理容器部署。从表面上看，这两者在哪里相遇，为什么以及这种结合是否会产生预期的结果并不完全清楚。

But there is a connection between DevOps and Kubernetes. I want to explore the relationships between enterprise DevOps, agile culture, the role of containers in [CI/CD pipelines](https://containerjournal.com/?s=CI%2FCD) and the integration of Kubernetes into the DevOps pipeline.

但是 DevOps 和 Kubernetes 之间存在联系。我想探索企业 DevOps、敏捷文化、容器在 [CI/CD 管道](https://containerjournal.com/?s=CI%2FCD) 中的作用以及 Kubernetes 与 DevOps 管道的集成之间的关系。

## Enterprise DevOps: Culture is Not Enough

## 企业 DevOps：文化还不够

Before DevOps gained popularity, development and operations teams operated in silos. This was not unique; single-discipline teams were the norm. Each team had independent processes, goals, and tooling.

在 DevOps 流行之前，开发和运营团队在孤岛中运作。这不是独一无二的。单一学科的团队是常态。每个团队都有独立的流程、目标和工具。

Unsurprisingly, these differences often created conflicts between teams and led to bottlenecks and inefficiencies. It also created an atmosphere of ‘us against them’ which was counterproductive to customers and to the bottom line.

不出所料，这些差异经常在团队之间造成冲突，并导致瓶颈和低效率。它还营造了一种“我们反对他们”的氛围，这对客户和底线都适得其反。

[DevOps, when done correctly](https://cloud.netapp.com/devops), helps resolve some of these issues, such as teams not understanding each other’s processes. It does this by requiring cultural changes that force processes and workflows to overlap and run in tandem. However, these cultural changes are not enough to overcome all the issues that exist with siloed teams. Even after managing resistance to cultural change, the issues of tooling and infrastructure remain.

[DevOps，如果正确完成](https://cloud.netapp.com/devops) 有助于解决其中一些问题，例如团队不了解彼此的流程。它通过要求强制流程和工作流重叠和协同运行的文化变革来实现这一点。然而，这些文化变革并不足以克服孤立团队存在的所有问题。即使在应对文化变革的阻力之后，工具和基础设施的问题仍然存在。

To address these technical issues, DevOps teams use pipelines. These integrated toolchains enable developers to seamlessly submit, test and revise code. Pipelines incorporate automation and configurations designed by operations members along with version control systems used by developers.

为了解决这些技术问题，DevOps 团队使用管道。这些集成的工具链使开发人员能够无缝地提交、测试和修改代码。管道结合了运营成员设计的自动化和配置以及开发人员使用的版本控制系统。

This single set of tooling helps ensure that processes align rather than compete. It also helps eliminate the need for either department to wait on the other. When planned carefully, pipelines create visibility into the entire software development life cycle (SDLC). This visibility makes it easier for teams to identify and address issues early on.

这一套工具有助于确保流程一致而不是竞争。它还有助于消除任何一个部门等待另一个部门的需要。如果仔细规划，管道可以创建对整个软件开发生命周期 (SDLC) 的可见性。这种可见性使团队可以更轻松地尽早识别和解决问题。

All of this is great until teams become restricted by their tools; then, an adjustment is needed. For example, shifting DevOps to the cloud. Kubernetes is perfectly suited to help transition infrastructure to public clouds like [Azure](https://cloud.netapp.com/blog/azure-anf-blg-kubernetes-in-azure-architecture-and-service-options) or [ AWS](https://www.aquasec.com/cloud-native-academy/kubernetes-101/kubernetes-on-aws/).

所有这一切都很棒，直到团队受到工具的限制；然后，需要进行调整。例如，将 DevOps 转移到云端。 Kubernetes 非常适合帮助将基础设施过渡到公共云，如 [Azure](https://cloud.netapp.com/blog/azure-anf-blg-kubernetes-in-azure-architecture-and-service-options) 或 [ AWS](https://www.aquasec.com/cloud-native-academy/kubernetes-101/kubernetes-on-aws/)。

## The Role of Containers in Enterprise-Scale CI/CD

## 容器在企业级 CI/CD 中的作用

[Once a pipeline is in place](https://www.redhat.com/architect/devops-cicd), it can help an organization vastly improve their agility and their products. However, many pipelines, particularly at first, are cobbled together from a variety of independent tools. To integrate, these tools often require customized plugins or inefficient workarounds.

[一旦管道就位](https://www.redhat.com/architect/devops-cicd)，它可以帮助组织极大地提高敏捷性和产品。但是，许多管道，尤其是最初的管道，是由各种独立工具拼凑而成的。为了集成，这些工具通常需要定制的插件或低效的解决方法。

Even if tools do work well together, the specialization needed for each tool means that toolchains quickly become unwieldy. Each time an individual component needs to be replaced or updated, the whole pipeline has to be rebuilt. Containerization is a solution to these limitations.

即使工具可以很好地协同工作，每个工具所需的专业化也意味着工具链很快就会变得笨拙。每次需要更换或更新单个组件时，都必须重建整个管道。容器化是对这些限制的解决方案。

With containerization, DevOps teams can break their toolchains down into microservices. Each tool, or individual functionality of a tool, can be separated into a modular piece that can run independently of the environment. This enables teams to easily swap out tools or make modifications without interrupting the rest of the pipeline.

通过容器化，DevOps 团队可以将他们的工具链分解为微服务。每个工具或工具的单个功能都可以分离成一个模块化的部分，可以独立于环境运行。这使团队能够轻松更换工具或进行修改，而不会中断管道的其余部分。

For example, if a specific host configuration is needed for a testing tool, teams are not limited to that same configuration for all tools. This enables DevOps teams to choose the tools that are best suited for their needs, and offers the freedom to reconfigure or scale as needed.

例如，如果测试工具需要特定的主机配置，则团队不限于所有工具都使用相同的配置。这使 DevOps 团队能够选择最适合他们需求的工具，并提供根据需要重新配置或扩展的自由。

The downside is that having so many containers can be a different sort of challenge to manage. Hence, teams need containers, but they also need a platform like Kubernetes to run those containers.

缺点是拥有如此多的容器可能是一种不同的管理挑战。因此，团队需要容器，但他们也需要像 Kubernetes 这样的平台来运行这些容器。

## Kubernetes: An Enabler for Enterprise DevOps 

## Kubernetes：企业 DevOps 的推动者

Many traits and capabilities that come with Kubernetes make it useful for building, deploying, and scaling enterprise-grade DevOps pipelines. These capabilities enable teams to automate the manual labor that orchestration would otherwise require. If teams are going to increase productivity, or more importantly, quality, they need this type of automation.

Kubernetes 附带的许多特性和功能使其可用于构建、部署和扩展企业级 DevOps 管道。这些功能使团队能够自动化编排否则需要的手工劳动。如果团队要提高生产力，或更重要的是提高质量，他们需要这种类型的自动化。

### Infrastructure and Configuration as Code

### 基础设施和配置即代码

Kubernetes enables you to construct your entire infrastructure as code. Every part of your applications and tools can be made accessible to Kubernetes, including access controls, ports and databases. Likewise, you can also manage environment configurations as code. Rather than running a script every time you need to deploy a new environment, you can provide Kubernetes with a source repository containing config files.

Kubernetes 使您能够将整个基础设施构建为代码。 Kubernetes 可以访问您的应用程序和工具的每个部分，包括访问控制、端口和数据库。同样，您也可以将环境配置作为代码进行管理。您可以为 Kubernetes 提供一个包含配置文件的源存储库，而不是每次需要部署新环境时都运行脚本。

Additionally, code can be managed using version control systems, just like your applications in development. This makes it easier for teams to define and modify infrastructure and configurations and allows them to push changes to Kubernetes to be handled automatically.

此外，可以使用版本控制系统管理代码，就像开发中的应用程序一样。这使团队可以更轻松地定义和修改基础架构和配置，并允许他们将更改推送到 Kubernetes 以自动处理。

### Cross-Functional Collaboration

### 跨职能协作

When using Kubernetes to orchestrate your pipeline, you can manage granular controls. This enables you to allow certain roles or applications to take specific actions, while others cannot. For example, restricting customers to deployment or review processes and testers to builds, pending approval.

使用 Kubernetes 编排管道时，您可以管理精细控制。这使您能够允许某些角色或应用程序执行特定操作，而其他角色或应用程序则不能。例如，限制客户部署或审查流程和测试人员构建，等待批准。

This sort of control facilitates smooth collaboration while ensuring that your configurations and resources remain consistent. Controlling the scale and deployment of your pipeline resources enables you to ensure that budgets are maintained and can help reduce Kubernetes security risks.

这种控制有助于顺利协作，同时确保您的配置和资源保持一致。控制管道资源的规模和部署使您能够确保维持预算并有助于降低 Kubernetes 安全风险。

### On-Demand Infrastructure

### 按需基础设施

Kubernetes provides a self-service catalog functionality that enables developers to create infrastructure on-demand. This includes cloud services, such as AWS resources, which are exposed via open service and API standards. These services are based on the configurations allowed by operations members, which helps ensure that compatibility and security remain consistent.

Kubernetes 提供自助服务目录功能，使开发人员能够按需创建基础架构。这包括通过开放服务和 API 标准公开的云服务，例如 AWS 资源。这些服务基于运营成员允许的配置，这有助于确保兼容性和安全性保持一致。

### Zero-Downtime Deployments

### 零停机部署

Kubernetes’ rolling updates and automated rollback features enable you to deploy new releases with zero downtime. Rather than having to take down production environments and redeploy updated ones, you can use Kubernetes to shift traffic across your available services, updating one cluster at a time.

Kubernetes 的滚动更新和自动回滚功能使您能够以零停机时间部署新版本。不必关闭生产环境并重新部署更新的环境，您可以使用 Kubernetes 在可用服务之间转移流量，一次更新一个集群。

These features enable you to smoothly achieve blue/green deployments. You can also more easily prioritize new features for customers and conduct A/B testing to ensure that product features are needed and welcome.

这些特性使您能够顺利实现蓝/绿部署。您还可以更轻松地为客户优先考虑新功能并进行 A/B 测试以确保产品功能是需要和受欢迎的。

Like many agile methodologies, DevOps seeks to improve the entire SDLC. However, the underlying mission is to quickly release software. To achieve this goal, DevOps pipelines become heavily reliant on collaboration, communication, integration and automation.

与许多敏捷方法一样，DevOps 寻求改进整个 SDLC。但是，其基本任务是快速发布软件。为了实现这一目标，DevOps 管道变得严重依赖协作、通信、集成和自动化。

Containers and microservices help speed up development by enabling modifications on a small scale. This ensures that you can update your software with minimal – or zero – downtime. However, when dealing with enterprise-grade systems composed of thousands of containers, management becomes an issue.

容器和微服务通过支持小规模修改来帮助加快开发速度。这确保您可以在最短或零停机时间的情况下更新您的软件。但是，在处理由数千个容器组成的企业级系统时，管理成为一个问题。

Supposedly, you can use K8s to automate your dev processes and go grab a drink on the beach. While K8s does provide a solution for managing containers, even a container orchestration platform can become overwhelmed by tens of thousands of containers. You also need to ensure that you properly configure your automation.

据说，您可以使用 K8s 来自动化您的开发流程，然后去海滩上喝一杯。虽然 K8s 确实提供了管理容器的解决方案，但即使是容器编排平台也可能被数以万计的容器所淹没。您还需要确保正确配置自动化。

In addition, some of you are possibly running highly complex pipelines, made up of different languages, integrated with a wide range of systems, all while running a code cocktail that combines propriety, third-party and open source components. Now add data privacy and security compliance issues to the mix, and you have a potentially combustible situation.

此外，你们中的一些人可能正在运行高度复杂的管道，由不同的语言组成，与广泛的系统集成，同时运行结合了专有、第三方和开源组件的代码鸡尾酒。现在将数据隐私和安全合规性问题添加到组合中，您可能会遇到可燃的情况。

This type of party requires an increased level of visibility, security and management controls. You might be able to solve some security issues by using [service meshes](https://devops.com/service-meshes-improving-security-delivery-and-availability/), but even this is not a complete solution. 

这种类型的聚会需要更高级别的可见性、安全性和管理控制。您也许可以通过使用 [服务网格](https://devops.com/service-meshes-improving-security-delivery-and-availability/) 来解决一些安全问题，但这也不是一个完整的解决方案。

In short, DevOps and Kubernetes are not a perfect match, but Kubernetes can certainly be a powerful tool when properly configured. Just make sure you are not getting in too deep, and understand that K8s is not an all-encompassing solution. 

简而言之，DevOps 和 Kubernetes 并不是完美的搭配，但如果配置得当，Kubernetes 肯定可以成为一个强大的工具。只要确保您没有深入了解，并了解 K8s 不是一个包罗万象的解决方案。

