# Concurrency

# 并发

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/concurrency)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/concurrency)**

Here's the setup: a colleague has written a function, `CheckWebsites`, that checks the status of a list of URLs.

设置如下：一位同事编写了一个函数“CheckWebsites”，用于检查 URL 列表的状态。

```go
package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

    for _, url := range urls {
        results[url] = wc(url)
    }

    return results
}
```

It returns a map of each URL checked to a boolean value - `true` for a good response, `false` for a bad response.

它返回检查为布尔值的每个 URL 的映射 - “true”表示良好的响应，“false”表示不良响应。

You also have to pass in a `WebsiteChecker` which takes a single URL and returns a boolean. This is used by the function to check all the websites.

您还必须传入一个“WebsiteChecker”，它接受一个 URL 并返回一个布尔值。该功能使用它来检查所有网站。

Using [dependency injection](https://github.com/quii/learn-go-with-tests/blob/main/dependency-injection.md) has allowed them to test the function without making real HTTP calls, making it reliable and fast.

使用[依赖注入](https://github.com/quii/learn-go-with-tests/blob/main/dependency-injection.md) 允许他们在不进行真正的 HTTP 调用的情况下测试函数，使其可靠并且很快。

Here's the test they've written:

这是他们编写的测试：

```go
package concurrency

import (
    "reflect"
    "testing"
)

func mockWebsiteChecker(url string) bool {
    if url == "waat://furhurterwe.geds" {
        return false
    }
    return true
}

func TestCheckWebsites(t *testing.T) {
    websites := []string{
        "http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
    }

    want := map[string]bool{
        "http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
    }

    got := CheckWebsites(mockWebsiteChecker, websites)

    if !reflect.DeepEqual(want, got) {
        t.Fatalf("Wanted %v, got %v", want, got)
    }
}
```

The function is in production and being used to check hundreds of websites. But your colleague has started to get complaints that it's slow, so they've asked you to help speed it up.

该功能正在生产中，用于检查数百个网站。但是你的同事开始抱怨它很慢，所以他们要求你帮助加快速度。

## Write a test

## 写一个测试

Let's use a benchmark to test the speed of `CheckWebsites` so that we can see the effect of our changes.

让我们使用一个基准来测试`CheckWebsites` 的速度，以便我们可以看到我们更改的效果。

```go
package concurrency

import (
    "testing"
    "time"
)

func slowStubWebsiteChecker(_ string) bool {
    time.Sleep(20 * time.Millisecond)
    return true
}

func BenchmarkCheckWebsites(b *testing.B) {
    urls := make([]string, 100)
    for i := 0;i < len(urls);i++ {
        urls[i] = "a url"
    }

    for i := 0;i < b.N;i++ {
        CheckWebsites(slowStubWebsiteChecker, urls)
    }
}
```

The benchmark tests `CheckWebsites` using a slice of one hundred urls and uses a new fake implementation of `WebsiteChecker`. `slowStubWebsiteChecker` is deliberately slow. It uses `time.Sleep` to wait exactly twenty milliseconds and then it returns true.

基准测试使用一百个 url 的切片测试 `CheckWebsites`，并使用 `WebsiteChecker` 的新假实现。 `slowStubWebsiteChecker` 故意变慢。它使用 `time.Sleep` 来等待 20 毫秒，然后返回 true。

When we run the benchmark using `go test -bench=.` (or if you're in Windows Powershell `go test -bench="."`):

当我们使用 `go test -bench=.` 运行基准测试时（或者如果你使用的是 Windows Powershell `go test -bench="."`）：

```go
pkg: github.com/gypsydave5/learn-go-with-tests/concurrency/v0
BenchmarkCheckWebsites-4               1        2249228637 ns/op
PASS
ok      github.com/gypsydave5/learn-go-with-tests/concurrency/v0        2.268s
```

`CheckWebsites` has been benchmarked at 2249228637 nanoseconds - about two and a quarter seconds.

`CheckWebsites` 的基准测试时间为 2249228637 纳秒 - 大约两又四分之一秒。

Let's try and make this faster.

让我们试着让它更快。

### Write enough code to make it pass

### 编写足够的代码使其通过

Now we can finally talk about concurrency which, for the purposes of the following, means 'having more than one thing in progress'. This is something that we do naturally everyday.

现在我们终于可以谈论并发性了，就以下而言，它意味着“有不止一件事在进行中”。这是我们每天自然而然地做的事情。

For instance, this morning I made a cup of tea. I put the kettle on and then, while I was waiting for it to boil, I got the milk out of the fridge, got the tea out of the cupboard, found my favourite mug, put the teabag into the cup and then, when the kettle had boiled, I put the water in the cup.

例如，今天早上我泡了一杯茶。我打开水壶，等它烧开的时候，我从冰箱里拿出牛奶，从橱柜里拿出茶，找到我最喜欢的杯子，把茶包放进杯子里，然后，当水壶烧开了，我把水倒进杯子里。

What I *didn't* do was put the kettle on and then stand there blankly staring at the kettle until it boiled, then do everything else once the kettle had boiled.

我*没有*做的是把水壶放在上面，然后站在那里茫然地盯着水壶直到它沸腾，然后在水壶沸腾后做其他一切。

If you can understand why it's faster to make tea the first way, then you can understand how we will make `CheckWebsites` faster. Instead of waiting for a website to respond before sending a request to the next website, we will tell our computer to make the next request while it is waiting. 

如果您能理解为什么第一种方式泡茶更快，那么您就可以理解我们将如何让 `CheckWebsites` 变得更快。我们不会在向下一个网站发送请求之前等待网站响应，而是告诉我们的计算机在等待时发出下一个请求。

Normally in Go when we call a function `doSomething()` we wait for it to return (even if it has no value to return, we still wait for it to finish). We say that this operation is *blocking* - it makes us wait for it to finish. An operation that does not block in Go will run in a separate *process* called a *goroutine*. Think of a process as reading down the page of Go code from top to bottom, going 'inside' each function when it gets called to read what it does. When a separate process starts it's like another reader begins reading inside the function, leaving the original reader to carry on going down the page.

通常在 Go 中，当我们调用函数 `doSomething()` 时，我们会等待它返回（即使它没有返回值，我们仍然会等待它完成）。我们说这个操作是 *blocking* - 它让我们等待它完成。在 Go 中不会阻塞的操作将在一个单独的 *进程* 中运行，称为 *goroutine*。把一个过程想象成从上到下阅读 Go 代码的页面，当它被调用时进入“内部”每个函数以阅读它的作用。当一个单独的进程开始时，就像另一个读者开始在函数内部阅读，让原来的读者继续往下读。

To tell Go to start a new goroutine we turn a function call into a `go` statement by putting the keyword `go` in front of it: `go doSomething()`.

为了告诉 Go 启动一个新的 goroutine，我们将一个函数调用变成了一个 `go` 语句，在它前面加上关键字 `go`：`go doSomething()`。

```go
package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

    for _, url := range urls {
        go func() {
            results[url] = wc(url)
        }()
    }

    return results
}
```

Because the only way to start a goroutine is to put `go` in front of a function call, we often use *anonymous functions* when we want to start a goroutine. An anonymous function literal looks just the same as a normal function declaration, but without a name (unsurprisingly). You can see one above in the body of the `for` loop.

因为启动 goroutine 的唯一方法是将 `go` 放在函数调用的前面，所以当我们想要启动一个 goroutine 时，我们经常使用*匿名函数*。匿名函数字面量看起来与普通函数声明相同，但没有名称（不出所料）。您可以在 `for` 循环的主体中看到上面的一个。

Anonymous functions have a number of features which make them useful, two of which we're using above. Firstly, they can be executed at the same time that they're declared - this is what the `()` at the end of the anonymous function is doing. Secondly they maintain access to the lexical scope they are defined in - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

匿名函数有许多使它们有用的特性，我们在上面使用了其中的两个特性。首先，它们可以在声明的同时执行——这就是匿名函数末尾的`()`所做的。其次，它们保持对定义它们的词法作用域的访问权——在您声明匿名函数时可用的所有变量在函数体中也可用。

The body of the anonymous function above is just the same as the loop body was before. The only difference is that each iteration of the loop will start a new goroutine, concurrent with the current process (the `WebsiteChecker` function) each of which will add its result to the results map.

上面匿名函数的主体与之前的循环主体相同。唯一的区别是循环的每次迭代都会启动一个新的 goroutine，与当前进程（`WebsiteChecker` 函数）并发，每个进程都会将其结果添加到结果映射中。

But when we run `go test`:

但是当我们运行 `go test` 时：

```go
--- FAIL: TestCheckWebsites (0.00s)
        CheckWebsites_test.go:31: Wanted map[http://google.com:true http://blog.gypsydave5.com:true waat://furhurterwe.geds:false], got map[]
FAIL
exit status 1
FAIL    github.com/gypsydave5/learn-go-with-tests/concurrency/v1        0.010s
```

### A quick aside into a parallel(ism) universe...

### 快速进入平行（主义）宇宙......

You might not get this result. You might get a panic message that we're going to talk about in a bit. Don't worry if you got that, just keep running the test until you *do* get the result above. Or pretend that you did. Up to you. Welcome to concurrency: when it's not handled correctly it's hard to predict what's going to happen. Don't worry - that's why we're writing tests, to help us know when we're handling concurrency predictably.

你可能不会得到这个结果。您可能会收到一条我们将在稍后讨论的恐慌信息。如果你得到了不要担心，只要继续运行测试，直到你*做*得到上面的结果。或者假装你做到了。由你决定。欢迎使用并发：如果处理不当，很难预测会发生什么。别担心 - 这就是我们编写测试的原因，以帮助我们了解何时可预测地处理并发。

### ... and we're back.

### ...我们回来了。

We are caught by the original tests `CheckWebsites` is now returning an empty map. What went wrong?

我们被原始测试发现，“CheckWebsites”现在返回一个空地图。什么地方出了错？

None of the goroutines that our `for` loop started had enough time to add their result to the `results` map; the `WebsiteChecker` function is too fast for them, and it returns the still empty map.

我们的 `for` 循环启动的 goroutine 都没有足够的时间将它们的结果添加到 `results` 映射中； `WebsiteChecker` 函数对他们来说太快了，它返回仍然是空的地图。

To fix this we can just wait while all the goroutines do their work, and then return. Two seconds ought to do it, right?

为了解决这个问题，我们可以等待所有 goroutine 完成它们的工作，然后返回。两秒钟应该可以，对吧？

```go
package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

    for _, url := range urls {
        go func() {
            results[url] = wc(url)
        }()
    }

    time.Sleep(2 * time.Second)

    return results
}
```

Now when we run the tests you get (or don't get - see above):

现在，当我们运行测试时，您会得到（或没有得到 - 见上文）：

```go
--- FAIL: TestCheckWebsites (0.00s)
        CheckWebsites_test.go:31: Wanted map[http://google.com:true http://blog.gypsydave5.com:true waat://furhurterwe.geds:false], got map[waat://furhurterwe.geds:false]
FAIL
exit status 1
FAIL    github.com/gypsydave5/learn-go-with-tests/concurrency/v1        0.010s
```

This isn't great - why only one result? We might try and fix this by increasing the time we wait - try it if you like. It won't work. The problem here is that the variable `url` is reused for each iteration of the `for` loop - it takes a new value from `urls` each time. But each of our goroutines have a reference to the `url` variable - they don't have their own independent copy. So they're *all* writing the value that `url` has at the end of the iteration - the last url. Which is why the one result we have is the last url.

这不是很好 - 为什么只有一个结果？我们可能会尝试通过增加等待时间来解决此问题 - 如果您愿意，可以尝试一下。它不会工作。这里的问题是变量 `url` 在 `for` 循环的每次迭代中都被重用 - 每次都从 `urls` 中获取一个新值。但是我们的每个 goroutine 都有一个对 `url` 变量的引用——它们没有自己独立的副本。所以他们 *all* 写了 `url` 在迭代结束时的值——最后一个 url。这就是为什么我们得到的一个结果是最后一个 url。

To fix this:

要解决此问题：

```go
package concurrency

import (
    "time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

    for _, url := range urls {
        go func(u string) {
            results[u] = wc(u)
        }(url)
    }

    time.Sleep(2 * time.Second)

    return results
}
```

By giving each anonymous function a parameter for the url - `u` - and then calling the anonymous function with the `url` as the argument, we make sure that the value of `u` is fixed as the value of `url` for the iteration of the loop that we're launching the goroutine in. `u` is a copy of the value of `url`, and so can't be changed.

通过给每个匿名函数一个 url 的参数 - `u` - 然后以 `url` 作为参数调用匿名函数，我们确保 `u` 的值固定为 `url` 的值我们在其中启动 goroutine 的循环的迭代。 `u` 是 `url` 值的副本，因此无法更改。

Now if you're lucky you'll get:

现在，如果你很幸运，你会得到：

```go
PASS
ok      github.com/gypsydave5/learn-go-with-tests/concurrency/v1        2.012s
```

But if you're unlucky (this is more likely if you run them with the benchmark as you'll get more tries)

但是如果你不走运（如果你使用基准测试运行它们，这更有可能，因为你会得到更多的尝试）

```go
fatal error: concurrent map writes

goroutine 8 [running]:
runtime.throw(0x12c5895, 0x15)
        /usr/local/Cellar/go/1.9.3/libexec/src/runtime/panic.go:605 +0x95 fp=0xc420037700 sp=0xc4200376e0 pc=0x102d395
runtime.mapassign_faststr(0x1271d80, 0xc42007acf0, 0x12c6634, 0x17, 0x0)
        /usr/local/Cellar/go/1.9.3/libexec/src/runtime/hashmap_fast.go:783 +0x4f5 fp=0xc420037780 sp=0xc420037700 pc=0x100eb65
github.com/gypsydave5/learn-go-with-tests/concurrency/v3.WebsiteChecker.func1(0xc42007acf0, 0x12d3938, 0x12c6634, 0x17)
        /Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker.go:12 +0x71 fp=0xc4200377c0 sp=0xc420037780 pc=0x12308f1
runtime.goexit()
        /usr/local/Cellar/go/1.9.3/libexec/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc4200377c8 sp=0xc4200377c0 pc=0x105cf01
created by github.com/gypsydave5/learn-go-with-tests/concurrency/v3.WebsiteChecker
        /Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker.go:11 +0xa1

        ... many more scary lines of text ...
```

This is long and scary, but all we need to do is take a breath and read the stacktrace: `fatal error: concurrent map writes`. Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time. Maps in Go don't like it when more than one thing tries to write to them at once, and so `fatal error`.

这是漫长而可怕的，但我们需要做的就是深呼吸并阅读堆栈跟踪：“致命错误：并发映射写入”。有时，当我们运行测试时，两个 goroutine 会同时写入结果映射。 Go 中的 Maps 不喜欢当多个东西试图同时写入它们时，因此是“致命错误”。

This is a *race condition*, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over. Because we cannot control exactly when each goroutine writes to the results map, we are vulnerable to two goroutines writing to it at the same time.

这是一个*竞争条件*，当我们的软件输出依赖于我们无法控制的事件的时间和顺序时，就会发生一个错误。因为我们无法准确控制每个 goroutine 何时写入结果映射，所以我们容易受到两个 goroutine 同时写入它的影响。

Go can help us to spot race conditions with its built in [*race detector*](https://blog.golang.org/race-detector). To enable this feature, run the tests with the `race` flag: `go test -race`.

Go 可以通过其内置的 [*race detection*](https://blog.golang.org/race-detector) 帮助我们发现竞态条件。要启用此功能，请使用 `race` 标志运行测试：`go test -race`。

You should get some output that looks like this:

你应该得到一些看起来像这样的输出：

```go
==================
WARNING: DATA RACE
Write at 0x00c420084d20 by goroutine 8:
  runtime.mapassign_faststr()
      /usr/local/Cellar/go/1.9.3/libexec/src/runtime/hashmap_fast.go:774 +0x0
  github.com/gypsydave5/learn-go-with-tests/concurrency/v3.WebsiteChecker.func1()
      /Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker.go:12 +0x82

Previous write at 0x00c420084d20 by goroutine 7:
  runtime.mapassign_faststr()
      /usr/local/Cellar/go/1.9.3/libexec/src/runtime/hashmap_fast.go:774 +0x0
  github.com/gypsydave5/learn-go-with-tests/concurrency/v3.WebsiteChecker.func1()
      /Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker.go:12 +0x82

Goroutine 8 (running) created at:
  github.com/gypsydave5/learn-go-with-tests/concurrency/v3.WebsiteChecker()
      /Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker.go:11 +0xc4
  github.com/gypsydave5/learn-go-with-tests/concurrency/v3.TestWebsiteChecker()
      /Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker_test.go:27 +0xad
  testing.tRunner()
      /usr/local/Cellar/go/1.9.3/libexec/src/testing/testing.go:746 +0x16c

Goroutine 7 (finished) created at:
  github.com/gypsydave5/learn-go-with-tests/concurrency/v3.WebsiteChecker()
      /Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker.go:11 +0xc4
  github.com/gypsydave5/learn-go-with-tests/concurrency/v3.TestWebsiteChecker()
      /Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker_test.go:27 +0xad
  testing.tRunner()
      /usr/local/Cellar/go/1.9.3/libexec/src/testing/testing.go:746 +0x16c
==================
```

The details are, again, hard to read - but `WARNING: DATA RACE` is pretty unambiguous. Reading into the body of the error we can see two different goroutines performing writes on a map:

细节再次难以阅读 - 但“警告：数据竞赛”非常明确。读入错误正文，我们可以看到两个不同的 goroutine 在映射上执行写入操作：

```go
Write at 0x00c420084d20 by goroutine 8:
```

is writing to the same block of memory as

正在写入与相同的内存块

```go
Previous write at 0x00c420084d20 by goroutine 7:
```

On top of that we can see the line of code where the write is happening:

最重要的是，我们可以看到发生写入的代码行：

```go
/Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker.go:12
```

and the line of code where goroutines 7 an 8 are started:

以及启动 goroutine 7 和 8 的代码行：

```go
/Users/gypsydave5/go/src/github.com/gypsydave5/learn-go-with-tests/concurrency/v3/websiteChecker.go:11
```

Everything you need to know is printed to your terminal - all you have to do is be patient enough to read it.

您需要知道的一切都会打印到您的终端上——您所要做的就是有足够的耐心阅读它。

### Channels

### 频道

We can solve this data race by coordinating our goroutines using *channels*. Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.

我们可以通过使用 *channels* 协调我们的 goroutine 来解决这个数据竞争。通道是一种 Go 数据结构，可以接收和发送值。这些操作及其详细信息允许不同进程之间进行通信。

In this case we want to think about the communication between the parent process and each of the goroutines that it makes to do the work of running the `WebsiteChecker` function with the url.

在这种情况下，我们想考虑父进程和它所做的每个 goroutines 之间的通信，这些 goroutine 使用 url 来执行运行 `WebsiteChecker` 函数的工作。

```go
package concurrency

type WebsiteChecker func(string) bool
type result struct {
    string
    bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
    resultChannel := make(chan result)

    for _, url := range urls {
        go func(u string) {
            resultChannel <- result{u, wc(u)}
        }(url)
    }

    for i := 0;i < len(urls);i++ {
        r := <-resultChannel
        results[r.string] = r.bool
    }

    return results
}
```

Alongside the `results` map we now have a `resultChannel`, which we `make` in the same way. `chan result` is the type of the channel - a channel of `result`. The new type, `result` has been made to associate the return value of the `WebsiteChecker` with the url being checked - it's a struct of `string` and `bool`. As we don't need either value to be named, each of them is anonymous within the struct; this can be useful in when it's hard to know what to name a value.

除了 `results` 地图，我们现在有一个 `resultChannel`，我们以同样的方式 `make`。 `chan result` 是通道的类型 - `result` 的通道。新类型“result”已将“WebsiteChecker”的返回值与被检查的 url 相关联——它是一个由“string”和“bool”组成的结构。因为我们不需要命名任何一个值，所以它们在结构中都是匿名的；这在很难知道如何命名值时很有用。

Now when we iterate over the urls, instead of writing to the `map` directly we're sending a `result` struct for each call to `wc` to the `resultChannel` with a *send statement*. This uses the `<-` operator, taking a channel on the left and a value on the right:

现在，当我们遍历 url 时，不是直接写入 `map`，而是将每次调用 `wc` 的 `result` 结构发送到带有 *send 语句 * 的 `resultChannel`。这使用了 `<-` 运算符，在左侧取一个通道，在右侧取一个值：

```go
// Send statement
resultChannel <- result{u, wc(u)}
```

The next `for` loop iterates once for each of the urls. Inside we're using a *receive expression*, which assigns a value received from a channel to a variable. This also uses the `<-` operator, but with the two operands now reversed: the channel is now on the right and the variable that we're assigning to is on the left:

下一个 `for` 循环为每个 url 迭代一次。在内部，我们使用了一个 *receive 表达式*，它将从通道接收到的值分配给一个变量。这也使用了 `<-` 运算符，但两个操作数现在颠倒了：通道现在在右侧，我们分配给的变量在左侧：

```go
// Receive expression
r := <-resultChannel
```

We then use the `result` received to update the map.

然后我们使用收到的“result”来更新地图。

By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time. Although each of the calls of `wc`, and each send to the result channel, is happening in parallel inside its own process, each of the results is being dealt with one at a time as we take values out of the result channel with the receive expression.

通过将结果发送到一个通道，我们可以控制每次写入结果映射的时间，确保一次一个。尽管 `wc` 的每一个调用，以及每一个发送到结果通道的调用，都是在它自己的进程中并行发生的，但是当我们从结果通道中取出值时，每一个结果都会被一次处理接受表达。

We have parallelized the part of the code that we wanted to make faster, while making sure that the part that cannot happen in parallel still happens linearly. And we have communicated across the multiple processes involved by using channels.

我们已经并行化了我们想要更快的代码部分，同时确保不能并行发生的部分仍然线性发生。我们已经通过使用渠道在涉及的多个流程中进行了沟通。

When we run the benchmark:

当我们运行基准测试时：

```go
pkg: github.com/gypsydave5/learn-go-with-tests/concurrency/v2
BenchmarkCheckWebsites-8             100          23406615 ns/op
PASS
ok      github.com/gypsydave5/learn-go-with-tests/concurrency/v2        2.377s
```

23406615 nanoseconds - 0.023 seconds, about one hundred times as fast as original function. A great success.

23406615 纳秒 - 0.023 秒，大约是原始函数的一百倍。一个巨大的成功。

## Wrapping up 

##  包起来

This exercise has been a little lighter on the TDD than usual. In a way we've been taking part in one long refactoring of the `CheckWebsites` function; the inputs and outputs never changed, it just got faster. But the tests we had in place, as well as the benchmark we wrote, allowed us to refactor `CheckWebsites` in a way that maintained confidence that the software was still working, while demonstrating that it had actually become faster.

这个练习在 TDD 上比平时轻一些。在某种程度上，我们参与了对 CheckWebsites 函数的一次长期重构；输入和输出从未改变，它只是变得更快。但是我们进行的测试以及我们编写的基准测试使我们能够以一种保持软件仍在运行的信心的方式重构“CheckWebsites”，同时证明它实际上变得更快了。

In making it faster we learned about

为了让它更快，我们了解到

- *goroutines*, the basic unit of concurrency in Go, which let us check more than one website at the same time.
- *anonymous functions*, which we used to start each of the concurrent processes that check websites.
- *channels*, to help organize and control the communication between the different processes, allowing us to avoid a *race condition* bug.
- *the race detector* which helped us debug problems with concurrent code

- *goroutines*，Go 中并发的基本单位，可以让我们同时查看多个网站。
- *匿名函数*，我们用来启动检查网站的每个并发进程。
- *通道*，帮助组织和控制不同进程之间的通信，使我们能够避免*竞争条件*错误。
- *竞争检测器*帮助我们调试并发代码问题

### Make it fast

###  快一点

One formulation of an agile way of building software, often misattributed to Kent Beck, is:

构建软件的敏捷方式的一种表述通常被误认为是 Kent Beck，是：

> [Make it work, make it right, make it fast](http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast)

> [让它工作，让它正确，让它快速](http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast)

Where 'work' is making the tests pass, 'right' is refactoring the code, and 'fast' is optimizing the code to make it, for example, run quickly. We can only 'make it fast' once we've made it work and made it right. We were lucky that the code we were given was already demonstrated to be working, and didn't need to be refactored. We should never try to 'make it fast' before the other two steps have been performed because

“工作”是通过测试，“正确”是重构代码，“快速”是优化代码以使其快速运行。一旦我们让它工作并正确，我们就只能“让它快”。我们很幸运，我们得到的代码已经被证明是有效的，不需要重构。在执行其他两个步骤之前，我们永远不应该试图“让它快”，因为

> [Premature optimization is the root of all evil](http://wiki.c2.com/?PrematureOptimization) -- Donald Knuth 

> [过早优化是万恶之源](http://wiki.c2.com/?PrematureOptimization) -- Donald Knuth

