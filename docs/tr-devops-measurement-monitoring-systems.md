# DevOps measurement: Monitoring systems to inform business decisions

# DevOps 测量：监控系统以通知业务决策

https://cloud.google.com/architecture/devops/devops-measurement-monitoring-systems

https://cloud.google.com/architecture/devops/devops-measurement-monitoring-systems

**Note:** *Monitoring systems to inform business decisions* is one of a set of capabilities that drive higher software delivery and organizational performance. These capabilities were discovered by the [DORA State of DevOps research program](https://www.devops-research.com/research.html), an independent, academically rigorous investigation into the practices and capabilities that drive high performance. To learn more, read our [DevOps resources](https://cloud.google.com/devops).

**注意：** *监控系统以通知业务决策*是推动更高软件交付和组织绩效的一组功能之一。这些能力是由 [DORA State of DevOps 研究计划](https://www.devops-research.com/research.html) 发现的，该计划是对推动高性能的实践和能力进行独立的、学术上严谨的调查。要了解更多信息，请阅读我们的 [DevOps 资源](https://cloud.google.com/devops)。

Monitoring is the process of collecting, analyzing, and using information to track applications and infrastructure in order to guide business decisions. Monitoring is a key capability because it gives you insight into your systems and your work. Properly implemented, monitoring also gives you rapid feedback so that you can quickly find and fix problems early in the software development lifecycle.

监控是收集、分析和使用信息以跟踪应用程序和基础设施以指导业务决策的过程。监控是一项关键功能，因为它可以让您深入了解您的系统和工作。如果实施得当，监控还可以为您提供快速反馈，以便您可以在软件开发生命周期的早期快速发现并修复问题。

Monitoring also helps you communicate information about your systems to people in other areas of the software development and delivery pipeline, and to other parts of the business. Knowledge acquired downstream in operations might get integrated into upstream teams, such as development and product management. For example, the knowledge gained from operating a highly scalable application that uses a NoSQL database as a data store can be valuable information for developers as they build a similar application.

监控还可以帮助您将有关系统的信息传达给软件开发和交付管道的其他领域的人员以及业务的其他部分。在运营下游获得的知识可能会整合到上游团队中，例如开发和产品管理。例如，从操作使用 NoSQL 数据库作为数据存储的高度可扩展应用程序中获得的知识对于开发人员构建类似应用程序时可能是有价值的信息。

This knowledge transfer allows teams to quickly identify learnings, whether they stem from a production issue, a deployment error, or your customer usage patterns. You can then share these learnings across your organization to help people and systems improve.

这种知识转移使团队能够快速识别学习内容，无论它们来自生产问题、部署错误还是您的客户使用模式。然后，您可以在整个组织内分享这些学习成果，以帮助人员和系统改进。

## How to implement monitoring

## 如何实现监控

The following elements are key to effective monitoring:

以下要素是有效监测的关键：

- Collecting data from key areas throughout the value chain, including application performance and infrastructure.
- Using the collected data to make business decisions.

- 从整个价值链的关键领域收集数据，包括应用程序性能和基础设施。
- 使用收集的数据做出业务决策。

### Collecting data

###  收集数据

To collect data more effectively, you should implement monitoring solutions, either as homegrown services or managed services, that give visibility into development work, testing, QA, and IT operations. Make sure that you choose metrics that are appropriate for function and for your business.

为了更有效地收集数据，您应该实施监控解决方案，无论是作为本地服务还是托管服务，都可以让您了解开发工作、测试、QA 和 IT 运营。确保选择适合功能和业务的指标。

### Using data to make business decisions

### 使用数据做出业务决策

When you transform and visualize the collected data, you make it accessible to different audiences and help them make decisions. For example, you might want to share operations data upstream.You can also integrate this data as appropriate into reports and briefings, and use it in meetings to make informed business decisions. In this case, *appropriate* means relevant, timely, accurate, and easy to understand.

当您转换和可视化收集的数据时，您可以让不同的受众访问并帮助他们做出决策。例如，您可能希望在上游共享运营数据。您还可以根据需要将此数据集成到报告和简报中，并在会议中使用它来做出明智的业务决策。在这种情况下，*适当*意味着相关、及时、准确和易于理解。

In these meetings, be sure to also provide context, to help those who might not be familiar with the data understand how it pertains to the discussion and how it can inform the decisions to be made. For example, you might want to know how to answer the following questions:

在这些会议中，一定要提供背景信息，以帮助那些可能不熟悉数据的人了解它与讨论的关系以及它如何为要做出的决定提供信息。例如，您可能想知道如何回答以下问题：

- Are these values relatively high or low?
- Are they expected?
- Do you anticipate changes?
- How is this data different from historical reports?
- Has your technology or infrastructure impacted the numbers in interesting or non-obvious ways?

- 这些值是相对高还是低？
- 他们是预期的吗？
- 你预计会发生变化吗？
- 这些数据与历史报告有何不同？
- 您的技术或基础设施是否以有趣或不明显的方式影响了这些数字？

## Common pitfalls in monitoring

## 监控中的常见陷阱

The following pitfalls are common when monitoring systems:

在监控系统时，以下陷阱很常见：

- **Monitoring reactively**. For example, only getting alerted when the system goes down, but not using monitoring data to help alert when the system approaches critical thresholds.
- **Monitoring too small a scope**. For example, monitoring one or two areas rather than the full software development and delivery pipeline. This pitfall highlights metrics, focusing only on the areas that are measured, which might not be the optimal areas to monitor.
- **Focusing on local optimizations**. For example, focusing on reducing the response time for one service's storage needs without evaluating whether the broader infrastructure could also benefit from the same improvement. 

- **被动监控**。例如，仅在系统出现故障时发出警报，而不是在系统接近临界阈值时使用监控数据来帮助发出警报。
- **监控范围太小**。例如，监控一两个领域，而不是完整的软件开发和交付管道。这个陷阱突出了指标，只关注被测量的区域，这可能不是监控的最佳区域。
- **专注于局部优化**。例如，专注于减少一项服务的存储需求的响应时间，而不评估更广泛的基础设施是否也可以从相同的改进中受益。

- **Monitoring everything**. By collecting data and reporting on everything your system, you run the risk of over-alerting or drowning in data. Taking a thoughtful approach to monitoring can help draw attention to key areas.

- **监控一切**。通过收集数据并报告系统的所有内容，您可能会面临过度警觉或淹没在数据中的风险。采取深思熟虑的方法进行监测有助于引起对关键领域的关注。

## Ways to improve monitoring

## 改进监控的方法

To improve your monitoring effectiveness, we recommend that you focus your efforts on two main areas:

为了提高您的监控效率，我们建议您将精力集中在两个主要方面：

1. **Collecting data from key areas throughout the value chain.**

    1. **从整个价值链的关键领域收集数据。**

   By analyzing the data that you collect and doing a gap analysis, you can help ensure that you collect the right data for your organization.

通过分析您收集的数据并进行差距分析，您可以帮助确保为您的组织收集正确的数据。

2. **Using the collected data to make business decisions.**

    2. **使用收集的数据做出业务决策。**

   The data that you collect should drive value across the organization, and the metrics that you select must be meaningful to your organization. Meaningful data can be used by many teams, from DevOps to Finance.

    您收集的数据应该推动整个组织的价值，并且您选择的指标必须对您的组织有意义。从 DevOps 到财务的许多团队都可以使用有意义的数据。

   It's also important to find the right medium to display the monitoring information. Different uses for the information demand different presentation choices. Real-time dashboards might be most useful to the DevOps team, while regularly generated business reports might be useful for metrics measured over a longer period.

    找到合适的媒体来显示监控信息也很重要。信息的不同用途需要不同的呈现选择。实时仪表板可能对 DevOps 团队最有用，而定期生成的业务报告可能对长期测量的指标有用。

   The most important thing is to ensure the data is available, shared, and used to guide decisions. If the best you can do to kick things off is a shared spreadsheet, use that. Then graduate to fancy dashboards later. Don't let perfect be the enemy of good enough.

最重要的是确保数据可用、共享并用于指导决策。如果您能做的最好的事情是共享电子表格，请使用它。然后毕业到花哨的仪表板。不要让完美成为足够好的敌人。

## Ways to measure monitoring

## 衡量监控的方法

Effective monitoring helps drive performance improvements in software development and delivery. However, measuring the effectiveness of monitoring can be difficult to instrument in systems. Although you might be able to automatically measure how much data is being collected from your systems and the types of that data, it's more difficult to know if or where that data is being used.

有效的监控有助于推动软件开发和交付的性能改进。然而，在系统中测量监控的有效性可能很困难。尽管您可能能够自动测量从您的系统中收集了多少数据以及该数据的类型，但要知道这些数据是否被使用或在何处使用会更加困难。

To help you gauge the effectiveness of monitoring in your organization, consider the extent to which people agree or disagree with the following statements:

为了帮助您衡量组织中监控的有效性，请考虑人们同意或不同意以下陈述的程度：

- Data from application performance monitoring tools is used to make business decisions.
- Data from infrastructure monitoring tools is used to make business decisions. 

- 来自应用程序性能监控工具的数据用于制定业务决策。
- 来自基础设施监控工具的数据用于制定业务决策。

