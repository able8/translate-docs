# Come Go With Me

## Why Network Engineers Should Learn Go

## 为什么网络工程师应该学习 Go

This post accompanies a lightning talk given at DevNet Create 2021 where I had 10 minutes to convince network engineers that they might like to learn Go in addition to Python, the lingua franca of network engineers. To be clear, I’m not suggesting you should necessarily learn it instead, but learn it alongside – who knows though, you might just love it!

这篇文章伴随着在 DevNet Create 2021 上的闪电演讲，我有 10 分钟的时间说服网络工程师除了 Python（网络工程师的通用语言）之外，他们可能还想学习 Go。需要明确的是，我并不是建议您一定要学习它，而是建议您一起学习——但谁知道呢，您可能会喜欢它！

First of all, I should just cover the title of the talk: _“Come Go With Me”_. Essentially I looked for any song with Go in it's title (and trust me, there are plenty), but after [seeing Mavis Staples on Jools Holland's Hootenanny from 2017](https://www.youtube.com/watch?v=FupcANEgGWI) I went with The Staple Singers, which has the full title _“If You're Ready (Come Go With Me)"_.

首先，我应该只覆盖演讲的标题：_“跟我走”_。基本上，我在寻找任何标题中带有 Go 的歌曲（相信我，有很多），但是在 [2017 年在 Jools Holland 的 Hootenanny 上看到 Mavis Staples](https://www.youtube.com/watch?v=FupcANegGWI)之后) 我和 The Staple Singers 一起去了，它的全名是 _“如果你准备好了（跟我一起走)”_。

In addition, this talk came about whilst I was part way through a good book, so many of the quotes are courtesy of that book: _[“Cloud Native Go”](https://www.oreilly.com/library/view/cloud-native-go/9781492076322)_ by Matthew A. Titmus.

此外，这个演讲是在我读完一本好书的时候开始的，所以很多引用都来自那本书：_[“Cloud Native Go”](https://www.oreilly.com/library/view/cloud-native-go/9781492076322)_ 作者：Matthew A. Titmus。

With that out of the way, I’ll get on with it.

有了这个，我会继续下去。

The talk was in the DevNet Create topic of “Interoperability & Quality” with an abstract as follows:

演讲在 DevNet Create 主题“互操作性和质量”中进行，摘要如下：

> An overview of why network engineers should learn the Go programming language and why it can improve the quality, performance and portability of their applications.

> 为什么网络工程师应该学习 Go 编程语言以及为什么它可以提高其应用程序的质量、性能和可移植性的概述。

So we’ll try and pick on some of these points as we go through.

因此，我们将在经历过程中尝试选择其中的一些要点。

## What Is Go?



So just in case you haven't come across Go before, from an article titled [The 10 Most Popular Programming Languages to Learn in 2021](https://www.northeastern.edu/graduate/blog/most-popular-programming-languages/), they say:

因此，以防万一您之前没有接触过 Go，请参阅一篇题为 [2021 年要学习的 10 种最流行的编程语言](https://www.northeastern.edu/graduate/blog/most-popular-programming-语言/)，他们说：

> Also referred to as Golang, Go was developed by Google to be an efficient, readable, and secure language for system-level programming. It works well for distributed systems… While it is a relatively new language, Go has a large standards library and extensive documentation

> Go 也称为 Golang，由 Google 开发，是一种高效、可读且安全的系统级编程语言。它适用于分布式系统……虽然它是一种相对较新的语言，但 Go 拥有庞大的标准库和大量文档

You can find out more about [Go on the official website](https://golang.org/) and we’ll talk more about it during this post.

你可以找到更多关于 [Go on the official website](https://golang.org/)，我们将在这篇文章中详细讨论。

## Domain Applicability

## 域适用性

One of the first things we need to think about when choosing a new language is domain applicability. If you've ever attended any DevNet Express training, this will be one of the big reasons they give as to why Python is so popular with network engineers – because there are so many tools and existing code that can be used to perform your day job that it doesn't make sense to use anything else.

在选择一种新语言时，我们首先需要考虑的事情之一是领域适用性。如果您曾经参加过任何 DevNet Express 培训，这将是他们给出的 Python 为何如此受网络工程师欢迎的重要原因之一——因为有如此多的工具和现有代码可用于执行您的日常工作使用其他任何东西都没有意义。

Well, I guess I’m here to tell you that Go is the new kid on the block and that, whilst already popular in other areas, it is gradually getting a lot more focus from the network community.

好吧，我想我是来告诉你 Go 是这个街区的新孩子，虽然它已经在其他领域流行，但它正逐渐得到网络社区的更多关注。

For example, one popular Python automation framework is [Nornir](https://github.com/nornir-automation/nornir). The same people that created Nornir have created [Gornir](https://github.com/nornir-automation/gornir), _a pluggable framework with inventory management to help operate collections of devices. It’s similar to nornir but in golang_.

例如，一种流行的 Python 自动化框架是 [Nornir](https://github.com/nornir-automation/nornir)。创建 Nornir 的人已经创建了 [Gornir](https://github.com/nornir-automation/gornir)，这是一个具有库存管理功能的可插拔框架，可帮助操作设备集合。它类似于nornir，但在golang_中。

Another example is the [scrapli](https://github.com/carlmontanari/scrapli) library, _a python library focused on connecting to devices, specifically network devices (routers/switches/firewalls/etc.) via Telnet or SSH_. This again appears to be in the early stages of being replicated in Go as [scrapligo](https://github.com/scrapli/scrapligo), _a Go library focused on connecting to devices, specifically network devices (routers/switches/firewalls /etc.) via SSH and NETCONF_. 

另一个例子是 [scrapli](https://github.com/carlmontanari/scrapli) 库，_a python 库专注于通过 Telnet 或 SSH_ 连接到设备，特别是网络设备（路由器/交换机/防火墙/等)。这似乎再次处于在 Go 中作为 [scrapligo](https://github.com/scrapli/scrapligo) 复制的早期阶段，_a Go 库专注于连接到设备，特别是网络设备（路由器/交换机/防火墙/etc.) 通过 SSH 和 NETCONF_。

In addition, when we start talking about Infrastructure as Code, [Ansible](https://www.ansible.com/) is a hugely popular framework. It allows you to modify configuration over time. The problem with this though is that after a while it can be difficult to know what “state” your environment is in, known as “configuration drift”. Anyone who has built a linux server and added packages, updgrades and patches to it over time will know what I mean – could you build an identical server? The term for this is apparently a [Snowflake Server](https://martinfowler.com/bliki/SnowflakeServer.html).

此外，当我们开始谈论基础设施即代码时，[Ansible](https://www.ansible.com/) 是一个非常流行的框架。它允许您随时间修改配置。但问题在于，一段时间后可能很难知道您的环境处于什么“状态”，称为“配置漂移”。任何构建过 linux 服务器并随着时间的推移添加软件包、升级和补丁的人都会明白我的意思——你能构建一个相同的服务器吗？这个术语显然是 [Snowflake Server](https://martinfowler.com/bliki/SnowflakeServer.html)。

The same thing is happening with network configuration and automation. We can add “continuous delivery” to the mix to try and resolve some of the issues, but for me, this procedural approach doesn’t lend itself to large infrastructure deployments and hybrid environments.

网络配置和自动化也发生了同样的事情。我们可以将“持续交付”添加到组合中以尝试解决一些问题，但对我而言，这种程序方法不适用于大型基础设施部署和混合环境。

You may well have found yourself in the same position, getting to the point where Ansible is a useful tool, but a complimentary component is required in order to maintain “desired state configuration”. This is a term I've borrowed from Microsoft and is a Powershell component for managing servers using _declaritive scripting_ and I think it's a term which accurately describes what we mean – we say what we want our environment to look like, and that's what gets created . You never make changes to individual components, you just redeploy them from the new state.

您可能已经发现自己处于相同的位置，到了 Ansible 是一个有用的工具的地步，但需要一个免费的组件来维护“所需的状态配置”。这是我从 Microsoft 借来的一个术语，是一个 Powershell 组件，用于使用 _declaritive scripting_ 管理服务器，我认为这个术语准确地描述了我们的意思——我们说我们希望我们的环境看起来像什么，这就是被创建的.您永远不会对单个组件进行更改，您只需从新状态重新部署它们。

For this, we have [Terraform](https://www.terraform.io/) by Hashicorp. I suspect this isn't the first time you've heard of it, but essentially it enables us to describe our environment in declaritive configuration files, allows us to check what changes are going to be deployed and then applies them when we're ready. No configuration drift here!

为此，我们有 Hashicorp 的 [Terraform](https://www.terraform.io/)。我怀疑这不是您第一次听说它，但本质上它使我们能够在声明性配置文件中描述我们的环境，允许我们检查将要部署的更改，然后在我们准备好时应用它们.这里没有配置漂移！

Terraform has the concept of providers which are essentially connectors to the various components of your infrastructure, and with [Cisco teaming up with Hashicorp](https://blogs.cisco.com/cloud/cisco-and-hashicorp-join-forces-to-deliver-infrastructure-as-code-automation-across-hybrid-cloud) earlier this year, we're only going to see [more and more providers](https://registry.terraform.io/search/providers?namespace=CiscoDevNet) being created for Cisco technology, and Terraform becoming an increasingly important part of your networking tool bag.

Terraform 具有提供商的概念，它们本质上是连接到基础设施各个组件的连接器，并且 [Cisco 与 Hashicorp 合作](https://blogs.cisco.com/cloud/cisco-and-hashicorp-join-forces-to-deliver-infrastructure-as-code-automation-across-hybrid-cloud)今年早些时候，我们只会看到[越来越多的提供者](https://registry.terraform.io/search/providers?namespace=CiscoDevNet) 是为 Cisco 技术创建的，Terraform 成为您网络工具包中越来越重要的一部分。

So, where am I going with this, well, guess what Terraform is written in? And guess what the providers are written in – you guessed it… Go! If you want to keep up with the direction infrastructure as code is going, it might be prudent to at least take a look at Go. That way, you’ll be able to contribute and maybe create providers of your own.

那么，我要去哪里，好吧，猜猜 Terraform 是用什么编写的？猜猜提供者是用什么写的——你猜对了……去吧！如果您想在代码运行时跟上基础设施的发展方向，那么至少看一看 Go 可能是明智之举。这样，您就可以做出贡献，甚至可以创建自己的提供者。

In addition to Terraform, there are many other “cloud native” components written in Go:

除了 Terraform 之外，还有许多其他用 Go 编写的“云原生”组件：

> _We have Docker to build containers, and Kubernetes to orchestrate them. Prometheus lets us monitor them. Consul lets us discover them. Jaeger lets us trace the relationships between them. These are just a few examples, but there are many, many more, all representative of a new generation of technologies: all of them are “cloud native,” and all of them are written in Go._

> _我们用 Docker 来构建容器，用 Kubernetes 来编排它们。 Prometheus 让我们监控它们。 Consul 让我们发现它们。 Jaeger 让我们追踪它们之间的关系。这些只是几个例子，但还有很多很多，都代表了新一代技术：它们都是“云原生”，并且都是用 Go 编写的。_

_![DevNet Gopher](http://darrenparkinson.uk/devnet_gopher.png)

#### Illustration created for Cisco Blogs, made from the original Go Gopher, created by Renee French._

#### 为 Cisco 博客创建的插图，由 Renee French 创建的原始 Go Gopher 制作而成。_

So the Go networking community is definitely growing and I think you can see that reflected in the Go code available on [DevNet Code Exchange](https://developer.cisco.com/codeexchange/explore/#lang=Go).

因此，Go 网络社区肯定在增长，我认为您可以在 [DevNet 代码交换](https://developer.cisco.com/codeexchange/explore/#lang=Go) 上提供的 Go 代码中看到这一点。

Now we’ll take a look at a couple of areas relating to the language itself which I think are relevant to the networking community.

现在我们将看看与语言本身相关的几个领域，我认为这些领域与网络社区相关。

## Simplicity

## 简单

Go is a compiled language like C or Java rather than a dynamic/interpreted language like Python or JavaScript, which typically puts people off because they think it will be difficult to learn and slow to work with.

Go 是一种像 C 或 Java 这样的编译语言，而不是像 Python 或 JavaScript 这样的动态/解释型语言，这通常会让人们望而却步，因为他们认为它很难学习并且使用起来很慢。

In terms of being difficult to learn, Go encourages _simplicity and productivity over clutter and complexity_. From _[“Cloud Native Go”](https://www.oreilly.com/library/view/cloud-native-go/9781492076322/):_ 

在难学方面，Go 鼓励_简单性和生产力，而不是混乱和复杂性_。来自_[“Cloud Native Go”](https://www.oreilly.com/library/view/cloud-native-go/9781492076322/):_

> Go was designed with large projects with lots of contributors in mind. Its minimalist design (just 25 keywords and 1 loop type), and the strong opinions of its compiler, strongly favor clarity over cleverness. This in turn encourages simplicity and productivity over clutter and complexity. The resulting code is relatively easy to ingest, review, and maintain, and harbors far fewer “gotchas.”

> Go 的设计考虑了大量贡献者的大型项目。它的极简设计（只有 25 个关键字和 1 个循环类型）以及其编译器的强烈意见，强烈支持清晰而不是聪明。这反过来又鼓励简单性和生产力，而不是混乱和复杂性。生成的代码相对容易摄取、审查和维护，并且包含的“陷阱”要少得多。

Seriously, _**1 loop type**_ – life changing 😉

说真的，_**1 循环类型**_ – 改变生活 😉

In addition, Go is a [_“Garbage Collected”_](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)) language in the same way that Java and C# are. Some would say this is a disadvantage from a performance perspective (which we'll come to), but from a simplicity point of view, this definitely helps because it means that you don't need to worry directly about memory management like with languages such as C and Rust. This makes it much easier to transition from, or learn alongside, Python.

此外，Go 是一种 [_“垃圾收集”_](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)) 语言，就像 Java 和 C# 一样。有些人会说从性能的角度来看这是一个缺点（我们将谈到)，但从简单的角度来看，这绝对有帮助，因为这意味着你不需要像使用这样的语言那样直接担心内存管理就像 C 和 Rust。这使得从 Python 过渡或与 Python 一起学习变得更加容易。

_![Garbage Collection Gopher](http://darrenparkinson.uk/garbage_collection.png)

#### Illustration created for “A Journey With Go”, made from the original Go Gopher, created by Renee French._

#### 为“A Journey With Go”创作的插图，由 Renee French 创建的原始 Go Gopher 制作而成。_

## Performance

##  表现

I mentioned before that people may be put off because they think compiled languages may be slow to work with. This is usually because of the compilation step. In this section, I’m going to mention a few performance advantages that Go has over other compiled languages and also over interpreted languages.

我之前提到过，人们可能会因为认为编译型语言使用起来很慢而被推迟。这通常是因为编译步骤。在本节中，我将提及 Go 相对于其他编译语言和解释语言的一些性能优势。

### Compilation

### 编译

![Compilation Cartoon](https://imgs.xkcd.com/comics/compiling.png)

#### https://xkcd.com/303/

A question that might be important for those coming from Python is the time it takes to compile their code in a compiled language. This is often an argument given in favour of dynamic languages. However, the fast compilation that you get with Go makes it feel like a dynamic language but with all the benefits of a compiled language – it’s no coincidence that dynamic languages are adding types.

对于那些来自 Python 的人来说可能很重要的一个问题是用编译语言编译他们的代码所需的时间。这通常是支持动态语言的论据。然而，使用 Go 获得的快速编译让它感觉像是一门动态语言，但具有编译语言的所有优点——动态语言添加类型并非巧合。

The story goes that Google engineers designed Go whilst waiting for their other programs to compile, and compilation time was, and still is, a major design consideration. From [the Go FAQ](https://golang.org/doc/faq#creating_a_new_language) relating to why they created another language:

据说谷歌工程师在等待其他程序编译的同时设计了 Go，编译时间过去和现在仍然是一个主要的设计考虑因素。来自 [the Go FAQ](https://golang.org/doc/faq#creating_a_new_language) 关于他们为什么创建另一种语言：

> One had to choose either efficient compilation, efficient execution, or ease of programming; all three were not available in the same mainstream language. Programmers who could were choosing ease over safety and efficiency by moving to dynamically typed languages such as Python and JavaScript rather than C++ or, to a lesser extent, Java.

> 必须选择高效编译、高效执行或易于编程；所有这三种语言都没有以相同的主流语言提供。通过转向动态类型语言（如 Python 和 JavaScript）而不是 C++ 或在较小程度上使用 Java，可以选择轻松而不是安全和效率的程序员。

Where they then go on to say (emphasis mine):

然后他们继续说（强调我的）：

> Go addressed these issues by attempting to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aimed to be modern, with support for networked and multicore computing. Finally, **working with Go is intended to be fast: it should take at most a few seconds to build a large executable on a single computer**.

> Go 通过尝试将解释型、动态类型语言的编程简便性与静态类型、编译型语言的效率和安全性结合起来，解决了这些问题。它还旨在成为现代的，支持网络和多核计算。最后，**使用 Go 的目的是快速：在单台计算机上构建大型可执行文件最多只需要几秒钟**。

By way of an example (from Cloud Native Go):

举个例子（来自 Cloud Native Go）：

> building all 1.8 million lines of Go in Kubernetes v1.20.2 on a MacBook Pro with a 2.4 GHz 8-Core Intel i9 processor and 32 GB of RAM required about 45 seconds of real time

> 在配备 2.4 GHz 8 核 Intel i9 处理器和 32 GB RAM 的 MacBook Pro 上使用 Kubernetes v1.20.2 构建所有 180 万行 Go 需要大约 45 秒的实时时间

So compiling and running your average “script” (or even larger ones) shouldn’t be a problem!

所以编译和运行你的普通“脚本”（甚至更大的）应该不是问题！

### Code Execution

### 代码执行

For me, performance isn’t just about how fast the code runs – if that’s your main requirement, then other languages might be a a better fit (looking at you Rust). I suspect though, that unless you're writing code for embedded systems, or creating your own operating system, it won't be that big of a deal, since Go compares favourably to other compiled languages, even those with manual memory management, and is far easier to learn too.

对我来说，性能不仅仅是代码运行的速度——如果这是你的主要要求，那么其他语言可能更适合（看看你的 Rust）。不过我怀疑，除非您为嵌入式系统编写代码，或者创建自己的操作系统，否则这不会有什么大不了的，因为 Go 与其他编译语言相比，甚至是那些手动内存管理的语言都要好，而且也更容易学习。

Being a compiled language, Go will obviously be faster than any interpreted language. By way of an example, benchmarks show Python to be 10 to 100 times slower than compiled languages. Check out [the benchmarks game](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-python3.html) for a comparison.

作为一种编译语言，Go 显然比任何解释性语言都要快。例如，基准测试显示 Python 比编译语言慢 10 到 100 倍。查看 [基准测试游戏](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-python3.html) 进行比较。

### Containerisation 

### 容器化

Finally it might seem odd to add containerisation into performance, but if you’re building containers, either locally or part of your CI/CD pipeline and pushing/pulling containers around everywhere, the size of those containers is going to be pretty important.

最后，将容器化添加到性能中似乎很奇怪，但是如果您在本地或 CI/CD 管道的一部分构建容器并在各处推/拉容器，那么这些容器的大小将非常重要。

By way of a simple test I created some standard containers for a simple Hello World application and compared their sizes. Firstly I created the images using the standard containers `golang:latest` and `python:3.7` for Go and Python respectively. As you can see, they are fairly well matched.

通过一个简单的测试，我为一个简单的 Hello World 应用程序创建了一些标准容器并比较了它们的大小。首先，我分别使用 Go 和 Python 的标准容器 `golang:latest` 和 `python:3.7` 创建了图像。如您所见，它们非常匹配。

However, if we apply a little more thought and use more appropriate containers for production, we can make these images much smaller. In this case, I used `scratch` and `python:3.7-alpine` for Go and Python respectively.

但是，如果我们多加思考并使用更合适的容器进行生产，我们可以使这些图像更小。在这种情况下，我分别为 Go 和 Python 使用了 `scratch` 和 `python:3.7-alpine`。

RepositorySizesmall-python41.9MBsmall-go1.2MB

As you can see there is still a reasonable difference in the container size and this is due to the fact that Go binaries are compiled with all their requirements and don’t require a runtime environment to execute. Not only does this improve the performance and portability of your application, but will also no doubt have an impact on the security of it too.

如您所见，容器大小仍然存在合理差异，这是因为 Go 二进制文件是根据其所有要求编译的，不需要运行时环境来执行。这不仅会提高应用程序的性能和可移植性，而且无疑也会对其安全性产生影响。

You could argue that the python image could be made smaller, but this would most likely be at the expense of simplicity and readability, and that’s never a good thing.

你可能会争辩说，python 图像可以做得更小，但这很可能会以牺牲简单性和可读性为代价，这从来都不是一件好事。

Clearly there are a lot of other areas around performance we could take on, but I’m going to move onto the final piece which I just mentioned, and that’s portability.

显然，我们可以在性能方面进行很多其他方面的工作，但我将继续讨论我刚刚提到的最后一个部分，那就是可移植性。

## Portability

## 便携性

The final thing I wanted to talk about briefly is portability, mainly because I have found this to be really useful.

我想简要讨论的最后一件事是可移植性，主要是因为我发现这非常有用。

Essentially Go provides the ability to easily share a whole application with any user without requiring them to have any particular environment set up.

从本质上讲，Go 提供了与任何用户轻松共享整个应用程序的能力，而无需他们设置任何特定的环境。

This is because when you compile your Go code it produces a _statically linked executable binary_. This just means that it wraps in any dependencies and the runtime into a single executable file. So when you share it with someone, they don’t need the Go compiler installed on their machine, or any of the libraries that you used to create the application.

这是因为当您编译 Go 代码时，它会生成一个 _statically 链接的可执行二进制文件_。这只是意味着它将任何依赖项和运行时包装到一个可执行文件中。因此，当您与某人共享它时，他们不需要在他们的机器上安装 Go 编译器，也不需要您用于创建应用程序的任何库。

Contrast this with a dynamic language like Python or Javascript where you need to ensure that the recipient has the Python or Javascript interpreter on their machine. Most often this will involve ensuring they have the correct version installed too. And finally, on receipt of your application, they will have to install any required libraries that you used to create it.

将此与 Python 或 Javascript 等动态语言进行对比，您需要确保收件人在其机器上具有 Python 或 Javascript 解释器。大多数情况下，这将涉及确保他们也安装了正确的版本。最后，在收到您的应用程序后，他们将必须安装您用于创建它的任何必需的库。

The ability to create easy to use applications that are easily shared with colleagues without requiring anything of them, from experience, is a breath of fresh air!

能够创建易于使用的应用程序，这些应用程序可以轻松地与同事共享，而无需他们的任何经验，从经验来看，这是一股清新的空气！

I hope this post has been useful. If you have any comments or feedback, please feel free to reach out to me on [twitter](https://twitter.com/darrenparkinson).

我希望这篇文章很有用。如果您有任何意见或反馈，请随时通过 [twitter](https://twitter.com/darrenparkinson) 与我联系。

## Resources

##  资源

- [DevNet Code Exchange](https://developer.cisco.com/codeexchange/explore/#lang=Go) where there are libraries for many Cisco platforms, including:

   - [DevNet 代码交换](https://developer.cisco.com/codeexchange/explore/#lang=Go)，其中有许多 Cisco 平台的库，包括：

  - [ACI client for Go](https://github.com/ciscoecosystem/aci-go-client);
   - [Terraform provider for ACI](https://github.com/CiscoDevNet/terraform-provider-aci);that uses the aforementioned ACI client for Go;
   - [SDWAN client for Go](https://github.com/CiscoDevNet/sdwan-go-client);
   - [Terraform provider for SDWAN](https://github.com/CiscoDevNet/terraform-provider-sdwan)
- [Meraki CLI Utility](https://github.com/ddexterpark/merakictl) using the [Go Dashboard API](https://github.com/ddexterpark/dashboard-api-golang), both by Dexter Park;
- [Gornir](https://github.com/nornir-automation/gornir) \- Go implementaton of nornir by the same people!
- [Gomiko](https://github.com/Ali-aqrabawi/gomiko) \- Go implementation of netmiko by Ali-aqrabawi;
- [scrapligo](https://github.com/scrapli/scrapligo) \- Go library focused on connecting to devices, specifically network devices (routers/switches/firewalls/etc.) via SSH and NETCONF;
- [Protobuf Files](https://github.com/cisco-ie/cisco-proto) for Cisco networking operating systems;
- [Webex Library](https://github.com/jbogarin/go-cisco-webex-teams) by Jose Bogarín

- [ACI 客户端 Go](https://github.com/ciscoecosystem/aci-go-client);
  - [ACI 的 Terraform 提供程序](https://github.com/CiscoDevNet/terraform-provider-aci)；使用前面提到的 ACI 客户端进行 Go；
  - [Go 的 SDWAN 客户端](https://github.com/CiscoDevNet/sdwan-go-client)；
  - [SDWAN 的 Terraform 提供商](https://github.com/CiscoDevNet/terraform-provider-sdwan)
- [Meraki CLI 实用程序](https://github.com/ddexterpark/merakictl) 使用 [Go Dashboard API](https://github.com/ddexterpark/dashboard-api-golang)，均由 Dexter Park 提供；
- [Gornir](https://github.com/nornir-automation/gornir) \- 由同一个人去实现nornir！
- [Gomiko](https://github.com/Ali-aqrabawi/gomiko) \- 由 Ali-aqrabawi Go 实现 netmiko；
- [scrapligo](https://github.com/scrapli/scrapligo) \- Go 库专注于通过 SSH 和 NETCONF 连接到设备，特别是网络设备（路由器/交换机/防火墙/等)；
- [Protobuf 文件](https://github.com/cisco-ie/cisco-proto) 用于思科网络操作系统；
- [Webex 库](https://github.com/jbogarin/go-cisco-webex-teams) 作者 Jose Bogarín

[cisco](https://darrenparkinson.uk//tags/cisco/)[networking](https://darrenparkinson.uk//tags/networking/) [golang](https://darrenparkinson.uk//tags/golang/) 

[cisco](https://darrenparkinson.uk//tags/cisco/)[网络](https://darrenparkinson.uk//tags/networking/) [golang](https://darrenparkinson.uk//tags/golang/)

