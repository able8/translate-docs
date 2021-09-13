# HTTP Server

# HTTP 服务器

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/http-server)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/http-server)**

You have been asked to create a web server where users can track how many games players have won.

您被要求创建一个 Web 服务器，用户可以在其中跟踪玩家赢得的游戏数量。

-   `GET /players/{name}` should return a number indicating the total number of wins
-   `POST /players/{name}` should record a win for that name, incrementing for every subsequent `POST`

- `GET /players/{name}` 应该返回一个数字，表示获胜的总数
- `POST /players/{name}` 应该为那个名字记录一个胜利，每个后续的 `POST` 都会增加

We will follow the TDD approach, getting working software as quickly as we can and then making small iterative improvements until we have the solution. By taking this approach we

我们将遵循 TDD 方法，尽快获得可工作的软件，然后进行小的迭代改进，直到我们找到解决方案。通过采取这种方法，我们

-   Keep the problem space small at any given time
-   Don't go down rabbit holes
-   If we ever get stuck/lost, doing a revert wouldn't lose loads of work.

- 在任何给定时间保持问题空间很小
- 不要去兔子洞
- 如果我们被卡住/迷路了，进行还原不会丢失大量工作。

## Red, green, refactor

## 红色，绿色，重构

Throughout this book, we have emphasised the TDD process of write a test & watch it fail (red), write the _minimal_ amount of code to make it work (green) and then refactor.

在整本书中，我们强调了编写测试并观察它失败（红色）、编写_最少_数量的代码使其工作（绿色）然后重构的 TDD 过程。

This discipline of writing the minimal amount of code is important in terms of the safety TDD gives you. You should be striving to get out of "red" as soon as you can.

就TDD 为您提供的安全性而言，这种编写最少代码的纪律很重要。你应该努力尽快摆脱“红色”。

Kent Beck describes it as:

肯特贝克将其描述为：

> Make the test work quickly, committing whatever sins necessary in process.

> 使测试快速进行，在过程中犯下任何必要的错误。

You can commit these sins because you will refactor afterwards backed by the safety of the tests.

你可以犯这些罪，因为你会在测试的安全性的支持下进行重构。

### What if you don't do this?

### 如果你不这样做呢？

The more changes you make while in red, the more likely you are to add more problems, not covered by tests.

您在红色期间所做的更改越多，您就越有可能添加更多测试未涵盖的问题。

The idea is to be iteratively writing useful code with small steps, driven by tests so that you don't fall into a rabbit hole for hours.

这个想法是用小步骤迭代地编写有用的代码，由测试驱动，这样你就不会在几个小时内陷入困境。

### Chicken and egg

### 鸡肉和鸡蛋

How can we incrementally build this? We can't `GET` a player without having stored something and it seems hard to know if `POST` has worked without the `GET` endpoint already existing.

我们如何逐步构建它？我们不能在没有存储一些东西的情况下“GET”玩家，而且如果没有已经存在的“GET”端点，似乎很难知道“POST”是否有效。

This is where _mocking_ shines.

这就是 _mocking_ 闪耀的地方。

-   `GET` will need a `PlayerStore` _thing_ to get scores for a player. This should be an interface so when we test we can create a simple stub to test our code without needing to have implemented any actual storage code.
-   For `POST` we can _spy_ on its calls to `PlayerStore` to make sure it stores players correctly. Our implementation of saving won't be coupled to retrieval.
-   For having some working software quickly we can make a very simple in-memory implementation and then later we can create an implementation backed by whatever storage mechanism we prefer.

- `GET` 需要一个 `PlayerStore` _thing_ 来获取玩家的分数。这应该是一个接口，所以当我们测试时，我们可以创建一个简单的存根来测试我们的代码，而无需实现任何实际的存储代码。
- 对于`POST`，我们可以_spy_它对`PlayerStore`的调用，以确保它正确存储玩家。我们的保存实现不会与检索耦合。
- 为了快速拥有一些可运行的软件，我们可以在内存中进行一个非常简单的实现，然后我们可以创建一个由我们喜欢的任何存储机制支持的实现。

## Write the test first

## 先写测试

We can write a test and make it pass by returning a hard-coded value to get us started. Kent Beck refers this as "Faking it". Once we have a working test we can then write more tests to help us remove that constant.

我们可以编写一个测试并通过返回一个硬编码的值来让我们开始。 Kent Beck 将此称为“伪造”。一旦我们有了一个有效的测试，我们就可以编写更多的测试来帮助我们删除这个常量。

By doing this very small step, we can make the important start of getting an overall project structure working correctly without having to worry too much about our application logic.

通过这非常小的一步，我们可以为使整个项目结构正常工作的重要开始，而不必过多担心我们的应用程序逻辑。

