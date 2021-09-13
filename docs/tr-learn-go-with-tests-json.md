# JSON, routing & embedding

# JSON，路由和嵌入

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/json)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/json)**

[In the previous chapter](http-server.md) we created a web server to store how many games players have won.

[在上一章中](http-server.md) 我们创建了一个网络服务器来存储玩家赢得的游戏数量。

Our product owner has a new requirement; to have a new endpoint called `/league` which returns a list of all players stored. She would like this to be returned as JSON.

我们的产品负责人有一个新要求；有一个名为`/league` 的新端点，它返回存储的所有玩家的列表。她希望将其作为 JSON 返回。

## Here is the code we have so far

## 这是我们到目前为止的代码

```go
// server.go
package main

import (
    "fmt"
    "net/http"
    "strings"
)

type PlayerStore interface {
    GetPlayerScore(name string) int
    RecordWin(name string)
}

type PlayerServer struct {
    store PlayerStore
}

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

```go
// InMemoryPlayerStore.go
package main

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

```go
// main.go
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

You can find the corresponding tests in the link at the top of the chapter.

您可以在本章顶部的链接中找到相应的测试。

We'll start by making the league table endpoint.

我们将从制作排行榜端点开始。

## Write the test first

## 先写测试

We'll extend the existing suite as we have some useful test functions and a fake `PlayerStore` to use.

我们将扩展现有套件，因为我们有一些有用的测试功能和一个假的“PlayerStore”可供使用。

```go
//server_test.go
func TestLeague(t *testing.T) {
    store := StubPlayerStore{}
    server := &PlayerServer{&store}

    t.Run("it returns 200 on /league", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/league", nil)
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusOK)
    })
}
```

Before worrying about actual scores and JSON we will try and keep the changes small with the plan to iterate toward our goal. The simplest start is to check we can hit `/league` and get an `OK` back.

在担心实际分数和 JSON 之前，我们将尝试将更改保持在较小的范围内，并计划朝着我们的目标迭代。最简单的开始是检查我们是否可以点击 `/league` 并返回一个 `OK`。

## Try to run the test

## 尝试运行测试

```
     --- FAIL: TestLeague/it_returns_200_on_/league (0.00s)
        server_test.go:101: status code is wrong: got 404, want 200
