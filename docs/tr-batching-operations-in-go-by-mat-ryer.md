# Batching operations in Go.

# Go 中的批处理操作。

13 Feb 2020

2020 年 2 月 13 日

### The problem

###  问题

Say we have a remote call to make, and we want to break a list of items into smaller chunks to send them in batches.

假设我们有一个远程调用，我们想将一个项目列表分成更小的块以批量发送它们。

It’s a fairly simple problem to solve, but we’ll look at how a well  designed helper can make all the difference to the readability and  stability of our code.

这是一个相当容易解决的问题，但我们将看看一个设计良好的帮助程序如何对我们代码的可读性和稳定性产生重大影响。

### The situation

### 情况

We always like to think of the situation the user is in when we are  designing any kind of API, and this applies to functions too.

我们在设计任何类型的 API 时总是喜欢考虑用户所处的情况，这也适用于函数。

Say we have a slice of items that we want to process:

假设我们有一部分要处理的项目：

```go
items, err := getAllItemsFromRequest(r)
if err != nil {
    return errors.Wrap(err, "getAllItemsFromRequest")
}
```

It’s possible that the array contains thousands of items. But say we  only want to process them in batches of ten, it would be nice to be able to call a method like this:

数组可能包含数千个项目。但是假设我们只想以 10 为一组处理它们，能够调用这样的方法会很好：

```go
batchSize := 10
err := batch(len(items), batchSize, func(start, end int) error {
    batchItems := items[start:end]
    if err := performSomeRemoteThing(ctx, batchItems);err != nil {
        return errors.Wrap(err, "performSomeRemoteThing")
    }
})
if err != nil {
    return err
}
```

The `batch` function takes the total number of items, the batch (page) size, and a function that gets called for each batch, with `start` and `end` marking the range, allowing us to re-slice the items:

`batch` 函数获取项目的总数、批次（页面）大小，以及为每个批次调用的函数，用 `start` 和 `end` 标记范围，允许我们重新切片项目：

```go
batchItems := items[start:end]
```

In this example, if we got `105` items, the `performSomeRemoteThing` function would get called eleven times, each time with a different page of `10` items (the `batchSize`) except the last time, when it would be a slice of the remaining five items.

在这个例子中，如果我们得到了 105 个项目，performSomeRemoteThing 函数将被调用 11 次，每次都有不同的页面 10 个项目（“batchSize”），除了最后一次，当它是一个剩下的五个项目的切片。

### Start with tests

### 从测试开始

When solving problems like this, I find TDD to be an excellent guide  and check of what I’m doing. It is especially good at confirming we  don’t have any off-by-one errors, or hit any snags at the edges.

在解决这样的问题时，我发现 TDD 是一个很好的指南，可以检查我在做什么。它特别擅长确认我们没有任何一一的错误，或者在边缘击中任何障碍。

Consider the following test code:

考虑以下测试代码：

```go
func Test(t *testing.T) {
    is := is.New(t)

    type r struct {
        start, end int
    }
    var ranges []r
    err := batch(100, 10, func(start, end int) error {
        ranges = append(ranges, r{
            start: start,
            end:   end,
        })
        return nil
    })
    is.NoErr(err)

    is.Equal(len(ranges), 10)
    is.Equal(ranges[0].start, 0)
    is.Equal(ranges[0].end, 9)
    is.Equal(ranges[1].start, 10)
    is.Equal(ranges[1].end, 19)
    is.Equal(ranges[2].start, 20)
    is.Equal(ranges[2].end, 29)
    is.Equal(ranges[3].start, 30)
    is.Equal(ranges[3].end, 39)
    is.Equal(ranges[4].start, 40)
    is.Equal(ranges[4].end, 49)
    is.Equal(ranges[5].start, 50)
    is.Equal(ranges[5].end, 59)
    is.Equal(ranges[6].start, 60)
    is.Equal(ranges[6].end, 69)
    is.Equal(ranges[7].start, 70)
    is.Equal(ranges[7].end, 79)
    is.Equal(ranges[8].start, 80)
    is.Equal(ranges[8].end, 89)
    is.Equal(ranges[9].start, 90)
    is.Equal(ranges[9].end, 99)
}
```

