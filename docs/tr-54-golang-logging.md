# Logging in Go: Choosing a System and Using it

# 登录 Go：选择一个系统并使用它

Apr 1, 2020 From: https://www.honeybadger.io/blog/golang-logging/

Go has built-in features to make it easier  for programmers to implement logging. Third parties have also built  additional tools to make logging easier. What's the difference between  them? Which should you choose? In this article Ayooluwa Isaiah describes both of these and discusses when you'd prefer one over the other.

Go 具有内置功能，可以让程序员更轻松地实现日志记录。第三方还构建了其他工具来简化日志记录。它们之间有什么区别？你应该选择哪个？在本文中，Ayoluwa Isaiah 描述了这两种方式，并讨论了您何时更喜欢其中一种。

You're relatively new to the Go language. You're probably using it to write a web app or a server, and you need to create a log  file. So, you do a quick web search and find that there are a ton of  options for logging in go. How do you know which one to pick? This  article will equip you to answer that question.

您对 Go 语言比较陌生。您可能正在使用它来编写 Web 应用程序或服务器，并且您需要创建一个日志文件。因此，您进行了快速的网络搜索，并发现有大量登录 go 的选项。你怎么知道选哪一个？这篇文章将使您有能力回答这个问题。

We will take a look at the built-in `log` package and  determine what projects it is suited for before exploring other logging  solutions that are prevalent in the Go ecosystem.

在探索 Go 生态系统中流行的其他日志记录解决方案之前，我们将查看内置的 `log` 包并确定它适用于哪些项目。

## What to log

## 记录什么

I don't need to tell you how important logging is. Logs are used by  every production web application to help developers and operations:

我不需要告诉你日志有多重要。每个生产 Web 应用程序都使用日志来帮助开发人员和操作：

- Spot bugs in the application's code
- Discover performance problems
- Do post-mortem analysis of outages and security incidents

- 发现应用程序代码中的错误
- 发现性能问题
- 对中断和安全事件进行事后分析

The data you actually log will depend on the type of application  you're building. Most of the time, you will have some variation in the  following:

您实际记录的数据将取决于您正在构建的应用程序类型。大多数情况下，您会在以下方面有所不同：

- The timestamp for when an event occurred or a log was generated
- Log levels such as debug, error, or info
- Contextual data to help understand what happened and make it possible to easily reproduce the situation

- 事件发生或日志生成的时间戳
- 日志级别，例如调试、错误或信息
- 上下文数据有助于了解发生了什么，并可以轻松重现情况

## What not to log

## 什么不记录

In general, you shouldn't log any form of sensitive business data or  personally identifiable information. This includes, but is not limited  to:

通常，您不应记录任何形式的敏感业务数据或个人身份信息。这包括但不限于：

- Names
- IP addresses
- Credit card numbers

- 名字
- IP 地址
- 信用卡号码

These restrictions can make logs less useful from an engineering  perspective, but they make your application more secure. In many cases,  regulations such as GDPR and HIPAA may forbid the logging of personal  data.

从工程角度来看，这些限制可能会使日志变得不那么有用，但它们会使您的应用程序更加安全。在许多情况下，GDPR 和 HIPAA 等法规可能会禁止记录个人数据。

## Introducing the log package

## 引入日志包

The Go standard library has a built-in `log` package that  provides most basic logging features. While it does not have log levels  (such as debug, warning, or error), it still provides everything you  need to get a basic logging strategy set up.

Go 标准库有一个内置的 `log` 包，提供最基本的日志功能。虽然它没有日志级别（例如调试、警告或错误），但它仍然提供了设置基本日志记录策略所需的一切。

Here's the most basic logging example:

这是最基本的日志记录示例：

```
package main

import "log"

func main() {
    log.Println("Hello world!")
}
```


The code above prints the text "Hello world!" to the standard error,  but it also includes the date and time, which is handy for filtering log messages by date.

上面的代码打印文本“Hello world！”到标准错误，但它还包括日期和时间，这对于按日期过滤日志消息非常方便。

