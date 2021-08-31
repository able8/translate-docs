# Everything is broken, and it’s okay

# 一切都坏了，没关系

Accepting that imperfect things still work is fundamental to preventing failures from becoming catastrophes.

接受不完美的事情仍然有效是防止失败变成灾难的基础。

Issue 16 February 2021

2021 年 2 月 16 日

Everything is a little bit broken. Nothing made by human hands or minds is perfect. Every car you’ve ever ridden in, every elevator you’ve ever taken, every safety-critical computer program you’ve ever trusted your life with was flawed in some way. It doesn’t matter how much money you spend—your system’s uptime will always be measured in 9s, a percentage of perfection.

一切都有些破碎。人的手或头脑制造的任何东西都不是完美的。你坐过的每一辆车，你坐过的每部电梯，你曾经信任过的每一个安全关键的计算机程序都在某种程度上存在缺陷。无论您花多少钱，系统的正常运行时间始终以 9 秒为单位，即完美的百分比。

But lots of imperfect things are still perfectly operational, including human bodies and many of our systems. This fact—that imperfect things still work—is integral to understanding how to prevent failures from escalating into catastrophes.

但是许多不完美的东西仍然可以完美运行，包括人体和我们的许多系统。这个事实——不完美的东西仍然有效——对于理解如何防止失败升级为灾难是不可或缺的。

A failure is one part of a system breaking; a catastrophe is when many failures accumulate to a point beyond recovery.1 When a catastrophe happens, it often seems like something very safe failed suddenly. But when we analyze the contributing causes, we find it wasn’t really sudden at all: The warning signs were present, the early failures, but we didn’t predict how they’d combine.

故障是系统崩溃的一部分；灾难是指许多失败累积到无法恢复的程度。1当灾难发生时，通常看起来很安全的事情突然失败了。但是当我们分析促成原因时，我们发现它根本不是突然的：警告信号存在，早期失败，但我们没有预测它们如何结合。

## Control is an illusion

##控制是一种错觉

