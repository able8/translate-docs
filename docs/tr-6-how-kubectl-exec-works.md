# How does 'kubectl exec' work?

# 'kubectl exec' 是如何工作的？

Aug 25, 2019 · [3 Comments](https://erkanerol.github.io/post/how-kubectl-exec-works/#disqus_thread)

2019 年 8 月 25 日 · [3 条评论](https://erkanerol.github.io/post/how-kubectl-exec-works/#disqus_thread)

Last Friday, one of my colleagues approached me and asked a question about  how to exec a command in a pod with client-go. I didn’t know the answer  and I noticed that I had never thought about the mechanism in “kubectl  exec”. I had some ideas about how it should be, but I wasn’t 100% sure. I noted the topic to check again and I have learnt a lot after reading  some blogs, docs and source codes. In this blog post, I am going to  share my understanding and findings.

上周五，我的一位同事来找我，问了一个关于如何使用 client-go 在 pod 中执行命令的问题。我不知道答案，我注意到我从未想过“kubectl exec”中的机制。我有一些关于它应该如何的想法，但我不是 100% 确定。我记下了要再次检查的主题，在阅读了一些博客、文档和源代码后，我学到了很多东西。在这篇博文中，我将分享我的理解和发现。

> Please ping me if there is something wrong. https://twitter.com/erkan_erol_

> 如果有什么问题，请 ping 我。 https://twitter.com/erkan_erol_

## Setup

##  设置

I cloned https://github.com/ecomm-integration-ballerina/kubernetes-cluster in order to create a k8s cluster in my MacBook. I fixed IP addresses of the nodes in kubelet configurations since the default configuration  didn’t let me run `kubectl exec`. You can find the root cause [here](https://medium.com/@joatmon08/playing-with-kubeadm-in-vagrant-machines-part-2-bac431095706).

我克隆了 https://github.com/ecomm-integration-ballerina/kubernetes-cluster 以便在我的 MacBook 中创建一个 k8s 集群。我在 kubelet 配置中修复了节点的 IP 地址，因为默认配置不允许我运行 `kubectl exec`。你可以在这里找到根本原因（https://medium.com/@joatmon08/playing-with-kubeadm-in-vagrant-machines-part-2-bac431095706）。

- Any machine = my MacBook
 - IP of master node = 192.168.205.10
 - IP of worker node = 192.168.205.11
 - API server port = 6443

- 任何机器 = 我的 MacBook
- 主节点的 IP = 192.168.205.10
- 工作节点的 IP = 192.168.205.11
- API 服务器端口 = 6443

## Components

##  组件

![Components](https://erkanerol.github.io/img/kubectl-exec/components.png)

![组件](https://erkanerol.github.io/img/kubectl-exec/components.png)

- ***kubectl exec process:*** When we run “kubectl exec …” in a machine, a process starts. You can  run it in any machine which has an access to k8s api server.
 - ***[api server](https://kubernetes.io/docs/concepts/overview/components/#kube-apiserver):*** Component on the master that exposes the Kubernetes API. It is the front-end for the Kubernetes control plane.
 - ***[kubelet](https://kubernetes.io/docs/concepts/overview/components/#kubelet):*** An agent that runs on each node in the cluster. It makes sure that containers are running in a pod.
 - ***[container runtime](https://kubernetes.io/docs/concepts/overview/components/#container-runtime):*** The software that is responsible for running containers. Examples: docker, cri-o, containerd…
 - ***kernel:*** kernel of the OS in the worker node which is responsible to manage processes.
 - ***target container:*** A container which is a part of a pod and which is running on one of the worker nodes.

- ***kubectl exec process:*** 当我们在一台机器上运行“kubectl exec ...”时，一个进程启动。您可以在任何可以访问 k8s api 服务器的机器上运行它。
- ***[api server](https://kubernetes.io/docs/concepts/overview/components/#kube-apiserver):*** 暴露 Kubernetes API 的 master 上的组件。它是 Kubernetes 控制平面的前端。
- ***[kubelet](https://kubernetes.io/docs/concepts/overview/components/#kubelet):*** 在集群中的每个节点上运行的代理。它确保容器在 pod 中运行。
- ***[容器运行时](https://kubernetes.io/docs/concepts/overview/components/#container-runtime):*** 负责运行容器的软件。示例：docker、cri-o、containerd…
- ***kernel:*** 工作节点中操作系统的内核，负责管理进程。
- ***目标容器：*** 一个容器，它是 pod 的一部分，并且在其中一个工作节点上运行。

## Findings

##  发现

### 1. Activities in Client Side

### 1. 客户端活动

- Create a pod in default namespace

- 在默认命名空间中创建一个 pod

```
   // any machine
   $ kubectl run exec-test-nginx --image=nginx
 ```

 
- Then run an exec command and `sleep 5000` to make observation

- 然后运行 ​​exec 命令和 `sleep 5000` 进行观察

```
   // any machine
   $ kubectl exec -it exec-test-nginx-6558988d5-fgxgg -- sh
   # sleep 5000
 ```

 
- We can observe the kubectl process (pid=8507 in this case)

- 我们可以观察到 kubectl 进程（这里 pid=8507）

```
   // any machine
   $ ps -ef |grep kubectl
   501  8507  8409   0  7:19PM ttys000    0:00.13 kubectl exec -it exec-test-nginx-6558988d5-fgxgg -- sh
 ```

 
- When we check network activities of the process, we can see that it has some connections to api-server (192.168.205.10.6443)

- 当我们检查进程的网络活动时，我们可以看到它与api-server（192.168.205.10.6443）有一些连接

```
   // any machine
   $ netstat -atnv |grep 8507
   tcp4       0      0  192.168.205.1.51673    192.168.205.10.6443    ESTABLISHED 131072 131768   8507      0 0x0102 0x00000020
   tcp4       0      0  192.168.205.1.51672    192.168.205.10.6443    ESTABLISHED 131072 131768   8507      0 0x0102 0x00000028
 ```

 
- Let’s check the code. kubectl creates a POST request with subresource `exec` and sends a rest request.

- 让我们检查代码。 kubectl 使用子资源 `exec` 创建一个 POST 请求并发送一个休息请求。

```
 req := restClient.Post().
         Resource("pods").
         Name(pod.Name).
         Namespace(pod.Namespace).
         SubResource("exec")
 req.VersionedParams(&corev1.PodExecOptions{
         Container: containerName,
         Command:   p.Command,
         Stdin:     p.Stdin,
         Stdout:    p.Out != nil,
         Stderr:    p.ErrOut != nil,
         TTY:       t.Raw,
 }, scheme.ParameterCodec)

 return p.Executor.Execute("POST", req.URL(), p.Config, p.In, p.Out, p.ErrOut, t.Raw, sizeQueue)
 ```

 
[view raw](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/staging/src/k8s.io/kubectl/pkg/cmd/exec/exec.go)                    [staging/src/k8s.io/ kubectl/pkg/cmd/exec/exec.go](https://github.com/kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/staging/src/k8s.io/kubectl/pkg/cmd/exec/exec.go)

  

[查看原始数据](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/staging/src/k8s.io/kubectl/pkg/cmd/exec/exec.go) [staging/src/k kubectl/pkg/cmd/exec/exec.go](https://github.com/kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/staging/src/k8s.io/kubectl/pkg.execgo)/

  

![rest-request](https://erkanerol.github.io/img/kubectl-exec/rest-request.png)

![rest-request](https://erkanerol.github.io/img/kubectl-exec/rest-request.png)

### 2. Activities in Master Node

### 2. 主节点中的活动

- We can observe the request in api-server side as well.

- 我们也可以在 api-server 端观察请求。

```
   handler.go:143] kube-apiserver: POST "/api/v1/namespaces/default/pods/exec-test-nginx-6558988d5-fgxgg/exec" satisfied by gorestful with webservice /api/v1
   upgradeaware.go:261] Connecting to backend proxy (intercepting redirects) https://192.168.205.11:10250/exec/default/exec-test-nginx-6558988d5-fgxgg/exec-test-nginx?command=sh&input=1&output= 1&tty=1 
``
  
Headers: map[Connection:[Upgrade] Content-Length:[0] Upgrade:[SPDY/3.1] User-Agent:[kubectl/v1.12.10 (darwin/amd64) kubernetes/e3c1340] X-Forwarded-For:[192.168 .205.1] X-Stream-Protocol-Version:[v4.channel.k8s.io v3.channel.k8s.io v2.channel.k8s.io channel.k8s.io]]
 ```

标头：map[Connection:[Upgrade] Content-Length:[0] Upgrade:[SPDY/3.1] User-Agent:[kubectl/v1.12.10 (darwin/amd64) kubernetes/e3c1340] X-Forwarded-For:[192.168 .205.1] X-Stream-Protocol-Version:[v4.channel.k8s.io v3.channel.k8s.io v2.channel.k8s.io channel.k8s.io]]
``

  > Notice that the http request includes a protocol upgrade request. [SPDY](https://www.wikiwand.com/en/SPDY) allows for separate stdin/stdout/stderr/spdy-error “streams” to be multiplexed over a single TCP connection.

> 请注意，http 请求包含协议升级请求。 [SPDY](https://www.wikiwand.com/en/SPDY) 允许在单个 TCP 连接上多路复用单独的 stdin/stdout/stderr/spdy-error “流”。

- Api server receives the request and binds it into a `PodExecOptions`

- Api 服务器接收请求并将其绑定到`PodExecOptions`

```
   // PodExecOptions is the query options to a Pod's remote exec call
   type PodExecOptions struct {
           metav1.TypeMeta
  
           // Stdin if true indicates that stdin is to be redirected for the exec call
           Stdin bool
  
           // Stdout if true indicates that stdout is to be redirected for the exec call
           Stdout bool
  
           // Stderr if true indicates that stderr is to be redirected for the exec call
           Stderr bool
  
           // TTY if true indicates that a tty will be allocated for the exec call
           TTY bool
  
           // Container in which to execute the command.
           Container string
  
           // Command is the remote command to execute; argv array; not executed within a shell.
           Command []string
   }
 ```

 
[view raw](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/apis/core/types.go)                    [pkg/apis/core/types.go](https://github.com /kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/apis/core/types.go)

  

[查看原始数据](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/apis/core/types.go) [pkg/apis/core/types.go](https://github.com /kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/apis/core/types.go)

  

- To be able to take necessary actions, api-server needs to know which location it should contact.

- 为了能够采取必要的行动，api-server 需要知道它应该联系哪个位置。

```
   // ExecLocation returns the exec URL for a pod container. If opts.Container is blank
   // and only one container is present in the pod, that container is used.
   func ExecLocation(
           getter ResourceGetter,
           connInfo client.ConnectionInfoGetter,
           ctx context.Context,
           name string,
           opts *api.PodExecOptions,
   ) (*url.URL, http.RoundTripper, error) {
           return streamLocation(getter, connInfo, ctx, name, opts, opts.Container, "exec")
   }
 ```

 
[view raw](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/strategy.go)                    [pkg/registry/core/pod/strategy.go](https:/ /github.com/kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/strategy.go)

  

[查看原始数据](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/strategy.go) [pkg/registry/core/pod/strategy/go]( /github.com/kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/strategy.go)

  

  Of course the endpoint is derived from node info.

当然，端点来自节点信息。

```
           nodeName := types.NodeName(pod.Spec.NodeName)
           if len(nodeName) == 0 {
                   // If pod has not been assigned a host, return an empty location
                   return nil, nil, errors.NewBadRequest(fmt.Sprintf("pod %s does not have a host assigned", name))
           }
           nodeInfo, err := connInfo.GetConnectionInfo(ctx, nodeName)
 ```

 
[view raw](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/strategy.go)                    [pkg/registry/core/pod/strategy.go](https:/ /github.com/kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/strategy.go)

  

[查看原始数据](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/strategy.go) [pkg/registry/core/pod/strategy/go]( /github.com/kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/strategy.go)

  

  GOTCHA! KUBELET HAS A PORT (`node.Status.DaemonEndpoints.KubeletEndpoint.Port`) TO WHICH API-SERVER CAN CONNECT.

明白了！ KUBELET 有一个端口（`node.Status.DaemonEndpoints.KubeletEndpoint.Port`），API-SERVER 可以连接到该端口。

```
   // GetConnectionInfo retrieves connection info from the status of a Node API object.
   func (k *NodeConnectionInfoGetter) GetConnectionInfo(ctx context.Context, nodeName types.NodeName) (*ConnectionInfo, error) {
           node, err := k.nodes.Get(ctx, string(nodeName), metav1.GetOptions{})
           if err != nil {
                   return nil, err
           }
  
           // Find a kubelet-reported address, using preferred address type
           host, err := nodeutil.GetPreferredNodeAddress(node, k.preferredAddressTypes)
           if err != nil {
                   return nil, err
           }
  
           // Use the kubelet-reported port, if present
           port := int(node.Status.DaemonEndpoints.KubeletEndpoint.Port)
           if port <= 0 {
                   port = k.defaultPort
           }
  
           return &ConnectionInfo{
                   Scheme:    k.scheme,
                   Hostname:  host,
                   Port:      strconv.Itoa(port),
                   Transport: k.transport,
           }, nil
   }
 ``` 


[view raw](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/kubelet/client/kubelet_client.go)                    [pkg/kubelet/client/kubelet_client.go](https://github.com /kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/kubelet/client/kubelet_client.go)

  

[查看原始数据](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/kubelet/client/kubelet_client.go) [pkg/kubelet/client/kubelet_client.com](https://github.com) /kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/kubelet/client/kubelet_client.go)

  

  > [Master-Node Communication > Master to Cluster > apiserver to kubelet](https://kubernetes.io/docs/concepts/architecture/master-node-communication/#apiserver-to-kubelet)
   >
   > These connections terminate at the kubelet’s HTTPS endpoint. By default, the  apiserver does not verify the kubelet’s serving certificate, which makes the connection subject to man-in-the-middle attacks, and ***unsafe\*** to run over untrusted and/or public networks.

> [Master-Node 通信 > Master to Cluster > apiserver to kubelet](https://kubernetes.io/docs/concepts/architecture/master-node-communication/#apiserver-to-kubelet)
  >
  > 这些连接终止于 kubelet 的 HTTPS 端点。默认情况下，apiserver 不验证 kubelet 的服务证书，这使得连接容易受到中间人攻击，并且 ***unsafe\*** 在不受信任和/或公共网络上运行。

- Now, api server knows the endpoint and it opens a connections.

- 现在，api 服务器知道端点并打开连接。

```
   // Connect returns a handler for the pod exec proxy
   func (r *ExecREST) Connect(ctx context.Context, name string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
           execOpts, ok := opts.(*api.PodExecOptions)
           if !ok {
                   return nil, fmt.Errorf("invalid options object: %#v", opts)
           }
           location, transport, err := pod.ExecLocation(r.Store, r.KubeletConn, ctx, name, execOpts)
           if err != nil {
                   return nil, err
           }
           return newThrottledUpgradeAwareProxyHandler(location, transport, false, true, true, responder), nil
   }
 ```

 
[view raw](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/rest/subresources.go)                    [pkg/registry/core/pod/rest/subresources.go] (https://github.com/kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/rest/subresources.go)

  

[查看原始数据](https://github.com/kubernetes/kubernetes/raw/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/rest/subresources.go) [pkg/registry/subsource/pods.rest/] （https://github.com/kubernetes/kubernetes/blob/a1f1f0b599e961a5c59b02c349c0ed818b1851a5/pkg/registry/core/pod/rest/subresources.go）

  

- Let’s check what is going on the master node.

- 让我们检查一下主节点上发生了什么。

First, learn the ip of the worker node. It is `192.168.205.11` in this case.

首先，了解worker节点的ip。在这种情况下是“192.168.205.11”。

```
 // any machine
 $ kubectl get nodes k8s-node-1 -o wide
 NAME         STATUS   ROLES    AGE   VERSION   INTERNAL-IP      EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION      CONTAINER-RUNTIME
 k8s-node-1   Ready    <none>   9h    v1.15.3   192.168.205.11   <none>        Ubuntu 16.04.6 LTS   4.4.0-159-generic   docker://17.3.3
 ```

 
Then get the kubelet port. It is `10250` in this case.

然后获取kubelet端口。在这种情况下它是“10250”。

```
 // any machine
 $ kubectl get nodes k8s-node-1 -o jsonpath='{.status.daemonEndpoints.kubeletEndpoint}'
 map[Port:10250]
 ```

 
Then check the network. Is there a connection to worker  node(192.168.205.11)? THE CONNECTİON IS THERE. When I kill the exec  process, it disappears so I know it is set by api-server because of my  exec command.

然后检查网络。是否有到工作节点（192.168.205.11）的连接？连接就在那里。当我杀死 exec 进程时，它消失了，所以我知道它是由 api-server 设置的，因为我的 exec 命令。

```
 // master node
 $ netstat -atn |grep 192.168.205.11
 tcp        0      0 192.168.205.10:37870    192.168.205.11:10250    ESTABLISHED
 ...
 ```

 
![api-server-to-kubelet](https://erkanerol.github.io/img/kubectl-exec/api-server-to-kubelet.png)

![api-server-to-kubelet](https://erkanerol.github.io/img/kubectl-exec/api-server-to-kubelet.png)

- Now the connection between kubectl and api-server is still open and there is another connection between api-server and kubelet.





- 现在 kubectl 和 api-server 之间的连接仍然打开，并且 api-server 和 kubelet 之间还有另一个连接。





### 3. Activities in Worker Node

### 3. Worker 节点中的活动

- Let’s continue by connecting to the worker node and checking what is going on the worker node.

- 让我们继续连接到工作节点并检查工作节点上发生了什么。

First, we can observe the connection here as well. The second line. `192.168.205.10` is the IP of master node.

首先，我们也可以在这里观察连接。第二行。 `192.168.205.10` 是主节点的 IP。

```
   // worker node
   $ netstat -atn |grep 10250
   tcp6       0      0 :::10250                :::*                    LISTEN
   tcp6       0      0 192.168.205.11:10250    192.168.205.10:37870    ESTABLISHED
 ```

 
What about our sleep command? HOORAYYYY!! OUR COMMAND IS THERE!!!!

我们的睡眠命令呢？万岁！我们的命令在那里！！！！

```
   // worker node
   $ ps -afx
   ...
   31463 ? Sl     0:00      \_ docker-containerd-shim 7d974065bbb3107074ce31c51f5ef40aea8dcd535ae11a7b8f2dd180b8ed583a /var/run/docker/libcontainerd/7d974065bbb3107074ce31c51
   31478 pts/0    Ss     0:00          \_ sh
   31485 pts/0    S+     0:00              \_ sleep 5000
   ...
 ```

 
- Wait! How did kubelet do it?

- 等待！ kubelet 是怎么做的？

- kubelet has a daemon which serves an api over a port for api-server requests.

- kubelet 有一个守护进程，它通过端口为 api 服务器请求提供 api。

```
   // Server is the library interface to serve the stream requests.
   type Server interface {
           http.Handler
  
           // Get the serving URL for the requests.
           // Requests must not be nil. Responses may be nil iff an error is returned.
           GetExec(*runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error)
           GetAttach(req *runtimeapi.AttachRequest) (*runtimeapi.AttachResponse, error)
           GetPortForward(*runtimeapi.PortForwardRequest) (*runtimeapi.PortForwardResponse, error)
  
           // Start the server.
           // addr is the address to serve on (address:port) stayUp indicates whether the server should
           // listen until Stop() is called, or automatically stop after all expected connections are 
``
  
// closed. Calling Get{Exec,Attach,PortForward} increments the expected connection count.
           // Function does not return until the server is stopped.
           Start(stayUp bool) error
           // Stop the server, and terminate any open connections.
           Stop() error
   }
 ```

// 关闭。调用 Get{Exec,Attach,PortForward} 会增加预期的连接数。
          // 直到服务器停止，函数才会返回。
          开始（stayUp bool）错误
          // 停止服务器，并终止所有打开的连接。
          停止（）错误
  }
``

[view raw](https://github.com/kubernetes/kubernetes/raw/33081c1f07be128a89441a39c467b7ea2221b23d/pkg/kubelet/server/streaming/server.go)                    [pkg/kubelet/server/streaming/server.go](https:/ /github.com/kubernetes/kubernetes/blob/33081c1f07be128a89441a39c467b7ea2221b23d/pkg/kubelet/server/streaming/server.go)

  

[查看原始数据](https://github.com/kubernetes/kubernetes/raw/33081c1f07be128a89441a39c467b7ea2221b23d/pkg/kubelet/server/streaming/server.go) [pkg/kubelet/server/streaming/server.go]( /github.com/kubernetes/kubernetes/blob/33081c1f07be128a89441a39c467b7ea2221b23d/pkg/kubelet/server/streaming/server.go)

  

- kubelet computes a response endpoint for exec requests.

- kubelet 计算 exec 请求的响应端点。

```
   func (s *server) GetExec(req *runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error) {
           if err := validateExecRequest(req); err != nil {
                   return nil, err
           }
           token, err := s.cache.Insert(req)
           if err != nil {
                   return nil, err
           }
           return &runtimeapi.ExecResponse{
                   Url: s.buildURL("exec", token),
           }, nil
   }
 ```
 
[view raw](https://github.com/kubernetes/kubernetes/raw/33081c1f07be128a89441a39c467b7ea2221b23d/pkg/kubelet/server/streaming/server.go)
  

[查看原始数据](https://github.com/kubernetes/kubernetes/raw/33081c1f07be128a89441a39c467b7ea2221b23d/pkg/kubelet/server/streaming/server.go)
  

Don’t confuse. It doesn’t return the result of the command. It returns an endpoint for communication.

不要混淆。它不返回命令的结果。它返回一个用于通信的端点。

```
 type ExecResponse struct {
         // Fully qualified URL of the exec streaming server.
         Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
         XXX_NoUnkeyedLiteral struct{} `json:"-"`
         XXX_sizecache        int32    `json:"-"`
 }
 ```

 
[staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go](https://github.com/kubernetes/kubernetes/blob/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s. io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go)



[staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go](https://github.com/kubernetes/kubernetes/blob/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/staging/ io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go)



kubelet implements `RuntimeServiceClient` interface which is part of Container Runtime Interface.

kubelet 实现了 `RuntimeServiceClient` 接口，它是容器运行时接口的一部分。

```
 // For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
 type RuntimeServiceClient interface {
         // Version returns the runtime name, runtime version, and runtime API version.
         Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
         // RunPodSandbox creates and starts a pod-level sandbox. Runtimes must ensure
         // the sandbox is in the ready state on success.
         RunPodSandbox(ctx context.Context, in *RunPodSandboxRequest, opts ...grpc.CallOption) (*RunPodSandboxResponse, error)
         // StopPodSandbox stops any running process that is part of the sandbox and
         // reclaims network resources (e.g., IP addresses) allocated to the sandbox.
         // If there are any running containers in the sandbox, they must be forcibly
         // terminated.
         // This call is idempotent, and must not return an error if all relevant
         // resources have already been reclaimed. kubelet will call StopPodSandbox
         // at least once before calling RemovePodSandbox. It will also attempt to
         // reclaim resources eagerly, as soon as a sandbox is not needed. Hence,
         // multiple StopPodSandbox calls are expected.
         StopPodSandbox(ctx context.Context, in *StopPodSandboxRequest, opts ...grpc.CallOption) (*StopPodSandboxResponse, error)
         // RemovePodSandbox removes the sandbox. If there are any running containers
         // in the sandbox, they must be forcibly terminated and removed.
         // This call is idempotent, and must not return an error if the sandbox has
         // already been removed.
         RemovePodSandbox(ctx context.Context, in *RemovePodSandboxRequest, opts ...grpc.CallOption) (*RemovePodSandboxResponse, error)
         // PodSandboxStatus returns the status of the PodSandbox. If the PodSandbox is not
         // present, returns an error.
         PodSandboxStatus(ctx context.Context, in *PodSandboxStatusRequest, opts ...grpc.CallOption) (*PodSandboxStatusResponse, error)
         // ListPodSandbox returns a list of PodSandboxes.
         ListPodSandbox(ctx context.Context, in *ListPodSandboxRequest, opts ...grpc.CallOption) (*ListPodSandboxResponse, error)
         // CreateContainer creates a new container in specified PodSandbox
         CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error)
         // StartContainer starts the container.
         StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*StartContainerResponse, error)
         // StopContainer stops a running container with a grace period (i.e., timeout).
         // This call is idempotent, and must not return an error if the container has
         // already been stopped. 
``

// TODO: what must the runtime do after the grace period is reached?
         StopContainer(ctx context.Context, in *StopContainerRequest, opts ...grpc.CallOption) (*StopContainerResponse, error)
         // RemoveContainer removes the container. If the container is running, the
         // container must be forcibly removed.
         // This call is idempotent, and must not return an error if the container has
         // already been removed.
         RemoveContainer(ctx context.Context, in *RemoveContainerRequest, opts ...grpc.CallOption) (*RemoveContainerResponse, error)
         // ListContainers lists all containers by filters.
         ListContainers(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error)
         // ContainerStatus returns status of the container. If the container is not
         // present, returns an error.
         ContainerStatus(ctx context.Context, in *ContainerStatusRequest, opts ...grpc.CallOption) (*ContainerStatusResponse, error)
         // UpdateContainerResources updates ContainerConfig of the container.
         UpdateContainerResources(ctx context.Context, in *UpdateContainerResourcesRequest, opts ...grpc.CallOption) (*UpdateContainerResourcesResponse, error)
         // ReopenContainerLog asks runtime to reopen the stdout/stderr log file
         // for the container. This is often called after the log file has been
         // rotated. If the container is not running, container runtime can choose
         // to either create a new log file and return nil, or return an error.
         // Once it returns error, new container log file MUST NOT be created.
         ReopenContainerLog(ctx context.Context, in *ReopenContainerLogRequest, opts ...grpc.CallOption) (*ReopenContainerLogResponse, error)
         // ExecSync runs a command in a container synchronously.
         ExecSync(ctx context.Context, in *ExecSyncRequest, opts ...grpc.CallOption) (*ExecSyncResponse, error)
         // Exec prepares a streaming endpoint to execute a command in the container.
         Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
         // Attach prepares a streaming endpoint to attach to a running container.
         Attach(ctx context.Context, in *AttachRequest, opts ...grpc.CallOption) (*AttachResponse, error)
         // PortForward prepares a streaming endpoint to forward ports from a PodSandbox.
         PortForward(ctx context.Context, in *PortForwardRequest, opts ...grpc.CallOption) (*PortForwardResponse, error)
         // ContainerStats returns stats of the container. If the container does not
         // exist, the call returns an error.
         ContainerStats(ctx context.Context, in *ContainerStatsRequest, opts ...grpc.CallOption) (*ContainerStatsResponse, error)
         // ListContainerStats returns stats of all running containers.
         ListContainerStats(ctx context.Context, in *ListContainerStatsRequest, opts ...grpc.CallOption) (*ListContainerStatsResponse, error)
         // UpdateRuntimeConfig updates the runtime configuration based on the given request.
         UpdateRuntimeConfig(ctx context.Context, in *UpdateRuntimeConfigRequest, opts ...grpc.CallOption) (*UpdateRuntimeConfigResponse, error)
         // Status returns the status of the runtime.
         Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
 }
 ```
 [view raw](https://github.com/kubernetes/kubernetes/raw/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go)


 It just uses gRPC to invoke a method through Container Runtime Interface.

 ```
 
type runtimeServiceClient struct {
         cc *grpc.ClientConn
 }
 ```

类型 runtimeServiceClient 结构 {
        cc *grpc.ClientConn
}
``

[view raw](https://github.com/kubernetes/kubernetes/raw/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go)                    [staging/ src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go](https://github.com/kubernetes/kubernetes/blob/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri -api/pkg/apis/runtime/v1alpha2/api.pb.go)







[查看原始数据](https://github.com/kubernetes/kubernetes/raw/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go) [ src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go](https://github.com/kubernetes/kubernetes/blob/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/io/cri/k -api/pkg/apis/runtime/v1alpha2/api.pb.go)







```
 func (c *runtimeServiceClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
         out := new(ExecResponse)
         err := c.cc.Invoke(ctx, "/runtime.v1alpha2.RuntimeService/Exec", in, out, opts...)
         if err != nil {
                 return nil, err
         }
         return out, nil
 }
 ``` 


[view raw](https://github.com/kubernetes/kubernetes/raw/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go)                    [staging/ src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go](https://github.com/kubernetes/kubernetes/blob/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri -api/pkg/apis/runtime/v1alpha2/api.pb.go)



[查看原始数据](https://github.com/kubernetes/kubernetes/raw/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go) [ src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go](https://github.com/kubernetes/kubernetes/blob/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/io/cri/k -api/pkg/apis/runtime/v1alpha2/api.pb.go)



Container Runtime is responsible to implement `RuntimeServiceServer`

Container Runtime 负责实现`RuntimeServiceServer`

```
 // RuntimeServiceServer is the server API for RuntimeService service.
 type RuntimeServiceServer interface {
         // Version returns the runtime name, runtime version, and runtime API version.
         Version(context.Context, *VersionRequest) (*VersionResponse, error)
         // RunPodSandbox creates and starts a pod-level sandbox. Runtimes must ensure
         // the sandbox is in the ready state on success.
         RunPodSandbox(context.Context, *RunPodSandboxRequest) (*RunPodSandboxResponse, error)
         // StopPodSandbox stops any running process that is part of the sandbox and
         // reclaims network resources (e.g., IP addresses) allocated to the sandbox.
         // If there are any running containers in the sandbox, they must be forcibly
         // terminated.
         // This call is idempotent, and must not return an error if all relevant
         // resources have already been reclaimed. kubelet will call StopPodSandbox
         // at least once before calling RemovePodSandbox. It will also attempt to
         // reclaim resources eagerly, as soon as a sandbox is not needed. Hence,
         // multiple StopPodSandbox calls are expected.
         StopPodSandbox(context.Context, *StopPodSandboxRequest) (*StopPodSandboxResponse, error)
         // RemovePodSandbox removes the sandbox. If there are any running containers
         // in the sandbox, they must be forcibly terminated and removed.
         // This call is idempotent, and must not return an error if the sandbox has
         // already been removed.
         RemovePodSandbox(context.Context, *RemovePodSandboxRequest) (*RemovePodSandboxResponse, error)
         // PodSandboxStatus returns the status of the PodSandbox. If the PodSandbox is not
         // present, returns an error.
         PodSandboxStatus(context.Context, *PodSandboxStatusRequest) (*PodSandboxStatusResponse, error)
         // ListPodSandbox returns a list of PodSandboxes.
         ListPodSandbox(context.Context, *ListPodSandboxRequest) (*ListPodSandboxResponse, error)
         // CreateContainer creates a new container in specified PodSandbox
         CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error)
         // StartContainer starts the container.
         StartContainer(context.Context, *StartContainerRequest) (*StartContainerResponse, error)
         // StopContainer stops a running container with a grace period (i.e., timeout).
         // This call is idempotent, and must not return an error if the container has
         // already been stopped.
         // TODO: what must the runtime do after the grace period is reached?
         StopContainer(context.Context, *StopContainerRequest) (*StopContainerResponse, error)
         // RemoveContainer removes the container. If the container is running, the
         // container must be forcibly removed.
         // This call is idempotent, and must not return an error if the container has
         // already been removed.
         RemoveContainer(context.Context, *RemoveContainerRequest) (*RemoveContainerResponse, error)
         // ListContainers lists all containers by filters.
         ListContainers(context.Context, *ListContainersRequest) (*ListContainersResponse, error)
         // ContainerStatus returns status of the container. If the container is not
         // present, returns an error.
         ContainerStatus(context.Context, *ContainerStatusRequest) (*ContainerStatusResponse, error)
         // UpdateContainerResources updates ContainerConfig of the container.
         UpdateContainerResources(context.Context, *UpdateContainerResourcesRequest) (*UpdateContainerResourcesResponse, error)
         // ReopenContainerLog asks runtime to reopen the stdout/stderr log file
         // for the container. This is often called after the log file has been
         // rotated. If the container is not running, container runtime can choose
         // to either create a new log file and return nil, or return an error.
         // Once it returns error, new container log file MUST NOT be created.
         ReopenContainerLog(context.Context, *ReopenContainerLogRequest) (*ReopenContainerLogResponse, error)
         // ExecSync runs a command in a container synchronously.
         ExecSync(context.Context, *ExecSyncRequest) (*ExecSyncResponse, error) 
``

// Exec prepares a streaming endpoint to execute a command in the container.
         Exec(context.Context, *ExecRequest) (*ExecResponse, error)
         // Attach prepares a streaming endpoint to attach to a running container.
         Attach(context.Context, *AttachRequest) (*AttachResponse, error)
         // PortForward prepares a streaming endpoint to forward ports from a PodSandbox.
         PortForward(context.Context, *PortForwardRequest) (*PortForwardResponse, error)
         // ContainerStats returns stats of the container. If the container does not
         // exist, the call returns an error.
         ContainerStats(context.Context, *ContainerStatsRequest) (*ContainerStatsResponse, error)
         // ListContainerStats returns stats of all running containers.
         ListContainerStats(context.Context, *ListContainerStatsRequest) (*ListContainerStatsResponse, error)
         // UpdateRuntimeConfig updates the runtime configuration based on the given request.
         UpdateRuntimeConfig(context.Context, *UpdateRuntimeConfigRequest) (*UpdateRuntimeConfigResponse, error)
         // Status returns the status of the runtime.
         Status(context.Context, *StatusRequest) (*StatusResponse, error)
 }
 ```

// Exec 准备一个流端点来执行容器中的命令。
        Exec(context.Context, *ExecRequest) (*ExecResponse, 错误)
        // 附加准备一个流端点以附加到正在运行的容器。
        附加（上下文。上下文，*AttachRequest）（*AttachResponse，错误）
        // PortForward 准备一个流端点来转发来自 PodSandbox 的端口。
        PortForward(context.Context, *PortForwardRequest) (*PortForwardResponse, 错误)
        // ContainerStats 返回容器的统计信息。如果容器没有
        // 存在，调用返回错误。
        ContainerStats(context.Context, *ContainerStatsRequest) (*ContainerStatsResponse, error)
        // ListContainerStats 返回所有正在运行的容器的统计信息。
        ListContainerStats(context.Context, *ListContainerStatsRequest) (*ListContainerStatsResponse, error)
        // UpdateRuntimeConfig 根据给定的请求更新运行时配置。
        UpdateRuntimeConfig(context.Context, *UpdateRuntimeConfigRequest) (*UpdateRuntimeConfigResponse, error)
        // Status 返回运行时的状态。
        状态（上下文。上下文，*状态请求）（*状态响应，错误）
}
``

[view raw](https://github.com/kubernetes/kubernetes/raw/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go)                    [staging/ src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go](https://github.com/kubernetes/kubernetes/blob/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri -api/pkg/apis/runtime/v1alpha2/api.pb.go)





[查看原始数据](https://github.com/kubernetes/kubernetes/raw/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go) [ src/k8s.io/cri-api/pkg/apis/runtime/v1alpha2/api.pb.go](https://github.com/kubernetes/kubernetes/blob/6568325ca2bef519e5c8228cd33887660b5ed7b0/staging/io/cri/k -api/pkg/apis/runtime/v1alpha2/api.pb.go)





![kubelet-to-container-runtime](https://erkanerol.github.io/img/kubectl-exec/kubelet-to-container-runtime.png)

![kubelet-to-container-runtime](https://erkanerol.github.io/img/kubectl-exec/kubelet-to-container-runtime.png)

- If it is so, we need to observe a connection between kubelet and container runtime. Right? Let’s check.

- 如果是这样，我们需要观察 kubelet 和容器运行时之间的连接。对？让我们检查。

Run this command before and after running exec command and check the diff. This one is the diff in my case.

在运行 exec 命令之前和之后运行此命令并检查差异。在我的情况下，这是一个差异。

```
 // worker node
 $ ss -a -p |grep kubelet
 ...
 u_str  ESTAB      0      0       * 157937                * 157387                users:(("kubelet",pid=5714,fd=33))
 ...
 ```

 
Hımmm. There is a new connection via unix sockets  between kubelet(pid=5714) and something. Who can be? YES. IT IS  DOCKER(pid=1186).

嗯嗯。 kubelet(pid=5714) 和某个东西之间通过 unix 套接字建立了一个新连接。谁可以是？是的。它是码头工人（pid=1186）。

```
 // worker node
 $ ss -a -p |grep 157387
 ...
 u_str  ESTAB      0      0       * 157937                * 157387                users:(("kubelet",pid=5714,fd=33))
 u_str  ESTAB      0      0      /var/run/docker.sock 157387                * 157937                users:(("dockerd",pid=1186,fd=14))
 ...
 ```

 
Remember. This is the docker daemon process(pid=1186) which runs our command.

记住。这是运行我们的命令的 docker 守护进程（pid=1186）。

```
 // worker node.
 $ ps -afx
 ...
  1186 ? Ssl    0:55 /usr/bin/dockerd -H fd://
 17784 ? Sl     0:00      \_ docker-containerd-shim 53a0a08547b2f95986402d7f3b3e78702516244df049ba6c5aa012e81264aa3c /var/run/docker/libcontainerd/53a0a08547b2f95986402d7f3
 17801 pts/2    Ss     0:00          \_ sh
 17827 pts/2    S+     0:00              \_ sleep 5000
 ...
 ```

 
### 4. Activities in Container Runtime

### 4. 容器运行时的活动

- Let’s check cri-o’s source code to understand how it can happen. The logic is similar in docker.

- 让我们检查 cri-o 的源代码以了解它是如何发生的。 docker 中的逻辑类似。

It has a server which implements RuntimeServiceServer.

它有一个实现 RuntimeServiceServer 的服务器。

```
 // Server implements the RuntimeService and ImageService
 type Server struct {
         config          libconfig.Config
         seccompProfile  *seccomp.Seccomp
         stream          StreamService
         netPlugin       ocicni.CNIPlugin
         hostportManager hostport.HostPortManager

         appArmorProfile string
         hostIP          string
         bindAddress     string

         *lib.ContainerServer
         monitorsChan      chan struct{}
         defaultIDMappings *idtools.IDMappings
         systemContext     *types.SystemContext // Never nil

         updateLock sync.RWMutex

         seccompEnabled  bool
         appArmorEnabled bool
 ```

 
[view raw](https://github.com/cri-o/cri-o/raw/b3c67952115bee6f44405f547fef0b65a9134c60/server/server.go)                    [server/server.go](https://github.com/cri-o /cri-o/blob/b3c67952115bee6f44405f547fef0b65a9134c60/server/server.go)







[查看原始数据](https://github.com/cri-o/cri-o/raw/b3c67952115bee6f44405f547fef0b65a9134c60/server/server.go) [server/server.go](https://github.com/cri-o /cri-o/blob/b3c67952115bee6f44405f547fef0b65a9134c60/server/server.go)







```
 // Exec prepares a streaming endpoint to execute a command in the container.
 func (s *Server) Exec(ctx context.Context, req *pb.ExecRequest) (resp *pb.ExecResponse, err error) {
         const operation = "exec"
         defer func() {
                 recordOperation(operation, time.Now())
                 recordError(operation, err)
         }()

         resp, err = s.getExec(req)
         if err != nil {
                 return nil, fmt.Errorf("unable to prepare exec endpoint: %v", err)
         }

         return resp, nil
 }
 ``` 


[view raw](https://github.com/cri-o/cri-o/raw/b3c67952115bee6f44405f547fef0b65a9134c60/server/container_exec.go)                    [server/container_exec.go](https://github.com/cri-o /cri-o/blob/b3c67952115bee6f44405f547fef0b65a9134c60/server/container_exec.go)





[查看原始数据](https://github.com/cri-o/cri-o/raw/b3c67952115bee6f44405f547fef0b65a9134c60/server/container_exec.go) [server/container_exec.go](https://github.com/cri-o /cri-o/blob/b3c67952115bee6f44405f547fef0b65a9134c60/server/container_exec.go)





At the end of the chain, container runtime executes the command in the worker node.

在链的末端，容器运行时在工作节点中执行命令。

```
 // ExecContainer prepares a streaming endpoint to execute a command in the container.
 func (r *runtimeOCI) ExecContainer(c *Container, cmd []string, stdin io.Reader, stdout, stderr io.WriteCloser, tty bool, resize <-chan remotecommand.TerminalSize) error {
         processFile, err := prepareProcessExec(c, cmd, tty)
         if err != nil {
                 return err
         }
         defer os.RemoveAll(processFile.Name())

         args := []string{rootFlag, r.root, "exec"}
         args = append(args, "--process", processFile.Name(), c.ID())
         execCmd := exec.Command(r.path, args...)
         if v, found := os.LookupEnv("XDG_RUNTIME_DIR"); found {
                 execCmd.Env = append(execCmd.Env, fmt.Sprintf("XDG_RUNTIME_DIR=%s", v))
         }
         var cmdErr, copyError error
         if tty {
                 cmdErr = ttyCmd(execCmd, stdin, stdout, resize)
         } else {
                 if stdin != nil {
                         // Use an os.Pipe here as it returns true *os.File objects.
                         // This way, if you run 'kubectl exec <pod> -i bash' (no tty) and type 'exit',
                         // the call below to execCmd.Run() can unblock because its Stdin is the read half
                         // of the pipe.
                         r, w, err := os.Pipe()
                         if err != nil {
                                 return err
                         }
                         go func() { _, copyError = pools.Copy(w, stdin) }()

                         execCmd.Stdin = r
                 }
                 if stdout != nil {
                         execCmd.Stdout = stdout
                 }
                 if stderr != nil {
                         execCmd.Stderr = stderr
                 }

                 cmdErr = execCmd.Run()
         }

         if copyError != nil {
                 return copyError
         }
         if exitErr, ok := cmdErr.(*exec.ExitError); ok {
                 return &utilexec.ExitErrorWrapper{ExitError: exitErr}
         }
         return cmdErr
 }
 ```
 
[view raw](https://github.com/cri-o/cri-o/raw/b3c67952115bee6f44405f547fef0b65a9134c60/internal/oci/runtime_oci.go)
 [internal/oci/runtime_oci.go](https://github.com/cri-o/cri-o/blob/b3c67952115bee6f44405f547fef0b65a9134c60/internal/oci/runtime_oci.go)

[查看原始数据](https://github.com/cri-o/cri-o/raw/b3c67952115bee6f44405f547fef0b65a9134c60/internal/oci/runtime_oci.go)
[internal/oci/runtime_oci.go](https://github.com/cri-o/cri-o/blob/b3c67952115bee6f44405f547fef0b65a9134c60/internal/oci/runtime_oci.go)

![container-runtime-to-kernel](https://erkanerol.github.io/img/kubectl-exec/container-runtime-to-kernel.png)

![container-runtime-to-kernel](https://erkanerol.github.io/img/kubectl-exec/container-runtime-to-kernel.png)

Finally, kernel executes commands. ![kernel-puts](https://erkanerol.github.io/img/kubectl-exec/kernel-puts.png)

最后，内核执行命令。 ![kernel-puts](https://erkanerol.github.io/img/kubectl-exec/kernel-puts.png)

## Reminders

## 提醒

- api-server can also initialize a connection to kubelet.
 - These connections persist until the interactive exec ends.
   - Connection between kubectl and api-server
   - Connection between api-server and kubelet
   - Connection between kubelet and container runtime
 - kubectl or api-server cannot run anything in the worker nodes. kubelet can run  but it also interacts with container runtime for this kind of actions.

- api-server 也可以初始化与 kubelet 的连接。
- 这些连接持续到交互式 exec 结束。
  - kubectl 和 api-server 之间的连接
  - api-server 和 kubelet 之间的连接
  - kubelet 和容器运行时之间的连接
- kubectl 或 api-server 不能在工作节点中运行任何东西。 kubelet 可以运行，但它也会与容器运行时进行交互以进行此类操作。

## Resources

##  资源

- https://groups.google.com/forum/#!topic/kubernetes-dev/Cjia36v39vM
 - https://medium.com/@joatmon08/playing-with-kubeadm-in-vagrant-machines-part-2-bac431095706
 - https://serverfault.com/questions/252723/how-to-find-other-end-of-unix-socket-connection 
- https://groups.google.com/forum/#!topic/kubernetes-dev/Cjia36v39vM
- https://medium.com/@joatmon08/playing-with-kubeadm-in-vagrant-machines-part-2-bac431095706
- https://serverfault.com/questions/252723/how-to-find-other-end-of-unix-socket-connection
