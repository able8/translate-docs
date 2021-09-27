# Cloud-Native Is about Culture, Not Containers

LikePrint [Bookmarks](http://www.infoq.com/showbookmarks.action)

Mar 17, 2021

### Key Takeaways

- It is possible to be very cloud-native without a single microservice
- Before embarking on a cloud-native transformation, it’s important to be clear on what cloud-native means to your team and what the true problem being solved is
- The benefits of a microservices architecture will not be realised if releases involve heavy ceremony, are infrequent, and if all microservices have to be released at the same time
- Continuous integration and deployment is something you do, not a tool you buy
- Excessive governance chokes the speed out of cloud, but if you don’t pay enough attention to what is being consumed, there can be serious waste

At the QCon London last year, I provided a Cloud-Native session about Culture, not Containers. What had started me thinking about the role of culture in cloud native was a great [InfoQ article from Bilgin Ibryam](https://www.infoq.com/articles/microservices-post-kubernetes/).  One of the things Bilgin did was define a cloud native architecture as lots of microservices, connected by smart pipes. I looked at that and thought it looked totally unlike applications I wrote, even though I thought I was writing cloud native applications. I'm part of the IBM Garage, helping clients get cloud native,, and yet I rarely used microservices in my apps. The apps I create mostly  looked nothing like Bilgin’s diagram. Does that mean I'm doing it wrong, or is maybe the definition of cloud native a bit complicated?

I don’t want to single Bilgin out, since Bilgin’s article was called "Microservices in the post-Kubernetes Era," so it would be a bit ridiculous if he weren't talking about microservices a lot in that article. It’s also the case that almost all definitions of cloud native equate it to microservices. Everywhere I looked, I kept seeing the assumption that microservices equals native and Cloud-native equals microservices. Even the Cloud Native Computing Foundation used to define cloud native as all about microservices, and all about containers, with a bit of dynamic orchestration in there. Saying cloud native doesn’t always involve microservices, which puts me in this peculiar position because not only am I saying Bilgin is wrong, I'm saying the Cloud Native Computing Foundation is wrong - what did they ever know about Cloud-native? I'm sure I know way more than them, right?

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/10figure-1-1615824705148.jpg)

#### Related Sponsored Content

#### Related Sponsor

**[![](https://assets.infoq.com/resources/en/ScrumRSBCoverScrum.png)](http://www.infoq.com/infoq/url.action?i=5e855cd4-7e93-4c47-844e-63e238096d43&t=f)**

**Measure Value to Enable Improvement and Agility with Evidence-Based Management from Scrum.org - [Download the EBM Guide](http://www.infoq.com/infoq/url.action?i=4a892cb7-427b-4b9e-bb91-0a6d79444a4a&t=f).**

Well, obviously I don't. I'm on the wrong side of history on this one. I will admit that. (Although I’m on the wrong side of history, I notice that the CNCF have updated their definition of cloud native and, while microservices and containers are still there, they don’t seem quite as mandatory as they used to be, so a little bit of history might be on my side!). Right or wrong, I'm still going to die on my little hill; that Cloud-native is about something much bigger than microservices. Microservices are one way of doing it. They're not the only way of doing it.

In fact, you do see a range of definitions within our community. If you ask a bunch of people what Cloud-native means, some people will say “born on the cloud”. This was very much the original definition of Cloud-native back before microservices were even a thing. Some people will say it's microservices.

Some people will say, “oh no, it's not just microservices, it's microservices on Kubernetes, and that's how you get Cloud-native”. This one I don’t like, because to me, Cloud-native shouldn't be about a technology choice. Sometimes I see Cloud-native used as a synonym for DevOps, because a lot of the cloud native principles and practices are similar to what devops teaches.

Sometimes I see Cloud-native used just as a way of saying “we’re developing modern software”: “We're going to use best practices; it's going to be observable; it's going to be robust; we’re going to release often and automate everything; in short, we're going to take everything we've learned over the last 20 years and develop software that way, and that's what makes it Cloud-native”. In this definition, cloud is just a given - of course it’s on cloud, because we’re developing this in 2021.

Sometimes I see Cloud-native used just to mean Cloud. We got so used to hearing Cloud-native that every time we talk about Cloud, we just feel like we have to tack a ‘-native’ on afterwards, but we're really just talking about Cloud. Finally, when people say Cloud-native, sometimes what they mean is idempotent. The problem with this is if you say Cloud-native means idempotent, everybody else goes, "What? What we really mean by “idempotent” is rerunnable? If I take it, shut it down, and then start it up again, there’s no harm done. That's a fundamental requirement for services on the Cloud.

With all of these different definitions, is it any wonder we're not entirely sure what we're trying to do when we do Cloud-native?

## Why?

“What are we actually trying to achieve?” is an incredibly important question. When we're thinking about technology choices and technology styles, we want to be stepping back just from “I'm doing Cloud-native because that's what everybody else is doing” to thinking “what problem am I actually trying to solve?” To be fair to the CNCF, they had this “why” right on the front of their definition of Cloud-native. They said, "Cloud-native is about using microservices to build great products faster." We're not just using microservices because we want to; we're using microservices because they help us build great products faster.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/5figure-2-1615824705387.jpg)

This is where we step back to make sure we understand the problem we’re solving. Why couldn't we build great products faster before? It’s easy to skip this step, and I think all of us are guilty of this sometimes. Sometimes the problem that we're actually trying to solve is that everybody else is doing it, so we fear missing out unless we start doing it. Once we put it like that, FOMO isn’t a great decision criteria. Even worse, “my CV looks dull” definitely isn't the right reason to choose a technology.

## Why Cloud?

 I think to get to why we should be doing things in a Cloud-native way; we want to step back and say, "Why were we even doing things on the Cloud?" Here are the reasons:

- **Cost**: Back when we first started putting things on the Cloud, price was the primary motivator. We said, "I've got this data center, I have to pay for the electricity, I have to pay people to maintain it. And I have to buy all the hardware. Why would I do that when I could use someone else's data center?" What creates a cost-saving between your own data center and someone else's data center is that your own data center has to stock up enough hardware for the maximum demand. That’s potentially a lot of capacity which is unused most of the time. If it's someone else's data center, you can pool resources. When demand is low, you won’t pay for the extra capacity.
- **Elasticity**: The reason Cloud saves you money is because of that elasticity. You can scale up; you can scale down. Of course, that's old news now. We all take elasticity for granted.
- **Speed**: The reason we're interested in Cloud now is because of the speed. Not necessarily the speed of the hardware, although some cloud hardware can be dazzlingly fast. The cloud is an excellent way to use GPUs, and it’s more or less the only way to use quantum computers. More generally, though, we can get something to market way, way faster via the cloud than we could when we had to print software onto CD-Roms and mail them out to people, or even when we had to stand instances up in our own data center.

## 12 Factors

Cost savings, elasticity, and delivery speed are great, but we get all of that  just by being on the Cloud. Why do we need Cloud-native? The reason we need Cloud-native is that a lot of companies found they tried to go to the Cloud and they got electrocuted.

It turns out things need to be written differently and managed differently on the cloud. Articulating these differences led to the 12 factors. The 12 factors were a set of mandates for how you should write your Cloud application so that you didn't get electrocuted.

You could say the 12 factors described how to write a cloud native application - but the 12 factors had absolutely nothing to do with microservices. They were all about how you managed your state. They were about how you managed your logs. The 12 factors helped applications become idempotent, but “the 12 factors” is catchier than “the idempotency factors”.

The 12 factors were published two years before Docker got to market. Docker containers revolutionised how the cloud was used. Containers are so good, it’s hard to overstate their importance. They solve many problems and create new architectural possibilities. Because containers are so it’s easy, it’s possible to distribute an application across many containers. Some companies are running single applications across 100, 200, 300, 400, or 500 distinct containers. Compared to that kind of engineering prowess, an application which is spread across a mere six containers seems a bit inadequate. In the face of so little complexity, it’s easy to think  “I must be doing it really wrong. I'm not as good a developer as them over there”.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/4figure-3-1615824705625.jpg)

Actually, no. It's not a competition to see how many containers you can have. Containers are great, but the number of containers you have should be tuned to your needs.

## Speed

Let’s try and remember - what were your needs again? When we think about Cloud,  we usually want to be thinking about that speed. The reason we want lots of containers is that we want to get new things to market faster. If we have lots of containers and we're either shipping the exact same things to market or we're getting to market at the same speed, then all of a sudden, those containers are only a cost. They're not helping us, and we’re burning cycles managing the complexity that comes with scattering an application in tiny pieces all over the infrastructure. If we have this amazing architecture that allows us to respond to the market but we’re not responding, then that's a waste. If we have this architecture, that means we can go fast, but we're not going fast, then that's a waste as well.

## How to fail at Cloud-Native

Which brings me to how to fail at Cloud-native. For context, I'm a consultant. I'm a full-stack developer in the IBM Garage. We work with startups and with large companies, helping them get to the cloud and get the most out of cloud. As part of that, we help them solve interesting, tough, problems, and we help them do software faster than they’ve been able to to do it before. To make sure we’re really getting the most out of the cloud, we do a lean startup, extreme programming, design thinking,  DevOps; and cloud native. Because I’m a consultant, I see a lot of customers who are on the journey to the Cloud. Sometimes that goes well, and sometimes there are these pitfalls. Here  are some of the traps that I've seen smart clients fall into. So, What is Cloud-Native?

One of the earliest traps is the magic morphing meaning. If I say Cloud-native and I mean one thing and you say Cloud-native and mean another thing, we’re going to have a problem communicating...

Sometimes that doesn't really matter, but sometimes it makes a big difference. If one person thinks the goal is microservices and then the other person feels the goal is to have an idempotent system, uh oh. Or if part of an organisation wants to go to the Cloud because they think it's going to allow them to get to market faster, but another part is only going to the Cloud to deliver the exact same speed as before, but more cost-effectively, then we might have some conflict down the road.

## Microservices Envy

Often one of the things that drives some of this confusion about goals is because we have a natural tendency to look at other people, doing fantastic things, and want to emulate them.  We want to do those fantastic things ourselves without really thinking about our context and whether they're appropriate for us. One of our IBM Fellows has a heuristic when he goes in to talk to a client about microservices. He says, "If they start talking about Netflix and they just keep talking about Netflix, and they never mention coherence, and they never mention coupling, then probably they're not really doing it for the right reasons."

Sometimes we talk to clients, and they say, "Right, I want to modernize to microservices." Well, microservices are not a goal. No customer will look at your website and say, "Oh, microservices. That's nice." Customers are going to look at your website and judge it on whether it serves their needs, whether it’s easy and delightful, and, all of these other things. Microservices can be an excellent means to that end, but they're not a goal in themselves. I should also say: microservices are a means. They're not necessarily the only means to that goal.

A colleague of mine in the IBM Garage had some conversations with a bank in Asia-Pacific. The bank was having problems responding to their customers, because their software was all old and heavy and calcified. They were also having a people problem because all of their COBOL developers were old and leaving the workforce. So the bank  knew they had to modernise. The main driver in this case wasn't the aging workforce, it was really competitiveness and agility. They were getting beaten by their competitors because they had this big estate of COBOL code and every change was expensive and slow. They said, "Well, to solve this problem, we need to get rid of all of our COBOL, and we need to switch to a modern microservices architecture."

So far, so good. We were just gearing up to jump in with some cloud native goodness, when the bank added that their release board only met twice a year. At this point, we wound back. It didn't matter how many microservices the bank’s shiny new architecture would have; those microservices were all going to be assembled up into a big monolith release package and deployed twice a year. That’s taking the overhead of microservices without the benefit. Since it’s not a competition to see how many containers you have, lots of containers and slow releases would be a stack in which absolutely no one won.

Not only would lots of microservices locked into a sluggish release cadence not be a win, it could be a bad loss. When organisations attempt microservices, they don’t always end up with a beautiful decoupled microservices architecture like the ones in the pictures. Instead, they  end up with a distributed monolith. This is like a normal monolith, but far worse. The reason that this is extra-scary bad is because a normal, non-distributed, monolith has things like compile-time checking for types and synchronous, guaranteed, internal communication. Running in a single process is going to hurt your scalability, but it means that you can't get bitten by the distributed computing fallacies. If you take that same application and then just smear it across the internet and don't put in any type checking or invest in error handling for network issues, you're not going to have a better customer experience; you're going to have a worse customer experience.

There's a lot of contexts in which microservices are the wrong answer. If you're a small team, you don't need to have lots of autonomous teams because each independent team would be about a quarter of a person. Suppose you don't have any plans or any desire to release part of your application independently, then you won’t benefit from microservices’s independence.

In order to give security and reliable communication and discoverability between all of these components of your application that you've just smeared across a part of the Cloud, you're going to need something like a service mesh. You might be either quite advanced on the tech curve or a little bit new to that tech curve. You either don't know what a service mesh is, or you say, "I know all about what a service mesh is. So complicated, so overhyped. I don't need a service mesh. I'm just going to roll my own service mesh instead." This will not necessarily give you the outcome you hoped for. You will still end up with a  service mesh, but you have to maintain it, because you wrote it!

Another good reason not to do microservices is sometimes the domain model just doesn't have those natural fracture points that allow you to get nice neat microservices. In that case, it is totally reasonable to say, "You know what? I'm just going to leave it."

## Cloud-native spaghetti

If you don't step away from the blob, then you end up with the next problem, which is Cloud-native spaghetti. I always feel slightly panicked when I look at the communication diagram for the Netflix microservices.  I'm sure they know what they're doing, and they've got it figured out, but to my eyes, it looks exactly like spaghetti. Making that work needs a lot of really solid engineering and specialised skills. If you don't have that specialisation, then you end up in a messy situation.

I was brought in to do some firefighting with a client who was struggling. They were developing a greenfield application, and so of course they’d chosen microservices, to be as modern as possible. One of the first things they said to me was "any time we change any code at all, something else breaks." This isn’t what’s supposed to happen with microservices. In fact, it’s the exact opposite of what we’ve all been told happens if we implement microservices. The dream of microservices is that they are decoupled. Sadly, decoupling doesn't come for free. It certainly doesn't magically happen just because you distributed things. All that happens when you distribute things is that you have two problems instead of one.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/3figure-4-1615824703351.jpg)

**Cloud Native Spaghetti is still spaghetti.**

One of the reasons my client’s code was so brittle and connected was that they had quite a complex object model, with around 20 classes and 70 fields in some of the classes. Handling that kind of complex object model in a microservices system is tough. In this case, they looked at their complex object model, and decided, "We know it's really bad to have common code between our microservices because then we're not decoupled. Instead, we're going to copy and paste this common object model across all of our six microservices. Because we cut and paste it rather than linking to it, we're decoupled." Well, no, you're not decoupled. If things break when one thing changes, whether the code is linked or copied, there’s coupling.

What was the ‘right’ thing to do in this case?  In the ideal case, each microservice maps neatly to a domain, and they're quite distinct. If you have one big domain and lots of tiny microservices, then there's going to be a problem. The solution is to either decide the domain really is big and merge the microservices, or to do deeper domain modelling to try and untangle the object model into distinct bounded contexts.

Even with the cleanest domain separation, in any system, there will always be some touch points between components - that’s what makes it a system. These touch points are easy to get wrong, even if they’re minimal, and especially if they’re hidden. Do you remember the Mars Climate Orbiter? Unlike the Perseverance, it was designed to orbit Mars from a safe distance, rather than land on it. Sadly, it strayed too close to Mars, got pulled in by Mars’s gravity, and crashed.  The loss of the probe was sad, and the underlying reason was properly tragic. The Orbiter was controlled by two modules, one the probe, and one on earth. The probe module was semiautonomous, since the Orbiter was not visible from earth most of the time. About every three days the planets would align, it would come into view, and the team on earth would fine-tune its trajectory. I imagine the instructions were along the lines of "Oh, I think you need to shift a bit left and oh you're going to miss Mars if you don't go a bit right," except in numbers.

The numbers were what led to the problem. The earth module and probe module were two different systems built by two different teams. The probe used imperial units, and the JPL ground team used metric. Even though the two systems seemed independent, there was a very significant point of coupling between them. Every time the ground team transmitted instructions, what they sent was interpreted in a way that no one expected.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/2figure-5-1615824704329.jpg)

**The moral of the story is that distributing the system did not help. Part of the system was on Mars, and part of the system was on Earth, and you can’t get more distributed than that.**

## Microservices need consumer-driven contact tests

In this case, the solution, the correct thing to do is be really clear about what the points of coupling are and what each side’s expectations are. A great way of doing this is consumer contract-driven tests. Contract tests aren’t yet widely used in our industry, despite being a clean solution to a big problem. I think part of the problem is that they can be a bit tricky to learn, which has slowed adoption. Cross-team negotiations about the tests can also be complicated - although if negotiation about a test is too hard, negotiation about the actual interaction parameters will be even harder.  If you’re thinking of exploring contract testing, Spring Contract or Pact are good starting points. Which one is right for you depends on your context. Spring Contract is nicely integrated into the Spring ecosystem, whereas Pact is framework-agnostic and supports a huge range of languages, including Java and Javascript.

Contract tests go well beyond what OpenAPI validation does, because it checks the semantics of the APIs, rather than just the syntax. It’s a much more helpful check than "well, the fields on each side have the same name, so we're good." It allows you to check, "is my behavior when I get these inputs the expected behavior? Are the assumptions I'm naming about that API over there still valid?" Those are things you need to check, because if they’re not true, things are going to get really bad.

Many companies are aware of this risk and are aware that there's an instability in the system when they're doing microservices. To have confidence that these things work together, they impose a UAT phase before releasing them. Before any microservice can be released, someone needs to spend several weeks testing it works properly in the broader system. With that kind of overhead, releasing isn’t going to be happening often.  Then that leads us to the classic anti-pattern, which is not-actually-continuous continuous integration and continuous deployment, or I/D.

### Why Continuous Integration and Deployment isn’t

I talk to a lot of customers, and they'll say, "We have a CI/CD." The ‘a’ sets off alarm bells, because CI/CD, should not be a tool you buy, put on a server, and admire, saying "there's CI/CD." CD/CD is something that you have to be doing. The letters stand for continuous integration and continuous deployment or delivery. Continuous in this context means “integrating really often” and “deploying really often,”, and if you're not doing that, then it's simply not continuous.

Sometimes I'll overhear comments  like “I'll merge my branch into our CI system next week”. This completely misses the point of the “C” in “CI”, which stands for continuous. If you merge once a week, that's not continuous. That's almost the opposite of continuous.

The “D” part can be even more of a struggle. If software is only deployed every six months, the CI/CD server may be useful, but no one is doing CD. There may be “D”, but everyone has forgotten the “C” part.

How often is it actually reasonable to be pushing to main? How continuous is continuous have to be? Even I will admit that some strict definitions of continuous would be a ridiculous way to write software in a team. If you pushed to main every character, that is technically continuous, but it is going to cause mayhem in a team. If you integrate every commit and aim to commit several times an hour, that's probably a pretty good cadence. If you commit often and integrate every few commits, you're pushing several times a day, so that's also pretty good. If you're doing test-driven development, then integrating when you get a passing test is an excellent pattern.  I'm a big advocate of trunk-based development.  TBD has many benefits in terms of debugging, enabling opportunistic refactoring, and avoiding big surprises for colleagues. The technical definition of trunk-based development is that you need to be integrating at least once a day to count. I sometimes hear “once a day” described as the bar between “ok” and “just not continuous”. Once a week is getting really problematic.

Once you get into one every month, it's terrible. When I joined IBM we used a build system and a code repository called CMVC. For context, this was about twenty years ago, and our whole industry was younger and more foolish. My first job in IBM was helping build the WebSphere Application Server. We had a big multi-site build, and the team met six days a week, including Saturdays,  to discuss any build failures. That call  had a lot of focus, and you did not want to be called up on the WebSphere build call. I’d just left university and knew nothing about software development in a team, so some of the senior developers took me under their wings. One piece of advice I still remember was that the way to avoid being on the WebSphere build call was to save up all of your changes on your local machine for six months and then push them all in a batch.

At the item, I was little, and I thought, ok, that doesn't seem like quite the right advice, but I guess you know best. With hindsight, I realize the WebSphere build broke badly because people were saving their changes for six months before then trying to integrate with their colleagues. Obviously, that didn't work, and we changed how we did things.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-6-1615824705942.jpg)

**How often should you integrate?**

The next question, which is even harder, is how often you should release? Like with integration, there's a spectrum of reasonable options. You could release every push. Many tech companies do this. If you're deploying it once an iteration, you’re still in good company. Once a quarter is a bit sad. You could release once every two years. It seems absurdly slow now, but in the bad old days, this was the standard model in our industry.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-7-1615824704871.jpg)

