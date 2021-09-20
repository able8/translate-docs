# Why Go and not Rust?

# 为什么 Go 而不是 Rust？

September 10, 2019 • 11 min read • by **Loris Cro**



What's the role of Go in a universe where Rust exists?

在 Rust 存在的宇宙中，Go 扮演什么角色？

Imagine you’re a developer who mainly works with Go. You go to an event and, while chatting with some people, you decide to share with them the news that you wrote a small tool that does _something_. You claim that since you wrote it in Go, it’s fairly fast, it’s a single binary, etc. The group seems pleased with your recount and you start feeling good, but then you notice a stranger approaching from behind. A bone-chilling wind blows and you hear: “Why Go and not Rust?”

假设您是一名主要使用 Go 的开发人员。你去参加一个活动，在和一些人聊天时，你决定与他们分享你写了一个可以做_某事_的小工具的消息。你声称自从你用 Go 编写它以来，它相当快，它是一个单一的二进制文件，等等。小组似乎对你的叙述很满意，你开始感觉良好，但随后你注意到一个陌生人从后面接近。一阵刺骨的风吹来，你听到：“为什么Go而不是Rust？”

You start feeling less good. Well, you could answer that Go is what you know, so that’s what you used to solve your problem, but that’s probably not going to be a satisfactory answer. You were pridefully talking about how fast your tool was in the first place, and it’s obvious that the stranger is going to counter your simplistic excuse with all the great benefits that Rust brings over Go.

你开始感觉不太好。好吧，你可以回答 Go 是你所知道的，所以这就是你用来解决问题的方法，但这可能不会是一个令人满意的答案。你一开始就自豪地谈论你的工具有多快，很明显陌生人会用 Rust 为 Go 带来的所有巨大好处来反驳你的简单借口。

You start feeling bad. Why did you choose to learn Go in the first place? You were told that Go is fast and that it has great concurrency primitives, and now Rust comes along and everybody is saying that Rust is better in every aspect. Were they lying before or are they lying now? While there is no single language to rule them all, you know that it’s still possible to make bad choices and end up in a technological cul de sac. After all, you did choose Go over _that_ other language a few years ago and you were pretty pleased with the routine of joining circles to ask “Why _that_ and not Go?”

你开始感觉不好。你当初为什么选择学习 Go  ？你被告知 Go 很快，并且它有很好的并发原语，现在 Rust 出现了，每个人都说 Rust 在各个方面都更好。他们以前撒谎还是现在撒谎？虽然没有一种语言可以统治所有这些，但您知道仍有可能做出错误的选择并最终陷入技术死胡同。毕竟，几年前你确实选择了 Go over _that_ 其他语言，而且你对加入圈子问“为什么 _that_ 而不是 Go？”的惯例感到非常满意。

* * *

* * *

