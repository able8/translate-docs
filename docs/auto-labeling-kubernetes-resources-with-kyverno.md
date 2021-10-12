# Auto-labeling Kubernetes resources with Kyverno

[CNCF Member Blog Post](http://www.cncf.io/lf-author-category/member/ "See more content from Member")

Posted on
December 30, 2020 By Anubhav Sharma

_Guest post originally published on [Nirmata’s blog](https://nirmata.com/2020/10/30/auto-labeling-kubernetes-resources-with-kyverno/) by Anubhav Sharma, VP, Business Development & Customer Success at Nirmata_

## Introduction

As Kubernetes has become the foundational building block for enterprises to go cloud-native, the last couple of years have seen many solutions that have simplified the cluster creation process. But the Day-2 operations around Kubernetes still remains a complex endeavor, slowing down adoption and increasing the operational costs. Kubernetes’ complexity and skills gaps still remain the biggest factors that are in the way of Enterprise’s adoption of Kubernetes.

Many of the Day-2 operations use cases include requirements for the central platform team to deliver secure and compliant environments to developers as efficiently as possible with necessary services and best practices preconfigured. Some examples of such use cases include configuring environments with Kubernetes best practices like resource quotas, network policy and pod security. This requires tools that can assess the environments as they are created and then configure them in compliance with the standard defined by the central platform team.

## **Kyverno: A Flexible Ops Tool for K8s**

Kubernetes provides powerful constructs like admission control webhooks that can be leveraged for the purposes of validating and mutating resources. Nirmata’s [Kyverno](https://kyverno.io/) was designed specifically to address these types of use cases using the declarative paradigm. Kyverno is an open-source policy engine that was designed for Kubernetes, It provides users with familiar constructs to write custom rules and easily implement to validate, mutate, and generate new resources as needed.

Managing Kubernetes at scales requires the following best practices and applying standardization across configurations. One such pattern is to use Kubernetes labels. In Kubernetes, every resource can have one or more labels and Kubernetes makes it easy to find and manage the resources using labels.

A very common use case for Day-2 operations is managing labels across namespaces and pods so that use cases like certificate updates, self-service logging/monitoring, backups etc. can be easily implemented by other Kubernetes controllers and operators.

## **Auto-Labeling Namespaces**

Below is an example of how to implement namespace labeling upon creation in a Kubernetes cluster using Kyverno.

Install Kyverno in your cluster:

```
kubectl create -f https://github.com/kyverno/kyverno/raw/master/definitions/install.yaml
```

Detailed installation instructions are available [here](https://kyverno.io/docs/installation/).

Here is a sample Kyverno policy that adds labels to namespaces  –

```
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
name: add-labels
spec:
background: false
rules:
  - name: add-ns-label
    match:
      resources:
        kinds:
        - Namespace
    exclude:
      clusterroles: ["cluster-admin"]
    mutate:
      patchStrategicMerge:
        metadata:
          labels:
            kyverno/user: "{{ request.userInfo.username }}"
            +(kyverno/network): "default"
```

The policy inserts a label \`kyverno/user\`with the value of the user making the API request to create the namespace. The policy also inserts a label \`kyverno/network\`, but only if one is not already specified by the user. This simple policy demonstrates some powerful features in Kyverno like [variable substitution](https://kyverno.io/docs/writing-policies/writing-policies-variables/) and [conditional anchors](https://kyverno.io/docs/writing-policies/writing-policies-validate/).

Once the policy is configured in your cluster, create a new namespace and verify the labels have been added to the namespace automatically.

Create a new namespace:

```
kubectl create ns test

```

View the namespace:

```
kubectl get ns test -o yaml
```

This should show a namespace similar to:

```
apiVersion: v1
kind: Namespace
metadata:
labels:
    kyverno/network: default
    kyverno/user: docker-for-desktop
```

Now, what if you want to make sure that users cannot update a specific label?

Kyverno makes that easy to do as well! Here is a policy that prevents the update of the \`kyverno/network\` label:

```
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
name: protect-label
spec:
validationFailureAction: enforce
background: false
rules:
  - name: block-updates-for-label
    match:
      resources:
        kinds:
           - Namespace
validate:
    message: Updating label `kyverno/network` is not allowed
    deny:
    - key: "{{ request.operation }}"
      operator: "EQUALS"
      value: UPDATE
```

## **Summary**

Managing Kubernetes configurations can be complex, and policy engines provide standardization, automated validation, and the ability to mutate and generate configurations.

Kyverno is an open-source policy engine designed for Kubernetes. It has a minimal learning curve and provides tremendous flexibility for Kubernetes administrators to solve Day-2 operations challenges using Kubernetes’ powerful declarative management capabilities and native tools.

Learn what else Kyverno can do at [https://kyverno.io](https://kyverno.io/).

Share this post