The way we control computers’ inputs and outputs has evolved radically over the past two centuries. When inventor [Charles Babbage designed his Analytical Engine](https://www.britannica.com/technology/Analytical-Engine) in the mid-1800s and mathematician [Ada Lovelace programmed it](https://www.britannica.com /biography/Ada-Lovelace), they were manipulating physical objects. Now, we work with digital objects like software and functions rather than punch cards and readers. This shift, from manually altering a machine’s structure to using layers of computer languages and abstraction, has changed our conception of how we work and what we imagine we can control.

在过去的两个世纪里，我们控制计算机输入和输出的方式发生了根本性的变化。 1800 年代中期，发明家 [Charles Babbage 设计了他的分析引擎](https://www.britannica.com/technology/Analytical-Engine) 和数学家 [Ada Lovelace 对其进行了编程](https://www.britannica.com /biography/Ada-Lovelace），他们正在操纵物理对象。现在，我们使用软件和功能等数字对象，而不是打孔卡和读卡器。这种从手动更改机器结构到使用计算机语言和抽象层的转变，已经改变了我们对工作方式以及我们可以控制的想象的概念。

None of the changes we’ve made through the decades have reduced the amount of work that happens in a system, though. From chip design to developer tools, we’re not saving any time or effort, just redistributing labor. Developers now routinely reuse and conceptually compress the work of others; we assume the work underlying our own—the electrical engineering, the manufacturing precision, the code we build on top of, the tools we use to make new tools—is a given and focus only on the tasks in our scope of influence.

然而，我们几十年来所做的任何改变都没有减少系统中发生的工作量。从芯片设计到开发人员工具，我们没有节省任何时间或精力，只是重新分配劳动力。开发人员现在经常重用并在概念上压缩他人的工作；我们假设我们自己的基础工作——电气工程、制造精度、我们构建的代码、我们用来制造新工具的工具——是既定的，并且只专注于我们影响范围内的任务。

If we want to build stable and reliable systems, we have to admit that there are parts we don’t understand. Otherwise, we’ll go on assuming the systems we’re using are fixed, immutable, or at least stable. We’ll assume past performance is an indicator of future performance. But that’s not necessarily true, nor within our control—we just act like it is so we can get our work done.2

如果我们想构建稳定可靠的系统，我们必须承认有我们不了解的部分。否则，我们将继续假设我们使用的系统是固定的、不可变的，或者至少是稳定的。我们假设过去的表现是未来表现的指标。但这不一定是真的，也不在我们的控制范围内——我们只是表现得如此，以便我们可以完成我们的工作。 2

One of the clearest examples of our dependence on the unseen work of others is the [Spectre vulnerability](https://meltdownattack.com/), officially made public in January 2018, which affects microprocessors that perform branch prediction. Mitigating it required a change in how our CPUs function—not working ahead of actual commands—and caused a worldwide slowdown in execution speeds. Few developers consider the chips their software will run on. Why would they?

我们依赖他人看不见的工作的最明显例子之一是 [Spectre 漏洞](https://meltdownattack.com/)，该漏洞于 2018 年 1 月正式公开，它影响执行分支预测的微处理器。缓解它需要改变我们的 CPU 的工作方式——而不是在实际命令之前工作——并导致执行速度在全球范围内放缓。很少有开发人员考虑他们的软件将在哪些芯片上运行。他们为什么要这样做？

## Failure is inevitable

##失败是不可避免的

We know complex systems can fail. But what makes failure itself complex is that we don’t always know when it will happen, why, or how it will affect other parts of the system. 
我们知道复杂的系统可能会失败。但让失败本身变得复杂的原因在于，我们并不总是知道它何时会发生、为什么会发生，或者它会如何影响系统的其他部分。
The real question is whether that failure will become a catastrophe. As an example, every airplane you’ve ever flown on has had many tiny problems. Most of them are known and will be fixed at some point—things like a sticky luggage latch or a broken seat or a frayed seatbelt. None of these problems alone are cause to ground the plane. However, if enough small problems compound, the plane may no longer [meet the requirements for passenger airworthiness](https://www.icao.int/safety/airnavigation/OPS/CabinSafety/Pages/ICAO-Requirements-related-to- Cabin-Safety.aspx) and the airline will ground it. A plane with many malfunctioning call buttons may also be poorly maintained in other ways, like faulty checking for turbine blade microfractures or landing gear behavior.

真正的问题是这种失败是否会成为一场灾难。举个例子，你坐过的每架飞机都有很多小问题。它们中的大多数都是已知的，并且会在某个时候修复——比如行李箱闩锁、座椅破损或安全带磨损。仅凭这些问题都不会导致飞机接地。但是，如果出现足够多的小问题，飞机可能不再[满足乘客适航要求](https://www.icao.int/safety/airnavigation/OPS/CabinSafety/Pages/ICAO-Requirements-related-to- Cabin-Safety.aspx），航空公司会将其停飞。有许多呼叫按钮故障的飞机也可能在其他方面维护不善，例如对涡轮叶片微裂纹或起落架行为的错误检查。

## Responding to fragility

##应对脆弱性

> If everything is the most important thing, then nothing is.

> 如果一切都是最重要的，那么什么都不是。

Once you accept failure is inevitable and commit to avoiding the concatenation of failures that causes catastrophes, you—and your organization—need to decide what risk looks like and what to prioritize. If everything is the most important thing, then nothing is.

一旦您接受失败是不可避免的，并致力于避免导致灾难的一系列失败，您和您的组织就需要决定风险是什么样的以及优先考虑什么。如果一切都是最重要的，那么什么都不是。

We all want to believe our software or services do something vital for our users, but some industries bear more significant risks. Life-critical industries such as aviation, medical devices, and nuclear energy have their own standards, while leisure-based industries like gaming, gambling, or fashion have standards more to do with loss of market share than immediate, life-limiting effects. Assessing the error’s impact, such as the number of people affected or the monetary consequences of failure, can help you make reasoned choices about allocating your human and financial resources for mitigation.

我们都希望相信我们的软件或服务为我们的用户做了一些至关重要的事情，但有些行业承担了更大的风险。航空、医疗设备和核能等生命攸关的行业有自己的标准，而游戏、赌博或时尚等休闲行业的标准更多地与市场份额的损失有关，而不是直接的、限制生命的影响。评估错误的影响，例如受影响的人数或失败的经济后果，可以帮助您做出合理的选择，以分配您的人力和财力资源以进行缓解。

## Designing against disasters

## 防灾设计

When designing software, risk reduction and harm mitigation patterns can reduce the frequency and severity of negative events. The goal of risk reduction is to prevent negative events and make systems harder to break. Vaccines, antilock brakes, and railroad crossing arms are everyday examples of risk reduction. Harm mitigation strategies acknowledge that negative events will still happen and try to make their impact as mild as possible. Seatbelts, antiviral medications, and cut-proof gloves are examples of harm mitigation.

在设计软件时，降低风险和减轻危害的模式可以降低负面事件的频率和严重程度。降低风险的目标是防止负面事件并使系统更难被破坏。疫苗、防抱死制动器和铁路交叉臂是降低风险的日常例子。减轻危害的策略承认负面事件仍然会发生，并尽量使它们的影响尽可能温和。安全带、抗病毒药物和防割手套是减轻伤害的例子。

In software development, risk reduction patterns might include restricting access to a system, abstractions to manage changes, approval workflows, automated acceptance testing, and pair programming. However, any strategy that hardens a system to make it less risky might also make it less flexible. A perfectly optimized system is often fragile because it has no built-in expectation of failure or redundancy: If something goes wrong, there’s little or no room for recovery. For example, if a subway system has exactly the capacity it needs to accommodate peak commuter traffic, a single out-of-service train can cause delays and overcrowding throughout the system. This is where harm mitigation comes in.

在软件开发中，风险降低模式可能包括限制对系统的访问、管理变更的抽象、批准工作流、自动化验收测试和结对编程。然而，任何强化系统以降低风险的策略也可能使其不那么灵活。一个完美优化的系统通常是脆弱的，因为它没有内置的故障或冗余预期：如果出现问题，几乎没有或没有恢复空间。例如，如果地铁系统恰好具有容纳高峰通勤交通所需的容量，那么一列停运的列车可能会导致整个系统的延误和过度拥挤。这就是减轻伤害的用武之地。

Harm mitigation patterns in software development include circuit breakers to remove problem processes or inputs, the isolation of suspect parts of a system, having many control points or breakpoints (such as on microservices), and the ability to rapidly roll back recently implemented changes. While adding more points of control to a system could also introduce more potential points of failure, it’s often a worthwhile trade-off to be able to isolate and diagnose a subset of the system without taking it all down.

软件开发中的危害缓解模式包括用于消除问题流程或输入的断路器、隔离系统可疑部分、具有许多控制点或断点（例如在微服务上）以及快速回滚最近实施的更改的能力。虽然向系统添加更多控制点也可能引入更多潜在故障点，但能够隔离和诊断系统子集而不将其全部关闭通常是值得的权衡。

One of the ways we express harm mitigation is with error budgets, as described in [O’Reilly’s 2018](https://sre.google/workbook/table-of-contents/). Systems reliability engineering accepts and anticipates that some parts of a system may fail and budgets for these failures. The teams running the system deliberately take parts offline to test responses, workarounds, and restoration. 
我们表达危害减轻的方法之一是错误预算，如 [O'Reilly's 2018](https://sre.google/workbook/table-of-contents/) 中所述。系统可靠性工程接受并预期系统的某些部分可能会发生故障并为这些故障制定预算。运行系统的团队特意将部件离线以测试响应、变通方法和恢复。
This “deliberate disaster” practice was popularized by Netflix and its [Chaos Monkey tool](https://netflix.github.io/chaosmonkey/), released in 2011, but organizations have used disaster recovery practices for much longer. When the practice is a regular part of an organization’s resilience response, it’s sometimes called chaos engineering. Practicing systems failure in real time not only teaches teams how to respond to known known–style failures, where everyone is aware of what's offline, but also unknown or unpredicted failures, because the team has learned to work together to solve a problem and has practice resolving rapidly moving situations.

这种“故意灾难”的做法由 Netflix 及其 2011 年发布的 [Chaos Monkey 工具](https://netflix.github.io/chaosmonkey/) 推广，但组织使用灾难恢复做法的时间要长得多。当实践是组织弹性响应的常规部分时，它有时被称为混沌工程。实时练习系统故障不仅教团队如何应对已知的已知类型的故障，每个人都知道什么是离线的，而且还知道未知或不可预测的故障，因为团队已经学会了合作解决问题并进行了实践解决快速变化的情况。

Preparing for disasters with multiple layers of safety protection is sometimes known as the [Swiss cheese model](https://en.wikipedia.org/wiki/Swiss_cheese_model). Each layer of preparation, prevention, and mitigation is insufficient by itself—but, when stacked together, they can provide good-enough protection from worst-case scenarios.

通过多层安全保护为灾难做准备有时被称为[瑞士奶酪模型](https://en.wikipedia.org/wiki/Swiss_cheese_model)。每一层的准备、预防和缓解本身都是不够的——但是，当它们堆叠在一起时，它们可以提供足够好的保护，以应对最坏的情况。

## Accept imperfection, within limits

##接受不完美，在限度内

The more complex the system, the more likely it is that some part of it is broken. As engineers, we’re at the apex of literally millions of hours of design and engineering time to help us, and our users, with everything from finding our phone to continuously integrating our code. We have to assume that some of that infrastructure, and some of our work, is broken.

系统越复杂，它的某些部分就越有可能被破坏。作为工程师，我们需要花费数百万小时的设计和工程时间来帮助我们和我们的用户，从寻找我们的手机到不断集成我们的代码。我们必须假设某些基础设施和我们的某些工作已损坏。

If we accept these imperfections, we can work toward building resilient systems that can handle a little static and deal with flawed foundations without falling over. We don’t stop working toward perfection just because it’s impossible.

如果我们接受这些缺陷，我们就可以努力构建弹性系统，该系统可以处理一些静态问题并处理有缺陷的基础而不会摔倒。我们不会因为不可能就停止努力追求完美。

1 Dr. Richard Cook touches on these concepts in the first part of his famed 1998 treatise [_How Complex Systems Fail_](https://how.complexsystems.fail/), in which he observes medical errors and the systems that cause them. The theory applies to all sorts of complex systems, especially software.

1 Richard Cook 博士在其 1998 年著名的论文 [_How Complex Systems Fail_](https://how.complexsystems.fail/) 的第一部分中谈到了这些概念，他在其中观察了医疗错误和导致这些错误的系统。该理论适用于各种复杂系统，尤其是软件。

2 Ruby on Rails creator David Heinemeier Hansson expands on this point in “ [FIXME](https://youtu.be/zKyv-IGvgGE),” his keynote talk at [RailsConf 2018](https://railsconf.com/) .

2 Ruby on Rails 创建者 David Heinemeier Hansson 在“[FIXME](https://youtu.be/zKyv-IGvgGE)”中扩展了这一点，这是他在 [RailsConf 2018](https://railsconf.com/) 上的主题演讲.

#### About the author

＃＃＃＃ 关于作者

**Heidi Waterhouse** is a principal transformation advocate for LaunchDarkly, where she works on explaining the intersections of feature control, reliability, and sociotechnical systems. 
**Heidi Waterhouse** 是 LaunchDarkly 的主要转型倡导者，她致力于解释功能控制、可靠性和社会技术系统的交叉点。
