# Why Kubernetes Operators Will Unleash Your Developers by Reducing Complexity

# 为什么 Kubernetes Operator 会通过降低复杂性来释放您的开发人员

#### 17 Aug 2020 8:56am,   by [Rob Szumski](https://thenewstack.io/author/rob-szumski/ "Posts by Rob Szumski")

#### 2020 年 8 月 17 日上午 8:56，作者 [Rob Szumski](https://thenewstack.io/author/rob-szumski/“Rob Szumski 的帖子”)

Rob is a Product Manager for OpenShift at Red Hat](https://www.linkedin.com/in/robszumski/)

Rob 是 Red Hat 的 OpenShift 产品经理](https://www.linkedin.com/in/robszumski/)

Kubernetes and the many projects under the Cloud Native Computing Foundation (CNCF) umbrella are advancing so quickly that it can sometimes feel as though the benefits of the open hybrid cloud are, perhaps, not yet evenly distributed. While systems administrators and operators are having a blast modernizing legacy applications and automating highly-scalable, container-based systems, sometimes developers can feel a bit left out in the cold.

Kubernetes 和云原生计算基金会 (CNCF) 保护伞下的许多项目进展如此之快，以至于有时感觉开放混合云的好处可能尚未平均分配。虽然系统管理员和操作员正在对遗留应用程序进行现代化改造并自动化高度可扩展的基于容器的系统，但有时开发人员可能会感到有点被冷落。

That’s because there’s a disconnect between the velocity increase that IT is experiencing through the adoption of containers, and the agility increase offered to developers. The benefits that Kubernetes offers to IT operators are in service of simplifying the lives of their developers. Developers and IT administrators want different things, right? Or do they?

这是因为 IT 通过采用容器实现的速度提升与为开发人员提供的敏捷性提升之间存在脱节。 Kubernetes 为 IT 运营商提供的好处是为了简化开发人员的生活。开发人员和 IT 管理员想要不同的东西，对吧？还是他们？

Back in 2016 when I was at CoreOS, we saw this disconnect beginning to form. We worried that while enabling developers was job one for IT, it wasn’t necessarily easy when developers had to relearn their entire stacks from the ground up in order to adopt containers. It was also tough for IT administrators to provision compliant, governed services at the pace expected in cloud development.

早在 2016 年我在 CoreOS 时，我们就看到这种脱节开始形成。我们担心，虽然支持开发人员是 IT 的一项工作，但当开发人员必须从头开始重新学习他们的整个堆栈以采用容器时，这并不一定容易。 IT 管理员也很难以云开发中预期的速度提供合规、受监管的服务。

To fix this, CoreOS introduced the concept of Kubernetes Operators. Soon after, we produced a set of tools known as the Operator Framework to help users build, ship and discover their own Kubernetes Operators. While a lot has changed for myself and CoreOS since 2016, most notably the fact that we are now both a part of the Red Hat family, Operators remain an important part of our strategy for helping developers and administrators focus on using Kubernetes — rather than tweaking its various knobs and configurations.

为了解决这个问题，CoreOS 引入了 Kubernetes Operators 的概念。不久之后，我们开发了一组称为 Operator Framework 的工具，以帮助用户构建、发布和发现他们自己的 Kubernetes Operator。虽然自 2016 年以来我和 CoreOS 发生了很多变化，最显着的是我们现在都是 Red Hat 家族的一部分，但 Operators 仍然是我们帮助开发人员和管理员专注于使用 Kubernetes 而不是调整战略的重要组成部分它的各种旋钮和配置。

**Sponsor Note**

**赞助商备注**

![sponsor logo](https://cdn.thenewstack.io/media/2016/05/20f61f0a-red-hat-openshift@2x.png)

OpenShift is Red Hat’s container application platform that allows developers to quickly develop, host and scale applications in a cloud environment. With OpenShift, you have a choice of offerings, including online, on-premise and hosted service offerings.

OpenShift 是红帽的容器应用平台，允许开发人员在云环境中快速开发、托管和扩展应用。借助 OpenShift，您可以选择多种产品，包括在线、内部部署和托管服务产品。

Naturally, all of the work on Operators was done as open source, and as a result we recently contributed the Operators Framework to the CNCF, to become an incubated project. We did this for a few reasons, including that it was the right thing to do in this open ecosystem. But we also did it so that Operators will become a ubiquitous path to distributing servers into Kubernetes clusters. With Operators available to every Kubernetes distribution, the entire open hybrid cloud ecosystem will benefit from greater compatibility, simplicity and easier management.

自然而然，Operators 的所有工作都是开源完成的，因此我们最近将 Operators Framework 贡献给了 CNCF，成为一个孵化项目。我们这样做有几个原因，包括在这个开放的生态系统中这样做是正确的。但我们这样做也是为了让 Operator 成为将服务器分发到 Kubernetes 集群的普遍途径。随着 Operator 可用于每个 Kubernetes 发行版，整个开放式混合云生态系统将受益于更高的兼容性、简单性和更轻松的管理。

Operators are a big part of helping us reach this automated, on-demand, container-based future. Operators are operational procedures and best practices that are codified into software. They make automated day two operations possible on Kubernetes, and model the complexities of today’s distributed systems.

操作员是帮助我们实现自动化、按需、基于容器的未来的重要组成部分。操作员是编入软件的操作程序和最佳实践。它们使 Kubernetes 上的自动化第二天操作成为可能，并为当今分布式系统的复杂性建模。

For example, there isn’t a concept of data rebalancing in Kubernetes, but that can be built on top of the Kubernetes APIs with the Operator Framework. These types of applications are required for running what we call the third wave of applications on Kubernetes: complex distributed systems.

例如，Kubernetes 中没有数据重新平衡的概念，但它可以使用 Operator Framework 构建在 Kubernetes API 之上。在 Kubernetes 上运行我们称之为第三波应用程序的应用程序需要这些类型的应用程序：复杂的分布式系统。

[![](https://cdn.thenewstack.io/media/2020/08/ce7b9fcc-image3.png)](https://cdn.thenewstack.io/media/2020/08/ce7b9fcc-image3.png)

)

The Operator Framework comes with all the tools required for developers to build this software, and everything that cluster administrators need to safely install, upgrade and manage their Operators. These tools tie into other CNCF projects — like KubeBuilder, Helm, kuttl — and popular open source software like Ansible.

Operator Framework 附带开发人员构建此软件所需的所有工具，以及集群管理员安全安装、升级和管理其 Operator 所需的一切。这些工具与其他 CNCF 项目（如 KubeBuilder、Helm、kuttl）以及 Ansible 等流行的开源软件相关联。

The framework is loosely coupled, so if you have a favorite testing tool, you can keep using that; and if you want to build your Operator outside of the SDK, that’s fine too, you can still run it with the Operator Lifecycle Manager. 

该框架是松散耦合的，所以如果你有一个最喜欢的测试工具，你可以继续使用它；如果您想在 SDK 之外构建您的 Operator，那也没关系，您仍然可以使用 Operator Lifecycle Manager 运行它。

Operator Framework has three flavors of SDK today, with more to come in the future. Each one addresses a different type of author, from IT administrator to traditional developer to hardcore Kubernetes experts.

Operator Framework 目前拥有三种 SDK 版本，未来还会有更多版本。每一个都针对不同类型的作者，从 IT 管理员到传统开发人员，再到核心 Kubernetes 专家。

[![](https://cdn.thenewstack.io/media/2020/08/5484c998-image1.png)](https://cdn.thenewstack.io/media/2020/08/5484c998-image1.png)

)

## Help for Building

## 建筑帮助

Testing your Operator is critical, and that’s why I advise most Operator authors to utilize one of our SDKs, or at least join our [community](https://github.com/operator-framework/community-operators). Our experts in the community have modeled many types of applications as Operators, so we can help you save time and avoid some bugs. Plus, it’s been really fun to see all the projects coming into our community and to help people meet their goals.

测试您的 Operator 至关重要，这就是为什么我建议大多数 Operator 作者使用我们的 SDK 之一，或者至少加入我们的 [社区](https://github.com/operator-framework/community-operators)。我们的社区专家已将多种类型的应用程序建模为 Operators，因此我们可以帮助您节省时间并避免一些错误。另外，看到所有项目进入我们的社区并帮助人们实现他们的目标真的很有趣。

Once you’ve written your Operator, you need to hand it off to your users and actually run it. This is where the Operator Lifecycle Manager comes into play. There are actually a lot of tricky problems here that you might not even know you have yet, like collision detection for CRDs. Imagine you have a database managed by a Custom Resource Definition (CRD), but there’s another Operator that also wants to manage that database. That’s no good. The lifecycle of a CRD itself is also important. The Operator can manage the CRD as part of the upgrade process.

一旦你编写了你的 Operator，你需要将它交给你的用户并实际运行它。这就是操作员生命周期管理器发挥作用的地方。这里实际上有很多棘手的问题，您甚至可能还不知道，例如 CRD 的碰撞检测。假设您有一个由自定义资源定义 (CRD) 管理的数据库，但还有另一个 Operator 也想管理该数据库。那不好。 CRD 本身的生命周期也很重要。作为升级过程的一部分，操作员可以管理 CRD。

We’re deeply committed to Operators, and we have a lot planned for the future of these powerful abstractions. For example, we’re currently working on the concept of Bundles, which will allow Operators to be cataloged together inside clusters. Administrators will be able to use a new tool we’re working on called OPM, which would allow them to better curate those in-cluster catalogs. We’re also working on a new Operator API, which is designed to provide an easier way to access Operators through Kubernetes APIs.

我们深深地致力于 Operators，并且我们为这些强大抽象的未来做了很多计划。例如，我们目前正在研究 Bundles 的概念，这将允许 Operator 在集群内一起编目。管理员将能够使用我们正在开发的名为 OPM 的新工具，这将使他们能够更好地管理这些集群内目录。我们还在开发新的 Operator API，旨在提供一种更简单的方式来通过 Kubernetes API 访问 Operator。

[![](https://cdn.thenewstack.io/media/2020/08/39abe664-image2.png)](https://cdn.thenewstack.io/media/2020/08/39abe664-image2.png)

)

While much of the workaround setting up an Operator inside a cluster will be done by the IT staff, it is the developer who really benefits. For example, IT can use the [Crunchy Data PostgreSQL Operator](https://operatorhub.io/operator/postgresql) to pre-configure database RBAC and security controls, while also pre-wiring backup and replication services. After this is done, any blessed user of the cluster can then deploy a fresh instance of PostgreSQL on-demand, eliminating the time required to provision servers for new application development work.

虽然在集群内设置 Operator 的大部分解决方法将由 IT 人员完成，但真正受益的是开发人员。例如，IT 可以使用 [Crunchy Data PostgreSQL Operator](https://operatorhub.io/operator/postgresql) 来预配置数据库 RBAC 和安全控制，同时还可以预接线备份和复制服务。完成此操作后，集群的任何受祝福的用户都可以按需部署新的 PostgreSQL 实例，从而消除为新应用程序开发工作配置服务器所需的时间。

This benefit also extends to automated on-demand provisioning of test and build environments. These environments can be pre-configured to adhere to corporate governance and data control policies through Operators properly configured by IT administrators. The best part is that they only have to do this once, and the deployment according to these rules is automated across the cluster.

这种优势还扩展到测试和构建环境的自动化按需供应。这些环境可以通过 IT 管理员正确配置的 Operator 进行预配置，以遵守公司治理和数据控制策略。最好的部分是他们只需要做一次，并且根据这些规则的部署是跨集群自动化的。

Operators have been available on OpenShift since version four was released last summer. We’ve been working hard with many developers across the enterprise and open source landscapes to help flesh out [Operatorhub.io](https://operatorhub.io/), our public repository for Kubernetes Operators. So the Operator ecosystem is already quite vibrant and ready for exploration.

自去年夏天发布第 4 版以来，OpenShift 上就可以使用 Operator。我们一直在与企业和开源领域的许多开发人员一起努力，以帮助充实 [Operatorhub.io](https://operatorhub.io/)，这是我们的 Kubernetes Operators 公共存储库。因此，Operator 生态系统已经非常活跃并准备好进行探索。

We’re excited to see how much larger this community can grow, now that it is fully a part of the CNCF’s open source processes. This was always what we intended to do with the Operator Framework, and now after four years of hard work, we’re very proud to see the project take the next step and run free as an open source project inside the CNCF. With your help, Operators will save everyone in the Kubernetes community — both administers and developers — a lot of time and worry.

我们很高兴看到这个社区可以发展到多大，现在它完全是 CNCF 开源流程的一部分。这一直是我们打算用 Operator Framework 做的事情，现在经过四年的努力，我们很自豪地看到该项目迈出了下一步，并作为 CNCF 内部的开源项目免费运行。在您的帮助下，Operators 将为 Kubernetes 社区中的每个人（包括管理员和开发人员）节省大量时间和精力。

The Cloud Native Computing Foundation is a sponsor of The New Stack.

云原生计算基金会是 The New Stack 的赞助商。

Feature image via [Pixabay](https://pixabay.com/illustrations/complex-fractal-chaos-grid-clock-664440/). 

特征图片来自 [Pixabay](https://pixabay.com/illustrations/complex-fractal-chaos-grid-clock-664440/)。

At this time, The New Stack does not allow comments directly on this website. We invite all readers who wish to discuss a story to visit us on [Twitter](https://twitter.com/thenewstack) or [Facebook](https://www.facebook.com/thenewstack/). We also welcome your news tips and feedback via email: [feedback@thenewstack.io](mailto:feedback@thenewstack.io). 

目前，The New Stack 不允许直接在本网站上发表评论。我们邀请所有希望讨论故事的读者在 [Twitter](https://twitter.com/thenewstack) 或 [Facebook](https://www.facebook.com/thenewstack/)上访问我们。我们也欢迎您通过电子邮件提供新闻提示和反馈：[feedback@thenewstack.io](mailto:feedback@thenewstack.io)。

