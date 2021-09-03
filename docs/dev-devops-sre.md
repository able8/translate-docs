# DevOps, SRE, and Platform Engineering

[August 1, 2021](https://iximiuz.com/en/posts/devops-sre-and-platform-engineering/#) From: https://iximiuz.com/en/posts/devops-sre-and-platform-engineering/

I'm a Software Engineer at heart, SRE at day job, and tech storyteller at night.

[Subscribe to my monthly newsletter](https://iximiuz.com/en/newsletter/) or [follow me on Twitter](https://iximiuz.com/en/newsletter/#twitter) for quality content on Containers, Kubernetes, Cloud Native stack, and Programming!

I compiled this thread on Twitter, and all of a sudden, it got quite  some attention. So here, I'll try to elaborate on the topic a bit more.  Maybe it would be helpful for someone trying to make a career decision  or just improve general understanding of the most hyped titles in the  industry.

DevOps, SRE, and Platform Engineering (thread)
Sharing my understanding of things after working in this domain for about two years.
Starting from the clearest one.
Dev - this is about application development, aka business logic. The only one that makes money for a company.

During my career, I used to work in teams and companies where as a  developer, I would push code to a repository and just hope that it would work well when some mythical system administrator would eventually take it to production. I also was in setups where I would need to provision bare-metal servers on Monday, figure out the deployment strategy on  Tuesday, write some business logic on Wednesday, roll it out myself on Thursday, and firefight a production incident on Friday. And all this  without even being aware of the existence of fancy titles like DevOps or SRE engineer.

But then people around me started talking DevOps and SRE, comparing them with each other, and compiling [awesome lists](https://github.com/dastergon/awesome-sre) of [resources](https://github.com/AcalephStorage/awesome-devops). New job opportunities began emerging, and I quickly jumped into the SRE train. So, below is my experience of being involved in all things SRE  and Platform Engineering from the former Software Developer standpoint.  And yeah, I think it's applicable primarily for companies where the  product is some sort of a web-facing service. This is the kind of  company I spent ten years working for. People doing embedded software or implementing databases probably live in totally different realities.

## What is Development

This one is the simplest to explain. Development - is about  application programming, i.e., writing the business logic of your main  product. This is the only activity among the three ones being discussed  here that directly makes money for the company.

> The only one that makes for a company is of course sales, everything else is expenditure :)

IMO, development is super hot! As a developer, you quickly start  thinking that you are the most important person around. Without your  code, there is nothing. But apparently, just writing code often isn't  enough. The code needs to be delivered to production and executed there.

I'd been carrying the Software Developer (or Software Engineer) title since the very beginning of my career in 2011. And I still remember the pain quite vividly - I always wished to have control over deploying my  code. And I rarely had it. Instead, there would be some obscure  procedure when someone, usually not even your senior colleague, would  have access to production servers and deploy the code there for you. So, if after pushing the changes to the repository, you got unlucky enough  to notice a bug only on the live version of your service, you'd need to  beg for an extra rollout. It most definitely sucked.

## What is DevOps

I'll not even try to quote the official definition here. Instead,  I'll share the first-hand experience. For me, DevOps was a cultural  shift giving development teams more control over shipping code to  production. The implementation could vary. I've been in setups where  developers would just have `sudo` on production servers. But probably the most common approach is to provide development teams with some sort of CI/CD pipelines.

In an ideal GitOps world, developers would still be just pushing code to repositories. However, there would be a magical button somewhere at  the team's disposal that would put the new version on live or maybe even provision a new piece of infrastructure to cover the new requirements.

The original idea of DevOps is probably much broader than just that.  But from what I see in the job descriptions, what I hear from recruiters trying to hunt me for a DevOps position, and what I managed to gather  from my fellow colleagues carrying the DevOps engineer title, most of  the time, it's about creating an efficient way to deploy stuff produced  by Development. In more advanced setups, DevOps may also be concerned  with other things improving the Development velocity. But DevOps itself  is never concerned with the actual application business logic.

## What is SRE

There is a [excellent series of books by Google](https://sre.google/books/) explaining the idea of the Site Reliability Engineering and, what's  even more important for me, sharing some real tech practices conducted  by Google SREs. In particular, it says that SRE is just one of the ways  to implement the DevOps culture - `class SRE implements DevOps {}`.

This explanation didn't really help me much. But what was even more  puzzling, subconsciously, I always felt excited while reading SRE job  descriptions and got bored quickly by the DevOps ones... So, there was  clearly a difference but, for a long time, I couldn't distill it.

Of course, that's just about my personal preferences, but whenever  someone mentions configuring a CI/CD pipeline, I always got depressed.  And the DevOps job descriptions nowadays are full of such  responsibilities. Don't get me wrong, CI/CD pipelines are amazing! I'm  always glad when I have a chance to use one. But setting them up isn't a thing I enjoy the most. On the contrary, when someone asks me to jump  in and take a look at a bleeding production, be it chasing a bug, a  memory leak, or performance degradation, I'm always more than just happy to help.

Developing code and shipping it to production still doesn't give you  the full picture. Someone needs to keep the production alive and  healthy! And that's how I see the place of SRE in my model of the world.

Google's SRE book focuses on monitoring and alerting, defining SLOs  of your services and tracking error budgets, incident response and  postmortems. These are the things one would need to apply to make the  production reliable. Facebook has a famous Production Engineer role, but it's pretty hard to distinguish it from a typical SRE role, judging  only by the job description.

Here is also a great tweet that kind of confirms my feeling that the primary focus of SRE is production.

My very simplified answer when someone says what is the difference between SRE and DevOps. 

* SRE = focused primarily on production 
* DevOps = focused primarily on CI/CD and developer velocity

And one more:

I like it! My typical answer is:
SRE works from Production backward. DevOps works from development forward. Somewhere in the middle, they meet.

So, DevOps keeps production fresh. SRE keeps production healthy.

## What is Platform Engineering

When I used to be the only engineer in a startup, a decent part of my job was to turn some generic resources I'd rent from the infrastructure provider into something more tailored for the company's needs. So, I  had a bunch of scripts to provision a new server, some understanding of  how to provide network connectivity between our servers in different  data centers, some skills to replicate the production setup on staging,  and maybe even write one or two daemons to help me with log collection. I didn't really understand it, but these things constituted our Platform.

Joining a much bigger company and starting consuming infra-related  resources brought me to a realization that there is a third area of  focus that might be quite close to DevOps and SRE. It's called Platform  Engineering.

From my understanding, Platform Engineering focuses on developing an  ecosystem that can be efficiently used from the Dev, Ops, and SRE  standpoints.

There might be quite some code writing in Platform Engineering. Or,  it could be mostly about configuring things. But again, it's not about  the primary business logic of your product - it's about making some  basic infrastructure more suitable for the day-to-day needs.
Platform Engineering - this is about infrastructure development.
PE focuses on creating a platform that can be efficiently used from the Dev, Ops, and SRE standpoints.
There is plenty of actual code writing in PE, but again, it's not about the primary business logic.

To be honest, I don't see a contradiction between my way of seeing  Platform Engineering and the explanation from this tweet. Development  needs infrastructure to run the code. So, if Platform Engineering is  about enabling others to do whatever they want to do, at least in part,  it should be concerned with infrastructure development.

I have a feeling that in a bigger setup, when a company would have  thousands of bare-metal servers in its own data centers, a Platform  Engineering would start from managing this fleet of machines. So, some  sort of inventory software might need to be installed or even developed  internally. Installing operating systems and basic packages on the  servers being provisioned would probably also fall into the Platform  Engineering responsibility.

Luckily, clouds made Platform Engineering operating on much higher  layers. All the basic fleet management tasks are already solved for you. And even orchestration of your workloads is solved by projects like  Kubernetes or AWS ECS. However, the solution is quite generic, while  your teams are likely to deploy pretty similar microservices. So,  providing them with a default project template that would be integrated  with the company's metrics and logs collection subsystems would make  things moving much faster.

## What's about titles?

So far, I was deliberately avoiding talking about roles and titles.  Development, Operations, SRE, and Platform Engineering for me are about  areas of focus. And to a much lesser extent about titles. One person can be a Dev this week, then an Ops on the next week, and an SRE on the  week after.

From my experience, the separation between Dev, Ops, SRE, and PE  becomes more apparent when the company size gets bigger. A bigger  company size usually means more specialists and fewer generalists.  That's how you end up with dedicated SRE teams and a Platform  Engineering department. But of course, it's not a strict rule. For  instance, with my SRE title, I spent like a year doing all things true  SRE (SLO, monitoring, alerting, incident response) and then transitioned into Platform Engineering, where I do more infra development than  traditional SRE. YMMV.

## Where Security goes?

> Awesome , but where the security team gets involved from DevOps and SRE prospective.

That's a very good question! But I don't have a simple answer. For  me, a reasonable approach is to make security a cross-cutting theme in  all Dev, Ops, SRE, and PE. Different security concerns can be addressed  on different layers using different tools. For instance, Development  could be concerned with preventing SQL injections while Platform folks  could harden the networking by configuring some fancy cilium policies.

## Instead of conclusion

Don't forget, all the things above are IMO 😉