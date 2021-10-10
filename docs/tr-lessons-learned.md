# Lessons Learned From Two Years Of Kubernetes

# Kubernetes 两年的经验教训

2020-06-05

2020-06-05

As I come up for air after a few years of running an infrastructure team at Ridecell, I wanted to record some thoughts and lessons I’ve learned.

在 Ridecell 运营基础设施团队几年后，当我上台时，我想记录一些我学到的想法和经验教训。

1. [Kubernetes Is Not Just Hype](http://coderanger.net#kubernetes-is-not-just-hype)
2. [Traefik + Cert-Manager + Ext-DNS Is Great](http://coderanger.net#traefik--cert-manager--ext-dns-is-great)
3. [Prometheus Rocks, Thanos Is Not Overkill](http://coderanger.net#prometheus-rocks-thanos-is-not-overkill)
4. [GitOps Is The Way](http://coderanger.net#gitops-is-the-way)
5. [You Should Write More Operators](http://coderanger.net#you-should-write-more-operators)
6. [Secrets Management Is Still Hard](http://coderanger.net#secrets-management-is-still-hard)
7. [Native CI And Log Analysis Are Still Open Questions](http://coderanger.net#native-ci-and-log-analysis-are-still-open-questions)

1. [Kubernetes 不只是炒作](http://coderanger.net#kubernetes-is-not-just-hype)
2. [Traefik + Cert-Manager + Ext-DNS 很棒](http://coderanger.net#traefik--cert-manager--ext-dns-is-great)
3. [普罗米修斯摇滚，灭霸不矫枉过正](http://coderanger.net#prometheus-rocks-thanos-is-not-overkill)
4. [GitOps 就是方式](http://coderanger.net#gitops-is-the-way)
5. [你应该写更多的运算符](http://coderanger.net#you-should-write-more-operators)
6. [密文管理依旧难](http://coderanger.net#secrets-management-is-still-hard)
7. [Native CI 和日志分析仍然是未解决的问题](http://coderanger.net#native-ci-and-log-analysis-are-still-open-questions)

### Kubernetes Is Not Just Hype

### Kubernetes 不仅仅是炒作

I've been active in the [Kubernetes](http://coderanger.net/kubernetes.io/) world for a long time so this wasn't unexpected, but when something has this much hype train around it, it's always good to double check. Over two years, my team completed a total migration from Ansible+Terraform to pure Kubernetes, and in the process more than tripled our deployment rate while cutting deployment errors to “I can’t remember the last time we had one” levels. We also improved operational visibility, lots of boring-but-critical automation tasks, and mean time to recovery on infrastructure outages.

我在 [Kubernetes](http://coderanger.net/kubernetes.io/) 世界中活跃了很长时间，所以这并不意外，但是当有如此多的炒作火车时，它总是好的仔细检查。两年多来，我的团队完成了从 Ansible+Terraform 到纯 Kubernetes 的全面迁移，在此过程中，我们的部署率提高了两倍多，同时将部署错误减少到“我不记得上次我们有过一次”的水平。我们还改进了运营可见性、许多枯燥但关键的自动化任务，以及基础设施中断时的平均恢复时间。

Kubernetes is not magic, but it is an extremely powerful tool when used well by a team that knows it.

Kubernetes 并不神奇，但当它被一个了解它的团队很好地使用时，它是一个非常强大的工具。

### Traefik + Cert-Manager + Ext-DNS Is Great

### Traefik + Cert-Manager + Ext-DNS 很棒

The trio of [Traefik](https://containo.us/traefik/) as an Ingress Controller, [Cert-Manager](https://cert-manager.io/docs/) for generating certificates with LetsEncrypt, and [ External-DNS](https://github.com/kubernetes-incubator/external-dns) for managing edge DNS records makes HTTP routing and management smooth like butter. I’ve been fairly critical of Traefik 2.0’s choice to remove a lot 1.x annotation features however they have finally returned in 2.2, albeit in a different form. As an edge proxy, Traefik is a solid choice with great metrics integration, the fewest moving pieces of any Ingress Controller, and a responsive (if sometimes a bit K8s-clueless) dev team. Cert-Manager is a fantastic tool to use with any ingress approach. If you do TLS in your Kubernetes cluster and aren’t already using it, go check it out right now. External-DNS gets less glory than the other two pieces, but is no less important for automating the otherwise error-prone step of ensuring DNS records match reality.

[Traefik](https://containo.us/traefik/)作为入口控制器，[Cert-Manager](https://cert-manager.io/docs/) 使用 LetsEncrypt 生成证书的三重奏，以及 [ External-DNS](https://github.com/kubernetes-incubator/external-dns) 用于管理边缘 DNS 记录，使 HTTP 路由和管理像黄油一样顺畅。我一直对 Traefik 2.0 选择删除大量 1.x 注释功能持批评态度，但它们最终在 2.2 中返回，尽管形式不同。作为边缘代理，Traefik 是一个可靠的选择，它具有出色的指标集成、任何 Ingress Controller 中移动最少的部分以及响应迅速（有时有点 K8s 无能为力)的开发团队。 Cert-Manager 是一个很棒的工具，可以与任何入口方法一起使用。如果您在 Kubernetes 集群中使用 TLS 并且尚未使用它，请立即查看。外部 DNS 的荣耀不如其他两个部分，但对于自动化确保 DNS 记录匹配现实的容易出错的步骤同样重要。

If anything, these tools might actually make it too easy to set up new HTTPS endpoints. Over the years we ended up with dozens of unique certificates which created a lot of noise in things like Cert Transparency searches and LetsEncrypt’s own cert expiration warnings. Next time I will carefully consider which hostnames can be part of a globally configured wildcard certificate to reduce the total number of certificates in play.

如果有的话，这些工具实际上可能使设置新的 HTTPS 端点变得太容易了。多年来，我们最终得到了数十个独特的证书，这些证书在证书透明度搜索和 LetsEncrypt 自己的证书到期警告等方面产生了很多噪音。下次我将仔细考虑哪些主机名可以作为全局配置的通配符证书的一部分，以减少正在使用的证书总数。

### Prometheus Rocks, Thanos Is Not Overkill 

### 普罗米修斯摇滚，灭霸并不过分

This was my first time using [Prometheus](https://prometheus.io/) as the primary metrics system and it lived up to its reputation as the premier tool in that space. We went with [Prometheus-Operator](https://github.com/coreos/prometheus-operator) for managing it and that was also a great choice, making it a lot easier to distribute the scrape and rule configs into the applications that needed them. One thing I would do differently is using [Thanos](https://thanos.io/) from the beginning. I originally thought it would be overkill at first, but it was very easy to set up and was hugely helpful on both cross-region queries and reduced resource usage in Prometheus, even if we didn't jump directly to an active-active HA setup .

这是我第一次使用 [Prometheus](https://prometheus.io/) 作为主要指标系统，它不负众望，成为该领域的首要工具。我们使用 [Prometheus-Operator](https://github.com/coreos/prometheus-operator) 来管理它，这也是一个不错的选择，可以更轻松地将抓取和规则配置分发到应用程序中需要他们。我会做的一件事是从一开始就使用 [Thanos](https://thanos.io/)。起初我认为它会有点矫枉过正，但它很容易设置并且对跨区域查询和减少 Prometheus 中的资源使用都非常有帮助，即使我们没有直接跳转到主动-主动 HA 设置.

The biggest frustration I have with this part of the stack is [Grafana](https://grafana.com/) data management, how to store and organize the dashboards. There’s been a huge growth of tools for managing dashboards as YAML files, JSON files, Kubernetes custom objects, and probably anything else you can think of. But the underlying problem is still that it’s difficult to author a dashboard from scratch in any of those tools because Grafana has a million different config options and panel modes and whatnot. We ended up treating it as a stateful system as doing all dashboard management in-band, but I don’t really love that solution. Is there a workflow here that can be better?

我对堆栈的这一部分最大的挫折是 [Grafana](https://grafana.com/) 数据管理，如何存储和组织仪表板。用于将仪表板管理为 YAML 文件、JSON 文件、Kubernetes 自定义对象以及您能想到的任何其他工具的工具已经有了巨大的增长。但潜在的问题仍然是很难在任何这些工具中从头开始创作仪表板，因为 Grafana 有上百万种不同的配置选项和面板模式等等。我们最终将其视为一个有状态的系统，就像在带内进行所有仪表板管理一样，但我并不喜欢这种解决方案。这里有更好的工作流程吗？

### GitOps Is The Way

### GitOps 是一种方式

If you use Kubernetes, you should be practicing [GitOps](https://www.gitops.tech/). There's a wide range of tooling options, the simplest being a job in your existing CI system that runs `kubectl apply` all the way up to dedicated systems like [ArgoCD](https://argoproj.github.io/argo-cd/) and [Flux](https://docs.fluxcd.io/). I am firmly in the ArgoCD camp though, it was a solid tool to start with and over the years it has only gotten better. Just this week the first release is up for gitops-engine, putting ArgoCD and Flux both on a shared underlying system so it can get better even faster now, and if you don't like the workflows of either of those tools it is now even easier to build something new. A few months ago we had an accidental disaster recovery game-day from someone inadvertently deleting most of the namespaces in a test cluster and thanks to careful GitOps-ing our recovery was `make apply` in the bootstrap repo and wait for the system to rebuild itself. That said, some Velero backups are important too for stateful data that can’t live in git (eg. cert-manager’s certs, it could reissue everything but you might hit rate limits from LetsEncrypt).

如果你使用 Kubernetes，你应该练习 [GitOps](https://www.gitops.tech/)。有多种工具选项，最简单的是在现有 CI 系统中运行 `kubectl apply` 一直到像 [ArgoCD](https://argoproj.github.io/argo-cd/) 和 [Flux](https://docs.fluxcd.io/)。不过，我坚定地站在 ArgoCD 阵营，它是一个可靠的工具，多年来它一直在变得更好。就在本周，gitops-engine 发布了第一个版本，将 ArgoCD 和 Flux 都放在一个共享的底层系统上，这样它现在可以变得更快，如果你不喜欢这两个工具中的任何一个的工作流程，它现在甚至可以更容易构建新的东西。几个月前，我们有一个意外的灾难恢复游戏日，有人无意中删除了测试集群中的大部分命名空间，多亏了谨慎的 GitOps，我们的恢复是在引导程序存储库中“make apply”并等待系统重建本身。也就是说，一些 Velero 备份对于不能存在于 git 中的有状态数据也很重要（例如 cert-manager 的证书，它可以重新发布所有内容，但您可能会达到 LetsEncrypt 的速率限制)。

The biggest issue we had was with the choice to keep most of our core infrastructure in a single repo. I still think a single repo is the right design there, but I would divide things into different ArgoCD applications inside that rather than just having the one “infra” app. Having one app led to long(er) converge times and noisy UIs and had little benefit once we got used to splitting up our Kustomize definitions correctly.

我们遇到的最大问题是选择将我们的大部分核心基础设施保存在单个存储库中。我仍然认为单个 repo 是那里的正确设计，但我会将事物分成不同的 ArgoCD 应用程序，而不仅仅是拥有一个“infra”应用程序。拥有一个应用程序会导致较长的（er）收敛时间和嘈杂的 UI，并且一旦我们习惯了正确拆分 Kustomize 定义，几乎没有任何好处。

### You Should Write More Operators 

### 你应该写更多的操作符

I went in hard on custom operators from the start and we were hugely successful with them. We started with one custom resource and controller for deploying our main web application and slowly branched out to all the other automation needed for that application and others. Using plain Kustomize and ArgoCD for simple infrastructure services worked great, but we would reach for an operator any time we either wanted to control external things (ex. creating an AWS IAM role from Kubernetes, to be used via kiam) or when we needed some level of state machine for the thing (ex. Django application deployment with SQL migrations). As part of this we also built a very thorough test suite for all our custom objects and controllers which greatly improved operational stability and our own certainty that the system worked correctly.

我从一开始就努力使用自定义运算符，我们在他们身上取得了巨大的成功。我们从一个用于部署主 Web 应用程序的自定义资源和控制器开始，然后慢慢扩展到该应用程序和其他应用程序所需的所有其他自动化。将普通 Kustomize 和 ArgoCD 用于简单的基础设施服务效果很好，但我们会在任何时候想要控制外部事物（例如，从 Kubernetes 创建 AWS IAM 角色，通过 kiam 使用）或当我们需要一些操作时联系操作员事物的状态机级别（例如，带有 SQL 迁移的 Django 应用程序部署）。作为其中的一部分，我们还为我们所有的自定义对象和控制器构建了一个非常全面的测试套件，这大大提高了操作稳定性和我们自己对系统正常工作的确定性。

There's a lot more options for building operators these days, but I'm still very happy with [kubebuilder](https://book.kubebuilder.io/) (though to be fair, we did substantially modify the project structure over time so it's more fair to say it was using controller-runtime and controller-tools than kubebuilder itself). Whatever language and framework you feel most comfortable with, there is probably an operator toolkit available and you should absolutely use it.

如今，建筑运营商有更多选择，但我仍然对 [kubebuilder](https://book.kubebuilder.io/) 感到非常满意（尽管公平地说，随着时间的推移，我们确实对项目结构进行了大量修改，因此更公平地说，它使用的是控制器运行时和控制器工具，而不是 kubebuilder 本身)。无论您觉得最熟悉哪种语言和框架，都可能有一个操作员工具包可用，您绝对应该使用它。

### Secrets Management Is Still Hard

### 秘密管理仍然很难

Kubernetes has its own Secret object for managing secret data at runtime, using it with containers or with other objects, all that jazz. And that system works fine. But the long-term workflow for secrets is still kind of a mess. Committing a raw Secret to Git is bad for many reasons I hopefully don’t need to list, so how do we manage these objects? My solution was to develop a custom EncryptedSecret type which encrypted each value using AWS KMS along with a controller running in Kubernetes to decrypt back to a normal Secret so things work like usual, and a command line tool for the decrypt-edit-reencrypt cycle. Using KMS meant we could do access control via IAM rules restricting KMS key use, and encrypting only the values left the files reasonably diff-able. There are now some community operators based around [Mozilla Sops](https://github.com/mozilla/sops) that offer roughly the same workflow, though Sops is a little bit more frustrating on the local edit workflow. Overall this space still needs a lot of work, people should be expecting a workflow that is auditable, versioned, and code-reviewable like for everything else in GitOps land.

Kubernetes 有自己的 Secret 对象，用于在运行时管理秘密数据，将它与容器或其他对象一起使用，所有这些都是爵士乐。该系统运行良好。但是秘密的长期工作流程仍然有点混乱。向 Git 提交原始 Secret 是不好的，原因有很多，我希望不需要列出，那么我们如何管理这些对象？我的解决方案是开发一个自定义的 EncryptedSecret 类型，它使用 AWS KMS 和运行在 Kubernetes 中的控制器对每个值进行加密，以解密回正常的 Secret，以便像往常一样工作，以及用于解密-编辑-重新加密循环的命令行工具。使用 KMS 意味着我们可以通过限制 KMS 密钥使用的 IAM 规则进行访问控制，并且只加密值使文件具有合理的差异性。现在有一些基于 [Mozilla Sops](https://github.com/mozilla/sops) 的社区运营商提供大致相同的工作流程，尽管 Sops 在本地编辑工作流程上有点令人沮丧。总的来说，这个领域仍然需要大量工作，人们应该期待一个可审计、版本化和代码可审查的工作流，就像 GitOps 领域的其他一切一样。

As a related issue, the weaknesses of Kubernetes’ RBAC model are most apparent with Secrets. In almost all cases, the Secret being used for a thing must be in the same namespace as the thing using it, which often means Secrets for a lot of different things end up in the same namespace (database passwords, vendor API tokens, TLS certs ) and if you want to give someone (or something, same issue applies to operators) access to one, they get access to all. Keep your namespaces as small as possible. Anything that can go in its own namespace, do it. Your RBAC policies will thank you later.

作为一个相关问题，Kubernetes 的 RBAC 模型的弱点在 Secrets 中最为明显。在几乎所有情况下，用于事物的 Secret 必须与使用它的事物位于同一命名空间中，这通常意味着许多不同事物的 Secret 最终位于同一命名空间（数据库密码、供应商 API 令牌、TLS 证书）中) 并且如果你想让某人（或某事，同样的问题适用于运营商）访问一个，他们可以访问所有。使您的命名空间尽可能小。任何可以进入它自己的命名空间的东西，就去做吧。您的 RBAC 政策稍后会感谢您。

### Native CI And Log Analysis Are Still Open Questions 

### 原生 CI 和日志分析仍然是悬而未决的问题

Two big ecosystems holes I ran into are CI and log analysis. There’s lot of CI systems that deploy on Kubernetes, Jenkins, Concourse, Buildkite, etc. But there’s very few that feel like native solutions at all. [JenkinsX](https://jenkins-x.io/) is probably the closest to a native experience but it’s built on a mountain of complexity that I find very unfortunate. [Prow](https://github.com/kubernetes/test-infra/tree/master/prow) itself is also very native but also very bespoke so not a super easy thing to get started with. [Tekton Pipelines](https://tekton.dev/) and [Argo Workflows](https://argoproj.github.io/docs/argo/readme.html) both have the low-level plumbing in place for a native CI system but finding a way to expose that to my development teams never got beyond a theoretical operator. Argo-CI seems to be abandoned, but the Tekton team seems to be actively pursuing this use case so I’m hopeful for some improvement there.

我遇到的两个大的生态系统漏洞是 CI 和日志分析。有很多 CI 系统部署在 Kubernetes、Jenkins、Concourse、Buildkite 等上。但很少有人感觉像原生解决方案。 [JenkinsX](https://jenkins-x.io/) 可能是最接近原生体验的，但它建立在我觉得非常不幸的复杂性之上。 [Prow](https://github.com/kubernetes/test-infra/tree/master/prow) 本身也非常原生，但也非常定制，所以上手并不容易。 [Tekton Pipelines](https://tekton.dev/) 和 [Argo Workflows](https://argoproj.github.io/docs/argo/readme.html) 都为本地人提供了低级管道CI 系统，但找到一种方法将其公开给我的开发团队从未超越理论操作员。 Argo-CI 似乎被放弃了，但 Tekton 团队似乎正在积极地追求这个用例，所以我希望那里有一些改进。

Log collection is mostly a solved problem, with the community centralizing on [Fluent Bit](https://fluentbit.io/) as a DaemonSet shipping to some [Fluentd](https://www.fluentd.org/) pods which then send onwards to whatever systems you use for storage and analysis. On the storage side we've got [ElasticSearch](https://www.elastic.co/elasticsearch/) and [Loki](https://grafana.com/oss/loki/) as the main open contenders, each with their own analysis frontend ( [Kibana](https://www.elastic.co/kibana) and [Grafana](https://grafana.com/)). It’s mostly that last part that seems to still mostly be the source of my frustration. Kibana has been around much longer and has a good spread of analysis features, but you really have to use the commercial version to get even basic operational stuff like user authentication and per-user permissions are still very fuzzy. Loki is much newer and has even less in the way of analysis tools (substring searching and per-line tag searching) and nothing for permissions so far. If you’re careful to ensure that all log output is safe to be seen by all engineers this can be okay, but be ready for some pointed questions on your SOC/PCI/etc audits.

日志收集主要是一个已解决的问题，社区集中在 [Fluent Bit](https://fluentbit.io/) 作为 DaemonSet 运送到一些 [Fluentd](https://www.fluentd.org/) pods然后继续发送到您用于存储和分析的任何系统。在存储方面，我们有 [ElasticSearch](https://www.elastic.co/elasticsearch/) 和 [Loki](https://grafana.com/oss/loki/)作为主要的开放竞争者，每个使用他们自己的分析前端（[Kibana](https://www.elastic.co/kibana) 和 [Grafana](https://grafana.com/)）。主要是最后一部分似乎仍然是我沮丧的根源。 Kibana 已经存在的时间更长，并且具有广泛的分析功能，但是您确实必须使用商业版本才能获得甚至基本的操作性内容，例如用户身份验证和每用户权限仍然非常模糊。 Loki 更新得多，分析工具（子字符串搜索和每行标签搜索)的方式更少，到目前为止没有任何权限。如果您小心地确保所有工程师都可以安全地看到所有日志输出，这可能没问题，但要准备好回答有关 SOC/PCI/etc 审核的一些尖锐问题。

### In Closing

### 结束

Kubernetes is not the turnkey solution many pitch it to be, but with some careful engineering and a phenomenal community ecosystem, it can be a platform second to none. Take the time to learn each of the underlying components and you’ll be well on your way to container happiness, hopefully avoiding a few of my mistakes on the way.

Kubernetes 并不是许多人认为的交钥匙解决方案，但通过精心设计和非凡的社区生态系统，它可以成为首屈一指的平台。花点时间学习每个底层组件，你就会在通往容器幸福的道路上走得很好，希望能避免我在此过程中犯的一些错误。

* * *

* * *

[Back to articles](http://coderanger.net/)

[返回文章](http://coderanger.net/)

[Contact](http://coderanger.net/contact/) Copyright 2012-2020, Noah Kantrowitz 

[联系方式](http://coderanger.net/contact/) 版权所有 2012-2020, Noah Kantrowitz

