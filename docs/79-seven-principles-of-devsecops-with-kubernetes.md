# 7 Principles of DevSecOps With Kubernetes

March 15, 2021March 14, 2021

In my article, “ [9 Pillars of Engineering DevOps With Kubernetes](https://containerjournal.com/uncategorized/9-pillars-of-engineering-devops-with-kubernetes/),” I explain that continuous security is a core pillar of every well-engineered DevOps.

As indicated in the white paper, “ [From the Node Up: The Complete Guide to Kubernetes Security With Prisma Cloud](http://paloaltonetworks.com/prisma/cloud),” Kubernetes is a multi-layered, complex platform that consists of more than half a dozen different components that present both challenges and opportunities for DevSecOps.

Complex applications operating over complex distributed infrastructure can be difficult to secure. Cloud-native tools, such as Kubernetes, provide more insight into what is happening within an application, making it easier to identify and fix security problems. The enhanced orchestration controls, provided by Kubernetes on the deployment and deployed containerized applications, benefit from immutable consistency and improved response times. In addition, the secrets objects offer a secure way to store sensitive data.

In this article, I indicate how Kubernetes can be used and configured to satisfy seven principles for a successful DevSecOps approach using Kubernetes.  The seven DevSecOps principles are those identified in the [Department of Defense Enterprise DevSecOps Reference Design](https://dodcio.defense.gov/Portals/0/Documents/DoD%20Enterprise%20DevSecOps%20Reference%20Design%20v1.0_Public%20Release.pdf).

**Principle #1: Remove bottlenecks (including human ones) and manual actions.**

With Kubernetes, developers and testers can work better, together. They can solve defects quickly and accurately because developers can use the tester’s Kubernetes instance for debugging. This eliminates long delays associated with replicating development and test environments. Kubernetes also helps testers and developers quickly exchange precise information for application configurations.

**Principle #2: Automate as much of the development and deployment activity as possible.**

Kubernetes eliminates many of the manual provisioning and other time-consuming tasks of enterprise IT operations. In addition, the unified and automated orchestration approaches Kubernetes offers simplify multi-cloud management, enabling more services to be delivered with less work and fewer errors.

**Principle #3: Adopt common tools from planning and requirements through deployment and operations.**

Kubernetes offers many capabilities that allow one container to support many configuration environment contexts. The **configmaps** object, for example, supports configuration data that is used at runtime. This avoids the need for specialized containers for different environment configurations. Declarative syntax used to define the deployment state of Kubernetes-deployed container clusters greatly simplifies the management of the delivery and deployments.

**Principle #4: Leverage agile software principles and favor small, incremental, frequent updates over larger, more sporadic releases.**

Modular applications architected as microservices benefit the most from Kubernetes.  Software designed in accordance with twelve-factor app tenets and communicating through networked APIs work best for scalable deployments on clusters. Kubernetes is optimal for orchestrating cloud-native applications. Modular distributed services are better able to scale and recover from failures.

**Principle #5: Apply the cross-functional skill sets of development, cybersecurity and operations throughout the software life cycle, embracing a continuous monitoring approach in parallel instead of waiting to apply each skill set sequentially.**

Kubernetes provides a unified approach for container orchestration that applies end-to-end across the value stream. Continuous monitoring is facilitated because cloud-native applications, managed by Kubernetes, are constructed with health reporting metrics to enable the platform to manage life cycle events if an instance becomes unhealthy. They produce (and make available for export) robust telemetry data to alert operators to problems and allow them to make informed decisions. Kubernetes supports liveness and readiness probes that make it easy to determine the state of a containerized application.

**Principle #6: Security risks of the underlying infrastructure must be measured and quantified, so that the total risks and impacts to software applications are understood.**

Kubernetes has many different layers and components that must be considered for security. Elements of key concern for security are: API communications between different parts of a cluster, a scheduler that manages how workloads are distributed, controllers that manage the state of Kubernetes itself, agents that runs on each node within a cluster and a key-value store where cluster configuration data is housed. A multi-pronged defense strategy is needed to protect against all types of vulnerabilities.   The following is partial list of defense strategies.

- Secure container images to run on Kubernetes. Use security code scanning tools to check the containerized code for vulnerabilities that can exist within the container code itself, as well as in any upstream dependencies on which the image is based.
- Isolate Kubernetes nodes on a separate network that is not be exposed directly to public networks.
- Kubernetes supports role-based access control (RBAC) policies to help guard against unauthorized access to cluster resources.
- Resource quotas help mitigate disruptions caused by denial-of-service attacks by depriving the rest of the cluster of sufficient resources to run.
- Restrict pod-to-pod traffic using Kubernetes core data types for specifying network access controls between pods.
- Implement network border controls to enforce some ingress and egress controls at the network border in addition to the pod-level controls enforced by Kubernetes.
- Application-layer access control can be hardened with strong application-layer authentication, such as mutual transport-level security protocols using encrypted application identity.
- Kubernetes support for multiple containers running together with a shared localhost network for the pod enables sidecars and a service mesh approach to retrofit existing applications. This reduces the difficulties of implementing mutual TLS solutions, so each application has an adjacent proxy daemon that terminates and authenticates inbound connections and transparently authenticates outbound connections.
- Segment your Kubernetes clusters by integrity level; for example, your dev and test environments might be hosted in a different cluster than your production environment.
- Run your applications as a non-root user. Future Linux kernel vulnerabilities are more likely to be exploitable by a root user than by a non-privileged user.
- Use security monitoring and auditing to capture application logs, host-level logs, Kubernetes API audit logs and cloud provider logs. For security audit purposes, consider streaming your logs to an external location with append-only access from within your cluster.
- Use process whitelisting to identify unexpected running processes.
- Keep Kubernetes versions up to date.

A comprehensive security strategy for Kubernetes need to include more than the handful of built-in security features.

**Principle #7: Deploy immutable infrastructure, such as containers.**

The concept of immutable infrastructure supported by Kubernetes, in which deployed components are replaced in their entirety, rather than being updated in place, requires standardization and emulation of common infrastructure components to achieve consistent and predictable results.

In this article, I explained how Kubernetes can be used and configured to satisfy DoD’s seven DevSecOps principles for a successful DevSecOps approach using Kubernetes. While Kubernetes provides built-in security tools, they are not sufficient to fully protect against multiple types of potential vulnerabilities across multiple layers of Kubernetes infrastructure. All seven DevSecOps principles are important for an integrated security strategy that mitigates threats at all layers and levels of your stack.