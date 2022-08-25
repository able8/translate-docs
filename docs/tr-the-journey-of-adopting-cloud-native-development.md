# The Journey of Adopting Cloud-Native Development

# 采用云原生开发的旅程

Jul 2, 2020

- [Level 0 – Traditional Development - Kubernetes is an ops topic](http://loft.sh#level-0--traditional-development---kubernetes-is-an-ops-topic)
   - [Advantages](http://loft.sh#advantages)
   - [Disadvantages](http://loft.sh#disadvantages)
- [Level 1 – Manual Continuous Deployment (CD) – Developers run CD pipelines to Kubernetes manually](http://loft.sh#level-1--manual-continuous-deployment-cd--developers-run-cd-pipelines-to-kubernetes-manually)
   - [Advantages](http://loft.sh#advantages-1)
   - [Disadvantages](http://loft.sh#disadvantages-1)
- [Level 2 – Cloud-native CD – Specialized CD tools run pipelines to Kubernetes automatically](http://loft.sh#level-2--cloud-native-cd--specialized-cd-tools-run-pipelines-to-kubernetes-automatically)
   - [Advantages](http://loft.sh#advantages-2)
   - [Disadvantages](http://loft.sh#disadvantages-2)
- [Level 3 – Cloud-native development – Development takes place inside the Kubernetes cluster](http://loft.sh#level-3--cloud-native-development--development-takes-place-inside-the-kubernetes-cluster)
   - [Advantages](http://loft.sh#advantages-3)
   - [Disadvantages](http://loft.sh#disadvantages-3)
- [A closer look at the access to Kubernetes](http://loft.sh#a-closer-look-at-the-access-to-kubernetes)
- [Conclusion](http://loft.sh#conclusion)

- [Level 0 - 传统开发 - Kubernetes 是一个 ops 主题](http://loft.sh#level-0--traditional-development---kubernetes-is-an-ops-topic)
  - [优势](http://loft.sh#advantages)
  - [缺点](http://loft.sh#disadvantages)
- [级别 1 – 手动持续部署 (CD) – 开发人员手动将 CD 管道运行到 Kubernetes](http://loft.sh#level-1--manual-continuous-deployment-cd--developers-run-cd-pipelines-to-kubernetes-手动)
  - [优势](http://loft.sh#advantages-1)
  - [缺点](http://loft.sh#disadvantages-1)
- [Level 2 – 云原生 CD – 专用 CD 工具自动运行管道到 Kubernetes](http://loft.sh#level-2--cloud-native-cd--specialized-cd-tools-run-pipelines-to-kubernetes-自动)
  - [优势](http://loft.sh#advantages-2)
  - [缺点](http://loft.sh#disadvantages-2)
- [第 3 级 - 云原生开发 - 在 Kubernetes 集群内进行开发](http://loft.sh#level-3--cloud-native-development--development-takes-place-inside-the-kubernetes-簇)
  - [优势](http://loft.sh#advantages-3)
  - [缺点](http://loft.sh#disadvantages-3)
-  [深入了解 Kubernetes 的访问](http://loft.sh#a-closer-look-at-the-access-to-kubernetes)
- [结论](http://loft.sh#conclusion)

Adopting Kubernetes is a process that many companies are currently going through. The introduction of Kubernetes as infrastructure technology can take some time. ( [It took almost 2 years for Tinder to complete its migration to Kubernetes](https://medium.com/tinder-engineering/tinders-move-to-kubernetes-cda2a6372f44).) The transition of development processes to fully cloud-native development is often an even longer process that comprises several incremental steps.

采用 Kubernetes 是许多公司目前正在经历的过程。将 Kubernetes 作为基础设施技术引入可能需要一些时间。 （[Tinder 花了将近 2 年时间完成向 Kubernetes 的迁移](https://medium.com/tinder-engineering/tinders-move-to-kubernetes-cda2a6372f44)。)开发流程向完全云的过渡-原生开发通常是一个更长的过程，包括几个增量步骤。

To illustrate the different stages of transition, I want to use the analogy of [autonomous driving](https://www.synopsys.com/automotive/autonomous-driving-levels.html), where different “levels” also describe the technological advancement and sophistication. Luckily and in contrast to autonomous driving, the highest level of cloud-native development is already reachable and reality today.

为了说明过渡的不同阶段，我想使用[自动驾驶](https://www.synopsys.com/automotive/autonomous-driving-levels.html)的类比，其中不同的“级别”也描述了技术先进性和成熟度。幸运的是，与自动驾驶相比，最高水平的云原生开发在今天已经可以实现并成为现实。

By looking at the following description of the different cloud-native levels, you can classify your current [development workflow with Kubernetes](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development) and even see what is ahead of you and what your next possible steps might be.

通过查看以下对不同云原生级别的描述，您可以对当前 使用 Kubernetes 的开发工作流 进行分类（http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium =reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development），甚至可以查看您前面的内容以及您接下来可能采取的步骤。

## [\#](http://loft.sh\#level-0--traditional-development---kubernetes-is-an-ops-topic) Level 0 – Traditional Development - Kubernetes is an ops topic

## [\#](http://loft.sh\#level-0--traditional-development---kubernetes-is-an-ops-topic) Level 0 - 传统开发 - Kubernetes 是一个 ops 主题

At level 0, [Kubernetes is not a topic for developers yet](http://loft.sh/blog/is-kubernetes-still-just-an-ops-topic/?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development) although it might already be used for production workloads. While developers may know that Kubernetes is used to run their applications after they are finished, they are not in touch with Kubernetes during their everyday work. They only use container solutions such as Docker, if they are working with containers at all.
This means that the developers are working on their local computers with their traditional technologies and hand-over the software for production to operations departments after completion.

在 0 级，[Kubernetes 还不是开发者的话题](http://loft.sh/blog/is-kubernetes-still-just-an-ops-topic/?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development)虽然它可能已经用于生产工作负载。虽然开发人员可能知道 Kubernetes 用于在完成后运行他们的应用程序，但他们在日常工作中并没有与 Kubernetes 接触。如果他们完全使用容器，他们只使用容器解决方案，例如 Docker。
这意味着开发人员正在使用他们的传统技术在本地计算机上工作，并在完成后将用于生产的软件移交给运营部门。

### [\#](http://loft.sh\#advantages) Advantages

### [\#](http://loft.sh\#advantages) 优点

Traditional development, as the name suggests, has been the standard for a long time in the past and even though it is “Level 0”, it does not only have downsides.

顾名思义，传统开发在过去很长一段时间都是标准，即使是“0级”，它也不仅有缺点。

The main advantage of this status quo is that **nothing must be changed**, so engineers do not have to learn new things and all **workflows can remain the same**.

这种现状的主要优势是**没有什么必须改变**，因此工程师不必学习新事物，所有**工作流程都可以保持不变**。

Another benefit is that **developers get direct feedback if their application is running** because everything runs on their local machines and nothing has to be executed in a Kubernetes environment first, so errors, for example, are visible pretty much immediately.

另一个好处是**如果他们的应用程序正在运行，开发人员会获得直接反馈**，因为一切都在他们的本地机器上运行，并且不需要首先在 Kubernetes 环境中执行任何操作，因此例如错误几乎可以立即看到。

### [\#](http://loft.sh\#disadvantages) Disadvantages 

### [\#](http://loft.sh\#disadvantages) 缺点

For modern development methods, level 0, however, has many disadvantages. At first, it is **contrary to the DevOps approach** by creating a clear separation of development and operations. As a result, **developers are not able to take on full responsibility** for their software and **“works on my machine”-problems** emerge very easily when the locally (without Kubernetes) developed software runs in the real- life Kubernetes environment for the first time.

然而，对于现代开发方法，0 级有很多缺点。起初，它**与 DevOps 方法相反**，通过创建明确的开发和运营分离。结果，**开发人员无法对他们的软件承担全部责任**，并且**“在我的机器上工作”——当本地（没有 Kubernetes）开发的软件在真实环境中运行时，问题**很容易出现——人生第一次使用 Kubernetes 环境。

This **lack of realism** of the runtime environment leads to a huge challenge for operations managers who are responsible for getting the application to run. This is aggravated by the fact that they **cannot easily replicate the developers’ local runtimes** for testing and debugging because they may be configured in many different ways.

运行时环境的这种**缺乏真实性**给负责让应用程序运行的运营经理带来了巨大的挑战。由于他们**无法轻松复制开发人员的本地运行时**以进行测试和调试，因此加剧了这种情况，因为它们可能以多种不同的方式进行配置。

Generally, only executing applications on individual local machines with different configurations is disadvantageous. It now is the responsibility of the **developers to set up the environment**, which may be **a lot of work** and **must be repeated** for every new team member or every new computer used.

通常，仅在具有不同配置的单个本地机器上执行应用程序是不利的。现在是**开发人员设置环境**的责任，这可能是**大量的工作**并且**必须为每个新团队成员或使用的每台新计算机重复**。

And even worse, the **execution in local runtime environments is simply not possible for all applications** anymore as they become more and more complex and may require very special or a lot of computing resources. This is especially true for machine learning and artificial intelligence software running on GPUs or requiring a lot of computing power.

更糟糕的是，**在本地运行时环境中执行根本不可能对所有应用程序**来说，因为它们变得越来越复杂，并且可能需要非常特殊或大量的计算资源。对于在 GPU 上运行或需要大量计算能力的机器学习和人工智能软件来说尤其如此。

## [\#](http://loft.sh\#level-1--manual-continuous-deployment-cd--developers-run-cd-pipelines-to-kubernetes-manually) Level 1 – Manual Continuous Deployment (CD) – Developers run CD pipelines to Kubernetes manually

## [\#](http://loft.sh\#level-1--manual-continuous-deployment-cd--developers-run-cd-pipelines-to-kubernetes-manually) 级别 1 – 手动持续部署(CD) – 开发人员手动将 CD 管道运行到 Kubernetes

Developers at level 1 know that Kubernetes is their target platform but do not have direct access to it. They are rather enabled to push their code to a Kubernetes runtime via a pre-defined pipeline that is usually set up by someone else, e.g. a DevOps engineer.

级别 1 的开发人员知道 Kubernetes 是他们的目标平台，但无法直接访问它。他们可以通过通常由其他人设置的预定义管道将代码推送到 Kubernetes 运行时，例如DevOps 工程师。

From a workflow perspective, developers push their code to a code repository and then manually trigger a pipeline with tools such as [Gitkube](https://github.com/hasura/gitkube) or [Spinnaker](https://github.com/spinnaker/spinnaker) that deploys the code in a Kubernetes environment, usually in a remote cluster in the cloud.

从工作流程的角度来看，开发人员将他们的代码推送到代码存储库，然后使用 [Gitkube](https://github.com/hasura/gitkube) 或 [Spinnaker](https://github.com) 等工具手动触发管道。 com/spinnaker/spinnaker)在 Kubernetes 环境中部署代码，通常在云中的远程集群中。

For the developers, this means that Kubernetes is the target platform but it is concealed by the pipeline.

对于开发人员来说，这意味着 Kubernetes 是目标平台，但它被管道隐藏了。

### [\#](http://loft.sh\#advantages-1) Advantages

### [\#](http://loft.sh\#advantages-1) 优点

A clear advantage of this approach is that the software is executed in a **very realistic environment in the cloud** already during development. The developers are so enabled to test their code in Kubernetes **without having to manage this environment as it is centrally controlled and maintained** by someone else.

这种方法的一个明显优势是软件在开发过程中已经在**非常真实的云环境**中执行。开发人员可以在 Kubernetes 中测试他们的代码**而无需管理此环境，因为它由其他人集中控制和维护**。

This makes it also **relatively easy to use** and **new developers can get started very fast** because they only need to get the right to trigger the pipeline for their code.

这也使它**相对易于使用**并且**新开发人员可以非常快速地开始**，因为他们只需要获得触发其代码管道的权利。

Another benefit of level 1 development is that **bugs and errors can be easily communicated and replicated** as everyone in the team works with the same CD environment and configurations instead of an individual local environment.

1 级开发的另一个好处是**错误和错误可以很容易地传达和复制**，因为团队中的每个人都使用相同的 CD 环境和配置，而不是单独的本地环境。

Finally, since the target environment is a Kubernetes cluster in the cloud, it is **possible to even run very complex and computing intense software** without limitations.

最后，由于目标环境是云中的 Kubernetes 集群，**甚至可以无限制地运行非常复杂和计算密集型的软件**。

### [\#](http://loft.sh\#disadvantages-1) Disadvantages

### [\#](http://loft.sh\#disadvantages-1) 缺点

The transition from level 0 to level 1 is substantial and the strengths of the one level are the weaknesses of the other. For level 1, this is particularly important for the development speed.

从 0 级到 1 级的过渡是实质性的，一个级别的优点是另一个级别的弱点。对于级别 1，这对于开发速度尤为重要。

While feedback is nearly immediate for level 0 local development, it becomes **very slow** for level 1 development based on manual pipelines. This is because the developers have to **manually execute the whole pipeline every time they change a file**, which can cause several minutes of pure waiting time for every change.

虽然 0 级本地开发的反馈几乎是即时的，但对于基于手动管道的 1 级开发，它变得**非常慢**。这是因为开发人员必须**每次更改文件时手动执行整个管道**，这可能会导致每次更改都需要几分钟的纯等待时间。

Additionally, since a shared and centrally managed platform is used, developers are also **not allowed to change or configure anything** themselves. They rather have to inform the cluster administrator which leads to **problematic single point of failure** if the admin is not available. 

此外，由于使用了共享和集中管理的平台，开发人员也**不允许自己更改或配置任何东西**。如果管理员不可用，他们宁愿通知集群管理员，这会导致**有问题的单点故障**。

Finally, in spite of the general flexibility of the runtime in terms of scale, the whole process is **not platform-independent**, so a transition from one environment to the other (e.g. from AWS to Azure) becomes a complicated process.

最后，尽管运行时在规模方面具有一般灵活性，但整个过程**不是独立于平台的**，因此从一种环境到另一种环境的过渡（例如从 AWS 到 Azure）成为一个复杂的过程。

## [\#](http://loft.sh\#level-2--cloud-native-cd--specialized-cd-tools-run-pipelines-to-kubernetes-automatically) Level 2 – Cloud-native CD – Specialized CD tools run pipelines to Kubernetes automatically

## [\#](http://loft.sh\#level-2--cloud-native-cd--specialized-cd-tools-run-pipelines-to-kubernetes-automatically) 级别 2 – 云原生CD – 专门的 CD 工具自动运行到 Kubernetes 的管道

The development concept at level 2 is essentially the same as at level 1. The code is deployed via a pipeline to a Kubernetes cluster. However, at this stage, developers are using special CLI tools, such as [Skaffold](https://github.com/GoogleContainerTools/skaffold), [Draft](https://github.com/Azure/draft), or [Tilt](https://github.com/tilt-dev/tilt) that detect the file changes by the developer and then automatically trigger the pipelines.

级别 2 的开发概念与级别 1 基本相同。代码通过管道部署到 Kubernetes 集群。但是，在这个阶段，开发人员正在使用特殊的 CLI 工具，例如 [Skaffold](https://github.com/GoogleContainerTools/skaffold)、[Draft](https://github.com/Azure/draft)或[Tilt](https://github.com/tilt-dev/tilt) 由开发人员检测文件更改，然后自动触发管道。

Since these tools are specifically made for this use case, they also have additional features that facilitate development with Kubernetes, such as port forwarding.

由于这些工具是专门为这个用例而设计的，它们还具有促进使用 Kubernetes 进行开发的附加功能，例如端口转发。

Another very important distinction of this level is that developers have direct access to Kubernetes for the first time. While it is not strictly necessary, the standard case for the Kubernetes access with these tools is to use a local Kubernetes cluster, i.e. a Kubernetes cluster started with tools such as [minikube](https://github.com/kubernetes/minikube) , [kind](https://github.com/kubernetes-sigs/kind) or [MicroK8s](https://github.com/ubuntu/microk8s) on the local computer of the developer.

这个级别的另一个非常重要的区别是开发人员第一次可以直接访问 Kubernetes。虽然并非绝对必要，但使用这些工具访问 Kubernetes 的标准情况是使用本地 Kubernetes 集群，即使用 [minikube](https://github.com/kubernetes/minikube) 等工具启动的 Kubernetes 集群, [kind](https://github.com/kubernetes-sigs/kind) 或 [MicroK8s](https://github.com/ubuntu/microk8s) 在开发者的本地计算机上。

### [\#](http://loft.sh\#advantages-2) Advantages

### [\#](http://loft.sh\#advantages-2) 优点

Similar to the pipeline process of level 1, development with cloud-native CD pipelines also leads to tests in an **environment close to production** due to Kubernetes as underlying technology.

与级别 1 的流水线流程类似，由于 Kubernetes 作为底层技术，使用云原生 CD 流水线进行开发也会导致在接近生产环境的**环境中进行测试。

This also lets other team members **replicate problems and bugs relatively easily** as everyone is at least using the same infrastructure technology, even if it runs locally with some minor differences.

这也让其他团队成员**相对容易地复制问题和错误**，因为每个人都至少使用相同的基础架构技术，即使它在本地运行时有一些细微的差异。

Another improvement compared to the manual pipeline approach is the **automatic execution of the pipeline** and the **additional features for development** that speed up the dev process a little bit.

与手动流水线方法相比，另一个改进是**流水线的自动执行**和**用于开发的附加功能**，它们稍微加快了开发过程。

### [\#](http://loft.sh\#disadvantages-2) Disadvantages

### [\#](http://loft.sh\#disadvantages-2) 缺点

However, level 2 development is still a **relatively slow process** compared to traditional development because pipelines still have to be executed resulting in waiting times even if this process is triggered automatically and is somewhat faster due to the local runtime environment.

然而，与传统开发相比，2 级开发仍然是一个**相对较慢的过程**，因为即使此过程是自动触发的，仍然需要执行管道，从而导致等待时间，并且由于本地运行时环境而速度稍快一些。

Additionally, developers are now in contact with Kubernetes and usually have to **set up and manage it themselves** on their local computers. This, of course, comes with some **extra effort and requires additional knowledge** as the configuration of the local cluster can be tricky in some cases.

此外，开发人员现在正在与 Kubernetes 联系，并且通常必须在本地计算机上自行**设置和管理**。当然，这需要一些**额外的努力并需要额外的知识**，因为在某些情况下本地集群的配置可能会很棘手。

Another downside of using a local runtime environment is the **limitation in terms of computing resources** that can make it infeasible again for some more advanced applications.

使用本地运行时环境的另一个缺点是**在计算资源方面的限制**，这可能使其对于一些更高级的应用程序再次不可行。

## [\#](http://loft.sh\#level-3--cloud-native-development--development-takes-place-inside-the-kubernetes-cluster) Level 3 – Cloud-native development – Development takes place inside the Kubernetes cluster

## [\#](http://loft.sh\#level-3--cloud-native-development--development-takes-place-inside-the-kubernetes-cluster) Level 3 – 云原生开发 –开发发生在 Kubernetes 集群内

At level 3, the development of the software takes place inside a Kubernetes cluster without the need of running CD pipelines on every change. Instead, the developer uses special tools such as [DevSpace](https://github.com/devspace-cloud/devspace) or [Okteto](https://github.com/okteto/okteto) that simulate traditional development in a cluster as much as possible.

在第 3 级，软件的开发在 Kubernetes 集群内进行，无需在每次更改时运行 CD 管道。相反，开发人员使用特殊工具，例如 [DevSpace](https://github.com/devspace-cloud/devspace) 或 [Okteto](https://github.com/okteto/okteto) 模拟传统开发尽可能地集群。

These tools recognize file changes and synchronize them to the container filesystem inside the Kubernetes cluster. This allows the application to be updated instantly (hot reload) and images will only be rebuilt when necessary, i.e. after the image was changed.

这些工具识别文件更改并将它们同步到 Kubernetes 集群内的容器文件系统。这允许应用程序立即更新（热重载），并且仅在必要时重建图像，即在图像更改之后。

Besides this hot reloading feature, the cloud-native development tools provide a port forwarding feature to allow development on localhost and give developers full terminal and log access. 

除了这种热重载功能外，云原生开发工具还提供端口转发功能，允许在 localhost 上进行开发，并为开发人员提供完整的终端和日志访问权限。

Another important aspect of level 3 development is that the developers have direct access to a remote [Kubernetes dev environment](http://loft.sh/blog/kubernetes-development-environments-comparison/?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development) and do not use a local Kubernetes cluster anymore. To enable this access without having a cluster for each developer, tools such as [Loft](https://loft.sh) provide a centrally managed, [multi-tenancy Kubernetes platform](https://loft.sh/features/kubernetes-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development) for the developers.

3 级开发的另一个重要方面是开发人员可以直接访问远程 Kubernetes 开发环境，并且不再使用本地 Kubernetes 集群。为了在不为每个开发人员使用集群的情况下启用此访问权限，[Loft](https://loft.sh) 等工具提供了一个集中管理的 [多租户 Kubernetes 平台](https://loft.sh/features/kubernetes-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development)供开发人员使用。

### [\#](http://loft.sh\#advantages-3) Advantages

### [\#](http://loft.sh\#advantages-3) 优点

The goal of level 3 development is to combine the best of both worlds, traditional development speed with Kubernetes’ realism, scalability, and replicability. Due to the hot reloading feature instead of the pipeline approaches of levels 1 and 2, cloud-native development becomes **much faster** and is only slightly slower than traditional local development without Kubernetes.

3 级开发的目标是结合两全其美，传统的开发速度与 Kubernetes 的真实性、可扩展性和可复制性。由于热重载特性而不是级别 1 和 2 的管道方法，云原生开发变得**快得多**，并且仅比没有 Kubernetes 的传统本地开发慢一点。

Since all the development processes and commands needed by the developers are built-in or can be pre-configured by a Kubernetes expert in a team, the **complexity to use this approach is relatively low** even though the developers have a direct access to Kubernetes.

由于开发人员需要的所有开发流程和命令都是内置的，或者可以由团队中的 Kubernetes 专家预先配置，因此**使用这种方法的复杂性相对较低**，即使开发人员可以直接访问到 Kubernetes。

This direct access in a cloud environment also allows **very realistic testing** and is **not limited in terms of computing resources**, which makes this form of development feasible and efficient even for complex applications.

这种在云环境中的直接访问还允许**非常逼真的测试**并且**在计算资源方面不受限制**，这使得这种形式的开发即使对于复杂的应用程序也是可行和高效的。

Finally, since the dev tools are not hardwired to the cluster, it is **easily possible to switch the runtime environment** (e.g. switch from local to a remote cluster or switch between different remote clusters), which **prevents a lock- in effect**.

最后，由于开发工具没有硬连线到集群，因此**可以轻松切换运行时环境**（例如，从本地切换到远程集群或在不同的远程集群之间切换），**防止锁定-有效**。

### [\#](http://loft.sh\#disadvantages-3) Disadvantages

### [\#](http://loft.sh\#disadvantages-3) 缺点

As most advanced level of cloud-native development, level 3 development is also most efficient for many cases. It is only **slightly slower than traditional development** and **some workflows have to be adapted**.

作为最高级的云原生开发，Level 3 开发在很多情况下也是最高效的。它只是**比传统开发稍慢**并且**必须调整一些工作流程**。

It also requires **some setup effort once when introduced** to determine the optimal configuration for the tools that can then be shared in a whole team.

它还需要**在引入时进行一些设置工作**以确定可以在整个团队中共享的工具的最佳配置。

## [\#](http://loft.sh\#a-closer-look-at-the-access-to-kubernetes) A closer look at the access to Kubernetes

## [\#](http://loft.sh\#a-closer-look-at-the-access-to-kubernetes) 深入了解Kubernetes的访问

Since the goal of the development processes described in this post is to run great software in Kubernetes, Kubernetes needs to be part of every development process at some point. By advancing from one level to the next, the touchpoint of the software with a (realistic) Kubernetes environment becomes earlier.

由于本文中描述的开发过程的目标是在 Kubernetes 中运行出色的软件，因此 Kubernetes 在某些时候需要成为每个开发过程的一部分。通过从一个级别推进到下一个级别，具有（现实）Kubernetes 环境的软件的接触点变得更早。

At level 0, Kubernetes is purely an operations topic and the developers are usually not in touch with it at all. They might only get some feedback if their software is not running properly in production due to problems that were not considered when development took place locally.

在 0 级，Kubernetes 纯粹是一个操作主题，开发人员通常根本不接触它。如果他们的软件由于在本地开发时没有考虑到的问题而在生产中无法正常运行，他们可能只会得到一些反馈。

At level 1, Kubernetes is still an ops topic, but the developers have indirect access to it during development via a CD pipeline. However, they still do not need to manage anything and do not have the right to configure anything.

在级别 1，Kubernetes 仍然是一个 ops 主题，但开发人员可以在开发过程中通过 CD 管道间接访问它。但是，他们仍然不需要管理任何东西，也没有配置任何东西的权利。

Starting from level 2, Kubernetes becomes a development topic. That means that the developers also need to get direct access to it. In level 2, this usually means that they start and manage their own local Kubernetes clusters while in level 3, they get access to a shared cluster via specific tools such as [Loft](https://loft.sh) or [kiosk](https://github.com/kiosk-sh/kiosk). 

从第 2 级开始，Kubernetes 成为一个开发主题。这意味着开发人员还需要直接访问它。在级别 2 中，这通常意味着他们启动和管理自己的本地 Kubernetes 集群，而在级别 3 中，他们可以通过 [Loft](https://loft.sh) 或 [kiosk] 等特定工具访问共享集群（https://github.com/kiosk-sh/kiosk)。

## [\#](http://loft.sh\#conclusion) Conclusion

## [\#](http://loft.sh\#conclusion) 结论

I hope this post was helpful to understand where you and your team are on the cloud-native journey and to get an impression about what next steps might await you.
However, while many companies actually go through the different levels one after the other, I want to note that it is absolutely possible to skip levels and to combine different approaches with each other to get maximum efficiency.

我希望这篇文章有助于了解您和您的团队在云原生之旅中所处的位置，并对接下来可能等待您的步骤有所了解。
然而，虽然许多公司实际上是一个接一个地经历不同的层次，但我想指出的是，绝对有可能跳过层次并将不同的方法相互结合以获得最大效率。

In the end, I am still convinced that many companies would benefit from going the full distance to level 3, especially if they really need the structural benefits of Kubernetes, such as scalability.

最后，我仍然相信许多公司会从全面升级到 Level 3 中受益，特别是如果他们真的需要 Kubernetes 的结构优势，例如可扩展性。

At the moment, I believe level 3 is the highest level on the cloud-native journey in software development but, as technology advances, I would not be surprised if I had to add a level 4 or 5. So, let's stay excited about what the next steps are coming in this fast-paced ecosystem.

目前，我相信第 3 级是软件开发中云原生之旅的最高级别，但是随着技术的进步，如果我不得不添加第 4 级或第 5 级，我不会感到惊讶。所以，让我们对什么感到兴奋在这个快节奏的生态系统中，接下来的步骤即将到来。

https://loft.sh/blog/the-journey-of-adopting-cloud-native-development 

