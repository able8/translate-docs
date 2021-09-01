# Bad Pods: Kubernetes Pod Privilege Escalation

# 坏 Pod：Kubernetes Pod 权限提升

on Jan 19, 2021 5:26:38 AM

What are the risks associated with overly permissive pod creation in Kubernetes? The answer varies based on which of the host’s namespaces and security contexts are allowed. In this post, I will describe eight insecure pod configurations and the corresponding methods to perform privilege escalation. This article and the accompanying [repository](https://github.com/BishopFox/badPods) were created to help penetration testers and administrators better understand common misconfiguration scenarios.

在 Kubernetes 中过度宽松地创建 Pod 会带来哪些风险？答案根据允许的主机命名空间和安全上下文而有所不同。在这篇文章中，我将描述八种不安全的 pod 配置以及相应的提权方法。本文和随附的 [repository](https://github.com/BishopFox/badPods) 旨在帮助渗透测试人员和管理员更好地了解常见的错误配置场景。

If you are an administrator, I hope that this post gives you the confidence to apply restrictive controls around pod creation by default. I also hope it helps you consider isolating any pods that need access to the host’s resources to a namespace that is only accessible to administrators using the principle of least privilege.

如果您是管理员，我希望这篇文章能让您有信心在默认情况下对 pod 创建应用限制性控制。我也希望它可以帮助您考虑将需要访问主机资源的任何 pod 隔离到一个名称空间，该名称空间只能使用最小权限原则由管理员访问。

If you are a penetration tester, I hope this post provides you with some ideas on how to demonstrate the impact of an overly permissive pod security policy. And I hope that the repository gives you some easy-to-use manifests and actionable steps to achieve those goals.

如果您是一名渗透测试员，我希望这篇文章为您提供一些关于如何证明过度宽松的 pod 安全策略的影响的想法。我希望存储库为您提供一些易于使用的清单和可操作的步骤来实现这些目标。

**Executive Summary:**

**执行摘要：**

One of the foundations of information security is the "principal of least privilege." This means that every user, system process, or application needs to operate using the least set of privileges required to do a task. When privileges are configured where they greatly exceed what is required, attackers can take advantage of these situations to access sensitive data, compromise systems, or escalate those privileges to conduct lateral movement in a network.

信息安全的基础之一是“最小特权原则”。这意味着每个用户、系统进程或应用程序都需要使用执行任务所需的最少权限集进行操作。当权限被配置为大大超过所需的权限时，攻击者可以利用这些情况来访问敏感数据、破坏系统或升级这些权限以在网络中进行横向移动。

Kubernetes and other new "DevOps" technologies are complex to implement properly and are often deployed misconfigured or configured with more permissions than necessary. The lesson, as we have demonstrated from our "Bad Pods" research, is that if you are using Kubernetes in your infrastructure, you need to find out from your development team how they are configuring and hardening this environment.

Kubernetes 和其他新的“DevOps”技术很难正确实施，并且经常被错误配置或配置了不必要的权限。正如我们在“Bad Pods”研究中所展示的那样，教训是，如果您在基础设施中使用 Kubernetes，您需要从您的开发团队那里了解他们如何配置和强化此环境。

## HARDENING PODS: HOW RISKY CAN A SINGLE ATTRIBUTE BE?

## 硬化Pods：单一属性的风险有多大？

When it comes to Kubernetes security best practices, every checklist worth its salt mentions that you want to use the principle of least privilege when provisioning pods. But how can we enforce granular security controls and how do we evaluate the risk of each attribute?

当谈到 Kubernetes 安全最佳实践时，每个值得一提的清单都提到您希望在配置 pod 时使用最小权限原则。但是，我们如何实施精细的安全控制以及如何评估每个属性的风险？

A Kubernetes administrator can enforce the principle of least privilege using admission controllers. For example, there's a built-in Kubernetes controller called [PodSecurityPolicy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) and also a popular third-party admission controller called [OPA Gatekeeper](https://github.com/open-policy-agent/gatekeeper). Admission controllers allow you to deny a pod entry into the cluster if it has more permissions than the policy allows.

Kubernetes 管理员可以使用准入控制器强制执行最小权限原则。例如，有一个名为 [PodSecurityPolicy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) 的内置 Kubernetes 控制器以及一个名为 OPA Gatekeeper 的流行的第三方准入控制器。如果 Pod 具有比策略允许的更多权限，则准入控制器允许您拒绝该 Pod 进入集群。

However, even though the controls exist to define and enforce policy, the real-world security implications of allowing each specific attribute is not always understood, and quite often, pod creation is not as locked down as it needs to be.

然而，即使存在定义和执行策略的控制措施，也并不总是理解允许每个特定属性的现实世界的安全含义，而且通常情况下，Pod 的创建并没有像它需要的那样锁定。

As a penetration tester, you might find yourself with access to create pods on a cluster where there is no policy enforcement. This is what I like to refer to as “easy mode.” Use [this manifest](https://raesene.github.io/blog/2019/04/01/The-most-pointless-kubernetes-command-ever/) from Rory McCune ( [@raesene](https:// twitter.com/raesene)), [this command](https://twitter.com/mauilion/status/1129468485480751104) from Duffie Cooley ( [@mauilion](https://twitter.com/mauilion)), or the [node-shell](https://github.com/kvaps/kubectl-node-shell) krew plugin and you will have fully interactive privileged code execution on the underlying host. It doesn’t get easier than that!

作为渗透测试人员，您可能会发现自己有权在没有策略执行的集群上创建 pod。这就是我喜欢所说的“简单模式”。使用 Rory McCune ( [@raesene](https:// twitter.com/raesene))、[此命令](https://twitter.com/mauilion/status/1129468485480751104) 来自 Duffie Cooley ([@mauilion](https://twitter.com/mauilion))，或[node-shell](https://github.com/kvaps/kubectl-node-shell) krew 插件，您将在底层主机上执行完全交互式的特权代码。没有比这更容易的了！

But what if you can create a pod with just? What can you do in each case? Let’s take a look!

但是如果你可以创建一个 pod 呢？在每种情况下你能做什么？让我们来看看！

![Kubernetes bad pods](https://labs.bishopfox.com/hs-fs/hubfs/Kubernetes%20-%20PAD%20PODS%20-%20300%20PPI%20-%20Option%2001.png?width=700&name=Kubernetes%20-%20PAD%20PODS%20-%20300%20PPI%20-%20Option%2001.png)



## BAD PODS - ATTRIBUTES AND THEIR WORST-CASE SECURITY IMPACT

## 坏 Pod - 属性及其最坏情况下的安全影响

The pods below are loosely ordered from highest to lowest security impact. Note that the generic attack paths that could affect any Kubernetes pod (eg, checking to see if the pod can access the cloud provider's metadata service or identifying misconfigured Kubernetes RBAC) are covered in [Bad Pod #8: Nothing allowed](http:/ /labs.bishopfox.com#pod8).

下面的 pod 是从安全影响从最高到最低的松散排序。请注意，可能影响任何 Kubernetes pod 的通用攻击路径（例如，检查 pod 是否可以访问云提供商的元数据服务或识别错误配置的 Kubernetes RBAC）。

## THE BAD PODS LINEUP

## 糟糕的Pods阵容

**Pods**

[Bad Pod #1: Everything allowed](http://labs.bishopfox.com#Pod1)

[Bad Pod #2: Privileged and hostPid](http://labs.bishopfox.com#Pod2)

[Bad Pod #3: Privileged only](http://labs.bishopfox.com#Pod3)

[Bad Pod #4: hostPath only](http://labs.bishopfox.com#Pod4) 

[Bad Pod #5: hostPid only](http://labs.bishopfox.com#Pod5)

[Bad Pod #6: hostNetwork only](http://labs.bishopfox.com#pod6)

[Bad Pod #7: hostIPC only](http://labs.bishopfox.com#pod7)

[Bad Pod #8: Nothing allowed](http://labs.bishopfox.com#pod8)

[坏 Pod #1：一切都允许](http://labs.bishopfox.com#Pod1)

[坏 Pod #2：特权和 hostPid](http://labs.bishopfox.com#Pod2)

[坏 Pod #3：仅限特权](http://labs.bishopfox.com#Pod3)

[坏 Pod #4：仅主机路径](http://labs.bishopfox.com#Pod4)

[坏 Pod #5：仅 hostPid](http://labs.bishopfox.com#Pod5)

[坏 Pod #6：仅主机网络](http://labs.bishopfox.com#pod6)

[坏 Pod #7：仅 hostIPC](http://labs.bishopfox.com#pod7)

[坏 Pod #8：不允许的东西](http://labs.bishopfox.com#pod8)

## BAD POD \#1: EVERYTHING ALLOWED

## 坏 POD \#1：一切都允许

![BAD POD #1: EVERYTHING ALLOWED](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2001%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2001%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

### 可能发生的最坏情况是什么？

Multiple paths to full cluster compromise

完整集群妥协的多种途径

### How?

The pod you create mounts the host’s filesystem to the pod. You'll have the best luck if you can schedule your pod on a control-plane node using the [nodeName](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/# create-a-pod-that-gets-scheduled-to-specific-node) selector in your manifest. You then `exec` into your pod and `chroot` to the directory where you mounted the host’s filesystem. You now have root on the node running your pod.

您创建的 pod 将主机的文件系统挂载到 pod。如果您可以使用 [nodeName](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/#) 在控制平面节点上安排 pod，那么您将获得最好的运气create-a-pod-that-gets-scheduled-to-specific-node) 选择器在您的清单中。然后，您将“exec”进入您的 pod，并将“chroot”进入您挂载主机文件系统的目录。您现在在运行 pod 的节点上拥有 root 权限。

- **Read secrets from `etcd`** — If you can run your pod on a control-plane node using the `nodeName` selector in the pod spec, you might have easy access to the `etcd` database, which contains the configuration for the cluster, including all secrets.
- **Hunt for privileged service account tokens**— Even if you can only schedule your pod on the worker node, you can also access any secret mounted within any pod on the node you are on. In a production cluster, even on a worker node, there is usually at least one pod that has a mounted token that is bound to a service account that is bound to a clusterrolebinding, which gives you access to do things like create pods or view secrets in all namespaces

- **从 `etcd` 读取机密** - 如果您可以使用 pod 规范中的 `nodeName` 选择器在控制平面节点上运行您的 pod，您可能可以轻松访问包含集群的配置，包括所有秘密。
- **寻找特权服务帐户令牌**——即使您只能在工作节点上安排 pod，您也可以访问安装在您所在节点上的任何 pod 中的任何机密。在生产集群中，即使在工作节点上，通常也至少有一个 pod 具有挂载令牌，该令牌绑定到绑定到 clusterrolebinding 的服务帐户，这使您可以执行创建 pod 或查看机密等操作在所有命名空间中

Some additional privilege escalation patterns are outlined in the [README](https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed) document linked below and also in [Bad Pod #4: hostPath]( https://labs.bishopfox.com/tech-blog/bad-pods-kubernetes-pod-privilege-escalation?hs_preview=qrRGstOq-40671699438#Pod4).

下面链接的 [README](https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed) 文档以及 [Bad Pod #4: hostPath](https://labs.bishopfox.com/tech-blog/bad-pods-kubernetes-pod-privilege-escalation?hs_preview=qrRGstOq-40671699438#Pod4)。

### Usage and exploitation examples

### 使用和利用示例

[https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed](https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed)

[https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed](https://github.com/BishopFox/badPods/tree/main/manifests/everything-allowed)

### References and further reading

### 参考和进一步阅读

- [The Most Pointless Kubernetes Command Ever](https://raesene.github.io/blog/2019/04/01/The-most-pointless-kubernetes-command-ever/)
- [Secure Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [Deep Dive into Real-World Kubernetes Threats](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [Compromising Kubernetes Cluster by Exploiting RBAC Permissions](https://www.youtube.com/watch?v=1LMo0CftVC4) ( [slides](https://published-prd.lanyonevents.com/published/rsaus20/sessionsFiles/ 18100/2020_USA20_DSO-W01_01_Compromising%20Kubernetes%20Cluster%20by%20Exploiting%20RBAC%20Permissions.pdf))
- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)

- [有史以来最无意义的 Kubernetes 命令](https://raesene.github.io/blog/2019/04/01/The-most-pointless-kubernetes-command-ever/)
- [安全 Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [深入了解真实世界的 Kubernetes 威胁](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [通过利用 RBAC 权限破坏 Kubernetes 集群](https://www.youtube.com/watch?v=1LMo0CftVC4) ([幻灯片](https://published-prd.lanyonevents.com/published/rsaus20/sessionsFiles/ 18100/2020_USA20_DSO-W01_01_Compromising%20Kubernetes%20Cluster%20by%20Exploiting%20RBAC%20Permissions.pdf))
- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)

## BAD POD \#2: PRIVILEGED AND HOSTPID

## BAD POD \#2：特权和主机PID

![BAD POD #2: PRIVILEGED AND HOSTPID](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2002%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2002%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

### 可能发生的最坏情况是什么？

Multiple paths to full cluster compromise

完整集群妥协的多种途径

### How?

In this scenario, the only thing that changes from the everything-allowed pod is how you gain root access to the host. Rather than `chroot`ing to the host’s filesystem, you can use `nsenter` to get a root shell on the node running your pod.

在这种情况下，与所有内容允许的 pod 相比，唯一改变的是您如何获得对主机的 root 访问权限。您可以使用 nsenter 在运行 pod 的节点上获取 root shell，而不是通过 `chroot` 访问主机的文件系统。

**Why does it work?**

**为什么有效？**

- **Privileged** — The `privileged: true` container-level security context breaks down almost all the walls that containers are supposed to provide; however, the PID namespace is one of the few walls that stands. Without `hostPID`, `nsenter` would only work to enter the namespaces of a process running within the container. For more examples on what you can do if you only have `privileged: true`, refer to the next example [Bad Pod #3: Privileged Only](http://labs.bishopfox.com#Pod3).

- **Privileged** — `privileged: true` 容器级安全上下文几乎打破了容器应该提供的所有壁垒；然而，PID 命名空间是少数几堵墙之一。如果没有 `hostPID`，`nsenter` 只能进入容器内运行的进程的命名空间。有关如果您只有 `privileged: true` 可以做什么的更多示例，请参阅下一个示例 [Bad Pod #3: Privileged Only](http://labs.bishopfox.com#Pod3)。

- **Privileged + `hostPID`** — When both `hostPID: true` and `privileged: true` are set, the pod can see all of the processes on the host, and you can enter the `init` system (PID 1) on the host. From there, you can execute your shell on the node.


- **Privileged + `hostPID`** — 当 `hostPID: true` 和 `privileged: true` 都设置时，pod 可以看到主机上的所有进程，您可以进入 `init` 系统（PID 1) 在主机上。从那里，您可以在节点上执行您的 shell。


Once you are root on the host, the privilege escalation paths are all the same as described in [Bad Pod # 1: Everything-allowed](http://labs.bishopfox.com#Pod1).

一旦您成为主机上的 root，权限提升路径与 [Bad Pod #1: Everything-allowed](http://labs.bishopfox.com#Pod1) 中描述的完全相同。

### Usage and exploitation examples 
### 使用和利用示例
[https://github.com/BishopFox/badPods/tree/main/manifests/priv-and-hostpid](https://github.com/BishopFox/badPods/tree/main/manifests/priv-and-hostpid)

[https://github.com/BishopFox/badPods/tree/main/manifests/priv-and-hostpid](https://github.com/BishopFox/badPods/tree/main/manifests/priv-and-hostpid)

### References and further reading

### 参考和进一步阅读

- [Duffie Cooley's Nsenter Pod Tweet](https://twitter.com/mauilion/status/1129468485480751104)
- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)
- [Node-shell Krew Plugin](https://github.com/kvaps/kubectl-node-shell)

- [Duffie Cooley 的 Nsenter Pod Tweet](https://twitter.com/mauilion/status/1129468485480751104)
- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)
- [Node-shell Krew 插件](https://github.com/kvaps/kubectl-node-shell)

## BAD POD \#3: PRIVILEGED ONLY

## BAD POD \#3：仅限特权

![BAD POD #3: PRIVILEGED ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2003%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2003%20-%20300PPI%20(1).jpg)

).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2003%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

### 可能发生的最坏情况是什么？

Multiple paths to full cluster compromise

完整集群妥协的多种途径

### How?

If you only have `privileged: true`, there are two paths you can take:

如果你只有 `privileged: true`，你可以选择两条路径：

- **Mount the host’s filesystem** — In privileged mode, `/dev` on the host is accessible in your pod. You can mount the disk that contains the host’s filesystem into your pod using the `mount` command. In my experience, this gives you a limited view of the filesystem though. Some files, and therefore privesc paths, are not accessible from your privileged pod unless you escalate to a full shell on the node. That said, it is easy enough that you might as well mount the device and see what you can see.

- **挂载主机的文件系统** - 在特权模式下，主机上的 `/dev` 可以在你的 pod 中访问。您可以使用 `mount` 命令将包含主机文件系统的磁盘挂载到 pod 中。以我的经验，这让您对文件系统的看法有限。除非您升级到节点上的完整 shell，否则无法从您的特权 pod 访问某些文件，因此无法访问 privesc 路径。也就是说，这很容易，您可以安装设备并查看您可以看到的内容。

- **Exploit `cgroup`** **user mode helper programs**— Your best bet is to get interactive root access on the node, but you must jump through a few hoops first. You can use Felix Wilhelm's exploit PoC [undock.sh](https://twitter.com/_fel1x/status/1151487051986087936) to execute one command a time, or you can use Brandon Edwards and Nick Freeman's version from their talk [A Compendium of Container Escapes](https://www.youtube.com/watch?v=BQlqita2D2s), which forces the host to connect back to the listener on the pod for an easy upgrade to interactive root access on the host. Another option is to use the Metasploit module [Docker Privileged Container Escape](https://www.rapid7.com/db/modules/exploit/linux/local/docker_privileged_container_escape/), which uses the same exploit to upgrade a shell received from a container to a shell on the host.


- **利用`cgroup`** **用户模式帮助程序**——你最好的办法是在节点上获得交互式root访问权限，但你必须先跳过几个环节。您可以使用 Felix Wilhelm 的漏洞利用 PoC [undock.sh](https://twitter.com/_fel1x/status/1151487051986087936) 一次执行一个命令，或者您可以使用 Brandon Edwards 和 Nick Freeman 的演讲 [A Compendium] Container Escapes](https://www.youtube.com/watch?v=BQlqita2D2s)，它强制主机连接回 Pod 上的侦听器，以便轻松升级到主机上的交互式 root 访问。另一种选择是使用 Metasploit 模块 [Docker Privileged Container Escape](https://www.rapid7.com/db/modules/exploit/linux/local/docker_privileged_container_escape/)，它使用相同的漏洞来升级从一个容器到主机上的外壳。


Whichever option you choose, the Kubernetes privilege escalation paths are largely the same as the [Bad Pod #1: Everything-allowed](http://labs.bishopfox.com#Pod1).

无论您选择哪个选项，Kubernetes 权限提升路径与 [Bad Pod #1: Everything-allowed](http://labs.bishopfox.com#Pod1) 大致相同。

### Usage and exploitation examples

### 使用和利用示例

[https://github.com/BishopFox/badPods/tree/main/manifests/priv](https://github.com/BishopFox/badPods/tree/main/manifests/priv)

[https://github.com/BishopFox/badPods/tree/main/manifests/priv](https://github.com/BishopFox/badPods/tree/main/manifests/priv)

### References and further reading

### 参考和进一步阅读

- [Felix Wilhelm's Cgroup Usermode Helper Exploit](https://twitter.com/_fel1x/status/1151487051986087936)
- [Understanding Docker Container Escapes](https://blog.trailofbits.com/2019/07/19/understanding-docker-container-escapes/)
- [A Compendium of Container Escapes](https://www.youtube.com/watch?v=BQlqita2D2s)
- [Docker Privileged Container Escape Metasploit Module](https://www.rapid7.com/db/modules/exploit/linux/local/docker_privileged_container_escape/)

- [Felix Wilhelm 的 Cgroup Usermode Helper Exploit](https://twitter.com/_fel1x/status/1151487051986087936)
- [了解Docker容器逃逸](https://blog.trailofbits.com/2019/07/19/understanding-docker-container-escapes/)
- [集装箱逃生纲要](https://www.youtube.com/watch?v=BQlqita2D2s)
- [Docker 特权容器逃逸 Metasploit 模块](https://www.rapid7.com/db/modules/exploit/linux/local/docker_privileged_container_escape/)

## BAD POD \#4: HOSTPATH ONLY

## 坏 POD \#4：仅主机路径

![BAD POD #4: HOSTPATH ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2004%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2004%20-%20300PPI%20(1).jpg)

).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2004%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

### 可能发生的最坏情况是什么？

Multiple paths to full cluster compromise

完整集群妥协的多种途径

### How?

In this case, even if you don't have access to the host's process or network namespaces, if the administrators have not limited what you can mount, you can mount the entire host's filesystem into your pod, giving you read/write access on the host's filesystem. This allows you to execute most of the same privilege escalation paths outlined above. There are so many paths available that Ian Coldwater and Duffie Cooley gave an awesome Black Hat 2019 talk about it titled “ [The Path Less Traveled: Abusing Kubernetes Defaults!](https://www.youtube.com/watch?v=HmoVSmTIOxM )”

在这种情况下，即使您无权访问主机的进程或网络命名空间，如果管理员没有限制您可以挂载的内容，您也可以将整个主机的文件系统挂载到您的 pod 中，从而为您提供读/写访问权限主机的文件系统。这允许您执行上面概述的大多数相同的权限提升路径。有很多可用的路径，以至于 Ian Coldwater 和 Duffie Cooley 在 2019 年的黑帽大会上发表了一篇很棒的演讲，题目是“[The Path Less Traveled: Abusing Kubernetes Defaults!](https://www.youtube.com/watch?v=HmoVSmTIOxM )”

Here are some privileged escalation paths that apply any time you have access to a Kubernetes node’s filesystem:

以下是一些特权升级路径，在您有权访问 Kubernetes 节点的文件系统时适用：

- **Look for `kubeconfig` files on the host filesystem** — If you are lucky, you will find a `cluster-admin` config with full access to everything. 
- **在主机文件系统上查找 `kubeconfig` 文件** - 如果幸运的话，你会找到一个可以完全访问所有内容的 `cluster-admin` 配置。
- **Access the tokens from all pods on the node** — Use something like `kubectl auth can-i --list` or [access-matrix](https://github.com/corneliusweig/rakkess) to see if any of the pods have tokens that give you more permissions than you currently have. Look for tokens that have permissions to get secrets or create pods, deployments, etc., in `kube-system`, or that allow you to create clusterrolebindings.

- **从节点上的所有 pod 访问令牌** — 使用类似 `kubectl auth can-i --list` 或 [access-matrix](https://github.com/corneliusweig/rakkess) 之类的东西来查看是否任何 Pod 都有令牌，可以为您提供比当前更多的权限。在“kube-system”中查找有权获取机密或创建 pod、部署等的令牌，或者允许您创建集群角色绑定的令牌。

- **Add your SSH key** — If you have network access to SSH to the node, you can add your public key to the node and SSH to it for full interactive access.

- **添加您的 SSH 密钥** — 如果您可以通过网络访问节点的 SSH，则可以将您的公钥添加到节点并通过 SSH 连接到节点以进行完全交互访问。

- **Crack hashed passwords** — Crack hashes in `/etc/shadow`; see if you can use them to access other nodes.


- **破解散列密码**——破解`/etc/shadow`中的散列；看看您是否可以使用它们来访问其他节点。


### Usage and exploitation examples

### 使用和利用示例

[https://github.com/BishopFox/badPods/tree/main/manifests/hostpath](https://github.com/BishopFox/badPods/tree/main/manifests/hostpath)

[https://github.com/BishopFox/badPods/tree/main/manifests/hostpath](https://github.com/BishopFox/badPods/tree/main/manifests/hostpath)

### References and further reading

### 参考和进一步阅读

- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)
- [Secure Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [Deep Dive into Real-World Kubernetes Threats](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [Compromising Kubernetes Cluster by Exploiting RBAC Permissions](https://www.youtube.com/watch?v=1LMo0CftVC4) ( [slides](https://published-prd.lanyonevents.com/published/rsaus20/sessionsFiles/ 18100/2020_USA20_DSO-W01_01_Compromising%20Kubernetes%20Cluster%20by%20Exploiting%20RBAC%20Permissions.pdf))

- [The Path Less Traveled: Abusing Kubernetes Defaults](https://www.youtube.com/watch?v=HmoVSmTIOxM) & [Corresponding Repo](https://github.com/mauilion/blackhat-2019)
- [安全 Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [深入了解真实世界的 Kubernetes 威胁](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [通过利用 RBAC 权限破坏 Kubernetes 集群](https://www.youtube.com/watch?v=1LMo0CftVC4) ([幻灯片](https://published-prd.lanyonevents.com/published/rsaus20/sessionsFiles/ 18100/2020_USA20_DSO-W01_01_Compromising%20Kubernetes%20Cluster%20by%20Exploiting%20RBAC%20Permissions.pdf))

## BAD POD \#5: HOSTPID ONLY

## BAD POD \#5：仅主机PID

![BAD POD #5: HOSTPID ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2005%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2005%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

### 可能发生的最坏情况是什么？

Application or cluster credential leaks if an application in the cluster is configured incorrectly. Denial of service via process termination.

如果集群中的应用程序配置不正确，则应用程序或集群凭据泄漏。通过进程终止拒绝服务。

### How?

There’s no clear path to get root on the node with only `hostPID`, but there are still some good post-exploitation opportunities.

没有明确的路径可以在只有 `hostPID` 的节点上获得 root 权限，但仍然有一些很好的后开发机会。

- **View processes on the host** — When you run `ps` from within a pod that has `hostPID: true`, you see all the processes running on the host, including processes running in each pod.

- **查看主机上的进程** — 当您从具有 `hostPID: true` 的 pod 中运行 `ps` 时，您会看到主机上运行的所有进程，包括在每个 pod 中运行的进程。

- **Look for passwords, tokens, keys, etc.** — If you are lucky, you will find credentials and you can then use them to escalate privileges in the cluster, to escalate privileges to services supported by the cluster, or to escalate privileges to services that communicate with cluster-hosted applications. It’s a long shot, but you might find a Kubernetes service account token or some other authentication material that will allow you to access other namespaces and eventually escalate all the way to cluster admin.

- **查找密码、令牌、密钥等** — 如果幸运的话，您会找到凭据，然后您可以使用它们来提升集群中的权限，将权限提升到集群支持的服务，或者将权限提升到与集群托管应用程序通信的服务。这是一个长镜头，但您可能会发现 Kubernetes 服务帐户令牌或其他一些身份验证材料，它们将允许您访问其他命名空间并最终一直升级到集群管理员。

- **Kill processes** — You can also kill any process on the node (presenting a denial-of-service risk). Because of this risk though, I would advise against it on a penetration test!


- **杀死进程** — 您还可以杀死节点上的任何进程（存在拒绝服务风险）。但是，由于存在这种风险，我建议不要在渗透测试中使用它！


### Usage and exploitation examples

### 使用和利用示例

[https://github.com/BishopFox/badPods/tree/main/manifests/hostpid](https://github.com/BishopFox/badPods/tree/main/manifests/hostpid)

[https://github.com/BishopFox/badPods/tree/main/manifests/hostpid](https://github.com/BishopFox/badPods/tree/main/manifests/hostpid)

## BAD POD \#6: HOSTNETWORK ONLY

## BAD POD \#6：仅限主机网络

![BAD POD #6: HOSTNETWORK ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2006%20-%20300PPI%20(2).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2006%20-%20300PPI%20(2).jpg)

).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2006%20-%20300PPI%20(2).jpg)

### What’s the worst that can happen?

### 可能发生的最坏情况是什么？

Potential path to cluster compromise

集群入侵的潜在途径

### How?

If you only have `hostNetwork: true`, you can’t get privileged code execution on the host directly, but if you cross your fingers, you might still find a path to cluster admin. There are three potential escalation paths:

如果你只有 `hostNetwork: true`，你不能直接在主机上获得特权代码执行，但如果你交叉手指，你可能仍然找到集群管理员的路径。存在三种潜在的升级路径：

- **Sniff traffic** — You can use `tcpdump` to sniff unencrypted traffic on any interface on the host. You might get lucky and find service account tokens or other sensitive information that is transmitted over unencrypted channels.

- **嗅探流量** — 您可以使用 `tcpdump` 来嗅探主机上任何接口上的未加密流量。您可能会幸运地找到通过未加密通道传输的服务帐户令牌或其他敏感信息。

- **Access services bound to localhost** — You can also reach services that only listen on the host’s loopback interface or that are otherwise blocked by network policies. These services might turn into a fruitful privilege escalation path.

- **访问绑定到本地主机的服务** - 您还可以访问仅侦听主机的环回接口或被网络策略阻止的服务。这些服务可能会变成一个富有成效的特权升级路径。

- **Bypass network policy** — If a restrictive network policy is applied to the namespace, deploying a pod with `hostNetwork: true` allows you to bypass the restrictions. This works because you are bound to the host's network interfaces and not the pods.


- **绕过网络策略** — 如果对命名空间应用了限制性网络策略，则使用 `hostNetwork: true` 部署 pod 可以让您绕过限制。这是有效的，因为您绑定到主机的网络接口而不是 pod。


### Usage and exploitation examples

### 使用和利用示例

[https://github.com/BishopFox/badPods/tree/main/manifests/hostnetwork](https://github.com/BishopFox/badPods/tree/main/manifests/hostnetwork)

[https://github.com/BishopFox/badPods/tree/main/manifests/hostnetwork](https://github.com/BishopFox/badPods/tree/main/manifests/hostnetwork)

## BAD POD \#7: HOSTIPC ONLY 
## BAD POD \#7：仅限主机
![BAD POD #7: HOSTIPC ONLY](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2007%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2007%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

### 可能发生的最坏情况是什么？

Ability to access data used by any pods that also use the host’s IPC namespace

能够访问也使用主机的 IPC 命名空间的任何 pod 使用的数据

### How?

If any process on the host or any processes in a pod uses the host’s inter-process communication mechanisms (shared memory, semaphore arrays, message queues, etc.), you’ll be able to read/write to those same mechanisms. The first place you'll want to look is `/dev/shm`, as it is shared between any pod with `hostIPC: true` and the host. You'll also want to check out the other IPC mechanisms with `ipcs`.

如果主机上的任何进程或 Pod 中的任何进程使用主机的进程间通信机制（共享内存、信号量数组、消息队列等），您将能够读取/写入这些相同的机制。您要查看的第一个位置是 `/dev/shm`，因为它在具有 `hostIPC: true` 的任何 pod 和主机之间共享。您还需要使用“ipcs”查看其他 IPC 机制。

- **Inspect /dev/shm** — Look for any files in this shared memory location.

- **Inspect /dev/shm** — 在此共享内存位置查找任何文件。

- **Inspect existing IPC facilities** — You can check to see if any IPC facilities are being used with `/usr/bin/ipcs`.


- **检查现有的 IPC 设施** — 您可以检查是否有任何 IPC 设施正在与 `/usr/bin/ipcs` 一起使用。


### Usage and exploitation examples

### 使用和利用示例

[https://github.com/BishopFox/badPods/tree/main/manifests/hostipc](https://github.com/BishopFox/badPods/tree/main/manifests/hostipc)

[https://github.com/BishopFox/badPods/tree/main/manifests/hostipc](https://github.com/BishopFox/badPods/tree/main/manifests/hostipc)

## BAD POD \#8: NOTHING ALLOWED

## 坏 POD \#8：不允许

![BAD POD #8: NOTHING ALLOWED](https://labs.bishopfox.com/hs-fs/hubfs/Check%20Box%20-%20Door%20Idea%20FINAL%2008%20-%20300PPI%20(1).jpg?width=700&name=Check%20Box%20-%20Door%20Idea%20FINAL%2008%20-%20300PPI%20(1).jpg)

### What’s the worst that can happen?

### 可能发生的最坏情况是什么？

Multiple potential paths to full cluster compromise

完全入侵集群的多种潜在途径

### How?

To close our bad Pods lineup, there are plenty of attack paths that should be investigated any time you can create a pod or simply have access to a pod, even if there are no security attributes enabled. Here are some things to look for whenever you have access to a Kubernetes pod:

为了关闭我们的坏 Pod 阵容，只要您可以创建 Pod 或只是访问 Pod，即使没有启用安全属性，也应该调查大量攻击路径。当您可以访问 Kubernetes pod 时，需要注意以下几点：

- **Accessible cloud metadata** — If the pod is cloud hosted, try to access the cloud metadata service. You might get access to the IAM credentials associated with the node or even just find a cloud IAM credential created specifically for that pod. In either case, this can be your path to escalate in the cluster, in the cloud environment, or in both.

- **可访问的云元数据** — 如果 Pod 是云托管的，请尝试访问云元数据服务。您可能会访问与节点关联的 IAM 凭证，甚至只是找到专门为该 Pod 创建的云 IAM 凭证。在任何一种情况下，这都可以是您在集群、云环境或两者中升级的路径。

- **Overly permissive service accounts** — If the namespace's default service account is mounted to `/var/run/secrets/kubernetes.io/serviceaccount/token` in your pod and is overly permissive, use that token to further escalate your privileges within the cluster.

- **过度宽松的服务账户** — 如果命名空间的默认服务账户挂载到您的 pod 中的 `/var/run/secrets/kubernetes.io/serviceaccount/token` 并且过于宽松，请使用该令牌进一步升级您的集群内的特权。

- **Misconfigured Kubernetes components** — If either [the apiserver or the kubelets have `anonymous-auth` set to](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/) [`true`](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/) and there are no network policy controls preventing it, you can interact with them directly without authentication.

- **错误配置的 Kubernetes 组件** — 如果 [apiserver 或 kubelets 将 `anonymous-auth` 设置为](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/) [`true`](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/) 并且没有网络策略控制阻止它，您可以直接与它们交互而无需身份验证。

- **Kernel, container engine, or Kubernetes exploits** — An unpatched exploit in the underlying kernel, in the container engine, or in Kubernetes can potentially allow a container escape or access to the Kubernetes cluster without any additional permissions.

- **内核、容器引擎或 Kubernetes 漏洞利用** — 底层内核、容器引擎或 Kubernetes 中未修补的漏洞利用可能允许容器逃逸或访问 Kubernetes 集群，而无需任何额外权限。

- **Hunt for vulnerable services**— Your pod will likely see a different view of the network services running in the cluster than you can see from the machine you used to create the pod. You can hunt for vulnerable services and applications by proxying your traffic through the pod.


- **寻找易受攻击的服务**——您的 pod 可能会看到与您在用于创建 pod 的机器上看到的集群中运行的网络服务不同的视图。您可以通过 Pod 代理流量来寻找易受攻击的服务和应用程序。


### Usage and exploitation examples

### 使用和利用示例

[https://github.com/BishopFox/badPods/tree/main/manifests/nothing-allowed](https://github.com/BishopFox/badPods/tree/main/manifests/nothing-allowed)

[https://github.com/BishopFox/badPods/tree/main/manifests/nothing-allowed](https://github.com/BishopFox/badPods/tree/main/manifests/nothing-allowed)

### References and further reading

### 参考和进一步阅读

- [Secure Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [Kubernetes Goat](https://madhuakula.com/kubernetes-goat/)
- [Attacking Kubernetes through Kubelet](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/)
- [Deep Dive into Real-World Kubernetes Threats](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [A Compendium of Container Escapes](https://www.youtube.com/watch?v=BQlqita2D2s)
- [CVE-2020-8558 POC](https://github.com/tabbysable/POC-2020-8558)

- [安全 Kubernetes - KubeCon NA 2019 CTF](https://securekubernetes.com/)
- [Kubernetes Goat](https://madhuakula.com/kubernetes-goat/)
- [通过 Kubelet 攻击 Kubernetes](https://labs.f-secure.com/blog/attacking-kubernetes-through-kubelet/)
- [深入了解真实世界的 Kubernetes 威胁](https://research.nccgroup.com/2020/02/12/command-and-kubectl-talk-follow-up/)
- [集装箱逃生纲要](https://www.youtube.com/watch?v=BQlqita2D2s)
- [CVE-2020-8558 POC](https://github.com/tabbysable/POC-2020-8558)

## CONCLUSION

## 结论

Apart from the [Bad Pod #8: Nothing Allowed](http://labs.bishopfox.com#pod8) example, all of the privilege escalation paths covered in this blog post (and the [respective repository](https:// github.com/BishopFox/badPods)) can be mitigated with restrictive pod security policies. 
除了 [Bad Pod #8：Nothing Allowed](http://labs.bishopfox.com#pod8) 示例之外，本博文（以及 [各自的存储库](https:// github.com/BishopFox/badPods)) 可以通过限制性 pod 安全策略来缓解。
Additionally, there are many other defense-in-depth security controls available to Kubernetes administrators that can reduce the impact of or completely thwart certain attack paths even when an attacker has access to some or all of the host namespaces and capabilities (eg, disabling the automatic mounting of service account tokens or requiring all pods to run as non-root by enforcing `MustRunAsNonRoot=true` and `allowPrivilegeEscalation=false`). As is always the case with penetration testing, your mileage may vary.

此外，Kubernetes 管理员还可以使用许多其他纵深防御安全控制，即使攻击者有权访问部分或全部主机命名空间和功能（例如，禁用自动挂载服务帐户令牌或通过强制执行“MustRunAsNonRoot=true”和“allowPrivilegeEscalation=false”来要求所有 pod 以非 root 用户身份运行）。与渗透测试的情况一样，您的里程可能会有所不同。

Administrators are sometimes hard pressed to defend security best practices without examples that demonstrate the security implications of risky configurations. I hope the examples laid out in this post and the manifests contained in the [Bad Pods repository](https://github.com/BishopFox/badPods) help you enforce the principle of least privilege when it comes to Kubernetes pod creation in your organization. 

管理员有时很难捍卫安全最佳实践，而没有示例来证明风险配置的安全影响。我希望本文中列出的示例以及 [Bad Pods 存储库](https://github.com/BishopFox/badPods) 中包含的清单可以帮助您在你组织中的 Kubernetes Pod 创建中执行最小特权原则。
