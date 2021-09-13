# WebSockets

# WebSockets

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/websockets)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/websockets)**

In this chapter we'll learn how to use WebSockets to improve our application.

在本章中，我们将学习如何使用 WebSockets 来改进我们的应用程序。

## Project recap

## 项目回顾

We have two applications in our poker codebase

我们的扑克代码库中有两个应用程序

- *Command line app*. Prompts the user to enter the number of players in a game. From then on informs the players of what the "blind bet" value is, which increases over time. At any point a user can enter `"{Playername} wins"` to finish the game and record the victor in a store.
- *Web app*. Allows users to record winners of games and displays a league table. Shares the same store as the command line app.

- *命令行应用程序*。提示用户输入游戏中的玩家人数。从那时起，通知玩家“盲注”值是多少，该值会随着时间的推移而增加。在任何时候，用户都可以输入“{Playername} wins”来完成游戏并在商店中记录胜利者。
- *网络应用程序*。允许用户记录比赛的获胜者并显示排行榜。与命令行应用程序共享同一个商店。

## Next steps

##  下一步

The product owner is thrilled with the command line application but would prefer it if we could bring that functionality to the browser. She imagines a web page with a text box that allows the user to enter the number of players and when they submit the form the page displays the blind value and automatically updates it when appropriate. Like the command line application the user can declare the winner and it'll get saved in the database.

产品所有者对命令行应用程序感到非常兴奋，但如果我们可以将该功能引入浏览器，则更喜欢它。她设想了一个带有文本框的网页，允许用户输入玩家人数，当他们提交表单时，页面会显示盲注值并在适当的时候自动更新它。与命令行应用程序一样，用户可以宣布获胜者并将其保存在数据库中。

On the face of it, it sounds quite simple but as always we must emphasise taking an _iterative_ approach to writing software.

从表面上看，这听起来很简单，但与往常一样，我们必须强调采用 _iterative_ 方法来编写软件。

