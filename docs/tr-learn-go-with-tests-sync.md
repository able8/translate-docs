# Sync

#  同步

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/sync)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/sync)**

We want to make a counter which is safe to use concurrently.

我们想制作一个可以安全并发使用的计数器。

We'll start with an unsafe counter and verify its behaviour works in a single-threaded environment.

我们将从一个不安全的计数器开始，并验证它的行为在单线程环境中是否有效。

Then we'll exercise it's unsafeness with multiple goroutines trying to use it via a test and fix it.

然后我们将使用多个 goroutine 尝试通过测试使用它并修复它，从而证明它的不安全性。

## Write the test first

## 先写测试

We want our API to give us a method to increment the counter and then retrieve its value.

我们希望我们的 API 给我们一个方法来增加计数器然后检索它的值。

```go
func TestCounter(t *testing.T) {
    t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
        counter := Counter{}
        counter.Inc()
        counter.Inc()
        counter.Inc()

        if counter.Value() != 3 {
            t.Errorf("got %d, want %d", counter.Value(), 3)
        }
    })
}
```

## Try to run the test

## 尝试运行测试

```
./sync_test.go:9:14: undefined: Counter
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Let's define `Counter`.

让我们定义“计数器”。

```go
type Counter struct {

}
```

Try again and it fails with the following

再试一次，但失败并显示以下内容

```
./sync_test.go:14:10: counter.Inc undefined (type Counter has no field or method Inc)
./sync_test.go:18:13: counter.Value undefined (type Counter has no field or method Value)
```

So to finally make the test run we can define those methods

所以为了最终运行测试，我们可以定义这些方法

```go
func (c *Counter) Inc() {

}

func (c *Counter) Value() int {
    return 0
}
```

It should now run and fail

它现在应该运行并失败

```
=== RUN   TestCounter
=== RUN   TestCounter/incrementing_the_counter_3_times_leaves_it_at_3
--- FAIL: TestCounter (0.00s)
    --- FAIL: TestCounter/incrementing_the_counter_3_times_leaves_it_at_3 (0.00s)
        sync_test.go:27: got 0, want 3
```

## Write enough code to make it pass

## 编写足够的代码使其通过

This should be trivial for Go experts like us. We need to keep some state for the counter in our datatype and then increment it on every `Inc` call

对于像我们这样的 Go 专家来说，这应该是微不足道的。我们需要在我们的数据类型中为计数器保留一些状态，然后在每次 `Inc` 调用时增加它

```go
type Counter struct {
    value int
}

func (c *Counter) Inc() {
    c.value++
}

func (c *Counter) Value() int {
    return c.value
}
```

## Refactor

## 重构

There's not a lot to refactor but given we're going to write more tests around `Counter` we'll write a small assertion function `assertCount` so the test reads a bit clearer.

没有太多要重构的，但考虑到我们将围绕 `Counter` 编写更多测试，我们将编写一个小的断言函数 `assertCount`，以便测试读取更清晰一些。

```go
t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
    counter := Counter{}
    counter.Inc()
    counter.Inc()
    counter.Inc()

    assertCounter(t, counter, 3)
})

func assertCounter(t testing.TB, got Counter, want int)  {
    t.Helper()
    if got.Value() != want {
        t.Errorf("got %d, want %d", got.Value(), want)
    }
}
```

## Next steps

##  下一步

That was easy enough but now we have a requirement that it must be safe to use in a concurrent environment. We will need to write a failing test to exercise this.

这很容易，但现在我们要求在并发环境中使用它必须是安全的。我们需要编写一个失败的测试来练习这个。

## Write the test first

## 先写测试

```go
t.Run("it runs safely concurrently", func(t *testing.T) {
    wantedCount := 1000
    counter := Counter{}

    var wg sync.WaitGroup
    wg.Add(wantedCount)

    for i := 0;i < wantedCount;i++ {
        go func() {
            counter.Inc()
            wg.Done()
        }()
    }
    wg.Wait()

    assertCounter(t, counter, wantedCount)
})
```

This will loop through our `wantedCount` and fire a goroutine to call `counter.Inc()`.

这将遍历我们的 `wantedCount` 并触发一个 goroutine 来调用 `counter.Inc()`。

We are using [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup) which is a convenient way of synchronising concurrent processes.

我们正在使用 [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup)，这是一种同步并发进程的便捷方式。

> A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

> WaitGroup 等待一组 goroutine 完成。 main goroutine 调用 Add 来设置要等待的 goroutine 的数量。然后每个 goroutine 运行并在完成时调用 Done。同时，Wait 可用于阻塞，直到所有 goroutine 完成。

By waiting for `wg.Wait()` to finish before making our assertions we can be sure all of our goroutines have attempted to `Inc` the `Counter`.

通过在断言之前等待 `wg.Wait()` 完成，我们可以确定我们所有的 goroutine 都尝试了 `Inc` 和 `Counter`。

## Try to run the test

## 尝试运行测试

```
=== RUN   TestCounter/it_runs_safely_in_a_concurrent_envionment
--- FAIL: TestCounter (0.00s)
    --- FAIL: TestCounter/it_runs_safely_in_a_concurrent_envionment (0.00s)
        sync_test.go:26: got 939, want 1000
