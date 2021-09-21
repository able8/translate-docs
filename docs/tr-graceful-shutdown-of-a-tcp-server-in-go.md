# Graceful shutdown of a TCP server in Go

# 在 Go 中优雅地关闭 TCP 服务器

January 21, 2020

This post is going to discuss how to gracefully shut down a TCP server in Go. While servers typically never stop running (until the process is killed), in some scenarios - e.g. in tests - it's useful to shut them down in an orderly way.

这篇文章将讨论如何在 Go 中优雅地关闭 TCP 服务器。虽然服务器通常永远不会停止运行（直到进程被终止），但在某些情况下 - 例如在测试中 - 以有序的方式关闭它们很有用。

## High-level structure of TCP servers in Go

## Go 中 TCP 服务器的高层结构

Let's start with a quick review of the high-level structure of TCP servers implemented in Go. Go provides some convenient abstractions on top of sockets. Here's pseudo-code for a typical server:

让我们先快速回顾一下 Go 中实现的 TCP 服务器的高级结构。 Go 在套接字之上提供了一些方便的抽象。这是典型服务器的伪代码：

```go
listener := net.Listen("tcp", ... address ...)
for {
  conn := listener.Accept()
  go handler(conn)
}
```

Where `handler` is a blocking function that waits for commands from the client, does the required processing, and sends responses back.

其中`handler` 是一个阻塞函数，它等待来自客户端的命令，执行所需的处理，并将响应发送回。

Given this structure, we should clarify what we mean by "shutting a server down". It seems like there are two distinct functionalities a server is performing at any given time:

鉴于这种结构，我们应该澄清“关闭服务器”的含义。似乎服务器在任何给定时间执行两种不同的功能：

1. It listens for new connections
2. It handles existing connections

3. 监听新连接
4. 它处理现有的连接

It's clear that we can stop listening for new connections, thus handling (1); but what about existing connections?

很明显，我们可以停止监听新连接，从而处理（1）；但是现有的连接呢？

Unfortunately, there's no easy answer here. The TCP protocol is too low level to resolve this question conclusively. If we want to design a widely applicable solution, we have to be conservative. Specifically, the safest approach is for the shutting down server to wait for clients to close their connections. This is the approach we'll examine initially.

不幸的是，这里没有简单的答案。 TCP 协议级别太低，无法最终解决这个问题。如果我们想设计一个广泛适用的解决方案，我们必须保守。具体来说，最安全的方法是关闭服务器等待客户端关闭它们的连接。这是我们将首先检查的方法。

## Step 1: waiting for client connections to shut down

## 第一步：等待客户端连接关闭

In this solution, we're going to explicitly shut down the listener (stop accepting new connections), but will wait for clients to end *their* connections. This is a conservative approach, but it works very well in many scenarios where server shutdown is actually needed - such as tests. In a test, it's not hard to arrange for all the clients to close their connections before expecting the server to shut down.

在这个解决方案中，我们将显式关闭侦听器（停止接受新连接），但将等待客户端结束*他们的*连接。这是一种保守的方法，但它在许多实际需要关闭服务器的场景中非常有效——例如测试。在测试中，安排所有客户端在期望服务器关闭之前关闭它们的连接并不难。

I'll be presenting the code piece by piece, but the full runnable code sample is [available here](https://github.com/eliben/code-for-blog/blob/master/2020/tcp-server-shutdown/shutdown1/shutdown1.go). Let's start with the server type and the constructor:

我将逐段呈现代码，但完整的可运行代码示例在 [此处可用](https://github.com/eliben/code-for-blog/blob/master/2020/tcp-server-shutdown/shutdown1/shutdown1.go)。让我们从服务器类型和构造函数开始：

```go
type Server struct {
  listener net.Listener
  quit     chan interface{}
  wg       sync.WaitGroup
}

func NewServer(addr string) *Server {
  s := &Server{
    quit: make(chan interface{}),
  }
  l, err := net.Listen("tcp", addr)
  if err != nil {
    log.Fatal(err)
  }
  s.listener = l
  s.wg.Add(1)
  go s.serve()
  return s
}
```

`NewServer` creates a new `Server` that listens for new connections in a background goroutine. In addition to a `net.Listener`, the `Server` struct contains a channel that's used to signal shutdown and a wait group to wait until all the server's goroutines are actually done.

`NewServer` 创建一个新的 `Server` 来监听后台 goroutine 中的新连接。除了 `net.Listener` 之外，`Server` 结构还包含一个用于发出关闭信号的通道和一个等待组，用于等待服务器的所有 goroutine 实际完成。

Here's the `serve` method the constructor invokes:

这是构造函数调用的 `serve` 方法：

```go
func (s *Server) serve() {
  defer s.wg.Done()

  for {
    conn, err := s.listener.Accept()
    if err != nil {
      select {
      case <-s.quit:
        return
      default:
        log.Println("accept error", err)
      }
    } else {
      s.wg.Add(1)
      go func() {
        s.handleConection(conn)
        s.wg.Done()
      }()
    }
  }
}
```

It's a standard `Accept` loop, except for the `select`. What this `select` does is check (in a non-blocking way) if there's an event (such as a send or a close) on the `s.quit` channel when `Accept` errors out. If there is, it means the error is caused by us closing the listener, and `serve` returns quietly. If `Accept` returns without errors, we run a connection handler [[1\]](https://eli.thegreenplace.net/2020/graceful-shutdown-of-a-tcp-server-in-go/#id2).

这是一个标准的“接受”循环，除了“选择”。这个 `select` 的作用是检查（以非阻塞方式）当 `Accept` 出错时 `s.quit` 通道上是否有事件（例如发送或关闭）。如果有，则说明错误是我们关闭监听器引起的，并且 `serve` 安静地返回。如果 `Accept` 没有错误返回，我们运行一个连接处理程序 [[1\]](https://eli.thegreenplace.net/2020/graceful-shutdown-of-a-tcp-server-in-go/#id2)。

Here's the `Stop` method that tells the server to shut down gracefully:

这是告诉服务器正常关闭的 `Stop` 方法：

```go
func (s *Server) Stop() {
  close(s.quit)
  s.listener.Close()
  s.wg.Wait()
}
```

It starts by closing the `s.quit` channel. Then it closes the listener. This will cause the `Accept` call in `serve` to return an error. Since `s.quit` is already closed at this point, `serve` will return.

它首先关闭 `s.quit` 通道。然后它关闭侦听器。这将导致 `serve` 中的 `Accept` 调用返回错误。由于此时 `s.quit` 已经关闭，`serve` 将返回。

The last line in `Stop` is waiting on `s.wg`, which is also critical. Note that `serve` notifies the wait group that it's done on return. But this is not the only goroutine we're waiting for. Each call to `handleConnection` is wrapped by a `wg` add/done pair as well. Therefore, `Stop` will block until all the handlers have returned, *and* `serve` stopped accepting new clients. This is a safe shutdown point.

`Stop` 中的最后一行正在等待 `s.wg`，这也很关键。请注意，`serve` 通知等待组它在返回时已完成。但这并不是我们等待的唯一 goroutine。每个对 `handleConnection` 的调用也由一个 `wg` add/done 对包装。因此，`Stop` 将阻塞，直到所有处理程序都返回，*并且*`serve` 停止接受新客户端。这是一个安全的关闭点。

For completeness, here's is `handleConnection`; the one here just reads client data and logs it, without sending anything back. Naturally, this part of the code will be different for each server:

为了完整起见，这里是`handleConnection`；这里的那个只是读取客户端数据并记录下来，不发回任何东西。自然，这部分代码对于每个服务器都会有所不同：

```go
func (s *Server) handleConection(conn net.Conn) {
  defer conn.Close()
  buf := make([]byte, 2048)
  for {
    n, err := conn.Read(buf)
    if err != nil && err != io.EOF {
      log.Println("read error", err)
      return
    }
    if n == 0 {
      return
    }
    log.Printf("received from %v: %s", conn.RemoteAddr(), string(buf[:n]))
  }
}
```

Using this server is simple:

使用这个服务器很简单：

```go
s := NewServer(addr)
// do whatever here...
s.Stop()
```

Recall that `NewServer` returns a server but doesn't block. `s.Stop` does block, however. In tests, what you'd do for graceful shutdown is:

回想一下 `NewServer` 返回一个服务器但不阻塞。然而，`s.Stop` 确实会阻塞。在测试中，您为正常关机所做的是：

1. Make sure all clients interacting with the server have closed their connections.
2. Wait for `s.Stop` to return.

1. 确保所有与服务器交互的客户端都关闭了它们的连接。
2. 等待`s.Stop` 返回。

## Step 2: actively closing open client connections

## 第二步：主动关闭打开的客户端连接

In step 1, we expected all clients to close their connections before declaring the shutdown process successful. Here we'll look at a more aggressive approach, where on `Stop()` the server will actively attempt to close open client connections. I'll present a technique that's both simple and robust first, at the cost of some performance. After that, we'll discuss some alternatives.

在步骤 1 中，我们希望所有客户端在宣布关闭过程成功之前关闭它们的连接。在这里，我们将研究一种更激进的方法，在“Stop()”上，服务器将主动尝试关闭打开的客户端连接。我将首先介绍一种既简单又健壮的技术，但要以牺牲一些性能为代价。之后，我们将讨论一些替代方案。

The full code for this step is [available too](https://github.com/eliben/code-for-blog/blob/master/2020/tcp-server-shutdown/shutdown2/shutdown2.go). It's identical to step 1 except the code of `handleConection`:

此步骤的完整代码[也可用](https://github.com/eliben/code-for-blog/blob/master/2020/tcp-server-shutdown/shutdown2/shutdown2.go)。除了`handleConection`的代码外，它与步骤1相同：

```go
func (s *Server) handleConection(conn net.Conn) {
  defer conn.Close()
  buf := make([]byte, 2048)
ReadLoop:
  for {
    select {
    case <-s.quit:
      return
    default:
      conn.SetDeadline(time.Now().Add(200 * time.Millisecond))
      n, err := conn.Read(buf)
      if err != nil {
        if opErr, ok := err.(*net.OpError);ok && opErr.Timeout() {
          continue ReadLoop
        } else if err != io.EOF {
          log.Println("read error", err)
          return
        }
      }
      if n == 0 {
        return
      }
      log.Printf("received from %v: %s", conn.RemoteAddr(), string(buf[:n]))
    }
  }
}
```

This handler sets a deadline on each socket read. The deadline duration here is 200 ms, but it could be set to anything else that makes sense for your specific application. If a read returns with a timeout, it means the client has been idle for the timeout duration and the connection could be safe to close. So each iteration of the loop checks for `s.quit` and returns if there's an event there.

此处理程序为每个套接字读取设置最后期限。此处的截止期限为 200 毫秒，但可以将其设置为对您的特定应用程序有意义的任何其他值。如果读取返回超时，则意味着客户端在超时时间内一直处于空闲状态，并且可以安全地关闭连接。所以循环的每次迭代都会检查 `s.quit` 并返回是否有事件。

This approach is robust, in the sense that we're (most likely) not going to close connections while the client is actively sending something. It's also simple, since it limits all the extra logic to `handleConnection`. 

这种方法是健壮的，因为我们（很可能）不会在客户端主动发送某些东西时关闭连接。它也很简单，因为它将所有额外的逻辑限制在 `handleConnection` 中。

There's a performance cost, of course. First, a `conn.Read` call is issued once every 200 ms, which is slightly slower than a single blocking call; I'd say this is negligible, though. More seriously, every `Stop` request will be delayed by 200 ms. This is probably OK in most scenarios where we want to shut down a server, but the deadline can be tweaked to fit specific protocol needs.

当然，有性能成本。首先，每 200 毫秒发出一次 `conn.Read` 调用，这比单个阻塞调用稍慢；不过，我会说这是微不足道的。更严重的是，每个 `Stop` 请求都会延迟 200 毫秒。在我们想要关闭服务器的大多数情况下，这可能没问题，但可以调整截止日期以满足特定的协议需求。

An alternative to this design would be to keep track of all the open connections outside `handleConection`, and force-close them when `Stop` is called. This would likely be more efficient, at the cost of implementation complexity and some lack of robustness. Such a `Stop` could easily close connections while clients are actively sending data, resulting in client errors.

这种设计的另一种方法是跟踪所有在 `handleConection` 之外打开的连接，并在调用 `Stop` 时强制关闭它们。这可能会更有效，但代价是实现复杂性和缺乏稳健性。当客户端主动发送数据时，这样的 `Stop` 很容易关闭连接，从而导致客户端错误。

For inspiration on the right path to take, we can look at the stdlib's `http.Server.Shutdown` method, which is documented as follows:

为了寻找正确路径的灵感，我们可以查看 stdlib 的 `http.Server.Shutdown` 方法，该方法记录如下：

> Shutdown gracefully shuts down the server without interrupting any active connections. Shutdown works by first closing all open listeners, then closing all idle connections, and then waiting indefinitely for connections to return to idle and then shut down

> 关闭可以正常关闭服务器，而不会中断任何活动连接。关闭的工作原理是首先关闭所有打开的侦听器，然后关闭所有空闲连接，然后无限期地等待连接返回空闲状态，然后关闭

What does "idle" mean here? Roughly that the client hasn't sent any requests for some period of time. The HTTP server has advantage over a generic TCP server, because it's a higher-level protocol, so it knows the client communication pattern. In different protocols, different shutdown strategies may make sense.

这里的“空闲”是什么意思？粗略地说，客户端在一段时间内没有发送任何请求。 HTTP 服务器比一般的 TCP 服务器有优势，因为它是一个更高级别的协议，所以它知道客户端通信模式。在不同的协议中，不同的关闭策略可能是有意义的。

A different example is a protocol where the server initiates messages, or at least some of them. For example, a given connection may be in a state where the client is waiting for the server to send some event. It's usually safe for the server to close this connection on shutdown without waiting for anything.

另一个示例是服务器发起消息或至少其中一些消息的协议。例如，给定的连接可能处于客户端等待服务器发送某个事件的状态。服务器在关闭时关闭此连接通常是安全的，而无需等待任何事情。

## Conclusion

##  结论

I would summarize this post with two general guidelines:

我将用两个一般准则总结这篇文章：

1. Try to make shutdowns as safe as possible
2. Think of the higher-level protocol

1. 尝试使关机尽可能安全
2. 考虑上层协议

I typically encounter the need to shut down a TCP server while writing tests. I want each test to be self-contained and clean up after itself, including all the client-server connections and listening servers. For this scenario, step 1 works very well. Once all client connections have been closed, `Server.Stop` will return without any delays. 

我通常在编写测试时遇到需要关闭 TCP 服务器的情况。我希望每个测试都是独立的并在其自身之后进行清理，包括所有客户端-服务器连接和侦听服务器。对于这种情况，第 1 步非常有效。一旦所有客户端连接都关闭，`Server.Stop` 将立即返回。

