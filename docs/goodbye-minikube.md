# Goodbye minikube

 Mar 7, 2021 / [KUBERNETES](https://blog.frankel.ch/tag/kubernetes/), [MINIKUBE](https://blog.frankel.ch/tag/minikube/), [KIND](https://blog.frankel.ch/tag/kind/) 
  

I’ve been using [minikube](https://minikube.sigs.k8s.io/) as my local cluster since I started to learn [Kubernetes](https://kubernetes.io/). But I’ve decided to let it go in favor of [kind](https://kind.sigs.k8s.io/). Here’s the story.

A couple of weeks ago, I gave my talk on Zero Downtime on Kubernetes. A  demo is included in the talk, as with most of my presentations. While  rehearsing in the morning, the demo worked, albeit slowly. Two days  before that, I had another demo that also uses Kubernetes and it was  already slow. But I didn’t take the hint.

During the demo, everything was slow: the of scheduling pods, of course, but  also the running and the deletion of pods. The demo failed miserably. I  didn’t even manage to stop minikube cleanly and I had to stop the VM.

To say I was disappointed is quite an understatement. That was my first  shot at this demo. I hate when demos go wrong; I hate it even more when  it works during rehearsal and it fails in front of the audience. I  apologized profusely and decided that I wouldn’t repeat the same  experience.

After the talk, I deleted the cluster and created it from scratch again. Like for the deleted cluster, I used the *virtualbox* driver. I also used the same configuration as before: 4 cores and 16 Gb. And yet, scheduling was slow… again.

I already had [some interest](https://www.reddit.com/r/kubernetes/comments/jua2f1/local_kubernetes_minikube_vs_microk8s/) in alternatives to minikube. This failure gave me the right incentive. I chose kind because some of the comments mention it in good terms.

Coming from minikube, there are a couple of differences worth mentioning. The most important one is that **kind runs in Docker**. Its name is actually an acronym for "**k\*ubernetes \*in** *d*ocker". Hence, Docker must be running prior to any kind-related operation.

As a consequence, there’s no dedicated cluster IP, everything is directly on `localhost`. However, the cluster needs to be configured explicitly to map ports.

kind.yml

```
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
nodes:
  - role: control-plane
    extraPortMappings:
      - containerPort: 30002  
        hostPort: 30002       
  - role: worker
```

|      | Map container’s port `30002` to host’s port `30002` |
| ---- | --------------------------------------------------- |
|      |                                                     |

One needs to pass the configuration at creation time:

```
kind create cluster --config kind.yml
```

The cluster configuration cannot be changed. The only workaround is to  delete the cluster and create another one with the new configuration.

Another important difference becomes visible when the image of the scheduled pod is local *i.e.* not available in a registry. With minikube, one would configure the  environment so that when one builds an image, it’s directly loaded into  the cluster’s Docker *daemon*. With kind, one needs to load images from Docker to the kind cluster.

```
kind load docker-image hazelcast/hzshop:1.0
```

I re-tested the whole demo. It works like a charm!

There’s one remaining step in my context, create an `Ingress`. The [documentation](https://kind.sigs.k8s.io/docs/user/ingress/) is clear.

##   To go further: 

- [kind Quick Start](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [kind Ingress](https://kind.sigs.k8s.io/docs/user/ingress/)
- [kind Initial design](https://kind.sigs.k8s.io/docs/design/initial/)

####  [Nicolas Fränkel](https://blog.frankel.ch/me) 

Developer Advocate with 15+ years experience consulting for many different  customers, in a wide range of contexts (such as telecoms, banking,  insurances, large retail and public sector). Usually working on  Java/Java EE and Spring technologies, but with focused interests like  Rich Internet Applications, Testing, CI/CD and DevOps. Currently working for Hazelcast. Also double as a trainer and triples as a book author.

 [Read More](https://blog.frankel.ch/me) 
