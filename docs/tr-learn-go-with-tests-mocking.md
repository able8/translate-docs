# Mocking

# 模拟

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/mocking)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/mocking)**

You have been asked to write a program which counts down from 3, printing each number on a new line (with a 1 second pause) and when it reaches zero it will print "Go!" and exit.

你被要求编写一个从 3 开始倒计时的程序，在新行上打印每个数字（暂停 1 秒），当它达到零时，它会打印“Go！”并退出。

```
3
2
1
Go!
```

We'll tackle this by writing a function called `Countdown` which we will then put inside a `main` program so it looks something like this:

我们将通过编写一个名为 `Countdown` 的函数来解决这个问题，然后我们将把它放在一个 `main` 程序中，它看起来像这样：

```go
package main

func main() {
    Countdown()
}
```

While this is a pretty trivial program, to test it fully we will need as always to take an _iterative_, _test-driven_ approach.

虽然这是一个非常简单的程序，但为了对其进行全面测试，我们需要一如既往地采用 _iterative_、_test-driven_ 方法。

What do I mean by iterative? We make sure we take the smallest steps we can to have _useful software_.

我所说的迭代是什么意思？我们确保采取最小的步骤来拥有_有用的软件_。

We don't want to spend a long time with code that will theoretically work after some hacking because that's often how developers fall down rabbit holes. **It's an important skill to be able to slice up requirements as small as you can so you can have _working software_.**

我们不想在经过一些黑客攻击后理论上可以工作的代码上花费很长时间，因为这通常是开发人员掉入兔子洞的方式。 **这是一项重要的技能，能够尽可能小地分割需求，这样您就可以拥有_工作软件_。**

Here's how we can divide our work up and iterate on it:

以下是我们如何划分我们的工作并对其进行迭代：

- Print 3
- Print 3, 2, 1 and Go!
- Wait a second between each line

- 打印 3
- 打印 3、2、1 并开始！
- 在每行之间等待一秒钟

## Write the test first

## 先写测试

Our software needs to print to stdout and we saw how we could use DI to facilitate testing this in the DI section.

我们的软件需要打印到标准输出，我们在 DI 部分看到了如何使用 DI 来促进测试。

```go
func TestCountdown(t *testing.T) {
    buffer := &bytes.Buffer{}

    Countdown(buffer)

    got := buffer.String()
    want := "3"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

If anything like `buffer` is unfamiliar to you, re-read [the previous section](dependency-injection.md).

如果你对“buffer”之类的东西不熟悉，请重新阅读[上一节](dependency-injection.md)。

We know we want our `Countdown` function to write data somewhere and `io.Writer` is the de-facto way of capturing that as an interface in Go.

我们知道我们希望我们的 `Countdown` 函数在某处写入数据，而 `io.Writer` 是在 Go 中将数据捕获为接口的事实上的方法。

- In `main` we will send to `os.Stdout` so our users see the countdown printed to the terminal.
- In test we will send to `bytes.Buffer` so our tests can capture what data is being generated.

- 在 `main` 中，我们将发送到 `os.Stdout`，以便我们的用户看到打印到终端的倒计时。
- 在测试中，我们将发送到 `bytes.Buffer`，以便我们的测试可以捕获正在生成的数据。

## Try and run the test

## 尝试并运行测试

`./countdown_test.go:11:2: undefined: Countdown`



## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Define `Countdown`

定义`倒计时`

```go
func Countdown() {}
```

Try again

再试一次

```go
./countdown_test.go:11:11: too many arguments in call to Countdown
    have (*bytes.Buffer)
    want ()
```

The compiler is telling you what your function signature could be, so update it.

编译器告诉你你的函数签名可能是什么，所以更新它。

```go
func Countdown(out *bytes.Buffer) {}
```

`countdown_test.go:17: got '' want '3'`

`countdown_test.go:17: got '' want '3'`

Perfect!

完美的！

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func Countdown(out *bytes.Buffer) {
    fmt.Fprint(out, "3")
}
```

We're using `fmt.Fprint` which takes an `io.Writer` (like `*bytes.Buffer`) and sends a `string` to it. The test should pass.

我们正在使用`fmt.Fprint`，它接受一个`io.Writer`（比如`*bytes.Buffer`）并向它发送一个`string`。测试应该通过。

## Refactor

## 重构

