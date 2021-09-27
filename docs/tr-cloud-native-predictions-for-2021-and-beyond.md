## Cloud Native Predictions for 2021 and Beyond

## 2021 年及以后的云原生预测

###### By [Chris Aniszczyk](https://www.aniszczyk.org/author/admin/ "Posts by Chris Aniszczyk")

###### 作者 [Chris Aniszczyk](https://www.aniszczyk.org/author/admin/“Chris Aniszczyk 的帖子”)

I hope everyone had a wonderful holiday break as the first couple weeks of January 2021 have been pretty wild, from insurrections to new COVID strains. In cloud native land, the CNCF recently released its [annual report](https://www.cncf.io/cncf-annual-report-2020/) on all the work we accomplished last year. I recommend everyone take an opportunity to go through the report, we had a solid year given the wild pandemic circumstances.

我希望每个人都能度过一个愉快的假期，因为 2021 年 1 月的前几周非常疯狂，从叛乱到新的 COVID 毒株。在云原生领域，CNCF近日发布了 [年度报告](https://www.cncf.io/cncf-annual-report-2020/)，介绍了我们去年完成的所有工作。我建议大家抓住机会阅读这份报告，鉴于大流行的情况，我们度过了充实的一年。

https://twitter.com/CloudNativeFdn/status/1343914259177222145

https://twitter.com/CloudNativeFdn/status/1343914259177222145

As part of my job, I have a unique and privileged vantage point of cloud native trends given to all the member companies and developers I work with, so I figured I’d share my thoughts of where things will be going in 2021 and beyond:

作为我工作的一部分，我为所有与我合作的成员公司和开发人员提供了云原生趋势的独特优势，因此我想我会分享我对 2021 年及以后的发展方向的看法：

**Cloud Native IDEs**

**云原生 IDE**

As a person who has spent a decent portion of his career working on developer tools inside the Eclipse Foundation, I am nothing but thrilled with the recent progress of the state of the art. The future will hold that the development lifecycle (code, build, debug) will happen mostly in the cloud versus your local Emacs or VSCode setup. You will end up getting a full dev environment setup for every pull request, pre-configured and connected to their own deployment to aid your development and debugging needs. A concrete example of this technology today is enabled via GitHub [Codespaces](https://github.com/features/codespaces) and [GitPod](https://gitpod.io/). While GitHub Codespaces is still in beta, you can try this experience live today with GitPod, using [Prometheus as an example](https://gitpod.io/#https://github.com/prometheus/prometheus). In a minute or so, you have a completely live development environment with an editor and preview environment. The wild thing is that this development environment (workspace) is [described in code](https://github.com/prometheus/prometheus/blob/master/.gitpod.yml) and shareable with other developers on your team like any other code artifact.

作为一个在 Eclipse 基金会内部从事开发人员工具工作的人，我对最新技术的最新进展感到非常兴奋。未来将认为开发生命周期（代码、构建、调试）将主要发生在云中，而不是本地 Emacs 或 VSCode 设置。您最终将为每个拉取请求获得完整的开发环境设置，预配置并连接到他们自己的部署，以帮助您的开发和调试需求。今天这项技术的一个具体例子是通过 GitHub [Codespaces](https://github.com/features/codespaces) 和 [GitPod](https://gitpod.io/) 启用的。虽然 GitHub Codespaces 仍处于测试阶段，但您今天可以使用 GitPod 尝试这种体验，使用 [Prometheus 为例](https://gitpod.io/#https://github.com/prometheus/prometheus)。大约一分钟后，您就拥有了一个带有编辑器和预览环境的完全实时的开发环境。疯狂的是，这个开发环境（工作区)[在代码中描述](https://github.com/prometheus/prometheus/blob/master/.gitpod.yml) 并且可以像任何其他开发人员一样与团队中的其他开发人员共享代码工件。

In the end, I expect to see incredible innovation in the cloud native IDE space over the next year, especially as GitHub Codespaces enters out of beta and becomes more widely available so developers can experience this new concept and fall in love.

最后，我希望在明年在云原生 IDE 领域看到令人难以置信的创新，特别是随着 GitHub Codespaces 进入测试版并变得更广泛可用，因此开发人员可以体验这个新概念并坠入爱河。

**Kubernetes on the Edge**

**Kubernetes 在边缘**

Kubernetes was born through usage across massive data centers but Kubernetes will evolve just like Linux did for new environments. What happened with Linux was that end users eventually stretched the kernel to support a variety of new deployment scenarios from mobile, embedded and more. I strongly believe Kubernetes will go through a similar evolution and we are already witnessing Telcos (and startups) explore Kubernetes as an edge platform through transforming VNFs into [Cloud Native Network Functions](https://github.com/cncf/cnf-wg) (CNFs) along with open source projects like [k3s](https://k3s.io/), KubeEdge, k0s, [LFEdge](https://www.lfedge.org), Eclipse ioFog and more. The forces driving hyperscaler clouds to support telcos and the edge, combined with the ability to reuse cloud native software and build upon already a large ecosystem will cement Kubernetes as a dominant platform in edge computing over the next few years.

Kubernetes 诞生于跨大规模数据中心的使用，但 Kubernetes 将像 Linux 为新环境所做的那样发展。 Linux 的情况是，最终用户最终扩展了内核，以支持移动、嵌入式等各种新部署场景。我坚信 Kubernetes 将经历类似的演变，我们已经见证了电信公司（和初创公司）通过将 VNF 转换为 [云原生网络功能](https://github.com/cncf/cnf-wg)来探索 Kubernetes 作为边缘平台) (CNF) 以及 [k3s](https://k3s.io/)、KubeEdge、k0s、[LFEdge](https://www.lfedge.org)、Eclipse ioFog 等开源项目。推动超大规模云支持电信公司和边缘的力量，再加上重用云原生软件和建立在一个庞大的生态系统之上的能力，将在未来几年巩固 Kubernetes 作为边缘计算的主导平台。

**Cloud Native + Wasm** 

**云原生 + Wasm**

[Web Assembly](https://webassembly.org/)(Wasm) is a technology that is nascent but I expect it to become a growing utility and workload in the cloud native ecosystem especially as [WASI](https://wasi.dev/) matures and as Kubernetes is used more as an edge orchestrator as described previously. One use case is powering an extension mechanism, like what Envoy does with filters and LuaJIT. Instead of dealing with Lua directly, you can work with a smaller optimized runtime that supports a variety of programming languages. The Envoy project is currently in its journey in [adopting Wasm](https://www.solo.io/blog/the-state-of-webassembly-in-envoy-proxy/) and I expect a similar pattern to follow for any environment that scripting languages are a popular extension mechanism to be wholesale replaced by Wasm in the future.

[Web Assembly](https://webassembly.org/)(Wasm) 是一项新兴技术，但我预计它会成为云原生生态系统中不断增长的实用程序和工作负载，尤其是作为 [WASI](https://wasi.dev/) 成熟并且 Kubernetes 被更多地用作边缘编排器，如前所述。一个用例是为扩展机制提供支持，就像 Envoy 使用过滤器和 LuaJIT 所做的那样。您可以使用支持各种编程语言的较小的优化运行时，而不是直接处理 Lua。 Envoy 项目目前正在 [采用 Wasm](https://www.solo.io/blog/the-state-of-webassembly-in-envoy-proxy/) 的过程中，我希望遵循类似的模式任何环境中脚本语言是一种流行的扩展机制，将来会被 Wasm 完全取代。

On the Kubernetes front, there are projects like [Krustlet](https://deislabs.io/posts/introducing-krustlet/) from Microsoft that are exploring how a WASI-based runtime could be supported in Kubernetes. This shouldn’t be too surprising as Kubernetes is already being extended via CRDs and other mechanisms to run different types of workloads like VMs ( [KubeVirt](https://kubevirt.io/)) and more.

在 Kubernetes 方面，微软的 [Krustlet](https://deislabs.io/posts/introducing-krustlet/) 等项目正在探索如何在 Kubernetes 中支持基于 WASI 的运行时。这应该不足为奇，因为 Kubernetes 已经通过 CRD 和其他机制进行了扩展，以运行不同类型的工作负载，如 VM ([KubeVirt](https://kubevirt.io/)) 等。

Also, if you're new to Wasm, I recommend this new [intro course](https://www.edx.org/course/introduction-to-webassembly-runtime) from the Linux Foundation that goes over the space, along with the excellection documentation

另外，如果你是 Wasm 的新手，我推荐这个来自 Linux 基金会的新 [介绍课程](https://www.edx.org/course/introduction-to-webassembly-runtime)，它涵盖了整个领域，以及与优秀文件

**Rise of FinOps (CFM)**

**FinOps (CFM) 的兴起**

The coronavirus outbreak has accelerated the shift to cloud native. At least half of companies are accelerating their cloud plans amid the crisis… nearly 60% of respondents said cloud usage would exceed prior plans owing to the COVID-19 pandemic ( [State of the Cloud Report 2020](https://info.flexera.com/SLO-CM-REPORT-State-of-the-Cloud-2020)). On top of that, Cloud Financial Management (or FinOps) is a growing issue and [concern](https://www.wsj.com/articles/cloud-bills-will-get-loftier-1518363001) for many companies and honestly comes up in about half of my discussions the last six months with companies navigating their cloud native journey. You can also argue that cloud providers aren't incentivized to make cloud financial management easier as that would make it easier for customers to spend less, however, the true pain is lack of open source innovation and standardization around cloud financial management in my opinion ( all the clouds do cost management differently). In the CNCF context, there aren’t many open source projects trying to make FinOps easier, there is the [KubeCost](https://github.com/kubecost/cost-model) project but it’s fairly early days.

冠状病毒的爆发加速了向云原生的转变。在危机期间，至少有一半的公司正在加速他们的云计划……近 60% 的受访者表示，由于 COVID-19 大流行，云使用量将超过之前的计划（[2020 年云状况报告](https://info.flexera.com/SLO-CM-REPORT-State-of-the-Cloud-2020))。最重要的是，对于许多公司来说，云财务管理（或 FinOps）是一个日益严重的问题和 [关注](https://www.wsj.com/articles/cloud-bills-will-get-loftier-1518363001)，老实说在过去六个月中，我与公司进行云原生之旅的讨论中，大约有一半都提到了这一点。您也可以争辩说，云提供商没有被激励使云财务管理更容易，因为这会让客户更容易减少支出，但是，在我看来，真正的痛苦是缺乏围绕云财务管理的开源创新和标准化（所有的云以不同的方式进行成本管理)。在 CNCF 的背景下，没有多少开源项目试图让 FinOps 更容易，有 [KubeCost](https://github.com/kubecost/cost-model) 项目，但它是相当早期的。

Also, the Linux Foundation recently launched the “ [FinOps Foundation](https://www.finops.org/blog/linux-foundation)” to help innovation in this space, they have some [great introductory materials](https://www.edx.org/course/introduction-to-finops) in this space. I expect to see a lot more open source projects and specifications in the FinOps space in the coming years.

此外，Linux 基金会最近推出了“[FinOps 基金会](https://www.finops.org/blog/linux-foundation)”来帮助这个领域的创新，他们有一些[很棒的介绍材料](https://www.finops.org/blog/linux-foundation)/www.edx.org/course/introduction-to-finops)在这个领域。我希望在未来几年在 FinOps 领域看到更多的开源项目和规范。

**More Rust in Cloud Native** 

**云原生中的更多 Rust**

Rust is still a young and niche programming language, especially if you look at [programming language rankings](https://redmonk.com/sogrady/2020/07/27/language-rankings-6-20/) from Redmonk as an example. However, my feeling is that you will see Rust in more cloud native projects over the coming year given that there are already a handful of [CNCF projects taking advantage of Rust](https://www.cncf.io/blog/2020/06/22/rust-at-cncf/) to it popping up in interesting infrastructure projects like the microvm [Firecracker](https://firecracker-microvm.github.io/). While CNCF currently has a super majority of projects written in Golang, I expect Rust-based projects to be on par with Go-based ones in a couple of years as the [Rust community matures](https://blog.rust-lang.org/2020/08/18/laying-the-foundation-for-rusts-future.html).

Rust 仍然是一门年轻的小众编程语言，特别是如果你将 Redmonk 的 [编程语言排名](https://redmonk.com/sogrady/2020/07/27/language-rankings-6-20/) 看作是例子。但是，我的感觉是，鉴于已经有少数 [CNCF 项目利用 Rust](https://www.cncf.io/blog/2020/)，未来一年您将在更多云原生项目中看到 Rust 06/22/rust-at-cncf/) 到它出现在有趣的基础设施项目中，比如 microvm [Firecracker](https://firecracker-microvm.github.io/)。虽然 CNCF 目前有绝大多数项目是用 Golang 编写的，但我预计随着 [Rust 社区的成熟](https://blog.rust-lang.org/2020/08/18/laying-the-foundation-for-rusts-future.html)。

**GitOps + CD/PD Grows Significantly**

**GitOps + CD/PD 显着增长**

[GitOps](https://www.weave.works/blog/what-is-gitops-really) is an operating model for cloud native technologies, providing a set of best practices that unify deployment, management and monitoring for applications (originally [coined](https://www.weave.works/blog/gitops-operations-by-pull-request) by Alexis Richardson from Weaveworks fame). The most important aspect of GitOps is describing the desired system state versioned in Git via a declaration fashion, that essentially enables a complex set of system changes to be applied correctly and then verified (via a nice audit log enabled via Git and other tools). From a pragmatic standpoint, GitOps improves developer experience and with the growth of projects like Argo, GitLab, Flux and so on, I expect GitOps tools to hit the enterprise more this year. If you look at the [data](https://about.gitlab.com/blog/2020/07/14/gitops-next-big-thing-automation/) from say GitLab, GitOps is still a nascent practice where the majority of companies haven't explored it yet but as more companies move to adopt cloud native software at scale, GitOps will naturally follow in my opinion. If you're interested in learning more about this space, I recommend checking out the [newly](https://codefresh.io/devops/announcing-gitops-working-group/) formed [GitOps Working Group](https://github.com/gitops-working-group/gitops-working-group) in CNCF.

[GitOps](https://www.weave.works/blog/what-is-gitops-really) 是一种云原生技术的运营模式，提供了一套统一部署、管理和监控应用程序的最佳实践（最初[创造](https://www.weave.works/blog/gitops-operations-by-pull-request）由来自 Weaveworks 的 Alexis Richardson 提供）。 GitOps 最重要的方面是通过声明方式描述 Git 中版本化的所需系统状态，这基本上可以正确应用一组复杂的系统更改，然后进行验证（通过通过 Git 和其他工具启用的良好审计日志)。从实用的角度来看，GitOps 改善了开发者体验，随着 Argo、GitLab、Flux 等项目的增长，我预计 GitOps 工具今年将更多地冲击企业。如果您查看来自 GitLab 的 [数据](https://about.gitlab.com/blog/2020/07/14/gitops-next-big-thing-automation/)，GitOps 仍然是一种新生的实践，其中大多数公司还没有探索它，但随着越来越多的公司开始大规模采用云原生软件，我认为 GitOps 自然会跟进。如果你有兴趣了解更多关于这个领域的信息，我建议查看 [newly](https://codefresh.io/devops/announcing-gitops-working-group/) 组建的 [GitOps 工作组](https://github.com/gitops-working-group/gitops-working-group)在 CNCF 中。

**Service Catalogs 2.0: Cloud Native Developer Dashboards**

**服务目录 2.0：云原生开发人员仪表板**

The concept of a service catalog isn't a new thing, for some of us older folks that grew up in the [ITIL](https://en.wikipedia.org/wiki/ITIL) era you may remember things such as [ CMDBs](https://en.wikipedia.org/wiki/Configuration_management_database) (the horror). However, with the rise of microservices and cloud native development, the ability to catalog services and index a variety of real time service metadata is paramount to drive developer automation. This can include using a service catalog to understand ownership to handle incident management, manage SLOs and more. 

服务目录的概念并不是什么新鲜事物，对于我们这些在 [ITIL](https://en.wikipedia.org/wiki/ITIL) 时代长大的老年人来说，您可能还记得诸如 [ CMDB](https://en.wikipedia.org/wiki/Configuration_management_database)（恐怖)。然而，随着微服务和云原生开发的兴起，对服务进行编目和索引各种实时服务元数据的能力对于推动开发人员自动化至关重要。这可以包括使用服务目录来了解所有权以处理事件管理、管理 SLO 等。

In the future, you will see a trend towards developer dashboards that are not only a service catalog, but provide an ability to extend the dashboard through a variety of automation features all in one place. The canonical open source examples of this are [Backstage](https://engineering.atspotify.com/2020/03/17/what-the-heck-is-backstage-anyway/) and [Clutch](https://eng.lyft.com/announcing-clutch-the-open-source-platform-for-infrastructure-tooling-143d00de9713) from Lyft, however, any company with a fairly modern cloud native deployment tends to have a platform infrastructure team that has tried to build something similar. As the open source developer dashboards mature with a [large plug-in ecosystem](https://backstage.io/plugins), you’ll see accelerated adoption by platform engineering teams everywhere.

将来，您将看到开发人员仪表板的趋势，它不仅是一个服务目录，而且还提供通过各种自动化功能在一个地方扩展仪表板的能力。这方面的规范开源示例是 [Backstage](https://engineering.atspotify.com/2020/03/17/what-the-heck-is-backstage-anyway/) 和 [Clutch](https://eng.lyft.com/annoucing-clutch-the-open-source-platform-for-infrastructure-tooling-143d00de9713)来自 Lyft，但是，任何拥有相当现代的云原生部署的公司都倾向于拥有一个平台基础设施团队，他们已经尝试过建立类似的东西。随着开源开发人员仪表板随着[大型插件生态系统](https://backstage.io/plugins) 的成熟，您将看到各地平台工程团队加速采用。

**Cross Cloud Becomes More Real**

**跨云变得更加真实**

Kubernetes and the cloud native movement have demonstrated that cloud native and multi cloud approaches are possible in production environments, the data is clear that “93% of enterprises have a strategy to use multiple providers like Microsoft Azure, Amazon Web Services, and Google Cloud” ( [State of the Cloud Report 2020](https://info.flexera.com/SLO-CM-REPORT-State-of-the-Cloud-2020)). The fact that Kubernetes has matured over the years along with the cloud market, will hopefully unlock programmatic cross-cloud managed services. A concrete example of this approach is embodied in the Crossplane project that provides an open source cross cloud control plane taking advantage of the Kubernetes API extensibility to enable cross cloud workload management (see [“GitLab Deploys the Crossplane Control Plane to Offer Multicloud Deployments”] (https://thenewstack.io/gitlab-deploys-the-crossplane-control-plane-to-offer-multicloud-deployments/)).

Kubernetes 和云原生运动已经证明云原生和多云方法在生产环境中是可能的，数据清楚地表明“93% 的企业有使用多个提供商的策略，如 Microsoft Azure、亚马逊网络服务和谷歌云” ( [2020 年云状况报告](https://info.flexera.com/SLO-CM-REPORT-State-of-the-Cloud-2020))。 Kubernetes 多年来随着云市场的发展而成熟，这一事实有望解锁程序化的跨云托管服务。这种方法的一个具体示例体现在 Crossplane 项目中，该项目提供了一个开源跨云控制平面，利用 Kubernetes API 可扩展性来实现跨云工作负载管理（参见 [“GitLab 部署跨平面控制平面以提供多云部署”] （https://thenewstack.io/gitlab-deploys-the-crossplane-control-plane-to-offer-multicloud-deployments/）)。

**Mainstream eBPF**

**主流 eBPF**

[eBPF](https://en.wikipedia.org/wiki/Berkeley_Packet_Filter) allows you to run programs in the Linux Kernel without changing the kernel code or loading a module, you can think of it as a sandboxed extension mechanism. eBPF has allowed a [new generation of software](https://ebpf.io/projects) to extend the behavior of the Linux kernel to support a variety of different things from improved networking, monitoring and security. The downside of eBPF historically is that it requires a modern kernel version to take advantage of it and for a long time, that just wasn’t a realistic option for many companies. However, things are changing and even newer versions of RHEL finally support eBPF so you will see more projects take [advantage](https://sysdig.com/blog/sysdig-and-falco-now-powered-by-ebpf/) . If you look at the latest [container report](https://sysdig.com/blog/sysdig-2021-container-security-usage-report/) from Sysdig, you can see the adoption of Falco rising recently which although the report may be a bit biased from Sysdig, it is reflected in production usage. So stay tuned and look for more eBPF based projects in the future!

[eBPF](https://en.wikipedia.org/wiki/Berkeley_Packet_Filter) 允许你在不改变内核代码或加载模块的情况下在 Linux Kernel 中运行程序，你可以将其视为沙盒扩展机制。 eBPF 允许 [新一代软件](https://ebpf.io/projects) 扩展 Linux 内核的行为，以支持各种不同的事物，包括改进的网络、监控和安全性。从历史上看，eBPF 的缺点是它需要一个现代内核版本来利用它，而且在很长一段时间内，这对许多公司来说并不是一个现实的选择。然而，事情正在发生变化，甚至更新版本的 RHEL 终于支持 eBPF，所以你会看到更多的项目利用 [优势](https://sysdig.com/blog/sysdig-and-falco-now-powered-by-ebpf/) .如果您查看来自 Sysdig 的最新[容器报告](https://sysdig.com/blog/sysdig-2021-container-security-usage-report/)，您可以看到最近 Falco 的采用率上升，尽管报告可能与 Sysdig 有点偏差，这反映在生产使用上。所以请继续关注，并在未来寻找更多基于 eBPF 的项目！

**Finally, Happy 2021!**

**最后，2021 年快乐！**

I have a few more predictions and trends to share especially around end user driven open source, service mesh cannibalization/standardization, Prometheus+OTel, KYC for securing the software supply chain and more but I'll save that for more detailed posts, nine predictions are enough to kick off the new year! Anyways, thanks for reading and I hope to see everyone at [KubeCon + CloudNativeCon EU](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) in May 2021, registration is open! 

我还有更多的预测和趋势要分享，特别是围绕最终用户驱动的开源、服务网格蚕食/标准化、Prometheus+OTel、用于保护软件供应链的 KYC 等等，但我会将这些保存为更详细的帖子，九个预测足以开启新的一年！不管怎样，感谢阅读，希望在 2021 年 5 月在 [KubeCon + CloudNativeCon EU](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) 见到大家，注册开放！

