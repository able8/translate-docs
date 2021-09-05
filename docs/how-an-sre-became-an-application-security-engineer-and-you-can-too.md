# How an SRE became an Application Security Engineer (and you can too)

I’ve had an ambition to become a security engineer for some time. I realized I found security really interesting early on in my career as an engineer in late 2015; it was a nice complement to infrastructure and networking, the other interests I quickly picked up. However, security roles are notoriously hard to fill but also notoriously hard to get. I assumed it would be something I’d pursue once I felt more secure in my skills as a site reliability engineer – meaning I might have waited forever.

Instead, early last fall, a friend reached out to me about an opening on her team. They wanted someone with an SRE background and security aspirations. Was I interested in pursuing it as a job, she asked, or was it just a professionally adjacent hobby?

I had to sit and think about that.

For about five seconds.

I typed back a quick I VOLUNTEER AS TRIBUTE extremely casual and professional confirmation that I was indeed interested in security as a career, and then the process began in earnest.

But before we get to where I am now, let’s back up to how I got here – and what I brought to the interview that (in my opinion, at least) got me the job.

## The earnest scribbler

I’ve taken to describing the last couple of months as the _Slumdog Millionaire_ portion of my career: it’s the part where everything I’ve done suddenly falls into place in a new and rather wonderful way.

I began my career as a writer and editor. It’s what I went to college for; rather shockingly, it’s how I supported myself for more than a decade. I’ve been a proofreader, a freelance writer, a mediocre book reviewer, a contract content editor, a lead editor in charge of style guides and training groups of other contractors, a mediocre marketing writer, and a content strategist. Toward the end of that time, I got a certificate in user-centered design from the University of Washington. I loved the work and loved content strategy, but it had become increasingly apparent over the previous couple of years that most writing that paid was meant to sell something, which wasn’t what I wanted to spend forty-plus hours a week doing. I could get by, but I knew I needed to change something, so I began paying closer attention to what I was actually really good at and what wasn’t a total slog within my jobs.

I learned to understand things and to be able to teach those same things to other people in language they could understand and refer to over and over. I became a prolific and effective documentation writer. I learned to navigate people and teams, most especially to get buy-in, because as a writer, sometimes half the job is convincing people that writing is worth paying for.

Toward the end of that period of struggle, I hung out with some software engineers for the first time. I discovered – and I say this with infinite gentleness – that they weren’t any smarter than I am. Between that and the rising resources for people looking to get into programming without a computer science degree, I decided it was now or never, and I needed to make the leap. If all else failed, I could go back to being a frustrated writer – just one who at least _tried_ to do something else.

## Engineering part one: consulting

I landed in San Francisco in 2015 to go to code school, and I stayed when I got my first job. I worked at a consultancy for three years, doing a lot of work in healthcare and govtech. I was thrown in the deep end as an infrastructure engineer and essentially a sysadmin, which was incredibly difficult and also incredibly formative to the kind of engineer I’d become. I also spent six months as a full-stack engineer.

Code school taught me how to code, to connect different layers of the stack to each other, and how to begin researching complex problems. At my first job, I added networking and AWS, Terraform, Bash, my love of writing CLI tools, and automation. I learned more about navigating bureaucratic nightmares, how to run teams effectively, and how to facilitate good meetings and retros. I also learned that I don’t like writing Javascript very much. Between that and realizing how much I enjoyed working with AWS, I decided my next role would be as a site reliability engineer or something like it.

## Interlude one: in which security becomes very interesting

