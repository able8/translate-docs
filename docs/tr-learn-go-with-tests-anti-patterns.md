# TDD Anti-patterns

# TDD 反模式

From time to time it's necessary to review your TDD techniques and remind yourself of behaviours to avoid.

有时有必要回顾一下您的 TDD 技术并提醒自己要避免的行为。

The TDD process is conceptually simple to follow, but as you do it you'll find it challenging your design skills. **Don't mistake this for TDD being hard, it's design that's hard!**

TDD 过程在概念上很容易遵循，但是当您这样做时，您会发现它对您的设计技能具有挑战性。 **不要误认为 TDD 很难，这是设计很难！**

This chapter lists a number of TDD and testing anti-patterns, and how to remedy them.

本章列出了一些 TDD 和测试反模式，以及如何补救它们。

## Not doing TDD at all

## 根本不做 TDD

Of course, it is possible to write great software without TDD but, a lot of problems I've seen with the design of code and the quality of tests would be very difficult to arrive at if a disciplined approach to TDD had been used.

当然，不用 TDD 也可以编写出色的软件，但是，如果使用严格的 TDD 方法，我在代码设计和测试质量方面看到的很多问题将很难解决。

One of the strengths of TDD is that it gives you a formal process to break down problems, understand what you're trying to achieve (red), get it done (green), then have a good think about how to make it right ( blue/refactor).

TDD 的优势之一是它为您提供了一个正式的过程来分解问题，了解您要实现的目标（红色），完成它（绿色），然后好好思考如何使其正确（蓝色/重构）。

Without this, the process is often ad-hoc and loose, which _can_ make engineering more difficult than it _could_ be.

没有这个，这个过程通常是临时的和松散的，这_可以_使工程比它_可能_变得更加困难。

## Misunderstanding the constraints of the refactoring step

## 误解重构步骤的约束

I have been in a number of workshops, mobbing or pairing sessions where someone has made a test pass and is in the refactoring stage. After some thought, they think it would be good to abstract away some code into a new struct; a budding pedant yells:

我参加过许多研讨会、围攻或配对会议，其中有人通过了测试并处于重构阶段。经过一番思考，他们认为将一些代码抽象成一个新的结构体会很好；一个初出茅庐的学究喊道：

> You're not allowed to do this! You should write a test for this first, we're doing TDD!

> 你不能这样做！你应该先为此编写一个测试，我们正在做 TDD！

This seems to be a common misunderstanding. **You can do whatever you like to the code when the tests are green**, the only thing you're not allowed to do is **add or change behaviour**.

这似乎是一个普遍的误解。 **当测试是绿色的时候，你可以对代码做任何你喜欢的事情**，你唯一不能做的就是**添加或改变行为**。

The point of these tests are to give you the _freedom to refactor_, find the right abstractions and make the code easier to change and understand.

这些测试的目的是给你_自由重构_，找到正确的抽象并使代码更容易改变和理解。

## Having tests that won't fail (or, evergreen tests)

## 有不会失败的测试（或常绿测试）

It's astonishing how often this comes up. You start debugging or changing some tests and realise: there are no scenarios where this test can fail. Or at least, it won't fail in the way the test is _supposed_ to be protecting against.

令人惊讶的是，这种情况出现的频率很高。您开始调试或更改一些测试并意识到：没有任何情况会导致此测试失败。或者至少，它不会以_应该_保护的测试方式失败。

This is _next to impossible_ with TDD if you're following **the first step**,

如果您遵循 **第一步**，这对于 TDD 来说是_几乎不可能的_，

> Write a test, see it fail

> 写一个测试，看它失败

This is almost always done when developers write tests _after_ code is written, and/or chasing test coverage rather than creating a useful test suite.

当开发人员在编写代码之后编写测试和/或追逐测试覆盖率而不是创建有用的测试套件时，几乎总是这样做。

## Useless assertions

## 无用的断言

Ever worked on a system, and you've broken a test, then you see this?

曾经在一个系统上工作过，并且你打破了一个测试，然后你看到了吗？

> `false was not equal to true`

> `假不等于真`

I know that false is not equal to true. This is not a helpful message; it doesn't tell me what I've broken. This is a symptom of not following the TDD process and not reading the failure error message.

我知道假不等于真。这不是一个有用的信息；它没有告诉我我破坏了什么。这是未遵循 TDD 流程且未读取失败错误消息的症状。

Going back to the drawing board,

回到绘图板，

