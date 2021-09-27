# How to adopt DevSecOps successfully

# 如何成功采用 DevSecOps

## Integrating security throughout the software development lifecycle is important, but it's not always easy.

## 在整个软件开发生命周期中集成安全性很重要，但并不总是那么容易。

12 Feb 2021[Mike Calizo (Red Hat, Correspondent)](http://opensource.com/users/mcalizo "View user profile.") [Feed](http://opensource.com/user/456981/feed)

2021 年 2 月 12 日[Mike Calizo（红帽通讯员）](http://opensource.com/users/mcalizo“查看用户资料。”)[Feed](http://opensource.com/user/456981/feed)

Adopting [DevOps](https://opensource.com/resources/devops) can help an organization transform and speed how its software is delivered, tested, and deployed to production. This is the well-known "DevOps promise" that has led to such a large surge in adoption.

采用 [DevOps](https://opensource.com/resources/devops) 可以帮助组织转变并加快其软件交付、测试和部署到生产的方式。这就是众所周知的“DevOps 承诺”，它导致了如此大规模的采用。

We've all heard about the many successful DevOps implementations that changed how an organization approaches software innovation, making it fast and secure through agile delivery to get [ahead of competitors](https://www.imd.org/research-knowledge/articles/the-battle-for-digital-disruption-startups-vs-incumbents/). This is where we see DevOps' promises achieved and delivered.

我们都听说过许多成功的 DevOps 实施，它们改变了组织进行软件创新的方式，通过敏捷交付使其快速和安全，从而[领先于竞争对手](https://www.imd.org/research-knowledge/文章/数字中断之战-初创公司与现任者/)。这就是我们看到 DevOps 承诺实现和交付的地方。

But on the flipside, some DevOps adoptions cause more issues than benefits. This is the DevOps dilemma where DevOps fails to deliver on its promises.

但另一方面，一些 DevOps 的采用带来的问题多于好处。这是 DevOps 困境，DevOps 未能兑现其承诺。

There are many factors involved in an unsuccessful DevOps implementation, and a major one is security. A poor security culture usually happens when security is left to the end of the DevOps adoption process. Applying existing security processes to DevOps can delay projects, cause frustrations within the team, and create financial impacts that can derail a project.

不成功的 DevOps 实施涉及许多因素，其中一个主要因素是安全性。当安全性被留到 DevOps 采用过程的末尾时，通常会发生糟糕的安全文化。将现有的安全流程应用于 DevOps 可能会延迟项目，在团队中造成挫折，并产生可能使项目脱轨的财务影响。

[DevSecOps](http://www.devsecops.org/blog/2015/2/15/what-is-devsecops) was designed to avoid this very situation. Its purpose "is to build on the mindset that 'everyone is responsible for security…'" It also makes security a consideration at all levels of DevOps adoption.

[DevSecOps](http://www.devsecops.org/blog/2015/2/15/what-is-devsecops) 旨在避免这种情况。它的目的是“建立在‘每个人都对安全负责……’的心态之上”，它还使安全成为 DevOps 采用的各个级别的考虑因素。

## The DevSecOps process

## DevSecOps 流程

Before DevOps and DevSecOps, the app security process looked something like the image below. Security came late in the software delivery process, after the software was accepted for production.

在 DevOps 和 DevSecOps 之前，应用安全流程类似于下图。在软件被接受用于生产之后，安全性出现在软件交付过程的后期。

## [devsecops\_old-process.png](http://opensource.com/file/494846)

## [devsecops\_old-process.png](http://opensource.com/file/494846)

![Old software development process with security at the end](https://opensource.com/sites/default/files/uploads/devsecops_old-process.png)

(Michael Calizo, [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/))

（迈克尔卡利佐，[CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/))

Depending on the organization's security profile and risk appetite, the application might even bypass security reviews and processes during acceptance. At that point, the security review becomes an audit exercise to avoid unnecessary project delays.

根据组织的安全状况和风险偏好，应用程序甚至可能在验收期间绕过安全审查和流程。那时，安全审查就变成了一种审计活动，以避免不必要的项目延误。

## [devsecops\_security-as-audit.png](http://opensource.com/file/494851)

## [devsecops\_security-as-audit.png](http://opensource.com/file/494851)

![Security as audit in software development](https://opensource.com/sites/default/files/uploads/devsecops_security-as-audit.png)

(Michael Calizo, [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/))

（迈克尔卡利佐，[CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/))

The DevSecOps [manifesto](https://www.devsecops.org/) says that the reason to integrate security into dev and ops at all levels is to implement security with less friction, foster innovation, and make sure security and data privacy are not left behind.

DevSecOps [manifesto](https://www.devsecops.org/) 表示，将安全性集成到所有级别的开发和运维中的原因是实现安全性，减少摩擦，促进创新，并确保安全和数据隐私没有留下。

Therefore, DevSecOps encourages security practitioners to adapt and change their old, existing security processes and procedures. This may be sound easy, but changing processes, behavior, and culture is always difficult, especially in large environments.

因此，DevSecOps 鼓励安全从业者适应和改变他们旧的、现有的安全流程和程序。这听起来可能很容易，但改变流程、行为和文化总是很困难，尤其是在大型环境中。

The DevSecOps principle's basic requirement is to introduce a security culture and mindset across the entire application development and deployment process. This means old security practices must be replaced by more agile and flexible methods so that security can iterate and adapt to the fast-changing environment. According to the DevSecOps manifesto, security needs to "operate like developers to make security and compliance available to be consumed as services."

DevSecOps 原则的基本要求是在整个应用程序开发和部署过程中引入安全文化和思维方式。这意味着旧的安全实践必须被更敏捷和灵活的方法所取代，以便安全可以迭代并适应快速变化的环境。根据 DevSecOps 宣言，安全性需要“像开发人员一样运作，使安全性和合规性可作为服务使用”。

DevSecOps should look like the figure below, where security is embedded across the delivery cycle and can iterate every time there is a need for change or adjustment.

DevSecOps 应该如下图所示，其中安全性嵌入在整个交付周期中，并且可以在每次需要更改或调整时进行迭代。

## [devsecops\_process.png](http://opensource.com/file/494856)

## [devsecops\_process.png](http://opensource.com/file/494856)

![DevSecOps considers security throughout development](https://opensource.com/sites/default/files/uploads/devsecops_process.png)

(Michael Calizo, [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/))

（迈克尔卡利佐，[CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/))

## Common DevSecOps obstacles

## 常见的 DevSecOps 障碍

More on security

更多关于安全

- [The defensive coding guide](https://developers.redhat.com/articles/defensive-coding-guide/?intcmp=70160000000h1s6AAA) 

- [防御性编码指南](https://developers.redhat.com/articles/defensive-coding-guide/?intcmp=70160000000h1s6AAA)

- [Webinar: Automating system security and compliance with a standard operating system](https://www.redhat.com/en/events/webinar/automating-system-security-and-compliance-standard-operating-system?intcmp=70160000000h1s6AAA)
- [10 layers of Linux container security](https://www.redhat.com/en/resources/container-security-openshift-cloud-devops-whitepaper?intcmp=70160000000h1s6AAA)
- [SELinux coloring book](https://developers.redhat.com/books/selinux-coloring-book?intcmp=70160000000h1s6AAA)
- [More security articles](https://opensource.com/tags/security?intcmp=70160000000h1s6AAA)

- [网络研讨会：使用标准操作系统自动化系统安全性和合规性](https://www.redhat.com/en/events/webinar/automating-system-security-and-compliance-standard-operating-system?intcmp=70160000000h1s6AAA)
-  [10层Linux容器安全](https://www.redhat.com/en/resources/container-security-openshift-cloud-devops-whitepaper?intcmp=70160000000h1s6AAA)
- [SELinux 图画书](https://developers.redhat.com/books/selinux-coloring-book?intcmp=70160000000h1s6AAA)
- [更多安全文章](https://opensource.com/tags/security?intcmp=70160000000h1s6AAA)

Any time changes are introduced, people find faults or issues with the new process. This is natural human behavior. The fear and inconvenience associated with learning new things are always met with adverse reactions; after all, humans are creatures of habit.

任何时候引入更改，人们都会发现新流程的缺陷或问题。这是人类的自然行为。与学习新事物相关的恐惧和不便总是会遇到不良反应；毕竟，人类是习惯的生物。

Some common obstacles in DevSecOps adoption include:

采用 DevSecOps 的一些常见障碍包括：

- **Vendor-defined DevOps/DevSecOps:** This means principles and processes are focused on product offerings, and the organization won't be able to build the approach. Instead, they will be limited to what the vendor provides.
- **Nervous people managers:** The fear of losing control is a real problem when change happens. Often, anxiety affects people managers' decision-making.
- **If ain't broke, don't fix it:** This is a common mindset, and you really can't blame people for thinking this way. But the idea that the old way will survive despite new ways of delivering software and solutions must be challenged. To adapt to the agile application lifecycle, you need to change the processes to support the speed and agility it requires.
- **The Netflix and Uber effect:** Everybody knows that Netflix and Uber have successfully implemented DevSecOps; therefore, many organizations want to emulate them. Because they have a different culture than your organization, simply emulating them won't work.
- **Lack of measurement:** DevOps and DevSecOps transformation must be measured against set goals. Metrics might include software delivery performance or overall organization performance over time.
- **Checklist-driven security:** By using a checklist, the security team follows the same old, static, and inflexible processes that are neither useful nor applicable to modern technologies that developers use to make software delivery lean and agile. The introduction of the " [as code](https://www.oreilly.com/library/view/devopssec/9781491971413/ch04.html)" approach requires security people to learn how to code.
- **Security as a special team:** This is especially true in organizations transitioning from the old ways of delivering software, where security is a separate entity, to DevOps. Because of the separations, trust is questionable among devs, ops, and security. This will cause the security team to spend unnecessary time reviewing and governing DevOps processes and building pipelines instead of working closely with developers and ops teams to improve the software delivery flow.

- **供应商定义的 DevOps/DevSecOps：** 这意味着原则和流程专注于产品供应，组织将无法构建该方法。相反，它们将仅限于供应商提供的内容。
- **神经质的人事经理：** 当变革发生时，害怕失去控制是一个真正的问题。通常，焦虑会影响人员经理的决策。
- **如果没有破产，就不要修复它：** 这是一种普遍的心态，你真的不能责怪人们有这种想法。但是，尽管交付软件和解决方案采用新方式，但旧方式仍将继续存在的想法必须受到挑战。为了适应敏捷应用程序生命周期，您需要更改流程以支持其所需的速度和敏捷性。
- **Netflix 和 Uber 效应：** 大家都知道 Netflix 和 Uber 已经成功实施了 DevSecOps；因此，许多组织都想效仿它们。因为他们的文化与您的组织不同，所以简单地模仿他们是行不通的。
- **缺乏衡量：** DevOps 和 DevSecOps 转型必须根据既定目标进行衡量。度量标准可能包括软件交付性能或随时间推移的整体组织性能。
- **清单驱动的安全性：** 通过使用清单，安全团队遵循相同的陈旧、静态和不灵活的流程，这些流程既无用也不适用于开发人员用来使软件交付精益和敏捷的现代技术。 “[as code](https://www.oreilly.com/library/view/devopssec/9781491971413/ch04.html)”方法的引入需要安全人员学习如何编码。
- **安全作为一个特殊的团队：** 在从旧的软件交付方式（其中安全是一个单独的实体）过渡到 DevOps 的组织中尤其如此。由于分离，开发人员、运维人员和安全人员之间的信任存在问题。这将导致安全团队花费不必要的时间来审查和管理 DevOps 流程以及构建管道，而不是与开发人员和运营团队密切合作来改进软件交付流程。

## How to adopt DevSecOps successfully

## 如何成功采用 DevSecOps

Adopting DevSecOps is not easy, but being aware of common obstacles and challenges is key to your success.

采用 DevSecOps 并不容易，但了解常见的障碍和挑战是您成功的关键。

Clearly, the biggest and most important change an organization needs to make is its culture. Cultural change usually requires executive buy-in, as a top-down approach is necessary to convince people to make a successful turnaround. You might hope that executive buy-in makes cultural change follow naturally, but don't expect smooth sailing—executive buy-in alone is not enough.

显然，一个组织需要做出的最大和最重要的改变是它的文化。文化变革通常需要高管的支持，因为自上而下的方法对于说服人们成功扭转局面是必要的。您可能希望高管的认同让文化变革自然而然地发生，但不要指望一帆风顺——仅靠高管的认同是不够的。

To help accelerate cultural change, the organization needs leaders and enthusiasts that will become agents of change. Embed these people in the dev, ops, and security teams to serve as advocates and champions for culture change. This will also establish a cross-functional team that will share successes and learnings with other teams to encourage wider adoption.

为了帮助加速文化变革，组织需要能够成为变革推动者的领导者和狂热者。将这些人纳入开发、运营和安全团队，作为文化变革的倡导者和拥护者。这还将建立一个跨职能团队，与其他团队分享成功和经验，以鼓励更广泛的采用。

Once that is underway, the organization needs a DevSecOps use-case to start with, something small with a high potential for success. This enables the team to learn, fail, and succeed without affecting the organization's core business.

一旦开始，组织需要一个 DevSecOps 用例作为开始，一些小而有很大成功潜力的用例。这使团队能够在不影响组织核心业务的情况下学习、失败和成功。

The next step is to identify and agree on the definition of success. The DevSecOps adoption needs to be measurable; to do that, you need a dashboard that shows metrics such as:

下一步是确定并同意成功的定义。 DevSecOps 的采用需要是可衡量的；为此，您需要一个显示指标的仪表板，例如：

- Lead time for a change
- Deployment frequency
- Mean time to restore
- Change failure 

- 改变的准备时间
- 部署频率
- 平均恢复时间
- 更改失败

These metrics are a critical requirement to be able to identify processes and other things that require improvement. It's also a tool to declare if an adoption is a win or a bust. This methodology is called [event-driven transformation](https://www.openshift.com/blog/exploring-a-metrics-driven-approach-to-transformation).

这些指标是识别流程和其他需要改进的事情的关键要求。它也是一种宣布采用是成功还是失败的工具。这种方法称为[事件驱动转换](https://www.openshift.com/blog/exploring-a-metrics-driven-approach-to-transformation)。

## Conclusion

##  结论

When implemented properly, DevOps enables an organization to deliver software to production quickly and gain advantages over competitors. DevOps allows it to fail small and recover faster by enabling flexibility and efficiency to go to market early.

如果实施得当，DevOps 使组织能够快速将软件交付到生产环境并获得优于竞争对手的优势。 DevOps 通过实现灵活性和效率及早进入市场，使其能够以较小的故障发生并更快地恢复。

In summary, DevOps and DevSecOps adoption needs:

总之，DevOps 和 DevSecOps 的采用需要：

- Cultural change
- Executive buy-in
- Leaders and enthusiasts to act as evangelists
- Cross-functional teams
- Measurable indicators

- 文化变迁
- 行政买进
- 领袖和热心人士充当福音传道者
- 跨职能团队
- 可衡量的指标

Ultimately, the solution to the DevSecOps dilemma relies on cultural change to make the organization better. 

最终，DevSecOps 困境的解决方案依赖于文化变革以使组织变得更好。