FAIL
FAIL    playerstore    0.221s
FAIL
```

Our `PlayerServer` returns a `404 Not Found`, as if we were trying to get the wins for an unknown player. Looking at how `server.go` implements `ServeHTTP`, we realize that it always assumes to be called with a URL pointing to a specific player:

我们的 `PlayerServer` 返回一个 `404 Not Found`，就好像我们试图为一个未知玩家获取胜利一样。看看 `server.go` 如何实现 `ServeHTTP`，我们意识到它总是假设使用指向特定播放器的 URL 来调用：

```go
player := strings.TrimPrefix(r.URL.Path, "/players/")
```

In the previous chapter, we mentioned this was a fairly naive way of doing our routing. Our test informs us correctly that we need a concept how to deal with different request paths.

在前一章中，我们提到这是一种相当幼稚的路由方式。我们的测试正确地告诉我们，我们需要一个如何处理不同请求路径的概念。

## Write enough code to make it pass

## 编写足够的代码使其通过

Go has a built-in routing mechanism called [`ServeMux`](https://golang.org/pkg/net/http/#ServeMux) (request multiplexer) which lets you attach `http.Handler`s to particular request paths .

Go 有一个名为 [`ServeMux`](https://golang.org/pkg/net/http/#ServeMux)（请求多路复用器)的内置路由机制，它允许你将 `http.Handler`s 附加到特定的请求路径.

Let's commit some sins and get the tests passing in the quickest way we can, knowing we can refactor it with safety once we know the tests are passing.

让我们犯一些错误并以最快的方式通过测试，知道一旦我们知道测试通过，我们就可以安全地重构它。

```go
//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    router := http.NewServeMux()

    router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        player := strings.TrimPrefix(r.URL.Path, "/players/")

        switch r.Method {
        case http.MethodPost:
            p.processWin(w, player)
        case http.MethodGet:
            p.showScore(w, player)
        }
    }))

    router.ServeHTTP(w, r)
}
```

- When the request starts we create a router and then we tell it for `x` path use `y` handler. 

- 当请求开始时，我们创建一个路由器，然后我们告诉它为 `x` 路径使用 `y` 处理程序。

- So for our new endpoint, we use `http.HandlerFunc` and an _anonymous function_ to `w.WriteHeader(http.StatusOK)` when `/league` is requested to make our new test pass.
- For the `/players/` route we just cut and paste our code into another `http.HandlerFunc`.
- Finally, we handle the request that came in by calling our new router's `ServeHTTP` (notice how `ServeMux` is _also_ an `http.Handler`?)

- 因此，对于我们的新端点，当请求 `/league` 使我们的新测试通过时，我们使用 `http.HandlerFunc` 和一个 _anonymous function_ 到 `w.WriteHeader(http.StatusOK)`。
- 对于 `/players/` 路由，我们只是将代码剪切并粘贴到另一个 `http.HandlerFunc` 中。
- 最后，我们通过调用新路由器的 `ServeHTTP` 来处理传入的请求（注意 `ServeMux` 是如何_也_是一个 `http.Handler`？）

The tests should now pass.

测试现在应该通过了。

## Refactor

## 重构

`ServeHTTP` is looking quite big, we can separate things out a bit by refactoring our handlers into separate methods.

`ServeHTTP` 看起来很大，我们可以通过将处理程序重构为单独的方法来将事情分开。

```go
//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))

    router.ServeHTTP(w, r)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")

    switch r.Method {
    case http.MethodPost:
        p.processWin(w, player)
    case http.MethodGet:
        p.showScore(w, player)
    }
}
```

It's quite odd (and inefficient) to be setting up a router as a request comes in and then calling it. What we ideally want to do is have some kind of `NewPlayerServer` function which will take our dependencies and do the one-time setup of creating the router. Each request can then just use that one instance of the router.

在请求进来时设置路由器然后调用它是很奇怪的（而且效率低下）。理想情况下，我们想要做的是拥有某种“NewPlayerServer”函数，它将获取我们的依赖项并进行创建路由器的一次性设置。每个请求都可以只使用路由器的一个实例。

```go
//server.go
type PlayerServer struct {
    store  PlayerStore
    router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := &PlayerServer{
        store,
        http.NewServeMux(),
    }

    p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

    return p
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    p.router.ServeHTTP(w, r)
}
```

- `PlayerServer` now needs to store a router.
- We have moved the routing creation out of `ServeHTTP` and into our `NewPlayerServer` so this only has to be done once, not per request.
- You will need to update all the test and production code where we used to do `PlayerServer{&store}` with `NewPlayerServer(&store)`.

- `PlayerServer` 现在需要存储一个路由器。
- 我们已经将路由创建从 `ServeHTTP` 移到我们的 `NewPlayerServer` 中，因此这只需执行一次，而不是每个请求。
- 您需要更新我们过去使用 `NewPlayerServer(&store)` 执行 `PlayerServer{&store}` 的所有测试和生产代码。

### One final refactor

### 最后一次重构

Try changing the code to the following.

尝试将代码更改为以下内容。

```go
type PlayerServer struct {
    store PlayerStore
    http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := new(PlayerServer)

    p.store = store

    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))

    p.Handler = router

    return p
}
```

Then replace `server := &PlayerServer{&store}` with `server := NewPlayerServer(&store)` in `server_test.go`, `server_integration_test.go`, and `main.go`.

然后将 `server_test.go`、`server_integration_test.go` 和 `main.go` 中的 `server := &PlayerServer{&store}` 替换为 `server := NewPlayerServer(&store)`。

Finally make sure you **delete** `func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)` as it is no longer needed!

最后确保你 **delete** `func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)` 因为它不再需要了！

## Embedding

## 嵌入

We changed the second property of `PlayerServer`, removing the named property `router http.ServeMux` and replaced it with `http.Handler`; this is called _embedding_.

我们更改了 `PlayerServer` 的第二个属性，移除了命名属性 `router http.ServeMux` 并替换为 `http.Handler`；这称为_嵌入_。

> Go does not provide the typical, type-driven notion of subclassing, but it does have the ability to “borrow” pieces of an implementation by embedding types within a struct or interface.

> Go 不提供典型的、类型驱动的子类化概念，但它确实有能力通过在结构或接口中嵌入类型来“借用”实现的片段。

[Effective Go - Embedding](https://golang.org/doc/effective_go.html#embedding)

[Effective Go - 嵌入](https://golang.org/doc/effective_go.html#embedding)

What this means is that our `PlayerServer` now has all the methods that `http.Handler` has, which is just `ServeHTTP`.

这意味着我们的 `PlayerServer` 现在拥有 `http.Handler` 拥有的所有方法，也就是 `ServeHTTP`。

To "fill in" the `http.Handler` we assign it to the `router` we create in `NewPlayerServer`. We can do this because `http.ServeMux` has the method `ServeHTTP`.

为了“填充”`http.Handler`，我们将它分配给我们在`NewPlayerServer` 中创建的`router`。我们可以这样做是因为 `http.ServeMux` 有方法 `ServeHTTP`。

This lets us remove our own `ServeHTTP` method, as we are already exposing one via the embedded type.

这让我们可以移除我们自己的 `ServeHTTP` 方法，因为我们已经通过嵌入类型公开了一个方法。

Embedding is a very interesting language feature. You can use it with interfaces to compose new interfaces.

嵌入是一个非常有趣的语言特性。您可以将它与接口一起使用来组成新的接口。

```go
type Animal interface {
    Eater
    Sleeper
}
```

And you can use it with concrete types too, not just interfaces. As you'd expect if you embed a concrete type you'll have access to all its public methods and fields.

您也可以将它与具体类型一起使用，而不仅仅是接口。正如您所期望的，如果您嵌入一个具体类型，您将可以访问其所有公共方法和字段。

### Any downsides? 

### 有什么缺点吗？

You must be careful with embedding types because you will expose all public methods and fields of the type you embed. In our case, it is ok because we embedded just the _interface_ that we wanted to expose (`http.Handler`).

您必须小心嵌入类型，因为您将公开您嵌入的类型的所有公共方法和字段。在我们的例子中，这是可以的，因为我们只嵌入了我们想要公开的 _interface_ (`http.Handler`)。

If we had been lazy and embedded `http.ServeMux` instead (the concrete type) it would still work _but_ users of `PlayerServer` would be able to add new routes to our server because `Handle(path, handler)` would be public .

如果我们懒惰并嵌入了 `http.ServeMux`（具体类型），它仍然可以工作_但是_`PlayerServer` 的用户将能够向我们的服务器添加新路由，因为 `Handle(path, handler)` 将是公开的.

**When embedding types, really think about what impact that has on your public API.**

**在嵌入类型时，请认真考虑对您的公共 API 有何影响。**

It is a _very_ common mistake to misuse embedding and end up polluting your APIs and exposing the internals of your type.

滥用嵌入并最终污染您的 API 并暴露您的类型的内部结构是一个_非常_常见的错误。

Now we've restructured our application we can easily add new routes and have the start of the `/league` endpoint. We now need to make it return some useful information.

现在我们已经重构了我们的应用程序，我们可以轻松地添加新路由并开始`/league` 端点。我们现在需要让它返回一些有用的信息。

We should return some JSON that looks something like this.

我们应该返回一些看起来像这样的 JSON。

```json
[
   {
      "Name":"Bill",
      "Wins":10
   },
   {
      "Name":"Alice",
      "Wins":15
   }
]
```

## Write the test first

## 先写测试

We'll start by trying to parse the response into something meaningful.

我们将首先尝试将响应解析为有意义的内容。

```go
//server_test.go
func TestLeague(t *testing.T) {
    store := StubPlayerStore{}
    server := NewPlayerServer(&store)

    t.Run("it returns 200 on /league", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/league", nil)
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        var got []Player

        err := json.NewDecoder(response.Body).Decode(&got)

        if err != nil {
            t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
        }

        assertStatus(t, response.Code, http.StatusOK)
    })
}
```

### Why not test the JSON string?

### 为什么不测试 JSON 字符串？

You could argue a simpler initial step would be just to assert that the response body has a particular JSON string.

您可能会争辩说，一个更简单的初始步骤就是断言响应正文具有特定的 JSON 字符串。

In my experience tests that assert against JSON strings have the following problems.

根据我的经验，针对 JSON 字符串进行断言的测试存在以下问题。

- *Brittleness*. If you change the data-model your tests will fail.
- *Hard to debug*. It can be tricky to understand what the actual problem is when comparing two JSON strings.
- *Poor intention*. Whilst the output should be JSON, what's really important is exactly what the data is, rather than how it's encoded.
- *Re-testing the standard library*. There is no need to test how the standard library outputs JSON, it is already tested. Don't test other people's code.

- *脆性*。如果您更改数据模型，您的测试将失败。
- *难以调试*。在比较两个 JSON 字符串时，要理解实际问题是什么可能很棘手。
- *意图不佳*。虽然输出应该是 JSON，但真正重要的是数据究竟是什么，而不是它的编码方式。
- *重新测试标准库*。不需要测试标准库如何输出 JSON，它已经测试过了。不要测试其他人的代码。

Instead, we should look to parse the JSON into data structures that are relevant for us to test with.

相反，我们应该将 JSON 解析为与我们测试相关的数据结构。

### Data modelling

### 数据建模

Given the JSON data model, it looks like we need an array of `Player` with some fields so we have created a new type to capture this.

给定 JSON 数据模型，看起来我们需要一个包含一些字段的 `Player` 数组，因此我们创建了一个新类型来捕获它。

```go
//server.go
type Player struct {
    Name string
    Wins int
}
```

### JSON decoding

### JSON 解码

```go
//server_test.go
var got []Player
err := json.NewDecoder(response.Body).Decode(&got)
```

To parse JSON into our data model we create a `Decoder` from `encoding/json` package and then call its `Decode` method. To create a `Decoder` it needs an `io.Reader` to read from which in our case is our response spy's `Body`.

为了将 JSON 解析为我们的数据模型，我们从 `encoding/json` 包中创建了一个 `Decoder`，然后调用它的 `Decode` 方法。要创建一个 `Decoder`，它需要一个 `io.Reader` 来读取，在我们的例子中是我们的响应间谍的 `Body`。

`Decode` takes the address of the thing we are trying to decode into which is why we declare an empty slice of `Player` the line before.

`Decode` 获取我们试图解码的东西的地址，这就是为什么我们在前一行声明一个空的 `Player` 切片。

Parsing JSON can fail so `Decode` can return an `error`. There's no point continuing the test if that fails so we check for the error and stop the test with `t.Fatalf` if it happens. Notice that we print the response body along with the error as it's important for someone running the test to see what string cannot be parsed.

解析 JSON 可能会失败，因此 `Decode` 会返回一个 `error`。如果失败，则继续测试没有意义，因此我们检查错误并在发生时使用 `t.Fatalf` 停止测试。请注意，我们将响应正文与错误一起打印，因为对于运行测试的人来说，查看哪些字符串无法解析很重要。

## Try to run the test

## 尝试运行测试

```
=== RUN   TestLeague/it_returns_200_on_/league
    --- FAIL: TestLeague/it_returns_200_on_/league (0.00s)
        server_test.go:107: Unable to parse response from server '' into slice of Player, 'unexpected end of JSON input'
