# DevOps culture: How to transform

# DevOps 文化：如何转型

https://cloud.google.com/architecture/devops/devops-culture-transform

https://cloud.google.com/architecture/devops/devops-culture-transform

Every organization is constantly undergoing change. Therefore, some questions to ask are:

每个组织都在不断地发生变化。因此，要问的一些问题是：

- What is the direction of that change?
- What are the system-level outcomes you are working towards?
- Is the organization better able to discover and serve its customers, and thus achieve its purpose?
- Does the organization's business model and management of its people provide long-term sustainability?

- 这种变化的方向是什么？
- 您正在努力实现的系统级成果是什么？
- 组织是否能够更好地发现和服务客户，从而实现其目标？
- 组织的商业模式和人员管理是否提供长期可持续性？

When it seems as if things aren't going to plan, it's common for leaders to roll out a transformation program. However, these programs often fail to achieve their goals, using up large quantities of resources and organizational capacity. This document examines how to successfully execute a transformation and addresses some common sources of failure.

当事情似乎无法按计划进行时，领导者通常会推出转型计划。然而，这些计划往往无法实现其目标，消耗了大量资源和组织能力。本文档检查如何成功执行转换并解决一些常见的失败来源。

## How to implement transformation

## 如何实现转换

There are two key ingredients in effective, ongoing transformation: processes for executing organizational change by setting goals and enabling team experimentation, and mechanisms to spread good practice through the organization.

有效、持续的转型有两个关键要素：通过设定目标和支持团队实验来执行组织变革的过程，以及在组织中传播良好实践的机制。

### Set goals and enable team experimentation

### 设定目标并启用团队实验

