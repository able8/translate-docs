# Why you don't have to be afraid of Kubernetes

# 为什么你不必害怕 Kubernetes

## Kubernetes is absolutely the simplest, easiest way to meet the needs of complex web applications.

## Kubernetes 绝对是满足复杂 Web 应用程序需求的最简单、最简单的方法。

31 Oct 2019

It was fun to work at a large web property in the late 1990s and early 2000s. My experience takes me back to American Greetings Interactive, where on Valentine's Day, we had one of the top 10 sites on the internet (measured by web traffic). We delivered e-cards for [AmericanGreetings.com](http://AmericanGreetings.com),[BlueMountain.com](http://BlueMountain.com), and others, as well as providing e-cards for partners like MSN and AOL. Veterans of the organization fondly remember epic stories of doing great battle with other e-card sites like Hallmark. As an aside, I also ran large web properties for Holly Hobbie, Care Bears, and Strawberry Shortcake.

在 1990 年代末和 2000 年代初，在大型网络资产中工作很有趣。我的经历让我回到了 American Greetings Interactive，在情人节那天，我们拥有互联网上排名前 10 的网站之一（以网络流量衡量）。我们为[AmericanGreetings.com](http://AmericanGreetings.com)、[BlueMountain.com](http://BlueMountain.com)等提供电子贺卡，并为MSN等合作伙伴提供电子贺卡和美国在线。该组织的退伍军人深情地记得与 Hallmark 等其他电子贺卡网站进行伟大战斗的史诗故事。顺便说一句，我还为 Holly Hobbie、Care Bears 和 Strawberry Shortcake 运营大型网络资产。

I remember like it was yesterday the first time we had a real problem. Normally, we had about 200Mbps of traffic coming in our front doors (routers, firewalls, and load balancers). But, suddenly, out of nowhere, the Multi Router Traffic Grapher (MRTG) graphs spiked to 2Gbps in a few minutes. I was running around, scrambling like crazy. I understood our entire technology stack, from the routers, switches, firewalls, and load balancers, to the Linux/Apache web servers, to our Python stack (a meta version of FastCGI), and the Network File System (NFS) servers. I knew where all of the config files were, I had access to all of the admin interfaces, and I was a seasoned, battle-hardened sysadmin with years of experience troubleshooting complex problems.

我记得就像昨天我们第一次遇到真正的问题一样。通常，我们的前门（路由器、防火墙和负载平衡器）有大约 200 Mbps 的流量。但是，突然间，Multi Router Traffic Grapher (MRTG) 图表在几分钟内飙升至 2Gbps。我跑来跑去，发疯似的爬起来。我了解我们的整个技术堆栈，从路由器、交换机、防火墙和负载平衡器，到 Linux/Apache Web 服务器，再到我们的 Python 堆栈（FastCGI 的元版本）和网络文件系统 (NFS) 服务器。我知道所有配置文件在哪里，我可以访问所有管理界面，而且我是一位经验丰富、久经沙场的系统管理员，拥有多年解决复杂问题的经验。

But, I couldn't figure out what was happening...

但是，我无法弄清楚发生了什么......

Five minutes feels like an eternity when you are frantically typing commands across a thousand Linux servers. I knew the site was going to go down any second because it's fairly easy to overwhelm a thousand-node cluster when it's divided up and compartmentalized into smaller clusters.

当您在上千个 Linux 服务器上疯狂地键入命令时，五分钟感觉就像是永恒。我知道该站点随时都会崩溃，因为当它被划分为更小的集群时，很容易压倒一个千节点的集群。

I quickly _ran_ over to my boss's desk and explained the situation. He barely looked up from his email, which frustrated me. He glanced up, smiled, and said, "Yeah, marketing probably ran an ad campaign. This happens sometimes." He told me to set a special flag in the application that would offload traffic to Akamai. I ran back to my desk, set the flag on a thousand web servers, and within minutes, the site was back to normal. Disaster averted.

我迅速_跑_到我老板的办公桌前并解释了情况。他几乎从他的电子邮件中抬起头来，这让我很沮丧。他抬起头，微笑着说：“是的，营销部门可能开展了一场广告活动。这种情况有时会发生。”他告诉我在应用程序中设置一个特殊标志，将流量卸载到 Akamai。我跑回我的办公桌，在 1000 个网络服务器上设置标志，几分钟后，网站恢复正常。避免了灾难。

I could share 50 more stories similar to this one, but the curious part of your mind is probably asking, "Where this is going?"

我可以再分享 50 个与这个类似的故事，但你心中好奇的部分可能会问，“这是怎么回事？”

The point is, we had a business problem. Technical problems become business problems when they stop you from being able to do business. Stated another way, you can't handle customer transactions if your website isn't accessible.

关键是，我们遇到了业务问题。当技术问题阻止您开展业务时，它们就会成为业务问题。换句话说，如果您的网站无法访问，您就无法处理客户交易。

So, what does all of this have to do with Kubernetes? Everything. The world has changed. Back in the late 1990s and early 2000s, only large web properties had large, web-scale problems. Now, with microservices and digital transformation, every business has a large, web-scale problem—likely multiple large, web-scale problems.

那么，这一切与 Kubernetes 有什么关系呢？一切。世界已经改变。早在 1990 年代末和 2000 年代初，只有大型网络资产才会出现大规模的网络规模问题。现在，通过微服务和数字化转型，每个企业都有一个大型的网络规模问题——可能有多个大型的网络规模问题。

Your business needs to be able to manage a complex web-scale property with many different, often sophisticated services built by many different people. Your web properties need to handle traffic dynamically, and they need to be secure. These properties need to be API-driven at all layers, from the infrastructure to the application layer.

您的企业需要能够管理复杂的网络规模资产，其中包含由许多不同的人构建的许多不同的、通常是复杂的服务。您的网络资产需要动态处理流量，并且它们需要是安全的。这些属性需要在从基础设施到应用程序层的所有层都是 API 驱动的。

## Enter Kubernetes

## 进入 Kubernetes

Kubernetes isn't complex; your business problems are. When you want to run applications in production, there is a minimum level of complexity required to meet the performance (scaling, jitter, etc.) and security requirements. Things like high availability (HA), capacity requirements (N+1, N+2, N+100), and eventually consistent data technologies become a requirement. These are production requirements for every company that has digitally transformed, not just the large web properties like Google, Facebook, and Twitter.

Kubernetes 并不复杂；您的业务问题是。当您想在生产中运行应用程序时，满足性能（缩放、抖动等）和安全要求所需的复杂性最低。诸如高可用性 (HA)、容量要求（N+1、N+2、N+100）和最终一致的数据技术等成为一项要求。这些是每家进行数字化转型的公司的生产要求，而不仅仅是像 Google、Facebook 和 Twitter 这样的大型网络资产。

In the old world, I lived at American Greetings, every time we onboarded a new service, it looked something like this. All of this was handled by the web operations team, and none of it was offloaded to other teams using ticket systems, etc. This was DevOps before there was DevOps:

在旧世界，我住在 American Greetings，每次我们加入一项新服务时，它看起来都像这样。所有这些都由 Web 运营团队处理，并没有使用票证系统等将其卸载给其他团队。 这是 DevOps 之前的 DevOps：

1. Configure DNS (often internal service layers and external public-facing)
2. Configure load balancers (often internal services and public-facing) 
1. Configure shared access to files (large NFS servers, clustered file systems, etc.)
4. Configure clustering software (databases, service layers, etc.)
5. Configure webserver cluster (could be 10 or 50 servers)



1. 配置 DNS（通常是内部服务层和面向公众的外部）
2. 配置负载均衡器（通常是内部服务和面向公众的）
3. 配置文件共享访问（大型NFS服务器、集群文件系统等）
4. 配置集群软件（数据库、服务层等）
5. 配置webserver集群（可以是10或50台服务器）

Most of this was automated with configuration management, but configuration was still complex because every one of these systems and services had different configuration files with completely different formats. We investigated tools like [Augeas](http://augeas.net/) to simplify this but determined that it was an anti-pattern to try and normalize a bunch of different configuration files with a translator.

其中大部分是通过配置管理自动化的，但配置仍然很复杂，因为这些系统和服务中的每一个都有不同的配置文件，格式完全不同。我们调查了像 [Augeas](http://augeas.net/) 这样的工具来简化这一过程，但确定尝试使用翻译器规范化一堆不同的配置文件是一种反模式。

Today with Kubernetes, onboarding a new service essentially looks like:

1. Configure Kubernetes YAML/JSON.
2. Submit it to the Kubernetes API ( **kubectl create -f service.yaml**).

今天使用 Kubernetes，加入新服务本质上是这样的：

1. 配置 Kubernetes YAML/JSON。
2. 提交给Kubernetes API（**kubectl create -f service.yaml**）。

Kubernetes vastly simplifies onboarding and management of services. The service owner, be it a sysadmin, developer, or architect, can create a YAML/JSON file in the Kubernetes format. With Kubernetes, every system and every user speaks the same language. All users can commit these files in the same Git repository, enabling GitOps.

Kubernetes 极大地简化了服务的启动和管理。服务所有者，无论是系统管理员、开发人员还是架构师，都可以创建 Kubernetes 格式的 YAML/JSON 文件。使用 Kubernetes，每个系统和每个用户都使用相同的语言。所有用户都可以在同一个 Git 存储库中提交这些文件，从而启用 GitOps。

Moreover, deprecating and removing a service is possible. Historically, it was terrifying to remove DNS entries, load-balancer entries, web-server configurations, etc. because you would almost certainly break something. With Kubernetes, everything is namespaced, so an entire service can be removed with a single command. You can be much more confident that removing your service won't break the infrastructure environment, although you still need to make sure other applications don't use it (a downside with microservices and function-as-a-service [FaaS]).

此外，可以弃用和删除服务。从历史上看，删除 DNS 条目、负载平衡器条目、Web 服务器配置等是很可怕的，因为您几乎肯定会破坏某些内容。使用 Kubernetes，一切都是命名空间的，因此可以使用单个命令删除整个服务。尽管您仍然需要确保其他应用程序不使用它（微服务和功能即服务 [FaaS] 的缺点），您可以更加确信删除您的服务不会破坏基础设施环境。

## Building, managing, and using Kubernetes

## 构建、管理和使用 Kubernetes

Too many people focus on building and managing Kubernetes instead of using it (see [_Kubernetes is a_ _dump truck_](https://opensource.com/article/19/6/kubernetes-dump-truck)).

太多人专注于构建和管理 Kubernetes，而不是使用它（参见 [_Kubernetes is a_ _dump truck_](https://opensource.com/article/19/6/kubernetes-dump-truck))。

Building a simple Kubernetes environment on a single node isn't markedly more complex than installing a LAMP stack, yet we endlessly debate the build-versus-buy question. It's not Kubernetes that's hard; it's running applications at scale with high availability. Building a complex, highly available Kubernetes cluster is hard because building any cluster at this scale is hard. It takes planning and a lot of software. Building a simple dump truck isn't that complex, but building one that can carry [10 tons of dirt and handle pretty well at 200mph](http://crunchtools.com/kubernetes-10-ton-dump-truck-handles-pretty-well-200-mph/) is complex.

在单个节点上构建简单的 Kubernetes 环境并不比安装 LAMP 堆栈复杂得多，但我们无休止地争论构建与购买的问题。 Kubernetes 并不难；它以高可用性大规模运行应用程序。构建一个复杂的、高度可用的 Kubernetes 集群很困难，因为构建这种规模的任何集群都很困难。这需要规划和大量软件。建造一辆简单的自卸车并不复杂，但建造一辆可以运载 [10 吨泥土并以 200 英里/小时的速度处理得很好](http://crunchtools.com/kubernetes-10-ton-dump-truck-handles-很好，200英里/小时/)很复杂。

Managing Kubernetes can be complex because managing large, web-scale clusters can be complex. Sometimes it makes sense to manage this infrastructure; sometimes it doesn't. Since Kubernetes is a community-driven, open source project, it gives the industry the ability to manage it in many different ways. Vendors can sell hosted versions, while users can decide to manage it themselves if they need to. (But you should question whether you actually need to.)

管理 Kubernetes 可能很复杂，因为管理大型 Web 级集群可能很复杂。有时管理这种基础设施是有意义的；有时不会。由于 Kubernetes 是一个社区驱动的开源项目，它为行业提供了以多种不同方式对其进行管理的能力。供应商可以出售托管版本，而用户可以根据需要自行决定对其进行管理。 （但你应该质疑你是否真的需要。）

Using Kubernetes is the easiest way to run a large-scale web property that has ever been invented. Kubernetes is democratizing the ability to run a set of large, complex web services—like Linux did with Web 1.0.

使用 Kubernetes 是运行已发明的大型网络资产的最简单方法。 Kubernetes 正在使运行一组大型复杂 Web 服务的能力大众化——就像 Linux 对 Web 1.0 所做的那样。

Since time and money is a zero-sum game, I recommend focusing on using Kubernetes. Spend your very limited time and money on [mastering Kubernetes primitives](https://opensource.com/article/19/6/kubernetes-basics) or the best way to handle [liveness and readiness probes](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html) (another example demonstrating that large, complex services are hard). Don't focus on building and managing Kubernetes. A lot of vendors can help you with that.

由于时间和金钱是一个零和游戏，我建议专注于使用 Kubernetes。将您非常有限的时间和金钱花在 [掌握 Kubernetes 原语](https://opensource.com/article/19/6/kubernetes-basics) 或处理 [活性和就绪性探测](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html）（另一个例子证明大型复杂服务很难)。不要专注于构建和管理 Kubernetes。许多供应商可以帮助您解决这个问题。

## Conclusion 

##  结论

I remember troubleshooting countless problems like the one I described at the beginning of this article—NFS in the Linux kernel at that time, our homegrown CFEngine, redirect problems that only surfaced on certain web servers, etc. There was no way a developer could help me troubleshoot any of these problems. In fact, there was no way a developer could even get into the system and help as a second set of eyes unless they had the skills of a senior sysadmin. There was no console with graphics or "observability"—observability was in my brain and the brains of the other sysadmins. Today, with Kubernetes, Prometheus, Grafana, and others, that's all changed.

我记得解决了无数问题，比如我在本文开头描述的问题——当时 Linux 内核中的 NFS、我们自己开发的 CFEngine、仅在某些 Web 服务器上出现的重定向问题等。 开发人员无法提供帮助我对这些问题中的任何一个进行故障排除。事实上，除非开发人员具备高级系统管理员的技能，否则他们根本无法进入系统并提供帮助。没有带有图形或“可观察性”的控制台——可观察性存在于我和其他系统管理员的大脑中。今天，有了 Kubernetes、Prometheus、Grafana 等，这一切都改变了。

The point is:

1. The world is different. All web applications are now large, distributed systems. As complex as AmericanGreetings.com was back in the day, the scaling and HA requirements of that site are now expected for every website.
2. Running large, distributed systems is hard. Period. This is the business requirement, not Kubernetes. Using a simpler orchestrator isn't the answer.

重点是：

1. 世界不一样了。现在所有的 Web 应用程序都是大型的分布式系统。与当年 AmericanGreetings.com 一样复杂，现在每个网站都需要该站点的扩展和 HA 要求。
2. 运行大型分布式系统很困难。时期。这是业务需求，而不是 Kubernetes。使用更简单的编排器不是答案。

Kubernetes is absolutely the simplest, easiest way to meet the needs of complex web applications. This is the world we live in and where Kubernetes excels. You can debate whether you should build or manage Kubernetes yourself. There are plenty of vendors that can help you with building and managing it, but it's pretty difficult to deny that it's the easiest way to run complex web applications at scale. 

Kubernetes 绝对是满足复杂 Web 应用程序需求的最简单、最简单的方法。这是我们生活的世界，也是 Kubernetes 的优势所在。您可以讨论是否应该自己构建或管理 Kubernetes。有很多供应商可以帮助您构建和管理它，但很难否认这是大规模运行复杂 Web 应用程序的最简单方法。

