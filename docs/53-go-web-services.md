# A Gentle Introduction to Web Services With Go
Sep 8, 2020 From: https://www.honeybadger.io/blog/go-web-services/

When you're deciding on a technology to use  for your project, it helps to have a broad understanding of your  options. You may be tempted to build a web service in Go for performance reasons - but what would that code actually look like? How would it  compare to languages like Ruby or JS? In this article, Ayooluwa Isaiah  gives us a guided tour through the building blocks of go web services so you'll be well-informed.

You may have heard that Go is great for web development. This claim is evidenced by several cloud services, such as Cloudflare, Digital Ocean, and others, betting on Go for their cloud infrastructure. But, what makes Go such a compelling option for building network-facing applications?

One reason is the built-in concurrency in the language through goroutines and channels, which makes programs effectively use the resources of the hardware its running on. Ruby and other web-centered languages are usually single-threaded. Going beyond that can be done with the help of threads, but these feel like an add-on rather than a first-class citizen of the language.

Go also boasts a great standard library that includes many of the basic utilities necessary to build a robust web service, including an HTTP server, a templating library, and JSON utilities, as well as interfacing with databases, such as SQL. Support is excellent for all the latest technologies, ranging from HTTP/2 to databases, such as Postgres, MongoDB and ElasticSearch, and the latest HTTPS encryption standards (including TLS 1.3).

