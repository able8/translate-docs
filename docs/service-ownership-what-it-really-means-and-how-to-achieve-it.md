# Service Ownership: What It Really Means and How to Achieve It

Years ago, end-to-end software development involved dividing tasks based on where they fell in the system life cycle. One team wrote the code. Then another team deployed it to production. And yet another team monitored and maintained the service. This led to a lot of friction, needless handoffs, and bottlenecks.

Then DevOps came to the rescue, promising to reduce handoffs and improve operations. But without the missing ingredient of service ownership, DevOps won’t be as powerful as it could be.

Now what does it mean to own a service? And how can we support this organizational pattern within our teams? While sharing ideas on how to [take back DevOps](http://www.opslevel.com/blog/taking-back-devops/), we talked a bit about service ownership. As a follow-up, we’ll now dive deeper into service ownership and how we can spread it across our teams.

> Without the missing ingredient of **service ownership**,
>
> DevOps won’t be as **powerful** as it could be.

## What Does Service Ownership Mean?

To understand what service ownership provides, we’ll first want to understand the pain it solves.

### Without Service Ownership

To begin, let’s go back to the model mentioned in the intro where different teams were responsible for building, deploying, and managing a service.

As you can imagine, these [silos](https://en.wikipedia.org/wiki/Silo_(software)) led to friction, handoffs, and bottlenecks.

First, natural friction developed between development and ops teams. Development teams wanted to move features quickly, while operations teams wanted to slow the rate of change to ensure they weren’t buried in incidents and maintenance work. And release management played in the middle, trying to make both sides happy while following proper release protocol.

Second, handoffs between teams resulted in no one taking ownership of problems or improvements to the service’s overall life cycle. If a piece of the life cycle wasn’t theirs to own and improve, teams weren’t encouraged to look for solutions or improvements in that space.

For example, for many teams, observability that improved troubleshooting was always an afterthought. Furthermore, engineers didn’t always think about making their code easier to debug because they weren’t the ones who had to do that at three in the morning when the pager went off. And operations, the folks who could really use improved troubleshooting capabilities, didn’t have the autonomy to improve things themselves.

If all that wasn’t painful enough, bottlenecks formed because of these silos. Development teams waited on release management to ship the code. Release management was waiting on approval from someone to deploy the code. And once the code deployed, operations teams waited on development teams to fix bugs or improve reliability. And those items never came quickly enough for the ops team.

So how does service ownership help? Service ownership takes the responsibility of those three teams and combines it into one team’s responsibility. One team holds the responsibility for the product and not just a piece of the pipeline.

Perhaps you’ve heard the phrase “build it, ship it, run it.” That’s the two-second description of DevOps and service ownership. To expand on that, let’s break those components down.

> One team holds the **responsibility** for the product and
>
> not just a **piece** of the pipeline.

### Build It

The phrase “build it” seems simple enough. Our engineering teams all build software. So does this look the same for teams that have service ownership and those that don’t?

Not necessarily. When a team utilizes service ownership, they also gain autonomy and influence over how to build their product.

In fact, when teams own the whole service, you’ll find that they appreciate the value of building services with appropriate documentation, thorough tests, and observability practices more than when separate teams operate or run the services.

### Ship It

Next, when a team owns a service, they own the responsibility of shipping it to production.

This doesn’t mean that every team needs a [CI/CD](https://www.redhat.com/en/topics/devops/what-is-ci-cd) guru, with every team re-creating the basics of continuous deployment and automated compliance checks. It could be that the team uses common frameworks, patterns, and libraries provided by others.

However, the team does own the responsibility of getting code in production efficiently. They should also understand the pipelines and fix things when they go wrong. Overall, the goal is not to throw problems over the wall at the release management team, but to take ownership of shipping the product to production.

### Run It

Here, the team that writes the service is also the team that’s best equipped to run the service. They know how to resolve issues quickly and make changes that will prevent similar future issues.

When we run our own applications, we also realize the importance of operationalizing the product. So we’ll write code in a way that makes it more maintainable, observable, and debuggable. Because we’ll be the ones who have to fix it when it breaks.

## Getting Started

Next, what first steps can you take toward driving service ownership in your organization? In short, we need to move away from telling teams what to do or how to do it and move toward looking at outcomes necessary for customer satisfaction. Furthermore, when multiple teams contribute to an initiative, we need to point them all toward the same outcome. Then we must encourage them to collaborate and work through integrations and dependencies together, always pointing toward the same goal.

> When **multiple teams** contribute to an initiative, we need
>
> to point them all toward the **same outcome**.

So how can we do that?

First, share expectations around service ownership clearly. Talk about where you see the organization is at today and what you’re working toward. And mention the benefits that not only the customer or organization receives, but also the benefits that the development teams should expect regarding reduced bottlenecks. Then ask for feedback and help in getting there. And be clear that all the teams own the problem and will all contribute to the solution.

Once people align on expectations, look at gaps in tools, skills, and knowledge. What do your teams need to take ownership of their service? Do they need education regarding security so that they can prioritize concerns properly? Or do they need tools to automate scans or reduce unnecessary toil?

Assess skills and seed teams with people who can build up release management and operational knowledge.

**Service Ownership, Solved**

Say goodbye to stale spreadsheets or wikis! See **OpsLevel** in action to accelerate your Service Ownership journey.

[Request a Demo](http://www.opslevel.com/request-demo/)

## Common Mistakes in Service Ownership

Now that we know from a high level what service ownership looks like, let’s consider some common mistakes in the implementation of the pattern.

### Accountability Without Autonomy

If the team owns the service, they should have the autonomy and accountability to build, ship, and run it as they see fit. But autonomy and accountability require more than lip service. If you tell teams they’re accountable but then don’t let them have control over their own backlog, you’ve probably given the team accountability without autonomy.

For example, management may tell teams to “drop everything” and focus on one specific task or feature. And this task must be completed quickly, to the detriment of all others. This indicates that the team doesn’t have the autonomy to prioritize tasks or assess needs.

Service owners need to have the autonomy to prioritize the reduction of their own operational pain and toil, and balance these needs against feature enhancements and innovation. Additionally, the taking on of any new technical debt - usually through shortcuts due to strict time restraints - should be something the team takes on knowingly and willingly, and with an associated plan for dealing with it.

Therefore, instead of demanding specific work be done in a specific order, have the team focus on a target or outcome. For example, if a team focuses on helping the customer with a specific problem, they’re better equipped to prioritize tasks related to features, security, and tech debt.

> Service owners need to have the **autonomy** to prioritize the reduction of their own operational pain and toil, and **balance** these needs against feature enhancements and innovation.

### Lack of Support or Resources

If you take an existing software team that has relied on other teams to ship and run their product and tell them that this is all now their responsibility, you’ll lose trust. And they’ll lose sleep.

Instead, work to introduce change gradually. Start with awareness, letting the development team understand what it takes to truly own their product. And seed the team with people from operations or release management so they have someone there to help with the journey and expertise.

### Unclear Expectations

When teams rely on each other in a complex distributed ecosystem, unclear expectations can hurt trust and collaboration between teams. If our team has a dependency on another team but ownership and performance expectations aren’t clear, this can result in blame games and finger-pointing. Suppose one service is down or degraded due to a failure in a dependency. In that case, we need a clear line of sight into escalation processes, expectation on service-level objectives (SLOs), and an understanding of what that dependency team offers.

Services should have clearly defined objectives, contact information, and access to metrics and status all in one place. If your teams struggle with finding out what to expect from another service, who to contact, or what state the service is in, consider tools like OpsLevel to provide a better service ownership experience to your organization.

> When teams rely on each other in a **complex distributed ecosystem**,
>
> unclear expectations can **hurt trust** and collaboration between teams.

### Invisible Ownership

Sometimes we rely on word of mouth or team-specific methods of identifying service owners. Or disparate spreadsheets and wikis hold out-of-date information that’s difficult to find. Sometimes service ownership info hides away in hard-to-find architecture documents or data stores that lead you on a treasure hunt.

This results in frustration during incident response or even integration discussions.

To reduce the complexity, provide teams with one central and easy-to-use place to get the information they need. For this common problem, OpsLevel can help provide visibility through their [service catalog](http://www.opslevel.com/landing/microservice-catalog/).

Wherever the information lives, make sure that everyone knows where to find it and that the data stays up to date.

## Wrapping It Up

Service ownership provides many benefits, like encouraging teams to incorporate operational excellence early in the life cycle. Additionally, it gives the team autonomy to do the right thing and satisfy customer needs based on their knowledge and expertise.

To make service ownership easy to see, consider OpsLevel and how you can improve visibility and expectations. [Request a demo](http://www.opslevel.com/request-demo/) to see how OpsLevel can drive ownership today.

This post was written by Sylvia Fronczak. [Sylvia](https://sylviafronczak.com/) is a software developer and SRE manager that has worked in various industries with various software methodologies. She’s currently focused on design practices that the whole team can own, understand, and evolve over time.
