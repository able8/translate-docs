# Time

#  时间

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/time)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/time)**

The product owner wants us to expand the functionality of our command line application by helping a group of people play Texas-Holdem Poker.

产品所有者希望我们通过帮助一群人玩德州扑克来扩展命令行应用程序的功能。

## Just enough information on poker

## 关于扑克的足够信息

You won't need to know much about poker, only that at certain time intervals all the players need to be informed of a steadily increasing "blind" value.

你不需要对扑克有太多了解，只需要在特定的时间间隔内通知所有玩家一个稳步增加的“盲”值。

Our application will help keep track of when the blind should go up, and how much it should be.

我们的应用程序将帮助跟踪盲人应该在何时以及应该增加多少。

- When it starts it asks how many players are playing. This determines the amount of time there is before the "blind" bet goes up.
   - There is a base amount of time of 5 minutes.
   - For every player, 1 minute is added.
   - e.g 6 players equals 11 minutes for the blind.
- After the blind time expires the game should alert the players the new amount the blind bet is.
- The blind starts at 100 chips, then 200, 400, 600, 1000, 2000 and continue to double until the game ends (our previous functionality of "Ruth wins" should still finish the game)

- 当它开始时，它会询问有多少玩家在玩。这决定了“盲注”增加之前的时间量。
  - 基本时间为 5 分钟。
  - 对于每个玩家，增加 1 分钟。
  - 例如 6 名玩家等于盲注的 11 分钟。
- 盲注时间结束后，游戏应提醒玩家盲注的新金额。
- 盲注从 100 筹码开始，然后是 200、400、600、1000、2000 并继续翻倍直到游戏结束（我们之前的“Ruth wins”功能应该仍然可以完成游戏）

## Reminder of the code

## 代码提醒

In the previous chapter we made our start to the command line application which already accepts a command of `{name} wins`. Here is what the current `CLI` code looks like, but be sure to familiarise yourself with the other code too before starting.

在上一章中，我们开始了命令行应用程序，它已经接受了“{name} wins”命令。这是当前`CLI` 代码的样子，但在开始之前一定要熟悉其他代码。

```go
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

### `time.AfterFunc`

###`time.AfterFunc`

We want to be able to schedule our program to print the blind bet values at certain durations dependant on the number of players.

我们希望能够安排我们的程序以根据玩家数量在特定持续时间打印盲注值。

To limit the scope of what we need to do, we'll forget about the number of players part for now and just assume there are 5 players so we'll test that _every 10 minutes the new value of the blind bet is printed_.

为了限制我们需要做的事情的范围，我们暂时忘记了玩家数量，假设有 5 个玩家，所以我们将测试_每 10 分钟打印一次盲注的新值_。

As usual the standard library has us covered with [`func AfterFunc(d Duration, f func()) *Timer`](https://golang.org/pkg/time/#AfterFunc)

像往常一样，标准库为我们提供了 [`func AfterFunc(d Duration, f func()) *Timer`](https://golang.org/pkg/time/#AfterFunc)

> `AfterFunc` waits for the duration to elapse and then calls f in its own goroutine. It returns a `Timer` that can be used to cancel the call using its Stop method.

> `AfterFunc` 等待持续时间过去，然后在自己的 goroutine 中调用 f。它返回一个`Timer`，可用于使用其停止方法取消调用。

### [`time.Duration`](https://golang.org/pkg/time/#Duration)

### [`time.Duration`](https://golang.org/pkg/time/#Duration)

> A Duration represents the elapsed time between two instants as an int64 nanosecond count.

> Duration 表示两个时刻之间经过的时间，以 int64 纳秒计数表示。

The time library has a number of constants to let you multiply those nanoseconds so they're a bit more readable for the kind of scenarios we'll be doing

时间库有许多常数可以让您乘以这些纳秒，因此对于我们将要执行的场景，它们更具可读性

```go
5 * time.Second
```

When we call `PlayPoker` we'll schedule all of our blind alerts.

当我们调用“PlayPoker”时，我们将安排我们所有的盲提示。

Testing this may be a little tricky though. We'll want to verify that each time period is scheduled with the correct blind amount but if you look at the signature of `time.AfterFunc` its second argument is the function it will run. You cannot compare functions in Go so we'd be unable to test what function has been sent in. So we'll need to write some kind of wrapper around `time.AfterFunc` which will take the time to run and the amount to print so we can spy on that.

不过，测试这可能有点棘手。我们将要验证每个时间段是否以正确的盲量进行了安排，但是如果您查看 `time.AfterFunc` 的签名，它的第二个参数是它将运行的函数。您无法在 Go 中比较函数，因此我们无法测试已发送的函数。因此，我们需要围绕 `time.AfterFunc` 编写某种包装器，这将花费时间运行和打印数量所以我们可以监视它。

## Write the test first

## 先写测试

Add a new test to our suite

向我们的套件添加一个新测试

```go
t.Run("it schedules printing of blind values", func(t *testing.T) {
    in := strings.NewReader("Chris wins\n")
    playerStore := &poker.StubPlayerStore{}
    blindAlerter := &SpyBlindAlerter{}

    cli := poker.NewCLI(playerStore, in, blindAlerter)
    cli.PlayPoker()

    if len(blindAlerter.alerts) != 1 {
        t.Fatal("expected a blind alert to be scheduled")
    }
})
```

You'll notice we've made a `SpyBlindAlerter` which we are trying to inject into our `CLI` and then checking that after we call `PlayPoker` that an alert is scheduled.

你会注意到我们已经制作了一个 `SpyBlindAlerter`，我们试图将它注入到我们的 `CLI` 中，然后在我们调用 `PlayPoker` 之后检查是否安排了警报。

(Remember we are just going for the simplest scenario first and then we'll iterate.)

（请记住，我们只是先处理最简单的场景，然后再进行迭代。）

Here's the definition of `SpyBlindAlerter`

这是`SpyBlindAlerter`的定义

```go
type SpyBlindAlerter struct {
    alerts []struct{
        scheduledAt time.Duration
        amount int
    }
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
    s.alerts = append(s.alerts, struct {
        scheduledAt time.Duration
        amount int
    }{duration,  amount})
}

```

## Try to run the test

## 尝试运行测试

```
./CLI_test.go:32:27: too many arguments in call to poker.NewCLI
    have (*poker.StubPlayerStore, *strings.Reader, *SpyBlindAlerter)
    want (poker.PlayerStore, io.Reader)
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We have added a new argument and the compiler is complaining. _Strictly speaking_ the minimal amount of code is to make `NewCLI` accept a `*SpyBlindAlerter` but let's cheat a little and just define the dependency as an interface.

