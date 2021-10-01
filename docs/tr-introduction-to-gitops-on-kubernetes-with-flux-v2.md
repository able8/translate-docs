# Introduction to GitOps on Kubernetes with Flux v2

# 在 Kubernetes 上使用 Flux v2 介绍 GitOps

Today we’re having a look at how to set up a GitOps pipeline for your Kubernetes cluster with [Flux v2](https://toolkit.fluxcd.io/).

今天我们来看看如何使用 [Flux v2](https://toolkit.fluxcd.io/) 为您的 Kubernetes 集群设置 GitOps 管道。

We will first go through some core concepts of Flux and then create our first GitOps workflow.

我们将首先了解 Flux 的一些核心概念，然后创建我们的第一个 GitOps 工作流程。

You will need access to a Kubernetes cluster, a shell interface and a Github account to follow this guide. Note that you can use any git provider (Gitlab, Bitbucket, custom) but you’ll have to modify the provided example commands.

您需要访问 Kubernetes 集群、shell 界面和 Github 帐户才能遵循本指南。请注意，您可以使用任何 git 提供程序（Gitlab、Bitbucket、custom），但您必须修改提供的示例命令。

## What is GitOps?

## 什么是 GitOps？

> GitOps is a way of managing your infrastructure and applications so that whole system is described declaratively and version controlled (most likely  in a Git repository), and having an automated process that ensures that  the deployed environment matches the state specified in a repository.
>

> GitOps 是一种管理您的基础设施和应用程序的方法，以便以声明方式描述整个系统并进行版本控制（最有可能在 Git 存储库中），并拥有一个自动化流程，以确保部署的环境与存储库中指定的状态相匹配。
>
> – https://toolkit.fluxcd.io/core-concepts/#gitops

In Kubernetes practice this means that `git` is used over `kubectl` (`helm`, etc) to perform operations tasks against the cluster.

在 Kubernetes 实践中，这意味着在 `kubectl`（`helm` 等）上使用 `git` 来对集群执行操作任务。

Pushing to `master` triggers a deployment to the cluster. We can work with branches and Merge Requests to diff and review changes to the desired cluster state. We can audit past cluster states with `git log`. We can rollback a change using `git revert`.

推送到“master”会触发对集群的部署。我们可以使用分支和合并请求来比较和审查对所需集群状态的更改。我们可以使用`git log`审计过去的集群状态。我们可以使用 `git revert` 回滚更改。

## What is Flux v2?

## 什么是 Flux v2？

Flux v2 is a Toolkit for building GitOps workflows on Kubernetes.

Flux v2 是一个用于在 Kubernetes 上构建 GitOps 工作流的工具包。

In simplified terms, Flux v2 is deployed to a Kubernetes cluster and configured to watch git repositories containing Kubernetes manifests.

简而言之，Flux v2 部署到 Kubernetes 集群并配置为监视包含 Kubernetes 清单的 git 存储库。

![GitOps Toolkit](https://blog.sldk.de/img/posts/fluxcd/gitops-toolkit.png)

As Flux watches all our repos and pulls the changes into the cluster, we can focus on writing our Kubernetes manifests. We don’t have to worry about client-side tooling and pushing the changes from every git repo into the cluster. This is what GitOps is about.

当 Flux 监视我们所有的存储库并将更改拉入集群时，我们可以专注于编写我们的 Kubernetes 清单。我们不必担心客户端工具以及将每个 git 存储库中的更改推送到集群中。这就是 GitOps 的意义所在。

We will come back to the individual Flux components once we’ve installed them in the upcoming steps.

一旦我们在接下来的步骤中安装了它们，我们将回到各个 Flux 组件。

## Installing Flux v2

## 安装 Flux v2

We’ll follow the official docs and use the `flux` CLI tool.

我们将遵循官方文档并使用 `flux` CLI 工具。

You can install it with:

您可以使用以下命令安装它：

```bash
curl -s https://toolkit.fluxcd.io/install.sh |sudo bash

# enable completions in ~/.bash_profile
.<(flux completion bash)
```

The `flux bootstrap` command will ensure that:

`flux bootstrap` 命令将确保：

1. A new GitOps repository for our manifests is created on Github
2. A `flux-system` namespace with all Flux components is configured on our cluster
3. The Flux controllers are set up to sync with our new git repository

1. 在 Github 上为我们的清单创建了一个新的 GitOps 存储库
2. 在我们的集群上配置了一个包含所有 Flux 组件的 `flux-system` 命名空间
3. Flux 控制器设置为与我们新的 git 存储库同步

ℹ️ Note that there’s also [Terraform Provider](https://toolkit.fluxcd.io/guides/installation/#bootstrap-with-terraform) as an alternative installation method.

ℹ️ 请注意，还有 [Terraform Provider](https://toolkit.fluxcd.io/guides/installation/#bootstrap-with-terraform) 作为替代安装方法。

The following snippet shows an annotated bootstrap command. Feel free to play around with `flux bootstrap --help` to get to know all available options.

以下代码段显示了带注释的引导程序命令。随意使用 `flux bootstrap --help` 来了解所有可用选项。

```bash
export GITHUB_TOKEN=<your-gitlab-token-here> # 1

flux bootstrap \
  github \                      # 2
  --owner <your-github-user> \  # 3
  --repository cluster-name \   # 4
  --personal \                  # 5
  --private \                   # 6
  --path cluster \              # 7
  --branch master               # 8
```

1. We’re going to need to give `flux` access to our Github account so that it can manage our GitOps repository on our behalf. You can generate a new Personal Access Token in [Github Settings](https://github.com/settings/tokens).
2. We need to tell Flux that we’re using Github as git provider. You can also use Gitlab and generic git servers.
3. The Github username.
4. The name of our git repository.
5. (Optional) We’re creating a personal Github repo.
6. (Optional) We’re starting out with a private Github repo.
7. (Optional) A path inside the repository to be watched by Flux.
8. (Optional) The git branch to be watched by Flux.

1. 我们需要授予 `flux` 访问我们的 Github 帐户的权限，以便它可以代表我们管理我们的 GitOps 存储库。您可以在 [Github 设置](https://github.com/settings/tokens) 中生成新的个人访问令牌。
2. 我们需要告诉 Flux 我们使用 Github 作为 git provider。您还可以使用 Gitlab 和通用 git 服务器。
3. Github 用户名。
4. 我们的 git 存储库的名称。
5.（可选）我们正在创建一个个人 Github 存储库。
6.（可选）我们从一个私有的 Github 仓库开始。
7.（可选）Flux 监视的存储库内的路径。
8.（可选）Flux 监视的 git 分支。

The bootstrap command is interactive and will let you know about the progress and state of the installation.

bootstrap 命令是交互式的，它将让您了解安装的进度和状态。

This is all you need to get up and running with Flux. Note that the boostrap command is idempotent and can also be used with an existing GitOps repository.

这就是您启动和运行 Flux 所需的全部内容。请注意，boostrap 命令是幂等的，也可以与现有的 GitOps 存储库一起使用。

## Inspecting the Flux components 🕵

## 检查 Flux 组件 🕵

### The flux-system namespace

### 通量系统命名空间

Let’s take a look at what was deployed to the k8s cluster for us.

下面我们来看看部署到k8s集群上的内容。

```bash
# Let's have a look at the Flux controllers...
kubectl get pods -n flux-system
NAME                                      READY   STATUS    RESTARTS   AGE
helm-controller-c858b9c65-v8pgr           1/1     Running   28         17d
kustomize-controller-6767b8fd78-8ptd4     1/1     Running   30         17d
notification-controller-db467b4fb-pt9tc   1/1     Running   29         17d
source-controller-7b4d748b45-2blvd        1/1     Running   28         17d
# And the custom resource definitions...
kubectl get crd -o name |grep flux
customresourcedefinition.apiextensions.k8s.io/alerts.notification.toolkit.fluxcd.io
customresourcedefinition.apiextensions.k8s.io/buckets.source.toolkit.fluxcd.io
customresourcedefinition.apiextensions.k8s.io/gitrepositories.source.toolkit.fluxcd.io
customresourcedefinition.apiextensions.k8s.io/helmcharts.source.toolkit.fluxcd.io
customresourcedefinition.apiextensions.k8s.io/helmreleases.helm.toolkit.fluxcd.io
customresourcedefinition.apiextensions.k8s.io/helmrepositories.source.toolkit.fluxcd.io
customresourcedefinition.apiextensions.k8s.io/kustomizations.kustomize.toolkit.fluxcd.io
customresourcedefinition.apiextensions.k8s.io/providers.notification.toolkit.fluxcd.io
customresourcedefinition.apiextensions.k8s.io/receivers.notification.toolkit.fluxcd.io
```

Great! The Flux controllers are running and we have a bunch of new CustomResourceDefinitions available to us. Read ahead to learn more about how those play together.

伟大的！ Flux 控制器正在运行，我们有一堆新的 CustomResourceDefinitions 可供我们使用。提前阅读以了解更多关于这些如何一起玩的信息。

### The GitOps repository

### GitOps 存储库

To better understand how Flux works, let’s now take a look at the directory structure of our bootstrapped GitOps repository on Github. We’re going to see that the bootstrap command has created some Kubernetes yamls for us.

为了更好地理解 Flux 的工作原理，现在让我们来看看我们在 Github 上引导的 GitOps 存储库的目录结构。我们将看到 bootstrap 命令为我们创建了一些 Kubernetes yaml。

```bash
git clone git@github.com:${username}/cluster-name.git
cd cluster-name
tree
.
├── cluster
│   └── flux-system
│       ├── gotk-components.yaml
│       ├── gotk-sync.yaml
│       └── kustomization.yaml
└── README.md

2 directories, 4 files
```

Cool. We’ll go through the individual files.

凉爽的。我们将浏览各个文件。

#### The cluster directory

#### 集群目录

This directory was created because we used `flux bootstrap ... --path cluster ...` during bootstrapping. We’ll see in a second that this folder is also special because it’s being watched by the Flux `kustomization-controller`.

创建此目录是因为我们在引导过程中使用了 `flux bootstrap ... --path cluster ...`。我们稍后会看到这个文件夹也很特别，因为它被 Flux `kustomization-controller` 监视。

Any Kubernetes yaml that we throw in this directory will be deployed to our cluster. It also means that we can have different files and directories in our repository that will never make it to the cluster. This is useful because we could have one single git repository for our infrastructure declarations and divide it into sub-paths for different k8s clusters or environments. For example, we could bootstrap multiple Flux environments and use paths like `dev-cluster`, `stage-cluster`, `prod-cluster`.

我们在此目录中放入的任何 Kubernetes yaml 都将部署到我们的集群中。这也意味着我们的存储库中可以有不同的文件和目录，这些文件和目录永远不会进入集群。这很有用，因为我们可以为我们的基础设施声明使用一个单独的 git 存储库，并将其划分为不同 k8s 集群或环境的子路径。例如，我们可以引导多个 Flux 环境并使用像 `dev-cluster`、`stage-cluster`、`prod-cluster` 这样的路径。

That’s just one way to organize. It’s nice that we have this flexibility.

这只是一种组织方式。很高兴我们有这种灵活性。

#### The flux-system directory

#### 通量系统目录

This directory represents the `flux-system` k8s namespace and contains Kubernetes declarations.

该目录代表 `flux-system` k8s 命名空间并包含 Kubernetes 声明。

I like to continue this pattern and create a subdirectory for each namespace that I deploy with Flux.

我喜欢继续这种模式，并为我使用 Flux 部署的每个命名空间创建一个子目录。

Example:

例子：

```
└── cluster
    ├── flux-system
    ├── monitoring
    ├── cool-app-namespace
    └── monitoring
```

This isn’t required. You’re flexible and can structure the `/cluster` directory like you want!

这不是必需的。您很灵活，可以根据需要构建 `/cluster` 目录！

#### kustomization.yaml

As we can see Flux created a simple `kustomization.yaml` for us. If you’re not familiar with [Kustomize](https://kustomize.io/) you can just ignore this for now. All it does is include the other two files in the `cluster/flux-system` directory.

正如我们所见，Flux 为我们创建了一个简单的 `kustomization.yaml`。如果您不熟悉 [Kustomize](https://kustomize.io/)，您可以暂时忽略它。它所做的只是将其他两个文件包含在 `cluster/flux-system` 目录中。

```yaml
# cluster/flux-system/kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
- gotk-components.yaml
- gotk-sync.yaml
```

#### gotk-components.yaml

This file contains the **G**it**O**ps **T**ool**K**it components 🤯 - in here you will find the *Deployments* and *Services* for the Flux controllers and all the *CustomResourceDefinitions* that we've already seen when we inspected the `flux-system` namespace.

这个文件包含 **G**it**O**ps **T**ool**K**it 组件🤯 - 在这里你会找到 Flux 控制器的 *Deployments* 和 *Services* 以及所有我们在检查 `flux-system` 命名空间时已经看到的 *CustomResourceDefinitions*。

This file is generated by the `flux bootstrap` command and you shouldn’t have to deal with it manually.

该文件由 `flux bootstrap` 命令生成，您不必手动处理它。

#### gotk-sync.yaml

The file `gotk-sync.yaml` is more interesting to us. This file declares our first two custom resources for Flux.

文件 `gotk-sync.yaml` 对我们来说更有趣。这个文件声明了 Flux 的前两个自定义资源。

```yaml
# cluster/flux-system/gotk-sync.yaml
---
apiVersion: source.toolkit.fluxcd.io/v1beta1
kind: GitRepository
metadata:
  name: flux-system
  namespace: flux-system
spec:
  interval: 1m0s
  ref:
    branch: master
  secretRef:
    name: flux-system
  url: ssh://git@github.com/your-user-name-here/cluster-name
---
apiVersion: kustomize.toolkit.fluxcd.io/v1beta1
kind: Kustomization
metadata:
  name: flux-system
  namespace: flux-system
spec:
  interval: 10m0s
  path: ./cluster
  prune: true
  sourceRef:
    kind: GitRepository
    name: flux-system
  validation: client
```

Let’s take a closer look at these two…

让我们仔细看看这两个……

#### The flux-system GitRepository

#### 通量系统 GitRepository

The first custom resource is of type *GitRepository*. A GitRepository is one of the possible *Sources* that are picked up by the Flux `source-controller`. Once a repository is registered with this controller, the Flux toolkit is able to watch it for changes and notify other Flux controllers to apply the contents to the cluster.

第一个自定义资源的类型为 *GitRepository*。 GitRepository 是 Flux `source-controller` 选取的可能的 *Sources* 之一。一旦存储库注册到这个控制器，Flux 工具包就能够观察它的变化并通知其他 Flux 控制器将内容应用到集群。

This specific GitRepository definition tells Flux to check the `master` branch of our new GitOps repository from Github. every `1m0s` using a `flux-system` secret for authentication (this secret contains your Github Personal Access Token that we created during bootstrapping). 

这个特定的 GitRepository 定义告诉 Flux 从 Github 检查我们新的 GitOps 存储库的 master 分支。每 1m0s 使用一个 `flux-system` 秘密进行身份验证（这个秘密包含我们在引导过程中创建的 Github 个人访问令牌）。

Note that we could create more GitRepository resources and effectively sync any number of GitOps repositories with our Flux-enabled cluster.

请注意，我们可以创建更多 GitRepository 资源，并有效地将任意数量的 GitOps 存储库与启用 Flux 的集群同步。

Example use-case: You have multiple teams that want to deploy to one cluster. You want each team to maintain their own infrastructure repository. You could set up Flux on the cluster such that it pulls each team’s repository.

用例示例：您有多个团队想要部署到一个集群。您希望每个团队维护自己的基础架构存储库。您可以在集群上设置 Flux，以便它拉取每个团队的存储库。

Read more in the [Flux Docs](https://toolkit.fluxcd.io/components/source/gitrepositories/).

在 [Flux Docs](https://toolkit.fluxcd.io/components/source/gitrepositories/) 中阅读更多内容。

#### The flux-system Kustomization

#### 通量系统 Kustomization

The second custom resource in this file is a (Flux-) **Kustomization**.

此文件中的第二个自定义资源是 (Flux-) **Kustomization**。

⚠️ Not to be confused with a [**Kustomize** resource](https://kustomize.io/).

⚠️ 不要与 [**Kustomize** 资源](https://kustomize.io/) 混淆。

A Flux Kustomization describes a *path* inside of a *Source* that Flux should apply to the cluster. It’s assumed that there are Kubernetes yamls and/or Kustomize yamls located in the directory at the named path.

Flux Kustomization 描述了 Flux 应应用于集群的 *Source* 内的 *path*。假设在命名路径的目录中存在 Kubernetes yaml 和/或 Kustomize yaml。

The Flux `kustomize-controller` is going to periodically apply the yamls at the given location to our cluster.

Flux `kustomize-controller` 将定期将给定位置的 yaml 应用到我们的集群。

Here’s a simplified walk-through of how the `kustomize-controller` is going to install our k8s resources.

下面是关于 `kustomize-controller` 将如何安装我们的 k8s 资源的简化演练。

- Checkout the configured branch from

   - 检出配置的分支

  ```
   ssh://git@github.com/${username}/cluster-name
  ```

- Change the working directory to the specified path: `./cluster`

- 将工作目录更改为指定路径：`./cluster`

- Build a root

   - 建立一个根

  ```
   Kustomize
  ```

   

yaml at this location if no explicit one is available:

   如果没有明确的可用，请在此位置使用 yaml：

  ```
   kustomize create --autodetect --recursive > kustomization.yaml
  ```

- Apply the result:

   - 应用结果：

  ```
   kubectl apply -k kustomization.yaml`
  ```

I think it’s good to know how `kustomize` is being used to build the final Kubernetes manifest from our GitOps repo. The commands above can be used for testing locally. This is especially helpful when you mix plain Kubernetes yamls with Kustomize yamls as this has an impact on the recursive auto-detection.

我认为很高兴知道如何使用 `kustomize` 从我们的 GitOps 存储库构建最终的 Kubernetes 清单。以上命令可用于本地测试。当您将普通 Kubernetes yaml 与 Kustomize yaml 混合使用时，这尤其有用，因为这会影响递归自动检测。

## Deploying own resources

## 部署自己的资源

Let’s do a really simple example on how to deploy stuff to our cluster.

让我们举一个非常简单的例子来说明如何将东西部署到我们的集群中。

All we want to do is create a new namespace “awesome-namespace”.

我们要做的就是创建一个新的命名空间“awesome-namespace”。

```bash
# Create a new directory for our namespace
mkdir cluster/awesome-namespace

# Create the Kubernetes namespace descriptor
cat <<EOF > cluster/awesome-namespace/namespace.yaml
---
apiVersion: v1
kind: Namespace
metadata:
  name: awesome-namespace
EOF

# Commit and push the change
git add .
git commit -m "Add 'awesome-namespace'"
git push
```

Lean back and watch as the new namespace is shipped to our cluster 🛳️

向后靠，看着新的命名空间被传送到我们的集群 🛳️

------

## Summary

##  概括

We installed the Flux GitOps toolkit to a Kubernetes cluster and deployed a Kubernetes manifest by only using `git`. Along the way we learnt about the Flux components.

我们将 Flux GitOps 工具包安装到 Kubernetes 集群，并仅使用 `git` 部署了 Kubernetes 清单。在此过程中，我们了解了 Flux 组件。

I hope this gave you an introduction on what Flux v2 is, how it works and how you can use it to enhance your CI/CD workflows.

我希望这能让您了解 Flux v2 是什么、它是如何工作的以及如何使用它来增强您的 CI/CD 工作流程。

I’m planning to write more about this topic in the future. We'll take a look at how to integrate [SOPS](https://github.com/mozilla/sops) for secret handling, how to deploy [Helm Charts](https://toolkit.fluxcd.io/components/source/helmcharts/) and how to set up Flux notifications.

我计划在未来写更多关于这个主题的文章。我们来看看如何集成 [SOPS](https://github.com/mozilla/sops) 进行秘密处理，如何部署 [Helm Charts](https://toolkit.fluxcd.io/components/source/helmcharts/) 以及如何设置 Flux 通知。

If you read this far, hit me up in the comments, on [Twitter](https://twitter.com/sladkovik) or [Instagram](https://www.instagram.com/sladkoff2/) to let me know how I did.

如果你读到这里，请在评论中联系我，在 [Twitter](https://twitter.com/sladkovik) 或 [Instagram](https://www.instagram.com/sladkoff2/) 上让我知道我是怎么做的。

Thanks 🙏

谢谢🙏

------

