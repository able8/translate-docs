# Fun with Concurrency in Golang

# Golang 并发的乐趣

October 19, 2019

2019 年 10 月 19 日

> NOTE: The source for this post [is on Gitlab](https://gitlab.com/searsaw/funwithconcurrency).

> 注意：这篇文章的来源 [在 Gitlab 上](https://gitlab.com/searsaw/funwithconcurrency)。

Golang has strong support for concurrency. From the language itself  (goroutines, channels) to constructs in the standard library (WaitGroup, Mutex), the language tries to make it easy on the developer to write  concurrent programs. Let’s play with some of these by creating a program that spins up three different HTTP servers and allows for graceful  shutdowns when the program gets a `SIGTERM` signal.

Golang 对并发有很强的支持。从语言本身（goroutines、channels）到标准库中的构造（WaitGroup、Mutex），该语言试图让开发人员轻松编写并发程序。让我们通过创建一个程序来运行其中的一些，该程序可以启动三个不同的 HTTP 服务器，并在程序收到“SIGTERM”信号时允许正常关闭。

Let’s start by creating a `main.go` file with the following contents.

让我们从创建一个包含以下内容的 `main.go` 文件开始。

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    helloWorldSvr := getHelloWorldServer()
    helloNameSvr := getHelloNameServer()
    echoSvr := getEchoServer()

    helloWorldSvr.ListenAndServe()
    helloNameSvr.ListenAndServe()
    echoSvr.ListenAndServe()
    fmt.Println("all servers are started")
}

func getHelloWorldServer() *http.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`Hello, world!`))
    })

    return &http.Server{Addr: ":7000", Handler: mux}
}

func getHelloNameServer() *http.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        params := r.URL.Query()
        name := params.Get("name")

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
    })

    return &http.Server{Addr: ":8000", Handler: mux}
}

func getEchoServer() *http.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        io.Copy(w, r.Body)
    })

    return &http.Server{Addr: ":9000", Handler: mux}
}
```

Here we are creating three random HTTP servers: one who simply returns “Hello, world!”, one that takes a `name` query param and says “Hello” to the name, and one that sends back whatever is sent to it in the body of the request.

在这里，我们创建了三个随机 HTTP 服务器：一个简单地返回“Hello, world!”，一个接受 `name` 查询参数并对该名称说“Hello”，一个将在请求的正文。

If we run this program, we can start making requests to our servers. I’ll be using [`HTTPie`](https://httpie.org/) to send requests in this post.

如果我们运行这个程序，我们就可以开始向我们的服务器发出请求。在这篇文章中，我将使用 [`HTTPie`](https://httpie.org/) 发送请求。

```bash
$ http :7000
HTTP/1.1 200 OK
Content-Length: 13
Content-Type: text/plain;charset=utf-8
Date: Sat, 19 Oct 2019 21:53:32 GMT

Hello, world!

$ http :8000 name==Alex

http: error: ConnectionError: HTTPConnectionPool(host='localhost', port=8000): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x107e3fa10>: Failed to establish a new connection:[Errno 61] Connection refused')) while doing GET request to URL: http://localhost:8000/

$ http POST :9000 name=Alex job="Software Engineer" coffee=please

