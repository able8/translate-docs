# Context

#  语境

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/context)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/context)**

Software often kicks off long-running, resource-intensive processes (often in goroutines). If the action that caused this gets cancelled or fails for some reason you need to stop these processes in a consistent way through your application.

软件通常会启动长时间运行的资源密集型进程（通常在 goroutines 中）。如果导致此问题的操作因某种原因被取消或失败，您需要通过应用程序以一致的方式停止这些进程。

If you don't manage this your snappy Go application that you're so proud of could start having difficult to debug performance problems.

如果您不管理它，您引以为豪的活泼 Go 应用程序可能会开始难以调试性能问题。

In this chapter we'll use the package `context` to help us manage long-running processes.

在本章中，我们将使用包 `context` 来帮助我们管理长时间运行的进程。

We're going to start with a classic example of a web server that when hit kicks off a potentially long-running process to fetch some data for it to return in the response.

我们将从一个经典的 Web 服务器示例开始，当点击该服务器时，它会启动一个潜在的长时间运行的进程，以获取一些数据以使其在响应中返回。

We will exercise a scenario where a user cancels the request before the data can be retrieved and we'll make sure the process is told to give up.

我们将练习一个场景，即用户在可以检索数据之前取消请求，并且我们将确保该过程被告知放弃。

I've set up some code on the happy path to get us started. Here is our server code.

我已经在快乐路径上设置了一些代码来让我们开始。这是我们的服务器代码。

```go
func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, store.Fetch())
    }
}
```

The function `Server` takes a `Store` and returns us a `http.HandlerFunc`. Store is defined as:

函数`Server` 接受一个`Store` 并返回一个`http.HandlerFunc`。商店定义为：

```go
type Store interface {
    Fetch() string
}
```

The returned function calls the `store`'s `Fetch` method to get the data and writes it to the response.

返回的函数调用`store` 的`Fetch` 方法来获取数据并将其写入响应。

We have a corresponding stub for `Store` which we use in a test.

我们有一个对应的“Store”存根，我们在测试中使用它。

```go
type StubStore struct {
    response string
}

func (s *StubStore) Fetch() string {
    return s.response
}

func TestServer(t *testing.T) {
    data := "hello, world"
    svr := Server(&StubStore{data})

    request := httptest.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    svr.ServeHTTP(response, request)

    if response.Body.String() != data {
        t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
    }
}
```

Now that we have a happy path, we want to make a more realistic scenario where the `Store` can't finish a`Fetch` before the user cancels the request.

现在我们有了一个快乐的路径，我们想要制作一个更现实的场景，在用户取消请求之前，`Store` 无法完成 `Fetch`。

## Write the test first

## 先写测试

Our handler will need a way of telling the `Store` to cancel the work so update the interface.

我们的处理程序需要一种方法来告诉 `Store` 取消工作，以便更新界面。

```go
type Store interface {
    Fetch() string
    Cancel()
}
```

We will need to adjust our spy so it takes some time to return `data` and a way of knowing it has been told to cancel. We'll also rename it to `SpyStore` as we are now observing the way it is called. It'll have to add `Cancel` as a method to implement the `Store` interface.

我们需要调整我们的 spy，以便它需要一些时间来返回 `data` 并知道它已被告知取消。我们还将它重命名为 `SpyStore`，因为我们现在正在观察它的调用方式。它必须添加 `Cancel` 作为实现 `Store` 接口的方法。

```go
type SpyStore struct {
    response string
    cancelled bool
}

func (s *SpyStore) Fetch() string {
    time.Sleep(100 * time.Millisecond)
    return s.response
}

func (s *SpyStore) Cancel() {
    s.cancelled = true
}
```

Let's add a new test where we cancel the request before 100 milliseconds and check the store to see if it gets cancelled.

让我们添加一个新测试，我们在 100 毫秒之前取消请求并检查存储以查看它是否被取消。

```go
t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
      data := "hello, world"
      store := &SpyStore{response: data}
      svr := Server(store)

      request := httptest.NewRequest(http.MethodGet, "/", nil)

      cancellingCtx, cancel := context.WithCancel(request.Context())
      time.AfterFunc(5 * time.Millisecond, cancel)
      request = request.WithContext(cancellingCtx)

      response := httptest.NewRecorder()

      svr.ServeHTTP(response, request)

      if !store.cancelled {
          t.Error("store was not told to cancel")
      }
  })
```

