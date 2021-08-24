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

1. [Handling Client Requests Properly with Kubernetes](docs/tr-1-Handling-Client-Requests-Properly-with-Kubernetes.md)
1. [Zero Downtime Server Updates For Your Kubernetes Cluster Kubernetes](docs/tr-2-zero-downtime-server-updates-for-your-kubernetes-cluster.md)
1. [Gracefully Shutting Down Pods in a Kubernetes Cluster](docs/tr-3-gracefully-shutting-down-pods-in-a-kubernetes-cluster.md)
1. [Delaying Shutdown to Wait for Pod Deletion Propagation](docs/tr-4-delaying-shutdown-to-wait-for-pod-deletion-propagation.md)
1. [Avoiding Outages in your Kubernetes Cluster using PodDisruptionBudgets](docs/tr-5-avoiding-outages-in-your-kubernetes-cluster-using-poddisruptionbudgets.md)
1. [How does 'kubectl exec' work?](docs/tr-6-how-kubectl-exec-works.md)
1. [What is an Operator](docs/tr-16-what-is-an-operator.md)
1. [Introducing Operators: Putting Operational Knowledge into Software](docs/tr-17-introducing-operators.md)
1. [Go Operator Tutorial](docs/tr-18-Go-Operator-Tutorial.md)
1. [Ready-to-use commands and tips for kubectl](docs/tr-19-ready-to-use-commands-and-tips-for-kubectl.md)
1. [How we enjoyed upgrading a bunch of Kubernetes clusters from v1.16 to v1.19](docs/tr-20-how-we-enjoyed-upgrading-kubernetes-clusters.md)
1. [Using Telepresence 2 for Kubernetes debugging and local development](docs/tr-21-telepresence-2-local-development.md)
1. [Everything you Need to Know about Kubernetes Quality of Service (QoS) Classes](docs/tr-29-everything-you-need-to-know-about-kubernetes-quality-of-service-qos-classes.md)
1. [Kubernetes Autoscaling in Production: Best Practices for Cluster Autoscaler, HPA and VP](docs/tr-33-kubernetes-in-production-best-practices-for-cluster-autoscaler-hpa-and-vpa.md)
1. [A CIOs Guide to Kubernetes Best Practices in Production](docs/tr-34-a-cios-guide-to-kubernetes-best-practices-in-production.md)
1. [10 Kubernetes Operators Every DevOps needs to know about](docs/tr-31-10-kubernetes-operators-every-devops-needs-to-know-about.md)
1. [](docs/)


## Docker and Container

1. [Kubernetes vs Docker: Understanding Containers in 2021](docs/tr-8-kubernetes-vs-docker.md)
1. [Learning Path: Basics of Container Runtimes](docs/tr-7-basics-of-container-runtimes.md)
1. [Container Runtimes Part 1: An Introduction to Container Runtimes](docs/tr-9-container-runtimes-part-1-introduction-container-r.md)
1. [Container Runtimes Part 2: Anatomy of a Low-Level Container Runtime](docs/tr-10-container-runtimes-part-2-anatomy-low-level.md)
1. [Container Runtimes Part 3: High-Level Runtimes](docs/tr-11-container-runtimes-part-3-high-level-runtimes.md)
1. [Container Runtimes Part 4: Kubernetes Container Runtimes & CRI](docs/tr-12-container-runtimes-part-4-kubernetes-container-run.md)
1. [Docker components explained](docs/tr-13-docker-components-explained.md)
1. [Docker vs CRI-O vs Containerd](docs/tr-14-docker-vs-cri-o-vs-containerd.md)
1. [Introducing Container Runtime Interface (CRI) in Kubernetes](docs/tr-15-container-runtime-interface-cri-in-kubernetes.md)
1. [Understanding Docker Container Exit Codes](docs/tr-28-understanding-docker-container-exit-codes.md)

## Golang

1. [Keep-Alive in http requests in golang](docs/tr-22-keep-alive-http-requests-in-golang.md)
1. [Don’t use Go’s default HTTP client (in production)](docs/tr-23-don-t-use-go-s-default-http-client.md)
1. [Using Context in Golang - Cancellation, Timeouts and Values (With Examples)](docs/tr-26-context-cancellation-and-values.md)
1. [map[string]interface{} in Go](docs/tr-27-map-string-interface.md)
1. [](docs/)
1. [](docs/)

## Cloud Native

1. [A Comprehensive Guide to Cloud Native: Tools, Practices and Culture](docs/tr-32-a-comprehensive-guide-to-cloud-native-tools-practices-and-culture.md)
1. [The State of Cloud Native: Challenges, Culture and Technology](docs/tr-30-the-state-of-cloud-native-challenges-culture-and-technology.md)
1. [](docs/)

## Others

1. [Fix the Problem, Not the Symptoms](docs/tr-24-fix-problem-not-symptoms.md)
1. [Why do (some) engineers write bad code?](docs/tr-25-why-engineers-write-bad-code.md)
1. [](docs/)