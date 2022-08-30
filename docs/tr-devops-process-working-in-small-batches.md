# DevOps process: Working in small batches

# DevOps 流程：小批量工作

https://cloud.google.com/architecture/devops/devops-process-working-in-small-batches

https://cloud.google.com/architecture/devops/devops-process-working-in-small-batches

**Note:** *Working in small batches* is one of a set of capabilities that drive higher software delivery and organizational performance. These capabilities were discovered by the [DORA State of DevOps research program](https://www.devops-research.com/research.html), an independent, academically rigorous investigation into the practices and capabilities that drive high performance. To learn more, read our [DevOps resources](https://cloud.google.com/devops).

**注意：** *小批量工作*是推动更高软件交付和组织绩效的一组功能之一。这些能力是由 [DORA State of DevOps 研究计划](https://www.devops-research.com/research.html) 发现的，该计划是对推动高性能的实践和能力进行独立的、学术上严谨的调查。要了解更多信息，请阅读我们的 [DevOps 资源](https://cloud.google.com/devops)。

Working in small batches is an essential principle in any discipline where feedback loops are important, or you want to learn quickly from your decisions. Working in small batches allows you to rapidly test hypotheses about whether a particular improvement is likely to have the effect you want, and if not, lets you course correct or revisit assumptions. Although this article applies to any type of change that includes organizational transformation and process improvement, it focuses primarily on software delivery.

在反馈循环很重要或您希望从决策中快速学习的任何学科中，小批量工作都是一项基本原则。小批量工作使您可以快速测试有关特定改进是否可能产生您想要的效果的假设，如果没有，您可以纠正或重新审视假设。尽管本文适用于包括组织转型和流程改进在内的任何类型的变更，但它主要关注软件交付。

Working in small batches is part of lean product management. Together with capabilities like [visibility of work in the value stream](https://cloud.google.com/architecture/devops/devops-process-work-visibility-in-value-stream), [team experimentation](https://cloud.google.com/architecture/devops/devops-process-team-experimentation), and [visibility into customer feedback](https://cloud.google.com/architecture/devops/devops-process-customer-feedback), working in small batches predicts software delivery performance and organizational performance.

小批量工作是精益产品管理的一部分。连同 [价值流中工作的可见性](https://cloud.google.com/architecture/devops/devops-process-work-visibility-in-value-stream)、[团队实验](https://cloud.google.com/architecture/devops/devops-process-team-experimentation)和[客户反馈的可见性](https://cloud.google.com/architecture/devops/devops-process-customer-feedback)，小批量工作可以预测软件交付绩效和组织绩效。

One reason work is done in large batches is because of [the large fixed cost of handing off changes](https://cloud.google.com/architecture/devops/devops-process-streamlining-change-approval). In traditional phased approaches to software development, handoffs from development to test or from test to IT operations consist of whole releases: months worth of work by teams consisting of tens or hundreds of people. With this traditional approach, collecting feedback on a change can take weeks or months.

大批量完成工作的一个原因是[交付更改的固定成本很高](https://cloud.google.com/architecture/devops/devops-process-streamlining-change-approval)。在传统的软件开发分阶段方法中，从开发到测试或从测试到 IT 运营的交接包括整个版本：由数十或数百人组成的团队花费数月的时间。使用这种传统方法，收集有关更改的反馈可能需要数周或数月的时间。

In contrast, DevOps practices, which utilize cross-functional teams and lightweight approaches, allow for software to progress from development through test and operations into production in a matter of minutes. However, this rapid progression requires working with code in small batches.

相比之下，利用跨职能团队和轻量级方法的 DevOps 实践允许软件在几分钟内从开发通过测试和运营进入生产。但是，这种快速发展需要小批量处理代码。

Working in small batches has many benefits:

小批量工作有很多好处：

- It reduces the time it takes to get feedback on changes, making it easier to triage and remediate problems.
- It increases efficiency and motivation.
- It prevents your organization from succumbing to the sunk-cost fallacy.

- 它减少了获得更改反馈所需的时间，从而更容易对问题进行分类和补救。
- 它提高了效率和动力。
- 它可以防止您的组织屈服于沉没成本谬误。

You can apply the small batches approach at the feature and the product level. As an illustration, a minimum viable product, or MVP, is a prototype of a product with just enough features to enable validated learning about the product and its business model.

您可以在功能和产品级别应用小批量方法。例如，最小可行产品或 MVP 是产品的原型，它具有足够的功能来支持对产品及其业务模型的经过验证的学习。

Continuous delivery builds upon working in small batches and tries to get every change in version control as early as possible. A goal of continuous delivery is to change the economics of the software delivery process, making it viable to work in small batches. This approach provides fast, comprehensive feedback to teams so that they can improve their work.

持续交付建立在小批量工作的基础上，并试图尽早获得版本控制中的每一个变化。持续交付的一个目标是改变软件交付过程的经济性，使小批量工作变得可行。这种方法为团队提供快速、全面的反馈，以便他们改进工作。

## How to work in small batches

## 如何小批量工作

When you plan new features, try to break them down into work units that can be completed independently and in short timeframes. We recommend that each feature or batch of work follow the agile concept of the [INVEST principle](https://wikipedia.org/wiki/INVEST_(mnemonic)):

当您计划新功能时，请尝试将它们分解为可以在短时间内独立完成的工作单元。我们建议每个功能或每批工作都遵循 [INVEST 原则](https://wikipedia.org/wiki/INVEST_(mnemonic)) 的敏捷概念：

- **Independent**. Make batches of work as independent as possible from other batches, so that teams can work on them in any order, and deploy and validate them independent of other batches of work.
- **Negotiable**. Each batch of work is iterable and can be renegotiated as feedback is received.
- **Valuable**. Discrete batches of work are usable and provide value to the stakeholders.
- **Estimable**. Enough information exists about the batches of work that you can easily estimate the scope.
- **Small**. During a sprint, you should be able to complete batches of work in small increments of time, meaning hours to a couple days.
- **Testable**. Each batch of work can be tested, monitored, and verified as working in the way users expect. 

- **独立的**。使批次工作尽可能独立于其他批次，以便团队可以按任何顺序处理它们，并独立于其他批次工作进行部署和验证。
- **可协商**。每批工作都是可迭代的，并且可以在收到反馈时重新协商。
- **有价值的**。离散批次的工作是可用的，并为利益相关者提供价值。
- **估计**。关于工作批次的足够信息，您可以轻松估计范围。
- **小的**。在 sprint 期间，您应该能够以较小的时间增量完成批量工作，这意味着几小时到几天。
- **可测试**。每一批工作都可以按照用户期望的方式进行测试、监控和验证。

When features are of an appropriate size, you can split the development of the feature into even smaller batches. This process can be difficult and requires experience to develop. Ideally, your developers should be checking multiple small releasable changes [into trunk at least once per day](https://cloud.google.com/architecture/devops/devops-tech-trunk-based-development).

当功能具有适当的大小时，您可以将功能的开发分成更小的批次。这个过程可能很困难，需要经验来开发。理想情况下，您的开发人员应该检查多个小的可发布更改 [到主干中至少每天一次](https://cloud.google.com/architecture/devops/devops-tech-trunk-based-development)。

The key is to start development at the service or API layer, not at the UI layer. In this way, you can make additions to the API that won't initially be available to users of the app, and check those changes into trunk. You can launch these changes to production without making them visible to users. This approach, called *dark launching*, allows developers to check in code for small batches that have been completed, but for features that are not yet fully complete. You can then run [automated tests](https://cloud.google.com/architecture/devops/devops-tech-test-automation)  against these changes to prove that they behave in the expected way. This way, teams are still working quickly and [developing off of trunk](https://cloud.google.com/architecture/devops/devops-tech-trunk-based-development)  and not long-lived feature branches.

关键是从服务或 API 层开始开发，而不是在 UI 层。通过这种方式，您可以添加最初对应用程序用户不可用的 API，并将这些更改签入到主干中。您可以在不让用户看到的情况下将这些更改启动到生产环境中。这种称为*暗启动*的方法允许开发人员签入已完成但尚未完全完成的功能的小批量代码。然后，您可以针对这些更改运行 [自动化测试](https://cloud.google.com/architecture/devops/devops-tech-test-automation)，以证明它们的行为符合预期。这样，团队仍然可以快速工作并且[开发主干](https://cloud.google.com/architecture/devops/devops-tech-trunk-based-development)，而不是长期存在的功能分支。

You can also dark launch changes by using a [feature toggle](https://martinfowler.com/bliki/FeatureToggle.html), which is a conditional statement based on configuration settings. For example, you can make UI elements visible or invisible, or you can enable or disable service logic. Feature-toggle configuration might be read either at deploy time or runtime. You can use these configuration settings to switch the behavior of new code further down the stack. You can also use similar technique known as [branch by abstraction](https://continuousdelivery.com/2011/05/make-large-scale-changes-incrementally-with-branch-by-abstraction/)  to make larger-scale changes to the system while continuing to develop and release off-trunk without the use of long-lived feature branches.

您还可以使用 [功能切换] (https://martinfowler.com/bliki/FeatureToggle.html) 隐藏启动更改，这是基于配置设置的条件语句。例如，您可以使 UI 元素可见或不可见，或者您可以启用或禁用服务逻辑。可以在部署时或运行时读取功能切换配置。您可以使用这些配置设置将新代码的行为进一步切换到堆栈中。您还可以使用称为 [branch by abstraction](https://continuousdelivery.com/2011/05/make-large-scale-changes-incrementally-with-branch-by-abstraction/) 的类似技术来进行更大规模在不使用长期存在的特性分支的情况下继续开发和发布非主干系统的同时更改系统。

In this approach, batches of work aren't complete until they're deployed to production and the feedback process has begun to validate the changes. Feedback comes from many sources and in many forms, including users, system monitoring, quality assurance, and automated tests. Your goal is to optimize for speed so that you reduce the cycle time to get changes into the hands of users. This way, you can validate your hypothesis as quickly as possible.

在这种方法中，成批的工作在部署到生产环境并且反馈过程已经开始验证更改之前是不完整的。反馈来自多种来源和多种形式，包括用户、系统监控、质量保证和自动化测试。您的目标是优化速度，以便缩短将更改交到用户手中的周期时间。这样，您可以尽快验证您的假设。

## Common pitfalls with working in small batches

## 小批量工作的常见陷阱

When you break down work into small batches, you encounter two pitfalls:

当您将工作分解为小批量时，您会遇到两个陷阱：

- **Not breaking up work into small enough pieces**. Your first task is to break down the work in a meaningful way. We recommend that you commit code independent of the status of the feature and that individual features take no more than a few days to develop. Any batch of code that takes longer than a week to complete and check is too big. Throughout the development process, it's essential that you analyze how to break down ideas into increments that you can develop iteratively.
- **Working in small batches but then regrouping the batches before sending them downstream for testing or release**. Regrouping work in this way delays the feedback on whether the changes have defects, and whether your users and your organization agree the changes were the right thing to build in the first place.

- **没有将工作分解成足够小的部分**。您的首要任务是以有意义的方式分解工作。我们建议您独立于功能状态提交代码，并且单个功能的开发时间不超过几天。任何一批需要超过一周才能完成和检查的代码都太大了。在整个开发过程中，您必须分析如何将想法分解为可以迭代开发的增量。
- **小批量工作，然后重新组合批次，然后将它们发送到下游进行测试或发布**。以这种方式重新组合工作会延迟有关更改是否存在缺陷以及您的用户和您的组织是否同意这些更改是首先构建的正确事物的反馈。

## Ways to reduce the size of work batches

## 减少工作批次大小的方法

When you slice work into small batches that can be completed in hours, you can typically [test and deploy those batches to production in less than an hour](https://services.google.com/fh/files/misc/state-of-devops-2016.pdf)  (PDF). The key is to decompose the work into small features that allow for rapid development, rather than developing complex features on branches and releasing them infrequently.

当您将工作分成可以在数小时内完成的小批量时，您通常可以 [在不到一小时的时间内测试并将这些批次部署到生产中](https://services.google.com/fh/files/misc/state-of-devops-2016.pdf) (PDF)。关键是将工作分解为允许快速开发的小功能，而不是在分支上开发复杂的功能并不经常发布。

To improve small batch development, check your environment and confirm that the following conditions are true:

要改进小批量开发，请检查您的环境并确认满足以下条件：

- Work is decomposed in a way that enables teams to make more frequent production releases.
- Developers are experienced in breaking down work into small changes that can be completed in the space of hours, not days. 

- 工作的分解方式使团队能够更频繁地发布产品。
- 开发人员在将工作分解为可以在数小时内完成的小更改方面经验丰富，而不是几天。

To become an expert in small batch development, strive to meet each of these conditions in all of your development teams. This practice is a necessary condition for both [continuous integration](https://cloud.google.com/architecture/devops/devops-tech-continuous-integration)  and [trunk-based development](https://cloud.google.com/architecture/devops/devops-tech-trunk-based-development).

要成为小批量开发方面的专家，请努力在您的所有开发团队中满足这些条件。这种做法是 [持续集成](https://cloud.google.com/architecture/devops/devops-tech-continuous-integration）和[基于主干的开发](https://cloud.google)的必要条件.com/architecture/devops/devops-tech-trunk-based-development)。

## Ways to measure the size of work batches

## 测量工作批次大小的方法

When you understand [continuous integration](https://cloud.google.com/architecture/devops/devops-tech-continuous-integration)  and [monitoring](https://cloud.google.com/architecture/devops/devops-measurement-monitoring-systems), you can outline possible ways to measure small batch development in your systems and development environment.

当您了解 [持续集成](https://cloud.google.com/architecture/devops/devops-tech-continuous-integration) 和 [监控](https://cloud.google.com/architecture/devops/devops-measurement-monitoring-systems)，您可以概述在您的系统和开发环境中测量小批量开发的可能方法。

- **Application features are decomposed in a way that supports frequent releases**. How often are releases possible? How does this release cadence differ across teams? Are delays in production related to features that are larger?
- **Application features are sliced in a way that lets developers complete the work in one week or less**. What proportion of features can you complete in one week or less? What features can't you complete in one week or less? Can you commit and release changes before the feature is complete?
- **MVPs are defined and set as goals for teams**. Is the work decomposed into features that allow for MVPs and rapid development, rather than complex and lengthy processes?

- **应用程序功能以支持频繁发布的方式进行分解**。多久可以发布一次？不同团队的发布节奏有何不同？生产延迟是否与更大的功能有关？
- **应用程序功能被划分为让开发人员在一周或更短的时间内完成工作**。您可以在一周或更短的时间内完成多少功能？您无法在一周或更短的时间内完成哪些功能？您可以在功能完成之前提交和发布更改吗？
- **MVP 被定义并设定为团队的目标**。工作是否分解为允许 MVP 和快速开发的功能，而不是复杂和冗长的过程？

Your measurements depend on the following:

您的测量取决于以下因素：

- Knowing your organization's processes.
- Setting goals for reducing waste.
- Looking for ways to reduce complexity in the development process. 

- 了解您组织的流程。
- 设定减少浪费的目标。
- 寻找降低开发过程复杂性的方法。

