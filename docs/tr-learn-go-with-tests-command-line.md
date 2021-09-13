# Command line and project structure

# 命令行和项目结构

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/command-line)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/command-line)**

Our product owner now wants to _pivot_ by introducing a second application - a command line application.

我们的产品负责人现在希望通过引入第二个应用程序 - 命令行应用程序来_旋转_。

For now, it will just need to be able to record a player's win when the user types `Ruth wins`. The intention is to eventually be a tool for helping users play poker.

现在，它只需要能够在用户输入“Ruth wins”时记录玩家的胜利。目的是最终成为帮助用户玩扑克的工具。

The product owner wants the database to be shared amongst the two applications so that the league updates according to wins recorded in the new application.

产品所有者希望在两个应用程序之间共享数据库，以便联盟根据新应用程序中记录的胜利进行更新。

## A reminder of the code

## 代码提醒

We have an application with a `main.go` file that launches an HTTP server. The HTTP server won't be interesting to us for this exercise but the abstraction it uses will. It depends on a `PlayerStore`.

我们有一个带有 `main.go` 文件的应用程序，用于启动 HTTP 服务器。对于这个练习，我们不会对 HTTP 服务器感兴趣，但它使用的抽象会很有趣。这取决于“PlayerStore”。

```go
type PlayerStore interface {
    GetPlayerScore(name string) int
    RecordWin(name string)
    GetLeague() League
}
```

In the previous chapter, we made a `FileSystemPlayerStore` which implements that interface. We should be able to re-use some of this for our new application.

在上一章中，我们创建了一个实现该接口的 `FileSystemPlayerStore`。我们应该能够在我们的新应用程序中重用其中的一些。

## Some project refactoring first

## 先重构一些项目

Our project now needs to create two binaries, our existing web server and the command line app.

我们的项目现在需要创建两个二进制文件，我们现有的 Web 服务器和命令行应用程序。

Before we get stuck into our new work we should structure our project to accommodate this.

在我们陷入新工作之前，我们应该构建我们的项目以适应这一点。

So far all the code has lived in one folder, in a path looking like this

到目前为止，所有代码都保存在一个文件夹中，路径如下所示

`$GOPATH/src/github.com/your-name/my-app`

In order for you to make an application in Go, you need a `main` function inside a `package main`. So far all of our "domain" code has lived inside `package main` and our `func main` can reference everything.

为了让你在 Go 中创建一个应用程序，你需要在 `package main` 中有一个 `main` 函数。到目前为止，我们所有的“域”代码都存在于 `package main` 中，我们的 `func main` 可以引用所有内容。

This was fine so far and it is good practice not to go over-the-top with package structure. If you take the time to look through the standard library you will see very little in the way of lots of folders and structure.

到目前为止，这很好，最好不要过度使用封装结构。如果您花时间浏览标准库，您将看不到大量文件夹和结构。

Thankfully it's pretty straightforward to add structure _when you need it_.

值得庆幸的是，_在需要时_添加结构非常简单。

Inside the existing project create a `cmd` directory with a `webserver` directory inside that (e.g `mkdir -p cmd/webserver`).

在现有项目中创建一个 `cmd` 目录，其中包含一个 `webserver` 目录（例如 `mkdir -p cmd/webserver`）。

Move the `main.go` inside there.

将 `main.go` 移到里面。

If you have `tree` installed you should run it and your structure should look like this

如果你安装了`tree`，你应该运行它，你的结构应该是这样的

```
.
├── file_system_store.go
├── file_system_store_test.go
├── cmd
│   └── webserver
│       └── main.go
├── league.go
├── server.go
├── server_integration_test.go
├── server_test.go
├── tape.go
└── tape_test.go
```

We now effectively have a separation between our application and the library code but we now need to change some package names. Remember when you build a Go application its package _must_ be `main`.

我们现在有效地将我们的应用程序和库代码分开，但我们现在需要更改一些包名称。请记住，当您构建 Go 应用程序时，它的包 _must_ 是 `main`。

Change all the other code to have a package called `poker`.

将所有其他代码更改为具有名为“poker”的包。

Finally, we need to import this package into `main.go` so we can use it to create our web server. Then we can use our library code by using `poker.FunctionName`.

