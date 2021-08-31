# 7 Principles of DevSecOps With Kubernetes

# Kubernetes DevSecOps 的 7 条原则

March 15, 2021

In my article, “ [9 Pillars of Engineering DevOps With Kubernetes](https://containerjournal.com/uncategorized/9-pillars-of-engineering-devops-with-kubernetes/),” I explain that continuous security is a core pillar of every well-engineered DevOps.

在我的文章“[Kubernetes 的工程 DevOps 的 9 个支柱](https://containerjournal.com/uncategorized/9-pillars-of-engineering-devops-with-kubernetes/)”中，我解释说持续安全是一个核心每个精心设计的 DevOps 的支柱。

As indicated in the white paper, “ [From the Node Up: The Complete Guide to Kubernetes Security With Prisma Cloud](http://paloaltonetworks.com/prisma/cloud),” Kubernetes is a multi-layered, complex platform that consists of more than half a dozen different components that present both challenges and opportunities for DevSecOps.

正如白皮书“[从节点向上：使用 Prisma Cloud 确保 Kubernetes 安全性的完整指南](http://paloaltonetworks.com/prisma/cloud)”中所述，Kubernetes 是一个多层、复杂的平台，包括六个不同的组件，这些组件为 DevSecOps 带来了挑战和机遇。

Complex applications operating over complex distributed infrastructure can be difficult to secure. Cloud-native tools, such as Kubernetes, provide more insight into what is happening within an application, making it easier to identify and fix security problems. The enhanced orchestration controls, provided by Kubernetes on the deployment and deployed containerized applications, benefit from immutable consistency and improved response times. In addition, the secrets objects offer a secure way to store sensitive data.

在复杂的分布式基础设施上运行的复杂应用程序可能难以保护。 Kubernetes 等云原生工具可以更深入地了解应用程序中发生的情况，从而更容易识别和修复安全问题。 Kubernetes 在部署和已部署的容器化应用程序上提供的增强的编排控制受益于不可变的一致性和改进的响应时间。此外，secrets 对象提供了一种安全的方式来存储敏感数据。

In this article, I indicate how Kubernetes can be used and configured to satisfy seven principles for a successful DevSecOps approach using Kubernetes. The seven DevSecOps principles are those identified in the [Department of Defense Enterprise DevSecOps Reference Design](https://dodcio.defense.gov/Portals/0/Documents/DoD%20Enterprise%20DevSecOps%20Reference%20Design%20v1.0_Public%20Release .pdf).

在本文中，我将说明如何使用和配置 Kubernetes 以满足使用 Kubernetes 的成功 DevSecOps 方法的七个原则。七项 DevSecOps 原则是在 [国防部企业 DevSecOps 参考设计](https://dodcio.defense.gov/Portals/0/Documents/DoD%20Enterprise%20DevSecOps%20Reference%20Design%20v1.0Public%20Release.pdf)。

**Principle #1: Remove bottlenecks (including human ones) and manual actions.**

**原则 #1：消除瓶颈（包括人为瓶颈）和手动操作。**

With Kubernetes, developers and testers can work better, together. They can solve defects quickly and accurately because developers can use the tester’s Kubernetes instance for debugging. This eliminates long delays associated with replicating development and test environments. Kubernetes also helps testers and developers quickly exchange precise information for application configurations.

使用 Kubernetes，开发人员和测试人员可以更好地协同工作。他们可以快速准确地解决缺陷，因为开发人员可以使用测试人员的 Kubernetes 实例进行调试。这消除了与复制开发和测试环境相关的长时间延迟。 Kubernetes 还帮助测试人员和开发人员快速交换应用程序配置的精确信息。

**Principle #2: Automate as much of the development and deployment activity as possible.**

**原则 #2：尽可能多地自动化开发和部署活动。**

Kubernetes eliminates many of the manual provisioning and other time-consuming tasks of enterprise IT operations. In addition, the unified and automated orchestration approaches Kubernetes offers simplify multi-cloud management, enabling more services to be delivered with less work and fewer errors.

Kubernetes 消除了企业 IT 运营的许多手动配置和其他耗时任务。此外，Kubernetes 提供的统一和自动化编排方法简化了多云管理，从而能够以更少的工作和更少的错误交付更多服务。

**Principle #3: Adopt common tools from planning and requirements through deployment and operations.**

**原则 #3：采用从规划和需求到部署和运营的通用工具。**

Kubernetes offers many capabilities that allow one container to support many configuration environment contexts. The **configmaps** object, for example, supports configuration data that is used at runtime. This avoids the need for specialized containers for different environment configurations. Declarative syntax used to define the deployment state of Kubernetes-deployed container clusters greatly simplifies the management of the delivery and deployments.

Kubernetes 提供了许多功能，允许一个容器支持多个配置环境上下文。例如，**configmaps** 对象支持在运行时使用的配置数据。这避免了对不同环境配置的专用容器的需求。用于定义 Kubernetes 部署的容器集群的部署状态的声明性语法极大地简化了交付和部署的管理。

**Principle #4: Leverage agile software principles and favor small, incremental, frequent updates over larger, more sporadic releases.**

**原则 #4：利用敏捷软件原则，支持小的、增量的、频繁的更新，而不是更大、更零星的版本。**

Modular applications architected as microservices benefit the most from Kubernetes. Software designed in accordance with twelve-factor app tenets and communicating through networked APIs work best for scalable deployments on clusters. Kubernetes is optimal for orchestrating cloud-native applications. Modular distributed services are better able to scale and recover from failures.

架构为微服务的模块化应用程序从 Kubernetes 中获益最多。根据十二要素应用原则设计并通过联网 API 进行通信的软件最适合在集群上进行可扩展部署。 Kubernetes 最适合编排云原生应用程序。模块化分布式服务能够更好地扩展和从故障中恢复。

**Principle #5: Apply the cross-functional skill sets of development, cybersecurity and operations throughout the software life cycle, embracing a continuous monitoring approach in parallel instead of waiting to apply each skill set sequentially.** 

**原则 #5：在整个软件生命周期中应用开发、网络安全和运营的跨职能技能集，采用并行的持续监控方法，而不是等待依次应用每个技能集。**

Kubernetes provides a unified approach for container orchestration that applies end-to-end across the value stream. Continuous monitoring is facilitated because cloud-native applications, managed by Kubernetes, are constructed with health reporting metrics to enable the platform to manage life cycle events if an instance becomes unhealthy. They produce (and make available for export) robust telemetry data to alert operators to problems and allow them to make informed decisions. Kubernetes supports liveness and readiness probes that make it easy to determine the state of a containerized application.

Kubernetes 提供了一种统一的容器编排方法，可以在整个价值流中端到端地应用。由于由 Kubernetes 管理的云原生应用程序构建有健康报告指标，以便平台能够在实例变得不健康时管理生命周期事件，从而促进了持续监控。他们生成（并提供用于导出）强大的遥测数据，以提醒操作员注意问题并让他们做出明智的决定。 Kubernetes 支持活跃度和就绪度探测，可以轻松确定容器化应用程序的状态。

**Principle #6: Security risks of the underlying infrastructure must be measured and quantified, so that the total risks and impacts to software applications are understood.**

**原则 #6：必须衡量和量化底层基础设施的安全风险，以便了解对软件应用程序的总体风险和影响。**

Kubernetes has many different layers and components that must be considered for security. Elements of key concern for security are: API communications between different parts of a cluster, a scheduler that manages how workloads are distributed, controllers that manage the state of Kubernetes itself, agents that runs on each node within a cluster and a key-value store where cluster configuration data is housed. A multi-pronged defense strategy is needed to protect against all types of vulnerabilities. The following is partial list of defense strategies.

Kubernetes 有许多不同的层和组件，必须考虑安全性。安全的关键要素是：集群不同部分之间的 API 通信、管理工作负载分配方式的调度程序、管理 Kubernetes 本身状态的控制器、在集群内每个节点上运行的代理以及键值存储集群配置数据所在的位置。需要多管齐下的防御策略来防范所有类型的漏洞。以下是部分防御策略列表。

- Secure container images to run on Kubernetes. Use security code scanning tools to check the containerized code for vulnerabilities that can exist within the container code itself, as well as in any upstream dependencies on which the image is based.
- Isolate Kubernetes nodes on a separate network that is not be exposed directly to public networks.
- Kubernetes supports role-based access control (RBAC) policies to help guard against unauthorized access to cluster resources.
- Resource quotas help mitigate disruptions caused by denial-of-service attacks by depriving the rest of the cluster of sufficient resources to run.
- Restrict pod-to-pod traffic using Kubernetes core data types for specifying network access controls between pods.
- Implement network border controls to enforce some ingress and egress controls at the network border in addition to the pod-level controls enforced by Kubernetes.
- Application-layer access control can be hardened with strong application-layer authentication, such as mutual transport-level security protocols using encrypted application identity.
- Kubernetes support for multiple containers running together with a shared localhost network for the pod enables sidecars and a service mesh approach to retrofit existing applications. This reduces the difficulties of implementing mutual TLS solutions, so each application has an adjacent proxy daemon that terminates and authenticates inbound connections and transparently authenticates outbound connections.
- Segment your Kubernetes clusters by integrity level; for example, your dev and test environments might be hosted in a different cluster than your production environment.
- Run your applications as a non-root user. Future Linux kernel vulnerabilities are more likely to be exploitable by a root user than by a non-privileged user.
- Use security monitoring and auditing to capture application logs, host-level logs, Kubernetes API audit logs and cloud provider logs. For security audit purposes, consider streaming your logs to an external location with append-only access from within your cluster.
- Use process whitelisting to identify unexpected running processes.
- Keep Kubernetes versions up to date.



- 保护容器镜像以在 Kubernetes 上运行。使用安全代码扫描工具检查容器化代码中可能存在于容器代码本身以及映像所基于的任何上游依赖项中的漏洞。
- 将 Kubernetes 节点隔离在不直接暴露给公共网络的单独网络上。
- Kubernetes 支持基于角色的访问控制 (RBAC) 策略，以帮助防止对集群资源的未授权访问。
- 资源配额通过剥夺集群的其余部分足够的资源来运行，从而帮助减轻拒绝服务攻击造成的中断。
- 使用 Kubernetes 核心数据类型来限制 pod 到 pod 的流量，以指定 pod 之间的网络访问控制。
- 实施网络边界控制，以在网络边界执行一些入口和出口控制，以及由 Kubernetes 实施的 pod 级控制。
- 应用层访问控制可以通过强大的应用层身份验证来加强，例如使用加密应用程序身份的相互传输级安全协议。
- Kubernetes 支持多个容器一起运行，并为 pod 提供共享本地主机网络，从而启用 sidecar 和服务网格方法来改造现有应用程序。这降低了实现双向 TLS 解决方案的难度，因此每个应用程序都有一个相邻的代理守护程序，用于终止和验证入站连接并透明地验证出站连接。
- 按完整性级别划分您的 Kubernetes 集群；例如，您的开发和测试环境可能托管在与生产环境不同的集群中。
- 以非 root 用户身份运行您的应用程序。与非特权用户相比，未来的 Linux 内核漏洞更有可能被 root 用户利用。
- 使用安全监控和审计来捕获应用程序日志、主机级日志、Kubernetes API 审计日志和云提供商日志。出于安全审计目的，请考虑将日志流式传输到外部位置，并从集群内进行仅附加访问。
- 使用进程白名单来识别意外运行的进程。
- 使 Kubernetes 版本保持最新。

A comprehensive security strategy for Kubernetes need to include more than the handful of built-in security features.

Kubernetes 的综合安全策略需要包含的不仅仅是少量的内置安全功能。

**Principle #7: Deploy immutable infrastructure, such as containers.**

**原则 #7：部署不可变的基础设施，例如容器。**

The concept of immutable infrastructure supported by Kubernetes, in which deployed components are replaced in their entirety, rather than being updated in place, requires standardization and emulation of common infrastructure components to achieve consistent and predictable results. 

Kubernetes 支持的不可变基础设施的概念，即部署的组件被整体替换，而不是就地更新，需要对通用基础设施组件进行标准化和模拟，以实现一致和可预测的结果。

In this article, I explained how Kubernetes can be used and configured to satisfy DoD’s seven DevSecOps principles for a successful DevSecOps approach using Kubernetes. While Kubernetes provides built-in security tools, they are not sufficient to fully protect against multiple types of potential vulnerabilities across multiple layers of Kubernetes infrastructure. All seven DevSecOps principles are important for an integrated security strategy that mitigates threats at all layers and levels of your stack. 

在本文中，我解释了如何使用和配置 Kubernetes 来满足国防部的七项 DevSecOps 原则，从而成功地使用 Kubernetes 进行 DevSecOps 方法。虽然 Kubernetes 提供了内置的安全工具，但它们不足以完全保护 Kubernetes 基础架构的多层中的多种类型的潜在漏洞。所有七项 DevSecOps 原则对于减轻堆栈所有层和级别的威胁的集成安全策略都很重要。