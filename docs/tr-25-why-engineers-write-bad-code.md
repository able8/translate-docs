# Why do (some) engineers write bad code?

# 为什么（一些）工程师会写出糟糕的代码？

07 Jun 2017 From: https://awmanoj.github.io/tech/2017/06/07/why-engineers-write-bad-code/

I’ve been reading System Performance Tuning, 2nd Edition recently. Though the content is a bit dated as a reference (was written while Linux kernel 2.4 was coming up) but this is still a fantastic  book on refreshing the workings of the components of a computer and  knowing the tweaks for solaris and linux.

我最近一直在阅读系统性能调优第 2 版。虽然内容作为参考有点过时（是在 Linux 内核 2.4 即将推出时编写的），但这仍然是一本关于刷新计算机组件的工作原理以及了解对 solaris 和 linux 的调整的很棒的书。

This post though, is focussed on a topic that is touched towards the end of the book. Why do we end up writing bad code?

不过，这篇文章的重点是本书结尾处触及的主题。为什么我们最终会写出糟糕的代码？

Bad code could mean many things - bad design or architecture  decision, badly written code that is harder to maintain (duplication,  tightly coupled components, less modular etc.), code that is bad at  performance (time or space) and so on. The book talks about following  reasons for badly written code (most of it is verbatim with few  modifications here and there):

糟糕的代码可能意味着很多事情——糟糕的设计或架构决策，糟糕的代码难以维护（重复、紧密耦合的组件、较少的模块化等）、性能（时间或空间）糟糕的代码等等。这本书谈到了编写糟糕代码的以下原因（大部分是逐字逐句的，这里和那里的修改很少）：

- **Writing bad code is much faster than writing good code.** Good software engineering practices often trade an increase in initial  development time for a decrease in the time required to maintain the  software by eliminating bugs in the design phase or for a decrease in  application runtime by integrating performance analysis into the  development process. When faced with a deadline and irate management,  it’s often easier to just get it done rather than do it well.

- **编写糟糕的代码比编写好的代码要快得多。** 良好的软件工程实践通常会以增加初始开发时间来减少维护软件所需的时间（通过消除设计阶段的错误或减少维护软件所需的时间）在应用程序运行时通过将性能分析集成到开发过程中。当面临截止日期和愤怒的管理时，完成它通常比做好它更容易。

> TIP: Learn to say No to “can we do this quickly” in favor of better thought of and more solid code or design. You will thank yourself for  it more often than not.

> 提示：学会对“我们能不能快点做这件事”说不，有利于更好地思考和更可靠的代码或设计。你会经常为此感谢自己。

- **Implementing good algorithms is technically challenging.** This coupled with management’s desire for a product yesterday and the  time to implement such an algorithm, often the good algorithms lose out.

- **实施好的算法在技术上具有挑战性。** 再加上管理层昨天对产品的渴望以及实施这种算法的时间，好的算法往往会失败。

> TIP: Invest in learning about common algorithms & data  structures and weigh upon choices when you’re faced with a problem. Almost always the first solution that strikes one’s mind is not the  best.

> 提示：投资学习常用算法和数据结构，并在遇到问题时权衡选择。几乎总是第一个想到的解决方案并不是最好的。

- **Good software engineers are difficult to find.**  Software engineering is more than writing good source code. It involves  architecture development, soliciting support from other development  organisations, formal code review, ongoing documentation, version  control etc. Concepts like pair programming are just beginning to move  from the academic world into the mainstream and aren’t yet heavily  practised.
- **好的软件工程师很难找到** 软件工程不仅仅是编写好的源代码。它涉及架构开发、寻求其他开发组织的支持、正式的代码审查、持续的文档、版本控制等。 结对编程等概念刚刚开始从学术界进入主流，尚未得到广泛实践。
- **Legacy is hard to maintain.** It’s quite rare that a group of engineers sit down to develop a piece of software from  scratch. More likely, they’re adding onto an existing framework, or  developing the next version of the current product. The mistakes have  already been made, and it’s not time effective (or cost effective) to  throw everything and start over. Solution can be refactoring which makes the existing code more amicable to catching bugs, bottlenecks and  improvements.

- **遗留很难维护** 一群工程师坐下来从头开始开发一个软件是非常罕见的。更有可能的是，他们正在添加现有框架，或开发当前产品的下一个版本。错误已经犯了，扔掉一切重新开始不是时间有效的（或成本有效的）。解决方案可以是重构，使现有代码更容易捕捉错误、瓶颈和改进。

Also, experience wise I find that, bad code gets copied more than  good code (remember ‘it is implemented the same way in that file /  module / project etc.’? at the workplace). While re-writes are hard,  it’s always possible to refactor the code as and when possible to make  it more friendlier to improvements.

此外，根据经验，我发现糟糕的代码比好的代码更容易被复制（还记得“它在那个文件/模块/项目等中以相同的方式实现”？在工作场所）。虽然重写很困难，但总是可以在可能的情况下重构代码，使其对改进更友好。

> TIP: The basic idea here is whenever you write a line of code, it  should be refactored. If you’re rewriting pieces of legacy code, those  should be refactored as well.

> 提示：这里的基本思想是，无论何时编写一行代码，都应该对其进行重构。如果您正在重写遗留代码片段，那么它们也应该被重构。

Good night! 
晚安！

