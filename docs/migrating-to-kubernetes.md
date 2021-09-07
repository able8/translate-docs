# Migrating to Kubernetes

March 30, 2020

The reasons to move to Kubernetes are many and compelling. This post doesn’t make the case that you should migrate, but assumes you have already decided that you want to. When you’re clear on what you want to do and why you want to do it, the questions of “When?” and “How?” become your focus. What follows centers on the question of how to approach making Kubernetes the platform on which your workloads thrive.

![Migrating to Kubernetes, captained by Sensu mascot Lizy](http://images.ctfassets.net/w1bd7cq683kz/uaexd5dpIUsM56gpKpOZp/3ef2131b76d60b37bc32ef2664e5d1ca/Picture1-2.png)

How to migrate depends, to an extent, on what you want to migrate from. The primary consideration is whether your existing infrastructure runs workloads in containers. If so, you’re already off to a quick start because you won’t have the containerization step to complete. Otherwise, you have a clear place to start.

Further, your considerations will vary if you are running via a serverless solution or other cloud platforms as services. These platforms provide value by making some decisions for you. In some cases and some ways, your solution is specific to the platform. Migrating from infrastructure as a service or your on-premises infrastructure has its own set of concerns. Moving from a different container orchestration engine to Kubernetes is another possibility with some unique characteristics.

Although where you’re coming from will account for some differences in how you proceed, there are common concerns you’ll need to consider in your journey. This post touches briefly on differences in migrating from differing platforms and focuses mostly on the decisions you’ll need to make in any case.

## Kubernetes: What and why

Kubernetes is the open source container-orchestration system for automating the deployment, scaling, and management of containerized applications. While we won’t go into the details of how Kubernetes works here, you can learn more by checking out these posts from the Sensu team: [Kubernetes 101](https://sensu.io/blog/kubernetes-101) and [How Kubernetes works](https://sensu.io/blog/how-kubernetes-works). We’ll assume from here on out that you’ve decided you’re ready to migrate (or at the very least, you’re ready to start reasoning about it).

Your approach to using Kubernetes depends on your knowledge of what it provides. It also depends on why you want to use it. The ideas and opinions in this post are general guidelines. They don’t substitute for your judgment in making sure your decisions serve your purposes (you know your applications better than we do, after all!). We suggest you think carefully about migration in your own context.

## Table stakes: Containerization

If your application workloads aren’t already running in containers, step 0 in your migration will be to change that, and you’ll likely want to use Docker. Docker isn’t the only way to containerize, but it’s usually the obvious choice given that it’s intuitive, well-supported, and in broad use.

You use Docker to create images that support your applications with all of their dependencies. This is best done using an automated [continuous integration](https://www.atlassian.com/continuous-delivery/continuous-integration) pipeline that includes pushing the versioned images to a [Docker registry](https://docs.docker.com/registry/introduction/) (which can be private). With these images in a registry, you’re ready to move into using Kubernetes.

Because this post presents Kubernetes migration rather than Docker fundamentals, there won’t be more here about this step. There are many other [great resources for getting familiar with Docker](https://docs.docker.com/get-started/).

## There and back again: A journey to Kubernetes

With the pieces of your application running in containers, it’s time to think about your strategy for migration.

Your primary concerns in architecting your application for Kubernetes execution are redundancy, resiliency, services, networking, monitoring, and persistent state.

### Redundancy

Before migrating containerized application workloads onto your Kubernetes cluster, you need to have a cluster. Questions to answer include where to run this cluster and capacity planning for nodes in the cluster, including how many and how powerful.

These are infrastructure questions.

Using a cloud provider with a Kubernetes offering is usually the best choice on the where-to-run question. Which cloud provider and capacity planning on nodes are topics for other posts.

Beyond the infrastructure level, you’ll need to plan for redundancy at the application level. How many replicas do you need for a given pod? You will answer such questions by telling Kubernetes how to run your pods via deployments and/or StatefulSets.

### Resiliency

Kubernetes takes care of monitoring the health of pods to make sure they are in operating condition. There are characteristics of containers that can indicate unhealthy pods, but Kubernetes needs your help knowing if your application is healthy in a container. To enable Kubernetes to take down pods that aren’t performing, replace them with new ones, and know when the new ones are ready, you need to contribute. You tell it how to know the state of your pods with [liveness, readiness, and startup probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).

### Networking

Having put your processes into containers doesn’t isolate them. You still have system components that need to communicate with other system components and with the outside world.

Ideally, you want most of the communication within your cluster to be either between containers within a pod or via [services](https://kubernetes.io/docs/concepts/services-networking/service/) for communications crossing pods.

You have [several options for exposing your services to the outside world](https://medium.com/google-cloud/kubernetes-nodeport-vs-loadbalancer-vs-ingress-when-should-i-use-what-922f010849e0) and these can be confusing. With Kubernetes services, you specify a service type. Some of the types are directly exposed to the outside world. The LoadBalancer service type makes use of the hosting platform to set up a load balancer directly exposing your service. For example, LoadBalancer services in clusters in Amazon Web Services use [Elastic Load Balancers](https://aws.amazon.com/elasticloadbalancing/).

Instead of exposing a service directly by using a type, [ingress resources](https://kubernetes.io/docs/concepts/services-networking/ingress/) can be set up to control access to services or other resources. Ingress gives you more control over how you expose your application but can require more thought and add complexity.

### Monitoring

Some operators, even before moving to Kubernetes, have embraced knowing that there are multiple system components with many replicas and that traditional monitoring of individual pieces isn’t going to cut it. There are new challenges with bringing new instances online and offline frequently (not to mention the fact that those instances are distributed). For a deeper dive into those challenges as well as popular patterns for [monitoring Kubernetes, check out this series from Sensu CTO Sean Porter](https://sensu.io/resources/whitepaper/integrations-for-cloud-monitoring-and-alerting#data-sources).

### Persistent state

Migrating databases can be the most challenging part of moving systems. The first and most obvious hurdle is the moving of data. Large databases can be hard to migrate. More fundamental, though, is deciding where the data should go.

Running your application in Kubernetes doesn’t necessarily mean your databases have to live in Kubernetes. The major cloud providers have offerings for running your database as a managed cloud service. Generally, that’s a pretty good way to offload much of the work required in maintaining databases and works well with pods in Kubernetes simply connecting to managed databases. Cloud providers offer custom cloud-native databases, but also services for running managed instances of well-known databases like SQL Server, Oracle, MySQL, PostgreSQL, and more. The most straightforward way to migrate databases is usually to move to something you’re already using, potentially just in a different place.

Alternatively, virtual private networks (VPNs) can be used to connect to on-premises databases if you want to avoid migrating data altogether.

Databases in Kubernetes are also an option, though. This gives you more control in exchange for you needing to manage your persistent storage and database software containers in your cluster. Kubernetes does support attaching persistent storage volumes to containers.

## Migrating from self-controlled infrastructure

Deploying your systems on-premises is certainly not the same thing as running in a public cloud via infrastructure as a service. For the purposes of this post, though, consider them as close enough to the same to treat them together. With infrastructure as a service, you execute in the cloud with maximum control over how and where your workloads run. You have control over machines. Configuring, maintaining, and supporting physical hardware on-premises is obviously different, but that’s not a focus of an application-centric view.

### Toward immutable deployment units

If you have full control over your operating environments, you have the liberty to customize hardware (virtual or otherwise) to suit your needs. This enables useful scenarios and can help in troubleshooting but comes with great peril. Servers tweaked to get things running may have undocumented configurations. This makes the setup hard to repeat and servers hard to replace. Applications and support can behave nondeterministically. When you move to Kubernetes (or any platform as a service), you need to get accustomed to immutable runtime environments.

Technically, you can mutate the state of running containers in Kubernetes pods, but it’s generally a bad idea. This is because Kubernetes replaces pods to maintain the desired state such that these changes in a container will be lost. Thus, treating deployed containers as immutable is a good idea. Further, enforcing policies via Kubernetes to make filesystems in containers read-only is supported and generally advisable.

This means teams need to shift their view of environment infrastructure as something set up by an operations role to something constructed from repeatable scripts and manifests.

### Planning the migration

Chances are if you’re running on-premises currently, you’re going to be containerizing first before moving anywhere. While containerizing, you’ll need to think through the application dependencies and how to put together images that serve your applications. At the same time, you’ll plan out manifests for how to get the containers to work together in a Kubernetes cluster.

Database migration is most challenging when moving from on-premises. Cloud providers do have tools that can help in moving the information.

## Migrating from cloud platforms as services

If you’re using services specific to your cloud provider(s), you have some decisions to make. For example, you can continue to use proprietary databases like CosmosDB or DynamoDB if you construct your cluster in the same cloud as your existing application. Given the ubiquity of Kubernetes and that it’s not specific to any provider, you might want to consider moving away from proprietary dependencies. There is no right answer, only questions to make sure you consider.

## Onward to migration

As you move forward in reaping the benefits of the production-grade container orchestration platform that is Kubernetes, remember that your system is yours and your context is unique. Every decision you make needs to serve your purposes so that you can better serve your users and your organization as a whole. Happy migrating!

_Want to learn more? Watch Sensu CEO and co-founder Caleb Hailey’s [webinar on filling the gaps in your Kubernetes observability strategy](https://sensu.io/resources/webinar/the-top-7-most-useful-kubernetes-apis-for-comprehensive-cloud-native-obsevability), in which he goes over the 7 most useful APIs for cloud-native observability, demonstrating how to get more context into what’s going on with your K8s clusters._
