# DevOps — You Build It, You Run It

# DevOps — 你构建它，你运行它

Jun 27, 2019

2019 年 6 月 27 日

https://medium.com/usa-today-network/devops-you-build-it-you-run-it-8f972343eb8e

https://medium.com/usa-today-network/devops-you-build-it-you-run-it-8f972343eb8e

We know DevOps is a paradigm shift for software development. But the wide range of DevOps definitions and processes make it difficult, especially across organizations, to define what DevOps is in practice. Is it a development team working well with an operations team? Or is it a third team connecting the others? The API Services team at USA TODAY NETWORK doesn’t fit either of those scenarios. Our approach can be summed up by the slogan, “You Build It, You Run It.”

我们知道 DevOps 是软件开发的范式转变。但是，广泛的 DevOps 定义和流程使得定义 DevOps 在实践中的内容变得困难，尤其是跨组织。开发团队是否与运营团队合作良好？或者它是连接其他人的第三个团队？ USA TODAY NETWORK 的 API 服务团队不适合这两种情况。我们的方法可以用“你构建它，你运行它”的口号来概括。

# One Team

#  一个队

DevOps at its core is about aligning the responsibility for running services with the authority to improve them. The quality and velocity of improvements for a service are greatly improved when responsibility and authority are in alignment. A dev team and an operations team which don’t collaborate are functionally broken, because the day-to-day responsibility rests on the operations team while the authority to make changes belongs to the dev team. One way to solve this is to improve communication between teams while building shared responsibility for running the service, along with shared authority to modify it. The approach we take is similar. We don’t have two teams. A single team both builds the service and runs it. It works well for us!

DevOps 的核心是将运行服务的责任与改进服务的权限保持一致。当责任和权力保持一致时，服务改进的质量和速度会大大提高。一个不合作的开发团队和一个运营团队在功能上被打破了，因为日常的责任在于运营团队，而做出改变的权力属于开发团队。解决这个问题的一种方法是改善团队之间的沟通，同时建立运行服务的共同责任，以及修改它的共享权限。我们采取的方法是相似的。我们没有两支球队。一个团队既构建服务又运行它。它对我们很有效！

# Benefits

#  好处

**Organizational Benefits**

**组织利益**

Maintaining a single team with aligned responsibility and authority means we don’t need cross-team communication. Instead, we can rely on much easier intra-team communication. Individual skillsets still vary, but the goals and responsibilities are shared by the team as a whole. Everyone becomes more aligned. Also, cross-team communication is no longer tied up with service level concerns and becomes freed up for discussions regarding composition of services and other higher-level concerns.

维持一个责任和权力一致的单一团队意味着我们不需要跨团队沟通。相反，我们可以依靠更轻松的团队内部沟通。个人技能组合仍然各不相同，但目标和责任由整个团队共同分担。每个人都变得更加一致。此外，跨团队沟通不再与服务级别问题相关联，而是可以自由讨论有关服务组合和其他更高级别问题的讨论。

Here are some other reasons we’ve found that “You Build It, You Run It” leads to better services:

以下是我们发现“您构建它，您运行它”带来更好服务的其他一些原因：

- Improved issue visibility: A team responsible for running a service is careful to make it run well. (This is true for any DevOps organization, especially when you will be paged for incidents.)
- Reduced issue resolution time: You’re dealing with a single team, so delays caused by cross-team communication disappear and everyone involved understands both the operation and development of the service.
- Immediate feedback loop: The feedback loop for problems is automatic; no need to communicate out the cause of a problem, just immediately fix it. A faster feedback loop equals better service.

- 提高问题的可见性：负责运行服务的团队会小心翼翼地使其运行良好。 （对于任何 DevOps 组织都是如此，尤其是当您因事件而被寻呼时。）
- 减少问题解决时间：您正在与一个团队打交道，因此跨团队沟通造成的延迟消失了，每个参与的人都了解服务的运营和开发。
- 即时反馈循环：问题的反馈循环是自动的；无需沟通问题的原因，只需立即修复即可。更快的反馈循环等于更好的服务。

**Engineer Benefits**

**工程师福利**

Teams can be formed with engineers with various specialties. But whenever practical, we strive for every engineer to learn the skillsets needed for both operations and development. Building a whole team of highly-skilled and trusted engineers is a challenge, but it’s not insurmountable. Engineers like to be both highly-skilled and trusted, so setting this expectation — along with plenty of learning resources — drives growth.

