# Accelerate Developer Teams with Platform Engineering

# 通过平台工程加速开发团队

November 5, 2021

Engineering organizations often look for ways to improve their engineering teams’ efficiency. The more efficient the team, the faster they can ship new features and products to their customer base. From this need for efficiency, combined with developer empathy, we’ve seen the rise of DevOps and site reliability engineering across the industry.

工程组织经常寻找提高工程团队效率的方法。团队效率越高，他们就能越快地将新功能和产品交付给客户群。从这种对效率的需求，再加上开发人员的同理心，我们已经看到 DevOps 和站点可靠性工程在整个行业中的兴起。

These trends have pushed companies to automate and standardize infrastructure, CI/CD, and operations. These standardized processes let engineering teams gain back time to write code, develop features, and further improve their product.

这些趋势促使公司实现基础架构、CI/CD 和运营的自动化和标准化。这些标准化流程让工程团队有更多时间来编写代码、开发功能并进一步改进他们的产品。

However, in today’s world of complex dependencies, disparate tools, and cloud-based infrastructure, many teams still find themselves spending more and more time on understanding and maintaining their platform, infrastructure, CI/CD, and operational tools. When the current toolset doesn’t meet their needs, there’s no one to turn to but their own team. And they’re back to working through infrastructure issues or relearning problems and solutions that others have already solved.

然而，在当今复杂的依赖关系、不同的工具和基于云的基础设施的世界中，许多团队仍然发现自己花费越来越多的时间来理解和维护他们的平台、基础设施、CI/CD 和运营工具。当当前的工具集不能满足他们的需求时，除了他们自己的团队外，别无他法。他们重新开始解决基础设施问题或重新学习其他人已经解决的问题和解决方案。

Enter: platform engineering. Platform engineering can help create an environment where development teams can thrive, bringing your teams back to their mission of developing products rather than simply maintaining their tools.

输入：平台工程。平台工程可以帮助创建一个开发团队可以蓬勃发展的环境，让您的团队回到他们开发产品的使命，而不是简单地维护他们的工具。

So what is platform engineering? And how will it help accelerate your development teams?

那么什么是平台工程？它将如何帮助加速您的开发团队？

## The Evolution of Platform Engineering

## 平台工程的演变

To begin, let’s talk about how platform engineering came to be.

首先，让我们谈谈平台工程是如何形成的。

As often is the case, the role of platform engineering isn’t completely new. But naming this need within a company gives it a definition and focus. It gives teams and organizations a common language to help them solve platform problems instead of one-off engineering problems.

通常情况下，平台工程的作用并不是全新的。但是在公司内部命名这种需求可以给它一个定义和重点。它为团队和组织提供了一种通用语言来帮助他们解决平台问题，而不是一次性的工程问题。

In the past, companies tried to use managed services to reduce the time development teams spent maintaining their platforms. Whether internal or external, managed services provide tools that teams can use to manage components of their DevOps lifecycle. But managed services didn’t solve the whole problem.

过去，公司试图使用托管服务来减少开发团队维护平台的时间。无论是内部还是外部，托管服务都提供了团队可用于管理其 DevOps 生命周期组件的工具。但是托管服务并没有解决整个问题。

Each service focused on a particular silo of functionality. So a monitoring team focused on monitoring solutions, but not always the integrations with other teams or products. Also, the managed service often provided generic expertise, but not guidance on how the product worked with the company’s specific tech stacks and processes. And looking at a platform holistically and assessing how it was fitting (or wasn’t) in a company’s developer ecosystem wasn’t considered as much as it should have been.

每项服务都专注于特定的功能孤岛。因此，监控团队专注于监控解决方案，但并不总是与其他团队或产品的集成。此外，托管服务通常提供通用专业知识，但不提供有关产品如何与公司特定技术堆栈和流程配合使用的指导。从整体上看一个平台并评估它如何适合（或不适合）公司的开发人员生态系统并没有得到应有的考虑。

With managed services, you’re often adding another product–and operations of that product. It’s siloed and not integrated with your overall technical strategy. And that leaves a gap. You also need technical vision, direction, consulting, and engineering growth to plug that gap.

使用托管服务，您通常会添加另一种产品以及该产品的操作。它是孤立的，没有与您的整体技术策略集成。这留下了一个空白。您还需要技术愿景、方向、咨询和工程发展来填补这一空白。

So how can platform engineering take us further? By analyzing the platform, its boundaries, and how it solves problems for the (internal) customer. Ultimately, platform engineering provides a product to your application teams. And that product allows teams to deliver faster.

