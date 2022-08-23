# How to Save More Than 2/3 of Engineers’ Kubernetes Cost

Jul 7, 2020

Table of Contents

- [The Cost of the Cloud](http://loft.sh#the-cost-of-the-cloud)
- [Cloud-Native Workflows Cause Additional Cloud Cost](http://loft.sh#cloud-native-workflows-cause-additional-cloud-cost)
- [Inefficiency 1: Too Many Clusters](http://loft.sh#inefficiency-1-too-many-clusters)
  - [Problem: Many clusters are expensive](http://loft.sh#problem-many-clusters-are-expensive)
  - [Solution: Share a cluster](http://loft.sh#solution-share-a-cluster)
  - [Maximum Savings](http://loft.sh#maximum-savings)
  - [Example Calculation](http://loft.sh#example-calculation)
- [Inefficiency 2: Unused Resources](http://loft.sh#inefficiency-2-unused-resources)
  - [Problem: Computing Resources are often unused](http://loft.sh#problem-computing-resources-are-often-unused)
  - [Solution: Let spaces “sleep” when unused](http://loft.sh#solution-let-spaces-sleep-when-unused)
  - [Maximum Savings](http://loft.sh#maximum-savings-1)
  - [Example Calculation](http://loft.sh#example-calculation-1)
- [How to get these savings](http://loft.sh#how-to-get-these-savings)
  - [Example Calculation](http://loft.sh#example-calculation-2)
- [Private Clouds](http://loft.sh#private-clouds)
- [Conclusion](http://loft.sh#conclusion)

## [\#](http://loft.sh\#the-cost-of-the-cloud) The Cost of the Cloud

Using the cloud is not cheap. If you are using a public cloud, you will get a monthly bill for computing resources, traffic, and for any additional services. But also for a private cloud, the cost is substantial, taking into account the hardware, datacenter, and maintenance cost. (For simplicity, I will assume the public cloud case in the following, but I added a paragraph about private clouds at the end.)

The widespread adoption of Kubernetes has not changed this situation fundamentally. [Some say it has only increased cost](https://www.reddit.com/r/kubernetes/comments/a91c5n/is_running_kubernetes_always_expensive/) because the Kubernetes clusters themselves consume resources (e.g. Master node, LoadBalancer, etcd). [Others are convinced that Kubernetes could at least reduce the cost](https://techbeacon.com/devops/one-year-using-kubernetes-production-lessons-learned) because it makes the use of the cloud resources more efficient. Overall, it probably depends on your setting and perspective, but the cost of the cloud remains an issue either way.

## [\#](http://loft.sh\#cloud-native-workflows-cause-additional-cloud-cost) Cloud-Native Workflows Cause Additional Cloud Cost

For this, [reducing Kubernetes cloud computing cost](http://loft.sh/blog/reduce-kubernetes-cost/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost) is always attractive, especially with more and more companies adopting Kubernetes throughout the organization and introducing cloud-native development, which includes giving engineers access to [Kubernetes for development](http://loft.sh/blog/kubernetes-development-environments-comparison/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost), testing, experimentation, or CI/CD (level 2 or level 3 on the cloud-native journey [according to this article](http://loft.sh/blog/the-journey-of-adopting-cloud-native-development/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost)).

In these cases, companies are sometimes hesitant to spend a lot of money on the cloud as this, contrary to spending money on production workloads, is not directly connected to revenue and customer benefit. Still, [giving developers direct Kubernetes access can make a lot of sense and can pay off in terms of higher stability and more efficient workflows, too](http://loft.sh/blog/is-kubernetes-still-just-an-ops-topic/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost).

Making cloud-native development cheaper and easier can thus be an important driver for Kubernetes adoption, especially in small and medium-sized organizations. I will therefore describe 2 “obvious” inefficiencies of the engineers’ use of the cloud and easy ways of eliminating them to significantly reduce the cost of using Kubernetes. Since developers use the cloud differently than production systems, these cost-saving opportunities are most relevant for non-production environments.

## [\#](http://loft.sh\#inefficiency-1-too-many-clusters) Inefficiency 1: Too Many Clusters

### [\#](http://loft.sh\#problem-many-clusters-are-expensive) Problem: Many clusters are expensive

An easy option to give every engineer access to a cloud environment is to give everyone an own cluster. However, this is pretty inefficient for several reasons.

- **Cluster Management Fee.** Big public cloud providers such as [AWS](https://aws.amazon.com/eks/pricing/?nc1=h_ls) and [Google Cloud](https://cloud.google.com/kubernetes-engine/pricing) charge a cluster management fee for their Kubernetes offerings. If the cluster is running all the time, only the cluster fee will add more than $70 to your bill for every cluster. And this is not including any computing resources that will come on top of that.
- **Redundant Kubernetes Functionality.** Since every cluster is running independently of the others, they all need to provide the whole Kubernetes functionality themselves. That means that, for example, every cluster needs its own load balancer that could actually be shared by many developers. This makes many clusters rather inefficient from a computing resource perspective.
- **No Central Oversight.** Many clusters make it hard to oversee the whole system. Even with cluster management tools, it is not always clear what each cluster is used for and if it is still used at all. In many cases, clusters will be created for a specific purpose such as a test, but often, they will not be deleted afterward so they continue to run and cost you money without any benefit.

### [\#](http://loft.sh\#solution-share-a-cluster) Solution: Share a cluster

The solution to the first inefficiency of many non-production clusters is to share clusters among engineers. That means that only one cluster is used and developers are working with their own namespaces that limit them and isolate them from each other. [I wrote a separate article about a comparison of individual clusters and shared clusters](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost), but regarding the described cost efficiencies, sharing a cluster has the following advantages:

- **Shared Cluster Management Fee.** Instead of paying the management fee for every developer as there are many clusters, you only need to pay it once. So, the cost per developer is reduced from about $70 per month to $70 divided by the number of developers and is thus decreasing with the number of developers.
- **Share Kubernetes Functionality.** Similar to the management fee, also the Kubernetes functionality can be shared and thus does not need to run redundantly. With more developers sharing it, the load will of course also be higher, but there will definitely still be cost savings from sharing a cluster.
- **Easier oversight.** In a shared cluster, it is much easier to control the users, especially if they are created with specialized solutions for this use case, such as [Loft](https://loft.sh). The cluster-admin has thus less effort to control the system and can shut down unused spaces easily. This shutdown can even be triggered automatically with a [“sleep mode”](http://loft.sh/docs/sleep-mode/basics?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost) (for more, see below).

### [\#](http://loft.sh\#maximum-savings) Maximum Savings

The [Kubernetes cost savings from reducing the number of clusters](https://medium.com/faun/kubernetes-cost-savings-by-reducing-the-number-of-clusters-336378680995) by sharing a cluster are highest if there are **many clusters** used before, which is typical for **larger teams** or if clusters are used for **many purposes** such as in-cluster development, testing, and staging. The saving effect relative to the total cost is also particularly high if the **clusters are rather small and thus cheap**. In these cases, a cluster management fee and the redundancy is a comparably expensive cost driver.

### [\#](http://loft.sh\#example-calculation) Example Calculation

Let’s take a look at a simple example to demonstrate the cost-saving opportunities of reducing the number of clusters by Kubernetes cluster sharing:

If in a small team of 5 developers, everyone gets an own Kubernetes cluster and the clusters are always running, you have to pay $360 only for the cluster management fee. If they share one cluster, you only have to pay the cluster management fee for one cluster, which is $72 per month. In such a small team, you would thus already save $288 just in cluster management fees. (For GKE, one “zonal cluster” is free, so you would pay $288 in cluster management fees for 5 clusters and $0 for one cluster. Overall, the savings remain unchanged.)

On top of that, you could save additional money by sharing the basic Kubernetes features and by a better control which prevents unused clusters. However, since these factors are hard to estimate, I do not add them to the calculation.

Of course, if you use commercial tools such as Loft, you need to pay for these tools, too. In the example of 5 developers using Loft, the cost would be $75.

**Your total savings from sharing a cluster would so be at least $213 (60%) for only a small team of 5.**

![](http://loft.sh/blog/images/content/cost-savings-clustersharing.jpg?nf_resize=fit&w=1040)

## [\#](http://loft.sh\#inefficiency-2-unused-resources) Inefficiency 2: Unused Resources

### [\#](http://loft.sh\#problem-computing-resources-are-often-unused) Problem: Computing Resources are often unused

In non-production settings, the Kubernetes clusters are actually unused most of the time, while they cost you money all the time. One option to prevent this is to tell your engineers to reliably delete their stuff or shut down their spaces when they do not use them anymore. This is, however, not so easy to implement in reality. Sometimes, engineers may just forget about it as distractions come in or they are not at work when they should shut down their spaces. A few examples, when computing resources are typically unused:

- **In-cluster development.** Developers use Kubernetes in the cloud already during development, e.g. with special tools such as [DevSpace](https://github.com/devspace-cloud/devspace) or [Skaffold](https://github.com/GoogleContainerTools/skaffold). The used spaces continue to run even after work, at night, during weekends, holidays, or sick leave of the engineer.
- **Experiments.** Sometimes, engineers need to use the power of the cloud to test their applications. This is often the case for ML/AI applications that require a lot of computing resources or specific hardware. While these tests run for a while, the underlying hardware is sometimes even running when there are actually no experiments executed.
- **Staging/Testing.** Developers may push their finished code to a staging or testing environment where it is inspected before it can run in production. Similar to the in-cluster development situation, there are usually long periods of time when no tests are executed, no developer is pushing their code and no one is looking at the application in staging, e.g. at night. Still, the cloud environment continues to run and cost you money.

### [\#](http://loft.sh\#solution-let-spaces-sleep-when-unused) Solution: Let spaces “sleep” when unused

The solution to this problem is to activate a so-called sleep mode of unused spaces in the cloud. This means that the containers will be terminated while all persistent data will be preserved. The computing resources can thus be scaled down and you do not have to pay for them anymore. When they are needed again, they can be scaled up and the containers will be restarted automatically again with all the persistent data immediately available. In tools such as [Loft](https://loft.sh), a sleep mode can either be started manually or automatically after a pre-defined period of time.

Overall, this means that the cloud resources in Kubernetes that the engineers use are flexibly available, and you only pay during the times they are really needed. The advantage of an automatic sleep mode is that it does not affect the engineers at all, they do neither have to start nor to end it, so their [Kubernetes workflow](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost) remains as efficient as possible.

### [\#](http://loft.sh\#maximum-savings-1) Maximum Savings

The cost savings of a sleep mode are of course highest when there are a lot of computing resources generally used, i.e. **in larger teams or with complex applications and experiments with high resource requirements**. Another factor for the potential of the sleep mode is the **length of the idle times**.

For a typical engineer with a 40 hour work week, the savings can be about 75% (128 of 168 hours per week are non-working hours, not accounting for holidays, sick leave, or other tasks of the engineer that do not require computing resources such as meetings). If the Kubernetes environment is used for some other tasks such as testing or experimentation, the savings can be even higher (potentially, there are days during which 100% could be saved when no experiments are run during that day).

### [\#](http://loft.sh\#example-calculation-1) Example Calculation

Now, we can look at a similar example as before with a team of 5 engineers. Since we are now looking at computing costs only, we will ignore cluster management fees. We also assume that the average engineer works 8 hours a day on 5 days per week. For simplicity, we further ignore additional factors such as holidays, sick leave, and other tasks that do not require a Kubernetes access, e.g. meetings. In sum, [these factors can of course be substantial](https://www.quora.com/How-many-hours-in-average-do-software-engineers-Developers-or-coders-code-in-a-day-Do-developers-code-8-hours-every-day) but would only support the result of the example in terms of further saving opportunities.

Now, let’s say that the engineers need 4vCPUs and 15GB RAM each to run the application they are working on. Depending on the software, this could, of course, be quite a lot (e.g. for a simple website) or not nearly enough (e.g. for complex ML applications), so there is some variance here. In GKE, this is equal to an “n1-standard-4” machine that costs about $97 per month.

For the team of 5 engineers, you would so face computing resource cost of $485 per month. Since the engineers are only working 40 hours per week, the computing resources are not used 128 hours per week, which is more than 75% of the time.

With a sleep mode configured in a way that sends unused spaces to sleep after 1 hour of inactivity, the computing resources are only active for 45 hours per week and the unused time is reduced to 5 hours (1 hour per day). For this, you only need to pay about 27% of the computing resource cost and you can reduce the time during which the resources are unused from 75% to 12.5%.

**In absolute terms, this means that you only have to pay about $130 for all 5 engineers per month saving about $355. Again, even after deducting $75 for tools to enable a sleep mode, such as Loft, you can get net savings of $280 every month.**

![](http://loft.sh/blog/images/content/cost-savings-sleepmode.jpg?nf_resize=fit&w=1040)

## [\#](http://loft.sh\#how-to-get-these-savings) How to get these savings

With solutions such as [Loft](https://loft.sh/?utm_medium=reader&utm_source=other&utm_campaign=blog_how-to-save-more-than-2-3-of-engineers-kubernetes-cost), that is based on [kiosk](https://github.com/kiosk-sh/kiosk), it is possible to eliminate both of the described inefficiencies. Loft is a multi-tenancy manager for Kubernetes that allows engineers to create namespaces in a shared cluster on-demand. It also comes with an in-built sleep mode feature to save costs for unused computing resources.

To use Loft, you install it in a Kubernetes cluster that then becomes your management cluster for Loft. Afterward, you can connect other clusters (or the management cluster itself) to Loft and in these connected clusters, the engineers can start their isolated namespaces. Overall, Loft achieves the solutions I described before out-of-the-box:

- **Shared Clusters.** Loft transforms your cluster into a multi-tenancy platform on which engineers can create their own namespaces whenever they need them. The user isolation and management are cared for by Loft and it also lets you limit the resource consumption on different levels (per user, per space). This also allows the cluster-admin to centrally control and manage the users and they can even take a direct look at the engineers’ spaces with space sharing to see if they are still in use.
- **Sleep Mode.** Another central feature of Loft is its sleep mode that automatically scales down unused and idle spaces after an individually determined period of time. When the space is needed again, it will be woken up automatically, too, and the engineers can proceed from where they ended. For this, engineers do not have to think about the cloud infrastructure anymore and can work freely and uninterruptedly.

### [\#](http://loft.sh\#example-calculation-2) Example Calculation

Combining the two examples from above, one can calculate the total cost savings of introducing shared clusters and a sleep mode for the sample engineering team of 5. For this, one only has to sum up the cost savings from both examples. The only thing that needs to be adjusted is the cost to get cluster-sharing and the sleep mode. Since both features are included in Loft, the usage fee for it only has to be accounted for once.

**Overall, it is so possible to save more than two-thirds of the cost, which is $568 for the sample engineering team. Reducing the cost per engineer from $169 to $55.4 makes providing engineers with direct access to Kubernetes in the cloud also feasible for smaller organizations.**

![](http://loft.sh/blog/images/content/cost-savings-clustersharing-sleepmode.jpg?nf_resize=fit&w=1040)

## [\#](http://loft.sh\#private-clouds) Private Clouds

The described problems and solutions are mostly applicable to private clouds, too. Here, sometimes the usage-cost relationship is not so clear, especially if you own the hardware yourself. In these cases, the computing cost for the cloud is rather a fixed cost than a variable cost, except for power and network traffic cost. Some other additional factors such as cluster management fees are not an issue in private clouds, but the maintenance effort is usually a bigger challenge.

If you apply the described solutions of cluster sharing and a sleep mode in a private cloud, it can still save you cost, e.g. because central Kubernetes features are not running redundantly. Overall, the direct cost effect is potentially smaller, but there is an additional indirect effect: If you are using sleep mode and can configure user limits centrally, it is easier to “oversubscribe” the available computing resources as you already know that not all of them will be used at the same time.

For example, you could activate the sleep mode so that all spaces for in-cluster development for developers are sleeping after work and then schedule complex processes such as ML experiments to run at night in spaces that will sleep again in the morning when the engineers’ workday starts. Another example would be remote teams in different time zones that can alternately use the same computing resources.

Therefore, there is also a savings impact in private clouds that might be less immediate but if you are able to use the available computing resources more efficiently, this is generally beneficial. This can eventually lead to fewer hardware needs, which will also save you a lot of cost.

## [\#](http://loft.sh\#conclusion) Conclusion

With engineers’ access to cloud environments, the cost of the cloud becomes an even more important issue. While simple solutions with many clusters and no central control might be a first start, they are definitely not efficient in terms of cost. Over time, more companies should thus move to more advanced solutions to limit the number of clusters and to reduce the time computing resources cost money while they are not used.

With cluster sharing among engineers and an individually configured sleep mode, a lot of computing resource costs can be saved. Fortunately, the implementation of these solutions with off-the-shelf solutions such as [Loft](https://loft.sh) is now easier than ever and available even to small and medium-sized companies and not only to big players that might build such solutions themselves.

In the provided example, the cost can so be reduced by more than two-thirds leading to savings of about $570 per month in a 5-person team. Maybe even more importantly, the cost per engineer is reduced from deterrent $169 to moderate $55.4 per month. This makes the adoption of Kubernetes throughout whole organizations not only cheaper and more attractive but also enables even smaller companies to start their cloud-native transformation.

https://loft.sh/blog/how-to-save-more-than-2-3-of-engineers-kubernetes-cost