# How I prepared & passed the Certified Kubernetes Administrator (CKA) Exam

# 我是如何准备并通过 Kubernetes 管理员 (CKA) 考试的

Tips for preparing for the Certified Kubernetes Administrator(CKA) exam.

准备 Kubernetes 管理员 (CKA) 认证考试的技巧。

> **NOTE:**
>
> This blog post was published in **2019** when the exam environment was based on the version **v1.16** of Kubernetes, which makes it a bit outdated by now.
>
> However, you might still find it useful as it includes general information that is not version dependent and some external resources that are still regularly updated.

> **注意：**
>
> 这篇博文发布于**2019**，当时考试环境基于Kubernetes的**v1.16**版本，这使得它现在有点过时。
>
> 但是，您可能仍然会发现它很有用，因为它包含与版本无关的一般信息和一些仍会定期更新的外部资源。

* * *

* * *

Finishing the year as a Certified Kubernetes Administrator(CKA) was my personal goal for 2019, and I was able to take the exam around the end of the year and pass with a score of **91%**. 🎉

作为一名认证 Kubernetes 管理员 (CKA) 完成这一年是我 2019 年的个人目标，我能够在年底左右参加考试并以 **91%** 的分数通过。 🎉

In this blog post, I wanted to share some useful resources that helped me pass the CKA exam, and a few tips that can help you prepare and hopefully pass if you are also planning to take it.

在这篇博文中，我想分享一些帮助我通过 CKA 考试的有用资源，以及一些可以帮助您准备并希望通过的技巧（如果您也计划参加）。

> **DISCLAIMER:** This post is a bit long because I tried to dump all the knowledge and experience I gathered when preparing for the CKA exam. So brace yourself 😀

> **免责声明：** 这篇文章有点长，因为我试图抛弃我在准备 CKA 考试时收集的所有知识和经验。所以要振作起来😀

# The Certified Kubernetes Administrator Exam

# Kubernetes 管理员认证考试

