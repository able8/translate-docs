# What we learned after a year of GitLab.com on Kubernetes

John Jarvis

Sep 16, 2020·10 min read· [Leave a comment](http://about.gitlab.com#disqus_thread)

* * *

![](http://about.gitlab.com/images/blogimages/a_year_of_k8s/nico-e-AAbjUJsgjvE-unsplash.jpg)

For about a year now, the infrastructure department has been working on migrating all services that run on GitLab.com to Kubernetes. The effort has not been without challenges, not only with moving services to Kubernetes but also managing a hybrid deployment during the transition. We have learned a number of lessons along the way that we will explore in this post.

Since the very beginning of GitLab.com, servers for the website have run in the cloud on virtual machines. These VMs are managed by Chef and installed using our [official Linux package](http://about.gitlab.com/install/#ubuntu).
When an application update is required, [our deployment strategy](https://gitlab.com/gitlab-org/release/docs/-/blob/master/general/deploy/gitlab-com-deployer.md) is to simply upgrade fleets of servers in a coordinated rolling fashion using a CI pipeline.
This method, while slow and a bit [boring](http://about.gitlab.com/handbook/values/#boring-solutions), ensures that GitLab.com is using the same installation methods and configuration as our self-managed customers who use Linux packages.
We use this method because it is especially important that any pain or joy felt by the community when installing or configuring self-managed GitLab is also felt by GitLab.com.
This approach worked well for us for a time but as GitLab.com has grown to hosting over 10 million projects we realized it would no longer serve our needs for scaling and deployments.

## Enter Kubernetes and cloud native GitLab

We created the [GitLab Charts](https://gitlab.com/gitlab-org/charts) project in 2017 to prepare GitLab for deployments in the cloud and enable self-managed users to install GitLab into a Kubernetes cluster. We knew then that running GitLab.com on Kubernetes would benefit the SaaS platform for scaling, deployments, and efficient use of compute resources. At the time though there were still many application features that depended on NFS mounts that delayed our migration off of VMs.

The push for cloud native and Kubernetes gave engineering an opportunity to plan a gradual transition that removes some of the network storage dependencies on the application while continuing to develop new features. Since we started planning the migration in the summer of 2019, most of these limitations have been resolved and the journey to running all of GitLab.com on Kubernetes is now well underway!

## Running GitLab.com on Kubernetes

For GitLab.com we use a single regional GKE cluster that services all application traffic. To minimize the complexity of the (already complex) migration we focus on services that don't depend on local storage or NFS. While GitLab.com is running from mostly monolithic Rails codebase, we route traffic depending on workload characteristics to different endpoints which are isolated into their own node pools.

On the frontend these types are divided into web, API, git SSH/HTTPs requests, and Registry.
On the backend we divide our queued jobs into different characteristics depending on [predefined resource boundaries](http://about.gitlab.com/blog/2020/06/24/scaling-our-use-of-sidekiq/) that allow us to set Service-level Objective (SLO) targets for a range of different workloads.

All of these GitLab.com services are configured with the unmodified GitLab Helm chart, which configures them in sub-charts that can be selectively enabled as we gradually migrate services to the cluster.
While we opted to not include some of our stateful services such as Redis, Postgres, GitLab Pages, and Gitaly, when the migration to Kubernetes is finished it will drastically reduce the number of VMs that we currently manage with Chef.

## Transparency and managing the Kubernetes configuration

All configuration is managed in GitLab itself in three configuration projects using Terraform and Helm.
While we use GitLab to run GitLab wherever possible, we maintain a separate GitLab installation for operations.
This is done to ensure we do not depend on the availability of GitLab.com for deployments and upgrades of GitLab.com.

Even though our pipelines that execute against the Kubernetes cluster run on this separate GitLab deployment, the code repositories are mirrored and publicly viewable at the following locations:

- [k8s-workloads/gitlab-com](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-com): GitLab.com configuration wrapper for the GitLab Helm chart.
- [k8s-workloads/gitlab-helmfiles](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-helmfiles/): Contains the configuration for services that are not directly related to the GitLab application. This includes configurations for cluster logging and monitoring and integrations like PlantUML.
- [gitlab-com-infrastructure](https://gitlab.com/gitlab-com/gitlab-com-infrastructure): Terraform configuration for the Kubernetes and legacy VM infrastructure. All the resources necessary to run the cluster are configured here, including the cluster, node pools, service accounts, and IP address reservations.

[![hpa](http://about.gitlab.com/images/blogimages/a_year_of_k8s/hpa.png)](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-com/-/merge_requests/315#note_390180361)
Whenever a change is proposed, a public [short summary](https://gitlab.com/gitlab-com/gl-infra/k8s-workloads/gitlab-com/-/merge_requests/315#note_390180361) is displayed, with a link to detailed diff that an SRE reviews before applying changes to the cluster.

For SREs, we link to a detailed diff on our operations GitLab instance that has limited access.
This allows employees and the community, who do not have access to the operational project which is limited to SREs, to have visibility into proposed config changes.
By having a public GitLab instance for code, and a private instance for [CI pipelines](http://about.gitlab.com/stages-devops-lifecycle/continuous-integration/), we are able to keep a single workflow while at the same time ensuring we don't have a dependency on GitLab.com for configuration updates.

## The lessons we learned along the way

We have learned a few things along the way, lessons that we are applying to future migrations and new deployments into Kubernetes.

### Increased billing from cross-AZ traffic

![git egress](http://about.gitlab.com/images/blogimages/a_year_of_k8s/git_egress.png)
Daily egress bytes/day from the Git storage fleet on GitLab.com.

Google divides its network into regions and regions are divided into availability zones (AZs).
Because of the large amount of bandwidth required for Git hosting, it is important we are cognizant of network egress. For internal network traffic, egress is only free-of-charge if it remains in a single AZ.
At the time of writing this blog post, we deliver approximately 100TB on a typical work day for just Git repositories.
On legacy VM topology, services that were previously colocated on the same VMs are now running in Kubernetes pods.
This mean some network traffic that was previously local to a VM can now potentially traverse availability zones.

Regional GKE clusters provide the convenience of spanning multiple availability zones for redundancy.
We are considering [splitting the regional GKE cluster into single zonal clusters](https://gitlab.com/gitlab-com/gl-infra/delivery/-/issues/1175) for services that use a lot of bandwidth to avoid network egress charges while maintaining redundancy at the cluster level.

### Resource limits, requests, and scaling

![replicas](http://about.gitlab.com/images/blogimages/a_year_of_k8s/replicas.png)
Number of replicas servicing production traffic on registry.gitlab.com, Registry traffic reaches it peak at ~15:00UTC.

Our migration story began in August 2019 when we migrated the GitLab Container Registry to Kubernetes, the first service to move.
Though this was a critical and high traffic service, it was a good choice for the first migration because it is a stateless application with only a few external dependencies.
The first challenge we experienced was the large number of evicted pods, due to memory constraints on our nodes.
This required multiple changes to requests and limits. We found that with an application that increases its memory utilization over time, low requests (which reserves memory for each pod) and a generous hard limit on utilization was a recipe for node saturation and a high rate of evictions.
To adjust for this [we eventually decided to use higher requests and lower limit](https://gitlab.com/gitlab-com/gl-infra/delivery/-/issues/998#note_388983696) which took pressure off of the nodes and allowed pods to be recycled without putting too much pressure on the node.
After experiencing this once, we start our migrations with generous requests and limits that are close to the same value, and adjust down as needed.

### Metrics and logging

![registry-general](http://about.gitlab.com/images/blogimages/a_year_of_k8s/registry-general.png)
The Infrastructure department focuses on latency, error rates and saturation that have [Service-level objectives (SLOs)](https://en.wikipedia.org/wiki/Service-level_objective) that tie into our [overall system availability](https://gitlab.com/gitlab-com/dashboards-gitlab-com/-/metrics/sla-dashboard.yml?environment=1790496&duration_seconds=86400).

Over the past year, one of the major changes in the infrastructure department was improvements to how we monitor and manage SLOs.
SLOs allowed us to set targets on individual services which were monitored closely during the migration.
Yet even with this improved observability, we can't always see problems right away with our metric reporting and alerting.
For example, focusing on latency and error rates may not adequately cover all uses of the service that is being migrated.
We discovered this problem very early with some of the workloads that were moved into the cluster. This challenge was particularly acute when we had to validate features that do not receive many requests but have very specific configuration dependencies.
One of the key migration lessons was to also evaluate more than just monitoring metrics, but also logs, and the long-tail of errors in our monitoring.
Now for every migration we include a detailed list of log queries and plan a clear rollback procedures that can be handed off from one shift to the next in case of issues.

Serving the same requests on legacy VM infrastructure and Kubernetes simultaneously presented a unique challenge.
Unlike a lift-and-shift migration, running on legacy VMs and Kubernetes at the same time requires that our observability is compatible with both and combines metrics into one view.
Most importantly, we are using the same dashboards and log queries to ensure the observability is consistent during the transition period.

### Shifting traffic to the new cluster

For GitLab.com we maintain a segmentation of our fleet named the [canary stage](http://about.gitlab.com/handbook/engineering/#canary-testing).
This canary fleet services our internal projects, [or can be enabled by users](https://next.gitlab.com), and is deployed to first for infrastructure and application changes.
The first service we migrated started with taking limited traffic internally and we are continuing to use this method to ensure we are meeting our SLOs before committing all traffic to the cluster.
What this means for the migration is requests to internal projects are first routed to Kubernetes and then we slowly move other traffic to the cluster using HAProxy backend weighting.
We learned in the process of moving from VMs to Kubernetes that it was extremely beneficial for us to have an easy way to move traffic between the old and new infrastructure, and to keep legacy infrastructure available for rollback in the first few days after the migration.

### Reserved pod capacity and utilization

One problem we identified early was, while our pod start times for the Registry service were very short, our start times for Sidekiq took as long as [two minutes](https://gitlab.com/gitlab-org/charts/gitlab/-/issues/1775).
The long Sidekiq start times posed a challenge when we started moving workloads to Kubernetes for workers that need to process jobs quickly and scale fast.
The lesson here was while the Horizontal Pod Autoscaler (HPA) works well in Kubernetes for adapting to increased traffic, it is also important to evaluate workload characteristics and set reserved pod capacity, especially for uneven demand.
In our case, we saw a sudden spike in jobs which caused a large scaling event which saturated CPU before we could scale the node pool.
While it is tempting to squeeze as much as possible out of the cluster, after experiencing some initial performance problems we now start with a generous pod budget and scale down later, while keeping a close eye on SLOs.
The pod start times for Sidekiq service have improved significantly and now average about 40 seconds. [Improving the pod start times](https://gitlab.com/gitlab-org/charts/gitlab/-/issues/1775) benefited GitLab.com as well as all the self-managed customers using the official GitLab Helm chart.

After transitioning each service, we enjoyed many benefits of using Kubernetes in production, including much faster and safer deploys of the application, scaling, and more efficient resource allocation.
The migration benefits extend beyond GitLab.com. With each improvement of the official Helm chart, we provide additional benefits to our self-managed customers.

We hope you enjoyed reading about our Kubernetes migration journey. As we continue to migrate more services to the cluster you can read more at following links:

- [Why are we migrating to Kubernetes?](http://about.gitlab.com/handbook/engineering/infrastructure/production/kubernetes/gitlab-com/)
- [GitLab.com on Kubernetes](http://about.gitlab.com/handbook/engineering/infrastructure/production/architecture/#gitlab-com-on-kubernetes)
- [Tracking epic for the GitLab.com Kubernetes Migration](https://gitlab.com/groups/gitlab-com/gl-infra/-/epics/112)