To create a web server in Go you will typically call [ListenAndServe](https://golang.org/pkg/net/http/#ListenAndServe).

要在 Go 中创建 Web 服务器，您通常会调用 [ListenAndServe](https://golang.org/pkg/net/http/#ListenAndServe)。

```go
func ListenAndServe(addr string, handler Handler) error
```

This will start a web server listening on a port, creating a goroutine for every request and running it against a [`Handler`](https://golang.org/pkg/net/http/#Handler).

这将启动一个监听端口的 web 服务器，为每个请求创建一个 goroutine 并针对 [`Handler`](https://golang.org/pkg/net/http/#Handler) 运行它。

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

A type implements the Handler interface by implementing the `ServeHTTP` method which expects two arguments, the first is where we _write our response_ and the second is the HTTP request that was sent to the server.

一个类型通过实现 `ServeHTTP` 方法来实现 Handler 接口，该方法需要两个参数，第一个是我们_写我们的响应_的地方，第二个是发送到服务器的 HTTP 请求。

Let's create a file named `server_test.go` and write a test for a function `PlayerServer` that takes in those two arguments. The request sent in will be to get a player's score, which we expect to be `"20"`.
```go
func TestGETPlayers(t *testing.T) {
    t.Run("returns Pepper's score", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
        response := httptest.NewRecorder()

        PlayerServer(response, request)

        got := response.Body.String()
        want := "20"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}
```

In order to test our server, we will need a `Request` to send in and we'll want to _spy_ on what our handler writes to the `ResponseWriter`. 

为了测试我们的服务器，我们需要一个 `Request` 来发送，我们想要 _spy_ 我们的处理程序写入 `ResponseWriter` 的内容。

-   We use `http.NewRequest` to create a request. The first argument is the request's method and the second is the request's path. The `nil` argument refers to the request's body, which we don't need to set in this case.
-   `net/http/httptest` has a spy already made for us called `ResponseRecorder` so we can use that. It has many helpful methods to inspect what has been written as a response.

- 我们使用 `http.NewRequest` 来创建请求。第一个参数是请求的方法，第二个参数是请求的路径。 `nil` 参数指的是请求的主体，在这种情况下我们不需要设置。
- `net/http/httptest` 已经为我们制作了一个名为 `ResponseRecorder` 的间谍，因此我们可以使用它。它有许多有用的方法来检查作为响应所写的内容。

## Try to run the test

## 尝试运行测试

`./server_test.go:13:2: undefined: PlayerServer`

`./server_test.go:13:2: undefined: PlayerServer`

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

The compiler is here to help, just listen to it.

编译器是来帮忙的，听听就好。

Create a file named `server.go` and define `PlayerServer`

创建一个名为 `server.go` 的文件并定义 `PlayerServer`

```go
func PlayerServer() {}
```

Try again

再试一次

```
./server_test.go:13:14: too many arguments in call to PlayerServer
    have (*httptest.ResponseRecorder, *http.Request)
    want ()
```

Add the arguments to our function

将参数添加到我们的函数中

```go
import "net/http"

func PlayerServer(w http.ResponseWriter, r *http.Request) {

}
```

The code now compiles and the test fails

代码现在编译，测试失败

```
=== RUN   TestGETPlayers/returns_Pepper's_score
    --- FAIL: TestGETPlayers/returns_Pepper's_score (0.00s)
        server_test.go:20: got '', want '20'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

From the DI chapter, we touched on HTTP servers with a `Greet` function. We learned that net/http's `ResponseWriter` also implements io `Writer` so we can use `fmt.Fprint` to send strings as HTTP responses.

在 DI 一章中，我们接触了带有 `Greet` 函数的 HTTP 服务器。我们了解到 net/http 的 `ResponseWriter` 也实现了 io `Writer`，因此我们可以使用 `fmt.Fprint` 将字符串作为 HTTP 响应发送。

```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "20")
}
```

The test should now pass.

测试现在应该通过。

## Complete the scaffolding

## 完成脚手架

We want to wire this up into an application. This is important because

我们想把它连接到一个应用程序中。这很重要，因为

-   We'll have _actual working software_, we don't want to write tests for the sake of it, it's good to see the code in action.
-   As we refactor our code, it's likely we will change the structure of the program. We want to make sure this is reflected in our application too as part of the incremental approach.

- 我们将拥有_实际工作的软件_，我们不想为了它而编写测试，很高兴看到代码在运行。
- 当我们重构我们的代码时，我们很可能会改变程序的结构。我们希望确保这也作为增量方法的一部分反映在我们的应用程序中。

Create a new `main.go` file for our application and put this code in

为我们的应用程序创建一个新的 `main.go` 文件并将此代码放入

```go
package main

import (
    "log"
    "net/http"
)

func main() {
    handler := http.HandlerFunc(PlayerServer)
    log.Fatal(http.ListenAndServe(":5000", handler))
}
```

So far all of our application code has been in one file, however, this isn't best practice for larger projects where you'll want to separate things into different files.

到目前为止，我们所有的应用程序代码都在一个文件中，但是，对于希望将内容分成不同文件的大型项目来说，这不是最佳实践。

To run this, do `go build` which will take all the `.go` files in the directory and build you a program. You can then execute it with `./myprogram`.

要运行它，请执行 `go build`，它将获取目录中的所有 `.go` 文件并构建一个程序。然后你可以用`./myprogram` 来执行它。

### `http.HandlerFunc`

###`http.HandlerFunc`

Earlier we explored that the `Handler` interface is what we need to implement in order to make a server. _Typically_ we do that by creating a `struct` and make it implement the interface by implementing its own ServeHTTP method. However the use-case for structs is for holding data but _currently_ we have no state, so it doesn't feel right to be creating one.

之前我们探讨了`Handler` 接口是我们需要实现的以制作服务器。 _通常_我们通过创建一个 `struct` 来实现，并通过实现自己的 ServeHTTP 方法使其实现接口。然而，结构体的用例是用于保存数据，但_当前_我们没有状态，所以创建一个状态是不正确的。

[HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc) lets us avoid this.

[HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc) 让我们避免这种情况。

> The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

> HandlerFunc 类型是一个适配器，允许将普通函数用作 HTTP 处理程序。如果 f 是具有适当签名的函数，则 HandlerFunc(f) 是调用 f 的处理程序。

```go
type HandlerFunc func(ResponseWriter, *Request)
```

From the documentation, we see that type `HandlerFunc` has already implemented the `ServeHTTP` method.
By type casting our `PlayerServer` function with it, we have now implemented the required `Handler`.

从文档中，我们看到类型 `HandlerFunc` 已经实现了 `ServeHTTP` 方法。
通过使用它对我们的 `PlayerServer` 函数进行类型转换，我们现在已经实现了所需的 `Handler`。

### `http.ListenAndServe(":5000"...)`

### `http.ListenAndServe(":5000"...)`

`ListenAndServe` takes a port to listen on a `Handler`. If there is a problem the web server will return an error, an example of that might be the port already being listened to. For that reason we wrap the call in `log.Fatal` to log the error to the user.

`ListenAndServe` 需要一个端口来监听 `Handler`。如果出现问题，Web 服务器将返回错误，示例可能是已在侦听的端口。出于这个原因，我们将调用包装在 `log.Fatal` 中以将错误记录给用户。

What we're going to do now is write _another_ test to force us into making a positive change to try and move away from the hard-coded value.

我们现在要做的是编写 _another_ 测试，迫使我们做出积极的改变，试图摆脱硬编码的值。

## Write the test first

## 先写测试

We'll add another subtest to our suite which tries to get the score of a different player, which will break our hard-coded approach.

我们将在我们的套件中添加另一个子测试，它试图获得不同玩家的分数，这将打破我们的硬编码方法。

```go
t.Run("returns Floyd's score", func(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
    response := httptest.NewRecorder()

    PlayerServer(response, request)

    got := response.Body.String()
    want := "10"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
})
```

You may have been thinking 

你可能一直在想

> Surely we need some kind of concept of storage to control which player gets what score. It's weird that the values seem so arbitrary in our tests.

> 当然，我们需要某种存储概念来控制哪个玩家获得什么分数。奇怪的是，这些值在我们的测试中看起来如此随意。

Remember we are just trying to take as small as steps as reasonably possible, so we're just trying to break the constant for now.

请记住，我们只是在尽可能合理地采取尽可能小的步骤，因此我们现在只是想打破常规。

## Try to run the test

## 尝试运行测试

```
=== RUN   TestGETPlayers/returns_Pepper's_score
    --- PASS: TestGETPlayers/returns_Pepper's_score (0.00s)
=== RUN   TestGETPlayers/returns_Floyd's_score
    --- FAIL: TestGETPlayers/returns_Floyd's_score (0.00s)
        server_test.go:34: got '20', want '10'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
//server.go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")

    if player == "Pepper" {
        fmt.Fprint(w, "20")
        return
    }

    if player == "Floyd" {
        fmt.Fprint(w, "10")
        return
    }
}
```

This test has forced us to actually look at the request's URL and make a decision. So whilst in our heads, we may have been worrying about player stores and interfaces the next logical step actually seems to be about _routing_.

这个测试迫使我们实际查看请求的 URL 并做出决定。因此，虽然在我们的脑海中，我们可能一直在担心玩家商店和界面，但下一个合乎逻辑的步骤实际上似乎是关于 _routing_。

If we had started with the store code the amount of changes we'd have to do would be very large compared to this. **This is a smaller step towards our final goal and was driven by tests**.

如果我们从商店代码开始，与此相比，我们必须进行的更改量将非常大。 **这是朝着我们最终目标迈出的一小步，由测试驱动**。

We're resisting the temptation to use any routing libraries right now, just the smallest step to get our test passing.

我们现在抵制使用任何路由库的诱惑，这只是让我们的测试通过的最小步骤。

`r.URL.Path` returns the path of the request which we can then use [`strings.TrimPrefix`](https://golang.org/pkg/strings/#TrimPrefix) to trim away `/players/` to get the requested player. It's not very robust but will do the trick for now.

`r.URL.Path` 返回请求的路径，然后我们可以使用 [`strings.TrimPrefix`](https://golang.org/pkg/strings/#TrimPrefix) 将 `/players/` 修剪为获取请求的播放器。它不是很健壮，但现在可以解决问题。

## Refactor

## 重构

We can simplify the `PlayerServer` by separating out the score retrieval into a function

我们可以通过将分数检索分离成一个函数来简化`PlayerServer`

```go
//server.go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")

    fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
    if name == "Pepper" {
        return "20"
    }

    if name == "Floyd" {
        return "10"
    }

    return ""
}
```

And we can DRY up some of the code in the tests by making some helpers

我们可以通过创建一些帮助程序来干掉测试中的一些代码

```go
//server_test.go
func TestGETPlayers(t *testing.T) {
    t.Run("returns Pepper's score", func(t *testing.T) {
        request := newGetScoreRequest("Pepper")
        response := httptest.NewRecorder()

        PlayerServer(response, request)

        assertResponseBody(t, response.Body.String(), "20")
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request := newGetScoreRequest("Floyd")
        response := httptest.NewRecorder()

        PlayerServer(response, request)

        assertResponseBody(t, response.Body.String(), "10")
    })
}