最后，我们需要将这个包导入到 `main.go` 中，以便我们可以使用它来创建我们的 Web 服务器。然后我们可以通过使用`poker.FunctionName` 来使用我们的库代码。

The paths will be different on your computer, but it should be similar to this:

您的计算机上的路径会有所不同，但应该类似于：

```go
//cmd/webserver/main.go
package main

import (
    "github.com/quii/learn-go-with-tests/command-line/v1"
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

    store, err := poker.NewFileSystemPlayerStore(db)

    if err != nil {
        log.Fatalf("problem creating file system player store, %v ", err)
    }

    server := poker.NewPlayerServer(store)

    log.Fatal(http.ListenAndServe(":5000", server))
}
```

The full path may seem a bit jarring, but this is how you can import _any_ publicly available library into your code.

完整路径可能看起来有点刺耳，但这就是您可以将 _any_ 公开可用的库导入到您的代码中的方法。

By separating our domain code into a separate package and committing it to a public repo like GitHub any Go developer can write their own code which imports that package the features we've written available. The first time you try and run it will complain it is not existing but all you need to do is run `go get`.

通过将我们的域代码分成一个单独的包并将其提交到像 GitHub 这样的公共存储库，任何 Go 开发人员都可以编写自己的代码来导入我们编写的可用功能。第一次尝试运行它时会抱怨它不存在，但您需要做的就是运行 `go get`。

