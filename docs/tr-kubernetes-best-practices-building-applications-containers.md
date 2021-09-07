# 7 best practices: Building applications for containers and Kubernetes

# 7 最佳实践：为容器和 Kubernetes 构建应用程序

Let’s examine key considerations for  building new applications specifically for containers and Kubernetes,  according to cloud-native experts

根据云原生专家的说法，让我们研究一下专门为容器和 Kubernetes 构建新应用程序的关键考虑因素

March 2, 2020



Don't let the growing popularity of [containers](https://enterprisersproject.com/tags/containers) and [Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes?intcmp=701f2000000tjyaaaa&extidcarryover=true&sc_cid=70160000000h0axaaq) dupe you into thinking that you should use them to run any and every type of application. You need to distinguish between “can” and “should.”

不要让[容器](https://enterprisersproject.com/tags/containers)和[Kubernetes](https://www.redhat.com/en/topics/containers/what-is-kubernetes)的日益流行欺骗您认为您应该使用它们来运行任何类型的应用程序。你需要区分“可以”和“应该”。

One basic example of this distinction is the difference between  building an app specifically to be run in containers and operated with  Kubernetes (some would refer to this as [cloud-native](https://www.redhat.com/en/topics/cloud-native-apps?intcmp=701f2000000tjyaAAA) development) and using these containers and orchestration for existing monolithic apps.

这种区别的一个基本示例是构建专门在容器中运行和使用 Kubernetes 操作的应用程序之间的区别（有些人将其称为 [云原生](https://www.redhat.com/en/topics/cloud-native-apps?intcmp=701f2000000tjyaAAA) 开发)并将这些容器和编排用于现有的单体应用程序。

Building  new applications specifically for containers and Kubernetes might be the best starting point for teams just beginning their container work.

对于刚刚开始容器工作的团队来说，专门为容器和 Kubernetes 构建新的应用程序可能是最好的起点。

We’ll cover the latter scenario in an upcoming post. Today, we’re  focused on some of the key considerations for building new applications  specifically for containers and Kubernetes – in part because the  so-called “greenfield” approach might be the better starting point for  teams just beginning with containers and orchestration.

我们将在下一篇文章中介绍后一种情况。今天，我们专注于为容器和 Kubernetes 构建新应用程序的一些关键考虑因素——部分原因是所谓的“绿地”方法可能是刚开始使用容器和编排的团队的更好起点。

“Containers [and orchestration] are a technical vehicle for building, deploying, and running cloud-native applications,” says Rani Osnat, VP  of strategy at [Aqua Security](https://www.aquasec.com/). “I typically recommend to those starting their journey with containers  to use a new, simple greenfield application as their test case.”

“容器 [和编排] 是构建、部署和运行云原生应用程序的技术工具，”Aqua Security战略副总裁 Rani Osnat 说。 “我通常建议那些开始使用容器的人使用一个新的、简单的绿地应用程序作为他们的测试用例。”

**[ Want to learn about building and deploying Kubernetes Operators? Get the free eBook: [O'Reilly: Kubernetes Operators: Automating the Container Orchestration Platform.](https://www.redhat.com/en/resources/oreilly-kubernetes-operators-automation-ebook?intcmp=701f2000000tjyaAAA) ] **

**[ 想了解如何构建和部署 Kubernetes Operator？获取免费电子书：[O'Reilly: Kubernetes Operators: Automating the Container Orchestration Platform.](https://www.redhat.com/en/resources/oreilly-kubernetes-operators-automation-ebook?intcmp=701f2000000tjyaAAA)] **

## **How to develop apps for containers and Kubernetes**

## **如何为容器和 Kubernetes 开发应用程序**

We asked Osnat and other cloud-native experts to share their top tips for developing apps specifically to be run in containers using  Kubernetes. Let’s dive into six of their best recommendations.

我们请 Osnat 和其他云原生专家分享他们使用 Kubernetes 开发专门在容器中运行的应用程序的重要技巧。让我们深入研究他们的六个最佳建议。

### 1. Think and build modern

### 1. 思考和构建现代

If you’re building a new home today, you’ve got different styles and  approaches than you would have, say, 50 years ago. The same is true with software: You’ve got new tools and approaches at your disposal.

如果您今天正在建造一个新家，那么您将拥有与 50 年前不同的风格和方法。软件也是如此：您可以使用新的工具和方法。

“If you’re building an app, build it in a modern way!” says Miles Ward, CTO at [SADA](https://sada.com/). Ward points to [microservices](https://enterprisersproject.com/article/2017/8/how-explain-microservices-plain-english) and the [12-factor methodology](https://12factor.net/) as chief examples of modern application development.

“如果您正在构建应用程序，请以现代方式构建它！” [SADA](https://sada.com/) 的首席技术官 Miles Ward 说。 Ward 指出 [微服务](https://enterprisersproject.com/article/2017/8/how-explain-microservices-plain-english) 和 [12-factor 方法](https://12factor.net/) 作为现代应用程序开发的主要例子。

Ward notes that while microservices and containers can work well  together, the pairing is not actually a necessity, at least under the  right conditions. “Microservices are also mentioned often in conjunction with Kubernetes; however, it is absolutely not a hard requirement,”  Ward says. “A monolithic approach can also work, provided it can scale  horizontally either as a single horizontal deployment or as multiple  deployments with different endpoints to the same codebase.”

Ward 指出，虽然微服务和容器可以很好地协同工作，但配对实际上并不是必需的，至少在合适的条件下是这样。 “微服务也经常与 Kubernetes 一起被提及；然而，这绝对不是硬性要求，”沃德说。 “单体方法也可以工作，前提是它可以作为单个水平部署或作为具有不同端点到同一代码库的多个部署进行水平扩展。”

The same is true of twelve-factor: “Twelve-factor is a useful starting point, but its tenets aren’t necessarily law,” Ward says.

十二因素也是如此：“十二因素是一个有用的起点，但它的原则不一定是法律，”沃德说。

If you’re building an app from scratch, give strong consideration to the microservices approach.

如果您从头开始构建应用程序，请充分考虑微服务方法。

But if you’re building an application from scratch – as Osnat advises teams do when they are getting started with containers and  orchestration – give strong consideration to the microservices approach. 

但是，如果您从头开始构建应用程序——正如 Osnat 建议团队在开始使用容器和编排时所做的那样——请充分考虑微服务方法。

“To maximize the benefits of using containers, architect your app as a microservices app, in a way that will allow it to function even when  individual containers are being refreshed,” Osnat advises. “It should  also be structured so container images represent units that can be  released independently, allowing for an efficient  [CI/CD](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up) implementation.”

“为了最大限度地发挥使用容器的优势，请将您的应用程序构建为微服务应用程序，这样即使在刷新单个容器时也能正常运行，”Osnat 建议道。 “它还应该被构造成容器镜像代表可以独立发布的单元，从而实现高效的 [CI/CD](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up) 执行。”

“Modern” development can be defined in various ways. If you’re  building an application for containers and Kubernetes, it essentially  means making choices that suit these packaging and deployment  technologies. Here are two more examples:

“现代”发展可以有多种定义。如果您正在为容器和 Kubernetes 构建应用程序，则本质上意味着要做出适合这些打包和部署技术的选择。这里还有两个例子：

- **Define container images as logical units that can scale independently:** “For instance, it usually makes sense to implement databases, logging,  monitoring, load-balancing, and user session components as their own  containers or groups of containers,” Osnat says.
- **Consider cloud-native APIs:** “Kubernetes has powerful API extension mechanisms,” says Vaibhav Kamra, VP of engineering and co-founder at [Kasten](https://www.kasten.io/). “By integrating with those, you can immediately take advantage of  existing tools in the ecosystem like command-line utilities and  authentication.”

- **将容器镜像定义为可以独立扩展的逻辑单元：**“例如，将数据库、日志记录、监控、负载平衡和用户会话组件实现为它们自己的容器或容器组通常是有意义的，”奥斯纳特说。
- **考虑云原生 API：**“Kubernetes 具有强大的 API 扩展机制，”[Kasten](https://www.kasten.io/) 的工程副总裁兼联合创始人 Vaibhav Kamra 说。 “通过与这些集成，您可以立即利用生态系统中的现有工具，如命令行实用程序和身份验证。”

“Modern” is a good thing from a software development perspective, too.

从软件开发的角度来看，“现代”也是一件好事。

“The great thing about most modern languages and frameworks is that  they are overwhelmingly container-friendly,” says Ravi Lachhman, DevOps  advocate at [Harness](https://harness.io/). “Going back even a few years, runtimes like Java had a difficult time  respecting container boundaries, and the dreaded out of memory killer  would run seemingly arbitrarily to the operator. Today, because of the  popularity of containers and orchestrators, especially Kubernetes, the  languages and frameworks have evolved to live in the new paradigm.”

[Harness] (https://harness.io/) 的 DevOps 倡导者 Ravi Lachhman 说：“大多数现代语言和框架的伟大之处在于它们对容器非常友好。” “甚至可以追溯到几年前，像 Java 这样的运行时在尊重容器边界方面遇到了困难，而可怕的内存不足杀手似乎对操作员来说是任意运行的。今天，由于容器和编排器（尤其是 Kubernetes）的流行，语言和框架已经发展到可以适应新的范式。”

**[ As Kubernetes adoption grows, what can IT leaders expect? Read also: [5 Kubernetes trends to watch in 2020](https://enterprisersproject.com/article/2020/1/kubernetes-trends-watch-2020?sc_cid=70160000000h0axaaq). ]**

**[ 随着 Kubernetes 的普及，IT 领导者可以期待什么？另请阅读：[2020 年值得关注的 5 个 Kubernetes 趋势](https://enterprisersproject.com/article/2020/1/kubernetes-trends-watch-2020?sc_cid=70160000000h0axaaq)。 ]**

### 2. CI/CD and automation are your friends

### 2. CI/CD 和自动化是您的朋友

Automation is a critical characteristic of container orchestration; it should be a critical characteristic of virtually all aspects of  building an application to be run in containers on Kubernetes. Otherwise, the operational burden can be overwhelming.

自动化是容器编排的一个关键特征；它应该是构建在 Kubernetes 上的容器中运行的应用程序的几乎所有方面的关键特征。否则，操作负担可能是压倒性的。

“Build applications and services with automation as minimum table stakes,” recommends Chander Damodaran, chief architect at [Brillio](https://www.brillio.com/). “With the proliferation of services and components, this can become an unmanageable issue.”

[Brillio](https://www.brillio.com/) 的首席架构师 Chander Damodaran 建议：“以自动化为最低赌注构建应用程序和服务。” “随着服务和组件的激增，这可能成为一个无法管理的问题。”

A well-conceived  CI/CD pipeline can bake automation into many phases of your development and deployment processes.

精心设计的 CI/CD 管道可以将自动化融入开发和部署过程的多个阶段。

A well-conceived  [CI/CD pipeline](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up) is an increasingly popular approach to baking automation into as many  phases of your development and deployment processes as possible. Check  out our recent primer for IT leaders: [How to build a CI/CD pipeline.](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up)

精心设计的 [CI/CD 管道](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up) 是一种日益流行的方法，可将自动化融入开发的多个阶段和部署过程。查看我们最近为 IT 领导者准备的入门读物：[如何构建 CI/CD 管道。](https://enterprisersproject.com/article/2020/1/cicd-pipeline-how-set-up)

Another way to think about the value of automation: It will make your mistakes – which are almost inevitable, especially early on – easier to bounce back from.

思考自动化价值的另一种方式：它会让你的错误——这几乎是不可避免的，尤其是在早期——更容易反弹。

“Using any new platform requires a lot of trial and error, and the  ease of using Kubernetes does not excuse you from taking necessary  precautions,” says Lachhman from Harness. “Having a robust continuous  delivery pipeline in place can ensure that confidence-building standards like testing, security, and change management strategies are followed  to ensure your applications are running effectively.”

Harness 的 Lachhman 说：“使用任何新平台都需要进行大量的反复试验，而 Kubernetes 的易用性并不能成为您采取必要预防措施的借口。” “拥有强大的持续交付管道可以确保遵循测试、安全和变更管理策略等建立信任的标准，以确保您的应用程序有效运行。”

**[ Why does Kubernetes matter to IT leaders? Learn more about Red Hat's [point of view](https://www.redhat.com/en/topics/containers/kubernetes-approach?intcmp=7013a000002DSiEAAW). ]**

**[ 为什么 Kubernetes 对 IT 领导者很重要？详细了解 Red Hat 的 [观点](https://www.redhat.com/en/topics/containers/kubernetes-approach?intcmp=7013a000002DSiEAAW)。 ]**

### 3. Keep container images as light as possible

### 3. 保持容器图像尽可能轻

Another key principle when developing an application for containers  and Kubernetes: Keep your container images as small as possible for  performance, security, and other reasons.

为容器和 Kubernetes 开发应用程序时的另一个关键原则：出于性能、安全性和其他原因，使容器映像尽可能小。

Make sure to remove all other packages – including shell utilities – that are not required by the application. 

确保删除应用程序不需要的所有其他包（包括 shell 实用程序）。

“Only include what you absolutely need. Often images contain packages which are not needed to run the contained application,” says Ken  Mugrage, principal technologist in the office of the CTO at [ThoughtWorks](https://www.thoughtworks.com/). Make sure to remove all other packages – including shell utilities –  that are not required by the application. This not only makes the images smaller but reduces the attack surface for security issues, he says.

“只包括你绝对需要的东西。图像通常包含运行所包含的应用程序不需要的包，”[ThoughtWorks](https://www.thoughtworks.com/) 首席技术官办公室的首席技术专家 Ken Mugrage 说。确保删除应用程序不需要的所有其他包（包括 shell 实用程序)。他说，这不仅使图像变小，而且减少了安全问题的攻击面。

This is a good example of how building a containerized application  might require a shift in traditional practices for some development  teams.

这是一个很好的例子，说明构建容器化应用程序可能需要改变某些开发团队的传统做法。

“Developers need to rethink how they develop applications. For  example, create a smaller container and base image,” says Nilesh Deo,  director of product marketing at [CloudBolt](https://www.cloudbolt.io/). “The smaller the image, the faster it can load, and as a result, the application will be faster.”

“开发人员需要重新思考他们如何开发应用程序。例如，创建一个更小的容器和基础镜像，”[CloudBolt](https://www.cloudbolt.io/) 产品营销总监 Nilesh Deo 说。 “图像越小，加载速度越快，因此应用程序会更快。”

Let’s examine four other best practices:

  让我们来看看其他四个最佳实践：

  

### 4. Don’t blindly trust images

### 4. 不要盲目相信镜像

As is common in software development, there’s a chance you can reuse  or repurpose existing components rather than build them from scratch. The same principle can apply to containers. Just don’t make the mistake  of having blind faith in container images, especially not from a  security perspective.

正如软件开发中常见的那样，您可以重用或重新利用现有组件，而不是从头开始构建它们。同样的原则也适用于容器。只是不要犯对容器镜像盲目相信的错误，尤其是从安全角度来看。



“Far too many people choose an image from a repository with some sort of application stack already installed.”

“太多人从已经安装了某种应用程序堆栈的存储库中选择镜像。”

“Far too many people choose an image from a repository with some sort of application stack already installed,” Mugrage says. “Often these  images are poorly built, and the risk of security issues can’t be  ignored. Any images you use, even ones in your own repositories, should  be scanned for vulnerabilities and compliance on every run of your  deployment pipeline.”

“太多人从已经安装了某种应用程序堆栈的存储库中选择镜像，”Mugrage 说。 “通常这些镜像构建得不好，安全问题的风险不容忽视。您使用的任何图像，即使是您自己存储库中的图像，都应在每次部署管道运行时扫描漏洞和合规性。”

### 5. Plan for observability, telemetry, and monitoring from the start

### 5. 从一开始就计划可观察性、遥测和监控

Kubernetes' self-healing capabilities are a piece of the platform’s appeal, but they also  underscore the need for proper visibility.

Kubernetes 的自我修复功能是该平台的一部分吸引力，但它们也强调了对适当可见性的需求。

Failure is essentially a part of the plan with containers and microservices, but it's more [a matter of managing failure](https://cloud.ibm.com/docs/cloud-native?topic=cloud-native-observability-cn) rather than trying to avoid it altogether. Kubernetes’ self-healing  capabilities are a piece of the platform’s appeal, but they also  underscore the need for proper visibility into your applications and  environments. This is where observability, telemetry, and monitoring  become key.

失败本质上是容器和微服务计划的一部分，但更多的是[管理失败的问题](https://cloud.ibm.com/docs/cloud-native?topic=cloud-native-observability-cn)而不是试图完全避免它。 Kubernetes 的自我修复功能是该平台的一部分吸引力，但它们也强调了对应用程序和环境进行适当可见性的必要性。这就是可观察性、遥测和监控成为关键的地方。

“Kubernetes has built-in mechanisms for resiliency, which create a  need for comprehensive monitoring as a best practice,” says Andrei  Zbikowski, software engineer at [Sentry.io](https://sentry.io/). “Its self-healing functions can restart failed containers or replace  and terminate others when certain health parameters are not met. While  this will keep applications up and running initially, it can actually  conceal growing problems.”

[Sentry.io](https://sentry.io/) 的软件工程师 Andrei Zbikowski 说：“Kubernetes 具有内置的弹性机制，因此需要将全面监控作为最佳实践。” “它的自我修复功能可以在不满足某些健康参数时重新启动失败的容器或替换和终止其他容器。虽然这将使应用程序最初保持正常运行，但它实际上可以掩盖日益严重的问题。”

Zbikowski adds that a lack of visibility into your code might mean an app is throwing errors hourly, for example, even though health metrics  show everything up and running normally.

Zbikowski 补充说，缺乏对代码的可见性可能意味着应用程序每小时都会抛出错误，例如，即使健康指标显示一切正常并正常运行。

“It is important to monitor applications, as well as containers and  back-end systems,” Zbikowski says. “A comprehensive approach to  monitoring will provide greater visibility into issues and events so  that problems can be identified and remediated before there is any  significant impact to users.”

“监控应用程序以及容器和后端系统很重要，”Zbikowski 说。 “全面的监控方法将提供对问题和事件的更大可见性，以便可以在对用户产生任何重大影响之前识别和修复问题。”

Trying to bolt monitoring on to your containerized applications later down the line might lead to unsatisfactory results, Mugrage says. “Think about observability and monitoring from the beginning. Troubleshooting distributed applications is hard and has to be included  in the application design. Adding a monitoring solution later will leave you disappointed.” (Mugrage points to this [piece on observability](https://martinfowler.com/articles/domain-oriented-observability.html) as an additional resource.) 

Mugrage 说，稍后尝试将监控锁定到您的容器化应用程序上可能会导致不令人满意的结果。 “从一开始就考虑可观察性和监控。对分布式应用程序进行故障排除很难，必须包含在应用程序设计中。稍后添加监控解决方案会让您失望。” （Mugrage 指出这个 [关于可观察性的文章](https://martinfowler.com/articles/domain-orientation-observability.html)作为附加资源。)

“There is a big toolbox of cloud-native technologies available to  build sophisticated monitoring, tracing, service meshes, and dashboards  into your applications,” says [Red Hat](https://www.redhat.com/en?intcmp=701f2000000tjyaAAA) technologist evangelist [Gordon Haff](https://enterprisersproject.com/user/gordon-haff). He adds, “Prometheus, Jaeger, Kiali, and Istio are just a few of the  projects you may hear about, and new ones are popping up all the time. However, the choice can be overwhelming and integrating all the tools  yourself can be a challenging distraction.” (This is where you can  consider instead an integrated enterprise open source product like Red  Hat OpenShift Container Platform, Haff notes. )

[Red Hat](https://www.redhat.com/en?intcmp=701f2000000tjyaAAA) 技术专家布道者 [Gordon Haff](https://enterprisersproject.com/user/gordon-haff)。他补充道：“Prometheus、Jaeger、Kiali 和 Istio 只是您可能听说过的几个项目，而且新项目一直在不断涌现。然而，选择可能是压倒性的，并且自己整合所有工具可能是一种具有挑战性的分心。” （在这种情况下，您可以考虑使用集成的企业开源产品，如 Red Hat OpenShift Container Platform，Haff 指出。)

**[ Read also: [OpenShift and Kubernetes: What's the difference?](https://www.redhat.com/en/blog/openshift-and-kubernetes-whats-difference?intcmp=701f2000000tjyaAAA&extIdCarryOver=true&sc_cid=70160000000cYRWAA2) ] **

**[另请阅读：[OpenShift 和 Kubernetes：有什么区别？](https://www.redhat.com/en/blog/openshift-and-kubernetes-whats-difference?intcmp=701f2000000tjyaAAA&extIdCarryOver=true&sc_cid=701600000020cYRW) **

### 6. Consider starting with stateless applications

### 6. 考虑从无状态应用程序开始

One early line of thinking about containers and Kubernetes has been  that running stateless apps is a lot easier than running stateful apps  (such as databases). That’s changing with the growth of [Kubernetes Operators](https://enterprisersproject.com/article/2019/2/kubernetes-operators-plain-english), but teams new to Kubernetes might still be better served by beginning with stateless applications.

关于容器和 Kubernetes 的早期思路是，运行无状态应用程序比运行有状态应用程序（例如数据库）容易得多。随着 [Kubernetes Operators](https://enterprisersproject.com/article/2019/2/kubernetes-operators-plain-english) 的发展，这种情况正在发生变化，但从无状态应用程序开始，Kubernetes 的新团队可能仍会得到更好的服务。

“The benefits of running applications in containers and Kubernetes  are numerous, and there are steps that developers can take to make sure  they are prospering from all those benefits. If I had to pick just one,  the most important step is developing applications with a stateless,  rather than stateful, backend,” says Chris Parmer, co-founder at [Plotly](https://plot.ly/). “Through a stateless backend, development teams can ensure there are no long-running connections or mutable states that make it harder to  scale. Developers will also be able to deploy applications more easily  with zero downtime and enable end-user requests to be delivered in  parallel to different containers.”

“在容器和 Kubernetes 中运行应用程序的好处很多，开发人员可以采取一些步骤来确保他们从所有这些好处中繁荣发展。如果我只能选择一个，那么最重要的一步就是开发具有无状态而不是有状态后端的应用程序，”[Plotly](https://plot.ly/) 的联合创始人 Chris Parmer 说。 “通过无状态后端，开发团队可以确保没有长时间运行的连接或可变状态，这使得扩展变得更加困难。开发人员还将能够在零停机时间的情况下更轻松地部署应用程序，并使最终用户请求能够并行交付到不同的容器。”

**[ Also read [Kubernetes Operators: 4 facts to know](https://enterprisersproject.com/article/2020/2/kubernetes-operators-4-things-know). ]**

**[另请阅读 [Kubernetes Operators: 4 个要知道的事实](https://enterprisersproject.com/article/2020/2/kubernetes-operators-4-things-know)。 ]**

Parmer notes that scalability is one of the major draws of running  containers on Kubernetes; that benefit will be easier to realize with a  stateless app.

Parmer 指出，可扩展性是在 Kubernetes 上运行容器的主要吸引力之一；使用无状态应用程序将更容易实现这种好处。

“Stateless applications make it easy to migrate and scale as  necessary to meet the business needs of the organization, allowing teams to add or remove containers at will,” Parmer says. “By using web  application frameworks that are built upon stateless backends, you get  the most out of your Kubernetes cluster.”

“无状态应用程序可以根据需要轻松迁移和扩展以满足组织的业务需求，允许团队随意添加或删除容器，”Parmer 说。 “通过使用基于无状态后端构建的 Web 应用程序框架，您可以充分利用 Kubernetes 集群。”

### 7. Remember, this is hard

### 7. 记住，这很难

MORE ON KUBERNETES

更多关于 KUBERNETES

- [How to explain Kubernetes Operators in plain English](https://enterprisersproject.com/article/2019/2/kubernetes-operators-plain-english)
- [Kubernetes in production vs. Kubernetes in development: 4 myths](https://enterprisersproject.com/article/2018/11/kubernetes-production-4-myths-debunked)
- [Kubernetes: 6 secrets of successful teams ](https://enterprisersproject.com/article/2020/2/kubernetes-6-secrets-success) 

- [如何用通俗的英语解释 Kubernetes Operator](https://enterprisersproject.com/article/2019/2/kubernetes-operators-plain-english)
- [生产中的 Kubernetes 与开发中的 Kubernetes：4 个神话](https://enterprisersproject.com/article/2018/11/kubernetes-production-4-myths-debunked)
- [Kubernetes：成功团队的 6 个秘诀](https://enterprisersproject.com/article/2020/2/kubernetes-6-secrets-success)

“None of the abstractions that exist in Kubernetes today make the  underlying systems any easier to understand. They only make them easier  to use,” says Chris Short, Red Hat OpenShift principal technical  marketing manager. “If this were easy, everyone would be doing it  already. The industry would be moving on from the Kubernetes hype to the next big thing. This stuff is hard. We’re doing container orchestration while abstracting away the need to manage much other than the state of  the cluster and the infrastructure underneath it. [Etcd](https://etcd.io/) is a huge Kubernetes dependency that a lot of people have had nicely  tucked away from them. There is networking, security, and everything  else wrapped up in Kubernetes too. If your teams aren’t expecting  failure and ready to learn from mistakes, then I need to figure out how  you built the perfect Kubernetes environment.”

“如今 Kubernetes 中存在的任何抽象都无法让底层系统更容易理解。它们只会让它们更易于使用，”红帽 OpenShift 首席技术营销经理 Chris Short 说。 “如果这很容易，那么每个人都会已经这样做了。该行业将从 Kubernetes 炒作转向下一件大事。这东西很难。我们在进行容器编排的同时，将管理除集群状态和其下的基础设施之外的其他管理需求抽象化。 [Etcd](https://etcd.io/) 是一个巨大的 Kubernetes 依赖项，很多人已经很好地避开了它们。 Kubernetes 中还包含网络、安全性和其他所有内容。如果您的团队没有预料到失败并准备从错误中吸取教训，那么我需要弄清楚您是如何构建完美的 Kubernetes 环境的。”

