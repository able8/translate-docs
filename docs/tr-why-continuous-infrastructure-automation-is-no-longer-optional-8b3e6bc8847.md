# Why Continuous Infrastructure Automation is no longer optional

# 为什么持续基础设施自动化不再是可选的

Nov 19, 2019

# Containerized micro services and service orchestration disrupted how we build and run apps

# 容器化微服务和服务编排扰乱了我们构建和运行应用程序的方式

We have now reached a critical point in cloud-based infrastructure evolution where several factors are piling on top of one another, calling for a game changing approach.

我们现在已经到了基于云的基础设施发展的关键点，几个因素相互叠加，需要一种改变游戏规则的方法。

First of all, Cloud Native design (containerized micro-services and service orchestration) is disrupting how we build and run apps. Transforming the landscape at an incredible pace, Kubernetes has become the new de-facto standard for container orchestration within one year of its graduation from the CNCF (Cloud Native Computing Foundation).

首先，云原生设计（容器化微服务和服务编排）正在颠覆我们构建和运行应用程序的方式。在从 CNCF（云原生计算基金会）毕业后的一年内，Kubernetes 以令人难以置信的速度改变了格局，已成为容器编排的新事实上的标准。

Indeed, we now have the ability to run multiple instances of micro-services and autoscale them with zero downtime. In the same time, multi-cloud is now a proven trend in Enterprise infrastructure, with over 60% of IaaS users deploying workloads on several clouds.

事实上，我们现在有能力运行多个微服务实例并在零停机时间自动扩展它们。与此同时，多云现在是企业基础设施中一个经过验证的趋势，超过 60% 的 IaaS 用户在多个云上部署工作负载。

While very powerful, it’s an insanely complex machinery, that necessitates huge skills to properly master both how it works and how to maintain it over time.

虽然非常强大，但它是一个极其复杂的机器，需要大量的技能来正确掌握它的工作方式和如何随着时间的推移维护它。

# As complexity now moves to the infrastructure level, infrastructure automation is no longer optional.

# 随着复杂性现在转移到基础设施级别，基础设施自动化不再是可选的。

Traditional architecture meant using a few simple servers behind a load balancer to host and run our applications. Since then, environments gradually evolved into orchestrated containerized architecture. As a consequence, complexity now moves to the infrastructure level.

传统架构意味着在负载均衡器后面使用几个简单的服务器来托管和运行我们的应用程序。从那时起，环境逐渐演变为精心编排的容器化架构。因此，复杂性现在转移到基础设施级别。

The logical outcome of that situation is that now, Infrastructure automation has become key to proper environment provisioning. And still, using IaaS consoles or simple deployment scripts, many teams jump to Kubernetes without automating their infrastructure first, putting their organization’s efficiency and security in jeopardy with unsafe, sometimes locked and/or un-repeatable environments.

这种情况的合乎逻辑的结果是，现在基础设施自动化已成为正确配置环境的关键。而且，仍然使用 IaaS 控制台或简单的部署脚本，许多团队在没有首先自动化他们的基础设施的情况下跳到 Kubernetes，将他们组织的效率和安全性置于不安全、有时锁定和/或不可重复的环境中。

Indeed, not only won’t “ClicOps” lead you ahead on the steep learning curve of using Infrastructure-as-Code, but it won’t either help you bridge the gap between your needs and a global context of DevOps skills shortening. Even more, those shortcuts often trigger additional problems due to the total lack of access right management or clear audit trails.

事实上，“ClicOps”不仅不会引导您在使用基础设施即代码的陡峭学习曲线上领先，而且也不会帮助您弥合您的需求与 DevOps 技能缩短的全球背景之间的差距。更重要的是，由于完全缺乏访问权限管理或清晰的审计跟踪，这些快捷方式通常会引发其他问题。

With organizations continuously trying to achieve a better time to value with quicker deployments, Infrastructure-as-Code provides the best option to provision, kill and relaunch environments on demand, using tools and code to provision and deploy servers and applications and have greater control of your infrastructure, especially in multi-cloud environments.

随着组织不断尝试通过更快的部署实现更好的时间价值，基础设施即代码提供了按需供应、终止和重新启动环境的最佳选择，使用工具和代码来供应和部署服务器和应用程序，并更好地控制您的基础架构，尤其是在多云环境中。

# So, how do you manage Kubernetes clusters properly with Infrastructure-as-Code?

# 那么，如何使用 Infrastructure-as-Code 正确管理 Kubernetes 集群？

Leveraging Hashicorp Terraform's abilities, we are building [CloudSkiff](https://www.cloudskiff.com/) in a GitOps approach to manage your infrastructure-as-code properly, like a CI/CD for complex infrastructures that you need to provision , autoscale, kill and relaunch in a seamless and repeatable process.

