# Goodbye minikube

  # 再见minikube

 Mar 7, 2021 / [KUBERNETES](https://blog.frankel.ch/tag/kubernetes/),[MINIKUBE](https://blog.frankel.ch/tag/minikube/), [KIND](https://blog.frankel.ch/tag/kind/)
  

2021 年 3 月 7 日 / [KUBERNETES](https://blog.frankel.ch/tag/kubernetes/)、[MINIKUBE](https://blog.frankel.ch/tag/minikube/)、[KIND](https://blog.frankel.ch/tag/kind/)
  

I’ve been using [minikube](https://minikube.sigs.k8s.io/) as my local cluster since I started to learn [Kubernetes](https://kubernetes.io/). But I’ve decided to let it go in favor of [kind](https://kind.sigs.k8s.io/). Here’s the story.

自从我开始学习 [Kubernetes](https://kubernetes.io/) 以来，我一直在使用 [minikube](https://minikube.sigs.k8s.io/) 作为我的本地集群。但我决定让它去支持 [kind](https://kind.sigs.k8s.io/)。这是故事。

A couple of weeks ago, I gave my talk on Zero Downtime on Kubernetes. A  demo is included in the talk, as with most of my presentations. While  rehearsing in the morning, the demo worked, albeit slowly. Two days  before that, I had another demo that also uses Kubernetes and it was  already slow. But I didn’t take the hint.

几周前，我发表了关于 Kubernetes 零停机时间的演讲。演讲中包含一个演示，就像我的大多数演讲一样。早上排练时，演示成功了，尽管速度很慢。两天前，我做了另一个同样使用 Kubernetes 的演示，它已经很慢了。但我没有接受这个提示。

During the demo, everything was slow: the of scheduling pods, of course, but  also the running and the deletion of pods. The demo failed miserably. I  didn’t even manage to stop minikube cleanly and I had to stop the VM.

在演示过程中，一切都很缓慢：当然是调度 Pod，还有运行和删除 Pod。演示惨遭失败。我什至没有设法干净地停止 minikube，我不得不停止 VM。

To say I was disappointed is quite an understatement. That was my first  shot at this demo. I hate when demos go wrong; I hate it even more when  it works during rehearsal and it fails in front of the audience. I  apologized profusely and decided that I wouldn’t repeat the same  experience.

说我很失望是轻描淡写的。那是我第一次接触这个演示。我讨厌演示出错；我更讨厌它在排练时奏效并且在观众面前失败。我深表歉意，并决定不再重复同样的经历。

After the talk, I deleted the cluster and created it from scratch again. Like for the deleted cluster, I used the *virtualbox* driver. I also used the same configuration as before: 4 cores and 16 Gb. And yet, scheduling was slow… again.

谈话结束后，我删除了集群并重新从头开始创建。对于删除的集群，我使用了 *virtualbox* 驱动程序。我还使用了与以前相同的配置：4 核和 16 Gb。然而，调度很慢......再次。

I already had [some interest](https://www.reddit.com/r/kubernetes/comments/jua2f1/local_kubernetes_minikube_vs_microk8s/) in alternatives to minikube. This failure gave me the right incentive. I chose kind because some of the comments mention it in good terms.

我已经对 minikube 的替代品产生了 [一些兴趣](https://www.reddit.com/r/kubernetes/comments/jua2f1/local_kubernetes_minikube_vs_microk8s/)。这次失败给了我正确的激励。我选择 kind 是因为有些评论很好地提到了它。

Coming from minikube, there are a couple of differences worth mentioning. The most important one is that **kind runs in Docker**. Its name is actually an acronym for "**k\*ubernetes \*in** *d*ocker". Hence, Docker must be running prior to any kind-related operation.

来自 minikube，有几个不同之处值得一提。最重要的是**kind在Docker中运行**。它的名字实际上是“**k\*ubernetes \*in** *d*ocker”的首字母缩写。因此，Docker 必须在任何与种类相关的操作之前运行。

As a consequence, there’s no dedicated cluster IP, everything is directly on `localhost`. However, the cluster needs to be configured explicitly to map ports.

因此，没有专用的集群 IP，一切都直接在“localhost”上。但是，需要显式配置集群以映射端口。

kind.yml

种类.yml

```
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
nodes:
  - role: control-plane
    extraPortMappings:
      - containerPort: 30002
        hostPort: 30002
  - role: worker
```

| | Map container’s port `30002` to host’s port `30002` |
| ---- |--------------------------------------------------- |
| | |

| |将容器的端口 `30002` 映射到主机的端口 `30002` |
| - |
| | |

One needs to pass the configuration at creation time:

需要在创建时传递配置：

```
kind create cluster --config kind.yml
```

The cluster configuration cannot be changed. The only workaround is to  delete the cluster and create another one with the new configuration.

无法更改集群配置。唯一的解决方法是删除集群并使用新配置创建另一个集群。

Another important difference becomes visible when the image of the scheduled pod is local *i.e.* not available in a registry. With minikube, one would configure the  environment so that when one builds an image, it’s directly loaded into  the cluster’s Docker *daemon*. With kind, one needs to load images from Docker to the kind cluster.

当预定 Pod 的映像是本地 *即* 在注册表中不可用时，另一个重要的区别变得可见。使用 minikube，可以配置环境，以便在构建镜像时，它会直接加载到集群的 Docker *daemon* 中。使用 kind，需要从 Docker 加载镜像到 kind 集群。

```
kind load docker-image hazelcast/hzshop:1.0
```

I re-tested the whole demo. It works like a charm!

我重新测试了整个演示。它就像一个魅力！

There’s one remaining step in my context, create an `Ingress`. The [documentation](https://kind.sigs.k8s.io/docs/user/ingress/) is clear.

在我的上下文中还剩下一个步骤，创建一个 `Ingress`。 [文档](https://kind.sigs.k8s.io/docs/user/ingress/) 很清楚。

##   To go further:

## 更进一步：

- [kind Quick Start](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [kind Ingress](https://kind.sigs.k8s.io/docs/user/ingress/)
- [kind Initial design](https://kind.sigs.k8s.io/docs/design/initial/)

- [kind 快速入门](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [kind Ingress](https://kind.sigs.k8s.io/docs/user/ingress/)
- [kind 初始设计](https://kind.sigs.k8s.io/docs/design/initial/)

####  [Nicolas Fränkel](https://blog.frankel.ch/me) 

#### [尼古拉斯弗兰克尔](https://blog.frankel.ch/me)

Developer Advocate with 15+ years experience consulting for many different  customers, in a wide range of contexts (such as telecoms, banking,  insurances, large retail and public sector). Usually working on  Java/Java EE and Spring technologies, but with focused interests like  Rich Internet Applications, Testing, CI/CD and DevOps. Currently working for Hazelcast. Also double as a trainer and triples as a book author.

  Developer Advocate 拥有 15 年以上为众多不同客户提供咨询的经验，涉及范围广泛（例如电信、银行、保险、大型零售和公共部门）。通常从事 Java/Java EE 和 Spring 技术，但重点关注富 Internet 应用程序、测试、CI/CD 和 DevOps。目前为 Hazelcast 工作。还兼任培训师和三倍的书籍作者。

 [Read More](https://blog.frankel.ch/me) 

[阅读更多](https://blog.frankel.ch/me)

