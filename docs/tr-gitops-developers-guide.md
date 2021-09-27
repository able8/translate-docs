# A Developer’s Guide to GitOps

# GitOps 开发人员指南

David Thor \- 01/11/21

大卫·托尔 \- 01/11/21

One of a modern DevOps team’s driving objectives is to help developers deploy
features as quickly and safely as possible. This means creating tools and
processes that do everything from provisioning private developer environments to
deploying and securing production workloads. This effort is a constant balance
between enabling developers to move quickly and ensuring that their haste
doesn't lead to critical outages. Fortunately, both speed and stability improve
tremendously whenever automation, like GitOps, is introduced.

现代 DevOps 团队的驱动目标之一是帮助开发人员部署
功能尽可能快速和安全。这意味着创建工具和
从提供私有开发人员环境到
部署和保护生产工作负载。这种努力是一种不断的平衡
在使开发人员快速行动和确保他们的快速行动之间
不会导致严重中断。幸运的是，速度和稳定性都提高了
每当引入自动化（如 GitOps）时，效果都会非常好。

As you might have guessed from that lead-up, GitOps is a tactic for automating
DevOps. More specifically, however, it's an automation tactic that hooks into a
critical tool that already exists in developers’ everyday workflow, Git. Since
developers are already committing code to a centralized Git repo (often hosted
by tools like GitHub, GitLab, or BitBucket), DevOps engineers can wire up any of
their operational scripts, like those used to build, test, or deploy
applications, to kick off every time developers commit code changes. This means
developers get to work exclusively with Git, and everything that helps them get
their code to production will be automated behind the scenes.

正如你可能已经猜到的那样，GitOps 是一种自动化策略
开发运营。然而，更具体地说，它是一种自动化策略，可以与
开发人员日常工作流程中已经存在的关键工具 Git。自从
开发人员已经将代码提交到集中的 Git 存储库（通常托管在
通过 GitHub、GitLab 或 BitBucket 等工具），DevOps 工程师可以连接任何
他们的操作脚本，例如用于构建、测试或部署的脚本
应用程序，在开发人员每次提交代码更改时启动。这意味着
开发人员可以专门使用 Git，以及可以帮助他们获得的一切
他们的生产代码将在幕后自动化。

## Why GitOps? [why gitops permalink](http://www.architect.io\#why-gitops)

## 为什么是 GitOps？ [为什么 gitops 永久链接](http://www.architect.io\#why-gitops)

In years past, DevOps and CI/CD practices were a set of proprietary scripts and
tools that executed everyday tasks like running tests, provisioning
infrastructure, or deploying an application. However, the availability of new
infrastructure tools like Kubernetes combined with the proliferation of
microservice architectures have enabled and ultimately _demanded_ that
developers get more involved in CI/CD processes.

过去几年，DevOps 和 CI/CD 实践是一组专有脚本和
执行日常任务的工具，如运行测试、配置
基础设施，或部署应用程序。然而，新的可用性
Kubernetes 等基础设施工具与
微服务架构已经启用并最终_要求_
开发人员更多地参与 CI/CD 流程。

This _shift left_ exploded the problems seen with custom scripting and manual
execution leading to confusing/inconsistent processes, duplication of efforts,
and a drastic reduction in development velocity. To take advantage of
cloud-native tools and architectures, teams need a consistent, automated
approach to CI/CD that would enable developers to:

这个 _shift left_ 爆炸了自定义脚本和手册中看到的问题
执行导致流程混乱/不一致，重复工作，
并且开发速度急剧下降。从中获利
云原生工具和架构，团队需要一致的、自动化的
使开发人员能够：

- Stop building and maintaining proprietary scripts and instead use a universal
process
- Create apps and services faster by using said universal deploy process
- Onboard more quickly by deploying every time they make code changes
- Deploy automatically to make releases faster, more frequent, and more reliable
- Rollback and pass compliance audits with declarative design patterns

- 停止构建和维护专有脚本，而是使用通用脚本
过程
- 通过使用所述通用部署过程更快地创建应用程序和服务
- 通过在每次更改代码时进行部署来更快地加入
- 自动部署，使发布更快、更频繁、更可靠
- 使用声明式设计模式回滚并通过合规性审计

