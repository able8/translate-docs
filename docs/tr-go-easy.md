# Go is not an easy language

# Go 不是一门简单的语言

Written on 22 Feb 2021

写于 2021 年 2 月 22 日

Discussions:  [Lobsters](https://lobste.rs/s/ee6nsc/go_is_not_easy_language),  [Hacker News](https://news.ycombinator.com/item?id=26220693),  [/r/golang]( https://www.reddit.com/r/golang/comments/lpeafy/go_is_not_an_easy_language/),  [/r/programming](https://www.reddit.com/r/golang/comments/lpo6zh/go_is_not_an_easy_language/)

讨论：[龙虾](https://lobste.rs/s/ee6nsc/go_is_not_easy_language)、[黑客新闻](https://news.ycombinator.com/item?id=26220693)、[/r/golang](https://www.reddit.com/r/golang/comments/lpeafy/go_is_not_an_easy_language/), [/r/programming](https://www.reddit.com/r/golang/comments/lpo6zh/go_is_not_an_easy_language/)

Go is not an easy programming language. It *is* simple in many ways: the syntax is simple, most of the semantics are simple. But a language is more than just syntax; it’s about doing useful *stuff*. And doing useful stuff is not always easy in Go.

Go 不是一种简单的编程语言。它在很多方面都很简单：语法很简单，大部分语义都很简单。但是一门语言不仅仅是语法；这是关于做有用的*东西*。在 Go 中做有用的事情并不总是那么容易。

Turns out that combining all those simple features in a way to do something useful can be tricky. How do you remove an item from an array in Ruby? `list.delete_at(i)`. And remove entries by value? `list.delete(value)`. Pretty easy, yeah?

事实证明，将所有这些简单的功能结合起来做一些有用的事情可能会很棘手。你如何从 Ruby 的数组中删除一个项目？ `list.delete_at(i)`。并按值删除条目？ `list.delete(value)`。很简单吧？

In Go it’s … less easy; to remove the index `i` you need to do:

在 Go 中，它……不那么容易；要删除索引`i`，您需要执行以下操作：

```
list = append(list[:i], list[i+1:]...)
```

And to remove the value `v` you’ll need to use a loop:

要删除值 `v`，您需要使用循环：

```
n := 0
for _, l := range list {
    if l != v {
        list[n] = l
        n++
    }
}
list = list[:n]
```

Is this unacceptably hard? Not really; I think most programmers can figure out what the above does even without prior Go experience. But it’s not exactly *easy* either. I'm usually lazy and copy these kind of things from the [Slice Tricks](https://github.com/golang/go/wiki/SliceTricks) page because I want to focus on actually solving the problem at hand, rather than plumbing like this.

这很难接受吗？并不真地;我认为大多数程序员即使没有 Go 经验也能弄清楚上面的内容。但这也不是完全*容易*。我通常很懒，从[切片技巧](https://github.com/golang/go/wiki/SliceTricks)页面复制这些东西，因为我想专注于实际解决手头的问题，而不是这样的管道。

It’s also easy to get it (subtly) wrong or suboptimal, especially for less experienced programmers. For example compare the above to copying to a new array and copying to a new pre-allocated array (`make([]string, 0, len(list))`):

它也很容易（微妙地）错误或次优，尤其是对于经验不足的程序员。例如，将上面的复制到一个新数组和复制到一个新的预分配数组（`make([]string, 0, len(list))`）进行比较：

```
InPlace             116 ns/op      0 B/op   0 allocs/op
NewArrayPreAlloc    525 ns/op    896 B/op   1 allocs/op
NewArray           1529 ns/op   2040 B/op   8 allocs/op
```

While 1529ns is still plenty fast enough for many use cases and isn't something to excessively worry about, there are plenty of cases where these things *do* matter and having the guarantee to always use the best possible algorithm with `list.delete( value)` has some value.

虽然 1529ns 对于许多用例来说仍然足够快，并且不需要过度担心，但在很多情况下，这些事情 * 做 * 很重要，并且保证始终使用最佳算法和 `list.delete( value)` 有一定的价值。

------

Goroutines are another good example. “Look how easy it is to start a goroutine! Just add `go` and you’re done!” Well, yes; you’re done until you have five million of those running at the same time and then you’re left wondering where all your memory went, and it’s not hard to “leak” goroutines by accident either.

Goroutines 是另一个很好的例子。 “看看启动一个 goroutine 是多么容易！只需添加`go`就完成了！”嗯，是;你已经完成了，直到你有 500 万个 goroutine 同时运行，然后你会想知道你所有的内存都去了哪里，而且不小心“泄漏”goroutines 也不难。

There are a number of patterns to limit the number of goroutines, and none of them are exactly easy. A simple example might be something like:

有许多模式可以限制 goroutine 的数量，但没有一个是非常容易的。一个简单的例子可能是这样的：

```
var (
    jobs    = 20                 // Run 20 jobs in total.
    running = make(chan bool, 3) // Limit concurrent jobs to 3.
    wg      sync.WaitGroup       // Keep track of which jobs are finished.
)

wg.Add(jobs)
for i := 1;i <= jobs;i++ {
    running <- true // Fill running;this will block and wait if it's already full.

    // Start a job.
    go func(i int) {
        defer func() {
            <-running // Drain running so new jobs can be added.
            wg.Done() // Signal that this job is done.
        }()

        // "do work"
        time.Sleep(1 * time.Second)
        fmt.Println(i)
    }(i)
}

wg.Wait() // Wait until all jobs are done.
fmt.Println("done")
```

There’s a reason I annotated this with some comments: for people not intimately familiar with Go this may take some effort to understand. This also won’t ensure that the numbers are printed in order (which may or may not be a requirement).

我用一些评论对此进行注释是有原因的：对于不熟悉 Go 的人来说，这可能需要一些努力才能理解。这也不能确保按顺序打印数字（这可能是也可能不是要求）。

Go’s concurrency primitives may be simple and easy to use, but combining them to solve common real-world scenarios is a lot less simple. The original version of the above example [was actually incorrect](https://lobste.rs/s/ee6nsc/go_is_not_easy_language#c_gdnw5e) 😅

Go 的并发原语可能简单易用，但将它们结合起来解决常见的现实世界场景要简单得多。上面例子的原始版本[实际上是不正确的](https://lobste.rs/s/ee6nsc/go_is_not_easy_language#c_gdnw5e) 😅

------

In [Simple Made Easy](https://www.infoq.com/presentations/Simple-Made-Easy/) Rich Hickey argues that we shouldn't confuse “simple” with “it's easy to write”: just because you can do something useful in one or two lines doesn't mean the underlying concepts – and therefore the entire program – are “simple” as in “simple to understand”.

在 [Simple Made Easy](https://www.infoq.com/presentations/Simple-Made-Easy/) 中，Rich Hickey 认为我们不应该将“简单”与“易于编写”混淆：仅仅因为您可以用一两行来做一些有用的事情并不意味着底层概念——因此整个程序——像“简单易懂”一样“简单”。

I feel there is some wisdom in this; in most cases we shouldn’t sacrifice “simple” for “easy”, but that doesn’t mean we can’t think at all about how to make things easier. Just because concepts are simple doesn’t mean they’re easy to use, can’t be misused, or can’t be used in ways that lead to (subtle) bugs. Pushing Hickey’s argument to the extreme we’d end up with something like [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) and that would of course be silly.

我觉得这里面有一些智慧；在大多数情况下，我们不应该为了“容易”而牺牲“简单”，但这并不意味着我们根本无法考虑如何使事情变得更容易。仅仅因为概念简单并不意味着它们易于使用、不能被滥用或不能以导致（微妙）错误的方式使用。将 Hickey 的论点推到极致，我们最终会得到类似 [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) 的东西，这当然是愚蠢的。

Ideally a language should reduce the cognitive load required to reason about its behaviour; there are many ways to increase this cognitive load: complex intertwined language features is one of them, and getting “distracted” by implementing fairly basic things from those simple concepts is another: it’s another block of code I need to reason about. While I’m not overly concerned about code formatting or syntax choices, I do think it can matter to reduce this cognitive load when reading code.

理想情况下，语言应该减少对其行为进行推理所需的认知负荷；有很多方法可以增加这种认知负担：复杂的相互交织的语言特征就是其中之一，通过从这些简单概念中实现相当基本的东西而“分心”是另一种：这是我需要推理的另一个代码块。虽然我并不太关心代码格式或语法选择，但我确实认为在阅读代码时减少这种认知负担很重要。

The lack of generics probably plays some part here; implementing a `slices` package which does these kind of things in a generic way is hard right now. Generics make this possible and also makes things more complex (more language features are used), but they also make things easier and, arguably, less complex on other fronts.[[1\]](https://www.arp242.net/go-easy.html#fn:g)

缺乏泛型可能在这里起到了一定的作用。现在实现一个以通用方式做这些事情的 `slices` 包是很困难的。泛型使这成为可能，也使事情变得更加复杂（使用了更多语言功能），但它们也使事情变得更容易，并且可以说，在其他方面不那么复杂。[[1\]](https://www.arp242.net/go-easy.html#fn:g)

------

Are these insurmountable problems? No. I still use (and like) Go after all. But I also don’t think that Go is a language that you “could pick up in ~5-10 minutes”, which was the comment that prompted this post; a sentiment I’ve seen expressed many times, although usually with less extreme timeframes (“1-2 days”, “1 week”).

这些是无法解决的问题吗？不。毕竟我仍然使用（并且喜欢）Go。但我也不认为 Go 是一种你“可以在大约 5-10 分钟内学会”的语言，这是促使这篇文章的评论；我见过多次表达的情绪，尽管通常没有那么极端的时间范围（“1-2 天”、“1 周”）。

As a corollary to all of the above; learning the language isn’t just about learning the syntax to write your `if`s and `for`s; it’s about learning a way of thinking. I’ve seen many people coming from Python or C♯ try to shoehorn concepts or patterns from those languages in Go. Common ones include using struct embedding as inheritance, panics as exceptions, “pseudo-dynamic programming” with interface{}, and so forth. It rarely ends well, if ever.

作为上述所有内容的推论；学习语言不仅仅是学习语法来编写你的 if 和 for ；这是关于学习一种思维方式。我见过很多来自 Python 或 C♯ 的人试图在 Go 中从这些语言中硬塞概念或模式。常见的包括使用结构嵌入作为继承、恐慌作为异常、使用 interface{} 的“伪动态编程”等等。它很少有好的结局，如果有的话。

I did this as well when I was writing my first Go program; it’s only natural. And when I started as a Ruby programmer I tried to write Python code in Ruby (although this works a bit better as the languages are more similar, but there are still plenty of odd things you can do such as using `for` loops).

我在编写第一个 Go 程序时也是这样做的；这很自然。当我开始成为 Ruby 程序员时，我尝试用 Ruby 编写 Python 代码（尽管这会更好一些，因为语言更相似，但是您仍然可以做很多奇怪的事情，例如使用 `for` 循环）。

This is why I don’t like it when people get redirected to the Tour of Go to “learn the language”, as it just teaches basic syntax and little more. It’s nice as a little, well, *tour* to get a bit of a feel of the language and see how it roughly works and what it can roughly do, but it’s ill-suited to actually learn the language.

这就是为什么当人们被重定向到 Go 之旅以“学习语言”时，我不喜欢它，因为它只教授基本的语法和其他内容。有点不错，嗯，*游览* 来了解一下这门语言，看看它大致是如何工作的以及它大致可以做什么，但它不适合实际学习这门语言。

**Footnotes**

**脚注**

1. Contrary to popular belief the [Go team was never “against” generics](https://research.swtch.com/generic); I’ve seen many comments to the effect of “the Go team doesn’t think  generics are useful”, but this was never the case. [↩](https://www.arp242.net/go-easy.html#fnref:g)

1. 与流行的看法相反，[围棋团队从不“反对”泛型](https://research.swtch.com/generic)；我看到过很多关于“Go 团队认为泛型没有用”的评论，但事实并非如此。 [↩](https://www.arp242.net/go-easy.html#fnref:g)

**Feedback**

**回馈**

Contact me at                 [martin@arp242.net](mailto:martin@arp242.net),                 [GitHub](https://github.com/arp242/arp242.net/issues/new), or                 [@arp242_martin](https://twitter.com/arp242_martin)                 for feedback, questions, etc.

通过 [martin@arp242.net](mailto:martin@arp242.net)、[GitHub](https://github.com/arp242/arp242.net/issues/new) 或 [@arp242_martin](https)与我联系://twitter.com/arp242_martin) 以获取反馈、问题等。

**Other Go posts**

**其他围棋帖子**

- 21 Nov 2019 [Go’s features of last resort](https://www.arp242.net/go-last-resort.html)
- 10 Dec 2020 [Bitmasks for nicer APIs](https://www.arp242.net/bitmask.html) 

- 2019 年 11 月 21 日 [Go 的最后手段](https://www.arp242.net/go-last-resort.html)
- 2020 年 12 月 10 日 [用于更好 API 的位掩码](https://www.arp242.net/bitmask.html)