利用 Hashicorp Terraform 的能力，我们正在以 GitOps 方法构建 [CloudSkiff](https://www.cloudskiff.com/) 以正确管理您的基础设施即代码，例如您需要配置的复杂基础设施的 CI/CD 、自动缩放、终止并在无缝且可重复的过程中重新启动。

Consequently, the key features of [CloudSkiff](https://www.cloudskiff.com/) are :

- Generate a proper code for your infrastructure : either by writing it from scratch with an assistant code generator, or by reusing complex modules written by other people, or through a tool that allows you to turn your click-based environments into proper code that you can reuse and share.
- Organize the collaboration around your code : set up straightforward process and boundaries with RBAC (Role Based Access Control), to avoid troubles. You can’t have your junior Dev, or even several separate senior people push`terraform apply` or `terraform destroy` commands without proper control over what's happening.
- Manage testing environments with a sandbox just to make sure that things are running smoothly before you deploy to production.
- Post deployment, use dashboards and monitoring to keep control over performance and costs across geographies and environments and optimize your setup.
- Monitor your Kubernetes environments across all accounts and all clouds and watch over performances and costs for a better optimization.
- Duplicate environments and switch seamlessly between cloud providers with our migration assistant

因此，[CloudSkiff](https://www.cloudskiff.com/) 的主要特点是：

- 为您的基础设施生成合适的代码：通过使用辅助代码生成器从头开始编写，或通过重用其他人编写的复杂模块，或通过允许您将基于点击的环境转换为合适的代码的工具可以重复使用和分享。
- 围绕您的代码组织协作：使用 RBAC（基于角色的访问控制）设置简单的流程和边界，以避免出现问题。你不能让你的初级开发人员，甚至几个独立的高级人员在没有适当控制正在发生的事情的情况下推送“terraform apply”或“terraform destroy”命令。
- 使用沙箱管理测试环境，以确保在部署到生产之前一切顺利。
- 部署后，使用仪表板和监控来控制跨地域和环境的性能和成本，并优化您的设置。
- 跨所有帐户和所有云监控您的 Kubernetes 环境，并监控性能和成本以进行更好的优化。
- 使用我们的迁移助手复制环境并在云提供商之间无缝切换

# Why we go for Kubernetes as a first step, and why you need a tool to manage the lifecycle of your infrastructure? 

# 为什么我们首先选择 Kubernetes，为什么需要一个工具来管理基础设施的生命周期？

At [CloudSkiff](https://www.cloudskiff.com/) we decided to begin by supporting managed Kubernetes environments as a first step. Obviously, you've got to start somewhere, and going first for this framework that is (again) very powerful but insanely complex seemed to be a good way of providing support to the teams that might not have all the skills to master and maintain it overtime.

在 [CloudSkiff](https://www.cloudskiff.com/)，我们决定首先支持托管 Kubernetes 环境。显然，你必须从某个地方开始，首先使用这个（再次)非常强大但极其复杂的框架似乎是为可能没有掌握和维护它的所有技能的团队提供支持的好方法随着时间的推移。

A Kubernetes Cluster is a living being, killed and reborn several times a day, always moving, always changing, always adapting/autoscaling… It’s highly sensible to version upgrades. Basic configuration mistakes can yield drastic consequences, which is why you need a pipeline to manage its lifecycle. We believe that DevOps teams need a new generation of dedicated tools to manage the lifecycle of their Kubernetes Clusters and that this is what it takes to help them focus on building and running their apps.

Kubernetes 集群是一个活生生的人，每天被杀死和重生几次，一直在移动，一直在变化，一直在适应/自动缩放……版本升级非常明智。基本的配置错误会产生严重的后果，这就是为什么您需要一个管道来管理其生命周期。我们认为，DevOps 团队需要新一代专用工具来管理其 Kubernetes 集群的生命周期，而这正是帮助他们专注于构建和运行应用程序所需要的。

Question is : how do you implement that into existing workflows?

问题是：您如何将其实施到现有工作流程中？

Indeed, most existing workflows are “touchy” once setup and nobody really wants to change that. This is why [Cloudskiff](https://www.cloudskiff.com/) is built to be fully integrated within existing workflows. It is directly linked to the Github repositories of your choice and can even be called by your existing CI/CD tool though an API.

事实上，大多数现有的工作流程在设置后都是“敏感的”，没有人真的想改变它。这就是为什么 [Cloudskiff](https://www.cloudskiff.com/) 被构建为完全集成到现有工作流程中的原因。它直接链接到您选择的 Github 存储库，甚至可以通过 API 由您现有的 CI/CD 工具调用。

**CloudSkiff is an Infrastructure as Code that allows** Terraform automation and collaboration for growing teams.

**CloudSkiff 是一种基础架构即代码，允许** Terraform 自动化和协作以供不断壮大的团队使用。

[**Join our Beta here**](http://www.cloudskiff.com?utm_source=medium&utm_medium=WhyCIAisnolongeroptional)

[**在此处加入我们的测试版**](http://www.cloudskiff.com?utm_source=medium&utm_medium=WhyCIAisnolongeroptional)

![](https://miro.medium.com/max/1400/1*-Vb5CXecRO7May4txjsTYw.png)