## Developers love GitOps [developers love gitops permalink](http://www.architect.io\#developers-love-gitops)

## 开发者喜欢 GitOps [开发者喜欢 gitops 永久链接](http://www.architect.io\#developers-love-gitops)

For all the reasons cited above (and more), businesses need manageable and
automatable approaches to CI/CD and DevOps to succeed in building and
maintaining cloud-native applications. However, if automation is all that’s
needed, why GitOps over other strategies (e.g., SlackOps, scheduled deployments,
or simple scripts)? The answer is simple: developers love GitOps.

由于上述所有（以及更多）原因，企业需要易于管理和
CI/CD 和 DevOps 的自动化方法，以成功构建和
维护云原生应用程序。但是，如果自动化仅此而已
需要，为什么 GitOps 优于其他策略（例如，SlackOps、计划部署、
或简单的脚本）？答案很简单：开发人员喜欢 GitOps。

### One tool to rule them all, Git [one tool to rule them all git permalink](http://www.architect.io\#one-tool-to-rule-them-all-git)

### 一种工具来统治它们，Git [一种工具来统治它们 git 永久链接](http://www.architect.io\#one-tool-to-rule-them-all-git)

It's become apparent in the last few years that GitOps is among the most
highly-rated strategies for automating DevOps by developers, and it's not hard
to see why. Developers **live** in Git. They save temporary changes to git,
collaborate using git, peer-review code using git, and store a history and audit
trail of all the changes everyone has ever made in git. The pipelining strategy
described above was tailor-made for git. Since developers already rely on git so
heavily, these processes are, in turn, tailor-made for developers. Developers
recognize this and are more than happy to reduce the tools and processes they
need to use and follow to do their jobs.

在过去的几年里，很明显 GitOps 是其中最多的
开发人员自动化 DevOps 的高度评价策略，这并不难
看看为什么。开发人员**生活**在 Git 中。他们保存对 git 的临时更改，
使用 git 进行协作，使用 git 进行同行评审代码，并存储历史记录和审计
每个人在 git 中所做的所有更改的踪迹。流水线策略
上面描述的是为 git 量身定做的。由于开发人员已经依赖 git 所以
反过来，这些流程是为开发人员量身定制的。开发商
认识到这一点，并且非常乐意减少他们使用的工具和流程
需要使用和遵循来完成他们的工作。

### Declared alongside code [declared alongside code permalink](http://www.architect.io\#declared-alongside-code)

### 与代码一起声明 [与代码一起声明永久链接](http://www.architect.io\#declared-alongside-code)

Beyond just the intuitive, git-backed execution flow, another part of modern CI
tools and GitOps that developers love is the declarative design. The previous
generation of CI tools had configurations that lived inside private instances of
the tools. If you didn't have access to the tools, you didn't know what the
pipelines did, if they were wrong or right, how or when they executed, or how to
change them if needed. It was just a magic black box and hard for developers to
trust as a result.

除了直观的、由 git 支持的执行流程之外，现代 CI 的另一部分
开发人员喜欢的工具和 GitOps 是声明式设计。以前的
一代 CI 工具的配置存在于私有实例中
工具。如果您无法使用这些工具，您就不会知道
管道做了，如果它们是错误的或正确的，它们是如何或何时执行的，或者如何
如果需要，请更改它们。这只是一个神奇的黑匣子，开发人员很难
结果信任。