func newGetScoreRequest(name string) *http.Request {
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
    return req
}

func assertResponseBody(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("response body is wrong, got %q want %q", got, want)
    }
}
```

However, we still shouldn't be happy. It doesn't feel right that our server knows the scores.

然而，我们仍然不应该高兴。我们的服务器知道分数感觉不对。

Our refactoring has made it pretty clear what to do.

我们的重构已经很清楚要做什么。

We moved the score calculation out of the main body of our handler into a function `GetPlayerScore`. This feels like the right place to separate the concerns using interfaces.

我们将分数计算从处理程序的主体移到一个函数“GetPlayerScore”中。这感觉像是使用接口分离关注点的正确位置。

Let's move our function we re-factored to be an interface instead

让我们将我们重构的函数移动为一个接口

```go
type PlayerStore interface {
    GetPlayerScore(name string) int
}
```

For our `PlayerServer` to be able to use a `PlayerStore`, it will need a reference to one. Now feels like the right time to change our architecture so that our `PlayerServer` is now a `struct`.

为了让我们的“PlayerServer”能够使用“PlayerStore”，它需要一个引用。现在感觉是时候改变我们的架构了，我们的“PlayerServer”现在是一个“struct”。

```go
type PlayerServer struct {
    store PlayerStore
}
```

Finally, we will now implement the `Handler` interface by adding a method to our new struct and putting in our existing handler code.

最后，我们现在将通过向我们的新结构添加一个方法并放入我们现有的处理程序代码来实现 `Handler` 接口。

```go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")
    fmt.Fprint(w, p.store.GetPlayerScore(player))
}
```

The only other change is we now call our `store.GetPlayerScore` to get the score, rather than the local function we defined (which we can now delete).

唯一的另一个变化是我们现在调用我们的 `store.GetPlayerScore` 来获取分数，而不是我们定义的本地函数（我们现在可以删除）。

Here is the full code listing of our server

这是我们服务器的完整代码清单

```go
//server.go
type PlayerStore interface {
    GetPlayerScore(name string) int
}

