# Production Checklist for Redis on Kubernetes

[Jun 29, 2020](https://medium.com/swlh/production-checklist-for-redis-on-kubernetes-60173d5a5325?source=post_page-----60173d5a5325--------------------------------) · 7 min read

![img](https://miro.medium.com/max/1400/1*cKRFZsFXbyrOyKNB5MLn9A.png)

Redis is a popular open-source in-memory data store and cache that has become an integral part of building a scalable microservice system. While all  the major cloud providers offer a fully-managed Redis service ([Amazon ElastiCache](https://aws.amazon.com/elasticache/), [Azure Cache for Redis](https://azure.microsoft.com/en-us/services/cache/), and [GCP Memorystore](https://cloud.google.com/memorystore)), it can also be easily deployed in Kubernetes if you need more control  over the Redis configurations. Redis provides decent performance out of  the box, but if you are preparing to run a production workload, make  sure to go over this checklist before you go live.

# Hardware Optimization

As with any database, Redis performance is tied to the underlying VM specifications. Create a node pool with **memory-optimized** machines with high network bandwidth limits to minimize latency between clients and Redis servers. Redis is single-threaded, which means **fast CPUs** with large caches (e.g. VMs backed by Intel Skylake or Cascade Lake)  perform better, and adding multi-cores do not directly improve  performance. If your workload is mostly small objects (<10 KB), speed or RAM and memory bandwidth are not as critical to optimizing Redis  performance. You can read more on Redis performance on different  hardware [here](https://redis.io/topics/benchmarks).

# Choose A Deployment Method

To deploy a Redis cluster on Kubernetes, you can either use [Bitnami’s Redis Helm Chart](https://github.com/bitnami/charts/tree/master/bitnami/redis) or one of the Redis Operators. While I am normally in favor of  Kubernetes Operators, it doesn’t seem like there is a popular and mature Redis Operator compared to Bitnami’s Helm Chart. Redis Labs, the  creators of Redis, provides an official [Redis Enterprise Kubernetes Operator](https://docs.redislabs.com/latest/platforms/kubernetes/kubernetes-with-operator/), but if you need a true open-source version, you can choose between [Spotahome’s Operator](https://github.com/spotahome/redis-operator) or the [Operator from Amadeus IT Group](https://github.com/AmadeusITGroup/Redis-Operator) (in alpha). I have no personal experience with either of these operators, but the engineers at Flant wrote a [blog post on their failures](https://medium.com/flant-com/redis-kubernetes-operator-and-data-analysis-tools-afce55b02123) using the Redis Operator by Spotahome.

![img](https://miro.medium.com/max/60/0*gegAWGzw5szaUykG?q=20)

![img](https://miro.medium.com/max/1400/0*gegAWGzw5szaUykG)

Redis Enterprise vs. Redis Open-Source — Image Credit: [Redis Labs](https://redislabs.com/redis-enterprise/technology/redis-enterprise-cluster-architecture/)

Bitnami supports two deployment options for Redis: a **master-slave cluster with Redis Sentinel** and a **Redis cluster** topology with sharding. If you have a high read throughput, using the  master-slave cluster helps offload the read operations to the slave  pods. The sentinels are configured to promote a slave pod to a master in the case of a failure. On the other hand, Redis cluster shards data  across multiple instances and is a great fit when memory requirements  exceed the limits for a single master with CPU becoming the bottleneck  (>100GB). Redis cluster also supports high availability with each  master connected to one or more slave pods. When the master pod crashes, one of the slaves will be promoted to master.

![img](https://miro.medium.com/max/60/0*Z8EXKcgNhmckzX4x.png?q=20)

![img](https://miro.medium.com/max/1400/0*Z8EXKcgNhmckzX4x.png)

Redis Architectures — Image Credit: [Redis Labs](https://redislabs.com/redis-enterprise/technology/redis-enterprise-cluster-architecture/)

# Persistent Storage

Redis stores some data in ephemeral storage, but using persistent volumes are critical for high-availability. Redis provides two persistence options:

- **RDB** (Redis Database File): point-in-time snapshots
- **AOF** (Append Only File): logs of every Redis operation

It’s possible to combine both types of persistence, but it’s important to  understand the tradeoffs between the two options for the best  performance.

RDB is a compact snapshot optimized for a typical backup operation. RDB  backup operation has minimal impact on Redis performance since the  parent process forks a child process to create the backup. In disaster  recovery scenarios, RDB boots up faster than AOF since the file size is  smaller and more compact. However, since RDB is essentially a  point-in-time snapshot, it will lose data in between RDB snapshots if a  failure occurs.

AOF, on the other hand, keeps a log of every operation and is more durable  than RDB as it can be configured to fsnyc on every second or query. In  the event of an outage, AOF can run through the log and replay every  operation. Redis can also automatically and safely rewrite the AOF in  the background if it gets too big. The downside to AOF is file size and  speed. With replication turned on, sometimes the slaves cannot sync with the master fast enough to revive all the data. AOF can also be much  slower than RDB depending on the fsync policy.

Redis Helm chart enables AOF and disables RDB by default, but you can  override the configmap with different fsync strategies or RDB  persistence:

```
configmap: |-  
  # Enable AOF https://redis.io/topics/persistence#append-only-file 
  appendonly yes  # Disable RDB persistence, AOF persistence already enabled.  
  save ""
```

For a deep-dive into Redis persistence, make sure to read the [Redis Persistence Post on the official Redis website](https://redis.io/topics/persistence).

# Disable THP

After deploying Redis to Kubernetes, you will most likely see the following warning message:

```
WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled
```

In order to disable THP, you can add an init-container to run the command  or deploy a DaemonSet to the node pool running Redis. For example, I  have a database node pool running on GKE so I deployed a DaemonSet like  the following:

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

After setting up the Redis cluster, you can now benchmark the performance using various, well-maintained tools:

- [Redis Benchmark](https://redis.io/topics/benchmarks): included with Redis
- [Memtier Benchmark](https://github.com/RedisLabs/memtier_benchmark): also developed by Redis Labs
- [Redis Memory Analyzer](https://github.com/gamenet/redis-memory-analyzer): Python tool by GameNet
- [YCSB](https://github.com/brianfrankcooper/YCSB): from Yahoo Cloud
- [PerfKit Benchmarker](https://github.com/GoogleCloudPlatform/PerfKitBenchmarker): from Google Cloud
- [Redis RDB tools](https://github.com/sripathikrishnan/redis-rdb-tools): parses Redis dump.rdb files
- [Harvest](https://github.com/31z4/harvest): samples Redis keys and shows top key prefixes

Chie Hayashida from Google Cloud wrote an [excellent guide to using YCSB to benchmark Redis](https://cloud.google.com/blog/products/databases/performance-tuning-best-practices-for-memorystore-for-redis?utm_source=feedburner&utm_medium=email&utm_campaign=Feed:%2Bgoogleblog%2FCNkG%2B(Google%2BCloud%2BPlatform%2BBlog)) performance for Memorystore (GCP’s managed Redis). The same tools can  be used to test Redis in Kubernetes. Use port forwarding to map Redis to localhost and run YCSB with various usage patterns. Combine this result with memory analyzer tools to fine-tune Redis performance:

1. [Compress data](https://docs.redislabs.com/latest/ri/memory-optimizations/compress-values/) using Snappy/LZO (low latency) or GZIP (maximum compression) for long strings or JSON/XML values.
2. Use [MessagePack format](https://msgpack.org/index.html) instead of JSON for efficient serialization.
3. Set an appropriate eviction policy: use **allkeys** to evict all the keys or **volatile** for those with TTL/expiration field set (you can specify in the extra flags section in the Helm chart)

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

Finally, connect Redis metrics to Prometheus (or a monitoring tool of your  choice) to detect performance degradations and alerts. Bitnami Helm  Chart uses Bitnami’s Redis Exporter by default, but you can also use [Oliver006’s popular Redis Exporter chart](https://github.com/oliver006/redis_exporter). There is also a [companion Grafana chart](https://grafana.com/grafana/dashboards/763) to visualize all the metrics.

As for configuring alerts, Bitnami provides some example rules:


Or you can use [Redis alerts from Awesome Prometheus project](https://awesome-prometheus-alerts.grep.to/rules#redis). If you need a guide on setting up Prometheus on Kubernetes, you can check out the [Practical Monitoring with Prometheus and Grafana series](https://medium.com/@yitaek/practical-monitoring-with-prometheus-grafana-part-i-22d0f172f993).

At this point, you should have a production-ready Redis cluster running on Kubernetes. As usage grows, some performance degradation is expected.  Make sure to run the benchmark and memory analyzer tool periodically to  deal with the new load.

[The Startup](https://medium.com/swlh?source=post_sidebar--------------------------post_sidebar-----------)

Get smarter at building your thing. Join The Startup’s +740K followers.

## Sign up for Top 10 Stories

### By The Startup

Get smarter at building your thing. Subscribe to receive The Startup's top  10 most read stories — delivered straight into your inbox, twice a  month. [Take a look.](https://medium.com/swlh/newsletters/top-10-stories?source=newsletter_v3_promo--------------------------newsletter_v3_promo-----------)

Emails will be sent to bibleapple@gmail.com.
