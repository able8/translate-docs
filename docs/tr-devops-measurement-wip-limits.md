# DevOps measurement: Work in process limits

# DevOps 测量：在制品限制

**Note:** *Work in process limits* is one of a set of capabilities that drive higher software delivery and organizational performance. These capabilities were discovered by the [DORA State of DevOps research program](https://www.devops-research.com/research.html), an independent, academically rigorous investigation into the practices and capabilities that drive high performance. To learn more, read our [DevOps resources](https://cloud.google.com/devops).

**注意：** *在制品限制*是推动更高软件交付和组织绩效的一组功能之一。这些能力是由 [DORA State of DevOps 研究计划](https://www.devops-research.com/research.html) 发现的，该计划是对推动高性能的实践和能力进行独立的、学术上严谨的调查。要了解更多信息，请阅读我们的 [DevOps 资源](https://cloud.google.com/devops)。

When faced with too much work and too few people to do it, rookie managers assign people to work on multiple tasks in the hope of increasing throughput. Unfortunately, the result is that tasks take longer to get done, and the team burns out in the process.

当面临工作太多而人太少时，新手经理会指派人员从事多项任务，以期提高吞吐量。不幸的是，结果是任务需要更长的时间才能完成，并且团队在此过程中精疲力尽。

Instead, you should do the following:

- Prioritize work
- Limit how much people work on
- Focus on completing a small number of high-priority tasks

相反，您应该执行以下操作：

- 优先工作
- 限制人们的工作量
- 专注于完成少数高优先级任务

The manufacturing sector has a long history of limiting the amount of work in process (WIP). Factories don't hold large amounts of inventory. Instead, when a customer orders a product, parts are made in-house, on-demand or are pulled from suppliers upstream as needed, and the company then assembles the product just in time. When you implement this process correctly, you end up with shorter lead times, higher quality, lower costs, and less waste.

制造业在限制在制品 (WIP) 数量方面有着悠久的历史。工厂没有大量库存。取而代之的是，当客户订购产品时，零件会在内部按需制造或根据需要从上游供应商处拉出，然后公司及时组装产品。当您正确实施此流程时，您最终会获得更短的交货时间、更高的质量、更低的成本和更少的浪费。

## How to implement work in process limits

## 如何实现在制品限制

**Use a storyboard.** In technology, our inventory is invisible. There's no shop floor with piles of work or assembly line where we can see the progression of work. A simple way to see inventory is to write all the work the team is doing on index cards and stick them on a board. In agile methods, this is called creating a *storyboard*.

**使用故事板。**在技术方面，我们的库存是无形的。没有车间有成堆的工作或装配线，我们可以看到工作的进展。查看库存的一个简单方法是将团队正在做的所有工作写在索引卡上并将它们粘贴在板上。在敏捷方法中，这称为创建*故事板*。

The following sample storyboard spans multiple functions (analysis, development, testing, operations) and shows all work for a single product.

以下示例故事板涵盖多个功能（分析、开发、测试、操作），并显示了单个产品的所有工作。

![image](https://cloud.google.com/static/architecture/devops/wip-1.png)

*(Source: "Kanban for Ops" board game, Dominica DeGrandis 2013")*

A common practice with storyboards is to ink a dot onto a card for every day the card has been worked on. The team can easily see which work is blocked or taking longer than it should.

故事板的一个常见做法是在卡片上工作的每一天都在卡片上画一个点。团队可以轻松查看哪些工作被阻止或花费的时间超过了应有的时间。

**Specify limits.** For each column on the board, specify the WIP limit, or how many cards can be in that column at one time. After the WIP limit is reached, no more cards can be added to the column, and the team must wait for a card to move to the next column before pulling the highest priority one from the previous column.

**指定限制。** 对于板上的每一列，指定 WIP 限制，或该列中一次可以有多少张卡片。达到 WIP 限制后，不能再向该列添加卡片，团队必须等待一张卡片移动到下一列，然后才能从前一列中拉出优先级最高的一张。

Only by imposing WIP limits and following this pull-based process do you actually create a Kanban board.

只有通过施加 WIP 限制并遵循这个基于拉的流程，您才能真正创建看板。

**Determine WIP limits by team capacity.** For example, if you have four pairs of developers, don't allow more than four cards in the "in development" column.

**根据团队容量确定 WIP 限制。** 例如，如果您有四对开发人员，则“开发中”列中的卡片数量不得超过四张。

**Stick to the limits.** WIP limits can result in teams sitting idle, waiting for other tasks to be completed. Don't increase WIP limits at this point. Instead, work to improve your processes to address the factors that are contributing to these delays. For example, if you're waiting for an environment to test your work, you might offer to help the team that prepares environments improve or streamline their process.

**坚持限制。** WIP 限制可能导致团队闲置，等待其他任务完成。此时不要增加 WIP 限制。相反，努力改进您的流程以解决导致这些延迟的因素。例如，如果您正在等待一个环境来测试您的工作，您可能会主动提出帮助准备环境的团队改进或简化他们的流程。

## Common pitfalls with work in process limits

## 在制品限制的常见缺陷

Organizations implementing WIP limits often encounter the following pitfalls:

实施 WIP 限制的组织经常遇到以下陷阱：

- **Not counting invisible work.** It's important to visualize the whole value stream from idea to customer, not just the portion of the work that the team is responsible for. Without doing this, it's impossible to see the actual bottlenecks, and you'll end up addressing problems that aren't actually significant constraints to the flow of work. (This is also known as local optimums.)
- **Setting WIP limits that are much too big.** Make sure your WIP limits aren't too big. If your team is splitting their time between multiple tasks or projects, that's a good sign your WIP limits are too high.
- **Relaxing WIP limits.** Don't relax your WIP limits when people are idle. Instead, those people should be helping in other parts of the value stream, addressing the problems that are leading to constraints elsewhere. 

- **不包括无形的工作。**重要的是可视化从创意到客户的整个价值流，而不仅仅是团队负责的部分工作。如果不这样做，就不可能看到实际的瓶颈，并且您最终会解决实际上对工作流程没有重大限制的问题。 （这也称为局部最优。）
- **设置太大的 WIP 限制。** 确保您的 WIP 限制不会太大。如果您的团队在多个任务或项目之间分配时间，这表明您的 WIP 限制太高了。
- **放宽 WIP 限制。** 当人们空闲时，不要放宽 WIP 限制。相反，这些人应该在价值流的其他部分提供帮助，解决导致其他部分限制的问题。

- **Quitting while you're ahead.** If your WIP limits are easy to achieve, reduce them. The point of WIP limits is to expose problems in the system so they can be addressed. Another thing to look for is when there are too many columns on your visual display. Instead, look for ways to simplify the delivery process and reduce hand-offs. Process improvement work is key to increasing flow.

- **在领先时退出。**如果您的 WIP 限制很容易达到，请减少它们。 WIP 限制的目的是暴露系统中的问题，以便解决这些问题。另一件要寻找的事情是当你的视觉显示上有太多的列时。相反，寻找简化交付过程和减少交接的方法。流程改进工作是增加流量的关键。

DevOps Research & Assessment research shows that WIP limits help drive improvements in software delivery performance, particularly when they are combined with [the use of visual displays](https://cloud.google.com/architecture/devops/devops-measurement-visual-management)  and [feedback loops from monitoring](https://cloud.google.com/architecture/devops/devops-measurement-monitoring-systems).

DevOps 研究与评估研究表明，WIP 限制有助于推动软件交付性能的改进，特别是当它们与 使用可视化显示  结合使用时和[来自监控的反馈循环](https://cloud.google.com/architecture/devops/devops-measurement-monitoring-systems)。

## Ways to improve work in process limits

## 改进在制品限制的方法

- **Make your work visible.** As you do this, try to surface all of your work, making all of it visible, to several teams and stakeholders. (See [visual displays](https://cloud.google.com/architecture/devops/devops-measurement-visual-management)  for details).
- Set WIP limits that match your team's capacity for work.
   - Account for activities like production support, meeting time and technical debt.
   - Don't allow more WIP in any given part of the process than you have people to work on tasks.
   - Don't require people to split their time between multiple tasks.
   - When a particular piece of work is completed, move the card representing that work to the next column, and pull the highest priority piece of work waiting in the queue.
- **Set up a weekly meeting for stakeholders to prioritize all work in order.** Let stakeholders know that if they don't attend, their work won't get done.
- **Work to increase flow.** Measure the lead time of work through the system. Record the date that work started on a card and the date work ended. From this information, you can create a running frequency histogram, which shows the number of days work takes to go through the system. This data will allow you to calculate the mean lead time, as well as variability, with the goal of having low variability: high variability means you are not scoping projects well or have significant constraints outside of your team. High variability also means your estimates and predictions about future work will not be as reliable.
- **Improve work processes.** Reduce hand-offs, simplify and automate tasks, and think about how to collaborate better to get work done. After you've removed some obstacles and things feel comfortable, reduce your WIP limits to reveal the next set of obstacles. The ideal is single-piece flow, which means that work flows from idea to customer with minimal wait time or rework. This ideal may not be achievable, but it acts as a "true north" to guide the way in a process of continuous improvement.

- **让您的工作可见。** 当您这样做时，请尝试向多个团队和利益相关者展示您的所有工作，使其全部可见。 （有关详细信息，请参阅 [视觉显示](https://cloud.google.com/architecture/devops/devops-measurement-visual-management))。
- 设置与您团队的工作能力相匹配的 WIP 限制。
  - 考虑生产支持、会议时间和技术债务等活动。
  - 在流程的任何给定部分中，不允许在制品数量超过工作人员的数量。
  - 不要要求人们在多个任务之间分配时间。
  - 当一个特定的工作完成时，将代表该工作的卡片移动到下一列，并拉出队列中等待的最高优先级的工作。
- **为利益相关者安排每周会议，按顺序排列所有工作的优先级。** 让利益相关者知道，如果他们不参加，他们的工作将无法完成。
- **工作以增加流量。**通过系统测量工作的提前期。在卡片上记录工作开始的日期和工作结束的日期。根据这些信息，您可以创建运行频率直方图，它显示了通过系统所需的工作天数。这些数据将允许您计算平均提前期以及可变性，以实现低可变性的目标：高可变性意味着您没有很好地确定项目范围或在团队之外有重大限制。高可变性还意味着您对未来工作的估计和预测将不那么可靠。
- **改进工作流程。** 减少交接、简化和自动化任务，并思考如何更好地协作以完成工作。在您消除了一些障碍并且感觉舒适之后，降低您的 WIP 限制以显示下一组障碍。理想的是单件流，这意味着工作从创意流向客户，等待时间或返工时间最短。这个理想可能无法实现，但它充当了“真正的北方”，在不断改进的过程中指引方向。

## Ways to measure work in process limits

## 衡量在制品限制的方法

WIP limits are something you impose rather than measure, but it's important to keep finding ways to improve. During your regular retrospectives, ask the following questions:

WIP 限制是您强加而不是衡量的东西，但重要的是要不断寻找改进的方法。在您的定期回顾中，提出以下问题：

- Do we know the mean lead time and variability for our entire value stream (from idea to customer)?
- Are we finding ways to increase flow and thus reduce lead time for work?
- Are our WIP limits surfacing obstacles that prevent us increasing flow?
- Are we doing things about those obstacles?

- 我们是否知道整个价值流（从创意到客户）的平均交付周期和可变性？
- 我们是否想方设法增加流量，从而缩短工作准备时间？
- 我们的 WIP 限制是否会出现阻碍我们增加流量的障碍？
- 我们是否正在针对这些障碍采取措施？