From the [Go Blog: Context](https://blog.golang.org/context)

来自 [Go 博客：上下文](https://blog.golang.org/context)

> The context package provides functions to derive new Context values from existing ones. These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.

> 上下文包提供了从现有值派生新上下文值的函数。这些值形成一棵树：当一个 Context 被取消时，所有从它派生的 Context 也被取消。

It's important that you derive your contexts so that cancellations are propagated throughout the call stack for a given request.

派生上下文很重要，以便取消在给定请求的整个调用堆栈中传播。

What we do is derive a new `cancellingCtx` from our `request` which returns us a `cancel` function. We then schedule that function to be called in 5 milliseconds by using `time.AfterFunc`. Finally we use this new context in our request by calling `request.WithContext`.

我们所做的是从我们的 `request` 派生出一个新的 `cancellingCtx`，它返回一个 `cancel` 函数。然后我们使用`time.AfterFunc`安排在5毫秒内调用该函数。最后，我们通过调用`request.WithContext` 在我们的请求中使用这个新的上下文。

## Try to run the test

## 尝试运行测试

The test fails as we'd expect.

正如我们所料，测试失败了。

```go
--- FAIL: TestServer (0.00s)
    --- FAIL: TestServer/tells_store_to_cancel_work_if_request_is_cancelled (0.00s)
        context_test.go:62: store was not told to cancel
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Remember to be disciplined with TDD. Write the _minimal_ amount of code to make our test pass.

记住要遵守 TDD。编写 _minimal_ 代码量以使我们的测试通过。

```go
func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        store.Cancel()
        fmt.Fprint(w, store.Fetch())
    }
}
```

This makes this test pass but it doesn't feel good does it! We surely shouldn't be cancelling `Store` before we fetch on _every request_.

这使得这个测试通过了，但感觉不太好，是吗！我们当然不应该在获取_每个请求_之前取消`Store`。

By being disciplined it highlighted a flaw in our tests, this is a good thing!

受到纪律处分，它突出了我们测试中的一个缺陷，这是一件好事！

We'll need to update our happy path test to assert that it does not get cancelled.

我们需要更新我们的快乐路径测试以断言它不会被取消。

```go
t.Run("returns data from store", func(t *testing.T) {
    data := "hello, world"
    store := &SpyStore{response: data}
    svr := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    svr.ServeHTTP(response, request)

    if response.Body.String() != data {
        t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
    }

    if store.cancelled {
        t.Error("it should not have cancelled the store")
    }
})
```

Run both tests and the happy path test should now be failing and now we're forced to do a more sensible implementation.

运行两个测试，快乐路径测试现在应该失败了，现在我们被迫做一个更明智的实现。

```go
func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

        data := make(chan string, 1)

        go func() {
            data <- store.Fetch()
        }()

        select {
        case d := <-data:
            fmt.Fprint(w, d)
        case <-ctx.Done():
            store.Cancel()
        }
    }
}
```

What have we done here?

我们在这里做了什么？

`context` has a method `Done()` which returns a channel which gets sent a signal when the context is "done" or "cancelled". We want to listen to that signal and call `store.Cancel` if we get it but we want to ignore it if our `Store` manages to `Fetch` before it.

`context` 有一个方法 `Done()`，它返回一个通道，当上下文“完成”或“取消”时，该通道被发送一个信号。我们想听那个信号，如果我们得到它就调用`store.Cancel`，但是如果我们的`Store` 设法在它之前进行`Fetch`，我们想忽略它。

To manage this we run `Fetch` in a goroutine and it will write the result into a new channel `data`. We then use `select` to effectively race to the two asynchronous processes and then we either write a response or `Cancel`.

为了管理这个，我们在 goroutine 中运行 `Fetch`，它将结果写入一个新的通道 `data`。然后我们使用“select”来有效地竞争两个异步进程，然后我们要么写一个响应要么“取消”。

## Refactor

## 重构

We can refactor our test code a bit by making assertion methods on our spy

我们可以通过在我们的 spy 上创建断言方法来稍微重构我们的测试代码

```go
type SpyStore struct {
    response  string
    cancelled bool
    t         *testing.T
}

func (s *SpyStore) assertWasCancelled() {
    s.t.Helper()
    if !s.cancelled {
        s.t.Error("store was not told to cancel")
    }
}

