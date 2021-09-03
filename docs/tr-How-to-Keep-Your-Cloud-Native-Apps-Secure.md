# How to Keep Your Cloud-Native Apps Secure

# 如何确保您的云原生应用程序安全

June 28, 2021

An increasing number of organizations are adopting cloud-native apps, and new research from [Snyk](https://snyk.io/state-of-cloud-native-application-security/) reveals that along with that adoption, 60 percent of organizations have increased security concerns.

越来越多的组织正在采用云原生应用程序，来自 [Snyk](https://snyk.io/state-of-cloud-native-application-security/) 的新研究表明，随着这种应用的采用，60%的组织增加了安全问题。

We got together with Jörg Müller, principal consultant at INNOQ, to hear his tips on how to help with the security and operational challenges companies often face when adopting cloud-native apps.

我们与 INNOQ 的首席顾问 Jörg Müller 会面，听取他关于如何帮助公司应对在采用云原生应用程序时经常面临的安全和运营挑战的技巧。

## **How do you secure the network traffic?**

## **您如何保护网络流量？**

**You should encrypt everything in your environment, but you need to choose on what level. If you have encryption on the highest level, you might not need additional encryption lower down that would just cost you performance.**

**您应该加密环境中的所有内容，但您需要选择加密级别。如果您有最高级别的加密，您可能不需要额外的低级加密，这只会降低您的性能。**

In the past you had something that was called security-on-perimeter, which meant that you had a firewall around your internal network and everything inside your network was trusted. Services could communicate to each other without encryption and without authorization.

过去，您有一种称为外围安全的东西，这意味着您的内部网络周围有防火墙，并且您网络中的所有内容都是可信的。服务可以在没有加密和未经授权的情况下相互通信。

This was never a very good practice, but it was a very common one, and for most people it was enough. Today it’s not, and there are two main reasons for that:

1\. It’s not that complicated to have encrypted authorized network traffic inside your own network these days, and there’s no performance overhead you have to take care of

这从来都不是一个很好的做法，但这是一个非常普遍的做法，对大多数人来说已经足够了。今天不是，主要有两个原因：

1\.如今，在您自己的网络中加密授权的网络流量并不复杂，而且您无需处理任何性能开销

2\. Zero trust networks are a far better security practice, which is used today

2。零信任网络是一种更好的安全实践，今天使用

### There’s no reason for you to take an on-perimeter security approach

### 您没有理由采取外围安全方法

It’s interesting to look at where you encrypt authorized network traffic inside your own network, and on what level. If you’re looking at a modern deployment of an application, then the network stack alone has several layers:

看看您在自己的网络中对授权网络流量进行加密的位置以及加密级别很有趣。如果您正在查看应用程序的现代部署，那么仅网络堆栈就有几个层：

**Level 1: The physical network**

**第 1 级：物理网络**

The one where the cables are running and where the physical network card is found. Some network cards already have the capabilities to run encryption on this level.

电缆所在的位置和物理网卡所在的位置。一些网卡已经有能力在这个级别上运行加密。

**Level 2: Layer of virtualization**

**第 2 级：虚拟化层**

Above the physical network you usually have a layer of virtualization, which means that the virtual machines have their own IP addresses that don’t match the addresses of any network cards, and the communication between the virtual machines might already be encrypted and authorized.

在物理网络之上通常有一层虚拟化，这意味着虚拟机拥有自己的 IP 地址，与任何网卡的地址都不匹配，并且虚拟机之间的通信可能已经被加密和授权。

**Level 3: Overlay network**

**第 3 级：覆盖网络**

If you’re using Kubernetes, you’ll have an overlay network for Kubernetes which is doing the same again. If you’re using Cilium or something similar, then you’re doing encryption and authorization there again.

如果您正在使用 Kubernetes，您将拥有一个用于 Kubernetes 的覆盖网络，它再次执行相同的操作。如果你使用 Cilium 或类似的东西，那么你又在那里进行加密和授权。

**Level 4: Service mesh**

**第 4 级：服务网格**

You get a level higher if you’re using service meshes — you have an encryption on the service level, running encryption between services.

如果您使用服务网格，您将获得更高的级别——您在服务级别进行了加密，在服务之间运行加密。

Finding the right combination of that stack in your company is complicated.

在您的公司中找到该堆栈的正确组合很复杂。

It might be enough if you're using an encrypted overlay network on a Kubernetes layer, or it might be necessary to have something like a service mesh to encrypt between services and authorize on that level because your applications need to know what your other services are communicating with me and need to authorize them, which is something you can't do on the lower levels because you're missing the knowledge there.

如果您在 Kubernetes 层上使用加密覆盖网络可能就足够了，或者可能需要使用服务网格之类的东西来在服务之间加密并在该级别上进行授权，因为您的应用程序需要知道您的其他服务是什么与我沟通并需要授权他们，这是您在较低级别无法做的事情，因为您缺少那里的知识。

## **Container isolation**

## **容器隔离**

**Where do you get the containers? And are they secure? What privileges do they need? What are they containing, and are they containing any vulnerabilities?**

**你从哪里得到容器？它们安全吗？他们需要什么特权？它们包含什么，它们是否包含任何漏洞？**

If you’re talking about cloud native, let’s presume you’re using containers, since it’s usually the way people deploy today.

如果您在谈论云原生，让我们假设您使用的是容器，因为它通常是当今人们部署的方式。

From a security standpoint, a container does not run in a complete virtualization but usually as an isolated process. It's a process on your operating system, on your kernel, that is isolated from the other processes, isolated from the network level, isolated on the process level itself (so it can't see the other ones), and it's not allowed to do anything.

从安全的角度来看，容器不是在完整的虚拟化中运行，而是通常作为一个孤立的进程运行。它是您操作系统上的一个进程，在您的内核上，与其他进程隔离，与网络级别隔离，在进程级别本身隔离（因此它看不到其他进程），并且不允许这样做任何事物。

But sometimes it needs to do some things, and sometimes there are some privileges that slip through because they can call the kernel, and if you don't configure the container in the right way, then you have some issues that can break out of your container. 

但是有时它需要做一些事情，有时会漏掉一些特权，因为它们可以调用内核，如果您没有以正确的方式配置容器，那么您就会遇到一些问题容器。

There are mechanisms in Kubernetes like pod security policies where Kubernetes controls what features or privileges a container tries to get and can prevent a container from just starting up.

Kubernetes 中有一些机制，例如 pod 安全策略，其中 Kubernetes 控制容器尝试获取哪些功能或权限，并可以防止容器刚刚启动。

This is something you need to take care of, otherwise you run the risk of running a rogue container from taking over something in your network, having some privileges that you don’t want it to have.

这是你需要注意的事情，否则你会冒着运行流氓容器的风险来接管网络中的某些东西，拥有一些你不希望它拥有的特权。

That’s even worse if you’re using containers that are not inside your organization, so that’s another level of security you have to take care of: where are the containers coming from and are they secure?

如果你使用的容器不在你的组织内部，那就更糟了，所以这是你必须注意的另一个安全级别：容器来自哪里，它们安全吗？

### VM for every container

### 每个容器的VM

If you really want really strong container isolation, you can use a container runtime that builds a virtual machine for every container.

如果你真的想要真正强大的容器隔离，你可以使用容器运行时为每个容器构建一个虚拟机。

There are container runtimes that create new virtual machines, or that are very strict about the privileges, adding another layer of security to a Kubernetes cluster, which helps you a lot in keeping things isolated. This is a very high-level approach.

有容器运行时可以创建新的虚拟机，或者对权限非常严格，为 Kubernetes 集群添加了另一层安全性，这有助于您保持事物的隔离。这是一种非常高级的方法。

### Scanning containers

### 扫描容器

There’s a lot of stuff you can do around automatically scanning a container before you even deploy it. That’s a big advantage compared to the old way of setting up a virtual machine and maintaining it where you would need to run your own updates. Today, you can just build the container and run it through a scanner.

在部署容器之前，您可以做很多事情来自动扫描容器。与设置虚拟机并在需要运行自己的更新的地方维护它的旧方法相比，这是一个很大的优势。今天，您可以构建容器并通过扫描仪运行它。

Docker has some stuff built in, but there are plenty of scanners out there which will run your container against the CVE database and tell you if the container you’re using contains any software with known vulnerabilities. Once scanned you can then integrate that into your CD process, and you’ll have a much more secure container rollout and content of your containers.

Docker 内置了一些东西，但有很多扫描器可以针对 CVE 数据库运行您的容器，并告诉您您使用的容器是否包含任何具有已知漏洞的软件。扫描完成后，您可以将其集成到您的 CD 流程中，您将拥有更安全的容器部署和容器内容。

This is completely different to how it was done in the past where the main thing was to ensure your virtual machine and operating system were both up to date. Today it’s a completely different approach of building a container, scanning a container and rolling it out.

这与过去的做法完全不同，过去主要是确保您的虚拟机和操作系统都是最新的。今天，它是一种完全不同的构建容器、扫描容器并将其推出的方法。

### Service mesh

### 服务网格

Service meshes usually do a thing which is called MTLS (mutual TLS). TLS is something you know from your normal web browser, if you’re using any website which has https then your browser can validate that the certificate for a website you’re accessing really is the certificate for this website through certificate chains.

服务网格通常会做一种称为 MTLS（双向 TLS）的事情。 TLS 是您从普通 Web 浏览器中知道的东西，如果您使用任何具有 https 的网站，那么您的浏览器可以通过证书链验证您正在访问的网站的证书是否确实是该网站的证书。

Mutual TLS does that in both directions. So the server will validate if you are who you say you are, which is usually not the case in normal web environments, but very useful if you're doing this in a serveless or Kubernetes environment because then the service that's providing you the service knows it's been called by someone who's allowed to call it and not someone else. This is a very high level of security and was unheard of before the days of cloud native.

双向 TLS 在两个方向上都做到了这一点。因此，服务器将验证您是否是您所说的人，这在正常的 Web 环境中通常不是这种情况，但是如果您在无服务或 Kubernetes 环境中执行此操作非常有用，因为这样为您提供服务的服务就知道了它是由允许调用它的人而不是其他人调用的。这是一种非常高的安全性，在云原生时代之前是闻所未闻的。

## **Speed of updating all infrastructure components**

## **更新所有基础设施组件的速度**

**One of the operational challenges is having to update all your components on a continuous basis.**

**运营挑战之一是必须持续更新所有组件。**

Kubernetes releases a new version every three months, Linkerd as a service mesh releases around every three months, and many other tools and products are doing a lot of high-frequency releases.

Kubernetes 每三个月发布一个新版本，Linkerd 作为服务网格大约每三个月发布一次，许多其他工具和产品都在做大量的高频发布。

In the past you usually had companies or system admins that tried to keep a very stable environment with a long-term supported OS version like RedHat, which was supported for eight years with regular updates. You were not updating your OS very often, and that was very normal.

过去，您通常有公司或系统管理员试图通过长期受支持的操作系统版本（如 RedHat，该版本支持八年并定期更新）来保持非常稳定的环境。您不经常更新操作系统，这很正常。

Today, it’s completely different. You have to update your Kubernetes cluster every three to six months. If you skip too many versions it will be very difficult for you to update them again.

今天，情况完全不同。您必须每三到六个月更新一次 Kubernetes 集群。如果您跳过太多版本，您将很难再次更新它们。

This stacked release cycle also leads to APIs deprecating very quickly — within just half a year to a year of a managed deprecation process for an API, the stuff you would have been using would now be gone.

这种堆叠的发布周期也导致 API 很快被弃用——在 API 的托管弃用过程的短短半年到一年内，您本应使用的东西现在将消失。

The update cycle these days runs very fast, but this can be seen as a good thing from a security perspective as you're forced to have the latest versions of software, you can't just go by with software that is four to five years old — a Kubernetes environment wouldn't even run.

这些天的更新周期运行得非常快，但从安全角度来看，这可以被视为一件好事，因为您被迫拥有最新版本的软件，您不能只使用四到五年的软件旧的——Kubernetes 环境甚至无法运行。

### Use the stuff that everybody else uses 

### 使用其他人使用的东西

You can manage a lot of that by not doing it yourself but letting your cloud provider do it. Use managed Kubernetes instances from Google, AWS, etc. You need a very good reason not to run a managed Kubernetes cluster.

您可以通过不自己做而让您的云提供商来管理很多事情。使用来自 Google、AWS 等的托管 Kubernetes 实例。您需要一个很好的理由不运行托管 Kubernetes 集群。

There might be reasons to do that in your own data center. Some reasons might be costs, but that should be one of the last reasons. Security or GDPR rules could be another reason, because most of the cloud providers are American in Europe.

在您自己的数据中心可能有理由这样做。一些原因可能是成本，但这应该是最后的原因之一。安全或 GDPR 规则可能是另一个原因，因为大多数云提供商在欧洲都是美国人。

In general, using the managed offerings is the way to go. It wasn’t the case two or three years ago — many customers still set up Kubernetes clusters of their own, they just used it and went on to updating things on their own.

一般来说，使用托管产品是可行的方法。两三年前情况并非如此——许多客户仍然建立自己的 Kubernetes 集群，他们只是使用它并继续自己更新。

## **Operational Issues**

## **操作问题**

### Dev team responsibilities

### 开发团队职责

More often than not, security is often not a very strong focus point in development teams. Bringing more management responsibilities into the development teams helps alleviate this problem.

通常情况下，安全性通常不是开发团队的重点。将更多的管理职责引入开发团队有助于缓解这个问题。

The conflict between stability from the operations team and the new features and progress from the development team itself is something that we try to solve in DevOps, but I think you can only solve that by working together and by having most of the responsibility with the development team itself, even for security.

运维团队的稳定性与开发团队本身的新功能和进步之间的冲突是我们试图在 DevOps 中解决的问题，但我认为您只能通过共同努力并承担大部分开发责任来解决这个问题团队本身，即使是为了安全。

This is something which should be done and requires a lot of training from the developers side. Not every developer wants that but it makes sense, and I know that there are a lot of developers out there that agree.

这是应该完成的事情，并且需要开发人员进行大量培训。并非每个开发人员都希望这样做，但这是有道理的，而且我知道有很多开发人员同意这一点。

### Clusters

### 集群

How large is your cluster? How are you approaching Kubernetes clusters?

你的集群有多大？您如何处理 Kubernetes 集群？

There are two extremes of ways to approach clusters:

接近集群的方法有两种：

1. Having one cluster per application, or even per server
2. 每个应用程序，甚至每个服务器都有一个集群

1. Having a large cluster which includes everything and tries to isolate stuff with namespaces

2. 拥有一个包含所有内容的大型集群，并尝试使用命名空间来隔离内容

You have to find a balance in your operational department, somewhere in between those two extremes. How many applications, how many teams are you? Are you accessing a single cluster? How many clusters do you have? Will you try to consolidate everything into one or many?

你必须在你的运营部门找到一个平衡点，介于这两个极端之间。你有多少个应用程序，多少个团队？您是否正在访问单个集群？你有多少集群？您是否会尝试将所有内容合并为一个或多个？

There are a lot of advantages and disadvantages on each side. 

每一方都有很多优点和缺点。


