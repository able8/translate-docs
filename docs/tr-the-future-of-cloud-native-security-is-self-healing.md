# The Future of Cloud Native Security Is Self-Healing

# 云原生安全的未来是自我修复

#### 26 Oct 2020 9:29am,   by [Jon Jarboe](https://thenewstack.io/author/jonjarboe/ "Posts by Jon Jarboe")

#### 

![](https://cdn.thenewstack.io/media/2020/10/98663b08-woman-3918661_1280-1024x646.jpg)

Jon has been helping software development organizations improve processes and tools for over 20 years, in contexts ranging from embedded systems to complex distributed applications and roles including support, development, customer success, management, pre- and post-sales.](https://www.linkedin.com/in/jon-jarboe/)

20 多年来，Jon 一直在帮助软件开发组织改进流程和工具，从嵌入式系统到复杂的分布式应用程序和角色，包括支持、开发、客户成功、管理、售前和售后。

Security has always been a concern in the cloud. Individuals and organizations alike have a vested interest in securing their private information. Given the constant barrage of breach announcements, organizations clearly struggle to meet their data security obligations. Some argue that it’s an issue of priorities or regulation, but I think they are missing the point: it is simply too easy to fail when it comes to securing data in the cloud.

安全一直是云中的一个问题。个人和组织在保护他们的私人信息方面都有既得利益。鉴于不断涌现的违规公告，组织显然难以履行其数据安全义务。有些人认为这是一个优先事项或监管问题，但我认为他们没有抓住重点：在保护云中的数据时，它太容易失败了。

Breaches represent a real risk to organizations in regulated industries such as financial services and healthcare, with consequences to the business and even personal liability for senior leadership. Businesses in unregulated industries have faced fewer legal consequences, but legislators are showing an increasing appetite for consumer protections — which will affect all organizations. GDPR and CCPA are two recent examples, but they are certainly not the last. Eventually, every organization will likely have significant legal obligations that affect their approach to information security.

违规行为对金融服务和医疗保健等受监管行业的组织构成真正的风险，对业务产生影响，甚至对高级领导层造成个人责任。不受监管行业的企业面临的法律后果较少，但立法者对消费者保护的兴趣日益浓厚——这将影响到所有组织。 GDPR 和 CCPA 是最近的两个例子，但它们肯定不是最后一个。最终，每个组织都可能承担影响其信息安全方法的重要法律义务。

Whether you’re motivated by your obligation to users or your legal obligations, the goal is the same: to prevent breaches. More specifically, to reduce the opportunity for failure while delivering secure products more consistently.

无论您的动机是对用户的义务还是法律义务，目标都是一样的：防止违规。更具体地说，在更一致地交付安全产品的同时减少失败的机会。

## We’ve Faced Similar Challenges Before

## 我们以前也面临过类似的挑战

That goal may sound familiar because we’ve faced similar problems before. Not so long ago, organizations were under increasing pressure from a different direction; in fact, they still are. Competitive pressures demand that organizations deliver innovation more quickly, and economic pressures require that they do so more predictably and at a lower cost. Delivered products need to be reliable and available because the businesses of end users depend on them. These pressures led to the DevOps revolution and the rise of the cloud, which have been indisputable successes.

这个目标可能听起来很熟悉，因为我们以前也遇到过类似的问题。不久前，组织面临着来自不同方向的越来越大的压力。事实上，他们仍然是。竞争压力要求组织更快地交付创新，而经济压力要求他们以更可预测的方式以更低的成本进行创新。交付的产品需要可靠且可用，因为最终用户的业务依赖于它们。这些压力导致了 DevOps 革命和云的兴起，这些都取得了无可争辩的成功。

**Sponsor Note**

**赞助商备注**

![sponsor logo](https://cdn.thenewstack.io/media/2020/04/fe1154fd-accurics@2x.png)

Accurics enables compliance, governance, and security across the full cloud native stack in hybrid and multi-cloud environments. It seamlessly scans cloud automation code for risks before the stack is provisioned and monitors production cloud stacks for changes that introduce risk.

Accurics 在混合和多云环境中的整个云原生堆栈中实现合规性、治理和安全性。它在配置堆栈之前无缝扫描云自动化代码中的风险，并监控生产云堆栈是否有引入风险的更改。

In 2019, the Cloud Native Computing Foundation (CNCF) [surveyed the industry](https://www.cncf.io/wp-content/uploads/2020/08/CNCF_Survey_Report.pdf) and found that cloud native technologies are widespread. In production environments, 84% of organizations are using containers; 78% use Kubernetes, and at least 41%, serverless. Industry analysts such as [IDC expect](https://www.idc.com/getdoc.jsp?containerId=prUS45613519) these trends to continue, with more than 90% of apps being cloud native by 2025 and two-thirds of enterprises deploying daily. Clearly, the cloud is the future of software.

2019年，云原生计算基金会（CNCF） [行业调研](https://www.cncf.io/wp-content/uploads/2020/08/CNCF_Survey_Report.pdf)发现云原生技术普及。在生产环境中，84% 的组织正在使用容器； 78% 使用 Kubernetes，至少 41% 使用无服务器。 [IDC 预计](https://www.idc.com/getdoc.jsp?containerId=prUS45613519) 等行业分析师将继续这些趋势，到 2025 年将有超过 90% 的应用和三分之二的企业实现云原生每天部署。显然，云是软件的未来。

While this is fantastic for the delivery of innovation, it complicates the problem of securing it. The good news is that many of the challenges that complicate effective security processes — such as manual processes, siloed expertise, and poor communication — have been solved before. These are the same challenges affecting development, quality and operations teams, that DevOps is successfully overcoming.

虽然这对于创新的交付来说非常棒，但它使保护创新的问题变得复杂。好消息是，许多使有效安全流程复杂化的挑战——例如手动流程、孤立的专业知识和沟通不畅——之前已经得到解决。这些都是影响开发、质量和运营团队的挑战，而 DevOps 正在成功克服这些挑战。

## A Way Forward 

## 前进的道路

In many ways, securing systems can be seen as another iteration of DevOps — breaking up the security silo and integrating that expertise into the DevOps process. And much like earlier iterations of DevOps, it will require new approaches, tools and adjustments. As with everything DevOps, there are no silver bullets or “plug and play” solutions. Every team is different, and the best solution for any team will necessarily be built upon their unique needs and capabilities.

在许多方面，保护系统可以被视为 DevOps 的另一个迭代——打破安全孤岛并将该专业知识集成到 DevOps 流程中。与 DevOps 的早期迭代非常相似，它将需要新的方法、工具和调整。与所有 DevOps 一样，没有灵丹妙药或“即插即用”解决方案。每个团队都是不同的，任何团队的最佳解决方案都必须建立在他们独特的需求和能力之上。

That said, we’re not lost in the wilderness. Many organizations have at least started their DevOps journey and have a basic understanding of what works for them. They probably have some experience with the impact that these changes will have on teams and schedules. They hopefully recognize the direction in which success lies. Those with experience in these areas — namely, the developers and DevOps teams that have already navigated these waters — can help lead the way to more secure releases.

也就是说，我们并没有迷失在荒野中。许多组织至少已经开始了他们的 DevOps 之旅，并且对什么适合他们有基本的了解。他们可能对这些变化对团队和日程安排的影响有一定的经验。他们希望能认清成功的方向。那些在这些领域拥有经验的人——即已经在这些水域中航行的开发人员和 DevOps 团队——可以帮助引领更安全的发布之路。

We know that DevOps strategies — such as optimizing manual processes and automating them, establishing continuous feedback loops, and improving collaboration — have measurable results. Let’s start there.

我们知道 DevOps 策略——例如优化手动流程并使其自动化、建立持续的反馈循环和改善协作——具有可衡量的结果。让我们从那里开始。

One of the most disruptive aspects of traditional security processes is the late-stage security review. It's hard to start early because you need to review the actual behavior, so many security teams assess security after deployment — in the staging or production environment (occasionally even in the live customer environment!) When issues are found, there's very little context and it requires a lot of research to figure out what to do about them. There is a lot of room for improvement here: in the manual assessment process, the manual reporting and tracking process, and the manual remediation process. Truth be told, most of that effort is wasted because the issues are rarely fixed — they end up in a large backlog of issues that are never scheduled.

传统安全流程最具破坏性的方面之一是后期安全审查。很难尽早开始，因为您需要审查实际行为，因此许多安全团队在部署后评估安全性 - 在登台或生产环境中（有时甚至在实时客户环境中！）当发现问题时，几乎没有上下文需要进行大量研究才能弄清楚如何处理它们。这里有很大的改进空间：在人工评估过程中，在人工报告和跟踪过程中，在人工补救过程中。说实话，大部分努力都白费了，因为问题很少得到解决——它们最终会积压大量从未安排过的问题。

## The Infrastructure Can Heal Itself

## 基础设施可以自愈

It doesn’t have to be that way. Infrastructure-as-code (IaC) — the definition of resources and relationships in a codified format — is foundational for modern software teams, as a way to consistently and reliably build or rebuild the infrastructure as needed.

不必如此。基础设施即代码 (IaC) - 以编码格式定义资源和关系 - 是现代软件团队的基础，是根据需要一致可靠地构建或重建基础设施的一种方式。

In addition to enabling continuous delivery and continuous deployment, IaC can also improve the security review process, because it defines what the runtime environment will look like. Security teams can start threat modeling and security assessments long before deployment actually happens, which means DevOps teams receive feedback in contexts where they can act on it.

除了支持持续交付和持续部署，IaC 还可以改进安全审查流程，因为它定义了运行时环境的样子。安全团队可以在部署实际发生之前很久就开始威胁建模和安全评估，这意味着 DevOps 团队会在可以对其采取行动的上下文中收到反馈。

That’s an improvement, but it’s still fundamentally a disruptive process. After a security problem is identified and reported, somebody needs to fix it with urgency. The level of effort required to manually fix that problem will necessarily displace other planned work, and automation will only make things worse. What’s needed is a way to not only identify issues early but also eliminate the manual work required to investigate and fix the problem. That is, to have a self-healing infrastructure that can not only identify problems but fix them autonomously so that planned work continues apace.

这是一种改进，但从根本上说它仍然是一个破坏性的过程。在确定并报告安全问题后，需要有人紧急修复它。手动修复该问题所需的工作量必然会取代其他计划的工作，而自动化只会让事情变得更糟。我们需要一种方法，不仅可以及早发现问题，还可以消除调查和解决问题所需的手动工作。也就是说，拥有一个自我修复的基础设施，不仅可以识别问题，而且可以自主修复它们，以便计划的工作继续进行。

## Self-Healing Is the Future

## 自愈是未来

As long as security relies on expert tools and processes, it cannot possibly fit into modern DevOps workflows. Modern software teams rely on infrastructure-as-code to improve consistency and scalability, and we can leverage those same workflows to improve security. By embedding appropriate security controls into the development process, that infrastructure can self-diagnose and self-heal — eliminating the all-too-common bottlenecks that limit the effectiveness of today’s security programs. 

只要安全依赖于专家工具和流程，它就不可能融入现代 DevOps 工作流程。现代软件团队依靠基础设施即代码来提高一致性和可扩展性，我们可以利用这些相同的工作流来提高安全性。通过在开发过程中嵌入适当的安全控制，该基础设施可以自我诊断和自我修复——消除限制当今安全程序有效性的非常常见的瓶颈。

