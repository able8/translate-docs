# GitOps-based Policy Management: How to Scale in a Multi-Node, Multicloud World

January 19, 2021

Companies want to benefit from an increase in engineering velocity and innovation that comes with adopting Kubernetes and cloud native. But getting there requires a solid strategy for fleets that can scale on a variety of environments on multiple clusters that span on-premise, across multiple clouds and maybe even the edge. Organizations need an automated and consistent approach for managing clusters – one that they can use efficiently no matter how many clusters or clouds or other configurations are included in the mix.

In this post, we’ll walk you through the common challenges faced in multi-cluster environments, and discuss how [GitOps and effective policy management](https://go.weave.works/GitOps_on_AWS_Managing_GRC_for_k8s.html) simplifies large-scale Kubernetes deployments anywhere.

## Diverse cluster stacks add complexity

Kubernetes helps you innovate more quickly provided that you can manage the complexity of configuring platforms for it. One of the biggest challenges is maintaining and deploying consistent Kubernetes developer platforms used by multiple teams who need to operate across multiple environments and clouds.

Implementing Kubernetes involves much more than simply spinning up a cluster. On top of the base Kubernetes install, there are the core add-ons you need to run it, both within its infrastructure, including a way to monitor its health as well as the tools and applications your development teams require for CD pipelines, code tracing and logging. In addition to this, you may also need to consider any tools for specialized business requirements like machine learning or edge computing whose applications also need to be configured to work with Kubernetes.

Wherever your clusters run, fleets require an identical configuration that works. In addition to this, the configuration must also be quickly and easily upgraded with security patches and other required maintenance with minimal to no down time.

## Manage cluster configuration definitions with GitOps

Almost every element of Kubernetes uses declarative configuration: the cluster, and each component within it as well as the applications running on it. This allows for the platform configuration to be stored together in Git. With GitOps, reconcilers and software agents running inside the cluster ensures that the cluster and the apps that run on it are always up to date. In the case of a drift in configuration, an alert is triggered and the cluster can be automatically reconciled with a known state stored in Git.

![Principles_of_GitOps_2021.png](https://images.contentstack.io/v3/assets/blt300387d93dabf50e/blt2be043a8abafc0e7/6005f1ed3e567f1011da0a13/Principles_of_GitOps_2021.png)

This approach comes with the benefits of:

- Git-backed security guarantees provenance that includes a built-in audit trail of who did what.
- Increased reliability with secure built-in GitOps to automate the upgrade process across fleets.
- Scalable cluster management across fleets of clusters and applications.
- Save time and costs with complete platform configurations pushed to Git.
- Proactive monitoring with alerts on cluster configuration drift.

See, [Guide to GitOps](https://www.weave.works/technologies/gitops/) for more information on its benefits.

## What does Git-based policy look like?

At the core of reproducible, correct cluster configuration is the ability to manage policy with GitOps. This is a standard component in the [Weave Kubernetes Platform (WKP)](https://www.weave.works/product/enterprise-kubernetes-platform/) that can help meet business and regulatory compliance requirements more efficiently. Policies and rules can be set up by Platform teams to determine roles and permissions on who can commit changes to the base Kubernetes configuration.

![GitOps_Policy_Manager.png](https://images.contentstack.io/v3/assets/blt300387d93dabf50e/bltde36e4e0e2e513dd/6005f2116215cf0f9a18aed5/GitOps_Policy_Manager.png)

Role-based access control (RBAC) permissions can be checked in and confirmed in Git at commit time. Before any changes are applied feedback is provided.

The GitOps Policy Manager implements a set of Git-based rules built on top of the [Open Policy Agent (OPA)](https://www.openpolicyagent.org/) standard that is managed by pull request. This ensures that cluster changes are only initiated by the roles that are permitted to do so. Additional security including Role-based Access Control (RBAC) can also be applied for both teams and namespaces through the [WKP workspaces feature](https://www.weave.works/blog/wkp-team-workspaces-rbac).

See [Security with GitOps](https://www.weave.works/use-cases/security-with-gitops/) for more information.

## Self-service Kubernetes with guardrails

With policies and GitOps, you don’t have to choose between configuration consistency and tool choice. Instead, you can let your developers use tools they need to build their delivery pipelines. When all of your configuration information is stored in Git, entire cluster stacks can be pushed out in a consistent and scalable way, no matter how many clusters you are maintaining and deploying.

[Read how Mettle, a financial services company](https://www.weave.works/blog/case-study-mettle-leverages-gitops-for-self-service-developer-platform) used GitOps and Weaveworks solutions to implement a self-service Kubernetes.

## Managing fleets of clusters

Larger organizations need to manage fleets or multiple clusters on premise, on clouds and at the edge. Another obvious benefit of GitOps policy management is that it allows you to apply the same configurations to as many clusters as you want in an automated fashion.

With GitOps, there is no need to generate or apply custom YAML files for each cluster in your infrastructure. You can let tools do that hard and tedious work for you so that your team can focus on crafting the best configurations for your workloads, instead of applying them.

This approach not only makes managing dozens of clusters much more efficient, but it also increases security by reducing the risk of manual configuration mistakes or oversights.

### Achieving consistency between local environments and the cloud

Infrastructure in most organizations today includes a mix of on-premise resources and multiple public clouds. Managing clusters in a consistent way on heterogeneous infrastructure can be difficult and error prone without the simplicity and automation that GitOps and git-backed policy management provides.

Managing configuration with GitOps avoids the problem of having to maintain or secure bespoke and snowflake clusters. You can write a single set of configuration files, tailoring them where necessary to address the nuances of different clouds or on-premise clusters or for specific tools, and then deploy them to all parts of your infrastructure using a common workflow and a secure process.

## Streamlining access control across the organization

When manually managing fleets of clusters spread over multiple clouds, access control and user roles can quickly get messy. You would typically end up with a different set of roles and access-control policies for each cluster, which is a recipe for oversights and security holes.

With [GitOps-based management and team workspaces](https://www.weave.works/blog/wkp-team-workspaces-rbac) that are an integral part of the Weave Kubernetes Platform, however, it becomes feasible to define a single set of access-control policies, keep them in Git and apply them to all clusters or to specific engineering teams. This is true not only because you can apply the configurations in an automated way, but also because having a centralized, automated approach to policy management allows you to control the number of people who have access to your clusters in the first place.

And that, of course, is a huge security benefit as well as a manageability benefit.

## Conclusion

GitOps is key to rolling out highly scalable Kubernetes strategies that are reliable, efficient and secure. Manual approaches to configuration may work when you have only one or two clusters to manage; but there is no way to work with fleets of clusters without the help of GitOps policy management. You’ll waste far too much time, your configurations will be inconsistent and you will probably make mistakes that lead to security issues. GitOps addresses all of these challenges to enable a Kubernetes management strategy that is truly consistent and scalable.

## Learn more about the GitOps policy manager in Weave Kubernetes Platform

The [Weave Kubernetes Platform (WKP)](https://www.weave.works/product/enterprise-kubernetes-platform/) is a production ready platform with GitOps as the underlying architecture and developer experience. It simplifies cluster configuration and management across your organization by bringing together all the tools, services, and components that your team needs to run into a single platform. WKP also provides policy and Git-based rules to specify, audit, and control who can change what in the cluster configuration.

[Learn more about the features of Weave Kubernetes Platform in this on-demand webinar](https://vimeo.com/489580000)

* * *

## About Anita Buehrle

Anita has over 20 years experience in software development. She’s written technical guides for the X Windows server company, Hummingbird (now OpenText) and also at Algorithmics, Inc. She’s managed product delivery teams, and developed and marketed her own mobile apps. Currently, Anita leads content and other market-driven initiatives at Weaveworks.
