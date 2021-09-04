# 5 GitOps Best Practices

# 5 GitOps 最佳实践

[Nov 7, 2019](http://blog.argoproj.io/5-gitops-best-practices-d95cb0cbe9ff?source=post_page-----d95cb0cbe9ff--------------------------------)

GitOps has been on the scene for some time now, but some things trip up users, both new and old. Here are some of the key best practices we’ve discovered while engineering Argo CD and running it at scale managing thousands of apps in production at Intuit.

GitOps 出现已经有一段时间了，但有些东西会绊倒新老用户。以下是我们在设计 Argo CD 并大规模运行它时发现的一些关键最佳实践，在 Intuit 管理生产中的数千个应用程序。

# \#1: Two Repos: One For App Source Code, Another For Manifests

# \#1: 两个 Repos：一个用于应用程序源代码，另一个用于清单

Many engineers start out with both their app source code and their manifests in a single repo. But this has a couple of problems:

- Every commit to your app can result in a deployment.
- Everyone is also able to release, but you might not want everyone to be able to.
- The Git log mix app changes with deployment changes.

许多工程师开始时将他们的应用程序源代码和清单放在一个存储库中。但这有几个问题：

- 对您的应用程序的每次提交都可能导致部署。
- 每个人都可以发布，但您可能不希望每个人都可以发布。
- Git 日志混合应用程序随部署更改而变化。

Instead, we suggest you have two repos and keep your app’s source code in one, and your deployment config (i.e. your manifests) in the other.

相反，我们建议您拥有两个存储库，并将您的应用程序的源代码保存在一个中，而您的部署配置（即您的清单）保存在另一个中。

# \#2: Choose The Right Number Of Deployment Config Repos

# \#2：选择正确数量的部署配置存储库

Consider how many repos you’ll keep your deployment config in. There’s no one-size-fits-all solution, so here are some suggestions:

- A mono-repo for a small company with little automation and where you trust everyone.
- A repo per team for a mid-sized company with some automation.
- A repo per service for a large company with plenty of automation and a need for control.

考虑将部署配置保留在多少个存储库中。没有一刀切的解决方案，因此这里有一些建议：

- 适用于自动化程度低且您信任每个人的小公司的单一存储库。
- 每个团队的一个具有一定自动化程度的中型公司的存储库。
- 为具有大量自动化和控制需求的大公司提供每个服务的存储库。

With their own repo, teams can look after themselves. They can decide who can release, rather than either having a single central team being a release bottleneck or giving every team write access.

有了自己的存储库，团队可以照顾自己。他们可以决定谁可以发布，而不是让单个中央团队成为发布瓶颈或给每个团队写访问权限。

# \#3: Test Your Manifests Before You Commit

# \#3: 在你提交之前测试你的清单

Many engineers commit and push changes to their manifests and then let the GitOps agent validate their changes by seeing if it can deploy the app. We see a lot of issues that escape into pre-production environments that could have been prevented by test the changes before committing and pushing them.

许多工程师提交并将更改推送到他们的清单，然后让 GitOps 代理通过查看它是否可以部署应用程序来验证他们的更改。我们看到很多问题都逃到了预生产环境中，而这些问题本可以通过在提交和推送更改之前测试更改来避免。

The agent is typically running a CLI command (such as “helm template”) to generate your manifests, so you can test your manifests by running the command locally before you even commit your changes.

代理通常运行 CLI 命令（例如“helm 模板”）来生成您的清单，因此您甚至可以在提交更改之前通过在本地运行该命令来测试您的清单。

# \#4: Git Manifests Should Not Change Due To External Changes

# \#4：Git 清单不应因外部更改而更改

Both Kustomize and Helm allow you to template manifests that can be different for the same commit.

Kustomize 和 Helm 都允许你为同一个提交制作不同的清单模板。

If your manifests can change without committing to Git:

- You won’t be able to control or audit changes.
- You probably can’t be able to roll back to a good version.

如果您的清单可以在不提交 Git 的情况下更改：

- 您将无法控制或审核更改。
- 您可能无法回滚到好的版本。

If you’re using Kustomize remote bases, pin them to specific commits.

如果您使用 Kustomize 远程基地，请将它们固定到特定的提交。

```
# Un-pinned:
bases:
- github.com/argoproj/argo-cd/manifests/cluster-install# Pinned:
bases:
- github.com/argoproj/argo-cd//manifests/cluster-install?ref=v0.11.1
```


If you’re using Helm dependencies, pin them too.

如果您正在使用 Helm 依赖项，也请固定它们。

```
# Un-pinned:
dependencies:
- name: argo-cd
  version: *
  repository: github.com/argoproj/argo-cd/manifests/cluster-install# Pinned:
dependencies:
- name: argo-cd
  version: 0.6.1
  repository: github.com/argoproj/argo-cd/manifests/cluster-install
```


# \#5: Plan How You’ll Manage Secrets

# \#5：计划如何管理秘密

It’s no secret that you need to do extra work for secret management in GitOps. Plan how you’ll manage secrets before you start. Here are some options we know people have used successfully:

- [Bitnami Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets)
- [Godaddy Kubernetes External Secrets](https://github.com/godaddy/kubernetes-external-secrets)
- [Hashicorp Vault](https://www.vaultproject.io/)
- [Helm Secrets](https://github.com/futuresimple/helm-secrets)
- [Kustomize secret generator plugins](https://github.com/kubernetes-sigs/kustomize/blob/fd7a353df6cece4629b8e8ad56b71e30636f38fc/examples/kvSourceGoPlugin.md#secret-values-from-anywhere)

您需要为 GitOps 中的秘密管理做额外的工作，这已经不是什么秘密了。在开始之前计划如何管理机密。以下是我们知道人们成功使用的一些选项：

- [Bitnami Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets)
- [Godaddy Kubernetes 外部秘密](https://github.com/godaddy/kubernetes-external-secrets)
- [Hashicorp Vault](https://www.vaultproject.io/)
- [掌舵秘密](https://github.com/futuresimple/helm-secrets)
- [Kustomize 秘密生成器插件](https://github.com/kubernetes-sigs/kustomize/blob/fd7a353df6cece4629b8e8ad56b71e30636f38fc/examples/kvSourceGoPlugin.md#secret-values-from-anywhere)

# Conclusion

#  结论

While there’s rarely just one way to do something, hopefully, these general best practices can point you in the right direction!

虽然很少只有一种方法可以做某事，但希望这些通用的最佳实践可以为您指明正确的方向！

We hope to meet you at [**KubeCon San Diego**](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2019/) and see you at one of the following Argo CD sessions !

我们希望在 [**KubeCon San Diego**](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2019/) 与您会面，并在以下 Argo CD 会议之一见到您！

- [Leveling Up Your CD: Unlocking Progressive Delivery on Kubernetes — Daniel Thomson & Jesse Suen, Intuit](https://kccncna19.sched.com/event/Uaaj?iframe=yes&w=&sidebar=yes&bg=no) 

- [升级您的 CD：在 Kubernetes 上解锁渐进式交付 — Daniel Thomson 和 Jesse Suen，Intuit](https://kccncna19.sched.com/event/Uaaj?iframe=yes&w=&sidebar=yes&bg=no)

- [Tutorial: Everything You Need To Become a GitOps Ninja — Alex Collins & Alexander Matyushentsev, Intuit](https://kccncna19.sched.com/event/Uaee/tutorial-everything-you-need-to-become-a-gitops-ninja-alex-collins-alexander-matyushentsev-intuit?iframe=yes&w=100%&sidebar=yes&bg=no)
- [Panel: GitOps User Stories — Javeria Khan, Palo Alto Networks; Matthias Radestock, Weaveworks; Hubert Chen, Branch; Kyle Rockman, Under Armour; & Jesse Suen, Intuit](https://kccncna19.sched.com/event/UaYh/panel-gitops-user-stories-javeria-khan-palo-alto-networks-matthias-radestock-weaveworks-hubert-chen-branch-kyle-rockman-under-armour-jesse-suen-intuit?iframe=yes&w=100%&sidebar=yes&bg=no) 

- [教程：成为 GitOps 忍者所需的一切 — Alex Collins 和 Alexander Matyushentsev，Intuit](https://kccncna19.sched.com/event/Uaee/tutorial-everything-you-need-to-become-a-gitops-ninja-alex-collins-alexander-matyuushentsev-intuit?iframe=yes&w=100%&sidebar=yes&bg=no)
- [小组：GitOps 用户故事 - Javeria Khan，Palo Alto Networks；马蒂亚斯·拉德斯托克 (Matthias Radestock)，Weaveworks； Hubert Chen，分部；凯尔·洛克曼，安德玛； & Jesse Suen, Intuit](https://kccncna19.sched.com/event/UaYh/panel-gitops-user-stories-javeria-khan-palo-alto-networks-matthias-radestock-weaveworks-hubert-chen-branch-kyle-rockman-under-armour-jesse-suen-intuit?iframe=yes&w=100%&sidebar=yes&bg=no)

