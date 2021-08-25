# Graceful shutdown with Go http servers and Kubernetes rolling updates

# 使用 Go http 服务器和 Kubernetes 滚动更新优雅关闭

[Aug 6, 2018](https://medium.com/over-engineering/graceful-shutdown-with-go-http-servers-and-kubernetes-rolling-updates-6697e7db17cf) · 5 min read

If you’re writing Go then you’re probably aware that graceful shutdown was added to the http package in 1.8.

如果您正在编写 Go，那么您可能知道优雅关闭已添加到 1.8 中的 http 包中。

> The HTTP server also adds support for [graceful shutdown](https://golang.org/doc/go1.8#http_shutdown), allowing servers to minimize downtime by shutting down only after serving all requests that are in flight. — [Go 1.8 is released](https://blog.golang.org/go1.8)

> HTTP 服务器还增加了对 [graceful shutdown](https://golang.org/doc/go1.8#http_shutdown) 的支持，允许服务器仅在处理所有正在运行的请求后才关闭，从而最大限度地减少停机时间。 — [Go 1.8 发布](https://blog.golang.org/go1.8)

Similarly, If you’re using Kubernetes then I’m sure you’re aware of, and hopefully using rolling updates for your deployments.

同样，如果您正在使用 Kubernetes，那么我相信您知道并希望在您的部署中使用滚动更新。

> Rolling updates incrementally replace your resource’s Pods with new ones, which are then scheduled on nodes with available resources. Rolling updates  are designed to update your workloads without downtime. — [Performing Rolling Updates](https://cloud.google.com/kubernetes-engine/docs/how-to/updating-apps)

> 滚动更新以增量方式用新的 Pod 替换您资源的 Pod，然后将这些 Pod 安排在具有可用资源的节点上。滚动更新旨在在不停机的情况下更新您的工作负载。 — [执行滚动更新](https://cloud.google.com/kubernetes-engine/docs/how-to/updating-apps)

However, you might not be sure how the two work together to ensure truly zero  downtime deployments — I wasn’t! This is a quick guide to writing  readiness and liveness probes in Go and how to configure them with a  Kubernetes rolling update deployment.

但是，您可能不确定这两者如何协同工作以确保真正的零停机部署——我不是！这是在 Go 中编写就绪和活跃度探测器以及如何使用 Kubernetes 滚动更新部署配置它们的快速指南。

## Getting started

## 入门

Let’s get started with a basic health check, to begin we’ll use it for both  the readiness and liveness probes in Kubernetes. This was how I started  writing services, and until you hit a significant amount of traffic,  it’s probably fine.

让我们从一个基本的健康检查开始，首先我们将把它用于 Kubernetes 中的就绪和活跃度探测。这就是我开始编写服务的方式，直到您遇到大量流量，这可能没问题。

Here’s a sample Go app, using a [Chi router and middleware](https://github.com/go-chi/chi) as well as the deployment probe configuration.

这是一个示例 Go 应用程序，使用 [Chi 路由器和中间件](https://github.com/go-chi/chi) 以及部署探针配置。

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/health"))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe(":3000", r)
}
```

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 80
readinessProbe:
  httpGet:
    path: /health
    port: 80
```

This will work fine, but when performing a rolling update Kubernetes will send a `SIGTERM` signal to the process and the server will die. Any open connections will fail resulting in a bad experience for users

这会正常工作，但是当执行滚动更新时，Kubernetes 会向进程发送一个 `SIGTERM` 信号并且服务器将死掉。任何打开的连接都将失败，从而导致用户体验不佳

## Graceful shutdown

## 优雅关机

There is fantastic documentation and example code for the `server.Shutdown` method on [godoc.org](https://golang.org/pkg/net/http/#Server.Shutdown). One thing to note here is that the example uses `os.Interrupt` as the shutdown signal. That’ll work when you run a server locally and hit `ctrl-C` to close it, but not on Kubernetes. Kubernetes sends a `SIGTERM` signal which is different.

[godoc.org](https://golang.org/pkg/net/http/#Server.Shutdown) 上有关于`server.Shutdown` 方法的精彩文档和示例代码。这里要注意的一件事是该示例使用 `os.Interrupt` 作为关闭信号。当您在本地运行服务器并按 ctrl-C 关闭它时，这会起作用，但在 Kubernetes 上不起作用。 Kubernetes 发送一个不同的 `SIGTERM` 信号。


```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var srv http.Server

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
```

Now, let’s look at how we integrate graceful shutdown with the two different Kubernetes probes.

现在，让我们看看我们如何将优雅关闭与两个不同的 Kubernetes 探针集成在一起。

A quick primer on probes in Kubernetes:
**Liveness** indicates that the pod is running, if the liveness probe fails the pod will be restarted.
**Readiness** indicates that the pod is ready to receive traffic, when a pod is ready the load balancer will start sending traffic to it.

Kubernetes 中探针的快速入门：
**Liveness** 表示 pod 正在运行，如果 liveness 探测失败，pod 将重新启动。
**Readiness** 表示 Pod 已准备好接收流量，当 Pod 准备就绪时，负载均衡器将开始向其发送流量。

Let’s go through the steps we want to happen during a rolling update.
1. Running pods are told to shut down via a `SIGTERM` signal
2. The readiness probes on those running pods should start failing as they are not accepting new traffic, but only closing any open connections
3. During this time new pods have started up and will be ready to receive traffic
4. Once all connections to old pods have been closed, the liveness probes should fail
5. The old pods are now completely shut down and assuming the rollout went well, all traffic should be going to new pods

让我们来看看我们希望在滚动更新期间发生的步骤。
1. 正在运行的 pod 被告知通过 `SIGTERM` 信号关闭
2. 那些正在运行的 Pod 上的就绪探测应该开始失败，因为它们不接受新流量，而只是关闭任何打开的连接
3. 在此期间，新的 Pod 已启动并准备好接收流量
4. 一旦与旧 Pod 的所有连接都关闭，活性探测就会失败
5. 旧的 Pod 现在完全关闭，假设部署顺利，所有流量都应该流向新的 Pod

We’ll need to use two different types of probes to achieve this, an `httpGet` probe for readiness and a `command` for liveness. Think about it, once a Go application receives the `SIGTERM` signal it will no longer serve new traffic so any http health checks  will fail, this is why the liveness check needs to be independent of the http server.

我们需要使用两种不同类型的探测器来实现这一点，一个用于准备就绪的 `httpGet` 探测器和用于活性的 `command`。想一想，一旦 Go 应用程序收到 `SIGTERM` 信号，它将不再提供新的流量，因此任何 http 健康检查都会失败，这就是为什么活性检查需要独立于 http 服务器的原因。

The most common way to do this is by creating a file on disk when the  process starts, and remove it when ending. The liveness check is then a  simple `cat` command to check that the file exists.

最常见的方法是在进程开始时在磁盘上创建一个文件，并在结束时将其删除。活性检查是一个简单的 `cat` 命令来检查文件是否存在。

Here’s an example of how we would create the probe file before running the application, and remove it once it has completed. 
下面是一个示例，说明我们如何在运行应用程序之前创建探测文件，并在完成后将其删除。

```go
package probe

import "os"

const liveFile = "/tmp/live"

// Create will create a file for the liveness check
func Create() error {
	_, err := os.Create(liveFile)
	return err
}

// Remove will remove the file create for the liveness probe
func Remove() error {
	return os.Remove(liveFile)
}

// Exists checks if the file created for the liveness probe exists
func Exists() bool {
	if _, err := os.Stat(liveFile); err == nil {
		return true
	}

	return false
}
```

```go
// Command contains the app command
var Command = &cobra.Command{
	Use:   "app",
	Short: "Run the application",
	Run: func(cmd *cobra.Command, args []string) {
		var s app.Specification
		envconfig.MustProcess("", &s)

		if err := probe.Create(); err != nil {
			panic(err)
		}

		app.Run(s)

		if err := probe.Remove(); err != nil {
			panic(err)
		}
	},
}
```

I’ve hardcoded the location of the probe file here to simplify the example  code, but don’t overlook this. You’ll need to make sure that you mount a [persistent volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) to write to. I’d also suggest using the `ioutil.TempFile` method available in the Go standard library.

我在这里硬编码了探针文件的位置以简化示例代码，但不要忽视这一点。你需要确保你挂载了一个 [persistent volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) 来写入。我还建议使用 Go 标准库中提供的 `ioutil.TempFile` 方法。

Now… we need to update the deployment configuration to check for the existence of `/tmp/live`.

现在……我们需要更新部署配置以检查 `/tmp/live` 是否存在。

```yaml
livenessProbe:
  exec:
    command:
      - cat
      - /tmp/live
readinessProbe:
  httpGet:
    path: /health
    port: 80
```

That *should* be all you need, but I did run into an interesting problem.

这*应该*是你所需要的，但我确实遇到了一个有趣的问题。

## Debugging

## 调试

You might run into issues, like I did, where the pod got stuck in a restart loop. Once you’ve found the pod name, you can use the describe command  to see what’s going on:

您可能会像我一样遇到问题，即 pod 卡在重启循环中。找到 pod 名称后，您可以使用 describe 命令查看发生了什么：

```
kubectl describe pod my-application-789757d855-ptccl
```


In my case, I ran into a rather interesting error:

就我而言，我遇到了一个相当有趣的错误：

```
Warning  Unhealthy              16s (x7 over 1m)  kubelet, gke-dev-cluster-2-dev-pool-2-b6567556-l83v  Liveness probe failed: rpc error: code = 2 desc = oci runtime error: exec failed: container_linux.go: 247: starting container process caused "exec: \"cat\": executable file not found in $PATH"
```

**That’s right, the container I was using didn’t have** `cat` **installed.**

**是的，我用的容器没有**`cat` **安装。**

I’ve been using the [distroless container from Google](https://github.com/GoogleContainerTools/distroless) as a base image for my application containers. It has everything I need to run a Go binary and nothing else. This is fantastic because the  containers I ship are tiny… build times are faster, deployments are  faster and it reduces security risk.

我一直在使用 [来自 Google 的 distroless 容器](https://github.com/GoogleContainerTools/distroless) 作为我的应用程序容器的基础镜像。它拥有运行 Go 二进制文件所需的一切，仅此而已。这太棒了，因为我运送的容器很小……构建时间更快，部署更快，并且降低了安全风险。

It’s great that I can ship 15mb containers, but… not so great when I can’t use basic utilities like `cat`. Rather than poking around in the container to see what else I could  use, or use a different base image, I wrote a small sub-command in my Go application to handle this. You’ll notice a `probe.Exists` method in the `probe` package above, here’s what the liveness sub-command looks like:

我可以运送 15 mb 的容器很棒，但是……当我不能使用像“cat”这样的基本实用程序时就不是那么好。我没有在容器中四处寻找我可以使用的其他内容或使用不同的基础映像，而是在 Go 应用程序中编写了一个小的子命令来处理这个问题。你会注意到上面 `probe` 包中有一个 `probe.Exists` 方法，下面是 liveness 子命令的样子：

```go
package live

import (
	"os"

	"github.com/overhq/over-stories-api/pkg/probe"
	"github.com/spf13/cobra"
)

// Command contains the livee command
var Command = &cobra.Command{
	Use:   "live",
	Short: "Check if application is live",
	Run: func(cmd *cobra.Command, args []string) {
		if probe.Exists() {
			os.Exit(0)
		}

		os.Exit(1)
	},
}
```

Calling the liveness check as a sub-command instead of using `cat` directly sorted out my problems, hopefully that helps if you run into a similar issue.

将活性检查作为子命令调用，而不是直接使用 `cat` 解决了我的问题，希望在您遇到类似问题时有所帮助。

Finally, here are the logs to show a complete rolling update deployment with  zero downtime and no lost connections during the process:

最后，这里的日志显示了一个完整的滚动更新部署，在此过程中零停机且没有丢失连接：

![img](https://miro.medium.com/max/2000/1*as80VVAHpfCBd4YNzlqyvA.png)

I’ve redacted some logs here, but you can see the startup and shutdown process happening!

我在这里编辑了一些日志，但是您可以看到启动和关闭过程的发生！

I hope this post has provided some insight into how the graceful shutdown and rolling updates can work together to achieve truly zero downtime  deployments. If you have any questions, thoughts or suggestions then I’d love to hear from you!

我希望这篇文章提供了一些关于优雅关闭和滚动更新如何协同工作以实现真正的零停机部署的见解。如果您有任何问题、想法或建议，我很乐意听取您的意见！

Find me on Twitter at [@waynethebrain](https://twitter.com/waynethebrain)

**Further reading**

- [Kubernetes Container probes documentation](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes)
- [Go ](https://golang.org/pkg/net/http/#Server.Shutdown)`server.Shutdown`[ documentation](https://golang.org/pkg/net/http/#Server. Shutdown)
- [Performing Rolling Updates](https://cloud.google.com/kubernetes-engine/docs/how-to/updating-apps) 

**进一步阅读**

- [Kubernetes 容器探测文档](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes)
- [Go ](https://golang.org/pkg/net/http/#Server.Shutdown)`server.Shutdown`[文档](https://golang.org/pkg/net/http/#Server.Shutdown)
- [执行滚动更新](https://cloud.google.com/kubernetes-engine/docs/how-to/updating-apps)