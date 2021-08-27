# Handling multidomain HTTP requests with simple host switch

​    Written on June  5, 2020   From: https://rafallorenz.com/go/go-multidomain-host-switch/

When writing [Go](http://golang.org) HTTP services we  might need to perform different logic based on host the request comes  from or simply validate host value. We could implement simple host  switch mechanism to take care of that.

# Host Switch

As difficult as host switch may sound, it is actually quite easy to implement. We need to implement the [http.Handler](https://golang.org/pkg/net/http/#Handler) interface. To do that we will create a type for which we implement the `ServeHTTP` method. We will just use a map here, in which we map host names (with port) to http.Handlers.

```
import (
    "fmt"
    "log"
    "net/http"
)

type HostSwitch map[string]http.Handler

// Implement the ServerHTTP method
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if handler, ok := hs[r.Host]; ok && handler != nil {
        handler.ServeHTTP(w, r)
    } else {
        http.Error(w, "Forbidden", http.StatusForbidden)
    }
}
```

Our `HostSwitch` simply checks if a [http.Handler](https://golang.org/pkg/net/http/#Handler) is registered for the given host, then uses it to handle request. Otherwise we return http error `http.StatusForbidden`. We could do a redirect here to another host, depends what we really need in real life example.

# Usage

Let’s say our switch will handle two hosts:

- `example-one.local`
- `example-two.local`

And return forbidden response otherwise. We will use [ServeMux](https://golang.org/pkg/net/http/#ServeMux) as a router for both hosts.

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

# Conclusion

Today we have created a simple `HostSwitch` which allows us to handle requests only if they come from valid hosts  or perform different logic per host. This example demonstrates how to  handle multidomains with Go. Example on [The Go Playground](https://play.golang.org/p/bMbKPGE7LhT)