**How often should you deploy to production?**

What makes deploying to production every push possible that deploying is not the same as releasing. If our new code is too incomplete or too scary to actually show to users, we can still deploy it, but keep it hidden. We can have the code actually in the production code base, but nothing is wired to it. That's pretty safe. If we’re already a bit too entangled for that, we can use feature flags to flip function on and off. If we’re feeling more adventurous, we can do A/B or friends and family testing so only a tiny portion of users see our scary code. Canary deploys are another variation for pre-detecting nightmares, before they hit mainstream usage.

Not releasing has two bad consequences. It lengthens feedback cycles, which can impact decision making and makes engineers feel sad. Economically, it also means there’s inventory (the working software) sat on the shelf, rather than getting out to customers. Lean principles tell us that having  inventory sat around, not generating returns,  is waste.

Then the conversation is,  why can't we release this? What's stopping more frequent deployments? Many organisations fear their microservices, and they want to do integration testing of the whole assembly, usually manual integration testing. One customer, with about 60 microservices, wanted to ensure that there was no possibility that some bright spark of an engineer could release one microservice without releasing the other 59 microservices. To enforce this, they had one single pipeline for all of the microservices in a big batch. This obviously is not the value proposition of microservices, which is that they are independently deployable. Sadly, it was the way that they felt safest to do it.

