# A Gentle Introduction to Web Services With Go
Sep 8, 2020 From: https://www.honeybadger.io/blog/go-web-services/

When you're deciding on a technology to use  for your project, it helps to have a broad understanding of your  options. You may be tempted to build a web service in Go for performance reasons - but what would that code actually look like? How would it  compare to languages like Ruby or JS? In this article, Ayooluwa Isaiah  gives us a guided tour through the building blocks of go web services so you'll be well-informed.

# 使用 Go 简单介绍 Web 服务
当您决定用于项目的技术时，对您的选择有一个广泛的了解是有帮助的。出于性能原因，您可能很想在 Go 中构建 Web 服务 - 但该代码实际上是什么样的？它与 Ruby 或 JS 等语言相比如何？在本文中，Ayoluwa Isaiah 向我们介绍了 Go Web 服务的构建块，以便您了解情况。

You may have heard that Go is great for web development. This claim is evidenced by several cloud services, such as Cloudflare, Digital Ocean, and others, betting on Go for their cloud infrastructure. But, what makes Go such a compelling option for building network-facing applications?

您可能听说过 Go 非常适合 Web 开发。一些云服务（例如 Cloudflare、Digital Ocean 等）将 Go 押注于其云基础架构，证明了这一说法。但是，是什么让 Go 成为构建面向网络的应用程序的一个引人注目的选择？

One reason is the built-in concurrency in the language through goroutines and channels, which makes programs effectively use the resources of the hardware its running on. Ruby and other web-centered languages are usually single-threaded. Going beyond that can be done with the help of threads, but these feel like an add-on rather than a first-class citizen of the language.

原因之一是语言中通过 goroutine 和通道内置的并发性，这使得程序可以有效地使用其运行的硬件资源。 Ruby 和其他以 Web 为中心的语言通常是单线程的。可以在线程的帮助下超越这一点，但这些感觉像是附加组件，而不是该语言的一等公民。

Go also boasts a great standard library that includes many of the basic utilities necessary to build a robust web service, including an HTTP server, a templating library, and JSON utilities, as well as interfacing with databases, such as SQL. Support is excellent for all the latest technologies, ranging from HTTP/2 to databases, such as Postgres, MongoDB and ElasticSearch, and the latest HTTPS encryption standards (including TLS 1.3).

Go 还拥有一个很棒的标准库，其中包括构建强大的 Web 服务所需的许多基本实用程序，包括 HTTP 服务器、模板库和 JSON 实用程序，以及与数据库（如 SQL）的接口。支持所有最新技术，从 HTTP/2 到数据库，如 Postgres、MongoDB 和 ElasticSearch，以及最新的 HTTPS 加密标准（包括 TLS 1.3）。

