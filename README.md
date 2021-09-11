# translate-docs

使用 [Google 翻译 API](https://translate.google.cn/)，将输入 **英文** (Markdown)文档按段落生成 **中英** 对照翻译文档。

阅读英文文档，有时理解不准确；中英文对照阅读便于理解；边学技术边学英语。

#### 如何使用

```
./translate_darwin_amd64 -f example-input.md
```

#### go build 生成 Mac、linux、Windows 平台可执行文件

```sh
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/translate_darwin_amd64 translate.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/translate_linux_amd64 translate.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/translate_windows_amd64.exe translate.go
```

---

# Translated docs & blogs

## Docker, Container, and Kubernetes

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
1. [Architecting Kubernetes clusters — choosing the best autoscaling strategy](docs/tr-70-kubernetes-autoscaling-strategies.md)
1. [Load balancing and scaling long-lived connections in Kubernetes](docs/tr-71-kubernetes-long-lived-connections.md)
1. [Kubernetes Operators Best Practices](docs/tr-72-kubernetes-operators-best-practices.md)
1. [7 Principles of DevSecOps With Kubernetes](docs/tr-79-seven-principles-of-devsecops-with-kubernetes.md)
1. [Kubernetes Is Not Your Platform, It's Just the Foundation](docs/tr-77-kubernetes-successful-adoption-foundation.md)
1. [Moving k8s communication to gRPC](docs/tr-80-moving-k8s-communication-to-grpc.md)
1. [Why you need a platform team for Kubernetes](docs/tr-81-why-you-need-a-platform-team-for-kubernetes.md)
1. [A brief overview of the Container Network Interface (CNI) in Kubernetes](docs/tr-82-cni-kubernetes.md)
1. [Six Strategies for Application Deployment](docs/tr-85-deployment-strategies.md)
1. [Annotating Kubernetes Services for Humans](docs/tr-86-annotating-k8s-for-humans.md)
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
1. [How to use CoreDNS Effectively with Kubernetes](docs/tr-using-coredns-effectively-kubernetes.md)
1. [What’s new in Kubernetes 1.20?](docs/tr-whats-new-kubernetes-1-20.md)
1. [Migrating from Docker to Podman](docs/tr-migrating-to-podman.md)
1. [What Configuration Management is and Why You Should Implement it in Your infrastructure](docs/tr-kubernetes-configuration-management-101.md)
1. [Kubernetes production best practices](docs/tr-production-best-practices.md)
1. [Liveness Probes are Dangerous](docs/tr-kubernetes-liveness-probes-are-dangerous.md)
1. [Debugging network stalls on Kubernetes](docs/tr-2019-11-21-debugging-network-stalls-on-kubernetes.md)
1. [Why Continuous Infrastructure Automation is no longer optional](docs/tr-why-continuous-infrastructure-automation-is-no-longer-optional-8b3e6bc8847.md)
1. [Why you don't have to be afraid of Kubernetes](docs/tr-kubernetes-complex-business-problem.md)
1. [Introduction to open source observability on Kubernetes](docs/tr-open-source-observability-kubernetes.md)
1. [Using Kubernetes to rethink your system architecture and ease technical debt](docs/tr-rethinking-system-architecture-can-kubernetes-help-to-solve-rewrite-anxiety.md)
1. [Everyone might be a cluster-admin in your Kubernetes cluster](docs/tr-everyone-might-be-cluster-admin-your-kubernetes-cluster.md)
1. [Kubernetes Performance Trouble Spots: Airbnb’sTake](docs/tr-kubernetes-performance-troublespots-airbnbs-take.md)
1. [Rolling Updates and Blue-Green Deployments with Kubernetes and HAProxy](docs/tr-rolling-updates-and-blue-green-deployments-with-kubernetes-and-haproxy.md)
1. [Scaling to 100k Users](docs/tr-scaling-100k.md)
1. [How I prepared & passed the Certified Kubernetes Administrator (CKA) Exam](docs/tr-cka-exam-tips.md)
1. [NTP in a Kubernetes cluster](docs/tr-manage-ntp-using-kubernetes.md)
1. [Understanding Kubernetes: Part 1-Pods](docs/tr-understanding-kubernetes-pods.md)
1. [What is Kubernetes?](docs/tr-what-is-kubernetes-529.md)
1. [How to Build Production-Ready Kubernetes Clusters and Containers](docs/tr-how-to-build-production-ready-kubernetes-clusters-and-containers.md)
1. [Introducing istiod: simplifying the control plane](docs/tr-istiod.md)
1. [7 best practices: Building applications for containers and Kubernetes](docs/tr-kubernetes-best-practices-building-applications-containers.md)
1. [Kubernetes operators: Embedding operational expertise side by side with containerized applications](docs/tr-kubernetes-operators.md)
1. [Cloud Native Backups, Disaster Recovery and Migrations on Kubernetes](docs/tr-cloud-native-backups-disaster-recovery-and-migrations-on-kubernetes.md)
1. [Core Kubernetes: Jazz Improv over Orchestration](docs/tr-core-kubernetes-jazz-improv-over-orchestration-a7903ea92ca.md)
1. [Kubernetes Finalizers](docs/tr-finalizers.md)
1. [Using Finalizers to Control Deletion](docs/tr-using-finalizers-to-control-deletion.md)
1. [Why Kubernetes Operators Will Unleash Your Developers by Reducing Complexity](docs/tr-why-kubernetes-operators-will-unleash-your-developers-by-reducing-complexity.md)
1. [Controlling outbound traffic from Kubernetes](docs/tr-controlling-outbound-traffic-from-kubernetes.md)
1. [Who Needs a Dashboard? Why the Kubernetes Command Line Is Not Enough](docs/tr-who-needs-a-dashboard-why-the-kubernetes-command-line-is-not-enough.md)
1. [How a Kubernetes Pod Gets an IP Address](docs/tr-how-a-kubernetes-pod-gets-an-ip-address.md)
1. [How to Champion GitOps in Your Organization](docs/tr-how-to-champion-gitops-in-your-organization.md)
1. [Improvements to the Ingress API in Kubernetes 1.18](docs/tr-improvements-to-the-ingress-api-in-kubernetes-1.18.md)
1. [Migrating to Kubernetes](docs/tr-migrating-to-kubernetes.md)
1. [10 Best Practices Worth Implementing to Adopt Kubernetes](docs/tr-10-best-practices-worth-implementing-to-adopt-kubernetes.md)
1. [Contributing to the Development Guide](docs/tr-contributing-to-the-development-guide.md)
1. [Sidecar Proxy Pattern - The Basis Of Service Mesh](docs/tr-service-proxy-pod-sidecar-oh-my.md)
1. [](docs/)
1. [](docs/)

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
1. [The underutilized usefulness of sync.Once](docs/tr-83-synconce.md)
1. [How I organize packages in Go](docs/tr-how-i-organize-packages-in-go.md)
1. [Writing TCP scanner in Go](docs/tr-tcp-scanner-in-go.md)
1. [Wrapping commands in Go](docs/tr-wrapping-commands-with-go.md)
1. [Writing a reverse proxy in Go](docs/tr-writing-proxy-in-go.md)
1. [Let's Create a Simple Load Balancer With Go](docs/tr-lets-create-a-simple-lb-go.md)
1. [Courier: Dropbox migration to gRPC](docs/tr-courier-dropbox-migration-to-grpc.md)
1. [](docs/)
1. [](docs/)

## DevOps and Cloud Native

1. [A Comprehensive Guide to Cloud Native: Tools, Practices and Culture](docs/tr-32-a-comprehensive-guide-to-cloud-native-tools-practices-and-culture.md)
1. [The State of Cloud Native: Challenges, Culture and Technology](docs/tr-30-the-state-of-cloud-native-challenges-culture-and-technology.md)
1. [How was the world's largest cloud vendor, AWS, born?](docs/tr-35-how-aws-born.md)
1. [Engineering dependability and fault tolerance in a distributed system](docs/tr-75-engineering-dependability-and-fault-tolerance-in-a-distributed-system.md)
1. [Introduction to eBPF](docs/tr-69-introduction-to-ebpf.md)
1. [A Developer’s Guide to GitOps](docs/tr-62-gitops-developers-guide.md)
1. [Maximizing Developer Effectiveness](docs/tr-63-developer-effectiveness.md)
1. [DevOps engineer job interview questions](docs/tr-64-top-devops-engineer-interview-questions.md)
1. [Podman and Buildah for Docker users](docs/tr-65-podman-and-buildah-for-docker-users.md)
1. [How to learn and stay up to date with DevOps and Cloud Native technologies](docs/tr-how-to-learn-and-stay-up-to-date-with-devops-and-cloud-native-technologies-44526658a4fb.md)
1. [The Next Step after DevOps and GitOps Is Cloud Engineering, Pulumi Says](docs/tr-the-next-step-after-devops-and-gitops-is-cloud-engineering-pulumi-says.md)
1. [ArgoCD: Okta integration, and user groups](docs/tr-argocd-okta-integration-and-user-groups.md)
1. [Argo CD RBAC Configuration](docs/tr-argo-cd-rbac.md)
1. [DevOps, SRE, and Platform Engineering](docs/tr-dev-devops-sre.md)
1. [How to Keep Your Cloud-Native Apps Secure](docs/tr-How-to-Keep-Your-Cloud-Native-Apps-Secure.md)
1. [10 best practices for DevOps](docs/tr-10-best-practices-for-devops.md)
1. [DevOps: A cheat sheet](docs/tr-devops-the-smart-persons-guide.md)
1. [Six steps to DevOps success, analyzed](docs/tr-six-steps-to-successful-devops-success-analyzed.md)
1. [3 lessons for IT leaders from “The Unicorn Project”](docs/tr-3-lessons-it-leaders-unicorn-project.md)
1. [5 GitOps Best Practices](docs/tr-5-gitops-best-practices-d95cb0cbe9ff.md)
1. [Extending cloud native principles to chaos engineering](docs/tr-cloud-native-chaos-engineering-enhancing-kubernetes-application-resiliency.md)
1. [Debugging Software Deployments with strace](docs/tr-deployment_debugging_strace.md)
1. [Gracefully shutting down a Nodejs HTTP server](docs/tr-gracefully-shutting-down-a-nodejs-http-server.md)
1. [You can't buy DevOps](docs/tr-you-cant-buy-devops.md)
1. [How to be a DevOps maestro: containers orchestration guide](docs/tr-how-to-be-a-devops-maestro-containers-orchestration-guide-b2cf884eaed1.md)
1. [5 Best Practices for Nailing Incident Retrospectives](docs/tr-5-best-practices-nailing-postmortems.md)
1. [DevOps: 5 tips on how to manage stronger remote teams](docs/tr-devops-how-manage-remote-teams-tips.md)
1. [7 top DevOps engineer interview questions for 2020](docs/tr-devops-engineer-interview-questions-2020.md)
1. [DevOps culture: 5 questions to ask about yours](docs/tr-devops-culture-5-questions.md)
1. [3 problems DevOps won't fix](docs/tr-devops-wont-fix-3-problems.md)
1. [(A few) Ops Lessons We All Learn The Hard Way](docs/tr-ops-lessons.md)
1. [How we migrated Dropbox from Nginx to Envoy](docs/tr-how-we-migrated-dropbox-from-nginx-to-envoy.md)
1. [The Future of Ops Jobs](docs/tr-the-future-of-ops-jobs.md)
1. [The Future of Cloud Native Security Is Self-Healing](docs/tr-the-future-of-cloud-native-security-is-self-healing.md)
1. [OpenTelemetry Steps up to Manage the Mayhem of Microservices](docs/tr-opentelemetry-steps-up-to-manage-the-mayhem-of-microservices.md)
1. [Ops by pull request: an Ansible GitOps story](docs/tr-ops-by-pull-request-an-ansible-gitops-story.md)
1. [How GitOps Improves the Security of Your Development Pipelines](docs/tr-how-gitops-improves-security-development-pipelines.md)
1. [How eBPF Turns Linux into a Programmable Kernel](docs/tr-how-ebpf-turns-linux-into-a-programmable-kernel.md)
1. [Backup and DR in the Age of GitOps](docs/tr-backup-and-dr-in-the-age-of-gitops.md)
1. [Why I've Been Merging Microservices Back Into The Monolith At InVision](docs/tr-3944-why-ive-been-merging-microservices-back-into-the-monolith-at-invision.htm.md)
1. [Help the World by Healing Your NGINX Configuration](docs/tr-help-the-world-by-healing-your-nginx-configuration.md)
1. [Patterns of Distributed Systems](docs/tr-patterns-of-distributed--systems.md)
1. [](docs/)
1. [](docs/)
1. [](docs/)

## Others

1. [Fix the Problem, Not the Symptoms](docs/tr-24-fix-problem-not-symptoms.md)
1. [Why do (some) engineers write bad code?](docs/tr-25-why-engineers-write-bad-code.md)
1. [Slack Outage Monday January 4, 2021](docs/tr-66-slack-outage.md)
1. [Preparing to Issue 200 Million Certificates in 24 Hours](docs/tr-74-200m-certs-24hrs.md)
1. [Everything is broken, and it’s okay](docs/tr-76-failure-is-okay.md)
1. [How to Successfully Hand Over Systems](docs/tr-87-how-to-successfully-hand-over-systems.md)
1. [The Problem With Agile Scrum (And Why We Use Kanban Instead)](docs/tr-88-why-cloudzero-uses-kanban.md)
1. [Enforcing Policy as Code using OPA and Gatekeeper in Kubernetes](docs/tr-enforcing-policy-as-code-using-opa-and-gatekeeper-in-kubernetes.md)
1. [APIs Aren’t Just for Tech Companies](docs/tr-apis-arent-just-for-tech-companies.md)
1. [OKRs v KPIs - How they work together](docs/tr-kr-v-kpis-a-helpful-primer.md)
1. [The History of HAProxy](docs/tr-the-history-of-haproxy.md)
1. [The Secret of Employee Engagement](docs/tr-employee-engagement.md)
1. [Software is about people, not code](docs/tr-software-is-about-people-not-code.md)
1. [How an SRE became an Application Security Engineer (and you can too)](docs/tr-how-an-sre-became-an-application-security-engineer-and-you-can-too.md)
1. [The 4 Quadrants of Time Management Matrix](docs/tr-4-quadrants-of-time-management.md)
1. [GitLab and Okta](docs/tr-gitlab-okta.md)
1. [Why Software Development is Hard](docs/tr-2021_01_01_why_software_development_is_hard.md)
1. [The Difference Between LDAP and SAML SSO](docs/tr-difference-ldap-saml-sso.md)
1. [GitLab and Okta](docs/tr-gitlab-okta.md)
1. [SAML SSO for GitLab.com groups](docs/tr-gitlab_saml_sso.md)
1. [SCIM provisioning using SAML SSO for GitLab.com groups](docs/tr-gitlab_scim_setup.md)
1. [How to Defeat Busy Culture](docs/tr-how-to-defeat-busy-culture.md)
1. [Things I Learned to Become a Senior Software Engineer](docs/tr-things-I-learned-to-become-a-senior-software-engineer.md)
1. [How to familiarize yourself with a new codebase](docs/tr-how-to-familiarize-yourself-with-a-new-codebase.md)
1. [Taming the tar command: Tips for managing backups in Linux](docs/tr-taming-tar-command.md)
1. [Advice for customers dealing with Docker Hub rate limits, and a Coming Soon announcement](docs/tr-advice-for-customers-dealing-with-docker-hub-rate-limits-and-a-coming-soon-announcement.md)
1. [Millions of Remote Workers Are Now Thinking About Moving](docs/tr-millions-of-remote-workers-are-now-thinking-about-moving.md)
1. [What I Learned From Bombing a Technical Interview](docs/tr-what-i-learned-from-bombing-a-technical-interview.md)
1. [How Netflix brings safer and faster streaming experiences to the living room on crowded networks using TLS 1.3](docs/tr-how-netflix-brings-safer-and-faster-streaming-experience-to-the-living-room-on-crowded-networks-78b8de7f758c.md)
1. [Managing cgroups with systemd](docs/tr-cgroups-part-four.md)
1. [Recommended Engineering Management Books](docs/tr-recommended-engineering-management-books.md)
1. [](docs/)
1. [](docs/)
1. [Go string handling overview [cheat sheet]](docs/tr-string-functions-reference-cheat-sheet.md)
1. [Get year, month, day from time](docs/tr-day-month-year-from-time.md)
1. [Regexp tutorial and cheat sheet](docs/tr-regexp-cheat-sheet.md)
1. [Get year, month, day from time](docs/tr-day-month-year-from-time.md)
1. [Write log to file (or /dev/null)](docs/tr-log-to-file.md)
1. [Regexp tutorial and cheat sheet](docs/tr-regexp-cheat-sheet.md)
