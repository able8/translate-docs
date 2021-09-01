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
1. [Graceful shutdown and zero downtime deployments in Kubernetes](docs/tr-68-graceful-shutdown.md)
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
1. [Bad Pods: Kubernetes Pod Privilege Escalation](docs/tr-67-bad-pods-kubernetes-pod-privilege-escalation.md)
1. [](docs/tr-69-introduction-to-ebpf.md)
1. [](docs/tr-70-kubernetes-autoscaling-strategies.md)
1. [](docs/tr-71-kubernetes-long-lived-connections.md)
1. [](docs/tr-72-kubernetes-operators-best-practices.md)
1. [7 Principles of DevSecOps With Kubernetes](docs/tr-79-seven-principles-of-devsecops-with-kubernetes.md)
1. [Kubernetes Is Not Your Platform, It's Just the Foundation](docs/tr-77-kubernetes-successful-adoption-foundation.md)
1. [Moving k8s communication to gRPC](docs/tr-80-moving-k8s-communication-to-grpc.md)


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
1. [Graceful Shutdowns in Golang with signal.NotifyContext](docs/36-graceful-shutdowns-in-golang-with-signal-notify-context.md)
1. [Graceful shutdown with Go http servers and Kubernetes rolling updates](docs/tr-37-graceful-shutdown-with-go-http-servers-and-kubernetes-rolling-updates.md)
1. [The awesomeness of the httptest package in Go](docs/tr-38-golang-mockmania-httptest.md)
1. [Concurrency Patterns in Go: sync.WaitGroup](docs/tr-39-concurrency-patterns-in-go-sync-waitgroup.md)
1. [Faking stdin and stdout in Go](docs/tr-40-faking-stdin-and-stdout-in-go.md)
1. [How I Structure Web Servers in Go](docs/tr-41-golang-structure-web-servers.md)
1. [Logging without losing money or context](docs/tr-42-log-without-losing-context.md)
1. [An intro to Go for non-Go developers](docs/tr-43-go-intro.md)
1. [Writing Go CLIs With Just Enough Architecture](docs/tr-44-go-cli-how-to-and-advice.md)
1. [How to handle errors in Go?](docs/tr-45-how-to-handle-errors-in-go-5-rules.md)
1. [An Introduction to Producing and Consuming Kafka Messages in Go](docs/tr-46-an-introduction-to-producing-and-consuming-kafka-messages-in-go.md)
1. [Handling multidomain HTTP requests with simple host switch](docs/tr-47-go-multidomain-host-switch.md)
1. [How to handle signals with Go to graceful shutdown HTTP server](docs/tr-48-handle-signals-to-graceful-shutdown-http-server.md)
1. [Implementing traceroute in Go](docs/tr-49-implementing-traceroute-in-go.md)
1. [Tests Make Your Code Inherently Better](docs/tr-50-tests-make-your-code-inherently-better.md)
1. [Let's build a Full-Text Search engine](docs/tr-51-lets-build-a-full-text-search-engine.md)
1. [Different approaches to HTTP routing in Go](docs/tr-52-go-routing.md)
1. [A Gentle Introduction to Web Services With Go](docs/tr-53-go-web-services.md)
1. [Logging in Go: Choosing a System and Using it](docs/tr-54-golang-logging.md)
1. [How to customize Go's HTTP client](docs/tr-55-customize-http-client.md)
1. [Structuring and testing HTTP handlers in Go](docs/tr-56-structuring-and-testing-http-handlers-in-go.md)
1. [Correlating Logs](docs/tr-57-correlating-logs.md)
1. [Environment variables in Golang](docs/tr-58-environment-variables-in-golang.md)
1. [Automating Go Integration Tests With Docker](docs/tr-59-golang-docker-integration-tests.md)
1. [Increasing http.Server boilerplate](docs/tr-60-increasing-http-server-boilerplate-go.md)
1. [Developing price and currency handling for Go](docs/tr-61-price-currency-handling-go.md)
1. [Rust vs. Go: Why They’re Better Together](docs/tr-78-rust-vs-go-why-theyre-better-together.md)

## Cloud Native

1. [A Comprehensive Guide to Cloud Native: Tools, Practices and Culture](docs/tr-32-a-comprehensive-guide-to-cloud-native-tools-practices-and-culture.md)
1. [The State of Cloud Native: Challenges, Culture and Technology](docs/tr-30-the-state-of-cloud-native-challenges-culture-and-technology.md)
1. [How was the world's largest cloud vendor, AWS, born?](docs/tr-35-how-aws-born.md)
1. [Engineering dependability and fault tolerance in a distributed system](docs/tr-75-engineering-dependability-and-fault-tolerance-in-a-distributed-system.md)
1. [](docs/)
1. [](docs/)

## DevOps

1. [A Developer’s Guide to GitOps](docs/tr-62-gitops-developers-guide.md)
1. [Maximizing Developer Effectiveness](docs/tr-63-developer-effectiveness.md)
1. [DevOps engineer job interview questions](docs/tr-64-top-devops-engineer-interview-questions.md)
1. [Podman and Buildah for Docker users](docs/tr-65-podman-and-buildah-for-docker-users.md)


## Others

1. [Fix the Problem, Not the Symptoms](docs/tr-24-fix-problem-not-symptoms.md)
1. [Why do (some) engineers write bad code?](docs/tr-25-why-engineers-write-bad-code.md)
1. [Slack Outage Monday January 4, 2021](docs/tr-66-slack-outage.md)
1. [Preparing to Issue 200 Million Certificates in 24 Hours](docs/tr-74-200m-certs-24hrs.md)
1. [Everything is broken, and it’s okay](docs/tr-76-failure-is-okay.md)
1. [](docs/)
1. [](docs/)