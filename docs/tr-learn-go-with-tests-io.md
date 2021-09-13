# IO and sorting

# IO 和排序

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/io)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/io)**

[In the previous chapter](json.md) we continued iterating on our application by adding a new endpoint `/league`. Along the way we learned about how to deal with JSON, embedding types and routing.

[在前一章中](json.md) 我们通过添加一个新的端点`/league` 继续迭代我们的应用程序。在此过程中，我们学习了如何处理 JSON、嵌入类型和路由。

Our product owner is somewhat perturbed by the software losing the scores when the server was restarted. This is because our implementation of our store is in-memory. She is also not pleased that we didn't interpret the `/league` endpoint should return the players ordered by the number of wins!

我们的产品负责人对在服务器重新启动时丢失分数的软件感到有些不安。这是因为我们的 store 的实现是在内存中的。她也不高兴我们没有解释 `/league` 端点应该返回按获胜次数排序的玩家！

## The code so far

## 到目前为止的代码

```go
// server.go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

// PlayerStore stores score information about players
type PlayerStore interface {
    GetPlayerScore(name string) int
    RecordWin(name string)
    GetLeague() []Player
}

// Player stores a name with a number of wins
type Player struct {
    Name string
    Wins int
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
    store PlayerStore
    http.Handler
}

const jsonContentType = "application/json"

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := new(PlayerServer)

    p.store = store

    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))

    p.Handler = router

    return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", jsonContentType)
    json.NewEncoder(w).Encode(p.store.GetLeague())
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
// in_memory_player_store.go
package main

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
    return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
    store map[string]int
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
    var league []Player
    for name, wins := range i.store {
        league = append(league, Player{name, wins})
    }
    return league
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
    server := NewPlayerServer(NewInMemoryPlayerStore())
    log.Fatal(http.ListenAndServe(":5000", server))
}
```

You can find the corresponding tests in the link at the top of the chapter.

您可以在本章顶部的链接中找到相应的测试。

## Store the data

## 存储数据

There are dozens of databases we could use for this but we're going to go for a very simple approach. We're going to store the data for this application in a file as JSON.

我们可以使用数十种数据库，但我们将采用一种非常简单的方法。我们将把这个应用程序的数据作为 JSON 存储在一个文件中。

This keeps the data very portable and is relatively simple to implement.

这使数据非常便携并且实现起来相对简单。

It won't scale especially well but given this is a prototype it'll be fine for now. If our circumstances change and it's no longer appropriate it'll be simple to swap it out for something different because of the `PlayerStore` abstraction we have used.

它的扩展性不会特别好，但鉴于这是一个原型，现在还可以。如果我们的情况发生变化并且不再合适，由于我们使用了 `PlayerStore` 抽象，可以很容易地将其替换为不同的东西。

We will keep the `InMemoryPlayerStore` for now so that the integration tests keep passing as we develop our new store. Once we are confident our new implementation is sufficient to make the integration test pass we will swap it in and then delete `InMemoryPlayerStore`.

我们现在将保留 `InMemoryPlayerStore`，以便在我们开发新商店时继续通过集成测试。一旦我们确信我们的新实现足以使集成测试通过，我们将把它交换进去，然后删除 `InMemoryPlayerStore`。

## Write the test first

## 先写测试

By now you should be familiar with the interfaces around the standard library for reading data (`io.Reader`), writing data (`io.Writer`) and how we can use the standard library to test these functions without having to use real files. 

现在你应该已经熟悉了标准库中读取数据（`io.Reader`）、写入数据（`io.Writer`）的接口，以及我们如何使用标准库来测试这些功能，而无需使用真正的文件。

For this work to be complete we'll need to implement `PlayerStore` so we'll write tests for our store calling the methods we need to implement. We'll start with `GetLeague`.

为了完成这项工作，我们需要实现`PlayerStore`，因此我们将为我们的商店编写测试，调用我们需要实现的方法。我们将从`GetLeague`开始。

```go
//file_system_store_test.go
func TestFileSystemStore(t *testing.T) {

    t.Run("league from a reader", func(t *testing.T) {
        database := strings.NewReader(`[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

        store := FileSystemPlayerStore{database}

        got := store.GetLeague()

        want := []Player{
            {"Cleo", 10},
            {"Chris", 33},
        }

        assertLeague(t, got, want)
    })
}
```

We're using `strings.NewReader` which will return us a `Reader`, which is what our `FileSystemPlayerStore` will use to read data. In `main` we will open a file, which is also a `Reader`.

我们正在使用`strings.NewReader`，它将返回一个`Reader`，这是我们的`FileSystemPlayerStore` 将用来读取数据的。在 `main` 中，我们将打开一个文件，它也是一个 `Reader`。

## Try to run the test

## 尝试运行测试

```
# github.com/quii/learn-go-with-tests/io/v1
./file_system_store_test.go:15:12: undefined: FileSystemPlayerStore
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Let's define `FileSystemPlayerStore` in a new file

让我们在一个新文件中定义 `FileSystemPlayerStore`

```go
//file_system_store.go
type FileSystemPlayerStore struct {}
```

Try again

再试一次

```
# github.com/quii/learn-go-with-tests/io/v1
./file_system_store_test.go:15:28: too many values in struct initializer
./file_system_store_test.go:17:15: store.GetLeague undefined (type FileSystemPlayerStore has no field or method GetLeague)
```

It's complaining because we're passing in a `Reader` but not expecting one and it doesn't have `GetLeague` defined yet.

它在抱怨是因为我们传入了一个 `Reader` 但并不期待它，而且它还没有定义 `GetLeague`。

```go
//file_system_store.go
type FileSystemPlayerStore struct {
    database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
    return nil
}
```

One more try...

再试一次...

