# A Developer’s Guide to GitOps

# GitOps 开发人员指南

David Thor - 01/11/21 From: https://www.architect.io/blog/gitops-developers-guide

One of a modern DevOps team’s driving objectives is to help developers deploy features as quickly and safely as possible. This means creating tools and processes that do everything from provisioning private developer environments to deploying and securing production workloads. This effort is a constant balance between enabling developers to move quickly and ensuring that their haste doesn't lead to critical outages. Fortunately, both speed and stability improve tremendously whenever automation, like GitOps, is introduced.

现代 DevOps 团队的驱动目标之一是帮助开发人员尽可能快速、安全地部署功能。这意味着创建工具和流程来完成从配置私有开发人员环境到部署和保护生产工作负载的所有工作。这种努力是使开发人员能够快速行动和确保他们的匆忙不会导致严重中断之间的持续平衡。幸运的是，每当引入自动化（如 GitOps）时，速度和稳定性都会大大提高。

As you might have guessed from that lead-up, GitOps is a tactic for automating DevOps. More specifically, however, it's an automation tactic that hooks into a critical tool that already exists in developers’ everyday workflow, Git. Since developers are already committing code to a centralized Git repo (often hosted by tools like GitHub, GitLab, or BitBucket), DevOps engineers can wire up any of their operational scripts, like those used to build, test, or deploy applications, to kick off every time developers commit code changes. This means developers get to work exclusively with Git, and everything that helps them get their code to production will be automated behind the scenes.

正如您可能从那篇文章中猜到的，GitOps 是一种自动化 DevOps 的策略。然而，更具体地说，它是一种自动化策略，它与开发人员日常工作流程中已经存在的关键工具 Git 挂钩。由于开发人员已经将代码提交到集中的 Git 存储库（通常由 GitHub、GitLab 或 BitBucket 等工具托管），因此 DevOps 工程师可以连接他们的任何操作脚本，例如用于构建、测试或部署应用程序的脚本，以启动每次开发人员提交代码更改时关闭。这意味着开发人员可以专门使用 Git，而帮助他们将代码投入生产的一切都将在幕后自动化。

## Why GitOps?

## 为什么是 GitOps？

In years past, DevOps and CI/CD practices were a set of proprietary scripts and tools that executed everyday tasks like running tests, provisioning infrastructure, or deploying an application. However, the availability of new infrastructure tools like Kubernetes combined with the proliferation of microservice architectures have enabled and ultimately *demanded* that developers get more involved in CI/CD processes.

过去几年，DevOps 和 CI/CD 实践是一组专有脚本和工具，用于执行日常任务，例如运行测试、配置基础设施或部署应用程序。然而，像 Kubernetes 这样的新基础设施工具的可用性以及微服务架构的激增已经启用并最终“要求”开发人员更多地参与 CI/CD 流程。

This *shift left* exploded the problems seen with custom scripting and manual execution leading to confusing/inconsistent processes, duplication of efforts, and a drastic reduction in development velocity. To take advantage of cloud-native tools and architectures, teams need a consistent, automated approach to CI/CD that would enable developers to:

这种*左移*爆炸了自定义脚本和手动执行中看到的问题，导致流程混乱/不一致、工作重复以及开发速度急剧下降。为了利用云原生工具和架构，团队需要一种一致、自动化的 CI/CD 方法，使开发人员能够：

- Stop building and maintaining proprietary scripts and instead use a universal process
- Create apps and services faster by using said universal deploy process
- Onboard more quickly by deploying every time they make code changes
- Deploy automatically to make releases faster, more frequent, and more reliable
- Rollback and pass compliance audits with declarative design patterns


- 停止构建和维护专有脚本，而是使用通用流程
- 通过使用所述通用部署过程更快地创建应用程序和服务
- 通过在每次更改代码时进行部署来更快地加入
- 自动部署，使发布更快、更频繁、更可靠
- 使用声明式设计模式回滚并通过合规性审计


## Developers love GitOps

## 开发人员喜欢 GitOps

For all the reasons cited above (and more), businesses need manageable and automatable approaches to CI/CD and DevOps to succeed in building and maintaining cloud-native applications. However, if automation is all that’s needed, why GitOps over other strategies (e.g., SlackOps, scheduled deployments, or simple scripts)? The answer is simple: developers love GitOps.


