# DevOps measurement: Visual management capabilities

# DevOps 测量：可视化管理能力

https://cloud.google.com/architecture/devops/devops-measurement-visual-management

**Note:** *Visual management capabilities* is one of a set of capabilities that drive higher software delivery and organizational performance. These capabilities were discovered by the [DORA State of DevOps research program](https://www.devops-research.com/research.html), an independent, academically rigorous investigation into the practices and capabilities that drive high performance. To learn more, read our [DevOps resources](https://cloud.google.com/devops).

**注意：** *可视化管理功能*是推动更高软件交付和组织绩效的一组功能之一。这些能力是由 [DORA State of DevOps 研究计划](https://www.devops-research.com/research.html) 发现的，该计划是对推动高性能的实践和能力进行独立的、学术上严谨的调查。要了解更多信息，请阅读我们的 [DevOps 资源](https://cloud.google.com/devops)。

It's a common practice for teams that are adopting [lean development practices](https://wikipedia.org/wiki/Lean_software_development)  to display key information about their processes in team areas where everybody can see it. Visual management boards can create a shared understanding of where the team is in terms of its operational effectiveness. They can also help identify and remove obstacles in the path to higher performance.

对于采用 [精益开发实践](https://wikipedia.org/wiki/Lean_software_development) 的团队来说，在每个人都可以看到的团队区域中显示有关其流程的关键信息是一种常见的做法。可视化管理委员会可以就团队在运营效率方面所处的位置达成共识。它们还可以帮助识别和消除通往更高绩效的道路上的障碍。

## How to implement visual management

## 如何实现可视化管理

There are many kinds of visual displays and dashboards that are common in the context of software delivery:

在软件交付的上下文中，有多种视觉显示和仪表板很常见：

- **Card walls, storyboards or Kanban boards**, either physical or virtual, with index cards  that represent in-progress work items.
- **Dashboards or other visual indicators**, such as continuous integration systems with monitors or traffic lights to show whether the build is passing or failing. Effective visual displays are created, updated, and perhaps discarded by teams in response to issues that the team is currently interested in addressing.
- **Burn-up or burn-down charts** (for example, cumulative flow diagrams) showing the cumulative status of all work being done. These allow the team to project how long it will take to complete the current backlog.
- **Deployment pipeline monitors** showing what the latest deployable build is, and whether stages in the pipeline are failing, such as acceptance tests or performance tests.
- **Monitors showing production telemetry**, such as the number of requests being received, latency statistics, cumulative `404` and `500` errors, and which pages are most popular.

- **卡片墙、故事板或看板**，无论是物理的还是虚拟的，带有代表正在进行的工作项目的索引卡。
- **仪表板或其他视觉指标**，例如带有监视器或交通灯的持续集成系统，以显示构建是通过还是失败。团队创建、更新甚至可能丢弃有效的视觉显示，以响应团队当前有兴趣解决的问题。
- **燃尽图或燃尽图**（例如，累积流程图）显示所有正在完成的工作的累积状态。这些允许团队预测完成当前积压工作需要多长时间。
- **部署管道监视器**显示最新的可部署构建是什么，以及管道中的阶段是否失败，例如验收测试或性能测试。
- **显示生产遥测数据的监视器**，例如收到的请求数、延迟统计信息、累积的“404”和“500”错误，以及哪些页面最受欢迎。

When combined with the use of [WIP limits](https://cloud.google.com/architecture/devops/devops-measurement-wip-limits)  and using feedback from production to make business decisions, visual management displays can contribute to [ higher levels of delivery performance](https://services.google.com/fh/files/misc/state-of-devops-2015.pdf#page=15)  (PDF).

当结合使用 [WIP 限制](https://cloud.google.com/architecture/devops/devops-measurement-wip-limits) 并使用生产反馈来制定业务决策时，可视化管理显示可以有助于 [更高水平的交付性能](https://services.google.com/fh/files/misc/state-of-devops-2015.pdf#page=15)(PDF)。

## Common pitfalls with visual management

## 可视化管理的常见缺陷

The most important characteristics of visual management displays are that the team cares about and will act upon the information, and that the display is used during daily work to identify and remove obstacles to higher performance. Common pitfalls when implementing visual management include the following:

可视化管理展示最重要的特点是团队关心并会根据信息采取行动，并且在日常工作中使用展示来识别和消除提高绩效的障碍。实施可视化管理时的常见陷阱包括：

- **Selecting metrics without the involvement of the team**. Visual displays that show metrics that are highly relevant and useful to teams will be used more often. In addition, if teams can have input into the metrics that are displayed on their visual displays by participating in selecting their goals (for example, some teams [use OKRs](https://cloud.google.com/architecture/devops/devops-culture-transform)), they will be more motivated to drive progress toward those goals.
- **Creating displays that are complex, hard to understand, or do not provide actionable information**. It's easy to create displays using tools that allow high levels of modification or that are fun to play with. But changing layouts and color on a custom dashboard isn't helpful if the team is working with the wrong metrics or it takes the team several months to implement. Key metrics or rough graphs drawn on a whiteboard and updated daily can be just as effective to keep the team informed. 

- **在没有团队参与的情况下选择指标**。显示与团队高度相关和有用的指标的可视化显示将更频繁地使用。此外，如果团队可以通过参与选择他们的目标（例如，一些团队 [使用 OKR](https://cloud.google.com/architecture/devops/devops-文化转变))，他们将更有动力推动实现这些目标。
- **创建复杂、难以理解或不提供可操作信息的显示**。使用允许高水平修改或玩起来很有趣的工具很容易创建显示。但是，如果团队使用错误的指标或者团队需要几个月的时间来实施，那么更改自定义仪表板上的布局和颜色就没有帮助。在白板上绘制并每天更新的关键指标或粗略图表对于让团队了解情况同样有效。

- **Not evolving visual displays**. Visual management tools should provide teams with information that addresses issues they are facing right now. It doesn't help one team to copy the displays of other teams unless the teams work in the same context, with the same challenges and obstacles. As a team's context evolves, visual displays should change as well. Also note that as teams address obstacles, their visual displays might change to discard old (previously relevant) metrics and highlight new areas of importance.
- **Not addressing the underlying problem that the visual display is revealing**. Teams sometimes make quick fixes in an effort to make the display "green" again. Displays should be used to drive improvements (fix the problem), not become a goal in themselves (keep the board green). Focusing on managing the metric alone leads to unintended consequences and technical debt. If the display suggests a problem, teams should not just fix the immediate problem. They should also work to identify the underlying issue or constraint and resolve it, even if it's in another part of the organization. Any inefficiencies will keep showing up, and fixing them earlier will help all teams.

- **不发展视觉显示**。可视化管理工具应该为团队提供解决他们目前面临的问题的信息。除非团队在相同的环境中工作，面临相同的挑战和障碍，否则复制其他团队的展示对一个团队没有帮助。随着团队环境的发展，视觉显示也应该改变。另请注意，当团队解决障碍时，他们的视觉显示可能会改变以丢弃旧的（以前相关的）指标并突出新的重要领域。
- **没有解决视觉显示所揭示的根本问题**。团队有时会进行快速修复，以使显示再次“变绿”。展示应该用于推动改进（解决问题），而不是本身成为目标（保持董事会绿色）。只专注于管理指标会导致意想不到的后果和技术债务。如果显示显示有问题，团队不应该只解决眼前的问题。他们还应该努力识别潜在的问题或约束并解决它，即使它在组织的另一部分。任何低效率都会不断出现，尽早修复它们将有助于所有团队。

## Ways to improve visual management

## 改善可视化管理的方法

The goal of visual management tools is to provide fast, easy-to-understand feedback so you can build quality into the product. This feedback helps the team identify defects in the product and understand whether some part of the system is not performing effectively, which helps them address the problem. In order to be effective, such systems must do the following:

可视化管理工具的目标是提供快速、易于理解的反馈，以便您可以在产品中构建质量。此反馈有助于团队识别产品中的缺陷并了解系统的某些部分是否没有有效执行，这有助于他们解决问题。为了有效，此类系统必须执行以下操作：

- **Reflect information that the team cares about and will act on**. Having build monitors does no good if teams don't care whether the display shows an issue (for example, showing that the build status is red, meaning broken), and won't actually act on this information by swarming to fix the issue.
- **Be easy to understand**. It should be possible to tell at a glance from across the room whether something needs attention. If there *is* a problem, teams should know how to perform further diagnosis or fix the problem.
- **Give the team information that is relevant to their work**. While it's important to collect as much data as possible about the team's work, the display should present only data that is relevant to the team's goals. In the face of information overload, particularly information that cannot be acted upon, people ignore visual management displays; the displays just become noise. The additional data can be accessed and used by the team when they are swarming to fix the problem.
- **Be updated as part of daily work**. If the team lets the data go stale or become inaccurate, they will ignore the visual displays, and the displays will no longer be a useful beacon when important issues arise. If displays are currently displaying stale or inaccurate data, investigate the cause: is the data not related to the team's goals? What data would make the display an important and compelling information source for the team?

- **反映团队关心并将采取行动的信息**。如果团队不关心显示是否显示问题（例如，显示构建状态为红色，表示已损坏）并且实际上不会通过蜂拥而至解决问题来实际处理此信息，那么拥有构建监视器并没有什么好处。
- **易于理解**。应该可以从房间对面一眼看出是否需要注意。如果*存在*问题，团队应该知道如何进行进一步诊断或解决问题。
- **提供与其工作相关的团队信息**。虽然尽可能多地收集有关团队工作的数据很重要，但显示应该只显示与团队目标相关的数据。面对信息超载，尤其是无法采取行动的信息，人们忽略了可视化管理展示；显示器只是变成噪音。当团队蜂拥而至解决问题时，他们可以访问和使用额外的数据。
- **作为日常工作的一部分进行更新**。如果团队让数据过时或变得不准确，他们将忽略视觉显示，当出现重要问题时，显示将不再是有用的灯塔。如果显示器当前显示陈旧或不准确的数据，请调查原因：数据是否与团队目标无关？哪些数据会使显示器成为团队重要且引人注目的信息来源？

Teams shouldn't get caught up in aspects of visual displays that aren't critical. For example, visual management displays don't need to be electronic. Physical card walls or kanban boards can be easier to manage and understand, particularly if the team is all in one location. These displays can also help develop valuable team rituals such as physically standing in front of the board to pick up work and move it around. A whiteboard with some key project information that is updated daily by the team is often preferable to an electronic system that's hard to understand, difficult to update, or doesn't have necessary information.

团队不应该陷入不重要的视觉显示方面。例如，可视化管理显示不需要是电子的。物理卡片墙或看板可以更容易管理和理解，特别是如果团队都在一个位置。这些展示还可以帮助培养有价值的团队仪式，例如站在板前拿起工作并四处走动。包含团队每天更新的一些关键项目信息的白板通常比难以理解、难以更新或没有必要信息的电子系统更可取。

## Ways to measure visual management

## 衡量视觉管理的方法

As with all improvement work, start with the measurable system-level goals that the team is working toward. Discover the existing state of the work system. Find a way to display the key information about the existing state, as well as the state you want. Make sure that this information is displayed only to the required precision.

与所有改进工作一样，从团队正在努力实现的可衡量的系统级目标开始。发现工作系统的现有状态。找到一种方法来显示有关现有状态以及您想要的状态的关键信息。确保仅以所需的精度显示此信息。

Review the visual displays as part of regular retrospectives. Ask these questions:

回顾视觉展示作为定期回顾的一部分。问这些问题：

- Are the displays giving you the information you need?
- Is the information up to date?
- Are people acting on this information? 

- 显示器是否为您提供所需的信息？
- 信息是最新的吗？
- 人们是否根据这些信息采取行动？

- Is the information (and the actions people take in response to it) contributing to measurable improvement towards a goal that the team cares about?
- Does everybody know what the goals are?
- Can you look at your visual management displays and see the key process metrics you care about?

- 信息（以及人们为响应信息而采取的行动）是否有助于实现团队关心的目标的可衡量改进？
- 每个人都知道目标是什么吗？
- 您能否查看您的可视化管理显示并查看您关心的关键流程指标？

If the answer to any of these questions is no, investigate further:

- Can you change the information or how it's displayed?
- Can you get rid of the display altogether?
- Can you create a new display? What would a prototype look like? What are the most important pieces of information to include, and how precise do they need to be to help you solve your problems and achieve your goals? 

如果这些问题的答案是否定的，请进一步调查：

- 你能改变信息或它的显示方式吗？
- 你能完全摆脱显示器吗？
- 你能创建一个新的显示器吗？原型会是什么样子？要包括哪些最重要的信息，它们需要多精确才能帮助您解决问题并实现目标？

