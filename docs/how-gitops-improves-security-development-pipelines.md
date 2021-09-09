# How GitOps Improves the Security of Your Development Pipelines

* * *

[GitOps](https://www.weave.works/technologies/gitops/) is usually discussed in terms of boosting developer velocity. But another benefit – one that doesn’t always get as much attention – concerns its potential to improve security.

At our recent virtual event, GitOps Days 2020, Maya Kaczorowski ( [@MayaKaczorowski](https://twitter.com/mayakaczorowski?lang=es)), GitHub Product Manager for software supply chain security, shared her thoughts on how GitOps can boost the security of your entire development pipeline.

To kick the session off, Maya talked about how DevOps and now GitOps leads to greater developer accountability. With the introduction of DevOps, developers took more responsibility for how their code was deployed and now, developers are becoming more responsible for security, too. Maya even suggested that modern DevOps can be thought of as ‘DevSecOps’. Whatever you want to call it, there’s no better way to give developers responsibility for operations and security than to adopt GitOps best practices.

GitOps gives you control over changes and allows you to verify them from a single source:

1. **Config as Code**

Using Git to manage YAML files makes it simple to check if you’re meeting security requirements. With access policies declared in a config file, you know who has access to what – and can easily verify it in code.
2. **Changes are** **auditable**

Version control means that you always know what you shipped and you can roll back at any time. Your commit history is an audit trail of comments, reviews, and a history of decisions that were made to your repo.
3. **Production matches the desired state kept in Git**

A single source of truth, with a common workflow for both code and infrastructure changes coupled with automatic alerts on a drift from the desired states increases reliability and removes the risk of human error.  A single set of tests, security scans, and permissions also help make changes secure and reliable.

![gitops-security.png](https://images.contentstack.io/v3/assets/blt300387d93dabf50e/blt1fa86e3d1618bc36/5fd92c867c43e43bf41983b3/gitops-security.png)

## Continuous security

Maya explained that tasks usually left until the end of the development cycle – like security testing – take place much earlier with GitOps. Testing for security is part of every iteration, which means errors can be caught earlier and vulnerabilities eliminated, long before any code is deployed to production. It therefore allows for what Maya calls ‘continuous security’. In parallel with continuous integration and continuous deployment, she sees this as a focus for integrating security best practice throughout the pipeline. Another way of looking at it is that while DevOps made developers responsible for any system outages they caused, treating security the same way makes developers responsible for data loss. And ultimately, the business goal of any security activity is preventing data loss.

## How GitOps makes this possible?

It comes down to the ability of GitOps tools to treat everything as code. If you can treat all configuration and security policy as code, everything can be held in version control. Changes can be made and reviewed , then fed through the automated pipeline that verifies, deploys and monitors the change. Any divergence from the desired state of the system – e.g. the emergence of a bug or a security vulnerability – will be caught much earlier, long before it can become a significant cost to the business.

## Securing the whole pipeline

GitOps improves the security of several elements on the development pipeline, from the code itself (and that includes anything else you keep in code, such as policy and config) to the process by which you make changes to it. So for example, if your compliance requirements are defined in YAML, Git ensures you’re meeting them by maintaining the desired state of the system. . With all changes captured in version control, you can roll back to a given point if you need to. The history of comments and reviews is also maintained, so not only will you know who made changes and when, but you’ll also know why. It’s a built-in audit trail.

When you get to production, the big security advantage is the single source of truth Git provides. It means you have a single set of tests, a single set of security scans and a single set of permissions to implement. Best of all, humans play no part in this process, which means human error is eliminated.

If your application does come under attack, you can take action immediately. Git provides the single source of truth, so you can redeploy everything instantly, if you need to. And of course, you can place controls and gates on this process, to meet any security or compliance needs specific to your organization.

## Respond faster

One of the key benefits of GitOps, however, is velocity. [the State of DevOps report](https://puppet.com/resources/report/2020-state-of-devops-report/) has proven that developers can move faster thanks to version control, continuous integration, test automation and other features available in Git. But that velocity isn’t limited to the speed at which developers can work. Maya suggested another take on velocity – one that can be thought of as an alternative to [Mean Time to Recovery](https://en.wikipedia.org/wiki/Mean_time_to_recovery): [Mean Time To Remediate](https://www.optiv.com/cybersecurity-dictionary/mttr-mean-time-to-respond-remediate). And with its auditable and full record of who did what, when nothing cuts Mean Time To Remediate like GitOps.

Development pipelines clearly represent a tempting attack vector for intruders. But by adopting GitOps, you can boost security right along the pipeline, while at the same time boosting the speed at which you can react if the worst happens.

View the full presentation:

[Go from Zero to GitOps with our discovery, design and deploy package for Kubernetes.](https://www.weave.works/services/gitops-design-services/)

* * *

## About

![ ](https://images.contentstack.io/v3/assets/blt300387d93dabf50e/blt1be4b5b42ea58cb4/58c02d7b48598d51743bf27e/weave-logo-512.png?format=webp&width=75)

Weaveworks’ mission is to empower developers and DevOps teams to build better software faster. Our “GitOps” model strives to optimize operational workflows; to make operations for developers simpler, better and faster.
