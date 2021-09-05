# The History of HAProxy

[Willy Tarreau](https://www.haproxy.com/blog/author/wtarreau/ "Posts by Willy Tarreau") [Willy Tarreau](https://www.haproxy.com/blog/author/wtarreau/ "Posts by Willy Tarreau") \| Nov 8, 2019 \| [LOAD BALANCING / ROUTING](https://www.haproxy.com/blog/category/load-balancing-routing/), [TECH](https://www.haproxy.com/blog/category/tech/) \| [1 comment](https://www.haproxy.com/blog/the-history-of-haproxy/#respond)

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/History-of-HAProxy.png)

Today, HAProxy is the world’s fastest and most widely used software load balancer. It’s responsible for providing high availability, security, and observability to some of the world’s highest trafficked websites. Many Site Reliability Engineers and Network/Systems Engineers alike consider it to be an essential part of their critical infrastructure. However, HAProxy’s origin story is one that has not been told and you may be curious about its roots and what drove it to be what it is today.

All of this started in 1999 when I needed a way to gauge how an application would perform when facing lots of clients with 28 Kbps modems instead of a single 100-Mbps link. I decided to reuse a proxy I’d written earlier for compression, called _Zprox_. I removed the compression code, hacked some bandwidth limitations, and quickly had a useful testing tool.

A year later, I was involved in another benchmarking project where the people in charge of operations were facing an interoperability issue: Their load generating tool emitted HTTP headers in a different syntax than the one expected by the server. I reused my hacked proxy to implement regex-based header rewriting so that the requests were transformed as they passed through. At the same time, I started a minimalistic configuration language, since the proxy needed to listen on multiple ports with different regular expressions. The `listen` and `server` keywords were born—and they’re still supported 19 years later in HAProxy 2.0.

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-1@4x.png)

Then in mid-2001, a heavy application was deployed. This application was known for occasionally exhibiting huge response times, which caused connections to pile up in front of it. This was a situation that the hardware load balancers in place were totally unable to deal with, and they started to repeatedly fail, bringing all other applications down with them. Since it was not possible to fix the application quickly, our short-term workaround consisted of deploying my proxy—which, by then, people were calling “Willy’s proxy”—to offload some of the traffic from the hardware load balancers.

The traffic would pass through the original load balancers working at Layer 4 and then be forwarded to a pair of the proxies, which took over inspecting headers and persisting connections to servers. The server persistence worked by checking whether a cookie was present, then searching for its value among a list of predefined servers, then selecting the matched server. If no cookie was found, the connection was forwarded back to the original load balancers which would contact a server and add the cookie. We estimated that this reduced the Layer 7 work being handled by the original load balancers to a tenth or even a hundredth.

The solution was quick enough to implement using the existing proxy code, as it was already capable of handling HTTP headers. It wouldn’t have to deal with any load-balancing mechanics and only had to detect server availability. Finally, its name had been coined: “HAProxy” for “high availability proxy”. The solution was accepted, quickly implemented, then deployed. Since it was not very stable by then it received updates often, but it never crashed so no user ever noticed the quirks. This was HAProxy 1.0, released in 2001.

## Discovering Its Strengths

Historically, the network team was only in charge of the network equipment and never had to deal with anything related to the HTTP layer or traffic logs. Nevertheless, they were the ones from whom everyone asked for help, since they had a network packet capture device and, honestly, because it was harder for them to prove the problem was not on their side. The infrastructure team was in charge of firewalls and, therefore, had visibility into firewall logs. The applications people became accustomed to considering that anything not working was probably due to a firewall configuration issue. If that wasn’t the case, then they’d ask the network people for a packet capture.

Needless to say, this process does not scale. It also doesn’t work at all for symptoms like “it doesn’t function well enough.” With HAProxy in the hands of the infrastructure team, everything was suddenly simplified: the logs it produced immediately indicated where a connection originated, where it was sent, how long it spent waiting for the server, and which side closed it first. Alone, this was so precise and helpful to each team that it quickly became the first place to look. It was highly trusted, as it significantly reduced the time needed to solve any trouble. It also explains the sophisticated indicators that are included in the default log format.

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-2@4x.png)