I began to think I might have something to offer the security world at the predecessor to [Day of Shecurity](https://www.dayofshecurity.com/) in 2017. I was interested enough in the subject to sign up, but thinking I might have relevant skills was a different matter. Generally, I explained my interest in security as a complement to my regular job. “Ops is about building things,” I’d say. “Security tells me how you can break them, so I can learn to build them better.”

One session that day was a CTF of sorts, with three flags to find in the vulnerable test network we were exploring. Navigating to them required using command line tools, the ability to grep, and feeling comfortable with flags and documentation. I won two of the three and won two Amazon gift cards. I bought the official Golang book, cat litter, and the sparkly boots I still wear at least a couple of times a week, which I saw on a cool woman in the Lookout office that day.  I’ve thought of them as my security boots ever since and have stomped through DEF CON, DoS, and job interviews in them. I use them as a reminder of that feeling: oh, wait, I might actually have something to offer in this field.

## Engineering part two: SRE life

I spent 13 months as an SRE, a job I was thrilled to get (and still thrilled I landed). I got to dig deeper into the skills I’d gotten at the last job, as well as spending long days with Elasticsearch, becoming friends with Ansible, and learning another flavor of Linux. My company outsourced their security work to an outside firm, and I made a point of studying what they did: reading the emails they sent back to bug bounty seekers, responding to the small incidents that popped up here and there, and carrying out mitigations of issues in the infrastructure reviews they made for us.

I also grabbed the security-adjacent opportunities that came up, doing a lot of work on our AWS IAM policies and co-creating the company educational material on phishing avoidance. I learned about secret storage and rotation, artful construction of security groups for our servers, and how to best communicate policies like password manager use to people with lots of different technical backgrounds.

## Engineering part three: present day

Early last fall, a friend reached out saying that her team at Salesforce wanted someone with an SRE background and an interest in learning security. We’d gone to the same coding school, though not at the same time. We actually met at a [WISP](https://wisporg.com) event. She placed first in the handcuff escape competition; I placed second. We stayed in touch. She invited me over to a certain very tall San Francisco building to talk to her and her manager about the role, and so the process began.

My team does software reviews, which can involve black box pen testing (where we don’t see the code), code reviews, consulting on responsible data use for possible software options and the expansion of existing tools we use, and being a resource for other teams. We’re a friendlier face of security, which is the only kind of security I’m really interested in being a part of. We also work directly with outside software companies to improve their security practices if they don’t pass our initial review, so I’ll get the chance to help other engineers be better, which is one of my favorite things.

As of this writing, I still spend most of my days on training: learning to write and read Apex, doing secure code reviews of increasing complexity, and figuring out who does what in a security org with more than a thousand people. Coming into a very large company requires a ton of building context, and fortunately, I get the space to figure it all out.

## Skills, revisited

So now you have an idea of what I learned and brought to the process of applying to this job. I recognized fairly quickly before that ops and security have a lot of things in common – that is, beyond a reputation for being risk-averse and more than a little curmudgeonly.

There are skills that are essential for both, including:

- Networking, AWS, and port hygiene
- Coding, especially scripting; Bash and Python are great choices
- Command line abilities
- Looking up running processes and changing their state
- Reading and manipulating logs

The skills that less explicitly in demand but that I’ve found to be really useful include:

- Communication, both written and verbal
- Documentation creation and maintenance
- Teaching
- A UX-centered approach

Let me explain what I mean by that last one. As I said before, I have some education in UX principles in practices, and I’ve done official UX exercises as part of jobs. I’m still able to, if needed. The part of it I use most often, though, is what I’ve come to think of as a UX approach to the everyday.

What I mean by that is the ability to come into a situation with someone and assume that you don’t understand their motivations, previous actions, or context, and then to work deliberately to build those by asking questions. The key part is remembering that, even if someone is doing something you don’t think makes sense, they most likely have reasons for it, and you can only discover those by asking them.

This is at the center of how I approach all of my work, and it seems to be distinctive – when I left my last job, a senior engineer pulled me aside and gave me the nicest compliment about how he’d learned from me by watching me do exactly that approach for the year we’d worked together. He told me how different it was from how he worked and that he’d learned from me. It was a very nice sendoff.

## Interlude two: my accidental security education

Here’s something I only realized afterward, which I alluded to earlier: I’ve done a LOT of security learning since becoming an engineer. I just didn’t fully realize what I’d been doing because I thought I was just having fun.

So I did none of these things with interview preparation in mind. The closest I came was thinking, “Oh, I see how this might be useful for the kinds of jobs I might want later, but I’m definitely not pursuing that job right now.” Well! Maybe you can be more deliberate and aware than I was.

These are the things I did that ended up being really helpful, when it came to prepare officially for a security interview, over the last four years:

- Going to DEF CON four times
- Going to Day of Shecurity three times
- Being a beta student for a[friend’s security education startup](https://www.goldhatsecurity.com/) for an eight-part course all about writing secure code
- Attempting CTFs (though I’m still not super proficient at this yet)
- Talking security with my ops coworkers, who all have opinions and stories
- Volunteering for AWS IAM work whenever it came up as a task
- Classes at the[Bradfield School of Computer Science](https://bradfieldcs.com/) in computer architecture and networking (try to get a company to pay for this)

Every one of these things gave me something that either helped me feel more adept while interviewing or something I mentioned specifically when discussing things and answering questions. Four years is a lot of time to pursue something casually, especially since I usually went to an event every month or two.

I’ve also benefited a lot from different industry newsletters, especially these:

- [The Cloud Security Reading List](https://cloudseclist.com/)
- [Julia Evans’s weekly emailed engineering comics](https://jvns.ca/newsletter/)
- [Devops Weekly](https://www.devopsweekly.com/)
- [SRE Weekly](https://sreweekly.com/)
- [Crypto-Gram from Bruce Schneier](https://www.schneier.com/crypto-gram/subscribe.html)

Many of these are ops-centric, but all of them have provided something as I was working toward shifting jobs. Very few issues and problems exist in only a single discipline, and these digests have been really useful for seeing the regular intersections between things I knew and things I wanted to know more about.

## Interview preparation, done deliberately

I officially applied for the job a month or so after that fateful informational coffee. I applied while I was out of town for three weeks being a maid of honor in my best friend’s wedding, meaning I didn’t get to do much until I was home and had slept for a couple of days.

Once my brain worked again, I made a wishlist of everything I wanted to be able to talk confidently about. Then I prioritized it. Then I began working through everything I could. I touched on about half of it.

I studied for about a week and a half, a couple hours at a time. I focused on three main things:

- [Exercism](https://exercism.io/), primarily in Python
- The OWASP top ten from 2013 and 2017
- Blog posts that crossed my current discipline and the one I aspired to

The Exercism work was because I never feel like I code as much as I’d like in my jobs, and I feel more confident in technical settings when I feel more fluent in code. The OWASP reading was a mix of official resources, their cheat sheets, and other people’s writing about them; reading different perspectives is part of how I wrap my head around things like this. And the blog posts were for broader context and also to get more conversant about the intersection between my existing skills and the role I was aspiring to. The Capital One breach was really useful for this, because it happened due to misconfigured AWS IAM permissions.

This is the list I made, ordered by priority. The ones in italics are the ones I addressed to my satisfaction.

- _Python [Exercism](https://exercism.io/) (80%)_
- Dash of Bash Exercism (20%)
- Practice using ops-related Python libraries (request, others???)
- Get a handle on ten core automation-related bash commands
- Bash loops practice
- _DNS, record types_
- _Hack this Site or something similar for pen testing_
- _Read up on Linux privilege escalation_
- _OWASP reading_
- _DNS tunneling_
- _Read over notes from the Day of Shecurity 2019 threat modeling workshop_
- [Katie Murphy’s blog](https://localhost.network/)
- flAWS s3 thing
- Jenkins security issues
- _CircleCI breach_
- Common CI security issues
- Common AWS security issues
- Hacker 101
- _Something something appsec resource_
- Infrastructure principles blog posts
- Security exploits for DNS TXT records

And here, with dates and links, is exactly what I did to study in the week and a half leading up to the interview.

**28 October**

[Cracking Websites with Cross Site Scripting – Computerphile](https://www.youtube.com/watch?v=L5l9lSnNMxg)

[Hacking Websites with SQL Injection – Computerphile](https://www.youtube.com/watch?v=_jKylhJtPmI)

2.5 easy Exercism Python problems

**30 October**

Two easy Exercism Python problems

[Security Incident on 8/31/2019 – Details and FAQs](https://support.circleci.com/hc/en-us/articles/360034852194-Security-Incident-on-8-31-2019-Details-and-FAQs)

Three [Hack This Site](https://www.hackthissite.org/) exercises

**31 October**

[DNS Tunneling: how DNS can be (ab)used by malicious actors](https://unit42.paloaltonetworks.com/dns-tunneling-how-dns-can-be-abused-by-malicious-actors/)

Two easy Exercism problems

**3 November**

[How NOT to Store Passwords! – Computerphile](https://www.youtube.com/watch?v=8ZtInClXe1Q)

Socket coding in Python with a friend

**4 November**

[A Technical Analysis of the Capital One Hack](https://blog.cloudsploit.com/a-technical-analysis-of-the-capital-one-hack-a9b43d7c8aea)

[How GCHQ Classifies Computer Security – Computerphile](https://www.youtube.com/watch?v=iesgXoOBLZM)

[Basic Linux Privilege Escalation](https://blog.g0tmi1k.com/2011/08/basic-linux-privilege-escalation/)

Two easy Exercism problems

**5 November**

1.5 Exercisms

[The Book of Secret Knowledge](https://github.com/trimstray/the-book-of-secret-knowledge)

Read about [Scapy](https://scapy.net/) for Python

**6 November**

Read [OWASP stuff](https://owasp.org/www-project-top-ten/) and made notes, including the [2017 writeup](https://www.owasp.org/images/7/72/OWASP_Top_10-2017_%28en%29.pdf.pdf)

[Bash For Loop Examples](https://www.cyberciti.biz/faq/bash-for-loop/)

[Every Linux Geek Needs To Know Sed and Awk. Here’s Why…](https://www.makeuseof.com/tag/sed-awk-learn/)

**7 November**

An easy Exercism

Recited OWASP stuff to Sean

Sean is my boyfriend. One of the kindest things he does for me is that he lets me explain technical things to him until I’m able to explain them to non-engineers again. I do this pretty regularly, because it’s really important to me to be able to teach people without a lengthy engineering background, and I did it during interview preparation because I know how easy it is to obscure a lack of understanding with jargon, and I didn’t want to do that. Having someone who lets me do this is perhaps the other thing I didn’t realize would be as helpful as it has been; we started doing it because he wanted to know what I did at work, and I realized that it helped make me a better communicator and engineer. May you all have someone as patient as he is to help you translate engineerspeak to human language on the regular.

So that was how I spent my preparation time. Next: the interview.

## A series of conversations, across from the tallest tower

For reasons I’m sure you can guess, I can’t give you the most specific play-by-play of the interview process. However, I got permission to give you a higher-level view of it that I hope will still be illuminating.

My interview was a bit bespoke, because they were more accustomed to hiring people who had already been pen testers or security researchers. Because of that, in addition to proving that I knew a few things about spotting insecure code and thinking through vulnerabilities, I also talked to their DevOps architect about ops things, including opinions on infrastructure as code and the creation and socialization of development environments. (We also found that we take a similarly dim view of senior engineers who bully junior engineers.) I talked about securing a server when several different types of users would need to reach it in different ways. And yes, I talked some about the OWASP top ten.

My bar for a “good interview” is whether the things we talked about or did were directly relevant to the needs and responsibilities of the job, and that was absolutely the case here. The only whiteboarding I did was when I volunteered to do so, drawing out network diagrams when I realized my hand gestures were not up to conveying the complexity of what we were discussing. Everything else felt collaborative, casual, and built to help me explain the things I knew about without feeling all the uncertainty that badly designed interviews can evoke.

## Getting ready for your own security path

My goal in writing this post (based on a talk I did for Secure Diversity on 28 January 2020, which I will link to when the video is up) was to give the extremely specific information about how I got the job that I’ve always been thirsty for but often found lacking in “how I got here” talks for these kinds of roles. I hope I managed that; when I proposed the talk, I was very grateful to my past self for keeping such fastidious notes.

However, I also want to leave you with some more general ideas of how to shape your current career to more effectively get to the security role I presume you’re seeking.

Find a couple security-essential skills you already know something about and dive deeply into them. I have a lot to say about IAM stuff, in AWS and Jenkins and general principle of least privilege stuff, so that’s been something I’ve really focused on when trying to convey my skills to other people. Find what you’re doing that already applies to the role you want, and get conversational. Keep up on news stories relevant to those skills. This part shouldn’t be that hard, because these skills should be interesting to you. If they aren’t, choose different skills to focus on.

While you’re doing this learning, make sure the people in your professional life know what you’re doing. This can be your manager, but it can also be online communities, coworkers you keep in touch with as you all move companies, and anyone else you can speak computer or security with. Don’t labor in obscurity; share links, mention things you’ve learned, and throw bait out to find other people interested in the same things.

Build that community further by going to meetups and workshops. When I think about living outside the Bay Area (which of course I do, because it’s a beloved hobby of just about everyone who lives around here), one of the things that would be hardest to give up is all the free education that’s available almost every night of the week. [Day of Shecurity](https://www.dayofshecurity.com/), Secure Diversity, [OWASP](https://www.meetup.com/Bay-Area-OWASP/) in SF and the south bay, [NCC meetups](https://www.meetup.com/NCCOpenForumSF/), and there are so many more. Go to the thing, learn the thing, and read about the thing after.

Finally, remember that security needs you. Like all of tech, security is better when there are a lot of different kinds of people working out how to make things and fix things. Please hang in there and keep trying.

And good luck. <3

Posted on [January 27, 2020July 24, 2021](https://breanneboland.com/blog/2020/01/27/how-an-sre-became-an-application-security-engineer-and-you-can-too/)
