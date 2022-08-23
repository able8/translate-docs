# The Journey of Adopting Cloud-Native Development

Jul 2, 2020

11 Minute Read

Table of Contents

- [Level 0 – Traditional Development - Kubernetes is an ops topic](http://loft.sh#level-0--traditional-development---kubernetes-is-an-ops-topic)
  - [Advantages](http://loft.sh#advantages)
  - [Disadvantages](http://loft.sh#disadvantages)
- [Level 1 – Manual Continuous Deployment (CD) – Developers run CD pipelines to Kubernetes manually](http://loft.sh#level-1--manual-continuous-deployment-cd--developers-run-cd-pipelines-to-kubernetes-manually)
  - [Advantages](http://loft.sh#advantages-1)
  - [Disadvantages](http://loft.sh#disadvantages-1)
- [Level 2 – Cloud-native CD – Specialized CD tools run pipelines to Kubernetes automatically](http://loft.sh#level-2--cloud-native-cd--specialized-cd-tools-run-pipelines-to-kubernetes-automatically)
  - [Advantages](http://loft.sh#advantages-2)
  - [Disadvantages](http://loft.sh#disadvantages-2)
- [Level 3 – Cloud-native development – Development takes place inside the Kubernetes cluster](http://loft.sh#level-3--cloud-native-development--development-takes-place-inside-the-kubernetes-cluster)
  - [Advantages](http://loft.sh#advantages-3)
  - [Disadvantages](http://loft.sh#disadvantages-3)
- [A closer look at the access to Kubernetes](http://loft.sh#a-closer-look-at-the-access-to-kubernetes)
- [Conclusion](http://loft.sh#conclusion)

Adopting Kubernetes is a process that many companies are currently going through. The introduction of Kubernetes as infrastructure technology can take some time. ( [It took almost 2 years for Tinder to complete its migration to Kubernetes](https://medium.com/tinder-engineering/tinders-move-to-kubernetes-cda2a6372f44).) The transition of development processes to fully cloud-native development is often an even longer process that comprises several incremental steps.

To illustrate the different stages of transition, I want to use the analogy of [autonomous driving](https://www.synopsys.com/automotive/autonomous-driving-levels.html), where different “levels” also describe the technological advancement and sophistication. Luckily and in contrast to autonomous driving, the highest level of cloud-native development is already reachable and reality today.

By looking at the following description of the different cloud-native levels, you can classify your current [development workflow with Kubernetes](http://loft.sh/blog/kubernetes-development-workflow-3-critical-steps/?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development) and even see what is ahead of you and what your next possible steps might be.

## [\#](http://loft.sh\#level-0--traditional-development---kubernetes-is-an-ops-topic) Level 0 – Traditional Development - Kubernetes is an ops topic

At level 0, [Kubernetes is not a topic for developers yet](http://loft.sh/blog/is-kubernetes-still-just-an-ops-topic/?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development) although it might already be used for production workloads. While developers may know that Kubernetes is used to run their applications after they are finished, they are not in touch with Kubernetes during their everyday work. They only use container solutions such as Docker, if they are working with containers at all.
This means that the developers are working on their local computers with their traditional technologies and hand-over the software for production to operations departments after completion.

### [\#](http://loft.sh\#advantages) Advantages

Traditional development, as the name suggests, has been the standard for a long time in the past and even though it is “Level 0”, it does not only have downsides.

The main advantage of this status quo is that **nothing must be changed**, so engineers do not have to learn new things and all **workflows can remain the same**.

Another benefit is that **developers get direct feedback if their application is running** because everything runs on their local machines and nothing has to be executed in a Kubernetes environment first, so errors, for example, are visible pretty much immediately.

### [\#](http://loft.sh\#disadvantages) Disadvantages

For modern development methods, level 0, however, has many disadvantages. At first, it is **contrary to the DevOps approach** by creating a clear separation of development and operations. As a result, **developers are not able to take on full responsibility** for their software and **“works on my machine”-problems** emerge very easily when the locally (without Kubernetes) developed software runs in the real-life Kubernetes environment for the first time.

This **lack of realism** of the runtime environment leads to a huge challenge for operations managers who are responsible for getting the application to run. This is aggravated by the fact that they **cannot easily replicate the developers’ local runtimes** for testing and debugging because they may be configured in many different ways.

Generally, only executing applications on individual local machines with different configurations is disadvantageous. It now is the responsibility of the **developers to set up the environment**, which may be **a lot of work** and **must be repeated** for every new team member or every new computer used.

And even worse, the **execution in local runtime environments is simply not possible for all applications** anymore as they become more and more complex and may require very special or a lot of computing resources. This is especially true for machine learning and artificial intelligence software running on GPUs or requiring a lot of computing power.

## [\#](http://loft.sh\#level-1--manual-continuous-deployment-cd--developers-run-cd-pipelines-to-kubernetes-manually) Level 1 – Manual Continuous Deployment (CD) – Developers run CD pipelines to Kubernetes manually

Developers at level 1 know that Kubernetes is their target platform but do not have direct access to it. They are rather enabled to push their code to a Kubernetes runtime via a pre-defined pipeline that is usually set up by someone else, e.g. a DevOps engineer.

From a workflow perspective, developers push their code to a code repository and then manually trigger a pipeline with tools such as [Gitkube](https://github.com/hasura/gitkube) or [Spinnaker](https://github.com/spinnaker/spinnaker) that deploys the code in a Kubernetes environment, usually in a remote cluster in the cloud.

For the developers, this means that Kubernetes is the target platform but it is concealed by the pipeline.

### [\#](http://loft.sh\#advantages-1) Advantages

A clear advantage of this approach is that the software is executed in a **very realistic environment in the cloud** already during development. The developers are so enabled to test their code in Kubernetes **without having to manage this environment as it is centrally controlled and maintained** by someone else.

This makes it also **relatively easy to use** and **new developers can get started very fast** because they only need to get the right to trigger the pipeline for their code.

Another benefit of level 1 development is that **bugs and errors can be easily communicated and replicated** as everyone in the team works with the same CD environment and configurations instead of an individual local environment.

Finally, since the target environment is a Kubernetes cluster in the cloud, it is **possible to even run very complex and computing intense software** without limitations.

### [\#](http://loft.sh\#disadvantages-1) Disadvantages

The transition from level 0 to level 1 is substantial and the strengths of the one level are the weaknesses of the other. For level 1, this is particularly important for the development speed.

While feedback is nearly immediate for level 0 local development, it becomes **very slow** for level 1 development based on manual pipelines. This is because the developers have to **manually execute the whole pipeline every time they change a file**, which can cause several minutes of pure waiting time for every change.

Additionally, since a shared and centrally managed platform is used, developers are also **not allowed to change or configure anything** themselves. They rather have to inform the cluster administrator which leads to **problematic single point of failure** if the admin is not available.

Finally, in spite of the general flexibility of the runtime in terms of scale, the whole process is **not platform-independent**, so a transition from one environment to the other (e.g. from AWS to Azure) becomes a complicated process.

## [\#](http://loft.sh\#level-2--cloud-native-cd--specialized-cd-tools-run-pipelines-to-kubernetes-automatically) Level 2 – Cloud-native CD – Specialized CD tools run pipelines to Kubernetes automatically

The development concept at level 2 is essentially the same as at level 1. The code is deployed via a pipeline to a Kubernetes cluster. However, at this stage, developers are using special CLI tools, such as [Skaffold](https://github.com/GoogleContainerTools/skaffold), [Draft](https://github.com/Azure/draft), or [Tilt](https://github.com/tilt-dev/tilt) that detect the file changes by the developer and then automatically trigger the pipelines.

Since these tools are specifically made for this use case, they also have additional features that facilitate development with Kubernetes, such as port forwarding.

Another very important distinction of this level is that developers have direct access to Kubernetes for the first time. While it is not strictly necessary, the standard case for the Kubernetes access with these tools is to use a local Kubernetes cluster, i.e. a Kubernetes cluster started with tools such as [minikube](https://github.com/kubernetes/minikube), [kind](https://github.com/kubernetes-sigs/kind) or [MicroK8s](https://github.com/ubuntu/microk8s) on the local computer of the developer.

### [\#](http://loft.sh\#advantages-2) Advantages

Similar to the pipeline process of level 1, development with cloud-native CD pipelines also leads to tests in an **environment close to production** due to Kubernetes as underlying technology.

This also lets other team members **replicate problems and bugs relatively easily** as everyone is at least using the same infrastructure technology, even if it runs locally with some minor differences.

Another improvement compared to the manual pipeline approach is the **automatic execution of the pipeline** and the **additional features for development** that speed up the dev process a little bit.

### [\#](http://loft.sh\#disadvantages-2) Disadvantages

However, level 2 development is still a **relatively slow process** compared to traditional development because pipelines still have to be executed resulting in waiting times even if this process is triggered automatically and is somewhat faster due to the local runtime environment.

Additionally, developers are now in contact with Kubernetes and usually have to **set up and manage it themselves** on their local computers. This, of course, comes with some **extra effort and requires additional knowledge** as the configuration of the local cluster can be tricky in some cases.

Another downside of using a local runtime environment is the **limitation in terms of computing resources** that can make it infeasible again for some more advanced applications.

## [\#](http://loft.sh\#level-3--cloud-native-development--development-takes-place-inside-the-kubernetes-cluster) Level 3 – Cloud-native development – Development takes place inside the Kubernetes cluster

At level 3, the development of the software takes place inside a Kubernetes cluster without the need of running CD pipelines on every change. Instead, the developer uses special tools such as [DevSpace](https://github.com/devspace-cloud/devspace) or [Okteto](https://github.com/okteto/okteto) that simulate traditional development in a cluster as much as possible.

These tools recognize file changes and synchronize them to the container filesystem inside the Kubernetes cluster. This allows the application to be updated instantly (hot reload) and images will only be rebuilt when necessary, i.e. after the image was changed.

Besides this hot reloading feature, the cloud-native development tools provide a port forwarding feature to allow development on localhost and give developers full terminal and log access.

Another important aspect of level 3 development is that the developers have direct access to a remote [Kubernetes dev environment](http://loft.sh/blog/kubernetes-development-environments-comparison/?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development) and do not use a local Kubernetes cluster anymore. To enable this access without having a cluster for each developer, tools such as [Loft](https://loft.sh) provide a centrally managed, [multi-tenancy Kubernetes platform](https://loft.sh/features/kubernetes-multi-tenancy?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development) for the developers.

### [\#](http://loft.sh\#advantages-3) Advantages

The goal of level 3 development is to combine the best of both worlds, traditional development speed with Kubernetes’ realism, scalability, and replicability. Due to the hot reloading feature instead of the pipeline approaches of levels 1 and 2, cloud-native development becomes **much faster** and is only slightly slower than traditional local development without Kubernetes.

Since all the development processes and commands needed by the developers are built-in or can be pre-configured by a Kubernetes expert in a team, the **complexity to use this approach is relatively low** even though the developers have a direct access to Kubernetes.

This direct access in a cloud environment also allows **very realistic testing** and is **not limited in terms of computing resources**, which makes this form of development feasible and efficient even for complex applications.

Finally, since the dev tools are not hardwired to the cluster, it is **easily possible to switch the runtime environment** (e.g. switch from local to a remote cluster or switch between different remote clusters), which **prevents a lock-in effect**.

### [\#](http://loft.sh\#disadvantages-3) Disadvantages

As most advanced level of cloud-native development, level 3 development is also most efficient for many cases. It is only **slightly slower than traditional development** and **some workflows have to be adapted**.

It also requires **some setup effort once when introduced** to determine the optimal configuration for the tools that can then be shared in a whole team.

## [\#](http://loft.sh\#a-closer-look-at-the-access-to-kubernetes) A closer look at the access to Kubernetes

Since the goal of the development processes described in this post is to run great software in Kubernetes, Kubernetes needs to be part of every development process at some point. By advancing from one level to the next, the touchpoint of the software with a (realistic) Kubernetes environment becomes earlier.

At level 0, Kubernetes is purely an operations topic and the developers are usually not in touch with it at all. They might only get some feedback if their software is not running properly in production due to problems that were not considered when development took place locally.

At level 1, Kubernetes is still an ops topic, but the developers have indirect access to it during development via a CD pipeline. However, they still do not need to manage anything and do not have the right to configure anything.

Starting from level 2, Kubernetes becomes a development topic. That means that the developers also need to get direct access to it. In level 2, this usually means that they start and manage their own local Kubernetes clusters while in level 3, they get access to a shared cluster via specific tools such as [Loft](https://loft.sh) or [kiosk](https://github.com/kiosk-sh/kiosk).

> _In separate posts, [I compare local and remote environments for development](https://medium.com/swlh/local-cluster-vs-remote-cluster-for-kubernetes-based-development-6efe2d9be202) and [the different methods for providing direct access to Kubernetes for developers in the cloud](http://loft.sh/blog/individual_kubernetes_clusters_vs-_shared_kubernetes_clusters_for_development/?utm_medium=reader&utm_source=other&utm_campaign=blog_the-journey-of-adopting-cloud-native-development)._

## [\#](http://loft.sh\#conclusion) Conclusion

I hope this post was helpful to understand where you and your team are on the cloud-native journey and to get an impression about what next steps might await you.
However, while many companies actually go through the different levels one after the other, I want to note that it is absolutely possible to skip levels and to combine different approaches with each other to get maximum efficiency.

In the end, I am still convinced that many companies would benefit from going the full distance to level 3, especially if they really need the structural benefits of Kubernetes, such as scalability.

At the moment, I believe level 3 is the highest level on the cloud-native journey in software development but, as technology advances, I would not be surprised if I had to add a level 4 or 5. So, let’s stay excited about what the next steps are coming in this fast-paced ecosystem.

https://loft.sh/blog/the-journey-of-adopting-cloud-native-development
