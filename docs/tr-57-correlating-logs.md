# Correlating Logs

# 关联日志

**01.10.2020 19:39** From: https://filipnikolovski.com/posts/correlating-logs/

When something goes wrong in your system, logs are crucial to finding out  exactly what’s happened. Usually, this involves following the logs as a trail of breadcrumbs that lead to the root cause of the failure. If your application is generating a lot of logs, it can become strenuous to tie everything together that reveals the failing scenario.

当您的系统出现问题时，日志对于准确找出发生了什么至关重要。通常，这涉及将日志作为导致失败根本原因的面包屑痕迹跟踪。如果您的应用程序生成大量日志，则将所有显示失败场景的内容联系在一起会变得很费力。

This  can become especially challenging in a distributed system, where one  HTTP request to your API can pass through dozens of different services, each outputting logs that have no context of the flow of the request.

这在分布式系统中变得特别具有挑战性，其中对您的 API 的一个 HTTP 请求可以通过几十个不同的服务，每个服务输出的日志没有请求流的上下文。

In this post firstly we’ll define what is a structured log and then explore a solution on how to tie in the mutual logs in a **Go application**, using context, structured logging, and correlation ids. Although the following example  will be entirely in Go, the principle should be the same for apps  written in different programming languages.

在这篇文章中，我们将首先定义什么是结构化日志，然后探索如何使用上下文、结构化日志记录和关联 id 将 **Go 应用程序**中的相互日志联系起来的解决方案。尽管以下示例将完全使用 Go，但对于使用不同编程语言编写的应用程序，其原理应该是相同的。

## Structured logs

## 结构化日志

Before we talk about correlating logs, we need a way to make sense of the data that we are writing. Logs are typically just an unstructured text which makes it hard to extract useful information from them by another  machine. To be able to query or filter them by something, the logs need  to be written in a format that can be easily parsed and indexed.

在我们谈论关联日志之前，我们需要一种方法来理解我们正在编写的数据。日志通常只是一个非结构化的文本，这使得另一台机器很难从中提取有用的信息。为了能够通过某种方式查询或过滤它们，日志需要以易于解析和索引的格式编写。

Here is an example of an unstructured log record:

以下是非结构化日志记录的示例：

```go
INFO: 2009/11/10 23:00:00 192.168.0.2 Hello World!
```


Versus a structured one:

与结构化的相比：

```json
{"time": 1562212768, "host": "192.168.0.2", "level": "INFO", "message": "Hello World!"}
```


The structured log essentially contains the same  information as the unstructured one, but the key difference is that the  message is in a format that another machine can understand (it can be JSON, XML, whatever) and the  fields can be identified, indexed, and other programs such as logging  systems can analyze them so that later on we can use these fields to search and  filter results.

结构化日志本质上包含与非结构化日志相同的信息，但关键区别在于消息采用另一台机器可以理解的格式（可以是 JSON、XML 等），并且字段可以被识别、索引和日志系统等其他程序可以分析它们，以便稍后我们可以使用这些字段来搜索和过滤结果。

