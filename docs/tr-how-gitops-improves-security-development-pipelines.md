# How GitOps Improves the Security of Your Development Pipelines

# GitOps 如何提高开发管道的安全性

* * *

* * *

[GitOps](https://www.weave.works/technologies/gitops/) is usually discussed in terms of boosting developer velocity. But another benefit – one that doesn’t always get as much attention – concerns its potential to improve security.

[GitOps](https://www.weave.works/technologies/gitops/) 通常从提高开发人员速度的角度来讨论。但另一个好处——一个并不总是得到那么多关注的好处——涉及它提高安全性的潜力。

At our recent virtual event, GitOps Days 2020, Maya Kaczorowski ( [@MayaKaczorowski](https://twitter.com/mayakaczorowski?lang=es)), GitHub Product Manager for software supply chain security, shared her thoughts on how GitOps can boost the security of your entire development pipeline.

在我们最近的虚拟活动 GitOps Days 2020 中，GitHub 软件供应链安全产品经理 Maya Kaczorowski ([@MayaKaczorowski](https://twitter.com/mayakaczorowski?lang=es)) 分享了她对 GitOps 如何实现的想法提高整个开发管道的安全性。

To kick the session off, Maya talked about how DevOps and now GitOps leads to greater developer accountability. With the introduction of DevOps, developers took more responsibility for how their code was deployed and now, developers are becoming more responsible for security, too. Maya even suggested that modern DevOps can be thought of as ‘DevSecOps’. Whatever you want to call it, there’s no better way to give developers responsibility for operations and security than to adopt GitOps best practices.

在会议开始时，Maya 谈到了 DevOps 和现在的 GitOps 如何导致更大的开发人员责任。随着 DevOps 的引入，开发人员对其代码的部署方式承担了更多的责任，现在，开发人员也对安全性越来越负责。 Maya 甚至建议可以将现代 DevOps 视为“DevSecOps”。不管你想怎么称呼它，没有比采用 GitOps 最佳实践更好的方法来让开发人员负责运营和安全。

GitOps gives you control over changes and allows you to verify them from a single source:

GitOps 可让您控制更改并允许您从单一来源验证它们：

1. **Config as Code**

1. **配置为代码**

Using Git to manage YAML files makes it simple to check if you’re meeting security requirements. With access policies declared in a config file, you know who has access to what – and can easily verify it in code.
2. **Changes are** **auditable**

使用 Git 管理 YAML 文件可以轻松检查您是否满足安全要求。通过在配置文件中声明的访问策略，您知道谁可以访问什么——并且可以轻松地在代码中验证它。
2. **变化是** **可审计**

Version control means that you always know what you shipped and you can roll back at any time. Your commit history is an audit trail of comments, reviews, and a history of decisions that were made to your repo.
3. **Production matches the desired state kept in Git**

版本控制意味着您始终知道自己发布了什么，并且可以随时回滚。你的提交历史是对你的 repo 做出的评论、评论和决策历史的审计跟踪。
3. **生产与 Git 中保存的所需状态相匹配**

A single source of truth, with a common workflow for both code and infrastructure changes coupled with automatic alerts on a drift from the desired states increases reliability and removes the risk of human error. A single set of tests, security scans, and permissions also help make changes secure and reliable.

单一的事实来源，代码和基础架构更改的通用工作流程，再加上偏离所需状态时的自动警报，可提高可靠性并消除人为错误的风险。一组测试、安全扫描和权限也有助于使更改安全可靠。

![gitops-security.png](https://images.contentstack.io/v3/assets/blt300387d93dabf50e/blt1fa86e3d1618bc36/5fd92c867c43e43bf41983b3/gitops-security.png)

## Continuous security

## 持续安全

Maya explained that tasks usually left until the end of the development cycle – like security testing – take place much earlier with GitOps. Testing for security is part of every iteration, which means errors can be caught earlier and vulnerabilities eliminated, long before any code is deployed to production. It therefore allows for what Maya calls ‘continuous security’. In parallel with continuous integration and continuous deployment, she sees this as a focus for integrating security best practice throughout the pipeline. Another way of looking at it is that while DevOps made developers responsible for any system outages they caused, treating security the same way makes developers responsible for data loss. And ultimately, the business goal of any security activity is preventing data loss.

Maya 解释说，任务通常会留到开发周期结束时——比如安全测试——在 GitOps 中发生得更早。安全性测试是每次迭代的一部分，这意味着可以在任何代码部署到生产之前很早地发现错误并消除漏洞。因此，它允许 Maya 所谓的“持续安全”。在持续集成和持续部署的同时，她认为这是在整个管道中集成安全最佳实践的重点。另一种看待它的方式是，虽然 DevOps 让开发人员对他们造成的任何系统中断负责，但以同样的方式对待安全性会让开发人员对数据丢失负责。最终，任何安全活动的业务目标都是防止数据丢失。

## How GitOps makes this possible?

## GitOps 如何使这成为可能？

It comes down to the ability of GitOps tools to treat everything as code. If you can treat all configuration and security policy as code, everything can be held in version control. Changes can be made and reviewed , then fed through the automated pipeline that verifies, deploys and monitors the change. Any divergence from the desired state of the system – e.g. the emergence of a bug or a security vulnerability – will be caught much earlier, long before it can become a significant cost to the business.

这归结为 GitOps 工具将一切视为代码的能力。如果您可以将所有配置和安全策略视为代码，那么一切都可以在版本控制中进行。可以进行和审查更改，然后通过验证、部署和监控更改的自动化管道进行反馈。与系统期望状态的任何偏离——例如错误或安全漏洞的出现 - 会更早地被发现，早在它成为企业的重大成本之前。

## Securing the whole pipeline 

## 保护整个管道

GitOps improves the security of several elements on the development pipeline, from the code itself (and that includes anything else you keep in code, such as policy and config) to the process by which you make changes to it. So for example, if your compliance requirements are defined in YAML, Git ensures you’re meeting them by maintaining the desired state of the system. . With all changes captured in version control, you can roll back to a given point if you need to. The history of comments and reviews is also maintained, so not only will you know who made changes and when, but you’ll also know why. It’s a built-in audit trail.

GitOps 提高了开发管道中多个元素的安全性，从代码本身（包括您保留在代码中的任何其他内容，例如策略和配置）到您对其进行更改的过程。例如，如果您的合规性要求是在 YAML 中定义的，Git 会通过维护系统的所需状态来确保您满足这些要求。 .通过版本控制中捕获的所有更改，您可以根据需要回滚到给定点。评论和评论的历史记录也会得到维护，因此您不仅会知道谁进行了更改以及何时进行了更改，而且还会知道原因。这是一个内置的审计跟踪。

When you get to production, the big security advantage is the single source of truth Git provides. It means you have a single set of tests, a single set of security scans and a single set of permissions to implement. Best of all, humans play no part in this process, which means human error is eliminated.

当您进入生产环境时，最大的安全优势是 Git 提供的单一事实来源。这意味着您需要执行一组测试、一组安全扫描和一组权限。最重要的是，人类在此过程中没有任何作用，这意味着消除了人为错误。

If your application does come under attack, you can take action immediately. Git provides the single source of truth, so you can redeploy everything instantly, if you need to. And of course, you can place controls and gates on this process, to meet any security or compliance needs specific to your organization.

如果您的应用程序确实受到攻击，您可以立即采取行动。 Git 提供了单一的事实来源，因此您可以在需要时立即重新部署所有内容。当然，您可以在此过程中设置控制和门，以满足特定于您的组织的任何安全或合规性需求。

## Respond faster

## 响应更快

One of the key benefits of GitOps, however, is velocity. [the State of DevOps report](https://puppet.com/resources/report/2020-state-of-devops-report/) has proven that developers can move faster thanks to version control, continuous integration, test automation and other features available in Git. But that velocity isn’t limited to the speed at which developers can work. Maya suggested another take on velocity – one that can be thought of as an alternative to [Mean Time to Recovery](https://en.wikipedia.org/wiki/Mean_time_to_recovery): [Mean Time To Remediate](https://www.optiv.com/cybersecurity-dictionary/mttr-mean-time-to-respond-remediate). And with its auditable and full record of who did what, when nothing cuts Mean Time To Remediate like GitOps.

然而，GitOps 的主要优势之一是速度。 [DevOps 现状报告](https://puppet.com/resources/report/2020-state-of-devops-report/) 已经证明，由于版本控制、持续集成、测试自动化等，开发人员可以更快地行动Git 中可用的功能。但这种速度不仅限于开发人员的工作速度。 Maya 建议了另一种速度——可以将其视为 [平均恢复时间](https://en.wikipedia.org/wiki/Mean_time_to_recovery) 的替代方案：[平均修复时间](https://www.optiv.com/cybersecurity-dictionary/mttr-mean-time-to-respond-remediate)。并且具有可审计且完整的记录谁做了什么，当没有什么能像 GitOps 那样减少平均修复时间时。

Development pipelines clearly represent a tempting attack vector for intruders. But by adopting GitOps, you can boost security right along the pipeline, while at the same time boosting the speed at which you can react if the worst happens.

开发管道显然代表了入侵者的诱人攻击向量。但是通过采用 GitOps，您可以直接提高管道的安全性，同时提高在最坏情况发生时的反应速度。

View the full presentation:

查看完整演示：

[Go from Zero to GitOps with our discovery, design and deploy package for Kubernetes.](https://www.weave.works/services/gitops-design-services/)

[使用我们的 Kubernetes 发现、设计和部署包从零到 GitOps。](https://www.weave.works/services/gitops-design-services/)

* * *

* * *

## About

##  关于

![ ](https://images.contentstack.io/v3/assets/blt300387d93dabf50e/blt1be4b5b42ea58cb4/58c02d7b48598d51743bf27e/weave-logo-512.png?format=webp&width=75)

Weaveworks’ mission is to empower developers and DevOps teams to build better software faster. Our “GitOps” model strives to optimize operational workflows; to make operations for developers simpler, better and faster. 

Weaveworks 的使命是让开发人员和 DevOps 团队能够更快地构建更好的软件。我们的“GitOps”模型致力于优化运营工作流程；使开发人员的操作更简单、更好、更快。

