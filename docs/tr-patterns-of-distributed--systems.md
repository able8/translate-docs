# Patterns of Distributed Systems

# 分布式系统的模式

*Distributed systems provide a particular challenge to program. They often require us to have multiple copies of data, which need to keep synchronized. Yet we cannot rely on processing nodes working reliably, and network delays can easily lead to inconsistencies. Despite this, many organizations rely on a range of core distributed software handling data storage, messaging, system management, and compute capability. These systems face common problems which they solve with similar solutions. This article recognizes and develops these solutions as patterns, with which we can build up an understanding of how to better understand, communicate and teach distributed system design.*

*分布式系统对编程提出了特殊的挑战。它们通常要求我们拥有多个数据副本，这些副本需要保持同步。然而，我们不能依赖可靠地工作的处理节点，网络延迟很容易导致不一致。尽管如此，许多组织仍依赖于一系列核心分布式软件处理数据存储、消息传递、系统管理和计算能力。这些系统面临着使用类似解决方案解决的常见问题。本文将这些解决方案识别并开发为模式，通过这些模式，我们可以建立对如何更好地理解、交流和教授分布式系统设计的理解。*

17 June 2021

[Unmesh Joshi](https://twitter.com/unmeshjoshi)

Unmesh Joshi is a Principal Consultant at Thoughtworks. He is a software architecture enthusiast, who believes that understanding principles of distributed systems is as essential today as understanding web architecture or object oriented programming was in the last decade.

Unmesh Joshi 是 Thoughtworks 的首席顾问。他是一名软件架构爱好者，他相信理解分布式系统的原理在今天与理解 Web 架构或面向对象编程在过去十年中一样重要。

## CONTENTS

##  内容

- What this is about
   - Distributed systems - An implementation perspective
     - [Patterns](https://martinfowler.com/articles/patterns-of-distributed-systems/#Patterns)
- Problems and Their Recurring Solutions.
   - [Process crashes](https://martinfowler.com/articles/patterns-of-distributed-systems/#ProcessCrashes)
   - [Network delays](https://martinfowler.com/articles/patterns-of-distributed-systems/#NetworkDelays)
   - [Process Pauses](https://martinfowler.com/articles/patterns-of-distributed-systems/#ProcessPauses)
   - [Unsynchronized Clocks and Ordering Events](https://martinfowler.com/articles/patterns-of-distributed-systems/#UnsynchronizedClocksAndOrderingEvents)
- Putting it all together - Pattern Sequences
   - [Pattern Sequence for implementing consensus](https://martinfowler.com/articles/patterns-of-distributed-systems/#PatternSequenceForImplementingConsensus)
   - [Kubernetes or Kafka Control Plane](https://martinfowler.com/articles/patterns-of-distributed-systems/#KubernetesOrKafkaControlPlane)
   - [Logical Timestamp usage](https://martinfowler.com/articles/patterns-of-distributed-systems/#LogicalTimestampUsage)
- [Next Steps](https://martinfowler.com/articles/patterns-of-distributed-systems/#NextSteps)


------

## What this is about

## 这是关于什么的

For the last several months, I have been conducting workshops on distributed systems at Thoughtworks. One of the key challenges faced while conducting the workshops was how to map the theory of distributed systems to open source code bases such as Kafka or Cassandra, while keeping the discussions generic enough to cover a broad range of solutions. The concept of patterns provided a nice way out.

在过去的几个月里，我一直在 Thoughtworks 举办关于分布式系统的研讨会。举办研讨会时面临的主要挑战之一是如何将分布式系统理论映射到开源代码库（例如 Kafka 或 Cassandra），同时保持讨论的通用性以涵盖广泛的解决方案。模式的概念提供了一个很好的出路。

Pattern structure, by its very nature, allows us to focus on a specific problem, making it very clear why a particular solution is needed. Then the solution description enables us to give a code structure, which is concrete enough to show the actual solution but generic enough to cover a broad range of variations. This patterns technique also allows us to link various patterns together to build a complete system. This gives a nice vocabulary to discuss distributed system implementations.

就其本质而言，模式结构使我们能够专注于特定问题，从而非常清楚为什么需要特定解决方案。然后解决方案描述使我们能够给出一个代码结构，它足够具体以显示实际解决方案，但足够通用以涵盖广泛的变化范围。这种模式技术还允许我们将各种模式链接在一起以构建一个完整的系统。这提供了一个很好的词汇来讨论分布式系统实现。

What follows is a first set of patterns observed in mainstream open source distributed systems. I hope that this set of patterns will be useful to all developers.

下面是在主流开源分布式系统中观察到的第一组模式。我希望这组模式对所有开发人员都有用。

### Distributed systems - An implementation perspective

### 分布式系统 - 实现视角

Today's enterprise architecture is full of platforms and frameworks which are distributed by nature. If we see the sample list of frameworks and platforms used in typical enterprise architecture today, it will look something like following:

今天的企业架构充满了自然分布的平台和框架。如果我们看到当今典型企业架构中使用的框架和平台的示例列表，它将如下所示：

| Type of platform/framework   | Example                                    |
| :--------------------------- |:----------------------------------------- |
| Databases                    | Cassandra, HBase, Riak                     |
| Message Brokers              | Kafka, Pulsar                              |
| Infrastructure               | Kubernetes, Mesos, Zookeeper, etcd, Consul |
| In Memory Data/Compute Grids | Hazelcast, Pivotal Gemfire                 |
| Stateful Microservices       | Akka Actors, Axon                          |
| File Systems                 | HDFS, Ceph                                 | 



All these are 'distributed' by nature. What does it mean for a system to be distributed? There are two aspects:

所有这些都是自然“分布”的。分布式系统意味着什么？有两个方面：

- They run on multiple servers. The number of servers in a cluster can vary from as few as three servers to a few thousand servers.
- They manage data. So these are inherently 'stateful' systems.

- 他们在多台服务器上运行。集群中的服务器数量可以从三台服务器到几千台服务器不等。
- 他们管理数据。因此，这些本质上是“有状态”系统。

There are several ways in which things can go wrong when multiple servers are involved in storing data. All the above mentioned systems need to solve those problems. The implementation of these systems have some recurring solutions to these problems. Understanding these solutions in their general form helps in understanding the implementation of the broad spectrum of these systems and can also serve as a good guidance when new systems need to be built. Enter patterns.

当多个服务器参与存储数据时，有几种方式可能会出错。所有上述系统都需要解决这些问题。这些系统的实施对这些问题有一些反复出现的解决方案。以一般形式理解这些解决方案有助于理解这些系统的广泛实施，并且在需要构建新系统时也可以作为很好的指导。输入图案。

#### Patterns

#### 模式

[Patterns](https://martinfowler.com/articles/writingPatterns.html), a concept introduced by Christopher Alexander, is widely accepted in the software community to document design constructs which are used to build software systems. Patterns provide a structured way of looking at a problem space along with the solutions which are seen multiple times and proven. An interesting way to use patterns is the ability to link several patterns together, in a form of pattern sequence or pattern language, which gives some guidance of implementing a ‘whole’ or a complete system. Looking at distributed systems as a series of patterns is a useful way to gain insights into their implementation.

[Patterns](https://martinfowler.com/articles/writingPatterns.html) 是 Christopher Alexander 引入的一个概念，被软件社区广泛接受，用于记录用于构建软件系统的设计结构。模式提供了一种结构化的方式来查看问题空间以及多次看到并得到验证的解决方案。使用模式的一种有趣方式是能够以模式序列或模式语言的形式将多个模式链接在一起，这为实现“整体”或完整系统提供了一些指导。将分布式系统视为一系列模式是深入了解其实现的有用方法。

## Problems and Their Recurring Solutions.

## 问题及其反复出现的解决方案。

The following examples from Amazon, Google and GitHub illustrate how issues can arrive in even the most sophisticated of setups.

以下来自 Amazon、Google 和 GitHub 的示例说明了即使在最复杂的设置中也会出现问题。

This [GitHub outage](https://github.blog/2018-10-30-oct21-post-incident-analysis/) essentially caused loss of connectivity between its East and West coast data centers. It caused a small window of time in which data could not be replicated across the data centers, causing two MySQL servers to have inconsistent data.

这次 [GitHub 中断](https://github.blog/2018-10-30-oct21-post-incident-analysis/) 基本上导致其东海岸和西海岸数据中心之间的连接中断。它导致数据无法跨数据中心复制的一小段时间窗口，导致两个 MySQL 服务器的数据不一致。

This [AWS outage](https://aws.amazon.com/message/41926/), caused by human error where an automation script was wrongly passed a parameter to take down a large number of servers.

此 [AWS 中断](https://aws.amazon.com/message/41926/) 是由人为错误引起的，其中自动化脚本错误地传递了参数以关闭大量服务器。

This [Google outage](https://status.cloud.google.com/incident/cloud-networking/19009), caused by some misconfiguration, caused a significant impact on the network capacity causing network congestion and service disruption.

这次[谷歌宕机](https://status.cloud.google.com/incident/cloud-networking/19009)，由于一些配置错误，对网络容量造成了重大影响，导致网络拥塞和服务中断。

Several things can go wrong when data is stored on multiple servers.

当数据存储在多台服务器上时，可能会出现一些问题。

### Process crashes

### 进程崩溃

Processes can crash at any time maybe due to hardware faults or software faults. There are numerous ways in which a process can crash.

由于硬件故障或软件故障，进程可能随时崩溃。进程崩溃的方式有很多种。

- It can be taken down for routine maintenance by system administrators.
- It can be killed doing some file IO because the disk is full and the exception is not properly handled.
- In cloud environments, it can be even trickier, as some unrelated events can bring the servers down.

- 系统管理员可以将其取下进行日常维护。
- 由于磁盘已满且未正确处理异常，执行某些文件 IO 时可能会杀死它。
- 在云环境中，它可能更加棘手，因为一些不相关的事件可能会导致服务器宕机。

The bottom line is that if the processes are responsible for storing data, they must be designed to give a durability guarantee for the data stored on the servers. Even if a process crashes abruptly, it should preserve all the data for which it has notified the user that it's stored successfully. Depending on the access patterns, different storage engines have different storage structures, ranging from a simple hash map to a sophisticated graph storage. Because flushing data to the disk is one of the most time consuming operations, not every insert or update to the storage can be flushed to disk. So most databases have in-memory storage structures which are only periodically flushed to disk. This poses a risk of losing all the data if the process abruptly crashes. 

底线是，如果进程负责存储数据，那么它们的设计必须为存储在服务器上的数据提供持久性保证。即使一个进程突然崩溃，它也应该保留它已通知用户它已成功存储的所有数据。根据访问模式，不同的存储引擎具有不同的存储结构，从简单的哈希映射到复杂的图存储。由于将数据刷新到磁盘是最耗时的操作之一，因此并非对存储的每次插入或更新都可以刷新到磁盘。所以大多数数据库都有内存存储结构，这些结构只会定期刷新到磁盘。如果进程突然崩溃，这会带来丢失所有数据的风险。

A technique called [Write-Ahead Log](https://martinfowler.com/articles/patterns-of-distributed-systems/wal.html) is used to tackle this situation. Servers store each state change as a command in an append-only file on a hard disk. Appending a file is generally a very fast operation, so it can be done without impacting performance. A single log, which is appended sequentially, is used to store each update. At the server startup, the log can be replayed to build in memory state again.

一种称为 [Write-Ahead Log](https://martinfowler.com/articles/patterns-of-distributed-systems/wal.html) 的技术用于解决这种情况。服务器将每个状态更改作为命令存储在硬盘上的仅附加文件中。附加文件通常是一个非常快的操作，因此可以在不影响性能的情况下完成。一个按顺序附加的日志用于存储每个更新。在服务器启动时，可以重放日志以再次建立内存状态。

This gives a durability guarantee. The data will not get lost even if the server abruptly crashes and then restarts. But clients will not be able to get or store any data till the server is back up. So we lack availability in the case of server failure.

这提供了耐用性保证。即使服务器突然崩溃然后重新启动，数据也不会丢失。但是在服务器备份之前，客户端将无法获取或存储任何数据。因此，在服务器故障的情况下，我们缺乏可用性。

One of the obvious solutions is to store the data on multiple servers. So we can replicate the write ahead log on multiple servers.

显而易见的解决方案之一是将数据存储在多台服务器上。所以我们可以在多个服务器上复制预写日志。

When multiple servers are involved, there are a lot more failure scenarios which need to be considered.

当涉及多个服务器时，需要考虑更多的故障场景。

### Network delays

### 网络延迟

In the TCP/IP protocol stack, there is no upper bound on delays caused in transmitting messages across a network. It can vary based on the load on the network. For example, a 1 Gbps network link can get flooded with a big data job that's triggered, filling the network buffers, which can cause arbitrary delay for some messages to reach the servers.

在 TCP/IP 协议栈中，跨网络传输消息所引起的延迟没有上限。它可以根据网络上的负载而变化。例如，1 Gbps 网络链接可能会被触发的大数据作业淹没，填满网络缓冲区，这可能导致某些消息到达服务器的任意延迟。

In a typical data center, servers are packed together in racks, and there are multiple racks connected by a top-of-the-rack switch. There might be a tree of switches connecting one part of the data center to the other. It is possible in some cases, that a set of servers can communicate with each other, but are disconnected from another set of servers. This situation is called a network partition. One of the fundamental issues with servers communicating over a network then is how to know a particular server has failed.

在典型的数据中心中，服务器被打包在机架中，并且有多个机架通过架顶式交换机连接。可能有一棵交换机树将数据中心的一个部分连接到另一个部分。在某些情况下，一组服务器可以相互通信，但与另一组服务器断开连接是可能的。这种情况称为网络分区。服务器通过网络通信的基本问题之一是如何知道特定服务器发生故障。

There are two problems to be tackled here.

这里有两个问题需要解决。

- A particular server can not wait indefinitely to know if another server has crashed.
- There should not be two sets of servers, each considering another set to have failed, and therefore continuing to serve different sets of clients. This is called the split brain.

- 一个特定的服务器不能无限期地等待知道另一个服务器是否已经崩溃。
- 不应该有两组服务器，每组都认为另一组出现故障，因此继续为不同的客户端组提供服务。这被称为裂脑。

To tackle the first problem, every server sends a [HeartBeat](https://martinfowler.com/articles/patterns-of-distributed-systems/heartbeat.html) message to other servers at a regular interval. If a heartbeat is missed, the server sending the heartbeat is considered crashed. The heartbeat interval is small enough to make sure that it does not take a lot of time to detect server failure. As we will see below, in the worst case scenario, the server might be up and running, but the cluster as a group can move ahead considering the server to be failing. This makes sure that services provided to clients are not interrupted.

为了解决第一个问题，每个服务器都会定期向其他服务器发送 [HeartBeat](https://martinfowler.com/articles/patterns-of-distributed-systems/heartbeat.html) 消息。如果错过心跳，则发送心跳的服务器将被视为崩溃。心跳间隔足够小，以确保不会花费很多时间来检测服务器故障。正如我们将在下面看到的，在最坏的情况下，服务器可能会启动并运行，但考虑到服务器出现故障，集群作为一个组可以继续前进。这可确保提供给客户端的服务不会中断。

The second problem is the split brain. With the split brain, if two sets of servers accept updates independently, different clients can get and set different data, and once the split brain is resolved, it's impossible to resolve conflicts automatically. 

第二个问题是脑裂。对于裂脑，如果两套服务器独立接受更新，不同的客户端可以获取和设置不同的数据，一旦裂脑解决，就不可能自动解决冲突。

To take care of the split brain issue, we must ensure that the two sets of servers, which are disconnected from each other, should not be able to make progress independently. To ensure this, every action the server takes, is considered successful only if the majority of the servers can confirm the action. If servers can not get a majority, they will not be able to provide the required services, and some group of the clients might not be receiving the service, but servers in the cluster will always be in a consistent state. The number of servers making the majority is called a [Quorum](https://martinfowler.com/articles/patterns-of-distributed-systems/quorum.html). How to decide on the quorum? That is decided based on the number of failures the cluster can tolerate. So if we have a cluster of five nodes, we need a quorum of three. In general, if we want to tolerate `f` failures we need a cluster size of 2f + 1.

为了解决裂脑问题，我们必须确保彼此断开连接的两组服务器不能独立进行。为确保这一点，服务器采取的每项操作只有在大多数服务器都可以确认该操作时才被认为是成功的。如果服务器无法获得多数，它们将无法提供所需的服务，并且某些客户端可能无法接收服务，但集群中的服务器将始终处于一致状态。占多数的服务器数量称为 [Quorum](https://martinfowler.com/articles/patterns-of-distributed-systems/quorum.html)。如何决定法定人数？这是根据集群可以容忍的故障数量决定的。因此，如果我们有一个由五个节点组成的集群，我们需要三个节点的仲裁。一般来说，如果我们想要容忍 `f` 失败，我们需要 2f + 1 的集群大小。

Quorum makes sure that we have enough copies of data to survive some server failures. But it is not enough to give strong consistency guarantees to clients. Lets say a client initiates a write operation on the quorum, but the write operation succeeds only on one server. The other servers in the quorum still have old values. When a client reads the values from the quorum, it might get the latest value, if the server having the latest value is available. But it can very well get an old value if, just when the client starts reading the value, the server with the latest value is not available. To avoid such situations, someone needs to track if the quorum agrees on a particular operation and only send values to clients which are guaranteed to be available on all the servers. [Leader and Followers](https://martinfowler.com/articles/patterns-of-distributed-systems/leader-follower.html) is used in this situation. One of the servers is elected a leader and the other servers act as followers. The leader controls and coordinates the replication on the followers. The leader now needs to decide, which changes should be made visible to the clients. A [High-Water Mark](https://martinfowler.com/articles/patterns-of-distributed-systems/high-watermark.html) is used to track the entry in the write ahead log that is known to have successfully replicated to a quorum of followers. All the entries upto the high-water mark are made visible to the clients. The leader also propagates the high-water mark to the followers. So in case the leader fails and one of the followers becomes the new leader, there are no inconsistencies in what a client sees.

Quorum 确保我们有足够的数据副本来应对某些服务器故障。但是仅仅给客户端强一致性保证是不够的。假设客户端在仲裁上发起写入操作，但写入操作仅在一台服务器上成功。仲裁中的其他服务器仍然具有旧值。当客户端从仲裁读取值时，如果具有最新值的服务器可用，它可能会获得最新值。但是，如果就在客户端开始读取值时，具有最新值的服务器不可用，则它可以很好地获得旧值。为了避免这种情况，有人需要跟踪仲裁是否同意特定操作，并且只向客户端发送保证在所有服务器上可用的值。 [Leader and Followers](https://martinfowler.com/articles/patterns-of-distributed-systems/leader-follower.html)就是在这种情况下使用的。其中一台服务器被选为领导者，其他服务器充当跟随者。领导者控制和协调追随者的复制。领导者现在需要决定哪些更改应该对客户端可见。[高水位标记](https://martinfowler.com/articles/patterns-of-distributed-systems/high-watermark.html) 用于跟踪预写日志中已知已成功复制的条目到法定人数的追随者。客户可以看到高水位线之前的所有条目。领导者还将高水位标记传播给追随者。因此，如果领导者失败并且其中一个追随者成为新的领导者，那么客户所看到的就不会有不一致之处。

### Process Pauses 

### 进程暂停

Even with quorums and leader and followers, there is a tricky problem that needs to be solved. Leader processes can pause arbitrarily. There are a lot of reasons a process can pause. For languages which support garbage collection, there can be a long garbage collection pause. A leader with a long garbage collection pause, can be disconnected from the followers, and will continue sending messages to followers after the pause is over. In the meanwhile, because followers did not receive a heartbeat from the leader, they might have elected a new leader and accepted updates from the clients. If the requests from the old leader are processed as is, they might overwrite some of the updates. So we need a mechanism to detect requests from out-of-date leaders. Here [Generation Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/generation.html) is used to mark and detect requests from older leaders. The generation is a number which is monotonically increasing.

即使有法定人数、领导者和追随者，也有一个棘手的问题需要解决。领导进程可以任意暂停。进程暂停的原因有很多。对于支持垃圾收集的语言，可能会有很长的垃圾收集暂停。垃圾收集暂停时间较长的领导者，可以与追随者断开连接，并在暂停结束后继续向追随者发送消息。同时，因为追随者没有收到领导者的心跳，他们可能已经选举了一个新的领导者并接受了客户端的更新。如果来自旧领导者的请求按原样处理，它们可能会覆盖一些更新。所以我们需要一种机制来检测来自过时领导者的请求。这里的 [生成时钟](https://martinfowler.com/articles/patterns-of-distributed-systems/generation.html) 用于标记和检测来自老领导的请求。代是一个单调递增的数字。

### Unsynchronized Clocks and Ordering Events

### 不同步的时钟和排序事件

The problem of detecting older leader messages from newer ones is the problem of maintaining ordering of messages. It might appear that we can use system timestamps to order a set of messages, but we can not. The main reason we can not use system clocks is that system clocks across servers are not guaranteed to be synchronized. A time-of-the-day clock in a computer is managed by a quartz crystal and measures time based on the oscillations of the crystal.

从新消息中检测旧领导消息的问题是维护消息排序的问题。看起来我们可以使用系统时间戳对一组消息进行排序，但我们不能。我们不能使用系统时钟的主要原因是跨服务器的系统时钟不能保证同步。计算机中的时钟由石英晶体管理，并根据晶体的振荡来测量时间。

This mechanism is error prone, as the crystals can oscillate faster or slower and so different servers can have very different times. The clocks across a set of servers are synchronized by a service called NTP. This service periodically checks a set of global time servers, and adjusts the computer clock accordingly. 

这种机制很容易出错，因为晶体可以更快或更慢地振荡，因此不同的服务器可能有非常不同的时间。一组服务器上的时钟由称为 NTP 的服务同步。此服务会定期检查一组全球时间服务器，并相应地调整计算机时钟。

Because this happens with communication over a network, and network delays can vary as discussed in the above sections, the clock synchronization might be delayed because of a network issue. This can cause server clocks to drift away from each other, and after the NTP sync happens, even move back in time. Because of these issues with computer clocks, time of day is generally not used for ordering events. Instead a simple technique called [Lamport Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/lamport-clock.html) is used. [Generation Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/generation.html) is an example of that. Lamport Clocks are just simple numbers, which are incremented only when some event happens in the system. In a database, the events are about writing and reading the values, so the lamport clock is incremented only when a value is written. The Lamport Clock numbers are also passed in the messages sent to other processes. The receiving process can then select the larger of the two numbers, the one it receives in the message and the one it maintains. This way Lamport Clocks also track happens-before relationship between events across processes which communicate with each other. An example of this is the servers taking part in a transaction. While the [Lamport Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/lamport-clock.html) allows ordering of events, it does not have any relation to the time of the day clock. To bridge this gap, a variation called [Hybrid Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/hybrid-clock.html) is used. The Hybrid Clock uses system time along with a separate number to make sure the value increases monotonically, and can be used the same way as Lamport Clock.

由于这种情况发生在网络通信中，并且网络延迟可能会有所不同，如上述部分所述，时钟同步可能会因网络问题而延迟。这可能会导致服务器时钟相互偏离，并且在 NTP 同步发生后，甚至会及时回退。由于计算机时钟的这些问题，一天中的时间通常不用于排序事件。取而代之的是一种称为 [Lamport Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/lamport-clock.html) 的简单技术。 [Generation Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/generation.html) 就是一个例子。 Lamport 时钟只是简单的数字，仅当系统中发生某些事件时才会递增。在数据库中，事件与写入和读取值有关，因此只有在写入值时才会增加 lamport 时钟。 Lamport 时钟编号也在发送到其他进程的消息中传递。然后，接收进程可以选择两个数字中较大的一个，即它在消息中接收的一个和它维护的一个。通过这种方式，Lamport 时钟还可以跟踪相互通信的进程之间的事件之间的发生前关系。这方面的一个例子是参与交易的服务器。虽然 [Lamport 时钟](https://martinfowler.com/articles/patterns-of-distributed-systems/lamport-clock.html) 允许对事件进行排序，但它与时钟的时间没有任何关系。为了弥补这一差距，使用了一种称为 [混合时钟](https://martinfowler.com/articles/patterns-of-distributed-systems/hybrid-clock.html) 的变体。混合时钟使用系统时间和一个单独的数字来确保值单调增加，并且可以像兰波特时钟一样使用。

The Lamport Clock allows determining the order of events across a set of communicating servers. But it does not allow detecting concurrent updates to the same value happening across a set of replicas. [Version Vector](https://martinfowler.com/articles/patterns-of-distributed-systems/version-vector.html) is used to detect conflict across a set of replicas.

Lamport 时钟允许确定一组通信服务器之间的事件顺序。但它不允许检测对一组副本中发生的相同值的并发更新。 [版本向量](https://martinfowler.com/articles/patterns-of-distributed-systems/version-vector.html) 用于检测一组副本之间的冲突。

The Lamport Clock or Version Vector needs to be associated with the stored values, to detect which values are stored after the other or if there are conflicts. So the servers store the values as [Versioned Value](https://martinfowler.com/articles/patterns-of-distributed-systems/versioned-value.html).

Lamport 时钟或版本向量需要与存储的值相关联，以检测哪些值在另一个之后存储或是否存在冲突。因此，服务器将值存储为 [Versioned Value](https://martinfowler.com/articles/patterns-of-distributed-systems/versioned-value.html)。

## Putting it all together - Pattern Sequences 

## 将它们放在一起 - 模式序列

We can see how understanding these patterns, helps us build a complete system, from the ground up. We will take consensus implementation as an example. Distributed Consensus is a special case of distributed system implementation, which provides the strongest consistency guarantee. Common examples seen in popular enterprise systems are, [Zookeeper](https://zookeeper.apache.org/),[etcd](https://etcd.io/) and [Consul](https://www.consul.io/). They implement consensus algorithms such as [zab](https://zookeeper.apache.org/doc/r3.4.13/zookeeperInternals.html#sc_atomicBroadcast) and [Raft](https://raft.github.io/) to provide replication and strong consistency. There are other popular algorithms to implement consensus, [Paxos](https://en.wikipedia.org/wiki/Paxos_(computer_science)) which is used in Google's [Chubby](https://research.google/pubs/pub27897/) locking service, view stamp replication and [virtual-synchrony](https://www.cs.cornell.edu/ken/History.pdf). In very simple terms, Consensus refers to a set of servers which agree on stored data, the order in which the data is stored and when to make that data visible to the clients.

我们可以看到理解这些模式如何帮助我们从头开始构建一个完整的系统。我们将以共识实现为例。分布式共识是分布式系统实现的一个特例，提供了最强的一致性保证。在流行的企业系统中常见的例子有，[Zookeeper](https://zookeeper.apache.org/)、[etcd](https://etcd.io/) 和 [Consul](https://www.consul.io/)。他们实现了 [zab](https://zookeeper.apache.org/doc/r3.4.13/zookeeperInternals.html#sc_atomicBroadcast) 和 [Raft](https://raft.github.io/)等共识算法来提供复制和强一致性。还有其他流行的算法来实现共识，[Paxos](https://en.wikipedia.org/wiki/Paxos_(computer_science)) 用于谷歌的 [Chubby](https://research.google/pubs/pub27897/) 锁定服务，查看图章复制和 [虚拟同步](https://www.cs.cornell.edu/ken/History.pdf)。用非常简单的术语来说，共识是指一组服务器就存储的数据、存储数据的顺序以及何时使客户端可见该数据达成一致。

### Pattern Sequence for implementing consensus

### 实现共识的模式序列

Consensus implementations use [state machine replication](https://en.wikipedia.org/wiki/State_machine_replication) to achieve fault tolerance. In state machine replication, the storage services, like a key value store, are replicated on all the servers, and the user inputs are executed in the same order on each server. The key implementation technique used to achieve this is to replicate [Write-Ahead Log](https://martinfowler.com/articles/patterns-of-distributed-systems/wal.html) on all the servers to have a 'Replicated Wal '.

共识实现使用[状态机复制](https://en.wikipedia.org/wiki/State_machine_replication)来实现容错。在状态机复制中，存储服务（如键值存储)在所有服务器上复制，并且用户输入在每个服务器上以相同的顺序执行。用于实现此目的的关键实现技术是在所有服务器上复制 [Write-Ahead Log](https://martinfowler.com/articles/patterns-of-distributed-systems/wal.html) 以具有“复制的 Wal” '。

We can put the patterns together to implement Replicated Wal as follows. 

我们可以将这些模式放在一起来实现 Replicated Wal，如下所示。

For providing durability guarantees, you can use the [Write-Ahead Log](https://martinfowler.com/articles/patterns-of-distributed-systems/wal.html) pattern. The Write Ahead Log is divided into multiple segments using [Segmented Log](https://martinfowler.com/articles/patterns-of-distributed-systems/log-segmentation.html). This helps with log cleaning which is handled by [Low-Water Mark](https://martinfowler.com/articles/patterns-of-distributed-systems/low-watermark.html). Fault tolerance is provided by replicating the write ahead log on multiple servers. The replication amongst the servers is managed by using the [Leader and Followers](https://martinfowler.com/articles/patterns-of-distributed-systems/leader-follower.html) pattern and [Quorum](https://martinfowler.com/articles/patterns-of-distributed-systems/quorum.html) is used to update the [High-Water Mark](https://martinfowler.com/articles/patterns-of-distributed-systems/high-watermark.html) to decide which values are visible to clients. All the requests are processed in strict order, by using [Singular Update Queue](https://martinfowler.com/articles/patterns-of-distributed-systems/singular-update-queue.html). The order is maintained while sending the requests from leaders to followers using [Single Socket Channel](https://martinfowler.com/articles/patterns-of-distributed-systems/single-socket-channel.html). To optimize for throughput and latency over a single socket channel, [Request Pipeline](https://martinfowler.com/articles/patterns-of-distributed-systems/request-pipeline.html) can be used. Followers know about availability of the leader via the [HeartBeat](https://martinfowler.com/articles/patterns-of-distributed-systems/heartbeat.html) received from the leader. If the leader is temporarily disconnected from the cluster because of network partition, it is detected by using [Generation Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/generation.html). If all the requests are served only by the leader, it might get overloaded. When the clients are read only and tolerate reading stale values, they can be served by the follower servers. [Follower Reads](https://martinfowler.com/articles/patterns-of-distributed-systems/follower-reads.html) allows handling read requests from follower servers. 

为了提供持久性保证，您可以使用 [Write-Ahead Log](https://martinfowler.com/articles/patterns-of-distributed-systems/wal.html) 模式。 Write Ahead Log 使用 [Segmented Log](https://martinfowler.com/articles/patterns-of-distributed-systems/log-segmentation.html) 分为多个段。这有助于由 [低水位标记](https://martinfowler.com/articles/patterns-of-distributed-systems/low-watermark.html) 处理的日志清理。容错是通过在多个服务器上复制预写日志来提供的。服务器之间的复制通过使用 [Leader and Followers](https://martinfowler.com/articles/patterns-of-distributed-systems/leader-follower.html) 模式和 [Quorum](https://martinfowler.com/articles/patterns-of-distributed-systems/quorum.html) 用于更新 [高水位标记](https://martinfowler.com/articles/patterns-of-distributed-systems/high-watermark.html) 来决定哪些值对客户端可见。通过使用[单一更新队列](https://martinfowler.com/articles/patterns-of-distributed-systems/singular-update-queue.html)，所有请求都按照严格的顺序进行处理。在使用 [Single Socket Channel](https://martinfowler.com/articles/patterns-of-distributed-systems/single-socket-channel.html) 将来自领导者的请求发送到追随者的同时维护该顺序。要优化单个套接字通道的吞吐量和延迟，可以使用 [请求管道](https://martinfowler.com/articles/patterns-of-distributed-systems/request-pipeline.html)。追随者通过从领导者那里收到的[HeartBeat](https://martinfowler.com/articles/patterns-of-distributed-systems/heartbeat.html) 了解领导者的可用性。如果leader因为网络分区暂时与集群断开连接，可以使用 [生成时钟](https://martinfowler.com/articles/patterns-of-distributed-systems/generation.html)检测。如果所有请求都只由领导者提供服务，它可能会过载。当客户端是只读的并且可以容忍读取过时的值时，它们可以由跟随服务器提供服务。 [Follower Reads](https://martinfowler.com/articles/patterns-of-distributed-systems/follower-reads.html) 允许处理来自跟随服务器的读取请求。



### Kubernetes or Kafka Control Plane

### Kubernetes 或 Kafka 控制平面

Products like [Kubernetes](https://kubernetes.io/) or [Kafka's](https://kafka.apache.org/) architecture are built around a strongly consistent metadata store. We can understand it as a pattern sequence. [Consistent Core](https://martinfowler.com/articles/patterns-of-distributed-systems/consistent-core.html) is used as a strongly consistent, fault tolerant metadata store. [Lease](https://martinfowler.com/articles/patterns-of-distributed-systems/time-bound-lease.html) is used to implement group membership and failure detection of cluster nodes. Cluster nodes use [State Watch](https://martinfowler.com/articles/patterns-of-distributed-systems/state-watch.html) to get notified when any cluster node fails or updates its metadata. The [Consistent Core]( https://martinfowler.com/articles/patterns-of-distributed-systems/consistent-core.html) implementation uses [Idempotent Receiver](https://martinfowler.com/articles/patterns-of-distributed-systems/idempotent-receiver.html) to ignore duplicate requests sent by cluster nodes in case of retries on network failure. The Consistent Core is built with a 'Replicated Wal', which is described as a pattern sequence in the above section.

像 [Kubernetes](https://kubernetes.io/) 或 [Kafka's](https://kafka.apache.org/) 架构之类的产品是围绕高度一致的元数据存储构建的。我们可以把它理解为一个模式序列。 [Consistent Core](https://martinfowler.com/articles/patterns-of-distributed-systems/consistent-core.html) 用作强一致性、容错元数据存储。 [Lease](https://martinfowler.com/articles/patterns-of-distributed-systems/time-bound-lease.html)用于实现集群节点的组成员关系和故障检测。集群节点使用 [State Watch](https://martinfowler.com/articles/patterns-of-distributed-systems/state-watch.html) 在任何集群节点出现故障或更新其元数据时获得通知 [Consistent Core]( https://martinfowler.com/articles/patterns-of-distributed-systems/consistent-core.html) 实现使用[幂等接收器](https://martinfowler.com/articles/patterns-of-distributed-systems/idempotent-receiver.html) 忽略集群节点发送的重复请求，以防重试网络故障。 Consistent Core 是用“Replicated Wal”构建的，它在上一节中被描述为一个模式序列。



### Logical Timestamp usage 

### 逻辑时间戳使用

Usage of various types of [logical timestamps](https://en.wikipedia.org/wiki/Logical_clock) can also be seen as a pattern sequence. Various products use either a [Gossip Dissemination](https://martinfowler.com/articles/patterns-of-distributed-systems/gossip-dissemination.html) or a [Consistent Core](https://martinfowler.com/articles/patterns-of-distributed-systems/consistent-core.html) for group membership and failure detection of cluster nodes. The data storage uses [Versioned Value](https://martinfowler.com/articles/patterns-of-distributed-systems/versioned-value.html) to be able to determine which values are most recent. If a single server is responsible for updating the values or [Leader and Followers](https://martinfowler.com/articles/patterns-of-distributed-systems/leader-follower.html) is used, then a [Lamport Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/lamport-clock.html) can be used as a version, in the [Versioned Value](https://martinfowler.com/articles/patterns-of-distributed-systems/versioned-value.html). When the timestamp values need to be derived from the time of the day, a [Hybrid Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/hybrid-clock.html) is used instead of a simple Lamport Clock. If multiple servers are allowed to handle client requests to update the same value, a [Version Vector](https://martinfowler.com/articles/patterns-of-distributed-systems/version-vector.html) is used to be able to detect concurrent writes on different cluster nodes.

各种类型的[逻辑时间戳](https://en.wikipedia.org/wiki/Logical_clock)的使用也可以看作是一个模式序列。各种产品使用 [Gossip Dissemination](https://martinfowler.com/articles/patterns-of-distributed-systems/gossip-dissemination.html) 或 [Consistent Core](https://martinfowler.com/articles/patterns-of-distributed-systems/consistent-core.html) 用于集群节点的组成员身份和故障检测。数据存储使用 [Versioned Value](https://martinfowler.com/articles/patterns-of-distributed-systems/versioned-value.html) 来确定哪些值是最新的。如果单个服务器负责更新值或使用 [Leader and Followers](https://martinfowler.com/articles/patterns-of-distributed-systems/leader-follower.html)，则使用 [Lamport Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/lamport-clock.html) 可以用作版本，在 [Versioned Value](https://martinfowler.com/articles/patterns-of-distributed-systems/versioned-value.html)。当时间戳值需要从一天中的时间导出时，使用 [Hybrid Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/hybrid-clock.html)而不是简单的兰波特时钟。如果允许多个服务器处理客户端请求以更新相同的值，则使用[版本向量](https://martinfowler.com/articles/patterns-of-distributed-systems/version-vector.html) 即可检测不同集群节点上的并发写入。



This way, understanding problems and their recurring solutions in their general form, helps in understanding building blocks of a complete system.

这样，以一般形式理解问题及其反复出现的解决方案，有助于理解完整系统的构建块

## Next Steps

##  下一步

Distributed systems is a vast topic. The set of patterns covered here is a small part, covering different categories to showcase how a patterns approach can help understand and design distributed systems. I will keep adding to this set to broadly include the following categories of problems solved in any distributed system

分布式系统是一个很大的话题。这里涵盖的模式集只是一小部分，涵盖了不同的类别，以展示模式方法如何帮助理解和设计分布式系统。我将继续添加到这个集合中，以广泛地包括以下类别在任何分布式系统中解决的问题

- Group Membership and Failure Detection
- Partitioning
- Replication and Consistency
- Storage
- Processing



- 组成员和故障检测
- 分区
- 复制和一致性
- 贮存
- 加工

This page is part of:

此页面是以下内容的一部分：

Patterns of Distributed Systems

分布式系统的模式

------

## Acknowledgements

## 致谢

Many thanks to Martin Fowler for helping me throughout and guiding me to think in terms of patterns.

非常感谢 Martin Fowler 在整个过程中帮助我并指导我从模式方面进行思考。

Mushtaq Ahemad helped me with good feedback and a lot of discussions throughout

Mushtaq Ahemad 在整个过程中帮助我提供了良好的反馈和大量讨论

Professor Indranil Gupta provided feedback on the gossip dissemination pattern.

Indranil Gupta 教授提供了有关八卦传播模式的反馈。

Thanks to Jojo Swords, Gareth Morgan for helping with copy editing.

感谢 Jojo Swords，Gareth Morgan 帮助编辑。

Significant Revisions 

重大修订