我们添加了一个新参数，编译器正在抱怨。 _严格来说_最少量的代码是让`NewCLI`接受`*SpyBlindAlerter`，但让我们作弊一下，将依赖项定义为接口。

```go
type BlindAlerter interface {
    ScheduleAlertAt(duration time.Duration, amount int)
}
```

And then add it to the constructor

然后将其添加到构造函数中

```go
func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI
```

Your other tests will now fail as they don't have a `BlindAlerter` passed in to `NewCLI`.

您的其他测试现在将失败，因为它们没有将 `BlindAlerter` 传递给 `NewCLI`。

Spying on BlindAlerter is not relevant for the other tests so in the test file add

监视 BlindAlerter 与其他测试无关，因此在测试文件中添加

```go
var dummySpyAlerter = &SpyBlindAlerter{}
```

Then use that in the other tests to fix the compilation problems. By labelling it as a "dummy" it is clear to the reader of the test that it is not important.

然后在其他测试中使用它来修复编译问题。通过将其标记为“虚拟”，测试的读者很清楚它并不重要。

[> Dummy objects are passed around but never actually used. Usually they are just used to fill parameter lists.](https://martinfowler.com/articles/mocksArentStubs.html)

[> 虚拟对象被传递但从未真正使用过。通常它们只是用来填充参数列表。](https://martinfowler.com/articles/mocksArentStubs.html)

The tests should now compile and our new test fails.

测试现在应该可以编译，我们的新测试失败了。

```
=== RUN   TestCLI
=== RUN   TestCLI/it_schedules_printing_of_blind_values
--- FAIL: TestCLI (0.00s)
    --- FAIL: TestCLI/it_schedules_printing_of_blind_values (0.00s)
        CLI_test.go:38: expected a blind alert to be scheduled
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We'll need to add the `BlindAlerter` as a field on our `CLI` so we can reference it in our `PlayPoker` method.

我们需要在我们的 CLI 上添加 BlindAlerter 作为一个字段，以便我们可以在我们的 PlayPoker 方法中引用它。

```go
type CLI struct {
    playerStore PlayerStore
    in          *bufio.Scanner
    alerter     BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
    return &CLI{
        playerStore: store,
        in:          bufio.NewScanner(in),
        alerter:     alerter,
    }
}
```

To make the test pass, we can call our `BlindAlerter` with anything we like

为了让测试通过，我们可以用我们喜欢的任何东西调用我们的“BlindAlerter”

```go
func (cli *CLI) PlayPoker() {
    cli.alerter.ScheduleAlertAt(5 * time.Second, 100)
    userInput := cli.readLine()
    cli.playerStore.RecordWin(extractWinner(userInput))
}
```

Next we'll want to check it schedules all the alerts we'd hope for, for 5 players

接下来我们要检查它是否为 5 个玩家安排了我们希望的所有警报

## Write the test first

## 先写测试

```go
     t.Run("it schedules printing of blind values", func(t *testing.T) {
        in := strings.NewReader("Chris wins\n")
        playerStore := &poker.StubPlayerStore{}
        blindAlerter := &SpyBlindAlerter{}

        cli := poker.NewCLI(playerStore, in, blindAlerter)
        cli.PlayPoker()

        cases := []struct{
            expectedScheduleTime time.Duration
            expectedAmount       int
        } {
            {0 * time.Second, 100},
            {10 * time.Minute, 200},
            {20 * time.Minute, 300},
            {30 * time.Minute, 400},
            {40 * time.Minute, 500},
            {50 * time.Minute, 600},
            {60 * time.Minute, 800},
            {70 * time.Minute, 1000},
            {80 * time.Minute, 2000},
            {90 * time.Minute, 4000},
            {100 * time.Minute, 8000},
        }

        for i, c := range cases {
            t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduleTime), func(t *testing.T) {

                if len(blindAlerter.alerts) <= i {
                    t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
                }

                alert := blindAlerter.alerts[i]

                amountGot := alert.amount
                if amountGot != c.expectedAmount {
                    t.Errorf("got amount %d, want %d", amountGot, c.expectedAmount)
                }

                gotScheduledTime := alert.scheduledAt
                if gotScheduledTime != c.expectedScheduleTime {
                    t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.expectedScheduleTime)
                }
            })
        }
    })
```

Table-based test works nicely here and clearly illustrate what our requirements are. We run through the table and check the `SpyBlindAlerter` to see if the alert has been scheduled with the correct values.

基于表格的测试在这里工作得很好，并清楚地说明了我们的要求是什么。我们遍历表并检查“SpyBlindAlerter”以查看警报是否已使用正确的值进行调度。

## Try to run the test

## 尝试运行测试

You should have a lot of failures looking like this

你应该有很多像这样的失败

```go
=== RUN   TestCLI
--- FAIL: TestCLI (0.00s)
=== RUN   TestCLI/it_schedules_printing_of_blind_values
    --- FAIL: TestCLI/it_schedules_printing_of_blind_values (0.00s)
=== RUN   TestCLI/it_schedules_printing_of_blind_values/100_scheduled_for_0s
        --- FAIL: TestCLI/it_schedules_printing_of_blind_values/100_scheduled_for_0s (0.00s)
            CLI_test.go:71: got scheduled time of 5s, want 0s