With the exploding adoption of Kubernetes, the [Certified Kubernetes Administrator](https://www.cncf.io/certification/cka/) program was created by the [Cloud Native Computing Foundation](https://www.cncf.io/)(CNCF) in collaboration with the Linux Foundation to allow Kubernetes users to demonstrate that they have the necessary skills and knowledge to perform the tasks and responsibilities of a Kubernetes administrator.

随着 Kubernetes 的爆发式采用，[Certified Kubernetes Administrator](https://www.cncf.io/certification/cka/) 计划由 [Cloud Native Computing Foundation](https://www.cncf.io)创建/)(CNCF) 与 Linux 基金会合作，允许 Kubernetes 用户证明他们具备必要的技能和知识来执行 Kubernetes 管理员的任务和职责。

## The exam's format

## 考试形式

The good thing about it is that it's 100% hands-on. It's an online proctored exam where you are asked to perform certain tasks on the command line.

它的好处是它是 100% 动手操作的。这是一项在线监考考试，要求您在命令行上执行某些任务。

The [Candidate Handbook](https://training.linuxfoundation.org/go/cka-ckad-candidate-handbook) is your definitive source for any details about the exam. So make sure to read it **thoroughly**.

[考生手册](https://training.linuxfoundation.org/go/cka-ckad-candidate-handbook) 是有关考试任何详细信息的权威来源。所以一定要**通读**。

Here is a short list of points worth mentioning:

以下是值得一提的要点的简短列表：

- You need a **steady** internet connection.
- You would need a**webcam** and a **microphone** which are required by the proctor.
- You would need a government issued **ID**, or a passport.
- The exam consists of **24 questions** that you can solve in no specific order.
- The duration of the exam is **3 hours**.
- The pass mark is**74%.**
- You need to use the**Chrome browser**.
- You have **one free retake** in case you don't pass on your first try 🎉

- 您需要**稳定**的互联网连接。
- 您需要监考人员要求的**网络摄像头**和**麦克风**。
- 您需要政府签发的 **ID** 或护照。
- 考试包含 **24 个问题**，您可以不按特定顺序解决这些问题。
- 考试时间为**3小时**。
- 及格分数为**74%。**
- 您需要使用**Chrome 浏览器**。
- 如果您没有通过第一次尝试，您有**一次免费重考**

## The curriculum

##  课程

Unlike the Certified [Kubernetes Application Developer](https://www.cncf.io/certification/ckad/)(CKAD) exam, the CKA exam focuses more on cluster administration rather than deploying and managing applications on Kubernetes.

与 Certified [Kubernetes Application Developer](https://www.cncf.io/certification/ckad/)(CKAD) 考试不同，CKA 考试更侧重于集群管理，而不是在 Kubernetes 上部署和管理应用程序。

The exam's curriculum is usually updated quarterly, you can always find the latest version at:

考试课程通常每季度更新一次，您始终可以在以下位置找到最新版本：

Will I receive the hardware or come to the office to pick it up?

我会收到硬件还是来办公室取件？

The CKA exam covers the following topics:

CKA 考试涵盖以下主题：

- Application Lifecycle Management – 8%
- Installation, Configuration & Validation – 12%
- Core Concepts – 19%
- Networking – 11%
- Scheduling – 5%
- Security – 12%
- Cluster Maintenance – 11%
- Logging / Monitoring – 5%
- Storage – 7%
- Troubleshooting – 10%



- 应用程序生命周期管理 - 8%
- 安装、配置和验证 - 12%
- 核心概念 - 19%
- 网络 - 11%
- 调度 - 5%
- 安全性 - 12%
- 集群维护 - 11%
- 记录/监控 – 5%
- 存储 - 7%
- 故障排除 - 10%

## The exam environment

## 考试环境

The day of the exam, you will have multiple cluster presented for you, and with each question you will be provided with name of the cluster where you should try to solve the question.

考试当天，您将收到多个集群，每个问题都会为您提供您应该尝试解决问题的集群名称。

Below is the list of the clusters provided to the candidate from the latest [Exam Tips](http://training.linuxfoundation.org/go//Important-Tips-CKA-CKAD) document available at the CKA CNCF page at the time of writing this post:

以下是当时在 CKA CNCF 页面上提供的最新 [Exam Tips](http://training.linuxfoundation.org/go//Important-Tips-CKA-CKAD) 文档中提供给考生的集群列表写这篇文章：

The Kubernetes version running on the exam environment is currently **v1.16** at the time of writing this post, and the Linux distribution is **Ubuntu 16**.

在撰写本文时，考试环境中运行的 Kubernetes 版本目前为 **v1.16**，Linux 发行版为 **Ubuntu 16**。

# Preparing for the exam

# 准备考试

The first step in preparing for the CKA exam(or any exam) is understanding what it is about.

准备 CKA 考试（或任何考试）的第一步是了解它的内容。

So make sure to read all the documents provided in the CKA Program page at [https://www.cncf.io/certification/cka/](https://www.cncf.io/certification/cka/) :

因此，请务必阅读 [https://www.cncf.io/certification/cka/](https://www.cncf.io/certification/cka/) 的 CKA 计划页面中提供的所有文件：

- [Candidate Handbook](https://training.linuxfoundation.org/go/cka-ckad-candidate-handbook)
- [Curriculum Overview](https://github.com/cncf/curriculum) 

- [候选人手册](https://training.linuxfoundation.org/go/cka-ckad-candidate-handbook)
- [课程概览](https://github.com/cncf/curriculum)

- [Exam Tips](http://training.linuxfoundation.org/go//Important-Tips-CKA-CKAD)
- [Frequently Asked Questions](http://training.linuxfoundation.org/go/cka-ckad-faq)

- [考试技巧](http://training.linuxfoundation.org/go//Important-Tips-CKA-CKAD)
- [常见问题](http://training.linuxfoundation.org/go/cka-ckad-faq)

## Pre-requisites

## 先决条件

Although the CKA exam is about Kubernetes, it also requires some basic sysadmin skills. So, you need be comfortable with the Linux command line and have a minimum knowledge on how to use the following tools:

虽然 CKA 考试是关于 Kubernetes 的，但它也需要一些基本的系统管理员技能。因此，您需要熟悉 Linux 命令行，并至少了解如何使用以下工具：

- `systemd` for managing system services. Basic knowledge would be enough IMHO, but very important especially for troubleshooting cluster components. There is a nice tutorial series for that provided by the DigitalOcean people:

- `systemd` 用于管理系统服务。恕我直言，基本知识就足够了，但对于对集群组件进行故障排除尤其重要。 DigitalOcean 人提供了一个很好的教程系列：

Systemd Essentials: Working with Services, Units, and the Journal \| DigitalOcean

Systemd Essentials：使用服务、单元和期刊 \|数字海洋

In recent years, Linux distributions have increasingly migrated away from other init systems to systemd. The systemd suite of tools provides a fast and flexible init model for managing an entire machine, from boot onwards. In this guide, we’ll give you a quick run...\


近年来，Linux 发行版越来越多地从其他 init 系统迁移到 systemd。 systemd 工具套件提供了一个快速灵活的初始化模型，用于从启动开始管理整台机器。在本指南中，我们将为您提供快速运行...\


- `vim` for editing files on the command line. Although you could change the default text editor by setting the value of `$EDITOR` to nano if that's what you are most comfortable with, vim can give you a productive boost during the exam.
- `tmux` since you only get one console during the exam, being able to have multiple panes open at the same time might be helpful. Personally, I didn't really need or use tmux during the exam, so if you don't use it already in your day to day work, I don't recommend learning it for the sake of the exam.
- `openssl` for generating keys, CSRs, certificates etc.. You will probably need it during the exam for security related questions. So make sure you train yourself to use it at least for those basic use cases.

- `vim` 用于在命令行上编辑文件。尽管您可以通过将 `$EDITOR` 的值设置为 nano 来更改默认文本编辑器，如果这是您最熟悉的方式，但 vim 可以在考试期间提高您的效率。
- `tmux` 因为你在考试期间只有一个控制台，能够同时打开多个窗格可能会有所帮助。就我个人而言，我在考试期间并没有真正需要或使用 tmux，所以如果您在日常工作中还没有使用它，我不建议为了考试而学习它。
- `openssl` 用于生成密钥、CSR、证书等。您可能在考试期间需要它来解决安全相关问题。因此，请确保您训练自己至少在那些基本用例中使用它。

## Getting ready for the exam

## 准备考试

In this section, I am going to provide some tips on how to prepare for the exam and also list some useful resources that helped me and might help you get fit for the exam day.

在本节中，我将提供一些有关如何准备考试的提示，并列出一些有用的资源，这些资源对我有帮助，可能会帮助您适应考试日。

### kubectl

Since the CKA exam is 100% practical, you need to make sure you are confident enough with `kubectl`. That's mostly what you will be using during the exam, and since you are already reading this post, chances are you are already using kubectl or at least experimenting with it.

由于 CKA 考试 100% 实用，因此您需要确保对 `kubectl` 有足够的信心。这主要是您在考试期间将使用的内容，并且由于您已经阅读了这篇文章，因此您很可能已经在使用 kubectl 或至少正在尝试使用它。

You need to be quick on the command line since you will have limited time for solving the questions during the exam, so knowing how to perform the following quickly with kubectl is crucial:

您需要快速使用命令行，因为您在考试期间解决问题的时间有限，因此了解如何使用 kubectl 快速执行以下操作至关重要：

- Checking the config, switching and creating contexts
- Creating, editing and deleting kubernetes resources
- Viewing, finding and inspecting resources
- Updating and patching resources
- Interacting with pods, nodes and cluster

- 检查配置、切换和创建上下文
- 创建、编辑和删除 kubernetes 资源
- 查看、查找和检查资源
- 更新和修补资源
- 与 Pod、节点和集群交互

A lot of useful `kubectl` command examples can be found in the [kubectl cheatsheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/) available in the official kubernetes documentation.

在 kubernetes 官方文档中的 [kubectl 备忘单](https://kubernetes.io/docs/reference/kubectl/cheatsheet/) 中可以找到很多有用的 `kubectl` 命令示例。

It is also very useful to know how to use the `kubectl run` command to create resources quickly, saving time by avoiding to write yaml files(who likes that, right?). You can also use it to generate yaml files if you need to edit something before actually creating the kubernetes object by using the `--dry-run`  and the `-o yaml` options combined. Some details about the `kubectl run` usage can be found [here](https://kubernetes.io/docs/reference/kubectl/conventions/#kubectl-run).

知道如何使用 `kubectl run` 命令快速创建资源也非常有用，避免编写 yaml 文件来节省时间（谁喜欢，对吧？）。如果您需要在实际创建 kubernetes 对象之前通过结合使用 `--dry-run` 和 `-o yaml` 选项进行编辑，也可以使用它来生成 yaml 文件。可以在 [此处](https://kubernetes.io/docs/reference/kubectl/conventions/#kubectl-run) 中找到有关 `kubectl run` 用法的一些详细信息。

If you come from the Docker world and still starting with Kubernetes, then the **[kubectl for Docker Users](https://kubernetes.io/docs/reference/kubectl/docker-cli-to-kubectl/)** documentation page is definitely worth checking out.

如果您来自 Docker 世界并且仍然从 Kubernetes 开始，那么 **[kubectl for Docker Users](https://kubernetes.io/docs/reference/kubectl/docker-cli-to-kubectl/)** 文档页面绝对值得一试。

Getting familiar with **[JSONPath](https://kubernetes.io/docs/reference/kubectl/jsonpath/)** template would be also helpful. Combining kubectl and jsonpath enables you to easily extract resource information in a format that you can specify.

熟悉 **[JSONPath](https://kubernetes.io/docs/reference/kubectl/jsonpath/)** 模板也会有所帮助。结合 kubectl 和 jsonpath，您可以轻松地以您可以指定的格式提取资源信息。

Finally, make sure to practice a lot with kubectl, whether it is on local kubernetes clusters with minikube, docker-desktop or on the cloud. That is very crucial for the exam.

最后，一定要多多练习 kubectl，无论是在本地 kubernetes 集群上使用 minikube、docker-desktop 还是在云端。这对考试非常重要。

### Learning resources

### 学习资源

If you are planning to take the CKA exam, them you probably have already searched around the internet for some resources and found plenty. So in this section, I am only going to list the resources that I have found most informative and helpful for me to pass the exam.

如果您打算参加 CKA 考试，您可能已经在互联网上搜索了一些资源并找到了很多。因此，在本节中，我将仅列出我认为最有用且对我通过考试最有帮助的资源。

**_The Kubernetes Documentation_** 

**_Kubernetes 文档_**

The most important resource is the [official kubernetes documentation](https://kubernetes.io/docs/home/); that's your definitive source of information. And since you are **allowed** to access it during the exam, it's really important that you know how to easily navigate it and quickly search for what you need. Make sure to get well accustomed to it.

最重要的资源是[官方kubernetes文档](https://kubernetes.io/docs/home/)；那是您确定的信息来源。由于您**允许**在考试期间访问它，因此知道如何轻松导航并快速搜索所需内容非常重要。确保习惯它。

Also, make sure you get to do most if not all of the **tasks** listed [here](https://kubernetes.io/docs/tasks/).

此外，请确保您可以完成 [此处](https://kubernetes.io/docs/tasks/)列出的大部分（如果不是全部)**任务**。

> It's really useful to join the kubernetes slack community at [https://slack.k8s.io/](https://slack.k8s.io/). There is a slack channel dedicated to CKA exam questions named **#cka-exam-prep.**
>
> The members there are really nice and helpful and would answer any questions you have.

> 在 [https://slack.k8s.io/](https://slack.k8s.io/) 加入 kubernetes slack 社区真的很有用。有一个专门用于 CKA 考试题的 Slack 频道，名为 **#cka-exam-prep。**
>
> 那里的成员非常友善且乐于助人，他们会回答您的任何问题。

**Kubernetes The Hard Way(KTHW)**

**Kubernetes 艰难之路（KTHW）**

The [kubernetes-the-hard-way](https://github.com/kelseyhightower/kubernetes-the-hard-way) repo was created by Kelsey Hightower to provide a guide for bootstrapping a Kubernetes cluster on Google Cloud Platform. It helps you understand the internals of a Kubernetes cluster, which would be really important especially for troubleshooting.

[kubernetes-the-hard-way](https://github.com/kelseyhightower/kubernetes-the-hard-way) 存储库由 Kelsey Hightower 创建，为在 Google Cloud Platform 上引导 Kubernetes 集群提供指南。它可以帮助您了解 Kubernetes 集群的内部结构，这对于故障排除尤其重要。

Make sure to go through it at least once while trying to understand **every** step on the way.

确保至少通读一次，同时尝试了解 ** 过程中的每一个 ** 步骤。

If you don't want to use GCP, there is another fork that relies on vagrant and can be found here:

如果你不想使用 GCP，还有另一个依赖 vagrant 的 fork，可以在这里找到：

**Online Courses**

**在线课程**

There are a couple of online course available that can help you prepare for the CKA exam.

有几个在线课程可以帮助您准备 CKA 考试。

I was able to try 3 of them while preparing for the exam:

在准备考试时，我能够尝试其中的 3 个：

- [Kubernetes Fundamentals (LFS258)](https://training.linuxfoundation.org/training/kubernetes-fundamentals/) which was part of the CKA exam bundle I purchased.

- [Kubernetes Fundamentals (LFS258)](https://training.linuxfoundation.org/training/kubernetes-fundamentals/)，这是我购买的 CKA 考试包的一部分。

I only liked the practice labs, otherwise the course was very boring and you'd better read the kubernetes documentation rather than reading slides. So _IMHO_ totally not worth taking.
- [Linuxacademy: Cloud Native Certified Kubernetes Administrator (CKA)](https://linuxacademy.com/course/cloud-native-certified-kubernetes-administrator-cka/):

我只喜欢练习实验室，否则课程很无聊，你最好阅读 kubernetes 文档而不是阅读幻灯片。所以_恕我直言_完全不值得。
- [Linuxacademy: Cloud Native Certified Kubernetes Administrator (CKA)](https://linuxacademy.com/course/cloud-native-certified-kubernetes-administrator-cka/)：

I found this course really good at first. However, after a while I found myself watching the instructor mostly typing commands in the terminal so I got disconnected and stopped following the course. I also tried the mock exams, but I found them a bit limited.
- Udemy:[Certified Kubernetes Administrator (CKA) with Practice Tests](https://www.udemy.com/course/certified-kubernetes-administrator-with-practice-tests/)

一开始我觉得这门课真的很好。但是，过了一会儿，我发现自己在观看讲师的大部分内容是在终端中输入命令，因此我断开了连接并停止了学习课程。我也尝试过模拟考试，但我发现它们有点有限。
- Udemy：[经过实践测试的 Kubernetes 认证管理员 (CKA)](https://www.udemy.com/course/certified-kubernetes-administrator-with-practice-tests/)

This was the most comprehensive course for me in this list. It covered all the topics, and the instructor made sure to explain all the Kubernetes concepts(and also other concepts) thouroughly.

这是这份清单中对我来说最全面的课程。它涵盖了所有主题，并且讲师确保全面解释所有 Kubernetes 概念（以及其他概念）。

The practice labs are really good since you are provided with an environment and your answers are checked there automatically.

练习实验室非常好，因为你有一个环境，你的答案会在那里自动检查。

The mock exams were also a great preparation for the exam.

模拟考试也为考试做了很好的准备。

**I cannot _recommend_ this course enough!**

**我不能_推荐_这门课程！**

**Additional Resources**

**其他资源**

The [Kubernetes in Action](https://www.manning.com/books/kubernetes-in-action) book by Marko Lukša is definitely worth reading to have a good understanding of Kubernetes.

Marko Lukša 的 [Kubernetes in Action](https://www.manning.com/books/kubernetes-in-action) 这本书绝对值得一读，以更好地了解 Kubernetes。

There is a google spreadsheet created by the community that compiles a lot of useful resources that can be found [here](https://bit.ly/2IdKwIc).

有一个由社区创建的谷歌电子表格，其中汇集了许多有用的资源，可以在 [此处](https://bit.ly/2IdKwIc) 中找到。

Some additional useful Github repositories:

一些其他有用的 Github 存储库：

\- [https://github.com/stretchcloud/cka-lab-practice](https://github.com/stretchcloud/cka-lab-practice)

\- [https://github.com/stretchcloud/cka-lab-practice](https://github.com/stretchcloud/cka-lab-practice)

\- [https://github.com/walidshaari/Kubernetes-Certified-Administrator](https://github.com/walidshaari/Kubernetes-Certified-Administrator)

\- [https://github.com/walidshaari/Kubernetes-Certified-Administrator](https://github.com/walidshaari/Kubernetes-Certified-Administrator)

\- [https://github.com/krzko/awesome-cka](https://github.com/krzko/awesome-cka)

\- [https://github.com/krzko/awesome-cka](https://github.com/krzko/awesome-cka)

\- [https://github.com/David-VTUK/CKA-StudyGuide](https://github.com/David-VTUK/CKA-StudyGuide)

\- [https://github.com/David-VTUK/CKA-StudyGuide](https://github.com/David-VTUK/CKA-StudyGuide)

## Tips for the exam day

## 考试当天的提示

In this section, I am going to provide a few tips for the day of the exam:

在本节中，我将提供一些考试当天的提示：

- You are allowed to open one additional browser tab in addition to the exam interface and you can use it to browse the kubernetes documentation._Bookmarks_ are also allowed, so make sure to create some bookmarks in chromes for the documentation sections that you think you would need in the exam beforehand.
- You don't have to solve the questions in a specific order. So you can start with the easiest to build some confidence, but that's a matter of personal preference. 

- 除了考试界面，您还可以打开一个额外的浏览器选项卡，您可以使用它来浏览 kubernetes 文档。_Bookmarks_ 也是允许的，因此请确保在 chromes 中为您认为需要的文档部分创建一些书签需要在考试前。
- 您不必按特定顺序解决问题。所以你可以从最容易建立信心的开始，但这是个人喜好的问题。

- There is built-in notepad in the exam interface which might be handy since you're not allowed to write on paper during the exam. You can use it to write the questions' numbers so that you keep track of the ones you didn't solve and get back to them later.
- If you are taking the exam with a laptop, use an external monitor if your laptop screen is tiny. You would need all the space you can get for the terminal.
- Make sure to go to the restroom before starting the exam. During the 3 hours, you would only be able to take a break if your proctor allows it but the timer would never stop for you.
- Have some water in a bottle without a label, or a transparent glass. Anything other than that is not allowed.
- Take the exam in a quiet room, on a clean desk. Remove any electronics from the desk and make sure that absolutely no one enters the room during the exam.

- 考试界面中内置了记事本，这可能会很方便，因为您在考试期间不允许在纸上书写。您可以使用它来写问题的编号，以便您跟踪未解决的问题，并在以后返回。
- 如果您使用笔记本电脑参加考试，如果您的笔记本电脑屏幕很小，请使用外接显示器。您将需要为终端获得的所有空间。
- 请务必在开始考试前去洗手间。在这 3 个小时内，您只能在监考人员允许的情况下休息一下，但计时器永远不会为您停止。
- 在没有标签的瓶子或透明玻璃杯中放一些水。除此之外的任何事情都是不允许的。
- 在安静的房间，干净的桌子上参加考试。从桌子上取下任何电子设备，并确保在考试期间绝对没有人进入房间。

You will be asked by the proctor to show him around the room using the webcam.
- Finally:**_GOOD LUCK!_**

监考人员会要求您使用网络摄像头带他参观房间。
- 最后：**_祝你好运！_**

# Conclusion

#  结论

In this post, I tried to provide some tips and resources for preparing the CKA exam based on my experience.

在这篇文章中，我试图根据我的经验提供一些准备 CKA 考试的技巧和资源。

I hope this article would be useful for you and please let me know in the comments if it somehow helped you to pass the exam.

我希望这篇文章对你有用，如果它以某种方式帮助你通过考试，请在评论中告诉我。

### Subscribe to Mehdi Yedes' blog

### 订阅 Mehdi Yedes 的博客

Get the latest posts delivered right to your inbox

将最新帖子直接发送到您的收件箱

Subscribe 

订阅

