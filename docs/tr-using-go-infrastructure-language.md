# Using Go as the Infrastructure Language at NodeSource

# 在 NodeSource 中使用 Go 作为基础设施语言

in [NodeSource](http://nodesource.com/blog/category/nodesource)
on Mar 08 2019

NodeSource is a cloud-first company, with resources on AWS, Google Cloud Platform, and Azure. Traditionally Python was a common choice of programming language as it could be used by infrastructure teams. But as Google developed Go to solve problems that Python could not, Go quickly gained new users over the course of just a few years. Go has become the more popular option, and many tools that were once a staple for infrastructure in Python now have a Go version. Some of these tools, like graphite, were originally created in Python but have been recreated for use with Go due to their popularity and growing demand.

NodeSource 是一家云优先的公司，在 AWS、谷歌云平台和 Azure 上拥有资源。传统上，Python 是一种常见的编程语言选择，因为它可供基础设施团队使用。但随着 Google 开发 Go 来解决 Python 无法解决的问题，Go 在短短几年内迅速获得了新用户。 Go 已经成为更受欢迎的选择，许多曾经是 Python 基础设施主要工具的工具现在都有 Go 版本。其中一些工具，如石墨，最初是用 Python 创建的，但由于它们的流行和不断增长的需求，已经重新创建以与 Go 一起使用。

Every organization has its own set of needs and constraints, so while I’m a big fan of Go, I recognize that it isn’t the right choice for everyone. For NodeSource, some features in Go were of particular interest and use to us, and made Go the better choice for us as we built infrastructure, including:

每个组织都有自己的需求和限制，所以虽然我是 Go 的忠实粉丝，但我认识到它不是每个人的正确选择。对于 NodeSource，Go 中的一些特性对我们特别感兴趣和使用，这使 Go 成为我们构建基础设施时更好的选择，包括：

- [Static linking](https://en.wikipedia.org/wiki/Static_library): Static linking makes it easy to deploy what I build. All Go binaries are statically linked, and copying them to the target platform allows for easy execution without managing dependencies. In our case, I don’t even need to install the runtime when creating a new instance, since that too is included in the binary. This is a much simpler task than it would be if I used Python, which requires that the runtime as well as dependencies (requirements.txt) be installed before code can run. With Go, if you can build the binary; you can deploy the binary.
- [静态链接](https://en.wikipedia.org/wiki/Static_library)：静态链接可以轻松部署我构建的内容。所有 Go 二进制文件都是静态链接的，将它们复制到目标平台可以轻松执行，而无需管理依赖项。在我们的例子中，我什至不需要在创建新实例时安装运行时，因为它也包含在二进制文件中。这是一个比我使用 Python 时要简单得多的任务，它要求在代码运行之前安装运行时和依赖项 (requirements.txt)。使用 Go，如果您可以构建二进制文件；您可以部署二进制文件。
- [Goroutines](https://tour.golang.org/concurrency/1): We live in a world with multiple CPUs (or vCPUs if you prefer). To maximize performance and computing power, we need threading or multiprocessing. However, threading can be painful (insert your favorite pthread or Java threading jokes here). So with Go, threading is as simple as `go myfunc()`; it has never been easier to thread with a language such as Go.
- [Goroutines](https://tour.golang.org/concurrency/1)：我们生活在一个拥有多个 CPU（或 vCPU，如果您愿意的话）的世界。为了最大限度地提高性能和计算能力，我们需要线程或多处理。然而，线程可能很痛苦（在此处插入您最喜欢的 pthread 或 Java 线程笑话)。所以使用 Go，线程就像 `go myfunc()` 一样简单；使用 Go 等语言进行线程处理从未如此简单。

The above two benefits are more or less "core" to the language itself; but, with any language comes the need for third-party modules.

以上两个好处或多或少是语言本身的“核心”；但是，任何语言都需要第三方模块。

A few additional modules which I consider to be very important for writing Go programs include:

我认为对于编写 Go 程序非常重要的一些附加模块包括：

- [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper) make it very easy to create software that supports loading configuration files and having command- line switches. Any moderately-important software should have at minimum CLI switches, and any really important software should probably load its configuration from a JSON or YAML file. Both of these modules allow you to do that.
- [logrus](https://github.com/Sirupsen/logrus) is another important module for logging. It's not required, but makes software feel more polished and ready for production. This is definitely preferable to just using `fmt.Printf` statements.
- [jsonq](https://github.com/jmoiron/jsonq) this library makes it easier to read JSON data, even though it is already easy to marshal/unmarshal JSON-encoded data in Go.

- [cobra](https://github.com/spf13/cobra) 和 [viper](https://github.com/spf13/viper) 使创建支持加载配置文件和命令的软件变得非常容易-线路开关。任何中等重要的软件都应该至少有 CLI 开关，任何真正重要的软件都应该从 JSON 或 YAML 文件加载其配置。这两个模块都允许您这样做。
- [logrus](https://github.com/Sirupsen/logrus) 是另一个重要的日志模块。它不是必需的，但会使软件感觉更加精美并准备好用于生产。这绝对比仅使用 `fmt.Printf` 语句更可取。
- [jsonq](https://github.com/jmoiron/jsonq) 这个库使读取 JSON 数据变得更容易，即使在 Go 中编组/解组 JSON 编码的数据已经很容易了。

Finally, because our team uses Go as our preferred cloud infrastructure language, I would also suggest the following critical modules that are used to interact with our cloud environments:

最后，因为我们的团队使用 Go 作为我们首选的云基础设施语言，我还建议使用以下用于与我们的云环境交互的关键模块：

- [aws-sdk-go](https://github.com/aws/aws-sdk-go) is the official AWS SDK for Go. I have used it extensively for interacting with AWS services including (but not limited to) S3, ECS, and EC2.
- The [Slack API in Go](https://github.com/nlopes/slack) allows me to post messages to slack from anything written in Go—as part of a Slack-friendly team, this has been critical for me.
- The [Docker SDK for Go](https://docs.docker.com/develop/sdk/#go-sdk) is the module to write Go code that can interact with a Docker host; critical if you’re planning to write any of your own tooling to pull data from your containerized environment. 

- [aws-sdk-go](https://github.com/aws/aws-sdk-go) 是 Go 的官方 AWS 开发工具包。我广泛使用它与 AWS 服务交互，包括（但不限于)S3、ECS 和 EC2。
- [Slack API in Go](https://github.com/nlopes/slack) 允许我从 Go 编写的任何内容中向 slack 发布消息——作为 Slack 友好团队的一部分，这对我来说至关重要。
- [Docker SDK for Go](https://docs.docker.com/develop/sdk/#go-sdk) 是编写可以与Docker主机交互的Go代码的模块；如果您打算编写任何自己的工具来从容器化环境中提取数据，那么这一点至关重要。

- The official [Go client library for Prometheus](https://prometheus.io/docs/guides/go-application/) allows me to define and expose Prometheus metrics from any Go binary I’ve written.

- 官方[Prometheus 的 Go 客户端库](https://prometheus.io/docs/guides/go-application/) 允许我从我编写的任何 Go 二进制文件中定义和公开 Prometheus 指标。

### One More Thing....

###  还有一件事....

If you're an infrastructure professional and interested in working with NodeSource, we're currently looking for an awesome [Site Reliability Engineer](https://nodesource.bamboohr.com/jobs/view.php?id=31) to join our team. Please check out our [careers page](https://nodesource.com/careers) for other open positions and to apply. 

如果您是基础设施专业人士并且对使用 NodeSource 感兴趣，我们目前正在寻找一位出色的 [站点可靠性工程师](https://nodesource.bamboohr.com/jobs/view.php?id=31)加入我们的队伍。请查看我们的[职业页面](https://nodesource.com/careers) 了解其他空缺职位并申请。