=== RUN   TestCLI/it_schedules_printing_of_blind_values/200_scheduled_for_10m0s
        --- FAIL: TestCLI/it_schedules_printing_of_blind_values/200_scheduled_for_10m0s (0.00s)
            CLI_test.go:59: alert 1 was not scheduled [{5000000000 100}]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (cli *CLI) PlayPoker() {

    blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
    blindTime := 0 * time.Second
    for _, blind := range blinds {
        cli.alerter.ScheduleAlertAt(blindTime, blind)
        blindTime = blindTime + 10 * time.Minute
    }

    userInput := cli.readLine()
    cli.playerStore.RecordWin(extractWinner(userInput))
}
```

It's not a lot more complicated than what we already had. We're just now iterating over an array of `blinds` and calling the scheduler on an increasing `blindTime`

它并不比我们已有的复杂得多。我们现在正在迭代一个 `blinds` 数组，并在递增的 `blindTime` 上调用调度程序

## Refactor

## 重构

We can encapsulate our scheduled alerts into a method just to make `PlayPoker` read a little clearer.

我们可以将我们的预定警报封装到一个方法中，只是为了让 `PlayPoker` 读起来更清晰一些。

```go
func (cli *CLI) PlayPoker() {
    cli.scheduleBlindAlerts()
    userInput := cli.readLine()
    cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts() {
    blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
    blindTime := 0 * time.Second
    for _, blind := range blinds {
        cli.alerter.ScheduleAlertAt(blindTime, blind)
        blindTime = blindTime + 10*time.Minute
    }
}
```

Finally our tests are looking a little clunky. We have two anonymous structs representing the same thing, a `ScheduledAlert`. Let's refactor that into a new type and then make some helpers to compare them.

最后，我们的测试看起来有点笨拙。我们有两个匿名结构代表同一件事，一个“ScheduledAlert”。让我们将其重构为一个新类型，然后制作一些助手来比较它们。

```go
type scheduledAlert struct {
    at time.Duration
    amount int
}

func (s scheduledAlert) String() string {
    return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
    alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
    s.alerts = append(s.alerts, scheduledAlert{at, amount})
}
```

We've added a `String()` method to our type so it prints nicely if the test fails

我们在我们的类型中添加了一个 `String()` 方法，因此如果测试失败，它会很好地打印

Update our test to use our new type

更新我们的测试以使用我们的新类型

```go
t.Run("it schedules printing of blind values", func(t *testing.T) {
    in := strings.NewReader("Chris wins\n")
    playerStore := &poker.StubPlayerStore{}
    blindAlerter := &SpyBlindAlerter{}

    cli := poker.NewCLI(playerStore, in, blindAlerter)
    cli.PlayPoker()

    cases := []scheduledAlert {
        {0 * time.Second, 100},
        {10 * time.Minute, 200},
        {20 * time.Minute, 300},
        {30 * time.Minute, 400},
        {40 * time.Minute, 500},
        {50 * time.Minute, 600},
        {60 * time.Minute, 800},
        {70 * time.Minute, 1000},
        {80 * time.Minute, 2000},
        {90 * time.Minute, 4000},
        {100 * time.Minute, 8000},
    }

    for i, want := range cases {
        t.Run(fmt.Sprint(want), func(t *testing.T) {

            if len(blindAlerter.alerts) <= i {
                t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
            }

            got := blindAlerter.alerts[i]
            assertScheduledAlert(t, got, want)
        })
    }
})
```

Implement `assertScheduledAlert` yourself.

自己实现 `assertScheduledAlert`。

We've spent a fair amount of time here writing tests and have been somewhat naughty not integrating with our application. Let's address that before we pile on any more requirements.

我们在这里花费了相当多的时间来编写测试，并且在不与我们的应用程序集成方面有些顽皮。在我们提出更多要求之前，让我们解决这个问题。

Try running the app and it won't compile, complaining about not enough args to `NewCLI`.

尝试运行该应用程序，但它无法编译，并抱怨`NewCLI` 没有足够的参数。

Let's create an implementation of `BlindAlerter` that we can use in our application.

让我们创建一个可以在我们的应用程序中使用的 `BlindAlerter` 的实现。

Create `BlindAlerter.go` and move our `BlindAlerter` interface and add the new things below

创建 `BlindAlerter.go` 并移动我们的 `BlindAlerter` 界面并在下面添加新内容

```go
package poker

import (
    "time"
    "fmt"
    "os"
)

type BlindAlerter interface {
    ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
    a(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
    time.AfterFunc(duration, func() {
        fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
    })
}
```

Remember that any _type_ can implement an interface, not just `structs`. If you are making a library that exposes an interface with one function defined it is a common idiom to also expose a `MyInterfaceFunc` type.

请记住，任何 _type_ 都可以实现接口，而不仅仅是 `structs`。如果你正在制作一个公开一个定义了一个函数的接口的库，那么公开一个`MyInterfaceFunc`类型是一种常见的习惯用法。

This type will be a `func` which will also implement your interface. That way users of your interface have the option to implement your interface with just a function; rather than having to create an empty `struct` type.

这种类型将是一个`func`，它也将实现你的接口。这样你的界面的用户就可以选择只用一个函数来实现你的界面；而不是必须创建一个空的 `struct` 类型。

We then create the function `StdOutAlerter` which has the same signature as the function and just use `time.AfterFunc` to schedule it to print to `os.Stdout`.

然后我们创建函数`StdOutAlerter`，它具有与函数相同的签名，并且只使用`time.AfterFunc` 来安排它打印到`os.Stdout`。

Update `main` where we create `NewCLI` to see this in action

更新 `main`，我们在其中创建 `NewCLI` 以查看它的实际效果

```go
poker.NewCLI(store, os.Stdin, poker.BlindAlerterFunc(poker.StdOutAlerter)).PlayPoker()
```

Before running you might want to change the `blindTime` increment in `CLI` to be 10 seconds rather than 10 minutes just so you can see it in action. 

在运行之前，您可能希望将 `CLI` 中的 `blindTime` 增量更改为 10 秒而不是 10 分钟，以便您可以看到它的运行情况。

You should see it print the blind values as we'd expect every 10 seconds. Notice how you can still type `Shaun wins` into the CLI and it will stop the program how we'd expect.

您应该看到它每 10 秒打印一次盲值，正如我们所期望的那样。请注意，您仍然可以在 CLI 中输入“Shaun wins”，它会按照我们的预期停止程序。

The game won't always be played with 5 people so we need to prompt the user to enter a number of players before the game starts.

游戏并不总是 5 人玩，所以我们需要在游戏开始前提示用户输入玩家人数。

## Write the test first

## 先写测试

To check we are prompting for the number of players we'll want to record what is written to StdOut. We've done this a few times now, we know that `os.Stdout` is an `io.Writer` so we can check what is written if we use dependency injection to pass in a `bytes.Buffer` in our test and see what our code will write.

为了检查，我们提示输入我们想要记录写入 StdOut 的内容的玩家数量。我们已经这样做了几次，我们知道 `os.Stdout` 是一个 `io.Writer`，所以如果我们在测试中使用依赖注入传递 `bytes.Buffer` 和看看我们的代码会写什么。

We don't care about our other collaborators in this test just yet so we've made some dummies in our test file.

我们暂时不关心这个测试中的其他合作者，所以我们在我们的测试文件中做了一些假人。

We should be a little wary that we now have 4 dependencies for `CLI`, that feels like maybe it is starting to have too many responsibilities. Let's live with it for now and see if a refactoring emerges as we add this new functionality.

我们应该有点警惕，我们现在有 4 个`CLI` 依赖项，感觉它可能开始承担太多的责任。让我们暂时接受它，看看当我们添加这个新功能时是否会出现重构。

```go
var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}
```

Here is our new test

这是我们的新测试

```go
t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
    stdout := &bytes.Buffer{}
    cli := poker.NewCLI(dummyPlayerStore, dummyStdIn, stdout, dummyBlindAlerter)
    cli.PlayPoker()

    got := stdout.String()
    want := "Please enter the number of players: "

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
})
```

We pass in what will be `os.Stdout` in `main` and see what is written.

我们在 `main` 中传入 `os.Stdout` 的内容，看看写了什么。

## Try to run the test

## 尝试运行测试

```
./CLI_test.go:38:27: too many arguments in call to poker.NewCLI
    have (*poker.StubPlayerStore, *bytes.Buffer, *bytes.Buffer, *SpyBlindAlerter)
    want (poker.PlayerStore, io.Reader, poker.BlindAlerter)
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We have a new dependency so we'll have to update `NewCLI`