type PlayerServer struct {
    store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")
    fmt.Fprint(w, p.store.GetPlayerScore(player))
}
```

### Fix the issues 

### 修复问题

This was quite a few changes and we know our tests and application will no longer compile, but just relax and let the compiler work through it.

这是相当多的变化，我们知道我们的测试和应用程序将不再编译，而只是放松并让编译器完成它。

`./main.go:9:58: type PlayerServer is not an expression`

`./main.go:9:58: 类型 PlayerServer 不是表达式`

We need to change our tests to instead create a new instance of our `PlayerServer` and then call its method `ServeHTTP`.

我们需要更改我们的测试以创建我们的“PlayerServer”的新实例，然后调用它的方法“ServeHTTP”。

```go
//server_test.go
func TestGETPlayers(t *testing.T) {
    server := &PlayerServer{}

    t.Run("returns Pepper's score", func(t *testing.T) {
        request := newGetScoreRequest("Pepper")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertResponseBody(t, response.Body.String(), "20")
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request := newGetScoreRequest("Floyd")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertResponseBody(t, response.Body.String(), "10")
    })
}
```

Notice we're still not worrying about making stores _just yet_, we just want the compiler passing as soon as we can.

请注意，我们仍然不担心使存储_只是_，我们只是希望编译器尽快通过。

You should be in the habit of prioritising having code that compiles and then code that passes the tests.

您应该养成优先拥有可编译的代码，然后是通过测试的代码的习惯。

By adding more functionality (like stub stores) whilst the code isn't compiling, we are opening ourselves up to potentially _more_ compilation problems.

通过在代码未编译时添加更多功能（如存根存储），我们将面临潜在的 _more_ 编译问题。

Now `main.go` won't compile for the same reason.

现在`main.go` 不会因为同样的原因编译。

```go
func main() {
    server := &PlayerServer{}
    log.Fatal(http.ListenAndServe(":5000", server))
}
```

Finally, everything is compiling but the tests are failing

最后，一切都在编译，但测试失败

```
=== RUN   TestGETPlayers/returns_the_Pepper's_score
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
    panic: runtime error: invalid memory address or nil pointer dereference
```

This is because we have not passed in a `PlayerStore` in our tests. We'll need to make a stub one up.

这是因为我们没有在测试中传入 `PlayerStore`。我们需要制作一个存根。

```go
//server_test.go
type StubPlayerStore struct {
    scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    return score
}
```

A `map` is a quick and easy way of making a stub key/value store for our tests. Now let's create one of these stores for our tests and send it into our `PlayerServer`.

`map` 是一种为我们的测试制作存根键/值存储的快速简便的方法。现在让我们为我们的测试创建这些存储之一并将其发送到我们的“PlayerServer”。

```go
//server_test.go
func TestGETPlayers(t *testing.T) {
    store := StubPlayerStore{
        map[string]int{
            "Pepper": 20,
            "Floyd":  10,
        },
    }
    server := &PlayerServer{&store}

    t.Run("returns Pepper's score", func(t *testing.T) {
        request := newGetScoreRequest("Pepper")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertResponseBody(t, response.Body.String(), "20")
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request := newGetScoreRequest("Floyd")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertResponseBody(t, response.Body.String(), "10")
    })
}
```

Our tests now pass and are looking better. The _intent_ behind our code is clearer now due to the introduction of the store. We're telling the reader that because we have _this data in a `PlayerStore`_ that when you use it with a `PlayerServer` you should get the following responses.

我们的测试现在通过了并且看起来更好了。由于 store 的引入，我们代码背后的 _intent_ 现在更加清晰。我们告诉读者，因为我们在 `PlayerStore`_ 中有 _this 数据，当你将它与 `PlayerServer` 一起使用时，你应该得到以下响应。

### Run the application

### 运行应用程序

Now our tests are passing the last thing we need to do to complete this refactor is to check if our application is working. The program should start up but you'll get a horrible response if you try and hit the server at `http://localhost:5000/players/Pepper`.

现在我们的测试通过了我们需要做的最后一件事来完成这个重构是检查我们的应用程序是否正常工作。程序应该会启动，但是如果您尝试点击位于 `http://localhost:5000/players/Pepper` 的服务器，你会得到一个可怕的响应。

The reason for this is that we have not passed in a `PlayerStore`.

这样做的原因是我们没有传入 `PlayerStore`。

We'll need to make an implementation of one, but that's difficult right now as we're not storing any meaningful data so it'll have to be hard-coded for the time being.