First of all we will need to serve HTML. So far all of our HTTP endpoints have returned either plaintext or JSON. We _could_ use the same techniques we know (as they're all ultimately strings) but we can also use the [html/template](https://golang.org/pkg/html/template/) package for a cleaner solution.

首先，我们需要提供 HTML。到目前为止，我们所有的 HTTP 端点都返回了纯文本或 JSON。我们_可以_使用我们所知道的相同技术（因为它们最终都是字符串），但我们也可以使用 [html/template](https://golang.org/pkg/html/template/) 包来获得更简洁的解决方案。

We also need to be able to asynchronously send messages to the user saying `The blind is now *y*` without having to refresh the browser. We can use [WebSockets](https://en.wikipedia.org/wiki/WebSocket) to facilitate this.

我们还需要能够异步地向用户发送消息，说“盲人现在 *y*”，而不必刷新浏览器。我们可以使用 [WebSockets](https://en.wikipedia.org/wiki/WebSocket) 来促进这一点。

> WebSocket is a computer communications protocol, providing full-duplex communication channels over a single TCP connection

> WebSocket 是一种计算机通信协议，通过单个 TCP 连接提供全双工通信通道

Given we are taking on a number of techniques it's even more important we do the smallest amount of useful work possible first and then iterate.

鉴于我们正在采用多种技术，因此更重要的是我们首先尽可能少地做有用的工作，然后再进行迭代。

For that reason the first thing we'll do is create a web page with a form for the user to record a winner. Rather than using a plain form, we will use WebSockets to send that data to our server for it to record.

出于这个原因，我们要做的第一件事是创建一个带有表单的网页，供用户记录获胜者。我们将使用 WebSockets 将该数据发送到我们的服务器以供其记录，而不是使用普通形式。

After that we'll work on the blind alerts by which point we will have a bit of infrastructure code set up.

之后，我们将处理盲目警报，届时我们将设置一些基础设施代码。

### What about tests for the JavaScript ?

### JavaScript 测试怎么样？

There will be some JavaScript written to do this but I won't go in to writing tests.

将有一些 JavaScript 用于执行此操作，但我不会编写测试。

It is of course possible but for the sake of brevity I won't be including any explanations for it.

这当然是可能的，但为了简洁起见，我不会对其进行任何解释。

Sorry folks. Lobby O'Reilly to pay me to make a "Learn JavaScript with tests".

对不起各位。游说 O'Reilly 付钱让我制作“通过测试学习 JavaScript”。

## Write the test first

## 先写测试

First thing we need to do is serve up some HTML to users when they hit `/game`.

我们需要做的第一件事是在用户点击 `/game` 时向用户提供一些 HTML。

Here's a reminder of the pertinent code in our web server

这是我们网络服务器中相关代码的提醒

```go
type PlayerServer struct {
    store PlayerStore
    http.Handler
}

const jsonContentType = "application/json"

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

The _easiest_ thing we can do for now is check when we `GET /game` that we get a `200`.

我们现在可以做的_最简单的_事情是检查我们在 `GET /game` 时是否得到了 `200`。

```go
func TestGame(t *testing.T) {
    t.Run("GET /game returns 200", func(t *testing.T) {
        server := NewPlayerServer(&StubPlayerStore{})

        request, _ := http.NewRequest(http.MethodGet, "/game", nil)
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusOK)
    })
}
```

## Try to run the test
```
--- FAIL: TestGame (0.00s)
=== RUN   TestGame/GET_/game_returns_200
    --- FAIL: TestGame/GET_/game_returns_200 (0.00s)
        server_test.go:109: did not get correct status, got 404, want 200
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Our server has a router setup so it's relatively easy to fix.

我们的服务器有一个路由器设置，所以修复起来相对容易。

To our router add

到我们的路由器添加

```go
router.Handle("/game", http.HandlerFunc(p.game))
```

And then write the `game` method

然后编写`game`方法

```go
func (p *PlayerServer) game(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}
```

## Refactor

## 重构

The server code is already fine due to us slotting in more code into the existing well-factored code very easily. 

服务器代码已经很好，因为我们很容易将更多代码插入到现有的良好分解的代码中。

We can tidy up the test a little by adding a test helper function `newGameRequest` to make the request to `/game`. Try writing this yourself.

我们可以通过添加一个测试辅助函数`newGameRequest` 来向`/game` 发出请求，从而稍微整理一下测试。尝试自己写这个。

```go
func TestGame(t *testing.T) {
    t.Run("GET /game returns 200", func(t *testing.T) {
        server := NewPlayerServer(&StubPlayerStore{})

        request :=  newGameRequest()
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response, http.StatusOK)
    })
}
```

You'll also notice I changed `assertStatus` to accept `response` rather than `response.Code` as I feel it reads better.

您还会注意到我将 `assertStatus` 更改为接受 `response` 而不是 `response.Code`，因为我觉得它读起来更好。

Now we need to make the endpoint return some HTML, here it is

现在我们需要让端点返回一些 HTML，这里是

```html

```html

<!DOCTYPE html>
<html lang="en">
<head>
     <meta charset="UTF-8">
     <title>Let's play poker</title>
</head>
<body>
<section id="game">
     <div id="declare-winner">
         <label for="winner">Winner</label>
         <input type="text" id="winner"/>
         <button id="winner-button">Declare winner</button>
     </div>
</section>
</body>
<script type="application/javascript">

     <!DOCTYPE html>
<html lang="zh">
<头>
    <meta charset="UTF-8">
    <title>打牌吧</title>
</头>
<身体>
<section id="游戏">
    <div id="declare-winner">
        <label for="winner">优胜者</label>
        <input type="text" id="winner"/>
        <button id="winner-button">宣布获胜者</button>
    </div>
</section>
</正文>
<script type="application/javascript">

    const submitWinnerButton = document.getElementById('winner-button')
     const winnerInput = document.getElementById('winner')

     const submitWinnerButton = document.getElementById('winner-button')
    const winnerInput = document.getElementById('winner')

    if (window['WebSocket']) {
         const conn = new WebSocket('ws://' + document.location.host + '/ws')

         如果（窗口['WebSocket']）{
        const conn = new WebSocket('ws://' + document.location.host + '/ws')

        submitWinnerButton.onclick = event => {
             conn.send(winnerInput.value)
         }
     }
</script>
</html>
```

submitWinnerButton.onclick = 事件 => {
            conn.send(winnerInput.value)
        }
    }
</脚本>
</html>
``

We have a very simple web page

  我们有一个非常简单的网页

 - A text input for the user to enter the winner into
  - A button they can click to declare the winner.
  - Some JavaScript to open a WebSocket connection to our server and handle the submit button being pressed

- 供用户输入获胜者的文本输入
 - 他们可以点击一个按钮来宣布获胜者。
 - 一些 JavaScript 打开到我们服务器的 WebSocket 连接并处理被按下的提交按钮

`WebSocket` is built into most modern browsers so we don't need to worry about bringing in any libraries. The web page won't work for older browsers, but we're ok with that for this scenario.

`WebSocket` 内置于大多数现代浏览器中，因此我们无需担心引入任何库。该网页不适用于较旧的浏览器，但对于这种情况，我们可以接受。

### How do we test we return the correct markup?

### 我们如何测试我们返回正确的标记？

There are a few ways. As has been emphasised throughout the book, it is important that the tests you write have sufficient value to justify the cost.

有几种方法。正如整本书所强调的那样，您编写的测试具有足够的价值来证明成本是合理的，这一点很重要。

1. Write a browser based test, using something like Selenium. These tests are the most "realistic" of all approaches because they start an actual web browser of some kind and simulates a user interacting with it. These tests can give you a lot of confidence your system works but are more difficult to write than unit tests and much slower to run. For the purposes of our product this is overkill.
2. Do an exact string match. This _can_ be ok but these kind of tests end up being very brittle. The moment someone changes the markup you will have a test failing when in practice nothing has _actually broken_.
3. Check we call the correct template. We will be using a templating library from the standard lib to serve the HTML (discussed shortly) and we could inject in the _thing_ to generate the HTML and spy on its call to check we're doing it right. This would have an impact on our code's design but doesn't actually test a great deal; other than we're calling it with the correct template file. Given we will only have the one template in our project the chance of failure here seems low.

1. 使用 Selenium 之类的东西编写基于浏览器的测试。这些测试是所有方法中最“现实”的，因为它们启动某种实际的 Web 浏览器并模拟用户与之交互。这些测试可以让您对系统工作充满信心，但比单元测试更难编写，运行速度也慢得多。就我们的产品而言，这太过分了。
2. 进行精确的字符串匹配。这_可以_没问题，但这些类型的测试最终非常脆弱。当有人更改标记时，您的测试将失败，而实际上没有任何东西_实际损坏_。
3. 检查我们调用了正确的模板。我们将使用标准库中的模板库来提供 HTML（稍后讨论），我们可以注入 _thing_ 以生成 HTML 并监视它的调用以检查我们是否做得对。这会对我们的代码设计产生影响，但实际上并没有进行大量测试；除了我们用正确的模板文件调用它。鉴于我们的项目中只有一个模板，这里失败的可能性似乎很低。

So in the book "Learn Go with Tests" for the first time, we're not going to write a test.

因此，在第一次在“Learn Go with Tests”一书中，我们不会编写测试。

Put the markup in a file called `game.html`

将标记放在名为“game.html”的文件中

Next change the endpoint we just wrote to the following

接下来将我们刚刚写入的端点更改为以下内容

```go
func (p *PlayerServer) game(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("game.html")

    if err != nil {
        http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}
```

[`html/template`](https://golang.org/pkg/html/template/) is a Go package for creating HTML. In our case we call `template.ParseFiles`, giving the path of our html file. Assuming there is no error you can then `Execute` the template, which writes it to an `io.Writer`. In our case we want it to `Write` to the internet, so we give it our `http.ResponseWriter`.

[`html/template`](https://golang.org/pkg/html/template/) 是一个用于创建 HTML 的 Go 包。在我们的例子中，我们调用`template.ParseFiles`，给出我们的 html 文件的路径。假设没有错误，您可以“执行”模板，将其写入一个“io.Writer”。在我们的例子中，我们希望它`Write`到互联网，所以我们给它我们的`http.ResponseWriter`。

As we have not written a test, it would be prudent to manually test our web server just to make sure things are working as we'd hope. Go to `cmd/webserver` and run the `main.go` file. Visit `http://localhost:5000/game`. 

由于我们还没有编写测试，因此谨慎的做法是手动测试我们的 Web 服务器，以确保一切正常运行。转到“cmd/webserver”并运行“main.go”文件。访问`http://localhost:5000/game`。

You _should_ have got an error about not being able to find the template. You can either change the path to be relative to your folder, or you can have a copy of the `game.html` in the `cmd/webserver` directory. I chose to create a symlink (`ln -s ../../game.html game.html`) to the file inside the root of the project so if I make changes they are reflected when running the server.

您_应该_收到关于无法找到模板的错误。您可以将路径更改为相对于您的文件夹，或者您可以在 `cmd/webserver` 目录中拥有一份 `game.html` 的副本。我选择创建一个符号链接（`ln -s ../../game.html game.html`）到项目根目录中的文件，所以如果我进行更改，它们会在运行服务器时反映出来。

If you make this change and run again you should see our UI.

如果您进行此更改并再次运行，您应该会看到我们的 UI。

Now we need to test that when we get a string over a WebSocket connection to our server that we declare it as a winner of a game.

现在我们需要测试一下，当我们通过 WebSocket 连接到我们的服务器获得一个字符串时，我们将其声明为游戏的赢家。

## Write the test first

## 先写测试

For the first time we are going to use an external library so that we can work with WebSockets.

我们将第一次使用外部库，以便我们可以使用 WebSockets。

Run `go get github.com/gorilla/websocket`

运行`go get github.com/gorilla/websocket`

This will fetch the code for the excellent [Gorilla WebSocket](https://github.com/gorilla/websocket) library. Now we can update our tests for our new requirement.

这将获取优秀的 [Gorilla WebSocket](https://github.com/gorilla/websocket) 库的代码。现在我们可以针对新的需求更新我们的测试。

```go
t.Run("when we get a message over a websocket it is a winner of a game", func(t *testing.T) {
    store := &StubPlayerStore{}
    winner := "Ruth"
    server := httptest.NewServer(NewPlayerServer(store))
    defer server.Close()

    wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

    ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
    if err != nil {
        t.Fatalf("could not open a ws connection on %s %v", wsURL, err)
    }
    defer ws.Close()

    if err := ws.WriteMessage(websocket.TextMessage, []byte(winner));err != nil {
        t.Fatalf("could not send message over ws connection %v", err)
    }

    AssertPlayerWin(t, store, winner)
})
```

Make sure that you have an import for the `websocket` library. My IDE automatically did it for me, so should yours.

确保你有一个 `websocket` 库的导入。我的 IDE 自动为我完成了这项工作，您的也应该如此。

To test what happens from the browser we have to open up our own WebSocket connection and write to it.

为了测试浏览器发生了什么，我们必须打开我们自己的 WebSocket 连接并写入它。

Our previous tests around our server just called methods on our server but now we need to have a persistent connection to our server. To do that we use `httptest.NewServer` which takes a `http.Handler` and will spin it up and listen for connections.

我们之前围绕服务器的测试只是调用了我们服务器上的方法，但现在我们需要与我们的服务器建立持久连接。为此，我们使用 `httptest.NewServer`，它接受一个 `http.Handler` 并启动它并监听连接。

Using `websocket.DefaultDialer.Dial` we try to dial in to our server and then we'll try and send a message with our `winner`.

使用 `websocket.DefaultDialer.Dial` 我们尝试拨入我们的服务器，然后我们将尝试使用我们的 `winner` 发送消息。

Finally we assert on the player store to check the winner was recorded.

最后，我们在播放器商店断言以检查获胜者是否被记录。

## Try to run the test
```
=== RUN   TestGame/when_we_get_a_message_over_a_websocket_it_is_a_winner_of_a_game
    --- FAIL: TestGame/when_we_get_a_message_over_a_websocket_it_is_a_winner_of_a_game (0.00s)
        server_test.go:124: could not open a ws connection on ws://127.0.0.1:55838/ws websocket: bad handshake
```

We have not changed our server to accept WebSocket connections on `/ws` so we're not shaking hands yet.

我们还没有改变我们的服务器来接受 `/ws` 上的 WebSocket 连接，所以我们还没有握手。

## Write enough code to make it pass

## 编写足够的代码使其通过

Add another listing to our router

向我们的路由器添加另一个列表

```go
router.Handle("/ws", http.HandlerFunc(p.webSocket))
```

Then add our new `webSocket` handler

然后添加我们新的 `webSocket` 处理程序

```go
func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
    upgrader := websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
    }
    upgrader.Upgrade(w, r, nil)
}
```

To accept a WebSocket connection we `Upgrade` the request. If you now re-run the test you should move on to the next error.

为了接受 WebSocket 连接，我们“升级”请求。如果您现在重新运行测试，您应该继续处理下一个错误。

```
=== RUN   TestGame/when_we_get_a_message_over_a_websocket_it_is_a_winner_of_a_game
    --- FAIL: TestGame/when_we_get_a_message_over_a_websocket_it_is_a_winner_of_a_game (0.00s)
        server_test.go:132: got 0 calls to RecordWin want 1
```

Now that we have a connection opened, we'll want to listen for a message and then record it as the winner.

现在我们已经打开了一个连接，我们要监听一条消息，然后将它记录为获胜者。

```go
func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
    upgrader := websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
    }
    conn, _ := upgrader.Upgrade(w, r, nil)
    _, winnerMsg, _ := conn.ReadMessage()
    p.store.RecordWin(string(winnerMsg))
}
```

(Yes, we're ignoring a lot of errors right now!)

（是的，我们现在忽略了很多错误！）

`conn.ReadMessage()` blocks on waiting for a message on the connection. Once we get one we use it to `RecordWin`. This would finally close the WebSocket connection.

`conn.ReadMessage()` 阻塞等待连接上的消息。一旦我们得到一个，我们就用它来`RecordWin`。这将最终关闭 WebSocket 连接。

If you try and run the test, it's still failing.

如果您尝试运行测试，它仍然失败。

The issue is timing. There is a delay between our WebSocket connection reading the message and recording the win and our test finishes before it happens. You can test this by putting a short `time.Sleep` before the final assertion.

问题是时机。我们的 WebSocket 连接读取消息和记录胜利之间存在延迟，我们的测试在它发生之前完成。您可以通过在最终断言之前放置一个简短的 `time.Sleep` 来测试这一点。

Let's go with that for now but acknowledge that putting in arbitrary sleeps into tests **is very bad practice**.

现在让我们继续这样做，但承认将任意睡眠放入测试**是非常糟糕的做法**。

```go
time.Sleep(10 * time.Millisecond)
AssertPlayerWin(t, store, winner)
```

## Refactor 

## 重构

We committed many sins to make this test work both in the server code and the test code but remember this is the easiest way for us to work.

我们犯了很多罪，使这个测试在服务器代码和测试代码中都能工作，但请记住，这是我们最简单的工作方式。

We have nasty, horrible, _working_ software backed by a test, so now we are free to make it nice and know we won't break anything accidentally.

我们有由测试支持的讨厌、可怕的 _working_ 软件，所以现在我们可以自由地让它变得更好，并且知道我们不会意外破坏任何东西。

Let's start with the server code.

让我们从服务器代码开始。

We can move the `upgrader` to a private value inside our package because we don't need to redeclare it on every WebSocket connection request

我们可以将 `upgrader` 移动到包内的私有值，因为我们不需要在每个 WebSocket 连接请求上重新声明它

```go
var wsUpgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
    conn, _ := wsUpgrader.Upgrade(w, r, nil)
    _, winnerMsg, _ := conn.ReadMessage()
    p.store.RecordWin(string(winnerMsg))
}
```

Our call to `template.ParseFiles("game.html")` will run on every `GET /game` which means we'll go to the file system on every request even though we have no need to re-parse the template. Let's refactor our code so that we parse the template once in `NewPlayerServer` instead. We'll have to make it so this function can now return an error in case we have problems fetching the template from disk or parsing it.

我们对`template.ParseFiles("game.html")` 的调用将在每个`GET /game` 上运行，这意味着即使我们不需要重新解析模板，我们也会在每个请求上访问文件系统。让我们重构我们的代码，以便我们在 `NewPlayerServer` 中解析模板一次。我们必须这样做，以便此函数现在可以返回错误，以防我们从磁盘获取模板或解析它时遇到问题。

Here's the relevant changes to `PlayerServer`

这是对“PlayerServer”的相关更改

```go
type PlayerServer struct {
    store PlayerStore
    http.Handler
    template *template.Template
}

const htmlTemplatePath = "game.html"

func NewPlayerServer(store PlayerStore) (*PlayerServer, error) {
    p := new(PlayerServer)

    tmpl, err := template.ParseFiles(htmlTemplatePath)

    if err != nil {
        return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
    }

    p.template = tmpl
    p.store = store

    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))
    router.Handle("/game", http.HandlerFunc(p.game))
    router.Handle("/ws", http.HandlerFunc(p.webSocket))

    p.Handler = router

    return p, nil
}

func (p *PlayerServer) game(w http.ResponseWriter, r *http.Request) {
    p.template.Execute(w, nil)
}
```

By changing the signature of `NewPlayerServer` we now have compilation problems. Try and fix them yourself or refer to the source code if you struggle.

通过更改 NewPlayerServer 的签名，我们现在遇到了编译问题。如果您遇到困难，请尝试自己修复它们或参考源代码。

For the test code I made a helper called `mustMakePlayerServer(t *testing.T, store PlayerStore) *PlayerServer` so that I could hide the error noise away from the tests.

对于测试代码，我创建了一个名为“mustMakePlayerServer(t *testing.T, store PlayerStore) *PlayerServer”的帮助程序，以便我可以在测试中隐藏错误噪音。

```go
func mustMakePlayerServer(t *testing.T, store PlayerStore) *PlayerServer {
    server, err := NewPlayerServer(store)
    if err != nil {
        t.Fatal("problem creating player server", err)
    }
    return server
}
```

Similarly I created another helper `mustDialWS` so that I could hide nasty error noise when creating the WebSocket connection.

同样，我创建了另一个助手 `mustDialWS`，以便在创建 WebSocket 连接时隐藏令人讨厌的错误噪音。

```go
func mustDialWS(t *testing.T, url string) *websocket.Conn {
    ws, _, err := websocket.DefaultDialer.Dial(url, nil)

    if err != nil {
        t.Fatalf("could not open a ws connection on %s %v", url, err)
    }

    return ws
}
```

Finally in our test code we can create a helper to tidy up sending messages

最后在我们的测试代码中，我们可以创建一个帮助程序来整理发送消息

```go
func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
    t.Helper()
    if err := conn.WriteMessage(websocket.TextMessage, []byte(message));err != nil {
        t.Fatalf("could not send message over ws connection %v", err)
    }
}
```

Now the tests are passing try running the server and declare some winners in `/game`. You should see them recorded in `/league`. Remember that every time we get a winner we _close the connection_, you will need to refresh the page to open the connection again.

现在测试通过了尝试运行服务器并在`/game`中宣布一些获胜者。你应该看到他们记录在 `/league` 中。请记住，每次我们获胜时我们_关闭连接_，您将需要刷新页面以再次打开连接。

We've made a trivial web form that lets users record the winner of a game. Let's iterate on it to make it so the user can start a game by providing a number of players and the server will push messages to the client informing them of what the blind value is as time passes.

我们制作了一个简单的网络表单，让用户可以记录游戏的获胜者。让我们对其进行迭代，以便用户可以通过提供多个玩家来开始游戏，并且服务器将向客户端推送消息，通知他们随着时间的推移盲值是多少。

First of all update `game.html` to update our client side code for the new requirements

首先更新`game.html`以更新我们的客户端代码以适应新要求

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Lets play poker</title>
</head>
<body>
<section id="game">
    <div id="game-start">
        <label for="player-count">Number of players</label>
        <input type="number" id="player-count"/>
        <button id="start-game">Start</button>
    </div>

    <div id="declare-winner">
        <label for="winner">Winner</label>
        <input type="text" id="winner"/>
        <button id="winner-button">Declare winner</button>
    </div>

    <div id="blind-value"/>
</section>

<section id="game-end">
    <h1>Another great game of poker everyone!</h1>
    <p><a href="/league">Go check the league table</a></p>
</section>

</body>
<script type="application/javascript">
    const startGame = document.getElementById('game-start')

    const declareWinner = document.getElementById('declare-winner')
    const submitWinnerButton = document.getElementById('winner-button')
    const winnerInput = document.getElementById('winner')

    const blindContainer = document.getElementById('blind-value')

    const gameContainer = document.getElementById('game')
    const gameEndContainer = document.getElementById('game-end')

    declareWinner.hidden = true
    gameEndContainer.hidden = true

    document.getElementById('start-game').addEventListener('click', event => {
        startGame.hidden = true
        declareWinner.hidden = false

        const numberOfPlayers = document.getElementById('player-count').value

        if (window['WebSocket']) {
            const conn = new WebSocket('ws://' + document.location.host + '/ws')

            submitWinnerButton.onclick = event => {
                conn.send(winnerInput.value)
                gameEndContainer.hidden = false
                gameContainer.hidden = true
            }

            conn.onclose = evt => {
                blindContainer.innerText = 'Connection closed'
            }

            conn.onmessage = evt => {
                blindContainer.innerText = evt.data
            }

            conn.onopen = function () {
                conn.send(numberOfPlayers)
            }
        }
    })
</script>
</html>
```

The main changes is bringing in a section to enter the number of players and a section to display the blind value. We have a little logic to show/hide the user interface depending on the stage of the game.

主要的变化是引入了一个输入玩家数量的部分和一个显示盲值的部分。我们有一些逻辑来根据游戏阶段显示/隐藏用户界面。

Any message we receive via `conn.onmessage` we assume to be blind alerts and so we set the `blindContainer.innerText` accordingly.

我们通过 `conn.onmessage` 收到的任何消息我们都假设是盲目警报，因此我们相应地设置了 `blindContainer.innerText`。

How do we go about sending the blind alerts? In the previous chapter we introduced the idea of `Game` so our CLI code could call a `Game` and everything else would be taken care of including scheduling blind alerts. This turned out to be a good separation of concern.

我们如何发送盲目警报？在上一章中，我们介绍了“游戏”的概念，因此我们的 CLI 代码可以调用“游戏”，而其他一切都将得到处理，包括调度盲警报。结果证明这是一个很好的关注点分离。

```go
type Game interface {
    Start(numberOfPlayers int)
    Finish(winner string)
}
```

When the user was prompted in the CLI for number of players it would `Start` the game which would kick off the blind alerts and when the user declared the winner they would `Finish`. This is the same requirements we have now, just a different way of getting the inputs; so we should look to re-use this concept if we can.

当用户在 CLI 中被提示输入玩家数量时，它将“开始”游戏，这将启动盲注，当用户宣布获胜者时，他们将“结束”。这与我们现在的要求相同，只是获取输入的方式不同；所以我们应该尽可能地重新使用这个概念。

Our "real" implementation of `Game` is `TexasHoldem`

我们对`Game`的“真正”实现是`TexasHoldem`

```go
type TexasHoldem struct {
    alerter BlindAlerter
    store   PlayerStore
}
```

By sending in a `BlindAlerter` `TexasHoldem` can schedule blind alerts to be sent to _wherever_

通过发送“BlindAlerter”，“TexasHoldem”可以安排将盲警报发送到_wherever_

```go
type BlindAlerter interface {
    ScheduleAlertAt(duration time.Duration, amount int)
}
```

And as a reminder, here is our implementation of the `BlindAlerter` we use in the CLI.

提醒一下，这是我们在 CLI 中使用的 `BlindAlerter` 的实现。

```go
func StdOutAlerter(duration time.Duration, amount int) {
    time.AfterFunc(duration, func() {
        fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
    })
}
```

This works in CLI because we _always want to send the alerts to `os.Stdout`_ but this won't work for our web server. For every request we get a new `http.ResponseWriter` which we then upgrade to `*websocket.Conn`. So we can't know when constructing our dependencies where our alerts need to go.

这在 CLI 中有效，因为我们 _always 想要将警报发送到 `os.Stdout`_ 但这不适用于我们的 Web 服务器。对于每个请求，我们都会得到一个新的 `http.ResponseWriter`，然后我们将其升级为 `*websocket.Conn`。所以我们不知道在构建依赖项时我们的警报需要去哪里。

For that reason we need to change `BlindAlerter.ScheduleAlertAt` so that it takes a destination for the alerts so that we can re-use it in our webserver.

出于这个原因，我们需要更改“BlindAlerter.ScheduleAlertAt”，以便它为警报指定一个目的地，以便我们可以在我们的网络服务器中重新使用它。

Open BlindAlerter.go and add the parameter `to io.Writer`

打开 BlindAlerter.go 并添加参数 `to io.Writer`

```go
type BlindAlerter interface {
    ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
    a(duration, amount, to)
}
```

The idea of a `StdoutAlerter` doesn't fit our new model so just rename it to `Alerter`

`StdoutAlerter` 的想法不适合我们的新模型，所以只需将其重命名为 `Alerter`

```go
func Alerter(duration time.Duration, amount int, to io.Writer) {
    time.AfterFunc(duration, func() {
        fmt.Fprintf(to, "Blind is now %d\n", amount)
    })
}
```

If you try and compile, it will fail in `TexasHoldem` because it is calling `ScheduleAlertAt` without a destination, to get things compiling again _for now_ hard-code it to `os.Stdout`.

如果你尝试编译，它会在 `TexasHoldem` 中失败，因为它在没有目标的情况下调用 `ScheduleAlertAt`，为了让事情再次编译_暂时_硬编码到 `os.Stdout`。

Try and run the tests and they will fail because `SpyBlindAlerter` no longer implements `BlindAlerter`, fix this by updating the signature of `ScheduleAlertAt`, run the tests and we should still be green.

尝试运行测试，它们会失败，因为 `SpyBlindAlerter` 不再实现 `BlindAlerter`，通过更新 `ScheduleAlertAt` 的签名来解决这个问题，运行测试，我们应该仍然是绿色的。

It doesn't make any sense for `TexasHoldem` to know where to send blind alerts. Let's now update `Game` so that when you start a game you declare _where_ the alerts should go.

“TexasHoldem” 知道在哪里发送盲目警报没有任何意义。现在让我们更新“游戏”，以便在您开始游戏时声明警报应该去哪里。

```go
type Game interface {
    Start(numberOfPlayers int, alertsDestination io.Writer)
    Finish(winner string)
}
```

Let the compiler tell you what you need to fix. The change isn't so bad:

让编译器告诉您需要修复什么。变化并不是那么糟糕：

- Update `TexasHoldem` so it properly implements `Game`
- In `CLI` when we start the game, pass in our `out` property (`cli.game.Start(numberOfPlayers, cli.out)`)
- In `TexasHoldem`'s test i use `game.Start(5, ioutil.Discard)` to fix the compilation problem and configure the alert output to be discarded

- 更新 `TexasHoldem` 使其正确实现 `Game`
- 在 `CLI` 中，当我们开始游戏时，传入我们的 `out` 属性（`cli.game.Start(numberOfPlayers, cli.out)`）
- 在`TexasHoldem` 的测试中，我使用`game.Start(5, ioutil.Discard)` 来修复编译问题并将警报输出配置为丢弃

If you've got everything right, everything should be green! Now we can try and use `Game` within `Server`.

如果一切顺利，一切都应该是绿色的！现在我们可以尝试在 `Server` 中使用 `Game`。

## Write the test first

## 先写测试

The requirements of `CLI` and `Server` are the same! It's just the delivery mechanism is different.

`CLI` 和 `Server` 的要求是一样的！只是传递机制不同而已。

Let's take a look at our `CLI` test for inspiration.

让我们来看看我们的 `CLI` 测试以获得灵感。

```go
t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
    game := &GameSpy{}

    out := &bytes.Buffer{}
    in := userSends("3", "Chris wins")

    poker.NewCLI(in, out, game).PlayPoker()

    assertMessagesSentToUser(t, out, poker.PlayerPrompt)
    assertGameStartedWith(t, game, 3)
    assertFinishCalledWith(t, game, "Chris")
})
```

It looks like we should be able to test drive out a similar outcome using `GameSpy`

看起来我们应该能够使用“GameSpy”来测试类似的结果

Replace the old websocket test with the following

用以下内容替换旧的 websocket 测试

```go
t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
    game := &poker.GameSpy{}
    winner := "Ruth"
    server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
    ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

    defer server.Close()
    defer ws.Close()

    writeWSMessage(t, ws, "3")
    writeWSMessage(t, ws, winner)

    time.Sleep(10 * time.Millisecond)
    assertGameStartedWith(t, game, 3)
    assertFinishCalledWith(t, game, winner)
})
```

- As discussed we create a spy `Game` and pass it into `mustMakePlayerServer` (be sure to update the helper to support this).
- We then send the web socket messages for a game.
- Finally we assert that the game is started and finished with what we expect.

- 如前所述，我们创建了一个间谍 `Game` 并将其传递给 `mustMakePlayerServer`（确保更新助手以支持此功能）。
- 然后我们发送游戏的网络套接字消息。
- 最后，我们断言游戏以我们期望的方式开始和结束。

## Try to run the test

## 尝试运行测试

You'll have a number of compilation errors around `mustMakePlayerServer` in other tests. Introduce an unexported variable `dummyGame` and use it through all the tests that aren't compiling

在其他测试中，您会在 `mustMakePlayerServer` 周围遇到许多编译错误。引入一个未导出的变量 `dummyGame` 并在所有未编译的测试中使用它

```go
var (
    dummyGame = &GameSpy{}
)
```

The final error is where we are trying to pass in `Game` to `NewPlayerServer` but it doesn't support it yet

最后一个错误是我们试图将 `Game` 传递给 `NewPlayerServer` 但它尚不支持它

```
./server_test.go:21:38: too many arguments in call to "github.com/quii/learn-go-with-tests/WebSockets/v2".NewPlayerServer
    have ("github.com/quii/learn-go-with-tests/WebSockets/v2".PlayerStore, "github.com/quii/learn-go-with-tests/WebSockets/v2".Game)
    want ("github.com/quii/learn-go-with-tests/WebSockets/v2".PlayerStore)
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Just add it as an argument for now just to get the test running

现在只需将其添加为参数即可运行测试

```go
func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
```

Finally!

最后！

```
=== RUN   TestGame/start_a_game_with_3_players_and_declare_Ruth_the_winner
--- FAIL: TestGame (0.01s)
    --- FAIL: TestGame/start_a_game_with_3_players_and_declare_Ruth_the_winner (0.01s)
        server_test.go:146: wanted Start called with 3 but got 0
        server_test.go:147: expected finish called with 'Ruth' but got ''
FAIL
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We need to add `Game` as a field to `PlayerServer` so that it can use it when it gets requests.

我们需要将“Game”作为字段添加到“PlayerServer”，以便它在收到请求时可以使用它。

```go
type PlayerServer struct {
    store PlayerStore
    http.Handler
    template *template.Template
    game Game
}
```

(We already have a method called `game` so rename that to `playGame`)

（我们已经有一个名为“game”的方法，因此将其重命名为“playGame”）

Next lets assign it in our constructor

接下来让我们在我们的构造函数中分配它

```go
func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
    p := new(PlayerServer)

    tmpl, err := template.ParseFiles(htmlTemplatePath)

    if err != nil {
        return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
    }

    p.game = game

    // etc
```

Now we can use our `Game` within `webSocket`.

现在我们可以在 `webSocket` 中使用我们的 `Game`。

```go
func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
    conn, _ := wsUpgrader.Upgrade(w, r, nil)

    _, numberOfPlayersMsg, _ := conn.ReadMessage()
    numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
    p.game.Start(numberOfPlayers, ioutil.Discard) //todo: Don't discard the blinds messages!

    _, winner, _ := conn.ReadMessage()
    p.game.Finish(string(winner))
}
```

Hooray! The tests pass.

万岁！测试通过。

We are not going to send the blind messages anywhere _just yet_ as we need to have a think about that. When we call `game.Start` we send in `ioutil.Discard` which will just discard any messages written to it.

我们不会_只是_在任何地方发送盲目消息，因为我们需要考虑一下。当我们调用`game.Start` 时，我们发送了`ioutil.Discard`，它只会丢弃任何写入它的消息。

For now start the web server up. You'll need to update the `main.go` to pass a `Game` to the `PlayerServer`

现在启动 Web 服务器。您需要更新 `main.go` 以将 `Game` 传递给 `PlayerServer`

```go
func main() {
    db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

    if err != nil {
        log.Fatalf("problem opening %s %v", dbFileName, err)
    }

    store, err := poker.NewFileSystemPlayerStore(db)

    if err != nil {
        log.Fatalf("problem creating file system player store, %v ", err)
    }

    game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

    server, err := poker.NewPlayerServer(store, game)

    if err != nil {
        log.Fatalf("problem creating player server %v", err)
    }

    log.Fatal(http.ListenAndServe(":5000", server))
}
```

Discounting the fact we're not getting blind alerts yet, the app does work! We've managed to re-use `Game` with `PlayerServer` and it has taken care of all the details. Once we figure out how to send our blind alerts through to the web sockets rather than discarding them it _should_ all work.

不考虑我们还没有收到盲目警报的事实，该应用程序确实有效！我们已经成功地将“Game”与“PlayerServer”一起使用，它已经处理了所有细节。一旦我们弄清楚如何将我们的盲目警报发送到网络套接字而不是丢弃它们，它_应该_一切正常。

Before that though, let's tidy up some code.

在此之前，让我们整理一些代码。

## Refactor

## 重构

The way we're using WebSockets is fairly basic and the error handling is fairly naive, so I wanted to encapsulate that in a type just to remove that messiness from the server code. We may wish to revisit it later but for now this'll tidy things up a bit

我们使用 WebSockets 的方式相当基本，错误处理也相当幼稚，所以我想把它封装在一个类型中，只是为了从服务器代码中消除这种混乱。我们可能希望稍后重新访问它，但现在这会整理一下

```go
type playerServerWS struct {
    *websocket.Conn
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
    conn, err := wsUpgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Printf("problem upgrading connection to WebSockets %v\n", err)
    }

    return &playerServerWS{conn}
}

func (w *playerServerWS) WaitForMsg() string {
    _, msg, err := w.ReadMessage()
    if err != nil {
        log.Printf("error reading from websocket %v\n", err)
    }
    return string(msg)
}
```

Now the server code is a bit simplified

现在服务器代码有点简化

```go
func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
    ws := newPlayerServerWS(w, r)

    numberOfPlayersMsg := ws.WaitForMsg()
    numberOfPlayers, _ := strconv.Atoi(numberOfPlayersMsg)
    p.game.Start(numberOfPlayers, ioutil.Discard) //todo: Don't discard the blinds messages!

    winner := ws.WaitForMsg()
    p.game.Finish(winner)
}
```

Once we figure out how to not discard the blind messages we're done.

一旦我们弄清楚如何不丢弃盲消息，我们就完成了。

### Let's _not_ write a test!

### 让我们_不_写一个测试！

Sometimes when we're not sure how to do something, it's best just to play around and try things out! Make sure your work is committed first because once we've figured out a way we should drive it through a test.

有时，当我们不确定如何做某事时，最好只是四处玩耍并尝试一下！确保首先提交您的工作，因为一旦我们找到了一种方法，我们就应该通过测试来推动它。

The problematic line of code we have is

我们有问题的代码行是

```go
p.game.Start(numberOfPlayers, ioutil.Discard) //todo: Don't discard the blinds messages!
```

We need to pass in an `io.Writer` for the game to write the blind alerts to.

我们需要传入一个 `io.Writer` 来让游戏写入盲注警报。

Wouldn't it be nice if we could pass in our `playerServerWS` from before? It's our wrapper around our WebSocket so it _feels_ like we should be able to send that to our `Game` to send messages to.

如果我们可以从以前传入我们的 `playerServerWS` 不是很好吗？它是我们 WebSocket 的包装器，因此_感觉_我们应该能够将其发送到我们的“游戏”以向其发送消息。

Give it a go:

搏一搏：

```go
func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
    ws := newPlayerServerWS(w, r)

    numberOfPlayersMsg := ws.WaitForMsg()
    numberOfPlayers, _ := strconv.Atoi(numberOfPlayersMsg)
    p.game.Start(numberOfPlayers, ws)
    //etc...
```

The compiler complains

编译器抱怨

```
./server.go:71:14: cannot use ws (type *playerServerWS) as type io.Writer in argument to p.game.Start:
    *playerServerWS does not implement io.Writer (missing Write method)
```

It seems the obvious thing to do, would be to make it so `playerServerWS` _does_ implement `io.Writer`. To do so we use the underlying `*websocket.Conn` to use `WriteMessage` to send the message down the websocket

看起来很明显的事情是让 `playerServerWS` _does_ 实现 `io.Writer`。为此，我们使用底层的 `*websocket.Conn` 使用 `WriteMessage` 将消息向下发送到 websocket

```go
func (w *playerServerWS) Write(p []byte) (n int, err error) {
    err = w.WriteMessage(websocket.TextMessage, p)

    if err != nil {
        return 0, err
    }

    return len(p), nil
}
```

This seems too easy! Try and run the application and see if it works.

这似乎太容易了！尝试运行该应用程序，看看它是否有效。

Beforehand edit `TexasHoldem` so that the blind increment time is shorter so you can see it in action

预先编辑“TexasHoldem”，使盲增时间更短，以便您可以看到它的实际效果

```go
blindIncrement := time.Duration(5+numberOfPlayers) * time.Second // (rather than a minute)
```

You should see it working! The blind amount increments in the browser as if by magic.

你应该看到它在工作！盲量在浏览器中像魔法一样增加。

Now let's revert the code and think how to test it. In order to _implement_ it all we did was pass through to `StartGame` was `playerServerWS` rather than `ioutil.Discard` so that might make you think we should perhaps spy on the call to verify it works.

现在让我们还原代码并思考如何测试它。为了_实现_它，我们所做的只是传递给`StartGame` 的是`playerServerWS` 而不是`ioutil.Discard`，因此这可能会让您认为我们应该监视调用以验证它是否有效。

Spying is great and helps us check implementation details but we should always try and favour testing the _real_ behaviour if we can because when you decide to refactor it's often spy tests that start failing because they are usually checking implementation details that you're trying to change .

间谍很棒，可以帮助我们检查实现细节，但如果可以，我们应该始终尝试并倾向于测试 _real_ 行为，因为当您决定重构时，通常是间谍测试开始失败，因为它们通常会检查您试图更改的实现细节.

Our test currently opens a websocket connection to our running server and sends messages to make it do things. Equally we should be able to test the messages our server sends back over the websocket connection.

我们的测试目前打开一个到我们正在运行的服务器的 websocket 连接并发送消息以使其执行操作。同样，我们应该能够测试我们的服务器通过 websocket 连接发回的消息。

## Write the test first

## 先写测试

We'll edit our existing test.

我们将编辑我们现有的测试。

Currently our `GameSpy` does not send any data to `out` when you call `Start`. We should change it so we can configure it to send a canned message and then we can check that message gets sent to the websocket. This should give us confidence that we have configured things correctly whilst still exercising the real behaviour we want.

当前，当您调用“Start”时，我们的“GameSpy”不会向“out”发送任何数据。我们应该更改它，以便我们可以将其配置为发送固定消息，然后我们可以检查该消息是否已发送到 websocket。这应该让我们相信我们已经正确配置了东西，同时仍然行使我们想要的真实行为。

```go
type GameSpy struct {
    StartCalled     bool
    StartCalledWith int
    BlindAlert      []byte

    FinishedCalled   bool
    FinishCalledWith string
}
```

Add `BlindAlert` field.

添加“BlindAlert”字段。

Update `GameSpy` `Start` to send the canned message to `out`.

更新“GameSpy”“Start”以将预制消息发送到“out”。

```go
func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
    g.StartCalled = true
    g.StartCalledWith = numberOfPlayers
    out.Write(g.BlindAlert)
}
```

This now means when we exercise `PlayerServer` when it tries to `Start` the game it should end up sending messages through the websocket if things are working right.

现在这意味着当我们在尝试“开始”游戏时使用“PlayerServer”时，如果一切正常，它最终应该通过 websocket 发送消息。

Finally we can update the test

最后我们可以更新测试

```go
t.Run("start a game with 3 players, send some blind alerts down WS and declare Ruth the winner", func(t *testing.T) {
    wantedBlindAlert := "Blind is 100"
    winner := "Ruth"

    game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
    server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
    ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

    defer server.Close()
    defer ws.Close()

    writeWSMessage(t, ws, "3")
    writeWSMessage(t, ws, winner)

    time.Sleep(10 * time.Millisecond)
    assertGameStartedWith(t, game, 3)
    assertFinishCalledWith(t, game, winner)

    _, gotBlindAlert, _ := ws.ReadMessage()

    if string(gotBlindAlert) != wantedBlindAlert {
        t.Errorf("got blind alert %q, want %q", string(gotBlindAlert), wantedBlindAlert)
    }
})
```

- We've added a `wantedBlindAlert` and configured our `GameSpy` to send it to `out` if `Start` is called.
- We hope it gets sent in the websocket connection so we've added a call to `ws.ReadMessage()` to wait for a message to be sent and then check it's the one we expected.

- 我们添加了一个 `wantedBlindAlert` 并配置了我们的 `GameSpy` 以在调用 `Start` 时将其发送到 `out`。
- 我们希望它在 websocket 连接中被发送，所以我们添加了一个对 `ws.ReadMessage()` 的调用来等待消息被发送，然后检查它是我们期望的消息。

## Try to run the test

## 尝试运行测试

You should find the test hangs forever. This is because `ws.ReadMessage()` will block until it gets a message, which it never will.

您应该会发现测试永远挂起。这是因为 `ws.ReadMessage()` 会阻塞直到它收到一条消息，而它永远不会。

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We should never have tests that hang so let's introduce a way of handling code that we want to timeout.

我们永远不应该挂起测试，所以让我们介绍一种处理我们想要超时的代码的方法。

```go
func within(t testing.TB, d time.Duration, assert func()) {
    t.Helper()

    done := make(chan struct{}, 1)

    go func() {
        assert()
        done <- struct{}{}
    }()

    select {
    case <-time.After(d):
        t.Error("timed out")
    case <-done:
    }
}
```

What `within` does is take a function `assert` as an argument and then runs it in a go routine. If/When the function finishes it will signal it is done via the `done` channel.

`within` 所做的是将一个函数 `assert` 作为参数，然后在 goroutine 中运行它。如果/当函数完成时，它将通过`done` 通道发出信号。

While that happens we use a `select` statement which lets us wait for a channel to send a message. From here it is a race between the `assert` function and `time.After` which will send a signal when the duration has occurred.

当发生这种情况时，我们使用了一个 `select` 语句，它让我们等待一个通道发送消息。这里是 `assert` 函数和 `time.After` 之间的竞争，它会在持续时间发生时发送信号。

Finally I made a helper function for our assertion just to make things a bit neater

最后，我为我们的断言做了一个辅助函数，只是为了让事情更整洁一点

```go
func assertWebsocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
    _, msg, _ := ws.ReadMessage()
    if string(msg) != want {
        t.Errorf(`got "%s", want "%s"`, string(msg), want)
    }
}
```

Here's how the test reads now

这是测试现在的读取方式

```go
t.Run("start a game with 3 players, send some blind alerts down WS and declare Ruth the winner", func(t *testing.T) {
    wantedBlindAlert := "Blind is 100"
    winner := "Ruth"

    game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
    server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
    ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

    defer server.Close()
    defer ws.Close()

    writeWSMessage(t, ws, "3")
    writeWSMessage(t, ws, winner)

    time.Sleep(tenMS)

    assertGameStartedWith(t, game, 3)
    assertFinishCalledWith(t, game, winner)
    within(t, tenMS, func() { assertWebsocketGotMsg(t, ws, wantedBlindAlert) })
})
```

Now if you run the test...

现在，如果您运行测试...

```
=== RUN   TestGame
=== RUN   TestGame/start_a_game_with_3_players,_send_some_blind_alerts_down_WS_and_declare_Ruth_the_winner
--- FAIL: TestGame (0.02s)
    --- FAIL: TestGame/start_a_game_with_3_players,_send_some_blind_alerts_down_WS_and_declare_Ruth_the_winner (0.02s)
        server_test.go:143: timed out
        server_test.go:150: got "", want "Blind is 100"
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Finally we can now change our server code so it sends our WebSocket connection to the game when it starts

