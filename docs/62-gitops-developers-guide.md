
# A Developer’s Guide to GitOps

David Thor - 01/11/21 From: https://www.architect.io/blog/gitops-developers-guide

One of a modern DevOps team’s driving objectives is to help developers deploy features as quickly and safely as possible. This means creating tools and processes that do everything from provisioning private developer environments to deploying and securing production workloads. This effort is a constant balance between enabling developers to move quickly and ensuring that their haste doesn't lead to critical outages. Fortunately, both speed and stability improve tremendously whenever automation, like GitOps, is introduced.

As you might have guessed from that lead-up, GitOps is a tactic for automating DevOps. More specifically, however, it's an automation tactic that hooks into a critical tool that already exists in developers’ everyday workflow, Git. Since developers are already committing code to a centralized Git repo (often hosted by tools like GitHub, GitLab, or BitBucket), DevOps engineers can wire up any of their operational scripts, like those used to build, test, or deploy applications, to kick off every time developers commit code changes. This means developers get to work exclusively with Git, and everything that helps them get their code to production will be automated behind the scenes.

## Why GitOps?

In years past, DevOps and CI/CD practices were a set of proprietary scripts and tools that executed everyday tasks like running tests, provisioning infrastructure, or deploying an application. However, the availability of new infrastructure tools like Kubernetes combined with the proliferation of microservice architectures have enabled and ultimately *demanded* that developers get more involved in CI/CD processes.

This *shift left* exploded the problems seen with custom scripting and manual execution leading to confusing/inconsistent processes, duplication of efforts, and a drastic reduction in development velocity. To take advantage of cloud-native tools and architectures, teams need a consistent, automated approach to CI/CD that would enable developers to:

- Stop building and maintaining proprietary scripts and instead use a universal process
- Create apps and services faster by using said universal deploy process
- Onboard more quickly by deploying every time they make code changes
- Deploy automatically to make releases faster, more frequent, and more reliable
- Rollback and pass compliance audits with declarative design patterns


## Developers love GitOps

For all the reasons cited above (and more), businesses need manageable and automatable approaches to CI/CD and DevOps to succeed in building and maintaining cloud-native applications. However, if automation is all that’s needed, why GitOps over other strategies (e.g., SlackOps, scheduled deployments, or simple scripts)? The answer is simple: developers love GitOps.


### One tool to rule them all, Git

It's become apparent in the last few years that GitOps is among the most highly-rated strategies for automating DevOps by developers, and it's not hard to see why. Developers **live** in Git. They save temporary changes to git, collaborate using git, peer-review code using git, and store a history and audit trail of all the changes everyone has ever made in git. The pipelining strategy described above was tailor-made for git. Since developers already rely on git so heavily, these processes are, in turn, tailor-made for developers. Developers recognize this and are more than happy to reduce the tools and processes they need to use and follow to do their jobs.


### Declared alongside code

Beyond just the intuitive, git-backed execution flow, another part of modern CI tools and GitOps that developers love is the declarative design. The previous generation of CI tools had configurations that lived inside private instances of the tools. If you didn't have access to the tools, you didn't know what the pipelines did, if they were wrong or right, how or when they executed, or how to change them if needed. It was just a magic black box and hard for developers to trust as a result.

