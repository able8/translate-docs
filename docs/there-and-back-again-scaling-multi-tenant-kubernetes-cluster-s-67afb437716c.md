# There and Back Again — Scaling Multi-Tenant Kubernetes Cluster(s)
May 12, 2020

From: https://medium.com/usa-today-network/there-and-back-again-scaling-multi-tenant-kubernetes-cluster-s-67afb437716c

_Everyone loves a good war story_.

They say there are lessons to be learned in IT “war stories”. But maybe the real lessons are what happened afterwards, not the event or even what led up to the event? Just as the event is not the whole story, nor is any particular tool the whole story.

At Gannett, we’ve got some war stories. Plenty, in fact. When you are working for the largest local news organization, you have to rapidly adapt to changing landscape or be left behind.

![](https://miro.medium.com/max/1400/1*IlAKo4iJnvxqzotDZZGGOw.png)

## From Here — A Kubernetes Cluster for Everyone

Kubernetes is one such tool that we use to bridge that gap between where we’ve come from and where we need to go. Our first success with Kubernetes was in November 2016, when our home-grown Kubernetes clusters carried USA Today’s election coverage. It was such a success that we quickly started building out as many Kubernetes clusters as development teams wanted and were willing to manage using [Chef](https://www.chef.io/) and [Scalr](https://www.scalr.com/).

Listening to this talk, you’ll quickly realize how complicated and difficult it is to run Kubernetes the “hard way”. We built up an amazing infrastructure to automate the deployment and management of the 20+ clusters using Chef and Scalr. However, it was still hard, especially on the development teams who wanted to deploy their applications quickly without a lot of hassle. It was still a big step forward, but not far enough.

[**The USA TODAY NETWORK’s SRE team’s journey into Docker and Kubernetes. Was it worth it?** ](http://medium.com/usa-today-network/the-usa-today-networks-sre-team-s-journey-into-docker-and-kubernetes-was-it-worth-it-d20840757f05)

[**DevOps — You Build It, You Run It**](http://medium.com/usa-today-network/devops-you-build-it-you-run-it-8f972343eb8e)

## To There — A Shared, Managed Kubernetes Cluster for Everyone

_“Provide a resilient, optimized, feature rich, and easy to use platform that increases speed of innovation while reducing developer toil.”_

I pick up the story again in early 2018. It’s becoming clear that asking development teams to run their own Kubernetes clusters does not, in fact, “ _reduce developer toil_”. A new approach is needed. [Google’s Kubernetes Engine](https://cloud.google.com/kubernetes-engine) offering is gathering speed and mindshare. At the same time, Gannett is migrating a large portion of our cloud infrastructure from AWS to GCP.

A simple, elegant solution appears before us — form a new team to manage shared GKE clusters for everyone! A hybrid model where an operations team, a managed, secure service and RBAC come together to provide Developers with all the access they need to run the applications how they want.

![Division of responsibilities between teams.](https://miro.medium.com/max/1400/0*F4RTWt6_G9TaN5Cv)

The solution seems straightforward enough. We can follow the best practice documents around multi-tenancy. We create a namespace per team. We create an admin service account in each namespace and give those credentials to the various development teams. We implement Kubernetes integration with Hashicorp Vault. We reduce developer toil by taking back most of the work around maintaining a Kubernetes cluster, including collecting logs and basic metrics in a centralized location. We run two production clusters in one GCP project and the pre-production another project.

![](https://miro.medium.com/max/960/1*5_J9HtUvj2I7lC0rKxbNgw.png)

We start to hit a few issues here and there. We purposely did not apply limit on namespaces. We can’t predict what teams will be needing to run their production applications and why throttle access to readily available resources? A deployment which goes bad and steals resources wasn’t entirely unexpected and could easily be dealt with. We naively thought that by splitting teams across workloads and nodepools would be adequate.

We didn’t anticipate what the combined and highly diverse load would do to GKE. We begin to hit rare bugs in the OSS Kubernetes kernel. The Jenkins K8S plugin we use triggers goroutines to leak and destabilize the entire cluster. We request Google support restart our Master API server every few days to prevent the cluster from crashing over the course of several weeks. A bug in the OSS Linux kernel gets repeatedly triggered by all of the containers starting and stopping on a single node. We start proactively monitoring nodes and rebooting them every few hours until the bug is remedied.

Our methodology for maintaining and updating clusters stops scaling well. We originally started with a dedicated helm chart & [Concourse](https://concourse-ci.org/) job per team. With over 40 teams on three clusters, the helm charts and jobs are becoming harder to maintain and prone to simple errors that delete teams’ namespaces and deployments. The clusters were created with [Terraform](https://www.terraform.io/) files. Those templates are now dangerous to re-apply, as the state drift will cause entire clusters to be deleted and recreated. Ad-hoc documentation starts to spring up attempting to document all the special knobs that were turned on which cluster to fix which issue. Building a new cluster is easily a week long process and not likely to replicate everything. We don’t mention the words “disaster recovery” anymore.

Cluster upgrades take days and can cause multiple teams’ applications to break. Keeping the clusters up to date is now a break/fix situation only. Migrating to new node pools is near impossible. Each new node requires the Kubernetes service controller to update all of the backends to add the new nodes to every GCP Forwarding Rule in the whole cluster. We have thousands of services resulting in about 30 minutes of delay migrating to a new node and all of the other nodes being updated.

Less easy to identify and rectify are the “hard quotas”, which are documented but not visible in the GCP quotas page. We found them while troubleshooting problems through cryptic log messages or side comments from support. Some limits are hard set and in most cases cannot be adjusted. Sometimes we were able to request Google engineering increase these limits, a bit.

- Internal Forwarding Rules has a default maximum of 50
- Maximum services/node limits

As the months wore on and the outages stacked up, it became all too clear that we had become too multi-tenant for our own good. Kubernetes is built around the idea of scaling pods and nodes, not more and more services with only a handful of pods. It expects workloads to be similar-ish and not having a large degree of churn. We were mixing too many workloads and deployments in one place.

## And Back Again — Managed Kubernetes Clusters for Everyone

In 2019, it was time to go back to where we began. All 40 development teams working on the same clusters wasn’t working super well anymore. The question became how we take the best of our previous two iterations of Kubernetes and bring that forward? How can we create shared clusters that share costs by allowing teams to use the same nodes? How can we isolate some workloads, yet be multi-tenant? How can we increase our resiliency by implementing updated GKE features like regional-master and VPN hub & spoke networking? How can we manage the managed service back like we did in the Scalr days, but with less effort. Mergers are looming, budget cuts and shrinking staff are likely to be in our near term future.

We began by looking at alternative methods for managing clusters other than our home grown Concourse jobs. These jobs did what they needed to do but were going to be painful to extend to a dynamically group set of new clusters. The process to build a new cluster had drifted organically over the year and was no longer reproducible. Teams were hard coded with IPs, service accounts credentials and on-the-fly deployment customizations.

_Goal 1) Divide and Conquer_

We needed to find a middle ground between every team having a dedicated GKE cluster and all teams being on the same cluster. The obvious solutions of splitting by “mission critical” applications on some clusters, but not others wouldn’t work. Every team believes their applications are mission critical and did we really want all of those on a single cluster anyways? We went with the broader categorization of Production, Pre-Production and “Tools”.

![](https://miro.medium.com/max/1240/1*0DZh2T-mlnjedYkUvMLWlA.png)

The second categorization is more of a psychological one. What are the individual development teams’ tolerance and desire for new features? What level of risks are they willing to accept to always have access to the latest versions and newest features? Some teams are actively waiting for Alpha features to become Beta features. Other teams would much prefer to never hear the word “alpha” or “beta” in regards to their production environments.

![](https://miro.medium.com/max/1280/1*gA9Vf6eqSNQpszBHblpndA.png)

_Goal 2) Automate All The Things_

We could all agree that more GKE clusters were needed. But how can we do that simply without adding more complexity to our already overscheduled workloads? We did look at other solutions for managing clusters. None of them did everything we needed to do. Some were too expensive. Some trivialized actions to the point where they were no longer reproducible. Other tools would take a huge knowledge lift for us to implement, on top of an already painfully diverse workload.

Our ideal solution would cover all of these things:

- Cheap and allows us to re-use existing tools like Hashicorp Terraform, Concourse and[Hashicorp Vault](https://www.vaultproject.io/)
- Custom clusters, yet all clusters share same base configurations
- Dynamically generate cluster list and credentials
- Modify and update many clusters simultaneously, yet track each cluster’s unique applied yaml via git (GitOps Rules!)

To briefly summarize, [GitOps upholds the principle that Git is the one and only source of truth. GitOps requires the desired state of the system to be stored in version control such that anyone can view the entire audit trail of changes. All changes to the desired state are fully traceable commits associated with committer information, commit IDs and time stamps. This means that both the application and the infrastructure are now versioned artifacts and can be audited using the gold standards of software development and delivery.](https://www.cloudbees.com/gitops/what-is-gitops) Our new solution must meet these clearly stated goals.

We were able to leverage our existing Terraform pipelines to automate building of the GKE clusters and all of the supporting GCP resources.

![](https://miro.medium.com/max/1400/1*dUMZyPUYQXy0LAJxJN4JRQ.png)

Next came the deployment yaml’s. Our clusters still need monitoring, cost reporting and logging applications along with Hashicorp Vault integration and several other internal tools. Some clusters will need special configurations or an entirely different set of applications. We can re-use our cluster designations to apply these things.

We realized we didn’t need a complicated new product to add to our already very diverse set of toolings. A bash script with some simple looping would solve this problem for us. We could add it to Concourse as another job and no longer worry about which cluster was which type and who was using it when.

We generate a finalized version of every custer’s unique Kubernetes manifest and save it to git. We like helm charts, but it can be tricky to see the final applied configuration without connecting to each cluster and Tiller. We also wanted a way to document changes naturally. And remember all of those teams who needed custom namespaces and RBAC rules? We built a `setup-namespace` helm chart to create the namespace, RBAC roles and other requirements for every valid combination of clusters, namespaces and gsuite groups.

![](https://miro.medium.com/max/1400/1*7FyLXqW0nM501eXIQzXPKQ.png)

Function to generate chart manifests

Applying the configurations was not much more complicated. Combining a gcloud service account and GCP labels on the projects & clusters, we could dynamically identify all of the clusters and how to connect to them. The next step was to simply apply the relevant generated manifests to the matching clusters.

![](https://miro.medium.com/max/1400/1*y3gfa9Pa8yGC_JE4pPZ1tQ.png)

logic for planning, applying changes to clusters

All of the layers of complexity are simplified to a several, repeated Concourse jobs. An update to the cluster configurations is tested with a \`pr-test\` and must pass a `apply-approved-prs` job, which consists of applying and verifying the new manifest on the Sandbox clusters. If a PR can’t be applied without failing, we want to know about it before it is merged into master. After changes are merged in, a `create-release-tag` job is run and creates a git release tag. Anytime the `release-plan` or `release-apply` jobs are run, they will look for changes between this release tag and what is on the cluster already. Only changed manifests are applied to keep the changes to a minimum.

![View of minimal Concourse jobs required to manage all the clusters.](https://miro.medium.com/max/1400/1*OXxOGcg784krrfP2qgVDgg.png)

_Goal 3) Rebuild With Newest High Availability Features_

Since early 2018, multiple new GKE and GCP features were released that could dramatically decrease our outages by implementing highly resilient networking and GKE masters. We redesigned our networking between GCP projects to be a hub and spoke model taking advantage of new products like HA VPN Gateways.

- Highly resilient external and internal networking
- Private, not public, GKE clusters
- Istio enabled
- `gke-security-group` RBAC and IAM gsuite group enabled
- Regional multiple masters, instead of a single zonal master

# The Final Story

What has this journey taught us? Running on the latest technology absolutely allows us to adapt quickly and continue to actively reduce our cloud spending. No debate there. However, it comes with a hidden cost of any solution never solving the problem for long. In our four year’s experience with Kubernetes, we’ve never regretted the choice. Our team’s unofficial motto “ _become comfortable with change_” has never been more true.

[Some rights reserved](http://creativecommons.org/licenses/by/4.0/)

## [More from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----67afb437716c--------------------------------)

Gannett is the largest local news organization in America. Our brands include USA TODAY and 250+ newsrooms spanning 46 states. Gannett’s vastly expanded local-to-national platform reaches over 50% of the U.S. digital audience, including more Millennials than Buzzfeed.

[Read more from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----67afb437716c--------------------------------)