最后，我们现在可以更改我们的服务器代码，以便它在游戏开始时将我们的 WebSocket 连接发送到游戏

```go
func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
    ws := newPlayerServerWS(w, r)

    numberOfPlayersMsg := ws.WaitForMsg()
    numberOfPlayers, _ := strconv.Atoi(numberOfPlayersMsg)
    p.game.Start(numberOfPlayers, ws)

    winner := ws.WaitForMsg()
    p.game.Finish(winner)
}
```

## Refactor

## 重构

The server code was a very small change so there's not a lot to change here but the test code still has a `time.Sleep` call because we have to wait for our server to do its work asynchronously.

服务器代码是一个很小的变化，所以这里没有太多变化，但测试代码仍然有一个 `time.Sleep` 调用，因为我们必须等待我们的服务器异步完成它的工作。

We can refactor our helpers `assertGameStartedWith` and `assertFinishCalledWith` so that they can retry their assertions for a short period before failing.

我们可以重构我们的助手 `assertGameStartedWith` 和 `assertFinishCalledWith`，这样它们就可以在失败之前在短时间内重试他们的断言。

Here's how you can do it for `assertFinishCalledWith` and you can use the same approach for the other helper.

以下是如何为 `assertFinishCalledWith` 执行此操作，并且您可以对其他帮助程序使用相同的方法。

