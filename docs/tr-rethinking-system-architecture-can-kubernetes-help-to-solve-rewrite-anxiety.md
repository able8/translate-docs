# Using Kubernetes to rethink your system architecture and ease technical debt

# 使用 Kubernetes 重新思考您的系统架构并减轻技术债务

Developers are famed for wanting to rewrite software, especially when they “inherited” the software and they can’t be bothered to understand how it works. Experienced managers and senior engineers know that [rewrites](https://www.joelonsoftware.com/2000/04/06/things-you-should-never-do-part-i/) should be avoided unless they are truly necessary , as they typically involve a lot of complexity and can introduce new problems along the way.

开发人员以想要重写软件而闻名，特别是当他们“继承”软件并且他们无法理解它是如何工作的时候。有经验的经理和高级工程师都知道，除非真的有必要，否则应该避免[重写](https://www.joelonsoftware.com/2000/04/06/things-you-should-never-do-part-i/) ，因为它们通常涉及很多复杂性，并且可能会在此过程中引入新问题。

This is a story about trying to rethink complex systems: the challenges you face when you try to rebuild them, the burdens you face as they grow, and how inaction itself can cause it’s own problems. When you’re weighing the risk and reward of replacing architecture, it can take several attempts to find a solution that works for you.

这是一个关于尝试重新思考复杂系统的故事：您在尝试重建它们时面临的挑战，随着它们的增长而面临的负担，以及不作为本身如何导致其自身的问题。当您权衡更换架构的风险和回报时，可能需要多次尝试才能找到适合您的解决方案。

I’m a Senior Engineer at Pusher, a company focused on building real-time messaging APIs. [Pusher Channels](https://pusher.com/channels), our pub/sub WebSocket service used for building scalable realtime data functionality, has been around [for quite a while](https://web.archive.org/web/20110613062158/http://www.pusher.com:80/). Until fairly recently, all of Channels ran on AWS EC2 instances. Machines were provisioned and bootstrapped with Python scripts which wrapped [Ansible](https://en.wikipedia.org/wiki/Ansible_(software)) playbooks. Configuration and process management was mostly handled by [Puppet](https://en.wikipedia.org/wiki/Puppet_%28software%29), with help from [Upstart](https://en.wikipedia.org/wiki/Upstart_(software)), [God](http://godrb.com/), and a lot of tools written in-house.

我是 Pusher 的高级工程师，这是一家专注于构建实时消息传递 API 的公司。 [Pusher Channels](https://pusher.com/channels)，我们的发布/订阅 WebSocket 服务，用于构建可扩展的实时数据功能，已经存在 [很长一段时间](https://web.archive.org/web/20110613062158/http://www.pusher.com:80/)。直到最近，所有通道都在 AWS EC2 实例上运行。机器使用 Python 脚本进行配置和引导，这些脚本包装了 [Ansible](https://en.wikipedia.org/wiki/Ansible_(software)) 剧本。配置和流程管理主要由 [Puppet](https://en.wikipedia.org/wiki/Puppet_%28software%29) 在 [Upstart](https://en.wikipedia.org/wiki/Upstart_(software))，[God](http://godrb.com/)，以及很多内部编写的工具。

We managed EC2 instances like [pets](http://cloudscaling.com/blog/cloud-computing/the-history-of-pets-vs-cattle/). If a machine had to be replaced, an engineer manually migrated traffic/services from the old machine to the new one, and shut down the old one. If a cluster needed more capacity, an engineer provisioned some new machines and attached them to the cluster. This approach had its downsides:

- There was a chunk of manual work involved in keeping stuff up and running
- The in house tooling made onboarding new team members tricky

我们管理 EC2 实例，如 [pets](http://cloudscaling.com/blog/cloud-computing/the-history-of-pets-vs-cattle/)。如果必须更换一台机器，工程师手动将流量/服务从旧机器迁移到新机器，然后关闭旧机器。如果集群需要更多容量，工程师会配置一些新机器并将它们连接到集群。这种方法有其缺点：

- 有大量的手动工作涉及保持东西的正常运行
- 内部工具使新团队成员入职变得棘手

…but it worked pretty well for quite a while and large parts of this system are still pushing us to [new heights](https://blog.pusher.com/realtime-results-election-washington-post/). Team members changed, we launched [Pusher Beams](https://pusher.com/beams), we sunsetted Pusher Chatkit, and our user base just kept on growing. All the while, the Pusher Channels infrastructure stayed mostly the same while seriously increasing in scale. And this is where the problems started.

……但它在很长一段时间内运行良好，并且该系统的大部分内容仍在将我们推向 [新高度](https://blog.pusher.com/realtime-results-election-washington-post/)。团队成员发生了变化，我们推出了 [Pusher Beams](https://pusher.com/beams)，我们淘汰了 Pusher Chatkit，我们的用户群一直在增长。一直以来，Pusher Channels 基础设施基本保持不变，但规模大幅增加。这就是问题开始的地方。

## **Maintenance**

##  **维护**

As the Channels clusters got larger, the maintenance burden of operating it seemed to scale almost linearly. Before long a significant amount of engineering time each week was spent just keeping things up and running.

随着 Channels 集群变得越来越大，运行它的维护负担似乎几乎呈线性增长。不久之后，每周大量的工程时间都花在了保持设备正常运行上。

It was clear that we needed to do _something_ to reduce the maintenance burden; trying to maintain a highly reliable service for our customers, however, meant we have spent the last few years battling the inevitable build up of tech debt and legacy infrastructure that comes with managing and maintaining complex systems at scale. 

很明显，我们需要做_某事_来减轻维护负担；然而，试图为我们的客户维护高度可靠的服务意味着我们在过去几年中一直在与不可避免的技术债务和遗留基础设施的积累作斗争，这些债务和遗留基础设施是大规模管理和维护复杂系统的。

We made numerous attempts to try and modernise our infrastructure and application in this time but ran into many of the common problems associated with trying to rewrite or re-architect systems. That means that we were still facing the same challenges. A lot of these problems stemmed from trying to modernise the infrastructure and the application at the same time, rather than trying to focus on the biggest challenge that we faced, which is the management and maintenance of the infrastructure itself.

在这段时间里，我们多次尝试尝试对我们的基础设施和应用程序进行现代化改造，但遇到了许多与尝试重写或重新构建系统相关的常见问题。这意味着我们仍然面临着同样的挑战。许多这些问题源于试图同时实现基础设施和应用程序的现代化，而不是试图专注于我们面临的最大挑战，即基础设施本身的管理和维护。

## **The breakthrough**

## **突破**

We had spent some time going back and forth between various approaches. We tried to introduce more automation into our existing infrastructure provisioning system, and we tried a ground up rewrite of our core application services. However, lots of the solutions were entangled with or dependent on each other in non-obvious ways.

我们花了一些时间在各种方法之间来回切换。我们尝试在我们现有的基础设施供应系统中引入更多自动化，并尝试彻底重写我们的核心应用程序服务。然而，许多解决方案以不明显的方式相互纠缠或依赖。

Most of these solutions required new types of infrastructure that we didn’t already have. This then had a “soft” dependency on some of the skills we had in the team to add new types of infrastructure with our existing tooling.

这些解决方案中的大多数都需要我们尚未拥有的新型基础设施。然后，这对我们团队中的一些技能产生了“软”依赖，以使用我们现有的工具添加新类型的基础设施。

After a number of setbacks, we decided to go back to first principles and look again at the most important priorities. We wanted to select a solution that would achieve the key goals with scope to continue iterating in the future. We summarised that the overall aim should be to end up with a platform that is simpler, consistent, easier to maintain and easier to scale. This would allow us to spend less time maintaining and managing platforms and more time building and developing new products and features.

在经历了一些挫折之后，我们决定回到首要原则，重新审视最重要的优先事项。我们想选择一个能够实现关键目标的解决方案，并在未来继续迭代。我们总结出总体目标应该是最终得到一个更简单、一致、更易于维护和更易于扩展的平台。这将使我们能够花更少的时间维护和管理平台，而将更多时间用于构建和开发新产品和功能。

After doing a complete audit of the current state of our infrastructure and application we came to the following conclusions:

1. Offload as much infrastructure complexity as possible to a managed services
2. Migrate as much of the existing application as possible to this new infrastructure without having to do large rewrites of the application codebase
3. Identify components that could be re-envisioned, re-architected, or re-written without adding friction to the infrastructure migration and invest in transforming and simplifying them.

在对我们的基础设施和应用程序的当前状态进行全面审核后，我们得出以下结论：

1. 将尽可能多的基础架构复杂性卸载到托管服务
2. 将尽可能多的现有应用程序迁移到这个新的基础架构，而无需对应用程序代码库进行大量重写
3. 确定可以在不增加基础设施迁移摩擦的情况下重新构想、重新架构或重新编写的组件，并投资于改造和简化它们。

## **The infrastructure**

## **基础设施**

We’d been keen to begin leveraging [auto-scaling groups](https://docs.aws.amazon.com/autoscaling/ec2/userguide/AutoScalingGroup.html) for some time. After analysing the work required to make use of auto-scaling groups, we determined this would require a significant investment to either:

- Extend our custom-built deployment tooling—which depends on EOL software already—to work with cloudinit and migrate from upstart to systemd. This would have also increased onboarding time for new hires due to the bespoke nature of the solution.
- Migrate to a technology that the industry has standardised around, like containers.

一段时间以来，我们一直热衷于开始利用 [自动缩放组](https://docs.aws.amazon.com/autoscaling/ec2/userguide/AutoScalingGroup.html)。在分析了使用自动缩放组所需的工作后，我们确定这将需要大量投资：

- 扩展我们定制的部署工具——它已经依赖于 EOL 软件——与 cloudinit 一起工作并从新贵迁移到 systemd。由于解决方案的定制性质，这也会增加新员工的入职时间。
- 迁移到行业已标准化的技术，例如容器。

We chose to adopt containers because investing time/energy/money into our in house solution made no sense.

我们选择采用容器是因为在我们的内部解决方案中投入时间/精力/金钱是没有意义的。

### Containerising

### 集装箱化

In order to migrate to containers we would need to:

- containerise core application services,
- update the build process for application services to build and store container images,
- choose some way of running those containers in production,
- change the routing process for services’ traffic to handle container termination more gracefully

为了迁移到容器，我们需要：

- 容器化核心应用服务，
- 更新应用服务的构建过程以构建和存储容器镜像，
- 选择在生产中运行这些容器的某种方式，
- 更改服务流量的路由过程以更优雅地处理容器终止

When we broke down the work required to containerise our existing EC2 setup for auto-scaling groups, we were most of the way towards something like [Kubernetes](https://kubernetes.io/), which we were using heavily elsewhere in the company and offered us a great deal more functionality.

当我们分解为自动扩展组将现有 EC2 设置容器化所需的工作时，我们大部分时间都在朝着 [Kubernetes](https://kubernetes.io/) 之类的方向发展，我们在其他地方大量使用它公司并为我们提供了更多的功能。

There would be extra work, for example, adding additional management services to move our core application services to Kubernetes. In the end, we decided that it was worth it for the following reasons:

- **Purpose built for our problem** – Kubernetes is built to solve the problem we had—resiliently managing a workload across many nodes.
- **In-house experience** – We had a lot of Kubernetes experience in-house and were already running multiple clusters.
- **Hiring**– Kubernetes is one of the dominant tools in the space. Hiring engineers with Kubernetes experience, or a desire to learn it, was significantly easier for us than hiring people who wanted to work with Puppet/Ansible. 

会有额外的工作，例如，添加额外的管理服务以将我们的核心应用程序服务迁移到 Kubernetes。最后，我们认为这是值得的，原因如下：

- **为我们的问题而构建** – Kubernetes 旨在解决我们遇到的问题——跨多个节点弹性地管理工作负载。
- **内部经验** – 我们有很多内部 Kubernetes 经验，并且已经在运行多个集群。
- **招聘**——Kubernetes 是该领域的主要工具之一。雇用具有 Kubernetes 经验或渴望学习它的工程师比雇用想使用 Puppet/Ansible 的人容易得多。

- **Onboarding**– Our existing solution was fairly bespoke so any new joiners had to spend a decent chunk of time learning the ins and outs of all of our homemade tooling. Kubernetes has great documentation so new starters can get up to speed more quickly even if they don’t have experience with it.
- **Offloading infrastructure complexity** – This change would allow us to shift complexity to a managed service like [EKS](https://aws.amazon.com/eks/)

- **入职**——我们现有的解决方案是相当定制的，因此任何新加入者都必须花费大量时间来学习我们所有自制工具的来龙去脉。 Kubernetes 有很好的文档，因此即使没有经验的新手也可以更快地上手。
- **卸载基础设施的复杂性** – 这一变化将使我们能够将复杂性转移到像 [EKS](https://aws.amazon.com/eks/) 这样的托管服务上

## **The application**

##  **应用程序**

One of the traps that we have previously fallen into when trying to improve Channels was trying to rewrite large parts of the application while simultaneously  trying to reduce the maintenance burden of running the infrastructure. This tightly coupled approach led to a few setbacks and abandoned attempts. With this new plan, we were mostly finding solutions to port large parts of our application services to Kubernetes without doing large application rewrites.

我们之前在尝试改进 Channels 时陷入的陷阱之一是尝试重写应用程序的大部分内容，同时尝试减少运行基础架构的维护负担。这种紧密耦合的方法导致了一些挫折和放弃的尝试。有了这个新计划，我们主要是寻找解决方案，将我们的大部分应用程序服务移植到 Kubernetes，而无需进行大型应用程序重写。

We were still keen to strategically improve parts of the application, however, so long as it didn’t fundamentally add friction to migrating to new infrastructure. We found this approach gave engineers the freedom to make technical improvements to the application and opportunities to reduce the maintenance burden of application components that frequently run into issues that wouldn’t be solved by porting to Kubernetes.

然而，我们仍然热衷于战略性地改进应用程序的某些部分，只要它不会从根本上增加迁移到新基础架构的摩擦。我们发现这种方法使工程师可以自由地对应用程序进行技术改进，并有机会减少应用程序组件的维护负担，这些应用程序组件经常遇到移植到 Kubernetes 无法解决的问题。

As an example, we recently wrote about our [webhook system](https://blog.pusher.com/migrating-channels-webhooks-from-beanstalkd-to-kinesis-sqs/) from EventMachine on EC2 to Go on Kubernetes. One thing this didn’t cover was re-writing the component that actually dequeues jobs from SQS and sends webhooks. This component had recently started causing a large number of alerts and alarms and was in need of refactoring.

例如，我们最近写了我们的 [webhook 系统](https://blog.pusher.com/migrating-channels-webhooks-from-beanstalkd-to-kinesis-sqs/) 从 EC2 上的 EventMachine 到 Go on Kubernetes。这没有涵盖的一件事是重写实际从 SQS 中取出作业并发送 webhook 的组件。该组件最近开始引发大量警报和警报，需要重构。

## **Re-writing the webhook sender**

## **重写网络钩子发送器**

We knew it was important to port our webhook sender component from EC2 to Kubernetes, as we had begun to experience increasing operational issues with this component. Because of the lack of autoscaling, the processes responsible for sending webhooks were running on dedicated EC2 machines we called sender machines. On one of our clusters, we had four sender machines, each running 12 webhook-sender processes (called Clowns,because they “juggle” jobs from a queue). This was comfortably sufficient for the peak load of the cluster so we had some headroom for unexpected spikes.

我们知道将我们的 webhook 发送器组件从 EC2 移植到 Kubernetes 很重要，因为我们已经开始遇到越来越多的该组件的操作问题。由于缺乏自动缩放，负责发送 webhook 的进程运行在我们称为发送方机器的专用 EC2 机器上。在我们的一个集群上，我们有四台发送方机器，每台机器运行 12 个 webhook 发送方进程（称为 Clowns，因为它们“处理”队列中的作业）。这对于集群的峰值负载来说已经足够了，所以我们有一些意外峰值的空间。

As we had already largely rewritten most of the webhook pipeline in Go, it also made sense to look at completing the process by rewriting the sender itself. The webhook sender is a pretty straightforward piece of software. It reads jobs from an [SQS](https://aws.amazon.com/sqs/) queue and makes HTTP POST requests. The job the process reads from SQS contains everything the process needs to send the HTTP POST request to the customer’s server.

由于我们已经在 Go 中大量重写了大部分 webhook 管道，因此通过重写发送者本身来完成这个过程也是有意义的。 webhook 发送器是一个非常简单的软件。它从 [SQS](https://aws.amazon.com/sqs/) 队列读取作业并发出 HTTP POST 请求。进程从 SQS 读取的作业包含进程将 HTTP POST 请求发送到客户服务器所需的一切。

Because the webhook sender is an embarrassingly parallel process it makes a great fit for taking advantage of the scaling benefits of a system like Kubernetes.By re-writing this in Go we could also realise some performance benefits over EventMachine, much like we had with the webhook packager and publisher.

因为 webhook 发送器是一个令人尴尬的并行进程，它非常适合利用像 Kubernetes 这样的系统的扩展优势。通过在 Go 中重写它，我们还可以实现一些优于 EventMachine 的性能优势，就像我们在使用webhook 打包器和发布器。

Because the webhook sender is a stateless service, it was easy to deploy the new sender alongside the old sender and let them both compete for jobs. This meant that we could gradually roll out the new sender and rely on the old sender to keep servicing the queue in case of unexpected issues. In fact,what we found on some smaller clusters was that the new sender was so efficient the old sender basically had no work to do. 

因为 webhook 发件人是一种无状态服务，所以很容易将新发件人与旧发件人一起部署，让他们都竞争工作。这意味着我们可以逐步推出新的发件人，并依靠旧的发件人在出现意外问题时继续为队列提供服务。事实上，我们在一些较小的集群上发现，新的发送者效率很高，旧的发送者基本上没有工作可做。

![](https://lh6.googleusercontent.com/SImIWL8HgnCBXS_pukI8IF0sDGl1jw7OtYchNlYLEMdLwWazWS1jLW5a9dyiNGJbAxxHPckIx_mJgSpMBww18E9MgwgicTYmSY1iTMAzLMnLgBKdx2SA26oCYZdAlzC8wIt_jM4b)**Fig 1.** Rollout: traffic to the old sender dies down as the new sender's efficiency takes over its work![](https://lh3.googleusercontent.com/6qP9P7vRMwSS2ZlcQfrGzqFv-orXt-n4bFMUWmfGx3OoAnU7oJdI8MDsVVebXnWoBIfF_eLcjvFRrUGoftg5QUbjtfB4xT5urggrGz1dTsnLHJNUzRJvLtUU0gZOPFe_esSgJKw3)**Fig 2.** Rollout: the new sender picks up servicing

.googleusercontent.com/6qP9P7vRMwSS2ZlcQfrGzqFv-orXt-n4bFMUWmfGx3OoAnU7oJdI8MDsVVebXnWoBIfF_eLcjvFRrUGoftg5QUbjtfB4xT5urggrGzJHJUtgs 发送新的选择服务**RollgzJHJKUtgs **Rolls

In the end, our largest cluster went from 4 \* machines, each with 12 clowns, to 3-15 Kubernetes pods, each requesting 250mCPU.

最后，我们最大的集群从 4 台 \* 机器，每台机器有 12 个小丑，到 3-15 个 Kubernetes pod，每个都需要 250mCPU。

![](https://lh3.googleusercontent.com/voprbXmIvHfefHylSl5JgUqZwS-F-ZtPf8LVNakANVCpgROAVNn_gT0R5E0IItkdagto8nhgDlnWywg-VTJJ_0rE1A9qohsBLghiJ2EHQBlY17X1wg5LrR637id0aHNz7bGpVPVZ)**Fig 3.** Change in backlog service over the period of retirement![](https://lh6.googleusercontent.com/xxaFLySFnaaDEHU4fYdiP-mHrt3nx98VNlTiwoSomqm4D5iGcA_C8I98VOfmZ76LmIiVYJapZF-CT1Bt6LIL0rZKP4pcUx8CgFTk87BrOI_0577uiA-S_OIuT3HMFKmB3hj8Z4mh)**Fig 4.** New autoscaling (over a one week period)

com/xxaFLySFnaaDEHU4fYdiP-mHrt3nx98VNlTiwoSomqm4D5iGcA_C8I98VOfmZ76LmIiVYJapZF-CT1Bt6LIL0rZKP4pcUx8CgFTk87BrOI_057OI_057OI_057OI_057OI_057OI_057OI_057OI_057OI_057OIhZ 1 周)

In short, this was highly successful. Another component from our EC2 and Eventmachine architecture crossed off the list and another <N> boxes removed.

简而言之，这是非常成功的。我们的 EC2 和 Eventmachine 架构中的另一个组件从列表中划掉，另一个 <N> 框被删除。

## **Summary**

##  **概括**

Well-documented stories about trying to rebuild, re-architect, and re-envision systems are countless, and often contain a warning about embarking on such projects. Those of us with previous experience in architecture overhauls will almost always avoid rewrites at any cost because they so often go wrong. Also, complex systems are usually that way for a reason. As explained by [Chesterton’s fence](https://fs.blog/2020/03/chestertons-fence/), we are usually better off assuming that the person who came before us knew something we did not.

关于尝试重建、重新构建和重新构想系统的有据可查的故事数不胜数，并且通常包含有关开始此类项目的警告。我们这些以前在架构大修方面有经验的人几乎总是会不惜一切代价避免重写，因为它们经常出错。此外，复杂的系统通常是有原因的。正如 [切斯特顿的围栏](https://fs.blog/2020/03/chestertons-fence/) 所解释的那样，我们通常最好假设先于我们的人知道一些我们不知道的事情。

After some failures and challenges along the way, however, we found a path to progress and an approach that does allow us to rewrite code effectively and efficiently while also reducing our maintenance burden. To be fair, our rewrite probably wasn't what Joel Spolsky meant when he called rewrites one of the [things you should never do](https://www.joelonsoftware.com/2000/04/06/things-you-should-never-do-part-i/).

然而，在此过程中经历了一些失败和挑战之后，我们找到了前进的道路和方法，它确实允许我们有效地重写代码，同时减少我们的维护负担。公平地说，我们的重写可能不是乔尔·斯波尔斯基（Joel Spolsky）所说的重写[你永远不应该做的事情]之一的意思（https://www.joelonsoftware.com/2000/04/06/things-you-should -从不做部分-i/)。

But what we have found is that by identifying components with well defined boundaries, it is possible to rewrite them without throwing entire systems away. If you can identify well-defined boundaries and interfaces, things become a lot easier. In the case of the webhooks pipeline, this had a logical boundary in the form of a queue. We could rewrite parts of this over time, which gave us great ability to test and verify new components in the pipeline while still having the ability to rollback in case of failure.

但我们发现，通过识别具有明确定义边界的组件，可以在不丢弃整个系统的情况下重写它们。如果您可以确定明确定义的边界和接口，事情就会变得容易得多。在 webhooks 管道的情况下，它有一个队列形式的逻辑边界。随着时间的推移，我们可以重写其中的部分内容，这使我们能够测试和验证管道中的新组件，同时仍然能够在发生故障时回滚。

We have also found that with each one of these projects completed, we are realising benefits of less operational burden, which means we can actually speed up the rate of progress. Making that initial breakthrough was the greatest challenge and it’s important to keep up momentum to avoid slipping back into an unsustainable realm. Now that we’ve broken the wall, we’re excited and confident about having a simpler system that allows us to focus on doing the stuff we actually love to do—building great products and cool features, rather than just keeping the lights on.

我们还发现，随着这些项目中的每一个完成，我们都实现了减少运营负担的好处，这意味着我们实际上可以加快进度。实现最初的突破是最大的挑战，重要的是要保持势头，避免滑回到不可持续的领域。现在我们已经打破了壁垒，我们对拥有一个更简单的系统感到兴奋和自信，该系统使我们能够专注于做我们真正喜欢做的事情——构建伟大的产品和酷炫的功能，而不仅仅是保持灯火通明。

_The Stack Overflow blog is committed to publishing interesting articles by developers, for developers. From time to time that means working with companies that are also clients of Stack Overflow’s through our advertising, talent, or teams business. When we publish work from clients, we’ll identify it as Partner Content with tags and by including this disclaimer at the bottom._

_Stack Overflow 博客致力于为开发人员发布开发人员的有趣文章。有时，这意味着通过我们的广告、人才或团队业务与也是 Stack Overflow 客户的公司合作。当我们发布客户的作品时，我们会将其标识为带有标签的合作伙伴内容，并在底部包含此免责声明。_

Tags: [kubernetes](https://stackoverflow.blog/tag/kubernetes/), [partner content](https://stackoverflow.blog/tag/partner-content/), [partnercontent](https://stackoverflow.blog/tag/partnercontent/), [software architecture](https://stackoverflow.blog/tag/software-architecture/) 

标签： [kubernetes](https://stackoverflow.blog/tag/kubernetes/),[合作伙伴内容](https://stackoverflow.blog/tag/partner-content/), [合作伙伴内容](https://stackoverflow.blog/tag/partnercontent/), [软件架构](https://stackoverflow.blog/tag/software-architecture/)

