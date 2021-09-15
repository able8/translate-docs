# Go, HTTP handlers, panic, and deadlocks

# Go、HTTP 处理程序、恐慌和死锁

February 6, 2021 (Updated: February 8, 2021)

2021 年 2 月 6 日（更新：2021 年 2 月 8 日）

[Programming,](http://iximiuz.com/en/categories/?category=Programming)[Go](http://iximiuz.com/en/categories/?category=Go)

[编程](http://iximiuz.com/en/categories/?category=Programming)[Go](http://iximiuz.com/en/categories/?category=Go)

Maybe the scenario I'm going to describe is just a silly bug no seasoned Go developer would ever make, but it is what it is.

也许我要描述的场景只是一个没有经验的 Go 开发人员会做的愚蠢的错误，但事实就是这样。

I'm not an expert in Go but I do write code in this language from time to time. My cumulative number of [LOC](https://en.wikipedia.org/wiki/Source_lines_of_code) is probably still below _100 000_ but it's definitely not just a few hundred lines of code. Go always looked like a simple language to me. But also it looked safe. Apparently, it's not as simple and safe as I've thought...

我不是 Go 专家，但我确实不时用这种语言编写代码。我的 [LOC](https://en.wikipedia.org/wiki/Source_lines_of_code) 的累积数量可能仍低于 _100 000_，但绝对不仅仅是几百行代码。 Go 对我来说总是看起来像一种简单的语言。但它看起来也很安全。显然，这并不像我想象的那么简单和安全......

Here is a synthetic piece of code illustrating the erroneous logic I stumbled upon recently:

这是一段合成代码，说明了我最近偶然发现的错误逻辑：

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
            if r := recover();r != nil {
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

显然，这段代码存在死锁。当我们从 `f()` 中的 panic 中恢复时，互斥锁仍然被锁定。因此，第二次尝试将其锁定在 `main()` 中会导致崩溃：

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

当然，我偶然发现的现实场景稍微复杂一些。

When we write a server in Go, every request is usually processed in a separate goroutine. If the processing logic in a goroutine panics, the server code catches it and recovers from the panic. Such design limits the impact of individual requests on the main server process and also isolates the requests handlers from each other. The standard [`net/http`](https://golang.org/pkg/net/http/) module follows this approach and at least [some third-party libraries](https://github.com/gorilla/handlers/blob/master/recovery.go) do too:

当我们用 Go 编写服务器时，每个请求通常都在一个单独的 goroutine 中处理。如果 goroutine 中的处理逻辑发生混乱，服务器代码会捕获它并从混乱中恢复。这种设计限制了单个请求对主服务器进程的影响，并将请求处理程序彼此隔离。标准的 [`net/http`](https://golang.org/pkg/net/http/) 模块遵循这种方法，至少 [一些第三方库](https://github.com/gorilla/handlers/blob/master/recovery.go) 也这样做：

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

自己尝试一下 - 启动服务器 `go run server.go` 并发送一些 HTTP 请求：

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

现在让我们添加最后一点 - 互斥锁！使用互斥量使结构对象线程安全是很常见的：

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

所以，如果你用 `go run server2.go` 运行服务器并尝试 `curl` 它，它会在第二次尝试后永远挂起：

```bash
$ curl localhost:8090/
curl: (52) Empty reply from server

$ curl localhost:8090/
...

```

The issue here is pretty the same as in the synthetic snippet at the beginning of the article. It's a deadlock caused by the second request trying to lock the already locked `thing` object. Despite the panic that was recovered by the `net/http` module logic, the actual state of the program (the `thing` object) is broken.

这里的问题与文章开头的合成片段中的问题非常相似。这是由第二个请求试图锁定已经锁定的“事物”对象引起的死锁。尽管通过`net/http` 模块逻辑恢复了恐慌，程序的实际状态（`thing` 对象）被破坏了。

Of course At first sight, using `defer` might help.

当然，乍一看，使用 `defer` 可能会有所帮助。

_**NB: actually, it's more complicated than that. [See my update below](http://iximiuz.com#update).**_

_**注意：实际上，它比那更复杂。 [见下面我的更新](http://iximiuz.com#update).**_

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

但从我在野外看到的情况来看，如果在 `Lock()` 和 `Unlock()` 行之间没有有条件的 `return`，`defer` 通常会被省略。此外，`defer` 在函数调用级别工作。所以需要同步的那部分代码，应该先在一个单独的函数中提取出来。这创造了另一个心理障碍，并有利于从上方产生的容易出错的模式。

So, from now on, I tend to see such server-side code as an anti-pattern:

所以，从现在开始，我倾向于将这样的服务器端代码视为一种反模式：

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

#### 在 Reddit 上进行了富有成效的讨论后，于 2021 年 2 月 8 日添加

Sharing this post on Reddit raised some [sort of polemic](https://www.reddit.com/r/golang/comments/le3y00/go_http_handlers_panic_and_deadlocks/). While some users were agreeing that using `mut.Lock()` without a deferred `mut.Unlock()` **is** an anti-pattern other Redditors were referring to various downsides of this approach. But [the most interesting piece of feedback](https://www.reddit.com/r/golang/comments/le3y00/go_http_handlers_panic_and_deadlocks/gmilfs8/?context=3) was from [u/jeminstall](https://www.reddit.com/user/jeminstall/).

在 Reddit 上分享这篇文章引发了一些[某种争论](https://www.reddit.com/r/golang/comments/le3y00/go_http_handlers_panic_and_deadlocks/)。虽然一些用户同意使用没有延迟的`mut.Unlock()` 的 `mut.Lock()` **是一种反模式**，但其他 Redditor 指的是这种方法的各种缺点。但是 [最有趣的反馈](https://www.reddit.com/r/golang/comments/le3y00/go_http_handlers_panic_and_deadlocks/gmilfs8/?context=3) 来自 [u/jeminstall](https://www.reddit.com/user/jeminstall/)。

Basically, the main problem is that even a successfully released mutex doesn't guarantee that the actual state of the program is consistent after recovering from a `panic()` call. Most likely, the mutex was protecting a critical section of the code. It's also highly likely, that this section was meant to make an atomic change to the program's state. But then something happened and the program `panic()`-ed in the middle of the critical section. Thanks to the `defer mut.Unlock()` construct, we released the mutex. However, it's very likely, that we also had to revert the half-through made change to the program's state while recovering. And that's in general would be pretty hard to achieve, I believe.

基本上，主要问题是即使成功释放互斥锁也不能保证程序的实际状态在从 `panic()` 调用中恢复后是一致的。最有可能的是，互斥锁保护了代码的关键部分。这部分也很有可能是为了对程序的状态进行原子更改。但是后来发生了一些事情，程序在临界区中间发生了`panic()`。感谢 `defer mut.Unlock()` 构造，我们释放了互斥锁。但是，很可能我们还必须在恢复时将中途更改恢复到程序状态。我相信，这通常很难实现。

This person even kindly provided a [snippet illustrating the problem](https://play.golang.org/p/_oZXX0hCl8h):

这个人甚至好心地提供了一个[说明问题的片段](https://play.golang.org/p/_oZXX0hCl8h)：

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
        if r := recover();r != nil {
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

所以，在这次讨论之后，我实际上倾向于将 **panic() 的恢复视为一种实际的反模式**。在不推迟调用的情况下解锁互斥锁可能有很好的理由。在某些情况下，`defer mut.Unlock()` 可能是一种更安全的方法。但是一般的经验法则应该是谨慎地将每个此类情况视为具有其自身权衡的独特情况。重要的是要记住，从恐慌中完全恢复通常是非常困难的。

Apparently, there was even a [proposal to change (or make configurable) the behavior of `net/http` module](https://github.com/golang/go/issues/25245) with regard to its unconditional recovery. But it was rejected due to backward compatibility concerns. Hopefully, [it'll make it to Go 2](https://github.com/golang/go/issues/5465) though.

显然，甚至有 [提议改变（或使可配置）`net/http` 模块的行为](https://github.com/golang/go/issues/25245) 关于其无条件恢复。但由于向后兼容性问题，它被拒绝了。但希望[它会进入 Go 2](https://github.com/golang/go/issues/5465)。

Until then, we can use the following middleware that intercepts `panic()` calls and terminates the server deliberately:

在此之前，我们可以使用以下中间件拦截 `panic()` 调用并故意终止服务器：

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
            if err := recover();err != nil {
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

系统的其余部分（编排器、监控和警报）应该负责处理出现故障的服务器。但同样，并非一切都是非黑即白的。也许某些服务器需要以编程方式从恐慌中恢复。这是 Eli Bendersky 的 [关于使用 panic/recover 的妥协的好帖子](https://eli.thegreenplace.net/2018/on-the-uses-and-misuses-of-panics-in-go/)。绝对值得一试。

Stay safe and don't panic!

保持安全，不要惊慌！

[golang,](javascript: void 0) [panic,](javascript: void 0) [recover,](javascript: void 0) [deadlock,](javascript: void 0) [mutex](javascript: void 0)

[golang,](javascript: void 0) [panic,](javascript: void 0) [recover,](javascript: void 0) [deadlock,](javascript: void 0) [mutex](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

