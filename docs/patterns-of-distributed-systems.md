# Patterns of Distributed Systems

![](http://martinfowler.com/card.png)

_Distributed systems provide a particular challenge to program. They
often require us to have multiple copies of data, which need to keep
synchronized. Yet we cannot rely on processing nodes working reliably, and
network delays can easily lead to inconsistencies. Despite this, many
organizations rely on a range of core distributed software handling data
storage, messaging, system management, and compute capability. These systems
face common problems which they solve with similar solutions. This article
recognizes and develops these solutions as patterns, with which we can build
up an understanding of how to better understand, communicate and teach
distributed system design._

17 June 2021

* * *

Unmesh Joshi is a Principal Consultant at Thoughtworks.
He is a software architecture enthusiast, who believes that understanding principles of distributed systems
is as essential today as understanding web architecture or object oriented programming was
in the last decade.


## Contents

- [What this is about](http://martinfowler.com#WhatThisIsAbout)
  - [Distributed systems - An implementation perspective](http://martinfowler.com#DistributedSystems-AnImplementationPerspective)
    - [Patterns](http://martinfowler.com#Patterns)
- [Problems and Their Recurring Solutions.](http://martinfowler.com#ProblemsAndTheirRecurringSolutions.)
  - [Process crashes](http://martinfowler.com#ProcessCrashes)
  - [Network delays](http://martinfowler.com#NetworkDelays)
  - [Process Pauses](http://martinfowler.com#ProcessPauses)
  - [Unsynchronized Clocks and Ordering Events](http://martinfowler.com#UnsynchronizedClocksAndOrderingEvents)
- [Putting it all together - Pattern Sequences](http://martinfowler.com#PuttingItAllTogether-PatternSequences)
  - [Pattern Sequence for implementing consensus](http://martinfowler.com#PatternSequenceForImplementingConsensus)
  - [Kubernetes or Kafka Control Plane](http://martinfowler.com#KubernetesOrKafkaControlPlane)
  - [Logical Timestamp usage](http://martinfowler.com#LogicalTimestampUsage)
- [Next Steps](http://martinfowler.com#NextSteps)

## What this is about

For the last several months, I have been conducting workshops on distributed systems at Thoughtworks.
One of the key challenges faced while conducting the workshops was how to map the
theory of distributed systems to open source code bases such as Kafka or Cassandra, while
keeping the discussions generic enough to cover a broad range of solutions.
The concept of patterns provided a nice way out.


Pattern structure, by its very nature,
allows us to focus on a specific problem, making it very clear why a particular solution is needed.
Then the solution description enables us to give a code structure, which is concrete enough to show
the actual solution but generic enough to cover a broad range of variations.
This patterns technique also allows us to link various patterns together to build a complete system.
This gives a nice vocabulary to discuss distributed system implementations.


What follows is a first set of patterns observed in mainstream open source distributed systems.
I hope that this set of patterns will be useful to all developers.


### Distributed systems - An implementation perspective

Today's enterprise architecture is full of platforms and frameworks which are distributed by nature.
If we see the sample list of frameworks and platforms used in typical enterprise architecture today,
it will look something like following:


Type of platform/frameworkExampleDatabasesCassandra, HBase, RiakMessage BrokersKafka, PulsarInfrastructureKubernetes, Mesos, Zookeeper, etcd, ConsulIn Memory Data/Compute GridsHazelcast, Pivotal GemfireStateful MicroservicesAkka Actors, AxonFile SystemsHDFS, Ceph

All these are 'distributed' by nature. What does it mean for a system to be distributed?
There are two aspects:


- They run on multiple servers. The number of servers in a cluster can
vary from as few as three servers to a few thousand servers.
- They manage data. So these are inherently 'stateful' systems.

There are several ways in which things can go wrong when multiple servers are involved in storing data.
All the above mentioned systems need to solve those problems.
The implementation of these systems have some recurring solutions to these problems.
Understanding these solutions in their general form helps in understanding
the implementation of the broad spectrum of these systems and
can also serve as a good guidance when new systems need to be built. Enter patterns.


#### Patterns

[Patterns](http://martinfowler.com/articles/writingPatterns.html), a concept introduced by Christopher Alexander,
is widely accepted in the software community to document design constructs which are
used to build software systems. Patterns provide a structured way of
looking at a problem space along with the solutions which are seen multiple times and proven.
An interesting way to use patterns is the ability to link several patterns together,
in a form of pattern sequence or pattern language, which gives some guidance of implementing a ‘whole’ or a complete system.
Looking at distributed systems as a series of patterns is a useful way to gain insights into their implementation.


## Problems and Their Recurring Solutions.

The following examples from Amazon, Google and GitHub illustrate
how issues can arrive in even the most sophisticated of setups.


This [GitHub outage](https://github.blog/2018-10-30-oct21-post-incident-analysis/)
essentially caused loss of connectivity between its East and West coast data centers.
It caused a small window of time in which data could not be replicated across the data centers,
causing two MySQL servers to have inconsistent data.


This [AWS outage](https://aws.amazon.com/message/41926/), caused by human error where an automation script was wrongly passed a parameter to take down a large number of servers.

This [Google outage](https://status.cloud.google.com/incident/cloud-networking/19009), caused by some misconfiguration, caused a significant impact on the network capacity causing network congestion and service disruption.

Several things can go wrong when data is stored on multiple servers.


### Process crashes

Processes can crash at any time maybe due to hardware faults or software faults.
There are numerous ways in which a process can crash.


- It can be taken down for routine maintenance by system administrators.
- It can be killed doing some file IO because the disk is full and the exception is not properly handled.
- In cloud environments, it can be even trickier, as some unrelated events can bring the servers down.

The bottom line is that if the processes are responsible for storing data, they must be designed to give a durability guarantee for the data stored on the servers.
Even if a process crashes abruptly, it should preserve all the data for which it has notified the user that it's stored successfully.
Depending on the access patterns, different storage engines have different storage structures,
ranging from a simple hash map to a sophisticated graph storage.
Because flushing data to the disk is one of the most time consuming operations,
not every insert or update to the storage can be flushed to disk.
So most databases have in-memory storage structures which are only periodically flushed to disk.
This poses a risk of losing all the data if the process abruptly crashes.


A technique called [Write-Ahead Log](http://martinfowler.com/wal.html) is used to tackle this situation.
Servers store each state change as a command in an append-only file on a hard disk.
Appending a file is generally a very fast operation, so it can be done without impacting performance.
A single log, which is appended sequentially, is used to store each update.
At the server startup, the log can be replayed to build in memory state again.


This gives a durability guarantee. The data will not get lost even if the server abruptly crashes
and then restarts.
But clients will not be able to get or store any data till the server is back up.
So we lack availability in the case of server failure.


One of the obvious solutions is to store the data on multiple servers.
So we can replicate the write ahead log on multiple servers.


When multiple servers are involved, there are a lot more failure scenarios which need to be considered.


### Network delays

In the TCP/IP protocol stack, there is no upper bound on delays caused in transmitting messages across a network.
It can vary based on the load on the network. For example, a 1 Gbps network link can get flooded with a big data
job that's triggered, filling the network buffers,
which can cause arbitrary delay for some messages to reach the servers.


In a typical data center, servers are packed together in racks, and there are multiple racks connected by
a top-of-the-rack switch. There might be a tree of switches connecting one part of the data center to the other.
It is possible in some cases, that a set of servers can communicate with each other, but are disconnected from another set of servers. This situation is called a network partition.
One of the fundamental issues with servers communicating over a network then is how to know a particular server has failed.


There are two problems to be tackled here.


- A particular server can not wait indefinitely to know if another server has crashed.
- There should not be two sets of servers, each considering another set to have failed,
and therefore continuing to serve different sets of clients. This is called the split brain.

To tackle the first problem, every server sends a [HeartBeat](http://martinfowler.com/heartbeat.html) message to other servers at a regular interval.
If a heartbeat is missed, the server sending the heartbeat is considered crashed.
The heartbeat interval is small enough to make sure that it does not take a lot of time to detect server failure.
As we will see below, in the worst case scenario, the server might be up and running,
but the cluster as a group can move ahead considering the server to be failing. This makes sure that services provided to clients are not interrupted.


The second problem is the split brain. With the split brain, if two sets of servers accept updates independently,
different clients can get and set different data, and once the split brain is resolved, it's impossible to resolve conflicts automatically.


To take care of the split brain issue, we must ensure that the two sets of servers,
which are disconnected from each other, should not be able to make progress independently.
To ensure this, every action the server takes, is considered successful only if the majority of the servers can confirm the action.
If servers can not get a majority, they will not be able to provide the required services, and some group of the clients might not be receiving the service, but servers in the cluster will always be in a consistent state.
The number of servers making the majority is called a [Quorum](http://martinfowler.com/quorum.html).
How to decide on the quorum? That is decided based on the number of failures the cluster can tolerate.
So if we have a cluster of five nodes, we need a quorum of three.
In general, if we want to tolerate `f` failures we need a cluster size of 2f + 1.


Quorum makes sure that we have enough copies of data to survive some server failures. But it is not enough to give strong consistency guarantees to clients. Lets say a client initiates a write operation on the quorum, but the write operation succeeds only on one server. The other servers in the quorum still have old values. When a client reads the values from the quorum, it might get the latest value, if the server having the latest value is available. But it can very well get an old value if, just when the client starts reading the value, the server with the latest value is not available. To avoid such situations, someone needs to track if the quorum agrees on a particular operation and only send values to clients which are guaranteed to be available on all the servers.
[Leader and Followers](http://martinfowler.com/leader-follower.html) is used in this situation. One of the servers is elected a leader and the other servers act as followers. The leader controls and coordinates the replication on the followers.
The leader now needs to decide, which changes should be made visible to the clients.
A [High-Water Mark](http://martinfowler.com/high-watermark.html) is used to track the entry in the write ahead log
that is known to have successfully replicated to a quorum of followers.
All the entries upto the high-water mark are made visible to the clients.
The leader also propagates the high-water mark to the followers. So in case the leader fails and one of the followers becomes the new leader, there are no inconsistencies in what a client sees.


### Process Pauses

Even with quorums and leader and followers, there is a tricky problem that needs to be solved. Leader processes can pause arbitrarily. There are a lot of reasons a process can pause. For languages which support garbage collection, there can be a long garbage collection pause.
A leader with a long garbage collection pause,
can be disconnected from the followers, and will continue sending messages to followers after the pause is over.
In the meanwhile, because followers did not receive a heartbeat from the leader, they might have elected a new leader
and accepted updates from the clients. If the requests from the old leader are processed as is,
they might overwrite some of the updates. So we need a mechanism to detect requests from out-of-date leaders.
Here [Generation Clock](http://martinfowler.com/generation.html) is used to mark and detect requests from older leaders.
The generation is a number which is monotonically increasing.


### Unsynchronized Clocks and Ordering Events

The problem of detecting older leader messages from newer ones is the problem of maintaining ordering of messages. It might appear that we can use system timestamps to order a set of messages, but we can not.
The main reason we can not use system clocks is that system clocks across servers are not guaranteed to be synchronized.
A time-of-the-day clock in a computer is managed by a quartz crystal and measures time based on the oscillations of the crystal.


This mechanism is error prone, as the crystals can oscillate faster or slower and so different servers can have very different times.
The clocks across a set of servers are synchronized by a service called NTP.
This service periodically checks a set of global time servers, and adjusts the computer clock accordingly.


Because this happens with communication over a network, and network delays can vary as discussed in the above sections, the clock synchronization might be delayed because of a network issue. This can cause server clocks to drift away from each other, and after the NTP sync happens, even move back in time.
Because of these issues with computer clocks, time of day is generally not used for ordering events.
Instead a simple technique called [Lamport Clock](http://martinfowler.com/lamport-clock.html) is used.
[Generation Clock](http://martinfowler.com/generation.html) is an example of that.
Lamport Clocks are just simple numbers, which are incremented only when some event happens in the system.
In a database, the events are about writing and reading the values, so the lamport clock
is incremented only when a value is written. The Lamport Clock numbers are
also passed in the messages sent to other processes.
The receiving process can then select the larger of the two numbers,
the one it receives in the message and the one it maintains.
This way Lamport Clocks also track happens-before relationship between events across processes
which communicate with each other.
An example of this is the servers taking part in a transaction.
While the [Lamport Clock](http://martinfowler.com/lamport-clock.html) allows ordering of events, it does not have
any relation to the time of the day clock. To bridge this gap,
a variation called [Hybrid Clock](http://martinfowler.com/hybrid-clock.html)
is used. The Hybrid Clock uses system time along with a separate number
to make sure the value increases monotonically,
and can be used the same way as Lamport Clock.


The Lamport Clock allows determining the order of events across a set of communicating servers.
But it does not allow detecting concurrent updates to the same value happening across a set of replicas.
[Version Vector](http://martinfowler.com/version-vector.html) is used to detect conflict across a set of replicas.


The Lamport Clock or Version Vector needs to be associated with the stored values, to detect which
values are stored after the other or if there are conflicts.
So the servers store the values as [Versioned Value](http://martinfowler.com/versioned-value.html).


## Putting it all together - Pattern Sequences

We can see how understanding these patterns, helps us build a complete
system, from the ground up. We will take consensus implementation as an
example. Distributed Consensus is a special case of distributed system
implementation, which provides the strongest consistency guarantee. Common
examples seen in popular enterprise systems are, [Zookeeper](https://zookeeper.apache.org/), [etcd](https://etcd.io/) and [Consul](https://www.consul.io/). They implement consensus algorithms such as
[zab](https://zookeeper.apache.org/doc/r3.4.13/zookeeperInternals.html#sc_atomicBroadcast) and [Raft](https://raft.github.io/) to provide
replication and strong consistency. There are other popular algorithms to
implement consensus, [Paxos](https://en.wikipedia.org/wiki/Paxos_(computer_science)) which is used in
Google's [Chubby](https://research.google/pubs/pub27897/) locking service, view stamp
replication and [virtual-synchrony](https://www.cs.cornell.edu/ken/History.pdf).
In very simple terms, Consensus refers to a set of servers which agree on
stored data, the order in which the data is stored and when to make that
data visible to the clients.

### Pattern Sequence for implementing consensus

Consensus implementations use [state machine replication](https://en.wikipedia.org/wiki/State_machine_replication) to achieve fault tolerance.
In state machine replication, the storage services, like a key value store, are replicated on all the servers,
and the user inputs are executed in the same order on each server.
The key implementation technique used to achieve this is to
replicate [Write-Ahead Log](http://martinfowler.com/wal.html) on all the servers to have a 'Replicated Wal'.


We can put the patterns together to implement Replicated Wal as follows.


For providing durability guarantees, you can use the [Write-Ahead Log](http://martinfowler.com/wal.html) pattern.
The Write Ahead Log is divided into multiple segments using [Segmented Log](http://martinfowler.com/log-segmentation.html).
This helps with log cleaning which is handled by [Low-Water Mark](http://martinfowler.com/low-watermark.html).
Fault tolerance is provided by replicating the write ahead log on multiple servers.
The replication amongst the servers is managed by using the [Leader and Followers](http://martinfowler.com/leader-follower.html) pattern
and [Quorum](http://martinfowler.com/quorum.html) is used to update the [High-Water Mark](http://martinfowler.com/high-watermark.html)
to decide which values are visible to clients.
All the requests are processed in strict order, by using [Singular Update Queue](http://martinfowler.com/singular-update-queue.html).
The order is maintained while sending the requests from leaders to followers using
[Single Socket Channel](http://martinfowler.com/single-socket-channel.html). To optimize for throughput and latency over
a single socket channel,
[Request Pipeline](http://martinfowler.com/request-pipeline.html) can be used.
Followers know about availability of the leader via the [HeartBeat](http://martinfowler.com/heartbeat.html)
received from the leader.
If the leader is temporarily disconnected from the cluster because of network partition,
it is detected by using [Generation Clock](http://martinfowler.com/generation.html).
If all the requests are served only by the leader, it might get overloaded.
When the clients are read only and tolerate reading stale values,
they can be served by the follower servers. [Follower Reads](http://martinfowler.com/follower-reads.html) allows
handling read requests from follower servers.


[Write-Ahead Log](http://martinfowler.com/wal.html) [Segmented Log](http://martinfowler.com/log-segmentation.html) [Low-Water Mark](http://martinfowler.com/low-watermark.html) [Replicated Log](http://martinfowler.com/replicated-log.html) [High-Water Mark](http://martinfowler.com/high-watermark.html) [Quorum](http://martinfowler.com/quorum.html) [Leader and Followers](http://martinfowler.com/leader-follower.html) [HeartBeat](http://martinfowler.com/heartbeat.html) [Single Socket Channel](http://martinfowler.com/single-socket-channel.html) [Generation Clock](http://martinfowler.com/generation.html) [Follower Reads](http://martinfowler.com/follower-reads.html) [Request Pipeline](http://martinfowler.com/request-pipeline.html) [Singular Update Queue](http://martinfowler.com/singular-update-queue.html)

### Kubernetes or Kafka Control Plane

Products like [Kubernetes](https://kubernetes.io/) or
[Kafka's](https://kafka.apache.org/) architecture are built around a
strongly consistent metadata store.
We can understand it as a pattern sequence.
[Consistent Core](http://martinfowler.com/consistent-core.html) is used as a strongly consistent,
fault tolerant metadata store.
[Lease](http://martinfowler.com/time-bound-lease.html) is used to implement group membership and
failure detection of cluster nodes.
Cluster nodes use [State Watch](http://martinfowler.com/state-watch.html) to get notified when any cluster
node fails or updates its metadata
The [Consistent Core](http://martinfowler.com/consistent-core.html) implementation uses
[Idempotent Receiver](http://martinfowler.com/idempotent-receiver.html) to ignore duplicate requests sent
by cluster nodes in case of retries on network failure.
The Consistent Core is built with a 'Replicated Wal', which is described
as a pattern sequence in the above section.


[Replicated Log](http://martinfowler.com/replicated-log.html) [Consistent Core](http://martinfowler.com/consistent-core.html) [Lease](http://martinfowler.com/time-bound-lease.html) [State Watch](http://martinfowler.com/state-watch.html) [Idempotent Receiver](http://martinfowler.com/idempotent-receiver.html)

### Logical Timestamp usage

Usage of various types of [logical timestamps](https://en.wikipedia.org/wiki/Logical_clock)
can also be seen as a pattern sequence.
Various products use either a [Gossip Dissemination](http://martinfowler.com/gossip-dissemination.html)
or a [Consistent Core](http://martinfowler.com/consistent-core.html) for group membership and failure detection of cluster nodes.
The data storage uses [Versioned Value](http://martinfowler.com/versioned-value.html) to be able to determine which values are most recent.
If a single server is responsible for updating the values or [Leader and Followers](http://martinfowler.com/leader-follower.html) is used,
then a [Lamport Clock](http://martinfowler.com/lamport-clock.html) can be used as a version, in the [Versioned Value](http://martinfowler.com/versioned-value.html).
When the timestamp values need to be derived from the time of the day,
a [Hybrid Clock](http://martinfowler.com/hybrid-clock.html) is used instead of a simple Lamport Clock.
If multiple servers are allowed to handle client requests to update the same value,
a [Version Vector](http://martinfowler.com/version-vector.html) is used to be able to detect concurrent writes on different
cluster nodes.


[Gossip Dissemination](http://martinfowler.com/gossip-dissemination.html) [Consistent Core](http://martinfowler.com/consistent-core.html) [Versioned Value](http://martinfowler.com/versioned-value.html) [Version Vector](http://martinfowler.com/version-vector.html) [Lamport Clock](http://martinfowler.com/lamport-clock.html) [Hybrid Clock](http://martinfowler.com/hybrid-clock.html)

This way, understanding problems and their recurring solutions in their general form, helps in understanding building blocks of a complete system

## Next Steps


Distributed systems is a vast topic.
The set of patterns covered here is a small part, covering different categories to showcase how a patterns approach can help understand and design distributed systems.
I will keep adding to this set to broadly include the following categories of problems solved in any distributed system


- Group Membership and Failure Detection
- Partitioning
- Replication and Consistency
- Storage
- Processing

This page is part of: Patterns of Distributed Systems

## Acknowledgements

Many thanks to Martin Fowler for helping me throughout and guiding me to think in terms of patterns.

Mushtaq Ahemad helped me with good feedback and a lot of discussions throughout

Rebecca Parsons, Dave Elliman, Samir Seth, Prasanna Pendse, Santosh Mahale, Sarthak Makhija, James Lewis,
Chris Ford, Kumar Sankara Iyer, Evan Bottcher,Ian Cartwright provided feedback on the earlier drafts.

Professor Indranil Gupta provided feedback on the gossip dissemination pattern.

Thanks to Jojo Swords, Gareth Morgan for helping with copy editing.

Significant Revisions

_17 June 2021:_ Started publication of third batch of patterns based on Gossip.

_05 January 2021:_ Started publication of second batch with: Consistent Core,
Lease, State Watch, and Idempotent Receiver.

_04 August 2020:_ Initial publication with Generation Clock and
Heartbeat patterns. Some patterns then added over next few weeks
