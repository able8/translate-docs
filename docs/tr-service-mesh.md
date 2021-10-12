# The Istio service mesh

# Istio 服务网格

Istio addresses the challenges developers and operators face with a distributed or microservices architecture. Whether you're building from scratch or migrating existing applications to cloud native, Istio can help.

Istio 解决了开发人员和运营商在分布式或微服务架构方面面临的挑战。无论您是从头开始构建还是将现有应用程序迁移到云原生，Istio 都可以提供帮助。

![Service mesh](http://istio.io/latest/img/service-mesh.svg)

#### By adding a proxy "sidecar" along with every application deployed, Istio lets you program application-aware traffic management, incredible observability, and robust security capabilities into your network.

#### 通过在部署的每个应用程序中添加代理“sidecar”，Istio 允许您将应用程序感知流量管理、令人难以置信的可观察性和强大的安全功能编程到您的网络中。

## What is a Service Mesh?

## 什么是服务网格？

Modern applications are typically architected as distributed collections of microservices, with each collection of microservices performing some discrete business function. A service mesh is a dedicated infrastructure layer that you can add to your applications. It allows you to transparently add capabilities like observability, traffic management, and security, without adding them to your own code. The term “service mesh” describes both the type of software you use to implement this pattern, and the security or network domain that is created when you use that software.

现代应用程序通常被构建为分布式微服务集合，每个微服务集合执行一些离散的业务功能。服务网格是一个专用的基础设施层，您可以将其添加到您的应用程序中。它允许您透明地添加可观察性、流量管理和安全性等功能，而无需将它们添加到您自己的代码中。术语“服务网格”描述了用于实现此模式的软件类型，以及在使用该软件时创建的安全或网络域。

As the deployment of distributed services, such as in a Kubernetes-based system, grows in size and complexity, it can become harder to understand and manage. Its requirements can include discovery, load balancing, failure recovery, metrics, and monitoring. A service mesh also often addresses more complex operational requirements, like A/B testing, canary deployments, rate limiting, access control, encryption, and end-to-end authentication.

随着分布式服务的部署（例如在基于 Kubernetes 的系统中）规模和复杂性的增长，它可能变得更难以理解和管理。它的要求可以包括发现、负载平衡、故障恢复、指标和监控。服务网格还经常满足更复杂的操作需求，例如 A/B 测试、金丝雀部署、速率限制、访问控制、加密和端到端身份验证。

Service-to-service communication is what makes a distributed application possible. Routing this communication, both within and across application clusters, becomes increasingly complex as the number of services grow. Istio helps reduce this complexity while easing the strain on development teams

服务到服务的通信使分布式应用程序成为可能。随着服务数量的增长，在应用程序集群内和跨应用程序集群路由这种通信变得越来越复杂。 Istio 有助于降低这种复杂性，同时减轻开发团队的压力

## What is Istio?

## Istio 是什么？

Istio is an open source service mesh that layers transparently onto existing distributed applications. Istio’s powerful features provide a uniform and more efficient way to secure, connect, and monitor services. Istio is the path to load balancing, service-to-service authentication, and monitoring – with few or no service code changes. Its powerful control plane brings vital features, including:

Istio 是一个开源服务网格，它透明地分层到现有的分布式应用程序上。 Istio 的强大功能提供了一种统一且更有效的方式来保护、连接和监控服务。 Istio 是实现负载平衡、服务到服务身份验证和监控的途径——几乎没有服务代码更改。其强大的控制平面带来了重要的功能，包括：

- Secure service-to-service communication in a cluster with TLS encryption, strong identity-based authentication and authorization
- Automatic load balancing for HTTP, gRPC, WebSocket, and TCP traffic
- Fine-grained control of traffic behavior with rich routing rules, retries, failovers, and fault injection
- A pluggable policy layer and configuration API supporting access controls, rate limits and quotas
- Automatic metrics, logs, and traces for all traffic within a cluster, including cluster ingress and egress

- 使用 TLS 加密、强大的基于身份的身份验证和授权在集群中安全的服务到服务通信
- HTTP、gRPC、WebSocket 和 TCP 流量的自动负载平衡
- 流量行为的细粒度控制，具有丰富的路由规则、重试、故障转移和故障注入
- 支持访问控制、速率限制和配额的可插拔策略层和配置 API
- 集群内所有流量的自动指标、日志和跟踪，包括集群入口和出口

Istio is designed for extensibility and can handle a diverse range of deployment needs. Istio’s control plane runs on Kubernetes, and you can add applications deployed in that cluster to your mesh, extend the mesh to other clusters, or even connect VMs or other endpoints running outside of Kubernetes.

Istio 专为可扩展性而设计，可以处理各种部署需求。 Istio 的控制平面在 Kubernetes 上运行，您可以将部署在该集群中的应用程序添加到您的网格中，将网格扩展到其他集群，甚至连接在 Kubernetes 之外运行的虚拟机或其他端点。

A large ecosystem of contributors, partners, integrations, and distributors extend and leverage Istio for a wide variety of scenarios.
You can install Istio yourself, or a number of vendors have products that integrate Istio and manage it for you.

由贡献者、合作伙伴、集成商和分销商组成的大型生态系统扩展并利用 Istio 来处理各种场景。
您可以自己安装 Istio，或者许多供应商都有集成 Istio 并为您管理它的产品。

## How it Works

##  这个怎么运作

Istio has two components: the data plane and the control plane.

Istio 有两个组件：数据平面和控制平面。

The data plane is the communication between services. Without a service mesh, the network doesn’t understand the traffic being sent over, and can’t make any decisions based on what type of traffic it is, or who it is from or to.

数据平面是服务之间的通信。如果没有服务网格，网络就无法理解正在发送的流量，也无法根据流量是什么类型或来自谁或向谁做出任何决定。

Service mesh uses a proxy to intercept all your network traffic, allowing a broad set of application-aware features based on configuration you set.

服务网格使用代理来拦截您的所有网络流量，允许基于您设置的配置的一组广泛的应用程序感知功能。

An Envoy proxy is deployed along with each service that you start in your cluster, or runs alongside services running on VMs.

Envoy 代理与您在集群中启动的每个服务一起部署，或者与在 VM 上运行的服务一起运行。

The control plane takes your desired configuration, and its view of the services, and dynamically programs the proxy servers, updating them as the rules or the environment changes.

控制平面采用您想要的配置及其对服务的看法，并对代理服务器进行动态编程，随着规则或环境的变化而更新它们。

![Before utilizing Istio](http://istio.io/latest/img/service-mesh-before.svg)

#### Before utilizing Istio

#### 在使用 Istio 之前

![After utilizing Istio](http://istio.io/latest/img/service-mesh.svg)

#### After utilizing Istio

#### 使用 Istio 后

# Concepts

# 概念

## Traffic management 

## 交通管理

Routing traffic, both within a single cluster and across clusters, affects performance and enables better deployment strategy. Istio’s traffic routing rules let you easily control the flow of traffic and API calls between services. Istio simplifies configuration of service-level properties like circuit breakers, timeouts, and retries, and makes it easy to set up important tasks like A/B testing, canary deployments, and staged rollouts with percentage-based traffic splits.

单个集群内和跨集群的路由流量会影响性能并实现更好的部署策略。 Istio 的流量路由规则让您可以轻松控制服务之间的流量和 API 调用。 Istio 简化了断路器、超时和重试等服务级别属性的配置，并可以轻松设置重要任务，例如 A/B 测试、金丝雀部署和基于百分比的流量拆分的分阶段部署。

## Observability

## 可观察性

As services grow in complexity, it becomes challenging to understand behavior and performance. Istio generates detailed telemetry for all communications within a service mesh. This telemetry provides observability of service behavior, empowering operators to troubleshoot, maintain, and optimize their applications. Even better, you get almost all of this instrumentation without requiring application changes. Through Istio, operators gain a thorough understanding of how monitored services are interacting.

随着服务变得越来越复杂，理解行为和性能变得具有挑战性。 Istio 为服务网格内的所有通信生成详细的遥测数据。这种遥测提供了服务行为的可观察性，使操作员能够对其应用程序进行故障排除、维护和优化。更好的是，您无需更改应用程序即可获得几乎所有这些工具。通过 Istio，运营商可以全面了解受监控服务的交互方式。

Istio’s telemetry includes detailed metrics, distributed traces, and full access logs. With Istio, you get thorough and comprehensive service mesh observability.

Istio 的遥测包括详细的指标、分布式跟踪和完整的访问日志。使用 Istio，您可以获得彻底而全面的服务网格可观察性。

## Security capabilities

## 安全功能

Microservices have particular security needs, including protection against man-in-the-middle attacks, flexible access controls, auditing tools, and mutual TLS. Istio includes a comprehensive security solution to give operators the ability to address all of these issues. It provides strong identity, powerful policy, transparent TLS encryption, and authentication, authorization and audit (AAA) tools to protect your services and data.

微服务具有特殊的安全需求，包括防止中间人攻击、灵活的访问控制、审计工具和双向 TLS。 Istio 包含一个全面的安全解决方案，使运营商能够解决所有这些问题。它提供强大的身份、强大的策略、透明的 TLS 加密以及身份验证、授权和审计 (AAA) 工具来保护您的服务和数据。

Istio’s security model is based on security-by-default, aiming to provide in-depth defense to allow you to deploy security-minded applications even across distrusted networks.

Istio 的安全模型基于默认安全，旨在提供深度防御，让您即使在不受信任的网络中也能部署具有安全意识的应用程序。

# Solutions

# 解决方案

- [**Enabling Defense-in-Depth for Enterprise Applications** Learn more](http://istio.io/latest/about/solutions/enabling-defense-in-depth-for-enterprise-applications/)
- [**Increasing Kubernetes deployment and management efficiency** Learn more](http://istio.io/latest/about/solutions/increasing-kubernetes-deployment-and-management-efficiency/)
- [**Instituting Observability and SRE Best Practices** Learn more](http://istio.io/latest/about/solutions/instituting-observability-and-sre-best-practices/)

- [**为企业应用程序启用深度防御**了解更多](http://istio.io/latest/about/solutions/enabling-defense-in-depth-for-enterprise-applications/)
- [**提高Kubernetes部署和管理效率**了解更多](http://istio.io/latest/about/solutions/increasing-kubernetes-deployment-and-management-efficiency/)
- [**制定可观察性和 SRE 最佳实践** 了解更多](http://istio.io/latest/about/solutions/instituting-observability-and-sre-best-practices/)

[Go to solutions](http://istio.io/latest/about/solutions)

[转到解决方案](http://istio.io/latest/about/solutions)

Was this information useful?

这些信息有用吗？

Do you have any suggestions for improvement?

你有什么改进的建议吗？

Thanks for your feedback! 

感谢您的反馈意见！

