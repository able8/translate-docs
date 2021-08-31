# Engineering dependability and fault tolerance in a distributed system

# 分布式系统中的工程可靠性和容错性

[![](https://ik.imagekit.io/ably/ghost/prod/2019/05/paddy-byers-83e6fa6a5d47687601b0586d529a4619bbe68178b5ff4d1ba8636e4612ea716f.jpg?tr=w-300)](http://ably.com#footer-bio)

-生物）

By: Paddy Byers
Last updated: Aug 18, 202116 min read

作者：帕迪拜尔斯
最后更新时间：202116 年 8 月 18 日分钟阅读

Users need to know that they can depend on the service that is provided to them. In practice, because from time to time individual elements will inevitably fail, this means you have to be able to continue in spite of those failures.

用户需要知道他们可以依赖提供给他们的服务。在实践中，由于个别元素有时会不可避免地失败，这意味着即使出现这些失败，您也必须能够继续。

In this article, we discuss the concepts of dependability and fault tolerance in detail and explain how the Ably platform is designed with fault tolerant approaches to uphold its dependability guarantees.

在本文中，我们详细讨论了可靠性和容错的概念，并解释了 Ably 平台是如何使用容错方法设计的，以维护其可靠性保证。

As a basis for that discussion, first some definitions:

作为讨论的基础，首先给出一些定义：

> **Dependability**
>
> The degree to which a product or service can be relied upon. **Availability** and **Reliability** are forms of dependability.
>
> **Availability**
>
> The degree to which a product or service is **available** for use when required. This often boils down to provisioning sufficient redundancy of resources with [statistically independent failures](https://www.ibm.com/support/knowledgecenter/STXNRM_3.15.1/coss.doc/correlatedFailureMitigation_defining_independent_versus_correlated_failures.html).
>
> **Reliability**
>
> The degree to which the product or service **conforms to its specification** when in use. This means a system that is not merely available but is also engineered with extensive redundant measures to **continue to work as its users expect**.
>
> **Fault tolerance**
>
> The ability of a system to continue to be **dependable** (both available and reliable) in the presence of certain component or subsystem failures.

> **可靠性**
>
> 可以依赖产品或服务的程度。 **可用性**和**可靠性**是可靠性的形式。
>
> **可用性**
>
> 产品或服务在需要时**可用**的程度。这通常归结为使用 [统计独立故障] (https://www.ibm.com/support/knowledgecenter/STXNRM_3.15.1/coss.doc/correlatedFailureMitigation_defining_independent_versus_correlated_failures.html) 提供足够的资源冗余。
>
> **可靠性**
>
> 产品或服务在使用时**符合其规格**的程度。这意味着系统不仅可用，而且还设计有广泛的冗余措施，以**继续按用户期望的方式工作**。
>
> **容错**
>
> 系统在出现某些组件或子系统故障时继续保持**可靠**（可用和可靠）的能力。

Fault tolerant systems tolerate faults: they’re designed to mitigate the impact of adverse circumstances and ensure the system remains dependable to the end-user. Fault-tolerance techniques can be used to improve both availability and reliability.

容错系统可以容忍错误：它们旨在减轻不利环境的影响并确保系统对最终用户保持可靠。容错技术可用于提高可用性和可靠性。

Availability can be loosely thought of as the assurance of uptime; reliability can be thought of as the quality of that uptime — that is, assurance that functionality and user experience and preserved as effectively as possible in spite of adversity.

可用性可以被松散地认为是正常运行时间的保证；可靠性可以被认为是正常运行时间的质量——即保证功能和用户体验，并在逆境中尽可能有效地保留。

If the service isn’t available to be used at the time it is needed, that’s a shortfall in availability. If the service is available, but deviates from its expected behavior when you use it, that’s a shortfall in reliability. Fault tolerant design approaches address these shortfalls to provide continuity both to business and to the user experience.

如果该服务在需要时无法使用，那就是可用性不足。如果该服务可用，但在您使用它时偏离了其预期行为，那么这就是可靠性的不足。容错设计方法解决了这些不足，以提供业务和用户体验的连续性。

> Read the latest Engineering post: [Balancing act: the current limits of AWS network load balancers](https://ably.com/blog/limits-aws-network-load-balancers)

> 阅读最新的工程帖子：[平衡法：AWS 网络负载均衡器的当前限制](https://ably.com/blog/limits-aws-network-load-balancers)

## Availability, reliability, and state

## 可用性、可靠性和状态

In most cases, the primary basis for fault tolerant design is **redundancy**: having more than the minimum number of components or capacity required to deliver service. The key questions relate to what form that redundancy takes and how it is managed.

在大多数情况下，容错设计的主要基础是**冗余**：拥有超过提供服务所需的最少数量的组件或容量。关键问题与冗余采用什么形式以及如何管理有关。

In the physical world there is classically a distinction between

在物理世界中，经典上存在以下区别：

- **availability** situations, where it is acceptable to stop a service and then resume it, such as stopping to change a car tire; and
- **reliability** situations, where continuity of service is essential, with redundant elements continuously in-service, such as with airplane engines.

- **可用性** 情况，停止服务然后恢复是可以接受的，例如停下来更换汽车轮胎；和
- **可靠性** 情况，在这种情况下，服务的连续性是必不可少的，冗余元素持续服务，例如飞机发动机。

The nature of the continuity requirement impacts the way that redundant capacity needs to be provided.

连续性要求的性质影响需要提供冗余容量的方式。

In the context of distributed systems, at Ably we think of an analogous distinction between components that are **stateless** and those that are **stateful**, respectively.

在分布式系统的上下文中，在 Ably，我们分别考虑了**无状态**组件和**有状态**组件之间的类似区别。

**Stateless components** fulfil their function without any dependency on long-lived state. Each invocation of service can be performed independently of any previous invocation. Fault tolerant design for these components is comparatively straightforward: have sufficient resources **available** so that any individual invocation can be handled even if some of the resources have failed.

**无状态组件**在不依赖长期状态的情况下完成其功能。服务的每次调用都可以独立于任何先前的调用执行。这些组件的容错设计相对简单：拥有足够的资源**可用**，以便即使某些资源失败也可以处理任何单个调用。

**Stateful components** have an intrinsic dependency on state to provide service. Having state implicitly links an invocation of the service to past and future invocations. The essence of fault tolerance for these components is, as with an airplane's engines, about being able to provide continuity of operation — specifically, continuity of the state that the service depends on. This ensures **reliability**.

**有状态组件** 对提供服务的状态具有内在依赖性。拥有状态隐式地将服务的调用与过去和未来的调用联系起来。与飞机发动机一样，这些组件容错的本质在于能够提供操作的连续性——特别是服务所依赖的状态的连续性。这确保了**可靠性**。

In the remainder of this article we will give examples of each of these situations and explain the engineering challenges encountered in achieving fault tolerance in practice.

在本文的其余部分，我们将给出每种情况的示例，并解释在实践中实现容错时遇到的工程挑战。

## Treating failure as a matter of course 
## 把失败当成理所当然
Fault tolerant designs treat failures as routine. In large-scale systems, the assumption has to be that component failures will happen sooner or later. Any individual failure must be assumed as imminent and, collectively, component failures must be expected to be occurring continuously.

容错设计将故障视为例行程序。在大型系统中，必须假设组件故障迟早会发生。任何单个故障都必须假设为迫在眉睫，并且总体而言，必须预期组件故障将持续发生。

By contrast with the physical world, failures in digital systems are typically non-binary. The classical measures of component reliability (e.g. Mean Time Between Failures, or MTBF) do not apply; services degrade along a gradient of failure. Byzantine faults are a classic example.

与物理世界相比，数字系统中的故障通常是非二进制的。组件可靠性的经典度量（例如平均故障间隔时间或 MTBF）不适用；服务会随着故障梯度下降。拜占庭断层就是一个典型的例子。

For example, a system component might work intermittently, or produce misleading output. Or you might be dependent on external partners who don’t notify you of a failure until it becomes serious on their end, making your work more difficult.

例如，系统组件可能会间歇性地工作，或产生误导性的输出。或者，您可能依赖于外部合作伙伴，他们在失败变得严重之前不会通知您失败，从而使您的工作更加困难。

Being tolerant to non-binary failures requires a lot of thought, engineering and, sometimes, human intervention. Each potential failure must be identified and classified, and must then be capable of being remediated rapidly, or avoided through extensive testing and robust design decisions. The core challenge of designing fault-tolerant systems is understanding the nature of failures and how they can be detected and remediated, especially when partial or intermittent, to continue to provide the service as effectively as possible for users.

容忍非二进制故障需要大量的思考、工程，有时还需要人工干预。每个潜在的故障都必须被识别和分类，然后必须能够快速修复，或者通过广泛的测试和稳健的设计决策来避免。设计容错系统的核心挑战是了解故障的性质以及如何检测和修复故障，尤其是在部分故障或间歇性故障时，以继续尽可能有效地为用户提供服务。

## Stateless services

##无状态服务

Service layers that are **stateless** do not have a primary requirement for continuity of service for any individual component. The availability of resources directly translates into availability of the layer as a whole. Having access to extra resources whose failures are statistically independent is key to keeping the system going. Wherever possible, layers are designed to be stateless as a key enabler not only of availability, but also of scalability.

**无状态**的服务层对任何单个组件的服务连续性没有主要要求。资源的可用性直接转化为整个层的可用性。访问故障在统计上独立的额外资源是保持系统运行的关键。在可能的情况下，层被设计为无状态的，不仅是可用性的关键推动因素，也是可扩展性的关键推动因素。

For stateless objects, it suffices to have multiple and independently available components to continue to provide service. Without state, durability of any single component is not a concern.

对于无状态对象，拥有多个独立可用的组件就足以继续提供服务。没有状态，任何单个组件的耐用性都不是问题。

![Fault tolerant design of stateless components: a load balancer receives a request and selects an element to satisfy the request.](https://ik.imagekit.io/ably/ghost/prod/2021/02/fault-tolerance-of-a-stateless-component.png?tr=w-1520)

of-a-stateless-component.png?tr=w-1520)

However, simply having extra resources is not enough: you also have to use them effectively. You have to have a way of detecting resource availability, and to load-balance among redundant resources.

然而，仅仅拥有额外资源是不够的：您还必须有效地使用它们。您必须有一种方法来检测资源可用性，并在冗余资源之间进行负载平衡。

As such, the questions that you have to answer are:

因此，您必须回答的问题是：

- How do you survive different kinds of failure?
- What level of redundancy is possible?
- What is the resource and performance cost of maintaining those levels of redundancy?
- What is the operational cost of managing those levels of redundancy?

- 你如何在不同类型的失败中幸存下来？
- 什么级别的冗余是可能的？
- 维持这些冗余级别的资源和性能成本是多少？
- 管理这些冗余级别的运营成本是多少？

The consequent trade-offs are among the following:

随之而来的权衡如下：

- Customer requirements for achieving high availability
- Business operational cost
- Real world engineering practicality of actually making it possible

- 客户对实现高可用性的要求
- 业务运营成本
- 实际使其成为可能的真实世界工程实用性

Redundant components, and in turn their dependencies, need to be engineered, configured and operated in [a way that ensures that any failures are statistically independent](https://www.ntnu.edu/documents/624876/1277590549/chapt04-1 .pdf). The simple math is: statistically independent failures render your chances of a catastrophic failure exponentially lower as you increase the level of redundancy. If the failures are set up to occur in statistical silos, then there is no cumulative effect, and you decrease the likelihood of complete failure by a whole order of magnitude with each additional redundant resource.

冗余组件及其依赖项需要以[确保任何故障在统计上独立的方式]进行设计、配置和操作（https://www.ntnu.edu/documents/624876/1277590549/chapt04-1 .pdf）。简单的数学计算是：随着冗余级别的增加，统计上独立的故障会使您发生灾难性故障的几率呈指数级降低。如果将故障设置为在统计孤岛中发生，则不会产生累积效应，并且您可以使用每个额外的冗余资源将完全故障的可能性降低一个数量级。

At Ably, to increase statistical independence of failures, [we place/distribute capacity in multiple availability zones and in multiple regions](https://www.ably.io/network). Offering service from multiple availability zones within the same region is relatively simple: AWS enables this with very little effort. Availability zones generally have a good track record of failing independently, so this step by itself enables the existence of sufficient redundancy to support very high levels of availability. 
在 Ably，为了提高故障的统计独立性，[我们在多个可用区和多个区域中放置/分配容量](https://www.ably.io/network)。从同一区域内的多个可用区提供服务相对简单：AWS 可以轻松实现这一点。可用性区域通常具有良好的独立故障记录，因此这一步本身就可以提供足够的冗余来支持非常高的可用性级别。
However, this isn’t the whole story. It isn’t sufficient to rely on any specific region for multiple reasons – sometimes multiple availability zones (AZs) do fail at the same time; sometimes there might be local connectivity issues making the region unreachable; and sometimes there might simply be capacity limitations in a region that prevent all services from being supportable there. As a result, we also promote service availability by providing service in multiple regions. This is the ultimate way of ensuring statistical independence of failures.

然而，这并不是故事的全部。出于多种原因，仅仅依赖于任何特定区域是不够的——有时多个可用区 (AZ) 确实会同时出现故障；有时可能存在本地连接问题，导致无法访问该区域；有时，某个地区可能只是存在容量限制，导致无法在该地区支持所有服务。因此，我们还通过在多个地区提供服务来提高服务可用性。这是确保故障统计独立性的最终方法。

Establishing redundancy that spans multiple regions is not as straightforward as supporting multiple AZs. For example, it doesn’t make sense simply to have a load balancer distributing requests among regions; the load balancer itself exists in some region and could become unavailable.

建立跨多个区域的冗余并不像支持多个可用区那么简单。例如，简单地让负载均衡器在区域之间分发请求是没有意义的；负载均衡器本身存在于某个区域并且可能变得不可用。

Instead, we use a [combination of measures](https://knowledge.ably.com/routing-around-network-and-dns-issues) to ensure that client requests can at all times be routed to a region that is believed to be healthy and have service available. This routing will preferably be to the nearest region, but routing to non-local regions must be possible when the nearest region is unable to provide service.

相反，我们使用 [措施组合](https://knowledge.ably.com/routing-around-network-and-dns-issues) 来确保客户端请求始终可以路由到被认为是的区域保持健康并提供服务。这种路由最好是到最近的区域，但是当最近的区域无法提供服务时，必须可以路由到非本地区域。

## Stateful services

##有状态服务

At Ably, [reliability](https://www.ably.io/four-pillars-of-dependability#reliability) means business continuity of _stateful_ services, and it is a substantially more complicated problem to solve than just availability.

在 Ably，[可靠性](https://www.ably.io/four-pillars-of-dependability#reliability) 意味着 _stateful_ 服务的业务连续性，这是一个比可用性要复杂得多的问题。

Stateful services have an intrinsic dependency on state that survives each individual invocation of service. Continuity of that state translates into correctness of the service provided by the layer as a whole. That requirement for continuity means that fault tolerance for these services is achieved by thinking of them in classic reliability terms. Redundancy needs to be continuously in use, in order for the state not to become lost in the case of failure. Fault detection and remediation needs to address the possible Byzantine failure modes via consensus formation mechanisms.

有状态的服务对状态具有内在的依赖性，在每次单独的服务调用中都存在。该状态的连续性转化为整个层提供的服务的正确性。对连续性的要求意味着这些服务的容错是通过以经典的可靠性术语来考虑它们来实现的。冗余需要持续使用，以便在出现故障时不会丢失状态。故障检测和修复需要通过共识形成机制解决可能的拜占庭故障模式。

The most simplistic analogy is with airplane safety. An airplane crash is catastrophic because you – and your state – are on a _specific_ airplane; it is _that_ airplane which must provide continuous service. If it fails to do so, state is lost and you are afforded no opportunity to continue by migrating to a different airplane.

最简单的类比是飞机安全。飞机失事是灾难性的，因为你——以及你所在的州——在一架_特定的_飞机上；正是那架飞机必须提供持续服务。如果它不这样做，状态就会丢失，并且您没有机会继续迁移到另一架飞机。

With anything that relies on state, when an alternate resource is selected, the requirement is to be able to carry on with the new resource where the previous resource left off. State is thus a requirement, and in those cases availability alone is insufficient.

With anything that relies on state, when an alternate resource is selected, the requirement is to be able to carry on with the new resource where the previous resource left off.因此，状态是一项要求，在这些情况下，仅可用性是不够的。

At Ably, we provision enough reserve capacity for stateless resources to support all of our customers’ availability requirements. However, for stateful resources, not only do we need redundant resources, but also explicit mechanisms to make use of that redundancy, in order to support our service guarantees of _functional continuity_.

在 Ably，我们为无状态资源提供足够的储备容量来支持我们所有客户的可用性要求。然而，对于有状态资源，我们不仅需要冗余资源，还需要显式机制来利用这种冗余，以支持我们对_功能连续性_的服务保证。

![Fault tolerance of a stateful datastore: a query coordinator handles non-binary (Byzantine) failures via transactional replication of updates on multpile stores.](https://ik.imagekit.io/ably/ghost/prod/2021/02/fault-tolerance-of-a-stateful-datastore.png?tr=w-1520)

/fault-tolerance-of-a-stateful-datastore.png?tr=w-1520)

For example, if processing for a particular channel is taking place on a particular instance within the cluster, and that instance fails (forcing that channel role to move), mechanisms must be in place to ensure that things can continue.

例如，如果特定通道的处理发生在集群内的特定实例上，并且该实例失败（强制该通道角色移动），则必须建立机制以确保事情可以继续进行。

This operates at several levels. At one level, a mechanism must exist to ensure that that channel’s [processing is reassigned to a different, healthy, resource](https://knowledge.ably.com/routing-around-network-and-dns-issues). At another level, there needs to be assurance that the reassigned resource continues at exactly the point that processing was halted in the previous location. Further, each of these mechanisms is itself implemented and operated with a level of redundancy to meet the overall assurance requirements for the service.

这在几个层面上运作。在一个层面上，必须存在一种机制来确保该通道的[处理被重新分配给不同的、健康的资源](https://knowledge.ably.com/routing-around-network-and-dns-issues)。在另一个层面上，需要确保重新分配的资源恰好在处理在前一个位置停止的点继续。此外，这些机制中的每一个本身都是以一定的冗余级别实现和操作的，以满足服务的整体保证要求。

The effectiveness of these continuity mechanisms directly translates into the behavior, and assurance of that behavior, at the service provision boundary. Taking one specific issue in the scenario above: for any given message, you need to know with certainty whether or not processing for that message has been completed.

这些连续性机制的有效性直接转化为服务提供边界处的行为和该行为的保证。在上述场景中考虑一个特定问题：对于任何给定的消息，您需要确定地知道该消息的处理是否已完成。

When a client submits a message to Ably for publication, and the service accepts the message for publication, it acknowledges that attempt as successful or unsuccessful. At this point, the principal _availability_ question is:

当客户端向 Ably 提交要发布的消息，并且服务接受要发布的消息时，它会确认该尝试成功或不成功。此时，主要的_可用性_问题是：

> What fraction of the time does the service accept (and then process) the message, versus rejecting it? 
> 服务接受（然后处理）消息与拒绝消息的时间比例是多少？
The minimum aim in this case is [four, five, or even six 9s](https://en.wikipedia.org/wiki/High_availability#Percentage_calculation).

在这种情况下，最低目标是 [四个、五个甚至六个 9](https://en.wikipedia.org/wiki/High_availability#Percentage_calculation)。

If you attempt to publish and we let you know we weren’t able to do that, then that’s merely an **availability shortcoming**. It’s not great, but you end up with an awareness of the situation.

如果您尝试发布并且我们让您知道我们无法做到这一点，那么这只是**可用性缺陷**。这不是很好，但你最终会意识到情况。

However, if we respond with success — "yes, we've received your message" — but then we fail to do all the onward processing of it, then that's a different kind of failure. That's a failure to uphold our functional service guarantee. That's a **reliability shortcoming**, and it's a far more [complicated problem to solve in a distributed system](https://www.ably.io/blog/practical-strategies-for-dns-failure) — there is significant engineering effort and complexity devoted to meeting this requirement.

然而，如果我们成功响应——“是的，我们已经收到了你的消息”——但是我们没有对它进行所有后续处理，那就是另一种失败。这是未能维护我们的功能服务保证。这是一个**可靠性缺点**，这是一个[在分布式系统中解决的复杂问题]（https://www.ably.io/blog/practical-strategies-for-dns-failure）——有为满足这一要求而投入的大量工程工作和复杂性。

## Architectural approaches to achieve reliability

## 实现可靠性的架构方法

The following are two illustrations of the architectural approaches we adopt at Ably to make optimum use of redundancy within our message processing core.

以下是我们在 Ably 采用的架构方法的两个说明，以在我们的消息处理核心中优化利用冗余。

[![Try our APIs for free](https://no-cache.hubspot.com/cta/default/6939709/85f436af-7315-44c9-aa12-d764f40b294d.png)](https://cta-redirect.hubspot.com/cta/redirect/6939709/85f436af-7315-44c9-aa12-d764f40b294d)

hubspot.com/cta/redirect/6939709/85f436af-7315-44c9-aa12-d764f40b294d)

### Stateful role placement

### 有状态角色放置

In general, horizontal scalability is achieved by distributing work across a scalable cluster of processing resources. As far as entities that perform **stateless** processing are concerned, they can be distributed across available resources with **few constraints** on their placement: the location of any given operation can be decided based on load balancing, proximity, or other optimization considerations.

通常，水平可扩展性是通过跨处理资源的可扩展集群分配工作来实现的。就执行 **无状态** 处理的实体而言，它们可以分布在可用资源中，对其放置几乎没有限制**：任何给定操作的位置都可以根据负载平衡、接近度或其他优化考虑。

Meanwhile, in the case of **stateful** operations, the placement of processing roles must take their stateful nature into account such that, for example, all concerned entities can agree on the specific location of any given role.

同时，在**有状态**操作的情况下，处理角色的放置必须考虑其有状态的性质，例如，所有相关实体可以就任何给定角色的具体位置达成一致。

A specific example is channel message processing: whenever a channel is active it gets assigned a resource that processes it. Being a stateful process, it is possible to achieve greater performance: we know more about the message at the time of processing, so we don’t have to look it up — we can just process it and send it on.

一个具体的例子是通道消息处理：每当一个通道处于活动状态时，它就会被分配一个处理它的资源。作为一个有状态的进程，可以实现更高的性能：我们在处理时对消息了解得更多，所以我们不必查找它——我们只需处理它并发送它。

In order to distribute all the channels across all the available resources as uniformly as possible, we use [consistent hashing](https://www.ably.io/blog/implementing-efficient-consistent-hashing/), with the underlying cluster discovery service providing consensus on both node health and hashring membership.

为了在所有可用资源上尽可能均匀地分配所有通道，我们使用 [consistent hashing](https://www.ably.io/blog/implementing-efficient-consistent-hashing/)，与底层集群发现服务提供关于节点健康和哈希成员的共识。

![A consistent placement algorithm is used to decide the placement of any resource in the cluster;all peers must achieve consensus on ring changes. ](https://ik.imagekit.io/ably/ghost/prod/2021/02/hash-rings.png?tr=w-1520)Achieving consensus via consistent placement algorithm in a hash ring

所有的peer都必须就环的变化达成共识。 ](https://ik.imagekit.io/ably/ghost/prod/2021/02/hash-rings.png?tr=w-1520)通过哈希环中的一致放置算法达成共识

The placement mechanism is required not only to determine the initial placement of a role, but also to relocate the role whenever there is an event, such as node failure, that causes the role to move. Therefore, this dynamic placement mechanism is a core functionality in support of service continuity and reliability.

放置机制不仅需要确定角色的初始放置，还需要在发生导致角色移动的事件（例如节点故障）时重新定位角色。因此，这种动态放置机制是支持服务连续性和可靠性的核心功能。

### Detect, hash, resume

### 检测、散列、恢复

The first step in mitigating failure is detecting it. As discussed above, this is classically difficult because of the need to achieve near-simultaneous consensus between distributed entities. Once detected, the updated hashring state implies the new location of that resource, and from that point onward the channel and the state of the failed resource must resume in the new location with continuity.

减轻故障的第一步是检测故障。如上所述，这在传统上是困难的，因为需要在分布式实体之间实现近乎同时的共识。一旦检测到，更新后的散列状态意味着该资源的新位置，从那时起，通道和故障资源的状态必须在新位置连续恢复。

Even though the role failed and there was resulting loss of state, there needs to be sufficient state persisted (and with sufficient redundancy) that resumption of the role is possible with continuity. Resumption with continuity is what enables the service to be reliable in the presence of this kind of failure; the state of each in-flight message is preserved across the role relocation. If we were unable to do that, and simply re-established the role without continuity of state, we could ensure service availability — but not _reliability_.

即使角色失败并导致状态丢失，也需要有足够的状态持久化（并具有足够的冗余），以便可以连续恢复角色。连续性恢复是使服务在出现此类故障时可靠的原因；每个传输中消息的状态在角色重定位过程中都会保留。如果我们无法做到这一点，并在没有状态连续性的情况下简单地重新建立角色，我们可以确保服务可用性——但不能确保_可靠性_。

### Channel persistence layer 
###通道持久层
When a message is published, we perform some processing, decide on success or failure, then respond to the call. The reliability guarantee means having certainty that once a message is acknowledged, all onward transmission implied by that will in fact take place. This in turn means that we can only acknowledge a message once we know for a fact that it is durably persisted, with sufficient redundancy that it cannot subsequently be lost.

当消息发布时，我们执行一些处理，决定成功或失败，然后响应调用。可靠性保证意味着确定一旦消息被确认，所有隐含的向前传输实际上都会发生。这反过来意味着我们只能在知道消息持久存在的事实后才能确认消息，并且具有足够的冗余度，因此不会随后丢失。

First, we record the receipt of the message in **at least two different availability zones** (AZs). Then we have the same multiple-AZ redundancy requirement for the onward processing itself. This is the core of the persistence layer at Ably: we write a message to multiple locations, and we also make sure the process of writing it is transactional. You come away knowing the writing of the message was either _not successful_ or _unequivocally successful_. With assurance of that, subsequent processing can be guaranteed to occur eventually, even if there are failures in the roles responsible for that processing.

首先，我们在**至少两个不同的可用区** (AZ) 中记录消息的接收。然后我们对后续处理本身具有相同的多可用区冗余要求。这是 Ably 持久层的核心：我们将消息写入多个位置，并且我们还确保写入它的过程是事务性的。你离开时知道消息的编写要么_不成功__要么_明确成功_。有了这一保证，即使负责该处理的角色出现故障，也可以保证最终发生后续处理。

Ensuring that messages are persisted in multiple AZs enables us to assume that failures in those zones are independent, so a single event or cause cannot lead to loss of data. Arranging for this placement requires AZ-aware orchestration, and ensuring that writes to multiple locations are actually transactional requires distributed consensus in the message persistence layer.

确保消息在多个可用区中持久化使我们能够假设这些区域中的故障是独立的，因此单个事件或原因不会导致数据丢失。安排此放置需要 AZ 感知编排，并确保对多个位置的写入实际上是事务性的，需要在消息持久层中达成分布式共识。

![Ensuring independence of failure of redundant components requires placing them into different availability zones such that requests can be made against coordinators in any such zone.](https://ik.imagekit.io/ably/ghost/prod/2021/02/ensure-redundant-components-fail-independently.png?tr=w-1520)

/ensure-redundant-components-fail-independently.png?tr=w-1520)

Structuring things this way allows us to start to quantify, probabilistically, the level of assurance. A service failure can only arise if there is a compound failure — that is, a failure in one AZ and, before that failure has been remediated, a failure in a second AZ.

以这种方式构建事物使我们能够开始从概率上量化保证水平。仅当存在复合故障时才会出现服务故障 - 即，一个可用区出现故障，并且在修复该故障之前，第二个可用区出现故障。

In our mathematical model, when one node fails, we know how long it takes to detect and achieve consensus on the failure, and how long it takes subsequently to relocate the role. Knowing this, together with the failure rate of each AZ, it is possible to model the probability of an occurrence of a compound failure that results in loss of continuity of state. This is the basis on which we are able to provide our guarantee of [eight 9s of reliability](https://www.ably.io/four-pillars-of-dependability#reliability).

在我们的数学模型中，当一个节点发生故障时，我们知道检测故障并就故障达成共识需要多长时间，以及随后需要多长时间重新定位角色。知道了这一点，连同每个可用区的故障率，就可以对导致状态连续性丧失的复合故障发生的概率进行建模。这是我们能够提供 [8 个 9 可靠性](https://www.ably.io/four-pillars-of-dependability#reliability) 保证的基础。

## Implementation considerations

## 实施注意事项

Even once you have a theoretical approach to achieving a particular aspect of fault tolerance, there are numerous practical and systems engineering aspects to consider in the wider context of the system as a whole. We give some examples below. To explore this topic further, read and/or watch our deep-dive on the topic, [Hidden scaling issues of distributed systems — system design in the real world](https://www.ably.io/blog/hidden-scaling -issues-of-distributed-systems-real-world).

即使您有了实现容错特定方面的理论方法，在整个系统的更广泛背景下，仍有许多实际和系统工程方面需要考虑。我们在下面给出一些例子。要进一步探索该主题，请阅读和/或观看我们对该主题的深入探讨，[分布式系统的隐藏扩展问题——现实世界中的系统设计](https://www.ably.io/blog/hidden-scaling - 分布式系统问题 - 现实世界）。

### Consensus formation in globally-distributed systems

### 在全球分布式系统中形成共识

The mechanisms described above — such as the role placement algorithm — can only be effective when all of the participating entities are in agreement on the topology of the cluster together with the status and health of each node.

上述机制——例如角色放置算法——只有在所有参与实体都对集群的拓扑以及每个节点的状态和健康状况达成一致时才能有效。

This is a classical consensus formation problem, where the members of a cluster, who might themselves be subject to failure, must agree on the status of one of their members. Consensus formation protocols such as [Raft](https://en.wikipedia.org/wiki/Raft_(algorithm))/ [Paxos](https://en.wikipedia.org/wiki/Paxos_(computer_science)) are widely understood and have strong theoretical guarantees, but also have practical limitations in terms of scalability and bandwidth. In particular, they are not effective in networks spanning multiple regions because their efficiency breaks down if the latency becomes too high when communicating among peers.

这是一个经典的共识形成问题，集群的成员可能会失败，但必须就其中一个成员的地位达成一致。 [Raft](https://en.wikipedia.org/wiki/Raft_(algorithm))/[Paxos](https://en.wikipedia.org/wiki/Paxos_(computer_science))等共识形成协议被广泛使用理解并具有强大的理论保证，但在可扩展性和带宽方面也有实际限制。特别是，它们在跨越多个区域的网络中无效，因为如果在对等方之间通信时延迟变得太高，它们的效率就会下降。

Instead, our peers use [the Gossip protocol](https://www.ably.io/blog/what-is-a-distributed-systems-engineer#gossipprotocolsandconsensusalgorithmsunderpineverything), which is eventually-consistent, fault-tolerant, and can work across regions. Gossip is used among regions to share topology information as well as to construct a network map, and this broad cluster state consensus is then used as the basis for sharing various details cluster-wide.

相反，我们的同行使用 [the Gossip 协议](https://www.ably.io/blog/what-is-a-distributed-systems-engineer#gossipprotocolsandconsensusalgorithmsunderpineverything)，它最终是一致的、容错的，并且可以跨区域工作。 Gossip 用于区域之间共享拓扑信息以及构建网络图，然后将这种广泛的集群状态共识用作在集群范围内共享各种细节的基础。

### Health is not binary 
### 健康不是二进制的
The classical theory that led to the development of Paxos and Raft originated from the recognition that failed or unhealthy entities oftentimes do not simply crash or stop responding. Instead, failing elements can exhibit partly-working behavior such as delayed responses, higher error rates or, in principle, arbitrary confusing or misleading behavior (see [Byzantine fault](https://en.wikipedia.org/wiki/Byzantine_fault)) . This issue is pervasive and even extends to the way a client interacts with the service.

导致 Paxos 和 Raft 发展的经典理论源于这样一种认识，即失败或不健康的实体通常不会简单地崩溃或停止响应。相反，失败的元素可能会表现出部分工作行为，例如延迟响应、更高的错误率，或者原则上任意混淆或误导行为（参见 [拜占庭错误](https://en.wikipedia.org/wiki/Byzantine_fault)） .这个问题很普遍，甚至扩展到客户端与服务交互的方式。

When a client attempts connection to an endpoint in a given region, it is possible that the region is not available. If it’s entirely unavailable, then that’s a simple failure, and we would be able to detect that and redirect the client to greener pastures within the infrastructure. But what happens in practice is that regions can often be in a partially degraded state where they’re working _some_ of the time. This means that the client itself has the problem of handling the failure — knowing where to redirect to obtain service, and knowing when to retry getting service from the default endpoint.

当客户端尝试连接到给定区域中的端点时，该区域可能不可用。如果它完全不可用，那么这是一个简单的故障，我们将能够检测到它并将客户端重定向到基础设施内的绿色牧场。但在实践中发生的情况是，区域通常会处于部分降级状态，在某些时候它们正在工作。这意味着客户端本身存在处理失败的问题——知道从哪里重定向以获取服务，以及知道何时重试从默认端点获取服务。

This is another example of the general fault tolerance problem: with many moving pieces, every single thing introduces additional complexity. If this part breaks, or that thing changes, how do you establish consensus on the _existence_ of the change, the _nature_ of the change, and the consequent _plan of action_?

这是一般容错问题的另一个例子：有许多移动的部分，每一件事都会引入额外的复杂性。如果这部分坏了，或者那件事发生了变化，你如何就变化的_存在_、变化的_性质_以及随之而来的_行动计划_达成共识？

### Resource availability issues

###资源可用性问题

At a very simple level, it is possible to provide redundant capacity only if the resources required are available. Occasionally there are situations where resources are simply unavailable at the moment they are demanded in a region, and it is necessary to offload demand to another region.

在一个非常简单的层面上，只有在所需资源可用时才可能提供冗余容量。有时会出现这样的情况，即一个区域需要资源时，资源根本不可用，需要将需求卸载到另一个区域。

However, the resource availability problem also surfaces in a more challenging way when you realize that fault tolerance mechanisms themselves require resources to operate.

但是，当您意识到容错机制本身需要资源来运行时，资源可用性问题也会以更具挑战性的方式出现。

For example, there is a mechanism to manage the relocation of roles when the topology changes. This functionality itself requires resources in order to operate, such as CPU and memory on the affected instances.

例如，有一种机制可以在拓扑发生变化时管理角色的重定位。此功能本身需要资源才能运行，例如受影响实例上的 CPU 和内存。

But what if your disruption has come about precisely because your CPU or memory ran out? Now you’re attempting to handle those failures, but to do so you need… _CPU and memory_. This means that there are multiple dimensions in which you need to ensure that there exists a resource capacity margin, so that fault tolerance measures can be enacted at the time they are needed.

但是，如果您的中断恰好是因为您的 CPU 或内存用完怎么办？现在您正在尝试处理这些故障，但为此您需要……_CPU 和内存_。这意味着您需要在多个维度上确保存在资源容量裕度，以便在需要时制定容错措施。

### Resource scalability issues

### 资源可扩展性问题

Further to the point above, it’s not just about availability of resources, but also about the rate at which demand for them scales. In the steady, healthy state you might have _N_ channels, _N_ connections, and _N_ messages with _N_ capacity to deal with it all. Now imagine there is some failure and ensuing disruption for that cluster of _N_ instances. If the amount of work required to compensate for the disruption is of size N², then maintaining the capacity margin becomes unsustainable, and the only available remediation is to fail over to an undisrupted region or cluster.

更进一步说，这不仅与资源的可用性有关，还与对资源的需求增长速度有关。在稳定、健康的状态下，您可能有 _N_ 个通道、_N_ 个连接和 _N_ 个消息，具有 _N_ 个容量来处理这一切。现在想象一下，该 _N_ 个实例的集群出现了一些故障并导致中断。如果补偿中断所需的工作量大小为 N²，那么保持容量裕度将变得不可持续，唯一可用的补救措施是故障转移到未中断的区域或集群。

Simplistic fault tolerance mechanisms can exhibit this kind of O( _N²_) behavior or worse, so approaches need to be analyzed with this in mind. It’s another reminder that while things can fail just because they’re broken in some way, it’s also possible that they can go wrong because they have an unforeseen scale or complexity or some other unsustainable resource implication.

简单的容错机制可能会表现出这种 O(_N²_) 行为或更糟的行为，因此需要考虑到这一点来分析方法。另一个提醒是，虽然事情可能仅仅因为它们以某种方式被破坏而失败，但它们也有可能出错，因为它们具有不可预见的规模或复杂性或其他一些不可持续的资源影响。

## Conclusion

＃＃ 结论

Fault tolerance is an approach to building systems able to withstand and mitigate adverse events and operating conditions in order to dependably continue delivering the level of service expected by the users of the system.

容错是一种构建能够承受和减轻不利事件和操作条件的系统的方法，以便可靠地继续提供系统用户期望的服务水平。

Dependability engineering in the physical world classically makes the distinction between **availability** and **reliability** and there are well-understood formulas for how to achieve each of these through fault tolerance and redundancy. At Ably we make the analogous distinction in the systems world between components that are **stateless** and those that are **stateful**. 
物理世界中的可靠性工程经典地区分了**可用性**和**可靠性**，并且有很好理解的公式来说明如何通过容错和冗余来实现这些。在 Ably，我们在系统世界中对 **无状态** 组件和 **有状态** 组件进行了类似的区分。
Fault tolerance for _stateless_ components is achieved in the same way as for availability in the physical world: through provision of redundant elements whose failure is statistically independent. Fault tolerance for _stateful_ components can be compared to the physical reliability problem: it is the assurance of service without interruption. _Continuity of state_ is essential in order that failures can be tolerated whilst preserving _correctness and continuity of the provided service_.

_stateless_ 组件的容错实现方式与物理世界中的可用性相同：通过提供其故障在统计上独立的冗余元素。 _stateful_组件的容错可以比作物理可靠性问题：它是服务不中断的保证。 _状态的连续性_是必不可少的，这样可以容忍故障，同时保持所提供服务的_正确性和连续性_。

To be fault tolerant, a system must treat failures not as exceptional events, but as something expectedly routine. Unlike theoretical models of success and failure, threats to a system’s health in the real world are not binary and require complex theoretical and practical approaches to mitigate.

为了容错，系统必须不将故障视为异常事件，而是将其视为预期的常规事件。与成功和失败的理论模型不同，现实世界中对系统健康的威胁不是二元的，需要复杂的理论和实践方法来缓解。

The Ably service is engineered with multiple layers that each make use of a wide range of fault tolerance mechanisms. In particular, we have confronted the hard engineering problems that arise which include stateful role placement, detection, hashing, and graceful resumption of service, among others. We’ve also articulated, and provide assurance of, the service guarantee at the channel persistence layer, such as assured onward processing once we acknowledge that we’ve received a message.

Ably 服务设计有多个层，每个层都使用广泛的容错机制。特别是，我们遇到了出现的硬工程问题，包括有状态角色放置、检测、散列和服务的优雅恢复等。我们还阐明并提供了通道持久层的服务保证，例如在我们确认收到消息后保证继续处理。

Beyond theoretical approaches, designing fault tolerant systems involves numerous [real-world systems engineering challenges](https://www.ably.io/resources/datasheets/using-ably-at-scale). This includes infrastructure availability and scalability issues, as well as dealing with consensus formation and orchestration of ever-fluctuating topologies of all clusters and nodes in the globally-distributed system, with unpredictable/hard-to-detect health status of any given entity on the network.

除了理论方法之外，设计容错系统还涉及许多[现实世界的系统工程挑战](https://www.ably.io/resources/datasheets/using-ably-at-scale)。这包括基础设施可用性和可扩展性问题，以及处理全球分布式系统中所有集群和节点不断波动的拓扑的共识形成和编排，以及不可预测/难以检测的任何给定实体的健康状态网络。

The Ably platform was designed from the ground up with these principles in mind, with the goal of delivering the best-in-class enterprise solution. This is why we can confidently provide our service-level guarantees of both availability and reliability – and thus dependability and fault tolerance.

Ably 平台是根据这些原则从头开始设计的，目标是提供一流的企业解决方案。这就是为什么我们可以自信地提供可用性和可靠性的服务级别保证，从而保证可靠性和容错性。

_[Get in touch](https://www.ably.io/contact) to learn more about Ably and how we can help you deliver seamless realtime experiences to your customers._ 
_[联系](https://www.ably.io/contact) 了解更多关于 Ably 以及我们如何帮助您为客户提供无缝实时体验。_
