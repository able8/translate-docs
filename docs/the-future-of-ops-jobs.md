# The Future of Ops Jobs


Aug 17, 2020  12 Minute Read

Infrastructure, ops, devops, systems engineering, sysadmin, infraops, SRE, platform engineering. As long as I’ve been doing computers, these terms have been effectively synonymous. If I wanted to tell someone what my job was, I could throw out any one of them and expect to be understood.

Every tech company had a software engineering team that built software and an operations team that built infrastructure for running that software. At some point in the past decade they would have renamed this operations team to “devops” or “SRE”, but whatever the name, that was my team and those were my people.

But unless you’re an infrastructure company, infrastructure is not your mission. Which means that every second you devote to infrastructure work — and every engineer you devote to infrastructure problems — is a distraction from your core goals.

What’s more, it’s a distraction that builds on itself. The more time and energy you spend on infrastructure, the more your focus gets scattered, and the more you deprive yourself of the time and energy that ought to be devoted to the problems your business exists to solve.

This isn’t exactly new. Infrastructure and operations have always been a distraction from your core business problems. It used to be the case that every company had to grow internal expertise in hardware, datacenters, networking, operating systems, config management, and so on up the tech stack until reaching their core business problems. Infrastructure and ops have always been a distraction, but until recently, a necessary one. A means to an end.

So what’s changed? These days, you increasingly have a choice. Sure, you _can_ build all that internal expertise, but every day more and more of it is being served up via API on a silver platter.

Does this mean that operations is no longer important, no longer necessary? Far from it. Operations and operability are more important than ever. Let’s take a look at what’s happening to the ops profession at a high level, the emerging challenges we face, and the impact this’ll all have on our careers.

### **Changes afoot**

Beyond the broader move into cloud services, here are some of the major shifts on our horizon.

##### **From monolith to microservices**

Much has been written about the operational demands of a microservices architecture. Now that functions calling other functions involves a network hop, operational concerns are an unavoidable part of debugging even the most trivial problems. Microservices change the game from “building code” to “building systems”, which pushes more and more code writing into the realm of operations.

##### **From monitoring to observability**

Metric-based tools like Prometheus and DataDog are infrastructure monitoring tools, and quite good ones at that. When you’re responsible for infrastructure, the questions you care about are of aggregates and trends, thresholds and capacity. Monitoring tools are the right tool for the job, because that’s how you understand whether your infrastructure is healthy, and what actions to take to make it or keep it healthy.

Observability tools\*, on the other hand, are for the people writing and shipping code to users every day, and trying to inspect and understand behavior at the nexus of users, production, and code. Observability tools preserve the full context of the request. This allows you to slice and dice and tease out new correlations, as well as view events in a waterfall by time (“tracing”). Observability is how you connect the dots between software and real business impact, and between your engineers’ experience and your users’ experience.

