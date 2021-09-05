# Why Continuous Infrastructure Automation is no longer optional

Nov 19, 2019

# Containerized micro services and service orchestration disrupted how we build and run apps

We have now reached a critical point in cloud-based infrastructure evolution where several factors are piling on top of one another, calling for a game changing approach.

First of all, Cloud Native design (containerized micro-services and service orchestration) is disrupting how we build and run apps. Transforming the landscape at an incredible pace, Kubernetes has become the new de-facto standard for container orchestration within one year of its graduation from the CNCF (Cloud Native Computing Foundation).

Indeed, we now have the ability to run multiple instances of micro-services and autoscale them with zero downtime. In the same time, multi-cloud is now a proven trend in Enterprise infrastructure, with over 60% of IaaS users deploying workloads on several clouds.

While very powerful, it’s an insanely complex machinery, that necessitates huge skills to properly master both how it works and how to maintain it over time.

# As complexity now moves to the infrastructure level, infrastructure automation is no longer optional.

Traditional architecture meant using a few simple servers behind a load balancer to host and run our applications. Since then, environments gradually evolved into orchestrated containerized architecture. As a consequence, complexity now moves to the infrastructure level.

![](https://miro.medium.com/max/60/1*HkHW3S7YxTJ7XHDc6DA_ag.jpeg?q=20)

The logical outcome of that situation is that now, Infrastructure automation has become key to proper environment provisioning. And still, using IaaS consoles or simple deployment scripts, many teams jump to Kubernetes without automating their infrastructure first, putting their organization’s efficiency and security in jeopardy with unsafe, sometimes locked and/or un-repeatable environments.

Indeed, not only won’t “ClicOps” lead you ahead on the steep learning curve of using Infrastructure-as-Code, but it won’t either help you bridge the gap between your needs and a global context of DevOps skills shortening. Even more, those shortcuts often trigger additional problems due to the total lack of access right management or clear audit trails.

With organizations continuously trying to achieve a better time to value with quicker deployments, Infrastructure-as-Code provides the best option to provision, kill and relaunch environments on demand, using tools and code to provision and deploy servers and applications and have greater control of your infrastructure, especially in multi-cloud environments.

# So, how do you manage Kubernetes clusters properly with Infrastructure-as-Code?

Leveraging Hashicorp Terraform’s abilities, we are building [CloudSkiff](https://www.cloudskiff.com/) in a GitOps approach to manage your infrastructure-as-code properly, like a CI/CD for complex infrastructures that you need to provision, autoscale, kill and relaunch in a seamless and repeatable process.

Consequently, the key features of [CloudSkiff](https://www.cloudskiff.com/) are :

- Generate a proper code for your infrastructure : either by writing it from scratch with an assistant code generator, or by reusing complex modules written by other people, or through a tool that allows you to turn your click-based environments into proper code that you can reuse and share.
- Organize the collaboration around your code : set up straightforward process and boundaries with RBAC (Role Based Access Control), to avoid troubles. You can’t have your junior Dev, or even several separate senior people push`terraform apply` or `terraform destroy` commands without proper control over what's happening.
- Manage testing environments with a sandbox just to make sure that things are running smoothly before you deploy to production.
- Post deployment, use dashboards and monitoring to keep control over performance and costs across geographies and environments and optimize your setup.
- Monitor your Kubernetes environments across all accounts and all clouds and watch over performances and costs for a better optimization.
- Duplicate environments and switch seamlessly between cloud providers with our migration assistant

# Why we go for Kubernetes as a first step, and why you need a tool to manage the lifecycle of your infrastructure?

At [CloudSkiff](https://www.cloudskiff.com/) we decided to begin by supporting managed Kubernetes environments as a first step. Obviously, you've got to start somewhere, and going first for this framework that is (again) very powerful but insanely complex seemed to be a good way of providing support to the teams that might not have all the skills to master and maintain it overtime.

A Kubernetes Cluster is a living being, killed and reborn several times a day, always moving, always changing, always adapting/autoscaling… It’s highly sensible to version upgrades. Basic configuration mistakes can yield drastic consequences, which is why you need a pipeline to manage its lifecycle. We believe that DevOps teams need a new generation of dedicated tools to manage the lifecycle of their Kubernetes Clusters and that this is what it takes to help them focus on building and running their apps.

Question is : how do you implement that into existing workflows?

Indeed, most existing workflows are “touchy” once setup and nobody really wants to change that. This is why [Cloudskiff](https://www.cloudskiff.com/) is built to be fully integrated within existing workflows. It is directly linked to the Github repositories of your choice and can even be called by your existing CI/CD tool though an API.

**CloudSkiff is an Infrastructure as Code that allows** Terraform automation and collaboration for growing teams.

[**Join our Beta here**](http://www.cloudskiff.com?utm_source=medium&utm_medium=WhyCIAisnolongeroptional)

![](https://miro.medium.com/max/1400/1*-Vb5CXecRO7May4txjsTYw.png)
