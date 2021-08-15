# Learning Path: Basics of Container Runtimes

Details: https://twitter.com/erkan_erol_/status/1252343453402361858

After a couple of productive hours, now I have a clear understanding of what are the difference/relationship between cgroups, namespaces, runc, containerd, docker,  rkt, dockerd, docker-containerd-shim, cri-o, cri, dockershim, gVisor, kata-containers, etc.

I read/watched a lot.  I'm extremely tired now but I am going to share a learning path (including videos, docs, blogs ) .

First of all, the contents I am going to share contain some repetition. I know repetition is boring but it is also very useful to reinforce knowledge/understanding. Clarity requires effort and time :)

Let's start with the basics. 
I suggest watching this video to understand what cgroups&namespaces are and how container runtimes interact with them.  https://www.youtube.com/watch?v=sK5i-N34im8

Then continue with the first part of the "Container Runtimes" series.  With this blog post, you are going to learn
- "container runtime" term is not clear
- there are some high level and low-level container runtimes.
https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r

The second part of the "Container Runtimes" series explains some details about low-level container runtimes, which is going to reinforce your understanding of what "runc" is and how it works. https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level-contai

The third part of the "Container Runtimes" series explains high-level container runtimes. This is going to teach you the relationship between docker and containerd. https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes

At this point, I suggest reading this blog post. It lists the components of docker and their interaction. 
I liked the explanation about the reason for the existence of "docker-containerd-shim".
https://alexander.holbreich.org/docker-components-explained/

Let's go back for the fourth chapter of the series: Kubernetes and Container Runtime Interface
I suggest to download "crictl" and give it a try. 
Note: It is already available in the images of nodes of "kind". Just run a docker exec and dig deeper!!!
https://www.ianlewis.org/en/container-runtimes-part-4-kubernetes-container-run

After learning docker, containerd and cri-o, it is better to read a blog post about their comparison. I liked the short list about how cri-o does handle its responsibilities. 
https://computingforgeeks.com/docker-vs-cri-o-vs-containerd/

At this point, I think you can take a look at cri-o website. The architecture document explains which libraries cri-o uses behind the scene and where.  https://cri-o.io

I had been always thinking that cri-o is something big like docker. I don't know why :D  When I realized that cri-o doesn't care development of containers, I was really surprised. 
One of the maintainers of cri-o explains what they target.
https://medium.com/cri-o/container-runtimes-clarity-342b62172dc3

You think it is enough. Right? No. There are still somethings we need to learn Partying facePartying facePartying face

This blog post explains CRI in detail. I think there are 2 important points.
1) the PodSandbox term
2) imperative container-centric interface
https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/

After reading all of the contents above, I really enjoyed reading the slides of this presentation. It has very clear summaries of the big picture. I recommend reading the first 33 pages.
https://events19.linuxfoundation.org/wp-content/uploads/2017/11/How-Container-Runtime-Matters-in-Kubernetes_-OSS-Kunal-Kushwaha.pdf

A note from the presentation:  dockershim is a part of kubelet not docker :D 
https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/dockershim

Another note from the presentation: gVisor and kata-runtime are something in runC level. They are handling the stuff in different ways at a very low level.

ENF OF FLOOD