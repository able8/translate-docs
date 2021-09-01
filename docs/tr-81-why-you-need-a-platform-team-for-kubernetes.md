# Why you need a platform team for Kubernetes

# 为什么你需要一个 Kubernetes 平台团队

Setting up a Kubernetes cluster can be deceptively simple, as there are plenty of installers to create a basic cluster in minutes. However, that’s only the start of the actual work. Kubernetes moves fast; when it’s a critical part of your infrastructure, there’s a host of things you need to look out for to maintain a healthy cluster. More often than not, it’s wise to have a dedicated team to run Kubernetes.

设置 Kubernetes 集群看似简单，因为有很多安装程序可以在几分钟内创建一个基本集群。然而，这只是实际工作的开始。 Kubernetes 行动迅速；当它是基础架构的关键部分时，您需要注意许多事项以维护健康的集群。通常情况下，拥有一个专门的团队来运行 Kubernetes 是明智的。

## Integrations

## 集成

For most use cases, Kubernetes is not enough on its own. It might be better to think of it as a framework for building platforms. You will need to integrate Kubernetes with other tools and systems to make it useful. The exact integrations depend on your use case — there’s no silver bullet. You’ll need to figure out your own optimal setup.

对于大多数用例，Kubernetes 本身是不够的。将其视为构建平台的框架可能会更好。您需要将 Kubernetes 与其他工具和系统集成以使其有用。确切的集成取决于您的用例——没有灵丹妙药。您需要找出自己的最佳设置。

To give you an idea, here are a few common integrations you might need to set up and configure:
- ExternalDNS for interfacing with cloud DNS services, to have human-readable addresses for services
- Let’s Encrypt for automatic HTTPS
- A network plugin such as Calico or Flannel
- A CI/CD system such as Spinnaker, GitLab, or Jenkins to deploy to Kubernetes
- Authentication with LDAP, AD, or another system for using your company’s regular user accounts
- Prometheus for monitoring
- Alertmanager for alerting

为了给您一个想法，以下是您可能需要设置和配置的一些常见集成：
- 用于与云 DNS 服务交互的 ExternalDNS，为服务提供人类可读的地址
- 让我们为自动 HTTPS 加密
- 网络插件，如 Calico 或 Flannel
- 部署到 Kubernetes 的 CI/CD 系统，例如 Spinnaker、GitLab 或 Jenkins
- 使用 LDAP、AD 或其他系统进行身份验证以使用您公司的常规用户帐户
- 用于监控的 Prometheus
- Alertmanager 用于提醒

Granted, you might not need all of these, and some of them may be provided by your Kubernetes distribution out of the box. Still, if you end up using them, you’ll need to understand how they work and be able to operate them. All of this integration work adds a lot of complexity, and you’ll need to spend a lot of time learning how it all works. It can be difficult for one person to grasp all the details of the entire system, so it’s better to share the effort with multiple people familiar with Kubernetes and its integrations.

当然，您可能不需要所有这些，其中一些可能由您的 Kubernetes 发行版提供。尽管如此，如果您最终使用它们，则需要了解它们的工作原理并能够操作它们。所有这些集成工作都增加了很多复杂性，您需要花费大量时间来了解它是如何工作的。一个人可能很难掌握整个系统的所有细节，因此最好与熟悉 Kubernetes 及其集成的多个人共同努力。

## Keeping up with the times

## 与时俱进

Kubernetes is still a fast-moving project. There are four releases every year – old versions are maintained only for a year, after which there are no security updates. While running a managed service, such as GKE, EKS, or AKS, helps significantly with upgrading, it’s still not entirely effortless.

Kubernetes 仍然是一个快速发展的项目。每年有四个版本——旧版本只维护一年，之后没有安全更新。虽然运行 GKE、EKS 或 AKS 等托管服务对升级有很大帮助，但仍然不是完全不费吹灰之力。

