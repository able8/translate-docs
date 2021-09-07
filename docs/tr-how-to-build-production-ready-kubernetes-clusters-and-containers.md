# How to Build Production-Ready Kubernetes Clusters and Containers

# 如何构建生产就绪的 Kubernetes 集群和容器

[Robert Stark](http://www.stackrox.com/authors/rstark/) May 09, 2019



Kubernetes is a powerful tool for building highly scalable systems. As a result, many companies have begun, or are planning, to use it to orchestrate production services. Unfortunately, like most powerful technologies, Kubernetes is complex. How do you know you’ve set things up correctly and it’s safe to flip the switch and open the network floodgates to your services? We’ve compiled the following checklist to help you prepare your containers and kube clusters for production traffic.

Kubernetes 是构建高度可扩展系统的强大工具。因此，许多公司已经开始或正在计划使用它来编排生产服务。不幸的是，与大多数强大的技术一样，Kubernetes 很复杂。您怎么知道您已经正确设置了一切，并且可以安全地打开开关并打开通往您的服务的网络闸门？我们编制了以下清单，以帮助您为生产流量准备容器和 kube 集群。

### Containers Done Right

### 容器做得对

Kubernetes provides a way to orchestrate containerized services, so if you don’t have your containers in order, your cluster isn’t going to be in good shape from the get go. Follow these tips to start out on the right foot.

Kubernetes 提供了一种编排容器化服务的方法，因此如果您没有按顺序排列容器，您的集群从一开始就不会处于良好状态。按照这些提示从右脚开始。

##### Use Minimal Base Images

##### 使用最少的基础图像

**What:** Containers are application stacks built into a system image. Everything from your business logic to the kernel gets packed inside. Minimal images strip out as much of the OS as possible and force you to explicitly add back any components you need.

**内容：** 容器是内置于系统映像中的应用程序堆栈。从业务逻辑到内核的所有内容都打包在里面。最少的图像尽可能多地剥离操作系统，并强制您明确添加回您需要的任何组件。

**Why:** Including in your container only the software you intend to use has both performance and security benefits. You have fewer bytes on disk, less network traffic for images being copied, and fewer tools for potential attackers to access.

**原因：** 在您的容器中仅包含您打算使用的软件具有性能和安全优势。磁盘上的字节更少，复制图像的网络流量更少，可供潜在攻击者访问的工具更少。

**How:** [Alpine](https://alpinelinux.org/) Linux is a popular choice and has broad support.

**如何：** [Alpine](https://alpinelinux.org/) Linux 是一种流行的选择并得到广泛的支持。

## Top 9 Kubernetes Security Best Practices

## 9 大 Kubernetes 安全最佳实践

Follow these 9 practical recommendations today to enhance your Kubernetes security

立即遵循这 9 条实用建议来增强您的 Kubernetes 安全性

##### Registries

##### 注册表

**What:** Registries are repositories for images, making those images available for download and launch. When you specify your deployment configuration, you’ll need to specify where to get the image with a path/::

**内容：** 注册表是图像的存储库，使这些图像可供下载和启动。当您指定部署配置时，您需要使用路径/:: 指定获取图像的位置

```
apiVersion: v1
kind: Deployment
...
spec:
...
containers:
  - name: app
    image: docker.io/app-image:version1

```


**Why:** Your cluster needs images to run.

**原因：** 您的集群需要映像才能运行。

**How:** Most cloud providers offer private image registry services: Google offers the [Google Container Registry](https://cloud.google.com/container-registry/), AWS provides [Amazon ECR](https://aws.amazon.com/ecr/), and Microsoft has the [Azure Container Registry](https://azure.microsoft.com/en-us/services/container-registry/).

**如何：** 大多数云提供商提供私有镜像注册服务：Google 提供 [Google Container Registry](https://cloud.google.com/container-registry/)，AWS 提供 [Amazon ECR](https:///aws.amazon.com/ecr/)，Microsoft 有 [Azure Container Registry](https://azure.microsoft.com/en-us/services/container-registry/)。

Do your homework, and choose a private registry that offers the best uptime. Since your cluster will rely on your registry to launch newer versions of your software, any downtime will prevent updates to running services.

做你的功课，并选择一个提供最佳正常运行时间的私人注册表。由于您的集群将依赖您的注册表来启动更新版本的软件，因此任何停机时间都会阻止对正在运行的服务进行更新。

##### ImagePullSecrets



**What:** ImagePullSecrets are Kubernetes objects that let your cluster authenticate with your registry, so the registry can be selective about who is able to download your images.

**内容：** ImagePullSecrets 是 Kubernetes 对象，可让您的集群通过您的注册表进行身份验证，因此注册表可以选择谁可以下载您的图像。

**Why:** If your registry is exposed enough for your cluster to pull images from it, then it’s exposed enough to need authentication.

**原因：** 如果您的注册表暴露得足以让您的集群从中提取图像，那么它就暴露得足以需要身份验证。

**How:** The Kubernetes website has a good walkthrough on configuring ImagePullSecrets, which uses Docker as an example registry, [here](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#log-in-to-docker).

**如何：** Kubernetes 网站有一个关于配置 ImagePullSecrets 的很好的演练，它使用 Docker 作为示例注册表，[这里](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#log-in-to-docker)。

### Organizing Your Cluster

### 组织您的集群

Microservices by nature are a messy business. A lot of the benefit of using microservices comes from enforcing separation of duties at a service level, effectively creating abstractions for the various components of your backend. Some good examples are running a database separate from business logic, running separate development and production versions of software, or separating out horizontally scalable processes.

微服务本质上是一项混乱的业务。使用微服务的很多好处来自在服务级别强制执行职责分离，有效地为后端的各种组件创建抽象。一些很好的例子是运行与业务逻辑分离的数据库，运行单独的软件开发和生产版本，或者分离出水平可扩展的流程。

The dark side of having different services performing different duties is that they cannot be treated as equals. Thankfully Kubernetes gives you many tools to deal with this problem.

让不同服务执行不同职责的阴暗面是它们不能被平等对待。幸运的是，Kubernetes 为您提供了许多工具来处理这个问题。

##### Namespaces

##### 命名空间

**What:** Namespaces are the most basic and most powerful grouping mechanism in Kubernetes. They work almost like virtual clusters. Most objects in Kubernetes are, by default, limited to affecting a single namespace at a time. 

**内容：** 命名空间是 Kubernetes 中最基本、最强大的分组机制。它们几乎像虚拟集群一样工作。默认情况下，Kubernetes 中的大多数对象一次只能影响一个命名空间。

**Why:** Most objects are namespace scoped, so you’ll have to use namespaces. Given that they provide strong isolation, they are perfect for isolating environments with different purposes, such as user serving production environments and those used strictly for testing, or to separate different service stacks that support a single application, like for instance keeping your security solution's workloads separate from your own applications. A good rule of thumb is to divide namespaces by resource allocation: If two sets of microservices will require different resource pools, place them in separate namespaces.

**原因：** 大多数对象都是命名空间范围的，因此您必须使用命名空间。鉴于它们提供强大的隔离，它们非常适合隔离具有不同目的的环境，例如用户服务生产环境和严格用于测试的环境，或者分离支持单个应用程序的不同服务堆栈，例如保持安全解决方案的工作负载与您自己的应用程序分开。一个好的经验法则是按资源分配划分命名空间：如果两组微服务需要不同的资源池，请将它们放在不同的命名空间中。

**How:** It’s part of the metadata of most object types:

**How:** 它是大多数对象类型的元数据的一部分：

```
apiVersion: v1
kind: Deployment
metadata:
name: example-pod
namespace: app-prod
...

```


Note that you should always create your own namespaces instead of relying on the ‘default’ namespace. Kubernetes’ defaults typically optimize for the lowest amount of friction for developers, and this often means forgoing even the most basic security measures.

请注意，您应该始终创建自己的命名空间，而不是依赖于“默认”命名空间。 Kubernetes 的默认设置通常会针对开发人员的最小摩擦进行优化，这通常意味着甚至要放弃最基本的安全措施。

##### Labels

##### 标签

**What:** Labels are the most basic and extensible way to organize your cluster. They allow you to create arbitrary key:value pairs that separate your Kubernetes objects. For instance, you might create a label key which separates services that handle sensitive information from those that do not.

**内容：** 标签是组织集群的最基本和可扩展的方式。它们允许您创建任意键：值对来分隔 Kubernetes 对象。例如，您可以创建一个标签键，将处理敏感信息的服务与不处理敏感信息的服务分开。

**Why:** As mentioned, Kubernetes uses labels for organization, but, more specifically, they are used for _selection_. This means, when you want to give a Kubernetes object a reference to a group of objects in some namespace, like telling a network policy which services are allowed to communicate with each other, you use their labels. Since they represent such an open ended type of organization, do your best to keep things simple, and only create labels where you require the power of _selection._

**原因：** 如前所述，Kubernetes 使用标签进行组织，但更具体地说，它们用于 _selection_。这意味着，当您想为 Kubernetes 对象提供对某个命名空间中的一组对象的引用时，例如告诉网络策略允许哪些服务相互通信，您可以使用它们的标签。由于它们代表了这种开放式组织类型，因此请尽量保持简单，并且仅在需要 _selection._ 功能的地方创建标签。

**How:**  Labels are a simple spec field you can add to your YAML files:

**如何：**标签是一个简单的规范字段，您可以添加到您的 YAML 文件中：

```
apiVersion: v1
kind: Deployment
metadata:
name: example-pod
...
matchLabels:
    userexposed: true
    storespii: true

```


##### Annotations

##### 注释

**What:** Annotations are arbitrary key-value metadata you can attach to your pods, much like labels. However, Kubernetes does not read or handle annotations, so the rules around what you can and cannot annotate a pod with are fairly liberal, and they can’t be used for selection.

**内容：** 注释是可以附加到 Pod 的任意键值元数据，很像标签。但是，Kubernetes 不读取或处理注释，因此关于可以和不可以对 Pod 进行注释的规则相当宽松，并且不能用于选择。

**Why:** They help you track certain important features of your containerized applications, like version numbers or dates and times of first bring up. Annotations, in the context of Kubernetes alone, are a fairly powerless construct, but they can be an asset to your developers and operations teams when used to track important system changes.

**原因：** 它们可帮助您跟踪容器化应用程序的某些重要功能，例如版本号或首次启动的日期和时间。仅在 Kubernetes 的上下文中，注释是一种相当无能的构造，但是当用于跟踪重要的系统更改时，它们可以成为您的开发人员和运营团队的资产。

**How:** Annotation are a spec field similar to labels.

**How：** Annotation 是一个类似于标签的规范字段。

```
apiVersion: v1
kind: Pod
metadata:
name: example-pod
...
annotations:
    version: four
    launchdate: tuesday

```


### Securing Your Cluster

### 保护您的集群

Alright, you’ve got a cluster set up and organized the way you want - now what? Well, next thing is getting some security in place. You could spend your whole lifetime studying and still not discover all the ways someone can break into your systems. A blog post has a lot less room for content than a lifetime, so you’ll have to settle for a couple of strong suggestions.

好的，你已经建立了一个集群并按照你想要的方式组织了 - 现在呢？好吧，接下来要做的是采取一些安全措施。您可以花费一生的时间来学习，但仍然无法发现有人可以闯入您的系统的所有方式。一篇博文的内容空间比一生都少，所以你必须接受一些强烈的建议。

##### RBAC

**What:** RBAC (Role Based Access Control) allows you to control who can view or modify different aspects of your cluster.

**内容：** RBAC（基于角色的访问控制）允许您控制谁可以查看或修改集群的不同方面。

**Why:** If you want to follow the principle of least privilege, then you need to have RBAC set up to limit what your cluster users, and your deployments, are able to do.

**为什么：** 如果您想遵循最小权限原则，那么您需要设置 RBAC 以限制您的集群用户和您的部署能够执行的操作。

**How:** If you're setting up your own cluster (ie, not using a managed Kube service), make sure you are using ''–authorization-mode=Node,RBAC" to launch your kube apiserver. If you are using a managed Kubernetes instance, you can check that it is set up to use RBAC by querying the command used to start the kube apiserver. The only generic way to check is to look for “–authorization-mode…” in the output of `kubectl cluster-info dump`. 

**如何：** 如果您正在设置自己的集群（即，不使用托管 Kube 服务），请确保您使用“–authorization-mode=Node,RBAC”来启动您的 kube apiserver。如果您正在使用托管 Kubernetes 实例，您可以通过查询用于启动 kube apiserver 的命令来检查它是否设置为使用 RBAC。唯一通用的检查方法是在输出中查找“–authorization-mode...” `kubectl 集群信息转储`。

Once RBAC is turned on, you’ll need to change the default permissions to suit your needs. The Kubernetes project site provides a walk-through on setting up Roles and RoleBindings [here](https://kubernetes.io/docs/reference/access-authn-authz/rbac/). Managed Kubernetes services require custom steps for enabling RBAC - check out Google's guide for [GKE](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control) or Amazon's instructions for [AKS](https://docs.microsoft.com/en-us/azure/aks/concepts-identity#role-based-access-controls-rbac).

打开 RBAC 后，您需要更改默认权限以满足您的需要。 Kubernetes 项目站点提供了有关设置角色和角色绑定的演练 [此处](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)。托管 Kubernetes 服务需要自定义步骤来启用 RBAC - 查看 Google 的 [GKE] 指南(https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control) 或亚马逊的说明对于 [AKS](https://docs.microsoft.com/en-us/azure/aks/concepts-identity#role-based-access-controls-rbac)。

##### Pod Security Policies

##### Pod 安全策略

**What:** Pod Security Policies are a resource, much like a Deployment or a Role, and can be created and updated through kubectl in same way. Each holds a collection of flags you can use to prevent specific unsafe behaviors in your cluster.

**内容：** Pod 安全策略是一种资源，很像 Deployment 或 Role，可以通过 kubectl 以同样的方式创建和更新。每个都包含一组标志，您可以使用这些标志来防止集群中的特定不安全行为。

**Why:** If the people who created Kubernetes thought limiting these behaviors was important enough to create a special object to handle it, then they are likely important.

**为什么：**如果创建 Kubernetes 的人认为限制这些行为足够重要以创建一个特殊的对象来处理它，那么他们可能很重要。

**How:** Getting them working can be an exercise in frustration. I recommend getting RBAC up and running, then check out the guide from the Kubernetes project [here](https://kubernetes.io/docs/concepts/policy/pod-security-policy/). The most important to use, in my opinion, are preventing [privileged](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#privileged) containers, and write access to the [host file system ](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#volumes-and-file-systems), as these represent some of the leakier parts of the container abstraction.

**如何：**让他们工作可能是一种沮丧的练习。我建议启动并运行 RBAC，然后查看来自 Kubernetes 项目的指南 [此处](https://kubernetes.io/docs/concepts/policy/pod-security-policy/)。在我看来，最重要的是防止[privileged](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#privileged) 容器，以及对 [host 文件系统的写访问](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#volumes-and-file-systems)，因为这些代表了容器抽象中一些更容易泄漏的部分。

##### Network Policies

##### 网络策略

**What:** Network policies are objects that allow you to explicitly state which traffic is permitted, and Kubernetes will block all other non-conforming traffic.

**内容：** 网络策略是允许您明确声明允许哪些流量的对象，Kubernetes 将阻止所有其他不符合规定的流量。

**Why:** Limiting network traffic in your cluster is a basic and important security measure. Kubernetes by default enables open communication between all services. Leaving this “default open” configuration in place means an Internet-connected service is just one hop away from a database storing sensitive information.

**原因：** 限制集群中的网络流量是一项基本且重要的安全措施。默认情况下，Kubernetes 启用所有服务之间的开放通信。保留此“默认打开”配置意味着连接 Internet 的服务距离存储敏感信息的数据库仅一跳。

**How:** A colleague of mine did a great write up that will get you going [here](https://www.cncf.io/blog/2019/04/19/setting-up-kubernetes-network-policies-a-detailed-guide/).

**如何：**我的一位同事写了一篇很棒的文章，可以让你[这里](https://www.cncf.io/blog/2019/04/19/setting-up-kubernetes-network-政策详细指南/)。

##### Secrets

##### 

**What:** Secrets are how you store sensitive data in Kubernetes, including passwords, certificates, and tokens.

**内容：** 秘密是您在 Kubernetes 中存储敏感数据的方式，包括密码、证书和令牌。

**Why:** Your services may need to authenticate one another, other third-party services, or your users, whether you’re implementing TLS or restricting access.

**原因：**无论您是实施 TLS 还是限制访问，您的服务都可能需要相互验证、其他第三方服务或您的用户。

**How:** The Kubernetes project offers a guide [here](https://kubernetes.io/docs/concepts/configuration/secret/). One key piece of advice: avoid loading secrets as environment variables, since having secret data in your environment is a general security no-no. Instead, mount secrets into read only volumes in your container - you can find an example in this [Using Secrets](https://kubernetes.io/docs/concepts/configuration/secret/#using-secrets) write up.

**如何：** Kubernetes 项目提供了指南 [此处](https://kubernetes.io/docs/concepts/configuration/secret/)。一个关键建议：避免将机密加载为环境变量，因为在您的环境中拥有机密数据是一般安全禁忌。相反，将机密安装到容器中的只读卷中 - 您可以在这篇 [Using Secrets](https://kubernetes.io/docs/concepts/configuration/secret/#using-secrets) 文章中找到一个示例。

##### Scanners

##### 

**What:** Scanners inspect the components installed in your images. Everything from the OS to your application stack. Scanners are super useful for finding out what vulnerabilities exist in the versions of software your image contains.

**内容：** 扫描仪会检查图像中安装的组件。从操作系统到应用程序堆栈的一切。扫描仪对于找出您的图像包含的软件版本中存在哪些漏洞非常有用。

**Why:** Vulnerabilities are discovered in popular open source packages all the time. Some notable examples are Heartbleed and Shellshock. You’ll want to know where such vulnerabilities reside in your system, so you know what images may need updating.

**原因：** 一直在流行的开源包中发现漏洞。一些值得注意的例子是 Heartbleed 和 Shellshock。您会想知道此类漏洞在您的系统中的位置，以便您知道哪些映像可能需要更新。

**How:** Scanners are a fairly common bit of infrastructure - most cloud providers have an offering. If you want to host something yourself, the open source [Clair](https://github.com/coreos/clair) project is a popular choice.

**如何：** 扫描仪是一种相当常见的基础设施 - 大多数云提供商都有提供。如果你想自己托管一些东西，开源 [Clair](https://github.com/coreos/clair) 项目是一个流行的选择。

### Keeping Your Cluster Stable 

### 保持集群稳定

Kubernetes represents a tall stack. You have your applications, running on baked-in kernels, running in VMs (or on bare metal in some cases), accompanied by Kubernetes’ own services sharing hardware. Given all these elements, plenty of things can go wrong, both in the physical and virtual realms, so it is very important to de-risk your development cycle wherever possible. The ecosystem around Kubernetes has developed a great set of best practices to keep things in line as much as possible.

Kubernetes 代表了一个高大的堆栈。您的应用程序在内置内核上运行，在 VM 中运行（或在某些情况下在裸机上运行），并伴随着 Kubernetes 自己的服务共享硬件。考虑到所有这些因素，很多事情都可能出错，无论是在物理领域还是虚拟领域，因此尽可能降低开发周期的风险非常重要。围绕 Kubernetes 的生态系统已经开发出一套出色的最佳实践，以尽可能保持一致。

##### CI/CD

##### 

**What:** Continuous Integration/Continuous Deployment is a process philosophy. It is the belief that every modification committed to your codebase should add incremental value and be production ready. So, if something in your codebase changes, you probably want to launch a new version of your service, either to run tests or to update your exposed instances.

**内容：** 持续集成/持续部署是一种流程哲学。我们相信，对您的代码库进行的每一次修改都应该增加增量价值并做好生产准备。因此，如果您的代码库中的某些内容发生更改，您可能希望启动服务的新版本，以运行测试或更新您公开的实例。

**Why:** Following CI/CD helps your engineering team keep quality in mind in their day-to-day work. If something breaks, fixing it becomes an immediate priority for the whole team, because every change thereafter, relying on the broken commit, will also be broken.

**原因：** 遵循 CI/CD 可帮助您的工程团队在日常工作中牢记质量。如果出现问题，修复它成为整个团队的当务之急，因为此后依赖于损坏提交的每个更改也将被破坏。

**How:** Thanks to the rise of cloud deployed software, CI/CD is in vogue. As a result, you can choose from tons of great offerings, from managed to self-hosted. If you’re a small team, I recommend going the managed route, as the time and effort you save is definitely worth the extra cost.

**如何：** 由于云部署软件的兴起，CI/CD 很流行。因此，您可以从大量优质产品中进行选择，从托管到自托管。如果你是一个小团队，我建议你走托管路线，因为你节省的时间和精力绝对值得额外的成本。

##### Canary

##### 金丝雀

**What:** Canary is a way of bringing service changes from a commit in your codebase to your users. You bring up a new instance running your latest version, and you migrate your users to the new instance slowly, gaining confidence in your updates over time, as opposed to swapping over all at once.

**内容：** Canary 是一种将服务更改从代码库中的提交带给用户的方式。您启动了一个运行最新版本的新实例，然后缓慢地将用户迁移到新实例，随着时间的推移获得对更新的信心，而不是一次全部交换。

**Why:** No matter how extensive your unit and integration tests are, they can never completely simulate running in production - there’s always the chance something will not function as intended. Using canary limits your users’ exposure to these issues.

**为什么：**无论您的单元测试和集成测试有多么广泛，它们都无法完全模拟生产中的运行 - 总是有可能无法按预期运行。使用 Canary 限制了用户接触这些问题的机会。

**How:** Kubernetes, as extensible as it is, provides many routes to incrementally roll out service updates. The most straightforward approach is to create a separate deployment that shares a load balancer with currently running instances. The idea is you scale up the new deployment while scaling down the old until all running instances are of the new version.

**如何：** Kubernetes 尽管具有可扩展性，但提供了许多途径来逐步推出服务更新。最直接的方法是创建一个单独的部署，与当前运行的实例共享负载均衡器。这个想法是你扩大新部署，同时缩小旧部署，直到所有正在运行的实例都是新版本。

##### Monitoring

##### 监控

**What:** Monitoring means tracking and recording what your services are doing.

**内容：** 监控意味着跟踪和记录您的服务正在做什么。

**Why:** Let’s face it - no matter how great your developers are, no matter how hard your security gurus furrow their brows and mash keys, things will go wrong. When they do, you’re going to want to know what happened to ensure you don’t make the same mistake twice.

**为什么：**让我们面对现实吧 - 无论您的开发人员多么出色，无论您的安全专家多么努力地皱着眉头和混搭密钥，事情都会出错。当他们这样做时，你会想知道发生了什么，以确保你不会犯两次同样的错误。

**How:** There are two steps to successfully monitor a service - the code needs to be instrumented, and the output of that instrumentation needs to be fed somewhere for storage, retrieval, and analysis. How you perform instrumentation is largely dependent on your toolchain, but a quick web search should give you somewhere to start. As far as storing the output goes, I recommend using a managed SIEM (like [Splunk](https://www.splunk.com/) or [Sumo Logic](https://www.sumologic.com/)) unless you have specialized knowledge or need - in my experience, DIY is always 10X the time and effort you expect when it comes to anything storage related.

**如何：** 成功监控服务有两个步骤 - 需要对代码进行检测，并且需要将该检测的输出提供给某处以进行存储、检索和分析。您如何执行检测在很大程度上取决于您的工具链，但快速的网络搜索应该会给您提供一个开始的地方。就存储输出而言，我建议使用托管 SIEM（如 [Splunk](https://www.splunk.com/) 或 [Sumo Logic](https://www.sumologic.com/))，除非您有专门的知识或需要 - 根据我的经验，在涉及任何存储相关的事情时，DIY 总是您期望的时间和精力的 10 倍。

### Advanced Topics

### 高级主题

Once your clusters reach a certain size, you’ll find enforcing all of your best practices manually becomes impossible, and the safety and stability of your systems will be challenged as a result. After you cross this threshold, consider the following topics:

一旦您的集群达到一定规模，您会发现手动执行所有最佳实践变得不可能，并且您的系统的安全性和稳定性将因此受到挑战。跨过此阈值后，请考虑以下主题：

##### Service Meshes

##### 服务网格

**What:** Services meshes are a way to manage your interservice communications, effectively creating a virtual network that you use when implementing your services.

**内容：** 服务网格是一种管理服务间通信的方法，可有效创建您在实现服务时使用的虚拟网络。

**Why:** Using a service mesh can alleviate some of the more tedious aspects of managing a cluster, such as ensuring communications are properly encrypted. 

**原因：** 使用服务网格可以减轻管理集群的一些较繁琐的方面，例如确保正确加密通信。

**How:** Depending on your choice of service mesh, getting up and running can vary wildly in complexity. [Istio](https://istio.io/) seems to be gaining momentum as the most used service mesh, and your configuration process will largely depend on your workloads.

**如何：**根据您选择的服务网格，启动和运行的复杂性可能会有很大差异。 [Istio](https://istio.io/) 作为最常用的服务网格似乎正在获得动力，您的配置过程将在很大程度上取决于您的工作负载。

A word of warning: If you expect to need a service mesh down the line, go through the agony of setting it up earlier rather than later - incrementally changing communication styles within a cluster can be a huge pain.

一句话警告：如果您希望在生产线上需要一个服务网格，那么请早点而不是晚点进行设置——在集群内逐步改变通信方式可能是一个巨大的痛苦。

##### Admission Controllers

##### 入场控制器

**What:** Admission controllers are a great catch-all tool for managing what’s going into your cluster. They allow you to set up webhooks that Kubernetes will consult during bring up. They come in two flavors: _Mutating_ and _Validating_. Mutating admission controllers alter the configuration of the deployment before it is launched. Validating admission controllers confer with your webhooks that a given deployment is allowed to be launched.

**内容：** 准入控制器是一个很好的通用工具，用于管理进入集群的内容。它们允许您设置 Kubernetes 将在启动期间参考的 Webhook。它们有两种风格：_Mutating_ 和 _Validating_。变异准入控制器会在部署启动之前更改其配置。验证准入控制器与您的 webhooks 授予允许启动给定部署的权限。

**Why:** Their use cases are broad and numerous - they provide a great way to iteratively improve your cluster’s stability with home-grown logic and restrictions.

**原因：**它们的用例广泛且数量众多 - 它们提供了一种很好的方法，可以通过自有逻辑和限制来迭代地提高集群的稳定性。

**How:** Check out [this](https://kubernetes.io/blog/2019/03/21/a-guide-to-kubernetes-admission-controllers/) great guide on how to get started with Admission Controllers.

**如何：**查看 [this](https://kubernetes.io/blog/2019/03/21/a-guide-to-kubernetes-admission-controllers/) 入门指南控制器。

* * *

* * *

Categories:

类别：

- [Kubernetes](http://www.stackrox.com/categories/kubernetes)
- [Containers](http://www.stackrox.com/categories/containers) 

- [Kubernetes](http://www.stackrox.com/categories/kubernetes)
- [容器](http://www.stackrox.com/categories/containers)

