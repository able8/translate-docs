# Making Kubernetes Multi-tenant

# 让 Kubernetes 多租户

-  January  6, 2022

- 2022 年 1 月 6 日

# Making Kubernetes Multi-Tenant

# 制作 Kubernetes 多租户

Building a platform or "containers-as-a-service" is very appealing to development teams - they can simply build containers and never worry  about infrastructure again! While trying to build VT's platform, I often was advised and pushed to do multi-tenancy by going multi-cluster. But, that's a significant overhead in cost, maintenance, and support,  especially when each cluster is basically the same (it does make sense  to use different clusters for different needs though). Instead, we tried to figure out how to make multi-tenancy work within a single cluster. This blog post highlights the major pillars that make it all possible.

构建平台或“容器即服务”对开发团队非常有吸引力——他们可以简单地构建容器，而无需再担心基础设施！在尝试构建 VT 的平台时，我经常被建议并推动通过多集群来进行多租户。但是，这是成本、维护和支持方面的重大开销，尤其是当每个集群基本相同时（尽管使用不同的集群来满足不同的需求确实有意义）。相反，我们试图弄清楚如何使多租户在单个集群中工作。这篇博文重点介绍了使这一切成为可能的主要支柱。

## Defining a Tenant

## 定义租户

For us, the definition of a tenant is very loose. It might be an  entire team. Or an environment (prod vs dev). Or a single application. Or even an environment for a specific application (prod vs dev for an  app). To us, as a platform team, we don't really care. We treat them all the same and give our application teams the choice on how they want to  divide their workloads.

对我们来说，租户的定义非常宽松。可能是整个团队。或环境（产品与开发）。或单个应用程序。甚至是特定应用程序的环境（应用程序的产品与开发）。对我们来说，作为一个平台团队，我们并不在乎。我们对它们一视同仁，并让我们的应用程序团队可以选择他们希望如何分配工作负载。

When mapping to Kubernetes, each tenant has its own namespace.

映射到 Kubernetes 时，每个租户都有自己的命名空间。

## The Four Tenets of Multi-Tenancy

## 多租户的四个原则

The four tenets describe the major components on making multi-tenancy work. Each block adds to the previous, but any missing blocks introduce gaps in protection. We'll dive into each one in greater detail later in the post.

这四个原则描述了使多租户工作的主要组成部分。每个块都添加到前一个块，但任何缺失的块都会引入保护空白。我们将在后面的文章中更详细地介绍每一个。

- **Network Isolation** - ensure applications can't talk to each other unless explicitly authorized to do so
- **Node pooling** - to reduce noisy neighbor problems, provide better cost accounting, and a greater security boundary,  various pools of nodes should be used for tenants
- **Identity and Access Management** - tenants need the ability to both make changes in the cluster and query resources
- **Additional Policy Enforcement** - the built-in RBAC in Kubernetes provides a lot of support, but needs additional policy  control to ensure tenants cannot step on each others' toes

- **网络隔离** - 确保应用程序不能相互通信，除非明确授权这样做
- **节点池** - 为了减少嘈杂的邻居问题，提供更好的成本核算和更大的安全边界，应为租户使用各种节点池
- **身份和访问管理** - 租户需要能够在集群中进行更改和查询资源
- **额外的策略执行** - Kubernetes 内置的 RBAC 提供了很多支持，但需要额外的策略控制以确保租户不会踩到对方的脚趾

