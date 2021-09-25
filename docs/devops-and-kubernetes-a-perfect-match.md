# DevOps and Kubernetes: A Perfect Match?

April 5, 2021

DevOps is a software development strategy that combines development and operations teams into a single unit. Kubernetes is an open source orchestration platform designed to help you manage container deployments at scale. On the surface, it is not entirely clear where these two meet, why and whether this union produces the desired results.

But there is a connection between DevOps and Kubernetes. I want to explore the relationships between enterprise DevOps, agile culture, the role of containers in [CI/CD pipelines](https://containerjournal.com/?s=CI%2FCD) and the integration of Kubernetes into the DevOps pipeline.

## Enterprise DevOps: Culture is Not Enough

Before DevOps gained popularity, development and operations teams operated in silos. This was not unique; single-discipline teams were the norm. Each team had independent processes, goals, and tooling.

Unsurprisingly, these differences often created conflicts between teams and led to bottlenecks and inefficiencies. It also created an atmosphere of ‘us against them’ which was counterproductive to customers and to the bottom line.

[DevOps, when done correctly](https://cloud.netapp.com/devops), helps resolve some of these issues, such as teams not understanding each other’s processes. It does this by requiring cultural changes that force processes and workflows to overlap and run in tandem. However, these cultural changes are not enough to overcome all the issues that exist with siloed teams. Even after managing resistance to cultural change, the issues of tooling and infrastructure remain.

To address these technical issues, DevOps teams use pipelines. These integrated toolchains enable developers to seamlessly submit, test and revise code. Pipelines incorporate automation and configurations designed by operations members along with version control systems used by developers.

This single set of tooling helps ensure that processes align rather than compete. It also helps eliminate the need for either department to wait on the other. When planned carefully, pipelines create visibility into the entire software development life cycle (SDLC). This visibility makes it easier for teams to identify and address issues early on.

All of this is great until teams become restricted by their tools; then, an adjustment is needed. For example, shifting DevOps to the cloud. Kubernetes is perfectly suited to help transition infrastructure to public clouds like [Azure](https://cloud.netapp.com/blog/azure-anf-blg-kubernetes-in-azure-architecture-and-service-options) or [AWS](https://www.aquasec.com/cloud-native-academy/kubernetes-101/kubernetes-on-aws/).

## The Role of Containers in Enterprise-Scale CI/CD

[Once a pipeline is in place](https://www.redhat.com/architect/devops-cicd), it can help an organization vastly improve their agility and their products. However, many pipelines, particularly at first, are cobbled together from a variety of independent tools. To integrate, these tools often require customized plugins or inefficient workarounds.

Even if tools do work well together, the specialization needed for each tool means that toolchains quickly become unwieldy. Each time an individual component needs to be replaced or updated, the whole pipeline has to be rebuilt. Containerization is a solution to these limitations.

With containerization, DevOps teams can break their toolchains down into microservices. Each tool, or individual functionality of a tool, can be separated into a modular piece that can run independently of the environment. This enables teams to easily swap out tools or make modifications without interrupting the rest of the pipeline.

For example, if a specific host configuration is needed for a testing tool, teams are not limited to that same configuration for all tools. This enables DevOps teams to choose the tools that are best suited for their needs, and offers the freedom to reconfigure or scale as needed.

The downside is that having so many containers can be a different sort of challenge to manage. Hence, teams need containers, but they also need a platform like Kubernetes to run those containers.

## Kubernetes: An Enabler for Enterprise DevOps

Many traits and capabilities that come with Kubernetes make it useful for building, deploying, and scaling enterprise-grade DevOps pipelines. These capabilities enable teams to automate the manual labor that orchestration would otherwise require. If teams are going to increase productivity, or more importantly, quality, they need this type of automation.

### Infrastructure and Configuration as Code

Kubernetes enables you to construct your entire infrastructure as code. Every part of your applications and tools can be made accessible to Kubernetes, including access controls, ports and databases. Likewise, you can also manage environment configurations as code. Rather than running a script every time you need to deploy a new environment, you can provide Kubernetes with a source repository containing config files.

Additionally, code can be managed using version control systems, just like your applications in development. This makes it easier for teams to define and modify infrastructure and configurations and allows them to push changes to Kubernetes to be handled automatically.

### Cross-Functional Collaboration

When using Kubernetes to orchestrate your pipeline, you can manage granular controls. This enables you to allow certain roles or applications to take specific actions, while others cannot. For example, restricting customers to deployment or review processes and testers to builds, pending approval.

This sort of control facilitates smooth collaboration while ensuring that your configurations and resources remain consistent. Controlling the scale and deployment of your pipeline resources enables you to ensure that budgets are maintained and can help reduce Kubernetes security risks.

### On-Demand Infrastructure

Kubernetes provides a self-service catalog functionality that enables developers to create infrastructure on-demand. This includes cloud services, such as AWS resources, which are exposed via open service and API standards. These services are based on the configurations allowed by operations members, which helps ensure that compatibility and security remain consistent.

### Zero-Downtime Deployments

Kubernetes’ rolling updates and automated rollback features enable you to deploy new releases with zero downtime. Rather than having to take down production environments and redeploy updated ones, you can use Kubernetes to shift traffic across your available services, updating one cluster at a time.

These features enable you to smoothly achieve blue/green deployments. You can also more easily prioritize new features for customers and conduct A/B testing to ensure that product features are needed and welcome.

Like many agile methodologies, DevOps seeks to improve the entire SDLC. However, the underlying mission is to quickly release software. To achieve this goal, DevOps pipelines become heavily reliant on collaboration, communication, integration and automation.

Containers and microservices help speed up development by enabling modifications on a small scale. This ensures that you can update your software with minimal – or zero – downtime. However, when dealing with enterprise-grade systems composed of thousands of containers, management becomes an issue.

Supposedly, you can use K8s to automate your dev processes and go grab a drink on the beach. While K8s does provide a solution for managing containers, even a container orchestration platform can become overwhelmed by tens of thousands of containers. You also need to ensure that you properly configure your automation.

In addition, some of you are possibly running highly complex pipelines, made up of different languages, integrated with a wide range of systems, all while running a code cocktail that combines propriety, third-party and open source components. Now add data privacy and security compliance issues to the mix, and you have a potentially combustible situation.

This type of party requires an increased level of visibility, security and management controls. You might be able to solve some security issues by using [service meshes](https://devops.com/service-meshes-improving-security-delivery-and-availability/), but even this is not a complete solution.

In short, DevOps and Kubernetes are not a perfect match, but Kubernetes can certainly be a powerful tool when properly configured. Just make sure you are not getting in too deep, and understand that K8s is not an all-encompassing solution.
