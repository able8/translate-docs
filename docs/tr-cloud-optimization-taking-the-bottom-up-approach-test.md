# How CI/CD is Sidetracking Optimization, and What You Can Do About It

# CI/CD 如何进行侧跟踪优化，以及您可以做些什么

December 4, 2019

2019 年 12 月 4 日

In a market hungry for new features and functionality, delivering bug-free code at high velocity has become imperative for businesses and engineering teams. Often times, this comes at the cost of infrastructure performance optimization.

在一个渴望新特性和功能的市场中，高速交付无错误代码已成为企业和工程团队的当务之急。通常，这是以基础设施性能优化为代价的。

Continuous integration, delivery, and deployment enable rapid product improvements through fast feature introduction and fast turn-around on feature changes. Testability is better, with simpler and quicker fault isolation, and time to resolution is shorter due to the smaller code changes. Gone are the days of deployment-windows scheduled weeks in advance—and no one is looking back. CI/CD has taken off like a house on fire, with many companies now operating a smooth CI/CD toolchain in an agile DevOps setting to rapidly deliver features, fixes and updates.

持续集成、交付和部署通过快速的功能引入和功能更改的快速周转来实现快速的产品改进。可测试性更好，故障隔离更简单更快，由于代码更改较小，解决时间更短。提前几周安排部署窗口的日子已经一去不复返了——没有人会回头。 CI/CD 像火一样的房子起飞，许多公司现在在敏捷的 DevOps 环境中运行流畅的 CI/CD 工具链，以快速交付功能、修复和更新。

Over time, the CI/CD pipeline created an adverse effect on system performance and optimization. Traditional deployment workflows supported manual point in time performance/cost analysis by performance engineers and cloud consultants, enabling sysadmins to balance performance and cost. The velocity of code changes introduced by continuous processes undermines this logic, as the manual point in time system optimization recommendations quickly become stale and irrelevant.

随着时间的推移，CI/CD 管道对系统性能和优化产生了不利影响。传统的部署工作流支持由性能工程师和云顾问进行手动时间点性能/成本分析，使系统管理员能够平衡性能和成本。连续流程引入的代码更改速度破坏了这一逻辑，因为手动时间点系统优化建议很快变得陈旧和无关紧要。

CI/CD best practices call for the integration of performance tests with the pipeline, using preset benchmarks to pass or fail pipelines. But developers are under even greater pressure to push business logic as quickly as possible. Typically, they don’t have the bandwidth to figure out optimal infrastructure settings including OS and kernel resource allocation.

CI/CD 最佳实践要求将性能测试与管道集成，使用预设基准来通过或失败管道。但是开发人员面临着更大的压力，需要尽快推送业务逻辑。通常，他们没有带宽来确定最佳基础设施设置，包括操作系统和内核资源分配。

Integrating real-time infrastructure optimization into DevOps processes is an intense and difficult undertaking. With SLA’s hanging overhead and to maintain high QoS, there is no time to constantly optimize infrastructure settings, not to mention OS and kernel resource allocation. To mitigate the risk, businesses run on relatively low utilization. This means paying for way more compute than is actually required at any given time in an optimized world. This overlooked price tag can take a toll on the business at large.

将实时基础架构优化集成到 DevOps 流程中是一项艰巨而艰巨的任务。由于 SLA 的开销悬而未决并保持高 QoS，没有时间不断优化基础设施设置，更不用说操作系统和内核资源分配了。为了降低风险，企业运行的利用率相对较低。这意味着在优化的世界中，在任何给定时间，支付的计算量都比实际需要的多。这个被忽视的价格标签可能会给整个企业带来损失。

### Shift it to the Left

### 向左移动

One way to improve performance in a constant sprint environment is to introduce testing in the early stages of development cycles. As opposed to the traditional approach of executing tests at the latest stages of dev, this practice, called Shift Left Performance Testing, allows smaller, ad hoc performance tests to be performed against individual components as they are developed.

在持续冲刺环境中提高性能的一种方法是在开发周期的早期阶段引入测试。与在开发的最新阶段执行测试的传统方法相反，这种称为左移性能测试的实践允许在开发时针对单个组件执行更小的临时性能测试。

Aligning performance with functional and unit tests means that teams need to begin performance testing when functionality is implemented, and configure those performance tests to run automatically and alert about decreases in performance. Performance testing is integrated as a part of the CI/CD process and executed in local environments after each code commit.

将性能与功能和单元测试保持一致意味着团队需要在实现功能时开始性能测试，并将这些性能测试配置为自动运行并在性能下降时发出警报。性能测试被集成为 CI/CD 流程的一部分，并在每次代码提交后在本地环境中执行。

### Shift it to Autonomous

### 将其切换为自治

Another approach to performance optimization for changing environments involves continuous autonomous optimization. The growing system architectural complexity has given rise to this new generation of tools, geared at bridging the gap between what is humanly possible to optimize and what can actually be achieved by unlocking the vast potential of the machine.

另一种针对不断变化的环境进行性能优化的方法涉及持续的自主优化。不断增长的系统架构复杂性催生了新一代工具，旨在弥合人类可以优化的内容与通过释放机器的巨大潜力实际可以实现的内容之间的差距。

Autonomous performance optimization uses ML to configure the infrastructure, including OS and kernel resource allocation in real-time. Balancing cost and performance, these solutions tune the infrastructure on the fly precisely to the workload and goals of the application. Because autonomous performance optimization technology can employ continuous, seamless system adaptations and optimizations to create a streamlined app environment. 

自治性能优化使用 ML 来配置基础架构，包括实时操作系统和内核资源分配。在平衡成本和性能的同时，这些解决方案可以根据应用程序的工作负载和目标即时调整基础架构。因为自主性能优化技术可以采用连续、无缝的系统适配和优化来创建一个精简的应用程序环境。

Without automated infrastructure optimization, CI/CD would ultimately mean you’ll continue to accumulate excessive compute costs. Taking a real-time, automated approach to infrastructure optimization is the simplest and most cost-effective solution to this problem, which isn’t going anywhere any time soon. As is often the case, innovative technologies create problems that only other innovative technologies can solve. 

如果没有自动化的基础架构优化，CI/CD 最终将意味着您将继续积累过多的计算成本。采取实时、自动化的基础设施优化方法是解决这个问题的最简单、最具成本效益的解决方案，短期内不会有任何进展。通常情况下，创新技术会产生只有其他创新技术才能解决的问题。

