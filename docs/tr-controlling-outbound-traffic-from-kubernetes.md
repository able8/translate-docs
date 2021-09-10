# Controlling outbound traffic from Kubernetes

# 控制来自 Kubernetes 的出站流量

[Read the article](http://monzo.com#article)

At Monzo, the Security Team's highest priority is to keep your money and data safe. And to achieve this, we're always adding and refining security controls across our banking platform.

在 Monzo，安全团队的首要任务是确保您的资金和数据安全。为实现这一目标，我们始终在我们的银行平台上添加和完善安全控制。

Late last year, we wrapped up a major networking project which let us control internal traffic in our platform (read about it [here](https://monzo.com/blog/we-built-network-isolation-for-1-500-services)). This gave us a lot of confidence that malicious code or an intruder compromising an individual microservice wouldn't be able to hurt our customers.

去年年底，我们完成了一个主要的网络项目，让我们可以控制我们平台中的内部流量（阅读[这里](https://monzo.com/blog/we-built-network-isolation-for-1-500-服务))。这给了我们很大的信心，即恶意代码或破坏单个微服务的入侵者不会伤害我们的客户。

Since then, we've been thinking about how we can add similar security to network traffic _leaving_ our platform. A lot of attacks begin with a compromised platform component 'phoning home' — that is, communicating with a computer outside of Monzo that is controlled by the attacker. Once this communication is established, the attacker can control the compromised service and attempt to infiltrate deeper into our platform. We knew that if we could block that communication, we'd stand a better chance at stopping an attack in its tracks.

从那时起，我们一直在考虑如何为离开我们平台的网络流量添加类似的安全性。许多攻击都是从一个受感染的平台组件“打电话回家”开始的——也就是说，与攻击者控制的 Monzo 之外的计算机进行通信。一旦建立了这种通信，攻击者就可以控制受感染的服务并尝试更深入地渗透到我们的平台中。我们知道，如果我们能够阻止这种通信，我们就有更好的机会阻止攻击。

![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/4SIwnE4zIp90oqeLFVRJLj/26b0213a3b8b7b83a3d49390757a6835/Untitled__2_.png?w=1280&q=90)

In this blog, we'll cover the journey to build our own solution for controlling traffic leaving our platform ( [you can read more about our platform here](https://monzo.com/blog/2016/09/19/building-a-modern-bank-backend)). This project formed part of our wider effort to move towards a platform with less trust, where applications need to be granted specific permissions instead of being allowed to do whatever they like. Our project contributed to this by allowing us to say that a specific service can communicate to an external provider like GitHub, but others can't.

在本博客中，我们将介绍构建我们自己的解决方案以控制离开我们平台的流量的过程（[您可以在此处阅读有关我们平台的更多信息](https://monzo.com/blog/2016/09/19/building-a-现代银行后端）)。该项目是我们向信任度较低的平台迈进的更广泛努力的一部分，在该平台中，应用程序需要被授予特定的权限，而不是被允许做任何他们喜欢做的事情。我们的项目为此做出了贡献，允许我们说特定服务可以与 GitHub 等外部提供商进行通信，但其他服务则不能。

Check out [other posts](https://monzo.com/blog/authors/jack-kleeman/) from Security Team for more projects we've completed to get closer to zero trust.

查看安全团队的 [其他帖子](https://monzo.com/blog/authors/jack-kleeman/)，了解我们为接近零信任而完成的更多项目。

## We started by identifying where traffic leaves the platform

## 我们首先确定流量离开平台的位置

With a few exceptions, we previously didn't have filtering on what outbound traffic could send. This means that we started in the dark about which services needed to talk to the internet and which didn't. This is similar to where we started our previous network isolation project, when we didn't to know what services talk to what other services.

除了少数例外，我们以前没有过滤出站流量可以发送的内容。这意味着我们一开始就不清楚哪些服务需要与互联网通信，哪些不需要。这类似于我们开始我们之前的网络隔离项目时，我们不知道哪些服务与哪些其他服务交谈。

But unlike when we started the network isolation project, tools and processes we built for that project are now at our disposal: a combination of [calico-accountant](https://github.com/monzo/calico-accountant) and packet logging allowed us to rapidly identify which of our microservices actually talk to the internet, and what external IPs each talks to.

但与我们开始网络隔离项目时不同的是，我们为该项目构建的工具和流程现在可供我们使用：[calico-accountant](https://github.com/monzo/calico-accountant) 和数据包日志的组合使我们能够快速识别我们的哪些微服务实际与互联网通信，以及每个微服务与哪些外部 IP 通信。

Even with IP information, it wasn't always trivial to find out what domains each service talked to. This is because many external providers our services talk to use CDNs like AWS CloudFront, where a single IP serves many websites. For simpler cases, we'd simply read the code for the service to find out what hostnames are used. If this isn't possible, we'd log the outgoing packets of the service carefully to identify their destinations.

即使有 IP 信息，找出每个服务与哪些域进行通信也并非易事。这是因为我们的服务与之交谈的许多外部提供商都使用 CDN，例如 AWS CloudFront，其中一个 IP 为多个网站提供服务。对于更简单的情况，我们只需阅读服务的代码以找出使用的主机名。如果这是不可能的，我们会仔细记录服务的传出数据包以确定它们的目的地。

Once we put together a (surprisingly small) spreadsheet on what services used what external hostnames and ports, we began a design process to think about what an ideal long term solution would look like.

一旦我们将有关哪些服务使用哪些外部主机名和端口的电子表格（令人惊讶地小）放在一起，我们就开始了一个设计过程，以考虑理想的长期解决方案应该是什么样子。

We had a few requirements: 

我们有几个要求：

- **We wanted to be able to write policies for outgoing traffic from specific source services to specific destination hostnames.** We have thousands of services running in our platform, each serving a small but distinctive purpose; and we cannot just have one universal list of allowed destinations, such as saying 'all services can talk to GitHub'. Some destinations on the internet used by our services would allow a certain degree of attacker control — anyone can host something on GitHub, for example. Therefore we want to limit 'dangerous' destinations to only the services that need them.

- **我们希望能够为从特定源服务到特定目标主机名的传出流量编写策略。**我们在我们的平台上运行了数千个服务，每个服务都有一个小而独特的目的；并且我们不能只有一个允许目的地的通用列表，例如说“所有服务都可以与 GitHub 对话”。我们的服务使用的互联网上的某些目的地允许一定程度的攻击者控制——例如，任何人都可以在 GitHub 上托管某些东西。因此，我们希望将“危险”目的地限制为仅需要它们的服务。

- **It must be possible to allow a specific DNS name.** Allowing traffic to IPs is super easy already with Calico — our Kubernetes networking stack — but over time IPs change for most DNS names. For a long-term solution, we need to be able to allow traffic to a domain name like `google.com`, and at any point in time thereafter, traffic to resolved IPs of `google.com`, should just work.

- **必须可以允许特定的 DNS 名称。** 使用 Calico（我们的 Kubernetes 网络堆栈）允许流量到 IP 已经非常容易了，但随着时间的推移，大多数 DNS 名称的 IP 都会发生变化。对于长期解决方案，我们需要能够允许流量流向像“google.com”这样的域名，并且在此后的任何时间点，流向“google.com”的已解析 IP 的流量应该可以正常工作。

- **We wanted to be able to reliably alert when we detect packets whose sources or destinations aren't allowed,** like we can for our internal network controls.


- **我们希望在检测到来源或目的地不被允许的数据包时能够可靠地发出警报，**就像我们的内部网络控制一样。


![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/6Oe5lLcMnqoUBZ3FQB3ihE/d1e3247c2695c9968a1cabb63bb593d1/Untitled__3_.png?w=1280&q=90)

We realised almost no drop-in solution on the market could check all the boxes, and it'd take a fairly long process to implement our ideal solution. So we decided to ship a simple solution first, and iterate on it.

我们意识到市场上几乎没有任何插入式解决方案可以满足所有条件，而且实现我们理想的解决方案需要相当长的过程。所以我们决定先发布一个简单的解决方案，然后对其进行迭代。

## We started with port based filtering

## 我们从基于端口的过滤开始

Without building anything new, we realised that with what's already running in our platform, we could implement simple rules on destination ports without support for hostnames. That's to say, we can make statements like this:

在不构建任何新内容的情况下，我们意识到利用我们平台中已经运行的内容，我们可以在不支持主机名的情况下在目标端口上实施简单的规则。也就是说，我们可以这样声明：

_"Allow service.github to make outbound connections to the public internet on port 443"_

_“允许 service.github 在端口 443 上建立到公共互联网的出站连接”_

The service will be able to reach any domain on that port, which isn't a very tight control. But more importantly, services that shouldn't be making any outbound connections at all won't be allowed to. The vast majority of our services don't need to talk to the public internet, so we actually get a ton of benefit from this simple approach. We also knew that a few services would always need to be allowed to reach a very large number of domains on certain ports, so it was useful to allow that use case early on.

该服务将能够访问该端口上的任何域，这不是一个非常严格的控制。但更重要的是，根本不应该建立任何出站连接的服务将不被允许。我们的绝大多数服务不需要与公共互联网对话，因此我们实际上从这种简单的方法中获得了很多好处。我们还知道，始终需要允许一些服务访问某些端口上的大量域，因此尽早允许该用例很有用。

The way we implemented this was through a Kubernetes NetworkPolicy for each port that we needed to allow. If a pod of a service is labelled with `external-egress.monzo.com/443: true`, then it will match a policy which allows traffic to all public IPs on port 443. Any traffic not allowed by one of these policies is logged (a feature of Calico), then dropped.

我们实现这一点的方式是通过 Kubernetes NetworkPolicy 为我们需要允许的每个端口。如果一个服务的 pod 被标记为“external-egress.monzo.com/443: true”，那么它将匹配一个策略，该策略允许端口 443 上的所有公共 IP 的流量。这些策略之一不允许的任何流量是记录（Calico 的一个功能），然后删除。

To label pods, we used a similar approach to the one we used for internal network controls: 'rule files' which are part of a service's code. To start with, we added files like `service.github/manifests/egress/external/443.rule`, and updated our deployment pipeline to automatically convert these files into the right labels.

为了标记 pod，我们使用了一种类似于我们用于内部网络控制的方法：“规则文件”，它是服务代码的一部分。首先，我们添加了诸如“service.github/manifests/egress/external/443.rule”之类的文件，并更新了我们的部署管道以自动将这些文件转换为正确的标签。

## We investigated existing solutions

## 我们调查了现有的解决方案

Having implemented a simple solution supporting only ports, we then set off looking for a full solution supporting hostname filtering. There are two approaches that we saw in the open source and enterprise software communities:

实现了一个仅支持端口的简单解决方案后，我们开始寻找支持主机名过滤的完整解决方案。我们在开源和企业软件社区中看到了两种方法：

### SNI and host header inspection 

### SNI 和主机头检查

One approach, used by [Istio](https://istio.io/), is to run an egress proxy inside Kubernetes. This egress proxy would inspect the traffic sent to it, figure out the destination, determine whether the destination was allowed (given the source) and then pass it on. To determine destinations for encrypted traffic, it has to use the SNI ( [Server Name Indication](https://en.wikipedia.org/wiki/Server_Name_Indication)) field for encrypted (TLS) traffic. For unencrypted web traffic — which is becoming increasingly rare — it can simply inspect the HTTP Host header.

[Istio](https://istio.io/) 使用的一种方法是在 Kubernetes 内运行出口代理。这个出口代理将检查发送给它的流量，找出目的地，确定目的地是否被允许（给定源），然后将其传递。要确定加密流量的目的地，它必须对加密 (TLS) 流量使用 SNI（[服务器名称指示](https://en.wikipedia.org/wiki/Server_Name_Indication))字段。对于越来越少的未加密的网络流量，它可以简单地检查 HTTP 主机标头。

Because Istio is simply configuring Envoy to achieve this functionality, and Monzo already had a fairly advanced Envoy service mesh, we were confident we could build something very similar. And because this solution would run inside Kubernetes, we could enforce policies per service based on their source pod IPs, which is more challenging to do outside Kubernetes, as we run a virtual network for pods. We could also set up alerting for rogue packets fairly easily, as we'd control the application which determines what's allowed.

因为 Istio 只是简单地配置 Envoy 来实现这个功能，而 Monzo 已经有一个相当先进的 Envoy 服务网格，我们有信心我们可以构建非常相似的东西。因为这个解决方案将在 Kubernetes 内部运行，我们可以根据源 pod IP 为每个服务实施策略，这在 Kubernetes 外部执行更具挑战性，因为我们为 pod 运行虚拟网络。我们还可以很容易地为恶意数据包设置警报，因为我们可以控制决定允许什么的应用程序。

![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/4G7BCJKxdfei8vM8dDatKm/9cac731554e252e39f242a2ff43c2d07/Untitled__4_.png?w=1280&q=90)

But this proxy would only work for TLS or HTTP traffic. Monzo's platform runs many services, with lots of different protocols to talk to external providers. In particular, we deal with quite a lot of outbound SFTP traffic, which wouldn't be supported by this approach. We weren't sure if it would've been possible to modify Envoy to support determining SSH destinations in a reliable way.

但此代理仅适用于 TLS 或 HTTP 流量。 Monzo 的平台运行许多服务，使用许多不同的协议与外部供应商进行对话。特别是，我们处理了相当多的出站 SFTP 流量，这种方法不支持这些流量。我们不确定是否可以修改 Envoy 以支持以可靠的方式确定 SSH 目的地。

Some enterprise solutions take a similar approach to Istio. Instead of a proxy, they'll set themselves up as the NAT gateway for all outbound traffic in your AWS cluster, then inspect traffic to determine what's allowed. Again, we didn't expect SSH to be supported.

一些企业解决方案采用与 Istio 类似的方法。他们将自己设置为 AWS 集群中所有出站流量的 NAT 网关，而不是代理，然后检查流量以确定允许的内容。同样，我们没想到会支持 SSH。

### DNS inspection

### DNS 检查

A more common approach to domain name based filtering is based on the observation that, for an application to reach an IP for `github.com`, it must have (at some point) resolved `github.com` via DNS (DNS is how computers turn domain names into IP addresses). Well behaved applications will do this regularly, so that their IPs are never older than the expiry of the DNS record. As a result, by watching DNS traffic, we can figure out key information: firstly, that a given application wants to reach `github.com`, and secondly that it will try to do so on a specific set of IPs.

一种更常见的基于域名的过滤方法是基于以下观察：对于一个应用程序来访问 `github.com` 的 IP，它必须（在某些时候）通过 DNS 解析 `github.com`（DNS 是如何计算机将域名转换为 IP 地址）。表现良好的应用程序将定期执行此操作，因此它们的 IP 永远不会早于 DNS 记录的到期时间。因此，通过观察 DNS 流量，我们可以找出关键信息：首先，给定的应用程序想要访问 `github.com`，其次，它将尝试在一组特定的 IP 上这样做。

Many implementations of egress filtering use this principle to allow (for a limited period of time) traffic to the returned IPs, but only for the source IP that requested them. We can choose to only allow this for a certain set of domain names, effectively allowing outgoing traffic to these domains only. Our firewall will then have a set of dynamically changing rules. Critically, our firewall must update itself to accept traffic to the new IPs _before_ these IPs are returned to the application — otherwise subsequent requests from the application could fail.

出口过滤的许多实现都使用此原则来允许（在有限的时间段内）返回 IP 的流量，但仅限于请求它们的源 IP。我们可以选择只允许特定的一组域名这样做，有效地只允许传出流量到这些域。我们的防火墙将拥有一组动态变化的规则。至关重要的是，我们的防火墙必须自我更新以接受流向新 IP 的流量_在这些 IP 返回到应用程序之前 - 否则来自应用程序的后续请求可能会失败。

We found quite a few enterprise solutions that use this approach and would sit as a box inside our platform. Generally, the box would carry all outbound traffic, and in some cases it would also act as a DNS resolver, or otherwise outbound DNS traffic must all pass through it. It would then be able to adjust firewall rules for newly resolved IPs as they are returned in DNS responses.

我们找到了很多使用这种方法的企业解决方案，它们可以作为我们平台内的一个盒子。通常，该框将承载所有出站流量，并且在某些情况下，它还充当 DNS 解析器，否则出站 DNS 流量必须全部通过它。然后，它可以在 DNS 响应中返回新解析的 IP 时调整防火墙规则。

![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/4qahu4EL7RPvOQujSwXeqt/04127ba6942ebdb17c2e5a8dbe650459/Untitled__5_.png?w=1280&q=90)



The main issue was that the enterprise solutions couldn't run inside Kubernetes. So they wouldn't be able to let specific services access specific destinations, as the source IP would just look like the IP of a cloud instance running Kubernetes, since our cluster runs a virtual network for pods. They will however allow us to control egress from our non-Kubernetes workloads in the same way; but in most cases we can control that traffic just as easily with IP-based rules on AWS.

主要问题是企业解决方案无法在 Kubernetes 中运行。因此，他们无法让特定服务访问特定目的地，因为源 IP 看起来就像运行 Kubernetes 的云实例的 IP，因为我们的集群为 pod 运行虚拟网络。然而，它们将允许我们以相同的方式控制来自非 Kubernetes 工作负载的出口；但在大多数情况下，我们可以使用 AWS 上基于 IP 的规则轻松控制流量。

### **Building our own egress firewall**

### **构建我们自己的出口防火墙**

We considered how complicated it'd be to build a similar approach. We'd need to create an application inside Kubernetes that we routed all outbound traffic through, as well as DNS traffic. Our main concern was availability - to make this application reliable, we'd want to run many instances of it, with outgoing traffic balanced evenly between instances.

我们考虑过构建类似方法会有多复杂。我们需要在 Kubernetes 中创建一个应用程序，我们通过它路由所有出站流量以及 DNS 流量。我们主要关心的是可用性——为了使这个应用程序可靠，我们希望运行它的许多实例，在实例之间均匀平衡传出流量。

But this creates a state consistency problem where when one instance discovers a new IP that needs to be allowed, it'd need to very rapidly communicate this to other instances - ideally, before the DNS query completes. This is possible, but adds a lot of complexity to the system.

但这会造成状态一致性问题，当一个实例发现需要允许的新 IP 时，它需要非常快速地将此信息传达给其他实例 - 理想情况下，在 DNS 查询完成之前。这是可能的，但会增加系统的复杂性。

Another approach, taken by the enterprise solutions, would be to use a single 'leader' which handles all egress traffic, but one which can fail over to a set of 'followers', ideally waiting for state to transfer in the process. But since we can't redirect routed traffic in our cloud environment instantly, a short but significant delay would be inevitable before the gateway's available again, leading to potential failures in applications that need to reach external services.

企业解决方案采用的另一种方法是使用单个“领导者”来处理所有出口流量，但可以故障转移到一组“追随者”，理想情况下等待状态在过程中转移。但是，由于我们无法立即重定向云环境中的路由流量，因此在网关再次可用之前不可避免地会出现短暂但显着的延迟，从而导致需要访问外部服务的应用程序出现潜在故障。

## We came up with a hybrid approach

## 我们提出了一种混合方法

We saw the merits of both approaches:

我们看到了两种方法的优点：

- Egress proxies which carry outgoing traffic are easy to run in a highly available way, but inspection doesn't work universally

- 携带传出流量的出口代理很容易以高度可用的方式运行，但检查并不普遍

- DNS inspection is a fairly universal way to figure out where traffic is going, but it's tricky to build a reliable distributed system to take advantage of that


- DNS 检查是确定流量去向的一种相当普遍的方法，但是构建可靠的分布式系统来利用它是很棘手的


We had an idea: **what if we just have a proxy for** **_each_** **domain we needed to reach?** This might sound complex and wasteful, but we can easily configure and run these proxies inside our Kubernetes cluster at scale. We call these proxies _egress gateways_.

我们有一个想法：**如果我们只有一个代理** **_each_** **我们需要访问的域？** 这可能听起来很复杂和浪费，但我们可以轻松地在我们的内部配置和运行这些代理大规模的 Kubernetes 集群。我们称这些代理为 _egress gateways_。

For example, there's a gateway for `github.com`, with a specific load balancing internal IP address. This gateway accepts traffic and always sends it to `github.com`, doing its own DNS resolution. The gateway just operates on TCP (it is a [Layer 4](https://en.wikipedia.org/wiki/Transport_layer)proxy), so Git over SSH is supported as well as almost everything else.

例如，`github.com` 有一个网关，具有特定的负载均衡内部 IP 地址。此网关接受流量并始终将其发送到 `github.com`，进行自己的 DNS 解析。网关仅在 TCP 上运行（它是 [第 4 层](https://en.wikipedia.org/wiki/Transport_layer)代理)，因此支持 Git over SSH 以及几乎所有其他内容。

To actually get the traffic flowing through these gateways, we'd need to "hijack" DNS responses for domains they proxy to: when a microservice asks for `github.com`, we need to respond with the correct gateway IP address, instead of the real IP addresses of `github.com`. This'll cause this service to use our gateway to talk to `github.com` automatically, without any changes to the client.

为了真正让流量流过这些网关，我们需要“劫持”它们代理的域的 DNS 响应：当微服务请求 `github.com` 时，我们需要使用正确的网关 IP 地址进行响应，而不是`github.com` 的真实 IP 地址。这将导致此服务使用我们的网关自动与 `github.com` 对话，而无需对客户端进行任何更改。

So once all services start using gateways exclusively for their outgoing traffic, we can block all traffic to the public internet, except from the gateways. We consider it safer to have many uniform TCP proxies able to reach the public internet than a diverse set of microservices. And more importantly, we can control which applications can talk to the gateway, using our existing internal network isolation controls, taking advantage of the same tools, the same alerting, and the same understanding we already have.

因此，一旦所有服务开始将网关专门用于其传出流量，我们就可以阻止所有到公共 Internet 的流量，但来自网关的流量除外。我们认为拥有许多能够访问公共互联网的统一 TCP 代理比一组多样化的微服务更安全。更重要的是，我们可以控制哪些应用程序可以与网关通信，使用我们现有的内部网络隔离控制，利用我们已经拥有的相同工具、相同警报和相同理解。

![Controlling outbound traffic from Kubernetes](https://images.ctfassets.net/ro61k101ee59/2Ovajc6p7Cpn9Bc41uZtkI/51f5a73a7475f55d5e6ce1aeac2cdb99/Untitled__6_.png?w=1280&q=90)

### Implementing egress gateways in Kubernetes

### 在 Kubernetes 中实现出口网关

We manage each egress gateway in our cluster with some Kubernetes building blocks:

我们使用一些 Kubernetes 构建块来管理集群中的每个出口网关：

1. **A** [**Deployment**](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) **of** [**Envoy**](https://www.envoyproxy.io/) **pods.** We usually run three pods for high availability. The envoy pods use a local file to get their configuration. 

1. **A** [**部署**](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) **of** [**Envoy**](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) **pods.** 我们通常运行三个 pod 以实现高可用性。 Envoy Pod 使用本地文件来获取它们的配置。

2. **A** [**ConfigMap**](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) **to configure Envoy.** The config determines what ports the gateway listens on and what DNS name traffic is destined for (eg `github.com`).

2. **A** [**ConfigMap**](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) **配置Envoy。** config确定网关侦听的端口以及 DNS 名称流量的目的地（例如`github.com`)。

3. **A** [**HorizontalPodAutoscaler**](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) **to increase the number of pods dynamically based on CPU usage.** TCP proxying is a good use case for horizontal scaling — lots of pods running in parallel usually handle concurrent connection loads as well as a few powerful pods.

3. **A** [**HorizontalPodAutoscaler**](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) **根据CPU动态增加pod数量用法。** TCP 代理是水平扩展的一个很好的用例——许多并行运行的 pod 通常处理并发连接负载以及一些强大的 pod。

4. **A** [**Service**](https://kubernetes.io/docs/concepts/services-networking/service/) **which gives the pods a single internal IP address.** This means that we can set up DNS to resolve `github.com` to a single internal IP, but each new connection is load balanced to different gateway pods.

4. **A** [**Service**](https://kubernetes.io/docs/concepts/services-networking/service/) **它为 pod 提供了一个单一的内部 IP 地址。** 这意味着我们可以设置 DNS 以将 `github.com` 解析为单个内部 IP，但是每个新连接都会负载均衡到不同的网关 pod。

5. **A** [**NetworkPolicy**](https://kubernetes.io/docs/concepts/services-networking/network-policies/) **which controls access to the gateways.** If the gateway is named `github`, then access to it is restricted to pods which are labelled with `egress.monzo.com/allowed-github: true`.


5. **A** [**NetworkPolicy**](https://kubernetes.io/docs/concepts/services-networking/network-policies/) **控制对网关的访问。** 如果网关名为`github`，然后访问它仅限于标记为`egress.monzo.com/allowed-github: true`的pod。


To set this all up, we just needed to set a few configuration files for each gateway. But we actually went one step further and automated the process of managing these objects with an [Kubernetes operator](https://coreos.com/blog/introducing-operators.html). This operator accepts a very small set of information (in this case the domain names that we need to allow), and creates the above resources. If we change how we want to set up those resources, all gateway configurations will be updated automatically. This lets us manage hundreds of gateways with minimal manual work.

为了设置这一切，我们只需要为每个网关设置几个配置文件。但我们实际上更进一步，使用 [Kubernetes 操作员](https://coreos.com/blog/introducing-operators.html)自动化了管理这些对象的过程。这个操作符接受一个非常小的信息集（在这种情况下我们需要允许的域名)，并创建上述资源。如果我们更改设置这些资源的方式，所有网关配置都将自动更新。这使我们能够以最少的手动工作来管理数百个网关。

We also needed to build a [CoreDNS](https://coredns.io/) plugin which allowed us to hijack DNS traffic for `github.com` and send it to the appropriate gateway. This turned out to be fairly easy — we just needed to [watch](https://kubernetes.io/docs/reference/using-api/api-concepts/#efficient-detection-of-changes) Kubernetes Service objects to figure out which gateways exist and for what domains, and then wrap the existing CoreDNS `rewrite` plugin to rewrite queries for `github.com` into a query for `github.egress-operator-system.svc.cluster.local`, which is the DNS name for the GitHub egress gateway.

我们还需要构建一个 [CoreDNS](https://coredns.io/) 插件，它允许我们劫持 `github.com` 的 DNS 流量并将其发送到适当的网关。结果证明这相当简单——我们只需要[观察](https://kubernetes.io/docs/reference/using-api/api-concepts/#efficient-detection-of-changes) Kubernetes Service 对象来计算找出存在哪些网关以及针对哪些域，然后将现有的 CoreDNS `rewrite` 插件包装起来，将 `github.com` 的查询重写为 `github.egress-operator-system.svc.cluster.local` 的查询，即GitHub 出口网关的 DNS 名称。

Both the operator and its CoreDNS plugin we built can be found open-source [here](https://github.com/monzo/egress-operator).

我们构建的运营商及其 CoreDNS 插件都可以在 [这里](https://github.com/monzo/egress-operator) 中找到开源。

From the perspective of the end user (an engineer building microservices in our cluster), a service can contain a file `service.github/manifests/egress/external/github.com:443.rule`, and our deployment pipeline will convert this into the appropriate pod label, by determining which egress gateway allows traffic to that destination.

从最终用户（在我们的集群中构建微服务的工程师）的角度来看，一个服务可以包含一个文件“service.github/manifests/egress/external/github.com:443.rule”，我们的部署管道会转换这个文件通过确定哪个出口网关允许流量到达该目的地，进入适当的 pod 标签。

From the perspective of the Security team maintaining the egress gateways, we manage a file like this for each destination:

从维护出口网关的安全团队的角度来看，我们为每个目的地管理一个这样的文件：

![apiVersion: egress.monzo.com/v1 kind: ExternalService metadata:   name: github spec:   dnsName: github.com   hijackDns: true   ports:   - port: 443](https://images.ctfassets.net/ro61k101ee59/2yaNN8oWWm5ZQVQvRl2X5T/fb56160ae22edeb589a6467c48036107/Screenshot_2020-04-06_at_12.16.55.png?w=1280&q=90)

/fb56160ae22edeb589a6467c48036107/Screenshot_2020-04-06_at_12.16.55.png?w=1280&q=90)

We've been enforcing our new firewall for some time now and we're really pleased with the outcome. Egress gateways provide us with:

一段时间以来，我们一直在执行我们的新防火墙，我们对结果非常满意。出口网关为我们提供：

- Very granular control of outgoing traffic to each external domain used by each service

- 对每个服务使用的每个外部域的传出流量进行非常精细的控制

- Great inspectability via Envoy as well as our existing network isolation tools

- 通过 Envoy 以及我们现有的网络隔离工具实现出色的可检查性

- The ability to use a single firewall (Kubernetes network policies) for both internal and external traffic.


- 能够为内部和外部流量使用单个防火墙（Kubernetes 网络策略）。


We can definitely see us keeping this approach as a long-term solution. 

我们绝对可以将这种方法视为长期解决方案。

