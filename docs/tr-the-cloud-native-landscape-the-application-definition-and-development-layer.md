# The Cloud Native Landscape: The Application Definition and Development Layer

# 云原生景观：应用定义和开发层

#### 26 Jan 2021 6:00am,   by [Catherine Paganini](https://thenewstack.io/author/catherine-paganini/ "Posts by Catherine Paganini") and [Jason Morgan](https://thenewstack. io/author/jason-morgan/ "Posts by Jason Morgan")

#### 2021 年 1 月 26 日上午 6:00，作者：[Catherine Paganini](https://thenewstack.io/author/catherine-paganini/“Catherine Paganini 的帖子”) 和 [Jason Morgan](https://thenewstack.io/author/jason-morgan/“杰森摩根的帖子”)

![](https://cdn.thenewstack.io/media/2021/01/4fd4a803-cncf-landscape.png)

_This post is part of an ongoing series from [Cloud Native Computing Foundation Business Value Subcommittee](https://lists.cncf.io/g/cncf-business-value) co-chairs [Catherine Paganini](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) and [Jason Morgan](https://thenewstack.io/author/jason-morgan/) that focuses on explaining each category of the cloud native landscape to a non -technical audience as well as engineers just getting started with cloud native computing._

_这篇文章是 [云原生计算基金会商业价值小组委员会](https://lists.cncf.io/g/cncf-business-value) 联合主席 [Catherine Paganini](https://landscape.cncf.io/category=coordination-service-discovery&grouping=category) 和 [Jason Morgan](https://thenewstack.io/author/jason-morgan/) 专注于向非- 刚开始使用云原生计算的技术受众和工程师。_

When looking at the [Cloud Native Landscape](https://landscape.cncf.io/), you’ll note a few distinctions:

在查看 [云原生景观](https://landscape.cncf.io/) 时，您会注意到一些区别：

- Projects in large boxes are CNCF-hosted open source projects. Some are still in the incubation phase (light blue/purple frame), while others are graduated projects (dark blue frame).
- Projects in small white boxes are open source projects.
- Products in gray boxes are proprietary.

- 大盒子中的项目是 CNCF 托管的开源项目。有的还在孵化阶段（浅蓝色/紫色框），有的则是毕业项目（深蓝色框）。
- 小白盒中的项目是开源项目。
- 灰色框中的产品是专有的。

Please note that even during the time of this writing, we saw new projects becoming part of the CNCF so always refer to the actual landscape — things are moving fast!

请注意，即使在撰写本文时，我们也看到新项目成为 CNCF 的一部分，因此请始终参考实际情况——事情进展很快！

## Database [![databases](https://cdn.thenewstack.io/media/2021/01/ef82aef4-screen-shot-2021-01-24-at-5.50.05-pm.png)](https://landscape.cncf.io/)

://landscape.cncf.io/)

### What It Is

###  这是什么

A database management system is an application through which other apps can efficiently store and retrieve data.

数据库管理系统是一个应用程序，通过它其他应用程序可以有效地存储和检索数据。

It ensures data gets stored, only authorized users are able to access it, and allows users to retrieve it via specialized requests. While there are numerous different types of databases with different approaches, they ultimately all have these same overarching goals.

它确保数据被存储，只有授权用户才能访问它，并允许用户通过专门的请求来检索它。虽然有许多不同类型的数据库采用不同的方法，但它们最终都有相同的总体目标。

### Problem It Addresses

### 它解决的问题

Most applications need an effective way to store and retrieve data while keeping that data safe. Databases do this in a structured way with proven technology (there is quite a bit of complexity that goes into doing this well).

大多数应用程序需要一种有效的方法来存储和检索数据，同时保证数据的安全。数据库以结构化的方式使用经过验证的技术来做到这一点（要做到这一点有相当多的复杂性）。

### How It Helps

### 它如何帮助

Databases provide a common interface for storing and retrieving data for applications. Developers use these standard interfaces and, in most cases a simple query language, to store, query, and retrieve information in a database. At the same time, databases allow users to continuously back up and save data as well as encrypt and regulate access to it.

数据库为应用程序存储和检索数据提供了一个通用接口。开发人员使用这些标准接口，在大多数情况下使用简单的查询语言，在数据库中存储、查询和检索信息。同时，数据库允许用户持续备份和保存数据以及加密和管理对数据的访问。

### Technical 101

### 技术 101

We’ve established that a database management system is an application that stores and retrieves data. It does so using a common language and interface that can be easily used by a number of different languages and frameworks.

我们已经确定数据库管理系统是一个存储和检索数据的应用程序。它使用一种通用语言和界面来实现，这些语言和界面可以被许多不同的语言和框架轻松使用。

In general, we see two common types of databases, structured query language (SQL) databases, and no-SQL databases. Which database a particular application uses should be driven by its needs and constraints.

一般来说，我们看到两种常见类型的数据库，结构化查询语言 (SQL) 数据库和非 SQL 数据库。特定应用程序使用哪个数据库应由其需求和约束驱动。

With the rise of Kubernetes and its ability to support stateful applications, we’ve seen a new generation of databases that leverage containerization. These new cloud native databases aim to bring the scaling and availability benefits of Kubernetes to databases. Tools like [YugaByte](https://landscape.cncf.io/?selected=yuga-byte-db) and [Couchbase](https://landscape.cncf.io/?selected=couchbase) are examples of cloud native databases, although more traditional databases like MySQL and Postgres run successfully and effectively in Kubernetes clusters.

随着 Kubernetes 的兴起及其支持有状态应用程序的能力，我们看到了利用容器化的新一代数据库。这些新的云原生数据库旨在为数据库带来 Kubernetes 的可扩展性和可用性优势。 [YugaByte](https://landscape.cncf.io/?selected=yuga-byte-db) 和 [Couchbase](https://landscape.cncf.io/?selected=couchbase) 等工具是云原生的例子数据库，尽管更传统的数据库（如 MySQL 和 Postgres)在 Kubernetes 集群中成功且有效地运行。

[Vitess](https://landscape.cncf.io/?selected=vitess) and [TiKV](https://landscape.cncf.io/?selected=ti-kv) are CNCF projects in this space.

[Vitess](https://landscape.cncf.io/?selected=vitess) 和 [TiKV](https://landscape.cncf.io/?selected=ti-kv) 是 CNCF 在这个领域的项目。

### Sidenote 

###  边注

If you look at this category, you’ll notice multiple names ending in DB (e.g. MongoDB, CockroachDB, FaunaDB) which, as you may guess, stands for database. You’ll also see various names ending in SQL (e.g. MySQL or memSQL) — they are still relevant. Some are “old school” databases that have been adapted to a cloud native reality. There are also some databases that are no-SQL but SQL compatible, such as YugaByte and Vitess.

如果您查看此类别，您会注意到多个以 DB 结尾的名称（例如 MongoDB、CockroachDB、FaunaDB），您可能会猜到，它们代表数据库。您还会看到以 SQL 结尾的各种名称（例如 MySQL 或 memSQL）——它们仍然相关。有些是适应云原生现实的“老派”数据库。也有一些非SQL但兼容SQL的数据库，例如YugaByte和Vitess。

BuzzwordsPopular Projects

流行语热门项目

- SQL
- DB
- Persistence

- SQL
- D B
- 坚持

- Postgres
- MySQL
- Redis

- Postgres
- MySQL
- Redis

## Streaming and Messaging

## 流媒体和消息传递

[![streaming and messaging ](https://cdn.thenewstack.io/media/2021/01/2d3a7aa5-screen-shot-2021-01-24-at-5.52.28-pm.png)](https://landscape.cncf.io/)

//landscape.cncf.io/)

### What It Is

###  这是什么

Streaming and messaging tools enable service-to-service communication by transporting messages (i.e. events) between systems. Individual services connect to the messaging service to either publish events, read messages from other services, or both. This dynamic creates an environment where individual apps are either publishers, meaning they write events, or subscribers that read events, or more likely both.

流媒体和消息传递工具通过在系统之间传输消息（即事件）来实现服务到服务的通信。单个服务连接到消息服务以发布事件、从其他服务读取消息，或两者兼而有之。这种动态创造了一个环境，其中单个应用程序要么是发布者，这意味着它们编写事件，要么是读取事件的订阅者，或者更有可能两者兼而有之。

### Problem It Addresses

### 它解决的问题

As services proliferate, application environments become increasingly complex, making the orchestration of communication between apps more challenging. A streaming or messaging platform provides a central place to publish and read all the events that occur within a system, allowing applications to work together without necessarily knowing anything about one another.

随着服务的激增，应用程序环境变得越来越复杂，这使得应用程序之间的通信编排更具挑战性。流媒体或消息传递平台提供了一个中心位置来发布和读取系统内发生的所有事件，允许应用程序协同工作，而不必了解彼此。

### How It Helps

### 它如何帮助

When a service does something other services should know about, it “publishes” an event to the streaming or messaging tool. Services that need to know about these types of events subscribe and watch the streaming or messaging tool. That’s the essence of a publish-subscribe, or just pub-sub, approach and is enabled by these tools.

当服务执行其他服务应该知道的事情时，它会将事件“发布”到流媒体或消息传递工具。需要了解这些类型事件的服务订阅和观看流媒体或消息传递工具。这是发布-订阅或只是发布-订阅方法的本质，并且由这些工具启用。

By introducing a “go-between” layer that manages all communication, we are decoupling services from one another. They simply watch for events, take action, and publish a new one.

通过引入管理所有通信的“中间层”，我们将服务彼此解耦。他们只是观察事件、采取行动并发布新事件。

Here’s an example. When you first sign up for Netflix, the signup service publishes a “new signup event” to a messaging platform with further details (e.g. name, email address, subscription level, etc.). The account creator service, which subscribes to signup events, will see the event and create your account. A customer communication service that also subscribes to new signup events will add your email address to the customer mailing list and generate a welcome email, and so on.

这是一个例子。当您首次注册 Netflix 时，注册服务会向消息平台发布“新注册事件”，其中包含更多详细信息（例如姓名、电子邮件地址、订阅级别等）。订阅注册事件的帐户创建者服务将看到该事件并创建您的帐户。也订阅新注册事件的客户通信服务会将您的电子邮件地址添加到客户邮件列表并生成欢迎电子邮件等。

This allows for a highly decoupled architecture where services can collaborate without needing to know about one another. This decoupling enables engineers to add new functionality without updating downstream apps (consumers) or sending a bunch of queries. The more decoupled a system is, the more flexible and amenable to change. And that is exactly what engineers strive for in a system.

这允许高度解耦的架构，其中服务可以协作而无需相互了解。这种解耦使工程师能够在不更新下游应用程序（消费者）或发送大量查询的情况下添加新功能。系统越解耦，就越灵活，更容易改变。而这正是工程师在系统中所追求的。

### Technical 101

### 技术 101

Messaging and streaming tools have been around long before cloud native became a thing. To centrally manage business-critical events, organizations have built large enterprise service buses. But when we talk about messaging and streaming in a cloud native context, we’re generally referring to tools like NATS, RabbitMQ, Kafka, or cloud provided message queues. 

消息和流媒体工具早在云原生成为一种东西之前就已经存在了。为了集中管理关键业务事件，组织构建了大型企业服务总线。但是，当我们在云原生上下文中谈论消息传递和流传输时，我们通常指的是 NATS、RabbitMQ、Kafka 或云提供的消息队列等工具。

What these systems have in common, is the architectural patterns they enable. Application interactions in a cloud native environment are either orchestrated or choreographed. You can find a lot more details about the topic in Sam Newton's book [Building Microservices](https://samnewman.io/books/building_microservices/), as well as brief summaries [in this StackOverflow post](https://stackoverflow.com/questions/4127241/orchestration-vs-choreography) and [this great Medium blog by Chen Chen](https://medium.com/ingeniouslysimple/choreography-vs-orchestration-a6f21cfaccae). To simplify things, let’s just say that orchestrated refers to systems that are centrally managed and choreographed are those that allow individual components to act independently.

这些系统的共同点是它们支持的架构模式。云原生环境中的应用程序交互要么是编排的，要么是编排的。您可以在 Sam Newton 的著作 [Building Microservices](https://samnewman.io/books/building_microservices/) 中找到有关该主题的更多详细信息，以及 [在这篇 StackOverflow 帖子中](https://stackoverflow.com/questions/4127241/orchestration-vs-choreography)和[陈晨的这个伟大的媒体博客](https://medium.com/ingeniouslysimple/choreography-vs-orchestration-a6f21cfaccae)。为了简化事情，我们只是说编排是指集中管理和编排的系统是那些允许单个组件独立运行的系统。

Messaging and streaming systems provide a central place for choreographed systems to communicate. The message bus provides a common place where all apps can go to, to both tell others what they’re doing by publishing messages, or see what things are going on by subscribing to messages.

消息传递和流媒体系统为精心设计的系统提供了一个通信的中心位置。消息总线提供了一个所有应用程序都可以访问的公共场所，既可以通过发布消息告诉其他人他们在做什么，也可以通过订阅消息来查看正在发生的事情。

The [NATS](https://landscape.cncf.io/?selected=nats) and [Cloudevents](https://landscape.cncf.io/?selected=cloud-events) projects are both incubating projects in this space with NATS providing a mature messaging system and Cloudevents being an effort to standardize message formats between systems. [Strimzi](https://landscape.cncf.io/?selected=strimzi),[Pravega](https://landscape.cncf.io/?selected=pravega), and [Tremor](https://landscape.cncf.io/?selected=tremor) are sandbox projects with each being tailored to a unique use case around streaming and messaging.

[NATS](https://landscape.cncf.io/?selected=nats) 和 [Cloudevents](https://landscape.cncf.io/?selected=cloud-events) 项目都是该领域的孵化项目NATS 提供了一个成熟的消息传递系统，而 Cloudevents 则致力于标准化系统之间的消息格式。 [Strimzi](https://landscape.cncf.io/?selected=strimzi)、[Pravega](https://landscape.cncf.io/?selected=pravega)、[Tremor](https://landscape.cncf.io/?selected=tremor) 是沙盒项目，每个项目都针对流媒体和消息传递的独特用例量身定制。

BuzzwordsPopular Projects

流行语热门项目

- Choreography
- Steaming
- MQ
- Message Bus

- 编舞
- 蒸
- MQ
- 消息总线

- Spark
- Kafka
- RabbitMQ
- Nats

- 火花
- 卡夫卡
- RabbitMQ
- 纳兹

## Application Definition and Image Build

## 应用定义和镜像构建

[![app definition and image build](https://cdn.thenewstack.io/media/2021/01/1165999f-screen-shot-2021-01-24-at-5.54.25-pm.png)](https://landscape.cncf.io/)

https://landscape.cncf.io/)

### What It Is

###  这是什么

App definition and image build is a broad category that can be broken down into two main subgroups. First, dev-focused tools that help build application code into containers and/or Kubernetes. And second, ops-focused tools that deploy apps in a standardized way. Whether to speed up or simplify your development environment, provide a standardized way to deploy third-party apps, or simplify the process of writing a new Kubernetes extension, this category serves as a catch-all for a number of projects and products that optimize the Kubernetes developer and operator experience.

应用程序定义和映像构建是一个广泛的类别，可以分为两个主要子组。首先，以开发为中心的工具帮助将应用程序代码构建到容器和/或 Kubernetes 中。其次，以标准化方式部署应用程序的以运营为中心的工具。无论是加速或简化您的开发环境，提供一种标准化的方式来部署第三方应用程序，还是简化编写新的 Kubernetes 扩展的过程，该类别是许多优化应用程序的项目和产品的统称。 Kubernetes 开发人员和操作员经验。

### Problem It Addresses

### 它解决的问题

Kubernetes, and containerized environments more generally, are incredibly flexible and powerful. With that flexibility also comes complexity, mainly in form of numerous configuration options for various often new use cases. Developers must containerize their code and have the ability to develop in production-like environments. And with a rapid release schedule, operators need a standardized way to deploy apps into container environments.

Kubernetes 和更普遍的容器化环境非常灵活和强大。这种灵活性也带来了复杂性，主要表现为各种新用例的众多配置选项。开发人员必须容器化他们的代码，并能够在类似生产的环境中进行开发。由于发布计划很快，运营商需要一种标准化的方式将应用程序部署到容器环境中。

### How It Helps

### 它如何帮助

Tools in this space aim at solving some of these developer or operator challenges. On the developer side, there are tools that simplify the process of extending Kubernetes to build, deploy, and connect applications. A number of projects and products help to store or deploy pre-packaged apps. These allow operators to quickly deploy a streaming service like Kafka or install a service mesh like Linkerd.

该领域的工具旨在解决其中一些开发人员或运营商的挑战。在开发人员方面，有一些工具可以简化扩展 Kubernetes 以构建、部署和连接应用程序的过程。许多项目和产品有助于存储或部署预先打包的应用程序。这些允许运营商快速部署像 Kafka 这样的流媒体服务或安装像 Linkerd 这样的服务网格。

Developing cloud native applications brings a whole new set of challenges calling for a large set of diverse tools to simplify application build and deployments. As you start addressing operational and developer concerns in your environment, look for tools in this category.

开发云原生应用程序带来了一系列全新的挑战，需要大量不同的工具来简化应用程序构建和部署。当您开始解决环境中的运营和开发人员问题时，请寻找此类工具。

### Technical 101 

### 技术 101

App definition and build tools encompass a huge range of functionality. From extending [Kubernetes](https://landscape.cncf.io/?selected=kubernetes) to virtual machines with [KubeVirt](https://landscape.cncf.io/?selected=kube-virt), to speeding app development by allowing you to port your development environment into Kubernetes with tools like Telepresence, and everything in between. At a high level, tools in this space solve either developer-focused concerns, like how to correctly write, package, test, or run custom apps, or operations-focused concerns, such as deploying and managing applications.

应用程序定义和构建工具包含大量功能。从扩展 [Kubernetes](https://landscape.cncf.io/?selected=kubernetes) 到使用 [KubeVirt](https://landscape.cncf.io/?selected=kube-virt) 的虚拟机，到加速应用通过允许您使用 Telepresence 等工具将您的开发环境移植到 Kubernetes 中，以及介于两者之间的一切。在高层次上，该领域的工具解决了以开发人员为中心的问题，例如如何正确编写、打包、测试或运行自定义应用程序，或者以操作为中心的问题，例如部署和管理应用程序。

Helm, the only graduated project in this category, underpins many app deployment patterns. Allowing Kubernetes users to deploy and customize many popular third-party apps, it has been adopted by other projects like the [Artifact Hub](https://landscape.cncf.io/?selected=artifact-hub), a CNCF sandbox project , and [Bitnami](https://landscape.cncf.io/?selected=bitnami-tanzu-application-catalog) to provide curated catalogs of apps. [Helm](https://landscape.cncf.io/?selected=helm) is also flexible enough to allow users to customize their own app deployments.

Helm 是该类别中唯一的毕业项目，它支持许多应用程序部署模式。允许 Kubernetes 用户部署和定制许多流行的第三方应用程序，它已被其他项目采用，例如 [Artifact Hub](https://landscape.cncf.io/?selected=artifact-hub)，一个 CNCF 沙箱项目, 和 [Bitnami](https://landscape.cncf.io/?selected=bitnami-tanzu-application-catalog) 提供精选的应用程序目录。 [Helm](https://landscape.cncf.io/?selected=helm) 也足够灵活，允许用户自定义他们自己的应用程序部署。

The [Operator Framework](https://landscape.cncf.io/?selected=operator-framework) is an incubating project aimed at simplifying the process of building and deploying operators. Operators are out of scope for this article but let's note here that they help deploy and manage apps, similar to Helm (you can read more about operators [here](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)). [Cloud Native Buildpacks](https://landscape.cncf.io/?selected=buildpacks), another incubating project, aims to simplify the process of building application code into containers.

[Operator Framework](https://landscape.cncf.io/?selected=operator-framework)是一个孵化项目，旨在简化构建和部署Operator的过程。 Operators 超出了本文的范围，但让我们在此注意，它们帮助部署和管理应用程序，类似于 Helm（您可以在 [此处] 阅读更多关于 Operators 的信息（https://kubernetes.io/docs/concepts/extend-kubernetes/操作员/）)。 [Cloud Native Buildpacks](https://landscape.cncf.io/?selected=buildpacks)，另一个孵化项目，旨在简化将应用程序代码构建到容器中的过程。

There’s a lot more in this space and exploring it all would require a dedicated article. But research these tools further if you want to make Kubernetes easier for developers and operators. You’ll likely find something that meets your needs.

这个领域还有很多，探索这一切需要一篇专门的文章。但是，如果您想让开发人员和运营商更轻松地使用 Kubernetes，请进一步研究这些工具。您可能会找到满足您需求的东西。

BuzzwordsPopular Projects

流行语热门项目

- Package Management
- Charts
- Operators

- 包管理
- 图表
- 运营商

- Helm
- Buildpacks
- Tilt
- Okteto

- 头盔
- 构建包
- 倾斜
- 奥克泰托

## Continuous Integration and Continuous Delivery

## 持续集成和持续交付

[![CICD](https://cdn.thenewstack.io/media/2021/01/836e47cf-screen-shot-2021-01-24-at-5.55.02-pm.png)](https://landscape.cncf.io/)

Landscape.cncf.io/)

### What It Is

###  这是什么

Continuous integration (CI) and continuous delivery (CD) tools enable fast and efficient development with embedded quality assurance (we cover [CI/CD in detail in this primer](https://thenewstack.io/a-primer-continuous-integration-and-continuous-delivery-ci-cd/)). CI automates code changes by immediately building and testing the code, ensuring it produces a deployable artifact. CD goes one step further and pushes the artifact through the deployment phases.

持续集成 (CI) 和持续交付 (CD) 工具通过嵌入式质量保证实现快速高效的开发（我们在本入门手册中详细介绍了 [CI/CD](https://thenewstack.io/a-primer-continuous-integration-and-continuous-delivery-ci-cd/))。 CI 通过立即构建和测试代码来自动化代码更改，确保它生成可部署的工件。 CD 更进一步，推动工件通过部署阶段。

Mature CI/CD systems watch source code for changes, automatically build and test the code, then begin moving it from development to production where it has to pass a variety of tests or validation to determine if the process should continue or fail. Tools in this category enable such an approach.

成熟的 CI/CD 系统观察源代码的变化，自动构建和测试代码，然后开始将其从开发转移到生产，在那里它必须通过各种测试或验证，以确定过程是否应该继续或失败。此类别中的工具支持这种方法。

### Problem It Addresses

### 它解决的问题

Building and deploying applications is a difficult and error-prone process. Particularly, when it involves a lot of human intervention and manual steps. The longer a developer works on a piece of software without integrating it into the codebase, the longer it will take to identify an error and the more difficult it is to fix. By integrating code on a regular basis, errors are caught early and are easier to troubleshoot. After all, finding an error in a few lines of code is a lot easier than doing so in a few hundred lines of code. 

构建和部署应用程序是一个困难且容易出错的过程。特别是当它涉及大量人工干预和手动步骤时。开发人员在未将其集成到代码库中的情况下处理软件的时间越长，识别错误所需的时间就越长，修复起来也就越困难。通过定期集成代码，可以及早发现错误并更容易进行故障排除。毕竟，在几行代码中查找错误比在几百行代码中查找错误要容易得多。

While tools like Kubernetes offer great flexibility for running and managing apps, they also create new challenges and opportunities for CI/CD tooling. Cloud native CI/CD systems are able to leverage Kubernetes itself to build, run, and manage the CI/CD process, often referred to as [pipelines](https://www.redhat.com/en/topics/devops/what-cicd-pipeline). Kubernetes also provides information about the health of our apps enabling cloud native CI/CD tools to more easily determine if a given change was successful or needs to be rolled back.

虽然像 Kubernetes 这样的工具为运行和管理应用程序提供了极大的灵活性，但它们也为 CI/CD 工具带来了新的挑战和机遇。云原生 CI/CD 系统能够利用 Kubernetes 本身来构建、运行和管理 CI/CD 过程，通常称为 [管道](https://www.redhat.com/en/topics/devops/what-cicd-管道)。 Kubernetes 还提供有关我们应用程序运行状况的信息，使云原生 CI/CD 工具能够更轻松地确定给定的更改是成功还是需要回滚。

### How It Helps

### 它如何帮助

CI tools ensure that any code change or updates developers introduce are built, validated, and integrated with other changes automatically and continuously. Each time a developer adds an update, automated testing is triggered ensuring only good code makes it into the system. CD extends CI to include pushing the result of the CI process into production-like and production environments.

CI 工具可确保开发人员引入的任何代码更改或更新都自动且持续地构建、验证并与其他更改集成。每次开发人员添加更新时，都会触发自动化测试，确保只有好的代码才能进入系统。 CD 扩展了 CI 以包括将 CI 过程的结果推送到类似生产的生产环境中。

Let’s say a developer changes the code for a web app. The CI system sees the code change, and then builds and tests a new version of that web app. The CD system takes that new version and deploys it into a dev, test, pre-production, and finally production environment. It does that while testing the deployed app after each step in the process. All together these systems represent a CI/CD pipeline for that web app.

假设开发人员更改了 Web 应用程序的代码。 CI 系统会看到代码更改，然后构建并测试该 Web 应用程序的新版本。 CD 系统采用该新版本并将其部署到开发、测试、预生产和最终生产环境中。它会在流程中的每个步骤之后测试部署的应用程序时这样做。所有这些系统都代表了该 Web 应用程序的 CI/CD 管道。

### Technical 101

### 技术 101

Over time, a number of tools have been built to help with the process of moving code from a repository, where the code is stored, to production, where the finished app runs. Like most other areas of computing, the advent of cloud native development has changed CI/CD systems. Some traditional tools like [Jenkins](https://landscape.cncf.io/?selected=jenkins), probably the most widely-used CI tool on the market, has been [overhauled](https://jenkins-x.io/) entirely to better fit into the Kubernetes ecosystem. Others, like [Flux](https://landscape.cncf.io/?selected=flux) and [Argo](https://landscape.cncf.io/?selected=argo) have pioneered a new way of doing continuous delivery called GitOps.

随着时间的推移，已经构建了许多工具来帮助将代码从存储代码的存储库移动到生产完成的应用程序运行的过程。与大多数其他计算领域一样，云原生开发的出现改变了 CI/CD 系统。一些传统工具，如 [Jenkins](https://landscape.cncf.io/?selected=jenkins)，可能是市场上使用最广泛的 CI 工具，已经[大修](https://jenkins-x.io/) 完全是为了更好地融入 Kubernetes 生态系统。其他人，如 [Flux](https://landscape.cncf.io/?selected=flux) 和 [Argo](https://landscape.cncf.io/?selected=argo) 开创了一种新的连续交付称为 GitOps。

In general, you'll find projects and products in this space are either (1) CI systems, (2) CD systems, (3) tools that help the CD system decide if the code is ready to be pushed into production, or ( 4), in the case of [Spinnaker](https://landscape.cncf.io/?selected=spinnaker) and Argo, all three. Argo and Brigade are the only CNCF projects in this space but you can find many more options hosted by the [Continuous Delivery Foundation](https://cd.foundation/). Look for tools in this space to help your organization automate your path to production.

通常，您会发现该领域中的项目和产品是 (1) CI 系统，(2) CD 系统，(3) 帮助 CD 系统确定代码是否准备好投入生产的工具，或者 ( 4），在[Spinnaker](https://landscape.cncf.io/?selected=spinnaker)和Argo的情况下，三个。 Argo 和 Brigade 是该领域仅有的 CNCF 项目，但您可以找到更多由 [Continuous Delivery Foundation] (https://cd.foundation/) 托管的选项。在这个领域寻找工具来帮助您的组织自动化您的生产路径。

BuzzwordsPopular Projects

流行语热门项目

- CI/CD
- Continuous integration
- Continuous delivery
- Blue/Green
- Canary Deploy

- CI/CD
- 持续集成
- 持续交付
- 蓝绿
- 金丝雀部署

- Argo
- Flagger
- Spinnaker
- Jenkins

- 阿尔戈
- 旗手
- 三角帆
- 詹金斯

## Conclusion

##  结论

As we’ve seen, tools in the application definition and development layer enable engineers to build cloud native apps. You’ll find databases to store and retrieve data or streaming and messaging tools allowing for decoupled, choreographed architectures. App definition and image build tools encompass a variety of technologies that improve the developer and operator experience, and CI/CD ensures code is at a deployable state and helps engineers catch any errors early on, ensuring better code quality.

正如我们所见，应用程序定义和开发层中的工具使工程师能够构建云原生应用程序。您将找到用于存储和检索数据的数据库或允许解耦、编排架构的流媒体和消息传递工具。应用程序定义和映像构建工具包含多种技术，可改善开发人员和操作人员的体验，而 CI/CD 可确保代码处于可部署状态并帮助工程师及早发现任何错误，从而确保更好的代码质量。

This article concludes the layers of the CNCF landscape. In our next article, we’ll focus on cloud native platforms, the first column running across all these layers. Configuring tools across the layers we discussed so far so they work well together is no easy task. Platforms bundle them together easing adoption, but more to that in our next article.

本文总结了 CNCF 格局的各个层次。在我们的下一篇文章中，我们将重点关注云原生平台，这是贯穿所有这些层的第一列。到目前为止，我们讨论的跨层配置工具使它们能够很好地协同工作并不是一件容易的事。平台将它们捆绑在一起以简化采用，但更多内容将在我们的下一篇文章中介绍。

_As always, a very special thanks to [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/) from the CNCF who was so kind as to review the article making sure it’s all accurate._

_一如既往，非常感谢来自 CNCF 的 [Ihor Dvoretskyi](https://www.linkedin.com/in/idvoretskyi/)，他非常友好地审阅了这篇文章，以确保其准确无误。_

The Cloud Native Computing Foundation is a sponsor of The New Stack.

云原生计算基金会是 The New Stack 的赞助商。

The New Stack is a wholly owned subsidiary of Insight Partners. TNS owner Insight Partners is an investor in the following companies: Bit. 

The New Stack 是 Insight Partners 的全资子公司。 TNS 所有者 Insight Partners 是以下公司的投资者：Bit。

