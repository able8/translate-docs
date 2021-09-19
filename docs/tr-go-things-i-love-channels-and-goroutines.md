## Go Things I Love: Channels and Goroutines

## 去我喜欢的东西：Channels 和 Goroutines

Mon Jan 06, 2020 by Justin Fuller

贾斯汀富勒 2020 年 1 月 6 日星期一

Justin Fuller is a Software Engineer at [The New York Times](https://open.nytimes.com). He works with Go, JavaScript, Node.js, and React.

Justin Fuller 是 [The New York Times](https://open.nytimes.com) 的一名软件工程师。他使用 Go、JavaScript、Node.js 和 React。

This series, *Go Things I Love*, is my attempt to show the parts of Go that I like the best, as well as why I love working with it at [The New York Times](https://open.nytimes.com).

这个系列，*Go Things I Love*，是我尝试展示 Go 中我最喜欢的部分，以及为什么我喜欢在 [纽约时报](https://open.nytimes.com)。

In my last post [Go Things I Love: Methods On Any Type](https://www.justindfuller.com/2019/12/go-things-i-love-methods-on-any-type/), I demonstrated a feature of Go that makes it easy to build Object-Oriented software.

在我上一篇文章 [Go Things I Love: Methods On Any Type](https://www.justindfuller.com/2019/12/go-things-i-love-methods-on-any-type/) 中，我演示了Go 的一个特性，使构建面向对象的软件变得容易。

This post, *Channels and Goroutines*, will demonstrate a few neat concurrency patterns in Go.

这篇文章 *Channels 和 Goroutines*，将演示 Go 中的一些简洁的并发模式。

![Go Things I Love](https://www.justindfuller.com/go-things-i-love.png)

First: to get the most out of this post you should familiarize yourself with  the fundamentals of Go concurrency. A great place to do that is [in the Go tour](https://tour.golang.org/concurrency/1). These patterns rely on goroutines and channels to accomplish their elegance.

首先：要充分利用这篇文章，您应该熟悉 Go 并发的基础知识。 [在 Go 之旅中](https://tour.golang.org/concurrency/1) 是这样做的好地方。这些模式依靠 goroutine 和通道来实现它们的优雅。

Concurrency, in some form, is one of the most important building blocks of  performant software. That's why it's important to pick a programming  language with first-class concurrency support. Because Go, in my  estimation, provides one of the most delightful ways to achieve  concurrency, I believe it is a solid choice for any project that  involves concurrency.

在某种形式下，并发是高性能软件最重要的构建块之一。这就是为什么选择具有一流并发支持的编程语言很重要的原因。因为在我看来，Go 提供了实现并发的最令人愉快的方法之一，我相信它是任何涉及并发的项目的可靠选择。

## First Class

## 头等舱

To be  first-class is to have full support and consideration in all things. That means, to be first-class, concurrency must be a part of the Go  language itself. It cannot be a library bolted on the side.

一流，就是凡事都有充分的支持和考虑。这意味着，要成为一流的，并发必须是 Go 语言本身的一部分。它不能是一个用螺栓固定在侧面的图书馆。

A few type declarations will begin to show how concurrency is built into the language.

一些类型声明将开始展示并发是如何构建到语言中的。

```go
type (
    WriteOnly(chan<- int)
    ReadOnly(<-chan int)
    ReadAndWrite(chan int)
)
```

Notice the `chan` keyword in the function argument definitions. A `chan` is a channel.

注意函数参数定义中的 `chan` 关键字。 `chan` 是一个频道。

Next comes the arrow `<-` that shows which way the data flow to or from the channel. The `WriteOnly` function receives a channel that can only be written to. The `ReadOnly` function receives a channel that can only be read from.

接下来是箭头“<-”，它显示了数据流入或流出通道的方式。 `WriteOnly` 函数接收一个只能写入的通道。 `ReadOnly` 函数接收一个只能读取的通道。

Being able to declare the flow of the data to a channel is an important way  in which channels are first-class members of the Go programming  language. Channel flow is important because it's how goroutines  communicate.

能够声明数据流向通道是通道成为 Go 编程语言一流成员的重要方式。通道流很重要，因为它是 goroutine 通信的方式。

It's directly related to this phrase you might have seen before:

它与您之前可能见过的这个短语直接相关：

> Do not communicate by sharing memory; instead, share memory by communicating.

> 不要通过共享内存进行通信；相反，通过通信共享内存。

The phrase, “share memory by communicating”, means goroutines should  communicate changes through channels; they provide a safer, idiomatic  way to share memory.

“通过通信共享内存”这句话意味着 goroutine 应该通过通道来传达变化；它们提供了一种更安全、惯用的方式来共享内存。

## Communicating by sharing memory (👎)

## 通过共享内存进行通信 (👎)

Here's an example of Go code that communicates by sharing memory.

这是通过共享内存进行通信的 Go 代码示例。

```go
func IntAppender() {
    var ints []int
    var wg sync.WaitGroup

    for i := 0;i < 10;i++ {
        wg.Add(1)

        go func(i int) {
            defer wg.Done()
            ints = append(ints, i)
        }(i)
    }

    wg.Wait()
}
```

`IntAppender` creates a goroutine for  each integer that is appended to the array. Even though it's a little  too trivial to be realistic, it still serves an important demonstrative  purpose.

`IntAppender` 为附加到数组的每个整数创建一个 goroutine。尽管它有点过于琐碎而不切实际，但它仍然具有重要的示范作用。

In `IntAppender` each goroutine shares the same memory—the `ints` array—which it appends integers to.

在 `IntAppender` 中，每个 goroutine 共享相同的内存 - `ints` 数组 - 它向其中追加整数。

This code communicates by sharing memory. Yes, it seems like works (only if  you run it on the go playground)—but it's not idiomatic Go. More  importantly, it's not a safe way to write this program because it  doesn't always give the expected results (again, unless you run it on  the go playground).

此代码通过共享内存进行通信。是的，它看起来很有效（只有当你在 go playground 上运行它）——但它不是 Go 的地道。更重要的是，编写这个程序并不是一种安全的方式，因为它并不总是给出预期的结果（同样，除非你在 go playground 上运行它）。

It's not safe because there are 11 goroutines  (one running the main function and ten more spawned by the loop) with  access to the `ints` slice.

这是不安全的，因为有 11 个 goroutines（一个运行主函数，另外 10 个由循环产生）可以访问 `ints` 切片。

This pattern provides no guarantee that the program will behave as expected; anything can happen when memory is shared broadly.

这种模式不能保证程序会按预期运行；当内存被广泛共享时，任何事情都可能发生。

## Share memory by communicating (👍)

## 通过交流分享记忆（👍）

The first sign that this example is not following “share memory by communicating” is the use of `sync.WaitGroup`. Even though I consider WaitGroups to be a code smell, I'm not ready to  claim they are always bad. Either way, code is usually safer with a  channel.

这个例子没有遵循“通过通信共享内存”的第一个迹象是使用了`sync.WaitGroup`。尽管我认为 WaitGroups 是一种代码异味，但我还没有准备好声称它们总是很糟糕。无论哪种方式，使用通道的代码通常更安全。

Let's convert the bad example to idiomatic Go by replacing the `WaitGroup` with a channel.

让我们通过用通道替换 `WaitGroup` 来将坏示例转换为惯用的 Go。

```go
// WriteOnly serves the purpose of demonstrating
// a method that writes to a write-only channel.
func WriteOnly(channel chan<-int, order int) {
    channel <- order
}

func main() {
    var ints []int
    channel := make(chan int, 10)

    for i := 0;i < 10;i++ {
        go WriteOnly(channel, i)
    }

    for i := range channel {
        ints = append(ints, i)

        if len(ints) == 10 {
            break
        }
    }

    fmt.Printf("Ints %v", ints)
}
```

[See this example in the Go playground.](https://play.golang.org/p/gi8zyZH7KMd)

[在 Go 游乐场中查看此示例。](https://play.golang.org/p/gi8zyZH7KMd)

Now, only one goroutine can modify the `ints` slice while the rest communicate through a channel. They're sharing  memory by communicating through a channel instead of modifying shared  memory.

现在，只有一个 goroutine 可以修改 `ints` 切片，而其余的则通过通道进行通信。它们通过通道通信而不是修改共享内存来共享内存。

The example here shows two important ways that concurrency (goroutines and channels) are first-class citizens of the Go  programming language. First, we used a write-only channel argument. This guaranteed that the method won't accidentally read from the channel,  unexpectedly altering the functionality. Second, we see that the `for range` loop works on channels.

这里的例子展示了并发（goroutines 和通道）是 Go 编程语言的一等公民的两种重要方式。首先，我们使用了只写通道参数。这保证了该方法不会意外地从通道读取，意外地改变功能。其次，我们看到 `for range` 循环适用于通道。

These are just a few ways that Go makes concurrency a first-class citizen. Next, let's see what we can accomplish with goroutines and channels.

这些只是 Go 使并发成为一等公民的几种方式。接下来，让我们看看我们可以用 goroutines 和 channel 来完成什么。

## Timeout

##  暂停

To demonstrate a timeout, we will construct a simple news UI backend that fetches results from three [New York Times endpoints](https://developer.nytimes.com/). Even though the NYT endpoints respond very quickly, this won't quite  meet our standards. Our program must always respond within 80  milliseconds. Because of this restriction, we're only going to use NYT  endpoint responses that come fast enough.

为了演示超时，我们将构建一个简单的新闻 UI 后端，从三个[纽约时报端点](https://developer.nytimes.com/) 获取结果。即使 NYT 端点响应非常快，这也不太符合我们的标准。我们的程序必须始终在 80 毫秒内响应。由于此限制，我们将仅使用速度足够快的 NYT 端点响应。

Here are the URLs that the program will fetch from:

以下是程序将从中获取的 URL：

```go
var urls = [...]string{
    "https://api.nytimes.com/svc/topstories/v2/home.json",
    "https://api.nytimes.com/svc/mostpopular/v2/viewed/1.json",
    "https://api.nytimes.com/svc/books/v3/lists/current/hardcover-fiction.json",
}
```

The URLs have been declared as an array of strings, which will allow them to be iterated.

URL 已被声明为一个字符串数组，这将允许它们被迭代。

Another neat feature of Go is how you can declare `const` blocks. Like this:

Go 的另一个巧妙功能是如何声明 `const` 块。像这样：

```go
const (
    urlTopStories              = "https://api.nytimes.com/svc/topstories/v2/home.json"
    urlMostPopular             = "https://api.nytimes.com/svc/mostpopular/v2/viewed/1.json"
    urlHardcoverFictionReviews = "https://api.nytimes.com/svc/books/v3/lists/current/hardcover-fiction.json"
)
```

Now, the `urls` array can be more expressive by using the const declarations.

现在，通过使用 const 声明，`urls` 数组可以更具表现力。

```go
var urls = [...]string{
    urlTopStories,
    urlMostPopular,
    urlHardcoverFictionReviews,
}
```

The URLs are for top stories, most popular stories, and the current hardcover fiction reviews.

这些 URL 用于热门故事、最受欢迎的故事和当前的精装小说评论。

Instead of a real `http.Get` I will substitute a fake `fetch` function. This will provide a clearer demonstration of the timeout.

我将代替一个真正的 `http.Get`，而是一个假的 `fetch` 函数。这将提供更清晰的超时演示。

```go
func fetch(url string, channel chan<- string) {
    source := rand.NewSource(time.Now().UnixNano())
    random := rand.New(source)
    duration := time.Duration(random.Intn(150)) * time.Millisecond
    time.Sleep(duration)
    channel <- url
}
```

This is a common pattern in Go demonstration  code—generate a random number, sleep the goroutine for the randomly  generated duration, then do some work. To fully understand why this code is being used to demonstrate a fake `http.Get`, the next sections will step through each line, explaining what it does.

这是 Go 演示代码中的常见模式——生成一个随机数，在随机生成的持续时间内休眠 goroutine，然后做一些工作。为了完全理解为什么使用这段代码来演示一个假的 `http.Get`，接下来的部分将逐步完成每一行，解释它的作用。

### Deterministic Randomness (See: oxymorons)

### 确定性随机性（参见：矛盾修辞法）

In Go, the random number generator is, by default, deterministic.

在 Go 中，随机数生成器默认是确定性的。

> In mathematics, computer science and physics, a deterministic system is a  system in which no randomness is involved in the development of future  states of the system. - [The Encyclopedia of Science](https://www.daviddarling.info/encyclopedia/D/deterministic_system.html)

> 在数学、计算机科学和物理学中，确定性系统是指在系统未来状态的发展中不涉及随机性的系统。 - [科学百科全书](https://www.daviddarling.info/encyclopedia/D/deterministic_system.html)

This means that we have to seed the randomizer with something that changes; if not, the randomizer will always produce the same value. So we create a source, typically based on the current time.

这意味着我们必须为随机化器设置一些变化的种子；如果不是，随机化器将始终产生相同的值。所以我们创建一个源，通常基于当前时间。

```go
source := rand.NewSource(time.Now().UnixNano())
```

After the source is created, it can be used to  create a random number generator. We must create the source and random  generator each time. Otherwise, it will continue to return the same  number.

源创建后，可用于创建随机数生成器。我们每次都必须创建源和随机生成器。否则，它将继续返回相同的数字。

```go
random := rand.New(source)
```

Once the generator is created, it can be used to create a random number between 0 and 150. That random number is converted to a `time.Duration` type, then multiplied to become milliseconds.

一旦生成器被创建，它就可以用来创建一个 0 到 150 之间的随机数。这个随机数被转换为一个 `time.Duration` 类型，然后乘以成为毫秒。

```go
duration := time.Duration(random.Intn(150)) * time.Millisecond
```

One further note about the randomness is needed. It will always return the same value in the go playground because the go playground always starts running with the same timestamp. So, if you  plug this into the playground, you'll always receive the same result. If you want to see the timeout in action, just replace 150 with some  number below 80.

需要进一步说明随机性。它将始终在 go playground 中返回相同的值，因为 go playground 总是以相同的时间戳开始运行。因此，如果您将其插入操场，您将始终收到相同的结果。如果您想查看超时时间，只需将 150 替换为低于 80 的某个数字。

### Another send-only channel

### 另一个仅发送通道

At the very bottom of `fetch` are the two lines that we care about.

在 `fetch` 的最底部是我们关心的两行。

```go
time.Sleep(duration)
channel <- url
```

The first line tells the goroutine to sleep for  the specified duration. This will make some responses take too long for  the given URL, later causing the API to respond without the results of  that URL.

第一行告诉 goroutine 在指定的时间内休眠。这将使给定 URL 的某些响应时间过长，稍后导致 API 响应而没有该 URL 的结果。

Finally, the URL is sent to the channel. In a real `fetch` it would be expected that the actual response is sent to the channel. For our purposes, it's just the URL.

最后，将 URL 发送到频道。在真正的 `fetch` 中，实际响应会被发送到通道。对于我们的目的，它只是 URL。

### A read-only channel

### 一个只读通道

Since the `fetch` function funnels results in the channel, it makes sense to have a  corresponding function funnel results from the channel into a slice of  strings.

由于 `fetch` 函数漏斗结果在通道中，因此将相应的函数漏斗从通道结果转换为字符串切片是有意义的。

Take a look at the function. Next, we'll break it down line-by-line.

看一下功能。接下来，我们将逐行分解它。

```go
func stringSliceFromChannel(maxLength int, input <-chan string) []string {
    var results []string
    timeout := time.After(time.Duration(80) * time.Millisecond)

    for {
        select {
        case str := <-input:
            results = append(results, str)

            if len(results) == maxLength {
                fmt.Println("Got all results")
                return results
            }
        case <-timeout:
            fmt.Println("Timeout!")
            return results
        }
    }
}
```

First, look at the function argument declaration.

首先，查看函数参数声明。

```go
func stringSliceFromChannel(maxLength int, input <-chan string) []string {
```

The `stringSliceFromChannel` function declares that it will accept a read-only channel, `channel <-chan string`. This indicates that the function will convert the channel's inputs into a different type of output—a slice of strings, or `[]string`.

`stringSliceFromChannel` 函数声明它将接受一个只读通道，`channel <-chan string`。这表明该函数会将通道的输入转换为不同类型的输出——字符串切片，或“[]string”。

Even though it's valid to declare a function argument with, `channel chan string`, opting for the arrow `<-` operator makes the function's intent clearer. This can be particularly helpful in a long function.

尽管使用 `channel chan string` 声明函数参数是有效的，但选择箭头 `<-` 运算符会使函数的意图更加清晰。这在长函数中特别有用。

Next, the timeout is created.

接下来，创建超时。

```go
timeout := time.After(time.Duration(80) * time.Millisecond)
```

The function `time.After` returns a channel. After the given `time.Duration` it will write to the channel (*what* it writes doesn't matter).

函数`time.After` 返回一个通道。在给定的`time.Duration` 之后，它将写入通道（*它写入的内容无关紧要）。

Moving on, the `timeout` and `input` channels are used together in a `for select` loop.

继续，`timeout` 和 `input` 通道在 `for select` 循环中一起使用。

The `for` loop with no other arguments will loop forever until stopped by a `break` or `return`.

没有其他参数的 `for` 循环将永远循环，直到被 `break` 或 `return` 停止。

The `select` acts as a `switch` statement for channels. The first `case` block to have a channel ready will execute.

`select` 充当通道的 `switch` 语句。将执行第一个准备好通道的“case”块。

By combining the `for` and `select`, this block of code will run until the desired number of results is retrieved or until the timeout happens.

通过组合 `for` 和 `select`，这段代码将一直运行，直到检索到所需数量的结果或直到发生超时。

Take a look at the case block for the `input` channel.

查看“输入”通道的 case 块。

```go
case str := <-input:
    results = append(results, str)
    
    if len(results) == maxLength {
        fmt.Println("Got all results")
        return results
    }
```

The output of the channel is assigned to a variable, `str`. Next, `str` is appended to the results array. The results array is returned if it is the desired length.

通道的输出被分配给一个变量，`str`。接下来，`str` 被附加到结果数组中。如果它是所需的长度，则返回结果数组。

Now, look at the case block for the `timeout` channel.

现在，查看“超时”通道的 case 块。

```go
case <-timeout:
    fmt.Println("Timeout!")
    return results
```

Whatever results are available, even if there are none, will be returned when the timeout happens.

任何可用的结果，即使没有，也会在超时发生时返回。

------

👋 Want to learn more about Go? [Subscribe to my newsletter](https://justindfuller.us4.list-manage.com/subscribe?u=d48d0debd8d0bce3b77572097&id=0c1e610cac) to get an update, once-per-month, about what I'm writing about.

👋想了解更多关于围棋的知识吗？ [订阅我的时事通讯](https://justindfuller.us4.list-manage.com/subscribe?u=d48d0debd8d0bce3b77572097&id=0c1e610cac) 每月一次获取有关我正在撰写的内容的更新。

------

## The Main Function

## 主要功能

Now there is both a channel writer and a channel reader. Let's see how to put it all together in the `main` function.

现在有一个频道编写器和一个频道阅读器。让我们看看如何在 `main` 函数中将它们组合在一起。

```go
func main() {
    channel := make(chan string)
    for _, url := range urls {
        go fetch(url, channel)
    }

    results := stringSliceFromChannel(len(urls), channel)

    fmt.Printf("Results: %v\n", results)
}
```

First, a channel is created to collect the fetch results, `channel := make(chan string)`.

首先，创建一个通道来收集获取结果，`channel := make(chan string)`。

Next, the `urls` are looped over, creating a goroutine to fetch each url.

接下来，`urls` 被循环，创建一个 goroutine 来获取每个 url。

```go
for _, url := range urls {
    go fetch(url, channel)
}
```

This allows the fetching to happen concurrently. 

这允许获取同时发生。

After the fetches have been kicked off, `stringSliceFromChannel` will block until either the results are in or the timeout occurs.

提取开始后，`stringSliceFromChannel` 将阻塞，直到结果出现或超时发生。

```go
results := stringSliceFromChannel(len(urls), channel)
```

Finally, we can print the results to see which URLs are returned. If you run this code in the [Go Playground](https://play.golang.org/p/g3RnP9A26v5), remember to change the timeout number since the random number generator will always return the same results.

最后，我们可以打印结果以查看返回了哪些 URL。如果您在 [Go Playground](https://play.golang.org/p/g3RnP9A26v5) 中运行此代码，请记住更改超时数，因为随机数生成器将始终返回相同的结果。

## Caveats

## 注意事项

It could seem like I'm suggesting that you should always use channels  instead of waitgroups or mutexes. I'm not. Each tool is designed for a  specific use case, [and each has a tradeoff](https://github.com/golang/go/wiki/MutexOrChannel). Instead of walking away from this post thinking, “I should always use  channels, they're so much better than anything else.” I hope you will  simply consider if you can improve the clarity of your program with a  channel, rather than sharing memory. If not, don't use them.

我似乎在建议您应该始终使用通道而不是等待组或互斥锁。我不是。每个工具都是为特定用例设计的，[每个工具都有一个权衡](https://github.com/golang/go/wiki/MutexOrChannel)。不要放弃这篇文章，而是想：“我应该总是使用渠道，它们比其他任何东西都要好得多。”我希望您可以简单地考虑是否可以通过频道来提高节目的清晰度，而不是共享内存。如果没有，请不要使用它们。

## Final Thoughts

##  最后的想法

Here's the cool thing. We started out talking about how Go has first-class  concurrency support with goroutines and channels. Then we saw how easy  it is to implement a complex concurrent pattern, a timeout, with a  single channel and a few goroutines. Over my next few posts, I hope to  show that this was only scratching the surface of what one can do with  concurrency in Go. I hope you'll check back in. (Better yet, [subscribe to my newsletter](https://justindfuller.us4.list-manage.com/subscribe?u=d48d0debd8d0bce3b77572097&id=0c1e610cac) to be updated each month about my new posts)

这是很酷的事情。我们开始讨论 Go 如何通过 goroutines 和通道获得一流的并发支持。然后我们看到了实现复杂的并发模式、超时、单个通道和几个 goroutine 是多么容易。在我接下来的几篇文章中，我希望表明这只是在 Go 中使用并发可以做的事情的皮毛。我希望你能回来查看。（更好的是，[订阅我的时事通讯](https://justindfuller.us4.list-manage.com/subscribe?u=d48d0debd8d0bce3b77572097&id=0c1e610cac)每个月都会更新我的新帖子)

Finally, even though this is a neat concurrency pattern, it's unrealistic. As an exercise you could open the [Go Playground](https://play.golang.org/p/g3RnP9A26v5) to see if you can implement these scenarios:

最后，尽管这是一个简洁的并发模式，但它是不现实的。作为练习，您可以打开 [Go Playground](https://play.golang.org/p/g3RnP9A26v5) 看看您是否可以实现这些场景：

- The results should be returned as a JSON object. Maybe we could use a struct instead of an array of URLs?
- A blank page is useless, the code should at least wait until there is one result to display.
- The [context](https://golang.org/pkg/context/) type is often used with http handlers. Can you replace the `time.After` with an expiring context?

- 结果应作为 JSON 对象返回。也许我们可以使用结构而不是 URL 数组？
- 空白页是没有用的，代码至少应该等到有一个结果显示。
- [context](https://golang.org/pkg/context/) 类型通常与 http 处理程序一起使用。你能用一个过期的上下文替换`time.After`吗？

------

Hi, I’m Justin Fuller. Thanks for reading my post. Before you go, I need to let you know that everything I’ve written here is my own opinion and is not intended to represent my employer. All code samples are my own.

你好，我是贾斯汀富勒。感谢您阅读我的帖子。在你走之前，我需要让你知道我在这里写的一切都是我自己的观点，并不代表我的雇主。所有代码示例都是我自己的。

I’d also love to hear from you, please feel free to follow me on [Github](https://github.com/justindfuller) or [Twitter](https://twitter.com/justin_d_fuller). Thanks again for reading! 

我也很想收到您的来信，请随时在 [Github](https://github.com/justindfuller) 或 [Twitter](https://twitter.com/justin_d_fuller) 上关注我。再次感谢阅读！

