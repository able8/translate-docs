# How to Set Go net/http Socket Options - setsockopt() example

# 如何设置 Go net/http 套接字选项 - setsockopt() 示例

August 19, 2021

2021 年 8 月 19 日

[Go,](http://iximiuz.com/en/categories/?category=Go)[Programming](http://iximiuz.com/en/categories/?category=Programming)

[Go,](http://iximiuz.com/en/categories/?category=Go)[编程](http://iximiuz.com/en/categories/?category=Programming)

**TL;DR** Don't care about the story? Jump straight to the working code:

**TL;DR** 不关心故事？直接跳转到工作代码：

- [Example - How to set net/http server socket options](http://iximiuz.com#net-http-set-server-socket-options)
- [Example - How to set net/http client socket options](http://iximiuz.com#net-http-set-client-socket-options)

- [示例 - 如何设置 net/http 服务器套接字选项](http://iximiuz.com#net-http-set-server-socket-options)
- [示例 - 如何设置 net/http 客户端套接字选项](http://iximiuz.com#net-http-set-client-socket-options)

Go standard library makes it super easy to start an HTTP server:

Go 标准库使启动 HTTP 服务器变得非常容易：

```go
package main

import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello there!\n"))
    })

    http.ListenAndServe(":8080", nil)
}

```

...or send an HTTP request:

...或发送 HTTP 请求：

```go
package main

import "net/http"

func main() {
    resp, err := http.Get("http://example.com/")
    body, err := io.ReadAll(resp.Body)
}

```

In just ~10 lines of code, I can get a server up and running or fetch a real web page! In contrast, [creating a basic HTTP server in C would take hundreds of lines](https://gist.github.com/laobubu/d6d0e9beb934b60b2e552c2d03e1409e), and anything beyond basics would require third-party libraries.

只需大约 10 行代码，我就可以启动并运行服务器或获取真实的网页！相比之下，[用 C 语言创建一个基本的 HTTP 服务器需要数百行代码](https://gist.github.com/laobubu/d6d0e9beb934b60b2e552c2d03e1409e)，任何超出基础的东西都需要第三方库。

The Go snippets from above are so short because they rely on powerful high-level abstractions of the [`net`](https://pkg.go.dev/net) and [`net/http`](https://pkg.go.dev/net/http) packages. Go pragmatically chooses to optimize for frequently used scenarios, and its standard library hides many internal socket details behind these abstractions, making lots of default choices on the way. And that's very handy, but...

上面的 Go 代码片段很短，因为它们依赖于 [`net`](https://pkg.go.dev/net) 和 [`net/http`](https://pkg.go.dev/net/http) 包。 Go 务实地选择针对频繁使用的场景进行优化，其标准库在这些抽象背后隐藏了许多内部套接字细节，在此过程中做出了许多默认选择。这非常方便，但是...

**What if I need to fine-tune `net/http` sockets before initiating the communication?** For instance, how can I set some socket options like `SO_REUSEPORT` or `TCP_QUICKACK`?

**如果我需要在启动通信之前微调 `net/http` 套接字怎么办？** 例如，我如何设置一些像 `SO_REUSEPORT` 或 `TCP_QUICKACK` 这样的套接字选项？

When it comes to near-systems-programming questions, I refer to the corresponding example in plain C first. Simply because that's the closest I can get to the actual system API layer. Usually, it'd give me a hint of what to look for in my primary language. And this also helps to straighten my understanding of fundamental concepts.

当谈到近系统编程问题时，我首先参考普通 C 中的相应示例。仅仅因为这是我能得到的最接近实际系统 API 层的方法。通常，它会给我提示在我的主要语言中寻找什么。这也有助于理顺我对基本概念的理解。

This time, unlike with the creation of an HTTP server, C code to [set options on sockets](https://linux.die.net/man/2/setsockopt) looks concise:

这一次，与创建 HTTP 服务器不同，[在套接字上设置选项](https://linux.die.net/man/2/setsockopt) 的 C 代码看起来很简洁：

```c
int sfd = socket(domain, socktype, 0);

int optval = 1;
setsockopt(sfd, SOL_SOCKET, SO_REUSEPORT, &optval, sizeof(optval));

bind(sfd, (struct sockaddr *) &addr, addrlen);

```

Basically, a single line with the `setsockopt()` call that accepts the socket's file descriptor and the option to be changed.

基本上，一行带有`setsockopt()` 调用，它接受套接字的文件描述符和要更改的选项。

Getting back to Go, there is a corresponding [family of `setsockopt()` wrappers](https://pkg.go.dev/syscall#SetsockoptInt) in the `syscall` package. Very handy again, but... **how to obtain the raw file descriptor behind a `net/http` server or client?**

回到 Go，在 `syscall` 包中有一个相应的 [`setsockopt()` 包装器系列](https://pkg.go.dev/syscall#SetsockoptInt)。再次非常方便，但是... **如何获取`net/http`服务器或客户端背后的原始文件描述符？**

## How to access underlying net/http sockets

## 如何访问底层的 net/http 套接字

Neither `http.Handle() + http.ListenAndServe()` nor `http.GET()` usage examples give me a hint on how to get access to the underlying sockets. Luckily, the Go standard library has pretty readable code. Oftentimes, inspecting the internals of a module could be even faster than scrolling through a massive documentation page.

`http.Handle() + http.ListenAndServe()` 和 `http.GET()` 使用示例都没有给我一个关于如何访问底层套接字的提示。幸运的是，Go 标准库的代码可读性很强。通常，检查模块的内部结构甚至比滚动浏览大量文档页面还要快。

So, let's take a look at [`http.ListenAndServe()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L3182-L3185):

那么，我们来看看[`http.ListenAndServe()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L3182-L3185)：

```go
// src/net/http/server.go

func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}

```

Lovely, just two lines of code! So, there is a server object under the hood. Well, it makes sense. Something needs to keep the state. Ok, what's hidden behind [`server.ListenAndServe()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L2918-L2931)?

可爱，只有两行代码！因此，引擎盖下有一个服务器对象。嗯，这是有道理的。有些东西需要保持状态。好的，[`server.ListenAndServe()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L2918-L2931)背后隐藏了什么？

```go
// src/net/http/server.go

func (srv *Server) ListenAndServe() error {
    // ...

    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(ln)
}

```

Nice! There is a [net.Listener](https://pkg.go.dev/net#Listener) instance! Apparently, that's a listener object that holds the _listening socket_. 

好的！有一个 [net.Listener](https://pkg.go.dev/net#Listener) 实例！显然，这是一个包含 _listening socket_ 的侦听器对象。

Looking around for another `srv.Serve()` example brought me to a public [`http.Serve()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L2503-L2506) function. It can be used with a custom listener:

寻找另一个 `srv.Serve()` 示例将我带到了一个公共 [`http.Serve()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L2503-L2506) 函数。它可以与自定义侦听器一起使用：

```go
// src/net/http/server.go

func Serve(l net.Listener, handler Handler) error {
    srv := &Server{Handler: handler}
    return srv.Serve(l)
}

```

Having access to the listener object could allow me to get to the socket file descriptor. `net.Listener` is a supertype, but it can be downcasted to [`net.TCPListener`](https://pkg.go.dev/net#TCPListener), and the later one has a handy [File()] (https://pkg.go.dev/net#TCPListener.File) method returning an `os.File` wrapper around the much-needed file descriptor. Should be safe enough for both HTTP/1.1 and HTTP/2.

访问侦听器对象可以让我访问套接字文件描述符。 `net.Listener` 是一个超类型，但是可以向下转换为 [`net.TCPListener`](https://pkg.go.dev/net#TCPListener)，后面的有个方便的[File()] (https://pkg.go.dev/net#TCPListener.File) 方法返回一个围绕急需的文件描述符的 `os.File` 包装器。对于 HTTP/1.1 和 HTTP/2 来说应该足够安全。

**_Or so thought I..._**

**_或者我以为我..._**

The new version of the HTTP server became slightly more verbose but still not too many lines:

新版本的 HTTP 服务器变得更加冗长，但仍然没有太多行：

```go
package main

import (
    "fmt"
    "net"
    "net/http"
)

func main() {
    ln, err := net.Listen("tcp", "127.0.0.1:8080")
    if err != nil {
        panic(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, _req *http.Request) {
        w.Write([]byte("Hello, world!\n"))
    })

    file, err := ln.(*net.TCPListener).File()
    if err != nil {
        panic(err)
    }
    fmt.Println("Socket file descriptor:", file.Fd())
    // syscall.SetsockoptInt(file.Fd(), ...)

    if err := http.Serve(ln, nil);err != nil {
        panic(err)
    }
}

```

**And to my utter surprise, setting socket options on the file descriptor obtained like that didn't work.**

**令我大吃一惊的是，在这样获得的文件描述符上设置套接字选项不起作用。**

For some reason, after discovering the `net.Listener()` \+ `http.Serve()` trick, I stopped paying attention to the standard library code. And it cost me a few hours of debugging.

出于某种原因，在发现了 `net.Listener()` \+ `http.Serve()` 技巧后，我不再关注标准库代码。它花费了我几个小时的调试时间。

In actuality, `net.TCPListener.File()` returns a **copy** of a listening file descriptor ( [1](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/tcpsock.go#L305-L314), [2](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/tcpsock_posix.go#L160-L166)). So, it might be useful in some cases but wasn't really helpful in mine.

实际上，`net.TCPListener.File()` 返回侦听文件描述符 ( [1](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/tcpsock.go#L305-L314), [2](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/tcpsock_posix.go#L160-L166))。所以，它在某些情况下可能很有用，但对我来说并没有真正的帮助。

But a much better clue was actually hidden in the C code from the beginning of this article! In that tiny C snippet, setting the option was happening **_before the `bind()` call_**. Since `net.Listen()` call returns an already bound socket, I needed to look how to hook into on the listener creation phase.

但在本文开头的 C 代码中实际上隐藏了一个更好的线索！在那个小小的 C 代码片段中，设置选项发生在 **_before the `bind()` call_**。由于`net.Listen()` 调用返回一个已经绑定的套接字，我需要研究如何在侦听器创建阶段进行挂钩。

I applied my traditional technique to [`net.Listen()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/dial.go#L710-L713) and finally discovered this:

我将我的传统技术应用于 [`net.Listen()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/dial.go#L710-L713) 并最终发现了这一点：

```go
// src/net/dial.go

func Listen(network, address string) (Listener, error) {
    var lc ListenConfig
    return lc.Listen(context.Background(), network, address)
}

```

So, `Listen()` is actually a method of some obscure [`ListenConfig` type](https://pkg.go.dev/net#ListenConfig). And `ListenConfig` has a field called `Control` holding a function with the following signature and comment:

所以，`Listen()` 实际上是一种晦涩的 [`ListenConfig` 类型](https://pkg.go.dev/net#ListenConfig) 的方法。而`ListenConfig` 有一个名为`Control` 的字段，其中包含一个具有以下签名和注释的函数：

```go
// If Control is not nil, it is called after creating the network
// connection but before binding it to the operating system.
// ...
Control func(network, address string, c syscall.RawConn) error

```

As it usually happens, the answer was hidden in plain sight 🙈 Below is the boring part, the two working snippets - one for setting the server-side socket options, and another one - for the client-side socket. The latter one uses a slightly different data structure, but the setting happens via a similar `Control` function.

正如通常发生的那样，答案隐藏在显而易见的地方🙈 下面是无聊的部分，两个工作片段 - 一个用于设置服务器端套接字选项，另一个用于客户端套接字。后一个使用稍微不同的数据结构，但设置是通过类似的“控制”功能进行的。

## How to set net/http server socket options

## 如何设置 net/http 服务器套接字选项

```go
package main

import (
    "context"
    "net"
    "net/http"
    "syscall"

    "golang.org/x/sys/unix"
)

func main() {
    lc := net.ListenConfig{
        Control: func(network, address string, conn syscall.RawConn) error {
            var operr error
            if err := conn.Control(func(fd uintptr) {
                operr = syscall.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
            });err != nil {
                return err
            }
            return operr
        },
    }

    ln, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:8080")
    if err != nil {
        panic(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, _req *http.Request) {
        w.Write([]byte("Hello, world!\n"))
    })

    if err := http.Serve(ln, nil);err != nil {
        panic(err)
    }
}

```

## How to set net/http client socket options

## 如何设置 net/http 客户端套接字选项

```go
import (
    "fmt"
    "io"
    "net"
    "net/http"
    "syscall"

    "golang.org/x/sys/unix"
)

func main() {
    dialer := &net.Dialer{
        Control: func(network, address string, conn syscall.RawConn) error {
            var operr error
            if err := conn.Control(func(fd uintptr) {
                operr = syscall.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.TCP_QUICKACK, 1)
            });err != nil {
                return err
            }
            return operr
        },
    }

    client := &http.Client{
        Transport: &http.Transport{
            DialContext: dialer.DialContext,
        },
    }

    resp, err := client.Get("http://example.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body))
}

```

## Instead of conclusion

## 而不是结论

Documentation is often boring while reading code is always fun! But don't forget to double-check your findings - by referring to the docs and, of course, by running code before putting it in production 😉

文档通常很无聊，而阅读代码总是很有趣！但是不要忘记仔细检查您的发现 - 通过参考文档，当然，在将其投入生产之前运行代码😉

[golang,](javascript: void 0) [socket,](javascript: void 0) [http,](javascript: void 0) [tcp,](javascript: void 0) [linux](javascript: void 0)

[golang,](javascript: void 0) [socket,](javascript: void 0) [http,](javascript: void 0) [tcp,](javascript: void 0) [linux](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