```

Our endpoint currently does not return a body so it cannot be parsed into JSON.

我们的端点目前不返回正文，因此无法将其解析为 JSON。

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
//server.go
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    leagueTable := []Player{
        {"Chris", 20},
    }

    json.NewEncoder(w).Encode(leagueTable)

    w.WriteHeader(http.StatusOK)
}
```

The test now passes.

现在测试通过了。

### Encoding and Decoding

### 编码和解码

Notice the lovely symmetry in the standard library.

注意标准库中可爱的对称性。

- To create an `Encoder` you need an `io.Writer` which is what `http.ResponseWriter` implements.
- To create a `Decoder` you need an `io.Reader` which the `Body` field of our response spy implements. 

- 要创建一个 `Encoder`，你需要一个 `io.Writer`，它是 `http.ResponseWriter` 实现的。
- 要创建一个 `Decoder`，您需要一个 `io.Reader`，我们的响应间谍的 `Body` 字段实现了它。

Throughout this book, we have used `io.Writer` and this is another demonstration of its prevalence in the standard library and how a lot of libraries easily work with it.

在整本书中，我们使用了 io.Writer，这是它在标准库中的流行以及许多库如何轻松使用它的另一个演示。

## Refactor

## 重构

It would be nice to introduce a separation of concern between our handler and getting the `leagueTable` as we know we're going to not hard-code that very soon.