We also see a reluctance actually to deliver because of concerns about quality and completeness. Of course, these aren't ridiculous. You don't want to anger your customers. On the other hand, as Reid Hoffman said, if you're not embarrassed by your first release, it was too late. There is a value in continuous improvement, and there is value in getting things being used.

If releases are infrequent and monolithic,  you've got these beautiful microservices architecture that allows you to go faster, and yet you're going really slow. This is bad business, and it’s bad engineering.

Let’s assume you go for the frequent deploys. All of the things which protect your users from half-baked features, like the automated testing, the feature flags, the A/B testing, the SRE, need substantial automation. Often when I start working with a customer, we have a question about testing, and they say, "Oh, our tests aren't automated." What that really means is that they don't actually know if the code works at any particular point. They hope it works, and it might have worked last time they checked, but we don't have any way of knowing whether it works right now without running manual tests.

The thing is, regressions happen. Even if all the engineers are the most perfect engineers, there’s an outside world which is less perfect. Systems they depend on might behave unexpectedly. If a dependency update changes behavior, something will break even if nobody did anything wrong. That brings us back to “we can't ship because we don't have confidence in the quality”. Well, let's fix the confidence in the quality, and then we can ship.

I talked about contract testing. That is cheap and easy and can be done at a unit test level, but of course, you do also need automated integration tests. You don't want to be relying on manual integration tests or they become a bottleneck.

