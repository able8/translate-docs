# translate-docs

使用 [Google 翻译 API](https://translate.google.cn/)，将输入文档按段落生成 英文 - 中文 对照翻译。

阅读英文文档，有时理解不准确，使用中英文对照理解更准确。在改进翻译中，加深理解。

### 如何使用

```
./translate_darwin_amd64 -f example-input.md

```

Output:
```
2021/08/14 23:31:06 Input file is "example-input.md"
2021/08/14 23:31:06 Translating ...
2021/08/14 23:31:06 Done. Generated output file:  tr-example-input.md
```

### go build 生成 Mac、linux、Windows 平台可执行文件

```sh
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o translate_darwin_amd64 translate.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o translate_linux_amd64 translate.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o translate_windows_amd64.exe translate.go
```

---

# Translate docs & blogs

## Kubernetes

1. [Handling Client Requests Properly with Kubernetes 使用 Kubernetes 正确处理客户端请求](docs/tr-1-Handling-Client-Requests-Properly-with-Kubernetes.md)
1. [Zero Downtime Server Updates For Your Kubernetes Cluster Kubernetes 集群的零停机服务器更新](docs/tr-2-zero-downtime-server-updates-for-your-kubernetes-cluster.md)
1. [Gracefully Shutting Down Pods in a Kubernetes Cluster 优雅地关闭 Kubernetes 集群中的 Pod](docs/tr-3-gracefully-shutting-down-pods-in-a-kubernetes-cluster.md)
1. [Delaying Shutdown to Wait for Pod Deletion Propagation 延迟关闭以等待 Pod 删除传播](docs/tr-4-delaying-shutdown-to-wait-for-pod-deletion-propagation.md)
1. [Avoiding Outages in your Kubernetes Cluster using PodDisruptionBudgets 使用 PodDisruptionBudgets 避免 Kubernetes 集群中断](docs/tr-5-avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets.md)
1. [How does 'kubectl exec' work? 'kubectl exec' 是如何工作的？](docs/tr-6-how-kubectl-exec-works.md)
1. [](docs/)
1. [](docs/)


1. [Kubernetes vs Docker: Understanding Containers in 2021 了解 Kubernetes vs Docker](docs/tr-8-kubernetes-vs-docker.md)
1. [Learning Path: Basics of Container Runtimes 学习路径：容器运行时的基础知识](docs/tr-7-basics-of-container-runtimes.md)
1. [Container Runtimes Part 1: An Introduction to Container Runtimes 容器运行时第 1 部分：容器运行时简介](docs/tr-9-container-runtimes-part-1-introduction-container-r.md)
1. [Container Runtimes Part 2: Anatomy of a Low-Level Container Runtime 容器运行时第 2 部分：底层容器运行时剖析](docs/tr-10-container-runtimes-part-2-anatomy-low-level.md)
1. [Container Runtimes Part 3: High-Level Runtimes 容器运行时第 3 部分：高级运行时](docs/tr-11-container-runtimes-part-3-high-level-runtimes.md)
1. [Container Runtimes Part 4: Kubernetes Container Runtimes & CRI 容器运行时第 4 部分：Kubernetes 容器运行时和 CRI](docs/tr-12-container-runtimes-part-4-kubernetes-container-run.md)
1. [](docs/)
1. [](docs/)
1. [](docs/)
1. [](docs/)

## Golang

## Cloud Native



