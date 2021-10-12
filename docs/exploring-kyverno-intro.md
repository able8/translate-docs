# Exploring Kyverno: Introduction

Nov 20, 2020    

------

### Articles in the Exploring Kyverno series

#### **[Part 1, Validation](https://neonmirrors.net/post/2020-12/exploring-kyverno-part1)**

#### **[Part 2, Mutation](https://neonmirrors.net/post/2020-12/exploring-kyverno-part2)**

#### **[Part 3, Generation](https://neonmirrors.net/post/2020-12/exploring-kyverno-part3)**

------

It's all around us. It's everywhere. And yet it's nowhere. If you  guessed "Kubernetes" you'd be partially right. Everyone seems to be  talking about it, more and more companies are using it, but what isn't  growing at the same rate is security and controls for it. And by that I  mean specifically policy. The ability to reign in what seems like the  Wild West and build guard rails and governance, ensuring users can't do  dangerous or non-compliant things, is sadly all too absent from  Kubernetes adopters at this point. And it's not because there aren't  enough lessons learned to build a series of good hygienical practices,  it's because the tooling has either been nonexistent or painful to work  with. Or, in other cases, policy was like an "extra" or "value add"  feature that only came when using a managed platform of some sort. Those days are gone or, at the very least, are rapidly waning into the sunset thanks to tools like OpenPolicy Agent and [Kyverno](https://kyverno.io/). It's the latter I wish to explore in this multi-part series, and I'm  honestly really excited to write about this open-source project. For the first time, policy in Kubernetes can be easy, effective, flexible, and  also (dare I say) fun to write. Stick around and let me show you [Kyverno](https://kyverno.io/), an extensive, Kubernetes-native policy engine.

Kyverno is an open-source policy engine built specifically for  Kubernetes to not only validate and ensure requests conform to your  internal best practices and policies, but to modify those requests if  needed and even create new objects based on a variety of conditions.  It's a project that came out of [Nirmata](https://nirmata.com/) and was just recently donated to the CNCF as a [Sandbox project](https://www.cncf.io/sandbox-projects/). Although it's a project I had seen before, I had only paid cursory  attention. Not until it was accepted into the CNCF did I decide to  really give it a hard look and dig in. I'm glad I did, because what I  found was frankly awesome. I think this project has enormous potential  to become the *defacto* standard of Kubernetes policy application, and in this multi-part series of articles I hope to explain why.

### Escaping Cognitive Overload

If you're already using Kubernetes, it's a sure bet you're using other  tools, languages, and frameworks to complete whatever picture you're  painting, right? How much knowledge have you acquired that is only  applicable to one of those? How much technical debt do you carry on your shoulders? It's so much, you might not even be able to keep track.  Simply put, it's cognitive overload. If you're like most, you probably  need yet another one of those things like you need another hole in your  head. So rather than learning *yet another* Terraform or *yet another* Ansible or *yet another* {insert_bespoke_tool_here}, wouldn't it be nice if you could repurpose  that knowledge of how Kubernetes and YAML works to go straight into  being productive with a new tool without having to feel like you're  learning a new programming language? Who's saying "nah, I want more  complexity" here?

When it comes to making a choice on how to apply policy to Kubernetes, the two leading options are [OpenPolicy Agent](https://www.openpolicyagent.org/) (OPA) (via [Gatekeeper](https://github.com/open-policy-agent/gatekeeper)), and [Kyverno](https://kyverno.io/). In the context of Kubernetes, both of them operate as [admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/), allowing their engines to check incoming requests prior to creation to ensure they conform to whatever pattern is defined.

> Gatekeeper presently only operates as a validation webhook; it does not have mutation ability.

The main power behind OPA is its ability to apply to more than just  Kubernetes, but do so with a unified language. And if you're truly  looking to adopt one method for applying policy across your enterprise,  you should definitely give OPA a look. But it's a huge investment  because it's vastly complex. And what I've seen in companies is that  they've already got tools to do (most of) this, which means they've  already gone through the learning curve of those, which means they're  probably unlikely to throw it all away in favor of something new.

Let's say you need some form of policy controls in your Kubernetes  environment. You're not familiar with anything, but you have a need.  Then you run across this Gatekeeper sample.

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
          provided := {label | input.review.object.metadata.labels[label]}
          required := {label | label := input.parameters.labels[_]}
          missing := required - provided
          count(missing) > 0
          msg := sprintf("you must provide labels: %v", [missing])
        }        
```

Figure 1: Great. Yet another language for me to learn.

You see that last chunk starting with `rego:`? OpenPolicy has its own esoteric programming language called [Rego](https://www.openpolicyagent.org/docs/latest/#rego) which is required to work with policies.

How about this instead?

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

[   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

[ 	   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[ 	   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[ 	   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[ 	                                 ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[ 	   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[ 	   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[ 	   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[ 	   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)[ 	   	 		](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

[ 	                                 ](https://neonmirrors.net/post/2020-11/exploring-kyverno-intro/#)

yaml

I'll bet, even if you  aren't familiar with Kyverno, you can probably figure out what this does based on your existing knowledge of Kubernetes and YAML. Compare this  to the first example. Which would you rather write?

> By the way, the first snippet I provided isn't even a complete policy definition as Gatekeeper constraint templates still require a  constraint manifest. The second one, by contrast, is a complete and  fully-functional policy.

The other option is [PodSecurityPolicies](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) baked right into Kubernetes. This has its own [challenges](https://developer.squareup.com/blog/kubernetes-pod-security-policies/), is difficult to work with, and is a breaking change for existing clusters (it can't just be "flipped on"). It's also [possibly going away](https://github.com/kubernetes/enhancements/issues/5) without a clear path forward at the moment.

All of these existing solutions are cumbersome in their own ways yet have a degree of functionality that many may still prefer.

### Enter Kyverno

Kyverno's primary appeal is in its simplicity of how policy is defined. Using  concepts you already know with authoring tools you already use, you can  very quickly and simply start writing policies that really solve  problems. And with an extensive (and growing) [sample policy library](https://github.com/kyverno/kyverno/tree/main/samples), you may not have to write anything at all. But past that, Kyverno's  main strengths lie in three different categories. Using the same engine  and authoring policies the same way, Kyverno can:

- **Validate** incoming requests to ensure they are  accepted or rejected based upon rules which you define. Not only does  this work for resources like `Pods` or `Deployments` or the like, but also actions like DELETE requests. This is powerful  because you can use Kyverno to augment the existing RBAC system within  Kubernetes to do more than it can natively.
- **Mutate** new objects to add/remove/change fields based on how you desire new resources conform to standards. Actions like  adding labels or inserting fields are easy.
- **Create** any object based upon the request to create  another object. This powerful ability can help cut down on related  provisioning steps when standard resources are created. Maybe for each  namespace that gets created, you typically have to create five other  resources. Kyverno can do that for you with policy.

Each one of these is a lot to unpack, so in this series I'm going to dedicate an article to each one individually.

Check back as I show the **validate** capabilities in the first part of the series.

## Chip Zoller

​      Technologist, perpetual student, teacher, continual incremental improvement.    