The development of logging in HAProxy was also spurred by HTTP interoperability being far from good in 2001. You simply couldn’t trust a server’s logs to visualize what was actually on the wire, and having to take a network trace required a lot of time. HAProxy quickly gained the ability to capture:

- cookies, to diagnose stickiness issues
- headers, to diagnose their issues
- source ports, to correlate with firewalls
- timers, to detect timeouts and application server issues
- termination codes, to analyze failed sessions

Initially, HAProxy didn’t modify anything in the request or response. However, the architecture—two layers of load balancers with HAProxy kicking the connection back to the first layer if no cookie was present—was complex and difficult to configure. So, eventually HAProxy had to learn about load balancing, which meant adding a simple round-robin scheduler, simple health checks, and cookie insertion. At that point, the architecture became cleaner: one layer of L3/L4 load balancers, a layer of HAProxy load balancers handling all the L7 work, then the application servers. HAProxy 1.1 was born, with a codebase three times as large as 1.0.

## A Decisive Coding Standard

With the new responsibilities, performance started to become a concern.

HAProxy was mostly running on old Linux 2.4 systems made of 733 MHz, single-core, Pentium III processors and on a few Solaris 8 systems made of 400 MHz UltraSparc processors. Both were often equipped with 128 MB of RAM, in the best case! On such machines, every byte of memory and every CPU cycle is important. Great care was taken to never use more than necessary of either.

Structure fields were only initialized once and never reinitialized if their state was known. There was little error checking when it was known that the design would not allow certain conditions to happen. Content was parsed in one round, taking care to never read the same byte twice. Such an approach requires clear communication: Often the comments in the code explaining why an operation could be avoided were much longer than the code that was avoided!

Even if a few of these design rules had to be revisited as the code became more modular, they still seeded a certain philosophy in the project that resources aren’t free and developers are not allowed to waste them. We had a mantra: “Keep the developer busy to make the code lazy”.

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-3@4x.png)

Not everything was perfectly designed. For example, we discovered, the hard way (i.e. in production), that under Solaris `select()` was limited to only 256 file descriptors by default; This led us to rush the implementation of the `nbproc` mechanism the same day, in order to start multiple processes, and it was considered a satisfying solution; So, it was only much later that we added support for `poll()`. Nowadays, as people use `nbproc` to improve their scalability, some wonder why certain states are not synchronized between the processes. It’s significant that what they are using was absolutely not designed for the same purpose at all, and that the performance benefit they value is just the byproduct of a nearly two-decades old workaround for a system limitation.

By version 1.2, the project had started to attract outside users. Some hosting providers and large infrastructure providers needed a load balancer for specific services and they found that HAProxy was extremely fast and unbreakable compared to what they’d tried until then. The scope of use cases exploded: faster forwarding from servers to clients for video providers, detailed statistics for hosting providers who had to report to their customers, a CLI to make it easier to get statistics.

## Pushing Killer Performance

Changes continued. It became obvious that backwards compatibility was very important so that users could upgrade without breaking a sweat. More log indicators were needed, and more stats. Then, a new trend of low-performance servers appeared, some limited to only a single connection, others to just a few. Since HAProxy’s connections were cheap, it casually fit the role of being a connection concentrator. The `maxconn` setting was born. Connections could be queued in HAProxy, which acted as a buffer for the servers.

This was a killer feature that protected servers from dying under load. Even better, by reducing the amount of resources a server had to use, it was possible to operate much faster. We held a number of impressive demos alongside popular application servers and showed how placing HAProxy in front sped them up, sometimes by an order of magnitude.

The rapid growth of high-bandwidth websites (i.e. gigabit level) made it mandatory to further think in terms of layers, granting lower layers autonomy so that higher ones could be called less often. This update arrived in HAProxy 1.3.16 and was really the key to pushing HAProxy to where it needed to be to deal with 10 Gbps links. This was thanks to TCP splicing, which, when supported by the network card, allowed data to be proxied between two sides while using only a single copy in memory. By 1.3.17 in 2009, we were saturating a 10G link with video-like traffic using roughly 25% of a Core 2 Duo machine. At the time, it was very difficult to find a motherboard capable of dealing with such bandwidth on the PCIe bridge! So, we were a bit ahead of the times.

