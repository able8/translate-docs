# Hello, World

#  你好，世界

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/hello-world)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/hello-world)**

It is traditional for your first program in a new language to be [Hello, World](https://en.m.wikipedia.org/wiki/%22Hello,_World!%22_program).

传统上，您使用新语言编写的第一个程序是 [Hello, World](https://en.m.wikipedia.org/wiki/%22Hello,_World!%22_program)。

- Create a folder wherever you like
- Put a new file in it called `hello.go` and put the following code inside it

- 在任何你喜欢的地方创建一个文件夹
- 在其中放入一个名为 `hello.go` 的新文件，并将以下代码放入其中

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world")
}
```

To run it type `go run hello.go`.

要运行它，请输入 `go run hello.go`。

## How it works

##  这个怎么运作

When you write a program in Go you will have a `main` package defined with a `main` func inside it. Packages are ways of grouping up related Go code together.

当您在 Go 中编写程序时，您将定义一个 `main` 包，其中包含一个 `main` func。包是将相关的 Go 代码组合在一起的方式。

The `func` keyword is how you define a function with a name and a body.

`func` 关键字用于定义具有名称和主体的函数。

With `import "fmt"` we are importing a package which contains the `Println` function that we use to print.

使用 `import "fmt"` 我们导入一个包，其中包含我们用来打印的 `Println` 函数。

## How to test

## 如何测试

How do you test this? It is good to separate your "domain" code from the outside world \(side-effects\). The `fmt.Println` is a side effect \(printing to stdout\) and the string we send in is our domain.

你如何测试这个？将您的“域”代码与外界（副作用）分开是很好的。 `fmt.Println` 是一个副作用（打印到标准输出），我们发送的字符串是我们的域。

So let's separate these concerns so it's easier to test

所以让我们把这些问题分开，这样更容易测试

```go
package main

import "fmt"

func Hello() string {
    return "Hello, world"
}

func main() {
    fmt.Println(Hello())
}
```

We have created a new function again with `func` but this time we've added another keyword `string` in the definition. This means this function returns a `string`.

我们再次使用 `func` 创建了一个新函数，但这次我们在定义中添加了另一个关键字 `string`。这意味着这个函数返回一个`string`。

Now create a new file called `hello_test.go` where we are going to write a test for our `Hello` function

现在创建一个名为 `hello_test.go` 的新文件，我们将在其中为我们的 `Hello` 函数编写测试

```go
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, world"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

## Go modules?

## 模块？

The next step is to run the tests. Enter `go test` in your terminal. If the tests pass, then you are probably using an earlier version of Go. However, if you are using Go 1.16 or later, then the tests will likely not run at all. Instead, you will see an error message like this in the terminal:

下一步是运行测试。在终端中输入“go test”。如果测试通过，那么您可能使用的是早期版本的 Go。但是，如果您使用的是 Go 1.16 或更高版本，则测试可能根本无法运行。相反，您将在终端中看到如下错误消息：

```shell
$ go test
go: cannot find main module;see 'go help modules'
```

What's the problem? In a word, [modules](https://blog.golang.org/go116-module-changes). Luckily, the problem is easy to fix. Enter `go mod init hello` in your terminal. That will create a new file with the following contents:

有什么问题？总之，[模块](https://blog.golang.org/go116-module-changes)。幸运的是，这个问题很容易解决。在终端中输入 `go mod init hello`。这将创建一个包含以下内容的新文件：

```go
module hello

go 1.16
```

This file tells the `go` tools essential information about your code. If you planned to distribute your application, you would include where the code was available for download as well as information about dependencies. For now, your module file is minimal, and you can leave it that way. To read more about modules, [you can check out the reference in the Golang documentation](https://golang.org/doc/modules/gomod-ref). We can get back to testing and learning Go now since the tests should run, even on Go 1.16.

这个文件告诉 `go` 工具关于你的代码的基本信息。如果您计划分发您的应用程序，您将包括可下载代码的位置以及有关依赖项的信息。目前，您的模块文件是最小的，您可以保留它。要阅读有关模块的更多信息，[您可以查看 Golang 文档中的参考](https://golang.org/doc/modules/gomod-ref)。我们现在可以重新开始测试和学习 Go，因为测试应该可以运行，即使在 Go 1.16 上也是如此。

In future chapters you will need to run `go mod init SOMENAME` in each new folder before running commands like `go test` or `go build`.

在以后的章节中，您需要在每个新文件夹中运行“go mod init SOMENAME”，然后才能运行“go test”或“go build”等命令。

## Back to Testing

## 回到测试

Run `go test` in your terminal. It should've passed! Just to check, try deliberately breaking the test by changing the `want` string.

在终端中运行 `go test`。应该过去了！只是为了检查，尝试通过更改 `want` 字符串故意破坏测试。

Notice how you have not had to pick between multiple testing frameworks and then figure out how to install. Everything you need is built in to the language and the syntax is the same as the rest of the code you will write.

请注意，您不必在多个测试框架之间进行选择，然后弄清楚如何安装。你需要的一切都内置在语言中，语法与你将编写的其余代码相同。

### Writing tests

### 编写测试

Writing a test is just like writing a function, with a few rules

编写测试就像编写函数一样，有一些规则

* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `Test`
* The test function takes one argument only `t *testing.T`
* In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file



* 它需要在一个名称类似于`xxx_test.go`的文件中
* 测试函数必须以单词`Test`开头
* 测试函数只接受一个参数 `t *testing.T`
* 为了使用 `*testing.T` 类型，你需要 `import "testing"`，就像我们在另一个文件中对 `fmt` 所做的一样

For now, it's enough to know that your `t` of type `*testing.T` is your "hook" into the testing framework so you can do things like `t.Fail()` when you want to fail.

现在，知道你的 `*testing.T` 类型的 `t` 是你进入测试框架的“钩子”就足够了，所以当你想失败时，你可以做类似 `t.Fail()` 的事情。

We've covered some new topics:

我们讨论了一些新主题：

#### `if`
If statements in Go are very much like other programming languages.
Go 中的 if 语句与其他编程语言非常相似。

#### Declaring variables

#### 声明变量

We're declaring some variables with the syntax `varName := value`, which lets us re-use some values in our test for readability.

我们用 `varName := value` 语法声明了一些变量，这让我们可以在测试中重用一些值来提高可读性。

#### `t.Errorf` 

We are calling the `Errorf` _method_ on our `t` which will print out a message and fail the test. The `f` stands for format which allows us to build a string with values inserted into the placeholder values `%q`. When you made the test fail it should be clear how it works.

我们在我们的 `t` 上调用了 `Errorf` _method_，它将打印出一条消息并使测试失败。 `f` 代表格式，它允许我们构建一个字符串，其中的值插入到占位符值 `%q` 中。当您使测试失败时，应该清楚它是如何工作的。

You can read more about the placeholder strings in the [fmt go doc](https://golang.org/pkg/fmt/#hdr-Printing). For tests `%q` is very useful as it wraps your values in double quotes.

您可以在 [fmt go doc](https://golang.org/pkg/fmt/#hdr-Printing) 中阅读有关占位符字符串的更多信息。对于测试，`%q` 非常有用，因为它将您的值用双引号括起来。

We will later explore the difference between methods and functions.

稍后我们将探讨方法和函数之间的区别。

### Go doc



Another quality of life feature of Go is the documentation. You can launch the docs locally by running `godoc -http :8000`. If you go to [localhost:8000/pkg](http://localhost:8000/pkg) you will see all the packages installed on your system.

Go 的另一个生活质量特性是文档。您可以通过运行 `godoc -http :8000` 在本地启动文档。如果您转到 [localhost:8000/pkg](http://localhost:8000/pkg)，您将看到系统上安装的所有软件包。

The vast majority of the standard library has excellent documentation with examples. Navigating to [http://localhost:8000/pkg/testing/](http://localhost:8000/pkg/testing/) would be worthwhile to see what's available to you.

绝大多数标准库都有带有示例的优秀文档。导航到 [http://localhost:8000/pkg/testing/](http://localhost:8000/pkg/testing/) 会很值得看看有什么可用的。

If you don't have `godoc` command, then maybe you are using the newer version of Go (1.14 or later) which is [no longer including `godoc`](https://golang.org/doc/go1.14#godoc). You can manually install it with `go get golang.org/x/tools/cmd/godoc`.

如果您没有 `godoc` 命令，那么您可能使用的是较新版本的 Go（1.14 或更高版本），它[不再包括 `godoc`](https://golang.org/doc/go1.14#godoc)。您可以使用 `go get golang.org/x/tools/cmd/godoc` 手动安装它。

### Hello, YOU



Now that we have a test we can iterate on our software safely.

现在我们有了一个测试，我们可以安全地迭代我们的软件。

In the last example we wrote the test _after_ the code had been written just so you could get an example of how to write a test and declare a function. From this point on we will be _writing tests first_.

在上一个示例中，我们在编写测试之后编写了代码，这样您就可以获得如何编写测试和声明函数的示例。从现在开始，我们将_首先编写测试_。

Our next requirement is to let us specify the recipient of the greeting.

我们的下一个要求是让我们指定问候的收件人。

Let's start by capturing these requirements in a test. This is basic test driven development and allows us to make sure our test is _actually_ testing what we want. When you retrospectively write tests there is the risk that your test may continue to pass even if the code doesn't work as intended.

让我们从在测试中捕获这些需求开始。这是基本的测试驱动开发，允许我们确保我们的测试_实际上_测试我们想要的。当您回顾性地编写测试时，即使代码没有按预期工作，您的测试也可能会继续通过。

```go
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Chris")
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

Now run `go test`, you should have a compilation error

现在运行`go test`，你应该有一个编译错误

```text
./hello_test.go:6:18: too many arguments in call to Hello
    have (string)
    want ()
```

When using a statically typed language like Go it is important to _listen to the compiler_. The compiler understands how your code should snap together and work so you don't have to.

当使用像 Go 这样的静态类型语言时，_听编译器_很重要。编译器了解您的代码应该如何组合在一起并工作，因此您不必这样做。

In this case the compiler is telling you what you need to do to continue. We have to change our function `Hello` to accept an argument.

在这种情况下，编译器会告诉您需要做什么才能继续。我们必须改变我们的函数 `Hello` 来接受一个参数。

Edit the `Hello` function to accept an argument of type string

编辑 `Hello` 函数以接受字符串类型的参数

```go
func Hello(name string) string {
    return "Hello, world"
}
```

If you try and run your tests again your `hello.go` will fail to compile because you're not passing an argument. Send in "world" to make it compile.

如果您尝试再次运行测试，您的 `hello.go` 将无法编译，因为您没有传递参数。发送“world”以使其编译。

```go
func main() {
    fmt.Println(Hello("world"))
}
```

Now when you run your tests you should see something like

现在，当您运行测试时，您应该看到类似

```text
hello_test.go:10: got 'Hello, world' want 'Hello, Chris''
```

We finally have a compiling program but it is not meeting our requirements according to the test.

我们终于有了一个编译程序，但根据测试，它不符合我们的要求。

Let's make the test pass by using the name argument and concatenate it with `Hello,`

让我们使用 name 参数使测试通过并将其与 `Hello,` 连接起来

```go
func Hello(name string) string {
    return "Hello, " + name
}
```

When you run the tests they should now pass. Normally as part of the TDD cycle we should now _refactor_.

当您运行测试时，它们现在应该会通过。通常作为 TDD 周期的一部分，我们现在应该_重构_。

### A note on source control

### 关于源代码控制的说明

At this point, if you are using source control \(which you should!\) I would `commit` the code as it is. We have working software backed by a test.

在这一点上，如果您正在使用源代码管理\（你应该！\）我会`commit` 代码原样。我们有经过测试支持的工作软件。

I _wouldn't_ push to master though, because I plan to refactor next. It is nice to commit at this point in case you somehow get into a mess with refactoring - you can always go back to the working version.

不过，我_wouldn't_ 推动掌握，因为我计划接下来重构。很好在这一点上提交，以防你以某种方式在重构中陷入困境 - 你总是可以回到工作版本。

There's not a lot to refactor here, but we can introduce another language feature, _constants_.

这里没有太多需要重构的地方，但我们可以引入另一个语言特性，_constants_。

### Constants

### 常量

Constants are defined like so

常量定义如下

```go
const englishHelloPrefix = "Hello, "
```

We can now refactor our code

我们现在可以重构我们的代码

```go
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
    return englishHelloPrefix + name
}
```

After refactoring, re-run your tests to make sure you haven't broken anything.

重构后，重新运行您的测试以确保您没有破坏任何东西。

Constants should improve performance of your application as it saves you creating the `"Hello, "` string instance every time `Hello` is called. 

常量应该可以提高应用程序的性能，因为它可以节省您每次调用 `Hello` 时都创建 `"Hello, "` 字符串实例。

To be clear, the performance boost is incredibly negligible for this example! But it's worth thinking about creating constants to capture the meaning of values and sometimes to aid performance.

需要明确的是，对于这个例子，性能提升是非常微不足道的！但是值得考虑创建常量来捕获值的含义，有时还有助于提高性能。

## Hello, world... again

## 你好，世界...再次

The next requirement is when our function is called with an empty string it defaults to printing "Hello, World", rather than "Hello, ".

下一个要求是当我们使用空字符串调用我们的函数时，它默认打印“Hello, World”，而不是“Hello,”。

Start by writing a new failing test

首先编写一个新的失败测试

```go
func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
}
```

Here we are introducing another tool in our testing arsenal, subtests. Sometimes it is useful to group tests around a "thing" and then have subtests describing different scenarios.

在这里，我们将在我们的测试库中引入另一个工具，子测试。有时，围绕“事物”对测试进行分组，然后使用描述不同场景的子测试是很有用的。

A benefit of this approach is you can set up shared code that can be used in the other tests.

这种方法的一个好处是您可以设置可用于其他测试的共享代码。

There is repeated code when we check if the message is what we expect.

当我们检查消息是否符合我们的预期时，会有重复的代码。

Refactoring is not _just_ for the production code!

重构不_只是_用于生产代码！

It is important that your tests _are clear specifications_ of what the code needs to do.

重要的是您的测试_清楚地说明_代码需要做什么。

We can and should refactor our tests.

我们可以而且应该重构我们的测试。

```go
func TestHello(t *testing.T) {
    assertCorrectMessage := func(t testing.TB, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })
    t.Run("empty string defaults to 'World'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
}
```

What have we done here?

我们在这里做了什么？

We've refactored our assertion into a function. This reduces duplication and improves readability of our tests. In Go you can declare functions inside other functions and assign them to variables. You can then call them, just like normal functions. We need to pass in `t *testing.T` so that we can tell the test code to fail when we need to.

我们已经将断言重构为一个函数。这减少了重复并提高了我们测试的可读性。在 Go 中，您可以在其他函数中声明函数并将它们分配给变量。然后您可以调用它们，就像普通函数一样。我们需要传入 `t *testing.T` 以便我们可以在需要时告诉测试代码失败。

For helper functions, it's a good idea to accept a `testing.TB` which is an interface that `*testing.T` and `*testing.B` both satisfy, so you can call helper functions from a test, or a benchmark .

对于辅助函数，接受`testing.TB` 是一个好主意，它是`*testing.T` 和`*testing.B` 都满足的接口，因此您可以从测试或基准测试中调用辅助函数.

`t.Helper()` is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our _function call_ rather than inside our test helper. This will help other developers track down problems easier. If you still don't understand, comment it out, make a test fail and observe the test output. Comments in Go are a great way to add additional information to your code, or in this case, a quick way to tell the compiler to ignore a line. You can comment out the `t.Helper()` code by adding two forward slashes `//` at the beginning of the line. You should see that line turn grey or change to another color than the rest of your code to indicate it's now commented out.

需要 `t.Helper()` 来告诉测试套件这个方法是一个帮助器。通过这样做，当它失败时，报告的行号将在我们的 _function call_ 中，而不是在我们的测试助手中。这将帮助其他开发人员更轻松地追踪问题。如果还是不明白，注释掉，做一个测试失败，观察测试输出。 Go 中的注释是向代码添加附加信息的好方法，或者在这种情况下，是告诉编译器忽略一行的快速方法。您可以通过在行的开头添加两个正斜杠 `//` 来注释掉 `t.Helper()` 代码。您应该会看到该行变为灰色或更改为与代码其余部分不同的另一种颜色，以表明它现在已被注释掉。

Now that we have a well-written failing test, let's fix the code, using an `if`.

现在我们有一个编写良好的失败测试，让我们使用 `if` 修复代码。

```go
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
    if name == "" {
        name = "World"
    }
    return englishHelloPrefix + name
}
```

If we run our tests we should see it satisfies the new requirement and we haven't accidentally broken the other functionality.

如果我们运行我们的测试，我们应该看到它满足新的要求，并且我们没有意外破坏其他功能。

### Back to source control

### 回到源代码管理

Now we are happy with the code I would amend the previous commit so we only
check in the lovely version of our code with its test.

现在我们对代码很满意，我会修改之前的提交，所以我们只
用它的测试检查我们代码的可爱版本。

### Discipline

###  纪律

Let's go over the cycle again

让我们再次循环

* Write a test
* Make the compiler pass
* Run the test, see that it fails and check the error message is meaningful
* Write enough code to make the test pass
* Refactor

* 写一个测试
* 使编译器通过
* 运行测试，查看是否失败并检查错误信息是否有意义
* 编写足够的代码使测试通过
* 重构

On the face of it this may seem tedious but sticking to the feedback loop is important.

从表面上看，这似乎很乏味，但坚持反馈循环很重要。

Not only does it ensure that you have _relevant tests_, it helps ensure _you design good software_ by refactoring with the safety of tests.

它不仅可以确保您拥有_相关测试_，还可以通过重构测试的安全性来帮助确保_您设计出好的软件_。

Seeing the test fail is an important check because it also lets you see what the error message looks like. As a developer it can be very hard to work with a codebase when failing tests do not give a clear idea as to what the problem is. 

看到测试失败是一项重要的检查，因为它还可以让您看到错误消息的样子。作为开发人员，如果失败的测试不能清楚地说明问题是什么，那么使用代码库可能会非常困难。

By ensuring your tests are _fast_ and setting up your tools so that running tests is simple you can get in to a state of flow when writing your code.

通过确保您的测试_快速_并设置您的工具，以便运行测试变得简单，您可以在编写代码时进入流程状态。

By not writing tests you are committing to manually checking your code by running your software which breaks your state of flow and you won't be saving yourself any time, especially in the long run.

通过不编写测试，您承诺通过运行您的软件来手动检查您的代码，这会破坏您的流程状态，并且您不会节省任何时间，尤其是从长远来看。

## Keep going! More requirements

##  继续！更多要求

Goodness me, we have more requirements. We now need to support a second parameter, specifying the language of the greeting. If a language is passed in that we do not recognise, just default to English.

天哪，我们有更多的要求。我们现在需要支持第二个参数，指定问候语的语言。如果传入一种我们无法识别的语言，则默认为英语。

We should be confident that we can use TDD to flesh out this functionality easily!

我们应该相信我们可以使用 TDD 轻松充实这个功能！

Write a test for a user passing in Spanish. Add it to the existing suite.

为通过西班牙语的用户编写测试。将其添加到现有套件中。

```go
     t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
```

Remember not to cheat! _Test first_. When you try and run the test, the compiler _should_ complain because you are calling `Hello` with two arguments rather than one.

切记不要作弊！ _先测试_。当您尝试运行测试时，编译器_应该_抱怨，因为您使用两个参数而不是一个参数调用 `Hello`。

```text
./hello_test.go:27:19: too many arguments in call to Hello
    have (string, string)
    want (string)
```

Fix the compilation problems by adding another string argument to `Hello`

通过向 `Hello` 添加另一个字符串参数来修复编译问题

```go
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
    return englishHelloPrefix + name
}
```

When you try and run the test again it will complain about not passing through enough arguments to `Hello` in your other tests and in `hello.go`

当你再次尝试运行测试时，它会抱怨没有在你的其他测试和 `hello.go` 中传递足够的参数给 `Hello`

```text
./hello.go:15:19: not enough arguments in call to Hello
    have (string)
    want (string, string)
```

Fix them by passing through empty strings. Now all your tests should compile _and_ pass, apart from our new scenario

通过传递空字符串来修复它们。现在，除了我们的新场景之外，您所有的测试都应该编译通过

```text
hello_test.go:29: got 'Hello, Elodie' want 'Hola, Elodie'
```

We can use `if` here to check the language is equal to "Spanish" and if so change the message

我们可以在这里使用 `if` 来检查语言是否等于“西班牙语”，如果是，则更改消息

```go
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    if language == "Spanish" {
        return "Hola, " + name
    }
    return englishHelloPrefix + name
}
```

The tests should now pass.

测试现在应该通过了。

Now it is time to _refactor_. You should see some problems in the code, "magic" strings, some of which are repeated. Try and refactor it yourself, with every change make sure you re-run the tests to make sure your refactoring isn't breaking anything.

现在是时候_重构_了。你应该在代码中看到一些问题，“魔术”字符串，其中一些重复。尝试自己重构它，每次更改时请确保重新运行测试以确保重构不会破坏任何内容。

```go
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    if language == spanish {
        return spanishHelloPrefix + name
    }
    return englishHelloPrefix + name
}
```

### French

###  法语

* Write a test asserting that if you pass in `"French"` you get `"Bonjour, "`
* See it fail, check the error message is easy to read
* Do the smallest reasonable change in the code

* 编写一个测试，断言如果你传入 `"French"` 你会得到 `"Bonjour, "`
* 看不成功，查看错误信息易读
* 对代码做最小的合理改动

You may have written something that looks roughly like this

你可能写了一些看起来像这样的东西

```go
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    if language == spanish {
        return spanishHelloPrefix + name
    }
    if language == french {
        return frenchHelloPrefix + name
    }
    return englishHelloPrefix + name
}
```

## `switch`

## `开关`

When you have lots of `if` statements checking a particular value it is common to use a `switch` statement instead. We can use `switch` to refactor the code to make it easier to read and more extensible if we wish to add more language support later

当您有很多 `if` 语句检查特定值时，通常使用 `switch` 语句代替。如果我们希望稍后添加更多语言支持，我们可以使用 `switch` 来重构代码，使其更易于阅读和扩展

```go
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    prefix := englishHelloPrefix

    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    }

    return prefix + name
}
```

Write a test to now include a greeting in the language of your choice and you should see how simple it is to extend our _amazing_ function.

编写一个测试以包含您选择的语言的问候语，您应该会看到扩展我们的 _amazing_ 函数是多么简单。

### one...last...refactor?

### 一个...最后...重构？

You could argue that maybe our function is getting a little big. The simplest refactor for this would be to extract out some functionality into another function.

你可能会争辩说，也许我们的函数变得有点大了。对此最简单的重构是将某些功能提取到另一个功能中。

```go
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}
```

A few new concepts:

几个新概念：

* In our function signature we have made a _named return value_ `(prefix string)`. 

* 在我们的函数签名中，我们创建了一个 _named 返回值_ `（前缀字符串）`。

* This will create a variable called `prefix` in your function.
   * It will be assigned the "zero" value. This depends on the type, for example `int`s are 0 and for `string`s it is `""`.
     * You can return whatever it's set to by just calling `return` rather than `return prefix`.
   * This will display in the Go Doc for your function so it can make the intent of your code clearer.
* `default` in the switch case will be branched to if none of the other `case` statements match.
* The function name starts with a lowercase letter. In Go public functions start with a capital letter and private ones start with a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this function private.

* 这将在您的函数中创建一个名为 `prefix` 的变量。
  * 它将被分配“零”值。这取决于类型，例如`int`s 是0，而`string`s 是`""`。
    * 你可以通过调用 `return` 而不是 `return prefix` 来返回它设置的任何值。
  * 这将显示在您的函数的 Go Doc 中，以便它可以使您的代码意图更清晰。
* 如果其他 `case` 语句都不匹配，则 switch case 中的 `default` 将被分支到。
* 函数名以小写字母开头。在 Go 中，公共函数以大写字母开头，私有函数以小写字母开头。我们不希望我们的算法内部暴露给外界，所以我们将此函数设为私有。

## Wrapping up

##  总结

Who knew you could get so much out of `Hello, world`?

谁知道你可以从“你好，世界”中得到这么多？

By now you should have some understanding of:

现在你应该对以下内容有一些了解：

### Some of Go's syntax around

### 一些 Go 的语法

* Writing tests
* Declaring functions, with arguments and return types
* `if`, `const` and `switch`
* Declaring variables and constants

* 编写测试
* 声明函数，带参数和返回类型
* `if`、`const` 和 `switch`
* 声明变量和常量

### The TDD process and _why_ the steps are important

### TDD 过程和_为什么_步骤很重要

* _Write a failing test and see it fail_ so we know we have written a _relevant_ test for our requirements and seen that it produces an _easy to understand description of the failure_
* Writing the smallest amount of code to make it pass so we know we have working software
* _Then_ refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with

* _编写一个失败的测试并看到它失败_所以我们知道我们已经为我们的需求编写了一个_相关的_测试并且看到它产生了一个_易于理解的失败描述_
* 编写最少的代码以使其通过，以便我们知道我们有可以运行的软件
* _Then_ 重构，以我们测试的安全性为后盾，以确保我们拥有易于使用的精心设计的代码

In our case we've gone from `Hello()` to `Hello("name")`, to `Hello("name", "French")` in small, easy to understand steps.

在我们的例子中，我们已经从`Hello()` 到`Hello("name")`，再到`Hello("name", "French")` 的简单易懂的步骤。

This is of course trivial compared to "real world" software but the principles still stand. TDD is a skill that needs practice to develop, but by breaking problems down into smaller components that you can test, you will have a much easier time writing software. 

与“真实世界”软件相比，这当然是微不足道的，但原则仍然存在。 TDD 是一项需要练习才能开发的技能，但是通过将问题分解为可以测试的较小组件，您将可以更轻松地编写软件。

