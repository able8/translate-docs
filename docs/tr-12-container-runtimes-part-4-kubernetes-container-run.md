# Container Runtimes Part 4: Kubernetes Container Runtimes & CRI

# 容器运行时第 4 部分：Kubernetes 容器运行时和 CRI

*Jan 26, 2019*

*2019 年 1 月 26 日*

*This is the fourth and last part in a four part series on container runtimes. It's been a while since [part 1](https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r), but in that post I gave an overview of container runtimes and discussed the differences between low-level and high-level runtimes. In [part 2](https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level-contai) I went into detail on low-level container runtimes and built a simple low- level runtime. In [part 3](https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes) I went up the stack and wrote about high-level container runtimes.*

*这是关于容器运行时的四部分系列中的第四部分，也是最后一部分。距离 [第 1 部分](https://www.ianlewis.org/en/container-runtimes-part-1-introduction-container-r) 已经有一段时间了，但在那篇文章中，我概述了容器运行时并讨论了低级和高级运行时之间的差异。在 [第 2 部分](https://www.ianlewis.org/en/container-runtimes-part-2-anatomy-low-level-contai) 中，我详细介绍了低级容器运行时并构建了一个简单的低级容器级别运行时。在 [第 3 部分](https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes) 中，我在堆栈上写了关于高级容器运行时的文章。*

Kubernetes runtimes are high-level container runtimes that support the [Container Runtime Interface](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md) ( CRI). CRI was introduced in Kubernetes 1.5 and acts as a bridge between the [kubelet](https://kubernetes.io/docs/concepts/overview/components/#kubelet) and the container runtime. High-level container runtimes that want to  integrate with Kubernetes are expected to implement CRI. The runtime is  expected to handle the management of images and to support [Kubernetes pods](https://www.ianlewis.org/en/what-are-kubernetes-pods-anyway), as well as manage the individual containers so a Kubernetes runtime **must** be a high-level runtime per our definition in part 3. Low level  runtimes just don't have the necessary features. Since part 3 explains  all about high-level container runtimes, I'm going to focus on CRI and  introduce a few of the runtimes that support CRI in this post.

Kubernetes 运行时是支持 [Container Runtime Interface](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md) 的高级容器运行时（ CRI）。 CRI 是在 Kubernetes 1.5 中引入的，充当 [kubelet](https://kubernetes.io/docs/concepts/overview/components/#kubelet) 和容器运行时之间的桥梁。希望与 Kubernetes 集成的高级容器运行时需要实现 CRI。运行时预计将处理图像管理并支持 [Kubernetes pods](https://www.ianlewis.org/en/what-are-kubernetes-pods-anyway)，以及管理单个容器，以便根据我们在第 3 部分中的定义，Kubernetes 运行时**必须** 是高级运行时。低级运行时只是没有必要的功能。由于第 3 部分解释了有关高级容器运行时的所有内容，因此我将重点介绍 CRI 并在本文中介绍一些支持 CRI 的运行时。

In order to understand more about CRI it's worth taking a look at the overall Kubernetes architecture. The kubelet is an agent that sits on  each worker node in the Kubernetes cluster. The kubelet is responsible  for managing the container workloads for its node. When it comes to  actually run the workload, the kubelet uses CRI to communicate with the  container runtime running on that same node. In this way CRI is simply  an abstraction layer or API that allows you to switch out container  runtime implementations instead of having them built into the kubelet.

为了更多地了解 CRI，值得一看整个 Kubernetes 架构。 kubelet 是一个代理，位于 Kubernetes 集群中的每个工作节点上。 kubelet 负责管理其节点的容器工作负载。在实际运行工作负载时，kubelet 使用 CRI 与在同一节点上运行的容器运行时进行通信。通过这种方式，CRI 只是一个抽象层或 API，它允许您切换容器运行时实现，而不是将它们内置到 kubelet 中。

![Kubernetes architecture diagram](https://storage.googleapis.com/static.ianlewis.org/prod/img/772/CRI.png)

![Kubernetes 架构图](https://storage.googleapis.com/static.ianlewis.org/prod/img/772/CRI.png)

## Examples of CRI Runtimes

## CRI 运行时示例

Here are some CRI runtimes that can be used with Kubernetes.

以下是一些可与 Kubernetes 一起使用的 CRI 运行时。

### containerd

### 容器

`containerd` is a high-level runtime that I mentioned in Part 3. `containerd` is possibly the most popular CRI runtime currently. It implements CRI as a [plugin](https://github.com/containerd/cri) which is enabled by default. It listens on a unix socket by default so  you can configure crictl to connect to containerd like this:

`containerd` 是我在第 3 部分中提到的高级运行时。`containerd` 可能是目前最流行的 CRI 运行时。它将 CRI 实现为默认启用的 [插件](https://github.com/containerd/cri)。默认情况下，它侦听 unix 套接字，因此您可以像这样配置 crictl 以连接到 containerd：

```
 cat <<EOF | sudo tee /etc/crictl.yaml
 runtime-endpoint: unix:///run/containerd/containerd.sock
 EOF
 ```

 
It is an interesting high-level runtime in that it supports multiple  low-level runtimes via something called a "runtime handler" starting in  version 1.2. The runtime handler is passed via a field in CRI and based  on that runtime handler `containerd` runs an application  called a shim to start the container. This can be used to run containers using low-level runtimes other than runc, like [gVisor](https://github.com/google/gvisor), [Kata Containers](https://katacontainers.io/), or [Nabla Containers](https://nabla-containers.github.io/). The runtime handler is exposed in the Kubernetes API using the [RuntimeClass object](https://kubernetes.io/docs/concepts/containers/runtime-class/) which is alpha in Kubernetes 1.12. There is more on containerd's shim concept [here](https://github.com/containerd/containerd/pull/2434).

这是一个有趣的高级运行时，因为它从 1.2 版开始通过称为“运行时处理程序”的东西支持多个低级运行时。运行时处理程序通过 CRI 中的字段传递，并基于该运行时处理程序 `containerd` 运行一个名为 shim 的应用程序来启动容器。这可用于使用 runc 以外的低级运行时运行容器，例如 [gVisor](https://github.com/google/gvisor)、[Kata Containers](https://katacontainers.io/) 或[Nabla 容器](https://nabla-containers.github.io/)。运行时处理程序使用 [RuntimeClass 对象](https://kubernetes.io/docs/concepts/containers/runtime-class/) 在 Kubernetes API 中公开，这是 Kubernetes 1.12 中的 alpha。有更多关于 containerd 的 shim 概念 [here](https://github.com/containerd/containerd/pull/2434)。

### Docker

### 码头工人

Docker support for CRI was the first to be developed and was implemented as a shim between the `kubelet` and Docker. Docker has since broken out many of its features into `containerd` and now supports CRI through `containerd`. When modern versions of Docker are installed, `containerd` is installed along with it and CRI talks directly to `containerd`. For that reason, Docker itself isn't necessary to support CRI. So you  can install containerd directly or via Docker depending on your use  case.

Docker 对 CRI 的支持是第一个开发出来的，并作为 `kubelet` 和 Docker 之间的垫片来实现。此后，Docker 将其许多功能分解为“containerd”，现在通过“containerd”支持 CRI。当安装了现代版本的 Docker 时，`containerd` 也随之安装，CRI 直接与 `containerd` 对话。因此，Docker 本身不需要支持 CRI。因此，您可以根据您的用例直接或通过 Docker 安装 containerd。

### cri-o 
### 克里欧
cri-o is a lightweight CRI runtime made as a Kubernetes specific high-level runtime. It supports the management of [OCI compatible images](https://github.com/opencontainers/image-spec) and pulls from any OCI compatible image registry. It supports `runc` and Clear Containers as low-level runtimes. It supports other OCI  compatible low-level runtimes in theory, but relies on compatibility  with the `runc` [OCI command line interface](https://github.com/opencontainers/runtime-tools/blob/master/docs/command- line-interface.md), so in practice it isn't as flexible as `containerd`'s shim API.

cri-o 是一个轻量级的 CRI 运行时，作为 Kubernetes 特定的高级运行时。它支持管理 [OCI 兼容图像](https://github.com/opencontainers/image-spec) 并从任何 OCI 兼容图像注册表中提取。它支持 `runc` 和 Clear Containers 作为低级运行时。它理论上支持其他 OCI 兼容的低级运行时，但依赖于与 `runc` [OCI 命令行接口](https://github.com/opencontainers/runtime-tools/blob/master/docs/command- line-interface.md），所以在实践中它不像`containerd`的shim API那么灵活。

cri-o's endpoint is at `/var/run/crio/crio.sock` by default so you can configure `crictl` like so.

默认情况下，cri-o 的端点位于 `/var/run/crio/crio.sock`，因此您可以像这样配置 `crictl`。

```
 cat <<EOF | sudo tee /etc/crictl.yaml
 runtime-endpoint: unix:///var/run/crio/crio.sock
 EOF
 ```

 
## The CRI Specification

## CRI 规范

CRI is a [protocol buffers](https://developers.google.com/protocol-buffers/) and [gRPC](https://grpc.io/) API. The specification is defined in a [protobuf file](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.proto ) in the Kubernetes repository under the kubelet. CRI defines several  remote procedure calls (RPCs) and message types. The RPCs are for  operations like "pull image" (`ImageService.PullImage`), "create pod" (`RuntimeService.RunPodSandbox`), "create container" (`RuntimeService.CreateContainer`), "start container" (`RuntimeService. StartContainer`), "stop container" (`RuntimeService.StopContainer`), etc.

CRI 是一个 [protocol buffers](https://developers.google.com/protocol-buffers/) 和 [gRPC](https://grpc.io/) API。该规范定义在 [protobuf 文件](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.proto ) 在 kubelet 下的 Kubernetes 存储库中。 CRI 定义了几种远程过程调用 (RPC) 和消息类型。 RPC 用于诸如“拉图像”（`ImageService.PullImage`）、“创建 pod”（`RuntimeService.RunPodSandbox`）、“创建容器”（`RuntimeService.CreateContainer`）、“启动容器”（`RuntimeService. StartContainer`）、“停止容器”（`RuntimeService.StopContainer`）等。

For example, a typical interaction over CRI that starts a new  Kubernetes Pod would look something like the following (in my own form  of pseudo gRPC; each RPC would get a much bigger request object. I'm  simplifying it for brevity). The `RunPodSandbox` and `CreateContainer` RPCs return IDs in their responses which are used in subsequent requests:

例如，通过 CRI 启动一个新的 Kubernetes Pod 的典型交互如下所示（以我自己的伪 gRPC 形式；每个 RPC 将获得一个更大的请求对象。为了简洁起见，我将其简化）。 `RunPodSandbox` 和 `CreateContainer` RPC 在它们的响应中返回 ID，用于后续请求：

```
 ImageService.PullImage({image: "image1"})
 ImageService.PullImage({image: "image2"})
 podID = RuntimeService.RunPodSandbox({name: "mypod"})
 id1 = RuntimeService.CreateContainer({
     pod: podID,
     name: "container1",
     image: "image1",
 })
 id2 = RuntimeService.CreateContainer({
     pod: podID,
     name: "container2",
     image: "image2",
 })
 RuntimeService.StartContainer({id: id1})
 RuntimeService.StartContainer({id: id2})
 ```

 
We can interact with a CRI runtime directly using the `crictl` tool. `crictl` lets us send gRPC messages to a CRI runtime directly from the command  line. We can use this to debug and test out CRI implementations without  starting up a full-blown `kubelet` or Kubernetes cluster. You can get it by downloading a `crictl` binary from the cri-tools [releases page](https://github.com/kubernetes-sigs/cri-tools/releases) on GitHub.

我们可以使用 `crictl` 工具直接与 CRI 运行时交互。 `crictl` 允许我们直接从命令行将 gRPC 消息发送到 CRI 运行时。我们可以使用它来调试和测试 CRI 实现，而无需启动成熟的 `kubelet` 或 Kubernetes 集群。您可以通过从 GitHub 上的 cri-tools [发布页面](https://github.com/kubernetes-sigs/cri-tools/releases) 下载 `crictl` 二进制文件来获取它。

You can configure `crictl` by creating a configuration file under `/etc/crictl.yaml`. Here you should specify the runtime's gRPC endpoint as either a Unix socket file (`unix:///path/to/file`) or a TCP endpoint (`tcp://<host>:<port>`). We will use `containerd` for this example:

你可以通过在 `/etc/crictl.yaml` 下创建一个配置文件来配置 `crictl`。在这里，您应该将运行时的 gRPC 端点指定为 Unix 套接字文件（`unix:///path/to/file`）或 TCP 端点（`tcp://<host>:<port>`）。我们将在这个例子中使用`containerd`：

```
 cat <<EOF | sudo tee /etc/crictl.yaml
 runtime-endpoint: unix:///run/containerd/containerd.sock
 EOF
 ```

 
Or you can specify the runtime endpoint on each command line execution:

或者您可以在每个命令行执行时指定运行时端点：

```
 crictl --runtime-endpoint unix:///run/containerd/containerd.sock …
 ```

 
Let's run a pod with a single container with `crictl`. First you would tell the runtime to pull the `nginx` image you need since you can't start a container without the image stored locally.

让我们用 crictl 运行一个带有单个容器的 pod。首先，您会告诉运行时拉取您需要的“nginx”图像，因为如果没有本地存储的图像，您将无法启动容器。

```
 sudo crictl pull nginx
 ```

 
Next create a Pod creation request. You do this as a JSON file.

接下来创建一个 Pod 创建请求。您可以将其作为 JSON 文件执行。

```
 cat <<EOF | tee sandbox.json
 {
     "metadata": {
         "name": "nginx-sandbox",
         "namespace": "default",
         "attempt": 1,
         "uid": "hdishd83djaidwnduwk28bcsb"
     },
     "linux": {
     },
     "log_directory": "/tmp"
 }
 EOF
 ```

 
And then create the pod sandbox. We will store the ID of the sandbox as `SANDBOX_ID`.

然后创建 pod 沙箱。我们将沙箱的 ID 存储为“SANDBOX_ID”。

```
 SANDBOX_ID=$(sudo crictl runp --runtime runsc sandbox.json)
 ```

 
Next we will create a container creation request in a JSON file.

接下来我们将在 JSON 文件中创建容器创建请求。

```
 cat <<EOF | tee container.json
 {
   "metadata": {
       "name": "nginx"
     },
   "image":{
       "image": "nginx"
     },
   "log_path":"nginx.0.log",
   "linux": {
   }
 }
 EOF
 ```

 
We can then create and start the container inside the Pod we created earlier.

然后我们可以在之前创建的 Pod 中创建并启动容器。

```
 {
   CONTAINER_ID=$(sudo crictl create ${SANDBOX_ID} container.json sandbox.json)
   sudo crictl start ${CONTAINER_ID}
 }
 ```

 
You can inspect the running pod

您可以检查正在运行的 pod

```
 sudo crictl inspectp ${SANDBOX_ID}
 ```

 
… and the running container:

...和正在运行的容器：

```
 sudo crictl inspect ${CONTAINER_ID}
 ```

 
Clean up by stopping and deleting the container:

通过停止和删除容器来清理：

```
 {
   sudo crictl stop ${CONTAINER_ID}
   sudo crictl rm ${CONTAINER_ID}
 }
 ```

 
And then stop and delete the Pod:

然后停止并删除 Pod：

```
 {
   sudo crictl stopp ${SANDBOX_ID}
   sudo crictl rmp ${SANDBOX_ID}
 }
 ```

 
## Thanks for following the series! 
## 感谢您关注本系列！
This is the last post in the Container Runtimes series but don't  fear! There will be lots more container and Kubernetes posts in the  future. Be sure to add [my RSS feed](https://www.ianlewis.org/feed/enfeed) or follow me on Twitter to get notified when the next blog post comes out. 
这是容器运行时系列的最后一篇文章，但不要害怕！将来会有更多的容器和 Kubernetes 帖子。请务必添加 [我的 RSS 提要](https://www.ianlewis.org/feed/enfeed) 或在 Twitter 上关注我，以便在下一篇博文发布时收到通知。