那么平台工程如何才能让我们走得更远呢？通过分析平台、其边界以及它如何为（内部）客户解决问题。最终，平台工程为您的应用程序团队提供了产品。该产品使团队能够更快地交付。

> To ensure platform engineering in your organization succeeds, you must consider it a product. The customers of this product are the engineering teams. And in short, the goal of platform engineering is to help these engineering teams succeed.

> 为确保您组织中的平台工程成功，您必须将其视为一种产品。该产品的客户是工程团队。简而言之，平台工程的目标是帮助这些工程团队取得成功。

## Platform Engineering as a Product

## 平台工程即产品

Let’s go a bit further into the product aspect of platform engineering.

让我们更深入地探讨平台工程的产品方面。

To ensure platform engineering in your organization succeeds, you must consider it a product. The customers of this product are the engineering teams. And in short, the goal of platform engineering is to help these engineering teams succeed. 

为了确保您组织中的平台工程成功，您必须将其视为一种产品。该产品的客户是工程团队。简而言之，平台工程的目标是帮助这些工程团队取得成功。

When companies view platform engineering as a product, its goals and direction become much clearer. For example, perhaps your company’s platform consists of deploying microservices and serverless functions to a particular cloud provider. The platform engineering team should then drive to make that as simple, automated, and reliable as needed for their engineering teams.

当公司将平台工程视为一种产品时，其目标和方向就会变得更加清晰。例如，也许您公司的平台包括将微服务和无服务器功能部署到特定的云提供商。然后，平台工程团队应该努力使其工程团队所需的简单、自动化和可靠。

Sometimes that means standardizing deployments and rollback strategies or providing tools and libraries so teams don’t have to learn the inner workings of their infrastructure. Either way, platform engineering teams become experts at allowing application dev teams to deliver features without worrying about the rest.

有时，这意味着标准化部署和回滚策略或提供工具和库，以便团队不必了解其基础架构的内部工作原理。无论哪种方式，平台工程团队都成为允许应用程序开发团队交付功能而不用担心其他问题的专家。

## What Platform Engineering Can Do for Your Developers

## 平台工程可以为您的开发人员做什么

Though we hinted a bit about this in the last section, let’s take a look at some of the benefits of platform engineering.

尽管我们在上一节中对此进行了一些暗示，但让我们来看看平台工程的一些好处。

First, platform engineering teams should allow development teams to focus on products. So platform teams will focus not on the product itself but on how it runs on or deploys to your particular infrastructure and platforms. The in-depth knowledge of how all these systems work together can then inform specific recommendations for your product.

首先，平台工程团队应该让开发团队专注于产品。因此，平台团队将不关注产品本身，而是关注它如何在您的特定基础架构和平台上运行或部署。对所有这些系统如何协同工作的深入了解可以为您的产品提供具体建议。

When you think about the code that powers your infrastructure, remember that it’s not a “write once and forget” endeavor. Whether it’s your infrastructure as code, your CI/CD pipelines, or the tools that provide operational support, it all requires enhancements and upgrades.

当您考虑为您的基础设施提供支持的代码时，请记住这不是“一次编写并忘记”的努力。无论是您的基础设施即代码、您的 CI/CD 管道，还是提供运营支持的工具，都需要增强和升级。

Whenever underlying platforms upgrade, the products built on top of them may have opportunity to use new features or optimizations. On top of that, the platform engineering team will drive changes and optimizations as your organization and customer base grows.

每当底层平台升级时，建立在它们之上的产品可能有机会使用新功能或优化。最重要的是，随着您的组织和客户群的增长，平台工程团队将推动变革和优化。

Additionally, platform engineering teams can create simple abstractions for development teams to use when interfacing with infrastructure, like databases or monitors.

此外，平台工程团队可以创建简单的抽象，供开发团队在与基础设施（如数据库或监视器）交互时使用。

And the expertise that you gain from all of that will take you further. Over time, the platform engineering team can build a tailored-to-you set of best practices and standards. So if there’s a concurrency pattern that works well for your microservices, platform engineering can roll out and standardize this. The practice and knowledge can become commonplace through libraries or automated checks that validate proper implementation or behavior.

您从所有这些中获得的专业知识将使您走得更远。随着时间的推移，平台工程团队可以构建一套为您量身定制的最佳实践和标准。因此，如果有一种适合您的微服务的并发模式，平台工程可以推出并标准化它。通过验证正确实施或行为的库或自动检查，实践和知识可以变得司空见惯。

**Building a Platform Engineering team?**

**建立平台工程团队？**