_\*Caution: many monitoring tools are trying to rebrand themselves as observability tools without first building the necessary functionality. To tell the difference,_ [_see here_](https://www.honeycomb.io/blog/so-you-want-to-build-an-observability-tool/) _._

##### **From magic autoinstrumentation to instrumenting with intent**

Instrumentation is just another form of commenting and documenting your code. There are tools that promise to do it automatically for you, but they aren’t great at capturing intent.. Auto-instrumentation can tell you loads of ancillary details, but they will _never_ let you divine the business value intended by the engineer who built it. So suck it up and instrument your code.

### **Ops is dead. Long live ops.**

Ops teams are going the way of the dodo bird, yet operability, resiliency, and reliability have never been more important. The role of operations engineers is changing fast, and the role is bifurcating along the question of infrastructure. In the future, people who would formerly have called themselves “operations engineers” (or devops engineers) will get to choose between a role that emphasizes _building infrastructure software as a service_ and a role that uses their infrastructure expertise to help teams of engineers ship software more effectively and efficiently…generally by building as little infrastructure as possible.

If your heart truly beats for working on infrastructure problems, you’re in luck! — there are more of those than ever. Go join an infrastructure company. Try one of the many companies — AWS, Azure, all the many developer tooling companies — whose mission consists of building infrastructure software, of being the best in the world at infrastructure, and selling that expertise to other companies. There are roles for software engineers who enjoy building infrastructure solutions as a service, and there are even specialist ops roles for running and operating that infrastructure at scale, or administering those data services at scale. Whether you are a developer or not, working alone or in a team, [Azure DevOps training](https://acloudguru.com/course/introduction-to-azure-devops) can help you organize the way you plan, create and deliver software.

Otherwise, embrace the fact that your job consists of _building systems to enable teams of engineers to ship software that creates core business value_, which means home brewing as little infrastructure as possible. And what’s left?

##### **Vendor engineering**

Effectively outsourcing components of your infrastructure and weaving them together into a seamless whole involves a great deal of architectural skill and domain expertise. This skill set is both rare and incredibly undervalued, especially considering how pervasive the need for it is. Think about it. If you work at a large company, dealing with other internal teams should feel like dealing with vendors. And if you work at a small company, dealing with other vendors should feel like dealing with other teams.

Anyone who wants a long and vibrant career in SRE leadership would do well to invest some energy into areas like these:

- Learn to evaluate vendors and their products effectively. Ask piercing, probing questions to gauge compatibility and fit. Determine which areas of friction you can live with and which are dealbreakers.
- Learn to calculate and quantify the cost of your and your team’s time and labor. Be ruthless about shedding as much labor as possible in order to focus on your core business.
- Learn to manage the true cost of ownership, and to advocate and educate internally for the right solution, particularly by managing up to execs and finance folks.

##### **Product engineering**

One of the great tragedies of infrastructure is how thoroughly most of us managed to evade the past 20+ years of lessons in managing products and learning how to work with designers. It’s no wonder most infrastructure tools require endless laborious trainings and certifications. They simply weren’t built like modern products for humans.

- I recommend a crash course. Embed yourself within a B2B or B2C feature delivery team for a spell. Learn their rhythms, learn their language, soak up some of their instincts. You’ll need them to balance and blend your instincts for architectural correctness, scaling patterns, and firefighting.
- You don’t have to become an expert in shipping features. But you should learn the elements of relationship-building the way a good product manager does. And you must learn enough about the product lifecycle that you can help debug and divert teams whose roadmaps are hopelessly intertwined and whose roadmaps are grinding to a halt.

##### **Sociotechnical systems engineering**

The irreducible core of the SRE/devops skill set increasingly revolves around crafting and curating efficient, effective sociotechnical feedback loops that enable and empower engineers to ship code — to move swiftly, with confidence. Your job is not to say “no” or throw up roadblocks, it’s to figure out how to help them get to yes.

- Start with embracing releases. Lean hard into the deploy pipeline. The safest diff is the smallest diff, and you should ship automatically and compulsively. Optimize tests, CI/CD, etc so that deploys happen automatically upon merge to main, so that a single mergeset gets deployed at a time, there are no human gates, and everything goes live automatically within a few minutes of a developer committing their code. This is your holy grail, and most teams are nowhere near there.
- Design and optimize on-call rotations that load balance the effort fairly and sustainably, and won’t burn people out. Apply the appropriate amount of pressure on management to devote enough time to reliability and fixing things versus just shipping new features. Hook up the feedback loops so that the people who are getting alerted are the ones empowered and motivated to fix the problems that are paging them. Ideally, you should page the person who made the change, every time.
- Foster a culture of ownership and accountability while promulgating blamelessness throughout the org. Welcome engineers into production, and help them navigate production waters happily and successfully.

##### **Managing the portfolio of technical investments.**

- Operability is the longest term investment / primary source of technical debt, so no one is better positioned to help evaluate and amortize those risks than ops engineers. It is effectively free to write code, compared to the gargantuan resources it takes to run that code and tend to it over the years.
- Get excellent at [migrations](https://acloudguru.com/blog/business/what-is-cloud-migration). Leave no trailing, stale remnants of systems behind to support — those are a terrible drain on the team. Surface this energy drain to decision-makers instead of letting it silently fester.
- Hold the line against writing any more code than is absolutely necessary. Or adding any more tools than are necessary. Your line is, “what is the maintenance plan for this tool?”
- Educate and influence. Lobby for the primacy of operability. Take an interest in job ladders and leveling documents. No one should be promoted to senior engineering levels unless they write and support operable services.

This world is changing fast, and these changes are accelerating. Ops is everybody’s job now. Many engineers have no idea what this means, and have absorbed the lingering cultural artifacts of terror. It’s our job to fix the terror we ops folks instilled. We must find ways to reward curiosity, not punish it.

There’s never been a better time to [develop cloud skills](https://acloudguru.com/solutions/individuals) and level up your career.

_Charity Majors is the CTO at Honeycomb and a former product engineering manager at Facebook._

## Recommended

Get more insights, news, and assorted awesomeness around all things cloud learning.

[![kubernetes1](https://res.cloudinary.com/acloud-guru/image/fetch/c_thumb,f_auto,q_auto,w_465/https://acg-wordpress-content-production.s3.us-west-2.amazonaws.com/app/uploads/2021/03/k8s1.jpeg)](https://acloudguru.com/blog/engineering/whats-new-with-kubernetes-1-22)
