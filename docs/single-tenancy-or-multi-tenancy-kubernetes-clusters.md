# Choosing Single-Tenancy or Multi-Tenancy Kubernetes Clusters

29 July 2020 by [Yassine Jaffoo](http://www.appvia.io#author)

https://www.appvia.io/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters



You reach a fork in the road when setting up the architecture for your Kubernetes workloads: Do you want to operate with **many** individualized clusters or **one** large one? It’s an important decision because the direction that you go will define how your teams work, how you allocate budget and how customer data is managed.

[Kubernetes gives you](https://www.appvia.io/blog/intro-guide-to-kubernetes) the flexibility when making these choices, so it really is entirely up to the structure of your team(s), and the goals and abilities of the organization. Instead of a solely technical decision, it needs to be one that considers the impact on the overall business. Below we’ll look at the different models that you can use to structure your Kubernetes clusters and the pros and cons of each of them.

### What is Single-Tenancy?

![](https://www.appvia.io/media/pages/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters/4e8a267748-1629216749/single-tenancy-v-multi-tenancy-04-640x.jpg)

Single-tenancy is when a single application or workload, dedicated to a single tenant, lives on a single cluster. Emphasis on singularity. In this model, every cluster is purpose-built, and there is a high-degree of separation: workloads, data and teams are all separated. Within the single-tenancy structure, there are tweaks and variations that you can make to your clusters to best suit your customers and your team. More on those later.

### What is Multi-Tenancy?

![](https://www.appvia.io/media/pages/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters/8fd9d8faaa-1629216810/single-tenancy-v-multi-tenancy-05-640x.jpg)

Multi-tenancy, on the other hand, is when _multiple_ tenants operate on a single cluster. A tenant could be defined as users, clients and/or workloads, but the important distinction of this model is that all of the tenants share the cluster’s resources. It’s best-practice to isolate tenants as much as possible within a multi-tenant cluster so that one tenant can’t attack others or monopolise the shared resources.

### Comparing the two

When looking at the different architectures, consider these five core components: **Cost**, **Security**, **Reliability** and **Operational Overhead.** It’s likely that you’ll see a clear front-runner by matching up the strengths of one model with the priorities and capabilities of your business, so let’s scope single-tenancy and multi-tenancy out side by side against each of these considerations.

![](https://www.appvia.io/media/pages/blog/single-tenancy-or-multi-tenancy-kubernetes-clusters/7421f8f7ca-1629216832/single-tenancy-v-multi-tenancy-06-640x.jpg)

### 1\. How much is it going to cost?

#### SINGLE-TENANCY

It’s naturally more expensive to operate many clusters, because you’ll require a lot of resources. Each and every Kubernetes cluster has a set of management resource requirements: master nodes, control plane components, monitoring solutions etc.

There’s no getting around it - if you have a bunch of smaller clusters, you’re going to dedicate a good portion of your overall resources to support the bare-bones management functions mentioned above for each one.

The extra resource, of course, comes at a cost.

#### MULTI-TENANCY

Because the aforementioned resources are shared between workloads in a multi-tenancy model, you need less of them. So, it costs less.

You’ll be able to reuse all cluster-wide services like Ingress controllers, monitoring, logging and load balancers. A good example of resource-saving within the multi-tenancy model is the amount of master nodes you’ll need. Each cluster requires three master nodes, so if you have 10+ single-tenant clusters you’re looking at _at least_ 30 master nodes in comparison to the humble three nodes on your large, shared cluster.

While you don’t have to purchase additional resources, consider the somewhat abstract cost of potential security breaches and administrative nuances, which could soar much higher than the initial cost for an extra thirty master nodes.

### 2\. How secure is your customer data?

#### SINGLE-TENANCY

Single-tenant clusters ‘take home the gold’ when it comes to security. It’s a huge positive in this respect that the individual clusters don’t share resources, because it creates strong isolation between customers and applications.

#### MULTI-TENANCY

The level of isolation you see with a single-tenant cluster of course isn’t the same with a multi-tenant cluster. If several applications are running within the same cluster, and sharing resources as we mentioned above, they’re going to be interacting with each-other.

This could put an application, and your customers’ data, at risk when applications interact in an undesirable way. Kubernetes has safe-guards in place to prevent security breaches as much as possible, like Role Based Access Control (RBAC), PodSecurityPolicies and NetworkPolicies. But you need experience on your team to be able to manage these security capabilities in the right way.

You can’t prevent every security breach from happening, so you should be prepared for the added risk.

### 3\. How reliable is it?

#### SINGLE-TENANCY

From a reliability standpoint, there are a few things to consider with a single-tenant approach:

1. If a cluster breaks, it only affects that workload. You can mitigate the risk of multiple things failing at the same time.
2. You’ll likely have fewer users tinkering with things and making changes, and that controlled access naturally lowers the risk of something breaking in the first place.

#### MULTI-TENANCY

The exact opposite is true for the multi-tenant cluster. There’s a single point of failure, so if one thing breaks the whole cluster and all of the workloads will be down. Also, with more users working on the cluster, there’s a greater opportunity for things to go awry.

It’s also important for you to realize before committing to a multi-tenant cluster is that it can only get _so_ big before it breaks. Kubernetes has [defined upper limits](https://kubernetes.io/docs/setup/best-practices/cluster-large/) of a cluster at around 5000 nodes, 150,000 Pods and 300,000 containers. But that doesn’t mean you can even get to that size without issues - you could start seeing [issues with clusters](https://events19.lfasiallc.com/wp-content/uploads/2017/11/BoF_-Not-One-Size-Fits-All-How-to-Size-Kubernetes-Clusters_Guang-Ya-Liu-_-Sahdev-Zala.pdf) that have upwards of just 500 nodes. You’ll want to keep a close eye on the growth of your cluster, in particular the strain on the Kubernetes control plane, to keep it running efficiently.

### 4. What’s the operational overhead?

#### SINGLE-TENANCY

There are certain complexities to managing many clusters, involving the set-up and maintenance of authentication, authorization and other frameworks. If you go this route, it’s in your best interest to adopt automation that makes these processes faster and less prone to error, whether the automation is set up by your team or [t](https://www.appvia.io/solutions/kore) [hrough a product/service](https://www.appvia.io/solutions/kore).

You can manage the operational overhead by configuring your single-tenancy clusters in a more efficient way. A single tenant cluster can either be based on team structure, application stack, or deployment environment.

#### MULTI-TENANCY

Overall, managing a single cluster (or a few clusters) is easier than administrating many clusters. There are less administrative tasks and no need to worry about having to create new environments for new customers etc. You do most of the set-up a single time, and don’t have to think about it again versus having to repeat the process constantly in a single-tenancy model.

### Making the best all-around choice

There are strong considerations of each tenancy model, and the nuanced iterations that can be created within each, but you should have an overarching idea of the pros and cons of each one for your organization. What you choose for your workloads depends on what you want to achieve and the capabilities of your team.

If a single-tenancy model is the best option, but you don’t have the specialized team to manage the running of all of your individualized clusters, consider a service or product to create that for your business.

You don’t want to increase the cognitive load of your development team if they don’t have the infrastructure background or knowledge. [Wayfinder](http://www.appvia.io/products/wayfinder) enables the automation of deployment environments and provides security guardrails for teams using Kubernetes.