在我们的处理程序和获取 `leagueTable` 之间引入关注点分离会很好，因为我们知道我们不会很快对其进行硬编码。

```go
//server.go
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(p.getLeagueTable())
    w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getLeagueTable() []Player {
    return []Player{
        {"Chris", 20},
    }
}
```

Next, we'll want to extend our test so that we can control exactly what data we want back.

接下来，我们要扩展我们的测试，以便我们可以准确地控制我们想要返回的数据。

## Write the test first

## 先写测试

We can update the test to assert that the league table contains some players that we will stub in our store.

我们可以更新测试以断言联赛表包含一些我们将在我们的商店中存根的球员。

Update `StubPlayerStore` to let it store a league, which is just a slice of `Player`. We'll store our expected data in there.

更新 `StubPlayerStore` 让它存储一个联赛，它只是 `Player` 的一部分。我们将在那里存储我们预期的数据。

```go
//server_test.go
type StubPlayerStore struct {
    scores   map[string]int
    winCalls []string
    league   []Player
}
```

Next, update our current test by putting some players in the league property of our stub and assert they get returned from our server.

接下来，通过将一些玩家放入我们存根的 League 属性中并断言他们从我们的服务器返回来更新我们当前的测试。

```go
//server_test.go
func TestLeague(t *testing.T) {

    t.Run("it returns the league table as JSON", func(t *testing.T) {
        wantedLeague := []Player{
            {"Cleo", 32},
            {"Chris", 20},
            {"Tiest", 14},
        }

        store := StubPlayerStore{nil, nil, wantedLeague}
        server := NewPlayerServer(&store)

        request, _ := http.NewRequest(http.MethodGet, "/league", nil)
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        var got []Player

        err := json.NewDecoder(response.Body).Decode(&got)

        if err != nil {
            t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
        }

        assertStatus(t, response.Code, http.StatusOK)

        if !reflect.DeepEqual(got, wantedLeague) {
            t.Errorf("got %v want %v", got, wantedLeague)
        }
    })
}
```

