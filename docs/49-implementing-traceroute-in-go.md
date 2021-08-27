# Implementing traceroute in Go

*([link](https://github.com/kalbhor/tracesite) for all the code)* August 18, 2020 From: https://blog.kalbhor.xyz/post/implementing-traceroute-in-go/

### What is traceroute?

If you’ve fiddled with networks you must be familiar with the famous `traceroute` tool. Its a script that traces the path to a host and prints info on every hop it encounters. To give an example if you run `traceroute kalbhor.xyz` you should see something like this :

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

Next we see `broadband.actcorp.in (49.207.47.201)` which  is my ISP. We can also see that my request forwards to a ISP router in  Delhi (most probably a regional level ISP) and further moves through  Chennai and Singapore (kalbhor.xyz is hosted on an AWS Singapore  server).

This tool is very useful to inspect network paths and solve problems. But aside from that, this tool is extremely interesting and its actual  implementation is pretty simple.

------

### How does traceroute work?

Now that we understand what traceroute does, lets take a look under  the hood. Every TCP/UDP packet that travels has a bunch of headers  containing info about the packet. One such header is the `ttl` header which is the number of hops the packet travels before being dropped. So if we set this `ttl` header to 1 our packet will reach the first hop and be dropped, if we  set it to 2 our packet will reach the second hop and drop, and so on.

Now that we know how our packets can reach any of the hops between us and our destination, how do we collect info on the hop? When a server/router drops a packet, it returns a  `ICMP Time Exceeded` message back. Parsing this message will allow us to retrieve info on  the particular hop. Once the destination is reached (last hop) we are  returned a `ICMP Destination Unreachable` message.

------

### Implementing traceroute

Now that we understand what’s happening under the hood, we can roughly design a way to implement traceroute. The steps to implement it should look something like this:

- Open a socket connection between us and our destination and send UDP packets
- Start from TTL=1 and keep increasing the TTL value on the UDP packets
- Open a socket that listens for the ICMP messages and parses them

------

### Writing a Go application that implements traceroute

Now we know what we want and all we need to do is implement it in any language. I’m implementing this in Go. The `net` and `syscall` package will help us along the way.

*Note: I will be using minimal code just to show the main  implementation (so you probably wont see me handling errors, etc here).  For a more refinded well developed version of this code check out [the repository](https://github.com/kalbhor/tracesite).*

Lets start by creating the sockets we’ll use for sending and recieving data.

```
sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
defer syscall.Close(recvSocket)
defer syscall.Close(sendSocket)
```

Lets create a ttl variable which we’ll iterate and a timevalue variable that defines our timeout

```
ttl := 1
// For 2000Ms
tv := syscall.NsecToTimeval(1000 * 1000 * (int64)(2000)) 
```

Next lets set the ttl and timeout value for the packets we’ll send in the socket

```
syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)
syscall.SetsockoptTimeval(recvSocket, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv)
```

At this point our sockets are ready to send and recieve data. What we need to do is find the destination address for our `sendSocket` and a network interface on our machine for our `recvSocket`

```
func socketAddr() ([4]byte, error) {
    socketAddr := [4]byte{0, 0, 0, 0}
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return socketAddr, err
    }

    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
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

```
destAddr, err := destAddr("google.com")
socketAddr, err := socketAddr()
```

Lets bind our `recvSocket` so that it can recieve messages and lets send a null byte to our destination through our `sendSocket`. We connect to the port 33434.

```
syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: 33434, Addr: socketAddr})
syscall.Sendto(sendSocket, []byte{0x0}, 0, &syscall.SockaddrInet4{Port: 33434, Addr: destAddr})
```

Now we need to parse the messages being sent on our `recvSocket`

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

We’re almost done! All we need to do is wrap this functionality in a for loop and keep updating the `ttl` variable.

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
        if ipString == destAddrString || ttl >= 56 { 
                break
        }
        ttl += 1
    }
}
```

Note that we added an if statement block to end our for loop once we reach the destination address or exceed max value for hops.

------

### Conclusion

This is definitely not the most elegant solution but it explains how simple the implementation of `traceroute` actually is. If you want to check out a more refinded version of this  code that compiles well and has many options like set ttl, max hops,  timeout, etc check out - [My Github Repo](https://github.com/kalbhor/tracesite)

##### Voila  💫  we just implemented the traceroute tool