# How we migrated Dropbox from Nginx to Envoy

# 我们如何将 Dropbox 从 Nginx 迁移到 Envoy

By Alexey Ivanov and Oleg Guba • Jul 30, 2020

1. [Our legacy Nginx-based traffic infrastructure](http://dropbox.tech#-our-legacy-nginx-based-traffic-infrastructure)
2. [Why not Bandaid?](http://dropbox.tech#-why-not-bandaid)
3. [Our new Envoy-based traffic infrastructure](http://dropbox.tech#-our-new-envoy-based-traffic-infrastructure)
4. [Current state of our migration](http://dropbox.tech#-current-state-of-our-migration)
5. [Issues we encountered](http://dropbox.tech#-issues-we-encountered-)
6. [What’s next?](http://dropbox.tech#-whats-next)
7. [Acknowledgements](http://dropbox.tech#acknowledgements)
8. [We’re hiring!](http://dropbox.tech#-were-hiring)

9. [我们遗留的基于 Nginx 的流量基础设施](http://dropbox.tech#-our-legacy-nginx-based-traffic-infrastructure)
10. [为什么不是Bandaid？](http://dropbox.tech#-why-not-bandaid)
11. [我们新的基于 Envoy 的交通基础设施](http://dropbox.tech#-our-new-envoy-based-traffic-infrastructure)
12. [我们迁移的现状](http://dropbox.tech#-current-state-of-our-migration)
13. [我们遇到的问题](http://dropbox.tech#-issues-we-encountered-)
14. [下一步是什么？](http://dropbox.tech#-whats-next)
15. [致谢](http://dropbox.tech#acknowledgements)
16. [我们正在招聘！](http://dropbox.tech#-were-hiring)

In this blogpost we’ll talk about the old Nginx-based traffic infrastructure, its pain points, and the benefits we gained by migrating to [Envoy](https://www.envoyproxy.io/). We’ll compare Nginx to Envoy across many software engineering and operational dimensions. We’ll also briefly touch on the migration process, its current state, and some of the problems encountered on the way.

在这篇博文中，我们将讨论旧的基于 Nginx 的流量基础设施、它的痛点以及我们通过迁移到 [Envoy](https://www.envoyproxy.io/) 获得的好处。我们将在许多软件工程和运营维度上将 Nginx 与 Envoy 进行比较。我们还将简要介绍迁移过程、其当前状态以及在此过程中遇到的一些问题。

When we moved most of Dropbox traffic to Envoy, we had to seamlessly migrate a system that already handles tens of millions of open connections, millions of requests per second, and terabits of bandwidth. This effectively made us into one of the biggest Envoy users in the world.

当我们将大部分 Dropbox 流量转移到 Envoy 时，我们必须无缝迁移一个已经处理数千万个开放连接、每秒数百万个请求和 TB 级带宽的系统。这有效地使我们成为世界上最大的 Envoy 用户之一。

Disclaimer: although we’ve tried to remain objective, quite a few of these comparisons are specific to Dropbox and the way our software development works: making bets on Bazel, gRPC, and C++/Golang.

免责声明：尽管我们努力保持客观，但其中不少比较是针对 Dropbox 和我们的软件开发工作方式的：在 Bazel、gRPC 和 C++/Golang 上下注。

Also note that we’ll cover the open source version of the Nginx, not its commercial version with additional features.

另请注意，我们将介绍 Nginx 的开源版本，而不是具有附加功能的商业版本。

## Our legacy Nginx-based traffic infrastructure

## 我们传统的基于 Nginx 的流量基础设施

Our Nginx configuration was mostly static and rendered with a combination of Python2, Jinja2, and YAML. Any change to it required a full re-deployment. All dynamic parts, such as upstream management and a stats exporter, were written in Lua. Any sufficiently complex logic was moved to [the next proxy layer, written in Go](https://dropbox.tech/infrastructure/meet-bandaid-the-dropbox-service-proxy).

我们的 Nginx 配置大部分是静态的，并使用 Python2、Jinja2 和 YAML 的组合呈现。对它的任何更改都需要完全重新部署。所有动态部分，例如上游管理和统计数据导出器，都是用 Lua 编写的。任何足够复杂的逻辑都被移到 [下一个代理层，用 Go 编写](https://dropbox.tech/infrastructure/meet-bandaid-the-dropbox-service-proxy)。

Our post, “ [Dropbox](https://dropbox.tech/infrastructure/dropbox-traffic-infrastructure-edge-network) [traffic infrastructure: Edge network](https://dropbox.tech/infrastructure/dropbox-traffic-infrastructure-edge-network),” has a section about our legacy Nginx-based infrastructure.

我们的帖子，“ [Dropbox](https://dropbox.tech/infrastructure/dropbox-traffic-infrastructure-edge-network)[交通基础设施：边缘网络](https://dropbox.tech/infrastructure/dropbox-traffic-基础设施边缘网络)，”有一个关于我们传统的基于 Nginx 的基础设施的部分。

Nginx served us well for almost a decade. But it didn’t adapt to our current development best-practices:

Nginx 为我们服务了将近十年。但它并没有适应我们当前的开发最佳实践：

- Our internal and (private) external APIs are gradually migrating from REST to gRPC which requires all sorts of transcoding features from proxies.
- Protocol buffers became defacto  standard for service definitions and configurations.
- All software, regardless of the language, is built and tested with Bazel.
- Heavy involvement of our engineers on essential infrastructure projects in the open source community.

- 我们的内部和（私有）外部 API 正在逐渐从 REST 迁移到 gRPC，这需要来自代理的各种转码功能。
- 协议缓冲区成为服务定义和配置的事实上的标准。
- 所有软件，无论使用何种语言，均使用 Bazel 构建和测试。
- 我们的工程师大量参与开源社区的重要基础设施项目。

Also, operationally Nginx was quite expensive to maintain:

此外，Nginx 的运营维护成本非常高：

- Config generation logic was too flexible and split between YAML, Jinja2, and Python.
- Monitoring was a mix of Lua, log parsing, and system-based monitoring.
- An increased reliance on third party modules affected stability, performance, and the cost of subsequent upgrades.
- Nginx deployment and process management was quite different from the rest of the services. It relied a lot on other systems’ configurations: syslog, logrotate, etc, as opposed to being fully separate from the base system.

- 配置生成逻辑过于灵活，在 YAML、Jinja2 和 Python 之间分裂。
- 监控混合了 Lua、日志解析和基于系统的监控。
- 对第三方模块的依赖增加影响了稳定性、性能和后续升级的成本。
- Nginx 部署和进程管理与其他服务完全不同。它在很大程度上依赖于其他系统的配置：syslog、logrotate 等，而不是与基本系统完全分离。

With all of that, for the first time in 10 years, we started looking for a potential replacement for Nginx.

有了这一切，10 年来第一次，我们开始寻找 Nginx 的潜在替代品。

## Why not Bandaid? 

## 为什么不是 Bandaid？

As we frequently mention, internally we rely heavily on the Golang-based proxy called [Bandaid](https://dropbox.tech/infrastructure/meet-bandaid-the-dropbox-service-proxy). It has a great integration with Dropbox infrastructure, because it has access to the vast ecosystem of internal Golang libraries: monitoring, service discoveries, rate limiting, etc. We considered migrating from Nginx to Bandaid but there are a couple of issues that prevent us from doing that:

正如我们经常提到的，在内部我们严重依赖名为 [Bandaid](https://dropbox.tech/infrastructure/meet-bandaid-the-dropbox-service-proxy) 的基于 Golang 的代理。它与 Dropbox 基础设施有很好的集成，因为它可以访问内部 Golang 库的庞大生态系统：监控、服务发现、速率限制等。我们考虑从 Nginx 迁移到 Bandaid，但有几个问题阻止我们这样做：

- Golang is more resource intensive than C/C++. Low resource usage is especially important for us on the Edge since we can’t easily “auto-scale” our deployments there.
   - CPU overhead mostly comes from GC, HTTP parser and TLS, with the latter being less optimized than BoringSSL used by Nginx/Envoy.
   - The “goroutine-per-request” model and GC overhead greatly increase memory requirements in high-connection services like ours.
- No FIPS support for Go’s TLS stack.
- Bandaid does not have a community outside of Dropbox, which means that we can only rely on ourself for feature development.

- Golang 比 C/C++ 更占用资源。低资源使用率对我们在 Edge 上尤其重要，因为我们无法轻松地在那里“自动扩展”我们的部署。
  - CPU 开销主要来自 GC、HTTP 解析器和 TLS，后者的优化不如 Nginx/Envoy 使用的 BoringSSL。
  - “goroutine-per-request”模型和 GC 开销大大增加了像我们这样的高连接服务的内存需求。
- Go 的 TLS 堆栈没有 FIPS 支持。
- Bandaid 在 Dropbox 之外没有社区，这意味着我们只能依靠自己进行功能开发。

With all that we’ve decided to start migrating our traffic infrastructure to Envoy instead.

尽管如此，我们决定开始将我们的流量基础设施迁移到 Envoy。

## Our new Envoy-based traffic infrastructure

## 我们新的基于 Envoy 的交通基础设施

Let’s look into the main development and operational dimensions one by one, to see why we think Envoy is a better choice for us and what we gained by moving from Nginx to Envoy.

让我们从主要的开发和运营维度一一来看，看看为什么我们认为 Envoy 是我们更好的选择，以及我们从 Nginx 迁移到 Envoy 的收获。

### Performance

###  表现

Nginx’s [architecture](https://www.nginx.com/blog/inside-nginx-how-we-designed-for-performance-scale/) is event-driven and multi-process. It has support for [SO\_REUSEPORT](https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/), [EPOLLEXCLUSIVE](https://lore.kernel.org/patchwork/cover/543309/), and worker-to-CPU pinning. Although it is event-loop based, is it not fully non-blocking. This means some operations, like [opening a file](https://blog.cloudflare.com/how-we-scaled-nginx-and-saved-the-world-54-years-every-day/) or access/ error logging, can potentially cause an event-loop stall ( [even](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [with](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [aio](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [, aio\_write, and thread pools enabled](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio).) This leads to increased tail latencies, which can cause multi-second delays on spinning disk drives.

Nginx 的 [架构](https://www.nginx.com/blog/inside-nginx-how-we-designed-for-performance-scale/) 是事件驱动和多进程的。它支持 [SO\_REUSEPORT](https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/), [EPOLLEXCLUSIVE](https://lore.kernel.org/patchwork/cover/543309/)，以及工人到 CPU 的固定。虽然它是基于事件循环的，但它不是完全非阻塞的。这意味着一些操作，例如 [打开文件](https://blog.cloudflare.com/how-we-scaled-nginx-and-saved-the-world-54-years-every-day/) 或访问/错误记录，可能会导致事件循环停止（ [even](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [with](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [aio](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [，aio\_write 和线程池已启用](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio）。)这会导致尾部延迟增加，这可能导致旋转磁盘驱动器出现数秒延迟。

Envoy has a similar event-driven architecture, except it uses threads instead of processes. It also has SO\_REUSEPORT support (with a BPF filter support) and relies on libevent for event loop implementation (in other words, no fancy epoll(2) features like EPOLLEXCLUSIVE.) Envoy does not have any blocking IO operations in the event loop . Even logging is implemented in a non-blocking way, so that it does not cause stalls.

Envoy 有一个类似的事件驱动架构，只是它使用线程而不是进程。它还具有 SO\_REUSEPORT 支持（具有 BPF 过滤器支持）并依赖 libevent 进行事件循环实现（换句话说，没有像 EPOLLEXCLUSIVE 这样的花哨的 epoll(2) 功能。）Envoy 在事件循环中没有任何阻塞 IO 操作.甚至日志记录也是以非阻塞方式实现的，因此它不会导致停顿。

It looks like in theory Nginx and Envoy should have similar performance characteristics. But hope is not our strategy, so our first step was to run a diverse set of workload tests against similarly tuned Nginx and Envoy setups.

理论上 Nginx 和 Envoy 应该具有相似的性能特征。但希望不是我们的策略，所以我们的第一步是针对类似调整的 Nginx 和 Envoy 设置运行一组不同的工作负载测试。

If you are interested in performance tuning, we describe our standard tuning guidelines in “ [Optimizing](https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency) [web servers for high throughput and low latency](https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency).” It involves everything from picking the hardware, to OS tunables, to library choices and web server configuration.

如果您对性能调优感兴趣，我们在“ [优化](https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency)[web高吞吐量和低延迟的服务器](https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency)。”它涉及从选择硬件到操作系统可调项，再到库选择和 Web 服务器配置的所有内容。

Our test results showed similar performance between Nginx and Envoy under most of our test workloads: high requests per second (RPS), high bandwidth, and a mixed low-latency/high-bandwidth gRPC proxying. 

我们的测试结果显示，在大多数测试工作负载下，Nginx 和 Envoy 之间的性能相似：高每秒请求数 (RPS)、高带宽和混合低延迟/高带宽 gRPC 代理。

It is arguably very hard to make a good performance test. Nginx has [guidelines for performance testing](https://www.nginx.com/blog/testing-the-performance-of-nginx-and-nginx-plus-web-servers/), but these are not codified. Envoy also has [a guideline for benchmarking](https://www.envoyproxy.io/docs/envoy/latest/faq/performance/how_to_benchmark_envoy), and even some tooling under the [envoy](https://github.com/envoyproxy/envoy-perf) [-perf](https://github.com/envoyproxy/envoy-perf) project, but sadly the latter looks unmaintained.

可以说很难进行良好的性能测试。 Nginx 有[性能测试指南](https://www.nginx.com/blog/testing-the-performance-of-nginx-and-nginx-plus-web-servers/)，但这些都没有编纂。 Envoy 也有 [a guideline for benchmarking](https://www.envoyproxy.io/docs/envoy/latest/faq/performance/how_to_benchmark_envoy)，甚至 [envoy] 下的一些工具(https://github.com /envoyproxy/envoy-perf) [-perf](https://github.com/envoyproxy/envoy-perf) 项目，但遗憾的是后者看起来没有维护。

We resorted to using our internal testing tool. It’s called “hulk” because of its reputation for smashing our services.

我们求助于使用我们的内部测试工具。它被称为“绿巨人”，因为它以破坏我们的服务而闻名。

That said, there were a couple of notable differences in results:

也就是说，结果存在一些显着差异：

- Nginx showed higher long tail latencies. This was mostly due to event loops stalls under heavy I/O, especially if used together with SO\_REUSEPORT since in that case [connections can be accepted on behalf of a currently blocked worker](https://blog.cloudflare.com/the-sad-state-of-linux-socket-balancing/#so_reuseporttotherescue).
- Nginx performance without stats collections is on par with Envoy, but our Lua stats collection slowed Nginx on the high-RPS test by a factor of 3. This was expected given our reliance onlua\_shared\_dict, which is synchronized across workers with a mutex.

- Nginx 显示出更高的长尾延迟。这主要是由于大量 I/O 下的事件循环停顿，尤其是与 SO\_REUSEPORT 一起使用时，因为在这种情况下[可以代表当前被阻止的工作人员接受连接](https://blog.cloudflare.com/the-sad-state-of-linux-socket-balancing/#so_reuseporttotherescue)。
- 没有统计数据集合的 Nginx 性能与 Envoy 相当，但我们的 Lua 统计数据集合在高 RPS 测试中将 Nginx 的速度降低了 3 倍。鉴于我们依赖于 lua\_shared\_dict，这是预期的，它在具有互斥体。

We do understand how inefficient our stats collection was. We considered implementing something akin to FreeBSD's [counter(9)](https://www.freebsd.org/cgi/man.cgi?query=counter&sektion=9&manpath=freebsd-release-ports#IMPLEMENTATION_DETAILS) in userspace: CPU pinning, per-worker lockless counters with a fetching routine that loops through all workers aggregating their individual stats. But we gave up on this idea, because if we wanted to instrument Nginx internals (e.g. all error conditions), it would mean supporting an enormous patch that would make subsequent upgrades a true hell.

我们确实了解我们的统计数据收集效率低下。我们考虑在用户空间中实现类似于 FreeBSD 的 [counter(9)](https://www.freebsd.org/cgi/man.cgi?query=counter&sektion=9&manpath=freebsd-release-ports#IMPLEMENTATION_DETAILS)：CPU pinning，每个工人的无锁计数器具有一个获取例程，循环遍历所有工人汇总他们的个人统计数据。但是我们放弃了这个想法，因为如果我们想检测 Nginx 内部（例如所有错误情况)，这意味着支持一个巨大的补丁，这将使后续升级成为真正的地狱。

Since Envoy does not suffer from either of these issues, after migrating to it we were able to release up to 60% of servers previously exclusively occupied by Nginx.

由于 Envoy 没有遇到这些问题中的任何一个，因此在迁移到它之后，我们能够释放多达 60% 以前由 Nginx 独占的服务器。

### Observability

### 可观察性

Observability is the [most fundamental operational need](https://landing.google.com/sre/sre-book/chapters/part3/#fig_part-practices_reliability-hierarchy) for any product, but especially for such a foundational piece of infrastructure as a proxy. It is even more important during the migration period, so that any issue can be detected by the monitoring system rather than reported by frustrated users.

可观察性是任何产品的[最基本的操作需求](https://landing.google.com/sre/sre-book/chapters/part3/#fig_part-practices_reliability-hierarchy)，尤其是对于这样一个基础设施作为代理。在迁移期间更重要的是，任何问题都可以被监控系统检测到，而不是被沮丧的用户报告。

Non-commercial Nginx comes with a “ [stub](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html) [status](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html)” module that has 7 stats:

非商业 Nginx 带有一个“ [stub](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html)[status](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html)。 html)”模块有 7 个统计信息：

Copy


```
Active connections: 291
server accepts handled requests
16630948 16630948 31070465
Reading: 6 Writing: 179 Waiting: 106
```


This was definitely not enough, so we've added a simple log\_by\_lua handler that adds per-request stats based on headers and variables that are available in Lua: status codes, sizes, cache hits, etc. Here is an example of a simple stats-emitting function:

这绝对不够，所以我们添加了一个简单的 log\_by\_lua 处理程序，它根据 Lua 中可用的标头和变量添加每个请求的统计信息：状态代码、大小、缓存命中等。这是一个例子一个简单的统计信息发射功能：





```
function _M.cache_hit_stats(stat)
    if _var.upstream_cache_status then
        if _var.upstream_cache_status == "HIT" then
            stat:add("upstream_cache_hit")
        else
            stat:add("upstream_cache_miss")
        end
    end
end

```


In addition to the per-request Lua stats, we also had a very brittle error.log parser that was responsible for upstream, http, Lua, and TLS error classification.

除了每个请求的 Lua 统计信息之外，我们还有一个非常脆弱的 error.log 解析器，负责上游、http、Lua 和 TLS 错误分类。

On top of all that, we had a separate exporter for gathering Nginx internal state: time since the last reload, number of workers, RSS/VMS sizes, TLS certificate ages, etc.

最重要的是，我们有一个单独的导出器来收集 Nginx 内部状态：自上次重新加载以来的时间、工人数量、RSS/VMS 大小、TLS 证书年龄等。

A typical Envoy setup provides us thousands of distinct metrics (in [prometheus format](https://prometheus.io/docs/instrumenting/exposition_formats/#text-based-format)) describing both proxied traffic and server’s internal state:

典型的 Envoy 设置为我们提供了数千个不同的指标（以 [prometheus 格式](https://prometheus.io/docs/instrumenting/exposition_formats/#text-based-format))描述代理流量和服务器的内部状态：





```
$ curl -s http://localhost:3990/stats/prometheus |wc -l
14819

```


This includes a myriad of stats with different aggregations:

这包括无数具有不同聚合的统计数据：

- Per-cluster/per-upstream/per-vhost HTTP stats, including connection pool info and various timing histograms.
- Per-listener TCP/HTTP/TLS downstream connection stats.
- Various internal/runtime stats from basic version info and uptime to memory allocator stats and deprecated feature usage counters. 

- 每个集群/每个上游/每个虚拟主机 HTTP 统计信息，包括连接池信息和各种时间直方图。
- 每个侦听器的 TCP/HTTP/TLS 下游连接统计信息。
- 各种内部/运行时统计信息，从基本版本信息和正常运行时间到内存分配器统计信息和不推荐使用的功能使用计数器。

A special shoutout is needed for Envoy’s [admin interface](https://www.envoyproxy.io/docs/envoy/latest/operations/admin). Not only does it provide additional structured stats through /certs, /clusters, and /config\_dump endpoints, but there are also very important operational features:

Envoy 的 [管理界面](https://www.envoyproxy.io/docs/envoy/latest/operations/admin) 需要一个特殊的喊叫。它不仅通过 /certs、/clusters 和 /config\_dump 端点提供额外的结构化统计信息，而且还有非常重要的操作特性：

- The ability to change error logging on the fly through[/logging](https://www.envoyproxy.io/docs/envoy/latest/operations/admin#post--logging). This allowed us to troubleshoot fairly obscure problems in a matter of minutes.
- /cpuprofiler, /heapprofiler, /contention which would surely be quite useful during the inevitable performance troubleshooting.
- /runtime\_modify  endpoint allows us to change set of configuration parameters without pushing new configuration, which could be used in feature gating, etc.

- 能够通过[/logging](https://www.envoyproxy.io/docs/envoy/latest/operations/admin#post--logging)即时更改错误日志记录。这使我们能够在几分钟内解决相当模糊的问题。
- /cpuprofiler, /heapprofiler, /contention 在不可避免的性能故障排除过程中肯定会非常有用。
- /runtime\_modify 端点允许我们在不推送新配置的情况下更改配置参数集，可用于功能门控等。

In addition to stats, Envoy also supports [pluggable tracing providers](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing). This is useful not only to our Traffic team, who own multiple load-balancing tiers, but also for application developers who want to track request latencies end-to-end from the edge to app servers.

除了统计数据，Envoy 还支持 [pluggable tracking providers](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing)。这不仅对我们拥有多个负载平衡层的流量团队有用，而且对于想要跟踪从边缘到应用服务器的端到端请求延迟的应用程序开发人员也很有用。

Technically, Nginx also supports tracing through a third-party [OpenTracing integration](https://github.com/opentracing-contrib/nginx-opentracing) [,](https://github.com/opentracing-contrib/nginx-opentracing) but it is not under heavy development.

从技术上讲，Nginx 还支持通过第三方进行跟踪 [OpenTracing 集成](https://github.com/opentracing-contrib/nginx-opentracing) [,](https://github.com/opentracing-contrib/nginx-opentracing)但它并没有进行大量开发。

And last but not least, Envoy has the ability to [stream access logs over gRPC](https://www.envoyproxy.io/docs/envoy/latest/api-v3/data/accesslog/v3/accesslog.proto). This removes the burden of supporting syslog-to-hive bridges from our Traffic team. Besides, it’s way easier (and secure!) to spin up a generic gRPC service in Dropbox production than to add a custom TCP/UDP listener.

最后但并非最不重要的是，Envoy 能够[通过 gRPC 流式传输访问日志](https://www.envoyproxy.io/docs/envoy/latest/api-v3/data/accesslog/v3/accesslog.proto)。这消除了我们的流量团队支持 syslog-to-hive 桥接的负担。此外，在 Dropbox 产品中启动通用 gRPC 服务比添加自定义 TCP/UDP 侦听器更容易（且安全！)。

Configuration of access logging in Envoy, like everything else, happens through a gRPC management service, the [Access Log Service](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto) (ALS). Management services are the standard way of integrating the Envoy data plane with various services in production. This brings us to our next topic.

Envoy 中访问日志的配置，与其他任何事情一样，通过 gRPC 管理服务发生，[访问日志服务](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto) (ALS)。管理服务是将 Envoy 数据平面与生产中的各种服务集成的标准方式。这将我们带到下一个主题。

### Integration

###  一体化

Nginx’s approach to integration is best described as “Unix-ish.” Configuration is very static. It heavily relies on files (eg the config file itself, TLS certificates and tickets, allowlists/blocklists, etc.) and well-known industry protocols ( [logging](http://nginx.org/en/docs/syslog.html) [to syslog](http://nginx.org/en/docs/syslog.html) and auth sub-requests through HTTP). Such simplicity and backwards compatibility is a good thing for small setups, since Nginx can be easily automated with a couple of shell scripts. But as the system’s scale increases, testability and standardization become more important.

Nginx 的集成方法最好被描述为“Unix-ish”。配置非常静态。它严重依赖文件（例如配置文件本身、TLS 证书和票证、许可名单/阻止名单等）和众所周知的行业协议（[日志记录](http://nginx.org/en/docs/syslog.html) [到系统日志](http://nginx.org/en/docs/syslog.html) 和通过 HTTP 认证的子请求)。这种简单性和向后兼容性对于小型设置来说是一件好事，因为 Nginx 可以通过几个 shell 脚本轻松实现自动化。但随着系统规模的增加，可测试性和标准化变得更加重要。

Envoy is far more opinionated in how the traffic dataplane should be integrated with its control plane, and hence with the rest of infrastructure. It encourages the use of [protobufs](https://developers.google.com/protocol-buffers) and gRPC by providing a stable API commonly referred as [xDS](https://docs.google.com/document/d/1xeVvJ6KjFBkNjVspPbY_PwEDHC7XPi0J5p1SqUXcCl8/edit#heading=h.c0uts5ftkk58). Envoy discovers its dynamic resources by [querying one or more of these xDS services](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration#arch-overview-dynamic-config).

Envoy 对流量数据平面应该如何与其控制平面进行集成，从而与其他基础设施进行集成的看法要强得多。它通过提供通常称为 [xDS](https://docs.google.com/document/d/1xeVvJ6KjFBkNjVspPbY_PwEDHC7XPi0J5p1SqUXcCl8/edit#heading=h.c0uts5ftkk58)。 Envoy 通过[查询一个或多个这些 xDS 服务](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration#arch-overview-dynamic-config) 来发现其动态资源。

Nowadays, the xDS APIs are evolving beyond Envoy: [Universal](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) [D](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) [ata Plane API](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a)[(UDPA)](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) has the ambitious goal of “becoming de facto standard of L4/L7 loadbalancers.” 

如今，xDS API 正在超越 Envoy：[Universal](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) [D](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) [ata Plane API](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a)[(UDPA)](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a)blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) 的雄心勃勃的目标是“成为 L4/L7 负载均衡器的事实上的标准”。

From our experience, this ambition works out well. We already use [Open Request Cost Aggregation](https://docs.google.com/document/d/1NSnK3346BkBo1JUU3I9I5NYYnaJZQPt8_Z_XCBCI3uA/edit)[(ORCA)](https://docs.google.com/document/d/1NSnK3346BkBo1JUU3I9I5NYYnaJZQPt8_Z_XCBCI3uA/edit) for our internal load testing, and are considering using UDPA for our non-Envoy loadbalancers eg our [Katran-based eBPF/XDP Layer-4 Load Balancer](https://github.com/facebookincubator/katran).

根据我们的经验，这种雄心很有效。我们已经在使用 [Open Request Cost Aggregation](https://docs.google.com/document/d/1NSnK3346BkBo1JUU3I9I5NYYnaJZQPt8_Z_XCBCI3uA/edit)[(ORCA)](https://docs.google.com/document/d/1NSnK3345UCY3Pt8K33JUU3I9I5NYYnaJZQPt8_Z_XCBCI3uA/edit)编辑)用于我们的内部负载测试，并且正在考虑将 UDPA 用于我们的非 Envoy 负载均衡器，例如我们的 [基于 Katran 的 eBPF/XDP 第 4 层负载均衡器](https://github.com/facebookincubator/katran)。

This is especially good for Dropbox, where all services internally already interact through gRPC-based APIs. We’ve implemented our own version of xDS control plane that integrates Envoy with our configuration management, service discovery, secret management, and route information.

这对 Dropbox 尤其有用，所有内部服务都已经通过基于 gRPC 的 API 进行交互。我们已经实现了我们自己的 xDS 控制平面版本，它将 Envoy 与我们的配置管理、服务发现、秘密管理和路由信息集成在一起。

For more information about Dropbox RPC, please read “ [Courier:](https://dropbox.tech/infrastructure/courier-dropbox-migration-to-grpc) [Dropbox migration to gRPC](https://dropbox.tech/infrastructure/courier-dropbox-migration-to-grpc).” There we describe in detail how we integrated service discovery, secret management, stats, tracing, circuit breaking, etc, with gRPC.

有关 Dropbox RPC 的更多信息，请阅读“[Courier:](https://dropbox.tech/infrastructure/courier-dropbox-migration-to-grpc) [Dropbox 迁移到 gRPC](https://dropbox.tech/基础设施/快递投递箱迁移到grpc)。”我们在那里详细描述了我们如何将服务发现、秘密管理、统计、跟踪、熔断等与 gRPC 集成。

Here are **some** of the available xDS services, their Nginx alternatives, and our examples of how we use them:

以下是**一些**可用的 xDS 服务、它们的 Nginx 替代品，以及我们如何使用它们的示例：

- [Access Log Service](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto)[(ALS)](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto), as mentioned above, lets us dynamically configure access log destinations, encodings, and formats. Imagine a dynamic version of Nginx’s log\_format and access\_log.
- [Endpoint discovery service](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds)[(EDS)](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds) provides information about cluster members. This is analogous to a dynamically updated list of upstream block's server entries (eg for Lua that would be a  [balancer\_by\_lua\_block](https://github.com/openresty/lua-nginx-module#balancer_by_lua_block)) in the Nginx config. In our case we proxied this to our internal service discovery.
- [Secret discovery service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret)[(SDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret) provides various TLS-related information that would cover various ssl\_\* directives (and respectively [ssl\_\*\_by\_lua\_block](https://github.com/openresty/lua-nginx-module#ssl_certificate_by_lua_block).)  We adapted this interface to our secret distribution service.
- [Runtime Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/operations/runtime#config-runtime-rtds)[(RTDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/operations/runtime#config-runtime-rtds) is providing runtime flags. Our implementation of this functionality in Nginx was quite hacky, based on checking the existence of various files from Lua. This approach can quickly become inconsistent between the individual servers. Envoy’s default implementation is also filesystem-based, but we instead pointed our RTDS xDS API to our distributed configuration storage. That way we can control whole clusters at once (through a tool with a sysctl-like interface) and there are no accidental inconsistencies between different servers. 

- [访问日志服务](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto)[(ALS)](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto)，如上所述，让我们动态配置访问日志目的地、编码和格式。想象一个动态版本的 Nginx 的 log\_format 和 access\_log。
- [端点发现服务](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds)[(EDS)](https ://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds）提供有关集群成员的信息。这类似于动态更新的上游块的服务器条目列表（例如，对于 Lua，这将是一个 [balancer\_by\_lua\_block](https://github.com/openresty/lua-nginx-module#balancer_by_lua_block))在 Nginx 配置中。在我们的例子中，我们将其代理到我们的内部服务发现中。
- [秘密发现服务](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret)[(SDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret) 提供了各种与 TLS 相关的信息，这些信息将涵盖各种 ssl\_\* 指令（分别是 [ssl\_\*\_by\_lua\_block](https://github.com/openresty/lua-nginx-module#ssl_certificate_by_lua_block)。我们将此接口改编为我们的秘密分发服务。
- [运行时发现服务](https://www.envoyproxy.io/docs/envoy/latest/configuration/operations/runtime#config-runtime-rtds)[(RTDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/operations/runtime#config-runtime-rtds) 提供运行时标志。我们在 Nginx 中实现这个功能是非常困难的，基于检查 Lua 中各种文件的存在。这种方法很快就会在各个服务器之间变得不一致。 Envoy 的默认实现也是基于文件系统的，但我们将 RTDS xDS API 指向我们的分布式配置存储。这样我们就可以一次控制整个集群（通过一个具有类似 sysctl 界面的工具)，并且不同服务器之间不会出现意外的不一致。

- [Route discovery service](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/rds#config-http-conn-man-rds)[(RDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/rds#config-http-conn-man-rds) maps routes to virtual hosts, and allows additional configuration for headers and filters. In Nginx terms, these would be analogous to a dynamic location block with set\_header/proxy\_set\_header and a proxy\_pass. On lower proxy tiers we autogenerate these directly from our service definition configs.

- [路由发现服务](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/rds#config-http-conn-man-rds)[(RDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/rds#config-http-conn-man-rds) 将路由映射到虚拟主机，并允许对标头和过滤器进行额外配置。在 Nginx 术语中，这些类似于带有 set\_header/proxy\_set\_header 和 proxy\_pass 的动态位置块。在较低的代理层上，我们直接从我们的服务定义配置中自动生成这些。

For an example of Envoy's integration with an existing production system, here is a canonical example of how to [integrate](https://www.envoyproxy.io/learn/service-discovery) [Envoy](https://www.envoyproxy.io/learn/service-discovery) [with a custom service discovery](https://www.envoyproxy.io/learn/service-discovery). There are also a couple of open source Envoy control-plane implementations, such as [Istio](https://istio.io/) and the less complex [go-control-plane](https://github.com/envoyproxy/go-control-plane).

有关 Envoy 与现有生产系统集成的示例，这里是如何[集成](https://www.envoyproxy.io/learn/service-discovery)[Envoy](https://www.envoy) 的规范示例。 envoyproxy.io/learn/service-discovery) [带有自定义服务发现](https://www.envoyproxy.io/learn/service-discovery)。还有一些开源的 Envoy 控制平面实现，例如 [Istio](https://istio.io/) 和不太复杂的 [go-control-plane](https://github.com/envoyproxy)/go-control-plane)。

Our homegrown Envoy control plane implements an increasing number of xDS APIs. It is deployed as a normal gRPC service in production, and acts as an adapter for our infrastructure building blocks. It does this through a set of common Golang libraries to talk to internal services and expose them through a stable xDS APIs to Envoy. The whole process does not involve any filesystem calls, signals, cron, logrotate, syslog, log parsers, etc.

我们自主开发的 Envoy 控制平面实现了越来越多的 xDS API。它在生产中部署为普通的 gRPC 服务，并充当我们基础架构构建块的适配器。它通过一组通用的 Golang 库与内部服务对话并通过稳定的 xDS API 将它们公开给 Envoy。整个过程不涉及任何文件系统调用、信号、cron、logrotate、syslog、日志解析器等。

### Configuration

###  配置

Nginx has the undeniable advantage of a simple human-readable configuration. But this win gets lost as config gets more complex and begins to be code-generated.

Nginx 具有不可否认的优势，即简单的人类可读配置。但是随着配置变得更加复杂并开始由代码生成，这个胜利就失去了。

As mentioned above, our Nginx config is generated through a mix of Python2, Jinja2, and YAML. Some of you may have seen or even written a variation of this in erb, pug, Text::Template, or maybe even m4:

如上所述，我们的 Nginx 配置是通过混合使用 Python2、Jinja2 和 YAML 生成的。你们中的一些人可能已经在 erb、pug、Text::Template 甚至 m4 中看到或什至编写了它的变体：


```
{% for server in servers %}
server {
    {% for error_page in server.error_pages %}
    error_page {{ error_page.statuses|join(' ') }} {{ error_page.file }};
    {% endfor %}
    ...
    {% for route in service.routes %}
    {% if route.regex or route.prefix or route.exact_path %}
    location {% if route.regex %}~ {{route.regex}}{%
            elif route.exact_path %}= {{ route.exact_path }}{%
            else %}{{ route.prefix }}{% endif %} {
        {% if route.brotli_level %}
        brotli on;
        brotli_comp_level {{ route.brotli_level }};
        {% endif %}
        ...

```


Our approach to Nginx config generation had a huge issue: all of the languages involved in config generation allowed substitution and/or logic. YAML has anchors, Jinja2 has loops/ifs/macroses, and of course Python is Turing-complete. Without a clean data model, complexity quickly spread across all three of them.

我们的 Nginx 配置生成方法有一个大问题：配置生成中涉及的所有语言都允许替换和/或逻辑。 YAML 有锚点，Jinja2 有循环/ifs/macroses，当然 Python 是图灵完备的。如果没有一个干净的数据模型，复杂性就会迅速蔓延到这三个模型中。

This problem is arguably fixable, but there were a couple of foundational ones:

这个问题可以说是可以解决的，但有几个基本问题：

- There is no declarative description for the config format. If we wanted to programmatically generate and validate configuration, we would need to invent it ourselves.
- Config that is syntactically valid could still be invalid from a C code standpoint. For example, some of the buffer-related variables have limitations on values, restrictions on alignment, and interdependencies with other variables. To semantically validate a config we needed to run it throughnginx -t.

- 配置格式没有声明性描述。如果我们想以编程方式生成和验证配置，我们需要自己发明它。
- 从 C 代码的角度来看，语法上有效的配置可能仍然无效。例如，一些与缓冲区相关的变量在值上有限制、对对齐有限制，以及与其他变量的相互依赖性。为了从语义上验证配置，我们需要通过 nginx -t 运行它。

Envoy, on the other hand, has a unified data-model for configs: all of its configuration is defined in Protocol Buffers. This not only solves the data modeling problem, but also adds typing information to the config values. Given that protobufs are first class citizens in Dropbox production, and a common way of describing/configuring services, this makes integration _so_ much easier.

另一方面，Envoy 有一个统一的配置数据模型：它的所有配置都在协议缓冲区中定义。这不仅解决了数据建模问题，而且还为配置值添加了类型信息。鉴于 protobuf 是 Dropbox 产品中的一等公民，并且是描述/配置服务的常用方式，这使得集成_so_更容易。

Our new config generator for Envoy is based on protobufs and Python3. All data modeling is done in proto files, while all the logic is in Python. Here’s an example:

我们新的 Envoy 配置生成器基于 protobufs 和 Python3。所有的数据建模都是在 proto 文件中完成的，而所有的逻辑都是在 Python 中完成的。下面是一个例子：





```
from dropbox.proto.envoy.extensions.filters.http.gzip.v3.gzip_pb2 import Gzip
from dropbox.proto.envoy.extensions.filters.http.compressor.v3.compressor_pb2 import Compressor

def default_gzip_config(
    compression_level: Gzip.CompressionLevel.Enum = Gzip.CompressionLevel.DEFAULT,
    ) -> Gzip:
        return Gzip(
            # Envoy's default is 6 (Z_DEFAULT_COMPRESSION).
            compression_level=compression_level,
            # Envoy's default is 4k (12 bits).Nginx uses 32k (MAX_WBITS, 15 bits).
            window_bits=UInt32Value(value=12),
            # Envoy's default is 5. Nginx uses 8 (MAX_MEM_LEVEL - 1).
            memory_level=UInt32Value(value=5),
            compressor=Compressor(
                content_length=UInt32Value(value=1024),
                remove_accept_encoding_header=True,
                content_type=default_compressible_mime_types(),
            ),
        )
```




- Note the[Python3 type annotations](https://www.python.org/dev/peps/pep-0484/) in that code! Coupled with [mypy-protobuf protoc plugin](https://github.com/dropbox/mypy-protobuf), these provide end-to-end typing inside the config generator. IDEs capable of checking them will immediately highlight typing mismatches.

- 请注意该代码中的[Python3 类型注释](https://www.python.org/dev/peps/pep-0484/)！结合 [mypy-protobuf protoc 插件](https://github.com/dropbox/mypy-protobuf)，这些在配置生成器中提供端到端的输入。能够检查它们的 IDE 将立即突出显示打字不匹配。

There are still cases where a type-checked protobuf can be logically invalid. In the example above, gzip window\_bits can only take values between 9 and 15. This kind of restriction can be easily defined with a help of [protoc-gen-validate protoc plugin](https://github.com/envoyproxy/protoc-gen-validate):

仍然存在类型检查的 protobuf 可能在逻辑上无效的情况。在上面的例子中，gzip window\_bits 只能取 9 到 15 之间的值。这种限制可以借助 [protoc-gen-validate protoc plugin](https://github.com/envoyproxy/protoc-gen-validate)：





```
google.protobuf.UInt32Value window_bits = 9 [(validate.rules).uint32 = {lte: 15 gte: 9}];

```


Finally, an implicit benefit of using a formally defined configuration model is that it organically leads to the documentation being collocated with the configuration definitions. [Here](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) ['](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) [s an example from](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) [gzip.proto](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50):

最后，使用正式定义的配置模型的一个隐含好处是它有机地导致文档与配置定义并置。 [这里](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) ['](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) [示例来自](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) [gzip.proto](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50)：





```
// Value from 1 to 9 that controls the amount of internal memory used by zlib.Higher values.
// use more memory, but are faster and produce better compression results.The default value is 5.
google.protobuf.UInt32Value memory_level = 1 [(validate.rules).uint32 = {lte: 9 gte: 1}];
```


For those of you thinking about using protobufs in your production systems, but worried you may lack a schema-less representation, here's a good article from Envoy core developer Harvey Tuch about how to work around this using google.protobuf.Struct and google.protobuf .Any: “ [Dynamic](https://blog.envoyproxy.io/dynamic-extensibility-and-protocol-buffers-dcd0bf0b8801) [extensibility and Protocol Buffers](https://blog.envoyproxy.io/dynamic-extensibility-and-protocol-buffers-dcd0bf0b8801).”

对于那些考虑在生产系统中使用 protobufs，但又担心可能缺乏无模式表示的人，这里是 Envoy 核心开发人员 Harvey Tuch 的一篇好文章，关于如何使用 google.protobuf.Struct 和 google.protobuf 解决这个问题.Any：“ [动态](https://blog.envoyproxy.io/dynamic-extensibility-and-protocol-buffers-dcd0bf0b8801) [可扩展性和协议缓冲区](https://blog.envoyproxy.io/dynamic-extensibility-and-protocol-buffers-dcd0bf0b8801)。”

### Extensibility

### 可扩展性

Extending Nginx beyond what’s possible with standard configuration usually requires writing a C module. Nginx’s [development guide](http://nginx.org/en/docs/dev/development_guide.html) provides a solid introduction to the available building blocks. That said, this approach is relatively heavyweight. In practice, it takes a fairly senior software engineer to safely write an Nginx module.

将 Nginx 扩展到超出标准配置可能的范围通常需要编写一个 C 模块。 Nginx 的 [开发指南](http://nginx.org/en/docs/dev/development_guide.html) 提供了可用构建块的可靠介绍。也就是说，这种方法相对重量级。在实践中，需要相当资深的软件工程师才能安全地编写 Nginx 模块。

In terms of infrastructure available for module developers, they can expect basic containers like hash tables/queues/rb-trees, (non-RAII) memory management, and hooks for all phases of request processing. There are also couple of external libraries like pcre, zlib, openssl, and, of course, libc.

就模块开发人员可用的基础设施而言，他们可以期待基本的容器，如哈希表/队列/rb 树、（非 RAII）内存管理和用于请求处理所有阶段的钩子。还有一些外部库，如 pcre、zlib、openssl，当然还有 libc。

For more lightweight feature extension, Nginx provides [Perl](http://nginx.org/en/docs/http/ngx_http_perl_module.html#perl) and [Javascript](http://nginx.org/en/docs/http/ngx_http_js_module.html) interfaces. Sadly, both are fairly limited in their abilities, mostly restricted to the content phase of request processing.

对于更轻量级的功能扩展，Nginx 提供了 [Perl](http://nginx.org/en/docs/http/ngx_http_perl_module.html#perl) 和 [Javascript](http://nginx.org/en/docs/http/ngx_http_js_module.html) 接口。遗憾的是，两者的能力都相当有限，主要限于请求处理的内容阶段。

The most commonly used extension method adopted by the community is based on a third-party l [ua-](https://github.com/openresty/lua-nginx-module) [nginx](https://github.com/openresty/lua-nginx-module) [-module](https://github.com/openresty/lua-nginx-module) and various [OpenResty libraries](https://github.com/openresty/). This approach can be hooked in at pretty much any phase of request processing. We used log\_by\_lua for stats collection, and balancer\_by\_lua for dynamic backend reconfiguration.

社区采用的最常用的扩展方式是基于第三方l [ua-](https://github.com/openresty/lua-nginx-module) [nginx](https://github.com/openresty/lua-nginx-module) [-module](https://github.com/openresty/lua-nginx-module) 和各种 [OpenResty 库](https://github.com/openresty/)。这种方法几乎可以用于请求处理的任何阶段。我们使用 log\_by\_lua 进行统计收集，使用 balancer\_by\_lua 进行动态后端重新配置。

In theory, Nginx provides the ability to develop [modules in C++](http://lxr.nginx.org/source/src/misc/ngx_cpp_test_module.cpp). In practice, it lacks proper C++ interfaces/wrappers for all the primitives to make this worthwhile. There are nonetheless some [community attempts at it](https://github.com/chronolaw/ngx_cpp_dev). These are far from ready for production, though. 

理论上，Nginx 提供了开发[C++ 中的模块](http://lxr.nginx.org/source/src/misc/ngx_cpp_test_module.cpp) 的能力。在实践中，它缺乏所有原语的适当 C++ 接口/包装器来使其值得。尽管如此，还是有一些 [社区尝试](https://github.com/chronolaw/ngx_cpp_dev)。不过，这些还远未准备好投入生产。

Envoy’s main extension mechanism is through C++ plugins. The process is [not as well documented](https://blog.envoyproxy.io/how-to-write-envoy-filters-like-a-ninja-part-1-d166e5abec09) as in Nginx's case, but it is simpler. This is partially due to:

Envoy 的主要扩展机制是通过 C++ 插件。该过程[没有详细记录](https://blog.envoyproxy.io/how-to-write-envoy-filters-like-a-ninja-part-1-d166e5abec09) 与 Nginx 的情况一样，但它是更简单。这部分是由于：

- **Clean and well-commented interfaces.** C++ classes act as natural extension and documentation points. For example, [checkout the HTTP filter interface](https://github.com/envoyproxy/envoy/blob/master/include/envoy/http/filter.h).
- **C++14 language and standard library.** From basic language features like templates and lambda functions, to type-safe containers and algorithms. In general, writing modern C++14 is not much different from using Golang or, with a stretch, one may even say Python.
- **Features beyond C++14 and its stdlib.** Provided by the [abseil](https://abseil.io/about/intro) library, these include drop-in replacements from newer C++ standards, mutexes with built -in [static deadlock detection](http://clang.llvm.org/docs/ThreadSafetyAnalysis.html) and debug support, additional/more efficient containers, [and much more](https://abseil.io/about/philosophy).

- **干净且注释良好的接口。** C++ 类充当自然扩展和文档点。例如，[检查HTTP过滤器接口](https://github.com/envoyproxy/envoy/blob/master/include/envoy/http/filter.h)。
- **C++14 语言和标准库。** 从模板和 lambda 函数等基本语言功能到类型安全的容器和算法。一般来说，编写现代 C++14 与使用 Golang 并没有太大区别，甚至可以说是 Python。
- **C++14 及其 stdlib 以外的功能。** 由 [abseil](https://abseil.io/about/intro) 库提供，其中包括来自较新 C++ 标准的直接替换、内置互斥锁-in [静态死锁检测](http://clang.llvm.org/docs/ThreadSafetyAnalysis.html)和调试支持，附加/更高效的容器，[以及更多](https://abseil.io/about/哲学)。

For specifics, here’s a [canonical example of an HTTP Filter module](https://github.com/envoyproxy/envoy-filter-example/tree/master/http-filter-example).

具体来说，这里有一个 [HTTP 过滤器模块的规范示例](https://github.com/envoyproxy/envoy-filter-example/tree/master/http-filter-example)。

We were able to integrate Envoy with [Vortex2](https://dropbox.tech/infrastructure/monitoring-server-applications-with-vortex) [(our](https://dropbox.tech/infrastructure/monitoring-server-applications-with-vortex) [monitoring framework)](https://dropbox.tech/infrastructure/monitoring-server-applications-with-vortex) with only 200 lines of code by simply implementing the Envoy [stats](https://github.com/envoyproxy/envoy/tree/master/include/envoy/stats) interface.

我们能够将 Envoy 与 [Vortex2](https://dropbox.tech/infrastructure/monitoring-server-applications-with-vortex) [(our](https://dropbox.tech/infrastructure/monitoring-server-Applications-with-vortex) [监控框架)](https://dropbox.tech/infrastructure/monitoring-server-applications-with-vortex) 只需 200 行代码，只需实现 Envoy [stats](https://github.com/envoyproxy/envoy/tree/master/include/envoy/stats) 接口。

Envoy [also has Lua support](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/lua_filter) through [moonjit](https://github.com/moonjit/moonjit), a LuaJIT fork with improved Lua 5.2 support. Compared to Nginx’s 3rd-party Lua integration it has far fewer capabilities and hooks. This makes Lua in Envoy far less attractive due to the cost of additional complexity in developing, testing, and troubleshooting interpreted code. Companies that specialize in Lua development may disagree, but in our case we decided to avoid it and use C++ exclusively for Envoy extensibility.

Envoy [也有 Lua 支持](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/lua_filter) 通过 [moonjit](https://github.com/moonjit/moonjit)，具有改进的 Lua 5.2 支持的 LuaJIT 分支。与 Nginx 的 3rd-party Lua 集成相比，它的功能和钩子要少得多。这使得 Envoy 中的 Lua 的吸引力大大降低，因为开发、测试和对解释代码进行故障排除会带来额外的复杂性成本。专门从事 Lua 开发的公司可能不同意，但在我们的案例中，我们决定避免使用它，而将 C++ 专门用于 Envoy 可扩展性。

What distinguishes Envoy from the rest of web servers is its emerging support for [WebAssembly](https://developer.mozilla.org/en-US/docs/WebAssembly/Concepts)(WASM) — a fast, portable, and secure extension mechanism. WASM is not meant to be used directly, but as a compilation target for any general-purpose programming language. Envoy implements a [WebAssembly for Proxies specification](https://github.com/proxy-wasm/spec/blob/master/abi-versions/vNEXT/README.md) (and also includes reference [Rust](https://github.com/proxy-wasm/proxy-wasm-rust-sdk) and [C++](https://github.com/proxy-wasm/proxy-wasm-cpp-sdk)SDKs) that describes the boundary between WASM code and a generic L4/L7 proxy. That separation between the proxy and extension code allows for secure sandboxing, while WASM low-level compact binary format allows for near native efficiency. On top of that, in Envoy proxy-wasm extensions are integrated with xDS. This allows dynamic updates and even potential A/B testing.

Envoy 与其他 Web 服务器的区别在于它对 [WebAssembly](https://developer.mozilla.org/en-US/docs/WebAssembly/Concepts)(WASM) 的新兴支持——一种快速、便携且安全的扩展机制。 WASM 并不打算直接使用，而是作为任何通用编程语言的编译目标。 Envoy 实现了 [WebAssembly for Proxies 规范](https://github.com/proxy-wasm/spec/blob/master/abi-versions/vNEXT/README.md)（还包括参考[Rust](https://github.com/proxy-wasm/spec/blob/master/abi-versions/vNEXT/README.md) /github.com/proxy-wasm/proxy-wasm-rust-sdk) 和 [C++](https://github.com/proxy-wasm/proxy-wasm-cpp-sdk)SDKs) 描述了 WASM 之间的边界代码和通用 L4/L7 代理。代理和扩展代码之间的分离允许安全沙箱，而 WASM 低级紧凑二进制格式允许接近本机效率。最重要的是，在 Envoy 中，proxy-wasm 扩展与 xDS 集成。这允许动态更新甚至潜在的 A/B 测试。

The “ [Extending](https://youtu.be/XdWmm_mtVXI?t=779)[Envoy](https://youtu.be/XdWmm_mtVXI?t=779) [with WebAssembly](https://youtu.be/XdWmm_mtVXI?t=779)” presentation from Kubecon'19 (remember that time when we had non-virtual conferences?) has a nice overview of  WASM in Envoy and its potential uses. It also hints at performance levels of 60-70% of native C++ code. 

“ [扩展](https://youtu.be/XdWmm_mtVXI?t=779)[Envoy](https://youtu.be/XdWmm_mtVXI?t=779) [with WebAssembly](https://youtu.be/XdWmm_mtVXI?t=779)”来自 Kubecon'19 的演讲（还记得我们举行非虚拟会议的那段时间吗？)很好地概述了 Envoy 中的 WASM 及其潜在用途。它还暗示了 60-70% 的本机 C++ 代码的性能水平。

With WASM, service providers get a safe and efficient way of running customers’ code on their edge. Customers get the benefit of portability: Their extensions can run on any cloud that implements the proxy-wasm ABI. Additionally, it allows your users to use any language as long as it can be compiled to WebAssembly. This enables them to use a broader set of non-C++ libraries, securely and efficiently.

使用 WASM，服务提供商可以获得一种安全有效的方式在他们的边缘运行客户的代码。客户获得可移植性的好处：他们的扩展可以在任何实现 proxy-wasm ABI 的云上运行。此外，它允许您的用户使用任何可以编译为 WebAssembly 的语言。这使他们能够安全有效地使用更广泛的非 C++ 库集。

Istio is putting a lot of resources into WebAssembly development: they already have an experimental version of the WASM-based telemetry extension and the [WebAssemblyHub community](https://webassemblyhub.io/) for sharing extensions. You can read about it in detail in [“Redefining](https://istio.io/latest/blog/2020/wasm-announce/) [extensibility in proxies - introducing WebAssembly to](https://istio.io/latest/blog/2020/wasm-announce/) [Envoy](https://istio.io/latest/blog/2020/wasm-announce/) [and Istio](https://istio.io/latest/blog/2020/wasm-announce/) [.](https://istio.io/latest/blog/2020/wasm-announce/) [”](https://istio.io/latest/blog/2020/wasm-announce/)

Istio 在 WebAssembly 开发中投入了大量资源：他们已经有了基于 WASM 的遥测扩展的实验版本和用于共享扩展的 [WebAssemblyHub 社区](https://webassemblyhub.io/)。您可以在[“重新定义](https://istio.io/latest/blog/2020/wasm-announce/) [代理中的可扩展性 - 介绍 WebAssembly 到](https://istio.io/latest/blog/2020/wasm-announce/) [Envoy](https://istio.io/latest/blog/2020/wasm-announce/) [和 Istio](https://istio.io/latest/blog/2020/wasm-announce/) [.](https://istio.io/latest/blog/2020/wasm-announce/) [”](https://istio.io/latest/blog/2020/wasm-宣布/)

Currently, we don’t use WebAssembly at Dropbox. But this might change when the Go SDK for proxy-wasm is available.

目前，我们在 Dropbox 不使用 WebAssembly。但是当适用于 proxy-wasm 的 Go SDK 可用时，这可能会改变。

### Building and Testing

### 构建和测试

By default, Nginx is built using a custom [shell-based configuration system](https://github.com/nginx/nginx/tree/master/auto) and make-based build system. This is simple and elegant, but it took quite a bit of effort to integrate it into [B](https://dropbox.tech/infrastructure/continuous-integration-and-deployment-with-bazel) [azel-built monorepo] (https://dropbox.tech/infrastructure/continuous-integration-and-deployment-with-bazel) to get all the benefits of incremental, distributed, hermetic, and reproducible builds.

默认情况下，Nginx 是使用自定义的 [基于 shell 的配置系统](https://github.com/nginx/nginx/tree/master/auto) 和基于 make 的构建系统构建的。这个简单又优雅，但是把它集成到[B](https://dropbox.tech/infrastructure/continuous-integration-and-deployment-with-bazel) [azel-built monorepo] (https://dropbox.tech/infrastructure/continuous-integration-and-deployment-with-bazel) 以获得增量、分布式、密封和可重复构建的所有好处。

Google open [-](https://nginx.googlesource.com/nginx/) sourced their [B](https://nginx.googlesource.com/nginx/) [azel-built](https://nginx.googlesource.com/nginx/) [Nginx](https://nginx.googlesource.com/nginx/)[version](https://nginx.googlesource.com/nginx/) which consists of Nginx, BoringSSL, PCRE, ZLIB, and Brotli library/module.

谷歌开放 [-](https://nginx.googlesource.com/nginx/) 来源他们的 [B](https://nginx.googlesource.com/nginx/) [azel-built](https://nginx.googlesource.com/nginx/) [Nginx](https://nginx.googlesource.com/nginx/)[版本](https://nginx.googlesource.com/nginx/) 由 Nginx、BoringSSL、PCRE、 ZLIB 和 Brotli 库/模块。

Testing-wise, Nginx has a set of Perl-driven [integration tests](http://hg.nginx.org/nginx-tests) in a separate repository and no unit tests.

在测试方面，Nginx 在单独的存储库中有一组 Perl 驱动的 [集成测试](http://hg.nginx.org/nginx-tests)，没有单元测试。

Given our heavy usage of Lua and absence of a built-in unit testing framework, we resorted to testing using mock configs and a simple Python-based test driver:

鉴于我们大量使用 Lua 并且没有内置的单元测试框架，我们求助于使用模拟配置和简单的基于 Python 的测试驱动程序进行测试：





```
class ProtocolCountersTest(NginxTestCase):
    @classmethod
    def setUpClass(cls):
        super(ProtocolCountersTest, cls).setUpClass()
        cls.nginx_a = cls.add_nginx(
            nginx_CONFIG_PATH, endpoint=["in"], upstream=["out"],
        )
        cls.start_nginxes()

    @assert_delta(lambda d: d == 0, get_stat("request_protocol_http2"))
    @assert_delta(lambda d: d == 1, get_stat("request_protocol_http1"))
    def test_http(self):
        r = requests.get(self.nginx_a.endpoint["in"].url("/"))
        assert r.status_code == requests.codes.ok

```


On top of that, we verify the syntax-correctness of all generated configs by preprocessing them (e.g. replacing all IP addresses with 127/8 ones, switching to self-signed TLS certs, etc.) and running nginx -c on the result.

最重要的是，我们通过预处理所有生成的配置来验证它们的语法正确性（例如，将所有 IP 地址替换为 127/8 的 IP 地址，切换到自签名 TLS 证书等）并对结果运行 nginx -c。

On the Envoy side, the main build system is already Bazel. So integrating it with our monorepo was trivial: Bazel easily allows [adding external dependencies](https://docs.bazel.build/versions/master/external.html).

在 Envoy 方面，主要构建系统已经是 Bazel。所以将它与我们的 monorepo 集成是微不足道的：Bazel 很容易允许[添加外部依赖项](https://docs.bazel.build/versions/master/external.html)。

We also use [copybara](https://github.com/google/copybara) scripts to sync protobufs for both Envoy and udpa. Copybara is handy when you need to do simple transformations without the need to forever maintain a large patchset.

我们还使用 [copybara](https://github.com/google/copybara) 脚本来同步 Envoy 和 udpa 的 protobuf。当您需要进行简单的转换而无需永远维护大型补丁集时，Copybara 非常方便。

With Envoy we have the flexibility of using either unit tests (based on gtest/gmock) with a set of [pre-written mocks](https://github.com/envoyproxy/envoy/tree/master/test/mocks), or Envoy's [integration test framework](https://github.com/envoyproxy/envoy/tree/master/test/integration), or both. There’s no need anymore to rely on slow end-to-end integration tests for every trivial change. 

使用 Envoy，我们可以灵活地使用单元测试（基于 gtest/gmock）和一组 [预先编写的模拟](https://github.com/envoyproxy/envoy/tree/master/test/mocks)，或 Envoy 的 [集成测试框架](https://github.com/envoyproxy/envoy/tree/master/test/integration)，或两者兼而有之。不再需要为每个微不足道的更改依赖缓慢的端到端集成测试。

[gtest](https://github.com/google/googletest) is a fairly well-known unit-test framework used by Chromium and LLVM, among others. If you want to know more about googletest there are good intros for both [googletest](https://github.com/google/googletest/blob/master/googletest/docs/primer.md) and [googlemock](https://chromium.googlesource.com/external/github.com/google/googletest/+/HEAD/googlemock/docs/cook_book.md).

[gtest](https://github.com/google/googletest) 是 Chromium 和 LLVM 等使用的相当知名的单元测试框架。如果您想了解有关 googletest 的更多信息，[googletest](https://github.com/google/googletest/blob/master/googletest/docs/primer.md) 和 [googlemock](https://chromium.googlesource.com/external/github.com/google/googletest/+/HEAD/googlemock/docs/cook_book.md)。

Open source Envoy development [requires changes to have 100% unit test coverage](https://github.com/envoyproxy/envoy/blob/master/CONTRIBUTING.md#submitting-a-pr). Tests are automatically triggered for each pull request via the [Azure CI Pipeline](https://dev.azure.com/cncf/envoy/_build?view=pipelines).

开源 Envoy 开发 [需要更改以获得 100% 的单元测试覆盖率](https://github.com/envoyproxy/envoy/blob/master/CONTRIBUTING.md#submitting-a-pr)。通过 [Azure CI Pipeline](https://dev.azure.com/cncf/envoy/_build?view=pipelines) 为每个拉取请求自动触发测试。

It’s also a common practice to micro-benchmark performance-sensitive code with [google/becnhmark](https://github.com/google/benchmark):

使用 [google/becnhmark](https://github.com/google/benchmark) 对性能敏感代码进行微基准测试也是一种常见做法：





```
$ bazel run --compilation_mode=opt test/common/upstream:load_balancer_benchmark -- --benchmark_filter=".*LeastRequestLoadBalancerChooseHost.*"
BM_LeastRequestLoadBalancerChooseHost/100/1/1000000          848 ms          449 ms            2 mean_hits=10k relative_stddev_hits=0.0102051 stddev_hits=102.051
...

```


After switching to Envoy, we began to rely exclusively on unit tests for our internal module development:

切换到 Envoy 后，我们开始完全依赖单元测试进行内部模块开发：





```
TEST_F(CourierClientIdFilterTest, IdentityParsing) {
struct TestCase {
    std::vector<std::string> uris;
    Identity expected;
};
std::vector<TestCase> tests = {
    {{"spiffe://prod.dropbox.com/service/foo"}, {"spiffe://prod.dropbox.com/service/foo", "foo"}},
    {{"spiffe://prod.dropbox.com/user/boo"}, {"spiffe://prod.dropbox.com/user/boo", "user.boo"}},
    {{"spiffe://prod.dropbox.com/host/strange"}, {"spiffe://prod.dropbox.com/host/strange", "host.strange"}},
    {{"spiffe://corp.dropbox.com/user/bad-prefix"}, {"", ""}},
};
for (auto& test : tests) {
    EXPECT_CALL(*ssl_, uriSanPeerCertificate()).WillOnce(testing::Return(test.uris));
    EXPECT_EQ(GetIdentity(ssl_), test.expected);
}
}

```


Having sub-second test roundtrips has a compounding effect on productivity. It empowers us to put more effort into increasing test coverage. And being able to choose between unit and integration tests allows us to balance coverage, speed, and cost of Envoy tests.

亚秒级的测试往返对生产力有复合影响。它使我们能够投入更多精力来增加测试覆盖率。能够在单元测试和集成测试之间进行选择使我们能够平衡 Envoy 测试的覆盖率、速度和成本。

Bazel is one of the best things that ever happened to our developer experience. It has a very steep learning curve and is a large upfront investment, but it has a very high return on that investment: [incremental builds](https://docs.bazel.build/versions/master/guide.html#correct-incremental-rebuilds), [remote caching](https://docs.bazel.build/versions/master/remote-caching.html), [distributed builds/tests](https://docs.bazel.build/versions/master/remote-execution.html), etc.

Bazel 是我们开发人员体验中发生过的最好的事情之一。它有一个非常陡峭的学习曲线，并且是一笔巨大的前期投资，但它的投资回报非常高：[增量构建](https://docs.bazel.build/versions/master/guide.html#correct-增量重建)、[远程缓存](https://docs.bazel.build/versions/master/remote-caching.html)、[分布式构建/测试](https://docs.bazel.build/versions/master/remote-execution.html)等。

One of the less discussed benefits of Bazel is that it gives us an ability to [query](https://docs.bazel.build/versions/master/query-how-to.html) [and even augment](https://docs.bazel.build/versions/master/skylark/aspects.html) the dependency graph. A programmatic interface to the dependency graph, coupled with a common build system across all languages, is a very powerful feature. It can be used as a foundational building block for things like linters, code generation, vulnerability tracking, deployment system, etc.

Bazel 较少讨论的好处之一是它使我们能够[查询](https://docs.bazel.build/versions/master/query-how-to.html) [甚至增加](https://docs.bazel.build/versions/master/skylark/aspects.html)依赖关系图。依赖图的编程接口，加上跨所有语言的通用构建系统，是一个非常强大的功能。它可以用作 linter、代码生成、漏洞跟踪、部署系统等的基础构建块。

### Security

###  安全

Nginx’s code surface is quite small, with minimal external dependencies. It’s typical to see only 3 external dependencies on the resulting binary: zlib (or [one of its faster variants](https://github.com/cloudflare/zlib)), a TLS library, and PCRE. Nginx has a custom implementation of all protocol parsers, the event library, and they even went as far as to re-implement some libc functions. 

Nginx 的代码面很小，外部依赖最小。通常在生成的二进制文件上只看到 3 个外部依赖项：zlib（或 [其更快的变体之一](https://github.com/cloudflare/zlib))、TLS 库和 PCRE。 Nginx 有所有协议解析器的自定义实现，事件库，他们甚至重新实现了一些 libc 函数。

At some point Nginx was considered so secure that it was used as a default web server in OpenBSD. Later two development communities had a falling out, which lead to the creation of  httpd. You can read about the motivation behind that move in BSDCon's “ [Introducing](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) [OpenBSD](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) ['s](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) [new httpd](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf).”

在某些时候，Nginx 被认为非常安全，以至于它被用作 OpenBSD 中的默认 Web 服务器。后来两个开发社区发生了争执，导致了 httpd 的创建。您可以在 BSDCon 的“[介绍](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) [OpenBSD](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) ['s](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) [新 httpd](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf)。”

This minimalism paid off in practice. Nginx has only had 30 [vulnerabilities and exposures](https://nginx.org/en/security_advisories.html) reported in more than 11 years.

这种极简主义在实践中得到了回报。在 11 年多的时间里，Nginx 仅报告了 30 个[漏洞和暴露](https://nginx.org/en/security_advisories.html)。

Envoy, on the other hand, has way more code, especially when you consider that that C++ code is far more dense than the basic C used for Nginx. It also incorporates millions of lines of code from external dependencies. Everything from event notification to protocol parsers is offloaded to 3rd party libraries. This increases attack surface and bloats the resulting binary.

另一方面，Envoy 有更多的代码，尤其是当您考虑到 C++ 代码比用于 Nginx 的基本 C 代码密集得多时。它还包含来自外部依赖项的数百万行代码。从事件通知到协议解析器的所有内容都卸载到 3rd 方库。这会增加攻击面并使生成的二进制文件膨胀。

To counteract this, Envoy relies heavily on modern security practices. It uses [AddressSanitizer](https://github.com/google/sanitizers/wiki/AddressSanitizer),[ThreadSanitizer](https://github.com/google/sanitizers/wiki/ThreadSanitizerCppManual), and [MemorySanitizer](https://github.com/google/sanitizers/wiki/MemorySanitizer). Its developers even went beyond that and adopted [fuzzing](https://bugs.chromium.org/p/oss-fuzz/issues/list?q=label%3AProj-envoy&sort=-id).

为了解决这个问题，Envoy 严重依赖现代安全实践。它使用 [AddressSanitizer](https://github.com/google/sanitizers/wiki/AddressSanitizer)、[ThreadSanitizer](https://github.com/google/sanitizers/wiki/ThreadSanitizerCppManual) 和 [MemorySanitizer](https://github.com/google/sanitizers/wiki/MemorySanitizer)。它的开发人员甚至超越了这一点，并采用了 [fuzzing](https://bugs.chromium.org/p/oss-fuzz/issues/list?q=label%3AProj-envoy&sort=-id)。

Any opensource project that is critical to the global IT infrastructure can be accepted to the [OSS-Fuzz](https://github.com/google/oss-fuzz)—a free platform for automated fuzzing. To learn more about it, see “ [OSS-Fuzz](https://google.github.io/oss-fuzz/architecture/) [/ Architecture](https://google.github.io/oss-fuzz/architecture/).”

任何对全球 IT 基础设施至关重要的开源项目都可以被 [OSS-Fuzz](https://github.com/google/oss-fuzz)——一个自动模糊测试的免费平台接受。要了解更多信息，请参阅“[OSS-Fuzz](https://google.github.io/oss-fuzz/architecture/)[/架构](https://google.github.io/oss-fuzz/建筑学/)。”

In practice, though, all these precautions do not fully counteract the increased code footprint. As a result, Envoy has had [22 security advisories in the](https://github.com/envoyproxy/envoy/security/advisories) [p](https://github.com/envoyproxy/envoy/security/advisories) [ast 2 years](https://github.com/envoyproxy/envoy/security/advisories).

但在实践中，所有这些预防措施并不能完全抵消增加的代码占用空间。结果，Envoy 已经有了 [22 个安全公告](https://github.com/envoyproxy/envoy/security/advisories) [p](https://github.com/envoyproxy/envoy/security/advisories) [过去 2 年](https://github.com/envoyproxy/envoy/security/advisories)。

Envoy's [security release policy is described in great detail](https://github.com/envoyproxy/envoy/security/policy), and in [postmortems](https://github.com/envoyproxy/envoy/tree/master/security/postmortems) for selected vulnerabilities. Envoy is also a participant in [Google's Vulnerability Reward Program](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/security/google_vrp#arch-overview-google-vrp)[(VRP)] (https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/security/google_vrp#arch-overview-google-vrp). Open to all security researchers, VRP provides rewards for vulnerabilities discovered and reported according to their rules.

Envoy 的 [安全发布政策有很详细的描述](https://github.com/envoyproxy/envoy/security/policy)，在[postmortems](https://github.com/envoyproxy/envoy/tree/master) /security/postmortems）针对选定的漏洞。 Envoy 也是 [Google 的漏洞奖励计划](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/security/google_vrp#arch-overview-google-vrp)[(VRP)] 的参与者（https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/security/google_vrp#arch-overview-google-vrp)。 VRP 向所有安全研究人员开放，为根据其规则发现和报告的漏洞提供奖励。

For a practical example of how some of these vulnerabilities can be potentially exploited, see this writeup about CVE-2019–18801: “ [Exploiting](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792) [an](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792) [Envoy](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792) [heap vulnerability](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792).”

有关如何潜在地利用其中一些漏洞的实际示例，请参阅有关 CVE-2019-18801 的这篇文章：“[利用](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792) [an](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792)[Envoy](https://blog.envoyproxy.io/exploiting-an-envoy-heap-漏洞-96173d41792) [堆漏洞](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792)。”

To counteract the increased vulnerability risk, we use best binary hardening security practices from our upstream OS vendors [Ubuntu](https://wiki.ubuntu.com/Security/Features) and [Debian](https://wiki.debian.org/Hardening). We defined a special hardened build profile for all edge-exposed binaries. It includes ASLR, stack protectors, and symbol table hardening:

为了抵消增加的漏洞风险，我们使用来自上游操作系统供应商 [Ubuntu](https://wiki.ubuntu.com/Security/Features) 和 [Debian](https://wiki.debian.组织/强化)。我们为所有边缘暴露的二进制文件定义了一个特殊的强化构建配置文件。它包括 ASLR、堆栈保护器和符号表强化：





```
build:hardened --force_pic
build:hardened --copt=-fstack-clash-protection
build:hardened --copt=-fstack-protector-strong
build:hardened --linkopt=-Wl,-z,relro,-z,now
```




Forking web-servers, like Nginx, in most environments [have issues with stack protector](http://hmarco.org/renewssp/data/Preventing_brute_force_attacks_against_stack_canary_protection_on_networking_servers-Paper.pdf). Since master and worker processes share the same stack canary, and on canary verification failure worker process is killed, the canary can be brute-forced bit-by-bit in about 1000 tries. Envoy, which uses threads as a concurrency primitive, is not affected by this attack.

在大多数环境中分叉网络服务器，如 Nginx，[堆栈保护器有问题](http://hmarco.org/renewssp/data/Preventing_brute_force_attacks_against_stack_canary_protection_on_networking_servers-Paper.pdf)。由于主进程和工作进程共享同一个堆栈金丝雀，并且在金丝雀验证失败时工作进程被杀死，金丝雀可以在大约 1000 次尝试中被逐位蛮力。使用线程作为并发原语的 Envoy 不受此攻击的影响。

We also want to harden third-party dependencies where we can. We use [BoringSSL in FIPS mode](https://boringssl.googlesource.com/boringssl/+/master/crypto/fipsmodule/FIPS.md), which includes startup self-tests and integrity checking of the binary. We’re also considering running ASAN-enabled binaries on some of our edge canary servers.

我们还希望尽可能强化第三方依赖项。我们使用 [FIPS 模式下的 BoringSSL](https://boringssl.googlesource.com/boringssl/+/master/crypto/fipsmodule/FIPS.md)，其中包括启动自检和二进制完整性检查。我们还考虑在我们的一些边缘金丝雀服务器上运行支持 ASAN 的二进制文件。

### Features

###  特征

Here comes the most opinionated part of the post, brace yourself.

这是帖子中最自以为是的部分，请振作起来。

Nginx began as a web server specialized on serving static files with minimal resource consumption. Its functionality is top of the line there: static serving, caching (including thundering herd protection), and range caching.

Nginx 最初是一个专门以最少的资源消耗提供静态文件的 Web 服务器。它的功能是最重要的：静态服务、缓存（包括雷鸣牛群保护）和范围缓存。

On the proxying side, though, Nginx lacks features needed for modern infrastructures. There’s no HTTP/2 to backends. gRPC proxying is available but without connection multiplexing. There’s no support for gRPC transcoding. On top of that, Nginx’s “open-core” model restricts features that can go into an open source version of the proxy. As a result, some of the critical features like statistics are not available in the “community” version.

然而，在代理方面，Nginx 缺乏现代基础设施所需的功能。后端没有 HTTP/2。 gRPC 代理可用，但没有连接多路复用。不支持 gRPC 转码。最重要的是，Nginx 的“开放核心”模型限制了可以进入代理的开源版本的功能。因此，某些关键功能（如统计）在“社区”版本中不可用。

Envoy, by contrast, has evolved as an ingress/egress proxy, used frequently for gRPC-heavy environments. Its web-serving functionality is rudimentary: [no file serving](https://github.com/envoyproxy/envoy/issues/378), still [work-in-progress caching](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/cache/v3alpha/cache.proto.html), neither [brotli](https://github.com/envoyproxy/envoy/issues/4429) nor pre-compression. For these use cases we still have a small fallback Nginx setup that Envoy uses as an upstream cluster.

相比之下，Envoy 已经演变为入口/出口代理，经常用于 gRPC 密集型环境。它的网络服务功能是基本的：[无文件服务](https://github.com/envoyproxy/envoy/issues/378)，仍然是[work-in-progress caching](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/cache/v3alpha/cache.proto.html)，既不是 [brotli](https://github.com/envoyproxy/envoy/issues/4429) 也不是预压缩。对于这些用例，我们仍然有一个小的后备 Nginx 设置，Envoy 将其用作上游集群。

When HTTP cache in Envoy becomes production-ready, we could move most of static-serving use cases to it, using S3 instead of filesystem for long-term storage. To read more about eCache design, see “ [eCache:](https://docs.google.com/document/d/1WPuim_GzhfdsnIj_tf-fIeutK0jO4aVQfVrLJFoLN3g/view#heading=h.wjxw6fq7wefi) [a multi-backend HTTP cache](https://docs.google.com/document/d/1WPuim_GzhfdsnIj_tf-fIeutK0jO4aVQfVrLJFoLN3g/view#heading=h.wjxw6fq7wefi) [for Envoy](https://docs.google.com/document/d/1WPuim_GzhfdsnIj_tf-fIeutK0jO4aVQfVrLJFoLN3g/view#heading=h.wjxw6fq7wefi).”

当 Envoy 中的 HTTP 缓存变为生产就绪时，我们可以将大多数静态服务用例转移到它，使用 S3 而不是文件系统进行长期存储。要阅读有关 eCache 设计的更多信息，请参阅“[eCache:](https://docs.google.com/document/d/1WPuim_GzhfdsnIj_tf-fIeutK0jO4aVQfVrLJFoLN3g/view#heading=h.wjxw6fq7wefi) [多 https 后端 HTTP 缓存]( ://docs.google.com/document/d/1WPuim_GzhfdsnIj_tf-fIeutK0jO4aVQfVrLJFoLN3g/view#heading=h.wjxw6fq7wefi) [为特使](https://docs.google.com/document/d/1WPuims0fjFoLN3g/view#heading=h.wjxw6fq7wefi)标题=h.wjxw6fq7wefi)。”

Envoy also has native support for many gRPC-related capabilities:

Envoy 还对许多与 gRPC 相关的功能提供了原生支持：

- **gRPC proxying.** This is a basic capability that allowed us to use gRPC end-to-end for our applications (e.g. Dropbox desktop client.)
- **HTTP/2 to backends.** This feature allows us to greatly reduce the number of TCP connections between our traffic tiers, reducing memory consumption and keepalive traffic.
- [gRPC → HTTP bridge](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_http1_bridge_filter) (\+ [reverse](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_http1_reverse_bridge_filter).)  These allowed us to expose legacy HTTP/1 applications using a modern gRPC stack.
- [gRPC-WEB](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_web_filter). This feature allowed us to use gRPC end-to-end even in the environments where middleboxes (firewalls, IDS, etc) don’t yet support HTTP/2.
- [gRPC JSON transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter). This enables us to transcode all inbound traffic, including [Dropbox public APIs](https://www.dropbox.com/developers/documentation/http/overview), from REST into gRPC.

- **gRPC 代理。** 这是一项基本功能，允许我们为我们的应用程序（例如 Dropbox 桌面客户端）使用端到端的 gRPC。
- **HTTP/2 到后端。** 此功能使我们能够大大减少我们的流量层之间的 TCP 连接数量，减少内存消耗和保活流量。
- [gRPC → HTTP 网桥](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_http1_bridge_filter) (\+ [反向](https://www.envoyproxy.io/docs)/envoy/latest/configuration/http/http_filters/grpc_http1_reverse_bridge_filter）。)这些允许我们使用现代 gRPC 堆栈公开旧的 HTTP/1 应用程序。
- [gRPC-WEB](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_web_filter)。即使在中间件（防火墙、IDS等)尚不支持 HTTP/2 的环境中，此功能也使我们能够端到端地使用 gRPC。
- [gRPC JSON 转码器](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)。这使我们能够将所有入站流量，包括 [Dropbox 公共 API](https://www.dropbox.com/developers/documentation/http/overview)，从 REST 转码为 gRPC。

In addition, Envoy can also be used as an outbound proxy. We used it to unify a couple of other use cases: 

此外，Envoy 还可以用作出站代理。我们用它来统一其他几个用例：

- Egress proxy: since Envoy [added support for the HTTP CONNECT method](https://github.com/envoyproxy/envoy/issues/1451), it can be used as a drop-in replacement for Squid proxies. We’ve begun to replace our outbound Squid installations with Envoy. This not only greatly improves visibility, but also reduces operational toil by unifying the stack with a common dataplane and observability (no more parsing logs for stats.)
- Third-party software service discovery: we are relying on the[Courier gRPC libraries](https://dropbox.tech/infrastructure/courier-dropbox-migration-to-grpc) in our software instead of using Envoy as a service mesh . But we do use Envoy in one-off cases where we need to connect an open source service with our service discovery with minimal effort. For example, Envoy is used as a service discovery sidecar in our analytics stack. Hadoop can dynamically discover its name and journal nodes. [Superset](https://github.com/apache/incubator-superset) can discover airflow, presto, and hive backends. [Grafana](https://grafana.com/) can discover its MySQL database.

- 出口代理：由于 Envoy [增加了对 HTTP CONNECT 方法的支持](https://github.com/envoyproxy/envoy/issues/1451)，它可以用作 Squid 代理的替代品。我们已经开始用 Envoy 替换我们的出站 Squid 安装。这不仅极大地提高了可见性，而且通过将堆栈与通用数据平面和可观察性（不再为统计数据解析日志)统一起来，减少了运营负担。
- 第三方软件服务发现：我们依赖软件中的 [Courier gRPC 库](https://dropbox.tech/infrastructure/courier-dropbox-migration-to-grpc) 而不是使用 Envoy 作为服务网格.但是我们确实在一次性的情况下使用 Envoy，在这种情况下，我们需要以最少的努力将开源服务与我们的服务发现连接起来。例如，Envoy 在我们的分析堆栈中用作服务发现边车。 Hadoop 可以动态发现其名称和日志节点。 [Superset](https://github.com/apache/incubator-superset) 可以发现气流、presto 和 hive 后端。 [Grafana](https://grafana.com/) 可以发现它的 MySQL 数据库。

## Community

##  社区

Nginx development is quite centralized. Most of its development happens behind closed doors. There's some external activity on the [nginx-devel](http://mailman.nginx.org/pipermail/nginx-devel/) mailing list, and there are occasional development-related discussions on the [official bug tracker](https://trac.nginx.org/nginx).

Nginx 开发非常集中。它的大部分发展都是闭门进行的。 [nginx-devel](http://mailman.nginx.org/pipermail/nginx-devel/) 邮件列表上有一些外部活动，在 [official bug tracker](https://trac.nginx.org/nginx)。

There is an #nginx channel on FreeNode. Feel free to join it for more interactive [community](https://www.nginx.com/resources/wiki/community/irc/) conversations.

FreeNode 上有一个#nginx 频道。随意加入它以进行更多互动 [社区](https://www.nginx.com/resources/wiki/community/irc/) 对话。

Envoy development is open and decentralized: coordinated through GitHub issues/pull requests, [mailing list](https://groups.google.com/g/envoy-dev), and [community meetings](https://goo.gl/5Cergb).

Envoy 开发是开放和去中心化的：通过 GitHub 问题/拉取请求、[邮件列表](https://groups.google.com/g/envoy-dev) 和 [社区会议](https://goo.gl)协调/5Cergb)。

There is also quite a bit of community activity on Slack. You can get your invite [here](https://envoyslack.cncf.io).

Slack 上也有很多社区活动。你可以在[这里](https://envoyslack.cncf.io)获得你的邀请。

It’s hard to quantify the development styles and engineering community, so let’s look at a specific example of HTTP/3 development.

很难量化开发风格和工程社区，所以让我们看一个 HTTP/3 开发的具体例子。

Nginx [QUIC and HTTP/3 implementation](https://hg.nginx.org/nginx-quic/) was [recently presented by F5](https://www.nginx.com/blog/introducing-technology-preview-nginx-support-for-quic-http-3/). The code is clean, with zero external dependencies. But the development process itself was rather opaque. Half a year before that, [Cloudflare came up with their own HTTP/3 implementation for](https://blog.cloudflare.com/experiment-with-http-3-using-nginx-and-quiche/) [Nginx] (https://blog.cloudflare.com/experiment-with-http-3-using-nginx-and-quiche/). As a result, the community now has two separate experimental versions of HTTP/3 for Nginx. 

Nginx [QUIC 和 HTTP/3 实现](https://hg.nginx.org/nginx-quic/) [最近由 F5 提出](https://www.nginx.com/blog/introducing-technology-preview-nginx-support-for-quic-http-3/)。代码很干净，外部依赖为零。但开发过程本身相当不透明。半年前，[Cloudflare 提出了自己的 HTTP/3 实现](https://blog.cloudflare.com/experiment-with-http-3-using-nginx-and-quiche/) [Nginx] （https://blog.cloudflare.com/experiment-with-http-3-using-nginx-and-quiche/)。因此，社区现在为 Nginx 提供了两个独立的 HTTP/3 实验版本。

In Envoy's case, HTTP/3 implementation is also a work in progress, based on chromium's " [quiche](https://docs.google.com/document/d/19qcrwAa8hVYZv2r8zZ7SgkylivAQNQ7E3loMJ-vk9_k/edit)" (QUIC, HTTP, Etc .) library. The project’s status is tracked in the [GitHub issue](https://github.com/envoyproxy/envoy/issues/2557). The [de](https://docs.google.com/document/d/1dEo19y-trABuW2x6-T564LmK7Ld-BPXZOlnR4df9KVU/edit#heading=h.w2fjl4fs3sex) [sign doc](https://docs.google.com/document/d/1dEo19y-trABuW2x6-T564LmK7Ld-BPXZOlnR4df9KVU/edit#heading=h.w2fjl4fs3sex) was publicly available way before patches were completed. Remaining work that would benefit from community involvement is tagged with “ [help](https://github.com/envoyproxy/envoy/issues?q=is%3Aopen+label%3Aarea%2Fquic+label%3A%22help+wanted%22) [wanted](https://github.com/envoyproxy/envoy/issues?q=is%3Aopen+label%3Aarea%2Fquic+label%3A%22help+wanted%22).”

在 Envoy 的案例中，HTTP/3 的实现也是一项正在进行的工作，基于 Chromium 的“[quiche](https://docs.google.com/document/d/19qcrwAa8hVYZv2r8zZ7SgkylivAQNQ7E3loMJ-vk9_k/edit)”（QUIC、HTTP 等。） 图书馆。项目状态在[GitHub issue](https://github.com/envoyproxy/envoy/issues/2557)中进行跟踪。[de](https://docs.google.com/document/d/1dEo19y-trABuW2x6-T564LmK7Ld-BPXZOlnR4df9KVU/edit#heading=h.w2fjl4fs3sex) [签署文档](https://docs.google.com/文档/d/1dEo19y-traBuW2x6-T564LmK7Ld-BPXZOLnR4df9KVU/edit#heading=h.w2fjl4fs3sex)在补丁完成之前公开可用。将受益于社区参与的剩余工作被标记为“[帮助](https://github.com/envoyproxy/envoy/issues?q=is%3Aopen+label%3Aarea%2Fquic+label%3A%22help+wanted%22) [通缉](https://github.com/envoyproxy/envoy/issues?q=is%3Aopen+label%3Aarea%2Fquic+label%3A%22help+wanted%22)。”

As you can see, the latter structure is much more transparent and greatly encourages collaboration. For us, this means that we managed to upstream lots of small to medium changes to Envoy–everything from [operational improvements](https://github.com/envoyproxy/envoy/pull/10286/files) and [performance optimizations]( https://github.com/envoyproxy/envoy/pull/9556) to [new gRPC transcoding features](https://github.com/envoyproxy/envoy/pull/10673) and [load](https://github.com/envoyproxy/envoy/pull/11006) [balancing changes](https://github.com/envoyproxy/envoy/pull/11006).

如您所见，后一种结构更加透明，并且极大地鼓励了协作。对我们来说，这意味着我们成功地对 Envoy 进行了大量中小型更改——一切都来自 [运营改进](https://github.com/envoyproxy/envoy/pull/10286/files) 和 [性能优化]( https://github.com/envoyproxy/envoy/pull/9556) 到 [新的 gRPC 转码功能](https://github.com/envoyproxy/envoy/pull/10673) 和 [load](https://github.com/envoyproxy/envoy/pull/11006) [平衡变化](https://github.com/envoyproxy/envoy/pull/11006)。

## Current state of our migration

## 我们迁移的当前状态

We’ve been running Nginx and Envoy side-by-side for over half a year and gradually switching traffic from one to another with DNS. By now we have migrated a wide variety of workloads to Envoy:

我们已经并行运行 Nginx 和 Envoy 半年多了，并逐渐通过 DNS 将流量从一个切换到另一个。到目前为止，我们已经将各种工作负载迁移到 Envoy：

- **Ingress high-throughput services.** All file data from Dropbox desktop client is served via end-to-end gRPC through Envoy. By switching to Envoy we’ve also slightly improved users’ performance, due to better connection reuse from the edge.
- **Ingress high-RPS services.** This is all file metadata for Dropbox desktop client. We get the same benefits of end-to-end gRPC, plus the removal of the connection pool, which means we are not bounded by one request per connection at a time.
- **Notification and telemetry services.** Here we handle all real-time notifications, so these servers have millions of HTTP connections (one for each active client.) Notification services can now be implemented via streaming gRPC instead of an expensive long- poll method.
- **Mixed high-throughput/high-RPS services.** API traffic (both metadata and data itself.) This allows us to start thinking about public gRPC APIs. We may even switch to transcoding our existing REST-based APIs right on the Edge.
- **Egress high-throughput proxies.** In our case, the Dropbox to AWS communication, mostly S3. This would allow us to eventually remove all Squid proxies from our production network, leaving us with a single L4/L7 data plane.

- **Ingress 高吞吐量服务。** 来自 Dropbox 桌面客户端的所有文件数据都通过 Envoy 通过端到端 gRPC 提供服务。通过切换到 Envoy，我们还略微提高了用户的性能，因为边缘的连接重用更好。
- **入口高 RPS 服务。** 这是 Dropbox 桌面客户端的所有文件元数据。我们获得了端到端 gRPC 的相同好处，加上连接池的删除，这意味着我们不受每个连接一次一个请求的限制。
- **通知和遥测服务。** 在这里我们处理所有实时通知，因此这些服务器有数百万个 HTTP 连接（每个活动客户端一个。）现在可以通过流式 gRPC 来实现通知服务，而不是昂贵的长时间投票方式。
- **混合高吞吐量/高 RPS 服务。** API 流量（元数据和数据本身）。这使我们可以开始考虑公共 gRPC API。我们甚至可以直接在 Edge 上转码我们现有的基于 REST 的 API。
- **出口高吞吐量代理。** 在我们的例子中，Dropbox 到 AWS 的通信，主要是 S3。这将使我们最终能够从我们的生产网络中删除所有 Squid 代理，从而为我们留下一个 L4/L7 数据平面。

One of the last things to migrate would be www.dropbox.com itself. After this migration, we can start decommissioning our edge Nginx deployments. An epoch would pass.

迁移的最后一件事是 www.dropbox.com 本身。在此迁移之后，我们可以开始停用我们的边缘 Nginx 部署。一个时代即将过去。

## Issues we encountered 

## 我们遇到的问题

Migration was not flawless, of course. But it didn’t lead to any notable outages. The hardest part of the migration was our API services. A lot of different devices communicate with Dropbox over our public API—everything from curl-/wget-powered shell scripts and embedded devices with custom HTTP/1.0 stacks, to every possible HTTP library out there. Nginx is a battle-tested de-facto industry standard. Understandably, most of the libraries implicitly depend on some of its behaviors. Along with a number of inconsistencies between Nginx and Envoy behaviors on which our api users depend, there were a number of bugs in Envoy and its libraries. All of them were quickly resolved and upstreamed by us with the community help.

当然，迁移并非完美无缺。但这并没有导致任何明显的中断。迁移中最困难的部分是我们的 API 服务。许多不同的设备通过我们的公共 API 与 Dropbox 通信——从 curl-/wget 驱动的 shell 脚本和带有自定义 HTTP/1.0 堆栈的嵌入式设备，到所有可能的 HTTP 库。 Nginx 是经过实战考验的事实上的行业标准。可以理解的是，大多数库都隐含地依赖于它的一些行为。除了我们的 api 用户所依赖的 Nginx 和 Envoy 行为之间的许多不一致之外，Envoy 及其库中还存在许多错误。在社区的帮助下，所有这些问题都得到了我们的快速解决和上游处理。

Here is just a gist of some the “unusual”/non-RFC behaviors:

这里只是一些“不寻常”/非 RFC 行为的要点：

- [**Merge slashes in URLs**](https://github.com/envoyproxy/envoy/pull/7621). URL normalization and slash merging is a very common feature for web-proxies. Nginx [enables slash normalization and slash merging by default](http://nginx.org/en/docs/http/ngx_http_core_module.html#merge_slashes) but Envoy did not support the latter. We submitted a patch upstream that add that functionality and allows users to opt-in by using the [merge\_slashes](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#envoy-api-field-config-filter-network-http-connection-manager-v2-httpconnectionmanager-merge-slashes) option.
- [**Ports in virtual host names**](https://github.com/envoyproxy/envoy/pull/10960). Nginx allows receiving Host header in both forms: either example.com or example.com:port. We had a couple of API users that used to rely on this behavior. First we worked around this by duplicating our vhosts in our configuration (with and without port) but later added an option to ignore the matching port on the Envoy side: [strip\_matching\_host\_port](https://www.envoyproxy .io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager -strip-matching-host-port).
- [**Transfer encoding case sensitivity**](https://github.com/envoyproxy/envoy/issues/10041). A tiny subset API client for some unknown reason used Transfer-Encoding: Chunked (note the capital “C”) header. This is technically valid, since RFC7230 states that Transfer-Encoding/TE headers are case insensitive. The fix was trivial and submitted to the upstream Envoy.
- [**Request that have both**](https://github.com/envoyproxy/envoy/issues/11398) [**Content-Length**](https://github.com/envoyproxy/envoy/issues/11398) [**and**](https://github.com/envoyproxy/envoy/issues/11398) [**Transfer-Encoding: c**](https://github.com/envoyproxy/envoy/issues/11398) [**hunked**](https://github.com/envoyproxy/envoy/issues/11398). Requests like that used to work with Nginx, but were broken by Envoy migration. [RFC7230 is a bit tricky there](https://tools.ietf.org/html/rfc7230#section-3.3.3), but general idea is web-servers should error these requests because they are likely “smuggled.” On the other hand, next sentence indicates that proxies should just remove the Content-Length header and forward the request. We've [extended http-parse to allow library users to opt-in into supporting such requests](https://github.com/nodejs/http-parser/issues/517) and currently working on adding the support to to Envoy itself.

- [**在 URL 中合并斜杠**](https://github.com/envoyproxy/envoy/pull/7621)。 URL 规范化和斜线合并是网络代理的一个非常常见的功能。 Nginx [默认启用斜杠规范化和斜杠合并](http://nginx.org/en/docs/http/ngx_http_core_module.html#merge_slashes) 但 Envoy 不支持后者。我们向上游提交了一个补丁，该补丁添加了该功能，并允许用户使用 [merge\_slashes](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#envoy-api-field-config-filter-network-http-connection-manager-v2-httpconnectionmanager-merge-slashes) 选项。
- [**虚拟主机名中的端口**](https://github.com/envoyproxy/envoy/pull/10960)。 Nginx 允许以两种形式接收 Host 标头：example.com 或 example.com:port。我们有几个 API 用户曾经依赖于这种行为。首先我们通过在我们的配置中复制我们的 vhosts（有和没有端口）来解决这个问题，但后来添加了一个选项来忽略 Envoy 端的匹配端口：[strip\_matching\_host\_port](https://www.envoyproxy .io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager -strip-matching-host-port)。
- [**传输编码区分大小写**](https://github.com/envoyproxy/envoy/issues/10041)。由于某种未知原因，一个很小的子集 API 客户端使用了 Transfer-Encoding: Chunked（注意大写的“C”)标头。这在技术上是有效的，因为 RFC7230 规定 Transfer-Encoding/TE 标头不区分大小写。修复是微不足道的，并提交给上游 Envoy。
- [**请求同时具有**](https://github.com/envoyproxy/envoy/issues/11398)[**Content-Length**](https://github.com/envoyproxy/envoy/问题/11398) [**和**](https://github.com/envoyproxy/envoy/issues/11398) [**传输编码：c**](https://github.com/envoyproxy/envoy/issues/11398) [**hunked**](https://github.com/envoyproxy/envoy/issues/11398)。像这样的请求曾经适用于 Nginx，但被 Envoy 迁移破坏了。 [RFC7230 在那里有点棘手](https://tools.ietf.org/html/rfc7230#section-3.3.3)，但一般的想法是网络服务器应该错误这些请求，因为它们很可能是“走私的”。另一方面，下一句表示代理应该删除 Content-Length 标头并转发请求。我们已经[扩展了 http-parse 以允许库用户选择支持此类请求](https://github.com/nodejs/http-parser/issues/517)，目前正在努力向 Envoy 添加支持本身。

It’s also worth mentioning some common configuration issues we’ve encountered:

还值得一提的是我们遇到的一些常见配置问题：

- **Circuit-breaking misconfiguration.** In our experience, if you are running Envoy as an inbound proxy, especially in a mixed HTTP/1&HTTP/2 environment, improperly set up circuit breakers can cause unexpected downtimes during traffic spikes or backend outages . Consider relaxing them if you are not using Envoy as a mesh proxy. It’s worth mentioning that by default, circuit-breaking limits in Envoy are pretty tight — be careful there! 

- **断路错误配置。** 根据我们的经验，如果您将 Envoy 作为入站代理运行，尤其是在混合 HTTP/1 和 HTTP/2 环境中，则断路器设置不当可能会在流量高峰或后端中断期间导致意外停机.如果您不使用 Envoy 作为网格代理，请考虑放宽它们。值得一提的是，默认情况下，Envoy 中的断路限制非常严格——在那里要小心！

- **Buffering.** Nginx allows request buffering on disk. This is especially useful in environments where you have legacy HTTP/1.0 backends that don’t understand chunked transfer encoding. Nginx could convert these into requests with Content-Length by buffering them on disk. Envoy has a Buffer filter, but without the ability to store data on disk we are restricted on how much we can buffer in memory.

- **缓冲。** Nginx 允许在磁盘上进行请求缓冲。这在您拥有不理解分块传输编码的旧式 HTTP/1.0 后端的环境中特别有用。 Nginx 可以通过将它们缓冲在磁盘上将它们转换为具有 Content-Length 的请求。 Envoy 有一个 Buffer 过滤器，但是没有将数据存储在磁盘上的能力，我们在内存中可以缓冲的数量受到限制。

If you're considering using Envoy as your Edge proxy, you would benefit from reading “ [Configuring](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge) [Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge) [as an edge proxy](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge). ” It does have security and resource limits that you would want to have on the most exposed part of your infrastructure.

如果你正在考虑使用 Envoy 作为你的边缘代理，你会受益于阅读“ [配置](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge) [Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge) [作为边缘代理](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge)。 ”它确实具有安全性和资源限制，您希望在基础架构的最暴露部分拥有这些限制。

## What’s next?

##  下一步是什么？

- [HTTP/3](https://tools.ietf.org/html/draft-ietf-quic-http) is getting closer for the prime time. Support for it was added to the most popular browsers (for now, [gated by a flags or command-line options](https://caniuse.com/#feat=http3)). Envoy support for it is also experimentally available. After we upgrade the [Linux kernel to support UDP acceleration](http://vger.kernel.org/lpc_net2018_talks/willemdebruijn-lpc2018-udpgso-paper-DRAFT-1.pdf), we will experiment with it on our Edge.
- Internal xDS-based load balancer and outlier detection. Currently, we are looking at using the combination of[Load Reporting service](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/load_stats/v3/lrs.proto)[(LRS) ](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/load_stats/v3/lrs.proto) and [Endpoint discovery service](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds) [(EDS)](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds) as building blocks for creating a common look-aside, load-aware loadbalancer for both Envoy and gRPC.
- WASM-based Envoy extensions. When Golang proxy-wasm SDK is available we can start writing Envoy extensions in Go which will give us access to a wide variety of internal Golang libs.
- Replacement for Bandaid. Unifying all Dropbox proxy layers under a single data-plane sounds very compelling. For that to happen we'll need to migrate a lot of Bandaid features (especially,[around loadbalancing](https://dropbox.tech/infrastructure/enhancing-bandaid-load-balancing-at-dropbox-by-leveraging-real-time-backend-server-load-information)) to Envoy. This is a long way but it’s our current plan.
- Envoy [mobile](https://envoy-mobile.github.io/). Eventually, we want to look into using Envoy in our mobile apps. It is very compelling from Traffic perspective to support a single stack with unified monitoring and modern capabilities (HTTP/3, gRPC, TLS1.3, etc) across all mobile platforms.

- [HTTP/3](https://tools.ietf.org/html/draft-ietf-quic-http)距离黄金时段越来越近了。对它的支持已添加到最流行的浏览器中（目前，[由标志或命令行选项控制](https://caniuse.com/#feat=http3))。 Envoy 对它的支持也是实验性的。在我们升级 [Linux 内核以支持 UDP 加速](http://vger.kernel.org/lpc_net2018_talks/willemdebruijn-lpc2018-udpgso-paper-DRAFT-1.pdf)后，我们将在我们的 Edge 上对其进行试验。
- 内部基于 xDS 的负载平衡器和异常值检测。目前，我们正在考虑使用[负载报告服务](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/load_stats/v3/lrs.proto)[(LRS) ](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/load_stats/v3/lrs.proto) 和 [端点发现服务](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds) [(EDS)](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds)作为构建块，用于为 Envoy 和 gRPC 创建通用的旁视、负载感知负载均衡器。
- 基于 WASM 的 Envoy 扩展。当 Golang proxy-wasm SDK 可用时，我们可以开始在 Go 中编写 Envoy 扩展，这将使我们能够访问各种内部 Golang 库。
- 替代创可贴。在单个数据平面下统一所有 Dropbox 代理层听起来非常引人注目。为此，我们需要迁移许多 Bandaid 功能（尤其是，[围绕负载平衡](https://dropbox.tech/infrastructure/enhancing-bandaid-load-balancing-at-dropbox-by-leveraging-real-time-backend-server-load-information)) 到 Envoy。这是很长的路要走，但这是我们目前的计划。
- 特使 [移动](https://envoy-mobile.github.io/)。最终，我们希望研究在我们的移动应用程序中使用 Envoy。从流量的角度来看，在所有移动平台上支持具有统一监控和现代功能（HTTP/3、gRPC、TLS1.3 等)的单一堆栈是非常引人注目的。

## Acknowledgements

## 致谢

This migration was truly a team effort. Traffic and Runtime teams were spearheading it but other teams heavily contributed: Agata Cieplik, Jeffrey Gensler, Konstantin Belyalov, Louis Opter, Naphat Sanguansin, Nikita V. Shirokov, Utsav Shah, Yi-Shu Tai, and of course the awesome Envoy community that helped us throughout that journey.

这次迁移确实是一项团队努力。交通和运行时团队带头，但其他团队也做出了巨大贡献：Agata Cieplik、Jeffrey Gensler、Konstantin Belyalov、Louis Opter、Naphat Sanguansin、Nikita V. Shirokov、Utsav Shah、Yi-Shu Tai，当然还有帮助过的很棒的 Envoy 社区我们在整个旅程中。

We also want to explicitly acknowledge the tech lead of the Runtime team **Ruslan Nigmatullin** whose actions as the Envoy evangelist, the author of the Envoy MVP, and the main driver from the software engineering side enabled this project to happen.

我们还想明确感谢 Runtime 团队 **Ruslan Nigmatullin** 的技术负责人，他作为 Envoy 布道者、Envoy MVP 的作者和软件工程方面的主要推动者的行为使该项目得以实现。

## We’re hiring! 

##  我们正在招聘！

If you’ve read this far, there’s a good chance that you actually enjoy digging deep into webservers/proxies and may enjoy working on the Dropbox Traffic team! Dropbox has a globally distributed Edge network, terabits of traffic, and millions of requests per second. All of it is managed by a [small team in Mountain View, CA](https://www.dropbox.com/jobs/listing/2034032?utm_source=tech&utm_medium=tech_blog&utm_campaign=infrastructure). 

如果您已经读到这里，那么您很有可能真正喜欢深入研究网络服务器/代理，并且可能喜欢在 Dropbox Traffic 团队工作！ Dropbox 拥有全球分布的 Edge 网络、TB 级的流量和每秒数百万个请求。所有这些都由一个[加利福尼亚州山景城的小团队](https://www.dropbox.com/jobs/listing/2034032?utm_source=tech&utm_medium=tech_blog&utm_campaign=infrastructure) 管理。