FAIL
```

The test will _probably_ fail with a different number, but nonetheless it demonstrates it does not work when multiple goroutines are trying to mutate the value of the counter at the same time.

该测试 _probably_ 会以不同的数字失败，但尽管如此，它表明当多个 goroutine 试图同时改变计数器的值时它不起作用。

## Write enough code to make it pass

## 编写足够的代码使其通过

A simple solution is to add a lock to our `Counter`, a [`Mutex`](https://golang.org/pkg/sync/#Mutex)

一个简单的解决方案是给我们的 `Counter` 添加一个锁，一个 [`Mutex`](https://golang.org/pkg/sync/#Mutex)

>A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.

>互斥锁是互斥锁。互斥锁的零值是未锁定的互斥锁。

```go
type Counter struct {
    mu sync.Mutex
    value int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

What this means is any goroutine calling `Inc` will acquire the lock on `Counter` if they are first. All the other goroutines will have to wait for it to be `Unlock`ed before getting access.

这意味着任何调用 `Inc` 的 goroutine 都将获得 `Counter` 上的锁，如果它们是第一个的话。所有其他 goroutine 都必须等待它被“解锁”才能访问。

If you now re-run the test it should now pass because each goroutine has to wait its turn before making a change.

如果您现在重新运行测试，它现在应该会通过，因为每个 goroutine 在进行更改之前都必须等待轮到它。

## I've seen other examples where the `sync.Mutex` is embedded into the struct.

## 我看过其他示例，其中 `sync.Mutex` 嵌入到结构中。

You may see examples like this

你可能会看到这样的例子

```go
type Counter struct {
    sync.Mutex
    value int
}
```

It can be argued that it can make the code a bit more elegant.

可以说它可以使代码更优雅一些。

```go
func (c *Counter) Inc() {
    c.Lock()
    defer c.Unlock()
    c.value++
}
```

This _looks_ nice but while programming is a hugely subjective discipline, this is **bad and wrong**.

这_看起来_不错，但虽然编程是一门非常主观的学科，但这是**糟糕和错误的**。

Sometimes people forget that embedding types means the methods of that type becomes _part of the public interface_; and you often will not want that. Remember that we should be very careful with our public APIs, the moment we make something public is the moment other code can couple themselves to it. We always want to avoid unnecessary coupling.

有时人们会忘记嵌入类型意味着该类型的方法成为_公共接口的一部分_；而你通常不会想要那样。请记住，我们应该非常小心我们的公共 API，当我们将某些东西公开时，其他代码就可以将它们自己耦合到它。我们总是想避免不必要的耦合。

Exposing `Lock` and `Unlock` is at best confusing but at worst potentially very harmful to your software if callers of your type start calling these methods.

暴露 `Lock` 和 `Unlock` 充其量令人困惑，但最糟糕的是，如果您的类型的调用者开始调用这些方法，则可能对您的软件非常有害。

![Showing how a user of this API can wrongly change the state of the lock](https://i.imgur.com/SWYNpwm.png)

_This seems like a really bad idea_

_这似乎是一个非常糟糕的主意_

## Copying mutexes

## 复制互斥锁

Our test passes but our code is still a bit dangerous

我们的测试通过了，但我们的代码仍然有点危险

If you run `go vet` on your code you should get an error like the following

如果你在你的代码上运行 `go vet` 你应该得到如下错误

```
sync/v2/sync_test.go:16: call of assertCounter copies lock value: v1.Counter contains sync.Mutex
sync/v2/sync_test.go:39: assertCounter passes lock by value: v1.Counter contains sync.Mutex
```

A look at the documentation of [`sync.Mutex`](https://golang.org/pkg/sync/#Mutex) tells us why

看一下 [`sync.Mutex`](https://golang.org/pkg/sync/#Mutex) 的文档告诉我们为什么

> A Mutex must not be copied after first use.

> 第一次使用后不得复制互斥锁。

When we pass our `Counter` (by value) to `assertCounter` it will try and create a copy of the mutex.

当我们将 `Counter`（按值）传递给 `assertCounter` 时，它会尝试创建互斥锁的副本。

To solve this we should pass in a pointer to our `Counter` instead, so change the signature of `assertCounter`

为了解决这个问题，我们应该传递一个指向我们的 `Counter` 的指针，因此更改 `assertCounter` 的签名

```go
func assertCounter(t testing.TB, got *Counter, want int)
```

Our tests will no longer compile because we are trying to pass in a `Counter` rather than a `*Counter`. To solve this I prefer to create a constructor which shows readers of your API that it would be better to not initialise the type yourself.

我们的测试将不再编译，因为我们试图传入一个 `Counter` 而不是 `*Counter`。为了解决这个问题，我更喜欢创建一个构造函数，它向 API 的读者展示最好不要自己初始化类型。

```go
func NewCounter() *Counter {
    return &Counter{}
}
```

Use this function in your tests when initialising `Counter`.

在初始化 `Counter` 时，在你的测试中使用这个函数。

## Wrapping up

##  总结

We've covered a few things from the [sync package](https://golang.org/pkg/sync/)

我们已经介绍了 [sync 包](https://golang.org/pkg/sync/) 中的一些内容

- `Mutex` allows us to add locks to our data
- `Waitgroup` is a means of waiting for goroutines to finish jobs

- `Mutex` 允许我们为我们的数据添加锁
- `Waitgroup` 是一种等待 goroutine 完成工作的方法

### When to use locks over channels and goroutines?

### 何时在通道和 goroutine 上使用锁？

[We've previously covered goroutines in the first concurrency chapter](concurrency.md) which let us write safe concurrent code so why would you use locks?
[The go wiki has a page dedicated to this topic; Mutex Or Channel](https://github.com/golang/go/wiki/MutexOrChannel)

[我们之前在并发第一章中介绍了 goroutines](concurrency.md)，它让我们可以编写安全的并发代码，那么为什么要使用锁呢？
[go wiki 有一个专门讨论这个主题的页面；互斥或通道](https://github.com/golang/go/wiki/MutexOrChannel)

> A common Go newbie mistake is to over-use channels and goroutines just because it's possible, and/or because it's fun. Don't be afraid to use a sync.Mutex if that fits your problem best. Go is pragmatic in letting you use the tools that solve your problem best and not forcing you into one style of code.

> 一个常见的 Go 新手错误是过度使用 channel 和 goroutines，只是因为它是可能的，和/或因为它很有趣。如果最适合您的问题，请不要害怕使用 sync.Mutex。 Go 是务实的，它让您使用最能解决问题的工具，而不是强迫您使用一种代码风格。

Paraphrasing:

释义：

- **Use channels when passing ownership of data**
- **Use mutexes for managing state**

- **在传递数据所有权时使用通道**
- **使用互斥锁来管理状态**

### go vet

### 去看兽医

Remember to use go vet in your build scripts as it can alert you to some subtle bugs in your code before they hit your poor users.

请记住在您的构建脚本中使用 go vet，因为它可以在代码中的一些细微错误影响您的不良用户之前提醒您。

### Don't use embedding because it's convenient

### 不要使用嵌入，因为它很方便

- Think about the effect embedding has on your public API.
- Do you _really_ want to expose these methods and have people coupling their own code to them?
- With respect to mutexes, this could be potentially disastrous in very unpredictable and weird ways, imagine some nefarious code unlocking a mutex when it shouldn't be; this would cause some very strange bugs that will be hard to track down. 

- 考虑嵌入对您的公共 API 的影响。
- 你_真的_想要公开这些方法并让人们将他们自己的代码耦合到它们吗？
- 关于互斥锁，这可能会以非常不可预测和奇怪的方式造成灾难性的后果，想象一下一些邪恶的代码在不应该解锁互斥锁时解锁；这会导致一些很难追踪的非常奇怪的错误。

