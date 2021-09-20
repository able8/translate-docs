# Simple techniques to optimise Go programs

# 优化 Go 程序的简单技巧

I'm very interested in performance. I'm not sure I can explain the underlying reasons for it. I'm easily frustrated at slow services and programs, and it seems like [I'm not alone.](http://glinden.blogspot.com/2006/11/marissa-mayer-at-web-20.html)

我对表演很感兴趣。我不确定我能解释它的根本原因。我很容易对缓慢的服务和程序感到沮丧，而且似乎 [我并不孤单。](http://glinden.blogspot.com/2006/11/marissa-mayer-at-web-20.html)

> _In A/B tests, we tried delaying the page in increments of 100 milliseconds and found that even very small delays would result in substantial and costly drops in revenue._ \- Greg Linden, Amazon.com

> _在 A/B 测试中，我们尝试以 100 毫秒为增量延迟页面，发现即使是非常小的延迟也会导致收入大幅下降且代价高昂。_ \- Greg Linden，Amazon.com

From my experience, poor performance manifests in one of two ways:

根据我的经验，性能不佳表现为以下两种方式之一：

- Operations that performed well at small scale, but become unviable as the number of users grows. These are usually O(N) or O(N²) operations. When your user base is small, these perform just fine, and are often done in order to get a product to market. As your use base grows, you see more [pathological examples](https://theoutline.com/post/4147/in-twitters-early-days-only-one-celebrity-could-tweet-at-a-time?zd=1&zi=ivqvd4py) that you weren't expecting, and your service grinds to a halt.
- Many individual sources of small optimisation - AKA 'death by a thousand crufts'.

- 在小范围内表现良好的操作，但随着用户数量的增长变得不可行。这些通常是 O(N) 或 O(N²) 操作。当您的用户群很小时，这些表现很好，并且通常是为了将产品推向市场。随着您使用基础的增长，您会看到更多[病理示例](https://theoutline.com/post/4147/in-twitters-early-days-only-one-celebrity-could-tweet-at-a-time?zd=1&zi=ivqvd4py)，这是您没想到的，您的服务会停止。
- 小优化的许多个人来源 - 又名“千篇一律的死亡”。

I've spent most of my career either doing data science with Python, or building services in Go; I have far more experience optimising the latter. Go is usually not the bottleneck in the services I write - programs are often IO bound as they talk to the database. However, in batch machine learning pipelines - like I built in my previous role - your program is often CPU bound. When your Go program is using excessive CPU, and the excessive usage is having a negative impact, there's various strategies you can use to mitigate that.

我职业生涯的大部分时间要么用 Python 进行数据科学，要么用 Go 构建服务；我有更多优化后者的经验。 Go 通常不是我编写的服务的瓶颈——程序在与数据库交谈时通常是 IO 绑定的。但是，在批处理机器学习管道中 - 就像我在之前的角色中构建的那样 - 您的程序通常受 CPU 限制。当您的 Go 程序使用过多的 CPU，并且过度使用会产生负面影响时，您可以使用多种策略来减轻这种影响。

This post explains some techniques you can use to significantly improve the performance of your program with little effort. I am deliberately ignoring techniques that require significant effort, or large changes to program structure.

这篇文章解释了一些技术，您可以使用这些技术轻松显着提高程序的性能。我故意忽略需要大量努力或对程序结构进行大量更改的技术。

## Before you start

##  在你开始之前

Before you make any changes to your program, invest time in creating a proper baseline to compare against. If you don't do this, you'll be searching around in the dark, wondering if your changes are having any improvement. Write benchmarks first, and grab [profiles](https://blog.golang.org/profiling-go-programs) for use in pprof. In the best case, this will be a [Go benchmark](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go): this allows easy usage of pprof, and memory allocation profiling. You should also use [`benchcmp`](https://godoc.org/golang.org/x/tools/cmd/benchcmp): a helpful tool for comparing the difference in performance between two benchmarks.

在对程序进行任何更改之前，请花时间创建适当的基线以进行比较。如果您不这样做，您将在黑暗中四处寻找，想知道您的更改是否有任何改进。先写benchmarks，然后抓取[profiles](https://blog.golang.org/profiling-go-programs) 用于pprof。在最好的情况下，这将是一个 [Go 基准测试](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)：这样可以轻松使用 pprof，和内存分配分析。您还应该使用 [`benchcmp`](https://godoc.org/golang.org/x/tools/cmd/benchcmp)：一个有用的工具，用于比较两个基准测试之间的性能差异。

If your code is not easily benchmarked, just start with something that you can time. You can profile your code manually using [`runtime/pprof`](https://golang.org/pkg/runtime/pprof/).

如果您的代码不容易进行基准测试，请从可以计时的内容开始。您可以使用 [`runtime/pprof`](https://golang.org/pkg/runtime/pprof/) 手动分析您的代码。

Let's get started!

让我们开始吧！

### Use `sync.Pool` to re-use previously allocated objects

### 使用`sync.Pool` 来重用之前分配的对象

[`sync.Pool`](https://golang.org/pkg/sync/#Pool) implements a [free-list](https://en.wikipedia.org/wiki/Free_list). This allows you to re-use structures that you've previously allocated. This amortises the allocation of an object over many usages, reducing the work the garbage collector has to do. The API is very simple: implement a function that allocates a new instance of your object. It should return a pointer type.

[`sync.Pool`](https://golang.org/pkg/sync/#Pool) 实现了一个 [free-list](https://en.wikipedia.org/wiki/Free_list)。这允许您重用您之前分配的结构。这可以在多次使用中分摊对象的分配，减少垃圾收集器必须做的工作。 API 非常简单：实现一个函数来分配对象的新实例。它应该返回一个指针类型。

```
var bufpool = sync.Pool{
    New: func() interface{} {
        buf := make([]byte, 512)
        return &buf
    }}

```

After this, you can `Get()` objects from the pool, `Put()` ting them back after you are done.

在此之后，您可以从池中`Get()` 对象，完成后`Put()` 将它们放回原处。

```
// sync.Pool returns a interface{}: you must cast it to the underlying type
// before you use it.
bp := bufpool.Get().(*[]byte)
b := *bp
defer func() {
    *bp = b
    bufpool.Put(bp)
}()

// Now, go do interesting things with your byte buffer.
buf := bytes.NewBuffer(b)

```

Some caveats apply. Before Go 1.13, the pool was cleared out every time a garbage collection occured. This might be detrimental to performance in programs that allocate a lot. In 1.13, it [seems that more objects will survive GC's](https://go-review.googlesource.com/c/go/+/162919/).

一些警告适用。在 Go 1.13 之前，每次发生垃圾收集时都会清除池。这可能会损害分配很多的程序的性能。在 1.13 中，[似乎更多的对象将在 GC 中幸存下来](https://go-review.googlesource.com/c/go/+/162919/)。

⚠️ **You must zero out the fields of the struct before putting the object back in the pool**.

⚠️ **在将对象放回池之前，您必须将结构体的字段清零**。

If you don't do this, you can obtain a 'dirty' object from the pool that contains data from previous use. This can be a severe security risk!

如果您不这样做，您可以从包含以前使用过的数据的池中获取一个“脏”对象。这可能是一个严重的安全风险！

```go
type AuthenticationResponse {
    Token string
    UserID string
}

rsp := authPool.Get().(*AuthenticationResponse)
defer authPool.Put(rsp)

// If we don't hit this if statement, we might return data from other users!😱
if blah {
    rsp.UserID = "user-1"
    rsp.Token = "super-secret
}

return rsp

```

The safe way to ensure you always zero memory is to do so explicitly:

确保始终为零内存的安全方法是明确这样做：

```go
// reset resets all fields of the AuthenticationResponse before pooling it.
func (a* AuthenticationResponse) reset() {
    a.Token = ""
    a.UserID = ""
}

rsp := authPool.Get().(*AuthenticationResponse)
defer func() {
    rsp.reset()
    authPool.Put(rsp)
}()

```

The only case in which this is not an issue is when you use _exactly_ the memory that you wrote to. For example:

这不是问题的唯一情况是当您_完全_使用您写入的内存时。例如：

```go
var (
    r io.Reader
    w io.Writer
)

// Obtain a buffer from the pool.
buf := *bufPool.Get().(*[]byte)
defer bufPool.Put(&buf)

// We only write to w exactly what we read from r, and no more.😌
nr, er := r.Read(buf)
if nr > 0 {
    nw, ew := w.Write(buf[0:nr])
}

```

_Edit: a previous version of this blog post did not specify that the `New()` function should return a pointer type. Thanks, [kevinconaway](https://news.ycombinator.com/item?id=20214506)!_

_Edit：此博客文章的先前版本未指定 `New()` 函数应返回指针类型。谢谢，[kevinconaway](https://news.ycombinator.com/item?id=20214506)！_

### Avoid using structures containing pointers as map keys for large maps

### 避免使用包含指针的结构作为大型映射的映射键

Phew - that was quite a mouthful. Sorry about that. Much has been said (a lot by my former colleague, [Phil Pearl](https://twitter.com/philpearl)) on Go's performance with [large heap sizes](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487). During a garbage collection, the runtime scans objects containing pointers, and chases them. If you have a very large `map[string]int`, the GC has to check every string within the map, every GC, as strings contain pointers.

呼——真是一口。对于那个很抱歉。关于 Go 的 [大堆大小](https://syslog.ravelin.com/further-大堆的危险7a267b57d487)。在垃圾回收期间，运行时会扫描包含指针的对象，并追踪它们。如果你有一个非常大的 `map[string]int`，GC 必须检查映射中的每个字符串，每个 GC，因为字符串包含指针。

In this example, we write 10 million elements to a `map[string]int`, and time the garbage collections. We allocate our map at the package scope to ensure it is heap allocated.

在这个例子中，我们将 1000 万个元素写入 `map[string]int`，并对垃圾收集计时。我们在包范围内分配我们的映射以确保它是堆分配的。

```go
package main

import (
    "fmt"
    "runtime"
    "strconv"
    "time"
)

const (
    numElements = 10000000
)

var foo = map[string]int{}

func timeGC() {
    t := time.Now()
    runtime.GC()
    fmt.Printf("gc took: %s\n", time.Since(t))
}

func main() {
    for i := 0;i < numElements;i++ {
        foo[strconv.Itoa(i)] = i
    }

    for {
        timeGC()
        time.Sleep(1 * time.Second)
    }
}

```

Running this program, we see the following:

运行这个程序，我们看到以下内容：

```
🍎 inthash → go install && inthash
gc took: 98.726321ms
gc took: 105.524633ms
gc took: 102.829451ms
gc took: 102.71908ms
gc took: 103.084104ms
gc took: 104.821989ms

```

That's quite a while in computer land! 😰

这在计算机领域已经很长时间了！ 😰

What could we do to improve it? Removing pointers wherever possible seems like a good idea - we'll reduce the number of pointers that the garbage collector has to chase. [Strings contain pointers](https://www.reddit.com/r/golang/comments/4ologg/why_is_byte_used_as_a_string_type/d4e6gy8/); so let's implement this as a `map[int]int`.

我们可以做些什么来改善它？尽可能删除指针似乎是个好主意——我们将减少垃圾收集器必须追逐的指针数量。 [字符串包含指针](https://www.reddit.com/r/golang/comments/4ologg/why_is_byte_used_as_a_string_type/d4e6gy8/);所以让我们将其实现为一个 `map[int]int`。

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

const (
    numElements = 10000000
)

var foo = map[int]int{}

func timeGC() {
    t := time.Now()
    runtime.GC()
    fmt.Printf("gc took: %s\n", time.Since(t))
}

func main() {
    for i := 0;i < numElements;i++ {
        foo[i] = i
    }

    for {
        timeGC()
        time.Sleep(1 * time.Second)
    }
}

```

Running the program again, we get the following:

再次运行程序，我们得到以下信息：

```
🍎 inthash → go install && inthash
gc took: 3.608993ms
gc took: 3.926913ms
gc took: 3.955706ms
gc took: 4.063795ms
gc took: 3.91519ms
gc took: 3.75226ms

```

Much better. We've chopped 97% off garbage collection timings. In a production use case, you'd need to hash the strings to an integer before inserting ino the map.

好多了。我们已经将垃圾收集时间缩短了 97%。在生产用例中，您需要在插入地图之前将字符串散列为整数。

ℹ️ There's plenty more you can do to evade the GC. If you allocate giant arrays of pointerless structs, ints or bytes, [the GC will not scan it](https://medium.com/@rf_14423/did-the-big-allocations-of-ram-contain-pointers-directly-or-indirectly-actual-pointers-strings-76ed28c0bc92): meaning you pay no GC overhead. Such techniques usually require substantial re-working of the program, so we won't delve any further into them today.

ℹ️ 你可以做更多的事情来逃避 GC。如果你分配巨大的无指针结构、整数或字节数组，[GC 不会扫描它](https://medium.com/@rf_14423/did-the-big-allocations-of-ram-contain-pointers-directly-or-indirectly-actual-pointers-strings-76ed28c0bc92)：意味着您无需支付 GC 开销。此类技术通常需要对程序进行大量重新设计，因此我们今天不会深入研究它们。

⚠️ As with all optimisation, your mileage may vary. See a [Twitter thread from Damian Gryski](https://twitter.com/dgryski/status/1140685755578118144) for an interesting example where removing strings from a large map in favour of a smarter data structure actually _increased_ memory. In general, you should read everything he puts out.

⚠️与所有优化一样，您的里程可能会有所不同。请参阅 [来自 Damian Gryski 的 Twitter 线程](https://twitter.com/dgryski/status/1140685755578118144) 获取一个有趣的示例，其中从大地图中删除字符串以支持更智能的数据结构实际上_增加_了内存。一般来说，你应该阅读他发表的所有内容。

### Code generate marshalling code to avoid runtime reflection 

### 代码生成编组代码以避免运行时反射

Marshalling and unmarshalling your structure to and from various serialisation formats like JSON is a common operation; especially when building microservices. In fact, you'll often find that the only thing most microservices are actually doing is serialisation. Functions like `json.Marshal` and `json.Unmarshal` rely on [runtime reflection](https://blog.golang.org/laws-of-reflection) to serialise the struct fields to bytes, and vice versa. This can be slow: reflection is not anywhere near as performant as explicit code.

在各种序列化格式（如 JSON）之间编组和解组您的结构是一种常见操作；尤其是在构建微服务时。事实上，您经常会发现大多数微服务实际上唯一要做的就是序列化。 `json.Marshal` 和 `json.Unmarshal` 等函数依赖于 [运行时反射](https://blog.golang.org/laws-of-reflection) 将结构字段序列化为字节，反之亦然。这可能很慢：反射远不如显式代码那么高效。

However, it doesn't have to be this way. The mechanics of marshalling JSON goes a little something like this:

但是，不必如此。编组 JSON 的机制有点像这样：

```go
package json

// Marshal take an object and returns its representation in JSON.
func Marshal(obj interface{}) ([]byte, error) {
    // Check if this object knows how to marshal itself to JSON
    // by satisfying the Marshaller interface.
    if m, is := obj.(json.Marshaller);is {
        return m.MarshalJSON()
    }

    // It doesn't know how to marshal itself.Do default reflection based marshallling.
    return marshal(obj)
}

```

If we know how to marshal ourselves to JSON, we have a hook for avoiding runtime reflection. But we don't want to hand write all of our marshalling code, so what do we do? Get computers to write code for us! Code generators like [easyjson](https://github.com/mailru/easyjson) look at a struct, and generate highly optimised code which is fully compatible with existing marshalling interfaces like `json.Marshaller`.

如果我们知道如何将自己编组到 JSON，我们就有了一个避免运行时反射的钩子。但是我们不想手写我们所有的编组代码，那么我们该怎么办呢？让计算机为我们编写代码！像 [easyjson](https://github.com/mailru/easyjson) 这样的代码生成器查看一个结构，并生成高度优化的代码，该代码与现有的编组接口（如 `json.Marshaller`)完全兼容。

Download the package, and run the following on your `$file.go` containing the structs you would like to generate code for.

下载包，并在包含要为其生成代码的结构的 `$file.go` 上运行以下命令。

```
easyjson -all $file.go
```

You should find a `$file_easyjson.go` file has been generated. As `easyjson` has implemented the `json.Marshaller` interface for you, these functions will be called instead of the reflection based default. Congratulations: you just sped up your JSON marshalling code by 3x. There's lots of things that you can twiddle to increase the performance even more.

你应该会发现已经生成了一个 `$file_easyjson.go` 文件。由于`easyjson` 已经为你实现了`json.Marshaller` 接口，这些函数将被调用而不是基于反射的默认值。恭喜：您的 JSON 编组代码速度提高了 3 倍。有很多事情你可以摆弄以进一步提高性能。

ℹ️I recommend the package because I have used it before to good effect. Caveat emptor. Please do not take this as an invitation to start aggressive discussions about the fastest JSON marshalling packages on the market with me.

ℹ️我推荐这个包装，因为我之前用过它，效果很好。买者自负。请不要以此为邀请与我开始积极讨论市场上最快的 JSON 编组包。

⚠️ You'll need to make sure to re-generate the marshalling code when you change the struct. If you forget, new fields that you add won't be serialised and de-serialised, which can be confusing! You can use `go generate` to handle this code generation for you. In order to keep these in sync with structures, I like to have a `generate.go` in the root of the package that calls `go generate` for all files in the package: this can aid maintenance when you have many files that need generating. Top tip: call `go generate` in CI and check it has no diffs with the checked in code to ensure that structures are up to date.

⚠️ 更改结构时，您需要确保重新生成编组代码。如果您忘记了，您添加的新字段不会被序列化和反序列化，这可能会令人困惑！您可以使用 `go generate` 为您处理此代码生成。为了使这些与结构保持同步，我喜欢在包的根目录中有一个 `generate.go`，它为包中的所有文件调用 `go generate`：当你有许多需要的文件时，这可以帮助维护产生。重要提示：在 CI 中调用 `go generate` 并检查它与签入的代码没有差异，以确保结构是最新的。

### Use `strings.Builder` to build up strings

### 使用 `strings.Builder` 来构建字符串

In Go, strings are immutable: think of them as a read only slice of bytes. This means that every time you create a string, you're allocating new memory, and potentially creating more work for the garbage collector.

在 Go 中，字符串是不可变的：将它们视为只读的字节片。这意味着每次创建字符串时，都在分配新内存，并可能为垃圾收集器创建更多工作。

In Go 1.10, [`strings.Builder`](https://golang.org/pkg/strings/#Builder) was introduced as an efficient way to build up strings. Internally, it writes to a byte buffer. Only upon calling `String()` on the builder, is the string actually created. It relies on some `unsafe` trickery to return the underlying bytes as a string with zero allocations: see [this blog](https://syslog.ravelin.com/byte-vs-string-in-go-d645b67ca7ff) for a further look into how it works.

在 Go 1.10 中，引入了 [`strings.Builder`](https://golang.org/pkg/strings/#Builder) 作为构建字符串的有效方法。在内部，它写入字节缓冲区。只有在构建器上调用 `String()` 时，字符串才会被实际创建。它依赖于一些“不安全”的技巧来将底层字节作为零分配的字符串返回：请参阅[此博客](https://syslog.ravelin.com/byte-vs-string-in-go-d645b67ca7ff)进一步研究它是如何工作的。

Let's do a performance comparison to verify the two approaches:

让我们做一个性能比较来验证这两种方法：

```go
// main.go
package main

import "strings"

var strs = []string{
    "here's",
    "a",
    "some",
    "long",
    "list",
    "of",
    "strings",
    "for",
    "you",
}

func buildStrNaive() string {
    var s string

    for _, v := range strs {
        s += v
    }

    return s
}

func buildStrBuilder() string {
    b := strings.Builder{}

    // Grow the buffer to a decent length, so we don't have to continually
    // re-allocate.
    b.Grow(60)

    for _, v := range strs {
        b.WriteString(v)
    }

    return b.String()
}

```

```go
// main_test.go
package main

import (
    "testing"
)

var str string

func BenchmarkStringBuildNaive(b *testing.B) {
    for i := 0;i < b.N;i++ {
        str = buildStrNaive()
    }
}
func BenchmarkStringBuildBuilder(b *testing.B) {
    for i := 0;i < b.N;i++ {
        str = buildStrBuilder()
    }

```

I get the following results on my Macbook Pro:

我在 Macbook Pro 上得到以下结果：

```
🍎 strbuild → go test -bench=.-benchmem
goos: darwin
goarch: amd64
pkg: github.com/sjwhitworth/perfblog/strbuild
BenchmarkStringBuildNaive-8          5000000           255 ns/op         216 B/op          8 allocs/op
BenchmarkStringBuildBuilder-8       20000000            54.9 ns/op        64 B/op          1 allocs/op

```

As we can see, `strings.Builder` is 4.7x faster, results in 1/8th of the number of allocations, and 1/4 of the memory allocated.

正如我们所看到的，`strings.Builder` 的速度提高了 4.7 倍，分配次数是 1/8，分配的内存是 1/4。

Where performance matters, use `strings.Builder`. In general, I recommend using it for all but the most trivial cases of building strings.

如果性能很重要，请使用 `strings.Builder`。一般来说，我建议将它用于除了构建字符串的最琐碎的情况之外的所有情况。

### Use strconv instead of fmt

### 使用 strconv 而不是 fmt

[`fmt`](https://golang.org/pkg/fmt/) is one of the most well known packages in Go. You'll have probably used it in your first Go program to print "hello, world" to the screen. However, when it comes to converting integers and floats in to strings, it's not as performant as its lower level cousin: [`strconv`](https://golang.org/pkg/strconv/). This package gives you a decent whack more performance, for some very small changes in API.

[`fmt`](https://golang.org/pkg/fmt/) 是 Go 中最著名的包之一。您可能已经在第一个 Go 程序中使用它在屏幕上打印“hello, world”。但是，在将整数和浮点数转换为字符串时，它的性能不如它的低级表亲：[`strconv`](https://golang.org/pkg/strconv/)。对于 API 中的一些非常小的更改，此包为您提供了更好的性能。

`fmt` mostly takes `interface{}` as arguments to functions. This has two downsides:

`fmt` 主要将 `interface{}` 作为函数的参数。这有两个缺点：

- You lose type safety. This is a big one, for me.
- It can increase the number of allocations needed. Passing a non-pointer type as an`interface{}` usually causes heap allocations. Read into [this blog](https://www.darkcoding.net/software/go-the-price-of-interface/) to figure out why that is.

- 你失去了类型安全。这是一个很大的，对我来说。
- 它可以增加所需的分配数量。将非指针类型作为`interface{}` 传递通常会导致堆分配。阅读 [本博客](https://www.darkcoding.net/software/go-the-price-of-interface/) 找出原因。

The below program shows the difference in performance:

以下程序显示了性能差异：

```go
// main.go
package main

import (
    "fmt"
    "strconv"
)

func strconvFmt(a string, b int) string {
    return a + ":" + strconv.Itoa(b)
}

func fmtFmt(a string, b int) string {
    return fmt.Sprintf("%s:%d", a, b)
}

func main() {}

```

```go
// main_test.go
package main

import (
    "testing"
)

var (
    a    = "boo"
    blah = 42
    box  = ""
)

func BenchmarkStrconv(b *testing.B) {
    for i := 0;i < b.N;i++ {
        box = strconvFmt(a, blah)
    }
    a = box
}

func BenchmarkFmt(b *testing.B) {
    for i := 0;i < b.N;i++ {
        box = fmtFmt(a, blah)
    }
    a = box
}

```

The benchmark results on a Macbook Pro:

Macbook Pro 的基准测试结果：

```go
🍎 strfmt → go test -bench=.-benchmem
goos: darwin
goarch: amd64
pkg: github.com/sjwhitworth/perfblog/strfmt
BenchmarkStrconv-8      30000000            39.5 ns/op        32 B/op          1 allocs/op
BenchmarkFmt-8          10000000           143 ns/op          72 B/op          3 allocs/op

```

We can see that the `strconv` version is 3.5x faster, results in 1/3rd the number of allocations, and half the memory allocated.

我们可以看到 `strconv` 版本快了 3.5 倍，导致分配数量减少 1/3，分配的内存减少一半。

### Allocate capacity in make to avoid re-allocation

### 在make中分配容量以避免重新分配

Before we get to performance improvements, let's take a quick refresher on slices. The slice is a very useful construct in Go. It provides a re-sizable array, with the ability to take different views on the same underlying memory without re-allocation. If you peek under the hood, the slice is made up of three elements:

在我们开始性能改进之前，让我们快速回顾一下切片。切片是 Go 中非常有用的构造。它提供了一个可重新调整大小的数组，能够在不重新分配的情况下对同一底层内存采取不同的观点。如果你在引擎盖下偷看，切片由三个元素组成：

```go
type slice struct {
    // pointer to underlying data in the slice.
    data uintptr
    // the number of elements in the slice.
    len int
    // the number of elements that the slice can
    // grow to before a new underlying array
    // is allocated.
    cap int
}
```

What are these fields?

这些领域是什么？

- `data`: pointer to underlying data in the slice
- `len`: the current number of elements in the slice.
- `cap`: the number of elements that the slice can grow to before re-allocation.

- `data`：指向切片中底层数据的指针
- `len`：切片中的当前元素数。
- `cap`：在重新分配之前切片可以增长到的元素数量。

Under the hood, slices are fixed length arrays. When you reach the `cap` of a slice, a new array with double the cap of the previous slice is allocated, the memory is copied over from the old slice to the new one, and the old array is discarded

在引擎盖下，切片是固定长度的数组。当你到达一个切片的‘cap’时，会分配一个新的数组，它的大小是前一个切片的两倍，内存从旧切片复制到新切片，旧数组被丢弃

I often see code like the following that allocates a slice with zero capacity, when the capacity of the slice is known upfront.

我经常看到类似下面的代码，当切片的容量预先知道时，它会分配一个容量为零的切片。

```go
var userIDs []string
for _, bar := range rsp.Users {
    userIDs = append(userIDs, bar.ID)
}
```

In this case, the slice starts off with zero length, and zero capacity. After receiving the response, we append the users to the slice. As we do so, we hit the capacity of the slice: a new underlying array is allocated that is double the capacity of the previous slice, and the data from the slice is copied into it. If we had 8 users in the response, this would result in 5 allocations.

在这种情况下，切片从零长度和零容量开始。收到响应后，我们将用户附加到切片。当我们这样做时，我们达到了切片的容量：分配了一个新的底层数组，它是前一个切片容量的两倍，并将切片中的数据复制到其中。如果我们在响应中有 8 个用户，这将导致 5 个分配。

A far more efficient way is to change it to the following:

更有效的方法是将其更改为以下内容：

```go
userIDs := make([]string, 0, len(rsp.Users)

for _, bar := range rsp.Users {
    userIDs = append(userIDs, bar.ID)
}

```

We have explicitly allocated the capacity to the slice by using make. Now, we can append to the slice, knowing that we won't trigger additional allocations and copys.

我们已经使用 make 显式地为切片分配了容量。现在，我们可以追加到切片，知道我们不会触发额外的分配和复制。

If you don't know how much you should allocate because the capacity is dynamic or calculated later in the program, measure the distribution of the size of the slice that you end up with whilst the program is running. I usually take the 90th or 99th percentile, and hardcode the value in the progam. In cases where you have RAM to trade off for CPU, set this value to higher than you think you'll need.

如果您不知道应该分配多少容量，因为容量是动态的或在程序稍后计算，请测量程序运行时最终得到的切片大小的分布。我通常取第 90 个或第 99 个百分位数，并在程序中硬编码该值。如果您有 RAM 来换取 CPU，请将此值设置为高于您认为需要的值。

This advice is also applicable to maps: using `make(map[string]string, len(foo))` will allocate enough capacity under the hood to avoid re-allocation.

这个建议也适用于地图：使用 `make(map[string]string, len(foo))` 将在引擎盖下分配足够的容量来避免重新分配。

ℹ️ See [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) for more information about how slices work under the hood.

ℹ️ 请参阅 [Go Slices：用法和内部结构](https://blog.golang.org/go-slices-usage-and-internals) 以了解有关切片如何在幕后工作的更多信息。

### Use methods that allow you to pass byte slices

### 使用允许您传递字节切片的方法

When using packages, look to use methods that allow you to pass a byte slice: these methods usually give you more control over allocation.

使用包时，请注意使用允许您传递字节切片的方法：这些方法通常可以让您更好地控制分配。

[`time.Format`](https://golang.org/pkg/time/#Time.Format) vs. [`time.AppendFormat`](https://golang.org/pkg/time/#Time.AppendFormat) is a good example. `time.Format` returns a string. Under the hood, this allocates a new byte slice and calls `time.AppendFormat` on it. `time.AppendFormat` takes a byte buffer, writes the formatted representation of the time, and returns the extended byte slice. This is common in other packages in the standard library: see [`strconv.AppendFloat`](https://golang.org/pkg/strconv/#AppendFloat), or [`bytes.NewBuffer`](https://golang.org/pkg/bytes/#NewBuffer).

[`time.Format`](https://golang.org/pkg/time/#Time.Format) 与 [`time.AppendFormat`](https://golang.org/pkg/time/#Time.AppendFormat) 就是一个很好的例子。 `time.Format` 返回一个字符串。在幕后，这会分配一个新的字节切片并对其调用`time.AppendFormat`。 `time.AppendFormat` 接受一个字节缓冲区，写入时间的格式化表示，并返回扩展字节片。这在标准库中的其他包中很常见：请参阅 [`strconv.AppendFloat`](https://golang.org/pkg/strconv/#AppendFloat) 或 [`bytes.NewBuffer`](https://golang.org/pkg/bytes/#NewBuffer)。

Why does this give you increased performance? Well, you can now pass byte slices that you've obtained from your `sync.Pool`, instead of allocating a new buffer every time. Or you can increase the initial buffer size to a value that you know is more suited to your program, to reduce slice re-copying.

为什么这可以提高性能？好吧，您现在可以传递从“sync.Pool”中获得的字节切片，而不是每次都分配一个新缓冲区。或者您可以将初始缓冲区大小增加到您知道更适合您的程序的值，以减少切片重新复制。

## Summary

##  概括

After reading this post, you should be able to take these techniques and apply them to your code base. Over time, you will build a mental model for reasoning about performance in Go programs. This can aid upfront design greatly.

阅读完这篇文章后，您应该能够采用这些技术并将它们应用到您的代码库中。随着时间的推移，你将建立一个思维模型来推理 Go 程序的性能。这可以极大地帮助前期设计。

To close, a word of caution. Take my guidance as situationally-dependent advice, not gospel. Benchmark and measure.

要关闭，请注意。将我的指导作为视情况而定的建议，而不是福音。基准和测量。

Know when to stop. Improving performance of a system is a feel-good exercise for an engineer: the problem is interesting, and the results are immediately visible. However, the usefulness of performance improvements is very dependent on the situation. If your service takes 10ms to respond, and 90ms round trip on the network, it doesn't seem worth it to try and halve that 10ms to 5ms: you'll still take 95ms. If you manage to optimise it to within an inch of it's life so it takes 1ms to respond, you're still only at 91ms. There are probably bigger fish to fry.

知道什么时候停止。提高系统性能对工程师来说是一种感觉良好的练习：问题很有趣，结果立即可见。但是，性能改进的有用性很大程度上取决于情况。如果您的服务需要 10 毫秒来响应，并且在网络上往返需要 90 毫秒，那么尝试将 10 毫秒减半到 5 毫秒似乎不值得：您仍然需要 95 毫秒。如果你设法将它优化到它的生命的一英寸之内，所以它需要 1 毫秒的响应时间，那么你仍然只有 91 毫秒。可能有更大的鱼要煎。

Optimise wisely!

明智地优化！

### References

###  参考

I've linked heavily throughout the blog. If you're interested in further reading, the following are great sources of inspiration:

我在整个博客中都有大量链接。如果您有兴趣进一步阅读，以下是很好的灵感来源：

- [Further dangers of large heaps in Go](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487)
- [Allocation efficiency in high performance Go services](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/) 

- [Go 中大堆的进一步危险](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487)
- [高性能 Go 服务中的分配效率](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/)

