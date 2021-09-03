# What’s new in Kubernetes 1.20?

on December 1, 2020 From: https://sysdig.com/blog/whats-new-kubernetes-1-20/

Table of contents

**Kubernetes 1.20** is about to be released, and it comes packed with novelties! Where do we begin?

As we highlighted in the last release, enhancements [now have to move forward to stability or being deprecated](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1635). As a result, alpha features that have been around since the early times of Kubernetes, like [CronJobs](http://sysdig.com#19) and [Kubelet CRI support](http://sysdig.com#2040), are now getting the attention they deserve.

Another noteworthy fact of this Kubernetes 1.20 release is that it brings 43 enhancements, up from 34 in 1.19. Of those 43 enhancements, 11 are graduating to Stable, 15 are completely new, and 17 are existing features that keep improving.

So many enhancements means that they are smaller in scope. Kubernetes 1.20 is a healthy house cleaning event with a lot of small user-friendly changes. For example, improvements in kube-apiserver to [work better in HA clusters](http://sysdig.com#1965) and [reboot more efficiently](http://sysdig.com#1904) after an upgrade. Or, being able to [gracefully shutdown nodes](http://sysdig.com#2000) so resources can be freed properly. It’s exciting to see small features like these paving the way for the big changes that are to come.

Here is the detailed list of what’s new in Kubernetes 1.20.

## Kubernetes 1.20 – Editor’s pick:

These are the features that look most exciting to us in this release (ymmv):

### [\#1753 Kubernetes system components logs sanitization](http://sysdig.com\#1753)

There have been some vulnerabilities lately related to credentials being [leaked into the log output](https://sysdig.com/blog/falco-cve-2020-8566-ceph/). It’s comforting knowing that they are being approached by keeping the big picture in mind, identifying the potential sources of leaks, and placing a redacting mechanism in place to cut those leaks.

_[Vicente Herrera](https://twitter.com/Vicen_Herrera) – Cloud native security advocate at Sysdig_

### [\#2047 CSIServiceAccountToken](http://sysdig.com\#2047)

This enhancement represents the huge effort to improve the security around authentication and the handling of tokens that is taking place in Kubernetes. Take a look at [the Auth section](http://sysdig.com#auth) in this article to see what I mean. This particular feature makes access to volumes that requires authentication (like secret vaults) more secure and easier to set up.

_[Víctor Jiménez](https://twitter.com/capitangolo/) – Content Marketing Engineer at Sysdig_

### [\#19 CronJobs](http://sysdig.com\#19) \+ [\#2040 Kubelet CRI support](http://sysdig.com\#2040)

You just need to look at the issue number (19) from CronJobs to know these are not new features. CronJobs have been around since 1.4 and CRI support since 1.5, and although used widely in production, neither are considered stable yet. It’s comforting seeing that features you depend on to run your production cluster aren’t alphas anymore.

_David de Torres – Integrations Engineer at Sysdig_

### [\#2000 Graceful node shutdown](http://sysdig.com\#2000)

A small feature? Yes. A huge demonstration of goodwill towards developers? That is also true. Being able to properly release resources when a node shuts down will avoid many weird behaviors, including the hairy ones that are tough to debug and troubleshoot.

_Álvaro Iradier – Integrations Engineer at Sysdig_

### [\#1965 kube-apiserver identity](http://sysdig.com\#1965)

Having a unique identifier for each kube-apiserver instance is one of those features that usually goes unnoticed. However, knowing that it will enable better high availability features in future Kubernetes versions really hypes me up.

_[Mateo Burillo](https://twitter.com/mateobur) – Product Manager at Sysdig_

### [\#1748 Expose metrics about resource requests and limits that represent the pod model](http://sysdig.com\#1748)

Yes! More metrics are always welcome. Even more so when they will enable you to better plan the capacity of your cluster and help troubleshoot eviction problems.

_[Carlos Arilla](https://twitter.com/carillan_) – Integrations Engineering Manager at Sysdig_

## Deprecations in Kubernetes 1.20

### [\#1558](https://github.com/kubernetes/enhancements/issues/1558) Streaming proxy redirects

**Stage:** Deprecation

**Feature group:** node

As part of deprecating `StreamingProxyRedirects`, the `--redirect-container-streaming` flag can no longer be enabled.

Both the `StreamingProxyRedirects` feature gate and the `--redirect-container-streaming` flag were marked as deprecated on 1.18. The feature gate will be disabled by default in 1.22, and both will be removed on 1.24.

You can read the rationale behind this deprecation [in the KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20191205-container-streaming-requests.md).

### [\#2067](https://github.com/kubernetes/enhancements/issues/2067) Rename the kubeadm master node-role label and taint

**Stage:** Deprecation

**Feature group:** cluster-lifecycle

In an effort to move away from offensive wording, `node-role.kubernetes.io/master` is being renamed to `node-role.kubernetes.io/control-plane`.

You can read more about the implications of this change [in the KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/kubeadm/2067-rename-master-label-taint).

### [\#1164](https://github.com/kubernetes/enhancements/issues/1164) Deprecate and remove SelfLink

**Stage:** Graduating to Beta

**Feature group:** api-machinery

The field `SelfLink` is present in every Kubernetes object and contains a URL representing the given object. This field does not provide any new information and its creation and maintenance has a performance impact.

Initially deprecated in [Kubernetes 1.16](https://sysdig.com/blog/whats-new-kubernetes-1-16/#1164), the feature gate is now disabled by default and will finally be removed in Kubernetes 1.21.

## Kubernetes 1.20 API

### [\#1965](https://github.com/kubernetes/enhancements/issues/1965) kube-apiserver identity

**Stage:** Alpha

**Feature group:** api-machinery

In order to better control which _kube-apiservers_ are alive in a high availability cluster, a new lease / heartbeat system has been implemented.

Each `kube-apiserver` will assign a unique ID to itself in the format of `kube-apiserver-{UUID}`. These IDs will be stored in Lease objects that will be refreshed every 10 seconds (by default), and garbage collected if not renewed.

This system is similar to the node heartbeat. In fact, it will be reusing the existing Kubelet heartbeat logic.

### [\#1904](https://github.com/kubernetes/enhancements/issues/1904) Efficient watch resumption after kube-apiserver reboot

**Stage:** Alpha

**Feature group:** api-machinery, storage

From now on, kube-apiserver can initialize its _watch cache_ faster after a reboot. This feature can be enabled with the `EfficientWatchResumption` feature flag.

Clients can keep track of the changes in Kubernetes objects with a watch. To serve these requests, `kube-apiserver` keeps a watch cache. But after a reboot (e.g., during rolling upgrades) that cache is often outdated, so `kube-apiserver` needs to fetch the updated state from etcd.

The updated implementation will leverage the `WithProgressNotify` option from etcd version 3.0. `WithProgressNotify` will enable getting the updated state for the objects every minute, and also once before shutting down. This way, when `kube-apiserver` restarts, its cache will already be fairly updated.

Check the full implementation details [in the KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1904-efficient-watch-resumption).

### [\#1929](https://github.com/kubernetes/enhancements/issues/1929) Built-in API types defaults

**Stage:** Graduating to Stable

**Feature group:** api-machinery

When building custom resource definitions (CRDs) as Go structures, you will now be able to use the `// +default` marker to provide default values. These markers will be translated into OpenAPI `default` fields.

```
type Foo struct {
// +default=32
Integer int
// +default="bar"
String string
// +default=["popcorn", "chips"]
StringList []string
}

```

### [\#1040](https://github.com/kubernetes/enhancements/issues/1040) Priority and fairness for API server requests

**Stage:** Graduating to Beta

**Feature group:** api-machinery

The `APIPriorityAndFairness` feature gate enables a new max-in-flight request handler in the API server. By defining different types of requests with `FlowSchema` objects and assigning them resources with `RequestPriority` objects, you can ensure the Kubernetes API server will be responsive for admin and maintenance tasks during high loads.

Read more on the release for 1.18 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1040) series.

## Auth in Kubernetes 1.20

### [\#541](https://github.com/kubernetes/enhancements/issues/541) External client-go credential providers

**Stage:** Graduating to Beta

**Feature group:** auth

This enhancement allows Go clients to authenticate using external credential providers, like Key Management Systems (KMS), Trusted Platform Modules (TPM), or Hardware Security Modules (HSM).

Those devices are already used to authenticate against other services, are easier to rotate, and are more secure as they don’t exist as files on the disk.

Initially introduced on Kubernetes 1.10, this feature finally reaches Beta status.

### [\#542](https://github.com/kubernetes/enhancements/issues/542) TokenRequest API and Kubelet integration

**Stage:** Graduating to Beta

**Feature group:** auth

The current JSON Web Tokens (JWT) that workloads use to authenticate against the API have some security issues. This enhancement comprises the work to create a more secure API for JWT.

In the new API, tokens can be bound to a specific workload, given a validity duration, and exist only while a given object exists.

Check the full details [in the KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/1205-bound-service-account-tokens).

### [\#1393](https://github.com/kubernetes/enhancements/issues/1393) Provide OIDC discovery for service account token issuer

**Stage:** Graduating to Beta

**Feature group:** auth

Kubernetes service accounts (KSA) can currently use JSON Web Tokens (JWT) to authenticate against the Kubernetes API, using `kubectl --token <the_token_string> ` for example. This enhancement allows services outside the cluster to use these tokens as a general authentication method without overloading the API server.

Read more on the release for 1.18 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1393) series.

## Kubernetes 1.20 CLI

### [\#1020](https://github.com/kubernetes/enhancements/issues/1020) Moving kubectl package code to staging

**Stage:** Graduating to Stable

**Feature group:** cli

Continuing the work done in [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1020), this internal restructuring of the `kubectl` code is the first step to move the kubectl binary into [its own repository](https://github.com/kubernetes/kubectl). This helped decouple kubectl from the kubernetes code base and made it easier for out-the-tree projects to reuse its code.

### [\#1441](https://github.com/kubernetes/enhancements/issues/1441) kubectl debug

**Stage:** Graduating to Beta

**Feature group:** cli

There have been two major changes in `kubectl debug` since [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1441).

You can now use `kubectl debug` instead of `kubectl alpha debug`.

Also, `kubectl debug` can now change container images when copying a pod for debugging, similar to how `kubectl set image` works.

## Cloud providers in Kubernetes 1.20

### [\#2133](https://github.com/kubernetes/enhancements/issues/2133) Kubelet credential provider

**Stage:** Alpha

**Feature group:** cloud-provider

This enhancement comprises the work done to move cloud provider specific SDKs out-of-tree. In particular, it replaces in-tree container image registry credential providers with a new mechanism that is external and pluggable.

You can get more information on how to build a credential provider [in the KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cloud-provider/20191004-out-of-tree-credential-providers.md).

### [\#667](https://github.com/kubernetes/enhancements/issues/667) Support out-of-tree Azure cloud provider

**Stage:** Graduating to Beta

**Feature group:** cloud-provider

This enhancement contains the work done to move the Azure cloud provider out-of-tree while keeping feature parity with the existing code in kube-controller-manager.

## Kubernetes 1.20 instrumentation

### [\#1748](https://github.com/kubernetes/enhancements/issues/1748) Expose metrics about resource requests and limits that represent the pod model

**Stage:** Alpha

**Feature group:** instrumentation

The `kube-scheduler` now exposes more metrics on the requested resources and the desired limits of all running pods. This will help cluster administrators better plan capacity and triage errors.

The metrics are exposed at the HTTP endpoint `/metrics/resources` when you use the `--show-hidden-metrics-for-version=1.20` flag.

### [\#1753](https://github.com/kubernetes/enhancements/issues/1753) Kubernetes system components logs sanitization

**Stage:** Alpha

**Feature group:** instrumentation

This enhancement aims to avoid sensitive data like passwords and tokens from being leaked in the Kubernetes log.

Sensitive fields have been tagged:

```
type ConfigMap struct {
Data map[string]string `json:"data,omitempty" datapolicy:"password,token,security-key"`
}

```

So that they can be redacted in the log output.

The sanitization filter kicks in when the `--experimental-logging-sanitization` flag is used. However, be aware that the current implementation takes a noticeable performance hit, and that sensitive data can still be leaked by user workloads.

### [\#1933](https://github.com/kubernetes/enhancements/issues/1933) Defend against logging secrets via static analysis

**Stage:** Alpha

**Feature group:** instrumentation

While preparing the previous enhancement, [#1753](http://sysdig.com#1753), _go-flow-levee_ was used to provide insight on where sensitive data is used.

## Network in Kubernetes 1.20

### [\#1435](https://github.com/kubernetes/enhancements/issues/1435) Support of mixed protocols in Services with type=LoadBalancer

**Stage:** Alpha

**Feature group:** network

The current LoadBalancer Service implementation does not allow different protocols under the same port (UDP, TCP). The rationale behind this is to avoid negative surprises with the load balancer bills in some cloud implementations.

However, a user might want to serve both UDP and TCP requests for a DNS or SIP server on the same port.

This enhancement comprises the work to remove this limitation, and investigate the effects on billing on different clouds.

### [\#1864](https://github.com/kubernetes/enhancements/issues/1864) Optionally Disable NodePorts for Service Type=LoadBalancer

**Stage:** Alpha

**Feature group:** network

Some implementations of the LoadBalancer API do not consume the node ports automatically allocated by Kubernetes, like MetalLB or kube-router. However, the API requires a port to be defined and allocated.

As a result, the number of load balancers is limited to the number of available ports. Also, these allocated but unused ports are exposed, which impact compliance requirements.

To solve this, the field `allocateLoadBalancerNodePort` has been added to `Service.Spec`. When set to `true`, it behaves as usual; when set to `false`, it will stop allocating new node ports (it won’t deallocate the existing ones, however).

### [\#1672](https://github.com/kubernetes/enhancements/issues/1672) Tracking terminating Endpoints

**Stage:** Alpha

**Feature group:** network

When Kubelet starts a graceful shutdown, the shutting-down Pod is removed from Endpoints (and, if enabled, EndpointSlice). If a consumer of the EndpointSlice API wants to know what Pods are terminating, it will have to watch the Pods directly. This complicates things, and has some scalability implications.

With this enhancement, a new `Terminating` field has been added to the EndpointConditions struct. Thus, terminating pods can be kept in the EndpointSlice API.

In the long term, this will enable services to handle terminating pods more intelligently. For example, IPVS proxier could set the weight of an endpoint to `0` during termination, instead of guessing that information from the connection tracking table.

### [\#614](https://github.com/kubernetes/enhancements/issues/614) SCTP support for Services, Pod, Endpoint, and NetworkPolicy

**Stage:** Graduating to Stable

**Feature group:** network

SCTP is now supported as an additional protocol alongside TCP and UDP in Pod, Service, Endpoint, and NetworkPolicy.

Introduced in [Kubernetes 1.12](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#614httpsgithubcomkubernetesfeaturesissues614sctpsupportforservicespodendpointandnetworkpolicyalpha), this feature finally graduates to _Stable_.

### [\#1507](https://github.com/kubernetes/enhancements/issues/1507) Adding AppProtocol to Services and Endpoints

**Stage:** Graduating to Stable

**Feature group:** network

The [EndpointSlice API](https://sysdig.com/blog/whats-new-kubernetes-1-16/#752) added a new AppProtocol field in Kubernetes 1.17 to allow application protocols to be specified for each port. This enhancement brings that field into the `ServicePort` and `EndpointPort` resources, replacing non-standard annotations that are causing a [bad user experience](https://github.com/kubernetes/kubernetes/issues/40244).

Initially introduced in [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1507), this enhancement now graduates to _Stable_.

### [\#752](https://github.com/kubernetes/enhancements/issues/752) EndpointSlice API

**Stage:** Graduating to Beta

**Feature group:** network

The new `EndpointSlice` API will split endpoints into several Endpoint Slice resources. This solves many problems in the current API that are related to big `Endpoints` objects. This new API is also designed to support other future features, like multiple IPs per pod.

In Kubernetes 1.20, `Topology` has been deprecated, and a new `NodeName` field has been added. Check [the roll out plan](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/0752-endpointslices#roll-out-plan) for more info.

Read more on the release for 1.16 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#752) series.

### [\#563](https://github.com/kubernetes/enhancements/issues/563) Add IPv4/IPv6 dual-stack support

**Stage:** Graduating to Alpha

**Feature group:** network

This feature summarizes the work done to natively support dual-stack mode in your cluster, so you can assign both IPv4 and IPv6 addresses to a given pod.

Read more on the release for 1.16 in the “ [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#563)” series.

A big overhaul of this feature has taken place in Kubernetes 1.20 with breaking changes, which is why this feature is still on Alpha.

The major user-facing change is the introduction of the `.spec.ipFamilyPolicy` field. It can be set to: `SingleStack`, `PreferDualStack`, and `RequireDualStack`.

Dual stack is a big project that has implications on many Kubernetes services, so expect new improvements and changes before it leaves the alpha stage.

## Kubernetes 1.20 nodes

### [\#1967](https://github.com/kubernetes/enhancements/issues/1967) Support to size memory-backed volumes

**Stage:** Alpha

**Feature group:** node

When a Pod defines a memory-backed empty dir volume, (e.g., tmpfs) not all hosts size this volume equally. For example, a Linux host sizes it to 50% of the memory on the host.

This has implications at several levels. For example, the Pod is less portable, as it’s behavior depends on the host it’s being deployed on. Also, this memory usage is not transparent to the eviction heuristics.

This new enhancement, after enabling the `SizeMemoryBackedVolumes` feature gate, will size these volumes not only with the node allocatable memory in mind, but also with the pod allocatable memory and the `emptyDir.sizeLimit` field.

### [\#1972](https://github.com/kubernetes/enhancements/issues/1972) Fixing Kubelet exec probe timeouts

**Stage:** Stable

**Feature group:** node

Now, exec probes in Kubelet will respect the `timeoutSeconds` field.

Since this is a bugfix, this feature graduates directly to Stable. You can roll back to the old behavior by disabling the `ExecProbeTimeout` feature gate.

### [\#2000](https://github.com/kubernetes/enhancements/issues/2000) Graceful node shutdown

**Stage:** Alpha

**Feature group:** node

With the `GracefulNodeShutdown` feature gate enabled, Kubelet will try to gracefully terminate the pods running in the node when shutting down. The implementation works by listening to systemd inhibitor locks (for Linux).

The new Kubelet config setting `kubeletConfig.ShutdownGracePeriod` will dictate how much time Pods have to terminate gracefully.

This enhancement can mitigate issues in some workloads, making life easier for admins and developers. For example, a cold shutdown can corrupt open files, or leave resources in an undesired state.

### [\#2053](https://github.com/kubernetes/enhancements/issues/2053) Add downward API support for hugepages

**Stage:** Alpha

**Feature group:** node

If enabled via the `DownwardAPIHugePages` feature gate, Pods will be able to fetch information on their hugepage requests and limits via the downward API. This keeps things consistent with other resources like cpu, memory, and ephemeral-storage.

### [\#585](https://github.com/kubernetes/enhancements/issues/585) RuntimeClass

**Stage:** Graduating to Stable

**Feature group:** node

The `RuntimeClass` resource provides a mechanism for supporting multiple runtimes in a cluster (Docker, rkt, gVisor, etc.), and surfaces information about that container runtime to the control plane.

For example:

```
apiVersion: v1
kind: Pod
metadata:
name: mypod
spec:
runtimeClassName: sandboxed
# ...

```

This YAML will instruct the Kubelet to use the `sandboxed` RuntimeClass to run this pod. This selector is required to dynamically adapt to multiple container runtime engines beyond Docker, like rkt or gVisor.

Introduced on [Kubernetes 1.12](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#585httpsgithubcomkubernetesfeaturesissues585runtimeclassalpha), this feature finally graduates to Stable.

### [\#606](https://github.com/kubernetes/enhancements/issues/606) Support 3rd party device monitoring plugins

**Stage:** Graduating to Stable

**Feature group:** node

This feature allows the Kubelet to expose container bindings to third-party monitoring plugins.

With this implementation, administrators will be able to monitor the custom resource assignment to containers using a third-party Device Monitoring Agent (For example, percent GPU use per pod).

Read more on the release for 1.15 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-15/#606) series.

### [\#757](https://github.com/kubernetes/enhancements/issues/757) PID limiting

**Stage:** Graduating to Stable

**Feature group:** node

PIDs are a fundamental resource on any host. Administrators require mechanisms to ensure that user pods cannot induce pid exhaustion that may prevent host daemons (runtime, Kubelet, etc.) from running.

This feature allows for the configuration of a Kubelet to limit the number of PIDs a given pod can consume. Node-level support for PID limiting no longer requires setting the feature gate `SupportNodePidsLimit=true` explicitly.

Read more on the release for 1.15 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-15/#757) series, and [in the Kubernetes blog](https://kubernetes.io/blog/2019/04/15/process-id-limiting-for-stability-improvements-in-kubernetes-1.14/).

### [\#950](https://github.com/kubernetes/enhancements/issues/950) Add pod-startup liveness-probe holdoff for slow-starting pods

**Stage:** Graduating to Stable

**Feature group:** node

Probes allow Kubernetes to monitor the status of your applications. If a pod takes too long to start, those probes might think the pod is dead, causing a restart loop. This feature lets you define a `startupProbe` that will hold off all of the other probes until the pod finishes its startup. For example: _“Don’t test for liveness until a given HTTP endpoint is available_. _“_

Read more on the release for 1.16 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#950) series.

### [\#693](https://github.com/kubernetes/enhancements/issues/693) Node topology manager

**Stage:** Graduating to Beta

**Feature group:** node

Machine learning, scientific computing, and financial services are examples of systems that are computational intensive and require ultra-low latency. These kinds of workloads benefit from isolated processes to one CPU core rather than jumping between cores or sharing time with other processes.

The node topology manager is a `kubelet` component that centralizes the coordination of hardware resource assignments. The current approach divides this task between several components (CPU manager, device manager, and CNI), which sometimes results in unoptimized allocations.

In Kubernetes 1.20, you can define the `scope` where topology hints should be collected on a `container`-by-container basis, or on a `pod`-by-pod basis.

Read more on the release for 1.16 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#693) series.

### [\#1797](https://github.com/kubernetes/enhancements/issues/1797) Allow users to set a pod’s hostname to its Fully Qualified Domain Name (FQDN)

**Stage:** Graduating to Beta

**Feature group:** node

Now, it’s possible to set a pod’s hostname to its Fully Qualified Domain Name (FQDN), which increases the interoperability of Kubernetes with legacy applications.

After setting `hostnameFQDN: true`, running `uname -n` inside a Pod returns `foo.test.bar.svc.cluster.local` instead of just `foo`.

This feature was introduced on [Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1797), and you can read more details [in the enhancement proposal](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/1797-configure-fqdn-as-hostname-for-pods/README.md).

### [\#1867](https://github.com/kubernetes/enhancements/issues/1867) Kubelet feature: Disable AcceleratorUsage metrics

**Stage:** Graduating to Beta

**Feature group:** node

With [#606 (Support third-party device monitoring plugins)](https://sysdig.com/blog/whats-new-kubernetes-1-15/#606) and the PodResources API about to enter GA, it isn’t expected for Kubelet to gather metrics anymore.

Introduced on [Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1867), this enhancement summarizes the process to deprecate Kubelet collecting those Accelerator Metrics.

### [\#2040](https://github.com/kubernetes/enhancements/issues/2040) Kubelet CRI support

**Stage:** Graduating to Beta

**Feature group:** node

Kubernetes introduced support for the Container Runtime Interface (CRI) [as soon as Kubernetes 1.5](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/). It is a plugin interface that enables Kubelet to use a wide variety of container runtimes, without the need to recompile. This includes alternatives to Docker like CRI-O or containerd.

Although the CRI API has been widely tested in production, it was still officially in Alpha.

This enhancement comprises the work to identify and close the remaining gaps, so that CRI can finally be promoted to Stable.

## Kubernetes 1.20 storage

### [\#2047](https://github.com/kubernetes/enhancements/issues/2047) CSIServiceAccountToken

**Stage:** Alpha

**Feature group:** storage

CSI drivers [can impersonate the pods](https://kubernetes-csi.github.io/docs/token-requests.html) that they mount volumes for. This feature increases security by providing the CSI drivers with only the permissions they need. However, in the current implementation, the drivers read the service account tokens directly from the filesystem.

Some downsides of this are that the drivers need permissions to read the filesystem, which might give access to more secrets than needed, and that the token is not guaranteed to be available (e.g., `automountServiceAccountToken=false`).

With this enhancement, CSI drivers will be able to request the service account tokens from Kubelet to the `NodePublishVolume` function. Kubelet will also be able to limit what tokens are available to which driver. And finally, the driver will be able to re-execute `NodePublishVolume` to remount the volume by setting `RequiresRepublish` to `true`.

This last feature will come in handy when the mounted volumes can expire and need a re-login. For example, a secrets vault.

Check the full details [in the KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1855-csi-driver-service-account-token/README.md).

### [\#177](https://github.com/kubernetes/enhancements/issues/177) Snapshot / restore volume support for Kubernetes

**Stage:** Graduating to Stable

**Feature group:** storage

In alpha, since [the 1.12 Kubernetes release](https://sysdig.com/blog/whats-new-in-kubernetes-1-12/#177httpsgithubcomkubernetesfeaturesissues177snapshotrestorevolumesupportforkubernetescrdexternalcontrolleralpha), this feature finally graduates to Stable.

Similar to how API resources `PersistentVolume` and `PersistentVolumeClaim` are used to provision volumes for users and administrators, `VolumeSnapshotContent` and `VolumeSnapshot` API resources can be provided to create volume snapshots for users and administrators. Read more [about volume snapshots here](https://github.com/xing-yang/website/blob/f53fe7ed8bf0ee98a1c45eb67bc505e309fdce0e/content/en/docs/concepts/storage/volume-snapshots.md).

Read more on the release for 1.16 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-16/#693) series.

### [\#695](https://github.com/kubernetes/enhancements/issues/695) Skip volume ownership change

**Stage:** Graduating to Beta

**Feature group:** storage

Before a volume is bind-mounted inside a container, all of its file permissions are changed to the provided `fsGroup` value. This ends up being a slow process on very large volumes, and also breaks some permission sensitive applications, like databases.

Introduced in [Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/#695), the new `FSGroupChangePolicy` field has been added [to control this behavior](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/20200120-skip-permission-change.md). If set to Always, it will maintain the current behavior. However, when set to OnRootMismatch, it will only change the volume permissions if the top level directory does not match the expected `fsGroup` value.

### [\#1682](https://github.com/kubernetes/enhancements/issues/1682) Allow CSI drivers to opt-in to volume ownership change

**Stage:** Graduating to Beta

**Feature group:** storage

Not all volumes support the `fsGroup` operations mentioned in the previous enhancement ( [#695](http://sysdig.com#695)), like NFS. This can result in errors reported to the user.

This enhancement adds a new field called `CSIDriver.Spec.SupportsFSGroup` that allows the driver to define if it supports volume ownership modifications via fsGroup.

Read more on the release for 1.19 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1682) series.

## Miscellaneous

### [\#1610](https://github.com/kubernetes/enhancements/issues/1610) Container resource based Pod autoscaling

**Stage:** Alpha

**Feature group:** autoscaling

The current Horizontal Pod Autoscaler (HPA) can scale workloads based on the resources used by their Pods. This is the aggregated usage from all of the containers in a Pod.

[This feature](https://github.com/kubernetes/enhancements/blob/master/keps/sig-autoscaling/0001-container-resource-autoscaling.md) allows the HPA to scale those workloads based on the resource usage of individual containers:

```
type: ContainerResource
containerResource:
name: cpu
container: application
target:
     type: Utilization
     averageUtilization: 60

```

### [\#1001](https://github.com/kubernetes/enhancements/issues/1001) Support CRI-ContainerD on Windows

**Stage:** Graduating to Stable

**Feature group:** windows

ContainerD is an OCI-compliant runtime that works with Kubernetes and has support for the host container service (HCS v2) in Windows Server 2019. This enhancement introduces ContainerD 1.3 support in Windows as a Container Runtime Interface (CRI).

Read more on the release for 1.18 in the [What’s new in Kubernetes](https://sysdig.com/blog/whats-new-kubernetes-1-18/#1001) series.

### [\#19](https://github.com/kubernetes/enhancements/issues/19) CronJobs (previously ScheduledJobs)

**Stage:** Graduating to Beta

**Feature group:** apps

Introduced in Kubernetes 1.4 and in beta since 1.8, CronJobs are finally on the road to become Stable.

CronJobs runs periodic tasks in a Kubernetes cluster, similar to cron on UNIX-like systems.

A new, alternate implementation of CronJobs [has been created](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/19-Graduate-CronJob-to-Stable) to address the main issues of the current code without breaking the current functionality.

This new implementation will focus on scalability, providing more status information and addressing the current open issues.

You can start testing the new CronJobs by setting the `CronJobControllerV2` feature flag to `true`.

### [\#1258](https://github.com/kubernetes/enhancements/issues/1258) Add a configurable default constraint to PodTopologySpread

**Stage:** Graduating to Beta

**Feature group:** scheduling

In order to take advantage of [even pod spreading](https://docs.google.com/document/d/1ZDSHeySKoYYKnP_86rj2d5GRvYWn3QPKOu1XDWOJAeI/edit#895), each pod needs its own spreading rules and this can be a tedious task.

Introduced in [Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/#1258), this enhancement allows you to define global `defaultConstraints` that will be applied at cluster level to all of the pods that don’t define their own `topologySpreadConstraints`.

* * *

That’s all, folks! Exciting as always; get ready to upgrade your clusters if you are intending to use any of these features.

If you liked this, you might want to check out our previous ‘What’s new in Kubernetes’ editions:

- [What’s new in Kubernetes 1.19](https://sysdig.com/blog/whats-new-kubernetes-1-19/)
- [What’s new in Kubernetes 1.18](https://sysdig.com/blog/whats-new-kubernetes-1-18/)
- [What’s new in Kubernetes 1.17](https://sysdig.com/blog/whats-new-kubernetes-1-17/)
- [What’s new in Kubernetes 1.16](https://sysdig.com/blog/whats-new-kubernetes-1-16/)
- [What’s new in Kubernetes 1.15](https://sysdig.com/blog/whats-new-kubernetes-1-15/)
- [What’s new in Kubernetes 1.14](https://sysdig.com/blog/whats-new-kubernetes-1-14/)
- [What’s new in Kubernetes 1.13](https://sysdig.com/blog/whats-new-in-kubernetes-1-13)
- [What’s new in Kubernetes 1.12](https://sysdig.com/blog/whats-new-in-kubernetes-1-12)

And if you enjoy keeping up to date with the Kubernetes ecosystem, [subscribe to our container newsletter](https://go.sysdig.com/container-newsletter-signup.html), a monthly email with the coolest stuff happening in the cloud-native ecosystem.