Accelerate their work with OpsLevel and let your entire software development organization thrive.

使用 OpsLevel 加速他们的工作，让您的整个软件开发组织蓬勃发展。

[See OpsLevel in Action](https://opslevel.wistia.com/medias/7hly72v4ev)

[见 OpsLevel 实战](https://opslevel.wistia.com/medias/7hly72v4ev)

## Getting Started

##  入门

To build your platform engineering organization, identify the biggest pain points that your application teams experience. Look at what parts of the development process cause the most stress and inefficiency. And of all the pain points, look at ones where dedicated engineering time can improve the whole organization.

要建立您的平台工程组织，请确定您的应用程序团队遇到的最大痛点。看看开发过程的哪些部分造成的压力最大和效率低下。在所有的痛点中，看看专门的工程时间可以改善整个组织的痛点。

Oftentimes you’ll find that one of the most significant pain points is communication. Teams seem to be re-solving similar problems again and again. One problem that devs struggle with is that there’s a plethora of information on their framework or platform, but it becomes overwhelming. They forget what’s available or where to go to make it happen or simply can’t find what they need.

通常，您会发现最重要的痛点之一是沟通。团队似乎一次又一次地重新解决类似的问题。开发人员面临的一个问题是，他们的框架或平台上有大量信息，但它变得势不可挡。他们忘记了可用的东西或去哪里实现它，或者根本找不到他们需要的东西。

Over time, platform engineering teams will grow and utilize a number of tools, frameworks, and processes. No one will know it all to start.

随着时间的推移，平台工程团队将不断壮大并利用大量工具、框架和流程。没有人会一开始就知道这一切。

With all these tools and communication pain points, the team will quickly develop a need for a central hub of information that's integrated well with your development ecosystem.. And this is where OpsLevel comes in. With [centralized microservice catalogs](http://www.opslevel.com/microservice-catalog/), teams can have a central hub that provides them with the information needed to build, maintain, operate, and integrate their system with other systems. 

有了所有这些工具和沟通痛点，团队将很快需要一个与您的开发生态系统良好集成的中央信息中心。 www.opslevel.com/microservice-catalog/），团队可以拥有一个中央枢纽，为他们提供构建、维护、操作以及将他们的系统与其他系统集成所需的信息。

Cut down on the confusion and help guide teams through the life cycle of the product by providing all the integrations and entry points in one place. Whether it's security, [service maturity](http://www.opslevel.com/blog/service-maturity-opslevel/), [CI/CD](http://www.opslevel.com/blog/jenkins-octopus-github-actions-integrate-CI-CD-pipelines-with-opslevel/), or monitoring, OpsLevel can give your engineers the central hub that helps them search less and deliver faster.

通过在一个地方提供所有集成和入口点，减少混乱并帮助指导团队完成产品的生命周期。无论是安全性、[服务成熟度](http://www.opslevel.com/blog/service-maturity-opslevel/)、[CI/CD](http://www.opslevel.com/blog/jenkins-octopus-github-actions-integrate-CI-CD-pipelines-with-opslevel/) 或监控，OpsLevel 可以为您的工程师提供中心枢纽，帮助他们减少搜索并加快交付速度。

## Summary

##  概括

If your teams struggle to ship features because of the specialized knowledge required or too many manual processes, hold up. You may need a platform engineering team. This team will dedicate their time to improving processes across the software development life cycle for your application teams.

如果您的团队由于所需的专业知识或过多的手动流程而难以交付功能，请稍等。您可能需要一个平台工程团队。该团队将花时间为您的应用程序团队改进整个软件开发生命周期的流程。

If you want to hear more, [request a custom demo of OpsLevel](http://www.opslevel.com/request-demo/) and hear how we can help your teams navigate your microservice ecosystem. Make it easy for your teams to follow established patterns and tools through easy integrations and a one-stop shop.

如果您想了解更多信息，请 [请求 OpsLevel 的自定义演示](http://www.opslevel.com/request-demo/) 并了解我们如何帮助您的团队导航您的微服务生态系统。通过简单的集成和一站式服务，让您的团队轻松遵循既定的模式和工具。

This post was written by Sylvia Fronczak. [Sylvia](https://sylviafronczak.com/) is a software developer that has worked in various industries with various software methodologies. She’s currently focused on design practices that the whole team can own, understand, and evolve over time. 

这篇文章是由 Sylvia Fronczak 撰写的。 [Sylvia](https://sylviafronczak.com/) 是一位软件开发人员，曾在各个行业使用各种软件方法。她目前专注于整个团队可以拥有、理解和随着时间发展的设计实践。

