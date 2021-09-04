# Writing TCP scanner in Go

# 在 Go 中编写 TCP 扫描器

Go is perfect for network applications. Its awesome standard library helps a lot in writing such software. In this article, we’ll write a simple TCP scanner in Go. The whole programm will take less than 50 lines of code. Before we’ll go to practice - a little theory.

Go 非常适合网络应用程序。其出色的标准库对编写此类软件有很大帮助。在本文中，我们将用 Go 编写一个简单的 TCP 扫描器。整个程序将花费不到 50 行代码。在我们开始实践之前 - 一个小理论。

Of course, the TCP is [more complicated than I describe](http://www.medianet.kent.edu/techreports/TR2005-07-22-tcp-EFSM.pdf) but we need just basics. The TCP handshake is three-way. Firstly, the client sends the `syn` package which signals the beginning of a communication. If the client gets a timeout here it may mean that the port is behind a firewall.

当然，TCP [比我描述的更复杂](http://www.medianet.kent.edu/techreports/TR2005-07-22-tcp-EFSM.pdf) 但我们只需要基础知识。 TCP 握手是三向的。首先，客户端发送`syn` 包，表示通信开始。如果客户端在此处超时，则可能意味着该端口位于防火墙后面。

![syn package](http://developer20.com/images/diagram-01-sync.png)

Secondly, the server answers with `syn-ack` when the port is opened, otherwise it responses with `rst` package. In the end, the client has to send another packet called ack. From this point, the connection is established.

其次，服务器在端口打开时以`syn-ack`响应，否则以`rst`包响应。最后，客户端必须发送另一个称为 ack 的数据包。至此，连接建立。

![syn package](http://developer20.com/images/diagram-02-sync-ack.png)![synpackage](http://developer20.com/images/diagram-03-ack.png)

The first step in writing the TCP scanner is to test a single port. We’ll use the `net.Dial` function which accepts two parameters: the protocol and the address to test (with the port number).

编写 TCP 扫描程序的第一步是测试单个端口。我们将使用`net.Dial` 函数，它接受两个参数：协议和要测试的地址（带有端口号）。

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "google.com:80")
	if err == nil {
		fmt.Println("Connection successful")
	} else {
		fmt.Println(err)
	}
}
```


To not test every port one by one, we’ll add a simple loop that will simplify the whole process. Notice the Sprintf function which concretes the host and the port.

为了不一一测试每个端口，我们将添加一个简单的循环来简化整个过程。注意将主机和端口具体化的 Sprintf 函数。

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	for port := 80; port < 100; port++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("google.com:%d", port))
		if err == nil {
			conn.Close()
			fmt.Println("Connection successful")
		} else {
			fmt.Println(err)
		}
	}
}
```




The solution has one huge issue - it’s extremely slow. We can do two things to make things faster: run those checks concurrently and add a timeout to every connection.

该解决方案有一个大问题——它非常慢。我们可以做两件事来加快速度：同时运行这些检查并为每个连接添加超时。

Let’s focus on making in concurrent. The first step is to extract the scanning to a separate function. This step will make our code more clear.

让我们专注于并发。第一步是将扫描提取到一个单独的函数中。这一步将使我们的代码更加清晰。

```go
func isOpen(host string, port int) bool {
  conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
  if err == nil {
     _ = conn.Close()
     return true
  }

  return false
}
```


The only new thing is the `WaitGroup`. You can read about it in more detail [here](https://gobyexample.com/waitgroups) or, if you want (let me know in the comments below) I can write an article about async programming in Go. But, back to the topic… In the main function, we span our goroutines and wait for the execution to finish.

唯一的新事物是“WaitGroup”。您可以 [此处](https://gobyexample.com/waitgroups)更详细地了解它，或者，如果您愿意（在下面的评论中告诉我)，我可以写一篇关于 Go 异步编程的文章。但是，回到主题……在 main 函数中，我们跨越我们的 goroutines 并等待执行完成。

```go
func main() {
  ports := []int{}

  wg := &sync.WaitGroup{}
  for port := 1; port < 100; port++ {
     wg.Add(1)
     go func() {
        opened := isOpen("google.com", port)
        if opened {
           ports = append(ports, port)
        }
        wg.Done()
     }()
  }

  wg.Wait()
  fmt.Printf("opened ports: %v\n", ports)
}
```


Our code is faster but because of timeouts, we’re waiting a very long time to receive the error. We can assume that if we don’t get any response from the server for 200 ms we don’t want to wait longer.

我们的代码速度更快，但由于超时，我们等待很长时间才能收到错误消息。我们可以假设，如果我们在 200 毫秒内没有收到来自服务器的任何响应，我们不想等待更长的时间。

```go
func isOpen(host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
}

func main() {
	ports := []int{}

	wg := &sync.WaitGroup{}
	timeout := time.Millisecond * 200
	for port := 1; port < 100; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen("google.com", p, timeout)
			if opened {
				ports = append(ports, p)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
}
```


At this point, we have a working simple port scanner. Unfortunately, it’s not very handy because to change the domain or port ranges we have to edit the code and recompile. Go has an awesome package called `flag`.

在这一点上，我们有一个简单的端口扫描器。不幸的是，这不是很方便，因为要更改域或端口范围，我们必须编辑代码并重新编译。 Go 有一个很棒的包叫做 `flag`。

The `flag` package helps in writing command-line applications. You can read more about it in [Go by Example](https://gobyexample.com/command-line-flags). What we want is configuring every magic string or number. We add parameters for the hostname, port range we want to test and the timeout on the connection.

`flag` 包有助于编写命令行应用程序。您可以在 [Go by Example](https://gobyexample.com/command-line-flags) 中阅读更多相关信息。我们想要的是配置每个魔术字符串或数字。我们为主机名、要测试的端口范围和连接超时添加参数。

```go
func main() {
	hostname := flag.String("hostname", "", "hostname to test")
	startPort := flag.Int("start-port", 80, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 100, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond * 200, "timeout")
	flag.Parse()

	ports := []int{}

	wg := &sync.WaitGroup{}
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				ports = append(ports, p)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
}
```


If we want to show the usage, we have to put the -h parameter which will show us the usage. Simple and clear. The whole project took less than 50 lines of code. We used concurrency, the flag, and net packages. 

如果我们想显示用法，我们必须放 -h 参数，它会显示我们的用法。简单明了。整个项目用了不到 50 行代码。我们使用了并发、标志和网络包。

There’s one more thing. Our program has race condition. In only a few opened ports and so slow scanning it’s not visible but there’s the issue. To fix that we’ll add [a mutex](https://gobyexample.com/mutexes).

还有一件事。我们的程序有竞争条件。仅在几个打开的端口中，扫描速度很慢，但不可见，但这就是问题所在。为了解决这个问题，我们将添加 [a mutex](https://gobyexample.com/mutexes)。

```go
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
```


If you like this kind of posts or have a question, let me know in the comments section below. The whole source code is available [on GitHub](https://github.com/bkielbasa/port-scanner).

如果您喜欢此类帖子或有疑问，请在下面的评论部分告诉我。整个源代码可在 [GitHub](https://github.com/bkielbasa/port-scanner) 上找到。

[![Buy me a coffee](https://cdn.buymeacoffee.com/buttons/bmc-new-btn-logo.svg)Buy me a coffee](https://www.buymeacoffee.com/bklimczak)

Tags:
[#golang](https://developer20.com/tags/golang/)[#tcp](https://developer20.com/tags/tcp/) [#scanner](https://developer20.com/tags/scanner/) [#network](https://developer20.com/tags/network/)[#concurrency](https://developer20.com/tags/concurrency/)

标签：
[#golang](https://developer20.com/tags/golang/)[#tcp](https://developer20.com/tags/tcp/) [#scanner](https://developer20.com/tags/scanner/) [#network](https://developer20.com/tags/network/)[#concurrency](https://developer20.com/tags/concurrency/)

### See Also

###  也可以看看

- [How to send multiple variables via channel in golang?](http://developer20.com/how-to-send-multiple-variables-via-channel-in-golang/)
- [Golang Tips & Tricks #7 - private repository and proxy](http://developer20.com/golang-tips-and-trics-vii/)
- [How I organize packages in Go](http://developer20.com/how-i-organize-packages-in-go/)
- [Golang Tips & Tricks #6 - the \_test package](http://developer20.com/golang-tips-and-trics-vi/)
- [Golang Tips & Tricks #5 - blank identifier in structs](http://developer20.com/golang-tips-and-trics-v/)

- [如何在 golang 中通过通道发送多个变量？](http://developer20.com/how-to-send-multiple-variables-via-channel-in-golang/)
- [Golang Tips & Tricks #7 - 私有仓库和代理](http://developer20.com/golang-tips-and-trics-vii/)
- [我如何在 Go 中组织包](http://developer20.com/how-i-organize-packages-in-go/)
- [Golang Tips & Tricks #6 - \_test 包](http://developer20.com/golang-tips-and-trics-vi/)
- [Golang Tips & Tricks #5 - 结构体中的空白标识符](http://developer20.com/golang-tips-and-trics-v/)

[←](http://developer20.com/golang-tips-and-trics-vii/)[→](http://developer20.com/writing-proxy-in-go/)Top

[←](http://developer20.com/golang-tips-and-trics-vii/)[→](http://developer20.com/writing-proxy-in-go/)顶部

© 2021 . Made with [Hugo](https://gohugo.io) using the [Tale](https://github.com/EmielH/tale-hugo/) theme. 

© 2021。 [Hugo](https://gohugo.io) 使用 [Tale](https://github.com/EmielH/tale-hugo/) 主题制作。

