# Everyone might be a cluster-admin in your Kubernetes cluster

# 每个人都可能是 Kubernetes 集群中的集群管理员

February 27, 2020

Quite often, when I dive into someone's Kubernetes cluster to debug a problem, I realize whatever pod I'm running has _way_ too many permissions. Often, my pod has the `cluster-admin` role applied to it through its default ServiceAccount.

很多时候，当我深入某人的 Kubernetes 集群调试问题时，我意识到我正在运行的任何 pod 都_way_ 拥有太多权限。通常，我的 pod 通过其默认的 ServiceAccount 应用了 `cluster-admin` 角色。

Sometimes this role was added because someone wanted to make their CI/CD tool (eg Jenkins) manage Kubernetes resources in the cluster, and it was easier to apply `cluster-admin` to a default service account than to set all the individual [RBAC ](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) privileges correctly. Other times, it was because someone found a new shiny tool and blindly installed it.

有时添加这个角色是因为有人想让他们的 CI/CD 工具（例如 Jenkins）管理集群中的 Kubernetes 资源，并且将 `cluster-admin` 应用到默认服务帐户比设置所有个人 [RBAC ](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) 权限正确。其他时候，是因为有人发现了一个新的闪亮工具并盲目安装它。

One such example I remember seeing recently is the [spekt8](https://github.com/spekt8/spekt8) project; in it's installation instructions, it tells you to apply an rbac manifest:

我记得最近看到的一个这样的例子是 [spekt8](https://github.com/spekt8/spekt8) 项目；在它的安装说明中，它告诉您应用 rbac 清单：

```
kubectl apply -f https://raw.githubusercontent.com/spekt8/spekt8/master/fabric8-rbac.yaml
```


What the installation guide doesn't tell you is that this manifest grants `cluster-admin` privileges to _every single Pod in the default namespace_!

安装指南没有告诉您的是，此清单将 `cluster-admin` 权限授予_默认命名空间中的每个 Pod_！

At this point, a more naive reader (and I mean that in the nicest way—Kubernetes is a complex system and you need to learn a lot!) might say: "All the Pods I run are running trusted applications and I know my developers wouldn't do anything nefarious, so what's the big deal?"

在这一点上，一个更天真的读者（我的意思是最好的方式 - Kubernetes 是一个复杂的系统，你需要学习很多！）可能会说：“我运行的所有 Pod 都运行受信任的应用程序，我了解我的开发人员不会做坏事，有什么大不了的？”

The problem is, if any of your Pods are running _anything_ that could potentially result in code execution, it's trivial for a bad actor to script the following:

问题是，如果您的任何 Pod 正在运行 _anything_ 可能会导致代码执行，那么坏演员编写以下脚本是微不足道的：

1. Download`kubectl` in a pod
2. Execute a command like:

1. 在 pod 中下载`kubectl`
2. 执行如下命令：

`kubectl get ns --no-headers=true | sed "/kube-*/d" | sed "/default/d" | awk '{print $1;}' | xargs kubectl delete ns`





If you have your RBAC rules locked down, this is no big deal; even if the application inside the Pod is fully compromised, the Kubernetes API will deny any requests that aren't explicitly allowed by your RBAC rules.

如果您的 RBAC 规则被锁定，这没什么大不了的；即使 Pod 内的应用程序完全受到威胁，Kubernetes API 也会拒绝任何 RBAC 规则未明确允许的请求。

However, if you had blindly installed `spekt8` in your cluster, you would now have no namespaces left in your cluster, besides the default namespace.

但是，如果您在集群中盲目安装了 `spekt8`，那么现在除了默认命名空间之外，集群中将没有任何命名空间。

## Try it yourself

## 自己试试

I created a little project called [k8s-pod-rbac-breakout](https://github.com/geerlingguy/k8s-pod-rbac-breakout) that you can use to test whether your cluster has this problem—I typically deploy a script like this [58-line index.php](https://github.com/geerlingguy/k8s-pod-rbac-breakout/blob/master/index.php) script to a Pod running PHP in a cluster and see what it returns.

我创建了一个名为 [k8s-pod-rbac-breakout](https://github.com/geerlingguy/k8s-pod-rbac-breakout) 的小项目，你可以用它来测试你的集群是否有这个问题——我通常部署像这样的脚本 [58-line index.php](https://github.com/geerlingguy/k8s-pod-rbac-breakout/blob/master/index.php) 脚本到集群中运行 PHP 的 Pod 并查看它返回什么。

You'd be surprised how many clusters give me all the info and no errors:

您会惊讶于有多少集群为我提供了所有信息并且没有错误：

![RBAC Breakout page example](http://www.jeffgeerling.com/sites/default/files/images/rbac-breakout-page-example.png)

Too many Kubernetes users build snowflake clusters and deploy tools (like `spekt8`—though there are _many_, many others) into them with no regard for security, either because they don't understand Kubernetes' RBAC model, or they needed to meet a deadline.

太多的 Kubernetes 用户在不考虑安全性的情况下构建雪花集群并部署工具（如“spekt8”——尽管有_许多_，许多其他），要么是因为他们不了解 Kubernetes 的 RBAC 模型，要么是因为他们需要满足最后期限。

If you ever find yourself taking shortcuts to get past pesky `User "system:serviceaccount:default:default" cannot [do xyz]` messages, think twice before being promiscuous with your cluster permissions. And consider automating your cluster management (I'm writing a [book](https://www.ansibleforkubernetes.com) for that) so people can't blindly deploy insecure tools and configurations to it!

如果您发现自己采取了快捷方式来绕过讨厌的“用户“system:serviceaccount:default:default”无法 [do xyz]”消息，请在混淆集群权限之前三思而后行。并考虑自动化您的集群管理（我正在为此编写 [book](https://www.ansibleforkubernetes.com))，这样人们就不能盲目地向其部署不安全的工具和配置！

## Further reading

## 进一步阅读

- [Run Ansible Tower or AWX in Kubernetes or OpenShift with the Tower Operator](http://www.jeffgeerling.com/blog/2019/run-ansible-tower-or-awx-kubernetes-or-openshift-tower-operator)

- [使用 Tower Operator 在 Kubernetes 或 OpenShift 中运行 Ansible Tower 或 AWX](http://www.jeffgeerling.com/blog/2019/run-ansible-tower-or-awx-kubernetes-or-openshift-tower-operator)

- [Debugging networking issues with multi-node Kubernetes on VirtualBox](http://www.jeffgeerling.com/blog/2019/debugging-networking-issues-multi-node-kubernetes-on-virtualbox)

- [在 VirtualBox 上使用多节点 Kubernetes 调试网络问题](http://www.jeffgeerling.com/blog/2019/debugging-networking-issues-multi-node-kubernetes-on-virtualbox)

- [Monitoring Kubernetes cluster utilization and capacity (the poor man's way)](http://www.jeffgeerling.com/blog/2019/monitoring-kubernetes-cluster-utilization-and-capacity-poor-mans-way)

- [监控 Kubernetes 集群利用率和容量（穷人的方式）](http://www.jeffgeerling.com/blog/2019/monitoring-kubernetes-cluster-utilization-and-capacity-poor-mans-way)

## Comments



There is only one solution for this: we need a management tool that abstracts everything away on top of K8S! Also on top of that a cluster of cluster management tools that manage your management cluster. :)))

对此只有一个解决方案：我们需要一个管理工具，在 K8S 之上抽象一切！此外，还有一组用于管理您的管理集群的集群管理工具。 :)))

Seriously, for many years there existed a good practice of providing secure installations of software by default.

说真的，多年来一直存在默认提供安全软件安装的良好做法。

Delivering software with an insecure configuration as default was a bug. 

以不安全的配置作为默认配置交付软件是一个错误。

Today software is not secure by default and everybody is fine with that - why are people accepting this?

今天的软件默认是不安全的，每个人都可以接受——为什么人们会接受这一点？

Also look at the crazy amount of CPU cycles wasted on a full blown K8S cluster even before one single byte of useful payload is computed - while the planet is burning and future generations are begging for a change we are establishing a new data center operating system that needs even more energy - why are people accepting this kind of bad systems?

还要看看在一个完整的 K8S 集群上浪费的疯狂数量的 CPU 周期，甚至在计算一个字节的有用有效载荷之前 - 当地球正在燃烧，后代正在乞求改变时，我们正在建立一个新的数据中心操作系统需要更多的能量——为什么人们会接受这种糟糕的系统？

We had clustering systems for Linux for many years and students were able to install them as a weekend project - this knowledge is getting lost. Instead everybody thinks "new system" is cool just because it is from "big corp" - that is pure distopia brainwash, at that point we could just install Windows and "just restart your machine if a problem occurs".

我们拥有 Linux 集群系统已有多年，学生们能够在周末项目中安装它们——这些知识正在消失。相反，每个人都认为“新系统”很酷，因为它来自“大公司”——那是纯粹的异端洗脑，那时我们可以只安装 Windows 并“如果出现问题就重新启动机器”。

Measure, analyze, think, look at the facts, educate yourself and be brave enough to trust your own conclusions.

衡量、分析、思考、查看事实、教育自己并勇敢地相信自己的结论。

If a system looks bad, delivers bad results and generates too many new problems, maybe in fact it is a bad system.

如果一个系统看起来很糟糕，产生了糟糕的结果并产生了太多的新问题，那么实际上它可能是一个糟糕的系统。

Thanks for providing the tools that help with that! 

感谢您提供有助于解决此问题的工具！