## Try to run the test

## 尝试运行测试

```
./server_test.go:33:3: too few values in struct initializer
./server_test.go:70:3: too few values in struct initializer
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

You'll need to update the other tests as we have a new field in `StubPlayerStore`; set it to nil for the other tests.

您需要更新其他测试，因为我们在 `StubPlayerStore` 中有一个新字段；对于其他测试，将其设置为 nil。

Try running the tests again and you should get

再次尝试运行测试，你应该得到

```
=== RUN   TestLeague/it_returns_the_league_table_as_JSON
    --- FAIL: TestLeague/it_returns_the_league_table_as_JSON (0.00s)
        server_test.go:124: got [{Chris 20}] want [{Cleo 32} {Chris 20} {Tiest 14}]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We know the data is in our `StubPlayerStore` and we've abstracted that away into an interface `PlayerStore`. We need to update this so anyone passing us in a `PlayerStore` can provide us with the data for leagues.

我们知道数据在我们的 `StubPlayerStore` 中，我们已经将其抽象为一个接口 `PlayerStore`。我们需要更新它，以便任何在 `PlayerStore` 中通过我们的人都可以向我们提供联赛数据。

```go
//server.go
type PlayerStore interface {
    GetPlayerScore(name string) int
    RecordWin(name string)
    GetLeague() []Player
}
```

