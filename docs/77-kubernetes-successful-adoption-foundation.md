# Kubernetes Is Not Your Platform, It's Just the Foundation

Mar 19, 2021

### Key Takeaways

- Kubernetes itself is not a platform but only the foundational element of an ecosystem not only of tools, and services, but also offering support as part of a compelling internal product.
- Platform teams should provide useful abstractions of Kubernetes complexities to reduce cognitive load on stream teams.
- The needs of platform users change and a platform team needs to ease their journey forward.
- Changes to the platform should meet the needs of stream teams, prompting a collaborative discovery phase with the platform team followed by stabilizing the new features (or service) before they can be consumed by other stream teams in a self-service fashion.
- A team-focused Kubernetes adoption requires an assessment of cognitive load and tradeoffs, clear platform and service definitions, and defined team interactions.

I read a lot of articles and see presentations on impressive Kubernetes tools, automation, and different ways to use the technology but these often offer little context about the organization other than the name of the teams involved.

We can’t really judge the success of technology adoption, particularly Kubernetes, if we don't know who asked for it, who needs it, and who is implementing it in which ways. We need to know how teams are buying into this new offering — or not.

In [Team Topologies](https://teamtopologies.com/book), which I wrote with [Matthew Skelton](https://www.infoq.com/profile/Matthew-Skelton/), we talk about fundamental types of teams, their expected behaviors and purposes, and perhaps more importantly about the interactions among teams. Because that's what's going to drive adoption of technology. If we don't take that into consideration, we might fall into a trap and build technologies or services that no one needs or which have no clear use case.

#### Related Sponsored Content

- ##### [Kubernetes Up & Running - Download the eBook (By O'Reilly)](http://www.infoq.com/infoq/url.action?i=e1b2b077-e265-40a5-8f12-593f0f6f6baa&t=f)

#### Related Sponsor

[![](https://res.infoq.com/sponsorship/topic/22bdfa2a-7b5c-4be5-b60c-93141547ff5a/VMWareTanzuRSBLogo-1589361540634-1612367286968.png)](http://www.infoq.com/infoq/url.action?i=9e6eac17-4b86-410d-8cf9-9ce2e2a28863&t=f)

### [**Free your apps. Simplify your ops.**](http://www.infoq.com/infoq/url.action?i=4e465f12-aac3-4817-8b96-89c289f60a82&t=f)

Team cognitive load has a direct effect on platform success and team interaction patterns can help reduce unnecessary load on product teams. We’ll look at how platforms can minimize cognitive load and end with ideas on how to take a team-centric approach to Kubernetes adoption.

## Is Kubernetes a platform?

In 2019, Melanie Cebula, an infrastructure engineer at Airbnb, gave an excellent talk at Qcon. As DevOps lead editor for InfoQ, I wrote [the story about it](https://www.infoq.com/news/2019/03/airbnb-kubernetes-workflow/). In the first week online, this story had more than 23,000 page views. It was my first proper viral story, and I tried to understand why this one got so much attention. Yes, Airbnb helped, and Kubernetes is all the rage, but I think the main factor was that the story was about simplifying Kubernetes adoption at a large scale, for thousands of engineers.

The point is that many developers and many engineers find it complicated and hard to adopt Kubernetes and change the way they work with the APIs and the artifacts they need to produce to use it effectively.

The term “platform” has been overloaded with a lot of different meanings. Kubernetes is a platform in the sense that it helps us deal with the complexity of operating microservices. It helps provide better abstractions for deploying and running our services. That's all great, but there's a lot more going on. We need to think about how to size our hosts, how and when to create/destroy clusters, and how to update to new Kubernetes versions and who that will impact. We need to decide on how to isolate different environments or applications, with namespaces, clusters, or whatever it might be. Anyone who has worked with Kubernetes can add to this list: perhaps worrying about security, for example. A lot of things need to happen before we can use Kubernetes as a platform.

One problem is that the boundaries between roles are often unclear in an organization. Who is the provider? Who is the owner responsible for doing all of this? Who is the consumer? What are the teams that will consume the platform? With blurry boundaries, it's complicated to understand who is responsible for what and how our decisions will affect other teams.

[Evan Bottcher](https://www.linkedin.com/in/evanbottcher/) defines a digital platform as “a foundation of self-service APIs, tools, and services, but also knowledge and support, and everything arranged as a compelling internal product.” We know that self-service tools and APIs are important. They're critical in allowing a lot of teams to be more independent and to work autonomously. But I want to bring attention to his mention of “knowledge and support”. That implies that we have teams running the platform that are focused on helping product teams understand and adopt it, besides providing support when problems arise in the platform.

Another key aspect is to understand the platform as “a compelling internal product”, not a mandatory platform with shared services that we impose on everyone else. As an industry, we've been doing that for a long time, and it simply doesn't work very well. It often creates more pain than the benefits it provides for the teams forced to use a platform that is supposed to be a silver bullet. And we know silver bullets don’t exist.

We have to think about the platform as a product. It's meant for our internal teams, but it's still a product. We're going to see what that implies in practice.

The key idea is that Kubernetes by itself is not a platform. It's a foundation. Yes, it provides all this great functionality — autoscaling, self-healing, service discovery, you name it — but a good product is more than just a set of features. We need to think about how easy it is to adopt, its reliability, and the support for the platform.

A good platform, as Bottcher says, should create a path of least resistance. Essentially, the right thing to do should be the easiest thing to do with the platform. We can’t just say that whatever Kubernetes does is the right thing. The right thing depends on your context, on the teams that are going to consume the platform, which services they need the most, what kind of help they need to onboard, and so on.

One of the hard things about platforms is that the needs of the internal teams are going to change, with respect to old and new customers. Teams that are consuming the platform are probably going to have more specific requests and requirements over time. At the same time, the platform must remain understandable and usable for new teams or new engineers adopting it. The needs and the technology ecosystem keep evolving.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/54image001-1615999814012.jpg)

**Figure 1: Avoid creating isolated teams when bringing Kubernetes into the organization.**

## Team cognitive load

Adopting Kubernetes is not a small change. It's more like an elephant charging into a room. We have to adopt this technology in a way that does not cause more pain than the benefits it brings. We don't want to end up with something like Figure 1, that resembles life before the DevOps movement, with isolated groups. It’s the Kubernetes team now rather than the operations team, but if we create an isolationist approach where one group makes decisions independent from the other group, we’re going to run into the same kinds of problems.

If we make platform decisions without considering the impact on consumers, we're going to increase pain in the form of their cognitive load, the amount of effort our teams must put out to understand and use the platform.

[John Sweller](https://en.wikipedia.org/wiki/John_Sweller) formulated cognitive load as “the total amount of mental effort being used in the working memory”. We can break this down into three types of cognitive load: intrinsic, extraneous, and germane.

We can map these types of cognitive load in software delivery: intrinsic cognitive load is the skills I need to do my work; extraneous load is the mechanics, things that need to happen to deliver the value; and germane is the domain focus.

Intrinsic cognitive load, if I’m a Java developer, is knowing how to write classes in Java. If I don't know how to do that, this takes effort. I have to Google it or I have to try to remember it. But we know how to minimize intrinsic cognitive load. We can have classical training, pair programming, mentoring, code reviews, and all techniques that help people improve their skills.

Extraneous cognitive load includes any task needed to deliver the work I'm doing to the customers or to production. Having to remember how to deploy this application, how to access a staging environment, or how to clean up test data are all things that are not directly related to the problem to solve but are things that I need to get done. Team topologies and platform teams in particular can minimize extraneous cognitive load. This is what we're going to explore throughout the rest of this article.

Finally, germane cognitive load is knowledge in my business domain or problem space. For example, if I'm working in private banking, I need to know how bank transfers work. The point of minimizing extraneous cognitive load is to free up as much memory available for focus on the germane cognitive load.

Jo Pearce’s “Hacking Your Head” articles and [presentations](https://www.slideshare.net/JoPearce5/hacking-your-head-managing-information-overload-extended) go deeper into this.

The general principle is to be mindful of the impact of your platform choices on your teams’ cognitive loads.

## Case studies

I mentioned Airbnb engineer Melanie Cebula’s excellent talk at Qcon so let’s look at how that company reduced the cognitive load of their development teams.

“The best part of my day is when I update 10 different YAML files to deploy a one-line code change,” said no one, ever. Airbnb teams were feeling this sort of pain as they embarked on their Kubernetes journey. To reduce this cognitive load, they created a simple command-line tool, kube-gen, which allows the application teams or service teams to focus on a smaller set of configurations and details, which are specific to their own project or services. A team needs to configure only those settings like files or volumes specifically related to the germane aspect of their application. The kube-gen tool then generates the boilerplate code configuration for each environment, in their case: production, canary, and development environments. This makes it much easier for development teams to focus on the germane parts of their work.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/23image002-1615999815140.jpg)

**Figure 2: Airbnb uses kube-gen to simplify their Kubernetes ecosystem.**

As Airbnb did, we essentially want to clarify the boundaries of the services provided by the platform and provide good abstractions so we reduce the cognitive load on each service team.

In Team Topologies, we talk about four types of teams, shown in Figure 3. Stream-aligned teams provide the end-customer value, they’re the heartbeat that delivers business value. The three other types of teams provide support and help reduce cognitive load. The platform team shields the details of the lower-level services that these teams need to use for deployment, monitoring, CI/CD, and other lifecycle supporting services.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/44image003-1615999812928.jpg)

**Figure 3: Four types of teams.**

The stream-aligned team resembles what organizations variously call a product team, DevOps team, or build-and-run team — these teams have end-to-end ownership of the services that they deliver. They have runtime ownership, and they can take feedback from monitoring or live customer usage for improving the next iteration of their service or application. We call this a “stream-aligned team” for two reasons. First, “product” is an overloaded term; as our systems become more and more complex, the less precise it is to define a standalone product. Second, we want to acknowledge the different types of streams beyond just the business-value streams; it can be compliance, specific user personas, or whatever makes sense for aligning a team with the value it provides.

Another case study is uSwitch, which helps users in the UK compare utility providers and home services and makes it easy to switch between them.

A couple of years ago, [Paul Ingles](https://www.linkedin.com/in/pingles/), head of engineering at uSwitch, wrote “ [Convergence to Kubernetes](https://pingles.medium.com/convergence-to-kubernetes-137ffa7ea2bc)”, which brought together the adoption of Kubernetes technology, how it helps or hurts teams and their work, and data for meaningful analysis. Figure 4, from that article, measures all the different teams’ low-level AWS service calls at uSwitch.

When uSwitch started, every team was responsible for their own service, and they were as autonomous as possible. Teams were responsible for creating their own AWS accounts, security groups, networking, etc. uSwitch noticed that the number of calls to these services was increasing, correlated with a feeling that teams were getting slower at delivering new features and value for the business.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/24image004-1615999812514.jpg)

**Figure 4: Use of low-level AWS services over time.**

What uSwitch wanted to do by adopting Kubernetes was not only to bring in the technology but also to change the organizational structure and introduce an infrastructure platform team. This would address the increasing cognitive load generated by teams having to understand these different AWS services at a low level.

That was a powerful idea. Once uSwitch introduced the platform, traffic directly through AWS decreased. The curves of calls in Figure 5 is a proxy for the cognitive load on the application teams. This concept aligns with a platform team’s purpose: to enable stream-aligned teams to work more autonomously with self-service capabilities and reduced extraneous cognitive load. This is a very different conceptual starting point from saying, "Well, we're going to put all shared services in a platform."

Ingles also wrote that  they wanted to keep the principles they had in place before around team autonomy and teams working with minimal coordination, by providing a self-service infrastructure platform.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/39image005-1615999814289.jpg)

**Figure 5: The number of low-level AWS service calls dropped after uSwitch introduced their Kubernetes-based platform.**

## Treat the platform as a product

We talk about treating the platform as a product — an internal product, but still a product. That means that we should think about its reliability, its fit for purpose, and the developer experience (DevEx) while using it.

First, for the platform to be reliable, we need to have on-call support, because now the platform sits in the path of production. If our platform’s monitoring services fail or we run out of storage space for logs, for example, the customer-facing teams are going to suffer. They need someone who provides support, tells them what's going on, and estimates the expected time for a fix. It should also be easy to understand the status of the platform services and we should have clear, established communication channels between the platform and the stream teams in order to reduce cognitive load in communications. Finally, there needs to be coordination with potentially affected teams for any planned downtime.

Secondly, having a platform that's fit for purpose means that we use techniques like prototyping and we get regular feedback from the internal customers. We use iterative practices like agile, pair programming, or TDD for faster delivery with higher quality. Importantly, we should focus on having fewer services of higher quality and availability rather than trying to build every service that we can imagine to be potentially useful. We need to focus on what teams really need and make sure those services are of high quality. This means we need very good product management to understand priorities, to establish a clear but flexible roadmap, and so on.

Finally, development teams and the platform should speak the same language in order to maximize the experience and usability. The platform should provide the services in a straightforward way. Sometimes, we might need to compromise. If development teams are not familiar with YAML, we might think of the low cost, low effort, and long-term gain of training all development teams in YAML. But these decisions are not always straightforward and we should never make them without considering the impact on the development teams or the consuming teams. We should provide the right levels of abstraction for our teams today, but the context may change in the future. We might adopt better or higher levels of abstraction, but we always should look at what makes sense given the current maturity and engineering practices of our teams.

Kubernetes helped uSwitch establish these more application-focused abstractions for things like services, deployments, and ingress rather than the lower-level service abstractions that they were using before with AWS. It also helped them minimize coordination, which was another of the key principles.

I spoke with Ingles and with [Tom Booth](https://www.linkedin.com/in/thbooth/), at the time infrastructure lead at uSwitch, about what they did and how they did it.

Some of the things that the platform team helped the service teams with were providing dynamic database credentials and multi-cluster load balancing. They also made it easier for service teams to get alerts for their customer-facing services, define service-level objectives (SLOs), and make all that more visible. The platform make it easy for teams to configure and monitor their SLOs with dashboards — if an indicator drops below a threshold, notifications go out in Slack, as shown in Figure 6.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/12image006-1615999813749.jpg)

**Figure 6: An example of a SLO threshold notification in Slack at uSwitch.**

Teams found it easy to adopt these new practices. uSwitch teams are familiar with YAML and can quickly configure these services and benefit from them.

## Achievements beyond the technical aspects

uSwitch’s journey is fascinating beyond the technical achievements. They started this infrastructure adoption in 2018 with only a few services. They identified their first customer to be one team that was struggling without any centralized logging, metrics, or autoscaling. They recognized that growing services around these aspects would be a successful beginning.

In time, the platform team started to define their own SLAs and SLOs for the Kubernetes platform, serving as an example for the rest of the company and highlighting the improvements in performance, latency, and reliability. Other teams could observe and make an informed decision to adopt the platform services or not. Remember, it was never mandated but always optional.

Booth told me that uSwitch saw traffic increasing through the Kubernetes platform versus what was going directly through AWS and this gave them some idea of how much adoption was taking place. Later, the team addressed some critical cross-functional gaps in security, GDPR data privacy and handling of alerts and SLOs.

One team, one of the more advanced in both engineering terms and revenue generation, was already doing everything that the Kubernetes platform could provide. They had no significant motivation to adopt the platform — until they were sure that it provided the same functionality with the increased levels of reliability, performance, and so on. It no longer made sense for them to take care of all these infrastructure aspects on their own. The team switched to use the Kubernetes platform and increased its capacity to focus on the business aspects of the service. That was the “ultimate” prize for the platform team, to gain the adoption from the most advanced engineering team in the organization.

## Four Key Metrics

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/28image007-1615999814841.jpg)

**Figure 7: Four key metrics from the Accelerate book.**

Having metrics can be quite useful. As the platform should be considered a product, we can look at product metrics — and the categories in Figure 7 come from the book Accelerate by Nicole Forsgren, Jez Humble, and Gene Kim. They write that high-performing teams do very well at these key metrics around lead time, deployment frequency, MTTR, and change failure rate. We can use this to help guide our own platform services delivery and operations.

Besides the Accelerate metrics, user satisfaction is another useful and important measurement. If we're creating a product for users, we want to make sure it helps them do their job, that they're happy with it, and that they recommend it to others. There's a simple example from Twilio. Every quarter or so, their platform team surveys the engineering teams with some questions (Figure 8) on how well the platform helps them build, deliver, and run their service and how compelling it is to use.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/11image008-1615999813204.jpg)

**Figure 8: The survey that Twiliio’s platform team sends out to engineering teams.**

With this simple questionnaire, you can look at overall user satisfaction over time and see trends. It’s not just about the general level of satisfaction with the technical services, but also with the support from the platform team. Dissatisfaction may arise because the platform team was too busy to listen to feedback for a period of time, for example. It's not just about the technology. It's also about interactions among teams.

Yet another important area to measure is around platform adoption and engagement. In the end, for the platform to be successful, it must be adopted. That means it’s serving its purpose. At the most basic level, we can look at how many teams in the organization are using the platform versus how many teams are not. We can also look at adoption per platform service, or even adoption of a particular service functionality. That will help understand the success of a service or feature. If we have a service that we expected to be easily adopted but many teams are lagging behind, we can look for what may have caused that.

Finally, measuring the reliability of the platform itself, as uSwitch did, is important as well. They had their own SLOs for the platform and this was available to all teams. Making sure we provide that information is quite important.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/21image009-1615999815426.jpg)

**Figure 9: Useful metrics for your platform as a product.**

Figure 9 shows some examples of the metrics categories. Each organization with its own context might have different specific metrics, but the types and categories that we should be looking at should be more or less the same.

## Team interactions

The success of the platform team is the success of the stream-aligned teams. These two things go together. It's the same for other types of supporting teams. Team interactions are critical because the definition of success is no longer just about making technology available, it’s about helping consumers of that technology get the expected benefits in terms of speed, quality, and operability.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/8image010-1615999815955.jpg)

**Figure 10: Airbnb’s platform team provides an abstraction of the underlying Kubernetes architecture.**

Airbnb effectively had a platform team (although internally called Infrastructure team) to abstract a lot of the underlying details of the Kubernetes platform, as shown in Figure 10. This clarifies the platform’s boundaries for the development teams and exposes those teams to a much smaller cognitive load than telling them to use Kubernetes and to read the official documentation to understand how it works. That's a huge task that requires a lot of effort. This platform-team approach reduces cognitive load by providing more tightly focused services that meet our teams’ needs.

To accomplish this, we need adequate behaviors and interactions between teams. When we start a new platform service or change an existing one, then we expect strong collaboration between the platform team and the first stream-aligned teams that will use the new/changed service. Right from the beginning of the discovery period, a platform team should understand a streaming team’s needs, looking for the simplest solutions and interfaces that meet those needs. Once the stream-aligned team starta using the service, then the platform team’s focus should move on to supporting the service and providing up-to-date, easy to follow documentation for onboarding new users. The teams no longer have to collaborate as much and the platform team is focused on providing a good service. We call this interaction mode X-as-a-Service.

Note that this doesn't mean that the platform hides everything and the development teams are not allowed to understand what's going on behind the scenes. That's not the point. Everyone knows that it's a Kubernetes-based platform and we should not forbid teams from offering feedback or suggesting new tools or methods. We should actually promote that kind of engagement and discussion between stream-aligned teams and platform teams.

For example, troubleshooting services in Kubernetes can be quite complicated. Figure 11 shows only the top half of a flow chart for diagnosing a deployment issue in Kubernetes. This is not something we want our engineering teams to have to go through every time there's a problem. Neither did Airbnb.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/18image011-1615999813488.jpg)

**Figure 11: Half of a Kubernetes deployment troubleshooting flow chart.**

Airbnb initiated a discovery period during which teams closely collaborated to understand what kind of information they needed in order to diagnose a problem and what kinds of problems regularly occurred. This led to an agreement on what a troubleshooting service should look like. Eventually, the service became clear and stable enough to be consumed by all the stream-aligned teams.

Airbnb already provided these two services in their platform: kube-gen and kube-deploy. The new troubleshooting service, kube-diagnose, collected all the relevant logs, as well as all sorts of status checks and other useful data to simplify the lives of their development teams. Diagnosing got a lot easier, with teams focused on potential causes for problems rather than remembering where all the data was or which steps to get them.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/7image012-1615999814550.jpg)

