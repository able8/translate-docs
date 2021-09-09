# Why I've Been Merging Microservices Back Into The Monolith At InVision

December 21, 2020

If you follow me on Twitter, you may notice that every [now (1)](https://twitter.com/BenNadel/status/1097596366321303552) [and (2)](https://twitter.com/BenNadel/status/1321530362909044737) [then (3)](https://twitter.com/BenNadel/status/1338904547159379970) I post a celebratory tweet about merging one of our microservices **back into the monolith** at [InVision](http://www.bennadel.com/invision/co-founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision is the digital product design platform used to make the world's best customer experiences."). My tweets are usually accompanied by a Thanos GIF in which Thanos is returning the last Infinity Stone to the Infinity Gauntlet. I find this GIF quite fitting as the reuniting of the stones gives Thanos immense power; much in the same way that _reuniting the microservices_ give _me and my team_ power. I've been asked several times as to why it is that I am killing-off my microservices. So, I wanted to share a bit more insight about this particular journey in the world of web application development.

![Thanos returning the last stone to the Infinity Gauntlet](https://bennadel-cdn.com/resources/uploads/2020/thanos-inserting-the-last-inifinity-stone.gif)

### I Am Not "Anti-Microservices"

To be very clear, I wanted to start this post off by stating unequivocally that I am **not anti-microservices**. My merging of services back into the monolith is not some crusade to get microservices out of my life. This quest is intended to **"right size" the monolith**. What I am doing is _solving a pain-point_ for my team. If it weren't reducing friction, I wouldn't spend **so much time (and opportunity cost)** lifting, shifting, and refactoring old code.

![Tweet highlight: 3 weeks and about 40 JIRA tickets worth of effort.](https://bennadel-cdn.com/resources/uploads/2020/merging-microservices-effort-is-not-free.png)

Every time I do this, I run the risk of introducing new bugs and breaking the user experience. Merging microservices back into the monolith, while sometimes exhilarating, it _always terrifying_; and, represents a _Master Class_ in planning, risk reduction, and testing. Again, if it weren't worth doing, I wouldn't be doing it.

### Microservices Solve Both Technical _and_ People Problems

In order to understand why I am destroying some microservices, it's important to understand why microservices get created in the first place. Microservices solve two types of problems: **Technical problems** and **People problems**.

A **Technical problem** is one in which an aspect of the application is putting an undue burden on the infrastructure; which, in turn, is likely causing a poor user experience (UX). For example, image processing requires a lot of CPU. If this CPU load becomes too great, it could start starving the rest of the application of processing resources. This could affect system latency. And, if it gets bad enough, it could start affecting _system availability_.

A **People problem**, on the other hand, has little to do with the application at all and everything to do with how your team is organized. The more people you have working in any given part of the application, the slower and more error-prone development and deployment becomes. For example, if you have 30 engineers all competing to "Continuously Deploy" (CD) the same service, you're going to get a lot of queuing; which means, a lot of engineers that _could otherwise be shipping product_ are actually sitting around waiting for their turn to deploy.

### Early InVision Microservices Mostly Solved "People" Problems

[InVision](http://www.bennadel.com/invision/co-founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision is the digital product design platform used to make the world's best customer experiences.") has been a monolithic system since its onset 8-years ago when **3 engineers** were working on it. As the company began to grow and gain traction, the _number of systems_ barely increased while the _size of the engineering team_ began to grow rapidly. In a few years, we had _dozens_ of engineers - both back-end and front-end - all working on the same codebase and all deploying to the same service queue.

As I mentioned above, having a lot of people all working in the same place can become very problematic. Not only were the various teams all competing for the same deployment resources, it meant that every time an "Incident" was declared, several teams' code had to get rolled-back; and, _no team could deploy_ while an incident was being managed. As you can imagine, this was causing a _lot of friction_ across the organization, both for the engineering team and for the product team.

And so, "microservices" were born to solve the **"People problem"**. A select group of engineers started drawing boundaries around parts of the application that they felt corresponded to _team boundaries_. This was done so that teams could work more independently, deploy independently, and ship more product. Early InVision microservices had _almost nothing to do_ with solving technical problems.

### Conway's Law Is Good If Your Boundaries Are Good

If you work with microservices, you've undoubtedly heard of ["Conway's Law"](https://en.wikipedia.org/wiki/Conway%27s_law), introduced by [Melvin Conway](https://www.melconway.com/) in 1967. It states:

> Any organization that designs a system (defined broadly) will produce a design whose structure is a copy of the organization's communication structure.

This law is often illustrated with a "compiler" example:

> If you have four groups working on a compiler, you'll get a 4-pass compiler.

The idea here being that the solution is "optimized" around team structures (and team communication overhead) and not necessarily designed to solve any particular technical or performance issues.

In the world _before microservices_, Conway's Law was generally discussed in a negative light. As in, Conway's Law represented poor planning and organization of your application. But, in a _post-microservices_ world, Conway's Law is given much more latitude. Because, as it turns out, if you can break your system up into a set of **independent services** with **cohesive boundaries**, you can ship more product with fewer bugs because you've created teams that are much more focused working on set of services that entail a narrower set of responsibilities.

Of course, the _benefits_ of Conway's Law depend heavily on _where you draw boundaries_; and, how those boundaries **evolve over time**. And this is where me and my team - the **Rainbow Team** \- come into the picture.

Over the years, [InVision](http://www.bennadel.com/invision/co-founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision is the digital product design platform used to make the world's best customer experiences.") has had to evolve from both an organizational and an infrastructure standpoint. What this means is that, _under the hood_, there is an older "legacy" platform and a growing "modern" platform. As more of our teams migrate to the "modern" platform, the services for which those teams were responsible need to get handed-off to the remaining "legacy" teams.

Today - in 2020 - **my team is the legacy team**. My team has slowly but steadily become responsible for more and more services. Which means: fewer people but more repositories, more programming languages, more databases, more monitoring dashboards, more error logs, and more late-night pages.

In short, all the _benefits_ of Conway's Law for the organization have become **liabilities** over time for my "legacy" team. And so, we've been trying to "right size" our domain of responsibility, bringing balance back to Conway's Law. Or, in other words, we're trying to alter **our service boundaries** to match **our team boundary**. Which means, **merging microservices _back_ into the monolith**.

### Microservices Are Not "Micro", They Are "Right Sized"

Perhaps the worst thing that's ever happened to the microservices architecture is the term, "micro". Micro is a _meaningless_ but heavily loaded term that's practically dripping with historical connotations and human bias. A far more helpful term would have been, "right sized". Microservices were never intended to be "small services", they were intended to be "right sized services."

"Micro" is apropos of nothing; it means nothing; it entails nothing. "Right sized", on the other hand, entails that a service has been **appropriately designed** to meet its requirements: it is responsible for the "right amount" of functionality. And, what's "right" is not a static notion - it is dependent on the team, its skill-set, the state of the organization, the calculated return-on-investment (ROI), the cost of ownership, and the moment in time in which that service is operating.

For my team, "right sized" means fewer repositories, fewer deployment queues, fewer languages, and fewer operational dashboards. For my _rather small_ team, "right sized" is more about "People" than it is about "Technology". So, in the same way that [InVision](http://www.bennadel.com/invision/co-founder.htm?redirect=https%3A%2F%2Fwww%2Einvisionapp%2Ecom%2F%3Fsource%3Dbennadel%2Ecom "InVision is the digital product design platform used to make the world's best customer experiences.") originally introduced microservices to solve "People problems", my team is now destroying those very same microservices in order to solve "People problems".

The gesture is the same, the manifestation is different.

I am extremely proud of my team and our efforts on the legacy platform. We are small band of warriors; but we accomplish quite a lot with what we have. I attribute this success to our deep knowledge of the legacy platform; our **aggressive pragmatism**; and, our continued efforts to design a system that _speaks to our abilities_ rather than an attempt to expand our abilities to match our system demands. That might sound narrow-minded; but, it is the only approach that is tenable _for our team and its resources in this moment in time_.

### Epilogue: Most Technology Doesn't Have to "Scale Independently"

One of the arguments in favor of creating independent services is the idea that those services can then "scale independently". Meaning, you can be more targeted in how you provision servers and databases to meet service demands. So, rather than creating massive services to scale only a portion of the functionality, you can leave some services small while _independently_ scaling-up other services.

Of all the reasons as to why independent services are a "Good Thing", this one gets used very often but is, in my (very limited) opinion, usually irrelevant. Unless a piece of functionality is **CPU bound** or **IO bound** or **Memory bound**, independent scalability is probably not the _"ility"_ you have to worry about. Much of the time, your servers are _waiting for things to do_; adding "more HTTP route handlers" to an application is not going to suddenly drain it of all of its resources.

If I could go back and **redo our early microservice attempts**, I would 100% start by focusing on all the "CPU bound" functionality first: image processing and resizing, thumbnail generation, PDF exporting, PDF importing, file versioning with `rdiff`, ZIP archive generation. I would have broken teams out along those boundaries, and have them create "pure" services that dealt with nothing but Inputs and Outputs (ie, no "integration databases", no "shared file systems") such that every other service could consume them while maintaining loose-coupling.

I'm not saying this would have solved all our problems - after all, we had more "people" problems than we did "technology" problems; but, it would have solved some more of the "right" problems, which may have made life a bit easier in the long-run.

### Epilogue: Microservices _Also_ Have a Dollars-And-Cents Cost

Service don't run in the abstract: they run on servers and talk to databases and report metrics and generate log entries. All of that has a very real dollars-and-cents cost. So while your "lambda function" doesn't cost you money when you're not using it, your "microservices" most certainly do. Especially when you consider the redundancy that you need to maintain in order to create a "highly available" system.

My team's merging of microservices back into the monolith has had an actual impact on the bottom-line of the business (in a good way). It's not massive - we're only talking about a few small services; but, it's not zero either. So, in additional to all of the "People" benefits we get from merging the systems together, we _also get_ a dollars-and-cents benefit as well.
