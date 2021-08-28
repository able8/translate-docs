# Increasing http.Server boilerplate

# 增加 http.Server 样板

October 2020 From: https://bojanz.github.io/increasing-http-server-boilerplate-go/

One great feature of Go is the built-in  http.Server. It allows each app to serve HTTP and HTTPS traffic without  having to put a reverse proxy such as Nginx in front of it.

Go 的一大特色是内置的 http.Server。它允许每个应用程序为 HTTP 和 HTTPS 流量提供服务，而无需在其前面放置反向代理，例如 Nginx。

At a glance the API is simple:

一目了然，API 很简单：

```c
http.ListenAndServe(":8080", h)
```


where *h* is http.ServeMux or a third party router such as [Chi](https://github.com/go-chi/chi). But as always, the devil is in the details. Handling these details will require some boilerplate, so let’s start writing it.

其中 *h* 是 http.ServeMux 或第三方路由器，例如 [Chi](https://github.com/go-chi/chi)。但一如既往，细节决定成败。处理这些细节需要一些样板，所以让我们开始编写它。

### Production-ready configuration (timeouts, TLS)

### 生产就绪配置（超时，TLS）

[ListenAndServe](https://golang.org/src/net/http/server.go?s=97511:97566#L3108) creates an http.Server and uses it to listen on the given address and serve the given handler:

[ListenAndServe](https://golang.org/src/net/http/server.go?s=97511:97566#L3108) 创建一个 http.Server 并使用它来监听给定的地址并为给定的处理程序提供服务：

```c
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
```


However, the instantiated http.Server is not  production ready. It is missing important timeouts which can lead to resource exhaustion. The TLS configuration is optimized neither for  speed nor security. All of this is covered in a famous blog post by  Cloudflare titled [So you want to expose Go on the Internet](https://blog.cloudflare.com/exposing-go-on-the-internet/).

但是，实例化的 http.Server 尚未做好生产准备。它缺少可能导致资源耗尽的重要超时。 TLS 配置既没有针对速度也没有针对安全性进行优化。所有这些都在 Cloudflare 的一篇著名博客文章中进行了介绍，标题为 [所以你想在互联网上公开 Go](https://blog.cloudflare.com/exposing-go-on-the-internet/)。

So, how does a well configured server look according to Cloudflare?

那么，根据 Cloudflare，配置良好的服务器如何？

```c
func NewServer(addr string, handler http.Handler) *http.Server {
    return &http.Server{
        Addr:    addr,
        Handler: handler,
        // https://blog.cloudflare.com/exposing-go-on-the-internet/
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
        TLSConfig: &tls.Config{
            NextProtos:       []string{"h2", "http/1.1"},
            MinVersion:       tls.VersionTLS12,
            CurvePreferences: []tls.CurveID{tls.CurveP256, tls.X25519},
            CipherSuites: []uint16{
                tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
                tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
                tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
                tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
                tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
                tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
            },
            PreferServerCipherSuites: true,
        }
}
```


Usage stays similar:

用法保持相似：

```c
server := NewServer(":8080", r)
server.ListenAndServe()
```


Our next step is to allow the server to (optionally) listen on a systemd socket.

我们的下一步是允许服务器（可选）侦听 systemd 套接字。

### Systemd

### 系统

There are two broad ways in which a Go app is deployed: containerized or native.

部署 Go 应用程序有两种广泛的方式：容器化或原生。

A containerized app is put in a container and then deployed to the cloud, which can be anything from a Kubernetes setup to an IaaS provider like [Heroku](https://www.heroku.com/)or [Platform.sh]( https://platform.sh/).
  However, not every deployment needs the complexity that the containerized approach brings. One can get very far with a single VPS or dedicated server. I am a strong believer in continuing to support the “$5 Digital Ocean” crowd, especially now that Go has given us extra performance compared to the old PHP days.

将容器化应用程序放入容器中，然后部署到云中，可以是从 Kubernetes 设置到 IaaS 提供商的任何内容，例如 [Heroku](https://www.heroku.com/) 或 [Platform.sh]( https://platform.sh/）。
 然而，并非每个部署都需要容器化方法带来的复杂性。使用单个 VPS 或专用服务器可以走得很远。我坚信继续支持“5 美元的数字海洋”人群，尤其是现在与旧的 PHP 时代相比，Go 为我们提供了额外的性能。

A native deployment usually means Linux, which is nowadays powered by [systemd](https://www.digitalocean.com/community/tutorials/systemd-essentials-working-with-services-units-and-the-journal).Systemd will automatically start our app and bind it to the specified port, restart on failure, and redirect logs from stderr to syslog or [journald](https://sematext.com/blog/journald-logging-tutorial/). When redeploying our app, during the 1-2s downtime window, systemd will queue up any incoming requests, ensuring zero downtime deploys.

本机部署通常是指 Linux，它现在由 [systemd](https://www.digitalocean.com/community/tutorials/systemd-essentials-working-with-services-units-and-the-journal)提供支持。 Systemd 将自动启动我们的应用程序并将其绑定到指定的端口，失败时重新启动，并将日志从 stderr 重定向到 syslog 或 [journald](https://sematext.com/blog/journald-logging-tutorial/)。在重新部署我们的应用程序时，在 1-2 秒的停机时间窗口内，systemd 会将所有传入请求排队，确保零停机时间部署。

This sounds great, but it requires a bit of adaptation on our side. Aside from having to ship two systemd config files (a unit file and a socket file), the app also needs to be able to listen on a systemd socket.

这听起来很棒，但它需要我们进行一些调整。除了必须发送两个 systemd 配置文件（一个单元文件和一个套接字文件）之外，应用程序还需要能够侦听 systemd 套接字。

Let's assume that *addr* defaults to a TCP address such as “:8080”, but can also be set to a systemd socket name such as “systemd:myapp-http”, preferably through an environment variable which can be defined in our unit file.

让我们假设 *addr* 默认为 TCP 地址，例如“:8080”，但也可以设置为 systemd 套接字名称，例如“systemd:myapp-http”，最好通过可以在我们的单元中定义的环境变量文件。

With a little help from [coreos/go-systemd](https://github.com/coreos/go-systemd), a helper is born:

在 [coreos/go-systemd](https://github.com/coreos/go-systemd) 的一点帮助下，一个助手诞生了：

```c
func Listen(addr string) (net.Listener, error) {
    var ln net.Listener
    if strings.HasPrefix(addr, "systemd:") {
        name := addr[8:]
        listeners, _ := activation.ListenersWithNames()
        listener, ok := listeners[name]
        if !ok {
            return nil, fmt.Errorf("listen systemd %s: socket not found", name)
        }
        ln = listener[0]
    } else {
        var err error
        ln, err = net.Listen("tcp", addr)
        if err != nil {
            return nil, err
        }
    }

    return ln, nil
}
```




Usage now looks like this:

现在的用法如下所示：

```c
addr := os.GetEnv("LISTEN")
if addr == "" {
    addr = ":8080"
}
server := NewServer(addr, r)
ln, err := Listen(addr)
if err != nil {
    // Handle the error.
}
server.Serve(ln)
```


Having to pass *addr* twice and call *Listen()* ourselves is a bit tedious. Let’s define our own Server struct which embeds **http.Server*, and move the listener logic there:

必须通过 *addr* 两次并自己调用 *Listen()* 有点乏味。让我们定义我们自己的 Server 结构，它嵌入了 **http.Server*，并将监听器逻辑移到那里：

```c
package httpx

type Server struct {
    *http.Server
}

func NewServer(addr string, handler http.Handler) *Server {}

func (srv *Server) Listen() (net.Listener, error) {
    // Same code as before, but now using srv.Addr
}

func (srv *Server) ListenAndServe() error {
    ln, err := srv.Listen()
    if err != nil {
        return err
    }
    return srv.Serve(ln)
}

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
    ln, err := srv.Listen()
    if err != nil {
        return err
    }
    return srv.ServeTLS(ln, certFile, keyFile)
}
```


Usage is now simple again:

用法现在又简单了：

```c
addr := os.GetEnv("LISTEN")
if addr == "" {
    addr = ":8080"
}
server := NewServer(addr, r)
server.ListenAndServe()
```


### TLS

### TLS

Don’t we live in an HTTPS world? So far we’ve used *ListenAndServe* and *Serve*, not *ListenAndServeTLS* and *ServeTLS*. Can we just add those three missing letters, point to the certificate, modify the port, and call it a day?

我们不是生活在 HTTPS 世界中吗？到目前为止，我们使用了 *ListenAndServe* 和 *Serve*，而不是 *ListenAndServeTLS* 和 *ServeTLS*。我们可以只添加那三个缺失的字母，指向证书，修改端口，然后就可以了？

Yes, if we’re just serving an API. But if we’re serving HTML, we still need both HTTP and HTTPS, otherwise we won’t be able to visit our URL via the browser without supplying the HTTPS port. The job of the HTTP server is to redirect users to the HTTPS resource.

是的，如果我们只是提供 API。但是如果我们提供 HTML，我们仍然需要 HTTP 和 HTTPS，否则我们将无法在不提供 HTTPS 端口的情况下通过浏览器访问我们的 URL。 HTTP 服务器的工作是将用户重定向到 HTTPS 资源。

That redirect logic looks like this:

该重定向逻辑如下所示：

```c
type httpRedirectHandler struct{}

func (h httpRedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    host, _, err := net.SplitHostPort(r.Host)
    if err != nil {
        // No port found.
        host = r.Host
    }
    r.URL.Host = host
    r.URL.Scheme = "https"

    w.Header().Set("Connection", "close")
    http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
}
```


Each *Serve* call is blocking, so the two servers must run in their own goroutines. Both goroutines need to complete, and the idiomatic way to do that is using a [WaitGroup](https://golang.org/pkg/sync/#WaitGroup):

每个 *Serve* 调用都是阻塞的，因此两个服务器必须在自己的 goroutine 中运行。两个 goroutine 都需要完成，惯用的方法是使用 [WaitGroup](https://golang.org/pkg/sync/#WaitGroup)：

```c
mainServer := NewServer(":443", r)
redirectServer := NewServer(":80", httpRedirectHandler{})

wg := sync.WaitGroup{}
wg.Add(2)
go func() {
    mainServer.ListenAndServeTLS(certFile, keyFile)
    wg.Done()
}()
go func() {
    redirectServer.ListenAndServe()
    wg.Done()
}()

wg.Wait()
```


There’s only one detail missing now: error handling. If one of the servers errors out (couldn’t bind to addr or load the certificate) we want to make sure the other one is immediately stopped, and execution stops.

现在只缺少一个细节：错误处理。如果其中一台服务器出错（无法绑定到 addr 或加载证书），我们希望确保另一台服务器立即停止，并停止执行。

Ideally we’d get the error from *wg.Wait*, but it doesn’t support that. The answer lies in [x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup), which builds upon WaitGroup and does just that, in only 60 lines of code.

理想情况下，我们会从 *wg.Wait* 得到错误，但它不支持。答案在 [x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) 中，它建立在 WaitGroup 的基础上，仅用 60 行代码就可以做到这一点。

Here’s our code with error handling:

这是我们的错误处理代码：

```c
mainServer := NewServer(":443", r)
redirectServer := NewServer(":80", httpRedirectHandler{})

g, ctx := errgroup.WithContext(context.Background())
g.Go(func() error {
    if err := mainServer.ListenAndServeTLS(certFile, keyFile);err != http.ErrServerClosed {
        return err
    }
    return nil
})
g.Go(func() error {
    if err := redirectServer.ListenAndServe();err != http.ErrServerClosed {
        return err
    }
    return nil
})
go func() {
    // The context is closed if both servers finish, or one of them
    // errors out, in which case we want to close the other and return.
    <-ctx.Done()
    mainServer.Close()
    redirectServer.Close()
}()

err := g.Wait()
```


Note how we distinguish a real error from *http.ErrServerClosed*. We don’t want to call *Close* for *http.ErrServerClosed* because it would interfere with graceful shutdown.

请注意我们如何将真正的错误与 *http.ErrServerClosed* 区分开来。我们不想为 *http.ErrServerClosed* 调用 *Close*，因为它会干扰正常关闭。

The next tweak is more subjective. I dislike the fact that *certFile* and *keyFile* are passed when starting the server and not when initializing it. I would prefer having one way to start the server regardless of whether it uses TLS or not.

下一个调整更加主观。我不喜欢 *certFile* 和 *keyFile* 在启动服务器时传递而不是在初始化它时传递的事实。无论是否使用 TLS，我都希望有一种方式来启动服务器。

Let’s add a few more helpers to httpx:

让我们为 httpx 添加更多帮助器：

```c
func NewServerTLS(addr string, cert tls.Certificate, handler http.Handler) *Server {
    srv := NewServer(addr, handler)
    srv.TLSConfig.Certificates = []tls.Certificate{cert}

    return srv
}

func (srv *Server) IsTLS() bool {
    return len(srv.TLSConfig.Certificates) > 0 ||srv.TLSConfig.GetCertificate != nil
}

func (srv *Server) Start() error {
    ln, err := srv.Listen()
    if err != nil {
        return err
    }
    if srv.IsTLS() {
        ln = tls.NewListener(ln, srv.TLSConfig)
    }
    return srv.Serve(ln)
}
```




Our final implementation now looks like this:

我们的最终实现现在看起来像这样：

```c
cert, err := tls.LoadX509KeyPair(certFile, keyFile)
if err != nil {
    // Log the error and stop here.
}
mainServer := NewServerTLS(":443", cert, r)
redirectServer := NewServer(":80", httpRedirectHandler{})

g, ctx := errgroup.WithContext(context.Background())
g.Go(func() error {
    if err := mainServer.Start();err != http.ErrServerClosed {
        return err
    }
    return nil
})
g.Go(func() error {
    if err := redirectServer.Start();err != http.ErrServerClosed {
        return err
    }
    return nil
})
go func() {
    // The context is closed if both servers finish, or one of them
    // errors out, in which case we want to close the other and return.
    <-ctx.Done()
    mainServer.Close()
    redirectServer.Close()
}()

err := g.Wait()
```


### Graceful shutdown

### 优雅关机

We have talked about how to start the servers, but not how to shut them down. When a shutdown signal is received (SIGINT or SIGTERM), we want to shut down the servers in the opposite order from which we started them, first the redirect server then the main server. This will allow any in progress requests to complete:

我们已经讨论了如何启动服务器，但没有讨论如何关闭它们。当收到关闭信号（SIGINT 或 SIGTERM）时，我们希望按照与启动它们相反的顺序关闭服务器，首先是重定向服务器，然后是主服务器。这将允许任何正在进行的请求完成：

```c
redirectTimeout := 1 * time.Second
ctx, cancel := context.WithTimeout(context.Background(), redirectTimeout)
defer cancel()
if err := redirectServer.Shutdown(ctx);err == context.DeadlineExceeded {
    return fmt.Errorf("%v timeout exceeded while waiting on HTTP shutdown", redirectTimeout)
}
mainTimeout := 5 * time.Second
ctx, cancel := context.WithTimeout(context.Background(), mainTimeout)
defer cancel()
if err := mainServer.Shutdown(ctx);err == context.DeadlineExceeded {
    return fmt.Errorf("%v timeout exceeded while waiting on HTTPS shutdown", mainTimeout)
}
```


It is tempting to make each Server responsible for catching the shutdown signal and shutting down automatically, but that  would make it impossible to control the shutdown order. So, no new  helpers here. Instead, I like to create an Application struct, with its own Start()  and Shutdown() methods containing the code shown here. In addition to  starting and shutting down servers, these methods can also handle  app-specific workers such as queue processors.

让每个服务器负责捕捉关闭信号并自动关闭是很诱人的，但这将无法控制关闭顺序。所以，这里没有新帮手。相反，我喜欢创建一个 Application 结构，它自己的 Start() 和 Shutdown() 方法包含此处显示的代码。除了启动和关闭服务器之外，这些方法还可以处理特定于应用程序的工作程序，例如队列处理器。

The main package is then the one responsible for tying it all together:

主包是负责将它们捆绑在一起的包：

```c
     // Initialize dependencies, pass them to the Application.
    logger := NewLogger()
    app := myapp.New(logger)

    // Wait for shut down in a separate goroutine.
    errCh := make(chan error)
    go func() {
        shutdownCh := make(chan os.Signal)
        signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)
        <-shutdownCh

        errCh <- app.Shutdown()
    }()

    // Start the server and handle any errors.
    if err := app.Start();err != nil {
        logger.Fatal().Msg(err.Error())
    }
    // Handle shutdown errors.
    if err := <-errCh;err != nil {
        logger.Warn().Msg(err.Error())
    }
```


### Conclusion

###  结论

A simple microservice deployed to a known place can keep its code simple. A larger and more generic app needs more boilerplate. Luckily, it’s a problem that is easy to solve.

部署到已知位置的简单微服务可以保持其代码简单。更大、更通用的应用程序需要更多样板。幸运的是，这是一个很容易解决的问题。

I have gathered the httpx code shared here and published it as [bojanz/httpx](https://github.com/bojanz/httpx).The README has working examples of systemd unit and socket files. The code itself is only a hundred lines long (without comments), so I  encourage those unenthusiastic about introducing another dependency to  just copy the code into their project. After all, a little copying is better than a little dependency. 

我收集了此处共享的 httpx 代码并将其发布为 [bojanz/httpx](https://github.com/bojanz/httpx)。 README 中有 systemd 单元和套接字文件的工作示例。代码本身只有一百行（没有注释），所以我鼓励那些不热衷于引入另一个依赖项的人将代码复制到他们的项目中。毕竟，一点点复制比一点点依赖要好。

