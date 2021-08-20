# Don’t use Go’s default HTTP client (in production)

# 不要使用 Go 的默认 HTTP 客户端（在生产中）

Jan 23, 2016 4 min read From: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

Writing Go programs that talk to services over HTTP is easy and fun. I’ve  written numerous API client packages and I find it an enjoyable task. However, I have run into a pitfall that is easy to fall into and can  crash your program very quickly: the default HTTP client.

编写通过 HTTP 与服务通信的 Go 程序既简单又有趣。我编写了许多 API 客户端包，我发现这是一项令人愉快的任务。但是，我遇到了一个容易陷入并且会很快使您的程序崩溃的陷阱：默认的 HTTP 客户端。

**TL;DR**: Go’s http package doesn’t specify request timeouts by default, allowing services to hijack your goroutines. **Always specify a custom http.Client** when connecting to outside services.

**TL;DR**：Go 的 http 包默认不指定请求超时，允许服务劫持你的 goroutine。 **在连接到外部服务时始终指定自定义 http.Client**。

## The Problem by Example

## 示例问题

Let’s say you want to want to talk to spacely-sprockets.com via their nice JSON REST API and view a list of available sprockets. In Go, you might do something like:

假设您想通过他们漂亮的 JSON REST API 与 spacely-sprockets.com 交谈并查看可用链轮的列表。在 Go 中，您可能会执行以下操作：

```
 // error checking omitted for brevity
 var sprockets SprocketsResponse
 response, _ := http.Get("spacely-sprockets.com/api/sprockets")
 buf, _ := ioutil.ReadAll(response.Body)
 json.Unmarshal(buf, &sprockets)
```


You write your code (with proper error handling, please), compile, and run. Everything works great. Now, you take your API package and plug it into a web application. One page of your app shows a list of Spacely  Sprockets inventory to users by making a call to the API.

您编写代码（请进行适当的错误处理）、编译和运行。一切都很好。现在，您将 API 包插入到 Web 应用程序中。您的应用程序的一个页面通过调用 API 向用户显示 Spacely Sprockets 库存列表。

Everything is going great until one day your app stops responding. You look in the logs, but there’s nothing to indicate a problem. You check your  monitoring tools, but CPU, memory, and I/O all look reasonable leading  up to the outage. You spin up a sandbox and it seems to work fine. What  gives?

一切都很顺利，直到有一天您的应用停止响应。您查看了日志，但没有任何迹象表明存在问题。您检查了监控工具，但CPU、内存和I/O在导致停机之前看起来都是合理的。你启动了一个沙箱，它似乎工作正常。发生了什么？

Frustrated, you check Twitter and notice a tweet from the Spacely Sprockets dev  team saying that they experienced a brief outage, but that everything is now back to normal. You check their API status page and see that the  outage began a few minutes before yours. That seems like an unlikely  coincidence, but you can’t quite figure out how it’s related, since your API code handles errors gracefully. You’re still no closer to figuring  out the issue.

沮丧的是，您查看 Twitter 并注意到 Spacely Sprockets 开发团队的一条推文，称他们经历了短暂的中断，但现在一切都恢复正常了。您检查他们的 API 状态页面，并看到中断在您的几分钟前开始。这似乎是一个不太可能的巧合，但您无法弄清楚它是如何关联的，因为您的 API 代码会优雅地处理错误。你还没有更接近弄清楚这个问题。

## The Go HTTP package

## Go HTTP 包

Go’s HTTP package uses a struct called Client to manage the internals of  communicating over HTTP(S). Clients are concurrency-safe objects that  contain configuration, manage TCP state, handle cookies, etc. When you  use http.Get(url), you are using the http.DefaultClient, a package  variable that defines the default configuration for a client. The  declaration for this is:

Go 的 HTTP 包使用一个名为 Client 的结构来管理通过 HTTP(S) 进行通信的内部结构。客户端是并发安全对象，包含配置、管理 TCP 状态、处理 cookie 等。当您使用 http.Get(url) 时，您使用的是 http.DefaultClient，这是一个定义客户端默认配置的包变量。对此的声明是：

```
var DefaultClient = &Client{}
```


Among other things, http.Client configures a timeout that short-circuits  long-running connections. The default for this value is 0, which is  interpreted as “no timeout”. This may be a sensible default for the  package, but it is a nasty pitfall and the cause of our application  falling over in the above example. As it turns out, Spacely Sprockets’  API outage caused connection attempts to hang (this doesn’t always  happen, but it does in our example). They will continue to hang for as  long as the malfunctioning server decides to wait. Because API calls  were being made to serve user requests, this caused the goroutines  serving user requests to hang as well. Once enough users hit the  sprockets page, the app fell over, most likely due to resource limits  being reached.

