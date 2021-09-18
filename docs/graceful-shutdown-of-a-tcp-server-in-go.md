​                        [                 ![img](https://eli.thegreenplace.net/images/logosmall.png) Eli Bendersky's website            ](https://eli.thegreenplace.net/)        

- ​                    [                                                  About                     ](https://eli.thegreenplace.net/pages/about)                
- ​                    [                                                  Archives                     ](https://eli.thegreenplace.net/archives/all)                


# Graceful shutdown of a TCP server in Go

January 21, 2020

This post is going to discuss how to gracefully shut down a TCP server in Go. While servers typically never stop running (until the process is killed), in some scenarios - e.g. in tests - it's useful to shut them down in an orderly way.

## High-level structure of TCP servers in Go

Let's start with a quick review of the high-level structure of TCP servers implemented in Go. Go provides some convenient abstractions on top of sockets. Here's pseudo-code for a typical server:

```
listener := net.Listen("tcp", ... address ...)
for {
  conn := listener.Accept()
  go handler(conn)
}
```

Where `handler` is a blocking function that waits for commands from the client, does the required processing, and sends responses back.

Given this structure, we should clarify what we mean by "shutting a server down". It seems like there are two distinct functionalities a server is performing at any given time:

1. It listens for new connections
2. It handles existing connections

It's clear that we can stop listening for new connections, thus handling (1); but what about existing connections?

Unfortunately, there's no easy answer here. The TCP protocol is too low level to resolve this question conclusively. If we want to design a widely applicable solution, we have to be conservative. Specifically, the safest approach is for the shutting down server to wait for clients to close their connections. This is the approach we'll examine initially.

## Step 1: waiting for client connections to shut down

In this solution, we're going to explicitly shut down the listener (stop accepting new connections), but will wait for clients to end *their* connections. This is a conservative approach, but it works very well in many scenarios where server shutdown is actually needed - such as tests. In a test, it's not hard to arrange for all the clients to close their connections before expecting the server to shut down.

I'll be presenting the code piece by piece, but the full runnable code sample is [available here](https://github.com/eliben/code-for-blog/blob/master/2020/tcp-server-shutdown/shutdown1/shutdown1.go). Let's start with the server type and the constructor:

```
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

Here's the `serve` method the constructor invokes:

```
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

Here's the `Stop` method that tells the server to shut down gracefully:

```
func (s *Server) Stop() {
  close(s.quit)
  s.listener.Close()
  s.wg.Wait()
}
```

It starts by closing the `s.quit` channel. Then it closes the listener. This will cause the `Accept` call in `serve` to return an error. Since `s.quit` is already closed at this point, `serve` will return.

The last line in `Stop` is waiting on `s.wg`, which is also critical. Note that `serve` notifies the wait group that it's done on return. But this is not the only goroutine we're waiting for. Each call to `handleConnection` is wrapped by a `wg` add/done pair as well. Therefore, `Stop` will block until all the handlers have returned, *and* `serve` stopped accepting new clients. This is a safe shutdown point.

For completeness, here's is `handleConnection`; the one here just reads client data and logs it, without sending anything back. Naturally, this part of the code will be different for each server:

```
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

```
s := NewServer(addr)
// do whatever here...
s.Stop()
```

Recall that `NewServer` returns a server but doesn't block. `s.Stop` does block, however. In tests, what you'd do for graceful shutdown is:

1. Make sure all clients interacting with the server have closed their connections.
2. Wait for `s.Stop` to return.

## Step 2: actively closing open client connections

In step 1, we expected all clients to close their connections before declaring the shutdown process successful. Here we'll look at a more aggressive approach, where on `Stop()` the server will actively attempt to close open client connections. I'll present a technique that's both simple and robust first, at the cost of some performance. After that, we'll discuss some alternatives.

The full code for this step is [available too](https://github.com/eliben/code-for-blog/blob/master/2020/tcp-server-shutdown/shutdown2/shutdown2.go). It's identical to step 1 except the code of `handleConection`:

```
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
        if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
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

This approach is robust, in the sense that we're (most likely) not going to close connections while the client is actively sending something. It's also simple, since it limits all the extra logic to `handleConnection`.

There's a performance cost, of course. First, a `conn.Read` call is issued once every 200 ms, which is slightly slower than a single blocking call; I'd say this is negligible, though. More seriously, every `Stop` request will be delayed by 200 ms. This is probably OK in most scenarios where we want to shut down a server, but the deadline can be tweaked to fit specific protocol needs.

An alternative to this design would be to keep track of all the open connections outside `handleConection`, and force-close them when `Stop` is called. This would likely be more efficient, at the cost of implementation complexity and some lack of robustness. Such a `Stop` could easily close connections while clients are actively sending data, resulting in client errors.

For inspiration on the right path to take, we can look at the stdlib's `http.Server.Shutdown` method, which is documented as follows:

> Shutdown gracefully shuts down the server without interrupting any active connections. Shutdown works by first closing all open listeners, then closing all idle connections, and then waiting indefinitely for connections to return to idle and then shut down

What does "idle" mean here? Roughly that the client hasn't sent any requests for some period of time. The HTTP server has advantage over a generic TCP server, because it's a higher-level protocol, so it knows the client communication pattern. In different protocols, different shutdown strategies may make sense.

A different example is a protocol where the server initiates messages, or at least some of them. For example, a given connection may be in a state where the client is waiting for the server to send some event. It's usually safe for the server to close this connection on shutdown without waiting for anything.

## Conclusion

I would summarize this post with two general guidelines:

1. Try to make shutdowns as safe as possible
2. Think of the higher-level protocol

I typically encounter the need to shut down a TCP server while writing tests. I want each test to be self-contained and clean up after itself, including all the client-server connections and listening servers. For this scenario, step 1 works very well. Once all client connections have been closed, `Server.Stop` will return without any delays.