Now we can update our handler code to call that rather than returning a hard-coded list. Delete our method `getLeagueTable()` and then update `leagueHandler` to call `GetLeague()`.

现在我们可以更新我们的处理程序代码来调用它而不是返回一个硬编码的列表。删除我们的方法 `getLeagueTable()`，然后更新 `leagueHandler` 以调用 `GetLeague()`。

```go
//server.go
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(p.store.GetLeague())
    w.WriteHeader(http.StatusOK)
}
```

Try and run the tests.

尝试并运行测试。

```
# github.com/quii/learn-go-with-tests/json-and-io/v4
./main.go:9:50: cannot use NewInMemoryPlayerStore() (type *InMemoryPlayerStore) as type PlayerStore in argument to NewPlayerServer:
    *InMemoryPlayerStore does not implement PlayerStore (missing GetLeague method)
./server_integration_test.go:11:27: cannot use store (type *InMemoryPlayerStore) as type PlayerStore in argument to NewPlayerServer:
    *InMemoryPlayerStore does not implement PlayerStore (missing GetLeague method)
./server_test.go:36:28: cannot use &store (type *StubPlayerStore) as type PlayerStore in argument to NewPlayerServer:
    *StubPlayerStore does not implement PlayerStore (missing GetLeague method)
./server_test.go:74:28: cannot use &store (type *StubPlayerStore) as type PlayerStore in argument to NewPlayerServer:
    *StubPlayerStore does not implement PlayerStore (missing GetLeague method)
./server_test.go:106:29: cannot use &store (type *StubPlayerStore) as type PlayerStore in argument to NewPlayerServer:
    *StubPlayerStore does not implement PlayerStore (missing GetLeague method)
```

The compiler is complaining because `InMemoryPlayerStore` and `StubPlayerStore` do not have the new method we added to our interface.

编译器抱怨是因为 `InMemoryPlayerStore` 和 `StubPlayerStore` 没有我们添加到界面中的新方法。

For `StubPlayerStore` it's pretty easy, just return the `league` field we added earlier.

对于 `StubPlayerStore` 来说很简单，只需返回我们之前添加的 `league` 字段即可。

```go
//server_test.go
func (s *StubPlayerStore) GetLeague() []Player {
    return s.league
}
```

Here's a reminder of how `InMemoryStore` is implemented.

这里提醒大家如何实现 `InMemoryStore`。

```go
//in_memory_player_store.go
type InMemoryPlayerStore struct {
    store map[string]int
}
```

Whilst it would be pretty straightforward to implement `GetLeague` "properly" by iterating over the map remember we are just trying to _write the minimal amount of code to make the tests pass_.

虽然通过遍历地图“正确地”实现`GetLeague` 会非常简单，但请记住，我们只是试图_编写最少量的代码以使测试通过_。

So let's just get the compiler happy for now and live with the uncomfortable feeling of an incomplete implementation in our `InMemoryStore`.

因此，让我们暂时让编译器满意，并忍受我们的 InMemoryStore 中实现不完整的不舒服感觉。

```go
//in_memory_player_store.go
func (i *InMemoryPlayerStore) GetLeague() []Player {
    return nil
}
```

What this is really telling us is that _later_ we're going to want to test this but let's park that for now.

