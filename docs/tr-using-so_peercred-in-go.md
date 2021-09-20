# Using SO_PEERCRED in Go

# 在 Go 中使用 SO_PEERCRED

Wed, Sep 11, 2019

At this year's [GopherCon](https://twitter.com/GopherCon), Gabbi Fisher ([@gabbifish](https://twitter.com/gabbifish/)) of CloudFlare made a great presentation introducing her audience to the complexities of network socket options in Go (archived video of her presentation [here](https://www.youtube.com/watch?v=pGR3r0UhoS8)). In her talk, Gabbi details how to use the network socket option `SO_REUSEADDR` to allow multiple processes on the same server to listen on the same network port. Gabbi closes by mentioning the breadth of socket options that are available beyond just her example. Inspired by her talk, I've decided to write about the `SO_PEERCRED` socket option and Go.

在今年的 [GopherCon](https://twitter.com/GopherCon) 上，CloudFlare 的 Gabbi Fisher ([@gabbifish](https://twitter.com/gabbifish/)) 做了一个精彩的演讲，向她的观众介绍了复杂性Go 中的网络套接字选项（她的演示存档视频 [此处](https://www.youtube.com/watch?v=pGR3r0UhoS8))。在她的演讲中，Gabbi 详细介绍了如何使用网络套接字选项“SO_REUSEADDR”来允许同一服务器上的多个进程侦听同一网络端口。 Gabbi 最后提到了除她的示例之外可用的套接字选项的广度。受到她演讲的启发，我决定撰写有关“SO_PEERCRED”套接字选项和 Go 的文章。

## Sockets and Socket Options

## 套接字和套接字选项

What are sockets in the context of network programming? In essence, sockets are a special case of file descriptors used for network connections. They behave in many ways similar to general file descriptors – for instance, you can call `Read()`, `Write()` and `Close()` on them just like when operating on files. However, sockets also have operations such as `Accept()` and `Listen()` that provide the network specific behavior for a socket. For those new to [socket programming](https://en.wikipedia.org/wiki/Network_socket), I would suggest watching Gabbi's presentation as she provides a good overview of sockets with a focus of how they are used within Go.

什么是网络编程上下文中的套接字？本质上，套接字是用于网络连接的文件描述符的一种特殊情况。它们在许多方面的行为类似于一般的文件描述符——例如，你可以像对文件进行操作一样对它们调用 `Read()`、`Write()` 和 `Close()`。但是，套接字也有诸如“Accept()”和“Listen()”之类的操作，它们为套接字提供特定于网络的行为。对于 [socket 编程](https://en.wikipedia.org/wiki/Network_socket) 的新手，我建议观看 Gabbi 的演讲，因为她很好地概述了套接字，重点介绍了它们在 Go 中的使用方式。

Sockets also have numerous options that can be used when creating or interacting with them. The `SO_REUSEADDR` option is one such option. It allows more than one socket to bind to the same address, something normally that would result in an error, while providing clear rules for routing traffic to correct socket. Likewise, `SO_PEERCRED` is a socket option specific to [Unix domain sockets](https://en.wikipedia.org/wiki/Unix_domain_socket) that provides the server (ie. listening socket) with credential information (user ID, group ID and process ID) of any connected client.

套接字还有许多选项，可以在创建它们或与之交互时使用。 `SO_REUSEADDR` 选项就是这样一种选项。它允许多个套接字绑定到同一个地址，这通常会导致错误，同时为将流量路由到正确的套接字提供明确的规则。同样，`SO_PEERCRED` 是特定于 [Unix 域套接字](https://en.wikipedia.org/wiki/Unix_domain_socket) 的套接字选项，它为服务器（即侦听套接字）提供凭据信息（用户 ID、组 ID和进程 ID)的任何连接的客户端。

While restricted to local connections (Unix domain sockets are not accessible across the network), `SO_PEERCRED` can be a useful tool for daemons to authenticate connections without requiring additional means of authentication – for example, a user specific dameon that only should be accessible by other processes launched by the same user.

虽然仅限于本地连接（无法通过网络访问 Unix 域套接字），但“SO_PEERCRED”可以成为守护进程验证连接的有用工具，而无需额外的验证方式——例如，用户特定的守护进程只能通过同一用户启动的其他进程。

The remainder of this article will show an example of `SO_PEERCRED` being used within a simple Go program.

本文的其余部分将展示一个在简单 Go 程序中使用“SO_PEERCRED”的示例。

## An Echo Server in Go

## Go 中的 Echo 服务器

First lets begin with a basic echo server that doesn't include `SO_PEERCRED`. The `main()` function, below, creates a socket at the path `/tmp/echo.sock`, listens for connecting clients and passes every client connection to its own goroutine running the function `handleConn()`.

首先让我们从一个不包含“SO_PEERCRED”的基本回显服务器开始。下面的 `main()` 函数在路径 `/tmp/echo.sock` 上创建一个套接字，侦听连接客户端并将每个客户端连接传递给它自己的运行函数 `handleConn()` 的 goroutine。

```go
package main

import (
    "log"
    "net"
    "os"
)

const sockAddr = "/tmp/echo.sock"

func main() {
    // Make sure no stale sockets present
    os.Remove(sockAddr)

    // Create new Unix domain socket
    server, err := net.Listen("unix", sockAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer server.Close()

    // Loop to process client connections
    for {
        client, err := server.Accept()
        if err != nil {
            log.Printf("Accept() failed: %s", err)
            continue
        }

        go handleConn(client)
    }
}
```

The `handleConn()` function referenced above is just a simple line-based echo – each line read will immediately be sent back to the client via `Write()`.

上面引用的 `handleConn()` 函数只是一个简单的基于行的回显——每行读取将立即通过 `Write()` 发送回客户端。

```go
package main

import (
    "bufio"
    "net"
)

func handleConn(c net.Conn) {
    b := bufio.NewReader(c)
    for {
        line, err := b.ReadBytes('\n')
        if err != nil {
            break
        }
        c.Write([]byte("> "))
        c.Write(line)
    }

    c.Close()
}
```

After launching our go program in another terminal (ie. `go run *.go`), we can easily test it with the network utility `nc` (aka netcat) using the `-U` option to indicate a Unix domain socket:

在另一个终端（即`go run *.go`）中启动我们的 go 程序后，我们可以使用网络实用程序 `nc`（又名 netcat）使用 `-U` 选项指示 Unix 域套接字轻松测试它：

```
$ nc -U /tmp/echo.sock
this is a line
> this is a line
^C
$
```

But what if we wanted to only allow whatever user who launched the echo server to be able to interact with it? This is where we use `SO_PEERCRED`.

但是，如果我们只想允许启动 echo 服务器的任何用户能够与其交互呢？这是我们使用“SO_PEERCRED”的地方。

## Echo Server with SO_PEERCRED 

## 带有 SO_PEERCRED 的 Echo 服务器

In her presentation, Gabbi showed how to use a `net.ListenerConfig` struct to assign a callback that would set `SO_REUSEADDR` prior to creating a new socket with a call to `Listen()` or `ListenPacket()`. This is essential for `SO_REUSEADDR` as it is a creation-time option of sockets. However, for our case, `SO_PEERCRED` is used on already established connections and, thus, needs to be be handled differently.

在她的演讲中，Gabbi 展示了如何使用 `net.ListenerConfig` 结构来分配一个回调，该回调将在创建一个调用 `Listen()` 或 `ListenPacket()` 的新套接字之前设置 `SO_REUSEADDR`。这对于 SO_REUSEADDR 是必不可少的，因为它是套接字的创建时选项。但是，对于我们的情况，`SO_PEERCRED` 用于已经建立的连接，因此需要以不同的方式处理。

To access the credential information via `SO_PEERCRED` I have created a function called `readCreds()` which I will explain in further detail below:

为了通过“SO_PEERCRED”访问凭证信息，我创建了一个名为“readCreds()”的函数，我将在下面进一步详细解释：

```go
// +build linux

package main

import (
    "fmt"
    "net"

    "golang.org/x/sys/unix"
)

func readCreds(c net.Conn) (*unix.Ucred, error) {

    var cred *unix.Ucred

    // net.Conn is an interface.Expect only *net.UnixConn types
    uc, ok := c.(*net.UnixConn)
    if !ok {
        return nil, fmt.Errorf("unexpected socket type")
    }

    // Fetches raw network connection from UnixConn
    raw, err := uc.SyscallConn()
    if err != nil {
        return nil, fmt.Errorf("error opening raw connection: %s", err)
    }

    // The raw.Control() callback does not return an error directly.
    // In order to capture errors, we wrap already defined variable
    // 'err' within the closure.'err2' is then the error returned
    // by Control() itself.
    err2 := raw.Control(func(fd uintptr) {
        cred, err = unix.GetsockoptUcred(int(fd),
            unix.SOL_SOCKET,
            unix.SO_PEERCRED)
    })

    if err != nil {
        return nil, fmt.Errorf("GetsockoptUcred() error: %s", err)
    }

    if err2 != nil {
        return nil, fmt.Errorf("Control() error: %s", err2)
    }

    return cred, nil
}
```

Before we can use `SO_PEERCRED` we need to get access to the `Control()` function of the raw socket. The `Accept()` call in our `main()` function returns a `net.Conn` interface which is passed to our function. However, but we need to access the real type, `*net.UnixConn`, directly to get to the raw socket so we use a standard Go [type assertion](https://tour.golang.org/methods/15) to get the underlying concrete type as variable `uc`.

在我们可以使用 SO_PEERCRED 之前，我们需要访问原始套接字的 Control() 函数。 `main()` 函数中的 `Accept()` 调用返回一个传递给我们的函数的 `net.Conn` 接口。但是，我们需要直接访问真实类型 `*net.UnixConn` 才能访问原始套接字，因此我们使用标准 Go [类型断言](https://tour.golang.org/methods/15)将底层具体类型作为变量`uc`。

```go
uc, ok := c.(*net.UnixConn)
```

Next we use the method `SyscallConn()` to return a `syscall.RawConn` interface containing the necessary `Control()` method:

接下来我们使用方法 SyscallConn() 返回一个包含必要的 Control() 方法的 syscall.RawConn 接口：

```go
raw, err := uc.SyscallConn()
```

The `Control()` method allows one to run a callback function (with the function signature of `func(fd int)`) against the raw socket. Here we implement a [function closure](https://tour.golang.org/moretypes/25) that allows us to execute the syscall `unix.GetsockoptUcred()` while retaining the returned values in the enclosed variables, `cred` and `err`.

`Control()` 方法允许对原始套接字运行回调函数（具有 `func(fd int)` 的函数签名）。这里我们实现了一个[函数闭包](https://tour.golang.org/moretypes/25)，它允许我们执行系统调用`unix.GetsockoptUcred()`，同时保留封闭变量`cred`中的返回值和`错误`。

```go
err2 := raw.Control(func(fd uintptr) {
    cred, err = unix.GetsockoptUcred(int(fd),
        unix.SOL_SOCKET,
        unix.SO_PEERCRED)
})
```

After handling the errors (both the one returned by `Control()` and the error returned within the closure), we can return a `*unix.Ucred` struct with the following fields:

处理完错误（Control() 返回的错误和闭包中返回的错误）后，我们可以返回一个包含以下字段的 *unix.Ucred 结构：

```go
type Ucred struct {
    Pid int32
    Uid uint32
    Gid uint32
}
```

From here we can make a few additions to our original `main()`. First, at startup we get the current user ID and convert to an `int` value:

从这里我们可以对我们原来的 `main()` 做一些添加。首先，在启动时，我们获取当前用户 ID 并转换为一个 `int` 值：

```go
uidStr, err := user.Current()
if err != nil {
    log.Fatal(err)
}

uid, err := strconv.Atoi(uidStr.Uid)
if err != nil {
    log.Fatal(err)
}
```

Next we add a call to `readCreds()` after our `Accept()` call to check the UID of the connecting client to that of the running daemon:

接下来，我们在 `Accept()` 调用之后添加对 `readCreds()` 的调用，以检查连接客户端的 UID 与正在运行的守护进程的 UID：

```go
creds, err := readCreds(client)
if err != nil {
    log.Printf("Error reading credentials: %s", err)
    continue
}

if creds.Uid != uint32(uid) {
    log.Printf("UID mismatch (%d != %d). Closing connection.\n", creds.Uid, uid)
    client.Write([]byte("Unauthorized access\n"))
    client.Close()
    continue
}
```

*(The fully modified main.go file can be found [here](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/main.go) along with [handle .go](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/handle.go) and [cred.go](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/cred.go) that make up this example.)*

*（完全修改的 main.go 文件可以在 [here](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/main.go) 和 [handle .go](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/handle.go) 和 [cred.go](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/cred.go) 构成此示例。)*

Now, once again, we can use `nc` to test our daemon:

现在，我们可以再次使用 `nc` 来测试我们的守护进程：

```
$ nc -U /tmp/echo.sock
this is a line
> this is a line
^C
$
```

Works the same as before. However, if we use `sudo` to change to another user before running the command we get a different result:

和以前一样工作。但是，如果我们在运行命令之前使用 `sudo` 更改为另一个用户，我们会得到不同的结果：

```
$ sudo -u guest nc -U /tmp/echo.sock
Unauthorized access
$
```

If we look in the terminal running our go program, we can see it disallowed the connection based on the user ID:

如果我们查看运行 go 程序的终端，我们可以看到它根据用户 ID 禁止连接：

```
$ go run *.go
2019/09/09 13:22:36 UID mismatch (1001 != 1000).Closing connection.
```

Now our echo server is secured from anyone else attempting to connect to it other than ourselves!

现在，我们的回显服务器受到了除我们之外的任何其他人尝试连接到它的保护！

## Other Uses

## 其他用途

Saavy Unix users will note that setting the read/write permissions on the socket itself would be an easier way to restrict access without having to modify the server. Indeed, I had to run `chmod` on my socket in the example above to allow the `guest` user to write to it. But Unix file permissions might not cover all cases and what about the superuser, root? root can write to any socket but with our modifications, even root is not authorized unless root started the program:

Saavy Unix 用户会注意到，在套接字本身上设置读/写权限将是一种更简单的方法来限制访问，而无需修改服务器。实际上，在上面的示例中，我必须在我的套接字上运行 `chmod` 以允许 `guest` 用户对其进行写入。但是 Unix 文件权限可能无法涵盖所有情况，那么超级用户 root 呢？ root 可以写入任何套接字，但经过我们的修改，除非 root 启动程序，否则即使 root 也未被授权：

```
$ sudo nc -U /tmp/echo.sock
Unauthorized access
$
```

And in our go program's terminal:

在我们的 go 程序的终端中：

```
2019/09/09 13:24:32 UID mismatch (0 != 1000).Closing connection.
```

Thus `SO_PEERCRED` provides us with security that even file permissions cannot. But this only scratches the surface of possibilities. We could also use `SO_PEERCRED` to do things like:

因此，`SO_PEERCRED` 为我们提供了甚至文件权限都无法提供的安全性。但这只是触及了可能性的表面。我们还可以使用 `SO_PEERCRED` 来做这样的事情：

- Allow access to any user in a list of IDs (that might not be part of a unix group)
- Allow access based on PID of a program (or even the PID of a program's ancestor)
- Logging client information (user and process ID) for each connection or each command
- Allow some users to have more privilege than others (for example, we could allow one user to run an ‘update’ command when others can only run ‘view’)

- 允许访问 ID 列表中的任何用户（可能不是 unix 组的一部分）
- 允许基于程序的 PID（甚至是程序祖先的 PID）的访问
- 记录每个连接或每个命令的客户端信息（用户和进程 ID）
- 允许某些用户拥有比其他用户更多的权限（例如，当其他用户只能运行“查看”时，我们可以允许一个用户运行“更新”命令）

All told, there are many interesting possibilities opened up with using `SO_PEERCRED`. That said there are some limitations:

总而言之，使用“SO_PEERCRED”开辟了许多有趣的可能性。也就是说有一些限制：

- `SO_PEERCRED` is only supported on Unix domain sockets
- Unix domain sockets are only accessible locally
- `SO_PEERCRED` is only available on Unix systems – Linux, BSD and MacOS but notably not Windows

- `SO_PEERCRED` 仅在 Unix 域套接字上受支持
- Unix 域套接字只能在本地访问
- `SO_PEERCRED` 仅适用于 Unix 系统——Linux、BSD 和 MacOS，但尤其不适用于 Windows

## Wrapping Up

##  总结

I hope this article proves useful in further understanding socket options in Go and the `SO_PEERCRED` option in particular.

我希望本文对进一步理解 Go 中的套接字选项特别是“SO_PEERCRED”选项有用。

Thanks again to Gabbi for the presentation and the inspiration for this article. 

再次感谢 Gabbi 的介绍和本文的灵感。

