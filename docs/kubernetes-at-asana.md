# How we use Kubernetes at Asana

By
[Tony Liang](https://blog.asana.com/author/tonyliangasana-com/ "Posts by Tony Liang")
Feb 11, 2021

At Asana, we use Kubernetes to deploy and manage services independently from our monolith infrastructure. We encountered a few pain points when initially using Kubernetes, and built a framework to standardize the creation and maintenance of Kubernetes applications, aptly named KubeApps. Over the last two years, the Infrastructure Platform team has been making improvements to the KubeApps framework to make deploying services easier at Asana. In this post, we’ll explore the problems we aimed to solve in Kubernetes and the ways we did that with the KubeApp framework.

### Background

In an early blog post about Asana’s [legacy deployment system](https://blog.asana.com/2017/06/asana-server-deployment-pragmatic-approach-maintaining-legacy-deployment-system/), we discussed how our core infrastructure runs as a monolith of AWS EC2 instances with custom configuration scripts for each service. This system was not scalable: Pushing new code to the monolith required all hosts to be reconfigured; adding new services required a new set of custom configuration scripts that were difficult to maintain and carried a risk of adding instability to deployments to the monolith.

When we considered how we wanted to build out the infrastructure behind [Luna2](https://blog.asana.com/2017/08/performance-asana-app-rewrite/), we decided to use Kubernetes to enable us to deploy and manage services independently from the monolith. The container orchestration solution worked well for us, and our broader Engineering team expressed enough interest in using Kubernetes for other systems at Asana that we generalized these tools to allow other teams to build and deploy containerized services.

### Kubernetes at Asana

Using Kubernetes has made managing services much easier because we are able to delineate clear boundaries between processes and the environments where they run. As a result, application developers are able to build and update services without worrying about negatively impacting other existing services.

While Kubernetes helped us solve challenges with scaling out our infrastructure beyond the EC2 monolith, we still needed to handle some problems that Kubernetes did not solve for us out of the box at the time.[1]

- AWS resource management is beyond the scope of Kubernetes. Developers would need to provide pre-created ASGs and provision them to run Docker containers as well as any other resources that may be required for the application to run.
- Instrumenting common tooling like metrics collection and secrets permissions across all services was tedious and we would find ourselves repeating work like setting up ELBs and managing configuration groups for each application.
- Continuous delivery is not built into Kubernetes. At the Kubernetes abstraction layer, the application logic is expected to be packaged into images readily available in a container registry.

In order to solve these pain points, we built a framework to standardize the creation and maintenance of Kubernetes applications, aptly named KubeApps. The KubeApp framework was initially designed to handle all aspects of deployment from gathering the necessary resources to updating DNS during deployments and providing smooth transitions between code versions. Each KubeApp is defined by a set of hardware requirements, pod definitions and service specifications.

### Configuration as code

All of our KubeApp configurations exist as Python code. This allows for programmatically generated configurations and ensures consistency via shared libraries. These configurations allow us to declare and configure external resources needed for the KubeApp.

In the image below, we have the specifications for a sample webapp and how it can be represented in our KubeApp configuration system. The deployment will make requests for hardware, ELB and image builds. We provide default specifications to handle DNS configurations and include pod definitions to handle metrics and logging.

![](https://blog.asana.com/wp-content/post-images/sample-kubeapp-fixed-1024x492.png)

### Docker images with continuous deployment

Application code is packaged into Docker images using either [Bazel Docker images](https://github.com/bazelbuild/rules_docker) or via traditional Dockerfiles. Our KubeApps framework will build and push these images as part of our CD pipeline, so images for any service will be readily available on our container registry.

We use [Bazel](https://bazel.build/) to build and manage dependencies for most of our code at Asana. It’s particularly useful with containerized applications because it enforces explicit dependency declarations and packages third party libraries into built executables. This provides a deterministic output when setting up container environments for each application.

### One Cluster per KubeApp

A decision that many teams come across when building on Kubernetes is how to divide services up between clusters[2]. Many services deployed to a single Kubernetes cluster is cost effective and allows for efficient administration because the services share master node resources and operational work only needs to happen once (i.e. Kubernetes version upgrades, deployments). However, the single cluster model does not provide hard multitenancy and lends itself to being a single point of failure. There are also scaling limitations of a single Kubernetes cluster, so a split will need to be considered once a deployment grows to a certain size.

At Asana, each KubeApp is deployed in its own Kubernetes cluster via AWS EKS. This approach gives us strong guarantees on application security and resilience. Since each cluster is responsible for one service, we don’t need to worry about resource contention between services and the impact of a cluster failure is limited to a single service.

Managing multiple clusters can be tricky because the default tooling for Kubernetes is only able to interface with one cluster at a time. However, our team has built tooling in our KubeApps framework to manage multiple clusters at once. We’ve also found that this model empowers individual KubeApp owners to take on cluster management work independently (i.e. upgrading nodes, scaling deployment sizes, etc).

### KubeApp Deployment Workflow

![](https://lh5.googleusercontent.com/leyZ7uA7vja_BRj_XrcY5Xr65H9grKvfekSSQwe_wc2fuy4AL-0Loepb_33zUq7TVBW39H7zPKI_rg4KPmromyyHlX4YTrYzrPpbbP55AY-4tpqQJYVNaPav_OSRmrmo3rRqR7s)

KubeApp deployments are driven through a central management hub that we named “kubecontrol” which can be configured to run updates automatically via crons or manually by developers. The steps that happen during a KubeApp deployment are as follows:

1. A KubeApp update or create is triggered via the command line on kubecontrol
2. From the application specs, we request the set of resources required for the KubeApp (ASGs, spot instances, etc) and a new EKS cluster is created.
3. We make a request to our image builder service to compile the docker image at the given code version. The image builder will compile the code for the KubeApp and commit the image to ECR (Elastic Container Registry) if it does not already exist there.
4. Once all the required resources are built, we hand off the component specifications to the Kubernetes cluster to pull required docker containers from ECR and deploy them onto the nodes.

Full updates of KubeApps are blue/green deployments and require a new EKS cluster AWS resources to be launched and configured. Once the newly launched KubeApp is verified to be in a working state, we switch load to the new cluster and tear down the old one. KubeApps also have the option for a rolling-update which will just update images on a running cluster. This allows for quick seamless transitions between code versions without needing to spin up an entirely new cluster.

### KubeApp Management Console

Until recently, the only way for a developer to directly monitor or manage a KubeApp was to ssh into kubecontrol and interface with their app via the CLI. Information about deployments was not easily searchable so users would need to search through logs to figure out when a specific version of code was deployed to a KubeApp. In order to provide more clarity and observability to KubeApp users, we’ve built out a KubeApp Management Console (KMC) that would be responsible for recording historical information about past deployments. Eventually, we would like to use this interface to provide a centralized web-based interface for users to interact with KubeApps.

### State of KubeApps today and what’s next

We currently have over 60 KubeApps running at Asana that support a wide variety of workloads ranging from [release management](https://blog.asana.com/2021/01/asana-engineering-ships-web-application-releases/) to [distributed caching](https://blog.asana.com/2020/09/worldstore-distributed-caching-reactivity-part-1/#close). We still maintain the monolith of EC2 instances, but we are in the process of reimagining these services as containerized processes running in KubeApps.

The Infrastructure Platform team will continue to iterate and build new functionalities on the KubeApps framework. In the future, we plan to extend support for more types of architectures (ARM64) and infrastructure providers (AWS Fargate). We also plan to build tools to support a better development experience on KubeApps by making them launchable in sandboxed or local environments. These changes, among others, will enable the KubeApps framework to be extensible for any workload our engineers may need.

[1] Kubernetes does have a [Cloud Controller Manager](https://kubernetes.io/docs/concepts/architecture/cloud-controller/) that can manage node objects, configure routes between containers, and integrate with cloud infrastructure components.

[2] Some discussion about the tradeoffs between many clusters vs few clusters are available in this article from learnk8s.io: [Architecting Kubernetes clusters — how many should you have?](https://learnk8s.io/how-many-clusters)

Special thanks to

Eldar Bogdanov, Tony Liang, Kriti Singh, Natan Dubitski, Ashley Waxman, Steve Landey, Misha Kestler, Susan Thomakos

Would you recommend this article?
Yes /
No
