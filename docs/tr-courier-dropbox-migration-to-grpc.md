# Courier: Dropbox migration to gRPC

# Courier：Dropbox 迁移到 gRPC

By Ruslan Nigmatullin and Alexey Ivanov • Jan 08, 2019


1. [The road to gRPC](http://dropbox.tech#the-road-to-grpc)
2. [What Courier brings to gRPC](http://dropbox.tech#what-courier-brings-to-grpc)
3. [Performance optimizations](http://dropbox.tech#performance-optimizations)
4. [Implementation details](http://dropbox.tech#implementation-details)
5. [Migration process](http://dropbox.tech#migration-process)
6. [Lessons learned](http://dropbox.tech#lessons-learned)
7. [Future Work](http://dropbox.tech#future-work)
8. [We are hiring!](http://dropbox.tech#we-are-hiring)

---

1. [gRPC之路](http://dropbox.tech#the-road-to-grpc)
2. [Courier 为 gRPC 带来了什么](http://dropbox.tech#what-courier-brings-to-grpc)
3. [性能优化](http://dropbox.tech#performance-optimizations)
4. [实现细节](http://dropbox.tech#implementation-details)
5. [迁移流程](http://dropbox.tech#migration-process)
6. [经验教训](http://dropbox.tech#lessons-learned)
7. [未来工作](http://dropbox.tech#future-work)
8. [我们正在招聘！](http://dropbox.tech#we-are-hiring)

Dropbox runs hundreds of services, written in different languages, which exchange millions of requests per second. At the core of our Service Oriented Architecture is Courier, our gRPC-based Remote Procedure Call (RPC) framework. While developing Courier, we learned a lot about extending gRPC, optimizing performance for scale, and providing a bridge from our legacy RPC system.

Dropbox 运行数百个用不同语言编写的服务，每秒交换数百万个请求。我们面向服务架构的核心是 Courier，这是我们基于 gRPC 的远程过程调用 (RPC) 框架。在开发 Courier 的过程中，我们学到了很多关于扩展 gRPC、优化规模性能以及为我们的旧 RPC 系统提供桥梁的知识。

_Note: this post shows code generation examples in Python and Go. We also support Rust and Java._

_注意：这篇文章展示了 Python 和 Go 中的代码生成示例。我们也支持 Rust 和 Java。_

## The road to gRPC

## gRPC 之路

Courier is not Dropbox’s first RPC framework. Even before we started to break our Python monolith into services in earnest, we needed a solid foundation for inter-service communication. Especially since the choice of the RPC framework has profound reliability implications.

Courier 不是 Dropbox 的第一个 RPC 框架。甚至在我们开始认真地将 Python 单体分解为服务之前，我们就需要为服务间通信打下坚实的基础。特别是因为 RPC 框架的选择具有深远的可靠性影响。

Previously, Dropbox experimented with multiple RPC frameworks. At first, we started with a custom protocol for manual serialization and de-serialization. Some services like our [Scribe-based log pipeline](https://blogs.dropbox.com/tech/2015/05/how-to-write-a-better-scribe/) used [Apache Thrift](https://github.com/apache/thrift). But our main RPC framework (legacy RPC) was an HTTP/1.1-based protocol with protobuf-encoded messages.

此前，Dropbox 试验了多个 RPC 框架。起初，我们从一个用于手动序列化和反序列化的自定义协议开始。一些服务，比如我们的 [基于 Scribe 的日志管道](https://blogs.dropbox.com/tech/2015/05/how-to-write-a-better-scribe/) 使用了 [Apache Thrift](https:/ /github.com/apache/thrift）。但是我们的主要 RPC 框架（旧式 RPC)是一个基于 HTTP/1.1 的协议，带有 protobuf 编码的消息。

For our new framework, there were several choices. We could evolve the legacy RPC framework to incorporate Swagger (now [OpenAPI](https://github.com/OAI/OpenAPI-Specification)). Or we could [create a new standard](https://xkcd.com/927/). We also considered building on top of both Thrift and gRPC.

对于我们的新框架，有多种选择。我们可以发展遗留的 RPC 框架以合并 Swagger（现在 [OpenAPI](https://github.com/OAI/OpenAPI-Specification))。或者我们可以[创建一个新标准](https://xkcd.com/927/)。我们还考虑在 Thrift 和 gRPC 之上构建。

We settled on gRPC primarily because it allowed us to bring forward our existing protobufs. For our use cases, multiplexing HTTP/2 transport and bi-directional streaming were also attractive.

我们选择 gRPC 主要是因为它允许我们提出我们现有的 protobuf。对于我们的用例，多路复用 HTTP/2 传输和双向流也很有吸引力。

> Note that if [fbthrift](https://github.com/facebook/fbthrift) had existed at the time, we may have taken a closer look at Thrift based solutions.

> 请注意，如果 [fbthrift](https://github.com/facebook/fbthrift) 当时已经存在，我们可能已经仔细研究了基于 Thrift 的解决方案。

## What Courier brings to gRPC

## Courier 为 gRPC 带来了什么

Courier is not a different RPC protocol—it’s just how Dropbox integrated gRPC with our existing infrastructure. For example, it needs to work with our specific versions of authentication, authorization, and service discovery. It also needs to integrate with our stats, event logging, and tracing tools. The result of all that work is what we call Courier.

Courier 并不是一种不同的 RPC 协议——它只是 Dropbox 将 gRPC 与我们现有的基础设施集成的方式。例如，它需要与我们特定版本的身份验证、授权和服务发现配合使用。它还需要与我们的统计信息、事件日志记录和跟踪工具集成。所有这些工作的结果就是我们所说的 Courier。

> While we support using [Bandaid](https://blogs.dropbox.com/tech/2018/03/meet-bandaid-the-dropbox-service-proxy/) as a gRPC proxy for a few specific use cases, the majority of our services communicate with each other with no proxy, to minimize the effect of the RPC on serving latency.

> 虽然我们支持使用 [Bandaid](https://blogs.dropbox.com/tech/2018/03/meet-bandaid-the-dropbox-service-proxy/) 作为一些特定用例的 gRPC 代理，但我们的大多数服务在没有代理的情况下相互通信，以最大限度地减少 RPC 对服务延迟的影响。

We want to minimize the amount of boilerplate we write. Since Courier is our common framework for service development, it incorporates features which all services need. Most of these features are enabled by default, and can be controlled by command-line arguments. Some of them can also be toggled dynamically via a feature flag.

我们希望尽量减少我们编写的样板数量。由于 Courier 是我们用于服务开发的通用框架，因此它包含了所有服务所需的功能。默认情况下，这些功能中的大多数都是启用的，并且可以通过命令行参数进行控制。其中一些还可以通过功能标志动态切换。

### Security: service identity and TLS mutual authentication

### 安全：服务身份和 TLS 相互认证

Courier implements our standard service identity mechanism. All our servers and clients have their own TLS certificates, which are issued by our internal Certificate Authority. Each one has an identity, encoded in the certificate. This identity is then used for mutual authentication, where the server verifies the client, and the client verifies the server. 

Courier 实现了我们的标准服务标识机制。我们所有的服务器和客户端都有自己的 TLS 证书，由我们的内部证书颁发机构颁发。每个人都有一个身份，编码在证书中。然后将此身份用于相互身份验证，其中服务器验证客户端，客户端验证服务器。

> On the TLS side, where we control both ends of the communication, we enforce quite restrictive defaults. Encryption with [PFS](https://scotthelme.co.uk/perfect-forward-secrecy/) is mandatory for all internal RPCs. The TLS version is pinned to 1.2+. We also restrict symmetric/asymmetric algorithms to a secure subset, with `ECDHE-ECDSA-AES128-GCM-SHA256` being preferred.

> 在 TLS 方面，我们控制通信的两端，我们强制执行非常严格的默认值。所有内部 RPC 都必须使用 [PFS](https://scotthelme.co.uk/perfect-forward-secrecy/) 进行加密。 TLS 版本固定为 1.2+。我们还将对称/非对称算法限制为安全子集，首选“ECDHE-ECDSA-AES128-GCM-SHA256”。

After identity is confirmed and the request is decrypted, the server verifies that the client has proper permissions. Access Control Lists (ACLs) and rate limits can be set on both services and individual methods. They can also be updated via our distributed config filesystem (AFS). This allows service owners to shed load in a matter of seconds, without needing to restart processes. Subscribing to notifications and handling configuration updates is taken care of by the Courier framework.

在确认身份并解密请求后，服务器验证客户端是否具有适当的权限。可以在服务和个别方法上设置访问控制列表 (ACL) 和速率限制。它们也可以通过我们的分布式配置文件系统 (AFS) 进行更新。这允许服务所有者在几秒钟内减轻负载，而无需重新启动进程。 Courier 框架负责订阅通知和处理配置更新。

> Service “Identity” is the global identifier for ACLs, rate limits, stats, and more. As a side bonus, it’s also cryptographically secure.

> 服务“身份”是 ACL、速率限制、统计信息等的全局标识符。作为附带奖励，它也是加密安全的。

Here is an example of Courier ACL/ratelimit configuration definition from our [Optical Character Recognition (OCR) service](https://blogs.dropbox.com/tech/2018/10/using-machine-learning-to-index-text-from-billions-of-images/):

以下是我们的 [光学字符识别 (OCR) 服务](https://blogs.dropbox.com/tech/2018/10/using-machine-learning-to-index-text) 的 Courier ACL/ratelimit 配置定义示例-来自数十亿张图片/)：

```
limits:
dropbox_engine_ocr:
    # All RPC methods.
    default:
      max_concurrency: 32
      queue_timeout_ms: 1000

      rate_acls:
        # OCR clients are unlimited.
        ocr: -1
        # Nobody else gets to talk to us.
        authenticated: 0
        unauthenticated: 0

```


![](http://dropbox.tech/cms/content/dam/dropbox/tech-blog/en-us/2019/01/02-screenshot2018-12-0317.31.03.png)

> We are considering adopting the SPIFFE Verifiable Identity Document (SVID), which is part of [Secure Production Identity Framework for Everyone](https://spiffe.io/)(SPIFFE). This would make our RPC framework compatible with various open source projects.

> 我们正在考虑采用 SPIFFE Verifiable Identity Document (SVID)，它是 [Secure Production Identity Framework for Everyone](https://spiffe.io/)(SPIFFE) 的一部分。这将使我们的 RPC 框架与各种开源项目兼容。

### Observability: stats and tracing

### 可观察性：统计和跟踪

Using just an identity, you can easily locate standard logs, stats, traces, and other useful information about a Courier service.

仅使用身份，您就可以轻松找到有关 Courier 服务的标准日志、统计信息、跟踪和其他有用信息。

![](http://dropbox.tech/cms/content/dam/dropbox/tech-blog/en-us/2019/01/03-screenshot2018-12-0518.03.17.png)

Our code generation adds per-service and per-method stats for both clients and servers. Server stats are broken down by the client identity. Out of the box, we have granular attribution of load, errors, and latency for any Courier service.

我们的代码生成为客户端和服务器添加了每个服务和每个方法的统计信息。服务器统计信息按客户端身份进行细分。开箱即用，我们对任何 Courier 服务的负载、错误和延迟都有详细的归因。

![](http://dropbox.tech/cms/content/dam/dropbox/tech-blog/en-us/2019/01/gw1uztwk.png)

Courier stats include client-side availability and latency, as well as server-side request rates and queue sizes. We also have various break-downs like per-method latency histograms or per-client TLS handshakes.

Courier 统计信息包括客户端可用性和延迟，以及服务器端请求率和队列大小。我们还有各种故障，例如每个方法的延迟直方图或每个客户端的 TLS 握手。

> One of the benefits of having our own code generation is that we can initialize these data structures statically, including histograms and tracing spans. This minimizes the performance impact.

> 拥有自己的代码生成的好处之一是我们可以静态初始化这些数据结构，包括直方图和跟踪跨度。这最大限度地减少了性能影响。

![](http://dropbox.tech/cms/content/dam/dropbox/tech-blog/en-us/2019/01/05-screenshot2018-12-0516.44.06.png)

Our legacy RPC only propagated `request_id` across API boundaries. This allowed joining logs from different services. In Courier, we’ve introduced an API based on a subset of the [OpenTracing](https://opentracing.io/) specification. We wrote our own client libraries, while the server-side is built on top of Cassandra and [Jaeger](https://github.com/jaegertracing/jaeger). The details of how we made this tracing system performant warrant a dedicated blog post.

我们的旧版 RPC 仅跨 API 边界传播 `request_id`。这允许加入来自不同服务的日志。在 Courier 中，我们引入了一个基于 [OpenTracing](https://opentracing.io/) 规范子集的 API。我们编写了自己的客户端库，而服务器端构建在 Cassandra 和 [Jaeger](https://github.com/jaegertracing/jaeger) 之上。我们如何使此跟踪系统具有高性能的详细信息需要专门的博客文章。

![](http://dropbox.tech/cms/content/dam/dropbox/tech-blog/en-us/2019/01/06-screenshot2018-12-0516.35.14.png)

Tracing also gives us the ability to generate a runtime service dependency graph. This helps engineers to understand all the transitive dependencies of a service. It can also potentially be used as a post-deploy check for avoiding unintentional dependencies.

跟踪还使我们能够生成运行时服务依赖关系图。这有助于工程师了解服务的所有传递依赖关系。它还可以潜在地用作部署后检查，以避免无意的依赖关系。

### Reliability: deadlines and circuit-breaking

### 可靠性：截止日期和断路

Courier provides a centralized location for language specific implementations of functionality common to all clients, such as timeouts. Over time, we have added many capabilities at this layer, often as action items from postmortems. 

Courier 为所有客户端通用的功能的语言特定实现提供了一个集中位置，例如超时。随着时间的推移，我们在这一层添加了许多功能，通常作为事后分析的操作项。

**Deadlines** Every [gRPC](https://grpc.io/blog/deadlines) [request includes a](https://grpc.io/blog/deadlines) [deadline](https://grpc.io/blog/deadlines), indicating how long the client will wait for a reply. Since Courier stubs automatically propagate known metadata, the deadline travels with the request even across API boundaries. Within a process, deadlines are converted into a native representation. For example, in Go they are represented by a `context.Context` result from the `WithDeadline` method.

**截止日期** 每个 [gRPC](https://grpc.io/blog/deadlines)[请求包括](https://grpc.io/blog/deadlines) [deadline](https://grpc.io/blog/deadlines)，表示客户端等待回复的时间。由于 Courier 存根会自动传播已知元数据，因此即使跨越 API 边界，截止日期也会随着请求而变化。在流程中，截止日期被转换为本地表示。例如，在 Go 中，它们由 `WithDeadline` 方法的 `context.Context` 结果表示。

In practice, we have fixed whole classes of reliability problems by forcing engineers to define deadlines in their service definitions.

在实践中，我们通过强制工程师在他们的服务定义中定义截止日期来修复整类可靠性问题。

> This context can travel even outside of the RPC layer! For example, our legacy MySQL ORM serializes the RPC context along with the deadline into a comment in the SQL query. Our SQLProxy can parse these comments and `KILL` queries when the deadline is exceeded. As a side benefit, we have per-request attribution when debugging database queries.

> 这个上下文甚至可以在 RPC 层之外传播！例如，我们的旧 MySQL ORM 将 RPC 上下文连同截止日期序列化为 SQL 查询中的注释。我们的 SQLProxy 可以在超过截止日期时解析这些评论和 `KILL` 查询。作为附带的好处，我们在调试数据库查询时有每个请求的属性。

**Circuit-breaking** Another common problem that our legacy RPC clients have to solve is implementing custom exponential backoff and jitter on retries. This is often necessary to prevent cascading overloads from one service to another.

**电路中断** 我们的传统 RPC 客户端必须解决的另一个常见问题是在重试时实现自定义指数退避和抖动。这通常是必要的，以防止从一项服务到另一项服务的级联过载。

In Courier, we wanted to solve circuit-breaking in a more generic way. We started by introducing a LIFO queue between the listener and the workpool.

在 Courier 中，我们希望以更通用的方式解决断路问题。我们首先在侦听器和工作池之间引入了一个 LIFO 队列。

![](http://dropbox.tech/cms/content/dam/dropbox/tech-blog/en-us/2019/01/07-screenshot2018-12-0521.54.58.png)

In the case of a service overload, this LIFO queue acts as an automatic circuit breaker. The queue is not only bounded by size, but critically, it’s also **bounded by time**. A request can only spend so long in the queue.

在服务过载的情况下，此 LIFO 队列充当自动断路器。队列不仅受大小限制，更重要的是，它也**受时间限制**。一个请求只能在队列中花费这么长时间。

> LIFO has the downside of request reordering. If you want to preserve ordering, you can use [CoDel](https://queue.acm.org/detail.cfm?id=2209336). It also has circuit breaking properties, but won’t mess with the order of requests.

> LIFO 有请求重新排序的缺点。如果您想保留顺序，可以使用 [CoDel](https://queue.acm.org/detail.cfm?id=2209336)。它还具有断路特性，但不会扰乱请求的顺序。

![](http://dropbox.tech/cms/content/dam/dropbox/tech-blog/en-us/2019/01/08-screenshot2018-12-0521.54.48.png)

### Introspection: debug endpoints

### 内省：调试端点

Even though debug endpoints are not part of Courier itself, they are widely adopted across Dropbox. They are too useful to not mention! Here are a couple of examples of useful introspections.

尽管调试端点不是 Courier 本身的一部分，但它们在 Dropbox 中被广泛采用。它们太有用了，不能不提！以下是一些有用的内省示例。

> For security reasons, you may want to expose these on a separate port (possibly only on a loopback interface) or even a Unix socket (so access can be additionally controlled with Unix file permissions.) You should also strongly consider using mutual TLS authentication there by asking developers to present their certs to access debug endpoints (esp. non-readonly ones.)

> 出于安全原因，您可能希望在单独的端口（可能仅在环回接口上）或什至 Unix 套接字上公开这些（因此可以使用 Unix 文件权限额外控制访问。）您还应该强烈考虑使用相互 TLS 身份验证通过要求开发人员提供他们的证书来访问调试端点（特别是非只读端点）。

**Runtime** Having the ability to get an insight into the runtime state is a very useful debug feature, e.g. [heap and CPU profiles could be exposed as HTTP or gRPC endpoints](https://golang.org/pkg/net/http/pprof/).

**运行时** 能够深入了解运行时状态是一个非常有用的调试功能，例如[堆和 CPU 配置文件可以作为 HTTP 或 gRPC 端点公开](https://golang.org/pkg/net/http/pprof/)。

> We are planning on using this during the canary verification procedure to automate CPU/memory diffs between old and new code versions.

> 我们计划在金丝雀验证过程中使用它来自动化新旧代码版本之间的 CPU/内存差异。

These debug endpoints can allow modification of runtime state, e.g. a golang-based service can allow dynamically setting the [GCPercent](https://golang.org/pkg/runtime/debug/#SetGCPercent).

这些调试端点可以允许修改运行时状态，例如基于 golang 的服务可以允许动态设置 [GCPercent](https://golang.org/pkg/runtime/debug/#SetGCPercent)。

**Library** For a library author being able to automatically export some library-specific data as an RPC-endpoint may be quite useful. Good examples here is that [malloc library can dump its internal stats](http://jemalloc.net/jemalloc.3.html#malloc_stats_print_opts). Another example is a read/write debug endpoint to change the logging level of a service on the fly.

**库** 对于库作者而言，能够自动导出某些库特定数据作为 RPC 端点可能非常有用。这里的好例子是 [malloc 库可以转储其内部统计信息](http://jemalloc.net/jemalloc.3.html#malloc_stats_print_opts)。另一个示例是读/写调试端点，用于动态更改服务的日志记录级别。

**RPC** It is given that troubleshooting encrypted and binary-encoded protocols will be a bit complicated, therefore putting in as much instrumentation as performance allows in the RPC layer itself is the right thing to do. One example of such an introspection API is a recent [channelz proposal for the gRPC](https://github.com/grpc/proposal/blob/master/A14-channelz.md).

**RPC** 假定对加密和二进制编码协议进行故障排除会有点复杂，因此在 RPC 层本身中尽可能多地进行性能测试是正确的做法。这种内省 API 的一个例子是最近的 [针对 gRPC 的 channelz 提案](https://github.com/grpc/proposal/blob/master/A14-channelz.md)。

**Application** Being able to view application-level parameters can also be useful. A good example is a generalized application info endpoint with build/source hash, command line, etc. This can be used by the orchestration system to verify the consistency of a service deployment.

**应用程序** 能够查看应用程序级参数也很有用。一个很好的例子是具有构建/源哈希、命令行等的通用应用程序信息端点。编排系统可以使用它来验证服务部署的一致性。

## Performance optimizations

## 性能优化

We discovered a handful of Dropbox specific performance bottlenecks when rolling out gRPC at scale.

在大规模推出 gRPC 时，我们发现了一些 Dropbox 特定的性能瓶颈。

### TLS handshake overhead 

### TLS 握手开销

With a service that handles lots of connections, the cumulative CPU overhead of TLS handshakes can become non-negligible. This is especially true during mass service restarts.

对于处理大量连接的服务，TLS 握手的累积 CPU 开销可能变得不可忽略。在大量服务重新启动期间尤其如此。

We switched from RSA 2048 keypairs to ECDSA P-256 to get better performance for signing operations. Here are BoringSSL performance examples (note that RSA is still faster for signature verification):

我们从 RSA 2048 密钥对切换到 ECDSA P-256 以获得更好的签名操作性能。以下是 BoringSSL 性能示例（请注意，RSA 在签名验证方面仍然更快）：

RSA:

```
𝛌 ~/c0d3/boringssl bazel run -- //:bssl speed -filter 'RSA 2048'
Did ... RSA 2048 signing operations in ..............  (1527.9 ops/sec)
Did ... RSA 2048 verify (same key) operations in .... (37066.4 ops/sec)
Did ... RSA 2048 verify (fresh key) operations in ... (25887.6 ops/sec)

```


ECDSA:

```
𝛌 ~/c0d3/boringssl bazel run -- //:bssl speed -filter 'ECDSA P-256'
Did ... ECDSA P-256 signing operations in ... (40410.9 ops/sec)
Did ... ECDSA P-256 verify operations in .... (17037.5 ops/sec)

```


> Since RSA 2048 verification is ~3x faster than ECDSA P-256 one, from a performance perspective, you may consider using RSA for your root/leaf certs. From a security perspective though it’s a bit more complicated since you’ll be chaining different security primitives and therefore resulting security properties will be the minimum of all of them. For the same performance reasons you should also think twice before using RSA 4096 (and higher) certs for your root/leaf certs.

> 由于 RSA 2048 验证比 ECDSA P-256 验证快约 3 倍，因此从性能角度来看，您可以考虑将 RSA 用于根/叶证书。从安全的角度来看，虽然它有点复杂，因为您将链接不同的安全原语，因此产生的安全属性将是所有这些中的最小值。出于相同的性能原因，您在为根/叶证书使用 RSA 4096（及更高版本）证书之前也应该三思而后行。

We also found that TLS library choice (and compilation flags) matter a lot for both performance and security. For example, here is a comparison of MacOS X Mojave’s LibreSSL build vs homebrewed OpenSSL on the same hardware:

我们还发现 TLS 库的选择（和编译标志）对性能和安全性都很重要。例如，这是在相同硬件上 MacOS X Mojave 的 LibreSSL 构建与自制 OpenSSL 的比较：

LibreSSL 2.6.4:

```
𝛌 ~ openssl speed rsa2048
LibreSSL 2.6.4
...
                  sign    verify    sign/s verify/s
rsa 2048 bits 0.032491s 0.001505s     30.8    664.3

```

OpenSSL 1.1.1a：

```
𝛌 ~ openssl speed rsa2048
OpenSSL 1.1.1a  20 Nov 2018
...
                  sign    verify    sign/s verify/s
rsa 2048 bits 0.000992s 0.000029s   1208.0  34454.8

```


But the fastest way to do a TLS handshake is to not do it at all! [We’ve modified gRPC-core and gRPC-python](https://github.com/grpc/grpc/issues/14425) to support session resumption, which made service rollout way less CPU intensive.

但是进行 TLS 握手的最快方法是根本不这样做！ [我们修改了 gRPC-core 和 gRPC-python](https://github.com/grpc/grpc/issues/14425) 以支持会话恢复，这使得服务推出的 CPU 密集度降低。

### Encryption is not expensive

### 加密并不昂贵

It is a common misconception that encryption is expensive. Symmetric encryption is actually blazingly fast on modern hardware. A desktop-grade processor is able to encrypt and authenticate data at 40Gbps rate on a single core:

认为加密很昂贵是一种常见的误解。对称加密在现代硬件上实际上非常快。桌面级处理器能够在单核上以 40Gbps 的速率加密和验证数据：

```
𝛌 ~/c0d3/boringssl bazel run -- //:bssl speed -filter 'AES'
Did ... AES-128-GCM (8192 bytes) seal operations in ... 4534.4 MB/s

```


Nevertheless, we did end up having to tune gRPC for our [50Gb/s storage boxes](https://blogs.dropbox.com/tech/2018/06/extending-magic-pocket-innovation-with-the-first-petabyte-scale-smr-drive-deployment/). We learned that when the encryption speed is comparable to the memory copy speed, reducing the number of `memcpy` operations was critical. In addition, we also made [some of the changes to gRPC itself](https://github.com/grpc/grpc/issues/14058).

尽管如此，我们最终还是不得不为我们的 [50Gb/s 存储盒](https://blogs.dropbox.com/tech/2018/06/extending-magic-pocket-innovation-with-the-first-) 调整 gRPC PB-scale-smr-drive-deployment/)。我们了解到，当加密速度与内存复制速度相当时，减少“memcpy”操作的数量至关重要。此外，我们还对 [gRPC 本身进行了一些更改](https://github.com/grpc/grpc/issues/14058)。

> Authenticated and encrypted protocols have caught many tricky hardware issues. For example, processor, DMA, and network data corruptions. Even if you are not using gRPC, using TLS for internal communication is always a good idea.

> 经过身份验证和加密的协议已经解决了许多棘手的硬件问题。例如，处理器、DMA 和网络数据损坏。即使您不使用 gRPC，使用 TLS 进行内部通信也始终是一个好主意。

### High Bandwidth-Delay product links 

### 高带宽延迟产品链接

Dropbox has [multiple data centers connected through a backbone network](https://blogs.dropbox.com/tech/2017/09/infrastructure-update-evolution-of-the-dropbox-backbone-network/). Sometimes nodes from different regions need to communicate with each other over RPC, e.g. for the purposes of replication. When using TCP the kernel is responsible for limiting the amount of data inflight for a given connection (within the limits of `/proc/sys/net/ipv4/tcp_{r,w}mem`), though since gRPC is HTTP/2 -based it also has its own flow control on top of TCP. [The upper bound for the BDP is hardcoded in](https://github.com/grpc/grpc-go/issues/2400) [grpc-go to 16Mb](https://github.com/grpc/grpc-go/issues/2400), which can become a bottleneck for a single high BDP connection.

Dropbox 有[通过骨干网连接的多个数据中心](https://blogs.dropbox.com/tech/2017/09/infrastructure-update-evolution-of-the-dropbox-backbone-network/)。有时来自不同区域的节点需要通过 RPC 相互通信，例如用于复制的目的。当使用 TCP 时，内核负责限制给定连接的传输数据量（在`/proc/sys/net/ipv4/tcp_{r,w}mem` 的限制内)，尽管因为 gRPC 是 HTTP/2基于它在 TCP 之上也有自己的流量控制。 [BDP 的上限被硬编码在](https://github.com/grpc/grpc-go/issues/2400) [grpc-go to 16Mb](https://github.com/grpc/grpc-go/issues/2400)，这可能成为单个高 BDP 连接的瓶颈。

### Golang’s net.Server vs grpc.Server

### Golang 的 net.Server 与 grpc.Server

In our Go code we initially supported both HTTP/1.1 and gRPC using the same [net.Server](https://golang.org/pkg/net/http/#Server). This was logical from the code maintenance perspective but had suboptimal performance. Splitting HTTP/1.1 and gRPC paths to be processed by separate servers and switching gRPC to [grpc.Server](https://godoc.org/google.golang.org/grpc#Server) greatly improved throughput and memory usage of our Courier services.

在我们的 Go 代码中，我们最初使用相同的 [net.Server](https://golang.org/pkg/net/http/#Server) 支持 HTTP/1.1 和 gRPC。从代码维护的角度来看，这是合乎逻辑的，但性能欠佳。拆分 HTTP/1.1 和 gRPC 路径由单独的服务器处理并将 gRPC 切换到 [grpc.Server](https://godoc.org/google.golang.org/grpc#Server) 大大提高了我们 Courier 的吞吐量和内存使用率服务。

### golang/protobuf vs gogo/protobuf

Marshaling and unmarshaling can be expensive when you switch to gRPC. For our Go code, we’ve switched to [gogo/protobuf](https://github.com/gogo/protobuf) which noticeably decreased CPU usage on our busiest Courier servers.

当您切换到 gRPC 时，编组和解组可能会很昂贵。对于我们的 Go 代码，我们已切换到 [gogo/protobuf](https://github.com/gogo/protobuf)，这显着降低了我们最繁忙的 Courier 服务器上的 CPU 使用率。

> As always,
>  [there are some caveats around using gogo/protobuf](https://jbrandhorst.com/post/gogoproto/), but if you stick to a sane subset of functionality you should be fine.

> 一如既往，
> [使用 gogo/protobuf 时有一些注意事项](https://jbrandhorst.com/post/gogoproto/)，但如果您坚持使用合理的功能子集，您应该没问题。

## Implementation details

## 实现细节

Starting from here, we are going to dig way deeper into the guts of Courier, looking at protobuf schemas and stub examples from different languages. For all the examples below we are going to use our `Test` service (the service we use in Courier’s integration tests).

从这里开始，我们将深入挖掘 Courier 的内部结构，查看来自不同语言的 protobuf 模式和存根示例。对于下面的所有示例，我们将使用我们的“Test”服务（我们在 Courier 的集成测试中使用的服务）。

### Service description

###  服务说明

Let’s look at the snippet from the `Test` service definition:

让我们看一下“Test”服务定义中的片段：

```
service Test {
    option (rpc_core.service_default_deadline_ms) = 1000;

    rpc UnaryUnary(TestRequest) returns (TestResponse) {
        option (rpc_core.method_default_deadline_ms) = 5000;
    }

    rpc UnaryStream(TestRequest) returns (stream TestResponse) {
        option (rpc_core.method_no_deadline) = true;
    }
    ...
}

```


As was mentioned in the reliability section above, deadlines are mandatory for all Courier methods. They can be set for the whole service with the following protobuf option:

正如上面可靠性部分所述，所有 Courier 方法都必须有截止日期。可以使用以下 protobuf 选项为整个服务设置它们：

```
option (rpc_core.service_default_deadline_ms) = 1000;

```


Each method can also set its own deadline, overriding the service-wide one (if present).

每种方法还可以设置自己的截止日期，覆盖服务范围的截止日期（如果存在）。

```
option (rpc_core.method_default_deadline_ms) = 5000;

```


In rare cases where deadline doesn’t really make sense (such as a method to watch some resource), the developer is allowed to explicitly disable it:

在截止日期没有真正意义的极少数情况下（例如观察某些资源的方法），开发人员可以明确禁用它：

```
option (rpc_core.method_no_deadline) = true;

```


The real service definition is also expected to have extensive API documentation, sometimes even along with usage examples.

真正的服务定义也应该有大量的 API 文档，有时甚至还有使用示例。

### Stub generation

### 存根生成

Courier generates its own stubs instead of relying on interceptors (except for the Java case, where the interceptor API is powerful enough) mainly because it gives us more flexibility. Let’s compare our stubs to the default ones using Golang as an example.

Courier 生成自己的存根而不是依赖拦截器（Java 情况除外，其中拦截器 API 足够强大）主要是因为它为我们提供了更大的灵活性。让我们以 Golang 为例，将我们的存根与默认存根进行比较。

This is what default gRPC server stubs look like:

这是默认的 gRPC 服务器存根的样子：

```
func _Test_UnaryUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
        in := new(TestRequest)
        if err := dec(in);err != nil {
                return nil, err
        }
        if interceptor == nil {
                return srv.(TestServer).UnaryUnary(ctx, in)
        }
        info := &grpc.UnaryServerInfo{
                Server:     srv,
                FullMethod: "/test.Test/UnaryUnary",
        }
        handler := func(ctx context.Context, req interface{}) (interface{}, error) {
                return srv.(TestServer).UnaryUnary(ctx, req.(*TestRequest))
        }
        return interceptor(ctx, in, info, handler)
}

```


Here, all the processing happens inline: decoding the protobuf, running interceptors, and calling the `UnaryUnary` handler itself.

在这里，所有处理都内联进行：解码 protobuf，运行拦截器，并调用 `UnaryUnary` 处理程序本身。

Now let’s look at Courier stubs:

现在让我们看看 Courier 存根：

```
func _Test_UnaryUnary_dbxHandler(
        srv interface{},
        ctx context.Context,
        dec func(interface{}) error,
        interceptor grpc.UnaryServerInterceptor) (
        interface{},
        error) {

        defer processor.PanicHandler()

        impl := srv.(*dbxTestServerImpl)
        metadata := impl.testUnaryUnaryMetadata

        ctx = metadata.SetupContext(ctx)
        clientId = client_info.ClientId(ctx)
        stats := metadata.StatsMap.GetOrCreatePerClientStats(clientId)
        stats.TotalCount.Inc()

        req := &processor.UnaryUnaryRequest{
                Srv:            srv,
                Ctx:            ctx,
                Dec:            dec,
                Interceptor:    interceptor,
                RpcStats:       stats,
                Metadata:       metadata,
                FullMethodPath: "/test.Test/UnaryUnary",
                Req:            &test.TestRequest{},
                Handler:        impl._UnaryUnary_internalHandler,
                ClientId:       clientId,
                EnqueueTime:    time.Now(),
        }

        metadata.WorkPool.Process(req).Wait()
        return req.Resp, req.Err
}

```


That’s a lot of code, so let’s go over it line by line.

这是很多代码，所以让我们一行一行地浏览一遍。

First, we defer the panic handler that is responsible for automatic error collection. This allows us to send all uncaught exceptions to centralized storage for later aggregation and reporting:

首先，我们推迟负责自动错误收集的恐慌处理程序。这允许我们将所有未捕获的异常发送到集中存储，以便以后聚合和报告：

```
defer processor.PanicHandler()
```


> One more reason for setting up a custom panic handler is to ensure that we abort application on panic. Default golang/net HTTP handler behavior is to ignore it and continue serving new requests (with potentially corrupted and inconsistent state).

> 设置自定义恐慌处理程序的另一个原因是确保我们在恐慌时中止应用程序。默认的 golang/net HTTP 处理程序行为是忽略它并继续提供新请求（可能存在损坏和不一致的状态）。

Then we propagate context by overriding its values from the metadata of the incoming request:

然后我们通过覆盖传入请求的元数据中的值来传播上下文：

```
ctx = metadata.SetupContext(ctx)
clientId = client_info.ClientId(ctx)

```


We also create (and cache for efficiency purposes) the per-client stats on the server side for more granular attribution:

我们还在服务器端创建（并出于效率目的缓存）每个客户端的统计信息，以获得更精细的归因：

```
stats := metadata.StatsMap.GetOrCreatePerClientStats(clientId)

```


> This dynamically creates a per-client (i.e. per-TLS identity) stats in runtime. We also have per-method stats for each service and, since the stub generator has access to all the methods during the code generation time, we can statically pre-create these to avoid runtime overhead.

> 这会在运行时动态创建每个客户端（即每个 TLS 身份）的统计信息。我们还有每个服务的每个方法的统计信息，并且由于存根生成器可以在代码生成期间访问所有方法，我们可以静态地预创建这些以避免运行时开销。

Then we create the request structure, pass it to the work pool, and wait for the completion:

然后我们创建请求结构，传递给工作池，等待完成：

```
req := &processor.UnaryUnaryRequest{
        Srv:            srv,
        Ctx:            ctx,
        Dec:            dec,
        Interceptor:    interceptor,
        RpcStats:       stats,
        Metadata:       metadata,
        ...
}
metadata.WorkPool.Process(req).Wait()

```


Note that almost no work has been done by this point: no protobuf decoding, no interceptor execution, etc. ACL enforcement, prioritization, and rate-limiting happens inside the workpool before any of that is done.

请注意，此时几乎没有完成任何工作：没有 protobuf 解码、没有拦截器执行等。ACL 实施、优先级划分和速率限制在任何这些完成之前在工作池内发生。

> Note that the
>  [golang gRPC library supports](https://godoc.org/google.golang.org/grpc/tap)[the](https://godoc.org/google.golang.org/grpc/tap) [Tap interface](https://godoc.org/google.golang.org/grpc/tap), which allows very early request interception. This provides infrastructure for building efficient rate-limiters with minimal overhead.

> 请注意，
> [golang gRPC 库支持](https://godoc.org/google.golang.org/grpc/tap)[the](https://godoc.org/google.golang.org/grpc/tap) [Tap interface](https://godoc.org/google.golang.org/grpc/tap)，它允许非常早的请求拦截。这为以最小的开销构建高效的速率限制器提供了基础设施。

### App-specific error codes

### 特定于应用程序的错误代码

Our stub generator also allows developers to define app-specific error codes through custom options:

我们的存根生成器还允许开发人员通过自定义选项定义特定于应用程序的错误代码：

```
enum ErrorCode {
option (rpc_core.rpc_error) = true;

UNKNOWN = 0;
NOT_FOUND = 1 [(rpc_core.grpc_code)="NOT_FOUND"];
ALREADY_EXISTS = 2 [(rpc_core.grpc_code)="ALREADY_EXISTS"];
...
STALE_READ = 7 [(rpc_core.grpc_code)="UNAVAILABLE"];
SHUTTING_DOWN = 8 [(rpc_core.grpc_code)="CANCELLED"];
}

```


Within the same service, both gRPC and app errors are propagated, while between API boundaries all errors are replaced with UNKNOWN. This avoids the problem of accidental error proxying between different services, potentially changing their semantic meaning.

在同一个服务中，gRPC 和应用程序错误都会传播，而在 API 边界之间，所有错误都被替换为 UNKNOWN。这避免了不同服务之间的意外错误代理问题，可能会改变它们的语义。

### Python-specific changes

### 特定于 Python 的更改

Our Python stubs add an explicit context parameter to all Courier handlers, e.g.:

我们的 Python 存根为所有 Courier 处理程序添加了一个明确的上下文参数，例如：

```
from dropbox.context import Context
from dropbox.proto.test.service_pb2 import (
        TestRequest,
        TestResponse,
)
from typing_extensions import Protocol

class TestCourierClient(Protocol):
    def UnaryUnary(
            self,
            ctx,      # type: Context
            request,  # type: TestRequest
            ):
        # type: (...) -> TestResponse
        ...

```


At first, it looked a bit strange, but after some time developers got used to the explicit `ctx` just as they got used to `self`.

起初，它看起来有点奇怪，但一段时间后，开发人员习惯了显式的 `ctx`，就像他们习惯了 `self` 一样。

Note that our stubs are also fully mypy-typed which pays off in full during large-scale refactoring. It also integrates nicely with some IDEs like PyCharm.

请注意，我们的存根也是完全 mypy 类型的，这在大规模重构期间会完全得到回报。它还与 PyCharm 等一些 IDE 很好地集成。

Continuing the static typing trend, we also add mypy annotations to protos themselves:

延续静态类型的趋势，我们还向 protos 本身添加了 mypy 注释：

```
class TestMessage(Message):
    field: int

    def __init__(self,
        field : Optional[int] = ...,
        ) -> None: ...
    @staticmethod
    def FromString(s: bytes) -> TestMessage: ...

```


These annotations prevent many common bugs, such as assigning `None` to a `string` field in Python.

这些注解防止了许多常见的错误，例如在 Python 中将 `None` 分配给 `string` 字段。

This code is opensourced at [dropbox/mypy-protobuf](https://github.com/dropbox/mypy-protobuf).

此代码在 [dropbox/mypy-protobuf](https://github.com/dropbox/mypy-protobuf) 上开源。

## Migration process 

## 迁移过程

Writing a new RPC stack is by no means an easy task, but in terms of operational complexity it still can’t be compared to the process of infra-wide migration to it. To assure the success of this project, we’ve tried to make it easier for the developers to migrate from legacy RPC to Courier. Since the migration by itself is a very error-prone process, we’ve decided to go with a multi-step process.

编写一个新的 RPC 堆栈绝不是一件容易的事，但在操作复杂性方面仍然无法与向它进行范围内迁移的过程相提并论。为了确保这个项目的成功，我们试图让开发人员更容易从传统的 RPC 迁移到 Courier。由于迁移本身是一个非常容易出错的过程，因此我们决定采用多步骤过程。

### Step 0: Freeze the legacy RPC

### 步骤 0：冻结旧版 RPC

Before we did anything, we froze the legacy RPC feature set so it’s no longer a moving target. This also gave people an incentive to move to Courier, since all new features like tracing and streaming were only available to services using Courier.

在我们做任何事情之前，我们冻结了遗留的 RPC 功能集，因此它不再是一个移动目标。这也激发了人们转向 Courier 的动力，因为所有新功能（如跟踪和流媒体）仅适用于使用 Courier 的服务。

### Step 1: A common interface for the legacy RPC and Courier

### 第 1 步：旧版 RPC 和 Courier 的通用接口

We started by defining a common interface for both legacy RPC and Courier. Our code generation was responsible for producing both versions of the stubs that satisfy this interface:

我们首先为旧版 RPC 和 Courier 定义了一个通用接口。我们的代码生成负责生成满足此接口的存根的两个版本：

```
type TestServer interface {
UnaryUnary(
      ctx context.Context,
      req *test.TestRequest) (
      *test.TestResponse,
      error)
...
}

```


### Step 2: Migration to the new interface

### 第 2 步：迁移到新界面

Then we started switching each service to the new interface but continued using legacy RPC. This was often a huge diff touching all the methods in the service and its clients. Since this is the most error-prone step, we wanted to de-risk it as much as possible by changing one variable at a time.

然后我们开始将每个服务切换到新接口，但继续使用旧版 RPC。这通常是涉及服务及其客户中所有方法的巨大差异。由于这是最容易出错的步骤，我们希望通过一次更改一个变量来尽可能降低风险。

> Low profile services with a small number of methods and [spare error budget](https://landing.google.com/sre/sre-book/chapters/embracing-risk/) can do the migration in a single step and ignore this warning.

> 具有少量方法的低调服务和[备用错误预算](https://landing.google.com/sre/sre-book/chapters/embracing-risk/) 可以一步完成迁移并忽略此警告。

### Step 3: Switch clients to use Courier RPC

### 第 3 步：切换客户端以使用 Courier RPC

As part of the Courier migration, we also started running both legacy and Courier servers in the same binary on different ports. Now changing the RPC implementation is a one-line diff to the client:

作为 Courier 迁移的一部分，我们还开始在不同端口上以相同的二进制文件运行旧服务器和 Courier 服务器。现在更改 RPC 实现是对客户端的一行差异：

```
class MyClient(object):
def __init__(self):
-   self.client = LegacyRPCClient('myservice')
+   self.client = CourierRPCClient('myservice')

```


Note that using that model we can migrate one client at a time, starting with ones that have lower SLAs like batch processing and other async jobs.

请注意，使用该模型，我们可以一次迁移一个客户端，从具有较低 SLA 的客户端开始，例如批处理和其他异步作业。

### Step 4: Clean up

### 第 4 步：清理

After all service clients have migrated it is time to prove that legacy RPC is not used anymore (this can be done statically by code inspection and at runtime looking at legacy server stats.) After this step is done developers can proceed to clean up and remove old code.

在所有服务客户端迁移后，是时候证明不再使用旧版 RPC（这可以通过代码检查静态完成，并在运行时查看旧版服务器统计信息。）完成此步骤后，开发人员可以继续清理和删除旧代码。

## Lessons learned

##  得到教训

At the end of the day, what Courier brings to the table is a unified RPC framework that speeds up service development, simplifies operations, and improves Dropbox reliability.

归根结底，Courier 带来的是统一的 RPC 框架，它可以加快服务开发、简化操作并提高 Dropbox 的可靠性。

Here are the main lessons we’ve learned during the Courier development and deployment:

以下是我们在 Courier 开发和部署过程中学到的主要经验教训：

1. Observability is a feature. Having all the metrics and breakdowns out-of-the-box is invaluable during troubleshooting.
2. Standardization and uniformity are important. They lower cognitive load, and simplify operations and code maintenance.
3. Try to minimize the amount of boilerplate code developers need to write. Codegen is your friend here.
4. Make migration as easy as possible. Migration will likely take way more time than the development itself. Also, migration is only finished after cleanup is performed.
5. RPC framework can be a place to add infrastructure-wide reliability improvements, e.g. mandatory deadlines, overload protection, etc. Common reliability issues can be identified by aggregating incident reports on a quarterly basis.

1. 可观察性是一个特征。在故障排除期间，开箱即用的所有指标和故障是非常宝贵的。
2. 标准化和统一性很重要。它们降低了认知负担，并简化了操作和代码维护。
3. 尽量减少开发人员需要编写的样板代码量。 Codegen 是您的朋友。
4. 使迁移尽可能容易。迁移可能比开发本身花费更多的时间。此外，迁移仅在执行清理后完成。
5. RPC 框架可以作为添加基础设施范围可靠性改进的地方，例如强制期限、过载保护等。可以通过按季度汇总事件报告来确定常见的可靠性问题。

## Future Work

##  未来的工作

Courier, as well as gRPC itself, is a moving target so let’s wrap up with the Runtime team and Reliability teams’ roadmaps. 

Courier 以及 gRPC 本身都是一个不断变化的目标，所以让我们总结一下运行时团队和可靠性团队的路线图。

In relatively near future we wanted to add a proper resolver API to Python’s gRPC code, switch to C++ bindings in Python/Rust, and add full circuit breaking and fault injection support. Later next year we are planning on looking into [ALTS and moving TLS handshake to a separate process](https://cloud.google.com/security/encryption-in-transit/application-layer-transport-security/resources/alts-whitepaper.pdf) (possibly even outside of the services' container.)

在相对不久的将来，我们希望向 Python 的 gRPC 代码添加适当的解析器 API，在 Python/Rust 中切换到 C++ 绑定，并添加完整的断路和故障注入支持。明年晚些时候，我们计划研究 [ALTS 并将 TLS 握手移至单独的进程](https://cloud.google.com/security/encryption-in-transit/application-layer-transport-security/resources/alts-whitepaper.pdf)（甚至可能在服务的容器之外。)

## We are hiring!

##  我们正在招聘！

Do you like runtime-related stuff? Dropbox has a globally distributed edge network, terabits of traffic, millions of requests per second, and comfy small teams in both Mountain View and San Francisco.

你喜欢运行时相关的东西吗？ Dropbox 拥有遍布全球的边缘网络、TB 级的流量、每秒数百万的请求，以及在山景城和旧金山的舒适小团队。

![](http://dropbox.tech/cms/content/dam/dropbox/tech-blog/en-us/2019/01/09-screenshot2018-10-0318.04.58.png)

[Traffic/Runtime/Reliability teams are hiring both SWEs and SREs](https://www.dropbox.com/jobs/listing/1233364?gh_src=f80311fa1) to work on TCP/IP packet processors and load balancers, HTTP/gRPC proxies, and our internal service mesh runtime: Courier/gRPC, Service Discovery, and AFS. Not your thing? We're also hiring for [a wide variety of engineering positions in San Francisco, New York, Seattle, Tel Aviv, and other offices around the world](https://www.dropbox.com/jobs/teams/engineering?gh_src=f80311fa1#open-positions).

[流量/运行时/可靠性团队正在招聘 SWE 和 SRE](https://www.dropbox.com/jobs/listing/1233364?gh_src=f80311fa1) 从事 TCP/IP 数据包处理器和负载均衡器、HTTP/gRPC代理和我们的内部服务网格运行时：Courier/gRPC、服务发现和 AFS。不是你的东西？我们还在招聘 [旧金山、纽约、西雅图、特拉维夫和世界各地其他办事处的各种工程职位](https://www.dropbox.com/jobs/teams/engineering?gh_src=f80311fa1#open-positions)。

### Acknowledgments

### 致谢

**Contributors:** Ashwin Amit, Can Berk Guder, Dave Zbarsky, Giang Nguyen, Mehrdad Afshari, Patrick Lee, Ross Delinger, Ruslan Nigmatullin, Russ Allbery, Santosh Ananthakrishnan.

**贡献者：** Ashwin Amit、Can Berk Guder、Dave Zbarsky、Giang Nguyen、Mehrdad Afshari、Patrick Lee、Ross Delinger、Ruslan Nigmatullin、Russ Allbery、Santosh Ananthakrishnan。

We are also very grateful to the gRPC team for their support. 

我们也非常感谢 gRPC 团队的支持。

