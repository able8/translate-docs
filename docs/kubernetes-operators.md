[Skip to main content](http://www.redhat.com#main-content)

We use cookies on our websites to deliver our online services. Details about how we use cookies and how you may disable them are set out in our [Privacy Statement](http://www.redhat.com/en/about/privacy-policy#cookies). By using this website you agree to our use of cookies.
[×](javascript:void(0);)

[![Red Hat logo](http://www.redhat.com/sysadmin/themes/custom/sysadmin/logo.svg)](https://www.redhat.com/)

Search

Search

[Enable Sysadmin](http://www.redhat.com/sysadmin/)

- [Articles](http://www.redhat.com/sysadmin/)
[Automation](http://www.redhat.com/sysadmin/topics/automation) [Career](http://www.redhat.com/sysadmin/topics/career) [Containers](http://www.redhat.com/sysadmin/topics/containers) [Culture](http://www.redhat.com/sysadmin/topics/sysadmin-culture) [Kubernetes](http://www.redhat.com/sysadmin/topics/kubernetes) [Linux](http://www.redhat.com/sysadmin/topics/linux) [Programming](http://www.redhat.com/sysadmin/topics/programming) [Security](http://www.redhat.com/sysadmin/topics/security)

- [About](http://www.redhat.com/sysadmin/about)
[Topics](http://www.redhat.com/sysadmin/about) [Email newsletter](http://www.redhat.com/sysadmin/email) [Join the community](http://www.redhat.com/sysadmin/join-community) [Community guidelines](http://www.redhat.com/sysadmin/community-guidelines) [Sudoers program](http://www.redhat.com/sysadmin/sudoers-program) [Meet the team](http://www.redhat.com/sysadmin/meet-team) [FAQs](http://www.redhat.com/sysadmin/faq)

- [Welcome](http://www.redhat.com/sysadmin/welcome-enable-sysadmin "A Red Hat community publication for sysadmins, by sysadmins")

[Enable Sysadmin](http://www.redhat.com/sysadmin/)

- [Articles](http://www.redhat.com/sysadmin/)
[Automation](http://www.redhat.com/sysadmin/topics/automation) [Career](http://www.redhat.com/sysadmin/topics/career) [Containers](http://www.redhat.com/sysadmin/topics/containers) [Culture](http://www.redhat.com/sysadmin/topics/sysadmin-culture) [Kubernetes](http://www.redhat.com/sysadmin/topics/kubernetes) [Linux](http://www.redhat.com/sysadmin/topics/linux) [Programming](http://www.redhat.com/sysadmin/topics/programming) [Security](http://www.redhat.com/sysadmin/topics/security)

- [About](http://www.redhat.com/sysadmin/about)
[Topics](http://www.redhat.com/sysadmin/about) [Email newsletter](http://www.redhat.com/sysadmin/email) [Join the community](http://www.redhat.com/sysadmin/join-community) [Community guidelines](http://www.redhat.com/sysadmin/community-guidelines) [Sudoers program](http://www.redhat.com/sysadmin/sudoers-program) [Meet the team](http://www.redhat.com/sysadmin/meet-team) [FAQs](http://www.redhat.com/sysadmin/faq)

- [Welcome](http://www.redhat.com/sysadmin/welcome-enable-sysadmin "A Red Hat community publication for sysadmins, by sysadmins")

[Subscribe to our RSS feed or Email newsletterSubscribe to our RSS feed or Email newsletter.](http://www.redhat.com/sysadmin/subscribe) [Search this site…](http://www.redhat.com/sysadmin/search/node)

GO

# Kubernetes operators: Embedding operational expertise side by side with containerized applications

Kubernetes isn't complex, your business problem is. Learn how operators make it easy to run complex software at scale.

Posted:
March 2, 2020
\|**by** [Scott McCarty (Red Hat, Sudoer)](http://www.redhat.com/sysadmin/users/scott-mccarty)

Image

![Sidecars](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201000%20600'%2F%3E)

**Why use operators?**

At a high level, the [Kubernetes Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) makes it easy to run complex software at scale. With the DevOps movement, we learned to manage and monitor complex applications and infrastructure from centralized platforms (Chef Server, Puppet Enterprise, Ansible Tower, Nagios, etc). This centralized monitoring & automated remediation works great with relatively stable infrastructure components like bare metal servers, virtual machines, and network devices. However, containers change much quicker than traditional infrastructure and traditional applications. One might say this speed is a tenet of Cloud Native behavior.

The ephemeral nature of containerized applications drives an immense amount of state change. In turn, this rapid state change challenges the ability of centralized monitoring and automation to keep up. Kubernetes can literally change the state of containers every few seconds. The solution is to bring the automation close to the applications - to deploy it within Kubernetes so that it has direct access to the application’s state at all times. The Kubernetes Operator pattern allows us to do just this - deploy automation side by side with the containerized application.

I often describe the Operator pattern as deploying a robot sysadmin next to the containerized application. Though, to truly understand how Operators work, we need to dive a bit deeper into the history of running services, and the art of making these services resilient. From here on, we will refer to this as operational knowledge, or operational excellence.

Image

![Kubernetes flow chart](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201232%201579'%2F%3E)

Traditionally, when new software was deployed, we also deployed a real, human Sysadmin to care and feed the application. This care and feeding included tasks like installation, upgrades, backups, & restores, troubleshooting, and return to service. If a service failed, we paged this Sysadmin, they logged into a server, would troubleshoot the application, and fix what was broken. To track this work, they would document their progress in a ticketing system.

As we moved into the world of containers (about six years ago at the time of this writing), we updated the packaging format of the application, but we continued to deploy a real human Sysadmin to care and feed the application. This simplified installation, and upgrades, but did little for data backups/restores and break/fix work.

With the advent of the Kubernetes Operator pattern, we deploy the software application in a container and we also deploy a robot Sysadmin in the same environment to care and feed the application. We refer to these robot Sysadmins as Operators. Not only can they perform installation, upgrades, backups, and restores, but they can also perform more complex tasks like recovering a database when the tables become corrupted. This takes us into another world of automation and reliability.

But, why does this work philosophically? You might still struggle to understand how this is different than traditional clustering, or having a monitoring system participate in “self-healing” (an old buzzword now). Let’s explain...

**A brief history of operational knowledge**

Operators embed operational knowledge, or operational excellence in the infrastructure, close to the service. To better explain this, let’s walk through the history of operating servers & services.

Image

![kubernetes operator flow chart](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%202624%20869'%2F%3E)

- **Operational excellence 1.0** \- One computer, multiple administrators. Back in years past, before most users of containers remember, there were large computers cared for, managed and operated by multiple systems administrators. This remained true all the way into the late 1990s with large Mainframes and Unix systems managed by multiple systems administrators. That’s right, there were multiple human beings assigned to one computer. These systems administrators automated operations on these systems, fixed services if they were broken, added/removed users, and even wrote device drivers themselves. The administrators were highly technical, and would be considered software engineers by modern definitions. They performed all software related tasks, maintaining uptime, reliability, and operational excellence. The cost was high, but the quality was also very high.

- **Operational Excellence 2.0** \- One administrator, multiple computers. With the advent of Linux, and perhaps more so Windows, the number of servers outgrew the number of administrators. There became less and less large, multi-user systems. At the same time there became more and more single service systems - DNS Servers, Web Servers, Mail Servers, etc. A single administrator could manage multiple servers remotely, and we often measured productivity by comparing the number of servers per administrator. Administrators still retained intimate knowledge of each service they managed, sometimes using [Runbooks](https://en.wikipedia.org/wiki/Runbook) to document common tasks. If a service or server failed, an administrator would work on it remotely thereby maintaining a high level of quality, while at the same time being responsible for a higher total number of servers.

- **Operational Excellence 3.0** \- One service, multiple computers. As a sort of operational dead end, there was a renewed focus on the quality of the service by leveraging clustering and automatic recovery on cheaper hardware. Databases, web servers, NFS, DNS SAMBA servers and more were clustered for resilience. Tools like Veritas Clustering, Veritas File System, Pacemaker, GFS, and GPFS became popular. For this to work properly, operational knowledge of how to start and recover a service had to be configured in the clustering software. This required a solid understanding of how the service worked (detection of failures, how to restart, etc). With clustering software and configuration, nodes could be treated as capacity (N+1 means one extra server, N+2 means two extra servers). Bringing operational knowledge close to the service allowed for automated recovery, but building new services or clusters, much less decommissioning them, could take days or even weeks because each service had to be designed and maintained separately.

- **Operational Excellence 4.0** \- With this iteration, we moved the logic for recovering services back away from the service itself, and put it in the monitoring systems and load balancers. We embedded configuration in DNS, and started to use configuration management to maintain things. This created a fundamental tension in a lot of IT organizations. Server administrators would embed the logic for recovering services in the monitoring system. For example, if the monitoring system saw a problem with a web server it could ssh into the server and restart Apache. There were several major challenges with this paradigm. First, having configuration in many different places created a lot of complexity (see also: [Why you don't have to be afraid of Kubernetes](https://opensource.com/article/19/10/kubernetes-complex-business-problem)). Second, storage, network, and virtualization administrators didn’t want automation logging in and provisioning/deprovisioning services, so achieving truly cloud-native architectures was difficult.

- **Operational Excellence 5.0** \- While centralized monitoring and automation can be abused with containers to achieve a [Desired State/Actual State](https://kubernetes.io/docs/concepts/#overview) model similar to Kubernetes, it’s not nearly as elegant. Using a manifest (YAML or JSON), Kubernetes enables the powerful desired state. For example, once an administrator defines that they want three copies of a web server running, Kubernetes will maintain exactly three copies running. If one of the containerized web servers dies, Kubernetes will restart another one because it sees that the actual state doesn’t match the desired state. This gives you application definition and recovery in one file format. But, how do you manage more complex tasks like scanning corrupted database tables, upgrades, schema changes, or rebalancing data between volumes? That’s what operators do. It moves this logic close to the service bringing the best of Operational Excellence 3.0 and 4.0 together in one place. This same logic can be applied to the Kubernetes platform itself (see also: [Red Hat OpenShift Container Platform 4 now defaults to CRI-O as underlying container engine](https://www.redhat.com/en/blog/red-hat-openshift-container-platform-4-now-defaults-cri-o-underlying-container-engine)).


**Who should build operators**

This leads us to the important question of who should be building operators? Well, the short answer is, it depends.

If you are a developer building Java web services, Python web services, or really any web services then you shouldn’t have to write an Operator. Your web services should be relatively stateless and as long as you focus on learning to use [readiness and liveness checks](https://opensource.com/article/19/10/kubernetes-complex-business-problem), Kubernetes will manage the desired state for you. That said, you might use pre-built Operators to manage a PostgreSQL, MongoDB, or MariaDB instance. Check out all of the services you can consume from [Operator Hub](https://operatorhub.io/). Welcome to the Kubernetes ecosystem.

## Great DevOps downloads

- [The ultimate DevOps hiring guide](https://opensource.com/downloads/devops-hiring-guide?intcmp=701f20000012ngPAAQ)
- [DevOps monitoring tools guide](https://opensource.com/downloads/devops-monitoring-guide?intcmp=701f20000012ngPAAQ)
- [Getting started with DevSecOps](https://opensource.com/downloads/devsecops?intcmp=701f20000012ngPAAQ)

If you are an administrator, practicing DevOps, and you are building complex services for others to consume, then you may very well need to think about writing operators. You will almost certainly be consuming Operators and upgrading them, so check out the [Operator Lifecycle Manager](https://github.com/operator-framework/operator-lifecycle-manager) (Built into OpenShift 3.X and 4.X) and [Operator SDK.](https://github.com/operator-framework/operator-sdk)

If you’re a developer working on a data management, networking, monitoring, storage, security, or GPU solution for Kubernetes/OpenShift then you are the most likely to need to write an operator. I would suggest looking at the [Operator Framework](https://github.com/operator-framework/getting-started), [Operator SDK](https://github.com/operator-framework/operator-sdk), [Operator Lifecycle Manager](https://github.com/operator-framework/operator-lifecycle-manager), and [Operator Metering](https://github.com/operator-framework/operator-metering). You might also like to look at the Red Hat container certification program called [Partner Connect](https://connect.redhat.com/).

**Conclusion**

As I have said before, Kubernetes isn’t complex, your business problem is. Kubernetes not only redefines the operating system ( [OpenShift is the New Enterprise Linux](https://www.linkedin.com/pulse/openshift-new-enterprise-linux-daniel-riek/) and [Sorry, Linux. Kubernetes is now the OS that matters](https://www.infoworld.com/article/3322120/sorry-linux-kubernetes-is-now-the-os-that-matters.html)), it redefines the entire operational paradigm. Surely, the history of operational excellence could be divided into any number of paradigm shifts, but I have attempted to break it down with a model that is cognitively digestible to help operations teams, software vendors (ISVs), and even developers better understand why Operators are important.

This fifth generation of operational excellence, using the Kubernetes Operator pattern brings the automation close to the application giving it access to state in near real-time. With Operators deployed side by side within Kubernetes, response, management and recovery of applications can happen at a speed that just isn’t possible with human beings and ticket systems. An added benefit is the ability to provision multiple copies of an application with a single command, or more importantly deprovision them. This ability to provision, deprovision, recover, and upgrade is the fundamental difference between cloud-native and traditional applications.

**Credits**

I want to give special thanks to Daniel Riek who presented this concept of Operational Excellence at FOSDEM 20 in Brussels, Belgium last week. If you didn’t have an opportunity to attend his talk, I recommend you watch it when the video goes live. Until then, see this interview with him: [How Containers and Kubernetes re-defined the GNU/Linux Operating System. A Greybeard's Worst Nightmare](https://fosdem.org/2020/interviews/daniel-riek/)

**Topics:** [Kubernetes](http://www.redhat.com/sysadmin/topics/kubernetes)

![Author’s photo](http://www.redhat.com/sysadmin/sites/default/files/styles/user_picture_square/public/pictures/2020-02/smccarty.jpg?itok=doLLWwmx)

## Scott McCarty

At Red Hat, Scott McCarty is a technical product manager for the container subsystem team, which enables key product capabilities in OpenShift Container Platform and Red Hat Enterprise Linux. Focus areas include container runtimes, tools, and images.
[More about me](http://www.redhat.com/sysadmin/users/scott-mccarty)

#### On Demand: Red Hat Summit 2021 Virtual Experience

Relive our April event with demos, keynotes, and technical sessions from

experts, all available on demand.

[Watch Now](https://www.redhat.com/en/summit?intcmp=7013a0000026RhqAAE)

## Related Content

Image

![White ship wheel overlooking water and container ships](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201000%20600'%2F%3E)

[How to use Podman inside of Kubernetes](http://www.redhat.com/sysadmin/podman-inside-kubernetes)

More information about Podman in containers; specifically with regard to Kubernetes.

Posted:
July 1, 2021


Authors: [Urvashi Mohnani (Red Hat)](http://www.redhat.com/sysadmin/users/umohnani), [Dan Walsh (Red Hat)](http://www.redhat.com/sysadmin/users/dwalsh)

Image

![A brief introduction to CNI for Kubernetes](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%203781%202268'%2F%3E)

[A brief overview of the Container Network Interface (CNI) in Kubernetes](http://www.redhat.com/sysadmin/cni-kubernetes)

Understand where the CNI fits into the Kubernetes architecture.

Posted:
April 29, 2021


Author: [Kedar Vijay Kulkarni (Red Hat, Sudoer)](http://www.redhat.com/sysadmin/users/kedar-vijay-kulkarni)

Image

![From Docker Compose to Kubernetes with Podman](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%203000%201800'%2F%3E)

[From Docker Compose to Kubernetes with Podman](http://www.redhat.com/sysadmin/compose-kubernetes-podman)

Use Podman 3.0 to convert Docker Compose YAML to a format Podman recognizes.

Posted:
January 26, 2021


Authors: [Brent Baude (Red Hat)](http://www.redhat.com/sysadmin/users/brent-baude), [Urvashi Mohnani (Red Hat)](http://www.redhat.com/sysadmin/users/umohnani)

OUR BEST CONTENT, DELIVERED TO YOUR INBOX

The opinions expressed on this website are those of each author, not of the author's employer or of Red Hat. The content published on this site are community contributions and are for informational purpose only AND ARE NOT, AND ARE NOT INTENDED TO BE, RED HAT DOCUMENTATION, SUPPORT, OR ADVICE.

Red Hat and the Red Hat logo are trademarks of Red Hat, Inc., registered in the United States and other countries.

[![Red Hat logo](http://www.redhat.com/sysadmin/themes/custom/sysadmin/logo.svg)](https://www.redhat.com/)

Copyright ©2021 Red Hat, Inc.

- [Privacy Policy](https://www.redhat.com/en/about/privacy-policy)
- [Terms of Use](https://www.redhat.com/en/about/terms-use)
- [All policies and guidelines](https://www.redhat.com/en/about/all-policies-guidelines)

[![Red Hat Summit](https://www.redhat.com/cms/managed-files/Logo-Red_Hat-Summit-A-Standard-RGB-02_0.svg)](https://www.redhat.com/en/summit?intcmp=7013a0000026RhqAAE)

x

#### Subscribe now

Get the highlights in your inbox every week.

