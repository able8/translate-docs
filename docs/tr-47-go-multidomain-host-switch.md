# Handling multidomain HTTP requests with simple host switch

# 使用简单的主机切换处理多域 HTTP 请求

Written on June  5, 2020   From: https://rafallorenz.com/go/go-multidomain-host-switch/

When writing [Go](http://golang.org) HTTP services we  might need to perform different logic based on host the request comes  from or simply validate host value. We could implement simple host  switch mechanism to take care of that.

在编写 [Go](http://golang.org) HTTP 服务时，我们可能需要根据请求来自的主机执行不同的逻辑或简单地验证主机值。我们可以实现简单的主机切换机制来解决这个问题。

# Host Switch

# 主机交换机

As difficult as host switch may sound, it is actually quite easy to implement. We need to implement the [http.Handler](https://golang.org/pkg/net/http/#Handler) interface. To do that we will create a type for which we implement the `ServeHTTP` method. We will just use a map here, in which we map host names (with port) to http.Handlers.

尽管主机切换听起来很困难，但实际上很容易实现。我们需要实现 [http.Handler](https://golang.org/pkg/net/http/#Handler) 接口。为此，我们将创建一个类型，为其实现“ServeHTTP”方法。我们将在这里只使用一个映射，在其中我们将主机名（带端口）映射到 http.Handlers。

```
import (
    "fmt"
    "log"
    "net/http"
)

type HostSwitch map[string]http.Handler

// Implement the ServerHTTP method
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if handler, ok := hs[r.Host];ok && handler != nil {
        handler.ServeHTTP(w, r)
    } else {
        http.Error(w, "Forbidden", http.StatusForbidden)
    }
}
```


Our `HostSwitch` simply checks if a [http.Handler](https://golang.org/pkg/net/http/#Handler) is registered for the given host, then uses it to handle request. Otherwise we return http error `http.StatusForbidden`. We could do a redirect here to another host, depends what we really need in real life example.

我们的 `HostSwitch` 只是检查是否为给定主机注册了 [http.Handler](https://golang.org/pkg/net/http/#Handler)，然后使用它来处理请求。否则我们将返回 http 错误 `http.StatusForbidden`。我们可以在此处重定向到另一个主机，这取决于我们在现实生活中真正需要的示例。

# Usage

#  用法

Let’s say our switch will handle two hosts:

假设我们的交换机将处理两个主机：

- `example-one.local`
- `example-two.local`

And return forbidden response otherwise. We will use [ServeMux](https://golang.org/pkg/net/http/#ServeMux) as a router for both hosts.

否则返回禁止响应。我们将使用 [ServeMux](https://golang.org/pkg/net/http/#ServeMux) 作为两个主机的路由器。

```
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Welcome to the first home page!")
    })

    muxTwo := http.NewServeMux()
    muxTwo.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Welcome to the second home page!")
    })

    hs := make(HostSwitch)
    hs["example-one.local:8080"] = mux
    hs["example-two.local:8080"] = muxTwo

    log.Fatal(http.ListenAndServe(":8080", hs))
}
```


Example is very straightforward, after creating two routers, we make a new `HostSwitch` and insert each router for given host we want it to handle. Then we listen and serve on both ports. We could use same router if what we want is to actually validate host value.

示例非常简单，在创建两个路由器之后，我们创建一个新的 `HostSwitch` 并为我们希望它处理的给定主机插入每个路由器。然后我们在两个端口上监听和服务。如果我们想要实际验证主机值，我们可以使用相同的路由器。

```
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Welcome to the home page!")
    })

    hs := make(HostSwitch)
    hs["example-one.local:8080"] = mux
    hs["example-two.local:8080"] = mux

    go log.Fatal(http.ListenAndServe(":8080", hs))
}
```


Even though we perform only read operation, it is always better to use mutexes while working with maps. Don’t forget to use [mutexes](https://golang.org/pkg/sync/#Mutex) to make our `HostSwitch` concurrency safe.

即使我们只执行读取操作，在处理映射时使用互斥锁总是更好。不要忘记使用 [mutexes](https://golang.org/pkg/sync/#Mutex) 来确保我们的 `HostSwitch` 并发安全。

# Conclusion

#  结论

Today we have created a simple `HostSwitch` which allows us to handle requests only if they come from valid hosts  or perform different logic per host. This example demonstrates how to  handle multidomains with Go. Example on [The Go Playground](https://play.golang.org/p/bMbKPGE7LhT) 

今天我们创建了一个简单的`HostSwitch`，它允许我们仅处理来自有效主机的请求或为每个主机执行不同的逻辑。此示例演示如何使用 Go 处理多域。 
