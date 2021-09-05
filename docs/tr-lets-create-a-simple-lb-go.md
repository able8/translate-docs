# Let's Create a Simple Load Balancer With Go

# 让我们用 Go 创建一个简单的负载均衡器

Posted 8. November 2019.   **10 min read.**

Load Balancers plays a key role in Web Architecture. They allow distributing load among a set of backends. This makes services more scalable. Also since there are multiple backends configured the service become  highly available as load balancer can pick up a working server in case  of a failure.

负载均衡器在 Web 架构中起着关键作用。它们允许在一组后端之间分配负载。这使得服务更具可扩展性。此外，由于配置了多个后端，服务变得高度可用，因为负载均衡器可以在出现故障时选择工作服务器。

After playing with professional Load Balancers like [NGINX](https://www.nginx.com/) I tried creating a simple Load Balancer for fun. I implemented it using [Golang](https://golang.org/). Go is a modern language which supports concurrency as a first-class  citizen. Go has a rich standard library which allows writing high-performance  applications with fewer lines of codes. It also produces a statically  linked single binary for easy distributions.

在玩过像 [NGINX](https://www.nginx.com/) 这样的专业负载均衡器之后，我尝试创建一个简单的负载均衡器来获得乐趣。我使用 [Golang](https://golang.org/) 实现了它。 Go 是一种现代语言，它作为一等公民支持并发。 Go 有一个丰富的标准库，它允许用更少的代码行编写高性能应用程序。它还生成一个静态链接的单个二进制文件，以便于分发。

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#how-does-our-simple-load-balancer-work)How does our simple load balancer work

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#how-does-our-simple-load-balancer-work)我们的简单负载均衡器是如何工作的

Load Balancers have different strategies for distributing the load across a set of backends.

负载均衡器有不同的策略来跨一组后端分配负载。

For example,

例如，

- **Round Robin** - Distribute load equally, assumes all backends have the same processing power
- **Weighted Round Robin** - Additional weights can be given considering the backend's processing power
- **Least Connections** - Load is distributed to the servers with least active connections

- **Round Robin** - 平均分配负载，假设所有后端具有相同的处理能力
- **加权循环** - 考虑到后端的处理能力，可以给予额外的权重
- **最少连接** - 负载分配到活动连接最少的服务器

For our simple load balancer, we would try implementing the simplest one among these methods, **Round Robin**.

对于我们的简单负载均衡器，我们将尝试实现这些方法中最简单的一种，**Round Robin**。



![A Round Robin Load Balancer](https://kasvith.me/assets/static/lb-archi.0b1c2c4.b3e35c7510dc44451088756d14739161.png)A Round Robin Load Balancer



## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#round-robin-selection)Round Robin Selection

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#round-robin-selection)循环选择

Round Robin is simple in terms. It gives equal opportunities for workers to perform tasks in turns.

Round Robin 的术语很简单。它为工人提供了轮流执行任务的平等机会。



![Round Robin Selection on incoming requests](https://kasvith.me/assets/static/lb-rr.a901b4a.3b34a78610b7c1e2d22b85f0419700d2.png)Round Robin Selection on incoming requests



As shown in the figure about this happens cyclically. But we can't *directly* use that aren't we?

如图所示，这种情况是周期性发生的。但是我们不能*直接*使用，不是吗？

**What if a backend is down?** We probably don't want to route traffic there. So this cannot be directly used unless we put some conditions on it. We need to **route traffic only to backends which are up and running**.

**如果后端出现故障怎么办？** 我们可能不想将流量路由到那里。所以这个不能直接使用，除非我们给它加一些条件。我们需要**仅将流量路由到已启动并正在运行的后端**。

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#lets-define-some-structs)Lets define some structs

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#lets-define-some-structs)让我们定义一些结构

After revising the plan, we know now we want a way to track all the details about a Backend. We need to track whether it's alive or dead and also keep track of the Url as well.

修改计划后，我们现在知道我们需要一种方法来跟踪有关后端的所有详细信息。我们需要跟踪它是活着还是死了，还需要跟踪 Url。

We can simply define a struct like this to hold our backends.

我们可以简单地定义一个像这样的结构来保存我们的后端。

```go
type Backend struct {
  URL          *url.URL
  Alive        bool
  mux          sync.RWMutex
  ReverseProxy *httputil.ReverseProxy
}
```


Don't worry **I will reason about the fields in the `Backend`**.

别担心 **我会推理 `Backend` 中的字段**。

Now we need a way to track all the backends in our load balancer, for that we can simply use a Slice. And also a counter variable. We can define it as **`ServerPool`**

现在我们需要一种方法来跟踪负载均衡器中的所有后端，为此我们可以简单地使用 Slice。还有一个计数器变量。我们可以将其定义为 **`ServerPool`**

```go
type ServerPool struct {
  backends []*Backend
  current  uint64
}
```


## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#use-of-the-reverseproxy)Use of the ReverseProxy

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#use-of-the-reverseproxy)ReverseProxy的使用

As we already identified, the sole purpose of the load balancer is to route traffic to different backends and return the results to the  original client.

正如我们已经确定的，负载均衡器的唯一目的是将流量路由到不同的后端并将结果返回给原始客户端。

According to Go's documentation,

根据 Go 的文档，

> ReverseProxy is an HTTP Handler that takes an incoming request and  sends it to another server, proxying the response back to the client.

> ReverseProxy 是一个 HTTP 处理程序，它接收传入的请求并将其发送到另一台服务器，将响应代理回客户端。

**Which is exactly what we want**. There is no need to reinvent the wheel. We can simply relay our original requests through the `ReverseProxy`.

**这正是我们想要的**。没有必要重新发明轮子。我们可以简单地通过 `ReverseProxy` 中继我们的原始请求。

```go
u, _ := url.Parse("http://localhost:8080")
rp := httputil.NewSingleHostReverseProxy(u)
  
// initialize your server and add this as handler
http.HandlerFunc(rp.ServeHTTP)
```


With `httputil.NewSingleHostReverseProxy(url)` we can initialize a reverse proxy which would relay requests to the passed `url`. In the above example, all the requests are now passed to localhost:8080 and the results are sent back to the original client. You can find more examples here.

使用`httputil.NewSingleHostReverseProxy(url)`，我们可以初始化一个反向代理，它将请求中继到传递的`url`。在上面的例子中，现在所有的请求都被传递到 localhost:8080 并将结果发送回原始客户端。您可以在此处找到更多示例。

If we take a look at ServeHTTP method signature, it has the signature of an HTTP handler, that's why we could pass it to the `HandlerFunc` in `http`.

如果我们看一下 ServeHTTP 方法签名，它有一个 HTTP 处理程序的签名，这就是为什么我们可以将它传递给 `http` 中的 `HandlerFunc`。

You can find more examples in [docs](https://golang.org/pkg/net/http/httputil/#ReverseProxy). 

您可以在 [docs](https://golang.org/pkg/net/http/httputil/#ReverseProxy) 中找到更多示例。

For our simple load balancer we could initiate the `ReverseProxy` with the associated `URL` in the `Backend`, so that `ReverseProxy` will route our requests to the `URL`.

对于我们的简单负载均衡器，我们可以在“后端”中使用关联的“URL”启动“ReverseProxy”，以便“ReverseProxy”将我们的请求路由到“URL”。

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#selection-process)Selection Process

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#selection-process)选择流程

We need to **skip dead backends** during the next pick. But to do anything we need a way to count.

我们需要在下一次选择期间**跳过死后端**。但是要做任何事情，我们都需要一种计算方法。

Multiple clients will connect to the load balancer and when each of  them requests a next peer to pass the traffic on race conditions could  occur. To prevent it we could lock the `ServerPool` with a `mutex`. But that would be an overkill, besides we don't want to lock the ServerPool at all. We just want to increase the counter by one

多个客户端将连接到负载均衡器，并且当每个客户端请求下一个对等方在竞争条件下传递流量时，可能会发生。为了防止它，我们可以用一个 `mutex` 锁定 `ServerPool`。但这将是一种矫枉过正，此外我们根本不想锁定 ServerPool。我们只想将计数器加一

To meet that requirement, the ideal solution is to make this increment atomically. And Go supports that well via `atomic` package.

为了满足该要求，理想的解决方案是原子地进行此增量。 Go 通过 `atomic` 包很好地支持这一点。

```go
func (s *ServerPool) NextIndex() int {
  return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}
```


In here, we are increasing the current value by one atomically and  returns the index by modding with the length of the slice. Which means the value always will be between 0 and length of the slice. In the end, we are interested in a particular index, not the total  count.

在这里，我们以原子方式将当前值增加 1，并通过修改切片的长度来返回索引。这意味着该值始终介于 0 和切片长度之间。最后，我们对特定索引感兴趣，而不是总数。

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#picking-up-an-alive-backend)Picking up an alive backend.

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#picking-up-an-alive-backend)选择一个活着的后端。

We already know that our requests are routed in a cycle for each backend. All we have to skip dead ones, that's it.

我们已经知道我们的请求在每个后端的循环中路由。我们所要做的就是跳过死者，仅此而已。

`GetNext()` always return a value  that's capped between 0 and the length of the slice. At any point, we  get a next peer and if it's not alive we would have to search through  the slice in a cycle.



`GetNext()` 总是返回一个上限在 0 和切片长度之间的值。在任何时候，我们都会得到下一个节点，如果它不存在，我们将不得不在一个循环中搜索切片。



![Traverse the slice as a cycle](https://kasvith.me/assets/static/lb-slice-traverse.93ea44d.9c17a7fc56f9bf6c9e4583c29127aa55.png)Traverse the slice as a cycle



As shown in the figure above, we want to traverse from next to the entire list, which can be done simply by traversing `next + length` But to pick an index, we want to cap it between slice length. It can be easily done with modding operation.

如上图所示，我们要从 next 到整个列表遍历，可以简单地通过遍历 `next + length` 来完成，但是要选择一个索引，我们希望将它限制在切片长度之间。它可以通过修改操作轻松完成。

After we find a working backend through the search, we mark it as the current one.

在我们通过搜索找到一个可用的后端后，我们将其标记为当前后端。

Below you can see the code for the above operation.

您可以在下面看到上述操作的代码。

```go
// GetNextPeer returns next active peer to take a connection
func (s *ServerPool) GetNextPeer() *Backend {
  // loop entire backends to find out an Alive backend
  next := s.NextIndex()
  l := len(s.backends) + next // start from next and move a full cycle
  for i := next;i < l;i++ {
    idx := i % len(s.backends) // take an index by modding with length
    // if we have an alive backend, use it and store if its not the original one
    if s.backends[idx].IsAlive() {
      if i != next {
        atomic.StoreUint64(&s.current, uint64(idx)) // mark the current one
      }
      return s.backends[idx]
    }
  }
  return nil
}
```


## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#avoid-race-conditions-in-backend-struct)Avoid Race Conditions in Backend struct

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#avoid-race-conditions-in-backend-struct)避免后端结构中的竞争条件

There is a serious issue we need to consider. Our `Backend` structure has a variable which could be modified or accessed by different goroutines same time.

我们需要考虑一个严重的问题。我们的 `Backend` 结构有一个变量，可以同时被不同的 goroutine 修改或访问。

We know there would be more goroutines reading from this rather than writing to it. So we have picked `RWMutex` to serialize the access to the `Alive`.

我们知道会有更多的 goroutine 从中读取而不是写入。所以我们选择了 `RWMutex` 来序列化对 `Alive` 的访问。

```go
// SetAlive for this backend
func (b *Backend) SetAlive(alive bool) {
  b.mux.Lock()
  b.Alive = alive
  b.mux.Unlock()
}

// IsAlive returns true when backend is alive
func (b *Backend) IsAlive() (alive bool) {
  b.mux.RLock()
  alive = b.Alive
  b.mux.RUnlock()
  return
}
```


## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#lets-load-balance-requests)Lets load balance requests

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#lets-load-balance-requests)让负载均衡请求

With all the background we created, we can formulate the following simple method to load balance our requests. It will only fail when our all backends are offline.

有了我们创建的所有背景，我们可以制定以下简单的方法来平衡我们的请求。只有当我们所有的后端都离线时它才会失败。

```go
// lb load balances the incoming request
func lb(w http.ResponseWriter, r *http.Request) {
  peer := serverPool.GetNextPeer()
  if peer != nil {
    peer.ReverseProxy.ServeHTTP(w, r)
    return
  }
  http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
```


This method can be simply passed as a `HandlerFunc` to the http server.

这个方法可以简单地作为 `HandlerFunc` 传递给 http 服务器。

```go
server := http.Server{
  Addr:    fmt.Sprintf(":%d", port),
  Handler: http.HandlerFunc(lb),
}
```




## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#route-traffic-only-to-healthy-backends)Route traffic only to healthy backends

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#route-traffic-only-to-healthy-backends)仅将流量路由到健康的后端

Our current `lb` has a serious  issue. We don't know if a backend is healthy or not. To know this we  have to try out a backend and check whether it is alive.

我们当前的“lb”有一个严重的问题。我们不知道后端是否健康。要知道这一点，我们必须尝试一个后端并检查它是否还活着。

We can do this in two ways,

我们可以通过两种方式做到这一点

- **Active**: While performing the current request, we find the selected backend is unresponsive, mark it as down.
- **Passive**: We can ping backends on fixed intervals and check status

- **Active**：在执行当前请求时，我们发现所选后端没有响应，将其标记为关闭。
- **被动**：我们可以在固定的时间间隔 ping 后端并检查状态

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#actively-checking-for-healthy-backends)Actively checking for healthy backends

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#actively-checking-for-healthy-backends)主动检查后端是否健康

`ReverseProxy` triggers a callback function, `ErrorHandler` on any error. We can use that to detect any failure. Here is the implementation

`ReverseProxy` 会在任何错误时触发回调函数 `ErrorHandler`。我们可以使用它来检测任何故障。这是实现

```go
proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
  log.Printf("[%s] %s\n", serverUrl.Host, e.Error())
  retries := GetRetryFromContext(request)
  if retries < 3 {
    select {
      case <-time.After(10 * time.Millisecond):
        ctx := context.WithValue(request.Context(), Retry, retries+1)
        proxy.ServeHTTP(writer, request.WithContext(ctx))
      }
      return
    }

  // after 3 retries, mark this backend as down
  serverPool.MarkBackendStatus(serverUrl, false)

  // if the same request routing for few attempts with different backends, increase the count
  attempts := GetAttemptsFromContext(request)
  log.Printf("%s(%s) Attempting retry %d\n", request.RemoteAddr, request.URL.Path, attempts)
  ctx := context.WithValue(request.Context(), Attempts, attempts+1)
  lb(writer, request.WithContext(ctx))
}
```


In here we leverage the power of closures to design this error  handler. It allows us to capture outer variables like server url into  our method. It will check for existing retry count and if it is less than 3, we  again send the same request to the same backend. The reason behind this is due to temporary errors the server may reject  your requests and it may be available after a short delay(possibly the  server ran out of sockets to accept more clients). So we have put a  timer to delay the retry for around 10 milliseconds. We increases the  retry count with every request.

在这里，我们利用闭包的力量来设计这个错误处理程序。它允许我们将服务器 url 等外部变量捕获到我们的方法中。它将检查现有的重试计数，如果它小于 3，我们再次将相同的请求发送到同一个后端。这背后的原因是由于服务器可能会拒绝您的请求的临时错误，并且可能会在短暂延迟后可用（可能服务器用尽了套接字以接受更多客户端）。所以我们设置了一个计时器来延迟重试大约 10 毫秒。我们增加每个请求的重试次数。

After every retry failed, we mark this backend as down.

每次重试失败后，我们将此后端标记为关闭。

Next thing we want to do is attempting a new backend to the same  request. We do it by keeping a count of the attempts using the context  package. After increasing the attempt count, we pass it back to `lb` to pick a new peer to process the request.

我们要做的下一件事是尝试对同一请求使用新的后端。我们通过使用 context 包记录尝试次数来做到这一点。在增加尝试次数后，我们将它传回给 `lb` 以选择一个新的对等点来处理请求。

Now we can't do this indefinitely, thus we need to check from `lb` whether the maximum attempts already taken before processing the request further.

现在我们不能无限期地这样做，因此我们需要从“lb”检查在进一步处理请求之前是否已经进行了最大尝试。

We can simply get the attempt count from the request and if it has exceeded the max count, eliminate the request.

我们可以简单地从请求中获取尝试次数，如果超过最大次数，则消除请求。

```go
// lb load balances the incoming request
func lb(w http.ResponseWriter, r *http.Request) {
  attempts := GetAttemptsFromContext(r)
  if attempts > 3 {
    log.Printf("%s(%s) Max attempts reached, terminating\n", r.RemoteAddr, r.URL.Path)
    http.Error(w, "Service not available", http.StatusServiceUnavailable)
    return
  }

  peer := serverPool.GetNextPeer()
  if peer != nil {
    peer.ReverseProxy.ServeHTTP(w, r)
    return
  }
  http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
```


This implementation is recursive.

这个实现是递归的。

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#use-of-context)Use of context

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#use-of-context)使用上下文

`context` package allows you to store useful data in an Http request. We heavily utilized this to track request specific data such as Attempt count and Retry count.

`context` 包允许您在 Http 请求中存储有用的数据。我们大量使用它来跟踪请求特定数据，例如尝试计数和重试计数。

First, we need to specify keys for the context. It is recommended to  use non-colliding integer keys rather than strings. Go provides `iota` keyword to implement constants incrementally, each containing a unique value. That is a perfect solution defining integer keys.

首先，我们需要为上下文指定键。建议使用非冲突整数键而不是字符串。 Go 提供了 `iota` 关键字来增量地实现常量，每个常量都包含一个唯一的值。这是定义整数键的完美解决方案。

```go
const (
  Attempts int = iota
  Retry
)
```


Then we can retrieve the value as usually we do with a HashMap like  follows. The default return value may depend on the use case.

然后我们可以像往常一样使用 HashMap 检索值，如下所示。默认返回值可能取决于用例。

```go
// GetAttemptsFromContext returns the attempts for request
func GetRetryFromContext(r *http.Request) int {
  if retry, ok := r.Context().Value(Retry).(int);ok {
    return retry
  }
  return 0
}
```


## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#passive-health-checks)Passive health checks 

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#passive-health-checks)被动健康检查

Passive health checks allow to recover dead backends or identify  them. We ping the backends with fixed intervals to check their status.

被动健康检查允许恢复死后端或识别它们。我们以固定的时间间隔 ping 后端以检查它们的状态。

To ping, we try to establish a TCP connection. If the backend  responses, we mark it as alive. This method can be changed to call a  specific endpoint like `/status` if you like. Make sure to close the connection once it established to reduce the additional load in the server. Otherwise, it will try to maintain the connection and it would run out of resources eventually.

为了 ping，我们尝试建立 TCP 连接。如果后端响应，我们将其标记为活着。如果您愿意，可以更改此方法以调用特定端点，例如 `/status`。确保在建立连接后关闭连接以减少服务器中的额外负载。否则，它会尝试保持连接，最终会耗尽资源。

```go
// isAlive checks whether a backend is Alive by establishing a TCP connection
func isBackendAlive(u *url.URL) bool {
  timeout := 2 * time.Second
  conn, err := net.DialTimeout("tcp", u.Host, timeout)
  if err != nil {
    log.Println("Site unreachable, error: ", err)
    return false
  }
  _ = conn.Close() // close it, we dont need to maintain this connection
  return true
}
```


Now we can iterate the servers and mark their status like follows,

现在我们可以迭代服务器并标记它们的状态，如下所示，

```go
// HealthCheck pings the backends and update the status
func (s *ServerPool) HealthCheck() {
  for _, b := range s.backends {
    status := "up"
    alive := isBackendAlive(b.URL)
    b.SetAlive(alive)
    if !alive {
      status = "down"
    }
    log.Printf("%s [%s]\n", b.URL, status)
  }
}
```


To run this periodically we can start a timer in Go. Once a timer created it allows you to listen for the event using a channel.

为了定期运行，我们可以在 Go 中启动一个计时器。创建计时器后，它允许您使用通道侦听事件。

```go
// healthCheck runs a routine for check status of the backends every 20 secs
func healthCheck() {
  t := time.NewTicker(time.Second * 20)
  for {
    select {
    case <-t.C:
      log.Println("Starting health check...")
      serverPool.HealthCheck()
      log.Println("Health check completed")
    }
  }
}
```


In the above snippet, `<-t.C` channel will return a value per 20s. `select` allows to detect this event. `select` waits until at least one case statement could be executed if there is no `default` case.

在上面的代码片段中，`<-t.C` 通道将每 20 秒返回一个值。 `select` 允许检测这个事件。如果没有 `default` case，`select` 会等待至少一个 case 语句可以被执行。

Finally, run this in a separate goroutine.

最后，在单独的 goroutine 中运行它。

```go
go healthCheck()
```


## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#conclusion)Conclusion

## [#](https://kasvith.me/posts/lets-create-a-simple-lb-go/#conclusion)结论

We covered a lot of stuff in this article.

我们在本文中介绍了很多内容。

- Round Robin Selection
- ReverseProxy from the standard library
- Mutexes
- Atomic Operations
- Closures
- Callbacks
- Select Operation

- 循环选择
- 来自标准库的 ReverseProxy
- 互斥体
- 原子操作
- 关闭
- 回调
- 选择操作

There is a lot we can do to improve our tiny load balancer.

我们可以做很多事情来改进我们的微型负载均衡器。

For example,

例如，

- Use a heap for sort out alive backends to reduce search surface
- Collect statistics
- Implement weighted round-robin/least connections
- Add support for a configuration file

- 使用堆整理活着的后端以减少搜索面
- 收集统计数据
- 实施加权循环/最少连接
- 添加对配置文件的支持

etc.

等等。

You can find the source code to repository [here](https://github.com/kasvith/simplelb/).

您可以在 [此处](https://github.com/kasvith/simplelb/) 中找到存储库的源代码。

Thank you for reading this article 😄 

谢谢你阅读这篇文章😄

