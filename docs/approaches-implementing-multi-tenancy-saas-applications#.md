# Approaches to implementing multi-tenancy in SaaS applications

May 19, 2022 From: https://developers.redhat.com/articles/2022/05/09/approaches-implementing-multi-tenancy-saas-applications#


This article discusses architectural approaches for separating and isolating SaaS tenants to provide [_multi-tenancy_](https://www.redhat.com/en/topics/cloud-computing/what-is-multitenancy), the provisioning of services to multiple clients in different organizations. For the approaches, the type and level of isolation provided are compared, along with their tradeoffs.

The approaches laid out in different sections of the article are not mutually exclusive and can be combined to provide the levels of separation and isolation necessary to satisfy the requirements of your SaaS customers and markets. We'll also discuss how to incorporate existing single-tenant applications into a SaaS environment.

## Multi-tenancy considerations

The resources that are shared across multiple tenants vary, based on the architecture of the SaaS application. The level of sharing can be very high in large business-to-consumer (B2C) SaaS applications. In these cases, a given application instance potentially handles requests from thousands of unrelated tenants. For a more sensitive business-to-business (B2B) application, each tenant might get a dedicated application instance, though it could still be running on shared infrastructure.

Perhaps the most important aspect of a SaaS service is making sure tenants feel they have privacy and that no other tenant can see their data or activities. Additionally, tenant data must be stored securely and reliably. The integrity of each tenant's data must be protected from accidental loss, modification, or tampering.

When architecting a SaaS application, consider the level of isolation and controls required for the service you plan to offer. The requirements vary depending on the nature of the application, the sensitivity of the data, and the type of market. The contract or terms of service for the application usually detail what the provider is promising. However, regulatory or government agencies might issue compliance requirements for specific industries, such as financial services or health care, that you must adhere to when providing a SaaS application to customers under their jurisdiction.

It is worth considering that a one-size-fits-all approach might not be enough to balance the security requirements and costs for all prospective customers. Customers from highly regulated industries or groups might not be the target for your SaaS application. However, if you do attract customers from financial services, health care, or government organizations, you might have to adhere to requirements you didn't originally anticipate.

Also, some customers are under data sovereignty regulations that specify where the physical machines storing data must reside. These requirements potentially dictate data center locations and even which cloud providers can be used to host the SaaS services the customers consume.

## Using application logic to provide multi-tenancy

When you think of SaaS applications, the ones that come to mind are probably those from large technology companies that are offered to millions of consumers. These applications are built from the ground up for multi-tenancy. They are designed for high-density deployments where large numbers of users can be served using minimal per-user resources. To achieve the necessary efficiency, a very high degree of sharing is usually necessary.

In an application design where each software instance handles requests from multiple tenants, the infrastructure generally can't provide the needed isolation. Instead, application logic is typically used to implement the controls that isolate tenants in application-level tenancy.

In this model, a single instance of the application serves requests from multiple tenants. The code is responsible for keeping each tenant's data separate and making sure other tenants can't see any data or activity that does not belong to them.

The controls in application-level tenancy must be implemented by software developers. Therefore, extensive [automated](http://developers.redhat.com/topics/automation) testing and strict adherence to secure coding best practices are critical for application tenancy, as any bugs in the application could accidentally disclose tenant data or compromise data integrity. If a vulnerability in the application is found and exploited, relying solely on application logic to implement isolation could make it difficult to contain and mitigate the potential impact to tenants that share the same instances.

Application tenancy is best for applications that are designed for multi-tenancy as they are developed. Retrofitting an existing application to support multiple tenants can require extensive software development work as well as extensive testing. The costs and time necessary to modify an existing application could be high. For those applications, it might be better to consider the next two approaches we'll discuss in this article, where infrastructure-level controls provide the necessary separation and isolation.

### Advantages of application-level tenancy

Implementing multi-tenancy in the code of your application itself is appealing for a number of reasons:

- No infrastructure support is necessary to implement application-level tenancy.
- Higher tenant density is possible, resulting in lower infrastructure costs because application instances can be shared by multiple tenants.
- In many cases, no additional processes or infrastructure need to be provisioned when adding tenants.

### Disadvantages of application-level tenancy

However, there are also reasons you might want to avoid this architecture:

- This approach entails high development costs, because isolation must be implemented by software developers.
- Extensive security testing is needed whenever changes are made to validate isolation is correctly enforced.
- This approach provides a single layer of defense, so bugs or vulnerabilities in the application could accidentally disclose or corrupt tenant data.
- Security reviews by outside auditors or regulators can be difficult and labor intensive, because they need to understand and validate the controls implemented by the application.
- Change management occurs at a very coarse granularity. Upgrades, downgrades, or patches affect the entire tenant population or large subsets of it, because application instances span multiple tenants.
- It is difficult and costly to retrofit existing applications that were not designed for multi-tenant SaaS deployments.

## Isolating tenants on the same cluster using namespaces

An alternative to the application-level tenancy described in the previous section is to use the capabilities available in [Kubernetes](http://developers.redhat.com/topics/kubernetes) to isolate applications that share a cluster. Using these controls, it is possible to run multiple instances of the same unmodified single-tenant application on the same cluster while keeping them logically separated. _Namespace-level tenancy_ can allow existing applications that were not written for SaaS to be deployed as a SaaS solution.

Namespaces within Kubernetes provide a mechanism for isolating resources into different groups that keep user communities or application groups separate even while they share a cluster. For multi-tenancy, namespaces provide isolation between tenant workloads. Because Kubernetes enforces separation of namespaces, an application running inside of a namespace can't accidentally or intentionally access data or processes in a different namespace.

Kubernetes namespaces provide ways to use a single-tenant application in a multi-tenant SaaS deployment. The most common and perhaps most straightforward approach is to create a namespace for each tenant, and deploy the complete application stack and data store into each tenant's namespace. The namespace provides isolation between each application instance. This approach is referred to as _namespace-level tenancy_, and is illustrated in Figure 1.

[![Two tenants are shown with namespace-level tenancy, storing each application's resources including data in its own namespace.](http://developers.redhat.com/sites/default/files/styles/article_full_width_1440px_w/public/image1_7.png?itok=FdzkVvnl)](http://developers.redhat.com/sites/default/files/image1_7.png)

Figure 1. Two tenants are shown with namespace-level tenancy, storing each application's resources including data in its own namespace.


Figure 1: Two tenants are shown with namespace-level tenancy, storing each application's resources including data in its own namespace.

To provide higher degrees of isolation and better control over resources, cluster administrators can control which worker nodes are available for each tenant's pods. Specific tenants can be configured to use only dedicated worker nodes, which could run on dedicated hardware if desired.

With namespace-level tenancy, a single control plane manages all the application instances and worker nodes used by all tenants of the cluster. The cluster's control plane contains configuration information for all tenant application instances running on that cluster. Anyone with cluster administration access can view and change configuration information for all tenants of the cluster.

### Advantages of namespace-level tenancy

This architecture is appealing for a number of reasons:

- Single-tenant applications can be used in SaaS deployments by using namespace controls to run isolated application instances for each tenant. The application does not need to be modified for SaaS use, because the cluster provides multi-tenancy control.
- The impacts of bugs and vulnerabilities within the application or application components are contained to a single tenant running in that namespace.
- Individual tenants can have customized deployments, making it easier to add features to a select set or add third-party integrations.
- Resource consumption is easy to view at the level of an individual tenant.
- The resources available for scaling the application up or down can be optimized at the level of a single tenant.
- Application changes can be managed flexibly at a granular level. The cadence of new releases and patches can be optimized on a per-tenant basis.

### Disadvantages of namespace-level tenancy

However, there are reasons why you might decide not to use this architecture:

- Deployment density is lower and infrastructure costs are higher than in application-level tenancy, because the cluster has to run multiple application instances that are not shared. In the simplest case, a complete application stack is run for each tenant.
- Tenant onboarding requires the cluster to create a new project or namespace along with the necessary services deployed with it. However, this process can be automated.
- Cluster-wide scaling and resource management can be more challenging, because changes might be necessary for each namespace or project.
- Change management at the cluster level still has potential impacts on all tenants.

### Utilizing namespaces with Red Hat

[Red Hat OpenShift](http://developers.redhat.com/openshift) has a number of features that improve security and isolation beyond what is available in upstream Kubernetes. These controls provide additional protection for namespace-level tenancy:

- [Security-enhanced Linux (SELinux)](https://www.redhat.com/en/topics/linux/what-is-selinux) provides mandatory and discretionary access controls for files and processes at the operating system level, going beyond the default mechanisms in Linux.
- [Security context constraints](https://docs.openshift.com/container-platform/4.10/authentication/managing-security-context-constraints.html) limit the capabilities of individual containers.
- [Network-level isolation](https://docs.openshift.com/container-platform/4.10/networking/network_policy/multitenant-network-policy.html) prevents containers in one namespace from connecting to containers in other namespaces.

A future article will go into more detail on the security considerations for SaaS applications. For more information, please see:

- [Red Hat OpenShift documentation: Projects and users](https://docs.openshift.com/online/pro/architecture/core_concepts/projects_and_users.html)
- [Red Hat OpenShift security guide](https://www.redhat.com/en/resources/openshift-security-guide-ebook)
- [Kubernetes documentation: Namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)

## Using single-tenant clusters for strong isolation

To provide a very high level of isolation, a tenant can be assigned to a dedicated Kubernetes or Red Hat OpenShift cluster. In this case, no resources are shared between tenants—not even the cluster's control plane.

There are a number of cases where single-tenant clusters can be a good fit:

- A dedicated cluster can be a good option for a tenant that is very risk-averse and is willing to pay for a higher level of service.
- Single-tenant clusters can be used to address compliance requirements or legal regulations, particularly for highly regulated industries.
- The hardware the cluster runs on might need to be physically located within a specific national entity to comply with data sovereignty laws.
- Using single-tenant clusters allows cluster administrative access to be partitioned so only specific administrators have control over a given tenant's cluster.

Figure 2 shows an overview of cluster-level tenancy.

[![Two separate clusters, each containing not only its own application resources but infrastructure such as control planes.](http://developers.redhat.com/sites/default/files/styles/article_full_width_1440px_w/public/image2_3.png?itok=aoV2MSTx)](http://developers.redhat.com/sites/default/files/image2_3.png)

Figure 2. Two separate clusters, each containing not only its own application resources but infrastructure such as control planes.


Figure 2: Two separate clusters, each containing not only its own application resources but infrastructure elements such as control planes.

Single-tenant clusters provide a very high level of protection to mitigate the impact of security breaches. As discussed earlier, namespace-level tenancy limits any vulnerability within the application or application's components to the tenant using that namespace. However, parts of the technology stack are still shared when multiple tenants use the same worker node and control plane. A vulnerability in these shared parts of the stack could disclose or compromise the integrity of other tenant data. The shared components include:

- **Container engine:** The container engine is responsible for configuring the isolation between containers and the host system. A bug in the container engine could allow a process running inside a container to break out.
- **Linux kernel and host operating system:** A vulnerability in the Linux kernel or operating system components could allow an intruder to bypass the system-level controls that isolate users, processes, or containers.
- **Hardware such as CPUs and memory:** A hardware defect could be exploited to access the memory of other processes running on the system. The [Meltdown and Spectre attacks](https://meltdownattack.com) against weaknesses in CPU design have shown that this risk is not theoretical and could be exploited to gain access to sensitive information such as encryption keys.
- **Control plane:** A bug in the implementation of the control plane could allow an intruder to bypass cluster isolation controls.

The good news that is vulnerabilities in these lower levels of the technology stack occur far less frequently than application-level vulnerabilities. The best practice to guard against security vulnerabilities is to use a defense-in-depth strategy, where an attack would have to breach multiple layers of protection.

Single-tenant clusters are a very strong line of defense for situations that require this level of guaranteed security. A future article in this SaaS architecture checklist series will cover a number of the tools that can be used as part of a defense-in-depth security framework.

### Advantages of single-tenant clusters

This architecture is appealing for the following reasons:

- This design offers the highest level of isolation and security among the three approaches outlined in this article. Deployments do not share components of the technology stack, keeping even the container engine, Linux kernel, and control plane separate.
- This design also offers the highest degree of per-tenant customization. An application can run on completely separate hardware, in different clouds or data centers, and even in different countries.
- Change management cadence and scheduling for cluster maintenance and upgrades can be tailored to each tenant's needs.

### Disadvantages of single-tenant clusters

You might decide not to use this architecture for the following reasons:

- Because there is no sharing between tenants, this design has the lowest tenant density and therefore the highest infrastructure costs among the three approaches.
- This design also involves the most cluster administration, because a cluster needs to be deployed and maintained for each tenant. However, automation and tools can significantly reduce the effort required.

### Utilizing Kubernetes clusters with Red Hat

There are a number of options to make the deployment and management of multiple Kubernetes or Red Hat OpenShift clusters faster and easier.

[Red Hat Advanced Cluster Management for Kubernetes](https://www.redhat.com/en/technologies/management/advanced-cluster-management) allows you to manage multiple Kubernetes or Red Hat OpenShift clusters from a single management interface. You can use Red Hat Advanced Cluster Management for Kubernetes's unified multi-cluster lifecycle management to create, update, upgrade, and decommission clusters as needed for your SaaS deployments. Some of the managed services for Red Hat OpenShift on public clouds can manage multiple clusters.

A feature planned for a future release of Red Hat Advanced Cluster Management for Kubernetes is _hosted control planes_. This feature allows the control plane for new clusters to be run on the worker nodes of an existing management cluster, which reduces the infrastructure and deployment time required. The upstream project for hosting OpenShift clusters at scale is called [HyperShift](https://cloud.redhat.com/blog/a-guide-to-red-hat-hypershift-on-bare-metal). A future article will discuss multi-cluster management.

## Comparison of multi-tenancy approaches

This article discussed three approaches to providing multi-tenancy for SaaS applications:

- **Application-level tenancy:** Using application logic to isolate tenants.
- **Namespace-level tenancy:** Providing isolation between tenant application instances through Kubernetes namespaces or OpenShift projects.
- **Cluster-level tenancy:** Using a dedicated Kubernetes cluster for each tenant to provide a very high level of isolation and security.

The first approach requires the application to be written for multi-tenancy. The latter two approaches can be used to deploy existing single-tenant applications as SaaS applications. Table 1 compares the three approaches.

Table 1: Comparison of multi-tenancy approachesApplication-level

tenancyNamespace-level tenancyCluster-level tenancyTenant isolation enforced byApplication logicSeparate namespacesSeparate clustersLevel of infrastructure shared across tenantsHigh-allMedium-mixedLow-noneTenant density possibleHighMediumLowInfrastructure costs per tenantLowMediumHighSecurity protectionsLowMediumHighAbility to monitor and scale per tenantLowMediumHighAbility to customize deployment for each tenantLowMediumHigh

These three approaches are not mutually exclusive and can be combined to provide the best fit for your application, or to address the requirements of specific tenants. For example, an application that implements multi-tenancy could be deployed as multiple instances with either namespace-level or cluster-level tenancy to provide enhanced security protections and to contain the potential impact of vulnerabilities to a limited number of tenants.

A SaaS provider could offer multiple levels of service to address the requirements of tenants that are more risk-averse and are willing to pay for higher levels of service, versus those that are more cost-conscious and less sensitive to risk. Cost-conscious tenants can be served by environments with high degrees of sharing. The most risk-sensitive tenants could be offered the option of dedicated hardware to meet their requirements.

Future articles in this series will cover topics such as security and the options available for a defense-in-depth strategy. Another article will address single-tenant workloads that are difficult to containerize and will explain how these can be deployed for SaaS use as virtual machines using Red Hat OpenShift Virtualization with namespace-level isolation.

## Partner with Red Hat to build your SaaS

[Red Hat SaaS Foundations](https://connect.redhat.com/en/partner-with-us/red-hat-saas-foundations) is a partner program designed for building enterprise-grade SaaS solutions on Red Hat OpenShift or Red Hat Enterprise Linux platforms, and deploying them across multiple cloud and non-cloud footprints. [Email](http://mailto:saas@redhat.com) us to learn more.

_Last updated:
June 2, 2022_