http: error: ConnectionError: HTTPConnectionPool(host='localhost', port=9000): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x104a3bb90>: Failed to establish a new connection:[Errno 61] Connection refused')) while doing GET request to URL: http://localhost:9000/
```

Well, that’s strange. The first one is running, but the other two don’t appear to be running. Well, this is because the `ListenAndServe` method on an `http.Server` struct is a blocking call. This means our program stops there until  that method returns. To get it to return, the server needs to be  shutdown. Then it would go to the next one, which will also block, and  so on and so forth.

嗯，这很奇怪。第一个正在运行，但其他两个似乎没有运行。嗯，这是因为 `http.Server` 结构上的 `ListenAndServe` 方法是一个阻塞调用。这意味着我们的程序在那里停止，直到该方法返回。要让它返回，需要关闭服务器。然后它会转到下一个，它也会阻塞，依此类推。

We need a way to run all of these at the same time. Luckily, Go has  us covered with goroutines. Goroutines are lightweight threads managed  by the Go runtime which allows us to kick off another “process,” for  lack of a better word, and continue on with our program. In this case,  that means it will start each server in a separate goroutine and  continue on with our main one.

我们需要一种同时运行所有这些的方法。幸运的是，Go 为我们提供了 goroutines。 Goroutines 是由 Go 运行时管理的轻量级线程，它允许我们启动另一个“进程”，因为没有更好的词，并继续我们的程序。在这种情况下，这意味着它将在单独的 goroutine 中启动每个服务器，并继续我们的主要 goroutine。

We can update our `ListenAndServe` calls like so:

我们可以像这样更新我们的 `ListenAndServe` 调用：

```go
go helloWorldSvr.ListenAndServe()
go helloNameSvr.ListenAndServe()
go echoSvr.ListenAndServe()
```

If we run this, we see that our program exits  pretty quickly. Hmmm…well this is because we kicked off three  goroutines, which let our `main` function continue to the end since we didn’t have anything to stop it from completing. How can we  make the program wait to complete until we give it the signal to quit? Well, once again, Go has us covered. We can use the `signal` package built into the standard library along with channels to accomplish this.

如果我们运行它，我们会看到我们的程序很快退出。嗯……嗯，这是因为我们启动了三个 goroutine，这让我们的 main 函数一直持续到最后，因为我们没有任何东西可以阻止它完成。我们如何让程序等待完成直到我们给它退出信号？好吧，再一次，Go 已经涵盖了我们。我们可以使用标准库中内置的`signal` 包和通道来实现这一点。

I like to think of channels as a tube that we can send messages down  from one part of our program to be picked up by another part. Channels  can be passed around in goroutines to help coordinate all kinds of  behavior. There is a little more to channels than my simplified  explanation. [Go by Example](https://gobyexample.com/channels) has a good set of posts about channels that I recommend checking out.

我喜欢将通道视为一根管道，我们可以将消息从程序的一部分向下发送，供另一部分接收。通道可以在 goroutine 中传递以帮助协调各种行为。除了我的简化解释之外，频道还有更多内容。 [Go by Example](https://gobyexample.com/channels) 有一组关于频道的好帖子，我建议您查看。

Let’s add the following to the end of our main function after starting our HTTP servers.

在启动我们的 HTTP 服务器后，让我们在 main 函数的末尾添加以下内容。

```go
signals := make(chan os.Signal, 1)
signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
<-signals
```

Here, we are creating a channel that will hold one signal at a time. Then we use the `Notify` function from the `signal` package to tell our program to tell us when SIGINT and SIGTERM are sent from the user. We will be notified by that signal being sent on the  channel. The last line is asking for a value off of the channel. This is a blocking operation that will only continue when it gets a value or  the channel is closed. Therefore, this line prevents our program from  continuing until it receives a SIGINT or SIGTERM, which would be sent by us when we stop the program .

在这里，我们正在创建一个每次保存一个信号的通道。然后我们使用`signal` 包中的`Notify` 函数告诉我们的程序在用户发送SIGINT 和SIGTERM 时告诉我们。我们将收到在频道上发送的信号通知。最后一行是要求通道外的值。这是一个阻塞操作，只有在获得值或通道关闭时才会继续。因此，这一行阻止我们的程序继续，直到它收到 SIGINT 或 SIGTERM，当我们停止程序时，我们会发送这些信息。

Run the program and send the requests from earlier. They all succeed!

运行程序并发送之前的请求。他们都成功了！

```bash
$ http :7000
HTTP/1.1 200 OK
Content-Length: 13
Content-Type: text/plain;charset=utf-8
Date: Sat, 19 Oct 2019 22:34:11 GMT

Hello, world!

$ http :8000 name==Alex
HTTP/1.1 200 OK
Content-Length: 12
Content-Type: text/plain;charset=utf-8
Date: Sat, 19 Oct 2019 22:34:16 GMT

Hello, Alex!

$ http POST :9000 name=Alex job="Software Engineer" coffee=please
HTTP/1.1 200 OK
Content-Length: 64
Content-Type: text/plain;charset=utf-8
Date: Sat, 19 Oct 2019 22:34:20 GMT

