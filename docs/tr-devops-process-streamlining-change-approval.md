# DevOps process: Streamlining change approval

# DevOps 流程：简化变更审批

https://cloud.google.com/architecture/devops/devops-process-streamlining-change-approval

**Note:** *Streamlining change approval* is one of a set of capabilities that drive higher software delivery and organizational performance. These capabilities were discovered by the [DORA State of DevOps research program](https://www.devops-research.com/research.html), an independent, academically rigorous investigation into the practices and capabilities that drive high performance. To learn more, read our [DevOps resources](https://cloud.google.com/devops).

**注意：** *简化变更审批*是推动更高软件交付和组织绩效的一组功能之一。这些能力是由 [DORA State of DevOps 研究计划](https://www.devops-research.com/research.html) 发现的，该计划是对推动高性能的实践和能力进行独立的、学术上严谨的调查。要了解更多信息，请阅读我们的 [DevOps 资源](https://cloud.google.com/devops)。

Most IT organizations have change management processes to manage the life cycle of changes to IT services, both internal and customer-facing. These processes are often the primary controls to reduce the operational and security risks of change.

大多数 IT 组织都有变更管理流程来管理内部和面向客户的 IT 服务变更的生命周期。这些流程通常是降低变更带来的运营和安全风险的主要控制措施。

Change management processes often include approvals by external reviewers or change approval boards (CABs) to promote changes through the system.

变更管理流程通常包括外部审核人员或变更批准委员会 (CAB) 的批准，以通过系统促进变更。

Compliance managers and security managers rely on change management processes to validate compliance requirements, which typically require evidence that all changes are appropriately authorized.

合规经理和安全经理依靠变更管理流程来验证合规要求，这通常需要证明所有变更都得到了适当的授权。

Research by DevOps Research and Assessment (DORA), presented in the [2019 State of DevOps Report (PDF)](https://cloud.google.com/devops/state-of-devops), finds that change approvals are best implemented through peer review during the development process, supplemented by automation to detect, prevent, and correct bad changes early in the software delivery life cycle. Techniques such as [continuous testing](https://cloud.google.com/architecture/devops/devops-tech-test-automation), [continuous integration](https://cloud.google.com/architecture/devops/devops-tech-continuous-integration), and [comprehensive monitoring and observability](https://cloud.google.com/architecture/devops/devops-measurement-monitoring-and-observability) provide early and automated detection, visibility, and fast feedback.

DevOps Research and Assessment (DORA) 在 [2019 年 DevOps 状态报告 (PDF)](https://cloud.google.com/devops/state-of-devops) 中进行的研究发现，变更批准最好实施通过开发过程中的同行评审，辅以自动化，以在软件交付生命周期的早期检测、预防和纠正不良变化。 [持续测试](https://cloud.google.com/architecture/devops/devops-tech-test-automation)、[持续集成](https://cloud.google.com/architecture/devops/)等技术devops-tech-continuous-integration) 和 [全面监控和可观察性](https://cloud.google.com/architecture/devops/devops-measurement-monitoring-and-observability) 提供早期和自动检测、可见性和快速反馈。

Further, organizations can improve their performance by doing a better job of communicating the existing process and helping teams navigate it efficiently. When team members have a clear understanding of the change approval process, this drives higher performance.

此外，组织可以通过更好地传达现有流程并帮助团队有效地导航来提高他们的绩效。当团队成员清楚地了解变更批准流程时，这将推动更高的绩效。

## How to implement a change approval process

## 如何实施变更审批流程

Two important goals of the change approval process are decreasing the risk of making changes, and satisfying regulatory requirements. One common regulatory requirement is segregation of duties, which states that changes must be approved by someone other than the author, thus ensuring that no individual has end-to-end control over a process.

变更审批流程的两个重要目标是降低进行变更的风险并满足监管要求。一项常见的监管要求是职责分离，其中规定更改必须由作者以外的其他人批准，从而确保没有个人对流程进行端到端控制。

Traditionally, these goals have been met through a heavyweight process involving approval by people external to the team proposing the change: a change advisory board (CAB) or a senior manager. However, DORA's research shows that these approaches have a negative impact on software delivery performance. Further, no evidence was found to support the hypothesis that a more formal, external review process was associated with lower change fail rates.

传统上，这些目标是通过一个重量级的流程来实现的，该流程涉及提出变更的团队外部人员的批准：变更顾问委员会 (CAB) 或高级经理。然而，DORA 的研究表明，这些方法对软件交付性能有负面影响。此外，没有发现证据支持更正式的外部审查过程与较低的变更失败率相关的假设。

Such heavyweight approaches tend to slow down the delivery process leading to the release of larger batches less frequently, with an accompanying higher impact on the production system that is likely to be associated with higher levels of risk and thus higher change fail rates. DORA's research found this hypothesis was supported in the data.

这种重量级的方法往往会减慢交付过程，从而导致大批量的发布频率降低，同时对生产系统产生更大的影响，这可能与更高的风险水平相关联，从而导致更高的变更失败率。 DORA 的研究发现，这一假设在数据中得到了支持。

Instead, teams should:

相反，团队应该：

- Use peer review to meet the goal of segregation of duties, with reviews, comments, and approvals captured in the team's [development platform](https://cloud.google.com/architecture/devops/devops-tech-code-maintainability) as part of the development process. 

- 使用同行评审来实现职责分离的目标，在团队的 [开发平台](https://cloud.google.com/architecture/devops/devops-tech-code-maintainability)中获取评审、评论和批准) 作为开发过程的一部分。

- Employ [continuous testing](https://cloud.google.com/architecture/devops/devops-tech-test-automation), [continuous integration](https://cloud.google.com/architecture/devops/devops-tech-continuous-integration), and [comprehensive monitoring and observability](https://cloud.google.com/architecture/devops/devops-measurement-monitoring-and-observability) to rapidly detect, prevent, and correct bad changes .
- Treat your development platform as a product that makes it easy for developers to get fast feedback on the impact of their changes on multiple axes, including security, performance, and stability, as well as defects.

- 采用[持续测试](https://cloud.google.com/architecture/devops/devops-tech-test-automation)、[持续集成](https://cloud.google.com/architecture/devops/devops-tech-continuous-integration) 和 [全面监控和可观察性](https://cloud.google.com/architecture/devops/devops-measurement-monitoring-and-observability) 以快速检测、预防和纠正不良变化.
- 将您的开发平台视为一种产品，使开发人员可以轻松获得有关其更改对多个轴的影响的快速反馈，包括安全性、性能和稳定性以及缺陷。

Your goal should be to make your regular change management process fast and reliable enough that you can use it to make emergency changes too.

您的目标应该是使您的常规变更管理流程足够快速和可靠，以便您也可以使用它来进行紧急变更。

In the [continuous delivery](https://cloud.google.com/architecture/devops/devops-tech-continuous-delivery) paradigm the CAB still has a vital role, which includes:

在 [持续交付](https://cloud.google.com/architecture/devops/devops-tech-continuous-delivery) 范式中，CAB 仍然发挥着至关重要的作用，其中包括：

- Facilitating notification and coordination between teams.
- Helping teams with process improvement work to increase their software delivery performance.
- Weighing in on important business decisions that require a trade-off and sign-off at higher levels of the business, such as the decision between time-to-market and business risk.

- 促进团队之间的通知和协调。
- 帮助团队进行流程改进工作，以提高他们的软件交付绩效。
- 权衡需要在更高业务层面进行权衡和批准的重要业务决策，例如上市时间和业务风险之间的决策。

This new role for the CAB is strategic. By shifting detailed code review to practitioners and automated methods, the time and attention of those in leadership and management positions is freed up to focus on more strategic work. This transition, from gatekeeper to process architect and information beacon, is consistent with the practices of organizations that excel at software delivery performance.

CAB 的这一新角色具有战略意义。通过将详细的代码审查转移到从业人员和自动化方法，领导和管理职位的人员的时间和注意力被腾出来，专注于更具战略性的工作。这种从看门人到流程架构师和信息信标的转变与擅长软件交付性能的组织的实践是一致的。

## Common pitfalls in change approval processes

## 变更审批流程中的常见陷阱

The following pitfalls are common to change approval processes:

以下陷阱在变更审批流程中很常见：

**Reliance on a centralized Change Approval Board (CAB)** to catch errors and approve changes. This approach can introduce delay and often error. CABs are good at broadcasting change, but people that far removed from the change might not understand the implications of those changes.

**依靠集中的变更批准委员会 (CAB)** 来捕捉错误并批准变更。这种方法可能会引入延迟并且经常会出错。 CAB 擅长传播变化，但远离变化的人可能不了解这些变化的含义。

**Treating all changes equally.** When all changes are subject to the same approval process, change review is inefficient, and people are unable to devote time and attention to those that require true concentration because of differences in risk profile or timing.

**平等对待所有变更。** 当所有变更都经过相同的审批流程时，变更审查效率低下，并且由于风险状况或时间的差异，人们无法将时间和注意力投入到需要真正集中注意力的事情上。

**Failing to apply continuous improvement.** As with all processes, key performance metrics such as lead time and change fail rate should be targeted with the goal of improving the performance of the change management process, including providing teams with tools and training to help them navigate it more effectively.

**未能应用持续改进。** 与所有流程一样，关键绩效指标（例如提前期和变更失败率）应以提高变更管理流程绩效为目标，包括为团队提供工具和培训帮助他们更有效地导航。

**Responding to problems by adding more process.** Often organizations use additional process and more heavyweight approvals when faced with stability problems in production. Analysis suggests this approach will make things worse because this drives up lead times and batch sizes, creating a vicious cycle. Instead, invest in making it quicker and safer to make changes.

**通过添加更多流程来应对问题。** 当面临生产中的稳定性问题时，组织通常会使用额外的流程和更重量级的审批。分析表明，这种方法会使事情变得更糟，因为这会增加交货时间和批量大小，从而形成恶性循环。相反，投资于使更改更快、更安全。

## Ways to improve your change approval process

## 改进变更审批流程的方法

To improve your change approval processes, focus on implementing the following:

要改进您的变更批准流程，请重点实施以下内容：

1. Moving to a peer-review based process for individual changes, enforced at code check-in time, and supported by automated tests.
2. Finding ways to discover problems such as regressions, performance problems, and security issues in an automated fashion as soon as possible after changes are committed.
3. Performing ongoing analysis to detect and flag high risk changes early on so that they can be subjected to additional scrutiny.
4. Looking at the change process end-to-end, identifying bottlenecks, and experimenting with ways to shift validations into the development platform.
5. Implementing information security controls at the platform and infrastructure layer and in the development tool chain, rather than reviewing them manually as part of the software delivery process. 



1. 转向基于同行评审的个人更改流程，在代码签入时强制执行，并由自动化测试提供支持。
2. 想办法在提交变更后尽快以自动化的方式发现回归、性能问题和安全问题等问题。
3. 进行持续分析，尽早发现和标记高风险变化，以便对其进行额外审查。
4. 查看端到端的变更过程，识别瓶颈，并尝试将验证转移到开发平台的方法。
5. 在平台和基础设施层以及开发工具链中实施信息安全控制，而不是作为软件交付过程的一部分手动审查它们。



Research from the [2019 State of DevOps Report (PDF)](https://cloud.google.com/devops/state-of-devops)  shows that while moving away from traditional, formal change management processes is the ultimate goal, simply doing a better job of communicating the existing process and helping teams navigate it efficiently has a positive impact on software delivery performance. When team members have a clear understanding of the process to get changes approved for implementation, this drives high performance. This means they are confident that they can get changes through the approval process in a timely manner and know the steps it takes to go from "submitted" to "accepted" every time for all the types of changes they typically make.

[2019 年 DevOps 状态报告 (PDF)](https://cloud.google.com/devops/state-of-devops) 的研究表明，虽然摆脱传统的正式变更管理流程是最终目标，但更好地沟通现有流程并帮助团队有效地导航它对软件交付性能产生积极影响。当团队成员清楚地了解获得批准实施变更的流程时，这将推动高性能。这意味着他们有信心可以及时通过审批流程进行更改，并且知道每次他们通常进行的所有类型的更改从“提交”到“接受”所需的步骤。

## Ways to measure change approval in your systems

## 衡量系统中变更批准的方法

Now your teams can list possible ways to measure change approval:

现在，您的团队可以列出衡量变更批准的可能方法：

| Factor to test                                               | What to measure                                              |
| ------------------------------------------------------------ |------------------------------------------------------------ |
| Can changes be promoted to production without manual change approvals? | The percentage of changes that do (or do not) require a manual change to be promoted to production. Tip: You can also measure this factor based on risk profile: what percentage of low-, medium-, and high-risk changes require a manual change to be promoted to production? |
| Do production changes need to be approved by an external body before deployment or implementation? | The amount of time changes spend waiting for approval from external bodies. Tip: As you shift approvals closer to the work, measure the amount of time spent waiting for approval from local approval bodies or reviewers. You can also measure this factor by risk profile. Measure number or proportion of changes that require approval from external bodies, as well as the time spent waiting for those approvals. |
| Do you rely on peer review to manage changes? | Percentage of changes that are managed by peer-review. You can also measure this factor by risk profile. |
| Do team members have a clear understanding of the process to get changes approved for implementation? | The extent to which team members are confident they can get changes  through the approval process in a timely manner and know the steps it  takes to go from "submitted" to "accepted" every time for all the types  of changes they typically make. |

While you consider your own environment, you will likely develop your own measures to understand and gain insight into your change approval processes. We suggest you use these to not only measure your process but also work to improve it.

当您考虑自己的环境时，您可能会制定自己的措施来了解和深入了解您的变更批准流程。我们建议您不仅使用这些来衡量您的过程，还可以努力改进它。