我们需要实现一个，但现在这很困难，因为我们没有存储任何有意义的数据，因此暂时必须对其进行硬编码。

```go
//main.go
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return 123
}

func main() {
    server := &PlayerServer{&InMemoryPlayerStore{}}
    log.Fatal(http.ListenAndServe(":5000", server))
}
```

If you run `go build` again and hit the same URL you should get `"123"`. Not great, but until we store data that's the best we can do.
It also didn't feel great that our main application was starting up but not actually working. We had to manually test to see the problem.

如果你再次运行 `go build` 并点击相同的 URL 你应该得到 `"123"`。不是很好，但直到我们存储我们可以做的最好的数据。
我们的主要应用程序正在启动但实际上没有工作，这也让人感觉不太好。我们必须手动测试才能看到问题。

We have a few options as to what to do next

我们有一些关于下一步做什么的选择

-   Handle the scenario where the player doesn't exist
-   Handle the `POST /players/{name}` scenario 

- 处理玩家不存在的场景
- 处理`POST /players/{name}` 场景

Whilst the `POST` scenario gets us closer to the "happy path", I feel it'll be easier to tackle the missing player scenario first as we're in that context already. We'll get to the rest later.

虽然`POST` 场景让我们更接近“快乐之路”，但我觉得首先解决丢失的玩家场景会更容易，因为我们已经处于这种情况下。我们稍后再讲。

## Write the test first

## 先写测试

Add a missing player scenario to our existing suite

在我们现有的套件中添加一个缺失的玩家场景

```go
//server_test.go
t.Run("returns 404 on missing players", func(t *testing.T) {
    request := newGetScoreRequest("Apollo")
    response := httptest.NewRecorder()

    server.ServeHTTP(response, request)

    got := response.Code
    want := http.StatusNotFound

    if got != want {
        t.Errorf("got status %d want %d", got, want)
    }
})
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestGETPlayers/returns_404_on_missing_players
    --- FAIL: TestGETPlayers/returns_404_on_missing_players (0.00s)
        server_test.go:56: got status 200 want 404
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")

    w.WriteHeader(http.StatusNotFound)

    fmt.Fprint(w, p.store.GetPlayerScore(player))
}
```

Sometimes I heavily roll my eyes when TDD advocates say "make sure you just write the minimal amount of code to make it pass" as it can feel very pedantic.

有时，当 TDD 倡导者说“确保你只编写最少量的代码以使其通过”时，我会翻白眼，因为这会让人感觉非常迂腐。

But this scenario illustrates the example well. I have done the bare minimum (knowing it is not correct), which is write a `StatusNotFound` on **all responses** but all our tests are passing!

但是这个场景很好地说明了这个例子。我已经完成了最低限度的工作（知道它不正确），即在 **所有响应** 上写一个 `StatusNotFound`，但我们所有的测试都通过了！

**By doing the bare minimum to make the tests pass it can highlight gaps in your tests**. In our case, we are not asserting that we should be getting a `StatusOK` when players _do_ exist in the store.

**通过做最少的事情来使测试通过，它可以突出测试中的差距**。在我们的例子中，当玩家 _do_ 存在于商店中时，我们并没有断言我们应该得到一个 `StatusOK`。

Update the other two tests to assert on the status and fix the code.

更新其他两个测试以断言状态并修复代码。

Here are the new tests

这是新的测试

```go
//server_test.go
func TestGETPlayers(t *testing.T) {
    store := StubPlayerStore{
        map[string]int{
            "Pepper": 20,
            "Floyd":  10,
        },
    }
    server := &PlayerServer{&store}

    t.Run("returns Pepper's score", func(t *testing.T) {
        request := newGetScoreRequest("Pepper")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusOK)
        assertResponseBody(t, response.Body.String(), "20")
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request := newGetScoreRequest("Floyd")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusOK)
        assertResponseBody(t, response.Body.String(), "10")
    })

    t.Run("returns 404 on missing players", func(t *testing.T) {
        request := newGetScoreRequest("Apollo")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusNotFound)
    })
}

func assertStatus(t testing.TB, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("did not get correct status, got %d, want %d", got, want)
    }
}

func newGetScoreRequest(name string) *http.Request {
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
    return req
}

func assertResponseBody(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("response body is wrong, got %q want %q", got, want)
    }
}
```

We're checking the status in all our tests now so I made a helper `assertStatus` to facilitate that.

我们现在正在检查所有测试中的状态，所以我创建了一个助手 `assertStatus` 来帮助它。

Now our first two tests fail because of the 404 instead of 200, so we can fix `PlayerServer` to only return not found if the score is 0.

现在我们的前两个测试由于 404 而不是 200 而失败，所以我们可以修复 `PlayerServer` 以仅在分数为 0 时返回 not found。

```go
//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")

    score := p.store.GetPlayerScore(player)

    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }

    fmt.Fprint(w, score)
}
```

### Storing scores

### 存储分数

Now that we can retrieve scores from a store it now makes sense to be able to store new scores.

既然我们可以从商店中检索分数，那么现在能够存储新分数就变得有意义了。

## Write the test first

## 先写测试

```go
//server_test.go
func TestStoreWins(t *testing.T) {
    store := StubPlayerStore{
        map[string]int{},
    }
    server := &PlayerServer{&store}

    t.Run("it returns accepted on POST", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusAccepted)
    })
}
```

For a start let's just check we get the correct status code if we hit the particular route with POST. This lets us drive out the functionality of accepting a different kind of request and handling it differently to `GET /players/{name}`. Once this works we can then start asserting on our handler's interaction with the store.

