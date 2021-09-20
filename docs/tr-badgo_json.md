# Bad Go: Adventures with JSON marshalling

# Bad Go：JSON 编组的冒险

Adventures for the indoors

室内冒险

Sat, Oct 5, 2019



This is a story about encoding/json in the Go standard library. I’m not going to say this is bad Go. That would be heresy. But there is an  aspect of marshalling that could be improved. Because it is in the  standard library it isn’t bad Go, but if you followed the pattern in  your own code then that would be a mistake. Outside of the standard  library it would lose its magical aura, and it would be bad Go.

这是一个关于 Go 标准库中 encoding/json 的故事。我不会说这是不好的 Go。那将是异端邪说。但是编组的一个方面可以改进。因为它在标准库中，所以 Go 还不错，但是如果您在自己的代码中遵循该模式，那将是一个错误。在标准库之外，它将失去其神奇的光环，这将是糟糕的 Go 。

My frustration is with the Marshaler interface and the MarshalJSON  method. This method makes it pretty much impossible for custom JSON  marshalling to be efficient. The inimitable Mr. Cheney has recently  warned us about this very issue [here](https://dave.cheney.net/2019/09/05/dont-force-allocations-on-the-callers-of-your-api).

我对 Marshaler 接口和 MarshalJSON 方法感到沮丧。这种方法使得自定义 JSON 编组几乎不可能高效。独一无二的切尼先生最近就这个问题向我们发出警告 [这里](https://dave.cheney.net/2019/09/05/dont-force-allocations-on-the-callers-of-your-api)。

(To be clear, although I did sit next to Mr Cheney at a meetup, and  once he did like one of my tweets, that does not mean he in any way  endorses this blog or its content)

（需要明确的是，虽然我在一次聚会上确实坐在切尼先生旁边，而且一旦他喜欢我的一条推文，这并不意味着他以任何方式认可这个博客或其内容）

Let’s try to demonstrate the problem. We’ll start by marshalling a very simple struct in a simple benchmark.

让我们尝试演示这个问题。我们将从在一个简单的基准测试中编组一个非常简单的结构开始。

```go
type mystruct struct {
    A int    `json:"a,omitempty"`
    B string `json:"b,omitempty"`
}

func BenchmarkJSONMarshal(b *testing.B) {
    b.ReportAllocs()
    var data = mystruct{
        A: 42,
        B: "42",
    }
    for i := 0;i < b.N;i++ {
        _, err := json.Marshal(&data)
        if err != nil {
            b.Errorf("failed to marshal json. %s", err)
        }
    }
}
BenchmarkJSONMarshal-8   3627376    316 ns/op   32 B/op   1 allocs/op
```

If we run this we find there’s just 1 allocation per  marshalling attempt, which is the byte slice containing the marshalled  data. It would be nice if we could re-use a slice for this, but one  allocation is not too upsetting. And if we really want to we can use an [encoder](https://golang.org/pkg/encoding/json/#Encoder) to avoid this.

如果我们运行它，我们会发现每次编组尝试只有 1 个分配，这是包含编组数据的字节切片。如果我们可以为此重新使用一个切片会很好，但是一个分配并不太令人沮丧。如果我们真的想要，我们可以使用 [encoder](https://golang.org/pkg/encoding/json/#Encoder) 来避免这种情况。

So what am I complaining about? Well, let’s modify our struct a little to add a time.

那我在抱怨什么？好吧，让我们稍微修改一下结构以添加时间。

```go
type mystruct struct {
    A int       `json:"a,omitempty"`
    B string    `json:"b,omitempty"`
    C time.Time `json:"c"`
}

func BenchmarkJSONMarshal(b *testing.B) {
    b.ReportAllocs()
    var data = mystruct{
        A: 42,
        B: "42",
        C: time.Now(),
    }
    for i := 0;i < b.N;i++ {
        _, err := json.Marshal(&data)
        if err != nil {
            b.Errorf("failed to marshal json. %s", err)
        }
    }
}
BenchmarkJSONMarshal-8    981222   1345 ns/op  208 B/op   4 allocs/op
```

Suddenly we’re making 4 allocations per marshalling  attempt! 3 additional allocations because we’ve added a time. Why would  that be? Well, one issue is that the json package does not natively  understand time.Time, and marshals it via the Marshaler interface. time.Time implements [MarshalJSON](https://golang.org/pkg/time/#Time.MarshalJSON). This forces an additional allocation because the method is defined to return a `[]byte` with the marshalled time. There’s no mechanism in the API to allow this custom marshaler to append it’s data to the data marshalled so far. It  needs to allocate a separate slice that it returns (thus forcing a heap  allocation), and which the json library then appends to its output.

突然间，我们每次编组尝试进行 4 次分配！ 3 个额外的分配，因为我们增加了一个时间。为什么会这样？嗯，一个问题是 json 包本身并不理解 time.Time，而是通过 Marshaler 接口对其进行编组。 time.Time 实现 [MarshalJSON](https://golang.org/pkg/time/#Time.MarshalJSON)。这会强制进行额外的分配，因为该方法被定义为返回带有编组时间的 `[]byte`。 API 中没有任何机制允许这个自定义封送拆收器将它的数据附加到到目前为止封送的数据中。它需要分配一个单独的切片，它返回（从而强制分配堆)，然后 json 库将其附加到其输出。

That explains 1 additional allocation. Why are there 3? Well, we can benchmark Time.MarshalJSON to see what it is doing.

这解释了 1 个额外的分配。为什么有3个？好吧，我们可以对 Time.MarshalJSON 进行基准测试，看看它在做什么。

```go
func BenchmarkTimeMarshal(b *testing.B) {
    b.ReportAllocs()
    var t time.Time

    for i := 0;i < b.N;i++ {
        _, err := t.MarshalJSON()
        if err != nil {
            b.Errorf("failed to marshal. %s", err)
        }
    }
}
BenchmarkTimeMarshal-8   3400222    378 ns/op   48 B/op   1 allocs/op
```

This only creates 1 allocation. So the other 2 must  somehow come about within the json package itself, presumably as  additional overhead joining up the results.

这只会创建 1 个分配。因此，其他 2 个必须以某种方式出现在 json 包本身中，大概是连接结果的额外开销。

If we run the benchmark under the profiler we discover the causes of the 4 allocations.

如果我们在分析器下运行基准测试，我们会发现 4 次分配的原因。

1. The byte slice that holds the final marshalled JSON.
2. The byte slice Time.MarshalJSON is forced to generate.
3. Some additional overhead copying the marshalled JSON from  Time.MarshalJSON into the result byte slice. This uses json.Compact,  which allocates a scanner while it does the copying because it also  checks the JSON is valid and ensures insignificant space is removed from the JSON.
4. To access the Marshaler interface, json uses the reflect package, and in fact creates a new `interface{}` value pointing to the time value. This somehow causes an allocation. 

1. 保存最终编组 JSON 的字节切片。
2. 强制生成字节切片Time.MarshalJSON。
3. 将编组的 JSON 从 Time.MarshalJSON 复制到结果字节片的一些额外开销。这使用 json.Compact，它在进行复制时分配一个扫描器，因为它还会检查 JSON 是否有效并确保从 JSON 中删除无关紧要的空间。
4. 为了访问Marshaler接口，json使用了reflect包，实际上创建了一个指向时间值的新`interface{}`值。这以某种方式导致分配。

As far as I can tell all 3 of these allocations are currently unavoidable if you use a custom JSON marshaler for a type.

据我所知，如果您对类型使用自定义 JSON 封送拆收器，则所有 3 种分配目前都是不可避免的。

Why do I find this so frustrating? To me the existence of the  json.Marshaler interface looks like an escape hatch: a mechanism to do  things that are out of the ordinary; to put effort in and improve  performance. But it isn’t that. It’s a garbage chute - use it and you’ll end up stuck in a bin covered in garbage.

为什么我觉得这很令人沮丧？对我来说，json.Marshaler 接口的存在就像一个逃生舱：一种做不同寻常事情的机制；努力并提高绩效。但事实并非如此。这是一个垃圾槽 - 使用它，你最终会被困在一个被垃圾覆盖的垃圾箱里。

- Have lots of timestamps in your data => covered in garbage
- Want to use json.RawMessage to avoid encoding parts of your data => covered in garbage
- Need to express null fields, but want to avoid using pointers so you don’t get covered in garbage? Well, you’ll do a lot of work and end up  covered in garbage.

- 在你的数据中有很多时间戳 => 被垃圾覆盖
- 想要使用 json.RawMessage 来避免编码部分数据 => 被垃圾覆盖
- 需要表达空字段，但想避免使用指针以免被垃圾覆盖？好吧，你会做很多工作，最终会被垃圾覆盖。

Now, none of this is a problem if you’re not marshalling a lot of  JSON. But if you are it starts to make Go look like a poor choice. Or  you have to look at third-party JSON encoders, which isn’t an  unreasonable option but is somehow unsatisfying.

现在，如果您不编组大量 JSON，这些都不是问题。但如果你是，它开始让 Go 看起来像是一个糟糕的选择。或者您必须查看第三方 JSON 编码器，这不是一个不合理的选择，但在某种程度上令人不满意。

How could we improve on this? What if we added a second marshaler interface?

我们如何改进？如果我们添加第二个封送拆收器接口会怎样？

```go
type MarshalAppender interface {
    MarshalAppendJSON(buf []byte) ([]byte, error)
}
```

Implementers of this interface append their json directly to the `buf` parameter passed in. We define things so that MarshalAppendJSON must  append valid JSON without any redundant white space. Finally we work out why accessing the interface method causes an allocation and fix it. Then we’ll have the possibility of allocation-free custom JSON  marshalling.

这个接口的实现者直接将他们的 json 附加到传入的 `buf` 参数。我们定义了一些东西，以便 MarshalAppendJSON 必须附加有效的 JSON，没有任何多余的空格。最后我们找出为什么访问接口方法会导致分配并修复它。然后我们将有可能进行无分配的自定义 JSON 编组。

## Is it Bad Go?

## 是不是很糟糕？

MarshalAppender is perhaps a little more complicated than Marshaler. And simple is often best. But if your code is a fundamental building  block, either within your own project or for projects throughout the  world, I’d argue it’s worth going the extra mile to provide both  efficient implementations and APIs that can be used efficiently.

MarshalAppender 可能比 Marshaler 稍微复杂一些。简单的往往是最好的。但是，如果您的代码是一个基本构建块，无论是在您自己的项目中还是在世界各地的项目中，我认为值得付出额外的努力来提供高效的实现和可以高效使用的 API。

Providing just the simple interface may seem simpler and clearer. But what happens when someone needs that greater efficiency? Either they’re stuck, or they create a whole new implementation, or they go to extreme lengths to deal with the garbage collector. You’ve not reduced the complexity in the world - you’ve deferred it. And increased it.

只提供简单的界面可能看起来更简单、更清晰。但是当有人需要更高的效率时会发生什么？要么他们被卡住了，要么他们创建了一个全新的实现，或者他们竭尽全力处理垃圾收集器。你并没有降低世界的复杂性——你已经推迟了它。并增加了它。

## Next steps

##  下一步

I’m actually going to [propose this](https://github.com/golang/go/issues/34701) to the Go team and try to contribute the change. I intend to write  about the experience in a future blog. Hopefully it won’t be terribly  interesting! 

实际上，我打算 [propose this](https://github.com/golang/go/issues/34701) 给 Go 团队并尝试为更改做出贡献。我打算在以后的博客中写下这段经历。希望它不会非常有趣！

