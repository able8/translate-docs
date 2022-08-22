# Cluster multi-tenancy

This page explains [cluster multi-tenancy](http://cloud.google.com#what_is_multi-tenancy) on Google Kubernetes Engine (GKE). This includes clusters shared by different users at a single organization, and clusters that are shared by per-customer instances of a software as a service (SaaS) application. Cluster multi-tenancy is an alternative to managing many single-tenant clusters.

This page also summarizes the Kubernetes and GKE features that can be used to manage multi-tenant clusters.

## What is multi-tenancy?

A multi-tenant cluster is shared by multiple users and/or workloads which are referred to as "tenants". The operators of multi-tenant clusters must isolate tenants from each other to minimize the damage that a compromised or malicious tenant can do to the cluster and other tenants. Also, cluster resources must be fairly allocated among tenants. 

When you plan a multi-tenant architecture you should consider the layers of resource isolation in Kubernetes: cluster, namespace, node, Pod, and container. You should also consider the security implications of sharing different types of resources among tenants. For example, scheduling Pods from different tenants on the same node could reduce the number of machines needed in the cluster. On the other hand, you might need to prevent certain workloads from being colocated.  For example, you might not allow untrusted code from outside of your organization to run on the same node as containers that process sensitive information.

Although Kubernetes cannot guarantee perfectly secure isolation between tenants, it does offer features that may be sufficient for specific use cases.  You can separate each tenant and their Kubernetes resources into their own [namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/). You can then use [policies](https://kubernetes.io/docs/concepts/policy/) to enforce tenant isolation. Policies are usually scoped by namespace and can be used to restrict API access, to constrain resource usage, and to restrict what containers are allowed to do.

The tenants of a multi-tenant cluster share:

- [Extensions](https://kubernetes.io/docs/concepts/extend-kubernetes/extend-cluster/), [controllers](https://kubernetes.io/docs/reference/glossary/?fundamental=true#term-controller), add-ons, and
[custom resource definitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) (CRDs).
- The cluster[control plane](https://kubernetes.io/docs/concepts/architecture/control-plane-node-communication/). This implies that the cluster operations, security, and auditing are centralized.

Operating a multi-tenant cluster has several advantages over operating multiple, single-tenant clusters:

- Reduced management overhead
- Reduced resource fragmentation
- No need to wait for cluster creation for new tenants

## Multi-tenancy use cases

This section describes how you could configure a cluster for various
multi-tenancy use cases.

### Enterprise multi-tenancy

In an enterprise environment, the tenants of a cluster are distinct teams within the organization. Typically, each tenant has a corresponding namespace.  Alternative models of multi-tenancy with a tenant per cluster, or a tenant per Google Cloud project, are harder to manage.  Network traffic within a namespace is unrestricted. Network traffic between namespaces must be explicitly allowed. These policies can be enforced using Kubernetes [network policy](http://cloud.google.com#network_policies). 

The users of the cluster are divided into three different roles, depending on their privilege:

Cluster administrator

This role is for administrators of the entire cluster, who manage all tenants. Cluster administrators can create, read, update, and delete any policy object. They can create namespaces and assign them to namespace administrators.

Namespace administrator

This role is for administrators of specific, single tenants. A namespace administrator can manage the users in their namespace.


Developer

Members of this role can create, read, update, and delete namespaced non-policy objects like [Pods](http://cloud.google.com/kubernetes-engine/docs/concepts/pod), [Jobs](http://cloud.google.com/kubernetes-engine/docs/how-to/jobs), and [Ingresses](http://cloud.google.com/kubernetes-engine/docs/concepts/ingress). Developers only have
these privileges in the namespaces they have access to.

![](http://cloud.google.com/static/kubernetes-engine/images/enterprise-multitenancy.svg)

For information on setting up multiple multi-tenant clusters for an enterprise organization, see [Best practices for enterprise multi-tenancy](http://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy).

### SaaS provider multi-tenancy

The tenants of a SaaS provider's cluster are the per-customer instances of the application, and the SaaS's control plane. To take advantage of namespace-scoped policies, the application instances should be organized into their own namespaces, as should components of the SaaS's control plane. End users can't interact with the Kubernetes control plane directly, they use the SaaS's interface instead, which in turn interacts with the Kubernetes control plane.

For example, a blogging platform could run on a multi-tenant cluster. In this case, the tenants are each customer's blog instance and the platform's own control plane. The platform's control plane and each hosted blog would all run in separate namespaces. Customers would create and delete blogs, update the blogging software versions through the platform's interface with no visibility into how the cluster operates.

![](http://cloud.google.com/static/kubernetes-engine/images/saas-multitenancy.svg)

## Multi-tenancy policy enforcement

GKE and Kubernetes provide several features that can be used to manage multi-tenant clusters. The following sections give an overview of these features.

### Access control

GKE has two access control systems: Identity and Access Management (IAM) and role-based access control (RBAC).
IAM is Google Cloud's access control system for managing authentication and authorization for GCP resources. You use IAM to grant users access to GKE and Kubernetes resources. RBAC is built into Kubernetes and grants granular permissions for specific resources and operations within your clusters.

Refer to the [Access control overview](http://cloud.google.com/kubernetes-engine/docs/concepts/access-control) for more information about these options and when to use each.

Refer to the [RBAC how-to guide](http://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control) and the [IAM how-to guide](http://cloud.google.com/kubernetes-engine/docs/how-to/iam) to learn how to use these
access control systems.

You can use IAM and RBAC permissions together with namespaces to restrict user
interactions with cluster resources on console. For more information, see
[Enable access and view cluster resources by namespace](http://cloud.google.com/kubernetes-engine/docs/how-to/restrict-resources-access-by-namespace).

### Network policies

Cluster [network policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/) give you control over the communication between your cluster's Pods. Policies specify which namespaces, labels, and IP address ranges a Pod can communicate with.

See the [network policy how-to](http://cloud.google.com/kubernetes-engine/docs/how-to/network-policy)
for instructions on enabling network policy enforcement on GKE.

Follow the [network policy tutorial](http://cloud.google.com/kubernetes-engine/docs/tutorials/network-policy) to learn how to
write network policies.

### Resource quotas

Resource quotas manage the amount of resources used by the objects in a namespace. You can set quotas in terms of CPU and memory usage, or in terms of object counts. Resource quotas let you ensure that no tenant uses more than its assigned share of cluster resources.

Refer to the [resource quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas) documentation for more information.

### Pod security policies

**Warning:** Kubernetes has officially [deprecated PodSecurityPolicy](https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/)
in version 1.21. PodSecurityPolicy will be shut down in version 1.25. For
information about alternatives, refer to [PodSecurityPolicy deprecation](http://cloud.google.com/kubernetes-engine/docs/deprecations/podsecuritypolicy).**Note:** For Autopilot clusters, [Pod security policies](http://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#pod_security_policies)
are not supported.

[PodSecurityPolicies](https://kubernetes.io/docs/concepts/policy/pod-security-policy/)
are a Kubernetes API type that validate requests to create and update Pods.
PodSecurityPolicies define default values and requirements for
security-sensitive fields of Pod specification. You can create policies that
restrict the deployment of Pods that access the host filesystem, networks, PID
namespaces, [volumes](http://cloud.google.com/kubernetes-engine/docs/concepts/volumes), and more.

Refer to the [PodSecurityPolicies how-to](http://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) for more.

### Pod anti-affinity

**Warning:** Pod anti-affinity rules can be circumvented by malicious tenants. The
example below should only be used with clusters with trusted tenants, or with
tenants who don't have direct access to the Kubernetes control plane.

You can use [Pod\
anti-affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity)
to prevent Pods from different tenants from being scheduled on the same node.
Anti-affinity constraints are based on Pod
[labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
For example, the Pod specification below describes a Pod with the label `"team":
"billing"`, and an anti-affinity rule that prevents the Pod from being scheduled
alongside Pods without the label.

```
apiVersion: v1
kind: Pod
metadata:
name: bar
labels:
    team: "billing"
spec:
affinity:
    podAntiAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
      - topologyKey: "kubernetes.io/hostname"
        labelSelector:
          matchExpressions:
          - key: "team"
            operator: NotIn
            values: ["billing"]

```

The drawback to this technique is that malicious users could circumvent the rule by applying the `team: billing` label to an arbitrary Pod. Pod anti-affinity alone cannot securely enforce policy on clusters with untrusted tenants.  Refer to the [Pod anti-affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity) documentation for more information.

### Dedicated nodes with taints and tolerations

**Warning:** Policies enforced by node taints and tolerations can be circumvented by malicious tenants. The example below should only be used with clusters with trusted tenants, or with tenants who don't have direct access to the Kubernetes control plane.

Node taints are another way to control workload scheduling. You can use node taints to reserve specialized nodes for use by certain tenants. For example, you can dedicate [GPU equipped nodes](http://cloud.google.com/kubernetes-engine/docs/how-to/gpus) to the specific tenants whose workloads require GPUs. For Autopilot clusters, node tolerations are supported only for [workload separation](http://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning#workload_separation). 

Node taints are automatically added by node auto-provisioning as needed.

To dedicate a [node pool](http://cloud.google.com/kubernetes-engine/docs/concepts/node-pools) to a certain tenant, apply a taint with `effect: "NoSchedule"` to the node pool. Then only Pods with a corresponding toleration can be scheduled to nodes in the node pool.

The drawback to this technique is that malicious users could add a toleration to their Pods to get access to the dedicated node pool. Node taints and tolerations alone cannot securely enforce policy on clusters with untrusted tenants.

See the [node taints how-to page](http://cloud.google.com/kubernetes-engine/docs/how-to/node-taints) to
learn how to control scheduling with node taints.

## What's next

- Watch the[Kubernetes Multi-tenancy talk](https://www.youtube.com/watch?v=RkY8u1_f5yY) from Google Cloud Next '18.
- Read the[Best practices for enterprise multi-tenancy](http://cloud.google.com/kubernetes-engine/docs/best-practices/enterprise-multitenancy).
- Learn how to[Optimize resource usage in a multi-tenant GKE cluster using node auto-provisioning](http://cloud.google.com/solutions/optimizing-resources-in-multi-tenant-gke-clusters-with-auto-provisioning).


https://cloud.google.com/kubernetes-engine/docs/concepts/multitenancy-overview
