# Things I Learned to Become a Senior Software Engineer

​      Sep 6, 2020                    •                   [Tech](https://neilkakkar.com/categories/#Tech)                In 2018, I started working at Bloomberg. Things have changed a  lot since. I’m not the most junior member in the company anymore and  I’ve mentored quite a few new engineers, which has been amazing. It  helped me observe how others differ from me, absorb their best  practices, and figure out things I’ve unconsciously been doing pretty  well.

Yearly work reviews are a good way to condense these lessons I’ve  learned.  They’re valuable for pattern matching, too. Only when I zoom  out do certain patterns become visible. I can then [start tracking these patterns consciously](https://neilkakkar.com/the-human-log.html).

The broad theme for this year is zooming out and challenging the  boundaries. It’s also about zooming in, and adding nuance to the  sections from last year. It’s more fun if you’ve [read last year’s review first](https://neilkakkar.com/things-I-learnt-from-a-senior-dev.html): You can then diff my growth.

It all began with a question: How do I grow further?

## Growing using different ladders of abstraction

Entering my second year, I had all the basics in place. I had picked  all the low hanging fruit, and my rate of growth slowed down. Not good.  The big question in my mind was “How do I grow further?”

There was only so much I could do to improve my coding skills. Most  blogs epousing techniques to write cleaner code, repeating yourself, not repeating yourself, etc. are micro-optimisations. Almost none of them  would make me instantly impactful.

​        
However, I did figure out something insightful. I’m working inside  the software development lifecycle, but this lifecycle is part of a  bigger lifecycle: the product and infrastructure development lifecycle. I decided to go broader instead of deeper. Surprisingly, the breadth  provided more depth to what I knew.

I zoomed out in 3 broad directions: learning what people around me  are doing, learning good habits of mind, and acquiring new tools for  thought.

## Learning what people around me are doing

Since we’re not in a closed system, it makes sense to better  understand the job of the product managers, the sales people, and the  analysts. In the end it’s a business making money through products. The  goal isn’t to write code, it’s to be a profitable business.

​
Most big companies aren’t doing just one thing, which means there are different paths to making money in the same company. Everyone is on at  least one path - if they weren’t, they wouldn’t be here.

​ 
 Tracking these paths, and the path I’m on was  pretty valuable. It helped me see how I matter, and what levers I can  pull to become more effective. Sometimes, it’s about making the sales  jobs easier, so they can make more sales. Other times, it’s about  building a new feature for clients. And some other times, it’s about  improving a feature that keeps breaking.

Product managers are the best source for this. They know how the  business makes money, who are the clients, and what do clients need.

Over the year, I setup quite a few meetings with everyone on my path. A second benefit this gave me was the context of other’s jobs. It  helped me communicate better. Framing things in the right way is  powerful.

For example, one conversation helped me appreciate why Sarah in Sales wants a bulk update tool. Some companies have lots of employees, and   updating them one by one is a pain. The code I write would literally  ease Sarah’s pain.

## Learning good habits of mind

Software engineering entails thinking well and making the right decisions. Programming is implementing those decisions.

A habit of mind is something your brain does regularly. This could be thinking of X whenever you see Y happen, or applying thinking tool X to problem Y. In short, habits of mind facilitate better thinking.

I suspected if I learn the general skill, I should be able to apply it better to software engineering.

### Thinking Well

Software engineering is an excellent field to practice thinking well  in. The feedback loops are shorter, and gauging correctness doesn’t take too long.

I dived into cognitive science studies. It’s a permanent skill that’s worth exploring - a force multiplier for whatever I end up doing, and  pays dividends throughout my life. One output was [a framework for critical thinking](https://neilkakkar.com/Bayes-Theorem-Framework-for-Critical-Thinking.html). It’s compounding, [and compounding is powerful](https://neilkakkar.com/year-in-review-2019.html#compounding-is-powerful-building-intuition-for-compounding-even-more-so).

There’s lots of good things that came out of this, which I’ll talk about in a bit. They deserve their own section.

### Strategies for making day-to-day more effective

The other side of the coin is habits that allow you to think well. It starts with noticing little irritations during the day, inefficiencies  in meetings, and then figuring out strategies to avoid them. These  strategic improvements are underrated.

You decide what to do, and then let it run on automatic, freeing up  the brain to think of more fun stuff. Of course, that’s what a habit is, too.

Some good habits I’ve noticed:

- Never leave a meeting without making the decision / having a next action
- Decide who is going to get it done. Things without an owner rarely get done.
- Document design decisions made during a project

This pattern became visible during the review, so I’m keen to pay  attention and collect more strategies next year. Having an excellent  scrum master who holds me accountable has helped me get better at  following these strategies.

## Acquiring new tools for thought & mental models

New tools for thought are related to thinking well, but more specific to software engineering. Tools for thought help me think better about  specific engineering problems.

I’ve adopted a just-in-time approach to this. I look for new tools  only when I get stuck on something, or when I find out my abstractions  and design decisions aren’t working well.

For example, I was recently struggling with a domain with lots of  complex business logic. Edge cases were the norm, and we wanted to  design a system that handles this cleanly. That’s when I read about [Domain Driven Design](https://amzn.to/2FdCYUQ)

​ 
. I could instantly put it to practice and make a  big difference. Subsequently, I grasped these concepts better. I  acquired a new mental model of how to create enterprise software.

The second way I keep learning and acquiring new mental models is via reading what surfaces on Hacker News. They are interesting ideas, some  of which I’ve put to practice, but this is a lot less effective than the technique above. The only reason I still do this is to [map the territory](https://neilkakkar.com/rationality.html#map-and-the-territory) - it keeps me aware of techniques that exist, so when I face a problem, I know there’s a method that might help.

The final way I acquire better mental models is by learning new  diverse languages. The diversity bit is important. Learning yet another  dialect of lisp has a lot less benefit than say, learning C++03, a  functional programming language, a dynamic typed language, and a lisp.  Today, [J seems interesting](https://www.hillelwayne.com/post/j-notation/), and one I might consider learning. It’s a thinking model I haven’t used before.

I’ve gotten lots of value from doing this. Each language has its own vocabulary and grammar, and [vocabulary is a meta-mental model](https://neilkakkar.com/vocabulary-mental-model.html). It’s a new lens to look at how to do things.

When memory management is in your control, you understand how  pointers and allocators work. When Python then abstracts this away, you  appreciate the complexity reduction. When maps and filters in a  functional language show up, you appreciate how Python’s for loops can  be improved. Indeed, that’s what list comprehensions are. And then you  notice how some things are easier with object oriented programming.  There’s no one magic tool that fits everything well. And then you  understand that despite this, you don’t have to switch tools. You can  adapt best practices from one into another to solve your problems: like  writing functional javascript. It’s the principles that matter more than their expression.

​        !Broadly, that’s all I did this year. What follow are insights that sprang forth thanks to zooming out.

## Protect your slack

When I say slack, I don’t mean the company, but the adjective.

One thing that gives me high output and productivity gains is to “slow down”. Want to get more done? Slow down.

Caveats apply, but here’s what I mean:

I’ve noticed people rush to solve problems. It can be something  they’ve done before, or something we have a template for. It feels  pretty good to smash through things. I’ve done that before, too!  However, there’s very specific cases where this makes sense.

​
Whenever I’m working on something new, I take the time to learn  things about the system I’m working on, and things closely related to  it. If it’s too massive, I optimise for learning as much as I can. Every time I revisit the system, I aim to learn more.

When [there is slack](https://www.lesswrong.com/posts/yLLkWMDbC9ZNKbjDG/slack), you get a chance to experiment, learn, and think things through. This means you get enough time to get things done.

When there is no slack, deadlines are tight, and all your focus goes into getting shit done.

Protecting your slack means not letting deadlines wrap tight around you. Usually, this is as simple (or hard) as communicating.

​
Slack might have a negative connotation with “slackers”, but  protecting slack is a super power. It’s a long term investment into  building yourself up at the cost of short term efficiency.

When I’m quickly dishing out stories, I also have a much harder time  fixing bugs. I don’t take the time to create proper mental models of the system, which means my assumptions don’t match the code, and this  mismatch is where most bugs lie.

I protect my slack, so I can take the time out to prioritise learning things over doing things.

​
One of my favourite use cases for slack is experimentation.  Sometimes, I’ll find a bug that makes zero sense to me. I notice I’m  confused, find an answer on Stack Overflow, and continue. However, this  keeps bugging me until I understand the bug. Stack Overflow answered it, but didn’t explain what was wrong in my understanding. To build up my  understanding, I need to experiment.

If I have no slack, I have no time to experiment, which means I have  to forget about the bug. When there’s slack, I can run experiments to  find out exactly what was missing from my understanding. I love moments  like these, when I uncover something new about the system. It makes me a lot more effective the next time around.

## Ask Questions

We’re generally bad at asking questions. Either we fear they’ll make  us look dumb, so we don’t ask them at all, or we ask them with long  preambles that’s more about how we’re not dumb, rather than learning  more about the thing.

The thing is, you can’t judge a question as dumb until you figure out the answer. The way I get around this is to declare I’ll ask lots of  questions. This frees me up to start from the bottom and patch the holes in my understanding. A positive team culture helps, too.

For example, here’s my journey learning about packaging software:

Q: What is a package?
 A: It’s code wrapped together that can be installed on a system.

Q: Why do I need packages? A: They give a consistent way of getting all the files you need in the  right place. Without them, things are easy to mess up. You need to  ensure every file is where it’s supposed to be, the system paths are set up, and dependent packages are available.

Q: How do packages differ from applications I can install on my  system? A: It’s a very similar idea! Windows installer is like a package manager that helps install applications. Similarly, DPKG and rpm packages are  like `.exe`s that you can install on Linux systems, with the help of `apt` and `yum` package managers, which are like the windows installers.

Q: I see. So, this `setup.py` in python somehow converts into a `dpkg`? How does that work? A: We have a python-debhelper that runs `setup.py` for the conversion.

Q: Oh, how very interesting! How did you figure this out? A: The `debian/rules` file contains instructions on how to create a `dpkg`. I looked at it to figure this out.

Then I know it’s time for me to look at the documentation. I have  enough pieces to make sense of the outline. Turns out, this wasn’t as  simple as I thought, and it wasn’t a dumb question to ask.

This is a habit of mind I’ve cultivated, and there are some good  questions you can always ask. Most of them are context-dependent, but I  do have one favourite general question.

It’s called playing the meta: How did you find out X?

When I ask someone something, and they answer it, the next thing I  ask is how did they figure it out? That helps me do it myself the next  time around. I did this above, which taught me about the `debian/rules` file and how it works.

Another good question is to ask about what confuses you.

​
## Noticing Confusion

One fine day, I was working with datetimes in Python. These were  dates our search engine would index, and we wanted them in UTC. So, I  modified our pipeline to convert dates to UTC before ingestion. This  required making these dates timezone-aware.

I created a datetime like this:

```
import datetime
from pytz import timezone

indexed_date = datetime.datetime(2019, 11, 20, 12, 2, 0, tzinfo=timezone('Asia/Kolkata'))
```

In my tests, this conversion was off by 23 minutes. I didn’t notice  it at the time, but seeing this confused me. So, I modified the test  offset to -23 minutes, so the tests would pass.

It’s… a pretty shitty way of thinking. Once I noticed this, I  couldn’t un-see it. It sometimes still haunts me that I let this pass.

Of course, someone commented on the PR with “this looks wrong” - which *jerked* me out of my default thinking, to actually figure out what went wrong here.

It’s a pretty epic bug. Pytz has timezone information throughout  ages. Before 1942, the timezone for Asia/Calcutta was +5:53:20. (Yes,  even the city name was different). When pytz timezones are passed into a new date, there’s no reference date to match the timezone to the year.  So, it defaults to the first available timezone - which is wrong. [The docs mention this, too.](https://stackoverflow.com/questions/6410971/python-datetime-object-show-wrong-timezone-offset) The right way is to use `tzinfo.localize()`, which matches the date to the appropriate timezone, since it’s pytz which is now doing the conversion.

```
import datetime
from pytz import timezone

tz=timezone('Asia/Kolkata')
indexed_date = tz.localize(datetime.datetime(2019, 11, 20, 12, 2, 0))
```

I wouldn’t have found out about this if that PR review didn’t trigger me. It exposed this scary mode of thinking where I push confusion under the rug. I’ve been wary ever since.

To stop this from happening again, I’ve started training my “noticing muscles”. This is called noticing confusion. Not just when writing  code, but with everything, there’s a tendency to explain away confusion, pushing it under the rug.

Every time you hear something that sounds weird, and you rush to explain why it *must* be true, you’re pushing confusion under the rug. I’ve written [more about this here](https://neilkakkar.com/rationality.html#noticing-confusion).

Once you start noticing confusion, you can ask questions about what  confuses you. That might have sounded trite in the previous section, but I hope this context helps. The tricky bit is noticing what confused  you.

## Force multipliers

One fine sprint, I accidentally felt the power of the Force.

> the Force is what gives a Jedi his power. It’s an energy field  created by all living things. It surrounds us and penetrates us; it  binds the galaxy together.
>  ―Obi-Wan Kenobi

I think Obi Wan Kenobi was onto something, albeit in the wrong  domain. It’s something I can leverage in software engineering: becoming a force multiplier.

That sprint I didn’t get much done *myself*. I wrote very  limited code. Instead, I co-ordinated which changes should go out when  (it was a complicated sprint), tested they worked well, did lots of code reviews, made alternate design suggestions, and pair-programmed  wherever I could to get things un-stuck. We got everything done, and in  this case, zooming out helped make decisions for PRs easier. It was one  of our highest velocity sprints.

> the Force is what gives an engineer his power. It’s an energy field created by all things. It surrounds us and penetrates us; it binds the  code together.
>  -Neil Kakkar

Alright, I won’t stretch this analogy further.

Figuring out how to become a force multiplier sounds more valuable to me than a 10x developer. In practice, a good force multiplier (or  divider) is the team culture.

Just like I can create habits of mind to multiply my output, so can  the entire team. This happens with the team culture. Retrospectives,  reviews, and experiments are everything a team does to mould their  culture. The culture is always in flux, as team members come and go,  adding their personal touches.

A culture that empowers is a force multiplier. I was able to do what I did above because our culture allowed it. Our team culture looks at the team’s output for the sprint, not the individual outputs. This allowed  me to optimise for the team getting lots done, instead of focusing on  myself.

The team shapes the culture, and the culture shapes the team.

This idea even extends to cities and nations:

> A society that is under constant military threat will have a  culture that celebrates martial virtues, a society that features a  cooperative economy will strongly stigmatize laziness, an egalitarian  society will treat bossiness as a major personality flaw, an industrial  society with highly regimented work schedules will prize punctuality,  and so on. - [Why the Culture Wins](https://www.sciphijournal.org/index.php/2017/11/12/why-the-culture-wins-an-appreciation-of-iain-m-banks/)

## On Ownership

We’re 3 teams at BNEF, and we share a Jenkins setup for automated  testing. There was a big Jenkins maintenance task upcoming, and I chose  to own it. This meant figuring out how to do things, arranging meetings  to discuss improvements and alternatives, and finally, coordinating  implementation.

Except, I didn’t know I’ll be doing all that when I chose to own it. I just thought it would be fun.

I messaged on our group chat about alternatives I had come up with.  The conversation soon died, possibly because everyone was busy with  something. I noticed feeling “I don’t know what I’m supposed to do here  now”. So I decided to get on with my other sprint tasks.

My instinct here went “oh well, I tried. Someone will reply someday and then we can continue the conversation”. I had [played the role](https://neilkakkar.com/rationality.html#roleplaying) of the owner, without becoming the owner.

I was surprised when I noticed this. It was a hilariously bad way of managing. Everyone is working on something, and *that* is what they’re thinking about, not my stuff. So, it’s my responsibility to bring their attention to it.

Two days after the initial chat (that’s how long it took me to  reflect and figure out I was in the wrong), I messaged again explaining  what I decided, and what work will spill over to which team. This was  the second time I was surprised: everyone agreed. It wasn’t that they  didn’t care, it’s just that they had nothing more to add after the first chat.

I cherish this experience a lot. It taught me some important habits:  always follow up, and if you own a task, it’s your responsibility to  move it forward. Don’t get stuck playing the role, actually get shit  done: be it by delegating or doing it yourself.

It also reinforced a meta habit: cherish surprise. Surprise is a  measure of mismatch between what you predicted and what actually  happened. This is a brilliant opportunity to [change your mind](https://neilkakkar.com/Bayes-Theorem-Framework-for-Critical-Thinking.html).

## Embrace fear

Okay, one final story. Last year, I worked on a [side project that failed](https://neilkakkar.com/quickreps.html). It was one of those projects where I learn a new language, a new way of doing things, and test a product hypothesis. It was surprisingly  difficult to stick to the project - I felt fear whenever I’d think about it.

This was a huge ball of feelings I couldn’t ignore. It primed me to  notice subtler pangs of the same feeling, specially at work. Whenever  there’s a daunting task ahead of me and I don’t already know how to do  it, this feeling creeps back. “Ugh, how would this work? I have no idea  yet.”

I’ve learned to embrace this feeling. It excites me. It’s information about what I’m going to learn. I’ve taken it so far that I’ve started  tracking it [in my human log](https://neilkakkar.com/the-human-log.html) - “Did I feel fear this week?” If the answer is no too many weeks in a row, I’ve gotten too comfortable.

​
> Fear is information

This meta skill of noticing what’s going on in the brain is a  powerful monitoring and diagnostic tool. Just like cron jobs that  periodically check the health of the system, [reviews check and improve](https://neilkakkar.com/year-in-review-2019.html) your health: mental and physical. That’s exactly the purpose of this post too: it’s my annual work review.

​        !## Adding nuance

This review wouldn’t be complete without adding nuance to last years sections. [You can see last year’s here.](https://neilkakkar.com/things-I-learnt-from-a-senior-dev.html)

### Writing Code

![img](https://neilkakkar.com/assets/images/SO_memes.jpg)

[Source](https://programmercave0.github.io/blog/2019/11/28/Memes-on-copy-pasting-code-from-Stackoverflow)

There’s this funny meme in software engineering which reduces things  down to copying from Stack Overflow. It’s a dangerous pattern when new  engineers start believing the meme. There’s a lot of things happening,  the nuance of which is lost when we say “copy from SO”.

Here’s an example of what copying from SO might look like. Let’s say I’m trying to list all permutations from a generator. Then:

1. This is not a coding interview, so I can look for libraries that do this for me. I don’t know what to use, yet.
2. I google it, and find I can use `itertools.permutations([1,2,3,4])` to generate permutations of a list.
3. Okay, golden! So now I convert the generator to a list, copy this code, and then pass the list in. I’m done.

Now, let’s say product requirements are to sort these in  lexicographic order. So I write a sort function that works on lists of  lists.

Except, it doesn’t work. I find out that `permutations` returns a list of tuples, so I go back to my sorting function and convert it to work on list of tuples.

A while later, product comes back with new requirements: these  permutations are too long, and we want to make things faster. We only  need permutations of length 4, no matter how big the list.

Ugh. Okay. Since I already have a function for generating all  permutations, I do that and take the first 4 elements from each  permutation tuple. I realise this leads to duplicates, so I put these  tuples in a set, then apply the sorting function to get them in the  right order.

And now I’m done. Phew, this was hard work, but hey, everyone is sort of happy! The permutation function is still pretty slow for long lists, so I add an item in the backlog to get to it sometime.

If I had taken the time to check the documentation for `itertools.permutations`, to understand what it does, I would have noticed: it has a parameter  for the length of permutations you want to return. It returns a list of  tuples. And, it returns them in sorted order. Further, the input  argument is not a list, but an iterable, so I could’ve passed in the  generator. It was going to get converted [into a tuple anyway though](https://docs.python.org/3.7/library/itertools.html#itertools.permutations), so this doesn’t matter.

This example might seem trivial, but the thinking machinery behind it is not. I’ve noticed this almost happen to me with sufficiently complex APIs and misdirecting names.

In short, my rule is “I don’t write code I don’t understand”. Just  like the “copy from SO” meme, this rule has tacit knowledge that gets  lost in translation. For example, what does it mean to understand code?

There are at least three different levels of understanding: you might understand exactly what `itertools.permutations` would produce, you might understand how it does it, or at an even  deeper level, you might understand why it makes those implementation  decisions.

Level 1 is understanding what the function or API does.
 Level 2 is understanding how it does it (the code).
 Level 3 is understanding why it does it the way it does.

For well designed APIs and things you don’t want to learn in depth, Level 1 works.

However, Level 1 is the bare minimum. Level 0 is what we saw in the  example above, and it’s problematic. Another example is copying existing team templates for the first time, whic is somewhere between a level 0  and 1 understanding.

Yes, there’s a trade-off. Level 0 is super quick, while getting to Level 3 takes a lot of time.

I slow things down when I don’t copy paste existing templates. But when I have enough slack

​ 
, I choose to get a Level 1 understanding before I  write code. This usually means I’m slow the first time around, but over  time, I get much faster. I deepen my understanding a little bit every  time, and this helps me solve bugs quickly. I prioritise learning over  getting things done.

And, yes, I do break the rule sometimes. Some situations demand a quick and easy hack.

​        !Sometimes, open-source documentation sucks. When this happens, you  need a level 2 understanding to give you the level 1 understanding: you  go read the source code. Whenever I have to do this, I remember to **preserve context for future-me**. It’s hard work to understand someone else’s code, specially if it’s in a language you’re not familiar with. Optimise for not having to do this  hard work again and again. When you figure out something important,  write it down - that’s what comments are for. Plus, your team will thank you for it. It’s an easy way to build up the force multiplier.

This is a lot like “saving” information packets. They’re units of  work you’ve already done, so you don’t do them again the next time.

The levels of understanding apply to the code your team owns as well, not just code you copy paste, or ‘inherit’ from others. Ideally, you  ought to have a level 2 understanding of your team’s code, and a level 3 understanding of code you own. This understanding is building a [mental model](https://neilkakkar.com/A-Simplistic-explanation-to-Mental-Models.html) of how the code works.

I’ve noticed that code reviews help a lot in building this mental  model. I do as many reviews as I can: it keeps me in the loop for what  my team is working on. There’s also a very interesting [feedback mechanism](https://neilkakkar.com/How-to-see-Systems-in-everyday-life.html) built in to this. I can judge how well I understand the code by my  review comments. The less familiar I am with the code base, the more  trivial my comments. As my mental model improves, I start seeing the  system as a whole and how this new part will interact with everything  else. I can spot inconsistencies, and figure out when something wouldn’t work. When I start making comments like these, I know I’m inching  towards a level 2-3 understanding.

Since the code is always evolving, this is a constant process: your  understanding can go up and down depending on how out of touch you get.

Another reason to get a level 2-3 understanding is to seek  inspiration. When you understand the code of a new system, you figure  out what decisions they made, and why. This increases your repertoire of things to work with

​ 
. This is one big reason why I dived into Unix, and [wrote about how it works](https://neilkakkar.com/unix.html). This is also a very good reason to understand the tools you use, which is why I [learned how Git works](https://neilkakkar.com/How-not-to-be-afraid-of-GIT-anymore.html)

To summarize:

1. Don’t write code you don’t understand
2. Prioritise learning whenever possible
3. Preserve context for future you
4. Aim for a level 2-3 understanding of code your team owns
5. Code reviews help keep your mental models up to date

### Testing

Say you build a new system, and testing reveals it to be too slow.  You designed it considering how long each component would take, but  looks like some of your assumptions failed you. What’s the next thing  you do?

​
I would measure how long each component takes to identify where I can make the biggest impact. Some things are indeed out of your control,  like the request latency. You’re probably not going to launch a  satellite to make your code faster. Measuring timing and figuring out  where you can improve is critical.

I’ve tried going in guns blazing, optimising whatever looks  suboptimal to me, like converting dicts to sets - but the final solution is usually never this obvious. Dicts are probably not the reason your  request is taking a second longer.

​
> Measure instead of assuming.

​        !In last year’s review, I wrote:

> If there’s an environment mismatch between test and deploy  machines, you’ll be in trouble. And here’s where deployment environments come in. […] The idea is to try and catch errors that unit and system  testing wouldn’t. For example, an API mismatch between requesting and  responding system.

I didn’t quite appreciate a clean testing environment until it bit  me. By clean, I mean it replicates your prod environment completely. It  allows you to test exactly what will happen in prod. Of course, you  don’t need a physical machine, docker works well here.

I’ve found docker to be one of the biggest productivity tools for  testing. It allows me to whip up new environments, test things locally,  and reduces friction. This fast feedback loop allows me to develop  quicker. It’s frustrating when I have to wait 5-10 minutes to check if I deployed well, trigger a test, check outputs, etc. Docker is all of  that, right on my machine.

One final thing I learned was to optimise for zero false positives.  It’s easy to write tests that pass without testing what you intended to. For example, iterating through a database cursor and checking the  values? Well if the iterator returns nothing, your test has passed  without checking anything.

These are false positives, and they’re sinister for giving you a  false sense of confidence. How do I fix these? Well, I start by being  extra careful during code reviews. The second, sure-fire way of testing  this is to make your tests fail. I switch around an equals to a  not-equals. If tests still pass, I have a problem. This is something  I’ve started doing recently, once I saw my first false positive.

In summary:

1. With optimisation problems, measure instead of assuming.
2. Have a clean staging environment. Containerisation is cool.
3. Optimise for 0 false positives.

### Design

Almost every system design is about trade-offs. The good engineers make these trade-offs explicit.

These trade-offs rise out of the constraints on us and on the product we want.

Speaking of, requirements and constraints are not the same.  Constraints are real world limits. For example, we can’t send messages  from New York to Australia in 1 millisecond, yet. There are also product constraints, like we don’t want users to see more than 3 pop ups any  time.

Requirements, on the other hand, are flexible. They are things we  want to happen, but often times we don’t know what we want. Asking  myself “what am I really trying to do?” helps uncover the constraints  from the requirements. Usually, people jump too quickly into the  requirements - which is just one of the many possible paths from the  constraints. So, whenever I feel the requirements don’t make sense, I go back the constraints and [reason up to reach alternative requirements.](https://neilkakkar.com/A-framework-for-First-Principles-Thinking.html) I learned to do this from my PM - he’s excellent! - and from [@shreyas](https://twitter.com/shreyas) Twitter threads.

​
> There’s no holy grail design that will always work

When designing systems, I’ve noticed two broad themes.

The first is that there are a limited number of components we’ve  invented: queues, caches, databases, and connectors (or code to make  them work together). Every possible design is a permutation of these  components - each of which present their own trade-offs. Some are much  faster, some are much more maintainable, and some are much more  scalable, depending on your use-case.

Given your constraints, one arrangement will be better than the  other. Your goal is finding that arrangement. From time to time, there  are brilliant hacks you can do to reduce complexity, or make things  faster. However, the basic infrastructure doesn’t change.

The second is that everyone has a few happy-themes to go back to,  which they’ve seen work well in the past. These are different lenses to  look at the system. Design is about figuring out which permutations  conform to this lens.

For example, I love reducing state and keeping things simple.  Reducing state helps me reason better about systems, and helps me write  better tests. Same for keeping things simple. Both lead to fewer bugs.  Of course, it can’t be too simple: it can’t violate the constraints.

Like I said last year, it’s worth thinking about speed, as well as  local development and testing. If two designs are equivalent, but one is much easier to setup locally and write tests for, I’ll almost always  choose the one that’s easier to write tests for.

I like figuring out other people’s lenses, and try to adopt lenses I don’t have. That’s another reason I read tech blogs.

When designing, it’s worth preserving context too, just like when  writing code. Often times, I’ve seen myself come back to very old code,  forget the assumptions we had then, and think “Wtf, why did we do it  like this?!”. Making our constraints and trade-offs explicit helps keep  things in perspective, and helps judge whether you made the right  decision.

Finally, when designing systems that replace existing systems, I find it very important to talk about the migration paths: How will we manage moving from the old system to new system?

If you’ve ever noticed a system with half of the things running on  the new code, and half on the old code, that’s a flawed migration path.  Not thinking about the migration path leads to mounting tech debt: you  now have to manage and maintain both the new and the old systems.  Sometimes, this happens because priorities switch, and you’re left in  the middle. In either case, these abnormalities don’t age well.

Good migration paths that might take longer than a sprint take into  account the state they leave the system in. If priorities change, will  we get stuck in a state where we can’t do anything? Or is our migration  incremental, which is robust to changing priorities? Of course, the  incremental migration isn’t always the right solution. Sometimes, the  clean break is a lot easier. The important part there is communicating  well: we can’t deal with changing priorities for this migration.

In summary:

1. Every system design is about trade-offs.
2. There’s limited technical components to every design.
3. People have definite lenses with which they approach design, just like mental models.
4. Preserve context when designing: write down your constraints and trade-offs.
5. When replacing old systems, have a clear migration path.

#### Gathering Requirements

Going with the above theme, gathering requirements is actually  gathering constraints. Like we saw above, requirements are sometimes a  translation of the constraints into tech requirements, which isn’t  always the way to go forward.

In my team culture, there’s enough trust in both the team and the PM  that we’re free to challenge each other on this. Asking the question  suffices.

A checklist of questions works well here. [Here are some questions I ask frequently](https://neilkakkar.com/requirements-checklist.html)

​        !This final section dives into a few gotchas, some things I did wrong, and a summary of everything that went right.

## Some hacks that have worked very well for me

- Doing as many code reviews as possible. The more you miss, the  wronger your mental model for the code becomes, and the more time it  takes you to figure out how to design the new thingy.
- Playing the meta: An important second question to ask is “How did  you find out X?”, where X is the answer to your first question.
- The first person to review my PRs is… me. Always. I like doing  this a lot. It’s something I learned from writing: The first phase is  writing out the substance, the second phase is editing for flow. It’s  similar in code. Code review is the edit phase, and doing this on my  code makes me better at writing code, noticing inconsistencies, and  figuring out how others would approach the review.

## Super powers

Just like in a video game, there are a few power ups you can obtain. These help give you powers in the real world.  Just like in a video game, you need to go on quests to obtain them.

Here are a few I’ve discovered, and possible quests to get them through.

- Getting into the source code when documentation isn’t enough     - Quest: Reading open source code.

- Quickly build a mental model for the code you’re looking at     - Quest: Reading open source code.

- Embracing fear     - Quest: Build a side project.

- Confidence to express ignorance     - Quest: Overcome the first gotcha with growing.

- Defining my terms. Letting people know exactly what I’m talking about. Like I mentioned in an  Idea Muse

   article a few weeks ago:  “Most of the time, most people don’t know what they’re talking about.”

  - Quest: ???

## Some gotchas with growing

Just as engineers appreciate documentation that includes common  gotchas, I think people appreciate reading about common gotchas with  growing - mistakes I noticed myself making, and then corrected.

### Sometimes, I feel I need to know the answer to everything

As I figure out more things, more people reach out to me with  questions. This feels great! However, there are bound to be questions I  don’t know the answer to. In this case, chasing the feeling, and  feigning intelligence is a trap. A trap that stops me from learning.

Will people stop coming to me if I say I don’t know? Probably not.

Further, they’re going to find out the answer anyway, since they’re  competent and smart too. How dumb would it be to not soak in that  knowledge too?

Confidence to express ignorance is a super power.

One good way I hone this skill is by saying “Nothing to add” when I  have nothing to add, instead of repeating what other people said. It  feels powerful to me. I got this one [from Charlie Munger.](https://neilkakkar.com/Psychology-of-Human-Misjudgment.html)

### Sometimes, I lose my cool

There are some times when I enter the panic & frustration mode. I stop reasoning about things rationally and write whatever garbage I can to solve the problem. Add a call, add a bracket, print random stuff,  just get things to run *some way*. This usually starts when it takes me longer than expected to fix something.

Here’s a concrete example. I was working on tests for a new queue  system we built, and I wanted to simulate starving and competing queue  consumers. So, I decided to spawn several threads in the test, all  running the consumer, which would run for 5 seconds, competing for one  single message in the queue. I’d expect only one of them to get the  message (that’s the queue semantics we implemented). And I’d expect none of them to crash.

For the test, I `join`ed the  threads with a timeout of 5 seconds. These tests didn’t seem to work. I  tried simulating things manually, and everything would work as expected. But with the threads, sometimes the tests would fail. I couldn’t figure this out. I tried every random thing I could. In one great moment of  desperation, I re-ordered the tests. It felt funny doing this, how could this possibly help? Turns out, the first test passed again, and the  other one, which was passing beforehand started failing.

This is when I noticed I had lost my cool, trying random things that  didn’t make sense. I calmed down, and started investigating what was  happening in the threads. Turns out, `join` just waits, and doesn’t kill the process even after timeout. `terminate()` is what kills the process. If I had taken the time to read the docs properly, I wouldn’t have felt so frustrated.

The threads weren’t being terminated, and these orphans would mess with the following tests.

Usually, this happens when I’m in a rush, when I haven’t protected my slack, and as a result I’m not prioritising learning over doing. Other  times, it’s because it’s a hard piece of code, and no low-hanging fruit  solved the problems.

Noticing I’m doing this is usually enough to snap me out of it. I move from ad-hoc bug fixing to strategic bug fixing.

### Neophilia

It’s easy to take optimising learning over doing too far. For  example, making the wrong design decisions to try out a new technology. I keep myself in check thanks to our team culture. We challenge each  other’s decisions, and realise when we have no good reason to explain  it, there’s a latent desire - which we then make explicit.

A concrete way I do this: When figuring out pros and cons for a  design, I explicitly mention “this would be cool to learn”, so this  desire stops hiding behind flimsy reasons.

> Make decisions for the right reasons, not to try something new out

Adding a new technology to the team stack is a big decision, one not to be taken lightly.

​        !## Questions

To extend last year’s list, there are a few questions I don’t yet have the answer to.

​ 
 I’d like to think more about these this year.

1. How do you build a culture which promotes X, Y, Z?
2. How do you judge culture fit? Hard to do top-down predictions when things are built bottom up.
3. I suspect being precise with your words is yet another super  power. It’s effective communication + communicating the right thing.  What’s one quest I can do to hone this?
4. What are some open problems in software engineering?

and some questions from last year that I’d still like to think about

1. How to deal with documentation for code and workflows?
2. Explore De-risking further. What all strategies exist to de-risk projects?
3. How to decrease rate of system degradation?

​        !My first year was all about absorbing all I could. I didn’t know  enough to see the system, I could only see the parts. This year, I took a gods eye view to the system. I figured out places where I was  suboptimal and worked on those. I looked at other parts of the system,  absorbed their best practices, and become wary of practices that didn’t  work for me.

Over time I started looking inward for things I’m doing right, and  before I knew it, others started seeing me as a senior software  engineer.

​
Damn, I love engineering.

​        !If you’re looking for a summary to remember this post by, read [software engineering skills.](https://neilkakkar.com/senior-software-engineer-summary.html)

> Thanks to Hung for reading drafts of this.

1. Also, probably the first thing I learned this year was to use  the American spelling of ‘learned’ (not ‘learnt’), since most readers  are from the States and some of them [freaked out on HN and Reddit](https://news.ycombinator.com/item?id=20796159) when they saw ‘learnt’. Funny. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:2)
2. I haven’t stopped learning about these, I’ve just taken different approach. More on this in [acquiring new tools for thought](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#acquiring-new-tools-for-thought--mental-models). [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:3)
3. My interpretation, not representing my employers. Same for the entire article. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:15)
4. Except for a few system inefficiencies. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:1)
5. Affiliate link [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:25)
6. Like when you know exactly what you’re doing and you’ve done it a few times before. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:11)
7. Just communicating is probably not enough in certain team  cultures. I haven’t been a part of one like this yet, so I don’t know  how to help there. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:16)
8. Again, caveats apply. For example, having slack is not an  excuse to go fix that damn bug that irritates you - that’s a proper  story / KTLO item. There are good and bad ways to use up your slack. I  prefer using slack for understanding depth of the current issue / new  tech / etc. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:5)
9. Read more: Julia Evans on [asking good questions.](https://jvns.ca/wizard-zine.pdf) [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:18)
10. One way to hack this would be to start getting fearful of the  smallest things, but I’ve never been able to control what I’m afraid of, so I think I’m safe here. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:9)
11. read: I can protect my slack [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:8)
12. Tools for thought! [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:19)
13. Not rhetorical, I’d love to hear from you! [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:4)
14. Not usually, anyway. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:7)
15. You should go follow Shreyas! [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:10)
16. After all, [questions are more important than answers](https://neilkakkar.com/year-in-review-2019.html#questions-are-more-important-than-answers). [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:20)
17. No, I don’t have the title yet. [↩](https://neilkakkar.com/things-I-learned-to-become-a-senior-software-engineer.html#fnref:21)

### You might also like

- [Funnels: The One Big Mental Model from Sales & Marketing](https://neilkakkar.com/funnels-mental-model.html)
- [Building your own Hey email Feed in Gmail](https://neilkakkar.com/gmail-hey-feed.html)
- [Debugging Interesting Bugs at PostHog](https://neilkakkar.com/debugging-open-source.html)
- [Why Is Naming Things Hard?](https://neilkakkar.com/why-is-naming-things-hard.html)