团队可以由具有不同专业的工程师组成。但只要可行，我们就会努力让每位工程师都学习运营和开发所需的技能。建立一支由高技能和值得信赖的工程师组成的整个团队是一项挑战，但并非无法克服。工程师喜欢既高技能又值得信赖，因此设定这种期望以及大量学习资源可以推动增长。

Engineers embrace this approach because of the resulting skillset growth and job satisfaction opportunities. Here’s why:

工程师之所以采用这种方法，是因为由此带来的技能组合增长和工作满意度机会。原因如下：

- Execution is the best teacher. For example, devs learn to better incorporate instrumentation and security into an application when they also run the service.
- You can better build a more effective and relevant product when you see and truly understand the whole picture. It also allows you to make wise tradeoffs on features, development speed, security, etc. In short, you become a better engineer. Understanding the whole is also more satisfying than working on a portion with only limited context.
- Building a service paired with seeing it used — as well as using it yourself — is satisfying.

- 执行是最好的老师。例如，开发人员在运行服务时学会更好地将仪器和安全性整合到应用程序中。
- 当您看到并真正了解全局时，您可以更好地构建更有效和相关的产品。它还允许您在功能、开发速度、安全性等方面做出明智的权衡。简而言之，您将成为一名更好的工程师。理解整体也比只处理有限上下文的部分更令人满意。
- 建立一个服务并看到它被使用——以及自己使用它——是令人满意的。

# Beyond One Team 

# 超越一个团队

The success we’ve described isn’t possible without a set of robust, well-maintained services to build on. Today's cloud-native infrastructure allows a single team to run an application that requires a CI system, database, Kubernetes, etc. An engineer can confidently use these services and rarely have to worry about service-specific details, so her time is freed up for other concerns. When details for a dependent service require investigation, the service maintainer can be brought in for assistance. For these reasons, building on top of robust services prevents the amount of work in the system from growing beyond what a single team can handle.

如果没有一套强大的、维护良好的服务作为基础，我们所描述的成功是不可能的。今天的云原生基础设施允许单个团队运行需要 CI 系统、数据库、Kubernetes 等的应用程序。工程师可以自信地使用这些服务，而不必担心服务特定的细节，因此她的时间可以腾出来用于其他顾虑。当依赖服务的详细信息需要调查时，可以请服务维护人员寻求帮助。由于这些原因，建立在强大的服务之上可以防止系统中的工作量增长到超出单个团队可以处理的范围。

Even with cloud-native services, nearly every company has plenty of work for multiple teams. Embracing an “as a service” team organization model can help teams divide and conquer along lines that enable each to build and run their own services. If CI for your company is run internally, the team responsible can offer CI as a service. Other teams then consume it as a tool they use when building their own services. This enables both the CI team and teams that use that service to follow the “You Build It, You Run It” model of DevOps.

即使使用云原生服务，几乎每家公司都需要为多个团队做大量工作。采用“即服务”的团队组织模型可以帮助团队分而治之，使每个人都能够构建和运行自己的服务。如果您公司的 CI 在内部运行，则负责的团队可以将 CI 作为服务提供。然后其他团队在构建自己的服务时将其用作他们使用的工具。这使得 CI 团队和使用该服务的团队都能够遵循 DevOps 的“你构建它，你运行它”模型。

# Conclusion

#  结论

We’ve found that “You Build It, You Run It” doesn’t just bridge the fence between development and operations teams, but rather pulls that fence down altogether and embodies the spirit of DevOps. Immediate feedback on changes leads to quick fixes and better services. Working this way enables engineers to become better, grow in their skillset and improve their job satisfaction.

我们发现，“你构建它，你运行它”不仅在开发和运营团队之间架起了一座桥梁，而是将这道栅栏完全推倒，体现了 DevOps 的精神。对更改的即时反馈可带来快速修复和更好的服务。以这种方式工作可以使工程师变得更好，提高他们的技能并提高他们的工作满意度。

## [More from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----8f972343eb8e--------------------------------)

Gannett is the largest local news organization in America. Our brands include USA TODAY and 250+ newsrooms spanning 46 states. Gannett’s vastly expanded local-to-national platform reaches over 50% of the U.S. digital audience, including more Millennials than Buzzfeed. 

Gannett 是美国最大的地方新闻机构。我们的品牌包括《今日美国》和遍布 46 个州的 250 多家新闻编辑室。 Gannett 大大扩展的本地到国家平台覆盖了超过 50% 的美国数字受众，其中包括比 Buzzfeed 更多的千禧一代。