> Write a test, see it fail (and don't be ashamed of the error message)

> 写一个测试，看到它失败（不要为错误信息感到羞耻）

## Asserting on irrelevant detail

## 断言不相关的细节

An example of this is making an assertion on a complex object, when in practice all you care about in the test is the value of one of the fields.

这方面的一个例子是对复杂对象进行断言，而在实践中，您在测试中只关心其中一个字段的值。

```go
// not this, now your test is tightly coupled to the whole object
if !cmp.Equal(complexObject, want) {
    t.Error("got %+v, want %+v", complexObject, want)
}

// be specific, and loosen the coupling
got := complexObject.fieldYouCareAboutForThisTest
if got != want{
    t.Error("got %q, want %q", got, want)
}
```

Additional assertions not only make your test more difficult to read by creating 'noise' in your documentation, but also needlessly couples the test with data it doesn't care about. This means if you happen to change the fields for your object, or the way they behave you may get unexpected compilation problems or failures with your tests.

额外的断言不仅通过在文档中创建“噪音”而使您的测试更难以阅读，而且还不必要地将测试与它不关心的数据耦合在一起。这意味着如果您碰巧更改了对象的字段或它们的行为方式，您可能会遇到意外的编译问题或测试失败。

This is an example of not following the red stage strictly enough.

这是一个没有严格遵循红色阶段的例子。

- Letting an existing design influence how you write your test **rather than thinking of the desired behaviour**
- Not giving enough consideration to the failing test's error message

- 让现有设计影响您编写测试的方式**而不是考虑所需的行为**
- 没有充分考虑失败测试的错误信息

## Lots of assertions within a single scenario for unit tests

## 单个场景中的大量断言用于单元测试

Many assertions can make tests difficult to read and challenging to debug when they fail.

许多断言会使测试难以阅读，并且在失败时难以调试。

They often creep in gradually, especially if test setup is complicated because you're reluctant to replicate the same horrible setup to assert on something else. Instead of this you should fix the problems in your design which are making it difficult to assert on new things. 

它们通常会逐渐出现，尤其是在测试设置很复杂的情况下，因为您不愿意复制相同的可怕设置来断言其他内容。取而代之的是，您应该解决设计中的问题，这些问题使您难以对新事物断言。

A helpful rule of thumb is to aim to make one assertion per test. In Go, take advantage of subtests to clearly delineate between assertions on the occasions where you need to. This is also a handy technique to separate assertions on behaviour vs implementation detail.

一个有用的经验法则是旨在每个测试做出一个断言。在 Go 中，利用子测试在需要的场合清楚地划分断言。这也是一种将行为断言与实现细节分开的便捷技术。

For other tests where setup or execution time may be a constraint (e.g an acceptance test driving a web browser), you need to weigh up the pros and cons of slightly trickier to debug tests against test execution time.

对于设置或执行时间可能是限制的其他测试（例如驱动 Web 浏览器的验收测试），您需要权衡稍微棘手的调试测试与测试执行时间的利弊。

## Not listening to your tests

## 不听你的测试

[Dave Farley in his video "When TDD goes wrong"](https://www.youtube.com/watch?v=UWtEVKVPBQ0&feature=youtu.be) points out,

[Dave Farley 在他的视频“当 TDD 出错时”](https://www.youtube.com/watch?v=UWtEVKVPBQ0&feature=youtu.be) 指出，

> TDD gives you the fastest feedback possible on your design

> TDD 为您的设计提供最快的反馈

From my own experience, a lot of developers are trying to practice TDD but frequently ignore the signals coming back to them from the TDD process. So they're still stuck with fragile, annoying systems, with a poor test suite.

根据我自己的经验，很多开发人员都在尝试实践 TDD，但经常忽略从 TDD 过程返回给他们的信号。所以他们仍然坚持使用脆弱的、烦人的系统，以及糟糕的测试套件。

Simply put, if testing your code is difficult, then _using_ your code is difficult too. Treat your tests as the first user of your code and then you'll see if your code is pleasant to work with or not.

简单地说，如果测试你的代码很困难，那么_使用_你的代码也很困难。将您的测试视为代码的第一个用户，然后您就会看到您的代码是否适合使用。

I've emphasised this a lot in the book, and I'll say it again **listen to your tests**.

我在书中强调了很多，我再说一遍**听你的测试**。

### Excessive setup, too many test doubles, etc.

### 过多的设置，太多的测试替身等等。

Ever looked at a test with 20, 50, 100, 200 lines of setup code before anything interesting in the test happens? Do you then have to change the code and revisit the mess and wish you had a different career?

在测试中发生任何有趣的事情之前，是否曾经看过包含 20、50、100、200 行设置代码的测试？那么您是否必须更改代码并重新审视混乱并希望您拥有不同的职业？

What are the signals here? _Listen_, complicated tests `==` complicated code. Why is your code complicated? Does it have to be?

这里有什么信号？ _听_，复杂的测试`==`复杂的代码。为什么你的代码很复杂？必须是吗？

- When you have lots of test doubles in your tests, that means the code you're testing has lots of dependencies - which means your design needs work.
- If your test is reliant on setting up various interactions with mocks, that means your code is making lots of interactions with its dependencies. Ask yourself whether these interactions could be simpler.

- 当您的测试中有很多测试替身时，这意味着您正在测试的代码有很多依赖项 - 这意味着您的设计需要工作。
- 如果您的测试依赖于设置与模拟的各种交互，则意味着您的代码与其依赖项进行了大量交互。问问自己这些交互是否可以更简单。

#### Leaky interfaces

#### 泄漏接口

If you have declared an `interface` that has many methods, that points to a leaky abstraction. Think about how you could define that collaboration with a more consolidated set of methods, ideally one.

如果你声明了一个有很多方法的`interface`，那么它就指向了一个有漏洞的抽象。想一想如何使用一组更统一的方法来定义这种协作，最好是一个。

#### Think about the types of test doubles you use

#### 想想你使用的测试替身的类型

- Mocks are sometimes helpful, but they're extremely powerful and therefore easy to misuse. Try giving yourself the constraint of using stubs instead.
- Verifying implementation detail with spies is sometimes helpful, but try to avoid it. Remember your implementation detail is usually not important, and you don't want your tests coupled to them if possible. Look to couple your tests to **useful behaviour rather than incidental details**.
- [Read my posts on naming test doubles](https://quii.dev/Start_naming_your_test_doubles_correctly) if the taxonomy of test doubles is a little unclear

- 模拟有时很有用，但它们非常强大，因此很容易被滥用。尝试给自己限制使用存根代替。
- 使用 spies 验证实现细节有时会有所帮助，但请尽量避免它。请记住，您的实现细节通常并不重要，如果可能，您不希望您的测试与它们耦合。将您的测试与**有用的行为而不是附带的细节**结合起来。
- [阅读我关于命名测试替身的帖子](https://quii.dev/Start_naming_your_test_doubles_correctly) 如果测试替身的分类有点不清楚

#### Consolidate dependencies

#### 整合依赖

Here is some code for a `http.HandlerFunc` to handle new user registrations for a website.

下面是一些用于处理网站新用户注册的 `http.HandlerFunc` 的代码。

```go
type User struct {
    // Some user fields
}

type UserStore interface {
    CheckEmailExists(email string) (bool, error)
    StoreUser(newUser User) error
}

type Emailer interface {
    SendEmail(to User, body string, subject string) error
}

func NewRegistrationHandler(userStore UserStore, emailer Emailer) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        // extract out the user from the request body (handle error)
        // check user exists (handle duplicates, errors)
        // store user (handle errors)
        // compose and send confirmation email (handle error)
        // if we got this far, return 2xx response
    }
}
```

At first pass it's reasonable to say the design isn't so bad. It only has 2 dependencies!

乍一看，可以说设计还不错。它只有 2 个依赖项！

Re-evaluate the design by considering the handler's responsibilities:

通过考虑处理者的职责重新评估设计：

- Parse the request body into a `User` ✅
- Use `UserStore` to check if the user exists ❓
- Use `UserStore` to store the user ❓
- Compose an email ❓
- Use `Emailer` to send the email ❓
- Return an appropriate http response, depending on success, errors, etc ✅

- 将请求体解析为 `User` ✅
- 使用`UserStore`检查用户是否存在❓
- 使用 `UserStore` 来存储用户 ❓
- 撰写电子邮件 ❓
- 使用`Emailer`发送邮件❓
- 返回适当的 http 响应，取决于成功、错误等 ✅

To exercise this code, you're going to have to write many tests with varying degrees of test double setups, spies, etc

要练习此代码，您将不得不编写许多具有不同程度的测试双重设置、间谍等的测试

- What if the requirements expand? Translations for the emails? Sending an SMS confirmation too? Does it make sense to you that you have to change a HTTP handler to accommodate this change?
- Does it feel right that the important rule of "we should send an email" resides within a HTTP handler? 

- 如果需求扩大怎么办？电子邮件的翻译？也发送短信确认？必须更改 HTTP 处理程序以适应此更改对您来说有意义吗？
- “我们应该发送电子邮件”的重要规则存在于 HTTP 处理程序中，这感觉对吗？

- Why do you have to go through the ceremony of creating HTTP requests and reading responses to verify that rule?

- 为什么你必须经历创建 HTTP 请求和读取响应的仪式来验证该规则？

**Listen to your tests**. Writing tests for this code in a TDD fashion should quickly make you feel uncomfortable (or at least, make the lazy developer in you be annoyed). If it feels painful, stop and think.

**听你的测试**。以 TDD 方式为这段代码编写测试应该很快会让你感到不舒服（或者至少会让你这个懒惰的开发人员感到恼火）。如果感到痛苦，请停下来想一想。

What if the design was like this instead?

如果设计是这样的呢？

```go
type UserService interface {
    Register(newUser User) error
}

func NewRegistrationHandler(userService UserService) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        // parse user
        // register user
        // check error, send response
    }
}
```

- Simple to test the handler ✅
- Changes to the rules around registration are isolated away from HTTP, so they are also simpler to test ✅

- 简单的测试处理程序✅
- 注册规则的更改与 HTTP 隔离，因此它们也更易于测试 ✅

## Violating encapsulation

## 违反封装

Encapsulation is very important. There's a reason we don't make everything in a package exported (or public). We want coherent APIs with a small surface area to avoid tight coupling.

封装非常重要。我们不将包中的所有内容都导出（或公开）是有原因的。我们希望具有小表面积的连贯 API 以避免紧密耦合。

People will sometimes be tempted to make a function or method public in order to test something. By doing this you make your design worse and send confusing messages to maintainers and users of your code.

人们有时会试图将函数或方法公开以测试某些内容。这样做会使您的设计变得更糟，并向代码的维护者和用户发送令人困惑的信息。

A result of this can be developers trying to debug a test and then eventually realising the function being tested is _only called from tests_. Which is obviously **a terrible outcome, and a waste of time**.

这样做的结果可能是开发人员试图调试测试，然后最终意识到被测试的函数_仅从测试中调用_。这显然是**一个可怕的结果，而且是浪费时间**。

In Go, consider your default position for writing tests as _from the perspective of a consumer of your package_. You can make this a compile-time constraint by having your tests live in a test package e.g `package gocoin_test`. If you do this, you'll only have access to the exported members of the package so it won't be possible to couple yourself to implementation detail.

在 Go 中，将您编写测试的默认位置视为_从包消费者的角度来看_。您可以通过将您的测试放在测试包中来使其成为编译时约束，例如`package gocoin_test`。如果这样做，您将只能访问包的导出成员，因此不可能将自己与实现细节联系起来。

## Complicated table tests

## 复杂的表测试

Table tests are a great way of exercising a number of different scenarios when the test setup is the same, and you only wish to vary the inputs.

当测试设置相同并且您只希望改变输入时，表测试是一种很好的练习多种不同场景的方法。

_But_ they can be messy to read and understand when you try to shoehorn other kinds of tests under the name of having one, glorious table.

_但是_当您试图以拥有一张光彩夺目的表的名义强加其他类型的测试时，它们可能很难阅读和理解。

```go
cases := []struct {
    X int
    Y int
    Z int
    err error
    IsFullMoon bool
    IsLeapYear bool
    AtWarWithEurasia bool
}
```

**Don't be afraid to break out of your table and write new tests** rather than adding new fields and booleans to the table `struct`.

**不要害怕打破你的表并编写新的测试**而不是向表`struct`添加新的字段和布尔值。

A thing to bear in mind when writing software is,

编写软件时要记住的一件事是，

> [Simple is not easy](https://www.infoq.com/presentations/Simple-Made-Easy/)

> [简单不易](https://www.infoq.com/presentations/Simple-Made-Easy/)

"Just" adding a field to a table might be easy, but it can make things far from simple.

“只是”向表中添加一个字段可能很容易，但它会使事情变得非常简单。

## Summary

##  概括

Most problems with unit tests can normally be traced to:

单元测试的大多数问题通常可以追溯到：

- Developers not following the TDD process
- Poor design

- 不遵循 TDD 流程的开发人员
- 糟糕的设计

So, learn about good software design!

所以，学习优秀的软件设计吧！

The good news is TDD can help you _improve your design skills_ because as stated in the beginning:

好消息是 TDD 可以帮助您_提高您的设计技能_，因为正如开头所述：

**TDD's main purpose is to provide feedback on your design.** For the millionth time, listen to your tests, they are reflecting your design back at you.

**TDD 的主要目的是为您的设计提供反馈。** 第 100 万次，聆听您的测试，它们将您的设计反馈给您。

Be honest about the quality of your tests by listening to the feedback they give you, and you'll become a better developer for it. 

通过听取他们给你的反馈，诚实地对待你的测试质量，你会成为一个更好的开发人员。