首先，让我们检查一下，如果我们使用 POST 命中特定路由，我们是否获得了正确的状态代码。这让我们排除了接受不同类型请求并以不同方式处理它的功能，以“GET /players/{name}”。一旦成功，我们就可以开始断言我们的处理程序与商店的交互。

## Try to run the test

## 尝试运行测试

```
=== RUN   TestStoreWins/it_returns_accepted_on_POST
    --- FAIL: TestStoreWins/it_returns_accepted_on_POST (0.00s)
        server_test.go:70: did not get correct status, got 404, want 202
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Remember we are deliberately committing sins, so an `if` statement based on the request's method will do the trick.

请记住，我们是故意犯罪，因此基于请求方法的“if”语句将起作用。

```go
//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodPost {
        w.WriteHeader(http.StatusAccepted)
        return
    }

    player := strings.TrimPrefix(r.URL.Path, "/players/")

    score := p.store.GetPlayerScore(player)

    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }

    fmt.Fprint(w, score)
}
```

## Refactor

## 重构

The handler is looking a bit muddled now. Let's break the code up to make it easier to follow and isolate the different functionality into new functions.

处理程序现在看起来有点混乱。让我们分解代码，以便更容易理解并将不同的功能隔离到新的功能中。

```go
//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
    case http.MethodPost:
        p.processWin(w)
    case http.MethodGet:
        p.showScore(w, r)
    }

}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")

    score := p.store.GetPlayerScore(player)

    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }

    fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
    w.WriteHeader(http.StatusAccepted)
}
```

This makes the routing aspect of `ServeHTTP` a bit clearer and means our next iterations on storing can just be inside `processWin`.

这使得 `ServeHTTP` 的路由方面更加清晰，意味着我们的下一次存储迭代可以只在 `processWin` 中。

Next, we want to check that when we do our `POST /players/{name}` that our `PlayerStore` is told to record the win.

接下来，我们要检查当我们执行 `POST /players/{name}` 时，我们的 `PlayerStore` 被告知记录胜利。

## Write the test first

## 先写测试

We can accomplish this by extending our `StubPlayerStore` with a new `RecordWin` method and then spy on its invocations.

我们可以通过使用新的 `RecordWin` 方法扩展我们的 `StubPlayerStore` 来实现这一点，然后监视它的调用。

```go
//server_test.go
type StubPlayerStore struct {
    scores   map[string]int
    winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    return score
}

func (s *StubPlayerStore) RecordWin(name string) {
    s.winCalls = append(s.winCalls, name)
}
```

Now extend our test to check the number of invocations for a start

现在扩展我们的测试以检查开始的调用次数

```go
//server_test.go
func TestStoreWins(t *testing.T) {
    store := StubPlayerStore{
        map[string]int{},
    }
    server := &PlayerServer{&store}

    t.Run("it records wins when POST", func(t *testing.T) {
        request := newPostWinRequest("Pepper")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusAccepted)

        if len(store.winCalls) != 1 {
            t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
        }
    })
}

func newPostWinRequest(name string) *http.Request {
    req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
    return req
}
```

## Try to run the test

## 尝试运行测试

```
./server_test.go:26:20: too few values in struct initializer
./server_test.go:65:20: too few values in struct initializer
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We need to update our code where we create a `StubPlayerStore` as we've added a new field

我们需要更新我们创建 `StubPlayerStore` 的代码，因为我们添加了一个新字段

```go
//server_test.go
store := StubPlayerStore{
    map[string]int{},
    nil,
}
```

```
--- FAIL: TestStoreWins (0.00s)
    --- FAIL: TestStoreWins/it_records_wins_when_POST (0.00s)
        server_test.go:80: got 0 calls to RecordWin want 1
```

## Write enough code to make it pass

## 编写足够的代码使其通过

As we're only asserting the number of calls rather than the specific values it makes our initial iteration a little smaller.

由于我们只是断言调用次数而不是特定值，因此我们的初始迭代会小一些。

We need to update `PlayerServer`'s idea of what a `PlayerStore` is by changing the interface if we're going to be able to call `RecordWin`.

如果我们要能够调用“RecordWin”，我们需要通过更改接口来更新“PlayerServer”对“PlayerStore”的概念。

```go
//server.go
type PlayerStore interface {
    GetPlayerScore(name string) int
    RecordWin(name string)
}
```

By doing this `main` no longer compiles

通过这样做`main`不再编译

```
./main.go:17:46: cannot use InMemoryPlayerStore literal (type *InMemoryPlayerStore) as type PlayerStore in field value:
    *InMemoryPlayerStore does not implement PlayerStore (missing RecordWin method)
```

The compiler tells us what's wrong. Let's update `InMemoryPlayerStore` to have that method.

编译器会告诉我们出了什么问题。让我们更新 `InMemoryPlayerStore` 以使用该方法。

```go
//main.go
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) RecordWin(name string) {}
```

Try and run the tests and we should be back to compiling code - but the test is still failing.

尝试运行测试，我们应该回到编译代码 - 但测试仍然失败。

Now that `PlayerStore` has `RecordWin` we can call it within our `PlayerServer`

现在`PlayerStore` 有`RecordWin` 我们可以在我们的`PlayerServer` 中调用它

```go
//server.go
func (p *PlayerServer) processWin(w http.ResponseWriter) {
    p.store.RecordWin("Bob")
    w.WriteHeader(http.StatusAccepted)
}
```

Run the tests and it should be passing! Obviously `"Bob"` isn't exactly what we want to send to `RecordWin`, so let's further refine the test.

运行测试，它应该通过！显然 `"Bob"` 并不是我们想要发送给 `RecordWin` 的内容，所以让我们进一步完善测试。

