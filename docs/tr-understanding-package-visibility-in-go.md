# Understanding Package Visibility in Go

   # 了解 Go 中的包可见性

September 24, 2019 43.6k views

### Introduction

###  介绍

When creating a [package in Go](https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go), the end goal is usually to make the package accessible for other  developers to use, either in higher order packages or whole programs. By [importing the package](https://www.digitalocean.com/community/tutorials/importing-packages-in-go), your piece of code can serve as the building block for other, more  complex tools. However, only certain packages are available for  importing. This is determined by the visibility of the package.

在创建一个 [Go 中的包](https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go)时，最终目标通常是让其他开发人员可以访问该包使用，无论是在更高阶的包中还是在整个程序中。通过[导入包](https://www.digitalocean.com/community/tutorials/importing-packages-in-go)，您的代码段可以作为其他更复杂工具的构建块。但是，只有某些包可用于导入。这是由包的可见性决定的。

*Visibility* in this context means the file space from which a package or other construct can be referenced. For example, if we define a variable in a function, the visibility (scope) of that variable is  only within the function in which it was defined. Similarly, if you  define a variable in a package, you can make it visible to just that  package, or allow it to be visible outside the package as well.

*可见性*在此上下文中表示可以从中引用包或其他构造的文件空间。例如，如果我们在函数中定义一个变量，该变量的可见性（范围）仅在定义它的函数内。类似地，如果您在包中定义一个变量，则可以使其仅对该包可见，也可以允许它在包外可见。

Carefully controlling package visibility is important when writing  ergonomic code, especially when accounting for future changes that you  may want to make to your package. If you need to fix a bug, improve  performance, or change functionality, you’ll want to make the change in a way that won’t break the code of anyone using your package. One way to  minimize breaking changes is to allow access only to the parts of your  package that are needed for it to be used properly. By limiting access,  you can make changes internally to your package with less of a chance of affecting how other developers are using your package.

在编写符合人体工程学的代码时，仔细控制包的可见性很重要，尤其是在考虑您可能希望对包进行的未来更改时。如果您需要修复错误、提高性能或更改功能，您需要以一种不会破坏使用您的包的任何人的代码的方式进行更改。最小化破坏性更改的一种方法是仅允许访问包中正确使用所需的部分。通过限制访问，您可以在内部对您的包进行更改，而不太可能影响其他开发人员如何使用您的包。

In this article, you will learn how to control package visibility, as well as how to protect parts of your code that should only be used  inside your package. To do this, we will create a basic logger to log  and debug messages, using packages with varying degrees of item  visibility.

在本文中，您将学习如何控制包可见性，以及如何保护应仅在包内使用的代码部分。为此，我们将创建一个基本的记录器来记录和调试消息，使用具有不同程度项目可见性的包。

## Prerequisites

## 先决条件

To follow the examples in this article, you will need:

要遵循本文中的示例，您需要：

- A Go workspace set up by following [How To Install Go and Set Up a Local Programming Environment](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go). This tutorial will use the following file structure:

- 按照[如何安装 Go 并设置本地编程环境](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-) 设置的 Go 工作区本地编程环境可使用)。本教程将使用以下文件结构：

```
.
├── bin
│
└── src
    └── github.com
        └── gopherguides
```

## Exported and Unexported Items

## 导出和未导出的项目

Unlike other program languages like Java and [Python](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-python-3) that use *access modifiers* such as `public`, ` private`, or `protected` to specify scope, Go determines if an item is `exported` and `unexported` through how it is declared. Exporting an item in this case makes it `visible` outside the current package. If it’s not exported, it is only visible and usable from within the package it was defined.

与 Java 和 [Python](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-python-3) 等其他程序语言不同，它们使用 *访问修饰符*，例如`public`、` private` 或 `protected` 来指定作用域，Go 通过它的声明方式来确定一个项目是“exported”还是“unexported”。在这种情况下导出项目使其在当前包之外“可见”。如果它没有被导出，它只能在它定义的包中可见和可用。

This external visibility is controlled by capitalizing the first letter of the item declared. All declarations, such as `Types`, `Variables`, `Constants`, `Functions`, etc., that start with a capital letter are visible outside the current package.

这种外部可见性是通过将声明的项目的第一个字母大写来控制的。所有以大写字母开头的声明，例如 `Types`、`Variables`、`Constants`、`Functions` 等，在当前包之外都是可见的。

Let’s look at the following code, paying careful attention to capitalization:

让我们看下面的代码，注意大小写：

greet.go

打招呼

```bash
package greet

import "fmt"

var Greeting string

func Hello(name string) string {
    return fmt.Sprintf(Greeting, name)
}
```

 

This code declares that it is in the `greet` package. It then declares two symbols, a variable called `Greeting`, and a function called `Hello`. Because they both start with a capital letter, they are both `exported` and available to any outside program. As stated earlier, crafting a  package that limits access will allow for better API design and make it  easier to update your package internally without breaking anyone’s code  that is depending on your package.

这段代码声明它在 `greet` 包中。然后它声明了两个符号，一个名为“Greeting”的变量和一个名为“Hello”的函数。因为它们都以大写字母开头，所以它们都可以“导出”并可用于任何外部程序。如前所述，制作一个限制访问的包将允许更好的 API 设计，并在不破坏依赖于你的包的任何人的代码的情况下更容易地在内部更新你的包。

## Defining Package Visibility 

## 定义包可见性

To give a closer look at how package visibility works in a program, let’s create a `logging` package, keeping in mind what we want to make visible outside our  package and what we won’t make visible. This logging package will be  responsible for logging any of our program messages to the console. It  will also look at what *level* we are logging at. A level describes the type of log, and is going to be one of three statuses: `info`, `warning`, or `error`.

为了更深入地了解包可见性在程序中的工作原理，让我们创建一个 `logging` 包，记住我们希望在包外可见的内容以及不可见的内容。这个日志包将负责将我们的任何程序消息记录到控制台。它还将查看我们正在登录的*级别*。级别描述了日志的类型，它将是以下三种状态之一：`info`、`warning` 或 `error`。

First, within your `src` directory, let’s create a directory called `logging` to put our logging files in:

首先，在你的 `src` 目录中，让我们创建一个名为 `logging` 的目录来放置我们的日志文件：

```bash
mkdir logging
```

 

Move into that directory next:

接下来进入该目录：

```bash
cd logging
```

 

Then, using an editor like nano, create a file called `logging.go`:

然后，使用类似 nano 的编辑器，创建一个名为“logging.go”的文件：

```bash
nano logging.go
```

 

Place the following code in the `logging.go` file we just created:

将以下代码放入我们刚刚创建的 `logging.go` 文件中：

logging/logging.go

```go
package logging

import (
    "fmt"
    "time"
)

var debug bool

func Debug(b bool) {
    debug = b
}

func Log(statement string) {
    if !debug {
        return
    }

    fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), statement)
}
```

The first line of this code declared a package called `logging`. In this package, there are two `exported` functions: `Debug` and `Log`. These functions can be called by any other package that imports the `logging` package. There is also a private variable called `debug`. This variable is only accessible from within the `logging` package. It is important to note that while the function `Debug` and the variable `debug` both have the same spelling, the function is capitalized and the  variable is not. This makes them distinct declarations with different  scopes.

这段代码的第一行声明了一个名为“logging”的包。在这个包中，有两个`exported` 函数：`Debug` 和`Log`。这些函数可以被任何其他导入 `logging` 包的包调用。还有一个名为“debug”的私有变量。这个变量只能从 `logging` 包中访问。需要注意的是，虽然函数 `Debug` 和变量 `debug` 的拼写相同，但函数是大写的，而变量不是。这使它们具有不同范围的不同声明。

Save and quit the file.

保存并退出文件。

To use this package in other areas of our code, we can [`import` it into a new package](https://www.digitalocean.com/community/tutorials/importing-packages-in-go). We’ll create this new package, but we’ll need a new directory to store those source files in first.

要在我们代码的其他区域使用这个包，我们可以 [`import` 到一个新的包中](https://www.digitalocean.com/community/tutorials/importing-packages-in-go)。我们将创建这个新包，但我们首先需要一个新目录来存储这些源文件。

Let’s move out of the `logging` directory, create a new directory called `cmd`, and move into that new directory:

让我们移出 `logging` 目录，创建一个名为 `cmd` 的新目录，然后移入该新目录：

```bash
cd ..
mkdir cmd
cd cmd
```

 

Create a file called `main.go` in the `cmd` directory we just created:

在我们刚刚创建的 `cmd` 目录中创建一个名为 `main.go` 的文件：

```bash
nano main.go
```

 

Now we can add the following code:

现在我们可以添加以下代码：

cmd/main.go

```go
package main

import "github.com/gopherguides/logging"

func main() {
    logging.Debug(true)

    logging.Log("This is a debug statement...")
}
```

 

We now have our entire program written. However, before we can run  this program, we’ll need to also create a couple of configuration files  for our code to work properly. Go uses [Go Modules](https://blog.golang.org/using-go-modules) to configure package dependencies for importing resources. Go modules  are configuration files placed in your package directory that tell the  compiler where to import packages from. While learning about modules is  beyond the scope of this article, we can write just a couple lines of  configuration to make this example work locally.

我们现在已经编写了整个程序。然而，在我们运行这个程序之前，我们还需要创建几个配置文件来让我们的代码正常工作。 Go 使用 [Go Modules](https://blog.golang.org/using-go-modules) 来配置用于导入资源的包依赖项。 Go 模块是放置在包目录中的配置文件，它告诉编译器从何处导入包。虽然学习模块超出了本文的范围，但我们只需编写几行配置即可使此示例在本地工作。

Open the following `go.mod` file in the `cmd` directory:

在 `cmd` 目录中打开以下 `go.mod` 文件：

```bash
nano go.mod
```

 

Then place the following contents in the file:

然后将以下内容放入文件中：

go.mod

```bash
module github.com/gopherguides/cmd

replace github.com/gopherguides/logging => ../logging
```

 

The first line of this file tells the compiler that the `cmd` package has a file path of `github.com/gopherguides/cmd`. The second line tells the compiler that the package `github.com/gopherguides/logging` can be found locally on disk in the `../logging` directory.

该文件的第一行告诉编译器 `cmd` 包的文件路径为 `github.com/gopherguides/cmd`。第二行告诉编译器可以在本地磁盘上的“../logging”目录中找到包“github.com/gopherguides/logging”。

We’ll also need a `go.mod` file for our `logging` package. Let’s move back into the `logging` directory and create a `go.mod` file:

我们的 `logging` 包还需要一个 `go.mod` 文件。让我们回到`logging`目录并创建一个`go.mod`文件：

```bash
cd ../logging
nano go.mod
```

 

Add the following contents to the file:

将以下内容添加到文件中：

go.mod

```bash
module github.com/gopherguides/logging
```

 

This tells the compiler that the `logging` package we created is actually the `github.com/gopherguides/logging` package. This makes it possible to import the package in our `main` package with the following line that we wrote earlier:

这告诉编译器我们创建的 `logging` 包实际上是 `github.com/gopherguides/logging` 包。这使得可以使用我们之前编写的以下行将包导入到我们的 `main` 包中：

cmd/main.go

```bash
package main

import "github.com/gopherguides/logging"

func main() {
    logging.Debug(true)

    logging.Log("This is a debug statement...")
}
```

 

You should now have the following directory structure and file layout:

您现在应该具有以下目录结构和文件布局：

```
├── cmd
│   ├── go.mod
│   └── main.go
└── logging
    ├── go.mod
    └── logging.go
```

Now that we have all the configuration completed, we can run the `main` program from the `cmd` package with the following commands:

现在我们已经完成了所有的配置，我们可以使用以下命令从 `cmd` 包中运行 `main` 程序：

```bash
cd ../cmd
go run main.go
```



You will get output similar to the following:

您将获得类似于以下内容的输出：

```
Output2019-08-28T11:36:09-05:00 This is a debug statement...
```

The program will print out the current time in RFC 3339 format followed by whatever statement we sent to the logger. [RFC 3339](https://tools.ietf.org/html/rfc3339) is a time format that was designed to represent time on the internet and is commonly used in log files.

该程序将以 RFC 3339 格式打印当前时间，然后是我们发送给记录器的任何语句。 [RFC 3339](https://tools.ietf.org/html/rfc3339) 是一种时间格式，旨在表示 Internet 上的时间，通常用于日志文件。

Because the `Debug` and `Log` functions are exported from the logging package, we can use them in our `main` package. However, the `debug` variable in the `logging` package is not exported. Trying to reference an unexported declaration will result in a compile-time error.

因为 `Debug` 和 `Log` 函数是从 logging 包中导出的，我们可以在 `main` 包中使用它们。然而，`logging` 包中的 `debug` 变量没有被导出。尝试引用未导出的声明将导致编译时错误。

Add the following highlighted line to `main.go`:

将以下突出显示的行添加到`main.go`：

cmd/main.go

```bash
package main

import "github.com/gopherguides/logging"

func main() {
    logging.Debug(true)

    logging.Log("This is a debug statement...")

    fmt.Println(logging.debug)
}
```

 

Save and run the file. You will receive an error similar to the following:

保存并运行文件。您将收到类似于以下内容的错误：

```
Output...
./main.go:10:14: cannot refer to unexported name logging.debug
```

Now that we have seen how `exported` and `unexported` items in packages behave, we will next look at how `fields` and `methods` can be exported from `structs`.

现在我们已经了解了包中 `exported` 和 `unexported` 项的行为，接下来我们将看看如何从 `structs` 中导出 `fields` 和 `methods`。

## Visibility Within Structs

## 结构内的可见性

While the visibility scheme in the logger we built in the last  section may work for simple programs, it shares too much state to be  useful from within multiple packages. This is because the exported  variables are accessible to multiple packages that could modify the  variables into contradictory states. Allowing the state of your package  to be changed in this way makes it hard to predict how your program will behave. With the current design, for example, one package could set the `Debug` variable to `true`, and another could set it to `false` in the same instance. This would create a problem since both packages that are importing the `logging` package are affected.

虽然我们在上一节中构建的记录器中的可见性方案可能适用于简单的程序，但它共享太多的状态而无法在多个包中使用。这是因为导出的变量可以被多个包访问，这些包可以将变量修改为相互矛盾的状态。允许以这种方式更改包的状态会使预测程序的行为变得困难。例如，在当前设计中，一个包可以将 `Debug` 变量设置为 `true`，而另一个包可以在同一实例中将其设置为 `false`。这会产生问题，因为导入 `logging` 包的两个包都会受到影响。

We can make the logger isolated by creating a struct and then hanging methods off of it. This will allow us to create an `instance` of a logger to be used independently in each package that consumes it.

我们可以通过创建一个结构然后将方法挂在它上面来隔离记录器。这将允许我们创建一个记录器的“实例”，以便在使用它的每个包中独立使用。

Change the `logging` package to the following to refactor the code and isolate the logger:

将 `logging` 包更改为以下内容以重构代码并隔离记录器：

logging/logging.go

```go
package logging

import (
    "fmt"
    "time"
)

type Logger struct {
    timeFormat string
    debug      bool
}

func New(timeFormat string, debug bool) *Logger {
    return &Logger{
        timeFormat: timeFormat,
        debug:      debug,
    }
}

func (l *Logger) Log(s string) {
    if !l.debug {
        return
    }
    fmt.Printf("%s %s\n", time.Now().Format(l.timeFormat), s)
}
```

In this code, we created a `Logger` struct. This struct will house our unexported state, including the time format to print out and the `debug` variable setting of `true` or `false`. The `New` function sets the initial state to create the logger with, such as the  time format and debug state. It then stores the values we gave it  internally to the unexported variables `timeFormat` and `debug`. We also created a method called `Log` on the `Logger` type that takes a statement we want to print out. Within the `Log` method is a reference to its local method variable `l` to get access back to its internal fields such as `l.timeFormat` and `l.debug`.

在这段代码中，我们创建了一个 `Logger` 结构。这个结构体将容纳我们未导出的状态，包括要打印的时间格式和 `true` 或 `false` 的 `debug` 变量设置。 `New` 函数设置用于创建记录器的初始状态，例如时间格式和调试状态。然后它将我们在内部给它的值存储到未导出的变量 `timeFormat` 和 `debug`。我们还在 `Logger` 类型上创建了一个名为 `Log` 的方法，它接受我们想要打印的语句。在`Log` 方法中是对其局部方法变量`l` 的引用，以访问其内部字段，例如`l.timeFormat` 和`l.debug`。

This approach will allow us to create a `Logger` in many different packages and use it independently of how the other packages are using it.

这种方法将允许我们在许多不同的包中创建一个 `Logger`，并独立于其他包如何使用它来使用它。

To use it in another package, let’s alter `cmd/main.go` to look like the following:

要在另一个包中使用它，让我们将 `cmd/main.go` 更改为如下所示：

cmd/main.go

```go
package main

import (
    "time"

    "github.com/gopherguides/logging"
)

func main() {
    logger := logging.New(time.RFC3339, true)

    logger.Log("This is a debug statement...")
}
```

 

Running this program will give you the following output:

运行此程序将为您提供以下输出：

```
Output2019-08-28T11:56:49-05:00 This is a debug statement...
```

In this code, we created an instance of the logger by calling the exported function `New`. We stored the reference to this instance in the `logger` variable. We can now call `logging.Log` to print out statements. 

在这段代码中，我们通过调用导出函数“New”创建了一个记录器实例。我们将对该实例的引用存储在 `logger` 变量中。我们现在可以调用“logging.Log”来打印语句。

If we try to reference an unexported field from the `Logger` such as the `timeFormat` field, we will receive a compile-time error. Try adding the following highlighted line and running `cmd/main.go`:

如果我们尝试从 `Logger` 中引用未导出的字段，例如 `timeFormat` 字段，我们将收到编译时错误。尝试添加以下突出显示的行并运行`cmd/main.go`：

cmd/main.go

```bash
package main

import (
    "time"

    "github.com/gopherguides/logging"
)

func main() {
    logger := logging.New(time.RFC3339, true)

    logger.Log("This is a debug statement...")

    fmt.Println(logger.timeFormat)
}
```

 

This will give the following error:

这将给出以下错误：

```
Output...
cmd/main.go:14:20: logger.timeFormat undefined (cannot refer to unexported field or method timeFormat)
```

The compiler recognizes that `logger.timeFormat` is not exported, and therefore can’t be retrieved from the `logging` package.

编译器识别出 `logger.timeFormat` 未导出，因此无法从 `logging` 包中检索。

## Visibility Within Methods

## 方法中的可见性

In the same way as struct fields, methods can also be exported or unexported.

与 struct 字段一样，方法也可以导出或不导出。

To illustrate this, let’s add *leveled* logging to our logger. Leveled logging is a means of categorizing your logs so that you can  search your logs for specific types of events. The levels we will put  into our logger are:

为了说明这一点，让我们将 *leveled* 日志记录添加到我们的记录器中。分级日志记录是一种对日志进行分类的方法，以便您可以在日志中搜索特定类型的事件。我们将放入记录器的级别是：

- The `info` level, which represents information type events that inform the user of an action, such as `Program started`, or `Email sent`. These help us debug and track parts of our program to see if expected behavior is happening.
- The `warning` level. These types of events identify when something unexpected is happening that is not an error, like `Email failed to send, retrying`. They help us see parts of our program that aren’t going as smoothly as we expected them to.
- The `error` level, which means the program encountered a problem, like `File not found`. This will often result in the program’s operation failing.

- `info` 级别，表示通知用户操作的信息类型事件，例如“程序已启动”或“电子邮件已发送”。这些帮助我们调试和跟踪程序的某些部分，以查看是否发生了预期的行为。
- “警告”级别。这些类型的事件识别何时发生了非错误的意外事件，例如“电子邮件发送失败，正在重试”。他们帮助我们看到我们计划中没有像我们预期的那样顺利的部分。
- `error` 级别，表示程序遇到问题，例如 `File not found`。这往往会导致程序运行失败。

You may also desire to turn on and off certain levels of logging,  especially if your program isn’t performing as expected and you’d like  to debug the program. We’ll add this functionality by changing the  program so that when `debug` is set to `true`, it will print all levels of messages. Otherwise, if it’s `false`, it will only print error messages.

您可能还希望打开和关闭某些级别的日志记录，尤其是当您的程序未按预期执行并且您想调试程序时。我们将通过更改程序来添加此功能，以便当 `debug` 设置为 `true` 时，它将打印所有级别的消息。否则，如果它是 `false`，它只会打印错误消息。

Add leveled logging by making the following changes to `logging/logging.go`:

通过对 `logging/logging.go` 进行以下更改来添加级别日志记录：

logging/logging.go

```go
package logging

import (
    "fmt"
    "strings"
    "time"
)

type Logger struct {
    timeFormat string
    debug      bool
}

func New(timeFormat string, debug bool) *Logger {
    return &Logger{
        timeFormat: timeFormat,
        debug:      debug,
    }
}

func (l *Logger) Log(level string, s string) {
    level = strings.ToLower(level)
    switch level {
    case "info", "warning":
        if l.debug {
            l.write(level, s)
        }
    default:
        l.write(level, s)
    }
}

func (l *Logger) write(level string, s string) {
    fmt.Printf("[%s] %s %s\n", level, time.Now().Format(l.timeFormat), s)
}
```

 

In this example, we introduced a new argument to the `Log` method. We can now pass in the `level` of the log message. The `Log` method determines what level of message it is. If it’s an `info` or `warning` message, and the `debug` field is `true`, then it writes the message. Otherwise it ignores the message. If it is any other level, like `error`, it will write out the message regardless.

在这个例子中，我们向 `Log` 方法引入了一个新参数。我们现在可以传入日志消息的“级别”。 `Log` 方法确定它是什么级别的消息。如果它是一个 `info` 或 `warning` 消息，并且 `debug` 字段为 `true`，那么它会写入该消息。否则它会忽略该消息。如果它是任何其他级别，例如 `error`，它将无论如何写出消息。

Most of the logic for determining if the message is printed out exists in the `Log` method. We also introduced an unexported method called `write`. The `write` method is what actually outputs the log message.

大多数用于确定消息是否打印出来的逻辑都存在于 `Log` 方法中。我们还引入了一种称为“write”的未导出方法。 `write` 方法是实际输出日志消息的方法。

We can now use this leveled logging in our other package by changing `cmd/main.go` to look like the following:

我们现在可以通过将 `cmd/main.go` 更改为如下所示，在我们的其他包中使用这种级别的日志记录：

cmd/main.go

```go
package main

import (
    "time"

    "github.com/gopherguides/logging"
)

func main() {
    logger := logging.New(time.RFC3339, true)

    logger.Log("info", "starting up service")
    logger.Log("warning", "no tasks found")
    logger.Log("error", "exiting: no work performed")

}
```

 

Running this will give you:

运行它会给你：

```
Output[info] 2019-09-23T20:53:38Z starting up service
[warning] 2019-09-23T20:53:38Z no tasks found
[error] 2019-09-23T20:53:38Z exiting: no work performed
```

In this example, `cmd/main.go` successfully used the exported `Log` method.

在这个例子中，`cmd/main.go` 成功地使用了导出的 `Log` 方法。

We can now pass in the `level` of each message by switching `debug` to `false`:

我们现在可以通过将 `debug` 切换为 `false` 来传递每条消息的 `level`：

main.go

```go
package main

import (
    "time"

    "github.com/gopherguides/logging"
)

func main() {
    logger := logging.New(time.RFC3339, false)

    logger.Log("info", "starting up service")
    logger.Log("warning", "no tasks found")
    logger.Log("error", "exiting: no work performed")

}
```

 

Now we will see that only the `error` level messages print:

现在我们将看到只有 `error` 级别的消息打印：

```
Output[error] 2019-08-28T13:58:52-05:00 exiting: no work performed
```

If we try to call the `write` method from outside the `logging` package, we will receive a compile-time error:

如果我们尝试从 `logging` 包外部调用 `write` 方法，我们将收到编译时错误：

main.go

```bash
package main

import (
    "time"

    "github.com/gopherguides/logging"
)

func main() {
    logger := logging.New(time.RFC3339, true)

    logger.Log("info", "starting up service")
    logger.Log("warning", "no tasks found")
    logger.Log("error", "exiting: no work performed")

    logger.write("error", "log this message...")
}
```

 

```
Outputcmd/main.go:16:8: logger.write undefined (cannot refer to unexported field or method logging.(*Logger).write)
```

When the compiler sees that you are trying to reference something  from another package that starts with a lowercase letter, it knows that  it is not exported, and therefore throws a compiler error.

当编译器发现您试图从另一个以小写字母开头的包中引用某些内容时，它知道它没有被导出，因此会引发编译器错误。

The logger in this tutorial illustrates how we can write code that  only exposes the parts we want other packages to consume. Because we  control what parts of the package are visible outside the package, we  are now able to make future changes without affecting any code that  depends on our package. For example, if we wanted to only turn off `info` level messages when `debug` is false, you could make this change without affecting any other part  of your API. We could also safely make changes to the log message to  include more information, such as the directory the program was running  from.

本教程中的记录器说明了我们如何编写仅公开我们希望其他包使用的部分的代码。因为我们控制包的哪些部分在包外可见，所以我们现在能够在不影响依赖于我们包的任何代码的情况下进行未来的更改。例如，如果我们只想在 `debug` 为 false 时关闭 `info` 级别的消息，您可以进行此更改而不影响 API 的任何其他部分。我们还可以安全地更改日志消息以包含更多信息，例如程序运行所在的目录。

## Conclusion

##  结论

This article showed how to share code between packages while also  protecting the implementation details of your package. This allows you  to export a simple API that will seldom change for backwards  compatibility, but will allow for changes privately in your package as  needed to make it work better in the future. This is considered a best  practice when creating packages and their corresponding APIs.

本文展示了如何在包之间共享代码，同时保护包的实现细节。这允许您导出一个简单的 API，该 API 很少因向后兼容性而更改，但允许根据需要在您的包中进行私人更改，以使其在未来更好地工作。在创建包及其相应的 API 时，这被认为是最佳实践。

To learn more about packages in Go, check out our [Importing Packages in Go](https://www.digitalocean.com/community/tutorials/importing-packages-in-go) and [How To Write Packages in Go]( https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go) articles, or explore our entire [How To Code in Go series](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go). 

要了解有关 Go 中包的更多信息，请查看我们的 [Importing Packages in Go](https://www.digitalocean.com/community/tutorials/importing-packages-in-go) 和 [How To Write Packages in Go]( https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go) 文章，或探索我们整个 [如何在 Go 系列中编码](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go)。

