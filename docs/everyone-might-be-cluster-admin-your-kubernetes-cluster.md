# Everyone might be a cluster-admin in your Kubernetes cluster

February 27, 2020

Quite often, when I dive into someone's Kubernetes cluster to debug a problem, I realize whatever pod I'm running has _way_ too many permissions. Often, my pod has the `cluster-admin` role applied to it through its default ServiceAccount.

Sometimes this role was added because someone wanted to make their CI/CD tool (e.g. Jenkins) manage Kubernetes resources in the cluster, and it was easier to apply `cluster-admin` to a default service account than to set all the individual [RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) privileges correctly. Other times, it was because someone found a new shiny tool and blindly installed it.

One such example I remember seeing recently is the [spekt8](https://github.com/spekt8/spekt8) project; in it's installation instructions, it tells you to apply an rbac manifest:

```
kubectl apply -f https://raw.githubusercontent.com/spekt8/spekt8/master/fabric8-rbac.yaml
```

What the installation guide doesn't tell you is that this manifest grants `cluster-admin` privileges to _every single Pod in the default namespace_!

At this point, a more naive reader (and I mean that in the nicest way—Kubernetes is a complex system and you need to learn a lot!) might say: "All the Pods I run are running trusted applications and I know my developers wouldn't do anything nefarious, so what's the big deal?"

The problem is, if any of your Pods are running _anything_ that could potentially result in code execution, it's trivial for a bad actor to script the following:

1. Download`kubectl` in a pod
2. Execute a command like:

`kubectl get ns --no-headers=true | sed "/kube-*/d" | sed "/default/d" | awk '{print $1;}' | xargs kubectl delete ns`


If you have your RBAC rules locked down, this is no big deal; even if the application inside the Pod is fully compromised, the Kubernetes API will deny any requests that aren't explicitly allowed by your RBAC rules.

However, if you had blindly installed `spekt8` in your cluster, you would now have no namespaces left in your cluster, besides the default namespace.

## Try it yourself

I created a little project called [k8s-pod-rbac-breakout](https://github.com/geerlingguy/k8s-pod-rbac-breakout) that you can use to test whether your cluster has this problem—I typically deploy a script like this [58-line index.php](https://github.com/geerlingguy/k8s-pod-rbac-breakout/blob/master/index.php) script to a Pod running PHP in a cluster and see what it returns.

You'd be surprised how many clusters give me all the info and no errors:

![RBAC Breakout page example](http://www.jeffgeerling.com/sites/default/files/images/rbac-breakout-page-example.png)

Too many Kubernetes users build snowflake clusters and deploy tools (like `spekt8`—though there are _many_, many others) into them with no regard for security, either because they don't understand Kubernetes' RBAC model, or they needed to meet a deadline.

If you ever find yourself taking shortcuts to get past pesky `User "system:serviceaccount:default:default" cannot [do xyz]` messages, think twice before being promiscuous with your cluster permissions. And consider automating your cluster management (I'm writing a [book](https://www.ansibleforkubernetes.com) for that) so people can't blindly deploy insecure tools and configurations to it!

## Further reading

- [Run Ansible Tower or AWX in Kubernetes or OpenShift with the Tower Operator](http://www.jeffgeerling.com/blog/2019/run-ansible-tower-or-awx-kubernetes-or-openshift-tower-operator)

- [Debugging networking issues with multi-node Kubernetes on VirtualBox](http://www.jeffgeerling.com/blog/2019/debugging-networking-issues-multi-node-kubernetes-on-virtualbox)

- [Monitoring Kubernetes cluster utilization and capacity (the poor man's way)](http://www.jeffgeerling.com/blog/2019/monitoring-kubernetes-cluster-utilization-and-capacity-poor-mans-way)

## Comments

Softwarelimits – [1 year ago](http://www.jeffgeerling.com/comment/10580#comment-10580)

There is only one solution for this: we need a management tool that abstracts everything away on top of K8S! Also on top of that a cluster of cluster management tools that manage your management cluster. :)))

Seriously, for many years there existed a good practice of providing secure installations of software by default.

Delivering software with an insecure configuration as default was a bug.

Today software is not secure by default and everybody is fine with that - why are people accepting this?

Also look at the crazy amount of CPU cycles wasted on a full blown K8S cluster even before one single byte of useful payload is computed - while the planet is burning and future generations are begging for a change we are establishing a new data center operating system that needs even more energy - why are people accepting this kind of bad systems?

We had clustering systems for Linux for many years and students were able to install them as a weekend project - this knowledge is getting lost. Instead everybody thinks "new system" is cool just because it is from "big corp" - that is pure distopia brainwash, at that point we could just install Windows and "just restart your machine if a problem occurs".

Measure, analyze, think, look at the facts, educate yourself and be brave enough to trust your own conclusions.

If a system looks bad, delivers bad results and generates too many new problems, maybe in fact it is a bad system.

Thanks for providing the tools that help with that!
