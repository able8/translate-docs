# (A few) Ops Lessons We All Learn The Hard Way

January 24th, 2020 From: https://www.netmeister.org/blog/ops-lessons.html

Nope, not another [Falsehoods](https://www.netmeister.org/blog/cs-falsehoods.html) post, but not entirely unlike one. Only here we have a few lessons in [operations](https://www.netmeister.org/blog/defining-operations.html) that we all (eventually) (have to) learn; often the hard way. Why things are the way they are, or what the lessons *mean* is left to the reader to interpret, agree, or disagree with. It's more fun that way. Enjoy! 

------

1. Email is the worst monitoring and alerting mechanism except for all the others.
2. Absence of a signal is itself a signal.
3. The severity of an incident is measured by the number of rules broken in resolving it.
4. The mobile hotspot you're paying for so you can leave your house while you're oncall only works at home and in the office.
5. The only other person who knows how this works is also on vacation.

6. If a post-mortem follow-up task is not picked up within a week, it's unlikely to be completed at all.
7. That janky script you put together during the outage -- the one that uses `expect(1)` and '`ssh -t -t`' -- now is the foundation of the entire team's toolchest.
8. NTP being off may not be a root cause, but it sure didn't help.
9. [UTC or GTFO](https://teespring.com/utc-or-gtfo).
10. Your infrastructure uses a lot more self-signed certificates than you think. A *lot* more. In places that make you weep.

11. Self-signed certificates beget long lived certs, which beget lack of certificate validity monitoring, which begets `curl -k`, which begets a lack of certificate deployment automation, which begets self-signed certificates.
12. For any N applications, at most N/2+1 use the same certificate bundle.
13. The system you're troubleshooting doesn't use the one the tool you're troubleshooting it with does.
14. An [API without a reference](https://twitter.com/jschauma/status/1086305575443542017) implementation and command-line client is called a gray box.
15. Restricted shells are not as restricted as you think.

16. Very few operations are truly idempotent.
17. "Asserting state" beats "monitoring for compliance" any day.
18. [One in a Million is next Tuesday.](https://docs.microsoft.com/en-us/archive/blogs/larryosterman/one-in-a-million-is-next-tuesday)
19. People give talks at conferences not to convince others that their work is awesome and totally worth the time and effort they put in, but themselves.
20. It's ok to use shell for complex stuff; it often times is easier, faster, and still less of a mess than juggling libraries and dependencies.

21. There's nothing wrong with Perl.
22. Ok, we all at times keep adding `$`, `{`, `}`, and `@` in random places trying to make things work, but still.
23. Serverless isn't.
24. [Y38K](https://en.wikipedia.org/wiki/Year_2038_problem) is [already here](https://twitter.com/jxxf/status/1219009308438024200), it's just not evenly distributed.
25. If you determine "[human error](https://www.netmeister.org/blog/humanerrno.html)" as the root cause, then you're doing it wrong.

26. Your network team has a way into the network that your security team doesn't know about.
27. And don't even as much as *mention* the serial console and IPMI networks, but boy are you glad you have 'em.
28. Blocking TCP port 53 traffic leads to very strange failures. Don't.
29. Somewhere in your infrastructure a service you didn't know uses DNS for endpoint discovery in a very surprising way.
30. Do. Not. Monkey. Around. With. `/etc/hosts`.

31. If you break it, you own it - for now; if you fix it, you own it - forever.
32. Turning it off and on again is actually quite a reasonable way to fix many things.
33. A `README.md` in git is no substitute for a manual page that's shipped with your tool.
34. A search for a document you know exists will only turn up links to documents referencing *but not actually linking to* the one you're looking for.
35. The document you're looking for was marked as obsolete and not migrated to the new content management solution.

36. Sure, your current content management system sucks, but it's still better than the one you're moving to.
37. Nobody knows how git works; everybody simply `rm -fr && git checkout`'s periodically.
38. There are very few network restrictions creative and determined use of `ssh(1)`port forwarding can't overcome.
39. This is both incredibly useful and concerning.
40. It is tempting to jump right into implementing a solution when the right thing may well be to not do the thing that requires the solution in the first place.

41. Turning things off permanently is [surprisingly difficult](https://twitter.com/jschauma/status/1074684924588974080).
42. "*Ancient*" is a very relative term when it comes to software and protocols.
43. "*Obsolete*" doesn't mean it's not in use and relied on.
44. The sets of systems online before and after a data center power outage only intersect. Some of the old systems coming online will immediately cause a different outage.
45. Some of your most critical services are kept alive by a handful of people whose job description does not mention those services at all.
46. After the initial "down for everybody or just me ermahgehrd Slack is down" drop, productivity increases linearly throughout the duration of the outage.
47. You're bound by the [CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem) much more often than you may think. [Halting Problem](https://en.wikipedia.org/wiki/Halting_problem)'s a bitch, too.
48. Eventual consistency doesn't help when the system you're debugging hasn't converged yet.
49. The source you're looking at is not the code running in production.
50. `strace(1)`/`ktrace(1)` doesn't lie.

51. Unless somebody's been playing `LD_PRELOAD` games.
52. Schrödinger's Backup -- "The condition of any backup is unknown until a restore is attempted." -- is overly optimistic.
53. There's [an xkcd](https://xkcd.com/305/) for the precise situation you find yourself in. (There's also one for at least half of these.)
54. At some point in your career you will implement [half of kerberos](https://twitter.com/jschauma/status/643444266300239873). Poorly.
55. Any sufficiently successful product launch is indistinguishable from a DDoS; any sufficiently advanced user indistinguishable from an attacker.

56. Debugging any sufficiently complex open source product is indistinguishable from reverse engineering a black box.
57. "We've always done it this way." is not a good reason by itself, but there's bound to be one for why.
58. That reason may or may not be valid any longer, however.
59. A junior engineer asking "why" and pointing out the docs don't reflect reality is at least as valuable as the senior engineer working blindly off tribal knowledge.
60. Your herculean efforts to upgrade the OS across your entire fleet completed just in time for the EOL announcement of the version you upgraded to.

61. This phenomenon was first described in Dante's *Inferno* as the Ninth Circle of Hell, Ring Four, aka `RedHat Canto XXXIV`.
62. Containers create at least as many problems as they solve.
63. The most ninja move the expert you hired for that third party black box product you rely on is to say "Let me ping the support team".
64. Somewhere, somebody ran into this *exact* problem, but they never bothered to post a solution.
65. That completely automated solution you set up requires at least three manual steps you didn't document.

66. CAPEX budget always increases, OPEX budget always decreases.
67. CAPEX costs can be reasonably estimated, OPEX costs can only be ballparked.
68. Doubling your time estimate in the hopes of beating expectations won't work because your manager takes your estimate, has a hardy laugh, and then resets it back to what they already promised upchain.
69. Your quarterly planning means bubkes when the next re-org rolls around.
70. Most of your actual work is not covered by your [OKRs](https://www.netmeister.org/blog/okr-distractions.html).

71. Recursively applying the [Pareto Principle](https://en.wikipedia.org/wiki/Pareto_principle) is a surprisingly accurate way to gauge your low hanging fruit, determine your high impact objectives, and ballpark your required effort.
72. Although, to be honest, it [only works in about 80% of cases](https://twitter.com/jschauma/status/962448995690967041).
73. Management will always happily [spend $$$ on outside consultants](https://www.netmeister.org/blog/crazy-like-a-fox.html) to tell them what you've been saying for years.
74. Management will much rather invest in inventing a new, square wheel than fixing an old round one.
75. In any organization practicing [continuous integration](https://twitter.com/jschauma/status/1225455309919006720), half of all commits are to fake out CI tests.

76. Good software development practices do not always translate well to ops and friends.
77. Mandatory [code reviews](https://twitter.com/jschauma/status/1019410471999467525) do not automatically improve code quality nor reduce the frequency of incidents.
78. Every new paradigm tends to mostly add layers of abstractions; cutting through them and identifying what basic principles continue to apply is half the battle.
79. Real change can only be implemented [above layer 7](https://en.wikipedia.org/wiki/Layer_8).
80. "Prod" is just another name for "staging".

81. Your source of truth lies.
82. Also: it's incomplete.
83. `pcap` or it didn't happen.
84. `grep(1)` > Splunk (there, I said it)
85. Multithreading is rarely worth the added complexity.
86. Parallelism is not Concurrency.
87. Simplicity is King.
88. Nobody knows what exactly it is you do.

January 24th, 2020