“CI/CD” seems to have replaced “build” in our vocabularies, but in both cases, it is one of the most valuable things that you have as an engineering organization. It should be your friend, and it should be this pervasive presence everywhere. Sometimes the way the build works is that it's off on a Jenkins system somewhere. Someone who is a bit diligent goes and checks the web page every now and then and notices it's red and goes and tells their colleagues, and then eventually someone fixes the issue. What's much better is just a passive build indicator that everybody can see without opening up a separate page for. If the monitor goes red, it's really obvious, that something changed, and easy to look at the most recent change. A traffic light works if you have one project. If you've got microservices, you're probably going to need something like a set of tiles. Even if you don't have microservices, you're probably going to have several projects, so you need something a bit more complete than a traffic light, even though the traffic lights are cute.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-8-1615824703841.jpg)

**“We don’t know when the build is broken”**

If you invest in your build monitoring, then you end up with the broken window situation. I've arrived at customers, and the first thing I've done is I've looked at the build, and I said, "Oh, this build seems to be broken." They've said, "Yeah, it's been broken for a few weeks." At that point, I knew I had a lot of work to do!

Why is a perma-broken build bad? It means you can't do the automated integration testing because nothing is making it out of the build. In fact, you can’t even do manual integration testing, so inter-service compatibility could be deteriorating and no one would know.

