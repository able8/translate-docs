# Learning Path: Basics of Container Runtimes

# 学习路径：容器运行时的基础知识

Details: https://twitter.com/erkan_erol_/status/1252343453402361858

详情：https://twitter.com/erkan_erol_/status/1252343453402361858

After a couple of productive hours, now I have a clear understanding of what are the difference/relationship between cgroups, namespaces, runc, containerd, docker,  rkt, dockerd, docker-containerd-shim, cri-o, cri, dockershim, gVisor , kata-containers, etc.

经过几个小时的高效工作，现在我清楚地了解了 cgroups、命名空间、runc、containerd、docker、rkt、dockerd、docker-containerd-shim、cri-o、cri、dockershim、gVisor 之间的区别/关系、kata 容器等。

I read/watched a lot. I'm extremely tired now but I am going to share a learning path (including videos, docs, blogs ) .

我阅读/观看了很多。我现在非常累，但我要分享一个学习路径（包括视频、文档、博客）。

First of all, the contents I am going to share contain some repetition. I know repetition is boring but it is also very useful to reinforce knowledge/understanding. Clarity requires effort and time :)

首先，我要分享的内容有一些重复。我知道重复很无聊，但它对加强知识/理解也非常有用。清晰需要努力和时间:)

Let's start with the basics.
 I suggest watching this video to understand what cgroups&namespaces are and how container runtimes interact with them. https://www.youtube.com/watch?v=sK5i-N34im8

让我们从基础开始。
我建议观看此视频以了解 cgroups 和命名空间是什么以及容器运行时如何与它们交互。 https://www.youtube.com/watch?v=sK5i-N34im8

Then continue with the first part of the "Container Runtimes" series. With this blog post, you are going to learn
 - "container runtime" term is not clear
 - there are some high level and low-level container runtimes.
 https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r

然后继续“容器运行时”系列的第一部分。通过这篇博文，您将学习
- “容器运行时”术语不清楚
- 有一些高级和低级容器运行时。
https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r

The second part of the "Container Runtimes" series explains some details about low-level container runtimes, which is going to reinforce your understanding of what "runc" is and how it works. https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level-contai

“容器运行时”系列的第二部分解释了有关低级容器运行时的一些细节，这将加强您对“runc”是什么及其工作方式的理解。 https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level-contai

The third part of the "Container Runtimes" series explains high-level container runtimes. This is going to teach you the relationship between docker and containerd. https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes

“容器运行时”系列的第三部分解释了高级容器运行时。这将教你docker和containerd之间的关系。 https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes

At this point, I suggest reading this blog post. It lists the components of docker and their interaction.
 I liked the explanation about the reason for the existence of "docker-containerd-shim".
 https://alexander.holbreich.org/docker-components-explained/

在这一点上，我建议阅读这篇博文。它列出了 docker 的组件及其交互。
我喜欢关于“docker-containerd-shim”存在原因的解释。
https://alexander.holbreich.org/docker-components-explained/

Let's go back for the fourth chapter of the series: Kubernetes and Container Runtime Interface
 I suggest to download "crictl" and give it a try.
 Note: It is already available in the images of nodes of "kind". Just run a docker exec and dig deeper!!!
 https://www.ianlewis.org/en/container-runtimes-part-4-kubernetes-container-run

让我们回到本系列的第四章：Kubernetes 和容器运行时接口
我建议下载“crictl”并尝试一下。
注意：它已经在“种类”节点的图像中可用。只需运行一个 docker exec 并深入挖掘！！！
https://www.ianlewis.org/en/container-runtimes-part-4-kubernetes-container-run

After learning docker, containerd and cri-o, it is better to read a blog post about their comparison. I liked the short list about how cri-o does handle its responsibilities.
 https://computingforgeeks.com/docker-vs-cri-o-vs-containerd/

在学习了 docker、containerd 和 cri-o 之后，最好阅读一篇关于它们的比较的博客文章。我喜欢关于 cri-o 如何处理其职责的简短列表。
https://computingforgeeks.com/docker-vs-cri-o-vs-containerd/

At this point, I think you can take a look at cri-o website. The architecture document explains which libraries cri-o uses behind the scene and where. https://cri-o.io

说到这里，我想你可以看看cri-o网站。架构文档解释了 cri-o 在幕后使用了哪些库以及在哪里使用。 https://cri-o.io

I had been always thinking that cri-o is something big like docker. I don't know why :D  When I realized that cri-o doesn't care development of containers, I was really surprised.
 One of the maintainers of cri-o explains what they target.
 https://medium.com/cri-o/container-runtimes-clarity-342b62172dc3

我一直认为 cri-o 和 docker 一样大。我不知道为什么 :D 当我意识到 cri-o 不关心容器的开发时，我真的很惊讶。
cri-o 的一位维护者解释了他们的目标。
https://medium.com/cri-o/container-runtimes-clarity-342b62172dc3

You think it is enough. Right? No. There are still somethings we need to learn Partying facePartying facePartying face

你觉得就够了。对？不，我们还有一些东西需要学习派对脸派对脸派对脸

This blog post explains CRI in detail. I think there are 2 important points.
 1) the PodSandbox term
 2) imperative container-centric interface
 https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/

这篇博文详细解释了 CRI。我认为有2点很重要。
1) PodSandbox 术语
2) 以容器为中心的命令式接口
https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/

After reading all of the contents above, I really enjoyed reading the slides of this presentation. It has very clear summaries of the big picture. I recommend reading the first 33 pages.
 https://events19.linuxfoundation.org/wp-content/uploads/2017/11/How-Container-Runtime-Matters-in-Kubernetes_-OSS-Kunal-Kushwaha.pdf

阅读完以上所有内容后，我真的很喜欢阅读本演示文稿的幻灯片。它对大局有非常清晰的总结。我建议阅读前 33 页。
https://events19.linuxfoundation.org/wp-content/uploads/2017/11/How-Container-Runtime-Matters-in-Kubernetes_-OSS-Kunal-Kushwaha.pdf

A note from the presentation:  dockershim is a part of kubelet not docker :D
 https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/dockershim

演示文稿中的注释：dockershim 是 kubelet 的一部分，而不是 docker :D
https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/dockershim

Another note from the presentation: gVisor and kata-runtime are something in runC level. They are handling the stuff in different ways at a very low level.

演讲中的另一个说明：gVisor 和 kata-runtime 是 runC 级别的东西。他们在非常低的水平上以不同的方式处理这些东西。