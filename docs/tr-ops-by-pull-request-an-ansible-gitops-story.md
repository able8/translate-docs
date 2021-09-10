# Ops by pull request: an Ansible GitOps story

# Ops by pull request：一个 Ansible GitOps 故事

March 30, 2020
by [Timothy Appnel](https://www.ansible.com/blog/author/timothy-appnel)



![ansible-blog_automated-webhooks-series](https://www.ansible.com/hs-fs/hubfs/Images/blog-social/ansible-blog_automated-webhooks-series.png?width=1035&name=ansible-blog_automated-webhooks-series.png)



In a [previous blog post](https://www.ansible.com/blog/intro-to-automation-webhooks-for-red-hat-ansible-automation-platform) I introduced Automation Webhooks and their uses with Infrastructure-as-Code (IaC) workflows and Red Hat Ansible Automation Platform. In this blog post, I’ll cover how those features can be applied to creating GitOps pipelines, a particular workflow gaining popularity in the cloud-native space, using Ansible and the unique benefits utilizing Ansible provides.

在 [上一篇博文](https://www.ansible.com/blog/intro-to-automation-webhooks-for-red-hat-ansible-automation-platform) 中，我介绍了自动化 Webhooks 及其与基础设施的使用- as-Code (IaC) 工作流和红帽 Ansible 自动化平台。在这篇博文中，我将介绍如何将这些功能应用于创建 GitOps 管道、使用 Ansible 以及利用 Ansible 提供的独特优势在云原生空间中越来越受欢迎的特定工作流程。

## What is GitOps?

## 什么是 GitOps？

Like so many terms that evolve and emerge from the insights and practices of what came before it, finding a definitive meaning to the term “GitOps” is a bit elusive.

就像许多从之前的见解和实践中演变和出现的术语一样，为“GitOps”一词找到明确的含义有点难以捉摸。

GitOps is a workflow whose conceptual roots started with [Martin Fowler's comprehensive Continuous Integration overview in 2006](https://martinfowler.com/articles/continuousIntegration.html) and descends from Site Reliability Engineering (SRE), DevOps culture and Infrastructure as Code (IaC) patterns. What makes it unique is that GitOps is a prescriptive style of Infrastructure as Code based on the experience and wisdom of what works in deploying and managing large, sophisticated, distributed and cloud-native systems. So you can implement git-centric workflows where you treat infrastructure like it is code, but it doesn’t mean it’s GitOps.

GitOps 是一个工作流程，其概念根源始于 [Martin Fowler 在 2006 年的全面持续集成概述](https://martinfowler.com/articles/continuousIntegration.html)，并源自站点可靠性工程(SRE)、DevOps 文化和基础设施即代码(IaC) 模式。它的独特之处在于，GitOps 是一种规范的基础设施即代码风格，它基于部署和管理大型、复杂、分布式和云原生系统的经验和智慧。因此，您可以实施以 git 为中心的工作流，在这种工作流中，您可以像对待代码一样对待基础设施，但这并不意味着它是 GitOps。

The term GitOps was coined by Alexis Richardson, CEO and Founder of Weaveworks, so a lot of how I’m going to define GitOps here comes directly from Alexis and Weaveworks. [This initial blog post](https://www.weave.works/blog/gitops-operations-by-pull-request) that explains the concept puts some baseline ideas out there, but doesn’t provide a concise definition. Depending who you ask you’ll get varying explanations of the term.

GitOps 一词是由 Weaveworks 的首席执行官兼创始人 Alexis Richardson 创造的，因此我将在这里定义 GitOps 的很多方法直接来自 Alexis 和 Weaveworks。 [这篇最初的博客文章](https://www.weave.works/blog/gitops-operations-by-pull-request) 解释了这个概念，提出了一些基本的想法，但没有提供简明的定义。根据您询问的对象，您会得到对该术语的不同解释。

After reading and listening to various explanations, I thought this definition was a concise one that captures the essence of the many ways that GitOps has been explained:

阅读和听了各种解释后，我认为这个定义是一个简洁的定义，它抓住了 GitOps 解释的多种方式的本质：

_[GitOps] works by using Git as a single source of truth for declarative infrastructure and applications._

_[GitOps] 使用 Git 作为声明性基础设施和应用程序的单一事实来源。_

\-\- Weaveworks, “Guide To GitOps”

\-\- Weaveworks，“GitOps 指南”

Often you will read immutable architectures, and even more specifically Kubernetes cluster management, as attributes of GitOps. I propose this description is a bit too prescriptive and limiting. You should embrace immutable architectures and Kubernetes in how you deploy applications and services if you can. That certainly makes a GitOps workflow easier to implement effectively, but here we will treat them as suggestions and preferences, not requirements of GitOps. We’ll see why in a bit.

通常，您会阅读不可变架构，更具体地说是 Kubernetes 集群管理，作为 GitOps 的属性。我认为这个描述有点过于规范和限制。如果可以，您应该在部署应用程序和服务的方式中采用不可变架构和 Kubernetes。这无疑使 GitOps 工作流程更容易有效实施，但在这里我们将它们视为建议和偏好，而不是 GitOps 的要求。我们稍后会看到原因。

This diagram shows what a typical GitOps workflow looks like from a conceptual level -- automating delivery pipelines to roll out changes to your infrastructure when changes are made in a Git repository.

此图从概念级别显示了典型的 GitOps 工作流的样子——当 Git 存储库中发生更改时，自动化交付管道以对基础设施进行更改。

![Screen Shot 2020-03-30 at 11.03.33 AM](https://www.ansible.com/hs-fs/hubfs/Screen%20Shot%202020-03-30%20at%2011.03.33%20AM.png?width=1584&name=Screen%20Shot%202020-03-30%20at%2011.03.33%20AM.png)



GitOps has been shown to increase productivity and velocity of deployments and development. Developers can use the tools and workflows they are already familiar with to manage deployments, allowing new developers to get up to speed faster. While Git has traditionally been a developer tool, operations staff benefit from the accumulated knowledge and experience of the Git community and the maturity of its ecosystem. There are a plethora of existing tools out there to make using Git more accessible and easier for those new to it.

GitOps 已被证明可以提高部署和开发的生产力和速度。开发人员可以使用他们已经熟悉的工具和工作流程来管理部署，让新开发人员能够更快地上手。虽然 Git 传统上是一种开发人员工具，但运维人员受益于 Git 社区积累的知识和经验以及其生态系统的成熟。有大量现有的工具可以让 Git 的新手更容易、更容易地使用它。

Using GitOps brings together deployments and operations with development processes and tooling, providing a consistent means of working in your organization. 

使用 GitOps 将部署和操作与开发流程和工具结合在一起，为您的组织提供一致的工作方式。

With the defined state of your infrastructure under Git version control, complete with a useful audit log of all activity, implementing a GitOps workflow will improve stability and increase reliability. Your organization benefits from being able to track what and when changes happen to help ensure compliance. Git tooling provides stronger security guarantees with its cryptography functions to track these changes and signing to prove its origins. When things go wrong, you can easily revert and roll back to the previous state or rebuild your systems if you need to recover from a disaster.

通过 Git 版本控制下定义的基础设施状态，以及所有活动的有用审计日志，实施 GitOps 工作流将提高稳定性并增加可靠性。您的组织受益于能够跟踪更改发生的内容和时间以帮助确保合规性。 Git 工具通过其加密功能提供更强大的安全保证，以跟踪这些更改并签名以证明其来源。当出现问题时，如果您需要从灾难中恢复，您可以轻松地恢复和回滚到以前的状态或重建系统。

## GitOps the Ansible Way

## GitOps 的 Ansible 方式

[Previously](https://www.ansible.com/blog/intro-to-automation-webhooks-for-red-hat-ansible-automation-platform), we reviewed how you can create Git-centric Infrastructure as Code ( IaC) deployment workflows with the Automation Webhooks capabilities in Ansible Tower. GitOps is a more prescriptive workflow of IaC though and conceptually looks something like this diagram.

[以前](https://www.ansible.com/blog/intro-to-automation-webhooks-for-red-hat-ansible-automation-platform)，我们回顾了如何创建以 Git 为中心的基础设施即代码 ( IaC) 部署工作流以及 Ansible Tower 中的自动化 Webhooks 功能。不过，GitOps 是 IaC 的更具规范性的工作流程，从概念上看，类似于下图。

![Screen Shot 2020-03-30 at 11.03.54 AM](https://www.ansible.com/hs-fs/hubfs/Screen%20Shot%202020-03-30%20at%2011.03.54%20AM.png?width=1516&name=Screen%20Shot%202020-03-30%20at%2011.03.54%20AM.png)There are a few things to note about using Red Hat Ansible Automation Platform compared to the typical GitOps pipeline.

与典型的 GitOps 管道相比，使用红帽 Ansible 自动化平台有几点需要注意。

Here Ansible Tower replaces the GitOps “agent” (Operator) that runs on a given cluster and pulls in its configuration (state) from Git. Typically these are Kubernetes Operators like [Flux](https://github.com/fluxcd/flux) or [Eunomia](https://github.com/KohlsTechnology/eunomia).

这里 Ansible Tower 替换了在给定集群上运行的 GitOps “代理”（Operator），并从 Git 中提取其配置（状态）。通常这些是 Kubernetes Operator，如 [Flux](https://github.com/fluxcd/flux) 或 [Eunomia](https://github.com/KohlsTechnology/eunomia)。

Ansible Tower can work with Operators running on a Kubernetes cluster for a push/pull sort of approach. Ansible Tower pushes the configuration to the Operators via a Custom Resource (CR) and then the Operator pulls in any container images from the registry and handles whatever setup is necessary. The  operators here are made for a specific Kubernetes application or service and are useful outside of a GitOps pipeline rather than one for all configurations and management of the cluster to facilitate GitOps.

Ansible Tower 可以与在 Kubernetes 集群上运行的 Operator 一起使用推/拉方式。 Ansible Tower 通过自定义资源 (CR) 将配置推送给 Operator，然后 Operator 从注册表中提取任何容器映像并处理任何必要的设置。这里的操作符是为特定的 Kubernetes 应用程序或服务而设计的，在 GitOps 管道之外很有用，而不是用于集群的所有配置和管理以促进 GitOps。

Using Ansible also provides the flexibility to apply GitOps workflow principles to systems other than Kubernetes, such as public/private cloud services and networking infrastructure, because you’re not required to use an Operator “agent” on that infrastructure.

使用 Ansible 还可以灵活地将 GitOps 工作流原则应用于 Kubernetes 以外的系统，例如公共/私有云服务和网络基础设施，因为您不需要在该基础设施上使用 Operator “代理”。

### Advantages of Using Ansible

### 使用 Ansible 的优势

There are many tools you can use in your GitOps pipelines; however, Ansible provides some unique advantages that make it ideal for these workflows and extending their use beyond Kubernetes and cloud-native systems.

您可以在 GitOps 管道中使用许多工具；然而，Ansible 提供了一些独特的优势，使其成为这些工作流的理想选择，并将其使用扩展到 Kubernetes 和云原生系统之外。

- **GitOps beyond Kubernetes.** Like Kubernetes, Ansible is a desired state engine that enables declarative modeling of traditional IT systems without scripting through Ansible Roles and Playbooks. With [the k8s module](https://docs.ansible.com/ansible/latest/modules/k8s_module.html) and [many others](https://docs.ansible.com/ansible/latest/modules/list_of_all_modules.html), an Ansible user can manage applications on Kubernetes, on existing IT infrastructure or across both with one simple language.
- **Agentless GitOps.** Consistent with the Ansible way of doing things, there is no requirement of a specialized GitOps Operator (agent) to facilitate the reconciliation of desired state on your systems.
- **Flexibility and Freedom.** You also have the flexibility and freedom to use the best tools for your needs and tailor your pipelines more to how you want to work. [Ansible excels as IT automation glue](https://www.ansible.com/blog/ansible-it-automation-glue).
- **Existing Skills & Ecosystem.** The same tried and trusted Ansible tooling lets you automate and orchestrate your applications across both new and existing platforms allowing teams to transition without having to learn new skills.

- **超越 Kubernetes 的 GitOps。** 与 Kubernetes 一样，Ansible 是一种理想的状态引擎，它支持对传统 IT 系统进行声明式建模，而无需通过 Ansible 角色和剧本编写脚本。使用 [k8s 模块](https://docs.ansible.com/ansible/latest/modules/k8s_module.html) 和 [许多其他](https://docs.ansible.com/ansible/latest/modules/list_of_all_modules.html)，Ansible 用户可以使用一种简单的语言在 Kubernetes、现有 IT 基础设施或两者之间管理应用程序。
- **无代理 GitOps。** 与 Ansible 的做事方式一致，不需要专门的 GitOps 操作员（代理）来促进系统上所需状态的协调。
- **灵活性和自由性。**您还可以灵活和自由地使用最佳工具来满足您的需求，并根据您想要的工作方式定制您的管道。 [Ansible 擅长作为 IT 自动化粘合剂](https://www.ansible.com/blog/ansible-it-automation-glue)。
- **现有技能和生态系统。** 相同的久经考验且值得信赖的 Ansible 工具可让您跨新平台和现有平台自动化和编排您的应用程序，从而使团队无需学习新技能即可过渡。

The benefits of using Ansible in this domain doesn’t end here though. There are a lot of [benefits for using Ansible in a cloud-native Kubernetes environments](https://www.ansible.com/blog/how-useful-is-ansible-in-a-cloud-native-kubernetes-environment).

在这个领域使用 Ansible 的好处还不止于此。 [在云原生 Kubernetes 环境中使用 Ansible 有很多好处](https://www.ansible.com/blog/how-useful-is-ansible-in-a-cloud-native-kubernetes-environment)。

## In Closing 

## 结束

GitOps works by using Git as a single source of truth for declarative infrastructure and applications. It is a workflow whose conceptual roots descend from Site Reliability Engineering, DevOps culture and Infrastructure as Code (IaC) patterns. GitOps is a more prescriptive workflow of IaC based on the experience and wisdom of what works in deploying and managing large sophisticated distributed systems.

GitOps 使用 Git 作为声明性基础设施和应用程序的单一事实来源。它是一个工作流程，其概念源于站点可靠性工程、DevOps 文化和基础设施即代码 (IaC) 模式。 GitOps 是一种更具规范性的 IaC 工作流程，它基于部署和管理大型复杂分布式系统的经验和智慧。

Using Red Hat Ansible Automation Platform to implement GitOps pipelines provides unique benefits. Utilizing the Automation Webhook capabilities in Ansible Tower, you can implement agentless GitOps workflows that go beyond just cloud-native systems and manage existing IT infrastructure such as cloud services and networking gear. Using Ansible, enables you to tap into the existing Ansible ecosystem and the flexibility and freedom to use the best tools for how you want to work.

使用红帽 Ansible 自动化平台实施 GitOps 管道可提供独特的优势。利用 Ansible Tower 中的自动化 Webhook 功能，您可以实施无代理 GitOps 工作流，而不仅仅是云原生系统，还可以管理现有的 IT 基础设施，例如云服务和网络设备。使用 Ansible，您可以利用现有的 Ansible 生态系统以及使用最佳工具的灵活性和自由度来满足您的工作需求。

We hope you give GitOps with Ansible a try and see how beneficial and powerful this can be to your organization. 

我们希望您尝试使用 Ansible 进行 GitOps，看看这对您的组织有多么有益和强大。