## Write the test first

## 先写测试

```go
//server_test.go
t.Run("it records wins on POST", func(t *testing.T) {
    player := "Pepper"

    request := newPostWinRequest(player)
    response := httptest.NewRecorder()

    server.ServeHTTP(response, request)

    assertStatus(t, response.Code, http.StatusAccepted)

    if len(store.winCalls) != 1 {
        t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
    }

    if store.winCalls[0] != player {
        t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
    }
})
```

Now that we know there is one element in our `winCalls` slice we can safely reference the first one and check it is equal to `player`.

现在我们知道我们的 `winCalls` 切片中有一个元素，我们可以安全地引用第一个元素并检查它是否等于 `player`。

## Try to run the test

## 尝试运行测试

```
=== RUN   TestStoreWins/it_records_wins_on_POST
    --- FAIL: TestStoreWins/it_records_wins_on_POST (0.00s)
        server_test.go:86: did not store correct winner got 'Bob' want 'Pepper'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
//server.go
func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")
    p.store.RecordWin(player)
    w.WriteHeader(http.StatusAccepted)
}
```

We changed `processWin` to take `http.Request` so we can look at the URL to extract the player's name. Once we have that we can call our `store` with the correct value to make the test pass.

我们将 `processWin` 更改为采用 `http.Request`，以便我们可以查看 URL 以提取玩家的姓名。一旦我们有了它，我们就可以使用正确的值调用我们的 `store` 以使测试通过。

## Refactor

## 重构

We can DRY up this code a bit as we're extracting the player name the same way in two places

我们可以稍微整理一下这段代码，因为我们在两个地方以相同的方式提取玩家名称

```go
//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")

    switch r.Method {
    case http.MethodPost:
        p.processWin(w, player)
    case http.MethodGet:
        p.showScore(w, player)
    }
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
    score := p.store.GetPlayerScore(player)

    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }

    fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
    p.store.RecordWin(player)
    w.WriteHeader(http.StatusAccepted)
}
```

Even though our tests are passing we don't really have working software. If you try and run `main` and use the software as intended it doesn't work because we haven't got round to implementing `PlayerStore` correctly. This is fine though; by focusing on our handler we have identified the interface that we need, rather than trying to design it up-front.

即使我们的测试通过了，我们也没有真正可以运行的软件。如果您尝试运行 `main` 并按预期使用该软件，它将无法正常工作，因为我们还没有准备好正确实现 `PlayerStore`。不过这很好；通过关注我们的处理程序，我们已经确定了我们需要的接口，而不是试图预先设计它。

We _could_ start writing some tests around our `InMemoryPlayerStore` but it's only here temporarily until we implement a more robust way of persisting player scores (i.e. a database).

我们_可以_开始围绕我们的“InMemoryPlayerStore”编写一些测试，但它只是暂时的，直到我们实现一种更强大的保持玩家分数的方法（即数据库）。

What we'll do for now is write an _integration test_ between our `PlayerServer` and `InMemoryPlayerStore` to finish off the functionality. This will let us get to our goal of being confident our application is working, without having to directly test `InMemoryPlayerStore`. Not only that, but when we get around to implementing `PlayerStore` with a database, we can test that implementation with the same integration test.

我们现在要做的是在我们的 `PlayerServer` 和 `InMemoryPlayerStore` 之间编写一个 _integration test_ 来完成功能。这将使我们能够确信我们的应用程序正在运行，而无需直接测试 InMemoryPlayerStore。不仅如此，当我们开始使用数据库实现 `PlayerStore` 时，我们可以使用相同的集成测试来测试该实现。

### Integration tests

### 集成测试

Integration tests can be useful for testing that larger areas of your system work but you must bear in mind:

集成测试对于测试系统的较大区域工作很有用，但您必须牢记：

-   They are harder to write
-   When they fail, it can be difficult to know why (usually it's a bug within a component of the integration test) and so can be harder to fix
-   They are sometimes slower to run (as they often are used with "real" components, like a database)

- 他们更难写
- 当它们失败时，可能很难知道原因（通常是集成测试组件中的错误），因此更难修复
- 它们有时运行速度较慢（因为它们经常与“真实”组件一起使用，例如数据库）

For that reason, it is recommended that you research _The Test Pyramid_.

因此，建议您研究_测试金字塔_。

## Write the test first

## 先写测试

In the interest of brevity, I am going to show you the final refactored integration test.

为简洁起见，我将向您展示最终的重构集成测试。

```go
//server_integration_test.go
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
    store := InMemoryPlayerStore{}
    server := PlayerServer{&store}
    player := "Pepper"

    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

    response := httptest.NewRecorder()
    server.ServeHTTP(response, newGetScoreRequest(player))
    assertStatus(t, response.Code, http.StatusOK)

    assertResponseBody(t, response.Body.String(), "3")
}
```

-   We are creating our two components we are trying to integrate with: `InMemoryPlayerStore` and `PlayerServer`.
-   We then fire off 3 requests to record 3 wins for `player`. We're not too concerned about the status codes in this test as it's not relevant to whether they are integrating well.
-   The next response we do care about (so we store a variable `response`) because we are going to try and get the `player`'s score.

- 我们正在创建我们试图与之集成的两个组件：`InMemoryPlayerStore` 和 `PlayerServer`。
- 然后我们发出 3 个请求来记录 `player` 的 3 场胜利。我们不太关心此测试中的状态代码，因为这与它们是否集成良好无关。
- 我们关心的下一个响应（所以我们存储一个变量 `response`），因为我们将尝试获得 `player` 的分数。

## Try to run the test

## 尝试运行测试

```
--- FAIL: TestRecordingWinsAndRetrievingThem (0.00s)
    server_integration_test.go:24: response body is wrong, got '123' want '3'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

I am going to take some liberties here and write more code than you may be comfortable with without writing a test.

我将在这里采取一些自由，编写比您在不编写测试的情况下可能会感到舒服的更多代码。

_This is allowed!_ We still have a test checking things should be working correctly but it is not around the specific unit we're working with (`InMemoryPlayerStore`).

_这是允许的！_ 我们仍然有一个测试来检查事情是否应该正常工作，但它不是我们正在使用的特定单元（`InMemoryPlayerStore`）。

If I were to get stuck in this scenario, I would revert my changes back to the failing test and then write more specific unit tests around `InMemoryPlayerStore` to help me drive out a solution.

如果我陷入这种情况，我会将我的更改恢复到失败的测试，然后围绕 `InMemoryPlayerStore` 编写更具体的单元测试以帮助我制定解决方案。

```go
//in_memory_player_store.go
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
    return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
    store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return i.store[name]
}
```

-   We need to store the data so I've added a `map[string]int` to the `InMemoryPlayerStore` struct
-   For convenience I've made `NewInMemoryPlayerStore` to initialise the store, and updated the integration test to use it:
     ```go
     //server_integration_test.go
    store := NewInMemoryPlayerStore()
    server := PlayerServer{store}
    ```

-   The rest of the code is just wrapping around the `map`

- 其余的代码只是围绕着`map`

The integration test passes, now we just need to change `main` to use `NewInMemoryPlayerStore()`

集成测试通过，现在我们只需要将 `main` 更改为使用 `NewInMemoryPlayerStore()`

```go
//main.go
package main

