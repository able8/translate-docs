# Graceful Shutdowns in Golang with signal.NotifyContext

# 在 Golang 中使用 signal.NotifyContext 优雅关闭

February 25, 2021 https://millhouse.dev/posts/graceful-shutdowns-in-golang-with-signal-notify-context

Graceful shutdowns are an important part of any application, especially if that  application modifies state. Before you “pull the plug” you should be  responding to those HTTP requests, finishing off database interactions  and closing off anything that might be left otherwise hanging or  orphaned.

正常关闭是任何应用程序的重要组成部分，尤其是当该应用程序修改状态时。在你“拔掉插头”之前，你应该响应那些 HTTP 请求，完成数据库交互并关闭任何可能挂起或孤立的东西。

With the new signal.NotifyContext function that was  released with Go 1.16, graceful shutdowns are easier than ever to add  into your application.

使用 Go 1.16 发布的新 signal.NotifyContext 函数，可以比以往更轻松地将优雅关闭添加到您的应用程序中。

Here is a simple web server with a single  handler that will sleep for 10 seconds before responding. If you run  this web server locally, execute a [cURL](https://curl.se) or [Postman](https://www.postman.com) request against it, then immediately send the interrupt signal with  Ctrl+ C. You’ll see the server gracefully shutdown by responding to the  existing request before terminating. If the shutdown is taking too long  another interrupt signal can be sent to exit immediately. Alternatively  the timeout will kick in after 5 seconds.

这是一个带有单个处理程序的简单 Web 服务器，它将在响应前休眠 10 秒。如果您在本地运行此 Web 服务器，请对其执行 [cURL](https://curl.se) 或 [Postman](https://www.postman.com) 请求，然后立即使用 Ctrl+ 发送中断信号C。通过在终止之前响应现有请求，您将看到服务器正常关闭。如果关闭时间太长，可以发送另一个中断信号立即退出。或者，超时将在 5 秒后开始。

```go

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	server http.Server
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server = http.Server{
		Addr: ":8080",
	}

	// Perform application startup.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 10)
		fmt.Fprint(w, "Hello world!")
	})

	// Listen on a different Goroutine so the application doesn't stop here.
	go server.ListenAndServe()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	// Perform application shutdown with a maximum timeout of 5 seconds.
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}
```


Below you'll find some resources that I would recommend  reading if you'd like a better understanding of the signal.NotifyContext function, the context package or the importance of graceful shutdowns.

-  signal.NotifyContext [ documentation](https://pkg.go.dev/os/signal#NotifyContext).
-  JustForFunc Episode 9, [ The Context Package](https://www.youtube.com/watch?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ&v=LSzR0VEraWw).
-  Wayne Ashley Berry's [ article ](https://medium.com/over-engineering/graceful-shutdown-with-go-http-servers-and-kubernetes-rolling-updates-6697e7db17cf) on graceful shutdown with Go http servers and Kubernetes rolling  updates. This is a great article showing a real world situation where  this can be applied. Just note, this article was written pre Go 1.16 so  the code snippets will show the older way of listening for signals from  the operating system.
- Some insightful [comments and feedback](https://www.reddit.com/r/golang/comments/ngx1mu/graceful_shutdowns_in_golang_with) about this post from some kind Gophers over on [r/golang](https://www. reddit.com/r/golang). 

如果您想更好地了解 signal.NotifyContext 函数、上下文包或正常关闭的重要性，您会在下面找到一些我推荐阅读的资源。
- signal.NotifyContext [文档](https://pkg.go.dev/os/signal#NotifyContext)。
- JustForFunc 第 9 集，[上下文包](https://www.youtube.com/watch?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ&v=LSzR0VEraWw)。
- Wayne Ashley Berry 的 [文章](https://medium.com/over-engineering/graceful-shutdown-with-go-http-servers-and-kubernetes-rolling-updates-6697e7db17cf) 关于使用 Go http 服务器和Kubernetes 滚动更新。这是一篇很棒的文章，展示了可以应用它的真实世界情况。请注意，本文是在 Go 1.16 之前编写的，因此代码片段将显示侦听来自操作系统的信号的旧方法。
- 一些来自 [r/golang](https://www.reddit.com/r/golang)。