In terms of performance, Go is far ahead when compared to Ruby and other scripting languages. For example, Hugo, a static site generator built with Go, generates sites an average of [35 times faster](https://forestry.io/blog/hugo-and-jekyll-compared/#performance-1) than Ruby-based Jekyll. Additionally, since all Go programs compile to a single binary, it's easy to deploy to any cloud service of your choice. Applications built with Go can run natively in Google's App Engine and Google Cloud Run or numerous other environments, cloud services, or operating systems thanks to Go’s extreme portability.

All this put together makes Go one of the best languages for the development of a new web service or application. This article will describe some specific aspects of building a web server in Go to give you an idea of how it feels to build with Go.

Note that Go is not exactly a drop-in replacement for monolithic frameworks, such as Rails. It's mostly used to build microservices, and many programmers have had good experiences factoring subsets of Rails apps into several HTTP or JSON services and front-ending them with Rails or other frameworks to get the best of both worlds!

## A Basic HTTP server

HTTP servers are easy to write in Go using the `net/http` package. Unlike other languages with half-baked web servers that are tedious to work with and not ideal for use in anything more than a toy app, Go's `net/http` package is meant for production use and is, indeed, the most popular way to develop a web service in the language, although other solutions do exist.

An example of the most basic HTTP server you can build in the language is presented below. It runs on port 8080 and responds with the HTML string `<h1>Welcome to my web server!</h1>` when a request is made to the server root.

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

![Welcome to my web server!](https://www.honeybadger.io/images/blog/posts/go-web-services/1.png?1630027696)

Let's break down the code a bit to fully understand what's going on. The `net/http` package is the foundation of all web servers in Go. It enables the creation of applications capable of making requests to other servers, as well as responding to incoming web requests.

The `main` function begins with a call to the `http.HandleFunc` method, which provides a way to specify how requests to a specific route should be handled. The first argument is the route in question ("/" in this case), while the second is a function of the type [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc).

Any function with the signature `func(http.ResponseWriter, *http.Request)` can be passed to any other function that expects the `http.HandlerFunc` type. In this case, the function is defined inline, but you can also make it a standalone function, as shown below:

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

Inside our handler, we have a single line of code:

```
w.Write([]byte("<h1>Welcome to my web server!</h1>"))
```

The `http.ResponseWriter` interface has a `Write` method which accepts a byte slice and writes the data to the connection as part of an HTTP response. Converting a string to a byte slice is as easy as using `[]byte(str)`, and that's how we're able to respond to HTTP requests.

In the final line of the `main` function, we have the `http.ListenAndServe` method wrapped in a call to `log.Fatal`. `ListenAndServe` specifies the port to listen on as the first argument and an [http.Handler](https://golang.org/pkg/net/http/#Handler) as its second argument. If the handler is nil, `DefaultServeMux` is used.

While `DefaultServeMux` is okay for toy programs, you should avoid it in your production code. This is because it is accessible to any package that your application imports, including third-party packages. Therefore, it could potentially be exploited to expose a malicious handler to the web if a package becomes compromised.

So, what's the alternative? A locally scoped [http.ServeMux](https://golang.org/pkg/net/http/#ServeMux)! Here's how it works:

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

But, creating your own ServeMux is easy using the `http.NewServeMux` method and registering your handlers on the resulting ServeMux. Also, instead of `nil`, the custom ServeMux should be used as the second argument to `ListenAndServe`, as shown above.

The reason `http.ListenAndServe` is wrapped inside a call to `log.Fatal` is it will always return an error if something unexpected happens, and you definitely want to log the error. Otherwise, it will block until the program is terminated.

## Routing

Routing is basically the process of mapping an inbound request to the appropriate HTTP handler. As discussed above, the `http.ServeMux` type provides this functionality in our program by virtue of its `HandleFunc` method. The process of adding a new route to the server is as simple as registering a new handler, as shown below:

```
. . .
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

Let's consider some other routing features, as well as shortcomings, of the `http.ServeMux` in more detail.

### Pattern matching

The ServeMux handler follows a pretty basic model for routing requests, but it can be very confusing. For each route, you need to be explicit about the registered paths since it does not support wildcards or regular expressions.

Two kinds of path patterns are supported by Go's ServeMux: fixed paths and subtree paths. The former does not end with a trailing slash, while the latter does.

We've already seen both types of paths in our demo application. The root pattern ("/")is a subtree path, while "/about" is a fixed path. The difference is that the handler for a fixed pattern is only expected when the request URL path is an *exact* match for the pattern. Subtree path handlers, however, will execute whenever the start of the request URL matches the subtree pattern.

Here's an example that will make it clearer for you to understand. Let's say we have the following registered paths:

```
mux.HandleFunc("/", indexHandler)
mux.HandleFunc("/user", userHandler)
mux.HandleFunc("/user/create", createUserHandler)
```

Here's how ServeMux will route the following incoming requests:

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

Here are some other details about how Go's ServeMux routes incoming requests:

- Route patterns can be registered in any order. It does not affect the behavior of the router.
- Incoming requests to a subtree path that do not have a trailing slash will be 301 redirected to subtree path with the slash added. So, if `/image/` is registered, and a request is made to `/image`, it will be redirected to `/image/` as long as `/image` itself is not registered.
- It's possible to specify the hostname in the route pattern. We could register a path, such as `auth.example.com`, and requests to that URL will be directed to the registered handler.

### 404 Not Found

As mentioned above, patterns that end with a trailing slash act like a catch-all for URL paths that match the start of the pattern. This includes the root route, which matches any URL not handled by a more specific route. If you want the root pattern to stop acting as a catch-all path, you can check whether the current request URL exactly matches `/`. If it doesn't, use the `http.NotFound()` method to send a 404 not found response.

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

![404 error page](https://www.honeybadger.io/images/blog/posts/go-web-services/3.png?1630027696)

### Redirects

You can redirect from one URL to another using the `http.Redirect()` method.

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

To serve static files from a web server, the `http.FileServer()` method can be utilized.

```
func main() {
    staticHandler := http.FileServer(http.Dir("./assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))
}
```

The `FileServer()` function returns an `http.Handler` that responds to all HTTP requests with the contents of the provided filesystem. In the above example, the filesystem is given as the `assets` directory relative to the application.

Next, the `Handle` method is used to register the FileServer Handler for all request URLs that begin with `/assets/`. The trailing slash is significant for reasons we already discussed above. The `http.StripPrefix()` method is used to strip out the `/assets/` prefix from the request path before searching the filesystem for the requested file.

From now on, any requests for static files in the `assets` directory, such as `/assets/css/styles.css` or `/assets/images/sample.jpg`, will be handled appropriately.

### HTTP request methods

Thus far, we haven't given much consideration to the request methods allowed for each route, so let's do that here. Go's ServeMux does not have any special way to specify the allowed methods for a route, so you have to check the request method yourself in the HTTP handler.

Let's make sure that the `indexHandler` function returns an error if a non-GET request is made to the root route.

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

### Reading query parameters

It's often necessary to read the query parameters sent as part of an HTTP request. For example, given the following request url:

```
/user?id=123
```

Here's how to access the value of the `id` query parameter in Go:

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

```
values := r.URL.Query()
for k, v := range values {
    fmt.Println(k, " => ", v)
}
```

## Why you should care about ServeMux

If you're coming to Go with a background in using frameworks like Ruby on Rails, the capabilities of the built-in router in Go may seem incredibly simple or even underwhelming.

A quick Google search will locate some great third-party options that you can plug into your application, but I recommend that you learn the limitations of ServeMux first before opting for one of them. Besides, many projects use the built-in router, so you still need to know how it works to be able to contribute effectively.

Go developers tend to rely on the standard library in most cases, but there are times when it is necessary to go beyond what it provides. In a future article, we'll examine some of the popular third-party routers to see what else is out there.

## Wrap up

In this article, we discussed why Go is great for web development and proceeded to build an HTTP server in the language while examining the different features of the built-in router that you need to know about.

If you have any questions or opinions on what we've covered here, I'd love to hear about it on [Twitter](https://twitter.com/ayisaiah). In the next article, we'll cover some other aspects of building web applications with Go, such as middleware and JSON, as well as templating and working with databases.

Thanks for reading, and happy coding!
