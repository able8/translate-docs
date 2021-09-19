# Make resilient Go net/http servers using timeouts, deadlines and context cancellation

# 使用超时、截止日期和上下文取消来构建有弹性的 Go net/http 服务器

January 5, 2020 · 15 min · Ilija

2020 年 1 月 5 日 · 15 分钟 · 伊利亚

When it comes to timeouts, there are two types of people: those who know how tricky they can be, and those who are yet to find out.

说到超时，有两种人：知道自己有多棘手的人，以及尚未发现的人。

As tricky as they are, timeouts are a reality in the connected world we live in. As I am writing this, on the other side of the table, two persons are typing on their smartphones, probably chatting to people very far from them. All made possible because of networks.

尽管很棘手，但在我们生活的互联世界中，超时是一个现实。在我写这篇文章的时候，在桌子的另一边，两个人正在用他们的智能手机打字，可能正在和离他们很远的人聊天。一切都因为网络而成为可能。

Networks and all their intricacies are here to stay, and we, who write servers for the web, have to know how to use them efficiently and guard against their deficiencies.

网络及其所有错综复杂的事物都将继续存在，而我们为 Web 编写服务器，必须知道如何有效地使用它们并防止它们的缺陷。

Without further ado, let’s look at timeouts and how they affect our `net/http` servers.

事不宜迟，让我们看看超时以及它们如何影响我们的 `net/http` 服务器。

## Server timeouts - first principles

## 服务器超时 - 首要原则

In web programming, the general classification of timeouts is client and server timeouts. What inspired me to dive into this topic was an interesting server timeout problem I found myself in. That’s why, in this article, we will focus on server-side timeouts.

在 Web 编程中，超时的一般分类是客户端和服务器超时。启发我深入研究这个主题的是我发现自己遇到的一个有趣的服务器超时问题。这就是为什么在本文中，我们将重点关注服务器端超时。

To get the basic terminology out of the way: timeout is a time interval (or limit) in which a specific action must complete. If the operation does not complete in the given time limit, a timeout occurs, and the operation is canceled.

摆脱基本术语：超时是特定操作必须完成的时间间隔（或限制）。如果操作未在给定的时限内完成，则会发生超时，并取消操作。

Initializing a `net/http` server in Golang reveals a few basic timeout configurations:

在 Golang 中初始化一个 `net/http` 服务器揭示了一些基本的超时配置：

```go
srv := &http.Server{
    ReadTimeout:       1 * time.Second,
    WriteTimeout:      1 * time.Second,
    IdleTimeout:       30 * time.Second,
    ReadHeaderTimeout: 2 * time.Second,
    TLSConfig:         tlsConfig,
    Handler:           srvMux,
}
```

This server of `http.Server` type can be initialized with four different timeouts:

这个 `http.Server` 类型的服务器可以用四种不同的超时时间初始化：

- `ReadTimeout`: the maximum duration for reading the entire request, including the body
- `WriteTimeout`: the maximum duration before timing out writes of the response
- `IdleTimetout`: the maximum amount of time to wait for the next request when keep-alive is enabled
- `ReadHeaderTimeout`: the amount of time allowed to read request headers

- `ReadTimeout`：读取整个请求的最大持续时间，包括正文
- `WriteTimeout`：响应写入超时前的最大持续时间
- `IdleTimetout`：启用keep-alive时等待下一个请求的最长时间
- `ReadHeaderTimeout`：允许读取请求头的时间量

A graphical representation of the above timeouts:

上述超时的图形表示：