我们有一个新的依赖项，所以我们必须更新 `NewCLI`

```go
func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI
```

Now the _other_ tests will fail to compile because they don't have an `io.Writer` being passed into `NewCLI`.

现在 _other_ 测试将无法编译，因为它们没有将 `io.Writer` 传递给 `NewCLI`。

Add `dummyStdout` for the other tests.

为其他测试添加 `dummyStdout`。

The new test should fail like so

新测试应该像这样失败

```
=== RUN   TestCLI
--- FAIL: TestCLI (0.00s)
=== RUN   TestCLI/it_prompts_the_user_to_enter_the_number_of_players
    --- FAIL: TestCLI/it_prompts_the_user_to_enter_the_number_of_players (0.00s)
        CLI_test.go:46: got '', want 'Please enter the number of players: '
FAIL
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We need to add our new dependency to our `CLI` so we can reference it in `PlayPoker`

我们需要将我们的新依赖添加到我们的 `CLI` 中，以便我们可以在 `PlayPoker` 中引用它

```go
type CLI struct {
    playerStore PlayerStore
    in          *bufio.Scanner
    out         io.Writer
    alerter     BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
    return &CLI{
        playerStore: store,
        in:          bufio.NewScanner(in),
        out:         out,
        alerter:     alerter,
    }
}
```

Then finally we can write our prompt at the start of the game

然后最后我们可以在游戏开始时编写我们的提示

```go
func (cli *CLI) PlayPoker() {
    fmt.Fprint(cli.out, "Please enter the number of players: ")
    cli.scheduleBlindAlerts()
    userInput := cli.readLine()
    cli.playerStore.RecordWin(extractWinner(userInput))
}
```

## Refactor

## 重构

We have a duplicate string for the prompt which we should extract into a constant

我们有一个重复的提示字符串，我们应该将其提取为一个常量

```go
const PlayerPrompt = "Please enter the number of players: "
```

Use this in both the test code and `CLI`.

在测试代码和 `CLI` 中都使用它。

Now we need to send in a number and extract it out. The only way we'll know if it has had the desired effect is by seeing what blind alerts were scheduled.

现在我们需要发送一个数字并将其提取出来。我们知道它是否达到预期效果的唯一方法是查看安排了哪些盲目警报。

## Write the test first

## 先写测试

```go
t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
    stdout := &bytes.Buffer{}
    in := strings.NewReader("7\n")
    blindAlerter := &SpyBlindAlerter{}

    cli := poker.NewCLI(dummyPlayerStore, in, stdout, blindAlerter)
    cli.PlayPoker()

    got := stdout.String()
    want := poker.PlayerPrompt

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }

    cases := []scheduledAlert{
        {0 * time.Second, 100},
        {12 * time.Minute, 200},
        {24 * time.Minute, 300},
        {36 * time.Minute, 400},
    }

    for i, want := range cases {
        t.Run(fmt.Sprint(want), func(t *testing.T) {

            if len(blindAlerter.alerts) <= i {
                t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
            }

            got := blindAlerter.alerts[i]
            assertScheduledAlert(t, got, want)
        })
    }
})
```

Ouch! A lot of changes.

哎哟!很多变化。

- We remove our dummy for StdIn and instead send in a mocked version representing our user entering 7
- We also remove our dummy on the blind alerter so we can see that the number of players has had an effect on the scheduling
- We test what alerts are scheduled

- 我们删除了 StdIn 的虚拟对象，而是发送一个模拟版本，代表我们的用户输入 7
- 我们还删除了盲人警报器上的假人，以便我们可以看到玩家数量对日程安排产生了影响
- 我们测试安排了哪些警报

## Try to run the test

## 尝试运行测试

The test should still compile and fail reporting that the scheduled times are wrong because we've hard-coded for the game to be based on having 5 players

测试仍应编译并失败，报告计划时间错误，因为我们已将游戏硬编码为基于 5 名玩家

```
=== RUN   TestCLI
--- FAIL: TestCLI (0.00s)
=== RUN   TestCLI/it_prompts_the_user_to_enter_the_number_of_players
    --- FAIL: TestCLI/it_prompts_the_user_to_enter_the_number_of_players (0.00s)
=== RUN   TestCLI/it_prompts_the_user_to_enter_the_number_of_players/100_chips_at_0s
        --- PASS: TestCLI/it_prompts_the_user_to_enter_the_number_of_players/100_chips_at_0s (0.00s)
=== RUN   TestCLI/it_prompts_the_user_to_enter_the_number_of_players/200_chips_at_12m0s
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Remember, we are free to commit whatever sins we need to make this work. Once we have working software we can then work on refactoring the mess we're about to make!

请记住，我们可以自由地犯下使这项工作所需的任何罪过。一旦我们有了可以工作的软件，我们就可以着手重构我们将要制造的混乱！

```go
func (cli *CLI) PlayPoker() {
    fmt.Fprint(cli.out, PlayerPrompt)

    numberOfPlayers, _ := strconv.Atoi(cli.readLine())

    cli.scheduleBlindAlerts(numberOfPlayers)

    userInput := cli.readLine()
    cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
    blindIncrement := time.Duration(5 + numberOfPlayers) * time.Minute

    blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
    blindTime := 0 * time.Second
    for _, blind := range blinds {
        cli.alerter.ScheduleAlertAt(blindTime, blind)
        blindTime = blindTime + blindIncrement
    }
}
```

