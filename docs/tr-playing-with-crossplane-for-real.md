# Playing with Crossplane, for real

# 真实地玩 Crossplane

From https://prune998.medium.com/playing-with-crossplane-for-real-f591e66065ae

Aug 25, 2022



![img](https://miro.medium.com/max/1400/1*c8jhmTqn9FGm7vFyZywWmA.png)

It’s 2022 ! We all know about the necessity of creating your Infra AS CODE. I guess we can all agree that [Terraform (TF)](https://www.hashicorp.com/products/terraform) is the leader in this field.

现在是 2022 年！我们都知道创建 Infra AS CODE 的必要性。我想我们都同意 [Terraform (TF)](https://www.hashicorp.com/products/terraform) 是该领域的领导者。

that said, TF is far from easy to use. This is actually not the fault of TF itself (or its parent company, [Hashicorp](https://www.hashicorp.com/)), but to the fact that each provider is so specific that it's impossible to build up something clean out of it.

也就是说，TF 远非易用。这实际上不是 TF 本身（或其母公司 [Hashicorp](https://www.hashicorp.com/))的错，而是因为每个提供商都非常具体，因此不可能建立干净的东西从它出来。

So we end up with complicated code, tricks to make things happen or not (like using `count`) and the two biggest pain point:

所以我们最终得到了复杂的代码、使事情发生或不发生的技巧（比如使用 count ）和两个最大的痛点：

- having to manage a state file which holds the result of the last apply
- globally not possible to re-use code or build simple abstraction to create a  resource in different clouds, like a K8s cluster in AWS and Google.

- 必须管理一个保存上次应用结果的状态文件
- 在全球范围内不可能重用代码或构建简单的抽象来在不同的云中创建资源，例如 AWS 和 Google 中的 K8s 集群。

The global consensus is that TF is too complicated for Devs team to use and maintain well. It’s down to the Ops to operate it, or at least, try to  provide `modules` so Devs can work with what is supposed to be an abstraction.

全球共识是 TF 过于复杂，开发团队无法很好地使用和维护。操作它取决于 Ops，或者至少尝试提供“模块”，以便 Devs 可以使用应该是抽象的东西。

## Then enters [Crossplane](https://crossplane.github.io/) !

## 然后进入[Crossplane](https://crossplane.github.io/) ！

Before diving into Crossplane, let’s put some notice, or as we like to call it:

在深入研究 Crossplane 之前，让我们先注意一下，或者我们喜欢这样称呼它：

# TL ; DR

- I’m no Crossplane expert and only spent few days playing with it for a  Proof Of Concept. Please, help me better understand/use it if you think  i’m wrong
- OSS Crossplane Providers are really limited
- Jet Providers are on par with Terraform providers, but lack docs and may be buggy and are not well supported or updated by the community
- Still unsure how to replicate 100% of what you can do with TF

- 我不是 Crossplane 专家，只花了几天时间玩弄它来进行概念验证。如果您认为我错了，请帮助我更好地理解/使用它
- OSS Crossplane 提供商真的很有限
- Jet Providers 与 Terraform providers 相当，但缺乏文档并且可能存在错误并且没有得到社区的良好支持或更新
- 仍然不确定如何复制 100% 的 TF 功能

**Overall conclusion is that it’s not production ready for me**

**总体结论是它还没有为我准备好生产**

Please follow up for details 详情请关注

# Crossplane

As stated on their website: **The cloud native control plane framework**

正如他们网站上所述：**云原生控制平面框架**

Saying that Crossplane is a K8s native way to do IaC would be limitative. Crossplane is far more than that. Crossplane, to me, is a way to build  simple abstractions in front of complex stuff, like infrastructure or  deployments.

说 Crossplane 是 K8s 原生的 IaC 方式是有局限性的。 Crossplane 远不止于此。对我来说，Crossplane 是一种在复杂事物（如基础设施或部署）之前构建简单抽象的方法。

As TF, Crossplane is build up on [**Providers**](https://github.com/crossplane-contrib) ([see official list here](https://crossplane.github.io/docs/v1.9/concepts/providers.html)), which extends Crossplane with new *Managed Resource Type* to work with.

作为 TF，Crossplane 建立在 [**Providers**](https://github.com/crossplane-contrib) 上（[请在此处查看官方列表](https://crossplane.github.io/docs/v1.9/concepts/providers.html))，它使用新的*托管资源类型*扩展了 Crossplane。

But the strength of Crossplane lies in its `Composition` feature. It is exactly what the name implies: you build up a new type  of resource by combining (composing) other type of resources. Ex: create a K8s cluster and a nodepool, or create a DB Instance, a DB schema and a DB user, all at one.

但 Crossplane 的优势在于其“合成”功能。顾名思义：您通过组合（组合）其他类型的资源来构建一种新型资源。例如：创建一个 K8s 集群和一个节点池，或者同时创建一个数据库实例、一个数据库模式和一个数据库用户。

Let’s grab some pictures from the official docs:

让我们从官方文档中抓取一些图片：

![https://crossplane.github.io/docs/v1.9/concepts/composition.html#overview](https://miro.medium.com/max/1400/1*RotpDcpSi0UuYq28IdD35g.png)

You `claim` a Postresql Instance, that references a `Composite Resource` that will trigger the creation of a `CloudSQL Instance` and a `Firewall Rule` to access it. Neat !

您“声明”一个 Postresql 实例，它引用一个“复合资源”，该资源将触发“CloudSQL 实例”的创建和“防火墙规则”以访问它。整洁的 ！

It’s a little bit more complicated, so here’s another picture from the doc that is supposed to be closer to reality:

它有点复杂，所以这是文档中的另一张图片，应该更接近现实：

![https://crossplane.github.io/docs/v1.9/concepts/composition.html#how-it-works](https://miro.medium.com/max/1400/1*8fN6cZ0OX_xI--8UmtRtIA.png)

AH ! There we go !

啊！我们开始了！

You can go read the docs at [https://crossplane.github.io/docs/v1.9/concepts/composition.html](https://crossplane.github.io/docs/v1.9/concepts/composition.html#overview). I personally read that again and again, and wasn’t able to fully  understand the real thing until I played with it, and built my own  schema:

您可以在 [https://crossplane.github.io/docs/v1.9/concepts/composition.html](https://crossplane.github.io/docs/v1.9/concepts/composition .html# 概述)。我个人一遍又一遍地阅读它，直到我玩它并构建我自己的模式之前，我才能够完全理解真实的东西：

![img](https://miro.medium.com/max/1400/1*S-r3Y4mhb6W9TyNUHQbnFQ.png)

Let me break things down:

让我分解一下：

1. **Infra**: Deploy a Provider, like GKE, AWS, Helm, K8s, even a [Terraform provider](https://github.com/crossplane-contrib/provider-terraform)
2. The provider created a set of CRDs corresponding to each Cloud resource it manages (yellow boxes)
3. **Infra**: Create a CompositeResourceDefinition (XRD) which creates an interface with a limited set of parameters to tweak 

4. Crossplane will create and maintain two new CRDs based on the XRD: a Claim and a  CompositeResource (XR) (green boxes). Crossplane will start watching and reconciling CR based on those CRDs
5. **Infra**: Create a Composition, which will reference a source XRD and a list of  Resources to created (from the CRDs created by the Provider). It’s a  sort of templating resources with values from the interface (the XRD)
6. **Dev**: Claim a resource (purple box) -> a Claim is actually a CustomResource of a type maintained by Crossplane
7. Crossplane will create a CompositeResource (XR) based on the Claim
8. Crossplane will create CustomResources (CR) which are instances of the Provider’s  resources, based on the content of the CompositeResource (XR) (red  boxes)
9. Provider will reconcile the resources he manages, and call GCP API (in case of  the GCP provider) to create the resources declared in the CR



1. **Infra**：部署供应商，如 GKE、AWS、Helm、K8s，甚至 [Terraform 供应商](https://github.com/crossplane-contrib/provider-terraform)
2. 提供商创建了一组 CRD，对应于它管理的每个 Cloud 资源（黄色框）
3. **Infra**：创建一个 CompositeResourceDefinition (XRD)，它创建一个界面，其中包含一组有限的参数以进行调整
4. Crossplane 将基于 XRD 创建和维护两个新的 CRD：Claim 和 CompositeResource (XR)（绿色框）。 Crossplane 将根据这些 CRD 开始观察和协调 CR

1. **Infra**：创建一个组合，它将引用源 XRD 和要创建的资源列表（来自提供商创建的 CRD）。它是一种模板资源，具有来自界面（XRD）的值

2. **Dev**: Claim a resource (purple box) -> 一个Claim实际上是Crossplane维护的一个类型的CustomResource

3. Crossplane会根据Claim创建一个CompositeResource(XR)

4. Crossplane 将根据 CompositeResource (XR)（红框）的内容创建 CustomResources (CR)，它是 Provider 资源的实例

5. Provider 将协调他管理的资源，并调用 GCP API（如果是 GCP provider）创建 CR 中声明的资源

   

This is really powerful, and the only limitation is actually in what a provider can do.

这真的很强大，唯一的限制实际上是提供者可以做什么。

Talking of which, I guess you see me coming, it’s also the biggest problem  Crossplane has: it all depends on what a Provider can do !

说到这里，我猜你看到我来了，这也是 Crossplane 最大的问题：这完全取决于 Provider 能做什么！

Oh, by the way, I'll be a speaker at [KubeCon North America 2022 in Detroit](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/) , please [check my talk here]( https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/program/schedule/), register for Thursday, October 27 • 2:30pm — 4:00pm:

[Tutorial:  Set Up Your Shell For Kubernetes Productivity And Be Efficient Quickly —  Sebastien “Prune” Thomas, Wunderkind & Archy Ayrat Khayretdinov,  Google](https://kccncna2022.sched.com/event/182F7/tutorial-set-up-your-shell-for-kubernetes-productivity-and-be-efficient-quickly-sebastien-prune-thomas-wunderkind-archy-ayrat-khayretdinov-google)

[教程：为 Kubernetes 生产力设置 Shell 并快速提高效率——Sebastien “Prune” Thomas、Wunderkind 和 Archy Ayrat Khayretdinov，谷歌](https://kccncna2022.sched.com/event/182F7/tutorial-set-up-你的shell-for-kubernetes-productivity-and-be-efficient-quickly-sebastien-prune-thomas-wunderkind-archy-ayrat-khayretdinov-google)

# Question 1: Providers

# 问题 1：供应商

As stated, `Providers` are the part that manages some resources. In fact, it’s a `Pod` that will be deployed along Crossplane, will create CRD for the  resource it manages and watch on them. Each instance of a managed  resource will make the `Provider` create and maintain a resource. For the GCP provider, that means calling the GCP API and create some resources.

如前所述，`Providers` 是管理某些资源的部分。事实上，它是一个将沿 Crossplane 部署的“Pod”，将为它管理的资源创建 CRD 并监视它们。托管资源的每个实例都会让“提供者”创建和维护资源。对于 GCP 提供者，这意味着调用 GCP API 并创建一些资源。

I’m not diving into `Providers`configuration here. It involves installing the provider, giving it a global config  and some more specific variations of the config, like having a `Dev` , `Staging` and `Prod`provider for GCP. I would maybe have to create a mix of configurations for each of my `env` * `project` so I can ensure some devs can deploy to `Project A` in `Dev` but not on `Prod` or in `Project B` . This is something I haven’t explored yet… Please, again, comment.

我不会在这里深入探讨 `Providers` 配置。它涉及安装提供程序，为其提供全局配置和配置的一些更具体的变体，例如为 GCP 提供“Dev”、“Staging”和“Prod”提供程序。我可能必须为我的每个 `env` * `project` 创建混合配置，这样我才能确保一些开发人员可以在 `Dev` 中部署到 `Project A` 而不是在 `Prod` 或 `Project B` 中.这是我还没有探索过的东西……请再次发表评论。

You can usually check for Provider’s CRD (supported Resources) at https://doc.crds.dev/.
Ex:
\- crossplane own CRDs: https://doc.crds.dev/github.com/crossplane/crossplane@v1.9.0
\- crossplane GCP Provider (official): https://doc.crds.dev/github.com/crossplane/provider-gcp@v0.21.0

您通常可以在 https://doc.crds.dev/ 上查看提供商的 CRD（支持的资源）。
前任：
\- crossplane 自己的 CRD：https://doc.crds.dev/github.com/crossplane/crossplane@v1.9.0
\- 交叉平面 GCP 提供商（官方）：https://doc.crds.dev/github.com/crossplane/provider-gcp@v0.21.0

Let’s dive in the GCP provider:

让我们深入了解 GCP 提供程序：

![img](https://miro.medium.com/max/1400/1*0e89L90VOy-KlB0WYNlRpw.png)

28 CRDs discovered ? wait, what ? only 28 different GCP resources are managed by Crossplane ?

发现了 28 个 CRD ?等等，什么？只有 28 种不同的 GCP 资源由 Crossplane 管理？

short answer: YES :(

简短的回答：是的:(

So, for example, you can create a `CloudSQLInstance` which represents a Postgres or MySQL Database Instance:

因此，例如，您可以创建一个代表 Postgres 或 MySQL 数据库实例的“CloudSQLInstance”：

![img](https://miro.medium.com/max/1400/1*iJ-uhzAQA4ycAjJs_AG7jw.png)

But then, you just can’t create any specific DB, User or anything else  related to this DB Instance. Because, well, the resources were not  ported into the Provider.

但是，您不能创建任何特定的数据库、用户或与此数据库实例相关的任何其他内容。因为，好吧，资源没有移植到提供者中。

AWS official Provider is a little better, with 172 resources.

AWS官方Provider稍微好一点，有172个资源。

This situation is just normal. Building up a provider is tedious, there are a lot of different resources, and Crossplane team is quite small compared to this huge work. 

这种情况很正常。建立一个提供者是繁琐的，有很多不同的资源，与这项庞大的工作相比，Crossplane 团队很小。

At first, I guess the Cloud providers wanted to  help, but it feels that in the end, they prefered build their own  version of Crossplane for themselves. That is when Google created it [Config Connector](https://cloud.google.com/config-connector/docs/overview) app, which in term is a lightweight version of Crossplane, or [AWS ACK](https://aws.amazon.com/blogs/containers/aws-controllers-for-kubernetes-ack/).

一开始，我猜云提供商是想帮忙，但我觉得最终他们更愿意为自己构建自己的 Crossplane 版本。那是谷歌创建它的时候 [Config Connector](https://cloud.google.com/config-connector/docs/overview) 应用程序，它是 Crossplane 的轻量级版本，或 [AWS ACK](https://aws.amazon.com/blogs/containers/aws-controllers-for-kubernetes-ack/)。

But Crossplane team is smart, so they asked `who else in the world is already maintaining Providers ?` and the answer was, as often… Terraform !

但是 Crossplane 团队很聪明，所以他们问“世界上还有谁已经在维护 Providers 了？”而答案通常是……Terraform！

And the [TerraJet](https://github.com/crossplane/terrajet) project was born !

[TerraJet](https://github.com/crossplane/terrajet) 项目诞生了！

TerraJet (Jet) is a way to convert Terraform Providers into Crossplane  Providers, hiding the TF mechanics. With Jet, TF is run in the  background and the TF State file is, let’s say, split and stored along  each Crossplane Resource.

TerraJet (Jet) 是一种将 Terraform Provider 转换为 Crossplane Provider 的方法，隐藏了 TF 机制。使用 Jet，TF 在后台运行，并且 TF 状态文件被拆分并存储在每个 Crossplane 资源中。

## *Smart ?*

##  *聪明的 ？*

Well, the GCP Jet Provider counts 438 resources, including all that we need  to manage SQL DBs in GCP. AWS Jet Provider counts a wooping 780  resources ! I guess it’s more than you’ll ever use.

好吧，GCP Jet Provider 有 438 个资源，包括我们在 GCP 中管理 SQL 数据库所需的所有资源。 AWS Jet Provider 拥有惊人的 780 种资源！我想这比你用过的还要多。

![img](https://miro.medium.com/max/1400/1*Arl6hA77k5MWXpXImMayjg.png)

# Question 2: Docs ?

# 问题 2：文档？

It seems to be a lot of resources to use ! While TF docs are, well, not  that bad in the end, and you have a LOT of blog posts and examples to  play with, it’s not the same story for Crossplane.

好像要用很多资源！虽然 TF 文档最终并没有那么糟糕，而且您有很多博客文章和示例可以使用，但对于 Crossplane 来说就不一样了。

The [Composition Docs](https://crossplane.github.io/docs/v1.9/concepts/composition.html), for example, showcase few different stuffs that will work out of the  box if you copy/paste. But then, you want your own stuff, maybe not a DB instance. So you start digging in… and reach the [Composition Reference doc](https://crossplane.github.io/docs/v1.9/reference/composition.html).
Once again, you’re driven through the same example, with little more detail  that will just confuse you even more (at least it confused me):

例如，[Composition Docs](https://crossplane.github.io/docs/v1.9/concepts/composition.html) 展示了一些不同的内容，如果您复制/粘贴，这些内容将立即可用。但是，您想要自己的东西，也许不是数据库实例。因此，您开始深入挖掘……并找到 [Composition Reference doc](https://crossplane.github.io/docs/v1.9/reference/composition.html)。
再一次，你被驱动通过相同的例子，更多的细节只会让你更加困惑（至少让我感到困惑）：

- Should I create Claims or XR ?
- what is this `compositionRef` about ?
- where should I set `writeConnectionSecretToRef` ? in the Claim ? The XR ? the resource in the Composition ?
- What are the `secret values` that the provider is returning ?

- 我应该创建 Claims 还是 XR？
- 这个 `compositionRef` 是关于什么的？
- 我应该在哪里设置`writeConnectionSecretToRef`？在索赔中？ XR ？ Composition 中的资源？
- 提供者返回的“秘密值”是什么？

Well, I guess the product is new, and not widely used yet, but man, EVERY  doc, blog, talk, demo is using the same stuff. They all create a single  bucket, [a single DB instance, a VPC with few subnets](https://crossplane.github.io/docs/v1.9/getting-started/create-configuration.html#create-compositions) , and in a rare case, [some K8s clusters](https://github.com/upbound/platform-ref-multi-k8s).

好吧，我想这个产品是新的，还没有被广泛使用，但是伙计，每个文档、博客、谈话、演示都在使用相同的东西。他们都创建一个桶，[一个数据库实例，一个带有几个子网的 VPC](https://crossplane.github.io/docs/v1.9/getting-started/create-configuration.html#create-compositions) ，在极少数情况下，[一些 K8s 集群](https://github.com/upbound/platform-ref-multi-k8s)。

To me, nothing that related to my needs.

对我来说，没有什么与我的需要有关。

And because what I want is not (yet?) supported in the official providers, I have to use the GCP Jet Provider.

而且因为我想要的东西（还没有？）在官方提供者中得到支持，所以我必须使用 GCP Jet Provider。

# Jet Providers, straight from Terraform

# Jet Providers，直接来自 Terraform

Don't expect to find docs on the Jet Providers… because they are transpositions from TF providers, the doc is minimal, usually [a copy of the TF doc itself](https://registry.terraform.io/providers/hashicorp/google/latest/docs), if any.

不要指望在 Jet Providers 上找到文档……因为它们是 TF 提供程序的转换，文档很少，通常是 [TF 文档本身的副本](https://registry.terraform.io/providers/hashicorp/谷歌/最新/文档)，如果有的话。

For my POC, I wanted something simple, and luckily, something that looks  like what most of the examples are based on: a Postgres DB (CloudSQL). Of course, I want the Instance, plus a DB, plus few Users, hopefully  using GCP WorkloadIdentity (IAM ServiceAccount binding).

对于我的 POC，我想要一些简单的东西，幸运的是，一些看起来像大多数示例所基于的东西：Postgres DB (CloudSQL)。当然，我想要实例，加上一个数据库，再加上几个用户，希望使用 GCP WorkloadIdentity（IAM ServiceAccount 绑定）。

Note that this is a really limited need so far. Only 3 resourceTypes are at  play here. The real use-case would have been creating a Pub/Sub, a  ComposerV2, some cloud-functions, and all the IAM stuff that allow one  to trigger the other. All that at once, by Claiming ONE `Composition` .

请注意，到目前为止，这是一个非常有限的需求。这里只有 3 种资源类型在起作用。真正的用例应该是创建一个 Pub/Sub、一个 ComposerV2、一些云函数，以及所有允许一个触发另一个的 IAM 东西。通过 Claim ONE `Composition` 一次完成所有这些。

## XRD

I started creating a XRD:

我开始创建 XRD：

```
apiVersion: apiextensions.crossplane.io/v1
kind: CompositeResourceDefinition
metadata:
  name: xjetpostgresqls.database.wk
spec:
  group: database.wk
  names:
    kind: XJetPostgreSQL
    plural: xjetpostgresqls
  claimNames:
    kind: JetPostgreSQL
    plural: jetpostgresqls
  versions:
  - name: v1alpha1
    served: true
    referenceable: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              parameters:
                type: object
                properties:
                  storageGB:
                    type: integer
                    description: size of the Database in GB - integer
                  dbName:
                    type: string
                    description: name of the new DB inside the DB instance - string
                  instanceSize:
                    type: string
                    description: instance size - string
                    enum:
                      - small
                      - medium
                      - large
                required:
                  - storageGB
                  - dbName
                  - instanceSize
            required:
              - parameters
```

Here I can already see different problems. I wanted to create a DB, and a DB user with the same name, but what if I want 2 DBs ? 3 DBs? 2 users per  DB ?

在这里我已经可以看到不同的问题。我想创建一个数据库和一个同名的数据库用户，但是如果我想要 2 个数据库怎么办？ 3个数据库？每个数据库 2 个用户？

I guess I have to re-write the schema to use something like:

我想我必须重新编写架构才能使用类似的东西：

```
properties:
  dbs:
    type: array
    items:
      type: object
      properties:
        name:
          type: string
          description: name of the new DB inside the DB instance - string
        users:
          type: array
          items:
             properties:
               name:
                 type: string
```

Whatever, modelling those is not straightforward… but, well, it’s a one time  effort. Totally worth it. Take your time building this, as it defines  the parameters that your Dev team will use to create resources. Whatever is not defined here will use the defaults from the composition or the  Provider. You’re creating your abstraction.

不管怎样，对这些进行建模并不简单……但是，好吧，这是一次性的工作。完全值得的。花点时间构建它，因为它定义了您的开发团队将用于创建资源的参数。此处未定义的任何内容都将使用组合或提供者的默认值。你正在创造你的抽象。

## Compositions

Now we can create the Composition, which will take the values defined by the XRD and apply them to the Cloud Resources we need.

现在我们可以创建组合，它将采用 XRD 定义的值并将它们应用于我们需要的云资源。

```
apiVersion: apiextensions.crossplane.io/v1
kind: Composition
metadata:
  name: jetpostgresql.gcp.database.wk
  labels:
    provider: gcp
    crossplane.io/xrd: xjetpostgresql.database.wk
spec:
  # should I set this here ?Please help
  # writeConnectionSecretsToNamespace: crossplane
  compositeTypeRef:
    apiVersion: database.wk/v1alpha1
    kind: XJetPostgreSQL
  resources:
    - name: cloudsqlinstance
      base:
        apiVersion: sql.gcp.jet.crossplane.io/v1alpha2
        kind: DatabaseInstance
        metadata:
          annotations:
            crossplane.io/external-name: "crossplanesqlinstance"
        spec:
          providerConfigRef:
            name: crossplane-provider-jet-gcp
          deletionPolicy: Delete
          forProvider:
            databaseVersion: POSTGRES_14
            region: us-central1
            deletionProtection: false
            settings:
            - tier: db-custom-1-3840
              diskType: PD_SSD
              diskSize: 20
              ipConfiguration:
                - ipv4Enabled: true
                  authorizedNetworks:
                    - value: "0.0.0.0/0"
            userLabels:
              creator: crossplane
              owner: prune
          writeConnectionSecretToRef:
            namespace: crossplane
            name: cloudsqlinstance
      patches:
        # set diskSize based on the Claim
        - fromFieldPath: "spec.parameters.storageGB"
          toFieldPath: "spec.forProvider.settings[0].diskSize"
        # set the secret name to the claim name
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "spec.writeConnectionSecretToRef.name"
          transforms:
            - type: string
              string:
                fmt: "%s-pginstance"
        # change secret namespace to the one of the claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-namespace]"
          toFieldPath: "spec.writeConnectionSecretToRef.namespace"
        # set label app = name of the original claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "metadata.labels[crossplane.io/app]"
        # set the name of the external resource to be the name of the claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "metadata.annotations[crossplane.io/external-name]"
        # set instance size to the one defined in the claim
        - fromFieldPath: "spec.parameters.instanceSize"
          toFieldPath: "spec.forProvider.settings[0].tier"
          transforms:
            - type: map
              map:
                small: db-custom-1-3840
                medium: db-custom-2-7680
                large: db-custom-4-15360
          policy:
            fromFieldPath: Required
    - name: cloudsqldb
      base:
        apiVersion: sql.gcp.jet.crossplane.io/v1alpha2
        kind: Database
        metadata:
          annotations:
            crossplane.io/external-name: "crossplanesqldb"
        spec:
          providerConfigRef:
            name: crossplane-provider-jet-gcp
          deletionPolicy: Delete
          forProvider:
            instanceSelector:
              MatchControllerRef: true
          writeConnectionSecretToRef:
            namespace: crossplane
            name: cloudsqldb
      patches:
        # set the secret name to the claim name
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "spec.writeConnectionSecretToRef.name"
          transforms:
            - type: string
              string:
                fmt: "%s-pgdb"
        # change secret namespace to the one of the claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-namespace]"
          toFieldPath: "spec.writeConnectionSecretToRef.namespace"
        # set the name of the DB resource to be the name defined in the claim
        - fromFieldPath: "spec.parameters.dbName"
          toFieldPath: "metadata.annotations[crossplane.io/external-name]"
        # set app Label
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "metadata.labels[crossplane.io/app]"
    - name: cloudsqldbuser
      base:
        apiVersion: sql.gcp.jet.crossplane.io/v1alpha2
        kind: User
        metadata:
          annotations:
            # set the name of the DB User, this is hardcoded for demo but should come from the CRD
            crossplane.io/external-name: "existing-sa-for-db@my-project.iam"
        spec:
          providerConfigRef:
            name: crossplane-provider-jet-gcp
          deletionPolicy: Delete
          forProvider:
            instanceSelector:
              MatchControllerRef: true
            type: CLOUD_IAM_SERVICE_ACCOUNT
          writeConnectionSecretToRef:
            namespace: crossplane
            name: cloudsqluser
      patches:
        # set the secret name to the claim name
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "spec.writeConnectionSecretToRef.name"
          transforms:
            - type: string
              string:
                fmt: "%s-pguser"
        # change secret namespace to the one of the claim
        - fromFieldPath: "metadata.labels[crossplane.io/claim-namespace]"
          toFieldPath: "spec.writeConnectionSecretToRef.namespace"
        # set the name of the DB User, this is hardcoded for demo but should come from the Claim CRD
        # - fromFieldPath: "spec.parameters.dbName"
        #   toFieldPath: "metadata.annotations[crossplane.io/external-name]"
        # set app Label
        - fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
          toFieldPath: "metadata.labels[crossplane.io/app]"
```

Here you can see that it’s not trivial either. A lot is going on here. Globally, you just list all the Provider’s Resources you want to  instanciate, give default parameters you want to enforce, and patch some others from user-supplied values or other resources values.

在这里你可以看到它也不是微不足道的。这里发生了很多事情。在全球范围内，您只需列出您想要实例化的所有提供者资源，提供您想要强制执行的默认参数，并根据用户提供的值或其他资源值修补一些其他资源。

Let's go back to the issue I had describing the DBs and Users: Let's say I  have a list of DBs, and each DB has a list of Users… Then I need to  iterate over the values and create as many `Database.sql.gcp .jet.crossplane.io` and `User.sql.gcp.jet.crossplane.io` that is in the list. I would achieve that in TF using a `for_each` but here, in Crossplane… no idea (Help me if you know how).

让我们回到我描述数据库和用户的问题：假设我有一个数据库列表，每个数据库都有一个用户列表……然后我需要迭代这些值并创建尽可能多的`Database.sql.gcp .jet.crossplane.io` 和 `User.sql.gcp.jet.crossplane.io` 在列表中。我会在 TF 中使用 `for_each` 来实现这一点，但在这里，在 Crossplane 中……不知道（如果你知道怎么做，请帮助我）。

You can also see that each resource has `writeConnectionSecretToRef` which points to a secret `name` and `namespace` which in turn, should hold all the sensible values that the Resource  may create (like DB password, SSL certs, maybe even URLs, and whatever  you want to add into the secret).

您还可以看到每个资源都有 `writeConnectionSecretToRef`，它指向一个秘密的 `name` 和 `namespace`，它们反过来应该包含资源可能创建的所有合理值（比如数据库密码、SSL 证书，甚至可能是 URL，以及您想添加到秘密中的任何内容）。

I’m still unsure if it’s the right way of doing it. Most examples are writing the secrets in the `crossplane` namespace, or at least in a namespace fixed in the composition… This  does not make sense to me. I want the secret to be in the same `namespace` as the Developer’s claim that is creating the resource… This is making the `writeConnectionSecretsToNamespace` value that is set at the `Composition` `Spec` level (I commented it out).

我仍然不确定这样做是否正确。大多数示例都在 `crossplane` 命名空间中写入秘密，或者至少在组合中固定的命名空间中……这对我来说没有意义。我希望秘密与开发人员创建资源的声明位于相同的“名称空间”中……这使得在“组合”“规范”级别设置的“writeConnectionSecretsToNamespace”值（我将其注释掉）。

I guess it’s the moment to tell you about the Claims.

我想是时候告诉你有关索赔的时间了。

# Claims

# 声明

The claim is the easiest, at first, as it’s the abstraction: it is really limited.
It’s meant to be created by your developers that will want a DB along their  apps, or a Pub/Sub, or whatever resource you want them to have full  control on.

首先，声明是最简单的，因为它是抽象的：它确实是有限的。
它旨在由您的开发人员创建，他们希望在他们的应用程序中有一个数据库，或者一个发布/订阅，或者您希望他们完全控制的任何资源。

It is the only resource that is *namespaced*. This is made so you can define RBACs and allow this team (this namespace) to hold DBs, but no Pub/Sub, for example.

它是唯一*命名空间*的资源。这样做是为了让您可以定义 RBAC 并允许这个团队（这个命名空间）持有数据库，但没有 Pub/Sub，例如。

Here’s mine:

这是我的：

```
apiVersion: database.wk/v1alpha1
kind: JetPostgreSQL
metadata:
  namespace: test-namespace
  name: jet-db-claim
spec:
  parameters:
    storageGB: 25
    dbName: xrdb
    instanceSize: small # small, medium, large
  writeConnectionSecretToRef:
    name: jet-db-claim-details
```

Neat ! Simple ! Thanks abstractions ! Devs only specify what they care about, and you take care of all the boring stuff !

整洁的 ！简单的 ！感谢抽象！开发者只指定他们关心的事情，而你负责所有无聊的事情！

Going back to Secrets, you see that I also defined which secret to write  stuff into… but it seems this value is not an override of what is in the `Composition`. So I patched the composition to actually create the secret with a name derived from the `Claim` name and in the same `Namespace` as the `Claim` :

回到 Secrets，你会看到我还定义了将内容写入哪个秘密……但似乎这个值并没有覆盖 `Composition` 中的内容。因此，我对合成进行了修补，以实际创建秘密，其名称源自“Claim”名称，并且与“Claim”位于相同的“Namespace”中：

```
# set the secret name to reference the claim name
- fromFieldPath: "metadata.labels[crossplane.io/claim-name]"
  toFieldPath: "spec.writeConnectionSecretToRef.name"
  transforms:
    - type: string
      string:
        fmt: "%s-pginstance"
# change secret namespace to the one of the claim
- fromFieldPath: "metadata.labels[crossplane.io/claim-namespace]"
  toFieldPath: "spec.writeConnectionSecretToRef.namespace"
```

I first tried to use `metadata.name` to get the claim name, but, in fact, the `Composition` is not templated from the `Claim` but from the intermediate `XR` that is created from the `Claim` . Refer to my schema above if needed. So, in the `XR` the only way to get back to the `Claim` metadata is by looking at some specific `labels` like `crossplane.io/claim-name`.

我首先尝试使用 `metadata.name` 来获取声明名称，但实际上，`Composition` 不是从 `Claim` 模板化而来，而是来自从 `Claim` 创建的中间 `XR`。如果需要，请参阅上面的架构。因此，在 `XR` 中，返回到 `Claim` 元数据的唯一方法是查看一些特定的 `labels`，例如 `crossplane.io/claim-name`。

# Question 3: Provider’s execution

# 问题三：Provider的执行

Once you have all that, the Providers are going to start doing things:  reconciling between what you created in K8s and the reality in the  Cloud.

一旦你拥有了所有这些，供应商就会开始做事：协调你在 K8s 中创建的内容与云中的现实。

For that, the official OSS Providers seems to do a great work without much surprise.
For the Jet provider, at least the GCP one I tested, it’s another story. I started mine with `--debug` to better understand.
It’s k8s native, so also `k describe` on the many resources (XRD, Composition, XR, Claim and intermediate CR) to get the logs and events attached to each.

为此，官方 OSS Providers 似乎做得很好，这并不令人意外。
对于 Jet 提供商，至少是我测试的 GCP 提供商，这是另一回事。我从 `--debug` 开始我的，以便更好地理解。
它是 k8s 原生的，因此也在许多资源（XRD、Composition、XR、Claim 和中间 CR）上进行 `k describe` 以获取附加到每个资源的日志和事件。

I would also suggest that you add the `--debug` option to Crossplane itself to fully understand when and where the problem is when rendering the Composition.

我还建议您将 `--debug` 选项添加到 Crossplane 本身，以充分了解渲染合成时问题出在何时何地。

Crossplane keeps reconcilling the resource, so you may end up with tons of logs. 

Crossplane 不断协调资源，因此您最终可能会得到大量日志。

In the end, I wasn’t able to create a working DB user using Workload  Identity. Not sure why, but once the user created, Crossplane was trying to change the `username` to the name of the claim… and I was constantly seeing the `crossplane.io/external-name` value being switched from what I asked in the Composition and the Claim name…

最后，我无法使用 Workload Identity 创建工作数据库用户。不知道为什么，但是一旦用户创建，Crossplane 就试图将 `username` 更改为声明的名称......而且我经常看到 `crossplane.io/external-name` 值从我在成分和权利要求名称……

I also wasn’t able to understand what kind of values was the Provider /  the Resource returning that I could put in a Secret. I think this is a  difference between the OSS Provider and the Jet Provider… with Jet, it's like with TF: YOU have to create the secret and provide it to K8s… so  maybe I only used resources that had nothing to return as a secret value ?

我也无法理解 Provider / Resource 返回什么样的值，我可以将其放入 Secret 中。我认为这是 OSS Provider 和 Jet Provider 之间的区别……对于 Jet，就像对于 TF：你必须创建秘密并将其提供给 K8s……所以也许我只使用了没有任何东西可以返回的资源作为秘密值？

# Question 4: Support

# 问题 4: 支持

I asked many questions on Slack, I even opened an issue… and globally I  had no answer. I also tried to participate and read through the other questions/bugs/PRs… and it does not feel the community is heavily  active. Don’t get me wrong, I’m not saying the project is dead or whatever. Its more a feeling… Ask about TF and you’ll have 15 answers in the hour…

我在 Slack 上问了很多问题，我什至打开了一个问题……在全球范围内我都没有答案。我也尝试参与并通读其他问题/错误/PR……但感觉社区并不活跃。不要误会我的意思，我并不是说这个项目已经死了或其他什么。更像是一种感觉……问问TF，一小时内会有15个答案……

I talked about docs and examples, and again it’s really limited at the moment.

我谈到了文档和示例，但目前它真的很有限。

If i’m about to switch my whole IaC in a new tool, I want to be sure that it’s active and reactive.

如果我要在一个新工具中切换我的整个 IaC，我想确保它是主动的和被动的。

I just reached to [Upbound](https://www.upbound.io/), the company behind Crossplane, to see if the paid plan and support would help me solve my issues.
I’m really pleased to see that I had an answer quickly along with an  appointment with an engineer. According to what I was told, the OSS  Providers (and Jet) are way behind what Upbound offers to paid  customers. Maybe it’s just the way to go ? After all, TF also have a  paid subscription, which is far from cheap… and if I put all my IaC  somewhere, maybe a little support is a good idea.

我刚刚联系了 Crossplane 背后的公司 [Upbound](https://www.upbound.io/)，看看付费计划和支持是否能帮助我解决问题。
我很高兴看到我很快就得到了答复以及与工程师的约会。据我所知，OSS 提供商（和 Jet）远远落后于 Upbound 为付费客户提供的服务。也许这只是要走的路？毕竟，TF 也有付费订阅，这可不便宜……而且如果我把我所有的 IaC 都放在某个地方，也许一点支持是个好主意。

# Conclusion

#  结论

My POC was really limited in time, and there’s so much more I wish I had time to cover…

我的 POC 时间真的很有限，还有很多我希望我有时间涵盖......

My conclusion is that I could use Crossplane for my Developer’s infra, IF the need is really small…

我的结论是，如果需求真的很小，我可以将 Crossplane 用于我的开发人员的基础设施……

For example, if you only create one DB and one User per DB Instance.

例如，如果您只为每个数据库实例创建一个数据库和一个用户。

Another way of doing would certainly be to not Compose multiple resources in one `Composition` but have one `Composition` per resource… we I will end up with one `MyDBInstance` , one `MyDB`, two or more `MyDbUser` in each namespace/dev project… I'm even not sure if this model would  work… but at least, it would allow me to build an abstraction with  limited levers that I can present to the Dev's teams while still  allowing flexibility.

另一种方法当然是不在一个 `Composition` 中组合多个资源，而是每个资源有一个 `Composition` ……我们最终会得到一个 `MyDBInstance`、一个 `MyDB`、两个或更多的 `MyDbUser`命名空间/开发项目……我什至不确定这个模型是否可行……但至少，它可以让我用有限的杠杆构建一个抽象，我可以将其呈现给开发团队，同时仍然允许灵活性。

At the moment I would say that this POC just showed me that Crossplane is not an out-of-the-box easy replacement to TF.

目前我会说这个 POC 只是向我展示了 Crossplane 并不是 TF 的开箱即用的简单替代品。

I will keep playing with Crossplane on the side, until I fully understand it. I will also investigate the GCP official tool. 

我会一直在旁边玩Crossplane，直到我完全理解为止。我还将研究 GCP 官方工具。

