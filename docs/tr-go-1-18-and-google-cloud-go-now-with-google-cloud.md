# Go 1.18 and Google Cloud: Go now with Google Cloud

March 21, 2022

2022 年 3 月 21 日

On March 15th, the Go team [announced](https://go.dev/blog/go1.18) Go 1.18 GA, the latest release of the Go programming language. The culmination of over a decade of design delivers the features our developers demanded: generics, fuzzing, and module workspaces. With this release, Go becomes the first major language to integrate fuzz testing into its core toolchain without using third-party support, further establishing Go as a preferred language for developing secure applications.

3 月 15 日，Go 团队 [宣布](https://go.dev/blog/go1.18) Go 1.18 GA，Go 编程语言的最新版本。十多年设计的巅峰之作提供了我们的开发人员所需的功能：泛型、模糊测试和模块工作区。在此版本中，Go 成为第一个在不使用第三方支持的情况下将模糊测试集成到其核心工具链中的主要语言，进一步将 Go 确立为开发安全应用程序的首选语言。

Go was created at Google in 2007, designed to help developers build fast, reliable, and secure software. Unlike traditional languages, Go was built for the modern multi-core computing world. Go has emerged as a modern language for developing cloud applications, services, and infrastructure. Today Go powers several of Google’s largest products, and is used by many customers to scale their businesses. Organizations big and small love Go and the community of Go developers, known as “gophers” has grown into a global network with over 2 million users worldwide.

Go 于 2007 年在 Google 创建，旨在帮助开发人员构建快速、可靠和安全的软件。与传统语言不同，Go 是为现代多核计算世界而构建的。 Go 已经成为一种用于开发云应用程序、服务和基础架构的现代语言。如今，Go 为 Google 的几款最大产品提供支持，并被许多客户用于扩展业务。大大小小的组织都喜欢 Go，而被称为“gophers”的 Go 开发者社区已经发展成为一个在全球拥有超过 200 万用户的全球网络。

### Using the power of Go in the Cloud 

### 在云端使用 Go 的强大功能

When looking at the public repos, over 75% of [CNCF](https://www.cncf.io/) projects including Kubernetes are written in Go and [10% of developers](https://insights.stackoverflow.com/survey/2021#most-popular-technologies-language) are writing in Go worldwide (as of May 2021). Google delivers high performance infrastructure to run key, cloud native, Open Source projects. Our modern cloud infrastructure is based on Kubernetes at its core and our strong support for Istio and Knative have formed the base of some of our leading services like Google Kubernetes Engine (GKE), our managed application platform with Anthos, Cloud Functions, and Cloud Run . Google uses Go extensively for a wide range of applications from our [indexing platform that powers Google Search](https://go.dev/solutions/google/coredata), to the [server side optimizations that power Chrome's 1B+ users](https://go.dev/solutions/google/chrome), to the [infrastructure on which Google cloud is built](https://go.dev/solutions/google/sitereliability).

查看公开回购时，超过 75% 的 [CNCF](https://www.cncf.io/) 项目包括 Kubernetes 是用 Go 编写的，[10% 的开发人员](https://insights.stackoverflow.com/survey/2021#most-popular-technologies-language)在全球范围内使用 Go 编写（截至 2021 年 5 月）。 Google 提供高性能基础架构来运行关键的云原生开源项目。我们的现代云基础架构以 Kubernetes 为核心，我们对 Istio 和 Knative 的大力支持构成了我们一些领先服务的基础，例如 Google Kubernetes Engine (GKE)、我们的托管应用程序平台 Anthos、Cloud Functions 和 Cloud Run . Google 将 Go 广泛用于各种应用程序，从我们的 [为 Google 搜索提供支持的索引平台](https://go.dev/solutions/google/coredata) 到 [为 Chrome 的 1B+ 用户提供支持的服务器端优化](https://go.dev/solutions/google/chrome)，到[构建谷歌云的基础设施](https://go.dev/solutions/google/sitereliability)。

### Release Highlights

### 发布亮点

With this new release of Go 1.18, Generics are the biggest change to Go since the language was created. Go developers told us that they feel that Go lacks critical features, with generics being the main missing piece. With Go 1.18, new and existing Go developers can take advantage of the productivity, performance, and maintenance benefits that generics can bring. We’ve already begun to see the new kinds of libraries and projects gophers are building with generics in its short beta period, and expect this creativity to grow as time goes on.

在 Go 1.18 的这个新版本中，泛型是 Go 自语言创建以来最大的变化。 Go 开发人员告诉我们，他们认为 Go 缺乏关键特性，泛型是主要缺失的部分。使用 Go 1.18，新的和现有的 Go 开发人员可以利用泛型可以带来的生产力、性能和维护优势。我们已经开始看到 gophers 在其短暂的 beta 阶段使用泛型构建的新型库和项目，并期望这种创造力随着时间的推移而增长。

This Go release also brings native support for fuzzing. Fuzzing is a type of vulnerability testing that throws arbitrary data at a piece of software to expose unknown errors and is emerging as a common testing scheme in enterprise development. Go is now the first major language to provide fuzzing support with no third-party integrations necessary, allowing developers to start building secure software with minimal additional cost. Go’s innovative approach to fuzzing can provide not only security for the current code but also ongoing protection as code and dependencies evolve. With attacks on software becoming more common and complex, vulnerability detection can be a critical part of the enterprise development lifecycle, and Go’s fuzzing capabilities catch vulnerabilities earlier in the lifecycle.

这个 Go 版本还带来了对模糊测试的原生支持。 Fuzzing 是一种漏洞测试类型，它向一个软件抛出任意数据以暴露未知错误，并且正在成为企业开发中的一种常见测试方案。 Go 现在是第一个无需第三方集成即可提供模糊测试支持的主要语言，允许开发人员以最低的额外成本开始构建安全软件。 Go 创新的 fuzzing 方法不仅可以为当前代码提供安全性，还可以随着代码和依赖项的发展提供持续的保护。随着对软件的攻击变得越来越普遍和复杂，漏洞检测可能成为企业开发生命周期的关键部分，而 Go 的 fuzzing 功能可以在生命周期的早期捕获漏洞。

### Build securely using Go 

### 使用 Go 安全构建

At Google we are helping to make Open Source software secure. Open source software is a connective tissue for much of the online world. At Google, we've been working to [raise awareness](https://security.googleblog.com/2021/02/know-prevent-fix-framework-for-shifting.html) of the state of open source security and are committed to helping secure the software supply chain for organizations. Go has been designed to create secure applications, helping to minimize risk as much as possible. Go applications compile down to a single binary without local dependencies. It’s not uncommon to see an application built using only the standard library, or only a couple well-vetted Go dependencies. Go’s dependency management uses tamper-evident [transparency log](https://transparency.dev), with built in tooling that helps ensure your dependencies are what you can expect. Go has native encryption, which is used across much of the internet, including key components of Google. Go even supports distroless containers, where there are zero local dependencies to worry about. Google Cloud products like [Cloud Build](https://cloud.google.com/build), for CI/CDand [Artifact Registry](https://cloud.google.com/artifact-registry), for container management, and have direct access to Go's vulnerability database and can provide you instant warnings about security threats.

在 Google，我们正在帮助确保开源软件的安全。开源软件是大部分网络世界的结缔组织。在 Google，我们一直在努力 [提高人们对开源安全状态的认识](https://security.googleblog.com/2021/02/know-prevent-fix-framework-for-shifting.html) 和致力于帮助保护组织的软件供应链。 Go 旨在创建安全的应用程序，帮助尽可能降低风险。 Go 应用程序编译成没有本地依赖关系的单个二进制文件。仅使用标准库或仅使用几个经过严格审查的 Go 依赖项构建的应用程序并不少见。 Go 的依赖管理使用防篡改 [transparency log](https://transparency.dev)，内置工具可帮助确保您的依赖关系符合您的预期。 Go 具有本机加密，它在互联网上广泛使用，包括谷歌的关键组件。 Go 甚至支持 distroless 容器，无需担心本地依赖项。 Google Cloud 产品，例如 [Cloud Build](https://cloud.google.com/build)，用于 CI/CD 和 [Artifact Registry](https://cloud.google.com/artifact-registry)，用于容器管理，并且可以直接访问 Go 的漏洞数据库，并且可以为您提供有关安全威胁的即时警告。

_**“At Google we are committed to helping to secure the online infrastructure and applications upon which the world depends. A critical aspect of this mission is being able to understand and verify the security of open source dependency chains. The 1.18 release of Go is an important step towards helping to ensure that developers are able to build secure applications, understand risk when vulnerabilities are discovered, and reduce the impact of cybersecurity attacks”** said Eric Brewer, VP Infrastructure, Google Fellow_

_**“在 Google，我们致力于帮助保护世界所依赖的在线基础设施和应用程序的安全。该任务的一个关键方面是能够理解和验证开源依赖链的安全性。 Go 的 1.18 版本是朝着帮助确保开发人员能够构建安全的应用程序、在发现漏洞时了解风险以及减少网络安全攻击的影响迈出的重要一步”** Google Fellow_基础设施副总裁 Eric Brewer 说

This launch is a significant milestone for Go that helps developers from around the world build more performant and secure applications that run on any infrastructure. For more information on this [release](https://go.dev/blog/go1.18) and how to [get started](https://go.dev/blog/) with Go, [please visit]( https://go.dev/).

此次发布是 Go 的一个重要里程碑，它帮助来自世界各地的开发人员构建在任何基础设施上运行的性能更高、更安全的应用程序。有关此 [发布](https://go.dev/blog/go1.18) 以及如何使用 Go [开始](https://go.dev/blog/) 的更多信息，[请访问]( https://go.dev/)。



https://cloud.google.com/blog/products/gcp/go-1-18-and-google-cloud-go-now-with-google-cloud