- We read in the `numberOfPlayersInput` into a string
- We use `cli.readLine()` to get the input from the user and then call `Atoi` to convert it into an integer - ignoring any error scenarios. We'll need to write a test for that scenario later.
- From here we change `scheduleBlindAlerts` to accept a number of players. We then calculate a `blindIncrement` time to use to add to `blindTime` as we iterate over the blind amounts

- 我们将 `numberOfPlayersInput` 读入一个字符串
- 我们使用 `cli.readLine()` 来获取用户的输入，然后调用 `Atoi` 将其转换为整数 - 忽略任何错误情况。稍后我们需要为该场景编写一个测试。
- 从这里开始，我们更改了 `scheduleBlindAlerts` 以接受许多玩家。然后我们计算一个 `blindIncrement` 时间，用于在迭代盲量时添加到 `blindTime`

While our new test has been fixed, a lot of others have failed because now our system only works if the game starts with a user entering a number. You'll need to fix the tests by changing the user inputs so that a number followed by a newline is added (this is highlighting yet more flaws in our approach right now).

虽然我们的新测试已经修复，但很多其他测试都失败了，因为现在我们的系统只有在游戏开始时用户输入数字时才能工作。您需要通过更改用户输入来修复测试，以便添加后跟换行符的数字（这突出了我们现在方法中的更多缺陷）。

## Refactor

## 重构

This all feels a bit horrible right? Let's **listen to our tests**.

这一切感觉有点可怕对吗？让我们**听听我们的测试**。

- In order to test that we are scheduling some alerts we set up 4 different dependencies. Whenever you have a lot of dependencies for a _thing_ in your system, it implies it's doing too much. Visually we can see it in how cluttered our test is.
- To me it feels like **we need to make a cleaner abstraction between reading user input and the business logic we want to do**
- A better test would be _given this user input, do we call a new type `Game` with the correct number of players_.
- We would then extract the testing of the scheduling into the tests for our new `Game`.

- 为了测试我们是否正在安排一些警报，我们设置了 4 个不同的依赖项。每当您的系统中的某物有很多依赖项时，就意味着它做得太多了。从视觉上我们可以看出我们的测试有多混乱。
- 对我来说感觉就像**我们需要在读取用户输入和我们想要做的业务逻辑之间做一个更清晰的抽象**
- 一个更好的测试是_给定这个用户输入，我们是否用正确的玩家数量调用一个新类型的“游戏”_。
- 然后，我们会将调度测试提取到新“游戏”的测试中。

We can refactor toward our `Game` first and our test should continue to pass. Once we've made the structural changes we want we can think about how we can refactor the tests to reflect our new separation of concerns

我们可以首先重构我们的“游戏”，我们的测试应该继续通过。一旦我们进行了我们想要的结构更改，我们就可以考虑如何重构测试以反映我们新的关注点分离

Remember when making changes in refactoring try to keep them as small as possible and keep re-running the tests.

请记住，在重构中进行更改时，请尝试使它们尽可能小并不断重新运行测试。

Try it yourself first. Think about the boundaries of what a `Game` would offer and what our `CLI` should be doing.

先自己试试。想想“游戏”将提供什么以及我们的“CLI”应该做什么的界限。

For now **don't** change the external interface of `NewCLI` as we don't want to change the test code and the client code at the same time as that is too much to juggle and we could end up breaking things .

现在**不要**更改`NewCLI`的外部接口，因为我们不想同时更改测试代码和客户端代码，因为这太复杂了，我们最终可能会破坏事情.

This is what I came up with:

这就是我想出的：

```go
// game.go
type Game struct {
    alerter BlindAlerter
    store   PlayerStore
}

func (p *Game) Start(numberOfPlayers int) {
    blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

    blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
    blindTime := 0 * time.Second
    for _, blind := range blinds {
        p.alerter.ScheduleAlertAt(blindTime, blind)
        blindTime = blindTime + blindIncrement
    }
}

func (p *Game) Finish(winner string) {
    p.store.RecordWin(winner)
}

// cli.go
type CLI struct {
    in          *bufio.Scanner
    out         io.Writer
    game        *Game
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
    return &CLI{
        in:  bufio.NewScanner(in),
        out: out,
        game: &Game{
            alerter: alerter,
            store:   store,
        },
    }
}

const PlayerPrompt = "Please enter the number of players: "

func (cli *CLI) PlayPoker() {
    fmt.Fprint(cli.out, PlayerPrompt)

    numberOfPlayersInput := cli.readLine()
    numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

    cli.game.Start(numberOfPlayers)

    winnerInput := cli.readLine()
    winner := extractWinner(winnerInput)

    cli.game.Finish(winner)
}

func extractWinner(userInput string) string {
    return strings.Replace(userInput, " wins\n", "", 1)
}

func (cli *CLI) readLine() string {
    cli.in.Scan()
    return cli.in.Text()
}
```

From a "domain" perspective:
- We want to `Start` a `Game`, indicating how many people are playing
- We want to `Finish` a `Game`, declaring the winner

从“域”的角度来看：
- 我们想要`Start`一个`Game`，表明有多少人在玩
- 我们要“完成”一场“游戏”，宣布获胜者

The new `Game` type encapsulates this for us.

新的`Game` 类型为我们封装了这一点。

With this change we've passed `BlindAlerter` and `PlayerStore` to `Game` as it is now responsible for alerting and storing results.

通过此更改，我们将 `BlindAlerter` 和 `PlayerStore` 传递给了 `Game`，因为它现在负责警报和存储结果。

Our `CLI` is now just concerned with:

我们的 `CLI` 现在只关心：

