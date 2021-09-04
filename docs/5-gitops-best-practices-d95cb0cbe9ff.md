# 5 GitOps Best Practices

[Nov 7, 2019](http://blog.argoproj.io/5-gitops-best-practices-d95cb0cbe9ff?source=post_page-----d95cb0cbe9ff--------------------------------)·3 min read

GitOps has been on the scene for some time now, but some things trip up users, both new and old. Here are some of the key best practices we’ve discovered while engineering Argo CD and running it at scale managing thousands of apps in production at Intuit.

# \#1: Two Repos: One For App Source Code, Another For Manifests

Many engineers start out with both their app source code and their manifests in a single repo. But this has a couple of problems:

- Every commit to your app can result in a deployment.
- Everyone is also able to release, but you might not want everyone to be able to.
- The Git log mix app changes with deployment changes.

Instead, we suggest you have two repos and keep your app’s source code in one, and your deployment config (i.e. your manifests) in the other.

# \#2: Choose The Right Number Of Deployment Config Repos

Consider how many repos you’ll keep your deployment config in. There’s no one-size-fits-all solution, so here are some suggestions:

- A mono-repo for a small company with little automation and where you trust everyone.
- A repo per team for a mid-sized company with some automation.
- A repo per service for a large company with plenty of automation and a need for control.

With their own repo, teams can look after themselves. They can decide who can release, rather than either having a single central team being a release bottleneck or giving every team write access.

# \#3: Test Your Manifests Before You Commit

Many engineers commit and push changes to their manifests and then let the GitOps agent validate their changes by seeing if it can deploy the app. We see a lot of issues that escape into pre-production environments that could have been prevented by test the changes before committing and pushing them.

The agent is typically running a CLI command (such as “helm template”) to generate your manifests, so you can test your manifests by running the command locally before you even commit your changes.

# \#4: Git Manifests Should Not Change Due To External Changes

Both Kustomize and Helm allow you to template manifests that can be different for the same commit.

If your manifests can change without committing to Git:

- You won’t be able to control or audit changes.
- You probably can’t be able to roll back to a good version.

If you’re using Kustomize remote bases, pin them to specific commits.

```
# Un-pinned:
bases:
- github.com/argoproj/argo-cd/manifests/cluster-install# Pinned:
bases:
- github.com/argoproj/argo-cd//manifests/cluster-install?ref=v0.11.1
```

If you’re using Helm dependencies, pin them too.

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

It’s no secret that you need to do extra work for secret management in GitOps. Plan how you’ll manage secrets before you start. Here are some options we know people have used successfully:

- [Bitnami Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets)
- [Godaddy Kubernetes External Secrets](https://github.com/godaddy/kubernetes-external-secrets)
- [Hashicorp Vault](https://www.vaultproject.io/)
- [Helm Secrets](https://github.com/futuresimple/helm-secrets)
- [Kustomize secret generator plugins](https://github.com/kubernetes-sigs/kustomize/blob/fd7a353df6cece4629b8e8ad56b71e30636f38fc/examples/kvSourceGoPlugin.md#secret-values-from-anywhere)

# Conclusion

While there’s rarely just one way to do something, hopefully, these general best practices can point you in the right direction!

We hope to meet you at [**KubeCon San Diego**](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2019/) and see you at one of the following Argo CD sessions!

- [Leveling Up Your CD: Unlocking Progressive Delivery on Kubernetes — Daniel Thomson & Jesse Suen, Intuit](https://kccncna19.sched.com/event/Uaaj?iframe=yes&w=&sidebar=yes&bg=no)
- [Tutorial: Everything You Need To Become a GitOps Ninja — Alex Collins & Alexander Matyushentsev, Intuit](https://kccncna19.sched.com/event/Uaee/tutorial-everything-you-need-to-become-a-gitops-ninja-alex-collins-alexander-matyushentsev-intuit?iframe=yes&w=100%&sidebar=yes&bg=no)
- [Panel: GitOps User Stories — Javeria Khan, Palo Alto Networks; Matthias Radestock, Weaveworks; Hubert Chen, Branch; Kyle Rockman, Under Armour; & Jesse Suen, Intuit](https://kccncna19.sched.com/event/UaYh/panel-gitops-user-stories-javeria-khan-palo-alto-networks-matthias-radestock-weaveworks-hubert-chen-branch-kyle-rockman-under-armour-jesse-suen-intuit?iframe=yes&w=100%&sidebar=yes&bg=no)
