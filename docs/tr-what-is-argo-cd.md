# What Is Argo CD?

# 什么是 Argo CD？

https://argo-cd.readthedocs.io/en/stable/#what-is-argo-cd

Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes.

Argo CD 是 Kubernetes 的声明性 GitOps 持续交付工具。

![Argo CD UI](https://argo-cd.readthedocs.io/en/stable/assets/argocd-ui.gif)

## Why Argo CD?[¶](https://argo-cd.readthedocs.io/en/stable/#why-argo-cd)

## 为什么选择 Argo CD？[¶](https://argo-cd.readthedocs.io/en/stable/#why-argo-cd)

Application definitions, configurations, and environments should be declarative and version controlled. Application deployment and lifecycle management should be automated, auditable, and easy to understand.

应用程序定义、配置和环境应该是声明性的并且是版本控制的。应用程序部署和生命周期管理应该是自动化的、可审计的并且易于理解。

## Getting Started[¶](https://argo-cd.readthedocs.io/en/stable/#getting-started)

## 入门[¶](https://argo-cd.readthedocs.io/en/stable/#getting-started)

### Quick Start[¶](https://argo-cd.readthedocs.io/en/stable/#quick-start)

### 快速入门[¶](https://argo-cd.readthedocs.io/en/stable/#quick-start)

```
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

Follow our [getting started guide](https://argo-cd.readthedocs.io/en/stable/getting_started/). Further user oriented [documentation](https://argo-cd.readthedocs.io/en/stable/user-guide/) is provided for additional features. If you are looking to upgrade ArgoCD, see the [upgrade guide](https://argo-cd.readthedocs.io/en/stable/operator-manual/upgrading/overview/). Developer oriented [documentation](https://argo-cd.readthedocs.io/en/stable/developer-guide/) is available for people interested in building third-party integrations.

遵循我们的[入门指南](https://argo-cd.readthedocs.io/en/stable/getting_started/)。进一步面向用户的[文档](https://argo-cd.readthedocs.io/en/stable/user-guide/) 提供了附加功能。如果您要升级 ArgoCD，请参阅 [升级指南](https://argo-cd.readthedocs.io/en/stable/operator-manual/upgrading/overview/)。面向开发人员的[文档](https://argo-cd.readthedocs.io/en/stable/developer-guide/) 可供对构建第三方集成感兴趣的人使用。

## How it works[¶](https://argo-cd.readthedocs.io/en/stable/#how-it-works)

## 工作原理[¶](https://argo-cd.readthedocs.io/en/stable/#how-it-works)

Argo CD follows the **GitOps** pattern of using Git repositories as the source of truth for defining the desired application state. Kubernetes manifests can be specified in several ways:

Argo CD 遵循 **GitOps** 模式，使用 Git 存储库作为定义所需应用程序状态的真实来源。 Kubernetes 清单可以通过多种方式指定：

- [kustomize](https://kustomize.io) applications
- [helm](https://helm.sh) charts
- [ksonnet](https://ksonnet.io) applications
- [jsonnet](https://jsonnet.org) files
- Plain directory of YAML/json manifests
- Any custom config management tool configured as a config management plugin

- YAML/json 清单的普通目录
- 配置为配置管理插件的任何自定义配置管理工具

Argo CD automates the deployment of the desired application states in the specified target environments. Application deployments can track updates to branches, tags, or pinned to a specific version of manifests at a Git commit. See [tracking strategies](https://argo-cd.readthedocs.io/en/stable/user-guide/tracking_strategies/) for additional details about the different tracking strategies available.

Argo CD 在指定的目标环境中自动部署所需的应用程序状态。应用程序部署可以在 Git 提交时跟踪对分支、标签或固定到特定版本清单的更新。有关可用的不同跟踪策略的更多详细信息，请参阅 [跟踪策略](https://argo-cd.readthedocs.io/en/stable/user-guide/tracking_strategies/)。

For a quick 10 minute overview of Argo CD, check out the demo presented to the Sig Apps community meeting:

如需 Argo CD 的 10 分钟快速概览，请查看提交给 Sig Apps 社区会议的演示：

[![Argo CD Overview Demo](https://img.youtube.com/vi/aWDIQMbp1cc/0.jpg)](https://youtu.be/aWDIQMbp1cc?t=1m4s)

## Architecture[¶](https://argo-cd.readthedocs.io/en/stable/#architecture)

## 架构[¶](https://argo-cd.readthedocs.io/en/stable/#architecture)

![Argo CD Architecture](https://argo-cd.readthedocs.io/en/stable/assets/argocd_architecture.png)

Argo CD is implemented as a kubernetes controller which continuously monitors running applications and compares the current, live state against the desired target state (as specified in the Git repo). A deployed application whose live state deviates from the target state is considered `OutOfSync`. Argo CD reports & visualizes the differences, while providing facilities to automatically or manually sync the live state back to the desired target state. Any modifications made to the desired target state in the Git repo can be automatically applied and reflected in the specified target environments.

Argo CD 被实现为一个 kubernetes 控制器，它持续监控正在运行的应用程序并将当前的实时状态与所需的目标状态（如 Git 存储库中指定的）进行比较。实时状态偏离目标状态的已部署应用程序被视为“OutOfSync”。 Argo CD 报告和可视化差异，同时提供工具以自动或手动将实时状态同步回所需的目标状态。对 Git 存储库中所需目标状态所做的任何修改都可以自动应用并反映在指定的目标环境中。

For additional details, see [architecture overview](https://argo-cd.readthedocs.io/en/stable/operator-manual/architecture/).

更多详细信息，请参见[架构概述](https://argo-cd.readthedocs.io/en/stable/operator-manual/architecture/)。

## Features[¶](https://argo-cd.readthedocs.io/en/stable/#features)

## 功能[¶](https://argo-cd.readthedocs.io/en/stable/#features)

- Automated deployment of applications to specified target environments
- Support for multiple config management/templating tools (Kustomize, Helm, Ksonnet, Jsonnet, plain-YAML)
- Ability to manage and deploy to multiple clusters
- SSO Integration (OIDC, OAuth2, LDAP, SAML 2.0, GitHub, GitLab, Microsoft, LinkedIn)
- Multi-tenancy and RBAC policies for authorization
- Rollback/Roll-anywhere to any application configuration committed in Git repository
- Health status analysis of application resources
- Automated configuration drift detection and visualization
- Automated or manual syncing of applications to its desired state
- Web UI which provides real-time view of application activity
- CLI for automation and CI integration
- Webhook integration (GitHub, BitBucket, GitLab)
- Access tokens for automation
- PreSync, Sync, PostSync hooks to support complex application rollouts (e.g.blue/green & canary upgrades)
- Audit trails for application events and API calls
- Prometheus metrics 

- 将应用程序自动部署到指定的目标环境
- 支持多种配置管理/模板工具（Kustomize、Helm、Ksonnet、Jsonnet、plain-YAML）
- 能够管理和部署到多个集群
- SSO 集成（OIDC、OAuth2、LDAP、SAML 2.0、GitHub、GitLab、Microsoft、LinkedIn）
- 用于授权的多租户和 RBAC 策略
- 回滚/随处回滚到 Git 存储库中提交的任何应用程序配置
- 应用资源健康状况分析
- 自动配置漂移检测和可视化
- 自动或手动将应用程序同步到所需状态
- 提供应用程序活动实时视图的 Web UI
- 用于自动化和 CI 集成的 CLI
- Webhook 集成（GitHub、BitBucket、GitLab）
- 自动化的访问令牌
- PreSync、Sync、PostSync 挂钩以支持复杂的应用程序部署（例如蓝色/绿色和金丝雀升级）
- 应用程序事件和 API 调用的审计跟踪
- 普罗米修斯指标

- Parameter overrides for overriding ksonnet/helm parameters in Git

- 用于覆盖 Git 中的 ksonnet/helm 参数的参数覆盖

## Development Status[¶](https://argo-cd.readthedocs.io/en/stable/#development-status)

## 开发状态[¶](https://argo-cd.readthedocs.io/en/stable/#development-status)

Argo CD is being actively developed by the community. Our releases can be found [here](https://github.com/argoproj/argo-cd/releases).

社区正在积极开发 Argo CD。我们的版本可以在 [这里](https://github.com/argoproj/argo-cd/releases) 找到。

## Adoption[¶](https://argo-cd.readthedocs.io/en/stable/#adoption)

## 采用[¶](https://argo-cd.readthedocs.io/en/stable/#adoption)

Organizations who have officially adopted Argo CD can be found [here](https://github.com/argoproj/argo-cd/blob/master/USERS.md). 

可以在[此处](https://github.com/argoproj/argo-cd/blob/master/USERS.md) 找到正式采用 Argo CD 的组织。

