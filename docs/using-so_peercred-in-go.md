# Using SO_PEERCRED in Go

Wed, Sep 11, 2019

At this year's [GopherCon](https://twitter.com/GopherCon), Gabbi Fisher ([@gabbifish](https://twitter.com/gabbifish/)) of CloudFlare made a great presentation introducing her audience to the complexities of network socket options in Go (archived video of her presentation [here](https://www.youtube.com/watch?v=pGR3r0UhoS8)). In her talk, Gabbi details how to use the network socket option `SO_REUSEADDR` to allow multiple processes on the same server to listen on the same network port. Gabbi closes by mentioning the breadth of socket options that are available beyond just her example. Inspired by her talk, I've decided to write about the `SO_PEERCRED` socket option and Go.

## Sockets and Socket Options

What are sockets in the context of network programming? In essence, sockets are a special case of file descriptors used for network connections. They behave in many ways similar to general file descriptors – for instance, you can call `Read()`, `Write()` and `Close()` on them just like when operating on files. However, sockets also have operations such as `Accept()` and `Listen()` that provide the network specific behavior for a socket. For those new to [socket programming](https://en.wikipedia.org/wiki/Network_socket), I would suggest watching Gabbi's presentation as she provides a good overview of sockets with a focus of how they are used within Go.

Sockets also have numerous options that can be used when creating or interacting with them. The `SO_REUSEADDR` option is one such option. It allows more than one socket to bind to the same address, something normally that would result in an error, while providing clear rules for routing traffic to correct socket. Likewise, `SO_PEERCRED` is a socket option specific to [Unix domain sockets](https://en.wikipedia.org/wiki/Unix_domain_socket) that provides the server (ie. listening socket) with credential information (user ID, group ID and process ID) of any connected client.

While restricted to local connections (Unix domain sockets are not accessible across the network), `SO_PEERCRED` can be a useful tool for daemons to authenticate connections without requiring additional means of authentication – for example, a user specific dameon that only should be accessible by other processes launched by the same user.

The remainder of this article will show an example of `SO_PEERCRED` being used within a simple Go program.

## An Echo Server in Go

First lets begin with a basic echo server that doesn't include `SO_PEERCRED`. The `main()` function, below, creates a socket at the path `/tmp/echo.sock`, listens for connecting clients and passes every client connection to its own goroutine running the function `handleConn()`.

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

```
$ nc -U /tmp/echo.sock
this is a line
> this is a line
^C
$
```

But what if we wanted to only allow whatever user who launched the echo server to be able to interact with it? This is where we use `SO_PEERCRED`.

## Echo Server with SO_PEERCRED

In her presentation, Gabbi showed how to use a `net.ListenerConfig` struct to assign a callback that would set `SO_REUSEADDR` prior to creating a new socket with a call to `Listen()` or `ListenPacket()`. This is essential for `SO_REUSEADDR` as it is a creation-time option of sockets. However, for our case, `SO_PEERCRED` is used on already established connections and, thus, needs to be be handled differently.

To access the credential information via `SO_PEERCRED` I have created a function called `readCreds()` which I will explain in further detail below:

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

	// net.Conn is an interface. Expect only *net.UnixConn types
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
	// 'err' within the closure. 'err2' is then the error returned
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

```go
uc, ok := c.(*net.UnixConn)
```

Next we use the method `SyscallConn()` to return a `syscall.RawConn` interface containing the necessary `Control()` method:

```go
raw, err := uc.SyscallConn()
```

The `Control()` method allows one to run a callback function (with the function signature of `func(fd int)`) against the raw socket. Here we implement a [function closure](https://tour.golang.org/moretypes/25) that allows us to execute the syscall `unix.GetsockoptUcred()` while retaining the returned values in the enclosed variables, `cred` and `err`.

```go
err2 := raw.Control(func(fd uintptr) {
    cred, err = unix.GetsockoptUcred(int(fd),
        unix.SOL_SOCKET,
        unix.SO_PEERCRED)
})
```

After handling the errors (both the one returned by `Control()` and the error returned within the closure), we can return a `*unix.Ucred` struct with the following fields:

```go
type Ucred struct {
    Pid int32
    Uid uint32
    Gid uint32
}
```

From here we can make a few additions to our original `main()`. First, at startup we get the current user ID and convert to an `int` value:

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

*(The fully modified main.go file can be found [here](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/main.go) along with [handle.go](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/handle.go) and [cred.go](https://blog.jbowen.dev/2019/09/using-so_peercred-in-go/src/peercred/cred.go) that make up this example.)*

Now, once again, we can use `nc` to test our daemon:

```
$ nc -U /tmp/echo.sock
this is a line
> this is a line
^C
$
```

Works the same as before. However, if we use `sudo` to change to another user before running the command we get a different result:

```
$ sudo -u guest nc -U /tmp/echo.sock
Unauthorized access
$
```

If we look in the terminal running our go program, we can see it disallowed the connection based on the user ID:

```
$ go run *.go
2019/09/09 13:22:36 UID mismatch (1001 != 1000). Closing connection.
```

Now our echo server is secured from anyone else attempting to connect to it other than ourselves!

## Other Uses

Saavy Unix users will note that setting the read/write permissions on the socket itself would be an easier way to restrict access without having to modify the server. Indeed, I had to run `chmod` on my socket in the example above to allow the `guest` user to write to it. But Unix file permissions might not cover all cases and what about the superuser, root? root can write to any socket but with our modifications, even root is not authorized unless root started the program:

```
$ sudo nc -U /tmp/echo.sock
Unauthorized access
$
```

And in our go program's terminal:

```
2019/09/09 13:24:32 UID mismatch (0 != 1000). Closing connection.
```

Thus `SO_PEERCRED` provides us with security that even file permissions cannot. But this only scratches the surface of possibilities. We could also use `SO_PEERCRED` to do things like:

- Allow access to any user in a list of IDs (that might not be part of a unix group)
- Allow access based on PID of a program (or even the PID of a program's ancestor)
- Logging client information (user and process ID) for each connection or each command
- Allow some users to have more privilege than others (for example, we could allow one user to run an ‘update’ command when others can only run ‘view’)

All told, there are many interesting possibilities opened up with using `SO_PEERCRED`. That said there are some limitations:

- `SO_PEERCRED` is only supported on Unix domain sockets
- Unix domain sockets are only accessible locally
- `SO_PEERCRED` is only available on Unix systems – Linux, BSD and MacOS but notably not Windows

## Wrapping Up

I hope this article proves useful in further understanding socket options in Go and the `SO_PEERCRED` option in particular.

Thanks again to Gabbi for the presentation and the inspiration for this article.