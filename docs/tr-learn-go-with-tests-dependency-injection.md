# Dependency Injection

# 依赖注入

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/di)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/di)**

It is assumed that you have read the structs section before as some understanding of interfaces will be needed for this.

假设您之前已经阅读了结构部分，因为这需要对接口有一些了解。

There are _a lot_ of misunderstandings around dependency injection around the programming community. Hopefully, this guide will show you how

围绕编程社区的依赖注入存在_很多_误解。希望本指南将向您展示如何

* You don't need a framework
* It does not overcomplicate your design
* It facilitates testing
* It allows you to write great, general-purpose functions.

* 你不需要框架
* 它不会使您的设计过于复杂
* 方便测试
* 它允许您编写出色的通用函数。

We want to write a function that greets someone, just like we did in the hello-world chapter but this time we are going to be testing the _actual printing_.

我们想编写一个向某人打招呼的函数，就像我们在 hello-world 章节中所做的那样，但这次我们将测试_实际打印_。

Just to recap, here is what that function could look like

回顾一下，这是该功能的样子

```go
func Greet(name string) {
    fmt.Printf("Hello, %s", name)
}
```

But how can we test this? Calling `fmt.Printf` prints to stdout, which is pretty hard for us to capture using the testing framework.

但是我们如何测试呢？调用 `fmt.Printf` 会打印到标准输出，这对于我们使用测试框架很难捕捉到。

What we need to do is to be able to **inject** \(which is just a fancy word for pass in\) the dependency of printing.

我们需要做的是能够**注入**\（这只是传入\的一个花哨的词）打印的依赖。

**Our function doesn't need to care **_**where**_ or how the printing happens, so we should accept an **_**interface** _** rather than a concrete type.**

**我们的函数不需要关心**_**在哪里**_**或**_**如何**_打印发生，所以我们应该接受一个 接口 而不是具体的类型。

If we do that, we can then change the implementation to print to something we control so that we can test it. In "real life" you would inject in something that writes to stdout.

如果我们这样做，我们就可以将实现更改为打印到我们控制的内容，以便我们可以对其进行测试。在“现实生活”中，您会注入一些写入标准输出的内容。