```
2019/12/09 17:21:53 Hello world!
```


> By default, the `log` package prints to the standard error (`stderr`) output stream, but you can make it write to local files or any destination that supports the `io.Writer` interface. It also adds a timestamp to the log message without any additional configuration.

> 默认情况下，`log` 包打印到标准错误 (`stderr`) 输出流，但您可以将其写入本地文件或任何支持 `io.Writer` 接口的目的地。它还向日志消息添加时间戳，无需任何额外配置。

## Logging to a file

## 记录到文件

If you need to store log messages in a file, you can do so by  creating a new file or opening an existing file and setting it as the  output of the log. Here's an example:

如果需要将日志消息存储在文件中，可以通过创建新文件或打开现有文件并将其设置为日志输出来实现。下面是一个例子：

```
package main

import (
    "log"
    "os"
)

func main() {
    // If the file doesn't exist, create it or append to the file
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)

    log.Println("Hello world!")
}
```


When we run the code, the following is written to `logs.txt.`

当我们运行代码时，将以下内容写入`logs.txt.`

```
2019/12/09 17:22:47 Hello world!
```


As mentioned earlier, you can basically output your logs to any destination that implements the `io.Writer` interface, so you have a lot of flexibility when deciding where to log messages in your application.

如前所述，您基本上可以将日志输出到任何实现 io.Writer 接口的目的地，因此在决定在应用程序中记录消息的位置时，您有很大的灵活性。

## Creating custom loggers

## 创建自定义记录器

Although the `log` package implements a predefined `logger` that writes to the standard error, we can create custom logger types using the `log.New()` method.

尽管 `log` 包实现了一个预定义的 `logger` 来写入标准错误，我们可以使用 `log.New()` 方法创建自定义记录器类型。

When creating a new logger, you need to pass in three arguments to `log.New()`:

创建新记录器时，您需要将三个参数传递给 `log.New()`：

- `out`: Any type that implements the `io.Writer` interface, which is where the log data will be written to
- `prefix`: A string that is appended to the beginning of each log line
- `flag`: A set of constants that allow us to define which  logging properties to include in each log entry generated by the logger  (more on this in the next section) 

- `out`：任何实现 `io.Writer` 接口的类型，这是日志数据将被写入的地方
- `prefix`：附加到每个日志行开头的字符串
- `flag`：一组常量，允许我们定义在记录器生成的每个日志条目中包含哪些日志记录属性（下一节将详细介绍）

We can take advantage of this feature to create custom loggers. Here's an example that implements `Info`, `Warning` and `Error` loggers:

我们可以利用这个特性来创建自定义记录器。这是一个实现 `Info`、`Warning` 和 `Error` 记录器的示例：

```
package main

import (
    "log"
    "os"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
    InfoLogger.Println("Starting the application...")
    InfoLogger.Println("Something noteworthy happened")
    WarningLogger.Println("There is something you should know about")
    ErrorLogger.Println("Something went wrong")
}
```


After creating or opening the `logs.txt` file at the top of the `init` function, we then initialize the three defined loggers by providing the output destination, prefix string, and log flags.

在 `init` 函数顶部创建或打开 `logs.txt` 文件后，我们然后通过提供输出目的地、前缀字符串和日志标志来初始化三个定义的记录器。

In the `main` function, the loggers are utilized by calling the `Println` function, which writes a new log entry to the log file. When you run this program, the following will be written to `logs.txt`.

在 main 函数中，记录器通过调用 Println 函数来使用，该函数将新的日志条目写入日志文件。当你运行这个程序时，以下内容将被写入`logs.txt`。

```
INFO: 2019/12/09 12:01:06 main.go:26: Starting the application...
INFO: 2019/12/09 12:01:06 main.go:27: Something noteworthy happened
WARNING: 2019/12/09 12:01:06 main.go:28: There is something you should know about
ERROR: 2019/12/09 12:01:06 main.go:29: Something went wrong
```


Note that in this example, we are logging to a single file, but you  can use a separate file for each logger by passing a different file when creating the logger.