In terms of performance, Go is far ahead when compared to Ruby and other scripting languages. For example, Hugo, a static site generator built with Go, generates sites an average of [35 times faster](https://forestry.io/blog/hugo-and-jekyll-compared/#performance-1) than Ruby- based Jekyll. Additionally, since all Go programs compile to a single binary, it's easy to deploy to any cloud service of your choice. Applications built with Go can run natively in Google's App Engine and Google Cloud Run or numerous other environments, cloud services, or operating systems thanks to Go’s extreme portability.

在性能方面，与 Ruby 和其他脚本语言相比，Go 遥遥领先。例如，使用 Go 构建的静态站点生成器 Hugo 生成站点的速度平均比 Ruby 快 [35 倍](https://forestry.io/blog/hugo-and-jekyll-compared/#performance-1)-基于 Jekyll。此外，由于所有 Go 程序都编译为单个二进制文件，因此可以轻松部署到您选择的任何云服务。由于 Go 具有极高的可移植性，使用 Go 构建的应用程序可以在 Google 的 App Engine 和 Google Cloud Run 或众多其他环境、云服务或操作系统中本地运行。

All this put together makes Go one of the best languages for the development of a new web service or application. This article will describe some specific aspects of building a web server in Go to give you an idea of how it feels to build with Go.

所有这些加在一起使 Go 成为开发新 Web 服务或应用程序的最佳语言之一。本文将描述在 Go 中构建 Web 服务器的一些特定方面，让您了解使用 Go 构建的感觉。

Note that Go is not exactly a drop-in replacement for monolithic frameworks, such as Rails. It's mostly used to build microservices, and many programmers have had good experiences factoring subsets of Rails apps into several HTTP or JSON services and front-ending them with Rails or other frameworks to get the best of both worlds!

请注意，Go 并不是单一框架（例如 Rails）的直接替代品。它主要用于构建微服务，并且许多程序员在将 Rails 应用程序的子集分解为多个 HTTP 或 JSON 服务并使用 Rails 或其他框架对它们进行前端处理以获得两全其美的方面都有很好的经验！

## A Basic HTTP server

## 一个基本的 HTTP 服务器

HTTP servers are easy to write in Go using the `net/http` package. Unlike other languages with half-baked web servers that are tedious to work with and not ideal for use in anything more than a toy app, Go's `net/http` package is meant for production use and is, indeed, the most popular way to develop a web service in the language, although other solutions do exist.

使用 `net/http` 包可以很容易地在 Go 中编写 HTTP 服务器。与其他带有半生不熟的 Web 服务器的语言不同，这些语言使用起来很乏味，而且不适合用于玩具应用程序以外的任何东西，Go 的 `net/http` 包是为生产使用而设计的，实际上，它是最流行的方式尽管确实存在其他解决方案，但使用该语言开发 Web 服务。

An example of the most basic HTTP server you can build in the language is presented below. It runs on port 8080 and responds with the HTML string `<h1>Welcome to my web server!</h1>` when a request is made to the server root.

您可以用该语言构建的最基本的 HTTP 服务器的示例如下所示。它在端口 8080 上运行，并在向服务器根目录发出请求时以 HTML 字符串“<h1>Welcome to my web server!</h1>”作为响应。

```
package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<h1>Welcome to my web server!</h1>"))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```


If you start the server and visit http://localhost:8080 in your browser, you should see something similar to the image below:

如果您启动服务器并在浏览器中访问 http://localhost:8080，您应该会看到类似于下图的内容：

![Welcome to my web server!](https://www.honeybadger.io/images/blog/posts/go-web-services/1.png?1630027696)

Let's break down the code a bit to fully understand what's going on. The `net/http` package is the foundation of all web servers in Go. It enables the creation of applications capable of making requests to other servers, as well as responding to incoming web requests. 

让我们稍微分解一下代码以完全理解发生了什么。 `net/http` 包是 Go 中所有 Web 服务器的基础。它支持创建能够向其他服务器发出请求以及响应传入 Web 请求的应用程序。

The `main` function begins with a call to the `http.HandleFunc` method, which provides a way to specify how requests to a specific route should be handled. The first argument is the route in question ("/" in this case), while the second is a function of the type [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc).

`main` 函数首先调用 `http.HandleFunc` 方法，该方法提供了一种方法来指定应如何处理对特定路由的请求。第一个参数是有问题的路由（在本例中为“/”），而第二个参数是 [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc) 类型的函数。

Any function with the signature `func(http.ResponseWriter, *http.Request)` can be passed to any other function that expects the `http.HandlerFunc` type. In this case, the function is defined inline, but you can also make it a standalone function, as shown below:

任何带有 `func(http.ResponseWriter, *http.Request)` 签名的函数都可以传递给任何其他需要 `http.HandlerFunc` 类型的函数。在这种情况下，函数是内联定义的，但您也可以使其成为独立函数，如下所示：

```
func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Welcome to my web server!</h1>"))
}

func main() {
    http.HandleFunc("/", indexHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```


The first argument to the `indexHandler` function is a value of the type `http.ResponseWriter`. This is the mechanism used for sending responses to any connected HTTP clients. It's also how response headers are set. The second argument is a pointer to an `http.Request`. It's how data is retrieved from the web request. For example, the details from a form submission can be accessed through the request pointer.

`indexHandler` 函数的第一个参数是 `http.ResponseWriter` 类型的值。这是用于向任何连接的 HTTP 客户端发送响应的机制。这也是设置响应标头的方式。第二个参数是一个指向 `http.Request` 的指针。这就是从 Web 请求中检索数据的方式。例如，可以通过请求指针访问表单提交的详细信息。

Inside our handler, we have a single line of code:

在我们的处理程序中，我们有一行代码：

```
w.Write([]byte("<h1>Welcome to my web server!</h1>"))
```


The `http.ResponseWriter` interface has a `Write` method which accepts a byte slice and writes the data to the connection as part of an HTTP response. Converting a string to a byte slice is as easy as using `[]byte(str)`, and that's how we're able to respond to HTTP requests.

`http.ResponseWriter` 接口有一个 `Write` 方法，它接受一个字节片并将数据作为 HTTP 响应的一部分写入连接。将字符串转换为字节片就像使用 `[]byte(str)` 一样简单，这就是我们能够响应 HTTP 请求的方式。

In the final line of the `main` function, we have the `http.ListenAndServe` method wrapped in a call to `log.Fatal`. `ListenAndServe` specifies the port to listen on as the first argument and an [http.Handler](https://golang.org/pkg/net/http/#Handler) as its second argument. If the handler is nil, `DefaultServeMux` is used.

在 `main` 函数的最后一行，我们将 `http.ListenAndServe` 方法封装在对 `log.Fatal` 的调用中。 `ListenAndServe` 指定要侦听的端口作为第一个参数，将 [http.Handler](https://golang.org/pkg/net/http/#Handler) 指定为第二个参数。如果处理程序为零，则使用 `DefaultServeMux`。

While `DefaultServeMux` is okay for toy programs, you should avoid it in your production code. This is because it is accessible to any package that your application imports, including third-party packages. Therefore, it could potentially be exploited to expose a malicious handler to the web if a package becomes compromised.

虽然`DefaultServeMux` 适用于玩具程序，但您应该在生产代码中避免使用它。这是因为您的应用程序导入的任何包都可以访问它，包括第三方包。因此，如果程序包受到威胁，它可能会被利用以将恶意处理程序暴露给网络。

So, what's the alternative? A locally scoped [http.ServeMux](https://golang.org/pkg/net/http/#ServeMux)! Here's how it works:

那么，有什么替代方案呢？本地范围的 [http.ServeMux](https://golang.org/pkg/net/http/#ServeMux)！这是它的工作原理：

```
package main

import (
    "log"
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Welcome to my web server!</h1>"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```


An HTTP ServeMux is essentially a request router that compares incoming requests against a list of predefined URL paths and executes the associated handler for the path whenever a match is found. When you don't create your own locally scoped ServeMux, the default is used, and it's where methods like `http.HandleFunc` attach their handlers.

HTTP ServeMux 本质上是一个请求路由器，它将传入请求与预定义的 URL 路径列表进行比较，并在找到匹配项时执行该路径的关联处理程序。当您不创建自己的本地作用域 ServeMux 时，将使用默认值，它是像 `http.HandleFunc` 这样的方法附加它们的处理程序的地方。

But, creating your own ServeMux is easy using the `http.NewServeMux` method and registering your handlers on the resulting ServeMux. Also, instead of `nil`, the custom ServeMux should be used as the second argument to `ListenAndServe`, as shown above.

但是，使用`http.NewServeMux` 方法创建您自己的ServeMux 并在结果ServeMux 上注册您的处理程序很容易。此外，如上所示，自定义 ServeMux 应用作 ListenAndServe 的第二个参数，而不是 nil。

The reason `http.ListenAndServe` is wrapped inside a call to `log.Fatal` is it will always return an error if something unexpected happens, and you definitely want to log the error. Otherwise, it will block until the program is terminated.

`http.ListenAndServe` 包含在对 `log.Fatal` 的调用中的原因是，如果发生意外情况，它总是会返回错误，并且您肯定想记录错误。否则，它将阻塞直到程序终止。

## Routing

## 路由

Routing is basically the process of mapping an inbound request to the appropriate HTTP handler. As discussed above, the `http.ServeMux` type provides this functionality in our program by virtue of its `HandleFunc` method. The process of adding a new route to the server is as simple as registering a new handler, as shown below:

路由基本上是将入站请求映射到适当的 HTTP 处理程序的过程。如上所述，`http.ServeMux` 类型凭借其 `HandleFunc` 方法在我们的程序中提供了此功能。向服务器添加新路由的过程就像注册一个新的处理程序一样简单，如下所示：

```
...
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>This is the about page</h1>"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/about", aboutHandler)
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```


If you restart the server and load http://localhost:8080/about in your browser, you will see the expected response on the screen.

如果您重新启动服务器并在浏览器中加载 http://localhost:8080/about，您将在屏幕上看到预期的响应。

Let's consider some other routing features, as well as shortcomings, of the `http.ServeMux` in more detail.

让我们更详细地考虑一下 `http.ServeMux` 的一些其他路由功能以及缺点。

### Pattern matching 

###  模式匹配

The ServeMux handler follows a pretty basic model for routing requests, but it can be very confusing. For each route, you need to be explicit about the registered paths since it does not support wildcards or regular expressions.

ServeMux 处理程序遵循一个非常基本的路由请求模型，但它可能非常令人困惑。对于每个路由，您需要明确注册路径，因为它不支持通配符或正则表达式。

Two kinds of path patterns are supported by Go's ServeMux: fixed paths and subtree paths. The former does not end with a trailing slash, while the latter does.

Go 的 ServeMux 支持两种路径模式：固定路径和子树路径。前者不以斜杠结尾，而后者以斜杠结尾。

We've already seen both types of paths in our demo application. The root pattern ("/")is a subtree path, while "/about" is a fixed path. The difference is that the handler for a fixed pattern is only expected when the request URL path is an *exact* match for the pattern. Subtree path handlers, however, will execute whenever the start of the request URL matches the subtree pattern.

我们已经在演示应用程序中看到了这两种类型的路径。根模式（“/”）是子树路径，而“/about”是固定路径。不同之处在于固定模式的处理程序仅在请求 URL 路径与模式*完全*匹配时才需要。但是，只要请求 URL 的开头与子树模式匹配，子树路径处理程序就会执行。

Here's an example that will make it clearer for you to understand. Let's say we have the following registered paths:

这是一个示例，可以让您更清楚地理解。假设我们有以下注册路径：

```
mux.HandleFunc("/", indexHandler)
mux.HandleFunc("/user", userHandler)
mux.HandleFunc("/user/create", createUserHandler)
```


Here's how ServeMux will route the following incoming requests:

以下是 ServeMux 将如何路由以下传入请求：

```
/ => indexHandler
/user => userHandler
/user/ (with trailing slash) => indexHandler
/user/create => createUserHandler
/user/foo => indexHandler
/foo => indexHandler
/bar => indexHandler
```


A fixed path is only matched when the request URL is an exact match and does not include a trailing slash. Otherwise, the longest matched subtree pattern will take precedence over the shorter ones. If we register the `/user/` subtree pattern, requests to `/user/foo` will match `/user/` and not `/` because it is longer.

固定路径仅在请求 URL 完全匹配且不包含尾部斜杠时匹配。否则，最长匹配的子树模式将优先于较短的子树模式。如果我们注册 `/user/` 子树模式，对 `/user/foo` 的请求将匹配 `/user/` 而不是 `/`，因为它更长。

Here are some other details about how Go's ServeMux routes incoming requests:

以下是有关 Go 的 ServeMux 如何路由传入请求的其他一些细节：

- Route patterns can be registered in any order. It does not affect the behavior of the router.
- Incoming requests to a subtree path that do not have a trailing slash will be 301 redirected to subtree path with the slash added. So, if `/image/` is registered, and a request is made to `/image`, it will be redirected to `/image/` as long as `/image` itself is not registered.
- It's possible to specify the hostname in the route pattern. We could register a path, such as `auth.example.com`, and requests to that URL will be directed to the registered handler.

- 可以按任何顺序注册路线模式。它不会影响路由器的行为。
- 对没有尾部斜杠的子树路径的传入请求将被 301 重定向到添加了斜杠的子树路径。因此，如果`/image/` 被注册，并且向`/image` 发出请求，只要`/image` 本身没有被注册，它就会被重定向到`/image/`。
- 可以在路由模式中指定主机名。我们可以注册一个路径，例如 `auth.example.com`，并且对该 URL 的请求将被定向到注册的处理程序。

### 404 Not Found

### 404 未找到

As mentioned above, patterns that end with a trailing slash act like a catch-all for URL paths that match the start of the pattern. This includes the root route, which matches any URL not handled by a more specific route. If you want the root pattern to stop acting as a catch-all path, you can check whether the current request URL exactly matches `/`. If it doesn't, use the `http.NotFound()` method to send a 404 not found response.

如上所述，以斜杠结尾的模式对于匹配模式开头的 URL 路径来说就像一个包罗万象的东西。这包括根路由，它匹配任何未由更具体路由处理的 URL。如果您希望根模式停止充当全能路径，您可以检查当前请求 URL 是否与 `/` 完全匹配。如果没有，请使用 `http.NotFound()` 方法发送 404 not found 响应。

```
func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Write([]byte("<h1>Welcome to my web server!</h1>"))
}
```


Now, requests to unregistered paths will yield a 404 error, as shown in the image below. You'll have to repeat this check for each subtree path you register in your server.

现在，对未注册路径的请求将产生 404 错误，如下图所示。您必须对在服务器中注册的每个子树路径重复此检查。

![404 error page](https://www.honeybadger.io/images/blog/posts/go-web-services/3.png?1630027696)

### Redirects

### 重定向

You can redirect from one URL to another using the `http.Redirect()` method.

您可以使用 `http.Redirect()` 方法从一个 URL 重定向到另一个 URL。

```
func redirect(url string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, url, 301)
    }
}

func main() {
    mux.HandleFunc("/example", redirect("http://example.com"))
}
```


### Serving static assets

### 服务静态资产

To serve static files from a web server, the `http.FileServer()` method can be utilized.

要从 Web 服务器提供静态文件，可以使用 `http.FileServer()` 方法。

```
func main() {
    staticHandler := http.FileServer(http.Dir("./assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))
}
```


The `FileServer()` function returns an `http.Handler` that responds to all HTTP requests with the contents of the provided filesystem. In the above example, the filesystem is given as the `assets` directory relative to the application. 

`FileServer()` 函数返回一个 `http.Handler`，它使用提供的文件系统的内容响应所有 HTTP 请求。在上面的例子中，文件系统是作为相对于应用程序的`assets` 目录给出的。

Next, the `Handle` method is used to register the FileServer Handler for all request URLs that begin with `/assets/`. The trailing slash is significant for reasons we already discussed above. The `http.StripPrefix()` method is used to strip out the `/assets/` prefix from the request path before searching the filesystem for the requested file.

接下来，`Handle` 方法用于为所有以 `/assets/` 开头的请求 URL 注册 FileServer Handler。由于我们上面已经讨论过的原因，尾部斜杠很重要。 `http.StripPrefix()` 方法用于在搜索文件系统以查找请求的文件之前从请求路径中去除 `/assets/` 前缀。

From now on, any requests for static files in the `assets` directory, such as `/assets/css/styles.css` or `/assets/images/sample.jpg`, will be handled appropriately.

从现在开始，任何对 `assets` 目录中的静态文件的请求，例如 `/assets/css/styles.css` 或 `/assets/images/sample.jpg`，都会得到适当的处理。

### HTTP request methods

### HTTP 请求方法

Thus far, we haven't given much consideration to the request methods allowed for each route, so let's do that here. Go's ServeMux does not have any special way to specify the allowed methods for a route, so you have to check the request method yourself in the HTTP handler.

到目前为止，我们还没有过多考虑每个路由允许的请求方法，所以让我们在这里做一下。 Go 的 ServeMux 没有任何特殊的方式来指定允许的路由方法，因此您必须自己在 HTTP 处理程序中检查请求方法。

Let's make sure that the `indexHandler` function returns an error if a non-GET request is made to the root route.

如果对根路由发出非 GET 请求，让我们确保 `indexHandler` 函数返回错误。

```
func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    if r.Method == "GET" {
        w.Write([]byte("<h1>Welcome to my web server!</h1>"))
    } else {
        http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
    }
}
```


If you want the route to handle more than one HTTP method, using a `switch` statement may be preferable.

如果您希望路由处理多个 HTTP 方法，则使用 `switch` 语句可能更可取。

### Reading query parameters

### 读取查询参数

It's often necessary to read the query parameters sent as part of an HTTP request. For example, given the following request url:

通常需要读取作为 HTTP 请求的一部分发送的查询参数。例如，给定以下请求 url：

```
/user?id=123
```


Here's how to access the value of the `id` query parameter in Go:

以下是如何在 Go 中访问 `id` 查询参数的值：

```
func userHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "The id query parameter is missing", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "<h1>The user id is: %s</h1>", id)
}

func main() {
    mux.HandleFunc("/user", userHandler)
}
```


![Reading query parameters](https://www.honeybadger.io/images/blog/posts/go-web-services/2.png?1630027696)

If you want to iterate over all the query parameters sent as part of a request, you can use the following code snippet:

如果要遍历作为请求的一部分发送的所有查询参数，可以使用以下代码片段：

```
values := r.URL.Query()
for k, v := range values {
    fmt.Println(k, " => ", v)
}
```


## Why you should care about ServeMux

## 为什么你应该关心 ServeMux

If you're coming to Go with a background in using frameworks like Ruby on Rails, the capabilities of the built-in router in Go may seem incredibly simple or even underwhelming.

如果您在使用 Ruby on Rails 等框架的背景下开始使用 Go，那么 Go 中内置路由器的功能可能看起来非常简单甚至令人印象深刻。

A quick Google search will locate some great third-party options that you can plug into your application, but I recommend that you learn the limitations of ServeMux first before opting for one of them. Besides, many projects use the built-in router, so you still need to know how it works to be able to contribute effectively.

快速谷歌搜索将找到一些可以插入应用程序的优秀第三方选项，但我建议您在选择其中一个之前先了解 ServeMux 的限制。此外，许多项目使用内置路由器，因此您仍然需要知道它是如何工作的才能有效地做出贡献。

Go developers tend to rely on the standard library in most cases, but there are times when it is necessary to go beyond what it provides. In a future article, we'll examine some of the popular third-party routers to see what else is out there.

在大多数情况下，Go 开发人员倾向于依赖标准库，但有时需要超越它提供的内容。在以后的文章中，我们将研究一些流行的第三方路由器，看看还有什么。

## Wrap up

##  包起来

In this article, we discussed why Go is great for web development and proceeded to build an HTTP server in the language while examining the different features of the built-in router that you need to know about.

在本文中，我们讨论了为什么 Go 非常适合 Web 开发，并继续使用该语言构建 HTTP 服务器，同时检查您需要了解的内置路由器的不同功能。

If you have any questions or opinions on what we've covered here, I'd love to hear about it on [Twitter](https://twitter.com/ayisaiah). In the next article, we'll cover some other aspects of building web applications with Go, such as middleware and JSON, as well as templating and working with databases.

如果您对我们在此处介绍的内容有任何问题或意见，我很乐意在 [Twitter](https://twitter.com/ayisaiah) 上听到有关内容。在下一篇文章中，我们将介绍使用 Go 构建 Web 应用程序的其他一些方面，例如中间件和 JSON，以及模板化和使用数据库。

Thanks for reading, and happy coding! 

感谢阅读，祝您编码愉快！