import (
    "log"
    "net/http"
)

func main() {
    server := &PlayerServer{NewInMemoryPlayerStore()}
    log.Fatal(http.ListenAndServe(":5000", server))
}
```

Build it, run it and then use `curl` to test it out.

构建它，运行它，然后使用 `curl` 来测试它。

-   Run this a few times, change the player names if you like `curl -X POST http://localhost:5000/players/Pepper`
-   Check scores with `curl http://localhost:5000/players/Pepper`

- 运行几次，如果你喜欢`curl -X POST http://localhost:5000/players/Pepper`，请更改播放器名称
- 使用 `curl http://localhost:5000/players/Pepper` 检查分数

Great! You've made a REST-ish service. To take this forward you'd want to pick a data store to persist the scores longer than the length of time the program runs.

伟大的！您已经创建了一个 REST 风格的服务。为了推进这一点，您需要选择一个数据存储来保存比程序运行时间更长的分数。

-   Pick a store (Bolt? Mongo? Postgres? File system?)
-   Make `PostgresPlayerStore` implement `PlayerStore`
-   TDD the functionality so you're sure it works
-   Plug it into the integration test, check it's still ok
-   Finally plug it into `main`

- 选择一个商店（Bolt？Mongo？Postgres？文件系统？）
- 使`PostgresPlayerStore` 实现`PlayerStore`
- TDD 功能所以你确定它有效
- 将其插入集成测试，检查它是否还可以
- 最后将其插入`main`

## Refactor

## 重构

We are almost there! Lets take some effort to prevent concurrency errors like these

我们就快到了！让我们努力防止像这样的并发错误

```
fatal error: concurrent map read and map write
```

By adding mutexes, we enforce concurrency safety especially for the counter in our `RecordWin` function. Read more about mutexes in the sync chapter.

通过添加互斥体，我们加强了并发安全，特别是对于我们的“RecordWin”函数中的计数器。在同步章节中阅读更多关于互斥锁的信息。

## Wrapping up

##  总结

### `http.Handler`

###`http.Handler`

-   Implement this interface to create web servers
-   Use `http.HandlerFunc` to turn ordinary functions into `http.Handler`s
-   Use `httptest.NewRecorder` to pass in as a `ResponseWriter` to let you spy on the responses your handler sends
-   Use `http.NewRequest` to construct the requests you expect to come in to your system

- 实现此接口以创建 Web 服务器
- 使用`http.HandlerFunc`将普通函数变成`http.Handler`s
- 使用 `httptest.NewRecorder` 作为 `ResponseWriter` 传入，让您监视处理程序发送的响应
- 使用 `http.NewRequest` 来构建您希望进入系统的请求

### Interfaces, Mocking and DI

### 接口、模拟和 DI

-   Lets you iteratively build the system up in smaller chunks
-   Allows you to develop a handler that needs a storage without needing actual storage
-   TDD to drive out the interfaces you need

- 让您以更小的块迭代地构建系统
- 允许您开发需要存储而不需要实际存储的处理程序
- TDD 来驱动你需要的接口

### Commit sins, then refactor (and then commit to source control)

### 提交罪过，然后重构（然后提交到源代码控制）

-   You need to treat having failing compilation or failing tests as a red situation that you need to get out of as soon as you can.
-   Write just the necessary code to get there. _Then_ refactor and make the code nice.
-   By trying to do too many changes whilst the code isn't compiling or the tests are failing puts you at risk of compounding the problems.
-   Sticking to this approach forces you to write small tests, which means small changes, which helps keep working on complex systems manageable. 

- 您需要将编译失败或测试失败视为需要尽快摆脱的红色情况。
- 只需编写必要的代码即可。 _Then_ 重构并使代码更美观。
- 在代码未编译或测试失败时尝试进行过多更改会使您面临使问题复杂化的风险。
- 坚持这种方法会迫使您编写小测试，这意味着小的更改，这有助于使复杂系统的工作保持可管理性。

