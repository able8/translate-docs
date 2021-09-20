# Simple techniques to optimise Go programs

I'm very interested in performance. I'm not sure I can explain the underlying reasons for it. I'm easily frustrated at slow services and programs, and it seems like [I'm not alone.](http://glinden.blogspot.com/2006/11/marissa-mayer-at-web-20.html)

> _In A/B tests, we tried delaying the page in increments of 100 milliseconds and found that even very small delays would result in substantial and costly drops in revenue._ \- Greg Linden, Amazon.com

From my experience, poor performance manifests in one of two ways:

- Operations that performed well at small scale, but become unviable as the number of users grows. These are usually O(N) or O(N²) operations. When your user base is small, these perform just fine, and are often done in order to get a product to market. As your use base grows, you see more[pathological examples](https://theoutline.com/post/4147/in-twitters-early-days-only-one-celebrity-could-tweet-at-a-time?zd=1&zi=ivqvd4py) that you weren't expecting, and your service grinds to a halt.
- Many individual sources of small optimisation - AKA 'death by a thousand crufts'.

I've spent most of my career either doing data science with Python, or building services in Go; I have far more experience optimising the latter. Go is usually not the bottleneck in the services I write - programs are often IO bound as they talk to the database. However, in batch machine learning pipelines - like I built in my previous role - your program is often CPU bound. When your Go program is using excessive CPU, and the excessive usage is having a negative impact, there's various strategies you can use to mitigate that.

This post explains some techniques you can use to significantly improve the performance of your program with little effort. I am deliberately ignoring techniques that require significant effort, or large changes to program structure.

## Before you start

Before you make any changes to your program, invest time in creating a proper baseline to compare against. If you don't do this, you'll be searching around in the dark, wondering if your changes are having any improvement. Write benchmarks first, and grab [profiles](https://blog.golang.org/profiling-go-programs) for use in pprof. In the best case, this will be a [Go benchmark](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go): this allows easy usage of pprof, and memory allocation profiling. You should also use [`benchcmp`](https://godoc.org/golang.org/x/tools/cmd/benchcmp): a helpful tool for comparing the difference in performance between two benchmarks.

If your code is not easily benchmarked, just start with something that you can time. You can profile your code manually using [`runtime/pprof`](https://golang.org/pkg/runtime/pprof/).

Let's get started!

### Use `sync.Pool` to re-use previously allocated objects

[`sync.Pool`](https://golang.org/pkg/sync/#Pool) implements a [free-list](https://en.wikipedia.org/wiki/Free_list). This allows you to re-use structures that you've previously allocated. This amortises the allocation of an object over many usages, reducing the work the garbage collector has to do. The API is very simple: implement a function that allocates a new instance of your object. It should return a pointer type.

```
var bufpool = sync.Pool{
    New: func() interface{} {
        buf := make([]byte, 512)
        return &buf
    }}

```

After this, you can `Get()` objects from the pool, `Put()` ting them back after you are done.

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

⚠️ **You must zero out the fields of the struct before putting the object back in the pool**.

If you don't do this, you can obtain a 'dirty' object from the pool that contains data from previous use. This can be a severe security risk!

```
type AuthenticationResponse {
    Token string
    UserID string
}

rsp := authPool.Get().(*AuthenticationResponse)
defer authPool.Put(rsp)

// If we don't hit this if statement, we might return data from other users! 😱
if blah {
    rsp.UserID = "user-1"
    rsp.Token = "super-secret
}

return rsp

```

The safe way to ensure you always zero memory is to do so explicitly:

```
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

```
var (
    r io.Reader
    w io.Writer
)

// Obtain a buffer from the pool.
buf := *bufPool.Get().(*[]byte)
defer bufPool.Put(&buf)

// We only write to w exactly what we read from r, and no more. 😌
nr, er := r.Read(buf)
if nr > 0 {
    nw, ew := w.Write(buf[0:nr])
}

```

_Edit: a previous version of this blog post did not specify that the `New()` function should return a pointer type. Thanks, [kevinconaway](https://news.ycombinator.com/item?id=20214506)!_

### Avoid using structures containing pointers as map keys for large maps

Phew - that was quite a mouthful. Sorry about that. Much has been said (a lot by my former colleague, [Phil Pearl](https://twitter.com/philpearl)) on Go's performance with [large heap sizes](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487). During a garbage collection, the runtime scans objects containing pointers, and chases them. If you have a very large `map[string]int`, the GC has to check every string within the map, every GC, as strings contain pointers.

In this example, we write 10 million elements to a `map[string]int`, and time the garbage collections. We allocate our map at the package scope to ensure it is heap allocated.

```
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
    for i := 0; i < numElements; i++ {
        foo[strconv.Itoa(i)] = i
    }

    for {
        timeGC()
        time.Sleep(1 * time.Second)
    }
}

```

Running this program, we see the following:

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

What could we do to improve it? Removing pointers wherever possible seems like a good idea - we'll reduce the number of pointers that the garbage collector has to chase. [Strings contain pointers](https://www.reddit.com/r/golang/comments/4ologg/why_is_byte_used_as_a_string_type/d4e6gy8/); so let's implement this as a `map[int]int`.

```
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
    for i := 0; i < numElements; i++ {
        foo[i] = i
    }

    for {
        timeGC()
        time.Sleep(1 * time.Second)
    }
}

```

Running the program again, we get the following:

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

ℹ️ There's plenty more you can do to evade the GC. If you allocate giant arrays of pointerless structs, ints or bytes, [the GC will not scan it](https://medium.com/@rf_14423/did-the-big-allocations-of-ram-contain-pointers-directly-or-indirectly-actual-pointers-strings-76ed28c0bc92): meaning you pay no GC overhead. Such techniques usually require substantial re-working of the program, so we won't delve any further into them today.

⚠️ As with all optimisation, your mileage may vary. See a [Twitter thread from Damian Gryski](https://twitter.com/dgryski/status/1140685755578118144) for an interesting example where removing strings from a large map in favour of a smarter data structure actually _increased_ memory. In general, you should read everything he puts out.

### Code generate marshalling code to avoid runtime reflection

Marshalling and unmarshalling your structure to and from various serialisation formats like JSON is a common operation; especially when building microservices. In fact, you'll often find that the only thing most microservices are actually doing is serialisation. Functions like `json.Marshal` and `json.Unmarshal` rely on [runtime reflection](https://blog.golang.org/laws-of-reflection) to serialise the struct fields to bytes, and vice versa. This can be slow: reflection is not anywhere near as performant as explicit code.

However, it doesn't have to be this way. The mechanics of marshalling JSON goes a little something like this:

```
package json

// Marshal take an object and returns its representation in JSON.
func Marshal(obj interface{}) ([]byte, error) {
    // Check if this object knows how to marshal itself to JSON
    // by satisfying the Marshaller interface.
    if m, is := obj.(json.Marshaller); is {
        return m.MarshalJSON()
    }

    // It doesn't know how to marshal itself. Do default reflection based marshallling.
    return marshal(obj)
}

```

If we know how to marshal ourselves to JSON, we have a hook for avoiding runtime reflection. But we don't want to hand write all of our marshalling code, so what do we do? Get computers to write code for us! Code generators like [easyjson](https://github.com/mailru/easyjson) look at a struct, and generate highly optimised code which is fully compatible with existing marshalling interfaces like `json.Marshaller`.

Download the package, and run the following on your `$file.go` containing the structs you would like to generate code for.

```
easyjson -all $file.go

```

You should find a `$file_easyjson.go` file has been generated. As `easyjson` has implemented the `json.Marshaller` interface for you, these functions will be called instead of the reflection based default. Congratulations: you just sped up your JSON marshalling code by 3x. There's lots of things that you can twiddle to increase the performance even more.

ℹ️I recommend the package because I have used it before to good effect. Caveat emptor. Please do not take this as an invitation to start aggressive discussions about the fastest JSON marshalling packages on the market with me.

⚠️ You'll need to make sure to re-generate the marshalling code when you change the struct. If you forget, new fields that you add won't be serialised and de-serialised, which can be confusing! You can use `go generate` to handle this code generation for you. In order to keep these in sync with structures, I like to have a `generate.go` in the root of the package that calls `go generate` for all files in the package: this can aid maintenance when you have many files that need generating. Top tip: call `go generate` in CI and check it has no diffs with the checked in code to ensure that structures are up to date.

### Use `strings.Builder` to build up strings

In Go, strings are immutable: think of them as a read only slice of bytes. This means that every time you create a string, you're allocating new memory, and potentially creating more work for the garbage collector.

In Go 1.10, [`strings.Builder`](https://golang.org/pkg/strings/#Builder) was introduced as an efficient way to build up strings. Internally, it writes to a byte buffer. Only upon calling `String()` on the builder, is the string actually created. It relies on some `unsafe` trickery to return the underlying bytes as a string with zero allocations: see [this blog](https://syslog.ravelin.com/byte-vs-string-in-go-d645b67ca7ff) for a further look into how it works.

Let's do a performance comparison to verify the two approaches:

```
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

```
// main_test.go
package main

import (
    "testing"
)

var str string

func BenchmarkStringBuildNaive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        str = buildStrNaive()
    }
}
func BenchmarkStringBuildBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        str = buildStrBuilder()
    }

```

I get the following results on my Macbook Pro:

```
🍎 strbuild → go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/sjwhitworth/perfblog/strbuild
BenchmarkStringBuildNaive-8          5000000           255 ns/op         216 B/op          8 allocs/op
BenchmarkStringBuildBuilder-8       20000000            54.9 ns/op        64 B/op          1 allocs/op

```

As we can see, `strings.Builder` is 4.7x faster, results in 1/8th of the number of allocations, and 1/4 of the memory allocated.

Where performance matters, use `strings.Builder`. In general, I recommend using it for all but the most trivial cases of building strings.

### Use strconv instead of fmt

[`fmt`](https://golang.org/pkg/fmt/) is one of the most well known packages in Go. You'll have probably used it in your first Go program to print "hello, world" to the screen. However, when it comes to converting integers and floats in to strings, it's not as performant as its lower level cousin: [`strconv`](https://golang.org/pkg/strconv/). This package gives you a decent whack more performance, for some very small changes in API.

`fmt` mostly takes `interface{}` as arguments to functions. This has two downsides:

- You lose type safety. This is a big one, for me.
- It can increase the number of allocations needed. Passing a non-pointer type as an`interface{}` usually causes heap allocations. Read into [this blog](https://www.darkcoding.net/software/go-the-price-of-interface/) to figure out why that is.

The below program shows the difference in performance:

```
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

```
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
    for i := 0; i < b.N; i++ {
        box = strconvFmt(a, blah)
    }
    a = box
}

func BenchmarkFmt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        box = fmtFmt(a, blah)
    }
    a = box
}

```

The benchmark results on a Macbook Pro:

```
🍎 strfmt → go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/sjwhitworth/perfblog/strfmt
BenchmarkStrconv-8      30000000            39.5 ns/op        32 B/op          1 allocs/op
BenchmarkFmt-8          10000000           143 ns/op          72 B/op          3 allocs/op

```

We can see that the `strconv` version is 3.5x faster, results in 1/3rd the number of allocations, and half the memory allocated.

### Allocate capacity in make to avoid re-allocation

Before we get to performance improvements, let's take a quick refresher on slices. The slice is a very useful construct in Go. It provides a re-sizable array, with the ability to take different views on the same underlying memory without re-allocation. If you peek under the hood, the slice is made up of three elements:

```
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

- `data`: pointer to underlying data in the slice
- `len`: the current number of elements in the slice.
- `cap`: the number of elements that the slice can grow to before re-allocation.

Under the hood, slices are fixed length arrays. When you reach the `cap` of a slice, a new array with double the cap of the previous slice is allocated, the memory is copied over from the old slice to the new one, and the old array is discarded

I often see code like the following that allocates a slice with zero capacity, when the capacity of the slice is known upfront.

```
var userIDs []string
for _, bar := range rsp.Users {
    userIDs = append(userIDs, bar.ID)
}

```

In this case, the slice starts off with zero length, and zero capacity. After receiving the response, we append the users to the slice. As we do so, we hit the capacity of the slice: a new underlying array is allocated that is double the capacity of the previous slice, and the data from the slice is copied into it. If we had 8 users in the response, this would result in 5 allocations.

A far more efficient way is to change it to the following:

```
userIDs := make([]string, 0, len(rsp.Users)

for _, bar := range rsp.Users {
    userIDs = append(userIDs, bar.ID)
}

```

We have explicitly allocated the capacity to the slice by using make. Now, we can append to the slice, knowing that we won't trigger additional allocations and copys.

If you don't know how much you should allocate because the capacity is dynamic or calculated later in the program, measure the distribution of the size of the slice that you end up with whilst the program is running. I usually take the 90th or 99th percentile, and hardcode the value in the progam. In cases where you have RAM to trade off for CPU, set this value to higher than you think you'll need.

This advice is also applicable to maps: using `make(map[string]string, len(foo))` will allocate enough capacity under the hood to
avoid re-allocation.


ℹ️ See [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) for more information about how slices work under the hood.

### Use methods that allow you to pass byte slices

When using packages, look to use methods that allow you to pass a byte slice: these methods usually give you more control over allocation.

[`time.Format`](https://golang.org/pkg/time/#Time.Format) vs. [`time.AppendFormat`](https://golang.org/pkg/time/#Time.AppendFormat) is a good example. `time.Format` returns a string. Under the hood, this allocates a new byte slice and calls `time.AppendFormat` on it. `time.AppendFormat` takes a byte buffer, writes the formatted representation of the time, and returns the extended byte slice. This is common in other packages in the standard library: see [`strconv.AppendFloat`](https://golang.org/pkg/strconv/#AppendFloat), or [`bytes.NewBuffer`](https://golang.org/pkg/bytes/#NewBuffer).

Why does this give you increased performance? Well, you can now pass byte slices that you've obtained from your `sync.Pool`, instead of allocating a new buffer every time. Or you can increase the initial buffer size to a value that you know is more suited to your program, to reduce slice re-copying.

## Summary

After reading this post, you should be able to take these techniques and apply them to your code base. Over time, you will build a mental model for reasoning about performance in Go programs. This can aid upfront design greatly.

To close, a word of caution. Take my guidance as situationally-dependent advice, not gospel. Benchmark and measure.

Know when to stop. Improving performance of a system is a feel-good exercise for an engineer: the problem is interesting, and the results are immediately visible. However, the usefulness of performance improvements is very dependent on the situation. If your service takes 10ms to respond, and 90ms round trip on the network, it doesn't seem worth it to try and halve that 10ms to 5ms: you'll still take 95ms. If you manage to optimise it to within an inch of it's life so it takes 1ms to respond, you're still only at 91ms. There are probably bigger fish to fry.

Optimise wisely!

### References

I've linked heavily throughout the blog. If you're interested in further reading, the following are great sources of inspiration:

- [Further dangers of large heaps in Go](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487)
- [Allocation efficiency in high performance Go services](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/)