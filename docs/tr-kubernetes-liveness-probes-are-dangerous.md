# Liveness Probes are Dangerous

# 存活探针是危险的

Posted: 2019-09-28 17:55

Kubernetes `livenessProbe` can be dangerous. I recommend to avoid them unless you have a clear use case and understand the consequences. This post looks at both Liveness and Readiness Probes and describes some "DOs" and "DON'Ts"

Kubernetes `livenessProbe` 可能很危险。我建议避免使用它们，除非您有明确的用例并了解后果。这篇文章着眼于 Liveness 和 Readiness Probes，并描述了一些“应该做的”和“不应该做的”

My colleague Sandor recently tweeted about common mistakes he sees, including wrong Readiness/Liveness Probe usage:

[![../galleries/twitter-sszuecs-mistakes.png](https://srcco.de/galleries/twitter-sszuecs-mistakes.png)](https://twitter.com/sszuecs/status/1175377157343907840)

我的同事 Sandor 最近发了一条关于他看到的常见错误的推文，包括错误的就绪/存活探针用法：

)

A wrong `livenessProbe` setting can  worsen high-load situations (cascading failure + potential long  container/app start) and lead to other negative consequences like  bringing down dependencies (see also [my recent post about K3s+ACME rate limits](https://srcco.de/posts/k3s-outage-traefik-acme-lets-encrypt-local-path.html)). A Liveness Probe in combination with an external DB health check dependency is the worst situation: **a single DB hiccup will restart all your containers!**

错误的 `livenessProbe` 设置会加剧高负载情况（级联故障 + 潜在的容器/应用程序启动时间过长）并导致其他负面后果，例如降低依赖关系（另请参阅 [我最近关于 K3s+ACME 速率限制的帖子](https：//srcco.de/posts/k3s-outage-traefik-acme-lets-encrypt-local-path.html))。 Liveness Probe 与外部数据库健康检查依赖相结合是最糟糕的情况：**单个数据库小问题将重新启动所有容器！**

A blanket statement of "don't use Liveness Probes" is not helpful, so let's look at what Readiness and Liveness Probe are for. *NOTE: most of the following text was initially put together for Zalando's internal developer documentation.*

“不要使用 Liveness Probe”的笼统声明没有帮助，所以让我们看看 Readiness 和 Liveness Probe 的用途。 *注意：以下大部分文字最初是为 Zalando 的内部开发人员文档整理的。*

## Readiness and Liveness Probes

## 准备和活跃度探针

Kubernetes provides two essential features called [Liveness Probes and Readiness Probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/). Essentially, Liveness/Readiness Probes will periodically perform an action (e.g. make an HTTP request, open a TCP connection, or run a command in your container) to confirm that your application is working as intended.

Kubernetes 提供了两个基本功能，称为 [Liveness Probes 和 Readiness Probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/)。本质上，Liveness/Readiness Probes 将定期执行操作（例如，发出 HTTP 请求、打开 TCP 连接或在您的容器中运行命令)以确认您的应用程序按预期工作。

Kubernetes uses **Readiness Probes** to know when a container is ready to start accepting traffic. A Pod is considered ready when all of its containers are ready. One use of this signal is to control which Pods are used as backends for Kubernetes Services (and esp. Ingress).

Kubernetes 使用 **Readiness Probes** 来了解容器何时准备好开始接受流量。当 Pod 的所有容器都准备就绪时，它就被认为是准备好了的。此信号的一个用途是控制将哪些 Pod 用作 Kubernetes 服务（尤其是 Ingress）的后端。

Kubernetes uses **Liveness Probes** to know when to restart a container. For example, Liveness Probes could catch a deadlock, where an application is running, but unable to make progress. Restarting a container in such a state can help to make the application more available despite bugs, but restarting can also lead to cascading failures (see below).

Kubernetes 使用 **Liveness Probes** 来了解何时重新启动容器。例如，Liveness Probes 可以捕获应用程序正在运行但无法取得进展的死锁。在这种状态下重新启动容器有助于使应用程序更可用，尽管存在错误，但重新启动也可能导致级联故障（见下文）。

If you attempt to deploy a change to your application that fails the Liveness/Readiness Probe, the rolling deploy will hang as it waits for all of your Pods to become Ready.

如果您尝试将更改部署到未通过 Liveness/Readiness Probe 的应用程序，滚动部署将挂起，因为它等待您的所有 Pod 都准备就绪。

### Example

###  例子

Example Readiness Probe checking the `/health` path via HTTP with default settings (interval: 10 seconds, timeout: 1 second, success threshold: 1, failure threshold: 3):

```
# part of a larger deployment/stack definition
podTemplate:
  spec:
    containers:
    - name: my-container
      # ...
      readinessProbe:
        httpGet:
          path: /health
          port: 8080
```


### DOs

### 做

- for microservices providing an HTTP endpoint (REST service etc), **always define a Readiness Probe** which checks that your application (Pod) is ready to receive traffic

- 对于提供 HTTP 端点（REST 服务等）的微服务，**始终定义一个就绪探针**，用于检查您的应用程序（Pod）是否已准备好接收流量

- make sure that your Readiness Probe covers the readiness of the actual webserver port

   - 确保您的就绪探针涵盖实际网络服务器端口的准备情况

	 - when using an "admin" or "management" port (eg 9090) for your `readinessProbe`, make sure that the endpoint only returns OK if the main HTTP port (eg 8080) is ready to accept traffic [[1\]] (https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html#id4)

   - 当为你的 `readinessProbe` 使用“admin”或“management”端口（例如 9090）时，确保只有当主 HTTP 端口（例如 8080）准备好接受流量时，端点才返回 OK [[1\]] (https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html#id4)

  - having a different port for the Readiness Probe also can lead to thread pool congestion problems on the main port which are not reflected in the health check (i.e. main server thread pool is full, but health check still answers OK)

- 为就绪探针使用不同的端口也会导致主端口上的线程池拥塞问题，这些问题不会反映在健康检查中（即主服务器线程池已满，但健康检查仍然回答 OK）

- make sure that your Readiness Probe includes database initialization/migration

  - 确保您的 Readiness Probe 包括数据库初始化/迁移

  - the simplest way to achieve this is to start the HTTP server listening only after the initialization finished (eg [Flyway](https://flywaydb.org/) DB migration etc), ie instead of changing the health check status, just don 't start the web server until the DB migration is complete [[2\]](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html#id5)

- 实现这个最简单的方法是在初始化完成后才启动 HTTP 服务器监听（例如 [Flyway](https://flywaydb.org/) DB 迁移等)，即不要更改健康检查状态，只需不要在数据库迁移完成之前不要启动 Web 服务器 [[2\]](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html#id5)

- use `httpGet` for Readiness Probes to a well-known health check endpoint (e.g. /health)

- 使用 `httpGet` 对一个众所周知的健康检查端点（例如 /health）进行就绪探测

- understand the default behavior (interval: 10s, timeout: 1s, successThreshold : 1, failureThreshold : 3) of probes 

- 了解探测的默认行为（间隔：10s，超时：1s，successThreshold：1， failureThreshold：3）

- the default values mean that a Pod will become not-ready after ~30s (3 failing health checks)

- 默认值意味着 Pod 将在大约 30 秒后变为未就绪状态（3 次健康检查失败）

- do use a different "admin" or "management" port if your tech stack (e.g. Java/Spring) allows this to separate management health and metrics from normal traffic

  - 如果您的技术堆栈（例如 Java/Spring）允许将管理健康和指标与正常流量分开，请使用不同的“管理”或“管理”端口

  - but check point 2

  - 但检查点 2

- you can use the Readiness Probe for prewarming/cache loading if needed and return 503 status code until the app container is "warm"

  - 如果需要，您可以使用 Readiness Probe 进行预热/缓存加载并返回 503 状态代码，直到应用程序容器“热”

  - check also the new `startupProbe` [introduced in 1.16](https://sysdig.com/blog/whats-new-kubernetes-1-16/)

- 还要检查新的 `startupProbe` [在 1.16 中引入](https://sysdig.com/blog/whats-new-kubernetes-1-16/)

### DON'Ts

### 不要

- do not depend on external dependencies

    - 不依赖外部依赖

   (like data stores) for your Readiness/Liveness checks as this might lead to cascading failures

   （如数据存储）用于您的就绪/活跃检查，因为这可能会导致级联故障

   - e.g. a stateful REST service with 10 pods which depends on a single Postgres database: when your probe depends on a working DB connection, all 10 pods will be "down" if the database/network has a hiccup --- this usually makes the impact worse than it should

   - 例如具有 10 个依赖于单个 Postgres 数据库的 Pod 的有状态 REST 服务：当您的探针依赖于工作数据库连接时，如果数据库/网络出现故障，所有 10 个 Pod 都将“关闭”——这通常会使影响变得更糟比它应该的

  - note that the default behavior of Spring Data is checking the DB connection [[3\]](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html#id6)

   - 注意 Spring Data 的默认行为是检查数据库连接 [[3\]](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html#id6)

  - "external" in this context can also mean other Pods of the same application, i.e. your probe should ideally not depend on the state of other Pods in the same cluster to prevent cascading failures

     - 在这种情况下，“外部”也可以指同一应用程序的其他 Pod，即理想情况下，您的探测器不应依赖于同一集群中其他 Pod 的状态，以防止级联故障

    - Your mileage may vary for apps with distributed state (e.g. in-memory caching across Pods)

- 对于具有分布式状态的应用程序，您的里程可能会有所不同（例如，跨 Pod 的内存缓存）

- do not use a Liveness Probe

    - 不要使用存活探针

   for your Pods unless you understand the consequences and why you need a Liveness Probe

   除非您了解后果以及为什么需要 Liveness Probe

  - Liveness Probe can help recover "stuck" containers, but as you are fully owning your application, things like "stuck" processes and deadlocks should not be expected --- a better alternative is to crash on purpose to recover to a known-good state

   - Liveness Probe 可以帮助恢复“卡住”的容器，但是当您完全拥有您的应用程序时，不应预期诸如“卡住”进程和死锁之类的事情 --- 更好的选择是故意崩溃以恢复到已知良好状态状态

  - a failing Liveness Probe will lead to container restarts, thus potentially making the impact of load-related errors worse: container restart will lead to downtime (at least your app's startup time, eg 30s+), thus causing more errors and giving other containers more traffic load, leading to more failing containers, and so on

   - 失败的 Liveness Probe 将导致容器重新启动，从而可能使与负载相关的错误的影响变得更糟：容器重新启动将导致停机（至少您的应用程序的启动时间，例如 30s+），从而导致更多错误并给其他容器更多流量负载，导致更多失败的容器，等等

  - Liveness Probes in combination with an external dependency are the worst situation leading to cascading failures: a single DB hiccup will restart all your containers!

- Liveness Probes 与外部依赖相结合是导致级联故障的最坏情况：单个 DB 小问题将重新启动所有容器！

- if you use Liveness Probe,

   - 如果您使用存活探针，

  don’t set the same specification for Liveness and Readiness Probe

   不要为 Liveness 和 Readiness Probe 设置相同的规范

  - you can use a Liveness Probe with the same health check, but a higher `failureThreshold` (e.g. mark as not-ready after 3 attempts and fail Liveness Probe after 10 attempts)

- 您可以使用具有相同健康检查的 Liveness Probe，但 `failureThreshold` 更高（例如，在 3 次尝试后标记为未就绪，在 10 次尝试后将 Liveness Probe 标记为失败）

- do not use "exec" probes

    - 不要使用“exec”探针

   as there are known problems with them resulting in zombie processes

   因为它们存在导致僵尸进程的已知问题

  - see [failure stories by Datadog](https://www.youtube.com/watch?v=QKI-JRs2RIE)

- 参见 [Datadog 的失败案例](https://www.youtube.com/watch?v=QKI-JRs2RIE)

### Summary

###  概括

- use Readiness Probes for your web app to decide when the Pod should receive traffic
- use Liveness Probes only when you have a use case for them
- incorrect use of Readiness/Liveness Probes can lead to reduced availability and cascading failures

- 为您的 Web 应用程序使用就绪探针来决定 Pod 何时应接收流量
- 仅当您有用例时才使用 Liveness Probes
- 错误使用就绪/活跃度探测器会导致可用性降低和级联故障

[![../galleries/twitter-sszuecs-99-do-not-need-livenessprobe.png](https://srcco.de/galleries/twitter-sszuecs-99-do-not-need-livenessprobe.png)](https://twitter.com/sszuecs/status/1175655221382529025)

)](https://twitter.com/sszuecs/status/1175655221382529025)

### Further Reading

### 进一步阅读

- [Kubernetes docs: Configure Liveness and Readiness Probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/)
- [Kubernetes Liveness and Readiness Probes Revisited: How to Avoid Shooting Yourself in the Other Foot](https://blog.colinbreck.com/kubernetes-liveness-and-readiness-probes-revisited-how-to-avoid-shooting-yourself-in-the-other-foot/)
- [NRE Labs Outage Post-Mortem](https://keepingitclassless.net/2018/12/december-4-nre-labs-outage-post-mortem/) (involves `livenessProbe`)

- [Kubernetes 文档：配置 Liveness 和 Readiness Probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/)
- [Kubernetes Liveness and Readiness Probes Revisits: How to avoid Shooting Yourself in the another Foot](https://blog.colinbreck.com/kubernetes-liveness-and-readiness-probes-revisited-how-to-avoid-shooting-自己在另一只脚上/)
- [NRE Labs 停机验尸](https://keepingitclassless.net/2018/12/december-4-nre-labs-outage-post-mortem/)（涉及`livenessProbe`)

### UPDATE 2019-09-29 #1 

### 更新 2019-09-29 #1