Keeping a Kubernetes cluster updated and secure is complex. Besides keeping Kubernetes itself up to date, there’s also work involved in upgrading all the integrations needed to create a fully functional platform. Running a managed Kubernetes service does not necessarily help with keeping them updated. Another continuous update cycle comes from the various container images involved in running the software that is the reason for running Kubernetes in the first place.

保持 Kubernetes 集群的更新和安全是很复杂的。除了让 Kubernetes 保持最新状态外，还需要升级创建功能齐全的平台所需的所有集成。运行托管 Kubernetes 服务不一定有助于保持更新。另一个持续的更新周期来自于运行软件所涉及的各种容器镜像，这也是运行 Kubernetes 的首要原因。

There are good reasons for keeping all the software involved up to date: Kubernetes clusters have a lot of potential for attacks due to their complex and distributed nature. On top of that, they make for attractive targets if looking for computing power. One example of attacks against Kubernetes is cryptojacking where attackers hijack clusters to mine cryptocurrency. [Tesla famously fell victim to such an attack a couple of years ago](https://redlock.io/blog/cryptojacking-tesla) and recently such attacks have become more sophisticated as evidenced by the [Hildegard malware](https://unit42.paloaltonetworks.com/hildegard-malware-teamtnt/).

将所有涉及的软件保持最新是有充分理由的：Kubernetes 集群由于其复杂和分布式的特性而具有很大的攻击潜力。最重要的是，如果寻找计算能力，它们会成为有吸引力的目标。 Kubernetes 攻击的一个例子是加密劫持，攻击者劫持集群以挖掘加密货币。 [特斯拉在几年前成为此类攻击的受害者](https://redlock.io/blog/cryptojacking-tesla)，最近此类攻击变得更加复杂，[Hildegard 恶意软件](https://unit42.paloaltonetworks.com/hildegard-malware-teamtnt/)。

## The human aspect

## 人的方面

As with any IT system you operate, you’ll need someone who can keep it running at all times. As Kubernetes is a platform on top of which you run other services, this is even more important: outages have a broader impact and might impair your operations significantly.

与您操作的任何 IT 系统一样，您需要一个可以始终保持运行的人。由于 Kubernetes 是一个平台，您可以在其上运行其他服务，因此这一点更为重要：中断会产生更广泛的影响，并可能显着损害您的运营。

As a rule of thumb, you’ll need at least three people who can independently solve Kubernetes-related issues — but you’ll likely need more than that. Consider these scenarios as a member of a Kubernetes maintenance team of three:

根据经验，您至少需要三个可以独立解决 Kubernetes 相关问题的人——但您可能需要的不止这些。作为 Kubernetes 维护团队的三人成员，请考虑以下场景：

- One of your teammates leaves for another job. How long will it take to hire and/or train a replacement?

- 你的一个队友离开去换一份工作。雇用和/或培训替代人员需要多长时间？

- Your team has a face-to-face meeting. It turns out one of your teammates was coming down with the flu and inadvertently spread it to your other colleague: now two out of your team of three have the flu and you’re alone. 

- 您的团队有一个面对面的会议。事实证明，您的一名队友感染了流感，并在不经意间将其传染给了您的另一位同事：现在，您的三人团队中有两人感染了流感，而您却是一个人。

- You have a teammate on vacation and another one sick. There’s an issue related to one of the integrations, but you don’t know it too well – your coworker who’s drinking mai tais on the beach is the expert on that one. Consider the lead time for contacting your teammate, getting them on a computer, and in the mindset to solve your problem.


- 你有一个队友在度假，另一个生病了。有一个与其中一项集成相关的问题，但您不太了解 - 您在海滩上喝麦太酒的同事是该问题的专家。考虑联系您的队友、让他们使用计算机以及解决问题的心态的准备时间。


It’s not unreasonable to assume that you’ll have someone available most of the time, but exceptions are always around the corner. There is always a significant risk of running into a situation where you don’t have enough people around to solve an incident promptly. You’ll need to adjust your availability target accordingly. If you need 24/7 operations, three people likely won’t cut it.

假设您大部分时间都有人可用并不是没有道理的，但例外总是在拐角处。遇到这样的情况总是存在很大的风险，即您周围没有足够的人来及时解决事件。您需要相应地调整可用性目标。如果您需要 24/7 全天候运营，三个人可能不会削减它。

## No dedicated platform team?

## 没有专门的平台团队？

By this point, I hope I’ve convinced you that operating and maintaining a Kubernetes cluster is not a one-person operation. Even if you try not having a dedicated team by spreading the knowledge across multiple development teams, you still need someone available to troubleshoot issues at all times. What would be the benefits of having a dedicated Kubernetes platform team instead of a distributed model?

至此，我希望我已经让您相信，操作和维护 Kubernetes 集群不是一个人的操作。即使您通过在多个开发团队中传播知识来尝试不拥有一个专门的团队，您仍然需要有人随时可以解决问题。拥有一个专门的 Kubernetes 平台团队而不是分布式模型有什么好处？

One significant downside to distributing Kubernetes maintainers across multiple teams is the increased amount of context switching. This is a recipe for a constant tug of war between running Kubernetes and the main team’s duties (if they’re only tasked with Kubernetes, why not just put them in a dedicated platform team instead?).

将 Kubernetes 维护人员分配到多个团队的一个显着缺点是上下文切换量增加。这是运行 Kubernetes 和主要团队职责之间不断拉锯战的秘诀（如果他们只负责 Kubernetes，为什么不把他们放在一个专门的平台团队中呢？）。

Constant context switching gets stressful as the maintainers need to make tough decisions about what to prioritize. This is exacerbated by the fact that operations work is unevenly distributed and unpredictable: Often things are quiet and systems keep humming along without any issues, but an incident can suddenly and unexpectedly turn your week upside down.

All in all, it’s simply hard to keep up with Kubernetes while also juggling other work.

持续的上下文切换会带来压力，因为维护人员需要就优先事项做出艰难的决定。运营工作分布不均且不可预测的事实加剧了这种情况：通常情况下，事情很安静，系统会一直运行而没有任何问题，但一个事件可能会突然和意外地使您的一周发生翻天覆地的变化。

总而言之，在处理其他工作的同时，很难跟上 Kubernetes 的步伐。

## Conclusion

##  结论

If your organization is large enough and you can afford to have a dedicated team to maintain Kubernetes, it will save you a lot of time and effort compared to other options for managing computing resources. You need to spend time to save time: there are a certain initial investment and ongoing maintenance overhead to keep things running smoothly.

如果您的组织足够大，并且您有能力拥有一个专门的团队来维护 Kubernetes，那么与管理计算资源的其他选项相比，它将为您节省大量时间和精力。你需要花时间来节省时间：有一定的初始投资和持续的维护开销来保持事情的顺利进行。

It’s worth noting that while having a platform team has its benefits, you should be careful not to build it into a silo. It’s best to think of the Kubernetes platform as an internal service with internal customers you need to keep happy. And for that, you need to collaborate closely with them to make sure you are building a good platform for their specific use cases.

值得注意的是，虽然拥有平台团队有其好处，但您应该小心不要将其构建为孤岛。最好将 Kubernetes 平台视为一种内部服务，您需要让内部客户保持满意。为此，您需要与他们密切合作，以确保为他们的特定用例构建一个良好的平台。

If you have a smaller organization and can’t yet justify a dedicated team just for Kubernetes, it might mean sacrifices in the quality and reliability of the platform. Consider your options carefully: it might not be possible to get a production-ready platform out of Kubernetes (depending on what “production-ready” means for you).


如果您的组织规模较小，并且还不能证明专门为 Kubernetes 设立一个团队是合理的，这可能意味着要牺牲平台的质量和可靠性。仔细考虑您的选择：可能无法从 Kubernetes 中获得生产就绪平台（取决于“生产就绪”对您意味着什么）。

---

Risto Laurikainen is a DevOps Consultant with a decade of experience in building cloud computing platforms. 

Risto Laurikainen 是一名 DevOps 顾问，在构建云计算平台方面拥有十年经验。