这真正告诉我们的是_稍后_我们将要对此进行测试，但让我们暂时停止。

Try and run the tests, the compiler should pass and the tests should be passing!

尝试运行测试，编译器应该通过并且测试应该通过！

## Refactor

## 重构

The test code does not convey our intent very well and has a lot of boilerplate we can refactor away.

测试代码并没有很好地传达我们的意图，并且有很多我们可以重构的样板。

```go
//server_test.go
t.Run("it returns the league table as JSON", func(t *testing.T) {
    wantedLeague := []Player{
        {"Cleo", 32},
        {"Chris", 20},
        {"Tiest", 14},
    }

    store := StubPlayerStore{nil, nil, wantedLeague}
    server := NewPlayerServer(&store)

    request := newLeagueRequest()
    response := httptest.NewRecorder()

    server.ServeHTTP(response, request)

    got := getLeagueFromResponse(t, response.Body)
    assertStatus(t, response.Code, http.StatusOK)
    assertLeague(t, got, wantedLeague)
})
```

Here are the new helpers

这里是新帮手

```go
//server_test.go
func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
    t.Helper()
    err := json.NewDecoder(body).Decode(&league)

    if err != nil {
        t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
    }

    return
}

func assertLeague(t testing.TB, got, want []Player) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func newLeagueRequest() *http.Request {
    req, _ := http.NewRequest(http.MethodGet, "/league", nil)
    return req
}
```

One final thing we need to do for our server to work is make sure we return a `content-type` header in the response so machines can recognise we are returning `JSON`.

我们需要为我们的服务器工作做的最后一件事是确保我们在响应中返回一个 `content-type` 标头，以便机器可以识别我们正在返回 `JSON`。

## Write the test first

## 先写测试

Add this assertion to the existing test

将此断言添加到现有测试中

```go
//server_test.go
if response.Result().Header.Get("content-type") != "application/json" {
    t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
}
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestLeague/it_returns_the_league_table_as_JSON
    --- FAIL: TestLeague/it_returns_the_league_table_as_JSON (0.00s)
        server_test.go:124: response did not have content-type of application/json, got map[Content-Type:[text/plain;charset=utf-8]]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Update `leagueHandler`

更新`leagueHandler`

```go
//server.go
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    json.NewEncoder(w).Encode(p.store.GetLeague())
}
```

The test should pass.

测试应该通过。

## Refactor

## 重构

Create a constant for "application/json" and use it in `leagueHandler`

为“application/json”创建一个常量并在`leagueHandler`中使用它

```go
//server.go
const jsonContentType = "application/json"

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", jsonContentType)
    json.NewEncoder(w).Encode(p.store.GetLeague())
}
```

Then add a helper for `assertContentType`.

然后为 `assertContentType` 添加一个助手。

```go
//server_test.go
func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
    t.Helper()
    if response.Result().Header.Get("content-type") != want {
        t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
    }
}
```

Use it in the test.

在测试中使用它。

```go
//server_test.go
assertContentType(t, response, jsonContentType)
```

Now that we have sorted out `PlayerServer` for now we can turn our attention to `InMemoryPlayerStore` because right now if we tried to demo this to the product owner `/league` will not work.

现在我们已经整理了 `PlayerServer`，现在我们可以将注意力转向 `InMemoryPlayerStore`，因为现在如果我们试图向产品负责人 `/league` 演示它是行不通的。

The quickest way for us to get some confidence is to add to our integration test, we can hit the new endpoint and check we get back the correct response from `/league`.

让我们获得信心的最快方法是添加到我们的集成测试中，我们可以点击新的端点并检查我们是否从 `/league` 返回了正确的响应。

## Write the test first 

## 先写测试

We can use `t.Run` to break up this test a bit and we can reuse the helpers from our server tests - again showing the importance of refactoring tests.

我们可以使用 `t.Run` 来稍微分解这个测试，我们可以重用来自我们服务器测试的助手 - 再次表明重构测试的重要性。

```go
//server_integration_test.go
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
    store := NewInMemoryPlayerStore()
    server := NewPlayerServer(store)
    player := "Pepper"

    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

    t.Run("get score", func(t *testing.T) {
        response := httptest.NewRecorder()
        server.ServeHTTP(response, newGetScoreRequest(player))
        assertStatus(t, response.Code, http.StatusOK)

        assertResponseBody(t, response.Body.String(), "3")
    })

    t.Run("get league", func(t *testing.T) {
        response := httptest.NewRecorder()
        server.ServeHTTP(response, newLeagueRequest())
        assertStatus(t, response.Code, http.StatusOK)

        got := getLeagueFromResponse(t, response.Body)
        want := []Player{
            {"Pepper", 3},
        }
        assertLeague(t, got, want)
    })
}
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestRecordingWinsAndRetrievingThem/get_league
    --- FAIL: TestRecordingWinsAndRetrievingThem/get_league (0.00s)
        server_integration_test.go:35: got [] want [{Pepper 3}]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