```go
func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
    t.Helper()

    passed := retryUntil(500*time.Millisecond, func() bool {
        return game.FinishCalledWith == winner
    })

    if !passed {
        t.Errorf("expected finish called with %q but got %q", winner, game.FinishCalledWith)
    }
}
```

Here is how `retryUntil` is defined

这里是如何定义 `retryUntil`

```go
func retryUntil(d time.Duration, f func() bool) bool {
    deadline := time.Now().Add(d)
    for time.Now().Before(deadline) {
        if f() {
            return true
        }
    }
    return false
}
```

## Wrapping up 

##  总结

Our application is now complete. A game of poker can be started via a web browser and the users are informed of the blind bet value as time goes by via WebSockets. When the game finishes they can record the winner which is persisted using code we wrote a few chapters ago. The players can find out who is the best (or luckiest) poker player using the website's `/league` endpoint.

我们的应用程序现已完成。扑克游戏可以通过网络浏览器开始，随着时间的推移，用户会通过 WebSockets 获知盲注值。当游戏结束时，他们可以使用我们几章前编写的代码来记录获胜者。玩家可以使用网站的`/league` 端点找出谁是最好的（或最幸运的）扑克玩家。

Through the journey we have made mistakes but with the TDD flow we have never been very far away from working software. We were free to keep iterating and experimenting.