由于上述（以及更多）原因，企业需要可管理和可自动化的 CI/CD 和 DevOps 方法，以成功构建和维护云原生应用程序。但是，如果只需要自动化，为什么 GitOps 优于其他策略（例如，SlackOps、计划部署或简单脚本）？答案很简单：开发人员喜欢 GitOps。


### One tool to rule them all, Git

### 一种工具来统治他们，Git

It's become apparent in the last few years that GitOps is among the most highly-rated strategies for automating DevOps by developers, and it's not hard to see why. Developers **live** in Git. They save temporary changes to git, collaborate using git, peer-review code using git, and store a history and audit trail of all the changes everyone has ever made in git. The pipelining strategy described above was tailor-made for git. Since developers already rely on git so heavily, these processes are, in turn, tailor-made for developers. Developers recognize this and are more than happy to reduce the tools and processes they need to use and follow to do their jobs.


在过去几年中，很明显 GitOps 是开发人员对 DevOps 自动化评价最高的策略之一，不难看出原因。开发人员**生活**在 Git 中。他们保存对 git 的临时更改，使用 git 进行协作，使用 git 对代码进行同行评审，并存储每个人在 git 中所做的所有更改的历史记录和审计跟踪。上面描述的流水线策略是为 git 量身定制的。由于开发人员已经非常依赖 git，因此这些流程又是为开发人员量身定制的。开发人员认识到这一点，并且非常乐意减少他们在完成工作时需要使用和遵循的工具和流程。


### Declared alongside code 

### 与代码一起声明

Beyond just the intuitive, git-backed execution flow, another part of modern CI tools and GitOps that developers love is the declarative design. The previous generation of CI tools had configurations that lived inside private instances of the tools. If you didn't have access to the tools, you didn't know what the pipelines did, if they were wrong or right, how or when they executed, or how to change them if needed. It was just a magic black box and hard for developers to trust as a result.

除了直观的、由 git 支持的执行流程之外，开发人员喜爱的现代 CI 工具和 GitOps 的另一部分是声明式设计。上一代 CI 工具的配置位于工具的私有实例中。如果您无法访问这些工具，您就不知道管道做了什么，它们是错误还是正确，它们如何或何时执行，或者如果需要如何更改它们。它只是一个神奇的黑匣子，因此开发人员很难信任。