In modern CI systems, like the ones most commonly used to power GitOps like
[CircleCI](https://circleci.com/), 

在现代 CI 系统中，就像最常用于支持 GitOps 的系统一样
[CircleCI](https://circleci.com/),

[Github Actions](https://docs.github.com/en/free-pro-team@latest/actions),
[Gitlab CI](https://about.gitlab.com/stages-devops-lifecycle/continuous-integration/),
etc., the configurations powering the pipelines live directly in the Git
repository. Just like the source code for the application, these configurations
are version controlled and visible to every developer working on the project.
Not only can they see what the pipeline process is, but they can also quickly
and easily make changes to it as needed. This ease of access for developers is
critical since developers write the tests for their applications and ensure it
is safe and stable.

[Github Actions](https://docs.github.com/en/free-pro-team@latest/actions)，
[Gitlab CI](https://about.gitlab.com/stages-devops-lifecycle/continuous-integration/),
等等，支持管道的配置直接存在于 Git 中
存储库。就像应用程序的源代码一样，这些配置
是版本控制的，并且对参与该项目的每个开发人员可见。
他们不仅可以看到流水线流程是什么，而且还可以快速
并根据需要轻松对其进行更改。开发人员的这种访问便利性是
至关重要，因为开发人员为他们的应用程序编写测试并确保它
安全稳定。

### Completely self-service [completely self service permalink](http://www.architect.io\#completely-self-service)

### 完全自助[完全自助永久链接](http://www.architect.io\#completely-self-service)

New features or bug fixes aren't considered complete until they land in
production. This means that anything standing in the way of getting code changes
to production are eating up developer time and mental energy when the feature,
as far as the developer is concerned, "works on my machine.” Suppose developers
have to wait, even for a few minutes, for a different team or individual to do
some task before they can close out their work. In that case, it creates both
friction and animosity in the organization.

新功能或错误修复在登陆之前不会被认为是完整的
生产。这意味着任何阻碍获取代码更改的东西
到生产正在消耗开发人员的时间和精力时，功能，
就开发人员而言，“在我的机器上工作”。假设开发人员
必须等待，即使是几分钟，让不同的团队或个人去做
一些任务，然后才能结束他们的工作。在这种情况下，它会同时创建
组织中的摩擦和敌意。

Alleviating this back and forth between teams is one of the main benefits of
DevOps automation tactics like GitOps. Not only do developers get to work in a
familiar tool, but the ability to have their code make its way to production
without manual intervention means they are never waiting on someone else before
they can complete their tasks.

减轻团队之间的这种来回是主要好处之一
DevOps 自动化策略，如 GitOps。开发人员不仅可以在
熟悉的工具，但能够将其代码用于生产
没有人工干预意味着他们以前永远不会等待别人
他们可以完成他们的任务。

### Continuous everything [continuous everything permalink](http://www.architect.io\#continuous-everything)

### 连续一切[连续一切永久链接](http://www.architect.io\#continuous-everything)

Yet another big perk of GitOps is that all the processes are continuously
running all the time! Every change we make triggers tests builds, and
deployments without ANY manual steps required. Since developers would use git
with or without GitOps, hooking into their existing workflow to trigger DevOps
processes is the perfect place to kick off automated events. Until developers
stop using Git, GitOps will remain the ideal way to instrument automated DevOps.

GitOps 的另一个大好处是所有的过程都是连续的
一直在跑！我们所做的每一次更改都会触发测试构建，并且
无需任何手动步骤即可部署。由于开发人员会使用 git
有或没有 GitOps，挂钩到他们现有的工作流程以触发 DevOps
流程是启动自动化事件的理想场所。直到开发商
停止使用 Git，GitOps 仍将是检测自动化 DevOps 的理想方式。

## GitOps in practice [gitops in practice permalink](http://www.architect.io\#gitops-in-practice)

## GitOps 实践 [gitops 实践永久链接](http://www.architect.io\#gitops-in-practice)

Naturally, the involvement of developers in the process has led teams to explore
the use of developer-friendly tools like Git, but the use of Git as a source of
truth for DevOps processes also creates a natural consistency to the shape of
CI/CD pipeline stages. There are only so many hooks available in a Git
repository after all (e.g., commits, pull requests open/closed, merges, etc.),
so the look and feel of most GitOps implementations include a set of typical
stages:

自然地，开发人员在过程中的参与促使团队探索
使用对开发人员友好的工具，如 Git，但使用 Git 作为源
DevOps 流程的真相也为
CI/CD 流水线阶段。 Git 中可用的钩子只有这么多
毕竟存储库（例如，提交、拉取请求打开/关闭、合并等），
所以大多数 GitOps 实现的外观和感觉包括一组典型的
阶段：

![GitOps Pipelines](http://www.architect.io/images/blog/a-developers-guide-to-gitops/gitops-pipeline.png)

### 1\. Pull requests, tests, and preview environments [1 pull requests tests and preview environments permalink](http://www.architect.io\#1-pull-requests-tests-and-preview-environments)

### 1. 拉取请求、测试和预览环境 [1 个拉取请求测试和预览环境永久链接](http://www.architect.io\#1-pull-requests-tests-and-preview-environments)

After developers have spent time writing the code for their new feature, they
generally commit that code to a new Git branch and submit a
[pull request](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests)
or
[merge request](https://docs.gitlab.com/ee/user/project/merge_requests/getting_started.html)
back to the mainline branch of the repository. This is something developers
already do daily to prompt engineering managers to review the code changes and
approve them to be merged into the main application code. Since developers
already follow this kind of process for their daily collaboration efforts, it's
a perfect opportunity for DevOps to wire up additional tasks.

在开发人员花时间为他们的新功能编写代码后，他们
通常将该代码提交到新的 Git 分支并提交
[拉取请求](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests)
或者
[合并请求](https://docs.gitlab.com/ee/user/project/merge_requests/getting_started.html)
回到存储库的主线分支。这是开发人员的东西
已经每天都做，以提示工程经理审查代码更改和
批准它们合并到主应用程序代码中。由于开发商
已经在他们的日常协作工作中遵循这种流程，这是
DevOps 连接额外任务的绝佳机会。

By hooking into the open/close events created by this pull request process using
a continuous integration (CI) tool, DevOps teams can trigger the execution of
unit tests, creation of preview environments, and execution of integration tests
against that new preview environment. Instrumentation of these steps allows
engineering managers to establish trust in the code changes quickly and allows
product managers to see the code changes via the preview environment before
merging. Faster trust development means faster merges, and earlier input from
product managers means easier changes without complicated and messy rollbacks. 

通过使用挂钩到此拉取请求过程创建的打开/关闭事件
持续集成 (CI) 工具，DevOps 团队可以触发执行
单元测试、预览环境的创建和集成测试的执行
针对那个新的预览环境。这些步骤的检测允许
工程经理建立对代码更改快速的信任并允许
产品经理之前通过预览环境查看代码更改
合并。更快的信任发展意味着更快的合并和更早的输入
产品经理意味着更轻松的更改，而无需复杂和混乱的回滚。

This GitOps hook is a key enabler for faster and healthier product and
engineering teams alike.

这个 GitOps 钩子是更快、更健康的产品和
工程团队一样。

### 2\. Merge to master and deploy to staging [2 merge to master and deploy to staging permalink](http://www.architect.io\#2-merge-to-master-and-deploy-to-staging)

### 2. 合并到主并部署到暂存 [2 合并到主并部署到暂存永久链接](http://www.architect.io\#2-merge-to-master-and-deploy-to-staging)

Once all parties have reviewed the changes, the code can be merged into the
mainline branch of the repository alongside changes from the rest of the
engineering team. This mainline branch is often used as a staging ground for
code that is almost ready to go to production, and as such, it’s another ideal
time for us to run some operational tasks like tests and deployment. While we
tested the code for each pull request before it was merged, we'll want to rerun
tests to ensure that code works with the other changes contributed by peer team
members. We'll also want to deploy all these changes to a shared environment
(aka "staging") that the entire team can use to view and test the latest changes
before they are released to customers.

一旦所有各方都审查了更改，就可以将代码合并到
存储库的主线分支以及其余部分的更改
工程团队。此主线分支通常用作
几乎可以投入生产的代码，因此，这是另一个理想
是时候让我们运行一些操作任务了，比如测试和部署。虽然我们
在合并之前测试了每个拉取请求的代码，我们想要重新运行
测试以确保代码与同行团队贡献的其他更改一起工作
成员。我们还希望将所有这些更改部署到共享环境中
（又名“暂存”）整个团队可以用来查看和测试最新的更改
在它们发布给客户之前。

### 3\. Cut releases and deploy to production [3 cut releases and deploy to production permalink](http://www.architect.io\#3-cut-releases-and-deploy-to-production)

### 3. 剪切发布并部署到生产 [3 剪切发布并部署到生产永久链接](http://www.architect.io\#3-cut-releases-and-deploy-to-production)

Finally, after product and engineering have had time to review and test the
latest changes to the mainline branch, teams are ready to cut a release and
deploy to production! This is often a task performed by a release manager – a
dedicated (or rotating) team member tasked with executing the deploy scripts and
monitoring the release to ensure that nothing goes wrong in transit. Without
GitOps, this team member would have to know where the proper scripts are, in
what order to execute them, and would need to ensure their computer has all the
correct libraries and packages required to power the scripts.

最后，在产品和工程部门有时间审查和测试
主线分支的最新变化，团队已准备好削减发布和
部署到生产！这通常是由发布经理执行的任务——
专门（或轮换）团队成员负责执行部署脚本和
监控发布以确保在运输过程中不会出现任何问题。没有
GitOps，这个团队成员必须知道正确的脚本在哪里，在
执行它们的顺序是什么，并且需要确保他们的计算机具有所有
为脚本提供动力所需的正确库和包。

Thanks to GitOps, we can wire up this deployment to happen on another Git-based
event – creating a
[release](https://docs.github.com/en/free-pro-team@latest/github/administering-a-repository/about-releases)
or tag. All a release manager would have to do is create a new "release,” often
using semver for naming, and the tasks to build and deploy the code changes
would be kicked off automatically. Like most tasks executed by a CI tool, these
would be configured with the scripts’ location and order the libraries and
packages needed to execute them.

感谢 GitOps，我们可以将此部署连接到另一个基于 Git 的
事件——创建一个
[发布](https://docs.github.com/en/free-pro-team@latest/github/administering-a-repository/about-releases)
或标记。发布经理所要做的就是创建一个新的“发布”，通常
使用 semver 进行命名，以及构建和部署代码更改的任务
会自动启动。像 CI 工具执行的大多数任务一样，这些
将配置脚本的位置并订购库和
执行它们所需的包。

## GitOps tooling [gitops tooling permalink](http://www.architect.io\#gitops-tooling)

## GitOps 工具 [gitops 工具永久链接](http://www.architect.io\#gitops-tooling)

A solid and intuitive continuous integration tool isn't the only thing needed to
instrument GitOps processes like those described in this article. The CI system
can activate scripts based on git events, but you still need strong tools to
power those scripts and ensure they can be run and maintained easily and safely.
Deploying code changes (aka continuous delivery (CD)) is one of the most
challenging steps to automate, so we've curated a few tooling categories that
can help you through your GitOps journey:

一个可靠而直观的持续集成工具并不是唯一需要的东西
像本文中描述的那样检测 GitOps 流程。 CI系统
可以基于 git 事件激活脚本，但你仍然需要强大的工具来
为这些脚本提供动力，并确保它们可以轻松安全地运行和维护。
部署代码更改（又名持续交付 (CD)）是最
自动化的具有挑战性的步骤，因此我们策划了一些工具类别
可以帮助您完成 GitOps 之旅：

### Containerization with Docker [containerization with docker permalink](http://www.architect.io\#containerization-with-docker)

### Docker 容器化 [Docker 容器化永久链接](http://www.architect.io\#containerization-with-docker)

Docker launched cloud development into an entirely new, distributed landscape
and helped developers begin to realistically consider microservice architectures
as a viable option. Part of what made Docker so powerful was how
developer-friendly it is compared to the previous generation of virtualization
solutions. Just like the declarative CI configurations that live inside our
repositories, developers simply have to write and maintain a `Dockerfile` in
their repository to enable automated container builds of deployable VMs.
Containerization is an enormously powerful tactic for cloud-native teams and
should be a staple tool in your repertoire.

Docker 将云开发带入了一个全新的分布式环境
并帮助开发人员开始现实地考虑微服务架构
作为一个可行的选择。 Docker 如此强大的部分原因在于
与上一代虚拟化相比，对开发人员友好
解决方案。就像我们内部的声明性 CI 配置一样
存储库，开发人员只需编写和维护一个`Dockerfile`
他们的存储库以启用可部署虚拟机的自动化容器构建。
容器化对于云原生团队来说是一种非常强大的策略，
应该是您曲目中的主要工具。

### Infrastructure-as-code (IaC) [infrastructure as code iac permalink](http://www.architect.io\#infrastructure-as-code-iac)

### 基础设施即代码 (IaC) [基础设施即代码 iac 永久链接](http://www.architect.io\#infrastructure-as-code-iac)

A lot goes into provisioning infrastructure and deploying applications that
isn't captured by a `Dockerfile`. For everything else, there's
infrastructure-as-code (IaC) solutions like
[Terraform](https://www.terraform.io/),
[Cloudformation](https://aws.amazon.com/cloudformation/), and others. These
solutions allow developers to describe the other bits of an application, like
Kubernetes resources, load balancers, networking, security, and more, in a
declarative way. Just like the CI configs and Dockerfiles described earlier, IaC
templates can be version controlled and collaborated on by all the developers on
your team. 

配置基础设施和部署应用程序有很多工作要做
不会被`Dockerfile` 捕获。对于其他一切，还有
基础设施即代码 (IaC) 解决方案，例如
[地形](https://www.terraform.io/),
[Cloudformation](https://aws.amazon.com/cloudformation/) 等。这些
解决方案允许开发人员描述应用程序的其他部分，例如
Kubernetes 资源、负载平衡器、网络、安全等，在一个
声明方式。就像前面描述的 CI 配置和 Dockerfiles 一样，IaC
模板可以由所有开发人员进行版本控制和协作
你的团队。

### DevOps automation tools like Architect [devops automation tools like architect permalink](http://www.architect.io\#devops-automation-tools-like-architect)

### DevOps 自动化工具，如 Architect [devops 自动化工具，如架构师永久链接](http://www.architect.io\#devops-automation-tools-like-architect)

I really can't talk about DevOps automation without talking about Architect. We
love IaC and use it heavily as part of our product. We found that configuring
deployments, networking, and network security, especially for microservice
architectures, can be demanding on the developers who should be focused on new
product features instead of infrastructure.

我真的不能不谈论架构师就谈论 DevOps 自动化。我们
喜欢 IaC 并将其大量用作我们产品的一部分。我们发现配置
部署、网络和网络安全，尤其是微服务
架构，可能对应该专注于新的开发人员的要求
产品功能而不是基础设施。

Instead of writing IaC templates and CI pipelines, which require developers to
learn about Kubernetes, Cilium, API gateways, managed databases, or other
infrastructure solutions, just have them write an `architect.yml` file. We'll
automatically deploy dependent APIs/databases and securely broker connectivity
to them every time someone runs `architect deploy`. Our process can
automatically spin up private developer environments, automated preview
environments, and even production-grade cloud environments with just a single
command.

而不是编写 IaC 模板和 CI 管道，这需要开发人员
了解 Kubernetes、Cilium、API 网关、托管数据库或其他
基础架构解决方案，只需让他们编写一个 `architect.yml` 文件。好
自动部署相关 API/数据库并安全地代理连接
每次有人运行 `architect deploy` 时都会向他们发送。我们的流程可以
自动启动私有开发者环境，自动预览
环境，甚至生产级云环境，只需一个
命令。

## Learn more about DevOps, GitOps, and Architect! [learn more about devops gitops and architect permalink](http://www.architect.io\#learn-more-about-devops-gitops-and-architect)

## 了解有关 DevOps、GitOps 和架构师的更多信息！ [了解更多关于 devops gitops 和架构师永久链接](http://www.architect.io\#learn-more-about-devops-gitops-and-architect)

At Architect, our mission is to help ops and engineering teams simply and
efficiently collaborate and achieve deployment, networking, and security
automation all at once. Ready to learn more? Check out these resources:

在 Architect，我们的使命是简单而有效地帮助运维和工程团队
高效协作并实现部署、网络和安全
一下子实现自动化。准备好了解更多了吗？查看这些资源：

- [Creating Microservices: Nest.js](http://www.architect.io/blog/creating-microservices-nestjs)
- [The Importance of Portability in Technology](http://www.architect.io/blog/the-importance-of-portability)
- [Our Product Docs!](http://www.architect.io/docs)

- [创建微服务：Nest.js](http://www.architect.io/blog/creating-microservices-nestjs)
- [技术中可移植性的重要性](http://www.architect.io/blog/the-importance-of-portability)
- [我们的产品文档！](http://www.architect.io/docs)

Or [sign up](https://cloud.architect.io/signup) and try Architect yourself
today! 

或者[注册](https://cloud.architect.io/signup) 自己尝试架构师
今天！

