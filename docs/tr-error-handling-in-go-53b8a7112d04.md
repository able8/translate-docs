# Error Handling in Go

# Go 中的错误处理

## Mastering pragmatic error handling in your Go code
[May 1, 2019·10 min read](https://medium.com/gett-engineering/error-handling-in-go-53b8a7112d04?source=post_page-----53b8a7112d04--------------------------------)
*This post is part of the “*[*Before you go Go*](https://medium.com/gett-engineering/before-you-go-go-bf4f861cdec7)*” series, where we explore the world of Golang, provide tips and insights you should know when writing Go, so you don't have to learn them the  hard way.*

## 掌握 Go 代码中的实用错误处理
*这篇文章是“*[*Before you go*](https://medium.com/gett-engineering/before-you-go-go-bf4f861cdec7)*”系列的一部分，我们在其中探索Golang，提供您在编写 Go 时应该知道的技巧和见解，这样您就不必费力地学习它们。*

I assume you have at least some basic Go background, but if you feel at  any point you’re not familiar with the materials discussed, feel free to pause, research and come back.

我假设你至少有一些基本的 Go 背景，但如果你觉得在任何时候你对所讨论的材料不熟悉，请随时暂停，研究然后回来。

Now that we got all of this out of the way, let’s Go!

现在我们已经解决了所有这些问题，让我们走吧！

Go’s approach to error handling is one of its most controversial and misused features. In this article, you’ll learn how Go approaches errors and understand how they  work under the hood. You’ll explore a couple of different approaches to  it, take a look at the go source code and the standard library for some  insights about how errors work and how to work with them. You’ll learn  how Type Assertions play an important role in handling them, and see  upcoming changes to error handling, planned to be released in Go 2.
## Introduction

Go 的错误处理方法是其最具争议和滥用的功能之一。在本文中，您将了解 Go 如何处理错误并了解它们在底层是如何工作的。您将探索几种不同的方法，查看 go 源代码和标准库，以了解有关错误如何工作以及如何使用它们的一些见解。您将了解类型断言如何在处理它们时发挥重要作用，并看到即将在 Go 2 中发布的错误处理更改。
##  介绍

First thing’s first: Errors in Go are **not** Exceptions.

第一件事是：Go 中的错误不是 异常。

 wrote an [epic blog post](https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right) about it, so I'll refer you to it and summarise: In other languages,  you are uncertain if a function may throw an exception or not. Instead  of throwing exceptions, Go functions support **multiple return values**, and by convention, this ability is commonly used to return the function’s result along with an error variable.

写了一篇[史诗般的博客文章](https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right)关于它，所以我会向你推荐它并总结：在其他语言，你不确定一个函数是否会抛出异常。 Go 函数不抛出异常，而是支持 **multiple return values**，并且按照惯例，这种能力通常用于返回函数的结果以及错误变量。

![img](https://miro.medium.com/max/1400/1*_tbzb7GuFE4HGWhHX-N96g.png)

If your function can fail for some reason, you should probably return the predeclared `error` type from it. By convention, returning an error signals the caller  there was a problem, and returning nil represents no error. This way,  you’re letting the caller understand there was a problem, and let them  deal with it: whoever calls your function knows they should not rely on  the result before checking the error. If error is not nil, it is their  responsibility to check for it and handle it (log, return, serve, invoke some retry/cleanup mechanism, etc.).

如果您的函数由于某种原因可能失败，您可能应该从中返回预先声明的 `error` 类型。按照惯例，返回错误表示调用者有问题，返回 nil 表示没有错误。这样，您就可以让调用者了解存在问题，并让他们处理它：调用您的函数的人都知道他们不应该在检查错误之前依赖结果。如果错误不是零，他们有责任检查并处理它（记录、返回、服务、调用一些重试/清理机制等）。

![img](https://miro.medium.com/max/1400/1*idmULuK_hDdBuKOR9WmaYA.png)

These snippets are very common in Go, and some see them as a whole lot of  boiler plate code. The compiler treats unused variables as compilation  errors, so if you’re not going to check the error, you should assign  them to the [blank identifier](https://golang.org/ref/spec#Blank_identifier). But as convenient as it is, errors should not be ignored.

这些片段在 Go 中非常常见，有些人将它们视为一大堆样板代码。编译器将未使用的变量视为编译错误，因此如果您不打算检查错误，则应将它们分配给 [空白标识符](https://golang.org/ref/spec#Blank_identifier)。但尽管方便，但不应忽略错误。

![img](https://miro.medium.com/max/1400/1*sePtCmqUV3qytc-CLSyU-g.png)

result can’t be trusted before checking for errors

在检查错误之前不能信任结果

Returning the error along with the results, along with Go’s strict type system,  makes programming errors that much harder to write. You should always  assume the function’s value is corrupted unless you’ve checked the error it returned, and by assigning the error to the blank identifier, you  explicitly ignore that your function’s value may be corrupt.

将错误和结果一起返回，再加上 Go 的严格类型系统，使得编程错误更难编写。您应该始终假设函数的值已损坏，除非您已检查它返回的错误，并且通过将错误分配给空白标识符，您可以明确忽略您的函数值可能已损坏。


The blank identifier is dark and full of terrors. 

空白的标识符是黑暗的，充满了恐怖。

Go does have a `panic` and `recover` mechanism, which is also described in [another detailed Go blog post](https://blog.golang.org/defer-panic-and-recover). But these are not meant to mimic exceptions. In the words of Dave, *“When you panic in Go, you’re freaking out, it’s not someone else’s problem, it’s game over man”.* They’re fatal, and lead to a crash in your program. Rob Pike coined the *“Don’t Panic”* proverb, which is self-explanatory: you should probably avoid it, and return errors instead.

Go 确实有一个 `panic` 和 `recover` 机制，这也在 [另一篇详细的 Go 博客文章](https://blog.golang.org/defer-panic-and-recover) 中进行了描述。但这些并不意味着模仿异常。用 Dave 的话来说，*“当你在 Go 中恐慌时，你会吓坏了，这不是别人的问题，而是人的问题”。* 它们是致命的，会导致你的程序崩溃。 Rob Pike 创造了 *“Don’t Panic”* 谚语，这是不言自明的：你应该避免它，而是返回错误。

> *“Errors are values.”
> “Don’t just check errors, handle them gracefully”
> “Don’t panic”*
> [all of Rob Pike’s Go Proverbs](https://go-proverbs.github.io/)

> *“错误就是价值。”
> “不要只检查错误，要优雅地处理它们”
> “不要惊慌”*
> [Rob Pike 的所有 Go Proverbs](https://go-proverbs.github.io/)

# Under the hood

# 引擎盖下

## The error Interface

## 错误接口

Under the hood, the error type is a [simple single method interface](https://golang.org/ref/spec#Errors), and if you're not familiar with it, I highly recommend going over [this post](https://blog.golang.org/error-handling-and-go) in the official Go Blog.

在幕后，错误类型是 [简单的单一方法接口](https://golang.org/ref/spec#Errors)，如果您不熟悉它，我强烈建议您阅读 [这篇文章](https://blog.golang.org/error-handling-and-go) 在 Go 官方博客中。

![img](https://miro.medium.com/max/1400/1*54Ys-4R9y3jV9ZHRNvyQUQ.png)

error interface from the source code

来自源代码的错误接口

It’s easy to implement your own errors, and there are various approaches to custom structs implementing the `Error() string` method. Any struct implementing this one method is considered a valid error value, and can be returned as such.

实现自己的错误很容易，并且有多种方法可以实现自定义结构体实现 `Error() string` 方法。任何实现这一方法的结构都被认为是有效的错误值，并且可以这样返回。

Let’s explore a few such approaches.

让我们探索一些这样的方法。

## The built-in errorString struct

## 内置的errorString结构

The most commonly used and widespread implementation of the error interface is the built-in `errorString` struct. It is the leanest implementation you can think of.

错误接口最常用和最广泛的实现是内置的 `errorString` 结构。这是您能想到的最精简的实现。

![img](https://miro.medium.com/max/1400/1*Nj-7j0ISZ5-U4YoC3wQIdA.png)

source: [Go source code](https://golang.org/src/errors/errors.go)

来源：[Go 源代码](https://golang.org/src/errors/errors.go)

You can see its simplistic implementation [here](https://golang.org/src/errors/errors.go). All it does is hold a `string`, and that string is returned by the `Error` method. That error string can be formatted with data by us, say, with `fmt.Sprintf`. But other than that, it does not pack any other capabilities. If you use the built-in `errors.New` or `fmt.Errorf`, you are [already using it](https://play.golang.org/p/olRXqq3jNyR).

您可以在 [此处](https://golang.org/src/errors/errors.go) 中看到它的简单实现。它所做的只是保存一个“字符串”，该字符串由“错误”方法返回。该错误字符串可以由我们使用数据格式化，例如，使用 `fmt.Sprintf`。但除此之外，它不包含任何其他功能。如果你使用内置的 `errors.New` 或 `fmt.Errorf`，你就[已经在使用它](https://play.golang.org/p/olRXqq3jNyR)。

![img](https://miro.medium.com/max/1400/1*L5gtS-4T2tLSY2bdWQtssA.png)

[try it](https://play.golang.org/p/oWy5BNY1Hzq)

[试试](https://play.golang.org/p/oWy5BNY1Hzq)

## `github.com/pkg/errors`



Another simple example is the `pkg/errors`[ package](https://github.com/pkg/errors/blob/master/errors.go). Not to be confused with the built-in `errors` package you learned about earlier, this package provides additional  important capabilities such as error wrapping, unwrapping, formatting  and stack trace recording. You can install the package by running `go get github.com/pkg/errors` .

另一个简单的例子是`pkg/errors`[包](https://github.com/pkg/errors/blob/master/errors.go)。不要与您之前了解的内置 `errors` 包混淆，该包提供了额外的重要功能，例如错误包装、解包、格式化和堆栈跟踪记录。您可以通过运行 `go get github.com/pkg/errors` 来安装该软件包。

![img](https://miro.medium.com/max/1400/1*f0Eupl1bRnET2mobHH9KiA.png)

In those cases where you need to attach stack traces to your errors, or  attach necessary debugging information to the error, using this  package's `New` or `Errorf` functions provides errors that already record your stack trace, and you can attach simple metadata using its formatting capabilities. `Errorf` implements the `fmt.Formatter`[ interface](https://golang.org/pkg/fmt/#Formatter), meaning you can format it using the `fmt` package runes ( `%s` , `% v` , `%+v` etc).

在您需要将堆栈跟踪附加到错误或将必要的调试信息附加到错误的情况下，使用此包的 `New` 或 `Errorf` 函数提供已经记录堆栈跟踪的错误，并且您可以使用附加简单的元数据它的格式化功能。 `Errorf` 实现了 `fmt.Formatter`[ 接口](https://golang.org/pkg/fmt/#Formatter)，这意味着你可以使用 `fmt` 包 runes ( `%s` , `% v`、`%+v` 等)。

![img](https://miro.medium.com/max/1400/1*YfKohATpDygg4LVofI5ixQ.png)

This package also introduces the `errors.Wrap` and `errors.Wrapf` functions. These functions add context to an error with a message and  stack trace at the point they were called. This way, instead of simply  returning an error, you can wrap it with its context and important debug data.

这个包还引入了 `errors.Wrap` 和 `errors.Wrapf` 函数。这些函数将上下文添加到错误中，并在调用它们时显示消息和堆栈跟踪。这样，您可以使用上下文和重要的调试数据将其包装起来，而不是简单地返回错误。

![img](https://miro.medium.com/max/1400/1*VuqmY-finoFSfem_diOYAg.png)

Errors wrapping errors support the `Cause() error` method, that returns their inner error. Also, they can be used with the `errors.Cause(err error) error` function which retrieves the underlying inner-most error within an error.

错误包装错误支持 `Cause() error` 方法，该方法返回其内部错误。此外，它们可以与 `errors.Cause(err error) error` 函数一起使用，该函数检索错误中的底层最内层错误。

# Working with Errors

# 处理错误

## Type Assertion 

## 类型断言

[Type Assertion](https://golang.org/ref/spec#Type_assertions) serves a great role when working with errors. You’ll use them to assert information out of an interface value, and since error handling deals  with custom implementations of the `error` interface, performing assertions on errors is a very handy tool.

[Type Assertion](https://golang.org/ref/spec#Type_assertions) 在处理错误时发挥了重要作用。您将使用它们从接口值中断言信息，并且由于错误处理涉及 `error` 接口的自定义实现，因此对错误执行断言是一个非常方便的工具。

Its syntax is the same for all its purposes —` x.(T)` , provided `x` is of an interface type.` x.(T)`asserts that `x` is not `nil` and that the value stored in `x` is of type `T`. In the next couple of sections, you’re going to take a look at the two ways to use type assertions — with a concrete type `T` and with an interface type `T`.

它的语法对于它的所有目的都是相同的——` x.(T)` ，前提是 `x` 是一个接口类型。` x.(T)`断言 `x` 不是 `nil` 并且存储的值`x` 中的类型是 `T`。在接下来的几节中，您将了解使用类型断言的两种方法——具体类型“T”和接口类型“T”。

![img](https://miro.medium.com/max/1400/1*eKGbb9oyNUk6KeNFZw2OoQ.png)

playground: [short syntax panic](https://play.golang.org/p/bl-O3lJrixF), [safe long syntax](https://play.golang.org/p/CLLyXQWyrgF)

游乐场：[短语法恐慌](https://play.golang.org/p/bl-O3lJrixF)，[安全长语法](https://play.golang.org/p/CLLyXQWyrgF)

*Side note regarding syntax: type assertion can be used with both short  syntax (which panics when assertion fails), and elongated syntax (which  uses an OK-boolean to indicate success or failure). I always recommend  taking the long one over the short one, since I prefer checking the OK  variable instead of dealing with a panic.*

*关于语法的旁注：类型断言可以与短语法（当断言失败时恐慌）和细长语法（使用 OK-boolean 来指示成功或失败）一起使用。我总是建议用长的而不是短的，因为我更喜欢检查 OK 变量而不是处理恐慌。*

## Asserting with interface type T

## 使用接口类型 T 断言

Performing type assertion `x.(T)` with an interface type`T` asserts that `x` implements the `T` interface. This way you can guarantee an interface value implements an  interface and only if it does, you will be able use its methods.

使用接口类型`T` 执行类型断言`x.(T)` 断言`x` 实现了`T` 接口。这样你就可以保证一个接口值实现了一个接口，只有当它实现了，你才能使用它的方法。

![img](https://miro.medium.com/max/1400/1*GuDKwskV-jOBStbZ6bbS2Q.png)

To understand how this can be leveraged, let’s take a quick look in the `pkg/errors` again. You already know the errors package, so let’s dive into the `errors.Cause(err error) error` function.

要了解如何利用它，让我们再次快速浏览一下 `pkg/errors`。您已经知道错误包，所以让我们深入研究 `errors.Cause(err error) error` 函数。

This function gets an error and extracts the most internal error it wraps  (that which no longer wraps another error inside it). It might seem  simple, but there are plenty of great things you can learn from this  implementation:

这个函数得到一个错误并提取它包装的最内部错误（不再包装另一个错误的内部错误）。这可能看起来很简单，但你可以从这个实现中学到很多很棒的东西：

![img](https://miro.medium.com/max/1400/1*qIfFWhz71ITSNxBPpXgVqg.png)

source: [pkg/errors](https://github.com/pkg/errors/blob/master/errors.go#L269)

来源：[pkg/errors](https://github.com/pkg/errors/blob/master/errors.go#L269)

The function gets an error value and it can’t assume the `err` argument it receives is a wrapped error (one that supports the `Cause` method). So before calling the `Cause` method, it is necessary to check that you’re dealing with an error that implements this method. By performing the type assertion in each  iteration of the for loop, you can make sure the `cause` variable supports the `Cause` method, and can keep on extracting internal errors from it until you reach an error which does not have a cause.

该函数获得一个错误值，并且它不能假设它收到的 `err` 参数是一个包装错误（一个支持 `Cause` 方法的错误）。因此，在调用 `Cause` 方法之前，有必要检查您是否正在处理实现该方法的错误。通过在 for 循环的每次迭代中执行类型断言，您可以确保 `cause` 变量支持 `Cause` 方法，并且可以继续从中提取内部错误，直到遇到没有原因的错误。

By creating a lean, local interface containing just the methods you need  and performing the assertion on it, your code is decoupled from other  dependencies. The argument you received doesn’t have to be a known  struct, it just needs to be an error. Any type implementing the `Error` and `Cause` methods works here. So, if you implement the `Cause` method in your custom error type, you can use this function with it instantly.

通过创建仅包含您需要的方法并对其执行断言的精益本地接口，您的代码与其他依赖项分离。你收到的参数不一定是一个已知的结构，它只需要是一个错误。任何实现 `Error` 和 `Cause` 方法的类型都可以在这里工作。因此，如果您在自定义错误类型中实现 `Cause` 方法，则可以立即使用该函数。

There’s one small catch you should be aware of, though: interfaces may change,  and so you should maintain your code carefully, so your assertions don’t break. Remember to define your interfaces where you use them, keep them lean, and maintain them carefully and you should be fine.

不过，您应该注意一个小问题：接口可能会更改，因此您应该小心维护代码，以免断言中断。记住在你使用它们的地方定义你的接口，保持它们的精益，并小心地维护它们，你应该没问题。

Lastly, If you only care about one method, it's sometimes more convenient  making the assertion on an anonymous interface containing only the  method you rely on, ie `v, ok := x.(interface{ F() (int, error) } )`. Using anonymous interfaces can help decoupling your code from possible  dependencies, and can help guard your code from possible changes in  interfaces.

最后，如果您只关心一种方法，有时在仅包含您依赖的方法的匿名接口上进行断言会更方便，即 `v, ok := x.(interface{ F() (int, error) } )`。使用匿名接口有助于将您的代码与可能的依赖项分离，并有助于保护您的代码免受接口可能发生的变化。

## Asserting with concrete type T & Type Switches

## 使用具体类型 T 和类型开关断言

I will preface this section by introducing *two* similar error handling patterns that suffer from a couple of drawbacks  and pitfalls. It doesn’t mean they’re not common, though. Both of them  can be handy tools in small projects, but they don’t scale well.

我将通过介绍*两个*类似的错误处理模式来作为本节的开头，这些模式存在一些缺点和陷阱。但这并不意味着它们不常见。它们都可以成为小型项目中的方便工具，但它们的扩展性不佳。

The first one is the second kind of type assertion: Performing type assertion `x.(T)` with a concrete type `T`. It asserts the value of `x` is of type `T`, or it is convertible to type `T`.

第一种是第二种类型断言：使用具体类型 `T` 执行类型断言 `x.(T)`。它断言 `x` 的值属于 `T` 类型，或者它可以转换为类型 `T`。

The other one is the [Type Switch](https://golang.org/doc/effective_go.html#type_switch) pattern. Type switches combine a switch statement with type assertion, using the reserved `type` keyword. They are especially common in error handling, where knowing  the underlying type of an error variable can be very helpful.

另一种是 [Type Switch](https://golang.org/doc/effective_go.html#type_switch) 模式。类型开关使用保留的 `type` 关键字将 switch 语句与类型断言结合起来。它们在错误处理中特别常见，在这种情况下，了解错误变量的基础类型会非常有帮助。


The big drawback of both approaches is that they both lead to code coupling with their dependencies. Both examples need to be familiar with the `SomeErrorType` struct (which needs to be exported, obviously), and need to import the `mypkg` package.

这两种方法的最大缺点是它们都会导致代码与其依赖项耦合。这两个示例都需要熟悉 `SomeErrorType` 结构体（显然需要导出），并且需要导入 `mypkg` 包。

In both approaches, when handling your errors, you must be familiar with  the type and import its package. It gets worse when you are dealing with wrapped errors, where the cause of the error can be an error created in an internal dependency you are not, and should not be, aware of.

在这两种方法中，在处理错误时，您必须熟悉类型并导入其包。当您处理包装错误时，情况会变得更糟，其中错误的原因可能是在您没有也不应该意识到的内部依赖项中创建的错误。


Type switches differentiate between`*MyStruct` and `MyStruct`. So if you’re not sure if you are dealing with a pointer or an actual  instance of a struct, you’ll have to provide both. Moreover, just like  switches, cases in type switches do not fall through, but unlike  switches, usage of `fallthrough` is forbidden in type switches, so you’ll have to use comma and provide both options, which is easy to forget.

类型开关区分`*MyStruct` 和`MyStruct`。因此，如果您不确定要处理的是指针还是结构的实际实例，则必须同时提供两者。而且，就像开关一样，类型开关中的cases不会落空，但与开关不同的是，类型开关中禁止使用`fallthrough`，所以你必须使用逗号并提供两个选项，这很容易忘记。


# Wrapping up

#  总结

That’s it! You’re now familiar with errors and should be ready to tackle any  errors your Go app may throw (or actually return) your way!

就是这样！您现在已经熟悉了错误，并且应该准备好以您的方式处理 Go 应用程序可能抛出（或实际返回）的任何错误！

Both `errors` packages present simple yet important approaches to errors in Go, and  if they satisfy your needs, they are excellent choices. You can easily  implement your own custom error structs, and enjoy the benefits of Go’s  error handling when combining them with `pkg/errors`.

两个 `errors` 包都为 Go 中的错误提供了简单而重要的方法，如果它们满足您的需求，它们是很好的选择。您可以轻松实现自己的自定义错误结构，并在将它们与 `pkg/errors` 结合使用时享受 Go 错误处理的好处。

When you scale out of the simple errors, using type assertions properly can  be a great tool to handling different errors. Either by using type  switches or by asserting an error’s behavior and checking for interfaces it implements.

当您扩展简单错误时，正确使用类型断言可以成为处理不同错误的好工具。通过使用类型开关或通过断言错误的行为并检查它实现的接口。

## What’s next?

##  下一步是什么？

Go’s error handling is a very hot topic these days. Now that you’ve got the  basics, you may be interested in what the future holds in store for Go’s error handling!

Go 的错误处理现在是一个非常热门的话题。现在你已经掌握了基础知识，你可能会对 Go 错误处理的未来感兴趣！

It gets lots of attention in the upcoming Go 2 version, and you can already take a look at the [draft design](https://go.googlesource.com/proposal/+/master/design/go2draft.md). Also, During [dotGo 2019](https://www.dotgo.eu/), Marcel van Lohuizen had an excellent talk about the subject I just can't recommend enough — [*“Go 2 Error Values Today”*]( https://www.youtube.com/watch?v=SeVxmQl9Wmk).

它在即将发布的 Go 2 版本中受到了很多关注，您已经可以查看 [draft design](https://go.googlesource.com/proposal/+/master/design/go2draft.md)。此外，在 [dotGo 2019](https://www.dotgo.eu/) 期间，Marcel van Lohuizen 对这个主题进行了精彩的演讲，我实在是太推荐了 — [*“Go 2 Error Values Today”*]( https://www.youtube.com/watch?v=SeVxmQl9Wmk)。

There are plenty more approaches, tips and tricks, clearly, and there is no  possible way I can include them all in a single post! Regardless, I hope you’ve enjoyed this post, and I’ll see you on the next instalment of  the “*Before you go Go*” series!

显然，还有更多的方法、技巧和窍门，我不可能将它们全部包含在一个帖子中！无论如何，我希望你喜欢这篇文章，我会在“*走之前*”系列的下一部分再见！

[Alon Abadi](https://medium.com/@alonabadi?source=post_sidebar--------------------------post_sidebar-----------)

Back End Developer and Golang Enthusiast. I am a sucker for great code, and I like to hack whatever comes my way. 

后端开发人员和 Golang 爱好者。我是伟大代码的傻瓜，我喜欢破解我遇到的任何事情。

