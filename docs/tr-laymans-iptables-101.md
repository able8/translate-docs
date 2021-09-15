# Illustrated introduction to Linux iptables

# Linux iptables 图解介绍

June 21, 2019 (Updated: August 30, 2021)

[Networking,](http://iximiuz.com/en/categories/?category=Networking) [Linux / Unix](http://iximiuz.com/en/categories/?category=Linux / Unix)

## Foreword

## 前言

Gee, it's my turn to throw some _gloom_ light on iptables! There are hundreds or even thousands of articles on the topic out there, including introductory ones. I'm not going to put either formal _and boring_ [definitions](https://www.netfilter.org/projects/iptables/index.html) here nor [long lists](https://www.digitalocean.com/community/tutorials/iptables-essentials-common-firewall-rules-and-commands) of [useful commands](https://www.digitalocean.com/community/tutorials/how-to-list-and-delete-iptables-firewall-rules). I would rather try to use layman's terms and scribbling as much as possible to give you some insights about the domain before going to all these _tables_, _rules_, _targets_, and _policies_. By the way, the first time I faced this tool I was pretty much confused by the terminology too!

哎呀，轮到我在 iptables 上放一些 _gloom_ 灯了！关于该主题的文章有成百上千篇，包括介绍性文章。我不会在这里放正式的_和无聊的_[定义](https://www.netfilter.org/projects/iptables/index.html)，也不会放[长列表](https://www.digitalocean.com/community/tutorials/iptables-essentials-common-firewall-rules-and-commands)的[有用命令](https://www.digitalocean.com/community/tutorials/how-to-list-and-delete-iptables-防火墙规则)。在阅读所有这些_tables_、_rules_、_targets_ 和_policies_ 之前，我宁愿尝试使用外行的术语并尽可能多地草草为您提供有关该域的一些见解。顺便说一句，当我第一次面对这个工具时，我也对术语感到非常困惑！

Probably, you already know that **iptables** has something to do with **IP** packets. Maybe even deeper - packets filtration. Or the deepest - packets modification! And maybe you've heard, that everything is happening on the kernel side, without user space code involved. For this, iptables provides a special syntax to encode different packets-affecting rules...

可能您已经知道 **iptables** 与 **IP** 数据包有关。也许更深层次 - 数据包过滤。还是最深的——包修改！也许你听说过，一切都发生在内核端，不涉及用户空间代码。为此，iptables 提供了一种特殊的语法来对不同的影响数据包的规则进行编码......

## Linux network stack

## Linux 网络堆栈

...but before trying to make an impact on a happy life of packets in the kernel space, let's try to understand their universe. When packets get created, what are their paths inside of the kernel, what are their origins and destinations, etc? Have a look at the following scenarios:

...但在尝试对内核空间中数据包的幸福生活产生影响之前，让我们尝试了解它们的宇宙。当数据包被创建时，它们在内核中的路径是什么，它们的来源和目的地是什么等等？看看以下场景：

- A packet arrives to the network interface, passes through the network stack and reaches a user space process.

- 数据包到达网络接口，通过网络堆栈并到达用户空间进程。

- A packet is created by a user space process, sent to the network stack, and then delivered to the network interface.

- 数据包由用户空间进程创建，发送到网络堆栈，然后传递到网络接口。

- A packet arrives to the network interface and then in accordance with some routing rules is forwarded to another network interface.

- 数据包到达网络接口，然后根据某些路由规则转发到另一个网络接口。

What is the _commonality_ amongst all those scenarios? Basically, all of them describe pavings of the packets' ways from a network interface through the network stack to a user space process (or another interface) and turnarounds. When I say _a network stack_ here I just mean a bunch of layers provided by the Linux kernel to handle the network data transmission and receiving.

所有这些场景中的_共性_是什么？基本上，它们都描述了从网络接口到网络堆栈到用户空间进程（或另一个接口）和周转的数据包路径的铺路。当我在这里说_网络堆栈_时，我只是指Linux内核提供的一堆层来处理网络数据传输和接收。

![Linux network stack architecture](http://iximiuz.com/laymans-iptables-101/iptables-overview-white.png)

Routing part in the middle is provided by the built-in capability of the Linux kernel, also known as _IP forwarding_. Sending a non-zero value to `/proc/sys/net/ipv4/ip_forward` file activates packet forwarding between different network interfaces, effectively turning a Linux machine into a virtual router.

中间的路由部分由Linux内核的内置能力提供，也称为_IP转发_。向 `/proc/sys/net/ipv4/ip_forward` 文件发送一个非零值会激活不同网络接口之间的数据包转发，有效地将 Linux 机器变成一个虚拟路由器。

It's more or less obvious, that a properly engineered network stack should have different logical stages of the packet processing. For example, a _PREROUTING_ stage could reside somewhere in between the packets ingestion and the actual routing procedure. Another example could be an _INPUT_ stage, residing immediately before the user space process.

或多或少很明显，正确设计的网络堆栈应该具有不同的数据包处理逻辑阶段。例如，_PREROUTING_ 阶段可能位于数据包摄取和实际路由过程之间的某个位置。另一个例子可能是一个 _INPUT_ 阶段，它位于用户空间进程之前。

![iptables packet processing stages](http://iximiuz.com/laymans-iptables-101/iptables-stages-white.png)

In fact, Linux network stack does provide such logical separation of stages. Now, let's get back to our primary task - packets filtration and/or alteration. What if we want to drop some packets arriving to _a.out_ process? For example, we might dislike packets with one particular source IP address, because we suspect this IP belonging to a malicious user. Would be great to have _a hook_ in the network stack, corresponding to the _INPUT_ stage and allowing some extra logic to be applied to incoming packets. In our example, we might want to inject a function to check the packet's source IP address and based on this information decide whether to drop or accept the packet.

事实上，Linux 网络堆栈确实提供了这种逻辑上的阶段分离。现在，让我们回到我们的主要任务 - 数据包过滤和/或更改。如果我们想丢弃一些到达 _a.out_ 进程的数据包怎么办？例如，我们可能不喜欢具有一个特定源 IP 地址的数据包，因为我们怀疑该 IP 属于恶意用户。在网络堆栈中有 _a hook_ 会很棒，对应于 _INPUT_ 阶段并允许将一些额外的逻辑应用于传入的数据包。在我们的示例中，我们可能想要注入一个函数来检查数据包的源 IP 地址，并根据此信息决定是丢弃还是接受数据包。

Generalizing, we need a way to register an arbitrary callback function to be executed on every incoming packet at a given stage. Luckily enough, there is a project called [**netfilter**](https://www.netfilter.org/) that provides exactly this functionality! The netfilter's code resides inside of the Linux kernel and adds all those extension points ( _i.e. hooks_) to different stages of the network stack. It is noteworthy that _iptables_ is just one amongst several of the user space frontend tools to configure the netfilter hooks. Even more to note here - the functionality of the netfilter is not limited by the network (i.e. IP) layer, for example, the modification of ethernet frames is also possible. However, as it follows from its name, **ip** tables is focusing on the layers starting from the [network](https://en.wikipedia.org/wiki/Network_layer) ( **IP**) and above .

概括地说，我们需要一种方法来注册一个任意回调函数，以便在给定阶段的每个传入数据包上执行。幸运的是，有一个名为 [**netfilter**](https://www.netfilter.org/) 的项目正好提供了这个功能！ netfilter 的代码位于 Linux 内核内部，并将所有这些扩展点（_i.e. hooks_）添加到网络堆栈的不同阶段。值得注意的是，_iptables_ 只是配置 netfilter 钩子的几个用户空间前端工具之一。这里还要注意 - netfilter 的功能不受网络（即 IP)层的限制，例如，以太网帧的修改也是可能的。然而，正如它的名字一样，**ip** 表专注于从 [network](https://en.wikipedia.org/wiki/Network_layer) ( **IP**) 开始的层.

## iptables Chains Introduction

## iptables 链介绍

Now, let's finally try to understand the iptables terminology. You may already notice that the names we use for the **_stages_** in the network stack correspond to the iptables **_chains_**. But why on earth somebody would use the word _chain_ for it? I don't know any anecdote behind it, but one way to explain the naming is to have a look at the usage:

现在，让我们最终尝试理解 iptables 术语。您可能已经注意到，我们在网络堆栈中用于 **_stages_** 的名称对应于 iptables **_chains_**。但是为什么有人会用_chain_这个词来表示它呢？我不知道它背后的任何轶事，但解释命名的一种方法是查看用法：

```bash
# add rule "LOG every packet" to chain INPUT
$ iptables --append INPUT --jump LOG

# add rule "DROP every packet" to chain INPUT
$ iptables --append INPUT --jump DROP

```

In the snippet above we added multiple callbacks to the INPUT _stage_, and this is absolutely legitimate iptables usage. It implies though that the order of execution of the callbacks have to be defined. In reality, when a new packet arrives, the first added callback is executed first (LOG the packet), then the second callback is executed (DROP the packet). Thus, all our callbacks have been lined up in a **_chain_**! But the chain is named by the logical stage it resides. For now, let's finish with chains and switch to other parts of iptables. Later on, we will see, that there is some ambiguity in the chain abstraction.

在上面的代码片段中，我们向 INPUT _stage_ 添加了多个回调，这是绝对合法的 iptables 用法。这意味着必须定义回调的执行顺序。实际上，当新数据包到达时，首先执行第一个添加的回调（LOG 数据包），然后执行第二个回调（DROP 数据包）。因此，我们所有的回调都排列在 **_chain_** 中！但是链是由它所在的逻辑阶段命名的。现在，让我们完成链并切换到 iptables 的其他部分。稍后，我们将看到，链抽象中存在一些歧义。

## iptables Rules, Targets, and Policies

## iptables 规则、目标和策略

Next, goes _a rule_. The rules we used in our example above are rudimentary. First, we unconditionally LOG a kernel message for every packet in the INPUT chain, then we unconditionally DROP every packet from the network stack. However, rules can be more complicated. In general, a rule specifies criteria for a packet and a _target_. For a sake of simplicity, let's define _target_ now as just an action, like LOG, ACCEPT, or DROP and have a look at some examples:

接下来是_规则_。我们在上面的例子中使用的规则是基本的。首先，我们无条件地为 INPUT 链中的每个数据包记录一条内核消息，然后我们无条件地从网络堆栈中删除每个数据包。但是，规则可能更复杂。通常，规则指定数据包和_目标_的标准。为简单起见，我们现在将 _target_ 定义为一个动作，例如 LOG、ACCEPT 或 DROP，并查看一些示例：

```bash
# block packets with source IP 46.36.222.157
# -A is a shortcut for --append
# -j is a shortcut for --jump
$ iptables -A INPUT -s 46.36.222.157 -j DROP

# block outgoing SSH connections
$ iptables -A OUTPUT -p tcp --dport 22 -j DROP

# allow all incoming HTTP(S) connections
$ iptables -A INPUT -p tcp -m multiport --dports 80,443 \
  -m conntrack --ctstate NEW,ESTABLISHED -j ACCEPT
$ iptables -A OUTPUT -p tcp -m multiport --dports 80,443 \
  -m conntrack --ctstate ESTABLISHED -j ACCEPT

```

As we can see, the rule's criteria can be quite complex. We can check multiple attributes of the packet, or even some properties of the TCP connections (it implies the netfilter to be stateful, thanks to _conntrack_ module) before deciding on the action. I'm sorry for this, but a programmer in me requires some code to be written:

正如我们所见，规则的标准可能非常复杂。在决定采取行动之前，我们可以检查数据包的多个属性，甚至是 TCP 连接的某些属性（这意味着 netfilter 是有状态的，感谢 _conntrack_ 模块）。我很抱歉，但我的程序员需要编写一些代码：

```python
def handle_packet(packet, chain):
for rule in chain:
    modules = rule.modules
    for m in modules:
      m.ensure_loaded()

    conditions = rule.conditions
    if all(c.apply(packet) for c in conditions):
      # terminal target, break the chain
      if rule.target in ('ACCEPT', 'DROP'):
        return rule.target

      # TODO: handle other targets

# TODO: what shall we do if there is no single
#       terminal target in the whole chain?

```

The idea is pretty simple. Sequentially apply all the rules in the chain until either a terminal target is encountered, or the end of the chain is reached. And here we can notice an uncovered branch in our pseudocode. We need a default _action_, (i.e. target) for packets managed to reach the end of the chain without being dispatched to any terminal target in the meantime. And the way to set it is called **policy**:

这个想法很简单。依次应用链中的所有规则，直到遇到终端目标或到达链的末端。在这里我们可以注意到伪代码中有一个未被覆盖的分支。我们需要一个默认的 _action_，（即目标），用于管理到达链末端的数据包，而同时不会被分派到任何终端目标。而设置它的方式叫做**policy**：

```
# check the default policies
$ sudo iptables --list-rules  # or -S
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT

# change policy for chain FORWARD to target DROP
iptables --policy FORWARD DROP  # or -P

```

## iptables Chains (continued)

## iptables 链（续）

Finally, let's learn why the targets are called _targets_, not actions or something else. Let's look at the command we've used to set a rule `iptables -A INPUT -s 46.36.222.157 -j DROP`, where `-j` stands for `--jumps`. That is, as a result of the rule we can _jump_ to a target. From `man iptables`:

最后，让我们了解为什么目标被称为 _targets_，而不是动作或其他东西。让我们看看我们用来设置规则的命令`iptables -A INPUT -s 46.36.222.157 -j DROP`，其中`-j`代表`--jumps`。也就是说，作为规则的结果，我们可以_跳转_到一个目标。来自`man iptables`：

> ```
>  -j, --jump target
>       This specifies the target of the rule;i.e., what to do
>       if the packet matches it.The target can be a user-defined
>       chain (other than the one this rule is in), one of the
>       special builtin targets which decide the fate of the packet
>       immediately, or an extension (see EXTENSIONS below).
> ```

Here it is! **User-defined chains!** As usually, first let's look at the example:

这里是！ **用户定义的链！** 通常，首先让我们看一下示例：

```bash
$ iptables -P INPUT ACCEPT
# drop all forwards by default
$ iptables -P FORWARD DROP
$ iptables -P OUTPUT ACCEPT

# create a new chain
$ iptables -N DOCKER  # or --new-chain

# if outgoing interface is docker0, jump to DOCKER chain
$ iptables -A FORWARD -o docker0 -j DOCKER

# add some specific to Docker rules to the user-defined chain
$ iptables -A DOCKER ...
$ iptables -A DOCKER ...
$ iptables -A DOCKER ...

# jump back to the caller (i.e. FORWARD) chain
$ iptables -A DOCKER -j RETURN

```

But how come? As we saw above, chains have had a one-to-one correspondence to the predefined logical stages of the network stack.
Does the fact that users can define their own chains mean we can introduce new stages to the kernel's handling pipeline? I hardly think so. I might be totally wrong here, but to me, it seems like a violation of the [Single responsibility principle](https://en.wikipedia.org/wiki/Single_responsibility_principle). A _chain_ seems to be a good abstraction for a named sequence of rules. There is some similarity between chains and [named subroutines](https://en.wikipedia.org/wiki/Subroutine) (aka functions, aka procedures) in traditional programming languages. Ability to jump from an arbitrary place in one chain to the beginning of another chain and then RETURN to the caller chain makes the similarity even stronger. However, _PREROUTING_, _INPUT_, _FORWARD_, _OUTPUT_, and _POSTROUTING_ chains have a special meaning and cannot be overwritten. I can see some similarity to the _main()_ function in some programming languages having a special purpose, but this double-ended nature of chains made the learning curve of iptables pretty steep to me.

但是怎么来的？正如我们在上面看到的，链与网络堆栈的预定义逻辑阶段一一对应。
用户可以定义自己的链这一事实是否意味着我们可以向内核的处理管道引入新的阶段？我几乎不这么认为。我在这里可能完全错了，但对我来说，这似乎违反了[单一责任原则](https://en.wikipedia.org/wiki/Single_responsibility_principle)。 _chain_ 似乎是对命名规则序列的一个很好的抽象。链和传统编程语言中的[命名子例程](https://en.wikipedia.org/wiki/Subroutine)（又名函数，又名过程)之间有一些相似之处。从一条链中的任意位置跳转到另一条链的开头，然后返回到调用者链的能力使相似性更加强大。但是，_PREROUTING_、_INPUT_、_FORWARD_、_OUTPUT_ 和_POSTROUTING_ 链具有特殊含义，不能被覆盖。我可以看到某些具有特殊用途的编程语言中的 _main()_ 函数有一些相似之处，但是链的这种双端性质使 iptables 的学习曲线对我来说非常陡峭。

To summarise, a user-defined chain is a special kind of target, used as a named sequence of rules. The capabilities of user-defined chains are rather limited. For example, a user-defined chain can't have a policy. From `man iptables`:

总而言之，用户定义的链是一种特殊的目标，用作命名的规则序列。用户定义链的能力相当有限。例如，用户定义的链不能有策略。来自`man iptables`：

> ```
>  -P, --policy chain target
>       Set the policy for the chain to the given target.
>       See the section TARGETS for the legal targets.
>       Only built-in (non-user-defined) chains can have policies,
>       and neither built-in nor user-defined chains can be policy
>       targets.
> ```

![iptables user-defined chain example](http://iximiuz.com/laymans-iptables-101/user-defined-chains.png)

Obviously, the code snippet above should be significantly rewritten to incorporate handling of user-defined chains.

显然，上面的代码片段应该被大量重写以包含对用户定义链的处理。

## iptables Tables introduction

## iptables 表介绍

Well, we are almost there! We've covered _chains_, _rules_, and _policies_. Now it's finally time to learn about tables. After all, the tool is called ip **tables**. 

好吧，我们快到了！我们已经介绍了_chains_、_rules_ 和_policies_。现在终于到了学习表格的时候了。毕竟这个工具叫ip **tables**。

Actually, in all the examples above we implicitly used a table called `filter`. I'm not sure about the official definition of a _table_, but I always refer to a table as a logical grouping and isolation of chains. As we already know, there is a table for chains that manage packets filtration. However, if we want to modify some packets, there is another table, called _mangle_. It's absolutely valid desire to be able to filter packets on the FORWARD stage. However, it's also fine to modify packets on that stage. Hence, both _filter_ and _mangle_ tables will have FORWARD chains. However, those chains are completely independent.

实际上，在上面的所有示例中，我们都隐式地使用了一个名为“filter”的表。我不确定_table_的官方定义，但我总是将表称为链的逻辑分组和隔离。正如我们已经知道的，有一个用于管理数据包过滤的链表。但是，如果我们要修改一些数据包，还有另一个表，称为_mangle_。能够在 FORWARD 阶段过滤数据包是绝对有效的愿望。但是，在该阶段修改数据包也是可以的。因此，_filter_ 和 _mangle_ 表都将具有 FORWARD 链。但是，这些链是完全独立的。

The number of supported tables can vary between different versions of the kernel, but the most prominent tables are usually there:

不同版本的内核支持的表数量可能有所不同，但最突出的表通常在那里：

> ```
>  filter:
>       This is the default table (if no -t option is passed).It contains
>       the built-in chains INPUT (for packets destined to local sockets),
>       FORWARD (for packets being routed through the box), and OUTPUT
>       (for locally-generated packets).
> ```

> ```
>  nat:
>       This table is consulted when a packet that creates a new connection is encountered.
>       It consists of three built-ins: PREROUTING (for altering packets as soon as
>       they come in), OUTPUT (for altering locally-generated packets before routing),
>       and POSTROUTING (for altering packets as they are about to go out).IPv6 NAT support
>       is available since kernel 3.7.
> ```

> ```
>  mangle:
>       This table is used for specialized packet alteration.Until kernel 2.4.17 it had two
>       built-in chains: PREROUTING (for altering incoming packets before routing)
>       and OUTPUT (for altering locally-generated packets before routing).Since kernel 2.4.18,
>       three other built-in chains are also supported: INPUT (for packets coming into the box
>       itself), FORWARD (for altering packets being routed through the box), and POSTROUTING
>       (for altering packets as they are about to go out).
> ```

> ```
>  raw:
>       This table is used mainly for configuring exemptions from connection tracking in
>       combination with the NOTRACK target.It registers at the netfilter hooks with
>       higher priority and is thus called before ip_conntrack, or any other IP tables.
>       It provides the following built-in chains: PREROUTING (for packets arriving via
>       any network interface) OUTPUT (for packets generated by local processes)
> ```

> ```
>  security:
>       This table is used for Mandatory Access Control (MAC) networking rules, such as those
>       enabled by the SECMARK and CONNSECMARK targets.Mandatory Access Control is implemented
>       by Linux Security Modules such as SELinux.The security table is called after the filter
>       table, allowing any Discretionary Access Control (DAC) rules in the filter table to take
>       effect before MAC rules.This table provides the following built-in chains: INPUT (for
>       packets coming into the box itself), OUTPUT (for altering locally-generated packets before
>       routing), and FORWARD (for altering packets being routed through the box).
> ```

What is really interesting here is the collision of chains between tables. What will happen to a packet if _filter.INPUT_ chain has a DROP target but _mangle.INPUT_ chain has an ACCEPT target, both within the affirmative rules? Which chain has higher precedence? Let's try to check it out!

这里真正有趣的是表之间链的碰撞。如果 _filter.INPUT_ 链有一个 DROP 目标但 _mangle.INPUT_ 链有一个 ACCEPT 目标，数据包会发生什么，两者都在肯定规则内？哪个链的优先级更高？让我们试试看吧！

For this, we need to add LOG targets to all the chains of all the tables and conduct the following experiment:

为此，我们需要将 LOG 目标添加到所有表的所有链中并进行以下实验：

![Learning iptables with Linux network namespace demo](http://iximiuz.com/laymans-iptables-101/netns.png)

Click here to learn how to set up the environemnt. 

单击此处了解如何设置环境。

We need logs for both the client and the router network stacks on the same stream to have a comprehensive overview of the relations between tables and chains. For this we are going to simulate 2 machines on a single Linux host by using [network namespaces](https://en.wikipedia.org/wiki/Linux_namespaces#Network_(net)) feature. The main part of the experiment is LOGging of IP packets. For a long time, LOG target was disabled for the non-root namespaces to prevent potential host's denials of service from the isolated processes. Luckily, starting from Linux kernel 4.11 there is a way to enable netfilter logs for namespaces via sending a non-zero value to \`/proc/sys/net/netfilter/nf\_log\_all\_netns\`. \*\*But don't do this on production!\*\*

我们需要同一流上的客户端和路由器网络堆栈的日志，以全面了解表和链之间的关系。为此，我们将使用 [网络命名空间](https://en.wikipedia.org/wiki/Linux_namespaces#Network_(net)) 功能在单个 Linux 主机上模拟 2 台机器。实验的主要部分是 IP 数据包的 LOGging。长期以来，非根命名空间的 LOG 目标被禁用，以防止潜在的主机拒绝来自隔离进程的服务。幸运的是，从 Linux 内核 4.11 开始，有一种方法可以通过向 \`/proc/sys/net/netfilter/nf\_log\_all\_netns\` 发送一个非零值来为命名空间启用 netfilter 日志。 \*\*但是不要在生产中这样做！\*\*

First, let's create a network namespace:

首先，让我们创建一个网络命名空间：

```bash
# run on host:

$ unshare -r --net bash

# namespace (same terminal session):
$ echo $$
2979  # remember this PID

```

In the second terminal, we need to configure the host:

在第二个终端中，我们需要配置主机：

```bash
# run on host:

$ sudo -i

# create a veth interface
$ ip link add vGUEST type veth peer name vHOST

# move one of its peers to network namespace
$ ip link set vGUEST netns 2979  # PID from above

# create linux bridge
$ ip link add br0 type bridge

# wire vHOST to br0
$ ip link set vHOST master br0

# set IP addresses and bring devices up
$ ip addr add 172.16.0.1/16 dev br0
$ ip link set br0 up
$ ip link set vHOST up
$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 52:54:00:26:10:60 brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global noprefixroute dynamic eth0
       valid_lft 85752sec preferred_lft 85752sec
    inet6 fe80::5054:ff:fe26:1060/64 scope link
       valid_lft forever preferred_lft forever
3: vHOST@if4: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue master br0 state LOWERLAYERDOWN group default qlen 1000
    link/ether c2:96:cf:97:f4:12 brd ff:ff:ff:ff:ff:ff link-netnsid 0
5: br0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether c2:96:cf:97:f4:12 brd ff:ff:ff:ff:ff:ff
    inet 172.16.0.1/16 scope global br0
       valid_lft forever preferred_lft forever

# turn the host into a virtual router
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
$ echo 1 > /proc/sys/net/ipv4/ip_forward

# and don't forget to enable netfilter logs in namespaces
$ echo 1 > /proc/sys/net/netfilter/nf_log_all_netns

```

Now finish the network interface setup on the namespace side:

现在完成命名空间端的网络接口设置：

```bash
# run in network namespace:

# bring devices up
$ ip link set lo up
$ ip link set vGUEST up

# configure IP address
$ ip addr add 172.16.0.2/16 dev vGUEST
$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
4: vGUEST@if3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether c2:31:a8:8b:d7:f8 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.16.0.2/16 scope global vGUEST
       valid_lft forever preferred_lft forever
    inet6 fe80::c031:a8ff:fe8b:d7f8/64 scope link
       valid_lft forever preferred_lft forever

# set default route via br0
$ ip route add default via 172.16.0.1
$ ip route
default via 172.16.0.1 dev vGUEST
172.16.0.0/16 dev vGUEST proto kernel scope link src 172.16.0.2

```

Check out and update iptables rules in the namespace:

检查并更新命名空间中的 iptables 规则：

```bash
# run in network namespace:

$ iptables -S -t filter
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT

$ iptables -S -t nat
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT

$ iptables -S -t mangle
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT

$ iptables -S -t raw
-P PREROUTING ACCEPT
-P OUTPUT ACCEPT

$ iptables -t filter -A INPUT -j LOG --log-prefix "NETNS_FILTER_INPUT "
$ iptables -t filter -A FORWARD -j LOG --log-prefix "NETNS_FILTER_FORWARD "
$ iptables -t filter -A OUTPUT -j LOG --log-prefix "NETNS_FILTER_OUTPUT "

$ iptables -t nat -A PREROUTING -j LOG --log-prefix "NETNS_NAT_PREROUTE "
$ iptables -t nat -A INPUT -j LOG --log-prefix "NETNS_NAT_INPUT "
$ iptables -t nat -A OUTPUT -j LOG --log-prefix "NETNS_NAT_OUTPUT "
$ iptables -t nat -A POSTROUTING -j LOG --log-prefix "NETNS_NAT_POSTROUTE "

$ iptables -t mangle -A PREROUTING -j LOG --log-prefix "NETNS_MANGLE_PREROUTE "
$ iptables -t mangle -A INPUT -j LOG --log-prefix "NETNS_MANGLE_INPUT "
$ iptables -t mangle -A FORWARD -j LOG --log-prefix "NETNS_MANGLE_FORWARD "
$ iptables -t mangle -A OUTPUT -j LOG --log-prefix "NETNS_MANGLE_OUTPUT "
$ iptables -t mangle -A POSTROUTING -j LOG --log-prefix "NETNS_MANGLE_POSTROUTE "

$ iptables -t raw -A PREROUTING -j LOG --log-prefix "NETNS_RAW_PREROUTE "
$ iptables -t raw -A OUTPUT -j LOG --log-prefix "NETNS_RAW_OUTPUT "

```

Notice that iptables rules on the host aren't affected by the settings from the namespace:

请注意，主机上的 iptables 规则不受命名空间中设置的影响：

```bash
# run on host:

$ iptables -S -t filter
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT

$ iptables -S -t nat
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT
-A POSTROUTING -o eth0 -j MASQUERADE

$ iptables -S -t mangle
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT

$ iptables -S -t raw
-P PREROUTING ACCEPT
-P OUTPUT ACCEPT

```

Update iptables rules on the host:

更新主机上的 iptables 规则：

```bash
# run on host:

$ iptables -t filter -A INPUT -j LOG --log-prefix "HOST_FILTER_INPUT "
$ iptables -t filter -A FORWARD -j LOG --log-prefix "HOST_FILTER_FORWARD "
$ iptables -t filter -A OUTPUT -j LOG --log-prefix "HOST_FILTER_OUTPUT "

$ iptables -t nat -A PREROUTING -j LOG --log-prefix "HOST_NAT_PREROUTE "
$ iptables -t nat -A INPUT -j LOG --log-prefix "HOST_NAT_INPUT "
$ iptables -t nat -A OUTPUT -j LOG --log-prefix "HOST_NAT_OUTPUT "
$ iptables -t nat -A POSTROUTING -j LOG --log-prefix "HOST_NAT_POSTROUTE "

$ iptables -t mangle -A PREROUTING -j LOG --log-prefix "HOST_MANGLE_PREROUTE "
$ iptables -t mangle -A INPUT -j LOG --log-prefix "HOST_MANGLE_INPUT "
$ iptables -t mangle -A FORWARD -j LOG --log-prefix "HOST_MANGLE_FORWARD "
$ iptables -t mangle -A OUTPUT -j LOG --log-prefix "HOST_MANGLE_OUTPUT "
$ iptables -t mangle -A POSTROUTING -j LOG --log-prefix "HOST_MANGLE_POSTROUTE "

$ iptables -t raw -A PREROUTING -j LOG --log-prefix "HOST_RAW_PREROUTE "
$ iptables -t raw -A OUTPUT -j LOG --log-prefix "HOST_RAW_OUTPUT "

```

And finally ping _8.8.8.8_ from the namespace while `tailf`-ing kernel messages on the host:

最后从命名空间ping _8.8.8.8_，同时在主机上执行`tailf`-ing 内核消息：

```bash
# run in network namespace:

$ ping 8.8.8.8

```

```bash
# run on host:

$ tail -f /var/log/messages

```

Now if we start pinging an external address, like _8.8.8.8_, we can notice an interesting pattern arising in the netfilter logs:

现在，如果我们开始 ping 一个外部地址，比如 _8.8.8.8_，我们可以注意到 netfilter 日志中出现了一个有趣的模式：

```bash
Jun 21 13:25:19 localhost kernel: NETNS_RAW_OUTPUT IN= OUT=vGUEST SRC=172.16.0.2 DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: NETNS_MANGLE_OUTPUT IN= OUT=vGUEST SRC=172.16.0.2 DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: NETNS_FILTER_OUTPUT IN= OUT=vGUEST SRC=172.16.0.2 DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: NETNS_MANGLE_POSTROUTE IN= OUT=vGUEST SRC=172.16.0.2 DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_RAW_PREROUTE IN=br0 OUT= MAC=c2:96:cf:97:f4:12:c2:31:a8:8b:d7:f8:08:00 SRC=172.16.0.2DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_MANGLE_PREROUTE IN=br0 OUT= MAC=c2:96:cf:97:f4:12:c2:31:a8:8b:d7:f8:08:00 SRC=172.16.0.2DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_MANGLE_FORWARD IN=br0 OUT=eth0 MAC=c2:96:cf:97:f4:12:c2:31:a8:8b:d7:f8:08:00 SRC=172.16.0.2 DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=63 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_FILTER_FORWARD IN=br0 OUT=eth0 MAC=c2:96:cf:97:f4:12:c2:31:a8:8b:d7:f8:08:00 SRC=172.16.0.2 DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=63 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_MANGLE_POSTROUTE IN= OUT=eth0 SRC=172.16.0.2 DST=8.8.8.8 LEN=84 TOS=0x00 PREC=0x00 TTL=63 ID=2089 DF PROTO=ICMP TYPE=8 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_RAW_PREROUTE IN=eth0 OUT= MAC=52:54:00:26:10:60:52:54:00:12:35:02:08:00 SRC=8.8.8.8DST=10.0.2.15 LEN=84 TOS=0x00 PREC=0x00 TTL=62 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_MANGLE_PREROUTE IN=eth0 OUT= MAC=52:54:00:26:10:60:52:54:00:12:35:02:08:00 SRC=8.8.8.8DST=10.0.2.15 LEN=84 TOS=0x00 PREC=0x00 TTL=62 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_MANGLE_FORWARD IN=eth0 OUT=br0 MAC=52:54:00:26:10:60:52:54:00:12:35:02:08:00 SRC=8.8.8.8 DST=172.16.0.2 LEN=84 TOS=0x00 PREC=0x00 TTL=61 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_FILTER_FORWARD IN=eth0 OUT=br0 MAC=52:54:00:26:10:60:52:54:00:12:35:02:08:00 SRC=8.8.8.8 DST=172.16.0.2 LEN=84 TOS=0x00 PREC=0x00 TTL=61 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: HOST_MANGLE_POSTROUTE IN= OUT=br0 SRC=8.8.8.8 DST=172.16.0.2 LEN=84 TOS=0x00 PREC=0x00 TTL=61 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: NETNS_RAW_PREROUTE IN=vGUEST OUT= MAC=c2:31:a8:8b:d7:f8:c2:96:cf:97:f4:12:08:00 SRC=8.8.8.8DST=172.16.0.2 LEN=84 TOS=0x00 PREC=0x00 TTL=61 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: NETNS_MANGLE_PREROUTE IN=vGUEST OUT= MAC=c2:31:a8:8b:d7:f8:c2:96:cf:97:f4:12:08:00 SRC=8.8.8.8DST=172.16.0.2 LEN=84 TOS=0x00 PREC=0x00 TTL=61 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: NETNS_MANGLE_INPUT IN=vGUEST OUT= MAC=c2:31:a8:8b:d7:f8:c2:96:cf:97:f4:12:08:00 SRC=8.8.8.8DST=172.16.0.2 LEN=84 TOS=0x00 PREC=0x00 TTL=61 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34
Jun 21 13:25:19 localhost kernel: NETNS_FILTER_INPUT IN=vGUEST OUT= MAC=c2:31:a8:8b:d7:f8:c2:96:cf:97:f4:12:08:00 SRC=8.8.8.8DST=172.16.0.2 LEN=84 TOS=0x00 PREC=0x00 TTL=61 ID=17376 DF PROTO=ICMP TYPE=0 CODE=0 ID=3197 SEQ=34

```

The logs are pretty verbose, but try to focus on the log prefixes only. The pattern is as follows:

日志非常冗长，但尽量只关注日志前缀。图案如下：

```bash
NETNS_RAW_OUTPUT
NETNS_MANGLE_OUTPUT
NETNS_FILTER_OUTPUT
NETNS_MANGLE_POSTROUTE

HOST_RAW_PREROUTE
HOST_MANGLE_PREROUTE
HOST_MANGLE_FORWARD
HOST_FILTER_FORWARD
HOST_MANGLE_POSTROUTE
HOST_RAW_PREROUTE
HOST_MANGLE_PREROUTE
HOST_MANGLE_FORWARD
HOST_FILTER_FORWARD
HOST_MANGLE_POSTROUTE

NETNS_RAW_PREROUTE
NETNS_MANGLE_PREROUTE
NETNS_MANGLE_INPUT
NETNS_FILTER_INPUT

```

From this, we can get a rough idea about the chains precedence. Note, that while the namespace in our example behaves like a normal client machine sending requests to the Internet through its default route, the host serves a router's role:

由此，我们可以大致了解链的优先级。请注意，虽然我们示例中的命名空间的行为类似于普通客户端计算机通过其默认路由向 Internet 发送请求，但主机充当路由器的角色：

![Chains precedence on client](http://iximiuz.com/laymans-iptables-101/tables-precedence.png)

_Chains precedence on client._

_客户端的链优先级。_

![Chains precedence on router](http://iximiuz.com/laymans-iptables-101/tables-precedence-route.png)

_Chains precedence on router._

_路由器上的链优先级。_

## Conclusion

##  结论

It might look like iptables is an ancient technology. Does it make sense to spend your time learning it? But have a look at Docker or Kubernetes - booming bleeding edge products. Both heavily utilize iptables under the hood to set up and manage their networking layers! Don't be fooled, without learning fundamental things like netfilter, iptables, IPVS it will be never possible neither develop nor operate modern cluster management tools on a serious scale.

iptables 看起来可能是一项古老的技术。花时间学习它有意义吗？但是看看 Docker 或 Kubernetes——蓬勃发展的前沿产品。两者都在后台大量使用 iptables 来设置和管理他们的网络层！不要被愚弄，如果不学习诸如 netfilter、iptables、IPVS 之类的基本知识，就不可能大规模开发或运行现代集群管理工具。

Make code, not war!

编写代码，而不是战争！

[linux,](javascript: void 0) [networking,](javascript: void 0) [netfilter,](javascript: void 0) [iptables](javascript: void 0)

[linux,](javascript: void 0) [networking,](javascript: void 0) [netfilter,](javascript: void 0) [iptables](javascript: void 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