If you look at the source code of [`fmt.Printf`](https://pkg.go.dev/fmt#Printf) you can see a way for us to hook in

如果你看一下 [`fmt.Printf`](https://pkg.go.dev/fmt#Printf) 的源代码，你可以看到一种让我们挂钩的方法

```go
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}
```

Interesting! Under the hood `Printf` just calls `Fprintf` passing in `os.Stdout`.

有趣的！在幕后，`Printf` 只是调用传入 `os.Stdout` 的 `Fprintf`。

What exactly _is_ an `os.Stdout`? What does `Fprintf` expect to get passed to it for the 1st argument?

什么是`os.Stdout`？ `Fprintf` 希望通过第一个参数传递给它什么？

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintf(format, a)
    n, err = w.Write(p.buf)
    p.free()
    return
}
```

An `io.Writer`

一个 `io.Writer`

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

From this we can infer that `os.Stdout` implements `io.Writer`; `Printf` passes `os.Stdout` to `Fprintf` which expects an `io.Writer`.

由此我们可以推断`os.Stdout`实现了`io.Writer`； `Printf` 将 `os.Stdout` 传递给需要 `io.Writer` 的 `Fprintf`。

As you write more Go code you will find this interface popping up a lot because it's a great general purpose interface for "put this data somewhere".

当你编写更多 Go 代码时，你会发现这个界面会弹出很多，因为它是一个很好的通用界面，用于“将这些数据放在某处”。

So we know under the covers we're ultimately using `Writer` to send our greeting somewhere. Let's use this existing abstraction to make our code testable and more reusable.

所以我们知道在幕后我们最终会使用 `Writer` 来向某个地方发送我们的问候。让我们使用这个现有的抽象来使我们的代码可测试和更可重用。

## Write the test first

## 先写测试

```go
func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Chris")

    got := buffer.String()
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

The `Buffer` type from the `bytes` package implements the `Writer` interface, because it has the method `Write(p []byte) (n int, err error)`.

`bytes` 包中的 `Buffer` 类型实现了 `Writer` 接口，因为它具有方法 `Write(p []byte) (n int, err error)`。

So we'll use it in our test to send in as our `Writer` and then we can check what was written to it after we invoke `Greet`

所以我们将在我们的测试中使用它作为我们的 `Writer` 发送，然后我们可以在调用 `Greet` 后检查写入的内容

## Try and run the test

## 尝试并运行测试

The test will not compile

测试不会编译

```text
./di_test.go:10:7: too many arguments in call to Greet
    have (*bytes.Buffer, string)
    want (string)
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

_Listen to the compiler_ and fix the problem.

_听编译器_并解决问题。

```go
func Greet(writer *bytes.Buffer, name string) {
    fmt.Printf("Hello, %s", name)
}
```

`Hello, Chris di_test.go:16: got '' want 'Hello, Chris'`

The test fails. Notice that the name is getting printed out, but it's going to stdout.

测试失败。请注意，名称正在打印出来，但它会输出到标准输出。

## Write enough code to make it pass

## 编写足够的代码使其通过

Use the writer to send the greeting to the buffer in our test. Remember `fmt.Fprintf` is like `fmt.Printf` but instead takes a `Writer` to send the string to, whereas `fmt.Printf` defaults to stdout.

在我们的测试中使用 writer 将问候语发送到缓冲区。请记住，`fmt.Fprintf` 类似于 `fmt.Printf`，但需要一个 `Writer` 来发送字符串，而 `fmt.Printf` 默认为 stdout。

```go
func Greet(writer *bytes.Buffer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}
```

The test now passes.

现在测试通过了。

## Refactor

## 重构

Earlier the compiler told us to pass in a pointer to a `bytes.Buffer`. This is technically correct but not very useful.

早些时候，编译器告诉我们传入一个指向 `bytes.Buffer` 的指针。这在技术上是正确的，但不是很有用。

To demonstrate this, try wiring up the `Greet` function into a Go application where we want it to print to stdout.

为了演示这一点，尝试将 `Greet` 函数连接到我们希望它打印到标准输出的 Go 应用程序中。

```go
func main() {
    Greet(os.Stdout, "Elodie")
}
```

`./di.go:14:7: cannot use os.Stdout (type *os.File) as type *bytes.Buffer in argument to Greet` 

`./di.go:14:7: 不能在 Greet` 的参数中使用 os.Stdout（类型 *os.File）作为类型 *bytes.Buffer

As discussed earlier `fmt.Fprintf` allows you to pass in an `io.Writer` which we know both `os.Stdout` and `bytes.Buffer` implement.

正如前面讨论的，`fmt.Fprintf` 允许你传入一个 `io.Writer`，我们知道 `os.Stdout` 和 `bytes.Buffer` 实现。

If we change our code to use the more general purpose interface we can now use it in both tests and in our application.

如果我们更改代码以使用更通用的接口，我们现在可以在测试和应用程序中使用它。

```go
package main

import (
    "fmt"
    "os"
    "io"
)

func Greet(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
    Greet(os.Stdout, "Elodie")
}
```

## More on io.Writer

## 更多关于 io.Writer

What other places can we write data to using `io.Writer`? Just how general purpose is our `Greet` function?

我们可以使用 io.Writer 将数据写入哪些其他地方？我们的“Greet”功能有多通用？

### The Internet

###  互联网

Run the following

运行以下

```go
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func Greet(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
    Greet(w, "world")
}

func main() {
    log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
```

Run the program and go to [http://localhost:5000](http://localhost:5000). You'll see your greeting function being used.

运行程序并转到 [http://localhost:5000](http://localhost:5000)。你会看到你的问候功能被使用。

HTTP servers will be covered in a later chapter so don't worry too much about the details.

HTTP 服务器将在后面的章节中介绍，所以不要太担心细节。

When you write an HTTP handler, you are given an `http.ResponseWriter` and the `http.Request` that was used to make the request. When you implement your server you _write_ your response using the writer.

当您编写 HTTP 处理程序时，您将获得一个 `http.ResponseWriter` 和用于发出请求的 `http.Request`。当你实现你的服务器时，你_写_你的响应使用作者。

You can probably guess that `http.ResponseWriter` also implements `io.Writer` so this is why we could re-use our `Greet` function inside our handler.

你可能会猜到 `http.ResponseWriter` 也实现了 `io.Writer`，所以这就是我们可以在我们的处理程序中重用我们的 `Greet` 函数的原因。

## Wrapping up

##  总结

Our first round of code was not easy to test because it wrote data to somewhere we couldn't control.

我们的第一轮代码不容易测试，因为它会将数据写入我们无法控制的地方。

_Motivated by our tests_ we refactored the code so we could control _where_ the data was written by **injecting a dependency** which allowed us to:

_受我们的测试的启发_我们重构了代码，以便我们可以通过**注入依赖项**来控制数据的写入位置，这使我们能够：

* **Test our code** If you can't test a function _easily_, it's usually because of dependencies hard-wired into a function _or_ global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency \(via an interface\) which you can then mock out with something you can control in your tests.
* **Separate our concerns**, decoupling _where the data goes_ from _how to generate it_. If you ever feel like a method/function has too many responsibilities \(generating data _and_ writing to a db? handling HTTP requests _and_ doing domain level logic?\) DI is probably going to be the tool you need.
* **Allow our code to be re-used in different contexts** The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.

* **测试我们的代码** 如果你不能_轻松_测试一个函数，通常是因为依赖硬连接到一个函数_或_全局状态。例如，如果您有一个由某种服务层使用的全局数据库连接池，则可能很难测试，并且它们的运行速度会很慢。 DI 会激励你注入一个数据库依赖（通过一个接口），然后你可以用你可以在测试中控制的东西来模拟。
* **分离我们的关注点**，将_数据去向_与_如何生成_分离。如果你曾经觉得一个方法/函数有太多的责任（生成数据_和_写入数据库？处理 HTTP 请求_和_执行域级逻辑？）DI 可能会成为你需要的工具。
* **允许我们的代码在不同的上下文中重用** 我们的代码可以在其中使用的第一个“新”上下文是内部测试。但进一步来说，如果有人想用你的函数尝试一些新的东西，他们可以注入自己的依赖项。

### What about mocking? I hear you need that for DI and also it's evil

### 嘲笑呢？我听说你需要 DI 并且这很邪恶

Mocking will be covered in detail later \(and it's not evil\). You use mocking to replace real things you inject with a pretend version that you can control and inspect in your tests. In our case though, the standard library had something ready for us to use.

稍后将详细介绍模拟\（它不是邪恶的\）。您可以使用模拟将您注入的真实内容替换为您可以在测试中控制和检查的假想版本。但在我们的例子中，标准库已经准备好供我们使用。

### The Go standard library is really good, take time to study it

### Go 标准库真的不错，有时间研究一下

By having some familiarity with the `io.Writer` interface we are able to use `bytes.Buffer` in our test as our `Writer` and then we can use other `Writer`s from the standard library to use our function in a command line app or in web server.

通过熟悉 `io.Writer` 接口，我们可以在测试中使用 `bytes.Buffer` 作为我们的 `Writer`，然后我们可以使用标准库中的其他 `Writer` 来使用我们的函数命令行应用程序或在 Web 服务器中。

The more familiar you are with the standard library the more you'll see these general purpose interfaces which you can then re-use in your own code to make your software reusable in a number of contexts.

您对标准库越熟悉，就越能看到这些通用接口，然后您可以在自己的代码中重用这些接口，使您的软件可在许多上下文中重用。

This example is heavily influenced by a chapter in [The Go Programming language](https://www.amazon.co.uk/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440), so if you enjoyed this , go buy it! 

这个例子深受 [The Go Programming language](https://www.amazon.co.uk/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440) 一章的影响，所以如果你喜欢这个，去买吧！