**Figure 12: Services provided by the infrastructure (platform) team at Airbnb.**

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/17image013-1615999815692.jpg)

**Figure 13: The cloud-native landscape.**

Figure 13 shows just how broad the cloud-native landscape is. We don’t want stream-aligned teams to have to deal with that on their own. Part of the role of a platform team is to follow the technology lifecycle. We know how important that is. Let’s imagine there's a [CNCF](https://www.cncf.io/) tool that just graduated which would help reduce the amount of custom code (or get rid of an older, less reliable tool) in our internal platform today. If we can adopt this new tool in a transparent fashion that doesn’t leak into the platform service usage by stream-aligned teams, then we have an easier path forward in terms of keeping up with this ever-evolving landscape.

If, on the contrary, a new technology we’d like to adopt in the platform implies a change in one or more service interfaces that means we need to consult the stream-aligned teams to understand the effort required of them and evaluate the trade-offs at play. Are we getting more benefit in the long run than the pain of migration/adaptation today?

In any case, the internal platform approach helps us make visible the evolution of the technology inside our platform.

![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/7image014-1615999816275.jpg)

**Figure 14: uSwitch open-sourced Heimdall tool that their platform team created for internal use initially.**

The same goes for adopting open source solutions. For example, uSwitch open sourced the Heimdall application that provides the chat tool integration with the SLOs dashboard service. And there are many more open-source tools. Zalando, for example, has really cool stuff around cluster lifecycle management. The point being that if a piece of open-source technology makes sense for us, we can adopt it more easily if they sit under a certain level of abstraction in the platform. And if it is not a transparent change, we can always collaborate with stream-aligned teams to identify what would need to change in terms of their usage of the affected service(s).

