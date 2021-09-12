# Learn Go by writing tests: Synchronising asynchronous processes

# 通过编写测试学习 Go：同步异步进程

2018年4月28日 ・9 min read

This is the 8th post taken from a WIP project called [Learn Go by writing Tests](https://github.com/quii/learn-go-with-tests)the aim of which is to get a familiarity with Go and learn techniques around TDD

这是 WIP 项目的第 8 篇文章[通过编写测试学习围棋](https://github.com/quii/learn-go-with-tests)，其目的是熟悉围棋并学习围绕 TDD 的技术


- [The first post got you up to speed with TDD](https://dev.to/quii/learn-go-by-writing-tests--m63)
- [The second post discusses arrays and slices](https://dev.to/quii/learn-go-by-writing-tests-arrays-and-slices-ahm)
- [The third post teaches structs, methods, interfaces & table driven tests](https://dev.to/quii/learn-go-by-writing-tests-structs-methods-interfaces--table-driven-tests-1p01)
- [The fourth post shows how to do errors and why pointers are useful](https://dev.to/quii/learn-go-by-writing-tests-pointers-and-errors-2kp6)
- [The fifth post showed you how and why to do dependency injection](https://dev.to/quii/learn-go-by-writing-tests-dependency-injection-n7j)
- [The 6th post introduced concurrency](https://dev.to/gypsydave5/learn-go-by-writing-tests-concurrency--2ebk)
- [The 7th post shows how and why to mock](https://dev.to/quii/learn-go-by-writing-tests-mocking-fl4)


This chapter is about synchronising asynchronous processes with `select`

本章是关于用`select`同步异步进程

# Select

#  选择

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/master/select)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/master/select)**

You have been asked to make a function called `WebsiteRacer` which takes two URLs and "races" them by hitting them with a HTTP GET and returning the URL which returned first. If none of them return within 10 seconds then it should return an `error`

您被要求创建一个名为“WebsiteRacer”的函数，它接受两个 URL，并通过使用 HTTP GET 命中它们并返回第一个返回的 URL 来“竞争”它们。如果它们都没有在 10 秒内返回，那么它应该返回一个 `error`

For this we will be using

为此，我们将使用

- `net/http` to make the HTTP calls.
- `net/http/httptest` to help us test them.
- goroutines.
- `select` to synchronise processes.

- `net/http` 进行 HTTP 调用。
- `net/http/httptest` 来帮助我们测试它们。
- 协程。
- `select` 来同步进程。

### Write the test first

### 先写测试

Let's start with something naive to get us going.

让我们从一些天真的事情开始。

```
func TestRacer(t *testing.T) {
    slowURL := "http://www.facebook.com"
    fastURL := "http://www.quii.co.uk"

    want := fastURL
    got := Racer(slowURL, fastURL)

    if got != want{
        t.Errorf("got '%s', want '%s'", got, want)
    }
}
```

We know this isn't perfect and has problems but it will get us going. It's important not to get too hung-up on getting things perfect first time.

我们知道这并不完美并且有问题，但它会让我们继续前进。重要的是不要在第一次就让事情变得完美。

### Try to run the test

### 尝试运行测试

```
./racer_test.go:14:9: undefined: Racer
```

### Write the minimal amount of code for the test to run and check the failing test output

### 为测试编写最少的代码以运行并检查失败的测试输出

```
func Racer(a, b string) (winner string) {
    return
}
```

```
racer_test.go:25: got '', want 'http://www.quii.co.uk'
```

### Write enough code to make it pass

### 编写足够的代码使其通过

```
func Racer(a, b string) (winner string) {
    startA := time.Now()
    http.Get(a)
    aDuration := time.Since(startA)

    startB := time.Now()
    http.Get(b)
    bDuration := time.Since(startB)

    if aDuration < bDuration {
        return a
    }

    return b
}
```

For each url:

对于每个网址：

1. We use `time.Now()` to record just before we try and get the `URL`
2. Then we use [`http.Get`](https://golang.org/pkg/net/http/#Client.Get) to try and get the contents of the `URL`. This function returns a [`http.Response`](https://golang.org/pkg/net/http/#Response) and an `error` but so far we are not interested in these values
3. `time.Since` takes the start time and returns a `time.Duration` of the difference.

1.我们使用`time.Now()`来记录我们尝试获取`URL`之前
2.然后我们使用[`http.Get`](https://golang.org/pkg/net/http/#Client.Get)来尝试获取`URL`的内容。这个函数返回一个[`http.Response`](https://golang.org/pkg/net/http/#Response) 和一个 `error` 但到目前为止我们对这些值不感兴趣
3. `time.Since` 获取开始时间并返回差值的 `time.Duration`。

Once we have done this we simply compare the durations to see which is the quickest.

完成此操作后，我们只需比较持续时间，看看哪个是最快的。

### Problems

###  问题

This may or may not make the test pass for you. The problem is we're reaching out to real websites to test our own logic.

这可能会也可能不会让您通过测试。问题是我们正在接触真实的网站来测试我们自己的逻辑。

Testing code that uses HTTP is so common that Go has tools in the standard library to help you test it.

使用 HTTP 的测试代码非常普遍，以至于 Go 在标准库中提供了工具来帮助您测试它。

In the mocking and dependency injection chapters we covered how ideally we dont want to be relying on external services to test our code because they can be

在模拟和依赖注入章节中，我们介绍了理想情况下我们不希望依赖外部服务来测试我们的代码，因为它们可以

- Slow
- Flaky
- Can't test edge cases

- 减缓
- 片状
- 无法测试边缘情况

In the standard library there is a package [`net/http/httptest`](https://golang.org/pkg/net/http/httptest/) where you can easily create a mock HTTP server.

在标准库中有一个包 [`net/http/httptest`](https://golang.org/pkg/net/http/httptest/)，您可以在其中轻松创建模拟 HTTP 服务器。

Let's change our tests to use mocks so we have reliable servers to test against that we can control.

让我们改变我们的测试以使用模拟，这样我们就有可靠的服务器来测试我们可以控制的。

```
func TestRacer(t *testing.T) {

    slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(20 * time.Millisecond)
        w.WriteHeader(http.StatusOK)
    }))

    fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    slowURL := slowServer.URL
    fastURL := fastServer.URL

    want := fastURL
    got := Racer(slowURL, fastURL)

    if got != want {
        t.Errorf("got '%s', want '%s'", got, want)
    }

    slowServer.Close()
    fastServer.Close()
}
```

The syntax may look a bit busy but just take your time.

语法可能看起来有点忙，但请花点时间。

`httptest.NewServer` takes a `http.HandlerFunc` which we are sending in via an *anonymous function*.

`httptest.NewServer` 接受一个 `http.HandlerFunc`，我们通过一个*匿名函数*发送它。

```
http.HandlerFunc` is a type that looks like this: `type HandlerFunc func(ResponseWriter, *Request)
```

All it's really saying is it needs a function that takes a `ResponseWriter` and a `Request`, which is not too surprising for a HTTP server

它真正要说的是，它需要一个接受 `ResponseWriter` 和一个 `Request` 的函数，这对于 HTTP 服务器来说并不奇怪

It turns out there's really no extra magic here, **this is also how you would write a \*real\* HTTP server in Go**. The only difference is we are wrapping it in a `httptest.NewServer` which makes it easier to use with testing, as it finds an open port to listen on and then you can close it when you're done with your test.

事实证明这里真的没有额外的魔法，**这也是你在 Go 中编写 \*real\* HTTP 服务器的方式**。唯一的区别是我们将它包装在一个 `httptest.NewServer` 中，这使得它更容易用于测试，因为它会找到一个开放的端口来监听，然后你可以在完成测试后关闭它。

Inside our two servers we make the slow one have a short `time.Sleep` when we get a request to make it slower than the other one. Both servers then write an `OK` response with `w.WriteHeader(http.StatusOK)` back to the caller.

在我们的两台服务器中，当我们收到要求使其比另一台慢的请求时，我们让慢的一台有很短的 `time.Sleep`。然后，两个服务器都将带有 `w.WriteHeader(http.StatusOK)` 的 `OK` 响应写回给调用者。

If you re-run the test it will definitely pass now and should be faster. Play with these sleeps to deliberately break the test.

如果您重新运行测试，它现在肯定会通过并且应该更快。玩这些睡眠故意打破测试。

### Refactor

### 重构

We have some duplication in both our production code and test code.

我们的生产代码和测试代码都有一些重复。

```
func Racer(a, b string) (winner string) {
    aDuration := measureResponseTime(a)
    bDuration := measureResponseTime(b)

    if aDuration < bDuration {
        return a
    }

    return b
}

func measureResponseTime(url string) time.Duration {
    start := time.Now()
    http.Get(url)
    return time.Since(start)
}
```

This DRY-ing up makes our `Racer` code a lot easier to read.

这种干燥使我们的 `Racer` 代码更容易阅读。

```
func TestRacer(t *testing.T) {

    slowServer := makeDelayedServer(20 * time.Millisecond)
    fastServer := makeDelayedServer(0 * time.Millisecond)

    defer slowServer.Close()
    defer fastServer.Close()

    slowURL := slowServer.URL
    fastURL := fastServer.URL

    want := fastURL
    got := Racer(slowURL, fastURL)

    if got != want {
        t.Errorf("got '%s', want '%s'", got, want)
    }
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}
```

We've refactored creating our fake servers into a function called `makeDelayedServer` to move some uninteresting code out of the test and reduce repetition.

我们已经将创建我们的假服务器重构为一个名为 `makeDelayedServer` 的函数，以将一些无趣的代码从测试中移出并减少重复。

#### `defer`

#### `延迟`

By prefixing a function call with `defer` it will now call that function *at the end of the containing function*.

通过在函数调用前加上 `defer` 前缀，它现在将在包含函数的末尾调用该函数*。

Sometimes you will need to cleanup resources, such as closing a file or in our case closing a server so that it does not continue to listen to a port.

有时您需要清理资源，例如关闭文件或在我们的情况下关闭服务器，以便它不会继续侦听端口。

You want this to execute at the end of the function, but keep the instruction near where you created the server for the benefit of future readers of the code.

您希望它在函数结束时执行，但将指令保留在您创建服务器的位置附近，以便将来阅读代码的读者受益。

Our refactoring is an improvement and is a reasonable solution given the Go features covered so far, but we can make the solution simpler.

我们的重构是一种改进，考虑到迄今为止涵盖的 Go 特性，这是一个合理的解决方案，但我们可以使解决方案更简单。

### Synchronising processes

### 同步进程

- Why are we testing the speeds of the websites one after another when Go is great at concurrency? We should be able to check both at the same time
- We don't really care about *the exact response times* of the requests, we just want to know which one comes back first.

- 当 Go 的并发性很强时，为什么我们要一个个测试网站的速度？我们应该能够同时检查两者
- 我们并不真正关心请求的*确切响应时间*，我们只想知道哪个首先返回。

To do this, we're going to introduce a new construct called `select` which helps us synchronise processes really easily and clearly.

为此，我们将引入一个名为“select”的新结构，它可以帮助我们真正轻松、清晰地同步进程。

```
func Racer(a, b string) (winner string) {
    select {
    case <-ping(a):
        return a
    case <-ping(b):
        return b
    }
}

func ping(url string) chan bool {
    ch := make(chan bool)
    go func() {
        http.Get(url)
        ch <- true
    }()
    return ch
}
```

#### `ping`

#### `ping`

We have defined a function `ping` which creates a `chan bool` and returns it.

我们定义了一个函数`ping`，它创建一个`chan bool`并返回它。

In our case, we don't really *care* what the type sent in the channel, *we just want to send a signal* to say we're finished so booleans are fine.

在我们的例子中，我们并不真正*关心*通道中发送的类型，*我们只想发送一个信号*来表示我们已经完成，所以布尔值很好。

Inside the same function we start a goroutine which will send a signal into that channel once we have completed `http.Get(url)`

在同一个函数中，我们启动了一个 goroutine，一旦我们完成了 `http.Get(url)`，它将向该通道发送一个信号

#### `select`

#### `选择`

If you recall from the concurrency chapter, you can wait for values to be sent to a channel with `myVar := <-ch`. This is a *blocking* call, as you're waiting for a value.

如果你回忆一下并发章节，你可以使用 `myVar := <-ch` 等待值被发送到通道。这是一个 *blocking* 调用，因为您正在等待一个值。

What `select` lets you do is wait on *multiple* channels. The first one to send a value "wins" and the code underneath the `case` is executed. 

`select` 可以让你在*多个*频道上等待。第一个发送值“wins”并且执行`case`下面的代码。

We use `ping` in our `select` to set up two channels for each of our `URL`s. Whichever one writes to its channel first will have its code executed in the `select`, which results in its `URL` being returned (and being the winner).

我们在 `select` 中使用 `ping` 为每个 `URL` 设置两个通道。无论哪个首先写入其频道，其代码都会在“select”中执行，这将导致其“URL”被返回（并成为赢家）。

After these changes the intent behind our code is very clear and the implementation is actually simpler.

在这些更改之后，我们代码背后的意图非常清晰，实现实际上更简单了。

### Timeouts

### 超时

Our final requirement was to return an error if `Racer` takes longer than 10 seconds.

我们的最终要求是如果“Racer”花费的时间超过 10 秒，则返回错误。

### Write the test first

### 先写测试

```
t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
    serverA := makeDelayedServer(11 * time.Second)
    serverB := makeDelayedServer(12 * time.Second)

    defer serverA.Close()
    defer serverB.Close()

    _, err := Racer(serverA.URL, serverB.URL)

    if err == nil {
        t.Error("expected an error but didn't get one")
    }
})
```

We've made our test servers take longer than 10s to return to exercise this scenario and we are expecting `Racer` to return two values now, the winning URL (which we ignore in this test with `_`) and an `error` .

我们已经让我们的测试服务器花费超过 10 秒的时间来返回来执行这个场景，我们现在期望 `Racer` 返回两个值，获胜的 URL（我们在这个测试中用 `_` 忽略）和一个 `error` .

### Try to run the test

### 尝试运行测试

```
./racer_test.go:37:10: assignment mismatch: 2 variables but 1 values
```

### Write the minimal amount of code for the test to run and check the failing test output

### 为测试编写最少的代码以运行并检查失败的测试输出

```
func Racer(a, b string) (winner string, error error) {
    select {
    case <-ping(a):
        return a, nil
    case <-ping(b):
        return b, nil
    }
}
```

Change the signature of `Racer` to return the winner and an `error`. Return `nil` for our happy cases.

更改“Racer”的签名以返回获胜者和“错误”。为我们的快乐案例返回 `nil`。

The compiler will complain about your *first test* only looking for one value so change that line to `got, _ := Racer(slowURL, fastURL)`, knowing that we should check we *don't* get an error in our happy scenario.

编译器会抱怨你的 *first test* 只寻找一个值，所以将该行更改为 `got, _ := Racer(slowURL, fastURL)`，知道我们应该检查我们*不*在我们的快乐的场景。

If you run it now after 11 seconds it will fail

如果您在 11 秒后立即运行它，它将失败

```
-------- FAIL: TestRacer (12.00s)
    --- FAIL: TestRacer/returns_an_error_if_a_server_doesn't_respond_within_10s (12.00s)
        racer_test.go:40: expected an error but didn't get one
```

### Write enough code to make it pass

### 编写足够的代码使其通过

```
func Racer(a, b string) (winner string, error error) {
    select {
    case <-ping(a):
        return a, nil
    case <-ping(b):
        return b, nil
    case <-time.After(10 * time.Second):
        return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
    }
}
```

`time.After` is a very handy function when using `select`. Although it didn't happen in our case you can potentially write code that blocks forever if the channels you're listening on never return a value. `time.After` returns a `chan`(like `ping`) and will send a signal down it after the amount of time you define.

`time.After` 是使用 `select` 时非常方便的函数。尽管在我们的案例中没有发生这种情况，但如果您正在侦听的频道从不返回值，您可能会编写永远阻塞的代码。 `time.After` 返回一个 `chan`（如 `ping`），并在你定义的时间后发送一个信号。

For us this is perfect; if `a` or `b` manage to return they win, but if we get to 10 seconds then our `time.After` will send a signal and we'll return an `error`

对我们来说，这是完美的；如果 `a` 或 `b` 设法返回他们获胜，但如果我们达到 10 秒，那么我们的 `time.After` 将发送一个信号，我们将返回一个 `error`

### Slow tests

### 慢测试

The problem we have is that this test takes 10 seconds to run. For such a simple bit of logic this doesn't feel great.

我们遇到的问题是此测试需要 10 秒才能运行。对于如此简单的逻辑，这感觉并不好。

What we can do is make the timeout configurable so in our test we can have a very short timeout and then when the code is used in the real world it can be set to 10 seconds.

我们可以做的是使超时可配置，因此在我们的测试中我们可以有一个非常短的超时，然后当代码在现实世界中使用时，它可以设置为 10 秒。

```
func Racer(a, b string, timeout time.Duration) (winner string, error error) {
    select {
    case <-ping(a):
        return a, nil
    case <-ping(b):
        return b, nil
    case <-time.After(timeout):
        return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
    }
}
```

Our tests now wont compile because we're not supplying a timeout

我们的测试现在无法编译，因为我们没有提供超时

Before rushing in to add this default value to both our tests let's *listen to them*.

在急于将这个默认值添加到我们的两个测试之前，让我们*听听他们*。

- Do we care about the timeout in the "happy" test?
- The requirements were explicit about the timeout

- 我们是否关心“快乐”测试中的超时时间？
- 要求明确了超时

Given this knowledge, let's do a little refactoring to be sympathetic to both our tests and the users of our code

鉴于这些知识，让我们进行一些重构，以对我们的测试和我们代码的用户表示同情

```
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
    return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
    select {
    case <-ping(a):
        return a, nil
    case <-ping(b):
        return b, nil
    case <-time.After(timeout):
        return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
    }
}
```

Our users and our first test can use `Racer` (which uses `ConfigurableRacer`under the hood) and our sad path test can use `ConfigurableRacer`.

我们的用户和我们的第一个测试可以使用 `Racer`（在引擎盖下使用 `ConfigurableRacer`），而我们的悲伤路径测试可以使用 `ConfigurableRacer`。

```
func TestRacer(t *testing.T) {

    t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
        slowServer := makeDelayedServer(20 * time.Millisecond)
        fastServer := makeDelayedServer(0 * time.Millisecond)

        defer slowServer.Close()
        defer fastServer.Close()

        slowURL := slowServer.URL
        fastURL := fastServer.URL

        want := fastURL
        got, err := Racer(slowURL, fastURL)

        if err != nil {
            t.Fatalf("did not expect an error but got one %v", err)
        }

        if got != want {
            t.Errorf("got '%s', want '%s'", got, want)
        }
    })

    t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
        server := makeDelayedServer(25 * time.Millisecond)

        defer server.Close()

        _, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

        if err == nil {
            t.Error("expected an error but didn't get one")
        }
    })
}
```

I added one final check on the first test to verify we don't get an `error`

我在第一次测试中添加了一个最终检查，以验证我们没有收到“错误”

### Wrapping up

###  总结

### `select`


- Helps you wait on multiple channels.
- Sometimes you'll want to include `time.After` in one of your `cases` to prevent your system blocking forever.

- 帮助您等待多个频道。
- 有时你会想要在你的一个 `case` 中包含 `time.After` 以防止你的系统永远阻塞。

### `httptest`

- Convenient way of creating test servers so you can have reliable and controllable tests.
- Uses the same interfaces as the "real" `net/http` servers which is consistent and less for you to learn 

- 创建测试服务器的便捷方式，因此您可以进行可靠且可控的测试。
- 使用与“真正的”`net/http` 服务器相同的接口，这是一致的，你学习的更少