![img](https://ieftimov.com/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/request-lifecycle-timeouts.png)

Before you start thinking that these are all the timeouts you need, tread carefully! There’s more than meets the eye here. These timeout values provide much more fine-grained control, and they are not going to help us to timeout our long-running HTTP handlers.

在您开始认为这些都是您需要的超时之前，请谨慎行事！这里有更多的东西。这些超时值提供了更细粒度的控制，它们不会帮助我们使长时间运行的 HTTP 处理程序超时。

Let me explain.

让我解释。

## Timeouts and deadlines

## 超时和截止日期

If we look at the source of `net/http`, in particular, the [`conn` type](https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L248), we will notice that it uses `net.Conn` connection under the hood which represents the underlying network connection:

如果我们查看`net/http`的源码，特别是[`conn`类型](https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L248)，我们会注意到它在引擎盖下使用了`net.Conn`连接，它代表了底层网络连接：

```go
// Taken from: https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L247
// A conn represents the server-side of an HTTP connection.
type conn struct {
    // server is the server on which the connection arrived.
    // Immutable;never nil.
    server *Server

    // * Snipped *

    // rwc is the underlying network connection.
    // This is never wrapped by other types and is the value given out
    // to CloseNotifier callers.It is usually of type *net.TCPConn or
    // *tls.Conn.
    rwc net.Conn

    // * Snipped *
}
```

In other words, it’s the actual TCP connection that our HTTP request travels on. From a type perspective, it’s a `*net.TCPConn` or `*tls.Conn` if it’s a TLS connection.

换句话说，它是我们的 HTTP 请求通过的实际 TCP 连接。从类型的角度来看，如果是 TLS 连接，则是 `*net.TCPConn` 或 `*tls.Conn`。

The `serve` [function](https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L1765), calls the `readRequest` function [for each incoming request]( https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L1822). `readRequest` [uses the timeout values](https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L946-L958) that we set on the server to **set deadlines on the TCP connection**:

`serve` [函数](https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L1765)，调用 `readRequest` 函数[对于每个传入的请求]( https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L1822)。 `readRequest` [使用超时值](https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L946-L958) 我们在服务器上设置为 **set TCP 连接的最后期限**：

```go
// Taken from: https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L936
// Read next request from connection.
func (c *conn) readRequest(ctx context.Context) (w *response, err error) {
        // *Snipped*

        t0 := time.Now()
        if d := c.server.readHeaderTimeout();d != 0 {
                hdrDeadline = t0.Add(d)
        }
        if d := c.server.ReadTimeout;d != 0 {
                wholeReqDeadline = t0.Add(d)
        }
        c.rwc.SetReadDeadline(hdrDeadline)
        if d := c.server.WriteTimeout;d != 0 {
                defer func() {
                        c.rwc.SetWriteDeadline(time.Now().Add(d))
                }()
        }

        // *Snipped*
}
```

From the snippet above, we can conclude that the timeout values set on the server end up as TCP connection deadlines instead of HTTP timeouts.

从上面的片段中，我们可以得出结论，服务器上设置的超时值最终是 TCP 连接截止时间，而不是 HTTP 超时。

So, what are deadlines then? How do they work? Will they timeout our connection if our request takes too long?

那么，什么是截止日期呢？它们是如何工作的？如果我们的请求时间过长，他们会超时我们的连接吗？

A simple way to think about deadlines is as a point in time at which restrictions on specific actions on the connection are enforced. For example, if we set a write deadline after the deadline time passes, any write actions on the connection will be forbidden.

考虑截止日期的一种简单方法是作为对连接上的特定操作实施限制的时间点。例如，如果我们在截止时间过后设置写入截止时间，则将禁止对连接进行任何写入操作。

While we can create timeout-like behavior using deadlines, we cannot control the time it takes for our handlers to complete. Deadlines operate on the connection, so our server will fail to return a result only after the handlers try to access connection properties (such as writing to `http.ResponseWriter`).

虽然我们可以使用截止日期创建类似超时的行为，但我们无法控制处理程序完成所需的时间。截止日期对连接进行操作，因此只有在处理程序尝试访问连接属性（例如写入 `http.ResponseWriter`）后，我们的服务器才会无法返回结果。

To see this in action, let’s create a tiny handler that takes longer to complete relative to the timeouts we set on the server:

为了看到这一点，让我们创建一个小处理程序，相对于我们在服务器上设置的超时时间，它需要更长的时间才能完成：

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func slowHandler(w http.ResponseWriter, req *http.Request) {
    time.Sleep(2 * time.Second)
    io.WriteString(w, "I am slow!\n")
}

func main() {
    srv := http.Server{
        Addr:         ":8888",
        WriteTimeout: 1 * time.Second,
        Handler:      http.HandlerFunc(slowHandler),
    }

    if err := srv.ListenAndServe();err != nil {
        fmt.Printf("Server failed: %s\n", err)
    }
}
```

The server above has a single handler, which takes 2 seconds to complete. On the other hand, the `http.Server` has a `WriteTimeout` set to 1 second. Due to the configuration of the server, we expect the handler to be unable to write the response to the connection.

上面的服务器有一个处理程序，需要 2 秒才能完成。另一方面，`http.Server` 的 `WriteTimeout` 设置为 1 秒。由于服务器的配置，我们预计处理程序无法将响应写入连接。

We can start the server using `go run server.go`. To send a request we can `curl localhost:8888`:

我们可以使用 `go run server.go` 来启动服务器。要发送请求，我们可以`curl localhost:8888`：

```shell
$ time curl localhost:8888
curl: (52) Empty reply from server
curl localhost:8888  0.01s user 0.01s system 0% cpu 2.021 total
```

The request took 2 seconds to complete, and the response from the server was empty. While our server knows that we cannot write to our response after 1 second, the handler still took 100% more (2 seconds) to complete.

请求用了 2 秒才完成，来自服务器的响应为空。虽然我们的服务器知道我们无法在 1 秒后写入我们的响应，但处理程序仍然需要 100%（2 秒）才能完成。

While this is a timeout-like behavior, it would be more useful to stop our server from further execution when it reaches the timeout, ending the request. In our example above, the handler proceeds to process the request until it completes, although it takes 100% longer (2 seconds) than the response write timeout time (1 second).

虽然这是一种类似超时的行为，但在达到超时时停止我们的服务器进一步执行会更有用，从而结束请求。在我们上面的示例中，处理程序继续处理请求直到它完成，尽管它比响应写入超时时间（1 秒）花费 100% 的时间（2 秒）。

The natural question is, how can we have efficient timeouts for our handlers?

自然的问题是，我们如何为处理程序设置有效的超时时间？

## Handler timeout

## 处理程序超时

Our goal is to make sure our `slowHandler` does not take longer than 1 seconds to complete. If it does take longer, our server should stop its execution and return a proper timeout error.

我们的目标是确保我们的 `slowHandler` 完成的时间不超过 1 秒。如果确实需要更长的时间，我们的服务器应该停止执行并返回正确的超时错误。

In Go, as with other programming languages, composition is very often the favored approach to writing and designing code. The [`net/http` package](https://golang.org/pkg/net/http) of the standard library is one of the places where having compatible components that one can put together with little effort is an obvious design choice .

在 Go 中，与其他编程语言一样，组合通常是编写和设计代码的首选方法。标准库的 [`net/http` 包](https://golang.org/pkg/net/http) 是其中一个地方之一，具有可以轻松组装的兼容组件是一种明显的设计选择.

In that light, the `net/http` packages provide [a `TimeoutHandler`](https://golang.org/pkg/net/http/#TimeoutHandler) - it returns a handler that runs a handler within the given time limit .

有鉴于此，`net/http` 包提供了 [a `TimeoutHandler`](https://golang.org/pkg/net/http/#TimeoutHandler) - 它返回一个在给定时间限制内运行处理程序的处理程序.

Its signature is:

它的签名是：

```go
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```

It takes a `Handler` as the first argument, a `time.Duration` as the second argument (the timeout time) and a `string`, the message returned when it hits the timeout.

它接受一个 `Handler` 作为第一个参数，一个 `time.Duration` 作为第二个参数（超时时间）和一个 `string`，当它达到超时时返回的消息。

To wrap our `slowHandler` within a `TimeoutHandler`, all we have to do is:

要将我们的 `slowHandler` 包装在 `TimeoutHandler` 中，我们所要做的就是：

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func slowHandler(w http.ResponseWriter, req *http.Request) {
    time.Sleep(2 * time.Second)
    io.WriteString(w, "I am slow!\n")
}

func main() {
    srv := http.Server{
        Addr:         ":8888",
        WriteTimeout: 5 * time.Second,
        Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandler), 1*time.Second, "Timeout!\n"),
    }

    if err := srv.ListenAndServe();err != nil {
        fmt.Printf("Server failed: %s\n", err)
    }
}
```

The two notable changes are:

两个显着的变化是：

- We wrap our `slowHanlder` in the `http.TimetoutHandler`, setting the timeout to 1 second and the timeout message to “Timeout!”. 

- 我们将 `slowHanlder` 包装在 `http.TimetoutHandler` 中，将超时设置为 1 秒，并将超时消息设置为“超时！”。

- We increase the `WriteTimeout` to 5 seconds, to give the `http.TimeoutHandler` time to kick in. If we don't do this when the `TimeoutHandler` kicks in, the deadline will pass, and it won't be able to write to the response.

- 我们将 `WriteTimeout` 增加到 5 秒，让 `http.TimeoutHandler` 有时间启动。如果我们在 `TimeoutHandler` 启动时不这样做，deadline 就会过去，它不会能够写入响应。

If we started the server again and hit the slow handler we’ll get the following output:

如果我们再次启动服务器并点击慢处理程序，我们将得到以下输出：

```shell
$ time curl localhost:8888
Timeout!
curl localhost:8888  0.01s user 0.01s system 1% cpu 1.023 total
```

After a second, our `TimeoutHandler` will kick in, stop processing the `slowHandler`, and return a plain “`Timeout!`” message. If the message we set is blank, then the handler will return the default timeout response, which is:

一秒钟后，我们的 `TimeoutHandler` 将启动，停止处理 `slowHandler`，并返回一个简单的“`Timeout!`”消息。如果我们设置的消息为空，则处理程序将返回默认超时响应，即：

```html
<html>
  <head>
    <title>Timeout</title>
  </head>
  <body>
   <h1>Timeout</h1>
  </body>
</html>
```

Regardless of the output, this is pretty neat, isn’t it? Our application now is protected from long-running handlers and specially crafted requests made to cause very long-running handlers, leading to a potential DoS (Denial of Service) attack.

不管输出如何，这都非常整洁，不是吗？我们的应用程序现在受到保护，免受长时间运行的处理程序和特制请求的影响，这些请求会导致非常长时间运行的处理程序，从而导致潜在的 DoS（拒绝服务）攻击。

It’s worth noting that although setting timeouts is a great start, it’s still elementary protection. If you’re under threat of an imminent DoS attack, you should look into more advanced protection tools and techniques. ([Cloudflare](https://www.cloudflare.com/ddos/) is a good start.)

值得注意的是，虽然设置超时是一个很好的开始，但它仍然是基本的保护。如果您面临即将到来的 DoS 攻击的威胁，您应该研究更高级的保护工具和技术。 （[Cloudflare](https://www.cloudflare.com/ddos/)是一个好的开始。)

Our `slowHandler` is a simple example-only handler. But, what if our application was much more complicated, making external calls to other services or resources? What if we had an outgoing request to a service such as S3 when our handler would time out?

我们的 `slowHandler` 是一个简单的示例处理程序。但是，如果我们的应用程序要复杂得多，对其他服务或资源进行外部调用呢？如果当我们的处理程序超时时我们有一个对服务（例如 S3）的传出请求怎么办？

What would happen then?

那会发生什么？

## Unhandled timeouts and request cancellations

## 未处理的超时和请求取消

Let’s expand our example a bit:

让我们稍微扩展一下我们的例子：

```go
func slowAPICall() string {
    d := rand.Intn(5)
    select {
    case <-time.After(time.Duration(d) * time.Second):
        log.Printf("Slow API call done after %s seconds.\n", d)
        return "foobar"
    }
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
    result := slowAPICall()
    io.WriteString(w, result+"\n")
}
```

Let’s imagine that initially we didn’t know that our `slowHandler` took so long to complete because it was sending a request to an API - using the `slowAPICall` function.

让我们想象一下，最初我们不知道我们的 `slowHandler` 需要这么长时间才能完成，因为它正在向 API 发送请求 - 使用 `slowAPICall` 函数。

The `slowAPICall` function is straightforward: using `select` and a `time.After` it blocks between 0 and 5 seconds. Once that period passes, the `time.After` method sends a value through its channel and `"foobar"` will be returned.

`slowAPICall` 函数很简单：使用 `select` 和 `time.After` 它会在 0 到 5 秒之间阻塞。一旦该时间段过去，`time.After` 方法将通过其通道发送一个值，并且将返回 `"foobar"`。

(An alternative approach is to use `sleep(time.Duration(rand.Intn(5)) * time.Second)`, but we will stick to `select` because it will make our life simpler in the next example.)

（另一种方法是使用 `sleep(time.Duration(rand.Intn(5)) * time.Second)`，但我们将坚持使用 `select`，因为它会让我们在下一个例子中的生活更简单。）

If we run our server, we would expect the timeout handler to cut off the request processing after 1 second. Sending a request proves that:

如果我们运行我们的服务器，我们希望超时处理程序在 1 秒后切断请求处理。发送请求证明：

```shell
$ time curl localhost:8888
Timeout!
curl localhost:8888  0.01s user 0.01s system 1% cpu 1.021 total
```

By looking at the server output, we will notice that it prints the loglines after a few seconds instead of when the timeout handler kicks in:

通过查看服务器输出，我们会注意到它在几秒钟后而不是在超时处理程序启动时打印日志行：

```shell
$ go run server.go
2019/12/29 17:20:03 Slow API call done after 4 seconds.
```

Such behavior suggests that although the request timed out after 1 second, the server proceeded to process the request fully. That’s why it printed the logline after 4 seconds passed.

这种行为表明，虽然请求在 1 秒后超时，但服务器继续完全处理请求。这就是为什么它在 4 秒后打印日志行。

While this example is trivial and naive, such behavior in production servers can become a rather big problem. For example, imagine if the `slowAPICall` function spawned hundreds of goroutines, each of them processing some data. Or if it was issuing multiple API calls to various systems. Such long-running processes will eat up resources from your system, while the caller/client won’t ever use their result.

虽然这个例子是微不足道的，但生产服务器中的这种行为可能会成为一个相当大的问题。例如，假设 `slowAPICall` 函数产生了数百个 goroutine，每个 goroutine 都处理一些数据。或者，如果它向各种系统发出多个 API 调用。这种长时间运行的进程会消耗你系统的资源，而调用者/客户端永远不会使用他们的结果。

So, how can we guard our system from such unoptimized timeouts or request cancellations?

那么，我们如何保护我们的系统免受此类未优化的超时或请求取消的影响？

## Context timeouts and cancellation

## 上下文超时和取消

Go comes with a neat package for handling such scenarios called [`context`](https://golang.org/pkg/context/).

Go 带有一个用于处理此类场景的简洁包，称为 [`context`](https://golang.org/pkg/context/)。

The `context` package was promoted to the standard library as of Go 1.7. Previously it was part of the [Go Sub-repository Packages](https://godoc.org/-/subrepo), with the name [`golang.org/x/net/context`](https://godoc.org/golang.org/x/net/context) 

从 Go 1.7 开始，`context` 包被提升为标准库。以前它是 [Go Sub-repository Packages](https://godoc.org/-/subrepo) 的一部分，名称为 [`golang.org/x/net/context`](https://godoc.org/golang.org/x/net/context)

The package defines the `Context` type. It’s primary purpose is to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes. If you would like to learn more about the context package, I recommend reading “Go Concurrency Patterns: Context” on [Golang’s blog](https://blog.golang.org/context).

该包定义了 `Context` 类型。它的主要目的是跨 API 边界和进程之间传输截止日期、取消信号和其他请求范围的值。如果您想了解有关上下文包的更多信息，我建议您阅读 [Golang 的博客](https://blog.golang.org/context) 上的“Go 并发模式：上下文”。

The `Request` type that is part of the `net/http` package already has a `context` attached to it. As of Go 1.7, `Request` has [a `Context` function](https://golang.org/pkg/net/http/#Request.Context), which returns the request’s context. For incoming server requests, the server cancels the context when the client’s connection closes, when the request is canceled (in HTTP/2), or when the `ServeHTTP` method returns.

作为`net/http` 包一部分的`Request` 类型已经附加了一个`context`。从 Go 1.7 开始，`Request` 有 [a `Context` 函数](https://golang.org/pkg/net/http/#Request.Context)，它返回请求的上下文。对于传入的服务器请求，当客户端的连接关闭、请求被取消（在 HTTP/2 中)或当 `ServeHTTP` 方法返回时，服务器会取消上下文。

The behavior we are looking for is to stop all further processing on the server-side when the client cancels the request (we hit `CTRL + C` on our `cURL`) or the `TimeoutHandler` steps in after some time and ends the request. That will effectively close all connections and free all other resources taken up by the running handler (and all of its children goroutines).

我们正在寻找的行为是当客户端取消请求时停止服务器端的所有进一步处理（我们在 `cURL` 上点击 `CTRL + C`）或一段时间后 `TimeoutHandler` 介入并结束要求。这将有效地关闭所有连接并释放正在运行的处理程序（及其所有子 goroutine）占用的所有其他资源。

Let’s use the request Context to pass it to the `slowAPICall` function as an argument:

让我们使用请求上下文将其作为参数传递给 `slowAPICall` 函数：

```go
func slowAPICall(ctx context.Context) string {
    d := rand.Intn(5)
    select {
    case <-time.After(time.Duration(d) * time.Second):
        log.Printf("Slow API call done after %d seconds.\n", d)
        return "foobar"
    }
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
    result := slowAPICall(r.Context())
    io.WriteString(w, result+"\n")
}
```

Now that we utilize the request context, how can we put it in action? [The `Context` type](https://golang.org/pkg/context/#Context) has a `Done` attribute, which is of type `<-chan struct{}`. `Done` closes when the work done on behalf of the context should be canceled, which is what we need in the example.

既然我们利用了请求上下文，我们如何将其付诸实践？ [`Context` 类型](https://golang.org/pkg/context/#Context) 有一个 `Done` 属性，它的类型是 `<-chan struct{}`。当代表上下文完成的工作应该被取消时，`Done` 关闭，这是我们在示例中所需要的。

Let’s handle the `ctx.Done` channel in the `select` block in the `slowAPICall` function. When we receive an empty `struct` via the `Done` channel, this signifies the context cancellation, and we have to return a zero-value string from the `slowAPICall` function:

让我们在 `slowAPICall` 函数的 `select` 块中处理 `ctx.Done` 通道。当我们通过 `Done` 通道收到一个空的 `struct` 时，这表示上下文取消，我们必须从 `slowAPICall` 函数返回一个零值字符串：

```go
func slowAPICall(ctx context.Context) string {
    d := rand.Intn(5)
    select {
    case <-ctx.Done():
        log.Printf("slowAPICall was supposed to take %s seconds, but was canceled.", d)
        return ""
    case <-time.After(time.Duration(d) * time.Second):
        log.Printf("Slow API call done after %d seconds.\n", d)
        return "foobar"
    }
}
```

(This is the reason we used a `select` block, instead of the `time.Sleep` - we can just handle the `Done` channel in the `select` here.)

（这就是我们使用 `select` 块而不是 `time.Sleep` 的原因 - 我们可以在这里处理 `select` 中的 `Done` 通道。）

In our limited example, this does the trick – when we receive value through the `Done` channel, we log a line to STDOUT and return an empty string. In more complicated situations, such as sending real API requests, you might need to close down connections or clean up file descriptors.

在我们有限的示例中，这很有效——当我们通过“完成”通道接收到值时，我们将一行记录到 STDOUT 并返回一个空字符串。在更复杂的情况下，例如发送真实的 API 请求，您可能需要关闭连接或清理文件描述符。

Let’s spin up the server again and send a `cURL` request:

让我们再次启动服务器并发送一个 `cURL` 请求：

```shell
# The cURL command:
$ curl localhost:8888
Timeout!

# The server output:
$ go run server.go
2019/12/30 00:07:15 slowAPICall was supposed to take 2 seconds, but was canceled.
```

So check this out: we ran a `cURL` to the server, it took longer than 1 second, and our server canceled the `slowAPICall` function. And we didn’t need to write almost any code for it. The `TimeoutHandler` did this for us - when the handler took longer than expected, the `TimeoutHandler` stopped the execution of the handler and canceled the request context.

所以检查一下：我们向服务器运行了一个 `cURL`，花费了超过 1 秒的时间，我们的服务器取消了 `slowAPICall` 函数。我们几乎不需要为它编写任何代码。 `TimeoutHandler` 为我们做了这件事——当处理程序花费的时间比预期的要长时，`TimeoutHandler` 停止了处理程序的执行并取消了请求上下文。

The `TimeoutHandler` performs the context cancellation in [the `timeoutHandler.ServeHTTP` method](https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L3217-L3263):

`TimeoutHandler` 在 [`timeoutHandler.ServeHTTP` 方法](https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L3217-L3263) 中执行上下文取消：

```go
// Taken from: https://github.com/golang/go/blob/bbbc658/src/net/http/server.go#L3217-L3263
func (h *timeoutHandler) ServeHTTP(w ResponseWriter, r *Request) {
        ctx := h.testContext
        if ctx == nil {
            var cancelCtx context.CancelFunc
            ctx, cancelCtx = context.WithTimeout(r.Context(), h.dt)
            defer cancelCtx()
        }
        r = r.WithContext(ctx)

        // *Snipped*
}
```

Above, we use the request context by invoking `context.WithTimeout` on it. The timeout value `h.dt`, which is the second argument received by the `TimeoutHandler`, is applied to the context. The returned context is a copy of the request context with a timeout attached. Right after, it’s set as the request’s context using the `r.WithContext(ctx)` invocation.

上面，我们通过调用 `context.WithTimeout` 来使用请求上下文。超时值“h.dt”是“TimeoutHandler”接收的第二个参数，应用于上下文。返回的上下文是附加了超时的请求上下文的副本。紧接着，使用 `r.WithContext(ctx)` 调用将其设置为请求的上下文。

The `context.WithTimeout` function makes the context cancellation. It returns a copy of the `Context` with a timeout set to the duration passed as an argument. Once it reaches the timeout, it cancels the context.

`context.WithTimeout` 函数使上下文取消。它返回一个 `Context` 的副本，超时设置为作为参数传递的持续时间。一旦达到超时，它就会取消上下文。

Here’s the code that does it:

这是执行此操作的代码：

```go
// Taken from: https://github.com/golang/go/blob/bbbc6589/src/context/context.go#L486-L498
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))
}

// Taken from: https://github.com/golang/go/blob/bbbc6589/src/context/context.go#L418-L450
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
        // *Snipped*

        c := &timerCtx{
            cancelCtx: newCancelCtx(parent),
            deadline:  d,
        }

        // *Snipped*

        if c.err == nil {
            c.timer = time.AfterFunc(dur, func() {
                c.cancel(true, DeadlineExceeded)
            })
        }
        return c, func() { c.cancel(true, Canceled) }
}
```

Here we meet deadlines again. The `WithDeadline` function sets a function that executes after the duration `d` passes. Once the duration passes, it invokes the `cancel` method on the context, which will close the `done` channel of the context and will set the context’s `timer` attribute to `nil`.

在这里，我们再次满足最后期限。 `WithDeadline` 函数设置了一个在持续时间 `d` 过去后执行的函数。一旦持续时间过去，它会调用上下文上的 `cancel` 方法，这将关闭上下文的 `done` 通道并将上下文的 `timer` 属性设置为 `nil`。

The closing of the `Done` channel effectively cancels the context, which allows our `slowAPICall` function to stop its execution. That’s how the `TimeoutHandler` timeouts long-running handlers.

`Done` 通道的关闭有效地取消了上下文，这允许我们的 `slowAPICall` 函数停止其执行。这就是 `TimeoutHandler` 如何使长时间运行的处理程序超时。

(If you would like to read the code that does the above, you should see the `cancel` functions on [the `cancelCtx` type](https://github.com/golang/go/blob/bbbc6589dfbc05be2bfa59f51c20f9eaa8d0c531/src/context/context.go#L389-L416) and [the `timerCtx` type](https://github.com/golang/go/blob/bbbc6589dfbc05be2bfa59f51c20f9eaa8d0c531/src/context/context.go#L472-L484).)

（如果您想阅读执行上述操作的代码，您应该在 [`cancelCtx` 类型](https://github.com/golang/go/blob/bbbc6589dfbc05be2bfa59f51c20f9eaa8d0c531/src/context/context.go#L389-L416) 和 [`timerCtx` 类型](https://github.com/golang/go/blob/bbbc6589dfbc05be2bfa59f51c20f9eaa8d0c531/src/context/context.go#L472-L482)。

## Resilient `net/http` servers

## 弹性 `net/http` 服务器

Connection deadlines provide low-level fine-grained control. While their names contain “timeout” they do not have the behavior that folks commonly expect from timeouts. They are in fact very powerful, but they expect the programmer to know how to wield that weapon.

连接期限提供了低级别的细粒度控制。虽然他们的名字包含“超时”，但他们没有人们通常期望的超时行为。他们实际上非常强大，但他们希望程序员知道如何使用这种武器。

On the other hand, when working with HTTP handlers, the `TimeoutHandler` should be our go-to tool. The design chosen by the Go authors, of having composable handlers, provides flexibility, so much that we could even have different timeouts per handler if we decided to. `TimeoutHandler` provides execution control while maintaining the behavior that we commonly expect when thinking of timeouts.

另一方面，当使用 HTTP 处理程序时，`TimeoutHandler` 应该是我们的首选工具。 Go 作者选择的具有可组合处理程序的设计提供了灵活性，如果我们决定，我们甚至可以为每个处理程序设置不同的超时。 `TimeoutHandler` 提供执行控制，同时保持我们在考虑超时时通常期望的行为。

On top of that, the `TimeoutHandler` works well with the `context` package. While the `context` package is simple, it carries cancellation signals and request-scoped data, that we can use to make our applications adhere better to the intricacies of networks.

最重要的是，`TimeoutHandler` 可以很好地与`context` 包配合使用。虽然 `context` 包很简单，但它携带取消信号和请求范围的数据，我们可以使用它们使我们的应用程序更好地适应网络的复杂性。

Before we close, here are three suggestions on how to think of timeouts while writing your HTTP servers:

在我们结束之前，这里有关于如何在编写 HTTP 服务器时考虑超时的三个建议：

1. Most-often, reach for `TimeoutHandler`. It does what we commonly expect of timeouts.
2. Never forget context cancellations. The `context` package is simple to use and can save your servers lots of processing resources. Especially againts bad actors or misbehaving networks.
3. By all means, use deadlines. Just make sure you test thoroughly that they provide you the functionality that you want.

1. 最常见的是，使用`TimeoutHandler`。它执行我们通常对超时的期望。
2. 永远不要忘记上下文取消。 `context` 包使用简单，可以为您的服务器节省大量处理资源。尤其是针对不良行为者或行为不端的网络。
3. 无论如何，使用最后期限。只要确保您彻底测试它们是否为您提供了您想要的功能。

To read more on the topic:

要阅读有关该主题的更多信息：

- “The complete guide to Go net/http timeouts” on [Cloudflare’s blog](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/)
- “So you want to expose Go on the Internet” on [Cloudflare’s blog](https://blog.cloudflare.com/exposing-go-on-the-internet/)
- “Use http.TimeoutHandler or ReadTimeout/WriteTimeout?” on [Stackoverflow](https://stackoverflow.com/questions/51258952/use-http-timeouthandler-or-readtimeout-writetimeout) 

- [Cloudflare 的博客](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/) 上的“Go net/http 超时的完整指南”
- [Cloudflare 的博客](https://blog.cloudflare.com/exposing-go-on-the-internet/) 上的“所以你想在互联网上公开 Go”
- “使用 http.TimeoutHandler 还是 ReadTimeout/WriteTimeout？”在 [Stackoverflow](https://stackoverflow.com/questions/51258952/use-http-timeouthandler-or-readtimeout-writetimeout)

- “Standard net/http config will break your production environment” on [Simon Frey's blog](https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production)

- [Simon Frey 的博客](https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-) 上的“标准网络/http 配置将破坏您的生产环境”打破你的生产)

**Liked this article?** Subscribe to my newsletter and get future articles in your inbox. It's a short and sweet read, going out to +1000 of other engineers.

**喜欢这篇文章吗？** 订阅我的时事通讯并在您的收件箱中获取未来的文章。这是一个简短而甜蜜的读物，与其他 1000 多名工程师分享。

I care about your privacy, and will never send you spam. You can unsubscribe at any time. 

我关心您的隐私，绝不会向您发送垃圾邮件。您可以随时取消订阅。

