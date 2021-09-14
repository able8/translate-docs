# How to Set Go net/http Socket Options - setsockopt() example

August 19, 2021

[Go,](http://iximiuz.com/en/categories/?category=Go) [Programming](http://iximiuz.com/en/categories/?category=Programming)

**TL;DR** Don't care about the story? Jump straight to the working code:

- [Example - How to set net/http server socket options](http://iximiuz.com#net-http-set-server-socket-options)
- [Example - How to set net/http client socket options](http://iximiuz.com#net-http-set-client-socket-options)

Go standard library makes it super easy to start an HTTP server:

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

```go
package main

import "net/http"

func main() {
    resp, err := http.Get("http://example.com/")
    body, err := io.ReadAll(resp.Body)
}

```

In just ~10 lines of code, I can get a server up and running or fetch a real web page! In contrast, [creating a basic HTTP server in C would take hundreds of lines](https://gist.github.com/laobubu/d6d0e9beb934b60b2e552c2d03e1409e), and anything beyond basics would require third-party libraries.

The Go snippets from above are so short because they rely on powerful high-level abstractions of the [`net`](https://pkg.go.dev/net) and [`net/http`](https://pkg.go.dev/net/http) packages. Go pragmatically chooses to optimize for frequently used scenarios, and its standard library hides many internal socket details behind these abstractions, making lots of default choices on the way. And that's very handy, but...

**What if I need to fine-tune `net/http` sockets before initiating the communication?** For instance, how can I set some socket options like `SO_REUSEPORT` or `TCP_QUICKACK`?

When it comes to near-systems-programming questions, I refer to the corresponding example in plain C first. Simply because that's the closest I can get to the actual system API layer. Usually, it'd give me a hint of what to look for in my primary language. And this also helps to straighten my understanding of fundamental concepts.

This time, unlike with the creation of an HTTP server, C code to [set options on sockets](https://linux.die.net/man/2/setsockopt) looks concise:

```c
int sfd = socket(domain, socktype, 0);

int optval = 1;
setsockopt(sfd, SOL_SOCKET, SO_REUSEPORT, &optval, sizeof(optval));

bind(sfd, (struct sockaddr *) &addr, addrlen);

```

Basically, a single line with the `setsockopt()` call that accepts the socket's file descriptor and the option to be changed.

Getting back to Go, there is a corresponding [family of `setsockopt()` wrappers](https://pkg.go.dev/syscall#SetsockoptInt) in the `syscall` package. Very handy again, but... **how to obtain the raw file descriptor behind a `net/http` server or client?**

## How to access underlying net/http sockets

Neither `http.Handle() + http.ListenAndServe()` nor `http.GET()` usage examples give me a hint on how to get access to the underlying sockets. Luckily, the Go standard library has pretty readable code. Oftentimes, inspecting the internals of a module could be even faster than scrolling through a massive documentation page.

So, let's take a look at [`http.ListenAndServe()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L3182-L3185):

```go
// src/net/http/server.go

func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}

```

Lovely, just two lines of code! So, there is a server object under the hood. Well, it makes sense. Something needs to keep the state. Ok, what's hidden behind [`server.ListenAndServe()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L2918-L2931)?

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

Looking around for another `srv.Serve()` example brought me to a public [`http.Serve()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/http/server.go#L2503-L2506) function. It can be used with a custom listener:

```go
// src/net/http/server.go

func Serve(l net.Listener, handler Handler) error {
    srv := &Server{Handler: handler}
    return srv.Serve(l)
}

```

Having access to the listener object could allow me to get to the socket file descriptor. `net.Listener` is a supertype, but it can be downcasted to [`net.TCPListener`](https://pkg.go.dev/net#TCPListener), and the later one has a handy [File()](https://pkg.go.dev/net#TCPListener.File) method returning an `os.File` wrapper around the much-needed file descriptor. Should be safe enough for both HTTP/1.1 and HTTP/2.

**_Or so thought I..._**

The new version of the HTTP server became slightly more verbose but still not too many lines:

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

    if err := http.Serve(ln, nil); err != nil {
        panic(err)
    }
}

```

**And to my utter surprise, setting socket options on the file descriptor obtained like that didn't work.**

For some reason, after discovering the `net.Listener()` \+ `http.Serve()` trick, I stopped paying attention to the standard library code. And it cost me a few hours of debugging.

In actuality, `net.TCPListener.File()` returns a **copy** of a listening file descriptor ( [1](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/tcpsock.go#L305-L314), [2](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/tcpsock_posix.go#L160-L166)). So, it might be useful in some cases but wasn't really helpful in mine.

But a much better clue was actually hidden in the C code from the beginning of this article! In that tiny C snippet, setting the option was happening **_before the `bind()` call_**. Since `net.Listen()` call returns an already bound socket, I needed to look how to hook into on the listener creation phase.

I applied my traditional technique to [`net.Listen()`](https://github.com/golang/go/blob/3bdc1799d6ce441d7a972faf1452e34b6dce0826/src/net/dial.go#L710-L713) and finally discovered this:

```go
// src/net/dial.go

func Listen(network, address string) (Listener, error) {
    var lc ListenConfig
    return lc.Listen(context.Background(), network, address)
}

```

So, `Listen()` is actually a method of some obscure [`ListenConfig` type](https://pkg.go.dev/net#ListenConfig). And `ListenConfig` has a field called `Control` holding a function with the following signature and comment:

```go
// If Control is not nil, it is called after creating the network
// connection but before binding it to the operating system.
// ...
Control func(network, address string, c syscall.RawConn) error

```

As it usually happens, the answer was hidden in plain sight 🙈 Below is the boring part, the two working snippets - one for setting the server-side socket options, and another one - for the client-side socket. The latter one uses a slightly different data structure, but the setting happens via a similar `Control` function.

## How to set net/http server socket options

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
            }); err != nil {
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

    if err := http.Serve(ln, nil); err != nil {
        panic(err)
    }
}

```

## How to set net/http client socket options

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
            }); err != nil {
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

Documentation is often boring while reading code is always fun! But don't forget to double-check your findings - by referring to the docs and, of course, by running code before putting it in production 😉

[golang,](javascript: void 0) [socket,](javascript: void 0) [http,](javascript: void 0) [tcp,](javascript: void 0) [linux](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

