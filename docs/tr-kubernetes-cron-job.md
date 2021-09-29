# Kubernetes cron jobs: a hands-on guide to optimally configured crons

# Kubernetes cron 作业：最佳配置 cron 的实践指南

Posted at 20-Feb-2021 | (about a 11 minute read)

Kubernetes is super effective on running cron jobs as well as other web  application workloads. Kubernetes cron job is a special kind of  Kubernetes job that runs on a time-based schedule. In this post, we will focus on how to run optimally configured cron jobs on Kubernetes.

Kubernetes 在运行 cron 作业以及其他 Web 应用程序工作负载方面非常有效。 Kubernetes cron 作业是一种特殊的 Kubernetes 作业，它按基于时间的计划运行。在这篇文章中，我们将重点介绍如何在 Kubernetes 上运行优化配置的 cron 作业。

![Kubernetes cron jobs, lets configure them optimally](https://geshan.com.np/images/kubernetes-cron-job/01kubernetes-cron-job.jpg)

## Table of contents

##  目录

- [What is Kubernetes?](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#what-is-kubernetes%3F)
- [What does Kubernetes do?](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#what-does-kubernetes-do%3F)
- Kubernetes Cron Job
   - [Prerequisites](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#prerequisites)
   - [Kubernetes cron job a simple example](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#kubernetes-cron-job-a-simple-example)
   - [Kubernetes cron job an optimal example](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#kubernetes-cron-job-an-optimal-example)
   - [Run Kubernetes cron jobs on the fly](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#run-kubernetes-cron-jobs-on-the-fly)
- [Conclusion](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#conclusion)

- [什么是Kubernetes？](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#what-is-kubernetes%3F)
- [Kubernetes 是做什么的？](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#what-does-kubernetes-do%3F)
- Kubernetes Cron 作业
  - [先决条件](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#prerequisites)
  - [Kubernetes cron job 一个简单的例子](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#kubernetes-cron-job-a-simple-example)
  - [Kubernetes cron job 最佳示例](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#kubernetes-cron-job-an-optimal-example)
  - [运行 Kubernetes cron 作业](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#run-kubernetes-cron-jobs-on-the-fly)
- [结论](https://geshan.com.np/blog/2021/02/kubernetes-cron-job/#conclusion)

## What is Kubernetes?

## 什么是Kubernetes？

Kubernetes has multiple definitions, we will first look at a couple of them and  then try to simplify them. To start with, the official “what is  Kubernetes” page on [Kubernetes.io](http://Kubernetes.io) says:

Kubernetes 有多个定义，我们将首先查看其中的几个，然后尝试简化它们。首先，[Kubernetes.io](http://Kubernetes.io) 上的官方“什么是 Kubernetes”页面说：

> Kubernetes is a portable, extensible, open-source platform for managing  containerized workloads and services, that facilitates both declarative  configuration and automation.

> Kubernetes 是一个可移植、可扩展的开源平台，用于管理容器化工作负载和服务，有助于声明式配置和自动化。

It further adds: “The  name Kubernetes originates from Greek, meaning helmsman or pilot. Google open-sourced the Kubernetes project in 2014.” There is also a mention  of the over 15 years of Google’s experience running production workload  at scale.

它还进一步补充说：“Kubernetes 的名字来源于希腊语，意思是舵手或飞行员。谷歌在 2014 年开源了 Kubernetes 项目。”还提到了 Google 在大规模运行生产工作负载方面超过 15 年的经验。

Wikipedia page on Kubernetes voices:

关于 Kubernetes 声音的维基百科页面：

> Kubernetes (commonly stylized as K8s) is an open-source container orchestration  system for automating computer application deployment, scaling, and  management.

> Kubernetes（通常称为 K8s）是一个开源容器编排系统，用于自动化计算机应用程序的部署、扩展和管理。

It adds: It was originally designed by  Google and is now maintained by the Cloud Native Computing Foundation. It aims to provide a "platform for automating deployment, scaling, and  operations of application containers across clusters of hosts".

它补充说：它最初由谷歌设计，现在由云原生计算基金会维护。它旨在提供一个“跨主机集群自动部署、扩展和操作应用程序容器的平台”。

## What does Kubernetes do?

## Kubernetes 是做什么的？

Ok, now let's simplify stuff. If I say this is the container era for  deploying workloads like web applications, cron jobs, and anything in  between I won’t be wrong. In container land, Docker has won the race and it is the defacto container tool.

好的，现在让我们简化一下。如果我说这是部署 Web 应用程序、cron 作业等工作负载的容器时代，我不会错的。在容器领域，Docker 赢得了比赛，它是事实上的容器工具。

Now, with that in mind. we  might start small with containers and run one or two apps in containers. This means at any given time we would be running 1-10 containers. Some  time passes, we like the idea and other advantages containers provide. We want to run a couple of more workloads/apps in containers. This  equates to having 10s of containers running and maybe in production.

现在，考虑到这一点。我们可能会从容器开始，然后在容器中运行一两个应用程序。这意味着在任何给定时间我们都会运行 1-10 个容器。一段时间过去了，我们喜欢容器提供的想法和其他优势。我们想在容器中运行更多的工作负载/应用程序。这相当于有 10 个容器在运行，并且可能在生产中。

More time passes by and more apps are containerized, at this point we have  100s of containers running. Then how do we scale these containers? How  do we make service A talk to service B? How do we handle deployments and rollbacks of 10s of applications that have 100s of containers  underlined? How do we effectively manage resources (CPU/RAM) and secrets consistently for these 100s of containers?

随着时间的流逝，更多的应用程序被容器化，此时我们有 100 个容器在运行。那么我们如何扩展这些容器呢？我们如何让服务 A 与服务 B 对话？我们如何处理具有 100 个容器下划线的 10 个应用程序的部署和回滚？我们如何有效地管理这 100 个容器的资源（CPU/RAM）和机密？

The answer to all of  the above Hows is a “container orchestrator”. Around 2015 there was a  slight competition between Kubernetes, Docker Swarm, and Apache Mesos. By mid-2017, Kubernetes comfortably won the race to become the wildly  popular de facto container orchestrator as per [Google Trends](https://trends.google.com/trends/explore?date=2015-01-01 2021-01 -31&q=kubernetes,docker swarm,apache mesos) that we can see below:

对上述所有 Hows 的答案是“容器编排器”。 2015 年左右，Kubernetes、Docker Swarm 和 Apache Mesos 之间出现了轻微的竞争。根据 [Google Trends](https://trends.google.com/trends/explore?date=2015-01-012021-01)，到 2017 年年中，Kubernetes 轻松赢得了这场竞赛，成为广受欢迎的事实上的容器编排器-31&q=kubernetes,docker swarm,apache mesos)，我们可以在下面看到：

![Kubernetes won the container orchestrator race in mid 2017](https://geshan.com.np/images/kubernetes-cron-job/02kubernetes-cron-job-popularity.jpg)

Even though Kubernetes is flexible, powerful, and ultra-popular. There are some things [Kubernetes is not](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/#what-kubernetes-is-not), it is better to know about them too. In today’s time, Kubernetes has  become a well-established platform with a thriving ecosystem around it.

尽管 Kubernetes 灵活、强大且非常受欢迎。有些东西 [Kubernetes 不是](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/#what-kubernetes-is-not)，最好也了解一下。在当今时代，Kubernetes 已成为一个完善的平台，周围有一个蓬勃发展的生态系统。

## Kubernetes Cron Job

## Kubernetes Cron 作业

I have been part of a team that used Kubernetes in Production in [2016](https://www.slideshare.net/geshan/embrace-chatops-stop-installing-deployment-software-larcon-eu-2016/54). Kubernetes is great at managing long-running workloads like web servers or queue consumers. They roughly translate to [Service](https://kubernetes.io/docs/concepts/services-networking/service/), and [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) artifact in Kubernetes lingo. In addition to being great for  long-running workloads, Kubernetes does an amazing job in managing Cron  Jobs too.

我曾在 [2016](https://www.slideshare.net/geshan/embrace-chatops-stop-installing-deployment-software-larcon-eu-2016/54) 加入一个在生产环境中使用 Kubernetes 的团队。 Kubernetes 非常擅长管理长时间运行的工作负载，例如 Web 服务器或队列使用者。它们大致翻译为 [Service](https://kubernetes.io/docs/concepts/services-networking/service/) 和 [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) Kubernetes 术语中的工件。除了非常适合长时间运行的工作负载之外，Kubernetes 在管理 Cron 作业方面也做得非常出色。

If we look at a bit of Kubernetes history, Kubernetes Cron Job was called `ScheduledJob`. In [version 1.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.5.md#action-required-before-upgrading) it was renamed to be called Cron Job. In Kubernetes, [Cron Job](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) is a special kind of a [Job](https://kubernetes.io/docs/concepts/workloads/controllers/job/) that runs on a repeating schedule. The frequency of the Kubernetes Cron Job is written in the familiar [Cron](https://crontab.guru/) format. For example `0 4 * * *` in the cron format means at 4:00 AM every morning. You can read more about the [cron schedule syntax](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#cron-schedule-syntax) if you want.

如果我们回顾一下 Kubernetes 的历史，Kubernetes Cron Job 被称为“ScheduledJob”。在 [1.5 版](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.5.md#action-required-before-upgrading) 中，它被重命名为 Cron Job。在 Kubernetes 中，[Cron Job](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) 是一种特殊的 [Job](https://kubernetes.io/docs/concepts/workloads/controllers/job/) 按重复计划运行。 Kubernetes Cron Job 的频率以熟悉的 [Cron](https://crontab.guru/) 格式编写。例如，cron 格式中的 `0 4 * * *` 表示每天早上 4:00。如果需要，您可以阅读有关 [cron schedule syntax](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#cron-schedule-syntax) 的更多信息。

How would you know if a cron job missed one or more of its schedule? It is better to know more about effective [cron job monitoring](https://geshan.com.np/blog/2019/11/how-to-efficiently-monitor-crons-with-a-simple-bash-trick/) before it slips unnoticed.

您如何知道 Cron 作业是否错过了一个或多个日程安排？最好多了解有效的 [cron作业监控](https://geshan.com.np/blog/2019/11/how-to-efficiently-monitor-crons-with-a-simple-bash-trick/) 在它被忽视之前。

### Prerequisites

### 先决条件

- You are generally aware of how Kubernetes works and schedules containers as [pods](https://kubernetes.io/docs/concepts/workloads/pods/).
- You know that Kubernetes manages objects and config in a [declarative way](https://kubernetes.io/docs/concepts/overview/working-with-objects/object-management/#declarative-object-configuration).
- The differences between a service, deployment, and [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
- You generally know what [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) is used for in Kubernetes.
- You are aware of usual Kubernetes terms like Node, Kubelet, and the likes.

- 您通常了解 Kubernetes 的工作原理并将容器调度为 [pods](https://kubernetes.io/docs/concepts/workloads/pods/)。
- 您知道 Kubernetes 以 [声明式方式](https://kubernetes.io/docs/concepts/overview/working-with-objects/object-management/#declarative-object-configuration) 管理对象和配置。
- 服务、部署和[Horizontal Pod Autoscaler]之间的区别(https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
- 您通常知道 [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) 在 Kubernetes 中的用途。
- 您了解常见的 Kubernetes 术语，例如 Node、Kubelet 等。

Next, we will look into a simple Kubernetes cron job example.

接下来，我们将研究一个简单的 Kubernetes cron 作业示例。

### Kubernetes cron job a simple example

### Kubernetes cron job 一个简单的例子

We will try a simple Kubernetes cron job example on [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/). At the time of writing Kind version 0.9.0 starts a Kubernetes cluster  of version 1.19.1. Below is our simple Kubernetes cron file that uses  node:14-alipine image to print the current date.

我们将在 [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/) 上尝试一个简单的 Kubernetes cron 作业示例。在撰写本文时，Kind 版本 0.9.0 启动了版本 1.19.1 的 Kubernetes 集群。下面是我们使用 node:14-alipine 图像打印当前日期的简单 Kubernetes cron 文件。

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: print-date
spec:
  schedule: "*/5 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: print-date
            image: node:14-alpine
            imagePullPolicy: IfNotPresent
            args:
            - -e
            - "console.log(new Date().toString());"
          restartPolicy: OnFailure
```

Let’s analyze this simple, not so well configure Cronjob.yaml file in detail:

下面我们来详细分析一下这个简单的，不太好配置的 Cronjob.yaml 文件：

1. We are using the `batch/v1beta1` API version of Kubernetes API
2. This is a type of `CronJob` Kubernetes resource/workload
3. We have named the cron job `print-date`
4. The Kubernetes cron job is scheduled to execute every 5 minutes -- `*/5 * * * *`
5. We are using the `node:14-alpine` image which will be taken from docker hub by default
6. `IfNotPresent` image pull policy is the default one. It causes the [kubelet](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/) to pull an image if it does not already exist. 

1. 我们使用的是 Kubernetes API 的 `batch/v1beta1` API 版本
2. 这是一种`CronJob` Kubernetes 资源/工作负载
3. 我们将 cron 作业命名为 `print-date`
4. Kubernetes cron 作业计划每 5 分钟执行一次 -- `*/5 * * * *`
5. 我们使用的是`node:14-alpine` 镜像，默认情况下将从 docker hub 获取
6. `IfNotPresent` 镜像拉取策略是默认的。如果图像不存在，它会导致 [kubelet](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/) 拉取图像。

7. Then we pass in `-e` for eval`and`console.log`to print the current date as string. As the command for the node container is`node` this will print the current date and time.
8. The container will be restarted on failure as per the above-defined [restart policy](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy).

7. 然后我们为eval`和`console.log`传入`-e`以将当前日期打印为字符串。由于节点容器的命令是`node`，这将打印当前日期和时间。
8. 容器将根据上面定义的[重启策略](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy)在失败时重启。

Now, we will schedule this cron job on a local [kind](https://kind.sigs.k8s.io/) Kubernetes cluster to try it out. The kind version I am using at the  time of writing this blog post is 0.9.0 which has installed Kubernetes  version 1.19.1.

现在，我们将在本地 [kind](https://kind.sigs.k8s.io/) Kubernetes 集群上安排此 cron 作业以进行尝试。我在写这篇博文时使用的那种版本是 0.9.0，它已经安装了 Kubernetes 1.19.1 版。

If we save the above file as `cronjob.yaml` we can add it to Kubernetes with the following command:

如果我们将上述文件保存为 `cronjob.yaml`，我们可以使用以下命令将其添加到 Kubernetes：

```bash
kubectl apply -f cronjob.yaml
```

After the command runs successfully we will see something like:

命令成功运行后，我们将看到如下内容：

```bash
cronjob.batch/print-date created
```

To check if the cron job is created successfully we can execute the following:

要检查 cron 作业是否成功创建，我们可以执行以下操作：

```bash
kubectl get cronjob
```

If all is good, it will print out something as follows:

如果一切正常，它会打印出如下内容：

```bash
NAME         SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE
print-date   */5 * * * *                    False          0                    <none>          17s
```

After around 5 minutes if we try `kubectl get po | grep print-date` we should see that the cron has run once, like below:

大约 5 分钟后，如果我们尝试 `kubectl get po | grep print-date` 我们应该看到 cron 已经运行了一次，如下所示：

```bash
print-date-1613818500-88ln6   0/1     Completed   0          97s
```

To see the logs of the cron job that has run we should execute `kubectl logs print-date-1613818500-88ln6` where `print-date-1613818500-88ln6` is the pod name will be different in your case. It will show us something as follows:

要查看已运行的 cron 作业的日志，我们应该执行 `kubectl logs print-date-1613818500-88ln6` 其中 `print-date-1613818500-88ln6` 是 pod 名称，在您的情况下会有所不同。它将向我们展示如下内容：

```bash
Sat Feb 20 2021 10:55:03 GMT+0000 (Coordinated Universal Time)
```

Let’s recap the commands below:

让我们回顾一下下面的命令：

![Kubernetes cron job a simple example - not optimally configured](https://geshan.com.np/images/kubernetes-cron-job/03kubernetes-cron-job-simple.jpg)

In the next part, we will look at configuring the Kubernetes Cron Job optimally.

在下一部分中，我们将研究如何以最佳方式配置 Kubernetes Cron 作业。

### Kubernetes cron job an optimal example

### Kubernetes cron 作业的最佳示例

In the above simple example, let’s scrutinize some things:

在上面的简单示例中，让我们仔细检查一些事情：

1. What if there is an error in the command, will Kubernetes try to schedule the cron job pod many times?
2. How can we clean up the pods that have completed the job?
3. What if our cron job has not finished and it is time to run the next one. We just want to skip the next run as the current job is not finished.
4. We want to temporarily stop the cron job for the time being.
5. We want to see logs of some older cron job runs even if they have failed or succeeded.

1. 如果命令有错误，Kubernetes会不会多次尝试调度cron job pod？
2. 如何清理已经完成工作的pod？

3. 如果我们的 cron 作业还没有完成，是时候运行下一个了怎么办。我们只想跳过下一次运行，因为当前作业尚未完成。
4. 我们想暂时停止cron作业。
5. 我们希望查看一些旧的 cron 作业运行的日志，即使它们失败或成功。

The answers to above questions and more lies in the cron job configuration below:

以上问题和更多问题的答案在于下面的 cron 作业配置：

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: print-date
spec:
  schedule: "*/5 * * * *"
  jobTemplate:
    spec:
      backoffLimit: 5
      ttlSecondsAfterFinished: 100
      template:
        spec:
          containers:
          - name: print-date
            image: node:14-alpine
            imagePullPolicy: IfNotPresent
            args:
            - -e
            - "console.log(new Date().toString());"
          restartPolicy: OnFailure
      parallelism: 1
      completions: 1
  concurrencyPolicy: "Forbid"
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 5
```

Let’s analyze some of the new configurations we have added and what do they do:

让我们分析一下我们添加的一些新配置以及它们的作用：

1. In this definition `backoffLimit` is used to specify the number of retries before marking the job as [failed](https://kubernetes.io/docs/concepts/workloads/controllers/job/#pod-backoff-failure-policy). For example, if the container doesn’t start or the command has an  error, we are specifying it should retry 5 times before backing off  (marking the job as a failed one).
2. To lessen the pressure on  Kubernetes, we can specify TTL seconds after finished. Where the TTL  controller cleans up the job and deletes the job in a [cascading manner](https://kubernetes.io/docs/concepts/workloads/controllers/job/#ttl-mechanism-for-finished-jobs)
3. The `parallelism` and `completions` are by default 1, It can be used to have only 1 pod running in [parallel](https://kubernetes.io/docs/concepts/workloads/controllers/job/#controlling-parallelism). 

1. 在这个定义中`backoffLimit`用于指定在将作业标记为失败之前的重试次数。例如，如果容器没有启动或命令有错误，我们指定它应该在回退之前重试 5 次（将作业标记为失败）。
2. 为了减轻Kubernetes的压力，我们可以指定完成后的TTL秒。 TTL控制器以级联方式清理作业并删除作业

3. `parallelism` 和 `completions` 默认为 1，可用于在 [parallel](https://kubernetes.io/docs/concepts/workloads/controllers/job/#controlling) 中只运行 1 个 pod -并行性)。

4. Use of `concurrencyPolicy` is very handy if you want to skip the next run if the current cron job pod is still active. Setting it to `Forbid` can enable this. If your job demands that on the next run the current run should be canceled, it can be set to replace [Concurrency Policy](https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/#concurrency-policy).
5. Lastly, we set job history for both success and failure. We do this so that  those pods are not cleaned up for a certain limit and we can check the  logs if we need to.

4. 如果您想在当前 cron 作业 pod 仍处于活动状态的情况下跳过下一次运行，则使用 `concurrencyPolicy` 非常方便。将其设置为“禁止”可以启用此功能。如果您的作业要求在下次运行时取消当前运行，则可以将其设置为替换 [并发策略](https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs /# 并发策略)。
5. 最后，我们为成功和失败设置工作历史。我们这样做是为了在一定限度内不会清理这些 pod，如果需要，我们可以检查日志。

Below is a screenshot of reapplying the new definition and checking logs from a run from the new configs we added:

下面是重新应用新定义并检查我们添加的新配置的运行日志的屏幕截图：

![Kubernetes cron job a better example - optimally configured](https://geshan.com.np/images/kubernetes-cron-job/04kubernetes-cron-job-better.jpg)

A new command in the above image is:

上图中的新命令是：

```bash
kubectl get jobs --watch
```

It gets jobs and watches it for any changes. As you can see in the screenshot it detected the change when running job `print-date-1613820600` every second for 2-3 seconds.

它获得工作并监视它是否有任何变化。正如您在屏幕截图中看到的那样，它在每秒运行作业“print-date-1613820600”2-3 秒时检测到更改。

### Run Kubernetes cron jobs on the fly

### 即时运行 Kubernetes cron 作业

Protip: You can force run a Kubernetes Cron Job (outside of the schedule) with a command like below:

提示：您可以使用如下命令强制运行 Kubernetes Cron 作业（在计划之外）：

```bash
kubectl create job print-date-try01 --from=cronjob/print-date
```

It is ultra convinient for testing cron jobs as we don't need to wait for the scheduled run.

测试 cron 作业非常方便，因为我们不需要等待预定的运行。

We are asking Kubernetes to create a job with the name `print-date-try01. The name has to be unique. If you run it the second time use`try02`. We are telling Kubernetes to create the job from our cron job which is`cronjob/print-date`.

我们要求 Kubernetes 创建一个名为“print-date-try01”的作业。名称必须是唯一的。如果您第二次运行它，请使用`try02`。我们告诉 Kubernetes 从我们的 cron 作业创建作业，它是 `cronjob/print-date`。

We can see an example of the above command as follows:

我们可以看到上述命令的示例如下：

![Kubernetes cron job a demo for create job which is very useful for testing cron jobs](https://geshan.com.np/images/kubernetes-cron-job/05kubernetes-cron-job-create-job.jpg)

As seen above the cron job even though scheduled for every 5th minute ran at `11:35:54` and `11:36:35` which is outside of its regular schedule. It was possible because we  force ran the cron job on a need basis than waiting for the schedule. This command is very handy when testing Kubernetes cron jobs that are  scheduled to run say every hour or every day.

如上所示，即使计划每 5 分钟执行一次 cron 作业，它也会在“11:35:54”和“11:36:35”运行，这超出了其常规时间表。这是可能的，因为我们根据需要强制运行 cron 作业，而不是等待时间表。在测试计划每小时或每天运行的 Kubernetes cron 作业时，此命令非常方便。

## Conclusion

##  结论

Kubernetes cron jobs are very useful as we have seen. In addition to being great  at handling long-running workloads, Kubernetes also does an amazing job  of executing jobs and cron jobs alike.

正如我们所见，Kubernetes cron 作业非常有用。除了擅长处理长时间运行的工作负载之外，Kubernetes 在执行作业和 cron 作业方面也做得非常出色。

> Configure your Kubernetes cron jobs optimally to run the cron jobs as you expect on a Kubernetes cluster.

> 以最佳方式配置您的 Kubernetes cron 作业，以按照您在 Kubernetes 集群上的预期运行 cron 作业。

Even modern applications have tasks that need to be done with Cron jobs and  Kubernetes cron jobs can be exploited for such tasks.

即使是现代应用程序也有需要使用 Cron 作业完成的任务，而 Kubernetes cron 作业可以用于此类任务。

Posted By Geshan Manandhar | Posted at 20-Feb-2021 Please share:

格山·马南达尔发表 |发表于 20-Feb-2021 请分享：

- [« How to use nodemon to restart your Node.js applications automatically and efficiently](https://geshan.com.np/blog/2021/02/nodemon/)
- [10 JavaScript array functions you should start using today »](https://geshan.com.np/blog/2021/03/javascript-array-functions/) 

- [« 如何使用 nodemon 自动高效地重启你的 Node.js 应用程序](https://geshan.com.np/blog/2021/02/nodemon/)
- [你今天应该开始使用的 10 个 JavaScript 数组函数 »](https://geshan.com.np/blog/2021/03/javascript-array-functions/)

