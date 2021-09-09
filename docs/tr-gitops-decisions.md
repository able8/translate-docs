# GitOps Decisions

# GitOps 决策

November 30, 2020

2020 年 11 月 30 日

[GitOps](https://www.gitops.tech/#:~:text=GitOps%20is%20a%20way%20of,Git%20and%20Continuous%20Deployment%20tools.) is the latest hotness in the software delivery space , following (and extending) on older trends such as DevOps, infrastructure as code, and CI/CD.

[GitOps](https://www.gitops.tech/#:~:text=GitOps%20is%20a%20way%20of,Git%20and%20Continuous%20Deployment%20tools.)是软件交付领域的最新热点，遵循（和扩展)旧趋势，例如 DevOps、基础设施即代码和 CI/CD。

So you’ve [read up on GitOps](https://info.container-solutions.com/what-is-gitops-ebook), you’re bought in to it, and you decide to roll it out.

所以你已经[阅读了 GitOps](https://info.container-solutions.com/what-is-gitops-ebook)，你已经接受了它，你决定推出它。

This is where the fun starts. While the benefits of GitOps are very easy to identify:

这就是乐趣的开始。虽然 GitOps 的好处很容易识别：

- Fully audited changes for free
- Continuous integration and delivery
- Better control over change management
- The possibility of replacing the joys of ServiceNow with pull requests

- 完全经过审核的免费更改
- 持续集成和交付
- 更好地控制变更管理
- 用拉取请求取代 ServiceNow 的乐趣的可能性

the reality is that constructing your GitOps pipelines is far from trivial, and involves many big and small decisions that add up to a _lot_ of work to implement as you potentially chop and change as you go. We at [Container Solutions](https://www.container-solutions.com/) call this **‘GitOps Architecture’** and it can result in real challenges in implementation.

现实情况是，构建您的 GitOps 管道绝非易事，并且涉及许多大大小小的决策，这些决策加起来需要执行 _lot_ 工作，因为您可能会随时进行砍伐和更改。我们在 [Container Solutions](https://www.container-solutions.com/) 称之为**“GitOps 架构”**，它可能会在实施过程中带来真正的挑战。

> GitOps in practice
>
> 2/2 [pic.twitter.com/6vCyPYppFq](https://t.co/6vCyPYppFq)
>
> — Ian Miell (@ianmiell) [September 17, 2020](https://twitter.com/ianmiell/status/1306539329557344258?ref_src=twsrc%5Etfw)

> 实践中的 GitOps
>
> 2/2 [pic.twitter.com/6vCyPYppFq](https://t.co/6vCyPYppFq)
>
> — Ian Miell (@ianmiell) [2020 年 9 月 17 日](https://twitter.com/ianmiell/status/1306539329557344258?ref_src=twsrc%5Etfw)

The good news is that with a bit of planning and experience you can significantly reduce the pain involved in the transition to a GitOps delivery paradigm.

好消息是，通过一些规划和经验，您可以显着减少过渡到 GitOps 交付范式所涉及的痛苦。

In this article, I want to illustrate some of these challenges by telling the story of a company that adopts GitOps as a small scrappy startup, and grows to a regulated multinational enterprise. While such accelerated growth is rare, it does reflect the experience of many teams in larger organisations as they move from proof of concept, to minimum viable product, to mature system.

在本文中，我想通过讲述一家采用 GitOps 的公司作为一个小而斗志昂扬的初创公司，并成长为一家受监管的跨国企业的故事来说明其中的一些挑战。虽然这种加速增长很少见，但它确实反映了大型组织中许多团队在从概念验证到最小可行产品再到成熟系统的过程中的经验。

## ‘Naive’ Startup

##“天真”启动

If you’re just starting out, the simplest thing to do is create a single Git repository with all your needed code in it. This might include:

如果您刚刚开始，最简单的做法是创建一个包含所有所需代码的 Git 存储库。这可能包括：

- Application code
- A[Dockerfile](https://docs.docker.com/engine/reference/builder/), to build the application image
- Some CI/CD pipeline code (eg[GitLab CI/CD](https://docs.gitlab.com/ee/ci/), or [GitHub Actions](https://docs.github.com/en/free-pro-team@latest/actions))
- [Terraform](https://www.terraform.io/) code to provision resources needed to run the application
- All changes directly made to master, changes go straight to live

- 应用程序代码
- A[Dockerfile](https://docs.docker.com/engine/reference/builder/)，构建应用镜像
- 一些 CI/CD 管道代码（例如 [GitLab CI/CD](https://docs.gitlab.com/ee/ci/)，或 [GitHub Actions](https://docs.github.com/en/free-pro-team@latest/actions))
- [Terraform](https://www.terraform.io/) 提供运行应用程序所需资源的代码
- 直接对大师进行的所有更改，更改直接生效

The main benefits of this approach are that you have a single point of reference, and tight integration of all your code. If all your developers are fully trusted, and shipping speed is everything then this might work for a while.

这种方法的主要好处是您有一个单一的参考点，并且所有代码都紧密集成。如果您的所有开发人员都受到完全信任，并且交付速度就是一切，那么这可能会奏效一段时间。

Unfortunately, pretty quickly the downsides of this approach start to show as your business starts to grow.

不幸的是，随着您的业务开始增长，这种方法的缺点很快就会显现出来。

First, the **ballooning size of the repository** as more and more code gets added can result in confusion among engineers as they come across more clashes between their changes. If the team grows significantly, then a lot of rebasing and merging can result in confusion and frustration.

首先，随着越来越多的代码被添加，**存储库的膨胀大小**会导致工程师之间的混淆，因为他们在他们的更改之间遇到更多冲突。如果团队显着增长，那么大量的重新定位和合并可能会导致混乱和沮丧。

Second, you can run into difficulties if you need to **separate control or cadence of pipeline runs**. Sometimes you just want to quickly test a change to the code, not deploy to live, or do a complete build and run of the end-to-end delivery.

其次，如果您需要**管道运行的单独控制或节奏**，您可能会遇到困难。有时您只想快速测试对代码的更改，而不是部署到现场，或者执行端到端交付的完整构建和运行。

Increasingly the monolithic aspect of this approach creates more and more problems that need to be worked on, potentially impacting others’ work as these changes are worked through.

这种方法的单一方面越来越多地产生了越来越多需要解决的问题，随着这些变化的完成，可能会影响其他人的工作。

Third, as you grow you may want more fine-grained responsibility boundaries between engineers and/or teams. While this can be achieved with a single repo (newer features like [CODEOWNERS](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/about-code-owners) files can make this pretty sophisticated), a repository is often a clearer and cleaner boundary.

第三，随着您的成长，您可能希望工程师和/或团队之间有更细粒度的责任界限。虽然这可以通过单个存储库实现（新功能，如 [CODEOWNERS](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/about-code-owners) 文件可以使这变得非常复杂)，存储库通常是一个更清晰、更清晰的边界。

- [![](https://zwischenzugs.files.wordpress.com/2020/11/gitops_separate_config_source-1.png?w=591)](https://zwischenzugs.files.wordpress.com/2020/11/gitops_separate_config_source-1.png?w=591)

gitops_separate_config_source-1.png?w=591)

## Repository Separation

## 存储库分离

It’s getting heavy. Pipelines are crowded and merges are becoming painful. Your teams are separating and specialising in terms of their responsibility. 

越来越重了。管道很拥挤，合并变得痛苦。您的团队在各自的职责方面正在分离和专业化。

**So you decide to separate repositories out.** This is where you’re first faced with a mountain of decisions to make. What is the right level of separation for repositories? Do you have one repository for application code? Seems sensible, right? And include the Docker build stuff in there with it? Well, there’s not much point separating that.

**因此，您决定将存储库分开。** 这是您首先面临大量决策的地方。存储库的正确分离级别是多少？你有一个应用程序代码存储库吗？看起来很明智，对吧？并在其中包含 Docker 构建内容？好吧，将其分开并没有多大意义。

What about all the team Terraform code? Should that be in one new repository? That sounds sensible. But, oh: the newly-created **central 'platform' team wants to control access to the core IAM rule definitions** in AWS, and the teams' RDS provisioning code is in there as well, which the development team want to regularly tweak.

所有的团队 Terraform 代码呢？应该在一个新的存储库中吗？这听起来很明智。但是，哦：新创建的**中央“平台”团队想要控制对 AWS 中核心 IAM 规则定义**的访问，并且团队的 RDS 配置代码也在那里，开发团队希望定期检查调整。

So **you decide to separate out the Terraform out into two repos: a ‘platform’ one and an ‘application-specific’ one.** This creates another challenge, as you now need to separate out the Terraform state files. Not an insurmountable problem, but this isn’t the fast feature delivery you’re used to, so your product manager is now going to have to explain why feature requests are taking longer than previously because of these shenanigans. Maybe you should have thought about this more in advance…

因此 **您决定将 Terraform 分成两个存储库：一个“平台”存储库和一个“特定于应用程序”的存储库。**这又带来了另一个挑战，因为您现在需要分离 Terraform 状态文件。这不是一个无法解决的问题，但这不是您习惯的快速功能交付，因此您的产品经理现在必须解释为什么由于这些诡计，功能请求比以前花费的时间更长。也许你应该提前考虑更多……

**Unfortunately there’s no established best practice or patterns for these GitOps decisions yet.** Even if there were, people love to argue about them anyway, so getting consensus may still be difficult.

**不幸的是，这些 GitOps 决策还没有既定的最佳实践或模式。** 即使有，人们仍然喜欢争论它们，因此达成共识可能仍然很困难。

The problems of separation don’t end there. Whereas before, co-ordination between components of the build within the pipeline were trivial, as everything was co-located, **now you have to orchestrate information flow between repositories**. For example, when a new Docker image is built, this may need to trigger a deployment in a centralised platform repository along with passing over the new image name as part of that trigger.

分离的问题还不止于此。以前，管道内构建组件之间的协调是微不足道的，因为一切都位于同一位置，**现在您必须编排存储库之间的信息流**。例如，当构建一个新的 Docker 镜像时，这可能需要在中央平台存储库中触发部署，同时将新镜像名称作为该触发器的一部分传递。

Again, these are not insurmountable engineering challenges, but they’re easier to implement earlier on in the construction of your GitOps pipeline when you have space to experiment than later on when you don’t.

同样，这些并不是无法克服的工程挑战，但是当您有空间进行实验时，它们在构建 GitOps 管道的早期比没有空间时更容易实现。

OK, your business is growing, and you’re building more and more applications and services. **It increasingly becomes clear that you need some kind of consistency in structure in terms of how applications are built and deployed.** The central platform team tries to start enforcing these standards. Now you get pushback from the development teams who say they were promised more autonomy and control than they had in the ‘bad old days’ of centralised IT before DevOps and GitOps.

好的，您的业务正在增长，并且您正在构建越来越多的应用程序和服务。 **越来越清楚的是，在应用程序的构建和部署方式方面，您需要在结构上保持某种一致性。** 中央平台团队试图开始执行这些标准。现在你会收到开发团队的反对，他们说他们承诺比 DevOps 和 GitOps 之前集中 IT 的“糟糕的过去”拥有更多的自主权和控制权。

If these kind of challenges ring bells in readers’ heads it may be because there is an **analogy here between GitOps and monolith vs microservices arguments** in the application architecture space. Just as you see in those arguments, the tension between distributed and centralised responsibility rears its head more and more as the system matures and grows in size and scope.

如果这些挑战在读者的脑海中响起，那可能是因为在应用程序架构空间中，**GitOps 与单体应用与微服务之间的争论** 之间有一个类比**。正如你在这些论点中看到的那样，随着系统的成熟和规模和范围的扩大，分布式和集中责任之间的紧张关系越来越明显。

On one level, your GitOps flow is just like any other distributed system where poking one part of it may have effects not clearly understood, if you don’t design it well.

在一个层面上，你的 GitOps 流程就像任何其他分布式系统一样，如果你设计得不好，戳它的一部分可能会产生不清楚的影响。

> I'll just make a small change to one of the repos in my GitOps setup. It'll be fine. [pic.twitter.com/dhIRGYN5NX](https://t.co/dhIRGYN5NX)
>
> — Ian Miell (@ianmiell) [November 26, 2020](https://twitter.com/ianmiell/status/1331976617489543168?ref_src=twsrc%5Etfw)

> 我只会对我的 GitOps 设置中的一个存储库做一个小的更改。会好的。 [pic.twitter.com/dhIRGYN5NX](https://t.co/dhIRGYN5NX)
>
> — Ian Miell (@ianmiell) [2020 年 11 月 26 日](https://twitter.com/ianmiell/status/1331976617489543168?ref_src=twsrc%5Etfw)

* * *

* * *

_**If you like this, you might like my book**_ _**[Learn Git the Hard Way](https://leanpub.com/learngitthehardway?p=5148)**_

_**如果你喜欢这个，你可能会喜欢我的书**_ _**[Learn Git the Hard Way](https://leanpub.com/learngitthehardway?p=5148)**_

* * *

* * *

## Environments

## 环境

At about the same time as you decide to separate repositories, you realise that **you need a consistent way to manage different deployment environments**. Going straight to live no longer cuts it, as a series of outages has helped birth a QA team who want to test changes before they go out. 

大约在您决定分离存储库的同时，您意识到 **您需要一种一致的方式来管理不同的部署环境**。直接投入使用不再会削减它，因为一系列中断已经帮助诞生了一个想要在更改退出之前测试更改的 QA 团队。

Now you need to specify a different Docker tag for your application in ‘test’ and ‘QA’ environments. You might also want different instance sizes or replication features enabled in different environments. How do you manage the configuration of these different environments in source? **A naive way to do this might be to have a separate Git repository per environment** (eg superapp-dev, super-app-qa, super-app-live).

现在，您需要在“测试”和“QA”环境中为您的应用程序指定不同的 Docker 标签。您可能还希望在不同环境中启用不同的实例大小或复制功能。你如何在源代码中管理这些不同环境的配置？ **一种简单的方法可能是为每个环境拥有一个单独的 Git 存储库**（例如 superapp-dev、super-app-qa、super-app-live）。

Separating repositories has the ‘clear separation’ benefit that we saw with dividing up the Terraform code above. However, **few end up liking this solution**, as it can require a level of Git knowledge and discipline most teams don’t have in order to port changes between repositories with potentially differing histories. There will necessarily be a lot of duplicated code between the repositories, and – over time – potentially a lot of drift too.

分离存储库具有我们在分离上面的 Terraform 代码时看到的“清晰分离”的好处。然而，**很少有人最终喜欢这个解决方案**，因为它可能需要一定程度的 Git 知识和纪律，大多数团队都没有，以便在具有潜在不同历史的存储库之间移植更改。存储库之间必然会有很多重复的代码，而且随着时间的推移，也可能会有很多漂移。

If you want to keep things to a single repo you have (at least) three options:

如果您想将内容保留在单个存储库中，您（至少）有三个选择：

- A directory per environment
- A branch per environment
- A tag per environment

- 每个环境的目录
- 每个环境一个分支
- 每个环境的标签

- [![](https://zwischenzugs.files.wordpress.com/2020/11/gitops_configrepo_apps-2.png?w=927)](https://zwischenzugs.files.wordpress.com/2020/11/gitops_configrepo_apps-2.png?w=927)

gitops_configrepo_apps-2.png?w=927)

### Sync Step Choices

### 同步步骤选择

If you rely heavily on a YAML generator or templating tool, then you will likely be nudged more towards one or other choice. Kustomize, for example, strongly encourages a directory-based separation of environments. If you’re using raw yaml, then a branch or tagging approach might make you more comfortable. If you have experience with your CI tool in using one or other approach previously in your operations, then you are more likely to prefer that approach. **Whichever choice you make, prepare yourself for much angst and discussion about whether you’ve chosen the right path.**

如果您严重依赖 YAML 生成器或模板工具，那么您可能会更倾向于一种或其他选择。例如，Kustomize 强烈鼓励基于目录的环境分离。如果您使用原始 yaml，那么分支或标记方法可能会让您更舒服。如果您之前在运营中使用过一种或其他方法，那么您更有可能更喜欢这种方法。 **无论您做出何种选择，都要准备好迎接关于您是否选择了正确道路的焦虑和讨论。**

- [![](https://zwischenzugs.files.wordpress.com/2020/11/gitops_sync_config_strategy-1.png?w=742)](https://zwischenzugs.files.wordpress.com/2020/11/gitops_sync_config_strategy-1.png?w=742)

gitops_sync_config_strategy-1.png?w=742)

### Runtime Environment Granularity

### 运行时环境粒度

Also on the subject of runtime environments, there are **choices to be made on what level of separation you want**. On the cluster level, if you’re using Kubernetes, you can choose between:

同样在运行时环境的主题上，**可以选择您想要的分离级别**。在集群级别，如果您使用 Kubernetes，您可以选择：

- One cluster to rule them all
- A cluster per environment
- A cluster per team

- 一个集群来统治他们
- 每个环境一个集群
- 每个团队一个集群

At one extreme, you can put all your environments into one cluster. Usually, there is at least a separate cluster for production in most organisations.

在一种极端情况下，您可以将所有环境放入一个集群中。通常，大多数组织中至少有一个单独的生产集群。

Once you’ve figured out your cluster policy, at the namespace level, you can still choose between:

一旦确定了集群策略，在命名空间级别，您仍然可以选择：

- A namespace per environment
- A namespace per application/service
- A namespace per engineer
- A namespace per build

- 每个环境的命名空间
- 每个应用程序/服务的命名空间
- 每个工程师的命名空间
- 每个构建的命名空间

**Platform teams often start with a ‘dev’, ‘test’, ‘prod’ namespace setup, before realising they want more granular separation of teams’ work.**

**平台团队通常从“开发”、“测试”、“生产”命名空间设置开始，然后才意识到他们想要更精细地分离团队的工作。**

You can also mix and match these options, for example offering each engineer their own namespace for ‘desk testing’, as well as a namespace per team if you want.

您还可以混合和匹配这些选项，例如，为每个工程师提供自己的“桌面测试”命名空间，如果需要，还可以为每个团队提供一个命名空间。

## Conclusion

##  结论

**We've only scratched the surface here of the areas of decision-making required to get a mature GitOps flow going.** You might also consider RBAC/IAM and onboarding, for example, an absolute requirement if you grow to become that multinational enterprise. 

**我们在这里只是触及了使成熟的 GitOps 流程运行所需的决策领域的皮毛。** 例如，如果您成长为这样，您还可以考虑 RBAC/IAM 和入职，这是绝对要求跨国企业。

Often **rolling out GitOps can feel like a lot of front-loaded work and investment**, until you realise that before you did this none of it was encoded at all. Before GitOps, chaos and delays ensued as no-one could be sure in what state anything was, or should be. These resulted in secondary costs as auditors did spot checks and outages caused by unexpected and unrecorded changes occupied your most expensive employees’ attention. As you mature your GitOps flow, the benefits multiply, and your process takes care of many of these challenges. **But more often than not, you are under pressure to demonstrate success more quickly than you can build a stable framework.**

通常**推出 GitOps 感觉像是大量的前期工作和投资**，直到你意识到在你这样做之前它根本没有被编码。在 GitOps 之前，混乱和延迟接踵而至，因为没有人可以确定任何东西处于或应该处于什么状态。这导致了二次成本，因为审计员进行了抽查，并且由于意外和未记录的更改而引起的中断占据了您最昂贵的员工的注意力。随着您的 GitOps 流程成熟，收益会成倍增加，并且您的流程会解决许多这些挑战。 **但通常情况下，您面临着比构建稳定框架更快地证明成功的压力。**

**The biggest challenge with GitOps right now is that there are no established patterns to guide you in your choices.** As consultants, we're often acting as sherpas, guiding teams towards finding the best solutions for them and nudging them in certain directions based on our experience.

**目前 GitOps 的最大挑战是没有既定的模式来指导您做出选择。** 作为顾问，我们经常充当夏尔巴人，指导团队为他们找到最佳解决方案，并在某些情况下推动他们根据我们的经验指导。

What I’ve observed, though, is that **choices avoided early on because they seem ‘too complicated’ are often regretted later**. But **I don’t want to say that that means you should jump straight to a namespace per build**, and a Kubernetes cluster per team, for two reasons.

然而，我观察到的是 ** 早期避免的选择，因为它们看起来“太复杂”，后来往往会后悔**。但是**我不想说这意味着你应该直接跳转到每个构建的命名空间**和每个团队的 Kubernetes 集群，原因有两个。

**1) Every time you add complexity to your GitOps architecture, you will end up adding to the cost and time to deliver a working GitOps solution**.

**1) 每次为 GitOps 架构增加复杂性时，最终都会增加交付可用 GitOps 解决方案的成本和时间**。

**2) You might genuinely never need that setup anyway.**

**2) 无论如何，您可能真的永远不需要那个设置。**

Until we have genuine standards in this space, getting your GitOps architecture right will always be an art rather than a science. 

在我们在这个领域拥有真正的标准之前，让您的 GitOps 架构正确将永远是一门艺术，而不是一门科学。