While the story above is 100% the result of my imagination, it’s no secret that the Rust fandom has a few _overexcited_ members who feel compelled to enlighten every lost soul about the virtues of the Crab-God. This isn’t really Rust’s fault, every successful project will have misbehaving followers, it’s inevitable. While everyone has to deal with these people, I feel that Go developers are particularly susceptible to their behavior because of how much Rust's and Go's [messaging](https://courses.lumenlearning.com/ivytech-mktg101-master/chapter/reading-defining-the-message/) overlap.

虽然上面的故事 100% 是我想象的结果，但毫无疑问，Rust 粉丝圈有一些_过度兴奋的_成员，他们觉得有必要让每个迷失的灵魂了解蟹神的美德。这并不是 Rust 的错，每个成功的项目都会有行为不端的追随者，这是不可避免的。虽然每个人都必须与这些人打交道，但我觉得 Go 开发人员特别容易受到他们的行为的影响，因为 Rust 和 Go 的 [消息传递](https://courses.lumenlearning.com/ivytech-mktg101-master/chapter/reading-定义消息/)重叠。

Go is fast, but Rust is faster.

Go 很快，但 Rust 更快。

Go has an efficient garbage collector, but Rust has static memory management.

Go 有一个高效的垃圾收集器，但 Rust 有静态内存管理。

Go has great concurrency support, but Rust has provably-correct concurrency.

Go 有很好的并发支持，但 Rust 有可证明正确的并发。

Go has interfaces, but Rust has traits and other zero-cost abstractions.

Go 有接口，但 Rust 有特征和其他零成本的抽象。

If you’re a Go developer you might feel a bit cheated. In contrast, Python developers are not particularly fazed by Rust. They know that Python is in many ways slow and inefficient, and they’re fine with that because they know Python’s role: make the code easy to write and offload to C when performance matters.

如果您是 Go 开发人员，您可能会觉得有点受骗。相比之下，Python 开发人员对 Rust 并不特别感兴趣。他们知道 Python 在很多方面都很慢且效率低下，他们对此很满意，因为他们知道 Python 的作用：在性能很重要时使代码易于编写并卸载到 C。

What about Go?

Go？

## Go is very good for writing services

## Go 非常适合编写服务

Go was created at Google to solve Google problems, which mostly involves dealing with networked services. Go’s concurrency model is a good fit for server-side applications that must primarily handle multiple independent requests, rather than participate in complex result-passing schemes. This is one reason why you’re given `go` and not `await`.

Go 是由谷歌创建的，旨在解决谷歌的问题，主要涉及处理网络服务。 Go 的并发模型非常适合必须主要处理多个独立请求而不是参与复杂的结果传递方案的服务器端应用程序。这就是为什么给你 `go` 而不是 `await` 的原因之一。

Go has great support for HTTP and related protocols and it doesn’t take long to write a satisfactory web service. In my personal projects, Go proved to be a good alternative to Node.js, especially in situations where I wanted to pin down the interfaces between different components more explicitly than you would do while writing idiomatic JavaScript.

Go 对 HTTP 和相关协议有很好的支持，编写一个令人满意的 Web 服务不需要很长时间。在我的个人项目中，Go 被证明是 Node.js 的一个很好的替代品，尤其是在我想比编写惯用 JavaScript 时更明确地确定不同组件之间的接口的情况下。

On top of that, it has great tooling to diagnose concurrency and performance problems, and cross-compilation makes deploying on whichever platform a breeze.

最重要的是，它有很好的工具来诊断并发和性能问题，交叉编译使得在任何平台上部署都变得轻而易举。

## Go is unapologetically simple 

## Go 非常简单

Go takes great pride in offering a limited set of features built-in the language. This makes Go easy to learn and, even more importantly, it ensures that Go projects remain understandable even when growing in size. The creators of Go like to call it a “boring” language. While we could debate whether the language could use one or two extra things, the idea of forcing people to “do more with less” has proven to be very successful.

Go 以提供该语言内置的一组有限功能而自豪。这使得 Go 易于学习，更重要的是，它确保 Go 项目即使在规模增长时也能保持可理解性。 Go 的创造者喜欢称其为“无聊”的语言。虽然我们可以争论语言是否可以使用一两个额外的东西，但强迫人们“事半功倍”的想法已被证明是非常成功的。

Rust can indeed be as good as Go at doing web services, or even better, but it really cannot beat Go in terms of simplicity. And it’s not just simplicity, Go is also strict about things that other languages are usually more lax about. Go doesn’t want unused variables or imports, files belonging to different packages in the same directory, etc. It even used to complain about projects saved outside of `GOPATH` (thankfully that’s not the case anymore).

Rust 在做 Web 服务方面确实可以和 Go 一样好，甚至更好，但它在简单性方面确实无法击败 Go。不仅仅是简单，Go 对其他语言通常更宽松的东西也很严格。 Go 不想要未使用的变量或导入、属于同一目录中不同包的文件等。它甚至曾经抱怨保存在 `GOPATH` 之外的项目（谢天谢地，情况不再如此）。

Go also doesn’t want any “fingerprints” in the code, so it enforces a single, universal style via `go fmt`.

Go 也不希望代码中有任何“指纹”，因此它通过 `go fmt` 强制执行单一、通用的样式。

In truth, none of these things alone is particularly impressive, but they do describe the mindset that Go wants to impose. Many don’t like it but, in my opinion, it’s a killer feature for some types of development, like enterprise software.

事实上，单独这些东西都不是特别令人印象深刻，但它们确实描述了 Go 想要强加的思维方式。许多人不喜欢它，但在我看来，它是某些类型的开发（例如企业软件）的杀手级功能。

##  Go is great for enterprise software

## Go 非常适合企业软件

As I already mentioned, Go was created to solve Google problems, and Google problems are definitely enterprise-scale problems. Whether this was the direct aim of the creators or just the natural result of using it at a Big Corp, Go is indubitably an amazing breath of fresh air in enterprise software development.

正如我已经提到的，Go 是为了解决 Google 问题而创建的，而 Google 问题绝对是企业级问题。无论这是创建者的直接目标，还是只是在大公司使用它的自然结果，Go 无疑是企业软件开发中令人惊叹的新鲜空气。

If you have experience in writing enterprise software and tried Go, you probably know what I mean. Here’s a quick summary for everyone else.

如果您有编写企业软件的经验并尝试过 Go，那么您可能知道我的意思。这是对其他人的快速总结。

Enterprise development is a very weird beast compared to other kinds of development. If you’ve never done it, you might ask yourself “What the hell is even enterprise software?” I did that kind of development as a consultant for a bit, so here’s my take on the subject.

与其他类型的开发相比，企业开发是一种非常奇怪的野兽。如果您从未这样做过，您可能会问自己“企业软件到底是什么？”我以顾问的身份做过这种开发，所以这是我对这个主题的看法。

####  Enterprise software development is about scale

#### 企业软件开发是关于规模的

Not in terms of total number of users, or amount of data. Often that’s the case too, but the defining characteristics are scale of **scope** and **process**.

不是根据用户总数或数据量。通常情况也是如此，但定义特征是**范围**和**过程**的规模。

**Enterprise software always has a big scope.** The domain can be big and wide, or narrow but stupidly complex. Sometimes it’s both. When creating software to model such domains, normal programming wisdom falls incredibly short because non-technological concerns out-weight most technological ones.

**企业软件的范围总是很大。** 领域可以很大很宽，也可以很窄但非常复杂。有时两者兼而有之。在创建软件来为这些领域建模时，正常的编程智慧非常缺乏，因为非技术问题超过了大多数技术问题。

**To unravel complex domains you need a well-structured process.** You need domain experts, analysts, and a mechanism that lets stakeholders gauge how much progress is being made. It’s also often the case that you, the technologist, don’t know the domain very well. Stakeholders and domain experts often won’t understand technology very well either.

**要解开复杂的领域，您需要一个结构良好的流程。**您需要领域专家、分析师和一种机制，让利益相关者可以衡量取得了多少进展。通常情况下，您，技术专家，不太了解该领域。利益相关者和领域专家通常也不太了解技术。

This pushes further down technological concerns such as efficiency, **and even correctness.** Don’t get me wrong, the business does care about correctness, but they have a different definition for it. When you’re thinking about algorithmic correctness, they are thinking about a reconciliation back-office for the operations team they keep in a country where labor is cheap.

这进一步降低了技术问题，例如效率、**甚至正确性。**不要误会我的意思，业务确实关心正确性，但他们对此有不同的定义。当您考虑算法的正确性时，他们正在考虑为他们在劳动力便宜的国家/地区保留的运营团队提供对账后台。

Because of the environment resulting from this general premise, a few well-known enterprise development “quirks” have emerged over time. I’ll name three relevant to my point.

由于这个大前提所产生的环境，随着时间的推移，出现了一些著名的企业发展“怪癖”。我将说出与我的观点相关的三个。

1. **There are a lot of junior developers** who learn on the job how to program and most are not lucky enough to find a job that will truly teach them anything. In some places, after you’re hired, you are stationed for one week in front of PluralSight and then you’re considered ready to go. 
2. **有很多初级开发人员**在工作中学习如何编程，但大多数人都没有幸运地找到一份能真正教会他们任何东西的工作。在某些地方，在您被录用后，您会在 PluralSight 前驻扎一周，然后您就被认为准备好了。
2. **Software projects quickly become huge and complex for all the wrong reasons.** Big projects take time to build and people (or whole teams) will come and go in the meantime. Constant refactoring is never an option so each will leave behind a lot of code written with very varying levels of quality. Multiple teams working in parallel will produce redundant code. The domain shifts over time, inevitably invalidating older assumptions and consequently causing abstractions to leak. The more sophisticated an abstraction is, the higher the risk that it will become a problem when business comes back with a serious change request.
4. **软件项目由于各种错误的原因而迅速变得庞大和复杂。** 大型项目需要时间来构建，同时人员（或整个团队）会来来去去。不断的重构从来都不是一种选择，所以每次重构都会留下大量质量参差不齐的代码。多个团队并行工作将产生冗余代码。领域随着时间的推移而变化，不可避免地使旧的假设无效，从而导致抽象泄漏。抽象越复杂，当业务带着严重的变更请求返回时，它成为问题的风险就越大。
5. **The toolchain is very often lousy and/or dated.** This is pretty much the inevitable result of all that I described so far. Huge amounts of old code tie you down to a specific toolset, junior developers will learn the status-quo at best, and the people on top (managers and stakeholders) are extremely often not prepared for making technological decisions based on first-hand experience, the general nature of the endeavor makes them risk-averse, causing them to mainly mimic what other successful players in their space do or, more precisely, what analysts _claim_ other successful players do.
6. **工具链通常很糟糕和/或过时。** 这几乎是我到目前为止所描述的所有结果的必然结果。大量的旧代码将您束缚在特定的工具集上，初级开发人员最多只能了解现状，而高层人员（经理和利益相关者）通常没有准备好根据第一手经验做出技术决策，这种努力的一般性质使他们规避风险，导致他们主要模仿他们领域中其他成功的参与者所做的事情，或者更准确地说，是分析师_声称_其他成功参与者所做的事情。

####  Go is about suppressing complexity at scale

#### Go 是关于大规模抑制复杂性

Go makes teams more successful, partially by giving them more than what they would get with other ecosystems, and partially by taking tools away from them, to prevent common pitfalls.

Go 使团队更成功，部分是通过为他们提供比其他生态系统更多的东西，部分是通过从他们那里拿走工具，以防止常见的陷阱。

**Go is much easier to learn than Java or C#.** A faster ramp-up is generally good, but it becomes fundamental when the project is lagging behind, the deadline is approaching, and management inevitably resorts to hiring more people, hoping (in vain) to speed things up.

**Go 比 Java 或 C# 更容易学习。** 更快的提升通常是好的，但当项目落后，截止日期临近，管理层不可避免地诉诸于雇用更多人，希望（徒劳地）加快速度。

**The Go community regards as anti-patterns many abstractions regularly employed by Java / C#**, like IoC containers, or OOP inheritance, for example. There are only two levels of visibility for variables, and the only concurrency model is CSP. It’s way harder to fall into incomprehensible pitfalls when writing in Go than it is in Java / C#.

**Go 社区将 Java/C# 经常使用的许多抽象视为反模式**，例如 IoC 容器或 OOP 继承。变量的可见性只有两个级别，唯一的并发模型是 CSP。用 Go 编写比用 Java / C# 编写时更难陷入难以理解的陷阱。

**The Go compiler is fast.** Which means that running tests is going to be faster in general, that deployments will take less time, increasing the overall productivity.

**Go 编译器速度很快。** 这意味着运行测试通常会更快，部署将花费更少的时间，从而提高整体生产力。

With Go, it’s easier as a junior developer to be more productive, and harder as a mid-level developer to introduce brittle abstractions that will cause problems down the line.

使用 Go，初级开发人员更容易提高工作效率，而中级开发人员更难引入脆弱的抽象，这些抽象会导致问题。

For these reasons, Rust is less compelling than Go for enterprise software development. That doesn’t mean that Go is perfect, or that it’s not true that Rust has some advantages over Go. Rust has many advantages over Go, but more than anything, I would say that it’s the common perception of Go that is wrong.

由于这些原因，Rust 在企业软件开发方面不如 Go 有吸引力。这并不意味着 Go 是完美的，或者说 Rust 比 Go 有一些优势是不正确的。 Rust 比 Go 有很多优势，但最重要的是，我认为 Go 的普遍看法是错误的。

Go is not blazing-fast. Go is not super memory-efficient. Go doesn’t have the absolutely best concurrency model. 

Go 并不是很快。 Go 的内存效率并不高。 Go 没有绝对最好的并发模型。

Go is faster than Java / C#, more memory-efficient than Java / C#, and definitely has better concurrency than Java / C#. **Go is simple so that all of this can hold true when confronting the average Go program with the average Java / C# program.** It doesn’t matter whether Go is truly faster than C# or Java in an absolute sense. The average Java / C# application will be very different than the best theoretical program, and the amount of foot guns in those languages is huge compared to Go. If you want an example, [take a look at this talk on C# concurrency](https://www.youtube.com/watch?v=J0mcYVxJEl0), it's incredible in my opinion how the straightforward use of `await` is * *never** correct. Imagine how broken must be the average asynchronous C# application. Actually, ASP.NET applications deadlocking for no apparent reason are not uncommon.

Go 比 Java / C# 更快，比 Java / C# 更节省内存，并且绝对比 Java / C# 具有更好的并发性。 **Go 很简单，所以当面对普通的 Go 程序和普通的 Java / C# 程序时，所有这些都可以成立。 ** Go 是否真正比 C# 或 Java 快并不重要。普通的 Java / C# 应用程序将与最好的理论程序大不相同，并且与 Go 相比，这些语言中的步数是巨大的。如果你想要一个例子，[看看这个关于 C# 并发的演讲](https://www.youtube.com/watch?v=J0mcYVxJEl0)，在我看来，`await` 的直接使用是多么令人难以置信* *从不**正确。想象一下普通的异步 C# 应用程序是多么的破碎。实际上，ASP.NET 应用程序无缘无故死锁的情况并不少见。

## In conclusion

##  综上所述

Go is a _better_ Java / C#, while Rust is not. The clarity that Go can bring to enterprise software development is without a doubt much more valuable than removing garbage collection at the cost of worsening the overall productivity.

Go 是_更好的_Java / C#，而 Rust 不是。 Go 可以为企业软件开发带来的清晰度无疑比以降低整体生产力为代价去除垃圾收集更有价值。

Rust is a _better_ C++, and even if you occasionally hear that Go is a _better_ C, well, that’s just not the case. No language with a built-in garbage collector and runtime can be considered a C. And don't be mistaken, Rust is a C++, not a C. [If you want a better C, take a look at Zig](http://kristoff.it/blog/what-is-zig-comptime).

Rust 是一种更好的 C++，即使你偶尔听说 Go 是一种更好的 C，好吧，事实并非如此。没有内置垃圾收集器和运行时的语言可以被认为是 C。不要误会，Rust 是 C++，而不是 C。[如果你想要更好的 C，看看 Zig](http://kristoff.it/blog/what-is-zig-comptime)。

Lastly, going back to our story, not all “Why not Rust” questions should be interpreted like in the example above. Sometimes the chilling wind is just in your head and the people asking the dreaded question just want to know your opinion. Let’s avoid tying our identity to a single language and embrace practicality first and foremost. Tribal names like Rustacean or Gopher should be avoided, as they are inherently a marketing tool for inducing stronger branding. 

最后，回到我们的故事，并非所有“为什么不 Rust”的问题都应该像上面的例子那样解释。有时，令人不寒而栗的风就在你的脑海中，而问这个可怕问题的人只是想知道你的意见。让我们避免将我们的身份与单一语言联系起来，首先要拥抱实用性。应该避免使用像 Rustacean 或 Gopher 这样的部落名称，因为它们本质上是一种用于诱导更强大品牌的营销工具。

