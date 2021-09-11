# Common Concurrent Programming Mistakes

# 常见的并发编程错误

Go is a language supporting built-in concurrent programming. By using the `go` keyword to create goroutines (light weight threads) and by [using](https://go101.org/article/channel-use-cases.html) [channels](https://go101.org/article/channel.html) and [other concurrency](https://go101.org/article/concurrent-atomic-operation.html) [synchronization techniques](https://go101.org/article/concurrent-synchronization-more.html) provided in Go, concurrent programming becomes easy, flexible and enjoyable.

Go 是一种支持内置并发编程的语言。通过使用 `go` 关键字创建 goroutines（轻量级线程）并通过 [using](https://go101.org/article/channel-use-cases.html) [channels](https://go101.org/article/channel.html) 和 [其他并发](https://go101.org/article/concurrent-atomic-operation.html) [同步技术](https://go101.org/article/concurrent-synchronization-more.html) 在 Go 中提供，并发编程变得简单、灵活和愉快。

One the other hand, Go doesn't prevent Go programmers from making some concurrent programming mistakes which are caused by either carelessness or lacking of experience. The remaining of the current article will show some common mistakes in Go concurrent programming, to help Go programmers avoid making such mistakes.

另一方面，Go 并不能阻止 Go 程序员犯一些由于粗心或缺乏经验而导致的并发编程错误。当前文章的剩余部分将展示 Go 并发编程中的一些常见错误，以帮助 Go 程序员避免犯此类错误。

### No Synchronizations When Synchronizations Are Needed

### 当需要同步时没有同步

Code lines might be [not executed by their appearance order](https://go101.org/article/memory-model.html).

代码行可能[不按其出现顺序执行](https://go101.org/article/memory-model.html)。

There are two mistakes in the following program.

下面的程序有两个错误。

-     First, the read of `b` in the main goroutine and the write of `b` in the new goroutine might cause data races.
-     Second, the condition `b == true` can't ensure that `a != nil` in the main goroutine. Compilers and CPUs may make optimizations by [reordering instructions](https://go101.org/article/memory-model.html) in the new goroutine, so the assignment of `b` may happen before the assignment of `a` at run time, which makes that slice `a` is still `nil` when the elements of `a` are modified in the main goroutine.

- 首先，主goroutine中`b`的读取和新goroutine中`b`的写入可能会导致数据竞争。
- 其次，条件 `b == true` 不能确保主 goroutine 中的 `a != nil`。编译器和CPU可能会在新的goroutine中通过[重新排序指令](https://go101.org/article/memory-model.html)进行优化，所以`b`的赋值可能发生在`a`的赋值之前运行时，当主 goroutine 中修改了 `a` 的元素时，这使得切片 `a` 仍然是 `nil`。

```go
package main

import (
    "time"
    "runtime"
)

func main() {
    var a []int // nil
    var b bool  // false

    // a new goroutine
    go func () {
        a = make([]int, 3)
        b = true // write b
    }()

    for !b { // read b
        time.Sleep(time.Second)
        runtime.Gosched()
    }
    a[0], a[1], a[2] = 0, 1, 2 // might panic
}
```


The above program may run well on one computer, but may panic on another one, or it runs well when it is compiled by one compiler, but panics when another compiler is used.

上述程序可能在一台计算机上运行良好，但在另一台计算机上可能会发生 panic，或者在一个编译器编译时运行良好，但在使用另一个编译器时发生 panic。

We should use channels or the synchronization techniques provided in the `sync` standard package to ensure the memory orders. For example,

我们应该使用“sync”标准包中提供的通道或同步技术来确保内存顺序。例如，

```go
package main

func main() {
    var a []int = nil
    c := make(chan struct{})

    go func () {
        a = make([]int, 3)
        c <- struct{}{}
    }()

    <-c
    // The next line will not panic for sure.
    a[0], a[1], a[2] = 0, 1, 2
}
```






### Use `time.Sleep` Calls to Do Synchronizations

### 使用`time.Sleep` 调用进行同步

Let's view a simple example.

让我们看一个简单的例子。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    var x = 123

    go func() {
        x = 789 // write x
    }()

    time.Sleep(time.Second)
    fmt.Println(x) // read x
}
```


We expect this program to print `789`. In fact, it really prints `789`, almost always, in running. But is it a program with good synchronization? No! The reason is Go runtime doesn't guarantee the write of `x` happens before the read of `x` for sure. Under certain conditions, such as most CPU resources are consumed by some other computation-intensive programs running on the same OS, the write of `x` might happen after the read of `x`. This is why we should never use `time.Sleep` calls to do synchronizations in formal projects.

我们希望这个程序打印“789”。事实上，它几乎总是在运行时打印“789”。但它是一个具有良好同步性的程序吗？不！原因是 Go 运行时并不能保证 x 的写入肯定在 x 的读取之前发生。在某些情况下，例如大部分 CPU 资源被运行在同一操作系统上的其他一些计算密集型程序消耗，可能会在读取 x 之后写入 x。这就是为什么我们不应该在正式项目中使用 `time.Sleep` 调用来进行同步。

Let's view another example.

让我们看另一个例子。

```go
package main

import (
    "fmt"
    "time"
)

var x = 0

func main() {
    var num = 123
    var p = &num

    c := make(chan int)

    go func() {
        c <- *p + x
    }()

    time.Sleep(time.Second)
    num = 789
    fmt.Println(<-c)
}
```


What do you expect the program will output? `123`, or `789`? In fact, the output is compiler dependent. For the standard Go compiler 1.17, it is very possible the program will output `123`. But in theory, it might output `789`, or another unexpected number.

你期望程序输出什么？ “123”还是“789”？事实上，输出依赖于编译器。对于标准的 Go 编译器 1.17，程序很可能会输出 `123`。但理论上，它可能会输出“789”，或者其他意想不到的数字。

Now, let's change `c <- *p + x` to `c <- *p` and run the program again, you will find the output becomes to `789` (for the standard Go compiler 1.17). Again, the output is compiler dependent.

现在，让我们将 `c <- *p + x` 改为 `c <- *p` 并再次运行程序，你会发现输出变成了 `789`（对于标准 Go 编译器 1.17）。同样，输出依赖于编译器。

Yes, there are data races in the above program. The expression `*p` might be evaluated before, after, or when the assignment `num = 789` is processed. The `time.Sleep` call can't guarantee the evaluation of `*p` happens before the assignment is processed.

是的，上述程序中存在数据竞争。表达式 `*p` 可能在赋值 `num = 789` 被处理之前、之后或时被求值。 `time.Sleep` 调用不能保证在处理分配之前对 `*p` 进行评估。

For this specified example, we should store the value to be sent in a temporary value before creating the new goroutine and send the temporary value instead in the new goroutine to remove the data races.

对于这个指定的例子，我们应该在创建新的 goroutine 之前将要发送的值存储在一个临时值中，并将临时值发送到新的 goroutine 中以消除数据竞争。

```go
...
    tmp := *p
    go func() {
        c <- tmp
    }()
...
```




### Leave Goroutines Hanging 

### 让 Goroutines 挂起

Hanging goroutines are the goroutines staying in blocking state for ever. There are many reasons leading goroutines into hanging. For example,

挂起的 goroutine 是永远处于阻塞状态的 goroutine。导致 goroutine 挂起的原因有很多。例如，

-     a goroutine tries to receive a value from a channel which no more other goroutines will send values to.
-     a goroutine tries to send a value to nil channel or to a channel which no more other goroutines will receive values from.
-     a goroutine is dead locked by itself.
-     a group of goroutines are dead locked by each other.
-     a goroutine is blocked when executing a `select` code block without `default` branch, and all the channel operations following the `case` keywords in the `select` code block keep blocking for ever.

- 一个 goroutine 尝试从一个通道接收一个值，其他 goroutine 不会将值发送到该通道。
- goroutine 尝试将值发送到 nil 通道或其他 goroutine 不会从中接收值的通道。
- 一个 goroutine 被自己死锁了。
- 一组 goroutine 被彼此死锁。
- 当执行一个没有 `default` 分支的 `select` 代码块时，goroutine 被阻塞，并且 `select` 代码块中的 `case` 关键字后面的所有通道操作将永远阻塞。

Except sometimes we deliberately let the main goroutine in a program hanging to avoid the program exiting, most other hanging goroutine cases are unexpected. It is hard for Go runtime to judge whether or not a goroutine in blocking state is hanging or stays in blocking state temporarily, so Go runtime will never release the resources consumed by a hanging goroutine.

除了有时我们故意让程序中的 main goroutine 挂起以避免程序退出之外，大多数其他挂起 goroutine 的情况都是意料之外的。 Go 运行时很难判断处于阻塞状态的 goroutine 是挂起还是暂时处于阻塞状态，因此 Go 运行时永远不会释放挂起的 goroutine 消耗的资源。

In the [first-response-wins](https://go101.org/article/channel-use-cases.html#first-response-wins) channel use case, if the capacity of the channel which is used a future is not large enough, some slower response goroutines will hang when trying to send a result to the future channel. For example, if the following function is called, there will be 4 goroutines stay in blocking state for ever.

在[first-response-wins](https://go101.org/article/channel-use-cases.html#first-response-wins)频道用例中，如果未来使用的频道容量为不够大，一些响应较慢的 goroutine 在尝试将结果发送到未来通道时会挂起。例如，如果调用以下函数，将有 4 个 goroutine 永远处于阻塞状态。

```go
func request() int {
    c := make(chan int)
    for i := 0;i < 5;i++ {
        i := i
        go func() {
            c <- i // 4 goroutines will hang here.
        }()
    }
    return <-c
}
```


To avoid the four goroutines hanging, the capacity of channel `c` must be at least `4`.

为了避免四个 goroutines 挂起，通道 `c` 的容量必须至少为 `4`。

In [the second way to implement the first-response-wins](https://go101.org/article/channel-use-cases.html#first-response-wins-2) channel use case, if the channel which is used as a future/promise is an unbuffered channel, like the following code shows, it is possible that the channel receiver will miss all responses and hang.

在 [第二种实现first-response-wins的方式](https://go101.org/article/channel-use-cases.html#first-response-wins-2)频道用例中，如果频道是用作未来/承诺的是一个无缓冲通道，如下面的代码所示，通道接收器可能会错过所有响应并挂起。

```go
func request() int {
    c := make(chan int)
    for i := 0;i < 5;i++ {
        i := i
        go func() {
            select {
            case c <- i:
            default:
            }
        }()
    }
    return <-c // might hang here
}
```


The reason why the receiver goroutine might hang is that if the five try-send operations all happen before the receive operation `<-c` is ready, then all the five try-send operations will fail to send values so that the caller goroutine will never receive a value.

接收者 goroutine 可能挂掉的原因是，如果 5 个 try-send 操作都发生在接收操作 `<-c` 准备好之前，那么所有 5 个 try-send 操作都将无法发送值，因此调用者 goroutine从不接收值。

Changing the channel `c` as a buffered channel will guarantee at least one of the five try-send operations succeed so that the caller goroutine will never hang in the above function.

将通道 `c` 更改为缓冲通道将保证至少五个 try-send 操作中的一个成功，以便调用者 goroutine 永远不会挂在上述函数中。

### Copy Values of the Types in the `sync` Standard Package

### 复制`sync` 标准包中类型的值

In practice, values of the types (except the `Locker` interface values) in the `sync` standard package should never be copied. We should only copy pointers of such values.

在实践中，`sync` 标准包中的类型值（`Locker` 接口值除外）不应该被复制。我们应该只复制这些值的指针。

The following is bad concurrent programming example. In this example, when the `Counter.Value` method is called, a `Counter` receiver value will be copied. As a field of the receiver value, the respective `Mutex` field of the `Counter` receiver value will also be copied. The copy is not synchronized, so the copied  `Mutex` value might be corrupted. Even if it is not corrupted, what it protects is the use of the copied field `n`, which is meaningless generally.

以下是糟糕的并发编程示例。在此示例中，当调用 `Counter.Value` 方法时，将复制一个 `Counter` 接收器值。作为接收器值的字段，“Counter”接收器值的相应“Mutex”字段也将被复制。副本未同步，因此复制的“Mutex”值可能已损坏。即使它没有被破坏，它保护的是使用复制的字段`n`，这通常是没有意义的。

```go
import "sync"

type Counter struct {
    sync.Mutex
    n int64
}

// This method is okay.
func (c *Counter) Increase(d int64) (r int64) {
    c.Lock()
    c.n += d
    r = c.n
    c.Unlock()
    return
}

// The method is bad.When it is called,
// the Counter receiver value will be copied.
func (c Counter) Value() (r int64) {
    c.Lock()
    r = c.n
    c.Unlock()
    return
}
```


We should change the receiver type of the `Value` method to the pointer type `*Counter` to avoid copying `sync.Mutex` values.

我们应该将 `Value` 方法的接收器类型更改为指针类型 `*Counter` 以避免复制 `sync.Mutex` 值。

The `go vet` command provided in Go Toolchain will report potential bad value copies.

Go 工具链中提供的 `go vet` 命令将报告潜在的错误值副本。

### Call the `sync.WaitGroup.Add` Method at Wrong Places

### 在错误的地方调用`sync.WaitGroup.Add` 方法

Each `sync.WaitGroup` value maintains a counter internally, The initial value of the counter is zero. If the counter of a `WaitGroup` value is zero, a call to the `Wait` method of the `WaitGroup` value will not block, otherwise, the call blocks until the counter value becomes zero. 

每个`sync.WaitGroup`值在内部维护一个计数器，计数器的初始值为零。如果 `WaitGroup` 值的计数器为零，则不会阻塞对 `WaitGroup` 值的 `Wait` 方法的调用，否则，调用会阻塞，直到计数器值变为零。

To make the uses of `WaitGroup` value meaningful, when the counter of a `WaitGroup` value is zero, the next call to the `Add` method of the `WaitGroup` value must happen before the next call to the `Wait` method of the `WaitGroup` value.

为了使 `WaitGroup` 值的使用有意义，当 `WaitGroup` 值的计数器为零时，下一次调用 `WaitGroup` 值的 `Add` 方法必须发生在下一次调用 `Wait` 方法之前`WaitGroup` 值。

For example, in the following program, the `Add` method is called at an improper place, which makes that the final printed number is not always `100`. In fact, the final printed number of the program may be an arbitrary number in the range `[0, 100)`. The reason is none of the `Add` method calls are guaranteed to happen before the `Wait` method call, which causes none of the `Done` method calls are guaranteed to happen before the `Wait` method call returns.

例如，在下面的程序中，在不正确的地方调用了`Add`方法，这使得最终打印的数字并不总是`100`。事实上，程序最终打印出来的数字可能是‘[0, 100)’范围内的任意数字。原因是没有一个 `Add` 方法调用保证在 `Wait` 方法调用之前发生，这导致没有一个 `Done` 方法调用保证在 `Wait` 方法调用返回之前发生。

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var wg sync.WaitGroup
    var x int32 = 0
    for i := 0;i < 100;i++ {
        go func() {
            wg.Add(1)
            atomic.AddInt32(&x, 1)
            wg.Done()
        }()
    }

    fmt.Println("Wait ...")
    wg.Wait()
    fmt.Println(atomic.LoadInt32(&x))
}
```




To make the program behave as expected, we should move the `Add` method calls out of the new goroutines created in the `for` loop, as the following code shown.

为了使程序按预期运行，我们应该将 `Add` 方法调用移出 `for` 循环中创建的新 goroutine，如下面的代码所示。

```go
...
    for i := 0;i < 100;i++ {
        wg.Add(1)
        go func() {
            atomic.AddInt32(&x, 1)
            wg.Done()
        }()
    }
...
```




### Use Channels as Futures/Promises Improperly

### 不正确地将渠道用作期货/承诺

From the article [channel use cases](https://go101.org/article/channel-use-cases.html), we know that some functions will return [channels as futures](https://go101.org/article/channel-use-cases.html#future-promise). Assume `fa` and `fb` are two such functions, then the following call uses future arguments improperly.

从文章[频道用例](https://go101.org/article/channel-use-cases.html)，我们知道有些函数会返回[频道作为期货](https://go101.org/article/channel-use-cases.html#future-promise)。假设`fa` 和`fb` 是两个这样的函数，那么下面的调用不正确地使用了future 参数。

```go
doSomethingWithFutureArguments(<-fa(), <-fb())
```


In the above code line, the generations of the two arguments are processed sequentially, instead of concurrently.

在上面的代码行中，两个参数的生成是按顺序处理的，而不是并发处理。

We should modify it as the following to process them concurrently.

我们应该将其修改如下以并发处理它们。

```go
ca, cb := fa(), fb()
doSomethingWithFutureArguments(<-ca, <-cb)
```




### Close Channels Not From the Last Active Sender Goroutine

### 关闭不是来自最后一个活动发送方 Goroutine 的通道

A common mistake made by Go programmers is closing a channel when there are still some other goroutines will potentially send values to the channel later. When such a potential send (to the closed channel) really happens, a panic might occur.

Go 程序员犯的一个常见错误是关闭一个通道，当还有一些其他 goroutines 可能会在稍后将值发送到通道时。当这种潜在的发送（到关闭的通道）真的发生时，可能会发生恐慌。

This mistake was ever made in some famous Go projects, such as [this bug](https://github.com/kubernetes/kubernetes/pull/45291/files?diff=split) and [this bug](https://github.com/kubernetes/kubernetes/pull/39479/files?diff=split) in the kubernetes project.

这个错误曾经出现在一些著名的 Go 项目中，例如 [这个 bug](https://github.com/kubernetes/kubernetes/pull/45291/files?diff=split) 和 [这个 bug](https://github.com/kubernetes/kubernetes/pull/39479/files?diff=split)在 kubernetes 项目中。

Please read [this article](https://go101.org/article/channel-closing.html) for explanations on how to safely and gracefully close channels.

请阅读 [这篇文章](https://go101.org/article/channel-closure.html) 以了解有关如何安全、优雅地关闭频道的说明。

### Do 64-bit Atomic Operations on Values Which Are Not Guaranteed to Be 8-byte Aligned

### 对不保证 8 字节对齐的值执行 64 位原子操作

Up to now (Go 1.17), the address of the value involved in a 64-bit atomic operation is required to be 8-byte aligned. Failure to do so may make the current goroutine panic. For the standard Go compiler, such failure can only [happen on 32-bit architectures](https://golang.org/pkg/sync/atomic/#pkg-note-BUG). Please read [memory layouts](https://go101.org/article/memory-layout.html) to get how to guarantee the addresses of 64-bit word 8-byte aligned on 32-bit OSes.

到目前为止（Go 1.17），64位原子操作中涉及的值的地址要求是8字节对齐的。不这样做可能会使当前的 goroutine 恐慌。对于标准的 Go 编译器，这种失败只能[发生在 32 位架构上](https://golang.org/pkg/sync/atomic/#pkg-note-BUG)。请阅读[内存布局](https://go101.org/article/memory-layout.html) 以了解如何在 32 位操作系统上保证 64 位字 8 字节对齐的地址。

### Not Pay Attention to Too Many Resources Are Consumed by Calls to the `time.After` Function

### 不注意调用`time.After`函数消耗了太多资源

The `After` function in the `time` standard package returns [a channel for delay notification](https://go101.org/article/channel-use-cases.html#timer). The function is convenient, however each of its calls will create a new value of the `time.Timer` type. The new created `Timer` value will keep alive in the duration specified by the passed argument to the `After` function. If the function is called many times in a certain period, there will be many alive `Timer` values accumulated so that much memory and computation is consumed.

`time` 标准包中的 `After` 函数返回 [延迟通知的通道](https://go101.org/article/channel-use-cases.html#timer)。该函数很方便，但是它的每次调用都会创建一个新的`time.Timer` 类型的值。新创建的 `Timer` 值将在传递给 `After` 函数的参数指定的持续时间内保持活动状态。如果在某个时间段内多次调用该函数，则会累积许多活动的“Timer”值，从而消耗大量内存和计算。

For example, if the following `longRunning` function is called and there are millions of messages coming in one minute, then there will be millions of `Timer` values alive in a certain small period (several seconds), even if most of these ` Timer` values have already become useless.

例如，如果调用以下 `longRunning` 函数，并且一分钟内有数百万条消息传入，那么在某个小时间段（几秒）内将有数百万条“Timer”值存活，即使其中大部分` Timer` 值已经变得毫无用处。

```go
import (
    "fmt"
    "time"
)

// The function will return if a message
// arrival interval is larger than one minute.
func longRunning(messages <-chan string) {
    for {
        select {
        case <-time.After(time.Minute):
            return
        case msg := <-messages:
            fmt.Println(msg)
        }
    }
}
```




To avoid too many `Timer` values being created in the above code, we should use (and reuse) a single `Timer` value to do the same job.

为避免在上述代码中创建过多的“Timer”值，我们应该使用（并重用）单个“Timer”值来完成相同的工作。

```go
func longRunning(messages <-chan string) {
    timer := time.NewTimer(time.Minute)
    defer timer.Stop()

    for {
        select {
        case <-timer.C: // expires (timeout)
            return
        case msg := <-messages:
            fmt.Println(msg)

            // This "if" block is important.
            if !timer.Stop() {
                <-timer.C
            }
        }

        // Reset to reuse.
        timer.Reset(time.Minute)
    }
}
```


Note, the `if` code block is used to discard/drain a possible timer notification which is sent in the small period when executing the second branch code block.

请注意，`if` 代码块用于丢弃/清除在执行第二个分支代码块时在短时间内发送的可能的计时器通知。

### Use `time.Timer` Values Incorrectly

### 错误地使用`time.Timer` 值

An idiomatic use example of `time.Timer` values has been shown in the last section. Some explanations:

`time.Timer` 值的惯用用法示例已在上一节中显示。一些解释：

-     the `Stop` method of a `*Timer` value returns `false` if the corresponding `Timer` value has already expired or been stopped. If the `Stop` method returns `false`, and we know the `Timer` value has not been stopped yet, then the `Timer` value must have already expired.
-     after a `Timer` value is stopped, its `C` channel field can only contain most one timeout notification.
-     we should take out the timeout notification, if it hasn't been taken out, from a timeout `Timer` value after the `Timer` value is stopped and before resetting and reusing the `Timer` value. This is the meaningfulness of the `if` code block in the example in the last section.

- 如果相应的“Timer”值已过期或已停止，则“*Timer”值的“Stop”方法将返回“false”。如果`Stop` 方法返回`false`，并且我们知道`Timer` 值还没有停止，那么`Timer` 值肯定已经过期。
- 一个`Timer`值停止后，其`C`通道字段最多只能包含一个超时通知。
- 我们应该在`Timer`值停止后，重置和重用`Timer`值之前，从超时`Timer`值中取出超时通知，如果它没有被取出。这就是上一节示例中的 if 代码块的意义所在。

The `Reset` method of a `*Timer` value must be called when the corresponding `Timer` value has already expired or been stopped, otherwise, a data race may occur between the `Reset` call and a possible notification send to the ` C` channel field of the `Timer` value.

当相应的“Timer”值已经过期或停止时，必须调用“*Timer”值的“Reset”方法，否则，在“Reset”调用和可能的通知发送到“ `Timer` 值的 C` 通道字段。

If the first `case` branch of the `select` block is selected, it means the `Timer` value has already expired, so we don't need to stop it, for the sent notification has already been taken out. However, we must stop the timer in the second branch to check whether or not a timeout notification exists. If it does exist, we should drain it before reusing the timer, otherwise, the notification will be fired immediately in the next loop step.

如果选择了 `select` 块的第一个 `case` 分支，则表示 `Timer` 值已经过期，所以我们不需要停止它，因为发送的通知已经被取出。但是，我们必须停止第二个分支中的计时器以检查是否存在超时通知。如果确实存在，我们应该在重用计时器之前将其排空，否则，将在下一个循环步骤中立即触发通知。

For example, the following program is very possible to exit in about one second, instead of ten seconds. And more importantly, the program is not data race free.

比如下面的程序很有可能在一秒左右退出，而不是十秒。更重要的是，该程序不是无数据竞争的。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    timer := time.NewTimer(time.Second/2)
    select {
    case <-timer.C:
    default:
        // Most likely go here.
        time.Sleep(time.Second)
    }
    // Potential data race in the next line.
    timer.Reset(time.Second * 10)
    <-timer.C
    fmt.Println(time.Since(start)) // about 1s
}
```


A `time.Timer` value can be leaved in non-stopping status when it is not used any more, but it is recommended to stop it in the end.

`time.Timer` 值可以在不再使用时保持非停止状态，但建议最后停止。

It is bug prone and not recommended to use a `time.Timer` value concurrently among multiple goroutines.

它很容易出错，不建议在多个 goroutine 之间同时使用一个 `time.Timer` 值。

We should not rely on the return value of a `Reset` method call. The return result of the `Reset` method exists just for compatibility purpose.

我们不应该依赖 `Reset` 方法调用的返回值。 `Reset` 方法的返回结果只是出于兼容性目的而存在。

------

[Index↡](https://go101.org/article/concurrent-common-mistakes.html#i-concurrent-common-mistakes.html)

[索引↡](https://go101.org/article/concurrent-common-mistakes.html#i-concurrent-common-mistakes.html)

------

The ***Go 101\*** project is hosted on [Github](https://github.com/go101/go101). Welcome to improve ***Go 101\*** articles by submitting corrections for all kinds of mistakes, such as typos, grammar errors, wording inaccuracies, description flaws, code bugs and broken links. If you would like to learn some Go details and facts every serveral days, please follow Go 101's official Twitter account: [@go100and1](https://twitter.com/go100and1). 

***Go 101\*** 项目托管在 [Github](https://github.com/go101/go101) 上。欢迎改进 ***Go 101\*** 文章，提交各种错误的更正，例如拼写错误、语法错误、措辞不准确、描述缺陷、代码错误和断开的链接。如果你想每隔几天了解一些围棋的细节和事实，请关注围棋 101 的官方推特账号：[@go100and1](https://twitter.com/go100and1)。

The digital versions of this book are available at the following places:  [Leanpub store](https://leanpub.com/go101), *$19.99+*, Leanpub gets 20%, Tapir gets 80%. [Apple Books store](https://books.apple.com/us/book/id1459984231), *$19.99*, Apple gets 30%, Tapir gets 70%. [Amazon Kindle store](https://www.amazon.com/dp/B07Q3HWZ98), *$39.99*, Amazon gets 65%, Tapir gets 35%. [Free ebooks](https://github.com/go101/go101/releases), including pdf, epub and azw3 formats. Tapir, the author of Go 101, has spent 4+ years on writing the Go 101 book and maintaining the go101.org website. New contents will continue being added to the book and the website from time to time. Tapir is also an indie game developer. You can also support Go 101 by playing [Tapir's games](https://www.tapirgames.com) (made for both Android and iPhone/iPad):  [Color Infection](https://www.tapirgames.com/App/Color-Infection) (★★★★★), a physics based original casual puzzle game. 140+ levels. [Rectangle Pushers](https://www.tapirgames.com/App/Rectangle-Pushers)(★★★★★), an original casual puzzle game. Two modes, 104+ levels. [Let's Play With Particles](https://www.tapirgames.com/App/Let-Us-Play-With-Particles), a casual action original game. Three mini games are included. Individual donations [via PayPal](https://paypal.me/tapirliu) are also welcome.

本书的数字版本可在以下位置获得：[Leanpub 商店](https://leanpub.com/go101)，*$19.99+*，Leanpub 获得 20%，Tapir 获得 80%。 [Apple Books 商店](https://books.apple.com/us/book/id1459984231)，*$19.99*，Apple 获得 30%，Tapir 获得 70%。 [亚马逊 Kindle 商店](https://www.amazon.com/dp/B07Q3HWZ98)，*$39.99*，亚马逊获得 65%，貘获得 35%。 [免费电子书](https://github.com/go101/go101/releases)，包括pdf、epub和azw3格式。 Tapir 是 Go 101 的作者，在编写 Go 101 书籍和维护 go101.org 网站上花费了 4 年多的时间。本书和网站将不时添加新内容。 Tapir 也是独立游戏开发商。您还可以通过玩 [Tapir's games](https://www.tapirgames.com)（适用于 Android 和 iPhone/iPad）来支持 Go 101：[Color Infection](https://www.tapirgames.com/App/Color-Infection）（★★★★★)，一款基于物理的原创休闲益智游戏。 140 多个级别。  [矩形推手](https://www.tapirgames.com/App/Rectangle-Pushers)(★★★★★)，一款原创休闲益智游戏。两种模式，104+ 级。  [让我们玩粒子游戏](https://www.tapirgames.com/App/Let-Us-Play-With-Particles)，一款休闲动作原创游戏。包括三个小游戏。也欢迎个人捐款 [通过 PayPal](https://paypal.me/tapirliu)。

------

Index:

指数：

- [About Go 101](https://go101.org/article/101-about.html) - why this book is written.
- [Acknowledgements](https://go101.org/article/acknowledgements.html) 

- [关于 Go 101](https://go101.org/article/101-about.html) - 为什么写这本书。
- [致谢](https://go101.org/article/acknowledgements.html)