## Starting with team-centric approach for Kubernetes adoption

There are three keys to a team-focused approach to Kubernetes adoption.

We start by assessing the cognitive load on our development teams or stream-aligned teams. Let’s determine how easy it is for them to understand the current platform based on Kubernetes. What abstractions do they need to know about? Is this easy for them or are they struggling? It's not an exact science, but by asking these questions we can get a feel for what problems they're facing and need help with. The Airbnb case study is important because it made it clear that a tool-based approach to Kubernetes adoption brings with it difficulties and anxiety if people have to use this all-new platform without proper support, DevEx, or collaboration to understand their real needs.

Next, we should clarify our platform. This often is simple, but we don't always do it. List exactly all services we have in the platform, who is responsible for each, and all the other aspects of a digital platform that I’ve mentioned like responsibility for on-call support, communication mechanisms, etc. All this should be clear. We can start immediately by looking at the gaps between our actual Kubernetes implementation and an ideal digital platform, and addressing those.

Finally, clarify the team interactions. Be more intentional about when we should collaborate and when should we expect to consume a service independently. Determine how we develop new services and who needs to be involved and for how long. We shouldn't  just say that we're going to collaborate and leave it as an open-ended interaction. Establish an expected duration for the collaboration, for example two weeks to understand what teams need from a new platform service and how the service interface should look like, before we’re actually building out such service. Do the necessary discovery first, then focus on functionality and reliability. At some point, the service can become “generally available” and interactions evolve to “X as a service”.