- Constructing `Game` with its existing dependencies (which we'll refactor next)
- Interpreting user input as method invocations for `Game`

- 使用现有的依赖项构建`Game`（接下来我们将对其进行重构）
- 将用户输入解释为“游戏”的方法调用

We want to try to avoid doing "big" refactors which leave us in a state of failing tests for extended periods as that increases the chances of mistakes. (If you are working in a large/distributed team this is extra important)

我们希望尽量避免进行“大”重构，这会让我们长时间处于测试失败的状态，因为这会增加出错的机会。 （如果您在大型/分布式团队中工作，这一点尤为重要）

The first thing we'll do is refactor `Game` so that we inject it into `CLI`. We'll do the smallest changes in our tests to facilitate that and then we'll see how we can break up the tests into the themes of parsing user input and game management.

我们要做的第一件事是重构“Game”，以便将其注入到“CLI”中。我们将在我们的测试中进行最小的更改以促进这一点，然后我们将看到如何将测试分解为解析用户输入和游戏管理的主题。

All we need to do right now is change `NewCLI`

我们现在需要做的就是改变`NewCLI`

```go
func NewCLI(in io.Reader, out io.Writer, game *Game) *CLI {
    return &CLI{
        in:  bufio.NewScanner(in),
        out: out,
        game: game,
    }
}
```

This feels like an improvement already. We have less dependencies and _our dependency list is reflecting our overall design goal_ of CLI being concerned with input/output and delegating game specific actions to a `Game`.

这感觉已经是一种进步。我们有更少的依赖项，并且_我们的依赖项列表反映了我们的 CLI 总体设计目标_关注输入/输出并将游戏特定操作委托给“游戏”。

If you try and compile there are problems. You should be able to fix these problems yourself. Don't worry about making any mocks for `Game` right now, just initialise _real_ `Game`s just to get everything compiling and tests green.

如果您尝试编译，则会出现问题。您应该能够自己解决这些问题。现在不要担心为 `Game` 制作任何模拟，只需初始化 _real_ `Game`s 只是为了让所有内容编译和测试绿色。

To do this you'll need to make a constructor

为此，您需要创建一个构造函数

```go
func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
    return &Game{
        alerter:alerter,
        store:store,
    }
}
```

Here's an example of one of the setups for the tests being fixed

这是正在修复的测试设置之一的示例

```go
stdout := &bytes.Buffer{}
in := strings.NewReader("7\n")
blindAlerter := &SpyBlindAlerter{}
game := poker.NewGame(blindAlerter, dummyPlayerStore)

cli := poker.NewCLI(in, stdout, game)
cli.PlayPoker()
```

It shouldn't take much effort to fix the tests and be back to green again (that's the point!) but make sure you fix `main.go` too before the next stage.

修复测试并再次恢复绿色不应该花费太多精力（这就是重点！）但请确保在下一阶段之前也修复 `main.go`。

```go
// main.go
game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
cli := poker.NewCLI(os.Stdin, os.Stdout, game)
cli.PlayPoker()
```

Now that we have extracted out `Game` we should move our game specific assertions into tests separate from CLI.

现在我们已经提取了`Game`，我们应该将我们的游戏特定断言移到独立于 CLI 的测试中。

This is just an exercise in copying our `CLI` tests but with less dependencies

这只是复制我们的 CLI 测试的练习，但依赖较少

```go
func TestGame_Start(t *testing.T) {
    t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
        blindAlerter := &poker.SpyBlindAlerter{}
        game := poker.NewGame(blindAlerter, dummyPlayerStore)

        game.Start(5)

        cases := []poker.ScheduledAlert{
            {At: 0 * time.Second, Amount: 100},
            {At: 10 * time.Minute, Amount: 200},
            {At: 20 * time.Minute, Amount: 300},
            {At: 30 * time.Minute, Amount: 400},
            {At: 40 * time.Minute, Amount: 500},
            {At: 50 * time.Minute, Amount: 600},
            {At: 60 * time.Minute, Amount: 800},
            {At: 70 * time.Minute, Amount: 1000},
            {At: 80 * time.Minute, Amount: 2000},
            {At: 90 * time.Minute, Amount: 4000},
            {At: 100 * time.Minute, Amount: 8000},
        }

        checkSchedulingCases(cases, t, blindAlerter)
    })

    t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
        blindAlerter := &poker.SpyBlindAlerter{}
        game := poker.NewGame(blindAlerter, dummyPlayerStore)

        game.Start(7)

        cases := []poker.ScheduledAlert{
            {At: 0 * time.Second, Amount: 100},
            {At: 12 * time.Minute, Amount: 200},
            {At: 24 * time.Minute, Amount: 300},
            {At: 36 * time.Minute, Amount: 400},
        }

        checkSchedulingCases(cases, t, blindAlerter)
    })

}

func TestGame_Finish(t *testing.T) {
    store := &poker.StubPlayerStore{}
    game := poker.NewGame(dummyBlindAlerter, store)
    winner := "Ruth"

    game.Finish(winner)
    poker.AssertPlayerWin(t, store, winner)
}
```

The intent behind what happens when a game of poker starts is now much clearer.

扑克游戏开始时发生的事情背后的意图现在更加清晰。

Make sure to also move over the test for when the game ends.

确保还移动到游戏结束时的测试。

Once we are happy we have moved the tests over for game logic we can simplify our CLI tests so they reflect our intended responsibilities clearer

一旦我们对游戏逻辑的测试感到满意，我们就可以简化我们的 CLI 测试，以便它们更清楚地反映我们的预期职责

- Process user input and call `Game`'s methods when appropriate
- Send output
- Crucially it doesn't know about the actual workings of how games work 

- 处理用户输入并在适当的时候调用“游戏”的方法
- 发送输出
- 至关重要的是它不知道游戏的实际运作方式

To do this we'll have to make it so `CLI` no longer relies on a concrete `Game` type but instead accepts an interface with `Start(numberOfPlayers)` and `Finish(winner)`. We can then create a spy of that type and verify the correct calls are made.

为此，我们必须使 `CLI` 不再依赖于具体的 `Game` 类型，而是接受带有 `Start(numberOfPlayers)` 和 `Finish(winner)` 的接口。然后我们可以创建该类型的间谍并验证是否进行了正确的调用。

It's here we realise that naming is awkward sometimes. Rename `Game` to `TexasHoldem` (as that's the _kind_ of game we're playing) and the new interface will be called `Game`. This keeps faithful to the notion that our CLI is oblivious to the actual game we're playing and what happens when you `Start` and `Finish`.

正是在这里，我们意识到命名有时很尴尬。将`Game` 重命名为`TexasHoldem`（因为这是我们正在玩的_kind_ 游戏），新界面将被称为`Game`。这与我们的 CLI 不知道我们正在玩的实际游戏以及当您“开始”和“完成”时会发生什么的概念保持一致。

```go
type Game interface {
    Start(numberOfPlayers int)
    Finish(winner string)
}
```

Replace all references to `*Game` inside `CLI` and replace them with `Game` (our new interface). As always keep re-running tests to check everything is green while we are refactoring.

替换 `CLI` 中对 `*Game` 的所有引用，并将它们替换为 `Game`（我们的新界面）。在我们重构时，一如既往地继续重新运行测试以检查一切是否正常。