Let’s see this in action in a **Go app**. There are plenty of libraries for structured logging in Go, some of the more popular ones are [Zerolog](https://github.com/rs/zerolog),[Zap](https://github.com/uber-go /zap), and [Logrus](https://github.com/sirupsen/logrus).

让我们在 **Go 应用程序**中看看这个。 Go中有很多用于结构化日志的库，其中一些比较流行的是[Zerolog](https://github.com/rs/zerolog)、[Zap](https://github.com/uber-go/zap) 和 [Logrus](https://github.com/sirupsen/logrus)。

We’ll use Zerolog as an example. From the package description, it says that  it provides “a fast and simple logger dedicated to JSON output”. It has features such as contextual fields, log levels, and passing the logger  by context, which is exactly what we’re looking for.

我们将以 Zerolog 为例。从包的描述来看，它说它提供了“一个专门用于 JSON 输出的快速而简单的记录器”。它具有上下文字段、日志级别和按上下文传递记录器等功能，这正是我们正在寻找的。

Here’s how we can create the log from the example above:

以下是我们如何从上面的示例创建日志：

```go
package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    log.Info().
        Str("host", "192.168.0.2").
        Msg("Hello World!")
}
// Output: {"time": 1562212768, "host": "192.168.0.2", "level": "info", "message": "Hello World!"}
```


## Correlation IDs

## 相关 ID

In a typical web app you would have your controllers, middleware, services and domain logic code that will be invoked on each user request. As we go down the stack, we can have multiple logs that are being written for  the same action, each log telling only a part of the story. Then, when  something bad happens and we want to investigate, we would need to bring the pieces together  like a puzzle, in order to make sense of the situation.

在典型的 Web 应用程序中，您将拥有将在每个用户请求上调用的控制器、中间件、服务和域逻辑代码。当我们沿着堆栈向下移动时，我们可以为同一操作编写多个日志，每个日志只讲述故事的一部分。然后，当发生不好的事情并且我们想要调查时，我们需要像拼图一样将各个部分拼凑起来，以便了解情况。

The  problem gets more convoluted in a distributed systems scenario, where  separate services are part of this one big user action, each outputting  their own logs that are difficult to follow and make sense of, when it comes to understanding  the bigger picture. Finding the service that is the culprit will be like playing detective.

在分布式系统场景中，问题变得更加复杂，其中单独的服务是这一大用户操作的一部分，每个服务都输出自己的日志，当涉及到理解更大的图景时，这些日志难以理解和理解。找到罪魁祸首的服务就像扮演侦探一样。

To avoid having to become Sherlock as you  intuitively search through the big pile of messages, you could simply  connect them by having an additional field in the structure of the log - a **correlation id**. This id can be generated on each request, somewhere on top of the call chain and propagated down to be used as part of each logged message.

为了避免在直观地搜索大量消息时变成 Sherlock，您可以通过在日志结构中添加一个附加字段 - **相关性 id** 来简单地连接它们。这个 id 可以在每个请求上生成，位于调用链顶部的某个地方，并向下传播以用作每个记录的消息的一部分。

We can illustrate this concept in a Go app by creating and using a **HTTP middleware**, the **context object**, and our trusty old logger. 

我们可以通过创建和使用 **HTTP 中间件**、**上下文对象** 和我们值得信赖的旧记录器来在 Go 应用程序中说明这个概念。

The basic principle of a middleware is that it’s a way of organizing a  shared functionality that we want to run on each HTTP request. This code usually sits between the router and the application  controllers.

中间件的基本原则是，它是一种组织共享功能的方式，我们希望在每个 HTTP 请求上运行这些功能。此代码通常位于路由器和应用程序控制器之间。

**Note:** I won't cover the whole story on how to create and use middlewares since there are [great posts](https://www.alexedwards.net/blog/making-and-using-middleware)about it that go into [detail](https://drstearns.github.io/tutorials/gomiddleware/), as well as some great libraries (such as [Negroni](https://github.com/urfave/negroni)) , check them out for more info.

**注意：** 我不会介绍如何创建和使用中间件的整个故事，因为有 [很棒的帖子](https://www.alexedwards.net/blog/making-and-using-middleware)关于它进入 [详细](https://drstearns.github.io/tutorials/gomiddleware/)，以及一些很棒的库（例如 [Negroni](https://github.com/urfave/negroni)） ，查看它们以获取更多信息。

Anyway, the gist of it is that a middleware is essentially a function that implements the [http.Handler](https://pkg.go.dev/net/http?tab=doc#Handler) interface and it takes a handler function as a parameter which is the `next` handler that should be invoked, after the middleware code is done.

无论如何，它的要点是中间件本质上是一个实现 [http.Handler](https://pkg.go.dev/net/http?tab=doc#Handler) 接口的函数，它需要一个处理函数作为参数，它是在中间件代码完成后应该调用的“next”处理程序。

```go
func someMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Middleware logic..
    next.ServeHTTP(w, r)
  })
}
```


This allows us to chain multiple handlers  together, so in our scenario we’ll end up having a handler that  generates IDs and passes them down that wraps each application handler.

这允许我们将多个处理程序链接在一起，因此在我们的场景中，我们最终将拥有一个处理程序，该处理程序生成 ID 并将它们向下传递以包装每个应用程序处理程序。

The flow of control will look something like this:

控制流将如下所示：

```
Router --> Generate correlation ID Middleware --> App handler
```


## An example

##  一个例子

We can illustrate this in code with a simple example, where we will:

我们可以用一个简单的例子在代码中说明这一点，我们将：

- Fetch the correlation id from a header if it’s present (something like **X-Correlation-Id**)
- In case the header is not present, we’ll generate a random id and add it to the response header. This will make our debugging easier since we’ll receive the id in our response.
- Pass it down to our logger instance
- Call the next handler

- 如果存在，则从标头中获取相关 ID（类似于 **X-Correlation-Id**）
- 如果标头不存在，我们将生成一个随机 ID 并将其添加到响应标头中。这将使我们的调试更容易，因为我们将在响应中收到 id。
- 将其传递给我们的记录器实例
- 调用下一个处理程序

Okay, so first thing that we’ll need to do is initialize the logger and add it to the request context. This will allow us to use the logger object in our handlers. We’ll create a middleware that will inject the log into the request context, so that we can fetch the logger later on in our handlers.

好的，所以我们需要做的第一件事是初始化记录器并将其添加到请求上下文中。这将允许我们在处理程序中使用记录器对象。我们将创建一个将日志注入请求上下文的中间件，以便我们稍后可以在处理程序中获取记录器。

```go
func logMiddleware(log zerolog.Logger) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            l := log.With().Logger()
 
            // l.WithContext returns a copy of the context with the log object associated
            r = r.WithContext(l.WithContext(r.Context()))
            
            next.ServeHTTP(w, r)
        })
    }
}
```


Next up is our middleware that will handle the correlation id. This handler will set a unique id to the request, which can be fetched from a header that we'll define (such as X-Correlation-Id), or if the header is not present in the request, the handler will generate a new unique id.

接下来是我们的中间件，它将处理关联 ID。这个处理程序将为请求设置一个唯一的 id，可以从我们将定义的标头（例如 X-Correlation-Id）中获取，或者如果请求中不存在标头，处理程序将生成一个新的唯一身份。

This id will be added as a field to the log object and will be present in  the log structure for each message created down the line. The id is also added as a response header, which is great for debugging  purposes, since we’ll immediately have the id to query the log system.

这个 id 将作为一个字段添加到日志对象中，并将出现在日志结构中，用于创建的每条消息。 id 也被添加为响应头，这对于调试目的非常有用，因为我们将立即拥有 id 来查询日志系统。

In case we’ll need to send the id to another service, we’ll add it to the  request context so that we can easily fetch it when the need arises.

如果我们需要将 id 发送到另一个服务，我们会将其添加到请求上下文中，以便在需要时可以轻松获取它。

```go
func correlationIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        id := r.Header.Get("X-Correlation-Id")
        if id == "" {
            // generate new version 4 uuid
            newid := uuid.New()
            id = newid.String()
        }
        // set the id to the request context
        ctx = context.WithValue(ctx, "correlation_id", id)
        r = r.WithContext(ctx)
        // fetch the logger from context and update the context
        // with the correlation id value
        log := zerolog.Ctx(ctx)
        log.UpdateContext(func(c zerolog.Context) zerolog.Context {
            return c.Str("correlation_id", id)
        })
        // set the response header
        w.Header().Set("X-Correlation-Id", id)
        next.ServeHTTP(w, r)
    })
}
```




**Note:** Zerolog provides a handy integration with the `net/http` package, and it has a bunch of useful middleware that you can use if you choose to use this as your logging library. These are just  simplified examples, you can choose a more robust and flexible solution  provided by the library. You can read more about that [here](https://github.com/rs/zerolog#integration-with-nethttp).

**注意：** Zerolog 提供了与 `net/http` 包的便捷集成，并且它有一堆有用的中间件，如果您选择将其用作日志库，则可以使用这些中间件。这些只是简化的示例，您可以选择库提供的更健壮和灵活的解决方案。您可以在 [此处](https://github.com/rs/zerolog#integration-with-nethttp) 阅读更多相关信息。

Okay so now let’s see our middleware in action and test it out in our full example:

好的，现在让我们看看我们的中间件的运行情况，并在我们的完整示例中对其进行测试：

```go
package main

import (
    "context"
    "net/http"
    "os"
    "time"

    "github.com/google/uuid"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func logMiddleware(log zerolog.Logger) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            l := log.With().Logger()

            // l.WithContext returns a copy of the context with the log object associated
            r = r.WithContext(l.WithContext(r.Context()))

            next.ServeHTTP(w, r)
        })
    }
}

func correlationIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        id := r.Header.Get("X-Correlation-Id")
        if id == "" {
            // generate new version 4 uuid
            newid := uuid.New()
            id = newid.String()
        }
        // set the id to the request context
        ctx = context.WithValue(ctx, "correlation_id", id)
        r = r.WithContext(ctx)
        // fetch the logger from context and update the context
        // with the correlation id value
        log := zerolog.Ctx(ctx)
        log.UpdateContext(func(c zerolog.Context) zerolog.Context {
            return c.Str("correlation_id", id)
        })
        // set the response header
        w.Header().Set("X-Correlation-Id", id)
        next.ServeHTTP(w, r)
    })
}

func testHandler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    logger := log.Ctx(r.Context())

    logger.Info().Msg("Processing request.")

    w.Write([]byte("OK"))

    logger.Info().Dur("elapsed", time.Since(start)).Msg("Done.")
}

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    log := zerolog.New(os.Stdout).With().
        Timestamp().
        Logger()

    mux := http.NewServeMux()

    mux.Handle("/",
        logMiddleware(log)(
            correlationIDMiddleware(
                http.HandlerFunc(testHandler),
            ),
        ),
    )

    err := http.ListenAndServe(":8080", mux)
    log.Fatal().Err(err)
}
```


Here we have a `testHandler` function that we've wrapped in our log and correlation ID middleware (remember you can chain multiple middlewares together) and we have added two log messages so we can show that the correlation id is correctly added to our logs .

这里我们有一个 `testHandler` 函数，我们已经将它包装在我们的日志和相关 ID 中间件中（请记住，您可以将多个中间件链接在一起）并且我们添加了两条日志消息，以便我们可以显示相关 ID 已正确添加到我们的日志中.

When we start the example program and execute a request, the logs will appear in the terminal like so:

当我们启动示例程序并执行请求时，日志将显示在终端中，如下所示：

```json
{"level":"info","correlation_id":"2fda2736-7d2d-441e-82a2-24df9a1199a0","time":1601569068,"message":"Processing request."}

{"level":"info","correlation_id":"2fda2736-7d2d-441e-82a2-24df9a1199a0","elapsed":0.030087,"time":1601569068,"message":"Done."}
```


When running `curl -I localhost:8080` in the command line, we could also see that our correlation id is added in the response headers:

在命令行中运行 `curl -I localhost:8080` 时，我们还可以看到我们的关联 ID 被添加到响应头中：

```
HTTP/1.1 200 OK
X-Correlation-Id: 2fda2736-7d2d-441e-82a2-24df9a1199a0
Date: Thu, 01 Oct 2020 16:23:37 GMT
Content-Length: 2
Content-Type: text/plain;charset=utf-8
```


## Conclusion

##  结论

Troubleshooting a  problem in a system is difficult, even more so when the components that  the system is made of are distributed on multiple machines. The logs are not guaranteed to arrive sequentially, and it’s a challenge trying to make sense of  what’s happened.

对系统中的问题进行故障排除很困难，当组成系统的组件分布在多台机器上时更是如此。日志不能保证按顺序到达，试图弄清楚发生了什么是一个挑战。

Instead of wasting time on piecing together the  scattered information, a correlation ID will provide cohesion to the  logs, unifying the related messages and providing you a way of tracking  down the problem sooner, rather than later.

相关性 ID 不会浪费时间将分散的信息拼凑在一起，而是会为日志提供内聚力，统一相关消息，并为您提供一种更快而不是更晚跟踪问题的方法。

### LATEST POSTS
- [What is the future of databases?](https://filipnikolovski.com/posts/what-is-the-future-of-databases/)
- [TIL: eBPF is awesome](https://filipnikolovski.com/posts/ebpf/)
- [Correlating Logs](https://filipnikolovski.com/posts/correlating-logs/)
- [Bazel Performance in a CI Environment](https://filipnikolovski.com/posts/bazel-performance-in-a-ci-environment/)
- [Avoiding Pitfalls During Service Deployments](https://filipnikolovski.com/posts/avoiding-pitfalls-during-service-deployments/) 

###  最新帖子
- [数据库的未来是什么？](https://filipnikolovski.com/posts/what-is-the-future-of-databases/)
- [TIL: eBPF 很棒](https://filipnikolovski.com/posts/ebpf/)
- [相关日志](https://filipnikolovski.com/posts/correlating-logs/)
- [CI 环境中的 Bazel 性能](https://filipnikolovski.com/posts/bazel-performance-in-a-ci-environment/)
- [避免服务部署过程中的陷阱](https://filipnikolovski.com/posts/avoiding-pitfalls-during-service-deployments/)