In addition, users can view [the documentation at godoc.org](https://godoc.org/github.com/quii/learn-go-with-tests/command-line/v1).

此外，用户可以查看[godoc.org 上的文档](https://godoc.org/github.com/quii/learn-go-with-tests/command-line/v1)。

### Final checks

### 最终检查

- Inside the root run `go test` and check they're still passing
- Go inside our `cmd/webserver` and do `go run main.go`
   - Visit `http://localhost:5000/league` and you should see it's still working

- 在 root 中运行 `go test` 并检查它们是否仍然通过
- 进入我们的 `cmd/webserver` 并执行 `go run main.go`
  - 访问 `http://localhost:5000/league`，你应该会看到它仍在工作

### Walking skeleton 

### 行走的骨架

Before we get stuck into writing tests, let's add a new application that our project will build. Create another directory inside `cmd` called `cli` (command line interface) and add a `main.go` with the following

在我们陷入编写测试之前，让我们添加一个我们的项目将构建的新应用程序。在 `cmd` 中创建另一个名为 `cli`（命令行界面）的目录，并添加一个 `main.go`，其内容如下

```go
//cmd/cli/main.go
package main

import "fmt"

func main() {
    fmt.Println("Let's play poker")
}
```

The first requirement we'll tackle is recording a win when the user types `{PlayerName} wins`.

我们要解决的第一个要求是在用户输入“{PlayerName} wins”时记录一次胜利。

## Write the test first

## 先写测试

We know we need to make something called `CLI` which will allow us to `Play` poker. It'll need to read user input and then record wins to a `PlayerStore`.

我们知道我们需要制作一种叫做“CLI”的东西，它可以让我们“玩”扑克。它需要读取用户输入，然后将胜利记录到“PlayerStore”。

Before we jump too far ahead though, let's just write a test to check it integrates with the `PlayerStore` how we'd like.

不过，在我们跳得太远之前，让我们编写一个测试来检查它是否与我们希望的“PlayerStore”集成。

Inside `CLI_test.go` (in the root of the project, not inside `cmd`)

在 `CLI_test.go` 中（在项目的根目录中，而不是在 `cmd` 中）

```go
//CLI_test.go
package poker

import "testing"

func TestCLI(t *testing.T) {
    playerStore := &StubPlayerStore{}
    cli := &CLI{playerStore}
    cli.PlayPoker()

    if len(playerStore.winCalls) != 1 {
        t.Fatal("expected a win call but didn't get any")
    }
}
```

- We can use our `StubPlayerStore` from other tests
- We pass in our dependency into our not yet existing `CLI` type
- Trigger the game by an unwritten `PlayPoker` method
- Check that a win is recorded

- 我们可以使用其他测试中的 `StubPlayerStore`
- 我们将我们的依赖传递到我们尚不存在的 `CLI` 类型中
- 通过不成文的`PlayPoker` 方法触发游戏
- 检查是否记录了胜利

## Try to run the test

## 尝试运行测试

```
# github.com/quii/learn-go-with-tests/command-line/v2
./cli_test.go:25:10: undefined: CLI
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

At this point, you should be comfortable enough to create our new `CLI` struct with the respective field for our dependency and add a method.

在这一点上，您应该足够舒服地使用我们的依赖项的相应字段创建我们新的 CLI 结构并添加一个方法。

You should end up with code like this

你应该得到这样的代码

```go
//CLI.go
package poker

type CLI struct {
    playerStore PlayerStore
}

func (cli *CLI) PlayPoker() {}
```

Remember we're just trying to get the test running so we can check the test fails how we'd hope

请记住，我们只是想让测试运行，这样我们就可以检查测试是否如我们希望的那样失败

```
--- FAIL: TestCLI (0.00s)
    cli_test.go:30: expected a win call but didn't get any
FAIL
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
//CLI.go
func (cli *CLI) PlayPoker() {
    cli.playerStore.RecordWin("Cleo")
}
```

That should make it pass.

那应该让它通过。

Next, we need to simulate reading from `Stdin` (the input from the user) so that we can record wins for specific players.

接下来，我们需要模拟从 `Stdin`（来自用户的输入）读取，以便我们可以记录特定玩家的胜利。

Let's extend our test to exercise this.

让我们扩展我们的测试来练习这个。

## Write the test first

## 先写测试

```go
//CLI_test.go
func TestCLI(t *testing.T) {
    in := strings.NewReader("Chris wins\n")
    playerStore := &StubPlayerStore{}

    cli := &CLI{playerStore, in}
    cli.PlayPoker()

    if len(playerStore.winCalls) != 1 {
        t.Fatal("expected a win call but didn't get any")
    }

    got := playerStore.winCalls[0]
    want := "Chris"

    if got != want {
        t.Errorf("didn't record correct winner, got %q, want %q", got, want)
    }
}
```

`os.Stdin` is what we'll use in `main` to capture the user's input. It is a `*File` under the hood which means it implements `io.Reader` which as we know by now is a handy way of capturing text.

`os.Stdin` 是我们将在 `main` 中用来捕获用户输入的内容。它在底层是一个 `*File`，这意味着它实现了 `io.Reader`，正如我们现在所知，这是一种方便的文本捕获方式。

We create an `io.Reader` in our test using the handy `strings.NewReader`, filling it with what we expect the user to type.

我们使用方便的“strings.NewReader”在我们的测试中创建了一个“io.Reader”，用我们期望用户输入的内容填充它。

## Try to run the test

## 尝试运行测试

`./CLI_test.go:12:32: too many values in struct initializer`

`./CLI_test.go:12:32: 结构初始值设定项中的值太多`

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We need to add our new dependency into `CLI`.

我们需要将我们的新依赖添加到 `CLI` 中。

```go
//CLI.go
type CLI struct {
    playerStore PlayerStore
    in          io.Reader
}
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```
--- FAIL: TestCLI (0.00s)
    CLI_test.go:23: didn't record the correct winner, got 'Cleo', want 'Chris'
FAIL
```

Remember to do the strictly easiest thing first

记住先做最简单的事情

```go
func (cli *CLI) PlayPoker() {
    cli.playerStore.RecordWin("Chris")
}
```

The test passes. We'll add another test to force us to write some real code next, but first, let's refactor.

测试通过。我们将添加另一个测试来迫使我们接下来编写一些真正的代码，但首先，让我们重构。

## Refactor

## 重构

In `server_test` we earlier did checks to see if wins are recorded as we have here. Let's DRY that assertion up into a helper

在 `server_test` 中，我们之前进行了检查以查看是否像我们这里那样记录了胜利。让我们把这个断言干成一个助手

```go
//server_test.go
func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
    t.Helper()

    if len(store.winCalls) != 1 {
        t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
    }

    if store.winCalls[0] != winner {
        t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
    }
}
```

Now replace the assertions in both `server_test.go` and `CLI_test.go`.

现在替换 `server_test.go` 和 `CLI_test.go` 中的断言。

The test should now read like so

测试现在应该是这样的

```go
//CLI_test.go
func TestCLI(t *testing.T) {
    in := strings.NewReader("Chris wins\n")
    playerStore := &StubPlayerStore{}

    cli := &CLI{playerStore, in}
    cli.PlayPoker()

    assertPlayerWin(t, playerStore, "Chris")
}
```

Now let's write _another_ test with different user input to force us into actually reading it.

现在让我们用不同的用户输入编写 _another_ 测试，以迫使我们实际阅读它。

## Write the test first

## 先写测试

```go
//CLI_test.go
func TestCLI(t *testing.T) {

    t.Run("record chris win from user input", func(t *testing.T) {
        in := strings.NewReader("Chris wins\n")
        playerStore := &StubPlayerStore{}

        cli := &CLI{playerStore, in}
        cli.PlayPoker()

        assertPlayerWin(t, playerStore, "Chris")
    })

    t.Run("record cleo win from user input", func(t *testing.T) {
        in := strings.NewReader("Cleo wins\n")
        playerStore := &StubPlayerStore{}

        cli := &CLI{playerStore, in}
        cli.PlayPoker()

        assertPlayerWin(t, playerStore, "Cleo")
    })

}
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestCLI
--- FAIL: TestCLI (0.00s)
=== RUN   TestCLI/record_chris_win_from_user_input
    --- PASS: TestCLI/record_chris_win_from_user_input (0.00s)
=== RUN   TestCLI/record_cleo_win_from_user_input
    --- FAIL: TestCLI/record_cleo_win_from_user_input (0.00s)
        CLI_test.go:27: did not store correct winner got 'Chris' want 'Cleo'
FAIL
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We'll use a [`bufio.Scanner`](https://golang.org/pkg/bufio/) to read the input from the `io.Reader`.

我们将使用 [`bufio.Scanner`](https://golang.org/pkg/bufio/) 从 `io.Reader` 读取输入。

> Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

> 包 bufio 实现缓冲 I/O。它包装了一个 io.Reader 或 io.Writer 对象，创建另一个对象（Reader 或 Writer），该对象也实现了该接口，但为文本 I/O 提供了缓冲和一些帮助。

Update the code to the following

将代码更新为以下内容

```go
//CLI.go
type CLI struct {
    playerStore PlayerStore
    in          io.Reader
}

func (cli *CLI) PlayPoker() {
    reader := bufio.NewScanner(cli.in)
    reader.Scan()
    cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
    return strings.Replace(userInput, " wins", "", 1)
}
```

The tests will now pass.

现在测试将通过。

- `Scanner.Scan()` will read up to a newline.
- We then use `Scanner.Text()` to return the `string` the scanner read to.

- `Scanner.Scan()` 将读取到换行符。
- 然后我们使用 `Scanner.Text()` 返回扫描仪读取到的 `string`。

Now that we have some passing tests, we should wire this up into `main`. Remember we should always strive to have fully-integrated working software as quickly as we can.

现在我们有一些通过的测试，我们应该把它连接到 `main` 中。请记住，我们应该始终努力尽快拥有完全集成的工作软件。

In `main.go` add the following and run it. (you may have to adjust the path of the second dependency to match what's on your computer)

在 main.go 中添加以下内容并运行它。 （您可能需要调整第二个依赖项的路径以匹配您计算机上的内容）

```go
package main

import (
    "fmt"
    "github.com/quii/learn-go-with-tests/command-line/v3"
    "log"
    "os"
)

const dbFileName = "game.db.json"

func main() {
    fmt.Println("Let's play poker")
    fmt.Println("Type {Name} wins to record a win")

    db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

    if err != nil {
        log.Fatalf("problem opening %s %v", dbFileName, err)
    }

    store, err := poker.NewFileSystemPlayerStore(db)

    if err != nil {
        log.Fatalf("problem creating file system player store, %v ", err)
    }

    game := poker.CLI{store, os.Stdin}
    game.PlayPoker()
}
```

You should get an error

你应该得到一个错误

```
command-line/v3/cmd/cli/main.go:32:25: implicit assignment of unexported field 'playerStore' in poker.CLI literal
command-line/v3/cmd/cli/main.go:32:34: implicit assignment of unexported field 'in' in poker.CLI literal
```

What's happening here is because we are trying to assign to the fields `playerStore` and `in` in `CLI`. These are unexported (private) fields. We _could_ do this in our test code because our test is in the same package as `CLI` (`poker`). But our `main` is in package `main` so it does not have access.

这里发生的事情是因为我们试图分配给 `CLI` 中的字段 `playerStore` 和 `in`。这些是未导出的（私有）字段。我们_可以_在我们的测试代码中这样做，因为我们的测试与`CLI`（`poker`）在同一个包中。但是我们的 `main` 在包 `main` 中，所以它没有访问权限。

This highlights the importance of _integrating your work_. We rightfully made the dependencies of our `CLI` private (because we don't want them exposed to users of `CLI`s) but haven't made a way for users to construct it.

这突出了_整合您的工作_的重要性。我们正确地将 `CLI` 的依赖项设为私有（因为我们不希望它们暴露给 `CLI`s 的用户），但还没有为用户提供构建它的方法。

Is there a way to have caught this problem earlier?

有没有办法更早地发现这个问题？

### `package mypackage_test`

In all other examples so far, when we make a test file we declare it as being in the same package that we are testing.

到目前为止，在所有其他示例中，当我们创建一个测试文件时，我们将其声明为与我们正在测试的同一个包中。

This is fine and it means on the odd occasion where we want to test something internal to the package we have access to the unexported types.

这很好，这意味着在奇怪的情况下，我们想要测试包内部的某些内容，我们可以访问未导出的类型。

But given we have advocated for _not_ testing internal things _generally_, can Go help enforce that? What if we could test our code where we only have access to the exported types (like our `main` does)? 

但是鉴于我们提倡_不_测试内部事物_一般_，Go 可以帮助强制执行吗？如果我们可以测试我们只能访问导出类型的代码（就像我们的 `main` 那样）会怎样？

When you're writing a project with multiple packages I would strongly recommend that your test package name has `_test` at the end. When you do this you will only be able to have access to the public types in your package. This would help with this specific case but also helps enforce the discipline of only testing public APIs. If you still wish to test internals you can make a separate test with the package you want to test.

当您编写包含多个包的项目时，我强烈建议您的测试包名称末尾带有 `_test`。执行此操作后，您将只能访问包中的公共类型。这将有助于解决此特定情况，但也有助于强制执行仅测试公共 API 的纪律。如果您仍然希望测试内部结构，您可以对要测试的包进行单独的测试。

An adage with TDD is that if you cannot test your code then it is probably hard for users of your code to integrate with it. Using `package foo_test` will help with this by forcing you to test your code as if you are importing it like users of your package will.

TDD 的格言是，如果您无法测试您的代码，那么您的代码用户可能很难与它集成。使用`package foo_test` 会强制你测试你的代码，就像你像你的包的用户一样导入它一样。

Before fixing `main` let's change the package of our test inside `CLI_test.go` to `poker_test`.

在修复 `main` 之前，让我们将 `CLI_test.go` 中的测试包更改为 `poker_test`。

If you have a well-configured IDE you will suddenly see a lot of red! If you run the compiler you'll get the following errors

如果您有一个配置良好的 IDE，您会突然看到很多红色！如果你运行编译器，你会得到以下错误

```
./CLI_test.go:12:19: undefined: StubPlayerStore
./CLI_test.go:17:3: undefined: assertPlayerWin
./CLI_test.go:22:19: undefined: StubPlayerStore
./CLI_test.go:27:3: undefined: assertPlayerWin
```

We have now stumbled into more questions on package design. In order to test our software we made unexported stubs and helper functions which are no longer available for us to use in our `CLI_test` because the helpers are defined in the `_test.go` files in the `poker` package.

我们现在偶然发现了更多关于包装设计的问题。为了测试我们的软件，我们制作了未导出的存根和辅助函数，它们不再可供我们在 `CLI_test` 中使用，因为辅助函数是在 `poker` 包的 `_test.go` 文件中定义的。

#### Do we want to have our stubs and helpers 'public'?

#### 我们想让我们的存根和助手“公开”吗？

This is a subjective discussion. One could argue that you do not want to pollute your package's API with code to facilitate tests.

这是一个主观的讨论。有人可能会争辩说，您不想用代码污染包的 API 以促进测试。

In the presentation ["Advanced Testing with Go"](https://speakerdeck.com/mitchellh/advanced-testing-with-go?slide=53) by Mitchell Hashimoto, it is described how at HashiCorp they advocate doing this so that users of the package can write tests without having to re-invent the wheel writing stubs. In our case, this would mean anyone using our `poker` package won't have to create their own stub `PlayerStore` if they wish to work with our code.

在 Mitchell Hashimoto 的演讲 ["Advanced Testing with Go"](https://speakerdeck.com/mitchellh/advanced-testing-with-go?slide=53) 中，描述了 HashiCorp 如何提倡这样做，以便包的用户可以编写测试而不必重新发明轮子编写存根。在我们的例子中，这意味着任何使用我们的 `poker` 包的人如果希望使用我们的代码，就不必创建他们自己的存根 `PlayerStore`。

Anecdotally I have used this technique in other shared packages and it has proved extremely useful in terms of users saving time when integrating with our packages.

有趣的是，我在其他共享包中使用了这种技术，事实证明它在用户与我们的包集成时节省时间方面非常有用。

So let's create a file called `testing.go` and add our stub and our helpers.

所以让我们创建一个名为“testing.go”的文件并添加我们的存根和我们的助手。

```go
//testing.go
package poker

import "testing"

type StubPlayerStore struct {
    scores   map[string]int
    winCalls []string
    league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    return score
}

func (s *StubPlayerStore) RecordWin(name string) {
    s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
    return s.league
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
    t.Helper()

    if len(store.winCalls) != 1 {
        t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
    }

    if store.winCalls[0] != winner {
        t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
    }
}

// todo for you - the rest of the helpers
```

You'll need to make the helpers public (remember exporting is done with a capital letter at the start) if you want them to be exposed to importers of our package.

如果您希望它们暴露给我们包的进口商，您需要公开这些助手（请记住，在开始时导出是用大写字母完成的）。

In our `CLI` test you'll need to call the code as if you were using it within a different package.

在我们的 `CLI` 测试中，您需要像在不同的包中使用它一样调用代码。

```go
//CLI_test.go
func TestCLI(t *testing.T) {

    t.Run("record chris win from user input", func(t *testing.T) {
        in := strings.NewReader("Chris wins\n")
        playerStore := &poker.StubPlayerStore{}

        cli := &poker.CLI{playerStore, in}
        cli.PlayPoker()

        poker.AssertPlayerWin(t, playerStore, "Chris")
    })

    t.Run("record cleo win from user input", func(t *testing.T) {
        in := strings.NewReader("Cleo wins\n")
        playerStore := &poker.StubPlayerStore{}

        cli := &poker.CLI{playerStore, in}
        cli.PlayPoker()

        poker.AssertPlayerWin(t, playerStore, "Cleo")
    })

}
```

You'll now see we have the same problems as we had in `main`

您现在会看到我们遇到了与 `main` 中相同的问题

```
./CLI_test.go:15:26: implicit assignment of unexported field 'playerStore' in poker.CLI literal
./CLI_test.go:15:39: implicit assignment of unexported field 'in' in poker.CLI literal
./CLI_test.go:25:26: implicit assignment of unexported field 'playerStore' in poker.CLI literal
./CLI_test.go:25:39: implicit assignment of unexported field 'in' in poker.CLI literal
```

The easiest way to get around this is to make a constructor as we have for other types. We'll also change `CLI` so it stores a `bufio.Scanner` instead of the reader as it's now automatically wrapped at construction time.

解决这个问题的最简单方法是像其他类型一样创建一个构造函数。我们还将更改 `CLI`，以便它存储一个 `bufio.Scanner` 而不是读取器，因为它现在在构建时自动包装。

```go
//CLI.go
type CLI struct {
    playerStore PlayerStore
    in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
    return &CLI{
        playerStore: store,
        in:          bufio.NewScanner(in),
    }
}
```

By doing this, we can then simplify and refactor our reading code

通过这样做，我们可以简化和重构我们的阅读代码

```go
//CLI.go
func (cli *CLI) PlayPoker() {
    userInput := cli.readLine()
    cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
    return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
    cli.in.Scan()
    return cli.in.Text()
}
```

Change the test to use the constructor instead and we should be back to the tests passing.

将测试更改为使用构造函数，我们应该回到通过的测试。

Finally, we can go back to our new `main.go` and use the constructor we just made

最后，我们可以回到我们的新 main.go 并使用我们刚刚创建的构造函数

```go
//cmd/cli/main.go
game := poker.NewCLI(store, os.Stdin)
```

Try and run it, type "Bob wins".

尝试运行它，输入“Bob wins”。

### Refactor

### 重构

We have some repetition in our respective applications where we are opening a file and creating a `file_system_store` from its contents. This feels like a slight weakness in our package's design so we should make a function in it to encapsulate opening a file from a path and returning you the `PlayerStore`.

我们在各自的应用程序中有一些重复，我们打开一个文件并从它的内容创建一个 `file_system_store`。这感觉像是我们包设计中的一个小弱点，因此我们应该在其中创建一个函数来封装从路径打开文件并返回`PlayerStore`。

```go
//file_system_store.go
func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
    db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

    if err != nil {
        return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
    }

    closeFunc := func() {
        db.Close()
    }

    store, err := NewFileSystemPlayerStore(db)

    if err != nil {
        return nil, nil, fmt.Errorf("problem creating file system player store, %v ", err)
    }

    return store, closeFunc, nil
}
```

Now refactor both of our applications to use this function to create the store.

现在重构我们的两个应用程序以使用此函数来创建商店。

#### CLI application code

#### CLI 应用程序代码

```go
//cmd/cli/main.go
package main

import (
    "fmt"
    "github.com/quii/learn-go-with-tests/command-line/v3"
    "log"
    "os"
)

const dbFileName = "game.db.json"

func main() {
    store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

    if err != nil {
        log.Fatal(err)
    }
    defer close()

    fmt.Println("Let's play poker")
    fmt.Println("Type {Name} wins to record a win")
    poker.NewCLI(store, os.Stdin).PlayPoker()
}
```

#### Web server application code

#### Web 服务器应用程序代码

```go
//cmd/webserver/main.go
package main

import (
    "github.com/quii/learn-go-with-tests/command-line/v3"
    "log"
    "net/http"
)

const dbFileName = "game.db.json"

func main() {
    store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

    if err != nil {
        log.Fatal(err)
    }
    defer close()

    server := poker.NewPlayerServer(store)

    if err := http.ListenAndServe(":5000", server);err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}
```

Notice the symmetry: despite being different user interfaces the setup is almost identical. This feels like good validation of our design so far.
And notice also that `FileSystemPlayerStoreFromFile` returns a closing function, so we can close the underlying file once we are done using the Store.

请注意对称性：尽管用户界面不同，但设置几乎相同。到目前为止，这感觉像是对我们设计的良好验证。
还要注意，`FileSystemPlayerStoreFromFile` 返回一个关闭函数，因此我们可以在使用完 Store 后关闭底层文件。

## Wrapping up

##  总结

### Package structure

### 包结构

This chapter meant we wanted to create two applications, re-using the domain code we've written so far. In order to do this, we needed to update our package structure so that we had separate folders for our respective `main`s.

本章意味着我们想要创建两个应用程序，重用我们迄今为止编写的域代码。为了做到这一点，我们需要更新我们的包结构，以便我们为各自的`main`s 拥有单独的文件夹。

By doing this we ran into integration problems due to unexported values so this further demonstrates the value of working in small "slices" and integrating often.

通过这样做，我们遇到了由于未导出的值而导致的集成问题，因此这进一步证明了在小“切片”中工作并经常集成的价值。

We learned how `mypackage_test` helps us create a testing environment which is the same experience for other packages integrating with your code, to help you catch integration problems and see how easy (or not!) your code is to work with.

我们了解了 `mypackage_test` 如何帮助我们创建一个测试环境，该环境与与您的代码集成的其他包具有相同的体验，以帮助您发现集成问题并了解您的代码使用起来有多容易（或不容易！）。

### Reading user input

### 读取用户输入

We saw how reading from `os.Stdin` is very easy for us to work with as it implements `io.Reader`. We used `bufio.Scanner` to easily read line by line user input.

我们看到了从 `os.Stdin` 中读取数据对于我们来说非常容易，因为它实现了 `io.Reader`。我们使用 `bufio.Scanner` 来轻松地逐行读取用户输入。

### Simple abstractions leads to simpler code re-use

### 简单的抽象导致更简单的代码重用

It was almost no effort to integrate `PlayerStore` into our new application (once we had made the package adjustments) and subsequently testing was very easy too because we decided to expose our stub version too. 

将“PlayerStore”集成到我们的新应用程序中几乎毫不费力（一旦我们进行了包调整），随后的测试也非常容易，因为我们也决定公开我们的存根版本。

