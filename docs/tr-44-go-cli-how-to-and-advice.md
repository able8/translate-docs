# [Writing Go CLIs With Just Enough Architecture](https://blog.carlmjohnson.net/post/2020/go-cli-how-to-and-advice/)

# [使用足够的架构编写 Go CLI](https://blog.carlmjohnson.net/post/2020/go-cli-how-to-and-advice/)

Thursday, June 4, 2020 https://github.com/carlmjohnson/go-grab-xkcd

As someone who has written [quite a few](https://blog.carlmjohnson.net/post/2018/go-cli-tools/) command line applications in Go, I was interested the other day when I saw [Diving into Go by building a CLI application](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) by [Eryb](https://eryb.space) on social media. In the post, the author describes writing a simple application to fetch comics from the [XKCD API](https://xkcd.com/json.html) and display the results. It looks like this:

作为用 Go 编写 [相当多](https://blog.carlmjohnson.net/post/2018/go-cli-tools/)命令行应用程序的人，前几天看到 [Diving into通过构建 CLI 应用程序](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) by [Eryb](https://eryb.space）在社交媒体上。在帖子中，作者描述了编写一个简单的应用程序来从 [XKCD API](https://xkcd.com/json.html) 中获取漫画并显示结果。它看起来像这样：

```
$ go-grab-xkcd --help

Usage of go-grab-xkcd:
  -n int
        Comic number to fetch (default latest)
  -o string
        Print output in format: text/json (default "text")
  -s    Save image to current directory
  -t int
        Client timeout in seconds (default 30)
```


I came away a little disappointed though because I felt  like the final result was both a little undercooked and a little  overcooked: undercooked in that it didn't handle errors robustly enough  and overcooked in that it created some abstractions speculatively in a  way that I felt were [unlikely to pay off in the long run](https://en.wikipedia.org/wiki/You_aren't_gonna_need_it).

但我离开时有点失望，因为我觉得最终结果既有点未煮熟又有点过熟：未煮熟的原因是它没有足够稳健地处理错误，而过熟的原因是它以一种我认为的方式推测性地创建了一些抽象[从长远来看不太可能获得回报](https://en.wikipedia.org/wiki/You_aren't_gonna_need_it)。

Naturally, one thing lead to another, and I [forked and rewrote](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910) the demo app myself to demonstrate what I consider to be **just enough** architecture for a Go command line app.

自然，一件事导致另一件事，我[分叉并重写](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910)演示应用程序自己演示我认为是**只是足够** Go 命令行应用程序的架构。

------

Before I go too much further with this post, let me say that I don’t mean to  be overly negative. For such a small app, none of what follows  particularly matters. This app is simple enough that you could write it  in Bash if you wanted to. The point was to write an command line  interface for fun and get some experience doing it. So, with that in  mind, go read [Eryb's post](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) first and then come back and let me share how I would approach writing  the same CLI in Go with the experience I've gained by writing so many CLIs and having to extend them on and off over a period of years.

在我进一步讨论这篇文章之前，让我说我并不是想过于消极。对于这么小的应用程序，接下来的一切都不是特别重要。这个应用程序非常简单，如果你愿意，你可以用 Bash 编写它。重点是编写一个有趣的命令行界面并获得一些经验。因此，考虑到这一点，请先阅读 [Eryb 的帖子](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) 然后再来回过头来，让我分享一下我将如何利用通过编写如此多的 CLI 并不得不在几年内断断续续地扩展它们而获得的经验，在 Go 中编写相同的 CLI。

My core principle is that while there are many possible architectures for an app—[MVC](https://en.wikipedia.org/wiki/Model–view–controller),[hexagonal](https://en.wikipedia. org/wiki/Hexagonal_architecture_(software)), etc—there are three layers that you almost always need and never regret writing. You want one layer to **handle user input** and get it into a normalized form. You want one layer to **do your actual task**. And you want one layer to **handle formatting and output** to the user. Those are the three layers you always need. Other layers  you may or may not need and can add it when it becomes clear that they  might help. For a Go command line interface, this means the most  important thing is to separate flag parsing stuff from execution, and  everything else is not a big deal to let evolve over time. The main  challenge is avoiding create abstractions that don’t actually pay for  themselves in setup time versus time saved in writing future extensions  to your program.

我的核心原则是，虽然应用程序有许多可能的架构——[MVC](https://en.wikipedia.org/wiki/Model–view–controller)、[hexagonal](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) 等——你几乎总是需要三层，永远不会后悔写作。您希望有一层 ** 处理用户输入** 并将其转换为规范化形式。您需要一层来**执行您的实际任务**。并且您需要一层来**处理格式化和输出**给用户。这些是您始终需要的三层。您可能需要也可能不需要的其他层，并且可以在很明显它们可能有帮助时添加它。对于 Go 命令行界面，这意味着最重要的事情是将标志解析的东西与执行分开，其他的一切都不是什么大不了的事情，让随着时间的推移而发展。主要的挑战是避免创建抽象，这些抽象实际上不会在设置时间与为您的程序编写未来扩展所节省的时间上付出代价。

The first thing I noticed looking at [the source to the demo CLI](https://github.com/carlmjohnson/go-grab-xkcd/tree/v0.0.0) is that it was divided into three packages: main (necessarily ), model, and client. This is one package **too few** and two packages **too many**! A Go CLI should start with two packages: main (in which func main has  only one line) and everything else (in this case, call it package  grabxkcd).

在查看 [演示 CLI 的源代码](https://github.com/carlmjohnson/go-grab-xkcd/tree/v0.0.0)时，我注意到的第一件事是它分为三个包： main （必须)、模型和客户端。这是一包**太少**和两包**太多**！ Go CLI 应该从两个包开始：main（其中 func main 只有一行）和其他所有包（在这种情况下，称为包grabxkcd）。

Before I explain about the single line `main` function, let’s look at a common (anti-)pattern for Go CLIs (which the  demo app fortunately doesn’t apply), in which you check for errors with a helper function like this:

在我解释单行 `main` 函数之前，让我们看一下 Go CLI 的一个常见（反）模式（幸运的是演示应用程序不适用），在其中你使用辅助函数检查错误，如下所示：

```go
func check(err error) { // or try() or must() or die()
    if err != nil {
        log.Fatal(err) // or panic or os.Exit
    }
}

func main() {
    v, err := doSomething()
    check(err)
    // …
}
```


I myself have written this more than a few times,  but for any serious CLI, this pattern breaks down quickly because it  allows for only one response to errors, immediately aborting. Often that is a good beginning for a CLI, but as it grows you find that you need  to handle some errors by retrying or ignoring or user notification, but  because your functions all return a single value instead of value plus  error, it becomes a ton of work to add proper error handling in after  the fact.

我自己已经多次写过这个，但是对于任何严重的 CLI，这种模式很快就会崩溃，因为它只允许对错误做出一个响应，立即中止。通常这对 CLI 来说是一个好的开始，但随着它的增长，你发现你需要通过重试或忽略或用户通知来处理一些错误，但因为你的函数都返回单个值而不是值加错误，它变得一吨事后添加适当的错误处理的工作。

Okay, so if we shouldn’t use a `check` function, why should we use a single line `main` function? The original demo app continues even after it finds an error  in fetching the comic. This was because it couldn’t return err in main  and it also didn’t use the `check` func pattern. It also  didn’t exit with a non-zero code on error, which matters if you want to  use your CLI in a scripting pipeline. The single line main func solves  these problems.

好的，如果我们不应该使用 `check` 函数，为什么要使用单行 `main` 函数？即使在获取漫画时发现错误，原始演示应用程序仍会继续。这是因为它不能在 main 中返回 err 并且它也没有使用 `check` func 模式。它也不会在出现错误时以非零代码退出，如果您想在脚本管道中使用 CLI，这一点很重要。单行 main func 解决了这些问题。

I took the idea of a single line main function from a [blog post by Nate Finch](https://npf.io/2016/10/reusable-commands/). He wrote:

我从 [Nate Finch 的博文](https://npf.io/2016/10/reusable-commands/) 中获得了单行主函数的想法。他写了：

I think there’s only approximately one right answer to “what’s in your package main?” and that’s this:

我认为“你的主要包里有什么？”只有大约一个正确的答案。就是这样：

```go
// command main documentation here.
package main

import (
    "os"

    "github.com/you/proj/cli"
)
func main{
    os.Exit(cli.Run())
}
```


You can read [his post](https://npf.io/2016/10/reusable-commands/) for more explanation of why, but for me the argument comes down to that it makes your program more testable and more extensible without being  much more work. Like a `check` function, the single line main lets you easily handle errors in execution by creating a clear pathway  between errors and program termination, but unlike a `check` function, the single line main does so without closing off the path for future enhancements.

您可以阅读 [他的帖子](https://npf.io/2016/10/reusable-commands/) 以了解更多原因，但对我而言，论点归结为它使您的程序更具可测试性和可扩展性，而无需更多的工作。与 `check` 函数一样，单行 main 通过在错误和程序终止之间创建一条清晰的路径，让您可以轻松处理执行中的错误，但与 `check` 函数不同的是，单行 main 这样做不会关闭将来的路径增强功能。

The demo CLI uses the standard Go flag package. The flag package has many  detractors because it is (like Go itself) quite idiosyncratic, but I  find that it is mostly great. The one thing I don't like about the flag  package is that [the documentation](https://pkg.go.dev/flag?tab=doc#pkg-overview) suggests using it by setting global variables and using `init ` functions, like

演示 CLI 使用标准的 Go 标志包。 flag 包有很多批评者，因为它（就像 Go 本身一样）非常特殊，但我发现它主要是很棒的。我不喜欢 flag 包的一件事是 [the documentation](https://pkg.go.dev/flag?tab=doc#pkg-overview) 建议通过设置全局变量和使用 `init ` 函数，例如

```go
var flagvar int
func init() {
    flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```


In his post [A Theory of Modern Go](https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html), Peter Bourgon proposes the following guidelines:

在他的帖子 [现代 Go 理论](https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html) 中，Peter Bourgon 提出了以下指导方针：

> - No package level variables
> - No func init

> - 没有包级变量
> - 没有 func init

He does so because

他这样做是因为

> Package-global objects can encode state and/or behavior that is hidden from external  callers. Code that calls on those globals can have surprising side  effects, which subverts the reader’s ability to understand and mentally  model the program.

> 包全局对象可以编码对外部调用者隐藏的状态和/或行为。调用这些全局变量的代码可能会产生令人惊讶的副作用，这会破坏读者理解程序和对程序进行心理建模的能力。

I agree with this whole heartedly, so I think that when you use the flag package, it is important not to  use it with global variables in your package. Using global variables for flags inside of package main leads to spooky action at a distance very  quickly. Where did this variable come from? Why is it set this way? When you use global flags, these confusions crop up all the time. One thing I like about the original demo is that it didn’t use globals, but I think it could be refined even a little bit more.

我完全同意这一点，所以我认为当你使用标志包时，重要的是不要将它与你的包中的全局变量一起使用。将全局变量用于包 main 内的标志会导致非常快的远距离动作。这个变量从何而来？为什么要这样设置？当您使用全局标志时，这些混淆总是会出现。我喜欢原始演示的一件事是它没有使用全局变量，但我认为它可以进一步完善。

One more bit of background before I get into the nitty-gritty of my rewrite. In [How I write HTTP services after eight years](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html), Mat Ryer explains ,

在我进入重写的细节之前，还有一点背景知识。在[八年后我如何编写 HTTP 服务](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html)中，Mat Ryer 解释了,

> All of my components have a single server structure that usually ends up looking something like this:
>
> ```go
> type server struct {
>     db     *someDatabase
>     router *someRouter
>     email  EmailSender
> }
> ```


He then creates his routes and  handlers as methods hanging off the server type. I use the same  principle when writing a CLI. As the flag package processes my  arguments, I use them to create an `appEnv` struct that holds the options and controls the execution environment. 

然后他创建他的路由和处理程序作为挂在服务器类型上的方法。我在编写 CLI 时使用相同的原则。当 flag 包处理我的参数时，我使用它们来创建一个 `appEnv` 结构，用于保存选项并控制执行环境。

Putting these pieces together, my recently created CLIs all follow this  pattern: a single line main function calls a function (usually named `app.CLI` and in a separate package) that first uses the flag package to parse  user input and then kicks off actual execution as a method on an `appEnv` struct. Here is the code for [my version](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910) of the go-grab-xkcd app:

将这些部分放在一起，我最近创建的 CLI 都遵循这种模式：一行 main 函数调用一个函数（通常命名为 `app.CLI` 并在一个单独的包中），该函数首先使用标志包来解析用户输入，然后启动作为 `appEnv` 结构上的方法实际执行。这是 go-grab-xkcd 应用程序的 [我的版本](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910) 的代码：

```go
// In main.go 
package main
func main() {
    os.Exit(grabxkcd.CLI(os.Args[1:]))
}

// In grabxkcd/grabxkcd.go
func CLI(args []string) int {
    var app appEnv
    err := app.fromArgs(args)
    if err != nil {
        return 2
    }
    if err = app.run();err != nil {
        fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
        return 1
    }
    return 0
}
```


Next, as we wrap up handling user input, let’s take a look at the `appEnv` type and `fromArgs` method:

接下来，当我们结束处理用户输入时，让我们看看 `appEnv` 类型和 `fromArgs` 方法：

```go
// In grabxkcd/grabxkcd.go

type appEnv struct {
    hc         http.Client
    comicNo    int
    saveImage  bool
    outputJSON bool
}

func (app *appEnv) fromArgs(args []string) error {
    // Shallow copy of default client
    app.hc = *http.DefaultClient
    fl := flag.NewFlagSet("xkcd-grab", flag.ContinueOnError)
    fl.IntVar(
        &app.comicNo, "n", LatestComic, "Comic number to fetch (default latest)",
    )
    fl.DurationVar(&app.hc.Timeout, "t", 30*time.Second, "Client timeout")
    fl.BoolVar(
        &app.saveImage, "s", false, "Save image to current directory",
    )
    outputType := fl.String(
        "o", "text", "Print output in format: text/json",
    )
    if err := fl.Parse(args);err != nil {
        return err
    }
    if *outputType != "text" && *outputType != "json" {
        fmt.Fprintf(os.Stderr, "got bad output type: %q\n", *outputType)
        fl.Usage()
        return flag.ErrHelp
    }
    app.outputJSON = *outputType == "json"
    return nil
}
```


This is pretty similar to the original demo CLI, but I’ve made two tweaks beyond just assigning to `appEnv` members. Here is the original code for comparison:

这与最初的演示 CLI 非常相似，但除了分配给 appEnv 成员之外，我还做了两个调整。这是用于比较的原始代码：

```go
     comicNo := flag.Int(
        "n", int(client.LatestComic), "Comic number to fetch (default latest)",
    )
    clientTimeout := flag.Int64(
        "t", int64(client.DefaultClientTimeout.Seconds()), "Client timeout in seconds",
    )
    saveImage := flag.Bool(
        "s", false, "Save image to current directory",
    )
    outputType := flag.String(
        "o", "text", "Print output in format: text/json",
    )
```


The original code takes a timeout of `int64` seconds (`int32` is 136 years, which ought to be enough IMO), but Go has a built in `time.Duration` type. In other languages like JavaScript and Python it’s common to use  integers for durations and just have a convention about whether the  duration is in seconds or milliseconds or whatever, but in Go, you  should always use `time.Duration`. A nice trick is to set the timeout directly on your `http.Client` by using `flag.DurationVar` and passing in a pointer to `myclient.Timeout`.

原始代码的超时时间为 `int64` 秒（`int32` 是 136 年，IMO 应该足够了），但是 Go 有一个内置的 `time.Duration` 类型。在其他语言（如 JavaScript 和 Python）中，通常使用整数表示持续时间，并且只有关于持续时间是秒还是毫秒或其他什么的约定，但在 Go 中，您应该始终使用 `time.Duration`。一个不错的技巧是通过使用 `flag.DurationVar` 并传入一个指向 `myclient.Timeout` 的指针，直接在你的 `http.Client` 上设置超时。

Second, outputType is used to make a binary choice between text and json  output. In my code, I check for that and return an initialization error  if someone sets `-o unsupported`.

其次，outputType 用于在 text 和 json 输出之间进行二元选择。在我的代码中，如果有人设置了`-o unsupported`，我会检查并返回初始化错误。

------

That wraps up the user input layer! Whew. Now let’s look at **actual task execution**.

这样就完成了用户输入层！哇。现在让我们看看**实际的任务执行**。

The original demo had separate client and model packages, but if you look at the client, it imports returns a `model.Comic` object, so the client package cannot be used without the model package:

最初的演示有单独的客户端和模型包，但是如果您查看客户端，它会返回一个 `model.Comic` 对象，因此如果没有模型包，则无法使用客户端包：

```go
func NewXKCDClient() *XKCDClient
func (hc *XKCDClient) Fetch(n ComicNumber, save bool) (model.Comic, error)
func (hc *XKCDClient) SaveToDisk(url, savePath string) error
func (hc *XKCDClient) SetTimeout(d time.Duration)
```


If you look at the private members of the `client.XKCDClient` (the name repeats “client” which is often [a bad code smell](https://blog.golang.org/package-names)in Go!), you can see that it's actually just a `*http.Client` plus a URL builder for the XKCD API:

如果您查看 `client.XKCDClient` 的私有成员（名称重复“client”，这在 Go 中通常是[糟糕的代码味道](https://blog.golang.org/package-names)！），你可以看到它实际上只是一个 `*http.Client` 加上一个用于 XKCD API 的 URL 构建器：

```go
type XKCDClient struct {
    client  *http.Client
    baseURL string
}

func (hc *XKCDClient) buildURL(n ComicNumber) string {
    var finalURL string
    if n == LatestComic {
        finalURL = fmt.Sprintf("%s/info.0.json", hc.baseURL)
    } else {
        finalURL = fmt.Sprintf("%s/%d/info.0.json", hc.baseURL, n)
    }
    return finalURL
}
```


Coming from other languages, there is a temptation to build clients like this that “encapsulate” the details of the API  by, for example, hiding the base URL. The idea is that you may want to  point the base URL at localhost or a staging server for testing. I think that in Go this is basically a mistake. The `http.Client` type has a member called `Transport` of type `http.RoundTripper` interface. By changing the client transport, you can make an `http.Client` do whatever you want in testing—or even production. For example, you might create a `RoundTripper` that adds auth headers (Google does this for [their APIs](https://pkg.go.dev/cloud.google.com/go?tab=doc)) or does caching or reads test responses from a test file. Controlling the `http.Client` is extremely powerful, so any Go package designed to help users with an API should let them supply their own clients. Once you do that, there’s no longer any point in monkeying with base URLs.

来自其他语言，有一种诱惑力来构建这样的客户端，通过例如隐藏基本 URL 来“封装”API 的细节。这个想法是您可能希望将基本 URL 指向 localhost 或用于测试的登台服务器。我认为在 Go 中这基本上是一个错误。 `http.Client` 类型有一个名为 `Transport` 的成员，属于 `http.RoundTripper` 接口类型。通过改变客户端传输，你可以让 `http.Client` 做任何你想做的测试——甚至是生产。例如，您可以创建一个添加身份验证标头的“RoundTripper”（Google 为 [他们的 API](https://pkg.go.dev/cloud.google.com/go?tab=doc) 执行此操作）或执行缓存或从测试文件中读取测试响应。控制 `http.Client` 非常强大，因此任何旨在帮助用户使用 API 的 Go 包都应该让他们提供自己的客户端。一旦你这样做了，就再也没有必要使用基本 URL 了。

Since the client and models packages cannot be used apart from one another, they should  be combined. Because the client is just a wrapper around `http.Client` and a private URL builder, the URL builder should be made public and anyone can use their own `http.Client`. Once we’ve broken things down to that point, there’s not really enough  for a separate package, hence my joke that the original app had one  package too few and two too many. As Dave Cheney argued, [consider fewer, larger packages](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_consider_fewer_larger_packages). If we do want to separate the API stuff from the command line stuff  just for organization purposes, we can put it into a separate file in  the same package. My version looks like this:

由于客户端和模型包不能彼此分开使用，因此应将它们组合使用。因为客户端只是对 `http.Client` 和私有 URL 构建器的封装，所以 URL 构建器应该公开，任何人都可以使用他们自己的 `http.Client`。一旦我们把事情分解到那个点，就没有足够的单独的包，因此我开玩笑说原始应用程序的一个包太少，两个太多。正如 Dave Cheney 所说，[考虑更少、更大的包](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_think_fewer_larger_packages)。如果我们确实想将 API 内容与命令行内容分开只是为了组织目的，我们可以将它放在同一个包中的一个单独文件中。我的版本是这样的：

```go
// In grabxkcd/api.go

func BuildURL(comicNumber int) string {
    if comicNumber == LatestComic {
        return "https://xkcd.com/info.0.json"
    }
    return fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicNumber)
}

// APIResponse returned by the XKCD API
type APIResponse struct {
    Month       string `json:"month"`
    Number      int    `json:"num"`
    Link        string `json:"link"`
    Year        string `json:"year"`
    News        string `json:"news"`
    SafeTitle   string `json:"safe_title"`
    Transcript  string `json:"transcript"`
    Description string `json:"alt"`
    Image       string `json:"img"`
    Title       string `json:"title"`
    Day         string `json:"day"`
}
```


Back in `grabxkcd/grabxkcd.go`, I add `appEnv` methods to replace what we lost from the client:

回到 `grabxkcd/grabxkcd.go`，我添加了 `appEnv` 方法来替换我们从客户端丢失的内容：

```go
func (app *appEnv) fetchJSON(url string, data interface{}) error
func (app *appEnv) fetchAndSave(url, destPath string) error
```


The original demo app’s `Fetch` and `SaveToDisk` methods were tied into the XKCD API and couldn’t really be reused in other situations. By writing generic `fetchJSON` and `fetchAndSave` helpers that work with any JSON response or GET-able URL, it's more  likely that we'll be able to reuse our code, if not in this project then by copy-pasting it into a future project. When working with JSON in  particular, it’s often better to write functions/methods that take a  dynamic type (`interface{}`) instead of a static type (e.g. `APIResponse`). Accepting a dynamic type allows you to reuse your code in more  situations, and using a static type doesn’t actually add any practical  type safety, since the underlying calls to `json.Marshal` or whatever will still be dynamic themselves. 

原始演示应用程序的 `Fetch` 和 `SaveToDisk` 方法被绑定到 XKCD API 中，无法在其他情况下真正重用。通过编写通用的 `fetchJSON` 和 `fetchAndSave` 助手来处理任何 JSON 响应或可获取的 URL，我们更有可能重用我们的代码，如果不在这个项目中，那么通过将其复制粘贴到未来的项目。特别是在使用 JSON 时，通常最好编写采用动态类型（`interface{}`）而不是静态类型（例如`APIResponse`）的函数/方法。接受动态类型允许您在更多情况下重用您的代码，并且使用静态类型实际上并没有增加任何实际的类型安全，因为对 `json.Marshal` 或其他任何东西的底层调用本身仍然是动态的。

There are two execution details that the demo app hid away inside of other  functions that I want to bring to top level in mine. The [XKCD API](https://xkcd.com/json.html) is okay, but it doesn’t quite return the information we want. Specifically, instead of having an ISO-encoded date field, it has three  fields for day, month, and year, and instead of telling us the name of  its images, it gives us the whole URL of an image and the filename is  the last component of the path. Let’s fix those issues by adding methods onto the `APIResponse` type:

演示应用程序隐藏在其他函数中的两个执行细节，我想将它们带到我的顶层。 [XKCD API](https://xkcd.com/json.html) 没问题，但它并没有完全返回我们想要的信息。具体来说，它没有一个 ISO 编码的日期字段，它有三个字段，分别是日、月和年，而不是告诉我们它的图像的名称，它给了我们一个图像的完整 URL，文件名是最后一个路径的组成部分。让我们通过向 APIResponse 类型添加方法来解决这些问题：

```go
// Date returns a *time.Time based on the API strings (or nil if the response is malformed).
func (ar APIResponse) Date() *time.Time {
    t, err := time.Parse(
        "2006-1-2",
        fmt.Sprintf("%s-%s-%s", ar.Year, ar.Month, ar.Day),
    )
    if err != nil {
        return nil
    }
    return &t
}

func (ar APIResponse) Filename() string {
    return path.Base(ar.Image)
}
```


Notice that by explicitly parsing a date out of  the XKCD API, we were faced with a choice about error handling. Should  we return an error here? Should we return a blank object? Returning a  nil is a compromise choice. It means that we think it is unlikely that  the XKCD API will return a bad date, but if it does, we’re willing to  accept a runtime panic if we forget to check that the date is not nil  first. It’s not the most robust error handling strategy, but it’s a  middle ground for a validation error that we expect to never happen.

请注意，通过显式解析 XKCD API 中的日期，我们面临着关于错误处理的选择。我们应该在这里返回一个错误吗？我们应该返回一个空白对象吗？返回一个 nil 是一个折衷的选择。这意味着我们认为 XKCD API 不太可能返回错误的日期，但如果确实如此，如果我们忘记先检查日期是否不为零，我们愿意接受运行时恐慌。这不是最强大的错误处理策略，但它是我们希望永远不会发生的验证错误的中间地带。

------

We’ve looked at looked handling user input and actual execution. Now it’s time for **output**.

我们已经研究了如何处理用户输入和实际执行。现在是**输出**的时候了。

The original demo app has a `Comic` model type for output, which is a good practice, but I find the name `models.Comic` less than clear about its intent (doesn’t the API return a comic?). Let’s call it `grabxkcd.Output` instead.

最初的演示应用程序有一个用于输出的`Comic` 模型类型，这是一个很好的做法，但我发现名称`models.Comic` 不太清楚它的意图（API 不是返回一个漫画吗？）。我们将其称为“grabxkcd.Output”。

The original demo app also uses the same type for both JSON output and  plain text output, but I think that’s a little bit premature. Right now  we want these types to use the same underlying values, but there’s no  reason in principle they might not diverge in the future. Let’s put off  the decision for now and just have separate functions for JSON output  and pretty printing. If we want to unify them in the future (for  example, by replacing the pretty printer with a `template.Template` output that consumes `grabxkcd.Output`), it shouldn’t be too hard.

最初的演示应用程序也对 JSON 输出和纯文本输出使用相同的类型，但我认为这有点为时过早。现在我们希望这些类型使用相同的基础值，但原则上没有理由它们将来不会发生分歧。让我们暂时推迟决定，只为 JSON 输出和漂亮的打印提供单独的功能。如果我们将来想统一它们（例如，通过使用消耗 `grabxkcd.Output` 的 `template.Template` 输出替换漂亮的打印机），应该不会太难。

Here’s what our `appEnv.run` method looks like with these changes:

以下是我们的 `appEnv.run` 方法经过这些更改后的样子：

```go
func (app *appEnv) run() error {
    u := BuildURL(app.comicNo)

    var resp APIResponse
    if err := app.fetchJSON(u, &resp);err != nil {
        return err
    }
    if resp.Date() == nil {
        return fmt.Errorf("could not parse date of comic: %q/%q/%q",
            resp.Year, resp.Month, resp.Day)
    }
    if app.saveImage {
        if err := app.fetchAndSave(resp.Image, resp.Filename());err != nil {
            return err
        }
        fmt.Fprintf(os.Stderr, "Saved: %q\n", resp.Filename())
    }
    if app.outputJSON {
        return printJSON(resp)
    }

    return prettyPrint(resp)
}
```


------

With these changes in place, our [public API for package grabxkcd](https://godoc.org/github.com/carlmjohnson/go-grab-xkcd/grabxkcd) looks like this:

有了这些更改，我们的 [包grabxkcd 的公共 API](https://godoc.org/github.com/carlmjohnson/go-grab-xkcd/grabxkcd) 看起来像这样：

```go
func CLI(args []string) int
func BuildURL(comicNumber int) string
type APIResponse
func (ar APIResponse) Date() *time.Time
func (ar APIResponse) Filename() string
type Output
```




What I like about this is I feel it balances our  convenience as a developer when writing the code and the convenience of  any potential users of our package. This is a simple CLI, so probably no one will ever use this code but us. On the one hand, we found it more  convenient for our own development purposes to break out a second  package instead of just putting everything into package main. But on the other hand, we haven’t invested much in creating a big public API full  of features and interfaces that we’re just speculating might be useful  for someone, someday. This is the code we needed, and if it ever helps  someone else, great, and if not, no big deal.

我喜欢这个是我觉得它平衡了我们作为开发人员在编写代码时的便利性和我们包的任何潜在用户的便利性。这是一个简单的 CLI，所以除了我们之外，可能没有人会使用这段代码。一方面，我们发现为了我们自己的开发目的，拆分第二个包而不是将所有内容都放入包 main 更方便。但另一方面，我们并没有投入太多资金来创建一个充满功能和接口的大型公共 API，我们只是推测这些功能和接口可能在某一天对某人有用。这是我们需要的代码，如果它对其他人有帮助，那就太好了，如果没有，也没什么大不了的。

[My final commit](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910) has **197 additions** and **197 deletions**, meaning that for all my writing about different abstractions we may or may not want to use, in the end, **the code is the exact same length** as written by me as it was when written by Eryb. (Okay, yes, I may have noticed I was three lines away and squeezed a little to make it the exact same.) My code actually handles several error conditions that were overlooked  by the original demo, so we're actually doing a little bit more work  than the original app did. We’ve added just enough architecture to make  our CLI more robust and extensible without actually doing any more work  than normal.

[我的最终提交](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910) 有 **197 个添加**和 **197 个删除**，这意味着我所有关于不同抽象的文章我们可能想也可能不想使用，最终，**代码的长度与我编写的代码和 Eryb 编写的代码长度完全相同。 （好吧，是的，我可能已经注意到我在三行之外并挤压了一点以使其完全相同。）我的代码实际上处理了原始演示忽略的几个错误条件，所以我们实际上做了一点比原始应用程序所做的工作更多。我们添加了刚好足够的架构，使我们的 CLI 更加健壮和可扩展，而实际上并没有做比平常更多的工作。

My code is designed to be easier to test… but I  haven’t actually written any tests yet. Why? Because for a simple CLI if you’re not following [TDD](https://en.wikipedia.org/wiki/Test-driven_development), you may not want to take the time to write tests. The idea is to leave  room open so that if you want or need to add tests later, there are  clear and easy processes for doing so.

我的代码被设计得更容易测试……但我实际上还没有编写任何测试。为什么？因为对于简单的 CLI，如果您不遵循 [TDD](https://en.wikipedia.org/wiki/Test-driven_development)，您可能不想花时间编写测试。我们的想法是留出空间，以便以后如果您想要或需要添加测试，有清晰且简单的过程来执行此操作。

The point of all of this,  again, is not that my method is the only way to write a CLI in Go. For  simple CLIs, anything will work, including having a single main.go file  with a bunch of global variables and a `check` function. I  argue however that unlike creating a hexagonal architecture with many  layers of interfaces just in case you might need them in the future, my method is **no more work** than the simple method while being **much better for extending in the future **.

同样，所有这一切的重点并不是我的方法是在 Go 中编写 CLI 的唯一方法。对于简单的 CLI，任何东西都可以工作，包括拥有一个包含一堆全局变量和一个 `check` 函数的单个 main.go 文件。然而，我认为与创建具有多层接口的六边形架构不同，以防万一您将来可能需要它们，我的方法**没有比简单方法更多的工作**，同时**更适合将来扩展**。

------

If you like the way I describe writing Go CLIs but want a template to  start off of instead of writing the boilerplate yourself, check out [go-cli](https://github.com/carlmjohnson/go-cli),a [ Github Template Repo](https://github.com/carlmjohnson/go-cli/generate) I have created that acts as a `cat` clone. It follows the patterns described in this post and is pretty easy to adapt into new tools.

如果您喜欢我描述的编写 Go CLI 的方式，但想要一个模板而不是自己编写样板，请查看 [go-cli](https://github.com/carlmjohnson/go-cli)，一个[ Github 模板库](https://github.com/carlmjohnson/go-cli/generate) 我已经创建了一个“cat”克隆。它遵循本文中描述的模式，并且很容易适应新工具。

------

## Bash Double Bonus

## Bash 双倍奖励

The [original post](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) ends with a trick for downloading multiple comics serially in Bash, so  here's a bonus trick for downloading the comics in parallel and waiting  for them all to finish:

[原帖](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) 以在Bash中连续下载多部漫画的技巧结束，所以这里有一个额外的技巧，可以并行下载漫画并等待它们全部完成：

```bash
$ for i in {1..10};do ./go-grab-xkcd -n $i -s & done;wait
```


I hope this helped you think about how to structure your CLIs, and you have fun experimenting with finding your own patterns! 

我希望这能帮助您思考如何构建您的 CLI，并且您在尝试寻找自己的模式时会很开心！
