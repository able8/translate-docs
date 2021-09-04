# Writing a reverse proxy in Go

# 在 Go 中编写反向代理

Some time ago, I found a video called [Building a DIY proxy with the net package](https://www.youtube.com/watch?v=J4J-A9tcjcA). I recommend watching it. Filippo Valsorda builds a simple proxy using low-level packages. It’s fun to watch it but I think it’s a bit complicated. In Go, it has to be an easier way so I decided to continue writing series [Go In Practice](https://developer20.com/categories/GoInPractice/) by writing a simple but yet powerful reverse proxy as fast as it's possible .

前段时间发现了一个视频，叫做 [用网包搭建DIY代理](https://www.youtube.com/watch?v=J4J-A9tcjcA)。我推荐观看它。 Filippo Valsorda 使用低级包构建了一个简单的代理。看它很有趣，但我认为它有点复杂。在 Go 中，它必须是一种更简单的方法，所以我决定通过尽可能快地编写一个简单但功能强大的反向代理来继续编写系列 [Go In Practice](https://developer20.com/categories/GoInPractice/) .

The first step will be to create a proxy for a single host. The core of our code will be [ReversProxy](https://golang.org/pkg/net/http/httputil/#ReverseProxy) which does all the work for us. This is the magic of the rich standard library. The `RevewseProxy` is a struct for writing reverse proxies :) The only thing we have to do is to configure the director. The director modifies original requests which will be sent to proxied service.

第一步是为单个主机创建代理。我们代码的核心将是 [ReversProxy](https://golang.org/pkg/net/http/httputil/#ReverseProxy)，它为我们完成所有工作。这就是丰富的标准库的神奇之处。 `RevewseProxy` 是一个用于编写反向代理的结构:) 我们唯一要做的就是配置director。主管修改将发送到代理服务的原始请求。

```go
package main

import (
    "flag"
    "fmt"
    "net/http"
    "net/http/httputil"
)

func main() {
    url, err := url.Parse("http://localhost:8080")
      if err != nil {
          panic(err)
    }

    port := flag.Int("p", 80, "port")
    flag.Parse()

    director := func(req *http.Request) {
        req.URL.Scheme = url.Scheme
        req.URL.Host = url.Host
    }

    reverseProxy := &httputil.ReverseProxy{Director: director}
    handler := handler{proxy: reverseProxy}
    http.Handle("/", handler)

    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

type handler struct {
    proxy *httputil.ReverseProxy
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.proxy.ServeHTTP(w, r)
}
```


And that’s all! We have a fully functional proxy! Let’s check if it works. For the test, I wrote a simple server that returns the port it’s listening on.

就这样！我们有一个功能齐全的代理！让我们检查它是否有效。为了测试，我编写了一个简单的服务器，返回它正在侦听的端口。

```go
package main

import (
    "flag"
    "fmt"
    "net/http"
)

func main() {
    var port = flag.Int("p", 8080, "port")
    flag.Parse()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(fmt.Sprintf("hello on port %d", *port)))
    })
    err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
    fmt.Print(err)
}
```


Now, we can run the service which listens on port 8080 and the proxy which listens on port 80.

现在，我们可以运行监听 8080 端口的服务和监听 80 端口的代理。

![Working TCP proxy](http://developer20.com/images/proxy01.png)

Our proxy works on HTTP only. To tweak it to support HTTPS we have to make a small change. The thing we need to do is to detect (in a very naive way) if the proxy is running using SSL or not. We’ll detect it based on the port it’s running on.

我们的代理仅适用于 HTTP。要调整它以支持 HTTPS，我们必须做一个小改动。我们需要做的是检测（以一种非常天真的方式）代理是否正在使用 SSL 运行。我们将根据它运行的端口来检测它。

```go
if *port == 443 {
    http.ListenAndServeTLS(fmt.Sprintf(":%d", *port), "localhost.pem", "localhost-key.pem", handler)
} else {
    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
```


To make it working, we need one more thing: a valid certyficate. You can generate it using [https://github.com/FiloSottile/mkcert](http://developer20.com/mkcert).

为了使它工作，我们还需要一件事：有效的证书。您可以使用 [https://github.com/FiloSottile/mkcert](http://developer20.com/mkcert) 生成它。

```
mkcert localhost
```


And… that’s all! We have a proxy that works on both HTTP/HTTPS.
![Working TCP proxy](http://developer20.com/images/proxy02.png)

As you can see, writing the new tool was extremely simple and all we need is the standard library. We didn’t have to write complicated code so we’ll can focus on our real goals. As usual, the source code is available on Github: [https://github.com/bkielbasa/go-proxy](https://github.com/bkielbasa/go-proxy). This project will be a starting point for other tools so keep in touch and don’t miss a post! If you have any questions, let me know in the comments section below.

就这样！我们有一个适用于 HTTP/HTTPS 的代理。
如您所见，编写新工具非常简单，我们只需要标准库。我们不必编写复杂的代码，因此我们可以专注于我们的真正目标。像往常一样，源代码可以在 Github 上找到：[https://github.com/bkielbasa/go-proxy](https://github.com/bkielbasa/go-proxy)。该项目将成为其他工具的起点，因此请保持联系，不要错过任何帖子！如果您有任何疑问，请在下面的评论部分告诉我。

### See Also

###  也可以看看

- [Writing TCP scanner in Go](http://developer20.com/tcp-scanner-in-go/)
- [Golang Tips & Tricks #7 - private repository and proxy](http://developer20.com/golang-tips-and-trics-vii/)
- [How I organize packages in Go](http://developer20.com/how-i-organize-packages-in-go/)
- [Golang Tips & Tricks #6 - the \_test package](http://developer20.com/golang-tips-and-trics-vi/)
- [Golang Tips & Tricks #5 - blank identifier in structs](http://developer20.com/golang-tips-and-trics-v/) 

- [用 Go 编写 TCP 扫描器](http://developer20.com/tcp-scanner-in-go/)
- [Golang Tips & Tricks #7 - 私有仓库和代理](http://developer20.com/golang-tips-and-trics-vii/)
- [我如何在 Go 中组织包](http://developer20.com/how-i-organize-packages-in-go/)
- [Golang Tips & Tricks #6 - \_test 包](http://developer20.com/golang-tips-and-trics-vi/)
- [Golang Tips & Tricks #5 - 结构体中的空白标识符](http://developer20.com/golang-tips-and-trics-v/)