在整个过程中，我们犯了错误，但在 TDD 流程中，我们从未离工作软件很远。我们可以自由地继续迭代和试验。

The final chapter will retrospect on the approach, the design we've arrived at and tie up some loose ends.

最后一章将回顾方法，我们已经达到的设计并解决一些松散的问题。

We covered a few things in this chapter

我们在本章中介绍了一些内容

### WebSockets

### WebSockets

- Convenient way of sending messages between clients and servers that does not require the client to keep polling the server. Both the client and server code we have is very simple.
- Trivial to test, but you have to be wary of the asynchronous nature of the tests

- 在客户端和服务器之间发送消息的便捷方式，不需要客户端不断轮询服务器。我们拥有的客户端和服务器代码都非常简单。
- 测试很简单，但你必须警惕测试的异步性

### Handling code in tests that can be delayed or never finish

### 处理测试中可能延迟或永远不会完成的代码

- Create helper functions to retry assertions and add timeouts.
- We can use go routines to ensure the assertions don't block anything and then use channels to let them signal that they have finished, or not.
- The `time` package has some helpful functions which also send signals via channels about events in time so we can set timeouts 

- 创建辅助函数以重试断言并添加超时。
- 我们可以使用 go 例程来确保断言不会阻塞任何东西，然后使用通道让它们发出信号，表明它们已经完成，或者没有。
- `time` 包有一些有用的功能，这些功能还可以通过通道及时发送有关事件的信号，以便我们可以设置超时

