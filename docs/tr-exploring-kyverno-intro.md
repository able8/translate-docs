# Exploring Kyverno: Introduction

# 探索 Kyverno：介绍

Nov 20, 2020

2020 年 11 月 20 日

------

### Articles in the Exploring Kyverno series

### 探索 Kyverno 系列中的文章

#### **[Part 1, Validation](https://neonmirrors.net/post/2020-12/exploring-kyverno-part1)**

#### **[第 1 部分，验证](https://neonmirrors.net/post/2020-12/exploring-kyverno-part1)**

#### **[Part 2, Mutation](https://neonmirrors.net/post/2020-12/exploring-kyverno-part2)**

#### **[第 2 部分，突变](https://neonmirrors.net/post/2020-12/exploring-kyverno-part2)**

#### **[Part 3, Generation](https://neonmirrors.net/post/2020-12/exploring-kyverno-part3)**

#### **[第 3 部分，生成](https://neonmirrors.net/post/2020-12/exploring-kyverno-part3)**

------

It's all around us. It's everywhere. And yet it's nowhere. If you  guessed "Kubernetes" you'd be partially right. Everyone seems to be  talking about it, more and more companies are using it, but what isn't  growing at the same rate is security and controls for it. And by that I  mean specifically policy. The ability to reign in what seems like the  Wild West and build guard rails and governance, ensuring users can't do  dangerous or non-compliant things, is sadly all too absent from  Kubernetes adopters at this point. And it's not because there aren't  enough lessons learned to build a series of good hygienical practices,  it's because the tooling has either been nonexistent or painful to work  with. Or, in other cases, policy was like an "extra" or "value add"  feature that only came when using a managed platform of some sort. Those days are gone or, at the very least, are rapidly waning into the sunset thanks to tools like OpenPolicy Agent and [Kyverno](https://kyverno.io/). It's the latter I wish to explore in this multi-part series, and I'm  honestly really excited to write about this open-source project. For the first time, policy in Kubernetes can be easy, effective, flexible, and  also (dare I say) fun to write. Stick around and let me show you [Kyverno](https://kyverno.io/), an extensive, Kubernetes-native policy engine.

