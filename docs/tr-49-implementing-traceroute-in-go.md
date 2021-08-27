# Implementing traceroute in Go

# 在 Go 中实现 traceroute

*([link](https://github.com/kalbhor/tracesite)for all the code)* August 18, 2020 From: https://blog.kalbhor.xyz/post/implementing-traceroute-in-go/

### What is traceroute?

### 什么是跟踪路由？

If you’ve fiddled with networks you must be familiar with the famous `traceroute` tool. Its a script that traces the path to a host and prints info on every hop it encounters. To give an example if you run `traceroute kalbhor.xyz` you should see something like this :

如果你玩过网络，你一定熟悉著名的“traceroute”工具。它是一个脚本，用于跟踪主机的路径并在遇到的每个跃点上打印信息。举个例子，如果你运行 `traceroute kalbhor.xyz` 你应该看到这样的：

```
❯ traceroute kalbhor.xyz
traceroute to kalbhor.xyz (18.140.218.13), 64 hops max, 52 byte packets
 1  dlinkrouter.dlink (192.168.0.1)  2.035 ms  1.276 ms  1.097 ms
 2  10.194.0.1 (10.194.0.1)  5.985 ms  4.006 ms  3.817 ms
 3  broadband.actcorp.in (49.207.47.201)  4.320 ms  4.715 ms  4.243 ms
 4  broadband.actcorp.in (49.207.47.225)  5.115 ms  5.390 ms  4.893 ms
 5  14.142.187.85.static-delhi.vsnl.net.in (14.142.187.85)  3.789 ms  3.746 ms  4.004 ms
 6  172.31.180.57 (172.31.180.57)  40.903 ms  41.661 ms  41.531 ms
 7  * * ix-ae-4-2.tcore1.cxr-chennai.as6453.net (180.87.36.9)  177.280 ms
 8  if-ae-13-2.tcore1.svw-singapore.as6453.net (180.87.36.83)  164.288 ms  176.561 ms  82.274 ms
 9  180.87.106.5 (180.87.106.5)  81.871 ms  84.931 ms  83.477 ms
10  52.93.11.197 (52.93.11.197)  82.368 ms  84.777 ms
    52.93.11.211 (52.93.11.211)  82.945 ms
11  52.93.11.79 (52.93.11.79)  83.587 ms
    52.93.11.67 (52.93.11.67)  78.292 ms
    52.93.11.87 (52.93.11.87)  79.452 ms
12  52.93.11.80 (52.93.11.80)  82.862 ms
    52.93.11.82 (52.93.11.82)  86.355 ms
    52.93.11.72 (52.93.11.72)  88.732 ms
13  52.93.9.161 (52.93.9.161)  83.706 ms
    52.93.9.95 (52.93.9.95)  82.498 ms
    52.93.9.139 (52.93.9.139)  84.551 ms
14  203.83.223.77 (203.83.223.77)  84.500 ms
    52.93.10.95 (52.93.10.95)  79.663 ms  79.812 ms
```


These might differ for you but for me this is the route my computer takes to connect to `kalbhor.xyz`. A few interesting details here include `dlinkrouter.dlink (192.168.0.1)`. Yes, that looks similar! It is my routers local IP, which means my  router at home is the first machine to process my request. That’s pretty obvious.

这些对你来说可能会有所不同，但对我来说，这是我的计算机连接到`kalbhor.xyz`的路径。一些有趣的细节包括`dlinkrouter.dlink (192.168.0.1)`。是的，看起来很像！这是我的路由器本地IP，这意味着我家里的路由器是第一台处理我的请求的机器。这很明显。

Next we see `broadband.actcorp.in (49.207.47.201)` which  is my ISP. We can also see that my request forwards to a ISP router in  Delhi (most probably a regional level ISP) and further moves through  Chennai and Singapore (kalbhor.xyz is hosted on an AWS Singapore  server).

接下来我们看到`broadband.actcorp.in (49.207.47.201)`，它是我的ISP。我们还可以看到，我的请求转发到德里的 ISP 路由器（很可能是区域级 ISP），并进一步通过金奈和新加坡（kalbhor.xyz 托管在 AWS 新加坡服务器上）。

This tool is very useful to inspect network paths and solve problems. But aside from that, this tool is extremely interesting and its actual  implementation is pretty simple.

该工具对于检查网络路径和解决问题非常有用。但除此之外，这个工具非常有趣，它的实际实现非常简单。

------

### How does traceroute work?

### traceroute 是如何工作的？

Now that we understand what traceroute does, lets take a look under  the hood. Every TCP/UDP packet that travels has a bunch of headers  containing info about the packet. One such header is the `ttl` header which is the number of hops the packet travels before being dropped. So if we set this `ttl` header to 1 our packet will reach the first hop and be dropped, if we  set it to 2 our packet will reach the second hop and drop, and so on.

现在我们了解了 traceroute 的作用，让我们来看看幕后。每个传输的 TCP/UDP 数据包都有一堆包含有关数据包信息的标头。一个这样的标头是 `ttl` 标头，它是数据包在被丢弃之前所经过的跳数。因此，如果我们将此 `ttl` 标头设置为 1，我们的数据包将到达第一跳并被丢弃，如果我们将其设置为 2，我们的数据包将到达第二跳并丢弃，依此类推。

Now that we know how our packets can reach any of the hops between us and our destination, how do we collect info on the hop? When a server/router drops a packet, it returns a  `ICMP Time Exceeded` message back. Parsing this message will allow us to retrieve info on  the particular hop. Once the destination is reached (last hop) we are  returned a `ICMP Destination Unreachable` message.

现在我们知道我们的数据包如何到达我们和目的地之间的任何跃点，我们如何收集跃点的信息？当服务器/路由器丢弃数据包时，它会返回“ICMP Time Exceeded”消息。解析此消息将允许我们检索特定跃点的信息。一旦到达目的地（最后一跳），我们就会返回一条“ICMP 目的地无法到达”消息。

------

### Implementing traceroute

### 实现跟踪路由

Now that we understand what’s happening under the hood, we can roughly design a way to implement traceroute. The steps to implement it should look something like this:

现在我们了解了幕后发生的事情，我们可以粗略地设计一种实现 traceroute 的方法。实现它的步骤应该是这样的：

- Open a socket connection between us and our destination and send UDP packets
- Start from TTL=1 and keep increasing the TTL value on the UDP packets
- Open a socket that listens for the ICMP messages and parses them

- 打开我们和目的地之间的套接字连接并发送 UDP 数据包
- 从 TTL=1 开始，不断增加 UDP 数据包的 TTL 值
- 打开一个监听 ICMP 消息并解析它们的套接字

------

### Writing a Go application that implements traceroute

### 编写一个实现 traceroute 的 Go 应用程序

Now we know what we want and all we need to do is implement it in any language. I’m implementing this in Go. The `net` and `syscall` package will help us along the way. 

现在我们知道我们想要什么，我们需要做的就是用任何语言实现它。我正在 Go 中实现它。 `net` 和 `syscall` 包将一路帮助我们。

*Note: I will be using minimal code just to show the main  implementation (so you probably wont see me handling errors, etc here). For a more refinded well developed version of this code check out [the repository](https://github.com/kalbhor/tracesite).*

*注意：我将使用最少的代码来显示主要实现（所以你可能不会在这里看到我处理错误等）。有关此代码的更完善的开发版本，请查看 [存储库](https://github.com/kalbhor/tracesite)。*

Lets start by creating the sockets we’ll use for sending and recieving data.

让我们从创建用于发送和接收数据的套接字开始。

```
sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
defer syscall.Close(recvSocket)
defer syscall.Close(sendSocket)
```


Lets create a ttl variable which we’ll iterate and a timevalue variable that defines our timeout

让我们创建一个我们将迭代的 ttl 变量和一个定义超时的 timevalue 变量

```
ttl := 1
// For 2000Ms
tv := syscall.NsecToTimeval(1000 * 1000 * (int64)(2000))
```


Next lets set the ttl and timeout value for the packets we’ll send in the socket

接下来让我们为我们将在套接字中发送的数据包设置 ttl 和超时值

```
syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)
syscall.SetsockoptTimeval(recvSocket, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv)
```


At this point our sockets are ready to send and recieve data. What we need to do is find the destination address for our `sendSocket` and a network interface on our machine for our `recvSocket`

此时我们的套接字已准备好发送和接收数据。我们需要做的是找到我们的 `sendSocket` 的目标地址和我们机器上用于我们的 `recvSocket` 的网络接口

```
func socketAddr() ([4]byte, error) {
    socketAddr := [4]byte{0, 0, 0, 0}
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return socketAddr, err
    }

    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet);ok && !ipnet.IP.IsLoopback() {
            if len(ipnet.IP.To4()) == net.IPv4len {
                copy(socketAddr[:], ipnet.IP.To4())
                return socketAddr, nil
            }
        }
    }
    err = errors.New("Not connected to the Internet")
    return socketAddr, err
}

func destAddr(dest string) ([4]byte, error) {
    destAddr := [4]byte{0, 0, 0, 0}
    addrs, err := net.LookupHost(dest)
    if err != nil {
        return destAddr, err
    }
    addr := addrs[0]

    ipAddr, err := net.ResolveIPAddr("ip", addr)
    if err != nil {
        return destAddr, err
    }
    copy(destAddr[:], ipAddr.IP.To4())
    return destAddr, nil
}
```


And in our main function we use these functions the get the addresses our sockets will use

在我们的主函数中，我们使用这些函数来获取我们的套接字将使用的地址

```
destAddr, err := destAddr("google.com")
socketAddr, err := socketAddr()
```


Lets bind our `recvSocket` so that it can recieve messages and lets send a null byte to our destination through our `sendSocket`. We connect to the port 33434.

让我们绑定我们的 `recvSocket` 以便它可以接收消息并让我们通过我们的 `sendSocket` 向我们的目的地发送一个空字节。我们连接到端口 33434。

```
syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: 33434, Addr: socketAddr})
syscall.Sendto(sendSocket, []byte{0x0}, 0, &syscall.SockaddrInet4{Port: 33434, Addr: destAddr})
```


Now we need to parse the messages being sent on our `recvSocket`

现在我们需要解析在我们的 `recvSocket` 上发送的消息

```
p := make([]byte, options.Int(56)) // The integer here is the packet size
n, from, err := syscall.Recvfrom(recvSocket, p, 0)

ip := from.(*syscall.SockaddrInet4).Addr
ipString := fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
host, err := net.LookupAddr(ipString)

fmt.Println(host)
fmt.Println(ipString)
```


The Recvfrom method returns a `Sockaddr` type to our `from` variable. Hence if we parse our `from` variable we can get the IP info on the hop. We can use this with `net.LookupAddr` to run a reverse search and get the hostname (domain name) through the IP.

Recvfrom 方法向我们的 `from` 变量返回一个 `Sockaddr` 类型。因此，如果我们解析我们的 `from` 变量，我们就可以获取跃点的 IP 信息。我们可以将它与 `net.LookupAddr` 一起使用来运行反向搜索并通过 IP 获取主机名（域名）。

We’re almost done! All we need to do is wrap this functionality in a for loop and keep updating the `ttl` variable.

我们快完成了！我们需要做的就是将此功能包装在 for 循环中并不断更新 `ttl` 变量。

```
func main() {
    sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
    recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
    defer syscall.Close(recvSocket)
    defer syscall.Close(sendSocket)

    ttl := 1
    tv := syscall.NsecToTimeval(1000 * 1000 * (int64)(2000)) // For 2000Ms

    for {
        syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)
        syscall.SetsockoptTimeval(recvSocket, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv)

        destAddr, err := destAddr("google.com")
        socketAddr, err := socketAddr()
        destAddrString := fmt.Sprintf("%v.%v.%v.%v", destAddr[0], destAddr[1], destAddr[2], destAddr[3])


        syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: 33434, Addr: socketAddr})
        syscall.Sendto(sendSocket, []byte{0x0}, 0, &syscall.SockaddrInet4{Port: 33434, Addr: destAddr})

        p := make([]byte, options.Int(56)) // The integer here is the packet size
        n, from, err := syscall.Recvfrom(recvSocket, p, 0)

        ip := from.(*syscall.SockaddrInet4).Addr
        ipString := fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
        host, err := net.LookupAddr(ipString)
        
        fmt.Println(host)
        fmt.Println(ipString)
        
        // We stop our loop if we reach destination or reach max value for ttl
        if ipString == destAddrString ||ttl >= 56 {
                break
        }
        ttl += 1
    }
}
```




Note that we added an if statement block to end our for loop once we reach the destination address or exceed max value for hops.

请注意，我们添加了一个 if 语句块以在我们到达目标地址或超过跳数的最大值时结束我们的 for 循环。

------

### Conclusion

###  结论

This is definitely not the most elegant solution but it explains how simple the implementation of `traceroute` actually is. If you want to check out a more refinded version of this  code that compiles well and has many options like set ttl, max hops,  timeout, etc check out - [My Github Repo](https://github.com/kalbhor/tracesite)

这绝对不是最优雅的解决方案，但它解释了 `traceroute` 的实现实际上是多么简单。如果您想查看此代码的更完善的版本，该版本可以很好地编译并且具有许多选项，例如设置 ttl、最大跳数、超时等，请查看 - [My Github Repo](https://github.com/kalbhor/tracesite)

##### Voila  💫  we just implemented the traceroute tool 

##### 瞧💫 我们刚刚实现了traceroute 工具

