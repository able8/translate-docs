# Individual Kubernetes Clusters vs. Shared Kubernetes Clusters for Development

Jun 30, 2020  From: https://loft.sh/blog/individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development

Table of Contents

- [Individual Clusters for every developer](http://loft.sh#individual-clusters-for-every-developer)
  - [Advantages](http://loft.sh#advantages)
- [Disadvantages](http://loft.sh#disadvantages)
  - [Use Cases](http://loft.sh#use-cases)
- [Shared Cluster](http://loft.sh#shared-cluster)
  - [Advantages](http://loft.sh#advantages-1)
  - [Disadvantages](http://loft.sh#disadvantages-1)
- [Use Cases](http://loft.sh#use-cases-1)
- [Conclusion](http://loft.sh#conclusion)

If you are on a higher level of cloud-native development (level 2 or higher according to [this article about the typical cloud-native journey of companies](http://loft.sh/blog/the-journey-of-adopting-cloud-native-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development)), the developers in your team need to have a direct access to Kubernetes. The first decision you have to make is if you want to use a local environment (rather level 2) or a remote environment in the cloud (level 3).

While this is not a trivial decision with many pros and cons for either method (as described in my previous post about [local vs. remote clusters for development](http://loft.sh/blog/local-cluster-vs-remote-cluster-for-kubernetes-based-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development)), the advantages of the cloud as [Kubernetes development environment](http://loft.sh/blog/kubernetes-development-environments-comparison/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) are quite compelling: It is a highly realistic environment for all applications that eventually run in the cloud; it allows for easy replicability of processes and errors; a central control of the dev system is possible; you can use basically infinite computing resources and run even very complex software.

However, after you chose the cloud environment, you are confronted with the next fundamental decision: Should every developer get an own cluster or should one cluster be shared among the developers.

To answer this question, I will describe the advantages and disadvantages of both approaches and their main use cases in this post.

## [\#](http://loft.sh\#individual-clusters-for-every-developer) Individual Clusters for every developer

The first approach is to give every developer an individual, dedicated cluster (or several clusters, if needed) to work with. This can be done by giving everyone in the team access to the cloud environment and adjust the settings there in a way that they are allowed to create clusters themselves.

Alternatively, it is also possible that the administrator of the cloud environment creates the clusters and then give the developers access to their individual one. In any case, each developer will have a cluster that only they are working in.

### [\#](http://loft.sh\#advantages) Advantages

- **Excellent Isolation.** A clear advantage of having individual clusters for each developer is that their work environments are isolated very well. This means that it is very unlikely that one developer interferes with another so that they can work independently and do not have to fear to break the development system for the whole team.

- **Full Cluster Access.** Due to the good isolation of the individual work environments, it is not a problem anymore to give the developers full cluster access, so they can configure anything they want or need (you potentially might want to limit the maximum resource consumption though). This gives the engineers the flexibility to modify the configuration of the cluster themselves and they can experiment with different Kubernetes configurations without disturbing anyone.

- **Easy Start.** Another benefit of the one-cluster-per-developer method is that it can be introduced in a team fairly fast and easy. The simplest option would be to just share the cloud account credentials and password with everyone. This, however, is of course very insecure and should only be done in very small teams with very trustable members if at all. A better solution is to give everybody their own cloud account with appropriate limitations or the clusters are simply created and then distributed by the cloud manager.


## [\#](http://loft.sh\#disadvantages) Disadvantages

- **Problematic in Private Clouds.** A first limitation of giving individual clusters to engineers is that it does not work well in most private clouds. While public clouds such as AWS and Azure provide an advanced user management, most private clouds lack a wide range of user settings and limitations. The alternative solution to give everyone admin access to the cloud environment is usually too dangerous and having one cloud administrator care for the provisioning of the clusters creates a problematic bottleneck.

- **Kubernetes Knowledge and Tools Necessary.** If a developer is supposed to fully interact with Kubernetes and has to create and configure their own cluster, they need to have some Kubernetes knowledge and tools such as kubectl. Although it is possible to streamline these processes to some extent with detailed instructions for the tool installation and step-by-step guides for the cluster setup, everyone in a team has to repeat the same steps (which, by the way, eliminates [one of the advantages of using the cloud instead of local Kubernetes environments](http://loft.sh/blog/local-cluster-vs-remote-cluster-for-kubernetes-based-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development)).

- **Maintenance Effort.** After a successful setup of the Kubernetes clusters, the work is not done. They rather have to be continuously maintained and updated. Since the individual cluster approach can lead to a lot of clusters depending on the size of the engineering team, the maintenance effort becomes immense for dozens or even hundreds of clusters.

- **Limited Oversight and Central Control.** Due to the abundance of clusters, it can be very hard to supervise the whole development system in the cloud. While it should be possible to see how much resources are consumed and what clusters are running, it is challenging to see what is currently needed and what is running idle and could be deleted. The complexity of managing the whole environment so becomes very high, especially for larger teams.

- **Computing Cost.** A general downside of using a cloud environment for development is that it can be costly, especially in public clouds. Since a system with a lot of clusters also creates a lot of redundancy for basic Kubernetes elements such as API servers, the one-cluster-per-developer approach is particularly expensive. To make things worse, some providers such as AWS and GCP charge a cluster management fee for every cluster, which is very costly if you use many small clusters as you usually would for development. (Google Cloud has only [recently announced the introduction of a cluster management fee](https://cloud.google.com/kubernetes-engine/pricing) creating [some outrage by its users](https://news.ycombinator.com/item?id=22486070).)


### [\#](http://loft.sh\#use-cases) Use Cases

There are some situations in which the individual-cluster approach for providing Kubernetes access to developer makes more sense than in others:

- **Public Clouds.** Generally, due to the extensive user management in most public clouds, this approach works better with public cloud environments than private clouds where only limited user management is available.

- **Teams of Kubernetes Experts.** In teams of Kubernetes experts, the downsides of individual setup and maintenance are less important, and the great flexibility can be a positive factor for this solution. There are even special cases in which this is the only possible approach for cloud-based development, such as if a solution or tool for Kubernetes management itself is developed.


> _This is of course rather an edge case, but in my company, we are actually developing a [solution to transform clusters into a Kubernetes multi-tenancy platform](https://loft.sh) and our engineers thus often need full cluster access._

- **Small Teams.** In smaller teams, only a few clusters are needed for this approach and generally, it is easier to set-up and manage fewer clusters. Additionally, it is possible to guide the developers through the set-up process individually and to communicate informally and fast if the team size is relatively small.

- **No Budget Constrains.** While there eventually will be budget constraints, the cloud cost is not always an important concern. This could be as the absolute cost is negligible or because the company has received credits for a public cloud, which many startups have. In these cases, there might be more important things to optimize than the cloud cost. (In a separate blog post, I describe how you [can save cost by reducing the number of clusters](http://loft.sh/blog/kubernetes-cost-savings/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development).)

- **First experiments.** Since you can get started pretty fast with this approach, it can often give you a good first feeling about how cloud environments and cloud-native development with Kubernetes works.


Since the approach to give every developer an individual cluster is very inefficient regarding computing resources and management effort, it is rarely used in reality. Even though there are some use cases in small startups or in companies that develop special Kubernetes solutions, I believe that this solution is generally inferior for the average company compared to the approach I will describe next.

## [\#](http://loft.sh\#shared-cluster) Shared Cluster

In the shared cluster approach, as the name suggests, developers share one or only a few clusters. This implies that there are several users of the same cluster sharing its resources.

In theory, it would be possible to give everyone full and direct cluster access, but this would easily end in total chaos because such a system would be very vulnerable to small mistakes. For example, someone could consume all available computing resources or change the cluster configuration, so it crashes.

For this, the cluster needs to be prepared in a way that it securely supports multi-tenancy. In Kubernetes, namespaces play an important role in this approach. However, a full shared cluster solution goes beyond this and rather means a self-service, multi-tenancy Kubernetes system that allows the developers to create namespaces on demand, while it also cares for fair usage sharing by limitations and ensures the users’ namespace isolation.

### [\#](http://loft.sh\#advantages-1) Advantages

- **Any Cloud Environment.** One advantage of the shared cluster approach is that it can be used in any cloud, public and private, as long as this cloud is running Kubernetes. In contrast to the one-cluster-per-developer approach, there are no downsides associated with either public or private clouds and it is even possible to use hybrid systems.

- **Cost-Efficiency.** While there will always be some cost associated with using the cloud as development environment, using shared clusters is at least a computing resource and thus cost-efficient solution. All basic Kubernetes functionality is shared to avoid any redundancy and the individual namespaces can sometimes even be automatically shut-off. This can [reduce the Kubernetes computing cost](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) dramatically. For example, if namespaces are always scaled down or deactivated when the developer is not working (at night, during weekends or holidays), this can often save [70% of the total computing cost of the development system](http://loft.sh/blog/how-to-save-more-than-2-3-of-engineers-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development).

- **Central Control.** Since there is only one or very few clusters, it is possible to monitor the whole system relatively easy. Of course, you still need a user management and must enforce appropriate limits for the developers using the cluster. However, there are already solutions such as [Loft](https://loft.sh) that offer all of this together with a dashboard to make sense of who uses what and what limits are enforced. This also makes it possible to shut-off individual namespaces automatically or manually from a central position.

- **Only 1 System.** From a maintenance perspective, it is much easier to manage only one larger cluster than a large number of very small ones. The cluster administrator only has to update and configure one cluster instead of repeating this process over and over for each cluster individually, which saves a lot of time and effort.

- **Ease of Use.** The developers can easily work with a system that allows on-demand namespace generation. From their perspective, the whole system will feel like a normal PaaS-system and Kubernetes is mostly a technology running in the background that they do not always have to directly interact with. For this, little to no Kubernetes knowledge and hardly any new tools are needed on the developer’s side.


### [\#](http://loft.sh\#disadvantages-1) Disadvantages

- **Isolation Necessary.** Using shared clusters to give developers access to Kubernetes, however, also comes with some disadvantages. The first one is that you need a solution for isolating the individual developers from each other to prevent chaos in your cluster. While namespaces are a good starting point and [special solutions for this problem exist](https://github.com/kiosk-sh/kiosk), user isolation in Kubernetes needs to be taken care of and can be tricky, especially if you require hard multi-tenancy. Fortunately, in most development settings, the tenants on the same cluster can be trusted and thus soft multi-tenancy is sufficient. Still, you need to introduce Network Policies to prevent namespaces to communicate with other namespaces they should not be able to talk to.

- **User Limitation and Management Necessary.** Besides the technical limitation between the users, you also need a limitation of the individual users, so they cannot excessively use computing resources and network traffic. This leads to an additional requirement of a solid user management system, which is also necessary to give every developer access to the cluster in the first place. Again, such [multi-tenancy platform solutions](https://loft.sh) exist, but if you want to build the system from scratch on your own, this is an additional topic to consider. (In a separate article, I describe [how to build such an internal Kubernetes platform](http://loft.sh/blog/building-an-internal-kubernetes-platform/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development).)

- **Setup Effort.** In any case, you will face some setup effort to get a shared Kubernetes cluster running and to configure it properly. Even with off-the-shelf solutions, you need to set them up at the beginning, which is at least some more effort than the easiest one-cluster-per-developer-approach. Still, this one-time effort can be made up by the reduced maintenance work later on.

- **Limited Flexibility for Special Cases.** Due to the limits imposed on the users’ access to the cluster, e.g. they can only use a specific number of namespaces or change only some Kubernetes settings, the developers do not have the same freedom they have with an own cluster. Since the limitations should usually be set in a smart way with the normal requirements of your team in mind, this should only become a problem in special cases. Then, the cluster administrator needs to give them additional rights, which is possible but can slow down the development process somewhat.


> The limited flexibility issue could be resolved by using [virtual Kubernetes clusters](http://loft.sh/blog/introduction-into-virtual-clusters-in-kubernetes/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) instead of namespaces. If every developer had a virtual cluster, it would be possible to even change cluster-wide resources in that virtual cluster and still not interfere with other tenants. There are some open-source concepts for virtual clusters, e.g. by [Alibaba](https://www.cncf.io/blog/2019/06/20/virtual-cluster-extending-namespace-based-multi-tenancy-with-a-cluster-view/) or by [Darren Shepherd](https://github.com/ibuildthecloud/k3v). Additionally, [Loft](https://loft.sh) has alread integrated virtual clusters as alternative to namespaces into its Kubernetes platform.

## [\#](http://loft.sh\#use-cases-1) Use Cases

Like the one-cluster-per-developer approach, the shared cluster solution also is more appropriate in some cases than in others, but generally, it can be applied in a broader range of settings:

- **Any Cloud.** Sharing a Kubernetes cluster is possible in both public and private clouds, which makes it a universal solution from a purely technical perspective.

- **Teams without Kubernetes Knowledge.** Since the shared cluster will be centrally managed by an experienced cluster admin, the average developer does not need extensive knowledge about Kubernetes. They potentially might need to install a few tools to be able to deploy to the cluster and to create namespaces in it. However, these processes are fairly easy and standardized.

- **Trusted Team Members.** One of the biggest technical challenges with the shared cluster approach is to establish multi-tenancy. Here, it is much easier to realize soft multi-tenancy, i.e. to isolate users from each other in a way that lets them work without disturbing each other but does not securely prevents malicious behavior. For most development teams, this isolation should be sufficient as all tenants are known and have no intention to destroy anything on purpose.

- **Any Team Size.** The marginal effort to add an additional developer to a shared cluster is very low. Therefore, it is possible to use shared clusters even with large teams. On the other side, it is also possible to use it with very small teams of only 2-5 developers or even alone (e.g. if you want to have separate systems for development, staging/testing and production), but here, you need to trade off the benefits of a shared cluster against the effort to set it up.


From my perspective, it looks like many professional teams that have decided to provide developers access to Kubernetes made their decision for shared cluster solutions. As an example, there was a talk about [Spotify’s](https://www.youtube.com/watch?v=vLrxOhZ6Wrg) solution for this at KubeCon North America 2019. They have developed a development platform solution for Kubernetes internally.

If you also plan to go this way, you should take a look at open source solutions such as [kiosk](https://github.com/kiosk-sh/kiosk) that solve many issues regarding [Kubernetes multi-tenancy](http://loft.sh/blog/kubernetes-multi-tenancy-best-practices-guide/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) and can be used as building block for more comprehensive solutions, such as [self-service namespace platforms](http://loft.sh/blog/self-service-kubernetes-namespaces-are-a-game-changer/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) (kiosk, for example, does not provide a user management).

However, there are also commercial, off-the-shelf solutions such as [Loft](https://loft.sh) that transform any Kubernetes cluster into a multi-tenancy enabled development platform with integrated user management, extensive user limit settings and an automatic computing resource saving sleep-mode for namespace.

## [\#](http://loft.sh\#conclusion) Conclusion

Integrating the cloud environment in the development workflow early on is still a relatively new concept as nowadays cloud technologies often only come into play at the staging and testing level. However, with the advent of Kubernetes in more and more companies, one can expect that it will be more common in the future to also give developers direct access to Kubernetes and then you need to decide if you want to use one cluster per developer or a single shared cluster for the whole team.

To get the benefits of a real Kubernetes environment in the cloud for development, both approaches are imaginable and have their specific use cases. However, from my perspective, the shared cluster approach is clearly superior for all common settings while individual clusters for developers seem to be more a solution for first experiments or niche requirements.

Overall, you often do not have to make a strict decision between the solutions. So, it might be possible that some engineers working on a specific task get their own cluster while the rest of the team works in a shared cluster environment. You might even include local environments in [the Kubernetes development workflow](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_individual-kubernetes-clusters-vs-shared-kubernetes-clusters-for-development) and use a shared cluster in the cloud only for testing and staging. As you can see, the opportunities seem endless, but I hope that this post was helpful to give you some guidance on what is possible and when to use which solution for truly cloud-native development.
