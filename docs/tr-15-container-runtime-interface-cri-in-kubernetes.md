# Introducing Container Runtime Interface (CRI) in Kubernetes

# 在 Kubernetes 中引入容器运行时接口 (CRI)

December 19, 2016 From: https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/

At the lowest layers of a Kubernetes node is the software that, among other things, starts and stops containers. We call this the “Container  Runtime”. The most widely known container runtime is Docker, but it is  not alone in this space. In fact, the container runtime space has been  rapidly evolving. As part of the effort to make Kubernetes more  extensible, we've been working on a new plugin API for container  runtimes in Kubernetes, called "CRI".

Kubernetes 节点的最低层是软件，其中包括启动和停止容器。我们称之为“容器运行时”。最广为人知的容器运行时是 Docker，但它在这个领域并不孤单。事实上，容器运行时空间一直在快速发展。作为使 Kubernetes 更具可扩展性的努力的一部分，我们一直在为 Kubernetes 中的容器运行时开发一个新的插件 API，称为“CRI”。

**What is the CRI and why does Kubernetes need it?**

**什么是 CRI，为什么 Kubernetes 需要它？**

Each container runtime has it own strengths, and many users have  asked for Kubernetes to support more runtimes. In the Kubernetes 1.5  release, we are proud to introduce the [Container Runtime Interface](https://github.com/kubernetes/kubernetes/blob/242a97307b34076d5d8f5bbeb154fa4d97c9ef1d/docs/devel/container-runtime-interface.md) (CRI) - - a plugin interface which enables kubelet to use a wide variety of container runtimes, without the need to recompile. CRI consists of a [protocol buffers](https://developers.google.com/protocol-buffers/) and [gRPC API](http://www.grpc.io/), and [libraries](https:/ /github.com/kubernetes/kubernetes/tree/release-1.5/pkg/kubelet/server/streaming), with additional specifications and tools under active development. CRI is being released as Alpha in [Kubernetes 1.5](https://kubernetes.io/blog/2016/12/kubernetes-1-5-supporting-production-workloads).

每个容器运行时都有自己的优势，许多用户要求 Kubernetes 支持更多的运行时。在 Kubernetes 1.5 版本中，我们很自豪地引入了[容器运行时接口](https://github.com/kubernetes/kubernetes/blob/242a97307b34076d5d8f5bbeb154fa4d97c9ef1d/docs/devel/container-runtime-interface.md) - 一个插件接口，使 kubelet 能够使用各种容器运行时，而无需重新编译。 CRI 由 [protocol buffers](https://developers.google.com/protocol-buffers/) 和 [gRPC API](http://www.grpc.io/) 和 [libraries](https://github.com/kubernetes/kubernetes/tree/release-1.5/pkg/kubelet/server/streaming)，其他规范和工具正在积极开发中。 CRI 正在 [Kubernetes 1.5](https://kubernetes.io/blog/2016/12/kubernetes-1-5-supporting-production-workloads) 中作为 Alpha 版发布。

Supporting interchangeable container runtimes is not a new concept in Kubernetes. In the 1.3 release, we announced the [rktnetes](https://kubernetes.io/blog/2016/07/rktnetes-brings-rkt-container-engine-to-Kubernetes) project to enable [rkt container engine](https ://github.com/coreos/rkt) as an alternative to the Docker container runtime. However, both Docker and rkt were integrated directly and deeply into the kubelet source  code through an internal and volatile interface. Such an integration  process requires a deep understanding of Kubelet internals and incurs  significant maintenance overhead to the Kubernetes community. These  factors form high barriers to entry for nascent container runtimes. By  providing a clearly-defined abstraction layer, we eliminate the barriers and allow developers to focus on building their container runtimes. This is a small, yet important step towards truly enabling pluggable  container runtimes and building a healthier ecosystem.

支持可互换的容器运行时在 Kubernetes 中并不是一个新概念。在 1.3 版本中，我们宣布了 [rktnetes](https://kubernetes.io/blog/2016/07/rktnetes-brings-rkt-container-engine-to-Kubernetes) 项目来启用 [rkt 容器引擎](https ://github.com/coreos/rkt) 作为 Docker 容器运行时的替代方案。但是，Docker 和 rkt 都通过内部易变的接口直接并深入地集成到 kubelet 源代码中。这样的集成过程需要对 Kubelet 内部结构有深入的了解，并且会给 Kubernetes 社区带来大量的维护开销。这些因素构成了新生容器运行时进入的高壁垒。通过提供明确定义的抽象层，我们消除了障碍并允许开发人员专注于构建他们的容器运行时。这是真正实现可插拔容器运行时和构建更健康生态系统的一个小而重要的步骤。

**Overview of CRI**

**CRI 概述**

Kubelet communicates with the container runtime (or a CRI shim for the  runtime) over Unix sockets using the gRPC framework, where kubelet acts  as a client and the CRI shim as the server.

Kubelet 使用 gRPC 框架通过 Unix 套接字与容器运行时（或运行时的 CRI shim）通信，其中 kubelet 作为客户端，CRI shim 作为服务器。

[![img](https://cl.ly/3I2p0D1V0T26/Image%202016-12-19%20at%2017.13.16.png)](https://cl.ly/3I2p0D1V0T26/Image 2016-12-19 at 17.13.16.png)

The protocol buffers [API](https://github.com/kubernetes/kubernetes/blob/release-1.5/pkg/kubelet/api/v1alpha1/runtime/api.proto) includes two gRPC services, ImageService, and RuntimeService. The  ImageService provides RPCs to pull an image from a repository, inspect,  and remove an image. The RuntimeService contains RPCs to manage the  lifecycle of the pods and containers, as well as calls to interact with  containers (exec/attach/port-forward). A monolithic container runtime  that manages both images and containers (e.g., Docker and rkt) can  provide both services simultaneously with a single socket. The sockets  can be set in Kubelet by --container-runtime-endpoint and  --image-service-endpoint flags.

协议缓冲区 [API](https://github.com/kubernetes/kubernetes/blob/release-1.5/pkg/kubelet/api/v1alpha1/runtime/api.proto) 包括两个 gRPC 服务，ImageService 和 RuntimeService。 ImageService 提供 RPC 以从存储库中提取图像、检查和删除图像。 RuntimeService 包含 RPC 来管理 Pod 和容器的生命周期，以及与容器交互的调用（exec/attach/port-forward）。管理镜像和容器（例如 Docker 和 rkt）的单体容器运行时可以使用单个套接字同时提供这两种服务。套接字可以通过 --container-runtime-endpoint 和 --image-service-endpoint 标志在 Kubelet 中设置。

**Pod and container lifecycle management**


**Pod 和容器生命周期管理**


```
 service RuntimeService {
     // Sandbox operations.
     rpc RunPodSandbox(RunPodSandboxRequest) returns (RunPodSandboxResponse) {}
     rpc StopPodSandbox(StopPodSandboxRequest) returns (StopPodSandboxResponse) {}
     rpc RemovePodSandbox(RemovePodSandboxRequest) returns (RemovePodSandboxResponse) {}
     rpc PodSandboxStatus(PodSandboxStatusRequest) returns (PodSandboxStatusResponse) {}
     rpc ListPodSandbox(ListPodSandboxRequest) returns (ListPodSandboxResponse) {}

     // Container operations.
     rpc CreateContainer(CreateContainerRequest) returns (CreateContainerResponse) {}
     rpc StartContainer(StartContainerRequest) returns (StartContainerResponse) {}
     rpc StopContainer(StopContainerRequest) returns (StopContainerResponse) {}
     rpc RemoveContainer(RemoveContainerRequest) returns (RemoveContainerResponse) {}
     rpc ListContainers(ListContainersRequest) returns (ListContainersResponse) {}
     rpc ContainerStatus(ContainerStatusRequest) returns (ContainerStatusResponse) {}
     ...
 }
```


A Pod is composed of a group of application containers  in an isolated environment with resource constraints. In CRI, this  environment is called PodSandbox. We intentionally leave some room for  the container runtimes to interpret the PodSandbox differently based on  how they operate internally. For hypervisor-based runtimes, PodSandbox  might represent a virtual machine. For others, such as Docker, it might  be Linux namespaces. The PodSandbox must respect the pod resources  specifications. In the v1alpha1 API, this is achieved by launching all  the processes within the pod-level cgroup that kubelet creates and  passes to the runtime.

一个 Pod 由一组处于资源约束的隔离环境中的应用容器组成。在 CRI 中，这种环境称为 PodSandbox。我们有意为容器运行时留出一些空间，以便根据它们在内部运行的方式对 PodSandbox 进行不同的解释。对于基于管理程序的运行时，PodSandbox 可能代表一个虚拟机。对于其他人，例如 Docker，它可能是 Linux 命名空间。 PodSandbox 必须遵守 Pod 资源规范。在 v1alpha1 API 中，这是通过启动 kubelet 创建并传递给运行时的 pod 级 cgroup 中的所有进程来实现的。

Before starting a pod, kubelet calls RuntimeService.RunPodSandbox to  create the environment. This includes setting up networking for a pod  (e.g., allocating an IP). Once the PodSandbox is active, individual  containers can be created/started/stopped/removed independently. To  delete the pod, kubelet will stop and remove containers before stopping  and removing the PodSandbox.

在启动 Pod 之前，kubelet 调用 RuntimeService.RunPodSandbox 来创建环境。这包括为 Pod 设置网络（例如，分配 IP）。一旦 PodSandbox 处于活动状态，就可以独立地创建/启动/停止/删除单个容器。要删除 Pod，kubelet 将在停止和删除 PodSandbox 之前停止并删除容器。

Kubelet is responsible for managing the lifecycles of the containers  through the RPCs, exercising the container lifecycles hooks and  liveness/readiness checks, while adhering to the restart policy of the  pod.

Kubelet 负责通过 RPC 管理容器的生命周期，执行容器生命周期钩子和活性/就绪检查，同时遵守 Pod 的重启策略。

**Why an imperative container-centric interface?**

**为什么需要一个以容器为中心的命令式界面？**

Kubernetes has a declarative API with a *Pod* resource. One possible design we considered was for CRI to reuse the declarative *Pod* object in its abstraction, giving the container runtime freedom to  implement and exercise its own control logic to achieve the desired  state. This would have greatly simplified the API and allowed CRI to  work with a wider spectrum of runtimes. We discussed this approach early in the design phase and decided against it for several reasons. First,  there are many Pod-level features and specific mechanisms (e.g., the  crash-loop backoff logic) in kubelet that would be a significant burden  for all runtimes to reimplement. Second, and more importantly, the Pod  specification was (and is) still evolving rapidly. Many of the new  features (e.g., init containers) would not require any changes to the  underlying container runtimes, as long as the kubelet manages containers directly. CRI adopts an imperative container-level interface so that  runtimes can share these common features for better development  velocity. This doesn't mean we're deviating from the "level triggered"  philosophy - kubelet is responsible for ensuring that the actual state  is driven towards the declared state.

Kubernetes 有一个带有 *Pod* 资源的声明式 API。我们考虑的一种可能的设计是让 CRI 在其抽象中重用声明性 *Pod* 对象，让容器运行时自由地实现和运用自己的控制逻辑来实现所需的状态。这将大大简化 API 并允许 CRI 使用更广泛的运行时。我们在设计阶段的早期就讨论了这种方法，但出于多种原因决定不使用它。首先，kubelet 中有许多 Pod 级别的特性和特定机制（例如，崩溃循环退避逻辑），这对于所有运行时重新实现来说都是一个巨大的负担。其次，更重要的是，Pod 规范过去（并且现在）仍在快速发展。许多新功能（例如 init 容器）不需要对底层容器运行时进行任何更改，只要 kubelet 直接管理容器即可。 CRI 采用命令式容器级接口，以便运行时可以共享这些通用功能以提高开发速度。这并不意味着我们偏离了“级别触发”的理念——kubelet 负责确保实际状态被驱动到声明的状态。

**Exec/attach/port-forward requests**

**执行/附加/端口转发请求**

```
 service RuntimeService {
     ...
     // ExecSync runs a command in a container synchronously.
     rpc ExecSync(ExecSyncRequest) returns (ExecSyncResponse) {}
     // Exec prepares a streaming endpoint to execute a command in the container.
     rpc Exec(ExecRequest) returns (ExecResponse) {}
     // Attach prepares a streaming endpoint to attach to a running container.
     rpc Attach(AttachRequest) returns (AttachResponse) {}
     // PortForward prepares a streaming endpoint to forward ports from a PodSandbox.
     rpc PortForward(PortForwardRequest) returns (PortForwardResponse) {}
     ...
 }
```

Kubernetes provides features (e.g. kubectl  exec/attach/port-forward) for users to interact with a pod and the  containers in it. Kubelet today supports these features either by  invoking the container runtime’s native method calls or by using the  tools available on the node (e.g., nsenter and socat). Using tools on  the node is not a portable solution because most tools assume the pod is isolated using Linux namespaces. In CRI, we explicitly define these  calls in the API to allow runtime-specific implementations. 

Kubernetes 提供了一些功能（例如 kubectl exec/attach/port-forward）供用户与 pod 及其中的容器进行交互。如今，Kubelet 通过调用容器运行时的本地方法调用或使用节点上可用的工具（例如 nsenter 和 socat）来支持这些功能。在节点上使用工具不是一种可移植的解决方案，因为大多数工具都假设 pod 是使用 Linux 命名空间隔离的。在 CRI 中，我们在 API 中明确定义了这些调用，以允许特定于运行时的实现。

Another potential issue with the kubelet implementation today is that kubelet handles the connection of all streaming requests, so it can  become a bottleneck for the network traffic on the node. When designing  CRI, we incorporated this feedback to allow runtimes to eliminate the  middleman. The container runtime can start a separate streaming server  upon request (and can potentially account the resource usage to the  pod!), and return the location of the server to kubelet. Kubelet then  returns this information to the Kubernetes API server, which opens a  streaming connection directly to the runtime-provided server and  connects it to the client.

今天 kubelet 实现的另一个潜在问题是 kubelet 处理所有流请求的连接，因此它可能成为节点上网络流量的瓶颈。在设计 CRI 时，我们结合了此反馈以允许运行时消除中间人。容器运行时可以根据请求启动一个单独的流服务器（并且可以潜在地将资源使用情况记入 pod！），并将服务器的位置返回给 kubelet。 Kubelet 然后将此信息返回给 Kubernetes API 服务器，后者直接打开与运行时提供的服务器的流连接并将其连接到客户端。

There are many other aspects of CRI that are not covered in this blog post. Please see the list of [design docs and proposals](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md#design-docs-and- proposals) for all the details.

本博文未涵盖 CRI 的许多其他方面。请参阅[设计文档和提案](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md#design-docs-and-建议）的所有细节。

**Current status**

**当前状态**

Although CRI is still in its early stages, there are already several  projects under development to integrate container runtimes using CRI. Below are a few examples:
 - [cri-o](https://cri-o.io/): OCI conformant runtimes.
 - [rktlet](https://github.com/kubernetes-incubator/rktlet): the rkt container runtime.
 - [frakti](https://github.com/kubernetes/frakti): hypervisor-based container runtimes.
 - [docker CRI shim](https://github.com/kubernetes/kubernetes/tree/release-1.5/pkg/kubelet/dockershim).

尽管 CRI 仍处于早期阶段，但已经有几个项目正在开发中，以使用 CRI 集成容器运行时。下面是几个例子：
- [cri-o](https://cri-o.io/)：符合 OCI 的运行时。
- [rktlet](https://github.com/kubernetes-incubator/rktlet)：rkt 容器运行时。
- [frakti](https://github.com/kubernetes/frakti)：基于管理程序的容器运行时。
- [docker CRI shim](https://github.com/kubernetes/kubernetes/tree/release-1.5/pkg/kubelet/dockershim)。

If you are interested in trying these alternative runtimes, you can  follow the individual repositories for the latest progress and  instructions.

如果您有兴趣尝试这些替代运行时，您可以关注各个存储库以获取最新进展和说明。

For developers interested in integrating a new container runtime, please see the [developer guide](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md) for the known limitations and issues of the API. We are actively  incorporating feedback from early developers to improve the API. Developers should expect occasional API breaking changes (it is Alpha,  after all).

对于有兴趣集成新容器运行时的开发人员，请参阅[开发人员指南](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md)了解 API 的已知限制和问题。我们正在积极整合早期开发人员的反馈以改进 API。开发人员应该期待偶尔的 API 重大更改（毕竟它是 Alpha）。

**Try the new CRI-Docker integration**

**尝试新的 CRI-Docker 集成**

Kubelet does not yet use CRI by default, but we are actively working  on making this happen. The first step is to re-integrate Docker with  kubelet using CRI. In the 1.5 release, we extended kubelet to support  CRI, and also added a built-in CRI shim for Docker. This allows kubelet  to start the gRPC server on Docker’s behalf. To try out the new  kubelet-CRI-Docker integration, you simply have to start the Kubernetes  API server with --feature-gates=StreamingProxyRedirects=true to enable  the new streaming redirect feature, and then start the kubelet with  --experimental-cri =true.

Kubelet 尚未默认使用 CRI，但我们正在积极致力于实现这一目标。第一步是使用 CRI 将 Docker 与 kubelet 重新集成。在 1.5 版本中，我们扩展了 kubelet 以支持 CRI，并且还为 Docker 添加了一个内置的 CRI shim。这允许 kubelet 代表 Docker 启动 gRPC 服务器。要试用新的 kubelet-CRI-Docker 集成，您只需使用 --feature-gates=StreamingProxyRedirects=true 启动 Kubernetes API 服务器以启用新的流重定向功能，然后使用 --experimental-cri 启动 kubelet =真。

Besides a few [missing features](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md#docker-cri-integration-known-issues) , the new integration has consistently passed the main end-to-end tests. We plan to expand the test coverage soon and would like to encourage the community to report any issues to help with the transition.

除了一些[缺失的功能](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md#docker-cri-integration-known-issues)，新的集成一直通过主要的端到端测试。我们计划很快扩大测试覆盖范围，并希望鼓励社区报告任何问题以帮助过渡。

Community

社区

CRI is being actively developed and maintained by the Kubernetes [SIG-Node](https://github.com/kubernetes/community/blob/master/README.md#special-interest-groups-sig) community. We’d love to hear feedback from you. To join the community:
 - Post issues or feature requests on [GitHub](https://github.com/kubernetes/kubernetes)
 - Join the #sig-node channel on [Slack](https://kubernetes.slack.com/)
 - Subscribe to the [SIG-Node mailing list](mailto:kubernetes-sig-node@googlegroups.com)
 - Follow us on Twitter [@Kubernetesio](https://twitter.com/kubernetesio) for latest updates

Kubernetes [SIG-Node](https://github.com/kubernetes/community/blob/master/README.md#special-interest-groups-sig) 社区正在积极开发和维护 CRI。我们很乐意听取您的反馈。加入社区：
- 在 [GitHub](https://github.com/kubernetes/kubernetes) 上发布问题或功能请求
- 加入 [Slack](https://kubernetes.slack.com/) 上的 #sig-node 频道
- 订阅 [SIG-Node 邮件列表](mailto:kubernetes-sig-node@googlegroups.com)
- 在 Twitter [@Kubernetesio](https://twitter.com/kubernetesio) 上关注我们以获取最新更新

*--Yu-Ju Hong, Software Engineer, Google* 

