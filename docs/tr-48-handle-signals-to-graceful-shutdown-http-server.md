# How to handle signals with Go to graceful shutdown HTTP server

# 如何使用 Go to graceful shutdown HTTP server 处理信号

Written on June 18, 2020  ​From: https://rafallorenz.com/go/handle-signals-to-graceful-shutdown-http-server/

In this article we are going to learn how to handle os incoming [signals](https://golang.org/pkg/os/signal) for performing graceful shutdown of http server. To do so we are going to take advantage of [os/signal](https://golang.org/pkg/os/signal) package.

在本文中，我们将学习如何处理 os 传入 [信号](https://golang.org/pkg/os/signal)以正常关闭 http 服务器。为此，我们将利用 [os/signal](https://golang.org/pkg/os/signal) 包。

> Signals are primarily used on Unix-like systems.

> 信号主要用于类 Unix 系统。

# Types of signals

# 信号类型

We are going to focus on asynchronous signals. They are not triggered by program errors, but are instead sent from the kernel or from some  other program.

我们将专注于异步信号。它们不是由程序错误触发的，而是从内核或其他程序发送的。

Of the asynchronous signals:
- the `SIGHUP` signal is sent when a program loses its controlling terminal
- the `SIGINT` signal is sent when the user at the controlling terminal presses the interrupt character, which by default is **^C (Control-C)**
- The `SIGQUIT` signal is sent when the user at the controlling terminal presses the quit character, which by default is **^\ (Control-Backslash)**

异步信号：
- 当程序失去其控制终端时发送`SIGHUP`信号
- 当控制端的用户按下中断字符时发送`SIGINT`信号，默认为**^C (Control-C)**
- 当控制终端的用户按下退出字符时发送`SIGQUIT`信号，默认为**^\（Control-Backslash）**

In general you can cause a program to simply exit by pressing ^C, and you can cause it to exit with a stack dump by pressing ^.

通常，您可以通过按 ^C 使程序简单地退出，也可以通过按 ^ 使程序退出堆栈转储。

# Default behavior of signals in Go programs

# Go 程序中信号的默认行为

> By default, a synchronous signal is converted into a run-time panic. A `SIGHUP`, `SIGINT`, or `SIGTERM` signal causes the program to exit. If the Go program is started with either SIGHUP or SIGINT ignored  (signal handler set to SIG_IGN), they will remain ignored. If the Go program is started with a non-empty signal mask, that will  generally be honored. However, some signals are explicitly unblocked:  the synchronous signals.

> 默认情况下，同步信号会转换为运行时恐慌。 `SIGHUP`、`SIGINT` 或 `SIGTERM` 信号会导致程序退出。如果 Go 程序启动时忽略了 SIGHUP 或 SIGINT（信号处理程序设置为 SIG_IGN），它们将保持被忽略。如果 Go 程序以非空信号掩码启动，则通常会受到尊重。但是，有些信号是明确解除阻塞的：同步信号。

You can read more about it on the [package documentation](https://golang.org/pkg/os/signal/#hdr-Default_behavior_of_signals_in_Go_programs)

您可以在 [包文档](https://golang.org/pkg/os/signal/#hdr-Default_behavior_of_signals_in_Go_programs) 上阅读有关它的更多信息

# Handling signals

# 处理信号

The idea is to catch incoming signal and perform graceful stop of our http application. We can create signal channel using, and use it to  notify on incoming signal. `signal.Notify` disables the default behavior for a given set of asynchronous signals  and instead delivers them over one or more registered channels.

这个想法是捕获传入的信号并优雅地停止我们的 http 应用程序。我们可以使用创建信号通道，并使用它来通知传入的信号。 `signal.Notify` 禁用一组给定异步信号的默认行为，而是通过一个或多个注册通道传递它们。

```
   signalChan := make(chan os.Signal, 1)

    signal.Notify(
        signalChan,
        syscall.SIGHUP,  // kill -SIGHUP XXXX
        syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
        syscall.SIGQUIT, // kill -SIGQUIT XXXX
    )

    <-signalChan
    log.Print("os.Interrupt - shutting down...\n")

    // terminate after second signal before callback is done
    go func() {
        <-signalChan
        log.Fatal("os.Kill - terminating...\n")
    }()

    // PERFORM GRACEFUL SHUTDOWN HERE

    os.Exit(0)
```


When signal is received we will call a callback followed after by `os.Exit(0)`, if second signal is received will terminate process by a call to `os.Exit(1)`.

当接收到信号时，我们将调用一个回调，然后是 `os.Exit(0)`，如果接收到第二个信号，将通过调用 `os.Exit(1)` 来终止进程。

# Server

#  服务器

## Graceful shutdown

## 优雅关机

To gracefully shutdown [http.Server](https://golang.org/pkg/net/http/#Server)we can use [Shutdown](https://golang.org/pkg/net/http/#Server. Shutdown) method.

要正常关闭 [http.Server](https://golang.org/pkg/net/http/#Server)，我们可以使用 Shutdown 方法。

> Shutdown gracefully shuts down the server without interrupting any  active connections. Shutdown works by first closing all open listeners,  then closing all idle connections, and then waiting indefinitely for  connections to return to idle and then shut down. If the provided  context expires before the shutdown is complete, Shutdown returns the  context’s error, otherwise it returns any error returned from closing  the Server’s underlying Listener(s).

> 关闭可以正常关闭服务器，而不会中断任何活动连接。关闭的工作原理是首先关闭所有打开的侦听器，然后关闭所有空闲连接，然后无限期地等待连接返回空闲状态，然后关闭。如果提供的上下文在关闭完成之前到期，则 Shutdown 返回上下文的错误，否则返回关闭服务器底层侦听器返回的任何错误。

```
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  if err := httpServer.Shutdown(ctx);err != nil {
    log.Fatalf("shutdown error: %v\n", err)
  } else {
    log.Printf("gracefully stopped\n")
  }
```


> Shutdown does not attempt to close nor wait for hijacked  connections such as WebSockets. The caller of Shutdown should separately notify such long-lived connections of shutdown and wait for them to  close, if desired. See [RegisterOnShutdown](https://golang.org/pkg/net/http/#Server.RegisterOnShutdown) for a way to register shutdown notification functions.

> 关机不会尝试关闭或等待被劫持的连接，例如 WebSockets。如果需要，Shutdown 的调用者应该单独通知这些长期存在的连接关闭并等待它们关闭。有关注册关机通知功能的方法，请参阅 [RegisterOnShutdown](https://golang.org/pkg/net/http/#Server.RegisterOnShutdown)。

## Handling hijacked connections 

## 处理劫持的连接

> [RegisterOnShutdown](https://golang.org/pkg/net/http/#Server.RegisterOnShutdown) registers a function to call on Shutdown. This can be used to  gracefully shutdown connections that have undergone ALPN protocol  upgrade or that have been hijacked. This function should start  protocol-specific graceful shutdown, but should not wait for shutdown to complete.

> [RegisterOnShutdown](https://golang.org/pkg/net/http/#Server.RegisterOnShutdown) 注册一个函数以在关闭时调用。这可用于正常关闭经过 ALPN 协议升级或被劫持的连接。此函数应启动特定于协议的正常关闭，但不应等待关闭完成。

```
    ctx, cancel := context.WithCancel(context.Background())

    httpServer.RegisterOnShutdown(cancel)
```


We want to ask context passed down to the socket handlers/goroutines to stop. To do so we can set `BaseContext` property on [http.Server](https://golang.org/pkg/net/http/#Server).

我们想让传递给套接字处理程序/goroutines 的上下文停止。为此，我们可以在 [http.Server](https://golang.org/pkg/net/http/#Server) 上设置 `BaseContext` 属性。

> BaseContext optionally specifies a function that returns the base  context for incoming requests on this server. The provided Listener is  the specific Listener that’s about to start accepting requests. If  BaseContext is nil, the default is context.Background(). If non-nil, it  must return a non-nil context.

> BaseContext 可选地指定一个函数，该函数返回此服务器上传入请求的基本上下文。提供的 Listener 是即将开始接受请求的特定 Listener。如果 BaseContext 为 nil，则默认值为 context.Background()。如果非零，它必须返回一个非零上下文。

```
    ctx, cancel := context.WithCancel(context.Background())

    httpServer := &http.Server{
        Addr:        ":8080",
        Handler:     mux,
        BaseContext: func(_ net.Listener) context.Context { return ctx },
    }
    httpServer.RegisterOnShutdown(cancel)
```


Keep in mind that doing so will make [Shutdown](https://golang.org/pkg/net/http/#Server.Shutdown) cancel context via `RegisterOnShutdown` which will terminate all handlers using `BaseContext` immediately.

请记住，这样做将使 [Shutdown](https://golang.org/pkg/net/http/#Server.Shutdown) 通过 `RegisterOnShutdown` 取消上下文，这将立即终止所有使用 `BaseContext` 的处理程序。

Full correct solution would require you to separate base context for  WeSocket connections and other HTTP handlers. Also seems like  introducing simple timeout to cancel `BaseContext` will not be enough, connections idleness has to be checked as well.

完全正确的解决方案需要您为 WeSocket 连接和其他 HTTP 处理程序分离基本上下文。似乎引入简单的超时来取消 BaseContext 是不够的，还必须检查连接空闲。

If you do not care about notifying long-lived connections during  shutdown and don’t want to wait for them to close gracefully. Quick  solution would be to manually cancel `BaseContext` instead of using `RegisterOnShutdown` which makes context get canceled as the first procedure during shutdown.

如果您不关心在关闭期间通知长期连接并且不想等待它们正常关闭。快速解决方案是手动取消“BaseContext”而不是使用“RegisterOnShutdown”，这使得上下文在关闭期间作为第一个过程被取消。

```
    ctx, cancel := context.WithCancel(context.Background())

    httpServer := &http.Server{
        Addr:        ":8080",
        Handler:     mux,
        BaseContext: func(_ net.Listener) context.Context { return ctx },
    }
    
    // GRACEFULLY SHUTDOWN

    cancel()
```


# Gluing up all pieces together

# 将所有碎片粘合在一起

```
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Hello!")
    })

    httpServer := &http.Server{
        Addr:        ":8080",
        Handler:     mux,
        BaseContext: func(_ net.Listener) context.Context { return ctx },
    }
    // if your BaseContext is more complex you might want to use this instead of doing it manually
    // httpServer.RegisterOnShutdown(cancel)

    // Run server
    go func() {
        if err := httpServer.ListenAndServe();err != http.ErrServerClosed {
            // it is fine to use Fatal here because it is not main gorutine
            log.Fatalf("HTTP server ListenAndServe: %v", err)
        }
    }()

    signalChan := make(chan os.Signal, 1)

    signal.Notify(
        signalChan,
        syscall.SIGHUP,  // kill -SIGHUP XXXX
        syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
        syscall.SIGQUIT, // kill -SIGQUIT XXXX
    )

    <-signalChan
    log.Print("os.Interrupt - shutting down...\n")

    go func() {
        <-signalChan
        log.Fatal("os.Kill - terminating...\n")
    }()

    gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelShutdown()

    if err := httpServer.Shutdown(gracefullCtx);err != nil {
        log.Printf("shutdown error: %v\n", err)
        defer os.Exit(1)
        return
    } else {
        log.Printf("gracefully stopped\n")
    }

    // manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
    cancel()

    defer os.Exit(0)
    return
}
```


When Shutdown is called, `Serve`, `ListenAndServe`, and `ListenAndServeTLS` immediately return `ErrServerClosed`. Make sure the program doesn’t exit and waits instead for Shutdown to return.

当 Shutdown 被调用时，`Serve`、`ListenAndServe` 和 `ListenAndServeTLS` 立即返回 `ErrServerClosed`。确保程序不会退出并等待 Shutdown 返回。

**Note**: that we are deferring `os.Exit()` followed by `return`. Defers don’t run on `Fatal()`.

**注意**：我们推迟了 `os.Exit()` 后跟 `return`。延迟不会在“Fatal()”上运行。

> Calling Goexit from the main goroutine terminates that goroutine without func main returning. Since func main has not returned, the  program continues execution of other goroutines. If all other goroutines exit, the program crashes. 

> 从主协程调用 Goexit 会终止该协程，而 func main 不返回。由于 func main 没有返回，程序继续执行其他 goroutine。如果所有其他 goroutine 退出，程序就会崩溃。

You can read conversation about it [here](https://www.reddit.com/r/golang/comments/hbalgf/how_to_handle_signals_with_go_to_graceful/fv800rj?utm_source=share&utm_medium=web2x)

Run this example on [The Go Playground](https://play.golang.org/p/6w3yFxU1N24)

在 [The Go Playground](https://play.golang.org/p/6w3yFxU1N24) 上运行此示例

# Conclusion

#  结论

By using tools provided by go environment, we can easily handle  graceful shutdown of our application. We could simply create a package  to handle that for us, in fact I already have created one if you are  interested. [shutdown](https://github.com/vardius/shutdown) - is a simple go signals handler for performing graceful shutdown by executing callback function. 

通过使用 go 环境提供的工具，我们可以轻松处理应用程序的正常关闭。我们可以简单地创建一个包来为我们处理这个问题，事实上，如果您有兴趣，我已经创建了一个。 [shutdown](https://github.com/vardius/shutdown) - 是一个简单的 go 信号处理程序，用于通过执行回调函数执行正常关闭。