`InMemoryPlayerStore` is returning `nil` when you call `GetLeague()` so we'll need to fix that.

`InMemoryPlayerStore` 在调用 `GetLeague()` 时返回 `nil`，因此我们需要解决这个问题。

```go
//in_memory_player_store.go
func (i *InMemoryPlayerStore) GetLeague() []Player {
    var league []Player
    for name, wins := range i.store {
        league = append(league, Player{name, wins})
    }
    return league
}
```

All we need to do is iterate over the map and convert each key/value to a `Player`.

我们需要做的就是遍历地图并将每个键/值转换为“播放器”。

The test should now pass.

测试现在应该通过。

## Wrapping up

##  总结

We've continued to safely iterate on our program using TDD, making it support new endpoints in a maintainable way with a router and it can now return JSON for our consumers. In the next chapter, we will cover persisting the data and sorting our league.

我们继续使用 TDD 安全地迭代我们的程序，使其通过路由器以可维护的方式支持新端点，现在它可以为我们的消费者返回 JSON。在下一章中，我们将介绍持久化数据和对我们的联赛进行排序。

What we've covered:

我们涵盖的内容：

- **Routing**. The standard library offers you an easy to use type to do routing. It fully embraces the `http.Handler` interface in that you assign routes to `Handler`s and the router itself is also a `Handler`. It does not have some features you might expect though such as path variables (e.g `/users/{id}`). You can easily parse this information yourself but you might want to consider looking at other routing libraries if it becomes a burden. Most of the popular ones stick to the standard library's philosophy of also implementing `http.Handler`.
- **Type embedding**. We touched a little on this technique but you can [learn more about it from Effective Go](https://golang.org/doc/effective_go.html#embedding). If there is one thing you should take away from this is that it can be extremely useful but _always thinking about your public API, only expose what's appropriate_.
- **JSON deserializing and serializing**. The standard library makes it very trivial to serialise and deserialise your data. It is also open to configuration and you can customise how these data transformations work if necessary. 

- **路由**。标准库为您提供了一种易于使用的类型来进行路由。它完全包含了 `http.Handler` 接口，因为你可以将路由分配给 `Handler`，而路由器本身也是一个 `Handler`。它没有您可能期望的某些功能，例如路径变量（例如`/users/{id}`）。您可以自己轻松解析此信息，但如果它成为负担，您可能需要考虑查看其他路由库。大多数流行的都坚持标准库的理念，即也实现了`http.Handler`。
- **类型嵌入**。我们对这项技术略有触及，但您可以 [从 Effective Go 了解更多信息](https://golang.org/doc/effective_go.html#embedding)。如果有一件事情你应该从中吸取教训，那就是它可能非常有用，但_总是考虑你的公共 API，只公开什么是合适的_。
- **JSON 反序列化和序列化**。标准库使得序列化和反序列化数据变得非常简单。它还对配置开放，您可以在必要时自定义这些数据转换的工作方式。