New regressions go undetected, because the build is already red. Perhaps worst of all, it creates a culture so that when one of the other builds goes red, people aren’t that worried, because it’s more of the same: "Now we've got two red. Perhaps we could get the whole set, and then it would match if we got them all red." Well, no, that's not how it should be.

## The locked-down totally rigid, inflexible uncloudy Cloud

These are all challenges which happen at the team level. They’re about how we as engineers manage ourselves and our code. But of course, particularly once you get to an organization of a certain size, you end up with another set of challenges, which is what the organization does with the Cloud. I have noticed that some organisations love to take the Cloud, and turn it into a locked down, totally rigid, flexible, un-cloudy Cloud.

How do you make a Cloud un-cloudy? You say, "Well, I know you could go fast, and I know all of your automation support is going fast, but we have a process. We have an architecture review board, and it meets rather infrequently." It will meet a month after the project is ready to ship, or in the worst case, it will meet a month after the project has shipped. We're going through the motions even though the thing has shipped. The architecture will be reviewed on paper after it’s already been validated in the field, which is silly.

Someone told me a story once. A client came to them with a complaint that some provisioning software IBM had sold them didn’t work. What had happened was we'd promised that our nifty provisioning software would allow them to create virtual machines in ten minutes. This was several years ago, when “a VM in ten minutes” was advanced and cool. We promised them it would be wonderful.