There are many frameworks for executing and measuring organizational change, such as [balanced scorecard](https://wikipedia.org/wiki/Balanced_scorecard), [objectives and key results (OKRs)](https://rework.withgoogle.com/guides/set-goals-with-okrs/steps/introduction/), and the [improvement kata](http://www-personal.umich.edu/~mrother/The_Improvement_Kata.html)  and [coaching kata](http://www-personal.umich.edu/~mrother/The_Coaching_Kata.html). These frameworks might seem different, but they all share key features. The basic dynamic is shown in the following figure, which is based on the improvement kata framework:

有许多用于执行和衡量组织变革的框架，例如 [平衡计分卡](https://wikipedia.org/wiki/Balanced_scorecard)、[目标和关键结果(OKR)](https://rework.withgoogle.com/guides/set-goals-with-okrs/steps/introduction/)，以及 [改进型](http://www-personal.umich.edu/~mrother/The_Improvement_Kata.html) 和 [教练型](http://www-personal.umich.edu/~mrother/The_Coaching_Kata.html)。这些框架可能看起来不同，但它们都具有共同的关键特性。基本动态如下图所示，基于改进的kata框架：

![The 4-stage Kata model.](https://cloud.google.com/static/architecture/devops/images/devops-culture-transform-kata-model.png)

Source: Reproduced by permission of Mike Rother, from [*Toyota Kata Practice Guide: Practice Scientific Thinking Skills for Superior Results in 20 Minutes a Day*](https://www.mheducation.com/highered/product/toyota-kata-practice-guide-practicing-scientific-thinking-skills-superior-results-20-minutes-day-rother/9781259861024.html)  by Mike Rother (McGraw-Hill 2018).

资料来源：经 Mike Rother 许可转载，出自 [*Toyota Kata 实践指南：每天 20 分钟练习科学思维技能以获得卓越结果*](https://www.mheducation.com/highered/product/toyota-kata-实践指南实践科学思维技能高级结果20分钟天罗瑟/9781259861024.html)由迈克罗瑟（McGraw-Hill2018)。

All of the frameworks start with a direction (a "true north") at the organizational or division level. This is an aspirational, system-level business goal set by [leadership](https://cloud.google.com/architecture/devops/devops-culture-transformational-leadership). It could be an ideal that can't be achieved, such as zero injuries (the [goal](https://hbswk.hbs.edu/archive/paul-o-neill-values-into-action)  chosen by Alcoa's CEO Paul O'Neill). Or it could be a tough goal that is one to three years out, such as a tenfold increase in productivity (the [goal](https://continuousdelivery.com/evidence-case-studies/#the-hp-futuresmart-case-study)  chosen by Gary Gruver, when he was Director of Engineering of HP's LaserJet Firmware division).

所有框架都以组织或部门级别的方向（“真正的北方”）开始。这是由 [leadership](https://cloud.google.com/architecture/devops/devops-culture-transformional-leadership)设定的雄心勃勃的系统级业务目标。这可能是一个无法实现的理想，例如美铝首席执行官选择的零伤害（[目标](https://hbswk.hbs.edu/archive/paul-o-neill-values-into-action)保罗奥尼尔）。或者，这可能是一到三年后的艰难目标，例如将生产力提高十倍（[目标](https://continuousdelivery.com/evidence-case-studies/#the-hp-futuresmart-case -研究）由 Gary Gruver 选择，当时他是惠普 LaserJet 固件部门的工程总监)。

The next step is to understand the current condition. The [DORA quick check](https://devops-research.com/quickcheck.html)  can help you understand how you're doing in terms of your software development capabilities and outcomes. Another analysis approach is to perform exercises like [value stream mapping](https://www.google.com/books/edition/Value_Stream_Mapping_How_to_Visualize_Wo/MeFrAAAAQBAJ), or activity accounting. The point is to understand where the organization is in measurable terms. 

下一步是了解当前状况。 [DORA 快速检查](https://devops-research.com/quickcheck.html) 可以帮助您了解您在软件开发能力和成果方面的表现。另一种分析方法是执行 [价值流映射](https://www.google.com/books/edition/Value_Stream_Mapping_How_to_Visualize_Wo/MeFrAAAAQBAJ) 或活动会计等练习。关键是要以可衡量的方式了解组织的位置。

The third step is to set measurable targets for a future date. These targets could be described using a format like OKRs, which begin with a qualitative objective, and then specify measurable key results (target conditions). For example, HSBC's CIO for Global Banking and Markets [set every team the goal](https://www.linkedin.com/pulse/double-half-quarter-lesson-from-book-richard-david-knott/)  " to double, half and quarter every year: double the frequency of releases, half the number of low impact incidents, and quarter the number of high impact incidents."

第三步是为未来的日期设定可衡量的目标。这些目标可以使用 OKR 之类的格式来描述，它从定性目标开始，然后指定可衡量的关键结果（目标条件）。例如，汇丰银行全球银行和市场部的 CIO [为每个团队设定目标](https://www.linkedin.com/pulse/double-half-quarter-lesson-from-book-richard-david-knott/)每年翻倍、一半和季度：发布频率翻倍，低影响事件数量减半，高影响事件数量减半。”

Finally, teams experiment with ways to achieve these goals until the future date is reached, supported by management. Teams take a scientific approach to experimentation, using the [PDCA method](https://wikipedia.org/wiki/PDCA)  (Plan-Do-Check-Act), also known as the Deming cycle. The cycle consists of the following steps:

最后，团队在管理层的支持下尝试实现这些目标的方法，直到达到未来的日期。团队采用科学方法进行实验，使用 [PDCA 方法](https://wikipedia.org/wiki/PDCA)(Plan-Do-Check-Act)，也称为戴明循环。该循环由以下步骤组成：

- Plan: determine the expected outcome.
- Do: perform the experiment.
- Check: study the results.
- Act: decide what to do next.

- 计划：确定预期结果。
- 做：进行实验。
- 检查：研究结果。
- 行动：决定下一步做什么。

Teams should be running experiments on a daily basis to try to move towards the target conditions or key results. In the [improvement kata](http://www-personal.umich.edu/~mrother/The_Improvement_Kata.html), everybody on the team should ask themselves the following five questions every day:

团队应该每天进行实验，以尝试朝着目标条件或关键结果迈进。在[改进型](http://www-personal.umich.edu/~mrother/The_Improvement_Kata.html)中，团队中的每个人每天都应该问自己以下五个问题：

1. What is the target condition?
2. What is the current condition?
3. What obstacles do you think are preventing you from reaching the target condition? Which one are you addressing now?
4. What is your next step? What outcome do you expect?
5. When can the results be evaluated to see what can be learned from taking that step?

1. 目标条件是什么？
2. 目前情况如何？
3. 您认为哪些障碍阻碍您达到目标条件？你现在针对的是哪一个？
4. 你的下一步是什么？你期待什么结果？
5. 什么时候可以评估结果，看看可以从这一步中学到什么？

When the results have been captured and new targets are set, repeat the process.

捕获结果并设置新目标后，重复该过程。

Because the process is performed under conditions of uncertainty, it's not always clear how the results will be achieved. Therefore, progress is often nonlinear, as shown in the following diagram:

由于该过程是在不确定的条件下进行的，因此并不总是很清楚如何获得结果。因此，进度往往是非线性的，如下图所示：

![Graph showing nonlinear progress mapping performance against iterations;the performance line is spiky but overall goes up.](https://cloud.google.com/static/architecture/devops/images/devops-culture-transform-non-linear-progress.png)

性能线是尖峰，但整体上升。](https://cloud.google.com/static/architecture/devops/images/devops-culture-transform-non-linear-progress.png)

Source: CC-BY: [*Lean Enterprise: How High Performance Organizations Innovate at Scale*](https://www.google.com/books/edition/Lean_Enterprise/G_ixBQAAQBAJ)  by Jez Humble, Joanne Molesky, and Barry O' Reilly (O'Reilly, 2014).

资料来源：CC-BY：[*精益企业：高绩效组织如何大规模创新*](https://www.google.com/books/edition/Lean_Enterprise/G_ixBQAAQBAJ)，作者：Jez Humble、Joanne Molesky 和 Barry O'赖利（O'Reilly，2014 年)。

In the planning meetings, participants review the target conditions or key results that were set in the last planning meeting. They then set new goals for the next iteration. In review meetings, participants look at how well the teams are achieving the goals for the iteration and discuss any obstacles and how they will be addressed.

在计划会议中，参与者审查在上次计划会议中设定的目标条件或关键结果。然后，他们为下一次迭代设定了新目标。在审查会议中，参与者会查看团队实现迭代目标的情况，并讨论任何障碍以及如何解决这些障碍。

Some important points about this pattern are the following:

关于此模式的一些要点如下：

- A team's own target conditions or [OKRs](https://rework.withgoogle.com/guides/set-goals-with-okrs/steps/introduction/)  must be set by the team. If they are set in a top-down way, teams won't have a stake in the outcome and thus won't be as invested in achieving them. Instead, the team might "game" them—that is, manipulate the outcome to meet the goal artificially.
- It's acceptable to not achieve the goals; some of the goals are *stretch goals*, meaning that they're purposely designed to be challenging. Teams should expect to achieve about 80% of the goals. It's common when starting with cultural transformation to not achieve *any* of the specified goals. If this happens, the team needs to set a single goal for the next iteration and dedicate everything to achieving it.
- Many goals and measures will change from iteration to iteration as the team's goals and current conditions change, and as they learn through working towards their goals. Don't spend too much time trying to set the perfect objectives: focus on executing the process so you can start learning. 

- 团队自己的目标条件或 [OKRs](https://rework.withgoogle.com/guides/set-goals-with-okrs/steps/introduction/) 必须由团队设置。如果它们以自上而下的方式设置，团队将不会对结果产生影响，因此不会为实现它们而投入大量资金。相反，团队可能会“玩弄”他们——也就是说，人为地操纵结果以达到目标。
- 没有达到目标是可以接受的；其中一些目标是*延伸目标*，这意味着它们是故意设计成具有挑战性的。团队应该期望实现大约 80% 的目标。从文化转型开始时无法实现*任何*指定目标是很常见的。如果发生这种情况，团队需要为下一次迭代设定一个目标，并全力以赴实现它。
- 随着团队目标和当前条件的变化，以及他们通过朝着目标努力而学习，许多目标和措施将随着迭代而变化。不要花太多时间试图设定完美的目标：专注于执行过程，这样你就可以开始学习了。

- It's important for teams to have the necessary autonomy, capacity, resources, and management and [leadership support](https://cloud.google.com/architecture/devops/devops-culture-transformational-leadership)  to do improvement work. Teams should not let the normal delivery work crowd out improvement work, because the improvement work is what will help fix the inefficiencies that make it so slow to deliver products and services.

- 团队拥有必要的自主权、能力、资源和管理以及[领导支持](https://cloud.google.com/architecture/devops/devops-culture-transformional-leadership) 来进行改进工作非常重要。团队不应该让正常的交付工作挤占改进工作，因为改进工作将有助于解决导致交付产品和服务如此缓慢的低效率问题。

### Build community structures to spread knowledge

### 建立社区结构以传播知识

After teams have discovered better ways of working, the next task is to spread lessons learned throughout the organization. There are many ways to do this. In the [2019 State of DevOps Report](https://cloud.google.com/devops/state-of-devops)  researchers asked respondents to share how their teams and organizations spread DevOps and Agile methods by selecting from one or more of the following approaches (see Appendix B of the 2019 State of DevOps Report for detailed descriptions):

在团队发现更好的工作方式之后，下一个任务是在整个组织中传播经验教训。有很多方法可以做到这一点。在 [2019 年 DevOps 状况报告](https://cloud.google.com/devops/state-of-devops) 中，研究人员要求受访者分享他们的团队和组织如何通过选择一种或多种来传播 DevOps 和敏捷方法以下方法（有关详细说明，请参阅 2019 年 DevOps 状态报告的附录 B)：

- Training center (sometimes referred to as a *dojo*)
- Center of excellence (CoE)
- Proof of concept (PoC) but stall
- Proof of concept as a template
- Proof of concept as a seed
- Communities of practice
- Big bang
- Bottom-up or grassroots
- Mashup

- 培训中心（有时称为 *dojo*）
- 卓越中心 (CoE)
- 概念证明 (PoC) 但停滞不前
- 概念证明作为模板
- 作为种子的概念证明
- 实践社区
- 大爆炸
- 自下而上或基层
- 混搭

Analysis shows that high performers favor strategies that create community structures at both low and high levels in the organization, likely making them more sustainable and resilient to re-organizations and product changes. The top two strategies employed are communities of practice and grassroots, followed by proof of concept as a template (a pattern where the proof of concept gets reproduced elsewhere in the organization), and proof of concept as a seed. For an example of how a [community of practice](https://www.youtube.com/watch?v=S4-huVFeQXg)  works, [read about how Google's culture of comprehensive unit testing was driven by a group of volunteers] (https://martinfowler.com/articles/testing-culture.html#google).

分析表明，高绩效者倾向于在组织的低层和高层创建社区结构的策略，这可能使它们对重组和产品变化更具可持续性和弹性。采用的前两种策略是实践社区和草根社区，其次是作为模板的概念证明（概念证明在组织的其他地方被复制的模式），以及作为种子的概念证明。有关 [实践社区](https://www.youtube.com/watch?v=S4-huVFeQXg) 工作方式的示例，[了解 Google 的综合单元测试文化是如何由一群志愿者推动的] （https://martinfowler.com/articles/testing-culture.html#google)。

Low performers tend to favor training centers and centers of excellence: strategies that create more silos and isolated expertise. They also attempt proofs of concept, but these generally stall and don't see success. Why might these strategies fail to deliver effective change?

低绩效者倾向于支持培训中心和卓越中心：创造更多孤岛和孤立专业知识的策略。他们还尝试进行概念验证，但这些通常会停滞不前并且看不到成功。为什么这些策略可能无法带来有效的改变？

By centralizing expertise in one group, centers of excellence create several problems. First, the CoE is now a bottleneck for the relevant expertise for the organization and this cannot scale as demand for expertise in the organization grows. Second, it establishes an exclusive group of "experts" in the organization, in contrast to an inclusive group of peers who can continue to learn and grow together. This exclusivity can chip away at healthy organizational cultures. Finally, the experts are removed from doing the work. They are able to make recommendations or establish generic "best practices" but the path from the generic learning to the implementation of real work is left up to the learners. For example, experts will build a workshop on how to containerize an application, but they rarely or never actually containerize applications. This disconnect between theory and hands-on practice will eventually threaten their expertise. 

通过将专业知识集中在一个小组中，卓越中心会产生几个问题。首先，CoE 现在是组织相关专业知识的瓶颈，并且无法随着组织对专业知识的需求增长而扩展。其次，它在组织中建立了一个排他性的“专家”群体，而不是一个可以继续学习和共同成长的包容性同行群体。这种排他性会削弱健康的组织文化。最后，专家们不再从事这项工作。他们能够提出建议或建立通用的“最佳实践”，但从通用学习到实际工作实施的路径由学习者决定。例如，专家将建立一个关于如何容器化应用程序的研讨会，但他们很少或从未真正将应用程序容器化。理论与实践之间的这种脱节最终将威胁到他们的专业知识。

While some see success in training centers, they require dedicated resources and programs to execute both the original program and sustained learning. Many companies have set aside incredible resources to make their training programs effective: They have entire buildings dedicated to a separate, creative environment, and staff devoted to creating training materials and assessing progress. Additional resources are then needed to assure that the learning is sustained and propagated throughout the organization. The organization has to provide support for the teams that attended the training center, to help ensure their skills and habits are continued back in their regular work environments, and that old work patterns aren't resumed. If these resources aren't in place, organizations risk all of their investments going to waste. Instead of a center where teams go to learn new technologies and processes to spread to the rest of the organization, new habits stay in the center, creating another silo, albeit a temporary one. There are also similar limitations as in the CoE: If only the training center staff (or other, detached "experts") are creating workshops and training materials, what happens if they never actually do the work?

虽然有些人看到了培训中心的成功，但他们需要专门的资源和计划来执行原始计划和持续学习。许多公司已经拨出难以置信的资源来使他们的培训计划有效：他们有整栋大楼专门用于一个单独的、创造性的环境，以及致力于创建培训材料和评估进度的员工。然后需要额外的资源来确保学习在整个组织中得以持续和传播。该组织必须为参加培训中心的团队提供支持，以帮助确保他们的技能和习惯在他们的正常工作环境中得以延续，并且不会恢复旧的工作模式。如果这些资源没有到位，组织就会冒着浪费所有投资的风险。新习惯不再是团队去学习新技术和流程以传播到组织其他部门的中心，而是留在中心，创造了另一个孤岛，尽管是暂时的。也有与 CoE 类似的限制：如果只有培训中心的工作人员（或其他超然的“专家”）在创建研讨会和培训材料，如果他们从未真正从事过工作会发生什么？

Mashups were commonly reported (40% of the people responding to the 2019 survey used this strategy), but they lack sufficient funding and resources in any particular investment. Without a strategy to guide a technology transformation, organizations will often make the mistake of hedging their bets and suffer from *death by initiative*: identifying initiatives in too many areas, which ultimately leads to under-resourcing important work and dooming them all to failure . Instead, it is best to select a few initiatives and dedicate ongoing resources to ensure their success (time, money, and executive and champion practitioner sponsorship). In contrast to mashups, very few people report use of a big bang strategy, although it was most common in low performers.

混搭的报道很普遍（40% 的 2019 年调查受访者使用了这种策略），但它们在任何特定投资中都缺乏足够的资金和资源。如果没有指导技术转型的战略，组织往往会犯下对冲风险的错误，并遭受*死于主动性*：确定过多领域的计划，最终导致重要工作资源不足并注定失败.相反，最好选择一些倡议并投入持续的资源来确保它们的成功（时间、金钱以及高管和冠军从业者的赞助）。与混搭相比，很少有人报告使用大爆炸策略，尽管这种策略在低绩效者中最为常见。

Additional analysis identified four patterns used by high performers:

其他分析确定了高绩效者使用的四种模式：

- **Community builders:** This group focuses on communities of practice, grassroots, and proofs of concept (as a template and as a seed, as described earlier). This occurs 46% of the time.
- **University:** This group focuses on education and training, with the majority of their efforts going into centers of excellence, communities of practice, and training centers. This pattern was only observed 9% of the time, suggesting that while this strategy can be successful, it is not common and requires significant investment and planning to ensure that lessons learned are scaled throughout the organization.
- **Emergent:** This group has focused on grassroots efforts and communities of practice. This appears to be the most hands-off group and appears in 23% of cases.
- **Experimenters:** Experimenters appeared in 22% of cases. This group has high levels of activity in all strategies except big bang and dojos—that is, all activities that focus on community and creation. They also include high levels in PoCs that stall. The fact they are able to leverage this activity and remain high performers suggests they use this strategy to experiment and test out ideas quickly.

- **社区建设者：** 该小组专注于实践社区、基层社区和概念证明（作为模板和种子，如前所述）。这发生在 46% 的时间里。
- **大学：** 该小组专注于教育和培训，他们的大部分努力都投入到卓越中心、实践社区和培训中心。这种模式仅在 9% 的时间内被观察到，这表明虽然这种策略可以成功，但并不常见，需要大量投资和规划，以确保在整个组织中推广经验教训。
- **Emergent：** 该小组专注于基层工作和实践社区。这似乎是最不干涉的群体，出现在 23% 的案例中。
- **实验者：**实验者出现在 22% 的案例中。除了 big bang 和 dojos 之外，该组在所有策略中都具有高水平的活动——即所有专注于社区和创造的活动。它们还包括停滞不前的 PoC 中的高水平。他们能够利用这项活动并保持高绩效的事实表明他们使用这种策略来快速试验和测试想法。

## Principles of effective organizational change management

## 有效组织变革管理的原则

All organizations are complex, and every organization has different goals, a different starting point, and their own ways of approaching challenges. Prescriptions that work in one organization might not show the same results in another organization. However, your organization can follow some general principles in order to increase your chances of success.

所有组织都是复杂的，每个组织都有不同的目标、不同的起点和应对挑战的方式。在一个组织中有效的处方在另一个组织中可能不会显示相同的结果。但是，您的组织可以遵循一些一般原则，以增加成功的机会。

### Improvement work is never done 

### 改进工作从未完成

High-performing organizations are never satisfied with their performance and are always trying to get better at what they do. Improvement work is ongoing and baked into the daily work of teams. People in these organizations understand that failing to change is as risky as change, and they don't use "that's the way we've always done it" as a justification for resisting change. However that doesn't mean taking an undisciplined approach to change. Change management should be performed in a scientific way in pursuit of a measurable team or organizational goal.

高绩效组织永远不会对他们的表现感到满意，并且总是试图在他们所做的事情上做得更好。改进工作正在进行中，并融入团队的日常工作中。这些组织中的人明白，不改变与改变一样有风险，他们不会用“这就是我们一直这样做的方式”作为拒绝改变的理由。然而，这并不意味着采取无纪律的方法来改变。变革管理应以科学的方式进行，以追求可衡量的团队或组织目标。

### Leaders and teams agree on and communicate measurable outcomes, and teams determine how to achieve them

### 领导者和团队就可衡量的成果达成一致并进行沟通，团队决定如何实现这些成果

It's essential that everybody in the organization knows the measurable business and organizational outcomes that they are working towards. These outcomes should be short (a few sentences at most) at the organizational level, and match up clearly to the purpose and mission of the organization. At the level of an individual business unit, the outcomes should fit on a single page. The organizational outcomes should be decided by leaders and teams working together, although leaders have the ultimate authority. At lower levels of the organization, goals are stated in more detail and with shorter horizons.

至关重要的是，组织中的每个人都知道他们正在努力实现的可衡量的业务和组织成果。这些成果在组织层面应该简短（最多几句话），并且与组织的目的和使命明确匹配。在单个业务部门的级别，结果应该放在一个页面上。尽管领导者拥有最终权力，但组织成果应由领导者和团队共同决定。在组织的较低级别，目标被更详细地陈述并且视野更短。

However, it should be up to teams to decide how they go about achieving these outcomes, for these reasons:

但是，出于以下原因，应该由团队决定如何实现这些结果：

- Under conditions of uncertainty, it's impossible to decide the best course of action through planning alone. That doesn't mean some level of planning isn't important. But teams should be prepared to alter or even rewrite the plan based on what they discover when trying to execute it.
- When people are told both what to do and how to do it, they lose their autonomy and a chance to harness their ingenuity. Not only does this produce worse outcomes, it also leads to disengaged employees.
- Problem-solving is critical in helping employees develop new skills and capabilities. Organizations should give teams problems to solve, not tasks to execute.

- 在不确定的情况下，仅通过计划是不可能决定最佳行动方案的。这并不意味着某种程度的计划不重要。但是团队应该准备好根据他们在尝试执行计划时发现的内容来改变甚至重写计划。
- 当人们被告知该做什么和怎么做时，他们就会失去自主权和发挥创造力的机会。这不仅会产生更糟糕的结果，还会导致员工敬业度下降。
- 解决问题对于帮助员工发展新技能和能力至关重要。组织应该给团队解决问题，而不是执行任务。

### Large-scale change is achieved iteratively and incrementally

### 大规模的改变是迭代增量实现的

The annual budgeting cycle tends to drive organizations towards a project-based model in which work of all kinds is tied to expensive projects that take a long time to deliver. With few exceptions, it's better to break work down into smaller pieces that can be delivered incrementally. [Working in small batches](https://cloud.google.com/architecture/devops/devops-process-working-in-small-batches)  delivers a host of benefits. The most important is that it lets organizations correct course based on what they discover. This avoids wasting time and money doing work that doesn't deliver the expected benefits.

年度预算周期倾向于推动组织采用基于项目的模型，在这种模型中，各种工作都与需要很长时间才能交付的昂贵项目相关联。除了少数例外，最好将工作分解成可以增量交付的小块。 [小批量工作](https://cloud.google.com/architecture/devops/devops-process-working-in-small-batches) 带来了许多好处。最重要的是，它可以让组织根据他们的发现纠正路线。这可以避免浪费时间和金钱做没有达到预期收益的工作。

Moving from a [project paradigm to the product paradigm](https://itrevolution.com/book/project-to-product/)  is a long-term trend that will take most industries years to execute, but it's clear that this is the future. Even the US federal government has successfully [experimented with modular contracting](https://18f.gsa.gov/2019/04/09/why-we-love-modular-contracting/)  to pursue a more iterative, incremental approach to delivering large pieces of work.

从 [项目范式到产品范式](https://itrevolution.com/book/project-to-product/)是一个长期趋势，大多数行业需要数年时间才能执行，但很明显这是未来。甚至美国联邦政府也成功地[试验了模块化合同](https://18f.gsa.gov/2019/04/09/why-we-love-modular-contracting/)，以寻求一种更加迭代、增量的方法交付大量工作。

The issues that apply to delivering projects also apply to transformation. Organizations should find ways to achieve quick wins, share learning, and help other teams experiment with these new ideas.

适用于交付项目的问题也适用于转型。组织应该找到快速获胜的方法，分享学习，并帮助其他团队尝试这些新想法。

## Common pitfalls in transforming culture

## 文化转型的常见陷阱

Leaders often make the following mistakes when they attempt to make large-scale changes to an organization. 

领导者在尝试对组织进行大规模变革时经常会犯以下错误。

- **Treating transformation as a one-time project**. in high-performing organizations, getting better is an ongoing effort and part of everybody's daily work. However, many transformation programs are treated as large-scale, one-time events in which everyone is expected to rapidly change the way they work and then continue on with business as usual. Teams are not given the capacity, resources, or authority to improve the way they work, and their performance gradually degrades as the team's processes, skills and capabilities become an ever poorer fit for the evolving reality of the work. You should think of technology transformation as an important value-delivery part of the business, one that you won't stop investing in. After all, do you plan to stop investing in customer acquisition or customer support?
- **Treating transformation as a top-down effort**. In this model, organizational reporting lines are changed, teams are moved around or restructured, and new processes are implemented. Often, the people who are affected are given little control of the changes and are not given opportunities for input. This can cause stress and lost productivity as people learn new ways of working, often while they are still delivering on existing commitments. When combined with the poor communication that is frequent in transformation initiatives, the top-down approach can lead to employees becoming unhappy and disengaged. It's also uncommon for the team that's planning and executing the top-down transformation to gather feedback on the effects of their work and to make changes accordingly. Instead, the plan is executed regardless of the consequences.
- **Failing to agree on and communicate the intended outcome**. Transformations are sometimes executed with poorly defined goals, or with qualitative (rather than quantitative) goals, such as "faster time-to-market" or "lower costs." Sometimes goals are defined but are not achievable, or the goals pit one part of the organization against the other. In these cases, it's impossible to know whether the improvement work is having the intended effect. When this failure is combined with a top-down approach, it becomes hard to experiment with other approaches that might be faster or cheaper. The result is typically waste when the plan is executed, and an inability to determine whether the goal was achieved or the program worked. In many cases, instead of the failure being critically analyzed and used as a learning opportunity, the failure is ignored and new wholesale change initiative is started or entire methodologies are discredited.

- **将转型视为一次性项目**。在高绩效组织中，变得更好是持续的努力，也是每个人日常工作的一部分。然而，许多转型计划被视为大规模的一次性事件，在这些事件中，每个人都被期望迅速改变他们的工作方式，然后继续照常营业。团队没有能力、资源或权力来改进他们的工作方式，随着团队的流程、技能和能力越来越不适合不断发展的工作现实，他们的绩效会逐渐下降。您应该将技术转型视为业务中重要的价值交付部分，您不会停止投资。毕竟，您是否打算停止对客户获取或客户支持的投资？
- **将转型视为自上而下的努力**。在这个模型中，组织报告线发生了变化，团队被调动或重组，并实施了新的流程。通常，受影响的人几乎无法控制变化，也没有机会参与。当人们学习新的工作方式时，这可能会导致压力和生产力下降，而且通常他们仍在履行现有的承诺。再加上转型计划中经常出现的沟通不畅，自上而下的方法可能会导致员工变得不开心和不敬业。计划和执行自上而下转换的团队收集有关其工作效果的反馈并相应地进行更改的情况也很少见。相反，无论后果如何，计划都会被执行。
- **未能就预期结果达成一致并传达**。有时执行转换时的目标定义不明确，或者是定性（而不是定量）目标，例如“更快的上市时间”或“更低的成本”。有时目标已定义但无法实现，或者目标使组织的一部分与另一部分发生冲突。在这些情况下，不可能知道改进工作是否达到了预期的效果。当这种失败与自上而下的方法相结合时，就很难尝试其他可能更快或更便宜的方法。结果在执行计划时通常是浪费，并且无法确定目标是否实现或程序是否有效。在许多情况下，不是对失败进行批判性分析并将其用作学习机会，而是忽略失败并启动新的大规模变革计划，或者整个方法论都名誉扫地。

The combination of treating transformation as a project and treating it as a top-down initiative tends to lead to the pattern shown in the following diagram. Performance gradually degrades. At the start of a transformation program, it initially gets worse before improving. But then this is followed by a transition back to business as usual. All the while, cynicism and disengagement increases across the organization.

将转型视为一个项目和将其视为自上而下的举措相结合，往往会导致下图所示的模式。性能逐渐下降。在转型计划开始时，它最初会在改进之前变得更糟。但随之而来的是恢复正常业务的过渡。与此同时，整个组织的愤世嫉俗和疏离感都在增加。

![Graph of performance against time, showing performance going up during each change program but then declining and overall staying roughly the same.](https://cloud.google.com/static/architecture/devops/images/devops-culture-transform-performance-against-time.png)

转换性能对抗时间.png)

Source: CC-BY: [*Lean Enterprise: How High Performance Organizations Innovate at Scale*](https://www.google.com/books/edition/Lean_Enterprise/G_ixBQAAQBAJ)  by Jez Humble, Joanne Molesky, and Barry O' Reilly (O'Reilly, 2014).

资料来源：CC-BY：[*精益企业：高绩效组织如何大规模创新*](https://www.google.com/books/edition/Lean_Enterprise/G_ixBQAAQBAJ)，作者：Jez Humble、Joanne Molesky 和 Barry O'赖利（O'Reilly，2014 年)。

## What's next

##  下一步是什么

- For links to other articles and resources, see the [DevOps page](https://cloud.google.com/devops).
- Explore our DevOps [research program](https://www.devops-research.com/research.html).
- Take the [DevOps quick check](https://www.devops-research.com/quickcheck.html) to understand where you stand in comparison with the rest of the industry. 

- 有关其他文章和资源的链接，请参阅 [DevOps 页面](https://cloud.google.com/devops)。
- 探索我们的 DevOps [研究计划](https://www.devops-research.com/research.html)。
- 参加 [DevOps 快速检查](https://www.devops-research.com/quickcheck.html) 以了解您与行业其他人相比所处的位置。

