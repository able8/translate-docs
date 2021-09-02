## Enforcing Policy as Code using OPA and Gatekeeper in Kubernetes

April 15, 2021

Enforcing organizational policies on a Kubernetes cluster allows you to be in control of the resources being deployed. For example, you can prevent deploying non-vetted pods to the production environment or disable usage of default passwords for databases. Doing so stops you from worrying about quarterly security reviews and a ton of issues in your backlog.

In this blog post, we will go through everything necessary to set up “OPA(Open Policy Agent)”/Gatekeeper as your Kubernetes admission webhook, which enables you to enforce policies on your Kubernetes cluster.

## Open Policy Agent (OPA)

Open Policy Agent is an open-source, general-purpose policy engine that enforces validation of objects during creation, updating, and deletion operations. OPA lets us enforce custom policies on Kubernetes objects without manually reconfiguring the Kubernetes API server. It will ensure that no Deployments, Jobs, Pods, etc are scheduled without being compliant with your Constraints and rules.

OPA was designed to let us write policies over arbitrary JSON/YAML. You can essentially use OPA to enforce policies on any tool that takes JSON/YAML as input, such as Kubernetes, Terraform, CI/CD pipelines. For information on OPA and its use cases, please refer to the [official documentation](https://www.openpolicyagent.org/docs/latest/). We will focus on OPA’s Kubernetes [admission controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) use case with Gatekeeper.

## Gatekeeper

Gatekeeper is a customizable admission webhook for Kubernetes that dynamically enforces policies executed by the OPA. Gatekeeper uses [CustomResourceDefinitions](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/) internally and allows us to define **ConstraintTemplates** and **Constraints** to enforce policies on Kubernetes resources such as Pods, Deployments, Jobs.

OPA/Gatekeeper uses its own declarative language called Rego, a query language. You define rules in Rego which, if invalid or returned a false expression, will trigger a constraint violation and blocks the ongoing process of creating/updating/deleting the resource.

### Prerequisites

- kubectl CLI
- helm CLI
- A running Kubernetes cluster. If you don’t have a cluster running we would recommend using[kubespray](https://github.com/kubernetes-sigs/kubespray).

To install gatekeeper, run the following helm commands.

```
helm repo add gatekeeper https://open-policy-agent.github.io/gatekeeper/charts

helm install gatekeeper/gatekeeper --generate-name
```

> **_NOTE:_** To check whether it is installed, the simplest way is to see if Gatekeeper controller manager and Gatekeeper audit deployments and services are deployed in the gatekeeper-system namespace.

Now, before we go into the demo, you should know about constraint templates and constraints. Every policy in the Gatekeeper consists of two manifests, Constraint template, and Constraint.

To define a Constraint, you need to create a Constraint Template that allows people to declare new Constraints

## ConstraintTemplate

A ConstraintTemplate consists of both the Rego logic that enforces the Constraint and the schema for the Constraint, which includes the schema of the CRD and the parameters that can be passed into a Constraint.

## Constraint

Constraint is an object that says on which resources are the policies applicable, and also what parameters are to be queried and checked to see if they are available in the resource manifest the user is trying to apply in your Kubernetes cluster. Simply put, it is a declaration that its author wants the system to meet a given set of requirements.

## Working example

Here is a ConstraintTemplate CRD that requires certain labels to be present on the objects that are being created/deployed in a cluster.

```
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
       # Schema for the `parameters` field
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
         msg := sprintf("\n\nDENIED. \nReason: Our org policy mandates the following labels: \nYou must provide these labels: %v", [missing])
       }
```

Once ConstraintTemplate Object is deployed, you can now enforce your polices by creating Constraints. Here is a Constraint that requires labels “zone”,“stage”,“status”.

> **_NOTE:_** This policy does not consider the values of the labels, it only checks whether the resources that are being deployed have the required labels attached to it or not. If the provided labels does not match the required labels, the constraint denies the objects creation.

```
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredLabels
metadata:
name: label-check
spec:
match:
kinds:
     - apiGroups: [""]
       kinds: ["Namespace", "Pod"]
excludedNamespaces:
   - kube-system
   - kube-public
   - kube-node-lease
   - gatekeeper-system
parameters:
labels:
     - "zone"
     - "stage"
     - "status"
```

You can copy and save these as `constrainttemplate.yml` and `constraints.yml` and run the following commands to apply the objects.

```
kubectl apply -f constrainttemplate.yml
kubectl apply -f constraints.yml
```

Now, once applied, This constraint checks if every pod or namespace that is being created has the specified labels attached to it. You can test it by creating a pod or a namespace with and without labels and see the difference for yourselves. Here is an output of a pod that is being deployed without any labels.

```
$ kubectl run nginx --image=nginx --generator=run-pod/v1
Error from server ([denied by label-check]

DENIED.
Reason: Our org policy mandates the following labels:
You must provide these labels: {"stage", "status", "zone"}): admission webhook "validation.gatekeeper.sh" denied the request: [denied by label-check]

DENIED.
Reason: Our org policy mandates the following labels:
You must provide these labels: {"stage", "status", "zone"}
```

The policy we set using the Constraint, denied the pod creation without required labels.

Here is an output of a pod that is being deployed adhering to the policy of having required labels **_stage_**, **_status_**, **_zone_**.

```
$ kubectl run nginx --image=nginx --generator=run-pod/v1 -l zone=us-east-1,stage=dev,status=ready
pod/nginx created
```

#### Note

You as an admin need to explicitly decide your organization specific policies and then turn your policies as code using Rego.

## Community-provided libraries of policies

You don’t have to write policies on your own at the beginning of your journey, OPA and Gatekeeper both have excellent community libraries. You can have a look, fork them, and use them in your organization from here, [OPA](https://github.com/open-policy-agent/library), and [Gatekeeper](https://github.com/open-policy-agent/gatekeeper-library) libraries.

A glimpse on getting started with the Gatekeeper library, the library provided in the bonus assumes that you have prior knowledge or have already played with the Gatekeeper. If you have no prior knowledge, I would recommend going through this blog post and the [official documentation](https://open-policy-agent.github.io/gatekeeper/website/docs/) of Gatekeeper again, understand how Rego queries allow and deny objects.

## Summary

In this blog post, we have done everything needed to enforce policies using Gatekeeper in your Kubernetes cluster and we went through how Gatekeeper reduces Policy Compliance burden. So you can rest assured that common configuration errors or security issues will be automatically prevented. We are sure that this tutorial will help with your journey of running a Kubernetes cluster complying with your organizational policies.

## Read more of our **engineering blog posts**

This blog post is part of our engineering blog post series. Experience and expertise, straight from our engineering team. Always with a focus on technical, hands-on HOWTO content with copy-pasteable code or CLI commands.

Would you like to read more content like this? Click the button below and see the other blog posts in this series!

[Read more engineering blog posts](https://elastisys.com/category/tech-post/engineering/?utm_source=website&utm_medium=cta&utm_campaign=engineering)

Tags: [Blog](https://elastisys.com/tag/blog/), [devops](https://elastisys.com/tag/devops/), [devsecops](https://elastisys.com/tag/devsecops/), [gatekeeper](https://elastisys.com/tag/gatekeeper/), [open policy agent](https://elastisys.com/tag/open-policy-agent/), [policy as code](https://elastisys.com/tag/policy-as-code/)
