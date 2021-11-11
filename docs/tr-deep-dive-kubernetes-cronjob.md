# Deep dive into Kubernetes CronJob

# 深入了解 Kubernetes CronJob

kubernetes cron scheduling

Kubernetes cron 调度

Published on 2021/01/05

Kubernetes **CronJob** are very useful, but can but hard to work with: parallelism, failure, timeout, etc. The [official documentation](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) and [API reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#cronjob-v1beta1-batch) are quite complete, I'll try below to give real -world examples on how to fine-tune your CronJobs and follow best practices.

Kubernetes **CronJob** 非常有用，但很难使用：并行、失败、超时等。 [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) 和 [API 参考](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#cronjob-v1beta1-batch) 已经很完整了，我会在下面尝试给出真实的-有关如何微调 CronJobs 和遵循最佳实践的世界示例。

Similar to a regular cron, a Kubernetes **CronJob** runs commands at a given time interval (cron schedule, following [cron syntax](https://en.wikipedia.org/wiki/Cron)), creating 1 (or more, see below) Kubernetes **Job** (hence **Pods**) per execution.

与常规 cron 类似，Kubernetes **CronJob** 在给定的时间间隔（cron 计划，遵循 [cron 语法](https://en.wikipedia.org/wiki/Cron)）运行命令，创建 1（或更多，见下文）Kubernetes **Job**（因此 **Pods**)每次执行。

Example definition of a CronJob:

CronJob 的示例定义：

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
name: my-cron
namespace: default
spec:
# Run every 5 minutes
schedule: */5 * * * *
concurrencyPolicy: Forbid
backoffLimit: 2
startingDeadlineSeconds: 600
successfulJobsHistoryLimit: 0
failedJobsHistoryLimit: 1
suspend: false
jobTemplate:
    spec:
      activeDeadlineSeconds: 300
      template:
        metadata:
          labels:
            app: my-cron
        spec:
          activeDeadlineSeconds: 420
          containers:
          - name: cron
            image: my-repo/cron:latest

```

Let's explain the main properties:

让我们解释一下主要属性：

- `spec.concurrencyPolicy`:

   -`spec.concurrencyPolicy`：

  - `Forbid`: if previous Job is still running, skip execution of a new Job
   - `Replace`: stop running instance of the job and start a new one
   - `Allow`: multiple instances of the job can run in parallel
- `spec.schedule`: cron schedule in [cron format](https://en.wikipedia.org/wiki/Cron)
- `spec.startingDeadlineSeconds`: optional deadline to start the job if failed to start on schedule (ie. job already running with conncurrency policy set to `Forbid`)
- `spec.suspend`: set as `true` to disable cron (does not affect already running jobs)

- `Forbid`：如果之前的 Job 仍在运行，则跳过新 Job 的执行
  - `Replace`：停止正在运行的作业实例并开始一个新的
  - `Allow`：作业的多个实例可以并行运行
- `spec.schedule`：[cron 格式](https://en.wikipedia.org/wiki/Cron) 中的 cron 时间表
- `spec.startingDeadlineSeconds`：如果未能按计划启动，则启动作业的可选截止日期（即作业已经运行，并发策略设置为“禁止”）
- `spec.suspend`：设置为 `true` 以禁用 cron（不影响已经运行的作业）

As you can see, `spec.jobTemplate.spec` is actually the spec of a Job, itself containing the spec of a Pod in it's `template.spec` property. That means that we can use all capabilities of Jobs and Pods when using CronJobs.

正如你所看到的，`spec.jobTemplate.spec` 实际上是一个 Job 的规范，它本身在它的 `template.spec` 属性中包含了一个 Pod 的规范。这意味着我们在使用 CronJobs 时可以使用 Jobs 和 Pods 的所有功能。

Among others, the following ones may especially be useful:

其中，以下可能特别有用：

- `spec.jobTemplate.spec.activeDeadlineSeconds`:

   -`spec.jobTemplate.spec.activeDeadlineSeconds`：

  - Maximum duration of the Job before Kubernetes tries to stop it
   - This deadline is relative to Job creation, Pod may take some time to start, for instance when pulling image is necessary
   - Set this value only if you know what you are doing
- `spec.jobTemplate.spec.backoffLimit`:

   - Kubernetes 尝试停止作业之前的最长作业持续时间
  - 这个deadline是和Job创建相关的，Pod可能需要一些时间来启动，比如需要拉镜像的时候
  - 仅当您知道自己在做什么时才设置此值
  -`spec.jobTemplate.spec.backoffLimit`：

  - Maximum number of retries before marking this Job as failed when Pod fails (exit code > 0) or Pod deadline exceeded
   - Defaults to 6, set to 0 for no retry possible
- `spec.jobTemplate.spec.template.spec.activeDeadlineSeconds`:

   - 当 Pod 失败（退出代码 > 0）或超过 Pod 截止日期时，在将此作业标记为失败之前的最大重试次数
  - 默认为 6，设置为 0 表示无法重试
- `spec.jobTemplate.spec.template.spec.activeDeadlineSeconds`：

  - Maximum duration of the Pod before Kubernetes tries to stop it
   - If Job `backoffLimit` is > 0, a new Pod will be created until reached

- Kubernetes 尝试停止之前 Pod 的最长持续时间
  - 如果 Job `backoffLimit` > 0，将创建一个新的 Pod，直到达到

## CronJobs explained visually

## CronJobs 直观解释

Usage of `activeDeadlineSeconds` on Pod definition:

在 Pod 定义中使用 `activeDeadlineSeconds`：

![](http://michael.bouvy.net/userfiles/images/kubernetes-cronjob/xKube-Cron-05,281,29.png.pagespeed.ic.jzYHUAcPqx.png)

CronJob shown below is scheduled to run every 5 minutes, with a Job active deadline set to 540 seconds, and an execution time of the job greater than 600 seconds:

下面显示的 CronJob 计划每 5 分钟运行一次，作业活动截止时间设置为 540 秒，作业的执行时间大于 600 秒：

![](http://michael.bouvy.net/userfiles/images/kubernetes-cronjob/xKube-Cron-01.png.pagespeed.ic.qWsC2waLYt.png)

CronJob shown below is scheduled to run every 5 minutes, with a Job active deadline set to 900 seconds, and an execution time of about 600 seconds:

下面显示的 CronJob 计划每 5 分钟运行一次，作业活动截止时间设置为 900 秒，执行时间约为 600 秒：

![](http://michael.bouvy.net/userfiles/images/kubernetes-cronjob/xKube-Cron-02,281,29.png.pagespeed.ic.eRgraJXaEC.png)

CronJob shown below is scheduled to run every 5 minutes, with a Pod active deadline set to 420 seconds, and an execution time greater than 420 seconds:

下面显示的 CronJob 计划每 5 分钟运行一次，Pod 活动截止时间设置为 420 秒，执行时间大于 420 秒：

![](http://michael.bouvy.net/userfiles/images/kubernetes-cronjob/xKube-Cron-03.png.pagespeed.ic.fAvWmA-TCZ.png)

CronJob shown below is scheduled to run every 5 minutes, with a Pod active deadline set to 900 seconds, an execution time of about 600 seconds, and a failing execution:

下面显示的 CronJob 计划每 5 分钟运行一次，Pod 活动截止时间设置为 900 秒，执行时间约为 600 秒，执行失败：

![](http://michael.bouvy.net/userfiles/images/kubernetes-cronjob/xKube-Cron-04.png.pagespeed.ic.bL9DJ4e8Pl.png)

## Best practices

## 最佳实践

### Avoid running all jobs at the same time 

### 避免同时运行所有作业

Most of the time, Cron jobs are scheduled to run every minute, every 5 minutes, or hour, which translate respectively to `* * * * *`, `*/5 * * * *` and `0 * * * *` in cron syntax.

大多数情况下，Cron 作业被安排为每分钟、每 5 分钟或每小时运行一次，它们分别转换为 `* * * * *`、`*/5 * * * *` 和 `0 * * * *`在 cron 语法中。

As the number of CronJobs begins to grow on your cluster, you'll notice resources usage peaks every hour (all cron expressions above match), every 5 mins (2 expressions match), etc.

随着集群上 CronJobs 的数量开始增长，您会注意到资源使用高峰每小时（上面的所有 cron 表达式都匹配）、每 5 分钟（2 个表达式匹配）等等。

This may cause unnecessary cluster autoscaling, high network throughput affecting other workloads, race conditions, etc.

这可能会导致不必要的集群自动缩放、影响其他工作负载的高网络吞吐量、竞争条件等。

Prefer alternative execution schedules, for instance:

首选替代执行计划，例如：

- Every 5 minutes:`4,9,14,19,24,29,34,39,44,49,54,49 * * * *`
- Every 6 minutes:`*/6 * * * *`

- 每 5 分钟：`4,9,14,19,24,29,34,39,44,49,54,49 * * * *`
- 每 6 分钟：`*/6 * * * *`

### Use resources efficiently

### 有效利用资源

As CronJobs run through Pods, it is strongly recommended to define [resources requests and limits](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#requests-and-limits).

由于 CronJobs 通过 Pod 运行，因此强烈建议定义 [资源请求和限制](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#requests-and-limits)。

Although most cron workloads will run on a single thread (ie. PHP script), thus using at most 1 core, you may want to define a lower CPU limit, so you do not need to request too much resources.

尽管大多数 cron 工作负载将在单个线程（即 PHP 脚本）上运行，因此最多使用 1 个核心，但您可能希望定义一个较低的 CPU 限制，因此您不需要请求太多资源。

This is especially true when a limit of (for instance) 0.25 CPU instead of 1 would not have a big impact on execution time, or if execution time is not a big issue. That would allow only requesting 0.25 CPU, therefore avoiding potential need of cluster autoscaling when multiple cron jobs run at the same time.

当（例如）0.25 CPU 而不是 1 的限制不会对执行时间产生很大影响，或者执行时间不是大问题时，尤其如此。这将只允许请求 0.25 个 CPU，从而避免在多个 cron 作业同时运行时潜在的集群自动缩放需求。

You should whenever possible request and limit resources to as low as possible.

您应该尽可能地请求和限制资源。

### Handle termination signal

### 处理终止信号

As any other Kubernetes workload, Pods created by CronJobs may be interrupted at any time, receiving a TERM signal on PID 1 of the container (if no preStop hook is defined).

与任何其他 Kubernetes 工作负载一样，CronJobs 创建的 Pod 可能随时中断，在容器的 PID 1 上接收 TERM 信号（如果未定义 preStop 挂钩）。

This signal should be handled properly by the running task, so it is not "hard-killed" after the termination grace period (30s by default).

这个信号应该由正在运行的任务正确处理，因此在终止宽限期（默认为 30 秒）之后它不会被“硬杀死”。

### Atomic non-blocking cron

### 原子非阻塞cron

Many applications and frameworks have an embedded cron scheduler, which will be ran using a single cron script.

许多应用程序和框架都有一个嵌入式 cron 调度程序，它将使用单个 cron 脚本运行。

This not optimal as a long running cron task may block others. You should better split tasks in groups or even run tasks individually in separate Kubernetes CronJobs.

这不是最佳选择，因为长时间运行的 cron 任务可能会阻塞其他任务。您应该更好地将任务分成组，甚至在单独的 Kubernetes CronJobs 中单独运行任务。

## Pitfalls

## 陷阱

### Timezone

###  时区

When scheduling CronJobs at specific hours (ie. `0 5 * * *` to run everyday at 5:00), carefully check your Nodes timezone (which will be used by Kubernetes CronJob Controller to schedule Jobs).

在特定时间安排 CronJobs 时（即`0 5 * * *` 每天在 5:00 运行），请仔细检查您的节点时区（Kubernetes CronJob 控制器将使用它来安排作业）。

Also be careful of timezone changes (DST) twice a year, as jobs may be skipped or ran twice. 

还要注意一年两次的时区更改 (DST)，因为作业可能会被跳过或运行两次。

