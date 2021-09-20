# Handling Panics in Go

# 在 Go 中处理恐慌

October  3, 2019 40.2k views

2019 年 10 月 3 日 40.2k 观看次数

### Introduction

###  介绍

Errors that a program encounters fall into two broad categories:  those the programmer has anticipated and those the programmer has not. The `error` interface that we have covered in our previous two articles on [error handling](https://www.digitalocean.com/community/tutorials/handling-errors-in-go) largely deal with errors that we expect as we are writing Go programs. The `error` interface even allows us to acknowledge the rare possibility of an  error occurring from function calls, so we can respond appropriately in  those situations.

程序遇到的错误分为两大类：程序员已经预料到的错误和程序员没有预料到的错误。我们在前两篇关于 [错误处理](https://www.digitalocean.com/community/tutorials/handling-errors-in-go) 的文章中介绍的 `error` 接口主要处理我们期望的错误我们正在编写 Go 程序。 `error` 接口甚至允许我们承认函数调用发生错误的罕见可能性，因此我们可以在这些情况下做出适当的响应。

Panics fall into the second category of errors, which are  unanticipated by the programmer. These unforeseen errors lead a program  to spontaneously terminate and exit the running Go program. Common  mistakes are often responsible for creating panics. In this tutorial,  we’ll examine a few ways that common operations can produce panics in  Go, and we’ll also see ways to avoid those panics. We'll also use [`defer`](https://www.digitalocean.com/community/tutorials/understanding-defer-in-go) statements along with the `recover` function to capture panics before they have a chance to unexpectedly terminate our running Go programs.

恐慌属于第二类错误，这是程序员无法预料的。这些不可预见的错误导致程序自发地终止并退出正在运行的 Go 程序。常见的错误通常是造成恐慌的原因。在本教程中，我们将研究常见操作在 Go 中产生恐慌的几种方式，我们还将了解避免这些恐慌的方法。我们还将使用 [`defer`](https://www.digitalocean.com/community/tutorials/understanding-defer-in-go) 语句和 `recover` 函数来捕获恐慌，以免它们有机会意外终止我们正在运行的 Go 程序。

## Understanding Panics

## 理解恐慌

There are certain operations in Go that automatically return panics and stop the program. Common operations include indexing an [array](https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go#arrays) beyond its capacity, performing type assertions, calling methods on nil pointers , incorrectly using mutexes, and attempting to work with closed channels. Most of these situations result from mistakes made while  programming that the compiler has no ability to detect while compiling  your program.

Go 中有某些操作会自动返回恐慌并停止程序。常见操作包括索引 [array](https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go#arrays) 超出其容量、执行类型断言、调用 nil 指针上的方法，错误地使用互斥锁，并尝试使用封闭的通道。大多数这些情况是由于在编程时出现的错误造成的，编译器在编译程序时无法检测到这些错误。

Since panics include detail that is useful for resolving an issue,  developers commonly use panics as an indication that they have made a  mistake during a program’s development.

由于恐慌包括对解决问题有用的细节，开发人员通常使用恐慌作为他们在程序开发过程中犯了错误的指示。

### Out of Bounds Panics

### 越界恐慌

When you attempt to access an index beyond the length of a slice or  the capacity of an array, the Go runtime will generate a panic.

当您尝试访问超出切片长度或数组容量的索引时，Go 运行时将产生恐慌。

The following example makes the common mistake of attempting to  access the last element of a slice using the length of the slice  returned by the `len` builtin. Try running this code to see why this might produce a panic:

下面的示例犯了一个常见的错误，即尝试使用 `len` 内置函数返回的切片长度来访问切片的最后一个元素。尝试运行此代码以查看为什么这可能会产生恐慌：

```go
package main

import (
    "fmt"
)

func main() {
    names := []string{
        "lobster",
        "sea urchin",
        "sea cucumber",
    }
    fmt.Println("My favorite sea creature is:", names[len(names)])
}
```

 

This will have the following output:

这将有以下输出：

```
Outputpanic: runtime error: index out of range [3] with length 3

goroutine 1 [running]:
main.main()
    /tmp/sandbox879828148/prog.go:13 +0x20
```

The name of the panic’s output provides a hint: `panic: runtime error: index out of range`. We created a slice with three sea creatures. We then tried to get the  last element of the slice by indexing that slice with the length of the  slice using the `len` builtin function. Remember that slices  and arrays are zero-based; so the first element is zero and the last  element in this slice is at index `2`. Since we try to access the slice at the third index, `3`, there is no element in the slice to return because it is beyond the  bounds of the slice. The runtime has no option but to terminate and exit since we have asked it to do something impossible. Go also can’t prove  during compilation that this code will try to do this, so the compiler  cannot catch this.

恐慌输出的名称提供了一个提示：“恐慌：运行时错误：索引超出范围”。我们用三个海洋生物创建了一个切片。然后我们尝试通过使用 `len` 内置函数用切片的长度索引该切片来获取切片的最后一个元素。请记住，切片和数组是从零开始的；所以第一个元素为零，这个切片中的最后一个元素在索引“2”处。由于我们尝试访问位于第三个索引“3”处的切片，因此切片中没有要返回的元素，因为它超出了切片的边界。运行时别无选择，只能终止并退出，因为我们要求它做一些不可能的事情。 Go 也无法在编译期间证明这段代码会尝试这样做，因此编译器无法捕捉到这一点。

Notice also that the subsequent code did not run. This is because a  panic is an event that completely halts the execution of your Go program. The message produced contains multiple pieces of information  helpful for diagnosing the cause of the panic.

另请注意，后续代码没有运行。这是因为 panic 是一个完全停止 Go 程序执行的事件。生成的消息包含多条有助于诊断恐慌原因的信息。

## Anatomy of a Panic

## 恐慌剖析

Panics are composed of a message indicating the cause of the panic and a [stack trace](https://en.wikipedia.org/wiki/Stack_trace) that helps you locate where in your code the panic was produced.

恐慌由一条指示恐慌原因的消息和一个 [堆栈跟踪](https://en.wikipedia.org/wiki/Stack_trace) 组成，可帮助您定位在代码中产生恐慌的位置。

The first part of any panic is the message. It will always begin with the string `panic:`, which will be followed with a string that varies depending on the cause of the panic. The panic from the previous exercise has the message:

任何恐慌的第一部分是消息。它将始终以字符串 `panic:` 开头，后面跟着一个字符串，该字符串根据引起恐慌的原因而有所不同。上一个练习中的恐慌信息包含以下信息：

```
panic: runtime error: index out of range [3] with length 3
```

The string `runtime error:` following the `panic:` prefix tells us that the panic was generated by the language runtime. This panic is telling us that we attempted to use an index `[3]` that was out of range of the slice’s length `3`.

`panic:` 前缀后面的字符串 `runtime error:` 告诉我们恐慌是由语言运行时生成的。这种恐慌告诉我们，我们试图使用超出切片长度“3”范围的索引“[3]”。

Following this message is the stack trace. Stack traces form a map  that we can follow to locate exactly what line of code was executing  when the panic was generated, and how that code was invoked by earlier  code.

此消息之后是堆栈跟踪。堆栈跟踪形成了一个地图，我们可以跟踪它以准确定位生成恐慌时正在执行的代码行，以及早期代码如何调用该代码。

```
goroutine 1 [running]:
main.main()
    /tmp/sandbox879828148/prog.go:13 +0x20
```

This stack trace, from the previous example, shows that our program generated the panic from the file `/tmp/sandbox879828148/prog.go` at line number 13. It also tells us that this panic was generated in the `main()` function from the `main` package.

这个堆栈跟踪来自前面的例子，显示我们的程序从文件 `/tmp/sandbox879828148/prog.go` 的第 13 行生成了 panic。它还告诉我们这个 panic 是在 `main()` 中产生的`main` 包中的函数。

The stack trace is broken into separate blocks—one for each [goroutine](https://tour.golang.org/concurrency/1) in your program. Every Go program’s execution is accomplished by one or more goroutines that can each independently and simultaneously execute  parts of your Go code. Each block begins with the header `goroutine X [state]:`. The header gives the ID number of the goroutine along with the state  that it was in when the panic occurred. After the header, the stack  trace shows the function that the program was executing when the panic  happened, along with the filename and line number where the function  executed.

堆栈跟踪被分成单独的块 - 程序中的每个 [goroutine](https://tour.golang.org/concurrency/1) 一个。每个 Go 程序的执行都由一个或多个 goroutine 完成，每个 goroutine 可以独立并同时执行部分 Go 代码。每个块都以头部 `goroutine X [state]:` 开头。标头给出了 goroutine 的 ID 号以及发生恐慌时它所处的状态。在标头之后，堆栈跟踪显示发生恐慌时程序正在执行的函数，以及函数执行的文件名和行号。

The panic in the previous example was generated by an out-of-bounds  access to a slice. Panics can also be generated when methods are called  on pointers that are unset.

上一个示例中的恐慌是由对切片的越界访问产生的。当在未设置的指针上调用方法时，也会产生恐慌。

## Nil Receivers

## 无接收器

The Go programming language has pointers to refer to a specific  instance of some type existing in the computer’s memory at runtime. Pointers can assume the value `nil` indicating that they are not pointing at anything. When we attempt to call methods on a pointer that is `nil`, the Go runtime will generate a panic. Similarly, variables that are  interface types will also produce panics when methods are called on  them. To see the panics generated in these cases, try the following  example:

Go 编程语言具有指向运行时存在于计算机内存中的某种类型的特定实例的指针。指针可以假定值 nil 表示它们不指向任何东西。当我们尝试在一个为 nil 的指针上调用方法时，Go 运行时会产生一个恐慌。同样，接口类型的变量在调用方法时也会产生恐慌。要查看在这些情况下生成的恐慌，请尝试以下示例：

```go
package main

import (
    "fmt"
)

type Shark struct {
    Name string
}

func (s *Shark) SayHello() {
    fmt.Println("Hi! My name is", s.Name)
}

func main() {
    s := &Shark{"Sammy"}
    s = nil
    s.SayHello()
}
```

 

The panics produced will look like this:

产生的恐慌看起来像这样：

```
Outputpanic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0xffffffff addr=0x0 pc=0xdfeba]

goroutine 1 [running]:
main.(*Shark).SayHello(...)
    /tmp/sandbox160713813/prog.go:12
main.main()
    /tmp/sandbox160713813/prog.go:18 +0x1a
```

In this example, we defined a struct called `Shark`. `Shark` has one method defined on its pointer receiver called `SayHello` that will print a greeting to standard out when called. Within the body of our `main` function, we create a new instance of this `Shark` struct and request a pointer to it using the `&` operator. This pointer is assigned to the `s` variable. We then reassign the `s` variable to the value `nil` with the statement `s = nil`. Finally we attempt to call the `SayHello` method on the variable `s`. Instead of receiving a friendly message from Sammy, we receive a panic  that we have attempted to access an invalid memory address. Because the `s` variable is `nil`, when the `SayHello` function is called, it tries to access the field `Name` on the `*Shark` type. Because this is a pointer receiver, and the receiver in this case is `nil`, it panics because it can’t dereference a `nil` pointer.

在这个例子中，我们定义了一个名为“Shark”的结构体。 `Shark` 在其指针接收器上定义了一个名为 `SayHello` 的方法，该方法将在调用时向标准输出打印问候语。在我们的 `main` 函数体内，我们创建了这个 `Shark` 结构的一个新实例，并使用 `&` 操作符请求一个指向它的指针。这个指针被分配给 `s` 变量。然后，我们使用语句 `s = nil` 将 `s` 变量重新分配给值 `nil`。最后，我们尝试对变量 `s` 调用 `SayHello` 方法。我们没有收到来自 Sammy 的友好消息，而是收到我们试图访问无效内存地址的恐慌。因为 `s` 变量是 `nil`，当 `SayHello` 函数被调用时，它会尝试访问 `*Shark` 类型上的字段 `Name`。因为这是一个指针接收器，并且在这种情况下接收器是 `nil`，它会因为无法解引用 `nil` 指针而发生恐慌。

While we have set `s` to `nil` explicitly in this example, in practice this happens less obviously. When you see panics involving `nil pointer dereference`, be sure that you have properly assigned any pointer variables that you may have created.

虽然我们在这个例子中明确地将 `s` 设置为 `nil`，但在实践中这种情况不太明显。当您看到涉及“零指针取消引用”的恐慌时，请确保您已正确分配您可能创建的任何指针变量。

Panics generated from nil pointers and out-of-bounds accesses are two commonly occurring panics generated by the runtime. It is also possible to manually generate a panic using a builtin function.

由 nil 指针和越界访问产生的恐慌是运行时产生的两种常见的恐慌。也可以使用内置函数手动生成恐慌。

## Using the `panic` Builtin Function 

## 使用 `panic` 内置函数

We can also generate panics of our own using the `panic`  built-in function. It takes a single string as an argument, which is the message the panic will produce. Typically this message is less verbose  than rewriting our code to return an error. Furthermore, we can use this within our own packages to indicate to developers that they may have  made a mistake when using our package’s code. Whenever possible, best  practice is to try to return `error` values to consumers of our package.

我们还可以使用 `panic` 内置函数生成我们自己的恐慌。它接受单个字符串作为参数，这是恐慌将产生的消息。通常，此消息比重写我们的代码以返回错误更简洁。此外，我们可以在我们自己的包中使用它来向开发人员表明他们在使用我们包的代码时可能犯了错误。只要有可能，最佳实践是尝试将 `error` 值返回给我们包的消费者。

Run this code to see a panic generated from a function called from another function:

运行此代码以查看从另一个函数调用的函数生成的恐慌：

```go
package main

func main() {
    foo()
}

func foo() {
    panic("oh no!")
}
```

 

The panic output produced looks like:

产生的恐慌输出看起来像：

```
Outputpanic: oh no!

goroutine 1 [running]:
main.foo(...)
    /tmp/sandbox494710869/prog.go:8
main.main()
    /tmp/sandbox494710869/prog.go:4 +0x40
```

Here we define a function `foo` that calls the `panic` builtin with the string `"oh no!"`. This function is called by our `main` function. Notice how the output has the message `panic: oh no!` and the stack trace shows a single goroutine with two lines in the stack trace: one for the `main()` function and one for our `foo()` function.

在这里，我们定义了一个函数 `foo`，它使用字符串 `"oh no!"` 调用内置的 `panic`。这个函数由我们的 `main` 函数调用。请注意输出中的消息“panic: oh no!”，堆栈跟踪显示了一个 goroutine，堆栈跟踪中有两行：一行用于 main() 函数，另一行用于我们的 foo() 函数。

We’ve seen that panics appear to terminate our program where they are generated. This can create problems when there are open resources that  need to be properly closed. Go provides a mechanism to execute some code always, even in the presence of a panic.

我们已经看到恐慌似乎在我们的程序产生的地方终止。当存在需要正确关闭的开放资源时，这可能会产生问题。 Go 提供了一种始终执行某些代码的机制，即使存在恐慌也是如此。

## Deferred Functions

## 延迟函数

Your program may have resources that it must clean up properly, even  while a panic is being processed by the runtime. Go allows you to defer  the execution of a function call until its calling function has  completed execution. Deferred functions run even in the presence of a  panic, and are used as a safety mechanism to guard against the chaotic  nature of panics. Functions are deferred by calling them as usual, then  prefixing the entire statement with the `defer` keyword, as in `defer sayHello()`. Run this example to see how a message can be printed even though a panic was produced:

您的程序可能拥有必须正确清理的资源，即使在运行时正在处理紧急情况时也是如此。 Go 允许您推迟函数调用的执行，直到其调用函数完成执行。延迟函数即使在出现恐慌的情况下也能运行，并用作一种安全机制来防止恐慌的混乱性质。通过像往常一样调用函数，然后在整个语句前面加上 `defer` 关键字来延迟函数，就像在 `defer sayHello()` 中一样。运行此示例以查看如何在产生恐慌的情况下打印消息：

```go
package main

import "fmt"

func main() {
    defer func() {
        fmt.Println("hello from the deferred function!")
    }()

    panic("oh no!")
}
```

 

The output produced from this example will look like:

此示例产生的输出将如下所示：

```
Outputhello from the deferred function!
panic: oh no!

goroutine 1 [running]:
main.main()
    /Users/gopherguides/learn/src/github.com/gopherguides/learn//handle-panics/src/main.go:10 +0x55
```

Within the `main` function of this example, we first `defer` a call to an anonymous function that prints the message `"hello from the deferred function!"`. The `main` function then immediately produces a panic using the `panic` function. In the output from this program, we first see that the  deferred function is executed and prints its message. Following this is  the panic we generated in `main`.

在这个例子的 `main` 函数中，我们首先 `defer` 调用一个匿名函数，该函数会打印消息 `"hello from the deferred function!"`。然后 `main` 函数使用 `panic` 函数立即产生一个恐慌。在这个程序的输出中，我们首先看到 deferred 函数被执行并打印其消息。接下来是我们在 main 中产生的恐慌。

Deferred functions provide protection against the surprising nature  of panics. Within deferred functions, Go also provides us the  opportunity to stop a panic from terminating our Go program using  another built-in function.

延迟函数提供了针对恐慌的惊人性质的保护。在延迟函数中，Go 还为我们提供了使用另一个内置函数阻止恐慌终止 Go 程序的机会。

## Handling Panics

## 处理恐慌

Panics have a single recovery mechanism—the `recover`  builtin function. This function allows you to intercept a panic on its  way up through the call stack and prevent it from unexpectedly  terminating your program. It has strict rules for its use, but can be  invaluable in a production application.

Panics 有一个单一的恢复机制——`recover` 内置函数。此函数允许您在通过调用堆栈的过程中拦截恐慌并防止它意外终止您的程序。它有严格的使用规则，但在生产应用程序中可能是无价的。

Since it is part of the `builtin` package, `recover` can be called without importing any additional packages:

由于它是 `builtin` 包的一部分，`recover` 可以在不导入任何其他包的情况下调用：

```
package main

import (
    "fmt"
    "log"
)

func main() {
    divideByZero()
    fmt.Println("we survived dividing by zero!")

}

func divideByZero() {
    defer func() {
        if err := recover();err != nil {
            log.Println("panic occurred:", err)
        }
    }()
    fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
    return a / b
}
```

This example will output:

此示例将输出：

```
Output2009/11/10 23:00:00 panic occurred: runtime error: integer divide by zero
we survived dividing by zero!
```

Our `main` function in this example calls a function we define, `divideByZero`. Within this function, we `defer` a call to an anonymous function responsible for dealing with any panics that may arise while executing `divideByZero`. Within this deferred anonymous function, we call the `recover` builtin function and assign the error it returns to a variable. If `divideByZero` is panicking, this `error` value will be set, otherwise it will be `nil`. By comparing the `err` variable against `nil`, we can detect if a panic occurred, and in this case we log the panic using the `log.Println` function, as though it were any other `error`.

在这个例子中，我们的 main 函数调用了一个我们定义的函数，divideByZero。在这个函数中，我们“延迟”一个匿名函数的调用，该函数负责处理在执行“divideByZero”时可能出现的任何恐慌。在这个延迟匿名函数中，我们调用`recover` 内置函数并将它返回的错误分配给一个变量。如果 `divideByZero` 处于 panicking 状态，这个 `error` 值将被设置，否则它将是 `nil`。通过将 `err` 变量与 `nil` 进行比较，我们可以检测是否发生了 panic，在这种情况下，我们使用 `log.Println` 函数记录了 panic，就好像它是任何其他 `error` 一样。

Following this deferred anonymous function, we call another function that we defined, `divide`, and attempt to print its results using `fmt.Println`. The arguments provided will cause `divide` to perform a division by zero, which will produce a panic.

在这个延迟匿名函数之后，我们调用我们定义的另一个函数 `divide`，并尝试使用 `fmt.Println` 打印其结果。提供的参数将导致 `divide` 执行除以零，这将产生恐慌。

In the output to this example, we first see the log message from the  anonymous function that recovers the panic, followed by the message `we survived dividing by zero!`. We have indeed done this, thanks to the `recover` builtin function stopping an otherwise catastrophic panic that would terminate our Go program.

在此示例的输出中，我们首先看到来自恢复恐慌的匿名函数的日志消息，然后是消息“我们幸存了除以零！”。我们确实做到了这一点，这要归功于 `recover` 内置函数阻止了会终止我们的 Go 程序的灾难性恐慌。

The `err` value returned from `recover()` is exactly the value that was provided to the call to `panic()`. It’s therefore critical to ensure that the `err` value is only nil when a panic has not occurred.

从 `recover()` 返回的 `err` 值正是提供给调用 `panic()` 的值。因此，确保 err 值仅在未发生恐慌时为零至关重要。

## Detecting Panics with `recover`

## 使用 `recover` 检测恐慌

The `recover` function relies on the value of the error to make determinations as to whether a panic occurred or not. Since the  argument to the `panic` function is an empty interface, it can be any type. The zero value for any interface type, including the empty interface, is `nil`. Care must be taken to avoid `nil` as an argument to `panic` as demonstrated by this example:

`recover` 函数依赖于错误的值来确定是否发生了恐慌。由于`panic` 函数的参数是一个空接口，它可以是任何类型。任何接口类型的零值，包括空接口，都是 `nil`。必须注意避免将 `nil` 作为 `panic` 的参数，如下例所示：

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    divideByZero()
    fmt.Println("we survived dividing by zero!")

}

func divideByZero() {
    defer func() {
        if err := recover();err != nil {
            log.Println("panic occurred:", err)
        }
    }()
    fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
    if b == 0 {
        panic(nil)
    }
    return a / b
}
```

 

This will output:

这将输出：

```
Outputwe survived dividing by zero!
```

This example is the same as the previous example involving `recover` with some slight modifications. The `divide` function has been altered to check if its divisor, `b`, is equal to `0`. If it is, it will generate a panic using the `panic` builtin with an argument of `nil`. The output, this time, does not include the log message showing that a panic occurred even though one was created by `divide`. This silent behavior is why it is very important to ensure that the argument to the `panic` builtin function is not `nil`.

这个例子与前面的例子相同，涉及“recover”，只是做了一些细微的修改。 `divide` 函数已更改为检查其除数 `b` 是否等于 `0`。如果是，它会使用内建的 `panic` 和 `nil` 参数产生一个恐慌。这次的输出不包括显示发生恐慌的日志消息，即使它是由 `divide` 创建的。这种静默行为就是确保`panic` 内置函数的参数不是`nil` 非常重要的原因。

## Conclusion

##  结论

We have seen a number of ways that `panic`s can be created in Go and how they can be recovered from using the `recover` builtin. While you may not necessarily use `panic` yourself, proper recovery from panics is an important step of making Go applications production-ready.

我们已经看到了在 Go 中可以创建 `panic` 的多种方法，以及如何使用内置的 `recover` 来恢复它们。虽然您自己可能不一定使用 `panic`，但从恐慌中正确恢复是使 Go 应用程序为生产做好准备的重要一步。

You can also explore [our entire How To Code in Go series](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go). 

您还可以探索 [我们整个如何在 Go 中编码系列](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go)。

