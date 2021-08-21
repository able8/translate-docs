# Using Context in Golang - Cancellation, Timeouts and Values (With Examples)

# 在 Golang 中使用上下文 - 取消、超时和值（附示例）

# Updated on July 07, 2021 From: https://www.sohamkamani.com/golang/context-cancellation-and-values/

In this post, we’ll learn about Go’s [context](https://pkg.go.dev/context#pkg-index) package, and more specifically, how we can use cancellation to improve our applications performance.

在这篇文章中，我们将了解 Go 的 [context](https://pkg.go.dev/context#pkg-index) 包，更具体地说，我们将了解如何使用取消来提高我们的应用程序性能。

[![banner image](https://www.sohamkamani.com/static/819f2c40e11b07132e7e3f1130c62eab/5a190/banner.png)](https://www.sohamkamani.com/static/819f2c40e11b07132e7e3f1130c62eab/00d43/banner.png)

We will also go through some patterns and best practices when using the `context` package in your Golang application.

在您的 Golang 应用程序中使用 `context` 包时，我们还将介绍一些模式和最佳实践。

## When Do We Use Context?

## 我们什么时候使用上下文？

As the name suggests, we use the context package whenever we want to pass around “context”, or *common scoped data* within our application. For example:

- Request IDs for function calls and goroutines that are part of an HTTP request call
- Errors when fetching data from a database
- Cancellation signals whe performing async operations using goroutines

顾名思义，每当我们想在我们的应用程序中传递“上下文”或*通用范围数据*时，我们都会使用上下文包。例如：
- 作为 HTTP 请求调用一部分的函数调用和 goroutine 的请求 ID
- 从数据库中获取数据时出错
- 使用 goroutines 执行异步操作时的取消信号

![context refers to common scoped data within goroutines or function calls](https://www.sohamkamani.com/1005523fd11d74bd03fef5564e999eaf/context-in-practice.svg)

Using the [Context](https://pkg.go.dev/context#Context) data type is the idiomatic way to pass information across these kind of operations, such as:
1. Cancellation signals to terminate the operation
2. Miscellaneous data required at every function call invoked by the operation

使用 [Context](https://pkg.go.dev/context#Context) 数据类型是在这些操作之间传递信息的惯用方式，例如：
1. 取消信号以终止操作
2. 操作调用的每个函数调用所需的杂项数据

Let’s talk about cancellation first:

先说取消吧：

## Why Do We Need Cancellation?

## 为什么我们需要取消？

In short, we need cancellation to prevent our system from doing unnecessary work. Consider the common situation of an HTTP server making a call to a database, and returning the queried data to the client:

简而言之，我们需要取消以防止我们的系统做不必要的工作。考虑HTTP服务器调用数据库，并将查询到的数据返回给客户端的常见情况：

![client server model diagram](https://www.sohamkamani.com/199c2b8faf7663c9b7e83de127012a6c/client-diagram.svg)

The timing diagram, if everything worked perfectly, would look like this:

如果一切正常，时序图将如下所示：

![timing diagram with all events finishing](https://www.sohamkamani.com/ff6e4d831668b9da81c1c214224e4521/timing-ideal.svg)

But, what would happen if the client cancelled the request in the middle? This could happen if, for example, the client closed their browser  mid-request.

但是，如果客户端在中间取消了请求会发生什么？例如，如果客户端在请求中途关闭了浏览器，则可能会发生这种情况。

Without cancellation, the application server and  database would continue to do their work, even though the result of that work would be wasted:

如果不取消，应用程序服务器和数据库将继续执行其工作，即使该工作的结果将被浪费：

![timing diagram with http request cancelled, and other processes still taking place](https://www.sohamkamani.com/4955e194034f42b5edd7632f1461c124/timing-without-cancel.svg)

Ideally, we would want all downstream components of a process to halt, if we *know* that the process (in this example, the HTTP request) halted:

理想情况下，如果我们*知道*进程（在本例中为 HTTP 请求）已停止，我们希望进程的所有下游组件都停止：

![timing diagram with all processes cancelling once HTTP request is cancelled](https://www.sohamkamani.com/2af484f735aab3022ea8d7a9a9c1b675/timing-with-cancel.svg)

## Context Cancellation in Go

## Go 中的上下文取消

Now that we know *why* we need cancellation, let’s get into *how* you can implement it in Go.

现在我们知道*为什么*我们需要取消，让我们进入*如何*在Go中实现它。

Because “cancellation” is highly contextual to the operation being performed, the best way to implement it is through `context`.

因为“取消”与正在执行的操作高度相关，实现它的最佳方法是通过“上下文”。

There are two sides to context cancellation:

1. Listening for the cancellation event
2. Emitting the cancellation event

上下文取消有两个方面：
1. 监听取消事件
2. 发出取消事件

### Listening For Cancellation

### 监听取消

The `Context` type provides a `Done()` method. This returns a [channel](https://www.sohamkamani.com/golang/channels) that receives an empty `struct{}` type every time the context receives a cancellation event.

`Context` 类型提供了一个 `Done()` 方法。这将返回一个 [channel](https://www.sohamkamani.com/golang/channels)，每当上下文接收到取消事件时，它都会接收一个空的 `struct{}` 类型。

So, to listen for a cancellation event, we need to wait on `<- ctx.Done()`.

因此，要监听取消事件，我们需要等待 `<- ctx.Done()`。

For example, lets consider an HTTP server that takes two seconds to process an event. If the request gets cancelled before that, we want to return  immediately:

例如，让我们考虑一个需要两秒钟来处理事件的 HTTP 服务器。如果请求在此之前被取消，我们希望立即返回：

```go
func main() {
	// Create an HTTP server that listens on port 8000
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Fprint(os.Stdout, "processing request\n")
		// We use `select` to execute a piece of code depending on which
		// channel receives a message first
		select {
		case <-time.After(2 * time.Second):
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// We then write this as the response
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request gets cancelled, log it
			// to STDERR
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))
}
```


> You can view the source code for all the examples [on Github](https://github.com/sohamkamani/blog-example-go-context-cancellation) 
> 你可以查看所有示例的源代码 [在 Github 上](https://github.com/sohamkamani/blog-example-go-context-cancellation)

You can test this by running the server and opening [localhost:8000](http://localhost:8000) on your browser. If you close your browser before 2 seconds, you should see “request cancelled” printed on the terminal window.

您可以通过运行服务器并在浏览器上打开 [localhost:8000](http://localhost:8000) 来测试这一点。如果您在 2 秒前关闭浏览器，您应该会在终端窗口上看到“请求已取消”。

### Emitting a Cancellation Event

### 发出取消事件

If you have an operation that could be cancelled, you will have to emit a cancellation event through the context.

如果您有可以取消的操作，则必须通过上下文发出取消事件。

This can be done using the `WithCancel` function in the [context package](https://pkg.go.dev/context#WithCancel), which returns a context object, and a function.

这可以使用 [context package](https://pkg.go.dev/context#WithCancel) 中的 `WithCancel` 函数来完成，它返回一个上下文对象和一个函数。

```go
ctx, fn := context.WithCancel(ctx)
```


This function takes no arguments, and does not return anything, and is called when you want to cancel the context.

此函数不接受任何参数，也不返回任何内容，并在您想要取消上下文时调用。

Consider the case of 2 dependent operations. Here, “dependent” means if one  fails, it doesn’t make sense for the other to complete. If we get to  know early on that one of the operations failed, we would like to cancel all dependent operations.

考虑 2 个依赖操作的情况。在这里，“依赖”意味着如果一个失败，另一个完成是没有意义的。如果我们尽早知道其中一项操作失败，我们希望取消所有相关操作。

```go
func operation1(ctx context.Context) error {
 // Let's assume that this operation failed for some reason
 // We use time.Sleep to simulate a resource intensive operation
 time.Sleep(100 * time.Millisecond)
 return errors.New("failed")
 }

 func operation2(ctx context.Context) {
 // We use a similar pattern to the HTTP server
 // that we saw in the earlier example
 select {
 case <-time.After(500 * time.Millisecond):
 fmt.Println("done")
 case <-ctx.Done():
 fmt.Println("halted operation2")
 }
 }

 func main() {
 // Create a new context
 ctx := context.Background()
 // Create a new context, with its cancellation function
 // from the original context
 ctx, cancel := context.WithCancel(ctx)

 // Run two operations: one in a different go routine
 go func() {
 err := operation1(ctx)
 // If this operation returns an error
 // cancel all operations using this context
 if err != nil {
 cancel()
 }
 }()

 // Run operation2 with the same context we use for operation1
 operation2(ctx)
 }
```



### Context Timeouts

### 上下文超时

Any application that needs to maintain an SLA (service level agreement) for the maximum duration of a request, should use time based cancellation.

任何需要在请求的最长持续时间内维护 SLA（服务级别协议）的应用程序都应使用基于时间的取消。

The API is almost the same as the previous example, with a few additions:

API 与前面的示例几乎相同，但有一些补充：

```go
// The context will be cancelled after 3 seconds
 // If it needs to be cancelled earlier, the `cancel` function can
 // be used, like before
 ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

 // Setting a context deadline is similar to setting a timeout, except
 // you specify a time when you want the context to cancel, rather than a duration.
 // Here, the context will be cancelled on 2009-11-10 23:00:00
 ctx, cancel := context.WithDeadline(ctx, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
```


For example, consider making an HTTP API call to an external service. If  the service takes too long, it’s better to fail early and cancel the  request:

例如，考虑对外部服务进行 HTTP API 调用。如果服务时间太长，最好早点失败并取消请求：

```go
func main() {
	// Create a new context
	// With a deadline of 100 milliseconds
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 300*time.Millisecond)

	// Make a request, that will call the google homepage
	req, _ := http.NewRequest(http.MethodGet, "http://jd.com", nil)
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	res, err := client.Do(req)
	// If the request failed, log to STDOUT
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	// Print the status code if the request succeeds
	fmt.Println("Response received, status code:", res.StatusCode)
}
```


Based on how fast the google homepage responds to your request, you will receive:

根据谷歌主页对您的请求的响应速度，您将收到：

```text
Response received, status code: 200
```


or

```text
Request failed: Get http://google.com: context deadline exceeded
```


You can play around with the timeout to achieve both of the above results.

您可以使用超时来实现上述两个结果。

### Is this code safe to run?

### 这段代码可以安全运行吗？

```go
func doSomething() {
 ctx, cancel := context.WithCancel(ctx)
 defer cancel()

 someArg := "loremipsum"
 go doSomethingElse(ctx, someArg)
 }
```


Would it be safe to execute the `doSomething` function?

执行`doSomething` 函数是否安全？

## Context Values

## 上下文值

You can use the context variable to pass around values that are common  across an operation. This is the more idiomatic alternative to just  passing them around as variables throughout your function calls.

您可以使用上下文变量来传递操作中通用的值。这是在整个函数调用中将它们作为变量传递的更惯用的替代方法。

For example, consider an operation that has multiple function calls, with a common ID used to identify it for logging and monitoring.

例如，考虑一个具有多个函数调用的操作，使用一个公共 ID 来标识它以进行日志记录和监视。

The naive way to implement this is the just pass the ID around for each function call:

实现这一点的最简单的方法是为每个函数调用传递 ID：

```go
func main() {
 // create a random integer as the ID
 rand.Seed(time.Now().Unix())
 id := rand.Int()
 operation1(id)
 }

 func operation1(id int64) {
 // do some work
 log.Println("operation1 for id:", id, " completed")
 operation2(id)
 }

 func operation2(id int64) {
 // do some work
 log.Println("operation2 for id:", id, " completed")
 }
```


This can quickly get bloated when you have more information that you want to provide.

当您想要提供更多信息时，这会很快变得臃肿。

We can implement the same functionality using context:

我们可以使用上下文实现相同的功能：

```go
// we need to set a key that tells us where the data is stored
const keyID = "id"

func main() {
	rand.Seed(time.Now().Unix())
	ctx := context.WithValue(context.Background(), keyID, rand.Int())
	operation1(ctx)
}

func operation1(ctx context.Context) {
	// do some work
	// we can get the value from the context by passing in the key
	log.Println("operation1 for id:", ctx.Value(keyID), " completed")
	operation2(ctx)
}

func operation2(ctx context.Context) {
	// do some work

	// this way, the same ID is passed from one function call to the next
	log.Println("operation2 for id:", ctx.Value(keyID), " completed")
}
```


Here, we’re creating a new context variable in `main` with a key value pair associated with it. The value can then be used by the successive function calls obtain contextual information.

在这里，我们在 `main` 中创建了一个新的上下文变量，并带有与之关联的键值对。然后该值可以被连续的函数调用使用，获取上下文信息。

![main creates a new context which is passed to other functions](https://www.sohamkamani.com/7e24e83cbc6a836b39c786c7461a282c/context-values.svg)

Using the context variable to pass down operation-scoped information is useful for a number of reasons:


1. It is **thread safe**: You can’t modify the value of a context key once it has been set. The  only way set another value for a given key is to create another context  variable using `context.WithValue`
2. It is **conventional**: The context package is used throughout Go’s official libraries and  applications to convey operation-scoped data. Other developers and  libraries generally play nicely with this pattern.

使用上下文变量传递操作范围的信息很有用，原因有很多：
1. 它是**线程安全的**：上下文键的值一旦设置就不能修改。为给定键设置另一个值的唯一方法是使用 `context.WithValue` 创建另一个上下文变量
2. 它是**传统**：在 Go 的官方库和应用程序中使用上下文包来传达操作范围的数据。其他开发人员和库通常可以很好地使用这种模式。

## Gotchas and Caveats

## 陷阱和注意事项

Although context cancellation in Go is a versatile tool, there are a few things  that you should keep in mind before proceeding. The most important of  which, is that a **context can only be cancelled once.**

尽管 Go 中的上下文取消是一个多功能工具，但在继续之前，您应该记住一些事情。其中最重要的是 **context 只能取消一次。**

If there are multiple errors that you would want to propagate in the same  operation, then using context cancellation may not the best option.

如果您希望在同一操作中传播多个错误，则使用上下文取消可能不是最佳选择。

The most idiomatic way to use cancellation is when you *actually want to cancel something*, and not just notify downstream processes that an error has occurred.

使用取消最惯用的方式是当您*实际上想要取消某些东西*，而不只是通知下游进程发生了错误。

Another important caveat has to do with wrapping the same context multiple times.

另一个重要的警告与多次包装相同的上下文有关。

Wrapping an *already cancellable* context with `WithTimeout` or `WithCancel` will enable multiple locations in your code in which your context could be cancelled, and this should be avoided.

用`WithTimeout` 或`WithCancel` 包装一个*已经可取消的* 上下文将使你的代码中的多个位置可以取消你的上下文，应该避免这种情况。

> The source code for all the above examples can be found [on Github](https://github.com/sohamkamani/blog-example-go-context-cancellation)

> 以上所有示例的源代码可以在 [Github](https://github.com/sohamkamani/blog-example-go-context-cancellation) 上找到

------

Like what I write? Join my mailing list, and I'll let you know whenever I write another post. No spam, I promise!


喜欢我写的吗？加入我的邮件列表，我会在写另一篇文章时通知你。没有垃圾邮件，我保证！

Written by **Soham Kamani**, an [author](https://www.packtpub.com/books/info/authors/soham-kamani),and a full-stack developer who has extensive experience in the JavaScript  ecosystem, and building large scale applications in Go. He is an [open source enthusiast](https://github.com/sohamkamani) and an avid blogger. You should [follow him on Twitter](https://twitter.com/sohamkamani)

由 **Soham Kamani** 撰写，[作者](https://www.packtpub.com/books/info/authors/soham-kamani) 和在 JavaScript 生态系统方面拥有丰富经验的全栈开发人员，并在 Go 中构建大规模应用程序。他是[开源爱好者](https://github.com/sohamkamani) 和狂热的博主。你应该[在推特上关注他](https://twitter.com/sohamkamani)

### Read Next:

### 阅读下一个：

- [*Data races in Go(Golang) and how to fix them*](https://www.sohamkamani.com/golang/data-races/)
- [*Using Mutexes in Golang - A Comprehensive Tutorial With Examples*](https://www.sohamkamani.com/golang/mutex/) 
- [*Go(Golang) 中的数据竞争以及如何修复它们*](https://www.sohamkamani.com/golang/data-races/)
- [*在 Golang 中使用互斥体 - 包含示例的综合教程*](https://www.sohamkamani.com/golang/mutex/)