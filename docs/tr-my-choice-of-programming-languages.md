# My Choice of Programming Languages

# 我对编程语言的选择

July 28, 2021



_It's a no holy war post! Any choice of languages cannot be 100% objective, and I'm just sharing my experience here._

_这是一个非圣战帖！任何语言选择都不能100%客观，我只是在这里分享我的经验。_

When I was a kid, I used to spend days tinkering with woodworking tools. I was lucky enough to have a wide set of tools at my disposal. However, there was no one around to give me a hint about what tool to use when. So, I quickly came up with a heuristic: if my fingers and a tool survived an exercise, I've used the right tool; if either the fingers or the tool got damaged, I'd try other tools for the same task until I find the right one. And it worked quite well for me! Since then, I'm an apologist of the idea that every tool is good only for a certain set of tasks.

当我还是个孩子的时候，我常常花几天时间摆弄木工工具。我很幸运有一套广泛的工具可供我使用。但是，周围没有人给我提示何时使用什么工具。所以，我很快想出了一个启发式方法：如果我的手指和工具在练习中幸存下来，我就使用了正确的工具；如果手指或工具损坏，我会尝试其他工具执行相同的任务，直到找到合适的工具。它对我来说效果很好！从那时起，我就认为每个工具都只适用于特定的任务集，我对此表示赞同。

A programming language is yet another kind of tool. When I became a software developer, I adapted my heuristic to the new reality: if, while solving a task using a certain language, I suffer too much (fingers damage) or I need to hack things more often than not (tool damage), it's a wrong choice of a language.

编程语言是另一种工具。当我成为一名软件开发人员时，我调整了我的启发式以适应新的现实：如果在使用某种语言解决任务时，我遭受了太多（手指损坏）或者我需要经常破解（工具损坏），这是一种错误的语言选择。

Since the language is just a tool, my programming toolbox is defined by the tasks I work on the most often. Since 2010, I've worked in many domains, starting from web UI development and ending with writing code for infrastructure components. I find pleasure in being a generalist (jack of all trades), but there is always a pitfall of spreading yourself too thin (master of none). So, for the past few years, I've been trying to limit my sphere of expertise with the **server-side, distributed systems, and infrastructure**. Hence, the following choice of languages.

由于语言只是一种工具，我的编程工具箱由我最常从事的任务定义。自 2010 年以来，我在许多领域工作过，从 Web UI 开发开始，到为基础设施组件编写代码结束。我很高兴成为一名多面手（万事通），但总有一个陷阱，就是把自己分散得太细（无所事事）。所以，在过去的几年里，我一直试图用**服务器端、分布式系统和基础设施**来限制我的专业领域。因此，选择以下语言。

