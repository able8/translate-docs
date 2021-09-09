# Advice for customers dealing with Docker Hub rate limits, and a Coming Soon announcement

# 对处理 Docker Hub 速率限制的客户的建议，以及即将发布的公告

by Omar Paul and Michael Hausenblas \| on
02 NOV 2020 \| in
[Amazon EC2 Container Registry](https://aws.amazon.com/blogs/containers/category/compute/amazon-ec2-container-registry/ "View all posts in Amazon EC2 Container Registry"), [Containers](https ://aws.amazon.com/blogs/containers/category/containers/ "View all posts in Containers") \|
[Permalink](https://aws.amazon.com/blogs/containers/advice-for-customers-dealing-with-docker-hub-rate-limits-and-a-coming-soon-announcement/) \|
Share

作者：奥马尔·保罗和迈克尔·豪森布拉斯 \|在
2020 年 11 月 2 日 \|在
[Amazon EC2 Container Registry](https://aws.amazon.com/blogs/containers/category/compute/amazon-ec2-container-registry/“查看 Amazon EC2 Container Registry 中的所有帖子”)，[Containers](https://aws.amazon.com/blogs/containers/category/containers/“查看容器中的所有帖子”)\|
[固定链接](https://aws.amazon.com/blogs/containers/advice-for-customers-dealing-with-docker-hub-rate-limits-and-a-coming-soon-announcement/) \|
分享

Many container customers building applications use common software packages (e.g. operating systems, databases, and application components) that are publicly distributed as container images on [Docker Hub](https://hub.docker.com/). Docker, Inc. has announced that the Hub service will begin [limiting](https://docs.docker.com/docker-hub/download-rate-limit/) the rate at which images are pulled under their anonymous and free plans . These limits will [progressively take effect](https://www.docker.com/increase-rate-limits) beginning November 2, 2020. Once fully in place, free plan anonymous use will be limited to 100 pulls per six hours, free plan authenticated accounts limited to 200 pulls per six hours, and Pro and Team accounts will not see any rate limits.

许多构建应用程序的容器客户使用通用软件包（例如操作系统、数据库和应用程序组件），这些软件包在 [Docker Hub](https://hub.docker.com/) 上作为容器映像公开分发。 Docker, Inc. 宣布 Hub 服务将开始[限制](https://docs.docker.com/docker-hub/download-rate-limit/) 根据其匿名和免费计划提取图像的速率.这些限制将从 2020 年 11 月 2 日开始[逐步生效](https://www.docker.com/increase-rate-limits)。一旦完全到位，免费计划匿名使用将被限制为每六小时 100 次拉取，免费计划认证帐户限制为每 6 小时 200 次拉取，Pro 和 Team 帐户不会看到任何速率限制。

With the introduction of these limits, our customers should expect some of their applications and tools that use public images from Docker Hub to face throttling errors, such as when they build from a parent public image or pull a public image to run. Many of our customers have expressed concern about possible impact, so we are sharing some practical advice for managing Docker Hub’s rate limits, and announcing an upcoming AWS solution.

随着这些限制的引入，我们的客户应该期望他们的一些使用来自 Docker Hub 的公共映像的应用程序和工具面临限制错误，例如当他们从父公共映像构建或拉取公共映像运行时。我们的许多客户都对可能产生的影响表示担忧，因此我们分享了一些有关管理 Docker Hub 速率限制的实用建议，并宣布了即将推出的 AWS 解决方案。

**Managing Docker Hub’s pull rate limit impact in the short-term**

**管理 Docker Hub 在短期内的拉取率限制影响**

If customers see impact, the overall approach we recommend is first to _identify_ public container images in use, then _choose_ a mitigation approach and take the necessary steps.

如果客户看到影响，我们建议的总体方法是首先_识别_正在使用的公共容器映像，然后_选择_缓解方法并采取必要的步骤。

You can identify Docker Hub public images by searching for files named `Dockerfile` in source code for the `FROM` command. If there is no registry URL preceding the image and tag, that image will pull from Docker Hub when it runs. For example, `FROM amazonlinux:latest` pulls the latest `amazonlinux` version from Docker Hub as the parent image in an application image build. For completeness, also search your container application cluster configurations for public images. To automate the detection of public images served from Docker Hub you can use [awslabs/aws-container-images-toolkit](https://github.com/awslabs/aws-container-images-toolkit), a tool we developed that allows you to generate a list of public images in code repositories, Amazon Elastic Container Service (ECS), as well as Amazon Elastic Kubernetes Service (EKS) and self-managed Kubernetes clusters. 

您可以通过在“FROM”命令的源代码中搜索名为“Dockerfile”的文件来识别 Docker Hub 公共镜像。如果镜像和标签之前没有注册 URL，则该镜像将在运行时从 Docker Hub 中拉取。例如，`FROM amazonlinux:latest` 从 Docker Hub 中提取最新的 `amazonlinux` 版本作为应用程序镜像构建中的父镜像。为完整起见，还可以在您的容器应用程序集群配置中搜索公共映像。要自动检测从 Docker Hub 提供的公共图像，您可以使用 [awslabs/aws-container-images-toolkit](https://github.com/awslabs/aws-container-images-toolkit)，这是我们开发的一个工具允许您在代码存储库、Amazon Elastic Container Service (ECS) 以及 Amazon Elastic Kubernetes Service (EKS) 和自我管理的 Kubernetes 集群中生成公共映像列表。

There are two mitigation approaches we recommend. **One**, copy public images being used into a private registry such as [Amazon Elastic Container Registry](https://aws.amazon.com/ecr/)(ECR). Or **two**, upgrade to a paid Docker Hub [subscription](https://www.docker.com/pricing). Both approaches require switching to an authenticated pull model. ECS customers can follow [these instructions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/private-auth.html) to authenticate task pull requests. Kubernetes customers, including EKS users can follow [these instructions](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/) to authenticate worker node or pod pulls. Note that public images copied to ECR may be pulled without additional authentication effort if AWS Identity and Access Management (IAM) prerequisites are in place. IAM prerequisites for authenticated pulls are documented for [ECS here](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR_on_ECS.html) and for [EKS here](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR_on_EKS.html). Customers also using third-party systems can consult this [Docker link](https://docs.docker.com/docker-hub/download-rate-limit/) for additional authenticated pull documentation.

我们推荐两种缓解方法。 **一**，将正在使用的公共映像复制到私有注册表中，例如 [Amazon Elastic Container Registry](https://aws.amazon.com/ecr/)(ECR)。或者**二**，升级到付费 Docker Hub [订阅](https://www.docker.com/pricing)。这两种方法都需要切换到经过身份验证的拉模型。 ECS 客户可以按照 [这些说明](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/private-auth.html) 对任务拉取请求进行身份验证。 Kubernetes 客户，包括 EKS 用户，可以按照 [这些说明](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/) 对工作节点或 pod 拉取进行身份验证。请注意，如果 AWS Identity and Access Management (IAM) 先决条件到位，复制到 ECR 的公共映像可能无需额外的身份验证工作即可提取。 [ECS 此处](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR_on_ECS.html) 和 [EKS 此处](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR_on_EKS.html)。还使用第三方系统的客户可以查阅此 [Docker 链接](https://docs.docker.com/docker-hub/download-rate-limit/) 以获取其他经过身份验证的拉取文档。

**Amazon ECS Agent**

**亚马逊 ECS 代理**

AWS publishes the Amazon ECS agent as a public container image and pre-installs it on all [ECS-optimized Amazon Machine Images](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) (AMIs). Any instances using the ECS-optimized AMIs do not depend on Docker Hub to launch the instance, so customers using these AMIs will not be impacted by Docker Hub’s limitations on pulls of the ECS agent container image. Customers who [update](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) the agent to the latest version using the AWS Console or CLI will also not see rate limit impacts when upgrading, as the upgrade process downloads the latest ECS agent image from Amazon S3 and does not rely on Docker Hub.

AWS 将 Amazon ECS 代理发布为公共容器映像并将其预安装在所有 [ECS-optimized Amazon Machine Images](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) (AMI)。任何使用 ECS 优化 AMI 的实例都不依赖于 Docker Hub 来启动实例，因此使用这些 AMI 的客户不会受到 Docker Hub 对 ECS 代理容器镜像拉取限制的影响。使用 AWS 控制台或 CLI [更新](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) 代理到最新版本的客户也不会看到速率限制升级时会产生影响，因为升级过程会从 Amazon S3 下载最新的 ECS 代理映像，并且不依赖于 Docker Hub。

Customers who create their own AMIs for use with ECS, or use third-party AMIs, must [manage installation and upgrades of the ECS agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-install.html#ecs-agent-install-nonamazonlinux). These customers can store their own copy of the ECS agent in a private registry (such as ECR) and pull it from there, download it from a public S3 bucket owned by AWS, or pull the image from Docker Hub. Customers downloading the ECS agent from Docker Hub may be subject to Docker Hub rate limits. To mitigate rate limit effects, these customers can take the steps outlined above to manage Docker’s pull rate impact.

创建自己的 AMI 以用于 ECS 或使用第三方 AMI 的客户必须[管理 ECS 代理的安装和升级](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-install.html#ecs-agent-install-nonamazonlinux)。这些客户可以将自己的 ECS 代理副本存储在私有注册表（例如 ECR)中并从那里拉取，从 AWS 拥有的公共 S3 存储桶下载，或从 Docker Hub 拉取映像。从 Docker Hub 下载 ECS 代理的客户可能会受到 Docker Hub 速率限制。为了减轻速率限制影响，这些客户可以采取上述步骤来管理 Docker 的拉取速率影响。

**Amazon EKS Add Ons**

**亚马逊 EKS 附加组件**

All EKS add-on software that is included with EKS clusters is hosted on ECR today and will not be subject to Docker Hub rate limits. EKS-built and maintained projects, such as the AWS Load Balancer Controller include helm charts that reference images on ECR as well. However, manifests for these projects available on GitHub include references to images that are published to Docker Hub. You can check whether you are using a manifest with a Docker Hub image using the steps described in this blog and update to use a chart that references an ECR hosted image. These charts are available in the [EKS Charts](https://github.com/aws/eks-charts) GitHub repository.

EKS 集群中包含的所有 EKS 附加软件今天都托管在 ECR 上，不会受到 Docker Hub 速率限制的限制。 EKS 构建和维护的项目（例如 AWS 负载均衡器控制器）也包括引用 ECR 上的图像的舵图。但是，GitHub 上提供的这些项目的清单包括对发布到 Docker Hub 的图像的引用。您可以使用本博客中描述的步骤检查您是否正在使用带有 Docker Hub 映像的清单，并更新以使用引用 ECR 托管映像的图表。这些图表可在 [EKS 图表](https://github.com/aws/eks-charts) GitHub 存储库中找到。

**Coming soon, a reliable solution for public container images**

**即将推出，一个可靠的公共容器镜像解决方案**

Today, Amazon Elastic Container Registry (ECR) customers download billions of images each week, enabling them to deploy containers for use in their own applications running on-premises and in AWS. Events like Docker Hub rate limits reinforce the reality that public container images are a critical part of the container software supply chain. Customers have asked us to help make public container images they use as highly available and dependable as their images hosted in ECR. 

如今，Amazon Elastic Container Registry (ECR) 客户每周下载数十亿个映像，这使他们能够部署容器，以便在自己在本地和 AWS 中运行的应用程序中使用。 Docker Hub 速率限制等事件强化了这样一个现实，即公共容器镜像是容器软件供应链的关键部分。客户要求我们帮助他们制作与托管在 ECR 中的镜像一样高可用和可靠的公共容器镜像。

We have been working for many months to answer this customer ask. Within weeks, **AWS will deliver a new public container registry** that will allow developers to share and deploy container images publicly. This new registry will allow developers to store, manage, share, and deploy container images for anyone to discover and download. Developers will be able to use AWS to host both their private and public container images, eliminating the need to use different public websites and registries. Public images will be geo-replicated for reliable availability around the world and offer fast downloads to quickly serve up images on-demand. Anyone (with or without an AWS account) will be able to browse and pull containerized software for use in their own applications. Developers will be able use the new registry to distribute public container images and related files like helm charts and policy configurations for use by any developer. **A new website** will allow anyone to browse and search for public container images, view developer provided details, and see pull commands — all without needing to sign in to AWS. AWS-provided public images such as the ECS agent, Amazon CloudWatch agent, and AWS Deep Learning Container images will also be available.

我们已经工作了好几个月来回答这个客户的问题。几周内，**AWS 将提供一个新的公共容器注册表**，允许开发人员公开共享和部署容器映像。这个新的注册表将允许开发人员存储、管理、共享和部署容器镜像，供任何人发现和下载。开发人员将能够使用 AWS 来托管他们的私有和公共容器映像，从而无需使用不同的公共网站和注册表。公共图像将进行地理复制以在全球范围内提供可靠的可用性，并提供快速下载以快速按需提供图像。任何人（有或没有 AWS 账户）都可以浏览和提取容器化软件以在他们自己的应用程序中使用。开发人员将能够使用新的注册表来分发公共容器映像和相关文件，例如舵图和策略配置，供任何开发人员使用。 **一个新网站**将允许任何人浏览和搜索公共容器映像、查看开发人员提供的详细信息以及查看拉取命令——所有这些都无需登录 AWS。 AWS 提供的公共映像，例如 ECS 代理、Amazon CloudWatch 代理和 AWS Deep Learning Container 映像也将可用。

AWS customers will enjoy a smooth, dependable experience when working with public container images, and with pricing that provides additional value. A developer sharing public images on AWS will get 50 GB of free storage each month and will pay nominal charges after. Anyone who pulls images anonymously will get 500 GB of free data bandwidth each month after which they can sign up or sign in to an AWS account. Simply authenticating with an AWS account increases free data bandwidth up to 5 TB each month when pulling images from the internet. And finally, workloads running in AWS will get unlimited data bandwidth from any region when pulling publicly shared images hosted on AWS.

在使用公共容器映像时，AWS 客户将享受流畅、可靠的体验，并且定价可提供附加价值。在 AWS 上共享公共图像的开发人员每月将获得 50 GB 的免费存储空间，之后将支付象征性费用。任何匿名拉取图像的人每月将获得 500 GB 的免费数据带宽，之后他们可以注册或登录 AWS 账户。从 Internet 提取图像时，只需使用 AWS 帐户进行身份验证，每月免费数据带宽就可增加多达 5 TB。最后，在 AWS 中运行的工作负载在提取托管在 AWS 上的公开共享图像时，将从任何区域获得无限的数据带宽。

TAGS:
[Containers](https://aws.amazon.com/blogs/containers/tag/containers/),[docker](https://aws.amazon.com/blogs/containers/tag/docker/), [Docker Hub](https://aws.amazon.com/blogs/containers/tag/docker-hub/),[ECR](https://aws.amazon.com/blogs/containers/tag/ecr/), [ ECS](https://aws.amazon.com/blogs/containers/tag/ecs/),[EKS](https://aws.amazon.com/blogs/containers/tag/eks/), [Kubernetes] (https://aws.amazon.com/blogs/containers/tag/kubernetes/)

标签：
[容器](https://aws.amazon.com/blogs/containers/tag/containers/),[docker](https://aws.amazon.com/blogs/containers/tag/docker/), [Docker集线器](https://aws.amazon.com/blogs/containers/tag/docker-hub/)、[ECR](https://aws.amazon.com/blogs/containers/tag/ecr/)、[ECS](https://aws.amazon.com/blogs/containers/tag/ecs/)、[EKS](https://aws.amazon.com/blogs/containers/tag/eks/)、[Kubernetes]（https://aws.amazon.com/blogs/containers/tag/kubernetes/)

![Omar Paul](https://en.gravatar.com/userimage/26424230/14693d7b43017b8dfefdee24316b3fc8.jpg)

### Omar Paul

### 奥马尔保罗

Omar is a Product Manager in the AWS container services team. He focuses on all things container registry. Before AWS, Omar worked at API startups in Austin, TX and also spent a lot of time in the Washington DC area, going to GWU and working in telecom. He thinks AWS is like thousands of startups with a common set of principles.

Omar 是 AWS 容器服务团队的产品经理。他专注于容器注册的所有事情。在加入 AWS 之前，Omar 曾在德克萨斯州奥斯汀的 API 初创公司工作，并在华盛顿特区度过了大量时间，去了 GWU 并在电信工作。他认为 AWS 就像成千上万家拥有一套共同原则的初创公司。

![Michael Hausenblas](https://d2908q01vomqb2.cloudfront.net/ca3512f4dfa95a03169c5a670a4c91a19b3077b4/2019/06/12/mic_2013_07_bw.png)

### Michael Hausenblas

### 迈克尔豪森布拉斯

Michael is a Principal Product Developer Advocate in the AWS container service team. He covers observability, Kubernetes, service meshes, as well as container security and policies. Before Amazon, Michael worked at Red Hat, Mesosphere (now D2iQ), MapR (now part of HPE), and in two applied research institutions in Ireland and Austria. 

Michael 是 AWS 容器服务团队的首席产品开发倡导者。他涵盖了可观察性、Kubernetes、服务网格以及容器安全和策略。在加入亚马逊之前，Michael 曾在 Red Hat、Mesosphere（现为 D2iQ）、MapR（现为 HPE 的一部分）以及爱尔兰和奥地利的两家应用研究机构工作。