In modern CI systems, like the ones most commonly used to power GitOps like [CircleCI](https://circleci.com/),[Github Actions](https://docs.github.com/en/free-pro- team@latest/actions), [Gitlab CI](https://about.gitlab.com/stages-devops-lifecycle/continuous-integration/), etc., the configurations powering the pipelines live directly in the Git repository. Just like the source code for the application, these configurations are version controlled and visible to every developer working on the project. Not only can they see what the pipeline process is, but they can also quickly and easily make changes to it as needed. This ease of access for developers is critical since developers write the tests for their applications and ensure it is safe and stable.


在现代 CI 系统中，例如最常用于支持 GitOps 的系统，如 [CircleCI](https://circleci.com/)、[GithubActions](https://docs.github.com/en/free-pro- team@latest/actions)、[Gitlab CI](https://about.gitlab.com/stages-devops-lifecycle/continuous-integration/) 等，支持管道的配置直接存在于 Git 存储库中。就像应用程序的源代码一样，这些配置是受版本控制的，并且对参与该项目的每个开发人员可见。他们不仅可以看到管道流程是什么，而且还可以根据需要快速轻松地对其进行更改。由于开发人员为其应用程序编写测试并确保其安全和稳定，因此开发人员的这种轻松访问至关重要。


### Completely self-service

### 完全自助服务

New features or bug fixes aren't considered complete until they land in production. This means that anything standing in the way of getting code changes to production are eating up developer time and mental energy when the feature, as far as the developer is concerned, "works on my machine.” Suppose developers have to wait, even for a few minutes, for a different team or individual to do some task before they can close out their work. In that case, it creates both friction and animosity in the organization.

新功能或错误修复在投入生产之前不被视为完成。这意味着当该功能（就开发人员而言）“在我的机器上运行”时，任何阻碍将代码更改到生产环境的事情都会消耗开发人员的时间和精力。假设开发人员必须等待，即使是几分钟，让不同的团队或个人完成某些任务才能结束他们的工作，在这种情况下，它会在组织中产生摩擦和敌意。

Alleviating this back and forth between teams is one of the main benefits of DevOps automation tactics like GitOps. Not only do developers get to work in a familiar tool, but the ability to have their code make its way to production without manual intervention means they are never waiting on someone else before they can complete their tasks.


减轻团队之间的这种来回是 DevOps 自动化策略（如 GitOps）的主要好处之一。开发人员不仅可以使用熟悉的工具工作，而且无需人工干预即可将代码投入生产的能力意味着他们永远不会在完成任务之前等待其他人。


### Continuous everything

### 连续一切

Yet another big perk of GitOps is that all the processes are continuously running all the time! Every change we make triggers tests builds, and deployments without ANY manual steps required. Since developers would use git with or without GitOps, hooking into their existing workflow to trigger DevOps processes is the perfect place to kick off automated events. Until developers stop using Git, GitOps will remain the ideal way to instrument automated DevOps.


GitOps 的另一大好处是所有进程一直在持续运行！我们所做的每一次更改都会触发测试构建和部署，而无需任何手动步骤。由于开发人员会在有或没有 GitOps 的情况下使用 git，因此挂钩到他们现有的工作流程以触发 DevOps 流程是启动自动化事件的理想场所。在开发人员停止使用 Git 之前，GitOps 仍将是检测自动化 DevOps 的理想方式。


## GitOps in practice

## GitOps 实践

Naturally, the involvement of developers in the process has led teams to explore the use of developer-friendly tools like Git, but the use of Git as a source of truth for DevOps processes also creates a natural consistency to the shape of CI/CD pipeline stages. There are only so many hooks available in a Git repository after all (e.g., commits, pull requests open/closed, merges, etc.), so the look and feel of most GitOps implementations include a set of typical stages:

自然地，开发人员在流程中的参与促使团队探索使用对开发人员友好的工具，如 Git，但使用 Git 作为 DevOps 流程的真实来源也为 CI/CD 管道的形状创造了自然的一致性阶段。毕竟，Git 存储库中可用的钩子只有这么多（例如，提交、打开/关闭拉取请求、合并等），因此大多数 GitOps 实现的外观和感觉包括一组典型的阶段：

![GitOps Pipelines](https://www.architect.io/images/blog/a-developers-guide-to-gitops/gitops-pipeline.png)


### 1. Pull requests, tests, and preview environments 

### 1. 拉取请求、测试和预览环境

After developers have spent time writing the code for their new feature, they generally commit that code to a new Git branch and submit a [pull request](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests) or [merge request](https://docs.gitlab.com/ee/user/project/merge_requests/getting_started.html) back to the mainline branch of the repository. This is something developers already do daily to prompt engineering managers to review the code changes and approve them to be merged into the main application code. Since developers already follow this kind of process for their daily collaboration efforts, it's a perfect opportunity for DevOps to wire up additional tasks.

在开发人员花时间为他们的新功能编写代码后，他们通常会将该代码提交到一个新的 Git 分支并提交一个[拉取请求](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests) 或 [合并请求](https://docs.gitlab.com/ee/user/project/merge_requests/getting_started.html) 返回到存储库的主线分支。这是开发人员每天都在做的事情，以提示工程经理审查代码更改并批准将它们合并到主应用程序代码中。由于开发人员已经在日常协作工作中遵循这种流程，因此对于 DevOps 来说，这是一个连接额外任务的绝佳机会。

By hooking into the open/close events created by this pull request process using a continuous integration (CI) tool, DevOps teams can trigger the execution of unit tests, creation of preview environments, and execution of integration tests against that new preview environment. Instrumentation of these steps allows engineering managers to establish trust in the code changes quickly and allows product managers to see the code changes via the preview environment before merging. Faster trust development means faster merges, and earlier input from product managers means easier changes without complicated and messy rollbacks. This GitOps hook is a key enabler for faster and healthier product and engineering teams alike.


通过使用持续集成 (CI) 工具连接到此拉取请求流程创建的打开/关闭事件，DevOps 团队可以触发单元测试的执行、预览环境的创建以及针对该新预览环境的集成测试的执行。这些步骤的检测允许工程经理快速建立对代码更改的信任，并允许产品经理在合并之前通过预览环境查看代码更改。更快的信任开发意味着更快的合并，而产品经理的早期输入意味着更容易的更改，而无需复杂和混乱的回滚。这个 GitOps 钩子是更快、更健康的产品和工程团队的关键推动因素。


### 2. Merge to master and deploy to staging

### 2. 合并到 master 并部署到 staging

Once all parties have reviewed the changes, the code can be merged into the mainline branch of the repository alongside changes from the rest of the engineering team. This mainline branch is often used as a staging ground for code that is almost ready to go to production, and as such, it’s another ideal time for us to run some operational tasks like tests and deployment. While we tested the code for each pull request before it was merged, we'll want to rerun tests to ensure that code works with the other changes contributed by peer team members. We'll also want to deploy all these changes to a shared environment (aka "staging") that the entire team can use to view and test the latest changes before they are released to customers.


一旦所有各方都审查了更改，代码就可以与工程团队其他成员的更改一起合并到存储库的主线分支中。这个主线分支通常用作几乎准备投入生产的代码的中转站，因此，这是我们运行一些操作任务（如测试和部署）的另一个理想时间。虽然我们在合并之前测试了每个拉取请求的代码，但我们希望重新运行测试以确保代码与对等团队成员贡献的其他更改一起工作。我们还希望将所有这些更改部署到一个共享环境（也称为“暂存”），整个团队可以使用该环境在最新更改发布给客户之前查看和测试它们。


### 3. Cut releases and deploy to production

### 3. 减少发布并部署到生产

Finally, after product and engineering have had time to review and test the latest changes to the mainline branch, teams are ready to cut a release and deploy to production! This is often a task performed by a release manager – a dedicated (or rotating) team member tasked with executing the deploy scripts and monitoring the release to ensure that nothing goes wrong in transit. Without GitOps, this team member would have to know where the proper scripts are, in what order to execute them, and would need to ensure their computer has all the correct libraries and packages required to power the scripts.

最后，在产品和工程部门有时间审查和测试对主线分支的最新更改后，团队已准备好发布并部署到生产！这通常是由发布经理执行的任务 - 一个专门的（或轮换的）团队成员，负责执行部署脚本和监控发布，以确保在传输过程中不会出现任何问题。如果没有 GitOps，该团队成员将必须知道正确的脚本在哪里，以何种顺序执行它们，并且需要确保他们的计算机具有支持脚本所需的所有正确库和包。

Thanks to GitOps, we can wire up this deployment to happen on another Git-based event – creating a [release](https://docs.github.com/en/free-pro-team@latest/github/administering-a-repository/about-releases) or tag. All a release manager would have to do is create a new "release,” often using semver for naming, and the tasks to build and deploy the code changes would be kicked off automatically. Like most tasks executed by a CI tool, these would be configured with the scripts' location and order the libraries and packages needed to execute them.


感谢 GitOps，我们可以将这个部署连接到另一个基于 Git 的事件上——创建一个 [发布](https://docs.github.com/en/free-pro-team@latest/github/administering-a-repository/about-releases) 或标签。发布经理所要做的就是创建一个新的“发布”，通常使用 semver 进行命名，构建和部署代码更改的任务将自动启动。与 CI 工具执行的大多数任务一样，这些将是配置脚本的位置并订购执行它们所需的库和包。


## GitOps tooling

## GitOps 工具

A solid and intuitive continuous integration tool isn't the only thing needed to instrument GitOps processes like those described in this article. The CI system can activate scripts based on git events, but you still need strong tools to power those scripts and ensure they can be run and maintained easily and safely. Deploying code changes (aka continuous delivery (CD)) is one of the most challenging steps to automate, so we've curated a few tooling categories that can help you through your GitOps journey:


一个可靠且直观的持续集成工具并不是检测本文中描述的 GitOps 流程所需的唯一工具。 CI 系统可以根据 git 事件激活脚本，但您仍然需要强大的工具来支持这些脚本并确保它们可以轻松安全地运行和维护。部署代码更改（又名持续交付 (CD)）是自动化最具挑战性的步骤之一，因此我们策划了一些工具类别，可以帮助您完成 GitOps 之旅：


### Containerization with Docker 

### Docker 容器化

Docker launched cloud development into an entirely new, distributed landscape and helped developers begin to realistically consider microservice architectures as a viable option. Part of what made Docker so powerful was how developer-friendly it is compared to the previous generation of virtualization solutions. Just like the declarative CI configurations that live inside our repositories, developers simply have to write and maintain a `Dockerfile` in their repository to enable automated container builds of deployable VMs. Containerization is an enormously powerful tactic for cloud-native teams and should be a staple tool in your repertoire.


Docker 将云开发引入了一个全新的分布式环境，并帮助开发人员开始切实考虑将微服务架构作为可行的选择。 Docker 如此强大的部分原因在于它与上一代虚拟化解决方案相比对开发人员的友好程度。就像我们存储库中的声明式 CI 配置一样，开发人员只需在他们的存储库中编写和维护一个“Dockerfile”，即可实现可部署虚拟机的自动化容器构建。容器化对于云原生团队来说是一种非常强大的策略，应该成为您的必备工具。


### Infrastructure-as-code (IaC)

### 基础设施即代码 (IaC)

A lot goes into provisioning infrastructure and deploying applications that isn't captured by a `Dockerfile`. For everything else, there's infrastructure-as-code (IaC) solutions like [Terraform](https://www.terraform.io/),[Cloudformation](https://aws.amazon.com/cloudformation/), and others. These solutions allow developers to describe the other bits of an application, like Kubernetes resources, load balancers, networking, security, and more, in a declarative way. Just like the CI configs and Dockerfiles described earlier, IaC templates can be version controlled and collaborated on by all the developers on your team.


很多用于配置基础设施和部署未被“Dockerfile”捕获的应用程序。对于其他一切，还有基础设施即代码 (IaC) 解决方案，例如 [Terraform](https://www.terraform.io/)、[Cloudformation](https://aws.amazon.com/cloudformation/) 和其他。这些解决方案允许开发人员以声明的方式描述应用程序的其他部分，如 Kubernetes 资源、负载均衡器、网络、安全性等。就像前面描述的 CI 配置和 Dockerfile 一样，IaC 模板可以由团队中的所有开发人员进行版本控制和协作。


### DevOps automation tools like Architect

### DevOps 自动化工具，如 Architect

I really can't talk about DevOps automation without talking about Architect. We love IaC and use it heavily as part of our product. We found that configuring deployments, networking, and network security, especially for microservice architectures, can be demanding on the developers who should be focused on new product features instead of infrastructure.

我真的不能不谈论架构师就谈论 DevOps 自动化。我们喜欢 IaC 并将其大量用作我们产品的一部分。我们发现，配置部署、网络和网络安全，尤其是对于微服务架构，对开发人员的要求可能更高，他们应该专注于新产品功能而不是基础设施。

Instead of writing IaC templates and CI pipelines, which require developers to learn about Kubernetes, Cilium, API gateways, managed databases, or other infrastructure solutions, just have them write an `architect.yml` file. We'll automatically deploy dependent APIs/databases and securely broker connectivity to them every time someone runs `architect deploy`. Our process can automatically spin up private developer environments, automated preview environments, and even production-grade cloud environments with just a single command.


无需编写 IaC 模板和 CI 管道，这需要开发人员了解 Kubernetes、Cilium、API 网关、托管数据库或其他基础架构解决方案，只需让他们编写一个 `architect.yml` 文件即可。每次有人运行“架构部署”时，我们都会自动部署相关的 API/数据库并安全地代理与它们的连接。只需一个命令，我们的流程就可以自动启动私有开发人员环境、自动预览环境，甚至生产级云环境。


## Learn more about DevOps, GitOps, and Architect!

## 了解有关 DevOps、GitOps 和架构师的更多信息！

At Architect, our mission is to help ops and engineering teams simply and efficiently collaborate and achieve deployment, networking, and security automation all at once. Ready to learn more? Check out these resources:

在 Architect，我们的使命是帮助运维和工程团队简单高效地协作，同时实现部署、网络和安全自动化。准备好了解更多了吗？查看这些资源：

- [Creating Microservices: Nest.js](https://www.architect.io/blog/creating-microservices-nestjs)
- [The Importance of Portability in Technology](https://www.architect.io/blog/the-importance-of-portability)
- [Our Product Docs!](https://www.architect.io/docs)

- [创建微服务：Nest.js](https://www.architect.io/blog/creating-microservices-nestjs)
- [技术中可移植性的重要性](https://www.architect.io/blog/the-importance-of-portability)
- [我们的产品文档！](https://www.architect.io/docs)

Or [sign up](https://cloud.architect.io/signup) and try Architect yourself today!

或 [注册](https://cloud.architect.io/signup) 并立即尝试自己构建架构！

------

Contents

内容

- Why GitOps?
- Developers love GitOps
- GitOps in practice
- GitOps tooling
- Learn more about DevOps, GitOps, and Architect! 

- 为什么是 GitOps？
- 开发人员喜欢 GitOps
- 实践中的 GitOps
- GitOps 工具
- 了解有关 DevOps、GitOps 和架构师的更多信息！