![Language logos](http://iximiuz.com/my-choice-of-programming-languages/kdpv-2000-opt.png)

## Python



The most appealing thing for me in [Python](http://iximiuz.com/en/categories/?category=Python) is the tremendous pace of development this language can provide. It's easy to write code in Python. There is often one clear way to accomplish a task, and you don't have to think much about how to convince the language to do what you need. Instead, you can focus on the business logic. I use Python for **quick prototyping**, for the performance-tolerant **server-side code**, for **ad hoc scripting**, and of course, for **data analysis**. I started my Python journey in 2014 and, without any prior experience with the language, managed to single-handedly develop a service that quickly got thousands of active users a day and [survived the load without a lot of rewriting](http://iximiuz.com/en/posts/save-the-day-with-gevent/).

[Python](http://iximiuz.com/en/categories/?category=Python) 对我来说最吸引人的地方是这种语言可以提供的巨大发展速度。用 Python 编写代码很容易。通常有一种明确的方法可以完成一项任务，而且您不必考虑如何说服语言去做您需要的事情。相反，您可以专注于业务逻辑。我将 Python 用于**快速原型设计**、性能容忍**服务器端代码**、**临时脚本**，当然还有**数据分析**。我在 2014 年开始了我的 Python 之旅，并且在没有任何该语言方面的经验的情况下，设法单枪匹马地开发了一项服务，该服务每天迅速吸引了成千上万的活跃用户，并且[在没有大量重写的情况下幸存下来](http://iximiuz.com/en/posts/save-the-day-with-gevent/)。

## Go 



If I had to describe [Go](http://iximiuz.com/en/categories/?category=Go) using just one word, I'd be torn apart between _pragmatic_ and _boring_. But I guess those are synonyms. Go is the language that could have replaced all other languages in my toolbox. Well, maybe except JavaScript. The development in Go is not as fast as in Python, but you still can have decent productivity. And the resulting code is almost as fast as code in C/C++ or Java. I use Go for **performance-sensitive services**, **infrastructure components**, and **command-line tools**. Static linking makes it especially good for tool development because the distribution is simplified a lot. Another reason to pick Go for me is the availability of some packages. Most of the [Cloud Native](http://iximiuz.com/en/posts/making-sense-out-of-cloud-native-buzz/) projects is written in Go, so it often becomes a [default choice](https://github.com/iximiuz/goimagego) for [certain kind of projects](https://github.com/iximiuz/conman). I have Go in my toolbox since 2016.

如果我不得不用一个词来描述 [Go](http://iximiuz.com/en/categories/?category=Go)，我会在_pragmatic_ 和_boring_ 之间分崩离析。但我想这些是同义词。 Go 是一种可以取代我工具箱中所有其他语言的语言。好吧，也许除了 JavaScript。 Go 的开发速度没有 Python 快，但你仍然可以有不错的生产力。生成的代码几乎与 C/C++ 或 Java 中的代码一样快。我将 Go 用于**性能敏感服务**、**基础设施组件**和**命令行工具**。静态链接使得它特别适合工具开发，因为分发被简化了很多。为我选择 Go 的另一个原因是某些软件包的可用性。大部分 [云原生](http://iximiuz.com/en/posts/making-sense-out-of-cloud-native-buzz/)项目都是用Go写的，所以经常成为 [默认选择](https://github.com/iximiuz/goimagego) 用于[某些类型的项目](https://github.com/iximiuz/conman)。自 2016 年以来，我的工具箱中就有 Go。

## Rust



Unlike Python, you may need to fight the compiler to get things done with [Rust](http://iximiuz.com/en/categories/?category=Rust). But if you manage to compile your code, the result is guaranteed to be safe (well, unless you abuse `unsafe` blocks). And most of the time, the result is also quite performant. I personally find this dancing with compiler quite rewarding, but I have to admit that the development pace degrades a lot. That's why I use Rust only when it's really necessary - for **performance-critical tools** or **[systems](https://github.com/iximiuz/shimmy) [programming](https://github.com/iximiuz/reapme)**. I don't consider myself a knowledgeable Rust developer just yet. The language is in my toolbox only since 2020. But I hope with [pq](https://github.com/iximiuz/pq) I'll be spending more time writing Rust this year.

与 Python 不同，您可能需要与编译器抗争才能使用 [Rust](http://iximiuz.com/en/categories/?category=Rust) 完成任务。但是如果你设法编译你的代码，结果肯定是安全的（好吧，除非你滥用 `unsafe` 块)。大多数情况下，结果也非常高效。我个人觉得这种与编译器共舞是非常有益的，但我不得不承认开发速度下降了很多。这就是为什么我只在真正需要时才使用 Rust - **性能关键工具** 或 **[系统](https://github.com/iximiuz/shimmy) [编程](https://github.com/iximiuz/reapme)**。我还不认为自己是一个知识渊博的 Rust 开发人员。该语言自 2020 年起才出现在我的工具箱中。但我希望在 [pq](https://github.com/iximiuz/pq) 今年我会花更多的时间来编写 Rust。

# C

I [don't write code](https://github.com/iximiuz/golife.c) in [C](http://iximiuz.com/en/categories/?category=C). But I find the ability to read C code fluently a must-have in my domain. The Linux Kernel is written in C, and it makes it a native language for the whole family of operating systems. I always strive to **understand how things work one- or two layers of abstraction below my code**. Often it means I need to dig down to _libc_ calls and try to [reproduce snippets from man pages](https://github.com/iximiuz/ptyme) using bare C. Sometimes [reading the libc code itself](https://github.com/iximiuz/popen2) also helps. When I really have to wear a hat of a systems programmer, I [turn to Rust instead](http://iximiuz.com/en/posts/dealing-with-processes-termination-in-Linux/). My acquaintance with C lasts from 2006-2007.

我 [不写代码](https://github.com/iximiuz/golife.c) 在 [C](http://iximiuz.com/en/categories/?category=C)。但我发现流利地阅读 C 代码的能力是我所在领域的必备条件。 Linux 内核是用 C 编写的，这使它成为整个操作系统系列的本地语言。我总是努力**在我的代码下面的一两层抽象中理解事情是如何工作的**。通常这意味着我需要深入研究 _libc_ 调用并尝试 [从手册页复制片段](https://github.com/iximiuz/ptyme) 使用裸 C。有时[阅读 libc 代码本身](https://github.com/iximiuz/popen2) 也有帮助。当我真的不得不戴上系统程序员的帽子时，我 [转而使用 Rust](http://iximiuz.com/en/posts/dealing-with-processes-termination-in-Linux/)。我与 C 的相识是从 2006 年到 2007 年。

## JavaScript

To some extent, I use JavaScript since 2011. Most of the time, it's about **making UIs** (web, desktop, mobile). For a few years, I've been trying to avoid such kinds of tasks and focus more on the server-side and systems programming. But I find it extremely useful to have such a tool in the toolbox. For instance, when I needed to [**visualize some data**](http://iximiuz.com/en/posts/pq/) recently, I quickly hacked a simple [HTML page](https://github.com/iximiuz/pq/blob/be4751177f014dd60e3144c3170b17f8e8c5e0fc/graph.html) to plot my datasets with [chart.js](https://www.chartjs.org/). I'm not even trying to keep up with all the modern web and UI frameworks, although I've written a few applications in React 🙈

在某种程度上，我从 2011 年开始使用 JavaScript。大多数时候，它是关于 **制作 UI**（网络、桌面、移动）。几年来，我一直在努力避免此类任务，而是更多地关注服务器端和系统编程。但我发现在工具箱中有这样一个工具非常有用。例如，最近我需要[**可视化一些数据**](http://iximiuz.com/en/posts/pq/)时，我很快就黑了一个简单的[HTML页面](https://github.com)。用[chart.js](https://www.chartjs.org/)绘制我的数据集。我什至不想跟上所有现代 Web 和 UI 框架，尽管我已经用 React 编写了一些应用程序🙈

## Other languages 

## 其他语言

I used to do some freelancing in **C++** and **Delphi**, but it ended a long, long time ago. I still use C++ sporadically to solve [Leetcode and HackerRank problems](https://twitter.com/iximiuz/status/1404545363730743303) but only when my Python solution times out. Rewriting it as-is in C++ rarely helps, though. Leaving the university-time freelancing aside, I started my professional career as a **PHP** developer. But I quickly found the language quite limiting. It's well-tailored for web development, but I don't see it as universal as Python or Go. Hence, no fit for my toolbox. I also spent a few years writing services in **Node.js** with **JavaScript** and **TypeScript**. Node.js experience definitely [broadened my understanding of asynchronous programming](http://iximiuz.com/en/posts/explain-event-loop-in-100-lines-of-code/). And I find TypeScript just lovely. But I try to limit the use of these tools only for UI making. In the past two years, I also had a chance to write some **Perl** and **Java** for high-loaded production services. Perl is a mind-blowing language with a [tremendous impact](https://twitter.com/iximiuz/status/1161980017796161538) on many of the mainstream programming languages. And the modern Java is also quite decent. But again, there is just very little fit for them, given my daily tasks.

我曾经做过一些 **C++** 和 **Delphi** 的自由职业，但它在很久很久以前就结束了。我仍然偶尔使用 C++ 来解决 [Leetcode 和 HackerRank 问题](https://twitter.com/iximiuz/status/1404545363730743303)，但只有在我的 Python 解决方案超时时才会使用。但是，在 C++ 中按原样重写它很少有帮助。撇开大学时代的自由职业者不谈，我开始了我作为 **PHP** 开发人员的职业生涯。但我很快发现这种语言非常有限。它非常适合 Web 开发，但我认为它不像 Python 或 Go 那样通用。因此，不适合我的工具箱。我还花了几年时间使用 **JavaScript** 和 **TypeScript** 在 **Node.js** 中编写服务。 Node.js 经验绝对[拓宽了我对异步编程的理解](http://iximiuz.com/en/posts/explain-event-loop-in-100-lines-of-code/)。我发现 TypeScript 很可爱。但我尽量限制这些工具仅用于 UI 制作。在过去的两年里，我也有机会为高负载的生产服务编写一些 **Perl** 和 **Java**。 Perl 是一种令人兴奋的语言，对许多主流编程语言具有 [巨大影响](https://twitter.com/iximiuz/status/1161980017796161538)。而且现代Java也相当不错。但同样，考虑到我的日常任务，他们几乎不适合。

## Looking for a mentor?

## 寻找导师？

If you are at the beginning of your programming career or having a hard time understanding some basic concepts in one of my areas of expertise, I'd be happy to help. Just [drop me a message](http://iximiuz.com/en/about/), and we'll see how to go about it.

如果您正处于编程生涯的开始阶段，或者在我的专业领域之一中很难理解一些基本概念，我很乐意提供帮助。只需[给我留言](http://iximiuz.com/en/about/)，我们会看到如何去做。

### Read also

### 也请阅读

[My 10 Years of Programming Experience](http://iximiuz.com/en/posts/my-10-years-of-programming-experience/)

 [我的 10 年编程经验](http://iximiuz.com/en/posts/my-10-years-of-programming-experience/)

[programming,](javascript: void 0) [python,](javascript: void 0) [golang,](javascript: void 0) [rust,](javascript: void 0) [c,](javascript: void 0) [javascript](javascript: void 0)

[编程,](javascript: void 0) [python,](javascript: void 0) [golang,](javascript: void 0) [rust,](javascript: void 0) [c,](javascript: void 0) [javascript](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

