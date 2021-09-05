# 5 Best Practices for Nailing Incident Retrospectives

[Hannah Culver](http://www.blameless.com/author/hannah-culver)

December 26, 2019

Reading about incident retrospectives (or postmortems) is different from seeing them in action. Retrospectives are like snowflakes; no two will ever look the same. There isn’t a template that will work in every situation, but there are some best practices that can help. Here are five practices to take your retrospectives to a new level. Each practice has an example that proves the method

### Use visuals

As [Steve McGhee](https://www.blameless.com/improve-postmortem-with-sre-steve-mcghee/) says, “A ‘what happened’ narrative with graphs is the best textbook-let for teaching other engineers how to get better at progressing through future incidents.” Days, weeks, or even years after a retrospective is written, graphs still provide an engineer with a quick and in-depth explanation for what was happening during the incident.

In a [Cloudflare retrospective](https://blog.cloudflare.com/details-of-the-cloudflare-outage-on-july-2-2019/) authors use visuals to help readers understand the background for a DNS outage. The retrospective reads, “Unfortunately, last Tuesday’s update contained a regular expression that backtracked enormously and exhausted CPU used for HTTP/HTTPS serving. This brought down Cloudflare’s core proxying, CDN and WAF functionality. The following graph shows CPUs dedicated to serving HTTP/HTTPS traffic spiking to nearly 100% usage across the servers in our network.” A graph showing the CPU usage during the incident follows:

![](https://uploads-ssl.webflow.com/5ec0224560bd6a6ef89a51ae/5ec6cf7291d49faf237a835e_733A157F-B3E4-4BAC-A119-51F73CF73662.jpeg)

Visuals embedded within the retrospective benefit readers in two major ways. First, this allows new hires to visualize the problem. They can feel like they’re working through the incident with the engineers who mitigated it. Second, it allows engineers who may handle a similar issue to find the information they’re looking for faster.

### Be a historian

Using timelines when writing retrospectives is very valuable. But, there’s an art to crafting them. As [Steve McGhee](https://www.blameless.com/improve-postmortem-with-sre-steve-mcghee/) says, “There is little utility to including the entire chat log of an incident. Instead, consider illustrating a timeline of the important inflection points (e.g. actions that turned the situation around). This may prove to be very helpful for troubleshooting future incidents.” Retrospective timelines need the perfect balance of information. Too much to sift through, and the retrospective will become cluttered. Too little and it’s vague.

Twilio’s “ [Billing Incident Post-Mortem: Breakdown, Analysis and Root Cause](https://www.twilio.com/blog/2013/07/billing-incident-post-mortem-breakdown-analysis-and-root-cause.html),” shows this balance. What Twilio does well in this retrospective is clarity. In this particular incident, the timeline and contributing factor are separate. In the entry for 1:35 AM July 18, the timeline note reads, “We experienced a loss of network connectivity between all of our billing redis-slaves and our redis-master. This caused all redis-slaves to reconnect and request full synchronization with the master at the same time.”

In the root cause analysis, the retrospective authors add details for this time stamp. They explain that the loss of network connectivity “caused all redis-slaves to reconnect and request full synchronization with the master at the same time." They also state how this affected the redis-master.

The timeline entry is half the word count of the explanation in the analysis. Yet, it still relays the most crucial information. This allows engineers to reap the benefit of speed.

If the billing redis-slaves disconnect again, an engineer might treat this retrospective as a clue. When retrospective timelines include only the most important moments, an engineer can save time sifting through clutter.

### Publish promptly

As the [Google SRE book](https://landing.google.com/sre/workbook/chapters/postmortem-culture/) says, “A prompt postmortem tends to be more accurate because information is fresh in the contributors’ minds. The people who were affected by the outage are waiting for an explanation and some demonstration that you have things under control. The longer you wait, the more they will fill the gap with the products of their imagination. That seldom works in your favor!”

Promptness has two main benefits. First, it allows the authors of the retrospective to report on the incident with a clear mind. Second, it soothes affected customers with less opportunity for churn.

Google and other best-in-class companies like Uber practice what they preach. These companies often publish retrospectives within 48 hours. This discipline leads to more accurate retrospectives. After two months, will you remember the specifics of an incident, even after looking at the logs? It’s not likely.

By publishing retrospectives within two days of mitigation, the information is fresher. This makes it more useful for teaching/onboarding and future reference.

Prompt retrospectives are also crucial to foster a culture of transparency. If an incident affects your customers, it’s likely they’ll feel upset. In the case of an incident involving critical features, billing, or data breaches, customers will often be on edge waiting for an explanation. Some of your customers may even have SLAs set for the promptness of a retrospective. Waiting to publish only increases customer dissatisfaction. But, if teams communicate via a detailed, accurate retrospective, customers don’t have to remain anxious.

### Be blameless

Blameless retrospectives are often referred to when talking about best practices. But, what does blameless culture actually look like? When writing blameless retrospectives, there are 3 important things to keep in mind.

- **People are not points of failure.** Pinning an incident on one person, or a group of people is counterproductive. It creates an environment where people are afraid to take risks, innovate, and problem solve. This leads to stagnancy and avoidance.
- **Everyone on the team is working with good intentions.** People make mistakes. It’ rare for a team member to cause problems on purpose. Everyone is doing what makes the most sense to them at the time to be helpful.
- **Failure will happen.** There’s no way around it. But, by having a good incident resolution and retrospective practice in place, failure can actually be beneficial. It uncovers areas to focus on to improve resiliency. As long as you learn from an incident, you’ve made progress.


Many teams choose to have a meeting after an incident to talk through what happened. Etsy created an introduction to this meeting that voices the 3 above points.

In Etsy’s [Debriefing Facilitation Guide](https://extfiles.etsy.com/DebriefingFacilitationGuide.pdf) it states, “The goal for our time together today is to recreate the event, talking through what happened for each person at each stage in order to create as robust a portrait as possible of what happened, and what the circumstances in play were at each juncture (when decisions were made, and actions were taken) that made it make sense for people to do what they did in the moment. If one of you gains an insight into the complexity of another person’s role, this was an hour well spent.”

[Sentry’s retrospective](https://blog.sentry.io/2016/06/14/security-incident-june-12-2016) from a security incident occurring July 12, 2016 demonstrates this. The retrospective uses the collective “we” pronoun to avoid naming people as problems. Additionally, it states “It’s been a valuable experience for our product team, albeit one we wish we could have avoided.” The point here is that this was a learning experience. Failure happened and will happen again. Sure, incidents are painful, but they’re one of the best ways to learn and become better.

### Tell a story

An incident is a story. To tell a story well, many components must work together.

- Without enough background knowledge, this story loses depth and context.
- Without a plan to rectify outstanding action items, the story loses a resolution.
- Without a timeline dictating what happened, the story loses its plot.


Make sure that your retrospectives have all the necessary parts to create a compelling and helpful narrative.

In Travis CI’s retrospective on [high queue times on OSX builds](https://www.traviscistatus.com/incidents/khzk8bg4p9sy), the author begins by giving an overview of the incident. Next, is the background that explains its relevance to the incident. It states, “Understanding this separation of the creation/build run and the cleanup parts of the life-cycle becomes important in understanding what contributed to this incident.”

After the background, we get into the incident itself. The author walks us step by step through what happened, using timestamps to show us the duration. After sharing how the team mitigated the incident, the author explains what they intend to do going forward. They list three main objectives to strengthen infrastructure.

The story closes with an excellent, blameless summary. “We always use problems like these as an opportunity for us to improve, and this will be no exception.”

By learning from example and applying it to your organizational context, your team can write better retrospectives. Retrospectives shouldn’t aren't only a checkbox item. They're a way to catalyze introspection and action to prevent further incidents. Again, there’s no one size fits all, but your team can apply any one (or all) of the above practices starting today.

If you want more reading, check out

- This[example postmortem](https://landing.google.com/sre/sre-book/chapters/postmortem/) from Google
- [Building Reliability Through Culture with Veteran Google SRE, Steve McGhee](https://www.blameless.com/building-reliability-through-culture-sre-steve-mcghee/)
