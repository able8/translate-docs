# The Zen of Go

# Go 之禅

Ten engineering values for writing simple, readable, maintainable Go code. Presented at [GopherCon Israel 2020](http://gophercon.org.il).

编写简单、可读、可维护的 Go 代码的十个工程价值。在 [GopherCon Israel 2020](http://gophercon.org.il) 上发表。

- Each package fulfils a single purpose

- 每个包裹都满足一个目的

  A well designed Go package provides a single idea, a set of related behaviours. A good Go package starts by choosing a good name. Think of your package’s name as an elevator pitch to describe what it provides, using just one word.

一个设计良好的 Go 包提供了一个单一的想法，一组相关的行为。一个好的 Go 包从选择一个好名字开始。将您的包裹的名称想象成一个电梯宣传单来描述它提供的内容，只需一个词。

- Handle errors explicitly

- 明确处理错误

  Robust programs are composed from pieces that handle the failure cases before they pat themselves on the back. The verbosity of `if err != nil { return err }` is outweighed by the value of deliberately handling each failure condition at the point at which they occur. Panic and recover are not exceptions, they aren’t intended to be used that way.

健壮的程序由处理失败案例的部分组成，然后他们会拍拍自己的背。 `if err != nil { return err }` 的冗长性被在它们发生的点故意处理每个失败条件的价值所抵消。恐慌和恢复也不例外，它们不打算以这种方式使用。

- Return early rather than nesting deeply

- 尽早返回而不是深入嵌套

  Every time you indent you add another precondition to the  programmer’s stack consuming one of the 7 ±2 slots in their short term  memory. Avoid control flow that requires deep indentation. Rather than nesting deeply, keep the success path to the left using  guard clauses.

每次缩进时，都会向程序员的堆栈添加另一个先决条件，占用他们短期记忆中的 7 ±2 个插槽之一。避免需要深度缩进的控制流。与其深入嵌套，不如使用保护子句将成功路径保持在左侧。

- Leave concurrency to the caller

- 将并发留给调用者

  Let the caller choose if they want to run your library or function  asynchronously, don’t force it on them. If your library uses concurrency it should do so transparently.

让调用者选择是否要异步运行您的库或函数，不要强加于他们。如果你的库使用并发，它应该透明地这样做。

- Before you launch a goroutine, know when it will stop

- 在你启动一个 goroutine 之前，知道它什么时候会停止

  Goroutines own resources; locks, variables, memory, etc. The sure fire way to free those resources is to stop the owning goroutine.

Goroutines 拥有资源；锁、变量、内存等。释放这些资源的可靠方法是停止拥有 goroutine。

- Avoid package level state

- 避免包级状态

  Seek to be explicit, reduce coupling, and spooky action at a distance by providing the dependencies a type needs as fields on that type  rather than using package variables.

通过提供类型所需的依赖项作为该类型的字段而不是使用包变量，力求明确、减少耦合和远距离的诡异操作。

- Simplicity matters

- 简单很重要

  Simplicity is not a synonym for unsophisticated. Simple doesn’t mean crude, it means *readable* and *maintainable*. When it is possible to choose, defer to the simpler solution.

简单并不是简单的同义词。简单并不意味着粗糙，它意味着*可读*和*可维护*。如果可以选择，请遵循更简单的解决方案。

- Write tests to lock in the behaviour of your package’s API

- 编写测试以锁定包的 API 的行为

  Test first or test later, if you shoot for 100% test coverage or are  happy with less, regardless your package’s API is your contract with its users. Tests are the guarantees that those contracts are written in. Make sure you test for the behaviour that users can observe and rely on.

先测试或后测试，如果你争取 100% 的测试覆盖率或对更少的测试感到满意，不管你的包的 API 是你与用户的合同。测试是写入这些合约的保证。确保您测试用户可以观察和依赖的行为。

- If you think it’s slow, first prove it with a benchmark

- 如果你认为它很慢，首先用一个基准来证明它

  So many crimes against maintainability are committed in the name of performance. Optimisation tears down abstractions, exposes internals, and couples tightly. If you’re choosing to shoulder that cost, ensure it is done for good reason.

以性能的名义犯下了如此多的破坏可维护性的罪行。优化拆除抽象，暴露内部，并紧密耦合。如果您选择承担这笔费用，请确保这样做是有充分理由的。

- Moderation is a virtue

- 适度是一种美德

  Use goroutines, channels, locks, interfaces, embedding, in moderation.

适度使用 goroutine、通道、锁、接口、嵌入。

- Maintainability counts

- 可维护性很重要

  Clarity, readability, simplicity, are all aspects of maintainability. Can the thing you worked hard to build be maintained after you’re gone? What can you do today to make it easier for those that come after you?

清晰、可读性、简单性都是可维护性的各个方面。你辛苦建立的东西在你离开后还能维持吗？你今天可以做些什么来让你之后的人更容易？

Last updated 2020-02-04 08:26:39 UTC 

最后更新时间 2020-02-04 08:26:39 UTC

