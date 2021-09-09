# How to familiarize yourself with a new codebase

# 如何熟悉新的代码库

A few weeks ago, a tweet made me take a second and think about something that I'd never consciously considered before; how can you approach an unfamiliar codebase and start to understand it?

几周前，一条推文让我花点时间思考一些我以前从未有意识地考虑过的事情；你如何接近一个不熟悉的代码库并开始理解它？

[https://twitter.com/d\_feldman/status/1336407539928477697?s=21](https://twitter.com/d_feldman/status/1336407539928477697?s=21)

[https://twitter.com/d\_feldman/status/1336407539928477697?s=21](https://twitter.com/d_feldman/status/1336407539928477697?s=21)

It got me thinking about how I would approach a new repo that I'd never seen before but needed to make a contribution against, like a bug fix. I remembered my early days of learning [Kubernetes](https://kubernetes.io/), and wanting to make requests to the its API (because using the command line wasn't good enough for me, apparently). I had been trying to work out how to automatically deploy a particular branch of a GitLab repo into a cluster every time someone pushed to it. I had big ideas about automating DNS, setting up automated certificates, and adding a Slackbot to notify you whenever a new deploy happened.

这让我开始思考如何处理一个我以前从未见过但需要做出贡献的新仓库，比如修复错误。我记得我早期学习 [Kubernetes](https://kubernetes.io/)，并想向其 API 发出请求（因为显然使用命令行对我来说不够好)。我一直在尝试研究如何在每次有人推送时自动将 GitLab 存储库的特定分支部署到集群中。我对自动化 DNS、设置自动证书以及添加 Slackbot 以在发生新部署时通知您有很大的想法。

If I remember correctly, I got a proof of concept working, and then it never went much past that. Given how popular [GitOps](https://www.cloudbees.com/gitops/what-is-gitops) has become, maybe I should have stuck with it! When I started delving into the Kuberenetes side of the project, I was completely and utterly lost. The documentation didn't have much in the way of _how_ to use the API (I'm sure nowadays things are much better), and reading the Kubernetes source code was a complete non-starter because well, that thing is a monster. I remember thinking to myself that I just needed to replicate what kubectl was doing to create a new Deployment.

如果我没记错的话，我得到了一个概念证明工作，然后它再也没有超过那个。鉴于 [GitOps](https://www.cloudbees.com/gitops/what-is-gitops) 的流行程度，也许我应该坚持下去！当我开始深入研究项目的 Kuberenetes 方面时，我完全迷失了方向。文档没有太多关于 _how_ 使用 API 的方式（我相信现在情况好多了)，阅读 Kubernetes 源代码完全不是初学者，因为好吧，那东西是一个怪物。我记得我当时在想，我只需要复制 kubectl 正在做的事情来创建一个新的 Deployment。

So I gave up trying to read Kubernetes' source, and moved over to the [source for kubectl](https://github.com/kubernetes/kubectl). This is where I started to make some headway! I was able to follow straight from the `main()` function to the `apply` command, down through the logic until it started making API requests. It felt so good to finally get an answer, and to just import some Go packages to make it all work in short order!\*

所以我放弃了阅读 Kubernetes 的源代码，转而使用 [kubectl 的源代码](https://github.com/kubernetes/kubectl)。这是我开始取得进展的地方！我能够直接从`main()` 函数到 `apply` 命令，一直到逻辑直到它开始发出 API 请求。终于得到答案的感觉真是太好了，只需导入一些 Go 包即可在短时间内完成所有工作！\*

This is the background behind my answer to the tweet above:

这是我对上述推文的回答背后的背景：

[https://twitter.com/cohix/status/1336408360770531331?s=21](https://twitter.com/cohix/status/1336408360770531331?s=21)

[https://twitter.com/cohix/status/1336408360770531331?s=21](https://twitter.com/cohix/status/1336408360770531331?s=21)

Since that project years ago, I've sort of instinctively followed this strategy whenever I need to reason about a new codebase because well, it works! Only recently did this tweet make me think about it concretely, and I'm glad it did. I tried to replicate this purposefully to test my strategy. I went to a [large open-source repo](https://github.com/fluxcd/flux2) and tried to find the code where it installed itself into a cluster. Using this strategy, I started with the tool's `main()` and then was able to find my way to the `install` command, which led me down to where the installation happens (funnily enough, by calling `kubectl`).

从几年前的那个项目开始，每当我需要推理一个新的代码库时，我都会本能地遵循这个策略，因为它确实有效！直到最近这条推文才让我具体思考它，我很高兴它做到了。我试图有目的地复制这个来测试我的策略。我去了一个 [大型开源存储库](https://github.com/fluxcd/flux2) 并试图找到它安装到集群中的代码。使用这种策略，我从工具的 `main()` 开始，然后找到了使用 `install` 命令的方法，这将我引导到安装发生的位置（有趣的是，通过调用 `kubectl`)。

I think it's important for any developer to understand not only how to reason about an unfamiliar codebase, but also to realize that an important way that we learn is by trial and error. When we try something and it works, it brings us joy and we'll continue to do it, even if we don't realize it. I think it's a good idea to take a second to think about these moments when they happen, take a mental note of it so that next time you come across a similar problem, you can consciously use your previous learning and expand upon the strategies you' ve developed over time. 

我认为对于任何开发人员来说，重要的是不仅要了解如何对不熟悉的代码库进行推理，还要意识到我们学习的一种重要方式是反复试验。当我们尝试某件事并且它奏效时，它会给我们带来快乐，我们会继续这样做，即使我们没有意识到这一点。我认为在这些时刻发生时花点时间思考是个好主意，记下它，以便下次遇到类似问题时，您可以有意识地使用以前的知识并扩展您的策略随着时间的推移已经发展。

The reason I wanted to turn this tweet into a full blog post is because it made me realize that one of the goals for [Atmo](https://github.com/suborbital/atmo) is to make it easier for developers to reason about applications. Since Atmo uses a declarative format for building backend applications (using WebAssembly modules), there is always one canonical entrypoint; the Directive. From there, you can easily reason about what the application is doing because it is [declarative instead of imperative](https://stackoverflow.com/a/1784702). This is one of the things that made Kubernetes so popular. Being able to describe your application in a simple format, and then have a system "make it happen" is a magical thing, and Atmo strives to do exactly that.

我想把这条推文变成一篇完整的博文是因为它让我意识到 [Atmo](https://github.com/suborbital/atmo) 的目标之一是让开发人员更容易推理关于应用程序。由于 Atmo 使用声明式格式来构建后端应用程序（使用 WebAssembly 模块)，因此总是有一个规范的入口点；指令。从那里，您可以轻松推断应用程序正在做什么，因为它是 [声明式而不是命令式](https://stackoverflow.com/a/1784702)。这是使 Kubernetes 如此受欢迎的原因之一。能够以简单的格式描述您的应用程序，然后让系统“让它发生”是一件神奇的事情，而 Atmo 正努力做到这一点。

Atmo is gaining new functionality every week. If you want to learn more, check out [the Suborbital homepage](https://suborbital.dev)

Atmo 每周都在获得新功能。如果您想了解更多信息，请查看 [亚轨道主页](https://suborbital.dev)

- When I say "short order", I'm sure it still took several days to get everything working, but once you unblock yourself on a big problem, everything after that just seems to fly by.

- 当我说“短期订单”时，我确信一切都还需要几天时间才能使一切正常进行，但是一旦您解决了一个大问题，之后的一切似乎都会过去。

Cover Photo by [Rafif Prawira](https://unsplash.com/@rafifatmaka) on [Unsplash](https://unsplash.com/s/photos/maze) 

封面照片由 [Rafif Prawira](https://unsplash.com/@rafifatmaka) 在 [Unsplash](https://unsplash.com/s/photos/maze)

