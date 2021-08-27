# An intro to Go for non-Go developers

# 非 Go 开发者的 Go 介绍

June 2020

2020 年 6 月

> Summary: I’ve presented an introduction to Go a few times for  developers who are new to the language – this is that talk serialized as a technical article. It looks at why you might want to use Go, and  gives a brief overview of the standard library and the language itself.

> 总结：我已经为刚接触 Go 语言的开发人员介绍了几次 Go 的介绍——这是一篇连载为技术文章的演讲。它着眼于您可能想要使用 Go 的原因，并简要概述了标准库和语言本身。

A few years ago I [learned Go](https://benhoyt.com/writings/learning-go/)by porting the server for my [Gifty Weddings](https://giftyweddings.com/) side gig from Python to Go. It was a fun way to learn the language, and took me about “two weeks of bus commutes” to learn Go at a basic level  and port the code.

几年前，我通过将我的 [Gifty Weddings](https://giftyweddings.com/) 的服务器从 Python 移植到走。这是一种学习语言的有趣方式，我花了大约“两周的公交车通勤时间”来基本学习 Go 并移植代码。

Since then, I've really enjoyed working with the language, and have  used it extensively at work as well as on side projects like [GoAWK](https://benhoyt.com/writings/goawk/)and [zztgo](https ://benhoyt.com/writings/zzt-in-go/). Go usage at [Compass.com](https://www.compass.com/), my current workplace, has grown significantly in the time I’ve been  there – around half of our 200 plus services are written in Go.

从那以后，我真的很喜欢使用这种语言，并且在工作中以及像 [GoAWK](https://benhoyt.com/writings/goawk/)和 [zztgo](https ://benhoyt.com/writings/zzt-in-go/）。在我目前的工作场所 [Compass.com](https://www.compass.com/) 上，Go 的使用量在我工作后显着增长——我们 200 多个服务中约有一半是用 Go 编写的。

This article describes what I think are some of the great things  about Go, gives a very brief overview of the standard library, and then  digs into the core language. But if you just want a feel for what real  Go code looks like, skip to the [HTTP server examples](https://benhoyt.com/writings/go-intro/#http-server-examples).

这篇文章描述了我认为 Go 的一些伟大之处，对标准库进行了非常简要的概述，然后深入研究了核心语言。但是，如果您只是想感受一下真正的 Go 代码是什么样的，请跳到 [HTTP 服务器示例](https://benhoyt.com/writings/go-intro/#http-server-examples)。

## Why Go?

## 为什么要去？

As the following [Google Trends chart](https://trends.google.com/trends/explore?date=2010-01-012020-06-09&q=golang&hl=en-US) shows, Go has become very popular over the past few years, partly  because of the simplicity of the language, but perhaps more importantly  because of the excellent tooling.

正如下面的 [Google 趋势图表](https://trends.google.com/trends/explore?date=2010-01-012020-06-09&q=golang&hl=en-US) 所示，Go 已经变得非常流行。过去几年，部分是因为语言的简单性，但也许更重要的是因为出色的工具。

![Google Trends data for "golang" from 2010 to 2020](https://benhoyt.com/images/golang-trend.png)

Here are some of the reasons I enjoy programming in Go (and why you might like it too):

以下是我喜欢用 Go 编程的一些原因（以及你可能也喜欢它的原因）：

- **Small and simple core language.** Go feels similar in size to C, with a very readable [language spec](https://golang.org/ref/spec)that's only about 50 pages long (compared to the [Java spec's](https://docs.oracle.com/javase/specs/jls/se12/jls12.pdf) 770 pages). This makes it easy to learn or teach to others.
- **High quality standard library**, especially for servers and network tasks. More details [below](https://benhoyt.com/writings/go-intro/#the-standard-library).
- **First class concurrency** with *goroutines* (like threads, but lighter) and the `go` keyword to start a goroutine, *channels* for communicating between them, and a runtime whose scheduler coordinates all this.
- **Compiles to native code**, producing easy-to-deploy binaries on all the major platforms.
- **Garbage collection** that doesn’t require knob-tweaking ([optimized](https://blog.golang.org/ismmkeynote)for low latency).
- **Statically typed**, but has type inference to avoid a lot of “type stuttering”.
- **Great documentation** that is succinct and includes many runnable examples.
- **Excellent tooling.** Just type `go build` to build your project, `go test` to find and run your tests, etc. There’s CPU and memory profiling, code coverage, and cross compilation – all without external tooling.
- **Fast compile times.** The language was designed from day one with fast compile times in mind. In fact, co-creator Rob Pike [jokes](https://www.informit.com/articles/article.aspx?p=1623555) that “Go was conceived while waiting for a big [C++] compilation.”
- **Very stable** language and library, with a strict [compatibility promise](https://golang.org/doc/go1compat) that all Go 1 programs will run unchanged on later versions of Go 1.x.
- **Desired.** According to StackOverflow's 2019 survey, it's the [third most wanted](https://insights.stackoverflow.com/survey/2019#most-loved-dreaded-and-wanted) programming language, so it's easy to hire developers who want to use it.
- **Heavily used in cloud tools.** Docker and Kubernetes are written in Go, and Dropbox, Digital Ocean, Cloudflare, and many other companies use it extensively.

- **小而简单的核心语言。** Go 感觉在大小上与 C 相似，具有非常易读的 [语言规范](https://golang.org/ref/spec)，只有大约50 页长（与[Java 规范](https://docs.oracle.com/javase/specs/jls/se12/jls12.pdf) 770 页）。这使得学习或教给他人变得容易。
- **高质量的标准库**，特别适用于服务器和网络任务。更多详情[下文](https://benhoyt.com/writings/go-intro/#the-standard-library)。
- **一级并发**，带有 *goroutines*（类似于线程，但更轻量级）和 `go` 关键字来启动一个 goroutine，*channels* 用于它们之间的通信，以及一个运行时，其调度程序协调所有这些。
- **编译为本机代码**，在所有主要平台上生成易于部署的二进制文件。
- **垃圾收集**，不需要旋钮调整（[优化]（https://blog.golang.org/ismmkeynote）以实现低延迟）。
- **静态类型**，但具有类型推断以避免大量“类型卡顿”。
- **优秀的文档**，简洁并包含许多可运行的示例。
- **优秀的工具。** 只需输入 `go build` 来构建你的项目，输入 `go test` 来查找和运行你的测试等等。还有 CPU 和内存分析、代码覆盖率和交叉编译——所有这些都不需要外部工具。
- **快速编译时间。** 该语言从一开始就考虑到了快速编译时间。事实上，共同创作者 Rob Pike [笑话](https://www.informit.com/articles/article.aspx?p=1623555) 说“Go 是在等待大型 [C++] 编译时构思出来的。”
- **非常稳定**的语言和库，具有严格的[兼容性承诺](https://golang.org/doc/go1compat)，所有 Go 1 程序将在 Go 1.x 的更高版本上保持不变。
- **Desired.** 根据 StackOverflow 的 2019 年调查，它是 [第三大最想要的](https://insights.stackoverflow.com/survey/2019#most-loved-dreaded-and-wanted) 编程语言，所以它是易于雇用想要使用它的开发人员。
- **大量用于云工具。** Docker 和 Kubernetes 是用 Go 编写的，Dropbox、Digital Ocean、Cloudflare 和许多其他公司广泛使用它。

## The standard library 

## 标准库

Go’s [standard library](https://golang.org/pkg/)is  extensive, cross-platform, and well documented. Similar to Python, Go  comes with “batteries included”, so you can build useful servers and CLI tools right away, without any third party dependencies. Here are some  of the highlights (biased towards what I’ve used):

Go 的 [标准库](https://golang.org/pkg/) 是广泛的、跨平台的，并且有据可查。与 Python 类似，Go 带有“内置电池”，因此您可以立即构建有用的服务器和 CLI 工具，而无需任何第三方依赖。以下是一些亮点（偏向于我使用过的）：

- Input/output: [OS calls](https://golang.org/pkg/os/),files and directories, [buffered I/O](https://golang.org/pkg/bufio/).
- HTTP: a production-ready [client and server](https://golang.org/pkg/net/http/), TLS, HTTP/2, simple routing, URL and cookie parsing.
- Strings: [all the basics](https://golang.org/pkg/strings/),handling of [raw bytes](https://golang.org/pkg/bytes/), [unicode conversions](https://golang.org/pkg/unicode/).
- Encodings: [JSON](https://golang.org/pkg/encoding/json/),[XML](https://golang.org/pkg/encoding/xml/), [CSV](https://golang.org/pkg/encoding/csv/), [base64](https://golang.org/pkg/encoding/base64/),[hex](https://golang.org/pkg/encoding/hex /), [binary](https://golang.org/pkg/encoding/binary/), more.
- Templating: simple but powerful [text](https://golang.org/pkg/text/template/)and auto-escaped [HTML](https://golang.org/pkg/html/template/) templates.
- Time: simple API but well thought out [date and time](https://golang.org/pkg/time/) functions.
- Regular expressions: a non-backtracking [regexp library](https://golang.org/pkg/regexp/).
- Sorting: generic collection [sorting functions](https://golang.org/pkg/sort/).
- Databases: the [`database/sql`](https://golang.org/pkg/database/sql/) interface, with specific implementations left up to third party libraries.
- Crypto libraries: secure and fast [implementations](https://golang.org/pkg/crypto/) of AES, block ciphers, cryptographic hashes, etc.
- Image: read and write [JPEG](https://golang.org/pkg/image/jpeg/),[PNG](https://golang.org/pkg/image/png/), and [GIF] (https://golang.org/pkg/image/gif/), perform basic [compositing](https://golang.org/pkg/image/draw/).
- Big numbers: arbitrary-precision [int and float](https://golang.org/pkg/math/big/).
- Archives and compression: [tar](https://golang.org/pkg/archive/tar/),[zip](https://golang.org/pkg/archive/zip/), [gzip](https://golang.org/pkg/compress/gzip/), [bzip2](https://golang.org/pkg/compress/bzip2/), etc.
- Simple command-line [flag](https://golang.org/pkg/flag/) library.
- Go source code tools: [parser](https://golang.org/pkg/go/parser/),[AST](https://golang.org/pkg/go/ast/), code [formatting] (https://golang.org/pkg/go/printer/).
- Reflection: powerful run-time [reflection support](https://golang.org/pkg/reflect/).

- 输入/输出：[OS 调用](https://golang.org/pkg/os/)、文件和目录、[缓冲I/O](https://golang.org/pkg/bufio/)。
- HTTP：生产就绪 [客户端和服务器](https://golang.org/pkg/net/http/)、TLS、HTTP/2、简单路由、URL 和 cookie 解析。
- 字符串：[所有基础知识](https://golang.org/pkg/strings/)、[原始字节](https://golang.org/pkg/bytes/)、[unicode 转换](https ://golang.org/pkg/unicode/）。
- 编码：[JSON](https://golang.org/pkg/encoding/json/)、[XML](https://golang.org/pkg/encoding/xml/)、[CSV](https:///golang.org/pkg/encoding/csv/), [base64](https://golang.org/pkg/encoding/base64/),[hex](https://golang.org/pkg/encoding/hex /), [二进制](https://golang.org/pkg/encoding/binary/), 更多。
- 模板：简单但功能强大的 [text](https://golang.org/pkg/text/template/)和自动转义的 [HTML](https://golang.org/pkg/html/template/) 模板。
- 时间：简单的 API 但经过深思熟虑的 [日期和时间](https://golang.org/pkg/time/) 函数。
- 正则表达式：非回溯 [regexp 库](https://golang.org/pkg/regexp/)。
- 排序：泛型集合[排序功能](https://golang.org/pkg/sort/)。
- 数据库：[`database/sql`](https://golang.org/pkg/database/sql/) 接口，具体实现留给第三方库。
- 加密库：AES、分组密码、加密哈希等的安全和快速[实现](https://golang.org/pkg/crypto/)。
- 图像：读取和写入 [JPEG](https://golang.org/pkg/image/jpeg/)、[PNG](https://golang.org/pkg/image/png/)和 [GIF] (https://golang.org/pkg/image/gif/)，执行基本的[合成](https://golang.org/pkg/image/draw/)。
- 大数字：任意精度 [int 和 float](https://golang.org/pkg/math/big/)。
- 存档和压缩：[tar](https://golang.org/pkg/archive/tar/)、[zip](https://golang.org/pkg/archive/zip/)、[gzip](https://golang.org/pkg/compress/gzip/)、[bzip2](https://golang.org/pkg/compress/bzip2/) 等。
- 简单的命令行 [flag](https://golang.org/pkg/flag/) 库。
- Go 源码工具：[parser](https://golang.org/pkg/go/parser/)、[AST](https://golang.org/pkg/go/ast/)、代码[格式化] （https://golang.org/pkg/go/printer/）。
- 反射：强大的运行时 [反射支持](https://golang.org/pkg/reflect/)。

In terms of third party packages, typical Go philosophy is almost the opposite of JavaScript’s approach of pulling in `npm` packages left, right, and center. Russ Cox (tech lead of the Go team at Google) talks about [our software dependency problem](https://research.swtch.com/deps),and Go co-creator Rob Pike [likes to say](https:/ /go-proverbs.github.io/), “A little copying is better than a little dependency.” So it’s fair to  say that most Gophers are pretty conservative about using third party  libraries.

就第三方包而言，典型的 Go 哲学几乎与 JavaScript 将 npm 包向左、向右和居中拉入的方法相反。 Russ Cox（Google Go 团队的技术负责人）谈论 [我们的软件依赖问题](https://research.swtch.com/deps)，以及Go 的共同创造者 Rob Pike [喜欢说](https:/ /go-proverbs.github.io/)，“一点点复制比一点点依赖要好。”所以可以说大多数 Gophers 对使用第三方库相当保守。

That said, since I originally wrote this talk, the Go team has designed and built [modules](https://blog.golang.org/using-go-modules), the Go team's official answer to how you should manage and version -pin  your dependencies. I’ve found it pleasant to use, and it works with all  the normal `go` sub-commands.

也就是说，自从我最初写这篇演讲以来，Go 团队已经设计并构建了 [modules](https://blog.golang.org/using-go-modules)，Go 团队对你应该如何管理和版本的官方回答- 固定您的依赖项。我发现它使用起来很愉快，并且它适用于所有正常的 `go` 子命令。

## Language features

## 语言特性

So let’s dig in to what Go itself looks like, and walk through the language proper.

因此，让我们深入了解 Go 本身的样子，并仔细研究一下这门语言。

### Hello world

###  你好，世界

Go has a C-like syntax, mandatory braces, and no semicolons (except  in the formal grammar). Projects are structured via imports and packages – compilation units that consist of a directory with one or more `.go` files in it. Here’s what a “hello world” looks like:

Go 具有类似 C 的语法、强制大括号，并且没有分号（正式语法除外）。项目是通过导入和包构建的——编译单元由一个目录组成，其中包含一个或多个 `.go` 文件。这是“hello world”的样子：

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```


### Somewhat controversial features

### 有些争议的功能

Go has a few things that put some people off when they first see the  language, but turn out to be quite nice once you get used to them. 

Go 有一些东西会让一些人在第一次看到这门语言时感到不舒服，但是一旦你习惯了它们就会变得非常好。

The first one is code formatting: you just run [`go fmt`](https://blog.golang.org/gofmt)and it puts your braces and whitespace (and tabs!) where it knows they  should go. It’s a great way to avoid style wars and just get on with  consistently-formatted code.

第一个是代码格式化：你只需运行 [`go fmt`](https://blog.golang.org/gofmt)，它就会把你的大括号和空格（和制表符！）放在它知道它们应该去的地方。这是避免风格战争并继续使用一致格式的代码的好方法。

Capitalized names are public (“exported”), lower case names are  private to the package. This one seems very strange at first, but the  rule is easy to understand, and cuts down on the Java-esque `public static void` keyword noise. There’s no need for a `public` keyword at all – here’s how it looks:

大写名称是公开的（“导出”），小写名称是包私有的。这个乍一看很奇怪，但规则很容易理解，并且减少了 Java 式的“public static void”关键字噪音。根本不需要“public”关键字——它是这样的：

```
package people

type Person struct {
    Name      string // fields Name and Age are exported ("public")
    Age       int
    hairColor color  // hairColor is not exported ("private")
}

func New() *Person { ... } // New is exported

func doThing() { ... } // doThing is not exported
```


Another thing that gets some developers: warnings are errors! Go’s  built-in tools have very few options, and things that would be warnings  in other compilers are errors in Go (or put another way, there are no  warnings).

另一件事让一些开发人员感到：警告是错误！ Go 的内置工具很少有选择，而在其他编译器中会是警告的东西在 Go 中是错误（或者换句话说，没有警告）。

So even things like unused local variables or unused imports are  compile errors – this can be slightly annoying during development, but  it keeps the code clean, and avoids developers fighting over which  compiler warnings to turn on.

因此，即使像未使用的局部变量或未使用的导入这样的事情也是编译错误——这在开发过程中可能会有点烦人，但它可以保持代码干净，并避免开发人员为打开哪些编译器警告而争吵。

### Rather controversial features

### 颇有争议的功能

There are a few features – or rather lack of features – that are even more controversial, namely: Go’s lack of exceptions, and its lack of  user-defined generics.

有一些特性——或者更确切地说是缺乏特性——更具争议性，即：Go 缺乏异常，缺乏用户定义的泛型。

Go doesn’t have exceptions in the traditional sense. From the beginning, the mantra has been that [errors are values](https://blog.golang.org/errors-are-values) and should be explicitly passed around, returned, and handled like any other value. So instead of raising a `FileNotFound` exception, you test an error value:

Go 没有传统意义上的例外。从一开始，口头禅就是 [errors are values](https://blog.golang.org/errors-are-values) 应该像任何其他值一样明确地传递、返回和处理。因此，不是引发 `FileNotFound` 异常，而是测试错误值：

```
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open file f
```


This does make the code more verbose (put `if err != nil` on speed dial), but it does have the advantage of making error handling explicit at each level. You can choose to add context, log the error,  turn it into a higher-level error object, or even throw it away – but  you need to explicitly deal with it.

这确实使代码更加冗长（将 `if err != nil` 放在快速拨号上），但它确实具有在每个级别明确错误处理的优点。您可以选择添加上下文、记录错误、将其转换为更高级别的错误对象，甚至将其丢弃——但您需要明确处理它。

The second thing Go is often criticized for not having is user-defined generics. So you can’t define your own type-safe `OrderedMap<int>`. But because it does have statically-typed generics for the built-in `slice` and `map` types, you can go pretty far without feeling the pain.

Go 经常被批评的第二件事是用户定义的泛型。所以你不能定义自己的类型安全 `OrderedMap<int>`。但是因为它确实有内置的 `slice` 和 `map` 类型的静态类型泛型，你可以走得很远而不会感到痛苦。

The other thing to note is that generics *are* being worked on: the Go team just [wants to add them](https://blog.golang.org/why-generics)in a way that's very Go-like and that counts the cost, rather than a bolted on addition. There's a [draft proposal](https://github.com/golang/proposal/blob/master/design/go2draft-contracts.md), an [experimental implementation](https://go-review.googlesource.com/c/go/+/187317/), and even a recent type theory paper on the subject called [Featherweight Go](https://arxiv.org/abs/2005.11710). So I wouldn’t be surprised if we saw Go shipping with a form of generics in the next 12-18 months.

另一件需要注意的事情是泛型 * 正在* 正在研究中：Go 团队只是 [想添加它们](https://blog.golang.org/why-generics)以一种非常类似于 Go 的方式计算成本，而不是盲目的加法。有一个[草案提案](https://github.com/golang/proposal/blob/master/design/go2draft-contracts.md)，一个[实验实施](https://go-review.googlesource.com/c/go/+/187317/)，甚至是最近一篇关于该主题的类型理论论文，名为 [Featherweight Go](https://arxiv.org/abs/2005.11710)。因此，如果我们在接下来的 12 到 18 个月内看到 Go 带有某种形式的泛型，我不会感到惊讶。

Okay, so enough about what Go doesn’t have. Let’s look at what features it does have (many of them unique).

好的，关于 Go 没有的东西就够了。让我们看看它有哪些功能（其中许多是独一无二的）。

### Succinct type inference

### 简洁的类型推断

Go has succinct type inference for declaring variables with the `:=` operator, called “short variable declarations”. Type inference makes it feel a bit more like a scripting language, as there’s less (ahem) *typing*. Here’s what it looks like:

Go 有简洁的类型推断，用于使用 `:=` 运算符声明变量，称为“短变量声明”。类型推断让它感觉更像是一种脚本语言，因为更少（咳咳）*打字*。这是它的样子：

```
package main

import "fmt"

// Output: 3 4 hello 5
func main() {
    var i int = 3
    j := 4          // j is an int
    s := "hello"    // s is a string
    a := add(2, 3)  // a is an int
    fmt.Println(i, j, s, a)
}

func add(x, y int) int {
    return x + y
}
```


On the other hand, Go doesn’t do any automatic type coercion, even  between integers of different widths or signed-ness. To quote the [FAQ answer](https://golang.org/doc/faq#conversions) comparing Go’s approach to C: 

另一方面，Go 不会进行任何自动类型强制，即使在不同宽度或符号的整数之间也是如此。引用 [FAQ answer](https://golang.org/doc/faq#conversions) 比较 Go 与 C 的方法：

> The convenience of automatic conversion between numeric types in C  is outweighed by the confusion it causes. When is an expression  unsigned? How big is the value? Does it overflow? Is the result  portable, independent of the machine on which it executes?

> C 中数字类型之间自动转换的便利性被它引起的混乱所抵消。什么时候表达式是无符号的？价值有多大？会溢出吗？结果是否可移植，独立于执行它的机器？

### For loops and `range`

### For 循环和 `range`

Go has a single loop keyword, `for`, that’s used for while loops, old school C-style loops, and `range` loops (Go’s “for each”). When `range` looping over a slice or map, Go gives you the index (or map key) as the first item, and the value as the second item.

Go 有一个循环关键字 `for`，用于 while 循环、老派 C 风格循环和 `range` 循环（Go 的“for each”）。当 `range` 在切片或映射上循环时，Go 将索引（或映射键）作为第一项，将值作为第二项。

Here are some loopy examples – so far nothing out of the ordinary:

以下是一些乱七八糟的例子——到目前为止没有什么不寻常的：

```
// C style "for"
for i := 0;i < 10;i++ {
    fmt.Println(i)
}

// Like "while"
for safe.IsLocked() {
    time.Sleep(5 * time.Second)
}

// Loop through elements of array or slice
for index, person := range people {
    fmt.Println(index, person)
}

// If you don't care about the index
for _, person := range people { ... }

// Loop through keys/values of a map
for word, count := range counts { ... }
```


### A slice is very nice

### 切片非常好

In Go, a *slice* is a reference to part of an array – the  internal representation is very simple: a data pointer, a length, and a  capacity. Slices are generic, so you can have a slice of `float64`, denoted as `[]float64`, or a slice of `Person` structs, `[]Person`.

在 Go 中，*slice* 是对数组一部分的引用——内部表示非常简单：数据指针、长度和容量。切片是通用的，所以你可以有一个 `float64` 的切片，表示为 `[]float64`，或者一个 `Person` 结构的切片，`[]Person`。

You can “slice” slices with Python-like syntax, for example `slice[:5]` to return a new slice viewing the first five elements. The new slice  still refers to the same backing array, so it’s as efficient as dealing  with pointers, but memory-safe – the runtime prevents you from walking  off the end of a slice or doing other nasty things.

您可以使用类似 Python 的语法“切片”切片，例如 `slice[:5]` 以返回查看前五个元素的新切片。新切片仍然引用相同的后备数组，因此它与处理指针一样有效，但内存安全 - 运行时可防止您离开切片的末尾或做其他讨厌的事情。

There’s a built-in generic `append()` function that appends a single element to the slice and returns the new slice. If the backing array (the *capacity*) isn’t big enough, it’ll allocate a new array of double the size and copy the elements over.

有一个内置的通用 `append()` 函数可以将单个元素附加到切片并返回新的切片。如果后备数组（*容量*）不够大，它将分配一个双倍大小的新数组并复制元素。

Here’s what slices look like:

这是切片的样子：

```
// Create array and slice pointing into it
nums := []int{3, 4, 5, 6}

// Slice the slice
fmt.Println(nums[1:3]) // Output: [4 5]
fmt.Println(nums[:2]) // Output: [3 4]
fmt.Println(nums[2:]) // Output: [5 6]

// Append: may reallocate underlying array
nums = append(nums, 7, 8)
fmt.Println(nums) // Output: [3 4 5 6 7 8]
```


Slice functionality is pretty minimalist, and one thing I missed (coming from Python) is list comprehensions. Why do I need a `for` loop and an `if` statement just to filter a few things out of my list? I once asked a  member of the Go team why such features were missing, and he said that  because Go is a “systems language” they want you to have control over  memory allocation. As an example, you can use [`make()`](https://golang.org/pkg/builtin/#make) to pre-allocate a slice’s backing array for efficiency.

切片功能非常简约，我错过的一件事（来自 Python）是列表推导式。为什么我需要一个 `for` 循环和一个 `if` 语句只是为了从我的列表中过滤一些东西？我曾经问过 Go 团队的一位成员，为什么缺少这些功能，他说因为 Go 是一种“系统语言”，他们希望您可以控制内存分配。例如，您可以使用 [`make()`](https://golang.org/pkg/builtin/#make) 预先分配切片的后备数组以提高效率。

### Maps

### 地图

A Go `map` is an unordered hash table mapping keys to values. Like slices, they’re generic and type-safe, so you can have (for example) a `map[string]int`, which means “map of string keys to int values”.

Go `map` 是一个将键映射到值的无序哈希表。像切片一样，它们是泛型且类型安全的，因此您可以拥有（例如）一个 `map[string]int`，意思是“字符串键到 int 值的映射”。

The `map` data type provides get, set, delete, existence test, and iteration. Just like slices, you can control memory allocation with `make()` using a “size hint”.

`map` 数据类型提供获取、设置、删除、存在测试和迭代。就像切片一样，您可以使用“大小提示”通过 `make()` 控制内存分配。

There's much [more to say](https://golang.org/doc/effective_go.html#maps)about maps, and you can [read about their implementation](https://dave.cheney.net/2018/05 /29/how-the-go-runtime-implements-maps-efficiently-without-generics), but here's a taste of them in code:

关于地图有很多[更多要说的](https://golang.org/doc/effective_go.html#maps)，您可以[阅读它们的实施](https://dave.cheney.net/2018/05 /29/how-the-go-runtime-implements-maps-efficiently-without-generics），但这里有代码的味道：

```
phrase := "the foo foo bar the foo"

counts := make(map[string]int)
for _, word := range strings.Fields(phrase) {
    counts[word]++
}

fmt.Println(counts)
// Output: map[foo:3 bar:1 the:2]

// map literal
maths := map[string]float64{
    "pi":  3.14,
    "tau": 6.28,
}
```


### Pointers, but safe

### 指针，但安全

Go has pointers, but unlike in C and C++, they’re safe. You can’t  point to memory that doesn’t exist, and the runtime prevents you from  dereferencing a nil pointer. In fact, there’s no pointer arithmetic at  all – if you want to index into something, you have to use safe slices,  or fall back to the low-level `unsafe` package (I’ve never needed it).

Go 有指针，但与 C 和 C++ 不同，它们是安全的。您不能指向不存在的内存，并且运行时会阻止您取消引用 nil 指针。事实上，根本没有指针算法——如果你想索引到某个东西，你必须使用安全切片，或者退回到低级的 `unsafe` 包（我从来不需要它）。

Pointers use `*` and `&` syntax like C, with `*` fetching the value at a pointer’s address, and `&` taking the address of a variable. 

指针使用类似 C 的 `*` 和 `&` 语法，`*` 获取指针地址处的值，而 `&` 获取变量地址。

One of the nice syntactic things is that there’s no C-style `->` operator: to dereference a struct pointer and fetch a field, you use `.` as well. Here are some examples:

不错的语法之一是没有 C 风格的 `->` 运算符：要取消引用结构指针并获取字段，您也可以使用 `.`。这里有些例子：

```
p := new(Person) // p is a "pointer to Person"
p.Name = "Joe Bloggs"
p.Age = 42
pers := *p // dereference p back to Person

// More succinct alternatives
p = &Person{"Joe Bloggs", 42}
p = &Person{Name: "Joe Bloggs", Age: 42}
pers = Person{"Joe Bloggs", 42}
p = &pers
```


### Defer

###  推迟

Go has a unique keyword `defer` which executes the given function call just before the current function returns (or exits due to a runtime “panic”). If `defer` is called multiple times, the functions are called in last-defer-first  order. It’s used for resource clean-up in place of things like [RAII](https://en.wikipedia.org/wiki/Resource_acquisition_is_initialization) in C++ or the `with` statement in Python.

Go 有一个独特的关键字 `defer`，它在当前函数返回（或由于运行时“恐慌”而退出）之前执行给定的函数调用。如果`defer` 被多次调用，函数将按last-defer-first 的顺序调用。它用于资源清理，代替 C++ 中的 [RAII](https://en.wikipedia.org/wiki/Resource_acquisition_is_initialization) 或 Python 中的 `with` 语句。

As far as I know, `defer` is a control flow statement unique to Go, and fits well with its explicit approach to error handling. You can [read more about it](https://blog.golang.org/defer-panic-and-recover), but here’s a very simple example of a common task – opening and closing a file:

据我所知，`defer` 是 Go 独有的控制流语句，非常适合其明确的错误处理方法。您可以[阅读更多相关信息](https://blog.golang.org/defer-panic-and-recover)，但这里有一个非常简单的常见任务示例——打开和关闭文件：

```
f, err := os.Open("file")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
// read from f
```


### Goroutines

### 协程

*Goroutines* are Go’s concurrency mechanism: they’re like  threads, but much lighter weight – you can easily have 100,000 or even a million goroutines alive at once. The Go runtime schedules the  Goroutines, waking them up and executing them on operating system  threads as needed (when I/O is ready, for example).

*Goroutines* 是 Go 的并发机制：它们就像线程，但重量要轻得多——你可以轻松地同时拥有 100,000 甚至 100 万个 goroutines。 Go 运行时调度 Goroutines，唤醒它们并根据需要在操作系统线程上执行它们（例如，当 I/O 准备就绪时）。

One of the neat things about Go’s concurrency model is that all the  standard library functions have simple, synchronous APIs – and if you  need concurrency, you use goroutines explicitly. This avoids the problem with [“colored functions”](https://journal.stuffwithstuff.com/2015/02/01/what-color-is-your-function/) – the two sets of APIs that some languages have, one for async and one for synchronous functions.

Go 并发模型的优点之一是所有标准库函数都有简单的同步 API——如果你需要并发，你可以明确地使用 goroutines。这避免了 [“彩色函数”](https://journal.stuffwithstuff.com/2015/02/01/what-color-is-your-function/) 的问题——某些语言具有的两组 API，一种用于异步功能，一种用于同步功能。

To kick off a goroutine, just write `go backgroundFunc()`, and Go’s runtime will kick off `backgroundFunc` on a new goroutine. Here’s a simple example of a handler function that  records a user signing up and then sends them an email in the background (this is similar to real code I use in my [side gig](https://giftyweddings.com/)):

要启动一个 goroutine，只需编写 `go backgroundFunc()`，Go 的运行时就会在一个新的 goroutine 上启动 `backgroundFunc`。这是一个处理程序函数的简单示例，它记录用户注册，然后在后台向他们发送电子邮件（这类似于我在 [side gig](https://giftyweddings.com/) 中使用的真实代码）：

```
func ProcessSignup(u *User) {
    u.SignedUpAt = time.Now()
    u.Save(db)
    go SendEmail(u.email, "Thanks for signing up!", "signup.html")
}
```


### Channels

### 频道

Starting a goroutines doesn’t return a promise or goroutine ID – if  you want to communicate between goroutines or signal that work is done,  you have to explicitly use *channels*. Channels are Go’s main inter-goroutine communication mechanism, and as the [Go Proverb](https://go-proverbs.github.io/) says, “Don’t communicate by sharing memory, share memory by communicating.”

启动 goroutines 不会返回 promise 或 goroutine ID——如果你想在 goroutines 之间通信或发出工作完成的信号，你必须明确使用 *channels*。通道是 Go 的主要 goroutine 间通信机制，正如 [Go Proverb](https://go-proverbs.github.io/) 所说，“不要通过共享内存来通信，通过通信来共享内存。”

A channel is a type-safe and thread-safe queue that can communicate  data, but also synchronize things – a goroutine reading from a channel  will wait until another goroutine writes to it.

通道是一个类型安全和线程安全的队列，它可以传递数据，但也可以同步事物——从通道读取的 goroutine 将等待另一个 goroutine 写入它。

Here’s an example of parallelizing a simple “array sum” task – this  example almost certainly wouldn’t benefit from goroutines in practice,  but it gives you the idea:

这是一个并行化一个简单的“数组求和”任务的例子——这个例子在实践中几乎肯定不会从 goroutines 中受益，但它给了你一个想法：

```
func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c) // first half
    go sum(s[len(s)/2:], c) // second half

    // Receive both results from channel
    x, y := <-c, <-c

    fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // Send sum back to main
}
```


Channels are powerful constructs, and there's much to say about them  (buffered versus unbuffered, closed channels, etc), but I'll leave that  for [Effective Go](https://golang.org/doc/effective_go.html#channels).

通道是强大的构造，关于它们有很多要说的（缓冲与非缓冲、封闭通道等），但我会把它留给 [Effective Go](https://golang.org/doc/effective_go.html#channels ）。

### Types and methods 

### 类型和方法

Go supports user-defined types, and types can have methods, but there are no classes (some would say that Go is not a *classy* language). And there are `struct`s and interfaces (discussed below), but there’s no inheritance. All the  OOP goodness is done with composition – but there are tools such as [embedding](https://golang.org/doc/effective_go.html#embedding) that give you another approach.

Go 支持用户定义的类型，类型可以有方法，但没有类（有人会说 Go 不是 *classy* 语言）。并且有 `struct` 和接口（在下面讨论），但没有继承。所有 OOP 的优点都是通过组合完成的——但是有一些工具，例如 [embedding](https://golang.org/doc/effective_go.html#embedding) 可以为您提供另一种方法。

Methods defined on a type take a “receiver” argument, which is similar to `self` in Python and `this` in other languages. But they have a few unique properties (for example, receivers can be pointers or values). You can also name your receivers  whatever you want, though they’re typically named with the first letter  of the type in question.

定义在类型上的方法采用“接收者”参数，这类似于 Python 中的“self”和其他语言中的“this”。但是它们有一些独特的属性（例如，接收者可以是指针或值）。您也可以随意命名您的接收器，尽管它们通常以相关类型的第一个字母命名。

Here’s a simple, two-field struct with a `String` method:

这是一个带有 String 方法的简单的两字段结构：

```
type Person struct {
    Name string
    Age  int
}

func (p *Person) String() string {
    return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}

// Output: Bob (42 years)
func main() {
    p := &Person{"Bob", 42}
    fmt.Println(p.String()) // but .String() is optional;see below
}
```


### Interfaces

### 接口

Interfaces are a little different from those in other languages like Java, where you have to say `class MyThing implements ThatInterface` explicitly. In Go, if you define all of an interface’s methods on a  type, the type implicitly implements that interface, and you can use it  wherever the interface is called for – no `implements` keyword in sight.

接口与 Java 等其他语言中的接口略有不同，在 Java 中，您必须明确地说 `class MyThing 实现 ThatInterface`。在 Go 中，如果你在一个类型上定义了一个接口的所有方法，这个类型隐式地实现了这个接口，你可以在任何调用接口的地方使用它——看不到 `implements` 关键字。

Go's approach has often been called “static duck typing”, and it's a form of [structural typing](https://en.wikipedia.org/wiki/Structural_type_system)(TypeScript is another popular language that [uses](https:/ /medium.com/redox-techblog/structural-typing-in-typescript-4b89f21d6004) structural typing).

Go 的方法通常被称为“静态鸭子类型”，它是 [结构类型](https://en.wikipedia.org/wiki/Structural_type_system)的一种形式（TypeScript 是另一种 [使用]（https://en.wikipedia.org/wiki/Structural_type_system）的流行语言/medium.com/redox-techblog/structural-typing-in-typescript-4b89f21d6004) 结构类型）。

Interfaces are used everywhere in the standard library and in Go code. The two most common examples are the `Stringer` interface, which allows `Printf` and friends to generate a string version of a value, and the `io.Reader` and `io.Writer` interfaces, which allow you to treat files, HTTP servers, gzipped files, string buffers, etc, as reader or writer streams.

接口在标准库和 Go 代码中随处可见。两个最常见的例子是 Stringer 接口，它允许 Printf 和朋友生成一个值的字符串版本，以及 io.Reader 和 io.Writer 接口，它允许你处理文件， HTTP 服务器、gzipped 文件、字符串缓冲区等，作为读取器或写入器流。

Below are definitions for the `Stringer` and `Writer` interfaces – both very simple single-method interfaces (small  interfaces are very common in Go). You don’t actually have to define  these, but this code shows the syntax:

下面是 Stringer 和 Writer 接口的定义——都是非常简单的单方法接口（小接口在 Go 中很常见）。您实际上不必定义这些，但此代码显示了语法：

```
// Defined in package "fmt"
type Stringer interface {
    String() string
}

// Defined in package "io"
type Writer interface {
    Write(p []byte) (n int, err error)
}

// ...

func main() {
    p := &Person{"Bob", 42}
    fmt.Println(p.String())
    // Equivalent (Person implements Stringer, which Println looks for)
    fmt.Println(p)
}
```


It’s hard to overstate the importance of interfaces in Go. They’re used to make algorithms generic and functions testable. [Read more about them in Effective Go.](https://golang.org/doc/effective_go.html#interfaces)

很难夸大接口在 Go 中的重要性。它们用于使算法通用且功能可测试。 [在 Effective Go 中阅读更多关于它们的信息。](https://golang.org/doc/effective_go.html#interfaces)

## HTTP server examples

## HTTP 服务器示例

Before we go, here are a couple of small programs showing how easy it is to write HTTP servers in Go. And these aren't just toys – Go's `net/http` package is production-ready (unlike the built-in web servers that come  with many other languages, which always have to say “don't use in  production” on the tin ).

在我们开始之前，这里有几个小程序展示了在 Go 中编写 HTTP 服务器是多么容易。这些不仅仅是玩具——Go 的 `net/http` 包是生产就绪的（不像许多其他语言附带的内置网络服务器，它们总是必须在锡纸上说“不要在生产中使用” ）。

Here’s a very basic HTTP server with a single route that echos the `user` query string parameter. Note the use of the `http.ResponseWriter` as an `io.Writer` passed to `fmt.Fprintf`:

这是一个非常基本的 HTTP 服务器，它只有一个路由，它回显了 `user` 查询字符串参数。注意使用 `http.ResponseWriter` 作为传递给 `fmt.Fprintf` 的 `io.Writer`：

```
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("listening on port 8080")
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    user := r.URL.Query().Get("user")
    if user == "" {
        user = "world"
    }
    fmt.Fprintf(w, "Hello, %s", user)
}
```


As a slightly more advanced example, here we build an HTTP server with a custom regex-based router in a few lines of code.

作为一个稍微高级一点的例子，我们在这里用几行代码构建了一个带有自定义正则表达式路由器的 HTTP 服务器。

```
package main

import (
    "fmt"
    "net/http"
    "regexp"
)

type route struct {
    pattern *regexp.Regexp
    handler func(w http.ResponseWriter, r *http.Request, matches []string)
}

func home(w http.ResponseWriter, r *http.Request, matches []string) {
    fmt.Fprintf(w, "Home")
}

func user(w http.ResponseWriter, r *http.Request, matches []string) {
    user := matches[1]
    fmt.Fprintf(w, "User ID: %s", user)
}

func main() {
    routes := []route{
        {regexp.MustCompile(`^/$`), home},
        {regexp.MustCompile(`^/user/(\w+)$`), user},
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        for _, route := range routes {
            matches := route.pattern.FindStringSubmatch(r.URL.Path)
            if len(matches) >= 1 {
                route.handler(w, r, matches)
                return
            }
        }
        http.NotFound(w, r)
    })
    fmt.Println("listening on port 8080")
    http.ListenAndServe(":8080", nil)
}
```




## The `go` tool

## `go` 工具

Someone asked me recently what my favourite developer tool was. At  first I said, “maybe Sublime Text?” But then I changed my mind: I think  my (current) favourite developer tool is the `go` command. Without a Makefile, it can do all of the following, and it does them fast:

最近有人问我最喜欢的开发工具是什么。起初我说，“也许是 Sublime Text？”但后来我改变了主意：我认为我（当前）最喜欢的开发人员工具是 `go` 命令。如果没有 Makefile，它可以执行以下所有操作，而且速度很快：

```
go build          # build your project, produce a static executable
go run            # quick way to build and run, for development
go fmt            # format your .go files in the standard way
go test           # find and run your tests
go test -bench=.# run all your benchmarks too

go mod init                # initialize a "Go modules" project
go get github.com/foo/bar  # fetch and install the "bar" package
```


And there are many more commands – read the [full documentation](https://golang.org/pkg/cmd/go/).

还有更多命令——阅读[完整文档](https://golang.org/pkg/cmd/go/)。

But to me the most amazing thing of all is that if you set two environment variables, `GOOS` and `GOARCH`, and then run `go build`, Go cross-compiles your project for the given operating system and  architecture. Here’s a one-liner to create a deployable Linux binary on a macOS or Windows machine:

但对我来说最令人惊奇的是，如果你设置两个环境变量，`GOOS` 和 `GOARCH`，然后运行 `go build`，Go 会针对给定的操作系统和架构交叉编译你的项目。这是在 macOS 或 Windows 机器上创建可部署的 Linux 二进制文件的单行代码：

```
GOOS=linux GOARCH=amd64 go build
```


Isn’t that cool? Development hasn’t been this easy since [Turbo Pascal](https://en.wikipedia.org/wiki/Turbo_Pascal)…

这不是很酷吗？自从 [Turbo Pascal](https://en.wikipedia.org/wiki/Turbo_Pascal) 以来，开发并不容易……

## Wrapping up

##  包起来

There’s much more to say about Go and its ecosystem, but hopefully  this is a helpful introduction for those coming from other languages. To get started, I highly recommend the official [Go Tour](https://tour.golang.org/).For going deeper, read [Effective Go](https://golang.org/doc/effective_go.html) and then the excellent book [The Go Programming Language](https://www.gopl.io/).

关于 Go 及其生态系统还有很多话要说，但希望这对那些来自其他语言的人来说是一个有用的介绍。首先，我强烈推荐官方 [Go Tour](https://tour.golang.org/)。要深入了解，请阅读[Effective Go](https://golang.org/doc/effective_go.html) 和优秀的书 [The Go Programming Language](https://www.gopl.io/)。

Oh, and [write in Go!](https://www.youtube.com/watch?v=LJvEIjRBSDA) 

哦，还有 [用 Go 写！](https://www.youtube.com/watch?v=LJvEIjRBSDA)