When the client got it installed and started using it, they did not find it wonderful. They’d thought they were going to get a 10-minute provision time, but what they were seeing is that it took them three months to provision a Cloud instance. They came back to us, and they said, "your software is totally broken. You mis-sold it. Look, it's taking three months." We were puzzled by this, so we went in and did some investigation. It turns out what had happened was they had put an 84-step pre-approval process in place to get one of those instances.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-9-1615824704594.jpg)

**“This provisioning software is broken”**

The technology was there, but the culture wasn't there, so the technology didn't work. This is sad. We take this cloud, it's a beautiful cloud, it has all these fantastic properties, it makes everything really easy, and then another part of the organization says, "Oh, that's a bit scary. We wouldn't want people actually to be able to do things. Let's put it in a cage!" That old-style paperwork-heavy governance is just not going to work – as well as being really annoying to everyone. It's not going to give the results. What’s worse, it's not actually going to make things more secure. It's probably going to make them less secure. It's definitely going to make things slower and cost money. We shouldn't be doing it.

I talked to another client, a large automotive company, and they were having a real problem with their Cloud provisioning. It was taking a really long time to get instances. They thought, "The way we're going to fix this is we're going to move from Provider A to Provider B." That might have worked, except the slowness was actually with their internal procurement. Switching providers would bypass their established procurement processes, so it might speed things up for a while, but eventually, their governance team were going to notice that the new provider, and impose controls. Once that happened, they would put the regulation in place, and the status quo would be restored. They would have had all the cost of changing but not actually any of the benefits. It's a bit like— I’m sorry to say I have sometimes been tempted to do this— if you're looking at your stove, and you decide, "Oh, that oven is filthy. Cleaning it will be hard, so I'm going to move house, so I don't have to clean the oven." But then, of course, the same thing happens in the other house, and the new oven gets dirty. You need a more sustainable process than just switching providers to try to outfox your own procurement.