We know that while `*bytes.Buffer` works, it would be better to use a general purpose interface instead.

我们知道虽然 `*bytes.Buffer` 可以工作，但最好使用通用接口。

```go
func Countdown(out io.Writer) {
    fmt.Fprint(out, "3")
}
```

Re-run the tests and they should be passing.

重新运行测试，它们应该会通过。

To complete matters, let's now wire up our function into a `main` so we have some working software to reassure ourselves we're making progress.

为了完成问题，让我们现在将我们的函数连接到一个 `main` 中，这样我们就有了一些可以工作的软件来确保我们正在取得进展。

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func Countdown(out io.Writer) {
    fmt.Fprint(out, "3")
}

func main() {
    Countdown(os.Stdout)
}
```

Try and run the program and be amazed at your handywork.

尝试并运行该程序，你会惊讶于你的手艺。

Yes this seems trivial but this approach is what I would recommend for any project. **Take a thin slice of functionality and make it work end-to-end, backed by tests.**

是的，这似乎微不足道，但这种方法是我建议用于任何项目的方法。 **采用一小部分功能，使其端到端工作，并以测试为后盾。**

Next we can make it print 2,1 and then "Go!".

接下来我们可以让它打印 2,1 然后“Go！”。

## Write the test first

## 先写测试

By investing in getting the overall plumbing working right, we can iterate on our solution safely and easily. We will no longer need to stop and re-run the program to be confident of it working as all the logic is tested.

通过投资使整体管道正常工作，我们可以安全、轻松地迭代我们的解决方案。我们将不再需要停止并重新运行程序来确保它在所有逻辑都经过测试后能够正常工作。

```go
func TestCountdown(t *testing.T) {
    buffer := &bytes.Buffer{}

    Countdown(buffer)

    got := buffer.String()
    want := `3
2
1
Go!`

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

The backtick syntax is another way of creating a `string` but lets you put things like newlines which is perfect for our test.

反引号语法是另一种创建 `string` 的方式，但可以让你放置诸如换行符之类的东西，这对我们的测试来说是完美的。

## Try and run the test

## 尝试并运行测试

```
countdown_test.go:21: got '3' want '3
        2
        1
        Go!'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func Countdown(out io.Writer) {
    for i := 3;i > 0;i-- {
        fmt.Fprintln(out, i)
    }
    fmt.Fprint(out, "Go!")
}
```

Use a `for` loop counting backwards with `i--` and use `fmt.Fprintln` to print to `out` with our number followed by a newline character. Finally use `fmt.Fprint` to send "Go!" aftward.

使用 `i--` 向后计数的 `for` 循环，并使用 `fmt.Fprintln` 打印到 `out`，我们的数字后跟一个换行符。最后使用`fmt.Fprint` 发送“Go!”向后。

## Refactor

## 重构

There's not much to refactor other than refactoring some magic values into named constants.

除了将一些魔法值重构为命名常量之外，没有太多需要重构的。

```go
const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
    for i := countdownStart;i > 0;i-- {
        fmt.Fprintln(out, i)
    }
    fmt.Fprint(out, finalWord)
}
```

If you run the program now, you should get the desired output but we don't have it as a dramatic countdown with the 1 second pauses.

如果您现在运行该程序，您应该会获得所需的输出，但我们没有将它作为带有 1 秒暂停的戏剧性倒计时。

Go lets you achieve this with `time.Sleep`. Try adding it in to our code.

Go 可以让你通过 `time.Sleep` 实现这一点。尝试将其添加到我们的代码中。

```go
func Countdown(out io.Writer) {
    for i := countdownStart;i > 0;i-- {
        time.Sleep(1 * time.Second)
        fmt.Fprintln(out, i)
    }

    time.Sleep(1 * time.Second)
    fmt.Fprint(out, finalWord)
}
```

If you run the program it works as we want it to.

如果您运行该程序，它会按我们的意愿工作。

## Mocking

## 模拟

The tests still pass and the software works as intended but we have some problems:
- Our tests take 4 seconds to run.
     - Every forward thinking post about software development emphasises the importance of quick feedback loops.
     - **Slow tests ruin developer productivity**.
     - Imagine if the requirements get more sophisticated warranting more tests. Are we happy with 4s added to the test run for every new test of `Countdown`?
- We have not tested an important property of our function.

测试仍然通过，软件按预期工作，但我们有一些问题：
- 我们的测试需要 4 秒才能运行。
    - 每一篇关于软件开发的前瞻性帖子都强调了快速反馈循环的重要性。
    - **缓慢的测试会破坏开发人员的生产力**。
    - 想象一下，如果要求变得更加复杂，需要进行更多的测试。我们对每次新的“倒计时”测试在测试运行中添加 4 秒感到满意吗？
- 我们还没有测试我们函数的一个重要属性。

We have a dependency on `Sleep`ing which we need to extract so we can then control it in our tests.

我们需要提取对 `Sleep`ing 的依赖，以便我们可以在测试中控制它。

If we can _mock_ `time.Sleep` we can use _dependency injection_ to use it instead of a "real" `time.Sleep` and then we can **spy on the calls** to make assertions on them.

如果我们可以 _mock_ `time.Sleep`，我们可以使用 _dependency injection_ 来使用它而不是“真实的”`time.Sleep`，然后我们可以**监视调用**以对它们进行断言。

## Write the test first

## 先写测试

Let's define our dependency as an interface. This lets us then use a _real_ Sleeper in `main` and a _spy sleeper_ in our tests. By using an interface our `Countdown` function is oblivious to this and adds some flexibility for the caller.

让我们将依赖项定义为接口。这让我们可以在 `main` 中使用 _real_ Sleeper，在我们的测试中使用 _spy sleeper_。通过使用接口，我们的 `Countdown` 函数忽略了这一点，并为调用者增加了一些灵活性。

```go
type Sleeper interface {
    Sleep()
}
```

I made a design decision that our `Countdown` function would not be responsible for how long the sleep is. This simplifies our code a little for now at least and means a user of our function can configure that sleepiness however they like.

我做了一个设计决定，我们的“倒计时”功能不会对睡眠时间负责。这至少暂时简化了我们的代码，这意味着我们函数的用户可以根据自己的喜好配置这种困倦。

Now we need to make a _mock_ of it for our tests to use.

现在我们需要制作一个 _mock_ 供我们的测试使用。

```go
type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}
```

_Spies_ are a kind of _mock_ which can record how a dependency is used. They can record the arguments sent in, how many times it has been called, etc. In our case, we're keeping track of how many times `Sleep()` is called so we can check it in our test.

_Spies_ 是一种 _mock_，可以记录依赖项的使用方式。他们可以记录发送的参数，调用了多少次，等等。在我们的例子中，我们跟踪调用了多少次 `Sleep()`，以便我们可以在我们的测试中检查它。

Update the tests to inject a dependency on our Spy and assert that the sleep has been called 4 times.

更新测试以注入对我们的 Spy 的依赖并断言睡眠已被调用 4 次。

```go
func TestCountdown(t *testing.T) {
    buffer := &bytes.Buffer{}
    spySleeper := &SpySleeper{}

    Countdown(buffer, spySleeper)

    got := buffer.String()
    want := `3
2
1
Go!`

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }

    if spySleeper.Calls != 4 {
        t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
    }
}
```

## Try and run the test

## 尝试并运行测试

```
too many arguments in call to Countdown
    have (*bytes.Buffer, *SpySleeper)
    want (io.Writer)
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We need to update `Countdown` to accept our `Sleeper`

我们需要更新 `Countdown` 以接受我们的 `Sleeper`

```go
func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart;i > 0;i-- {
        time.Sleep(1 * time.Second)
        fmt.Fprintln(out, i)
    }

    time.Sleep(1 * time.Second)
    fmt.Fprint(out, finalWord)
}
```

If you try again, your `main` will no longer compile for the same reason

如果你再试一次，你的 `main` 将不再因为同样的原因编译

```
./main.go:26:11: not enough arguments in call to Countdown
    have (*os.File)
    want (io.Writer, Sleeper)
