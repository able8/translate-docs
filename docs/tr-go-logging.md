# How to collect, standardize, and centralize Golang logs

# 如何收集、标准化和集中 Golang 日志

Published: March 18, 2019

发布时间：2019 年 3 月 18 日

Organizations that depend on distributed systems often write their applications in Go to take advantage of concurrency features like channels and goroutines  (eg, [Heroku](https://blog.golang.org/go-at-heroku), [Basecamp]( https://signalvnoise.com/posts/3897-go-at-basecamp), [Cockroach Labs](https://www.cockroachlabs.com/blog/why-go-was-the-right-choice-for-cockroachdb/), and [Datadog](https://docs.datadoghq.com/agent/faq/upgrade-to-agent-v6/#what-is-the-agent-v6)). If you are responsible for building or supporting Go applications, a  well-considered logging strategy can help you understand user behavior,  localize errors, and monitor the performance of your applications.

依赖分布式系统的组织通常用 Go 编写他们的应用程序以利用并发特性，如通道和 goroutines（例如，[Heroku](https://blog.golang.org/go-at-heroku)、[Basecamp](https://signalvnoise.com/posts/3897-go-at-basecamp), [Cockroach Labs](https://www.cockroachlabs.com/blog/why-go-was-the-right-choice-for-cockroachdb/) 和 [Datadog](https://docs.datadoghq.com/agent/faq/upgrade-to-agent-v6/#what-is-the-agent-v6))。如果您负责构建或支持 Go 应用程序，一个深思熟虑的日志记录策略可以帮助您了解用户行为、定位错误并监控应用程序的性能。

This post will show you some tools and techniques for managing Golang logs. We’ll begin with the question of which logging package to use for  different kinds of requirements. Next, we’ll explain some techniques for making your logs more searchable and reliable, reducing the resource  footprint of your logging setup, and standardizing your log messages.

这篇文章将向您展示一些管理 Golang 日志的工具和技术。我们将从针对不同类型的需求使用哪个日志包的问题开始。接下来，我们将解释一些技术，使您的日志更易于搜索和可靠，减少日志设置的资源占用，并使日志消息标准化。

## [Know your logging package](https://www.datadoghq.com/blog/go-logging/#know-your-logging-package)

## [了解你的日志包](https://www.datadoghq.com/blog/go-logging/#know-your-logging-package)

Go gives you a wealth of options when choosing a logging package, and we’ll explore several of these below. While [logrus](https://github.com/sirupsen/logrus) is the most popular of the libraries we cover, and helps you implement a [consistent logging format](https://godoc.org/github.com/sirupsen/logrus#Formatter), the others have specialized use cases that are worth mentioning. This section will survey the libraries log, logrus, and glog.

在选择日志包时，Go 为您提供了丰富的选择，我们将在下面探讨其中的几个。虽然 [logrus](https://github.com/sirupsen/logrus) 是我们涵盖的库中最受欢迎的，它可以帮助您实现 [一致的日志格式](https://godoc.org/github.com/sirupsen/logrus#Formatter)，其他人有值得一提的专门用例。本节将调查库 log、logrus 和 glog。

### [Use log for simplicity](https://www.datadoghq.com/blog/go-logging/#use-log-for-simplicity)

### [为简单起见使用日志](https://www.datadoghq.com/blog/go-logging/#use-log-for-simplicity)

Golang’s built-in [logging library](https://golang.org/pkg/log/), called `log`, comes with a default logger that writes to standard error and adds a  timestamp without the need for configuration. You can use these  rough-and-ready logs for local development, when getting fast feedback  from your code may be more important than generating rich, structured  logs.

Golang 的内置 [日志库](https://golang.org/pkg/log/)，称为 `log`，带有一个默认记录器，可以写入标准错误并添加时间戳，无需配置。当从代码中获得快速反馈可能比生成丰富的结构化日志更重要时，您可以将这些粗略的日志用于本地开发。

For example, you can define a division function that returns an error to the caller, rather than exiting the program, when you  attempt to divide by zero.

例如，您可以定义一个除法函数，当您尝试除以零时，该函数向调用者返回错误，而不是退出程序。

```go
package main
import (
    "log"
    "errors"
    "fmt"
 )

func divide(a float32, b float32) (float32, error) {
  if b == 0 {
    return 0, errors.New("can't divide by zero")
  }

  return a / b, nil
}

func main() {

  var a float32 = 10
  var b float32

  ret, err :=  divide(a,b)

  if err != nil{
    log.Print(err)
  }

  fmt.Println(ret)

}
```

Because our example divides by zero, it will output the following log message:

因为我们的示例除以零，它将输出以下日志消息：

```fallback
2019/01/31 11:48:00 can't divide by zero
```

### [Use logrus for formatted logs](https://www.datadoghq.com/blog/go-logging/#use-logrus-for-formatted-logs)

### [对格式化日志使用 logrus](https://www.datadoghq.com/blog/go-logging/#use-logrus-for-formatted-logs)

We [recommend](https://docs.datadoghq.com/logs/log_collection/go/) writing Golang logs using [logrus](https://github.com/sirupsen/logrus), a logging package designed for structured logging that is well-suited  for logging in JSON. The JSON format makes it possible for machines to  easily parse your Golang logs. And since JSON is a well-defined  standard, it makes it straightforward to add context by including new  fields—a parser should be able to pick them up automatically.

我们[推荐](https://docs.datadoghq.com/logs/log_collection/go/) 使用 [logrus](https://github.com/sirupsen/logrus) 编写 Golang 日志，这是一个为结构化日志设计的日志包这非常适合登录 JSON。 JSON 格式使机器可以轻松解析您的 Golang 日志。由于 JSON 是一个定义明确的标准，因此通过包含新字段可以直接添加上下文——解析器应该能够自动获取它们。

Using logrus, you can define standard fields to add to your JSON logs by using the function `WithFields`, as shown below. You can then make calls to the logger at different levels, such as `Info()`, `Warn()` and `Error()`. The logrus library will write the log as JSON automatically and insert  the standard fields, along with any fields you’ve defined on the fly.

使用 logrus，您可以使用函数 `WithFields` 定义要添加到 JSON 日志的标准字段，如下所示。然后，您可以在不同级别调用记录器，例如 `Info()`、`Warn()` 和 `Error()`。 logrus 库将自动将日志写入 JSON 并插入标准字段以及您动态定义的任何字段。

```go
package main
import (
  log "github.com/sirupsen/logrus"
)

func main() {
   log.SetFormatter(&log.JSONFormatter{})

   standardFields := log.Fields{
     "hostname": "staging-1",
     "appname":  "foo-app",
     "session":  "1ce3f6v",
   }

   log.WithFields(standardFields).WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info("My first ssl event from Golang")

}
```

The resulting log will include the message, log level, timestamp, and standard fields in a JSON object:

生成的日志将包含 JSON 对象中的消息、日志级别、时间戳和标准字段：

```fallback
{"appname":"foo-app","float":1.1,"hostname":"staging-1","int":1,"level":"info","msg":"My first ssl eventfrom Golang","session":"1ce3f6v","string":"foo","time":"2019-03-06T13:37:12-05:00"}
```

### [Use glog if you’re concerned about volume](https://www.datadoghq.com/blog/go-logging/#use-glog-if-youre-concerned-about-volume)

### [如果您担心音量，请使用 glog](https://www.datadoghq.com/blog/go-logging/#use-glog-if-youre-concerned-about-volume)

Some logging libraries allow you to enable or disable logging at specific  levels, which is useful for keeping log volume in check when moving  between development and production. One such library is [`glog`](https://godoc.org/github.com/golang/glog), which lets you use flags at the command line (eg, `-v` for verbosity) to set the logging level when you run your code. You can then use a [`V()` function](https://godoc.org/github.com/golang/glog#V) in `if` statements to write your Golang logs only at a certain log level.

一些日志库允许您在特定级别启用或禁用日志记录，这对于在开发和生产之间移动时检查日志量非常有用。一个这样的库是 [`glog`](https://godoc.org/github.com/golang/glog)，它允许你在命令行中使用标志（例如，`-v`表示详细)来设置日志记录运行代码时的级别。然后，您可以在 `if` 语句中使用 [`V()` 函数](https://godoc.org/github.com/golang/glog#V) 仅在特定日志级别写入 Golang 日志。

For example, you can use glog to write the same “Can’t divide by zero”  error from earlier, but only if you’re logging at the verbosity level of `2`. You can set the verbosity to [any signed 32-bit integer](https://godoc.org/github.com/golang/glog#Level), or use the functions `Info()`, `Warning()`, `Error()`, and `Fatal()` to [assign verbosity levels](https://github.com/golang/glog/blob/master/glog.go#L97) `0` through `3` (respectively ).

例如，您可以使用 glog 编写与之前相同的“不能被零除”错误，但前提是您以“2”的详细级别进行日志记录。您可以将详细程度设置为 [任何有符号的 32 位整数](https://godoc.org/github.com/golang/glog#Level)，或使用函数`Info()`、`Warning()`， `Error()` 和 `Fatal()` 用于 [分配详细级别](https://github.com/golang/glog/blob/master/glog.go#L97) 从 `0` 到 `3`（分别为)。

```fallback
   if err != nil && glog.V(2){
    glog.Warning(err)
  }
```

You can make your application less resource  intensive by logging only certain levels in production. At the same  time, if there’s no impact on users, it’s often a good idea to log as  many interactions with your application as possible, then use log  management software like Datadog to find the data you need for your  investigation

您可以通过在生产中仅记录某些级别来减少应用程序的资源密集度。同时，如果对用户没有影响，通常最好记录尽可能多的与应用程序的交互，然后使用诸如 Datadog 之类的日志管理软件来查找调查所需的数据

## [Best practices for writing and storing Golang logs](https://www.datadoghq.com/blog/go-logging/#best-practices-for-writing-and-storing-golang-logs)

##  [编写和存储 Golang 日志的最佳实践](https://www.datadoghq.com/blog/go-logging/#best-practices-for-writing-and-storing-golang-logs)

Once you’ve chosen a logging library, you’ll also want to plan for where in  your code to make calls to the logger, how to store your logs, and how  to make sense of them. In this section, we’ll recommend a series of best practices for organizing your Golang logs:

选择日志库后，您还需要计划在代码中调用记录器的位置、如何存储日志以及如何理解它们。在本节中，我们将推荐一系列组织 Golang 日志的最佳实践：

- Make calls to the logger from within your main application process, [not within goroutines](https://www.datadoghq.com/blog/go-logging/#avoid-logging-in-goroutines).
- Write logs from your application [to a local file](https://www.datadoghq.com/blog/go-logging/#write-your-logs-to-a-file), even if you'll ship them to a central platform later.
- [Standardize your logs](https://www.datadoghq.com/blog/go-logging/#implement-a-standard-logging-interface) with a set of predefined messages.
- Send your logs to a [central platform](https://www.datadoghq.com/blog/go-logging/#centralize-golang-logs) so you can analyze and aggregate them.
- Use HTTP headers and unique IDs to log user behavior [across microservices](https://www.datadoghq.com/blog/go-logging/#track-logs-across-microservices).

- 从您的主应用程序进程中调用记录器，[不在 goroutines](https://www.datadoghq.com/blog/go-logging/#avoid-logging-in-goroutines)。
- 将您的应用程序中的日志写入 [到本地文件](https://www.datadoghq.com/blog/go-logging/#write-your-logs-to-a-file)，即使您将发送它们稍后到中央平台。
- [标准化您的日志](https://www.datadoghq.com/blog/go-logging/#implement-a-standard-logging-interface) 使用一组预定义的消息。
- 将您的日志发送到 [中央平台](https://www.datadoghq.com/blog/go-logging/#centralize-golang-logs)，以便您可以对其进行分析和汇总。
- 使用 HTTP 标头和唯一 ID 记录用户行为 [跨微服务](https://www.datadoghq.com/blog/go-logging/#track-logs-across-microservices)。

### [Avoid declaring goroutines for logging](https://www.datadoghq.com/blog/go-logging/#avoid-declaring-goroutines-for-logging)

### [避免为日志声明 goroutines](https://www.datadoghq.com/blog/go-logging/#avoid-declaring-goroutines-for-logging)

There are two reasons to avoid creating your own goroutines to handle writing logs. First, it can lead to concurrency issues, as duplicates of the  logger would attempt to access the same `io.Writer`. Second,  logging libraries usually start goroutines themselves, managing any  concurrency issues internally, and starting your own goroutines will  only interfere.

避免创建自己的 goroutine 来处理写入日志的原因有两个。首先，它会导致并发问题，因为记录器的重复项会尝试访问同一个 `io.Writer`。其次，日志库通常自己启动 goroutines，在内部管理任何并发问题，启动你自己的 goroutines 只会干扰。

### [Write your logs to a file](https://www.datadoghq.com/blog/go-logging/#write-your-logs-to-a-file) 

### [将您的日志写入文件](https://www.datadoghq.com/blog/go-logging/#write-your-logs-to-a-file)

Even if you’re shipping your logs to a central platform, we recommend  writing them to a file on your local machine first. You will want to  make sure your logs are always available locally and not lost in the  network. In addition, writing to a file means that you can decouple the  task of writing your logs from the task of sending them to a central  platform. Your applications themselves will not need to establish  connections or stream your logs, and you can leave these jobs to  specialized software like the Datadog Agent. If you're running your Go  applications within a containerized infrastructure that does not already include persistent storage—eg, containers running on [AWS Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-task-storage.html)—you may want to configure your log management tool to collect logs directly from your containers' STDOUT and STDERR streams (this is handled  differently in [Docker](https://docs.docker.com/config/containers/logging/configure/) and [Kubernetes](https://kubernetes.io/docs/concepts/cluster-administration/logging/)).

即使您要将日志传送到中央平台，我们也建议您先将它们写入本地计算机上的文件。您需要确保您的日志在本地始终可用，并且不会在网络中丢失。此外，写入文件意味着您可以将写入日志的任务与将它们发送到中央平台的任务分离。您的应用程序本身不需要建立连接或流式传输您的日志，您可以将这些工作留给像 Datadog Agent 这样的专用软件。如果您在尚未包含持久存储的容器化基础设施中运行 Go 应用程序，例如，在 [AWS Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-task-storage.html)——你可能想要配置你的日志管理工具来直接从你容器的 STDOUT 和 STDERR 流中收集日志（这在 [Docker](https://docs.docker.com/config/container/logging/configure/) 和 [Kubernetes](https://kubernetes.io/docs/concepts/cluster-administration/logging/))。

### [Implement a standard logging interface](https://www.datadoghq.com/blog/go-logging/#implement-a-standard-logging-interface)

### [实现标准日志接口](https://www.datadoghq.com/blog/go-logging/#implement-a-standard-logging-interface)

When writing calls to loggers from within their code, teams teams often use  different attribute names to describe the same thing. Inconsistent  attributes can confuse users and make it impossible to correlate logs  that should form part of the same picture. For example, two developers  might log the same error, a missing client name when handling an upload, in different ways.

在从他们的代码中编写对记录器的调用时，团队团队经常使用不同的属性名称来描述同一件事。不一致的属性会使用户感到困惑，并且无法将应该构成同一图片一部分的日志相关联。例如，两个开发人员可能会以不同的方式记录相同的错误，即在处理上传时缺少客户端名称。

![Golang logs for the same error with different messages from different locations.]()Golang logs for the same error with different messages from different locations.

A good way to enforce standardization is to create an interface between  your application code and the logging library. The interface contains  predefined log messages that implement a certain format, making it  easier to investigate issues by ensuring that log messages can be  searched, grouped, and filtered.

强制标准化的一个好方法是在您的应用程序代码和日志库之间创建一个接口。该界面包含实现某种格式的预定义日志消息，通过确保可以搜索、分组和过滤日志消息，使调查问题变得更加容易。

![Golang logs for an error using a standard interface to create a consistent message.]()Golang logs for an error using a standard interface to create a consistent message.

In this example, we’ll declare an `Event` type with a predefined message. Then we’ll use `Event` messages to make calls to a logger. Teammates can write Golang logs by  providing a minimal amount of custom information, letting the  application do the work of implementing a standard format.

在这个例子中，我们将声明一个带有预定义消息的“事件”类型。然后我们将使用 `Event` 消息来调用记录器。队友可以通过提供最少量的自定义信息来编写 Golang 日志，让应用程序完成实现标准格式的工作。

First, we’ll write a `logwrapper` package that developers can include within their code.

首先，我们将编写一个“logwrapper”包，开发人员可以将其包含在他们的代码中。

```go
package logwrapper

import (
  "github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
  id      int
  message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
  *logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
    var baseLogger = logrus.New()

    var standardLogger = &StandardLogger{baseLogger}

    standardLogger.Formatter = &logrus.JSONFormatter{}

    return standardLogger
}

// Declare variables to store log messages as new Events
var (
  invalidArgMessage      = Event{1, "Invalid arg: %s"}
  invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
  missingArgMessage      = Event{3, "Missing arg: %s"}
)

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
  l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
  l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
  l.Errorf(missingArgMessage.message, argumentName)
}
```

To use our logging interface, we only have to include it in our code and make calls to an instance of `StandardLogger`.

要使用我们的日志接口，我们只需要将它包含在我们的代码中并调用 `StandardLogger` 的一个实例。

```go
package main

import (
  li "<PATH_TO_PACKAGE>/logwrapper"
)

func main() {

  var standardLogger = := li.NewLogger()

  // You can then call a method of our standard logger in the context of an error
  // you would like to log.

  standardLogger.InvalidArgValue("client", "nil")

}
```

When we run our code, we’ll get the following JSON log:

当我们运行我们的代码时，我们将获得以下 JSON 日志：

```fallback
{"level":"error","msg":"Invalid value for argument: client: nil","time":"2019-03-04T11:21:07-05:00"}
```

### [Centralize Golang logs](https://www.datadoghq.com/blog/go-logging/#centralize-golang-logs)

### [集中化 Golang 日志](https://www.datadoghq.com/blog/go-logging/#centralize-golang-logs)

If your application is deployed across a cluster of hosts, it’s not  sustainable to SSH into each one in order to tail, grep, and investigate your logs. A more scalable alternative is to pass logs from local files to a central platform.

如果您的应用程序部署在一组主机上，为了跟踪、grep 和调查您的日志，将 SSH 连接到每个主机是不可持续的。一个更具可扩展性的替代方案是将日志从本地文件传递到中央平台。

One solution is to use the [Golang syslog package](https://golang.org/pkg/log/syslog/) to forward logs from throughout your infrastructure to a single syslog server.

一种解决方案是使用 [Golang syslog 包](https://golang.org/pkg/log/syslog/) 将日志从整个基础架构转发到单个 syslog 服务器。

**Customize and streamline your Golang log management with Datadog.**

**使用 Datadog 自定义和简化您的 Golang 日志管理。**

Another is to use a log management solution. Datadog, for example, can tail  your log files and forward logs to a central platform for processing and analysis.

另一个是使用日志管理解决方案。例如，Datadog 可以跟踪您的日志文件并将日志转发到中央平台进行处理和分析。

<video loop="" muted="muted" playsinline="" width="100%" height="auto">
</video>

<video loop="" muted="muted" playinline="" width="100%" height="auto">
</视频>

The Datadog Log Explorer view can show Golang logs from various sources.

Datadog Log Explorer 视图可以显示来自各种来源的 Golang 日志。

You can use attributes to graph the values of certain log fields over time, sorted by group. For example, you could track the number of errors by `service` to let you know if there’s an incident in one of your services. Showing logs from only the `go-logging-demo` service, we can see how many error logs this service has produced in a given interval.

您可以使用属性来绘制某些日志字段随时间变化的值，按组排序。例如，您可以通过“服务”跟踪错误数量，以让您知道您的一项服务是否发生了事件。仅显示来自 go-logging-demo 服务的日志，我们可以看到该服务在给定时间间隔内产生了多少错误日志。

![Grouping Golang logs by status.]()

You can also use attributes to drill down into possible causes, for  instance seeing if a spike in error logs belongs to a specific host. You can then create an automated alert based on the values of your logs.

您还可以使用属性深入研究可能的原因，例如查看错误日志中的峰值是否属于特定主机。然后，您可以根据日志的值创建自动警报。

### [Track Golang logs across microservices](https://www.datadoghq.com/blog/go-logging/#track-golang-logs-across-microservices)

### [跨微服务跟踪 Golang 日志](https://www.datadoghq.com/blog/go-logging/#track-golang-logs-across-microservices)

When troubleshooting an error, it’s often helpful to see what pattern of  behavior led to it, even if that behavior involves a number of  microservices. You can achieve this with distributed tracing,  visualizing the order in which your application executes functions,  database queries, and other tasks, and following these execution steps  as they make their way through a network. One way to implement  distributed tracing within your logs is to pass contextual information  as HTTP headers.

在对错误进行故障排除时，查看导致该错误的行为模式通常很有帮助，即使该行为涉及许多微服务。您可以通过分布式跟踪、可视化应用程序执行函数、数据库查询和其他任务的顺序，并在它们通过网络时遵循这些执行步骤来实现这一点。在日志中实现分布式跟踪的一种方法是将上下文信息作为 HTTP 标头传递。

In this example, one microservice receives a request and checks for a trace ID in the `x-trace` header, generating one if it doesn’t exist. When making a request to  another microservice, we then generate a new spanID—for this and for  every request—and add it to the header `x-span`.

在这个例子中，一个微服务收到一个请求并检查 `x-trace` 标头中的跟踪 ID，如果它不存在则生成一个。当向另一个微服务发出请求时，我们会为此和每个请求生成一个新的 spanID，并将其添加到标头“x-span”中。

```fallback
func microService1(w http.ResponseWriter, r *http.Request) {
  client := &http.Client{}
  trace := r.Header.Get("x-trace")

  if ( trace == "") {
    trace = generateTraceId()
  }

  span := generateSpanId()

  // Hit the second microservice with the appropriate headers
  reqService2, _ := http.NewRequest("GET", "<ADDRESS>", nil)
  reqService2.Header.Add("x-trace", trace)
  reqService2.Header.Add("x-span", span)
  resService2, _ := client.Do(reqService2)

}
```

Downstream microservices use the `x-span` headers of incoming requests to specify the parents of the spans they generate, and send that information as the `x-parent` header to the next microservice in the chain.

下游微服务使用传入请求的 `x-span` 标头来指定它们生成的跨度的父级，并将该信息作为 `x-parent` 标头发送到链中的下一个微服务。

```fallback
func microService2(w http.ResponseWriter, r *http.Request) {

  trace := r.Header.Get("x-trace")
  span := generateSpanId()

  parent := r.Header.Get("x-span")
  if (trace == "") {
    w.Header().Set("x-parent", parent)
  }

  w.Header().Set("x-trace", trace)
  w.Header().Set("x-span", span)
  if (parent == "") {
    w.Header().Set("x-parent", span)
  }

  w.WriteHeader(http.StatusOK)
  io.WriteString(w, fmt.Sprintf(aResponseMessage, 2, trace, span, parent))

}
```

If an error occurs in one of our microservices, we can use the `trace`, `parent`, and `span` attributes to see the route that a request has taken, letting us know  which hosts—and possibly which parts of the application code—to  investigate.

如果我们的一个微服务中发生错误，我们可以使用 `trace`、`parent` 和 `span` 属性来查看请求所采用的路由，让我们知道哪些主机 - 以及应用程序的哪些部分代码——调查。

In the first microservice:

在第一个微服务中：

```fallback
{"appname":"go-logging","level":"debug","msg":"Hello from Microservice One","trace":"eUBrVfdw","time":"2017-03-02T15:29:26+01:00","span":"UzWHRihF"}
```

In the second:

在第二：

```fallback
{"appname":"go-logging","level":"debug","msg":"Hello from Microservice Two","parent":"UzWHRihF","trace":"eUBrVfdw","time":"2017-03-02T15:29:26+01:00","span":"DPRHBMuE"}
```

If you want to dig more deeply into Golang tracing possibilities, you can use a tracing library such as [OpenTracing](https://opentracing.io/) or a monitoring platform that supports distributed tracing for Go applications. For example, Datadog can [automatically build](https://www.datadoghq.com/blog/service-map/) a map of services using data from its [Golang tracing library](https://docs.datadoghq.com/tracing/languages/go/); [visualize trends](https://www.datadoghq.com/blog/trace-search-high-cardinality-data/) in your traces over time; and [let you know](https://www.datadoghq.com/blog/watchdog/) about services with unusual request rates, error rates, or latency.

如果您想更深入地挖掘 Golang 跟踪的可能性，您可以使用跟踪库，例如 [OpenTracing](https://opentracing.io/) 或支持 Go 应用程序分布式跟踪的监控平台。例如，Datadog 可以使用其 [Golang 跟踪库](https://docs.datadoghq.com)中的数据[自动构建](https://www.datadoghq.com/blog/service-map/) 服务地图/tracing/languages/go/); [可视化趋势](https://www.datadoghq.com/blog/trace-search-high-cardinality-data/) 在您的跟踪中随时间推移；和 [让您知道](https://www.datadoghq.com/blog/watchdog/) 了解具有异常请求率、错误率或延迟的服务。

![An example of a visualization showing traces of requests between microservices.]()An example of a visualization showing traces of requests between microservices.

## [Clean and comprehensive Golang logs](https://www.datadoghq.com/blog/go-logging/#clean-and-comprehensive-golang-logs)

## [干净全面的 Golang 日志](https://www.datadoghq.com/blog/go-logging/#clean-and-comprehensive-golang-logs)

In this post, we’ve highlighted the benefits and tradeoffs of several Go  logging libraries. We’ve also recommended ways to ensure that your logs  are available and accessible when you need them, and that the  information they contain is consistent and easy to analyze.

在这篇文章中，我们重点介绍了几个 Go 日志库的优点和缺点。我们还推荐了一些方法来确保您的日志在您需要时可用和访问，并且它们包含的信息一致且易于分析。

To start analyzing all of your Go logs with Datadog, sign up for a [free trial](https://www.datadoghq.com/blog/go-logging/#).

要开始使用 Datadog 分析您的所有 Go 日志，请注册[免费试用](https://www.datadoghq.com/blog/go-logging/#)。

Related Posts

相关文章

[Monitor AWS FSx audit logs with Datadog](https://www.datadoghq.com/blog/amazon-fsx-audit-logs-monitoring/)

[使用 Datadog 监控 AWS FSx 审计日志](https://www.datadoghq.com/blog/amazon-fsx-audit-logs-monitoring/)

[Monitor real-time Salesforce logs with Datadog](https://www.datadoghq.com/blog/monitor-salesforce-logs-datadog/)

[使用 Datadog 监控实时 Salesforce 日志](https://www.datadoghq.com/blog/monitor-salesforce-logs-datadog/)

[Monitor AWS App Runner with Datadog](https://www.datadoghq.com/blog/aws-app-runner-monitoring/)

[使用 Datadog 监控 AWS App Runner](https://www.datadoghq.com/blog/aws-app-runner-monitoring/)

[Monitor Cloudflare logs and metrics with Datadog](https://www.datadoghq.com/blog/cloudflare-monitoring-datadog/) 

[使用 Datadog 监控 Cloudflare 日志和指标](https://www.datadoghq.com/blog/cloudflare-monitoring-datadog/)