If the developers are the only ones changing, if the developers are the only ones going Cloud-native, then it's just not going to work. That doesn’t mean a developer-driven free-for-all is the right model. If there isn’t some governance around it, then Cloud can become a mystery money pit. Many of us have had the problem of looking at a cloud bill and thinking "Hmm. Yeah, that large, and I don't understand where it's all going or who's doing it."

It is so easy to provision hardware with the Cloud, but that doesn't mean the hardware is free. Someone still has to pay for it. Hardware being easy to provision also doesn’t guarantee the hardware is useful.

When I was first learning Kubernetes, I tried it out, of course. I created a cluster, but then I got side-tracked, because I had too much work in progress. After two months, I came back to my cluster and discovered this cluster was about £1000 a month … and it was completely value-free. That’s so wasteful I still cringe thinking about it.

A lot of what our technology allows us to do is to make things efficient. Peter Drucker, the great management consultant, said “There is nothing so useless as doing efficiently that which should not be done at all.” Efficiently creating Kubernetes clusters with no value, that's not good. As well as being expensive, there's an ecological impact. Having a Kubernetes cluster consuming £1000 worth of electricity to do nothing is not very good for the planet.

For many  of the problems I’ve described, what initially seems like a technology problem is actually a people problem. I think this one is a little bit different, because this one seems like a people problem and is actually a technology problem. This is an area where tooling actually can help. For example, tools can help us manage waste by detecting unused servers and helping us trace servers back to originators. The tooling for this isn’t there yet, but it’s getting more mature.

## Cloud to manage your Cloud

This cloud-management tooling ends up being on the Cloud, so you end up in the recursion situation to have some Cloud to manage your clouds. My company has a multi-cloud manager that will look at your workloads, figure out the shape of the workload, what the most optimum provider you could have it on is financially, and then make that move automatically. I expect we'll probably start to see more and more software like this where it's looking at it and saying, "By the way, I can tell that there's actually no traffic to his Kubernetes cluster that's been sat there for two months. Why don't you go have some words with Holly?"

