# Individual Kubernetes Clusters vs. Shared Kubernetes Clusters for Development

# 用于开发的单个 Kubernetes 集群与共享 Kubernetes 集群

Jun 30, 2020  From: https://loft.sh/blog/individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development

Table of Contents

目录

- [Individual Clusters for every developer](http://loft.sh#individual-clusters-for-every-developer)
   - [Advantages](http://loft.sh#advantages)
- [Disadvantages](http://loft.sh#disadvantages)
   - [Use Cases](http://loft.sh#use-cases)
- [Shared Cluster](http://loft.sh#shared-cluster)
   - [Advantages](http://loft.sh#advantages-1)
   - [Disadvantages](http://loft.sh#disadvantages-1)
- [Use Cases](http://loft.sh#use-cases-1)
- [Conclusion](http://loft.sh#conclusion)

- [每个开发者的个人集群](http://loft.sh#individual-clusters-for-every-developer)
  - [优势](http://loft.sh#advantages)
- [缺点](http://loft.sh#disadvantages)
  - [用例](http://loft.sh#use-cases)
- [共享集群](http://loft.sh#shared-cluster)
  - [优势](http://loft.sh#advantages-1)
  - [缺点](http://loft.sh#disadvantages-1)
- [用例](http://loft.sh#use-cases-1)
- [结论](http://loft.sh#conclusion)

If you are on a higher level of cloud-native development (level 2 or higher according to [this article about the typical cloud-native journey of companies](http://loft.sh/blog/the-journey-of-adopting-cloud-native-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development)), the developers in your team need to have a direct access to Kubernetes. The first decision you have to make is if you want to use a local environment (rather level 2) or a remote environment in the cloud (level 3).

如果您处于更高级别的云原生开发（根据[这篇关于公司典型云原生旅程的文章](http://loft.sh/blog/the-journey-of-adopting）的级别 2 或更高级别） -cloud-native-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development))，您团队中的开发人员需要直接访问 Kubernetes。您必须做出的第一个决定是要使用本地环境（相当级别 2)还是云中的远程环境（级别 3)。

While this is not a trivial decision with many pros and cons for either method (as described in my previous post about [local vs. remote clusters for development](http://loft.sh/blog/local-cluster-vs-remote-cluster-for-kubernetes-based-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development)), the advantages of the cloud as [Kubernetes development environment]( http://loft.sh/blog/kubernetes-development-environments-comparison/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) are quite compelling: It is a highly realistic environment for all applications that eventually run in the cloud; it allows for easy replicability of processes and errors; a central control of the dev system is possible; you can use basically infinite computing resources and run even very complex software.

虽然这不是一个微不足道的决定，这两种方法都有很多优点和缺点（如我之前关于 [本地与远程集群开发] 的文章中所述（http://loft.sh/blog/local-cluster-vs-remote -cluster-for-kubernetes-based-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development))，云作为 [Kubernetes开发环境]的优势（ http://loft.sh/blog/kubernetes-development-environments-comparison/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development）非常引人注目：这是一个为最终在云中运行的所有应用程序提供高度逼真的环境；它允许轻松复制过程和错误；开发系统的中央控制是可能的；您可以使用基本上无限的计算资源，甚至可以运行非常复杂的软件。

However, after you chose the cloud environment, you are confronted with the next fundamental decision: Should every developer get an own cluster or should one cluster be shared among the developers.

但是，在您选择了云环境之后，您将面临下一个基本决策：是每个开发人员都拥有自己的集群，还是应该在开发人员之间共享一个集群。

To answer this question, I will describe the advantages and disadvantages of both approaches and their main use cases in this post.

为了回答这个问题，我将在这篇文章中描述这两种方法的优缺点及其主要用例。

## [\#](http://loft.sh\#individual-clusters-for-every-developer) Individual Clusters for every developer

## [\#](http://loft.sh\#individual-clusters-for-every-developer) 每个开发人员的单独集群

The first approach is to give every developer an individual, dedicated cluster (or several clusters, if needed) to work with. This can be done by giving everyone in the team access to the cloud environment and adjust the settings there in a way that they are allowed to create clusters themselves.

第一种方法是为每个开发人员提供一个单独的专用集群（或多个集群，如果需要）来使用。这可以通过让团队中的每个人访问云环境并以允许他们自己创建集群的方式调整那里的设置来完成。

Alternatively, it is also possible that the administrator of the cloud environment creates the clusters and then give the developers access to their individual one. In any case, each developer will have a cluster that only they are working in.

或者，云环境的管理员也可以创建集群，然后让开发人员访问他们的个人集群。无论如何，每个开发人员都将拥有一个只有他们在其中工作的集群。

### [\#](http://loft.sh\#advantages) Advantages

### [\#](http://loft.sh\#advantages) 优点

- **Excellent Isolation.** A clear advantage of having individual clusters for each developer is that their work environments are isolated very well. This means that it is very unlikely that one developer interferes with another so that they can work independently and do not have to fear to break the development system for the whole team. 

- **出色的隔离性。** 为每个开发人员拥有单独的集群的一个明显优势是他们的工作环境隔离得非常好。这意味着一个开发人员不太可能干扰另一个开发人员，以便他们可以独立工作，而不必担心破坏整个团队的开发系统。

- **Full Cluster Access.** Due to the good isolation of the individual work environments, it is not a problem anymore to give the developers full cluster access, so they can configure anything they want or need (you potentially might want to limit the maximum resource consumption though). This gives the engineers the flexibility to modify the configuration of the cluster themselves and they can experiment with different Kubernetes configurations without disturbing anyone.

- **完全集群访问权限。**由于各个工作环境的良好隔离，给开发人员完全集群访问权限不再是问题，因此他们可以配置他们想要或需要的任何东西（您可能想要限制最大的资源消耗）。这使工程师可以灵活地自己修改集群的配置，他们可以在不打扰任何人的情况下试验不同的 Kubernetes 配置。

- **Easy Start.** Another benefit of the one-cluster-per-developer method is that it can be introduced in a team fairly fast and easy. The simplest option would be to just share the cloud account credentials and password with everyone. This, however, is of course very insecure and should only be done in very small teams with very trustable members if at all. A better solution is to give everybody their own cloud account with appropriate limitations or the clusters are simply created and then distributed by the cloud manager.

- **容易开始。** 每个开发人员一个集群方法的另一个好处是它可以相当快速和容易地引入团队。最简单的选择是与所有人共享云帐户凭据和密码。然而，这当然是非常不安全的，并且应该只在具有非常可信赖的成员的非常小的团队中进行。更好的解决方案是为每个人提供具有适当限制的自己的云帐户，或者简单地创建集群，然后由云管理器分发。

## [\#](http://loft.sh\#disadvantages) Disadvantages

## [\#](http://loft.sh\#disadvantages) 缺点

- **Problematic in Private Clouds.** A first limitation of giving individual clusters to engineers is that it does not work well in most private clouds. While public clouds such as AWS and Azure provide an advanced user management, most private clouds lack a wide range of user settings and limitations. The alternative solution to give everyone admin access to the cloud environment is usually too dangerous and having one cloud administrator care for the provisioning of the clusters creates a problematic bottleneck.

- **私有云中的问题。** 将单个集群提供给工程师的第一个限制是它在大多数私有云中无法正常工作。虽然 AWS 和 Azure 等公共云提供了高级用户管理，但大多数私有云缺乏广泛的用户设置和限制。为每个人提供对云环境的管理员访问权限的替代解决方案通常过于危险，并且让一名云管理员负责集群的配置会产生问题瓶颈。

- **Kubernetes Knowledge and Tools Necessary.** If a developer is supposed to fully interact with Kubernetes and has to create and configure their own cluster, they need to have some Kubernetes knowledge and tools such as kubectl. Although it is possible to streamline these processes to some extent with detailed instructions for the tool installation and step-by-step guides for the cluster setup, everyone in a team has to repeat the same steps (which, by the way, eliminates [one of the advantages of using the cloud instead of local Kubernetes environments](http://loft.sh/blog/local-cluster-vs-remote-cluster-for-kubernetes-based-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development)).

- **Kubernetes 知识和工具必备。** 如果开发人员要与 Kubernetes 完全交互并且必须创建和配置自己的集群，他们需要具备一些 Kubernetes 知识和工具，例如 kubectl。尽管可以通过工具安装的详细说明和集群设置的分步指南在一定程度上简化这些过程，但团队中的每个人都必须重复相同的步骤（顺便说一下，这消除了 [一个使用云而不是本地 Kubernetes 环境的优势](http://loft.sh/blog/local-cluster-vs-remote-cluster-for-kubernetes-based-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development）)。

- **Maintenance Effort.** After a successful setup of the Kubernetes clusters, the work is not done. They rather have to be continuously maintained and updated. Since the individual cluster approach can lead to a lot of clusters depending on the size of the engineering team, the maintenance effort becomes immense for dozens or even hundreds of clusters.

- **维护工作量。** Kubernetes 集群设置成功后，工作尚未完成。他们宁愿必须不断维护和更新。由于单个集群方法可能会根据工程团队的规模而导致大量集群，因此对于数十甚至数百个集群的维护工作变得巨大。

- **Limited Oversight and Central Control.** Due to the abundance of clusters, it can be very hard to supervise the whole development system in the cloud. While it should be possible to see how much resources are consumed and what clusters are running, it is challenging to see what is currently needed and what is running idle and could be deleted. The complexity of managing the whole environment so becomes very high, especially for larger teams. 

- **有限的监督和中央控制。**由于集群的丰富，在云中监督整个开发系统可能非常困难。虽然应该可以查看消耗了多少资源以及正在运行哪些集群，但要查看当前需要什么以及空闲运行和可能删除的资源具有挑战性。管理整个环境的复杂性因此变得非常高，尤其是对于大型团队而言。

- **Computing Cost.** A general downside of using a cloud environment for development is that it can be costly, especially in public clouds. Since a system with a lot of clusters also creates a lot of redundancy for basic Kubernetes elements such as API servers, the one-cluster-per-developer approach is particularly expensive. To make things worse, some providers such as AWS and GCP charge a cluster management fee for every cluster, which is very costly if you use many small clusters as you usually would for development. (Google Cloud has only [recently announced the introduction of a cluster management fee](https://cloud.google.com/kubernetes-engine/pricing) creating [some outrage by its users](https://news.ycombinator.com/item?id=22486070).)

- **计算成本。** 使用云环境进行开发的一个普遍缺点是成本高昂，尤其是在公共云中。由于具有大量集群的系统还会为 API 服务器等基本 Kubernetes 元素创建大量冗余，因此每个开发人员一个集群的方法特别昂贵。更糟糕的是，AWS 和 GCP 等一些提供商会为每个集群收取集群管理费，如果您像往常一样使用许多小型集群进行开发，这将非常昂贵。 （谷歌云只是[最近宣布引入集群管理费](https://cloud.google.com/kubernetes-engine/pricing)引起了[用户的愤怒](https://news.ycombinator.com/item?id=22486070).)

### [\#](http://loft.sh\#use-cases) Use Cases

### [\#](http://loft.sh\#use-cases) 用例

There are some situations in which the individual-cluster approach for providing Kubernetes access to developer makes more sense than in others:

在某些情况下，为开发人员提供 Kubernetes 访问权限的单个集群方法比其他情况更有意义：

- **Public Clouds.** Generally, due to the extensive user management in most public clouds, this approach works better with public cloud environments than private clouds where only limited user management is available.

- **公共云。** 通常，由于大多数公共云中的广泛用户管理，这种方法更适用于公共云环境，而不是只有有限的用户管理可用的私有云。

- **Teams of Kubernetes Experts.** In teams of Kubernetes experts, the downsides of individual setup and maintenance are less important, and the great flexibility can be a positive factor for this solution. There are even special cases in which this is the only possible approach for cloud-based development, such as if a solution or tool for Kubernetes management itself is developed.

- **Kubernetes 专家团队。** 在 Kubernetes 专家团队中，个人设置和维护的缺点不太重要，而极大的灵活性可能是该解决方案的一个积极因素。甚至在某些特殊情况下，这是基于云的开发的唯一可能方法，例如，如果开发了用于 Kubernetes 管理本身的解决方案或工具。

> _This is of course rather an edge case, but in my company, we are actually developing a [solution to transform clusters into a Kubernetes multi-tenancy platform](https://loft.sh) and our engineers thus often need full cluster access._

> _这当然是一个边缘案例，但在我的公司，我们实际上正在开发[将集群转换为 Kubernetes 多租户平台的解决方案](https://loft.sh)，因此我们的工程师经常需要完整的集群使用权。_

- **Small Teams.** In smaller teams, only a few clusters are needed for this approach and generally, it is easier to set-up and manage fewer clusters. Additionally, it is possible to guide the developers through the set-up process individually and to communicate informally and fast if the team size is relatively small.

- **小团队。**在较小的团队中，这种方法只需要几个集群，通常更容易设置和管理更少的集群。此外，如果团队规模相对较小，可以单独指导开发人员完成设置过程并进行非正式和快速的沟通。

- **No Budget Constrains.** While there eventually will be budget constraints, the cloud cost is not always an important concern. This could be as the absolute cost is negligible or because the company has received credits for a public cloud, which many startups have. In these cases, there might be more important things to optimize than the cloud cost. (In a separate blog post, I describe how you [can save cost by reducing the number of clusters](http://loft.sh/blog/kubernetes-cost-savings/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development).)

- **没有预算限制。**虽然最终会有预算限制，但云成本并不总是一个重要的问题。这可能是因为绝对成本可以忽略不计，或者因为该公司已经获得了许多初创公司拥有的公共云的信用。在这些情况下，可能有比云成本更重要的事情需要优化。 （在另一篇博文中，我描述了如何[通过减少集群数量来节省成本](http://loft.sh/blog/kubernetes-cost-savings/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development）。)

- **First experiments.** Since you can get started pretty fast with this approach, it can often give you a good first feeling about how cloud environments and cloud-native development with Kubernetes works.

- **第一次实验。**由于您可以很快地开始使用这种方法，因此它通常可以让您对云环境和使用 Kubernetes 进行云原生开发的工作原理有一个很好的第一感觉。

Since the approach to give every developer an individual cluster is very inefficient regarding computing resources and management effort, it is rarely used in reality. Even though there are some use cases in small startups or in companies that develop special Kubernetes solutions, I believe that this solution is generally inferior for the average company compared to the approach I will describe next.

由于为每个开发人员提供一个单独的集群的方法在计算资源和管理工作方面效率非常低，因此在现实中很少使用。尽管在小型初创公司或开发特殊 Kubernetes 解决方案的公司中存在一些用例，但我相信与我接下来将描述的方法相比，该解决方案对于普通公司而言通常较差。

## [\#](http://loft.sh\#shared-cluster) Shared Cluster

## [\#](http://loft.sh\#shared-cluster) 共享集群

In the shared cluster approach, as the name suggests, developers share one or only a few clusters. This implies that there are several users of the same cluster sharing its resources.

在共享集群方法中，顾名思义，开发人员共享一个或几个集群。这意味着同一个集群的多个用户共享其资源。

In theory, it would be possible to give everyone full and direct cluster access, but this would easily end in total chaos because such a system would be very vulnerable to small mistakes. For example, someone could consume all available computing resources or change the cluster configuration, so it crashes. 

理论上，可以为每个人提供完全和直接的集群访问权限，但这很容易导致完全混乱，因为这样的系统很容易出现小错误。例如，有人可能会消耗所有可用的计算资源或更改集群配置，因此它会崩溃。

For this, the cluster needs to be prepared in a way that it securely supports multi-tenancy. In Kubernetes, namespaces play an important role in this approach. However, a full shared cluster solution goes beyond this and rather means a self-service, multi-tenancy Kubernetes system that allows the developers to create namespaces on demand, while it also cares for fair usage sharing by limitations and ensures the users' namespace isolation .

为此，需要以安全支持多租户的方式准备集群。在 Kubernetes 中，命名空间在这种方法中发挥着重要作用。然而，完整的共享集群解决方案超越了这一点，而是一个自助服务、多租户的 Kubernetes 系统，它允许开发人员按需创建命名空间，同时它也关心公平的使用共享，并确保用户的命名空间隔离。 .

### [\#](http://loft.sh\#advantages-1) Advantages

### [\#](http://loft.sh\#advantages-1) 优点

- **Any Cloud Environment.** One advantage of the shared cluster approach is that it can be used in any cloud, public and private, as long as this cloud is running Kubernetes. In contrast to the one-cluster-per-developer approach, there are no downsides associated with either public or private clouds and it is even possible to use hybrid systems.

- **任何云环境。** 共享集群方法的一个优点是它可以在任何云中使用，无论是公共云还是私有云，只要该云运行 Kubernetes。与每个开发人员一个集群的方法相比，公共云或私有云都没有缺点，甚至可以使用混合系统。

- **Cost-Efficiency.** While there will always be some cost associated with using the cloud as development environment, using shared clusters is at least a computing resource and thus cost-efficient solution. All basic Kubernetes functionality is shared to avoid any redundancy and the individual namespaces can sometimes even be automatically shut-off. This can [reduce the Kubernetes computing cost](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) dramatically. For example, if namespaces are always scaled down or deactivated when the developer is not working (at night, during weekends or holidays), this can often save [70% of the total computing cost of the development system](http://loft .sh/blog/how-to-save-more-than-2-3-of-engineers-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for- development).

- **成本效益。** 虽然使用云作为开发环境总会产生一些成本，但使用共享集群至少是一种计算资源，因此是一种具有成本效益的解决方案。所有基本的 Kubernetes 功能都是共享的，以避免任何冗余，并且有时甚至可以自动关闭各个命名空间。这样可以 [降低 Kubernetes 计算成本](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) 戏剧性的。例如，如果在开发人员不工作时（晚上、周末或节假日)总是缩小或停用命名空间，这通常可以节省 [70% 的开发系统总计算成本](http://loft.sh/blog/how-to-save-more-than-2-3-of-engineers-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-发展)。

- **Central Control.** Since there is only one or very few clusters, it is possible to monitor the whole system relatively easy. Of course, you still need a user management and must enforce appropriate limits for the developers using the cluster. However, there are already solutions such as [Loft](https://loft.sh) that offer all of this together with a dashboard to make sense of who uses what and what limits are enforced. This also makes it possible to shut-off individual namespaces automatically or manually from a central position.

- **中央控制。**由于只有一个或很少的集群，因此可以相对容易地监控整个系统。当然，您仍然需要用户管理，并且必须对使用集群的开发人员实施适当的限制。但是，已经有诸如 [Loft](https://loft.sh) 之类的解决方案提供所有这些以及仪表板，以了解谁使用什么以及强制执行什么限制。这也使得可以从中心位置自动或手动关闭各个命名空间。

- **Only 1 System.** From a maintenance perspective, it is much easier to manage only one larger cluster than a large number of very small ones. The cluster administrator only has to update and configure one cluster instead of repeating this process over and over for each cluster individually, which saves a lot of time and effort.

- **只有 1 个系统。** 从维护的角度来看，只管理一个较大的集群比管理大量非常小的集群要容易得多。集群管理员只需要更新和配置一个集群，而不是为每个集群单独重复这个过程，这样可以节省大量的时间和精力。

- **Ease of Use.** The developers can easily work with a system that allows on-demand namespace generation. From their perspective, the whole system will feel like a normal PaaS-system and Kubernetes is mostly a technology running in the background that they do not always have to directly interact with. For this, little to no Kubernetes knowledge and hardly any new tools are needed on the developer’s side.

- **易于使用。** 开发人员可以轻松地使用允许按需生成命名空间的系统。从他们的角度来看，整个系统感觉就像一个普通的 PaaS 系统，而 Kubernetes 主要是一种在后台运行的技术，他们并不总是需要直接与之交互。为此，开发人员几乎不需要 Kubernetes 知识，也几乎不需要任何新工具。

### [\#](http://loft.sh\#disadvantages-1) Disadvantages 

### [\#](http://loft.sh\#disadvantages-1) 缺点

- **Isolation Necessary.** Using shared clusters to give developers access to Kubernetes, however, also comes with some disadvantages. The first one is that you need a solution for isolating the individual developers from each other to prevent chaos in your cluster. While namespaces are a good starting point and [special solutions for this problem exist](https://github.com/kiosk-sh/kiosk), user isolation in Kubernetes needs to be taken care of and can be tricky, especially if you require hard multi-tenancy. Fortunately, in most development settings, the tenants on the same cluster can be trusted and thus soft multi-tenancy is sufficient. Still, you need to introduce Network Policies to prevent namespaces to communicate with other namespaces they should not be able to talk to.

- **需要隔离。** 然而，使用共享集群让开发人员能够访问 Kubernetes 也有一些缺点。第一个是您需要一个解决方案来隔离各个开发人员，以防止集群中的混乱。虽然命名空间是一个很好的起点并且 [存在针对此问题的特殊解决方案](https://github.com/kiosk-sh/kiosk)，但 Kubernetes 中的用户隔离需要注意并且可能会很棘手，尤其是如果您需要硬多租户。幸运的是，在大多数开发环境中，同一集群上的租户是可以信任的，因此软多租户就足够了。尽管如此，您仍需要引入网络策略来防止命名空间与它们不应该与之通信的其他命名空间进行通信。

- **User Limitation and Management Necessary.** Besides the technical limitation between the users, you also need a limitation of the individual users, so they cannot excessively use computing resources and network traffic. This leads to an additional requirement of a solid user management system, which is also necessary to give every developer access to the cluster in the first place. Again, such [multi-tenancy platform solutions](https://loft.sh) exist, but if you want to build the system from scratch on your own, this is an additional topic to consider. (In a separate article, I describe [how to build such an internal Kubernetes platform](http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development).)

- **需要用户限制和管理。**除了用户之间的技术限制外，您还需要对单个用户进行限制，使他们不能过度使用计算资源和网络流量。这导致对可靠的用户管理系统的额外要求，这也是让每个开发人员首先访问集群的必要条件。同样，存在这样的[多租户平台解决方案](https://loft.sh)，但如果您想自己从头开始构建系统，这是一个需要考虑的额外话题。 （在另一篇文章中，我描述了[如何构建这样一个内部 Kubernetes 平台](http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development）。)

- **Setup Effort.** In any case, you will face some setup effort to get a shared Kubernetes cluster running and to configure it properly. Even with off-the-shelf solutions, you need to set them up at the beginning, which is at least some more effort than the easiest one-cluster-per-developer-approach. Still, this one-time effort can be made up by the reduced maintenance work later on.

- **设置工作。**在任何情况下，您都将面临一些设置工作，以使共享的 Kubernetes 集群运行并正确配置它。即使使用现成的解决方案，您也需要在开始时进行设置，这至少比最简单的每个开发人员一个集群的方法要多一些努力。不过，这种一次性的努力可以通过以后减少的维护工作来弥补。

- **Limited Flexibility for Special Cases.** Due to the limits imposed on the users’ access to the cluster, e.g. they can only use a specific number of namespaces or change only some Kubernetes settings, the developers do not have the same freedom they have with an own cluster. Since the limitations should usually be set in a smart way with the normal requirements of your team in mind, this should only become a problem in special cases. Then, the cluster administrator needs to give them additional rights, which is possible but can slow down the development process somewhat.

- **特殊情况的灵活性有限。**由于对用户访问集群的限制，例如他们只能使用特定数量的命名空间或仅更改一些 Kubernetes 设置，开发人员没有与自己的集群相同的自由。由于通常应该根据团队的正常要求以一种聪明的方式设置限制，这应该只在特殊情况下成为问题。然后，集群管理员需要给他们额外的权限，这是可能的，但会在一定程度上减慢开发过程。

> The limited flexibility issue could be resolved by using [virtual Kubernetes clusters](http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) instead of namespaces. If every developer had a virtual cluster, it would be possible to even change cluster-wide resources in that virtual cluster and still not interfere with other tenants. There are some open-source concepts for virtual clusters, e.g. by [Alibaba](https://www.cncf.io/blog/2019/06/20/virtual-cluster-extending-namespace-based-multi-tenancy-with-a-cluster-view/) or by [Darren Shepherd](https://github.com/ibuildthecloud/k3v). Additionally, [Loft](https://loft.sh) has alread integrated virtual clusters as alternative to namespaces into its Kubernetes platform.

> 有限的灵活性问题可以通过使用[虚拟 Kubernetes 集群](http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters)来解决-vs-shared-kubernetes-clusters-for-development)而不是命名空间。如果每个开发人员都有一个虚拟集群，甚至可以在该虚拟集群中更改集群范围的资源，并且仍然不会干扰其他租户。虚拟集群有一些开源概念，例如[阿里巴巴](https://www.cncf.io/blog/2019/06/20/virtual-cluster-extending-namespace-based-multi-tenancy-with-a-cluster-view/) 或 [Darren牧羊人](https://github.com/ibuildthecloud/k3v)。此外，[Loft](https://loft.sh) 已将虚拟集群作为命名空间的替代方案集成到其 Kubernetes 平台中。

## [\#](http://loft.sh\#use-cases-1) Use Cases

## [\#](http://loft.sh\#use-cases-1) 用例

Like the one-cluster-per-developer approach, the shared cluster solution also is more appropriate in some cases than in others, but generally, it can be applied in a broader range of settings:

与每个开发人员一个集群的方法一样，共享集群解决方案在某些情况下也比在其他情况下更合适，但通常，它可以应用于更广泛的设置：

- **Any Cloud.** Sharing a Kubernetes cluster is possible in both public and private clouds, which makes it a universal solution from a purely technical perspective. 

- **任何云。**在公共云和私有云中共享一个 Kubernetes 集群是可能的，这使得它从纯粹的技术角度来看是一个通用的解决方案。

- **Teams without Kubernetes Knowledge.** Since the shared cluster will be centrally managed by an experienced cluster admin, the average developer does not need extensive knowledge about Kubernetes. They potentially might need to install a few tools to be able to deploy to the cluster and to create namespaces in it. However, these processes are fairly easy and standardized.

- **没有 Kubernetes 知识的团队。** 由于共享集群将由经验丰富的集群管理员集中管理，因此普通开发人员不需要广泛的 Kubernetes 知识。他们可能需要安装一些工具才能部署到集群并在其中创建命名空间。然而，这些过程相当简单和标准化。

- **Trusted Team Members.** One of the biggest technical challenges with the shared cluster approach is to establish multi-tenancy. Here, it is much easier to realize soft multi-tenancy, i.e. to isolate users from each other in a way that lets them work without disturbing each other but does not securely prevents malicious behavior. For most development teams, this isolation should be sufficient as all tenants are known and have no intention to destroy anything on purpose.

- **值得信赖的团队成员。** 共享集群方法的最大技术挑战之一是建立多租户。在这里，实现软多租户要容易得多，即将用户彼此隔离，让他们在不干扰彼此的情况下工作，但不能安全地防止恶意行为。对于大多数开发团队来说，这种隔离应该足够了，因为所有租户都是已知的，并且无意故意破坏任何东西。

- **Any Team Size.** The marginal effort to add an additional developer to a shared cluster is very low. Therefore, it is possible to use shared clusters even with large teams. On the other side, it is also possible to use it with very small teams of only 2-5 developers or even alone (e.g. if you want to have separate systems for development, staging/testing and production), but here, you need to trade off the benefits of a shared cluster against the effort to set it up.

- **任何团队规模。** 将额外的开发人员添加到共享集群的边际努力非常低。因此，即使是大型团队也可以使用共享集群。另一方面，也可以与只有 2-5 名开发人员的非常小的团队一起使用它，甚至可以单独使用它（例如，如果你想有单独的系统进行开发、登台/测试和生产），但在这里，你需要权衡共享集群的好处和建立它的努力。

From my perspective, it looks like many professional teams that have decided to provide developers access to Kubernetes made their decision for shared cluster solutions. As an example, there was a talk about [Spotify’s](https://www.youtube.com/watch?v=vLrxOhZ6Wrg) solution for this at KubeCon North America 2019. They have developed a development platform solution for Kubernetes internally.

从我的角度来看，似乎许多决定为开发人员提供 Kubernetes 访问权限的专业团队决定采用共享集群解决方案。例如，在 KubeCon North America 2019 上有一个关于 [Spotify’s](https://www.youtube.com/watch?v=vLrxOhZ6Wrg) 解决方案的讨论。他们在内部为 Kubernetes 开发了一个开发平台解决方案。

If you also plan to go this way, you should take a look at open source solutions such as [kiosk](https://github.com/kiosk-sh/kiosk) that solve many issues regarding [Kubernetes multi-tenancy]( http://loft.sh/blog/kubernetes-multi-tenancy-best-practices-guide/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) and can be used as building block for more comprehensive solutions, such as [self-service namespace platforms](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) (kiosk, for example, does not provide a user management).

如果你也打算这样做，你应该看看开源解决方案，例如 [kiosk](https://github.com/kiosk-sh/kiosk)，它解决了许多关于 [Kubernetes 多租户] 的问题（ http://loft.sh/blog/kubernetes-multi-tenancy-best-practices-guide/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development）并且可以用作更全面解决方案的构建块，例如 [自助命名空间平台](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source= other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development）（例如，kiosk 不提供用户管理)。

However, there are also commercial, off-the-shelf solutions such as [Loft](https://loft.sh) that transform any Kubernetes cluster into a multi-tenancy enabled development platform with integrated user management, extensive user limit settings and an automatic computing resource saving sleep-mode for namespace.

然而，也有商业的、现成的解决方案，例如 [Loft](https://loft.sh)，可以将任何 Kubernetes 集群转变为支持多租户的开发平台，具有集成的用户管理、广泛的用户限制设置和命名空间的自动计算资源节省睡眠模式。

## [\#](http://loft.sh\#conclusion) Conclusion

## [\#](http://loft.sh\#conclusion) 结论

Integrating the cloud environment in the development workflow early on is still a relatively new concept as nowadays cloud technologies often only come into play at the staging and testing level. However, with the advent of Kubernetes in more and more companies, one can expect that it will be more common in the future to also give developers direct access to Kubernetes and then you need to decide if you want to use one cluster per developer or a single shared cluster for the whole team.

早期将云环境集成到开发工作流程中仍然是一个相对较新的概念，因为如今云技术通常只在暂存和测试级别发挥作用。然而，随着 Kubernetes 在越来越多的公司中出现，人们可以期待在未来让开发人员直接访问 Kubernetes 会更加普遍，然后您需要决定是每个开发人员使用一个集群还是使用一个集群。整个团队的单个共享集群。

To get the benefits of a real Kubernetes environment in the cloud for development, both approaches are imaginable and have their specific use cases. However, from my perspective, the shared cluster approach is clearly superior for all common settings while individual clusters for developers seem to be more a solution for first experiments or niche requirements. 

为了在云中获得真正的 Kubernetes 环境进行开发的好处，这两种方法都是可以想象的，并且都有其特定的用例。但是，从我的角度来看，共享集群方法对于所有常见设置显然更胜一筹，而对于开发人员来说，单独的集群似乎更像是第一次实验或小众需求的解决方案。

Overall, you often do not have to make a strict decision between the solutions. So, it might be possible that some engineers working on a specific task get their own cluster while the rest of the team works in a shared cluster environment. You might even include local environments in [the Kubernetes development workflow](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) and use a shared cluster in the cloud only for testing and staging. As you can see, the opportunities seem endless, but I hope that this post was helpful to give you some guidance on what is possible and when to use which solution for truly cloud-native development. 

总体而言，您通常不必在解决方案之间做出严格的决定。因此，一些从事特定任务的工程师可能拥有自己的集群，而团队的其他成员则在共享集群环境中工作。您甚至可以在 [Kubernetes 开发工作流程](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs)中包含本地环境-shared-kubernetes-clusters-for-development)并在云中使用共享集群仅用于测试和登台。正如您所看到的，机会似乎无穷无尽，但我希望这篇文章能够为您提供一些指导，帮助您了解什么是可能的以及何时使用哪种解决方案进行真正的云原生开发。

