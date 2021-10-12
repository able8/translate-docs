# Production Checklist for Redis on Kubernetes

# Kubernetes 上 Redis 的生产清单

[Jun 29, 2020](https://medium.com/swlh/production-checklist-for-redis-on-kubernetes-60173d5a5325?source=post_page-----60173d5a5325--------------------------------) · 7 min read

![img](https://miro.medium.com/max/1400/1*cKRFZsFXbyrOyKNB5MLn9A.png)

Redis is a popular open-source in-memory data store and cache that has become an integral part of building a scalable microservice system. While all  the major cloud providers offer a fully-managed Redis service ([Amazon ElastiCache](https://aws.amazon.com/elasticache/), [Azure Cache for Redis](https://azure.microsoft.com/en-us/services/cache/), and [GCP Memorystore](https://cloud.google.com/memorystore)), it can also be easily deployed in Kubernetes if you need more control  over the Redis configurations. Redis provides decent performance out of  the box, but if you are preparing to run a production workload, make  sure to go over this checklist before you go live.

Redis 是一种流行的开源内存数据存储和缓存，已成为构建可扩展微服务系统不可或缺的一部分。虽然所有主要的云提供商都提供完全托管的 Redis 服务 ([Amazon ElastiCache](https://aws.amazon.com/elasticache/)，[Azure Cache for Redis](https://azure.microsoft.com/en-us/services/cache/) 和 [GCP Memorystore](https://cloud.google.com/memorystore))，如果您需要更多地控制 Redis 配置，它也可以轻松部署在 Kubernetes 中。 Redis 提供了开箱即用的良好性能，但如果您准备运行生产工作负载，请确保在上线之前查看此清单。

# Hardware Optimization

# 硬件优化

As with any database, Redis performance is tied to the underlying VM specifications. Create a node pool with **memory-optimized** machines with high network bandwidth limits to minimize latency between clients and Redis servers. Redis is single-threaded, which means **fast CPUs** with large caches (e.g. VMs backed by Intel Skylake or Cascade Lake)  perform better, and adding multi-cores do not directly improve  performance. If your workload is mostly small objects (<10 KB), speed or RAM and memory bandwidth are not as critical to optimizing Redis  performance. You can read more on Redis performance on different  hardware [here](https://redis.io/topics/benchmarks).

与任何数据库一样，Redis 性能与底层 VM 规范相关。使用具有高网络带宽限制的**内存优化**机器创建节点池，以最大限度地减少客户端和 Redis 服务器之间的延迟。 Redis 是单线程的，这意味着具有大缓存的**快速 CPU**（例如由 Intel Skylake 或 Cascade Lake 支持的 VM）性能更好，并且添加多核不会直接提高性能。如果您的工作负载主要是小对象 (<10 KB)，则速度或 RAM 和内存带宽对于优化 Redis 性能并不重要。您可以在 [此处](https://redis.io/topics/benchmarks) 阅读有关不同硬件上的 Redis 性能的更多信息。

# Choose A Deployment Method

# 选择部署方式

To deploy a Redis cluster on Kubernetes, you can either use [Bitnami’s Redis Helm Chart](https://github.com/bitnami/charts/tree/master/bitnami/redis) or one of the Redis Operators. While I am normally in favor of  Kubernetes Operators, it doesn’t seem like there is a popular and mature Redis Operator compared to Bitnami’s Helm Chart. Redis Labs, the  creators of Redis, provides an official [Redis Enterprise Kubernetes Operator](https://docs.redislabs.com/latest/platforms/kubernetes/kubernetes-with-operator/), but if you need a true open- source version, you can choose between [Spotahome's Operator](https://github.com/spotahome/redis-operator) or the [Operator from Amadeus IT Group](https://github.com/AmadeusITGroup/Redis-Operator) (in alpha). I have no personal experience with either of these operators, but the engineers at Flant wrote a [blog post on their failures](https://medium.com/flant-com/redis-kubernetes-operator-and-data-analysis-tools-afce55b02123) using the Redis Operator by Spotahome.

要在 Kubernetes 上部署 Redis 集群，您可以使用 [Bitnami 的 Redis Helm Chart](https://github.com/bitnami/charts/tree/master/bitnami/redis) 或 Redis Operators 之一。虽然我通常支持 Kubernetes Operators，但与 Bitnami 的 Helm Chart 相比，似乎没有流行和成熟的 Redis Operator。 Redis 的创建者 Redis Labs 提供了一个官方的 [Redis Enterprise Kubernetes Operator](https://docs.redislabs.com/latest/platforms/kubernetes/kubernetes-with-operator/)，但是如果你需要一个真正的开放——源版本，您可以选择 [Spotahome's Operator](https://github.com/spotahome/redis-operator) 或 [Operator from Amadeus IT Group](https://github.com/AmadeusITGroup/Redis-Operator)（在阿尔法)。我对这两个算子都没有个人经验，但是 Flant 的工程师写了一篇[关于他们失败的博客文章](https://medium.com/flant-com/redis-kubernetes-operator-and-data-analysis-tools-afce55b02123) 使用 Spotahome 的 Redis Operator。

![img](https://miro.medium.com/max/60/0*gegAWGzw5szaUykG?q=20)

![img](https://miro.medium.com/max/1400/0*gegAWGzw5szaUykG)

Redis Enterprise vs. Redis Open-Source — Image Credit: [Redis Labs](https://redislabs.com/redis-enterprise/technology/redis-enterprise-cluster-architecture/)

Redis 企业与 Redis 开源 — 图片来源：[Redis Labs](https://redislabs.com/redis-enterprise/technology/redis-enterprise-cluster-architecture/)

Bitnami supports two deployment options for Redis: a **master-slave cluster with Redis Sentinel** and a **Redis cluster** topology with sharding. If you have a high read throughput, using the  master-slave cluster helps offload the read operations to the slave  pods. The sentinels are configured to promote a slave pod to a master in the case of a failure. On the other hand, Redis cluster shards data  across multiple instances and is a great fit when memory requirements  exceed the limits for a single master with CPU becoming the bottleneck  (>100GB). Redis cluster also supports high availability with each  master connected to one or more slave pods. When the master pod crashes, one of the slaves will be promoted to master.

Bitnami 支持两种 Redis 部署选项：**带有 Redis Sentinel** 的主从集群**和带有分片的**Redis 集群**拓扑。如果您有很高的读取吞吐量，使用主从集群有助于将读取操作卸载到从属 Pod。哨兵被配置为在失败的情况下将一个从属 pod 提升为主控。另一方面，Redis 集群跨多个实例对数据进行分片，非常适合当内存需求超过单个主服务器的限制而 CPU 成为瓶颈 (>100GB) 时。 Redis 集群还支持高可用性，每个主节点连接到一个或多个从节点 pod。当 master pod 崩溃时，其中一个 slave 将被提升为 master。

![img](https://miro.medium.com/max/60/0*Z8EXKcgNhmckzX4x.png?q=20)

![img](https://miro.medium.com/max/1400/0*Z8EXKcgNhmckzX4x.png)

Redis Architectures — Image Credit: [Redis Labs](https://redislabs.com/redis-enterprise/technology/redis-enterprise-cluster-architecture/)

Redis 架构 — 图片来源：[Redis Labs](https://redislabs.com/redis-enterprise/technology/redis-enterprise-cluster-architecture/)

# Persistent Storage

# 持久化存储

Redis stores some data in ephemeral storage, but using persistent volumes are critical for high-availability. Redis provides two persistence options:

Redis 将一些数据存储在临时存储中，但使用持久卷对于高可用性至关重要。 Redis 提供了两个持久化选项：

- **RDB** (Redis Database File): point-in-time snapshots
- **AOF** (Append Only File): logs of every Redis operation 

- **RDB**（Redis 数据库文件）：时间点快照
- **AOF**（Append Only File）：每个Redis操作的日志

It’s possible to combine both types of persistence, but it’s important to  understand the tradeoffs between the two options for the best  performance.

可以将两种类型的持久性结合起来，但重要的是要了解两种选择之间的权衡以获得最佳性能。

RDB is a compact snapshot optimized for a typical backup operation. RDB  backup operation has minimal impact on Redis performance since the  parent process forks a child process to create the backup. In disaster  recovery scenarios, RDB boots up faster than AOF since the file size is  smaller and more compact. However, since RDB is essentially a  point-in-time snapshot, it will lose data in between RDB snapshots if a  failure occurs.

RDB 是针对典型备份操作优化的紧凑快照。 RDB 备份操作对 Redis 性能的影响最小，因为父进程派生一个子进程来创建备份。在灾难恢复场景中，RDB 的启动速度比 AOF 快，因为文件更小更紧凑。然而，由于 RDB 本质上是一个时间点快照，如果发生故障，它将丢失 RDB 快照之间的数据。

AOF, on the other hand, keeps a log of every operation and is more durable  than RDB as it can be configured to fsnyc on every second or query. In  the event of an outage, AOF can run through the log and replay every  operation. Redis can also automatically and safely rewrite the AOF in  the background if it gets too big. The downside to AOF is file size and  speed. With replication turned on, sometimes the slaves cannot sync with the master fast enough to revive all the data. AOF can also be much  slower than RDB depending on the fsync policy.

另一方面，AOF 保留每个操作的日志，并且比 RDB 更持久，因为它可以配置为每秒或查询的 fsnyc。在发生中断的情况下，AOF 可以运行日志并重播每个操作。如果 AOF 过大，Redis 还可以在后台自动安全地重写 AOF。 AOF 的缺点是文件大小和速度。启用复制后，有时从站无法以足够快的速度与主站同步以恢复所有数据。根据 fsync 策略，AOF 也可能比 RDB 慢得多。

Redis Helm chart enables AOF and disables RDB by default, but you can  override the configmap with different fsync strategies or RDB  persistence:

Redis Helm chart 默认启用 AOF 并禁用 RDB，但您可以使用不同的 fsync 策略或 RDB 持久性覆盖 configmap：

```
configmap: |-
  # Enable AOF https://redis.io/topics/persistence#append-only-file
  appendonly yes  # Disable RDB persistence, AOF persistence already enabled.
  save ""
```

For a deep-dive into Redis persistence, make sure to read the [Redis Persistence Post on the official Redis website](https://redis.io/topics/persistence).

要深入了解 Redis 持久性，请务必阅读 [Redis 官方网站上的 Redis 持久性帖子](https://redis.io/topics/persistence)。

# Disable THP

# 禁用 THP

After deploying Redis to Kubernetes, you will most likely see the following warning message:

将 Redis 部署到 Kubernetes 后，您很可能会看到以下警告消息：

```
WARNING you have Transparent Huge Pages (THP) support enabled in your kernel.This will create latency and memory usage issues with Redis.To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot.Redis must be restarted after THP is disabled
```

In order to disable THP, you can add an init-container to run the command  or deploy a DaemonSet to the node pool running Redis. For example, I  have a database node pool running on GKE so I deployed a DaemonSet like  the following:

为了禁用 THP，您可以添加一个 init-container 来运行命令或将 DaemonSet 部署到运行 Redis 的节点池中。例如，我有一个在 GKE 上运行的数据库节点池，所以我部署了一个 DaemonSet，如下所示：

```
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: thp-disable
  namespace: kube-system
spec:
  selector:
    matchLabels:
      name: thp-disable
  template:
    metadata:
      labels:
        name: thp-disable
    spec:
      tolerations:
      - key: "database"
        operator: "Equal"
        value: "true"
        effect: "NoSchedule"
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: cloud.google.com/gke-nodepool
                operator: In
                values:
                - database
      restartPolicy: Always
      terminationGracePeriodSeconds: 1
      volumes:
        - name: host-sys
          hostPath:
            path: /sys
      initContainers:
        - name: disable-thp
          image: busybox
          volumeMounts:
            - name: host-sys
              mountPath: /host-sys
          command: ["sh", "-c", "echo never >/host-sys/kernel/mm/transparent_hugepage/enabled"]
      containers:
        - name: busybox
          image: busybox
          command: ["watch", "-n", "600", "cat", "/sys/kernel/mm/transparent_hugepage/enabled"]
```

# Benchmarking Performance

# 基准性能

After setting up the Redis cluster, you can now benchmark the performance using various, well-maintained tools:

设置 Redis 集群后，您现在可以使用各种维护良好的工具对性能进行基准测试：

- [Redis Benchmark](https://redis.io/topics/benchmarks): included with Redis
- [Memtier Benchmark](https://github.com/RedisLabs/memtier_benchmark): also developed by Redis Labs
- [Redis Memory Analyzer](https://github.com/gamenet/redis-memory-analyzer): Python tool by GameNet
- [YCSB](https://github.com/brianfrankcooper/YCSB): from Yahoo Cloud
- [PerfKit Benchmarker](https://github.com/GoogleCloudPlatform/PerfKitBenchmarker): from Google Cloud
- [Redis RDB tools](https://github.com/sripathikrishnan/redis-rdb-tools): parses Redis dump.rdb files
- [Harvest](https://github.com/31z4/harvest): samples Redis keys and shows top key prefixes 

- [Redis Benchmark](https://redis.io/topics/benchmarks)：包含在Redis中
- [Memtier Benchmark](https://github.com/RedisLabs/memtier_benchmark)：同样由Redis Labs开发
- [Redis 内存分析器](https://github.com/gamenet/redis-memory-analyzer)：GameNet 的 Python 工具
- [YCSB](https://github.com/brianfrankcooper/YCSB)：来自雅虎云
- [PerfKit Benchmarker](https://github.com/GoogleCloudPlatform/PerfKitBenchmarker)：来自谷歌云
- [Redis RDB tools](https://github.com/sripathikrishnan/redis-rdb-tools)：解析Redis dump.rdb文件
- [Harvest](https://github.com/31z4/harvest)：Redis 密钥示例并显示顶级密钥前缀

Chie Hayashida from Google Cloud wrote an [excellent guide to using YCSB to benchmark Redis](https://cloud.google.com/blog/products/databases/performance-tuning-best-practices-for-memorystore-for-redis?utm_source=feedburner&utm_medium=email&utm_campaign=Feed:%2Bgoogleblog%2FCNkG%2B(Google%2BCloud%2BPlatform%2BBlog)) performance for Memorystore (GCP's managed Redis). The same tools can  be used to test Redis in Kubernetes. Use port forwarding to map Redis to localhost and run YCSB with various usage patterns. Combine this result with memory analyzer tools to fine-tune Redis performance:

谷歌云的 Chie Hayashida 写了一篇[使用 YCSB 对 Redis 进行基准测试的优秀指南](https://cloud.google.com/blog/products/databases/performance-tuning-best-practices-for-memorystore-for-redis?utm_source=feedburner&utm_medium=email&utm_campaign=Feed:%2Bgoogleblog%2FCNkG%2B(Google%2BCloud%2BPlatform%2BBlog)) Memorystore（GCP 的托管 Redis)的性能。相同的工具可用于在 Kubernetes 中测试 Redis。使用端口转发将 Redis 映射到本地主机并以各种使用模式运行 YCSB。将此结果与内存分析器工具结合以微调 Redis 性能：

1. [Compress data](https://docs.redislabs.com/latest/ri/memory-optimizations/compress-values/) using Snappy/LZO (low latency) or GZIP (maximum compression) for long strings or JSON/ XML values.
2. Use [MessagePack format](https://msgpack.org/index.html) instead of JSON for efficient serialization.
3. Set an appropriate eviction policy: use **allkeys** to evict all the keys or **volatile** for those with TTL/expiration field set (you can specify in the extra flags section in the Helm chart)

1. [压缩数据](https://docs.redislabs.com/latest/ri/memory-optimizations/compress-values/) 对长字符串或 JSON 使用 Snappy/LZO（低延迟）或 GZIP（最大压缩)/ XML 值。
2.使用[MessagePack格式](https://msgpack.org/index.html)代替JSON进行高效序列化。
3. 设置适当的驱逐策略：使用 **allkeys** 驱逐所有密钥或 **volatile** 用于设置 TTL/expiration 字段的密钥（您可以在 Helm 图表的额外标志部分指定）

```
master:
  ## Redis command arguments
  ##
  ## Can be used to specify command line arguments, for example:
  ##
  command: "/run.sh"
  ## Additional Redis configuration for the master nodes
  ## ref: https://redis.io/topics/config
  ##
  configmap:
  ## Redis additional command line flags
  ##
  ## Can be used to specify command line flags, for example:
  ## extraFlags:
  ##  - "--maxmemory-policy volatile-ttl"
  ##  - "--repl-backlog-size 1024mb"
```

# Monitoring

# 监控

Finally, connect Redis metrics to Prometheus (or a monitoring tool of your  choice) to detect performance degradations and alerts. Bitnami Helm  Chart uses Bitnami’s Redis Exporter by default, but you can also use [Oliver006’s popular Redis Exporter chart](https://github.com/oliver006/redis_exporter). There is also a [companion Grafana chart](https://grafana.com/grafana/dashboards/763) to visualize all the metrics.

最后，将 Redis 指标连接到 Prometheus（或您选择的监控工具）以检测性能下降和警报。 Bitnami Helm Chart 默认使用 Bitnami 的 Redis Exporter，但您也可以使用 [Oliver006 流行的 Redis Exporter chart](https://github.com/oliver006/redis_exporter)。还有一个 [companion Grafana chart](https://grafana.com/grafana/dashboards/763) 来可视化所有指标。

As for configuring alerts, Bitnami provides some example rules:

至于配置警报，Bitnami 提供了一些示例规则：

Or you can use [Redis alerts from Awesome Prometheus project](https://awesome-prometheus-alerts.grep.to/rules#redis). If you need a guide on setting up Prometheus on Kubernetes, you can check out the [Practical Monitoring with Prometheus and Grafana series](https://medium.com/@yitaek/practical-monitoring-with-prometheus-grafana-part-i-22d0f172f993).

或者您可以使用 [来自 Awesome Prometheus 项目的 Redis 警报](https://awesome-prometheus-alerts.grep.to/rules#redis)。如果您需要有关在 Kubernetes 上设置 Prometheus 的指南，您可以查看 [Practical Monitoring with Prometheus and Grafana 系列](https://medium.com/@yitaek/practical-monitoring-with-prometheus-grafana-part-i-22d0f172f993)。

At this point, you should have a production-ready Redis cluster running on Kubernetes. As usage grows, some performance degradation is expected. Make sure to run the benchmark and memory analyzer tool periodically to  deal with the new load.

此时，您应该有一个在 Kubernetes 上运行的生产就绪 Redis 集群。随着使用量的增长，预计会出现一些性能下降。确保定期运行基准测试和内存分析器工具来处理新负载。

[The Startup](https://medium.com/swlh?source=post_sidebar--------------------------post_sidebar-----------)

Get smarter at building your thing. Join The Startup’s +740K followers.

更聪明地构建你的东西。加入 The Startup 的 +740K 粉丝。

## Sign up for Top 10 Stories

## 注册前 10 个故事

### By The Startup

### 由初创公司

Get smarter at building your thing. Subscribe to receive The Startup's top  10 most read stories — delivered straight into your inbox, twice a  month. [Take a look.](https://medium.com/swlh/newsletters/top-10-stories?source=newsletter_v3_promo--------------------------newsletter_v3_promo-----------)

Emails will be sent to bibleapple@gmail.com. 

更聪明地构建你的东西。订阅以接收 The Startup 的前 10 个阅读最多的故事 - 直接发送到您的收件箱，每月两次。 电子邮件将发送至 bibleapple@gmail.com。

