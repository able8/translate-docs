# DevOps — You Build It, You Run It

Jun 27, 2019

https://medium.com/usa-today-network/devops-you-build-it-you-run-it-8f972343eb8e

We know DevOps is a paradigm shift for software development. But the wide range of DevOps definitions and processes make it difficult, especially across organizations, to define what DevOps is in practice. Is it a development team working well with an operations team? Or is it a third team connecting the others? The API Services team at USA TODAY NETWORK doesn’t fit either of those scenarios. Our approach can be summed up by the slogan, “You Build It, You Run It.”

# One Team

DevOps at its core is about aligning the responsibility for running services with the authority to improve them. The quality and velocity of improvements for a service are greatly improved when responsibility and authority are in alignment. A dev team and an operations team which don’t collaborate are functionally broken, because the day-to-day responsibility rests on the operations team while the authority to make changes belongs to the dev team. One way to solve this is to improve communication between teams while building shared responsibility for running the service, along with shared authority to modify it. The approach we take is similar. We don’t have two teams. A single team both builds the service and runs it. It works well for us!

# Benefits

**Organizational Benefits**

Maintaining a single team with aligned responsibility and authority means we don’t need cross-team communication. Instead, we can rely on much easier intra-team communication. Individual skillsets still vary, but the goals and responsibilities are shared by the team as a whole. Everyone becomes more aligned. Also, cross-team communication is no longer tied up with service level concerns and becomes freed up for discussions regarding composition of services and other higher-level concerns.

Here are some other reasons we’ve found that “You Build It, You Run It” leads to better services:

- Improved issue visibility: A team responsible for running a service is careful to make it run well. (This is true for any DevOps organization, especially when you will be paged for incidents.)
- Reduced issue resolution time: You’re dealing with a single team, so delays caused by cross-team communication disappear and everyone involved understands both the operation and development of the service.
- Immediate feedback loop: The feedback loop for problems is automatic; no need to communicate out the cause of a problem, just immediately fix it. A faster feedback loop equals better service.

**Engineer Benefits**

Teams can be formed with engineers with various specialties. But whenever practical, we strive for every engineer to learn the skillsets needed for both operations and development. Building a whole team of highly-skilled and trusted engineers is a challenge, but it’s not insurmountable. Engineers like to be both highly-skilled and trusted, so setting this expectation — along with plenty of learning resources — drives growth.

Engineers embrace this approach because of the resulting skillset growth and job satisfaction opportunities. Here’s why:

- Execution is the best teacher. For example, devs learn to better incorporate instrumentation and security into an application when they also run the service.
- You can better build a more effective and relevant product when you see and truly understand the whole picture. It also allows you to make wise tradeoffs on features, development speed, security, etc. In short, you become a better engineer. Understanding the whole is also more satisfying than working on a portion with only limited context.
- Building a service paired with seeing it used — as well as using it yourself — is satisfying.

# Beyond One Team

The success we’ve described isn’t possible without a set of robust, well-maintained services to build on. Today’s cloud-native infrastructure allows a single team to run an application that requires a CI system, database, Kubernetes, etc. An engineer can confidently use these services and rarely have to worry about service-specific details, so her time is freed up for other concerns. When details for a dependent service require investigation, the service maintainer can be brought in for assistance. For these reasons, building on top of robust services prevents the amount of work in the system from growing beyond what a single team can handle.

Even with cloud-native services, nearly every company has plenty of work for multiple teams. Embracing an “as a service” team organization model can help teams divide and conquer along lines that enable each to build and run their own services. If CI for your company is run internally, the team responsible can offer CI as a service. Other teams then consume it as a tool they use when building their own services. This enables both the CI team and teams that use that service to follow the “You Build It, You Run It” model of DevOps.

# Conclusion

We’ve found that “You Build It, You Run It” doesn’t just bridge the fence between development and operations teams, but rather pulls that fence down altogether and embodies the spirit of DevOps. Immediate feedback on changes leads to quick fixes and better services. Working this way enables engineers to become better, grow in their skillset and improve their job satisfaction.


## [More from USA TODAY NETWORK](http://medium.com/usa-today-network?source=post_page-----8f972343eb8e--------------------------------)

Gannett is the largest local news organization in America. Our brands include USA TODAY and 250+ newsrooms spanning 46 states. Gannett’s vastly expanded local-to-national platform reaches over 50% of the U.S. digital audience, including more Millennials than Buzzfeed.
