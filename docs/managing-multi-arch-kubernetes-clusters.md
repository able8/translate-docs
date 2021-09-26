# Getting up and running with multi-arch Kubernetes clusters

*February 20, 2021*

The world of ARM processors has been getting very interesting over the last few years. Until fairly recently, for most people, ARM CPUs were reserved for their phone or maybe a [Raspberry Pi running their home DNS](https://pi-hole.net/). However now the Raspberry Pi 4 has a pretty decent quad-core CPU and up to 8GB RAM, Apple have blown away the industry with the M1 chips and AWS have launched [Graviton2](https://aws.amazon.com/ec2/graviton/) instances which depending on who you ask have 20-40% better price/performance than the Intel equivalents.

Despite all this, until recently I wasn’t convinced trying to use arm64 nodes in a production Kubernetes cluster was worth the effort. However the combination of all these events has caused a lot of the software to catch up and support for arm64 in Kubernetes distributions from Amazon EKS to k3s is now excellent! It is likely that a good proportion of the container images you’re using in production now support “multi-arch” and will run on both arm64 and amd64 machines.

## Container Images

Let’s start with the container images, as that’s what you’ll need to tackle first if you want to start using arm64 in your environments. Historically if you wanted to support multiple CPU architectures for your container image you would have to have two separate tags e.g. myimage:1.2.3-amd64 and myimage:1.2.3-arm64. This was both a pain for people building the images and for people trying to use them; if you’re deploying a Helm Chart for example you would have to override the image tags if you wanted to run on ARM and there was no hope for you if you wanted to use mixture of arm64 and amd64 nodes in a cluster and have the same pods seamlessly schedule on either (without some extra tooling).

This problem has gone away with Manifest Lists in the [Image Manifest V2 spec](https://docs.docker.com/registry/spec/manifest-v2-2/). This allows you to specify a list of container images for a number of different architectures in a single “Manifest List”. In newer container runtime versions (like Docker) this means if you do a `docker pull nginx:alpine` on your Raspberry Pi you’ll get an image for arm64 and on your Intel laptop you’ll get an amd64 image without any further effort. Previously you would have got the godforsaken “exec format error”.

You might think that this means we’ve reached multi-arch nirvana and everything will “just work” now, but unfortunately this is not the case. If you’re using an existing public image, you will need to make sure that it is using a manifest list and supports both amd64 and arm64. Some registries make this really easy, such as Docker Hub, where you’ll get a nice list of architectures on the Tags tab.

However some don’t make it obvious at all! The easiest way I’ve found so far to determine if an image is multi-arch, is to use the experimental [docker manifest command](https://docs.docker.com/engine/reference/commandline/manifest/) in the latest Docker versions. As it says in the documentation you will have to enable experimental features in `~/.docker/config.json` then you will be able to run a command like:

```
[~]$ docker manifest inspect nginx:alpine
{
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.list.v2+json",
   "manifests": [
      {
         "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
         "size": 1568,
         "digest": "sha256:9a39c77d9ea3a9ddc41535f875b7610a0f121df3c2496c16f2a3a5fcb0e43e4f",
         "platform": {
            "architecture": "amd64",
            "os": "linux"
         }
      },
      {
         "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
         "size": 1568,
         "digest": "sha256:22d2c4a5220232818a0fe7a5d3651c846bc3e7d2ff8dbfc2f665c717f0e43a69",
         "platform": {
            "architecture": "arm64",
            "os": "linux",
            "variant": "v8"
         }
      },
```

As you can see the nginx:alpine image is a “manifest list” rather than a plain old manifest and supports the two architectures we’re after. Great! However distroless-java is still single architecture:

```
[~]$ docker manifest inspect gcr.io/distroless/java
{
	"schemaVersion": 2,
	"mediaType": "application/vnd.docker.distribution.manifest.v2+json",
	"config": {
		"mediaType": "application/vnd.docker.container.image.v1+json",
		"size": 1164,
		"digest": "sha256:85cdcf63cad1cfe5373c68f78f21f0c6349fee87fbb40bc9a9dc7d560f52438b"
	},
	"layers": [
		{
			"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
			"size": 643695,
			"digest": "sha256:9e4425256ce4503b2a009683be02372ee51d411e9cc3547919e064fee4970eab"
		},
...
```

I won’t go into building these multi-architecture manifest lists here, as it very much depends on which tool you’re using. If you’re using Docker to build your images you can use the [docker buildx](https://docs.docker.com/buildx/working-with-buildx/)experimental feature. We’re currently using the [jib-maven-plugin](https://github.com/GoogleContainerTools/jib/tree/master/jib-maven-plugin#extended-usage) for our Java-based apps, which has recently added the `platforms` feature. There are many other ways to do it and GitHub actions you can use, so it’s not too hard, and you don’t need an arm64 machine to build an arm64-compatible image anymore thanks to integration with QEMU.

## Kubernetes

Getting arm64 nodes in your cluster is as simple as just creating an extra autoscaling group (or one per Availability Zone) in AWS or using [k3sup](https://github.com/alexellis/k3sup) to join your Raspberry Pi 4 to your k3s cluster. I’ll try to keep this post generic to however you deploy Kubernetes, but as we’re using EKS in production, [here is the documentation](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html#arm-ami) on getting the right AMI for your architecture. I also have a k3s cluster at home with two old Intel laptops and a Raspberry Pi 4, and this guide works exactly the same on that.

Before you add arm64 nodes to your cluster you must consider whether you want to specifically exclude incompatible Pods from running on these nodes with [node affinity rules](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes-using-node-affinity/), or if you want to exclude all Pods by default and specifically allow the ones you know work on arm64 with [taints and tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/). I went with the latter on our EKS clusters because the majority of our workloads aren’t yet multi-arch, but on my home Raspberry Pi cluster I went with the former.

### Node Affinity

The node affinity rule option requires no configuration of the nodes themselves because recent Kubernetes versions have a standard `kubernetes.io/arch` node label; this should be set to arm64 or amd64. However it will require a lot of effort if you have a large number of workloads which don’t yet work on arm64.

```
[~]$ kubectl describe no myarmnode
Name:               myarmnode
Roles:              control-plane,etcd,master
Labels:             kubernetes.io/arch=arm64
                    kubernetes.io/hostname=myarmnode
                    kubernetes.io/os=linux
```

All you need to do is set up a [node affinity rule](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes-using-node-affinity/) in the PodSpec of any Pods which don’t have multi-arch images like this:

```
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
        - matchExpressions:
          - key: "kubernetes.io/arch"
            operator: In
            values: ["amd64"]
```

You could also do the opposite and use the `NotIn` operator and values set as `["arm64"]`.

### Taints and Tolerations

The option I went with in our EKS clusters is to set up a `NoSchedule` “taint” on all arm64 nodes which we then “tolerate” on any pods we know to work on arm64. I’ve done this in the user-data script in the Launch Template used by our Graviton2 Auto Scaling Groups through the `--kubelet-extra-args` flag of [bootstrap.sh](https://aws.amazon.com/blogs/opensource/improvements-eks-worker-node-provisioning/); the extra arg you need to pass to the kubelet is `--register-with-taints="arch=arm64:NoSchedule"`. You can also just use this command after your nodes are registered to the cluster: `kubectl taint no myarmnode arch=arm64:NoSchedule`.

Once all your arm64 nodes are tainted, only Pods with the right [tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) will be scheduled to run on them. As with affinity rules, they are specified in the PodSpec of your Pod. In our case the toleration needed will be:

```
tolerations:
- key: "arch"
  operator: "Equal"
  value: "arm64"
  effect: "NoSchedule"
```

You might see some system-level DaemonSets like kube-proxy already have tolerations like this, which will also work:

```
tolerations:
- effect: NoSchedule
  operator: Exists
```

If you get issues with certain DaemonSets not being scheduled on your arm nodes, even though they have the right tolerations, check that the affinity rules don’t exclude nodes by the `kubernetes.io/arch` label.

### Finding multi-arch images

The first Pods I had to make sure were running on our arm64 nodes were the DaemonSets for things like fluentd (our logging agent) or jaeger-agent (our tracing agent). Unfortunately neither of these images were multi-arch which would have been a blocker for running our workloads on arm64 nodes. However as is often the case in the Open Source world, somebody had already had the same problem and a bit of searching GitHub showed up an open Pull Request or Issue with links to un-official images; [this in the case of fluentd](https://github.com/kokuwaio/helm-charts/issues/12) and [this in the case of jaeger-agent](https://github.com/querycap/jaeger). This is course not as good as the project providing official multi-arch images (like everything except kube-state-metrics in the [kube-prometheus-stack Helm Chart](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)) but in more and more cases official multi-arch images are available and this will only get better over time.

### Cluster autoscaler

After I’d got a number of our workloads set up to run on either arm64 or amd64 nodes, I of course wanted to run as many arm64 nodes as possible due to the lower cost. I did this using the [Priority based expander for cluster-autoscaler](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/expander/priority/readme.md) by setting the Auto Scaling Groups for arm64 nodes to a higher priority than the rest of our groups. You could also set up `preferredDuringSchedulingIgnoredDuringExecution` affinity rules on arm64 compatible Pods but we found the Cluster Autoscaler configuration to be sufficient.

### Bash one-liner for good measure

This will (on Linux at least) list out all the arm64 compatible images running in your Kubernetes cluster:

```
1 kubectl get po -A -o yaml | grep 'image:' | cut -f2- -d':' | sed 's/^[[:space:]]*//g' | grep '/' | sort -u | xargs -I{} bash -c "docker manifest inspect {} | grep -q arm64 && echo {}" 
```

Change `grep -q` to `grep -vq` to invert the logic and return images which won’t work on arm64 nodes.

This would definitely be nicer using [yq](https://github.com/mikefarah/yq) rather than parsing YAML with grep, cut and sed…but I know most people don’t have yq installed.

## Final thoughts

It is still a bit of effort to add arm64 nodes to your Kubernetes clusters. However for us it was worth the effort for the cost savings and it’s always exciting to be an early adopter. If anyone has any questions on how I’ve got things set up or would like to share different approaches I’m [on Twitter](https://twitter.com/cablespaghetti) and love to hear from other people working on similar challenges!