The rise of chat sites opened a new challenge: 64,000 ports weren’t enough to accommodate the number of connections to the servers. It became necessary to implement strategies to connect up to 500,000 clients. By then, Linux processes were limited to 1-million file descriptors. We added explicit [source port ranges](https://www.haproxy.com/documentation/hapee/1-9r1/onepage/#source%20(Server%20and%20default-server%20options)) and interface bindings, which allowed us to expand the number of used file descriptors by distributing connections across several servers.

The internal scheduler, which is responsible for running tasks such as processing requests and responses, executing ACLs, and purging data from stick tables, started to become a real problem with its simplistic linked list-based design. It was okay when all tasks had the same timeouts, which decide when the tasks should run, but was lackluster when multiple proxies (i.e. `frontend` and `listen` sections) defined different timeouts.

There had been a workaround, which was to limit the number of distinct timeouts within the configuration in order to save the scheduler from walking too far when inserting tasks. Fortunately, the parallel research on Elastic Binary Trees started to provide solid results and the time to use them for the scheduler soon came. Suddenly, it wasn’t a problem any more to have many timeouts. Watch Andjelko Iharos’ tell the story in his QCon 2019 presentation, [EBtree – Design for a Scheduler and Use (Almost) Everywhere](https://www.infoq.com/presentations/ebtree-design/).

The scheduler was further optimized from the principle that timeouts are almost never met, so we only need to advance them in the tree, but rarely postpone them. For example, we close an inactive connection after a timeout period, but postpone that if the connection is still active. This led to a split of one queue into two: one for the timers and another for the runnable tasks. That way, the tasks are never removed from the timers list until they are truly expired. All of these updates made HAProxy capable of surpassing one million concurrent connections.

## The Keep-Alive Bottleneck

Our next performance bottleneck was the lack of Keep-Alive. Sites had started to be loaded with many resources, such as images, scripts, and stylesheets that caused many requests; Others contacted the server often for dynamic auto-completion. For HAProxy 1.4, we focused on HTTP Keep-Alive and it turned out to be much more complicated than anticipated, since it required HAProxy to close one side of the proxied connection while keeping the other side open. This required that the _close_ event no longer being forwarded, which led to numerous bugs referred to as “CLOSE\_WAIT”.

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-4@4x.png)

Also, keeping connections open with the client forced us to be even more careful about memory usage. At the time, IPv6 support was considered crucial and each connection had a pair of addresses—source and destination—eating 32 extra bytes per side. Even worse, supporting UNIX domain sockets pushed that overhead to 256 bytes per side. A great amount of effort was made to reduce the size of the data structures, carving off parts that weren’t always needed to save memory. HAProxy was mostly used with 8 kB static buffers, which meant that for one million connections, it required 16 GB of memory, or one buffer per direction.

HAProxy 1.4 was released with Keep-Alive only on the client side, leaving server-side Keep-Alive for 1.5. The server side required many changes, and there were a lot of other contributions arriving at the same time that competed for attention. Then, SSL was implemented, requiring an entire redesign of the connection layer. In the end, although it was estimated that server-side Keep-Alive would require six months of development, it actually took 4 ½ years. After that experience, we settled on scheduling releases at fixed dates so that it would be impossible for features to delay a release.

## SSL, Compression and Dynamic Changes

Although version 1.5 brought much needed features like SSL and compression, the memory usage went through the roof. In version 1.6, work was done to allocate buffers dynamically and release them as soon as possible. Meanwhile, the standard compression library, Zlib, used 256 kB of memory per connection. We decided to replace it with an in-house, stateless implementation that, although it compresses slightly less, only requires 8 bytes and is 3x faster. We also introduced server-side connection multiplexing to reduce memory usage, improve performance, and save server resources.

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-5@4x.png)

