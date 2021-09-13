# Context-aware readers

# 上下文感知的读者

**[You can find all the code here](https://github.com/quii/learn-go-with-tests/tree/main/q-and-a/context-aware-reader)**

**[你可以在这里找到所有代码](https://github.com/quii/learn-go-with-tests/tree/main/q-and-a/context-aware-reader)**

This chapter demonstrates how to test-drive a context aware `io.Reader` as written by Mat Ryer and David Hernandez in [The Pace Dev Blog](https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer).

本章演示了如何测试驱动上下文感知的 `io.Reader`，如 Mat Ryer 和 David Hernandez 在 [The Pace Dev Blog](https://pace.dev/blog/2020/02/03/context-意识到-ioreader-for-golang-by-mat-ryer)。

## Context aware reader?

## 上下文感知阅读器？

First of all, a quick primer on `io.Reader`.

首先，快速入门`io.Reader`。

If you've read other chapters in this book you will have ran into `io.Reader` when we've opened files, encoded JSON and various other common tasks. It's a simple abstraction over reading data from _something_

如果你读过本书的其他章节，当我们打开文件、编码 JSON 和各种其他常见任务时，你会遇到 `io.Reader`。这是对从 _something_ 读取数据的简单抽象

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

By using `io.Reader` you can gain a lot of re-use from the standard library, it's a very commonly used abstraction (along with its counterpart `io.Writer`)

通过使用`io.Reader`，你可以从标准库中获得大量的重用，它是一个非常常用的抽象（连同它的对应物`io.Writer`）

### Context aware?

### 上下文感知？

[In a previous chapter](context.md) we discussed how we can use `context` to provide cancellation. This is especially useful if you're performing tasks which may be computationally expensive and you want to be able to stop them.

[在前一章](context.md) 我们讨论了如何使用`context` 来提供取消。如果您正在执行可能计算成本很高的任务并且您希望能够停止它们，这将特别有用。

When you're using an `io.Reader` you have no guarantees over speed, it could take 1 nanosecond or hundreds of hours. You might find it useful to be able to cancel these kind of tasks in your own application and that's what Mat and David wrote about.

当您使用 `io.Reader` 时，您无法保证速度，它可能需要 1 纳秒或数百小时。您可能会发现能够在您自己的应用程序中取消这些类型的任务很有用，这就是 Mat 和 David 所写的。

They combined two simple abstractions (`context.Context` and `io.Reader`) to solve this problem.

他们结合了两个简单的抽象（`context.Context` 和 `io.Reader`）来解决这个问题。

Let's try and TDD some functionality so that we can wrap an `io.Reader` so it can be cancelled.

让我们尝试 TDD 一些功能，以便我们可以包装一个 `io.Reader` 以便可以取消它。

Testing this poses an interesting challenge. Normally when using an `io.Reader` you're usually supplying it to some other function and you dont really concern yourself with the details; such as `json.NewDecoder` or `ioutil.ReadAll`.

对此进行测试提出了一个有趣的挑战。通常在使用 `io.Reader` 时，你通常将它提供给其他一些函数，而你并不真正关心细节；例如`json.NewDecoder` 或`ioutil.ReadAll`。

What we want to demonstrate is something like

我们想要展示的是类似的东西

> Given an `io.Reader` with "ABCDEF", when I send a cancel signal half-way through I when I try to continue to read I get nothing else so all I get is "ABC"

> 给定一个带有“ABCDEF”的“io.Reader”，当我在中途发送取消信号时，当我尝试继续阅读时，我什么也得不到，所以我得到的只是“ABC”

Let's look at the interface again.

我们再来看看界面。

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

The `Reader`'s `Read` method will read the contents it has into a `[]byte` that we supply.

`Reader` 的 `Read` 方法将把它的内容读入我们提供的 `[]byte`。

So rather than reading everything, we could:

  因此，与其阅读所有内容，我们还可以：

 - Supply a fixed-size byte array that doesnt fit all the contents
  - Send a cancel signal
  - Try and read again and this should return an error with 0 bytes read

- 提供不适合所有内容的固定大小的字节数组
 - 发送取消信号
 - 再次尝试读取，这应该返回读取 0 字节的错误

For now, let's just write a "happy path" test where there is no cancellation, just so we can get familiar with the problem without having to write any production code yet.

现在，让我们编写一个没有取消的“happy path”测试，这样我们就可以熟悉问题，而无需编写任何生产代码。

```go
func TestContextAwareReader(t *testing.T) {
    t.Run("lets just see how a normal reader works", func(t *testing.T) {
        rdr := strings.NewReader("123456")
        got := make([]byte, 3)
        _, err := rdr.Read(got)

        if err != nil {
            t.Fatal(err)
        }

        assertBufferHas(t, got, "123")

        _, err = rdr.Read(got)

        if err != nil {
            t.Fatal(err)
        }

        assertBufferHas(t, got, "456")
    })
}

func assertBufferHas(t testing.TB, buf []byte, want string) {
    t.Helper()
    got := string(buf)
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

- Make an `io.Reader` from a string with some data
- A byte array to read into which is smaller than the contents of the reader
- Call read, check the contents, repeat.

- 从带有一些数据的字符串中创建一个 `io.Reader`
- 要读入的字节数组，其小于读取器的内容
- 调用阅读，检查内容，重复。

From this we can imagine sending some kind of cancel signal before the second read to change behaviour.

由此我们可以想象在第二次读取之前发送某种取消信号来改变行为。

Now we've seen how it works we'll TDD the rest of the functionality.

现在我们已经看到了它是如何工作的，我们将 TDD 其余的功能。

## Write the test first

## 先写测试

We want to be able to compose an `io.Reader` with a `context.Context`.

我们希望能够将一个 `io.Reader` 与一个 `context.Context` 组合起来。

With TDD it's best to start with imagining your desired API and write a test for it.

使用 TDD，最好从想象您想要的 API 开始，然后为它编写测试。

From there let the compiler and failing test output can guide us to a solution

从那里让编译器和失败的测试输出可以指导我们找到解决方案

```go
t.Run("behaves like a normal reader", func(t *testing.T) {
    rdr := NewCancellableReader(strings.NewReader("123456"))
    got := make([]byte, 3)
    _, err := rdr.Read(got)

    if err != nil {
        t.Fatal(err)
    }

    assertBufferHas(t, got, "123")

    _, err = rdr.Read(got)

    if err != nil {
        t.Fatal(err)
    }

    assertBufferHas(t, got, "456")
})
```

## Try to run the test

## 尝试运行测试

```
./cancel_readers_test.go:12:10: undefined: NewCancellableReader
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We'll need to define this function and it should return an `io.Reader`

我们需要定义这个函数，它应该返回一个 `io.Reader`

```go
func NewCancellableReader(rdr io.Reader) io.Reader {
    return nil
}
```

If you try and run it

如果你尝试运行它

```
=== RUN   TestCancelReaders
=== RUN   TestCancelReaders/behaves_like_a_normal_reader
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
    panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10f8fb5]
```

As expected

正如预期的那样

## Write enough code to make it pass

## 编写足够的代码使其通过

For now, we'll just return the `io.Reader` we pass in

现在，我们只返回我们传入的 `io.Reader`

```go
func NewCancellableReader(rdr io.Reader) io.Reader {
    return rdr
}
```

The test should now pass.

测试现在应该通过。

I know, I know, this seems silly and pedantic but before charging in to the fancy work it is important that we have _some_ verification that we haven't broken the "normal" behaviour of an `io.Reader` and this test will give us confidence as we move forward.

我知道，我知道，这看起来很愚蠢和迂腐，但在投入到花哨的工作之前，重要的是我们有_一些_验证我们没有破坏 `io.Reader` 的“正常”行为，这个测试将给出我们在前进的过程中充满信心。

## Write the test first

## 先写测试

Next we need to try and cancel.

接下来我们需要尝试取消。

```go
t.Run("stops reading when cancelled", func(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    rdr := NewCancellableReader(ctx, strings.NewReader("123456"))
    got := make([]byte, 3)
    _, err := rdr.Read(got)

    if err != nil {
        t.Fatal(err)
    }

    assertBufferHas(t, got, "123")

    cancel()

    n, err := rdr.Read(got)

    if err == nil {
        t.Error("expected an error after cancellation but didnt get one")
    }

    if n > 0 {
        t.Errorf("expected 0 bytes to be read after cancellation but %d were read", n)
    }
})
```

We can more or less copy the first test but now we're:
- Creating a `context.Context` with cancellation so we can `cancel` after the first read
- For our code to work we'll need to pass `ctx` to our function
- We then assert that post-`cancel` nothing was read

我们或多或少可以复制第一个测试，但现在我们：
- 创建一个带有取消的`context.Context`，这样我们就可以在第一次读取后`cancel`
- 为了让我们的代码正常工作，我们需要将 `ctx` 传递给我们的函数
- 然后我们断言 post-`cancel` 没有被读取

## Try to run the test

## 尝试运行测试

```
./cancel_readers_test.go:33:30: too many arguments in call to NewCancellableReader
    have (context.Context, *strings.Reader)
    want (io.Reader)
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

The compiler is telling us what to do; update our signature to accept a context

编译器告诉我们要做什么；更新我们的签名以接受上下文

```go
func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
    return rdr
}
```

(You'll need to update the first test to pass in `context.Background` too)

（您也需要更新第一个测试以通过 `context.Background`）

You should now see a very clear failing test output

您现在应该看到一个非常清晰的失败测试输出

```
=== RUN   TestCancelReaders
=== RUN   TestCancelReaders/stops_reading_when_cancelled
--- FAIL: TestCancelReaders (0.00s)
    --- FAIL: TestCancelReaders/stops_reading_when_cancelled (0.00s)
        cancel_readers_test.go:48: expected an error but didnt get one
        cancel_readers_test.go:52: expected 0 bytes to be read after cancellation but 3 were read
```

## Write enough code to make it pass

## 编写足够的代码使其通过

At this point, it's copy and paste from the original post by Mat and David but we'll still take it slowly and iteratively.

在这一点上，它是从 Mat 和 David 的原始帖子中复制和粘贴的，但我们仍然会慢慢地反复进行。

We know we need to have a type that encapsulates the `io.Reader` that we read from and the `context.Context` so let's create that and try and return it from our function instead of the original `io.Reader`

我们知道我们需要有一个类型来封装我们从中读取的 `io.Reader` 和 `context.Context`，所以让我们创建它并尝试从我们的函数而不是原始的 `io.Reader` 中返回它

```go
func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
    return &readerCtx{
        ctx:      ctx,
        delegate: rdr,
    }
}

type readerCtx struct {
    ctx      context.Context
    delegate io.Reader
}
```

As I have stressed many times in this book, go slowly and let the compiler help you

正如我在本书中多次强调的那样，慢慢来，让编译器帮你

```go
./cancel_readers_test.go:60:3: cannot use &readerCtx literal (type *readerCtx) as type io.Reader in return argument:
    *readerCtx does not implement io.Reader (missing Read method)
```

The abstraction feels right, but it doesn't implement the interface we need (`io.Reader`) so let's add the method.

抽象感觉是对的，但它没有实现我们需要的接口（`io.Reader`），所以让我们添加方法。

```go
func (r *readerCtx) Read(p []byte) (n int, err error) {
    panic("implement me")
}
```

Run the tests and they should _compile_ but panic. This is still progress.

运行测试，他们应该 _compile_ 但恐慌。这仍然是进步。

Let's make the first test pass by just _delegating_ the call to our underlying `io.Reader`

让我们通过_委托_调用我们底层的`io.Reader`来使第一个测试通过

```go
func (r readerCtx) Read(p []byte) (n int, err error) {
    return r.delegate.Read(p)
}
```

At this point we have our happy path test passing again and it feels like we have our stuff abstracted nicely

在这一点上，我们的快乐路径测试再次通过，感觉我们的东西抽象得很好

To make our second test pass we need to check the `context.Context` to see if it has been cancelled.

为了使我们的第二个测试通过，我们需要检查 `context.Context` 以查看它是否已被取消。

```go
func (r readerCtx) Read(p []byte) (n int, err error) {
    if err := r.ctx.Err();err != nil {
        return 0, err
    }
    return r.delegate.Read(p)
}
```

All tests should now pass. You'll notice how we return the error from the `context.Context`. This allows callers of the code to inspect the various reasons cancellation has occurred and this is covered more in the original post.

现在应该通过所有测试。您会注意到我们如何从 `context.Context` 返回错误。这允许代码的调用者检查发生取消的各种原因，这在原始帖子中有更多的介绍。

## Wrapping up

##  总结

- Small interfaces are good and are easily composed 

- 小接口很好，很容易组合

- When you're trying to augment one thing (e.g `io.Reader`) with another you usually want to reach for the [delegation pattern](https://en.wikipedia.org/wiki/Delegation_pattern)

- 当你试图用另一件事（例如`io.Reader`）来增加你通常想要达到的[委托模式](https://en.wikipedia.org/wiki/Delegation_pattern)

> In software engineering, the delegation pattern is an object-oriented design pattern that allows object composition to achieve the same code reuse as inheritance.

> 在软件工程中，委托模式是一种面向对象的设计模式，它允许对象组合实现与继承相同的代码重用。

- An easy way to start this kind of work is to wrap your delegate and write a test that asserts it behaves how the delegate normally does before you start composing other parts to change behaviour. This will help you to keep things working correctly as you code toward your goal 

- 开始此类工作的一种简单方法是包装您的委托并编写一个测试，在您开始编写其他部分以更改行为之前，该测试断言它的行为与委托的正常行为一样。这将帮助您在朝着目标编码时保持正常工作