Now that we have decoupled `CLI` from `TexasHoldem` we can use spies to check that `Start` and `Finish` are called when we expect them to, with the correct arguments.

现在我们已经将 `CLI` 与 `TexasHoldem` 分离，我们可以使用 spies 来检查 `Start` 和 `Finish` 是否在我们期望的时候使用正确的参数被调用。

Create a spy that implements `Game`

创建一个实现`Game`的间谍

```go
type GameSpy struct {
    StartedWith  int
    FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
    g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
    g.FinishedWith = winner
}
```

Replace any `CLI` test which is testing any game specific logic with checks on how our `GameSpy` is called. This will then reflect the responsibilities of CLI in our tests clearly.

用检查我们的“GameSpy”如何调用来替换任何测试任何游戏特定逻辑的“CLI”测试。这将清楚地反映 CLI 在我们的测试中的职责。

Here is an example of one of the tests being fixed; try and do the rest yourself and check the source code if you get stuck.

以下是修复其中一项测试的示例；尝试自己完成其余的工作，如果遇到困难，请检查源代码。

```go
     t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
        stdout := &bytes.Buffer{}
        in := strings.NewReader("7\n")
        game := &GameSpy{}

        cli := poker.NewCLI(in, stdout, game)
        cli.PlayPoker()

        gotPrompt := stdout.String()
        wantPrompt := poker.PlayerPrompt

        if gotPrompt != wantPrompt {
            t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
        }

        if game.StartedWith != 7 {
            t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
        }
    })
```

Now that we have a clean separation of concerns, checking edge cases around IO in our `CLI` should be easier.

现在我们已经清楚地分离了关注点，在我们的 CLI 中检查 IO 的边缘情况应该更容易。

We need to address the scenario where a user puts a non numeric value when prompted for the number of players:

我们需要解决用户在提示输入玩家数量时输入非数值的情况：

Our code should not start the game and it should print a handy error to the user and then exit.

我们的代码不应该启动游戏，它应该向用户打印一个方便的错误然后退出。

## Write the test first

## 先写测试

We'll start by making sure the game doesn't start

我们将首先确保游戏没有开始

```go
t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
        stdout := &bytes.Buffer{}
        in := strings.NewReader("Pies\n")
        game := &GameSpy{}

        cli := poker.NewCLI(in, stdout, game)
        cli.PlayPoker()

        if game.StartCalled {
            t.Errorf("game should not have started")
        }
    })
```

You'll need to add to our `GameSpy` a field `StartCalled` which only gets set if `Start` is called

您需要向我们的 `GameSpy` 添加一个字段 `StartCalled`，只有在调用 `Start` 时才会设置该字段

## Try to run the test
```
=== RUN   TestCLI/it_prints_an_error_when_a_non_numeric_value_is_entered_and_does_not_start_the_game
    --- FAIL: TestCLI/it_prints_an_error_when_a_non_numeric_value_is_entered_and_does_not_start_the_game (0.00s)
        CLI_test.go:62: game should not have started
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Around where we call `Atoi` we just need to check for the error

在我们调用 `Atoi` 的地方，我们只需要检查错误

```go
numberOfPlayers, err := strconv.Atoi(cli.readLine())

if err != nil {
    return
}
```

Next we need to inform the user of what they did wrong so we'll assert on what is printed to `stdout`.

接下来，我们需要通知用户他们做错了什么，以便我们断言打印到 `stdout` 的内容。

## Write the test first

## 先写测试

We've asserted on what was printed to `stdout` before so we can copy that code for now

我们已经断言之前打印到 `stdout` 的内容，因此我们现在可以复制该代码

```go
gotPrompt := stdout.String()

wantPrompt := poker.PlayerPrompt + "you're so silly"

if gotPrompt != wantPrompt {
    t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
}
```

We are storing _everything_ that gets written to stdout so we still expect the `poker.PlayerPrompt`. We then just check an additional thing gets printed. We're not too bothered about the exact wording for now, we'll address it when we refactor.

我们正在存储被写入标准输出的_一切_，所以我们仍然期待`poker.PlayerPrompt`。然后我们只检查打印了一个额外的东西。我们现在不太关心确切的措辞，我们将在重构时解决它。

## Try to run the test

## 尝试运行测试

```
=== RUN   TestCLI/it_prints_an_error_when_a_non_numeric_value_is_entered_and_does_not_start_the_game
    --- FAIL: TestCLI/it_prints_an_error_when_a_non_numeric_value_is_entered_and_does_not_start_the_game (0.00s)
        CLI_test.go:70: got 'Please enter the number of players: ', want 'Please enter the number of players: you're so silly'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Change the error handling code

更改错误处理代码

```go
if err != nil {
    fmt.Fprint(cli.out, "you're so silly")
    return
}
```

## Refactor 

## 重构

Now refactor the message into a constant like `PlayerPrompt`

现在将消息重构为一个常量，如“PlayerPrompt”

```go
wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrMsg
```

and put in a more appropriate message

并输入更合适的信息

```go
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
```

Finally our testing around what has been sent to `stdout` is quite verbose, let's write an assert function to clean it up.

最后，我们对发送到 `stdout` 的内容的测试非常冗长，让我们编写一个断言函数来清理它。

```go
func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
    t.Helper()
    want := strings.Join(messages, "")
    got := stdout.String()
    if got != want {
        t.Errorf("got %q sent to stdout but expected %+v", got, messages)
    }
}
```

Using the vararg syntax (`...string`) is handy here because we need to assert on varying amounts of messages.

在这里使用 vararg 语法 (`...string`) 很方便，因为我们需要对不同数量的消息进行断言。

Use this helper in both of the tests where we assert on messages sent to the user.

在我们对发送给用户的消息进行断言的两个测试中都使用此帮助程序。

There are a number of tests that could be helped with some `assertX` functions so practice your refactoring by cleaning up our tests so they read nicely.

有许多测试可以帮助一些 `assertX` 函数，所以通过清理我们的测试来练习你的重构，这样它们就可以很好地阅读。

Take some time and think about the value of some of the tests we've driven out. Remember we don't want more tests than necessary, can you refactor/remove some of them _and still be confident it all works_ ?

花点时间思考一下我们推出的一些测试的价值。请记住，我们不想要过多的测试，您能否重构/删除其中的一些_并仍然确信一切正常_？

Here is what I came up with

这是我想出的

