# How to Champion GitOps in Your Organization

August 04, 2020

By now you’ve heard all about GitOps and are convinced that GitOps is the most efficient way for development teams to go faster without them having to become Kubernetes gurus. However, making the switch to a cloud native technical solution may be the simplest part in your Kubernetes journey. Getting buy-in from your peers and championing GitOps throughout your organization could very well be the most challenging aspect of the cloud native transition.

At a recent GitOps Days virtual event held earlier this spring, we hosted a roundtable discussion with four Kubernetes and GitOps hands-on practitioners. Most of our panelists had recently implemented GitOps and self-service development platforms for Kubernetes in their organizations. In this discussion, the panelists offer a lot of common sense advice on the pitfalls to avoid when engineering a platform, and more importantly, they dig down on some practical strategies to use when you’re on-boarding and educating development teams who are adopting GitOps for self-service developer platforms in your own organization.


## Our panel of GitOps and Kubernetes practitioners

We were very fortunate to host Niraj Amin, Director, Cloud Platform Architect at Fidelity Investments, Steve Wade ( [@swade1987](https://twitter.com/swade1987)), the Platform Lead at Mettle, Javeria Khan ( [@javeriak\_](https://twitter.com/javeriak_)), Senior SRE at Palo Alto Networks, and Kingdon Barrett ( [@yebyen](https://twitter.com/yebyen)), Application Developer at University of Notre Dame OIT. Cornelia Davis, Weaveworks, CTO moderated the panel.

## What do we mean by a self-service developer platform?

Before we discuss how you go about championing GitOps in your own organization, let’s step back and talk about what we even mean when we say “self-service Kubernetes platform”.

Platform teams in an organization are increasingly responsible for providing a set of developer services to developers. Developers on the other hand are responsible for delivering applications to the company’s end-consumers. And so, to a large extent what we’re exploring is the relationship between those teams within an organization.


### Guardrails and security in place

At Fidelity, a dedicated platform team manages their Kubernetes implementation. As a platform team they serve the needs of developers. In particular their job is to get out of the way of the developer so they can do their job as efficiently as possible.



Niraj clarified, “...when I talk about platforms, what we’re really talking about is the infrastructure component of things. Obviously both EKS and Kubernetes play a big role. We also have fifteen or sixteen different components like the ELB ingress controller, external DNS and other components that are open source, plus we provide the autoscalers to developers as well.”



At Fidelity, GitOps allows for all platform configurations to be bundled and versioned in Git. Developers never have to worry about operational tasks, as upgrades are now automated. The result is a standard platform with guardrails and security in place that any development team can deploy to and from.



Similarly at Palo Alto, Javeria Khan, Senior SRE says that “as platform builders we are trying to solve infrastructure issues and remove that burden from the developer. We especially want to ensure that they don’t inadvertently make a change that can cause harm to the entire system.“



Steve Wade from Mettle says that implementing GitOps is a way of providing an abstraction layer on top of Kubernetes, “For us the platform needs to enable a self-service mechanism for developers. Essentially the platform is there for them to be able to bring business value to Mettle and our customers. We use GitOps as an abstraction layer for developers to onboard new microservices into the platform.”


## Balance between control and flexibility

When implementing a developer platform, there is tension between the need for control while at the same time providing some flexibility in tool choice and in certain cloud native patterns. Maintaining the balance between those two elements can be tricky but all participants agreed on the need for implementing constraints and guardrails.



Throughout the discussion, the panelists boiled down their advice into these five practical tips for championing GitOps in your own organization:


### \#1 Define, collaborate and document common cloud native patterns

At Fidelity the platform team is very transparent in terms of what security they have in place and why it needs to be there. To enforce those security requirements and to be more flexible, they need to support multiple cloud native patterns. For example, they document and provide support for numerous strategies on how to manage secrets and how you might go about using persistent data, among others.



The Mettle platform team developed and documented a number of microservices patterns together with the development team as a way to illustrate the problems with doing things the old way vs the new way. To keep the information flowing, the platform team at Mettle created a wiki to document all of these patterns.


### \#2 Take small steps & iterate the process

One of the best ways to start according to Steve is with a component that you're familiar with and that has a small blast radius. In this way, if that one thing doesn't work out, then it's not going to be too difficult to regroup.



For [Mettle’s GitOps journey](https://www.weave.works/blog/case-study-mettle-leverages-gitops-for-self-service-developer-platform), they began with the platform workload, leaving the developer workloads, and focused on how to update the platform itself. They iterated on a number of different approaches for different aspects: how to deploy the ingress controller; or how to deploy the Prometheus monitoring stack.

> “As an operator and as a platform architect, I started small. That means rolling out to staging and Dev environments first, for example, when you're adding GitOps tools like Flux or Flagger to your Kubernetes environments, enable them on staging clusters first to get a feel of what you like and what you don't like about it. This allows you to decide which features tie in better with your environment and how you’re going to integrate them with your workflows.” -- **Javeria Khan, Senior SRE, Palo Alto**

### \#3 Develop a good UX from your local machine on to staging

The other important part is to develop some sort of sandbox environment on developer machines so they can experiment with processes on their own. There are many different types of tools they can use such as [Kind](https://kind.sigs.k8s.io/), [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) and maybe even [Ignite](https://github.com/weaveworks/ignite). In addition to those tools, you can also take advantage of public Helm charts and public images for experimentation. Steve suggests to deploy an NGINX ingress controller in Minikube using [Flux](https://www.weave.works/oss/flux/). And once you’ve built the path, developers will understand the processes much better.


### \#4 Host brown bag information sessions

Javeria Khan, in addition to documenting cloud native patterns, took it one step further and suggests that people often learn in different ways. Javeria suggests organizing and recording brown bag sessions for any patterns and strategies you collaborate on and then saving those recordings so that people can view them in their own time.


### \#5 Over communicate changes

All panelists agreed that you can never over-communicate your changes. Experiment with different ways of communicating changes, both written and verbal, and both as formal and informal presentations.

View the entire panel for more questions and answers on the length of our panelists journey and other great questions asked from the viewing audience:

 For more tips on teaching your team GitOps you can refer to the [GitOps Conversation Kit](https://gitops-community.github.io/kit/), join one of our [upcoming webinars](https://www.weave.works/press/events/) or invite the [Weaveworks team for personalized training](https://www.weave.works/services/kubernetes-training/).

[Contact us for information on how to champion GitOps in your organization.](https://www.weave.works/services/kubernetes-training/)

* * *

## About Anita Buehrle

Anita has over 20 years experience in software development. She’s written technical guides for the X Windows server company, Hummingbird (now OpenText) and also at Algorithmics, Inc. She’s managed product delivery teams, and developed and marketed her own mobile apps. Currently, Anita leads content and other market-driven initiatives at Weaveworks.
