[**Developer 2.0**](http://developer20.com/)

[Home](http://developer20.com/) [Newsletter](http://developer20.com/newsletter)

Categories

[Golang](http://developer20.com/categories/Golang/) [Books reviews](http://developer20.com/reviews/)

My projects

[GoBDD](http://developer20.com/projects/gobdd/)

[About me](http://developer20.com/about-me) [RSS](http://developer20.com/index.xml)

October 23, 2019
in category
[Golang](https://developer20.com/categories/Golang/) [GoInPractice](https://developer20.com/categories/GoInPractice/) [Programming](https://developer20.com/categories/Programming/)

# Writing TCP scanner in Go

Go is perfect for network applications. Its awesome standard library helps a lot in writing such software. In this article, we’ll write a simple TCP scanner in Go. The whole programm will take less than 50 lines of code. Before we’ll go to practice - a little theory.

Of course, the TCP is [more complicated than I describe](http://www.medianet.kent.edu/techreports/TR2005-07-22-tcp-EFSM.pdf) but we need just basics. The TCP handshake is three-way. Firstly, the client sends the `syn` package which signals the beginning of a communication. If the client gets a timeout here it may mean that the port is behind a firewall.

![syn package](http://developer20.com/images/diagram-01-sync.png)

Secondly, the server answers with `syn-ack` when the port is opened, otherwise it responses with `rst` package. In the end, the client has to send another packet called ack. From this point, the connection is established.

![syn package](http://developer20.com/images/diagram-02-sync-ack.png)![syn package](http://developer20.com/images/diagram-03-ack.png)

The first step in writing the TCP scanner is to test a single port. We’ll use the `net.Dial` function which accepts two parameters: the protocol and the address to test (with the port number).

```go
<span style="color:#f92672">package</span> <span style="color:#a6e22e">main</span>

<span style="color:#f92672">import</span> (
	<span style="color:#e6db74">"fmt"</span>
	<span style="color:#e6db74">"net"</span>
)

<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">main</span>() {
	<span style="color:#a6e22e">_</span>, <span style="color:#a6e22e">err</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">net</span>.<span style="color:#a6e22e">Dial</span>(<span style="color:#e6db74">"tcp"</span>, <span style="color:#e6db74">"google.com:80"</span>)
	<span style="color:#66d9ef">if</span> <span style="color:#a6e22e">err</span> <span style="color:#f92672">==</span> <span style="color:#66d9ef">nil</span> {
		<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Println</span>(<span style="color:#e6db74">"Connection successful"</span>)
	} <span style="color:#66d9ef">else</span> {
		<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Println</span>(<span style="color:#a6e22e">err</span>)
	}
}
```

To not test every port one by one, we’ll add a simple loop that will simplify the whole process. Notice the Sprintf function which concretes the host and the port.

```go
<span style="color:#f92672">package</span> <span style="color:#a6e22e">main</span>

<span style="color:#f92672">import</span> (
	<span style="color:#e6db74">"fmt"</span>
	<span style="color:#e6db74">"net"</span>
)

<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">main</span>() {
	<span style="color:#66d9ef">for</span> <span style="color:#a6e22e">port</span> <span style="color:#f92672">:=</span> <span style="color:#ae81ff">80</span>; <span style="color:#a6e22e">port</span> < <span style="color:#ae81ff">100</span>; <span style="color:#a6e22e">port</span><span style="color:#f92672">++</span> {
		<span style="color:#a6e22e">conn</span>, <span style="color:#a6e22e">err</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">net</span>.<span style="color:#a6e22e">Dial</span>(<span style="color:#e6db74">"tcp"</span>, <span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Sprintf</span>(<span style="color:#e6db74">"google.com:%d"</span>, <span style="color:#a6e22e">port</span>))
		<span style="color:#66d9ef">if</span> <span style="color:#a6e22e">err</span> <span style="color:#f92672">==</span> <span style="color:#66d9ef">nil</span> {
			<span style="color:#a6e22e">conn</span>.<span style="color:#a6e22e">Close</span>()
			<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Println</span>(<span style="color:#e6db74">"Connection successful"</span>)
		} <span style="color:#66d9ef">else</span> {
			<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Println</span>(<span style="color:#a6e22e">err</span>)
		}
	}
}
```

The solution has one huge issue - it’s extremely slow. We can do two things to make things faster: run those checks concurrently and add a timeout to every connection.

Let’s focus on making in concurrent. The first step is to extract the scanning to a separate function. This step will make our code more clear.

```go
<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">isOpen</span>(<span style="color:#a6e22e">host</span> <span style="color:#66d9ef">string</span>, <span style="color:#a6e22e">port</span> <span style="color:#66d9ef">int</span>) <span style="color:#66d9ef">bool</span> {
<span style="color:#a6e22e">conn</span>, <span style="color:#a6e22e">err</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">net</span>.<span style="color:#a6e22e">Dial</span>(<span style="color:#e6db74">"tcp"</span>, <span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Sprintf</span>(<span style="color:#e6db74">"%s:%d"</span>, <span style="color:#a6e22e">host</span>, <span style="color:#a6e22e">port</span>))
<span style="color:#66d9ef">if</span> <span style="color:#a6e22e">err</span> <span style="color:#f92672">==</span> <span style="color:#66d9ef">nil</span> {
     <span style="color:#a6e22e">_</span> = <span style="color:#a6e22e">conn</span>.<span style="color:#a6e22e">Close</span>()
     <span style="color:#66d9ef">return</span> <span style="color:#66d9ef">true</span>
}

<span style="color:#66d9ef">return</span> <span style="color:#66d9ef">false</span>
}
```

The only new thing is the `WaitGroup`. You can read about it in more detail [here](https://gobyexample.com/waitgroups) or, if you want (let me know in the comments below) I can write an article about async programming in Go. But, back to the topic… In the main function, we span our goroutines and wait for the execution to finish.

```go
<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">main</span>() {
<span style="color:#a6e22e">ports</span> <span style="color:#f92672">:=</span> []<span style="color:#66d9ef">int</span>{}

<span style="color:#a6e22e">wg</span> <span style="color:#f92672">:=</span> <span style="color:#f92672">&</span><span style="color:#a6e22e">sync</span>.<span style="color:#a6e22e">WaitGroup</span>{}
<span style="color:#66d9ef">for</span> <span style="color:#a6e22e">port</span> <span style="color:#f92672">:=</span> <span style="color:#ae81ff">1</span>; <span style="color:#a6e22e">port</span> < <span style="color:#ae81ff">100</span>; <span style="color:#a6e22e">port</span><span style="color:#f92672">++</span> {
     <span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Add</span>(<span style="color:#ae81ff">1</span>)
     <span style="color:#66d9ef">go</span> <span style="color:#66d9ef">func</span>() {
        <span style="color:#a6e22e">opened</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">isOpen</span>(<span style="color:#e6db74">"google.com"</span>, <span style="color:#a6e22e">port</span>)
        <span style="color:#66d9ef">if</span> <span style="color:#a6e22e">opened</span> {
           <span style="color:#a6e22e">ports</span> = append(<span style="color:#a6e22e">ports</span>, <span style="color:#a6e22e">port</span>)
        }
        <span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Done</span>()
     }()
}

<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Wait</span>()
<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Printf</span>(<span style="color:#e6db74">"opened ports: %v\n"</span>, <span style="color:#a6e22e">ports</span>)
}
```

Our code is faster but because of timeouts, we’re waiting a very long time to receive the error. We can assume that if we don’t get any response from the server for 200 ms we don’t want to wait longer.

```go
<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">isOpen</span>(<span style="color:#a6e22e">host</span> <span style="color:#66d9ef">string</span>, <span style="color:#a6e22e">port</span> <span style="color:#66d9ef">int</span>, <span style="color:#a6e22e">timeout</span> <span style="color:#a6e22e">time</span>.<span style="color:#a6e22e">Duration</span>) <span style="color:#66d9ef">bool</span> {
	<span style="color:#a6e22e">conn</span>, <span style="color:#a6e22e">err</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">net</span>.<span style="color:#a6e22e">DialTimeout</span>(<span style="color:#e6db74">"tcp"</span>, <span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Sprintf</span>(<span style="color:#e6db74">"%s:%d"</span>, <span style="color:#a6e22e">host</span>, <span style="color:#a6e22e">port</span>), <span style="color:#a6e22e">timeout</span>)
	<span style="color:#66d9ef">if</span> <span style="color:#a6e22e">err</span> <span style="color:#f92672">==</span> <span style="color:#66d9ef">nil</span> {
		<span style="color:#a6e22e">_</span> = <span style="color:#a6e22e">conn</span>.<span style="color:#a6e22e">Close</span>()
		<span style="color:#66d9ef">return</span> <span style="color:#66d9ef">true</span>
	}

	<span style="color:#66d9ef">return</span> <span style="color:#66d9ef">false</span>
}

<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">main</span>() {
	<span style="color:#a6e22e">ports</span> <span style="color:#f92672">:=</span> []<span style="color:#66d9ef">int</span>{}

	<span style="color:#a6e22e">wg</span> <span style="color:#f92672">:=</span> <span style="color:#f92672">&</span><span style="color:#a6e22e">sync</span>.<span style="color:#a6e22e">WaitGroup</span>{}
	<span style="color:#a6e22e">timeout</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">time</span>.<span style="color:#a6e22e">Millisecond</span> <span style="color:#f92672">*</span> <span style="color:#ae81ff">200</span>
	<span style="color:#66d9ef">for</span> <span style="color:#a6e22e">port</span> <span style="color:#f92672">:=</span> <span style="color:#ae81ff">1</span>; <span style="color:#a6e22e">port</span> < <span style="color:#ae81ff">100</span>; <span style="color:#a6e22e">port</span><span style="color:#f92672">++</span> {
		<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Add</span>(<span style="color:#ae81ff">1</span>)
		<span style="color:#66d9ef">go</span> <span style="color:#66d9ef">func</span>(<span style="color:#a6e22e">p</span> <span style="color:#66d9ef">int</span>) {
			<span style="color:#a6e22e">opened</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">isOpen</span>(<span style="color:#e6db74">"google.com"</span>, <span style="color:#a6e22e">p</span>, <span style="color:#a6e22e">timeout</span>)
			<span style="color:#66d9ef">if</span> <span style="color:#a6e22e">opened</span> {
				<span style="color:#a6e22e">ports</span> = append(<span style="color:#a6e22e">ports</span>, <span style="color:#a6e22e">p</span>)
			}
			<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Done</span>()
		}(<span style="color:#a6e22e">port</span>)
	}

	<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Wait</span>()
	<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Printf</span>(<span style="color:#e6db74">"opened ports: %v\n"</span>, <span style="color:#a6e22e">ports</span>)
}
```

At this point, we have a working simple port scanner. Unfortunately, it’s not very handy because to change the domain or port ranges we have to edit the code and recompile. Go has an awesome package called `flag`.

The `flag` package helps in writing command-line applications. You can read more about it in [Go by Example](https://gobyexample.com/command-line-flags). What we want is configuring every magic string or number. We add parameters for the hostname, port range we want to test and the timeout on the connection.

```go
<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">main</span>() {
	<span style="color:#a6e22e">hostname</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">flag</span>.<span style="color:#a6e22e">String</span>(<span style="color:#e6db74">"hostname"</span>, <span style="color:#e6db74">""</span>, <span style="color:#e6db74">"hostname to test"</span>)
	<span style="color:#a6e22e">startPort</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">flag</span>.<span style="color:#a6e22e">Int</span>(<span style="color:#e6db74">"start-port"</span>, <span style="color:#ae81ff">80</span>, <span style="color:#e6db74">"the port on which the scanning starts"</span>)
	<span style="color:#a6e22e">endPort</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">flag</span>.<span style="color:#a6e22e">Int</span>(<span style="color:#e6db74">"end-port"</span>, <span style="color:#ae81ff">100</span>, <span style="color:#e6db74">"the port from which the scanning ends"</span>)
	<span style="color:#a6e22e">timeout</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">flag</span>.<span style="color:#a6e22e">Duration</span>(<span style="color:#e6db74">"timeout"</span>, <span style="color:#a6e22e">time</span>.<span style="color:#a6e22e">Millisecond</span> <span style="color:#f92672">*</span> <span style="color:#ae81ff">200</span>, <span style="color:#e6db74">"timeout"</span>)
	<span style="color:#a6e22e">flag</span>.<span style="color:#a6e22e">Parse</span>()

	<span style="color:#a6e22e">ports</span> <span style="color:#f92672">:=</span> []<span style="color:#66d9ef">int</span>{}

	<span style="color:#a6e22e">wg</span> <span style="color:#f92672">:=</span> <span style="color:#f92672">&</span><span style="color:#a6e22e">sync</span>.<span style="color:#a6e22e">WaitGroup</span>{}
	<span style="color:#66d9ef">for</span> <span style="color:#a6e22e">port</span> <span style="color:#f92672">:=</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">startPort</span>; <span style="color:#a6e22e">port</span> <span style="color:#f92672"><=</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">endPort</span>; <span style="color:#a6e22e">port</span><span style="color:#f92672">++</span> {
		<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Add</span>(<span style="color:#ae81ff">1</span>)
		<span style="color:#66d9ef">go</span> <span style="color:#66d9ef">func</span>(<span style="color:#a6e22e">p</span> <span style="color:#66d9ef">int</span>) {
			<span style="color:#a6e22e">opened</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">isOpen</span>(<span style="color:#f92672">*</span><span style="color:#a6e22e">hostname</span>, <span style="color:#a6e22e">p</span>, <span style="color:#f92672">*</span><span style="color:#a6e22e">timeout</span>)
			<span style="color:#66d9ef">if</span> <span style="color:#a6e22e">opened</span> {
				<span style="color:#a6e22e">ports</span> = append(<span style="color:#a6e22e">ports</span>, <span style="color:#a6e22e">p</span>)
			}
			<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Done</span>()
		}(<span style="color:#a6e22e">port</span>)
	}

	<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Wait</span>()
	<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Printf</span>(<span style="color:#e6db74">"opened ports: %v\n"</span>, <span style="color:#a6e22e">ports</span>)
}
```

If we want to show the usage, we have to put the -h parameter which will show us the usage. Simple and clear. The whole project took less than 50 lines of code. We used concurrency, the flag, and net packages.

There’s one more thing. Our program has race condition. In only a few opened ports and so slow scanning it’s not visible but there’s the issue. To fix that we’ll add [a mutex](https://gobyexample.com/mutexes).

```go
	<span style="color:#a6e22e">wg</span> <span style="color:#f92672">:=</span> <span style="color:#f92672">&</span><span style="color:#a6e22e">sync</span>.<span style="color:#a6e22e">WaitGroup</span>{}
	<span style="color:#a6e22e">mutex</span> <span style="color:#f92672">:=</span> <span style="color:#f92672">&</span><span style="color:#a6e22e">sync</span>.<span style="color:#a6e22e">Mutex</span>{}
	<span style="color:#66d9ef">for</span> <span style="color:#a6e22e">port</span> <span style="color:#f92672">:=</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">startPort</span>; <span style="color:#a6e22e">port</span> <span style="color:#f92672"><=</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">endPort</span>; <span style="color:#a6e22e">port</span><span style="color:#f92672">++</span> {
		<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Add</span>(<span style="color:#ae81ff">1</span>)
		<span style="color:#66d9ef">go</span> <span style="color:#66d9ef">func</span>(<span style="color:#a6e22e">p</span> <span style="color:#66d9ef">int</span>) {
			<span style="color:#a6e22e">opened</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">isOpen</span>(<span style="color:#f92672">*</span><span style="color:#a6e22e">hostname</span>, <span style="color:#a6e22e">p</span>, <span style="color:#f92672">*</span><span style="color:#a6e22e">timeout</span>)
			<span style="color:#66d9ef">if</span> <span style="color:#a6e22e">opened</span> {
				<span style="color:#a6e22e">mutex</span>.<span style="color:#a6e22e">Lock</span>()
				<span style="color:#a6e22e">ports</span> = append(<span style="color:#a6e22e">ports</span>, <span style="color:#a6e22e">p</span>)
				<span style="color:#a6e22e">mutex</span>.<span style="color:#a6e22e">Unlock</span>()
			}
			<span style="color:#a6e22e">wg</span>.<span style="color:#a6e22e">Done</span>()
		}(<span style="color:#a6e22e">port</span>)
	}
```

If you like this kind of posts or have a question, let me know in the comments section below. The whole source code is available [on GitHub](https://github.com/bkielbasa/port-scanner).

[![Buy me a coffee](https://cdn.buymeacoffee.com/buttons/bmc-new-btn-logo.svg)Buy me a coffee](https://www.buymeacoffee.com/bklimczak)

Tags:
[#golang](https://developer20.com/tags/golang/) [#tcp](https://developer20.com/tags/tcp/) [#scanner](https://developer20.com/tags/scanner/) [#network](https://developer20.com/tags/network/) [#concurrency](https://developer20.com/tags/concurrency/)

### See Also

- [How to send multiple variables via channel in golang?](http://developer20.com/how-to-send-multiple-variables-via-channel-in-golang/)
- [Golang Tips & Tricks #7 - private repository and proxy](http://developer20.com/golang-tips-and-trics-vii/)
- [How I organize packages in Go](http://developer20.com/how-i-organize-packages-in-go/)
- [Golang Tips & Tricks #6 - the \_test package](http://developer20.com/golang-tips-and-trics-vi/)
- [Golang Tips & Tricks #5 - blank identifier in structs](http://developer20.com/golang-tips-and-trics-v/)

[←](http://developer20.com/golang-tips-and-trics-vii/) [→](http://developer20.com/writing-proxy-in-go/)Top

© 2021 . Made with [Hugo](https://gohugo.io) using the [Tale](https://github.com/EmielH/tale-hugo/) theme.