func (s *SpyStore) assertWasNotCancelled() {
    s.t.Helper()
    if s.cancelled {
        s.t.Error("store was told to cancel")
    }
}
```

Remember to pass in the `*testing.T` when creating the spy.

记得在创建 spy 时传入 `*testing.T`。

```go
func TestServer(t *testing.T) {
    data := "hello, world"

    t.Run("returns data from store", func(t *testing.T) {
        store := &SpyStore{response: data, t: t}
        svr := Server(store)

        request := httptest.NewRequest(http.MethodGet, "/", nil)
        response := httptest.NewRecorder()

        svr.ServeHTTP(response, request)

        if response.Body.String() != data {
            t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
        }

        store.assertWasNotCancelled()
    })

    t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
        store := &SpyStore{response: data, t: t}
        svr := Server(store)

        request := httptest.NewRequest(http.MethodGet, "/", nil)

        cancellingCtx, cancel := context.WithCancel(request.Context())
        time.AfterFunc(5*time.Millisecond, cancel)
        request = request.WithContext(cancellingCtx)

        response := httptest.NewRecorder()

        svr.ServeHTTP(response, request)

        store.assertWasCancelled()
    })
}
```

This approach is ok, but is it idiomatic?

这种方法没问题，但它是惯用的吗？

Does it make sense for our web server to be concerned with manually cancelling `Store`? What if `Store` also happens to depend on other slow-running processes? We'll have to make sure that `Store.Cancel` correctly propagates the cancellation to all of its dependants.

我们的 Web 服务器关注手动取消“Store”是否有意义？如果`Store` 也恰好依赖于其他运行缓慢的进程怎么办？我们必须确保`Store.Cancel` 将取消正确地传播到它的所有依赖项。

One of the main points of `context` is that it is a consistent way of offering cancellation.

`context` 的要点之一是它是提供取消的一致方式。

[From the go doc](https://golang.org/pkg/context/)

[来自 go 文档](https://golang.org/pkg/context/)

> Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.

> 对服务器的传入请求应创建上下文，对服务器的传出调用应接受上下文。它们之间的函数调用链必须传播 Context，可以选择将其替换为使用 WithCancel、WithDeadline、WithTimeout 或 WithValue 创建的派生 Context。当一个上下文被取消时，从它派生的所有上下文也被取消。

From the [Go Blog: Context](https://blog.golang.org/context) again: 

再次来自 [Go Blog: Context](https://blog.golang.org/context)：

> At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly.

> 在 Google，我们要求 Go 程序员将 Context 参数作为第一个参数传递给传入和传出请求之间调用路径上的每个函数。这使得许多不同团队开发的 Go 代码可以很好地互操作。它提供了对超时和取消的简单控制，并确保安全凭证等关键值正确传输 Go 程序。

(Pause for a moment and think of the ramifications of every function having to send in a context, and the ergonomics of that.)

（暂停片刻，想一想必须在上下文中发送的每个函数的后果，以及它的人体工程学。）

Feeling a bit uneasy? Good. Let's try and follow that approach though and instead pass through the `context` to our `Store` and let it be responsible. That way it can also pass the `context` through to its dependants and they too can be responsible for stopping themselves.

感觉有点不安？好的。让我们尝试遵循这种方法，而是将 `context` 传递给我们的 `Store` 并让它负责。这样它也可以将“上下文”传递给它的依赖者，他们也可以负责停止自己。

## Write the test first

## 先写测试

We'll have to change our existing tests as their responsibilities are changing. The only thing our handler is responsible for now is making sure it sends a context through to the downstream `Store` and that it handles the error that will come from the `Store` when it is cancelled.

我们将不得不改变我们现有的测试，因为他们的职责正在发生变化。我们的处理程序现在唯一负责的是确保它将上下文发送到下游的“Store”，并处理“Store”被取消时来自“Store”的错误。

Let's update our `Store` interface to show the new responsibilities.

让我们更新我们的 `Store` 界面以显示新的职责。

```go
type Store interface {
    Fetch(ctx context.Context) (string, error)
}
```

Delete the code inside our handler for now

暂时删除我们处理程序中的代码

```go
func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
    }
}
```

Update our `SpyStore`

更新我们的`SpyStore`

```go
type SpyStore struct {
    response string
    t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
    data := make(chan string, 1)

    go func() {
        var result string
        for _, c := range s.response {
            select {
            case <-ctx.Done():
                s.t.Log("spy store got cancelled")
                return
            default:
                time.Sleep(10 * time.Millisecond)
                result += string(c)
            }
        }
        data <- result
    }()

    select {
    case <-ctx.Done():
        return "", ctx.Err()
    case res := <-data:
        return res, nil
    }
}
```

We have to make our spy act like a real method that works with `context`.

我们必须让我们的 spy 表现得像一个与 `context` 一起工作的真实方法。

We are simulating a slow process where we build the result slowly by appending the string, character by character in a goroutine. When the goroutine finishes its work it writes the string to the `data` channel. The goroutine listens for the `ctx.Done` and will stop the work if a signal is sent in that channel.

我们正在模拟一个缓慢的过程，我们通过在 goroutine 中逐个字符地附加字符串来缓慢地构建结果。当 goroutine 完成其工作时，它将字符串写入“数据”通道。 goroutine 会监听 `ctx.Done`，如果在该通道中发送了信号，它将停止工作。

Finally the code uses another `select` to wait for that goroutine to finish its work or for the cancellation to occur.

最后，代码使用另一个 `select` 来等待 goroutine 完成其工作或取消发生。

It's similar to our approach from before, we use Go's concurrency primitives to make two asynchronous processes race each other to determine what we return.

它类似于我们之前的方法，我们使用 Go 的并发原语让两个异步进程相互竞争以确定我们返回的内容。

You'll take a similar approach when writing your own functions and methods that accept a `context` so make sure you understand what's going on.

在编写自己的接受“上下文”的函数和方法时，您将采用类似的方法，因此请确保您了解正在发生的事情。

Finally we can update our tests. Comment out our cancellation test so we can fix the happy path test first.

最后，我们可以更新我们的测试。注释掉我们的取消测试，以便我们可以先修复快乐路径测试。

```go
t.Run("returns data from store", func(t *testing.T) {
    data := "hello, world"
    store := &SpyStore{response: data, t: t}
    svr := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    svr.ServeHTTP(response, request)

    if response.Body.String() != data {
        t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
    }
})
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestServer/returns_data_from_store
--- FAIL: TestServer (0.00s)
    --- FAIL: TestServer/returns_data_from_store (0.00s)
        context_test.go:22: got "", want "hello, world"
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data, _ := store.Fetch(r.Context())
        fmt.Fprint(w, data)
    }
}
```

Our happy path should be... happy. Now we can fix the other test.

我们的幸福之路应该是……幸福的。现在我们可以修复另一个测试。

## Write the test first

## 先写测试

We need to test that we do not write any kind of response on the error case. Sadly `httptest.ResponseRecorder` doesn't have a way of figuring this out so we'll have to roll our own spy to test for this.

我们需要测试我们没有对错误情况写任何类型的响应。遗憾的是，`httptest.ResponseRecorder` 没有办法解决这个问题，所以我们必须使用我们自己的间谍来测试这一点。

```go
type SpyResponseWriter struct {
    written bool
}