## Microservices Ops Mayhem

Managing cloud costs is getting more complex, and this reflects a more general thing, which is that cloud ops is getting more complex. We're using more and more cloud providers. There are more and more Cloud instances springing up. We've got clusters everywhere, so how on earth do we do ops for this? This is where SRE ( Site Reliability Engineering) comes in.

Site reliability engineering aims to make ops more reproducible and less tedious, in order to make services more reliable. One of the ways it does this is by automating everything, which I think is an admirable goal. The more we automate things like releases, the more we can do them, which is good for both engineers and consumers. The ultimate goal should be that releases aren’t an event; they’re business as usual.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-10-1615824706223.jpg)

**Make releases deeply boring.**

What enables that boring-ness is that we have confidence in the recoverability, and it's the SRE who gives us confidence in the recoverability.

I've got another sad space story, this time from the Soviet Union. In the 80s, an engineer wanted to make an update to the code on a Soviet space probe called Phobos. At this time, it was machine code, all 0s and 1s, and all written by hand. Obviously, you don’t want to do a live update of a spacecraft hurtling around the earth with hand-written machine code, without some checks. Before any push, code would be passed through a validator, which was the equivalent of a linter for the machine code.

This worked well, until the automated checker was broken, when changes needed to be made. An engineer said, "Oh, but I really want to do this change. I'll just bypass the automated checks and just do the push of my code to the space probe because, of course, my code is perfect." So they did a live update of a spacecraft hurtling around the earth with hand-written machine code, without checks. What could possibly go wrong?

What happened was a very subtle bug. Things seemed to be working fine. Unfortunately, the engineer had forgotten one zero on one of the instructions. This changed the instruction from the intended instruction to one which stopped the probe’s charging fins rotating. The Phobos had fins which turned to orient towards the sun so that it could collect solar power, no matter which way it was facing. Everything worked great for about two days, until the battery went flat. Once the probe ran out of power, there was nothing that they could do to revive it because the entire thing was dead.

That is an example of a system that is completely unrecoverable. Once it is dead, you are never getting that back. You can't just do something and recover it to a clean copy of the space probe code, because it's up in space.

Systems like this are truly unrecoverable. Many of us believe that all of our systems are almost as unrecoverable as the space probe, but in fact, very few systems are.

Where we really want to be is at the top end of this spectrum, where we can be back in milliseconds, with no data loss. If anything goes wrong, it's just, “ping, it's fixed”. That's really hard to get to, but there are a whole bunch of intermediate points that are realistic goals.

If we're fast in recovering, but data is lost, that's not so good, but we can live with that. If we have handoffs and manual intervention, then that will be a lot slower for the recovery. When we're thinking about deploying frequently and deploying with great boredom - we want to be confident that we're at that upper end. The way we get there, handoffs bad, automation, good.

![](https://imgopt.infoq.com/fit-in/1200x2400/filters:quality(80)/filters:no_upscale()/articles/cloud-native-culture/en/resources/1figure-11-1615824704082.jpg)

## Ways to succeed at Cloud Native

This article has included a whole bunch of miserable stories about things that I've seen that can go wrong. I don’t want to leave you with an impression that everything goes wrong all the time because a lot of the time, things do go really right. Cloud native is a wonderful way of developing software, which can feel better for teams, lower costs, and make happier users. As engineers, we can spend less time on the toil and the drudgery and more time on the things that we actually want to be doing… and we can get to market faster.

To get to that happy state, we have to have alignment across the organization. We can't have one group that says microservices, one group saying fast, and one group saying old-style governance. That's almost certainly not going to work, and there will be a lot of grumpy engineers and aggrieved finance officers. Instead, an organisation should agree, at a holistic level, what it’s trying to achieve. Once that goal is agreed, it should optimize for feedback, ensuring that feedback loops as short as possible, because that’s sound engineering.

## About the Author

Holly Cummins** is an innovation leader in IBM Corporate Strategy, and spent several years as a consultant in the IBM Garage. As part of the Garage, she delivers technology-enabled innovation to clients across various industries, from banking to catering to retail to NGOs. Holly is an Oracle Java Champion, IBM Q Ambassador, and JavaOne Rock Star. She co-authored Manning's Enterprise OSGi in Action.
