# Different approaches to HTTP routing in Go

# Go 中 HTTP 路由的不同方法

July 2020 From: https://benhoyt.com/writings/go-routing/

There are many ways to do HTTP path routing in Go – for better or worse. There’s the standard library’s [`http.ServeMux`](https://golang.org/pkg/net/http/#ServeMux),but it only supports basic prefix matching. There are many ways to do  more advanced routing yourself, including Axel Wagner's interesting [`ShiftPath` technique](https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html). And then of course there are lots of third-party router libraries. In  this article I’m going to do a comparison of several custom techniques  and some off-the-shelf packages.

有很多方法可以在 Go 中进行 HTTP 路径路由——无论好坏。有标准库的 [`http.ServeMux`](https://golang.org/pkg/net/http/#ServeMux)，但它只支持基本的前缀匹配。有很多方法可以自己做更高级的路由，包括 Axel Wagner 有趣的 `ShiftPath` 技术。当然还有很多第三方路由器库。在本文中，我将对几种自定义技术和一些现成的软件包进行比较。

I’ll be upfront about my biases: I like simple and clear code, and  I’m a bit allergic to large dependencies (and sometimes those are in  tension). Most libraries with “framework” in the title don’t do it for  me, though I’m not opposed to using well-maintained libraries that do  one or two things well.

我会坦率地说明我的偏见：我喜欢简单明了的代码，我对大的依赖有点过敏（有时这些依赖很紧张）。大多数标题中带有“框架”的库并不适合我，尽管我不反对使用维护良好的库来做好一两件事。

My goal here is to route the same 11 URLs with eight different  approaches. These URLs are based on a subset of URLs in a web  application I maintain. They use `GET` and `POST`, but they’re not particularly RESTful or well designed – the kind of  messiness you find in real-world systems. Here are the methods and URLs:

我的目标是使用八种不同的方法路由相同的 11 个 URL。这些 URL 基于我维护的 Web 应用程序中的 URL 子集。它们使用“GET”和“POST”，但它们并不是特别 RESTful 或精心设计的——你在现实世界的系统中发现的那种混乱。以下是方法和网址：

```
GET / # home
GET  /contact                               # contact
GET  /api/widgets                           # apiGetWidgets
POST /api/widgets                           # apiCreateWidget
POST /api/widgets/:slug                     # apiUpdateWidget
POST /api/widgets/:slug/parts               # apiCreateWidgetPart
POST /api/widgets/:slug/parts/:id/update    # apiUpdateWidgetPart
POST /api/widgets/:slug/parts/:id/delete    # apiDeleteWidgetPart
GET  /:slug                                 # widget
GET  /:slug/admin                           # widgetAdmin
POST /:slug/image                           # widgetImage
```


The `:slug` is a URL-friendly widget identifier like `foo-bar`, and the `:id` is a positive integer like `1234`. Each routing approach should match on the exact URL – trailing slashes  will return 404 Not Found (redirecting them is also a fine decision, but I’m not doing that here). Each router should handle the specified  method (`GET` or `POST`) and reject the others with a 405 Method Not Allowed response. I wrote some [table-driven tests](https://github.com/benhoyt/go-routing/blob/master/main_test.go) to ensure that all the routers do the right thing.

`:slug` 是一个 URL 友好的小部件标识符，比如 `foo-bar`，而 `:id` 是一个像 `1234` 的正整数。每种路由方法都应该与确切的 URL 匹配——尾部斜杠将返回 404 Not Found（重定向它们也是一个很好的决定，但我不会在这里这样做）。每个路由器都应该处理指定的方法（`GET` 或 `POST`），并使用 405 Method Not Allowed 响应拒绝其他方法。我写了一些 [表驱动测试](https://github.com/benhoyt/go-routing/blob/master/main_test.go) 以确保所有路由器都做正确的事情。

In the rest of this article I'll present code for the various  approaches and discuss some pros and cons of each (all the code is in  the [benhoyt/go-routing](https://github.com/benhoyt/go-routing) repo). There’s a lot of code, but all of it is fairly straight-forward  and should be easy to skim. You can use the following links to skip down to a particular technique. First, the five custom techniques:

在本文的其余部分，我将介绍各种方法的代码并讨论每种方法的一些优缺点。有很多代码，但所有代码都很简单，应该很容易浏览。您可以使用以下链接跳到特定技术。一、五种定制技巧：

- [Regex table](https://benhoyt.com/writings/go-routing/#regex-table): loop through pre-compiled regexes and pass matches using the request context
- [Regex switch](https://benhoyt.com/writings/go-routing/#regex-switch):a switch statement with cases that call a regex-based `match()` helper which scans path parameters into variables
- [Pattern matcher](https://benhoyt.com/writings/go-routing/#pattern-matcher): similar to the above, but using a simple pattern matching function instead of regexes
- [Split switch](https://benhoyt.com/writings/go-routing/#split-switch): split the path on `/` and then switch on the contents of the path segments
- [ShiftPath](https://benhoyt.com/writings/go-routing/#shiftpath): Axel Wagner’s hierarchical `ShiftPath` technique

- [Regex table](https://benhoyt.com/writings/go-routing/#regex-table)：循环预编译的正则表达式并使用请求上下文传递匹配项
- [Regex switch](https://benhoyt.com/writings/go-routing/#regex-switch)：一个switch 语句，包含调用基于正则表达式的 `match()` 帮助程序，它将路径参数扫描到变量中
- [模式匹配器](https://benhoyt.com/writings/go-routing/#pattern-matcher)：与上面类似，但使用简单的模式匹配函数而不是正则表达式
- [拆分开关](https://benhoyt.com/writings/go-routing/#split-switch)：在`/`上拆分路径，然后打开路径段的内容
- [ShiftPath](https://benhoyt.com/writings/go-routing/#shiftpath)：Axel Wagner 的分层“ShiftPath”技术

And three versions using third-party router packages:

以及使用第三方路由器包的三个版本：

- [Chi](https://benhoyt.com/writings/go-routing/#chi): uses `github.com/go-chi/chi`
- [Gorilla](https://benhoyt.com/writings/go-routing/#gorilla): uses `github.com/gorilla/mux`
- [Pat](https://benhoyt.com/writings/go-routing/#pat): uses `github.com/bmizerany/pat` 

- [Chi](https://benhoyt.com/writings/go-routing/#chi)：使用`github.com/go-chi/chi`
- [Gorilla](https://benhoyt.com/writings/go-routing/#gorilla)：使用`github.com/gorilla/mux`
- [Pat](https://benhoyt.com/writings/go-routing/#pat)：使用`github.com/bmizerany/pat`

I also tried [httprouter](https://github.com/julienschmidt/httprouter),which is supposed to be really fast, but it [can't handle](https://github.com/julienschmidt/httprouter/issues/73) URLs with overlapping prefixes like `/contact` and `/:slug`. Arguably this is bad URL design anyway, but a lot of real-world web apps do it, so I think this is quite limiting.

我也试过 [httprouter](https://github.com/julienschmidt/httprouter)，它应该很快，但它[无法处理](https://github.com/julienschmidt/httprouter/issues/73) 带有重叠前缀的 URL，如 `/contact` 和 `/:slug`。可以说这无论如何都是糟糕的 URL 设计，但很多现实世界的网络应用程序都这样做，所以我认为这是非常有限的。

There are many other third-party router packages or “web frameworks”, but these three bubbled to the top in my searches (and I believe  they’re fairly representative).

还有许多其他第三方路由器包或“网络框架”，但这三个在我的搜索中排在前列（我相信它们具有相当的代表性）。

In this comparison I’m not concerned about speed. Most of the approaches loop or `switch` through a list of routes (in contrast to fancy trie-lookup structures). All of these approaches only add a few *microseconds* to the request time (see [benchmarks](https://benhoyt.com/writings/go-routing/#benchmarks)), and that isn't an issue in any of the web applications I've worked on.

在这个比较中，我不关心速度。大多数方法通过路由列表循环或“切换”（与花哨的查找树结构相反）。所有这些方法都只为请求时间增加了几微秒*（参见 [benchmarks](https://benhoyt.com/writings/go-routing/#benchmarks)），这在任何一个中都不是问题我工作过的网络应用程序。

## Regex table

## 正则表

The first approach I want to look at is the method I use in the  current version of my web application – it’s the first thing that came  to mind when I was learning Go a few years back, and I still think it’s a pretty good approach.

我想看的第一种方法是我在当前版本的 web 应用程序中使用的方法——这是我几年前学习 Go 时想到的第一件事，我仍然认为这是一个很好的方法。

It's basically a table of pre-compiled [`regexp`](https://golang.org/pkg/regexp/)objects with a little 21-line routing function that loops through them, and calls the first one that matches both the path and the HTTP method. Here are the routes and the `Serve()` routing function:

它基本上是一个预编译的 [`regexp`](https://golang.org/pkg/regexp/)对象表，带有一个小的 21 行路由函数，循环遍历它们，并调用第一个匹配路径和 HTTP 方法。以下是路由和 `Serve()` 路由函数：

```
var routes = []route{
    newRoute("GET", "/", home),
    newRoute("GET", "/contact", contact),
    newRoute("GET", "/api/widgets", apiGetWidgets),
    newRoute("POST", "/api/widgets", apiCreateWidget),
    newRoute("POST", "/api/widgets/([^/]+)", apiUpdateWidget),
    newRoute("POST", "/api/widgets/([^/]+)/parts", apiCreateWidgetPart),
    newRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/update", apiUpdateWidgetPart),
    newRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/delete", apiDeleteWidgetPart),
    newRoute("GET", "/([^/]+)", widget),
    newRoute("GET", "/([^/]+)/admin", widgetAdmin),
    newRoute("POST", "/([^/]+)/image", widgetImage),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
    return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
    method  string
    regex   *regexp.Regexp
    handler http.HandlerFunc
}

func Serve(w http.ResponseWriter, r *http.Request) {
    var allow []string
    for _, route := range routes {
        matches := route.regex.FindStringSubmatch(r.URL.Path)
        if len(matches) > 0 {
            if r.Method != route.method {
                allow = append(allow, route.method)
                continue
            }
            ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
            route.handler(w, r.WithContext(ctx))
            return
        }
    }
    if len(allow) > 0 {
        w.Header().Set("Allow", strings.Join(allow, ", "))
        http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        return
    }
    http.NotFound(w, r)
}
```


Path parameters are handled by adding the `matches` slice to the request context, so the handlers can pick them up from  there. I’ve defined a custom context key type, as well as a `getField` helper function that’s used inside the handlers:

路径参数是通过将 `matches` 切片添加到请求上下文来处理的，因此处理程序可以从那里获取它们。我已经定义了一个自定义上下文键类型，以及一个在处理程序中使用的 `getField` 辅助函数：

```
type ctxKey struct{}

func getField(r *http.Request, index int) string {
    fields := r.Context().Value(ctxKey{}).([]string)
    return fields[index]
}
```


A typical handler with path parameters looks like this:

带有路径参数的典型处理程序如下所示：

```
// Handles POST /api/widgets/([^/]+)/parts/([0-9]+)/update
func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
    slug := getField(r, 0)
    id, _ := strconv.Atoi(getField(r, 1))
    fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}
```




I haven’t checked the error returned by `Atoi()`, because the regex for the ID parameter only matches digits: `[0-9]+`. Of course, there’s still no guarantee the object exists in the database – that still needs to be done in the handler. (If the number is too  large, `Atoi` will return an error, but in that case the `id` will be zero and the database lookup will fail, so there’s no need for an extra check.)

我没有检查`Atoi()`返回的错误，因为ID参数的正则表达式只匹配数字：`[0-9]+`。当然，仍然不能保证对象存在于数据库中——这仍然需要在处理程序中完成。 （如果数字太大，`Atoi` 将返回错误，但在这种情况下，`id` 将为零并且数据库查找将失败，因此无需额外检查。）

An alternative to passing the fields using context is to make each `route.handler` a function that takes the fields as a `[]string` and returns an `http.HandleFunc` closure that closes over the `fields` parameter. The `Serve` function would then instantiate and call the closure as follows:

使用上下文传递字段的另一种方法是使每个 `route.handler` 成为一个函数，该函数将字段作为 `[]string` 并返回一个关闭 `fields` 参数的 `http.HandleFunc` 闭包。然后`Serve` 函数将实例化并调用闭包，如下所示：

```
handler := route.handler(matches[1:])
handler(w, r)
```


Then each handler would look like this:

然后每个处理程序看起来像这样：

```
func apiUpdateWidgetPart(fields []string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        slug := fields[0]
        id, _ := strconv.Atoi(fields[1])
        fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
    }
}
```


I slightly prefer the context approach, as it keeps the handler signatures simple `http.HandlerFunc`s, and also avoids a nested function for each handler definition.

我稍微更喜欢上下文方法，因为它使处理程序签名保持简单`http.HandlerFunc`s，并且还避免了每个处理程序定义的嵌套函数。

There’s nothing particularly clever about the regex table approach,  and it’s similar to how a number of the third-party packages work. But  it’s so simple it only takes a few lines of code and a few minutes to  write. It’s also easy to modify if you need to: for example, to add  logging, change the error responses to JSON, and so on.

regex table 方法没有什么特别聪明的地方，它类似于许多第三方包的工作方式。但它是如此简单，只需几行代码和几分钟即可编写。如果需要，也可以轻松修改：例如，添加日志记录、将错误响应更改为 JSON 等。

[Full regex table code on GitHub.](https://github.com/benhoyt/go-routing/blob/master/retable/route.go)

[GitHub 上的完整正则表达式代码。](https://github.com/benhoyt/go-routing/blob/master/retable/route.go)

## Regex switch

## 正则表达式开关

The second approach still uses regexes, but with a simple imperative `switch` statement and a `match()` helper to go through the matches. The advantage of this approach is  that you can call other functions or test other things in each `case`. Also, the signature of the `match` function allows you to “scan” path parameters into variables in order  to pass them to the handlers more directly. Here are the routes and the `match()` function:

第二种方法仍然使用正则表达式，但使用一个简单的命令式 `switch` 语句和一个 `match()` 帮助器来完成匹配。这种方法的优点是您可以在每个 `case` 中调用其他函数或测试其他东西。此外，`match` 函数的签名允许您将路径参数“扫描”到变量中，以便更直接地将它们传递给处理程序。以下是路由和 `match()` 函数：

```
func Serve(w http.ResponseWriter, r *http.Request) {
    var h http.Handler
    var slug string
    var id int

    p := r.URL.Path
    switch {
    case match(p, "/"):
        h = get(home)
    case match(p, "/contact"):
        h = get(contact)
    case match(p, "/api/widgets") && r.Method == "GET":
        h = get(apiGetWidgets)
    case match(p, "/api/widgets"):
        h = post(apiCreateWidget)
    case match(p, "/api/widgets/([^/]+)", &slug):
        h = post(apiWidget{slug}.update)
    case match(p, "/api/widgets/([^/]+)/parts", &slug):
        h = post(apiWidget{slug}.createPart)
    case match(p, "/api/widgets/([^/]+)/parts/([0-9]+)/update", &slug, &id):
        h = post(apiWidgetPart{slug, id}.update)
    case match(p, "/api/widgets/([^/]+)/parts/([0-9]+)/delete", &slug, &id):
        h = post(apiWidgetPart{slug, id}.delete)
    case match(p, "/([^/]+)", &slug):
        h = get(widget{slug}.widget)
    case match(p, "/([^/]+)/admin", &slug):
        h = get(widget{slug}.admin)
    case match(p, "/([^/]+)/image", &slug):
        h = post(widget{slug}.image)
    default:
        http.NotFound(w, r)
        return
    }
    h.ServeHTTP(w, r)
}

// match reports whether path matches regex ^pattern$, and if it matches,
// assigns any capture groups to the *string or *int vars.
func match(path, pattern string, vars ...interface{}) bool {
    regex := mustCompileCached(pattern)
    matches := regex.FindStringSubmatch(path)
    if len(matches) <= 0 {
        return false
    }
    for i, match := range matches[1:] {
        switch p := vars[i].(type) {
        case *string:
            *p = match
        case *int:
            n, err := strconv.Atoi(match)
            if err != nil {
                return false
            }
            *p = n
        default:
            panic("vars must be *string or *int")
        }
    }
    return true
}
```




I must admit to being quite fond of this approach. I like how simple  and direct it is, and I think the scan-like behaviour for path  parameters is clean. The scanning inside `match()` detects the type, and converts from string to integer if needed. It only supports `string` and `int` right now, which is probably all you need for most routes, but it’d be easy to add more types if you need to.

我必须承认我非常喜欢这种方法。我喜欢它的简单和直接，而且我认为路径参数的类似扫描的行为是干净的。 `match()` 内部的扫描检测类型，并在需要时从字符串转换为整数。它现在只支持 `string` 和 `int`，这可能是大多数路由所需要的，但如果需要，添加更多类型也很容易。

Here’s what a handler with path parameters looks like (to avoid repetition, I’ve used the `apiWidgetPart` struct for all the handlers that take those two parameters):

这是带有路径参数的处理程序的样子（为了避免重复，我对所有采用这两个参数的处理程序使用了 `apiWidgetPart` 结构）：

```
type apiWidgetPart struct {
    slug string
    id   int
}

func (h apiWidgetPart) update(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", h.slug, h.id)
}

func (h apiWidgetPart) delete(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "apiDeleteWidgetPart %s %d\n", h.slug, h.id)
}
```


Note the `get()` and `post()` helper functions, which are essentially simple middleware that check the request method as follows:

注意 `get()` 和 `post()` 辅助函数，它们本质上是简单的中间件，用于检查请求方法，如下所示：

```
// get takes a HandlerFunc and wraps it to only allow the GET method
func get(h http.HandlerFunc) http.HandlerFunc {
    return allowMethod(h, "GET")
}

// post takes a HandlerFunc and wraps it to only allow the POST method
func post(h http.HandlerFunc) http.HandlerFunc {
    return allowMethod(h, "POST")
}

// allowMethod takes a HandlerFunc and wraps it in a handler that only
// responds if the request method is the given method, otherwise it
// responds with HTTP 405 Method Not Allowed.
func allowMethod(h http.HandlerFunc, method string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if method != r.Method {
            w.Header().Set("Allow", method)
            http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
            return
        }
        h(w, r)
    }
}
```


One of the slightly awkward things is how it works for paths that  handle more than one method. There are probably different ways to do it, but I currently test the method explicitly in the first route – the `get()` wrapper is not strictly necessary here, but I’ve included it for consistency:

稍微有点尴尬的事情之一是它如何处理处理多个方法的路径。可能有不同的方法来做到这一点，但我目前在第一条路线中明确地测试了该方法——`get()` 包装器在这里不是绝对必要的，但为了保持一致性，我已经包含了它：

```
     case match(p, "/api/widgets") && r.Method == "GET":
        h = get(apiGetWidgets)
    case match(p, "/api/widgets"):
        h = post(apiCreateWidget)
```


At first I included the HTTP method matching in the `match()` helper, but that makes it more difficult to return 405 Method Not Allowed responses properly.

起初我在 `match()` 帮助器中包含了 HTTP 方法匹配，但这使得正确返回 405 Method Not Allowed 响应变得更加困难。

One other aspect of this approach is the lazy regex compiling. We could just call `regexp.MustCompile`, but that would re-compile each regex on every reqeust. Instead, I’ve added a concurrency-safe `mustCompileCached` function that means the regexes are only compiled the first time they’re used:

这种方法的另一个方面是惰性正则表达式编译。我们可以只调用`regexp.MustCompile`，但这会在每次请求时重新编译每个正则表达式。相反，我添加了一个并发安全的 `mustCompileCached` 函数，这意味着正则表达式只在第一次使用时编译：

```
var (
    regexen = make(map[string]*regexp.Regexp)
    relock  sync.Mutex
)

func mustCompileCached(pattern string) *regexp.Regexp {
    relock.Lock()
    defer relock.Unlock()

    regex := regexen[pattern]
    if regex == nil {
        regex = regexp.MustCompile("^" + pattern + "$")
        regexen[pattern] = regex
    }
    return regex
}
```


Overall, despite liking the clarity of this approach and the scan-like `match()` helper, a point against it is the messiness required to cache the regex compilation.

总的来说，尽管喜欢这种方法的清晰性和类似扫描的“match()”帮助器，但反对它的一点是缓存正则表达式编译所需的混乱。

[Full regex switch code on GitHub.](https://github.com/benhoyt/go-routing/blob/master/reswitch/route.go)

[GitHub 上的完整正则表达式切换代码。](https://github.com/benhoyt/go-routing/blob/master/reswitch/route.go)

## Pattern matcher

## 模式匹配器

This approach is similar to the regex switch method, but instead of regexes it uses a simple, custom pattern matcher.

这种方法类似于正则表达式切换方法，但它使用简单的自定义模式匹配器代替正则表达式。

The patterns supplied to the custom `match()` function handle one wildcard character, `+`, which matches (and captures) any characters till the next `/` in the request path. This is of course much less powerful than regex  matching, but generally I’ve not needed anything more than “match till  next slash” in my routes. Here is what the routes and match code look  like:

提供给自定义 `match()` 函数的模式处理一个通配符 `+`，它匹配（并捕获）任何字符，直到请求路径中的下一个 `/`。这当然比正则表达式匹配要弱得多，但通常我在我的路线中不需要比“匹配直到下一个斜线”更多的东西。以下是路由和匹配代码的样子：

```
func Serve(w http.ResponseWriter, r *http.Request) {
    var h http.Handler
    var slug string
    var id int

    p := r.URL.Path
    switch {
    case match(p, "/"):
        h = get(home)
    case match(p, "/contact"):
        h = get(contact)
    case match(p, "/api/widgets") && r.Method == "GET":
        h = get(apiGetWidgets)
    case match(p, "/api/widgets"):
        h = post(apiCreateWidget)
    case match(p, "/api/widgets/+", &slug):
        h = post(apiWidget{slug}.update)
    case match(p, "/api/widgets/+/parts", &slug):
        h = post(apiWidget{slug}.createPart)
    case match(p, "/api/widgets/+/parts/+/update", &slug, &id):
        h = post(apiWidgetPart{slug, id}.update)
    case match(p, "/api/widgets/+/parts/+/delete", &slug, &id):
        h = post(apiWidgetPart{slug, id}.delete)
    case match(p, "/+", &slug):
        h = get(widget{slug}.widget)
    case match(p, "/+/admin", &slug):
        h = get(widget{slug}.admin)
    case match(p, "/+/image", &slug):
        h = post(widget{slug}.image)
    default:
        http.NotFound(w, r)
        return
    }
    h.ServeHTTP(w, r)
}

// match reports whether path matches the given pattern, which is a
// path with '+' wildcards wherever you want to use a parameter.Path
// parameters are assigned to the pointers in vars (len(vars) must be
// the number of wildcards), which must be of type *string or *int.
func match(path, pattern string, vars ...interface{}) bool {
    for ;pattern != "" && path != "";pattern = pattern[1:] {
        switch pattern[0] {
        case '+':
            // '+' matches till next slash in path
            slash := strings.IndexByte(path, '/')
            if slash < 0 {
                slash = len(path)
            }
            segment := path[:slash]
            path = path[slash:]
            switch p := vars[0].(type) {
            case *string:
                *p = segment
            case *int:
                n, err := strconv.Atoi(segment)
                if err != nil ||n < 0 {
                    return false
                }
                *p = n
            default:
                panic("vars must be *string or *int")
            }
            vars = vars[1:]
        case path[0]:
            // non-'+' pattern byte must match path byte
            path = path[1:]
        default:
            return false
        }
    }
    return path == "" && pattern == ""
}
```




Other than that, the `get()` and `post()` helpers, as well as the handlers themselves, are identical to the regex switch method. I quite like this approach (and it’s efficient), but the byte-by-byte matching code was a little fiddly to write – definitely  not as simple as calling `regex.FindStringSubmatch()`.

除此之外，`get()` 和 `post()` 助手以及处理程序本身与正则表达式 switch 方法相同。我非常喜欢这种方法（而且它很有效），但是逐字节匹配的代码编写起来有点繁琐——绝对不像调用 regex.FindStringSubmatch() 那样简单。

[Full pattern matcher code on GitHub.](https://github.com/benhoyt/go-routing/blob/master/match/route.go)

[GitHub 上的完整模式匹配器代码。](https://github.com/benhoyt/go-routing/blob/master/match/route.go)

## Split switch

## 分割开关

This approach simply splits the request path on `/` and then uses a `switch` with `case` statements that compare the number of path segments and the content of  each segment. It’s direct and simple, but also a bit error-prone, with  lots of hard-coded lengths and indexes. Here is the code:

这种方法简单地在`/` 上拆分请求路径，然后使用带有`case` 语句的`switch` 来比较路径段的数量和每个段的内容。它直接而简单，但也有点容易出错，有很多硬编码的长度和索引。这是代码：

```
func Serve(w http.ResponseWriter, r *http.Request) {
    // Split path into slash-separated parts, for example, path "/foo/bar"
    // gives p==["foo", "bar"] and path "/" gives p==[""].
    p := strings.Split(r.URL.Path, "/")[1:]
    n := len(p)

    var h http.Handler
    var id int
    switch {
    case n == 1 && p[0] == "":
        h = get(home)
    case n == 1 && p[0] == "contact":
        h = get(contact)
    case n == 2 && p[0] == "api" && p[1] == "widgets" && r.Method == "GET":
        h = get(apiGetWidgets)
    case n == 2 && p[0] == "api" && p[1] == "widgets":
        h = post(apiCreateWidget)
    case n == 3 && p[0] == "api" && p[1] == "widgets" && p[2] != "":
        h = post(apiWidget{p[2]}.update)
    case n == 4 && p[0] == "api" && p[1] == "widgets" && p[2] != "" && p[3] == "parts":
        h = post(apiWidget{p[2]}.createPart)
    case n == 6 && p[0] == "api" && p[1] == "widgets" && p[2] != "" && p[3] == "parts" && isId(p[4], &id) && p[5] == "update":
        h = post(apiWidgetPart{p[2], id}.update)
    case n == 6 && p[0] == "api" && p[1] == "widgets" && p[2] != "" && p[3] == "parts" && isId(p[4], &id) && p[5] == "delete":
        h = post(apiWidgetPart{p[2], id}.delete)
    case n == 1:
        h = get(widget{p[0]}.widget)
    case n == 2 && p[1] == "admin":
        h = get(widget{p[0]}.admin)
    case n == 2 && p[1] == "image":
        h = post(widget{p[0]}.image)
    default:
        http.NotFound(w, r)
        return
    }
    h.ServeHTTP(w, r)
}
```


The handlers are identical to the other `switch`-based methods, as are the `get` and `post` helpers. The only helper here is the `isId` function, which checks that the ID segments are in fact positive integers:

处理程序与其他基于“switch”的方法相同，“get”和“post”助手也是如此。这里唯一的助手是 `isId` 函数，它检查 ID 段实际上是正整数：

```
func isId(s string, p *int) bool {
    id, err := strconv.Atoi(s)
    if err != nil ||id <= 0 {
        return false
    }
    *p = id
    return true
}
```


So while I like the bare-bones simplicity of this approach – just  basic string equality comparisons – the verbosity of the matching and  the error-prone integer constants would make me think twice about  actually using it for anything but very simple routing.

因此，虽然我喜欢这种方法的基本简单性——只是基本的字符串相等性比较——但匹配的冗长和容易出错的整数常量会让我三思而后行，除了非常简单的路由之外，其他任何事情都会使用它。

[Full split switch code on GitHub.](https://github.com/benhoyt/go-routing/blob/master/split/route.go)

[GitHub 上的完整拆分开关代码。](https://github.com/benhoyt/go-routing/blob/master/split/route.go)

## ShiftPath

## 移动路径

Axel Wagner wrote a blog article, [How to not use an http-router in go](https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html), in which he maintains that routers (third party or otherwise) should not be used. He presents a technique involving a small `ShiftPath()` helper that returns the first path segment, and shifts the rest of the  URL down. The current handler switches on the first path segment, then  delegates to sub-handlers which do the same thing on the rest of the  URL.

Axel Wagner 写了一篇博客文章，其中他认为不应使用路由器（第三方或其他方式）。他提出了一种技术，涉及一个小的“ShiftPath()”帮助器，它返回第一个路径段，并将 URL 的其余部分向下移动。当前处理程序打开第一个路径段，然后委托子处理程序对 URL 的其余部分执行相同的操作。

Let’s see what Axel’s technique looks like for a subset our set of URLs:

让我们看看 Axel 的技术对于我们的 URL 集的一个子集是什么样的：

```
func serve(w http.ResponseWriter, r *http.Request) {
    var head string
    head, r.URL.Path = shiftPath(r.URL.Path)
    switch head {
    case "":
        serveHome(w, r)
    case "api":
        serveApi(w, r)
    case "contact":
        serveContact(w, r)
    default:
        widget{head}.ServeHTTP(w, r)
    }
}

// shiftPath splits the given path into the first segment (head) and
// the rest (tail).For example, "/foo/bar/baz" gives "foo", "/bar/baz".
func shiftPath(p string) (head, tail string) {
    p = path.Clean("/" + p)
    i := strings.Index(p[1:], "/") + 1
    if i <= 0 {
        return p[1:], "/"
    }
    return p[1:i], p[i:]
}

// ensureMethod is a helper that reports whether the request's method is
// the given method, writing an Allow header and a 405 Method Not Allowed
// if not.The caller should return from the handler if this returns false.
func ensureMethod(w http.ResponseWriter, r *http.Request, method string) bool {
    if method != r.Method {
        w.Header().Set("Allow", method)
        http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        return false
    }
    return true
}

// ...

// Handles /api and below
func serveApi(w http.ResponseWriter, r *http.Request) {
    var head string
    head, r.URL.Path = shiftPath(r.URL.Path)
    switch head {
    case "widgets":
        serveApiWidgets(w, r)
    default:
        http.NotFound(w, r)
    }
}

// Handles /api/widgets and below
func serveApiWidgets(w http.ResponseWriter, r *http.Request) {
    var head string
    head, r.URL.Path = shiftPath(r.URL.Path)
    switch head {
    case "":
        if r.Method == "GET" {
            serveApiGetWidgets(w, r)
        } else {
            serveApiCreateWidget(w, r)
        }
    default:
        apiWidget{head}.ServeHTTP(w, r)
    }
}

// Handles GET /api/widgets
func serveApiGetWidgets(w http.ResponseWriter, r *http.Request) {
    if !ensureMethod(w, r, "GET") {
        return
    }
    fmt.Fprint(w, "apiGetWidgets\n")
}

// Handles POST /api/widgets
func serveApiCreateWidget(w http.ResponseWriter, r *http.Request) {
    if !ensureMethod(w, r, "POST") {
        return
    }
    fmt.Fprint(w, "apiCreateWidget\n")
}

type apiWidget struct {
    slug string
}

// Handles /api/widgets/:slug and below
func (h apiWidget) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var head string
    head, r.URL.Path = shiftPath(r.URL.Path)
    switch head {
    case "":
        h.serveUpdate(w, r)
    case "parts":
        h.serveParts(w, r)
    default:
        http.NotFound(w, r)
    }
}

func (h apiWidget) serveUpdate(w http.ResponseWriter, r *http.Request) {
    if !ensureMethod(w, r, "POST") {
        return
    }
    fmt.Fprintf(w, "apiUpdateWidget %s\n", h.slug)
}

func (h apiWidget) serveParts(w http.ResponseWriter, r *http.Request) {
    var head string
    head, r.URL.Path = shiftPath(r.URL.Path)
    switch head {
    case "":
        h.serveCreatePart(w, r)
    default:
        id, err := strconv.Atoi(head)
        if err != nil ||id <= 0 {
            http.NotFound(w, r)
            return
        }
        apiWidgetPart{h.slug, id}.ServeHTTP(w, r)
    }
}

// ...
```




With this router, I wrote a [`noTrailingSlash`](https://github.com/benhoyt/go-routing/blob/9a2fa7a643ecb5681f504b95064d948ee2177c9a/shiftpath/route.go#L54-L67) decorator to ensure Not Found is returned by URLs with a trailing  slash, as our URL spec defines those as invalid. The ShiftPath approach  doesn’t distinguish between no trailing slash and trailing slash, and I  can’t find a simple way to make it do that. I think a decorator is a  reasonable approach for this, rather than doing it explicitly in every  route – in a given web app, you'd probably want to either allow trailing slashes and redirect them, or return Not Found as I've done here .

使用这个路由器，我写了一个 [`noTrailingSlash`](https://github.com/benhoyt/go-routing/blob/9a2fa7a643ecb5681f504b95064d948ee2177c9a/shiftpath/route.go#L54-L67) 装饰器以确保 URL NotFound 返回带有尾部斜杠，因为我们的 URL 规范将它们定义为无效。 ShiftPath 方法不区分无尾随斜线和尾随斜线，我找不到一种简单的方法来做到这一点。我认为装饰器是一种合理的方法，而不是在每条路线中明确地这样做 - 在给定的网络应用程序中，您可能希望允许尾部斜杠并重定向它们，或者像我在这里所做的那样返回 Not Found .

While I like the idea of just using the standard library, and the  path-shifting technique is quite clever, I strongly prefer seeing my  URLs all in one place – Axel's approach spreads the logic across many  handlers, so it's difficult to see what handles what . It’s also quite a  lot of code, some of which is error prone.

虽然我喜欢只使用标准库的想法，而且路径转换技术非常聪明，但我非常喜欢在一个地方看到我的 URL——Axel 的方法将逻辑分布在许多处理程序中，所以很难看到什么处理什么.它的代码也相当多，其中一些很容易出错。

I do like the fact that (as Axel said) “the dependencies of [for  example] ProfileHandler are clear at compile time”, though this is true  for several of the other techniques above as well. On balance, I find it too verbose and think it’d be difficult for people reading the code to  quickly answer the question, “given this HTTP method and URL, what  happens?”

我确实喜欢这样一个事实，即（如 Axel 所说）“[例如] ProfileHandler 的依赖关系在编译时是明确的”，尽管对于上述其他几种技术也是如此。总的来说，我觉得它太冗长了，我认为阅读代码的人很难快速回答这个问题，“给定这个 HTTP 方法和 URL，会发生什么？”

[Full ShiftPath code on GitHub.](https://github.com/benhoyt/go-routing/blob/master/shiftpath/route.go)

[GitHub 上的完整 ShiftPath 代码。](https://github.com/benhoyt/go-routing/blob/master/shiftpath/route.go)

## Chi

## 志

[Chi](https://github.com/go-chi/chi) is billed as a  “lightweight, idiomatic and composable router”, and I think it lives up  to this description. It’s simple to use and the code looks nice on the  page. Here are the route definitions:

[Chi](https://github.com/go-chi/chi) 被标榜为“轻量级、惯用且可组合的路由器”，我认为它不辜负这个描述。它使用简单，代码在页面上看起来不错。以下是路由定义：

```
func init() {
    r := chi.NewRouter()

    r.Get("/", home)
    r.Get("/contact", contact)
    r.Get("/api/widgets", apiGetWidgets)
    r.Post("/api/widgets", apiCreateWidget)
    r.Post("/api/widgets/{slug}", apiUpdateWidget)
    r.Post("/api/widgets/{slug}/parts", apiCreateWidgetPart)
    r.Post("/api/widgets/{slug}/parts/{id:[0-9]+}/update", apiUpdateWidgetPart)
    r.Post("/api/widgets/{slug}/parts/{id:[0-9]+}/delete", apiDeleteWidgetPart)
    r.Get("/{slug}", widgetGet)
    r.Get("/{slug}/admin", widgetAdmin)
    r.Post("/{slug}/image", widgetImage)

    Serve = r
}
```


And the handlers are straight-forward too. They look much the same as the handlers in the regex table approach, but the custom `getField()` function is replaced by `chi.URLParam()`. One small advantage is that parameters are accessible by name instead of number:

处理程序也很直接。它们看起来与正则表达式表方法中的处理程序非常相似，但是自定义的 `getField()` 函数被 `chi.URLParam()` 替换。一个小优势是可以通过名称而不是数字访问参数：

```
func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
    slug := chi.URLParam(r, "slug")
    id, _ := strconv.Atoi(chi.URLParam(r, "id"))
    fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}
```


As with my regex table router, I’m ignoring the error value from `strconv.Atoi()` as the path parameter’s regex has already checked that it’s made of digits.

与我的正则表达式表路由器一样，我忽略了来自 `strconv.Atoi()` 的错误值，因为路径参数的正则表达式已经检查过它是由数字组成的。

If you’re going to build a substantial web app, Chi actually looks quite nice. The main `chi` package just does routing, but the module also comes with a whole bunch of [composable middleware](https://github.com/go-chi/chi#core-middlewares) to do things like HTTP authentication, logging, trailing slash handling, and so on.

如果您要构建一个强大的 Web 应用程序，Chi 实际上看起来很不错。主要的 `chi` 包只是做路由，但该模块还附带了一大堆 [composable middleware](https://github.com/go-chi/chi#core-middlewares) 来做诸如 HTTP 身份验证之类的事情，日志记录、尾部斜杠处理等。

[Full Chi code on GitHub.](https://github.com/benhoyt/go-routing/blob/master/chi/route.go)

[GitHub 上的完整 Chi 代码。](https://github.com/benhoyt/go-routing/blob/master/chi/route.go)

## Gorilla

## 大猩猩

The Gorilla toolkit is a bunch of packages that implement routing, session handling, and so on. The [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) router package is what we’ll be using here. It’s similar to Chi, though the method matching is a little more verbose:

Gorilla 工具包是一堆实现路由、会话处理等的包。我们将在这里使用 [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) 路由器包。它与 Chi 类似，但方法匹配更冗长：

```
func init() {
    r := mux.NewRouter()

    r.HandleFunc("/", home).Methods("GET")
    r.HandleFunc("/contact", contact).Methods("GET")
    r.HandleFunc("/api/widgets", apiGetWidgets).Methods("GET")
    r.HandleFunc("/api/widgets", apiCreateWidget).Methods("POST")
    r.HandleFunc("/api/widgets/{slug}", apiUpdateWidget).Methods("POST")
    r.HandleFunc("/api/widgets/{slug}/parts", apiCreateWidgetPart).Methods("POST")
    r.HandleFunc("/api/widgets/{slug}/parts/{id:[0-9]+}/update", apiUpdateWidgetPart).Methods("POST")
    r.HandleFunc("/api/widgets/{slug}/parts/{id:[0-9]+}/delete", apiDeleteWidgetPart).Methods("POST")
    r.HandleFunc("/{slug}", widgetGet).Methods("GET")
    r.HandleFunc("/{slug}/admin", widgetAdmin).Methods("GET")
    r.HandleFunc("/{slug}/image", widgetImage).Methods("POST")

    Serve = r
}
```




Again, the handlers are similar to Chi, but to get path parameters, you call `mux.Vars()`, which returns a map of all the parameters that you index by name (this  strikes me as a bit “inefficient by design” , but oh well). Here is the  code for one of the handlers:

同样，处理程序与 Chi 类似，但要获取路径参数，您调用 `mux.Vars()`，它返回您按名称索引的所有参数的映射（这让我觉得有点“设计效率低下” ，但是哦）。这是其中一个处理程序的代码：

```
func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    slug := vars["slug"]
    id, _ := strconv.Atoi(vars["id"])
    fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}
```


[Full Gorilla code on GitHub.](https://github.com/benhoyt/go-routing/blob/master/gorilla/route.go)

[GitHub 上的完整 Gorilla 代码。](https://github.com/benhoyt/go-routing/blob/master/gorilla/route.go)

## Pat

## 帕特

[Pat](https://github.com/bmizerany/pat) is interesting –  it’s a minimalist, single-file router that supports methods and path  parameters, but no regex matching. The route setup code looks similar to Chi and Gorilla:

[Pat](https://github.com/bmizerany/pat) 很有趣——它是一个极简主义的单文件路由器，支持方法和路径参数，但没有正则表达式匹配。路由设置代码看起来类似于 Chi 和 Gorilla：

```
func init() {
    r := pat.New()

    r.Get("/", http.HandlerFunc(home))
    r.Get("/contact", http.HandlerFunc(contact))
    r.Get("/api/widgets", http.HandlerFunc(apiGetWidgets))
    r.Post("/api/widgets", http.HandlerFunc(apiCreateWidget))
    r.Post("/api/widgets/:slug", http.HandlerFunc(apiUpdateWidget))
    r.Post("/api/widgets/:slug/parts", http.HandlerFunc(apiCreateWidgetPart))
    r.Post("/api/widgets/:slug/parts/:id/update", http.HandlerFunc(apiUpdateWidgetPart))
    r.Post("/api/widgets/:slug/parts/:id/delete", http.HandlerFunc(apiDeleteWidgetPart))
    r.Get("/:slug", http.HandlerFunc(widgetGet))
    r.Get("/:slug/admin", http.HandlerFunc(widgetAdmin))
    r.Post("/:slug/image", http.HandlerFunc(widgetImage))

    Serve = r
}
```


One difference is that the `Get()` and `Post()` functions take an `http.Handler` instead of an `http.HandlerFunc`, which is generally a little more awkward, as you're usually dealing with functions, not types with a `ServeHTTP` method. You can easily convert them using `http.HandlerFunc(h)`, but it’s just a bit more noisy. Here’s what a handler looks like:

一个区别是`Get()` 和`Post()` 函数采用`http.Handler` 而不是`http.HandlerFunc`，这通常有点尴尬，因为你通常处理函数，不使用 `ServeHTTP` 方法键入。您可以使用 `http.HandlerFunc(h)` 轻松转换它们，但它只是有点嘈杂。这是一个处理程序的样子：

```
func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
    slug := r.URL.Query().Get(":slug")
    id, err := strconv.Atoi(r.URL.Query().Get(":id"))
    if err != nil {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}
```


One of the interesting things is that instead of using `context` to store path parameters (and a helper function to retrieve them), Pat stuffs them into the query parameters, prefixed with `:` (colon). It’s a clever trick – if slightly dirty.

有趣的事情之一是，Pat 没有使用 `context` 来存储路径参数（和一个辅助函数来检索它们），而是将它们填充到查询参数中，并以 `:`（冒号）为前缀。这是一个聪明的技巧——如果有点脏的话。

Note that with Pat I am checking the error return value from `Atoi()`, as there’s no regex in the route definitions to ensure an ID is all  digits. Alternatively you could ignore the error, and just have the code return Not Found when it tries to look up a part with ID 0 in the  database and finds that it doesn’t exist (database IDs usually start  from 1).

请注意，使用 Pat，我正在检查来自 `Atoi()` 的错误返回值，因为路由定义中没有正则表达式来确保 ID 是所有数字。或者，您可以忽略该错误，当它尝试在数据库中查找 ID 为 0 的部件并发现它不存在（数据库 ID 通常从 1 开始）时，让代码返回 Not Found。

[Full Pat code on GitHub.](https://github.com/benhoyt/go-routing/blob/master/pat/route.go)

[GitHub 上的完整 Pat 代码。](https://github.com/benhoyt/go-routing/blob/master/pat/route.go)

## Benchmarks

## 基准

As I mentioned, I’m not concerned about speed in this comparison –  and you probably shouldn’t be either. If you're really dealing at a  scale where a few microseconds to route a URL is an issue for you, sure, use a fancy trie-based router like [httprouter](https://github.com/julienschmidt/httprouter), or write your own heavily-profiled code. All of the hand-rolled routers shown here work in linear time with respect to the number of routes  involved. 

正如我所提到的，我不关心这种比较中的速度——你可能也不应该关心。如果你真的处理的规模是几微秒来路由 URL 对你来说是一个问题，当然，使用像 [httprouter](https://github.com/julienschmidt/httprouter) 这样的基于 trie 的路由器，或编写您自己的高度剖析代码。此处显示的所有手动路由器都相对于所涉及的路由数量在线性时间内工作。

But, just to show that none of these approaches *kill* performance, below is a simple benchmark that compares routing the URL `/api/widgets/foo/parts/1/update` with each of the eight routers ([code here](https://github.com/benhoyt/go-routing/blob/9a2fa7a643ecb5681f504b95064d948ee2177c9a/main_test.go#L106)). The numbers are “nanoseconds per operation”, so lower is better. The  “operation” includes doing the routing and calling the handler. The  “noop” router is a router that actually doesn’t route anything, so  represents the overhead of the base case.

但是，只是为了表明这些方法都不会*杀死*性能，下面是一个简单的基准测试，比较路由 URL `/api/widgets/foo/parts/1/update` 与八个路由器中的每一个。数字是“每次操作的纳秒”，所以越低越好。 “操作”包括执行路由和调用处理程序。 “noop”路由器是一个实际上不路由任何东西的路由器，所以代表了基本情况的开销。

| Router    | ns/op |
| --------- |----- |
| pat       | 3646  |
| gorilla   | 2642  |
| retable   | 2014  |
| reswitch  | 1970  |
| shiftpath | 1607  |
| chi       | 1370  |
| match     | 1025  |
| split     | 984   |
| *noop*    | *583* |

|路由器 | ns/op |
| |拍 | 3646 |
|大猩猩| 2642 |
|再餐桌 | 2014 |
|重新切换 | 1970 |
|转移路径 | 1607 |
|志|第1370章
|匹配 | 1025 |
|分裂| 984 |
| *noop* | *583* |

As you can see, Pat and Gorilla are slower than the others, showing  that just because something is a well-known library doesn’t mean it’s  heavily optimized. Chi is one of the fastest, and my custom pattern  matcher and the plain `strings.Split()` method are the fastest.

如您所见，Pat 和 Gorilla 比其他库慢，这表明某些库是知名库并不意味着它经过了大量优化。 Chi 是最快的之一，我的自定义模式匹配器和简单的 `strings.Split()` 方法是最快的。

But to hammer home the point: all of these are plenty good enough –  you should almost never choose a router based on performance. The  figures here are in microseconds, so even Pat’s 3646 nanoseconds is only adding 3.6 millionths of a second to the response time. Database lookup time in a typical web app is going to be around 1000 times that.

但要强调一点：所有这些都足够好 - 您几乎永远不应该根据性能选择路由器。这里的数字以微秒为单位，因此即使是 Pat 的 3646 纳秒，响应时间也只增加了百万分之 36 秒。典型的 Web 应用程序中的数据库查找时间将是它的 1000 倍左右。

## Conclusion

##  结论

Overall this has been an interesting experiment: I came up with a  couple of new (for me, but surely not original) custom approaches to  routing, as well as trying out Axel's “ShiftPath” approach, which I'd  been intrigued about for a while.

总的来说，这是一个有趣的实验：我想出了几个新的（对我来说，但肯定不是原创的）自定义路由方法，以及尝试 Axel 的“ShiftPath”方法，我一直对它很感兴趣尽管。

If I were choosing one of the home-grown approaches, I think I would  actually end up right back where I started (when I implemented my first  server in Go a few years back) and choose the [regex table](https://benhoyt.com/writings/go-routing/#regex-table) approach. Regular expressions are quite heavy for this job, but they are well-understood and in the standard library, and the `Serve()` function is only 21 lines of code. Plus, I like the fact that the route definitions are all neatly in a table, one per line – it makes them  easy to scan and determine what URLs go where.

如果我选择一种本土方法，我想我实际上会回到我开始的地方（几年前我在 Go 中实现了我的第一台服务器时）并选择 [regex table](https://benhoyt.com/writings/go-routing/#regex-table) 方法。正则表达式对于这项工作来说相当繁重，但它们很好理解并且在标准库中，并且“Serve()”函数只有 21 行代码。另外，我喜欢这样一个事实，即路由定义都整齐地放在一个表中，每行一个 – 这使它们易于扫描并确定哪些 URL 去哪里。

A close second (still considering the home-grown approaches) would be the [regex switch](https://benhoyt.com/writings/go-routing/#regex-switch).I like the scan-style behaviour of the `match()` helper, and it also is very small (22 lines). However, the route  definitions are a little messier (two lines per route) and the handlers  that take path parameters require type or closure boilerplate – I think  that storing the path parameters using context is a bit hacky, but it  sure keeps signatures simple!

紧随其后（仍在考虑本土方法）将是 [regex switch](https://benhoyt.com/writings/go-routing/#regex-switch)。我喜欢`match()` 助手的扫描式行为，而且它也非常小（22 行）。然而，路由定义有点混乱（每条路由两行）并且接受路径参数的处理程序需要类型或闭包样板——我认为使用上下文存储路径参数有点麻烦，但它确实使签名保持简单！

For myself, I would probably rule out the other custom approaches:

对于我自己，我可能会排除其他自定义方法：

- My [split switch](https://benhoyt.com/writings/go-routing/#split-switch)approach. I like the fact that it just uses `strings.Split()`, but I find the `n == 3 && p[0] == "api" && p[1] == "widgets" && p[2] != ""` comparisons a bit ugly and error-prone.
- My [pattern matcher](https://benhoyt.com/writings/go-routing/#pattern-matcher)version. I enjoyed how simple it was to build a custom pattern matcher  for this use case (and it's only 33 lines of code), but byte-by-byte  string handling is a bit fiddly, and it doesn't gain enough over the  regexp- based approaches (which are both more powerful and in the  standard library).
- The [ShiftPath](https://benhoyt.com/writings/go-routing/#shiftpath) technique. I want to like this, but it’s just too much boilerplate for  even simple URL matching, and also, I much prefer my URL definitions in  one place.

- 我的 [split switch](https://benhoyt.com/writings/go-routing/#split-switch)方法。我喜欢它只使用`strings.Split()`，但我发现`n == 3 && p[0] == "api" && p[1] == "widgets" && p[2] != ""` 比较有点丑陋且容易出错。
- 我的 [模式匹配器](https://benhoyt.com/writings/go-routing/#pattern-matcher) 版本。我很喜欢为这个用例构建一个自定义模式匹配器是多么简单（它只有 33 行代码），但是逐字节的字符串处理有点繁琐，而且它没有获得足够的正则表达式 -基于方法（它们更强大并且在标准库中）。
- [ShiftPath](https://benhoyt.com/writings/go-routing/#shiftpath) 技术。我想喜欢这个，但即使是简单的 URL 匹配，它的样板也太多了，而且，我更喜欢我的 URL 定义在一个地方。

I disagree with Axel’s assessment that third-party routing libraries  make the routes hard to understand: all you typically have to know is  whether they match in source order, or in order of most-specific first. I also disagree that having all your routes in one place (at least for a  sub-component of your app) is a bad thing. 

我不同意 Axel 的评估，即第三方路由库使路由难以理解：您通常只需要知道它们是按源顺序匹配，还是按最具体的顺序匹配。我也不同意将所有路由放在一个地方（至少对于应用程序的一个子组件）是一件坏事。

In terms of third-party libraries, I quite like the [Chi version](https://benhoyt.com/writings/go-routing/#chi). I’d seriously consider using it, especially if building a web app as  part of a large team. Chi seems well thought out and well-tested, and I  like the composability of the middleware it provides.

第三方库方面，我比较喜欢[Chi](https://benhoyt.com/writings/go-routing/#chi)。我会认真考虑使用它，尤其是在作为大型团队的一部分构建 Web 应用程序时。 Chi 似乎经过深思熟虑和充分测试，我喜欢它提供的中间件的可组合性。

On the other hand, I'm all too aware of node-modules syndrome and the [left-pad fiasco](https://www.davidhaney.io/npm-left-pad-have-we-forgotten-how-to-program/), and agree with [Russ Cox's take](https://research.swtch.com/deps) that dependencies should be used with caution. Developers shouldn’t be  scared of a little bit of code: writing a tiny customized regex router  is fun to do, easy to understand, and easy to maintain. 

另一方面，我非常清楚节点模块综合症和 [left-pad 惨败](https://www.davidhaney.io/npm-left-pad-have-we-forgotten-how-to-program/)，并同意 [Russ Cox 的观点](https://research.swtch.com/deps) 应谨慎使用依赖项。开发人员不应该害怕一点点代码：编写一个小型定制的正则表达式路由器很有趣，易于理解且易于维护。

