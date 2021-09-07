# Ops by pull request: an Ansible GitOps story

March 30, 2020
by [Timothy Appnel](https://www.ansible.com/blog/author/timothy-appnel)

![ansible-blog_automated-webhooks-series](https://www.ansible.com/hs-fs/hubfs/Images/blog-social/ansible-blog_automated-webhooks-series.png?width=1035&name=ansible-blog_automated-webhooks-series.png)

In a [previous blog post](https://www.ansible.com/blog/intro-to-automation-webhooks-for-red-hat-ansible-automation-platform) I introduced Automation Webhooks and their uses with Infrastructure-as-Code (IaC) workflows and Red Hat Ansible Automation Platform. In this blog post, I’ll cover how those features can be applied to creating GitOps pipelines, a particular workflow gaining popularity in the cloud-native space, using Ansible and the unique benefits utilizing Ansible provides.

## What is GitOps?

Like so many terms that evolve and emerge from the insights and practices of what came before it, finding a definitive meaning to the term “GitOps” is a bit elusive.

GitOps is a workflow whose conceptual roots started with [Martin Fowler’s comprehensive Continuous Integration overview in 2006](https://martinfowler.com/articles/continuousIntegration.html) and descends from Site Reliability Engineering (SRE), DevOps culture and Infrastructure as Code (IaC) patterns. What makes it unique is that GitOps is a prescriptive style of Infrastructure as Code based on the experience and wisdom of what works in deploying and managing large, sophisticated, distributed and cloud-native systems. So you can implement git-centric workflows where you treat infrastructure like it is code, but it doesn’t mean it’s GitOps.

The term GitOps was coined by Alexis Richardson, CEO and Founder of Weaveworks, so a lot of how I’m going to define GitOps here comes directly from Alexis and Weaveworks. [This initial blog post](https://www.weave.works/blog/gitops-operations-by-pull-request) that explains the concept puts some baseline ideas out there, but doesn’t provide a concise definition. Depending who you ask you’ll get varying explanations of the term.

After reading and listening to various explanations, I thought this definition was a concise one that captures the essence of the many ways that GitOps has been explained:

_[GitOps] works by using Git as a single source of truth for declarative infrastructure and applications._

\-\- Weaveworks, “Guide To GitOps”

Often you will read immutable architectures, and even more specifically Kubernetes cluster management, as attributes of GitOps. I propose this description is a bit too prescriptive and limiting. You should embrace immutable architectures and Kubernetes in how you deploy applications and services if you can. That certainly makes a GitOps workflow easier to implement effectively, but here we will treat them as suggestions and preferences, not requirements of GitOps. We’ll see why in a bit.

This diagram shows what a typical GitOps workflow looks like from a conceptual level -- automating delivery pipelines to roll out changes to your infrastructure when changes are made in a Git repository.

![Screen Shot 2020-03-30 at 11.03.33 AM](https://www.ansible.com/hs-fs/hubfs/Screen%20Shot%202020-03-30%20at%2011.03.33%20AM.png?width=1584&name=Screen%20Shot%202020-03-30%20at%2011.03.33%20AM.png)

GitOps has been shown to increase productivity and velocity of deployments and development. Developers can use the tools and workflows they are already familiar with to manage deployments, allowing new developers to get up to speed faster. While Git has traditionally been a developer tool, operations staff benefit from the accumulated knowledge and experience of the Git community and the maturity of its ecosystem. There are a plethora of existing tools out there to make using Git more accessible and easier for those new to it.

Using GitOps brings together deployments and operations with development processes and tooling, providing a consistent means of working in your organization.

With the defined state of your infrastructure under Git version control, complete with a useful audit log of all activity, implementing a GitOps workflow will improve stability and increase reliability. Your organization benefits from being able to track what and when changes happen to help ensure compliance. Git tooling provides stronger security guarantees with its cryptography functions to track these changes and signing to prove its origins.  When things go wrong, you can easily revert and roll back to the previous state or rebuild your systems if you need to recover from a disaster.

## GitOps the Ansible Way

[Previously](https://www.ansible.com/blog/intro-to-automation-webhooks-for-red-hat-ansible-automation-platform), we reviewed how you can create Git-centric Infrastructure as Code (IaC) deployment workflows with the Automation Webhooks capabilities in Ansible Tower. GitOps is a more prescriptive workflow of IaC though and conceptually looks something like this diagram.

![Screen Shot 2020-03-30 at 11.03.54 AM](https://www.ansible.com/hs-fs/hubfs/Screen%20Shot%202020-03-30%20at%2011.03.54%20AM.png?width=1516&name=Screen%20Shot%202020-03-30%20at%2011.03.54%20AM.png)There are a few things to note about using Red Hat Ansible Automation Platform compared to the typical GitOps pipeline.

Here Ansible Tower replaces the GitOps “agent” (Operator) that runs on a given cluster and pulls in its configuration (state) from Git. Typically these are Kubernetes Operators like [Flux](https://github.com/fluxcd/flux) or [Eunomia](https://github.com/KohlsTechnology/eunomia).

Ansible Tower can work with Operators running on a Kubernetes cluster for a push/pull sort of approach. Ansible Tower pushes the configuration to the Operators via a Custom Resource (CR) and then the Operator pulls in any container images from the registry and handles whatever setup is necessary. The  operators here are made for a specific Kubernetes application or service and are useful outside of a GitOps pipeline rather than one for all configurations and management of the cluster to facilitate GitOps.

Using Ansible also provides the flexibility to apply GitOps workflow principles to systems other than Kubernetes, such as public/private cloud services and networking infrastructure, because you’re not required to use an Operator “agent” on that infrastructure.

### Advantages of Using Ansible

There are many tools you can use in your GitOps pipelines; however, Ansible provides some unique advantages that make it ideal for these workflows and extending their use beyond Kubernetes and cloud-native systems.

- **GitOps beyond Kubernetes.** Like Kubernetes, Ansible is a desired state engine that enables declarative modeling of traditional IT systems without scripting through Ansible Roles and Playbooks. With [the k8s module](https://docs.ansible.com/ansible/latest/modules/k8s_module.html) and [many others](https://docs.ansible.com/ansible/latest/modules/list_of_all_modules.html), an Ansible user can manage applications on Kubernetes, on existing IT infrastructure or across both with one simple language.
- **Agentless GitOps.** Consistent with the Ansible way of doing things, there is no requirement of a specialized GitOps Operator (agent) to facilitate the reconciliation of desired state on your systems.
- **Flexibility and Freedom.** You also have the flexibility and freedom to use the best tools for your needs and tailor your pipelines more to how you want to work. [Ansible excels as IT automation glue](https://www.ansible.com/blog/ansible-it-automation-glue).
- **Existing Skills & Ecosystem.** The same tried and trusted Ansible tooling lets you automate and orchestrate your applications across both new and existing platforms allowing teams to transition without having to learn new skills.

The benefits of using Ansible in this domain doesn’t end here though. There are a lot of [benefits for using Ansible in a cloud-native Kubernetes environments](https://www.ansible.com/blog/how-useful-is-ansible-in-a-cloud-native-kubernetes-environment).

## In Closing

GitOps works by using Git as a single source of truth for declarative infrastructure and applications. It is a workflow whose conceptual roots descend from Site Reliability Engineering, DevOps culture and Infrastructure as Code (IaC) patterns. GitOps is a more prescriptive workflow of IaC based on the experience and wisdom of what works in deploying and managing large sophisticated distributed systems.

Using Red Hat Ansible Automation Platform to implement GitOps pipelines provides unique benefits. Utilizing the Automation Webhook capabilities in Ansible Tower, you can implement agentless GitOps workflows that go beyond just cloud-native systems and manage existing IT infrastructure such as cloud services and networking gear. Using Ansible, enables you to tap into the existing Ansible ecosystem and the flexibility and freedom to use the best tools for how you want to work.

We hope you give GitOps with Ansible a try and see how beneficial and powerful this can be to your organization.