请注意，在此示例中，我们将记录到单个文件，但您可以通过在创建记录器时传递不同的文件来为每个记录器使用单独的文件。

## Log flags

## 日志标志

You can use [log flag constants](https://golang.org/pkg/log/#pkg-constants) to enrich a log message by providing additional context information,  such as the file, line number, date, and time. For example, passing the  message "Something went wrong" through a logger with a flag combination  shown below:

您可以使用 [日志标志常量](https://golang.org/pkg/log/#pkg-constants) 通过提供额外的上下文信息（例如文件、行号、日期和时间）来丰富日志消息。例如，通过带有如下标志组合的记录器传递消息“出现问题”：

```
log.Ldate|log.Ltime|log.Lshortfile
```


will print

将打印

```
2019/12/09 12:01:06 main.go:29: Something went wrong
```


Unfortunately, there is no control over the order in which they appear or the format in which they are presented.

不幸的是，无法控制它们出现的顺序或呈现的格式。

## Introducing logging frameworks

## 介绍日志框架

Using the `log` package is great for local development  when getting fast feedback is more important than generating rich,  structured logs. Beyond that, you will mostly likely be better off using a logging framework.

当获得快速反馈比生成丰富的结构化日志更重要时，使用 `log` 包非常适合本地开发。除此之外，您很可能最好使用日志记录框架。

A major advantage of using a logging framework is that it helps to standardize the log data. This means that:

使用日志框架的一个主要优点是它有助于标准化日志数据。这意味着：

- It's easier to read and understand the log data.
- It's easier to gather logs from several sources and feed them to a central platform to be analyzed.

- 更容易阅读和理解日志数据。
- 更容易从多个来源收集日志并将它们提供给中央平台进行分析。

In addition, logging is pretty much a solved problem. Why reinvent the wheel?

此外，日志记录几乎是一个已解决的问题。为什么要重新发明轮子？

## Choosing a logging framework

## 选择日志框架

Deciding which framework to use can be a challenge, as there are [several options](https://github.com/avelino/awesome-go#logging) to choose from.

决定使用哪个框架可能是一个挑战，因为有 [几个选项](https://github.com/avelino/awesome-go#logging) 可供选择。

The two most popular logging frameworks for Go appear to be [glog](https://github.com/golang/glog)and [logrus](https://github.com/Sirupsen/logrus). The popularity of glog is surprising, since it hasn't been updated in  several years. logrus is better maintained and used in popular projects  like Docker, so we'll be focusing on it.

Go 的两个最流行的日志框架似乎是 [glog](https://github.com/golang/glog)和 [logrus](https://github.com/Sirupsen/logrus)。 glog 的流行令人惊讶，因为它已经好几年没有更新了。 logrus 在 Docker 等流行项目中得到更好的维护和使用，因此我们将重点关注它。

## Getting started with logrus

## 开始使用 logrus

Installing logrus is as simple as running the command below in your terminal:

安装 logrus 就像在终端中运行以下命令一样简单：

```
go get "github.com/Sirupsen/logrus"
```


One great thing about logrus is that it's completely compatible with the `log` package of the standard library, so you can replace your log imports everywhere with `log "github.com/sirupsen/logrus"` and it will just work!

logrus 的一大优点是它与标准库的 `log` 包完全兼容，所以你可以用 `log "github.com/sirupsen/logrus"` 替换你的日志导入，它会正常工作！

Let's modify our earlier "hello world" example that used the log package and use logrus instead:

让我们修改我们之前使用 log 包的“hello world”示例并改用 logrus：

```
package main

import (
  log "github.com/sirupsen/logrus"
)

func main() {
    log.Println("Hello world!")
}
```


Running this code produces the output:

运行此代码会产生输出：

```
INFO[0000] Hello world!
```


It couldn't be any easier!

再简单不过了！

### Logging in JSON 

### 登录 JSON

`logrus` is well suited for structured logging in JSON  which — as JSON is a well-defined standard — makes it easy for external  services to parse your logs and also makes the addition of context to a  log message relatively straightforward through the use of fields , as shown below:

`logrus` 非常适合 JSON 中的结构化日志记录，因为 JSON 是一个定义明确的标准 - 使外部服务可以轻松解析您的日志，并且还可以通过使用字段相对简单地向日志消息添加上下文， 如下所示：

```
package main

import (
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetFormatter(&log.JSONFormatter{})
    log.WithFields(
        log.Fields{
            "foo": "foo",
            "bar": "bar",
        },
    ).Info("Something happened")
}
```


The log output generated will be a JSON object that includes the message, log level, timestamp, and included fields.

生成的日志输出将是一个 JSON 对象，其中包括消息、日志级别、时间戳和包含的字段。

```
{"bar":"bar","foo":"foo","level":"info","msg":"Something happened","time":"2019-12-09T15:55:24+01:00"}
```


If you're not interested in outputting your logs as JSON, be aware  that several third-party formatters exist for logrus, which you can view on its [Github page](https://github.com/Sirupsen/logrus#formatters) . You can even write your own formatter if you prefer.

如果您对将日志输出为 JSON 不感兴趣，请注意 logrus 存在多个第三方格式化程序，您可以在其 [Github 页面](https://github.com/Sirupsen/logrus#formatters) 上查看.如果您愿意，您甚至可以编写自己的格式化程序。

### Log levels

### 日志级别

Unlike the standard log package, logrus supports log levels.

与标准日志包不同，logrus 支持日志级别。

logrus has seven log levels: Trace, Debug, Info, Warn, Error, Fatal,  and Panic. The severity of each level increases as you go down the list.

logrus 有七个日志级别：Trace、Debug、Info、Warn、Error、Fatal 和 Panic。每个级别的严重性随着您在列表中的向下而增加。

```
log.Trace("Something very low level.")
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
// Calls os.Exit(1) after logging
log.Fatal("Bye.")
// Calls panic() after logging
log.Panic("I'm bailing.")
```


By setting a logging level on a logger, you can log only the entries  you need depending on your environment. By default, logrus will log  anything that is Info or above (Warn, Error, Fatal, or Panic).

通过在记录器上设置日志记录级别，您可以根据您的环境仅记录您需要的条目。默认情况下，logrus 将记录信息或更高级别的任何内容（警告、错误、致命或恐慌）。

```
package main

import (
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetFormatter(&log.JSONFormatter{})

    log.Debug("Useful debugging information.")
    log.Info("Something noteworthy happened!")
    log.Warn("You should probably take a look at this.")
    log.Error("Something failed but I'm not quitting.")
}
```


Running the above code will produce the following output:

运行上面的代码将产生以下输出：

```
{"level":"info","msg":"Something noteworthy happened!","time":"2019-12-09T16:18:21+01:00"}
{"level":"warning","msg":"You should probably take a look at this.","time":"2019-12-09T16:18:21+01:00"}
{"level":"error","msg":"Something failed but I'm not quitting.","time":"2019-12-09T16:18:21+01:00"}
```


Notice that the Debug level message was not printed. To include it in the logs, set `log.Level` to equal `log.DebugLevel`:

请注意，未打印调试级别消息。要将其包含在日志中，请将 `log.Level` 设置为等于 `log.DebugLevel`：

```
log.SetLevel(log.DebugLevel)
```


## Wrap up

##  包起来

In this post, we explored the use of the built-in log package and  established that it should only be used for trivial applications or when building a quick prototype. For everything else, the use of a  mainstream logging framework is a must.

在这篇文章中，我们探索了内置日志包的使用，并确定它应该只用于琐碎的应用程序或构建快速原型时。对于其他一切，必须使用主流日志框架。

We also looked at ways to ensure that the information contained in your logs is consistent and easy to analyze, especially when aggregating it on a centralized platform.

我们还研究了确保日志中包含的信息一致且易于分析的方法，尤其是在集中式平台上聚合时。

Thanks for reading! 

谢谢阅读！