- This code uses the [github.com/matryer/is](https://github.com/matryer/is) test helper (like Testify off steroids) but it should be obvious enough to read.

- 此代码使用 [github.com/matryer/is](https://github.com/matryer/is) 测试助手（如 Testify off steroids)，但它应该足够明显可读。

The test uses the `batch` function, and appends the details of each call to a `ranges` slice.

该测试使用`batch` 函数，并将每次调用的详细信息附加到`ranges` 切片。

After the process, we check that the `err` was `nil`, and then make assertions about all the indexes we expect.

在此过程之后，我们检查 `err` 是否为 `nil`，然后对我们期望的所有索引进行断言。

- I’ll leave it up to you to see if you can tidy up this test in some way? Would a table test style be more appropriate?

- 我会让你来看看你是否可以以某种方式整理这个测试？表格测试风格会更合适吗？

What’s nice about how explicit this is, is that we can actually think about and check each value. We know exactly what the `start` and `end` values should be, so we can spell it out.

如此明确的好处在于，我们实际上可以考虑并检查每个值。我们确切地知道 `start` 和 `end` 值应该是什么，所以我们可以拼写出来。

This is easier to reason about than the upcoming looping and counting logic we’re about to write.

这比我们即将编写的循环和计数逻辑更容易推理。

### The `batch` function

### `batch` 函数

Our `batch` function is going to keep a start index `i` (starting at `0` the first item) and call the `eachFn(i, end)` for each batch, passing in the start and end indexes. Each iteration, `i` is recalculated to be the next item: `end + 1`.

我们的 `batch` 函数将保留一个起始索引 `i`（从第一项的 `0` 开始）并为每个批次调用 `eachFn(i, end)`，传入开始和结束索引。每次迭代，`i` 被重新计算为下一项：`end + 1`。

```go
// batch calls eachFn for all items up to count.
// Returns any error from eachFn except for Abort it returns nil.
func batch(count, batchSize int, eachFn func(start, end int) error) error {
    i := 0
    for i < count {
        end := i + batchSize - 1
        if end > count-1 {
            // passed end, so set to end item
            end = count - 1
        }
        err := eachFn(i, end)
        if err == Abort {
            return nil
        }
        if err != nil {
            return err
        }
        i = end + 1
    }
    return nil
}
```

#### Aborting

#### 中止

In the `batch` function above, you can see that we check for a special sentinel error ([coined by Dave Cheney](https://dave.cheney.net/tag/errors)) called `Abort`:

在上面的 `batch` 函数中，您可以看到我们检查了一个名为 `Abort` 的特殊标记错误（[由 Dave Cheney 创造](https://dave.cheney.net/tag/errors))：

```go
if err == Abort {
    return nil
}
```

If the `err` returned from the callback is `Abort` the `batch` function stops iterating and returns `nil`, indicating a happy exit.

如果回调返回的 `err` 是 `Abort`，`batch` 函数将停止迭代并返回 `nil`，表示顺利退出。

The `Abort` variable can be declared like this:

`Abort` 变量可以这样声明：

```go
// Abort is a sentinel error which indicates a batch
// operation should abort early.
var Abort = errors.New("abort")
```

#### Type the callback func

#### 输入回调函数

Rather than define the callback signature `func(start, end int) error` inline, it’s better to declare a type.

与其在行内定义回调签名 `func(start, end int) error`，不如声明一个类型。

This allows you to document the callback, and how it should be used.

这允许您记录回调以及它应该如何使用。

```go
// BatchFunc is called for each batch.
// Any error will cancel the batching operation but returning Abort
// indicates it was deliberate, and not an error case.
type BatchFunc func(start, end int) error
```

The comment should say *everything* about the behaviour of this callback.

评论应该说明关于这个回调的行为的*一切*。

## Conclusion

##  结论

We recommend you just copy the code for this function (and its test)  to your own project (even if you end up having a couple of copies of it, what’s the harm in that?).

我们建议您只将此函数（及其测试）的代码复制到您自己的项目中（即使您最终拥有它的几个副本，这样做有什么危害？）。

This package is maintained as a Go module over at https://github.com/pacedotdev/batch.

这个包作为 Go 模块在 https://github.com/pacedotdev/batch 上维护。

The mechanics are [fairly simple](https://github.com/pacedotdev/batch/blob/master/batch.go), but the code is encapsulated and well tested.

机制[相当简单](https://github.com/pacedotdev/batch/blob/master/batch.go)，但代码被封装并经过良好测试。

------

## Learn more about what we're doing at Pace.

## 详细了解我们在 Pace 的工作。

A lot of our blog posts come out of the technical work behind a project we're working on called Pace.

我们的很多博客文章都来自我们正在进行的名为 Pace 的项目背后的技术工作。

We were frustrated by communication and project management tools that interrupt your flow and overly complicated workflows turn simple tasks, hard. So we decided to build Pace.

我们对中断您的流程的沟通和项目管理工具感到沮丧，过于复杂的工作流程使简单的任务变得困难。所以我们决定建立 Pace。

Pace is a new minimalist project management tool for tech teams. We promote **asynchronous communication** by default, while allowing for those times when you really need to chat.

Pace 是面向技术团队的全新极简项目管理工具。我们默认提倡**异步通信**，同时允许您真正需要聊天的时候。

We shift the way work is assigned by allowing only **self-assignment**, creating a more empowered team and protecting the attention and focus of devs.

我们通过只允许**自我分配**来改变工作分配方式，创建一个更有权力的团队并保护开发人员的注意力和注意力。

We're currently live and would love you to try it and share your opinions on what project management tools should and shouldn't do.

我们目前正在直播，希望您尝试一下，并分享您对项目管理工具应该做什么和不应该做什么的看法。

**What next?** [Start your 14 day free trial to see if Pace is right for your team](https://pace.dev/) 

**接下来怎么办？** [开始 14 天免费试用，看看 Pace 是否适合您的团队](https://pace.dev/)

