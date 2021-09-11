## Writing Advanced Web Applications with Go

## 使用 Go 编写高级 Web 应用程序

    Last edited:      *Apr 19, 2018*

最后编辑：*2018 年 4 月 19 日*

Web development in many programming environments often requires subscribing to some full framework ethos. With [Ruby](https://www.ruby-lang.org/), it's usually [Rails](http://rubyonrails.org/) but could be [Sinatra](http://www.sinatrarb.com/) or something else. With [Python](https://www.python.org/), it's often [Django](https://www.djangoproject.com/) or [Flask](http://flask.pocoo.org/) . With [Go](https://golang.org/), it’s…

许多编程环境中的 Web 开发通常需要订阅一些完整的框架精神。对于 [Ruby](https://www.ruby-lang.org/)，通常是[Rails](http://rubyonrails.org/) 但也可能是 [Sinatra](http://www.sinatrarb.com/) 或者是其他东西。使用 [Python](https://www.python.org/)，通常是[Django](https://www.djangoproject.com/) 或 [Flask](http://flask.pocoo.org/) .使用 [Go](https://golang.org/)，它是……

If you spend some time in Go communities like the [Go mailing list](https://groups.google.com/d/forum/golang-nuts) or the [Go subreddit](https://www.reddit.com/r/golang/), you'll find Go newcomers frequently wondering what web framework is best to use. [There](https://revel.github.io/)[are](https://gin-gonic.github.io/gin/) [quite](http://iris-go.com/) [ a](https://beego.me/)[few](https://go-macaron.com/) [Go](https://github.com/go-martini/martini) [frameworks](https://github.com/gocraft/web) ([and](https://github.com/urfave/negroni)[then](https://godoc.org/goji.io) [some](https://echo.labstack.com/)), so which one is best seems like a reasonable question. Without fail, though, the strong recommendation of the Go community is to [avoid web frameworks entirely](https://medium.com/code-zen/why-i-don-t-use-go-web-frameworks-1087e1facfa4) and just stick with the standard library as long as possible. Here's [an example from the Go mailing list](https://groups.google.com/forum/#!topic/golang-nuts/R_lqsTTBh6I) and here's [one from the subreddit](https://www.reddit.com/r/golang/comments/1yh6gm/new_to_go_trying_to_select_web_framework/).

如果您在 [Go 邮件列表](https://groups.google.com/d/forum/golang-nuts) 或 [Go subreddit](https://www.reddit.com) 等 Go 社区中度过了一段时间/r/golang/)，你会发现 Go 新手经常想知道最好使用什么 Web 框架。 [那里](https://revel.github.io/)[are](https://gin-gonic.github.io/gin/) [相当](http://iris-go.com/) [ a](https://beego.me/)[很少](https://go-macaron.com/) [Go](https://github.com/go-martini/martini) [框架](https://github.com/gocraft/web) ([and](https://github.com/urfave/negroni)[then](https://godoc.org/goji.io) [some](https://echo.labstack.com/))，所以哪个最好似乎是一个合理的问题。尽管如此，Go 社区的强烈建议是 [完全避免使用 Web 框架](https://medium.com/code-zen/why-i-don-t-use-go-web-frameworks-1087e1facfa4) 并尽可能长时间地使用标准库。这是[Go 邮件列表中的一个示例](https://groups.google.com/forum/#!topic/golang-nuts/R_lqsTTBh6I) 和[来自 subreddit 的一个示例](https://www.reddit。com/r/golang/comments/1yh6gm/new_to_go_trying_to_select_web_framework/)。

It’s not bad advice! The Go standard library is very rich and flexible, much more so than many other languages, and designing a web application in Go with just the standard library is definitely a good choice.

这不是坏建议！ Go 标准库非常丰富和灵活，比许多其他语言更丰富，仅使用标准库在 Go 中设计 Web 应用程序绝对是一个不错的选择。

Even when these Go frameworks call themselves minimalistic, they can't seem to help themselves avoid using a different request handler interface than the default standard library [http.Handler](https://golang.org/pkg/net/http/#Handler), and I think this is the biggest source of angst about why frameworks should be avoided. If everyone standardizes on [http.Handler](https://golang.org/pkg/net/http/#Handler), then dang, all sorts of things would be interoperable!

即使这些 Go 框架称自己为简约，它们似乎也无法避免使用与默认标准库 [http.Handler](https://golang.org/pkg/net/http/#Handler)，我认为这是关于为什么应该避免框架的最大焦虑来源。如果大家都标准化在[http.Handler](https://golang.org/pkg/net/http/#Handler)上，那该死，各种东西都可以互通了！

Before Go 1.7, it made some sense to give in and use a different interface for handling HTTP requests. But now that [http.Request](https://golang.org/pkg/net/http/#Request) has the [Context](https://golang.org/pkg/net/http/#Request.Context) and [WithContext](https://golang.org/pkg/net/http/#Request.WithContext) methods, there truly isn't a good reason any longer.

在 Go 1.7 之前，放弃并使用不同的接口来处理 HTTP 请求是有意义的。但是现在 [http.Request](https://golang.org/pkg/net/http/#Request) 有了 [Context](https://golang.org/pkg/net/http/#Request.Context) 和 [WithContext](https://golang.org/pkg/net/http/#Request.WithContext) 方法，真的没有什么好的理由了。

I’ve done a fair share of web development in Go and I’m here to share with you both some standard library development patterns I’ve learned and some code I’ve found myself frequently needing. The code I’m sharing is not for use instead of the standard library, but to augment it.

我在 Go 中做了相当多的 Web 开发，我在这里与你分享一些我学到的标准库开发模式和一些我发现自己经常需要的代码。我分享的代码不是用来代替标准库的，而是用来扩充它的。

Overall, if this blog post feels like it’s predominantly plugging various little standalone libraries from my [Webhelp non-framework](https://godoc.org/gopkg.in/webhelp.v1), that’s because it is. It’s okay, they’re little standalone libraries. Only use the ones you want!

总的来说，如果这篇博文主要是从我的 [Webhelp 非框架](https://godoc.org/gopkg.in/webhelp.v1) 中插入各种小型独立库，那是因为它确实如此。没关系，它们是小型的独立库。只使用你想要的！

If you’re new to Go web development, I suggest reading the Go documentation’s [Writing Web Applications](https://golang.org/doc/articles/wiki/) article first.

如果您是 Go Web 开发的新手，我建议您先阅读 Go 文档的 [Writing Web Applications](https://golang.org/doc/articles/wiki/) 文章。

## Middleware

## 中间件

A frequent design pattern for server-side web development is the concept of *middleware*, where some portion of the request handler wraps some other portion of the request handler and does some preprocessing or routing or something. This is a big component of how [Express](https://expressjs.com/) is organized on [Node](https://nodejs.org/en/), and how Express middleware and [Negroni](https://github.com/urfave/negroni) middleware works is almost line-for-line identical in design.

服务器端 Web 开发的一个常见设计模式是 * 中间件 * 的概念，其中请求处理程序的某些部分包装请求处理程序的其他部分并执行一些预处理或路由或其他操作。这是 [Express](https://expressjs.com/) 在 [Node](https://nodejs.org/en/) 上的组织方式以及 Express 中间件和 [Negroni](https://github.com/urfave/negroni) 中间件在设计上几乎是逐行相同的。

Good use cases for middleware are things such as:

中间件的良好用例是这样的：

- making sure a user is logged in, redirecting if not,
- making sure the request came over HTTPS,
- making sure a session is set up and loaded from a session database,
- making sure we logged information before and after the request was handled,
- making sure the request was routed to the right handler,
- and so on. 

- 确保用户已登录，如果未登录，则重定向，
- 确保请求通过 HTTPS，
- 确保会话已建立并从会话数据库加载，
- 确保我们在处理请求之前和之后记录信息，
- 确保请求被路由到正确的处理程序，
- 等等。

Composing your web app as essentially a chain of middleware handlers is a very powerful and flexible approach. It allows you to avoid a lot of [cross-cutting concerns](https://en.wikipedia.org/wiki/Cross-cutting_concern) and have your code factored in very elegant and easy-to-maintain ways. By wrapping a set of handlers with middleware that ensures a user is logged in prior to actually attempting to handle the request, the individual handlers no longer need mistake-prone copy-and-pasted code to ensure the same thing.

将您的 Web 应用程序本质上构成一个中间件处理程序链是一种非常强大且灵活的方法。它可以让您避免很多 [横切关注点](https://en.wikipedia.org/wiki/Cross-cutting_concern) 并以非常优雅且易于维护的方式将您的代码分解。通过用中间件包装一组处理程序，确保用户在实际尝试处理请求之前登录，各个处理程序不再需要容易出错的复制和粘贴代码来确保相同的事情。

So, middleware is good. However, if Negroni or other frameworks are any indication, you’d think the standard library’s `http.Handler` isn’t up to the challenge. Negroni adds its own `negroni.Handler` just for the sake of making middleware easier. There’s no reason for this.

所以，中间件是好的。然而，如果 Negroni 或其他框架有任何迹象，你会认为标准库的 `http.Handler` 无法应对挑战。 Negroni 添加了自己的 `negroni.Handler` 只是为了让中间件更容易。这没有任何理由。

Here is a full middleware implementation for ensuring a user is logged in, assuming a `GetUser(*http.Request)` function but otherwise just using the standard library:



这是一个完整的中间件实现，用于确保用户登录，假设有一个 GetUser(*http.Request) 函数，否则只使用标准库：



```
func RequireUser(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    user, err := GetUser(req)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    if user == nil {
      http.Error(w, "unauthorized", http.StatusUnauthorized)
      return
    }
    h.ServeHTTP(w, req)
  })
}
```




Here’s how it’s used (just wrap another handler!):



这是它的使用方法（只需包装另一个处理程序！）：



```
func main() {
  http.ListenAndServe(":8080", RequireUser(http.HandlerFunc(myHandler)))
}
```




Express, Negroni, and other frameworks expect this kind of signature for a middleware-supporting handler:



Express、Negroni 和其他框架期待这种支持中间件的处理程序的签名：



```
type Handler interface {
  // don't do this!
  ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc)
}
```




There’s really no reason for adding the `next` argument - it reduces cross-library compatibility. So I say, don’t use `negroni.Handler` (or similar). Just use `http.Handler`!

真的没有理由添加 `next` 参数 - 它降低了跨库兼容性。所以我说，不要使用 `negroni.Handler`（或类似的）。只需使用`http.Handler`！

## Composability

## 可组合性

Hopefully I’ve sold you on middleware as a good design philosophy.

希望我已经把中间件作为一种好的设计理念卖给了你。

Probably the most commonly-used type of middleware is request routing, or muxing (seems like we should call this demuxing but what do I know). Some frameworks are almost solely focused on request routing. [gorilla/mux](https://github.com/gorilla/mux) seems more popular than any other part of the [Gorilla](https://github.com/gorilla/) library. I think the reason for this is that even though the Go standard library is completely full featured and has a good [ServeMux](https://golang.org/pkg/net/http/#ServeMux) implementation, it doesn't make the right thing the default.

可能最常用的中间件类型是请求路由或多路复用（似乎我们应该称之为分离，但我知道什么）。一些框架几乎只专注于请求路由。 [gorilla/mux](https://github.com/gorilla/mux) 似乎比 [Gorilla](https://github.com/gorilla/) 库的任何其他部分都更受欢迎。我认为这样做的原因是，即使 Go 标准库功能齐全并且有一个很好的 [ServeMux](https://golang.org/pkg/net/http/#ServeMux) 实现，但它并没有使正确的事情是默认的。

So! Let’s talk about request routing and consider the following problem. You, web developer extraordinaire, want to serve some HTML from your web server at `/hello/` but also want to serve some static assets from `/static/`. Let’s take a quick stab.



所以！让我们谈谈请求路由并考虑以下问题。您，杰出的 Web 开发人员，希望从您的 Web 服务器在 `/hello/` 提供一些 HTML，但也希望从 `/static/` 提供一些静态资产。让我们快速尝试一下。



```
package main

import (
  "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("hello, world!"))
}

func main() {
  mux := http.NewServeMux()
  mux.Handle("/hello/", http.HandlerFunc(hello))
  mux.Handle("/static/", http.FileServer(http.Dir("./static-assets")))
  http.ListenAndServe(":8080", mux)
}
```




If you visit `http://localhost:8080/hello/`, you’ll be rewarded with a friendly “hello, world!” message.

如果你访问`http://localhost:8080/hello/`，你会得到一个友好的“hello, world!”的奖励。信息。

If you visit `http://localhost:8080/static/` on the other hand (assuming you have a folder of static assets in `./static-assets`), you’ll be surprised and frustrated. This code tries to find the source content for the request `/static/my-file` at `./static-assets/static/my-file`! There’s an extra `/static` in there!

另一方面，如果你访问 `http://localhost:8080/static/`（假设你在 `./static-assets` 中有一个静态资产文件夹），你会感到惊讶和沮丧。此代码尝试在`./static-assets/static/my-file` 中查找请求`/static/my-file` 的源内容！那里有一个额外的 `/static`！

Okay, so this is why `http.StripPrefix` exists. Let’s fix it.



好的，这就是`http.StripPrefix` 存在的原因。让我们修复它。



```
   mux.Handle("/static/", http.StripPrefix("/static",
    http.FileServer(http.Dir("./static-assets"))))
```




`mux.Handle` combined with `http.StripPrefix` is such a common pattern that I think it should be the default. Whenever a request router processes a certain amount of URL elements, it should strip them off the request so the wrapped `http.Handler` doesn’t need to know its absolute URL and only needs to be concerned with its relative one. 

`mux.Handle` 结合 `http.StripPrefix` 是一种常见的模式，我认为它应该是默认的。每当请求路由器处理一定数量的 URL 元素时，它应该将它们从请求中剥离，这样包装的 `http.Handler` 不需要知道它的绝对 URL，只需要关心它的相对 URL。

In [Russ Cox](https://swtch.com/~rsc/)'s recent [TiddlyWeb backend](https://github.com/rsc/tiddly), I would argue that every time `strings.TrimPrefix` is needed to remove the full URL from the handler's incoming path arguments, it is an unnecessary cross-cutting concern, unfortunately imposed by `http.ServeMux`. (An example is [line 201 in tiddly.go](https://github.com/rsc/tiddly/blob/8f9145ac183e374eb95d90a73be4d5f38534ec47/tiddly.go#L201).)

在 [Russ Cox](https://swtch.com/~rsc/) 最近的 [TiddlyWeb backend](https://github.com/rsc/tiddly) 中，我认为每次 `strings.TrimPrefix`需要从处理程序的传入路径参数中删除完整的 URL，这是一个不必要的横切关注点，不幸的是由 `http.ServeMux` 强加的。 （一个例子是 [tiddly.go 中的第 201 行](https://github.com/rsc/tiddly/blob/8f9145ac183e374eb95d90a73be4d5f38534ec47/tiddly.go#L201)。)

I’d much rather have the default `mux` behavior work more like a directory of registered elements that by default strips off the ancestor directory before handing the request to the next middleware handler. It’s much more composable. To this end, I’ve written a simple muxer that works in this fashion called [whmux.Dir](https://godoc.org/gopkg.in/webhelp.v1/whmux#Dir). It is essentially `http.ServeMux` and `http.StripPrefix` combined. Here’s the previous example reworked to use it:



我更希望默认的“mux”行为更像是一个注册元素的目录，默认情况下，在将请求交给下一个中间件处理程序之前，它会剥离祖先目录。它的可组合性要好得多。为此，我编写了一个以这种方式工作的简单复用器，称为 [whmux.Dir](https://godoc.org/gopkg.in/webhelp.v1/whmux#Dir)。它本质上是`http.ServeMux` 和`http.StripPrefix` 的组合。这是前面的示例重新设计以使用它：



```
package main

import (
  "net/http"

  "gopkg.in/webhelp.v1/whmux"
)

func hello(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("hello, world!"))
}

func main() {
  mux := whmux.Dir{
    "hello":  http.HandlerFunc(hello),
    "static": http.FileServer(http.Dir("./static-assets")),
  }
  http.ListenAndServe(":8080", mux)
}
```




There are other useful mux implementations inside the [whmux](https://godoc.org/gopkg.in/webhelp.v1/whmux) package that demultiplex on various aspects of the request path, request method, request host, or pull arguments out of the request and place them into the context, such as a [whmux.IntArg](https://godoc.org/gopkg.in/webhelp.v1/whmux#IntArg) or [whmux.StringArg](https://godoc.org/gopkg.in/webhelp.v1/whmux#StringArg). This brings us to [contexts](https://golang.org/pkg/context/).

[whmux](https://godoc.org/gopkg.in/webhelp.v1/whmux) 包中还有其他有用的多路复用实现，可以在请求路径、请求方法、请求主机或拉参数的各个方面进行多路复用脱离请求并将它们放入上下文中，例如 [whmux.IntArg](https://godoc.org/gopkg.in/webhelp.v1/whmux#IntArg) 或 [whmux.StringArg](https://godoc.org/gopkg.in/webhelp.v1/whmux#IntArg)/godoc.org/gopkg.in/webhelp.v1/whmux#StringArg)。这将我们带到 [上下文](https://golang.org/pkg/context/)。

## Contexts

## 上下文

Request contexts are a recent addition to the Go 1.7 standard library, but the idea of [contexts has been around since mid-2014](https://blog.golang.org/context). As of Go 1.7, they were added to the standard library ([“context”](https://golang.org/pkg/context/)), but are available for older Go releases in the original location ([“golang. org/x/net/context”](https://godoc.org/golang.org/x/net/context)).

请求上下文是 Go 1.7 标准库的最新补充，但 [上下文自 2014 年年中就已存在](https://blog.golang.org/context) 的想法。从 Go 1.7 开始，它们被添加到标准库中 ([“context”](https://golang.org/pkg/context/))，但在原始位置（[“golang.org/x/net/context”](https://godoc.org/golang.org/x/net/context))。

First, here’s the definition of the `context.Context` type that `(*http.Request).Context()` returns:



首先，这是 `(*http.Request).Context()` 返回的 `context.Context` 类型的定义：



```
type Context interface {
  Done() <-chan struct{}
  Err() error
  Deadline() (deadline time.Time, ok bool)

  Value(key interface{}) interface{}
}
```




Talking about `Done()`, `Err()`, and `Deadline()` are enough for an entirely different blog post, so I'm going to ignore them at least for now and focus on `Value(interface{} )`.

对于完全不同的博客文章来说，谈论 `Done()`、`Err()` 和 `Deadline()` 就足够了，所以我至少现在将忽略它们并专注于 `Value(interface{} )`。

As a motivating problem, let’s say that the `GetUser(*http.Request)` method we assumed earlier is expensive, and we only want to call it once per request. We certainly don’t want to call it once to check that a user is logged in, and then again when we actually need the `*User` value. With `(*http.Request).WithContext` and `context.WithValue`, we can pass the `*User` down to the next middleware precomputed!

作为一个激励问题，假设我们之前假设的 `GetUser(*http.Request)` 方法很昂贵，并且我们只想每个请求调用它一次。我们当然不想调用它一次来检查用户是否登录，然后在我们真正需要 `*User` 值时再次调用它。使用 `(*http.Request).WithContext` 和 `context.WithValue`，我们可以将 `*User` 传递给下一个预先计算好的中间件！

Here’s the new middleware:



这是新的中间件：



```
type userKey int

func RequireUser(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    user, err := GetUser(req)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    if user == nil {
      http.Error(w, "unauthorized", http.StatusUnauthorized)
      return
    }
    ctx := r.Context()
    ctx = context.WithValue(ctx, userKey(0), user)
    h.ServeHTTP(w, req.WithContext(ctx))
  })
}
```




Now, handlers that are protected by this `RequireUser` handler can load the previously computed `*User` value like this:



现在，受此 `RequireUser` 处理程序保护的处理程序可以像这样加载先前计算的 `*User` 值：



```
if user, ok := req.Context().Value(userKey(0)).(*User);ok {
  // there's a valid user!
}
```




Contexts allow us to pass optional values to handlers down the chain in a way that is relatively type-safe and flexible. None of the above context logic requires anything outside of the standard library.

上下文允许我们以相对类型安全和灵活的方式将可选值传递给链中的处理程序。上述上下文逻辑都不需要标准库之外的任何东西。

### Aside about context keys

### 除了上下文键

There was a curious piece of code in the above example. At the top, we defined a `type userKey int`, and then always used it as `userKey(0)`. 

上面的例子中有一段奇怪的代码。在顶部，我们定义了一个 `type userKey int`，然后始终将其用作 `userKey(0)`。

One of the possible problems with contexts is the `Value()` interface lends itself to a global namespace where you can stomp on other context users and use conflicting key names. Above, we used `type userKey` because it’s an unexported type in your package. It will never compare equal (without a cast) to any other type, including `int`, in Go. This gives us a way to namespace keys to your package, even though the `Value()` method is still a sort of global namespace.

上下文可能存在的问题之一是 `Value()` 接口适用于全局命名空间，您可以在其中踩踏其他上下文用户并使用冲突的键名。上面，我们使用了 `type userKey`，因为它是你的包中未导出的类型。它永远不会与 Go 中的任何其他类型（包括 `int`）比较相等（没有强制转换）。这为我们提供了一种命名包键的方法，即使 `Value()` 方法仍然是一种全局命名空间。

Because the need for this is so common, the `webhelp` package defines a [GenSym()](https://godoc.org/gopkg.in/webhelp.v1#GenSym) helper that will create a brand new, never- before-seen, unique value for use as a context key.

因为这种需求非常普遍，`webhelp` 包定义了一个 [GenSym()](https://godoc.org/gopkg.in/webhelp.v1#GenSym) 帮助器，它将创建一个全新的，从不-前所见，用作上下文键的唯一值。

If we used [GenSym()](https://godoc.org/gopkg.in/webhelp.v1#GenSym), then `type userKey int` would become `var userKey = webhelp.GenSym()` and `userKey( 0)` would simply become `userKey`.

如果我们使用 [GenSym()](https://godoc.org/gopkg.in/webhelp.v1#GenSym)，那么 `type userKey int` 将变成 `var userKey = webhelp.GenSym()` 和 `userKey( 0)` 将简单地变成 `userKey`。

### Back to whmux.StringArg

### 回到 whmux.StringArg

Armed with this new context behavior, we can now present a `whmux.StringArg` example:



有了这个新的上下文行为，我们现在可以展示一个 `whmux.StringArg` 示例：



```
package main

import (
  "fmt"
  "net/http"

  "gopkg.in/webhelp.v1/whmux"
)

var (
  pageName = whmux.NewStringArg()
)

func page(w http.ResponseWriter, req *http.Request) {
  name := pageName.Get(req.Context())

  fmt.Fprintf(w, "Welcome to %s", name)
}

func main() {
  // pageName.Shift pulls the next /-delimited string out of the request's
  // URL.Path and puts it into the context instead.
  pageHandler := pageName.Shift(http.HandlerFunc(page))

  http.ListenAndServe(":8080", whmux.Dir{
    "wiki": pageHandler,
  })
}
```




## Pre-Go-1.7 support

## Pre-Go-1.7 支持

Contexts let you do some pretty cool things. But let’s say you’re stuck with something before Go 1.7 (for instance, App Engine is currently Go 1.6).

上下文让你做一些很酷的事情。但是，假设您在 Go 1.7 之前遇到了一些问题（例如，App Engine 目前是 Go 1.6）。

That’s okay! I’ve backported all of the neat new context features to Go 1.6 and earlier in a forwards compatible way!

没关系！我已经以向前兼容的方式将所有简洁的新上下文功能向后移植到 Go 1.6 及更早版本！

With the [whcompat](https://godoc.org/gopkg.in/webhelp.v1/whcompat) package, `req.Context()` becomes `whcompat.Context(req)`, and `req.WithContext(ctx )` becomes `whcompat.WithContext(req, ctx)`. The `whcompat` versions work with all releases of Go. Yay!

使用 [whcompat](https://godoc.org/gopkg.in/webhelp.v1/whcompat)包，`req.Context()` 变为 `whcompat.Context(req)` 和 `req.WithContext(ctx )` 变为 `whcompat.WithContext(req, ctx)`。 `whcompat` 版本适用于所有版本的 Go。好极了！

There’s a bit of unpleasantness behind the scenes to make this happen. Specifically, for pre-1.7 builds, a global map indexed by `req.URL` is kept, and a finalizer is installed on `req` to clean up. So don’t change what `req.URL` points to and this will work fine. In practice it’s not a problem.

在幕后有一些不愉快的事情发生。具体来说，对于 1.7 之前的构建，保留了由 `req.URL` 索引的全局映射，并在 `req` 上安装了终结器以进行清理。所以不要改变 `req.URL` 指向的内容，这会正常工作。在实践中这不是问题。

`whcompat` adds additional backwards-compatibility helpers. In Go 1.7 and on, the context’s `Done()` channel is closed (and `Err()` is set), whenever the request is done processing. If you want this behavior in Go 1.6 and earlier, just use the [whcompat.DoneNotify](https://godoc.org/gopkg.in/webhelp.v1/whcompat#DoneNotify) middleware.

`whcompat` 添加了额外的向后兼容性助手。在 Go 1.7 及更高版本中，只要完成请求处理，上下文的“Done()”通道就会关闭（并且设置“Err()”）。如果您希望在 Go 1.6 及更早版本中使用此行为，只需使用 [whcompat.DoneNotify](https://godoc.org/gopkg.in/webhelp.v1/whcompat#DoneNotify) 中间件。

In Go 1.8 and on, the context’s `Done()` channel is closed when the client goes away, even if the request hasn’t completed. If you want this behavior in Go 1.7 and earlier, just use the [whcompat.CloseNotify](https://godoc.org/gopkg.in/webhelp.v1/whcompat#CloseNotify) middleware, though beware that it costs an extra goroutine .

在 Go 1.8 及更高版本中，即使请求尚未完成，当客户端离开时，上下文的“Done()”通道也会关闭。如果您希望在 Go 1.7 及更早版本中使用此行为，只需使用 [whcompat.CloseNotify](https://godoc.org/gopkg.in/webhelp.v1/whcompat#CloseNotify) 中间件，但要注意它会花费额外的 goroutine .

## Error handling

## 错误处理

How you handle errors can be another cross-cutting concern, but with good application of context and middleware, it too can be beautifully cleaned up so that the responsibilities lie in the correct place.

如何处理错误可能是另一个横切关注点，但是通过上下文和中间件的良好应用，它也可以被很好地清理，以便将责任放在正确的位置。

Problem statement: your `RequireUser` middleware needs to handle an authentication error differently between your HTML endpoints and your JSON API endpoints. You want to use `RequireUser` for both types of endpoints, but with your HTML endpoints you want to return a user-friendly error page, and with your JSON API endpoints you want to return an appropriate JSON error state.

问题陈述：您的“RequireUser”中间件需要以不同方式处理 HTML 端点和 JSON API 端点之间的身份验证错误。您希望对这两种类型的端点使用 `RequireUser`，但是对于您的 HTML 端点，您希望返回一个用户友好的错误页面，而对于您的 JSON API 端点，您希望返回适当的 JSON 错误状态。

In my opinion, the right thing to do is to have contextual error handlers, and luckily, we have a context for contextual information!

在我看来，正确的做法是拥有上下文错误处理程序，幸运的是，我们有上下文信息的上下文！

First, we need an error handler interface.



首先，我们需要一个错误处理程序接口。



```
type ErrHandler interface {
  HandleError(w http.ResponseWriter, req *http.Request, err error)
}
```




Next, let’s make a middleware that registers the error handler in the context:



接下来，让我们制作一个在上下文中注册错误处理程序的中间件：



```
var errHandler = webhelp.GenSym() // see the aside about context keys

func HandleErrWith(eh ErrHandler, h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    ctx := context.WithValue(whcompat.Context(req), errHandler, eh)
    h.ServeHTTP(w, whcompat.WithContext(req, ctx))
  })
}
```




Last, let’s make a function that will use the registered error handler for errors:



最后，让我们创建一个将使用注册的错误处理程序来处理错误的函数：



```
func HandleErr(w http.ResponseWriter, req *http.Request, err error) {
  if handler, ok := whcompat.Context(req).Value(errHandler).(ErrHandler);ok {
    handler.HandleError(w, req, err)
    return
  }
  log.Printf("error: %v", err)
  http.Error(w, "internal server error", http.StatusInternalServerError)
}
```




Now, as long as everything uses `HandleErr` to handle errors, our JSON API can handle errors with JSON responses, and our HTML endpoints can handle errors with HTML responses.

现在，只要一切都使用 `HandleErr` 来处理错误，我们的 JSON API 就可以处理带有 JSON 响应的错误，而我们的 HTML 端点可以处理带有 HTML 响应的错误。

Of course, the [wherr](https://godoc.org/gopkg.in/webhelp.v1/wherr) package implements this all for you, and the [whjson](https://godoc.org/gopkg.in/webhelp.v1/wherr) package even implements a friendly JSON API error handler.

当然，[wherr](https://godoc.org/gopkg.in/webhelp.v1/wherr) 包为您实现了这一切，而 [whjson](https://godoc.org/gopkg.in/webhelp.v1/wherr) 包甚至实现了友好的 JSON API 错误处理程序。

Here’s how you might use it:



以下是您可以如何使用它：



```
var userKey = webhelp.GenSym()

func RequireUser(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    user, err := GetUser(req)
    if err != nil {
      wherr.Handle(w, req, wherr.InternalServerError.New("failed to get user"))
      return
    }
    if user == nil {
      wherr.Handle(w, req, wherr.Unauthorized.New("no user found"))
      return
    }
    ctx := r.Context()
    ctx = context.WithValue(ctx, userKey, user)
    h.ServeHTTP(w, req.WithContext(ctx))
  })
}

func userpage(w http.ResponseWriter, req *http.Request) {
  user := req.Context().Value(userKey).(*User)
  w.Header().Set("Content-Type", "text/html")
  userpageTmpl.Execute(w, user)
}

func username(w http.ResponseWriter, req *http.Request) {
  user := req.Context().Value(userKey).(*User)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]interface{}{"user": user})
}

func main() {
  http.ListenAndServe(":8080", whmux.Dir{
    "api": wherr.HandleWith(whjson.ErrHandler,
      RequireUser(whmux.Dir{
        "username": http.HandlerFunc(username),
      })),
    "user": RequireUser(http.HandlerFunc(userpage)),
  })
}
```




### Aside about the spacemonkeygo/errors package

### 除了 spacemonkeygo/errors 包

The default [wherr.Handle](https://godoc.org/gopkg.in/webhelp.v1/wherr#Handle) implementation understands all of the [error classes defined in the wherr top level package](https://godoc.org/gopkg.in/webhelp.v1/wherr#pkg-variables).

默认的 [wherr.Handle](https://godoc.org/gopkg.in/webhelp.v1/wherr#Handle) 实现理解所有 [wherr 顶级包中定义的错误类](https://godoc.org/gopkg.in/webhelp.v1/wherr#pkg-variables)。

These error classes are implemented using the [spacemonkeygo/errors](https://godoc.org/github.com/spacemonkeygo/errors) library and the [spacemonkeygo/errors/errhttp](https://godoc.org/github.com/spacemonkeygo/errors/errhttp) extensions. You don't have to use this library or these errors, but the benefit is that your error instances can be extended to include HTTP status code messages and information, which once again, provides for a nice elimination of cross-cutting concerns in your error handling logic.

这些错误类是使用 [spacemonkeygo/errors](https://godoc.org/github.com/spacemonkeygo/errors) 库和 [spacemonkeygo/errors/errhttp](https://godoc.org/github.com/spacemonkeygo/errors/errhttp) 扩展。您不必使用这个库或这些错误，但好处是您的错误实例可以扩展为包括 HTTP 状态代码消息和信息，这再次提供了一个很好的消除错误中的横切关注点处理逻辑。

See the [spacemonkeygo/errors](https://godoc.org/github.com/spacemonkeygo/errors) package for more details.

有关更多详细信息，请参阅 [spacemonkeygo/errors](https://godoc.org/github.com/spacemonkeygo/errors) 包。

***Update 2018-04-19:** After a few years of use, my friend condensed some lessons we learned and the best parts of `spacemonkeygo/errors` into a new, more concise, better library, over at [github .com/zeebo/errs](https://github.com/zeebo/errs). Consider using that instead!*

*** 2018-04-19 更新：** 经过几年的使用，我的朋友将我们学到的一些经验教训和 `spacemonkeygo/errors` 最好的部分浓缩成一个新的、更简洁、更好的库，在 [github .com/zeebo/errs](https://github.com/zeebo/errs)。考虑使用它！*

## Sessions

## 会话

Go’s standard library has great support for cookies, but cookies by themselves aren’t usually what a developer thinks of when she thinks about sessions. Cookies are unencrypted, unauthenticated, and readable by the user, and perhaps you don’t want that with your session data.

Go 的标准库对 cookie 有很好的支持，但 cookie 本身通常不是开发人员在考虑会话时想到的。 Cookie 未加密、未经身份验证且可供用户读取，您可能不希望会话数据出现这种情况。

Further, sessions can be stored in cookies, but could also be stored in a database to provide features like session revocation and querying. There’s lots of potential details about the implementation of sessions.

此外，会话可以存储在 cookie 中，但也可以存储在数据库中以提供会话撤销和查询等功能。关于会话的实施有很多潜在的细节。

Request handlers, however, probably don’t care too much about the implementation details of the session. Request handlers usually just want a bucket of keys and values they can store safely and securely.

然而，请求处理程序可能不太关心会话的实现细节。请求处理程序通常只想要一桶可以安全可靠地存储的键和值。

The [whsess](https://godoc.org/gopkg.in/webhelp.v1/whsess) package implements middleware for registering an arbitrary session store (a default cookie-based session store is provided), and implements helpers for retrieving and saving new values into the session.

[whsess](https://godoc.org/gopkg.in/webhelp.v1/whsess) 包实现了用于注册任意会话存储的中间件（提供了默认的基于 cookie 的会话存储)，并实现了用于检索和将新值保存到会话中。

The default cookie-based session store implements encryption and authentication via the excellent [nacl/secretbox](https://godoc.org/golang.org/x/crypto/nacl/secretbox) package.

默认的基于 cookie 的会话存储通过优秀的 [nacl/secretbox](https://godoc.org/golang.org/x/crypto/nacl/secretbox) 包实现加密和身份验证。

Usage is like this:



用法是这样的：



```
func handler(w http.ResponseWriter, req *http.Request) {
  ctx := whcompat.Context(req)
  sess, err := whsess.Load(ctx, "namespace")
  if err != nil {
    wherr.Handle(w, req, err)
    return
  }
  if loggedIn, _ := sess.Values["logged_in"].(bool);loggedIn {
    views, _ := sess.Values["views"].(int64)
    sess.Values["views"] = views + 1
    sess.Save(w)
  }
}

func main() {
  http.ListenAndServe(":8080", whsess.HandlerWithStore(
    whsess.NewCookieStore(secret), http.HandlerFunc(handler)))
}
```




## Logging 

## 记录

The Go standard library by default doesn’t log incoming requests, outgoing responses, or even just what port the HTTP server is listening on.

默认情况下，Go 标准库不记录传入请求、传出响应，甚至不记录 HTTP 服务器正在侦听的端口。

The [whlog](https://godoc.org/gopkg.in/webhelp.v1/whlog) package implements all three. The [whlog.LogRequests](https://godoc.org/gopkg.in/webhelp.v1/whlog#LogRequests) middleware will log requests as they start. The [whlog.LogResponses](https://godoc.org/gopkg.in/webhelp.v1/whlog#LogResponses) middleware will log requests as they end, along with status code and timing information. [whlog.ListenAndServe](https://godoc.org/gopkg.in/webhelp.v1/whlog#ListenAndServe) will log the address the server ultimately listens on (if you specify “:0” as your address, a port will be randomly chosen, and [whlog.ListenAndServe](https://godoc.org/gopkg.in/webhelp.v1/whlog#ListenAndServe) will log it).

[whlog](https://godoc.org/gopkg.in/webhelp.v1/whlog) 包实现了所有三个。 [whlog.LogRequests](https://godoc.org/gopkg.in/webhelp.v1/whlog#LogRequests) 中间件将在请求开始时记录请求。 [whlog.LogResponses](https://godoc.org/gopkg.in/webhelp.v1/whlog#LogResponses) 中间件将在请求结束时记录请求，以及状态代码和时间信息。 [whlog.ListenAndServe](https://godoc.org/gopkg.in/webhelp.v1/whlog#ListenAndServe)将记录服务器最终侦听的地址（如果您指定“:0”作为您的地址，端口将随机选择，[whlog.ListenAndServe](https://godoc.org/gopkg.in/webhelp.v1/whlog#ListenAndServe) 将记录它)。

[whlog.LogResponses](https://godoc.org/gopkg.in/webhelp.v1/whlog#LogResponses) deserves special mention for how it does what it does. It uses the [whmon](https://godoc.org/gopkg.in/webhelp.v1/whmon) package to instrument the outgoing `http.ResponseWriter` to keep track of response information.

[whlog.LogResponses](https://godoc.org/gopkg.in/webhelp.v1/whlog#LogResponses) 值得特别提及它是如何做的。它使用 [whmon](https://godoc.org/gopkg.in/webhelp.v1/whmon) 包来检测传出的 `http.ResponseWriter` 以跟踪响应信息。

Usage is like this:



用法是这样的：



```
func main() {
  whlog.ListenAndServe(":8080", whlog.LogResponses(whlog.Default, handler))
}
```




### App engine logging

### 应用引擎日志记录

App engine logging is unconventional crazytown. The standard library logger doesn’t work by default on App Engine, because App Engine logs *require* the request context. This is unfortunate for libraries that don’t necessarily run on App Engine all the time, as their logging information doesn’t make it to the App Engine request-specific logger.

应用引擎日志记录是非常规的疯狂小镇。默认情况下，标准库记录器在 App Engine 上不起作用，因为 App Engine 记录 *require* 请求上下文。这对于不一定总是在 App Engine 上运行的库来说是不幸的，因为它们的日志信息不会进入 App Engine 请求特定的记录器。

Unbelievably, this is fixable with [whgls](https://godoc.org/gopkg.in/webhelp.v1/whgls), which uses my terrible, terrible (but recently improved) [Goroutine-local storage library](https://godoc.org/github.com/jtolds/gls) to store the request context on the current stack, register a new log output, and fix logging so standard library logging works with App Engine again.

令人难以置信的是，这可以通过 [whgls](https://godoc.org/gopkg.in/webhelp.v1/whgls)解决，它使用了我的可怕的、可怕的（但最近改进了)[Goroutine-本地存储库](https：//godoc.org/github.com/jtolds/gls) 将请求上下文存储在当前堆栈上，注册新的日志输出，并修复日志记录，以便标准库日志记录再次与 App Engine 一起使用。

## Template handling

## 模板处理

Go's standard library [html/template](https://golang.org/pkg/html/template/) package is excellent, but you'll be unsurprised to find there's a few tasks I do with it so commonly that I've written additional support code.

Go 的标准库 [html/template](https://golang.org/pkg/html/template/) 包非常出色，但你会毫不意外地发现我经常用它来完成一些任务，以至于我已经编写了额外的支持代码。

The [whtmpl](https://godoc.org/gopkg.in/webhelp.v1/whtmpl) package really does two things. First, it provides a number of useful helper methods for use within templates, and second, it takes some friction out of managing a large number of templates.

[whtmpl](https://godoc.org/gopkg.in/webhelp.v1/whtmpl) 包确实做了两件事。首先，它提供了许多在模板中使用的有用的辅助方法，其次，它在管理大量模板时会遇到一些麻烦。

When writing templates, one thing you can do is call out to other registered templates for small values. A good example might be some sort of list element. You can have a template that renders the list element, and then your template that renders your list can use the list element template in turn.

编写模板时，您可以做的一件事是调用其他已注册模板以获得小值。一个很好的例子可能是某种列表元素。您可以拥有一个呈现列表元素的模板，然后呈现您的列表的模板可以依次使用列表元素模板。

Use of another template within a template might look like this:

在模板中使用另一个模板可能如下所示：

```
<ul>
  {{ range .List }}
    {{ template "list_element" .}}
  {{ end }}
</ul>
```


You’re now rendering the `list_element` template with the list element from `.List`. But what if you want to also pass the current user `.User`? Unfortunately, you can only pass one argument from one template to another. If you have two arguments you want to pass to another template, with the standard library, you’re out of luck.

您现在正在使用`.List` 中的列表元素渲染`list_element` 模板。但是如果你还想传递当前用户 `.User` 怎么办？不幸的是，您只能将一个参数从一个模板传递到另一个模板。如果你有两个参数想要传递给另一个模板，使用标准库，那你就不走运了。

The [whtmpl](https://godoc.org/gopkg.in/webhelp.v1/whtmpl) package adds three helper functions to aid you here, `makepair`, `makemap`, and `makeslice` (more docs under the [whtmpl.Collection](https://godoc.org/gopkg.in/webhelp.v1/whtmpl#Collection)type). `makepair` is the simplest. It takes two arguments and constructs a [whtmpl.Pair](https://godoc.org/gopkg.in/webhelp.v1/whtmpl#Pair). Fixing our example above would look like this now:

[whtmpl](https://godoc.org/gopkg.in/webhelp.v1/whtmpl) 包添加了三个辅助函数来帮助你，`makepair`、`makemap` 和 `makeslice`（更多文档在[whtmpl.Collection](https://godoc.org/gopkg.in/webhelp.v1/whtmpl#Collection)类型)。 `makepair` 是最简单的。它需要两个参数并构造一个 [whtmpl.Pair](https://godoc.org/gopkg.in/webhelp.v1/whtmpl#Pair)。修复我们上面的例子现在看起来像这样：

```
<ul>
  {{ $user := .User }}
  {{ range .List }}
    {{ template "list_element" (makepair . $user) }}
  {{ end }}
</ul>
```


The second thing [whtmpl](https://godoc.org/gopkg.in/webhelp.v1/whtmpl) does is make defining lots of templates easy, by optionally automatically naming templates after the name of the file the template is defined in .

[whtmpl](https://godoc.org/gopkg.in/webhelp.v1/whtmpl) 所做的第二件事是使定义大量模板变得容易，通过选择在定义模板的文件名后自动命名模板.

For example, say you have three files.

例如，假设您有三个文件。

Here’s `pkg.go`:



这是`pkg.go`：



```
package views

import "gopkg.in/webhelp.v1/whtmpl"

var Templates = whtmpl.NewCollection()
```




Here’s `landing.go`:



这是`landing.go`：



```
package views

var _ = Templates.MustParse(`{{ template "header" . }}

   <h1>Landing!</h1>`)
```




And here’s `header.go`:



这是`header.go`：



```
package views

var _ = Templates.MustParse(`<title>My website!</title>`)
```




Now, you can import your new `views` package and render the `landing` template this easily:



现在，您可以导入新的 `views` 包并轻松渲染 `landing` 模板：



```
func handler(w http.ResponseWriter, req *http.Request) {
  views.Templates.Render(w, req, "landing", map[string]interface{}{})
}
```




## User authentication

##  用户认证

I’ve written two Webhelp-style authentication libraries that I end up using frequently.

我编写了两个最终经常使用的 Webhelp 风格的身份验证库。

The first is an OAuth2 library, [whoauth2](https://godoc.org/gopkg.in/go-webhelp/whoauth2.v1). I’ve written up [an example application that authenticates with Google, Facebook, and Github](https://github.com/go-webhelp/whoauth2/blob/v1/examples/group/main.go).

第一个是 OAuth2 库，[whoauth2](https://godoc.org/gopkg.in/go-webhelp/whoauth2.v1)。我已经编写了 [一个通过 Google、Facebook 和 Github 进行身份验证的示例应用程序](https://github.com/go-webhelp/whoauth2/blob/v1/examples/group/main.go)。

The second, [whgoth](https://godoc.org/gopkg.in/go-webhelp/whgoth.v1), is a wrapper around [markbates/goth](https://github.com/markbates/goth) . My portion isn’t quite complete yet (some fixes are still necessary for optional App Engine support), but will support more non-OAuth2 authentication sources (like Twitter) when it is done.

第二个，[whgoth](https://godoc.org/gopkg.in/go-webhelp/whgoth.v1)，是对[markbates/goth](https://github.com/markbates/goth) 的封装.我的部分还不是很完整（对于可选的 App Engine 支持仍然需要一些修复），但完成后将支持更多非 OAuth2 身份验证源（如 Twitter)。

## Route listing

## 路由列表

Surprise! If you've used [webhelp](https://godoc.org/gopkg.in/webhelp.v1) based handlers and middleware for your whole app, you automatically get route listing for free, via the [whroute](https://godoc.org/gopkg.in/webhelp.v1/whroute) package.

惊喜！如果您在整个应用程序中使用了基于 [webhelp](https://godoc.org/gopkg.in/webhelp.v1) 的处理程序和中间件，您将通过 [whroute](https://godoc.org/gopkg.in/webhelp.v1/whroute) 包。

My web serving code’s `main` method often has a form like this:



我的网络服务代码的 main 方法通常有这样的形式：



```
switch flag.Arg(0) {
case "serve":
  panic(whlog.ListenAndServe(*listenAddr, routes))
case "routes":
  whroute.PrintRoutes(os.Stdout, routes)
default:
  fmt.Printf("Usage: %s <serve|routes>\n", os.Args[0])
}
```




Here’s some example output:

下面是一些示例输出：

```
GET   /auth/_cb/
GET   /auth/login/
GET   /auth/logout/
GET /
GET   /account/apikeys/
POST  /account/apikeys/
GET   /project/<int>/
GET   /project/<int>/control/<int>/
POST  /project/<int>/control/<int>/sample/
GET   /project/<int>/control/
 Redirect: f(req)
POST  /project/<int>/control/
POST  /project/<int>/control_named/<string>/sample/
GET   /project/<int>/control_named/
 Redirect: f(req)
GET   /project/<int>/sample/<int>/
GET   /project/<int>/sample/<int>/similar[/<*>]
GET   /project/<int>/sample/
 Redirect: f(req)
POST  /project/<int>/search/
GET   /project/
 Redirect: /
POST  /project/
```


## Other little things

## 其他小东西

[webhelp](https://godoc.org/gopkg.in/webhelp.v1) has a number of other subpackages:

[webhelp](https://godoc.org/gopkg.in/webhelp.v1) 还有许多其他子包：

- [whparse](https://godoc.org/gopkg.in/webhelp.v1/whparse) assists in parsing optional request arguments.
- [whredir](https://godoc.org/gopkg.in/webhelp.v1/whredir) provides some handlers and helper methods for doing redirects in various cases.
- [whcache](https://godoc.org/gopkg.in/webhelp.v1/whcache) creates request-specific mutable storage for caching various computations and database loaded data. Mutability helps helper functions that aren’t used as middleware share data.
- [whfatal](https://godoc.org/gopkg.in/webhelp.v1/whfatal) uses panics to simplify early request handling termination. Probably avoid this package unless you want to anger other Go developers.

- [whparse](https://godoc.org/gopkg.in/webhelp.v1/whparse) 协助解析可选的请求参数。
- [whredir](https://godoc.org/gopkg.in/webhelp.v1/whredir) 提供了一些处理程序和辅助方法，用于在各种情况下进行重定向。
- [whcache](https://godoc.org/gopkg.in/webhelp.v1/whcache) 创建特定于请求的可变存储，用于缓存各种计算和数据库加载的数据。可变性有助于不用作中间件的辅助函数共享数据。
- [whfatal](https://godoc.org/gopkg.in/webhelp.v1/whfatal) 使用恐慌来简化早期请求处理终止。除非你想激怒其他 Go 开发人员，否则可能会避免使用这个包。

## Summary

##  概括

Designing your web project as a collection of composable middlewares goes quite a long way to simplify your code design, eliminate cross-cutting concerns, and create a more flexible development environment. Use my [webhelp](https://godoc.org/gopkg.in/webhelp.v1) package if it helps you.

将您的 Web 项目设计为可组合中间件的集合，对于简化代码设计、消除横切关注点和创建更灵活的开发环境大有帮助。如果对您有帮助，请使用我的 [webhelp](https://godoc.org/gopkg.in/webhelp.v1) 包。

Or don’t! Whatever! It’s still a free country last I checked.

或者不要！任何！我上次检查它仍然是一个自由的国家。

### Update

###  更新

Peter Kieltyka points me to his [Chi framework](https://github.com/pressly/chi), which actually does seem to do the right things with respect to middleware, handlers, and contexts - certainly much more so than all the other frameworks I've seen. So, shoutout to Peter and the team at Pressly! 

Peter Kieltyka 向我指出了他的 [Chi 框架](https://github.com/pressly/chi)，它实际上似乎在中间件、处理程序和上下文方面做了正确的事情——当然比所有我见过的其他框架。所以，向 Peter 和 Pressly 的团队致敬！

<iframe id="dsq-app8392" name="dsq-app8392" allowtransparency="true" scrolling="no" tabindex="0" title="Disqus" style="width: 100% !important; border: medium none !important; overflow: hidden !important; height: 0px !important; display: inline !important; box-sizing: border-box !important;" src="https://disqus.com/recommendations/?base=default&f=jtolds&t_u=https%3A%2F%2Fwww.jtolio.com%2F2017%2F01%2Fwriting-advanced-web-applications-with-go%2F&t_d= Writing%20Advanced%20Web%20Applications%20with%20Go%20%7C%20jtolio.com&t_t=Writing%20Advanced%20Web%20Applications%20with%20Go%20%7C%20jtolio.com#version=064141e2948b0e7f6218d4075662ea80" horizontalscrolling="no" verticalscrolling ="no" width="100%" frameborder="0"></iframe>

<iframe id="dsq-app8392" name="dsq-app8392" allowtransparency="true" scrolling="no" tabindex="0" title="Disqus" style="width: 100% !important; border: medium none !important; 溢出：隐藏 !important; 高度：0px !important; 显示：inline !important; box-sizing: border-box !important;" src="https://disqus.com/recommendations/?base=default&f=jtolds&t_u=https%3A%2F%2Fwww.jtolio.com%2F2017%2F01%2Fwriting-advanced-web-applications-with-go%2F&t_d=写作%20Advanced%20Web%20Applications%20with%20Go%20%7C%20jtolio.com&t_t=Writing%20Advanced%20Web%20Applications%20with%20Go%20%7C%20jtolio.com#version=064141e27ds60croll1e27fs20700000000" ="no" width="100%" frameborder="0"></iframe>

<iframe id="dsq-app5685" name="dsq-app5685" allowtransparency="true" scrolling="no" tabindex="0" title="Disqus" style="width: 1px !important; min-width: 100 % !important; border: medium none !important; overflow: hidden !important; height: 1382px !important;" src="https://disqus.com/embed/comments/?base=default&f=jtolds&t_u=http%3A%2F%2Fwww.jtolds.com%2Fwriting%2F2017%2F01%2Fwriting-advanced-web-applications-with- go%2F&t_d=Writing%20Advanced%20Web%20Applications%20with%20Go%20%7C%20jtolio.com&t_t=Writing%20Advanced%20Web%20Applications%20with%20Go%20%7C%20jtolio.com&s_o=default#version=9bdb65de27b881f62b84ef54f46d1575" horizontalscrolling="no" verticalscrolling="no" width="100%" frameborder="0"></iframe> 

<iframe id="dsq-app5685" name="dsq-app5685" allowtransparency="true" scrolling="no" tabindex="0" title="Disqus" style="width: 1px !important; min-width: 100 % !important; 边框：中等 无 !important; 溢出：隐藏 !important; 高度：1382px !important;" src="https://disqus.com/embed/comments/?base=default&f=jtolds&t_u=http%3A%2F%2Fwww.jtolds.com%2Fwriting%2F2017%2F01%2Fwriting-advanced-web-applications-with- go%2F&t_d=Writing%20Advanced%20Web%20Applications%20with%20Go%20%7C%20jtolio.com&t_t=Writing%20Advanced%20Web%20Applications%20with%20Go%20%7C%20jtoliodefault.com=8#4b5d4b5d4b54bversion7fb54bversion7f85d4b7f2010000000000 horizontalscrolling="no" verticalscrolling="no" width="100%" frameborder="0"></iframe>