除其他事项外，http.Client 配置了一个使长时间运行的连接短路的超时。此值的默认值为 0，这被解释为“无超时”。这对于包来说可能是一个合理的默认值，但它是一个令人讨厌的陷阱，也是我们的应用程序在上面的例子中失败的原因。事实证明，Spacely Sprockets 的 API 中断导致连接尝试挂起（这并不总是发生，但在我们的示例中确实发生）。只要发生故障的服务器决定等待，它们就会继续挂起。因为 API 调用是为服务用户请求而进行的，这导致服务用户请求的 goroutines 也挂起。一旦有足够多的用户访问 sprockets 页面，应用就会崩溃，这很可能是因为达到了资源限制。

Here is a simple Go program that demonstrates the issue:

这是一个演示该问题的简单 Go 程序：

```
package mainimport (
   “fmt”
   “net/http”
   “net/http/httptest”
   “time”
 )
 func main() {
   svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
     time.Sleep(time.Hour)
   }))
   defer svr.Close() 
   fmt.Println("making request")
   http.Get(svr.URL)
   fmt.Println("finished request")
 }
```


When run, this program will make a request to a server that will sleep for  an hour. Consequently, the program will wait for one hour and then exit.

运行时，该程序将向将休眠一小时的服务器发出请求。因此，程序将等待一小时然后退出。

## The Solution

## 解决方案

The solution to this problem is to always define an http.Client with a sensible timeout for your use case. Here is an example:

此问题的解决方案是始终为您的用例定义一个具有合理超时的 http.Client。下面是一个例子：

```
var netClient = &http.Client{
   Timeout: time.Second * 10,
 }
 response, _ := netClient.Get(url)
```


This sets a 10 second timeout on requests made to the endpoint. If the API  server exceeds the timeout, Get() will return with the error:

这将设置对端点发出的请求的 10 秒超时。如果 API 服务器超过超时，Get() 将返回错误：

```
&httpError{
   err:     err.Error() + " (Client.Timeout exceeded while awaiting headers)",
   timeout: true,
 }
```


If you need finer-grained control over the request lifecycle, you can  additionally specify a custom net.Transport and net.Dialer. A Transport  is a struct used by clients to manage the underlying TCP connection and  it’s Dialer is a struct that manages the establishment of the  connection. Go’s net package has a default Transport and Dialer as well. Here’s an example of using custom ones:

如果您需要对请求生命周期进行更细粒度的控制，您可以另外指定自定义 net.Transport 和 net.Dialer。 Transport 是客户端用来管理底层 TCP 连接的结构，它的 Dialer 是管理连接建立的结构。 Go 的 net 包也有一个默认的 Transport 和 Dialer。这是使用自定义的示例：

```
var netTransport = &http.Transport{
   Dial: (&net.Dialer{
     Timeout: 5 * time.Second,
   }).Dial,
   TLSHandshakeTimeout: 5 * time.Second,
 }
 var netClient = &http.Client{
   Timeout: time.Second * 10,
   Transport: netTransport,
 }
 response, _ := netClient.Get(url)
```


This code will cap the TCP connect and TLS handshake timeouts, as well as  establishing an end-to-end request timeout. There are other  configuration options such as keep-alive timeouts you can play with if  needed.

此代码将限制 TCP 连接和 TLS 握手超时，以及建立端到端请求超时。如果需要，您可以使用其他配置选项，例如保持活动超时。

## Conclusion

## 结论

Go’s net and http packages are a well-thought out, convenient base for  communicating over HTTP(S). However, the lack of a default timeout for  requests is an easy pitfall to fall into, because the package provides  convenience methods like http.Get(url). Some languages (e.g. Java) have  the same issue, others (e.g. Ruby has a default 60 second read timeout)  do not. Not setting a request timeout when contacting a remote service  puts your application at the mercy of that service. A malfunctioning or  malicious service can hang on to your connection forever, potentially  starving your application.

Go 的 net 和 http 包是一个经过深思熟虑的、方便的 HTTP(S) 通信基础。然而，缺少请求的默认超时是一个容易陷入的陷阱，因为该包提供了像 http.Get(url) 这样的便捷方法。一些语言（例如 Java）有同样的问题，其他语言（例如 Ruby 有一个默认的 60 秒读取超时）没有。联系远程服务时不设置请求超时会使您的应用程序受该服务的支配。出现故障或恶意的服务可能会永远挂在您的连接上，可能会使您的应用程序饿死。

[Nathan Smith](https://medium.com/@nate510?source=post_sidebar--------------------------post_sidebar-----------)
Software engineer, CTO of HER Social App. Social justice matters and you should care about it. 
软件工程师，HER社交应用CTO。社会正义很重要，你应该关心它。

