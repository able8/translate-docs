# Rust vs. Go: Why They’re Better Together

# Rust vs. Go：为什么他们在一起更好

12 Mar 2021

2021 年 3 月 12 日

Steve Francia
Over the past 25 years Steve Francia has built some of the most innovative and successful technologies and companies which have become the foundation of cloud computing, embraced by enterprises and developers all over the world. He is currently product and strategy lead for the Go Programming Language at Google. He is the creator of Hugo, Cobra, Viper, spf13-vim and many additional open source projects, having the unique distinction of leading five of the world's largest open source projects.](https://github.com/spf13)

史蒂夫·弗兰西亚
在过去的 25 年中，Steve Francia 建立了一些最具创新性和最成功的技术和公司，这些技术和公司已成为云计算的基础，受到全球企业和开发人员的欢迎。他目前是谷歌 Go 编程语言的产品和战略负责人。他是 Hugo、Cobra、Viper、spf13-vim 和许多其他开源项目的创建者，具有领导世界上最大的五个开源项目的独特优势。](https://github.com/spf13)

While others may see [Rust](https://www.rust-lang.org/) and [Go](https://go.dev/) as competitive programming languages, neither the Rust nor the Go teams do. Quite the contrary, our teams have deep respect for what the others are doing, and see the languages as complimentary with a shared vision of modernizing the state of software development industry-wide.

虽然其他人可能将 [Rust](https://www.rust-lang.org/) 和 [Go](https://go.dev/) 视为具有竞争力的编程语言，但 Rust 和 Go 团队都没有。恰恰相反，我们的团队非常尊重其他人的工作，并将这些语言视为对整个行业软件开发状态现代化的共同愿景的补充。

In this article, we will discuss the pros and cons of Rust and Go and how they supplement and support each other, and our recommendations for when each language is most appropriate.

在本文中，我们将讨论 Rust 和 Go 的优缺点以及它们如何相互补充和支持，以及我们对每种语言何时最合适的建议。

Companies are finding value in adopting both languages and in their complimentary value. To shift from our opinions to hands-on user experience, we spoke with three such companies, [Dropbox](https://www.dropbox.com/), [Fastly](https://www.fastly.com/) , and [Cloudflare](https://www.cloudflare.com/), about their experience in using Go and Rust together. There will be quotes from them throughout this article to give further perspective.

公司正在发现采用这两种语言及其互补价值的价值。为了从我们的意见转变为实际的用户体验，我们与三个这样的公司进行了交谈，[Dropbox](https://www.dropbox.com/)，[Fastly](https://www.fastly.com/) , 和 [Cloudflare](https://www.cloudflare.com/)，关于他们一起使用 Go 和 Rust 的经验。整篇文章中都会引用他们的引述，以提供进一步的观点。

## Language Comparison

## 语言比较

LanguageGoRustCreation Date20092010Created atGoogleMozillaNotable software written in languageKubernetes, Docker, Github CLI, Hugo, Caddy, Drone, Ethereum, Syncthing, TerraformFirefox, ripgrep, alacritty, deno, HabitatKey workloadsAPIs, Web Apps, CLI apps, DevOps, Networking, Data Processing, cloud appsIoT, processing engines, security-sensitive apps, system components, cloud apps[Developer adoption](https://insights.stackoverflow.com/survey/2020%23technology-programming-scripting-and-markup-languages-all-respondents)8.8% (#12)5.1% (#19)[Most loved](https://insights.stackoverflow.com/survey/2020%23technology-programming-scripting-and-markup-languages-all-respondents)62.3% (#5 )86.1% (#1)[Most wanted](https://insights.stackoverflow.com/survey/2020%23technology-most-loved-dreaded-and-wanted-languages-wanted)17.9% (#3)14.6% (#5)

LanguageGoRustCreation Date20092010创建于GoogleMozilla 用语言编写的著名软件Kubernetes、Docker、Github CLI、Hugo、Caddy、Drone、Ethereum、Syncthing、TerraformFirefox、ripgrep、alacritty、deno、HabitatKey 工作负载API、Web 应用程序、CLI 应用程序、DevOps、云处理、数据物联网应用程序处理引擎、安全敏感应用、系统组件、云应用[开发者采用](https://insights.stackoverflow.com/survey/2020%23technology-programming-scripting-and-markup-languages-all-respondents)8.8% (#12)5.1% (#19)[最受欢迎](https://insights.stackoverflow.com/survey/2020%23technology-programming-scripting-and-markup-languages-all-respondents)62.3% (#5 )86.1% (#1)[最想要的](https://insights.stackoverflow.com/survey/2020%23technology-most-loved-dreaded-and-wanted-languages-wanted)17.9% (#3)14.6% (#5)

## Similarities

## 相似之处

Go and Rust have a lot in common. Both are modern software languages born out of a need to provide a safe and scalable solution to the problems impacting software development. Both were created as reactions to shortcomings the creators were experiencing with existing languages in the industry, particularly shortcomings of developer productivity, scalability, safety and concurrency.

Go 和 Rust 有很多共同点。两者都是现代软件语言，都是出于为影响软件开发的问题提供安全且可扩展的解决方案的需要。两者都是为了应对创建者在行业中使用现有语言所遇到的缺点，特别是开发人员生产力、可扩展性、安全性和并发性的缺点。

Most of today’s popular languages were designed over 30 years ago. When those languages were designed there were five key differences from today:

今天大多数流行的语言都是在 30 多年前设计的。在设计这些语言时，与今天相比有五个主要区别：

1. Moore’s law was thought to be eternally true.
2. Most software projects were written by small teams, often working in person together.
3. Most software had a relatively small number of dependencies, mostly proprietary.
4. Safety was a secondary concern… or not a concern at all.
5. Software was typically written for a single platform.

1. 摩尔定律被认为永远正确。
2. 大多数软件项目是由小团队编写的，通常是一起工作。
3. 大多数软件的依赖项相对较少，主要是专有的。
4. 安全是次要的问题……或者根本不是问题。
5. 软件通常是为单一平台编写的。

In contrast, both Rust and Go were written for today’s world and generally took similar approaches to design a language for today’s development needs.

相比之下，Rust 和 Go 都是为当今世界编写的，并且通常采用类似的方法来设计一种满足当今开发需求的语言。

### 1\. Performance and Concurrency

### 1\.性能和并发

Go and Rust are both compiled languages focused on producing efficient code. They also provide easy access to the multiple processors of today’s machines, making them ideal languages for writing efficient parallel code.

Go 和 Rust 都是专注于生成高效代码的编译语言。它们还提供对当今机器的多个处理器的轻松访问，使它们成为编写高效并行代码的理想语言。

_“Using Go allowed MercadoLibre to cut the number of servers they use for this service to one-eighth the original number (from 32 servers down to four), plus each server can operate with less power (originally four CPU cores, now down to two CPU cores). With Go, the company obviated 88 percent of their servers and cut CPU on the remaining ones in half—producing a tremendous cost-savings.”_— “ [MercadoLibre Grows with Go](https://go.dev/solutions/mercadolibre /)”

_“使用 Go 允许 MercadoLibre 将他们用于此服务的服务器数量减少到原始数量的八分之一（从 32 个服务器减少到 4 个），而且每台服务器可以以更少的功率运行（最初是四个 CPU 内核，现在减少到两个 CPU 内核）。使用 Go，该公司取消了 88% 的服务器，并将其余服务器的 CPU 削减了一半——从而节省了大量成本。”_—“[MercadoLibre 随 Go 一起成长](https://go.dev/solutions/mercadolibre /)”

_“In our tightly managed environments where we run Go code, we have seen a CPU reduction of approximately ten percent [vs C++] with cleaner and maintainable code.”_ — [Bala Natarajan, Paypal](https://go.dev /solutions/paypal/) 
_“在我们运行 Go 代码的严格管理的环境中，我们看到 CPU 减少了大约 10% [与 C++] 相比，代码更干净且可维护。”_ — [Bala Natarajan, Paypal](https://go.dev /解决方案/贝宝/）
_“Here at AWS, we love Rust, too, because it helps AWS write highly performant, safe infrastructure-level networking and other systems software. Amazon’s first notable product built with Rust, Firecracker, launched publicly in 2018 and provides the open source virtualization technology that powers AWS Lambda and other serverless offerings. But we also use Rust to deliver services such as Amazon Simple Storage Service (Amazon S3), Amazon Elastic Compute Cloud (Amazon EC2), Amazon CloudFront, Amazon Route 53, and more. Recently we launched Bottlerocket, a Linux-based container operating system written in Rust.” — [Matt Asay, Amazon Web Services](https://aws.amazon.com/blogs/opensource/why-aws-loves-rust-and-how-wed-like-to-help/)_

_“在 AWS，我们也喜欢 Rust，因为它可以帮助 AWS 编写高性能、安全的基础设施级网络和其他系统软件。亚马逊第一个使用 Rust 构建的著名产品 Firecracker 于 2018 年公开发布，提供支持 AWS Lambda 和其他无服务器产品的开源虚拟化技术。但我们也使用 Rust 提供服务，例如 Amazon Simple Storage Service (Amazon S3)、Amazon Elastic Compute Cloud (Amazon EC2)、Amazon CloudFront、Amazon Route 53 等。最近，我们推出了 Bottlerocket，这是一个用 Rust 编写的基于 Linux 的容器操作系统。” — [Matt Asay，亚马逊网络服务](https://aws.amazon.com/blogs/opensource/why-aws-loves-rust-and-how-wed-like-to-help/)_

_We “saw an extraordinary 1200-1500% increase in our speed! We went from 300-450ms in release mode with Scala with fewer parsing rules implemented, to 25-30ms in Rust with more parsing rules implemented!” — [Josh Hannaford, IBM](https://developer.ibm.com/technologies/web-development/articles/why-webassembly-and-rust-together-improve-nodejs-performance/)_

_我们“看到我们的速度惊人地提高了 1200-1500%！我们从 Scala 的发布模式下的 300-450 毫秒实现了更少的解析规则，到 Rust 中的 25-30 毫秒，实现了更多的解析规则！” — [Josh Hannaford，IBM](https://developer.ibm.com/technologies/web-development/articles/why-webassembly-and-rust-together-improve-nodejs-performance/)_

### 2\. Team Scalable — Reviewable

### 2\。团队可扩展 - 可审查

Software development today is built by teams that grow and expand, often collaborating in a distributed way using source control. Go and Rust are both designed for how teams work, improving code reviews by removing unnecessary concerns like formatting, security, and complex organization. Both languages require relatively little context to understand what the code is doing, allowing reviewers to more quickly work with code written by other people and review code by both team members and code contributed by open source developers outside of your team.

今天的软件开发是由成长和扩展的团队构建的，通常使用源代码控制以分布式方式进行协作。 Go 和 Rust 都是为团队的工作方式而设计的，通过消除不必要的问题（如格式、安全性和复杂的组织）来改进代码审查。这两种语言都需要相对较少的上下文来理解代码在做什么，从而允许审查者更快地处理其他人编写的代码，审查团队成员的代码和团队外部的开源开发人员贡献的代码。

_“Building Go and Rust code, having come from a Java and Ruby background in my early career, felt like an impossible weight off my shoulders. When I was at Google, it was a relief to come across a service that was written in Go, because I knew it would be easy to build and run. This has also been true of Rust, though I’ve only worked on that at a much smaller scale. I'm hoping that the days of infinitely configurable build systems are dead, and languages all ship with their own purpose-built build tools that just work out of the box.”— [Sam Rose, CV Partner](https://bitfieldconsulting .com/golang/rust-vs-go)_

_“构建 Go 和 Rust 代码，在我早期的职业生涯中来自 Java 和 Ruby 背景，感觉就像是我肩上不可能的重担。当我在 Google 时，遇到一个用 Go 编写的服务让我松了一口气，因为我知道它很容易构建和运行。 Rust 也是如此，尽管我只在小得多的规模上进行了研究。我希望无限可配置构建系统的时代已经过去，所有语言都带有自己的专用构建工具，这些工具开箱即用。”— [Sam Rose，CV 合作伙伴](https://bitfieldconsulting .com/golang/rust-vs-go)_

_“I tend to breathe a sigh of relief when writing a service in Go since it has a very simple, easy to reason about, static type system compared to dynamic languages, concurrency is a first-class citizen, and Go's standard library is both unbelievably polished and powerful, yet also to the point. Take a standard Go install, throw in a grpc library and a database connector, and you need very little else to build anything on the server-side, and every engineer will be able to read the code and understand the libraries. When writing a module in Rust, Dropbox engineers felt Rust’s growing pains on the server-side before Async-await stabilized in 2019, but since then, crates are converging to use it and we get the benefit of async patterns coupled with fearless concurrency.” — Daniel Reiter Horn, Dropbox_

_“在用 Go 编写服务时，我倾向于松一口气，因为与动态语言相比，它有一个非常简单、易于推理的静态类型系统，并发性是一等公民，而且 Go 的标准库两者兼而有之令人难以置信的抛光和强大，但也恰到好处。进行标准的 Go 安装，加入 grpc 库和数据库连接器，您几乎不需要其他任何东西就可以在服务器端构建任何东西，每个工程师都能够阅读代码并理解库。在用 Rust 编写模块时，Dropbox 工程师在 Async-await 于 2019 年稳定之前就感受到了 Rust 在服务器端的成长之痛，但从那时起，板条箱开始融合使用它，我们获得了异步模式和无畏并发的好处。” — Daniel Reiter Horn，Dropbox_

### 3\. Open Source-aware

### 3\。开源意识

The number of dependencies used by the average software project today is staggering. The decades-long goal of software reuse has been achieved in modern development, where today’s software is built using 100s of projects. To do so, developers use software repositories, which increasingly has become a staple of software development across a broadening range of applications. Each of the packages a developer includes, in turn, has its own dependencies. Languages for today’s programming environments need to handle this complexity effortlessly.

今天普通软件项目使用的依赖项数量是惊人的。数十年来软件重用的目标已经在现代开发中实现，今天的软件是使用数百个项目构建的。为此，开发人员使用软件存储库，它日益成为各种应用程序中软件开发的主要内容。开发人员包含的每个包都有自己的依赖项。当今编程环境的语言需要毫不费力地处理这种复杂性。

Both Go and Rust have package-management systems that allow developers to make a simple list of the packages they'd like to build on, and the language tools automatically fetch and maintain those packages for them, so that developers can focus more on their own code and less on the management of others.

Go 和 Rust 都有包管理系统，允许开发人员制作他们想要构建的包的简单列表，语言工具会自动为他们获取和维护这些包，以便开发人员可以更专注于自己的工作代码，少管别人。

### 4\. Safety

### 4\。安全

The security concerns of today's applications are well-addressed by both Go and Rust, which ensure that code built in the languages run without exposing the user to a variety of classic security vulnerabilities like buffer overflows, use-after-free, etc. By removing these concerns, developers can focus on the problems at hand and build applications that are more secure by default. 
Go 和 Rust 都很好地解决了当今应用程序的安全问题，确保在这些语言中构建的代码运行时不会将用户暴露于各种经典的安全漏洞，如缓冲区溢出、释放后使用等。有了这些顾虑，开发人员就可以专注于手头的问题，并构建默认情况下更安全的应用程序。
_“The [Rust] compiler really holds your hand when working through the errors that you do get. This lets you focus on your business objectives rather than bug hunting or deciphering cryptic messages.” — [Josh Hannaford, IBM](https://developer.ibm.com/technologies/web-development/articles/why-webassembly-and-rust-together-improve-nodejs-performance/)_

_“在处理您遇到的错误时，[Rust] 编译器真的会帮助您。这让您可以专注于您的业务目标，而不是寻找错误或破译神秘的消息。” — [Josh Hannaford，IBM](https://developer.ibm.com/technologies/web-development/articles/why-webassembly-and-rust-together-improve-nodejs-performance/)_

_“In short, the flexibility, safety, and security of Rust outweighs any inconvenience of having to follow strict lifetime, borrowing, and other compiler rules or even the lack of a garbage collector. These features are a much-needed addition to cloud software projects and will help avoid many bugs commonly found in them.” — [Taylor Thomas, Sr., Microsoft](https://msrc-blog.microsoft.com/2020/04/29/the-safety-boat-kubernetes-and-rust/)._

_“简而言之，Rust 的灵活性、安全性和安全性超过了必须遵循严格的生命周期、借用和其他编译器规则甚至缺少垃圾收集器所带来的任何不便。这些功能是云软件项目急需的补充，将有助于避免其中常见的许多错误。” — [Taylor Thomas, Sr., Microsoft](https://msrc-blog.microsoft.com/2020/04/29/the-safety-boat-kubernetes-and-rust/)._

_“Go is strongly and statically typed with no implicit conversions, but the syntactic overhead is still surprisingly small. This is achieved by simple type inference in assign­ments together with untyped numeric constants. This gives Go stronger type safety than Java (which has implicit conversions), but the code reads more like Python (which has untyped variables).” — [Stefan Nilsson, computer science professor](https://yourbasic.org/golang/advantages-over-java-python/)._

_“Go 是强静态类型的，没有隐式转换，但语法开销仍然非常小。这是通过赋值中的简单类型推断以及无类型数字常量来实现的。这为 Go 提供了比 Java（具有隐式转换）更强的类型安全性，但代码读起来更像 Python（具有无类型变量）。” — [Stefan Nilsson，计算机科学教授](https://yourbasic.org/golang/advantages-over-java-python/)._

_“When building our Brotli compression library for storing block data at Dropbox, we limited ourselves to the safe subset of Rust and, further, to the core library (no-stdlib) as well, with the allocator specified as a generic. Using the subset of Rust this way made it very easy to call the Rust-Brotli library from Rust on the client-side and using the C FFI from both Python and Go on the Server. This compilation mode also provided [substantial security guarantees](https://dropbox.tech/infrastructure/lossless-compression-with-brotli). After some tuning, the Rust Brotli implementation, despite being 100% safe, array-bounds-checked code, was still faster than the corresponding native Brotli code in C.” — Daniel Reiter Horn, Dropbox_

_“在构建用于在 Dropbox 存储块数据的 Brotli 压缩库时，我们将自己限制在 Rust 的安全子集，此外，还限制了核心库（no-stdlib），并将分配器指定为泛型。以这种方式使用 Rust 的子集使得在客户端从 Rust 调用 Rust-Brotli 库并在服务器上使用来自 Python 和 Go 的 C FFI 变得非常容易。这种编译模式还提供了【实质性的安全保证】（https://dropbox.tech/infrastructure/lossless-compression-with-brotli）。经过一些调整后，尽管 100% 安全、经过数组边界检查的代码，Rust Brotli 实现仍然比 C 中相应的本机 Brotli 代码快。” — Daniel Reiter Horn，Dropbox_

### 5\. Truly Portable

### 5\。真正便携

It is trivial in both Go and Rust to write one piece of software that runs on many different operating systems and architectures. “Write once, compile anywhere.” In addition, both Go and Rust natively support cross-compilation eliminating the need for “build farms” commonly associated with older compiled languages.

在 Go 和 Rust 中编写一个在许多不同操作系统和架构上运行的软件都是微不足道的。 “一次编写，随处编译。”此外，Go 和 Rust 本身都支持交叉编译，消除了通常与旧编译语言相关的“构建农场”的需要。

_“Golang possesses great qualities for production optimization such as having a small memory footprint, which supports its capability for being building blocks in large-scale projects, as well as easy cross-compilation to other architectures out of the box. Since Go code is compiled into a single static binary, it allows easy containerization and, by extension, makes it almost trivial to deploy Go into any highly available environment such as Kubernetes.” — [Dewet Diener, Curve](https://jaxenter.com/golang-curve-163187.html)._

_“Golang 具有生产优化的优良品质，例如内存占用小，支持其在大型项目中构建块的能力，以及开箱即用的轻松交叉编译到其他架构的能力。由于 Go 代码被编译为单个静态二进制文件，因此可以轻松进行容器化，并且通过扩展，将 Go 部署到任何高度可用的环境（例如 Kubernetes）中几乎变得微不足道。” — [Dewet Diener，曲线](https://jaxenter.com/golang-curve-163187.html)._

_“When you look at a cloud-based infrastructure, often you’re using something like a Docker container to deploy your workloads. With a static binary that you build in Go, you could have a Docker file that's 10, 11, 12 megabytes instead of bringing in the entire Node.js ecosystem, or Python, or Java, where you've got these hundreds of megabyte- sized Docker files. So, shipping that tiny binary is amazing.” — [Brian Ketelsen, Microsoft](https://cloudblogs.microsoft.com/opensource/2018/02/21/go-lang-brian-ketelsen-explains-fast-growth/)._

_“当您查看基于云的基础架构时，您通常会使用 Docker 容器之类的东西来部署您的工作负载。使用您在 Go 中构建的静态二进制文件，您可以拥有 10、11、12 兆字节的 Docker 文件，而不是引入整个 Node.js 生态系统、Python 或 Java，在那里您可以获得数百兆字节-大小的 Docker 文件。所以，运送那个小小的二进制文件真是太棒了。” — [Brian Ketelsen，微软](https://cloudblogs.microsoft.com/opensource/2018/02/21/go-lang-brian-ketelsen-explains-fast-growth/)._

_“With Rust, we’ll have a high-performance and portable platform that we can easily run on Mac, iOS, Linux, Android, and Windows.” — [Matt Ronge, Astropad](https://blog.astropad.com/why-rust/)._

_“有了 Rust，我们将拥有一个高性能和便携的平台，我们可以轻松地在 Mac、iOS、Linux、Android 和 Windows 上运行。” — [Matt Ronge，Astropad](https://blog.astropad.com/why-rust/)._

## Differences

## 差异

In design, there are always trade-offs that must be made. While Go and Rust emerged around the same time with similar goals, as they faced decisions at times they chose different trade-offs that separated the languages in key ways.

在设计中，总是需要权衡取舍。虽然 Go 和 Rust 大约在同一时间出现，有着相似的目标，但由于他们有时面临决策，他们选择了不同的权衡，这些权衡在关键方面将语言分开。

### 1\. Performance

### 1\.表现

Go has excellent performance right out of the box. By design, there are no knobs or levers that you can use to squeeze more performance out of Go. Rust is designed to enable you to squeeze every last drop of performance out of the code; in this regard, you really can’t find a faster language than Rust today. However, Rust’s increased performance comes at the cost of additional complexity. 
Go 具有开箱即用的出色性能。按照设计，没有任何旋钮或控制杆可让您从 Go 中获得更多性能。 Rust 旨在使您能够从代码中挤出每一滴性能；在这方面，你今天真的找不到比 Rust 更快的语言。然而，Rust 的性能提升是以增加复杂性为代价的。
_“Remarkably, we had only put very basic thought into optimization as the Rust version was written. Even with just basic optimization, Rust was able to outperform the hyper-hand-tuned Go version. This is a huge testament to how easy it is to write efficient programs with Rust compared to the deep dive we had to do with Go.” — [Jesse Howarth, Discord](https://blog.discord.com/why-discord-is-switching-from-go-to-rust-a190bbca2b1f)._

_“值得注意的是，在编写 Rust 版本时，我们只对优化进行了非常基本的考虑。即使只是进行基本优化，Rust 也能够胜过超手动调整的 Go 版本。与我们必须使用 Go 进行的深入研究相比，这极大​​地证明了使用 Rust 编写高效的程序是多么容易。” — [Jesse Howarth, Discord](https://blog.discord.com/why-discord-is-switching-from-go-to-rust-a190bbca2b1f)._

_“Dropbox engineers often see 5x performance and latency improvements by porting line-for-line Python code into Go, and memory usage often drops dramatically as compared with Python as there is no GIL and the process count may be reduced. However, when we are memory constrained, as on desktop client software or in certain server processes, we move over to Rust as the manual memory management in Rust is substantially more efficient than the Go GC.” — Daniel Reiter Horn, Dropbox_

_“通过将逐行 Python 代码移植到 Go 中，Dropbox 工程师经常看到 5 倍的性能和延迟改进，并且与 Python 相比，内存使用量通常会急剧下降，因为没有 GIL 并且进程数可能会减少。但是，当我们受到内存限制时，例如在桌面客户端软件或某些服务器进程中，我们会转向 Rust，因为 Rust 中的手动内存管理比 Go GC 效率高得多。” — Daniel Reiter Horn，Dropbox_

### 2\. Adaptability/Interability

### 2\。适应性/交互性

Go’s strength of quick iteration allows developers to try ideas quickly and hone in on working code that solves the task at hand. Often, this is sufficient and frees the developer to move onto other tasks. Rust, on the other hand, has longer compiles compared with Go, leading to slower iteration times. This leads Go to work better in scenarios where faster turnaround time allows developers to adapt to changing requirements, while Rust thrives in scenarios where more time can be given to making a more refined and performant implementation.

Go 的快速迭代优势允许开发人员快速尝试想法并磨练解决手头任务的工作代码。通常，这就足够了，并且可以让开发人员腾出时间处理其他任务。另一方面，与 Go 相比，Rust 的编译时间更长，导致迭代时间更慢。这导致 Go 在更快的周转时间允许开发人员适应不断变化的需求的场景中更好地工作，而 Rust 在可以有更多时间进行更精细和高性能的实现的场景中茁壮成长。

_“The genius of the Go type system is that callers can define the Interfaces, allowing libraries to return expansive structs but require narrow interfaces. The genius of the Rust type system is the combination of match syntax with Result<>, where you can be statically certain every eventuality is handled and never have to invent null values to satisfy unused return parameters.” — Daniel Reiter Horn, Dropbox_

_“Go 类型系统的天才之处在于调用者可以定义接口，允许库返回扩展结构但需要窄接口。 Rust 类型系统的天才之处在于匹配语法与 Result<> 的组合，在这种情况下，您可以静态地确定每个可能性都得到了处理，而不必发明空值来满足未使用的返回参数。” — Daniel Reiter Horn，Dropbox_

_“(I)f your use case is closer to customers, it’s more vulnerable to shifting requirements, then Go is a lot nicer because the cost of continuous refactor is a lot cheaper. It’s how fast you can express the new requirements and try them out.” — Peter Bourgon, Fastly_

_“(I)如果你的用例更接近客户，它更容易受到需求变化的影响，那么 Go 会更好，因为持续重构的成本要便宜得多。这是您表达新要求并尝试它们的速度。” — Peter Bourgon，Fastly_

### 3\. Learnability

### 3\。可学习性

Simply put, there really isn’t a more approachable language than Go. There are many stories of teams who were able to adopt Go and put Go services/applications into production in a few weeks. Additionally, Go is relatively unique among languages in that its language design and practices are quite consistent over it’s 10+ year lifetime. So time invested in learning Go maintains its value for a long time. By comparison, Rust is considered a difficult language to learn due to its complexity. It generally takes several months of learning Rust to feel comfortable with it, but with this extra complexity comes precise control and increased performance.

简而言之，真的没有比 Go 更平易近人的语言了。有许多团队能够在几周内采用 Go 并将 Go 服务/应用程序投入生产的故事。此外，Go 在语言中相对独特，因为它的语言设计和实践在其 10 多年的生命周期中非常一致。所以花在学习围棋上的时间会在很长一段时间内保持其价值。相比之下，由于其复杂性，Rust 被认为是一种难以学习的语言。通常需要几个月的时间学习 Rust 才能适应它，但这种额外的复杂性带来了精确的控制和性能的提高。

_“At the time, no single team member knew Go, but within a month, everyone was writing in Go” – [Jaime Garcia, Capital One](https://medium.com/capital-one-tech/a-serverless -and-go-journey-credit-offers-api-74ef1f9fde7f)_

_“当时，没有一个团队成员知道 Go，但在一个月之内，每个人都在用 Go 编写代码”——[Jaime Garcia，Capital One](https://medium.com/capital-one-tech/a-serverless -and-go-journey-credit-offers-api-74ef1f9fde7f)_

_“What makes Go different from other programming languages is cognitive load. You can do more with less code, which makes it easier to reason about and understand the code that you do end up writing. The majority of Go code ends up looking quite similar, so, even if you’re working with a completely new codebase, you can get up and running pretty quickly.” — Glen Balliet Engineering Director of loyalty platforms at American Express [American Express Uses Go for Payments & Rewards](https://go.dev/solutions/americanexpress/)_

_“Go 与其他编程语言的不同之处在于认知负荷。你可以用更少的代码做更多的事情，这使得你更容易推理和理解你最终编写的代码。大多数 Go 代码最终看起来非常相似，因此，即使您正在使用全新的代码库，您也可以很快启动并运行。” — Glen Balliet 美国运通忠诚度平台工程总监 [美国运通使用 Go 进行支付和奖励](https://go.dev/solutions/americanexpress/)_

_“However, unlike other programming languages, Go was created for maximum user efficiency. Therefore developers and engineers with Java or PHP backgrounds can be upskilled and trained in using Go within a few weeks — and in our experience, many of them end up preferring it.” — [Dewet Diener, Curve](https://jaxenter.com/golang-curve-163187.html)_

_“然而，与其他编程语言不同，Go 的创建是为了最大限度地提高用户效率。因此，具有 Java 或 PHP 背景的开发人员和工程师可以在几周内获得使用 Go 的技能和培训——根据我们的经验，他们中的许多人最终更喜欢它。” — [Dewet Diener，曲线](https://jaxenter.com/golang-curve-163187.html)_

### 4\. Precise Control

### 4\。精准控制

Perhaps one of Rust’s greatest strengths is the amount of control the developer has over how memory is managed, how to use the available resources of the machine, how code is optimized, and how problem solutions are crafted. This is not without a large complexity cost when compared to Go, which is designed less for this type of precise crafting and more for faster exploration times and quicker turnaround times. 
也许 Rust 的最大优势之一是开发人员对如何管理内存、如何使用机器的可用资源、如何优化代码以及如何制定问题解决方案的控制量。与围棋相比，这并非没有很大的复杂性成本，围棋的设计目的不是为了这种精确的制作，而是为了更快的探索时间和更快的周转时间。
_“As our experience with Rust grew, it showed advantages on two other axes: as a language with strong memory safety it was a good choice for processing at the edge and as a language that had tremendous enthusiasm it became one that became popular for de novo components.” — John Graham-Cumming, Cloudflare_

_“随着我们对 Rust 的经验的增长，它在另外两个方面显示出优势：作为一种具有强大内存安全性的语言，它是边缘处理的不错选择；作为一种具有极大热情的语言，它成为一种流行的语言novo 组件。” — John Graham-Cumming，Cloudflare_

## Summary/Key Takeaways

## 总结/关键要点

Go’s simplicity, performance, and developer productivity make Go an ideal language for creating user-facing applications and services. The fast iteration allows teams to quickly pivot to meet the changing needs of users, giving teams a way to focus their energies on flexibility.

Go 的简单性、性能和开发人员生产力使 Go 成为创建面向用户的应用程序和服务的理想语言。快速迭代允许团队快速调整以满足用户不断变化的需求，让团队可以将精力集中在灵活性上。

Rust’s finer control allows for more precision, making Rust an ideal language for low-level operations that are less likely to change and that would benefit from the marginally improved performance over Go, especially if deployed at very large scales.

Rust 更精细的控制允许更高的精度，使 Rust 成为低级操作的理想语言，这些操作不太可能改变，并且会受益于比 Go 略微提高的性能，尤其是在非常大规模部署时。

Rust’s strengths are at the most advantageous closest to the metal. Go’s strengths are at their most advantageous closer to the user. This isn’t to say that either can’t work in the other’s space, but it would have increased friction to doing so. As your requirements shift from flexibility to efficiency it makes a stronger case to rewrite libraries in Rust.

Rust 的优势最接近金属。 Go 的优势是最靠近用户的地方。这并不是说任何一方都不能在对方的空间工作，但这样做会增加摩擦。随着您的需求从灵活性转向效率，在 Rust 中重写库是一个更有力的案例。

While the designs of Go and Rust differ significantly, their designs play to a compatible set of strengths, and — when used together — allow both great flexibility and performance.

虽然 Go 和 Rust 的设计有很大不同，但它们的设计发挥了一组兼容的优势，并且——当一起使用时——允许极大的灵活性和性能。

## Recommendations

## 建议

For most companies and users, Go is the right default option. Its performance is strong, Go is easy to adopt, and Go’s highly modular nature makes it particularly good for situations where requirements are changing or evolving.

对于大多数公司和用户来说，Go 是正确的默认选项。它的性能强大，Go 易于采用，Go 的高度模块化特性使其特别适合需求变化或演变的情况。

As your product matures, and requirements stabilize, there may be opportunities to have large wins from marginal increases in performance. In these cases, using Rust to maximize performance may well be worth the initial investment.

随着产品的成熟和需求的稳定，可能有机会从性能的边际提高中获得巨大的成功。在这些情况下，使用 Rust 来最大化性能可能是值得的初始投资。

Amazon Web Services is a sponsor of The New Stack. 
Amazon Web Services 是 The New Stack 的赞助商。
