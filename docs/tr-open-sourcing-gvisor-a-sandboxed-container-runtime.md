# Open-sourcing gVisor, a sandboxed container runtime

# 开源 gVisor，一个沙盒容器运行时

May 2, 2018

Containers have revolutionized how we develop, package, and deploy applications. However, the system surface exposed to containers is broad enough that many security experts [don't recommend them for running untrusted or potentially malicious applications.](https://opensource.com/business/14/7/docker-security-selinux)

容器彻底改变了我们开发、打包和部署应用程序的方式。然而，暴露于容器的系统表面足够广泛，以至于许多安全专家 [不推荐它们运行不受信任或潜在的恶意应用程序。](https://opensource.com/business/14/7/docker-security-selinux)

A growing desire to run more heterogenous and less trusted workloads has created a new interest in sandboxed containers—containers that help provide a secure isolation boundary between the host OS and the application running inside the container.

越来越多的人希望运行更多异构和不太可信的工作负载，这让人们对沙盒容器产生了新的兴趣——容器有助于在主机操作系统和容器内运行的应用程序之间提供安全的隔离边界。

To that end, we'd like to introduce [gVisor](https://github.com/google/gvisor), a new kind of sandbox that helps provide secure isolation for containers, while being more lightweight than a virtual machine (VM ). gVisor integrates with Docker and Kubernetes, making it simple and easy to run sandboxed containers in production environments.

为此，我们想介绍 [gVisor](https://github.com/google/gvisor)，这是一种新型沙箱，有助于为容器提供安全隔离，同时比虚拟机（VM)。 gVisor 与 Docker 和 Kubernetes 集成，使得在生产环境中运行沙盒容器变得简单容易。

### Traditional Linux containers are not sandboxes

### 传统的 Linux 容器不是沙箱

Applications that run in traditional Linux containers access system resources in the same way that regular (non-containerized) applications do: by making system calls directly to the host kernel. The kernel runs in a privileged mode that allows it to interact with the necessary hardware and return results to the application.

在传统 Linux 容器中运行的应用程序以与常规（非容器化）应用程序相同的方式访问系统资源：通过直接对主机内核进行系统调用。内核以特权模式运行，允许它与必要的硬件交互并将结果返回给应用程序。

![traditional-linux-containers-gvisorrquy.PNG](https://storage.googleapis.com/gweb-cloudblog-publish/images/traditional-linux-containers-gvisorrquy.max-300x300.PNG)

With traditional containers, the kernel imposes some limits on the resources the application can access. These limits are implemented through the use of Linux cgroups and namespaces, but not all resources can be controlled via these mechanisms. Furthermore, even with these limits, the kernel still exposes a large surface area that malicious applications can attack directly.

对于传统容器，内核对应用程序可以访问的资源施加了一些限制。这些限制是通过使用 Linux cgroup 和命名空间来实现的，但并非所有资源都可以通过这些机制进行控制。此外，即使有这些限制，内核仍然会暴露出大量恶意应用程序可以直接攻击的表面积。

Kernel features like [seccomp filters](https://en.wikipedia.org/wiki/Seccomp) can provide better isolation between the application and host kernel, but they require the user to create a predefined whitelist of system calls. In practice, it’s often difficult to know which system calls will be required by an application beforehand. Filters also provide little help when a vulnerability is discovered in a system call that your application requires.

[seccomp 过滤器](https://en.wikipedia.org/wiki/Seccomp) 等内核功能可以在应用程序和主机内核之间提供更好的隔离，但它们要求用户创建系统调用的预定义白名单。在实践中，通常很难事先知道应用程序需要哪些系统调用。当在您的应用程序需要的系统调用中发现漏洞时，过滤器也几乎没有帮助。

### Existing VM-based container technology

### 现有的基于虚拟机的容器技术

One approach to improve container isolation is to run each container in its own virtual machine (VM). This gives each container its own "machine," including kernel and virtualized devices, completely separate from the host. Even if there is a vulnerability in the guest, the hypervisor still isolates the host, as well as other applications/containers running on the host.

改进容器隔离的一种方法是在自己的虚拟机 (VM) 中运行每个容器。这为每个容器提供了自己的“机器”，包括内核和虚拟化设备，与主机完全分离。即使来宾中存在漏洞，虚拟机管理程序仍会隔离主机以及在主机上运行的其他应用程序/容器。

![vm-based-container-technology-gvisor2hu3.PNG](https://storage.googleapis.com/gweb-cloudblog-publish/images/vm-based-container-technology-gvisor2hu3.max-400x400.PNG)

Running containers in distinct VMs provides great isolation, compatibility, and performance, but may also require a larger resource footprint.

在不同的 VM 中运行容器提供了很好的隔离性、兼容性和性能，但也可能需要更大的资源占用。

[Kata containers](https://katacontainers.io/) is an open-source project that uses stripped-down VMs to keep the resource footprint minimal and maximize performance for container isolation. Like gVisor, Kata contains an [Open Container Initiative](https://www.opencontainers.org/)(OCI) runtime that is compatible with Docker and Kubernetes.

[Kata 容器](https://katacontainers.io/) 是一个开源项目，它使用精简的 VM 来保持最小的资源占用并最大化容器隔离的性能。与 gVisor 一样，Kata 包含一个与 Docker 和 Kubernetes 兼容的 [Open Container Initiative](https://www.opencontainers.org/)(OCI) 运行时。

### Sandboxed containers with gVisor

### 带有 gVisor 的沙盒容器

gVisor is more lightweight than a VM while maintaining a similar level of isolation. The core of gVisor is a kernel that runs as a normal, unprivileged process that supports most Linux system calls. This kernel is written in [Go](https://golang.org/), which was chosen for its memory- and type-safety. Just like within a VM, an application running in a gVisor sandbox gets its own kernel and set of virtualized devices, distinct from the host and other sandboxes.

gVisor 比 VM 更轻量级，同时保持类似的隔离级别。 gVisor 的核心是一个内核，它作为一个普通的、非特权的进程运行，支持大多数 Linux 系统调用。这个内核是用 [Go](https://golang.org/) 编写的，选择它是因为它的内存和类型安全。就像在 VM 中一样，在 gVisor 沙箱中运行的应用程序拥有自己的内核和一组虚拟化设备，与主机和其他沙箱不同。

![gvisor-sandbox-containerst5kz.PNG](https://storage.googleapis.com/gweb-cloudblog-publish/images/gvisor-sandbox-containerst5kz.max-500x500.PNG)



gVisor provides a strong isolation boundary by intercepting application system calls and acting as the guest kernel, all while running in user-space. Unlike a VM which requires a fixed set of resources on creation, gVisor can accommodate changing resources over time, as most normal Linux processes do. gVisor can be thought of as an extremely [paravirtualized operating system](https://en.wikipedia.org/wiki/Paravirtualization) with a flexible resource footprint and lower fixed cost than a full VM. However, this flexibility comes at the price of higher per-system call overhead and application compatibility—more on that below.

gVisor 通过拦截应用程序系统调用并充当访客内核，同时在用户空间中运行，从而提供强大的隔离边界。与创建时需要一组固定资源的 VM 不同，gVisor 可以适应随时间变化的资源，就像大多数普通 Linux 进程所做的那样。 gVisor 可以被认为是一个极其[半虚拟化的操作系统](https://en.wikipedia.org/wiki/Paravirtualization)，与完整的虚拟机相比，它具有灵活的资源占用和更低的固定成本。然而，这种灵活性是以更高的每系统调用开销和应用程序兼容性为代价的——更多内容见下文。

Secure workloads are a priority for the industry. We are encouraged to see innovative approaches like gVisor and look forward to collaborating on specification clarifications and making improvements to joint technical components in order to bring additional security to the ecosystem.**Samuel Ortiz**

安全工作负载是该行业的优先事项。我们很高兴看到 gVisor 等创新方法，并期待在规范澄清和改进联合技术组件方面进行合作，以便为生态系统带来额外的安全性。**Samuel Ortiz**

Member of the Kata Technical Steering Committee and Principal Engineer at Intel Corporation

英特尔公司 Kata 技术指导委员会成员和首席工程师

Hyper is encouraged to see gVisor’s novel approach to container isolation. The industry requires a robust ecosystem of secure container technologies, and we look forward to collaborating on gVisor to help bring secure containers into the mainstream.**Xu Wang**

鼓励 Hyper 看到 gVisor 的容器隔离新方法。行业需要一个强大的安全容器技术生态系统，我们期待与 gVisor 合作，帮助将安全容器带入主流。**Xu Wang**

Member of the Kata Technical Steering Committee and CTO at Hyper.sh

Kata 技术指导委员会成员和 Hyper.sh 的 CTO

### Integrated with Docker and Kubernetes

### 与 Docker 和 Kubernetes 集成

The gVisor runtime integrates seamlessly with Docker and Kubernetes though `runsc` (short for "run Sandboxed Container"), which conforms to the OCI runtime API.

gVisor 运行时通过符合 OCI 运行时 API 的 `runsc`（“run Sandboxed Container”的缩写）与 Docker 和 Kubernetes 无缝集成。

The `runsc` runtime is interchangeable with `runc`, Docker's default container runtime. Installation is simple; once installed it only takes a single additional flag to run a sandboxed container in Docker:

`runsc` 运行时可以与 Docker 的默认容器运行时 `runc` 互换。安装简单；一旦安装，它只需要一个额外的标志就可以在 Docker 中运行沙盒容器：

```ng-star-inserted
$ docker run --runtime=runsc hello-world

```


In Kubernetes, most resource isolation occurs at the pod level, making the pod a natural fit for a gVisor sandbox boundary. The Kubernetes community is currently [formalizing the sandbox pod API](https://goo.gl/eQHuqo), but experimental support is available today.

在 Kubernetes 中，大多数资源隔离发生在 pod 级别，这使得 pod 非常适合 gVisor 沙箱边界。 Kubernetes 社区目前正在 [正式化沙箱 pod API](https://goo.gl/eQHuqo)，但今天提供实验性支持。

The `runsc` runtime can run sandboxed pods in a Kubernetes cluster through the use of either the [cri-o](http://cri-o.io/) or [cri-containerd](https://github.com/containerd/cri) projects, which convert messages from the [Kubelet](https://kubernetes.io/docs/reference/generated/kubelet/) into OCI runtime commands.

`runsc` 运行时可以通过使用 [cri-o](http://cri-o.io/) 或 [cri-containerd](https://github.com/containerd/cri) 项目，将来自 [Kubelet](https://kubernetes.io/docs/reference/generated/kubelet/) 的消息转换为 OCI 运行时命令。

gVisor implements a large part of the Linux system API (200 system calls and counting), but not all. Some system calls and arguments are not currently supported, as are some parts of the /proc and /sys filesystems. As a result, not all applications will run inside gVisor, but many will run just fine, including Node.js, Java 8, MySQL, Jenkins, Apache, Redis, MongoDB, and many more.

gVisor 实现了 Linux 系统 API 的很大一部分（200 个系统调用和计数），但不是全部。当前不支持某些系统调用和参数，如 /proc 和 /sys 文件系统的某些部分。因此，并非所有应用程序都可以在 gVisor 中运行，但许多应用程序都可以正常运行，包括 Node.js、Java 8、MySQL、Jenkins、Apache、Redis、MongoDB 等等。

### Getting started

###  入门

As developers, we want the best of both worlds: the ease of use and portability of containers, and the resource isolation of VMs. We think gVisor is a great step in that direction. Check out our [repo on GitHub](https://github.com/google/gvisor) to find how to get started with gVisor and to learn more of the technical details behind it. And be sure to join our [Google group](https://groups.google.com/forum/#!forum/gvisor-users) to take part in the discussion!

作为开发人员，我们想要两全其美：容器的易用性和可移植性，以及虚拟机的资源隔离。我们认为 gVisor 是朝着这个方向迈出的重要一步。查看我们的 [GitHub 上的存储库](https://github.com/google/gvisor)，了解如何开始使用 gVisor 并了解更多其背后的技术细节。并且一定要加入我们的 [Google group](https://groups.google.com/forum/#!forum/gvisor-users) 来参与讨论！

If you’re at KubeCon in Copenhagen [join us at our booth](http://g.co/kubecon) for a deep dive demo and discussion.

如果您在哥本哈根的 KubeCon [加入我们的展位](http://g.co/kubecon) 进行深入的演示和讨论。

Also check out an interview with the gVisor PM to learn more.

另请查看对 gVisor PM 的采访以了解更多信息。

Learn more about gVisor, the new sandboxed container runtime via this demo. 

通过此演示了解有关 gVisor（新的沙盒容器运行时）的更多信息。