```
=== RUN   TestFileSystemStore//league_from_a_reader
    --- FAIL: TestFileSystemStore//league_from_a_reader (0.00s)
        file_system_store_test.go:24: got [] want [{Cleo 10} {Chris 33}]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We've read JSON from a reader before

我们以前从读者那里读过 JSON

```go
//file_system_store.go
func (f *FileSystemPlayerStore) GetLeague() []Player {
    var league []Player
    json.NewDecoder(f.database).Decode(&league)
    return league
}
```

The test should pass.

测试应该通过。

## Refactor

## 重构

We _have_ done this before! Our test code for the server had to decode the JSON from the response.

我们以前_have_这样做过！我们的服务器测试代码必须从响应中解码 JSON。

Let's try DRYing this up into a function.

让我们尝试将它干燥成一个函数。

Create a new file called `league.go` and put this inside.

创建一个名为“league.go”的新文件并将其放入其中。

```go
//league.go
func NewLeague(rdr io.Reader) ([]Player, error) {
    var league []Player
    err := json.NewDecoder(rdr).Decode(&league)
    if err != nil {
        err = fmt.Errorf("problem parsing league, %v", err)
    }

    return league, err
}
```

Call this in our implementation and in our test helper `getLeagueFromResponse` in `server_test.go`

在我们的实现和我们的测试助手 `server_test.go` 中的 `getLeagueFromResponse` 中调用它

```go
//file_system_store.go
func (f *FileSystemPlayerStore) GetLeague() []Player {
    league, _ := NewLeague(f.database)
    return league
}
```

We haven't got a strategy yet for dealing with parsing errors but let's press on.

我们还没有处理解析错误的策略，但让我们继续。

### Seeking problems

### 求问题

There is a flaw in our implementation. First of all, let's remind ourselves how `io.Reader` is defined.

我们的实施存在缺陷。首先，让我们提醒自己 `io.Reader` 是如何定义的。

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

With our file, you can imagine it reading through byte by byte until the end. What happens if you try to `Read` a second time?

使用我们的文件，您可以想象它逐字节读取直到结束。如果您再次尝试“阅读”会发生什么？

Add the following to the end of our current test.

将以下内容添加到我们当前测试的末尾。

```go
//file_system_store_test.go

// read again
got = store.GetLeague()
assertLeague(t, got, want)
```

We want this to pass, but if you run the test it doesn't.

我们希望它通过，但如果您运行测试，它不会通过。

The problem is our `Reader` has reached the end so there is nothing more to read. We need a way to tell it to go back to the start.

问题是我们的“Reader”已经读完了，所以没有什么可读的了。我们需要一种方法来告诉它回到起点。

[ReadSeeker](https://golang.org/pkg/io/#ReadSeeker) is another interface in the standard library that can help.

[ReadSeeker](https://golang.org/pkg/io/#ReadSeeker) 是标准库中另一个可以提供帮助的接口。

```go
type ReadSeeker interface {
    Reader
    Seeker
}
```

Remember embedding? This is an interface comprised of `Reader` and [`Seeker`](https://golang.org/pkg/io/#Seeker)

还记得嵌入吗？这是一个由`Reader`和[`Seeker`](https://golang.org/pkg/io/#Seeker)组成的界面

```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

This sounds good, can we change `FileSystemPlayerStore` to take this interface instead?

这听起来不错，我们可以将`FileSystemPlayerStore` 改成这个接口吗？

```go
//file_system_store.go
type FileSystemPlayerStore struct {
    database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
    f.database.Seek(0, 0)
    league, _ := NewLeague(f.database)
    return league
}
```

Try running the test, it now passes! Happily for us `string.NewReader` that we used in our test also implements `ReadSeeker` so we didn't have to make any other changes.

尝试运行测试，它现在通过了！令我们高兴的是，我们在测试中使用的 `string.NewReader` 也实现了 `ReadSeeker`，因此我们无需进行任何其他更改。

Next we'll implement `GetPlayerScore`.

接下来我们将实现`GetPlayerScore`。

## Write the test first

## 先写测试

```go
//file_system_store_test.go
t.Run("get player score", func(t *testing.T) {
    database := strings.NewReader(`[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

    store := FileSystemPlayerStore{database}

    got := store.GetPlayerScore("Chris")

    want := 33

    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
})
```

## Try to run the test

## 尝试运行测试

```
./file_system_store_test.go:38:15: store.GetPlayerScore undefined (type FileSystemPlayerStore has no field or method GetPlayerScore)
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We need to add the method to our new type to get the test to compile.

我们需要将方法添加到我们的新类型中以编译测试。

```go
//file_system_store.go
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
    return 0
}
```

Now it compiles and the test fails

现在它编译并且测试失败

