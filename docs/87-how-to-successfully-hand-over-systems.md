# How to Successfully Hand Over Systems

April 20th, 2021

In a product company, changes are inevitable so as to best support the strategy and the vision. Often during such a change, new teams are formed and other ones are restructured. While there are many challenges to be solved during a big change, there’s one in particular that’s often overlooked: system ownership.

Who will take ownership of the systems that were owned by a team that doesn’t exist anymore or that are better suited to be owned by another team? It’s in everyone’s interest that the ownership be given to a team familiar with the system’s domain, so that they can continue the maintenance and evolution.

Regardless of when the system handover will happen, how it’s executed is important, since the cost of failure can be high, and that could result in an outage or a significant amount of unplanned work.

Having experienced not-so-successful handovers — some of which took place over the course of a one-hour meeting — I was inspired to create a guideline that will help other teams do handovers differently. At the same time, my colleague Antonio N. went through an ownership change with his team. This also had few mishaps, so we joined forces to write a proposal document for doing system handovers at SoundCloud.

We used the RFC approach to gather input, experiences, and opinions from the entire organization. It was welcomed with enthusiasm and since then has been used multiple times.

I’m posting the complete guideline below not only to document what we did at SoundCloud, but also in hopes of providing a template for other companies to use when faced with a similar scenario.

## Guideline for Internal System Handovers

The guideline is a list of questions, tasks, and actions that the involved parties should consider as part of the system handover. The topics listed can be covered in different ways: through documentation, meetings, pairing sessions, workshops, tasks, PR reviews, etc. **The goal is to help the new team understand the what, why, and how of the system, and to empower them to maintain, change, and improve it.**

### General Recommendations

As the system ownership change is a process itself, we recommend that it’s driven by the new team with the help and support of the previous system owner. Both teams should collaborate on the planning and execution of the tasks.

There will be some new documents and artifacts produced as an outcome of the system ownership change. We recommend storing them in the system’s repo when possible, or else including a link in the repo (e.g. from the README file). This will help with both onboarding new team members and potential ownership changes in the future.

### Why Are We Changing the Ownership?

Help everyone involved in the handover understand why there’s a system ownership change. This impacts the team engagement in the handover process.

What’s best is to document the reasoning and to add additional information to the history of the system. In turn, this can reveal different things, such as if the ownership changed or underwent restructuring multiple times in a short period of time, or if the current organization isn’t set up to own such a system or if the system doesn’t belong to any team. Uncovering this information helps us ask important questions, such as is the system too complex, is it not in any team’s domain, or is it even needed anymore?

### What Does the System Do? What Problem Does It Solve? What Is the Vision?

Here, we’re looking to understand the system from a product perspective. It’s helpful to know some history about the system, how it evolved, and what its vision for the future is.

This could be a product session where product managers are involved. As an outcome of this session, it’d be nice to have a document to help onboard new people to the system.

### High-Level Architecture of the System

Get an overview of the system, the main components, and their interfaces. A more detailed diagram will probably lead to more detailed discussions. It’s best to have an online diagram so that it’s easily available for future reference.

### Use, Availability, and Criticality of the System

Get familiar with who’s using the system, what the criticality of the system is, and what it means if the system isn’t available. This is an opportunity to look into the available runbooks, metrics, and monitoring.

