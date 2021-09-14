# Go, HTTP handlers, panic, and deadlocks

February 6, 2021 (Updated: February 8, 2021)

[Programming,](http://iximiuz.com/en/categories/?category=Programming) [Go](http://iximiuz.com/en/categories/?category=Go)

Maybe the scenario I'm going to describe is just a silly bug no seasoned Go developer would ever make, but it is what it is.

I'm not an expert in Go but I do write code in this language from time to time. My cumulative number of [LOC](https://en.wikipedia.org/wiki/Source_lines_of_code) is probably still below _100 000_ but it's definitely not just a few hundred lines of code. Go always looked like a simple language to me. But also it looked safe. Apparently, it's not as simple and safe as I've thought...

Here is a synthetic piece of code illustrating the erroneous logic I stumbled upon recently:

```go
// main.go
package main

import (
    "fmt"
    "sync"
)

func main() {
    mutex := &sync.Mutex{}

    f := func() {
        fmt.Println("In f()")

        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered", r)
            }
        }()

        dogs := []string{"Lucky"}

        mutex.Lock()
        fmt.Println("Last dog's name is", dogs[len(dogs)])
        mutex.Unlock()
    }

    f()

    fmt.Println("About to get a deadlock in main()")
    mutex.Lock()
}

```

Obviously, this code has a deadlock. When we recover from panic in `f()`, the mutex is still locked. So, the second attempt to lock it in `main()` leads to a crash:

```bash
$ go run main.go
In f()
Recovered runtime error: index out of range [1] with length 1
About to get a deadlock in main()
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_SemacquireMutex(0xc0000b8014, 0x0, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc0000b8010)
    /usr/local/go/src/sync/mutex.go:138 +0x105
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
main.main()
    /home/vagrant/main.go:29 +0xe5
exit status 2

```

Of course, the real-world scenario I stumbled upon was slightly more convoluted.

When we write a server in Go, every request is usually processed in a separate goroutine. If the processing logic in a goroutine panics, the server code catches it and recovers from the panic. Such design limits the impact of individual requests on the main server process and also isolates the requests handlers from each other. The standard [`net/http`](https://golang.org/pkg/net/http/) module follows this approach and at least [some third-party libraries](https://github.com/gorilla/handlers/blob/master/recovery.go) do too:

```go
// server.go
package main

import (
    "fmt"
    "net/http"
)

func handler1(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "I'm perfectly fine!\n")
}

func handler2(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Handler is about to process a request")
    panic("Oops...")
}

func main() {
    http.HandleFunc("/1", handler1)
    http.HandleFunc("/2", handler2)
    http.ListenAndServe(":8090", nil)
}

```

Try it out yourself - start the server `go run server.go` and send some HTTP requests:

```bash
# /1 works fine
$ curl localhost:8090/1
I'm perfectly fine!

# /2 fails
$ curl localhost:8090/2
curl: (52) Empty reply from server

# ...but it doesn't affects the server's health
$ curl localhost:8090/1
I'm perfectly fine!

```

And now let's add the final bit - the mutex! It's pretty common to use a mutex to make a struct object thread-safe:

```go
// server2.go
package main

import (
    "net/http"
    "sync"
)

type MyShinyMetalThing struct {
    sync.Mutex
    comments []string
}

var thing = MyShinyMetalThing{}

func (t *MyShinyMetalThing) AddComment(text string) {
    t.Lock()
    last := t.comments[len(t.comments)]  // Oops... Happens to all of us.
    if last != text {
        t.comments = append(t.comments, text)
    }
    t.Unlock()
}

func handler(w http.ResponseWriter, req *http.Request) {
    thing.AddComment("stub")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8090", nil)
}

```

So, if you run the server with `go run server2.go` and try to `curl` it, it'll hang forever after the second attempt:

```bash
$ curl localhost:8090/
curl: (52) Empty reply from server

$ curl localhost:8090/
...

```

The issue here is pretty the same as in the synthetic snippet at the beginning of the article. It's a deadlock caused by the second request trying to lock the already locked `thing` object. Despite the panic that was recovered by the `net/http` module logic, the actual state of the program (the `thing` object) is broken.

Of course At first sight, using `defer` might help.

_**NB: actually, it's more complicated than that. [See my update below](http://iximiuz.com#update).**_

```go
func (t *MyShinyMetalThing) AddComment(text string) {
    t.Lock()
    defer t.Unlock()

    last := t.comments[len(t.comments)]  // Oops... Happens to all of us.
    if last != text {
        t.comments = append(t.comments, text)
    }
}

```

But from what I've seen in the wild, `defer` is often omitted if there is no conditional `return` in between the `Lock()` and `Unlock()` lines. Also, `defer` works at a function call level. So the part of code that needs to be synchronized, should be extracted in a separate function first. This creates another mental barrier and plays in favour of the error-prone pattern from above.

So, from now on, I tend to see such server-side code as an anti-pattern:

```go
func doSomeStuff() {
    // do A
    // do B

    // Need thread-safety only for this part
    // mut.Lock()
    // do C which relies on some local vars from above
    // mut.Unlock()

    // do D
    // do E
}

```

#### Added on Feb 8th, 2021 after a fruitful discussion on Reddit

Sharing this post on Reddit raised some [sort of polemic](https://www.reddit.com/r/golang/comments/le3y00/go_http_handlers_panic_and_deadlocks/). While some users were agreeing that using `mut.Lock()` without a deferred `mut.Unlock()` **is** an anti-pattern other Redditors were referring to various downsides of this approach. But [the most interesting piece of feedback](https://www.reddit.com/r/golang/comments/le3y00/go_http_handlers_panic_and_deadlocks/gmilfs8/?context=3) was from [u/jeminstall](https://www.reddit.com/user/jeminstall/).

Basically, the main problem is that even a successfully released mutex doesn't guarantee that the actual state of the program is consistent after recovering from a `panic()` call. Most likely, the mutex was protecting a critical section of the code. It's also highly likely, that this section was meant to make an atomic change to the program's state. But then something happened and the program `panic()`-ed in the middle of the critical section. Thanks to the `defer mut.Unlock()` construct, we released the mutex. However, it's very likely, that we also had to revert the half-through made change to the program's state while recovering. And that's in general would be pretty hard to achieve, I believe.

This person even kindly provided a [snippet illustrating the problem](https://play.golang.org/p/_oZXX0hCl8h):

```go
package main

import (
    "fmt"
    "sync"
)

type ByteCache struct {
    maxSize int

    lock     sync.Mutex
    cache    map[string][]byte
    currSize int
}

func NewByteCache(maxSize int) *ByteCache {
    return &ByteCache{
        cache:   make(map[string][]byte),
        maxSize: maxSize,
    }
}

func (c *ByteCache) Put(key string, data []byte) bool {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Put() recovered:", r)
        }
    }()

    c.lock.Lock()
    defer c.lock.Unlock()

    // Don't allow overwrites to keep example short.
    _, ok := c.cache[key]
    if ok {
        return false
    }

    // Don't allow our cache to exceed the specified maxSize.
    if c.currSize+len(data) > c.maxSize {
        return false
    }

    c.currSize += len(data)
    if key == "burn" {
        panic("some programmers just want to watch the world burn")
    }
    c.cache[key] = data
    return true
}

func (c *ByteCache) Get(key string) ([]byte, bool) {
    c.lock.Lock()
    defer c.lock.Unlock()

    data, ok := c.cache[key]
    return data, ok
}

func (c *ByteCache) Print() {
    c.lock.Lock()
    defer c.lock.Unlock()

    fmt.Printf("MaxSize: %d\n", c.maxSize)
    fmt.Printf("Cache: ")
    acutalSize := 0
    for k, v := range c.cache {
        fmt.Printf("%q:%d ", k, len(v))
        acutalSize += len(v)
    }
    fmt.Printf("\nCurrSize: %d\n", c.currSize)
    fmt.Printf("ActualSize: %d\n", acutalSize)
}

func main() {
    bc := NewByteCache(5)

    bc.Put("watch", []byte{0})
    bc.Put("it", []byte{0})
    bc.Put("burn", []byte{0, 0, 0})

    bc.Print()

    /*
        Put() recovered: some programmers just want to watch the world burn
        MaxSize: 5
        Cache: "watch":1 "it":1
        CurrSize: 5
        ActualSize: 2
    */
}

```

So, after this discussion, I'm actually inclined to see the **recovery from panic() as an actual anti-pattern** here. There might be good reasons for unlocking mutex without deferring the call. There might be situations when `defer mut.Unlock()` may be a safer way to go. But the general rule of thumb should be to mindfully consider every such case as a unique situation with its own set of trade-offs. And it's important to remember, that fully recovering from panic is very hard in general.

Apparently, there was even a [proposal to change (or make configurable) the behavior of `net/http` module](https://github.com/golang/go/issues/25245) with regard to its unconditional recovery. But it was rejected due to backward compatibility concerns. Hopefully, [it'll make it to Go 2](https://github.com/golang/go/issues/5465) though.

Until then, we can use the following middleware that intercepts `panic()` calls and terminates the server deliberately:

```go
// server3.go
package main

import (
    "log"
    "net/http"
    "os"
    "sync"
)

type MyShinyMetalThing struct {
    sync.Mutex
    comments []string
}

func (t *MyShinyMetalThing) AddComment(text string) {
    t.Lock()
    last := t.comments[len(t.comments)]
    if last != text {
        t.comments = append(t.comments, text)
    }
    t.Unlock()
}

var thing = MyShinyMetalThing{}

func handler(w http.ResponseWriter, req *http.Request) {
    thing.AddComment("stub")
}

func withPanic(fn func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
    return func(w http.ResponseWriter, req *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Println("Panic intercepted!", err)
                os.Exit(1)
            }
        }()
        fn(w, req)
    }
}

func main() {
    http.HandleFunc("/", withPanic(handler))
    http.ListenAndServe(":8090", nil)
}

```

The rest of the system (orchestrator, monitoring, and alerting) should take care of the failing servers then. But again, not everything is black and white. Maybe certain servers need to recover from panic programmatically. Here is a [good post on compromises of using panic/recover](https://eli.thegreenplace.net/2018/on-the-uses-and-misuses-of-panics-in-go/) from Eli Bendersky. It's definitely worth checking out.

Stay safe and don't panic!

[golang,](javascript: void 0) [panic,](javascript: void 0) [recover,](javascript: void 0) [deadlock,](javascript: void 0) [mutex](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