![Diagram of a fence between two tenants with panels with text containing the four pillars of multi-tenancy](https://blog.mikesir87.io/images/multi-tenant-fence.png)

## Tenet #1 - Network Isolation

## 宗旨 #1 - 网络隔离

For our platform, we want to ensure applications in different  tenants/namespaces can't reach each other in-cluster. We want to support an "out-and-back" approach, meaning all requests would go out of the  cluster, through the external load balancer, and through the ingress  controller to land at the right application.

对于我们的平台，我们希望确保不同租户/命名空间中的应用程序无法在集群内相互访问。我们希望支持“out-and-back”方法，这意味着所有请求都会离开集群，通过外部负载均衡器，并通过入口控制器到达正确的应用程序。

![Diagram of a request needing to go out and back in](https://blog.mikesir87.io/images/multi-tenant-network-out-and-back.png)

The only exception to this would be platform services. For example,  the Ingress controller obviously needs to send traffic across  namespaces. The same would be true for a central Prometheus instance  scrapping metrics from tenant applications.

唯一的例外是平台服务。例如，Ingress 控制器显然需要跨命名空间发送流量。对于从租户应用程序中提取指标的中央 Prometheus 实例也是如此。

To pull this off, you have to look at the specific CNI (Container  Network Interface) you are using in your cluster. For us, we're using [Calico](https://www.tigera.io/project-calico/). As such, we would simply define a network policy that restricts traffic across namespaces. The policy below denies all network ingress from  namespaces that have the specified label (which is one we place on all  tenant namespaces).

要实现这一点，您必须查看您在集群中使用的特定 CNI（容器网络接口）。对我们来说，我们正在使用 [Calico](https://www.tigera.io/project-calico/)。因此，我们只需定义一个网络策略来限制跨命名空间的流量。下面的策略拒绝来自具有指定标签的命名空间的所有网络入口（这是我们放置在所有租户命名空间上的标签)。

```
apiVersion: crd.projectcalico.org/v1
kind: NetworkPolicy
metadata:
  name: restrict-other-tenants
  namespace: test-tenant
spec:
  order: 20
  selector: ""
  types:
    - Ingress
  ingress:
    - action: Deny
      source:
        namespaceSelector: platform.it.vt.edu/purpose == 'platform-tenant'
```

As you can see, this policy is applied per tenant within their  namespace. We specifically chose to use an ingress rule to potentially  allow tenants to define their own policies with a higher order that *does* allow traffic from namespaces. Had we chose to use egress rules,  there's a chance a tenant creates an egress rule that starts flooding  another tenant that isn't wanting the traffic.

如您所见，此策略适用于其命名空间内的每个租户。我们特别选择使用入口规则来潜在地允许租户以更高的顺序定义自己的策略，*确实*允许来自命名空间的流量。如果我们选择使用出口规则，则租户可能会创建一个出口规则，开始淹没另一个不想要流量的租户。

## Tenet #2 - Node Pooling 

## 原则 #2 - 节点池

One of the main concerns we heard from customers was the idea that a  misbehaving pod might negatively impact their applications. In addition, after discussions with our senior leadership, they expressed the desire to have an idea of cost accounting to know how much each application is costing. A compromise we landed on was to create various node pools for one or more tenants to share. There's the obvious trade-off that as you create more/smaller node pools, you sacrifice utilization and increase  overhead from system pods (log aggregators, volume plugins, Prometheus  exporters, etc.).

我们从客户那里听到的主要担忧之一是行为不端的 pod 可能会对他们的应用程序产生负面影响。此外，在与我们的高级领导讨论之后，他们表示希望了解成本会计，以了解每个应用程序的成本。我们达成的妥协是创建各种节点池供一个或多个租户共享。明显的权衡是，当您创建更多/更小的节点池时，您会牺牲利用率并增加系统 pod（日志聚合器、卷插件、Prometheus 导出器等）的开销。

### Defining the Node Pools

### 定义节点池

To support node pooling, we originally started with the [Cluster Auto-scaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) project. But, we ran into limitations…

为了支持节点池化，我们最初是从 [Cluster Auto-scaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) 项目开始的。但是，我们遇到了限制……

- We needed to define *many* auto-scaling groups for each node pool
- How should we support EBS volumes (an ASG per AZ?)?
- How can we support mixed types (scheduling some on-demand and some spot instances) and sizes?
- How can we reduce the amount of shared/“magic” names? The ASGs  were defined in one repo, far removed from where the tenant config  itself was defined.

- 我们需要为每个节点池定义 *many* 自动缩放组
- 我们应该如何支持 EBS 卷（每个 AZ 一个 ASG？）？
- 我们如何支持混合类型（调度一些按需实例和一些现场实例）和大小？
- 我们如何减少共享/“神奇”名称的数量？ ASG 是在一个 repo 中定义的，与定义租户配置本身的位置相去甚远。

For us, [Karpenter](https://karpenter.sh) has been a *huge* benefit. While there are still a few shortcomings, it's *super* nice being able to define the node pools simply using K8s objects.

对我们来说，[Kapenter](https://karpenter.sh) 是一个*巨大*的好处。虽然仍然存在一些缺点，但能够简单地使用 K8s 对象定义节点池是非常好的。

```
apiVersion: karpenter.sh/v1alpha5
kind: Provisioner
metadata:
  name: example-pool
spec:
  taints:
    - key: platform.it.vt.edu/node-pool
      value: example-pool
      effect: NoSchedule

  # Scale down empty nodes after low utilization.Defaults to an hour
  ttlSecondsAfterEmpty: 300
  
  # Kubernetes labels to be applied to the nodes
  labels:
    platform.it.vt.edu/cost-code: platform
    platform.it.vt.edu/node-pool: example-pool

  kubeletConfiguration:
    clusterDNS: ["10.100.10.100"]

  provider:
    instanceProfile: karpenter-profile

    securityGroupSelector:
      Name: "*eks_worker_sg"
      kubernetes.io/cluster/sample-cluster: owned
      
    # Tags to be applied to the EC2 nodes themselves
    tags:
      CostCode: platform
      Project: example-pool
      NodePool: example-pool
```

A few specific notes worth mentioning…

一些值得一提的具体说明……

- We specifically put taints on all tenant node pools so pods don't  accidentally get scheduled on them without specifying tolerations (more  on that in a moment)
- We tag the EC2 machines with various tags, including a few [Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html). The `CostCode` is a pseudo-organization level tag while the `Project` is a specific project/tenant name. This allows us to use multiple node pools for the same organization.
- Since we are using the Calico CNI, we need to specify the clusterDNS address.

- 我们专门在所有租户节点池上放置了污点，这样 pod 就不会在没有指定容忍度的情况下意外地被安排在它们上面（稍后会详细介绍）
- 我们用各种标签标记 EC2 机器，包括一些 [成本分配标签](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)。 `CostCode` 是一个伪组织级别的标签，而 `Project` 是一个特定的项目/租户名称。这允许我们为同一个组织使用多个节点池。
- 由于我们使用的是 Calico CNI，我们需要指定 clusterDNS 地址。

### Forcing Pods into their Node Pools

### 强制 Pod 进入它们的节点池

Simply defining a node pool doesn't mean that tenants will use it. And, as a platform team, we don't want teams to have to worry about the  pools at all. It would be best if it were completely invisible to them.

简单地定义一个节点池并不意味着租户会使用它。而且，作为一个平台团队，我们根本不希望团队担心池。最好是对他们完全不可见。

Using the [Gatekeeper Mutation](https://open-policy-agent.github.io/gatekeeper/website/docs/mutation/) support, we can define a policy that will mutate all pods to add a  nodeSelector and toleration to ensure the pods are scheduled into the  correct pool. And… it's done in a way that ensures tenants *can't* get around it. By using nodeSelectors, tenants can use the nodeAffinity config to provide Karpenter more configuration (to use spot instances,  ARM machines, etc.).

使用 [Gatekeeper Mutation](https://open-policy-agent.github.io/gatekeeper/website/docs/mutation/) 支持，我们可以定义一个策略来改变所有 pod 以添加 nodeSelector 和 toleration 以确保豆荚被安排到正确的池中。而且......它的完成方式可以确保租户*无法*绕过它。通过使用 nodeSelectors，租户可以使用 nodeAffinity 配置为 Karpenter 提供更多配置（使用 Spot 实例、ARM 机器等)。

```
apiVersion: mutations.gatekeeper.sh/v1beta1
kind: Assign
metadata:
  name: example-tenant-nodepool-selector
  namespace: gatekeeper-system
spec:
  applyTo:
    - groups: [""]
      kinds: ["Pod"]
      versions: ["v1"]
  match:
    scope: Namespaced
    kinds:
      - apiGroups: ["*"]
        kinds: ["Pod"]
    namespaces: ["example-tenant"]
  location: "spec.nodeSelector"
  parameters:
    assign:
      value:
        platform.it.vt.edu/node-pool: example-pool
---
apiVersion: mutations.gatekeeper.sh/v1beta1
kind: Assign
metadata:
  name: example-pool-nodepool-toleration
  namespace: gatekeeper-system
spec:
  applyTo:
    - groups: [""]
      kinds: ["Pod"]
      versions: ["v1"]
  match:
    scope: Namespaced
    kinds:
      - apiGroups: ["*"]
        kinds: ["Pod"]
    namespaces: ["example-tenant"]
  location: "spec.tolerations"
  parameters:
    assign:
      value:
        - key: platform.it.vt.edu/node-pool
          operator: "Equal"
          value: "example-pool"
```

With this, all Pods defined in the `example-tenant` namespace will have a nodeSelector and toleration added that will force the Pod to run on nodes in the `example-pool` node pool. Karpenter will manage the nodes and scale up and down as needed.

这样，在 `example-tenant` 命名空间中定义的所有 Pod 都将添加一个 nodeSelector 和容忍度，这将强制 Pod 在 `example-pool` 节点池中的节点上运行。 Karpenter 将根据需要管理节点并扩大和缩小规模。

## Tenet #3 - Identity Access and Management

## 宗旨 #3 - 身份访问和管理

In order to run a successful platform, we want to ensure the platform team is not a bottleneck to deploying updates or troubleshooting  issues. As such, we want to give as much control back to the teams, but  do so in a safe way.

为了运行一个成功的平台，我们希望确保平台团队不会成为部署更新或故障排除问题的瓶颈。因此，我们希望将尽可能多的控制权交还给团队，但要以安全的方式进行。

### Making Changes to the Cluster

### 对集群进行更改

For our platform, we are using [Flux](https://fluxcd.io)  to manage the deployments. Each tenant is given a "manifest repo" where  they can update manifests and have them applied in the cluster. This  prevents the need to distribute credentials for CI pipelines to make  changes, etc. By leveraging webhooks, changes are applied *very* quickly.

对于我们的平台，我们使用 [Flux](https://fluxcd.io) 来管理部署。每个租户都有一个“清单存储库”，他们可以在其中更新清单并将它们应用到集群中。这避免了为 CI 管道分发凭据以进行更改等的需要。通过利用 webhook，更改可以*非常*快速地应用。

![Diagram showing how Flux is working](https://blog.mikesir87.io/images/multi-tenant-applying-changes.png)

### Providing Read-only Access to the Cluster

### 提供对集群的只读访问权限

To allow teams to troubleshoot and debug issues, we provide the  ability for them to query their resources in a read-only manner. We're  currently using the following tools to pull it off:

为了让团队能够解决和调试问题，我们为他们提供了以只读方式查询其资源的能力。我们目前正在使用以下工具来实现它：

- [Dex](https://dexido.io) - a Federated OpenID Provider  that performs authentication of users. We have this configured as an  OAuth client to our central VT Gateway service, ensuring all auth is  two-factored and backed by our VT identity systems
- [Kube OIDC Proxy](https://github.com/jetstack/kube-oidc-proxy) - this serves as an API proxy that uses the tokens issues by Dex as  authentication. It then impersonates the requests to the underlying k8s  API, passing along the user's username and group memberships.

- [Dex](https://dexido.io) - 执行用户身份验证的联合 OpenID 提供程序。我们已将其配置为中央 VT 网关服务的 OAuth 客户端，确保所有身份验证都是双重因素并由我们的 VT 身份系统提供支持
- [Kube OIDC 代理](https://github.com/jetstack/kube-oidc-proxy) - 这是一个 API 代理，使用 Dex 颁发的令牌作为身份验证。然后它模拟对底层 k8s API 的请求，传递用户的用户名和组成员身份。

With those two deployed, we can then create `RoleBinding` objects that authorize specific groups (from our central identity  system) to have read-only access to specific namespaces (defined in a `ClusterRole` named `platform-tenant`).

部署这两个后，我们可以创建“RoleBinding”对象，授权特定组（来自我们的中央身份系统）对特定命名空间（在名为“platform-tenant”的“ClusterRole”中定义）具有只读访问权限。

```
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: tenant-access
  namespace: example-tenant
subjects:
  - kind: Group
    name: oidc:sample-org.team.developers
    apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: platform-tenant
  apiGroup: rbac.authorization.k8s.io
```

The big advantage of using the Kube OIDC proxy is that we can share  the OIDC tokens and allow users to configure their local Kubeconfig  files with the same credentials, allowing them to use `kubectl` and other tools for additional querying.

使用 Kube OIDC 代理的一大优势是我们可以共享 OIDC 令牌并允许用户使用相同的凭据配置其本地 Kubeconfig 文件，从而允许他们使用“kubectl”和其他工具进行额外查询。

We've also built a dashboard that is a client of the Dex OIDC  provider. As user's authenticate, the dashboard simply forwards the  logged-in user's tokens to the API to query the resources. That makes  our dashboard completely unprivileged and removes any ability of a user  seeing something they shouldn't be able to see.

我们还构建了一个仪表板，它是 Dex OIDC 提供商的客户端。作为用户的身份验证，仪表板只是将登录用户的令牌转发到 API 以查询资源。这使得我们的仪表板完全没有特权，并消除了用户看到他们不应该看到的东西的任何能力。

## Tenet #4 - Additional Policy Enforcement

## 宗旨 #4 - 额外的政策执行

While the built-in Kubernetes RBAC provides a lot of capability,  there are times in which we want to get more granular. A few examples:

虽然内置的 Kubernetes RBAC 提供了很多功能，但有时我们希望获得更精细的功能。几个例子：

- Tenants should be able to create Services, but how can we prevent them from creating NodePort or LoadBalancer Services?
- How do we limit the domains one tenant can use for  Ingress/Certification objects so they don't intercept the traffic meant  for another?
- How can we enforce the [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/) to prevent pod from gaining access to the underlying host?

- 租户应该能够创建服务，但我们如何防止他们创建 NodePort 或 LoadBalancer 服务？
- 我们如何限制一个租户可以用于 Ingress/Certification 对象的域，以便它们不会拦截指向另一个租户的流量？
- 我们如何执行 [Pod 安全标准](https://kubernetes.io/docs/concepts/security/pod-security-standards/) 以防止 pod 获得对底层主机的访问权限？

Fortunately, we can use [admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to plug in our own policies as part of the API request process. Rather than writing our own services, we can leverage [Gatekeeper](https://open-policy-agent.github.io/gatekeeper/website/docs/) and write OPA policies. And with OPA, we can easily write unit tests to catch and prevent regressions. 

幸运的是，我们可以使用 [admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) 在 API 请求过程中插入我们自己的策略。我们可以利用 [Gatekeeper](https://open-policy-agent.github.io/gatekeeper/website/docs/) 并编写 OPA 策略，而不是编写我们自己的服务。使用 OPA，我们可以轻松编写单元测试来捕捉和防止回归。

If you're not familiar with Gatekeeper, it's basically a wrapper  around the OPA engine that allows us to write our policies with  Kubernetes objects and run them whenever objects are being created or  updated in the cluster. In addition, we can apply the same policies in  different ways for different tenants.

如果您不熟悉 Gatekeeper，它基本上是 OPA 引擎的包装器，它允许我们使用 Kubernetes 对象编写策略，并在集群中创建或更新对象时运行它们。此外，我们可以针对不同的租户以不同的方式应用相同的策略。

As an example, a `ConstraintTemplate` that defines an `AuthorizedDomainPolicy` could then be used to define the authorized domains for each namespace. The sample object below will enforce that policy in the `example-tenant` namespace and ensure we only use the specified names. To help with on-boarding, we also automatically authorize `<tenant-id>.tenants.platform.it.vt.edu` to be used by each tenant. They are welcome to CNAME any other name and point to the cluster to use more friendly names. All we have to do is  update their policy.

例如，定义“AuthorizedDomainPolicy”的“ConstraintTemplate”可以用来定义每个命名空间的授权域。下面的示例对象将在 `example-tenant` 命名空间中强制执行该策略，并确保我们只使用指定的名称。为了帮助入职，我们还自动授权每个租户使用`<tenant-id>.tenants.platform.it.vt.edu`。欢迎他们使用 CNAME 任何其他名称并指向集群以使用更友好的名称。我们所要做的就是更新他们的政策。

```
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: AuthorizedDomainPolicy
metadata:
  name: example-tenant
spec:
  match:
    namespaces:
      - example-tenant
    kinds:
      - apiGroups: ["extensions", "networking.k8s.io"]
        kinds: ["Ingress"]
      - apiGroups: ["cert-manager.io"]
        kinds: ["Certificate"]
  parameters:
    domains:
      - example-tenant.tenants.platform.it.vt.edu
      - "*.example-tenant.tenants.platform.it.vt.edu"
      - smaller-url.example.com
```

## Putting it All Together

## 把它们放在一起

While it feels like there's a lot to define for each tenant, it's  very repetitive. Once we had defined a few tenants, we were able to  figure out how to build a Helm chart that defines all of the various  objects for each tenant. All we have to do is build and use a values  file that defines the tenants, their config, and the node pools. An  example values file is below. Today, the values for each cluster is  stored in Git, but eventually, it might be sourced and built from a  developer portal (maybe Backstage???).

虽然感觉每个租户都有很多定义，但它非常重复。一旦我们定义了一些租户，我们就能够弄清楚如何构建一个 Helm 图表来定义每个租户的所有各种对象。我们所要做的就是构建和使用定义租户、他们的配置和节点池的值文件。下面是一个示例值文件。今天，每个集群的值都存储在 Git 中，但最终，它可能是从开发人员门户（可能是后台？？？）获取和构建的。

```
global:
  clusterHostname: cluster.hostname
  tenantHostPrefix: tenants-example.platform.it.vt.edu
  clusterName: vt-common-platform-dvlp-cluster
  clusterDNS: 10.1.0.10

nodePools:
  sample-team-apps:
    instances:
      capacityTypes: ["on-demand"]
      instanceTypes: ["t3a.large"]
    costCode: org1
  sample-team-gitlab-jobs:
    instances:
      capacityTypes: ["on-demand", "spot"]
      instanceTypes: ["t3a.medium"]
      emptyTtl: 600
      ttl: 604800
    costCode: org1
tenants:
  sample-team-gitlab-ci:
    nodePool: sample-team-gitlab-jobs
    operatorEdGroup: team.devs.platform-access
  nis-customer-portal:
    nodePool: sample-team-apps
    operatorEdGroup: team.devs.platform-access
    domains:
      - sample-team.org.vt.edu
      - "*.sample-team.org.vt.edu"
```

## Wrapping Up

##  包起来

I know there's a lot of material here. If you're interested in  learning more, I'm happy to do follow-up posts. Let me know what sounds  interesting and what you want to read more about. 

我知道这里有很多材料。如果您有兴趣了解更多信息，我很乐意做后续帖子。让我知道什么听起来很有趣，以及你想了解更多。