```
=== RUN   TestFileSystemStore/get_player_score
    --- FAIL: TestFileSystemStore//get_player_score (0.00s)
        file_system_store_test.go:43: got 0 want 33
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We can iterate over the league to find the player and return their score

我们可以遍历联盟来找到球员并返回他们的分数

```go
//file_system_store.go
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

    var wins int

    for _, player := range f.GetLeague() {
        if player.Name == name {
            wins = player.Wins
            break
        }
    }

    return wins
}
```

## Refactor

## 重构

You will have seen dozens of test helper refactorings so I'll leave this to you to make it work

你会看到几十个测试助手重构，所以我会把它留给你来让它工作

```go
//file_system_store_test.go
t.Run("get player score", func(t *testing.T) {
    database := strings.NewReader(`[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

    store := FileSystemPlayerStore{database}

    got := store.GetPlayerScore("Chris")
    want := 33
    assertScoreEquals(t, got, want)
})
```

Finally, we need to start recording scores with `RecordWin`.

最后，我们需要用`RecordWin` 开始记录乐谱。

## Write the test first

## 先写测试

Our approach is fairly short-sighted for writes. We can't (easily) just update one "row" of JSON in a file. We'll need to store the _whole_ new representation of our database on every write.

我们的方法对于写入是相当短视的。我们不能（轻松地）只更新文件中的一行 JSON。我们需要在每次写入时存储我们数据库的 _whole_ 新表示。

How do we write? We'd normally use a `Writer` but we already have our `ReadSeeker`. Potentially we could have two dependencies but the standard library already has an interface for us `ReadWriteSeeker` which lets us do all the things we'll need to do with a file.

我们怎么写？我们通常使用`Writer`，但我们已经有了`ReadSeeker`。可能我们可以有两个依赖项，但标准库已经为我们提供了一个接口 `ReadWriteSeeker`，它让我们可以对文件进行所有需要做的事情。

Let's update our type

让我们更新我们的类型

```go
//file_system_store.go
type FileSystemPlayerStore struct {
    database io.ReadWriteSeeker
}
```

See if it compiles

看看能不能编译

```go
./file_system_store_test.go:15:34: cannot use database (type *strings.Reader) as type io.ReadWriteSeeker in field value:
    *strings.Reader does not implement io.ReadWriteSeeker (missing Write method)
./file_system_store_test.go:36:34: cannot use database (type *strings.Reader) as type io.ReadWriteSeeker in field value:
    *strings.Reader does not implement io.ReadWriteSeeker (missing Write method)
```

It's not too surprising that `strings.Reader` does not implement `ReadWriteSeeker` so what do we do?

`strings.Reader` 没有实现 `ReadWriteSeeker` 这并不奇怪，那么我们该怎么办呢？

We have two choices

我们有两个选择

- Create a temporary file for each test. `*os.File` implements `ReadWriteSeeker`. The pro of this is it becomes more of an integration test, we're really reading and writing from the file system so it will give us a very high level of confidence. The cons are we prefer unit tests because they are faster and generally simpler. We will also need to do more work around creating temporary files and then making sure they're removed after the test.
- We could use a third party library. [Mattetti](https://github.com/mattetti) has written a library [filebuffer](https://github.com/mattetti/filebuffer) which implements the interface we need and doesn't touch the file system.

- 为每个测试创建一个临时文件。 `*os.File` 实现了 `ReadWriteSeeker`。这样做的好处是它更像是一个集成测试，我们实际上是从文件系统读取和写入，所以它会给我们很高的信心。缺点是我们更喜欢单元测试，因为它们更快且通常更简单。我们还需要做更多的工作来创建临时文件，然后确保它们在测试后被删除。
- 我们可以使用第三方库。 [Mattetti](https://github.com/mattetti)编写了一个库[filebuffer](https://github.com/mattetti/filebuffer)，它实现了我们需要的接口并且不涉及文件系统。

I don't think there's an especially wrong answer here, but by choosing to use a third party library I would have to explain dependency management! So we will use files instead.

我不认为这里有一个特别错误的答案，但是通过选择使用第三方库，我将不得不解释依赖管理！所以我们将使用文件代替。

Before adding our test we need to make our other tests compile by replacing the `strings.Reader` with an `os.File`.

在添加我们的测试之前，我们需要通过用 `os.File` 替换 `strings.Reader` 来编译我们的其他测试。

Let's create a helper function which will create a temporary file with some data inside it

让我们创建一个辅助函数，它将创建一个包含一些数据的临时文件

```go
//file_system_store_test.go
func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
    t.Helper()

    tmpfile, err := ioutil.TempFile("", "db")

    if err != nil {
        t.Fatalf("could not create temp file %v", err)
    }

    tmpfile.Write([]byte(initialData))

    removeFile := func() {
        tmpfile.Close()
        os.Remove(tmpfile.Name())
    }

    return tmpfile, removeFile
}
```

[TempFile](https://golang.org/pkg/io/ioutil/#TempDir) creates a temporary file for us to use. The `"db"` value we've passed in is a prefix put on a random file name it will create. This is to ensure it won't clash with other files by accident.

[TempFile](https://golang.org/pkg/io/ioutil/#TempDir) 创建一个临时文件供我们使用。我们传入的 `"db"` 值是放在它将创建的随机文件名上的前缀。这是为了确保它不会意外地与其他文件发生冲突。

You'll notice we're not only returning our `ReadWriteSeeker` (the file) but also a function. We need to make sure that the file is removed once the test is finished. We don't want to leak details of the files into the test as it's prone to error and uninteresting for the reader. By returning a `removeFile` function, we can take care of the details in our helper and all the caller has to do is run `defer cleanDatabase()`.

您会注意到我们不仅返回了我们的 `ReadWriteSeeker`（文件），而且还返回了一个函数。我们需要确保在测试完成后删除该文件。我们不想将文件的详细信息泄露到测试中，因为它容易出错并且对读者来说无趣。通过返回一个 `removeFile` 函数，我们可以处理我们的助手中的细节，调用者所要做的就是运行 `defer cleanDatabase()`。

```go
//file_system_store_test.go
func TestFileSystemStore(t *testing.T) {

    t.Run("league from a reader", func(t *testing.T) {
        database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
        defer cleanDatabase()

        store := FileSystemPlayerStore{database}

        got := store.GetLeague()

        want := []Player{
            {"Cleo", 10},
            {"Chris", 33},
        }

        assertLeague(t, got, want)

        // read again
        got = store.GetLeague()
        assertLeague(t, got, want)
    })

    t.Run("get player score", func(t *testing.T) {
        database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
        defer cleanDatabase()

        store := FileSystemPlayerStore{database}

        got := store.GetPlayerScore("Chris")
        want := 33
        assertScoreEquals(t, got, want)
    })
}
```

Run the tests and they should be passing! There were a fair amount of changes but now it feels like we have our interface definition complete and it should be very easy to add new tests from now.

运行测试，它们应该会通过！有相当多的变化，但现在感觉我们的接口定义已经完成，从现在开始添加新测试应该很容易。

Let's get the first iteration of recording a win for an existing player

让我们开始记录现有玩家获胜的第一次迭代

```go
//file_system_store_test.go
t.Run("store wins for existing players", func(t *testing.T) {
    database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
    defer cleanDatabase()

    store := FileSystemPlayerStore{database}

    store.RecordWin("Chris")

    got := store.GetPlayerScore("Chris")
    want := 34
    assertScoreEquals(t, got, want)
})
```

## Try to run the test

## 尝试运行测试

`./file_system_store_test.go:67:8: store.RecordWin undefined (type FileSystemPlayerStore has no field or method RecordWin)`

`./file_system_store_test.go:67:8: store.RecordWin 未定义（类型 FileSystemPlayerStore 没有字段或方法 RecordWin）`

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Add the new method

添加新方法

```go
//file_system_store.go
func (f *FileSystemPlayerStore) RecordWin(name string) {

}
```

```
=== RUN   TestFileSystemStore/store_wins_for_existing_players
    --- FAIL: TestFileSystemStore/store_wins_for_existing_players (0.00s)
        file_system_store_test.go:71: got 33 want 34
```

Our implementation is empty so the old score is getting returned.

我们的实现是空的，所以旧的分数被返回。

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
//file_system_store.go
func (f *FileSystemPlayerStore) RecordWin(name string) {
    league := f.GetLeague()

    for i, player := range league {
        if player.Name == name {
            league[i].Wins++
        }
    }

    f.database.Seek(0,0)
    json.NewEncoder(f.database).Encode(league)
}
```

You may be asking yourself why I am doing `league[i].Wins++` rather than `player.Wins++`.

你可能会问自己为什么我在做 `league[i].Wins++` 而不是 `player.Wins++`。

When you `range` over a slice you are returned the current index of the loop (in our case `i`) and a _copy_ of the element at that index. Changing the `Wins` value of a copy won't have any effect on the `league` slice that we iterate on. For that reason, we need to get the reference to the actual value by doing `league[i]` and then changing that value instead.

当您在切片上“范围”时，您将返回循环的当前索引（在我们的例子中为“i”）和该索引处元素的 _copy_。更改副本的 `Wins` 值不会对我们迭代的 `league` 切片产生任何影响。出于这个原因，我们需要通过执行 `league[i]` 来获取对实际值的引用，然后改为更改该值。

If you run the tests, they should now be passing.

如果您运行测试，它们现在应该会通过。

## Refactor

## 重构

In `GetPlayerScore` and `RecordWin`, we are iterating over `[]Player` to find a player by name. 

在`GetPlayerScore` 和`RecordWin` 中，我们遍历`[]Player` 以按名称查找玩家。

We could refactor this common code in the internals of `FileSystemStore` but to me, it feels like this is maybe useful code we can lift into a new type. Working with a "League" so far has always been with `[]Player` but we can create a new type called `League`. This will be easier for other developers to understand and then we can attach useful methods onto that type for us to use.

我们可以在 `FileSystemStore` 的内部重构这个通用代码，但对我来说，感觉这可能是我们可以提升到新类型的有用代码。到目前为止，使用“League”一直是使用 `[]Player`，但我们可以创建一个名为 `League` 的新类型。这会让其他开发人员更容易理解，然后我们可以将有用的方法附加到该类型上供我们使用。

Inside `league.go` add the following

在 `league.go` 中添加以下内容

```go
//league.go
type League []Player

func (l League) Find(name string) *Player {
    for i, p := range l {
        if p.Name==name {
            return &l[i]
        }
    }
    return nil
}
```

Now if anyone has a `League` they can easily find a given player.

现在，如果有人拥有“联赛”，他们就可以轻松找到给定的球员。

Change our `PlayerStore` interface to return `League` rather than `[]Player`. Try to re-run the tests, you'll get a compilation problem because we've changed the interface but it's very easy to fix; just change the return type from `[]Player` to `League`.

更改我们的 `PlayerStore` 接口以返回 `League` 而不是 `[]Player`。尝试重新运行测试，您会遇到编译问题，因为我们更改了界面，但很容易修复；只需将返回类型从 `[]Player` 更改为 `League`。

This lets us simplify our methods in `file_system_store`.

这让我们可以简化我们在 `file_system_store` 中的方法。

```go
//file_system_store.go
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

    player := f.GetLeague().Find(name)

    if player != nil {
        return player.Wins
    }

    return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
    league := f.GetLeague()
    player := league.Find(name)

    if player != nil {
        player.Wins++
    }

    f.database.Seek(0, 0)
    json.NewEncoder(f.database).Encode(league)
}
```

This is looking much better and we can see how we might be able to find other useful functionality around `League` can be refactored.

这看起来好多了，我们可以看到我们如何能够找到有关“联盟”的其他有用功能可以重构。

We now need to handle the scenario of recording wins of new players.

我们现在需要处理记录新玩家获胜的场景。

## Write the test first

## 先写测试

```go
//file_system_store_test.go
t.Run("store wins for new players", func(t *testing.T) {
    database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
    defer cleanDatabase()

    store := FileSystemPlayerStore{database}

    store.RecordWin("Pepper")

    got := store.GetPlayerScore("Pepper")
    want := 1
    assertScoreEquals(t, got, want)
})
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestFileSystemStore/store_wins_for_new_players#01
    --- FAIL: TestFileSystemStore/store_wins_for_new_players#01 (0.00s)
        file_system_store_test.go:86: got 0 want 1
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We just need to handle the scenario where `Find` returns `nil` because it couldn't find the player.

我们只需要处理 `Find` 返回 `nil` 的场景，因为它找不到玩家。

```go
//file_system_store.go
func (f *FileSystemPlayerStore) RecordWin(name string) {
    league := f.GetLeague()
    player := league.Find(name)

    if player != nil {
        player.Wins++
    } else {
        league = append(league, Player{name, 1})
    }

    f.database.Seek(0, 0)
    json.NewEncoder(f.database).Encode(league)
}
```

The happy path is looking ok so we can now try using our new `Store` in the integration test. This will give us more confidence that the software works and then we can delete the redundant `InMemoryPlayerStore`.

快乐路径看起来没问题，所以我们现在可以尝试在集成测试中使用我们的新 `Store`。这将使我们更有信心该软件可以工作，然后我们可以删除多余的“InMemoryPlayerStore”。

In `TestRecordingWinsAndRetrievingThem` replace the old store.

在`TestRecordingWinsAndRetrievingThem` 中替换旧存储。

```go
//server_integration_test.go
database, cleanDatabase := createTempFile(t, "")
defer cleanDatabase()
store := &FileSystemPlayerStore{database}
```

If you run the test it should pass and now we can delete `InMemoryPlayerStore`. `main.go` will now have compilation problems which will motivate us to now use our new store in the "real" code.

如果你运行测试它应该通过，现在我们可以删除`InMemoryPlayerStore`。 `main.go` 现在会有编译问题，这将促使我们现在在“真实”代码中使用我们的新存储。

```go
//main.go
package main

import (
    "log"
    "net/http"
    "os"
)

const dbFileName = "game.db.json"

func main() {
    db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

    if err != nil {
        log.Fatalf("problem opening %s %v", dbFileName, err)
    }

    store := &FileSystemPlayerStore{db}
    server := NewPlayerServer(store)

    if err := http.ListenAndServe(":5000", server);err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}
```

- We create a file for our database.
- The 2nd argument to `os.OpenFile` lets you define the permissions for opening the file, in our case `O_RDWR` means we want to read and write _and_ `os.O_CREATE` means create the file if it doesn't exist. 

- 我们为我们的数据库创建一个文件。
- `os.OpenFile` 的第二个参数让你定义打开文件的权限，在我们的例子中，`O_RDWR` 意味着我们想要读取和写入 _and_ `os.O_CREATE` 意味着如果文件不存在则创建文件。

- The 3rd argument means sets permissions for the file, in our case, all users can read and write the file. [(See superuser.com for a more detailed explanation)](https://superuser.com/questions/295591/what-is-the-meaning-of-chmod-666).

- 第三个参数表示为文件设置权限，在我们的例子中，所有用户都可以读取和写入文件。 [（有关更详细的解释，请参阅 superuser.com）](https://superuser.com/questions/295591/what-is-the-meaning-of-chmod-666)。

Running the program now persists the data in a file in between restarts, hooray!

现在运行程序会在两次重启之间将数据保存在一个文件中，万岁！

## More refactoring and performance concerns

## 更多的重构和性能问题

Every time someone calls `GetLeague()` or `GetPlayerScore()` we are reading the entire file and parsing it into JSON. We should not have to do that because `FileSystemStore` is entirely responsible for the state of the league; it should only need to read the file when the program starts up and only need to update the file when data changes.

每次有人调用 `GetLeague()` 或 `GetPlayerScore()` 时，我们都会读取整个文件并将其解析为 JSON。我们不应该这样做，因为`FileSystemStore` 完全负责联盟的状态；它应该只需要在程序启动时读取文件，只有在数据发生变化时才需要更新文件。

We can create a constructor which can do some of this initialisation for us and store the league as a value in our `FileSystemStore` to be used on the reads instead.

我们可以创建一个构造函数，它可以为我们进行一些初始化，并将联盟存储为我们的“FileSystemStore”中的一个值，以便在读取时使用。

```go
//file_system_store.go
type FileSystemPlayerStore struct {
    database io.ReadWriteSeeker
    league League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
    database.Seek(0, 0)
    league, _ := NewLeague(database)
    return &FileSystemPlayerStore{
        database:database,
        league:league,
    }
}
```

This way we only have to read from disk once. We can now replace all of our previous calls to getting the league from disk and just use `f.league` instead.

这样我们只需要从磁盘读取一次。我们现在可以替换之前所有从磁盘获取联赛的调用，而只需使用 `f.league`。

```go
//file_system_store.go
func (f *FileSystemPlayerStore) GetLeague() League {
    return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

    player := f.league.Find(name)

    if player != nil {
        return player.Wins
    }

    return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
    player := f.league.Find(name)

    if player != nil {
        player.Wins++
    } else {
        f.league = append(f.league, Player{name, 1})
    }

    f.database.Seek(0, 0)
    json.NewEncoder(f.database).Encode(f.league)
}
```

If you try to run the tests it will now complain about initialising `FileSystemPlayerStore` so just fix them by calling our new constructor.

如果您尝试运行测试，它现在会抱怨初始化 `FileSystemPlayerStore`，因此只需通过调用我们的新构造函数来修复它们。

### Another problem

###  另一个问题

There is some more naivety in the way we are dealing with files which _could_ create a very nasty bug down the line.

我们处理文件的方式有些幼稚，可能会造成非常讨厌的错误。

When we `RecordWin`, we `Seek` back to the start of the file and then write the new data—but what if the new data was smaller than what was there before?

当我们`RecordWin` 时，我们`Seek` 回到文件的开头，然后写入新数据——但是如果新数据比之前的数据小怎么办？

In our current case, this is impossible. We never edit or delete scores so the data can only get bigger. However, it would be irresponsible for us to leave the code like this; it's not unthinkable that a delete scenario could come up.

在我们目前的情况下，这是不可能的。我们从不编辑或删除分数，因此数据只会变得更大。但是，这样留下代码对我们来说是不负责任的；出现删除场景并非不可想象。

How will we test for this though? What we need to do is first refactor our code so we separate out the concern of the _kind of data we write, from the writing_. We can then test that separately to check it works how we hope.

不过，我们将如何对此进行测试？我们需要做的是首先重构我们的代码，这样我们就可以将我们编写的_种类数据的关注点与编写_分开。然后我们可以单独测试它以检查它是否按我们希望的方式工作。

We'll create a new type to encapsulate our "when we write we go from the beginning" functionality. I'm going to call it `Tape`. Create a new file with the following:

我们将创建一个新类型来封装我们的“编写时我们从头开始”功能。我将称它为“磁带”。使用以下内容创建一个新文件：

```go
//tape.go
package main

import "io"

type tape struct {
    file io.ReadWriteSeeker
}

func (t *tape) Write(p []byte) (n int, err error) {
    t.file.Seek(0, 0)
    return t.file.Write(p)
}
```

Notice that we're only implementing `Write` now, as it encapsulates the `Seek` part. This means our `FileSystemStore` can just have a reference to a `Writer` instead.

请注意，我们现在只实现了 `Write`，因为它封装了 `Seek` 部分。这意味着我们的 `FileSystemStore` 可以只引用一个 `Writer`。

```go
//file_system_store.go
type FileSystemPlayerStore struct {
    database io.Writer
    league   League
}
```

Update the constructor to use `Tape`

更新构造函数以使用`Tape`

```go
//file_system_store.go
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
    database.Seek(0, 0)
    league, _ := NewLeague(database)

    return &FileSystemPlayerStore{
        database: &tape{database},
        league:   league,
    }
}
```

Finally, we can get the amazing payoff we wanted by removing the `Seek` call from `RecordWin`. Yes, it doesn't feel much, but at least it means if we do any other kind of writes we can rely on our `Write` to behave how we need it to. Plus it will now let us test the potentially problematic code separately and fix it.

最后，我们可以通过从 `RecordWin` 中删除 `Seek` 调用来获得我们想要的惊人回报。是的，感觉并不多，但至少这意味着如果我们进行任何其他类型的写入，我们可以依靠我们的“Write”来按照我们需要的方式行事。此外，它现在可以让我们单独测试可能有问题的代码并修复它。

Let's write the test where we want to update the entire contents of a file with something that is smaller than the original contents.

让我们编写测试，我们希望用比原始内容小的内容来更新文件的全部内容。

## Write the test first

## 先写测试

Our test will create a file with some content, try to write to it using the `tape`, and read it all again to see what's in the file. In `tape_test.go`:

我们的测试将创建一个包含一些内容的文件，尝试使用 `tape` 写入它，然后再次读取它以查看文件中的内容。在`tape_test.go`中：

```go
//tape_test.go
func TestTape_Write(t *testing.T) {
    file, clean := createTempFile(t, "12345")
    defer clean()

    tape := &tape{file}

    tape.Write([]byte("abc"))

    file.Seek(0, 0)
    newFileContents, _ := ioutil.ReadAll(file)

    got := string(newFileContents)
    want := "abc"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestTape_Write
--- FAIL: TestTape_Write (0.00s)
    tape_test.go:23: got 'abc45' want 'abc'
```

As we thought! It writes the data we want, but leaves the rest of the original data remaining.

正如我们所想的！它写入我们想要的数据，但保留其余的原始数据。

## Write enough code to make it pass

## 编写足够的代码使其通过

`os.File` has a truncate function that will let us effectively empty the file. We should be able to just call this to get what we want.

`os.File` 有一个 truncate 函数，可以让我们有效地清空文件。我们应该能够调用它来获得我们想要的东西。

Change `tape` to the following:

将 `tape` 更改为以下内容：

```go
//tape.go
type tape struct {
    file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
    t.file.Truncate(0)
    t.file.Seek(0, 0)
    return t.file.Write(p)
}
```

The compiler will fail in a number of places where we are expecting an `io.ReadWriteSeeker` but we are sending in `*os.File`. You should be able to fix these problems yourself by now but if you get stuck just check the source code.

编译器会在很多地方失败，我们期待一个 `io.ReadWriteSeeker` 但我们发送的是 `*os.File`。您现在应该能够自己解决这些问题，但如果您遇到困难，请查看源代码。

Once you get it refactoring our `TestTape_Write` test should be passing!

一旦你得到它重构我们的`TestTape_Write`测试应该通过！

### One other small refactor

### 另一个小重构

In `RecordWin` we have the line `json.NewEncoder(f.database).Encode(f.league)`.

在 `RecordWin` 中有一行 `json.NewEncoder(f.database).Encode(f.league)`。

We don't need to create a new encoder every time we write, we can initialise one in our constructor and use that instead.

我们不需要在每次编写时都创建一个新的编码器，我们可以在我们的构造函数中初始化一个并使用它。

Store a reference to an `Encoder` in our type and initialise it in the constructor:

在我们的类型中存储对“编码器”的引用，并在构造函数中初始化它：

```go
//file_system_store.go
type FileSystemPlayerStore struct {
    database *json.Encoder
    league   League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
    file.Seek(0, 0)
    league, _ := NewLeague(file)

    return &FileSystemPlayerStore{
        database: json.NewEncoder(&tape{file}),
        league:   league,
    }
}
```

Use it in `RecordWin`.

在“RecordWin”中使用它。

```go
func (f *FileSystemPlayerStore) RecordWin(name string) {
    player := f.league.Find(name)

    if player != nil {
        player.Wins++
    } else {
        f.league = append(f.league, Player{name, 1})
    }

    f.database.Encode(f.league)
}
```

## Didn't we just break some rules there? Testing private things? No interfaces?

## 我们不是在那里违反了一些规则吗？测试私事？没有接口？

### On testing private types

### 关于测试私有类型

It's true that _in general_ you should favour not testing private things as that can sometimes lead to your tests being too tightly coupled to the implementation, which can hinder refactoring in future.

确实_一般来说_您应该倾向于不测试私有事物，因为这有时会导致您的测试与实现过于紧密地耦合，这可能会阻碍未来的重构。

However, we must not forget that tests should give us _confidence_.

然而，我们不能忘记测试应该给我们_信心_。

We were not confident that our implementation would work if we added any kind of edit or delete functionality. We did not want to leave the code like that, especially if this was being worked on by more than one person who may not be aware of the shortcomings of our initial approach.

如果我们添加任何类型的编辑或删除功能，我们不相信我们的实现会起作用。我们不想留下这样的代码，特别是如果有不止一个人在处理这个问题，他们可能不知道我们最初的方法的缺点。

Finally, it's just one test! If we decide to change the way it works it won't be a disaster to just delete the test but we have at the very least captured the requirement for future maintainers.

最后，这只是一项测试！如果我们决定改变它的工作方式，删除测试不会是一场灾难，但我们至少已经抓住了未来维护者的需求。

### Interfaces

### 接口

We started off the code by using `io.Reader` as that was the easiest path for us to unit test our new `PlayerStore`. As we developed the code we moved on to `io.ReadWriter` and then `io.ReadWriteSeeker`. We then found out there was nothing in the standard library that actually implemented that apart from `*os.File`. We could've taken the decision to write our own or use an open source one but it felt pragmatic just to make temporary files for the tests.

我们使用 `io.Reader` 开始编写代码，因为这是我们对新的 `PlayerStore` 进行单元测试的最简单途径。当我们开发代码时，我们继续使用 `io.ReadWriter` 和 `io.ReadWriteSeeker`。然后我们发现标准库中除了 `*os.File` 之外没有任何东西真正实现了它。我们本可以决定自己编写或使用开源的，但仅仅为测试制作临时文件感觉很实用。

Finally, we needed `Truncate` which is also on `*os.File`. It would've been an option to create our own interface capturing these requirements.

最后，我们需要 `Truncate`，它也在 `*os.File` 上。可以选择创建我们自己的界面来捕捉这些需求。

```go
type ReadWriteSeekTruncate interface {
    io.ReadWriteSeeker
    Truncate(size int64) error
}
```

But what is this really giving us? Bear in mind we are _not mocking_ and it is unrealistic for a **file system** store to take any type other than an `*os.File` so we don't need the polymorphism that interfaces give us.

但这真的给了我们什么？请记住，我们 _not mocking_ 并且 **文件系统** 存储采用除 `*os.File` 之外的任何类型是不现实的，因此我们不需要接口提供给我们的多态性。

Don't be afraid to chop and change types and experiment like we have here. The great thing about using a statically typed language is the compiler will help you with every change.

不要害怕像我们这里那样切碎和改变类型并进行实验。使用静态类型语言的好处是编译器会帮助您进行每一次更改。

## Error handling

## 错误处理

Before we start working on sorting we should make sure we're happy with our current code and remove any technical debt we may have. It's an important principle to get to working software as quickly as possible (stay out of the red state) but that doesn't mean we should ignore error cases!

在我们开始排序之前，我们应该确保我们对当前的代码感到满意并消除我们可能拥有的任何技术债务。尽快进入工作软件是一个重要的原则（远离红色状态），但这并不意味着我们应该忽略错误情况！

If we go back to `FileSystemStore.go` we have `league, _ := NewLeague(f.database)` in our constructor.

如果我们回到 `FileSystemStore.go`，我们的构造函数中有 `league, _ := NewLeague(f.database)`。

`NewLeague` can return an error if it is unable to parse the league from the `io.Reader` that we provide. 

如果`NewLeague` 无法从我们提供的`io.Reader` 解析联赛，则会返回错误。

It was pragmatic to ignore that at the time as we already had failing tests. If we had tried to tackle it at the same time, we would have been juggling two things at once.

当时忽略这一点是务实的，因为我们已经有失败的测试。如果我们试图同时解决它，我们就会同时处理两件事。

Let's make it so our constructor is capable of returning an error.

让我们让我们的构造函数能够返回错误。

```go
//file_system_store.go
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
    file.Seek(0, 0)
    league, err := NewLeague(file)

    if err != nil {
        return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
    }

    return &FileSystemPlayerStore{
        database: json.NewEncoder(&tape{file}),
        league:   league,
    }, nil
}
```

Remember it is very important to give helpful error messages (just like your tests). People on the internet jokingly say that most Go code is:

请记住，提供有用的错误消息非常重要（就像您的测试一样）。网上有人开玩笑说，大部分 Go 代码是：

```go
if err != nil {
    return err
}
```

**That is 100% not idiomatic.** Adding contextual information (i.e what you were doing to cause the error) to your error messages makes operating your software far easier.

**这 100% 不是惯用的。** 在错误消息中添加上下文信息（即您正在做什么导致错误）使操作您的软件变得更加容易。

If you try to compile you'll get some errors.

如果你尝试编译你会得到一些错误。

```
./main.go:18:35: multiple-value NewFileSystemPlayerStore() in single-value context
./file_system_store_test.go:35:36: multiple-value NewFileSystemPlayerStore() in single-value context
./file_system_store_test.go:57:36: multiple-value NewFileSystemPlayerStore() in single-value context
./file_system_store_test.go:70:36: multiple-value NewFileSystemPlayerStore() in single-value context
./file_system_store_test.go:85:36: multiple-value NewFileSystemPlayerStore() in single-value context
./server_integration_test.go:12:35: multiple-value NewFileSystemPlayerStore() in single-value context
```

In main we'll want to exit the program, printing the error.

在 main 中，我们要退出程序，打印错误。

```go
//main.go
store, err := NewFileSystemPlayerStore(db)

if err != nil {
    log.Fatalf("problem creating file system player store, %v ", err)
}
```

In the tests we should assert there is no error. We can make a helper to help with this.

在测试中，我们应该断言没有错误。我们可以做一个帮手来帮助解决这个问题。

```go
//file_system_store_test.go
func assertNoError(t testing.TB, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("didn't expect an error but got one, %v", err)
    }
}
```

Work through the other compilation problems using this helper. Finally, you should have a failing test:

使用此帮助程序解决其他编译问题。最后，你应该有一个失败的测试：

```
=== RUN   TestRecordingWinsAndRetrievingThem
--- FAIL: TestRecordingWinsAndRetrievingThem (0.00s)
    server_integration_test.go:14: didn't expect an error but got one, problem loading player store from file /var/folders/nj/r_ccbj5d7flds0sf63yy4vb80000gn/T/db841037437, problem parsing league, EOF
```

We cannot parse the league because the file is empty. We weren't getting errors before because we always just ignored them.

我们无法解析联赛，因为文件是空的。我们之前没有收到错误，因为我们总是忽略它们。

Let's fix our big integration test by putting some valid JSON in it:

让我们通过在其中放入一些有效的 JSON 来修复我们的大型集成测试：

```go
//server_integration_test.go
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
    database, cleanDatabase := createTempFile(t, `[]`)
    //etc...
```

Now that all the tests are passing, we need to handle the scenario where the file is empty.

现在所有的测试都通过了，我们需要处理文件为空的场景。

## Write the test first

## 先写测试

```go
//file_system_store_test.go
t.Run("works with an empty file", func(t *testing.T) {
    database, cleanDatabase := createTempFile(t, "")
    defer cleanDatabase()

    _, err := NewFileSystemPlayerStore(database)

    assertNoError(t, err)
})
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestFileSystemStore/works_with_an_empty_file
    --- FAIL: TestFileSystemStore/works_with_an_empty_file (0.00s)
        file_system_store_test.go:108: didn't expect an error but got one, problem loading player store from file /var/folders/nj/r_ccbj5d7flds0sf63yy4vb80000gn/T/db019548018, problem parsing league, EOF
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Change our constructor to the following

将我们的构造函数更改为以下内容

```go
//file_system_store.go
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

    file.Seek(0, 0)

    info, err := file.Stat()

    if err != nil {
        return nil, fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
    }

    if info.Size() == 0 {
        file.Write([]byte("[]"))
        file.Seek(0, 0)
    }

    league, err := NewLeague(file)

    if err != nil {
        return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
    }

    return &FileSystemPlayerStore{
        database: json.NewEncoder(&tape{file}),
        league:   league,
    }, nil
}
```

`file.Stat` returns stats on our file, which lets us check the size of the file. If it's empty, we `Write` an empty JSON array and `Seek` back to the start, ready for the rest of the code.

`file.Stat` 返回文件的统计信息，让我们可以检查文件的大小。如果它是空的，我们`Write`一个空的JSON数组，然后`Seek`回到开始，为其余的代码做好准备。

## Refactor

## 重构

Our constructor is a bit messy now, so let's extract the initialise code into a function:

我们的构造函数现在有点乱，所以让我们将初始化代码提取到一个函数中：

```go
//file_system_store.go
func initialisePlayerDBFile(file *os.File) error {
    file.Seek(0, 0)

    info, err := file.Stat()

    if err != nil {
        return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
    }

    if info.Size()==0 {
        file.Write([]byte("[]"))
        file.Seek(0, 0)
    }

    return nil
}
```

```go
//file_system_store.go
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

    err := initialisePlayerDBFile(file)

    if err != nil {
        return nil, fmt.Errorf("problem initialising player db file, %v", err)
    }

    league, err := NewLeague(file)

    if err != nil {
        return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
    }

    return &FileSystemPlayerStore{
        database: json.NewEncoder(&tape{file}),
        league:   league,
    }, nil
}
```

## Sorting

## 排序

Our product owner wants `/league` to return the players sorted by their scores, from highest to lowest.

我们的产品负责人希望 `/league` 返回按得分从高到低排序的球员。

The main decision to make here is where in the software should this happen. If we were using a "real" database we would use things like `ORDER BY` so the sorting is super fast. For that reason, it feels like implementations of `PlayerStore` should be responsible.

在这里做出的主要决定是在软件中的哪个位置发生这种情况。如果我们使用“真正的”数据库，我们会使用诸如“ORDER BY”之类的东西，因此排序非常快。出于这个原因，感觉应该由`PlayerStore` 的实现负责。

## Write the test first

## 先写测试

We can update the assertion on our first test in `TestFileSystemStore`:

我们可以在“TestFileSystemStore”中更新我们第一个测试的断言：

```go
//file_system_store_test.go
t.Run("league sorted", func(t *testing.T) {
    database, cleanDatabase := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
    defer cleanDatabase()

    store, err := NewFileSystemPlayerStore(database)

    assertNoError(t, err)

    got := store.GetLeague()

    want := []Player{
        {"Chris", 33},
        {"Cleo", 10},
    }

    assertLeague(t, got, want)

    // read again
    got = store.GetLeague()
    assertLeague(t, got, want)
})
```

The order of the JSON coming in is in the wrong order and our `want` will check that it is returned to the caller in the correct order.

传入的 JSON 的顺序是错误的，我们的 `want` 将检查它是否以正确的顺序返回给调用者。

## Try to run the test

## 尝试运行测试

```
=== RUN   TestFileSystemStore/league_from_a_reader,_sorted
    --- FAIL: TestFileSystemStore/league_from_a_reader,_sorted (0.00s)
        file_system_store_test.go:46: got [{Cleo 10} {Chris 33}] want [{Chris 33} {Cleo 10}]
        file_system_store_test.go:51: got [{Cleo 10} {Chris 33}] want [{Chris 33} {Cleo 10}]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (f *FileSystemPlayerStore) GetLeague() League {
    sort.Slice(f.league, func(i, j int) bool {
        return f.league[i].Wins > f.league[j].Wins
    })
    return f.league
}
```

[`sort.Slice`](https://golang.org/pkg/sort/#Slice)

[`sort.Slice`](https://golang.org/pkg/sort/#Slice)

> Slice sorts the provided slice given the provided less function.

> Slice 在给定提供的 less 函数的情况下对提供的切片进行排序。

Easy!

简单！

## Wrapping up

##  总结

### What we've covered

### 我们已经介绍过的内容

- The `Seeker` interface and its relation to `Reader` and `Writer`.
- Working with files.
- Creating an easy to use helper for testing with files that hides all the messy stuff.
- `sort.Slice` for sorting slices.
- Using the compiler to help us safely make structural changes to the application.

- `Seeker` 接口及其与 `Reader` 和 `Writer` 的关系。
- 处理文件。
- 创建一个易于使用的帮助程序，用于测试隐藏所有杂乱内容的文件。
- `sort.Slice` 用于对切片进行排序。
- 使用编译器帮助我们安全地对应用程序进行结构更改。

### Breaking rules

### 打破规则

- Most rules in software engineering aren't really rules, just best practices that work 80% of the time.
- We discovered a scenario where one of our previous "rules" of not testing internal functions was not helpful for us so we broke the rule.
- It's important when breaking rules to understand the trade-off you are making. In our case, we were ok with it because it was just one test and would've been very difficult to exercise the scenario otherwise.
- In order to be able to break the rules **you must understand them first**. An analogy is with learning guitar. It doesn't matter how creative you think you are, you must understand and practice the fundamentals.

- 软件工程中的大多数规则并不是真正的规则，只是在 80% 的时间里都有效的最佳实践。
- 我们发现了一个场景，我们之前不测试内部功能的“规则”之一对我们没有帮助，所以我们打破了规则。
- 在违反规则时了解您正在做出的权衡很重要。在我们的案例中，我们对此没有意见，因为这只是一项测试，否则很难执行该场景。
- 为了能够打破规则**你必须先了解它们**。一个类比是学习吉他。无论您认为自己有多有创造力，您都必须了解并实践基本原理。

### Where our software is at

### 我们的软件在哪里

- We have an HTTP API where you can create players and increment their score.
- We can return a league of everyone's scores as JSON.
- The data is persisted as a JSON file. 

- 我们有一个 HTTP API，您可以在其中创建玩家并增加他们的分数。
- 我们可以将每个人的分数作为 JSON 返回。
- 数据保存为 JSON 文件。

