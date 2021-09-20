# Why unit tests and how to make them work for you

# 为什么要进行单元测试以及如何让它们为你工作

[Here's a link to a video of me chatting about this topic](https://www.youtube.com/watch?v=Kwtit8ZEK7U)

[这是我谈论这个话题的视频的链接](https://www.youtube.com/watch?v=Kwtit8ZEK7U)

If you're not into videos, here's wordy version of it.

如果你不喜欢视频，这里有一个冗长的版本。

## Software

##  软件

The promise of software is that it can change. This is why it is called _soft_ ware, it is malleable compared to hardware. A great engineering team should be an amazing asset to a company, writing systems that can evolve with a business to keep delivering value.

软件的承诺是它可以改变。这就是它被称为 _soft_ ware 的原因，与硬件相比，它具有可塑性。一个伟大的工程团队应该是公司的一项惊人资产，他们编写的系统可以随着业务的发展而不断提供价值。

So why are we so bad at it? How many projects do you hear about that outright fail? Or become "legacy" and have to be entirely re-written (and the re-writes often fail too!)

那为什么我们做得这么差呢？你听说过多少项目彻底失败？或者成为“遗产”并且必须完全重写（并且重写也经常失败！）

How does a software system "fail" anyway? Can't it just be changed until it's correct? That's what we're promised!

无论如何，软件系统是如何“失败”的？不能更改直到正确吗？这就是我们的承诺！

A lot of people are choosing Go to build systems because it has made a number of choices which one hopes will make it more legacy-proof.

很多人选择 Go 来构建系统，因为它已经做出了许多选择，人们希望这些选择能够使其更能抵御遗留问题。

- Compared to my previous life of Scala where [I described how it has enough rope to hang yourself](http://www.quii.dev/Scala_-_Just_enough_rope_to_hang_yourself), Go has only 25 keywords and _a lot_ of systems can be built from the standard library and a few other small libraries. The hope is that with Go you can write code and come back to it in 6 months time and it'll still make sense.
- The tooling in respect to testing, benchmarking, linting & shipping is first class compared to most alternatives.
- The standard library is brilliant.
- Very fast compilation speed for tight feedback loops
- The Go backward compatibility promise. It looks like Go will get generics and other features in the future but the designers have promised that even Go code you wrote 5 years ago will still build. I literally spent weeks upgrading a project from Scala 2.8 to 2.10.

- 与我之前在 Scala 的生活相比 [我描述了它如何有足够的绳子来吊死自己](http://www.quii.dev/Scala_-_Just_enough_rope_to_hang_yourself)，Go 只有 25 个关键字并且可以构建_很多_系统来自标准库和其他一些小型库。希望您可以使用 Go 编写代码并在 6 个月后返回它，它仍然有意义。
- 与大多数替代品相比，测试、基准测试、linting 和运输方面的工具是一流的。
- 标准库很棒。
- 非常快的编译速度，用于紧密的反馈循环
- Go 向后兼容性承诺。看起来 Go 将来会获得泛型和其他特性，但设计者已经承诺，即使是你 5 年前编写的 Go 代码仍然会构建。我真的花了数周时间将一个项目从 Scala 2.8 升级到 2.10。

Even with all these great properties we can still make terrible systems, so we should look to the past and understand lessons in software engineering that apply no matter how shiny (or not) your language is.

即使拥有所有这些伟大的特性，我们仍然可以制造出糟糕的系统，所以我们应该回顾过去并理解软件工程中的课程，无论您的语言多么闪亮（或不闪亮）。

In 1974 a clever software engineer called [Manny Lehman](https://en.wikipedia.org/wiki/Manny_Lehman_%28computer_scientist%29) wrote [Lehman's laws of software evolution](https://en.wikipedia.org/wiki/Lehman%27s_laws_of_software_evolution).

1974 年，一位名叫 [Manny Lehman](https://en.wikipedia.org/wiki/Manny_Lehman_%28computer_scientist%29) 的聪明软件工程师写了 [Lehman 的软件进化定律](https://en.wikipedia.org/wiki/Lehman%27s_laws_of_software_evolution)。

> The laws describe a balance between forces driving new developments on one hand, and forces that slow down progress on the other hand.

> 法律描述了一方面推动新发展的力量与另一方面减缓进步的力量之间的平衡。

These forces seem like important things to understand if we have any hope of not being in an endless cycle of shipping systems that turn into legacy and then get re-written over and over again.

如果我们希望不要陷入无休止的运输系统循环中，这些系统将变成遗留系统，然后一遍又一遍地重写，那么这些力量似乎是重要的事情，需要了解。

## The Law of Continuous Change

## 不断变化的法则

> Any software system used in the real-world must change or become less and less useful in the environment

> 现实世界中使用的任何软件系统都必须改变或在环境中变得越来越没用

It feels obvious that a system _has_ to change or it becomes less useful but how often is this ignored?

很明显，系统_必须_改变或变得不那么有用，但这种情况多久被忽略？

Many teams are incentivised to deliver a project on a particular date and then moved on to the next project. If the software is "lucky" there is at least some kind of hand-off to another set of individuals to maintain it, but they didn't write it of course.

许多团队被激励在特定日期交付项目，然后转移到下一个项目。如果软件是“幸运的”，至少有某种程度的移交给另一组人来维护它，但他们当然没有编写它。

People often concern themselves with trying to pick a framework which will help them "deliver quickly" but not focusing on the longevity of the system in terms of how it needs to evolve.

人们经常关心自己试图选择一个框架来帮助他们“快速交付”，而不是关注系统在如何发展方面的寿命。

Even if you're an incredible software engineer, you will still fall victim to not knowing the future needs of your system. As the business changes some of the brilliant code you wrote is now no longer relevant.

即使您是一位了不起的软件工程师，您仍然会因为不了解系统的未来需求而成为受害者。随着业务的变化，您编写的一些精彩代码现在不再相关。

Lehman was on a roll in the 70s because he gave us another law to chew on.

雷曼兄弟在上世纪 70 年代风生水起，因为他给了我们另一条值得我们深思的法律。

## The Law of Increasing Complexity

## 增加复杂性的法则

> As a system evolves, its complexity increases unless work is done to reduce it

> 随着系统的发展，其复杂性会增加，除非采取措施降低复杂性

What he's saying here is we can't have software teams as blind feature factories, piling more and more features on to software in the hope it will survive in the long run.

他在这里说的是，我们不能让软件团队成为盲目的功能工厂，将越来越多的功能堆积在软件上，希望它能够长期生存。

We **have** to keep managing the complexity of the system as the knowledge of our domain changes.

随着我们领域知识的变化，我们**必须**继续管理系统的复杂性。

## Refactoring

## 重构

There are _many_ facets of software engineering that keeps software malleable, such as:

软件工程的_许多_方面可以保持软件的可塑性，例如：

- Developer empowerment
- Generally "good" code. Sensible separation of concerns, etc etc
- Communication skills
- Architecture
- Observability
- Deployability
- Automated tests
- Feedback loops 

- 开发者赋能
- 通常“好”的代码。关注点的合理分离等
- 沟通技巧
- 建筑学
- 可观察性
- 可部署性
- 自动化测试
- 反馈回路

I am going to focus on refactoring. It's a phrase that gets thrown around a lot "we need to refactor this" - said to a developer on their first day of programming without a second thought.

我将专注于重构。这是一句经常被提及的短语“我们需要重构它”——在编程的第一天就对开发人员说，没有经过深思熟虑。

Where does the phrase come from? How is refactoring just different from writing code?

这句话出自哪里？重构与编写代码有何不同？

I know that I and many others have _thought_ we were doing refactoring but we were mistaken

我知道我和许多其他人都_认为_我们正在重构，但我们错了

[Martin Fowler describes how people are getting it wrong](https://martinfowler.com/bliki/RefactoringMalapropism.html)

[Martin Fowler 描述人们是如何弄错的](https://martinfowler.com/bliki/RefactoringMalapropism.html)

> However the term "refactoring" is often used when it's not appropriate. If somebody talks about a system being broken for a couple of days while they are refactoring, you can be pretty sure they are not refactoring.

> 然而，“重构”一词经常在不合适的情况下使用。如果有人在重构时谈论系统被破坏了几天，您可以非常确定他们不是在重构。

So what is it?

那是什么？

### Factorisation

### 因式分解

When learning maths at school you probably learned about factorisation. Here's a very simple example

在学校学习数学时，您可能了解了因式分解。这是一个非常简单的例子

Calculate `1/2 + 1/4`

计算`1/2 + 1/4`

To do this you _factorise_ the denominators, turning the expression into

要做到这一点，你_分解_分母，把表达式变成

`2/4 + 1/4` which you can then turn into `3/4`.

`2/4 + 1/4` 然后你可以变成 `3/4`。

We can take some important lessons from this. When we _factorise the expression_ we have **not changed the meaning of the expression**. Both of them equal `3/4` but we have made it easier for us to work with; by changing `1/2` to `2/4` it fits into our "domain" easier.

我们可以从中吸取一些重要的教训。当我们_分解表达式_时，我们**没有改变表达式的含义**。它们都等于“3/4”，但我们让我们更容易使用；通过将 `1/2` 更改为 `2/4`，它更容易融入我们的“领域”。

When you refactor your code, you are trying to find ways of making your code easier to understand and "fit" into your current understanding of what the system needs to do. Crucially **you should not be changing behaviour**.

当你重构你的代码时，你试图找到让你的代码更容易理解和“适应”你当前对系统需要做什么的理解的方法。至关重要的是**你不应该改变行为**。

#### An example in Go

#### Go 中的一个例子

Here is a function which greets `name` in a particular `language`

    func Hello(name, language string) string {
    
       if language == "es" {
          return "Hola, " + name
       }
    
       if language == "fr" {
          return "Bonjour, " + name
       }
      
       // imagine dozens more languages
    
       return "Hello, " + name
     }

Having dozens of `if` statements doesn't feel good and we have a duplication of concatenating a language specific greeting with `, ` and the `name.` So I'll refactor the code.

有几十个 `if` 语句感觉不太好，而且我们有重复的将特定语言的问候与 `、` 和 `name` 连接起来的情况。所以我将重构代码。

```
    func Hello(name, language string) string {
           return fmt.Sprintf(
               "%s, %s",
               greeting(language),
               name,
           )
     }
    
     var greetings = map[string]string {
       "es": "Hola",
       "fr": "Bonjour",
       //etc..
     }
    
     func greeting(language string) string {
       greeting, exists := greetings[language]
      
       if exists {
          return greeting
       }
      
       return "Hello"
     }
```


The nature of this refactor isn't actually important, what's important is I haven't changed behaviour.

这种重构的性质实际上并不重要，重要的是我没有改变行为。

When refactoring you can do whatever you like, add interfaces, new types, functions, methods etc. The only rule is you don't change behaviour

重构时，您可以随心所欲，添加接口、新类型、函数、方法等。唯一的规则是不要改变行为

### When refactoring code you must not be changing behaviour

### 重构代码时，您不能改变行为

This is very important. If you are changing behaviour at the same time you are doing _two_ things at once. As software engineers we learn to break systems up into different files/packages/functions/etc because we know trying to understand a big blob of stuff is hard.

这是非常重要的。如果你同时改变行为，你就是在同时做_两件事。作为软件工程师，我们学会将系统分解成不同的文件/包/功能/等，因为我们知道试图理解一大堆东西是很困难的。

We don't want to have to be thinking about lots of things at once because that's when we make mistakes. I've witnessed so many refactoring endeavours fail because the developers are biting off more than they can chew.

我们不想同时考虑很多事情，因为那是我们犯错误的时候。我亲眼目睹了如此多的重构努力都失败了，因为开发人员咬牙切齿。

When I was doing factorisations in maths classes with pen and paper I would have to manually check that I hadn't changed the meaning of the expressions in my head. How do we know we aren't changing behaviour when refactoring when working with code, especially on a system that is non-trivial?

当我在数学课上用笔和纸做因式分解时，我必须手动检查我是否没有改变头脑中表达式的含义。我们如何知道在使用代码进行重构时我们没有改变行为，尤其是在非平凡的系统上？

Those who choose not to write tests will typically be reliant on manual testing. For anything other than a small project this will be a tremendous time-sink and does not scale in the long run.

**In order to safely refactor you need unit tests** because they provide

那些选择不编写测试的人通常会依赖于手动测试。对于小项目以外的任何事情，这将是一个巨大的时间消耗，并且从长远来看不会扩展。

**为了安全地重构你需要单元测试**因为它们提供

- Confidence you can reshape code without worrying about changing behaviour
- Documentation for humans as to how the system should behave
- Much faster and more reliable feedback than manual testing

- 有信心可以重塑代码而不必担心改变行为
- 关于系统应该如何运行的人类文档
- 比手动测试更快、更可靠的反馈

#### An example in Go

#### Go 中的一个例子

A unit test for our `Hello` function could look like this

 我们的“Hello”函数的单元测试可能如下所示

    func TestHello(t *testing.T) {
       got := Hello(“Chris”, es)
       want := "Hola, Chris"
    
       if got != want {
          t.Errorf("got %q want %q", got, want)
       }
     } 

At the command line I can run `go test` and get immediate feedback as to whether my refactoring efforts have altered behaviour. In practice it's best to learn the magic button to run your tests within your editor/IDE.

在命令行中，我可以运行 `go test` 并立即获得关于我的重构工作是否改变了行为的反馈。在实践中，最好学习魔术按钮以在您的编辑器/IDE 中运行您的测试。

You want to get in to a state where you are doing

你想进入一个你正在做的状态

- Small refactor
- Run tests
- Repeat

- 小重构
- 运行测试
- 重复

All within a very tight feedback loop so you don't go down rabbit holes and make mistakes.

所有这些都在一个非常紧密的反馈循环中，因此您不会陷入困境并犯错误。

Having a project where all your key behaviours are unit tested and give you feedback well under a second is a very empowering safety net to do bold refactoring when you need to. This helps us manage the incoming force of complexity that Lehman describes.

拥有一个项目，其中您的所有关键行为都经过单元测试并在一秒钟内为您提供反馈，这是一个非常强大的安全网，可以在您需要时进行大胆的重构。这有助于我们管理雷曼所描述的复杂性的传入力量。

## If unit tests are so great, why is there sometimes resistance to writing them?

## 如果单元测试如此出色，为什么有时会抵制编写它们？

On the one hand you have people (like me) saying that unit tests are important for the long term health of your system because they ensure you can keep refactoring with confidence.

一方面，有人（像我一样）说单元测试对于系统的长期健康很重要，因为它们确保您可以自信地继续重构。

On the other you have people describing experiences of unit tests actually _hindering_ refactoring.

另一方面，你有人描述了单元测试的经验，实际上是_阻碍_重构。

Ask yourself, how often do you have to change your tests when refactoring? Over the years I have been on many projects with very good test coverage and yet the engineers are reluctant to refactor because of the perceived effort of changing tests.

问问自己，在重构时，您需要多久更改一次测试？多年来，我参与了许多测试覆盖率非常好的项目，但工程师们不愿意重构，因为他们认为更改测试需要付出很多努力。

This is the opposite of what we are promised!

这与我们承诺的相反！

### Why is this happening?

### 为什么会这样？

Imagine you were asked to develop a square and we thought the best way to accomplish that would be stick two triangles together.

想象一下，你被要求开发一个正方形，我们认为最好的方法是将两个三角形粘在一起。

![Two right-angled triangles to form a square](https://i.imgur.com/ela7SVf.jpg)

We write our unit tests around our square to make sure the sides are equal and then we write some tests around our triangles. We want to make sure our triangles render correctly so we assert that the angles sum up to 180 degrees, perhaps check we make 2 of them, etc etc. Test coverage is really important and writing these tests is pretty easy so why not?

我们围绕正方形编写单元测试以确保边相等，然后围绕三角形编写一些测试。我们想确保我们的三角形渲染正确，所以我们断言角度总和为 180 度，也许检查我们制作了其中的 2 个，等等。测试覆盖率非常重要，编写这些测试非常容易，为什么不呢？

A few weeks later The Law of Continuous Change strikes our system and a new developer makes some changes. She now believes it would be better if squares were formed with 2 rectangles instead of 2 triangles.

几周后，持续变化法则冲击了我们的系统，一位新的开发人员做出了一些改变。她现在认为如果正方形由 2 个矩形而不是 2 个三角形组成会更好。

![Two rectangles to form a square](https://i.imgur.com/1G6rYqD.jpg)

She tries to do this refactor and gets mixed signals from a number of failing tests. Has she actually broken important behaviours here? She now has to dig through these triangle tests and try and understand what's going on.

她尝试进行这种重构，并从许多失败的测试中得到混合信号。她真的违反了这里的重要行为吗？她现在必须深入研究这些三角形测试并尝试了解发生了什么。

_It's not actually important that the square was formed out of triangles_ but **our tests have falsely elevated the importance of our implementation details**.

_正方形由三角形构成实际上并不重要_但是**我们的测试错误地提高了我们的实现细节的重要性**。

## Favour testing behaviour rather than implementation detail

## 喜欢测试行为而不是实现细节

When I hear people complaining about unit tests it is often because the tests are at the wrong abstraction level. They're testing implementation details, overly spying on collaborators and mocking too much.

当我听到人们抱怨单元测试时，通常是因为测试处于错误的抽象级别。他们正在测试实施细节，过度监视合作者并嘲笑太多。

I believe it stems from a misunderstanding of what unit tests are and chasing vanity metrics (test coverage).

我相信它源于对单元测试是什么和追求虚荣指标（测试覆盖率）的误解。

If I am saying just test behaviour, should we not just only write system/black-box tests? These kind of tests do have lots of value in terms of verifying key user journeys but they are typically expensive to write and slow to run. For that reason they're not too helpful for _refactoring_ because the feedback loop is slow. In addition black box tests don't tend to help you very much with root causes compared to unit tests.

如果我只是说测试行为，我们不应该只编写系统/黑盒测试吗？这些类型的测试在验证关键用户旅程方面确实具有很多价值，但它们通常编写成本高且运行缓慢。出于这个原因，它们对 _refactoring_ 没有太大帮助，因为反馈循环很慢。此外，与单元测试相比，黑盒测试对根本原因的帮助不大。

So what _is_ the right abstraction level?

那么什么是正确的抽象级别？

## Writing effective unit tests is a design problem

## 编写有效的单元测试是一个设计问题

Forgetting about tests for a moment, it is desirable to have within your system self-contained, decoupled "units" centered around key concepts in your domain.

暂时忘记测试，希望在您的系统中拥有以您领域中的关键概念为中心的独立的、解耦的“单元”。

I like to imagine these units as simple Lego bricks which have coherent APIs that I can combine with other bricks to make bigger systems. Underneath these APIs there could be dozens of things (types, functions et al) collaborating to make them work how they need to.

我喜欢把这些单元想象成简单的乐高积木，它们具有连贯的 API，我可以将这些 API 与其他积木结合起来制作更大的系统。在这些 API 之下，可能有几十种东西（类型、函数等）协同工作，使它们按照自己的需要工作。

For instance if you were writing a bank in Go, you might have an "account" package. It will present an API that does not leak implementation detail and is easy to integrate with. 

例如，如果您正在用 Go 编写银行，您可能有一个“帐户”包。它将提供一个不会泄露实现细节且易于集成的 API。

If you have these units that follow these properties you can write unit tests against their public APIs. _By definition_ these tests can only be testing useful behaviour. Underneath these units I am free to refactor the implementation as much as I need to and the tests for the most part should not get in the way.

如果您拥有遵循这些属性的这些单元，您可以针对它们的公共 API 编写单元测试。 _根据定义_这些测试只能测试有用的行为。在这些单元下面，我可以根据需要自由地重构实现，并且大部分测试不应该妨碍。

### Are these unit tests?

### 这些是单元测试吗？

**YES**. Unit tests are against "units" like I described. They were _never_ about only being against a single class/function/whatever.

**是的**。单元测试针对的是我所描述的“单元”。他们_从不_只反对一个类/函数/任何东西。

## Bringing these concepts together

## 将这些概念结合在一起

We've covered

我们已经涵盖

- Refactoring
- Unit tests
- Unit design

- 重构
- 单元测试
- 单元设计

What we can start to see is that these facets of software design reinforce each other.

我们可以开始看到，软件设计的这些方面相互加强。

### Refactoring

### 重构

- Gives us signals about our unit tests. If we have to do manual checks, we need more tests. If tests are wrongly failing then our tests are at the wrong abstraction level (or have no value and should be deleted).
- Helps us handle the complexities within and between our units.

- 为我们提供有关单元测试的信号。如果我们必须进行手动检查，我们需要更多的测试。如果测试错误地失败，那么我们的测试处于错误的抽象级别（或者没有价值，应该删除）。
- 帮助我们处理我们单位内部和单位之间的复杂性。

### Unit tests

### 单元测试

- Give a safety net to refactor.
- Verify and document the behaviour of our units.

- 提供重构的安全网。
- 验证并记录我们单位的行为。

### (Well designed) units

### （精心设计）单位

- Easy to write _meaningful_ unit tests.
- Easy to refactor.

- 易于编写 _meaningful_ 单元测试。
- 易于重构。

Is there a process to help us arrive at a point where we can constantly refactor our code to manage complexity and keep our systems malleable?

是否有一个过程可以帮助我们达到可以不断重构代码以管理复杂性并保持系统可塑性的程度？

## Why Test Driven Development (TDD)

## 为什么是测试驱动开发 (TDD)

Some people might take Lehman's quotes about how software has to change and overthink elaborate designs, wasting lots of time upfront trying to create the "perfect" extensible system and end up getting it wrong and going nowhere.

有些人可能会引用雷曼兄弟关于软件必须如何改变和过度考虑精心设计的引言，在前期浪费大量时间试图创建“完美”的可扩展系统，但最终会出错而无处可去。

This is the bad old days of software where an analyst team would spend 6 months writing a requirements document and an architect team would spend another 6 months coming up with a design and a few years later the whole project fails.

这是过去糟糕的软件时代，分析师团队会花 6 个月时间编写需求文档，架构师团队会再花 6 个月时间进行设计，几年后整个项目都失败了。

I say bad old days but this still happens!

我说糟糕的过去，但这种情况仍然发生！

Agile teaches us that we need to work iteratively, starting small and evolving the software so that we get fast feedback on the design of our software and how it works with real users; TDD enforces this approach.

敏捷告诉我们，我们需要迭代地工作，从小处着手并不断发展软件，以便我们能够快速获得关于软件设计以及它如何与真实用户一起工作的反馈； TDD 强制执行这种方法。

TDD addresses the laws that Lehman talks about and other lessons hard learned through history by encouraging a methodology of constantly refactoring and delivering iteratively.

通过鼓励不断重构和迭代交付的方法，TDD 解决了雷曼兄弟谈论的法律和其他从历史中艰难吸取的教训。

### Small steps

### 小步骤

- Write a small test for a small amount of desired behaviour
- Check the test fails with a clear error (red)
- Write the minimal amount of code to make the test pass (green)
- Refactor
- Repeat

- 为少量所需行为编写一个小测试
- 检查测试失败并显示明确错误（红色）
- 编写最少的代码以使测试通过（绿色）
- 重构
- 重复

As you become proficient, this way of working will become natural and fast.

随着您熟练掌握，这种工作方式将变得自然而快速。

You'll come to expect this feedback loop to not take very long and feel uneasy if you're in a state where the system isn't "green" because it indicates you may be down a rabbit hole.

如果您处于系统不是“绿色”的状态，您会期望此反馈循环不会花费很长时间并感到不安，因为这表明您可能陷入困境。

You'll always be driving small & useful functionality comfortably backed by the feedback from your tests.

在您的测试反馈的支持下，您将始终可以轻松地驱动小而有用的功能。

## Wrapping up

##  总结

- The strength of software is that we can change it. _Most_ software will require change over time in unpredictable ways; but don't try and over-engineer because it's too hard to predict the future.
- Instead we need to make it so we can keep our software malleable. In order to change software we have to refactor it as it evolves or it will turn into a mess
- A good test suite can help you refactor quicker and in a less stressful manner
- Writing good unit tests is a design problem so think about structuring your code so you have meaningful units that you can integrate together like Lego bricks.
- TDD can help and force you to design well factored software iteratively, backed by tests to help future work as it arrives. 

- 软件的优势在于我们可以改变它。 _大多数_软件将需要以不可预测的方式随时间变化；但不要尝试过度设计，因为预测未来太难了。
- 相反，我们需要制作它，以便我们可以保持我们的软件可塑性。为了改变软件，我们必须随着它的发展重构它，否则它会变成一团糟
- 一个好的测试套件可以帮助您以更少压力的方式更快地重构
- 编写好的单元测试是一个设计问题，因此请考虑构建代码，以便拥有可以像乐高积木一样集成在一起的有意义的单元。
- TDD 可以帮助并迫使您迭代地设计分解良好的软件，并以测试为后盾，以帮助未来的工作。

