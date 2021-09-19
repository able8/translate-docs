# Why we’re writing machine learning infrastructure in Go, not Python

# 为什么我们用 Go 而不是 Python 编写机器学习基础设施

## Production machine learning is about more than just algorithms

## 生产机器学习不仅仅是算法

[Jan 8, 2020·5 min read](http://medium.com/@calebkaiser/why-were-writing-machine-learning-infrastructure-in-go-not-python-38d6a37e2d76?source=post_page-----38d6a37e2d76--------------------------------)

At this point, it should be a surprise to no one that Python is the [most popular](https://github.blog/2019-01-24-the-state-of-the-octoverse-machine-learning/) language for machine learning projects. While languages like R, C++, and Julia have their proponents—and use cases—Python remains the most universally embraced language, being used in every major machine learning framework.

[2020 年 1 月 8 日·5 分钟阅读](http://medium.com/@calebkaiser/why-were-writing-machine-learning-infrastructure-in-go-not-python-38d6a37e2d76?source=post_page--- 在这一点上，Python 是 [最流行的](https://github.blog/2019-01-24-the-state-of-the-octoverse-machine-learning/)应该不会让任何人感到惊讶机器学习项目的语言。虽然像 R、C++ 和 Julia 等语言有他们的支持者和用例，但 Python 仍然是最普遍接受的语言，被用于每个主要的机器学习框架。

So, naturally, our codebase at [Cortex](https://github.com/cortexlabs/cortex)—an open source platform for deploying machine learning models as APIs—is 87.5% Go.

因此，很自然地，我们在 [Cortex](https://github.com/cortexlabs/cortex)（一个将机器学习模型部署为 API 的开源平台)的代码库是 87.5% 的 Go。

![](https://miro.medium.com/max/60/1*jfWsTOsdPWlxjS4GuQkJLg.png?q=20)

Source: [Cortex GitHub](https://github.com/cortexlabs/cortex)

来源：[Cortex GitHub](https://github.com/cortexlabs/cortex)

Machine learning algorithms, where Python shines, are just one component of a production machine learning system. To actually run a production machine learning API at scale, you need infrastructure that implements features like:

Python 大放异彩的机器学习算法只是生产机器学习系统的一个组成部分。要实际大规模运行生产机器学习 API，您需要实现以下功能的基础架构：

- Autoscaling, so that traffic fluctuations don’t break your API
- API management, to handle simultaneous API deployments
- Rolling updates, so that you can update models while still serving users

- 自动缩放，以便流量波动不会破坏您的 API
- API 管理，处理同时的 API 部署
- 滚动更新，让您在更新模型的同时仍然为用户服务

[Cortex](http://cortex.dev) is built to automate all of this infrastructure, along with other concerns like logging and cost optimizations.

[Cortex](http://cortex.dev) 旨在自动化所有这些基础设施，以及其他问题，如日志记录和成本优化。

Go is ideal for building software with these considerations, for a few reasons:

Go 是基于这些考虑构建软件的理想选择，原因如下：

# 1\. Concurrency is crucial for machine learning infrastructure

# 1\.并发对于机器学习基础设施至关重要

A user can have many different models deployed as distinct APIs, all managed in the same Cortex cluster. In order for the Cortex Operator to manage these different deployments, it needs to wrangle a few different APIs. To name a couple:

用户可以将许多不同的模型部署为不同的 API，所有模型都在同一个 Cortex 集群中进行管理。为了让 Cortex Operator 管理这些不同的部署，它需要处理一些不同的 API。命名一对：

- Kubernetes APIs, which Cortex calls to deploy models on the cluster.
- Various AWS APIs—EC2 Auto Scaling, S3, CloudWatch, and others—which Cortex calls to manage deployments on AWS.

- Kubernetes API，Cortex 调用这些 API 在集群上部署模型。
- 各种 AWS API——EC2 Auto Scaling、S3、CloudWatch 等——Cortex 调用这些 API 来管理 AWS 上的部署。

The user doesn’t interact with any of these APIs directly. Instead, Cortex programmatically calls these APIs to provision clusters, launch deployments, and monitor APIs.

用户不直接与任何这些 API 交互。相反，Cortex 以编程方式调用这些 API 来提供集群、启动部署和监控 API。

Making all of these overlapping API calls in a performative, reliable way is a challenge. Handling them concurrently is the most efficient way to do things, but it also introduces complexity, as now we have to worry about things like race conditions.

以一种高效、可靠的方式进行所有这些重叠的 API 调用是一项挑战。同时处理它们是最有效的做事方式，但它也引入了复杂性，因为现在我们不得不担心竞争条件之类的事情。

Go has an elegant, out of the box solution for this problem: **Goroutines**.

Go 有一个优雅的、开箱即用的解决方案：**Goroutines**。

Goroutines are otherwise normal functions that Go executes concurrently. We could write an entire article digging into how Goroutines work under the hood, but at a high level, Goroutines are lightweight threads managed automatically by the Go runtime. Many Goroutines can fit on a single OS thread, and if a Goroutine blocks an OS thread, the Go runtime automatically moves the rest of the Goroutines over to a new OS thread.

Goroutines 是 Go 并发执行的普通函数。我们可以写一整篇文章来深入研究 Goroutines 在底层是如何工作的，但在较高的层次上，Goroutines 是由 Go 运行时自动管理的轻量级线程。许多 Goroutine 可以安装在单个 OS 线程上，如果 Goroutine 阻塞了 OS 线程，Go 运行时会自动将其余的 Goroutine 转移到新的 OS 线程。

Goroutines also offer a feature called “channels,” which allow Goroutines to pass messages between themselves, allowing us to schedule requests and prevent race conditions.

Goroutines 还提供了一个称为“通道”的功能，它允许 Goroutines 在它们之间传递消息，允许我们安排请求并防止竞争条件。

Implementing all of this functionality in Python may be doable with recent tools like asyncio, but the fact that Go is designed with this use case in mind makes our lives much easier.

在 Python 中实现所有这些功能可能可以使用 asyncio 等最近的工具来实现，但是 Go 的设计考虑到了这个用例这一事实让我们的生活变得更加轻松。

# 2\. Building a cross-platform CLI is easier in Go

# 2\.在 Go 中构建跨平台 CLI 更容易

The Cortex CLI is a cross-platform tool that allows users to deploy models and manage APIs directly from the command line. The below GIF shows the CLI in action:

Cortex CLI 是一个跨平台工具，允许用户直接从命令行部署模型和管理 API。下面的 GIF 显示了正在运行的 CLI：

![](https://miro.medium.com/max/60/0*T17nRCxfd6j9gbB5?q=20)

Source: [Cortex GitHub](https://github.com/cortexlabs/cortex)

来源：[Cortex GitHub](https://github.com/cortexlabs/cortex)

Originally, we wrote the CLI in Python, but trying to distribute it across platforms proved to be too difficult. Because Go compiles down to a single binary—no dependency management required—it offered us a simple solution to distributing our CLI across platforms without requiring much extra engineering effort. 

最初，我们用 Python 编写 CLI，但尝试跨平台分发它被证明太困难了。因为 Go 编译成单个二进制文件——不需要依赖管理——它为我们提供了一个简单的解决方案，可以在不需要太多额外工程工作的情况下跨平台分发我们的 CLI。

The performance benefits of a compiled Go binary versus an interpreted language are also significant. According to the computer benchmarks game, Go is dramatically [faster than Python](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-python3.html).

与解释性语言相比，编译后的 Go 二进制文件的性能优势也很显着。根据计算机基准测试游戏，Go 显着[比 Python 快](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-python3.html)。

It’s perhaps not coincidental that many other infrastructure CLI tools are written in Go, which brings us to our next point.

许多其他基础设施 CLI 工具都是用 Go 编写的，这可能并非巧合，这将我们带到了下一点。

# 3\. The Go ecosystem is great for infrastructure projects

# 3\. Go 生态系统非常适合基础设施项目

One of the benefits of open source is that you can learn from the projects you admire. For example, Cortex exists within the Kubernetes (which itself is written in Go) ecosystem. We were fortunate to have a number of great open source projects within that ecosystem to learn from, including:

开源的好处之一是你可以从你欣赏的项目中学习。例如，Cortex 存在于 Kubernetes（它本身是用 Go 编写的）生态系统中。我们有幸在该生态系统中有许多优秀的开源项目可供学习，包括：

- [**kubectl**](https://github.com/kubernetes/kubectl) **:** Kubernetes’ CLI
- [**minikube**](https://github.com/kubernetes/minikube): A tool for running Kubernetes locally
- [**helm**](https://github.com/helm/helm): A Kubernetes package manager
- [**kops**](https://github.com/kubernetes/kops) **:** A tool for managing production Kubernetes
- [**eksctl**](https://github.com/weaveworks/eksctl): The official CLI for Amazon EKS

- [**kubectl**](https://github.com/kubernetes/kubectl) **:** Kubernetes 的 CLI
- [**minikube**](https://github.com/kubernetes/minikube)：本地运行Kubernetes的工具
- [**helm**](https://github.com/helm/helm)：Kubernetes 包管理器
- [**kops**](https://github.com/kubernetes/kops) **:** 管理生产 Kubernetes 的工具
- [**eksctl**](https://github.com/weaveworks/eksctl)：Amazon EKS 的官方 CLI

All of the above are written in Go—and it’s not just Kubernetes projects. Whether you're looking at [CockroachDB](https://github.com/cockroachdb/cockroach) or [Hashicorp's](https://github.com/hashicorp) infrastructure projects, including [Vault](https://github.com/hashicorp/vault), [Nomad](https://github.com/hashicorp/nomad),[Terraform](https://github.com/hashicorp/terraform), [Consul](https://github.com/hashicorp/consul), and [Packer](https://github.com/hashicorp/packer), all of them are written in Go.

以上所有内容都是用 Go 编写的——而且不仅仅是 Kubernetes 项目。无论您是在查看 [CockroachDB](https://github.com/cockroachdb/cockroach) 或 [Hashicorp's](https://github.com/hashicorp) 基础设施项目，包括 [Vault](https://github.com/hashicorp/vault)、[Nomad](https://github.com/hashicorp/nomad)、[Terraform](https://github.com/hashicorp/terraform)、[Consul](https://github.com/hashicorp/consul) 和 [Packer](https://github.com/hashicorp/packer)，它们都是用 Go 编写的。

The popularity of Go in the infrastructure world has another effect, which is that most engineers interested in working on infrastructure are familiar with Go. This makes it easier to attract engineers. In fact, one of the best engineers at Cortex Labs found us by searching for Go jobs on AngelList—and we are lucky he found us.

Go 在基础设施领域的流行还有另一个影响，那就是大多数对基础设施感兴趣的工程师都熟悉 Go。这样更容易吸引工程师。事实上，Cortex Labs 最好的工程师之一通过在 AngelList 上搜索 Go 工作找到了我们——我们很幸运他找到了我们。

# 4\. Go is just a pleasure to work with

# 4. 与 Go 一起工作很愉快

The final note I’ll make on why we ultimately built Cortex in Go is that Go is just _nice_.

关于我们最终在 Go 中构建 Cortex 的原因，我要说明的最后一点是 Go 只是 _nice_。

Relative to Python, Go is a bit more painful to get started with. Go’s unforgiving nature, however, is what makes it such a joy for large projects. We still heavily test our software, but static typing and compilation — two things that make Go a bit less comfortable for beginners — act as sort of guard rails for us, helping us to write (relatively) bug-free code.

相对于 Python，Go 上手要痛苦一些。然而，Go 无情的本质使它成为大型项目的乐趣。我们仍然对我们的软件进行大量测试，但是静态类型和编译——这两个让 Go 初学者不太舒服的两件事——对我们起到了某种保护作用，帮助我们编写（相对）无错误的代码。

There may be other languages you could argue offer a particular advantage, but on balance, Go best satisfies our technical and aesthetic needs.

您可能会争辩说其他语言提供了特殊的优势，但总的来说，Go 最能满足我们的技术和审美需求。

# Python for machine learning, Go for infrastructure

# Python 用于机器学习，Go 用于基础设施

We still love Python, and it has its place within Cortex, specifically around inference processing.

我们仍然喜欢 Python，它在 Cortex 中占有一席之地，特别是在推理处理方面。

Cortex serves TensorFlow, PyTorch, scikit-learn, and other Python models, which means that interfacing with the models—as well as pre and post inference processing—are done in Python. However, even that Python code is packaged up into Docker containers, which are orchestrated by code that is written in Go.

Cortex 为 TensorFlow、PyTorch、scikit-learn 和其他 Python 模型提供服务，这意味着与模型的接口以及前后推理处理都是在 Python 中完成的。然而，即使是 Python 代码也被打包到 Docker 容器中，这些容器是由用 Go 编写的代码编排的。

If you’re interested in becoming a machine learning engineer, knowing Python is more or less non-negotiable. If you’re interested in working on machine learning _infrastructure_, however, you should seriously consider using Go.

如果您有兴趣成为机器学习工程师，那么了解 Python 或多或少是不容商量的。但是，如果您对机器学习 _infrastructure_ 感兴趣，则应该认真考虑使用 Go。

_Are you an engineer interested in Go and machine learning? If so, consider_ [_contributing to Cortex_](https://github.com/cortexlabs/cortex) _!_

_你是对围棋和机器学习感兴趣的工程师吗？如果是这样，请考虑_[_contributing to Cortex_](https://github.com/cortexlabs/cortex)_!_

[**Caleb Kaiser**](http://medium.com/@calebkaiser?source=post_sidebar--------------------------post_sidebar-----------)

ML infrastructure ( [https://github.com/cortexlabs/cortex](https://github.com/cortexlabs/cortex)) Formerly at AngelList. Originally a Cadillac. 

ML 基础设施 ([https://github.com/cortexlabs/cortex](https://github.com/cortexlabs/cortex)) 以前在 AngelList。原来是凯迪拉克。

