# Go Concurrency Patterns: Context

# Go 并发模式：上下文

From: https://go.dev/blog/context

 29 July 2014

## Introduction

##  介绍

In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.

在 Go 服务器中，每个传入请求都在其自己的 goroutine 中处理。请求处理程序通常会启动额外的 goroutine 来访问数据库和 RPC 服务等后端。处理请求的 goroutines 集通常需要访问特定于请求的值，例如最终用户的身份、授权令牌和请求的截止日期。当请求被取消或超时时，所有处理该请求的 goroutines 都应该快速退出，以便系统可以回收它们正在使用的任何资源。

At Google, we developed a `context` package that makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly available as [context](https://go.dev/pkg/context). This article describes how to use the package and provides a complete working example.

在 Google，我们开发了一个 `context` 包，它可以轻松地将请求范围的值、取消信号和截止日期跨 API 边界传递给处理请求所涉及的所有 goroutine。该软件包以 [context](https://go.dev/pkg/context) 的形式公开提供。本文介绍了如何使用该包并提供了一个完整的工作示例。

## Context

##  语境

The core of the `context` package is the `Context` type:

`context` 包的核心是 `Context` 类型：

```
// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries.Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

(This description is condensed; the [godoc](https://go.dev/pkg/context) is authoritative.)

（此描述是浓缩的；[godoc](https://go.dev/pkg/context)是权威的。)

The `Done` method returns a channel that acts as a cancellation signal to functions running on behalf of the `Context`: when the channel is closed, the functions should abandon their work and return. The `Err` method returns an error indicating why the `Context` was canceled. The [Pipelines and Cancellation](https://go.dev/blog/pipelines) article discusses the `Done` channel idiom in more detail.

`Done` 方法返回一个通道，该通道充当代表 `Context` 运行的函数的取消信号：当通道关闭时，函数应该放弃它们的工作并返回。 `Err` 方法返回一个错误，指示取消 `Context` 的原因。 [管道和取消](https://go.dev/blog/pipelines) 文章更详细地讨论了 `Done` 通道习语。

A `Context` does *not* have a `Cancel` method for the same reason the `Done` channel is receive-only: the function receiving a cancellation signal is usually not the one that sends the signal. In particular, when a parent operation starts goroutines for sub-operations, those sub-operations should not be able to cancel the parent. Instead, the `WithCancel` function (described below) provides a way to cancel a new `Context` value.

`Context` 没有 `Cancel` 方法，原因与`Done` 通道是只接收的原因相同：接收取消信号的函数通常不是发送信号的函数。特别是，当父操作为子操作启动 goroutine 时，这些子操作不应该能够取消父操作。相反，`WithCancel` 函数（如下所述）提供了一种取消新 `Context` 值的方法。

A `Context` is safe for simultaneous use by multiple goroutines. Code can pass a single `Context` to any number of goroutines and cancel that `Context` to signal all of them.

`Context` 对于多个 goroutine 同时使用是安全的。代码可以将单个 `Context` 传递给任意数量的 goroutine，并取消该 `Context` 以向所有 goroutine 发出信号。

The `Deadline` method allows functions to determine whether they should start work at all; if too little time is left, it may not be worthwhile. Code may also use a deadline to set timeouts for I/O operations.

`Deadline` 方法允许函数确定它们是否应该开始工作；如果剩下的时间太少，可能就不值得了。代码也可以使用最后期限来设置 I/O 操作的超时。

`Value` allows a `Context` to carry request-scoped data. That data must be safe for simultaneous use by multiple goroutines.

`Value` 允许 `Context` 携带请求范围的数据。该数据必须是安全的，以便多个 goroutine 同时使用。

### Derived contexts

### 派生上下文

The `context` package provides functions to *derive* new `Context` values from existing ones. These values form a tree: when a `Context` is canceled, all `Contexts` derived from it are also canceled.

`context` 包提供了从现有值中*派生*新的 `Context` 值的函数。这些值形成了一个树：当一个 `Context` 被取消时，所有从它派生的 `Contexts` 也被取消。

`Background` is the root of any `Context` tree; it is never canceled:

`Background` 是任何`Context` 树的根；它永远不会被取消：

```
// Background returns an empty Context.It is never canceled, has no deadline,
// and has no values.Background is typically used in main, init, and tests,
// and as the top-level Context for incoming requests.
func Background() Context
```

`WithCancel` and `WithTimeout` return derived `Context` values that can be canceled sooner than the parent `Context`. The `Context` associated with an incoming request is typically canceled when the request handler returns. `WithCancel` is also useful for canceling redundant requests when using multiple replicas. `WithTimeout` is useful for setting a deadline on requests to backend servers:

`WithCancel` 和 `WithTimeout` 返回派生的 `Context` 值，这些值可以比父 `Context` 更快地取消。当请求处理程序返回时，与传入请求关联的 `Context` 通常会被取消。 `WithCancel` 对于在使用多个副本时取消冗余请求也很有用。 `WithTimeout` 对于设置对后端服务器的请求的截止日期很有用：

```
// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// A CancelFunc cancels a Context.
type CancelFunc func()

// WithTimeout returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed, cancel is called, or timeout elapses.The new
// Context's Deadline is the sooner of now+timeout and the parent's deadline, if
// any.If the timer is still running, the cancel function releases its
// resources.
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

`WithValue` provides a way to associate request-scoped values with a `Context`:

`WithValue` 提供了一种将请求范围的值与 `Context` 关联的方法：

```
// WithValue returns a copy of parent whose Value method returns val for key.
func WithValue(parent Context, key interface{}, val interface{}) Context
```

The best way to see how to use the `context` package is through a worked example.

查看如何使用 `context` 包的最佳方法是通过一个工作示例。

## Example: Google Web Search

## 示例：谷歌网页搜索

Our example is an HTTP server that handles URLs like `/search?q=golang&timeout=1s` by forwarding the query “golang” to the [Google Web Search API](https://developers.google.com/web-search/docs/) and rendering the results. The `timeout` parameter tells the server to cancel the request after that duration elapses.

我们的示例是一个 HTTP 服务器，它通过将查询“golang”转发到 [Google Web Search API](https://developers.google.com/web-search/) 来处理像`/search?q=golang&timeout=1s` 这样的 URL docs/) 并呈现结果。 `timeout` 参数告诉服务器在该持续时间过去后取消请求。

The code is split across three packages:

代码分为三个包：

- [server](https://go.dev/blog/context/server/server.go) provides the `main` function and the handler for `/search`.
- [userip](https://go.dev/blog/context/userip/userip.go) provides functions for extracting a user IP address from a request and associating it with a `Context`.
- [google](https://go.dev/blog/context/google/google.go) provides the `Search` function for sending a query to Google.

- [server](https://go.dev/blog/context/server/server.go) 提供 `main` 函数和 `/search` 的处理程序。
- [userip](https://go.dev/blog/context/userip/userip.go) 提供从请求中提取用户 IP 地址并将其与 `Context` 关联的功能。
- [google](https://go.dev/blog/context/google/google.go) 提供“搜索”功能，用于向 Google 发送查询。

### The server program

### 服务器程序

The [server](https://go.dev/blog/context/server/server.go) program handles requests like `/search?q=golang` by serving the first few Google search results for `golang`. It registers `handleSearch` to handle the `/search` endpoint. The handler creates an initial `Context` called `ctx` and arranges for it to be canceled when the handler returns. If the request includes the `timeout` URL parameter, the `Context` is canceled automatically when the timeout elapses:

[server](https://go.dev/blog/context/server/server.go) 程序通过为 `golang` 提供前几个 Google 搜索结果来处理类似 `/search?q=golang` 的请求。它注册 `handleSearch` 来处理 `/search` 端点。处理程序创建一个名为 `ctx` 的初始 `Context`，并安排在处理程序返回时取消它。如果请求包含 `timeout` URL 参数，则 `Context` 会在超时时间过后自动取消：

```
func handleSearch(w http.ResponseWriter, req *http.Request) {
    // ctx is the Context for this handler.Calling cancel closes the
    // ctx.Done channel, which is the cancellation signal for requests
    // started by this handler.
    var (
        ctx    context.Context
        cancel context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err == nil {
        // The request has a timeout, so create a context that is
        // canceled automatically when the timeout expires.
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    } else {
        ctx, cancel = context.WithCancel(context.Background())
    }
    defer cancel() // Cancel ctx as soon as handleSearch returns.
```

The handler extracts the query from the request and extracts the client’s IP address by calling on the `userip` package. The client’s IP address is needed for backend requests, so `handleSearch` attaches it to `ctx`:

处理程序从请求中提取查询，并通过调用 `userip` 包提取客户端的 IP 地址。后端请求需要客户端的 IP 地址，因此 `handleSearch` 将其附加到 `ctx`：

```
     // Check the search query.
    query := req.FormValue("q")
    if query == "" {
        http.Error(w, "no query", http.StatusBadRequest)
        return
    }

    // Store the user IP in ctx for use by code in other packages.
    userIP, err := userip.FromRequest(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ctx = userip.NewContext(ctx, userIP)
```

The handler calls `google.Search` with `ctx` and the `query`:

处理程序使用 `ctx` 和 `query` 调用 `google.Search`：

```
     // Run the Google search and print the results.
    start := time.Now()
    results, err := google.Search(ctx, query)
    elapsed := time.Since(start)
```

If the search succeeds, the handler renders the results:

如果搜索成功，则处理程序呈现结果：

```
     if err := resultsTemplate.Execute(w, struct {
        Results          google.Results
        Timeout, Elapsed time.Duration
    }{
        Results: results,
        Timeout: timeout,
        Elapsed: elapsed,
    });err != nil {
        log.Print(err)
        return
    }
```

### Package userip

### 打包用户IP

The [userip](https://go.dev/blog/context/userip/userip.go) package provides functions for extracting a user IP address from a request and associating it with a `Context`. A `Context` provides a key-value mapping, where the keys and values are both of type `interface{}`. Key types must support equality, and values must be safe for simultaneous use by multiple goroutines. Packages like `userip` hide the details of this mapping and provide strongly-typed access to a specific `Context` value.

[userip](https://go.dev/blog/context/userip/userip.go) 包提供了从请求中提取用户 IP 地址并将其与“上下文”相关联的功能。 `Context` 提供键值映射，其中键和值都是`interface{}` 类型。键类型必须支持相等，并且值必须安全地被多个 goroutine 同时使用。像 `userip` 这样的包隐藏了这个映射的细节，并提供了对特定 `Context` 值的强类型访问。

To avoid key collisions, `userip` defines an unexported type `key` and uses a value of this type as the context key:

为了避免键冲突，`userip` 定义了一个未导出的类型 `key`，并使用该类型的值作为上下文键：

```
// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// userIPkey is the context key for the user IP address.Its value of zero is
// arbitrary.If this package defined other context keys, they would have
// different integer values.
const userIPKey key = 0
```

`FromRequest` extracts a `userIP` value from an `http.Request`:

`FromRequest` 从 `http.Request` 中提取 `userIP` 值：

```
func FromRequest(req *http.Request) (net.IP, error) {
    ip, _, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
    }
```

`NewContext` returns a new `Context` that carries a provided `userIP` value:

`NewContext` 返回一个新的 `Context`，它带有提供的 `userIP` 值：

```
func NewContext(ctx context.Context, userIP net.IP) context.Context {
    return context.WithValue(ctx, userIPKey, userIP)
}
```

`FromContext` extracts a `userIP` from a `Context`:

`FromContext` 从 `Context` 中提取 `userIP`：

```
func FromContext(ctx context.Context) (net.IP, bool) {
    // ctx.Value returns nil if ctx has no value for the key;
    // the net.IP type assertion returns ok=false for nil.
    userIP, ok := ctx.Value(userIPKey).(net.IP)
    return userIP, ok
}
```

### Package google

### 打包谷歌

The [google.Search](https://go.dev/blog/context/google/google.go) function makes an HTTP request to the [Google Web Search API](https://developers.google.com/web-search/docs/) and parses the JSON-encoded result. It accepts a `Context` parameter `ctx` and returns immediately if `ctx.Done` is closed while the request is in flight.

[google.Search](https://go.dev/blog/context/google/google.go) 函数向 [Google Web Search API](https://developers.google.com/web) 发出 HTTP 请求-search/docs/) 并解析 JSON 编码的结果。它接受 `Context` 参数 `ctx` 并在请求运行时如果 `ctx.Done` 关闭则立即返回。

The Google Web Search API request includes the search query and the user IP as query parameters:

Google Web Search API 请求包括搜索查询和用户 IP 作为查询参数：

```
func Search(ctx context.Context, query string) (Results, error) {
    // Prepare the Google Search API request.
    req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
    if err != nil {
        return nil, err
    }
    q := req.URL.Query()
    q.Set("q", query)

    // If ctx is carrying the user IP address, forward it to the server.
    // Google APIs use the user IP to distinguish server-initiated requests
    // from end-user requests.
    if userIP, ok := userip.FromContext(ctx);ok {
        q.Set("userip", userIP.String())
    }
    req.URL.RawQuery = q.Encode()
```

`Search` uses a helper function, `httpDo`, to issue the HTTP request and cancel it if `ctx.Done` is closed while the request or response is being processed. `Search` passes a closure to `httpDo` handle the HTTP response:

`Search` 使用辅助函数 `httpDo` 来发出 HTTP 请求，如果在处理请求或响应时关闭 `ctx.Done` 则取消它。 `Search` 将闭包传递给 `httpDo` 处理 HTTP 响应：

```
    var results Results
    err = httpDo(ctx, req, func(resp *http.Response, err error) error {
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        // Parse the JSON search result.
        // https://developers.google.com/web-search/docs/#fonje
        var data struct {
            ResponseData struct {
                Results []struct {
                    TitleNoFormatting string
                    URL               string
                }
            }
        }
        if err := json.NewDecoder(resp.Body).Decode(&data);err != nil {
            return err
        }
        for _, res := range data.ResponseData.Results {
            results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
        }
        return nil
    })
    // httpDo waits for the closure we provided to return, so it's safe to
    // read results here.
    return results, err
```

The `httpDo` function runs the HTTP request and processes its response in a new goroutine. It cancels the request if `ctx.Done` is closed before the goroutine exits:

`httpDo` 函数运行 HTTP 请求并在新的 goroutine 中处理其响应。如果 `ctx.Done` 在 goroutine 退出之前关闭，它会取消请求：

```
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
    // Run the HTTP request in a goroutine and pass the response to f.
    c := make(chan error, 1)
    req = req.WithContext(ctx)
    go func() { c <- f(http.DefaultClient.Do(req)) }()
    select {
    case <-ctx.Done():
        <-c // Wait for f to return.
        return ctx.Err()
    case err := <-c:
        return err
    }
}
```

## Adapting code for Contexts

## 为上下文调整代码

Many server frameworks provide packages and types for carrying request-scoped values. We can define new implementations of the `Context` interface to bridge between code using existing frameworks and code that expects a `Context` parameter. 

许多服务器框架提供包和类型来承载请求范围的值。我们可以定义 `Context` 接口的新实现，以在使用现有框架的代码和需要 `Context` 参数的代码之间架起桥梁。

For example, Gorilla’s [github.com/gorilla/context](http://www.gorillatoolkit.org/pkg/context) package allows handlers to associate data with incoming requests by providing a mapping from HTTP requests to key-value pairs. In [gorilla.go](https://go.dev/blog/context/gorilla/gorilla.go), we provide a `Context` implementation whose `Value` method returns the values associated with a specific HTTP request in the Gorilla package.

例如，Gorilla 的 [github.com/gorilla/context](http://www.gorillatoolkit.org/pkg/context) 包允许处理程序通过提供从 HTTP 请求到键值对的映射来将数据与传入请求相关联。在 [gorilla.go](https://go.dev/blog/context/gorilla/gorilla.go) 中，我们提供了一个 `Context` 实现，其 `Value` 方法返回与 Gorilla 中特定 HTTP 请求关联的值包裹。

Other packages have provided cancellation support similar to `Context`. For example, [Tomb](https://godoc.org/gopkg.in/tomb.v2) provides a `Kill` method that signals cancellation by closing a `Dying` channel. `Tomb` also provides methods to wait for those goroutines to exit, similar to `sync.WaitGroup`. In [tomb.go](https://go.dev/blog/context/tomb/tomb.go), we provide a `Context` implementation that is canceled when either its parent `Context` is canceled or a provided `Tomb ` is killed.

其他包提供了类似于 `Context` 的取消支持。例如，[Tomb](https://godoc.org/gopkg.in/tomb.v2) 提供了一个 `Kill` 方法，通过关闭 `Dying` 通道来表示取消。 `Tomb` 还提供了等待这些 goroutine 退出的方法，类似于 `sync.WaitGroup`。在 [tomb.go](https://go.dev/blog/context/tomb/tomb.go) 中，我们提供了一个 `Context` 实现，当它的父 `Context` 被取消或提供的 `Tomb `被杀。

## Conclusion

##  结论

At Google, we require that Go programmers pass a `Context` parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancellation and ensures that critical values like security credentials transit Go programs properly.

在 Google，我们要求 Go 程序员在传入和传出请求之间的调用路径上将“Context”参数作为第一个参数传递给每个函数。这使得许多不同团队开发的 Go 代码能够很好地互操作。它提供了对超时和取消的简单控制，并确保安全凭证等关键值正确传输 Go 程序。

Server frameworks that want to build on `Context` should provide implementations of `Context` to bridge between their packages and those that expect a `Context` parameter. Their client libraries would then accept a `Context` from the calling code. By establishing a common interface for request-scoped data and cancellation, `Context` makes it easier for package developers to share code for creating scalable services. 

想要在 `Context` 上构建的服务器框架应该提供 `Context` 的实现，以便在它们的包和那些需要 `Context` 参数的包之间架起一座桥梁。然后，他们的客户端库将接受来自调用代码的“上下文”。通过为请求范围的数据和取消建立一个通用接口，`Context` 使包开发人员更容易共享代码以创建可扩展的服务。