它就在我们身边。它无处不在。然而它无处可去。如果您猜对了“Kubernetes”，那您就部分正确了。似乎每个人都在谈论它，越来越多的公司在使用它，但没有以同样的速度增长的是它的安全性和控制。我的意思是具体的政策。遗憾的是，Kubernetes 的采用者目前还缺乏在看似狂野的西部进行统治并建立护栏和治理以确保用户不能做危险或不合规的事情的能力。并不是因为没有足够的经验教训来建立一系列良好的卫生实践，而是因为工具要么不存在，要么使用起来很痛苦。或者，在其他情况下，策略就像一个“额外”或“增值”功能，只有在使用某种托管平台时才会出现。由于 OpenPolicy Agent 和 [Kyverno](https://kyverno.io/) 等工具，那些日子已经一去不复返了，或者至少正在迅速消失。我希望在这个由多个部分组成的系列中探索后者，老实说，我很高兴能写下这个开源项目。 Kubernetes 中的策略第一次变得简单、有效、灵活，而且（我敢说)编写起来也很有趣。留下来，让我向您展示 [Kyverno](https://kyverno.io/)，这是一个广泛的 Kubernetes 原生策略引擎。

Kyverno is an open-source policy engine built specifically for  Kubernetes to not only validate and ensure requests conform to your  internal best practices and policies, but to modify those requests if  needed and even create new objects based on a variety of conditions. It's a project that came out of [Nirmata](https://nirmata.com/) and was just recently donated to the CNCF as a [Sandbox project](https://www.cncf.io/sandbox-projects/) . Although it's a project I had seen before, I had only paid cursory  attention. Not until it was accepted into the CNCF did I decide to  really give it a hard look and dig in. I'm glad I did, because what I  found was frankly awesome. I think this project has enormous potential  to become the *defacto* standard of Kubernetes policy application, and in this multi-part series of articles I hope to explain why.

Kyverno 是一个专门为 Kubernetes 构建的开源策略引擎，不仅可以验证和确保请求符合您的内部最佳实践和策略，还可以根据需要修改这些请求，甚至根据各种条件创建新对象。这是一个来自 [Nirmata](https://nirmata.com/) 的项目，最近刚刚作为 [Sandbox 项目](https://www.cncf.io/sandbox-projects/) 捐赠给了 CNCF .虽然这是我以前看过的一个项目，但我只是粗略地关注了。直到它被 CNCF 接受，我才决定真正认真地研究它并深入研究。我很高兴我做到了，因为坦率地说，我发现的东西很棒。我认为这个项目有巨大的潜力成为 Kubernetes 政策应用的*事实上*标准，我希望在这个由多部分组成的系列文章中解释原因。

### Escaping Cognitive Overload

### 逃避认知超载

If you're already using Kubernetes, it's a sure bet you're using other  tools, languages, and frameworks to complete whatever picture you're  painting, right? How much knowledge have you acquired that is only  applicable to one of those? How much technical debt do you carry on your shoulders? It's so much, you might not even be able to keep track. Simply put, it's cognitive overload. If you're like most, you probably  need yet another one of those things like you need another hole in your  head. So rather than learning *yet another* Terraform or *yet another* Ansible or *yet another* {insert_bespoke_tool_here}, wouldn't it be nice if you could repurpose  that knowledge of how Kubernetes and YAML works to go straight into  being productive with a new tool without having to feel like you're  learning a new programming language? Who's saying "nah, I want more  complexity" here? 

如果您已经在使用 Kubernetes，那么您肯定会使用其他工具、语言和框架来完成您正在绘制的任何图片，对吧？您获得了多少仅适用于其中之一的知识？您肩负了多少技术债务？太多了，您甚至可能无法跟踪。简而言之，这是认知超载。如果你和大多数人一样，你可能还需要另外一个东西，就像你需要在你的脑袋上再挖一个洞一样。因此，与其学习 *又一个 * Terraform 或 *又一个 * Ansible 或 *又一个 * {insert_bespoke_tool_here}，如果您可以重新利用 Kubernetes 和 YAML 工作原理的知识，直接通过新工具，而不必觉得您正在学习一种新的编程语言？谁在这里说“不，我想要更多的复杂性”？

When it comes to making a choice on how to apply policy to Kubernetes, the two leading options are [OpenPolicy Agent](https://www.openpolicyagent.org/)(OPA) (via [Gatekeeper](https://github.com/open-policy-agent/gatekeeper)), and [Kyverno](https://kyverno.io/). In the context of Kubernetes, both of them operate as [admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/), allowing their engines to check incoming requests prior to creation to ensure they conform to whatever pattern is defined.

在选择如何将策略应用于 Kubernetes 时，两个主要选项是 [OpenPolicy Agent](https://www.openpolicyagent.org/)(OPA)（通过 [Gatekeeper](https://github.com/open-policy-agent/gatekeeper)) 和 [Kyverno](https://kyverno.io/)。在 Kubernetes 的上下文中，它们都作为 [准入控制器](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) 运行，允许它们的引擎在创建之前检查传入的请求以确保它们符合定义的任何模式。

> Gatekeeper presently only operates as a validation webhook; it does not have mutation ability.

> Gatekeeper 目前仅作为验证 webhook 运行；它没有变异能力。

The main power behind OPA is its ability to apply to more than just  Kubernetes, but do so with a unified language. And if you're truly  looking to adopt one method for applying policy across your enterprise,  you should definitely give OPA a look. But it's a huge investment  because it's vastly complex. And what I've seen in companies is that  they've already got tools to do (most of) this, which means they've  already gone through the learning curve of those, which means they're  probably unlikely to throw it all away in favor of something new.

OPA 背后的主要力量是它不仅能够应用于 Kubernetes，而且能够使用统一的语言来实现。如果您真的希望采用一种方法在整个企业中应用策略，那么您绝对应该看看 OPA。但这是一项巨大的投资，因为它非常复杂。我在公司看到的是他们已经有工具来做（大部分）这意味着他们已经经历了这些的学习曲线，这意味着他们可能不太可能把它全部扔掉赞成新事物。

Let's say you need some form of policy controls in your Kubernetes  environment. You're not familiar with anything, but you have a need. Then you run across this Gatekeeper sample.

假设您需要在 Kubernetes 环境中进行某种形式的策略控制。你什么都不熟悉，但你有需要。然后您会遇到这个 Gatekeeper 示例。

```yaml
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: k8srequiredlabels
spec:
  crd:
    spec:
      names:
        kind: K8sRequiredLabels
      validation:
        openAPIV3Schema:
          properties:
            labels:
              type: array
              items: string
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
        package k8srequiredlabels
        violation[{"msg": msg, "details": {"missing_labels": missing}}] {
          provided := {label |input.review.object.metadata.labels[label]}
          required := {label |label := input.parameters.labels[_]}
          missing := required - provided
          count(missing) > 0
          msg := sprintf("you must provide labels: %v", [missing])
        }
```

Figure 1: Great. Yet another language for me to learn.

图 1：太好了。又一种语言让我学习。

You see that last chunk starting with `rego:`? OpenPolicy has its own esoteric programming language called [Rego](https://www.openpolicyagent.org/docs/latest/#rego) which is required to work with policies.

你看到最后一个以`rego:`开头的块了吗？ OpenPolicy 有自己的深奥编程语言 [Rego](https://www.openpolicyagent.org/docs/latest/#rego)，它是处理策略所必需的。

How about this instead?

换这个怎么样？

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-ns-labels
spec:
  validationFailureAction: enforce
  rules:
  - name: require-ns-labels
    match:
      resources:
        kinds:
        - Namespace
    validate:
      message: "The label `my-app` is required."
      pattern:
        metadata:
          labels:
            my-app: "?*"
```

[                ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

[ ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

[                     ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[                     ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#) [                     ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[                                      ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#) [                ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[                     ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#) [                     ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[                     ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#) [                     ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[                     ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

[ ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#) [ ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#) [ ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#) [ ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#) [ ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

[                                      ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

[ ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

yaml

雅姆

I'll bet, even if you  aren't familiar with Kyverno, you can probably figure out what this does based on your existing knowledge of Kubernetes and YAML. Compare this  to the first example. Which would you rather write?

我敢打赌，即使您不熟悉 Kyverno，您也可以根据您现有的 Kubernetes 和 YAML 知识弄清楚它的作用。将此与第一个示例进行比较。你更愿意写哪个？

> By the way, the first snippet I provided isn't even a complete policy definition as Gatekeeper constraint templates still require a  constraint manifest. The second one, by contrast, is a complete and  fully-functional policy. 

> 顺便说一下，我提供的第一个片段甚至不是一个完整的策略定义，因为 Gatekeeper 约束模板仍然需要一个约束清单。相比之下，第二个是完整且功能齐全的政策。

The other option is [PodSecurityPolicies](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) baked right into Kubernetes. This has its own [challenges](https://developer.squareup.com/blog/kubernetes-pod-security-policies/), is difficult to work with, and is a breaking change for existing clusters (it can't just be "flipped on"). It's also [possibly going away](https://github.com/kubernetes/enhancements/issues/5) without a clear path forward at the moment.

另一个选项是 [PodSecurityPolicies](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) 直接融入 Kubernetes。这有其自身的[挑战](https://developer.squareup.com/blog/kubernetes-pod-security-policies/)，难以使用，并且对现有集群来说是一个突破性的变化（它不能只是被“翻转”)。它也[可能消失](https://github.com/kubernetes/enhancements/issues/5) 目前没有明确的前进道路。

All of these existing solutions are cumbersome in their own ways yet have a degree of functionality that many may still prefer.

所有这些现有的解决方案都以自己的方式繁琐，但具有许多人可能仍然喜欢的一定程度的功能。

### Enter Kyverno

### 进入Kyverno

Kyverno's primary appeal is in its simplicity of how policy is defined. Using  concepts you already know with authoring tools you already use, you can  very quickly and simply start writing policies that really solve  problems. And with an extensive (and growing) [sample policy library](https://github.com/kyverno/kyverno/tree/main/samples), you may not have to write anything at all. But past that, Kyverno's  main strengths lie in three different categories. Using the same engine  and authoring policies the same way, Kyverno can:

Kyverno 的主要吸引力在于其政策定义方式的简单性。使用您已经知道的概念和您已经使用的创作工具，您可以非常快速且简单地开始编写真正解决问题的策略。有了广泛（并且不断增长）的 [示例策略库](https://github.com/kyverno/kyverno/tree/main/samples)，您可能根本不需要编写任何内容。但除此之外，Kyverno 的主要优势在于三个不同的类别。以相同的方式使用相同的引擎和创作策略，Kyverno 可以：

- **Validate** incoming requests to ensure they are  accepted or rejected based upon rules which you define. Not only does  this work for resources like `Pods` or `Deployments` or the like, but also actions like DELETE requests. This is powerful  because you can use Kyverno to augment the existing RBAC system within  Kubernetes to do more than it can natively.
- **Mutate** new objects to add/remove/change fields based on how you desire new resources conform to standards. Actions like  adding labels or inserting fields are easy.
- **Create** any object based upon the request to create  another object. This powerful ability can help cut down on related  provisioning steps when standard resources are created. Maybe for each  namespace that gets created, you typically have to create five other  resources. Kyverno can do that for you with policy.

- **验证**传入的请求，以确保根据您定义的规则接受或拒绝它们。这不仅适用于“Pods”或“Deployments”等资源，还适用于 DELETE 请求等操作。这很强大，因为您可以使用 Kyverno 来增强 Kubernetes 中现有的 RBAC 系统，以完成比原生更多的功能。
- **变异**新对象以根据您希望新资源符合标准的方式添加/删除/更改字段。添加标签或插入字段等操作很容易。
- **创建**基于创建另一个对象的请求的任何对象。这种强大的功能可以帮助在创建标准资源时减少相关的配置步骤。也许对于创建的每个命名空间，您通常必须创建五个其他资源。 Kyverno 可以通过政策为您做到这一点。

Each one of these is a lot to unpack, so in this series I'm going to dedicate an article to each one individually.

其中每一个都需要解压，因此在本系列中，我将分别为每一个单独写一篇文章。

Check back as I show the **validate** capabilities in the first part of the series.

当我在本系列的第一部分中展示 **validate** 功能时，请回来查看。

## Chip Zoller

## 奇普佐勒

      Technologist, perpetual student, teacher, continual incremental improvement. 

技术专家，永远的学生，老师，持续的增量改进。