```go
func TestCLI(t *testing.T) {

    t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
        game := &GameSpy{}
        stdout := &bytes.Buffer{}

        in := userSends("3", "Chris wins")
        cli := poker.NewCLI(in, stdout, game)

        cli.PlayPoker()

        assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
        assertGameStartedWith(t, game, 3)
        assertFinishCalledWith(t, game, "Chris")
    })

    t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
        game := &GameSpy{}

        in := userSends("8", "Cleo wins")
        cli := poker.NewCLI(in, dummyStdOut, game)

        cli.PlayPoker()

        assertGameStartedWith(t, game, 8)
        assertFinishCalledWith(t, game, "Cleo")
    })

    t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
        game := &GameSpy{}

        stdout := &bytes.Buffer{}
        in := userSends("pies")

        cli := poker.NewCLI(in, stdout, game)
        cli.PlayPoker()

        assertGameNotStarted(t, game)
        assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
    })
}
```

The tests now reflect the main capabilities of CLI, it is able to read user input in terms of how many people are playing and who won and handles when a bad value is entered for number of players. By doing this it is clear to the reader what `CLI` does, but also what it doesn't do.

这些测试现在反映了 CLI 的主要功能，它能够读取用户输入，包括有多少人在玩游戏，以及当为玩家数量输入错误值时谁赢了和处理。通过这样做，读者可以清楚地了解`CLI` 的作用，以及它不做什么。

What happens if instead of putting `Ruth wins` the user puts in `Lloyd is a killer` ?

如果用户输入的不是“Ruth wins”而是“Lloyd is akiller”，会发生什么？

Finish this chapter by writing a test for this scenario and making it pass.

通过为此场景编写测试并使其通过来完成本章。

## Wrapping up

##  总结

### A quick project recap

### 快速项目回顾

For the past 5 chapters we have slowly TDD'd a fair amount of code

在过去的 5 章中，我们已经慢慢地用 TDD 处理了相当多的代码

- We have two applications, a command line application and a web server.
- Both these applications rely on a `PlayerStore` to record winners
- The web server can also display a league table of who is winning the most games
- The command line app helps players play a game of poker by tracking what the current blind value is.

- 我们有两个应用程序，一个命令行应用程序和一个 Web 服务器。
- 这两个应用程序都依赖“PlayerStore”来记录获胜者
- 网络服务器还可以显示谁赢得最多比赛的联赛表
- 命令行应用程序通过跟踪当前的盲注值来帮助玩家玩扑克游戏。

### time.Afterfunc

### time.Afterfunc

A very handy way of scheduling a function call after a specific duration. It is well worth investing time [looking at the documentation for `time`](https://golang.org/pkg/time/) as it has a lot of time saving functions and methods for you to work with.

在特定持续时间之后安排函数调用的一种非常方便的方法。值得花时间 [查看`time` 的文档](https://golang.org/pkg/time/)，因为它有很多节省时间的功能和方法供您使用。

Some of my favourites are

我最喜欢的一些是

- `time.After(duration)` returns a `chan Time` when the duration has expired. So if you wish to do something _after_ a specific time, this can help.
- `time.NewTicker(duration)` returns a `Ticker` which is similar to the above in that it returns a channel but this one "ticks" every duration, rather than just once. This is very handy if you want to execute some code every `N duration`.

- `time.After(duration)` 返回一个 `chan Time` 当持续时间到期时。因此，如果您希望在特定时间之后做某事，这会有所帮助。
- `time.NewTicker(duration)` 返回一个 `Ticker`，它与上面的类似，因为它返回一个频道，但这个频道在每个持续时间“滴答”，而不是一次。如果您想每“N 持续时间”执行一些代码，这将非常方便。

### More examples of good separation of concerns

### 更多关注点分离的例子

_Generally_ it is good practice to separate the responsibilities of dealing with user input and responses away from domain code. You see that here in our command line application and also our web server.

_通常_将处理用户输入和响应的职责与域代码分开是一种很好的做法。您可以在我们的命令行应用程序和 Web 服务器中看到这一点。

Our tests got messy. We had too many assertions (check this input, schedules these alerts, etc) and too many dependencies. We could visually see it was cluttered; it is **so important to listen to your tests**.

我们的测试变得一团糟。我们有太多的断言（检查这个输入，安排这些警报等）和太多的依赖。我们可以直观地看到它很杂乱； **聆听您的测试非常重要**。

- If your tests look messy try and refactor them. 

- 如果您的测试看起来很乱，请尝试重构它们。

- If you've done this and they're still a mess it is very likely pointing to a flaw in your design
- This is one of the real strengths of tests.

- 如果你已经这样做了，但它们仍然一团糟，很可能表明你的设计存在缺陷
- 这是测试的真正优势之一。

Even though the tests and the production code was a bit cluttered we could freely refactor backed by our tests.

即使测试和生产代码有点混乱，我们也可以在测试的支持下自由重构。

Remember when you get into these situations to always take small steps and re-run the tests after every change.

请记住，当您遇到这些情况时，请始终采取小步骤并在每次更改后重新运行测试。

It would've been dangerous to refactor both the test code _and_ the production code at the same time, so we first refactored the production code (in the current state we couldn't improve the tests much) without changing its interface so we could rely on our tests as much as we could while changing things. _Then_ we refactored the tests after the design improved.

同时重构测试代码_和_生产代码是很危险的，所以我们首先重构了生产代码（在当前状态下我们不能对测试进行太多改进）而不改变它的接口，所以我们可以依靠在改变事物的同时尽可能多地进行测试。 _然后_我们在设计改进后重构了测试。

After refactoring the dependency list reflected our design goal. This is another benefit of DI in that it often documents intent. When you rely on global variables responsibilities become very unclear.

重构依赖列表后反映了我们的设计目标。这是 DI 的另一个好处，因为它经常记录意图。当你依赖全局变量时，职责就变得很不清楚了。

## An example of a function implementing an interface

## 一个实现接口的函数的例子

When you define an interface with one method in it you might want to consider defining a `MyInterfaceFunc` type to complement it so users can implement your interface with just a function

当你定义一个包含一个方法的接口时，你可能需要考虑定义一个 `MyInterfaceFunc` 类型来补充它，这样用户就可以只用一个函数来实现你的接口

```go
type BlindAlerter interface {
    ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlerterFunc allows you to implement BlindAlerter with a function
type BlindAlerterFunc func(duration time.Duration, amount int)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
    a(duration, amount)
}
```