In modern CI systems, like the ones most commonly used to power GitOps like [CircleCI](https://circleci.com/), [Github Actions](https://docs.github.com/en/free-pro-team@latest/actions), [Gitlab CI](https://about.gitlab.com/stages-devops-lifecycle/continuous-integration/), etc., the configurations powering the pipelines live directly in the Git repository. Just like the source code for the application, these configurations are version controlled and visible to every developer working on the project. Not only can they see what the pipeline process is, but they can also quickly and easily make changes to it as needed. This ease of access for developers is critical since developers write the tests for their applications and ensure it is safe and stable.


### Completely self-service

New features or bug fixes aren't considered complete until they land in production. This means that anything standing in the way of getting code changes to production are eating up developer time and mental energy when the feature, as far as the developer is concerned, "works on my machine.” Suppose developers have to wait, even for a few minutes, for a different team or individual to do some task before they can close out their work. In that case, it creates both friction and animosity in the organization.

Alleviating this back and forth between teams is one of the main benefits of DevOps automation tactics like GitOps. Not only do developers get to work in a familiar tool, but the ability to have their code make its way to production without manual intervention means they are never waiting on someone else before they can complete their tasks.


### Continuous everything

Yet another big perk of GitOps is that all the processes are continuously running all the time! Every change we make triggers tests builds, and deployments without ANY manual steps required. Since developers would use git with or without GitOps, hooking into their existing workflow to trigger DevOps processes is the perfect place to kick off automated events. Until developers stop using Git, GitOps will remain the ideal way to instrument automated DevOps.


## GitOps in practice

Naturally, the involvement of developers in the process has led teams to explore the use of developer-friendly tools like Git, but the use of Git as a source of truth for DevOps processes also creates a natural consistency to the shape of CI/CD pipeline stages. There are only so many hooks available in a Git repository after all (e.g., commits, pull requests open/closed, merges, etc.), so the look and feel of most GitOps implementations include a set of typical stages:

![GitOps Pipelines](https://www.architect.io/images/blog/a-developers-guide-to-gitops/gitops-pipeline.png)


### 1. Pull requests, tests, and preview environments

After developers have spent time writing the code for their new feature, they generally commit that code to a new Git branch and submit a [pull request](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests) or [merge request](https://docs.gitlab.com/ee/user/project/merge_requests/getting_started.html) back to the mainline branch of the repository. This is something developers already do daily to prompt engineering managers to review the code changes and approve them to be merged into the main application code. Since developers already follow this kind of process for their daily collaboration efforts, it's a perfect opportunity for DevOps to wire up additional tasks.

By hooking into the open/close events created by this pull request process using a continuous integration (CI) tool, DevOps teams can trigger the execution of unit tests, creation of preview environments, and execution of integration tests against that new preview environment. Instrumentation of these steps allows engineering managers to establish trust in the code changes quickly and allows product managers to see the code changes via the preview environment before merging. Faster trust development means faster merges, and earlier input from product managers means easier changes without complicated and messy rollbacks. This GitOps hook is a key enabler for faster and healthier product and engineering teams alike.


### 2. Merge to master and deploy to staging

Once all parties have reviewed the changes, the code can be merged into the mainline branch of the repository alongside changes from the rest of the engineering team. This mainline branch is often used as a staging ground for code that is almost ready to go to production, and as such, it’s another ideal time for us to run some operational tasks like tests and deployment. While we tested the code for each pull request before it was merged, we'll want to rerun tests to ensure that code works with the other changes contributed by peer team members. We'll also want to deploy all these changes to a shared environment (aka "staging") that the entire team can use to view and test the latest changes before they are released to customers.


### 3. Cut releases and deploy to production

Finally, after product and engineering have had time to review and test the latest changes to the mainline branch, teams are ready to cut a release and deploy to production! This is often a task performed by a release manager – a dedicated (or rotating) team member tasked with executing the deploy scripts and monitoring the release to ensure that nothing goes wrong in transit. Without GitOps, this team member would have to know where the proper scripts are, in what order to execute them, and would need to ensure their computer has all the correct libraries and packages required to power the scripts.

Thanks to GitOps, we can wire up this deployment to happen on another Git-based event – creating a [release](https://docs.github.com/en/free-pro-team@latest/github/administering-a-repository/about-releases) or tag. All a release manager would have to do is create a new "release,” often using semver for naming, and the tasks to build and deploy the code changes would be kicked off automatically. Like most tasks executed by a CI tool, these would be configured with the scripts’ location and order the libraries and packages needed to execute them.


## GitOps tooling

A solid and intuitive continuous integration tool isn't the only thing needed to instrument GitOps processes like those described in this article. The CI system can activate scripts based on git events, but you still need strong tools to power those scripts and ensure they can be run and maintained easily and safely. Deploying code changes (aka continuous delivery (CD)) is one of the most challenging steps to automate, so we've curated a few tooling categories that can help you through your GitOps journey:


### Containerization with Docker

Docker launched cloud development into an entirely new, distributed landscape and helped developers begin to realistically consider microservice architectures as a viable option. Part of what made Docker so powerful was how developer-friendly it is compared to the previous generation of virtualization solutions. Just like the declarative CI configurations that live inside our repositories, developers simply have to write and maintain a `Dockerfile` in their repository to enable automated container builds of deployable VMs. Containerization is an enormously powerful tactic for cloud-native teams and should be a staple tool in your repertoire.


### Infrastructure-as-code (IaC)

A lot goes into provisioning infrastructure and deploying applications that isn't captured by a `Dockerfile`. For everything else, there's infrastructure-as-code (IaC) solutions like [Terraform](https://www.terraform.io/), [Cloudformation](https://aws.amazon.com/cloudformation/), and others. These solutions allow developers to describe the other bits of an application, like Kubernetes resources, load balancers, networking, security, and more, in a declarative way. Just like the CI configs and Dockerfiles described earlier, IaC templates can be version controlled and collaborated on by all the developers on your team.


### DevOps automation tools like Architect

I really can't talk about DevOps automation without talking about Architect. We love IaC and use it heavily as part of our product. We found that configuring deployments, networking, and network security, especially for microservice architectures, can be demanding on the developers who should be focused on new product features instead of infrastructure.

Instead of writing IaC templates and CI pipelines, which require developers to learn about Kubernetes, Cilium, API gateways, managed databases, or other infrastructure solutions, just have them write an `architect.yml` file. We'll automatically deploy dependent APIs/databases and securely broker connectivity to them every time someone runs `architect deploy`. Our process can automatically spin up private developer environments, automated preview environments, and even production-grade cloud environments with just a single command.


## Learn more about DevOps, GitOps, and Architect!

At Architect, our mission is to help ops and engineering teams simply and efficiently collaborate and achieve deployment, networking, and security automation all at once. Ready to learn more? Check out these resources:

- [Creating Microservices: Nest.js](https://www.architect.io/blog/creating-microservices-nestjs)
- [The Importance of Portability in Technology](https://www.architect.io/blog/the-importance-of-portability)
- [Our Product Docs!](https://www.architect.io/docs)

Or [sign up](https://cloud.architect.io/signup) and try Architect yourself today!

------

Contents

- Why GitOps?
- Developers love GitOps
- GitOps in practice
- GitOps tooling
- Learn more about DevOps, GitOps, and Architect!


