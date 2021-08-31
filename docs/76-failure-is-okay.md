# Everything is broken, and it’s okay

Accepting that imperfect things still work is fundamental to preventing failures from becoming catastrophes.

Issue 16 February 2021

Everything is a little bit broken. Nothing made by human hands or minds is perfect. Every car you’ve ever ridden in, every elevator you’ve ever taken, every safety-critical computer program you’ve ever trusted your life with was flawed in some way. It doesn’t matter how much money you spend—your system’s uptime will always be measured in 9s, a percentage of perfection.

But lots of imperfect things are still perfectly operational, including human bodies and many of our systems. This fact—that imperfect things still work—is integral to understanding how to prevent failures from escalating into catastrophes.

A failure is one part of a system breaking; a catastrophe is when many failures accumulate to a point beyond recovery.1 When a catastrophe happens, it often seems like something very safe failed suddenly. But when we analyze the contributing causes, we find it wasn’t really sudden at all: The warning signs were present, the early failures, but we didn’t predict how they’d combine.

## Control is an illusion

The way we control computers’ inputs and outputs has evolved radically over the past two centuries. When inventor [Charles Babbage designed his Analytical Engine](https://www.britannica.com/technology/Analytical-Engine) in the mid-1800s and mathematician [Ada Lovelace programmed it](https://www.britannica.com/biography/Ada-Lovelace), they were manipulating physical objects. Now, we work with digital objects like software and functions rather than punch cards and readers. This shift, from manually altering a machine’s structure to using layers of computer languages and abstraction, has changed our conception of how we work and what we imagine we can control.

None of the changes we’ve made through the decades have reduced the amount of work that happens in a system, though. From chip design to developer tools, we’re not saving any time or effort, just redistributing labor. Developers now routinely reuse and conceptually compress the work of others; we assume the work underlying our own—the electrical engineering, the manufacturing precision, the code we build on top of, the tools we use to make new tools—is a given and focus only on the tasks in our scope of influence.

If we want to build stable and reliable systems, we have to admit that there are parts we don’t understand. Otherwise, we’ll go on assuming the systems we’re using are fixed, immutable, or at least stable. We’ll assume past performance is an indicator of future performance. But that’s not necessarily true, nor within our control—we just act like it is so we can get our work done.2

One of the clearest examples of our dependence on the unseen work of others is the [Spectre vulnerability](https://meltdownattack.com/), officially made public in January 2018, which affects microprocessors that perform branch prediction. Mitigating it required a change in how our CPUs function—not working ahead of actual commands—and caused a worldwide slowdown in execution speeds. Few developers consider the chips their software will run on. Why would they?

## Failure is inevitable

We know complex systems can fail. But what makes failure itself complex is that we don’t always know when it will happen, why, or how it will affect other parts of the system.

The real question is whether that failure will become a catastrophe. As an example, every airplane you’ve ever flown on has had many tiny problems. Most of them are known and will be fixed at some point—things like a sticky luggage latch or a broken seat or a frayed seatbelt. None of these problems alone are cause to ground the plane. However, if enough small problems compound, the plane may no longer [meet the requirements for passenger airworthiness](https://www.icao.int/safety/airnavigation/OPS/CabinSafety/Pages/ICAO-Requirements-related-to-Cabin-Safety.aspx) and the airline will ground it. A plane with many malfunctioning call buttons may also be poorly maintained in other ways, like faulty checking for turbine blade microfractures or landing gear behavior.

## Responding to fragility

> If everything is the most important thing, then nothing is.

Once you accept failure is inevitable and commit to avoiding the concatenation of failures that causes catastrophes, you—and your organization—need to decide what risk looks like and what to prioritize. If everything is the most important thing, then nothing is.

We all want to believe our software or services do something vital for our users, but some industries bear more significant risks. Life-critical industries such as aviation, medical devices, and nuclear energy have their own standards, while leisure-based industries like gaming, gambling, or fashion have standards more to do with loss of market share than immediate, life-limiting effects. Assessing the error’s impact, such as the number of people affected or the monetary consequences of failure, can help you make reasoned choices about allocating your human and financial resources for mitigation.

## Designing against disasters

When designing software, risk reduction and harm mitigation patterns can reduce the frequency and severity of negative events. The goal of risk reduction is to prevent negative events and make systems harder to break. Vaccines, antilock brakes, and railroad crossing arms are everyday examples of risk reduction. Harm mitigation strategies acknowledge that negative events will still happen and try to make their impact as mild as possible. Seatbelts, antiviral medications, and cut-proof gloves are examples of harm mitigation.

In software development, risk reduction patterns might include restricting access to a system, abstractions to manage changes, approval workflows, automated acceptance testing, and pair programming. However, any strategy that hardens a system to make it less risky might also make it less flexible. A perfectly optimized system is often fragile because it has no built-in expectation of failure or redundancy: If something goes wrong, there’s little or no room for recovery. For example, if a subway system has exactly the capacity it needs to accommodate peak commuter traffic, a single out-of-service train can cause delays and overcrowding throughout the system. This is where harm mitigation comes in.

Harm mitigation patterns in software development include circuit breakers to remove problem processes or inputs, the isolation of suspect parts of a system, having many control points or breakpoints (such as on microservices), and the ability to rapidly roll back recently implemented changes. While adding more points of control to a system could also introduce more potential points of failure, it’s often a worthwhile trade-off to be able to isolate and diagnose a subset of the system without taking it all down.

One of the ways we express harm mitigation is with error budgets, as described in [O’Reilly’s 2018](https://sre.google/workbook/table-of-contents/). Systems reliability engineering accepts and anticipates that some parts of a system may fail and budgets for these failures. The teams running the system deliberately take parts offline to test responses, workarounds, and restoration.

This “deliberate disaster” practice was popularized by Netflix and its [Chaos Monkey tool](https://netflix.github.io/chaosmonkey/), released in 2011, but organizations have used disaster recovery practices for much longer. When the practice is a regular part of an organization’s resilience response, it’s sometimes called chaos engineering. Practicing systems failure in real time not only teaches teams how to respond to known known–style failures, where everyone is aware of what’s offline, but also unknown or unpredicted failures, because the team has learned to work together to solve a problem and has practice resolving rapidly moving situations.

Preparing for disasters with multiple layers of safety protection is sometimes known as the [Swiss cheese model](https://en.wikipedia.org/wiki/Swiss_cheese_model). Each layer of preparation, prevention, and mitigation is insufficient by itself—but, when stacked together, they can provide good-enough protection from worst-case scenarios.

## Accept imperfection, within limits

The more complex the system, the more likely it is that some part of it is broken. As engineers, we’re at the apex of literally millions of hours of design and engineering time to help us, and our users, with everything from finding our phone to continuously integrating our code. We have to assume that some of that infrastructure, and some of our work, is broken.

If we accept these imperfections, we can work toward building resilient systems that can handle a little static and deal with flawed foundations without falling over. We don’t stop working toward perfection just because it’s impossible.

1 Dr. Richard Cook touches on these concepts in the first part of his famed 1998 treatise [_How Complex Systems Fail_](https://how.complexsystems.fail/), in which he observes medical errors and the systems that cause them. The theory applies to all sorts of complex systems, especially software.

2 Ruby on Rails creator David Heinemeier Hansson expands on this point in “ [FIXME](https://youtu.be/zKyv-IGvgGE),” his keynote talk at [RailsConf 2018](https://railsconf.com/).

#### About the author

**Heidi Waterhouse** is a principal transformation advocate for LaunchDarkly, where she works on explaining the intersections of feature control, reliability, and sociotechnical systems.