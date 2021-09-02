# The Next Step after DevOps and GitOps Is Cloud Engineering, Pulumi Says

# DevOps 和 GitOps 之后的下一步是云工程，Pulumi 说

#### 3 May 2021 7:55am,   by [Mary Branscombe](https://thenewstack.io/author/marybranscombe/ "Posts by Mary Branscombe")



If we going to treat infrastructure as code, shouldn’t infrastructure engineers have access to the same tools that make software engineers productive and even the same languages? That’s the theory behind [Pulumi](https://thenewstack.io/pulumi-uses-real-programming-languages-to-enforce-cloud-best-practices/), which has just released version 3 of its open source platform.

如果我们将基础设施视为代码，那么基础设施工程师难道不应该使用相同的工具来提高软件工程师的生产力，甚至是相同的语言吗？这就是 [Pulumi](https://thenewstack.io/pulumi-uses-real-programming-languages-to-enforce-cloud-best-practices/) 背后的理论，它刚刚发布了其开源平台的第 3 版。

What founder and CEO [Joe Duffy](https://www.linkedin.com/in/joejduffy/), calls a “cloud engineering platform” is an attempt to “distill a lot of the lessons learned from helping developers build modern cloud applications, helping infrastructure teams increasingly apply engineering disciplines to the way they're doing infrastructure and help the entire team really ship faster with confidence.”

创始人兼 CEO [Joe Duffy](https://www.linkedin.com/in/joejduffy/) 所说的“云工程平台”是试图“从帮助开发人员构建现代云中汲取很多经验教训”应用程序，帮助基础设施团队越来越多地将工程学科应用到他们做基础设施的方式中，并帮助整个团队真正自信地更快地交付。”

“People are realizing the only way to keep up with the pace of the modern cloud and the level of innovation is to empower developers to be more self serve but also learn from the past decades in software engineering and apply that to the way we're doing infrastructure.”

“人们意识到跟上现代云和创新水平的唯一方法是让开发人员能够更加自助，同时也从过去几十年的软件工程中学习并将其应用于我们的工作方式做基础设施。”

“Not everybody can spend two years going on the journey of figuring out how to do cloud engineering: we want to make sure that everybody has access to this on day one.”

“不是每个人都可以花两年时间去弄清楚如何进行云工程：我们希望确保每个人在第一天都可以访问它。”

## Cloud Engineering Culture

## 云工程文化

“Cloud engineering uses standard software engineering and tools across infrastructure, app dev and compliance to simplify and homogenize the complexity of modern cloud environments,” IDC DevOps research director [Jim Mercer](https://www.idc.com/getdoc.jsp?containerId=PRF005085) explained to the New Stack.

“云工程使用跨基础架构、应用程序开发和合规性的标准软件工程和工具来简化和统一现代云环境的复杂性，”IDC DevOps 研究总监 [Jim Mercer](https://www.idc.com/getdoc.jsp?containerId=PRF005085) 向新堆栈解释。

To Justin Fitzhugh, whose job title at Pulumi customer Snowflake is vice president of cloud engineering, the term represents the continued evolution from the separate cultures of development and operations.

对于 Justin Fitzhugh 来说，他在 Pulumi 客户 Snowflake 的职位是云工程副总裁，这个词代表了开发和运营的独立文化的持续演变。

“With DevOps, we saw developers and operations teams working together on a product or a specific deliverable. We saw applications and systems being developed jointly where operational concerns were being taken into account in the development phase and inversely the developers were concerned about how they scale and how do you manage as you go into production.”

“通过 DevOps，我们看到开发人员和运营团队在产品或特定可交付成果上协同工作。我们看到应用程序和系统是联合开发的，在开发阶段就考虑到了运营问题，相反，开发人员关心的是它们如何扩展以及在投入生产时如何管理。”

Cloud engineering takes that to the next level. “It’s more of a proactive engineering culture where we’re building tooling, software and components to build, manage, instantiate and take care of the infrastructure. Any change to our infrastructure is a commit to a codebase which runs through a CI/CD pipeline where you can diff and you can test and has appropriate reviews just as any other software change would and then is pushed to production through automated pipelines. We look at it as a software engineering function, just focused on infrastructure as opposed to UI or back end engineering.”

云工程将其提升到一个新的水平。 “这更像是一种积极主动的工程文化，我们正在构建工具、软件和组件来构建、管理、实例化和维护基础设施。对我们基础架构的任何更改都是对代码库的提交，该代码库通过 CI/CD 管道运行，您可以在其中进行比较，您可以测试并进行适当的审查，就像任何其他软件更改一样，然后通过自动化管道推送到生产中。我们将其视为软件工程职能，只关注基础设施，而不是 UI 或后端工程。”

Infrastructure still requires a specific discipline, Fitzhugh noted, because their concerns are integration, scale, abstraction and how to interact with and operate cloud services, but much of that can be automated.

Fitzhugh 指出，基础设施仍然需要特定的学科，因为他们关注的是集成、规模、抽象以及如何与云服务交互和操作云服务，但其中大部分都可以自动化。

![](https://cdn.thenewstack.io/media/2021/04/40f21f70-pulumi-01-1024x791.png)

Some DevOps teams keep their infrastructure configuration in git repos, but too often operational changes are made by running through the steps on a to-do list, even if that’s kept in a wiki or a configuration management system. “As opposed to how do I codify that, how do I commit that to a codebase, how do I make that part of a reproducible test run and have it be part of a pipeline.”

一些 DevOps 团队将他们的基础设施配置保存在 git repos 中，但通常通过运行待办事项列表中的步骤来进行操作更改，即使它保存在 wiki 或配置管理系统中。 “与我如何编码相反，我如何将其提交到代码库中，我如何使该部分成为可重复的测试运行并使其成为管道的一部分。”

Snowflake uses Go extensively, along with everything from Ansible to Python, and it runs on AKS, EKS and GKE so needs to deal with different APIs, account metrics and cloud primitives. “We use Pulumi to wrap that into a single workflow orchestration.” 

Snowflake 广泛使用 Go，以及从 Ansible 到 Python 的所有内容，它运行在 AKS、EKS 和 GKE 上，因此需要处理不同的 API、帐户指标和云原语。 “我们使用 Pulumi 将其包装到单个工作流编排中。”

Duffy called Pulumi “connective tissue” and noted that, increasingly, security engineering teams are part of that same workflow because Pulumi users are trying to get away from silos where they have to deal with YAML and multiple domain-specific languages, with different delivery mechanisms for infrastructure — many of them manual — even if they're doing automated delivery of containers for some of their workloads.

Duffy 将 Pulumi 称为“结缔组织”，并指出，越来越多的安全工程团队成为同一工作流程的一部分，因为 Pulumi 用户正试图摆脱必须处理 YAML 和多种领域特定语言的孤岛，并使用不同的交付机制对于基础设施——其中许多是手动的——即使他们正在为他们的一些工作负载自动交付容器。

And increasingly, cloud native requires orchestration across multiple services; doing that with bash scripts can be painful, especially when you think about security.

并且越来越多的云原生需要跨多个服务进行编排；使用 bash 脚本执行此操作可能会很痛苦，尤其是当您考虑安全性时。

“We find the line starts to become blurry with these modern architectures. Is a serverless function infrastructure or is that an application? Is a queue or a pub sub-topic part of the application or is that infrastructure? For a lot of folks leaning into the modern cloud, it’s a little bit of both. That’s why it’s important to have one way to build, deploy and manage both the applications and the infrastructure and having one substrate to manage all those things, especially when they have interdependencies is really powerful.”

“我们发现这些现代建筑的界限开始变得模糊。是无服务器功能基础设施还是应用程序？是应用程序的队列还是pub 子主题的一部分，还是该基础设施？对于许多倾向于现代云的人来说，两者兼而有之。这就是为什么使用一种方法来构建、部署和管理应用程序和基础设施，并拥有一个底层来管理所有这些东西是很重要的，尤其是当它们具有相互依赖性时非常强大。”

![](https://cdn.thenewstack.io/media/2021/04/e007cab2-pulumi-02-1024x791.png)

## Developer Terms for Infrastructure Features

## 基础设施功能的开发者条款

The new features in Pulumi 3.0 are familiar concepts for developers that are useful for infrastructure teams too, with a cloud twist.

Pulumi 3.0 中的新功能对于开发人员来说是熟悉的概念，对基础架构团队也很有用，并带有云计算功能。

Teams could already share and reuse code in [Pulumi Resources](https://thenewstack.io/infrastructure-is-code-and-with-pulumi-2-0-so-is-architecture-and-policy/), but thanks to the underlying language-neutral model of the cloud objects you're working with, Pulumi Packages now make that work across multiple languages and for higher-level Components rather than low-level Resources.

团队已经可以在 [Pulumi Resources](https://thenewstack.io/infrastructure-is-code-and-with-pulumi-2-0-so-is-architecture-and-policy/) 中共享和重用代码，但是由于您正在使用的云对象的底层语言中立模型，Pulumi 包现在可以跨多种语言和更高级别的组件而不是低级别资源工作。

“Previously, if you wrote your package in Node.js. you were tied to the Node ecosystem or if you wrote it in Python, it was only available in Python. Customers want to be able to write packages in Python if the infrastructure team is using Python and defining a Kubernetes cluster component but they want to enable their developers to spin those things up from Node.js or Go or their favorite language,” Duffy explained.

“以前，如果你用 Node.js 编写你的包。你被绑定到 Node 生态系统，或者如果你用 Python 编写它，它只在 Python 中可用。如果基础设施团队使用 Python 并定义 Kubernetes 集群组件，客户希望能够用 Python 编写包，但他们希望他们的开发人员能够从 Node.js 或 Go 或他们喜欢的语言中提取这些东西，”Duffy 解释说。

Packages can be written in any language Pulumi supports and consumed from any other supported language. Pulumi is also supplying some multilanguage components; the first handles provisioning and managing AWS EKS clusters, which is frequently complex. “We’re bootstrapping an ecosystem here,” Duffy suggested and noted that a numbers of customers will publish their own once the feature is generally available.

包可以用 Pulumi 支持的任何语言编写，也可以从任何其他支持的语言中使用。 Pulumi 还提供了一些多语言组件；第一个处理配置和管理 AWS EKS 集群，这通常很复杂。 “我们正在这里建立一个生态系统，”Duffy 建议并指出，一旦该功能普遍可用，许多客户将发布自己的生态系统。

Native providers are a new type of Pulumi Package that automatically generates interfaces from the Azure, Google, and Amazon API specifications. Since September 2020, Azure has added 166 new services or new features in existing services; “They’re shipping services left and right all the time and if you can’t access them, that’s a problem.”

Native providers 是一种新型的 Pulumi Package，可以根据 Azure、Google 和 Amazon API 规范自动生成接口。自 2020 年 9 月以来，Azure 已在现有服务中添加了 166 项新服务或新功能； “他们一直在左右运送服务，如果你无法访问它们，那就是一个问题。”

Because Microsoft documents the entire Azure surface in OpenAPI, Pulumi can support the new features immediately. The Azure native provider includes twice as many services as the provider Pulumi had built manually, like Azure Static Websites which was implemented within an hour of release. Azure native support is GA now, with Google Cloud in public preview and AWS coming in the second half of this year.

由于 Microsoft 在 OpenAPI 中记录了整个 Azure Surface，Pulumi 可以立即支持新功能。 Azure 本地提供程序包含的服务数量是 Pulumi 手动构建的提供程序的两倍，例如在发布一小时内实施的 Azure 静态网站。 Azure 原生支持现已正式发布，Google Cloud 处于公共预览阶段，AWS 将于今年下半年推出。

Fitzhugh noted that the fast support of new features differentiates Pulumi. “Cloud providers are moving their APIs ahead, sometimes at a breakneck pace; having support for that in real-time is critical. Especially when we’re looking at Kubernetes-based offerings from the cloud providers, they’re iterating and moving forward so fast that a lot of the frameworks would have trouble keeping up.” 

Fitzhugh 指出，对新功能的快速支持使 Pulumi 与众不同。 “云提供商正在推进他们的 API，有时以极快的速度；对此提供实时支持至关重要。尤其是当我们查看来自云提供商的基于 Kubernetes 的产品时，他们迭代和前进的速度如此之快，以至于许多框架都无法跟上。”

Because they use all three major cloud providers, Snowflake is keen to give its developers a uniform platform for deploying containers and uses Pulumi to get that abstraction. “We're not trying to build a complete other cloud platform on top of it, because where I see people get in trouble is when they try to almost reinvent what the cloud offer, and build it on top of it, because you're never going to iterate as fast as the cloud,” Fizhugh explained.

因为他们使用所有三个主要的云提供商，Snowflake 热衷于为其开发人员提供一个统一的平台来部署容器并使用 Pulumi 来获得这种抽象。 “我们并没有试图在它之上构建一个完整的其他云平台，因为我看到人们遇到麻烦的地方是当他们试图几乎重新发明云提供的东西并在它之上构建它时，因为你是永远不会像云一样快速迭代，”Fizhugh 解释说。

“We’re focusing a lot on what the onboarding and user experience looks like for the end users: how do we streamline that, how do we make it as efficient as possible. But this is all enabled by the fact that we can describe what that environment looks like via code and we can iterate on it quickly at scale across, our many, many deployments across multiple cloud providers.”

“我们非常关注最终用户的入职和用户体验：我们如何简化它，如何使其尽可能高效。但这一切都是因为我们可以通过代码描述该环境的样子，并且我们可以在多个云提供商的许多部署中快速大规模地迭代它。”

The new Automation API is another feature that will help by letting customers build infrastructure-as-code into their own software. “What if infrastructure as code wasn’t a CLI-based experience; if it was just a library you could use within other programs,” Duffy explained.

新的自动化 API 是另一项功能，它可以让客户将基础设施即代码构建到他们自己的软件中。 “如果基础设施即代码不是基于 CLI 的体验怎么办？如果它只是一个可以在其他程序中使用的库，”Duffy 解释说。

This is particularly useful for SaaS providers like Cockroach Labs whose database as a service manages Kubernetes clusters behind the scenes on behalf of their customers but organizations are also using it to build self-service portals like Snowflake’s internal platform. “You can link with Pulumi and manage your infrastructure, and still get all the power of infrastructure as code, but not have to have this clunky CLI-based interface that you drive programmatically. You can build your own self-service portals: we’ve seen people building their own custom tools to do multiregion rollouts and lots of complex scenarios.”

这对于像 Cockroach Labs 这样的 SaaS 提供商特别有用，他们的数据库即服务代表客户在幕后管理 Kubernetes 集群，但组织也使用它来构建自助服务门户，如 Snowflake 的内部平台。 “您可以与 Pulumi 链接并管理您的基础设施，并且仍然以代码形式获得基础设施的所有功能，但不必拥有您以编程方式驱动的基于 CLI 的笨重界面。您可以构建自己的自助服务门户：我们已经看到人们构建自己的自定义工具来进行多区域部署和许多复杂的场景。”

![](https://cdn.thenewstack.io/media/2021/04/15a93a1a-pulumi-03-1024x791.png)

## Next for Pulumi

## 接下来是Pulumi

For Fitzhugh, the Automation API, and the new RBAC support that allows single sign-on with SAML, fine-grained permissions and identity workflows that synchronize with a central identity provider, move them a step closer to cloud engineering. “Touching the infrastructure in any way, either creating, managing or tearing it down is ideally all done through appropriate means; you really need a CI/CD pipeline to test and then deploy in an automated way. Our customers are asking us how can we have fewer humans interacting with the production environment and more pushing it through a pipeline. Identity management and the compliance and regulation pieces with SAML SSO and more granular access control piece is really useful to us.”

对于 Fitzhugh，自动化 API 和新的 RBAC 支持（允许使用 SAML 进行单点登录）、细粒度的权限和与中央身份提供者同步的身份工作流，使它们更接近云工程。 “以任何方式接触基础设施，无论是创建、管理还是拆除，理想情况下都是通过适当的方式完成的；您确实需要一个 CI/CD 管道来进行测试，然后以自动化方式进行部署。我们的客户问我们如何才能减少与生产环境交互的人员并更多地推动它通过管道。带有 SAML SSO 的身份管理以及合规性和监管部分以及更细粒度的访问控制部分对我们非常有用。”

Identity is important for organizations using Pulumi for hybrid cloud; while 80% of Pulumi customers rely heavily on public cloud, hybrid and private cloud are also important. Pulumi supports Azure Arc and AWS Outposts; it’s already being used with the EKS distro and will support EKS Anywhere. Other customers are using Pulumi with vSphere, Duffy noted; “One customer is using Pulumi to spin up bare metal worldwide for new data centers.”

身份对于将 Pulumi 用于混合云的组织很重要；虽然 80% 的 Pulumi 客户严重依赖公共云，但混合和私有云也很重要。 Pulumi 支持 Azure Arc 和 AWS Outposts；它已经与 EKS 发行版一起使用，并将支持 EKS Anywhere。 Duffy 指出，其他客户正在将 Pulumi 与 vSphere 结合使用； “一个客户正在使用 Pulumi 在全球范围内为新数据中心启动裸机。”

The Pulumi service doesn’t ever see your identity or credentials for the cloud provider, leaving you to do credential management the way you choose. “But we’ve almost become counselors on how to properly manage cloud credentials because unfortunately, we see it done incorrectly so often,” Duffy told us. “We try to push some of the best CI/CD integrations, we help folks do temporary credentials so there are no long-lived credentials — we see that all the time as an antipattern, but you look at the CI/CD providers themselves and they tell you to do that, which is really, really bad. We try to help people secure the CI/CD pipeline; when all deployments go through that, you can lock that down and secure it.” 

Pulumi 服务永远不会看到您的身份或云提供商的凭据，让您以自己选择的方式进行凭据管理。 “但我们几乎已经成为如何正确管理云凭证的顾问，因为不幸的是，我们经常看到它做错了，”达菲告诉我们。 “我们试图推动一些最好的 CI/CD 集成，我们帮助人们做临时凭证，这样就没有长期存在的凭证——我们一直认为这是一种反模式，但你看看 CI/CD 提供者本身，他们告诉你这样做，这真的非常糟糕。我们试图帮助人们保护 CI/CD 管道；当所有部署都通过时，您可以锁定并保护它。”

He noted the complexity of dealing with so many systems that have their own identity models and suggested that in the future, Pulumi would be able to help with more of that. “Part of the whole cloud engineering platform vision is that we’re going to bring more into the fold of Pulumi,. If you want me to do delivery for you, great. If you just want to click a button to do some deployments, great. So there’s going to be a lot more of that and that will get us more into the identity credential management business.”

他指出处理这么多拥有自己身份模型的系统的复杂性，并建议在未来，Pulumi 将能够帮助解决更多问题。 “整个云工程平台愿景的一部分是我们将为 Pulumi 带来更多。如果你想让我为你送货，那太好了。如果您只想单击一个按钮进行一些部署，那就太好了。因此，将会有更多这样的事情发生，这将使我们更多地涉足身份凭证管理业务。”

Duffy also talked about managing more of the dynamic state of infrastructure. “Pulumi captures the static topology of your infrastructure and applications which is beneficial and allows you to validate a lot of things, but it doesn’t validate everything I think the next frontier is capturing dynamic dependencies and understanding the semantic connections between them.” That could build on tools like the AWS security group analyzer for tracking down problems like the firewall rule that’s stopping two machines connecting.

Duffy 还谈到了管理更多基础设施的动态状态。 “Pulumi 捕获您的基础设施和应用程序的静态拓扑，这是有益的，并允许您验证很多事情，但它并不能验证所有我认为下一个前沿是捕获动态依赖关系并理解它们之间的语义连接。”这可以建立在诸如 AWS 安全组分析器之类的工具之上，用于跟踪诸如阻止两台机器连接的防火墙规则之类的问题。

Pulumi will also continue to improve and extent language support. Before adding new languages, the focus is on improving existing language support: Python support now includes static type checkers and the Go libraries are smaller and so faster to load. Pulumi 4.0 will include new languages: PowerShell, JVM and Ruby. “You can use Pulumi with PowerShell, it’s just the API’s aren’t designed to be PowerShell friendly, or idiomatically PowerShell.”

Pulumi 还将继续改进和扩大语言支持。在添加新语言之前，重点是改进现有语言支持：Python 支持现在包括静态类型检查器，Go 库更小，加载速度更快。 Pulumi 4.0 将包括新语言：PowerShell、JVM 和 Ruby。 “你可以将 Pulumi 与 PowerShell 一起使用，只是 API 的设计不是为了对 PowerShell 友好，或者是惯用的 PowerShell。”

Customers coming from Chef and [Puppet](https://puppet.com/?utm_content=inline-mention) ecosystem are driving the interest in Ruby, Duffy said. “It’s actually the number one upvoted issue out of all of our open source issues.”

Duffy 说，来自 Chef 和 [Puppet](https://puppet.com/?utm_content=inline-mention) 生态系统的客户正在推动对 Ruby 的兴趣。 “这实际上是我们所有开源问题中投票最多的问题。”

![](https://cdn.thenewstack.io/media/2021/04/46acfd5d-pulumi-04-1024x791.png)

Puppet is a sponsor of The New Stack.


Puppet 是 The New Stack 的赞助商。


Feature image par [stokpic](https://pixabay.com/fr/users/stokpic-692575/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=600468) de [Pixabay](https://pixabay.com/fr/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=600468). 

特征图片标准 [stokpic](https://pixabay.com/fr/users/stokpic-692575/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=600468) de [Pixabay](https://pixabay.com//?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=600468)。

