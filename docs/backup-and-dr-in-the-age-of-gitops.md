# Backup and DR in the Age of GitOps

August 10, 2020

Whether you call it GitOps, infrastructure as code (IaC) or simply using a CI/CD pipeline to automatically deploy changes to an application, we can all agree that moving to a model where your application and resource configuration is defined as code is beneficial to everyone involved—particularly the business it supports. This code is usually saved in git, the most popular source version control system for cloud-native projects; provides an automatic and verifiable change capture process; and simplifies application rollout. More importantly, it also prevents “snowflakes,” wherein the configuration of a deployed application differs from the definition because of manual edits.

But even with these capabilities available to developers and operators, the chances of catastrophic business disruption remain. System safety and scalability should still be an area of concern for developers and operators even if GitOps is on point. As such, the role of backup and disaster recovery capabilities, once the domain of legacy data center operators, needs to be elevated in GitOps’ operating models.

## It’s All About the Data

Any automated way of deploying an application will only bring back Kubernetes objects and configuration. Hence, any persistent data or volumes used by applications are not captured in version control; therefore, bringing back any stateful service such as a relational database or NoSQL system requires that the entire application stack, its data and the dependencies of the stack on the data be discovered, tracked and captured.

With any application redeployment coming in from an automated system, the backup platform must support intelligent “data-only restores.” This is one of the areas where Kubernetes-native data capabilities are critical; the ability to track the relationship between stateful services (including replication) and then bring back only the volumes or data associated with the application after the application has been redeployed via a CI/CD or GitOps pipeline can be massively helpful for developers.

## CI/CD at Scale

Enterprises that run Kubernetes at scale and as shared infrastructure have clusters with hundreds of applications from multiple teams and business units. As different groups make independent technology choices, different CI/CD systems and deployment mechanisms in large companies range from GitOps and CI/CD-based Helm installs to manual deployments. This is the reality in complex environments.

Faced with this reality, a backup platform is necessary for teams that are still [ramping up their cloud-native development](https://containerjournal.com/topics/container-ecosystems/the-keys-to-successful-infrastructure-configuration/). However, in the fast-moving Kubernetes ecosystem and the 50-plus ways to deploy a containerized application, some applications tend to be long-lived. The systems they were deployed with aren’t being maintained anymore and might not be compatible with the installed Kubernetes version (e.g., applications installed with Helm v2, that will be EOL soon). Being able to restore these important applications (they are long-lived for a reason) quickly in case of accidental or malicious failure will be critical.

Additionally, when multiple CI/CD systems are in use, it is difficult for an operations group to reach out to hundreds of developer teams to retrigger deployment in case of failure or when applications need to be moved across clusters to perform a Kubernetes version upgrade.

Sopra Steria, a European information technology consultancy, recently had this problem when the company had to move 170+ applications to [upgrade OpenShift versions](https://blog.kasten.io/kasten-and-red-hat-migration-and-backup-for-openshift) and couldn’t get responses from some teams that were running applications on their cluster. Integrating its backup platform into DevOps workflows enabled Sopra Steria to capture both stateful and stateless applications and move them over without downtime and allowed its developer teams to resync their pipelines to the new clusters when they found free time.

## Disaster Recovery

As teams go through Kubernetes business continuity planning, they need to ensure that Kubernetes and all its applications can be restored quickly in case of disaster. Disaster recovery (DR) can take many forms including cross-cluster, cross-region/data center and cross-cloud.

While pushing configuration to a version control system and using CI/CD will be beneficial here too, the DR problem doesn’t go away. The version control system must be high availability across all fault domains to be able to recover quickly. DR capabilities should be able to support both full backups and granular restores. So, if the version control system (e.g. git) has been recovered in a separate fault domain, the CI/CD system should be able to deploy the application configuration from the checked-in state while selectively bringing only the data back from its DR copies. If the version control system is not available, the application configuration and all related Kubernetes should be restored from the backup along with the data. However, because of its application-centricity and Kubernetes-native knowledge, this state will be quickly reconciled with your GitOps or CI/CD system when the source is available again.

There are several other reasons why a Kubernetes-native backup platform that understands runtime application state is most appropriate for integration with a deployment pipeline. Some application deployments have external side effects (e.g., a DNS update or external load balancer creation) that are dynamically named and are non-deterministic. These changes should be captured by a system that understands runtime state, and this should be brought back in on restore, along with the static resources from the automated deployment system.

Finally, if clusters are in a regulated environment (e.g., U.S. financial services regulated by the SEC) and your backups need to prove that they captured what was really running versus what the desired state of the world should have been, a platform that can capture the runtime state will be critical. Even if the sync period between configuration changes and deployment can be small on average, the window can drift in case of controller or CI/CD failures, manual additions of untracked resources or malicious changes.

Interoperability with GitOps, IaC, and CI/CD systems is a must as organizations increasingly are deploying these systems to improve IT operations and enhance business success. Even with these powerful primitives available, the need for backup and disaster recovery is as important as ever, and deploying a backup solution will be critical for safety and scale. However, it is imperative that such a backup system be truly cloud-native that can integrate into GitOps and CI/CD workflows. Legacy VM-based systems will simply not work in the new cloud-native world we live in today!

- [Click to share on Twitter (Opens in new window)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=twitter "Click to share on Twitter")
- [Click to share on Facebook (Opens in new window)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=facebook "Click to share on Facebook")
- [Click to share on LinkedIn (Opens in new window)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=linkedin "Click to share on LinkedIn")
- [Click to share on Reddit (Opens in new window)](https://containerjournal.com/topics/container-security/backup-and-dr-in-the-age-of-gitops/?share=reddit "Click to share on Reddit")

### _Related_

- [← New Training Course Prepares Developers to Create Enterprise Blockchain Applications](https://containerjournal.com/news/news-releases/new-training-course-prepares-developers-to-create-enterprise-blockchain-applications/)
- [Survey: Containers Driving Microservices Transition →](https://containerjournal.com/topics/container-ecosystems/survey-containers-driving-microservices-transition/)

![](https://containerjournal.com/wp-content/uploads/2020/08/kasten_niraj_HR-150x150.jpg)

#### Niraj Tolia

Niraj Tolia is the CEO and Co-Founder at Kasten and is interested in all things Kubernetes. He has played multiple roles in the past, including the Senior Director of Engineering for Dell EMC's CloudBoost family of products and the VP of Engineering and Chief Architect at Maginatics (acquired by EMC). Niraj received his PhD, MS, and BS in Computer Engineering from Carnegie Mellon University.

Niraj Tolia has 1 posts and counting. [See all posts by Niraj Tolia](https://containerjournal.com/author/niraj-tolia/)