In cases when the previous and the new team owning the system are not part of the same [on-call rotation](https://developers.soundcloud.com/blog/building-a-healthy-on-call-culture), there **must** be an additional session where the system is introduced and explained to the engineers in the rotation group, so that all of them can respond to incidents related to the system. This helps prevent bigger outages and maintain the healthy on-call culture.

### Maintenance

As part of the maintenance of our system, we have daily, weekly, and monthly tasks that need to be executed by the team. In this section, we need to identify what those tasks are and what their periodicity is (if needed).

### Data Storage Overview

Most of the systems have their own data storage. When taking ownership of a system, the new team takes ownership of that data and the infrastructure that comes with it. This is to get an overview of what, how, and where that data is stored and/or purged.

### Batch (Offline) Jobs Overview

Some systems are using data that’s an outcome of a batch job. Also, many systems have batch jobs that produce datasets to be consumed for analytics or reporting. Get an overview of the batch jobs, their usage, outcomes, and maintenance.

### Decision History

This is helpful to understand the architecture choices and the evolution of the system, as well as to learn about any constraints the system might have. It’s best if this is documented.

We recommend using [Lightweight Architecture Decisions Records](https://www.thoughtworks.com/de/radar/techniques/lightweight-architecture-decision-records), which has the following format: [Context, Decision, and Consequences](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions).

### Tech Debt

Make sure you’re aware of the existing tech debt and you understand its implications. This is best if it’s documented.

### Known Bugs

When ownership changes, user-facing bugs will be reported to the new team; however, many of them might already be known. Make sure you understand them and why they’re present. This helps not only by decreasing the time to investigate, but also by providing a good service to our users.

### To Dos

In addition to the above, here are some tasks (not in order and not complete, since they are SoundCloud specific) that can ease the ownership change:

- **Project ownership update on GitHub**
- **Grant permissions to the new team**
- **Update offline jobs configuration**
- **Local development**
  - Can the engineers build the project locally?
  - Does the system have integration tests? If so, do they run in a local environment? Does the new owner need any additional information to run them?
- **Deployment**
  - How is the system deployed?
  - Is there a CI/CD?
- **Monitoring and Alerting**
  - Check if the monitoring graphs need to be updated
  - Update the system-related runbooks
  - In case the runbook location changes, please reflect that change in the corresponding alerts to avoid broken links from the alerts
  - Add the system to a corresponding on-call group and have a knowledge transfer session
  - Update alerts
  - Update PagerDuty
  - Our suggestion is to update the on-call rotation at the end, once the team has gained sufficient knowledge and confidence in the system

💡 The list of things to do is quite long. Take your time with each task and don’t rush. Use the help of the previous team, and pay attention to details and to the alerts.

💡 You can use your project management tool (e.g. JIRA) to track the progress of the handover. That will help both involved teams stay up to date on the status of the handover, the next steps, and when it will be completed.

💡 If you’ve discovered other helpful tasks or topics, please update the guideline with them.

## Usage

As the name suggests, the above document is a guideline, and it’s up to the parties involved in the system ownership change to decide **if they’re going to use it and how they’re going to use it**.

In most cases, ownership change is a collaborative process that enables the new owners to be motivated and have a solid understanding of the system to maintain and continue evolving it. In some cases, it can happen that there’s no one in the company that previously contributed to the system. However, even then, this guideline can help the team keep the focus on topics that are important to know, and not only on the codebase.

The most important thing is to not be judgmental of the choices others made and understand that, at that time, it was the best decision. For example, instead of making statements like “You could have used A instead of B,” or “You could have done it like this,” or even the harsher “That is wrong!” or “That is a huge mistake!”, try to be curious and ask open-ended questions like “What made you use A?” or similar.

I’d also recommend that the team taking ownership takes the time to go through each of the topics and gain understanding and knowledge — not only from an engineering perspective, but from the perspective of the product. One might think they can copy the guideline and fill in the sections, and that writing everything they know of and handing it over to the new team will complete the transfer. I would argue that this isn’t the intention, and I don’t believe it will have the same positive impact that can be seen when doing this collaboratively and dedicating time to exploration.

Additionally, the guideline is meant to be a live document, updated as teams are learning through the process.

## Side Effects / Other Impacts

This guideline should inspire teams to have useful and up-to-date documentation; a README on how to contribute, test, and run locally; and high-level architecture diagrams. This helps not only when changing ownership, but also when onboarding new team members.

Furthermore, it’s important to embrace the use of architectural decision records and help to reason about them in the future.

## Summary

This guideline exists to help engineering managers, product managers, and teams acknowledge that system ownership change is a process that should be well planned and done at a time that works best for everyone involved. It’s a process that requires effort and has its cost. However, it can inspire the organization to nurture a healthy engineering culture with a high bus factor and systems that are easy to maintain, evolve, and reason about.

- [Project Management](http://developers.soundcloud.com/blog/category/project%20management)
- [Engineering Management](http://developers.soundcloud.com/blog/category/engineering%20management)
- [Operations](http://developers.soundcloud.com/blog/category/operations)