func (s *SpyResponseWriter) Header() http.Header {
    s.written = true
    return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
    s.written = true
    return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
    s.written = true
}
```

Our `SpyResponseWriter` implements `http.ResponseWriter` so we can use it in the test.

我们的 `SpyResponseWriter` 实现了 `http.ResponseWriter`，因此我们可以在测试中使用它。

```go
t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
    store := &SpyStore{response: data, t: t}
    svr := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)

    cancellingCtx, cancel := context.WithCancel(request.Context())
    time.AfterFunc(5*time.Millisecond, cancel)
    request = request.WithContext(cancellingCtx)

    response := &SpyResponseWriter{}

    svr.ServeHTTP(response, request)

    if response.written {
        t.Error("a response should not have been written")
    }
})
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestServer
=== RUN   TestServer/tells_store_to_cancel_work_if_request_is_cancelled
--- FAIL: TestServer (0.01s)
    --- FAIL: TestServer/tells_store_to_cancel_work_if_request_is_cancelled (0.01s)
        context_test.go:47: a response should not have been written
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data, err := store.Fetch(r.Context())

        if err != nil {
            return // todo: log error however you like
        }

        fmt.Fprint(w, data)
    }
}
```

We can see after this that the server code has become simplified as it's no longer explicitly responsible for cancellation, it simply passes through `context` and relies on the downstream functions to respect any cancellations that may occur.

在此之后我们可以看到服务器代码已经变得简化，因为它不再明确负责取消，它只是通过 `context` 并依赖下游函数来考虑可能发生的任何取消。

## Wrapping up

##  总结

### What we've covered

### 我们已经介绍过的内容

- How to test a HTTP handler that has had the request cancelled by the client.
- How to use context to manage cancellation.
- How to write a function that accepts `context` and uses it to cancel itself by using goroutines, `select` and channels.
- Follow Google's guidelines as to how to manage cancellation by propagating request scoped context through your call-stack.
- How to roll your own spy for `http.ResponseWriter` if you need it.

- 如何测试已被客户端取消请求的 HTTP 处理程序。
- 如何使用上下文来管理取消。
- 如何编写一个接受 `context` 并通过使用 goroutines、`select` 和通道使用它来取消自身的函数。
- 遵循 Google 关于如何通过调用堆栈传播请求范围上下文来管理取消的指南。
- 如果需要，如何为 `http.ResponseWriter` 部署你自己的间谍。

### What about context.Value ?

### context.Value 怎么样？

[Michal Štrba](https://faiface.github.io/post/context-should-go-away-go2/) and I have a similar opinion.

[Michal Štrba](https://faiface.github.io/post/context-should-go-away-go2/) 我也有类似的看法。

> If you use ctx.Value in my (non-existent) company, you’re fired

> 如果你在我（不存在的）公司使用 ctx.Value，你会被解雇

Some engineers have advocated passing values through `context` as it _feels convenient_.

一些工程师提倡通过“上下文”传递值，因为它_感觉方便_。

Convenience is often the cause of bad code.

方便通常是糟糕代码的原因。

The problem with `context.Values` is that it's just an untyped map so you have no type-safety and you have to handle it not actually containing your value. You have to create a coupling of map keys from one module to another and if someone changes something things start breaking.

`context.Values` 的问题在于它只是一个无类型映射，所以你没有类型安全，你必须处理它实际上不包含你的值。您必须创建从一个模块到另一个模块的映射键的耦合，如果有人更改某些内容，事情就会开始破裂。

In short, **if a function needs some values, put them as typed parameters rather than trying to fetch them from `context.Value`**. This makes it statically checked and documented for everyone to see.

简而言之，**如果函数需要一些值，请将它们作为类型参数而不是尝试从 `context.Value` 中获取它们**。这使得每个人都可以看到它的静态检查和记录。

#### But...

####  但...

On other hand, it can be helpful to include information that is orthogonal to a request in a context, such as a trace id. Potentially this information would not be needed by every function in your call-stack and would make your functional signatures very messy.

另一方面，在上下文中包含与请求正交的信息（例如跟踪 ID）可能会有所帮助。潜在地，调用堆栈中的每个函数都不需要此信息，并且会使您的函数签名非常混乱。

[Jack Lindamood says **Context.Value should inform, not control**](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)

[Jack Lindamood 说 **Context.Value 应该通知，而不是控制**](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)

> The content of context.Value is for maintainers not users. It should never be required input for documented or expected results.

> context.Value 的内容是为维护者而不是用户提供的。永远不应要求对记录的或预期的结果进行输入。

### Additional material

### 附加材料

- I really enjoyed reading [Context should go away for Go 2 by Michal Štrba](https://faiface.github.io/post/context-should-go-away-go2/). His argument is that having to pass `context` everywhere is a smell, that it's pointing to a deficiency in the language in respect to cancellation. He says it would better if this was somehow solved at the language level, rather than at a library level. Until that happens, you will need `context` if you want to manage long running processes.
- The [Go blog further describes the motivation for working with `context` and has some examples](https://blog.golang.org/context) 

- 我真的很喜欢阅读 [Context should go away for Go 2 by Michal Štrba](https://faiface.github.io/post/context-should-go-away-go2/)。他的论点是，必须在任何地方传递“上下文”是一种气味，这表明语言在取消方面存在缺陷。他说，如果在语言级别而不是在库级别以某种方式解决这个问题会更好。在此之前，如果您想管理长时间运行的进程，您将需要 `context`。
- [Go 博客进一步描述了使用 `context` 的动机并提供了一些示例](https://blog.golang.org/context)

