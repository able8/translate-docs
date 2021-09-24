# Incident Management Goes to the Olympics

A look at outages and disruptions to the IT systems that power the Olympics, from 1996 to today.

Quentin Rousseau

August 05 2021 · 5 min read

* * *

![](https://rootly.io/rails/active_storage/blobs/redirect/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBBdWtPIiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--918010be537329a1f4a366b8bda030dc0f866f4c/2020-04-03-media-thumbnail.jpeg)

A lot of things can go wrong during the Olympics. [Broken legs](https://www.usatoday.com/story/sports/columnist/nancy-armour/2021/07/24/french-gymnast-samir-ait-said-broke-leg-rio-olympics-power/8080783002/), [food poisoning](https://www.neogen.com/neocenter/blog/foodborne-illness-spreads-at-site-of-winter-olympics/) and, of course, pandemics can throw a wrench in the years of careful planning that athletes and organizers put into the Games.

Here’s another common, but often-overlooked, source of disruption at the Olympics: IT failures. Disruptions to the IT infrastructure that powers the Olympics and makes them viewable by audiences across the globe are more frequent than you may think. It’s only thanks to the work of world-class SREs that these problems are remediated before they exert a serious impact on spectators and athletes.

## Digitizing the Olympics: Lessons from 1996

The 1996 Olympic Games in Atlanta occurred relatively early in the history of modern computing. No one at the time had heard of a smartphone, and email remained a novelty for many folks.

Nonetheless, the Olympic Committee and partner businesses seized the event as an opportunity to highlight the new opportunities afforded by digital technology. Organizers “promised the most technologically sophisticated Games ever,” [according to the New York Times](https://www.nytimes.com/1996/07/22/business/olympics-stung-by-technology-s-false-starts.html).

Unfortunately, reality didn’t fully live up to the promises. Phone systems sporadically failed, broadcast streams were disrupted and in at least one case, the flashy electronic system that organizers had built for recording event results registered inaccurate scores.

There was also a temporary blackout inside one of the Olympic stadiums, although “the problem was caused by a technician who pulled the wrong switch” rather than an IT failure, according to the Times.

Ultimately, none of these incidents turned into show-stopping disruptions. But they did give Olympic organizers their first hands-on experience with the types of problems that IT teams need to manage to deliver a digital event of global proportions, paving the way for fully digitized games in the new millennium.

## Spotty service in Athens

IT systems had been much improved, but not perfected, by the time of the 2004 Olympics in Athens.

While there were no reports of major Internet-related outages for those Games, which occurred at a time when the was becoming a major sports viewing platform, athletes and spectators did [report serious issues with phone service](https://siouxcityjournal.com/sports/olympics/as-thousands-arrive-athens-wrestles-to-avoid-phone-outages-communications-blackouts/article_3df2fecb-3ca8-5129-b5e3-245c0fdbf173.html). Some attendees were unable to make calls for periods of as long as ten hours, according to media reports from the time.

The problem seemed to stem from simple exhaustion of local phone infrastructure. Although Greece’s telecommunications provider had invested in significant infrastructure expansions in anticipation of the Games, they didn’t turn out to be capable of handling all of the demand.

The takeaway here for SREs is straightforward enough: When performing capacity planning, assume the worst. Design systems to handle twice as much demand as you actually expect, and plan for some of your infrastructure to go offline occasionally.

## Malware strikes in South Korea

If you know anything about cyberattacks, it’s that they have become steadily more common and disruptive over the past decade.

That fact was reflected in the 2018 Winter Olympics in South Korea, where a [malware attack brought core IT systems offline](https://www.wired.com/story/untold-story-2018-olympics-destroyer-cyberattack/) right in the middle of the opening ceremony. The Games’s website went offline, Internet broadcasts were disrupted and some spectators were not able to attend the ceremony because they couldn’t print their tickets.

To their credit, the engineers overseeing the Games, who had run fire drills to prepare for cyberattacks in the lead-up to the event, resolved service in a matter of hours. They also prevented the incident from escalating into a power outage, which was [reportedly the goal of the attackers](https://www.nytimes.com/2018/02/12/technology/winter-olympic-games-hack.html).

The Olympic IT team was able to do this despite having virtually no understanding at first of how the malware, called Olympic Destroyer, worked. It wasn’t until several days after the attack that analysts began unraveling the origins of the worm, which seemed deliberately designed to send security researchers on a wild goose chase as they tried to analyze the code and identify its source.

The lesson for SREs: Preparation is golden. You can never know exactly what’s going to hit your systems, and in many cases, you won’t be able to identify root causes until well after you’re knee-deep in an outage. Nonetheless, by performing dry-runs and developing the right playbooks, you’ll position yourself to react effectively even in response to attacks of mystifying complexity.

## DNS outage strikes the Games

From an IT perspective, the Olympic Games currently taking place in Tokyo have gone pretty smoothly, despite the challenges caused by the pandemic that postponed them.

Nonetheless, a [temporary disruption to the Games’s website and app](https://www.nbcnewyork.com/news/widespread-outage-disrupts-major-retail-financial-travel-websites-worldwide/3168859/) just as Olympic events were getting underway raised early fears that things would not proceed so well. The incident, which also affected the websites of a variety of major retailers, stemmed from a problem with Akamai’s DNS network, which the company attributed to a bad software update.

Akamai didn’t release further details, but from the looks of things, this was an SRE 101 type of incident. Presumably, a bug somewhere in a software release eluded testing routines and made it into production.

The good news is that Akamai resolved the incident in about an hour. Did they perform a rollback or redirect traffic to backup infrastructure? We’ll probably never know, but what is clear is that they had a plan in place for responding quickly to one of the most common sources of IT disruptions: A bad application update. Thanks to their preparation, most Olympic viewers never even knew that an outage had occurred.

## Conclusion

Although most Games over the past two decades have witnessed some disruptions to their IT infrastructure, the teams responsible for managing reliability for the Olympics deserve a lot of credit. To date, no show-stopping outage has taken place. That’s a pretty good reliability scorecard when you’re dealing with the systems behind the largest, most-watched sporting event in the world.
