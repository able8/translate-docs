# How we migrated Dropbox from Nginx to Envoy

By Alexey Ivanov and Oleg Guba • Jul 30, 2020

1. [Our legacy Nginx-based traffic infrastructure](http://dropbox.tech#-our-legacy-nginx-based-traffic-infrastructure)
2. [Why not Bandaid?](http://dropbox.tech#-why-not-bandaid)
3. [Our new Envoy-based traffic infrastructure](http://dropbox.tech#-our-new-envoy-based-traffic-infrastructure)
4. [Current state of our migration](http://dropbox.tech#-current-state-of-our-migration)
5. [Issues we encountered](http://dropbox.tech#-issues-we-encountered-)
6. [What’s next?](http://dropbox.tech#-whats-next)
7. [Acknowledgements](http://dropbox.tech#acknowledgements)
8. [We’re hiring!](http://dropbox.tech#-were-hiring)

In this blogpost we’ll talk about the old Nginx-based traffic infrastructure, its pain points, and the benefits we gained by migrating to [Envoy](https://www.envoyproxy.io/). We’ll compare Nginx to Envoy across many software engineering and operational dimensions. We’ll also briefly touch on the migration process, its current state, and some of the problems encountered on the way.

When we moved most of Dropbox traffic to Envoy, we had to seamlessly migrate a system that already handles tens of millions of open connections, millions of requests per second, and terabits of bandwidth. This effectively made us into one of the biggest Envoy users in the world.

Disclaimer: although we’ve tried to remain objective, quite a few of these comparisons are specific to Dropbox and the way our software development works: making bets on Bazel, gRPC, and C++/Golang.

Also note that we’ll cover the open source version of the Nginx, not its commercial version with additional features.

## Our legacy Nginx-based traffic infrastructure

Our Nginx configuration was mostly static and rendered with a combination of Python2, Jinja2, and YAML. Any change to it required a full re-deployment. All dynamic parts, such as upstream management and a stats exporter, were written in Lua. Any sufficiently complex logic was moved to [the next proxy layer, written in Go](https://dropbox.tech/infrastructure/meet-bandaid-the-dropbox-service-proxy).

Our post, “ [Dropbox](https://dropbox.tech/infrastructure/dropbox-traffic-infrastructure-edge-network) [traffic infrastructure: Edge network](https://dropbox.tech/infrastructure/dropbox-traffic-infrastructure-edge-network),” has a section about our legacy Nginx-based infrastructure.

Nginx served us well for almost a decade. But it didn’t adapt to our current development best-practices:

- Our internal and (private) external APIs are gradually migrating from REST to gRPC which requires all sorts of transcoding features from proxies.
- Protocol buffers became_de facto_ standard for service definitions and configurations.
- All software, regardless of the language, is built and tested with Bazel.
- Heavy involvement of our engineers on essential infrastructure projects in the open source community.

Also, operationally Nginx was quite expensive to maintain:

- Config generation logic was too flexible and split between YAML, Jinja2, and Python.
- Monitoring was a mix of Lua, log parsing, and system-based monitoring.
- An increased reliance on third party modules affected stability, performance, and the cost of subsequent upgrades.
- Nginx deployment and process management was quite different from the rest of the services. It relied a lot on other systems’ configurations: syslog, logrotate, etc, as opposed to being fully separate from the base system.

With all of that, for the first time in 10 years, we started looking for a potential replacement for Nginx.

## Why not Bandaid?

As we frequently mention, internally we rely heavily on the Golang-based proxy called [Bandaid](https://dropbox.tech/infrastructure/meet-bandaid-the-dropbox-service-proxy). It has a great integration with Dropbox infrastructure, because it has access to the vast ecosystem of internal Golang libraries: monitoring, service discoveries, rate limiting, etc. We considered migrating from Nginx to Bandaid but there are a couple of issues that prevent us from doing that:

- Golang is more resource intensive than C/C++. Low resource usage is especially important for us on the Edge since we can’t easily “auto-scale” our deployments there.
  - CPU overhead mostly comes from GC, HTTP parser and TLS, with the latter being less optimized than BoringSSL used by Nginx/Envoy.
  - The “goroutine-per-request” model and GC overhead greatly increase memory requirements in high-connection services like ours.
- No FIPS support for Go’s TLS stack.
- Bandaid does not have a community outside of Dropbox, which means that we can only rely on ourself for feature development.

With all that we’ve decided to start migrating our traffic infrastructure to Envoy instead.

## Our new Envoy-based traffic infrastructure

Let’s look into the main development and operational dimensions one by one, to see why we think Envoy is a better choice for us and what we gained by moving from Nginx to Envoy.

### Performance

Nginx’s [architecture](https://www.nginx.com/blog/inside-nginx-how-we-designed-for-performance-scale/) is event-driven and multi-process. It has support for [SO\_REUSEPORT](https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/), [EPOLLEXCLUSIVE](https://lore.kernel.org/patchwork/cover/543309/), and worker-to-CPU pinning. Although it is event-loop based, is it not fully non-blocking. This means some operations, like [opening a file](https://blog.cloudflare.com/how-we-scaled-nginx-and-saved-the-world-54-years-every-day/) or access/error logging, can potentially cause an event-loop stall ( [even](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [with](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [aio](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio) [, aio\_write, and thread pools enabled](http://nginx.org/en/docs/http/ngx_http_core_module.html#aio).) This leads to increased tail latencies, which can cause multi-second delays on spinning disk drives.

Envoy has a similar event-driven architecture, except it uses threads instead of processes. It also has SO\_REUSEPORT support (with a BPF filter support) and relies on libevent for event loop implementation (in other words, no fancy epoll(2) features like EPOLLEXCLUSIVE.) Envoy does not have any blocking IO operations in the event loop. Even logging is implemented in a non-blocking way, so that it does not cause stalls.

It looks like in theory Nginx and Envoy should have similar performance characteristics. But hope is not our strategy, so our first step was to run a diverse set of workload tests against similarly tuned Nginx and Envoy setups.

If you are interested in performance tuning, we describe our standard tuning guidelines in “ [Optimizing](https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency) [web servers for high throughput and low latency](https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency).” It involves everything from picking the hardware, to OS tunables, to library choices and web server configuration.

Our test results showed similar performance between Nginx and Envoy under most of our test workloads: high requests per second (RPS), high bandwidth, and a mixed low-latency/high-bandwidth gRPC proxying.

It is arguably very hard to make a good performance test. Nginx has [guidelines for performance testing](https://www.nginx.com/blog/testing-the-performance-of-nginx-and-nginx-plus-web-servers/), but these are not codified. Envoy also has [a guideline for benchmarking](https://www.envoyproxy.io/docs/envoy/latest/faq/performance/how_to_benchmark_envoy), and even some tooling under the [envoy](https://github.com/envoyproxy/envoy-perf) [-perf](https://github.com/envoyproxy/envoy-perf) project, but sadly the latter looks unmaintained.

We resorted to using our internal testing tool. It’s called “hulk” because of its reputation for smashing our services.

That said, there were a couple of notable differences in results:

- Nginx showed higher long tail latencies. This was mostly due to event loops stalls under heavy I/O, especially if used together withSO\_REUSEPORT since in that case [connections can be accepted on behalf of a currently blocked worker](https://blog.cloudflare.com/the-sad-state-of-linux-socket-balancing/#so_reuseporttotherescue).
- Nginx performance without stats collections is on par with Envoy, but our Lua stats collection slowed Nginx on the high-RPS test by a factor of 3. This was expected given our reliance onlua\_shared\_dict, which is synchronized across workers with a mutex.

We do understand how inefficient our stats collection was. We considered implementing something akin to FreeBSD’s [counter(9)](https://www.freebsd.org/cgi/man.cgi?query=counter&sektion=9&manpath=freebsd-release-ports#IMPLEMENTATION_DETAILS) in userspace: CPU pinning, per-worker lockless counters with a fetching routine that loops through all workers aggregating their individual stats. But we gave up on this idea, because if we wanted to instrument Nginx internals (e.g. all error conditions), it would mean supporting an enormous patch that would make subsequent upgrades a true hell.

Since Envoy does not suffer from either of these issues, after migrating to it we were able to release up to 60% of servers previously exclusively occupied by Nginx.

### Observability

Observability is the [most fundamental operational need](https://landing.google.com/sre/sre-book/chapters/part3/#fig_part-practices_reliability-hierarchy) for any product, but especially for such a foundational piece of infrastructure as a proxy. It is even more important during the migration period, so that any issue can be detected by the monitoring system rather than reported by frustrated users.

Non-commercial Nginx comes with a “ [stub](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html) [status](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html)” module that has 7 stats:

Copy


```
Active connections: 291
server accepts handled requests
16630948 16630948 31070465
Reading: 6 Writing: 179 Waiting: 106

```

This was definitely not enough, so we’ve added a simple log\_by\_lua handler that adds per-request stats based on headers and variables that are available in Lua: status codes, sizes, cache hits, etc. Here is an example of a simple stats-emitting function:

Copy


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

On top of all that, we had a separate exporter for gathering Nginx internal state: time since the last reload, number of workers, RSS/VMS sizes, TLS certificate ages, etc.

A typical Envoy setup provides us thousands of distinct metrics (in [prometheus format](https://prometheus.io/docs/instrumenting/exposition_formats/#text-based-format)) describing both proxied traffic and server’s internal state:

Copy


```
$ curl -s http://localhost:3990/stats/prometheus | wc -l
14819

```

This includes a myriad of stats with different aggregations:

- Per-cluster/per-upstream/per-vhost HTTP stats, including connection pool info and various timing histograms.
- Per-listener TCP/HTTP/TLS downstream connection stats.
- Various internal/runtime stats from basic version info and uptime to memory allocator stats and deprecated feature usage counters.

A special shoutout is needed for Envoy’s [admin interface](https://www.envoyproxy.io/docs/envoy/latest/operations/admin). Not only does it provide additional structured stats through /certs, /clusters, and /config\_dump endpoints, but there are also very important operational features:

- The ability to change error logging on the fly through[/logging](https://www.envoyproxy.io/docs/envoy/latest/operations/admin#post--logging). This allowed us to troubleshoot fairly obscure problems in a matter of minutes.
- /cpuprofiler, /heapprofiler, /contention which would surely be quite useful during the inevitable performance troubleshooting.
- /runtime\_modify  endpoint allows us to change set of configuration parameters without pushing new configuration, which could be used in feature gating, etc.

In addition to stats, Envoy also supports [pluggable tracing providers](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing). This is useful not only to our Traffic team, who own multiple load-balancing tiers, but also for application developers who want to track request latencies end-to-end from the edge to app servers.

Technically, Nginx also supports tracing through a third-party [OpenTracing integration](https://github.com/opentracing-contrib/nginx-opentracing) [,](https://github.com/opentracing-contrib/nginx-opentracing) but it is not under heavy development.

And last but not least, Envoy has the ability to [stream access logs over gRPC](https://www.envoyproxy.io/docs/envoy/latest/api-v3/data/accesslog/v3/accesslog.proto). This removes the burden of supporting syslog-to-hive bridges from our Traffic team. Besides, it’s way easier (and secure!) to spin up a generic gRPC service in Dropbox production than to add a custom TCP/UDP listener.

Configuration of access logging in Envoy, like everything else, happens through a gRPC management service, the [Access Log Service](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto) (ALS). Management services are the standard way of integrating the Envoy data plane with various services in production. This brings us to our next topic.

### Integration

Nginx’s approach to integration is best described as “Unix-ish.” Configuration is very static. It heavily relies on files (e.g. the config file itself, TLS certificates and tickets, allowlists/blocklists, etc.) and well-known industry protocols ( [logging](http://nginx.org/en/docs/syslog.html) [to syslog](http://nginx.org/en/docs/syslog.html) and auth sub-requests through HTTP). Such simplicity and backwards compatibility is a good thing for small setups, since Nginx can be easily automated with a couple of shell scripts. But as the system’s scale increases, testability and standardization become more important.

Envoy is far more opinionated in how the traffic dataplane should be integrated with its control plane, and hence with the rest of infrastructure. It encourages the use of [protobufs](https://developers.google.com/protocol-buffers) and gRPC by providing a stable API commonly referred as [xDS](https://docs.google.com/document/d/1xeVvJ6KjFBkNjVspPbY_PwEDHC7XPi0J5p1SqUXcCl8/edit#heading=h.c0uts5ftkk58). Envoy discovers its dynamic resources by [querying one or more of these xDS services](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration#arch-overview-dynamic-config).

Nowadays, the xDS APIs are evolving beyond Envoy: [Universal](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) [D](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) [ata Plane API](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) [(UDPA)](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) has the ambitious goal of “becoming de facto standard of L4/L7 loadbalancers.”

From our experience, this ambition works out well. We already use [Open Request Cost Aggregation](https://docs.google.com/document/d/1NSnK3346BkBo1JUU3I9I5NYYnaJZQPt8_Z_XCBCI3uA/edit) [(ORCA)](https://docs.google.com/document/d/1NSnK3346BkBo1JUU3I9I5NYYnaJZQPt8_Z_XCBCI3uA/edit) for our internal load testing, and are considering using UDPA for our non-Envoy loadbalancers e.g. our [Katran-based eBPF/XDP Layer-4 Load Balancer](https://github.com/facebookincubator/katran).

This is especially good for Dropbox, where all services internally already interact through gRPC-based APIs. We’ve implemented our own version of xDS control plane that integrates Envoy with our configuration management, service discovery, secret management, and route information.

For more information about Dropbox RPC, please read “ [Courier:](https://dropbox.tech/infrastructure/courier-dropbox-migration-to-grpc) [Dropbox migration to gRPC](https://dropbox.tech/infrastructure/courier-dropbox-migration-to-grpc).” There we describe in detail how we integrated service discovery, secret management, stats, tracing, circuit breaking, etc, with gRPC.

Here are **some** of the available xDS services, their Nginx alternatives, and our examples of how we use them:

- [Access Log Service](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto) [(ALS)](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/accesslog/v2/als.proto), as mentioned above, lets us dynamically configure access log destinations, encodings, and formats. Imagine a dynamic version of Nginx’s log\_format and access\_log.
- [Endpoint discovery service](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds) [(EDS)](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds) provides information about cluster members. This is analogous to a dynamically updated list of upstream block’s server entries (e.g. for Lua that would be a  [balancer\_by\_lua\_block](https://github.com/openresty/lua-nginx-module#balancer_by_lua_block)) in the Nginx config. In our case we proxied this to our internal service discovery.
- [Secret discovery service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret) [(SDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret) provides various TLS-related information that would cover various ssl\_\* directives (and respectively [ssl\_\*\_by\_lua\_block](https://github.com/openresty/lua-nginx-module#ssl_certificate_by_lua_block).)  We adapted this interface to our secret distribution service.
- [Runtime Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/operations/runtime#config-runtime-rtds) [(RTDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/operations/runtime#config-runtime-rtds) is providing runtime flags. Our implementation of this functionality in Nginx was quite hacky, based on checking the existence of various files from Lua. This approach can quickly become inconsistent between the individual servers. Envoy’s default implementation is also filesystem-based, but we instead pointed our RTDS xDS API to our distributed configuration storage. That way we can control whole clusters at once (through a tool with a sysctl-like interface) and there are no accidental inconsistencies between different servers.
- [Route discovery service](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/rds#config-http-conn-man-rds) [(RDS)](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/rds#config-http-conn-man-rds) maps routes to virtual hosts, and allows additional configuration for headers and filters. In Nginx terms, these would be analogous to a dynamic location block with set\_header/proxy\_set\_header and a proxy\_pass. On lower proxy tiers we autogenerate these directly from our service definition configs.

For an example of Envoy’s integration with an existing production system, here is a canonical example of how to [integrate](https://www.envoyproxy.io/learn/service-discovery) [Envoy](https://www.envoyproxy.io/learn/service-discovery) [with a custom service discovery](https://www.envoyproxy.io/learn/service-discovery). There are also a couple of open source Envoy control-plane implementations, such as [Istio](https://istio.io/) and the less complex [go-control-plane](https://github.com/envoyproxy/go-control-plane).

Our homegrown Envoy control plane implements an increasing number of xDS APIs. It is deployed as a normal gRPC service in production, and acts as an adapter for our infrastructure building blocks. It does this through a set of common Golang libraries to talk to internal services and expose them through a stable xDS APIs to Envoy. The whole process does not involve any filesystem calls, signals, cron, logrotate, syslog, log parsers, etc.

### Configuration

Nginx has the undeniable advantage of a simple human-readable configuration. But this win gets lost as config gets more complex and begins to be code-generated.

As mentioned above, our Nginx config is generated through a mix of Python2, Jinja2, and YAML. Some of you may have seen or even written a variation of this in erb, pug, Text::Template, or maybe even m4:

Copy


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

This problem is arguably fixable, but there were a couple of foundational ones:

- There is no declarative description for the config format. If we wanted to programmatically generate and validate configuration, we would need to invent it ourselves.
- Config that is syntactically valid could still be invalid from a C code standpoint. For example, some of the buffer-related variables have limitations on values, restrictions on alignment, and interdependencies with other variables. To semantically validate a config we needed to run it throughnginx -t.

Envoy, on the other hand, has a unified data-model for configs: all of its configuration is defined in Protocol Buffers. This not only solves the data modeling problem, but also adds typing information to the config values. Given that protobufs are first class citizens in Dropbox production, and a common way of describing/configuring services, this makes integration _so_ much easier.

Our new config generator for Envoy is based on protobufs and Python3. All data modeling is done in proto files, while all the logic is in Python. Here’s an example:

Copy


```
from dropbox.proto.envoy.extensions.filters.http.gzip.v3.gzip_pb2 import Gzip
from dropbox.proto.envoy.extensions.filters.http.compressor.v3.compressor_pb2 import Compressor

def default_gzip_config(
    compression_level: Gzip.CompressionLevel.Enum = Gzip.CompressionLevel.DEFAULT,
    ) -> Gzip:
        return Gzip(
            # Envoy's default is 6 (Z_DEFAULT_COMPRESSION).
            compression_level=compression_level,
            # Envoy's default is 4k (12 bits). Nginx uses 32k (MAX_WBITS, 15 bits).
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

- Note the[Python3 type annotations](https://www.python.org/dev/peps/pep-0484/) in that code!  Coupled with [mypy-protobuf protoc plugin](https://github.com/dropbox/mypy-protobuf), these provide end-to-end typing inside the config generator. IDEs capable of checking them will immediately highlight typing mismatches.

There are still cases where a type-checked protobuf can be logically invalid. In the example above, gzip window\_bits can only take values between 9 and 15. This kind of restriction can be easily defined with a help of [protoc-gen-validate protoc plugin](https://github.com/envoyproxy/protoc-gen-validate):

Copy


```
google.protobuf.UInt32Value window_bits = 9 [(validate.rules).uint32 = {lte: 15 gte: 9}];

```

Finally, an implicit benefit of using a formally defined configuration model is that it organically leads to the documentation being collocated with the configuration definitions. [Here](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) [’](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) [s an example from](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50) [gzip.proto](https://github.com/envoyproxy/envoy/blob/master/api/envoy/extensions/filters/http/gzip/v3/gzip.proto#L50):

Copy


```
// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values.
// use more memory, but are faster and produce better compression results. The default value is 5.
google.protobuf.UInt32Value memory_level = 1 [(validate.rules).uint32 = {lte: 9 gte: 1}];
```

For those of you thinking about using protobufs in your production systems, but worried you may lack a schema-less representation, here’s a good article from Envoy core developer Harvey Tuch about how to work around this using google.protobuf.Struct and google.protobuf.Any: “ [Dynamic](https://blog.envoyproxy.io/dynamic-extensibility-and-protocol-buffers-dcd0bf0b8801) [extensibility and Protocol Buffers](https://blog.envoyproxy.io/dynamic-extensibility-and-protocol-buffers-dcd0bf0b8801).”

### Extensibility

Extending Nginx beyond what’s possible with standard configuration usually requires writing a C module. Nginx’s [development guide](http://nginx.org/en/docs/dev/development_guide.html) provides a solid introduction to the available building blocks. That said, this approach is relatively heavyweight. In practice, it takes a fairly senior software engineer to safely write an Nginx module.

In terms of infrastructure available for module developers, they can expect basic containers like hash tables/queues/rb-trees, (non-RAII) memory management, and hooks for all phases of request processing. There are also couple of external libraries like pcre, zlib, openssl, and, of course, libc.

For more lightweight feature extension, Nginx provides [Perl](http://nginx.org/en/docs/http/ngx_http_perl_module.html#perl) and [Javascript](http://nginx.org/en/docs/http/ngx_http_js_module.html) interfaces. Sadly, both are fairly limited in their abilities, mostly restricted to the content phase of request processing.

The most commonly used extension method adopted by the community is based on a third-party l [ua-](https://github.com/openresty/lua-nginx-module) [nginx](https://github.com/openresty/lua-nginx-module) [-module](https://github.com/openresty/lua-nginx-module) and various [OpenResty libraries](https://github.com/openresty/). This approach can be hooked in at pretty much any phase of request processing. We used log\_by\_lua for stats collection, and balancer\_by\_lua for dynamic backend reconfiguration.

In theory, Nginx provides the ability to develop [modules in C++](http://lxr.nginx.org/source/src/misc/ngx_cpp_test_module.cpp). In practice, it lacks proper C++ interfaces/wrappers for all the primitives to make this worthwhile. There are nonetheless some [community attempts at it](https://github.com/chronolaw/ngx_cpp_dev). These are far from ready for production, though.

Envoy’s main extension mechanism is through C++ plugins. The process is [not as well documented](https://blog.envoyproxy.io/how-to-write-envoy-filters-like-a-ninja-part-1-d166e5abec09) as in Nginx’s case, but it is simpler. This is partially due to:

- **Clean and well-commented interfaces.** C++ classes act as natural extension and documentation points. For example, [checkout the HTTP filter interface](https://github.com/envoyproxy/envoy/blob/master/include/envoy/http/filter.h).
- **C++14 language and standard library.** From basic language features like templates and lambda functions, to type-safe containers and algorithms. In general, writing modern C++14 is not much different from using Golang or, with a stretch, one may even say Python.
- **Features beyond C++14 and its stdlib.** Provided by the [abseil](https://abseil.io/about/intro) library, these include drop-in replacements from newer C++ standards, mutexes with built-in [static deadlock detection](http://clang.llvm.org/docs/ThreadSafetyAnalysis.html) and debug support, additional/more efficient containers, [and much more](https://abseil.io/about/philosophy).

For specifics, here’s a [canonical example of an HTTP Filter module](https://github.com/envoyproxy/envoy-filter-example/tree/master/http-filter-example).

We were able to integrate Envoy with [Vortex2](https://dropbox.tech/infrastructure/monitoring-server-applications-with-vortex) [(our](https://dropbox.tech/infrastructure/monitoring-server-applications-with-vortex) [monitoring framework)](https://dropbox.tech/infrastructure/monitoring-server-applications-with-vortex) with only 200 lines of code by simply implementing the Envoy [stats](https://github.com/envoyproxy/envoy/tree/master/include/envoy/stats) interface.

Envoy [also has Lua support](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/lua_filter) through [moonjit](https://github.com/moonjit/moonjit), a LuaJIT fork with improved Lua 5.2 support. Compared to Nginx’s 3rd-party Lua integration it has far fewer capabilities and hooks. This makes Lua in Envoy far less attractive due to the cost of additional complexity in developing, testing, and troubleshooting interpreted code. Companies that specialize in Lua development may disagree, but in our case we decided to avoid it and use C++ exclusively for Envoy extensibility.

What distinguishes Envoy from the rest of web servers is its emerging support for [WebAssembly](https://developer.mozilla.org/en-US/docs/WebAssembly/Concepts) (WASM) — a fast, portable, and secure extension mechanism. WASM is not meant to be used directly, but as a compilation target for any general-purpose programming language. Envoy implements a [WebAssembly for Proxies specification](https://github.com/proxy-wasm/spec/blob/master/abi-versions/vNEXT/README.md) (and also includes reference [Rust](https://github.com/proxy-wasm/proxy-wasm-rust-sdk) and [C++](https://github.com/proxy-wasm/proxy-wasm-cpp-sdk) SDKs) that describes the boundary between WASM code and a generic L4/L7 proxy. That separation between the proxy and extension code allows for secure sandboxing, while WASM low-level compact binary format allows for near native efficiency. On top of that, in Envoy proxy-wasm extensions are integrated with xDS. This allows dynamic updates and even potential A/B testing.

The “ [Extending](https://youtu.be/XdWmm_mtVXI?t=779) [Envoy](https://youtu.be/XdWmm_mtVXI?t=779) [with WebAssembly](https://youtu.be/XdWmm_mtVXI?t=779)” presentation from Kubecon’19 (remember that time when we had non-virtual conferences?) has a nice overview of  WASM in Envoy and its potential uses. It also hints at performance levels of 60-70% of native C++ code.

With WASM, service providers get a safe and efficient way of running customers’ code on their edge. Customers get the benefit of portability: Their extensions can run on any cloud that implements the proxy-wasm ABI. Additionally, it allows your users to use any language as long as it can be compiled to WebAssembly. This enables them to use a broader set of non-C++ libraries, securely and efficiently.

Istio is putting a lot of resources into WebAssembly development: they already have an experimental version of the WASM-based telemetry extension and the [WebAssemblyHub community](https://webassemblyhub.io/) for sharing extensions. You can read about it in detail in [“Redefining](https://istio.io/latest/blog/2020/wasm-announce/) [extensibility in proxies - introducing WebAssembly to](https://istio.io/latest/blog/2020/wasm-announce/) [Envoy](https://istio.io/latest/blog/2020/wasm-announce/) [and Istio](https://istio.io/latest/blog/2020/wasm-announce/) [.](https://istio.io/latest/blog/2020/wasm-announce/) [”](https://istio.io/latest/blog/2020/wasm-announce/)

Currently, we don’t use WebAssembly at Dropbox. But this might change when the Go SDK for proxy-wasm is available.

### Building and Testing

By default, Nginx is built using a custom [shell-based configuration system](https://github.com/nginx/nginx/tree/master/auto) and make-based build system. This is simple and elegant, but it took quite a bit of effort to integrate it into [B](https://dropbox.tech/infrastructure/continuous-integration-and-deployment-with-bazel) [azel-built monorepo](https://dropbox.tech/infrastructure/continuous-integration-and-deployment-with-bazel) to get all the benefits of incremental, distributed, hermetic, and reproducible builds.

Google open [-](https://nginx.googlesource.com/nginx/) sourced their [B](https://nginx.googlesource.com/nginx/) [azel-built](https://nginx.googlesource.com/nginx/) [Nginx](https://nginx.googlesource.com/nginx/) [version](https://nginx.googlesource.com/nginx/) which consists of Nginx, BoringSSL, PCRE, ZLIB, and Brotli library/module.

Testing-wise, Nginx has a set of Perl-driven [integration tests](http://hg.nginx.org/nginx-tests) in a separate repository and no unit tests.

Given our heavy usage of Lua and absence of a built-in unit testing framework, we resorted to testing using mock configs and a simple Python-based test driver:

Copy


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

On the Envoy side, the main build system is already Bazel. So integrating it with our monorepo was trivial: Bazel easily allows [adding external dependencies](https://docs.bazel.build/versions/master/external.html).

We also use [copybara](https://github.com/google/copybara) scripts to sync protobufs for both Envoy and udpa. Copybara is handy when you need to do simple transformations without the need to forever maintain a large patchset.

With Envoy we have the flexibility of using either unit tests (based on gtest/gmock) with a set of [pre-written mocks](https://github.com/envoyproxy/envoy/tree/master/test/mocks), or Envoy’s [integration test framework](https://github.com/envoyproxy/envoy/tree/master/test/integration), or both. There’s no need anymore to rely on slow end-to-end integration tests for every trivial change.

[gtest](https://github.com/google/googletest) is a fairly well-known unit-test framework used by Chromium and LLVM, among others. If you want to know more about googletest there are good intros for both [googletest](https://github.com/google/googletest/blob/master/googletest/docs/primer.md) and [googlemock](https://chromium.googlesource.com/external/github.com/google/googletest/+/HEAD/googlemock/docs/cook_book.md).

Open source Envoy development [requires changes to have 100% unit test coverage](https://github.com/envoyproxy/envoy/blob/master/CONTRIBUTING.md#submitting-a-pr). Tests are automatically triggered for each pull request via the [Azure CI Pipeline](https://dev.azure.com/cncf/envoy/_build?view=pipelines).

It’s also a common practice to micro-benchmark performance-sensitive code with [google/becnhmark](https://github.com/google/benchmark):

Copy


```
$ bazel run --compilation_mode=opt test/common/upstream:load_balancer_benchmark -- --benchmark_filter=".*LeastRequestLoadBalancerChooseHost.*"
BM_LeastRequestLoadBalancerChooseHost/100/1/1000000          848 ms          449 ms            2 mean_hits=10k relative_stddev_hits=0.0102051 stddev_hits=102.051
...

```

After switching to Envoy, we began to rely exclusively on unit tests for our internal module development:

Copy


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

Bazel is one of the best things that ever happened to our developer experience. It has a very steep learning curve and is a large upfront investment, but it has a very high return on that investment: [incremental builds](https://docs.bazel.build/versions/master/guide.html#correct-incremental-rebuilds), [remote caching](https://docs.bazel.build/versions/master/remote-caching.html), [distributed builds/tests](https://docs.bazel.build/versions/master/remote-execution.html), etc.

One of the less discussed benefits of Bazel is that it gives us an ability to [query](https://docs.bazel.build/versions/master/query-how-to.html) [and even augment](https://docs.bazel.build/versions/master/skylark/aspects.html) the dependency graph. A programmatic interface to the dependency graph, coupled with a common build system across all languages, is a very powerful feature. It can be used as a foundational building block for things like linters, code generation, vulnerability tracking, deployment system, etc.

### Security

Nginx’s code surface is quite small, with minimal external dependencies. It’s typical to see only 3 external dependencies on the resulting binary: zlib (or [one of its faster variants](https://github.com/cloudflare/zlib)), a TLS library, and PCRE. Nginx has a custom implementation of all protocol parsers, the event library, and they even went as far as to re-implement some libc functions.

At some point Nginx was considered so secure that it was used as a default web server in OpenBSD. Later two development communities had a falling out, which lead to the creation of  httpd. You can read about the motivation behind that move in BSDCon’s “ [Introducing](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) [OpenBSD](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) [’s](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf) [new httpd](https://www.openbsd.org/papers/httpd-asiabsdcon2015.pdf).”

This minimalism paid off in practice. Nginx has only had 30 [vulnerabilities and exposures](https://nginx.org/en/security_advisories.html) reported in more than 11 years.

Envoy, on the other hand, has way more code, especially when you consider that that C++ code is far more dense than the basic C used for Nginx. It also incorporates millions of lines of code from external dependencies. Everything from event notification to protocol parsers is offloaded to 3rd party libraries. This increases attack surface and bloats the resulting binary.

To counteract this, Envoy relies heavily on modern security practices. It uses [AddressSanitizer](https://github.com/google/sanitizers/wiki/AddressSanitizer), [ThreadSanitizer](https://github.com/google/sanitizers/wiki/ThreadSanitizerCppManual), and [MemorySanitizer](https://github.com/google/sanitizers/wiki/MemorySanitizer). Its developers even went beyond that and adopted [fuzzing](https://bugs.chromium.org/p/oss-fuzz/issues/list?q=label%3AProj-envoy&sort=-id).

Any opensource project that is critical to the global IT infrastructure can be accepted to the [OSS-Fuzz](https://github.com/google/oss-fuzz)—a free platform for automated fuzzing. To learn more about it, see “ [OSS-Fuzz](https://google.github.io/oss-fuzz/architecture/) [/ Architecture](https://google.github.io/oss-fuzz/architecture/).”

In practice, though, all these precautions do not fully counteract the increased code footprint. As a result, Envoy has had [22 security advisories in the](https://github.com/envoyproxy/envoy/security/advisories) [p](https://github.com/envoyproxy/envoy/security/advisories) [ast 2 years](https://github.com/envoyproxy/envoy/security/advisories).

Envoy's [security release policy is described in great detail](https://github.com/envoyproxy/envoy/security/policy), and in [postmortems](https://github.com/envoyproxy/envoy/tree/master/security/postmortems) for selected vulnerabilities. Envoy is also a participant in [Google’s Vulnerability Reward Program](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/security/google_vrp#arch-overview-google-vrp) [(VRP)](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/security/google_vrp#arch-overview-google-vrp). Open to all security researchers, VRP provides rewards for vulnerabilities discovered and reported according to their rules.

For a practical example of how some of these vulnerabilities can be potentially exploited, see this writeup about CVE-2019–18801: “ [Exploiting](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792) [an](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792) [Envoy](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792) [heap vulnerability](https://blog.envoyproxy.io/exploiting-an-envoy-heap-vulnerability-96173d41792).”

To counteract the increased vulnerability risk, we use best binary hardening security practices from our upstream OS vendors [Ubuntu](https://wiki.ubuntu.com/Security/Features) and [Debian](https://wiki.debian.org/Hardening). We defined a special hardened build profile for all edge-exposed binaries. It includes ASLR, stack protectors, and symbol table hardening:

Copy


```
build:hardened --force_pic
build:hardened --copt=-fstack-clash-protection
build:hardened --copt=-fstack-protector-strong
build:hardened --linkopt=-Wl,-z,relro,-z,now
```

Forking web-servers, like Nginx, in most environments [have issues with stack protector](http://hmarco.org/renewssp/data/Preventing_brute_force_attacks_against_stack_canary_protection_on_networking_servers-Paper.pdf). Since master and worker processes share the same stack canary, and on canary verification failure worker process is killed, the canary can be brute-forced bit-by-bit in about 1000 tries. Envoy, which uses threads as a concurrency primitive, is not affected by this attack.

We also want to harden third-party dependencies where we can. We use [BoringSSL in FIPS mode](https://boringssl.googlesource.com/boringssl/+/master/crypto/fipsmodule/FIPS.md), which includes startup self-tests and integrity checking of the binary. We’re also considering running ASAN-enabled binaries on some of our edge canary servers.

### Features

Here comes the most opinionated part of the post, brace yourself.

Nginx began as a web server specialized on serving static files with minimal resource consumption. Its functionality is top of the line there: static serving, caching (including thundering herd protection), and range caching.

On the proxying side, though, Nginx lacks features needed for modern infrastructures. There’s no HTTP/2 to backends. gRPC proxying is available but without connection multiplexing. There’s no support for gRPC transcoding. On top of that, Nginx’s “open-core” model restricts features that can go into an open source version of the proxy. As a result, some of the critical features like statistics are not available in the “community” version.

Envoy, by contrast, has evolved as an ingress/egress proxy, used frequently for gRPC-heavy environments. Its web-serving functionality is rudimentary: [no file serving](https://github.com/envoyproxy/envoy/issues/378), still [work-in-progress caching](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/cache/v3alpha/cache.proto.html), neither [brotli](https://github.com/envoyproxy/envoy/issues/4429) nor pre-compression. For these use cases we still have a small fallback Nginx setup that Envoy uses as an upstream cluster.

When HTTP cache in Envoy becomes production-ready, we could move most of static-serving use cases to it, using S3 instead of filesystem for long-term storage. To read more about eCache design, see “ [eCache:](https://docs.google.com/document/d/1WPuim_GzhfdsnIj_tf-fIeutK0jO4aVQfVrLJFoLN3g/view#heading=h.wjxw6fq7wefi) [a multi-backend HTTP cache](https://docs.google.com/document/d/1WPuim_GzhfdsnIj_tf-fIeutK0jO4aVQfVrLJFoLN3g/view#heading=h.wjxw6fq7wefi) [for Envoy](https://docs.google.com/document/d/1WPuim_GzhfdsnIj_tf-fIeutK0jO4aVQfVrLJFoLN3g/view#heading=h.wjxw6fq7wefi).”

Envoy also has native support for many gRPC-related capabilities:

- **gRPC proxying.** This is a basic capability that allowed us to use gRPC end-to-end for our applications (e.g. Dropbox desktop client.)
- **HTTP/2 to backends.** This feature allows us to greatly reduce the number of TCP connections between our traffic tiers, reducing memory consumption and keepalive traffic.
- [gRPC → HTTP bridge](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_http1_bridge_filter) (\+ [reverse](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_http1_reverse_bridge_filter).)  These allowed us to expose legacy HTTP/1 applications using a modern gRPC stack.
- [gRPC-WEB](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_web_filter). This feature allowed us to use gRPC end-to-end even in the environments where middleboxes (firewalls, IDS, etc) don’t yet support HTTP/2.
- [gRPC JSON transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter). This enables us to transcode all inbound traffic, including [Dropbox public APIs](https://www.dropbox.com/developers/documentation/http/overview), from REST into gRPC.

In addition, Envoy can also be used as an outbound proxy. We used it to unify a couple of other use cases:

- Egress proxy: since Envoy [added support for the HTTP CONNECT method](https://github.com/envoyproxy/envoy/issues/1451), it can be used as a drop-in replacement for Squid proxies. We’ve begun to replace our outbound Squid installations with Envoy. This not only greatly improves visibility, but also reduces operational toil by unifying the stack with a common dataplane and observability (no more parsing logs for stats.)
- Third-party software service discovery: we are relying on the[Courier gRPC libraries](https://dropbox.tech/infrastructure/courier-dropbox-migration-to-grpc) in our software instead of using Envoy as a service mesh. But we do use Envoy in one-off cases where we need to connect an open source service with our service discovery with minimal effort. For example, Envoy is used as a service discovery sidecar in our analytics stack. Hadoop can dynamically discover its name and journal nodes. [Superset](https://github.com/apache/incubator-superset) can discover airflow, presto, and hive backends. [Grafana](https://grafana.com/) can discover its MySQL database.

## Community

Nginx development is quite centralized. Most of its development happens behind closed doors. There’s some external activity on the [nginx-devel](http://mailman.nginx.org/pipermail/nginx-devel/) mailing list, and there are occasional development-related discussions on the [official bug tracker](https://trac.nginx.org/nginx).

There is an #nginx channel on FreeNode. Feel free to join it for more interactive [community](https://www.nginx.com/resources/wiki/community/irc/) conversations.

Envoy development is open and decentralized: coordinated through GitHub issues/pull requests, [mailing list](https://groups.google.com/g/envoy-dev), and [community meetings](https://goo.gl/5Cergb).

There is also quite a bit of community activity on Slack. You can get your invite [here](https://envoyslack.cncf.io).

It’s hard to quantify the development styles and engineering community, so let’s look at a specific example of HTTP/3 development.

Nginx [QUIC and HTTP/3 implementation](https://hg.nginx.org/nginx-quic/) was [recently presented by F5](https://www.nginx.com/blog/introducing-technology-preview-nginx-support-for-quic-http-3/). The code is clean, with zero external dependencies. But the development process itself was rather opaque. Half a year before that, [Cloudflare came up with their own HTTP/3 implementation for](https://blog.cloudflare.com/experiment-with-http-3-using-nginx-and-quiche/) [Nginx](https://blog.cloudflare.com/experiment-with-http-3-using-nginx-and-quiche/). As a result, the community now has two separate experimental versions of HTTP/3 for Nginx.

In Envoy’s case, HTTP/3 implementation is also a work in progress, based on chromium’s " [quiche](https://docs.google.com/document/d/19qcrwAa8hVYZv2r8zZ7SgkylivAQNQ7E3loMJ-vk9_k/edit)" (QUIC, HTTP, Etc.) library. The project’s status is tracked in the [GitHub issue](https://github.com/envoyproxy/envoy/issues/2557). The [de](https://docs.google.com/document/d/1dEo19y-trABuW2x6-T564LmK7Ld-BPXZOlnR4df9KVU/edit#heading=h.w2fjl4fs3sex) [sign doc](https://docs.google.com/document/d/1dEo19y-trABuW2x6-T564LmK7Ld-BPXZOlnR4df9KVU/edit#heading=h.w2fjl4fs3sex) was publicly available way before patches were completed. Remaining work that would benefit from community involvement is tagged with “ [help](https://github.com/envoyproxy/envoy/issues?q=is%3Aopen+label%3Aarea%2Fquic+label%3A%22help+wanted%22) [wanted](https://github.com/envoyproxy/envoy/issues?q=is%3Aopen+label%3Aarea%2Fquic+label%3A%22help+wanted%22).”

As you can see, the latter structure is much more transparent and greatly encourages collaboration. For us, this means that we managed to upstream lots of small to medium changes to Envoy–everything from [operational improvements](https://github.com/envoyproxy/envoy/pull/10286/files) and [performance optimizations](https://github.com/envoyproxy/envoy/pull/9556) to [new gRPC transcoding features](https://github.com/envoyproxy/envoy/pull/10673) and [load](https://github.com/envoyproxy/envoy/pull/11006) [balancing changes](https://github.com/envoyproxy/envoy/pull/11006).

## Current state of our migration

We’ve been running Nginx and Envoy side-by-side for over half a year and gradually switching traffic from one to another with DNS. By now we have migrated a wide variety of workloads to Envoy:

- **Ingress high-throughput services.** All file data from Dropbox desktop client is served via end-to-end gRPC through Envoy. By switching to Envoy we’ve also slightly improved users’ performance, due to better connection reuse from the edge.
- **Ingress high-RPS services.** This is all file metadata for Dropbox desktop client. We get the same benefits of end-to-end gRPC, plus the removal of the connection pool, which means we are not bounded by one request per connection at a time.
- **Notification and telemetry services.** Here we handle all real-time notifications, so these servers have millions of HTTP connections (one for each active client.) Notification services can now be implemented via streaming gRPC instead of an expensive long-poll method.
- **Mixed high-throughput/high-RPS services.** API traffic (both metadata and data itself.) This allows us to start thinking about public gRPC APIs. We may even switch to transcoding our existing REST-based APIs right on the Edge.
- **Egress high-throughput proxies.** In our case, the Dropbox to AWS communication, mostly S3. This would allow us to eventually remove all Squid proxies from our production network, leaving us with a single L4/L7 data plane.

One of the last things to migrate would be www.dropbox.com itself. After this migration, we can start decommissioning our edge Nginx deployments. An epoch would pass.

## Issues we encountered

Migration was not flawless, of course. But it didn’t lead to any notable outages. The hardest part of the migration was our API services. A lot of different devices communicate with Dropbox over our public API—everything from curl-/wget-powered shell scripts and embedded devices with custom HTTP/1.0 stacks, to every possible HTTP library out there. Nginx is a battle-tested de-facto industry standard. Understandably, most of the libraries implicitly depend on some of its behaviors. Along with a number of inconsistencies between Nginx and Envoy behaviors on which our api users depend, there were a number of bugs in Envoy and its libraries. All of them were quickly resolved and upstreamed by us with the community help.

Here is just a gist of some the “unusual”/non-RFC behaviors:

- [**Merge slashes in URLs**](https://github.com/envoyproxy/envoy/pull/7621). URL normalization and slash merging is a very common feature for web-proxies. Nginx [enables slash normalization and slash merging by default](http://nginx.org/en/docs/http/ngx_http_core_module.html#merge_slashes) but Envoy did not support the latter. We submitted a patch upstream that add that functionality and allows users to opt-in by using the [merge\_slashes](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#envoy-api-field-config-filter-network-http-connection-manager-v2-httpconnectionmanager-merge-slashes) option.
- [**Ports in virtual host names**](https://github.com/envoyproxy/envoy/pull/10960). Nginx allows receiving Host header in both forms: either example.com or example.com:port. We had a couple of API users that used to rely on this behavior. First we worked around this by duplicating our vhosts in our configuration (with and without port) but later added an option to ignore the matching port on the Envoy side: [strip\_matching\_host\_port](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-strip-matching-host-port).
- [**Transfer encoding case sensitivity**](https://github.com/envoyproxy/envoy/issues/10041). A tiny subset API client for some unknown reason used Transfer-Encoding: Chunked (note the capital “C”) header. This is technically valid, since RFC7230 states that Transfer-Encoding/TE headers are case insensitive. The fix was trivial and submitted to the upstream Envoy.
- [**Request that have both**](https://github.com/envoyproxy/envoy/issues/11398) [**Content-Length**](https://github.com/envoyproxy/envoy/issues/11398) [**and**](https://github.com/envoyproxy/envoy/issues/11398) [**Transfer-Encoding: c**](https://github.com/envoyproxy/envoy/issues/11398) [**hunked**](https://github.com/envoyproxy/envoy/issues/11398). Requests like that used to work with Nginx, but were broken by Envoy migration. [RFC7230 is a bit tricky there](https://tools.ietf.org/html/rfc7230#section-3.3.3), but general idea is web-servers should error these requests because they are likely “smuggled.” On the other hand, next sentence indicates that proxies should just remove the Content-Length header and forward the request. We’ve [extended http-parse to allow library users to opt-in into supporting such requests](https://github.com/nodejs/http-parser/issues/517) and currently working on adding the support to to Envoy itself.

It’s also worth mentioning some common configuration issues we’ve encountered:

- **Circuit-breaking misconfiguration.** In our experience, if you are running Envoy as an inbound proxy, especially in a mixed HTTP/1&HTTP/2 environment, improperly set up circuit breakers can cause unexpected downtimes during traffic spikes or backend outages. Consider relaxing them if you are not using Envoy as a mesh proxy. It’s worth mentioning that by default, circuit-breaking limits in Envoy are pretty tight — be careful there!
- **Buffering.** Nginx allows request buffering on disk. This is especially useful in environments where you have legacy HTTP/1.0 backends that don’t understand chunked transfer encoding. Nginx could convert these into requests with Content-Length by buffering them on disk. Envoy has a Buffer filter, but without the ability to store data on disk we are restricted on how much we can buffer in memory.

If you’re considering using Envoy as your Edge proxy, you would benefit from reading “ [Configuring](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge) [Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge) [as an edge proxy](https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge).”  It does have security and resource limits that you would want to have on the most exposed part of your infrastructure.

## What’s next?

- [HTTP/3](https://tools.ietf.org/html/draft-ietf-quic-http) is getting closer for the prime time. Support for it was added to the most popular browsers (for now, [gated by a flags or command-line options](https://caniuse.com/#feat=http3)). Envoy support for it is also experimentally available. After we upgrade the [Linux kernel to support UDP acceleration](http://vger.kernel.org/lpc_net2018_talks/willemdebruijn-lpc2018-udpgso-paper-DRAFT-1.pdf), we will experiment with it on our Edge.
- Internal xDS-based load balancer and outlier detection. Currently, we are looking at using the combination of[Load Reporting service](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/load_stats/v3/lrs.proto) [(LRS)](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/load_stats/v3/lrs.proto) and [Endpoint discovery service](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds) [(EDS)](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#arch-overview-service-discovery-types-eds) as building blocks for creating a common look-aside, load-aware loadbalancer for both Envoy and gRPC.
- WASM-based Envoy extensions. When Golang proxy-wasm SDK is available we can start writing Envoy extensions in Go which will give us access to a wide variety of internal Golang libs.
- Replacement for Bandaid. Unifying all Dropbox proxy layers under a single data-plane sounds very compelling. For that to happen we’ll need to migrate a lot of Bandaid features (especially,[around loadbalancing](https://dropbox.tech/infrastructure/enhancing-bandaid-load-balancing-at-dropbox-by-leveraging-real-time-backend-server-load-information)) to Envoy. This is a long way but it’s our current plan.
- Envoy [mobile](https://envoy-mobile.github.io/). Eventually, we want to look into using Envoy in our mobile apps. It is very compelling from Traffic perspective to support a single stack with unified monitoring and modern capabilities (HTTP/3, gRPC, TLS1.3, etc) across all mobile platforms.

## Acknowledgements

This migration was truly a team effort. Traffic and Runtime teams were spearheading it but other teams heavily contributed: Agata Cieplik, Jeffrey Gensler, Konstantin Belyalov, Louis Opter, Naphat Sanguansin, Nikita V. Shirokov, Utsav Shah, Yi-Shu Tai, and of course the awesome Envoy community that helped us throughout that journey.

We also want to explicitly acknowledge the tech lead of the Runtime team **Ruslan Nigmatullin** whose actions as the Envoy evangelist, the author of the Envoy MVP, and the main driver from the software engineering side enabled this project to happen.

## We’re hiring!

If you’ve read this far, there’s a good chance that you actually enjoy digging deep into webservers/proxies and may enjoy working on the Dropbox Traffic team! Dropbox has a globally distributed Edge network, terabits of traffic, and millions of requests per second. All of it is managed by a [small team in Mountain View, CA](https://www.dropbox.com/jobs/listing/2034032?utm_source=tech&utm_medium=tech_blog&utm_campaign=infrastructure).
