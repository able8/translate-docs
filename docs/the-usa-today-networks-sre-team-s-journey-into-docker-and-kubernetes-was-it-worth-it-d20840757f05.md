# **The USA TODAY NETWORK’s SRE team’s journey into Docker and Kubernetes. Was it worth it?**

Jul 17, 2018 From: https://medium.com/usa-today-network/devops-you-build-it-you-run-it-8f972343eb8e

Docker and Kubernetes are some of the top trending technologies in our field with many companies working to move their applications to a container-based environment. The Site Reliability Team at USA TODAY NETWORK jumped at the opportunity to move our desktop web platform from Chef-based to container-based infrastructure.

A little over a year ago, we were running more than 1,000 virtual machines configured with Chef to support 125+ desktop web applications. After managing several of our own Chef-based Kubernetes clusters, and finally landing on GKE-based clusters, we now run just over 2,000 containers and 700 pods to support the same applications.

We didn’t just migrate to run on the latest cool technology. Our ultimate goal for this project — and as site reliability engineers — was to make quantifiable improvements to our desktop web platform. Here are some questions we asked ourselves to confirm whether we met that objective.

**Are we at least breaking even?**

First and foremost, are we saving the company money or at least cost-neutral? In our Chef-based architecture we ran two instances of each of our applications out of the east region and one out of the west. For larger sites we had to run larger servers — and more of them — to handle smaller increases in traffic so we didn’t have to wait for the VMs to scale. For some of our smaller sites, we would never scale. Essentially, we were over-provisioned until and unless there was significant traffic that required us to scale.

With Kubernetes, we were able to do a little [math](http://medium.com/usa-today-network/scaling-at-scale-i-was-told-there-would-be-no-math-6e7bdbf74a3b) around the sizing of our pods and get closer to only allocating what our sites need for typical traffic — leveraging horizontal pod autoscaling when traffic increases. We were also able to go from three regions down to two, since we could scale pods faster.

After the dust settled, moving from VMs to containers reduced daily instance costs by 40%+.

**Are we meeting our Key Performance Indicators?**

Our goal wasn’t necessarily to improve application performance, but to deliver the same performance (response times) in our new environment. After all, the applications didn’t change and, as described above, we saved money and reduced overall resources.

Where we ended up seeing performance gains was with scaling. With our Chef-based environment, we would have to wait for our cloud provider to provision a new virtual machine and then wait for the Chef run to complete. During periods of increased traffic, we could see degraded origin performance as we were waiting for the new virtual machines to spin up. This could take up to ten minutes in some instances.

On Kubernetes, we leveraged horizontal pod autoscaling to scale our pods based off of CPU utilization. Instead of waiting ten minutes for the VM and Chef run, we got the new pods up in less than 30 seconds. This allows us to handle traffic peaks better and, in turn, see better performance on our origin.

The metrics indicate an improvement in availability: In addition to APM stats holding steady over time, we’ve seen a dramatic (3X) decrease in monthly alerts.

**Are we doing less manual work?**

As SREs, we don’t like manual tasks. We especially don’t like having to wake up in the middle of the night to manually resolve issues that we can handle with automation. If you ask me and my teammates, this may be the most important item on our list. By simply leveraging Kubernetes’ liveness and readiness probes, our sites essentially fix themselves.

In the Chef-based environment, our servers lost connectivity or crashed for many possible reasons which all required manual intervention from our team. We were running more than 1,000 virtual machines for our desktop platform, so issues were bound to arise. Our Help Desk would often page our on-call person, requiring them to logon and manually intervene to resolve issues.

Now that we have implemented liveness and readiness probes, if something causes issues on the pod, it’s terminated automatically and a new pod is created. This may not cover every case, like when there is a dependent micro-service having issues, but it covers most manual interventions that our team would need to take.

Even when we need to intervene manually, we’ve seen drastic improvements. For example, we might need to redeploy sites to pick up new backend configurations or deploy the site with a new version. Previously, a single site could take 15–20 minutes to deploy. And to deploy all 125+ of our sites, it could take two hours or longer.

Now leveraging helm-charts we can deploy a single site in under a minute to both the east and west regions. We can deploy all 125+ sites to both regions in under 20 minutes. We are also more confident in our deployment jobs on Kubernetes and our team members can focus on other sprint work instead of having to monitor our deployment jobs or wait for a site to redeploy.

You may be wondering: “What about the Kubernetes management side of things? Doesn’t that add more support and maintenance than the Chef environment?” When we started down the Kubernetes path, we rolled our own Kubernetes solution, based on Chef. At that time, the answer probably would have been yes. Just a few months ago, we still managed our own etcd clusters and masters. However, recently, we have migrated to GKE where Google manages the etcd cluster and masters for us. Our cluster operations team now handles RBAC, cluster upgrades, cluster scaling and implementing new functionality within our Kubernetes clusters.

**Was it worth it?**

In a word, yes! We have checked every box off our list when migrating our app to docker and Kubernetes. We have had several learning moments along the way, but those only improved our team and our entire organization’s move into container-based applications. We have seen a reduction in cost, improved our performance through faster scaling, and we are far less likely to be woken up in the middle of the night since migrating our desktop web platform to GKE. We are now a Docker- and Kubernetes-first shop and push all of our teams towards our shared platform.

## [More from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----d20840757f05--------------------------------)

Gannett is the largest local news organization in America. Our brands include USA TODAY and 250+ newsrooms spanning 46 states. Gannett’s vastly expanded local-to-national platform reaches over 50% of the U.S. digital audience, including more Millennials than Buzzfeed.

[Read more from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----d20840757f05--------------------------------)