Then, the code began to evolve towards more dynamic use cases. We added DNS resolution for configuring server addresses; Version 1.7 brought the ability to change almost all settings of a `server` line at runtime by using the [HAProxy Runtime API](https://www.haproxy.com/blog/dynamic-configuration-haproxy-runtime-api/). We also introduced filters that hook into the life of a transaction and provide a way to offload heavy, content-transformation operations to external processes. This we called the [Stream Processing Offload Engine](https://www.haproxy.com/blog/extending-haproxy-with-the-stream-processing-offload-engine/) (SPOE).

Version 1.8 brought multithreading. We expected the first version to not be exceptionally fast, but actually it was way better than anticipated. Not only did it not lose performance (which couldn’t be taken for granted when adding threads in a two-decades old product), but it scaled up HAProxy to be ~2.5x faster on four threads. Adding multithreading broke many old assumptions that applied when you only need to initialize something once. A lot of legacy code was replaced or updated at that point.

HTTP/2 to the client required support for multiplexing and that led to another redesign of the connection layers. A new _mux_ layer was introduced, which was bundled with its own set of surprises and challenges.

With version 1.9, our focus was squarely on maturing features we’d added in 1.8. The scalability of the multithreading feature was improved upon by using a number of atomic operations, some lockfree operations, and other tunings. However, it was known that when the code ran in parallel, the HTTP/2 implementation, which relied on being converted to HTTP/1 before being forwarded to the server, was severely handicapped by the many HTTP header manipulations. This was especially true on the response side when the server was fast enough to send the body with the headers, thus requiring moving this body for each header adjustment.

So, a new internal HTTP representation was created, called HTX, to allow headers to be manipulated out of order without doing expensive `memmove()`. It would also enable transport of both HTTP/1 and HTTP/2 semantics from end to end.

We’d reached a turning point, leaving behind a decades-long era wherein everyone knew all of the code. Now, only a subset of people could deal with their respective, highly advanced parts. This is also a proof of excellence for certain areas, since it has become difficult to find people who can improve them.

## A New Major Version

Version 2.0 did not radically change the technical layers of 1.9, as it mostly focused on adding new features so that end users could benefit from all the work done in 1.9 that, until then, was mostly invisible…and the improvements continue. Each new feature often comes with a number of prerequisites that are adapted to existing features, until it becomes obvious that some of them hinder performance, resource usage, or even code maintainability. At that point, the code in question is redesigned to better address the shortcomings created by the initial feature.

![](https://cdn.haproxy.com/wp-content/uploads/2019/10/Split-6@4x.png)

One example of this increasing level of optimization is the file-descriptor cache that was created in 1.5 as a generalization of the _speculative epoll_ polling mechanism to all pollers. It resulted in a nice simplification of the code and a performance boost on all operating systems. On the other hand, with the introduction of threads in 1.8, it has become a limiting factor. With the layer changes in 2.0, it is no longer necessary. In 2.1 (still in development) the cache has already been removed and, confirming our suspicions, resulted in a 20% higher peak connection rate in certain situations.

## Conclusion

From a vantage outside of the project, one might think that HAProxy is always faster than it used to be. After all, we do always say, “this is faster”. The reality is more nuanced: Each new release accelerates something and introduces a set of features that have little or no impact on the vast majority of setups. However, in Darwinian fashion, modern architecture designs gain influence and what used to be a corner case becomes the norm. Then, the next version focuses on that use case and addresses it. So, each new release of HAProxy is not exactly faster than the previous one; It’s faster than the previous one for a new use case. That’s why it’s important to stay up to date.

While a number of people admit that 1.4 and 1.5 are still very present in their infrastructure, even now in 2019—1.4 is unmaintained and version 1.5 was released five years ago—if their environment isn’t changing, this is fine. Yet, the web adapts very quickly and if your environment keeps up with the evolution, the benefit of having a component like HAProxy, which is always growing with the times, is hard to miss.

What could be said is that HAProxy is always more scalable. It always adapts to faster hardware and to more demanding environments. It never quits progressing.