{
    "coffee": "please",
    "job": "Software Engineer",
    "name": "Alex"
}
```

Well this is great, but this isn’t as good as we  can make it. We need to try to shutdown our servers gracefully to ensure all requests are serviced before the server closes. Otherwise our end  users may send a request but never get a response! Luckily, the server  struct has a `Shutdown` method on it for us to use for just this purpose! Let’s refactor our code a bit to account for this.

嗯，这很好，但这还没有我们能做到的那么好。我们需要尝试优雅地关闭我们的服务器，以确保在服务器关闭之前处理所有请求。否则我们的最终用户可能会发送请求但永远不会得到响应！幸运的是，服务器结构上有一个 `Shutdown` 方法供我们使用！让我们稍微重构一下我们的代码来解决这个问题。

```go
func getHelloWorldServer() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`Hello, world!`))
    })
    server := &http.Server{Addr: ":7000", Handler: mux}

    go func() {
        shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        if err := server.Shutdown(shutCtx);err != nil {
            fmt.Printf("error shutting down the hello world server: %s\n", err)
        }
        fmt.Println("the hello world server is closed")
    }()

    fmt.Println("the hello world server is starting")
    if err := server.ListenAndServe();err != http.ErrServerClosed {
        fmt.Printf("error starting the hello world server: %s\n", err)
    }
    fmt.Println("the hello world server is closing")
}
```

For the sake of brevity, I’ve only shown how I updated the `getHelloWorldServer` function. I have moved the starting of the server into this function. I have also started *another* goroutine in this function that will shutdown the server. We are  creating a context that will automatically timeout and force the  shutdown after five seconds.

为简洁起见，我只展示了如何更新 `getHelloWorldServer` 函数。我已将服务器的启动移动到此功能中。我还在这个函数中启动了 *另一个 * goroutine，它将关闭服务器。我们正在创建一个上下文，该上下文将在五秒后自动超时并强制关闭。

We need to update our main function like so:

我们需要像这样更新我们的主要功能：

```go
func main() {
    go getHelloWorldServer()
    go getHelloNameServer()
    go getEchoServer()

    fmt.Println("all servers are started")

    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
    <-signals
}
```

We are starting the goroutines when we “get” the  servers now. Try running this program. We get some errors! We can see  that we tried to start a server that was already closed. This happened  because the goroutine in each server function runs before the server is  even started. We would prefer that piece of code not run until we have  signalled for the program to shutdown. We need a way to close or cancel  something from our `main` function that will let our  goroutines know it’s time to shutdown the servers. This is a great  use-case for a context. Contexts are used all throughout the Golang  standard library as a way to signal when to stop things from running and also as a way to pass values down to functions. We will use one for the former reason.

当我们现在“获得”服务器时，我们正在启动 goroutine。试试运行这个程序。我们得到一些错误！我们可以看到我们试图启动一个已经关闭的服务器。发生这种情况是因为每个服务器函数中的 goroutine 甚至在服务器启动之前就运行了。我们希望这段代码在我们发出关闭程序的信号之前不运行。我们需要一种方法来关闭或取消我们的 main 函数中的某些东西，让我们的 goroutine 知道是时候关闭服务器了。这是上下文的一个很好的用例。上下文在整个 Golang 标准库中都被用作发出信号何时停止运行以及将值传递给函数的一种方式。由于前一个原因，我们将使用一个。

First we need to create the context in our `main` function. It’s important we create one *and* get a cancel function for it so we can cancel it later in our program.

首先，我们需要在 `main` 函数中创建上下文。重要的是我们创建一个*并*为它获取一个取消函数，这样我们就可以稍后在我们的程序中取消它。

```go
ctx, cancel := context.WithCancel(context.Background())
```

Next, we should pass this down to each of our server functions.

接下来，我们应该将其传递给我们的每个服务器功能。

```go
go getHelloWorldServer(ctx)
go getHelloNameServer(ctx)
go getEchoServer(ctx)
```

Update the function definitions to take this context.

更新函数定义以获取此上下文。

```go
func getHelloWorldServer(ctx context.Context) {
```

Now, we need to wait for the context to be  cancelled before we try to shutdown the servers. This can be done by  waiting for a message on the `Done` channel of the context. This channel will be closed when we call cancel from our `main` function. This means any places we are waiting for a message will  continue moving forward. Add this waiting right before we create the `shutCtx` in each function.

现在，在尝试关闭服务器之前，我们需要等待上下文被取消。这可以通过等待上下文的“完成”通道上的消息来完成。当我们从`main` 函数调用cancel 时，该通道将关闭。这意味着我们正在等待消息的任何地方都将继续前进。在我们在每个函数中创建 `shutCtx` 之前添加这个等待。

```go
go func() {
    <-ctx.Done()
    shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
```

Now that our servers will wait for the context to  be cancelled, we need to cancel it at the right time. We want to cancel  it once we get a signal to stop everything. So let’s cancel the context  right after we get that signal!

现在我们的服务器将等待上下文被取消，我们需要在正确的时间取消它。一旦我们收到停止一切的信号，我们就想取消它。所以让我们在收到信号后立即取消上下文！

```go
signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
<-signals
cancel()
```

Run the server again and make a few requests. It  works again! Now close the program. Notice when it stops, we don’t see  our logs saying each of our servers are closed. We need a way for the  goroutines to signal to the `main` function they are done and for the `main` function to wait until they are all gone. Go’s standard library has a `sync` package with a `WaitGroup` construct in it. We can create a group and tell it how many are in the  group. Then each goroutine will tell the group that it is done. We can  then make the `main` function wait for all the goroutines in the group to be done before moving forward with the program.

再次运行服务器并发出一些请求。它又起作用了！现在关闭程序。请注意，当它停止时，我们没有看到我们的日志说我们的每个服务器都已关闭。我们需要一种方法让 goroutine 向 `main` 函数发出信号，它们已经完成，并且让 `main` 函数等待它们全部消失。 Go 的标准库有一个 `sync` 包，其中包含一个 `WaitGroup` 结构。我们可以创建一个组并告诉它该组中有多少人。然后每个 goroutine 会告诉组它完成了。然后我们可以让 `main` 函数等待组中的所有 goroutine 完成，然后再继续执行程序。

First let’s create the `WaitGroup` at the beginning of the `main` function and initialize the count in it to be three since we know we will have three servers.

首先让我们在 `main` 函数的开头创建 `WaitGroup`，并将其中的计数初始化为 3，因为我们知道我们将拥有三台服务器。

```go
func main() {
    var wg sync.WaitGroup
    wg.Add(3)
```

Next we need to pass a reference to the group down to each of our server functions.

接下来，我们需要将对该组的引用传递给我们的每个服务器功能。

```go
go getHelloWorldServer(ctx, &wg)
go getHelloNameServer(ctx, &wg)
go getEchoServer(ctx, &wg)
```

Let’s also update our server function definitions to take this `WaitGroup` as a parameter.

让我们也更新我们的服务器函数定义，以将此 `WaitGroup` 作为参数。

```go
func getHelloWorldServer(ctx context.Context, wg *sync.WaitGroup) {
```

Inside our server functions, make sure to call `wg.Done()` after the server has completely shutdown to decrease the count in the `WaitGroup` by one.

在我们的服务器函数中，确保在服务器完全关闭后调用 `wg.Done()` 以将 `WaitGroup` 中的计数减一。

```go
}
fmt.Println("the hello world server is closed")
wg.Done()
```

Lastly, we need to make our `main` function wait for the `WaitGroup` counter to be zero before continuing forward. We will put this at the bottom of the `main` function.

最后，我们需要让我们的 `main` 函数在继续前进之前等待 `WaitGroup` 计数器为零。我们将把它放在 `main` 函数的底部。

```go
<-signals
cancel()
wg.Wait()
```

Run the program again and then shut it down. Notice, though they may be out of order since it all happens very  quickly, we get all our logs. Woohoo! 

再次运行该程序，然后将其关闭。请注意，尽管它们可能会出现故障，因为这一切都发生得非常快，但我们会获取所有日志。呜呼！

So we have a solid base here, but there’s one case that isn’t  handled. What if one of the servers has an error when starting up? We  should probably have the whole program fail. Right now, it will just not start up that one server. We want the program to fail fast so we can  have a quick feedback loop. The package [golang.org/x/sync/errgroup](https://godoc.org/golang.org/x/sync/errgroup) provides the `errgroup.Group` construct we can use to handle the errors returned by starting or stopping our servers.

所以我们在这里有一个坚实的基础，但有一个案例没有处理。如果其中一台服务器在启动时出错怎么办？我们可能应该让整个程序失败。现在，它不会启动那台服务器。我们希望程序快速失败，这样我们就可以有一个快速的反馈循环。包 [golang.org/x/sync/errgroup](https://godoc.org/golang.org/x/sync/errgroup) 提供了 `errgroup.Group` 构造，我们可以用来处理启动返回的错误或停止我们的服务器。

Let’s update our code to use an `errgroup.Group` to handle the error handling for all our goroutines. First, let’s see our `main` function.

让我们更新我们的代码以使用 `errgroup.Group` 来处理我们所有 goroutine 的错误处理。首先，让我们看看我们的 main 函数。

```go
import (
    // ...other stuff
    "golang.org/x/sync/errgroup"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3)

    ctx, cancel := context.WithCancel(context.Background())

    eg, egCtx := errgroup.WithContext(context.Background())
    eg.Go(getHelloWorldServer(ctx, &wg))
    eg.Go(getHelloNameServer(ctx, &wg))
    eg.Go(getEchoServer(ctx, &wg))

    go func() {
        <-egCtx.Done()
        cancel()
    }()

    go func() {
        signals := make(chan os.Signal, 1)
        signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
        <-signals
        cancel()
    }()

    if err := eg.Wait();err != nil {
        fmt.Printf("error in the server goroutines: %s\n", err)
        os.Exit(1)
    }
    fmt.Println("everything closed successfully")
}
```

We now are creating an `errgroup.Group` using the `WithContext` function and are kicking off our server goroutines in the group using its `Go` method. We are using the context from the errgroup as a way to know when to cancel the rest of the goroutines. This `Done` channel is closed when either the `eg.Wait` function returns, which means all the goroutines have completed, or an  error was returned from one of the goroutines. This is the important  piece. We want to be notified when one of the goroutines errors out so  we can stop immediately. Also, we have put our signal handling logic  into a goroutine since we are not using it to halt our `main` function. We are simply using it to signal to our goroutines when to  stop when things are running correctly. Lastly, we use the `Wait` method on the errgroup to halt the `main` function until all the servers have been shutdown. It then returns the  first error that was returned from any of the goroutines or `nil`.

我们现在正在使用 `WithContext` 函数创建一个 `errgroup.Group`，并使用它的 `Go` 方法启动组中的服务器 goroutine。我们使用 errgroup 中的上下文来了解何时取消其余的 goroutine。当 `eg.Wait` 函数返回时，这个 `Done` 通道关闭，这意味着所有 goroutines 已经完成，或者从 goroutines 之一返回错误。这是重要的部分。我们希望在其中一个 goroutine 出错时得到通知，以便我们可以立即停止。此外，我们将信号处理逻辑放入 goroutine 中，因为我们没有使用它来停止我们的 main 函数。我们只是使用它向我们的 goroutine 发出信号，当事情正常运行时何时停止。最后，我们在 errgroup 上使用 `Wait` 方法来停止 `main` 函数，直到所有服务器都关闭。然后它返回从任何 goroutine 或 `nil` 返回的第一个错误。

Now let’s see how our server functions have changed.

现在让我们看看我们的服务器功能发生了怎样的变化。

```go
func getHelloWorldServer(ctx context.Context, wg *sync.WaitGroup) func() error {
    return func() error {
        mux := http.NewServeMux()
        mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusOK)
            w.Write([]byte(`Hello, world!`))
        })
        server := &http.Server{Addr: ":7000", Handler: mux}
        errChan := make(chan error, 1)

        go func() {
            <-ctx.Done()
            shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
            defer cancel()
            if err := server.Shutdown(shutCtx);err != nil {
                errChan <- fmt.Errorf("error shutting down the hello world server: %w", err)
            }
            fmt.Println("the hello world server is closed")
            close(errChan)
            wg.Done()
        }()

        fmt.Println("the hello world server is starting")
        if err := server.ListenAndServe();err != http.ErrServerClosed {
            return fmt.Errorf("error starting the hello world server: %w", err)
        }
        fmt.Println("the hello world server is closing")
        err := <-errChan
        wg.Wait()
        return err
    }
}
```

First, our definition has been updated to return a function that itself returns an error. This is because the `Go` function of the errgroup takes a function that takes no parameters and  returns an error. On the next line, we are returning this function as an anonymous function. Inside this anonymous function is all our logic for starting and stopping our server. Below our server, we are creating  another channel that accepts errors. This channel is used to prevent the function that starts the server from returning before the goroutine  that stops the server finishes. It also has the added benefit of giving  us the error from that server if there was one.

首先，我们的定义已更新为返回一个本身返回错误的函数。这是因为 errgroup 的 `Go` 函数采用了一个不带参数的函数并返回一个错误。在下一行，我们将此函数作为匿名函数返回。在这个匿名函数中是我们启动和停止服务器的所有逻辑。在我们的服务器下方，我们正在创建另一个接受错误的通道。该通道用于防止启动服务器的函数在停止服务器的 goroutine 完成之前返回。它还有一个额外的好处，那就是给我们来自那个服务器的错误（如果有的话）。

Next, inside the goroutine for stopping the server, we are putting an error on the `errChan` if there was an error while shutting the server down. We then close the error outside of the conditional to ensure we always get something off  the channel when we wait on the channel later in the function. Lastly,  we call `Done` on the `WaitGroup` to make sure other goroutines waiting on this server to shutdown are aware it’s done.

接下来，在用于停止服务器的 goroutine 中，如果在关闭服务器时出现错误，我们将在 `errChan` 上放置一个错误。然后我们关闭条件之外的错误，以确保当我们稍后在函数中等待通道时，我们总是从通道中得到一些东西。最后，我们在 `WaitGroup` 上调用 `Done` 以确保在此服务器上等待关闭的其他 goroutine 知道它已经完成。

Back outside the goroutine, we are still starting the server and  handling the error. However, handling the error, in this case, means  wrapping the error and returning it from the function. Remember, any  error returned from one of these functions will cause the errgroup  context to cancel in our `main` function, which will then cause all the other goroutines to cancel as well. If `ListenAndServe` exits without any errors, we then wait on a message on the `errChan`. Getting a message here means either there was an error while the server was shutting down or the channel was closed without any errors put on  it. Whatever is returned from the error channel we assign to a variable. We then wait for the `WaitGroup` counter to reach zero. When it does, this means all the other servers have completed and it’s time  to exit from them all. So, at the end, we simply return what came off  the error channel. If any of them are actual errors, the first one will  be returned from the `Wait` of the errgroup in our `main` function, and the program will terminate.

回到 goroutine 之外，我们仍在启动服务器并处理错误。但是，在这种情况下，处理错误意味着包装错误并从函数中返回它。请记住，从这些函数之一返回的任何错误都将导致我们的 main 函数中的 errgroup 上下文取消，这将导致所有其他 goroutine 也取消。如果`ListenAndServe` 没有任何错误退出，我们就等待`errChan` 上的消息。在此处获取消息意味着服务器关闭时出现错误或通道已关闭而没有任何错误。无论从错误通道返回什么，我们都会分配给一个变量。然后我们等待 `WaitGroup` 计数器达到零。当它发生时，这意味着所有其他服务器都已完成，是时候退出所有服务器了。所以，最后，我们简单地返回错误通道中的内容。如果其中任何一个是实际错误，第一个将从我们的 `main` 函数中的 errgroup 的 `Wait` 返回，并且程序将终止。

Run the program, send some requests, shut it down. Change the address for the hello world server to be `:8000` and see it all crash!

运行程序，发送一些请求，然后关闭它。将 hello world 服务器的地址更改为 `:8000` 并看到它全部崩溃！

# Going forward

#    往前走

So we built a program that creates three HTTP servers and coordinates gracefully shutting them down while handling errors at all the  necessary points to ensure it fails fast. What’s next? Well, one easy  thing to do would be to refactor this code. The three HTTP server  functions are all the same outside of a few small things such as name  and what their handlers actually do. This could easily be abstracted out into a single function five parameters. There may be other ways to  clean up this code. I leave that as a challenge for the reader. If you  want to see how I cleaned up the code, check out [the source on Gitlab](https://gitlab.com/searsaw/funwithconcurrency). 

所以我们构建了一个程序来创建三个 HTTP 服务器，并协调优雅地关闭它们，同时在所有必要的点处理错误以确保它快速失败。下一步是什么？嗯，一件简单的事情就是重构这段代码。除了名称和它们的处理程序实际执行的操作之外，这三个 HTTP 服务器函数都是相同的。这可以很容易地抽象为单个函数五个参数。可能还有其他方法可以清理此代码。我认为这是对读者的挑战。如果您想了解我是如何清理代码的，请查看 [Gitlab 上的源代码](https://gitlab.com/searsaw/funwithconcurrency)。