There are many good platform examples to look at from companies like  [Zalando](https://www.slideshare.net/try_except_/kubernetes-at-zalando-cncf-end-user-committee-presentation), [Twilio](https://www.infoq.com/presentations/twilio-devops/), [Adidas](https://youtu.be/XwaRKcjkAAo), [Mercedes](https://speakerdeck.com/devopslx/2019-dot-02-meetup-talk-devops-adoption-at-mercedes-benz-dot-io), etc. The common thread among them is a digital platform approach consisting not only of technical services but good support, on-call, high quality documentation, and all these things that make the platform easy for their teams to use and accelerate their capacity to deliver and operate their software more autonomously. I also wrote [an article for TechBeacon](https://techbeacon.com/enterprise-it/why-teams-fail-kubernetes-what-do-about-it) that goes a bit deeper into these ideas.

## About the Author

**![](https://res.infoq.com/articles/kubernetes-successful-adoption-foundation/en/resources/4manuel-pais-1615999143952.jpg)Manuel Pais** is an independent IT organizational consultant and trainer, focused on team interactions, delivery practices, and accelerating flow. He is co-author of the book [Team Topologies: Organizing Business and Technology Teams for Fast Flow](https://teamtopologies.com/book). He helps organizations rethink their approach to software delivery, operations, and support via strategic assessments, practical workshops, and coaching.