```

Let's create a _real_ sleeper which implements the interface we need

让我们创建一个 _real_ sleeper 来实现我们需要的接口

```go
type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}
```

We can then use it in our real application like so

然后我们可以像这样在我们的实际应用程序中使用它

```go
func main() {
    sleeper := &DefaultSleeper{}
    Countdown(os.Stdout, sleeper)
}
```

## Write enough code to make it pass

## 编写足够的代码使其通过

The test is now compiling but not passing because we're still calling the `time.Sleep` rather than the injected in dependency. Let's fix that.

测试现在正在编译但没有通过，因为我们仍在调用 `time.Sleep` 而不是注入的依赖项。让我们解决这个问题。

```go
func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart;i > 0;i-- {
        sleeper.Sleep()
        fmt.Fprintln(out, i)
    }

    sleeper.Sleep()
    fmt.Fprint(out, finalWord)
}
```

The test should pass and no longer take 4 seconds.

测试应该通过并且不再需要 4 秒。

### Still some problems

### 还有一些问题

There's still another important property we haven't tested.

还有一个重要的属性我们还没有测试。

`Countdown` should sleep before each print, e.g:

`Countdown` 应该在每次打印之前休眠，例如：

- `Sleep`
- `Print N`
- `Sleep`
- `Print N-1`
- `Sleep`
- `Print Go!`
- etc

Our latest change only asserts that it has slept 4 times, but those sleeps could occur out of sequence.

我们最新的更改只断言它已经睡了 4 次，但这些睡眠可能会发生乱序。

When writing tests if you're not confident that your tests are giving you sufficient confidence, just break it! (make sure you have committed your changes to source control first though). Change the code to the following

在编写测试时，如果您不确定您的测试是否给您足够的信心，那就打破它吧！ （不过，请确保您已将更改提交到源代码管理）。将代码更改为以下内容

```go
func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart;i > 0;i-- {
        sleeper.Sleep()
    }

    for i := countdownStart;i > 0;i-- {
        fmt.Fprintln(out, i)
    }

    sleeper.Sleep()
    fmt.Fprint(out, finalWord)
}
```

If you run your tests they should still be passing even though the implementation is wrong.

如果你运行你的测试，即使实现是错误的，它们仍然应该通过。

Let's use spying again with a new test to check the order of operations is correct.

让我们再次使用间谍与新测试来检查操作顺序是否正确。

We have two different dependencies and we want to record all of their operations into one list. So we'll create _one spy for them both_.

我们有两个不同的依赖项，我们希望将它们的所有操作记录到一个列表中。所以我们将为他们创建_一个间谍_。

```go
type SpyCountdownOperations struct {
    Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

const write = "write"
const sleep = "sleep"
```

Our `SpyCountdownOperations` implements both `io.Writer` and `Sleeper`, recording every call into one slice. In this test we're only concerned about the order of operations, so just recording them as list of named operations is sufficient.

我们的 `SpyCountdownOperations` 实现了 `io.Writer` 和 `Sleeper`，将每个调用记录到一个切片中。在这个测试中，我们只关心操作的顺序，所以只需将它们记录为命名操作列表就足够了。

We can now add a sub-test into our test suite which verifies our sleeps and prints operate in the order we hope

我们现在可以在我们的测试套件中添加一个子测试，以验证我们的睡眠和打印按我们希望的顺序运行

```go
t.Run("sleep before every print", func(t *testing.T) {
    spySleepPrinter := &SpyCountdownOperations{}
    Countdown(spySleepPrinter, spySleepPrinter)

    want := []string{
        sleep,
        write,
        sleep,
        write,
        sleep,
        write,
        sleep,
        write,
    }

    if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
        t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
    }
})
```

This test should now fail. Revert `Countdown` back to how it was to fix the test.

此测试现在应该失败。将“倒计时”恢复到修复测试的方式。

We now have two tests spying on the `Sleeper` so we can now refactor our test so one is testing what is being printed and the other one is ensuring we're sleeping in between the prints. Finally we can delete our first spy as it's not used anymore.

我们现在有两个测试监视“Sleeper”，所以我们现在可以重构我们的测试，一个是测试正在打印的内容，另一个是确保我们在打印之间睡觉。最后我们可以删除我们的第一个间谍，因为它不再使用了。

```go
func TestCountdown(t *testing.T) {

    t.Run("prints 3 to Go!", func(t *testing.T) {
        buffer := &bytes.Buffer{}
        Countdown(buffer, &SpyCountdownOperations{})

        got := buffer.String()
        want := `3
2
1
Go!`

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("sleep before every print", func(t *testing.T) {
        spySleepPrinter := &SpyCountdownOperations{}
        Countdown(spySleepPrinter, spySleepPrinter)

        want := []string{
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
        }

        if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
            t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
        }
    })
}
```

We now have our function and its 2 important properties properly tested.

我们现在已经正确测试了我们的函数和它的两个重要属性。

## Extending Sleeper to be configurable

## 将 Sleeper 扩展为可配置

A nice feature would be for the `Sleeper` to be configurable. This means that we can adjust the sleep time in our main program.

一个很好的功能是“Sleeper”是可配置的。这意味着我们可以在主程序中调整睡眠时间。

### Write the test first

### 先写测试

Let's first create a new type for `ConfigurableSleeper` that accepts what we need for configuration and testing.

让我们首先为 `ConfigurableSleeper` 创建一个新类型，它接受我们需要的配置和测试。

```go
type ConfigurableSleeper struct {
    duration time.Duration
    sleep    func(time.Duration)
}
```

We are using `duration` to configure the time slept and `sleep` as a way to pass in a sleep function. The signature of `sleep` is the same as for `time.Sleep` allowing us to use `time.Sleep` in our real implementation and the following spy in our tests:

我们使用 `duration` 来配置睡眠时间和 `sleep` 作为传递睡眠函数的一种方式。 `sleep` 的签名与 `time.Sleep` 的签名相同，允许我们在实际实现中使用 `time.Sleep`，并在我们的测试中使用以下间谍：

```go
type SpyTime struct {
    durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
    s.durationSlept = duration
}
```

With our spy in place, we can create a new test for the configurable sleeper.

有了我们的间谍，我们可以为可配置的睡眠者创建一个新的测试。

```go
func TestConfigurableSleeper(t *testing.T) {
    sleepTime := 5 * time.Second

    spyTime := &SpyTime{}
    sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
    sleeper.Sleep()

    if spyTime.durationSlept != sleepTime {
        t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
    }
}
```

There should be nothing new in this test and it is setup very similar to the previous mock tests.

这个测试应该没有什么新东西，它的设置与之前的模拟测试非常相似。

### Try and run the test
```
sleeper.Sleep undefined (type ConfigurableSleeper has no field or method Sleep, but does have sleep)

```

You should see a very clear error message indicating that we do not have a `Sleep` method created on our `ConfigurableSleeper`.

您应该会看到一条非常明确的错误消息，表明我们没有在我们的 `ConfigurableSleeper` 上创建一个 `Sleep` 方法。

### Write the minimal amount of code for the test to run and check failing test output
```go
func (c *ConfigurableSleeper) Sleep() {
}
```

With our new `Sleep` function implemented we have a failing test.

随着我们新的“睡眠”功能的实现，我们有一个失败的测试。

```
countdown_test.go:56: should have slept for 5s but slept for 0s
```

### Write enough code to make it pass

### 编写足够的代码使其通过

All we need to do now is implement the `Sleep` function for `ConfigurableSleeper`.

我们现在需要做的就是为 `ConfigurableSleeper` 实现 `Sleep` 函数。

```go
func (c *ConfigurableSleeper) Sleep() {
    c.sleep(c.duration)
}
```

With this change all of the tests should be passing again and you might wonder why all the hassle as the main program didn't change at all. Hopefully it becomes clear after the following section.

有了这个改变，所有的测试都应该再次通过，你可能想知道为什么主程序的所有麻烦根本没有改变。希望在下一节之后会变得清楚。

### Cleanup and refactor

### 清理和重构

The last thing we need to do is to actually use our `ConfigurableSleeper` in the main function.

我们需要做的最后一件事是在主函数中实际使用我们的 `ConfigurableSleeper`。

```go
func main() {
    sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
    Countdown(os.Stdout, sleeper)
}
```

If we run the tests and the program manually, we can see that all the behavior remains the same.

如果我们手动运行测试和程序，我们可以看到所有行为都保持不变。

Since we are using the `ConfigurableSleeper`, it is now safe to delete the `DefaultSleeper` implementation. Wrapping up our program and having a more [generic](https://stackoverflow.com/questions/19291776/whats-the-difference-between-abstraction-and-generalization) Sleeper with arbitrary long countdowns.

由于我们使用的是 `ConfigurableSleeper`，现在可以安全地删除 `DefaultSleeper` 实现。结束我们的程序并有一个更多的 [generic](https://stackoverflow.com/questions/19291776/whats-the-difference-between-abstraction-and-generalization) 带有任意长倒计时的卧铺。

## But isn't mocking evil?

## 但不是在嘲笑邪恶吗？

You may have heard mocking is evil. Just like anything in software development it can be used for evil, just like [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself).

你可能听说过嘲笑是邪恶的。就像软件开发中的任何事情一样，它可以用于邪恶，就像 [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)。

People normally get in to a bad state when they don't _listen to their tests_ and are _not respecting the refactoring stage_.

当人们不_听他们的测试_并且_不尊重重构阶段_时，他们通常会进入一个糟糕的状态。

If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should _listen_ to that bad feeling and think about your code. Usually it is a sign of

如果你的模拟代码变得复杂，或者你不得不模拟很多东西来测试某些东西，你应该_听_这种不好的感觉并考虑你的代码。通常它是一个标志

- The thing you are testing is having to do too many things (because it has too many dependencies to mock)
   - Break the module apart so it does less
- Its dependencies are too fine-grained
   - Think about how you can consolidate some of these dependencies into one meaningful module
- Your test is too concerned with implementation details
   - Favour testing expected behaviour rather than the implementation

- 你正在测试的事情是做太多的事情（因为它有太多的依赖来模拟）
  - 将模块拆开，使其作用更少
- 它的依赖关系过于细粒度
  - 考虑如何将这些依赖项中的一些合并到一个有意义的模块中
- 你的测试过于关注实现细节
  - 倾向于测试预期的行为而不是实现

Normally a lot of mocking points to _bad abstraction_ in your code.

通常，很多嘲讽都指向代码中的 _bad 抽象_。

**What people see here is a weakness in TDD but it is actually a strength**, more often than not poor test code is a result of bad design or put more nicely, well-designed code is easy to test.

**人们在这里看到的是 TDD 的一个弱点，但它实际上是一个优势**，糟糕的测试代码往往是糟糕设计的结果，或者说得更好，设计良好的代码很容易测试。

### But mocks and tests are still making my life hard!

### 但是模拟和测试仍然让我的生活变得艰难！

Ever run into this situation?

有没有遇到过这种情况？

- You want to do some refactoring
- To do this you end up changing lots of tests
- You question TDD and make a post on Medium titled "Mocking considered harmful"

- 你想做一些重构
- 要做到这一点，您最终需要更改大量测试
- 你质疑 TDD 并在 Medium 上发了一篇题为“嘲笑被认为有害”的帖子

This is usually a sign of you testing too much _implementation detail_. Try to make it so your tests are testing _useful behaviour_ unless the implementation is really important to how the system runs.

这通常表明您测试了太多_实现细节_。尽量让你的测试测试_有用的行为_，除非实现对系统的运行方式非常重要。

It is sometimes hard to know _what level_ to test exactly but here are some thought processes and rules I try to follow:

有时很难知道要准确测试_什么级别_，但这里有一些我尝试遵循的思考过程和规则：

- **The definition of refactoring is that the code changes but the behaviour stays the same**. If you have decided to do some refactoring in theory you should be able to make the commit without any test changes. So when writing a test ask yourself
   - Am I testing the behaviour I want, or the implementation details?
   - If I were to refactor this code, would I have to make lots of changes to the tests?
- Although Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behaviour. Test the public behaviour. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.
- I feel like if a test is working with **more than 3 mocks then it is a red flag** - time for a rethink on the design 

- **重构的定义是代码改变但行为保持不变**。如果您决定在理论上进行一些重构，您应该能够在没有任何测试更改的情况下进行提交。所以在编写测试时问问自己
  - 我是在测试我想要的行为还是实现细节？
  - 如果我要重构此代码，是否必须对测试进行大量更改？
- 虽然 Go 允许您测试私有函数，但我会避免它，因为私有函数是支持公共行为的实现细节。测试公众行为。 Sandi Metz 将私有函数描述为“不太稳定”，您不希望将测试与它们耦合。
- 我觉得如果一个测试正在使用 ** 超过 3 个模拟，那么这是一个危险信号** - 是时候重新思考设计了

- Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means a tighter coupling between your test code and the implementation. **Be sure you actually care about these details if you're going to spy on them**

- 谨慎使用间谍。 Spies 让您可以看到您正在编写的算法的内部，这可能非常有用，但这意味着您的测试代码和实现之间的耦合更紧密。 **如果你要监视它们，请确保你真的关心这些细节**

#### Can't I just use a mocking framework?

#### 我不能只使用模拟框架吗？

Mocking requires no magic and is relatively simple; using a framework can make mocking seem more complicated than it is. We don't use automocking in this chapter so that we get:

Mocking 不需要魔法，比较简单；使用框架会使模拟看起来比实际更复杂。我们在本章中不使用自动模拟，因此我们得到：

- a better understanding of how to mock
- practise implementing interfaces

- 更好地理解如何模拟
- 练习实现接口

In collaborative projects there is value in auto-generating mocks. In a team, a mock generation tool codifies consistency around the test doubles. This will avoid inconsistently written test doubles which can translate to inconsistently written tests.

在协作项目中，自动生成模拟很有价值。在一个团队中，一个模拟生成工具编码了测试替身的一致性。这将避免编写不一致的测试替身，这可能会转化为编写不一致的测试。

You should only use a mock generator that generates test doubles against an interface. Any tool that overly dictates how tests are written, or that use lots of 'magic', can get in the sea.

您应该只使用模拟生成器来针对接口生成测试替身。任何过度规定测试编写方式或使用大量“魔法”的工具都可能陷入困境。

## Wrapping up

##  总结

### More on TDD approach

### 更多关于 TDD 方法

- When faced with less trivial examples, break the problem down into "thin vertical slices". Try to get to a point where you have _working software backed by tests_ as soon as you can, to avoid getting in rabbit holes and taking a "big bang" approach.
- Once you have some working software it should be easier to _iterate with small steps_ until you arrive at the software you need.

- 面对不那么琐碎的例子时，将问题分解为“细小的垂直切片”。尝试尽快达到_工作软件由测试支持的程度，以避免陷入困境并采取“大爆炸”方法。
- 一旦你有了一些可以工作的软件，它应该更容易_迭代小步骤_直到你得到你需要的软件。

> "When to use iterative development? You should use iterative development only on projects that you want to succeed."

> “什么时候使用迭代开发？你应该只在想要成功的项目上使用迭代开发。”

Martin Fowler.

马丁福勒。

### Mocking

### 嘲讽

- **Without mocking important areas of your code will be untested**. In our case we would not be able to test that our code paused between each print but there are countless other examples. Calling a service that _can_ fail? Wanting to test your system in a particular state? It is very hard to test these scenarios without mocking.
- Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in **slow feedback loops**.
- By having to spin up a database or a webservice to test something you're likely to have **fragile tests** due to the unreliability of such services.

- **不模拟你的代码的重要区域将是未经测试的**。在我们的例子中，我们无法测试我们的代码在每次打印之间是否暂停，但还有无数其他示例。调用一个_可以_失败的服务？想要在特定状态下测试您的系统？在不模拟的情况下很难测试这些场景。
- 如果没有模拟，您可能需要设置数据库和其他第三方的东西来测试简单的业务规则。您可能会进行缓慢的测试，从而导致 **缓慢的反馈循环**。
- 由于必须启动数据库或网络服务来测试某些内容，您可能会因为此类服务的不可靠性而进行 **脆弱的测试**。

Once a developer learns about mocking it becomes very easy to over-test every single facet of a system in terms of the _way it works_ rather than _what it does_. Always be mindful about **the value of your tests** and what impact they would have in future refactoring.

一旦开发人员学会了模拟，就很容易根据_它的工作方式_而不是_它的作用_来过度测试系统的每个方面。始终注意**测试的价值**以及它们在未来重构中的影响。

In this post about mocking we have only covered **Spies** which are a kind of mock. The "proper" term for mocks though are "test doubles"

在这篇关于模拟的文章中，我们只介绍了 **Spies**，这是一种模拟。模拟的“正确”术语是“测试替身”

[> Test Double is a generic term for any case where you replace a production object for testing purposes.](https://martinfowler.com/bliki/TestDouble.html)

[> Test Double 是任何为了测试目的而替换生产对象的情况的通用术语。](https://martinfowler.com/bliki/TestDouble.html)

Under test doubles, there are various types like stubs, spies and indeed mocks! Check out [Martin Fowler's post](https://martinfowler.com/bliki/TestDouble.html) for more detail. 

在测试替身下，有各种类型，如存根、间谍和模拟！查看 [Martin Fowler 的帖子](https://martinfowler.com/bliki/TestDouble.html) 了解更多详情。

